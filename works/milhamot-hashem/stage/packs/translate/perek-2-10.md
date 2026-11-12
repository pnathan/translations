# Translation pack: milhamot-hashem perek-2-10 (Milchamot ha-Shem 2:19–2:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-10`.

## Rulings for this unit

Extracted from `works/milhamot-hashem/decisions.md` for the words of this unit; the rulings below are the whole of what you need. Read the full file only if a question is not answered here.

# Milchamot ha-Shem: decisions for every unit

> Owner answered Q4 and Q7 on 2026-11-10; Q1-3, Q5-6 stand at the defaults (see "Owner's answers and defaults" at the end).

Editorial decisions set before translation. A reviewer may question one in a
`question` patch; the project owner decides. Where a ruling was amended, only
the final rule is given. "(note)" means a token `note`; "first in a unit"
means at the first occurrence in the unit.

General rule for technical terms: the English names the thing the words name
and goes no further. Kabbalistic technical names are transliterated, because
every English equivalent takes a side: the Kabbalists' defence
("configuration", "the Infinite", "Divine Presence") or Kafich's charge
("god", "idol"). He makes his charge in his own words ("אליל", "אלוהות", "אל
זר"), and the English translates those where he uses them.

Transliteration: the project's plain scheme, as in Avot and Genesis (Shetach,
Chayei Sarah, Lekh Lekha): no diacritics; ח = ch, כ = kh, צ = tz, ק = k
("Machzor Vitry"). The author: Yichya Kafich (owner, 2026-11-10; "Yihya
Kafich" is named once in the front matter). Book titles with a conventional
English spelling keep it (Kisse Eliyahu, Sefer ha-Kavanot, Machberet
ha-Kodesh); the scheme applies where the project has no fixed form.

## Output format for this work (prose)

Read this whole file before translating. A segment is a paragraph of prose.
The pack has no morphemes, lemmas, lexicon or commentary, so:

- `tokens`: one entry per word: `id`, `gloss`, and a `note` where a learner
  needs one. No `sense`, no `lex`.
- `commentary`: `[]`. `context`: only for a realia a reader needs (a Yemenite
  custom, a book, a person); say what is established; cite no modern work.
  It may carry a short note on a story the author cites (the Pardes story),
  without a reference unless the author gives one.
- `grammar`: 2 to 6 notes per paragraph (in a long one, on the hardest
  sentences). `literal`: word for word; on a long paragraph the glosses in
  Hebrew order, hyphens joining the parts of one word; no reasoning (that goes
  in `note`).
- `phrases` divide every token of the paragraph, in order; in a long
  paragraph a phrase is a clause, not a word; each phrase's `english` must
  occur in the paragraph's `english`.
- Quotations follow the Quotations section; a Bible verse is translated from
  the Hebrew as the author quotes it, the reference in `apparatus`.
- Every Hebrew abbreviation is expanded in `gloss` (the reading line carries
  the full word).

## Names and titles

- **Authorities**: acronyms and book-names used as names stay, transliterated;
  first in a unit a note gives the expansion where certain (else says so) and
  who is meant; identify no one further. ר' + name → "Rabbi". An initial a
  surname fixes is expanded, else kept (a lone כ' → "K."), noted. Surnames as
  printed, one spelling across the work.
- **Honorifics**: translate every one, opponents' included; add none.
- **Book titles**: transliterated, the meaning in a note first in a unit;
  abbreviations written out; every word kept. Bible books in English;
  tractates transliterated. A misspelt title as printed, the likelier in the
  note and `apparatus`.
- ה׳: the Tetragrammaton → "the LORD"; cited as a name ("זעיר אנפין דאיקרי
  ה'") → "who is called 'the LORD'"; before a personal name → "Mr." ("מאת ה'
  אליעזר בן יהודה" → "from Mr. Eliezer Ben-Yehuda"); a numeral or page letter
  → the number. The pack splits it into ה plus ': read it in context.
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- עבודה **אחרת** → "other worship" (note first in a unit).
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ג"כ** → "also".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
## Quotations

- **Scope.** All quoted text (Zohar, Bavli, Yerushalmi, Targum, Biblical
  Aramaic, Midrash, medieval, Judeo-Arabic) from the author's words; no Hebrew,
  Aramaic or Arabic in `english`.
- **Marking.** A source's words introduced by a formula (וז"ל, כתב, אמרו,
  גרסינן, שנאמר, כדכתיב, שנינו, הכי קאמר, אמר ר' X) or closed by עכ"ל / ע"כ:
  double quotation marks; a paraphrase or report ("כתב ש…"): none; no
  formula: marks if the source is certain, else none and a note. Nesting:
  double, single, double. A reply by ויאמר inside a quotation: double marks;
  his report of a sage's question: unmarked. Over several paragraphs
  (2:18-2:25): each opens with a mark, only the last closes. Plain text (no
  italics, no labels); the language in `grammar` ("Zohar Aramaic",
  "Judeo-Arabic"). Formulas are translated, outside the marks; glosses inside
  a quotation stay in parentheses where they stand, translated.
- **Exegesis on a Hebrew word** (את; אם אין → "Ayin"): the word stays
  transliterated ("not 'to the LORD your God' but 'et the LORD your God'"),
  explained in `grammar`.
- **Cut or changed quotation.** Only what is on the page: supply nothing וכו'
  omits, complete no verse (no "before Me"), restore no source; a verse's
  fixed wording covers only the words quoted. His comma joining two verses
  stays; an ellipsis only where he marks one (וכו', וגו'). His wording against
  the source: follow him, the source in `apparatus`.
- **Bible verses**, quoted or alluded to (*melitzah*, unmarked): from the
  Hebrew as it stands, plain modern English, not King James; each in
  `apparatus`. An echoed verse (אל דבר ה', בטרם תחיל) governs the sense of its
  words; a poetic phrase
  translated, a note naming the referent; a plural stays plural, no subject
  added; ketiv translated, the qere noted.
- **Idioms**: by sense, the literal in `literal`; a saying he glosses
  himself: both translated, unmerged.
- **References** as he gives them, English form, Arabic numerals; add,
  complete, convert or correct nothing; never the 2015 PDF's; none for
  unreferenced Zohar, Mishnah or Talmud, even in `apparatus` (Bible verses
  only). פ' + numeral → "chapter" for a book without parashot or a numeral
  beyond them, else "parashah" (noted). ד' + number after a book: the
  numerals ("4, 12"), noted when he writes דף nearby.
- **Section numbers.** His sections (ב) at 3:1 to קמד) at 10:6): the label →
  "§2."; irregular ones as written ("קיה)" → "§115", letters noted; the gap
  קכט)–קלא) noted, not repaired). List markers (א), ב)) → "1)", "2)" (note).
