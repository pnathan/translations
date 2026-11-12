# Review pack: milhamot-hashem perek-2-09 (Milchamot ha-Shem 2:17–2:18)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-09`.

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


### Translation under review (2:17)

English: And greater than this is what they also do: they sacrifice before a king or a minister sheep and cattle. It is one of the inner worships. Whoever does it for one of the kinds of idolatry, even though its worship is not in that way, is liable, as we learned in Sanhedrin, chapter four, on the four death penalties: “The worshipper of idolatry: one who worships, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, etc.” And we say in the Gemara: On this: “What is ‘one who worships’?” Rabbi Yirmiyah said: “This is what it means: one who worships in the usual way, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, even not in the usual way.” And so all the decisors ruled that for the four worships one is liable even if not in the usual way. And this practice is still observed to this day among our brethren, the people of the villages and the towns: that if one of them had an enemy who seeks his harm, he brings a lamb and slaughters it before his minister, to save him from his enemy. And Scripture also says: “Offer it now to your governor, etc.” Our forebears also had this custom: when the king came to the province, they sacrificed an ox before him, and they took the price of the ox from the communal chest. And no one was heard to doubt this on account of idolatry, since one does not intend to serve him by this or to accept him upon oneself as a god; it is only a sign of honor that one makes to him. Our father Jacob, peace be upon him, said to his sons: “Take of the choicest of the land's produce in your vessels, and bring down to the man a gift,” and it is said: “And they brought him the gift that was in their hands to the house, and bowed down to him, faces, etc.” And concerning Saul, the LORD's chosen one, it is said: “And sons of Belial said, ‘How can this one save us?’ and did not bring him a gift.” And in Nedarim we learned: “One who vows by an offering and says, ‘I vowed only by the offerings of kings.’” So it is clear that bowing down, sacrifice, bringing a meal offering, and all four inner worships are permitted to do for a man who is not worshipped like Haman; but if he is worshipped like Haman it is forbidden, for Mordechai refused to bow down to Haman, because he made himself an object of idolatry.

