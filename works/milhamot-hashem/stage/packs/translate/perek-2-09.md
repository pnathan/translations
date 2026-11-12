# Translation pack: milhamot-hashem perek-2-09 (Milchamot ha-Shem 2:17–2:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-09`.

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
- **ע"ה** → "peace be upon him" (honorific).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
## Terms

- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **ריבה / לרבות** → "include", never "multiply".
- **שלשה פנים** → "three aspects" (note: "faces").
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- הגשת **מנחה** (argued from Gen 43 / 1 Sam 10) → "gift"; "meal offering" only
  in a sacrificial context.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אחד ... **ואחד** of two kinds of case → "alike".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **אחי** → "my brother" (second person kept).
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מיתות**: פ' ד' מיתות (פ' + a chapter title) → "Four Death Penalties".
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


## Milchamot ha-Shem 2:17

Hebrew (Wikisource): וגדולה מזו שעושים גם כן לזבוח לפני מלך או שר צאן ובקר והיא מעבודות שבפנים שהעושה לאחד ממיני ע"ז אף שאין עבודתה בכך חייב כדתנן בסנהדרין פ' ד' מיתות העובד ע"ז אחד העובד, ואחד המזבח, ואחד המקטר ואחד המנסך, ואחד המשתחוה וכו' ואמרינן בגמ' עלה מאי אחד העובד? אמר ר' ירמיה הכי קאמר אחד העובד כדרכה, ואחד המזבח ואחד המקטר ואחד המנסך ואחד המשתחוה, ואפי' שלא כדרכה. וכן פסקו כל הפוסקים דארבע עבודות חייב עליהן אפי' שלא כדרכה. ודבר זה נהוג עד היום אצל אחינו בני הכפרים והעירות שאם היה לאחד מהם אויב שהוא מבקש רעתו הוא מביא כבש ושוחטו לפני השר שלו להושיעו מאויבו. וקרא נמי כתיב: הקריבהו נא לפחתך וכו'. קדמונינו היה עליהם ג"כ החק הזה בבוא המלך אל המדינה לזבוח לפניו שור ונוטלים את דמי השור מקופה של צבור. ולא שמעו שום מפקפק בזה משום ע"ז, לפי שאין מכוין לעבדו בזה ולקבלו עליו באלוה, רק אות כבוד הוא עושה לו. ויעקב אבינו ע"ה אמר לבניו: קחו מזמרת הארץ בכליכם והורידו לאיש מנחה, ונאמ': ויביאו לו את המנחה אשר בידם הביתה וישתחו לו אפים וכו'. ובשאול בחיר ה' נאמר: ובני בליעל אמרו מה יושיענו זה ולא הביאו לו מנחה. ובנדרים שנינו: נדר בקרבן ואמר לא נדרתי אלא בקרבנות מלכים. אלמא שהשתחויה וזיבוח והגשת מנחה וכל ארבע עבודות פנים מותר לעשותן לאדם שאינו נעבד כהמן, שאם הוא נעבד כהמן אסור, שכן נמנע מרדכי להשתחות להמן מפני שעשה עצמו ע"ז.

| id | word | punctuation after |
|---|---|---|
| 2:17.1 | וגדולה |  |
| 2:17.2 | מזו |  |
| 2:17.3 | שעושים |  |
| 2:17.4 | גם |  |
| 2:17.5 | כן |  |
| 2:17.6 | לזבוח |  |
| 2:17.7 | לפני |  |
| 2:17.8 | מלך |  |
| 2:17.9 | או |  |
| 2:17.10 | שר |  |
| 2:17.11 | צאן |  |
| 2:17.12 | ובקר |  |
| 2:17.13 | והיא |  |
| 2:17.14 | מעבודות |  |
| 2:17.15 | שבפנים |  |
| 2:17.16 | שהעושה |  |
| 2:17.17 | לאחד |  |
| 2:17.18 | ממיני |  |
| 2:17.19 | ע"ז |  |
| 2:17.20 | אף |  |
| 2:17.21 | שאין |  |
| 2:17.22 | עבודתה |  |
| 2:17.23 | בכך |  |
| 2:17.24 | חייב |  |
| 2:17.25 | כדתנן |  |
| 2:17.26 | בסנהדרין |  |
| 2:17.27 | פ | ' |
| 2:17.28 | ד | ' |
| 2:17.29 | מיתות |  |
| 2:17.30 | העובד |  |
| 2:17.31 | ע"ז |  |
| 2:17.32 | אחד |  |
| 2:17.33 | העובד | , |
| 2:17.34 | ואחד |  |
| 2:17.35 | המזבח | , |
| 2:17.36 | ואחד |  |
| 2:17.37 | המקטר |  |
| 2:17.38 | ואחד |  |
| 2:17.39 | המנסך | , |
| 2:17.40 | ואחד |  |
| 2:17.41 | המשתחוה |  |
| 2:17.42 | וכו | ' |
| 2:17.43 | ואמרינן |  |
| 2:17.44 | בגמ | ' |
| 2:17.45 | עלה |  |
| 2:17.46 | מאי |  |
| 2:17.47 | אחד |  |
| 2:17.48 | העובד | ? |
| 2:17.49 | אמר |  |
| 2:17.50 | ר | ' |
| 2:17.51 | ירמיה |  |
| 2:17.52 | הכי |  |
| 2:17.53 | קאמר |  |
| 2:17.54 | אחד |  |
| 2:17.55 | העובד |  |
| 2:17.56 | כדרכה | , |
| 2:17.57 | ואחד |  |
| 2:17.58 | המזבח |  |
| 2:17.59 | ואחד |  |
| 2:17.60 | המקטר |  |
| 2:17.61 | ואחד |  |
| 2:17.62 | המנסך |  |
| 2:17.63 | ואחד |  |
| 2:17.64 | המשתחוה | , |
| 2:17.65 | ואפי | ' |
| 2:17.66 | שלא |  |
| 2:17.67 | כדרכה | . |
| 2:17.68 | וכן |  |
| 2:17.69 | פסקו |  |
| 2:17.70 | כל |  |
| 2:17.71 | הפוסקים |  |
| 2:17.72 | דארבע |  |
| 2:17.73 | עבודות |  |
| 2:17.74 | חייב |  |
| 2:17.75 | עליהן |  |
| 2:17.76 | אפי | ' |
| 2:17.77 | שלא |  |
| 2:17.78 | כדרכה | . |
| 2:17.79 | ודבר |  |
| 2:17.80 | זה |  |
| 2:17.81 | נהוג |  |
| 2:17.82 | עד |  |
| 2:17.83 | היום |  |
| 2:17.84 | אצל |  |
| 2:17.85 | אחינו |  |
| 2:17.86 | בני |  |
| 2:17.87 | הכפרים |  |
| 2:17.88 | והעירות |  |
| 2:17.89 | שאם |  |
| 2:17.90 | היה |  |
| 2:17.91 | לאחד |  |
| 2:17.92 | מהם |  |
| 2:17.93 | אויב |  |
| 2:17.94 | שהוא |  |
| 2:17.95 | מבקש |  |
| 2:17.96 | רעתו |  |
| 2:17.97 | הוא |  |
| 2:17.98 | מביא |  |
| 2:17.99 | כבש |  |
| 2:17.100 | ושוחטו |  |
| 2:17.101 | לפני |  |
| 2:17.102 | השר |  |
| 2:17.103 | שלו |  |
| 2:17.104 | להושיעו |  |
| 2:17.105 | מאויבו | . |
| 2:17.106 | וקרא |  |
| 2:17.107 | נמי |  |
| 2:17.108 | כתיב | : |
| 2:17.109 | הקריבהו |  |
| 2:17.110 | נא |  |
| 2:17.111 | לפחתך |  |
| 2:17.112 | וכו | '. |
| 2:17.113 | קדמונינו |  |
| 2:17.114 | היה |  |
| 2:17.115 | עליהם |  |
| 2:17.116 | ג"כ |  |
| 2:17.117 | החק |  |
| 2:17.118 | הזה |  |
| 2:17.119 | בבוא |  |
| 2:17.120 | המלך |  |
| 2:17.121 | אל |  |
| 2:17.122 | המדינה |  |
| 2:17.123 | לזבוח |  |
| 2:17.124 | לפניו |  |
| 2:17.125 | שור |  |
| 2:17.126 | ונוטלים |  |
| 2:17.127 | את |  |
| 2:17.128 | דמי |  |
| 2:17.129 | השור |  |
| 2:17.130 | מקופה |  |
| 2:17.131 | של |  |
| 2:17.132 | צבור | . |
| 2:17.133 | ולא |  |
| 2:17.134 | שמעו |  |
| 2:17.135 | שום |  |
| 2:17.136 | מפקפק |  |
| 2:17.137 | בזה |  |
| 2:17.138 | משום |  |
| 2:17.139 | ע"ז | , |
| 2:17.140 | לפי |  |
| 2:17.141 | שאין |  |
| 2:17.142 | מכוין |  |
| 2:17.143 | לעבדו |  |
| 2:17.144 | בזה |  |
| 2:17.145 | ולקבלו |  |
| 2:17.146 | עליו |  |
| 2:17.147 | באלוה | , |
| 2:17.148 | רק |  |
| 2:17.149 | אות |  |
| 2:17.150 | כבוד |  |
| 2:17.151 | הוא |  |
| 2:17.152 | עושה |  |
| 2:17.153 | לו | . |
| 2:17.154 | ויעקב |  |
| 2:17.155 | אבינו |  |
| 2:17.156 | ע"ה |  |
| 2:17.157 | אמר |  |
| 2:17.158 | לבניו | : |
| 2:17.159 | קחו |  |
| 2:17.160 | מזמרת |  |
| 2:17.161 | הארץ |  |
| 2:17.162 | בכליכם |  |
| 2:17.163 | והורידו |  |
| 2:17.164 | לאיש |  |
| 2:17.165 | מנחה | , |
| 2:17.166 | ונאמ | ': |
| 2:17.167 | ויביאו |  |
| 2:17.168 | לו |  |
| 2:17.169 | את |  |
| 2:17.170 | המנחה |  |
| 2:17.171 | אשר |  |
| 2:17.172 | בידם |  |
| 2:17.173 | הביתה |  |
| 2:17.174 | וישתחו |  |
| 2:17.175 | לו |  |
| 2:17.176 | אפים |  |
| 2:17.177 | וכו | '. |
| 2:17.178 | ובשאול |  |
| 2:17.179 | בחיר |  |
| 2:17.180 | ה | ' |
| 2:17.181 | נאמר | : |
| 2:17.182 | ובני |  |
| 2:17.183 | בליעל |  |
| 2:17.184 | אמרו |  |
| 2:17.185 | מה |  |
| 2:17.186 | יושיענו |  |
| 2:17.187 | זה |  |
| 2:17.188 | ולא |  |
| 2:17.189 | הביאו |  |
| 2:17.190 | לו |  |
| 2:17.191 | מנחה | . |
| 2:17.192 | ובנדרים |  |
| 2:17.193 | שנינו | : |
| 2:17.194 | נדר |  |
| 2:17.195 | בקרבן |  |
| 2:17.196 | ואמר |  |
| 2:17.197 | לא |  |
| 2:17.198 | נדרתי |  |
| 2:17.199 | אלא |  |
| 2:17.200 | בקרבנות |  |
| 2:17.201 | מלכים | . |
| 2:17.202 | אלמא |  |
| 2:17.203 | שהשתחויה |  |
| 2:17.204 | וזיבוח |  |
| 2:17.205 | והגשת |  |
| 2:17.206 | מנחה |  |
| 2:17.207 | וכל |  |
| 2:17.208 | ארבע |  |
| 2:17.209 | עבודות |  |
| 2:17.210 | פנים |  |
| 2:17.211 | מותר |  |
| 2:17.212 | לעשותן |  |
| 2:17.213 | לאדם |  |
| 2:17.214 | שאינו |  |
| 2:17.215 | נעבד |  |
| 2:17.216 | כהמן | , |
| 2:17.217 | שאם |  |
| 2:17.218 | הוא |  |
| 2:17.219 | נעבד |  |
| 2:17.220 | כהמן |  |
| 2:17.221 | אסור | , |
| 2:17.222 | שכן |  |
| 2:17.223 | נמנע |  |
| 2:17.224 | מרדכי |  |
| 2:17.225 | להשתחות |  |
| 2:17.226 | להמן |  |
| 2:17.227 | מפני |  |
| 2:17.228 | שעשה |  |
| 2:17.229 | עצמו |  |
| 2:17.230 | ע"ז | . |


## Milchamot ha-Shem 2:18

Hebrew (Wikisource): ובס' תשועת ישראל לר' מנשה בן ישראל כ' וז"ל: נדרשתי להודיע מנהגי הכבוד והכניעות אשר יכבדו בם היהודים את ספר התורה, כי צרינו אשר אינם יודעים מה זאת מכעיסים אותנו לאמר כי עובדי אלילים אנחנו, וכי נשתחוה לארון עץ וליריעות עור אשר בתוכה, לכן אערוך תשובתי בסדר הניאות כיד ה' הטובה עלי והוא כל איש ישראל יראה את עצמו מחוייב לקום ולעמוד בהוצאת ספר תורה מארון הקודש עדי תונח על התיבה להראותה לכל העם ולקרות בה. ומצינו כי כן עשו בימי נחמיה ככתוב וכפתחו עמדו כל העם. וזאת יעשו מפני היראה והכבוד אשר יכבדו דברי אלהים חיים הכתובים בה, ומאותו טעם יכופו ראשם בנשאם אותה מהארון עד התיבה, ולא משום דבר אחר מן הטעמים אשר נזכיר.

| id | word | punctuation after |
|---|---|---|
| 2:18.1 | ובס | ' |
| 2:18.2 | תשועת |  |
| 2:18.3 | ישראל |  |
| 2:18.4 | לר | ' |
| 2:18.5 | מנשה |  |
| 2:18.6 | בן |  |
| 2:18.7 | ישראל |  |
| 2:18.8 | כ | ' |
| 2:18.9 | וז"ל | : |
| 2:18.10 | נדרשתי |  |
| 2:18.11 | להודיע |  |
| 2:18.12 | מנהגי |  |
| 2:18.13 | הכבוד |  |
| 2:18.14 | והכניעות |  |
| 2:18.15 | אשר |  |
| 2:18.16 | יכבדו |  |
| 2:18.17 | בם |  |
| 2:18.18 | היהודים |  |
| 2:18.19 | את |  |
| 2:18.20 | ספר |  |
| 2:18.21 | התורה | , |
| 2:18.22 | כי |  |
| 2:18.23 | צרינו |  |
| 2:18.24 | אשר |  |
| 2:18.25 | אינם |  |
| 2:18.26 | יודעים |  |
| 2:18.27 | מה |  |
| 2:18.28 | זאת |  |
| 2:18.29 | מכעיסים |  |
| 2:18.30 | אותנו |  |
| 2:18.31 | לאמר |  |
| 2:18.32 | כי |  |
| 2:18.33 | עובדי |  |
| 2:18.34 | אלילים |  |
| 2:18.35 | אנחנו | , |
| 2:18.36 | וכי |  |
| 2:18.37 | נשתחוה |  |
| 2:18.38 | לארון |  |
| 2:18.39 | עץ |  |
| 2:18.40 | וליריעות |  |
| 2:18.41 | עור |  |
| 2:18.42 | אשר |  |
| 2:18.43 | בתוכה | , |
| 2:18.44 | לכן |  |
| 2:18.45 | אערוך |  |
| 2:18.46 | תשובתי |  |
| 2:18.47 | בסדר |  |
| 2:18.48 | הניאות |  |
| 2:18.49 | כיד |  |
| 2:18.50 | ה | ' |
| 2:18.51 | הטובה |  |
| 2:18.52 | עלי |  |
| 2:18.53 | והוא |  |
| 2:18.54 | כל |  |
| 2:18.55 | איש |  |
| 2:18.56 | ישראל |  |
| 2:18.57 | יראה |  |
| 2:18.58 | את |  |
| 2:18.59 | עצמו |  |
| 2:18.60 | מחוייב |  |
| 2:18.61 | לקום |  |
| 2:18.62 | ולעמוד |  |
| 2:18.63 | בהוצאת |  |
| 2:18.64 | ספר |  |
| 2:18.65 | תורה |  |
| 2:18.66 | מארון |  |
| 2:18.67 | הקודש |  |
| 2:18.68 | עדי |  |
| 2:18.69 | תונח |  |
| 2:18.70 | על |  |
| 2:18.71 | התיבה |  |
| 2:18.72 | להראותה |  |
| 2:18.73 | לכל |  |
| 2:18.74 | העם |  |
| 2:18.75 | ולקרות |  |
| 2:18.76 | בה | . |
| 2:18.77 | ומצינו |  |
| 2:18.78 | כי |  |
| 2:18.79 | כן |  |
| 2:18.80 | עשו |  |
| 2:18.81 | בימי |  |
| 2:18.82 | נחמיה |  |
| 2:18.83 | ככתוב |  |
| 2:18.84 | וכפתחו |  |
| 2:18.85 | עמדו |  |
| 2:18.86 | כל |  |
| 2:18.87 | העם | . |
| 2:18.88 | וזאת |  |
| 2:18.89 | יעשו |  |
| 2:18.90 | מפני |  |
| 2:18.91 | היראה |  |
| 2:18.92 | והכבוד |  |
| 2:18.93 | אשר |  |
| 2:18.94 | יכבדו |  |
| 2:18.95 | דברי |  |
| 2:18.96 | אלהים |  |
| 2:18.97 | חיים |  |
| 2:18.98 | הכתובים |  |
| 2:18.99 | בה | , |
| 2:18.100 | ומאותו |  |
| 2:18.101 | טעם |  |
| 2:18.102 | יכופו |  |
| 2:18.103 | ראשם |  |
| 2:18.104 | בנשאם |  |
| 2:18.105 | אותה |  |
| 2:18.106 | מהארון |  |
| 2:18.107 | עד |  |
| 2:18.108 | התיבה | , |
| 2:18.109 | ולא |  |
| 2:18.110 | משום |  |
| 2:18.111 | דבר |  |
| 2:18.112 | אחר |  |
| 2:18.113 | מן |  |
| 2:18.114 | הטעמים |  |
| 2:18.115 | אשר |  |
| 2:18.116 | נזכיר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

