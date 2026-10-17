package llm

import (
	"encoding/json"
	"testing"
)

func TestApplyTokenPath(t *testing.T) {
	tu := &TranslationUnit{Unit: "u", Model: "m", Segments: []SegmentTr{{Ref: "1:1", Tokens: []TokenGloss{{ID: "1:1.2", Gloss: "old"}}}}}
	ru := &ReviewUnit{Model: "r", Patches: []Patch{{Ref: "1:1", Path: "token:1:1.2.gloss", Value: json.RawMessage(`"new"`)}}}
	out, err := Apply(tu, ru)
	if err != nil {
		t.Fatal(err)
	}
	if g := out.Segments[0].Tokens[0].Gloss; g != "new" {
		t.Errorf("gloss = %q", g)
	}
	if tu.Segments[0].Tokens[0].Gloss != "old" {
		t.Error("Apply changed its input")
	}
}
