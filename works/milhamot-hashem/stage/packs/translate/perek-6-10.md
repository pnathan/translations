# Translation pack: milhamot-hashem perek-6-10 (Milchamot ha-Shem 6:15–6:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-10`.

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
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
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

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
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
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- סבא **דסבין** → "the Elder of Elders".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
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
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **זווג** → "coupling".
- תורה **שבעל** פה → "the Oral Torah".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **גבורה** as a name of God → "the Power".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **ידידי** → "my friend" (second person kept).
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
- **תא חזי** → "come and see".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 6:15

Hebrew (Wikisource): סז) הן עתה אערוך לך ידידי הקורא דברי הפילסוף הנז' המבזה לימוד המשנה והתלמוד, ומזלזל בשקלא וטריא דרבנן בתלמוד שהם לפי דעתו מהסטרא אחרא הנקרא מצרים (פועל הרע) שמעבידם ומטריחם בכמה קושיות ופירוקם כאשר יבין המשכיל בדבריו איך הוא משביח את הרמזים שלו בפסוק ונהר יוצא מעדן (זוהר בראשית דף כ"ו וכ"ז) בשם סבא דסבין (עתיק) שנגלה לר' אלעזר וחביריו. ור' אלעזר סיפר זה לאביו ר' שמעון שאמר להם סבא דסבין שאותן ארבעה שנכנסו לפרדס. השלשה מהם שלא למדו הסודות שלו שהם מאבני שיש טהור שהם מאבא וברתא והם מלוחות ראשונות שנשתברו, רק למדו משנה ותלמוד שהם מלוחות שניות שיש בהם טוב ורע והוא התר ואסור, לפיכך נסתכנו, אבל ר' עקיבא למד הסודות שהם ענייני רבוי האלוהות שהוא מרמזם באותיות שם הוי'ה ב"ה ושהסודות המבוארים בדבריו הם מאבני שיש טהור שכולו טוב והם ממדרגת לוחות ראשונות שנשתברו ואין בהם תערובת רע, ולפיכך נכנס בשלום ויצא בשלום, אבל המשנה והתלמוד ושאר מדרשי רז"ל הם מעורבים מטוב ורע. ראה ידידי ותן דעתך בדבריו איך הוא מבזה לימוד המשנה והתלמוד ומדרשי רז"ל ומעיד עליהם שהם עץ הדעת טוב ורע היפך דברי שלמה המלך ע"ה שאמר ברוח הקדש כי לקח טוב נתתי לכם תורתי אל תעזובו, ורז"ל במדרשיהם הרבו בשבח תורה שבכתב, ותורה שבעל פה, יותר כמו שכתבתי לעיל בש' הירושלמי בשם רשב"י.