- **תא חזי** → "come and see".
- **אבותיהם** for "אבתיכם": follow his wording; the verse in `apparatus`.
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
## Style

- Plain, exact, modern English, the words at their own strength (אליל
  "idol"). No sarcasm of your own: no scare quotes, "so-called", "sic", or
  added "!". Keep his every "!", rhetorical question and disclaimer, and the
  second person.
- Long periods may be divided: no clause dropped, no connective that changes
  the logic, the order kept. Hebrew full stops: שכל / שהוא opening a sentence:
  the stop kept, a main clause by his syntax (note on the ש); a fragment
  (ושהם, והפטרה, ולדעת, ורבינו): a main clause when it continues the verb
  before, a comma when it continues a list (note); a main clause after a stop
  (ואם ראיתי. הם מגלים): joined with a comma (note); אם ... או after a stop:
  as punctuated if the English has a main clause, else a colon to the
  sentence before or a direct alternative question; a period completed in the
  next segment: the stops kept, the continuation in the segment `note`,
  semicolons only where unreadable.
- ל before a person's name is possessive (לרשב"י → "Rabbi Shimon's"), never
  the subject.
- Numbers: letter numerals → Arabic, as written whatever the tradition counts
  (the usual count in a note). A year without thousands → full ("תר"ע" →
  "5670"); the civil year in `apparatus`. Where gematria is the point, the
  number in English, the letters in a note.
