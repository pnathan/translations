# Review pack: milhamot-hashem perek-2-06 (Milchamot ha-Shem 2:12–2:12)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-06`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **רחמים**: "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes
  of mercy", a different thirteen from the hermeneutical rules.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- דברי **תימה** → "astonishing words".
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- **תא חזי** → "come and see".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דמרפסן**: "הני מילי דמרפסן איגרא" → "words that split roofs".
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


## Milchamot ha-Shem 2:12

Hebrew (Wikisource): ואת זה חזיתי בתשובתך. ואספרה הני מילי דמרפסן איגרא כי השבת על דברי שאמרתי לך איך נניח הא"ס ושאר הפרצופים הקרובים אל הא"ס ונעבוד לז"א? והיתה תשובתך, לפי ששאר המידות מהם דין ומהם רחמים וזה כלול מכולם. ולזה רצה המאציל לפעול פעולותיו ע"י כלי זה וכו'. והבאת לי דמיון מן היד הגופנית שהיא עושה כל המלאכות, והראש שהוא למעלה אינו עושה כלום. גם כח הדבור שהוא פאר וכו', כל דבריך אלה אינן אלא דברי תימה! אני שאלתיך על אמרם שז"א הוא אלהינו ואותו נעבוד, ואתה משיבני מן הפועלים בנו ברצון פועלם. וכי כל הפועלים בנו ברצון פועלם אותם נעבוד? ונאמר שהם אלהינו? הלא השמש פועלת לנו כמה פעולות טובות להאיר לנו לחמם האויר והארץ ולהצמיח דשאים ולבשל הפירות. האם אותה נעבוד? הירח והכוכבים גם המה פועלים בתחתונים ברצון הבורא ית' שנאמר: וממגד תבואות שמש וממגד גרש ירחים. הנעבוד אותם? הארץ והמים והאש גם הם פועלים להצמיח דשא ואילנות ולגדלם. ובאש אנו עושים כל חפצינו לאפות ולבשל כל מאכלינו. וברוב המלאכות נופח באש פחם ומוציא כלי למעשהו. הנעבוד את אחד מהם? ורצונו ית' לפעול פעולותיו על ידי שלוחים אמצעיים פשוט ומבואר בכתובים. ישרצו המים, תוצא הארץ נפש חיה, תדשא הארץ דשא, ונאמר עושה מלאכיו רוחות, משרתיו אש לוהט, ועוד ועוד, בכל עושה שליחותו. ועם כל זה ציוונו השי"ת לעבדו לבדו, והזהיר אותנו שלא נעבוד אחד מבלעדיו. ובמעמד הנבחר בהגלותו על הר סיני נתן קולו לפני חילו, אנכי ה' אלהיך, לא יהיה לך אלהים אחרים על פני. ונאמר: אני ה' הוא שמי וכבודי לאחר לא אתן! ואני ואתה עמדנו לפני אומן שעשאנו ולא רצה הוא שנעבוד אל אחר קצר אפים (ז"א) מבלעדיו אעפ"י ששלוחו הוא לפעול איזה פעולה. ולמה לא נשמע בקולו לעבדו, ונלך ונעבוד לז"א אשר לא ציוונו?

| id | word | punctuation after |
|---|---|---|
| 2:12.1 | ואת |  |
| 2:12.2 | זה |  |
| 2:12.3 | חזיתי |  |
| 2:12.4 | בתשובתך | . |
| 2:12.5 | ואספרה |  |
| 2:12.6 | הני |  |
| 2:12.7 | מילי |  |
| 2:12.8 | דמרפסן |  |
| 2:12.9 | איגרא |  |
| 2:12.10 | כי |  |
| 2:12.11 | השבת |  |
| 2:12.12 | על |  |
| 2:12.13 | דברי |  |
| 2:12.14 | שאמרתי |  |
| 2:12.15 | לך |  |
| 2:12.16 | איך |  |
| 2:12.17 | נניח |  |
| 2:12.18 | הא"ס |  |
| 2:12.19 | ושאר |  |
| 2:12.20 | הפרצופים |  |
| 2:12.21 | הקרובים |  |
| 2:12.22 | אל |  |
| 2:12.23 | הא"ס |  |
| 2:12.24 | ונעבוד |  |
| 2:12.25 | לז"א | ? |
| 2:12.26 | והיתה |  |
| 2:12.27 | תשובתך | , |
| 2:12.28 | לפי |  |
| 2:12.29 | ששאר |  |
| 2:12.30 | המידות |  |
| 2:12.31 | מהם |  |
| 2:12.32 | דין |  |
| 2:12.33 | ומהם |  |
| 2:12.34 | רחמים |  |
| 2:12.35 | וזה |  |
| 2:12.36 | כלול |  |
| 2:12.37 | מכולם | . |
| 2:12.38 | ולזה |  |
| 2:12.39 | רצה |  |
| 2:12.40 | המאציל |  |
| 2:12.41 | לפעול |  |
| 2:12.42 | פעולותיו |  |
| 2:12.43 | ע"י |  |
| 2:12.44 | כלי |  |
| 2:12.45 | זה |  |
| 2:12.46 | וכו | '. |
| 2:12.47 | והבאת |  |
| 2:12.48 | לי |  |
| 2:12.49 | דמיון |  |
| 2:12.50 | מן |  |
| 2:12.51 | היד |  |
| 2:12.52 | הגופנית |  |
| 2:12.53 | שהיא |  |
| 2:12.54 | עושה |  |
| 2:12.55 | כל |  |
| 2:12.56 | המלאכות | , |
| 2:12.57 | והראש |  |
| 2:12.58 | שהוא |  |
| 2:12.59 | למעלה |  |
| 2:12.60 | אינו |  |
| 2:12.61 | עושה |  |
| 2:12.62 | כלום | . |
| 2:12.63 | גם |  |
| 2:12.64 | כח |  |
| 2:12.65 | הדבור |  |
| 2:12.66 | שהוא |  |
| 2:12.67 | פאר |  |
| 2:12.68 | וכו | ', |
| 2:12.69 | כל |  |
| 2:12.70 | דבריך |  |
| 2:12.71 | אלה |  |
| 2:12.72 | אינן |  |
| 2:12.73 | אלא |  |
| 2:12.74 | דברי |  |
| 2:12.75 | תימה | ! |
| 2:12.76 | אני |  |
| 2:12.77 | שאלתיך |  |
| 2:12.78 | על |  |
| 2:12.79 | אמרם |  |
| 2:12.80 | שז"א |  |
| 2:12.81 | הוא |  |
| 2:12.82 | אלהינו |  |
| 2:12.83 | ואותו |  |
| 2:12.84 | נעבוד | , |
| 2:12.85 | ואתה |  |
| 2:12.86 | משיבני |  |
| 2:12.87 | מן |  |
| 2:12.88 | הפועלים |  |
| 2:12.89 | בנו |  |
| 2:12.90 | ברצון |  |
| 2:12.91 | פועלם | . |
| 2:12.92 | וכי |  |
| 2:12.93 | כל |  |
| 2:12.94 | הפועלים |  |
| 2:12.95 | בנו |  |
| 2:12.96 | ברצון |  |
| 2:12.97 | פועלם |  |
| 2:12.98 | אותם |  |
| 2:12.99 | נעבוד | ? |
| 2:12.100 | ונאמר |  |
| 2:12.101 | שהם |  |
| 2:12.102 | אלהינו | ? |
| 2:12.103 | הלא |  |
| 2:12.104 | השמש |  |
| 2:12.105 | פועלת |  |
| 2:12.106 | לנו |  |
| 2:12.107 | כמה |  |
| 2:12.108 | פעולות |  |
| 2:12.109 | טובות |  |
| 2:12.110 | להאיר |  |
| 2:12.111 | לנו |  |
| 2:12.112 | לחמם |  |
| 2:12.113 | האויר |  |
| 2:12.114 | והארץ |  |
| 2:12.115 | ולהצמיח |  |
| 2:12.116 | דשאים |  |
| 2:12.117 | ולבשל |  |
| 2:12.118 | הפירות | . |
| 2:12.119 | האם |  |
| 2:12.120 | אותה |  |
| 2:12.121 | נעבוד | ? |
| 2:12.122 | הירח |  |
| 2:12.123 | והכוכבים |  |
| 2:12.124 | גם |  |
| 2:12.125 | המה |  |
| 2:12.126 | פועלים |  |
| 2:12.127 | בתחתונים |  |
| 2:12.128 | ברצון |  |
| 2:12.129 | הבורא |  |
| 2:12.130 | ית | ' |
| 2:12.131 | שנאמר | : |
| 2:12.132 | וממגד |  |
| 2:12.133 | תבואות |  |
| 2:12.134 | שמש |  |
| 2:12.135 | וממגד |  |
| 2:12.136 | גרש |  |
| 2:12.137 | ירחים | . |
| 2:12.138 | הנעבוד |  |
| 2:12.139 | אותם | ? |
| 2:12.140 | הארץ |  |
| 2:12.141 | והמים |  |
| 2:12.142 | והאש |  |
| 2:12.143 | גם |  |
| 2:12.144 | הם |  |
| 2:12.145 | פועלים |  |
| 2:12.146 | להצמיח |  |
| 2:12.147 | דשא |  |
| 2:12.148 | ואילנות |  |
| 2:12.149 | ולגדלם | . |
| 2:12.150 | ובאש |  |
| 2:12.151 | אנו |  |
| 2:12.152 | עושים |  |
| 2:12.153 | כל |  |
| 2:12.154 | חפצינו |  |
| 2:12.155 | לאפות |  |
| 2:12.156 | ולבשל |  |
| 2:12.157 | כל |  |
| 2:12.158 | מאכלינו | . |
| 2:12.159 | וברוב |  |
| 2:12.160 | המלאכות |  |
| 2:12.161 | נופח |  |
| 2:12.162 | באש |  |
| 2:12.163 | פחם |  |
| 2:12.164 | ומוציא |  |
| 2:12.165 | כלי |  |
| 2:12.166 | למעשהו | . |
| 2:12.167 | הנעבוד |  |
| 2:12.168 | את |  |
| 2:12.169 | אחד |  |
| 2:12.170 | מהם | ? |
| 2:12.171 | ורצונו |  |
| 2:12.172 | ית | ' |
| 2:12.173 | לפעול |  |
| 2:12.174 | פעולותיו |  |
| 2:12.175 | על |  |
| 2:12.176 | ידי |  |
| 2:12.177 | שלוחים |  |
| 2:12.178 | אמצעיים |  |
| 2:12.179 | פשוט |  |
| 2:12.180 | ומבואר |  |
| 2:12.181 | בכתובים | . |
| 2:12.182 | ישרצו |  |
| 2:12.183 | המים | , |
| 2:12.184 | תוצא |  |
| 2:12.185 | הארץ |  |
| 2:12.186 | נפש |  |
| 2:12.187 | חיה | , |
| 2:12.188 | תדשא |  |
| 2:12.189 | הארץ |  |
| 2:12.190 | דשא | , |
| 2:12.191 | ונאמר |  |
| 2:12.192 | עושה |  |
| 2:12.193 | מלאכיו |  |
| 2:12.194 | רוחות | , |
| 2:12.195 | משרתיו |  |
| 2:12.196 | אש |  |
| 2:12.197 | לוהט | , |
| 2:12.198 | ועוד |  |
| 2:12.199 | ועוד | , |
| 2:12.200 | בכל |  |
| 2:12.201 | עושה |  |
| 2:12.202 | שליחותו | . |
| 2:12.203 | ועם |  |
| 2:12.204 | כל |  |
| 2:12.205 | זה |  |
| 2:12.206 | ציוונו |  |
| 2:12.207 | השי"ת |  |
| 2:12.208 | לעבדו |  |
| 2:12.209 | לבדו | , |
| 2:12.210 | והזהיר |  |
| 2:12.211 | אותנו |  |
| 2:12.212 | שלא |  |
| 2:12.213 | נעבוד |  |
| 2:12.214 | אחד |  |
| 2:12.215 | מבלעדיו | . |
| 2:12.216 | ובמעמד |  |
| 2:12.217 | הנבחר |  |
| 2:12.218 | בהגלותו |  |
| 2:12.219 | על |  |
| 2:12.220 | הר |  |
| 2:12.221 | סיני |  |
| 2:12.222 | נתן |  |
| 2:12.223 | קולו |  |
| 2:12.224 | לפני |  |
| 2:12.225 | חילו | , |
| 2:12.226 | אנכי |  |
| 2:12.227 | ה | ' |
| 2:12.228 | אלהיך | , |
| 2:12.229 | לא |  |
| 2:12.230 | יהיה |  |
| 2:12.231 | לך |  |
| 2:12.232 | אלהים |  |
| 2:12.233 | אחרים |  |
| 2:12.234 | על |  |
| 2:12.235 | פני | . |
| 2:12.236 | ונאמר | : |
| 2:12.237 | אני |  |
| 2:12.238 | ה | ' |
| 2:12.239 | הוא |  |
| 2:12.240 | שמי |  |
| 2:12.241 | וכבודי |  |
| 2:12.242 | לאחר |  |
| 2:12.243 | לא |  |
| 2:12.244 | אתן | ! |
| 2:12.245 | ואני |  |
| 2:12.246 | ואתה |  |
| 2:12.247 | עמדנו |  |
| 2:12.248 | לפני |  |
| 2:12.249 | אומן |  |
| 2:12.250 | שעשאנו |  |
| 2:12.251 | ולא |  |
| 2:12.252 | רצה |  |
| 2:12.253 | הוא |  |
| 2:12.254 | שנעבוד |  |
| 2:12.255 | אל |  |
| 2:12.256 | אחר |  |
| 2:12.257 | קצר |  |
| 2:12.258 | אפים |  |
| 2:12.259 | ז"א | ) |
| 2:12.260 | מבלעדיו |  |
| 2:12.261 | אעפ"י |  |
| 2:12.262 | ששלוחו |  |
| 2:12.263 | הוא |  |
| 2:12.264 | לפעול |  |
| 2:12.265 | איזה |  |
| 2:12.266 | פעולה | . |
| 2:12.267 | ולמה |  |
| 2:12.268 | לא |  |
| 2:12.269 | נשמע |  |
| 2:12.270 | בקולו |  |
| 2:12.271 | לעבדו | , |
| 2:12.272 | ונלך |  |
| 2:12.273 | ונעבוד |  |
| 2:12.274 | לז"א |  |
| 2:12.275 | אשר |  |
| 2:12.276 | לא |  |
| 2:12.277 | ציוונו | ? |


### Translation under review (2:12)

English: And this I have seen in your reply. And I will tell these words that split roofs, since you replied to my words that I said to you. How shall we set aside Ein Sof and the rest of the partzufim close to Ein Sof and worship Ze'ir Anpin? And your reply was, because the other attributes are some judgment and some mercy, and this one includes them all. And for this reason the Emanator wished to perform His actions through this vessel, etc. And you brought me an analogy from the bodily hand, which does all the work, while the head, which is above, does nothing. Also the power of speech, which is glory, etc.; all these words of yours are nothing but words of astonishment! I asked you about their saying that Ze'ir Anpin is our God, and we worship him, and you answer me with those who act in him by the will of their action. Do we then worship all those who act in him by the will of their action? And shall we say that they are our God? Surely the sun acts for us in many good actions: to shine for us, to warm the air and the earth, to make grass sprout, and to ripen the fruit. Shall we worship her? The moon and the stars too act upon the lower beings by the will of the Creator, may He be blessed, as it is said: "and from the precious things of the yield of the sun, and from the precious things of the growth of the moons." Shall we worship them? The earth, the water and the fire also act to sprout grass and trees and to raise them. And with fire we do all our needs: to bake and to cook all our food. And in most crafts we blow into a charcoal fire and bring out a vessel for its work. Shall we worship one of them? And His will, may He be blessed, is to perform His actions through intermediary agents, as is plain and clear in Scripture. "Let the waters swarm, let the earth bring forth a living creature, let the earth sprout grass," and it is said: "He makes His messengers winds, His ministers a blazing fire," and more and more, in every case of one doing its errand. And yet, despite all this, the Name, may He be blessed, commanded us to worship Him alone, and warned us not to worship any one besides Him. And at the chosen standing, when He revealed Himself on Mount Sinai, He gave His voice before His host: "I am the LORD your God … You shall have no other gods before Me." And it is said: "I am the LORD, that is My name; and My glory I will not give to another!" You and I stood before the craftsman who made us, and He did not wish us to worship another god, quick to anger (Ze'ir Anpin), apart from Him, although His agent is to perform some action. And why do we not heed His voice to worship Him, and go and worship Ze'ir Anpin, which He did not command us?

Literal: and-(object-marker) this I-have-seen in-your-reply and-I-will-tell these words that-split roofs since you-replied to my-words that-I-said to-you how shall-we-set-aside Ein-Sof and-the-rest-of the-partzufim close to Ein-Sof and-worship to-Ze'ir-Anpin and-was your-reply because the-rest-of the-attributes of-them judgment and-some-of-them mercy and-this includes all-of-them and-for wished the-Emanator to-perform His-actions through vessel this etc. and-you-brought to-me analogy from the-hand bodily which does all the-work and-the-head which above does-not does anything also the-power of-speech which glory etc. all your-words these are-not but words astonishment I asked-you about their-saying that-Ze'ir-Anpin is our-God and-him we-worship and-you answer-me with those-who-act in-him by-the-will of-their-action and-is-it-that all those-who-act in-him by-the-will of-their-action them shall-we-worship and-shall-we-say that-they our-God surely the-sun acts for-us many actions good to-shine for-us to-warm the-air and-the-earth and-to-make-sprout grass and-to-ripen the-fruit whether her we-worship the-moon and-the-stars also they act below by-the-will the-Creator may-He-be-blessed as-it-is-said and-from-the-precious-things-of the-yield of-the-sun and-from-the-precious-things-of the-growth of-the-moons shall-we-worship them the-earth and-the-water and-the-fire also they act to-sprout grass and-trees and-to-raise-them and-with-fire we do all our-needs to-bake and-to-cook all our-food and-in-most the-crafts we-blow into-the-fire of-charcoal and-bring-out a-vessel for-its-work shall-we-worship (object-marker) one of-them and-His-will may-He-be-blessed is-to-perform His-actions through by-means-of agents intermediary plain and-clear in-Scripture let-swarm the-waters let-bring-forth the-earth living creature let-sprout the-earth grass and-it-is-said making His-messengers winds His-ministers fire blazing and-more and-more in-every one-doing its-errand and-yet all this commanded-us the-Name,-may-He-be-blessed to-worship alone and-warned us not we-worship any-one besides-Him and-at-the-standing the-chosen in-His-revealing on the-mountain Sinai He-gave His-voice before His-host I-am the-LORD your-God no shall-be to-you gods other before Me and-it-is-said I the-LORD that-is My-name and-My-glory to-another not I-will-give and-I and-you stood before craftsman who-made-us and-not wished He that-we-worship god another quick to-anger Ze'ir-Anpin apart-from-Him although that-His-agent is to-perform some action and-why not we-heed His-voice to-worship-Him and-go and-worship to-Ze'ir-Anpin which not He-commanded-us

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:12.1 | and (object marker) |  |  |  |
| 2:12.2 | this |  |  |  |
| 2:12.3 | I have seen |  |  |  |
| 2:12.4 | in your reply |  |  |  |
| 2:12.5 | and I will tell |  |  | ואספרה: literally 'and I will count'; here 'tell'. |
| 2:12.6 | these |  |  | Aramaic 'these'. The idiom הני מילי דמרפסן איגרא, literally 'these are words that split roofs', is translated by its sense (see the phrase's literal). |
| 2:12.7 | words |  |  |  |
| 2:12.8 | that split |  |  |  |
| 2:12.9 | roofs |  |  |  |
| 2:12.10 | since |  |  |  |
| 2:12.11 | you replied |  |  |  |
| 2:12.12 | to |  |  |  |
| 2:12.13 | my words |  |  |  |
| 2:12.14 | that I said |  |  |  |
| 2:12.15 | to you |  |  |  |
| 2:12.16 | how |  |  |  |
| 2:12.17 | shall we set aside |  |  | נניח: literally 'shall we leave' or 'set aside'. |
| 2:12.18 | Ein Sof |  |  | א"ס = אין סוף, literally 'without end'. The name takes no article in English. |
| 2:12.19 | and the rest of |  |  |  |
| 2:12.20 | the partzufim |  |  | partzuf (Kabbalistic entity), literally 'face'; transliterated. |
| 2:12.21 | close |  |  |  |
| 2:12.22 | to |  |  |  |
| 2:12.23 | Ein Sof |  |  | א"ס again: Ein Sof, as at 2:12.18. |
| 2:12.24 | and worship |  |  |  |
| 2:12.25 | to Ze'ir Anpin |  |  | ז"א = זעיר אנפין, literally 'Small Face' (Aramaic). Ze'ir Anpin is transliterated; 'Ze'ir' alone where the text says זעיר. |
| 2:12.26 | and was |  |  |  |
| 2:12.27 | your reply |  |  |  |
| 2:12.28 | because |  |  |  |
| 2:12.29 | the rest of |  |  |  |
| 2:12.30 | the attributes |  |  | המידות: literally 'the measures'; the attributes of the sefirot or of God in the Kabbalistic sense. |
| 2:12.31 | of them |  |  |  |
| 2:12.32 | judgment |  |  |  |
| 2:12.33 | and some of them |  |  |  |
| 2:12.34 | mercy |  |  |  |
| 2:12.35 | and this |  |  |  |
| 2:12.36 | includes |  |  |  |
| 2:12.37 | all of them |  |  |  |
| 2:12.38 | and for |  |  |  |
| 2:12.39 | wished |  |  |  |
| 2:12.40 | the Emanator |  |  |  |
| 2:12.41 | to perform |  |  |  |
| 2:12.42 | His actions |  |  |  |
| 2:12.43 | through |  |  | ע"י = על ידי, 'by means of'. |
| 2:12.44 | vessel |  |  |  |
| 2:12.45 | this |  |  |  |
| 2:12.46 | etc. |  |  | וכו' = וכולי: the author's sentence stops short. 'etc.' is kept, and nothing is supplied. |
| 2:12.47 | and you brought |  |  |  |
| 2:12.48 | to me |  |  |  |
| 2:12.49 | analogy |  |  |  |
| 2:12.50 | from |  |  |  |
| 2:12.51 | the hand |  |  |  |
| 2:12.52 | bodily |  |  |  |
| 2:12.53 | which |  |  |  |
| 2:12.54 | does |  |  |  |
| 2:12.55 | all |  |  |  |
| 2:12.56 | the work |  |  |  |
| 2:12.57 | and the head |  |  |  |
| 2:12.58 | which |  |  |  |
| 2:12.59 | above |  |  |  |
| 2:12.60 | does not |  |  |  |
| 2:12.61 | does |  |  |  |
| 2:12.62 | anything |  |  |  |
| 2:12.63 | also |  |  |  |
| 2:12.64 | the power |  |  |  |
| 2:12.65 | of speech |  |  |  |
| 2:12.66 | which |  |  |  |
| 2:12.67 | glory |  |  |  |
| 2:12.68 | etc. |  |  | וכו' = וכולי: kept as 'etc.', as at 2:12.46. |
| 2:12.69 | all |  |  |  |
| 2:12.70 | your words |  |  |  |
| 2:12.71 | these |  |  |  |
| 2:12.72 | are not |  |  |  |
| 2:12.73 | but |  |  |  |
| 2:12.74 | words |  |  |  |
| 2:12.75 | astonishment |  |  |  |
| 2:12.76 | I |  |  |  |
| 2:12.77 | asked you |  |  |  |
| 2:12.78 | about |  |  |  |
| 2:12.79 | their saying |  |  |  |
| 2:12.80 | that Ze'ir Anpin |  |  | ש + ז"א: 'that Ze'ir Anpin'. |
| 2:12.81 | is |  |  |  |
| 2:12.82 | our God |  |  |  |
| 2:12.83 | and him |  |  |  |
| 2:12.84 | we worship |  |  |  |
| 2:12.85 | and you |  |  |  |
| 2:12.86 | answer me |  |  |  |
| 2:12.87 | with |  |  |  |
| 2:12.88 | those who act |  |  |  |
| 2:12.89 | in him |  |  |  |
| 2:12.90 | by the will |  |  |  |
| 2:12.91 | of their action |  |  |  |
| 2:12.92 | and is it that |  |  |  |
| 2:12.93 | all |  |  |  |
| 2:12.94 | those who act |  |  |  |
| 2:12.95 | in him |  |  |  |
| 2:12.96 | by the will |  |  |  |
| 2:12.97 | of their action |  |  |  |
| 2:12.98 | them |  |  |  |
| 2:12.99 | shall we worship |  |  |  |
| 2:12.100 | and shall we say |  |  |  |
| 2:12.101 | that they |  |  |  |
| 2:12.102 | our God |  |  |  |
| 2:12.103 | surely |  |  |  |
| 2:12.104 | the sun |  |  |  |
| 2:12.105 | acts |  |  |  |
| 2:12.106 | for us |  |  |  |
| 2:12.107 | many |  |  |  |
| 2:12.108 | actions |  |  |  |
| 2:12.109 | good |  |  |  |
| 2:12.110 | to shine |  |  |  |
| 2:12.111 | for us |  |  |  |
| 2:12.112 | to warm |  |  |  |
| 2:12.113 | the air |  |  |  |
| 2:12.114 | and the earth |  |  |  |
| 2:12.115 | and to make sprout |  |  |  |
| 2:12.116 | grass |  |  |  |
| 2:12.117 | and to ripen |  |  |  |
| 2:12.118 | the fruit |  |  |  |
| 2:12.119 | whether |  |  |  |
| 2:12.120 | her |  |  |  |
| 2:12.121 | we worship |  |  |  |
| 2:12.122 | the moon |  |  |  |
| 2:12.123 | and the stars |  |  |  |
| 2:12.124 | also |  |  |  |
| 2:12.125 | they |  |  |  |
| 2:12.126 | act |  |  |  |
| 2:12.127 | below |  |  |  |
| 2:12.128 | by the will |  |  |  |
| 2:12.129 | the Creator |  |  |  |
| 2:12.130 | may He be blessed |  |  | ית' = יתברך, 'may He be blessed', placed after the name as the Hebrew places it. |
| 2:12.131 | as it is said |  |  |  |
| 2:12.132 | and from the precious things of |  |  |  |
| 2:12.133 | the yield |  |  |  |
| 2:12.134 | of the sun |  |  |  |
| 2:12.135 | and from the precious things of |  |  |  |
| 2:12.136 | the growth |  |  |  |
| 2:12.137 | of the moons |  |  |  |
| 2:12.138 | shall we worship |  |  |  |
| 2:12.139 | them |  |  |  |
| 2:12.140 | the earth |  |  |  |
| 2:12.141 | and the water |  |  |  |
| 2:12.142 | and the fire |  |  |  |
| 2:12.143 | also |  |  |  |
| 2:12.144 | they |  |  |  |
| 2:12.145 | act |  |  |  |
| 2:12.146 | to sprout |  |  |  |
| 2:12.147 | grass |  |  |  |
| 2:12.148 | and trees |  |  |  |
| 2:12.149 | and to raise them |  |  |  |
| 2:12.150 | and with fire |  |  |  |
| 2:12.151 | we |  |  |  |
| 2:12.152 | do |  |  |  |
| 2:12.153 | all |  |  |  |
| 2:12.154 | our needs |  |  |  |
| 2:12.155 | to bake |  |  |  |
| 2:12.156 | and to cook |  |  |  |
| 2:12.157 | all |  |  |  |
| 2:12.158 | our food |  |  |  |
| 2:12.159 | and in most |  |  |  |
| 2:12.160 | the crafts |  |  |  |
| 2:12.161 | we blow |  |  |  |
| 2:12.162 | into the fire |  |  |  |
| 2:12.163 | of charcoal |  |  |  |
| 2:12.164 | and bring out |  |  |  |
| 2:12.165 | a vessel |  |  |  |
| 2:12.166 | for its work |  |  |  |
| 2:12.167 | shall we worship |  |  |  |
| 2:12.168 | (object marker) |  |  |  |
| 2:12.169 | one |  |  |  |
| 2:12.170 | of them |  |  |  |
| 2:12.171 | and His will |  |  |  |
| 2:12.172 | may He be blessed |  |  | ית' = יתברך, placed after 'His will' as the Hebrew places it. |
| 2:12.173 | is to perform |  |  |  |
| 2:12.174 | His actions |  |  |  |
| 2:12.175 | through |  |  |  |
| 2:12.176 | by means of |  |  |  |
| 2:12.177 | agents |  |  |  |
| 2:12.178 | intermediary |  |  |  |
| 2:12.179 | plain |  |  |  |
| 2:12.180 | and clear |  |  |  |
| 2:12.181 | in Scripture |  |  |  |
| 2:12.182 | let swarm |  |  | Jussive in form: 'let the waters swarm' (Gen 1:20). |
| 2:12.183 | the waters |  |  |  |
| 2:12.184 | let bring forth |  |  |  |
| 2:12.185 | the earth |  |  |  |
| 2:12.186 | living |  |  |  |
| 2:12.187 | creature |  |  |  |
| 2:12.188 | let sprout |  |  |  |
| 2:12.189 | the earth |  |  |  |
| 2:12.190 | grass |  |  |  |
| 2:12.191 | and it is said |  |  |  |
| 2:12.192 | making |  |  |  |
| 2:12.193 | His messengers |  |  |  |
| 2:12.194 | winds |  |  |  |
| 2:12.195 | His ministers |  |  |  |
| 2:12.196 | fire |  |  |  |
| 2:12.197 | blazing |  |  |  |
| 2:12.198 | and more |  |  |  |
| 2:12.199 | and more |  |  |  |
| 2:12.200 | in every |  |  |  |
| 2:12.201 | one doing |  |  |  |
| 2:12.202 | its errand |  |  |  |
| 2:12.203 | and yet |  |  |  |
| 2:12.204 | all |  |  |  |
| 2:12.205 | this |  |  |  |
| 2:12.206 | commanded us |  |  |  |
| 2:12.207 | the Name, may He be blessed |  |  | השי"ת = השם יתברך. Rendered 'the Name, may He be blessed' by the default in the decisions (open question 3). |
| 2:12.208 | to worship |  |  |  |
| 2:12.209 | alone |  |  |  |
| 2:12.210 | and warned |  |  |  |
| 2:12.211 | us |  |  |  |
| 2:12.212 | not |  |  |  |
| 2:12.213 | we worship |  |  |  |
| 2:12.214 | any one |  |  |  |
| 2:12.215 | besides Him |  |  |  |
| 2:12.216 | and at the standing |  |  |  |
| 2:12.217 | the chosen |  |  |  |
| 2:12.218 | in His revealing |  |  |  |
| 2:12.219 | on |  |  |  |
| 2:12.220 | the mountain |  |  |  |
| 2:12.221 | Sinai |  |  |  |
| 2:12.222 | He gave |  |  |  |
| 2:12.223 | His voice |  |  |  |
| 2:12.224 | before |  |  |  |
| 2:12.225 | His host |  |  |  |
| 2:12.226 | I am |  |  |  |
| 2:12.227 | the LORD |  |  | ה' is the Tetragrammaton as written, here 'the LORD'. |
| 2:12.228 | your God |  |  |  |
| 2:12.229 | no |  |  |  |
| 2:12.230 | shall be |  |  |  |
| 2:12.231 | to you |  |  |  |
| 2:12.232 | gods |  |  |  |
| 2:12.233 | other |  |  |  |
| 2:12.234 | before |  |  |  |
| 2:12.235 | Me |  |  |  |
| 2:12.236 | and it is said |  |  |  |
| 2:12.237 | I |  |  |  |
| 2:12.238 | the LORD |  |  |  |
| 2:12.239 | that is |  |  |  |
| 2:12.240 | My name |  |  |  |
| 2:12.241 | and My glory |  |  |  |
| 2:12.242 | to another |  |  |  |
| 2:12.243 | not |  |  |  |
| 2:12.244 | I will give |  |  |  |
| 2:12.245 | and I |  |  |  |
| 2:12.246 | and you |  |  |  |
| 2:12.247 | stood |  |  |  |
| 2:12.248 | before |  |  |  |
| 2:12.249 | craftsman |  |  |  |
| 2:12.250 | who made us |  |  |  |
| 2:12.251 | and not |  |  |  |
| 2:12.252 | wished |  |  |  |
| 2:12.253 | He |  |  |  |
| 2:12.254 | that we worship |  |  |  |
| 2:12.255 | god |  |  |  |
| 2:12.256 | another |  |  |  |
| 2:12.257 | quick |  |  | קצר אפים: the author's calque of זעיר אנפין (Aramaic 'small face'), set against ארך אפים, 'slow to anger', in Exod 34:6; compare Prov 14:17. Rendered 'quick to anger' throughout. |
| 2:12.258 | to anger |  |  |  |
| 2:12.259 | Ze'ir Anpin |  |  |  |
| 2:12.260 | apart from Him |  |  |  |
| 2:12.261 | although |  |  | אעפ"י = אף על פי, 'although'. |
| 2:12.262 | that His agent |  |  |  |
| 2:12.263 | is |  |  |  |
| 2:12.264 | to perform |  |  |  |
| 2:12.265 | some |  |  |  |
| 2:12.266 | action |  |  |  |
| 2:12.267 | and why |  |  |  |
| 2:12.268 | not |  |  |  |
| 2:12.269 | we heed |  |  |  |
| 2:12.270 | His voice |  |  |  |
| 2:12.271 | to worship Him |  |  |  |
| 2:12.272 | and go |  |  |  |
| 2:12.273 | and worship |  |  |  |
| 2:12.274 | to Ze'ir Anpin |  |  |  |
| 2:12.275 | which |  |  |  |
| 2:12.276 | not |  |  |  |
| 2:12.277 | He commanded us |  |  |  |

Phrases:
- 2:12.1 2:12.2 2:12.3 2:12.4 | And this I have seen in your reply. | and-(object-marker) this I-have-seen in-your-reply
- 2:12.5 2:12.6 2:12.7 2:12.8 2:12.9 | And I will tell these words that split roofs, | these-words that-split roofs: an Aramaic idiom for forceful words, translated by its sense
- 2:12.10 2:12.11 2:12.12 2:12.13 2:12.14 2:12.15 | since you replied to my words that I said to you. | since you-replied to my-words that-I-said to-you
- 2:12.16 2:12.17 | How shall we set aside | how shall-we-set-aside
- 2:12.18 2:12.19 2:12.20 2:12.21 2:12.22 2:12.23 | Ein Sof and the rest of the partzufim close to Ein Sof | Ein-Sof and-the-rest-of the-partzufim close to Ein-Sof
- 2:12.24 2:12.25 | and worship Ze'ir Anpin? | and-worship to-Ze'ir-Anpin
- 2:12.26 2:12.27 | And your reply was, | and-was your-reply
- 2:12.28 2:12.29 2:12.30 2:12.31 2:12.32 2:12.33 2:12.34 | because the other attributes are some judgment and some mercy, | because the-rest-of the-attributes of-them judgment and-some-of-them mercy
- 2:12.35 2:12.36 2:12.37 | and this one includes them all. | and-this includes all-of-them
- 2:12.38 2:12.39 2:12.40 2:12.41 2:12.42 | And for this reason the Emanator wished to perform His actions | and-for wished the-Emanator to-perform His-actions
- 2:12.43 2:12.44 2:12.45 2:12.46 | through this vessel, etc. | through vessel this etc.
- 2:12.47 2:12.48 2:12.49 | And you brought me an analogy | and-you-brought to-me analogy
- 2:12.50 2:12.51 2:12.52 2:12.53 2:12.54 2:12.55 2:12.56 | from the bodily hand, which does all the work, | from the-hand bodily which does all the-work
- 2:12.57 2:12.58 2:12.59 2:12.60 2:12.61 2:12.62 | while the head, which is above, does nothing. | and-the-head which above does-not does anything
- 2:12.63 2:12.64 2:12.65 2:12.66 2:12.67 2:12.68 | Also the power of speech, which is glory, etc.; | also the-power of-speech which glory etc.
- 2:12.69 2:12.70 2:12.71 2:12.72 2:12.73 2:12.74 2:12.75 | all these words of yours are nothing but words of astonishment! | all your-words these are-not but words astonishment
- 2:12.76 2:12.77 2:12.78 2:12.79 2:12.80 | I asked you about their saying that Ze'ir Anpin | I asked-you about their-saying that-Ze'ir-Anpin
- 2:12.81 2:12.82 2:12.83 2:12.84 | is our God, and we worship him, | is our-God and-him we-worship
- 2:12.85 2:12.86 2:12.87 2:12.88 2:12.89 2:12.90 2:12.91 | and you answer me with those who act in him by the will of their action. | and-you answer-me with those-who-act in-him by-the-will of-their-action
- 2:12.92 2:12.93 2:12.94 2:12.95 2:12.96 2:12.97 2:12.98 2:12.99 | Do we then worship all those who act in him by the will of their action? | and-is-it-that all those-who-act in-him by-the-will of-their-action them shall-we-worship
- 2:12.100 2:12.101 2:12.102 | And shall we say that they are our God? | and-shall-we-say that-they our-God
- 2:12.103 2:12.104 2:12.105 2:12.106 2:12.107 2:12.108 2:12.109 | Surely the sun acts for us in many good actions: | surely the-sun acts for-us many actions good
- 2:12.110 2:12.111 2:12.112 2:12.113 2:12.114 2:12.115 2:12.116 2:12.117 2:12.118 | to shine for us, to warm the air and the earth, to make grass sprout, and to ripen the fruit. | to-shine for-us to-warm the-air and-the-earth and-to-make-sprout grass and-to-ripen the-fruit
- 2:12.119 2:12.120 2:12.121 | Shall we worship her? | whether her we-worship
- 2:12.122 2:12.123 2:12.124 2:12.125 2:12.126 2:12.127 | The moon and the stars too act upon the lower beings | the-moon and-the-stars also they act below
- 2:12.128 2:12.129 2:12.130 | by the will of the Creator, may He be blessed, | by-the-will the-Creator may-He-be-blessed
- 2:12.131 2:12.132 2:12.133 2:12.134 2:12.135 2:12.136 2:12.137 | as it is said: "and from the precious things of the yield of the sun, and from the precious things of the growth of the moons." | as-it-is-said: Deut 33:14 as the author quotes it, the two halves in the verse's order
- 2:12.138 2:12.139 | Shall we worship them? | shall-we-worship them
- 2:12.140 2:12.141 2:12.142 2:12.143 2:12.144 2:12.145 2:12.146 2:12.147 2:12.148 2:12.149 | The earth, the water and the fire also act to sprout grass and trees and to raise them. | the-earth and-the-water and-the-fire also they act to-sprout grass and-trees and-to-raise-them
- 2:12.150 2:12.151 2:12.152 2:12.153 2:12.154 2:12.155 2:12.156 2:12.157 2:12.158 | And with fire we do all our needs: to bake and to cook all our food. | and-with-fire we do all our-needs to-bake and-to-cook all our-food
- 2:12.159 2:12.160 2:12.161 2:12.162 2:12.163 2:12.164 2:12.165 2:12.166 | And in most crafts we blow into a charcoal fire and bring out a vessel for its work. | and-in-most the-crafts we-blow into-the-fire of-charcoal and-bring-out a-vessel for-its-work
- 2:12.167 | Shall we worship | shall-we-worship
- 2:12.168 | - | (object marker): no English counterpart
- 2:12.169 2:12.170 | one of them? | one of-them
- 2:12.171 2:12.172 2:12.173 2:12.174 | And His will, may He be blessed, is to perform His actions | and-His-will may-He-be-blessed is-to-perform His-actions
- 2:12.175 2:12.176 2:12.177 2:12.178 | through intermediary agents, | through by-means-of agents intermediary
- 2:12.179 2:12.180 2:12.181 | as is plain and clear in Scripture. | plain and-clear in-Scripture
- 2:12.182 2:12.183 2:12.184 2:12.185 2:12.186 2:12.187 2:12.188 2:12.189 2:12.190 | "Let the waters swarm, let the earth bring forth a living creature, let the earth sprout grass," | quoted with no formula: Gen 1:20, Gen 1:24, Gen 1:11 in jussive form
- 2:12.191 2:12.192 2:12.193 2:12.194 2:12.195 2:12.196 2:12.197 | and it is said: "He makes His messengers winds, His ministers a blazing fire," | Ps 104:4 as the author quotes it
- 2:12.198 2:12.199 2:12.200 2:12.201 2:12.202 | and more and more, in every case of one doing its errand. | and-more and-more in-every one-doing its-errand
- 2:12.203 2:12.204 2:12.205 | And yet, despite all this, | and-yet all this
- 2:12.206 2:12.207 2:12.208 2:12.209 | the Name, may He be blessed, commanded us to worship Him alone, | commanded-us the-Name,-may-He-be-blessed to-worship alone
- 2:12.210 2:12.211 2:12.212 2:12.213 2:12.214 2:12.215 | and warned us not to worship any one besides Him. | and-warned us not we-worship any-one besides-Him
- 2:12.216 2:12.217 2:12.218 2:12.219 2:12.220 2:12.221 | And at the chosen standing, when He revealed Himself on Mount Sinai, | and-at-the-standing the-chosen in-His-revealing on the-mountain Sinai
- 2:12.222 2:12.223 2:12.224 2:12.225 | He gave His voice before His host: | He-gave His-voice before His-host
- 2:12.226 2:12.227 2:12.228 | "I am the LORD your God … | Exod 20:2 as the author quotes it
- 2:12.229 2:12.230 2:12.231 2:12.232 2:12.233 2:12.234 2:12.235 | You shall have no other gods before Me." | Exod 20:3 as the author quotes it
- 2:12.236 2:12.237 2:12.238 | And it is said: "I am the LORD, | Isa 42:8, introduced by ונאמר
- 2:12.239 2:12.240 | that is My name; | that-is My-name
- 2:12.241 2:12.242 2:12.243 2:12.244 | and My glory I will not give to another!" | and-My-glory to-another not I-will-give
- 2:12.245 2:12.246 2:12.247 2:12.248 | You and I stood before | and-I and-you stood before
- 2:12.249 2:12.250 | the craftsman who made us, | craftsman who-made-us
- 2:12.251 2:12.252 2:12.253 2:12.254 | and He did not wish us to worship | and-not wished He that-we-worship
- 2:12.255 2:12.256 2:12.257 2:12.258 2:12.259 | another god, quick to anger (Ze'ir Anpin), | quick to anger: the author's calque of Ze'ir Anpin, set against 'slow to anger' (Exod 34:6)
- 2:12.260 2:12.261 2:12.262 2:12.263 2:12.264 2:12.265 2:12.266 | apart from Him, although His agent is to perform some action. | apart-from-Him although that-His-agent is to-perform some action
- 2:12.267 2:12.268 2:12.269 2:12.270 2:12.271 | And why do we not heed His voice to worship Him, | and-why not we-heed His-voice to-worship-Him
- 2:12.272 2:12.273 2:12.274 | and go and worship Ze'ir Anpin, | and-go and-worship to-Ze'ir-Anpin
- 2:12.275 2:12.276 2:12.277 | which He did not command us? | which not He-commanded-us

Grammar notes:
- ואת זה חזיתי: the object marker את with the conjunction introduces זה, 'this', ahead of the verb; the English brings it to the front.
- הלא, האם, וכי and the questions ending in הנעבוד and נעבוד are rhetorical questions that expect the answer 'no'.
- ונעבוד, נעבוד, נשמע and נלך are first-person plural forms ('we shall worship', 'we shall heed'), read as deliberative in these questions; the English keeps 'we' in each.
- ישרצו, תוצא and תדשא are jussive in form (Gen 1:20, 1:24, 1:11): 'let the waters swarm', 'let the earth bring forth', 'let the earth sprout'.
- והיתה תשובתך, לפי ששאר המידות ...: the writer summarises the respondent's reply in indirect speech, with no quotation and no quotation marks.

Translator's note: The writer, who asked the question (שאלתיך), speaks to the respondent as 'you' and reports the reply in indirect speech; the reply is not quoted. השי"ת is rendered 'the Name, may He be blessed' by the default in the decisions (open question 3).

Commentary rulings:
- apparatus: Deut 33:14, quoted after 'as it is said'.
- apparatus: Gen 1:20, Gen 1:24 and Gen 1:11, quoted with no formula.
- apparatus: Ps 104:4, introduced by 'and it is said'.
- apparatus: Exod 20:2-3, quoted with no formula; the ellipsis is the fixed rendering of the two verses.
- apparatus: Isa 42:8, introduced by 'and it is said'.
- apparatus: Exod 34:6, the verse that sets 'slow to anger' against 'quick to anger' (see token 2:12.257).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

