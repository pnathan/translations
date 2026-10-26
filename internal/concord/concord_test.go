package concord

import (
	"testing"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/work"
)

func TestRootKey(t *testing.T) {
	for in, want := range map[string]string{
		"ארץ":  "ארצ",
		"שׁבע": "שׁבע",
		"שׂבע": "שׂבע",
		"שבע":  "שבע",
	} {
		if got := rootKey(in); got != want {
			t.Errorf("rootKey(%s) = %s, want %s", in, got, want)
		}
	}
	if !hasUndotted("שבע") || hasUndotted("שׁבע") || !hasUndotted("שׁרש") {
		t.Error("hasUndotted")
	}
}

func TestDotFromHeadword(t *testing.T) {
	for _, c := range []struct{ root, head, want string }{
		{"עשר", "עֲשָׂרָה", "עשׂר"},
		{"עשר", "עָשִׁיר", "עשׁר"},
		{"שרש", "שֹׁרֶשׁ", "שׁרשׁ"},
		{"שמש", "שִׁמֵּשׂ", ""}, // mixed dots
		{"עשה", "פֹּעַל", ""},  // no ש in the headword
	} {
		if got := dotFromHeadword(c.root, c.head); got != c.want {
			t.Errorf("dotFromHeadword(%s, %s) = %q, want %q", c.root, c.head, got, c.want)
		}
	}
}

func src(slug string, toks []corpus.Token, lexicon map[string]*lex.Entry) *Source {
	seg := corpus.Segment{Ref: "1:1", Section: "s", Tokens: toks}
	f := &llm.SegmentTr{Ref: "1:1"}
	for _, t := range toks {
		f.Tokens = append(f.Tokens, llm.TokenGloss{ID: t.ID, Gloss: "g" + t.ID, Sense: "1"})
	}
	return &Source{Cfg: &work.Config{Slug: slug}, Corpus: &corpus.Corpus{Segments: []corpus.Segment{seg}},
		Lexicon: lexicon, Final: map[string]*llm.SegmentTr{"1:1": f}}
}

func TestBuildJoinsWorks(t *testing.T) {
	bible := src("gen", []corpus.Token{
		{ID: "1:1.1", Text: "שֶׁבַע", Lex: "bdb:seven", Root: "שׁבע"},
		{ID: "1:1.2", Text: "וַיִּשְׂבַּע", Lex: "bdb:sated", Root: "שׂבע"},
		{ID: "1:1.3", Text: "כִּי", Lex: "bdb:ki"},
	}, map[string]*lex.Entry{
		"bdb:seven": {Key: "bdb:seven", Headword: "שֶׁ֫בַע"},
		"bdb:sated": {Key: "bdb:sated", Headword: "שָׂבַע"},
		"bdb:ki":    {Key: "bdb:ki", Headword: "כִּי"},
	})
	mishnah := src("avot", []corpus.Token{
		{ID: "1:1.1", Text: "שִׁבְעָה", Lex: "jastrow:seven", Root: "שבע"},
		{ID: "1:1.2", Text: "שָׂבֵעַ", Lex: "jastrow:sated", Root: "שבע"},
		{ID: "1:1.3", Text: "שבועה", Lex: "jastrow:oath", Root: "שבע"}, // unpointed: undetermined
	}, map[string]*lex.Entry{
		"jastrow:seven": {Key: "jastrow:seven", Headword: "שִׁבְעָה"},
		"jastrow:sated": {Key: "jastrow:sated", Headword: "שָׂבֵעַ"},
		"jastrow:oath":  {Key: "jastrow:oath", Headword: "שבועה"},
	})
	c := Build([]*Source{bible, mishnah})
	group := func(tok string) *Root { return c.Roots[c.ByToken[tok]] }
	if g := group("avot:1:1.1"); g != group("gen:1:1.1") || g.Root != "שׁבע" || g.Count != 2 {
		t.Errorf("seven: %+v", g)
	}
	if g := group("avot:1:1.2"); g != group("gen:1:1.2") || g.Count != 2 {
		t.Errorf("sated: %+v", g)
	}
	if g := group("avot:1:1.3"); g.Root != "שבע" || g.Note == "" || g.Count != 1 {
		t.Errorf("oath should stay apart, undetermined: %+v", g)
	}
	if g := group("gen:1:1.3"); g.Root != "" || g.Key != "lemma:bdb:ki" {
		t.Errorf("rootless lemma: %+v", g)
	}
	if o := group("gen:1:1.1").Lemmas[0].Occ[0]; o.Gloss == "" || o.Sense != "1" {
		t.Errorf("occurrence lacks the reviewed gloss: %+v", o)
	}
}
