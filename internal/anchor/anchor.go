// Package anchor ties each commentary comment to the span of tokens it
// explains. A comment opens with a quotation of the text (the dibbur
// ha-matḥil), which Sefaria sets in bold; the quotation is matched against
// the consonants of the tokens of the verse it is filed under and of the
// verses near it. A comment that cannot be matched is anchored to its whole
// verse and marked so; no anchor is guessed.
package anchor

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/hebrew"
)

type Anchor struct {
	ID          string   `json:"id"`
	Commentator string   `json:"commentator"`
	Normative   bool     `json:"normative,omitempty"`
	Filed       string   `json:"filed"` // the ref Sefaria files it under
	Ref         string   `json:"ref"`   // the ref it is anchored to
	Tokens      []string `json:"tokens,omitempty"`
	DH          string   `json:"dh,omitempty"`
	Status      string   `json:"status"` // matched, shifted, verse, segment
	Score       float64  `json:"score,omitempty"`
	Text        string   `json:"text"`
}

var reBold = regexp.MustCompile(`(?s)<b>(.*?)</b>`)

type piece struct{ dh, body string }

// pieces splits a comment at each bold heading.
func pieces(h string) []piece {
	locs := reBold.FindAllStringSubmatchIndex(h, -1)
	if len(locs) == 0 {
		return []piece{{body: hebrew.StripHTML(h)}}
	}
	var out []piece
	if pre := hebrew.StripHTML(h[:locs[0][0]]); pre != "" {
		out = append(out, piece{body: pre})
	}
	for i, l := range locs {
		end := len(h)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}
		out = append(out, piece{dh: hebrew.StripHTML(h[l[2]:l[3]]), body: hebrew.StripHTML(h[l[1]:end])})
	}
	return out
}

var skipWord = map[string]bool{"וגו": true, "וכו": true, "כו": true, "וגומר": true, "וגו'": true, "וכו'": true}

func dhWords(dh string) []string {
	var out []string
	for _, w := range hebrew.Words(hebrew.Unpointed(dh)) {
		w = strings.Trim(w, ".,:;()[]")
		switch strings.TrimLeft(w, "ו") {
		case "ה'", "ה׳", "השם", "יי":
			w = "יהוה" // the Name, abbreviated
		}
		if skipWord[w] || skipWord[hebrew.Skeleton(w)] || hebrew.Skeleton(w) == "" {
			continue
		}
		if strings.ContainsAny(w, "'׳\"״") { // an abbreviation
			continue
		}
		out = append(out, w)
	}
	return out
}

// match finds where the quoted words q occur in toks: q[0] at a start
// token, each following word within maxGap tokens of the one before.
// Starts at or after cursor are tried first. It returns the matched token
// indices and the mean score.
func match(q []string, toks []corpus.Token, cursor int) ([]int, float64) {
	const maxGap = 3
	var best []int
	bestVal := 0.0
	for pass := 0; pass < 2 && best == nil; pass++ {
		for s := range toks {
			if pass == 0 && s < cursor {
				continue
			}
			sc := hebrew.Similar(q[0], toks[s].Text)
			if sc == 0 {
				continue
			}
			// A quoted word absent from the text (a variant reading) is
			// skipped; two absent words in a row end the run.
			idx, sum, at, misses := []int{s}, sc, s, 0
			for i := 1; i < len(q) && misses < 2; i++ {
				found := false
				for k := at + 1; k <= at+1+maxGap && k < len(toks); k++ {
					if x := hebrew.Similar(q[i], toks[k].Text); x > 0 {
						idx, sum, at, found = append(idx, k), sum+x, k, true
						break
					}
				}
				if found {
					misses = 0
				} else {
					misses++
				}
			}
			cover := float64(len(idx)) / float64(len(q))
			if len(q) >= 2 && len(idx) < 2 || cover < 0.5 {
				continue
			}
			mean := sum / float64(len(idx))
			val := cover + mean/10
			if mean >= 0.75 && val > bestVal {
				best, bestVal = idx, val
			}
		}
	}
	return best, bestVal
}

