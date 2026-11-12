# translations

Translations of Jewish sacred literature from the Hebrew, made by a
repeatable pipeline. Every word carries its morphology, lemma, root and the
lexicon's range of senses, and the classical commentators are anchored to
the words they explain. The editions serve readers and learners of Hebrew.

| Work | Text | Morphology | Commentators |
|------|------|------------|--------------|
| Genesis 1–50 (Bereshit to Vayechi; 18–50 in progress) | Miqra according to the Masorah (Sefaria), collated with the WLC | OSHB | 1–17: Rashi, Ibn Ezra (normative); from 18: Ibn Ezra (normative); Rashi, Ramban, Radak, Sforno, Onkelos |
| Pirkei Avot 1–6 | Torat Emet (Sefaria), collated with Vilna 1913 | model, reviewed | Bartenura, Rambam (normative); Rabbeinu Yonah |
| Milchamot ha-Shem 1–10 (Qafiḥ, 1931; sources fetched, not yet translated) | Wikisource via Sefaria, to be collated with the yahadut.org.il PDF | none (prose kind) | none |

## Pipeline

```
fetch ─► build ─► [morph] ─► translate ─► review ─► apply ─► render
Sefaria   corpus    model      model        model     merge     edition.md
OSHB      anchors   (Mishnah)  (Haiku)      (Sonnet)  patches   docs/<work>/*.html
BDB       lexicon                                                docs/lexicon (concord)
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

`tr concord WORK…` joins the final translations of the works named into
the concordance, `docs/lexicon/`: every root, its lemmas, each lemma's
senses, and under each sense every word whose reviewed translation chose
it, with links into the text. Roots join across works on their consonants
with the shin and sin dots kept; an undotted root (the Mishnah's model
roots) takes the dot of its lemma's pointed headword, else the one dotted
root with its consonants, else stays apart, marked undetermined. Nothing in
it is written by a model. The section pages link each word to its root, and
their 3D view swings the page back to show, behind it, every verse of the
corpus that uses the word's root, the root's words highlighted.

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
- The English is the plain translation of the words. A commentator decides
  between readings of the words; his explanation of what happened or why
  goes in the notes, never in the English line.
- Genesis from chapter 18: Ibn Ezra is the one normative commentator;
  Rashi is consulted, his midrash recorded. The ancient Near Eastern
  setting is rendered as the evidence shows it, and the `context` notes
  name disputes as disputes; they cite no modern work, since the pipeline
  fetches none. See `works/genesis/decisions.md`.

## Models

Model test on Genesis 1:1–5, both versions reviewed by Opus under one brief:
Haiku drew 40 error patches (223k tokens, 6 min), Sonnet at low effort 19
(145k tokens, 2 min). Sonnet translates; Opus reviews; Haiku tagged the
Avot morphology. The Haiku translations are kept in `stage/alt/haiku/`.

Genesis 18–50 is translated by Haiku (medium effort) and reviewed by
Sonnet, for cost. On Genesis 18:1–8 one Haiku translation drew 26 error
patches from Opus and 37 from Sonnet; each review found faults the other
missed (8 shared), and the unit's review merges both
(`stage/alt/opus-pilot/`, `stage/alt/sonnet-review/`). A Haiku review of
the same unit found 5, two of them wrong (`stage/alt/haiku-review/`).
Units are reviewed one by one; a consistency pass per parasha checks
names, recurring terms and quotations that cross units.

## Known limitations

- The abridged OpenScriptures BDB gives many senses no `<def>` (שָׁמַע Qal:
  "subj. pers. hum."). They are now extracted, their gloss the nearest
  definition above and their own text; a derived stem names the
  definition it derives from ("subj. God (Hiph of: die)"). All 6,879
  senses cited before the change resolve to the same gloss. Genesis 1–17
  was translated before it; some of its tokens lack a sense that now
  exists.
- The packs list every sense of an entry. They once cut entries at 24
  senses, hiding the derived stems of eight long verbs (שׁוּב, הָלַךְ, עָלָה);
  no final unit carried a derived-stem form of them without a sense.
- Sefaria has no Ramban on Genesis 50.
- OSHB links some forms to the wrong lemma (רָבָּה 18:20, a verb, to the
  adjective רַב); the review removes the sense and says so.
- Avot: a model lemma is attached to the first Sefaria candidate entry,
  which is sometimes the wrong homograph; the review repoints `lex` where
  the pack lists the right one.
- Avot: a reviewer chooses senses against the entry its pack shows, but
  its own `morph:` patches can change the lemma and with it the default
  entry. The build keeps every entry a translation or review names; seven
  tokens carry an editor patch (reason "editor: ...") that pins the entry
  the sense refers to. BDB Augmented Strong has 150 headwords shared by
  homographs, and the key `bdbs:<headword>` holds only one of them.
- Agents share one scratch directory unless told otherwise; prompts now
  give each agent its own.

## Licences

The pipeline code and the translations: CC0 (see `LICENSE`), except where a
source licence requires otherwise. The published pages and `edition.md`
reproduce the MAM text (CC BY-SA) and OSHB morphology (CC BY 4.0, credit the
Open Scriptures Hebrew Bible Project); they are distributed under CC BY-SA
4.0. The commentaries, Jastrow and BDB are in the public domain or CC0.
