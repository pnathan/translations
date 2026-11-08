package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pnathan/translations/internal/concord"
	"github.com/pnathan/translations/internal/work"
)

// concordance joins the final translations of WORK and the further works
// named in args into docs/lexicon.
func concordance(cfg *work.Config, args []string) error {
	cfgs := []*work.Config{cfg}
	for _, a := range args {
		c, err := work.Load(".", a)
		if err != nil {
			return err
		}
		cfgs = append(cfgs, c)
	}
	var srcs []*concord.Source
	var works []concord.Work
	for _, c := range cfgs {
		s, err := concord.Load(c)
		if err != nil {
			return fmt.Errorf("%s: %w", c.Slug, err)
		}
		srcs = append(srcs, s)
		works = append(works, concord.Work{Slug: c.Slug, Title: c.Title})
	}
	cc := concord.Build(srcs)
	dir := filepath.Join("docs", "lexicon")
	if err := cc.Write(dir, works, srcs); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "wrote %s: %d groups, %d words\n", dir, len(cc.Roots), len(cc.ByToken))
	return nil
}
