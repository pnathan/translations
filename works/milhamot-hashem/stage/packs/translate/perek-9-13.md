# Translation pack: milhamot-hashem perek-9-13 (Milchamot ha-Shem 9:16–9:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-13`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
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
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
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
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- תורה **שבעל** פה → "the Oral Torah".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **א"כ** → "if so".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
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
- **זכאה / דכאה / חולקיה / הולקיה / חולקיהון / חולקיהן**: זכאה חולקיה /
  חולקיהון → "happy is his / their portion"; זכאה חולקיהון דישראל → "happy is
  the portion of Israel". Quoted as sales talk (9:4, 9:16): in quotation
  marks, no irony added. Variant spellings ("הולקיה", "דכאה"): the evident
  word, recorded in `apparatus`.
- בחלק **שפתיו** → "with the smoothness of his lips".
- **מילדי** נכרים → "from the children of strangers".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 9:16

Hebrew (Wikisource): קלד) מכאן למדנו ממלתיה דרבא שכך היא דרכו של מסית שמפליג בשבח האלילים הרחוקים שכל היא מושכת השפע אליה, וכך חיה מטיבה לעובד אותה להשפיע מטובה שמשכה אליה לזה שעובדה ומשבחה ומפארה, וכך מריעה לפונים אליה עורף ולא השחין לה. ותורתינו הקדושה העירה אותנו באמרה הקרובים אליך או הרחוקים, שמהקרובים שהם פרצופי השמש והירח ושבתי וצדק ומאדים ונגה וכוכב, שבדו חוברי שמים עובדיהם כמו שכתב הרמב"ם ורבי עובדיה בפ' כל הצלמים שיחסו לכוכבים פרצופים, עד שאמרו פרצוף שבתי זקן שחור וישיש. ופרצוף החמה פרצוף מלך מוכתר יושב בעגלה, ופרצוף נגה נערה יפה לבושה חלי זהב, וכן מיחסים לכל הכוכבים והמזלות פרצופים רבים וכו'! מהפרצופים של הכוכבים הקרובים אלינו, נקיש ונלמד על הפרצופים של ספירות הרחוקים ממנו אשר בדא מחבר הזהר המסית אותנו לעבדם ולקרותם בשמותם של הקב"ה שאין בהם ממש! ועל כן לא נאבה לו ולא נשמע אליו, כמצות תורתינו הקדושה שלא לשמוע אל המסית שמטעים את דבריו ומשבח את העושים כמעשיו לשתף ולצרף עם אלהינו ב"ה את הפרצופים הנזכרים, ולהיות העבודה ביחוד לקצר האפים בצירוף שאר פרצופים. וליחדם בפיהם שהבל א', והוסיף על השלשה בשמע ישראל גם את את המלכות שהיא הרביעית עם השלשה שמות יהוה אלהינו יהוה. כדברי השלישים, ובמלת אחד מצורף את המלכוץ עמהם. ואחריו נמשכו המון המקובלים החדשים בתם לבבם ולא נתנו אל לבם להבין ששטתו זאת מילדי נכרים המתנגדים לאמונתינו ביחוד השי"ת, והאמינו לו בחלק שפתיו המורגל בפיו תמיד לומר זכאה חולקיה. דכאה חולקיהון דישראל, כדי למשוך ולגנוב דעת השומעים להדיחם, והבטיח שלעתיד קרוב לימות המשיח כלא יתפרנסון מחאי הבורא דיליה, שיבטלו כל הלכות של תורה שבעל פה, היפך דברי רז"ל שאמרו שכל הלכות של תורה של פה אינם בטילים לעולם (רמב"ם סוף הל' מגלה). ולא די זה אלא שהמסית הנזכר שם את רשב"י תנא דידן לעובר על דברי חכמים ח"ו שאמרו ולא במרכבה ואפי' ביחיד אלא א"כ היה חכם ומבין מדעתו מוסרין לו ראשי פרקים. ולפי דבריו שאלו הם מעשה מרכבה שהזכירו רז"ל, איך היה רשב"י דורש בהם בפומפי בין חביריו מה שלא התירו חכמים לעשות כן אפי' היו כולן חכמים גדולים? ואם תאמר שעתיקא קדושא נגלה אליו כמבואר בזהר בראשית, הא ודאי לא צייתנא ליה לעקור דבר מן התורה שקבלו חכמים, שאין נביא רשאי לחדש דבר מעתה, ובודאי נביא שקר הוא, כמו שמבואר בתלמוד בכמה מקומות שאם יבוא אליהו ויאמר וכו' אין שומעין לו, וזה אפי' כמידי דגזרו ביה רבנן, כ"ש וק"ו לאיסור עצמי שקבלו חכמים בו שאסור לעשות כן שאפילו יבוא אליהו ויאמר שמותר הוא שאין שומעין לו.

