package llm

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/pnathan/translations/internal/anchor"
	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/hebrew"
	"github.com/pnathan/translations/internal/lex"
)

// Env is what the checks read.
type Env struct {
	Corpus  *corpus.Corpus
	Lexicon map[string]*lex.Entry
	Anchors []anchor.Anchor
}

func (e *Env) anchorsOf(ref string) []anchor.Anchor {
	var out []anchor.Anchor
	for _, a := range e.Anchors {
		if a.Ref == ref {
			out = append(out, a)
		}
	}
	return out
}

func (e *Env) anchor(id string) *anchor.Anchor {
	for i := range e.Anchors {
		if e.Anchors[i].ID == id {
			return &e.Anchors[i]
		}
	}
	return nil
}

var POS = map[string]bool{
	"noun": true, "verb": true, "adjective": true, "adverb": true, "pronoun": true,
	"preposition": true, "conjunction": true, "particle": true, "suffix": true, "article": true,
	"numeral": true, "interjection": true,
}

// CheckMorph validates a morphology unit.
func (e *Env) CheckMorph(mu *MorphUnit) []string {
	var errs []string
	u := e.Corpus.Unit(mu.Unit)
	if u == nil {
		return []string{"unknown unit " + mu.Unit}
	}
	if mu.Model == "" {
		errs = append(errs, "model is empty")
	}
	want := map[string]corpus.Token{}
	for _, r := range u.Refs {
		for _, t := range e.Corpus.Segment(r).Tokens {
			want[t.ID] = t
		}
	}
	seen := map[string]bool{}
	for _, a := range mu.Tokens {
		t, ok := want[a.ID]
		if !ok {
			errs = append(errs, a.ID+": not a token of this unit")
			continue
		}
		if seen[a.ID] {
			errs = append(errs, a.ID+": analysed twice")
		}
		seen[a.ID] = true
		errs = append(errs, checkAnalysis(t, a)...)
	}
	for id := range want {
		if !seen[id] {
			errs = append(errs, id+": no analysis")
		}
	}
	return errs
}

func checkAnalysis(t corpus.Token, a corpus.Analysis) []string {
	var errs []string
	if len(a.Morphemes) == 0 {
		return []string{a.ID + ": no morphemes"}
	}
	var forms strings.Builder
	for i, m := range a.Morphemes {
		forms.WriteString(m.Form)
		if !POS[m.POS] {
			errs = append(errs, fmt.Sprintf("%s: morpheme %d: pos %q not in the list", a.ID, i+1, m.POS))
		}
		if m.Lemma == "" && m.POS != "suffix" {
			errs = append(errs, fmt.Sprintf("%s: morpheme %d: no lemma", a.ID, i+1))
		}
		if m.Parse == "" {
			errs = append(errs, fmt.Sprintf("%s: morpheme %d: no parse", a.ID, i+1))
		}
	}
	if hebrew.Skeleton(forms.String()) != hebrew.Skeleton(t.Text) {
		errs = append(errs, fmt.Sprintf("%s: morpheme forms %q do not spell the token %q", a.ID, forms.String(), t.Text))
	}
	switch a.Lang {
	case "mishnaic", "hbo", "arc":
	default:
		errs = append(errs, fmt.Sprintf("%s: lang %q: want mishnaic, hbo or arc", a.ID, a.Lang))
	}
	return errs
}

var rePunct = regexp.MustCompile(`[^\p{L}\p{N}' ]+`)

func norm(s string) string {
	s = strings.ToLower(strings.ReplaceAll(s, "’", "'"))
	s = rePunct.ReplaceAllString(s, " ")
	return strings.Join(strings.Fields(s), " ")
}

func hasHebrew(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Hebrew, r) {
			return true
		}
	}
	return false
}

var relations = map[string]bool{"follows": true, "records": true, "declines": true, "not-applicable": true}

