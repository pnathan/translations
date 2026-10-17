// Package render writes the edition: a Markdown file per work for reading
// and review, and HTML pages per section for reading and learning.
package render

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/anchor"
	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/work"
)

// Edition joins the corpus with the final translation.
type Edition struct {
	Cfg     *work.Config
	Corpus  *corpus.Corpus
	Lexicon map[string]*lex.Entry
	Anchors []anchor.Anchor
	Final   map[string]*llm.SegmentTr // by ref
	Models  map[string]string         // unit → model provenance
	Reviews map[string]*llm.ReviewUnit
	Missing []string // units without a final translation
}

func (e *Edition) anchorsOf(ref string) []anchor.Anchor {
	var out []anchor.Anchor
	for _, a := range e.Anchors {
		if a.Ref == ref {
			out = append(out, a)
		}
	}
	return out
}

func (e *Edition) ruling(seg *llm.SegmentTr, id string) *llm.CommentUse {
	if seg == nil {
		return nil
	}
	for i := range seg.Commentary {
		if seg.Commentary[i].Anchor == id {
			return &seg.Commentary[i]
		}
	}
	return nil
}

func (e *Edition) gloss(seg *llm.SegmentTr, id string) *llm.TokenGloss {
	if seg == nil {
		return nil
	}
	for i := range seg.Tokens {
		if seg.Tokens[i].ID == id {
			return &seg.Tokens[i]
		}
	}
	return nil
}

func parseOf(t corpus.Token) string {
	var p []string
	for _, m := range t.Morphemes {
		x := m.POS
		if m.Parse != "" && m.Parse != m.POS {
			x += " (" + m.Parse + ")"
		}
		p = append(p, x)
	}
	return strings.Join(p, " + ")
}

func lemmaOf(t corpus.Token) string {
	var l []string
	for _, m := range t.Morphemes {
		if m.Lemma != "" {
			l = append(l, m.Lemma)
		}
	}
	return strings.Join(l, " + ")
}

func esc(s string) string { return strings.ReplaceAll(s, "|", "\\|") }

