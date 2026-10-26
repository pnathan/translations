package lex

import (
	"strings"
	"testing"
)

func TestReadBDBSensesWithoutDef(t *testing.T) {
	const x = `<lexicon><entry id="v.ed.aa"><w>שָׁמַע</w> <pos>vb</pos>. <def>hear</def>
<sense><stem>Qal</stem> 1052 <sense n="1">subj. pers. hum.</sense> <sense n="2">1st of men</sense></sense>
<sense><stem>Niph</stem>. <sense n="1"><def>be heard</def></sense></sense>
<sense><stem>Hiph</stem>. 63 <sense n="1">of man</sense></sense>
</entry></lexicon>`
	ix := &Index{bdb: map[string]*bdbEntry{}}
	if err := ix.readBDB(strings.NewReader(x)); err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for _, s := range ix.bdb["v.ed.aa"].Senses {
		got[s.Path] = s.Gloss
	}
	want := map[string]string{
		"Qal 1":  "hear — subj. pers. hum",
		"Qal 2":  "hear — 1st of men",
		"Niph 1": "be heard",
		"Hiph 1": "of man (Hiph of: hear)",
	}
	if len(got) != len(want) {
		t.Errorf("senses %v", got)
	}
	for p, g := range want {
		if got[p] != g {
			t.Errorf("%s = %q, want %q", p, got[p], g)
		}
	}
}
