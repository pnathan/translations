# Review pack: milhamot-hashem perek-2-11 (Milchamot ha-Shem 2:23–2:27)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-11`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **קצוצי** פאה (Jer 9:25) → "those with clipped temples" (context note: desert
  peoples).
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 2:23

Hebrew (Wikisource): חמישית העם היושבים באזיא אם יבוא אליהם מכתב פקודה מהמלך בכפם יתפשוהו, בפיהם ישקוהו, ועל ראשם יניחוהו, אף כי אנו נכבד את דברי אלהים חיים מלך על כל הארץ.

| id | word | punctuation after |
|---|---|---|
| 2:23.1 | חמישית |  |
| 2:23.2 | העם |  |
| 2:23.3 | היושבים |  |
| 2:23.4 | באזיא |  |
| 2:23.5 | אם |  |
| 2:23.6 | יבוא |  |
| 2:23.7 | אליהם |  |
| 2:23.8 | מכתב |  |
| 2:23.9 | פקודה |  |
| 2:23.10 | מהמלך |  |
| 2:23.11 | בכפם |  |
| 2:23.12 | יתפשוהו | , |
| 2:23.13 | בפיהם |  |
| 2:23.14 | ישקוהו | , |
| 2:23.15 | ועל |  |
| 2:23.16 | ראשם |  |
| 2:23.17 | יניחוהו | , |
| 2:23.18 | אף |  |
| 2:23.19 | כי |  |
| 2:23.20 | אנו |  |
| 2:23.21 | נכבד |  |
| 2:23.22 | את |  |
| 2:23.23 | דברי |  |
| 2:23.24 | אלהים |  |
| 2:23.25 | חיים |  |
| 2:23.26 | מלך |  |
| 2:23.27 | על |  |
| 2:23.28 | כל |  |
| 2:23.29 | הארץ | . |


### Translation under review (2:23)

English: "The fifth: the people who dwell in Asia, if a written command from the king comes to them, they take it in their hands, kiss it with their mouths and place it on their heads. How much more so shall we honor the words of the living God, King over all the earth.

Literal: fifth the-people the-dwelling in-Asia if comes to-them written-command from-the-king in-their-hands they-seize-it, with-their-mouths they-kiss-it, and-on their-heads they-place-it, how-much-more so we-shall-honor the-words-of God living King over all the-earth

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:23.1 | fifth |  |  | ordinal of a numbered series: 2:24 is the sixth, 2:25 the seventh |
| 2:23.2 | the people |  |  |  |
| 2:23.3 | dwelling |  |  |  |
| 2:23.4 | in Asia |  |  |  |
| 2:23.5 | if |  |  |  |
| 2:23.6 | comes |  |  | yiqtol after אם: the protasis of a conditional; the verbs that follow are yiqtol in a habitual sense |
| 2:23.7 | to them |  |  |  |
| 2:23.8 | written |  |  | מכתב, 'letter, writing' |
| 2:23.9 | command |  |  | פקודה, an order |
| 2:23.10 | from the king |  |  |  |
| 2:23.11 | in their hands |  |  |  |
| 2:23.12 | they take it |  |  | יתפשוהו: root תפש with the suffix ־הו, 'it'; the command is the object |
| 2:23.13 | with their mouths |  |  |  |
| 2:23.14 | they kiss it |  |  |  |
| 2:23.15 | and on |  |  |  |
| 2:23.16 | their heads |  |  |  |
| 2:23.17 | they place it |  |  |  |
| 2:23.18 | how much more |  |  | אף כי: 'how much more so', the apodosis of an argument from the lesser to the greater |
| 2:23.19 | so |  |  |  |
| 2:23.20 | we |  |  |  |
| 2:23.21 | shall honor |  |  | נכבד: cohortative, 'let us honor' |
| 2:23.22 | - |  |  | object marker את, no English |
| 2:23.23 | the words of |  |  |  |
| 2:23.24 | God |  |  | אלהים: the God of Israel, here 'the living God' (חיים) |
| 2:23.25 | living |  |  |  |
| 2:23.26 | King |  |  |  |
| 2:23.27 | over |  |  |  |
| 2:23.28 | all |  |  |  |
| 2:23.29 | the earth |  |  |  |

Phrases:
- 2:23.1 2:23.2 2:23.3 2:23.4 | The fifth: the people who dwell in Asia, | fifth the-people the-dwelling in-Asia
- 2:23.5 2:23.6 2:23.7 2:23.8 2:23.9 2:23.10 | if a written command from the king comes to them, | if comes to-them written-command from-the-king
- 2:23.11 2:23.12 | they take it in their hands, | in-their-hands they-seize-it
- 2:23.13 2:23.14 | kiss it with their mouths, | with-their-mouths they-kiss-it
- 2:23.15 2:23.16 2:23.17 | and place it on their heads. | and-on their-heads they-place-it
- 2:23.18 2:23.19 | How much more so | how-much-more so
- 2:23.20 2:23.21 | shall we honor | we shall-honor
- 2:23.22 | - | object marker, no English
- 2:23.23 2:23.24 2:23.25 | the words of the living God, | the-words-of God living
- 2:23.26 2:23.27 2:23.28 2:23.29 | King over all the earth. | King over all the-earth

Grammar notes:
- The paragraph is the fifth point of a numbered series, and its number opens the English as it opens the Hebrew.
- The conditional אם takes yiqtol verbs (they take, kiss, place): these describe a habitual practice, so the English uses the present.
- The argument is from the lesser to the greater (qal wa-chomer): if the people honor a king's written command, 'how much more so' the words of God. The cohortative נכבד keeps the 'we shall'.
- The three acts (seizing in the hands, kissing with the mouth, placing on the head) are parallel clauses, and the English keeps their order.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated, so the paragraph is translated as the base text (decisions, Q2).
- apparatus: Quotation: the double quotation marks open at this segment because the quoted letter is closed by עכ"ל at 2:25. The opening formula is not in this unit.
- apparatus: Allusions, the author's wording and not a quotation: 'the living God' (Jer 10:10) and 'King over all the earth' (cf. Zech 14:9).

## Milchamot ha-Shem 2:24

Hebrew (Wikisource): ששית כשהביאו שבעים ושנים זקנים את התורה לתלמי מלך מצרים, קם מכסאו והשתחוה ז' פעמים לקראתה, ואם עובד אלילים אשר לא הביא צוארו בעול תורתינו כבדה כל כך, אף כי אנחנו הבאים בבריתה ואשר ניתנה לנו למורשה.

| id | word | punctuation after |
|---|---|---|
| 2:24.1 | ששית |  |
| 2:24.2 | כשהביאו |  |
| 2:24.3 | שבעים |  |
| 2:24.4 | ושנים |  |
| 2:24.5 | זקנים |  |
| 2:24.6 | את |  |
| 2:24.7 | התורה |  |
| 2:24.8 | לתלמי |  |
| 2:24.9 | מלך |  |
| 2:24.10 | מצרים | , |
| 2:24.11 | קם |  |
| 2:24.12 | מכסאו |  |
| 2:24.13 | והשתחוה |  |
| 2:24.14 | ז | ' |
| 2:24.15 | פעמים |  |
| 2:24.16 | לקראתה | , |
| 2:24.17 | ואם |  |
| 2:24.18 | עובד |  |
| 2:24.19 | אלילים |  |
| 2:24.20 | אשר |  |
| 2:24.21 | לא |  |
| 2:24.22 | הביא |  |
| 2:24.23 | צוארו |  |
| 2:24.24 | בעול |  |
| 2:24.25 | תורתינו |  |
| 2:24.26 | כבדה |  |
| 2:24.27 | כל |  |
| 2:24.28 | כך | , |
| 2:24.29 | אף |  |
| 2:24.30 | כי |  |
| 2:24.31 | אנחנו |  |
| 2:24.32 | הבאים |  |
| 2:24.33 | בבריתה |  |
| 2:24.34 | ואשר |  |
| 2:24.35 | ניתנה |  |
| 2:24.36 | לנו |  |
| 2:24.37 | למורשה | . |