| id | word | punctuation after |
|---|---|---|
| 6:15.1 | סז | ) |
| 6:15.2 | הן |  |
| 6:15.3 | עתה |  |
| 6:15.4 | אערוך |  |
| 6:15.5 | לך |  |
| 6:15.6 | ידידי |  |
| 6:15.7 | הקורא |  |
| 6:15.8 | דברי |  |
| 6:15.9 | הפילסוף |  |
| 6:15.10 | הנז | ' |
| 6:15.11 | המבזה |  |
| 6:15.12 | לימוד |  |
| 6:15.13 | המשנה |  |
| 6:15.14 | והתלמוד | , |
| 6:15.15 | ומזלזל |  |
| 6:15.16 | בשקלא |  |
| 6:15.17 | וטריא |  |
| 6:15.18 | דרבנן |  |
| 6:15.19 | בתלמוד |  |
| 6:15.20 | שהם |  |
| 6:15.21 | לפי |  |
| 6:15.22 | דעתו |  |
| 6:15.23 | מהסטרא |  |
| 6:15.24 | אחרא |  |
| 6:15.25 | הנקרא |  |
| 6:15.26 | מצרים |  |
| 6:15.27 | פועל |  |
| 6:15.28 | הרע | ) |
| 6:15.29 | שמעבידם |  |
| 6:15.30 | ומטריחם |  |
| 6:15.31 | בכמה |  |
| 6:15.32 | קושיות |  |
| 6:15.33 | ופירוקם |  |
| 6:15.34 | כאשר |  |
| 6:15.35 | יבין |  |
| 6:15.36 | המשכיל |  |
| 6:15.37 | בדבריו |  |
| 6:15.38 | איך |  |
| 6:15.39 | הוא |  |
| 6:15.40 | משביח |  |
| 6:15.41 | את |  |
| 6:15.42 | הרמזים |  |
| 6:15.43 | שלו |  |
| 6:15.44 | בפסוק |  |
| 6:15.45 | ונהר |  |
| 6:15.46 | יוצא |  |
| 6:15.47 | מעדן |  |
| 6:15.48 | זוהר |  |
| 6:15.49 | בראשית |  |
| 6:15.50 | דף |  |
| 6:15.51 | כ"ו |  |
| 6:15.52 | וכ"ז | ) |
| 6:15.53 | בשם |  |
| 6:15.54 | סבא |  |
| 6:15.55 | דסבין |  |
| 6:15.56 | עתיק | ) |
| 6:15.57 | שנגלה |  |
| 6:15.58 | לר | ' |
| 6:15.59 | אלעזר |  |
| 6:15.60 | וחביריו | . |
| 6:15.61 | ור | ' |
| 6:15.62 | אלעזר |  |
| 6:15.63 | סיפר |  |
| 6:15.64 | זה |  |
| 6:15.65 | לאביו |  |
| 6:15.66 | ר | ' |
| 6:15.67 | שמעון |  |
| 6:15.68 | שאמר |  |
| 6:15.69 | להם |  |
| 6:15.70 | סבא |  |
| 6:15.71 | דסבין |  |
| 6:15.72 | שאותן |  |
| 6:15.73 | ארבעה |  |
| 6:15.74 | שנכנסו |  |
| 6:15.75 | לפרדס | . |
| 6:15.76 | השלשה |  |
| 6:15.77 | מהם |  |
| 6:15.78 | שלא |  |
| 6:15.79 | למדו |  |
| 6:15.80 | הסודות |  |
| 6:15.81 | שלו |  |
| 6:15.82 | שהם |  |
| 6:15.83 | מאבני |  |
| 6:15.84 | שיש |  |
| 6:15.85 | טהור |  |
| 6:15.86 | שהם |  |
| 6:15.87 | מאבא |  |
| 6:15.88 | וברתא |  |
| 6:15.89 | והם |  |
| 6:15.90 | מלוחות |  |
| 6:15.91 | ראשונות |  |
| 6:15.92 | שנשתברו | , |
| 6:15.93 | רק |  |
| 6:15.94 | למדו |  |
| 6:15.95 | משנה |  |
| 6:15.96 | ותלמוד |  |
| 6:15.97 | שהם |  |
| 6:15.98 | מלוחות |  |
| 6:15.99 | שניות |  |
| 6:15.100 | שיש |  |
| 6:15.101 | בהם |  |
| 6:15.102 | טוב |  |
| 6:15.103 | ורע |  |
| 6:15.104 | והוא |  |
| 6:15.105 | התר |  |
| 6:15.106 | ואסור | , |
| 6:15.107 | לפיכך |  |
| 6:15.108 | נסתכנו | , |
| 6:15.109 | אבל |  |
| 6:15.110 | ר | ' |
| 6:15.111 | עקיבא |  |
| 6:15.112 | למד |  |
| 6:15.113 | הסודות |  |
| 6:15.114 | שהם |  |
| 6:15.115 | ענייני |  |
| 6:15.116 | רבוי |  |
| 6:15.117 | האלוהות |  |
| 6:15.118 | שהוא |  |
| 6:15.119 | מרמזם |  |
| 6:15.120 | באותיות |  |
| 6:15.121 | שם |  |
| 6:15.122 | הוי'ה |  |
| 6:15.123 | ב"ה |  |
| 6:15.124 | ושהסודות |  |
| 6:15.125 | המבוארים |  |
| 6:15.126 | בדבריו |  |
| 6:15.127 | הם |  |
| 6:15.128 | מאבני |  |
| 6:15.129 | שיש |  |
| 6:15.130 | טהור |  |
| 6:15.131 | שכולו |  |
| 6:15.132 | טוב |  |
| 6:15.133 | והם |  |
| 6:15.134 | ממדרגת |  |
| 6:15.135 | לוחות |  |
| 6:15.136 | ראשונות |  |
| 6:15.137 | שנשתברו |  |
| 6:15.138 | ואין |  |
| 6:15.139 | בהם |  |
| 6:15.140 | תערובת |  |
| 6:15.141 | רע | , |
| 6:15.142 | ולפיכך |  |
| 6:15.143 | נכנס |  |
| 6:15.144 | בשלום |  |
| 6:15.145 | ויצא |  |
| 6:15.146 | בשלום | , |
| 6:15.147 | אבל |  |
| 6:15.148 | המשנה |  |
| 6:15.149 | והתלמוד |  |
| 6:15.150 | ושאר |  |
| 6:15.151 | מדרשי |  |
| 6:15.152 | רז"ל |  |
| 6:15.153 | הם |  |
| 6:15.154 | מעורבים |  |
| 6:15.155 | מטוב |  |
| 6:15.156 | ורע | . |
| 6:15.157 | ראה |  |
| 6:15.158 | ידידי |  |
| 6:15.159 | ותן |  |
| 6:15.160 | דעתך |  |
| 6:15.161 | בדבריו |  |
| 6:15.162 | איך |  |
| 6:15.163 | הוא |  |
| 6:15.164 | מבזה |  |
| 6:15.165 | לימוד |  |
| 6:15.166 | המשנה |  |
| 6:15.167 | והתלמוד |  |
| 6:15.168 | ומדרשי |  |
| 6:15.169 | רז"ל |  |
| 6:15.170 | ומעיד |  |
| 6:15.171 | עליהם |  |
| 6:15.172 | שהם |  |
| 6:15.173 | עץ |  |
| 6:15.174 | הדעת |  |
| 6:15.175 | טוב |  |
| 6:15.176 | ורע |  |
| 6:15.177 | היפך |  |
| 6:15.178 | דברי |  |
| 6:15.179 | שלמה |  |
| 6:15.180 | המלך |  |
| 6:15.181 | ע"ה |  |
| 6:15.182 | שאמר |  |
| 6:15.183 | ברוח |  |
| 6:15.184 | הקדש |  |
| 6:15.185 | כי |  |
| 6:15.186 | לקח |  |
| 6:15.187 | טוב |  |
| 6:15.188 | נתתי |  |
| 6:15.189 | לכם |  |
| 6:15.190 | תורתי |  |
| 6:15.191 | אל |  |
| 6:15.192 | תעזובו | , |
| 6:15.193 | ורז"ל |  |
| 6:15.194 | במדרשיהם |  |
| 6:15.195 | הרבו |  |
| 6:15.196 | בשבח |  |
| 6:15.197 | תורה |  |
| 6:15.198 | שבכתב | , |
| 6:15.199 | ותורה |  |
| 6:15.200 | שבעל |  |
| 6:15.201 | פה | , |
| 6:15.202 | יותר |  |
| 6:15.203 | כמו |  |
| 6:15.204 | שכתבתי |  |
| 6:15.205 | לעיל |  |
| 6:15.206 | בש | ' |
| 6:15.207 | הירושלמי |  |
| 6:15.208 | בשם |  |
| 6:15.209 | רשב"י | . |


