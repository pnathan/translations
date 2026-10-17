package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/anchor"
	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/hebrew"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/oshb"
	"github.com/pnathan/translations/internal/pack"
	"github.com/pnathan/translations/internal/sefaria"
	"github.com/pnathan/translations/internal/work"
)

// chapters reads the cached per-chapter responses of one text.
func chapters(cfg *work.Config, kind, version string) ([][][]string, error) {
	var out [][][]string
	for ch := 1; ch <= cfg.Chapters; ch++ {
		b, err := os.ReadFile(cfg.Path("raw", textName(kind, ch)))
		if os.IsNotExist(err) && strings.HasPrefix(kind, "commentary/") {
			out = append(out, nil) // the commentator wrote nothing on this chapter
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("%w (run tr fetch)", err)
		}
		v, err := sefaria.ParseText(b, version)
		if err != nil {
			return nil, err
		}
		segs, err := sefaria.Segments(v.Text)
		if err != nil {
			return nil, fmt.Errorf("%s chapter %d: %w", kind, ch, err)
		}
		out = append(out, segs)
	}
	return out, nil
}

func firsts(g [][][]string) [][]string {
	out := make([][]string, len(g))
	for i, ch := range g {
		for _, s := range ch {
			x := ""
			if len(s) > 0 {
				x = strings.Join(s, " ")
			}
			out[i] = append(out[i], x)
		}
	}
	return out
}

func build(cfg *work.Config) error {
	if err := os.MkdirAll(cfg.Path("stage"), 0o755); err != nil {
		return err
	}
	base, err := chapters(cfg, "text", cfg.Text.Version)
	if err != nil {
		return err
	}
	var (
		c       *corpus.Corpus
		lexicon map[string]*lex.Entry
	)
	switch cfg.Kind {
	case work.Tanakh:
		f, err := os.Open(cfg.Path("raw", "oshb", cfg.OSIS+".xml"))
		if err != nil {
			return err
		}
		verses, err := oshb.Read(f)
		f.Close()
		if err != nil {
			return err
		}
		open := func(n string) *os.File {
			f, err := os.Open(filepath.Join("sources", "openscriptures", n))
			if err != nil {
				fail(err)
			}
			return f
		}
		li, au, bd := open("LexicalIndex.xml"), open("AugIndex.xml"), open("BrownDriverBriggs.xml")
		ix, err := lex.Load(li, au, bd)
		li.Close()
		au.Close()
		bd.Close()
		if err != nil {
			return err
		}
		c, lexicon, err = corpus.BuildTanakh(cfg, firsts(base), verses, ix)
		if err != nil {
			return err
		}
	case work.Mishnah:
		wit := map[string][][]string{}
		for _, w := range cfg.Witnesses {
			g, err := chapters(cfg, "witness/"+slug(w.Name), w.Version)
			if err != nil {
				return err
			}
			wit[w.Name] = firsts(g)
		}
		morph, src, err := llm.LoadMorph(cfg)
		if err != nil {
			return err
		}
		c = corpus.BuildMishnah(cfg, firsts(base), wit, morph, src)
		lexicon = mishnahLexicon(cfg, c)
	}
	var anchors []anchor.Anchor
	for _, cm := range cfg.Commentators {
		g, err := chapters(cfg, "commentary/"+cm.Key, cm.Version)
		if err != nil {
			return err
		}
		var chs []anchor.Comments
		for i, ch := range g {
			chs = append(chs, anchor.Comments{Chapter: i + 1, Verses: ch})
		}
		anchors = append(anchors, anchor.Build(c, cm.Key, cm.Name, cm.Normative, chs)...)
	}
	if err := corpus.WriteJSON(cfg.Path("stage", "corpus.json"), c); err != nil {
		return err
	}
	if err := corpus.WriteJSON(cfg.Path("stage", "lexicon.json"), lexicon); err != nil {
		return err
	}
	if err := corpus.WriteJSON(cfg.Path("stage", "anchors.json"), anchors); err != nil {
		return err
	}
	report(c, anchors)
	return pack.Write(cfg, c, lexicon, anchors)
}

// mishnahLexicon resolves the model's lemmas against the cached Sefaria
// lookups and sets each morpheme's lexicon key to the first candidate.
func mishnahLexicon(cfg *work.Config, c *corpus.Corpus) map[string]*lex.Entry {
	out := map[string]*lex.Entry{}
	for si := range c.Segments {
		for ti := range c.Segments[si].Tokens {
			t := &c.Segments[si].Tokens[ti]
			for mi := range t.Morphemes {
				m := &t.Morphemes[mi]
				if m.Lemma == "" || isFunctionPOS(m.POS) || hebrew.Skeleton(m.Lemma) == "" {
					continue
				}
				p := cfg.Path("raw", "lexicon", hebrew.Skeleton(m.Lemma), hebrew.Unpointed(m.Lemma)+".json")
				b, err := os.ReadFile(p)
				if err != nil {
					continue
				}
				es, _ := lex.ParseWords(b, m.Lemma)
				for _, e := range es {
					out[e.Key] = e
				}
				if len(es) > 0 && m.Lex == "" {
					m.Lex = es[0].Key
				}
				if t.Lex == "" && m.Lex != "" {
					t.Lex = m.Lex
				}
			}
		}
	}
	return out
}

func report(c *corpus.Corpus, anchors []anchor.Anchor) {
	toks, vars := 0, 0
	for _, s := range c.Segments {
		toks += len(s.Tokens)
		vars += len(s.Variants)
	}
	st := map[string]map[string]int{}
	for _, a := range anchors {
		if st[a.Commentator] == nil {
			st[a.Commentator] = map[string]int{}
		}
		st[a.Commentator][a.Status]++
	}
	fmt.Fprintf(os.Stderr, "%d segments, %d units, %d tokens, %d variants\n", len(c.Segments), len(c.Units), toks, vars)
	var names []string
	for n := range st {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		fmt.Fprintf(os.Stderr, "  %-15s %v\n", n, st[n])
	}
}
