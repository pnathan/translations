package render

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
	"github.com/pnathan/translations/internal/work"
)

type pageWord struct {
	ID, Text, Gloss string
	Maqaf           bool
	Punct           string
	Anchored        bool
}

type pageComment struct {
	ID, Commentator, Relation, Summary, Hebrew, Span string
	Normative                                        bool
}

type pageSeg struct {
	Ref, English, Literal, Note, Para string
	Words                             []pageWord
	Grammar, Apparatus                []string
	Comments                          []pageComment
	Missing                           bool
}

type pageVocab struct {
	Lemma, Root, Gloss, Key string
	Count                   int
}

type page struct {
	Cfg        *work.Config
	Section    work.Section
	Sections   []work.Section
	Segs       []pageSeg
	Vocab      []pageVocab
	Data       template.JS
	Provenance template.HTML
	First      string
	Last       string
}

// token data for the word panel
type jsMorph struct {
	Form  string `json:"f"`
	Lemma string `json:"l,omitempty"`
	POS   string `json:"p,omitempty"`
	Parse string `json:"x,omitempty"`
	Lex   string `json:"k,omitempty"`
}

type jsTok struct {
	Text    string    `json:"t"`
	Morph   []jsMorph `json:"m,omitempty"`
	Root    string    `json:"r,omitempty"`
	Accents []string  `json:"a,omitempty"`
	Gloss   string    `json:"g,omitempty"`
	Sense   string    `json:"s,omitempty"`
	Lex     string    `json:"k,omitempty"`
	Note    string    `json:"n,omitempty"`
	Ketiv   string    `json:"kt,omitempty"`
	Alt     string    `json:"alt,omitempty"`
	Source  string    `json:"src,omitempty"`
	Notes   []string  `json:"wn,omitempty"`
	Comms   []string  `json:"c,omitempty"`
}

type jsData struct {
	Tokens map[string]jsTok      `json:"tokens"`
	Lex    map[string]*lex.Entry `json:"lex"`
}

