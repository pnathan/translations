# Translation pack: milhamot-hashem perek-9-05 (Milchamot ha-Shem 9:5–9:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-05`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **האלהי** before a name → "the divine" (note); never dropped.
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

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
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- גוף **דק** → "a subtle body".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **גבורה** as a name of God → "the Power".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- הבל **וריק** → "vanity and emptiness".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
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


## Milchamot ha-Shem 9:5

Hebrew (Wikisource): קכד) ועל הירא את דבר ה' ומשכיל בתורתו ובקבלת רבותינו במשנה ובתלמוד אל יפת בסתר לבו לעבוד שום צורה ופרצוף אשר לפי דבריהם חלק להם האלוה העליון מעצמותו וישם את חלק החכמה שבו נשמה לחכמה ואת חלק הבינה שיש בו נשמה לבינה הנקראת אימא ואת חלק הגדולה שבו שם בגדולה ואת חלק הגבורה שם בגבורה וכו'. ויהיו החלקים האלה להם לנשמות וגם את הגופים הרקים העליונים שמו אותם לנשמות לגופים בתחתונים מהם וקיימו וקבלו עליהם לעבדם ואמרו שבעברינו אותם אנחנו עובדים לנשמתם שהיא הסבה הראשונה ממש.

| id | word | punctuation after |
|---|---|---|
| 9:5.1 | קכד | ) |
| 9:5.2 | ועל |  |
| 9:5.3 | הירא |  |
| 9:5.4 | את |  |
| 9:5.5 | דבר |  |
| 9:5.6 | ה | ' |
| 9:5.7 | ומשכיל |  |
| 9:5.8 | בתורתו |  |
| 9:5.9 | ובקבלת |  |
| 9:5.10 | רבותינו |  |
| 9:5.11 | במשנה |  |
| 9:5.12 | ובתלמוד |  |
| 9:5.13 | אל |  |
| 9:5.14 | יפת |  |
| 9:5.15 | בסתר |  |
| 9:5.16 | לבו |  |
| 9:5.17 | לעבוד |  |
| 9:5.18 | שום |  |
| 9:5.19 | צורה |  |
| 9:5.20 | ופרצוף |  |
| 9:5.21 | אשר |  |
| 9:5.22 | לפי |  |
| 9:5.23 | דבריהם |  |
| 9:5.24 | חלק |  |
| 9:5.25 | להם |  |
| 9:5.26 | האלוה |  |
| 9:5.27 | העליון |  |
| 9:5.28 | מעצמותו |  |
| 9:5.29 | וישם |  |
| 9:5.30 | את |  |
| 9:5.31 | חלק |  |
| 9:5.32 | החכמה |  |
| 9:5.33 | שבו |  |
| 9:5.34 | נשמה |  |
| 9:5.35 | לחכמה |  |
| 9:5.36 | ואת |  |
| 9:5.37 | חלק |  |
| 9:5.38 | הבינה |  |
| 9:5.39 | שיש |  |
| 9:5.40 | בו |  |
| 9:5.41 | נשמה |  |
| 9:5.42 | לבינה |  |
| 9:5.43 | הנקראת |  |
| 9:5.44 | אימא |  |
| 9:5.45 | ואת |  |
| 9:5.46 | חלק |  |
| 9:5.47 | הגדולה |  |
| 9:5.48 | שבו |  |
| 9:5.49 | שם |  |
| 9:5.50 | בגדולה |  |
| 9:5.51 | ואת |  |
| 9:5.52 | חלק |  |
| 9:5.53 | הגבורה |  |
| 9:5.54 | שם |  |
| 9:5.55 | בגבורה |  |
| 9:5.56 | וכו | '. |
| 9:5.57 | ויהיו |  |
| 9:5.58 | החלקים |  |
| 9:5.59 | האלה |  |
| 9:5.60 | להם |  |
| 9:5.61 | לנשמות |  |
| 9:5.62 | וגם |  |
| 9:5.63 | את |  |
| 9:5.64 | הגופים |  |
| 9:5.65 | הרקים |  |
| 9:5.66 | העליונים |  |
| 9:5.67 | שמו |  |
| 9:5.68 | אותם |  |
| 9:5.69 | לנשמות |  |
| 9:5.70 | לגופים |  |
| 9:5.71 | בתחתונים |  |
| 9:5.72 | מהם |  |
| 9:5.73 | וקיימו |  |
| 9:5.74 | וקבלו |  |
| 9:5.75 | עליהם |  |
| 9:5.76 | לעבדם |  |
| 9:5.77 | ואמרו |  |
| 9:5.78 | שבעברינו |  |
| 9:5.79 | אותם |  |
| 9:5.80 | אנחנו |  |
| 9:5.81 | עובדים |  |
| 9:5.82 | לנשמתם |  |
| 9:5.83 | שהיא |  |
| 9:5.84 | הסבה |  |
| 9:5.85 | הראשונה |  |
| 9:5.86 | ממש | . |


