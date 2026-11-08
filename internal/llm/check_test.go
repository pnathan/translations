package llm

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/pnathan/translations/internal/anchor"
	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/lex"
)

func readingEnv() *Env {
	seg := corpus.Segment{Ref: "1:1", Tokens: []corpus.Token{
		{ID: "1:1.1", Text: "א", Lex: "bdb:x"}, {ID: "1:1.2", Text: "ב"}, {ID: "1:1.3", Text: "ג"},
	}}
	return &Env{
		Corpus:  &corpus.Corpus{Segments: []corpus.Segment{seg}},
		Lexicon: map[string]*lex.Entry{"bdb:x": {Key: "bdb:x", Senses: []lex.Sense{{Path: "1"}, {Path: "2"}}}},
		Anchors: []anchor.Anchor{{ID: "rashi:1:1:1", Ref: "1:1"}},
	}
}

func readingSeg(rs ...Reading) *SegmentTr {
	return &SegmentTr{Ref: "1:1", English: "a b c", Literal: "a b c", Grammar: []string{"g"},
		Tokens:   []TokenGloss{{ID: "1:1.1", Gloss: "a"}, {ID: "1:1.2", Gloss: "b"}, {ID: "1:1.3", Gloss: "c"}},
		Phrases:  []Phrase{{Tokens: []string{"1:1.1", "1:1.2", "1:1.3"}, English: "a b c", Literal: "x"}},
		Readings: rs}
}

func TestReadings(t *testing.T) {
	e := readingEnv()
	if errs := e.CheckSegment(readingSeg(Reading{Tokens: []string{"1:1.1", "1:1.2"}, English: "other", Source: "rashi:1:1:1",
		Senses: []ReadingSense{{ID: "1:1.1", Sense: "2"}}})); len(errs) != 0 {
		t.Errorf("good reading: %v", errs)
	}
	for name, c := range map[string]struct {
		r    Reading
		want string
	}{
		"gap":     {Reading{Tokens: []string{"1:1.1", "1:1.3"}, English: "x", Source: "plain"}, "not contiguous"},
		"source":  {Reading{Tokens: []string{"1:1.1"}, English: "x", Source: "rashi:9:9:9"}, "not an anchor"},
		"nosrc":   {Reading{Tokens: []string{"1:1.1"}, English: "x"}, "no source"},
		"hebrew":  {Reading{Tokens: []string{"1:1.1"}, English: "אב", Source: "plain"}, "Hebrew"},
		"sense":   {Reading{Tokens: []string{"1:1.1"}, English: "x", Source: "plain", Senses: []ReadingSense{{ID: "1:1.1", Sense: "9"}}}, "not a sense"},
		"outside": {Reading{Tokens: []string{"1:1.2"}, English: "x", Source: "plain", Senses: []ReadingSense{{ID: "1:1.1", Sense: "1"}}}, "outside the span"},
	} {
		errs := strings.Join(e.CheckSegment(readingSeg(c.r)), "; ")
		if !strings.Contains(errs, c.want) {
			t.Errorf("%s: want %q in %q", name, c.want, errs)
		}
	}
}

func TestApplyReadings(t *testing.T) {
	tu := &TranslationUnit{Segments: []SegmentTr{{Ref: "1:1"}}}
	ru := &ReviewUnit{Patches: []Patch{{Ref: "1:1", Path: "readings", Value: json.RawMessage(`[{"tokens":["1:1.1"],"english":"x","source":"plain"}]`)}}}
	out, err := Apply(tu, ru)
	if err != nil || len(out.Segments[0].Readings) != 1 {
		t.Fatalf("apply readings: %v %+v", err, out)
	}
}
