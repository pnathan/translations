// Package oshb reads the Open Scriptures Hebrew Bible (OSIS XML: the
// Westminster Leningrad Codex with lemma and morphology) and decodes its
// morphology codes.
package oshb

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Word is one <w> element of the WLC.
type Word struct {
	ID         string // immutable OSHB word id
	Text       string // pointed surface; "/" separates morphemes
	Lemma      string // e.g. "c/b/929", "1254 a"
	Morph      string // e.g. "HC/Rd/Ncfsa"
	Ketiv      bool
	Qere       *Word // the qere of a ketiv word
	MaqafAfter bool
	PaseqAfter bool
	Notes      []string // WLC textual notes on this word
}

// Verse is one verse of the WLC.
type Verse struct {
	OsisID string // "Gen.1.1"
	Words  []Word
	Para   string // "פ" or "ס" after the verse, if the WLC marks one
}

// Read parses an OSHB book file. Verses come in document order.
func Read(r io.Reader) ([]Verse, error) {
	dec := xml.NewDecoder(r)
	var (
		verses []Verse
		cur    *Verse
		inW    *Word
		wbuf   strings.Builder
		inNote bool
		noteN  string
		nbuf   strings.Builder
		inRdg  bool
		inSeg  string
	)
	attr := func(e xml.StartElement, k string) string {
		for _, a := range e.Attr {
			if a.Name.Local == k {
				return a.Value
			}
		}
		return ""
	}
	last := func() *Word {
		if cur == nil || len(cur.Words) == 0 {
			return nil
		}
		return &cur.Words[len(cur.Words)-1]
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "verse":
				if id := attr(t, "osisID"); id != "" {
					verses = append(verses, Verse{OsisID: id})
					cur = &verses[len(verses)-1]
				}
			case "w":
				inW = &Word{ID: attr(t, "id"), Lemma: attr(t, "lemma"), Morph: attr(t, "morph"), Ketiv: attr(t, "type") == "x-ketiv"}
				wbuf.Reset()
			case "note":
				inNote = true
				noteN = attr(t, "n")
				nbuf.Reset()
			case "rdg":
				inRdg = attr(t, "type") == "x-qere"
			case "seg":
				inSeg = attr(t, "type")
			}
		case xml.CharData:
			switch {
			case inW != nil:
				wbuf.Write(t)
			case inNote && !inRdg:
				nbuf.Write(t)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "w":
				inW.Text = wbuf.String()
				if inRdg {
					if lw := last(); lw != nil {
						q := *inW
						lw.Qere = &q
					}
				} else if cur != nil {
					cur.Words = append(cur.Words, *inW)
				}
				inW = nil
			case "rdg":
				inRdg = false
			case "note":
				inNote = false
				if noteN != "" {
					if lw := last(); lw != nil {
						lw.Notes = append(lw.Notes, fmt.Sprintf("WLC note %s: %s", noteN, strings.TrimSpace(nbuf.String())))
					}
				}
			case "seg":
				if lw := last(); lw != nil {
					switch inSeg {
					case "x-maqqef":
						lw.MaqafAfter = true
					case "x-paseq":
						lw.PaseqAfter = true
					}
				}
				if cur != nil {
					switch inSeg {
					case "x-pe":
						cur.Para = "פ"
					case "x-samekh":
						cur.Para = "ס"
					}
				}
				inSeg = ""
			case "verse":
				cur = nil
			}
		}
	}
	return verses, nil
}
