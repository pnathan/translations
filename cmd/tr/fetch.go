package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/hebrew"
	"github.com/pnathan/translations/internal/sefaria"
	"github.com/pnathan/translations/internal/work"
)

// Shared sources: the OpenScriptures files, used by every Tanakh work.
var openScriptures = map[string]string{
	"LexicalIndex.xml":      "https://raw.githubusercontent.com/openscriptures/HebrewLexicon/master/LexicalIndex.xml",
	"AugIndex.xml":          "https://raw.githubusercontent.com/openscriptures/HebrewLexicon/master/AugIndex.xml",
	"BrownDriverBriggs.xml": "https://raw.githubusercontent.com/openscriptures/HebrewLexicon/master/BrownDriverBriggs.xml",
}

func slug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
		} else if b.Len() > 0 && !strings.HasSuffix(b.String(), "-") {
			b.WriteRune('-')
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}

func textName(kind string, ch int) string { return fmt.Sprintf("%s/%02d.json", kind, ch) }

func fetch(cfg *work.Config, args []string) error {
	st, err := sefaria.Open(cfg.Path("raw"))
	if err != nil {
		return err
	}
	st.Refresh = len(args) > 0 && args[0] == "--refresh"
	defer st.Save()
	get := func(name, title, version string, ch int) error {
		ref := fmt.Sprintf("%s.%d", title, ch)
		b, err := st.Get(name, sefaria.TextURL(ref, version))
		if err != nil {
			return err
		}
		if _, err := sefaria.ParseText(b, version); err != nil {
			os.Remove(filepath.Join(st.Dir, name))
			return fmt.Errorf("%s: %w", ref, err)
		}
		return nil
	}
	for ch := 1; ch <= cfg.Chapters; ch++ {
		if err := get(textName("text", ch), cfg.Text.Sefaria, cfg.Text.Version, ch); err != nil {
			return err
		}
		for _, w := range cfg.Witnesses {
			if err := get(textName("witness/"+slug(w.Name), ch), w.Sefaria, w.Version, ch); err != nil {
				return err
			}
		}
		for _, c := range cfg.Commentators {
			if err := get(textName("commentary/"+c.Key, ch), c.Sefaria, c.Version, ch); err != nil {
				if !strings.Contains(err.Error(), "404") {
					return err
				}
				fmt.Fprintf(os.Stderr, "%s has no chapter %d\n", c.Name, ch)
			}
		}
		fmt.Fprintf(os.Stderr, "chapter %d fetched\n", ch)
	}
	if cfg.Kind == work.Tanakh {
		if _, err := st.Get("oshb/"+cfg.OSIS+".xml", cfg.OSHB); err != nil {
			return err
		}
		shared, err := sefaria.Open("sources/openscriptures")
		if err != nil {
			return err
		}
		for name, u := range openScriptures {
			if _, err := shared.Get(name, u); err != nil {
				return err
			}
		}
		return shared.Save()
	}
	// Mishnah: lexicon lookups for the lemmas of the morphology stage.
	c, err := corpus.Load(cfg.Path("stage", "corpus.json"))
	if err != nil {
		return nil // no corpus yet: the lookups follow the morphology stage
	}
	n := 0
	for _, lemma := range lemmas(c) {
		name := "lexicon/" + hebrew.Skeleton(lemma) + "/" + hebrew.Unpointed(lemma) + ".json"
		if _, err := st.Get(name, sefaria.WordsURL(lemma)); err != nil {
			return err
		}
		n++
	}
	fmt.Fprintf(os.Stderr, "%d lemmas looked up\n", n)
	return nil
}

// lemmas lists the distinct pointed lemmas of a mishnah corpus.
func lemmas(c *corpus.Corpus) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range c.Segments {
		for _, t := range s.Tokens {
			for _, m := range t.Morphemes {
				if m.Lemma == "" || seen[m.Lemma] || hebrew.Skeleton(m.Lemma) == "" || isFunctionPOS(m.POS) {
					continue
				}
				seen[m.Lemma] = true
				out = append(out, m.Lemma)
			}
		}
	}
	return out
}

// isFunctionPOS marks the morphemes that get a fixed gloss, not a lexicon entry.
func isFunctionPOS(p string) bool {
	switch p {
	case "conjunction", "suffix", "article":
		return true
	}
	return false
}