- Witnesses: square brackets ("[מתחלת באלוהיות]") are the transcriber's
  corrections: translate the corrected reading, record both. A word that does
  not fit but is close to a likelier one (יכינו for יבינו): by the context; a
  printed word echoing a verse's likelier word (נבוש for נמוש): the printed
  word; both words in a note and `apparatus`. Rule 3 covers every
  Wikisource/PDF difference that bears on the English.

## Translator's checklist (the pilot reviews found these faults most often)

Before running the check, read your English once against the Hebrew:
1. Every Hebrew word has its English; none is added (no "who received it", no
   parenthetical explanation). A Hebrew full stop is an English full stop, and
   every sentence has a main clause.
2. Recognise Bible verses, even in pieces (Isa 62:6, Job 37:18, Gen 43:14,
   Lev 26:44); translate from the Hebrew, check each verb's form, list the
   verse in `apparatus`.
3. Use the fixed renderings of this file for every term, name and abbreviation.
4. A passive participle is passive; a participle is not a perfect.
5. Capitalise only what the file capitalises. Begin each sentence with a
   capital.

## Owner's answers and defaults

Answered by the owner (2026-11-10):

- Q4: Judeo-Arabic and Aramaic are Sonnet's, not Haiku's: `models.json` lists
  the four Judeo-Arabic units (5:3, 7:24, 9:10, 2:13) and the 15 units with at
  least six Aramaic marker words (a heuristic; the Sonnet review covers every
  unit). The `lang` tags stay "heb".
- Q7: the project's transliteration scheme (above).

Not yet answered; the default stands and can be changed before the review
stage (a change to 1, 3 or 5 is a find-and-patch):

- Q1 title: "The Wars of the LORD" (Num 21:14).
- Q2 Wikisource additions: translated as base text, unmarked per word; one
  general line in `apparatus` on the unit's first segment until the 1931 scan
  is collated.
- Q3 השי"ת: "the Name, may He be blessed".
- Q5 קבלה האמתית: "the true tradition", with a note.
- Q6 author's name: patronymic kept as written, "Sali'".


## Milchamot ha-Shem 2:19

Hebrew (Wikisource): האחד הפרש גדול בין העובד ומתפלל לאיזה דבר, ובין המכבד איזה דבר לבד. להתפלל ולעבוד עבודה לאיזה דבר אסור לנו, יהיה אדם או מלאך, גשם או רוחני אבל לחלוק כבוד כמו שיכבד הנער את הזקן, העבד את אדוניו, והנקלה לנכבד מותר. וכן מצאנו באברהם אבינו אשר האמין בה' רצץ ושבר מצבות אלילים וכיחש בם, השתחווה ארצה לפני שלשה מלאכים הנצבים עליו אף כי נדמו לו לאנשים. וכן עשה יהושע ביריחו בראותו איש עומד לנגדו ויחשבהו למלאך נפל לפניו ארצה ואם אנשי לבב כמוהם אשר בדרכיהם נלך וממצות שפתותיהם לא נבוש עשו זאת ולא נחשב להם לעון ע"ז ברור הוא כי לכבד ספר תורה אינה עבודת גילולים.