func (e *Edition) section(sec work.Section) (*page, error) {
	p := &page{Cfg: e.Cfg, Section: sec, Sections: e.Cfg.Sections}
	data := jsData{Tokens: map[string]jsTok{}, Lex: map[string]*lex.Entry{}}
	vocab := map[string]*pageVocab{}
	for i := range e.Corpus.Segments {
		s := &e.Corpus.Segments[i]
		if s.Section != sec.Slug {
			continue
		}
		if p.First == "" {
			p.First = s.Ref
		}
		p.Last = s.Ref
		f := e.Final[s.Ref]
		ps := pageSeg{Ref: s.Ref, Para: s.Para, Missing: f == nil}
		if f != nil {
			ps.English, ps.Literal, ps.Note, ps.Grammar = f.English, f.Literal, f.Note, f.Grammar
			ps.Apparatus = append(ps.Apparatus, f.Apparatus...)
		}
		for _, v := range s.Variants {
			ps.Apparatus = append(ps.Apparatus, fmt.Sprintf("%s — %s reads %s (%s)", v.Base, v.Witness, v.Reading, v.Kind))
		}
		comms := map[string][]string{}
		for _, a := range e.anchorsOf(s.Ref) {
			c := pageComment{ID: a.ID, Commentator: a.Commentator, Hebrew: a.Text, Normative: a.Normative}
			if r := e.ruling(f, a.ID); r != nil {
				c.Relation, c.Summary = r.Relation, r.Summary
			}
			var words []string
			for _, id := range a.Tokens {
				comms[id] = append(comms[id], a.ID)
				for _, t := range s.Tokens {
					if t.ID == id {
						words = append(words, t.Text)
					}
				}
			}
			c.Span = strings.Join(words, " ")
			if c.Span == "" {
				c.Span = "the whole " + e.Cfg.Segment
			}
			ps.Comments = append(ps.Comments, c)
		}
		for _, t := range s.Tokens {
			g := e.gloss(f, t.ID)
			w := pageWord{ID: t.ID, Text: t.Text, Maqaf: t.Maqaf, Punct: t.Punct, Anchored: len(comms[t.ID]) > 0}
			jt := jsTok{Text: t.Text, Root: t.Root, Accents: t.Accents, Lex: t.Lex, Ketiv: t.Ketiv, Alt: t.Alt, Source: t.MorphSource, Notes: t.Notes, Comms: comms[t.ID]}
			if g != nil {
				w.Gloss = g.Gloss
				jt.Gloss, jt.Sense, jt.Note = g.Gloss, g.Sense, g.Note
				if g.Lex != "" {
					jt.Lex = g.Lex
				}
			}
			for _, m := range t.Morphemes {
				jt.Morph = append(jt.Morph, jsMorph{Form: m.Form, Lemma: m.Lemma, POS: m.POS, Parse: m.Parse, Lex: m.Lex})
				if m.Lex != "" {
					if le := e.Lexicon[m.Lex]; le != nil {
						data.Lex[m.Lex] = le
					}
				}
			}
			if jt.Lex != "" {
				if le := e.Lexicon[jt.Lex]; le != nil {
					data.Lex[jt.Lex] = le
					v := vocab[jt.Lex]
					if v == nil {
						v = &pageVocab{Lemma: le.Headword, Root: le.Root, Gloss: le.Gloss, Key: jt.Lex}
						vocab[jt.Lex] = v
					}
					v.Count++
				}
			}
			data.Tokens[t.ID] = jt
			ps.Words = append(ps.Words, w)
		}
		p.Segs = append(p.Segs, ps)
	}
	for _, v := range vocab {
		p.Vocab = append(p.Vocab, *v)
	}
	sort.Slice(p.Vocab, func(i, j int) bool {
		if p.Vocab[i].Count != p.Vocab[j].Count {
			return p.Vocab[i].Count > p.Vocab[j].Count
		}
		return p.Vocab[i].Lemma < p.Vocab[j].Lemma
	})
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// json.Marshal escapes <, > and & as < etc., so the blob is safe in a script.
	p.Data = template.JS(b)
	return p, nil
}

// HTML writes docs/<slug>/index.html and one page per section.
func (e *Edition) HTML(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	prov := template.HTML("<ul>" + mdList(e.provenance()) + "</ul>")
	for _, sec := range e.Cfg.Sections {
		p, err := e.section(sec)
		if err != nil {
			return err
		}
		if len(p.Segs) == 0 {
			continue
		}
		p.Provenance = prov
		var b bytes.Buffer
		if err := sectionTmpl.Execute(&b, p); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, sec.Slug+".html"), b.Bytes(), 0o644); err != nil {
			return err
		}
	}
	var b bytes.Buffer
	if err := indexTmpl.Execute(&b, map[string]any{"Cfg": e.Cfg, "Provenance": prov, "Missing": e.Missing}); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "index.html"), b.Bytes(), 0o644)
}

func mdList(s string) string {
	var out []string
	for _, l := range strings.Split(strings.TrimSpace(s), "\n") {
		l = strings.TrimPrefix(l, "- ")
		l = template.HTMLEscapeString(l)
		for strings.Count(l, "`") >= 2 {
			l = strings.Replace(l, "`", "<code>", 1)
			l = strings.Replace(l, "`", "</code>", 1)
		}
		out = append(out, "<li>"+l+"</li>")
	}
	return strings.Join(out, "")
}

var funcs = template.FuncMap{
	"inc":   func(i int) int { return i + 1 },
	"lower": strings.ToLower,
}

var sectionTmpl = template.Must(template.New("section").Funcs(funcs).Parse(baseCSS + sectionHTML))
var indexTmpl = template.Must(template.New("index").Funcs(funcs).Parse(baseCSS + indexHTML))
