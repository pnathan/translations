# Translation pack: milhamot-hashem perek-3-04 (Milchamot ha-Shem 3:7–3:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-04`.

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
- **ב"ה** → "blessed be He".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- כופר **בעיקר** → "denies the fundamental principle".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **תעמוד / רעך** (Lev 19:16) → "stand by the blood of your fellow" (no
  "idly").
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


## Milchamot ha-Shem 3:7

Hebrew (Wikisource): ז) וז"ל המקובל בעל היכל הברכה, מצות עשה על הב"ד לדרוש ולחקור את העדים בשבע חקירות וכו' עוד יש בדיקות שהוא שאלות מה היה לבוש, עפר הארץ שהיה עומד עליה שחור או חול, או חמר, ואם כחשו זה לזה בטל העדות. משרשי המצוה הזאת הכלה הכלולה עולימתא שפירתא דלית בה עינין, הנהגת העולם הזה שיהא בחירה חפשית עיניו כיונים אונאה לבני נשא. ועינין דילה בהיכל זכות, עיני העדה, שדנין בידיעה על פי שנים עדים, תרין עטרין, והדעת הוא סוד עינים כנודע, ואור זה מתפשט בשבע בחינות וכ' ע"ש. ובמצות לא תלך רכיל כתב וז"ל, שלא להגיד לחברו, שכנגדו דבר עליו רע, שנאמר, לא תלך רכיל. ומדיני המצוה אף שמדבר אמת, ויש עון גדול בזה. וגם בכלל לאו זה שלא ידבר כגנות חברו אף אמת והוא כאילו כפר בעיקר וכו'. והאריך עוד בעונש המספר לשון הרע ובשכר הנשמר ממנו, יעו"ש.

| id | word | punctuation after |
|---|---|---|
| 3:7.1 | ז | ) |
| 3:7.2 | וז"ל |  |
| 3:7.3 | המקובל |  |
| 3:7.4 | בעל |  |
| 3:7.5 | היכל |  |
| 3:7.6 | הברכה | , |
| 3:7.7 | מצות |  |
| 3:7.8 | עשה |  |
| 3:7.9 | על |  |
| 3:7.10 | הב"ד |  |
| 3:7.11 | לדרוש |  |
| 3:7.12 | ולחקור |  |
| 3:7.13 | את |  |
| 3:7.14 | העדים |  |
| 3:7.15 | בשבע |  |
| 3:7.16 | חקירות |  |
| 3:7.17 | וכו | ' |
| 3:7.18 | עוד |  |
| 3:7.19 | יש |  |
| 3:7.20 | בדיקות |  |
| 3:7.21 | שהוא |  |
| 3:7.22 | שאלות |  |
| 3:7.23 | מה |  |
| 3:7.24 | היה |  |
| 3:7.25 | לבוש | , |
| 3:7.26 | עפר |  |
| 3:7.27 | הארץ |  |
| 3:7.28 | שהיה |  |
| 3:7.29 | עומד |  |
| 3:7.30 | עליה |  |
| 3:7.31 | שחור |  |
| 3:7.32 | או |  |
| 3:7.33 | חול | , |
| 3:7.34 | או |  |
| 3:7.35 | חמר | , |
| 3:7.36 | ואם |  |
| 3:7.37 | כחשו |  |
| 3:7.38 | זה |  |
| 3:7.39 | לזה |  |
| 3:7.40 | בטל |  |
| 3:7.41 | העדות | . |
| 3:7.42 | משרשי |  |
| 3:7.43 | המצוה |  |
| 3:7.44 | הזאת |  |
| 3:7.45 | הכלה |  |
| 3:7.46 | הכלולה |  |
| 3:7.47 | עולימתא |  |
| 3:7.48 | שפירתא |  |
| 3:7.49 | דלית |  |
| 3:7.50 | בה |  |
| 3:7.51 | עינין | , |
| 3:7.52 | הנהגת |  |
| 3:7.53 | העולם |  |
| 3:7.54 | הזה |  |
| 3:7.55 | שיהא |  |
| 3:7.56 | בחירה |  |
| 3:7.57 | חפשית |  |
| 3:7.58 | עיניו |  |
| 3:7.59 | כיונים |  |
| 3:7.60 | אונאה |  |
| 3:7.61 | לבני |  |
| 3:7.62 | נשא | . |
| 3:7.63 | ועינין |  |
| 3:7.64 | דילה |  |
| 3:7.65 | בהיכל |  |
| 3:7.66 | זכות | , |
| 3:7.67 | עיני |  |
| 3:7.68 | העדה | , |
| 3:7.69 | שדנין |  |
| 3:7.70 | בידיעה |  |
| 3:7.71 | על |  |
| 3:7.72 | פי |  |
| 3:7.73 | שנים |  |
| 3:7.74 | עדים | , |
| 3:7.75 | תרין |  |
| 3:7.76 | עטרין | , |
| 3:7.77 | והדעת |  |
| 3:7.78 | הוא |  |
| 3:7.79 | סוד |  |
| 3:7.80 | עינים |  |
| 3:7.81 | כנודע | , |
| 3:7.82 | ואור |  |
| 3:7.83 | זה |  |
| 3:7.84 | מתפשט |  |
| 3:7.85 | בשבע |  |
| 3:7.86 | בחינות |  |
| 3:7.87 | וכ | ' |
| 3:7.88 | ע"ש | . |
| 3:7.89 | ובמצות |  |
| 3:7.90 | לא |  |
| 3:7.91 | תלך |  |
| 3:7.92 | רכיל |  |
| 3:7.93 | כתב |  |
| 3:7.94 | וז"ל | , |
| 3:7.95 | שלא |  |
| 3:7.96 | להגיד |  |
| 3:7.97 | לחברו | , |
| 3:7.98 | שכנגדו |  |
| 3:7.99 | דבר |  |
| 3:7.100 | עליו |  |
| 3:7.101 | רע | , |
| 3:7.102 | שנאמר | , |
| 3:7.103 | לא |  |
| 3:7.104 | תלך |  |
| 3:7.105 | רכיל | . |
| 3:7.106 | ומדיני |  |
| 3:7.107 | המצוה |  |
| 3:7.108 | אף |  |
| 3:7.109 | שמדבר |  |
| 3:7.110 | אמת | , |
| 3:7.111 | ויש |  |
| 3:7.112 | עון |  |
| 3:7.113 | גדול |  |
| 3:7.114 | בזה | . |
| 3:7.115 | וגם |  |
| 3:7.116 | בכלל |  |
| 3:7.117 | לאו |  |
| 3:7.118 | זה |  |
| 3:7.119 | שלא |  |
| 3:7.120 | ידבר |  |
| 3:7.121 | כגנות |  |
| 3:7.122 | חברו |  |
| 3:7.123 | אף |  |
| 3:7.124 | אמת |  |
| 3:7.125 | והוא |  |
| 3:7.126 | כאילו |  |
| 3:7.127 | כפר |  |
| 3:7.128 | בעיקר |  |
| 3:7.129 | וכו | '. |
| 3:7.130 | והאריך |  |
| 3:7.131 | עוד |  |
| 3:7.132 | בעונש |  |
| 3:7.133 | המספר |  |
| 3:7.134 | לשון |  |
| 3:7.135 | הרע |  |
| 3:7.136 | ובשכר |  |
| 3:7.137 | הנשמר |  |
| 3:7.138 | ממנו | , |
| 3:7.139 | יעו"ש | . |