| id | word | punctuation after |
|---|---|---|
| 2:19.1 | האחד |  |
| 2:19.2 | הפרש |  |
| 2:19.3 | גדול |  |
| 2:19.4 | בין |  |
| 2:19.5 | העובד |  |
| 2:19.6 | ומתפלל |  |
| 2:19.7 | לאיזה |  |
| 2:19.8 | דבר | , |
| 2:19.9 | ובין |  |
| 2:19.10 | המכבד |  |
| 2:19.11 | איזה |  |
| 2:19.12 | דבר |  |
| 2:19.13 | לבד | . |
| 2:19.14 | להתפלל |  |
| 2:19.15 | ולעבוד |  |
| 2:19.16 | עבודה |  |
| 2:19.17 | לאיזה |  |
| 2:19.18 | דבר |  |
| 2:19.19 | אסור |  |
| 2:19.20 | לנו | , |
| 2:19.21 | יהיה |  |
| 2:19.22 | אדם |  |
| 2:19.23 | או |  |
| 2:19.24 | מלאך | , |
| 2:19.25 | גשם |  |
| 2:19.26 | או |  |
| 2:19.27 | רוחני |  |
| 2:19.28 | אבל |  |
| 2:19.29 | לחלוק |  |
| 2:19.30 | כבוד |  |
| 2:19.31 | כמו |  |
| 2:19.32 | שיכבד |  |
| 2:19.33 | הנער |  |
| 2:19.34 | את |  |
| 2:19.35 | הזקן | , |
| 2:19.36 | העבד |  |
| 2:19.37 | את |  |
| 2:19.38 | אדוניו | , |
| 2:19.39 | והנקלה |  |
| 2:19.40 | לנכבד |  |
| 2:19.41 | מותר | . |
| 2:19.42 | וכן |  |
| 2:19.43 | מצאנו |  |
| 2:19.44 | באברהם |  |
| 2:19.45 | אבינו |  |
| 2:19.46 | אשר |  |
| 2:19.47 | האמין |  |
| 2:19.48 | בה | ' |
| 2:19.49 | רצץ |  |
| 2:19.50 | ושבר |  |
| 2:19.51 | מצבות |  |
| 2:19.52 | אלילים |  |
| 2:19.53 | וכיחש |  |
| 2:19.54 | בם | , |
| 2:19.55 | השתחווה |  |
| 2:19.56 | ארצה |  |
| 2:19.57 | לפני |  |
| 2:19.58 | שלשה |  |
| 2:19.59 | מלאכים |  |
| 2:19.60 | הנצבים |  |
| 2:19.61 | עליו |  |
| 2:19.62 | אף |  |
| 2:19.63 | כי |  |
| 2:19.64 | נדמו |  |
| 2:19.65 | לו |  |
| 2:19.66 | לאנשים | . |
| 2:19.67 | וכן |  |
| 2:19.68 | עשה |  |
| 2:19.69 | יהושע |  |
| 2:19.70 | ביריחו |  |
| 2:19.71 | בראותו |  |
| 2:19.72 | איש |  |
| 2:19.73 | עומד |  |
| 2:19.74 | לנגדו |  |
| 2:19.75 | ויחשבהו |  |
| 2:19.76 | למלאך |  |
| 2:19.77 | נפל |  |
| 2:19.78 | לפניו |  |
| 2:19.79 | ארצה |  |
| 2:19.80 | ואם |  |
| 2:19.81 | אנשי |  |
| 2:19.82 | לבב |  |
| 2:19.83 | כמוהם |  |
| 2:19.84 | אשר |  |
| 2:19.85 | בדרכיהם |  |
| 2:19.86 | נלך |  |
| 2:19.87 | וממצות |  |
| 2:19.88 | שפתותיהם |  |
| 2:19.89 | לא |  |
| 2:19.90 | נבוש |  |
| 2:19.91 | עשו |  |
| 2:19.92 | זאת |  |
| 2:19.93 | ולא |  |
| 2:19.94 | נחשב |  |
| 2:19.95 | להם |  |
| 2:19.96 | לעון |  |
| 2:19.97 | ע"ז |  |
| 2:19.98 | ברור |  |
| 2:19.99 | הוא |  |
| 2:19.100 | כי |  |
| 2:19.101 | לכבד |  |
| 2:19.102 | ספר |  |
| 2:19.103 | תורה |  |
| 2:19.104 | אינה |  |
| 2:19.105 | עבודת |  |
| 2:19.106 | גילולים | . |


## Milchamot ha-Shem 2:20

Hebrew (Wikisource): שנית היהודים נזהרים מאד בענייני עבודה אחרת, ואף להראות כמכבדה אסור להם. כמו שאמרו בתלמוד וברמב"ם אם לפת איש ארחות דרכו אצל בית ע"ז וישב לו קוץ ברגלו לא ישח אף לא יכרע וישתחוה להוציאו שם למען לא יתראה בעיני העוברים כמשתחוה ומכבד לע"ז, ואם היתה ההשתחויה לפני הארון כעבודה אחרת לא היו היהודים מחזיקים במנהג זה, וכיון שהחזיקו בו מימי אבותיהם אות אמת הוא כי אינה משובשת.

