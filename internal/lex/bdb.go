// Package lex reads the OpenScriptures Hebrew Lexicon (BDB with Strong's
// numbers and a lexical index) and holds the lexicon entries the pipeline
// attaches to tokens.
package lex

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Sense is one numbered sense, flattened. Path is "1", "1a", "Niph 2", …
type Sense struct {
	Path  string `json:"path"`
	Stem  string `json:"stem,omitempty"`
	Gloss string `json:"gloss"`
}

// Entry is a lexicon entry as the pipeline stores it.
type Entry struct {
	Key      string  `json:"key"` // e.g. "bdb:b.cw.aa", "jastrow:קָבַל I"
	Source   string  `json:"source"`
	Headword string  `json:"headword"`
	Root     string  `json:"root,omitempty"`
	POS      string  `json:"pos,omitempty"`
	Gloss    string  `json:"gloss,omitempty"` // short definition
	Senses   []Sense `json:"senses,omitempty"`
	Strong   string  `json:"strong,omitempty"`
}

// Index is the joined OpenScriptures lexicon.
type Index struct {
	li    map[string]*liEntry // lexical-index id → entry
	aug   map[string]string   // augmented Strong ("1254a") → lexical-index id
	bdb   map[string]*bdbEntry
	byStr map[string][]string // Strong number → lexical-index ids
}

type liEntry struct {
	ID, W, POS, Def, BDB, Strong, Aug string
	EtymType, EtymRoot, EtymText      string
}

type bdbEntry struct {
	ID, Head, POS, Def string
	Senses             []Sense
}

// Load reads LexicalIndex.xml, AugIndex.xml and BrownDriverBriggs.xml.
func Load(li, aug, bdb io.Reader) (*Index, error) {
	ix := &Index{li: map[string]*liEntry{}, aug: map[string]string{}, bdb: map[string]*bdbEntry{}, byStr: map[string][]string{}}
	if err := ix.readLI(li); err != nil {
		return nil, fmt.Errorf("lexical index: %w", err)
	}
	if err := ix.readAug(aug); err != nil {
		return nil, fmt.Errorf("aug index: %w", err)
	}
	if err := ix.readBDB(bdb); err != nil {
		return nil, fmt.Errorf("bdb: %w", err)
	}
	return ix, nil
}

func (ix *Index) readLI(r io.Reader) error {
	type xref struct {
		BDB    string `xml:"bdb,attr"`
		Strong string `xml:"strong,attr"`
		Aug    string `xml:"aug,attr"`
	}
	type etym struct {
		Type string `xml:"type,attr"`
		Root string `xml:"root,attr"`
		Text string `xml:",chardata"`
	}
	type entry struct {
		ID   string `xml:"id,attr"`
		W    string `xml:"w"`
		POS  string `xml:"pos"`
		Def  string `xml:"def"`
		Xref xref   `xml:"xref"`
		Etym etym   `xml:"etym"`
	}
	dec := xml.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if se, ok := tok.(xml.StartElement); ok && se.Name.Local == "entry" {
			var e entry
			if err := dec.DecodeElement(&e, &se); err != nil {
				return err
			}
			le := &liEntry{ID: e.ID, W: e.W, POS: e.POS, Def: e.Def, BDB: e.Xref.BDB, Strong: e.Xref.Strong, Aug: e.Xref.Aug,
				EtymType: e.Etym.Type, EtymRoot: e.Etym.Root, EtymText: strings.TrimSpace(e.Etym.Text)}
			ix.li[e.ID] = le
			if le.Strong != "" {
				ix.byStr[le.Strong] = append(ix.byStr[le.Strong], e.ID)
			}
		}
	}
}

func (ix *Index) readAug(r io.Reader) error {
	type w struct {
		Aug string `xml:"aug,attr"`
		ID  string `xml:",chardata"`
	}
	var doc struct {
		W []w `xml:"w"`
	}
	if err := xml.NewDecoder(r).Decode(&doc); err != nil {
		return err
	}
	for _, x := range doc.W {
		ix.aug[x.Aug] = strings.TrimSpace(x.ID)
	}
	return nil
}

var reWS = regexp.MustCompile(`\s+`)

func clean(s string) string { return strings.TrimSpace(reWS.ReplaceAllString(s, " ")) }

// reCount matches the occurrence counts that open a sense ("1052 ", not "1st").
var reCount = regexp.MustCompile(`^(?:\d+(?:[\s.,;:]+|$))+`)

