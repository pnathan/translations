package lex

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/pnathan/translations/internal/hebrew"
)

// Sefaria lexicon lookups (api/words). Only Jastrow and BDB entries are
// kept: both are in the public domain.

type sfEntry struct {
	Headword      string          `json:"headword"`
	ParentLexicon string          `json:"parent_lexicon"`
	StrongNumber  string          `json:"strong_number"`
	Content       json.RawMessage `json:"content"`
}

type sfSense struct {
	Definition string    `json:"definition"`
	Senses     []sfSense `json:"senses"`
	Grammar    struct {
		VerbalStem string `json:"verbal_stem"`
	} `json:"grammar"`
}

var keep = map[string]string{
	"Jastrow Dictionary":   "jastrow",
	"BDB Augmented Strong": "bdbs",
}

var reHomograph = regexp.MustCompile(`[\sIV²³¹]+$`)

func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

func flattenSenses(ss []sfSense, prefix string, stem string, out *[]Sense) {
	for i, s := range ss {
		p := fmt.Sprintf("%d", i+1)
		if prefix != "" {
			p = prefix + "." + p
		}
		st := stem
		if s.Grammar.VerbalStem != "" {
			st = s.Grammar.VerbalStem
		}
		if d := hebrew.StripHTML(s.Definition); d != "" && len(*out) < 16 {
			*out = append(*out, Sense{Path: p, Stem: st, Gloss: truncate(d, 260)})
		}
		flattenSenses(s.Senses, p, st, out)
	}
}

// ParseWords reads an api/words response and returns the Jastrow and BDB
// entries whose head word has the consonants of lemma.
func ParseWords(b []byte, lemma string) ([]*Entry, error) {
	var es []sfEntry
	if err := json.Unmarshal(b, &es); err != nil {
		var e struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(b, &e) == nil && e.Error != "" {
			return nil, nil
		}
		return nil, err
	}
	want := hebrew.Skeleton(lemma)
	var out []*Entry
	for _, e := range es {
		src, ok := keep[e.ParentLexicon]
		if !ok {
			continue
		}
		head := strings.TrimSpace(reHomograph.ReplaceAllString(e.Headword, ""))
		if hebrew.Skeleton(head) != want {
			continue
		}
		var c struct {
			Morphology string    `json:"morphology"`
			Senses     []sfSense `json:"senses"`
		}
		_ = json.Unmarshal(e.Content, &c)
		en := &Entry{Key: src + ":" + e.Headword, Source: e.ParentLexicon + " (Sefaria)", Headword: e.Headword, POS: c.Morphology, Strong: e.StrongNumber}
		flattenSenses(c.Senses, "", "", &en.Senses)
		uniquePaths(en.Senses)
		if len(en.Senses) > 0 {
			en.Gloss = en.Senses[0].Gloss
		}
		out = append(out, en)
	}
	return out, nil
}