// Markdown writes works/<slug>/edition.md.
func (e *Edition) Markdown(path string) error {
	var b strings.Builder
	cfg := e.Cfg
	fmt.Fprintf(&b, "# %s — %s\n\n", cfg.Title, cfg.He)
	fmt.Fprintf(&b, "Translated from the Hebrew (%s) by a model pipeline; every unit reviewed. Provenance at the end.\n\n", cfg.Text.Name)
	if len(e.Missing) > 0 {
		fmt.Fprintf(&b, "Units without a reviewed translation: %s.\n\n", strings.Join(e.Missing, ", "))
	}
	for _, sec := range cfg.Sections {
		var segs []*corpus.Segment
		for i := range e.Corpus.Segments {
			if e.Corpus.Segments[i].Section == sec.Slug {
				segs = append(segs, &e.Corpus.Segments[i])
			}
		}
		if len(segs) == 0 {
			continue
		}
		fmt.Fprintf(&b, "## %s — %s (%s %s–%s)\n\n### Reading text\n\n", sec.Name, sec.He, cfg.RefName, segs[0].Ref, segs[len(segs)-1].Ref)
		var para []string
		for _, s := range segs {
			if f := e.Final[s.Ref]; f != nil {
				para = append(para, fmt.Sprintf("<sup>%s</sup> %s", s.Ref, f.English))
			}
			if s.Para != "" || cfg.Kind == work.Mishnah {
				if len(para) > 0 {
					b.WriteString(strings.Join(para, " ") + "\n\n")
				}
				para = nil
			}
		}
		if len(para) > 0 {
			b.WriteString(strings.Join(para, " ") + "\n\n")
		}
		fmt.Fprintf(&b, "### %s by %s\n\n", sec.Name, cfg.Segment)
		for _, s := range segs {
			f := e.Final[s.Ref]
			fmt.Fprintf(&b, "#### %s %s\n\n> %s\n\n", cfg.RefName, s.Ref, s.Text)
			if f == nil {
				b.WriteString("_Not yet translated._\n\n")
				continue
			}
			fmt.Fprintf(&b, "**%s**\n\nLiteral: %s\n\n", f.English, f.Literal)
			b.WriteString("| word | gloss | lemma | root | parse | sense |\n|---|---|---|---|---|---|\n")
			for _, t := range s.Tokens {
				g := e.gloss(f, t.ID)
				gl, sn := "", ""
				if g != nil {
					gl, sn = g.Gloss, g.Sense
					if g.Note != "" {
						gl += " — " + g.Note
					}
				}
				fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s |\n", t.Text, esc(gl), esc(lemmaOf(t)), t.Root, esc(parseOf(t)), sn)
			}
			b.WriteString("\n")
			if len(f.Grammar) > 0 {
				b.WriteString("Grammar:\n\n")
				for _, g := range f.Grammar {
					fmt.Fprintf(&b, "- %s\n", g)
				}
				b.WriteString("\n")
			}
			if f.Note != "" {
				fmt.Fprintf(&b, "Note: %s\n\n", f.Note)
			}
			var cs []string
			for _, a := range e.anchorsOf(s.Ref) {
				if r := e.ruling(f, a.ID); r != nil && r.Summary != "" {
					cs = append(cs, fmt.Sprintf("- %s (%s): %s", a.Commentator, r.Relation, r.Summary))
				}
			}
			if len(cs) > 0 {
				b.WriteString("Commentary:\n\n" + strings.Join(cs, "\n") + "\n\n")
			}
			var ap []string
			ap = append(ap, f.Apparatus...)
			for _, v := range s.Variants {
				ap = append(ap, fmt.Sprintf("%s: %s reads %s (%s)", v.Base, v.Witness, v.Reading, v.Kind))
			}
			if len(ap) > 0 {
				b.WriteString("Apparatus:\n\n")
				for _, a := range ap {
					fmt.Fprintf(&b, "- %s\n", a)
				}
				b.WriteString("\n")
			}
		}
	}
	b.WriteString("## Provenance\n\n")
	b.WriteString(e.provenance())
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func (e *Edition) provenance() string {
	var b strings.Builder
	cfg := e.Cfg
	fmt.Fprintf(&b, "- Base text: %s, Sefaria version “%s”.\n", cfg.Text.Name, cfg.Text.Version)
	for _, w := range cfg.Witnesses {
		fmt.Fprintf(&b, "- Collated witness: %s, Sefaria version “%s”.\n", w.Name, w.Version)
	}
	if cfg.Kind == work.Tanakh {
		b.WriteString("- Morphology and lemmas: Open Scriptures Hebrew Bible (WLC 4.x with OSHB morphology, CC BY 4.0); the WLC is collated against the base text.\n")
		b.WriteString("- Lexicon: Brown–Driver–Briggs via the OpenScriptures Hebrew Lexicon.\n")
	} else {
		b.WriteString("- Morphology: a model analysis, reviewed (see each unit's provenance); lexicon entries from Jastrow and BDB (Sefaria).\n")
	}
	for _, c := range cfg.Commentators {
		n := ""
		if c.Normative {
			n = " (normative)"
		}
		fmt.Fprintf(&b, "- Commentary: %s%s, Sefaria version “%s”.\n", c.Name, n, c.Version)
	}
	var units []string
	for u := range e.Models {
		units = append(units, u)
	}
	sort.Strings(units)
	models := map[string][]string{}
	for _, u := range units {
		models[e.Models[u]] = append(models[e.Models[u]], u)
	}
	for m, us := range models {
		fmt.Fprintf(&b, "- Translation: %s — %d units.\n", m, len(us))
	}
	b.WriteString("- Fetched sources, with dates and SHA-256: `works/" + cfg.Slug + "/raw/manifest.json`.\n")
	return b.String()
}