### Translation under review (2:24)

English: The sixth: when the seventy-two elders brought the Torah to Ptolemy, king of Egypt, he rose from his throne and bowed down seven times to meet it, and if a worshipper of idols who has not put his neck under the yoke of our Torah, which is so heavy, how much more so we who enter into its covenant and to whom it was given as an inheritance.

Literal: sixth when-they-brought seventy and-two elders (object) the-Torah to-Ptolemy king of-Egypt rose from-his-throne and-bowed-down seven times to-meet-it, and-if worshipper-of idols who not put his-neck into-the-yoke of-our-Torah is-heavy so much, how-much-more so we entering in-its-covenant and-to-whom it-was-given to-us as-inheritance

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:24.1 | sixth |  |  |  |
| 2:24.2 | when they brought |  |  | כשהביאו: 'when' with a qatal; the verb comes before its subject |
| 2:24.3 | seventy |  |  |  |
| 2:24.4 | and two |  |  | שבעים ושנים: seventy-two, written 'seventy and two' |
| 2:24.5 | elders |  |  |  |
| 2:24.6 | - |  |  | object marker את, no English |
| 2:24.7 | the Torah |  |  |  |
| 2:24.8 | to Ptolemy |  |  | תלמי, the Hebrew spelling of the king's name |
| 2:24.9 | king |  |  |  |
| 2:24.10 | of Egypt |  |  |  |
| 2:24.11 | rose |  |  |  |
| 2:24.12 | from his throne |  |  | מכסאו: 'from his throne' |
| 2:24.13 | and bowed down |  |  | והשתחוה: the act of reverence; the verb is the same as in 2:26 |
| 2:24.14 | seven |  |  | ז׳ is the numeral seven, not a letter |
| 2:24.15 | times |  |  |  |
| 2:24.16 | to meet it |  |  | לקראתה: lit. 'to meet it', the greeting of a sacred object |
| 2:24.17 | and if |  |  |  |
| 2:24.18 | a worshipper |  |  | עובד אלילים: 'worshipper of idols' |
| 2:24.19 | of idols |  |  |  |
| 2:24.20 | who |  |  |  |
| 2:24.21 | has not |  |  |  |
| 2:24.22 | put |  |  | הביא צוארו בעול: lit. 'brought his neck into the yoke', an idiom of accepting a burden |
| 2:24.23 | his neck |  |  |  |
| 2:24.24 | under the yoke |  |  | בעול: the yoke, the image of taking on the Torah |
| 2:24.25 | of our Torah |  |  |  |
| 2:24.26 | which is heavy |  |  | כבדה: feminine, agreeing with תורתינו; the Torah is what is heavy, not the yoke |
| 2:24.27 | so |  |  |  |
| 2:24.28 | much |  |  |  |
| 2:24.29 | how much more |  |  | אף כי: 'how much more so', as in 2:23 |
| 2:24.30 | so |  |  |  |
| 2:24.31 | we |  |  |  |
| 2:24.32 | who enter |  |  | הבאים: participle, 'who enter' |
| 2:24.33 | into its covenant |  |  | בבריתה: the covenant of the Torah; 'its' is the Torah |
| 2:24.34 | and to whom |  |  |  |
| 2:24.35 | it was given |  |  |  |
| 2:24.36 | to us |  |  |  |
| 2:24.37 | as an inheritance |  |  | למורשה: 'as an inheritance'; see apparatus |

Phrases:
- 2:24.1 | The sixth: | sixth
- 2:24.2 2:24.3 2:24.4 2:24.5 | when the seventy-two elders brought | when-brought seventy and-two elders
- 2:24.6 2:24.7 | the Torah | (object) the-Torah
- 2:24.8 2:24.9 2:24.10 | to Ptolemy, king of Egypt, | to-Ptolemy king of-Egypt
- 2:24.11 2:24.12 | he rose from his throne | rose from-his-throne
- 2:24.13 2:24.14 2:24.15 2:24.16 | and bowed down seven times to meet it, | and-bowed-down seven times to-meet-it
- 2:24.17 2:24.18 2:24.19 | and if a worshipper of idols | and-if worshipper-of idols
- 2:24.20 2:24.21 2:24.22 2:24.23 2:24.24 2:24.25 | who has not put his neck under the yoke of our Torah, | who not brought his-neck into-the-yoke of-our-Torah
- 2:24.26 2:24.27 2:24.28 | which is so heavy, | is-heavy so much
- 2:24.29 2:24.30 | how much more so | how-much-more so
- 2:24.31 2:24.32 2:24.33 | we who enter into its covenant | we entering in-its-covenant
- 2:24.34 2:24.35 2:24.36 2:24.37 | and to whom it was given as an inheritance. | and-to-whom it-was-given to-us as-inheritance

Grammar notes:
- The temporal clause כשהביאו puts the verb first and the subject after it: 'when the seventy-two elders brought the Torah'.
- קם ... והשתחוה are two past acts in sequence (qatal with vav); the English keeps the sequence, 'rose ... and bowed down'.
- The clause introduced by ואם has no verb. The Hebrew leaves the king's act to be understood, and the English keeps the ellipsis; the apodosis 'how much more so we' answers it.
- כבדה is feminine and agrees with תורתינו, so 'which is so heavy' describes the Torah, not the yoke.

Commentary rulings:
- apparatus: Allusion, not quoted: Deut 33:4 for 'למורשה' ('an inheritance').
- apparatus: The number seventy-two and the translation story are the author's (see context).

## Milchamot ha-Shem 2:25

