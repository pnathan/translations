package oshb

import "testing"

func TestDecode(t *testing.T) {
	lang, ms := Decode("HC/Vqw3ms")
	if lang != "hbo" || len(ms) != 2 {
		t.Fatalf("Decode = %s %v", lang, ms)
	}
	if ms[0].POS != "conjunction" {
		t.Errorf("morpheme 1: %v", ms[0])
	}
	if want := "qal, sequential imperfect (wayyiqtol), 3rd person, masculine, singular"; ms[1].Parse != want {
		t.Errorf("morpheme 2: %q, want %q", ms[1].Parse, want)
	}
	_, ms = Decode("HVprfsa")
	if want := "piel, participle, feminine, singular, absolute"; ms[0].Parse != want {
		t.Errorf("participle: %q, want %q", ms[0].Parse, want)
	}
	_, ms = Decode("HVqrmsa")
	if want := "qal, participle active, masculine, singular, absolute"; ms[0].Parse != want {
		t.Errorf("qal participle: %q, want %q", ms[0].Parse, want)
	}
	_, ms = Decode("HR/Sp2fs")
	if ms[1].POS != "suffix" || ms[1].Parse != "pronominal, 2nd person, feminine, singular" {
		t.Errorf("suffix: %v", ms[1])
	}
}