Literal: and-great from-this that-they-do also so to-sacrifice before king or minister sheep and-cattle and-it of-worships of-the-inside that-the-one-who-does for-one of-kinds-of idolatry even though-not its-worship in-that liable as-we-learned in-Sanhedrin chapter four deaths the-worshipper-of idolatry one who-worships and-one the-altar and-one the-burner and-one the-pourer and-one the-bower etc and-we-say in-the-Gemara on-this what-is one who-worships said Rabbi Yirmiyah this-is what-it-means one who-worships in-its-way and-one the-altar and-one the-burner and-one the-pourer and-one the-bower even not in-its-way and-so all the-decisors ruled that-for the-four worships liable-on-them even not in-its-way and-this practised until the-day among our-brethren the-people-of the-villages and-the-towns that-if to-one of-them was enemy who-seeks his-harm he brings lamb and-slaughters-it before the-minister his to-save-him from-his-enemy and-Scripture also says offer-it now to-your-governor etc our-forebears had upon-them also the-custom this when came the-king to the-province to-sacrifice before-him ox and-they-took the-bloods-of-the-ox from-chest of-community and-not heard any doubter in-this on-account-of idolatry since not intends to-serve-him by-this and-to-accept-him upon-him as-a-god only sign of-honor it making to-him and-Jacob our-father peace-upon-him said to-his-sons take the-choicest of-the-land in-your-vessels and-bring-down to-the-man gift and-it-is-said and-they-brought to-him the-gift that in-their-hands to-the-house and-bowed to-him faces etc and-concerning-Saul chosen of-the-LORD it-is-said and-sons-of Belial said how will-save-us this and-not brought to-him gift and-in-Nedarim we-learned one vows by-offering and-says not I-vowed only by-offerings of-kings so-clear that-bowing and-sacrifice and-bringing meal-offering and-all the-four worships inner are-permitted to-do-them for-man who-is-not worshipped like-Haman but-if he is-worshipped like-Haman forbidden for refrained Mordechai to-bow-down to-Haman because he-made himself an-object-of idolatry

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:17.1 | and greater |  |  |  |
| 2:17.2 | than this |  |  |  |
| 2:17.3 | that they do |  |  |  |
| 2:17.4 | also |  |  |  |
| 2:17.5 | likewise |  |  |  |
| 2:17.6 | to sacrifice |  |  |  |
| 2:17.7 | before |  |  |  |
| 2:17.8 | a king |  |  |  |
| 2:17.9 | or |  |  |  |
| 2:17.10 | a minister |  |  |  |
| 2:17.11 | sheep |  |  |  |
| 2:17.12 | and cattle |  |  |  |
| 2:17.13 | and it is |  |  |  |
| 2:17.14 | one of the worships |  |  |  |
| 2:17.15 | of the inner kind |  |  | avodot penim, the 'inner' worships: the four acts the Mishnah names below; literally 'worships of the inside'. |
| 2:17.16 | that the one who does |  |  |  |
| 2:17.17 | for one |  |  |  |
| 2:17.18 | of the kinds of |  |  |  |
| 2:17.19 | idolatry |  |  | Literally 'strange worship' (avodah zarah); rendered 'idolatry' throughout. |
| 2:17.20 | even |  |  |  |
| 2:17.21 | though |  |  |  |
| 2:17.22 | its worship |  |  |  |
| 2:17.23 | in that way |  |  |  |
| 2:17.24 | is liable |  |  |  |
| 2:17.25 | as we learned |  |  |  |
| 2:17.26 | in Sanhedrin |  |  | Tractate Sanhedrin of the Mishnah; its chapter is given as the author writes it. |
| 2:17.27 | chapter |  |  |  |
| 2:17.28 | four |  |  |  |
| 2:17.29 | death penalties |  |  |  |
| 2:17.30 | the worshipper of |  |  |  |
| 2:17.31 | idolatry |  |  |  |
| 2:17.32 | one |  |  |  |
| 2:17.33 | who worships |  |  |  |
| 2:17.34 | and one |  |  |  |
| 2:17.35 | the altar |  |  | The Hebrew names the altar; the sense is the one who offers on it. |
| 2:17.36 | and one |  |  |  |
| 2:17.37 | who burns incense |  |  |  |
| 2:17.38 | and one |  |  |  |
| 2:17.39 | who pours libation |  |  |  |
| 2:17.40 | and one |  |  |  |
| 2:17.41 | who bows down |  |  |  |
| 2:17.42 | etc. |  |  |  |
| 2:17.43 | and we say |  |  |  |
| 2:17.44 | in the Gemara |  |  | The Gemara is the Talmud's discussion of the Mishnah. |
| 2:17.45 | on this |  |  | The Gemara asks what the Mishnah's first 'one who worships' means. |
| 2:17.46 | what is |  |  |  |
| 2:17.47 | one |  |  |  |
| 2:17.48 | who worships |  |  |  |
| 2:17.49 | said |  |  |  |
| 2:17.50 | Rabbi |  |  |  |
| 2:17.51 | Yirmiyah |  |  |  |
| 2:17.52 | this is |  |  |  |
| 2:17.53 | what it means |  |  |  |
| 2:17.54 | one |  |  |  |
| 2:17.55 | who worships |  |  |  |
| 2:17.56 | in the usual way |  |  | Literally 'in its way': in the usual manner. |
| 2:17.57 | and one |  |  |  |
| 2:17.58 | who sacrifices at the altar |  |  |  |
| 2:17.59 | and one |  |  |  |
| 2:17.60 | who burns incense |  |  |  |
| 2:17.61 | and one |  |  |  |
| 2:17.62 | who pours libation |  |  |  |
| 2:17.63 | and one |  |  |  |
| 2:17.64 | who bows down |  |  |  |
| 2:17.65 | even |  |  |  |
| 2:17.66 | not |  |  |  |
| 2:17.67 | in the usual way |  |  |  |
| 2:17.68 | and so |  |  |  |
| 2:17.69 | ruled |  |  |  |
| 2:17.70 | all |  |  |  |
| 2:17.71 | the decisors |  |  |  |
| 2:17.72 | that for the four |  |  |  |
| 2:17.73 | worships |  |  |  |
| 2:17.74 | one is liable |  |  |  |
| 2:17.75 | for them |  |  |  |
| 2:17.76 | even |  |  |  |
| 2:17.77 | not |  |  |  |
| 2:17.78 | in the usual way |  |  |  |
| 2:17.79 | and this |  |  |  |
| 2:17.80 | this |  |  |  |
| 2:17.81 | is practised |  |  |  |
| 2:17.82 | to |  |  |  |
| 2:17.83 | this day |  |  |  |
| 2:17.84 | among |  |  |  |
| 2:17.85 | our brethren |  |  |  |
| 2:17.86 | the people of |  |  |  |
| 2:17.87 | the villages |  |  |  |
| 2:17.88 | and the towns |  |  |  |
| 2:17.89 | that if |  |  |  |
| 2:17.90 | had |  |  |  |
| 2:17.91 | one of |  |  |  |
| 2:17.92 | them |  |  |  |
| 2:17.93 | an enemy |  |  |  |
| 2:17.94 | who |  |  |  |
| 2:17.95 | seeks |  |  |  |
| 2:17.96 | his harm |  |  |  |
| 2:17.97 | he |  |  |  |
| 2:17.98 | brings |  |  |  |
| 2:17.99 | a lamb |  |  |  |
| 2:17.100 | and slaughters it |  |  |  |
| 2:17.101 | before |  |  |  |
| 2:17.102 | the minister |  |  |  |
| 2:17.103 | his |  |  |  |
| 2:17.104 | to save him |  |  |  |
| 2:17.105 | from his enemy |  |  |  |
| 2:17.106 | and Scripture |  |  |  |
| 2:17.107 | also |  |  |  |
| 2:17.108 | says |  |  |  |
| 2:17.109 | offer it |  |  |  |
| 2:17.110 | now |  |  |  |
| 2:17.111 | to your governor |  |  |  |
| 2:17.112 | etc. |  |  |  |
| 2:17.113 | our forebears |  |  |  |
| 2:17.114 | had |  |  |  |
| 2:17.115 | upon them |  |  |  |
| 2:17.116 | also |  |  |  |
| 2:17.117 | the custom |  |  |  |
| 2:17.118 | this |  |  |  |
| 2:17.119 | when |  |  |  |
| 2:17.120 | the king |  |  |  |
| 2:17.121 | to |  |  |  |
| 2:17.122 | the province |  |  | medinah: a province or district. |
| 2:17.123 | to sacrifice |  |  |  |
| 2:17.124 | before him |  |  |  |
| 2:17.125 | an ox |  |  |  |
| 2:17.126 | and they took |  |  |  |
| 2:17.127 | - |  |  |  |
| 2:17.128 | the price |  |  | dmei: literally 'bloods'; here the price of the ox, which is the sense the sentence needs. The sense 'blood' is not taken. |
| 2:17.129 | of the ox |  |  |  |
| 2:17.130 | from the chest |  |  |  |
| 2:17.131 | of |  |  |  |
| 2:17.132 | the community |  |  |  |
| 2:17.133 | and no one |  |  |  |
| 2:17.134 | was heard |  |  |  |
| 2:17.135 | any |  |  |  |
| 2:17.136 | doubter |  |  |  |
| 2:17.137 | about this |  |  |  |
| 2:17.138 | on account of |  |  |  |
| 2:17.139 | idolatry |  |  |  |
| 2:17.140 | since |  |  |  |
| 2:17.141 | one does not |  |  |  |
| 2:17.142 | intend |  |  |  |
| 2:17.143 | to serve him |  |  |  |
| 2:17.144 | by this |  |  |  |
| 2:17.145 | or to accept |  |  |  |
| 2:17.146 | upon himself |  |  |  |
| 2:17.147 | as a god |  |  | be-eloah, 'as a god': the author's own term, applied here to the king in his argument. |
| 2:17.148 | only |  |  |  |
| 2:17.149 | a sign |  |  |  |
| 2:17.150 | of honor |  |  |  |
| 2:17.151 | it |  |  |  |
| 2:17.152 | one makes |  |  |  |
| 2:17.153 | to him |  |  |  |
| 2:17.154 | and Jacob |  |  |  |
| 2:17.155 | our father |  |  |  |
| 2:17.156 | peace be upon him |  |  |  |
| 2:17.157 | said |  |  |  |
| 2:17.158 | to his sons |  |  |  |
| 2:17.159 | take |  |  |  |
| 2:17.160 | the choicest |  |  | zimrat ha-aretz (Gen 43:11): the land's best produce. |
| 2:17.161 | of the land |  |  |  |
| 2:17.162 | in your vessels |  |  |  |
| 2:17.163 | and bring down |  |  |  |
| 2:17.164 | to the man |  |  |  |
| 2:17.165 | a gift |  |  |  |
| 2:17.166 | and it is said |  |  |  |
| 2:17.167 | and they brought |  |  |  |
| 2:17.168 | to him |  |  |  |
| 2:17.169 | - |  |  |  |
| 2:17.170 | the gift |  |  |  |
| 2:17.171 | that |  |  |  |
| 2:17.172 | in their hands |  |  |  |
| 2:17.173 | to the house |  |  |  |
| 2:17.174 | and bowed down |  |  |  |
| 2:17.175 | to him |  |  |  |
| 2:17.176 | faces |  |  | apayim, 'faces': the author's wording; see apparatus. |
| 2:17.177 | etc. |  |  |  |
| 2:17.178 | and concerning |  |  |  |
| 2:17.179 | the chosen |  |  |  |
| 2:17.180 | of the LORD |  |  |  |
| 2:17.181 | it is said |  |  |  |
| 2:17.182 | and sons of |  |  |  |
| 2:17.183 | Belial |  |  | Belial: worthless men, as in 1 Sam 10:27. |
| 2:17.184 | said |  |  |  |
| 2:17.185 | how |  |  |  |
| 2:17.186 | can save us |  |  |  |
| 2:17.187 | this one |  |  |  |
| 2:17.188 | and not |  |  |  |
| 2:17.189 | bring |  |  |  |
| 2:17.190 | to him |  |  |  |
| 2:17.191 | a gift |  |  |  |
| 2:17.192 | and in Nedarim |  |  |  |
| 2:17.193 | we learned |  |  |  |
| 2:17.194 | one vows |  |  |  |
| 2:17.195 | by an offering |  |  |  |
| 2:17.196 | and says |  |  |  |
| 2:17.197 | not |  |  |  |
| 2:17.198 | I vowed |  |  |  |
| 2:17.199 | only |  |  |  |
| 2:17.200 | by offerings |  |  |  |
| 2:17.201 | of kings |  |  |  |
| 2:17.202 | so it is clear |  |  |  |
| 2:17.203 | that bowing down |  |  |  |
| 2:17.204 | and sacrifice |  |  |  |
| 2:17.205 | and bringing |  |  |  |
| 2:17.206 | a meal offering |  |  | minchah here is a meal offering, the grain offering; the same word is 'gift' above. |
| 2:17.207 | and all |  |  |  |
| 2:17.208 | the four |  |  |  |
| 2:17.209 | worships |  |  |  |
| 2:17.210 | inner |  |  |  |
| 2:17.211 | are permitted |  |  |  |
| 2:17.212 | to do them |  |  |  |
| 2:17.213 | for a man |  |  |  |
| 2:17.214 | who is not |  |  |  |
| 2:17.215 | worshipped |  |  |  |
| 2:17.216 | like Haman |  |  | Haman, the villain of Esther. |
| 2:17.217 | but if |  |  |  |
| 2:17.218 | he |  |  |  |
| 2:17.219 | is worshipped |  |  |  |
| 2:17.220 | like Haman |  |  |  |
| 2:17.221 | forbidden |  |  |  |
| 2:17.222 | for |  |  |  |
| 2:17.223 | refrained |  |  |  |
| 2:17.224 | Mordechai |  |  | Mordechai, who refused to bow down to Haman (Esther 3:2). |
| 2:17.225 | to bow down |  |  |  |
| 2:17.226 | to Haman |  |  |  |
| 2:17.227 | because |  |  |  |
| 2:17.228 | he made |  |  |  |
| 2:17.229 | himself |  |  |  |
| 2:17.230 | an object of idolatry |  |  | Literally 'made himself strange worship'. |

