// Command tr is the translation pipeline.
//
//	tr fetch  WORK            fetch witnesses, commentaries and lexica (cached)
//	tr build  WORK            corpus, lexicon, anchors and packs
//	tr check  WORK [UNIT…]    validate model outputs (morph, translation, review)
//	tr apply  WORK            merge reviews into the final edition
//	tr render WORK            write the Markdown edition and the HTML pages
//	tr status WORK            the stage of every unit
//	tr show   WORK REF…       print the tokens of segments
//
// Run from the repository root.
package main

import (
	"fmt"
	"os"

	"github.com/pnathan/translations/internal/work"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: tr fetch|build|check|apply|render|status|show WORK [ARGS]")
		os.Exit(2)
	}
	cfg, err := work.Load(".", os.Args[2])
	if err != nil {
		fail(err)
	}
	args := os.Args[3:]
	switch os.Args[1] {
	case "fetch":
		err = fetch(cfg, args)
	case "build":
		err = build(cfg)
	case "check":
		err = check(cfg, args)
	case "apply":
		err = apply(cfg)
	case "render":
		err = renderWork(cfg)
	case "status":
		err = status(cfg)
	case "show":
		err = show(cfg, args)
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "tr:", err)
	os.Exit(1)
}
