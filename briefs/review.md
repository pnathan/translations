# Brief: review

You review one unit's translation, and for the Mishnah its morphology, against the Hebrew, the lexicon and the normative commentators. You do not rewrite the file under review. You write patches; the pipeline checks them and applies them.

## Input

- The review pack `works/<work>/stage/packs/review/<unit>.md`. It holds every segment's Hebrew, tokens, morphology, lexicon, variants and the full commentary, each followed by the translation under review.
- The files under review: `works/<work>/stage/translate/<unit>.json`, and for the Mishnah `works/<work>/stage/morph/<unit>.json`.
- `works/<work>/decisions.md` and `briefs/translate.md`: the translator's rules. Hold the translation to them.

## What to check

1. Fidelity. Every Hebrew word is accounted for. Person, number, gender, tense and aspect, and the binyan are carried over correctly. Nothing is added without a mark.
2. The glosses and senses. Each gloss fits its context, and each `sense` names the right sense of the entry. Check every sense: a translator may have filled them mechanically (for example, the first sense of each entry). Where the entry is the wrong homograph, patch `token:<id>.lex` to the right one if the pack lists it, else remove the sense and say so in the token's note.
3. The normative commentators. For each normative anchor of the unit, state whether the translation is consistent with the comment or contradicts it. Where the normative commentators (Genesis from chapter 18: Ibn Ezra; Avot: Bartenura and Rambam) agree on the plain sense against the English, the English is wrong, except as the work's decisions provide.
4. The translator's rulings on the commentary. Each `summary` must say what the comment says. A wrong summary is an error.
5. Morphology (Mishnah only). Each analysis is correct: segmentation, lemma, binyan, form, person, gender, number. Correct an analysis with a `morph:` patch.
6. Correctness only. Agreement with a published translation is not a fault. Patch only what is wrong, or what a better English would carry more exactly.
7. The learner notes are true, and enough for a student with one year of Hebrew.
8. The ancient setting (Genesis). The English does not read a later theology or a later world into a word, and the `context` notes are true: established facts stated as such, disputes named as disputes, nothing invented, no modern work cited by title, page or date. Remove a note you cannot stand behind; add one a reader needs.

## Output

Write `works/<work>/stage/review/<unit>.json`:

```json
{
  "unit": "bereshit-01",
  "model": "opus",
  "summary": "two to five sentences: the state of the unit and its main faults",
  "patches": [
    {"ref": "1:2", "path": "english", "value": "the whole new English of the segment", "severity": "error", "reason": "why"},
    {"ref": "1:2", "path": "phrases", "value": [{"tokens": ["1:2.1"], "english": "...", "literal": "..."}], "severity": "error", "reason": "the new English changed the fragments"},
    {"ref": "1:2", "path": "token:1:2.3.gloss", "value": "void", "severity": "style", "reason": "..."},
    {"ref": "1:2", "path": "commentary:rashi:1:2:1.summary", "value": "...", "severity": "error", "reason": "..."},
    {"ref": "1:2", "path": "commentary:ibn-ezra:1:2:3", "value": {"relation": "records", "summary": "..."}, "severity": "question", "reason": "a missing ruling"},
    {"ref": "1:1", "path": "morph:1:1.2", "value": {"lang": "mishnaic", "root": "קבל", "morphemes": [{"form": "קִבֵּל", "lemma": "קָבַל", "pos": "verb", "parse": "piel, perfect (qatal), 3rd person, masculine, singular"}]}, "severity": "error", "reason": "..."}
  ],
  "verdicts": [
    {"anchor": "rashi:1:1:2", "verdict": "consistent", "comment": "optional"}
  ]
}
```

- Paths: `english`, `literal`, `note`, `grammar` (a whole list), `phrases` (a whole list), `apparatus` (a whole list), `context` (a whole list), `token:<id>.gloss`, `token:<id>.sense`, `token:<id>.lex`, `token:<id>.note`, `commentary:<anchor>.relation`, `commentary:<anchor>.summary`, `commentary:<anchor>` (adds or replaces an entry), and `morph:<id>` (Mishnah only).
- `value`: a JSON string for a text field; a list or an object where the path names one.
- If you change `english`, check that every phrase's `english` still occurs in it. If one does not, patch `phrases` too.
- `severity`: `error` (wrong), `question` (open, with your proposed fix), `style` (better English, same sense).
- `verdicts`: one for every anchor the pack marks **normative**, in every segment of the unit. `verdict` is one of consistent, contradicts, not-applicable. A `contradicts` needs a patch that fixes the English, or a `comment` that says why the English stands.
- Do not patch what is correct.

## Check

Run `./bin/tr check <work> <unit>`. The check applies your patches and checks the result. Correct the review until it prints `<unit>: ok`. Then stop. Do not edit any other file.
