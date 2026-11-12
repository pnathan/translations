# Translation pack: milhamot-hashem perek-2-05 (Milchamot ha-Shem 2:10–2:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-05`.

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
- **חז"ל** → "our Sages of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **ריב"ש** → "Ribash" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ראב"ד** → "Ravad" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
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

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **השתלשלות** → "the chain of descent".
- **עלה / עלול** → "cause / effect".
- גוף **דק** → "a subtle body".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **קורין** אל → "call upon".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
- **רחמנא ליצלן** → "the Merciful One save us".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- תהו **ובהו** → "formlessness and void".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **בנדון** דידן → "in our case".
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


## Milchamot ha-Shem 2:10

Hebrew (Wikisource): ובירושלמי שבת פ' במה אשה אמרו חז"ל וריויה די רביעאה דמי לבר אלהין. אמר ר' ראובן: באותה שעה ירד מלאך וסטרו לההוא רשיעא על פיו ואמר ליה תקין מילך ובר אית ליה? חזר ואמר ברוך אלההון די שדרך מישך ועביד נגו די שלה מלאכיה ושיזיב לעבדוהי, בריה אין כתיב כאן אלא מלאכיה, ע"כ. ור"ל שנבוכדנאצר הרשע חשב שהמלאכים הם עלולים ומשתלשלים ממנו בטבע כבן מן האב כדברי המקובלים בספירות של כל העולמות, וכן הנשמות, ולזה סטרו המלאך. וחזר ואמר די שלח מלאכיה וכו'. הרי לך מפורש יוצא מדברי מדרש רבה והירושלמי היפך דברי המקובלים. ושכל הנמצאים המציאם הקב"ה בחסדו יש מאין מהצורה ראשונה עד יתוש קטן. ולא שהם נמשכים מעצמותו כמו שהבן מתהוה מכח הזרע הנמשך מאביו.

| id | word | punctuation after |
|---|---|---|
| 2:10.1 | ובירושלמי |  |
| 2:10.2 | שבת |  |
| 2:10.3 | פ | ' |
| 2:10.4 | במה |  |
| 2:10.5 | אשה |  |
| 2:10.6 | אמרו |  |
| 2:10.7 | חז"ל |  |
| 2:10.8 | וריויה |  |
| 2:10.9 | די |  |
| 2:10.10 | רביעאה |  |
| 2:10.11 | דמי |  |
| 2:10.12 | לבר |  |
| 2:10.13 | אלהין | . |
| 2:10.14 | אמר |  |
| 2:10.15 | ר | ' |
| 2:10.16 | ראובן | : |
| 2:10.17 | באותה |  |
| 2:10.18 | שעה |  |
| 2:10.19 | ירד |  |
| 2:10.20 | מלאך |  |
| 2:10.21 | וסטרו |  |
| 2:10.22 | לההוא |  |
| 2:10.23 | רשיעא |  |
| 2:10.24 | על |  |
| 2:10.25 | פיו |  |
| 2:10.26 | ואמר |  |
| 2:10.27 | ליה |  |
| 2:10.28 | תקין |  |
| 2:10.29 | מילך |  |
| 2:10.30 | ובר |  |
| 2:10.31 | אית |  |
| 2:10.32 | ליה | ? |
| 2:10.33 | חזר |  |
| 2:10.34 | ואמר |  |
| 2:10.35 | ברוך |  |
| 2:10.36 | אלההון |  |
| 2:10.37 | די |  |
| 2:10.38 | שדרך |  |
| 2:10.39 | מישך |  |
| 2:10.40 | ועביד |  |
| 2:10.41 | נגו |  |
| 2:10.42 | די |  |
| 2:10.43 | שלה |  |
| 2:10.44 | מלאכיה |  |
| 2:10.45 | ושיזיב |  |
| 2:10.46 | לעבדוהי | , |
| 2:10.47 | בריה |  |
| 2:10.48 | אין |  |
| 2:10.49 | כתיב |  |
| 2:10.50 | כאן |  |
| 2:10.51 | אלא |  |
| 2:10.52 | מלאכיה | , |
| 2:10.53 | ע"כ | . |
| 2:10.54 | ור"ל |  |
| 2:10.55 | שנבוכדנאצר |  |
| 2:10.56 | הרשע |  |
| 2:10.57 | חשב |  |
| 2:10.58 | שהמלאכים |  |
| 2:10.59 | הם |  |
| 2:10.60 | עלולים |  |
| 2:10.61 | ומשתלשלים |  |
| 2:10.62 | ממנו |  |
| 2:10.63 | בטבע |  |
| 2:10.64 | כבן |  |
| 2:10.65 | מן |  |
| 2:10.66 | האב |  |
| 2:10.67 | כדברי |  |
| 2:10.68 | המקובלים |  |
| 2:10.69 | בספירות |  |
| 2:10.70 | של |  |
| 2:10.71 | כל |  |
| 2:10.72 | העולמות | , |
| 2:10.73 | וכן |  |
| 2:10.74 | הנשמות | , |
| 2:10.75 | ולזה |  |
| 2:10.76 | סטרו |  |
| 2:10.77 | המלאך | . |
| 2:10.78 | וחזר |  |
| 2:10.79 | ואמר |  |
| 2:10.80 | די |  |
| 2:10.81 | שלח |  |
| 2:10.82 | מלאכיה |  |
| 2:10.83 | וכו | '. |
| 2:10.84 | הרי |  |
| 2:10.85 | לך |  |
| 2:10.86 | מפורש |  |
| 2:10.87 | יוצא |  |
| 2:10.88 | מדברי |  |
| 2:10.89 | מדרש |  |
| 2:10.90 | רבה |  |
| 2:10.91 | והירושלמי |  |
| 2:10.92 | היפך |  |
| 2:10.93 | דברי |  |
| 2:10.94 | המקובלים | . |
| 2:10.95 | ושכל |  |
| 2:10.96 | הנמצאים |  |
| 2:10.97 | המציאם |  |
| 2:10.98 | הקב"ה |  |
| 2:10.99 | בחסדו |  |
| 2:10.100 | יש |  |
| 2:10.101 | מאין |  |
| 2:10.102 | מהצורה |  |
| 2:10.103 | ראשונה |  |
| 2:10.104 | עד |  |
| 2:10.105 | יתוש |  |
| 2:10.106 | קטן | . |
| 2:10.107 | ולא |  |
| 2:10.108 | שהם |  |
| 2:10.109 | נמשכים |  |
| 2:10.110 | מעצמותו |  |
| 2:10.111 | כמו |  |
| 2:10.112 | שהבן |  |
| 2:10.113 | מתהוה |  |
| 2:10.114 | מכח |  |
| 2:10.115 | הזרע |  |
| 2:10.116 | הנמשך |  |
| 2:10.117 | מאביו | . |


