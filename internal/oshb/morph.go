package oshb

import "strings"

// Decoding of the OSHB morphology codes (https://hb.openscriptures.org/parsing/HebrewMorphologyCodes.html).
// A code starts with a language letter (H Hebrew, A Aramaic); morphemes are
// separated by "/".

var pos = map[byte]string{
	'A': "adjective", 'C': "conjunction", 'D': "adverb", 'N': "noun", 'P': "pronoun",
	'R': "preposition", 'S': "suffix", 'T': "particle", 'V': "verb",
}

var hebStem = map[byte]string{
	'q': "qal", 'N': "nifal", 'p': "piel", 'P': "pual", 'h': "hifil", 'H': "hofal",
	't': "hitpael", 'o': "polel", 'O': "polal", 'r': "hitpolel", 'm': "poel", 'M': "poal",
	'k': "palel", 'K': "pulal", 'Q': "qal passive", 'l': "pilpel", 'L': "polpal",
	'f': "hitpalpel", 'D': "nitpael", 'j': "pealal", 'i': "pilel", 'u': "hotpaal",
	'c': "tifil", 'v': "hishtafel", 'w': "nitpalel", 'y': "nitpoel", 'z': "hitpoel",
}

var aramStem = map[byte]string{
	'q': "peal", 'Q': "peil", 'u': "hitpeel", 'p': "pael", 'P': "itpaal", 'M': "hitpaal",
	'a': "afel", 'h': "hafel", 's': "safel", 'e': "shafel", 'H': "hofal", 'i': "itpeel",
	't': "hishtafal", 'v': "ishtafal", 'w': "hitafal", 'o': "polel", 'z': "itpoel",
	'r': "hitpolel", 'f': "hitpalpel", 'b': "hefal", 'c': "tifel", 'm': "poel",
	'l': "palpel", 'L': "itpalpel", 'O': "itpolel", 'G': "ittafal",
}

var conj = map[byte]string{
	'p': "perfect (qatal)", 'q': "sequential perfect (weqatal)", 'i': "imperfect (yiqtol)",
	'w': "sequential imperfect (wayyiqtol)", 'h': "cohortative", 'j': "jussive",
	'v': "imperative", 'r': "participle active", 's': "participle passive",
	'a': "infinitive absolute", 'c': "infinitive construct",
}

var person = map[byte]string{'1': "1st person", '2': "2nd person", '3': "3rd person"}
var gender = map[byte]string{'m': "masculine", 'f': "feminine", 'b': "both genders", 'c': "common"}
var number = map[byte]string{'s': "singular", 'p': "plural", 'd': "dual"}
var state = map[byte]string{'a': "absolute", 'c': "construct", 'd': "determined"}

var nounType = map[byte]string{'c': "common", 'g': "gentilic", 'p': "proper name"}
var properType = map[byte]string{'m': "person, masculine", 'f': "person, feminine", 'l': "location", 't': "title"}
var adjType = map[byte]string{'a': "adjective", 'c': "cardinal number", 'g': "gentilic", 'o': "ordinal number"}
var pronType = map[byte]string{'d': "demonstrative", 'f': "indefinite", 'i': "interrogative", 'p': "personal", 'r': "relative"}
var sufType = map[byte]string{'d': "directional he", 'h': "paragogic he", 'n': "paragogic nun", 'p': "pronominal"}
var partType = map[byte]string{
	'a': "affirmation", 'd': "definite article", 'e': "exhortation", 'i': "interrogative",
	'j': "interjection", 'm': "demonstrative", 'n': "negative", 'o': "direct object marker", 'r': "relative",
}

// Morpheme is the decoded form of one morpheme code.
type Morpheme struct {
	Code  string // the morpheme's code without the language letter
	POS   string
	Parse string // English description
}

// Decode splits a full code ("HC/Rd/Ncfsa") into decoded morphemes. lang
// is "hbo" or "arc".
func Decode(code string) (lang string, ms []Morpheme) {
	if code == "" {
		return "", nil
	}
	lang = "hbo"
	stems := hebStem
	if code[0] == 'A' {
		lang = "arc"
		stems = aramStem
	}
	for _, c := range strings.Split(code[1:], "/") {
		ms = append(ms, decodeOne(c, stems))
	}
	return lang, ms
}

func decodeOne(c string, stems map[byte]string) Morpheme {
	m := Morpheme{Code: c}
	if c == "" {
		return m
	}
	m.POS = pos[c[0]]
	rest := c[1:]
	var parts []string
	at := func(i int, tbl map[byte]string) {
		if i < len(rest) {
			if s, ok := tbl[rest[i]]; ok {
				parts = append(parts, s)
			} else if rest[i] != 'x' {
				parts = append(parts, string(rest[i]))
			}
		}
	}
	switch c[0] {
	case 'V':
		at(0, stems)
		// Active and passive participle differ only in the qal (and the
		// Aramaic peal); in the other stems "r" is the participle.
		if len(rest) > 1 && rest[1] == 'r' && rest[0] != 'q' {
			parts = append(parts, "participle")
		} else {
			at(1, conj)
		}
		if len(rest) > 1 && (rest[1] == 'r' || rest[1] == 's') {
			at(2, gender)
			at(3, number)
			at(4, state)
		} else {
			at(2, person)
			at(3, gender)
			at(4, number)
		}
	case 'N':
		if len(rest) > 0 && rest[0] == 'p' {
			parts = append(parts, "proper name")
			at(1, properType)
		} else {
			at(0, nounType)
			at(1, gender)
			at(2, number)
			at(3, state)
		}
	case 'A':
		at(0, adjType)
		at(1, gender)
		at(2, number)
		at(3, state)
	case 'P':
		at(0, pronType)
		at(1, person)
		at(2, gender)
		at(3, number)
	case 'S':
		at(0, sufType)
		at(1, person)
		at(2, gender)
		at(3, number)
	case 'T':
		at(0, partType)
	case 'R':
		if rest == "d" {
			parts = append(parts, "with definite article")
		}
	}
	m.Parse = strings.Join(parts, ", ")
	return m
}

// Prefix glosses for the one-letter lemma codes of OSHB.
var PrefixLemma = map[string]string{
	"b": "ב: in, with, by", "c": "ו: and", "d": "ה: the", "i": "ה: interrogative",
	"k": "כ: like, as", "l": "ל: to, for", "m": "מ: from, than", "s": "ש: who, which",
}