## Milchamot ha-Shem 6:16

Hebrew (Wikisource): ובכמה מקומות קורים רז"ל לתורה טוב כמו שאמרו על זנח ישראל טוב אויב ירדפו, ואין טוב אלא תורה שנ' כי לקח טוב נתתי לכם תורתי אל תעזובו. וכתיב עץ החיים היא למחזיקים בה, ולא עץ הדעת טוב ורע, ושוב חזר הפילסוף הנזכר לגנות את המשנה והתלמוד באמרו דבר אחר (שם דף כ"ז ע"א) ומשבח את סחורתו הסודות שהוא בודאם מלבו שהם מעץ החיים מזעיר אנפין דאצילות ששם אין שליטא לקליפות, ולא ימצא בהם קושיות ומשא ומתן של הלכה (אבל הם המצאת אלהים חדשים מקרוב באו לא שערום אבותינו והם עוקרים יסוד יחוד השם שהוא העיקר הגדול). אבל המשנה והתלמוד הם יוצאים ממטטרון דיצירה והוא זעיר אנפין דיצירה ששם נכנסו בן עזאי ובן זומא ואלישע, אחר ששם יש אחיזה לקליפות והם מעץ הדעת טוב ורע, מצד אחד טוב, והוא היתר כשר טהור, ומצד אחד רע והוא אסור ופסול וטמא (כאן גילה מצפוני לבו דניחא ליה בהפקירא ולא יזכר שם טמא ופסול ואסור החכם יבין). ועוד חזר שם לדרוש בגנות המשנה והתלמוד בשם מארי מתניתין וז"ל וימררו את חייהם (המצרים שהם הקליפות) בעבודה קשה, דא קושיא, בחומר, בקל וחומר, ובלבנים, בליבון הלכתא, ובכל עבודה בשדה, דא ברייתא, את כל עבודתם אשר עבדו בהם בפרך, דא משנה, אם תיבין בתיובתא. אתאמר בהון ויורהו ה' עץ, ודא עץ חיים, וביה ויתמתקו המים (פי' המקדש מלך בתחלה היו עומקים בפשטי התורה (דהיינו משנה וגמרא וברייתא שהזכיר לעיל) אזי וימררו את חייהם שאינם מקבלים שפע מזעיר דאצילות. ואחר ששבו בתשוּבה וחזרו ללמוד סודות התורה אז מושכים שפע משם מ"ה שהוא שקיו דאילנא שהוא פנימיות זעיר וכו') ודא משה משיח דאתמר ביה ומטה אלהים בידו, מטה דא מטטרון, מסטריה חיים, ומסטריה מיתה. כד אתהפיך למטה איהו עזר מסטרה דטוב. כך אתהפיך לחויא איהו כנגדו. מיד וינס משה מפניו. וקב"ה מסר ליה בידיה דמשה ואיהו אוריתא דבעל פה דביה איסור והתר (פי' המק"מ כי תורה שבע"פ היא במטטרון ששם טוב ורע שהם איסור והתר) הנך רואה בעיניך ידידי את דברי הפילסוף החובר מחבר הזהר איך הוא מכוין ליתן דופי במשנת חכמים ותלמודם המקובל בידם איש מפי איש ממשה רבינו ע"ה. וליחס את המשנה למטטרון שיש בו טוב ורע. ולא מפי הגבורה ניתנה אלא ממטטרון. והטוב שבה הוא התר טהור כשר, ורע שבה אסור וטמא ופסול. וחשב את העסק בה לעון אשר חטא. אשר אם ישובו ישראל בתשובה ויעזבו המשנה, ויתעסקו בסודות שערך לפניהם ברבוי האלוהות הקדושים לו פועלי הטוב, וגם בידיעת הקליפות אלהים אחרים טמאים פועלי הרע ימתקו להם המים המרים המאררים. כי אז לא ידעו עוד להבדיל בין טמא לטהור ובין קדש לחול, כאשר מצאנו ראינו האנשים העוסקים כל ימיהם בסודות הנזכרים ואינם משגיחים במשנה ותלמוד ופוסקים, שאינם יודעים כלל להבדיל בין הטמא ובין הטהור, כשר ופסול, אסור ומותר, ולא ידעו במה יכשלו, כי מוחם מלא מלי דחסידותא דאית בהו אפיקרותא, עבודת הפרצופים וכפירות שונים אשר לא צוה ה', ומוחין דקטנות ומוחין דגדלות, וזווגים וטפין הנוזלים מן הפרצופים הזכרים לרחם נשיהם ופילגשיהם, עד אפס מקום במוחם לנוח שם אסור ומותר. להבדיל בין הטמא ובין הטהור, ובין דבר הנאכל לדבר שאסור לאכלו אשר על זה נצטוינו בתורה, וקיימו בעצמם מאמר ר' נתן שאמר כי דבר ה' בזה זה שלא השגיח על המשנה (פי' רש"י שעשאה כפי שאינה עיקר. או שאינו מחזר לשנותה, סנהדרין דף צ"ט ע"ש ברש"י ומהרש"א).

| id | word | punctuation after |
|---|---|---|
| 6:16.1 | ובכמה |  |
| 6:16.2 | מקומות |  |
| 6:16.3 | קורים |  |
| 6:16.4 | רז"ל |  |
| 6:16.5 | לתורה |  |
| 6:16.6 | טוב |  |
| 6:16.7 | כמו |  |
| 6:16.8 | שאמרו |  |
| 6:16.9 | על |  |
| 6:16.10 | זנח |  |
| 6:16.11 | ישראל |  |
| 6:16.12 | טוב |  |
| 6:16.13 | אויב |  |
| 6:16.14 | ירדפו | , |
| 6:16.15 | ואין |  |
| 6:16.16 | טוב |  |
| 6:16.17 | אלא |  |
| 6:16.18 | תורה |  |
| 6:16.19 | שנ | ' |
| 6:16.20 | כי |  |
| 6:16.21 | לקח |  |
| 6:16.22 | טוב |  |
| 6:16.23 | נתתי |  |
| 6:16.24 | לכם |  |
| 6:16.25 | תורתי |  |
| 6:16.26 | אל |  |
| 6:16.27 | תעזובו | . |
| 6:16.28 | וכתיב |  |
| 6:16.29 | עץ |  |
| 6:16.30 | החיים |  |
| 6:16.31 | היא |  |
| 6:16.32 | למחזיקים |  |
| 6:16.33 | בה | , |
| 6:16.34 | ולא |  |
| 6:16.35 | עץ |  |
| 6:16.36 | הדעת |  |
| 6:16.37 | טוב |  |
| 6:16.38 | ורע | , |
| 6:16.39 | ושוב |  |
| 6:16.40 | חזר |  |
| 6:16.41 | הפילסוף |  |
| 6:16.42 | הנזכר |  |
| 6:16.43 | לגנות |  |
| 6:16.44 | את |  |
| 6:16.45 | המשנה |  |
| 6:16.46 | והתלמוד |  |
| 6:16.47 | באמרו |  |
| 6:16.48 | דבר |  |
| 6:16.49 | אחר |  |
| 6:16.50 | שם |  |
| 6:16.51 | דף |  |
| 6:16.52 | כ"ז |  |
| 6:16.53 | ע"א | ) |
| 6:16.54 | ומשבח |  |
| 6:16.55 | את |  |
| 6:16.56 | סחורתו |  |
| 6:16.57 | הסודות |  |
| 6:16.58 | שהוא |  |
| 6:16.59 | בודאם |  |
| 6:16.60 | מלבו |  |
| 6:16.61 | שהם |  |
| 6:16.62 | מעץ |  |
| 6:16.63 | החיים |  |
| 6:16.64 | מזעיר |  |
| 6:16.65 | אנפין |  |
| 6:16.66 | דאצילות |  |
| 6:16.67 | ששם |  |
| 6:16.68 | אין |  |
| 6:16.69 | שליטא |  |
| 6:16.70 | לקליפות | , |
| 6:16.71 | ולא |  |
| 6:16.72 | ימצא |  |
| 6:16.73 | בהם |  |
| 6:16.74 | קושיות |  |
| 6:16.75 | ומשא |  |
| 6:16.76 | ומתן |  |
| 6:16.77 | של |  |
| 6:16.78 | הלכה |  |
| 6:16.79 | אבל |  |
| 6:16.80 | הם |  |
| 6:16.81 | המצאת |  |
| 6:16.82 | אלהים |  |
| 6:16.83 | חדשים |  |
| 6:16.84 | מקרוב |  |
| 6:16.85 | באו |  |
| 6:16.86 | לא |  |
| 6:16.87 | שערום |  |
| 6:16.88 | אבותינו |  |
| 6:16.89 | והם |  |
| 6:16.90 | עוקרים |  |
| 6:16.91 | יסוד |  |
| 6:16.92 | יחוד |  |
| 6:16.93 | השם |  |
| 6:16.94 | שהוא |  |
| 6:16.95 | העיקר |  |
| 6:16.96 | הגדול | ). |
| 6:16.97 | אבל |  |
| 6:16.98 | המשנה |  |
| 6:16.99 | והתלמוד |  |
| 6:16.100 | הם |  |
| 6:16.101 | יוצאים |  |
| 6:16.102 | ממטטרון |  |
| 6:16.103 | דיצירה |  |
| 6:16.104 | והוא |  |
| 6:16.105 | זעיר |  |
| 6:16.106 | אנפין |  |
| 6:16.107 | דיצירה |  |
| 6:16.108 | ששם |  |
| 6:16.109 | נכנסו |  |
| 6:16.110 | בן |  |
| 6:16.111 | עזאי |  |
| 6:16.112 | ובן |  |
| 6:16.113 | זומא |  |
| 6:16.114 | ואלישע | , |
| 6:16.115 | אחר |  |
| 6:16.116 | ששם |  |
| 6:16.117 | יש |  |
| 6:16.118 | אחיזה |  |
| 6:16.119 | לקליפות |  |
| 6:16.120 | והם |  |
| 6:16.121 | מעץ |  |
| 6:16.122 | הדעת |  |
| 6:16.123 | טוב |  |
| 6:16.124 | ורע | , |
| 6:16.125 | מצד |  |
| 6:16.126 | אחד |  |
| 6:16.127 | טוב | , |
| 6:16.128 | והוא |  |
| 6:16.129 | היתר |  |
| 6:16.130 | כשר |  |
| 6:16.131 | טהור | , |
| 6:16.132 | ומצד |  |
| 6:16.133 | אחד |  |
| 6:16.134 | רע |  |
| 6:16.135 | והוא |  |
| 6:16.136 | אסור |  |
| 6:16.137 | ופסול |  |
| 6:16.138 | וטמא |  |
| 6:16.139 | כאן |  |
| 6:16.140 | גילה |  |
| 6:16.141 | מצפוני |  |
| 6:16.142 | לבו |  |
| 6:16.143 | דניחא |  |
| 6:16.144 | ליה |  |
| 6:16.145 | בהפקירא |  |
| 6:16.146 | ולא |  |
| 6:16.147 | יזכר |  |
| 6:16.148 | שם |  |
| 6:16.149 | טמא |  |
| 6:16.150 | ופסול |  |
| 6:16.151 | ואסור |  |
| 6:16.152 | החכם |  |
| 6:16.153 | יבין | ). |
| 6:16.154 | ועוד |  |
| 6:16.155 | חזר |  |
| 6:16.156 | שם |  |
| 6:16.157 | לדרוש |  |
| 6:16.158 | בגנות |  |
| 6:16.159 | המשנה |  |
| 6:16.160 | והתלמוד |  |
| 6:16.161 | בשם |  |
| 6:16.162 | מארי |  |
| 6:16.163 | מתניתין |  |
| 6:16.164 | וז"ל |  |
| 6:16.165 | וימררו |  |
| 6:16.166 | את |  |
| 6:16.167 | חייהם |  |
| 6:16.168 | המצרים |  |
| 6:16.169 | שהם |  |
| 6:16.170 | הקליפות | ) |
| 6:16.171 | בעבודה |  |
| 6:16.172 | קשה | , |
| 6:16.173 | דא |  |
| 6:16.174 | קושיא | , |
| 6:16.175 | בחומר | , |
| 6:16.176 | בקל |  |
| 6:16.177 | וחומר | , |
| 6:16.178 | ובלבנים | , |
| 6:16.179 | בליבון |  |
| 6:16.180 | הלכתא | , |
| 6:16.181 | ובכל |  |
| 6:16.182 | עבודה |  |
| 6:16.183 | בשדה | , |
| 6:16.184 | דא |  |
| 6:16.185 | ברייתא | , |
| 6:16.186 | את |  |
| 6:16.187 | כל |  |
| 6:16.188 | עבודתם |  |
| 6:16.189 | אשר |  |
| 6:16.190 | עבדו |  |
| 6:16.191 | בהם |  |
| 6:16.192 | בפרך | , |
| 6:16.193 | דא |  |
| 6:16.194 | משנה | , |
| 6:16.195 | אם |  |
| 6:16.196 | תיבין |  |
| 6:16.197 | בתיובתא | . |
| 6:16.198 | אתאמר |  |
| 6:16.199 | בהון |  |
| 6:16.200 | ויורהו |  |
| 6:16.201 | ה | ' |
| 6:16.202 | עץ | , |
| 6:16.203 | ודא |  |
| 6:16.204 | עץ |  |
| 6:16.205 | חיים | , |
| 6:16.206 | וביה |  |
| 6:16.207 | ויתמתקו |  |
| 6:16.208 | המים |  |
| 6:16.209 | פי | ' |
| 6:16.210 | המקדש |  |
| 6:16.211 | מלך |  |
| 6:16.212 | בתחלה |  |
| 6:16.213 | היו |  |
| 6:16.214 | עומקים |  |
| 6:16.215 | בפשטי |  |
| 6:16.216 | התורה |  |
| 6:16.217 | דהיינו |  |
| 6:16.218 | משנה |  |
| 6:16.219 | וגמרא |  |
| 6:16.220 | וברייתא |  |
| 6:16.221 | שהזכיר |  |
| 6:16.222 | לעיל | ) |
| 6:16.223 | אזי |  |
| 6:16.224 | וימררו |  |
| 6:16.225 | את |  |
| 6:16.226 | חייהם |  |
| 6:16.227 | שאינם |  |
| 6:16.228 | מקבלים |  |
| 6:16.229 | שפע |  |
| 6:16.230 | מזעיר |  |
| 6:16.231 | דאצילות | . |
| 6:16.232 | ואחר |  |
| 6:16.233 | ששבו |  |
| 6:16.234 | בתשוּבה |  |
| 6:16.235 | וחזרו |  |
| 6:16.236 | ללמוד |  |
| 6:16.237 | סודות |  |
| 6:16.238 | התורה |  |
| 6:16.239 | אז |  |
| 6:16.240 | מושכים |  |
| 6:16.241 | שפע |  |
| 6:16.242 | משם |  |
| 6:16.243 | מ"ה |  |
| 6:16.244 | שהוא |  |
| 6:16.245 | שקיו |  |
| 6:16.246 | דאילנא |  |
| 6:16.247 | שהוא |  |
| 6:16.248 | פנימיות |  |
| 6:16.249 | זעיר |  |
| 6:16.250 | וכו | ') |
| 6:16.251 | ודא |  |
| 6:16.252 | משה |  |
| 6:16.253 | משיח |  |
| 6:16.254 | דאתמר |  |
| 6:16.255 | ביה |  |
| 6:16.256 | ומטה |  |
| 6:16.257 | אלהים |  |
| 6:16.258 | בידו | , |
| 6:16.259 | מטה |  |
| 6:16.260 | דא |  |
| 6:16.261 | מטטרון | , |
| 6:16.262 | מסטריה |  |
| 6:16.263 | חיים | , |
| 6:16.264 | ומסטריה |  |
| 6:16.265 | מיתה | . |
| 6:16.266 | כד |  |
| 6:16.267 | אתהפיך |  |
| 6:16.268 | למטה |  |
| 6:16.269 | איהו |  |
| 6:16.270 | עזר |  |
| 6:16.271 | מסטרה |  |
| 6:16.272 | דטוב | . |
| 6:16.273 | כך |  |
| 6:16.274 | אתהפיך |  |
| 6:16.275 | לחויא |  |
| 6:16.276 | איהו |  |
| 6:16.277 | כנגדו | . |
| 6:16.278 | מיד |  |
| 6:16.279 | וינס |  |
| 6:16.280 | משה |  |
| 6:16.281 | מפניו | . |
| 6:16.282 | וקב"ה |  |
| 6:16.283 | מסר |  |
| 6:16.284 | ליה |  |
| 6:16.285 | בידיה |  |
| 6:16.286 | דמשה |  |
| 6:16.287 | ואיהו |  |
| 6:16.288 | אוריתא |  |
| 6:16.289 | דבעל |  |
| 6:16.290 | פה |  |
| 6:16.291 | דביה |  |
| 6:16.292 | איסור |  |
| 6:16.293 | והתר |  |
| 6:16.294 | פי | ' |
| 6:16.295 | המק"מ |  |
| 6:16.296 | כי |  |
| 6:16.297 | תורה |  |
| 6:16.298 | שבע"פ |  |
| 6:16.299 | היא |  |
| 6:16.300 | במטטרון |  |
| 6:16.301 | ששם |  |
| 6:16.302 | טוב |  |
| 6:16.303 | ורע |  |
| 6:16.304 | שהם |  |
| 6:16.305 | איסור |  |
| 6:16.306 | והתר | ) |
| 6:16.307 | הנך |  |
| 6:16.308 | רואה |  |
| 6:16.309 | בעיניך |  |
| 6:16.310 | ידידי |  |
| 6:16.311 | את |  |
| 6:16.312 | דברי |  |
| 6:16.313 | הפילסוף |  |
| 6:16.314 | החובר |  |
| 6:16.315 | מחבר |  |
| 6:16.316 | הזהר |  |
| 6:16.317 | איך |  |
| 6:16.318 | הוא |  |
| 6:16.319 | מכוין |  |
| 6:16.320 | ליתן |  |
| 6:16.321 | דופי |  |
| 6:16.322 | במשנת |  |
| 6:16.323 | חכמים |  |
| 6:16.324 | ותלמודם |  |
| 6:16.325 | המקובל |  |
| 6:16.326 | בידם |  |
| 6:16.327 | איש |  |
| 6:16.328 | מפי |  |
| 6:16.329 | איש |  |
| 6:16.330 | ממשה |  |
| 6:16.331 | רבינו |  |
| 6:16.332 | ע"ה | . |
| 6:16.333 | וליחס |  |
| 6:16.334 | את |  |
| 6:16.335 | המשנה |  |
| 6:16.336 | למטטרון |  |
| 6:16.337 | שיש |  |
| 6:16.338 | בו |  |
| 6:16.339 | טוב |  |
| 6:16.340 | ורע | . |
| 6:16.341 | ולא |  |
| 6:16.342 | מפי |  |
| 6:16.343 | הגבורה |  |
| 6:16.344 | ניתנה |  |
| 6:16.345 | אלא |  |
| 6:16.346 | ממטטרון | . |
| 6:16.347 | והטוב |  |
| 6:16.348 | שבה |  |
| 6:16.349 | הוא |  |
| 6:16.350 | התר |  |
| 6:16.351 | טהור |  |
| 6:16.352 | כשר | , |
| 6:16.353 | ורע |  |
| 6:16.354 | שבה |  |
| 6:16.355 | אסור |  |
| 6:16.356 | וטמא |  |
| 6:16.357 | ופסול | . |
| 6:16.358 | וחשב |  |
| 6:16.359 | את |  |
| 6:16.360 | העסק |  |
| 6:16.361 | בה |  |
| 6:16.362 | לעון |  |
| 6:16.363 | אשר |  |
| 6:16.364 | חטא | . |
| 6:16.365 | אשר |  |
| 6:16.366 | אם |  |
| 6:16.367 | ישובו |  |
| 6:16.368 | ישראל |  |
| 6:16.369 | בתשובה |  |
| 6:16.370 | ויעזבו |  |
| 6:16.371 | המשנה | , |
| 6:16.372 | ויתעסקו |  |
| 6:16.373 | בסודות |  |
| 6:16.374 | שערך |  |
| 6:16.375 | לפניהם |  |
| 6:16.376 | ברבוי |  |
| 6:16.377 | האלוהות |  |
| 6:16.378 | הקדושים |  |
| 6:16.379 | לו |  |
| 6:16.380 | פועלי |  |
| 6:16.381 | הטוב | , |
| 6:16.382 | וגם |  |
| 6:16.383 | בידיעת |  |
| 6:16.384 | הקליפות |  |
| 6:16.385 | אלהים |  |
| 6:16.386 | אחרים |  |
| 6:16.387 | טמאים |  |
| 6:16.388 | פועלי |  |
| 6:16.389 | הרע |  |
| 6:16.390 | ימתקו |  |
| 6:16.391 | להם |  |
| 6:16.392 | המים |  |
| 6:16.393 | המרים |  |
| 6:16.394 | המאררים | . |
| 6:16.395 | כי |  |
| 6:16.396 | אז |  |
| 6:16.397 | לא |  |
| 6:16.398 | ידעו |  |
| 6:16.399 | עוד |  |
| 6:16.400 | להבדיל |  |
| 6:16.401 | בין |  |
| 6:16.402 | טמא |  |
| 6:16.403 | לטהור |  |
| 6:16.404 | ובין |  |
| 6:16.405 | קדש |  |
| 6:16.406 | לחול | , |
| 6:16.407 | כאשר |  |
| 6:16.408 | מצאנו |  |
| 6:16.409 | ראינו |  |
| 6:16.410 | האנשים |  |
| 6:16.411 | העוסקים |  |
| 6:16.412 | כל |  |
| 6:16.413 | ימיהם |  |
| 6:16.414 | בסודות |  |
| 6:16.415 | הנזכרים |  |
| 6:16.416 | ואינם |  |
| 6:16.417 | משגיחים |  |
| 6:16.418 | במשנה |  |
| 6:16.419 | ותלמוד |  |
| 6:16.420 | ופוסקים | , |
| 6:16.421 | שאינם |  |
| 6:16.422 | יודעים |  |
| 6:16.423 | כלל |  |
| 6:16.424 | להבדיל |  |
| 6:16.425 | בין |  |
| 6:16.426 | הטמא |  |
| 6:16.427 | ובין |  |
| 6:16.428 | הטהור | , |
| 6:16.429 | כשר |  |
| 6:16.430 | ופסול | , |
| 6:16.431 | אסור |  |
| 6:16.432 | ומותר | , |
| 6:16.433 | ולא |  |
| 6:16.434 | ידעו |  |
| 6:16.435 | במה |  |
| 6:16.436 | יכשלו | , |
| 6:16.437 | כי |  |
| 6:16.438 | מוחם |  |
| 6:16.439 | מלא |  |
| 6:16.440 | מלי |  |
| 6:16.441 | דחסידותא |  |
| 6:16.442 | דאית |  |
| 6:16.443 | בהו |  |
| 6:16.444 | אפיקרותא | , |
| 6:16.445 | עבודת |  |
| 6:16.446 | הפרצופים |  |
| 6:16.447 | וכפירות |  |
| 6:16.448 | שונים |  |
| 6:16.449 | אשר |  |
| 6:16.450 | לא |  |
| 6:16.451 | צוה |  |
| 6:16.452 | ה | ', |
| 6:16.453 | ומוחין |  |
| 6:16.454 | דקטנות |  |
| 6:16.455 | ומוחין |  |
| 6:16.456 | דגדלות | , |
| 6:16.457 | וזווגים |  |
| 6:16.458 | וטפין |  |
| 6:16.459 | הנוזלים |  |
| 6:16.460 | מן |  |
| 6:16.461 | הפרצופים |  |
| 6:16.462 | הזכרים |  |
| 6:16.463 | לרחם |  |
| 6:16.464 | נשיהם |  |
| 6:16.465 | ופילגשיהם | , |
| 6:16.466 | עד |  |
| 6:16.467 | אפס |  |
| 6:16.468 | מקום |  |
| 6:16.469 | במוחם |  |
| 6:16.470 | לנוח |  |
| 6:16.471 | שם |  |
| 6:16.472 | אסור |  |
| 6:16.473 | ומותר | . |
| 6:16.474 | להבדיל |  |
| 6:16.475 | בין |  |
| 6:16.476 | הטמא |  |
| 6:16.477 | ובין |  |
| 6:16.478 | הטהור | , |
| 6:16.479 | ובין |  |
| 6:16.480 | דבר |  |
| 6:16.481 | הנאכל |  |
| 6:16.482 | לדבר |  |
| 6:16.483 | שאסור |  |
| 6:16.484 | לאכלו |  |
| 6:16.485 | אשר |  |
| 6:16.486 | על |  |
| 6:16.487 | זה |  |
| 6:16.488 | נצטוינו |  |
| 6:16.489 | בתורה | , |
| 6:16.490 | וקיימו |  |
| 6:16.491 | בעצמם |  |
| 6:16.492 | מאמר |  |
| 6:16.493 | ר | ' |
| 6:16.494 | נתן |  |
| 6:16.495 | שאמר |  |
| 6:16.496 | כי |  |
| 6:16.497 | דבר |  |
| 6:16.498 | ה | ' |
| 6:16.499 | בזה |  |
| 6:16.500 | זה |  |
| 6:16.501 | שלא |  |
| 6:16.502 | השגיח |  |
| 6:16.503 | על |  |
| 6:16.504 | המשנה |  |
| 6:16.505 | פי | ' |
| 6:16.506 | רש"י |  |
| 6:16.507 | שעשאה |  |
| 6:16.508 | כפי |  |
| 6:16.509 | שאינה |  |
| 6:16.510 | עיקר | . |
| 6:16.511 | או |  |
| 6:16.512 | שאינו |  |
| 6:16.513 | מחזר |  |
| 6:16.514 | לשנותה | , |
| 6:16.515 | סנהדרין |  |
| 6:16.516 | דף |  |
| 6:16.517 | צ"ט |  |
| 6:16.518 | ע"ש |  |
| 6:16.519 | ברש"י |  |
| 6:16.520 | ומהרש"א | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