| id | word | punctuation after |
|---|---|---|
| 2:20.1 | שנית |  |
| 2:20.2 | היהודים |  |
| 2:20.3 | נזהרים |  |
| 2:20.4 | מאד |  |
| 2:20.5 | בענייני |  |
| 2:20.6 | עבודה |  |
| 2:20.7 | אחרת | , |
| 2:20.8 | ואף |  |
| 2:20.9 | להראות |  |
| 2:20.10 | כמכבדה |  |
| 2:20.11 | אסור |  |
| 2:20.12 | להם | . |
| 2:20.13 | כמו |  |
| 2:20.14 | שאמרו |  |
| 2:20.15 | בתלמוד |  |
| 2:20.16 | וברמב"ם |  |
| 2:20.17 | אם |  |
| 2:20.18 | לפת |  |
| 2:20.19 | איש |  |
| 2:20.20 | ארחות |  |
| 2:20.21 | דרכו |  |
| 2:20.22 | אצל |  |
| 2:20.23 | בית |  |
| 2:20.24 | ע"ז |  |
| 2:20.25 | וישב |  |
| 2:20.26 | לו |  |
| 2:20.27 | קוץ |  |
| 2:20.28 | ברגלו |  |
| 2:20.29 | לא |  |
| 2:20.30 | ישח |  |
| 2:20.31 | אף |  |
| 2:20.32 | לא |  |
| 2:20.33 | יכרע |  |
| 2:20.34 | וישתחוה |  |
| 2:20.35 | להוציאו |  |
| 2:20.36 | שם |  |
| 2:20.37 | למען |  |
| 2:20.38 | לא |  |
| 2:20.39 | יתראה |  |
| 2:20.40 | בעיני |  |
| 2:20.41 | העוברים |  |
| 2:20.42 | כמשתחוה |  |
| 2:20.43 | ומכבד |  |
| 2:20.44 | לע"ז | , |
| 2:20.45 | ואם |  |
| 2:20.46 | היתה |  |
| 2:20.47 | ההשתחויה |  |
| 2:20.48 | לפני |  |
| 2:20.49 | הארון |  |
| 2:20.50 | כעבודה |  |
| 2:20.51 | אחרת |  |
| 2:20.52 | לא |  |
| 2:20.53 | היו |  |
| 2:20.54 | היהודים |  |
| 2:20.55 | מחזיקים |  |
| 2:20.56 | במנהג |  |
| 2:20.57 | זה | , |
| 2:20.58 | וכיון |  |
| 2:20.59 | שהחזיקו |  |
| 2:20.60 | בו |  |
| 2:20.61 | מימי |  |
| 2:20.62 | אבותיהם |  |
| 2:20.63 | אות |  |
| 2:20.64 | אמת |  |
| 2:20.65 | הוא |  |
| 2:20.66 | כי |  |
| 2:20.67 | אינה |  |
| 2:20.68 | משובשת | . |


## Milchamot ha-Shem 2:21

Hebrew (Wikisource): שלישית אחד ממנהגי הכבוד אשר כבדו לפנים את הע"ז היה הנשיקה וכן מצינו במקרא והשארתי בישראל שבעת אלפים כל הברכים אשר לא כרעו לבעל וכל הפה אשר לא נשק לו, ונאמר זובחי אדם עגלים ישקון. ואם כן הוא יהיו כל המנשקים את ספר התורה עובדי ע"ז? אבל לא כן הדבר, כי אין זה רק כבוד שיכבדו את תורתם ולא עבודה, ולכן ההשתחויה מול הארון הוא ג"כ רק כבוד ולא עבודה.

