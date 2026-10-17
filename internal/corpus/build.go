package corpus

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pnathan/translations/internal/hebrew"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/oshb"
	"github.com/pnathan/translations/internal/work"
)

var (
	reKQ   = regexp.MustCompile(`<span class="mam-kq"><span class="mam-kq-k">\(([^)<]*)\)</span>\s*<span class="mam-kq-q">\[([^\]<]*)\]</span></span>`)
	rePara = regexp.MustCompile(`<span class="mam-spi-(pe|samekh)">\{[פס]\}</span>`)
	reHold = regexp.MustCompile(`^§(\d+)§$`)
	reHeld = regexp.MustCompile(`§(\d+)§`)
)

type mamWord struct {
	text, ketiv string
}

// mamWords splits one MAM verse into words. A ketiv/qere pair is one word:
// the qere is read, the ketiv kept.
func mamWords(h string) (words []mamWord, para string, plain string) {
	var kq [][2]string
	h = reKQ.ReplaceAllStringFunc(h, func(m string) string {
		sm := reKQ.FindStringSubmatch(m)
		kq = append(kq, [2]string{sm[1], sm[2]})
		return fmt.Sprintf(" §%d§ ", len(kq)-1)
	})
	if m := rePara.FindStringSubmatch(h); m != nil {
		para = map[string]string{"pe": "פ", "samekh": "ס"}[m[1]]
	}
	h = rePara.ReplaceAllString(h, "")
	h = hebrew.StripHTML(h)
	for _, f := range strings.Fields(h) {
		if m := reHold.FindStringSubmatch(f); m != nil {
			var i int
			fmt.Sscan(m[1], &i)
			words = append(words, mamWord{text: kq[i][1], ketiv: kq[i][0]})
			continue
		}
		for _, w := range hebrew.Words(f) {
			words = append(words, mamWord{text: w})
		}
	}
	plain = reHeld.ReplaceAllStringFunc(strings.Join(strings.Fields(h), " "), func(m string) string {
		var i int
		fmt.Sscan(strings.Trim(m, "§"), &i)
		return kq[i][1]
	})
	return words, para, plain
}

// Alignment keys ignore the vowel letters, so plene and defective
// spellings of one word align; the difference is then recorded.
func (w mamWord) key() string {
	if w.ketiv != "" {
		return hebrew.Loose(w.ketiv)
	}
	return hebrew.Loose(w.text)
}

func oshbKey(w oshb.Word) string { return hebrew.Loose(w.Text) }

// lcs aligns two key sequences; it returns pairs (i, j) of matched indices.
func lcs(a, b []string) [][2]int {
	n, m := len(a), len(b)
	t := make([][]int, n+1)
	for i := range t {
		t[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i] == b[j] {
				t[i][j] = t[i+1][j+1] + 1
			} else {
				t[i][j] = max(t[i+1][j], t[i][j+1])
			}
		}
	}
	var out [][2]int
	for i, j := 0, 0; i < n && j < m; {
		switch {
		case a[i] == b[j]:
			out = append(out, [2]int{i, j})
			i++
			j++
		case t[i+1][j] >= t[i][j+1]:
			i++
		default:
			j++
		}
	}
	return out
}

func vowels(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r == 'ׇ':
			b.WriteRune('ָ')
		case r >= 'ְ' && r <= 'ּ', r == 'ׁ', r == 'ׂ':
			b.WriteRune(r)
		case hebrew.IsConsonant(r):
			b.WriteRune(r)
		}
	}
	return b.String()
}