// readBDB walks the BDB file. For each entry it keeps the head word, the
// part of speech, the <def> glosses outside any sense, and every <sense>
// flattened with its path and stem. A sense with no <def> of its own (the
// abridged file gives שָׁמַע Qal only "1 subj. pers. hum.") is kept when it
// has a number or no subsenses: its gloss is the nearest definition above
// it, then its own text; a derived stem that has none of its own names the
// entry's definition as the one it derives from ("subj. God (Hiph of: die)").
func (ix *Index) readBDB(r io.Reader) error {
	dec := xml.NewDecoder(r)
	type frame struct {
		n, stem string
		defs    []string
		text    strings.Builder // character data outside <stem>, not in a subsense
		kids    bool
	}
	var (
		cur     *bdbEntry
		stack   []*frame
		inDef   bool
		inW     bool
		inPOS   bool
		inStem  bool
		buf     strings.Builder
		topDefs []string
	)
	attr := func(e xml.StartElement, k string) string {
		for _, a := range e.Attr {
			if a.Name.Local == k {
				return a.Value
			}
		}
		return ""
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "entry":
				cur = &bdbEntry{ID: attr(t, "id")}
				stack, topDefs = nil, nil
			case "sense":
				if cur != nil {
					if len(stack) > 0 {
						stack[len(stack)-1].kids = true
					}
					stack = append(stack, &frame{n: attr(t, "n")})
				}
			case "def", "w", "pos", "stem":
				buf.Reset()
				switch t.Name.Local {
				case "def":
					inDef = true
				case "w":
					inW = cur != nil && cur.Head == ""
				case "pos":
					inPOS = true
				case "stem":
					inStem = true
				}
			}
		case xml.CharData:
			if inDef || inW || inPOS || inStem {
				buf.Write(t)
			}
			if len(stack) > 0 && !inStem {
				stack[len(stack)-1].text.Write(t)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "def":
				inDef = false
				d := clean(buf.String())
				if len(stack) > 0 {
					stack[len(stack)-1].defs = append(stack[len(stack)-1].defs, d)
				} else {
					topDefs = append(topDefs, d)
				}
			case "w":
				if inW && cur != nil {
					cur.Head = clean(buf.String())
				}
				inW = false
			case "pos":
				if cur != nil && cur.POS == "" {
					cur.POS = clean(buf.String())
				}
				inPOS = false
			case "stem":
				if len(stack) > 0 {
					stack[len(stack)-1].stem = clean(buf.String())
				}
				inStem = false
			case "sense":
				if cur == nil || len(stack) == 0 {
					break
				}
				f := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				var p []string
				stem := ""
				for _, g := range stack {
					if g.stem != "" {
						stem = g.stem
					}
					if g.n != "" {
						p = append(p, g.n)
					}
				}
				if f.stem != "" {
					stem = f.stem
				}
				if f.n != "" {
					p = append(p, f.n)
				}
				gloss := strings.Join(f.defs, ", ")
				if gloss == "" && (f.n != "" || !f.kids) {
					own := strings.TrimRight(reCount.ReplaceAllString(clean(f.text.String()), ""), " .,;:")
					base, inherited := strings.Join(topDefs, ", "), true
					for _, g := range stack {
						if len(g.defs) > 0 {
							base, inherited = strings.Join(g.defs, ", "), false
						}
					}
					switch {
					case own != "" && base != "" && inherited && stem != "" && stem != "Qal":
						// The entry's definition is the Qal's: name it as such.
						gloss = own + " (" + stem + " of: " + base + ")"
					case own != "" && base != "":
						gloss = base + " — " + own
					case own != "":
						gloss = own
					}
				}
				if gloss != "" {
					path := strings.Join(p, ".")
					if stem != "" {
						path = strings.TrimSpace(stem + " " + path)
					}
					cur.Senses = append(cur.Senses, Sense{Path: path, Stem: stem, Gloss: gloss})
				}
			case "entry":
				if cur != nil {
					uniquePaths(cur.Senses)
					cur.Def = strings.Join(topDefs, ", ")
					ix.bdb[cur.ID] = cur
				}
				cur = nil
			}
		}
	}
}

// root follows the etymology links of a lexical-index entry to its root.
func (ix *Index) root(id string, depth int) string {
	e := ix.li[id]
	if e == nil || depth > 4 {
		return ""
	}
	if e.EtymRoot != "" {
		return e.EtymRoot
	}
	if e.EtymType == "sub" && e.EtymText != "" {
		return ix.root(strings.Fields(strings.ReplaceAll(e.EtymText, ",", " "))[0], depth+1)
	}
	return ""
}

// Lookup resolves an OSHB lemma ("1254 a", "430", "1254+") to an entry.
func (ix *Index) Lookup(lemma string) (*Entry, bool) {
	l := strings.TrimSpace(strings.TrimSuffix(lemma, "+"))
	key := strings.ReplaceAll(l, " ", "")
	id, ok := ix.aug[key]
	if !ok {
		num := strings.Fields(l)
		if len(num) == 0 {
			return nil, false
		}
		ids := ix.byStr[num[0]]
		if len(ids) == 0 {
			return nil, false
		}
		id = ids[0]
	}
	le := ix.li[id]
	if le == nil {
		return nil, false
	}
	e := &Entry{Key: "bdb:" + le.BDB, Source: "BDB (OpenScriptures Hebrew Lexicon)", Headword: le.W, POS: le.POS, Gloss: le.Def, Root: ix.root(id, 0), Strong: le.Strong}
	if le.BDB == "" {
		e.Key = "li:" + id
	}
	if b := ix.bdb[le.BDB]; b != nil {
		if b.POS != "" {
			e.POS = b.POS
		}
		if b.Def != "" {
			e.Gloss = b.Def
		}
		e.Senses = b.Senses
	}
	return e, true
}

// Head returns BDB's headword for an entry key, or "".
func (ix *Index) Head(key string) string {
	if b := ix.bdb[strings.TrimPrefix(key, "bdb:")]; b != nil {
		return b.Head
	}
	return ""
}

// uniquePaths gives every sense a distinct path: an empty or repeated path
// takes a "#k" suffix, k its position from 1.
func uniquePaths(ss []Sense) {
	seen := map[string]bool{}
	for i := range ss {
		if ss[i].Path == "" || seen[ss[i].Path] {
			ss[i].Path = fmt.Sprintf("%s#%d", ss[i].Path, i+1)
		}
		seen[ss[i].Path] = true
	}
}
