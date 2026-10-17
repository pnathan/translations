package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/render"
	"github.com/pnathan/translations/internal/work"
)

func renderWork(cfg *work.Config) error {
	e, err := loadEnv(cfg)
	if err != nil {
		return err
	}
	ed := &render.Edition{Cfg: cfg, Corpus: e.Corpus, Lexicon: e.Lexicon, Anchors: e.Anchors,
		Final: map[string]*llm.SegmentTr{}, Models: map[string]string{}, Reviews: map[string]*llm.ReviewUnit{}}
	for _, u := range e.Corpus.Units {
		var tu llm.TranslationUnit
		if corpus.ReadJSON(llm.FinalPath(cfg, u.ID), &tu) != nil {
			ed.Missing = append(ed.Missing, u.ID)
			continue
		}
		ed.Models[u.ID] = tu.Model
		for i := range tu.Segments {
			ed.Final[tu.Segments[i].Ref] = &tu.Segments[i]
		}
	}
	if err := ed.Markdown(cfg.Path("edition.md")); err != nil {
		return err
	}
	dir := filepath.Join("docs", cfg.Slug)
	if err := ed.HTML(dir); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s (%d units missing)\n", cfg.Path("edition.md"), dir, len(ed.Missing))
	return nil
}
