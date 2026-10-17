package hebrew

import "testing"

func TestSkeleton(t *testing.T) {
	if got := Skeleton("בְּרֵאשִׁ֖ית"); got != "בראשית" {
		t.Errorf("Skeleton = %q", got)
	}
	if got := Skeleton("הָאָֽרֶץ׃"); got != "הארצ" {
		t.Errorf("final letter: Skeleton = %q", got)
	}
}

func TestSimilar(t *testing.T) {
	cases := []struct {
		q, t string
		min  float64
	}{
		{"בראשית", "בְּרֵאשִׁ֖ית", 1},
		{"וייצר", "וַיִּ֩יצֶר", 0.9}, // plene against defective
		{"תולדות", "תֹּלְד֥וֹת", 0.9},
		{"עז", "וְעֵ֥ז", 0.8}, // proclitic ו
	}
	for _, c := range cases {
		if got := Similar(c.q, c.t); got < c.min {
			t.Errorf("Similar(%q, %q) = %v, want ≥ %v", c.q, c.t, got, c.min)
		}
	}
	if got := Similar("ארץ", "שָׁמַיִם"); got != 0 {
		t.Errorf("unrelated words: %v", got)
	}
}

func TestAccents(t *testing.T) {
	a := Accents("הָאָֽרֶץ", true)
	if len(a) != 1 || a[0].Name != "silluq" {
		t.Errorf("Accents = %v, want silluq", a)
	}
	a = Accents("אֱלֹהִ֑ים", false)
	if len(a) != 1 || a[0].Name != "etnaḥta" || !a[0].Disjunctive {
		t.Errorf("Accents = %v, want etnaḥta", a)
	}
}

func TestWords(t *testing.T) {
	w := Words("עַל־פְּנֵ֣י תְה֑וֹם ׀ ׃")
	if len(w) != 3 {
		t.Errorf("Words = %q, want 3 words", w)
	}
}
