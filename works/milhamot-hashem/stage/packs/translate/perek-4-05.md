# Translation pack: milhamot-hashem perek-4-05 (Milchamot ha-Shem 4:8–4:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-05`.

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
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **כוונה** (in prayer) → "intention".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- אחד ... **ואחד** of two kinds of case → "alike".
- **סכלות** והוללות → "folly and madness".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ע"א / ע"ב** → "a / b".
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
- **סנהדרין** → "Sanhedrin 99"; סנהדרין רי"ף → "Sanhedrin, [in] the Rif".
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


## Milchamot ha-Shem 4:8

Hebrew (Wikisource): כה) ולפי דברי הזהר ומקדש מלך למה שינו השבעים ושנים זקנים לתלמי המלך כשהעתיקו לו את התורה כדגרסינן בפרק קמא דמגלה מעשה בתלמי המלך שכינס ע"ב זקינים והושיבם בע"ב בתים ולא גילה להם על מה כינסם. ונכנס אצל כל אחד ואחד מהם. ואמר להם כתבו לי תורת משה רבכם. ונתן הקב"ה עצה בלב כל אחד ואחד והסכימה דעתם לדעת אחת, וכתבו לו אלהים ברא בראשית את השמים ואת הארץ ופי' רש"י ז"ל שלא יאמר בראשית שם הוא, ושתי רשויות הן, האחד ברא את השני, וכן כתבו אעשה אדם בצלם ובדמות. וכן כ' הרב מהרש"א בפי' רש"י שלא יאמר שתי רשויות הן והוא לא יקבל מה שאמרו בסנהדרין כל מקום שפקרו המינים תשובתן בצדן. נעשה אדם, ויברא אלהים, הבה נרדה, וירד ה', ונכתבו מעיקרא בלשון רבים דאמר ר' יוחנן אין הקב"ה עושה דבר עד שנמלך בפימליא שלמעלה וכו' וענינו כמו שכ' רש"י בחומש שהוא ממדת הענוה שיהא הגדול נמלך בקטן והוא על פי בראשית רבה וכו'.

