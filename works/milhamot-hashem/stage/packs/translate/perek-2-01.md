# Translation pack: milhamot-hashem perek-2-01 (Milchamot ha-Shem 2:1–2:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-01`.

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
- **מלחמות ה׳** → *The Wars of the LORD* (Num 21:14; Q1); transliterated
  *Milchamot ha-Shem*.
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
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
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רוקח** → "Roke'ach" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **עה"ת** → "on the Torah"; after an unidentified book's name → "a book on the
  Torah" (do not guess Etz Chayyim).
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- **השפע** → "the flow".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ת"ל** → "Scripture says".
- **סתו"מ** → "Torah scrolls, tefillin and mezuzot" (note).
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אחד ... **ואחד** of two kinds of case → "alike".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **ידידי** → "my friend" (second person kept).
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
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 2:1

Hebrew (Wikisource): מלחמות ה'

| id | word | punctuation after |
|---|---|---|
| 2:1.1 | מלחמות |  |
| 2:1.2 | ה | ' |


## Milchamot ha-Shem 2:2

Hebrew (Wikisource): שאול שאל אחד מן התלמידים וז"ל יודיענו מ"ו העבודה והתפלה שאנו מתפללים למי שייכא. לעילת העילות ב"ו או לנאצל אשר נאצל ממנו (ז"א) כאשר פירש לנו באוזניכם מ"ו סאלם ן' דוד שנסע מכאן לארץ הקדש ע"פ האמונה האמתית כמפורש בס' מצרף האמונה ובס' עה"ת שהעובד לאין סוף עובד ללא אלהי אמת וללא אלהי משפט כדי שנדע את מי נעבוד ותשובתך מהרה אדוני ע"כ.

| id | word | punctuation after |
|---|---|---|
| 2:2.1 | שאול |  |
| 2:2.2 | שאל |  |
| 2:2.3 | אחד |  |
| 2:2.4 | מן |  |
| 2:2.5 | התלמידים |  |
| 2:2.6 | וז"ל |  |
| 2:2.7 | יודיענו |  |
| 2:2.8 | מ"ו |  |
| 2:2.9 | העבודה |  |
| 2:2.10 | והתפלה |  |
| 2:2.11 | שאנו |  |
| 2:2.12 | מתפללים |  |
| 2:2.13 | למי |  |
| 2:2.14 | שייכא | . |
| 2:2.15 | לעילת |  |
| 2:2.16 | העילות |  |
| 2:2.17 | ב"ו |  |
| 2:2.18 | או |  |
| 2:2.19 | לנאצל |  |
| 2:2.20 | אשר |  |
| 2:2.21 | נאצל |  |
| 2:2.22 | ממנו |  |
| 2:2.23 | ז"א | ) |
| 2:2.24 | כאשר |  |
| 2:2.25 | פירש |  |
| 2:2.26 | לנו |  |
| 2:2.27 | באוזניכם |  |
| 2:2.28 | מ"ו |  |
| 2:2.29 | סאלם |  |
| 2:2.30 | ן | ' |
| 2:2.31 | דוד |  |
| 2:2.32 | שנסע |  |
| 2:2.33 | מכאן |  |
| 2:2.34 | לארץ |  |
| 2:2.35 | הקדש |  |
| 2:2.36 | ע"פ |  |
| 2:2.37 | האמונה |  |
| 2:2.38 | האמתית |  |
| 2:2.39 | כמפורש |  |
| 2:2.40 | בס | ' |
| 2:2.41 | מצרף |  |
| 2:2.42 | האמונה |  |
| 2:2.43 | ובס | ' |
| 2:2.44 | עה"ת |  |
| 2:2.45 | שהעובד |  |
| 2:2.46 | לאין |  |
| 2:2.47 | סוף |  |
| 2:2.48 | עובד |  |
| 2:2.49 | ללא |  |
| 2:2.50 | אלהי |  |
| 2:2.51 | אמת |  |
| 2:2.52 | וללא |  |
| 2:2.53 | אלהי |  |
| 2:2.54 | משפט |  |
| 2:2.55 | כדי |  |
| 2:2.56 | שנדע |  |
| 2:2.57 | את |  |
| 2:2.58 | מי |  |
| 2:2.59 | נעבוד |  |
| 2:2.60 | ותשובתך |  |
| 2:2.61 | מהרה |  |
| 2:2.62 | אדוני |  |
| 2:2.63 | ע"כ | . |


