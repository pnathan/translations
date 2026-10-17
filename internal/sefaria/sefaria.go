// Package sefaria fetches texts and lexicon entries from the Sefaria API
// and keeps every response, unchanged, in a raw directory with a manifest
// (URL, fetch date, SHA-256). A response that is already present is read
// from disk; nothing is fetched twice unless the caller asks for it.
package sefaria

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const Base = "https://www.sefaria.org"

// Record is one manifest line.
type Record struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Fetched string `json:"fetched"`
	SHA256  string `json:"sha256"`
	Bytes   int    `json:"bytes"`
}

// Store is a raw directory with its manifest.
type Store struct {
	Dir      string
	Refresh  bool
	manifest map[string]Record
	client   *http.Client
	last     time.Time
}

func Open(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	s := &Store{Dir: dir, manifest: map[string]Record{}, client: &http.Client{Timeout: 120 * time.Second}}
	b, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err == nil {
		var recs []Record
		if err := json.Unmarshal(b, &recs); err != nil {
			return nil, fmt.Errorf("manifest: %w", err)
		}
		for _, r := range recs {
			s.manifest[r.Name] = r
		}
	}
	return s, nil
}

// Save writes the manifest, sorted by name.
func (s *Store) Save() error {
	var recs []Record
	for _, r := range s.manifest {
		recs = append(recs, r)
	}
	sort.Slice(recs, func(i, j int) bool { return recs[i].Name < recs[j].Name })
	b, _ := json.MarshalIndent(recs, "", "  ")
	return os.WriteFile(filepath.Join(s.Dir, "manifest.json"), append(b, '\n'), 0o644)
}

// Get returns the body stored under name, fetching u if it is absent.
func (s *Store) Get(name, u string) ([]byte, error) {
	p := filepath.Join(s.Dir, name)
	if !s.Refresh {
		if b, err := os.ReadFile(p); err == nil {
			return b, nil
		}
	}
	if d := time.Since(s.last); d < 300*time.Millisecond {
		time.Sleep(300*time.Millisecond - d)
	}
	var (
		b   []byte
		err error
	)
	for attempt, wait := 0, 2*time.Second; attempt < 4; attempt, wait = attempt+1, wait*2 {
		b, err = s.fetch(u)
		if err == nil {
			break
		}
		time.Sleep(wait)
	}
	s.last = time.Now()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return nil, err
	}
	if err := os.WriteFile(p, b, 0o644); err != nil {
		return nil, err
	}
	h := sha256.Sum256(b)
	s.manifest[name] = Record{Name: name, URL: u, Fetched: time.Now().UTC().Format("2006-01-02"), SHA256: hex.EncodeToString(h[:]), Bytes: len(b)}
	return b, nil
}

func (s *Store) fetch(u string) ([]byte, error) {
	resp, err := s.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GET %s: %s", u, resp.Status)
	}
	return b, nil
}

// TextURL is the v3 texts URL for one ref and one Hebrew version.
func TextURL(ref, version string) string {
	return Base + "/api/v3/texts/" + url.PathEscape(ref) + "?version=" + url.QueryEscape("hebrew|"+version) + "&return_format=default"
}

// WordsURL is the lexicon lookup URL for one word.
func WordsURL(word string) string {
	return Base + "/api/words/" + url.PathEscape(word) + "?never_split=1"
}

// Version is the part of a v3 response the pipeline reads.
type Version struct {
	VersionTitle string          `json:"versionTitle"`
	Language     string          `json:"language"`
	License      string          `json:"license"`
	VersionSrc   string          `json:"versionSource"`
	Text         json.RawMessage `json:"text"`
}

// ParseText returns the single Hebrew version of a v3 response. It fails if
// the response carries any other version or a non-Hebrew one.
func ParseText(b []byte, want string) (*Version, error) {
	var r struct {
		Versions []Version `json:"versions"`
		Error    string    `json:"error"`
	}
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Error != "" {
		return nil, fmt.Errorf("sefaria: %s", r.Error)
	}
	if len(r.Versions) != 1 {
		return nil, fmt.Errorf("sefaria: %d versions in response, want 1", len(r.Versions))
	}
	v := &r.Versions[0]
	if v.Language != "he" {
		return nil, fmt.Errorf("sefaria: version language %q, want he", v.Language)
	}
	if want != "" && v.VersionTitle != want {
		return nil, fmt.Errorf("sefaria: version %q, want %q", v.VersionTitle, want)
	}
	return v, nil
}

// Segments decodes the text field of a one-chapter response as segment →
// strings. A base text has one string per segment; a commentary has a list
// of comments per verse. A deeper level (sub-paragraphs) is joined.
func Segments(raw json.RawMessage) ([][]string, error) {
	var segs []json.RawMessage
	if err := json.Unmarshal(raw, &segs); err != nil {
		var one string
		if json.Unmarshal(raw, &one) == nil {
			return [][]string{{one}}, nil
		}
		return nil, err
	}
	out := make([][]string, len(segs))
	for j, sg := range segs {
		out[j] = flatten(sg)
	}
	return out, nil
}

func flatten(m json.RawMessage) []string {
	var s string
	if json.Unmarshal(m, &s) == nil {
		if s == "" {
			return nil
		}
		return []string{s}
	}
	var many []json.RawMessage
	if json.Unmarshal(m, &many) != nil {
		return nil
	}
	var out []string
	for _, x := range many {
		var one string
		if json.Unmarshal(x, &one) == nil {
			if one != "" {
				out = append(out, one)
			}
			continue
		}
		if sub := flatten(x); len(sub) > 0 {
			out = append(out, strings.Join(sub, " "))
		}
	}
	return out
}
