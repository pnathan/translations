package anchor

import (
	"testing"

	"github.com/pnathan/translations/internal/corpus"
)

func toks(words ...string) []corpus.Token {
	var out []corpus.Token
	for _, w := range words {
		out = append(out, corpus.Token{Text: w})
	}
	return out
}

func TestMatchGap(t *testing.T) {
	// Rashi 3:6 quotes "ונחמד להשכיל" over "וְנֶחְמָד הָעֵץ לְהַשְׂכִּיל".
	v := toks("וְתַּאֲוָה", "ה֣וּא", "לָעֵינַ֗יִם", "וְנֶחְמָ֤ד", "הָעֵץ֙", "לְהַשְׂכִּ֔יל")
	idx, _ := match(dhWords("ונחמד להשכיל."), v, 0)
	if len(idx) != 2 || idx[0] != 3 || idx[1] != 5 {
		t.Errorf("match = %v, want [3 5]", idx)
	}
}

func TestDHName(t *testing.T) {
	q := dhWords("את ה'.")
	if len(q) != 2 || q[1] != "יהוה" {
		t.Errorf("dhWords = %q", q)
	}
}

func TestPieces(t *testing.T) {
	p := pieces("<b>בְּרֵאשִׁית.</b> אָמַר רַבִּי יִצְחָק <b>ברא</b> עוד")
	if len(p) != 2 || p[0].dh != "בְּרֵאשִׁית." || p[1].dh != "ברא" {
		t.Errorf("pieces = %+v", p)
	}
}
