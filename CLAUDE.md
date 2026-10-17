# translations

Translations of Jewish sacred literature from the Hebrew, with word-level
analysis and anchored commentary. See `README.md` for the pipeline.

## Layout

```
cmd/tr/                 the pipeline CLI (fetch, build, check, apply, render, status, show)
internal/hebrew/        skeletons, word splitting, HTML removal, cantillation names
internal/oshb/          OSHB OSIS reader, morphology-code decoder
internal/lex/           OpenScriptures BDB index; Sefaria (Jastrow, BDB) lookups
internal/sefaria/       cached fetches with manifest
internal/corpus/        token stream, witness alignment, units
internal/anchor/        commentary-to-span anchoring by dibbur ha-matḥil
internal/pack/          Markdown packs for the model stages
internal/llm/           model-output schemas, checks, review patches
internal/render/        edition.md and the HTML pages
briefs/                 instructions for the model stages (morph, translate, review)
works/<work>/work.json  sources, commentators, sections, unit size
works/<work>/decisions.md  fixed editorial decisions for every unit
works/<work>/raw/       fetched sources, unchanged, with manifest.json
works/<work>/stage/     corpus, anchors, lexicon, packs, model outputs, final
sources/openscriptures/ shared lexicon files
docs/                   GitHub Pages root (deployed from main)
```

## Model stages

- Run from the repository root after `go build -o bin/tr ./cmd/tr`.
- Translation and morphology: Haiku subagents. Review: Opus subagents. One
  agent per unit; each runs `./bin/tr check WORK UNIT` until it passes.
- A model writes only its own stage file. It never edits the corpus, the
  packs or another stage's file. The review writes patches, never a new
  translation.
- After a stage: `./bin/tr build WORK` (regenerates packs; for the Mishnah it
  merges morphology, then `./bin/tr fetch WORK` looks up the new lemmas and a
  second `build` attaches them), then `./bin/tr apply WORK` and
  `./bin/tr render WORK`.

## Rules (do not violate)

1. Translate only from fetched witnesses. Never consult or reproduce a
   published translation.
2. Morphology and lexicon come from datasets where one exists. Model
   morphology is labelled on every word with its model and reviewer.
3. Record every variant between witnesses; never collapse them silently.
4. Normative commentators (marked in `work.json`) decide the plain sense;
   the check requires a ruling on every normative comment, and the review a
   verdict on each.
5. No NFC normalization of stored Hebrew.

## Adding a work

1. `works/<slug>/work.json` (copy an existing one) and `decisions.md`.
2. `tr fetch`, `tr build`; run the model stages unit by unit; `tr apply`;
   `tr render`.
3. Add a card to `docs/index.html`.
