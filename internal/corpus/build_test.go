package corpus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pnathan/translations/internal/work"
)

func proseCfg() *work.Config {
	return &work.Config{Slug: "p", Kind: work.Prose, UnitMax: 99, UnitWords: 10,
		Sections: []work.Section{{Slug: "c1", From: "1:1", To: "1:99"}, {Slug: "c2", From: "2:1", To: "2:99"}}}
}

func TestBuildProseUnitsByWords(t *testing.T) {
	w := func(n int) string { return strings.TrimSpace(strings.Repeat("דבר ", n)) }
	base := [][]string{{w(6), w(6), w(30), w(2)}, {w(3)}}
	c := BuildProse(proseCfg(), base, nil)
	var got []string
	for _, u := range c.Units {
		got = append(got, fmt.Sprintf("%s:%s", u.ID, strings.Join(u.Refs, ",")))
	}
	// 6+6 reaches 12 >= 10 and closes; the 30-word paragraph stands alone;
	// the last paragraph of chapter 1 starts a new unit; chapter 2 starts another.
	want := "c1-01:1:1,1:2 c1-02:1:3 c1-03:1:4 c2-01:2:1"
	if g := strings.Join(got, " "); g != want {
		t.Fatalf("units %q, want %q", g, want)
	}
	if tk := c.Segments[0].Tokens[0]; tk.ID != "1:1.1" || tk.Lang != "heb" || len(tk.Morphemes) != 0 {
		t.Fatalf("token %+v", tk)
	}
}

func TestBuildProseCollates(t *testing.T) {
	base := [][]string{{"אחד שני שלישי"}}
	wit := map[string][][]string{"PDF": {{"אחד שלישי"}}}
	c := BuildProse(proseCfg(), base, wit)
	v := c.Segments[0].Variants
	if len(v) != 1 || v[0].Base != "שני" || v[0].Reading != "—" || v[0].Kind != "word" {
		t.Fatalf("variants %+v", v)
	}
}
