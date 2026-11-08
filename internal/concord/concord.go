// Package concord builds the concordance of the finished works: every root,
// its lemmas, each lemma's senses, and every token of the texts that cites
// the lemma, with the gloss and the sense the reviewed translation chose
// there. It is derived from the corpus, the lexicon and the final
// translations; nothing in it is written by a model.
package concord

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/hebrew"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/work"
)

// Source is one work as the concordance reads it.
type Source struct {
	Cfg     *work.Config
	Corpus  *corpus.Corpus
	Lexicon map[string]*lex.Entry
	Final   map[string]*llm.SegmentTr // by ref; segments without one are skipped
}

// Load reads a work's corpus, lexicon and final translations.
func Load(cfg *work.Config) (*Source, error) {
	c, err := corpus.Load(cfg.Path("stage", "corpus.json"))
	if err != nil {
		return nil, err
	}
	s := &Source{Cfg: cfg, Corpus: c, Final: map[string]*llm.SegmentTr{}}
	if err := corpus.ReadJSON(cfg.Path("stage", "lexicon.json"), &s.Lexicon); err != nil {
		return nil, err
	}
	files, _ := filepath.Glob(llm.FinalPath(cfg, "*"))
	for _, f := range files {
		var tu llm.TranslationUnit
		if err := corpus.ReadJSON(f, &tu); err != nil {
			return nil, err
		}
		for i := range tu.Segments {
			s.Final[tu.Segments[i].Ref] = &tu.Segments[i]
		}
	}
	return s, nil
}

// Occurrence is one token that cites a lemma.
type Occurrence struct {
	Work    string `json:"w"`
	Section string `json:"sec"`
	Ref     string `json:"ref"`
	Token   string `json:"id"`
	Form    string `json:"f"`
	Gloss   string `json:"g,omitempty"`
	Sense   string `json:"s,omitempty"`
}

// Lemma is one lexicon entry and the tokens that cite it.
type Lemma struct {
	Key      string       `json:"key"`
	Headword string       `json:"hw"`
	POS      string       `json:"pos,omitempty"`
	Gloss    string       `json:"gloss,omitempty"`
	Source   string       `json:"src,omitempty"`
	Senses   []lex.Sense  `json:"senses,omitempty"`
	Note     string       `json:"note,omitempty"`
	Occ      []Occurrence `json:"occ"`
}

// Root groups the lemmas of one root. A lemma without a root is a group of
// its own, with Root empty.
type Root struct {
	Key    string   `json:"key"`  // join key: consonants, shin and sin dots kept
	Root   string   `json:"root"` // as the sources write it; "" for a lemma without a root
	Note   string   `json:"note,omitempty"`
	Lemmas []*Lemma `json:"lemmas"`
	Count  int      `json:"n"`
}

// Concordance is the joined index of all the works.
type Concordance struct {
	Roots []*Root
	// ByToken names the root group of every indexed token: "work:id" → index into Roots.
	ByToken map[string]int
}

const (
	shinDot = 'ׁ'
	sinDot  = 'ׂ'
)

// rootKey keeps the consonants (final forms as medial) and the dots that
// tell שׁ from שׂ; nothing else of the root as written.
func rootKey(r string) string {
	var b strings.Builder
	for _, c := range r {
		switch {
		case hebrew.IsConsonant(c):
			b.WriteString(hebrew.Skeleton(string(c)))
		case c == shinDot || c == sinDot:
			b.WriteRune(c)
		}
	}
	return b.String()
}

// hasUndotted reports whether a key has a ש without its dot.
func hasUndotted(k string) bool {
	rs := []rune(k)
	for i, c := range rs {
		if c == 'ש' && (i+1 == len(rs) || rs[i+1] != shinDot && rs[i+1] != sinDot) {
			return true
		}
	}
	return false
}

// dotFromHeadword dots the ש of an undotted root with the dot of the
// lemma's pointed headword (עשר, עֲשָׂרָה → עשׂר). The two must have the same
// number of ש and every ש of the headword the same dot; else it returns "".
func dotFromHeadword(root, head string) string {
	n := strings.Count(root, "ש")
	if n == 0 || strings.Count(head, "ש") != n {
		return ""
	}
	var dot rune
	rs := []rune(head)
	for i, c := range rs {
		if c != 'ש' {
			continue
		}
		var d rune
		for _, m := range rs[i+1:] {
			if m == shinDot || m == sinDot {
				d = m
				break
			}
			if !hebrew.IsMark(m) {
				break
			}
		}
		if d == 0 || dot != 0 && d != dot {
			return ""
		}
		dot = d
	}
	return strings.ReplaceAll(root, "ש", "ש"+string(dot))
}

type tokRef struct {
	work, id string
}

