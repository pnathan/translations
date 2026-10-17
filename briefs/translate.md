# Brief: translation

You translate one unit of Hebrew text into English. The result is the text of an edition for readers and for learners of Hebrew. Each word gets an analysis.

## Input

- The pack: `works/<work>/stage/packs/translate/<unit>.md`. It holds the Hebrew of each segment (verse or mishnah) and one table row per token: id, word, morphemes, lemmas with lexicon keys, parse, root and accents. It also holds variant readings, the commentary anchored to the segment, and the lexicon entries the unit cites.
- The work's decisions: `works/<work>/decisions.md`. Follow them in every unit.

## Rules

1. Translate from the Hebrew in the pack. Do not consult a published translation. Do not reproduce one from memory: if your English for a phrase is a famous rendering (KJV, JPS, RSV, NJPS, Artscroll, Koren, Soncino) and the Hebrew does not force it, write your own.
2. The morphology and lexicon in the pack are data from fetched sources. Do not contradict them. If you think one is wrong, keep the data and say so in `note`.
3. The commentators marked **normative** decide the plain sense. For Genesis these are Rashi and Ibn Ezra; for Avot, Bartenura and Rambam. Where they agree on a reading that bears on the English, the English follows it. Where they differ, choose one reading and record the other with `records`.
4. Do not leave out the Hebrew's difficulties. If a word is ambiguous, keep the ambiguity in the English, or choose one sense and give the other in `note` or `apparatus`.
5. Write English only in `english`: no Hebrew letters. Hebrew is allowed in the notes.

## Output

Write `works/<work>/stage/translate/<unit>.json`:

```json
{
  "unit": "bereshit-01",
  "model": "haiku",
  "segments": [
    {
      "ref": "1:1",
      "english": "the reading translation of the whole segment",
      "literal": "a word-for-word rendering, in Hebrew word order, hyphens joining one word's parts: in-beginning-of created God ...",
      "tokens": [
        {"id": "1:1.1", "gloss": "in the beginning of", "sense": "1", "note": "construct state: 'the beginning of'"},
        {"id": "1:1.2", "gloss": "created", "sense": "Qal 1"}
      ],
      "phrases": [
        {"tokens": ["1:1.1", "1:1.2", "1:1.3"], "english": "When God began to create", "literal": "'in-beginning-of created God'; the construct reading of Rashi"}
      ],
      "grammar": [
        "בְּרֵאשִׁית is in the construct state: Rashi reads 'in the beginning of (God's) creating'.",
        "The etnaḥta on אֱלֹהִים divides the verse in two."
      ],
      "note": "the translator's reasons for any choice that is not plain from the phrases",
      "commentary": [
        {"anchor": "rashi:1:1:2", "relation": "follows", "summary": "Rashi: רֵאשִׁית is always bound to the next word, so the verse is a time clause ..."}
      ],
      "apparatus": ["variant readings or open questions, one per line"]
    }
  ]
}
```

The example values show the format only; they are not decisions.

- `segments`: one per segment of the unit, in order.
- `tokens`: every token id of the segment, once each.
  - `gloss`: the English sense of the token in this context, with its prefixes ("and the earth").
  - `sense`: the path of the matching sense in the token's lexicon entry, written exactly as the pack prints it, without the backticks: `1`, `2.a`, `Qal 1`, `Niph 2`. Give it for every noun, verb, adjective and adverb whose entry lists senses. Leave it out for prefixes and particles that have no senses.
  - `lex`: only when the pack lists several entries for one lemma (homographs). Give the key of the one you mean.
  - `note`: give it where a learner needs it: a pausal form, a wayyiqtol, a construct chain, a weak root, an idiom.
- `phrases`: they divide all the tokens of the segment, in order, with no gaps and no overlaps. `english` must occur in `english` word for word; case and punctuation are ignored. Use `-` when the Hebrew has no English counterpart (for example, the object marker אֵת). `literal` gives the literal sense of the Hebrew and the reason for the English.
- `grammar`: 2 to 6 notes for a learner on this segment: syntax, verb forms, the accents that divide the segment, word order, wordplay.
- `commentary`: one entry for every anchor the pack marks **normative** on this segment. Add other anchors where they bear on the English. `relation` is one of:
  - `follows`: the English adopts this reading.
  - `records`: the English does not adopt it; the reading is a real alternative and is kept here.
  - `declines`: the English rejects it. Give the reason in `summary`.
  - `not-applicable`: the comment does not bear on the plain sense: homily, law, a cross-reference.

  `summary` gives in English, in one to three sentences, what the commentator says about the words. Quote his Hebrew where it matters.

## Check

From the repository root, run `./bin/tr check <work> <unit>` (or `go run ./cmd/tr check <work> <unit>`). Correct the file until it prints `<unit>: ok`. Then stop. Do not edit any other file.