Hebrew (Wikisource): שביעית יסודי ועיקרי דתינו הם מציאות בורא אחד יחיד קדמון אשר איננו גשם ולא כח בגשם, נותן התורה לעמו ישראל על ידי המחוקק הנאמן משה עבד ה' אדון כל הנביאים. הוא יתברך המציץ ומשגיח בעולמות שברא. ועיניו משוטטות בכל, וחוקר ובוחן לבות בני אדם ושופט בצדק על עלילות מצעדי גבר. ממנו לבד נבקש עזרה בצר לנו, והוא יפרה נצר מגזע ישי אשר יקבץ נדחי ישראל מארבע כנפות הארץ, ומכל קצוצי פאה אשר זרו אותם שמה, עצמות מתיהם היבשות תחינה ונבלתם יקומון. הן כל אלה עקרי ושרשי דתינו. ומה שמץ ע"ז נמצא בם? אף לדעת אותם אשר דעותיהם ואמונותיהם נבדלו מדעותינו ומאמונתנו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 2:25.1 | שביעית |  |
| 2:25.2 | יסודי |  |
| 2:25.3 | ועיקרי |  |
| 2:25.4 | דתינו |  |
| 2:25.5 | הם |  |
| 2:25.6 | מציאות |  |
| 2:25.7 | בורא |  |
| 2:25.8 | אחד |  |
| 2:25.9 | יחיד |  |
| 2:25.10 | קדמון |  |
| 2:25.11 | אשר |  |
| 2:25.12 | איננו |  |
| 2:25.13 | גשם |  |
| 2:25.14 | ולא |  |
| 2:25.15 | כח |  |
| 2:25.16 | בגשם | , |
| 2:25.17 | נותן |  |
| 2:25.18 | התורה |  |
| 2:25.19 | לעמו |  |
| 2:25.20 | ישראל |  |
| 2:25.21 | על |  |
| 2:25.22 | ידי |  |
| 2:25.23 | המחוקק |  |
| 2:25.24 | הנאמן |  |
| 2:25.25 | משה |  |
| 2:25.26 | עבד |  |
| 2:25.27 | ה | ' |
| 2:25.28 | אדון |  |
| 2:25.29 | כל |  |
| 2:25.30 | הנביאים | . |
| 2:25.31 | הוא |  |
| 2:25.32 | יתברך |  |
| 2:25.33 | המציץ |  |
| 2:25.34 | ומשגיח |  |
| 2:25.35 | בעולמות |  |
| 2:25.36 | שברא | . |
| 2:25.37 | ועיניו |  |
| 2:25.38 | משוטטות |  |
| 2:25.39 | בכל | , |
| 2:25.40 | וחוקר |  |
| 2:25.41 | ובוחן |  |
| 2:25.42 | לבות |  |
| 2:25.43 | בני |  |
| 2:25.44 | אדם |  |
| 2:25.45 | ושופט |  |
| 2:25.46 | בצדק |  |
| 2:25.47 | על |  |
| 2:25.48 | עלילות |  |
| 2:25.49 | מצעדי |  |
| 2:25.50 | גבר | . |
| 2:25.51 | ממנו |  |
| 2:25.52 | לבד |  |
| 2:25.53 | נבקש |  |
| 2:25.54 | עזרה |  |
| 2:25.55 | בצר |  |
| 2:25.56 | לנו | , |
| 2:25.57 | והוא |  |
| 2:25.58 | יפרה |  |
| 2:25.59 | נצר |  |
| 2:25.60 | מגזע |  |
| 2:25.61 | ישי |  |
| 2:25.62 | אשר |  |
| 2:25.63 | יקבץ |  |
| 2:25.64 | נדחי |  |
| 2:25.65 | ישראל |  |
| 2:25.66 | מארבע |  |
| 2:25.67 | כנפות |  |
| 2:25.68 | הארץ | , |
| 2:25.69 | ומכל |  |
| 2:25.70 | קצוצי |  |
| 2:25.71 | פאה |  |
| 2:25.72 | אשר |  |
| 2:25.73 | זרו |  |
| 2:25.74 | אותם |  |
| 2:25.75 | שמה | , |
| 2:25.76 | עצמות |  |
| 2:25.77 | מתיהם |  |
| 2:25.78 | היבשות |  |
| 2:25.79 | תחינה |  |
| 2:25.80 | ונבלתם |  |
| 2:25.81 | יקומון | . |
| 2:25.82 | הן |  |
| 2:25.83 | כל |  |
| 2:25.84 | אלה |  |
| 2:25.85 | עקרי |  |
| 2:25.86 | ושרשי |  |
| 2:25.87 | דתינו | . |
| 2:25.88 | ומה |  |
| 2:25.89 | שמץ |  |
| 2:25.90 | ע"ז |  |
| 2:25.91 | נמצא |  |
| 2:25.92 | בם | ? |
| 2:25.93 | אף |  |
| 2:25.94 | לדעת |  |
| 2:25.95 | אותם |  |
| 2:25.96 | אשר |  |
| 2:25.97 | דעותיהם |  |
| 2:25.98 | ואמונותיהם |  |
| 2:25.99 | נבדלו |  |
| 2:25.100 | מדעותינו |  |
| 2:25.101 | ומאמונתנו |  |
| 2:25.102 | עכ"ל | . |


### Translation under review (2:25)

English: "Seventh: the foundation and principle of our religion is the existence of one, unique, eternal Creator, who is not a body, nor a force in a body, who gives the Torah to His people Israel through the faithful lawgiver Moses, servant of the LORD, master of all the prophets. He, may He be blessed, looks on and exercises providence over the worlds that He created. And His eyes roam over all, He searches and examines the hearts of the sons of man and judges justly according to the deeds, the steps of a man. From Him alone we seek help in our distress, and He will bring forth a shoot from the stock of Jesse, who will gather the outcasts of Israel from the four corners of the earth, and from all the cut-off corners to which they scattered them there, the dry bones of their dead will live, and their corpses will rise. Indeed, all these are the principles and roots of our religion. And what trace of idolatry is found in them? Even concerning those whose opinions and beliefs differ from our opinions and from our faith." Thus far his words.

