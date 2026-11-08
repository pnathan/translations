// Package pack renders the Markdown a model reads for one unit: the text,
// every token with its morphology and lexicon entry, the variants, and the
// commentary anchored to its spans. The model's answer is a JSON file the
// pipeline checks; the pack itself is never parsed back.
package pack

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
	"github.com/pnathan/translations/internal/work"
)

const secondaryLimit = 1500 // runes of a non-normative comment in a translation pack

func cut(s string, n int) string {
	r := []rune(s)
	if n <= 0 || len(r) <= n {
		return s
	}
	return string(r[:n]) + " … [cut; full text in the review pack]"
}

func span(c *corpus.Corpus, a anchor.Anchor) string {
	if len(a.Tokens) == 0 {
		return "whole " + a.Ref
	}
	seg := c.Segment(a.Ref)
	var words []string
	for _, id := range a.Tokens {
		for _, t := range seg.Tokens {
			if t.ID == id {
				words = append(words, t.Text)
			}
		}
	}
	return fmt.Sprintf("%s «%s»", strings.Join(a.Tokens, " "), strings.Join(words, " "))
}

func morphCell(t corpus.Token) (forms, lemmas, parse string) {
	var f, l, p []string
	for _, m := range t.Morphemes {
		f = append(f, m.Form)
		x := m.Lemma
		if m.Lex != "" {
			x += " `" + m.Lex + "`"
		}
		l = append(l, x)
		q := m.POS
		if m.Parse != "" {
			q += ": " + m.Parse
		}
		p = append(p, q)
	}
	return strings.Join(f, " + "), strings.Join(l, " + "), strings.Join(p, " · ")
}

func writeSegment(b *strings.Builder, cfg *work.Config, c *corpus.Corpus, seg *corpus.Segment, anchors []anchor.Anchor, limit int, withMorph bool) {
	fmt.Fprintf(b, "## %s %s\n\n", cfg.RefName, seg.Ref)
	fmt.Fprintf(b, "Hebrew (%s): %s\n\n", cfg.Text.Name, seg.Text)
	if seg.Para != "" {
		fmt.Fprintf(b, "Paragraph mark after this verse: %s\n\n", seg.Para)
	}
	if withMorph {
		b.WriteString("| id | word | morphemes | lemmas | parse | root | accents |\n|---|---|---|---|---|---|---|\n")
		for _, t := range seg.Tokens {
			f, l, p := morphCell(t)
			acc := strings.Join(t.Accents, ", ")
			if t.Disjunctive {
				acc += " (disjunctive)"
			}
			w := t.Text
			if t.Maqaf {
				w += "־"
			}
			if t.Ketiv != "" {
				w += " (ketiv " + t.Ketiv + ")"
			}
			fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s |\n", t.ID, w, f, l, p, t.Root, acc)
		}
		b.WriteString("\n")
	} else {
		b.WriteString("| id | word | punctuation after |\n|---|---|---|\n")
		for _, t := range seg.Tokens {
			fmt.Fprintf(b, "| %s | %s | %s |\n", t.ID, t.Text, t.Punct)
		}
		b.WriteString("\n")
	}
	if len(seg.Variants) > 0 {
		b.WriteString("Variants:\n")
		for _, v := range seg.Variants {
			fmt.Fprintf(b, "- %s %s: %s reads %s; base %s\n", strings.Join(v.Tokens, " "), v.Kind, v.Witness, v.Reading, v.Base)
		}
		b.WriteString("\n")
	}
	for _, t := range seg.Tokens {
		for _, n := range t.Notes {
			fmt.Fprintf(b, "- %s: %s\n", t.ID, n)
		}
	}
	var here []anchor.Anchor
	for _, a := range anchors {
		if a.Ref == seg.Ref {
			here = append(here, a)
		}
	}
	if len(here) > 0 {
		b.WriteString("\n### Commentary anchored to this verse\n\n")
		for _, a := range here {
			tag := ""
			if a.Normative {
				tag = " **normative**"
			}
			text := a.Text
			if !a.Normative {
				text = cut(text, limit)
			}
			dh := ""
			if a.DH != "" {
				dh = " DH «" + a.DH + "»"
			}
			fmt.Fprintf(b, "- `%s` %s%s on %s (%s; filed at %s)%s: %s\n", a.ID, a.Commentator, tag, span(c, a), a.Status, a.Filed, dh, text)
		}
	}
	b.WriteString("\n")
}

