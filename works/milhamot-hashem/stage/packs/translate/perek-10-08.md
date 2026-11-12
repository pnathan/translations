# Translation pack: milhamot-hashem perek-10-08 (Milchamot ha-Shem 10:9–10:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-08`.

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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הנוצרים / הנצרים** → "the Christians".
- מאמיני **השלוש** → "the believers in the Trinity".
## Terms

- אל **אחר** → "another god".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- דברי **תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **חסר / חטר**: "(חסר)" → "(missing)", a lacuna sign; "(חטר)" is a variant of
  it, recorded.

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
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **אצהאר / גילוי**: ספר אצהאר אלחק = גילוי האמת; the title as each place writes
  it ("Iẓhar al-Chaqq", "Gilui ha-Emet"), a note that they are one book; its
  author unnamed.
- **העתקתו / תרגום**: his own Hebrew translations (10:7–10:9; "[תרגום הרב
  המחבר: …]") are separate segments, each from its own Hebrew; a difference
  in sense in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 10:9

Hebrew (Wikisource): העתק לשון ערבי בסי' קכ״ז ובס' גילוי האמת לאחד מחכמי הערביים דף רפ״א ורפ״ב. העתקתו ללשון הקדש. בטענתו נגד אמונת הנצרים כ' וז״ל מועתק וחכמי היהודים מאחר משה עד עתה אין מאמינים בשלוש, וברור הוא שהוא ית' ותאריו המורים על שלימותו קדומים בלי שינוי מצויים בו ית' לעולם ועד, ואלו היה השלוש אמת היה למשה רבינו ע״ה ושאר נביאי ישראל בארו תכלית הביאור, ומן התימה הגדול הוא שתהיה תורת משה שהיתה מחוייבת להאמינה ולהכנע לה בכל הנביאים עד זמן יש״ו ריקנית מביאור אמונה זו שהיא מביאה לחיי נצח כפי מחשבת השלישים, ולא יתכן לשום אדם להנצל מן הדין בזולתה יהיה נביא או זולת, ולא פירש משה ע״ה ולא שום נביא מנביאי ישראל אמונה זו בביאור מספיק שיובן ממנו אמונה זו בגלוי, ולא ישאר בה שום ספק, ויבאר משה דיניה שהיא אצל קדושי השלישים חסרה ופחותה מאד בביאור שלם, וישנה וישלש חוקיה פעם אחר פעם, ויזהיר על שמירתה אזהרה מעולה, ויחייב מיתה למי שלא יאמינה ע"כ.