## Milchamot ha-Shem 3:8

Hebrew (Wikisource): ח) ובאזהרת לא תעמוד על דם רעך, כ' אפי' עמידה כל דהו. לא תראה בצערו של חברך שאחד מלבין פניו ואתה שותק, עמדו ולא ענו עוד, אלא תהיה מוחה בכל יכלתך במי שמבזה חברך, או מדבר עליו רכילות ולשון הרע, אתה לא תהיה שומע ושותק, אלא תאמר לו רשע עוכר ישראל, למה תדבר סרה על דמות יוצרנו ועל שכינת עוזינו שהיא היו"ד שבעמיך, ומסתלקת בעון זה. ולכך נסמכה זה לזה, לא תלך רכיל ותאמר כיון שאיני מדבר אשמע ואשתוק ונהנה במה שחברו רשע מדבר, גם זה השומע רשע גמור. ועובר על לא תעמוד על דם רעך. ואין לו חלק באלהי ישראל כששומע רשע מדבר סרה על צדיק, או לוחש מלחש בשקרים על חברו רק תכף תדחה אותו ותאמר לו רשע לך מעלי, איני שומע לך, כי אם הוא שומע ושותק, וכל שכן אם נהנה, סופו נידון באש לא נופח, ולא יזכה לעמוד במחיצת הצדיקים, כשומע, כעונה, כמדבר, כולם יספו מן העולם בזה ובבא עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:8.1 | ח | ) |
| 3:8.2 | ובאזהרת |  |
| 3:8.3 | לא |  |
| 3:8.4 | תעמוד |  |
| 3:8.5 | על |  |
| 3:8.6 | דם |  |
| 3:8.7 | רעך | , |
| 3:8.8 | כ | ' |
| 3:8.9 | אפי | ' |
| 3:8.10 | עמידה |  |
| 3:8.11 | כל |  |
| 3:8.12 | דהו | . |
| 3:8.13 | לא |  |
| 3:8.14 | תראה |  |
| 3:8.15 | בצערו |  |
| 3:8.16 | של |  |
| 3:8.17 | חברך |  |
| 3:8.18 | שאחד |  |
| 3:8.19 | מלבין |  |
| 3:8.20 | פניו |  |
| 3:8.21 | ואתה |  |
| 3:8.22 | שותק | , |
| 3:8.23 | עמדו |  |
| 3:8.24 | ולא |  |
| 3:8.25 | ענו |  |
| 3:8.26 | עוד | , |
| 3:8.27 | אלא |  |
| 3:8.28 | תהיה |  |
| 3:8.29 | מוחה |  |
| 3:8.30 | בכל |  |
| 3:8.31 | יכלתך |  |
| 3:8.32 | במי |  |
| 3:8.33 | שמבזה |  |
| 3:8.34 | חברך | , |
| 3:8.35 | או |  |
| 3:8.36 | מדבר |  |
| 3:8.37 | עליו |  |
| 3:8.38 | רכילות |  |
| 3:8.39 | ולשון |  |
| 3:8.40 | הרע | , |
| 3:8.41 | אתה |  |
| 3:8.42 | לא |  |
| 3:8.43 | תהיה |  |
| 3:8.44 | שומע |  |
| 3:8.45 | ושותק | , |
| 3:8.46 | אלא |  |
| 3:8.47 | תאמר |  |
| 3:8.48 | לו |  |
| 3:8.49 | רשע |  |
| 3:8.50 | עוכר |  |
| 3:8.51 | ישראל | , |
| 3:8.52 | למה |  |
| 3:8.53 | תדבר |  |
| 3:8.54 | סרה |  |
| 3:8.55 | על |  |
| 3:8.56 | דמות |  |
| 3:8.57 | יוצרנו |  |
| 3:8.58 | ועל |  |
| 3:8.59 | שכינת |  |
| 3:8.60 | עוזינו |  |
| 3:8.61 | שהיא |  |
| 3:8.62 | היו"ד |  |
| 3:8.63 | שבעמיך | , |
| 3:8.64 | ומסתלקת |  |
| 3:8.65 | בעון |  |
| 3:8.66 | זה | . |
| 3:8.67 | ולכך |  |
| 3:8.68 | נסמכה |  |
| 3:8.69 | זה |  |
| 3:8.70 | לזה | , |
| 3:8.71 | לא |  |
| 3:8.72 | תלך |  |
| 3:8.73 | רכיל |  |
| 3:8.74 | ותאמר |  |
| 3:8.75 | כיון |  |
| 3:8.76 | שאיני |  |
| 3:8.77 | מדבר |  |
| 3:8.78 | אשמע |  |
| 3:8.79 | ואשתוק |  |
| 3:8.80 | ונהנה |  |
| 3:8.81 | במה |  |
| 3:8.82 | שחברו |  |
| 3:8.83 | רשע |  |
| 3:8.84 | מדבר | , |
| 3:8.85 | גם |  |
| 3:8.86 | זה |  |
| 3:8.87 | השומע |  |
| 3:8.88 | רשע |  |
| 3:8.89 | גמור | . |
| 3:8.90 | ועובר |  |
| 3:8.91 | על |  |
| 3:8.92 | לא |  |
| 3:8.93 | תעמוד |  |
| 3:8.94 | על |  |
| 3:8.95 | דם |  |
| 3:8.96 | רעך | . |
| 3:8.97 | ואין |  |
| 3:8.98 | לו |  |
| 3:8.99 | חלק |  |
| 3:8.100 | באלהי |  |
| 3:8.101 | ישראל |  |
| 3:8.102 | כששומע |  |
| 3:8.103 | רשע |  |
| 3:8.104 | מדבר |  |
| 3:8.105 | סרה |  |
| 3:8.106 | על |  |
| 3:8.107 | צדיק | , |
| 3:8.108 | או |  |
| 3:8.109 | לוחש |  |
| 3:8.110 | מלחש |  |
| 3:8.111 | בשקרים |  |
| 3:8.112 | על |  |
| 3:8.113 | חברו |  |
| 3:8.114 | רק |  |
| 3:8.115 | תכף |  |
| 3:8.116 | תדחה |  |
| 3:8.117 | אותו |  |
| 3:8.118 | ותאמר |  |
| 3:8.119 | לו |  |
| 3:8.120 | רשע |  |
| 3:8.121 | לך |  |
| 3:8.122 | מעלי | , |
| 3:8.123 | איני |  |
| 3:8.124 | שומע |  |
| 3:8.125 | לך | , |
| 3:8.126 | כי |  |
| 3:8.127 | אם |  |
| 3:8.128 | הוא |  |
| 3:8.129 | שומע |  |
| 3:8.130 | ושותק | , |
| 3:8.131 | וכל |  |
| 3:8.132 | שכן |  |
| 3:8.133 | אם |  |
| 3:8.134 | נהנה | , |
| 3:8.135 | סופו |  |
| 3:8.136 | נידון |  |
| 3:8.137 | באש |  |
| 3:8.138 | לא |  |
| 3:8.139 | נופח | , |
| 3:8.140 | ולא |  |
| 3:8.141 | יזכה |  |
| 3:8.142 | לעמוד |  |
| 3:8.143 | במחיצת |  |
| 3:8.144 | הצדיקים | , |
| 3:8.145 | כשומע | , |
| 3:8.146 | כעונה | , |
| 3:8.147 | כמדבר | , |
| 3:8.148 | כולם |  |
| 3:8.149 | יספו |  |
| 3:8.150 | מן |  |
| 3:8.151 | העולם |  |
| 3:8.152 | בזה |  |
| 3:8.153 | ובבא |  |
| 3:8.154 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