// Build joins the works. A root written without the dot on ש (the model
// roots of the Mishnah) takes the dot of its lemma's pointed headword;
// failing that, it joins the dotted root with its consonants when there is
// exactly one; when there are several (שׁבע and שׂבע) it stays a group of its
// own, marked undetermined.
func Build(srcs []*Source) *Concordance {
	type lemmaAcc struct {
		lem   *Lemma
		roots map[string]int // root as written → tokens
		root  string         // the root most tokens give
		key   string
	}
	lemmas := map[string]*lemmaAcc{} // work-independent lexicon key → lemma
	var order []string
	byTok := map[tokRef]string{} // token → lemma key
	for _, s := range srcs {
		for _, seg := range s.Corpus.Segments {
			f := s.Final[seg.Ref]
			if f == nil {
				continue
			}
			gl := map[string]*llm.TokenGloss{}
			for i := range f.Tokens {
				gl[f.Tokens[i].ID] = &f.Tokens[i]
			}
			for _, t := range seg.Tokens {
				g := gl[t.ID]
				key := t.Lex
				if g != nil && g.Lex != "" {
					key = g.Lex // the reviewed translation repointed the entry
				}
				e := s.Lexicon[key]
				if key == "" || e == nil {
					continue
				}
				// The token's root, unless the review repointed the entry
				// to one that names its own.
				root := t.Root
				if key != t.Lex && e.Root != "" || root == "" {
					root = e.Root
				}
				a := lemmas[key]
				if a == nil {
					a = &lemmaAcc{lem: &Lemma{Key: key, Headword: e.Headword, POS: e.POS, Gloss: e.Gloss, Source: e.Source, Senses: e.Senses}, roots: map[string]int{}}
					lemmas[key] = a
					order = append(order, key)
				}
				if root != "" {
					a.roots[root]++
				}
				o := Occurrence{Work: s.Cfg.Slug, Section: seg.Section, Ref: seg.Ref, Token: t.ID, Form: t.Text}
				if g != nil {
					o.Gloss, o.Sense = g.Gloss, g.Sense
				}
				a.lem.Occ = append(a.lem.Occ, o)
				byTok[tokRef{s.Cfg.Slug, t.ID}] = key
			}
		}
	}
	// One root per lemma: the one most of its tokens give. The others are
	// named in the group's note.
	notes := map[string]string{}
	for _, k := range order {
		a := lemmas[k]
		var rs []string
		for r := range a.roots {
			rs = append(rs, r)
		}
		sort.Slice(rs, func(i, j int) bool {
			if a.roots[rs[i]] != a.roots[rs[j]] {
				return a.roots[rs[i]] > a.roots[rs[j]]
			}
			return rs[i] < rs[j]
		})
		if len(rs) > 0 {
			a.root = rs[0]
		}
		if len(rs) > 1 {
			a.lem.Note = "tokens also give the root " + strings.Join(rs[1:], ", ")
		}
	}
	// Resolve the undotted keys.
	dotted := map[string][]string{} // skeleton → dotted keys
	for _, k := range order {
		a := lemmas[k]
		a.key = rootKey(a.root)
		if a.key != "" && !hasUndotted(a.key) && strings.ContainsAny(a.key, "ׁׂ") {
			sk := hebrew.Skeleton(a.key)
			if !contains(dotted[sk], a.key) {
				dotted[sk] = append(dotted[sk], a.key)
			}
		}
	}
	for _, k := range order {
		a := lemmas[k]
		if a.key == "" || !hasUndotted(a.key) {
			continue
		}
		if dr := dotFromHeadword(a.root, a.lem.Headword); dr != "" {
			a.root, a.key = dr, rootKey(dr)
			continue
		}
		switch d := dotted[hebrew.Skeleton(a.key)]; len(d) {
		case 0:
		case 1:
			a.key = d[0]
		default:
			notes[a.key] = "shin or sin undetermined: " + strings.Join(d, ", ")
		}
	}
	groups := map[string]*Root{}
	var roots []*Root
	for _, k := range order {
		a := lemmas[k]
		gk := a.key
		if gk == "" {
			gk = "lemma:" + k
		}
		r := groups[gk]
		if r == nil {
			r = &Root{Key: gk, Note: notes[gk]}
			if a.key != "" {
				r.Root = a.root
			}
			groups[gk] = r
			roots = append(roots, r)
		}
		// Prefer the dotted spelling for display.
		if a.key != "" && strings.ContainsAny(a.root, "ׁׂ") && !strings.ContainsAny(r.Root, "ׁׂ") {
			r.Root = a.root
		}
		r.Lemmas = append(r.Lemmas, a.lem)
		r.Count += len(a.lem.Occ)
	}
	for _, r := range roots {
		sort.SliceStable(r.Lemmas, func(i, j int) bool { return len(r.Lemmas[i].Occ) > len(r.Lemmas[j].Occ) })
	}
	sort.SliceStable(roots, func(i, j int) bool { return lessHe(sortKey(roots[i]), sortKey(roots[j])) })
	c := &Concordance{Roots: roots, ByToken: map[string]int{}}
	idx := map[string]int{}
	for i, r := range roots {
		idx[r.Key] = i
	}
	for tr, lk := range byTok {
		a := lemmas[lk]
		gk := a.key
		if gk == "" {
			gk = "lemma:" + lk
		}
		c.ByToken[tr.work+":"+tr.id] = idx[gk]
	}
	return c
}

func contains(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}

// sortKey orders a group by its root, or by its lemma's consonants.
func sortKey(r *Root) string {
	if r.Root != "" {
		return hebrew.Skeleton(r.Root)
	}
	return hebrew.Skeleton(r.Lemmas[0].Headword)
}

// lessHe orders Hebrew strings in alphabetical order (the code points of
// the medial letters are in alphabetical order once finals are mapped).
func lessHe(a, b string) bool { return a < b }

// Letter is the first consonant of a group, final forms as medial.
func (r *Root) Letter() string {
	k := []rune(sortKey(r))
	if len(k) == 0 {
		return ""
	}
	return string(k[0])
}