## Milchamot ha-Shem 2:3

Hebrew (Wikisource): תשובה. תשובתך מבוארת בספר כסא אליהו מדף י"ג עד דף י"ח ישמע חכם ויוסף לקח. והמקובל בידינו שהעבודה לאל אדון אחד להמשיך השפע ממנו אל עשר ספירות כמו שכתב לחם שלמה. עכ"ל המשיב.

| id | word | punctuation after |
|---|---|---|
| 2:3.1 | תשובה | . |
| 2:3.2 | תשובתך |  |
| 2:3.3 | מבוארת |  |
| 2:3.4 | בספר |  |
| 2:3.5 | כסא |  |
| 2:3.6 | אליהו |  |
| 2:3.7 | מדף |  |
| 2:3.8 | י"ג |  |
| 2:3.9 | עד |  |
| 2:3.10 | דף |  |
| 2:3.11 | י"ח |  |
| 2:3.12 | ישמע |  |
| 2:3.13 | חכם |  |
| 2:3.14 | ויוסף |  |
| 2:3.15 | לקח | . |
| 2:3.16 | והמקובל |  |
| 2:3.17 | בידינו |  |
| 2:3.18 | שהעבודה |  |
| 2:3.19 | לאל |  |
| 2:3.20 | אדון |  |
| 2:3.21 | אחד |  |
| 2:3.22 | להמשיך |  |
| 2:3.23 | השפע |  |
| 2:3.24 | ממנו |  |
| 2:3.25 | אל |  |
| 2:3.26 | עשר |  |
| 2:3.27 | ספירות |  |
| 2:3.28 | כמו |  |
| 2:3.29 | שכתב |  |
| 2:3.30 | לחם |  |
| 2:3.31 | שלמה | . |
| 2:3.32 | עכ"ל |  |
| 2:3.33 | המשיב | . |


## Milchamot ha-Shem 2:4

Hebrew (Wikisource): חזר השואל והשיב. תשובתך ידידי באה סתומה וחתומה ואין בה אלא מורה מקום, "לעיין בכסא אליהו מדף י"ג עד דף י"ח ישמע חכם ויוסף לקח והמקובל בידינו שהעבודה לאל אדון אחד" וכו'. ולא הבנתי דעתך וכוונתך בזה, אם זה המקובל בידך הוא כדברי כסא אליהו או היפך דבריו. כי כאשר עיינתי בכסא אליהו מקום שרשמת לי, לא מצאתי מבוקשי, אדרבא שם בדף ג' מדמה אחדות הבורא לשאר מיני האחדים, ששם כתב: כמו שזה הבית בכללותו נקרא בשם בית אחד. ואם תכנס בו תמצא כמה וכמה חדרים גדולים וקטנים וכמה מקומות זולתם. וכל א' ואחד מתארים לו תואר שם אחד. גם הבנין של עצים שיש בתוכו כמה תארים יש בו. גם תמצא עוד שזה הכותל קדש שנבנה היה חלקים כל אבן ואבן היא חלוקה מצרורות העפר לעצמו והסיד לעצמו כו'. ואחר שבא הבנאי וחבר כל חלקים אלו בחכמתו זה בתוך זה וזה אצל זה ונעשו כותל אחד וכו', נמצא אז הכל חתיכה אחת וחלק אחד. וכן הענין אצלנו וכו'. ושם בדף כ"ה וכ"ט מדמה אותו לגוף האדם המחובר מעצמות וגידים ובשר וראש ועין ואזן וחוטם ופה וידים ורגלים והכל נקרא בשם אדם ראובן או שמעון. ובדף כ"ה וכ"ו וכ"ז כתב: הכלל העולה כי הסבה הראשונה הנקרא בפי כל המקובלים אין סוף, הוא אשר האציל וברא ויצר ועשה והוא מסתתר תוך זעיר אנפין ואזי זעיר אנפין הוא השולט על כל הברואים ומנהיגם וזנם ומפרנסם בכח אין סוף שבו. ולכן הוא אלהינו ואנחנו עמו כי נשמותינו הם חלקו ואותו נעבוד והוא אלהי אבותינו שבידו הנהגת העולמות בשכר ועונש אבל האין סוף עם שאר הפרצופים שלמעלה מזעיר ונוקביה אם יסדרו בני אדם תפלתם להם בפרטות מבלי שיתפללו לזעיר אנפין גם אם יכוונו לנשמה המסתתרת בהם יהיו צועקים ואינם נענים. ואדרבה יהיו המתפללים אליהם נענשים כי זהו רצון הסבה הראשונה שיהיה הוא המוציא והמביא שפעו לתחתונים ואין עוד מלבדו ע"ש היטב. ודברי הר' כסא אליהו עומדים נגד דברי הרמב"ם ז"ל בפי' המשנה ובחבור ובמורה נבוכים ונגד דברי החסיד בעל חובת הלבבות בשער היחוד והרס"ג בס' האמונות והדעות והרוקח שכתבו שאחדותו ית' איננו כאחד הזוג ולא כאחד המין ולא כאיש האחד שנחלק לאחדים רבים. ולא אחד כמו הגוף הפשוט וכו' שמקבל החלוק לאין סוף. אבל השי"ת הוא אחד באחדות שאין כמותה. ג"כ כתב כסא אליהו שהאין סוף הוא נשמת עתיק, וארוך אנפין, ואבא ואימא, וזעיר ונוקביה, ורבותי ז"ל הנז"ל כתבו שאינו גוף ולא כח בגוף! ולפי דבריו הוא כח בגוף. ג"כ כתב שאריך אנפין אבא ואימא קדמו לזעיר אנפין שהוא אלהינו (לפי דבריו) והוא נקרא בן לאבא ואימא וכו'. ורז"ל אמרו שהוא קדמון לכל הנמצאים ושכל נמצא זולתו בלתי קדמון וכו'. ג"כ כתב שאין סוף ב"ה לא שייכא שום עבודה בו ולא שום קריאה כלל. ולא שייך בו שום שם לא הוי"ה ולא אדנות וכו' רק לז"א ונוקביה שבעולם האצילות בלבד ולא לא"א ואבא ואימא וזעיר ונוקביה שיש בעולמות הנאצלים למעלה מעולם דאצילות כמבואר שם ד' מ"ט סוף עמוד א' ובדף נ"ג....