Phrases:
- 2:17.1 2:17.2 2:17.3 2:17.4 2:17.5 2:17.6 2:17.7 2:17.8 2:17.9 2:17.10 2:17.11 2:17.12 | And greater than this is what they also do: they sacrifice before a king or a minister sheep and cattle. | and-great from-this that-they-do also so to-sacrifice before king or minister sheep and-cattle
- 2:17.13 2:17.14 2:17.15 | It is one of the inner worships. | and-it of-worships of-the-inside
- 2:17.16 2:17.17 2:17.18 2:17.19 2:17.20 2:17.21 2:17.22 2:17.23 2:17.24 | Whoever does it for one of the kinds of idolatry, even though its worship is not in that way, is liable, | that-the-one-who-does for-one of-kinds-of idolatry even though-not its-worship in-that liable
- 2:17.25 2:17.26 2:17.27 2:17.28 2:17.29 | as we learned in Sanhedrin, chapter four, on the four death penalties: | as-we-learned in-Sanhedrin chapter four deaths
- 2:17.30 2:17.31 2:17.32 2:17.33 2:17.34 2:17.35 2:17.36 2:17.37 2:17.38 2:17.39 2:17.40 2:17.41 2:17.42 | “The worshipper of idolatry: one who worships, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, etc.” | the-worshipper-of idolatry one who-worships and-one the-altar and-one the-burner and-one the-pourer and-one the-bower etc
- 2:17.43 2:17.44 | And we say in the Gemara: | and-we-say in-the-Gemara
- 2:17.45 2:17.46 2:17.47 2:17.48 | On this: “What is ‘one who worships’?” | on-this what-is one who-worships
- 2:17.49 2:17.50 2:17.51 2:17.52 2:17.53 | Rabbi Yirmiyah said: “This is what it means: | said Rabbi Yirmiyah this-is what-it-means
- 2:17.54 2:17.55 2:17.56 2:17.57 2:17.58 2:17.59 2:17.60 2:17.61 2:17.62 2:17.63 2:17.64 2:17.65 2:17.66 2:17.67 | one who worships in the usual way, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, even not in the usual way.” | one who-worships in-its-way and-one the-altar and-one the-burner and-one the-pourer and-one the-bower even not in-its-way
- 2:17.68 2:17.69 2:17.70 2:17.71 2:17.72 2:17.73 2:17.74 2:17.75 2:17.76 2:17.77 2:17.78 | And so all the decisors ruled that for the four worships one is liable even if not in the usual way. | and-so all the-decisors ruled that-for the-four worships liable-on-them even not in-its-way
- 2:17.79 2:17.80 2:17.81 2:17.82 2:17.83 2:17.84 2:17.85 2:17.86 2:17.87 2:17.88 | And this practice is still observed to this day among our brethren, the people of the villages and the towns: | and-this practised until the-day among our-brethren the-people-of the-villages and-the-towns
- 2:17.89 2:17.90 2:17.91 2:17.92 2:17.93 2:17.94 2:17.95 2:17.96 2:17.97 2:17.98 2:17.99 2:17.100 2:17.101 2:17.102 2:17.103 2:17.104 2:17.105 | that if one of them had an enemy who seeks his harm, he brings a lamb and slaughters it before his minister, to save him from his enemy. | that-if to-one of-them was enemy who-seeks his-harm he brings lamb and-slaughters-it before the-minister his to-save-him from-his-enemy
- 2:17.106 2:17.107 2:17.108 2:17.109 2:17.110 2:17.111 2:17.112 | And Scripture also says: “Offer it now to your governor, etc.” | and-Scripture also says offer-it now to-your-governor etc
- 2:17.113 2:17.114 2:17.115 2:17.116 2:17.117 2:17.118 2:17.119 2:17.120 2:17.121 2:17.122 2:17.123 2:17.124 2:17.125 | Our forebears also had this custom: when the king came to the province, they sacrificed an ox before him, | our-forebears had upon-them also the-custom this when came the-king to the-province to-sacrifice before-him ox
- 2:17.126 2:17.127 2:17.128 2:17.129 2:17.130 2:17.131 2:17.132 | and they took the price of the ox from the communal chest. | and-they-took the-bloods-of-the-ox from-chest of-community
- 2:17.133 2:17.134 2:17.135 2:17.136 2:17.137 2:17.138 2:17.139 | And no one was heard to doubt this on account of idolatry, | and-not heard any doubter in-this on-account-of idolatry
- 2:17.140 2:17.141 2:17.142 2:17.143 2:17.144 2:17.145 2:17.146 2:17.147 | since one does not intend to serve him by this or to accept him upon oneself as a god; | since not intends to-serve-him by-this and-to-accept-him upon-him as-a-god
- 2:17.148 2:17.149 2:17.150 2:17.151 2:17.152 2:17.153 | it is only a sign of honor that one makes to him. | only sign of-honor it making to-him
- 2:17.154 2:17.155 2:17.156 2:17.157 2:17.158 2:17.159 2:17.160 2:17.161 2:17.162 2:17.163 2:17.164 2:17.165 2:17.166 | Our father Jacob, peace be upon him, said to his sons: “Take of the choicest of the land's produce in your vessels, and bring down to the man a gift,” and it is said: | and-Jacob our-father peace-upon-him said to-his-sons take the-choicest of-the-land in-your-vessels and-bring-down to-the-man gift and-it-is-said
- 2:17.167 2:17.168 2:17.169 2:17.170 2:17.171 2:17.172 2:17.173 2:17.174 2:17.175 2:17.176 2:17.177 | “And they brought him the gift that was in their hands to the house, and bowed down to him, faces, etc.” | and-they-brought to-him the-gift that in-their-hands to-the-house and-bowed to-him faces etc
- 2:17.178 2:17.179 2:17.180 2:17.181 | And concerning Saul, the LORD's chosen one, it is said: | and-concerning-Saul chosen of-the-LORD it-is-said
- 2:17.182 2:17.183 2:17.184 2:17.185 2:17.186 2:17.187 2:17.188 2:17.189 2:17.190 2:17.191 | “And sons of Belial said, ‘How can this one save us?’ and did not bring him a gift.” | and-sons-of Belial said how will-save-us this and-not brought to-him gift
- 2:17.192 2:17.193 2:17.194 2:17.195 2:17.196 2:17.197 2:17.198 2:17.199 2:17.200 2:17.201 | And in Nedarim we learned: “One who vows by an offering and says, ‘I vowed only by the offerings of kings.’” | and-in-Nedarim we-learned one vows by-offering and-says not I-vowed only by-offerings of-kings
- 2:17.202 2:17.203 2:17.204 2:17.205 2:17.206 2:17.207 2:17.208 2:17.209 2:17.210 2:17.211 2:17.212 | So it is clear that bowing down, sacrifice, bringing a meal offering, and all four inner worships are permitted to do | so-clear that-bowing and-sacrifice and-bringing meal-offering and-all the-four worships inner are-permitted to-do-them
- 2:17.213 2:17.214 2:17.215 2:17.216 | for a man who is not worshipped like Haman; | for-man who-is-not worshipped like-Haman
- 2:17.217 2:17.218 2:17.219 2:17.220 2:17.221 | but if he is worshipped like Haman it is forbidden, | but-if he is-worshipped like-Haman forbidden
- 2:17.222 2:17.223 2:17.224 2:17.225 2:17.226 2:17.227 2:17.228 2:17.229 2:17.230 | for Mordechai refused to bow down to Haman, because he made himself an object of idolatry. | for refrained Mordechai to-bow-down to-Haman because he-made himself an-object-of idolatry