Literal: seventh foundation and-principle of-our-religion are existence of-Creator one unique eternal who is-not body nor force in-body giving the-Torah to-his-people Israel through the-hand-of the-lawgiver the-faithful Moses servant of-LORD master of-all the-prophets he may-he-be-blessed who-looks-on and-exercises-providence over the-worlds he-created and-his-eyes roam over-all and-searches and-examines the-hearts of-sons of-man and-judges justly according-to the-deeds the-steps-of a-man from-him alone we-seek help in-distress our and-he will-bear-fruit a-shoot from-stock-of Jesse who will-gather outcasts-of Israel from-four corners-of the-earth and-from all the-cut-off corners to-which they-scattered them there the-bones of-their-dead dry will-live and-their-corpses will-rise indeed all these the-principles and-roots-of our-religion and-what trace idolatry is-found in-them even concerning those whose opinions and-beliefs differ from-our-opinions and-from-our-faith thus far his-words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:25.1 | seventh |  |  |  |
| 2:25.2 | foundation |  |  |  |
| 2:25.3 | and principle |  |  |  |
| 2:25.4 | of our religion |  |  |  |
| 2:25.5 | is |  |  |  |
| 2:25.6 | existence |  |  |  |
| 2:25.7 | of Creator |  |  |  |
| 2:25.8 | one |  |  |  |
| 2:25.9 | unique |  |  |  |
| 2:25.10 | eternal |  |  | קדמון: eternal, 'without beginning' |
| 2:25.11 | who |  |  |  |
| 2:25.12 | is not |  |  |  |
| 2:25.13 | a body |  |  |  |
| 2:25.14 | nor |  |  |  |
| 2:25.15 | a force |  |  |  |
| 2:25.16 | in a body |  |  | גשם ולא כח בגשם: the form of the Rambam's principle, 'a body or a force in a body' |
| 2:25.17 | gives |  |  |  |
| 2:25.18 | the Torah |  |  |  |
| 2:25.19 | to His people |  |  |  |
| 2:25.20 | Israel |  |  |  |
| 2:25.21 | through |  |  |  |
| 2:25.22 | the hand of |  |  | על ידי: lit. 'upon the hand of', 'through' |
| 2:25.23 | the lawgiver |  |  |  |
| 2:25.24 | the faithful |  |  |  |
| 2:25.25 | Moses |  |  |  |
| 2:25.26 | servant |  |  |  |
| 2:25.27 | LORD |  |  | ה׳: the Tetragrammaton |
| 2:25.28 | master |  |  |  |
| 2:25.29 | of all |  |  |  |
| 2:25.30 | the prophets |  |  |  |
| 2:25.31 | He |  |  |  |
| 2:25.32 | may He be blessed |  |  | יתברך: 'may He be blessed' |
| 2:25.33 | who looks on |  |  | המציץ: participle, 'who looks on' |
| 2:25.34 | and exercises providence over |  |  | ומשגיח: participle; השגחה is 'providence' |
| 2:25.35 | the worlds |  |  |  |
| 2:25.36 | He created |  |  |  |
| 2:25.37 | and His eyes |  |  |  |
| 2:25.38 | roam |  |  |  |
| 2:25.39 | over all |  |  |  |
| 2:25.40 | and searches |  |  |  |
| 2:25.41 | and examines |  |  |  |
| 2:25.42 | the hearts |  |  |  |
| 2:25.43 | of the sons |  |  |  |
| 2:25.44 | of man |  |  |  |
| 2:25.45 | and judges |  |  |  |
| 2:25.46 | justly |  |  |  |
| 2:25.47 | according to |  |  |  |
| 2:25.48 | the deeds |  |  |  |
| 2:25.49 | the steps of |  |  |  |
| 2:25.50 | a man |  |  |  |
| 2:25.51 | from Him |  |  |  |
| 2:25.52 | alone |  |  |  |
| 2:25.53 | we seek |  |  |  |
| 2:25.54 | help |  |  |  |
| 2:25.55 | in distress |  |  |  |
| 2:25.56 | our |  |  |  |
| 2:25.57 | and He |  |  |  |
| 2:25.58 | will bring forth |  |  | יפרה: future; Isa 11:1 has this verb intransitively, 'bear fruit' |
| 2:25.59 | a shoot |  |  | נצר: 'shoot' or 'branch' |
| 2:25.60 | from the stock of |  |  |  |
| 2:25.61 | Jesse |  |  |  |
| 2:25.62 | who |  |  |  |
| 2:25.63 | will gather |  |  |  |
| 2:25.64 | the outcasts of |  |  | נדחי: 'the dispersed' |
| 2:25.65 | Israel |  |  |  |
| 2:25.66 | from the four |  |  |  |
| 2:25.67 | corners |  |  | כנפות: lit. 'wings', 'corners' |
| 2:25.68 | of the earth |  |  |  |
| 2:25.69 | and from all |  |  |  |
| 2:25.70 | the cut-off |  |  | קצוצי פאה: see context |
| 2:25.71 | corners |  |  |  |
| 2:25.72 | to which |  |  |  |
| 2:25.73 | they scattered |  |  |  |
| 2:25.74 | them |  |  |  |
| 2:25.75 | there |  |  |  |
| 2:25.76 | the bones |  |  |  |
| 2:25.77 | of their dead |  |  |  |
| 2:25.78 | dry |  |  |  |
| 2:25.79 | will live |  |  | תחינה: fem. plural yiqtol, the bones as subject |
| 2:25.80 | and their corpses |  |  |  |
| 2:25.81 | will rise |  |  |  |
| 2:25.82 | indeed |  |  |  |
| 2:25.83 | all |  |  |  |
| 2:25.84 | these |  |  |  |
| 2:25.85 | the principles |  |  |  |
| 2:25.86 | and roots of |  |  |  |
| 2:25.87 | our religion |  |  |  |
| 2:25.88 | and what |  |  |  |
| 2:25.89 | trace |  |  |  |
| 2:25.90 | idolatry |  |  | ע״ז = עבודה זרה, lit. 'strange worship' |
| 2:25.91 | is found |  |  |  |
| 2:25.92 | in them |  |  |  |
| 2:25.93 | even |  |  |  |
| 2:25.94 | concerning |  |  | לדעת: an infinitive used as 'concerning'; the sentence has no verb |
| 2:25.95 | those |  |  |  |
| 2:25.96 | whose |  |  |  |
| 2:25.97 | opinions |  |  |  |
| 2:25.98 | and beliefs |  |  |  |
| 2:25.99 | differ |  |  |  |
| 2:25.100 | from our opinions |  |  |  |
| 2:25.101 | and from our faith |  |  |  |
| 2:25.102 | thus far his words |  |  | עכ״ל: the closing formula of the quotation |

Phrases:
- 2:25.1 2:25.2 2:25.3 2:25.4 | Seventh: the foundation and principle of our religion | seventh foundation and-principle of-our-religion
- 2:25.5 2:25.6 2:25.7 2:25.8 2:25.9 2:25.10 | is the existence of one, unique, eternal Creator, | are existence of-Creator one unique eternal
- 2:25.11 2:25.12 2:25.13 2:25.14 2:25.15 2:25.16 | who is not a body, nor a force in a body, | who is-not body nor force in-body
- 2:25.17 2:25.18 2:25.19 2:25.20 | who gives the Torah to His people Israel | giving the-Torah to-his-people Israel
- 2:25.21 2:25.22 2:25.23 2:25.24 2:25.25 | through the faithful lawgiver Moses, | through the-hand-of the-lawgiver the-faithful Moses
- 2:25.26 2:25.27 2:25.28 2:25.29 2:25.30 | servant of the LORD, master of all the prophets. | servant of LORD master of-all the-prophets
- 2:25.31 2:25.32 2:25.33 2:25.34 2:25.35 2:25.36 | He, may He be blessed, looks on and exercises providence over the worlds that He created. | he may-he-be-blessed who-looks-on and-exercises-providence over the-worlds he-created
- 2:25.37 2:25.38 2:25.39 | And His eyes roam over all, | and-his-eyes roam over-all
- 2:25.40 2:25.41 2:25.42 2:25.43 2:25.44 2:25.45 2:25.46 | He searches and examines the hearts of the sons of man and judges justly | searches and-examines the-hearts of-sons of-man and-judges justly
- 2:25.47 2:25.48 2:25.49 2:25.50 | according to the deeds, the steps of a man. | according-to the-deeds the-steps-of a-man
- 2:25.51 2:25.52 2:25.53 2:25.54 2:25.55 2:25.56 | From Him alone we seek help in our distress, | from-Him alone we-seek help in-distress our
- 2:25.57 2:25.58 2:25.59 2:25.60 2:25.61 | and He will bring forth a shoot from the stock of Jesse, | and-he will-bear-fruit a-shoot from-stock-of Jesse
- 2:25.62 2:25.63 2:25.64 2:25.65 2:25.66 2:25.67 2:25.68 | who will gather the outcasts of Israel from the four corners of the earth, | who will-gather outcasts-of Israel from-four corners-of the-earth
- 2:25.69 2:25.70 2:25.71 2:25.72 2:25.73 2:25.74 2:25.75 | and from all the cut-off corners to which they scattered them there, | and-from all the-cut-off corners to-which they-scattered them there
- 2:25.76 2:25.77 2:25.78 2:25.79 2:25.80 2:25.81 | the dry bones of their dead will live, and their corpses will rise. | the-bones of-their-dead dry will-live and-their-corpses will-rise
- 2:25.82 2:25.83 2:25.84 2:25.85 2:25.86 2:25.87 | Indeed, all these are the principles and roots of our religion. | indeed all these the-principles and-roots-of our-religion
- 2:25.88 2:25.89 2:25.90 2:25.91 2:25.92 | And what trace of idolatry is found in them? | and-what trace idolatry is-found in-them
- 2:25.93 2:25.94 2:25.95 | Even concerning those | even concerning those
- 2:25.96 2:25.97 2:25.98 2:25.99 2:25.100 2:25.101 | whose opinions and beliefs differ from our opinions and from our faith. | whose opinions and-beliefs differ from-our-opinions and-from-our-faith
- 2:25.102 | Thus far his words. | thus far

Grammar notes:
- The paragraph is one sentence: 'the foundation ... is the existence of ...'. Its relative clauses (אשר, נותן) continue from the Creator, and the English keeps the sentence long.
- The verbs of providence (המציץ, ומשגיח, חוקר, ובוחן, ושופט) are participles with no finite verb; the English makes them finite with 'He'.
- והוא יפרה נצר: Isa 11:1 has this verb intransitively ('shall bear fruit'); here it takes נצר as its object, and the English follows the author ('bring forth a shoot').
- The last sentence has no verb ('even concerning those ...'); the English keeps the fragment, as the Hebrew does.