| id | word | punctuation after |
|---|---|---|
| 2:4.1 | חזר |  |
| 2:4.2 | השואל |  |
| 2:4.3 | והשיב | . |
| 2:4.4 | תשובתך |  |
| 2:4.5 | ידידי |  |
| 2:4.6 | באה |  |
| 2:4.7 | סתומה |  |
| 2:4.8 | וחתומה |  |
| 2:4.9 | ואין |  |
| 2:4.10 | בה |  |
| 2:4.11 | אלא |  |
| 2:4.12 | מורה |  |
| 2:4.13 | מקום | , |
| 2:4.14 | לעיין |  |
| 2:4.15 | בכסא |  |
| 2:4.16 | אליהו |  |
| 2:4.17 | מדף |  |
| 2:4.18 | י"ג |  |
| 2:4.19 | עד |  |
| 2:4.20 | דף |  |
| 2:4.21 | י"ח |  |
| 2:4.22 | ישמע |  |
| 2:4.23 | חכם |  |
| 2:4.24 | ויוסף |  |
| 2:4.25 | לקח |  |
| 2:4.26 | והמקובל |  |
| 2:4.27 | בידינו |  |
| 2:4.28 | שהעבודה |  |
| 2:4.29 | לאל |  |
| 2:4.30 | אדון |  |
| 2:4.31 | אחד | " |
| 2:4.32 | וכו | '. |
| 2:4.33 | ולא |  |
| 2:4.34 | הבנתי |  |
| 2:4.35 | דעתך |  |
| 2:4.36 | וכוונתך |  |
| 2:4.37 | בזה | , |
| 2:4.38 | אם |  |
| 2:4.39 | זה |  |
| 2:4.40 | המקובל |  |
| 2:4.41 | בידך |  |
| 2:4.42 | הוא |  |
| 2:4.43 | כדברי |  |
| 2:4.44 | כסא |  |
| 2:4.45 | אליהו |  |
| 2:4.46 | או |  |
| 2:4.47 | היפך |  |
| 2:4.48 | דבריו | . |
| 2:4.49 | כי |  |
| 2:4.50 | כאשר |  |
| 2:4.51 | עיינתי |  |
| 2:4.52 | בכסא |  |
| 2:4.53 | אליהו |  |
| 2:4.54 | מקום |  |
| 2:4.55 | שרשמת |  |
| 2:4.56 | לי | , |
| 2:4.57 | לא |  |
| 2:4.58 | מצאתי |  |
| 2:4.59 | מבוקשי | , |
| 2:4.60 | אדרבא |  |
| 2:4.61 | שם |  |
| 2:4.62 | בדף |  |
| 2:4.63 | ג | ' |
| 2:4.64 | מדמה |  |
| 2:4.65 | אחדות |  |
| 2:4.66 | הבורא |  |
| 2:4.67 | לשאר |  |
| 2:4.68 | מיני |  |
| 2:4.69 | האחדים | , |
| 2:4.70 | ששם |  |
| 2:4.71 | כתב | : |
| 2:4.72 | כמו |  |
| 2:4.73 | שזה |  |
| 2:4.74 | הבית |  |
| 2:4.75 | בכללותו |  |
| 2:4.76 | נקרא |  |
| 2:4.77 | בשם |  |
| 2:4.78 | בית |  |
| 2:4.79 | אחד | . |
| 2:4.80 | ואם |  |
| 2:4.81 | תכנס |  |
| 2:4.82 | בו |  |
| 2:4.83 | תמצא |  |
| 2:4.84 | כמה |  |
| 2:4.85 | וכמה |  |
| 2:4.86 | חדרים |  |
| 2:4.87 | גדולים |  |
| 2:4.88 | וקטנים |  |
| 2:4.89 | וכמה |  |
| 2:4.90 | מקומות |  |
| 2:4.91 | זולתם | . |
| 2:4.92 | וכל |  |
| 2:4.93 | א | ' |
| 2:4.94 | ואחד |  |
| 2:4.95 | מתארים |  |
| 2:4.96 | לו |  |
| 2:4.97 | תואר |  |
| 2:4.98 | שם |  |
| 2:4.99 | אחד | . |
| 2:4.100 | גם |  |
| 2:4.101 | הבנין |  |
| 2:4.102 | של |  |
| 2:4.103 | עצים |  |
| 2:4.104 | שיש |  |
| 2:4.105 | בתוכו |  |
| 2:4.106 | כמה |  |
| 2:4.107 | תארים |  |
| 2:4.108 | יש |  |
| 2:4.109 | בו | . |
| 2:4.110 | גם |  |
| 2:4.111 | תמצא |  |
| 2:4.112 | עוד |  |
| 2:4.113 | שזה |  |
| 2:4.114 | הכותל |  |
| 2:4.115 | קדש |  |
| 2:4.116 | שנבנה |  |
| 2:4.117 | היה |  |
| 2:4.118 | חלקים |  |
| 2:4.119 | כל |  |
| 2:4.120 | אבן |  |
| 2:4.121 | ואבן |  |
| 2:4.122 | היא |  |
| 2:4.123 | חלוקה |  |
| 2:4.124 | מצרורות |  |
| 2:4.125 | העפר |  |
| 2:4.126 | לעצמו |  |
| 2:4.127 | והסיד |  |
| 2:4.128 | לעצמו |  |
| 2:4.129 | כו | '. |
| 2:4.130 | ואחר |  |
| 2:4.131 | שבא |  |
| 2:4.132 | הבנאי |  |
| 2:4.133 | וחבר |  |
| 2:4.134 | כל |  |
| 2:4.135 | חלקים |  |
| 2:4.136 | אלו |  |
| 2:4.137 | בחכמתו |  |
| 2:4.138 | זה |  |
| 2:4.139 | בתוך |  |
| 2:4.140 | זה |  |
| 2:4.141 | וזה |  |
| 2:4.142 | אצל |  |
| 2:4.143 | זה |  |
| 2:4.144 | ונעשו |  |
| 2:4.145 | כותל |  |
| 2:4.146 | אחד |  |
| 2:4.147 | וכו | ', |
| 2:4.148 | נמצא |  |
| 2:4.149 | אז |  |
| 2:4.150 | הכל |  |
| 2:4.151 | חתיכה |  |
| 2:4.152 | אחת |  |
| 2:4.153 | וחלק |  |
| 2:4.154 | אחד | . |
| 2:4.155 | וכן |  |
| 2:4.156 | הענין |  |
| 2:4.157 | אצלנו |  |
| 2:4.158 | וכו | '. |
| 2:4.159 | ושם |  |
| 2:4.160 | בדף |  |
| 2:4.161 | כ"ה |  |
| 2:4.162 | וכ"ט |  |
| 2:4.163 | מדמה |  |
| 2:4.164 | אותו |  |
| 2:4.165 | לגוף |  |
| 2:4.166 | האדם |  |
| 2:4.167 | המחובר |  |
| 2:4.168 | מעצמות |  |
| 2:4.169 | וגידים |  |
| 2:4.170 | ובשר |  |
| 2:4.171 | וראש |  |
| 2:4.172 | ועין |  |
| 2:4.173 | ואזן |  |
| 2:4.174 | וחוטם |  |
| 2:4.175 | ופה |  |
| 2:4.176 | וידים |  |
| 2:4.177 | ורגלים |  |
| 2:4.178 | והכל |  |
| 2:4.179 | נקרא |  |
| 2:4.180 | בשם |  |
| 2:4.181 | אדם |  |
| 2:4.182 | ראובן |  |
| 2:4.183 | או |  |
| 2:4.184 | שמעון | . |
| 2:4.185 | ובדף |  |
| 2:4.186 | כ"ה |  |
| 2:4.187 | וכ"ו |  |
| 2:4.188 | וכ"ז |  |
| 2:4.189 | כתב | : |
| 2:4.190 | הכלל |  |
| 2:4.191 | העולה |  |
| 2:4.192 | כי |  |
| 2:4.193 | הסבה |  |
| 2:4.194 | הראשונה |  |
| 2:4.195 | הנקרא |  |
| 2:4.196 | בפי |  |
| 2:4.197 | כל |  |
| 2:4.198 | המקובלים |  |
| 2:4.199 | אין |  |
| 2:4.200 | סוף | , |
| 2:4.201 | הוא |  |
| 2:4.202 | אשר |  |
| 2:4.203 | האציל |  |
| 2:4.204 | וברא |  |
| 2:4.205 | ויצר |  |
| 2:4.206 | ועשה |  |
| 2:4.207 | והוא |  |
| 2:4.208 | מסתתר |  |
| 2:4.209 | תוך |  |
| 2:4.210 | זעיר |  |
| 2:4.211 | אנפין |  |
| 2:4.212 | ואזי |  |
| 2:4.213 | זעיר |  |
| 2:4.214 | אנפין |  |
| 2:4.215 | הוא |  |
| 2:4.216 | השולט |  |
| 2:4.217 | על |  |
| 2:4.218 | כל |  |
| 2:4.219 | הברואים |  |
| 2:4.220 | ומנהיגם |  |
| 2:4.221 | וזנם |  |
| 2:4.222 | ומפרנסם |  |
| 2:4.223 | בכח |  |
| 2:4.224 | אין |  |
| 2:4.225 | סוף |  |
| 2:4.226 | שבו | . |
| 2:4.227 | ולכן |  |
| 2:4.228 | הוא |  |
| 2:4.229 | אלהינו |  |
| 2:4.230 | ואנחנו |  |
| 2:4.231 | עמו |  |
| 2:4.232 | כי |  |
| 2:4.233 | נשמותינו |  |
| 2:4.234 | הם |  |
| 2:4.235 | חלקו |  |
| 2:4.236 | ואותו |  |
| 2:4.237 | נעבוד |  |
| 2:4.238 | והוא |  |
| 2:4.239 | אלהי |  |
| 2:4.240 | אבותינו |  |
| 2:4.241 | שבידו |  |
| 2:4.242 | הנהגת |  |
| 2:4.243 | העולמות |  |
| 2:4.244 | בשכר |  |
| 2:4.245 | ועונש |  |
| 2:4.246 | אבל |  |
| 2:4.247 | האין |  |
| 2:4.248 | סוף |  |
| 2:4.249 | עם |  |
| 2:4.250 | שאר |  |
| 2:4.251 | הפרצופים |  |
| 2:4.252 | שלמעלה |  |
| 2:4.253 | מזעיר |  |
| 2:4.254 | ונוקביה |  |
| 2:4.255 | אם |  |
| 2:4.256 | יסדרו |  |
| 2:4.257 | בני |  |
| 2:4.258 | אדם |  |
| 2:4.259 | תפלתם |  |
| 2:4.260 | להם |  |
| 2:4.261 | בפרטות |  |
| 2:4.262 | מבלי |  |
| 2:4.263 | שיתפללו |  |
| 2:4.264 | לזעיר |  |
| 2:4.265 | אנפין |  |
| 2:4.266 | גם |  |
| 2:4.267 | אם |  |
| 2:4.268 | יכוונו |  |
| 2:4.269 | לנשמה |  |
| 2:4.270 | המסתתרת |  |
| 2:4.271 | בהם |  |
| 2:4.272 | יהיו |  |
| 2:4.273 | צועקים |  |
| 2:4.274 | ואינם |  |
| 2:4.275 | נענים | . |
| 2:4.276 | ואדרבה |  |
| 2:4.277 | יהיו |  |
| 2:4.278 | המתפללים |  |
| 2:4.279 | אליהם |  |
| 2:4.280 | נענשים |  |
| 2:4.281 | כי |  |
| 2:4.282 | זהו |  |
| 2:4.283 | רצון |  |
| 2:4.284 | הסבה |  |
| 2:4.285 | הראשונה |  |
| 2:4.286 | שיהיה |  |
| 2:4.287 | הוא |  |
| 2:4.288 | המוציא |  |
| 2:4.289 | והמביא |  |
| 2:4.290 | שפעו |  |
| 2:4.291 | לתחתונים |  |
| 2:4.292 | ואין |  |
| 2:4.293 | עוד |  |
| 2:4.294 | מלבדו |  |
| 2:4.295 | ע"ש |  |
| 2:4.296 | היטב | . |
| 2:4.297 | ודברי |  |
| 2:4.298 | הר | ' |
| 2:4.299 | כסא |  |
| 2:4.300 | אליהו |  |
| 2:4.301 | עומדים |  |
| 2:4.302 | נגד |  |
| 2:4.303 | דברי |  |
| 2:4.304 | הרמב"ם |  |
| 2:4.305 | ז"ל |  |
| 2:4.306 | בפי | ' |
| 2:4.307 | המשנה |  |
| 2:4.308 | ובחבור |  |
| 2:4.309 | ובמורה |  |
| 2:4.310 | נבוכים |  |
| 2:4.311 | ונגד |  |
| 2:4.312 | דברי |  |
| 2:4.313 | החסיד |  |
| 2:4.314 | בעל |  |
| 2:4.315 | חובת |  |
| 2:4.316 | הלבבות |  |
| 2:4.317 | בשער |  |
| 2:4.318 | היחוד |  |
| 2:4.319 | והרס"ג |  |
| 2:4.320 | בס | ' |
| 2:4.321 | האמונות |  |
| 2:4.322 | והדעות |  |
| 2:4.323 | והרוקח |  |
| 2:4.324 | שכתבו |  |
| 2:4.325 | שאחדותו |  |
| 2:4.326 | ית | ' |
| 2:4.327 | איננו |  |
| 2:4.328 | כאחד |  |
| 2:4.329 | הזוג |  |
| 2:4.330 | ולא |  |
| 2:4.331 | כאחד |  |
| 2:4.332 | המין |  |
| 2:4.333 | ולא |  |
| 2:4.334 | כאיש |  |
| 2:4.335 | האחד |  |
| 2:4.336 | שנחלק |  |
| 2:4.337 | לאחדים |  |
| 2:4.338 | רבים | . |
| 2:4.339 | ולא |  |
| 2:4.340 | אחד |  |
| 2:4.341 | כמו |  |
| 2:4.342 | הגוף |  |
| 2:4.343 | הפשוט |  |
| 2:4.344 | וכו | ' |
| 2:4.345 | שמקבל |  |
| 2:4.346 | החלוק |  |
| 2:4.347 | לאין |  |
| 2:4.348 | סוף | . |
| 2:4.349 | אבל |  |
| 2:4.350 | השי"ת |  |
| 2:4.351 | הוא |  |
| 2:4.352 | אחד |  |
| 2:4.353 | באחדות |  |
| 2:4.354 | שאין |  |
| 2:4.355 | כמותה | . |
| 2:4.356 | ג"כ |  |
| 2:4.357 | כתב |  |
| 2:4.358 | כסא |  |
| 2:4.359 | אליהו |  |
| 2:4.360 | שהאין |  |
| 2:4.361 | סוף |  |
| 2:4.362 | הוא |  |
| 2:4.363 | נשמת |  |
| 2:4.364 | עתיק | , |
| 2:4.365 | וארוך |  |
| 2:4.366 | אנפין | , |
| 2:4.367 | ואבא |  |
| 2:4.368 | ואימא | , |
| 2:4.369 | וזעיר |  |
| 2:4.370 | ונוקביה | , |
| 2:4.371 | ורבותי |  |
| 2:4.372 | ז"ל |  |
| 2:4.373 | הנז"ל |  |
| 2:4.374 | כתבו |  |
| 2:4.375 | שאינו |  |
| 2:4.376 | גוף |  |
| 2:4.377 | ולא |  |
| 2:4.378 | כח |  |
| 2:4.379 | בגוף | ! |
| 2:4.380 | ולפי |  |
| 2:4.381 | דבריו |  |
| 2:4.382 | הוא |  |
| 2:4.383 | כח |  |
| 2:4.384 | בגוף | . |
| 2:4.385 | ג"כ |  |
| 2:4.386 | כתב |  |
| 2:4.387 | שאריך |  |
| 2:4.388 | אנפין |  |
| 2:4.389 | אבא |  |
| 2:4.390 | ואימא |  |
| 2:4.391 | קדמו |  |
| 2:4.392 | לזעיר |  |
| 2:4.393 | אנפין |  |
| 2:4.394 | שהוא |  |
| 2:4.395 | אלהינו |  |
| 2:4.396 | לפי |  |
| 2:4.397 | דבריו | ) |
| 2:4.398 | והוא |  |
| 2:4.399 | נקרא |  |
| 2:4.400 | בן |  |
| 2:4.401 | לאבא |  |
| 2:4.402 | ואימא |  |
| 2:4.403 | וכו | '. |
| 2:4.404 | ורז"ל |  |
| 2:4.405 | אמרו |  |
| 2:4.406 | שהוא |  |
| 2:4.407 | קדמון |  |
| 2:4.408 | לכל |  |
| 2:4.409 | הנמצאים |  |
| 2:4.410 | ושכל |  |
| 2:4.411 | נמצא |  |
| 2:4.412 | זולתו |  |
| 2:4.413 | בלתי |  |
| 2:4.414 | קדמון |  |
| 2:4.415 | וכו | '. |
| 2:4.416 | ג"כ |  |
| 2:4.417 | כתב |  |
| 2:4.418 | שאין |  |
| 2:4.419 | סוף |  |
| 2:4.420 | ב"ה |  |
| 2:4.421 | לא |  |
| 2:4.422 | שייכא |  |
| 2:4.423 | שום |  |
| 2:4.424 | עבודה |  |
| 2:4.425 | בו |  |
| 2:4.426 | ולא |  |
| 2:4.427 | שום |  |
| 2:4.428 | קריאה |  |
| 2:4.429 | כלל | . |
| 2:4.430 | ולא |  |
| 2:4.431 | שייך |  |
| 2:4.432 | בו |  |
| 2:4.433 | שום |  |
| 2:4.434 | שם |  |
| 2:4.435 | לא |  |
| 2:4.436 | הוי"ה |  |
| 2:4.437 | ולא |  |
| 2:4.438 | אדנות |  |
| 2:4.439 | וכו | ' |
| 2:4.440 | רק |  |
| 2:4.441 | לז"א |  |
| 2:4.442 | ונוקביה |  |
| 2:4.443 | שבעולם |  |
| 2:4.444 | האצילות |  |
| 2:4.445 | בלבד |  |
| 2:4.446 | ולא |  |
| 2:4.447 | לא"א |  |
| 2:4.448 | ואבא |  |
| 2:4.449 | ואימא |  |
| 2:4.450 | וזעיר |  |
| 2:4.451 | ונוקביה |  |
| 2:4.452 | שיש |  |
| 2:4.453 | בעולמות |  |
| 2:4.454 | הנאצלים |  |
| 2:4.455 | למעלה |  |
| 2:4.456 | מעולם |  |
| 2:4.457 | דאצילות |  |
| 2:4.458 | כמבואר |  |
| 2:4.459 | שם |  |
| 2:4.460 | ד | ' |
| 2:4.461 | מ"ט |  |
| 2:4.462 | סוף |  |
| 2:4.463 | עמוד |  |
| 2:4.464 | א | ' |
| 2:4.465 | ובדף |  |
| 2:4.466 | נ"ג | .... |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

