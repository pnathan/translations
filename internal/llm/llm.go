// Package llm defines the files the model stages write (morphology,
// translation, review), validates them against the corpus, and applies the
// review's patches. The model never edits the corpus; it writes these files
// and the pipeline checks and merges them.
package llm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/work"
)

// MorphUnit is the morphology stage's output for a mishnah unit.
type MorphUnit struct {
	Unit   string            `json:"unit"`
	Model  string            `json:"model"`
	Tokens []corpus.Analysis `json:"tokens"`
}

type TokenGloss struct {
	ID    string `json:"id"`
	Gloss string `json:"gloss"`           // sense in this context, English
	Lex   string `json:"lex,omitempty"`   // lexicon key, where the token has several candidates
	Sense string `json:"sense,omitempty"` // sense path within that entry
	Note  string `json:"note,omitempty"`  // a learner's note on this word
}

type Phrase struct {
	Tokens  []string `json:"tokens"`
	English string   `json:"english"` // the fragment of the reading translation; "-" for none
	Literal string   `json:"literal"` // the literal sense, and the reason for the English
}

type CommentUse struct {
	Anchor   string `json:"anchor"`
	Relation string `json:"relation"` // follows, records, declines, not-applicable
	Summary  string `json:"summary"`  // the comment's point, in English
}

type SegmentTr struct {
	Ref        string       `json:"ref"`
	English    string       `json:"english"`
	Literal    string       `json:"literal"`
	Tokens     []TokenGloss `json:"tokens"`
	Phrases    []Phrase     `json:"phrases"`
	Grammar    []string     `json:"grammar"`
	Note       string       `json:"note,omitempty"`
	Commentary []CommentUse `json:"commentary"`
	Apparatus  []string     `json:"apparatus,omitempty"`
}

type TranslationUnit struct {
	Unit     string      `json:"unit"`
	Model    string      `json:"model"`
	Segments []SegmentTr `json:"segments"`
}

// Patch replaces one field. Paths:
//
//	english | literal | note | grammar | phrases | apparatus
//	token:<id>.gloss | token:<id>.lex | token:<id>.sense | token:<id>.note
//	commentary:<anchor>.relation | commentary:<anchor>.summary | commentary:<anchor> (add)
//	morph:<id> (a whole analysis; mishnah only)
type Patch struct {
	Ref      string          `json:"ref"`
	Path     string          `json:"path"`
	Value    json.RawMessage `json:"value"`
	Severity string          `json:"severity"` // error, question, style
	Reason   string          `json:"reason"`
}

type Verdict struct {
	Anchor  string `json:"anchor"`
	Verdict string `json:"verdict"` // consistent, contradicts, not-applicable
	Comment string `json:"comment,omitempty"`
}

type Leak struct {
	Ref    string `json:"ref"`
	Phrase string `json:"phrase"`
	Note   string `json:"note"`
}

type ReviewUnit struct {
	Unit     string    `json:"unit"`
	Model    string    `json:"model"`
	Summary  string    `json:"summary"`
	Patches  []Patch   `json:"patches"`
	Verdicts []Verdict `json:"verdicts"`
	Leakage  []Leak    `json:"leakage,omitempty"`
}

// Variant, when set, puts the translation and review files of an
// experiment under stage/alt/<variant>/, apart from the edition's own.
var Variant string

func path(cfg *work.Config, stage, unit string) string {
	if Variant != "" && (stage == "translate" || stage == "review" || stage == "final") {
		return cfg.Path("stage", "alt", Variant, stage, unit+".json")
	}
	return cfg.Path("stage", stage, unit+".json")
}

// ReviewPackPath is the review pack of a unit, under the variant if one is set.
func ReviewPackPath(cfg *work.Config, unit string) string {
	if Variant != "" {
		return cfg.Path("stage", "alt", Variant, "packs", "review", unit+".md")
	}
	return cfg.Path("stage", "packs", "review", unit+".md")
}

func MorphPath(cfg *work.Config, unit string) string       { return path(cfg, "morph", unit) }
func TranslationPath(cfg *work.Config, unit string) string { return path(cfg, "translate", unit) }
func ReviewPath(cfg *work.Config, unit string) string      { return path(cfg, "review", unit) }
func FinalPath(cfg *work.Config, unit string) string       { return path(cfg, "final", unit) }

func exists(p string) bool { _, err := os.Stat(p); return err == nil }

// LoadMorph reads every morphology file and applies the morph patches of
// any review. It returns the analyses and their provenance by token id.
func LoadMorph(cfg *work.Config) (map[string]corpus.Analysis, map[string]string, error) {
	out := map[string]corpus.Analysis{}
	src := map[string]string{}
	files, _ := filepath.Glob(cfg.Path("stage", "morph", "*.json"))
	sort.Strings(files)
	for _, f := range files {
		var mu MorphUnit
		if err := corpus.ReadJSON(f, &mu); err != nil {
			return nil, nil, err
		}
		for _, a := range mu.Tokens {
			out[a.ID] = a
			src[a.ID] = "model:" + mu.Model
		}
		rp := ReviewPath(cfg, mu.Unit)
		if !exists(rp) {
			continue
		}
		var ru ReviewUnit
		if err := corpus.ReadJSON(rp, &ru); err != nil {
			return nil, nil, err
		}
		for _, a := range mu.Tokens {
			src[a.ID] += ", reviewed:" + ru.Model
		}
		for _, p := range ru.Patches {
			id, ok := strings.CutPrefix(p.Path, "morph:")
			if !ok {
				continue
			}
			var a corpus.Analysis
			if err := json.Unmarshal(p.Value, &a); err != nil {
				return nil, nil, fmt.Errorf("%s: patch %s: %w", rp, p.Path, err)
			}
			a.ID = id
			out[id] = a
			src[id] = "model:" + mu.Model + ", corrected:" + ru.Model
		}
	}
	return out, src, nil
}
