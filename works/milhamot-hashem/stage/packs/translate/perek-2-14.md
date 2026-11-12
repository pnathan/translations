# Translation pack: milhamot-hashem perek-2-14 (Milchamot ha-Shem 2:30–2:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-14`.

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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרמב"ן / רמב"ן** → "Ramban" (authority, transliterated; note first in a
  unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **ריב"ש** → "Ribash" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **הנוצרים / הנצרים** → "the Christians".
- מאמיני **השלוש** → "the believers in the Trinity".
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **רחמים**: "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes
  of mercy", a different thirteen from the hermeneutical rules.
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- עבודה **אחרת** → "other worship" (note first in a unit).
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **עלה / עלול** → "cause / effect".
- **כוונה** (in prayer) → "intention".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **המתפלספים** → "would-be philosophers".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 2:30

Hebrew (Wikisource): ומה שכתבת שז"א בן לאבא לאמא והוא אב לעלולים שאחריו, ויש לו אחות נוקבא דיליה, הכל משל וכינוי בספירות. ומה שכתבת שאין לעבוד שום פרצוף, זה ידוע בפי כל המקובלים, ואין צריכים לאזהרתך. ומה שכתבת בשם הרמב"ם חמשה הן הנקראים מינים וכו' מי הוא האומר שאין שם אלוה, או שהם שנים, או שהוא בעל גוף, או שאינו לבדו האחד או שעובד זולתו כדי להיות מליץ, או האומר שמצא סמנים או האומר שנברא העולם יש מיש ח"ו. אלא שכבר עלתה טינא בלבך, ואתה מעיין הספרים הקדושים ולוקח לך הדברים כפשטן וסומך על הבנתך. ושכחת מה שאמר חכם החכמים, ואל בינתך אל תשען. ומפני שכבר הודעת לי שאתה מבין תתיר לי ספקותי אשר שלחתי אליך ואשר אני עתיד לשלוח. ומה שכתבת והבנת רבים וגדולים חכמי לב הפורשים ממנה ראשונים ואחרונים, ח"ו שיש שום חכם פורש ממנה. ואם יש פורש ממנה הוא עצמו מודה בפה מלא שהוא מחסרון ידיעה בה. ומה שכתבת כמו שכ' הריב"ש וחו"י, שקר אתה דובר עליהם. והם שני עדים כשרים להעיד עליך שאתה דובר שקרים. וז"ל הריב"ש בקיצור כי ר' פרץ הכהן לא היה מדבר ולא מחשיב באותן הספירות. גם שמע מפיו שר' שמשון מקינון היה אומר אני מתפלל לדעת זה התינוק כלומר להוציא מלב המקובלים שהם מתפללים פעם לספירה אחת ופעם לספירה אחרת. גם בתפלת י"ח יש להם בכל אחת ואחת כוונה לספירה ידועה. וכל זה הוא דבר זר מאוד בעיני מי שאינו מקובל כמותם. וחושבים שזה אמונת שניות. וששמע אחד מן המתפלספים מספר בגנות המקובלים והיה אומר הנצרים מאמיני השלוש והמקובלים מאמיני העשיריות. ושאל להחכם הישיש דון יוסף ן' שושן והוא היה חכם בתלמוד. וראה פילוסופיא והיה מקובל וחסיד וגדול ומדקדק במצות. ואמר לו איך אתם המקובלים מכוונים בברכה אחת לספירה ידועה ובברכה אחת לספירה אחרת. ועוד הכי יש אלוהות לספירות? שיתפלל אדם להם? וענו לו חלילה שתהיה התפלה כי אם לשי"ת עילת העילות אלא שמכוין המחשבה להמשיך השפע לאותה ספירה המתיחסת לאותו דבר שהוא מבקש עליו כמו שתאמר שבברכת על הצדיקים יכוין לספירה הנקראת חסד שהיא מדת רחמים וכו'. והביא לו משל על זה. ענה ואמר והנה טוב מאד. ואמר איני תוקע עצמי באותה חכמה אחר שלא קיבלתיה מפי חכם מקובל. ואם ראיתי באורים על סודות הרמב"ן. הם מגלים טפח ומכסים טפחים. ואמר שחכם אחד פירש הקבלה שלא כד' הרמב"ן. ולכן אמר שדברי החכם ההוא יהיו לו לבדו. וחתם דבריו ולזה אני אומר שאין לסמוך בדברים כאלו אלא מפי חכם מקובל.

| id | word | punctuation after |
|---|---|---|
| 2:30.1 | ומה |  |
| 2:30.2 | שכתבת |  |
| 2:30.3 | שז"א |  |
| 2:30.4 | בן |  |
| 2:30.5 | לאבא |  |
| 2:30.6 | לאמא |  |
| 2:30.7 | והוא |  |
| 2:30.8 | אב |  |
| 2:30.9 | לעלולים |  |
| 2:30.10 | שאחריו | , |
| 2:30.11 | ויש |  |
| 2:30.12 | לו |  |
| 2:30.13 | אחות |  |
| 2:30.14 | נוקבא |  |
| 2:30.15 | דיליה | , |
| 2:30.16 | הכל |  |
| 2:30.17 | משל |  |
| 2:30.18 | וכינוי |  |
| 2:30.19 | בספירות | . |
| 2:30.20 | ומה |  |
| 2:30.21 | שכתבת |  |
| 2:30.22 | שאין |  |
| 2:30.23 | לעבוד |  |
| 2:30.24 | שום |  |
| 2:30.25 | פרצוף | , |
| 2:30.26 | זה |  |
| 2:30.27 | ידוע |  |
| 2:30.28 | בפי |  |
| 2:30.29 | כל |  |
| 2:30.30 | המקובלים | , |
| 2:30.31 | ואין |  |
| 2:30.32 | צריכים |  |
| 2:30.33 | לאזהרתך | . |
| 2:30.34 | ומה |  |
| 2:30.35 | שכתבת |  |
| 2:30.36 | בשם |  |
| 2:30.37 | הרמב"ם |  |
| 2:30.38 | חמשה |  |
| 2:30.39 | הן |  |
| 2:30.40 | הנקראים |  |
| 2:30.41 | מינים |  |
| 2:30.42 | וכו | ' |
| 2:30.43 | מי |  |
| 2:30.44 | הוא |  |
| 2:30.45 | האומר |  |
| 2:30.46 | שאין |  |
| 2:30.47 | שם |  |
| 2:30.48 | אלוה | , |
| 2:30.49 | או |  |
| 2:30.50 | שהם |  |
| 2:30.51 | שנים | , |
| 2:30.52 | או |  |
| 2:30.53 | שהוא |  |
| 2:30.54 | בעל |  |
| 2:30.55 | גוף | , |
| 2:30.56 | או |  |
| 2:30.57 | שאינו |  |
| 2:30.58 | לבדו |  |
| 2:30.59 | האחד |  |
| 2:30.60 | או |  |
| 2:30.61 | שעובד |  |
| 2:30.62 | זולתו |  |
| 2:30.63 | כדי |  |
| 2:30.64 | להיות |  |
| 2:30.65 | מליץ | , |
| 2:30.66 | או |  |
| 2:30.67 | האומר |  |
| 2:30.68 | שמצא |  |
| 2:30.69 | סמנים |  |
| 2:30.70 | או |  |
| 2:30.71 | האומר |  |
| 2:30.72 | שנברא |  |
| 2:30.73 | העולם |  |
| 2:30.74 | יש |  |
| 2:30.75 | מיש |  |
| 2:30.76 | ח"ו | . |
| 2:30.77 | אלא |  |
| 2:30.78 | שכבר |  |
| 2:30.79 | עלתה |  |
| 2:30.80 | טינא |  |
| 2:30.81 | בלבך | , |
| 2:30.82 | ואתה |  |
| 2:30.83 | מעיין |  |
| 2:30.84 | הספרים |  |
| 2:30.85 | הקדושים |  |
| 2:30.86 | ולוקח |  |
| 2:30.87 | לך |  |
| 2:30.88 | הדברים |  |
| 2:30.89 | כפשטן |  |
| 2:30.90 | וסומך |  |
| 2:30.91 | על |  |
| 2:30.92 | הבנתך | . |
| 2:30.93 | ושכחת |  |
| 2:30.94 | מה |  |
| 2:30.95 | שאמר |  |
| 2:30.96 | חכם |  |
| 2:30.97 | החכמים | , |
| 2:30.98 | ואל |  |
| 2:30.99 | בינתך |  |
| 2:30.100 | אל |  |
| 2:30.101 | תשען | . |
| 2:30.102 | ומפני |  |
| 2:30.103 | שכבר |  |
| 2:30.104 | הודעת |  |
| 2:30.105 | לי |  |
| 2:30.106 | שאתה |  |
| 2:30.107 | מבין |  |
| 2:30.108 | תתיר |  |
| 2:30.109 | לי |  |
| 2:30.110 | ספקותי |  |
| 2:30.111 | אשר |  |
| 2:30.112 | שלחתי |  |
| 2:30.113 | אליך |  |
| 2:30.114 | ואשר |  |
| 2:30.115 | אני |  |
| 2:30.116 | עתיד |  |
| 2:30.117 | לשלוח | . |
| 2:30.118 | ומה |  |
| 2:30.119 | שכתבת |  |
| 2:30.120 | והבנת |  |
| 2:30.121 | רבים |  |
| 2:30.122 | וגדולים |  |
| 2:30.123 | חכמי |  |
| 2:30.124 | לב |  |
| 2:30.125 | הפורשים |  |
| 2:30.126 | ממנה |  |
| 2:30.127 | ראשונים |  |
| 2:30.128 | ואחרונים | , |
| 2:30.129 | ח"ו |  |
| 2:30.130 | שיש |  |
| 2:30.131 | שום |  |
| 2:30.132 | חכם |  |
| 2:30.133 | פורש |  |
| 2:30.134 | ממנה | . |
| 2:30.135 | ואם |  |
| 2:30.136 | יש |  |
| 2:30.137 | פורש |  |
| 2:30.138 | ממנה |  |
| 2:30.139 | הוא |  |
| 2:30.140 | עצמו |  |
| 2:30.141 | מודה |  |
| 2:30.142 | בפה |  |
| 2:30.143 | מלא |  |
| 2:30.144 | שהוא |  |
| 2:30.145 | מחסרון |  |
| 2:30.146 | ידיעה |  |
| 2:30.147 | בה | . |
| 2:30.148 | ומה |  |
| 2:30.149 | שכתבת |  |
| 2:30.150 | כמו |  |
| 2:30.151 | שכ | ' |
| 2:30.152 | הריב"ש |  |
| 2:30.153 | וחו"י | , |
| 2:30.154 | שקר |  |
| 2:30.155 | אתה |  |
| 2:30.156 | דובר |  |
| 2:30.157 | עליהם | . |
| 2:30.158 | והם |  |
| 2:30.159 | שני |  |
| 2:30.160 | עדים |  |
| 2:30.161 | כשרים |  |
| 2:30.162 | להעיד |  |
| 2:30.163 | עליך |  |
| 2:30.164 | שאתה |  |
| 2:30.165 | דובר |  |
| 2:30.166 | שקרים | . |
| 2:30.167 | וז"ל |  |
| 2:30.168 | הריב"ש |  |
| 2:30.169 | בקיצור |  |
| 2:30.170 | כי |  |
| 2:30.171 | ר | ' |
| 2:30.172 | פרץ |  |
| 2:30.173 | הכהן |  |
| 2:30.174 | לא |  |
| 2:30.175 | היה |  |
| 2:30.176 | מדבר |  |
| 2:30.177 | ולא |  |
| 2:30.178 | מחשיב |  |
| 2:30.179 | באותן |  |
| 2:30.180 | הספירות | . |
| 2:30.181 | גם |  |
| 2:30.182 | שמע |  |
| 2:30.183 | מפיו |  |
| 2:30.184 | שר | ' |
| 2:30.185 | שמשון |  |
| 2:30.186 | מקינון |  |
| 2:30.187 | היה |  |
| 2:30.188 | אומר |  |
| 2:30.189 | אני |  |
| 2:30.190 | מתפלל |  |
| 2:30.191 | לדעת |  |
| 2:30.192 | זה |  |
| 2:30.193 | התינוק |  |
| 2:30.194 | כלומר |  |
| 2:30.195 | להוציא |  |
| 2:30.196 | מלב |  |
| 2:30.197 | המקובלים |  |
| 2:30.198 | שהם |  |
| 2:30.199 | מתפללים |  |
| 2:30.200 | פעם |  |
| 2:30.201 | לספירה |  |
| 2:30.202 | אחת |  |
| 2:30.203 | ופעם |  |
| 2:30.204 | לספירה |  |
| 2:30.205 | אחרת | . |
| 2:30.206 | גם |  |
| 2:30.207 | בתפלת |  |
| 2:30.208 | י"ח |  |
| 2:30.209 | יש |  |
| 2:30.210 | להם |  |
| 2:30.211 | בכל |  |
| 2:30.212 | אחת |  |
| 2:30.213 | ואחת |  |
| 2:30.214 | כוונה |  |
| 2:30.215 | לספירה |  |
| 2:30.216 | ידועה | . |
| 2:30.217 | וכל |  |
| 2:30.218 | זה |  |
| 2:30.219 | הוא |  |
| 2:30.220 | דבר |  |
| 2:30.221 | זר |  |
| 2:30.222 | מאוד |  |
| 2:30.223 | בעיני |  |
| 2:30.224 | מי |  |
| 2:30.225 | שאינו |  |
| 2:30.226 | מקובל |  |
| 2:30.227 | כמותם | . |
| 2:30.228 | וחושבים |  |
| 2:30.229 | שזה |  |
| 2:30.230 | אמונת |  |
| 2:30.231 | שניות | . |
| 2:30.232 | וששמע |  |
| 2:30.233 | אחד |  |
| 2:30.234 | מן |  |
| 2:30.235 | המתפלספים |  |
| 2:30.236 | מספר |  |
| 2:30.237 | בגנות |  |
| 2:30.238 | המקובלים |  |
| 2:30.239 | והיה |  |
| 2:30.240 | אומר |  |
| 2:30.241 | הנצרים |  |
| 2:30.242 | מאמיני |  |
| 2:30.243 | השלוש |  |
| 2:30.244 | והמקובלים |  |
| 2:30.245 | מאמיני |  |
| 2:30.246 | העשיריות | . |
| 2:30.247 | ושאל |  |
| 2:30.248 | להחכם |  |
| 2:30.249 | הישיש |  |
| 2:30.250 | דון |  |
| 2:30.251 | יוסף |  |
| 2:30.252 | ן | ' |
| 2:30.253 | שושן |  |
| 2:30.254 | והוא |  |
| 2:30.255 | היה |  |
| 2:30.256 | חכם |  |
| 2:30.257 | בתלמוד | . |
| 2:30.258 | וראה |  |
| 2:30.259 | פילוסופיא |  |
| 2:30.260 | והיה |  |
| 2:30.261 | מקובל |  |
| 2:30.262 | וחסיד |  |
| 2:30.263 | וגדול |  |
| 2:30.264 | ומדקדק |  |
| 2:30.265 | במצות | . |
| 2:30.266 | ואמר |  |
| 2:30.267 | לו |  |
| 2:30.268 | איך |  |
| 2:30.269 | אתם |  |
| 2:30.270 | המקובלים |  |
| 2:30.271 | מכוונים |  |
| 2:30.272 | בברכה |  |
| 2:30.273 | אחת |  |
| 2:30.274 | לספירה |  |
| 2:30.275 | ידועה |  |
| 2:30.276 | ובברכה |  |
| 2:30.277 | אחת |  |
| 2:30.278 | לספירה |  |
| 2:30.279 | אחרת | . |
| 2:30.280 | ועוד |  |
| 2:30.281 | הכי |  |
| 2:30.282 | יש |  |
| 2:30.283 | אלוהות |  |
| 2:30.284 | לספירות | ? |
| 2:30.285 | שיתפלל |  |
| 2:30.286 | אדם |  |
| 2:30.287 | להם | ? |
| 2:30.288 | וענו |  |
| 2:30.289 | לו |  |
| 2:30.290 | חלילה |  |
| 2:30.291 | שתהיה |  |
| 2:30.292 | התפלה |  |
| 2:30.293 | כי |  |
| 2:30.294 | אם |  |
| 2:30.295 | לשי"ת |  |
| 2:30.296 | עילת |  |
| 2:30.297 | העילות |  |
| 2:30.298 | אלא |  |
| 2:30.299 | שמכוין |  |
| 2:30.300 | המחשבה |  |
| 2:30.301 | להמשיך |  |
| 2:30.302 | השפע |  |
| 2:30.303 | לאותה |  |
| 2:30.304 | ספירה |  |
| 2:30.305 | המתיחסת |  |
| 2:30.306 | לאותו |  |
| 2:30.307 | דבר |  |
| 2:30.308 | שהוא |  |
| 2:30.309 | מבקש |  |
| 2:30.310 | עליו |  |
| 2:30.311 | כמו |  |
| 2:30.312 | שתאמר |  |
| 2:30.313 | שבברכת |  |
| 2:30.314 | על |  |
| 2:30.315 | הצדיקים |  |
| 2:30.316 | יכוין |  |
| 2:30.317 | לספירה |  |
| 2:30.318 | הנקראת |  |
| 2:30.319 | חסד |  |
| 2:30.320 | שהיא |  |
| 2:30.321 | מדת |  |
| 2:30.322 | רחמים |  |
| 2:30.323 | וכו | '. |
| 2:30.324 | והביא |  |
| 2:30.325 | לו |  |
| 2:30.326 | משל |  |
| 2:30.327 | על |  |
| 2:30.328 | זה | . |
| 2:30.329 | ענה |  |
| 2:30.330 | ואמר |  |
| 2:30.331 | והנה |  |
| 2:30.332 | טוב |  |
| 2:30.333 | מאד | . |
| 2:30.334 | ואמר |  |
| 2:30.335 | איני |  |
| 2:30.336 | תוקע |  |
| 2:30.337 | עצמי |  |
| 2:30.338 | באותה |  |
| 2:30.339 | חכמה |  |
| 2:30.340 | אחר |  |
| 2:30.341 | שלא |  |
| 2:30.342 | קיבלתיה |  |
| 2:30.343 | מפי |  |
| 2:30.344 | חכם |  |
| 2:30.345 | מקובל | . |
| 2:30.346 | ואם |  |
| 2:30.347 | ראיתי |  |
| 2:30.348 | באורים |  |
| 2:30.349 | על |  |
| 2:30.350 | סודות |  |
| 2:30.351 | הרמב"ן | . |
| 2:30.352 | הם |  |
| 2:30.353 | מגלים |  |
| 2:30.354 | טפח |  |
| 2:30.355 | ומכסים |  |
| 2:30.356 | טפחים | . |
| 2:30.357 | ואמר |  |
| 2:30.358 | שחכם |  |
| 2:30.359 | אחד |  |
| 2:30.360 | פירש |  |
| 2:30.361 | הקבלה |  |
| 2:30.362 | שלא |  |
| 2:30.363 | כד | ' |
| 2:30.364 | הרמב"ן | . |
| 2:30.365 | ולכן |  |
| 2:30.366 | אמר |  |
| 2:30.367 | שדברי |  |
| 2:30.368 | החכם |  |
| 2:30.369 | ההוא |  |
| 2:30.370 | יהיו |  |
| 2:30.371 | לו |  |
| 2:30.372 | לבדו | . |
| 2:30.373 | וחתם |  |
| 2:30.374 | דבריו |  |
| 2:30.375 | ולזה |  |
| 2:30.376 | אני |  |
| 2:30.377 | אומר |  |
| 2:30.378 | שאין |  |
| 2:30.379 | לסמוך |  |
| 2:30.380 | בדברים |  |
| 2:30.381 | כאלו |  |
| 2:30.382 | אלא |  |
| 2:30.383 | מפי |  |
| 2:30.384 | חכם |  |
| 2:30.385 | מקובל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

