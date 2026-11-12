# Translation pack: milhamot-hashem perek-5-20 (Milchamot ha-Shem 5:29–5:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-20`.

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
- **אדנות** as a name → "Adonai" (note: "lordship").
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
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
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ת"ל** → "Scripture says".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **מיתות**: פ' ד' מיתות (פ' + a chapter title) → "Four Death Penalties".
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


## Milchamot ha-Shem 5:29

Hebrew (Wikisource): נה) הן אלה קצות דרכי המקובלים החדשים בענין האלהות, אשר יראה הרואה שהם מתנגדים לגמרי לדברי רבותינו בבראשית רבה וילקוט המובא לעיל סי' ל"א ומדרש רבא פרשת יתרו וירושלמי שבת פרק במה אשה המובא בקונדריס שקדם דמלך בשר ודם שמלך יש לו אב או אח או בן וכו'. והאמת כדברי רז"ל כי אלהינו בהו"ש אחד ואין יחיד כיחודו, עוד כ' הרקח בשורש תשובה וז"ל אלו הן המינים האומר עולם כמנהגו נוהג ואין אלוה בעולם והאומר אין בורא קדמון לכל ( זהו דעת הפילסוף מחבר הזהר דמ"ל שעתיק נגלה לרשב"י והתיר לו לגלות הסוד שהאלוהות עלות וסבות רבות זו למעלה מזו ושאבא הוא הבורא והוא שאמר יהי אור, ויהי רקיע, ויקוו המים, ויהי מארת, אבל נעשה אדם אמרה אימא. ואדם קדמון שהוא למעלה מהם אמר ראו עתה כי אני אני הוא ואין אלהים עמדי) או האומר יש אלוה אחד עמו. (גם בזה שגגו המקובלים ונפתו למחבר הזהר שחושבים לכל אחד מן הפרצופים אלוה הוא ואשתו כדלעיל סי' כ' וכ"א וכ"ד וגם מאמינים במציאות עשר ספירות בקליפה הנקרא בשם אל אחד) או האומר יש לו דמות וצורה כאדם וגוף ותמונה (זה ג"כ דעת המקובלים שאמרו שזהו סוד ויברא אלהים את האדם בצלמו) או משתף דבר אחר עמו (ויתבאר לקטן סי' נ"ח איך המקובלים משתפים, וכבר ביארתי לעיל סי' מ"ח שהמקובלים ממאנים למור ד' בתים הראשונים שביגדל אלהים חי) או האומר שאין שם נבואה בעולם מן השי"ת או מכחיש תורת משה רבינו אפי' חבור אחד או דקדוק אחד או אומר אין הבורא יודע עשת מחשבות האדם, או אין האדם נדון לפי מעשיו (זה אומרים המקובלים על עתיק ואריך ופרצופים שלמעלה מהם. שאינם דנים את האדם לפי מעשיו. ואין אדם יוצא חייב בדיניהם כצדיק כרשע לפי שלא נתכנו להם עלילות, רק לז"א נתכנו עלילות סי' מ"ד ומ"ה). או שאומר אין העולם הבא ואין תשלום שכר, אלו המינים והאפיקורסין שמורידין אותן ולא מעלין אם לא עשו תשובה עכ"ל הרקח. ובספר מצות גדול לאוין א' המצוה הראשונה שלא לעולות במחשבה שיש שם אלוה זולתי השי"ת שנ' לא יהיה לך אלהים אחרים על פני, ותניא במכלתא לפי שנ' לא תעשה לך פסל. העשוי כבר מנין שלא יקיים ת"ל לא יהיה לך. ואין מקרא יוצא מדי פשוטו. ובפרק ד' מיתות אמרו שם כל המשתף שם שמים ודבר אחד נעקר מן העולם שנ' בלתי לה' לבדו. והמקרא הזה בא לעונש. וזה לאזהרה, ואומר ר' יעקב כי מה שנ' חרב לה' ולגדעון, אין זה שתוף לאלהות ולאדנות עכ"ל הסמ"ג.

| id | word | punctuation after |
|---|---|---|
| 5:29.1 | נה | ) |
| 5:29.2 | הן |  |
| 5:29.3 | אלה |  |
| 5:29.4 | קצות |  |
| 5:29.5 | דרכי |  |
| 5:29.6 | המקובלים |  |
| 5:29.7 | החדשים |  |
| 5:29.8 | בענין |  |
| 5:29.9 | האלהות | , |
| 5:29.10 | אשר |  |
| 5:29.11 | יראה |  |
| 5:29.12 | הרואה |  |
| 5:29.13 | שהם |  |
| 5:29.14 | מתנגדים |  |
| 5:29.15 | לגמרי |  |
| 5:29.16 | לדברי |  |
| 5:29.17 | רבותינו |  |
| 5:29.18 | בבראשית |  |
| 5:29.19 | רבה |  |
| 5:29.20 | וילקוט |  |
| 5:29.21 | המובא |  |
| 5:29.22 | לעיל |  |
| 5:29.23 | סי | ' |
| 5:29.24 | ל"א |  |
| 5:29.25 | ומדרש |  |
| 5:29.26 | רבא |  |
| 5:29.27 | פרשת |  |
| 5:29.28 | יתרו |  |
| 5:29.29 | וירושלמי |  |
| 5:29.30 | שבת |  |
| 5:29.31 | פרק |  |
| 5:29.32 | במה |  |
| 5:29.33 | אשה |  |
| 5:29.34 | המובא |  |
| 5:29.35 | בקונדריס |  |
| 5:29.36 | שקדם |  |
| 5:29.37 | דמלך |  |
| 5:29.38 | בשר |  |
| 5:29.39 | ודם |  |
| 5:29.40 | שמלך |  |
| 5:29.41 | יש |  |
| 5:29.42 | לו |  |
| 5:29.43 | אב |  |
| 5:29.44 | או |  |
| 5:29.45 | אח |  |
| 5:29.46 | או |  |
| 5:29.47 | בן |  |
| 5:29.48 | וכו | '. |
| 5:29.49 | והאמת |  |
| 5:29.50 | כדברי |  |
| 5:29.51 | רז"ל |  |
| 5:29.52 | כי |  |
| 5:29.53 | אלהינו |  |
| 5:29.54 | בהו"ש |  |
| 5:29.55 | אחד |  |
| 5:29.56 | ואין |  |
| 5:29.57 | יחיד |  |
| 5:29.58 | כיחודו | , |
| 5:29.59 | עוד |  |
| 5:29.60 | כ | ' |
| 5:29.61 | הרקח |  |
| 5:29.62 | בשורש |  |
| 5:29.63 | תשובה |  |
| 5:29.64 | וז"ל |  |
| 5:29.65 | אלו |  |
| 5:29.66 | הן |  |
| 5:29.67 | המינים |  |
| 5:29.68 | האומר |  |
| 5:29.69 | עולם |  |
| 5:29.70 | כמנהגו |  |
| 5:29.71 | נוהג |  |
| 5:29.72 | ואין |  |
| 5:29.73 | אלוה |  |
| 5:29.74 | בעולם |  |
| 5:29.75 | והאומר |  |
| 5:29.76 | אין |  |
| 5:29.77 | בורא |  |
| 5:29.78 | קדמון |  |
| 5:29.79 | לכל | ( |
| 5:29.80 | זהו |  |
| 5:29.81 | דעת |  |
| 5:29.82 | הפילסוף |  |
| 5:29.83 | מחבר |  |
| 5:29.84 | הזהר |  |
| 5:29.85 | דמ"ל |  |
| 5:29.86 | שעתיק |  |
| 5:29.87 | נגלה |  |
| 5:29.88 | לרשב"י |  |
| 5:29.89 | והתיר |  |
| 5:29.90 | לו |  |
| 5:29.91 | לגלות |  |
| 5:29.92 | הסוד |  |
| 5:29.93 | שהאלוהות |  |
| 5:29.94 | עלות |  |
| 5:29.95 | וסבות |  |
| 5:29.96 | רבות |  |
| 5:29.97 | זו |  |
| 5:29.98 | למעלה |  |
| 5:29.99 | מזו |  |
| 5:29.100 | ושאבא |  |
| 5:29.101 | הוא |  |
| 5:29.102 | הבורא |  |
| 5:29.103 | והוא |  |
| 5:29.104 | שאמר |  |
| 5:29.105 | יהי |  |
| 5:29.106 | אור | , |
| 5:29.107 | ויהי |  |
| 5:29.108 | רקיע | , |
| 5:29.109 | ויקוו |  |
| 5:29.110 | המים | , |
| 5:29.111 | ויהי |  |
| 5:29.112 | מארת | , |
| 5:29.113 | אבל |  |
| 5:29.114 | נעשה |  |
| 5:29.115 | אדם |  |
| 5:29.116 | אמרה |  |
| 5:29.117 | אימא | . |
| 5:29.118 | ואדם |  |
| 5:29.119 | קדמון |  |
| 5:29.120 | שהוא |  |
| 5:29.121 | למעלה |  |
| 5:29.122 | מהם |  |
| 5:29.123 | אמר |  |
| 5:29.124 | ראו |  |
| 5:29.125 | עתה |  |
| 5:29.126 | כי |  |
| 5:29.127 | אני |  |
| 5:29.128 | אני |  |
| 5:29.129 | הוא |  |
| 5:29.130 | ואין |  |
| 5:29.131 | אלהים |  |
| 5:29.132 | עמדי | ) |
| 5:29.133 | או |  |
| 5:29.134 | האומר |  |
| 5:29.135 | יש |  |
| 5:29.136 | אלוה |  |
| 5:29.137 | אחד |  |
| 5:29.138 | עמו | . |
| 5:29.139 | גם |  |
| 5:29.140 | בזה |  |
| 5:29.141 | שגגו |  |
| 5:29.142 | המקובלים |  |
| 5:29.143 | ונפתו |  |
| 5:29.144 | למחבר |  |
| 5:29.145 | הזהר |  |
| 5:29.146 | שחושבים |  |
| 5:29.147 | לכל |  |
| 5:29.148 | אחד |  |
| 5:29.149 | מן |  |
| 5:29.150 | הפרצופים |  |
| 5:29.151 | אלוה |  |
| 5:29.152 | הוא |  |
| 5:29.153 | ואשתו |  |
| 5:29.154 | כדלעיל |  |
| 5:29.155 | סי | ' |
| 5:29.156 | כ | ' |
| 5:29.157 | וכ"א |  |
| 5:29.158 | וכ"ד |  |
| 5:29.159 | וגם |  |
| 5:29.160 | מאמינים |  |
| 5:29.161 | במציאות |  |
| 5:29.162 | עשר |  |
| 5:29.163 | ספירות |  |
| 5:29.164 | בקליפה |  |
| 5:29.165 | הנקרא |  |
| 5:29.166 | בשם |  |
| 5:29.167 | אל |  |
| 5:29.168 | אחד | ) |
| 5:29.169 | או |  |
| 5:29.170 | האומר |  |
| 5:29.171 | יש |  |
| 5:29.172 | לו |  |
| 5:29.173 | דמות |  |
| 5:29.174 | וצורה |  |
| 5:29.175 | כאדם |  |
| 5:29.176 | וגוף |  |
| 5:29.177 | ותמונה |  |
| 5:29.178 | זה |  |
| 5:29.179 | ג"כ |  |
| 5:29.180 | דעת |  |
| 5:29.181 | המקובלים |  |
| 5:29.182 | שאמרו |  |
| 5:29.183 | שזהו |  |
| 5:29.184 | סוד |  |
| 5:29.185 | ויברא |  |
| 5:29.186 | אלהים |  |
| 5:29.187 | את |  |
| 5:29.188 | האדם |  |
| 5:29.189 | בצלמו | ) |
| 5:29.190 | או |  |
| 5:29.191 | משתף |  |
| 5:29.192 | דבר |  |
| 5:29.193 | אחר |  |
| 5:29.194 | עמו |  |
| 5:29.195 | ויתבאר |  |
| 5:29.196 | לקטן |  |
| 5:29.197 | סי | ' |
| 5:29.198 | נ"ח |  |
| 5:29.199 | איך |  |
| 5:29.200 | המקובלים |  |
| 5:29.201 | משתפים | , |
| 5:29.202 | וכבר |  |
| 5:29.203 | ביארתי |  |
| 5:29.204 | לעיל |  |
| 5:29.205 | סי | ' |
| 5:29.206 | מ"ח |  |
| 5:29.207 | שהמקובלים |  |
| 5:29.208 | ממאנים |  |
| 5:29.209 | למור |  |
| 5:29.210 | ד | ' |
| 5:29.211 | בתים |  |
| 5:29.212 | הראשונים |  |
| 5:29.213 | שביגדל |  |
| 5:29.214 | אלהים |  |
| 5:29.215 | חי | ) |
| 5:29.216 | או |  |
| 5:29.217 | האומר |  |
| 5:29.218 | שאין |  |
| 5:29.219 | שם |  |
| 5:29.220 | נבואה |  |
| 5:29.221 | בעולם |  |
| 5:29.222 | מן |  |
| 5:29.223 | השי"ת |  |
| 5:29.224 | או |  |
| 5:29.225 | מכחיש |  |
| 5:29.226 | תורת |  |
| 5:29.227 | משה |  |
| 5:29.228 | רבינו |  |
| 5:29.229 | אפי | ' |
| 5:29.230 | חבור |  |
| 5:29.231 | אחד |  |
| 5:29.232 | או |  |
| 5:29.233 | דקדוק |  |
| 5:29.234 | אחד |  |
| 5:29.235 | או |  |
| 5:29.236 | אומר |  |
| 5:29.237 | אין |  |
| 5:29.238 | הבורא |  |
| 5:29.239 | יודע |  |
| 5:29.240 | עשת |  |
| 5:29.241 | מחשבות |  |
| 5:29.242 | האדם | , |
| 5:29.243 | או |  |
| 5:29.244 | אין |  |
| 5:29.245 | האדם |  |
| 5:29.246 | נדון |  |
| 5:29.247 | לפי |  |
| 5:29.248 | מעשיו |  |
| 5:29.249 | זה |  |
| 5:29.250 | אומרים |  |
| 5:29.251 | המקובלים |  |
| 5:29.252 | על |  |
| 5:29.253 | עתיק |  |
| 5:29.254 | ואריך |  |
| 5:29.255 | ופרצופים |  |
| 5:29.256 | שלמעלה |  |
| 5:29.257 | מהם | . |
| 5:29.258 | שאינם |  |
| 5:29.259 | דנים |  |
| 5:29.260 | את |  |
| 5:29.261 | האדם |  |
| 5:29.262 | לפי |  |
| 5:29.263 | מעשיו | . |
| 5:29.264 | ואין |  |
| 5:29.265 | אדם |  |
| 5:29.266 | יוצא |  |
| 5:29.267 | חייב |  |
| 5:29.268 | בדיניהם |  |
| 5:29.269 | כצדיק |  |
| 5:29.270 | כרשע |  |
| 5:29.271 | לפי |  |
| 5:29.272 | שלא |  |
| 5:29.273 | נתכנו |  |
| 5:29.274 | להם |  |
| 5:29.275 | עלילות | , |
| 5:29.276 | רק |  |
| 5:29.277 | לז"א |  |
| 5:29.278 | נתכנו |  |
| 5:29.279 | עלילות |  |
| 5:29.280 | סי | ' |
| 5:29.281 | מ"ד |  |
| 5:29.282 | ומ"ה | ). |
| 5:29.283 | או |  |
| 5:29.284 | שאומר |  |
| 5:29.285 | אין |  |
| 5:29.286 | העולם |  |
| 5:29.287 | הבא |  |
| 5:29.288 | ואין |  |
| 5:29.289 | תשלום |  |
| 5:29.290 | שכר | , |
| 5:29.291 | אלו |  |
| 5:29.292 | המינים |  |
| 5:29.293 | והאפיקורסין |  |
| 5:29.294 | שמורידין |  |
| 5:29.295 | אותן |  |
| 5:29.296 | ולא |  |
| 5:29.297 | מעלין |  |
| 5:29.298 | אם |  |
| 5:29.299 | לא |  |
| 5:29.300 | עשו |  |
| 5:29.301 | תשובה |  |
| 5:29.302 | עכ"ל |  |
| 5:29.303 | הרקח | . |
| 5:29.304 | ובספר |  |
| 5:29.305 | מצות |  |
| 5:29.306 | גדול |  |
| 5:29.307 | לאוין |  |
| 5:29.308 | א | ' |
| 5:29.309 | המצוה |  |
| 5:29.310 | הראשונה |  |
| 5:29.311 | שלא |  |
| 5:29.312 | לעולות |  |
| 5:29.313 | במחשבה |  |
| 5:29.314 | שיש |  |
| 5:29.315 | שם |  |
| 5:29.316 | אלוה |  |
| 5:29.317 | זולתי |  |
| 5:29.318 | השי"ת |  |
| 5:29.319 | שנ | ' |
| 5:29.320 | לא |  |
| 5:29.321 | יהיה |  |
| 5:29.322 | לך |  |
| 5:29.323 | אלהים |  |
| 5:29.324 | אחרים |  |
| 5:29.325 | על |  |
| 5:29.326 | פני | , |
| 5:29.327 | ותניא |  |
| 5:29.328 | במכלתא |  |
| 5:29.329 | לפי |  |
| 5:29.330 | שנ | ' |
| 5:29.331 | לא |  |
| 5:29.332 | תעשה |  |
| 5:29.333 | לך |  |
| 5:29.334 | פסל | . |
| 5:29.335 | העשוי |  |
| 5:29.336 | כבר |  |
| 5:29.337 | מנין |  |
| 5:29.338 | שלא |  |
| 5:29.339 | יקיים |  |
| 5:29.340 | ת"ל |  |
| 5:29.341 | לא |  |
| 5:29.342 | יהיה |  |
| 5:29.343 | לך | . |
| 5:29.344 | ואין |  |
| 5:29.345 | מקרא |  |
| 5:29.346 | יוצא |  |
| 5:29.347 | מדי |  |
| 5:29.348 | פשוטו | . |
| 5:29.349 | ובפרק |  |
| 5:29.350 | ד | ' |
| 5:29.351 | מיתות |  |
| 5:29.352 | אמרו |  |
| 5:29.353 | שם |  |
| 5:29.354 | כל |  |
| 5:29.355 | המשתף |  |
| 5:29.356 | שם |  |
| 5:29.357 | שמים |  |
| 5:29.358 | ודבר |  |
| 5:29.359 | אחד |  |
| 5:29.360 | נעקר |  |
| 5:29.361 | מן |  |
| 5:29.362 | העולם |  |
| 5:29.363 | שנ | ' |
| 5:29.364 | בלתי |  |
| 5:29.365 | לה | ' |
| 5:29.366 | לבדו | . |
| 5:29.367 | והמקרא |  |
| 5:29.368 | הזה |  |
| 5:29.369 | בא |  |
| 5:29.370 | לעונש | . |
| 5:29.371 | וזה |  |
| 5:29.372 | לאזהרה | , |
| 5:29.373 | ואומר |  |
| 5:29.374 | ר | ' |
| 5:29.375 | יעקב |  |
| 5:29.376 | כי |  |
| 5:29.377 | מה |  |
| 5:29.378 | שנ | ' |
| 5:29.379 | חרב |  |
| 5:29.380 | לה | ' |
| 5:29.381 | ולגדעון | , |
| 5:29.382 | אין |  |
| 5:29.383 | זה |  |
| 5:29.384 | שתוף |  |
| 5:29.385 | לאלהות |  |
| 5:29.386 | ולאדנות |  |
| 5:29.387 | עכ"ל |  |
| 5:29.388 | הסמ"ג | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