## Milchamot ha-Shem 2:11

Hebrew (Wikisource): המורם מכל זה שאין לעבוד שום פרצוף וצורה מכל הנבראים הנמצאים בעולם העליון והשפל שהמציאם השי"ת בחסדו. כי לו לבדו ראוי לעבוד ולהתפלל, ועל זה באה המצוה והאזהרה בסיני: אנכי ה' אלהיך לא יהיה לך אלהים אחרים, כמו שכ' הרמב"ם ז"ל בפ' שני מהל' ע"ז. ובפ' ג' מהל' תשובה כתב וז"ל: ואלו שאין להם חלק לעולם הבא אלא נכרתין ואובדין על גודל רשעם וחטאתם לעולם ולעולמי עולמים: המינין והאפיקורסין וכו'. ובדין שאחריו כ': חמשה הן הנקראין מינים. האומר שאין שם אלוה ואין לעולם מנהיג, והאומר שיש שם מנהיג אבל הם שנים או יותר, והאומר שיש שם רבון אחד אלא שהוא גוף ובעל תמונה, והאומר שאינו לבדו הראשון וצור לכל. וכן העובד לזולתו כדי להיות מליץ בינו ובין רבון העולמים. כל אחד מחמשה אלו מין ע"כ. וכתב הרב לחם משנה על חלוקה ד' שכתב רבינו והאומר שאינו לבדו הראשון וכו', כ' הראב"ד, כאותו שאמר אלהיכם צייר גדול הוא אלא שמצא לו סמנים גדולים תהו ובהו וחשך ומים ורוח ובהם עשה מה שעשה, עכ"ל הראב"ד. ופי' הרב לח"מ שכוונתו להוסיף על דברי רבי' שאעפ"י שמודה שלא קדמה שום סבה לאלהינו שבשבילה נהיה, אלא שאומר שברא העולם יש מיש, הרי זה מין ע"ש. ובנדון דידן אמונת הקבלה החדשה כפי המבואר בספריהם לפי הבנתינו והבנת רבים וגדולים חכמי לב מרבותינו הפורשים ממנה ראשונים ואחרונים כהריב"ש וחו"י המובא בפתחי תשובה ושאר אחרונים נקבצו באו בה תלת וארבע מלי מהני דאמור רבנן, דהמאמין כן אין לו חלק לעה"ב. חדא רבוי האלוהות, א"ס ואדם קדמאה ואדם קדמון ועתיק ואריך אנפין ואבא ואימא וזעיר ונוקביה! שנית, שהם בעלי גוף דק, דהיינו אור, והא"ס נשמה לגופות הללו. שלישית, שהעבודה אינה לסבה הראשונה שהוא נקרא אצלם א"ס, רק לז"א שהוא סבה אחרונה שבאלו הסבות. רביעית, שהוא אמצעי הממשיך השפע מאלו המדות העליונות שהם עתיק ואריך אנפין ואבא ואימא, וקורין אותו בן לאו"א, והוא ג"כ אב לכתר דבריאה כפי השתלשלות העולמות לפי דעתם רחמנא ליצלן.