// senseOK reports whether path names a sense of entry key.
func (e *Env) senseOK(key, path string) bool {
	en := e.Lexicon[key]
	if en == nil {
		return false
	}
	for _, s := range en.Senses {
		if s.Path == path {
			return true
		}
	}
	return false
}

// CheckSegment validates one translated segment.
func (e *Env) CheckSegment(st *SegmentTr) []string {
	seg := e.Corpus.Segment(st.Ref)
	if seg == nil {
		return []string{st.Ref + ": unknown segment"}
	}
	var errs []string
	add := func(f string, a ...any) { errs = append(errs, st.Ref+": "+fmt.Sprintf(f, a...)) }
	if strings.TrimSpace(st.English) == "" {
		add("english is empty")
	}
	if hasHebrew(st.English) {
		add("english contains Hebrew letters")
	}
	if strings.TrimSpace(st.Literal) == "" {
		add("literal is empty")
	}
	if len(st.Grammar) == 0 {
		add("grammar has no learner notes")
	}
	tokIdx := map[string]int{}
	for i, t := range seg.Tokens {
		tokIdx[t.ID] = i
	}
	seen := map[string]bool{}
	for _, g := range st.Tokens {
		i, ok := tokIdx[g.ID]
		if !ok {
			add("token %s is not in this segment", g.ID)
			continue
		}
		if seen[g.ID] {
			add("token %s glossed twice", g.ID)
		}
		seen[g.ID] = true
		if strings.TrimSpace(g.Gloss) == "" {
			add("token %s: gloss is empty", g.ID)
		}
		key := g.Lex
		if key == "" {
			key = seg.Tokens[i].Lex
		}
		if g.Sense != "" && !e.senseOK(key, g.Sense) {
			add("token %s: sense %q is not a sense of %s", g.ID, g.Sense, key)
		}
		if g.Lex != "" && e.Lexicon[g.Lex] == nil {
			add("token %s: lexicon key %q unknown", g.ID, g.Lex)
		}
	}
	for _, t := range seg.Tokens {
		if !seen[t.ID] {
			add("token %s has no gloss", t.ID)
		}
	}
	// Phrases cover every token once, in order, and quote the English.
	next := 0
	english := norm(st.English)
	for k, p := range st.Phrases {
		if len(p.Tokens) == 0 {
			add("phrase %d has no tokens", k+1)
			continue
		}
		for _, id := range p.Tokens {
			i, ok := tokIdx[id]
			if !ok {
				add("phrase %d: token %s not in this segment", k+1, id)
				continue
			}
			if i != next {
				add("phrase %d: token %s out of order (want %s)", k+1, id, seg.Tokens[min(next, len(seg.Tokens)-1)].ID)
			}
			next = i + 1
		}
		if p.English != "-" && !strings.Contains(english, norm(p.English)) {
			add("phrase %d: %q does not occur in the english", k+1, p.English)
		}
		if strings.TrimSpace(p.Literal) == "" {
			add("phrase %d: literal is empty", k+1)
		}
	}
	if next != len(seg.Tokens) {
		add("phrases cover %d of %d tokens", next, len(seg.Tokens))
	}
	// Commentary: every normative comment anchored here is ruled on.
	used := map[string]bool{}
	for _, c := range st.Commentary {
		a := e.anchor(c.Anchor)
		if a == nil {
			add("commentary: unknown anchor %s", c.Anchor)
			continue
		}
		if a.Ref != st.Ref {
			add("commentary: anchor %s belongs to %s", c.Anchor, a.Ref)
		}
		if used[c.Anchor] {
			add("commentary: anchor %s twice", c.Anchor)
		}
		used[c.Anchor] = true
		if !relations[c.Relation] {
			add("commentary %s: relation %q not in follows|records|declines|not-applicable", c.Anchor, c.Relation)
		}
		if c.Relation != "not-applicable" && strings.TrimSpace(c.Summary) == "" {
			add("commentary %s: summary is empty", c.Anchor)
		}
	}
	for _, a := range e.anchorsOf(st.Ref) {
		if a.Normative && !used[a.ID] {
			add("commentary: normative anchor %s (%s) not ruled on", a.ID, a.Commentator)
		}
	}
	return errs
}

