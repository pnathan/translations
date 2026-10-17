package llm

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Apply returns the translation with the review's patches applied. Morph
// patches are skipped here; LoadMorph applies them.
func Apply(tu *TranslationUnit, ru *ReviewUnit) (*TranslationUnit, error) {
	b, _ := json.Marshal(tu)
	var out TranslationUnit
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	out.Model = tu.Model + ", reviewed by " + ru.Model
	for _, p := range ru.Patches {
		if strings.HasPrefix(p.Path, "morph:") {
			continue
		}
		var seg *SegmentTr
		for i := range out.Segments {
			if out.Segments[i].Ref == p.Ref {
				seg = &out.Segments[i]
			}
		}
		if seg == nil {
			return nil, fmt.Errorf("patch %s %s: no such segment", p.Ref, p.Path)
		}
		if err := applyOne(seg, p); err != nil {
			return nil, fmt.Errorf("patch %s %s: %w", p.Ref, p.Path, err)
		}
	}
	return &out, nil
}

func applyOne(seg *SegmentTr, p Patch) error {
	str := func() (string, error) {
		var s string
		err := json.Unmarshal(p.Value, &s)
		return s, err
	}
	var err error
	switch p.Path {
	case "english":
		seg.English, err = str()
		return err
	case "literal":
		seg.Literal, err = str()
		return err
	case "note":
		seg.Note, err = str()
		return err
	case "grammar":
		return json.Unmarshal(p.Value, &seg.Grammar)
	case "phrases":
		return json.Unmarshal(p.Value, &seg.Phrases)
	case "apparatus":
		return json.Unmarshal(p.Value, &seg.Apparatus)
	}
	if rest, ok := strings.CutPrefix(p.Path, "token:"); ok {
		// Token ids contain "." ("1:1.2"): the field follows the last one.
		i := strings.LastIndex(rest, ".")
		if i < 0 {
			return fmt.Errorf("want token:<id>.<field>")
		}
		id, field := rest[:i], rest[i+1:]
		var g *TokenGloss
		for i := range seg.Tokens {
			if seg.Tokens[i].ID == id {
				g = &seg.Tokens[i]
			}
		}
		if g == nil {
			return fmt.Errorf("no token %s", id)
		}
		s, err := str()
		if err != nil {
			return err
		}
		switch field {
		case "gloss":
			g.Gloss = s
		case "lex":
			g.Lex = s
		case "sense":
			g.Sense = s
		case "note":
			g.Note = s
		default:
			return fmt.Errorf("unknown token field %q", field)
		}
		return nil
	}
	if rest, ok := strings.CutPrefix(p.Path, "commentary:"); ok {
		id, field, hasField := strings.Cut(rest, ".")
		var c *CommentUse
		for i := range seg.Commentary {
			if seg.Commentary[i].Anchor == id {
				c = &seg.Commentary[i]
			}
		}
		if !hasField {
			var nc CommentUse
			if err := json.Unmarshal(p.Value, &nc); err != nil {
				return err
			}
			nc.Anchor = id
			if c != nil {
				*c = nc
			} else {
				seg.Commentary = append(seg.Commentary, nc)
			}
			return nil
		}
		if c == nil {
			return fmt.Errorf("no commentary entry %s (add it with path commentary:%s)", id, id)
		}
		s, err := str()
		if err != nil {
			return err
		}
		switch field {
		case "relation":
			c.Relation = s
		case "summary":
			c.Summary = s
		default:
			return fmt.Errorf("unknown commentary field %q", field)
		}
		return nil
	}
	return fmt.Errorf("unknown path")
}