func writeLexicon(b *strings.Builder, c *corpus.Corpus, refs []string, lexicon map[string]*lex.Entry) {
	keys := map[string]bool{}
	for _, r := range refs {
		for _, t := range c.Segment(r).Tokens {
			for _, m := range t.Morphemes {
				if m.Lex != "" {
					keys[m.Lex] = true
				}
				// Mishnah: list every homograph with the lemma's consonants.
				if strings.HasPrefix(m.Lex, "jastrow:") || strings.HasPrefix(m.Lex, "bdbs:") {
					for k, e := range lexicon {
						if e.Headword != "" && sameLemma(e.Headword, m.Lemma) {
							keys[k] = true
						}
					}
				}
			}
		}
	}
	var ks []string
	for k := range keys {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	b.WriteString("## Lexicon\n\nSenses are cited by entry key and path. BDB entries list senses by stem (binyan).\n\n")
	for _, k := range ks {
		e := lexicon[k]
		if e == nil {
			continue
		}
		fmt.Fprintf(b, "### `%s` %s", k, e.Headword)
		if e.POS != "" {
			fmt.Fprintf(b, " (%s)", e.POS)
		}
		if e.Root != "" {
			fmt.Fprintf(b, " root %s", e.Root)
		}
		if e.Gloss != "" {
			fmt.Fprintf(b, ": %s", e.Gloss)
		}
		b.WriteString("\n")
		// Every sense: the derived stems come last in BDB, and a cut
		// would hide the senses a Hiphil or Hophal form needs.
		for _, s := range e.Senses {
			st := ""
			if s.Stem != "" {
				st = s.Stem + " "
			}
			fmt.Fprintf(b, "- `%s` %s%s\n", s.Path, st, s.Gloss)
		}
		b.WriteString("\n")
	}
}

func sameLemma(a, b string) bool {
	return skeleton(a) == skeleton(b)
}

func header(b *strings.Builder, cfg *work.Config, u corpus.Unit, kind, brief, out string) {
	fmt.Fprintf(b, "# %s pack: %s %s (%s %s–%s)\n\n", kind, cfg.Slug, u.ID, cfg.RefName, u.Refs[0], u.Refs[len(u.Refs)-1])
	fmt.Fprintf(b, "Brief: `%s`. Write: `%s`. Check: `go run ./cmd/tr check %s %s`.\n\n", brief, out, cfg.Slug, u.ID)
}

// Write renders every pack the current state calls for.
func Write(cfg *work.Config, c *corpus.Corpus, lexicon map[string]*lex.Entry, anchors []anchor.Anchor) error {
	for _, d := range []string{"morph", "translate", "review"} {
		os.MkdirAll(cfg.Path("stage", "packs", d), 0o755)
	}
	os.MkdirAll(filepath.Dir(llm.ReviewPackPath(cfg, "x")), 0o755)
	counts := map[string]int{}
	for _, u := range c.Units {
		hasMorph := true
		for _, r := range u.Refs {
			for _, t := range c.Segment(r).Tokens {
				if len(t.Morphemes) == 0 {
					hasMorph = false
				}
			}
		}
		rel := func(p string) string { r, _ := filepath.Rel(".", p); return r }
		if !hasMorph {
			var b strings.Builder
			header(&b, cfg, u, "Morphology", "briefs/morph.md", rel(llm.MorphPath(cfg, u.ID)))
			for _, r := range u.Refs {
				writeSegment(&b, cfg, c, c.Segment(r), nil, 0, false)
			}
			if err := writeAtomic(cfg.Path("stage", "packs", "morph", u.ID+".md"), b.String()); err != nil {
				return err
			}
			counts["morph"]++
			continue
		}
		var b strings.Builder
		// The translation pack names the edition's path, whatever the variant.
		v := llm.Variant
		llm.Variant = ""
		header(&b, cfg, u, "Translation", "briefs/translate.md", rel(llm.TranslationPath(cfg, u.ID)))
		llm.Variant = v
		for _, r := range u.Refs {
			writeSegment(&b, cfg, c, c.Segment(r), anchors, secondaryLimit, true)
		}
		writeLexicon(&b, c, u.Refs, lexicon)
		if err := writeAtomic(cfg.Path("stage", "packs", "translate", u.ID+".md"), b.String()); err != nil {
			return err
		}
		counts["translate"]++
		var tu llm.TranslationUnit
		if corpus.ReadJSON(llm.TranslationPath(cfg, u.ID), &tu) != nil {
			continue
		}
		var r strings.Builder
		header(&r, cfg, u, "Review", "briefs/review.md", rel(llm.ReviewPath(cfg, u.ID)))
		for i, ref := range u.Refs {
			writeSegment(&r, cfg, c, c.Segment(ref), anchors, 0, true)
			if i < len(tu.Segments) {
				writeTranslation(&r, &tu.Segments[i])
			}
		}
		writeLexicon(&r, c, u.Refs, lexicon)
		if err := writeAtomic(llm.ReviewPackPath(cfg, u.ID), r.String()); err != nil {
			return err
		}
		counts["review"]++
	}
	fmt.Fprintf(os.Stderr, "packs: %v\n", counts)
	return nil
}

func writeTranslation(b *strings.Builder, s *llm.SegmentTr) {
	fmt.Fprintf(b, "### Translation under review (%s)\n\n", s.Ref)
	fmt.Fprintf(b, "English: %s\n\nLiteral: %s\n\n", s.English, s.Literal)
	b.WriteString("| token | gloss | lex | sense | note |\n|---|---|---|---|---|\n")
	for _, g := range s.Tokens {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n", g.ID, g.Gloss, g.Lex, g.Sense, g.Note)
	}
	b.WriteString("\nPhrases:\n")
	for _, p := range s.Phrases {
		fmt.Fprintf(b, "- %s | %s | %s\n", strings.Join(p.Tokens, " "), p.English, p.Literal)
	}
	b.WriteString("\nGrammar notes:\n")
	for _, g := range s.Grammar {
		fmt.Fprintf(b, "- %s\n", g)
	}
	if s.Note != "" {
		fmt.Fprintf(b, "\nTranslator's note: %s\n", s.Note)
	}
	b.WriteString("\nCommentary rulings:\n")
	for _, c := range s.Commentary {
		fmt.Fprintf(b, "- `%s` %s: %s\n", c.Anchor, c.Relation, c.Summary)
	}
	for _, a := range s.Apparatus {
		fmt.Fprintf(b, "- apparatus: %s\n", a)
	}
	b.WriteString("\n")
}

func skeleton(s string) string { return hebrew.Skeleton(s) }

// writeAtomic replaces a file in one rename, so an agent that reads a pack
// while the build runs never sees a partial file.
func writeAtomic(path, content string) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(content), 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