// CheckTranslation validates a translation unit.
func (e *Env) CheckTranslation(tu *TranslationUnit) []string {
	u := e.Corpus.Unit(tu.Unit)
	if u == nil {
		return []string{"unknown unit " + tu.Unit}
	}
	var errs []string
	if tu.Model == "" {
		errs = append(errs, "model is empty")
	}
	if len(tu.Segments) != len(u.Refs) {
		errs = append(errs, fmt.Sprintf("%d segments, want %d (%s)", len(tu.Segments), len(u.Refs), strings.Join(u.Refs, " ")))
	}
	for i := range tu.Segments {
		if i < len(u.Refs) && tu.Segments[i].Ref != u.Refs[i] {
			errs = append(errs, fmt.Sprintf("segment %d is %s, want %s", i+1, tu.Segments[i].Ref, u.Refs[i]))
		}
		errs = append(errs, e.CheckSegment(&tu.Segments[i])...)
	}
	return errs
}

var verdicts = map[string]bool{"consistent": true, "contradicts": true, "not-applicable": true}

// CheckReview validates a review against its translation: the patches apply
// and the result passes the translation check; every normative anchor has a
// verdict.
func (e *Env) CheckReview(ru *ReviewUnit, tu *TranslationUnit, mu *MorphUnit) []string {
	var errs []string
	if ru.Model == "" {
		errs = append(errs, "model is empty")
	}
	if strings.TrimSpace(ru.Summary) == "" {
		errs = append(errs, "summary is empty")
	}
	u := e.Corpus.Unit(ru.Unit)
	if u == nil {
		return append(errs, "unknown unit "+ru.Unit)
	}
	got := map[string]bool{}
	for _, v := range ru.Verdicts {
		if !verdicts[v.Verdict] {
			errs = append(errs, fmt.Sprintf("verdict %s: %q not in consistent|contradicts|not-applicable", v.Anchor, v.Verdict))
		}
		if e.anchor(v.Anchor) == nil {
			errs = append(errs, "verdict: unknown anchor "+v.Anchor)
		}
		got[v.Anchor] = true
	}
	for _, r := range u.Refs {
		for _, a := range e.anchorsOf(r) {
			if a.Normative && !got[a.ID] {
				errs = append(errs, fmt.Sprintf("no verdict on normative anchor %s (%s)", a.ID, a.Commentator))
			}
		}
	}
	for _, p := range ru.Patches {
		switch p.Severity {
		case "error", "question", "style":
		default:
			errs = append(errs, fmt.Sprintf("patch %s %s: severity %q", p.Ref, p.Path, p.Severity))
		}
		if strings.TrimSpace(p.Reason) == "" {
			errs = append(errs, fmt.Sprintf("patch %s %s: no reason", p.Ref, p.Path))
		}
	}
	final, err := Apply(tu, ru)
	if err != nil {
		return append(errs, err.Error())
	}
	for _, x := range e.CheckTranslation(final) {
		errs = append(errs, "after patches: "+x)
	}
	if mu != nil {
		toks := map[string]corpus.Token{}
		for _, r := range u.Refs {
			for _, t := range e.Corpus.Segment(r).Tokens {
				toks[t.ID] = t
			}
		}
		for _, p := range ru.Patches {
			id, ok := strings.CutPrefix(p.Path, "morph:")
			if !ok {
				continue
			}
			var a corpus.Analysis
			if err := json.Unmarshal(p.Value, &a); err != nil {
				errs = append(errs, fmt.Sprintf("patch %s: %v", p.Path, err))
				continue
			}
			a.ID = id
			t, ok := toks[id]
			if !ok {
				errs = append(errs, "patch "+p.Path+": not a token of this unit")
				continue
			}
			errs = append(errs, checkAnalysis(t, a)...)
		}
	}
	return errs
}