| id | word | punctuation after |
|---|---|---|
| 4:8.1 | כה | ) |
| 4:8.2 | ולפי |  |
| 4:8.3 | דברי |  |
| 4:8.4 | הזהר |  |
| 4:8.5 | ומקדש |  |
| 4:8.6 | מלך |  |
| 4:8.7 | למה |  |
| 4:8.8 | שינו |  |
| 4:8.9 | השבעים |  |
| 4:8.10 | ושנים |  |
| 4:8.11 | זקנים |  |
| 4:8.12 | לתלמי |  |
| 4:8.13 | המלך |  |
| 4:8.14 | כשהעתיקו |  |
| 4:8.15 | לו |  |
| 4:8.16 | את |  |
| 4:8.17 | התורה |  |
| 4:8.18 | כדגרסינן |  |
| 4:8.19 | בפרק |  |
| 4:8.20 | קמא |  |
| 4:8.21 | דמגלה |  |
| 4:8.22 | מעשה |  |
| 4:8.23 | בתלמי |  |
| 4:8.24 | המלך |  |
| 4:8.25 | שכינס |  |
| 4:8.26 | ע"ב |  |
| 4:8.27 | זקינים |  |
| 4:8.28 | והושיבם |  |
| 4:8.29 | בע"ב |  |
| 4:8.30 | בתים |  |
| 4:8.31 | ולא |  |
| 4:8.32 | גילה |  |
| 4:8.33 | להם |  |
| 4:8.34 | על |  |
| 4:8.35 | מה |  |
| 4:8.36 | כינסם | . |
| 4:8.37 | ונכנס |  |
| 4:8.38 | אצל |  |
| 4:8.39 | כל |  |
| 4:8.40 | אחד |  |
| 4:8.41 | ואחד |  |
| 4:8.42 | מהם | . |
| 4:8.43 | ואמר |  |
| 4:8.44 | להם |  |
| 4:8.45 | כתבו |  |
| 4:8.46 | לי |  |
| 4:8.47 | תורת |  |
| 4:8.48 | משה |  |
| 4:8.49 | רבכם | . |
| 4:8.50 | ונתן |  |
| 4:8.51 | הקב"ה |  |
| 4:8.52 | עצה |  |
| 4:8.53 | בלב |  |
| 4:8.54 | כל |  |
| 4:8.55 | אחד |  |
| 4:8.56 | ואחד |  |
| 4:8.57 | והסכימה |  |
| 4:8.58 | דעתם |  |
| 4:8.59 | לדעת |  |
| 4:8.60 | אחת | , |
| 4:8.61 | וכתבו |  |
| 4:8.62 | לו |  |
| 4:8.63 | אלהים |  |
| 4:8.64 | ברא |  |
| 4:8.65 | בראשית |  |
| 4:8.66 | את |  |
| 4:8.67 | השמים |  |
| 4:8.68 | ואת |  |
| 4:8.69 | הארץ |  |
| 4:8.70 | ופי | ' |
| 4:8.71 | רש"י |  |
| 4:8.72 | ז"ל |  |
| 4:8.73 | שלא |  |
| 4:8.74 | יאמר |  |
| 4:8.75 | בראשית |  |
| 4:8.76 | שם |  |
| 4:8.77 | הוא | , |
| 4:8.78 | ושתי |  |
| 4:8.79 | רשויות |  |
| 4:8.80 | הן | , |
| 4:8.81 | האחד |  |
| 4:8.82 | ברא |  |
| 4:8.83 | את |  |
| 4:8.84 | השני | , |
| 4:8.85 | וכן |  |
| 4:8.86 | כתבו |  |
| 4:8.87 | אעשה |  |
| 4:8.88 | אדם |  |
| 4:8.89 | בצלם |  |
| 4:8.90 | ובדמות | . |
| 4:8.91 | וכן |  |
| 4:8.92 | כ | ' |
| 4:8.93 | הרב |  |
| 4:8.94 | מהרש"א |  |
| 4:8.95 | בפי | ' |
| 4:8.96 | רש"י |  |
| 4:8.97 | שלא |  |
| 4:8.98 | יאמר |  |
| 4:8.99 | שתי |  |
| 4:8.100 | רשויות |  |
| 4:8.101 | הן |  |
| 4:8.102 | והוא |  |
| 4:8.103 | לא |  |
| 4:8.104 | יקבל |  |
| 4:8.105 | מה |  |
| 4:8.106 | שאמרו |  |
| 4:8.107 | בסנהדרין |  |
| 4:8.108 | כל |  |
| 4:8.109 | מקום |  |
| 4:8.110 | שפקרו |  |
| 4:8.111 | המינים |  |
| 4:8.112 | תשובתן |  |
| 4:8.113 | בצדן | . |
| 4:8.114 | נעשה |  |
| 4:8.115 | אדם | , |
| 4:8.116 | ויברא |  |
| 4:8.117 | אלהים | , |
| 4:8.118 | הבה |  |
| 4:8.119 | נרדה | , |
| 4:8.120 | וירד |  |
| 4:8.121 | ה | ', |
| 4:8.122 | ונכתבו |  |
| 4:8.123 | מעיקרא |  |
| 4:8.124 | בלשון |  |
| 4:8.125 | רבים |  |
| 4:8.126 | דאמר |  |
| 4:8.127 | ר | ' |
| 4:8.128 | יוחנן |  |
| 4:8.129 | אין |  |
| 4:8.130 | הקב"ה |  |
| 4:8.131 | עושה |  |
| 4:8.132 | דבר |  |
| 4:8.133 | עד |  |
| 4:8.134 | שנמלך |  |
| 4:8.135 | בפימליא |  |
| 4:8.136 | שלמעלה |  |
| 4:8.137 | וכו | ' |
| 4:8.138 | וענינו |  |
| 4:8.139 | כמו |  |
| 4:8.140 | שכ | ' |
| 4:8.141 | רש"י |  |
| 4:8.142 | בחומש |  |
| 4:8.143 | שהוא |  |
| 4:8.144 | ממדת |  |
| 4:8.145 | הענוה |  |
| 4:8.146 | שיהא |  |
| 4:8.147 | הגדול |  |
| 4:8.148 | נמלך |  |
| 4:8.149 | בקטן |  |
| 4:8.150 | והוא |  |
| 4:8.151 | על |  |
| 4:8.152 | פי |  |
| 4:8.153 | בראשית |  |
| 4:8.154 | רבה |  |
| 4:8.155 | וכו | '. |