// BuildTanakh joins the MAM text (Sefaria) with the OSHB words and the
// lexicon. mam is chapter → verse → HTML string.
func BuildTanakh(cfg *work.Config, mam [][]string, verses []oshb.Verse, ix *lex.Index) (*Corpus, map[string]*lex.Entry, error) {
	byID := map[string]oshb.Verse{}
	for _, v := range verses {
		byID[v.OsisID] = v
	}
	lexicon := map[string]*lex.Entry{}
	c := &Corpus{Work: cfg.Slug}
	for ci, ch := range mam {
		for vi, h := range ch {
			ref := fmt.Sprintf("%d:%d", ci+1, vi+1)
			sec := cfg.SectionOf(ref)
			if sec == nil {
				continue
			}
			ov, ok := byID[fmt.Sprintf("%s.%d.%d", cfg.OSIS, ci+1, vi+1)]
			if !ok {
				return nil, nil, fmt.Errorf("%s: no OSHB verse", ref)
			}
			mw, para, plain := mamWords(h)
			seg := Segment{Ref: ref, Section: sec.Slug, Text: plain, Para: para}
			ak := make([]string, len(ov.Words))
			for i, w := range ov.Words {
				ak[i] = oshbKey(w)
			}
			bk := make([]string, len(mw))
			for i, w := range mw {
				bk[i] = w.key()
			}
			pairs := lcs(ak, bk)
			matched := map[int]int{}
			for _, p := range pairs {
				matched[p[0]] = p[1]
			}
			usedB := map[int]bool{}
			for _, p := range pairs {
				usedB[p[1]] = true
			}
			for i, w := range ov.Words {
				id := fmt.Sprintf("%s.%d", ref, i+1)
				t := Token{ID: id, OSHB: w.ID, Maqaf: w.MaqafAfter, Notes: w.Notes, MorphSource: "OSHB"}
				wlc := strings.ReplaceAll(w.Text, "/", "")
				if w.Ketiv && w.Qere != nil {
					t.Ketiv = wlc
					wlc = strings.ReplaceAll(w.Qere.Text, "/", "")
				}
				if j, ok := matched[i]; ok {
					t.Text = mw[j].text
					if mw[j].ketiv != "" && t.Ketiv == "" {
						t.Ketiv = mw[j].ketiv
					}
					lemma := w.Lemma
					if mw[j].ketiv == "" && hebrew.Skeleton(t.Text) != hebrew.Skeleton(wlc) {
						seg.Variants = append(seg.Variants, Variant{Tokens: []string{id}, Witness: "WLC", Reading: wlc, Base: t.Text, Kind: "spelling"})
					} else if !strings.HasPrefix(lemma, "3068") && !strings.HasPrefix(lemma, "3069") && vowels(t.Text) != vowels(wlc) {
						seg.Variants = append(seg.Variants, Variant{Tokens: []string{id}, Witness: "WLC", Reading: wlc, Base: t.Text, Kind: "vowels"})
					}
				} else {
					t.Text = wlc
					seg.Variants = append(seg.Variants, Variant{Tokens: []string{id}, Witness: "MAM", Reading: "—", Base: wlc, Kind: "word"})
				}
				if w.Ketiv {
					seg.Variants = append(seg.Variants, Variant{Tokens: []string{id}, Witness: "ketiv", Reading: t.Ketiv, Base: t.Text, Kind: "ketiv-qere"})
				}
				fillOSHB(&t, w, lexicon, ix)
				seg.Tokens = append(seg.Tokens, t)
			}
			for j, w := range mw {
				if !usedB[j] {
					seg.Variants = append(seg.Variants, Variant{Witness: "WLC", Reading: "—", Base: w.text, Kind: "word"})
				}
			}
			if seg.Para == "" {
				seg.Para = ov.Para
			}
			for i := range seg.Tokens {
				acc := hebrew.Accents(seg.Tokens[i].Text, i == len(seg.Tokens)-1)
				for _, a := range acc {
					seg.Tokens[i].Accents = append(seg.Tokens[i].Accents, a.Name)
					if a.Disjunctive {
						seg.Tokens[i].Disjunctive = true
					}
				}
			}
			c.Segments = append(c.Segments, seg)
		}
	}
	c.Units = makeUnits(cfg, c.Segments)
	return c, lexicon, nil
}

// fillOSHB decodes the lemma and morphology of an OSHB word into morphemes.
func fillOSHB(t *Token, w oshb.Word, lexicon map[string]*lex.Entry, ix *lex.Index) {
	src := w
	if w.Ketiv && w.Qere != nil {
		src = *w.Qere
	}
	lang, ms := oshb.Decode(src.Morph)
	t.Lang = lang
	forms := strings.Split(src.Text, "/")
	lemmas := strings.Split(src.Lemma, "/")
	li := 0
	for k, m := range ms {
		mo := Morpheme{POS: m.POS, Parse: m.Parse, Code: m.Code}
		if k < len(forms) {
			mo.Form = forms[k]
		}
		if m.POS != "suffix" && li < len(lemmas) {
			l := lemmas[li]
			li++
			if g, ok := oshb.PrefixLemma[l]; ok {
				mo.Lemma = g
			} else if e, ok := ix.Lookup(l); ok {
				lexicon[e.Key] = e
				mo.Lemma = e.Headword
				mo.Lex = e.Key
				if t.Lex == "" {
					t.Lex = e.Key
					t.Root = e.Root
				}
			} else {
				mo.Lemma = "Strong " + l
			}
		}
		t.Morphemes = append(t.Morphemes, mo)
	}
}

func makeUnits(cfg *work.Config, segs []Segment) []Unit {
	var units []Unit
	count := map[string]int{}
	for i := range segs {
		s := &segs[i]
		var prev *Segment
		if i > 0 {
			prev = &segs[i-1]
		}
		start := len(units) == 0
		if !start {
			u := &units[len(units)-1]
			n := len(u.Refs)
			pch, _, _ := work.ParseRef(prev.Ref)
			sch, _, _ := work.ParseRef(s.Ref)
			switch {
			case u.Section != s.Section:
				start = true
			case n >= cfg.UnitMax:
				start = true
			case cfg.Kind == work.Tanakh && prev.Para != "" && n >= cfg.UnitMax/2:
				start = true
			case cfg.Kind == work.Mishnah && pch != sch:
				start = true
			}
		}
		if start {
			count[s.Section]++
			units = append(units, Unit{ID: fmt.Sprintf("%s-%02d", s.Section, count[s.Section]), Section: s.Section})
		}
		u := &units[len(units)-1]
		u.Refs = append(u.Refs, s.Ref)
		s.Unit = u.ID
	}
	return units
}