Commentary rulings:
- apparatus: Quotation: the double quotation marks close here with the formula עכ"ל. The opening is outside this unit (see 2:23 apparatus).
- apparatus: Allusions, the author's wording and not quotations: Jer 17:10 (חוקר לב בוחן) for 'searches and examines the hearts'; Prov 20:24 (מצעדי גבר) for 'the steps of a man'; Isa 11:1 for 'a shoot from the stock of Jesse'; Isa 11:12 (מארבע כנפות הארץ) for 'from the four corners of the earth'; Ezek 37 (the dry bones) and Isa 26:19 (נבלתי יקומון) for the resurrection clause; Jer 9:25 for 'cut-off corners'.

## Milchamot ha-Shem 2:26

Hebrew (Wikisource): ולפי דרכינו דרך רז"ל בבריתא הנז"ל אין צורך לכל זה שאין כוונתנו שאנו משתחוים מול הס"ת אלא למי שבחר בתורה, ובהדיא אנו אומרים בואו נשתחוה ונכרעה נברכה לפני ה' עושנו. והשתחוית אברהם אבינו למלאכים וכן יהושע, איבעית אימא כיוון שנדמו להם כאנשים קרינא בהו מאי דאמור רבנן בבריתא לא תשתחוה להם, להם אי אתה משתחוה אבל אתה משתחוה לאדם שכמותך כיון שנגלו להם בדמות אנשים ואיבעית אימא שכוונת אברהם ויהושע להשתחות למי ששלחם וכמו שכ' בעל העיקרים.

| id | word | punctuation after |
|---|---|---|
| 2:26.1 | ולפי |  |
| 2:26.2 | דרכינו |  |
| 2:26.3 | דרך |  |
| 2:26.4 | רז"ל |  |
| 2:26.5 | בבריתא |  |
| 2:26.6 | הנז"ל |  |
| 2:26.7 | אין |  |
| 2:26.8 | צורך |  |
| 2:26.9 | לכל |  |
| 2:26.10 | זה |  |
| 2:26.11 | שאין |  |
| 2:26.12 | כוונתנו |  |
| 2:26.13 | שאנו |  |
| 2:26.14 | משתחוים |  |
| 2:26.15 | מול |  |
| 2:26.16 | הס"ת |  |
| 2:26.17 | אלא |  |
| 2:26.18 | למי |  |
| 2:26.19 | שבחר |  |
| 2:26.20 | בתורה | , |
| 2:26.21 | ובהדיא |  |
| 2:26.22 | אנו |  |
| 2:26.23 | אומרים |  |
| 2:26.24 | בואו |  |
| 2:26.25 | נשתחוה |  |
| 2:26.26 | ונכרעה |  |
| 2:26.27 | נברכה |  |
| 2:26.28 | לפני |  |
| 2:26.29 | ה | ' |
| 2:26.30 | עושנו | . |
| 2:26.31 | והשתחוית |  |
| 2:26.32 | אברהם |  |
| 2:26.33 | אבינו |  |
| 2:26.34 | למלאכים |  |
| 2:26.35 | וכן |  |
| 2:26.36 | יהושע | , |
| 2:26.37 | איבעית |  |
| 2:26.38 | אימא |  |
| 2:26.39 | כיוון |  |
| 2:26.40 | שנדמו |  |
| 2:26.41 | להם |  |
| 2:26.42 | כאנשים |  |
| 2:26.43 | קרינא |  |
| 2:26.44 | בהו |  |
| 2:26.45 | מאי |  |
| 2:26.46 | דאמור |  |
| 2:26.47 | רבנן |  |
| 2:26.48 | בבריתא |  |
| 2:26.49 | לא |  |
| 2:26.50 | תשתחוה |  |
| 2:26.51 | להם | , |
| 2:26.52 | להם |  |
| 2:26.53 | אי |  |
| 2:26.54 | אתה |  |
| 2:26.55 | משתחוה |  |
| 2:26.56 | אבל |  |
| 2:26.57 | אתה |  |
| 2:26.58 | משתחוה |  |
| 2:26.59 | לאדם |  |
| 2:26.60 | שכמותך |  |
| 2:26.61 | כיון |  |
| 2:26.62 | שנגלו |  |
| 2:26.63 | להם |  |
| 2:26.64 | בדמות |  |
| 2:26.65 | אנשים |  |
| 2:26.66 | ואיבעית |  |
| 2:26.67 | אימא |  |
| 2:26.68 | שכוונת |  |
| 2:26.69 | אברהם |  |
| 2:26.70 | ויהושע |  |
| 2:26.71 | להשתחות |  |
| 2:26.72 | למי |  |
| 2:26.73 | ששלחם |  |
| 2:26.74 | וכמו |  |
| 2:26.75 | שכ | ' |
| 2:26.76 | בעל |  |
| 2:26.77 | העיקרים | . |


### Translation under review (2:26)

English: And according to our way, the way of our Rabbis of blessed memory in the aforementioned baraita, there is no need for all this, since our intention is not that we bow down before the Torah scroll, but to Him who chose the Torah. Explicitly, we say: "Come, let us bow down and kneel, let us bless before the LORD our Maker." As for the bowing down of Abraham our father to the angels, and so Joshua's. If you wish, say: since they appeared to them as men, we apply to them what the Rabbis said in the baraita: "'You shall not bow down to them.' To them you may not bow down, but you bow down to a man like yourself," since they were revealed to them in the form of men. And if you wish, say: the intention of Abraham and Joshua was to bow down to Him who sent them, as the author of the Ikkarim wrote.

