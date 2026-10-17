# translations

Translations of Jewish sacred literature from the Hebrew, made by a
repeatable pipeline. Every word carries its morphology, lemma, root and the
lexicon's range of senses, and the classical commentators are anchored to
the words they explain. The editions serve readers and learners of Hebrew.

| Work | Text | Morphology | Commentators |
|------|------|------------|--------------|
| Genesis 1:1–17:27 (Bereshit, Noach, Lekh Lekha) | Miqra according to the Masorah (Sefaria), collated with the WLC | OSHB | Rashi, Ibn Ezra (normative); Ramban, Radak, Sforno |
| Pirkei Avot 1–6 | Torat Emet (Sefaria), collated with Vilna 1913 | model, reviewed | Bartenura, Rambam (normative); Rabbeinu Yonah |

## Pipeline

```
fetch ─► build ─► [morph] ─► translate ─► review ─► apply ─► render
Sefaria   corpus    model      model        model     merge     edition.md
OSHB      anchors   (Mishnah)  (Haiku)      (Opus)    patches   docs/<work>/*.html
BDB       lexicon
          packs
```

1. `tr fetch WORK` fetches the base text, the witnesses, the commentaries and
   the lexica. Every response is stored unchanged under `works/WORK/raw/` with
   its URL, date and SHA-256 in `manifest.json`.
2. `tr build WORK` writes `stage/corpus.json` (tokens with fixed ids, OSHB
   morphology, BDB lexicon keys, accents, variants between witnesses),
   `stage/anchors.json` (each comment tied to the tokens its opening quotation
   names) and `stage/lexicon.json`. It then renders the Markdown packs the
   model stages read, under `stage/packs/`.
3. The model stages read a pack and a brief (`briefs/`) and write JSON:
   `stage/morph/` (Mishnah only), `stage/translate/`, `stage/review/`. The
   review writes patches; it does not edit the translation.
4. `tr check WORK [UNIT…]` validates every model file against the corpus:
   every token glossed, the phrases cover the tokens in order and quote the
   English, every sense exists in its lexicon entry, every normative comment
   ruled on, every review patch applies.
5. `tr apply WORK` merges the reviews into `stage/final/`.
6. `tr render WORK` writes `works/WORK/edition.md` and `docs/WORK/`.

`tr report WORK` writes `works/WORK/review.md`: patch counts per unit and
every open question the review put to the editor. `tr status WORK` lists
the stage of every unit; `tr show WORK REF…` prints
the tokens of a segment. Build with `go build -o bin/tr ./cmd/tr` and run
from the repository root. Go standard library only.

## Rules

- Translate from the fetched Hebrew only. Correctness is the test;
  agreement with a published translation is not a fault.
- Morphology and lexicon come from datasets wherever one exists (OSHB, BDB).
  The Mishnah has none: its morphology is a model stage, reviewed, and
  labelled as such on every word.
- Witnesses are collated, not collapsed: variants between MAM and the WLC,
  ketiv and qere, and Torat Emet against Vilna are recorded per token.
- No Unicode normalization of stored Hebrew (NFC reorders Hebrew marks).
  Matching uses consonant skeletons.

## Models

Model test on Genesis 1:1–5, both versions reviewed by Opus under one brief:
Haiku drew 40 error patches (223k tokens, 6 min), Sonnet at low effort 19
(145k tokens, 2 min). Sonnet translates; Opus reviews; Haiku tagged the
Avot morphology. The Haiku translations are kept in `stage/alt/haiku/`.

## Known limitations

- BDB senses without a `<def>` element in the OpenScriptures XML are not
  extracted, so some entries lack senses (מֵאָה "hundred" lists only
  "hundredth"). Reviewers removed such senses and said so in the token
  note. Re-extraction would renumber `#k` paths; it waits for a new run.
- Avot: a model lemma is attached to the first Sefaria candidate entry,
  which is sometimes the wrong homograph; the review repoints `lex` where
  the pack lists the right one.
- Agents share one scratch directory unless told otherwise; prompts now
  give each agent its own.

## Licences

The pipeline code and the translations: CC0 (see `LICENSE`), except where a
source licence requires otherwise. The published pages and `edition.md`
reproduce the MAM text (CC BY-SA) and OSHB morphology (CC BY 4.0, credit the
Open Scriptures Hebrew Bible Project); they are distributed under CC BY-SA
4.0. The commentaries, Jastrow and BDB are in the public domain or CC0.