| id | word | punctuation after |
|---|---|---|
| 10:9.1 | העתק |  |
| 10:9.2 | לשון |  |
| 10:9.3 | ערבי |  |
| 10:9.4 | בסי | ' |
| 10:9.5 | קכ״ז |  |
| 10:9.6 | ובס | ' |
| 10:9.7 | גילוי |  |
| 10:9.8 | האמת |  |
| 10:9.9 | לאחד |  |
| 10:9.10 | מחכמי |  |
| 10:9.11 | הערביים |  |
| 10:9.12 | דף |  |
| 10:9.13 | רפ״א |  |
| 10:9.14 | ורפ״ב | . |
| 10:9.15 | העתקתו |  |
| 10:9.16 | ללשון |  |
| 10:9.17 | הקדש | . |
| 10:9.18 | בטענתו |  |
| 10:9.19 | נגד |  |
| 10:9.20 | אמונת |  |
| 10:9.21 | הנצרים |  |
| 10:9.22 | כ | ' |
| 10:9.23 | וז״ל |  |
| 10:9.24 | מועתק |  |
| 10:9.25 | וחכמי |  |
| 10:9.26 | היהודים |  |
| 10:9.27 | מאחר |  |
| 10:9.28 | משה |  |
| 10:9.29 | עד |  |
| 10:9.30 | עתה |  |
| 10:9.31 | אין |  |
| 10:9.32 | מאמינים |  |
| 10:9.33 | בשלוש | , |
| 10:9.34 | וברור |  |
| 10:9.35 | הוא |  |
| 10:9.36 | שהוא |  |
| 10:9.37 | ית | ' |
| 10:9.38 | ותאריו |  |
| 10:9.39 | המורים |  |
| 10:9.40 | על |  |
| 10:9.41 | שלימותו |  |
| 10:9.42 | קדומים |  |
| 10:9.43 | בלי |  |
| 10:9.44 | שינוי |  |
| 10:9.45 | מצויים |  |
| 10:9.46 | בו |  |
| 10:9.47 | ית | ' |
| 10:9.48 | לעולם |  |
| 10:9.49 | ועד | , |
| 10:9.50 | ואלו |  |
| 10:9.51 | היה |  |
| 10:9.52 | השלוש |  |
| 10:9.53 | אמת |  |
| 10:9.54 | היה |  |
| 10:9.55 | למשה |  |
| 10:9.56 | רבינו |  |
| 10:9.57 | ע״ה |  |
| 10:9.58 | ושאר |  |
| 10:9.59 | נביאי |  |
| 10:9.60 | ישראל |  |
| 10:9.61 | בארו |  |
| 10:9.62 | תכלית |  |
| 10:9.63 | הביאור | , |
| 10:9.64 | ומן |  |
| 10:9.65 | התימה |  |
| 10:9.66 | הגדול |  |
| 10:9.67 | הוא |  |
| 10:9.68 | שתהיה |  |
| 10:9.69 | תורת |  |
| 10:9.70 | משה |  |
| 10:9.71 | שהיתה |  |
| 10:9.72 | מחוייבת |  |
| 10:9.73 | להאמינה |  |
| 10:9.74 | ולהכנע |  |
| 10:9.75 | לה |  |
| 10:9.76 | בכל |  |
| 10:9.77 | הנביאים |  |
| 10:9.78 | עד |  |
| 10:9.79 | זמן |  |
| 10:9.80 | יש״ו |  |
| 10:9.81 | ריקנית |  |
| 10:9.82 | מביאור |  |
| 10:9.83 | אמונה |  |
| 10:9.84 | זו |  |
| 10:9.85 | שהיא |  |
| 10:9.86 | מביאה |  |
| 10:9.87 | לחיי |  |
| 10:9.88 | נצח |  |
| 10:9.89 | כפי |  |
| 10:9.90 | מחשבת |  |
| 10:9.91 | השלישים | , |
| 10:9.92 | ולא |  |
| 10:9.93 | יתכן |  |
| 10:9.94 | לשום |  |
| 10:9.95 | אדם |  |
| 10:9.96 | להנצל |  |
| 10:9.97 | מן |  |
| 10:9.98 | הדין |  |
| 10:9.99 | בזולתה |  |
| 10:9.100 | יהיה |  |
| 10:9.101 | נביא |  |
| 10:9.102 | או |  |
| 10:9.103 | זולת | , |
| 10:9.104 | ולא |  |
| 10:9.105 | פירש |  |
| 10:9.106 | משה |  |
| 10:9.107 | ע״ה |  |
| 10:9.108 | ולא |  |
| 10:9.109 | שום |  |
| 10:9.110 | נביא |  |
| 10:9.111 | מנביאי |  |
| 10:9.112 | ישראל |  |
| 10:9.113 | אמונה |  |
| 10:9.114 | זו |  |
| 10:9.115 | בביאור |  |
| 10:9.116 | מספיק |  |
| 10:9.117 | שיובן |  |
| 10:9.118 | ממנו |  |
| 10:9.119 | אמונה |  |
| 10:9.120 | זו |  |
| 10:9.121 | בגלוי | , |
| 10:9.122 | ולא |  |
| 10:9.123 | ישאר |  |
| 10:9.124 | בה |  |
| 10:9.125 | שום |  |
| 10:9.126 | ספק | , |
| 10:9.127 | ויבאר |  |
| 10:9.128 | משה |  |
| 10:9.129 | דיניה |  |
| 10:9.130 | שהיא |  |
| 10:9.131 | אצל |  |
| 10:9.132 | קדושי |  |
| 10:9.133 | השלישים |  |
| 10:9.134 | חסרה |  |
| 10:9.135 | ופחותה |  |
| 10:9.136 | מאד |  |
| 10:9.137 | בביאור |  |
| 10:9.138 | שלם | , |
| 10:9.139 | וישנה |  |
| 10:9.140 | וישלש |  |
| 10:9.141 | חוקיה |  |
| 10:9.142 | פעם |  |
| 10:9.143 | אחר |  |
| 10:9.144 | פעם | , |
| 10:9.145 | ויזהיר |  |
| 10:9.146 | על |  |
| 10:9.147 | שמירתה |  |
| 10:9.148 | אזהרה |  |
| 10:9.149 | מעולה | , |
| 10:9.150 | ויחייב |  |
| 10:9.151 | מיתה |  |
| 10:9.152 | למי |  |
| 10:9.153 | שלא |  |
| 10:9.154 | יאמינה |  |
| 10:9.155 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