Literal: according-to our-way the-way-of our-Rabbis-of-blessed-memory in-the-baraita the-aforementioned there-is-no need for-all this since our-intention is-not that-we bow-down before the-Torah-scroll but to-whom chose the-Torah explicitly we say come let-us-bow-down and-kneel let-us-bless before the-LORD our-maker the-bowing-down of-Abraham our-father to-the-angels and-so Joshua if-you-wish say since they-appeared to-them as-men we-apply to-them what they-said the-Rabbis in-the-baraita you-shall-not bow-down to-them to-them you-may-not you bow-down but you bow-down to-a-man like-yourself since they-were-revealed to-them in-the-form of-men and-if-you-wish say the-intention of-Abraham and-Joshua to-bow-down to-whom they-sent-them as the-author-of the-Ikkarim wrote

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:26.1 | and according to |  |  |  |
| 2:26.2 | our way |  |  |  |
| 2:26.3 | the way of |  |  |  |
| 2:26.4 | our Rabbis of blessed memory |  |  | רז״ל: רבותינו זכרונם לברכה |
| 2:26.5 | the baraita |  |  |  |
| 2:26.6 | aforementioned |  |  | הנז״ל: הנזכר |
| 2:26.7 | there is no |  |  |  |
| 2:26.8 | need |  |  |  |
| 2:26.9 | for all |  |  |  |
| 2:26.10 | this |  |  |  |
| 2:26.11 | since not |  |  |  |
| 2:26.12 | our intention |  |  |  |
| 2:26.13 | that we |  |  |  |
| 2:26.14 | bow down |  |  |  |
| 2:26.15 | before |  |  |  |
| 2:26.16 | the Torah scroll |  |  | ס״ת: Sefer Torah, the Torah scroll |
| 2:26.17 | but |  |  |  |
| 2:26.18 | to Him who |  |  | למי does not name him; the context is worship of God |
| 2:26.19 | chose |  |  |  |
| 2:26.20 | the Torah |  |  |  |
| 2:26.21 | explicitly |  |  |  |
| 2:26.22 | we |  |  |  |
| 2:26.23 | say |  |  |  |
| 2:26.24 | come |  |  |  |
| 2:26.25 | let us bow down |  |  |  |
| 2:26.26 | and kneel |  |  |  |
| 2:26.27 | let us bless |  |  |  |
| 2:26.28 | before |  |  |  |
| 2:26.29 | the LORD |  |  | ה׳: the Tetragrammaton (Ps 95:6) |
| 2:26.30 | our Maker |  |  |  |
| 2:26.31 | the bowing down |  |  |  |
| 2:26.32 | of Abraham |  |  |  |
| 2:26.33 | our father |  |  |  |
| 2:26.34 | to the angels |  |  |  |
| 2:26.35 | and so |  |  |  |
| 2:26.36 | Joshua |  |  |  |
| 2:26.37 | if you wish |  |  |  |
| 2:26.38 | say |  |  |  |
| 2:26.39 | since |  |  |  |
| 2:26.40 | they appeared |  |  |  |
| 2:26.41 | to them |  |  |  |
| 2:26.42 | as men |  |  |  |
| 2:26.43 | we apply |  |  |  |
| 2:26.44 | to them |  |  | קרינא בהו: lit. 'we read them', the Talmudic 'we apply the verse to them' |
| 2:26.45 | what |  |  |  |
| 2:26.46 | they said |  |  |  |
| 2:26.47 | the Rabbis |  |  |  |
| 2:26.48 | in the baraita |  |  |  |
| 2:26.49 | you shall not |  |  |  |
| 2:26.50 | bow down |  |  |  |
| 2:26.51 | to them |  |  |  |
| 2:26.52 | to them |  |  |  |
| 2:26.53 | you may not |  |  | אי אתה משתחוה: the Talmudic 'you may not bow down' |
| 2:26.54 | you |  |  |  |
| 2:26.55 | bow down |  |  |  |
| 2:26.56 | but |  |  |  |
| 2:26.57 | you |  |  |  |
| 2:26.58 | bow down |  |  |  |
| 2:26.59 | to a man |  |  |  |
| 2:26.60 | like yourself |  |  |  |
| 2:26.61 | since |  |  |  |
| 2:26.62 | they were revealed |  |  |  |
| 2:26.63 | to them |  |  |  |
| 2:26.64 | in the form |  |  |  |
| 2:26.65 | of men |  |  |  |
| 2:26.66 | and if you wish |  |  |  |
| 2:26.67 | say |  |  |  |
| 2:26.68 | the intention |  |  |  |
| 2:26.69 | of Abraham |  |  |  |
| 2:26.70 | and Joshua |  |  |  |
| 2:26.71 | to bow down |  |  |  |
| 2:26.72 | to Him |  |  | as the note at 'to Him who' above |
| 2:26.73 | who sent them |  |  |  |
| 2:26.74 | as |  |  |  |
| 2:26.75 | wrote |  |  | שכ׳ = שכתב |
| 2:26.76 | the author of |  |  |  |
| 2:26.77 | the Ikkarim |  |  | Sefer ha-Ikkarim, 'Principles'; the author is named only by the title |

Phrases:
- 2:26.1 2:26.2 2:26.3 2:26.4 2:26.5 2:26.6 | And according to our way, the way of our Rabbis of blessed memory in the aforementioned baraita, | according-to our-way the-way-of our-Rabbis in-the-baraita the-aforementioned
- 2:26.7 2:26.8 2:26.9 2:26.10 | there is no need for all this, | there-is-no need for-all this
- 2:26.11 2:26.12 2:26.13 2:26.14 2:26.15 2:26.16 | since our intention is not that we bow down before the Torah scroll, | since our-intention is-not that-we bow-down before the-Torah-scroll
- 2:26.17 2:26.18 2:26.19 2:26.20 | but to Him who chose the Torah. | but to-whom chose the-Torah
- 2:26.21 2:26.22 2:26.23 | Explicitly, we say: | explicitly we say
- 2:26.24 2:26.25 2:26.26 2:26.27 2:26.28 2:26.29 2:26.30 | "Come, let us bow down and kneel, let us bless before the LORD our Maker." | come let-us-bow-down and-kneel let-us-bless before the-LORD our-maker
- 2:26.31 2:26.32 2:26.33 2:26.34 2:26.35 2:26.36 | As for the bowing down of Abraham our father to the angels, and so Joshua's. | the-bowing-down of-Abraham our-father to-the-angels and-so Joshua
- 2:26.37 2:26.38 | If you wish, say: | if-you-wish say
- 2:26.39 2:26.40 2:26.41 2:26.42 2:26.43 2:26.44 | since they appeared to them as men, we apply to them | since they-appeared to-them as-men we-apply to-them
- 2:26.45 2:26.46 2:26.47 2:26.48 | what the Rabbis said in the baraita: | what they-said the-Rabbis in-the-baraita
- 2:26.49 2:26.50 2:26.51 | "'You shall not bow down to them.' | you-shall-not bow-down to-them
- 2:26.52 2:26.53 2:26.54 2:26.55 | To them you may not bow down, | to-them you-may-not you bow-down
- 2:26.56 2:26.57 2:26.58 2:26.59 2:26.60 | but you bow down to a man like yourself." | but you bow-down to-a-man like-yourself
- 2:26.61 2:26.62 2:26.63 2:26.64 2:26.65 | since they were revealed to them in the form of men. | since they-were-revealed to-them in-the-form of-men
- 2:26.66 2:26.67 | And if you wish, say: | and-if-you-wish say
- 2:26.68 2:26.69 2:26.70 2:26.71 2:26.72 2:26.73 | the intention of Abraham and Joshua was to bow down to Him who sent them, | the-intention of-Abraham and-Joshua to-bow-down to-whom they-sent-them
- 2:26.74 2:26.75 2:26.76 2:26.77 | as the author of the Ikkarim wrote. | as the-author-of the-Ikkarim wrote

Grammar notes:
- בואו נשתחוה ונכרעה נברכה is Ps 95:6, given as the people's words: the cohortatives ('let us') stay in the English, and the verse is set in quotation marks.
- והשתחוית is a noun ('the bowing down'), so 'As for the bowing down' is a topic phrase; the English needs no verb until 'and so Joshua's'.
- איבעית אימא ('if you wish, say') introduces a second explanation. The author gives two, the first (the angels appeared as men) and the second (the intention was toward the One who sent them); the English keeps both in order.
- The baraita is quoted inside the sentence ('what the Rabbis said in the baraita: ...'), and the English returns to the author's syntax with 'since they were revealed to them'.
- אי אתה משתחוה is the Talmudic negative of permission, 'you may not bow down'; the English keeps the modal.

Commentary rulings:
- apparatus: Ps 95:6 for 'בואו נשתחוה ונכרעה נברכה לפני ה׳ עושנו'.
- apparatus: Exod 20:5 for 'לא תשתחוה להם' ('you shall not bow down to them'), the words of the baraita.
- apparatus: Cf. Gen 18:2 and Josh 5:14 for the bowing of Abraham and Joshua; the author alludes to them and does not quote them.

## Milchamot ha-Shem 2:27

