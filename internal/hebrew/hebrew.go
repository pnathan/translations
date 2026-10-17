// Package hebrew holds text operations on pointed Hebrew: consonant
// skeletons for matching, word splitting, HTML removal, and the names of
// the cantillation marks.
//
// No function here applies Unicode normalization to stored text. NFC
// reorders Hebrew combining marks (see the OSHB README); skeletons are
// computed from the consonants alone and need no normalization.
package hebrew

import (
	"html"
	"regexp"
	"strings"
	"unicode"
)

const (
	Maqaf     = '־'
	Paseq     = '׀'
	SofPasuq  = '׃'
	Meteg     = 'ֽ'
	Geresh    = '׳'
	Gershayim = '״'
)

// IsConsonant reports whether r is one of the 27 Hebrew letter forms.
func IsConsonant(r rune) bool { return r >= 'א' && r <= 'ת' }

// IsMark reports whether r is a point or accent (U+0591–U+05C7), other than
// the maqaf, paseq and sof pasuq, which are punctuation.
func IsMark(r rune) bool {
	if r < '֑' || r > 'ׇ' {
		return false
	}
	switch r {
	case Maqaf, Paseq, SofPasuq, '׆':
		return false
	}
	return true
}

var finals = map[rune]rune{'ך': 'כ', 'ם': 'מ', 'ן': 'נ', 'ף': 'פ', 'ץ': 'צ'}

// Skeleton returns the consonants of s, final forms mapped to medial forms.
func Skeleton(s string) string {
	var b strings.Builder
	for _, r := range s {
		if !IsConsonant(r) {
			continue
		}
		if m, ok := finals[r]; ok {
			r = m
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Loose returns the skeleton without the vowel letters ו and י, except in
// first position. Plene and defective spellings of one word agree under it.
func Loose(s string) string {
	sk := []rune(Skeleton(s))
	var b strings.Builder
	for i, r := range sk {
		if i > 0 && (r == 'ו' || r == 'י') {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Unpointed removes points and accents and keeps everything else.
func Unpointed(s string) string {
	var b strings.Builder
	for _, r := range s {
		if IsMark(r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Words splits s at white space and at the maqaf. Tokens with no Hebrew
// letter (paseq, punctuation) are dropped.
func Words(s string) []string {
	f := strings.FieldsFunc(s, func(r rune) bool { return unicode.IsSpace(r) || r == Maqaf })
	var out []string
	for _, w := range f {
		if Skeleton(w) != "" {
			out = append(out, w)
		}
	}
	return out
}

var (
	reFootnote = regexp.MustCompile(`(?s)<sup[^>]*class="footnote-marker"[^>]*>.*?</sup>\s*<i[^>]*class="footnote"[^>]*>.*?</i>`)
	reSuperRef = regexp.MustCompile(`<i data-commentator[^>]*></i>`)
	reTag      = regexp.MustCompile(`<[^>]+>`)
	reSpace    = regexp.MustCompile(`[ \t\r\n\x{00A0}\x{2009}]+`)
)

// StripHTML removes footnotes, super-commentary markers and tags, decodes
// entities and collapses white space.
func StripHTML(s string) string {
	s = reFootnote.ReplaceAllString(s, "")
	s = reSuperRef.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "<br>", " ")
	s = reTag.ReplaceAllString(s, "")
	s = html.UnescapeString(s)
	s = reSpace.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// Accent describes one cantillation mark.
type Accent struct {
	Name        string
	Disjunctive bool
}

var accents = map[rune]Accent{
	'֑': {"etnaḥta", true},
	'֒': {"segol", true},
	'֓': {"shalshelet", true},
	'֔': {"zaqef qaton", true},
	'֕': {"zaqef gadol", true},
	'֖': {"tipḥa", true},
	'֗': {"revia", true},
	'֘': {"zarqa", true},
	'֙': {"pashta", true},
	'֚': {"yetiv", true},
	'֛': {"tevir", true},
	'֜': {"geresh", true},
	'֝': {"geresh muqdam", true},
	'֞': {"gershayim", true},
	'֟': {"qarne para", true},
	'֠': {"telisha gedola", true},
	'֡': {"pazer", true},
	'֢': {"atnaḥ hafukh", false},
	'֣': {"munaḥ", false},
	'֤': {"mahpakh", false},
	'֥': {"merkha", false},
	'֦': {"merkha kefula", false},
	'֧': {"darga", false},
	'֨': {"qadma", false},
	'֩': {"telisha qetana", false},
	'֪': {"yeraḥ ben yomo", false},
	'֫': {"ole", true},
	'֬': {"iluy", false},
	'֭': {"deḥi", true},
	'֮': {"zinor", true},
}

// Accents returns the cantillation marks of a word in order of occurrence.
// last marks the final word of a verse: a meteg there is the silluq.
func Accents(word string, last bool) []Accent {
	var out []Accent
	for _, r := range word {
		if a, ok := accents[r]; ok {
			out = append(out, a)
		}
	}
	if last && strings.ContainsRune(word, Meteg) {
		out = append(out, Accent{"silluq", true})
	}
	return out
}

// Similar scores how well a quoted word q matches a text word t:
// 1 for equal skeletons, 0.9 for equal loose skeletons, 0.8 when one is the
// other with a one-letter proclitic, 0 otherwise.
func Similar(q, t string) float64 {
	qs, ts := Skeleton(q), Skeleton(t)
	if qs == "" || ts == "" {
		return 0
	}
	if qs == ts {
		return 1
	}
	if Loose(q) == Loose(t) {
		return 0.9
	}
	for _, p := range []string{"ו", "ה", "ב", "ל", "כ", "מ", "ש"} {
		if strings.TrimPrefix(qs, p) == ts || strings.TrimPrefix(ts, p) == qs {
			return 0.8
		}
		if Loose(strings.TrimPrefix(qs, p)) == Loose(ts) || Loose(strings.TrimPrefix(ts, p)) == Loose(qs) {
			return 0.7
		}
	}
	return 0
}
