# Brief: morphology (Mishnah)

The Mishnah has no published morphological dataset. You give the analysis of each token. A reviewer checks it later, and the edition labels it as a model analysis.

## Input

The pack `works/<work>/stage/packs/morph/<unit>.md`. It holds every segment's Hebrew and its tokens, with their ids.

## Output

Write `works/<work>/stage/morph/<unit>.json`:

```json
{
  "unit": "perek-1-01",
  "model": "haiku",
  "tokens": [
    {"id": "1:1.1", "lang": "mishnaic", "root": "", "morphemes": [
      {"form": "משֶׁה", "lemma": "מֹשֶׁה", "pos": "noun", "parse": "proper name, person, masculine"}]},
    {"id": "1:1.2", "lang": "mishnaic", "root": "קבל", "morphemes": [
      {"form": "קִבֵּל", "lemma": "קָבַל", "pos": "verb", "parse": "piel, perfect (qatal), 3rd person, masculine, singular"}]},
    {"id": "1:1.5", "lang": "mishnaic", "root": "מסר", "morphemes": [
      {"form": "וּ", "lemma": "ו: and", "pos": "conjunction", "parse": "conjunction"},
      {"form": "מְסָרָ", "lemma": "מָסַר", "pos": "verb", "parse": "qal, perfect (qatal), 3rd person, masculine, singular"},
      {"form": "הּ", "lemma": "", "pos": "suffix", "parse": "pronominal, 3rd person, feminine, singular"}]}
  ]
}
```

## Rules

- One entry for every token of the unit.
- `morphemes`: split off the prefixes (ו ה ב כ ל מ ש ד) and the pronominal suffixes. Copy the forms from the token as it is printed, with its points. The consonants of the forms, joined, must spell the token.
- `lemma`:
  - Verbs: the root in the pointed qal perfect, third person masculine singular (קָבַל, אָמַר), whatever the binyan of the word.
  - Nouns and adjectives: the pointed singular absolute (תּוֹרָה, גָּדוֹל).
  - Names: the name.
  - Prefixes: their gloss in the form `ו: and`, `ה: the`, `ב: in, with, by`, `כ: like, as`, `ל: to, for`, `מ: from, than`, `ש: who, which, that`.
  - Suffixes: an empty lemma.
- `pos`: one of noun, verb, adjective, adverb, pronoun, preposition, conjunction, particle, suffix, article, numeral, interjection.
- `parse`: use the terms of these models:
  - verb: `<binyan>, <form>, <person>, <gender>, <number>`, with binyan from qal, nifal, piel, pual, hifil, hofal, hitpael, nitpael, and form from perfect (qatal), imperfect (yiqtol), imperative, participle active, participle passive, infinitive construct.
  - participle: `piel, participle active, masculine, plural, absolute`.
  - noun: `common, feminine, singular, construct`, or `proper name, person, masculine`.
  - suffix: `pronominal, 3rd person, feminine, singular`.
- `lang`: `mishnaic`. Use `arc` for an Aramaic word (Hillel's sayings in 1:13 and 2:6; Ben Bag Bag and Ben He He in chapter 5) and `hbo` for a word inside a quotation from the Bible.
- `root`: the consonants of the root, without points; leave it empty for names and particles.

## Check

Run `./bin/tr check <work> <unit>`. Correct the file until it prints `<unit>: ok`. Then stop. Do not edit any other file.
