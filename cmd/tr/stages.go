package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pnathan/translations/internal/anchor"
	"github.com/pnathan/translations/internal/corpus"
	"github.com/pnathan/translations/internal/lex"
	"github.com/pnathan/translations/internal/llm"
	"github.com/pnathan/translations/internal/work"
)

func loadEnv(cfg *work.Config) (*llm.Env, error) {
	c, err := corpus.Load(cfg.Path("stage", "corpus.json"))
	if err != nil {
		return nil, fmt.Errorf("%w (run tr build)", err)
	}
	e := &llm.Env{Corpus: c}
	if err := corpus.ReadJSON(cfg.Path("stage", "lexicon.json"), &e.Lexicon); err != nil {
		return nil, err
	}
	if e.Lexicon == nil {
		e.Lexicon = map[string]*lex.Entry{}
	}
	var as []anchor.Anchor
	if err := corpus.ReadJSON(cfg.Path("stage", "anchors.json"), &as); err != nil {
		return nil, err
	}
	e.Anchors = as
	return e, nil
}

func unitsOf(e *llm.Env, args []string) []string {
	if len(args) > 0 {
		return args
	}
	var out []string
	for _, u := range e.Corpus.Units {
		out = append(out, u.ID)
	}
	return out
}

func check(cfg *work.Config, args []string) error {
	e, err := loadEnv(cfg)
	if err != nil {
		return err
	}
	bad := 0
	for _, id := range unitsOf(e, args) {
		var errs []string
		var mu *llm.MorphUnit
		if p := llm.MorphPath(cfg, id); exists(p) {
			mu = &llm.MorphUnit{}
			if err := corpus.ReadJSON(p, mu); err != nil {
				errs = append(errs, err.Error())
			} else {
				for _, x := range e.CheckMorph(mu) {
					errs = append(errs, "morph: "+x)
				}
			}
		}
		var tu *llm.TranslationUnit
		if p := llm.TranslationPath(cfg, id); exists(p) {
			tu = &llm.TranslationUnit{}
			if err := corpus.ReadJSON(p, tu); err != nil {
				errs = append(errs, err.Error())
				tu = nil
			} else {
				for _, x := range e.CheckTranslation(tu) {
					errs = append(errs, "translate: "+x)
				}
			}
		}
		if p := llm.ReviewPath(cfg, id); exists(p) {
			var ru llm.ReviewUnit
			switch {
			case corpus.ReadJSON(p, &ru) != nil:
				errs = append(errs, "review: cannot read "+p)
			case tu == nil:
				errs = append(errs, "review: no translation to review")
			default:
				for _, x := range e.CheckReview(&ru, tu, mu) {
					errs = append(errs, "review: "+x)
				}
			}
		}
		if len(errs) > 0 {
			bad++
			fmt.Printf("%s: %d problems\n", id, len(errs))
			for _, x := range errs {
				fmt.Println("  " + x)
			}
		} else if len(args) > 0 {
			fmt.Printf("%s: ok\n", id)
		}
	}
	if bad > 0 {
		return fmt.Errorf("%d units fail", bad)
	}
	return nil
}

func exists(p string) bool { _, err := os.Stat(p); return err == nil }

func apply(cfg *work.Config) error {
	e, err := loadEnv(cfg)
	if err != nil {
		return err
	}
	os.MkdirAll(cfg.Path("stage", "final"), 0o755)
	n, missing := 0, 0
	for _, u := range e.Corpus.Units {
		var tu llm.TranslationUnit
		var ru llm.ReviewUnit
		if corpus.ReadJSON(llm.TranslationPath(cfg, u.ID), &tu) != nil || corpus.ReadJSON(llm.ReviewPath(cfg, u.ID), &ru) != nil {
			missing++
			continue
		}
		final, err := llm.Apply(&tu, &ru)
		if err != nil {
			return fmt.Errorf("%s: %w", u.ID, err)
		}
		if errs := e.CheckTranslation(final); len(errs) > 0 {
			return fmt.Errorf("%s: after review: %s", u.ID, strings.Join(errs, "; "))
		}
		if err := corpus.WriteJSON(llm.FinalPath(cfg, u.ID), final); err != nil {
			return err
		}
		n++
	}
	fmt.Fprintf(os.Stderr, "%d units final, %d without translation or review\n", n, missing)
	return nil
}

func status(cfg *work.Config) error {
	e, err := loadEnv(cfg)
	if err != nil {
		return err
	}
	mark := func(p string) string {
		if exists(p) {
			return "x"
		}
		return "."
	}
	fmt.Println("unit            refs          morph translate review final")
	for _, u := range e.Corpus.Units {
		fmt.Printf("%-15s %-13s %-5s %-9s %-6s %s\n", u.ID, u.Refs[0]+"–"+u.Refs[len(u.Refs)-1],
			mark(llm.MorphPath(cfg, u.ID)), mark(llm.TranslationPath(cfg, u.ID)), mark(llm.ReviewPath(cfg, u.ID)), mark(llm.FinalPath(cfg, u.ID)))
	}
	return nil
}

func show(cfg *work.Config, args []string) error {
	c, err := corpus.Load(cfg.Path("stage", "corpus.json"))
	if err != nil {
		return err
	}
	for _, r := range args {
		s := c.Segment(r)
		if s == nil {
			return fmt.Errorf("no segment %s", r)
		}
		fmt.Printf("%s [%s] %s\n", s.Ref, s.Unit, s.Text)
		for _, t := range s.Tokens {
			var ms []string
			for _, m := range t.Morphemes {
				ms = append(ms, m.Form+"="+m.Lemma+" ("+m.POS+" "+m.Parse+")")
			}
			fmt.Printf("  %-8s %-16s %s\n", t.ID, t.Text, strings.Join(ms, " + "))
		}
	}
	return nil
}