Grammar notes:
- שעושים לזבוח: the relative clause שעושים ('that they do') governs the infinitive לזבוח ('to sacrifice'); the object is 'sheep and cattle' and the place is 'before a king or a minister'.
- והיא מעבודות שבפנים: והיא refers back to the act in the sentence before; עבודות פנים is the Mishnah's term for the inner worships.
- עלה מאי אחד העובד: the Gemara asks what the Mishnah's first 'one who worships' means; Rabbi Yirmiyah answers with הכי קאמר ('this is what it means').
- The Mishnah's list and Rabbi Yirmiyah's list each close with 'even not in the usual way' (ואפי' שלא כדרכה): the English repeats the phrase at both points.
- נעבד (passive, 'is worshipped'), נמנע ('refrained') and the adjective אסור ('forbidden') carry the last sentences; the English keeps the passive.
- שאם היה לאחד מהם אויב ... הוא מביא: a past verb in the conditional clause and a present verb in the main clause, with a habitual sense: 'if one of them had an enemy, he brings'.

Translator's note: The Mishnah's list and Rabbi Yirmiyah's gloss are given in full, as the author quotes them. 'Idolatry' and 'inner worships' follow the fixed terms. The chapter reference is the author's and is kept.

Commentary rulings:
- apparatus: Mal 1:8: 'הקריבהו נא לפחתך' (Offer it now to your governor), quoted by the author with וכו'.
- apparatus: Gen 43:11: 'קחו מזמרת הארץ בכליכם והורידו לאיש מנחה' (Take of the choicest of the land's produce in your vessels and bring down to the man a gift).
- apparatus: Gen 43:26: 'ויביאו לו את המנחה אשר בידם הביתה'; the author continues 'וישתחו לו אפים וכו'', and the English follows his wording, not the verse's continuation.
- apparatus: 1 Sam 10:27: 'ובני בליעל אמרו מה יושיענו זה ולא הביאו לו מנחה'; the verse has an extra clause before 'ולא הביאו' that the author omits, and the English omits it too.
- apparatus: Esth 3:2: Mordechai's refusal to bow to Haman, given by the author in his own words.
- apparatus: The Mishnah of Sanhedrin is cited as 'chapter four' (פ' ד'); the reference is kept as written.
- apparatus: The quotation from Nedarim gives no reference, and none is added.

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


### Translation under review (2:18)

English: And in the book Teshu'at Yisrael by Rabbi Menasseh ben Israel he wrote, and these are his words: “I have been required to make known the customs of honor and submission with which the Jews honor the Torah scroll, for our enemies, who do not know what this is, provoke us, saying that we are worshippers of idols, and that we bow down to a wooden ark and to the sheets of skin within it. Therefore I shall set out my reply in proper order, as the good hand of the LORD is upon me. namely, every man of Israel will see himself obliged to rise and stand at the removal of a Torah scroll from the holy ark until it is placed on the reading platform, to show it to all the people and to read from it. And we find that they did so in the days of Nehemiah, as it is written: ‘And when he opened, all the people stood.’ And this they do out of the fear and honor with which they honor the words of the living God written in it, and for that same reason they bow their heads when carrying it from the ark to the reading platform, and not because of any other of the reasons that we shall mention.”

Literal: and-in-the-book Teshu'at Yisrael by-Rabbi Menasseh son-of Israel wrote and-these-his-words I-have-been-required to-make-known customs-of honor and-submission with-which-they-honor with-them the-Jews object scroll-of the-Torah for our-enemies who not know what this provoke us saying that worshippers-of idols we and-that we-bow to-ark of-wood and-to-sheets of-skin which within-it therefore I-shall-arrange my-reply in-order proper as-hand-of the-LORD good upon-me and-it every man of-Israel will-see himself obliged to-rise and-stand at-removal-of scroll-of Torah from-ark the-holy until is-placed on the-reading-platform to-show-it to-all the-people and-to-read from-it and-we-find that-so they-did in-days-of Nehemiah as-it-is-written and-when-he-opened stood all the-people and-this they-do from the-fear and-the-honor which they-honor the-words-of God living written in-it and-for-that reason they-bow their-heads when-carrying-it from-the-ark until the-reading-platform and-not because-of other-thing of the-reasons which we-shall-mention

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:18.1 | and in the book |  |  |  |
| 2:18.2 | Teshu'at |  |  | Teshu'at Yisrael: literally 'the salvation of Israel', the title of the book. |
| 2:18.3 | Yisrael |  |  |  |
| 2:18.4 | Rabbi |  |  |  |
| 2:18.5 | Menasseh |  |  |  |
| 2:18.6 | ben |  |  |  |
| 2:18.7 | Israel |  |  |  |
| 2:18.8 | wrote |  |  |  |
| 2:18.9 | and these are his words: |  |  | vezal: the formula that opens the quotation, 'and these are his words'. |
| 2:18.10 | I have been required |  |  |  |
| 2:18.11 | to make known |  |  |  |
| 2:18.12 | the customs of |  |  |  |
| 2:18.13 | honor |  |  |  |
| 2:18.14 | and submission |  |  |  |
| 2:18.15 | with which |  |  |  |
| 2:18.16 | they honor |  |  |  |
| 2:18.17 | with them |  |  |  |
| 2:18.18 | the Jews |  |  |  |
| 2:18.19 | - |  |  |  |
| 2:18.20 | the scroll of |  |  |  |
| 2:18.21 | the Torah |  |  |  |
| 2:18.22 | for |  |  |  |
| 2:18.23 | our enemies |  |  |  |
| 2:18.24 | who |  |  |  |
| 2:18.25 | do not |  |  |  |
| 2:18.26 | know |  |  |  |
| 2:18.27 | what |  |  |  |
| 2:18.28 | this is |  |  |  |
| 2:18.29 | provoke |  |  |  |
| 2:18.30 | us |  |  |  |
| 2:18.31 | saying |  |  |  |
| 2:18.32 | that |  |  |  |
| 2:18.33 | worshippers of |  |  |  |
| 2:18.34 | idols |  |  |  |
| 2:18.35 | we are |  |  |  |
| 2:18.36 | and that |  |  |  |
| 2:18.37 | we bow down |  |  |  |
| 2:18.38 | to the ark |  |  |  |
| 2:18.39 | of wood |  |  |  |
| 2:18.40 | and to the sheets |  |  |  |
| 2:18.41 | of skin |  |  |  |
| 2:18.42 | which |  |  |  |
| 2:18.43 | within it |  |  |  |
| 2:18.44 | Therefore |  |  |  |
| 2:18.45 | I shall set in order |  |  |  |
| 2:18.46 | my reply |  |  |  |
| 2:18.47 | in order |  |  |  |
| 2:18.48 | proper |  |  |  |
| 2:18.49 | as the hand of |  |  |  |
| 2:18.50 | the LORD |  |  |  |
| 2:18.51 | good |  |  |  |
| 2:18.52 | upon me |  |  |  |
| 2:18.53 | namely |  |  |  |
| 2:18.54 | every |  |  |  |
| 2:18.55 | man |  |  |  |
| 2:18.56 | of Israel |  |  |  |
| 2:18.57 | will see |  |  |  |
| 2:18.58 | - |  |  |  |
| 2:18.59 | himself |  |  |  |
| 2:18.60 | obliged |  |  |  |
| 2:18.61 | to rise |  |  |  |
| 2:18.62 | and stand |  |  |  |
| 2:18.63 | at the removal of |  |  |  |
| 2:18.64 | a scroll of |  |  |  |
| 2:18.65 | Torah |  |  |  |
| 2:18.66 | from the ark |  |  |  |
| 2:18.67 | the holy |  |  |  |
| 2:18.68 | until |  |  |  |
| 2:18.69 | it is placed |  |  |  |
| 2:18.70 | on |  |  |  |
| 2:18.71 | the reading platform |  |  | tevah: the reading platform in the synagogue, from which the Torah is read. |
| 2:18.72 | to show it |  |  |  |
| 2:18.73 | to all |  |  |  |
| 2:18.74 | the people |  |  |  |
| 2:18.75 | and to read |  |  |  |
| 2:18.76 | from it |  |  |  |
| 2:18.77 | we find |  |  |  |
| 2:18.78 | that |  |  |  |
| 2:18.79 | so |  |  |  |
| 2:18.80 | they did |  |  |  |
| 2:18.81 | in the days of |  |  |  |
| 2:18.82 | Nehemiah |  |  |  |
| 2:18.83 | as it is written |  |  |  |
| 2:18.84 | and when he opened |  |  | Neh 8:5, in the author's wording. |
| 2:18.85 | stood |  |  |  |
| 2:18.86 | all |  |  |  |
| 2:18.87 | the people |  |  |  |
| 2:18.88 | And this |  |  |  |
| 2:18.89 | they do |  |  |  |
| 2:18.90 | out of |  |  |  |
| 2:18.91 | the fear |  |  |  |
| 2:18.92 | and the honor |  |  |  |
| 2:18.93 | with which |  |  |  |
| 2:18.94 | they honor |  |  |  |
| 2:18.95 | the words of |  |  |  |
| 2:18.96 | God |  |  | Elohim chayyim, 'the living God'; an allusion, see apparatus. |
| 2:18.97 | living |  |  |  |
| 2:18.98 | written |  |  |  |
| 2:18.99 | in it |  |  |  |
| 2:18.100 | and for that same |  |  |  |
| 2:18.101 | reason |  |  |  |
| 2:18.102 | they bow |  |  |  |
| 2:18.103 | their heads |  |  |  |
| 2:18.104 | when carrying |  |  |  |
| 2:18.105 | it |  |  |  |
| 2:18.106 | from the ark |  |  |  |
| 2:18.107 | until |  |  |  |
| 2:18.108 | the reading platform |  |  | tevah: the reading platform, as in token 71. |
| 2:18.109 | and not |  |  |  |
| 2:18.110 | because of |  |  |  |
| 2:18.111 | any |  |  |  |
| 2:18.112 | other |  |  |  |
| 2:18.113 | of |  |  |  |
| 2:18.114 | the reasons |  |  |  |
| 2:18.115 | that |  |  |  |
| 2:18.116 | we shall mention |  |  |  |

Phrases:
- 2:18.1 2:18.2 2:18.3 2:18.4 2:18.5 2:18.6 2:18.7 2:18.8 2:18.9 | And in the book Teshu'at Yisrael by Rabbi Menasseh ben Israel he wrote, and these are his words: | and-in-the-book Teshu'at Yisrael by-Rabbi Menasseh son-of Israel wrote and-these-his-words
- 2:18.10 2:18.11 2:18.12 2:18.13 2:18.14 2:18.15 2:18.16 2:18.17 2:18.18 | “I have been required to make known the customs of honor and submission with which the Jews honor | I-have-been-required to-make-known customs-of honor and-submission with-which-they-honor with-them the-Jews
- 2:18.19 2:18.20 2:18.21 | the Torah scroll, | object scroll-of the-Torah
- 2:18.22 2:18.23 2:18.24 2:18.25 2:18.26 2:18.27 2:18.28 2:18.29 2:18.30 2:18.31 2:18.32 2:18.33 2:18.34 2:18.35 | for our enemies, who do not know what this is, provoke us, saying that we are worshippers of idols, | for our-enemies who not know what this provoke us saying that worshippers-of idols we
- 2:18.36 2:18.37 2:18.38 2:18.39 2:18.40 2:18.41 2:18.42 2:18.43 | and that we bow down to a wooden ark and to the sheets of skin within it. | and-that we-bow to-ark of-wood and-to-sheets of-skin which within-it
- 2:18.44 2:18.45 2:18.46 2:18.47 2:18.48 2:18.49 2:18.50 2:18.51 2:18.52 | Therefore I shall set out my reply in proper order, as the good hand of the LORD is upon me. | therefore I-shall-arrange my-reply in-order proper as-hand-of the-LORD good upon-me
- 2:18.53 2:18.54 2:18.55 2:18.56 2:18.57 2:18.58 2:18.59 2:18.60 2:18.61 2:18.62 2:18.63 2:18.64 2:18.65 2:18.66 2:18.67 2:18.68 2:18.69 2:18.70 2:18.71 2:18.72 2:18.73 2:18.74 2:18.75 2:18.76 | namely, every man of Israel will see himself obliged to rise and stand at the removal of a Torah scroll from the holy ark until it is placed on the reading platform, to show it to all the people and to read from it. | and-it every man of-Israel will-see himself obliged to-rise and-stand at-removal-of scroll-of Torah from-ark the-holy until is-placed on the-reading-platform to-show-it to-all the-people and-to-read from-it
- 2:18.77 2:18.78 2:18.79 2:18.80 2:18.81 2:18.82 2:18.83 2:18.84 2:18.85 2:18.86 2:18.87 | And we find that they did so in the days of Nehemiah, as it is written: ‘And when he opened, all the people stood.’ | and-we-find that-so they-did in-days-of Nehemiah as-it-is-written and-when-he-opened stood all the-people
- 2:18.88 2:18.89 2:18.90 2:18.91 2:18.92 2:18.93 2:18.94 2:18.95 2:18.96 2:18.97 2:18.98 2:18.99 | And this they do out of the fear and honor with which they honor the words of the living God written in it, | and-this they-do from the-fear and-the-honor which they-honor the-words-of God living written in-it
- 2:18.100 2:18.101 2:18.102 2:18.103 2:18.104 2:18.105 2:18.106 2:18.107 2:18.108 | and for that same reason they bow their heads when carrying it from the ark to the reading platform, | and-for-that reason they-bow their-heads when-carrying-it from-the-ark until the-reading-platform
- 2:18.109 2:18.110 2:18.111 2:18.112 2:18.113 2:18.114 2:18.115 2:18.116 | and not because of any other of the reasons that we shall mention.” | and-not because-of other-thing of the-reasons which we-shall-mention

Grammar notes:
- נדרשתי להודיע: a niphal perfect, 'I have been required', with the infinitive להודיע, 'to make known'; the first person is the letter-writer's, quoted by the author.
- ככתוב וכפתחו עמדו כל העם: וכפתחו is a temporal clause ('when he opened'); עמדו is third plural, 'they stood'; the quotation is inside the quoted letter, so it takes single marks.
- והוא כל איש ישראל יראה: והוא introduces the content ('namely'); יראה is a yiqtol with a future sense, 'will see himself obliged'.
- ולא משום דבר אחר מן הטעמים אשר נזכיר: the last clause denies every other reason, so the reasons given are the only ones; the witness has no closing formula, and the quotation runs to the end of the paragraph.

Translator's note: The letter is the one Rabbi Menasseh ben Israel is quoted as writing; 'the reading platform' renders tevah. The witness has no closing formula, so the quotation runs to the end of the paragraph.

Commentary rulings:
- apparatus: Allusion to Ezra 7:28 ('the hand of the LORD upon me'); the English follows the author's wording.
- apparatus: Allusion to Jer 10:10 ('the living God').
- apparatus: Neh 8:5: the author writes 'וכפתחו עמדו כל העם'; his wording is followed.
- apparatus: The title of the book is transliterated per the project scheme; its meaning is at token 2.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