| id | word | punctuation after |
|---|---|---|
| 2:11.1 | המורם |  |
| 2:11.2 | מכל |  |
| 2:11.3 | זה |  |
| 2:11.4 | שאין |  |
| 2:11.5 | לעבוד |  |
| 2:11.6 | שום |  |
| 2:11.7 | פרצוף |  |
| 2:11.8 | וצורה |  |
| 2:11.9 | מכל |  |
| 2:11.10 | הנבראים |  |
| 2:11.11 | הנמצאים |  |
| 2:11.12 | בעולם |  |
| 2:11.13 | העליון |  |
| 2:11.14 | והשפל |  |
| 2:11.15 | שהמציאם |  |
| 2:11.16 | השי"ת |  |
| 2:11.17 | בחסדו | . |
| 2:11.18 | כי |  |
| 2:11.19 | לו |  |
| 2:11.20 | לבדו |  |
| 2:11.21 | ראוי |  |
| 2:11.22 | לעבוד |  |
| 2:11.23 | ולהתפלל | , |
| 2:11.24 | ועל |  |
| 2:11.25 | זה |  |
| 2:11.26 | באה |  |
| 2:11.27 | המצוה |  |
| 2:11.28 | והאזהרה |  |
| 2:11.29 | בסיני | : |
| 2:11.30 | אנכי |  |
| 2:11.31 | ה | ' |
| 2:11.32 | אלהיך |  |
| 2:11.33 | לא |  |
| 2:11.34 | יהיה |  |
| 2:11.35 | לך |  |
| 2:11.36 | אלהים |  |
| 2:11.37 | אחרים | , |
| 2:11.38 | כמו |  |
| 2:11.39 | שכ | ' |
| 2:11.40 | הרמב"ם |  |
| 2:11.41 | ז"ל |  |
| 2:11.42 | בפ | ' |
| 2:11.43 | שני |  |
| 2:11.44 | מהל | ' |
| 2:11.45 | ע"ז | . |
| 2:11.46 | ובפ | ' |
| 2:11.47 | ג | ' |
| 2:11.48 | מהל | ' |
| 2:11.49 | תשובה |  |
| 2:11.50 | כתב |  |
| 2:11.51 | וז"ל | : |
| 2:11.52 | ואלו |  |
| 2:11.53 | שאין |  |
| 2:11.54 | להם |  |
| 2:11.55 | חלק |  |
| 2:11.56 | לעולם |  |
| 2:11.57 | הבא |  |
| 2:11.58 | אלא |  |
| 2:11.59 | נכרתין |  |
| 2:11.60 | ואובדין |  |
| 2:11.61 | על |  |
| 2:11.62 | גודל |  |
| 2:11.63 | רשעם |  |
| 2:11.64 | וחטאתם |  |
| 2:11.65 | לעולם |  |
| 2:11.66 | ולעולמי |  |
| 2:11.67 | עולמים | : |
| 2:11.68 | המינין |  |
| 2:11.69 | והאפיקורסין |  |
| 2:11.70 | וכו | '. |
| 2:11.71 | ובדין |  |
| 2:11.72 | שאחריו |  |
| 2:11.73 | כ | ': |
| 2:11.74 | חמשה |  |
| 2:11.75 | הן |  |
| 2:11.76 | הנקראין |  |
| 2:11.77 | מינים | . |
| 2:11.78 | האומר |  |
| 2:11.79 | שאין |  |
| 2:11.80 | שם |  |
| 2:11.81 | אלוה |  |
| 2:11.82 | ואין |  |
| 2:11.83 | לעולם |  |
| 2:11.84 | מנהיג | , |
| 2:11.85 | והאומר |  |
| 2:11.86 | שיש |  |
| 2:11.87 | שם |  |
| 2:11.88 | מנהיג |  |
| 2:11.89 | אבל |  |
| 2:11.90 | הם |  |
| 2:11.91 | שנים |  |
| 2:11.92 | או |  |
| 2:11.93 | יותר | , |
| 2:11.94 | והאומר |  |
| 2:11.95 | שיש |  |
| 2:11.96 | שם |  |
| 2:11.97 | רבון |  |
| 2:11.98 | אחד |  |
| 2:11.99 | אלא |  |
| 2:11.100 | שהוא |  |
| 2:11.101 | גוף |  |
| 2:11.102 | ובעל |  |
| 2:11.103 | תמונה | , |
| 2:11.104 | והאומר |  |
| 2:11.105 | שאינו |  |
| 2:11.106 | לבדו |  |
| 2:11.107 | הראשון |  |
| 2:11.108 | וצור |  |
| 2:11.109 | לכל | . |
| 2:11.110 | וכן |  |
| 2:11.111 | העובד |  |
| 2:11.112 | לזולתו |  |
| 2:11.113 | כדי |  |
| 2:11.114 | להיות |  |
| 2:11.115 | מליץ |  |
| 2:11.116 | בינו |  |
| 2:11.117 | ובין |  |
| 2:11.118 | רבון |  |
| 2:11.119 | העולמים | . |
| 2:11.120 | כל |  |
| 2:11.121 | אחד |  |
| 2:11.122 | מחמשה |  |
| 2:11.123 | אלו |  |
| 2:11.124 | מין |  |
| 2:11.125 | ע"כ | . |
| 2:11.126 | וכתב |  |
| 2:11.127 | הרב |  |
| 2:11.128 | לחם |  |
| 2:11.129 | משנה |  |
| 2:11.130 | על |  |
| 2:11.131 | חלוקה |  |
| 2:11.132 | ד | ' |
| 2:11.133 | שכתב |  |
| 2:11.134 | רבינו |  |
| 2:11.135 | והאומר |  |
| 2:11.136 | שאינו |  |
| 2:11.137 | לבדו |  |
| 2:11.138 | הראשון |  |
| 2:11.139 | וכו | ', |
| 2:11.140 | כ | ' |
| 2:11.141 | הראב"ד | , |
| 2:11.142 | כאותו |  |
| 2:11.143 | שאמר |  |
| 2:11.144 | אלהיכם |  |
| 2:11.145 | צייר |  |
| 2:11.146 | גדול |  |
| 2:11.147 | הוא |  |
| 2:11.148 | אלא |  |
| 2:11.149 | שמצא |  |
| 2:11.150 | לו |  |
| 2:11.151 | סמנים |  |
| 2:11.152 | גדולים |  |
| 2:11.153 | תהו |  |
| 2:11.154 | ובהו |  |
| 2:11.155 | וחשך |  |
| 2:11.156 | ומים |  |
| 2:11.157 | ורוח |  |
| 2:11.158 | ובהם |  |
| 2:11.159 | עשה |  |
| 2:11.160 | מה |  |
| 2:11.161 | שעשה | , |
| 2:11.162 | עכ"ל |  |
| 2:11.163 | הראב"ד | . |
| 2:11.164 | ופי | ' |
| 2:11.165 | הרב |  |
| 2:11.166 | לח"מ |  |
| 2:11.167 | שכוונתו |  |
| 2:11.168 | להוסיף |  |
| 2:11.169 | על |  |
| 2:11.170 | דברי |  |
| 2:11.171 | רבי | ' |
| 2:11.172 | שאעפ"י |  |
| 2:11.173 | שמודה |  |
| 2:11.174 | שלא |  |
| 2:11.175 | קדמה |  |
| 2:11.176 | שום |  |
| 2:11.177 | סבה |  |
| 2:11.178 | לאלהינו |  |
| 2:11.179 | שבשבילה |  |
| 2:11.180 | נהיה | , |
| 2:11.181 | אלא |  |
| 2:11.182 | שאומר |  |
| 2:11.183 | שברא |  |
| 2:11.184 | העולם |  |
| 2:11.185 | יש |  |
| 2:11.186 | מיש | , |
| 2:11.187 | הרי |  |
| 2:11.188 | זה |  |
| 2:11.189 | מין |  |
| 2:11.190 | ע"ש | . |
| 2:11.191 | ובנדון |  |
| 2:11.192 | דידן |  |
| 2:11.193 | אמונת |  |
| 2:11.194 | הקבלה |  |
| 2:11.195 | החדשה |  |
| 2:11.196 | כפי |  |
| 2:11.197 | המבואר |  |
| 2:11.198 | בספריהם |  |
| 2:11.199 | לפי |  |
| 2:11.200 | הבנתינו |  |
| 2:11.201 | והבנת |  |
| 2:11.202 | רבים |  |
| 2:11.203 | וגדולים |  |
| 2:11.204 | חכמי |  |
| 2:11.205 | לב |  |
| 2:11.206 | מרבותינו |  |
| 2:11.207 | הפורשים |  |
| 2:11.208 | ממנה |  |
| 2:11.209 | ראשונים |  |
| 2:11.210 | ואחרונים |  |
| 2:11.211 | כהריב"ש |  |
| 2:11.212 | וחו"י |  |
| 2:11.213 | המובא |  |
| 2:11.214 | בפתחי |  |
| 2:11.215 | תשובה |  |
| 2:11.216 | ושאר |  |
| 2:11.217 | אחרונים |  |
| 2:11.218 | נקבצו |  |
| 2:11.219 | באו |  |
| 2:11.220 | בה |  |
| 2:11.221 | תלת |  |
| 2:11.222 | וארבע |  |
| 2:11.223 | מלי |  |
| 2:11.224 | מהני |  |
| 2:11.225 | דאמור |  |
| 2:11.226 | רבנן | , |
| 2:11.227 | דהמאמין |  |
| 2:11.228 | כן |  |
| 2:11.229 | אין |  |
| 2:11.230 | לו |  |
| 2:11.231 | חלק |  |
| 2:11.232 | לעה"ב | . |
| 2:11.233 | חדא |  |
| 2:11.234 | רבוי |  |
| 2:11.235 | האלוהות | , |
| 2:11.236 | א"ס |  |
| 2:11.237 | ואדם |  |
| 2:11.238 | קדמאה |  |
| 2:11.239 | ואדם |  |
| 2:11.240 | קדמון |  |
| 2:11.241 | ועתיק |  |
| 2:11.242 | ואריך |  |
| 2:11.243 | אנפין |  |
| 2:11.244 | ואבא |  |
| 2:11.245 | ואימא |  |
| 2:11.246 | וזעיר |  |
| 2:11.247 | ונוקביה | ! |
| 2:11.248 | שנית | , |
| 2:11.249 | שהם |  |
| 2:11.250 | בעלי |  |
| 2:11.251 | גוף |  |
| 2:11.252 | דק | , |
| 2:11.253 | דהיינו |  |
| 2:11.254 | אור | , |
| 2:11.255 | והא"ס |  |
| 2:11.256 | נשמה |  |
| 2:11.257 | לגופות |  |
| 2:11.258 | הללו | . |
| 2:11.259 | שלישית | , |
| 2:11.260 | שהעבודה |  |
| 2:11.261 | אינה |  |
| 2:11.262 | לסבה |  |
| 2:11.263 | הראשונה |  |
| 2:11.264 | שהוא |  |
| 2:11.265 | נקרא |  |
| 2:11.266 | אצלם |  |
| 2:11.267 | א"ס | , |
| 2:11.268 | רק |  |
| 2:11.269 | לז"א |  |
| 2:11.270 | שהוא |  |
| 2:11.271 | סבה |  |
| 2:11.272 | אחרונה |  |
| 2:11.273 | שבאלו |  |
| 2:11.274 | הסבות | . |
| 2:11.275 | רביעית | , |
| 2:11.276 | שהוא |  |
| 2:11.277 | אמצעי |  |
| 2:11.278 | הממשיך |  |
| 2:11.279 | השפע |  |
| 2:11.280 | מאלו |  |
| 2:11.281 | המדות |  |
| 2:11.282 | העליונות |  |
| 2:11.283 | שהם |  |
| 2:11.284 | עתיק |  |
| 2:11.285 | ואריך |  |
| 2:11.286 | אנפין |  |
| 2:11.287 | ואבא |  |
| 2:11.288 | ואימא | , |
| 2:11.289 | וקורין |  |
| 2:11.290 | אותו |  |
| 2:11.291 | בן |  |
| 2:11.292 | לאו"א | , |
| 2:11.293 | והוא |  |
| 2:11.294 | ג"כ |  |
| 2:11.295 | אב |  |
| 2:11.296 | לכתר |  |
| 2:11.297 | דבריאה |  |
| 2:11.298 | כפי |  |
| 2:11.299 | השתלשלות |  |
| 2:11.300 | העולמות |  |
| 2:11.301 | לפי |  |
| 2:11.302 | דעתם |  |
| 2:11.303 | רחמנא |  |
| 2:11.304 | ליצלן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