Hebrew (Wikisource): ושאלתך מענין הנבואה אעפ"י שאינו שייך לעניינינו ולא שאלת רק להטעותנו ולהשיאנו לדבר אחר, הנה זה מבואר מדרגות הנבואה ברמב"ם הל' יסודי התורה, ובחלק שני מספר המורה, ובספר האמונות והדעות להרס"ג. ולפי שמצאתיך חכם בעיניך אחזור לבקש ממך להשיב לי תשובה ניצחת מדברי רז"ל בתלמוד, כי רק אתה לבדך (לפי דעתך) לך נגלו כל תעלומות, ובינה יתירה ורוח חדש ניתן בקרבך. הורינו אני וחבירי איך הדברים יעלו בקנה אחד מלאות וטובות. כי נפשינו חשקה לשמוע לימודיך, ולהשתעשע בתורתך התמימה, ולך תהיה צדקה, וגם כבוד גדול תנחל מאיתנו.

| id | word | punctuation after |
|---|---|---|
| 2:27.1 | ושאלתך |  |
| 2:27.2 | מענין |  |
| 2:27.3 | הנבואה |  |
| 2:27.4 | אעפ"י |  |
| 2:27.5 | שאינו |  |
| 2:27.6 | שייך |  |
| 2:27.7 | לעניינינו |  |
| 2:27.8 | ולא |  |
| 2:27.9 | שאלת |  |
| 2:27.10 | רק |  |
| 2:27.11 | להטעותנו |  |
| 2:27.12 | ולהשיאנו |  |
| 2:27.13 | לדבר |  |
| 2:27.14 | אחר | , |
| 2:27.15 | הנה |  |
| 2:27.16 | זה |  |
| 2:27.17 | מבואר |  |
| 2:27.18 | מדרגות |  |
| 2:27.19 | הנבואה |  |
| 2:27.20 | ברמב"ם |  |
| 2:27.21 | הל | ' |
| 2:27.22 | יסודי |  |
| 2:27.23 | התורה | , |
| 2:27.24 | ובחלק |  |
| 2:27.25 | שני |  |
| 2:27.26 | מספר |  |
| 2:27.27 | המורה | , |
| 2:27.28 | ובספר |  |
| 2:27.29 | האמונות |  |
| 2:27.30 | והדעות |  |
| 2:27.31 | להרס"ג | . |
| 2:27.32 | ולפי |  |
| 2:27.33 | שמצאתיך |  |
| 2:27.34 | חכם |  |
| 2:27.35 | בעיניך |  |
| 2:27.36 | אחזור |  |
| 2:27.37 | לבקש |  |
| 2:27.38 | ממך |  |
| 2:27.39 | להשיב |  |
| 2:27.40 | לי |  |
| 2:27.41 | תשובה |  |
| 2:27.42 | ניצחת |  |
| 2:27.43 | מדברי |  |
| 2:27.44 | רז"ל |  |
| 2:27.45 | בתלמוד | , |
| 2:27.46 | כי |  |
| 2:27.47 | רק |  |
| 2:27.48 | אתה |  |
| 2:27.49 | לבדך |  |
| 2:27.50 | לפי |  |
| 2:27.51 | דעתך | ) |
| 2:27.52 | לך |  |
| 2:27.53 | נגלו |  |
| 2:27.54 | כל |  |
| 2:27.55 | תעלומות | , |
| 2:27.56 | ובינה |  |
| 2:27.57 | יתירה |  |
| 2:27.58 | ורוח |  |
| 2:27.59 | חדש |  |
| 2:27.60 | ניתן |  |
| 2:27.61 | בקרבך | . |
| 2:27.62 | הורינו |  |
| 2:27.63 | אני |  |
| 2:27.64 | וחבירי |  |
| 2:27.65 | איך |  |
| 2:27.66 | הדברים |  |
| 2:27.67 | יעלו |  |
| 2:27.68 | בקנה |  |
| 2:27.69 | אחד |  |
| 2:27.70 | מלאות |  |
| 2:27.71 | וטובות | . |
| 2:27.72 | כי |  |
| 2:27.73 | נפשינו |  |
| 2:27.74 | חשקה |  |
| 2:27.75 | לשמוע |  |
| 2:27.76 | לימודיך | , |
| 2:27.77 | ולהשתעשע |  |
| 2:27.78 | בתורתך |  |
| 2:27.79 | התמימה | , |
| 2:27.80 | ולך |  |
| 2:27.81 | תהיה |  |
| 2:27.82 | צדקה | , |
| 2:27.83 | וגם |  |
| 2:27.84 | כבוד |  |
| 2:27.85 | גדול |  |
| 2:27.86 | תנחל |  |
| 2:27.87 | מאיתנו | . |


### Translation under review (2:27)

English: And your question concerning prophecy, although it does not belong to our subject, and you did not ask only to mislead us and to divert us to another matter. Behold, this is set out in the degrees of prophecy in the Rambam, Laws of the Foundations of the Torah, and in part two of the book Moreh Nevukhim, and in the book Emunot ve-De'ot of Rabbi Saadia Gaon. Since I have found you wise in your eyes, I return to ask you to answer me a decisive answer from the words of our Rabbis of blessed memory in the Talmud, for you alone (in your opinion) all mysteries have been revealed to you, and greater understanding, and a new spirit has been given within you. Instruct us, I and my companions, how the matters can agree fully and well. For our soul has longed to hear your teachings, and to delight in your perfect Torah, and it will be to you a merit, and also great honor you will inherit from us.

