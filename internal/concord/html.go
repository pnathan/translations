package concord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/lex"
)

var letterSlug = map[string]string{
	"א": "alef", "ב": "bet", "ג": "gimel", "ד": "dalet", "ה": "he", "ו": "vav", "ז": "zayin",
	"ח": "het", "ט": "tet", "י": "yod", "כ": "kaf", "ל": "lamed", "מ": "mem", "נ": "nun",
	"ס": "samekh", "ע": "ayin", "פ": "pe", "צ": "tsadi", "ק": "qof", "ר": "resh", "ש": "shin", "ת": "tav",
}

// letters in alphabetical order
var letters = []string{"א", "ב", "ג", "ד", "ה", "ו", "ז", "ח", "ט", "י", "כ", "ל", "מ", "נ", "ס", "ע", "פ", "צ", "ק", "ר", "ש", "ת"}

// Work names a work for links and labels.
type Work struct {
	Slug, Title string
}

type senseGroup struct {
	Path, Gloss string
	Occ         []Occurrence
}

type lemmaView struct {
	*Lemma
	Groups []senseGroup
}

type rootView struct {
	R       *Root
	Index   int
	Views   []lemmaView
	Works   string
	Glosses string
}

type letterView struct {
	Letter, Slug string
	Roots        []rootView
	N            int
}

func senseGroups(l *Lemma) []senseGroup {
	gloss := map[string]string{}
	var order []string
	for _, s := range l.Senses {
		gloss[s.Path] = s.Gloss
		order = append(order, s.Path)
	}
	by := map[string][]Occurrence{}
	for _, o := range l.Occ {
		by[o.Sense] = append(by[o.Sense], o)
	}
	var out []senseGroup
	for _, p := range order {
		if len(by[p]) > 0 {
			out = append(out, senseGroup{Path: p, Gloss: gloss[p], Occ: by[p]})
			delete(by, p)
		}
	}
	// Senses the lexicon no longer lists, then the tokens with no sense.
	var rest []string
	for p := range by {
		if p != "" {
			rest = append(rest, p)
		}
	}
	sort.Strings(rest)
	for _, p := range rest {
		out = append(out, senseGroup{Path: p, Occ: by[p]})
	}
	if len(by[""]) > 0 {
		out = append(out, senseGroup{Occ: by[""]})
	}
	return out
}

func (c *Concordance) letterViews(works []Work) []letterView {
	byLetter := map[string]*letterView{}
	title := map[string]string{}
	for _, w := range works {
		title[w.Slug] = w.Title
	}
	for i, r := range c.Roots {
		l := r.Letter()
		lv := byLetter[l]
		if lv == nil {
			lv = &letterView{Letter: l, Slug: letterSlug[l]}
			byLetter[l] = lv
		}
		rv := rootView{R: r, Index: i}
		seen := map[string]bool{}
		var ws, gs []string
		for _, lm := range r.Lemmas {
			rv.Views = append(rv.Views, lemmaView{Lemma: lm, Groups: senseGroups(lm)})
			if lm.Gloss != "" && len(gs) < 3 {
				gs = append(gs, lm.Gloss)
			}
			for _, o := range lm.Occ {
				if !seen[o.Work] {
					seen[o.Work] = true
					ws = append(ws, title[o.Work])
				}
			}
		}
		rv.Works = strings.Join(ws, ", ")
		rv.Glosses = strings.Join(gs, "; ")
		lv.Roots = append(lv.Roots, rv)
		lv.N += r.Count
	}
	var out []letterView
	for _, l := range letters {
		if lv := byLetter[l]; lv != nil {
			out = append(out, *lv)
		}
	}
	return out
}

// jsIndex is the data of the 3D view and the word panel's links.
type jsIndex struct {
	Roots []jsRoot         `json:"roots"`
	Tok   map[string]int   `json:"tok"`
	Works map[string]jsWrk `json:"works"`
}