## Milchamot ha-Shem 4:9

Hebrew (Wikisource): ועיין במורה נבוכים פ"ו משני שהביא שתי נוסחאות. האחת שאין הקב"ה עושה דבר עד שמסתכל בפימליא שלמעלה והוה כמאמר אפלטון שהשי"ת יעיין בעולם השכלים וישפיע ממנו המציאות. והנוסחא השנית היא שאמרו במקום אחר אין הקב"ה עושה דבר עד שנמלך בפימיליא שלמעלה. ופימליא הוא החיל בלשון יון. ובבראשית רבה ובמדרש קהלת את אשר כבר עשוהו. עשהו לא נאמר אלא עשוהו. כביכול הוא ובית דינו נמנו על כל אבר ואבר שבך והושיבוהו על כנו שנ' הוא עשך ויכוננך, ובבראשית רבה אמרו ג"כ כל מקום שנ' וה' הוא ובית דנו, ואין הכוונה כאלו המאמרים מה שהשבוהו הפתאים שיש לו ית' דברים או מחשבה או שאלת עצה או התבוננות, והעזר בדעת אחרים, ואיך יעזר הבורא במה שברא, אבל זה כולו עניינו שאפי' חלקי המציאות עד בריאת האברים מבעלי חיים כפי מה שהם עליו. כל זה באמצעית מלאכים, כי הכחות כולם מלאכים. ומה מאד רע עורון הסכלות, ומה מאד מזיק, שאלו תאמר לאחד מאשר יחשבו שהם חכמי ישראל וכו' ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:9.1 | ועיין |  |
| 4:9.2 | במורה |  |
| 4:9.3 | נבוכים |  |
| 4:9.4 | פ"ו |  |
| 4:9.5 | משני |  |
| 4:9.6 | שהביא |  |
| 4:9.7 | שתי |  |
| 4:9.8 | נוסחאות | . |
| 4:9.9 | האחת |  |
| 4:9.10 | שאין |  |
| 4:9.11 | הקב"ה |  |
| 4:9.12 | עושה |  |
| 4:9.13 | דבר |  |
| 4:9.14 | עד |  |
| 4:9.15 | שמסתכל |  |
| 4:9.16 | בפימליא |  |
| 4:9.17 | שלמעלה |  |
| 4:9.18 | והוה |  |
| 4:9.19 | כמאמר |  |
| 4:9.20 | אפלטון |  |
| 4:9.21 | שהשי"ת |  |
| 4:9.22 | יעיין |  |
| 4:9.23 | בעולם |  |
| 4:9.24 | השכלים |  |
| 4:9.25 | וישפיע |  |
| 4:9.26 | ממנו |  |
| 4:9.27 | המציאות | . |
| 4:9.28 | והנוסחא |  |
| 4:9.29 | השנית |  |
| 4:9.30 | היא |  |
| 4:9.31 | שאמרו |  |
| 4:9.32 | במקום |  |
| 4:9.33 | אחר |  |
| 4:9.34 | אין |  |
| 4:9.35 | הקב"ה |  |
| 4:9.36 | עושה |  |
| 4:9.37 | דבר |  |
| 4:9.38 | עד |  |
| 4:9.39 | שנמלך |  |
| 4:9.40 | בפימיליא |  |
| 4:9.41 | שלמעלה | . |
| 4:9.42 | ופימליא |  |
| 4:9.43 | הוא |  |
| 4:9.44 | החיל |  |
| 4:9.45 | בלשון |  |
| 4:9.46 | יון | . |
| 4:9.47 | ובבראשית |  |
| 4:9.48 | רבה |  |
| 4:9.49 | ובמדרש |  |
| 4:9.50 | קהלת |  |
| 4:9.51 | את |  |
| 4:9.52 | אשר |  |
| 4:9.53 | כבר |  |
| 4:9.54 | עשוהו | . |
| 4:9.55 | עשהו |  |
| 4:9.56 | לא |  |
| 4:9.57 | נאמר |  |
| 4:9.58 | אלא |  |
| 4:9.59 | עשוהו | . |
| 4:9.60 | כביכול |  |
| 4:9.61 | הוא |  |
| 4:9.62 | ובית |  |
| 4:9.63 | דינו |  |
| 4:9.64 | נמנו |  |
| 4:9.65 | על |  |
| 4:9.66 | כל |  |
| 4:9.67 | אבר |  |
| 4:9.68 | ואבר |  |
| 4:9.69 | שבך |  |
| 4:9.70 | והושיבוהו |  |
| 4:9.71 | על |  |
| 4:9.72 | כנו |  |
| 4:9.73 | שנ | ' |
| 4:9.74 | הוא |  |
| 4:9.75 | עשך |  |
| 4:9.76 | ויכוננך | , |
| 4:9.77 | ובבראשית |  |
| 4:9.78 | רבה |  |
| 4:9.79 | אמרו |  |
| 4:9.80 | ג"כ |  |
| 4:9.81 | כל |  |
| 4:9.82 | מקום |  |
| 4:9.83 | שנ | ' |
| 4:9.84 | וה | ' |
| 4:9.85 | הוא |  |
| 4:9.86 | ובית |  |
| 4:9.87 | דנו | , |
| 4:9.88 | ואין |  |
| 4:9.89 | הכוונה |  |
| 4:9.90 | כאלו |  |
| 4:9.91 | המאמרים |  |
| 4:9.92 | מה |  |
| 4:9.93 | שהשבוהו |  |
| 4:9.94 | הפתאים |  |
| 4:9.95 | שיש |  |
| 4:9.96 | לו |  |
| 4:9.97 | ית | ' |
| 4:9.98 | דברים |  |
| 4:9.99 | או |  |
| 4:9.100 | מחשבה |  |
| 4:9.101 | או |  |
| 4:9.102 | שאלת |  |
| 4:9.103 | עצה |  |
| 4:9.104 | או |  |
| 4:9.105 | התבוננות | , |
| 4:9.106 | והעזר |  |
| 4:9.107 | בדעת |  |
| 4:9.108 | אחרים | , |
| 4:9.109 | ואיך |  |
| 4:9.110 | יעזר |  |
| 4:9.111 | הבורא |  |
| 4:9.112 | במה |  |
| 4:9.113 | שברא | , |
| 4:9.114 | אבל |  |
| 4:9.115 | זה |  |
| 4:9.116 | כולו |  |
| 4:9.117 | עניינו |  |
| 4:9.118 | שאפי | ' |
| 4:9.119 | חלקי |  |
| 4:9.120 | המציאות |  |
| 4:9.121 | עד |  |
| 4:9.122 | בריאת |  |
| 4:9.123 | האברים |  |
| 4:9.124 | מבעלי |  |
| 4:9.125 | חיים |  |
| 4:9.126 | כפי |  |
| 4:9.127 | מה |  |
| 4:9.128 | שהם |  |
| 4:9.129 | עליו | . |
| 4:9.130 | כל |  |
| 4:9.131 | זה |  |
| 4:9.132 | באמצעית |  |
| 4:9.133 | מלאכים | , |
| 4:9.134 | כי |  |
| 4:9.135 | הכחות |  |
| 4:9.136 | כולם |  |
| 4:9.137 | מלאכים | . |
| 4:9.138 | ומה |  |
| 4:9.139 | מאד |  |
| 4:9.140 | רע |  |
| 4:9.141 | עורון |  |
| 4:9.142 | הסכלות | , |
| 4:9.143 | ומה |  |
| 4:9.144 | מאד |  |
| 4:9.145 | מזיק | , |
| 4:9.146 | שאלו |  |
| 4:9.147 | תאמר |  |
| 4:9.148 | לאחד |  |
| 4:9.149 | מאשר |  |
| 4:9.150 | יחשבו |  |
| 4:9.151 | שהם |  |
| 4:9.152 | חכמי |  |
| 4:9.153 | ישראל |  |
| 4:9.154 | וכו | ' |
| 4:9.155 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

