// Package work reads a work's configuration (works/<slug>/work.json) and
// names the paths of its stages.
package work

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Kind of text: "tanakh" has OSHB morphology; "mishnah" has none, and its
// morphology is a model stage reviewed like a translation.
type Kind string

const (
	Tanakh  Kind = "tanakh"
	Mishnah Kind = "mishnah"
)

type Text struct {
	Sefaria string `json:"sefaria"` // index title, e.g. "Genesis", "Pirkei_Avot"
	Version string `json:"version"` // Hebrew versionTitle, pinned
	Name    string `json:"name"`    // short name for the apparatus
}

type Commentator struct {
	Key       string `json:"key"`     // "rashi"
	Name      string `json:"name"`    // "Rashi"
	He        string `json:"he"`      // "רש״י"
	Sefaria   string `json:"sefaria"` // "Rashi_on_Genesis"
	Version   string `json:"version"`
	Normative bool   `json:"normative"`    // the review must rule on every comment
	Targum    bool   `json:"version_text"` // a translation (targum): one entry per verse, not quote-matched
}

type Section struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	He   string `json:"he"`
	From string `json:"from"` // "1:1"
	To   string `json:"to"`
}

type Config struct {
	Slug         string        `json:"slug"`
	Title        string        `json:"title"`
	He           string        `json:"he"`
	Kind         Kind          `json:"kind"`
	RefName      string        `json:"refName"` // English book name for refs, "Genesis"
	OSIS         string        `json:"osis,omitempty"`
	OSHB         string        `json:"oshb,omitempty"` // URL of the OSHB book file
	Chapters     int           `json:"chapters"`       // number of chapters, from 1
	Text         Text          `json:"text"`
	Witnesses    []Text        `json:"witnesses,omitempty"` // further Sefaria versions, collated
	Commentators []Commentator `json:"commentators"`
	Sections     []Section     `json:"sections"`
	UnitMax      int           `json:"unitMax"` // most segments in one unit
	Segment      string        `json:"segment"` // "verse" or "mishnah"
	Dir          string        `json:"-"`
}

func Load(root, slug string) (*Config, error) {
	dir := filepath.Join(root, "works", slug)
	b, err := os.ReadFile(filepath.Join(dir, "work.json"))
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("%s/work.json: %w", slug, err)
	}
	c.Dir = dir
	return &c, nil
}

func (c *Config) Path(parts ...string) string {
	return filepath.Join(append([]string{c.Dir}, parts...)...)
}

// ParseRef reads "12:3".
func ParseRef(r string) (ch, v int, err error) {
	a, b, ok := strings.Cut(r, ":")
	if !ok {
		return 0, 0, fmt.Errorf("bad ref %q", r)
	}
	if ch, err = strconv.Atoi(a); err != nil {
		return
	}
	v, err = strconv.Atoi(b)
	return
}

// Less orders refs.
func Less(a, b string) bool {
	ac, av, _ := ParseRef(a)
	bc, bv, _ := ParseRef(b)
	return ac < bc || ac == bc && av < bv
}

// SectionOf names the section that holds ref.
func (c *Config) SectionOf(ref string) *Section {
	for i := range c.Sections {
		s := &c.Sections[i]
		if !Less(ref, s.From) && !Less(s.To, ref) {
			return s
		}
	}
	return nil
}
