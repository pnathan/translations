# Translation pack: milhamot-hashem perek-6-27 (Milchamot ha-Shem 6:41–6:41)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-27.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-27`.

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
- **אדנות** as a name → "Adonai" (note: "lordship").
- **ב"ה** → "blessed be He".
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
- **רמ"א** → "Rema" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
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
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
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


## Milchamot ha-Shem 6:41

Hebrew (Wikisource): עט) ולא די שעשו כן לחשוב כמה אלוהות בעולם האצילות, אלא שגם בעולם הבריאה, ועולם היצירה, ועולם העשייה. המציאו ברעיונותם כל הפרצופים הנז' אצלם בעולם האצילות, ואמר שהאין סוף המתלבש בהם לאו איהו חד בהון כמו בעולם האצילות כמו שכתב הר' כסא אליהו דף כ"א ע"א. ואע"פ שאמרו שהפרצופים שבבי"ע אין האין סוף בהם א' עכ"ז הם משבחים ומפארים להם, ומקבלים עליהם אלהותם כמבואר בסידור חסד חאברהם באמרם אין כאלהינו (מלכות דעשיה) אין כאדונינו (זעיר דעשייה) אין כמלכינו (אימא דעשייה) אין כמושיענו (אבא ואריך דעשייה) מי כאלהינו (מלכות דיצירה) מי כאדונינו (זעיר כיצירה) מי כמלכינו (אימא דיצירה) מי כמושיענו (אבא ואריך דיצירה) נודה לאלהינו (מלכות דבריאה) נודה לאדונינו (סעיר דבריאה) נודה למלכינו (אימא דבריאה) נודה למושיענו (אבא ואריך אנפין דבריאה) ברוך אלהינו (מלכות דאצילות) ברוך אדונינו (זעיר דאצילות) ברוך מלכינו (אימא דאצילות) ברןך מושעינו (אבא ואריך אנפין דאצילות) אצ ההוא אלהינו (כתר דאריך אנפין) אתה הוא אדונינו (ראש עתיק יומין) אתה הוא מלכינו (אדם קדמון) אתה הוא מושיענו (אור פנימי ואור מקיף המתלבש בו) כמו שיראה המעיין כל זה בסידור חסד לאברהם. הרי לך בהדיא שהם מקבלים עליהם מלכות ואדנות ואלהות כל פרצוף ופרצוף. אפי' פרצופים שבעולמות הפירוד לפי דבריהם שהם בריאה, יצירה ועשייה. ופליאה גדולה אנו מתפלאים עליכם שאתם נלחמים עמנו בכל המלחמה הזאת לבזות ולרדוף ולירד עמנו עד לחיינו. ואנו סובלים כל חרפות ובזוים לכבוד אבינו מלכינו יחיד ומיוחד מכל האחדים, בעבור מה? לא בעבור שנקבל ונודה לאמונתכם הלזו ש]ירשנו ממנה ולא חירפנו ולא בזינו אותה? ועם כל זה אתם מסתירים ומכסים את אמונתכם שלא לפרסמה. ומראים בתשובתכם כאלו אינכם מודים בה כעין גנב שמרה נשף לאמר לא תשורני עין. וסתר פנים תשימו לה! לכן הוכרחנו להביא כל זה להוכיח את הכחשות שאתם מכחישים אותנו בתשובתכם ואומרים ומי הוא שאמר שהאין סוף מתחלק? כל זה הוצרכנו להעתיק מספרי המקובלים החדשים. ולהביא מאמרי רז"ל נגדן להוכיח טעותכם או רמאותכם שאתם מסתירים איכות אמונתכם ומכסים אותה בטלית. ומכחישים את הידוע, ואינכם מבקשים האמת. ולמה זה תבושו לומר האמת שכן קבלנו לעבוד אלוהות רבים נבראים רק שאתם רודפים אחר הכבוד המדומה. לקנות שם נגד ההמון אשר לא ידעו ולא יבינו, בחשכה יתהלכו, ולשבח את עצמכם נגדם כאלו אתם נלחמים עם ה' ועם תורתו.