## Milchamot ha-Shem 9:6

Hebrew (Wikisource): והדעה הזאת לקוחה מהפילסופים החיצונים (שהזכירם הר' המורה פרק כ"ט לשלישי) שאמרו שהאלוה הוא נשמת הגלגלים והכוכבים והמזלות ולכן היו עובדים להם ואומרים שהם עובדים לנשמתם שהיא הסבה הראשונה לדבריהם. וברעיונותיו המציא את הפרצופים האלה מאדם קדמאה עד זעיר ונוקביה ושלשל אותם מהסבה הראשונה אשר צמצמה ברעיונו וסלקה למעלה וחזר ושלש ממנה כל אותן הפרצופים בעולמות אין קץ וגם בעולמות שלמטה מעולם האצילות. וישם את החוט המשתלשל בהם נשמה להם ויצו את קוראיו לבדם בטענת שלא את גופם עובדים רק לנשמתם כמו שעשו הפילסופים הראשונים החיצונים בכוכבים ובמזלות. ואת כל הפרצופים קרא בשם מעשה הרכבה וישם את חיות המרכבה אלהות בעבור החלק המשתלשל מהאלהים מאלהים המצומק אל הסביבות לפי דבריו. אשר מהחלק הזה היורד דרך צנור דק נתהוו כל הגופות עם עיגוליהם וממנו גם נשמותיהן ומלבושיהן ויהיו להם לאלוהות ויקראו שמותיהן בשמותיו של הקב"ה וכינוייו. והאלהים העליון לא השאירו לו לא שם ולא כינוי ועשוהו מסולק וריקן מכל שבח ותהלה וברכה לפי שאין לו ממי להתברך וגזרו אומר שהמתפלל אליו אין תפלתו תפלה ולא יענהו.

| id | word | punctuation after |
|---|---|---|
| 9:6.1 | והדעה |  |
| 9:6.2 | הזאת |  |
| 9:6.3 | לקוחה |  |
| 9:6.4 | מהפילסופים |  |
| 9:6.5 | החיצונים |  |
| 9:6.6 | שהזכירם |  |
| 9:6.7 | הר | ' |
| 9:6.8 | המורה |  |
| 9:6.9 | פרק |  |
| 9:6.10 | כ"ט |  |
| 9:6.11 | לשלישי | ) |
| 9:6.12 | שאמרו |  |
| 9:6.13 | שהאלוה |  |
| 9:6.14 | הוא |  |
| 9:6.15 | נשמת |  |
| 9:6.16 | הגלגלים |  |
| 9:6.17 | והכוכבים |  |
| 9:6.18 | והמזלות |  |
| 9:6.19 | ולכן |  |
| 9:6.20 | היו |  |
| 9:6.21 | עובדים |  |
| 9:6.22 | להם |  |
| 9:6.23 | ואומרים |  |
| 9:6.24 | שהם |  |
| 9:6.25 | עובדים |  |
| 9:6.26 | לנשמתם |  |
| 9:6.27 | שהיא |  |
| 9:6.28 | הסבה |  |
| 9:6.29 | הראשונה |  |
| 9:6.30 | לדבריהם | . |
| 9:6.31 | וברעיונותיו |  |
| 9:6.32 | המציא |  |
| 9:6.33 | את |  |
| 9:6.34 | הפרצופים |  |
| 9:6.35 | האלה |  |
| 9:6.36 | מאדם |  |
| 9:6.37 | קדמאה |  |
| 9:6.38 | עד |  |
| 9:6.39 | זעיר |  |
| 9:6.40 | ונוקביה |  |
| 9:6.41 | ושלשל |  |
| 9:6.42 | אותם |  |
| 9:6.43 | מהסבה |  |
| 9:6.44 | הראשונה |  |
| 9:6.45 | אשר |  |
| 9:6.46 | צמצמה |  |
| 9:6.47 | ברעיונו |  |
| 9:6.48 | וסלקה |  |
| 9:6.49 | למעלה |  |
| 9:6.50 | וחזר |  |
| 9:6.51 | ושלש |  |
| 9:6.52 | ממנה |  |
| 9:6.53 | כל |  |
| 9:6.54 | אותן |  |
| 9:6.55 | הפרצופים |  |
| 9:6.56 | בעולמות |  |
| 9:6.57 | אין |  |
| 9:6.58 | קץ |  |
| 9:6.59 | וגם |  |
| 9:6.60 | בעולמות |  |
| 9:6.61 | שלמטה |  |
| 9:6.62 | מעולם |  |
| 9:6.63 | האצילות | . |
| 9:6.64 | וישם |  |
| 9:6.65 | את |  |
| 9:6.66 | החוט |  |
| 9:6.67 | המשתלשל |  |
| 9:6.68 | בהם |  |
| 9:6.69 | נשמה |  |
| 9:6.70 | להם |  |
| 9:6.71 | ויצו |  |
| 9:6.72 | את |  |
| 9:6.73 | קוראיו |  |
| 9:6.74 | לבדם |  |
| 9:6.75 | בטענת |  |
| 9:6.76 | שלא |  |
| 9:6.77 | את |  |
| 9:6.78 | גופם |  |
| 9:6.79 | עובדים |  |
| 9:6.80 | רק |  |
| 9:6.81 | לנשמתם |  |
| 9:6.82 | כמו |  |
| 9:6.83 | שעשו |  |
| 9:6.84 | הפילסופים |  |
| 9:6.85 | הראשונים |  |
| 9:6.86 | החיצונים |  |
| 9:6.87 | בכוכבים |  |
| 9:6.88 | ובמזלות | . |
| 9:6.89 | ואת |  |
| 9:6.90 | כל |  |
| 9:6.91 | הפרצופים |  |
| 9:6.92 | קרא |  |
| 9:6.93 | בשם |  |
| 9:6.94 | מעשה |  |
| 9:6.95 | הרכבה |  |
| 9:6.96 | וישם |  |
| 9:6.97 | את |  |
| 9:6.98 | חיות |  |
| 9:6.99 | המרכבה |  |
| 9:6.100 | אלהות |  |
| 9:6.101 | בעבור |  |
| 9:6.102 | החלק |  |
| 9:6.103 | המשתלשל |  |
| 9:6.104 | מהאלהים |  |
| 9:6.105 | מאלהים |  |
| 9:6.106 | המצומק |  |
| 9:6.107 | אל |  |
| 9:6.108 | הסביבות |  |
| 9:6.109 | לפי |  |
| 9:6.110 | דבריו | . |
| 9:6.111 | אשר |  |
| 9:6.112 | מהחלק |  |
| 9:6.113 | הזה |  |
| 9:6.114 | היורד |  |
| 9:6.115 | דרך |  |
| 9:6.116 | צנור |  |
| 9:6.117 | דק |  |
| 9:6.118 | נתהוו |  |
| 9:6.119 | כל |  |
| 9:6.120 | הגופות |  |
| 9:6.121 | עם |  |
| 9:6.122 | עיגוליהם |  |
| 9:6.123 | וממנו |  |
| 9:6.124 | גם |  |
| 9:6.125 | נשמותיהן |  |
| 9:6.126 | ומלבושיהן |  |
| 9:6.127 | ויהיו |  |
| 9:6.128 | להם |  |
| 9:6.129 | לאלוהות |  |
| 9:6.130 | ויקראו |  |
| 9:6.131 | שמותיהן |  |
| 9:6.132 | בשמותיו |  |
| 9:6.133 | של |  |
| 9:6.134 | הקב"ה |  |
| 9:6.135 | וכינוייו | . |
| 9:6.136 | והאלהים |  |
| 9:6.137 | העליון |  |
| 9:6.138 | לא |  |
| 9:6.139 | השאירו |  |
| 9:6.140 | לו |  |
| 9:6.141 | לא |  |
| 9:6.142 | שם |  |
| 9:6.143 | ולא |  |
| 9:6.144 | כינוי |  |
| 9:6.145 | ועשוהו |  |
| 9:6.146 | מסולק |  |
| 9:6.147 | וריקן |  |
| 9:6.148 | מכל |  |
| 9:6.149 | שבח |  |
| 9:6.150 | ותהלה |  |
| 9:6.151 | וברכה |  |
| 9:6.152 | לפי |  |
| 9:6.153 | שאין |  |
| 9:6.154 | לו |  |
| 9:6.155 | ממי |  |
| 9:6.156 | להתברך |  |
| 9:6.157 | וגזרו |  |
| 9:6.158 | אומר |  |
| 9:6.159 | שהמתפלל |  |
| 9:6.160 | אליו |  |
| 9:6.161 | אין |  |
| 9:6.162 | תפלתו |  |
| 9:6.163 | תפלה |  |
| 9:6.164 | ולא |  |
| 9:6.165 | יענהו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