| id | word | punctuation after |
|---|---|---|
| 9:16.1 | קלד | ) |
| 9:16.2 | מכאן |  |
| 9:16.3 | למדנו |  |
| 9:16.4 | ממלתיה |  |
| 9:16.5 | דרבא |  |
| 9:16.6 | שכך |  |
| 9:16.7 | היא |  |
| 9:16.8 | דרכו |  |
| 9:16.9 | של |  |
| 9:16.10 | מסית |  |
| 9:16.11 | שמפליג |  |
| 9:16.12 | בשבח |  |
| 9:16.13 | האלילים |  |
| 9:16.14 | הרחוקים |  |
| 9:16.15 | שכל |  |
| 9:16.16 | היא |  |
| 9:16.17 | מושכת |  |
| 9:16.18 | השפע |  |
| 9:16.19 | אליה | , |
| 9:16.20 | וכך |  |
| 9:16.21 | חיה |  |
| 9:16.22 | מטיבה |  |
| 9:16.23 | לעובד |  |
| 9:16.24 | אותה |  |
| 9:16.25 | להשפיע |  |
| 9:16.26 | מטובה |  |
| 9:16.27 | שמשכה |  |
| 9:16.28 | אליה |  |
| 9:16.29 | לזה |  |
| 9:16.30 | שעובדה |  |
| 9:16.31 | ומשבחה |  |
| 9:16.32 | ומפארה | , |
| 9:16.33 | וכך |  |
| 9:16.34 | מריעה |  |
| 9:16.35 | לפונים |  |
| 9:16.36 | אליה |  |
| 9:16.37 | עורף |  |
| 9:16.38 | ולא |  |
| 9:16.39 | השחין |  |
| 9:16.40 | לה | . |
| 9:16.41 | ותורתינו |  |
| 9:16.42 | הקדושה |  |
| 9:16.43 | העירה |  |
| 9:16.44 | אותנו |  |
| 9:16.45 | באמרה |  |
| 9:16.46 | הקרובים |  |
| 9:16.47 | אליך |  |
| 9:16.48 | או |  |
| 9:16.49 | הרחוקים | , |
| 9:16.50 | שמהקרובים |  |
| 9:16.51 | שהם |  |
| 9:16.52 | פרצופי |  |
| 9:16.53 | השמש |  |
| 9:16.54 | והירח |  |
| 9:16.55 | ושבתי |  |
| 9:16.56 | וצדק |  |
| 9:16.57 | ומאדים |  |
| 9:16.58 | ונגה |  |
| 9:16.59 | וכוכב | , |
| 9:16.60 | שבדו |  |
| 9:16.61 | חוברי |  |
| 9:16.62 | שמים |  |
| 9:16.63 | עובדיהם |  |
| 9:16.64 | כמו |  |
| 9:16.65 | שכתב |  |
| 9:16.66 | הרמב"ם |  |
| 9:16.67 | ורבי |  |
| 9:16.68 | עובדיה |  |
| 9:16.69 | בפ | ' |
| 9:16.70 | כל |  |
| 9:16.71 | הצלמים |  |
| 9:16.72 | שיחסו |  |
| 9:16.73 | לכוכבים |  |
| 9:16.74 | פרצופים | , |
| 9:16.75 | עד |  |
| 9:16.76 | שאמרו |  |
| 9:16.77 | פרצוף |  |
| 9:16.78 | שבתי |  |
| 9:16.79 | זקן |  |
| 9:16.80 | שחור |  |
| 9:16.81 | וישיש | . |
| 9:16.82 | ופרצוף |  |
| 9:16.83 | החמה |  |
| 9:16.84 | פרצוף |  |
| 9:16.85 | מלך |  |
| 9:16.86 | מוכתר |  |
| 9:16.87 | יושב |  |
| 9:16.88 | בעגלה | , |
| 9:16.89 | ופרצוף |  |
| 9:16.90 | נגה |  |
| 9:16.91 | נערה |  |
| 9:16.92 | יפה |  |
| 9:16.93 | לבושה |  |
| 9:16.94 | חלי |  |
| 9:16.95 | זהב | , |
| 9:16.96 | וכן |  |
| 9:16.97 | מיחסים |  |
| 9:16.98 | לכל |  |
| 9:16.99 | הכוכבים |  |
| 9:16.100 | והמזלות |  |
| 9:16.101 | פרצופים |  |
| 9:16.102 | רבים |  |
| 9:16.103 | וכו | '! |
| 9:16.104 | מהפרצופים |  |
| 9:16.105 | של |  |
| 9:16.106 | הכוכבים |  |
| 9:16.107 | הקרובים |  |
| 9:16.108 | אלינו | , |
| 9:16.109 | נקיש |  |
| 9:16.110 | ונלמד |  |
| 9:16.111 | על |  |
| 9:16.112 | הפרצופים |  |
| 9:16.113 | של |  |
| 9:16.114 | ספירות |  |
| 9:16.115 | הרחוקים |  |
| 9:16.116 | ממנו |  |
| 9:16.117 | אשר |  |
| 9:16.118 | בדא |  |
| 9:16.119 | מחבר |  |
| 9:16.120 | הזהר |  |
| 9:16.121 | המסית |  |
| 9:16.122 | אותנו |  |
| 9:16.123 | לעבדם |  |
| 9:16.124 | ולקרותם |  |
| 9:16.125 | בשמותם |  |
| 9:16.126 | של |  |
| 9:16.127 | הקב"ה |  |
| 9:16.128 | שאין |  |
| 9:16.129 | בהם |  |
| 9:16.130 | ממש | ! |
| 9:16.131 | ועל |  |
| 9:16.132 | כן |  |
| 9:16.133 | לא |  |
| 9:16.134 | נאבה |  |
| 9:16.135 | לו |  |
| 9:16.136 | ולא |  |
| 9:16.137 | נשמע |  |
| 9:16.138 | אליו | , |
| 9:16.139 | כמצות |  |
| 9:16.140 | תורתינו |  |
| 9:16.141 | הקדושה |  |
| 9:16.142 | שלא |  |
| 9:16.143 | לשמוע |  |
| 9:16.144 | אל |  |
| 9:16.145 | המסית |  |
| 9:16.146 | שמטעים |  |
| 9:16.147 | את |  |
| 9:16.148 | דבריו |  |
| 9:16.149 | ומשבח |  |
| 9:16.150 | את |  |
| 9:16.151 | העושים |  |
| 9:16.152 | כמעשיו |  |
| 9:16.153 | לשתף |  |
| 9:16.154 | ולצרף |  |
| 9:16.155 | עם |  |
| 9:16.156 | אלהינו |  |
| 9:16.157 | ב"ה |  |
| 9:16.158 | את |  |
| 9:16.159 | הפרצופים |  |
| 9:16.160 | הנזכרים | , |
| 9:16.161 | ולהיות |  |
| 9:16.162 | העבודה |  |
| 9:16.163 | ביחוד |  |
| 9:16.164 | לקצר |  |
| 9:16.165 | האפים |  |
| 9:16.166 | בצירוף |  |
| 9:16.167 | שאר |  |
| 9:16.168 | פרצופים | . |
| 9:16.169 | וליחדם |  |
| 9:16.170 | בפיהם |  |
| 9:16.171 | שהבל |  |
| 9:16.172 | א | ', |
| 9:16.173 | והוסיף |  |
| 9:16.174 | על |  |
| 9:16.175 | השלשה |  |
| 9:16.176 | בשמע |  |
| 9:16.177 | ישראל |  |
| 9:16.178 | גם |  |
| 9:16.179 | את |  |
| 9:16.180 | את |  |
| 9:16.181 | המלכות |  |
| 9:16.182 | שהיא |  |
| 9:16.183 | הרביעית |  |
| 9:16.184 | עם |  |
| 9:16.185 | השלשה |  |
| 9:16.186 | שמות |  |
| 9:16.187 | יהוה |  |
| 9:16.188 | אלהינו |  |
| 9:16.189 | יהוה | . |
| 9:16.190 | כדברי |  |
| 9:16.191 | השלישים | , |
| 9:16.192 | ובמלת |  |
| 9:16.193 | אחד |  |
| 9:16.194 | מצורף |  |
| 9:16.195 | את |  |
| 9:16.196 | המלכוץ |  |
| 9:16.197 | עמהם | . |
| 9:16.198 | ואחריו |  |
| 9:16.199 | נמשכו |  |
| 9:16.200 | המון |  |
| 9:16.201 | המקובלים |  |
| 9:16.202 | החדשים |  |
| 9:16.203 | בתם |  |
| 9:16.204 | לבבם |  |
| 9:16.205 | ולא |  |
| 9:16.206 | נתנו |  |
| 9:16.207 | אל |  |
| 9:16.208 | לבם |  |
| 9:16.209 | להבין |  |
| 9:16.210 | ששטתו |  |
| 9:16.211 | זאת |  |
| 9:16.212 | מילדי |  |
| 9:16.213 | נכרים |  |
| 9:16.214 | המתנגדים |  |
| 9:16.215 | לאמונתינו |  |
| 9:16.216 | ביחוד |  |
| 9:16.217 | השי"ת | , |
| 9:16.218 | והאמינו |  |
| 9:16.219 | לו |  |
| 9:16.220 | בחלק |  |
| 9:16.221 | שפתיו |  |
| 9:16.222 | המורגל |  |
| 9:16.223 | בפיו |  |
| 9:16.224 | תמיד |  |
| 9:16.225 | לומר |  |
| 9:16.226 | זכאה |  |
| 9:16.227 | חולקיה | . |
| 9:16.228 | דכאה |  |
| 9:16.229 | חולקיהון |  |
| 9:16.230 | דישראל | , |
| 9:16.231 | כדי |  |
| 9:16.232 | למשוך |  |
| 9:16.233 | ולגנוב |  |
| 9:16.234 | דעת |  |
| 9:16.235 | השומעים |  |
| 9:16.236 | להדיחם | , |
| 9:16.237 | והבטיח |  |
| 9:16.238 | שלעתיד |  |
| 9:16.239 | קרוב |  |
| 9:16.240 | לימות |  |
| 9:16.241 | המשיח |  |
| 9:16.242 | כלא |  |
| 9:16.243 | יתפרנסון |  |
| 9:16.244 | מחאי |  |
| 9:16.245 | הבורא |  |
| 9:16.246 | דיליה | , |
| 9:16.247 | שיבטלו |  |
| 9:16.248 | כל |  |
| 9:16.249 | הלכות |  |
| 9:16.250 | של |  |
| 9:16.251 | תורה |  |
| 9:16.252 | שבעל |  |
| 9:16.253 | פה | , |
| 9:16.254 | היפך |  |
| 9:16.255 | דברי |  |
| 9:16.256 | רז"ל |  |
| 9:16.257 | שאמרו |  |
| 9:16.258 | שכל |  |
| 9:16.259 | הלכות |  |
| 9:16.260 | של |  |
| 9:16.261 | תורה |  |
| 9:16.262 | של |  |
| 9:16.263 | פה |  |
| 9:16.264 | אינם |  |
| 9:16.265 | בטילים |  |
| 9:16.266 | לעולם |  |
| 9:16.267 | רמב"ם |  |
| 9:16.268 | סוף |  |
| 9:16.269 | הל | ' |
| 9:16.270 | מגלה | ). |
| 9:16.271 | ולא |  |
| 9:16.272 | די |  |
| 9:16.273 | זה |  |
| 9:16.274 | אלא |  |
| 9:16.275 | שהמסית |  |
| 9:16.276 | הנזכר |  |
| 9:16.277 | שם |  |
| 9:16.278 | את |  |
| 9:16.279 | רשב"י |  |
| 9:16.280 | תנא |  |
| 9:16.281 | דידן |  |
| 9:16.282 | לעובר |  |
| 9:16.283 | על |  |
| 9:16.284 | דברי |  |
| 9:16.285 | חכמים |  |
| 9:16.286 | ח"ו |  |
| 9:16.287 | שאמרו |  |
| 9:16.288 | ולא |  |
| 9:16.289 | במרכבה |  |
| 9:16.290 | ואפי | ' |
| 9:16.291 | ביחיד |  |
| 9:16.292 | אלא |  |
| 9:16.293 | א"כ |  |
| 9:16.294 | היה |  |
| 9:16.295 | חכם |  |
| 9:16.296 | ומבין |  |
| 9:16.297 | מדעתו |  |
| 9:16.298 | מוסרין |  |
| 9:16.299 | לו |  |
| 9:16.300 | ראשי |  |
| 9:16.301 | פרקים | . |
| 9:16.302 | ולפי |  |
| 9:16.303 | דבריו |  |
| 9:16.304 | שאלו |  |
| 9:16.305 | הם |  |
| 9:16.306 | מעשה |  |
| 9:16.307 | מרכבה |  |
| 9:16.308 | שהזכירו |  |
| 9:16.309 | רז"ל | , |
| 9:16.310 | איך |  |
| 9:16.311 | היה |  |
| 9:16.312 | רשב"י |  |
| 9:16.313 | דורש |  |
| 9:16.314 | בהם |  |
| 9:16.315 | בפומפי |  |
| 9:16.316 | בין |  |
| 9:16.317 | חביריו |  |
| 9:16.318 | מה |  |
| 9:16.319 | שלא |  |
| 9:16.320 | התירו |  |
| 9:16.321 | חכמים |  |
| 9:16.322 | לעשות |  |
| 9:16.323 | כן |  |
| 9:16.324 | אפי | ' |
| 9:16.325 | היו |  |
| 9:16.326 | כולן |  |
| 9:16.327 | חכמים |  |
| 9:16.328 | גדולים | ? |
| 9:16.329 | ואם |  |
| 9:16.330 | תאמר |  |
| 9:16.331 | שעתיקא |  |
| 9:16.332 | קדושא |  |
| 9:16.333 | נגלה |  |
| 9:16.334 | אליו |  |
| 9:16.335 | כמבואר |  |
| 9:16.336 | בזהר |  |
| 9:16.337 | בראשית | , |
| 9:16.338 | הא |  |
| 9:16.339 | ודאי |  |
| 9:16.340 | לא |  |
| 9:16.341 | צייתנא |  |
| 9:16.342 | ליה |  |
| 9:16.343 | לעקור |  |
| 9:16.344 | דבר |  |
| 9:16.345 | מן |  |
| 9:16.346 | התורה |  |
| 9:16.347 | שקבלו |  |
| 9:16.348 | חכמים | , |
| 9:16.349 | שאין |  |
| 9:16.350 | נביא |  |
| 9:16.351 | רשאי |  |
| 9:16.352 | לחדש |  |
| 9:16.353 | דבר |  |
| 9:16.354 | מעתה | , |
| 9:16.355 | ובודאי |  |
| 9:16.356 | נביא |  |
| 9:16.357 | שקר |  |
| 9:16.358 | הוא | , |
| 9:16.359 | כמו |  |
| 9:16.360 | שמבואר |  |
| 9:16.361 | בתלמוד |  |
| 9:16.362 | בכמה |  |
| 9:16.363 | מקומות |  |
| 9:16.364 | שאם |  |
| 9:16.365 | יבוא |  |
| 9:16.366 | אליהו |  |
| 9:16.367 | ויאמר |  |
| 9:16.368 | וכו | ' |
| 9:16.369 | אין |  |
| 9:16.370 | שומעין |  |
| 9:16.371 | לו | , |
| 9:16.372 | וזה |  |
| 9:16.373 | אפי | ' |
| 9:16.374 | כמידי |  |
| 9:16.375 | דגזרו |  |
| 9:16.376 | ביה |  |
| 9:16.377 | רבנן | , |
| 9:16.378 | כ"ש |  |
| 9:16.379 | וק"ו |  |
| 9:16.380 | לאיסור |  |
| 9:16.381 | עצמי |  |
| 9:16.382 | שקבלו |  |
| 9:16.383 | חכמים |  |
| 9:16.384 | בו |  |
| 9:16.385 | שאסור |  |
| 9:16.386 | לעשות |  |
| 9:16.387 | כן |  |
| 9:16.388 | שאפילו |  |
| 9:16.389 | יבוא |  |
| 9:16.390 | אליהו |  |
| 9:16.391 | ויאמר |  |
| 9:16.392 | שמותר |  |
| 9:16.393 | הוא |  |
| 9:16.394 | שאין |  |
| 9:16.395 | שומעין |  |
| 9:16.396 | לו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

