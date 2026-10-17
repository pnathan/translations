// Package corpus builds the token stream of a work: segments (verses or
// mishnayot) of tokens with fixed ids, each token carrying its morphology,
// lexicon key and accents, and the variants between the witnesses.
package corpus

import (
	"encoding/json"
	"fmt"
	"os"
)

// Morpheme is one part of a token: a prefix, the stem word, or a suffix.
type Morpheme struct {
	Form  string `json:"form"`
	Lemma string `json:"lemma,omitempty"` // Hebrew lemma, or the prefix gloss
	Lex   string `json:"lex,omitempty"`   // lexicon key
	POS   string `json:"pos,omitempty"`
	Parse string `json:"parse,omitempty"` // English description
	Code  string `json:"code,omitempty"`  // OSHB code, when the source is OSHB
}

// Token is one written word (maqaf-joined words are separate tokens).
type Token struct {
	ID          string     `json:"id"`            // "1:1.3": ref, then 1-based position
	Text        string     `json:"text"`          // display form (the base text)
	Alt         string     `json:"alt,omitempty"` // the second witness's form, where it differs
	OSHB        string     `json:"oshb,omitempty"`
	Lang        string     `json:"lang,omitempty"` // hbo, arc, mishnaic
	Morphemes   []Morpheme `json:"morphemes,omitempty"`
	Lex         string     `json:"lex,omitempty"` // lexicon key of the main morpheme
	Root        string     `json:"root,omitempty"`
	Accents     []string   `json:"accents,omitempty"`
	Disjunctive bool       `json:"disjunctive,omitempty"`
	Maqaf       bool       `json:"maqaf,omitempty"` // joined by maqaf to the next token
	Punct       string     `json:"punct,omitempty"` // punctuation after the token (mishnah)
	Ketiv       string     `json:"ketiv,omitempty"` // written form, where the reading (Text) is the qere
	Notes       []string   `json:"notes,omitempty"`
	MorphSource string     `json:"morphSource,omitempty"` // "OSHB" or "model:<name>, reviewed"
}

// Variant records a difference between witnesses.
type Variant struct {
	Tokens  []string `json:"tokens,omitempty"`
	Witness string   `json:"witness"`
	Reading string   `json:"reading"`
	Base    string   `json:"base"`
	Kind    string   `json:"kind"` // consonants, vowels, accents, ketiv-qere, word
}

// Segment is a verse or a mishnah.
type Segment struct {
	Ref      string    `json:"ref"` // "1:1"
	Section  string    `json:"section"`
	Unit     string    `json:"unit"`
	Text     string    `json:"text"` // the base text, tags removed
	Tokens   []Token   `json:"tokens"`
	Para     string    `json:"para,omitempty"` // פ or ס after the verse
	Variants []Variant `json:"variants,omitempty"`
}

// Unit is the span one translator call handles.
type Unit struct {
	ID      string   `json:"id"`
	Section string   `json:"section"`
	Refs    []string `json:"refs"`
}

type Corpus struct {
	Work     string    `json:"work"`
	Units    []Unit    `json:"units"`
	Segments []Segment `json:"segments"`
}

func (c *Corpus) Segment(ref string) *Segment {
	for i := range c.Segments {
		if c.Segments[i].Ref == ref {
			return &c.Segments[i]
		}
	}
	return nil
}

func (c *Corpus) Unit(id string) *Unit {
	for i := range c.Units {
		if c.Units[i].ID == id {
			return &c.Units[i]
		}
	}
	return nil
}

func Load(path string) (*Corpus, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Corpus
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &c, nil
}

func WriteJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

func ReadJSON(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}