Literal: and-your-question concerning prophecy although is-not relevant to-our-subject and-not you-asked only to-mislead-us and-to-divert-us to-a-matter other. behold this is-set-out the-degrees-of prophecy in-the-Rambam Laws-of the-Foundations-of the-Torah, and-in part two of-the-book the-Moreh, and-in-the-book the-Emunot and-the-De'ot of-Rabbi-Saadia-Gaon. and-since I-found-you wise in-your-eyes, I-return to-ask of-you to-answer to-me an-answer decisive from-the-words-of our-Rabbis in-the-Talmud, for you alone (in-your-opinion) to-you were-revealed all mysteries, and-understanding greater and-spirit new has-been-given within-you. instruct-us I and-my-companions how the-matters will-agree in-one-reed in-full and-good. for our-soul longed to-hear your-teachings, and-to-delight in-your-Torah the-perfect, and-to-you will-be righteousness, and-also honor great you-will-inherit from-us

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:27.1 | and your question |  |  |  |
| 2:27.2 | concerning |  |  |  |
| 2:27.3 | prophecy |  |  |  |
| 2:27.4 | although |  |  | אעפ״י = אף על פי |
| 2:27.5 | is not |  |  |  |
| 2:27.6 | relevant |  |  |  |
| 2:27.7 | to our subject |  |  |  |
| 2:27.8 | and not |  |  |  |
| 2:27.9 | you asked |  |  |  |
| 2:27.10 | only |  |  |  |
| 2:27.11 | to mislead us |  |  |  |
| 2:27.12 | and to divert us |  |  |  |
| 2:27.13 | to a matter |  |  |  |
| 2:27.14 | other |  |  |  |
| 2:27.15 | behold |  |  |  |
| 2:27.16 | this |  |  |  |
| 2:27.17 | is set out |  |  |  |
| 2:27.18 | the degrees of |  |  |  |
| 2:27.19 | prophecy |  |  |  |
| 2:27.20 | in the Rambam |  |  |  |
| 2:27.21 | Laws of |  |  | הל׳ = הלכות: the Rambam's Mishneh Torah (the code); its first book is 'Laws of the Foundations of the Torah' |
| 2:27.22 | of the Foundations |  |  |  |
| 2:27.23 | of the Torah |  |  |  |
| 2:27.24 | and in part |  |  |  |
| 2:27.25 | two |  |  |  |
| 2:27.26 | of the book |  |  |  |
| 2:27.27 | Moreh Nevukhim |  |  | Moreh Nevukhim, the Guide of the Perplexed; the Hebrew is 'the book of the Moreh' |
| 2:27.28 | and in the book |  |  |  |
| 2:27.29 | Emunot |  |  | Emunot ve-De'ot, 'Beliefs and Opinions' |
| 2:27.30 | and De'ot |  |  |  |
| 2:27.31 | of Rabbi Saadia Gaon |  |  | רס״ג: Rabbi Saadia Gaon |
| 2:27.32 | and since |  |  |  |
| 2:27.33 | I have found you |  |  |  |
| 2:27.34 | wise |  |  |  |
| 2:27.35 | in your eyes |  |  | the idiom 'wise in one's own eyes' (Prov 3:7); the Hebrew keeps both readings |
| 2:27.36 | I return |  |  |  |
| 2:27.37 | to ask |  |  |  |
| 2:27.38 | of you |  |  |  |
| 2:27.39 | to answer |  |  |  |
| 2:27.40 | me |  |  |  |
| 2:27.41 | an answer |  |  |  |
| 2:27.42 | decisive |  |  |  |
| 2:27.43 | from the words of |  |  |  |
| 2:27.44 | our Rabbis of blessed memory |  |  | רז״ל: רבותינו זכרונם לברכה |
| 2:27.45 | in the Talmud |  |  |  |
| 2:27.46 | for |  |  |  |
| 2:27.47 | only |  |  |  |
| 2:27.48 | you |  |  |  |
| 2:27.49 | alone |  |  |  |
| 2:27.50 | in |  |  |  |
| 2:27.51 | your opinion |  |  | the parentheses are the author's and are kept |
| 2:27.52 | to you |  |  |  |
| 2:27.53 | were revealed |  |  |  |
| 2:27.54 | all |  |  |  |
| 2:27.55 | mysteries |  |  |  |
| 2:27.56 | and understanding |  |  |  |
| 2:27.57 | greater |  |  |  |
| 2:27.58 | and a spirit |  |  |  |
| 2:27.59 | new |  |  |  |
| 2:27.60 | has been given |  |  |  |
| 2:27.61 | within you |  |  |  |
| 2:27.62 | instruct us |  |  | imperative, with 'us' as object |
| 2:27.63 | I |  |  |  |
| 2:27.64 | and my companions |  |  |  |
| 2:27.65 | how |  |  |  |
| 2:27.66 | the matters |  |  |  |
| 2:27.67 | can agree |  |  | יעלו בקנה אחד: the idiom 'to agree' |
| 2:27.68 | in agreement |  |  | lit. 'in one reed' |
| 2:27.69 | one |  |  |  |
| 2:27.70 | fully |  |  |  |
| 2:27.71 | and well |  |  |  |
| 2:27.72 | for |  |  |  |
| 2:27.73 | our soul |  |  |  |
| 2:27.74 | has longed |  |  |  |
| 2:27.75 | to hear |  |  |  |
| 2:27.76 | your teachings |  |  |  |
| 2:27.77 | and to delight |  |  |  |
| 2:27.78 | in your Torah |  |  |  |
| 2:27.79 | perfect |  |  |  |
| 2:27.80 | and to you |  |  |  |
| 2:27.81 | will be |  |  |  |
| 2:27.82 | a merit |  |  | צדקה: 'righteousness', here the merit that the act earns |
| 2:27.83 | and also |  |  |  |
| 2:27.84 | honor |  |  |  |
| 2:27.85 | great |  |  |  |
| 2:27.86 | you will inherit |  |  |  |
| 2:27.87 | from us |  |  |  |

Phrases:
- 2:27.1 2:27.2 2:27.3 | And your question concerning prophecy, | and-your-question concerning prophecy
- 2:27.4 2:27.5 2:27.6 2:27.7 | although it does not belong to our subject, | although is-not relevant to-our-subject
- 2:27.8 2:27.9 2:27.10 2:27.11 2:27.12 2:27.13 2:27.14 | and you did not ask only to mislead us and to divert us to another matter. | and-not you-asked only to-mislead-us and-to-divert-us to-a-matter other
- 2:27.15 2:27.16 2:27.17 2:27.18 2:27.19 | Behold, this is set out in the degrees of prophecy | behold this is-set-out the-degrees-of prophecy
- 2:27.20 2:27.21 2:27.22 2:27.23 | in the Rambam, Laws of the Foundations of the Torah, | in-the-Rambam Laws-of the-Foundations-of the-Torah
- 2:27.24 2:27.25 2:27.26 2:27.27 | and in part two of the book Moreh Nevukhim, | and-in part two of-the-book the-Moreh
- 2:27.28 2:27.29 2:27.30 2:27.31 | and in the book Emunot ve-De'ot of Rabbi Saadia Gaon. | and-in-the-book the-Emunot and-the-De'ot of-Rabbi-Saadia-Gaon
- 2:27.32 2:27.33 2:27.34 2:27.35 | Since I have found you wise in your eyes, | and-since I-found-you wise in-your-eyes
- 2:27.36 2:27.37 2:27.38 2:27.39 2:27.40 | I return to ask you to answer me | I return to-ask of-you to-answer me
- 2:27.41 2:27.42 2:27.43 2:27.44 2:27.45 | a decisive answer from the words of our Rabbis of blessed memory in the Talmud, | an-answer decisive from-the-words-of our-Rabbis in-the-Talmud
- 2:27.46 2:27.47 2:27.48 2:27.49 2:27.50 2:27.51 | for you alone (in your opinion) | for only you alone (in-your-opinion)
- 2:27.52 2:27.53 2:27.54 2:27.55 | all mysteries have been revealed to you, | to-you were-revealed all mysteries
- 2:27.56 2:27.57 2:27.58 2:27.59 2:27.60 2:27.61 | and greater understanding, and a new spirit has been given within you. | and-understanding greater and-spirit new has-been-given within-you
- 2:27.62 2:27.63 2:27.64 2:27.65 2:27.66 2:27.67 2:27.68 2:27.69 2:27.70 2:27.71 | Instruct us, I and my companions, how the matters can agree fully and well. | instruct-us I and-my-companions how the-matters will-agree in-one-reed in-full and-good
- 2:27.72 2:27.73 2:27.74 2:27.75 2:27.76 | For our soul has longed to hear your teachings, | for our-soul has-longed to-hear your-teachings
- 2:27.77 2:27.78 2:27.79 | and to delight in your perfect Torah, | and-to-delight in-your-Torah the-perfect
- 2:27.80 2:27.81 2:27.82 | and it will be to you a merit, | and-to-you will-be righteousness
- 2:27.83 2:27.84 2:27.85 2:27.86 2:27.87 | and also great honor you will inherit from us. | and-also honor great you-will-inherit from-us

Grammar notes:
- The paragraph is a single letter sentence, and the list of works (the Rambam, the Moreh, Emunot ve-De'ot) is a series of 'in' phrases; the English keeps the series in the Hebrew's order.
- ולא שאלת רק להטעותנו: 'you did not ask except to mislead us'; the English keeps the restriction with 'only'.
- חכם בעיניך is an idiom that can mean 'wise in your own eyes' (Prov 3:7). The English keeps 'in your eyes', as the Hebrew does.
- הורינו אני וחבירי איך הדברים יעלו: the imperative is followed by a 'how' clause, so the English reads 'Instruct us ..., how the matters can agree'.

Commentary rulings:
- apparatus: Cf. Ezek 36:26 for 'ורוח חדש ניתן בקרבך' ('a new spirit ... within you'); the author does not quote it.
- apparatus: Cf. Prov 3:7 for 'חכם בעיניך' (see the token note at 'in your eyes').

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