// leading proposes quotations for an unbolded comment: its words up to the
// first comma or period (at most five), then its first two words, then its
// first word. Ibn Ezra's "וטעם X" and "מלת X" introduce X.
func leading(body string) []string {
	f := strings.Fields(body)
	for len(f) > 0 {
		switch strings.Trim(f[0], ",.:") {
		case "וטעם", "טעם", "ומלת", "מלת":
			f = f[1:]
			continue
		}
		break
	}
	var out []string
	var head []string
	for _, w := range f {
		head = append(head, w)
		if strings.ContainsAny(w, ",.:") || len(head) == 5 {
			break
		}
	}
	for n := len(head); n >= 1; n-- {
		if n == len(head) || n <= 2 {
			out = append(out, strings.Join(head[:n], " "))
		}
	}
	return out
}

// Comments is one commentator's text for one chapter: verse → comments.
type Comments struct {
	Chapter int
	Verses  [][]string
}

// Build anchors one commentator's comments to the corpus.
func Build(c *corpus.Corpus, key, name string, normative bool, chapters []Comments) []Anchor {
	var out []Anchor
	for _, ch := range chapters {
		for vi, comments := range ch.Verses {
			filed := fmt.Sprintf("%d:%d", ch.Chapter, vi+1)
			if c.Segment(filed) == nil {
				continue
			}
			n := 0
			cursor := map[string]int{}
			for _, h := range comments {
				ps := pieces(h)
				var cur *Anchor
				for pi, p := range ps {
					a, ok := place(c, filed, p, cursor)
					if !ok && pi > 0 && cur != nil {
						// A bold phrase inside a comment that is not a quotation.
						cur.Text += " " + p.dh + " " + p.body
						continue
					}
					n++
					a.ID = fmt.Sprintf("%s:%s:%d", key, filed, n)
					a.Commentator = name
					a.Normative = normative
					a.Filed = filed
					out = append(out, a)
					cur = &out[len(out)-1]
				}
			}
		}
	}
	return out
}

func place(c *corpus.Corpus, filed string, p piece, cursor map[string]int) (Anchor, bool) {
	var cands []string
	if p.dh != "" {
		cands = []string{p.dh}
		// A heading may join two quotations: "X. Y."; the first then stands alone.
		if i := strings.IndexAny(strings.TrimRight(p.dh, ". "), "."); i > 0 {
			cands = append(cands, p.dh[:i])
		}
	} else {
		cands = leading(p.body)
	}
	for _, dh := range cands {
		if a, ok := placeOne(c, filed, p, dh, cursor); ok {
			return a, true
		}
	}
	a := Anchor{Ref: filed, Text: p.body, DH: p.dh, Status: "segment"}
	if p.dh != "" {
		a.Status = "verse"
	}
	return a, false
}

func placeOne(c *corpus.Corpus, filed string, p piece, dh string, cursor map[string]int) (Anchor, bool) {
	a := Anchor{Ref: filed, Text: p.body, DH: p.dh}
	// A heading may quote several phrases: "ויברך, ויקדש".
	var parts [][]string
	for _, part := range strings.FieldsFunc(dh, func(r rune) bool { return r == ',' || r == ';' || r == '.' }) {
		if q := dhWords(part); len(q) > 0 {
			parts = append(parts, q)
		}
	}
	if len(parts) == 0 {
		return a, false
	}
	ch, v := 0, 0
	fmt.Sscanf(filed, "%d:%d", &ch, &v)
	for _, d := range []int{0, -1, 1, -2, 2} {
		ref := fmt.Sprintf("%d:%d", ch, v+d)
		seg := c.Segment(ref)
		if seg == nil {
			continue
		}
		var idx []int
		total, ok := 0.0, true
		cur := cursor[ref]
		for _, q := range parts {
			m, sc := match(q, seg.Tokens, cur)
			if m == nil {
				ok = false
				break
			}
			idx = append(idx, m...)
			total += sc
			cur = m[0]
		}
		if !ok {
			continue
		}
		sort.Ints(idx)
		a.Ref = ref
		a.Score = total / float64(len(parts))
		a.Status = "matched"
		if d != 0 {
			a.Status = "shifted"
		}
		if p.dh == "" {
			a.DH = dh
			a.Status += "-unbolded"
		}
		seen := map[int]bool{}
		for _, i := range idx {
			if !seen[i] {
				seen[i] = true
				a.Tokens = append(a.Tokens, seg.Tokens[i].ID)
			}
		}
		cursor[ref] = idx[0]
		return a, true
	}
	return a, false
}