type jsRoot struct {
	Root  string     `json:"r,omitempty"`
	Heads []string   `json:"h"`
	Gloss string     `json:"g,omitempty"`
	Page  string     `json:"p"`
	Occ   [][]string `json:"o"` // work, section, token id, gloss
}

type jsWrk struct {
	Title string `json:"t"`
}

// Verse is one segment as the 3D view shows it.
type Verse struct {
	Section string     `json:"sec"`
	English string     `json:"en"`
	Words   [][]string `json:"w"` // token id, text (with maqaf)
}

// Verses returns the segments of a work that have a final translation.
func Verses(s *Source) map[string]Verse {
	out := map[string]Verse{}
	for _, seg := range s.Corpus.Segments {
		f := s.Final[seg.Ref]
		if f == nil {
			continue
		}
		v := Verse{Section: seg.Section, English: f.English}
		for _, t := range seg.Tokens {
			x := t.Text
			if t.Maqaf {
				x += "־"
			}
			x += t.Punct
			v.Words = append(v.Words, []string{t.ID, x})
		}
		out[seg.Ref] = v
	}
	return out
}

// Write writes docs/lexicon: the index of roots, a page per initial
// letter, index.json and verses-<work>.json.
func (c *Concordance) Write(dir string, works []Work, srcs []*Source) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	lvs := c.letterViews(works)
	page := map[int]string{}
	for _, lv := range lvs {
		for _, r := range lv.Roots {
			page[r.Index] = lv.Slug + ".html#r" + fmt.Sprint(r.Index)
		}
	}
	for _, lv := range lvs {
		var b bytes.Buffer
		if err := letterTmpl.Execute(&b, map[string]any{"L": lv, "Letters": lvs}); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, lv.Slug+".html"), b.Bytes(), 0o644); err != nil {
			return err
		}
	}
	total := 0
	for _, r := range c.Roots {
		total += r.Count
	}
	var b bytes.Buffer
	if err := indexTmpl.Execute(&b, map[string]any{"Letters": lvs, "Works": works, "Groups": len(c.Roots), "Tokens": total}); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "index.html"), b.Bytes(), 0o644); err != nil {
		return err
	}
	ix := jsIndex{Tok: c.ByToken, Works: map[string]jsWrk{}}
	for _, w := range works {
		ix.Works[w.Slug] = jsWrk{Title: w.Title}
	}
	for i, r := range c.Roots {
		jr := jsRoot{Root: r.Root, Page: page[i]}
		for _, l := range r.Lemmas {
			jr.Heads = append(jr.Heads, lex.StripHomograph(l.Headword))
			if jr.Gloss == "" {
				jr.Gloss = l.Gloss
			}
			for _, o := range l.Occ {
				jr.Occ = append(jr.Occ, []string{o.Work, o.Section, o.Token, o.Gloss})
			}
		}
		ix.Roots = append(ix.Roots, jr)
	}
	if err := writeJSON(filepath.Join(dir, "index.json"), ix); err != nil {
		return err
	}
	for _, s := range srcs {
		if err := writeJSON(filepath.Join(dir, "verses-"+s.Cfg.Slug+".json"), Verses(s)); err != nil {
			return err
		}
	}
	return nil
}

func writeJSON(p string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(p, append(b, '\n'), 0o644)
}

// RefOf gives the segment ref of a token id ("18:1.3" → "18:1").
func RefOf(id string) string {
	if i := strings.LastIndexByte(id, '.'); i > 0 {
		return id[:i]
	}
	return id
}

var funcs = template.FuncMap{
	"link": func(o Occurrence) string { return "../" + o.Work + "/" + o.Section + ".html#t" + o.Token },
}

var letterTmpl = template.Must(template.New("letter").Funcs(funcs).Parse(css + letterHTML))
var indexTmpl = template.Must(template.New("index").Funcs(funcs).Parse(css + indexHTML))
