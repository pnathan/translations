# Runbook: one parasha per session

A session takes one parasha of a work (its issue names it) from fetched text
to final, reviewed units, on its own branch, and opens one pull request. It
is the coordinator: it does not translate or review; subagents do, one per
unit.

## Before you start

1. Read `CLAUDE.md`, `README.md`, this file, `briefs/translate.md`,
   `briefs/review.md` and the work's `decisions.md`.
2. Branch from `main`: `git checkout -b <work>-<parasha>`.
3. `go build -o bin/tr ./cmd/tr` and `./bin/tr status <work> | grep <parasha>`:
   the units, and which are already done.

## What a parasha branch may change

Only, per unit of the parasha:
`works/<work>/stage/translate/<unit>.json`,
`works/<work>/stage/review/<unit>.json`,
`works/<work>/stage/packs/review/<unit>.md`.

Never on the branch: code, briefs, `decisions.md`, `work.json`, the corpus,
lexicon or anchors, `stage/final/`, `edition.md`, `docs/`. `tr apply`,
`tr render` and `tr concord` run on `main` after the merge. A problem in the
pipeline, the lexicon or the decisions goes in the issue as a comment, for
`main`; the unit in question waits or carries a note.

`tr build` may be run on the branch (it writes the review packs a review
needs); commit only the review packs of the parasha's units.

## Models and budget

- Translation: Haiku, effort medium. Review and consistency: Sonnet. Opus
  only if the owner asks.
- At most six agents at a time. The usage limits are per account: more
  agents only reach them sooner. When an agent fails on a limit, its unit
  waits; relaunch after the reset with "if the file exists, complete it".
- Context: each agent reads each input once, the decisions only from the
  line where "Style" begins, writes its JSON from a script and never prints
  it, and pipes command output through `head`. Give it only what is specific
  to its unit beyond the brief.

## The unit cycle

1. Translate (Haiku) with the translation prompt below. Commit
   `translate/<unit>.json` when `tr check` passes.
2. `tr build`, then review (Sonnet) with the review prompt. Before you
   commit, read the unit's English after the patches (apply the patches in
   a scratch copy, or read the review's `english` patches) and check:
   the English is the plain translation (no commentator's gloss in it);
   quotation marks open and close once per speech, across verses; speech is
   introduced with a comma. Fix what you find with an appended patch whose
   reason starts "editor:".
3. Commit and push after every unit that passes. Commit work in progress
   too, as "wip", when a limit is near.

## Consistency pass

When every unit of the parasha is reviewed: extract the final English of
every verse in order (a TSV of ref, unit, English, made from the review
patches applied in a scratch copy) and a count of capitalized names. Give
these, the decisions and the list of stiff passages you noted, to one Sonnet
agent with the consistency prompt below. Its patches are appended to the
units' reviews with reasons starting "consistency:".

Then verify by script: every speech opens and closes once across the
parasha; no colon introduces speech; no "behold".

## Pull request

Title: `<work>: <Parasha> (<from>–<to>)`. Body: the units, patches by
severity, the editor and consistency patches, and every open question a
review raised. It closes the parasha's issue. After the merge, on `main`:
`tr build`, `tr apply`, `tr render`, `tr concord <work> …`, commit.

## Prompts

### Translation (Haiku, effort medium)

```
Translation stage, <work> unit <unit> (<refs>). Work in <repo>.
Context budget: read each input once and keep tool output short.
1. `cat briefs/translate.md` (rules and format; follow them exactly).
2. `sed -n '/^## Style/,$p' works/<work>/decisions.md` (the decisions).
3. `cat works/<work>/stage/packs/translate/<unit>.md` once; grep for a second look.
Write works/<work>/stage/translate/<unit>.json ("model": "haiku") from a
Python script in <scratch>/agents/<unit>/ (json.dump, ensure_ascii=False).
Never print the JSON. If the file exists, check and complete it.
Rules the check does not catch: every sense a printed path that fits, given
for every noun, verb, adjective and adverb whose entry lists senses, binyan
matching the parse; person, number, gender, binyan and verb form exact; the
literal line word by word; the English supplies no subject, object, verb or
explanation the Hebrew lacks (a copula or article English needs is
grammar); summaries say what the comment says; speech introduced with a
comma, opening in its first verse and closing in its last, inner speech in
single quotes; never "behold"; context notes only what is sure, disputes
named, no modern work cited, no verse cited for a wording not checked.
<unit-specific points: speeches crossing unit boundaries, disputed words>
Run `./bin/tr check <work> <unit> | head -30` until ok; fix only reported
fields. Write only that file; never go build, tr build/apply/render, git;
no tool but reading, writing and the shell. Final message: the check line
and the hardest choice in one sentence.
```

### Review (Sonnet)

```
Review stage, <work> unit <unit> (<refs>). Work in <repo>.
Context budget as above; write the review from a script in
<scratch>/agents/review-<unit>/; never print it.
1. `cat briefs/review.md`. 2. The decisions as above.
3. `cat works/<work>/stage/packs/review/<unit>.md` once.
Rules: the normative commentator(s) decide between readings of the words;
the English is the plain translation; a commentator's explanation goes in
the notes; English grammar is not addition. Check segment by segment: every
gloss and sense against its entry; the literal line word by word; every
summary of every commentator against the Hebrew; every note for truth;
context notes (true, disputes named, nothing cited that cannot be checked);
quotation marks across verses and units. Patch only what is wrong or what
better English carries more exactly. <unit-specific points>
Write works/<work>/stage/review/<unit>.json ("model": "sonnet", a verdict
per normative anchor). Run `./bin/tr check <work> <unit> | head -30` until
ok. Same restrictions. Final message: check line; patches by severity; the
most serious fault.
```

### Consistency (Sonnet)

```
Consistency pass over <Parasha> (<refs>, units <first> … <last>).
Inputs only: <tsv> (every verse's final English), <names> (capitalized
words with counts), the decisions, <notes> (stiff passages found). To see one
segment: python one-liner on the scratch final file; the Hebrew:
`./bin/tr show <work> REF | head -3`.
Check across the parasha: quotation spans across verses and units; commas
before speech; one English form per name and recurring term (verify the
Hebrew before changing one); the noted passages; plain English.
Fix by appending to works/<work>/stage/review/<unit>.json patches with path
"english" (the whole verse) and, when a phrase changes, "phrases" (the whole
list), reason "consistency: …". Run `./bin/tr check <work> <unit>` after
each unit. Never edit translate/ or final/. Final message: one line per
change, then the units changed with their check lines.
```