var reMishWord = regexp.MustCompile(`^([^\p{Hebrew}]*)(.*?)([^\p{Hebrew}\x{0591}-\x{05C7}]*)$`)

// mishnahTokens splits a mishnah into words and the punctuation after them.
func mishnahTokens(text string) []Token {
	var out []Token
	for _, f := range strings.Fields(text) {
		m := reMishWord.FindStringSubmatch(f)
		core, punct := f, ""
		if m != nil {
			core, punct = m[2], m[3]
		}
		if hebrew.Skeleton(core) == "" {
			if len(out) > 0 {
				out[len(out)-1].Punct += f
			}
			continue
		}
		out = append(out, Token{Text: core, Punct: punct})
	}
	return out
}

// Analysis is the model's morphology of one mishnah token.
type Analysis struct {
	ID        string     `json:"id"`
	Lang      string     `json:"lang"`
	Root      string     `json:"root,omitempty"`
	Morphemes []Morpheme `json:"morphemes"`
}

// BuildMishnah reads the base text and the collated witnesses. morph holds
// the model's analyses by token id; it may be empty before that stage.
func BuildMishnah(cfg *work.Config, base [][]string, witnesses map[string][][]string, morph map[string]Analysis, morphSource map[string]string) *Corpus {
	c := &Corpus{Work: cfg.Slug}
	for ci, ch := range base {
		for mi, h := range ch {
			ref := fmt.Sprintf("%d:%d", ci+1, mi+1)
			sec := cfg.SectionOf(ref)
			if sec == nil {
				continue
			}
			plain := hebrew.StripHTML(h)
			seg := Segment{Ref: ref, Section: sec.Slug, Text: plain}
			seg.Tokens = mishnahTokens(plain)
			for i := range seg.Tokens {
				t := &seg.Tokens[i]
				t.ID = fmt.Sprintf("%s.%d", ref, i+1)
				t.Lang = "mishnaic"
				if a, ok := morph[t.ID]; ok {
					t.Morphemes = a.Morphemes
					t.Root = a.Root
					if a.Lang != "" {
						t.Lang = a.Lang
					}
					t.MorphSource = morphSource[t.ID]
					for _, m := range a.Morphemes {
						if m.Lex != "" && t.Lex == "" {
							t.Lex = m.Lex
						}
					}
				}
			}
			for _, name := range sortedKeys(witnesses) {
				w := witnesses[name]
				if ci >= len(w) || mi >= len(w[ci]) {
					seg.Variants = append(seg.Variants, Variant{Witness: name, Reading: "—", Base: plain, Kind: "segment"})
					continue
				}
				collate(&seg, name, mishnahTokens(hebrew.StripHTML(w[ci][mi])))
			}
			c.Segments = append(c.Segments, seg)
		}
	}
	c.Units = makeUnits(cfg, c.Segments)
	return c
}

// collate records the words where a witness differs from the base text.
func collate(seg *Segment, name string, wt []Token) {
	ak := make([]string, len(seg.Tokens))
	for i, t := range seg.Tokens {
		ak[i] = hebrew.Loose(t.Text)
	}
	bk := make([]string, len(wt))
	for i, t := range wt {
		bk[i] = hebrew.Loose(t.Text)
	}
	pairs := lcs(ak, bk)
	pairs = append(pairs, [2]int{len(ak), len(bk)})
	i, j := 0, 0
	for _, p := range pairs {
		if p[0] > i || p[1] > j {
			var ids, base, read []string
			for k := i; k < p[0]; k++ {
				ids = append(ids, seg.Tokens[k].ID)
				base = append(base, seg.Tokens[k].Text)
			}
			for k := j; k < p[1]; k++ {
				read = append(read, wt[k].Text)
			}
			r := strings.Join(read, " ")
			if r == "" {
				r = "—"
			}
			b := strings.Join(base, " ")
			if b == "" {
				b = "—"
			}
			seg.Variants = append(seg.Variants, Variant{Tokens: ids, Witness: name, Reading: r, Base: b, Kind: "word"})
		}
		if p[0] < len(ak) {
			if hebrew.Skeleton(seg.Tokens[p[0]].Text) != hebrew.Skeleton(wt[p[1]].Text) {
				seg.Tokens[p[0]].Alt = wt[p[1]].Text
				seg.Variants = append(seg.Variants, Variant{Tokens: []string{seg.Tokens[p[0]].ID}, Witness: name, Reading: wt[p[1]].Text, Base: seg.Tokens[p[0]].Text, Kind: "spelling"})
			}
		}
		i, j = p[0]+1, p[1]+1
	}
}

func sortedKeys[V any](m map[string]V) []string {
	var k []string
	for x := range m {
		k = append(k, x)
	}
	for a := 1; a < len(k); a++ {
		for b := a; b > 0 && k[b] < k[b-1]; b-- {
			k[b], k[b-1] = k[b-1], k[b]
		}
	}
	return k
}