| id | word | punctuation after |
|---|---|---|
| 6:41.1 | עט | ) |
| 6:41.2 | ולא |  |
| 6:41.3 | די |  |
| 6:41.4 | שעשו |  |
| 6:41.5 | כן |  |
| 6:41.6 | לחשוב |  |
| 6:41.7 | כמה |  |
| 6:41.8 | אלוהות |  |
| 6:41.9 | בעולם |  |
| 6:41.10 | האצילות | , |
| 6:41.11 | אלא |  |
| 6:41.12 | שגם |  |
| 6:41.13 | בעולם |  |
| 6:41.14 | הבריאה | , |
| 6:41.15 | ועולם |  |
| 6:41.16 | היצירה | , |
| 6:41.17 | ועולם |  |
| 6:41.18 | העשייה | . |
| 6:41.19 | המציאו |  |
| 6:41.20 | ברעיונותם |  |
| 6:41.21 | כל |  |
| 6:41.22 | הפרצופים |  |
| 6:41.23 | הנז | ' |
| 6:41.24 | אצלם |  |
| 6:41.25 | בעולם |  |
| 6:41.26 | האצילות | , |
| 6:41.27 | ואמר |  |
| 6:41.28 | שהאין |  |
| 6:41.29 | סוף |  |
| 6:41.30 | המתלבש |  |
| 6:41.31 | בהם |  |
| 6:41.32 | לאו |  |
| 6:41.33 | איהו |  |
| 6:41.34 | חד |  |
| 6:41.35 | בהון |  |
| 6:41.36 | כמו |  |
| 6:41.37 | בעולם |  |
| 6:41.38 | האצילות |  |
| 6:41.39 | כמו |  |
| 6:41.40 | שכתב |  |
| 6:41.41 | הר | ' |
| 6:41.42 | כסא |  |
| 6:41.43 | אליהו |  |
| 6:41.44 | דף |  |
| 6:41.45 | כ"א |  |
| 6:41.46 | ע"א | . |
| 6:41.47 | ואע"פ |  |
| 6:41.48 | שאמרו |  |
| 6:41.49 | שהפרצופים |  |
| 6:41.50 | שבבי"ע |  |
| 6:41.51 | אין |  |
| 6:41.52 | האין |  |
| 6:41.53 | סוף |  |
| 6:41.54 | בהם |  |
| 6:41.55 | א | ' |
| 6:41.56 | עכ"ז |  |
| 6:41.57 | הם |  |
| 6:41.58 | משבחים |  |
| 6:41.59 | ומפארים |  |
| 6:41.60 | להם | , |
| 6:41.61 | ומקבלים |  |
| 6:41.62 | עליהם |  |
| 6:41.63 | אלהותם |  |
| 6:41.64 | כמבואר |  |
| 6:41.65 | בסידור |  |
| 6:41.66 | חסד |  |
| 6:41.67 | חאברהם |  |
| 6:41.68 | באמרם |  |
| 6:41.69 | אין |  |
| 6:41.70 | כאלהינו |  |
| 6:41.71 | מלכות |  |
| 6:41.72 | דעשיה | ) |
| 6:41.73 | אין |  |
| 6:41.74 | כאדונינו |  |
| 6:41.75 | זעיר |  |
| 6:41.76 | דעשייה | ) |
| 6:41.77 | אין |  |
| 6:41.78 | כמלכינו |  |
| 6:41.79 | אימא |  |
| 6:41.80 | דעשייה | ) |
| 6:41.81 | אין |  |
| 6:41.82 | כמושיענו |  |
| 6:41.83 | אבא |  |
| 6:41.84 | ואריך |  |
| 6:41.85 | דעשייה | ) |
| 6:41.86 | מי |  |
| 6:41.87 | כאלהינו |  |
| 6:41.88 | מלכות |  |
| 6:41.89 | דיצירה | ) |
| 6:41.90 | מי |  |
| 6:41.91 | כאדונינו |  |
| 6:41.92 | זעיר |  |
| 6:41.93 | כיצירה | ) |
| 6:41.94 | מי |  |
| 6:41.95 | כמלכינו |  |
| 6:41.96 | אימא |  |
| 6:41.97 | דיצירה | ) |
| 6:41.98 | מי |  |
| 6:41.99 | כמושיענו |  |
| 6:41.100 | אבא |  |
| 6:41.101 | ואריך |  |
| 6:41.102 | דיצירה | ) |
| 6:41.103 | נודה |  |
| 6:41.104 | לאלהינו |  |
| 6:41.105 | מלכות |  |
| 6:41.106 | דבריאה | ) |
| 6:41.107 | נודה |  |
| 6:41.108 | לאדונינו |  |
| 6:41.109 | סעיר |  |
| 6:41.110 | דבריאה | ) |
| 6:41.111 | נודה |  |
| 6:41.112 | למלכינו |  |
| 6:41.113 | אימא |  |
| 6:41.114 | דבריאה | ) |
| 6:41.115 | נודה |  |
| 6:41.116 | למושיענו |  |
| 6:41.117 | אבא |  |
| 6:41.118 | ואריך |  |
| 6:41.119 | אנפין |  |
| 6:41.120 | דבריאה | ) |
| 6:41.121 | ברוך |  |
| 6:41.122 | אלהינו |  |
| 6:41.123 | מלכות |  |
| 6:41.124 | דאצילות | ) |
| 6:41.125 | ברוך |  |
| 6:41.126 | אדונינו |  |
| 6:41.127 | זעיר |  |
| 6:41.128 | דאצילות | ) |
| 6:41.129 | ברוך |  |
| 6:41.130 | מלכינו |  |
| 6:41.131 | אימא |  |
| 6:41.132 | דאצילות | ) |
| 6:41.133 | ברןך |  |
| 6:41.134 | מושעינו |  |
| 6:41.135 | אבא |  |
| 6:41.136 | ואריך |  |
| 6:41.137 | אנפין |  |
| 6:41.138 | דאצילות | ) |
| 6:41.139 | אצ |  |
| 6:41.140 | ההוא |  |
| 6:41.141 | אלהינו |  |
| 6:41.142 | כתר |  |
| 6:41.143 | דאריך |  |
| 6:41.144 | אנפין | ) |
| 6:41.145 | אתה |  |
| 6:41.146 | הוא |  |
| 6:41.147 | אדונינו |  |
| 6:41.148 | ראש |  |
| 6:41.149 | עתיק |  |
| 6:41.150 | יומין | ) |
| 6:41.151 | אתה |  |
| 6:41.152 | הוא |  |
| 6:41.153 | מלכינו |  |
| 6:41.154 | אדם |  |
| 6:41.155 | קדמון | ) |
| 6:41.156 | אתה |  |
| 6:41.157 | הוא |  |
| 6:41.158 | מושיענו |  |
| 6:41.159 | אור |  |
| 6:41.160 | פנימי |  |
| 6:41.161 | ואור |  |
| 6:41.162 | מקיף |  |
| 6:41.163 | המתלבש |  |
| 6:41.164 | בו | ) |
| 6:41.165 | כמו |  |
| 6:41.166 | שיראה |  |
| 6:41.167 | המעיין |  |
| 6:41.168 | כל |  |
| 6:41.169 | זה |  |
| 6:41.170 | בסידור |  |
| 6:41.171 | חסד |  |
| 6:41.172 | לאברהם | . |
| 6:41.173 | הרי |  |
| 6:41.174 | לך |  |
| 6:41.175 | בהדיא |  |
| 6:41.176 | שהם |  |
| 6:41.177 | מקבלים |  |
| 6:41.178 | עליהם |  |
| 6:41.179 | מלכות |  |
| 6:41.180 | ואדנות |  |
| 6:41.181 | ואלהות |  |
| 6:41.182 | כל |  |
| 6:41.183 | פרצוף |  |
| 6:41.184 | ופרצוף | . |
| 6:41.185 | אפי | ' |
| 6:41.186 | פרצופים |  |
| 6:41.187 | שבעולמות |  |
| 6:41.188 | הפירוד |  |
| 6:41.189 | לפי |  |
| 6:41.190 | דבריהם |  |
| 6:41.191 | שהם |  |
| 6:41.192 | בריאה | , |
| 6:41.193 | יצירה |  |
| 6:41.194 | ועשייה | . |
| 6:41.195 | ופליאה |  |
| 6:41.196 | גדולה |  |
| 6:41.197 | אנו |  |
| 6:41.198 | מתפלאים |  |
| 6:41.199 | עליכם |  |
| 6:41.200 | שאתם |  |
| 6:41.201 | נלחמים |  |
| 6:41.202 | עמנו |  |
| 6:41.203 | בכל |  |
| 6:41.204 | המלחמה |  |
| 6:41.205 | הזאת |  |
| 6:41.206 | לבזות |  |
| 6:41.207 | ולרדוף |  |
| 6:41.208 | ולירד |  |
| 6:41.209 | עמנו |  |
| 6:41.210 | עד |  |
| 6:41.211 | לחיינו | . |
| 6:41.212 | ואנו |  |
| 6:41.213 | סובלים |  |
| 6:41.214 | כל |  |
| 6:41.215 | חרפות |  |
| 6:41.216 | ובזוים |  |
| 6:41.217 | לכבוד |  |
| 6:41.218 | אבינו |  |
| 6:41.219 | מלכינו |  |
| 6:41.220 | יחיד |  |
| 6:41.221 | ומיוחד |  |
| 6:41.222 | מכל |  |
| 6:41.223 | האחדים | , |
| 6:41.224 | בעבור |  |
| 6:41.225 | מה | ? |
| 6:41.226 | לא |  |
| 6:41.227 | בעבור |  |
| 6:41.228 | שנקבל |  |
| 6:41.229 | ונודה |  |
| 6:41.230 | לאמונתכם |  |
| 6:41.231 | הלזו |  |
| 6:41.232 | ש]ירשנו |  |
| 6:41.233 | ממנה |  |
| 6:41.234 | ולא |  |
| 6:41.235 | חירפנו |  |
| 6:41.236 | ולא |  |
| 6:41.237 | בזינו |  |
| 6:41.238 | אותה | ? |
| 6:41.239 | ועם |  |
| 6:41.240 | כל |  |
| 6:41.241 | זה |  |
| 6:41.242 | אתם |  |
| 6:41.243 | מסתירים |  |
| 6:41.244 | ומכסים |  |
| 6:41.245 | את |  |
| 6:41.246 | אמונתכם |  |
| 6:41.247 | שלא |  |
| 6:41.248 | לפרסמה | . |
| 6:41.249 | ומראים |  |
| 6:41.250 | בתשובתכם |  |
| 6:41.251 | כאלו |  |
| 6:41.252 | אינכם |  |
| 6:41.253 | מודים |  |
| 6:41.254 | בה |  |
| 6:41.255 | כעין |  |
| 6:41.256 | גנב |  |
| 6:41.257 | שמרה |  |
| 6:41.258 | נשף |  |
| 6:41.259 | לאמר |  |
| 6:41.260 | לא |  |
| 6:41.261 | תשורני |  |
| 6:41.262 | עין | . |
| 6:41.263 | וסתר |  |
| 6:41.264 | פנים |  |
| 6:41.265 | תשימו |  |
| 6:41.266 | לה | ! |
| 6:41.267 | לכן |  |
| 6:41.268 | הוכרחנו |  |
| 6:41.269 | להביא |  |
| 6:41.270 | כל |  |
| 6:41.271 | זה |  |
| 6:41.272 | להוכיח |  |
| 6:41.273 | את |  |
| 6:41.274 | הכחשות |  |
| 6:41.275 | שאתם |  |
| 6:41.276 | מכחישים |  |
| 6:41.277 | אותנו |  |
| 6:41.278 | בתשובתכם |  |
| 6:41.279 | ואומרים |  |
| 6:41.280 | ומי |  |
| 6:41.281 | הוא |  |
| 6:41.282 | שאמר |  |
| 6:41.283 | שהאין |  |
| 6:41.284 | סוף |  |
| 6:41.285 | מתחלק | ? |
| 6:41.286 | כל |  |
| 6:41.287 | זה |  |
| 6:41.288 | הוצרכנו |  |
| 6:41.289 | להעתיק |  |
| 6:41.290 | מספרי |  |
| 6:41.291 | המקובלים |  |
| 6:41.292 | החדשים | . |
| 6:41.293 | ולהביא |  |
| 6:41.294 | מאמרי |  |
| 6:41.295 | רז"ל |  |
| 6:41.296 | נגדן |  |
| 6:41.297 | להוכיח |  |
| 6:41.298 | טעותכם |  |
| 6:41.299 | או |  |
| 6:41.300 | רמאותכם |  |
| 6:41.301 | שאתם |  |
| 6:41.302 | מסתירים |  |
| 6:41.303 | איכות |  |
| 6:41.304 | אמונתכם |  |
| 6:41.305 | ומכסים |  |
| 6:41.306 | אותה |  |
| 6:41.307 | בטלית | . |
| 6:41.308 | ומכחישים |  |
| 6:41.309 | את |  |
| 6:41.310 | הידוע | , |
| 6:41.311 | ואינכם |  |
| 6:41.312 | מבקשים |  |
| 6:41.313 | האמת | . |
| 6:41.314 | ולמה |  |
| 6:41.315 | זה |  |
| 6:41.316 | תבושו |  |
| 6:41.317 | לומר |  |
| 6:41.318 | האמת |  |
| 6:41.319 | שכן |  |
| 6:41.320 | קבלנו |  |
| 6:41.321 | לעבוד |  |
| 6:41.322 | אלוהות |  |
| 6:41.323 | רבים |  |
| 6:41.324 | נבראים |  |
| 6:41.325 | רק |  |
| 6:41.326 | שאתם |  |
| 6:41.327 | רודפים |  |
| 6:41.328 | אחר |  |
| 6:41.329 | הכבוד |  |
| 6:41.330 | המדומה | . |
| 6:41.331 | לקנות |  |
| 6:41.332 | שם |  |
| 6:41.333 | נגד |  |
| 6:41.334 | ההמון |  |
| 6:41.335 | אשר |  |
| 6:41.336 | לא |  |
| 6:41.337 | ידעו |  |
| 6:41.338 | ולא |  |
| 6:41.339 | יבינו | , |
| 6:41.340 | בחשכה |  |
| 6:41.341 | יתהלכו | , |
| 6:41.342 | ולשבח |  |
| 6:41.343 | את |  |
| 6:41.344 | עצמכם |  |
| 6:41.345 | נגדם |  |
| 6:41.346 | כאלו |  |
| 6:41.347 | אתם |  |
| 6:41.348 | נלחמים |  |
| 6:41.349 | עם |  |
| 6:41.350 | ה | ' |
| 6:41.351 | ועם |  |
| 6:41.352 | תורתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