| id | word | punctuation after |
|---|---|---|
| 2:21.1 | שלישית |  |
| 2:21.2 | אחד |  |
| 2:21.3 | ממנהגי |  |
| 2:21.4 | הכבוד |  |
| 2:21.5 | אשר |  |
| 2:21.6 | כבדו |  |
| 2:21.7 | לפנים |  |
| 2:21.8 | את |  |
| 2:21.9 | הע"ז |  |
| 2:21.10 | היה |  |
| 2:21.11 | הנשיקה |  |
| 2:21.12 | וכן |  |
| 2:21.13 | מצינו |  |
| 2:21.14 | במקרא |  |
| 2:21.15 | והשארתי |  |
| 2:21.16 | בישראל |  |
| 2:21.17 | שבעת |  |
| 2:21.18 | אלפים |  |
| 2:21.19 | כל |  |
| 2:21.20 | הברכים |  |
| 2:21.21 | אשר |  |
| 2:21.22 | לא |  |
| 2:21.23 | כרעו |  |
| 2:21.24 | לבעל |  |
| 2:21.25 | וכל |  |
| 2:21.26 | הפה |  |
| 2:21.27 | אשר |  |
| 2:21.28 | לא |  |
| 2:21.29 | נשק |  |
| 2:21.30 | לו | , |
| 2:21.31 | ונאמר |  |
| 2:21.32 | זובחי |  |
| 2:21.33 | אדם |  |
| 2:21.34 | עגלים |  |
| 2:21.35 | ישקון | . |
| 2:21.36 | ואם |  |
| 2:21.37 | כן |  |
| 2:21.38 | הוא |  |
| 2:21.39 | יהיו |  |
| 2:21.40 | כל |  |
| 2:21.41 | המנשקים |  |
| 2:21.42 | את |  |
| 2:21.43 | ספר |  |
| 2:21.44 | התורה |  |
| 2:21.45 | עובדי |  |
| 2:21.46 | ע"ז | ? |
| 2:21.47 | אבל |  |
| 2:21.48 | לא |  |
| 2:21.49 | כן |  |
| 2:21.50 | הדבר | , |
| 2:21.51 | כי |  |
| 2:21.52 | אין |  |
| 2:21.53 | זה |  |
| 2:21.54 | רק |  |
| 2:21.55 | כבוד |  |
| 2:21.56 | שיכבדו |  |
| 2:21.57 | את |  |
| 2:21.58 | תורתם |  |
| 2:21.59 | ולא |  |
| 2:21.60 | עבודה | , |
| 2:21.61 | ולכן |  |
| 2:21.62 | ההשתחויה |  |
| 2:21.63 | מול |  |
| 2:21.64 | הארון |  |
| 2:21.65 | הוא |  |
| 2:21.66 | ג"כ |  |
| 2:21.67 | רק |  |
| 2:21.68 | כבוד |  |
| 2:21.69 | ולא |  |
| 2:21.70 | עבודה | . |


## Milchamot ha-Shem 2:22

Hebrew (Wikisource): רביעית הנסיון הורה אותנו כי בכל לשונות הגוים נמצאו מנהגי הכבוד יכבדו בם איש את רעהו בהיפגשם וכן יכופו ראשיהם זה לעומת זה, והיה גם זה לנו לאות שאינה עבודת גלולים כי אם מנהגי כבוד.

| id | word | punctuation after |
|---|---|---|
| 2:22.1 | רביעית |  |
| 2:22.2 | הנסיון |  |
| 2:22.3 | הורה |  |
| 2:22.4 | אותנו |  |
| 2:22.5 | כי |  |
| 2:22.6 | בכל |  |
| 2:22.7 | לשונות |  |
| 2:22.8 | הגוים |  |
| 2:22.9 | נמצאו |  |
| 2:22.10 | מנהגי |  |
| 2:22.11 | הכבוד |  |
| 2:22.12 | יכבדו |  |
| 2:22.13 | בם |  |
| 2:22.14 | איש |  |
| 2:22.15 | את |  |
| 2:22.16 | רעהו |  |
| 2:22.17 | בהיפגשם |  |
| 2:22.18 | וכן |  |
| 2:22.19 | יכופו |  |
| 2:22.20 | ראשיהם |  |
| 2:22.21 | זה |  |
| 2:22.22 | לעומת |  |
| 2:22.23 | זה | , |
| 2:22.24 | והיה |  |
| 2:22.25 | גם |  |
| 2:22.26 | זה |  |
| 2:22.27 | לנו |  |
| 2:22.28 | לאות |  |
| 2:22.29 | שאינה |  |
| 2:22.30 | עבודת |  |
| 2:22.31 | גלולים |  |
| 2:22.32 | כי |  |
| 2:22.33 | אם |  |
| 2:22.34 | מנהגי |  |
| 2:22.35 | כבוד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

