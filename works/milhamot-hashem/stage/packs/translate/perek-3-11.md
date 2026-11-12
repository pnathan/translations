# Translation pack: milhamot-hashem perek-3-11 (Milchamot ha-Shem 3:17–3:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-11`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **זלה"ה** → "may his memory be for the life of the world to come".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **תימן** → "Yemen".
- **בא"י** → "in the Land of Israel".

## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
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
- עבודה **אחרת** → "other worship" (note first in a unit).
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
- **עלה / עלול** → "cause / effect".
- **מחויב** המציאות → "the Necessary Existent".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **אפיקורוס / אפיקורסים** → "apikoros / apikorsim".
- תורה **שבעל** פה → "the Oral Torah".
- ימי **הבינים** → "the Middle Ages".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
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


## Milchamot ha-Shem 3:17

Hebrew (Wikisource): יז) ומן התורה שבכתב ושבעל פה הנמסרת לנו איש מפי איש שמשה רבינו ע"ה נביא ראיות חזקות ואמתיות שה' אלהינו הנקרא בשם הוי'ה ואדנות, ובדברי רז"ל בשם הקב"ה הוא הסבה הראשונה, והוא לבדו מחויב המציאות וכמו שיתבאר להלן שלא כמו שכ' בעל ספר הברית בשם יושר לבב ובשם האר"י שנמשכו אחרי דברי הפילוסוף מחבר הזהר ואמרו שהוא זעיר אנפין! מן התורה דכתיב מעונה אלהי קדם ענינו כי ה' אלהינו הוא הסבה הראשונה לכל ברואי מעלה ומטה והוא אשר לו לבדו נעבוד, כמו שכתב הרמב"ם ז"ל בפי' המשנה ובחיבור הרב סעדיה גאון ור' יהודה הלוי וזולתם. ונאמר אני ה' ראשון ומבלעדי אין אלהים. ודוד המלך ע"ה אמר כי מי אל מבלעדי ה' ומי צור מבלעדי אלהינו, כי הוא יתברך הסבה וההתחלה הפועלת לכל אשר זולתו כמו שביאר הרב המורה בפרק ט"ז מראשון עי"ש: אין סבה אחרת קודמת לו. וז"ל הר' החסיד בעל חובת הלבבות בשער היחוד, מן הדין שיהיה לעולם תחלה. אין תחלה לפניו, וראשון שאין לו ראשון, והוא אשר יצרו וחדשו מאין דבר, לא בדבר ולא על דבר כמו שנאמר אני ה' עשה כל נוטה שמים לבדי רוקע הארץ מי אתי. ואמר איוב נטה צפן על תהו תולה ארץ על בלי מה וכו'. הוא הקדמון הראשון אשר אין ראשית לראשיתו. ואין תכלית לקדמותו כמו שאמר הכתוב אני ראשון ואני אחרון ואומר אני ה' ראשון ואת אחרונים אני הוא עכ"ל. גם בתיקוני התפילות שתקנו אנשי כנה"ג נאמר אמת אתה הוא ראשון ואתה הוא אחרון וכו'.

| id | word | punctuation after |
|---|---|---|
| 3:17.1 | יז | ) |
| 3:17.2 | ומן |  |
| 3:17.3 | התורה |  |
| 3:17.4 | שבכתב |  |
| 3:17.5 | ושבעל |  |
| 3:17.6 | פה |  |
| 3:17.7 | הנמסרת |  |
| 3:17.8 | לנו |  |
| 3:17.9 | איש |  |
| 3:17.10 | מפי |  |
| 3:17.11 | איש |  |
| 3:17.12 | שמשה |  |
| 3:17.13 | רבינו |  |
| 3:17.14 | ע"ה |  |
| 3:17.15 | נביא |  |
| 3:17.16 | ראיות |  |
| 3:17.17 | חזקות |  |
| 3:17.18 | ואמתיות |  |
| 3:17.19 | שה | ' |
| 3:17.20 | אלהינו |  |
| 3:17.21 | הנקרא |  |
| 3:17.22 | בשם |  |
| 3:17.23 | הוי'ה |  |
| 3:17.24 | ואדנות | , |
| 3:17.25 | ובדברי |  |
| 3:17.26 | רז"ל |  |
| 3:17.27 | בשם |  |
| 3:17.28 | הקב"ה |  |
| 3:17.29 | הוא |  |
| 3:17.30 | הסבה |  |
| 3:17.31 | הראשונה | , |
| 3:17.32 | והוא |  |
| 3:17.33 | לבדו |  |
| 3:17.34 | מחויב |  |
| 3:17.35 | המציאות |  |
| 3:17.36 | וכמו |  |
| 3:17.37 | שיתבאר |  |
| 3:17.38 | להלן |  |
| 3:17.39 | שלא |  |
| 3:17.40 | כמו |  |
| 3:17.41 | שכ | ' |
| 3:17.42 | בעל |  |
| 3:17.43 | ספר |  |
| 3:17.44 | הברית |  |
| 3:17.45 | בשם |  |
| 3:17.46 | יושר |  |
| 3:17.47 | לבב |  |
| 3:17.48 | ובשם |  |
| 3:17.49 | האר"י |  |
| 3:17.50 | שנמשכו |  |
| 3:17.51 | אחרי |  |
| 3:17.52 | דברי |  |
| 3:17.53 | הפילוסוף |  |
| 3:17.54 | מחבר |  |
| 3:17.55 | הזהר |  |
| 3:17.56 | ואמרו |  |
| 3:17.57 | שהוא |  |
| 3:17.58 | זעיר |  |
| 3:17.59 | אנפין | ! |
| 3:17.60 | מן |  |
| 3:17.61 | התורה |  |
| 3:17.62 | דכתיב |  |
| 3:17.63 | מעונה |  |
| 3:17.64 | אלהי |  |
| 3:17.65 | קדם |  |
| 3:17.66 | ענינו |  |
| 3:17.67 | כי |  |
| 3:17.68 | ה | ' |
| 3:17.69 | אלהינו |  |
| 3:17.70 | הוא |  |
| 3:17.71 | הסבה |  |
| 3:17.72 | הראשונה |  |
| 3:17.73 | לכל |  |
| 3:17.74 | ברואי |  |
| 3:17.75 | מעלה |  |
| 3:17.76 | ומטה |  |
| 3:17.77 | והוא |  |
| 3:17.78 | אשר |  |
| 3:17.79 | לו |  |
| 3:17.80 | לבדו |  |
| 3:17.81 | נעבוד | , |
| 3:17.82 | כמו |  |
| 3:17.83 | שכתב |  |
| 3:17.84 | הרמב"ם |  |
| 3:17.85 | ז"ל |  |
| 3:17.86 | בפי | ' |
| 3:17.87 | המשנה |  |
| 3:17.88 | ובחיבור |  |
| 3:17.89 | הרב |  |
| 3:17.90 | סעדיה |  |
| 3:17.91 | גאון |  |
| 3:17.92 | ור | ' |
| 3:17.93 | יהודה |  |
| 3:17.94 | הלוי |  |
| 3:17.95 | וזולתם | . |
| 3:17.96 | ונאמר |  |
| 3:17.97 | אני |  |
| 3:17.98 | ה | ' |
| 3:17.99 | ראשון |  |
| 3:17.100 | ומבלעדי |  |
| 3:17.101 | אין |  |
| 3:17.102 | אלהים | . |
| 3:17.103 | ודוד |  |
| 3:17.104 | המלך |  |
| 3:17.105 | ע"ה |  |
| 3:17.106 | אמר |  |
| 3:17.107 | כי |  |
| 3:17.108 | מי |  |
| 3:17.109 | אל |  |
| 3:17.110 | מבלעדי |  |
| 3:17.111 | ה | ' |
| 3:17.112 | ומי |  |
| 3:17.113 | צור |  |
| 3:17.114 | מבלעדי |  |
| 3:17.115 | אלהינו | , |
| 3:17.116 | כי |  |
| 3:17.117 | הוא |  |
| 3:17.118 | יתברך |  |
| 3:17.119 | הסבה |  |
| 3:17.120 | וההתחלה |  |
| 3:17.121 | הפועלת |  |
| 3:17.122 | לכל |  |
| 3:17.123 | אשר |  |
| 3:17.124 | זולתו |  |
| 3:17.125 | כמו |  |
| 3:17.126 | שביאר |  |
| 3:17.127 | הרב |  |
| 3:17.128 | המורה |  |
| 3:17.129 | בפרק |  |
| 3:17.130 | ט"ז |  |
| 3:17.131 | מראשון |  |
| 3:17.132 | עי"ש | : |
| 3:17.133 | אין |  |
| 3:17.134 | סבה |  |
| 3:17.135 | אחרת |  |
| 3:17.136 | קודמת |  |
| 3:17.137 | לו | . |
| 3:17.138 | וז"ל |  |
| 3:17.139 | הר | ' |
| 3:17.140 | החסיד |  |
| 3:17.141 | בעל |  |
| 3:17.142 | חובת |  |
| 3:17.143 | הלבבות |  |
| 3:17.144 | בשער |  |
| 3:17.145 | היחוד | , |
| 3:17.146 | מן |  |
| 3:17.147 | הדין |  |
| 3:17.148 | שיהיה |  |
| 3:17.149 | לעולם |  |
| 3:17.150 | תחלה | . |
| 3:17.151 | אין |  |
| 3:17.152 | תחלה |  |
| 3:17.153 | לפניו | , |
| 3:17.154 | וראשון |  |
| 3:17.155 | שאין |  |
| 3:17.156 | לו |  |
| 3:17.157 | ראשון | , |
| 3:17.158 | והוא |  |
| 3:17.159 | אשר |  |
| 3:17.160 | יצרו |  |
| 3:17.161 | וחדשו |  |
| 3:17.162 | מאין |  |
| 3:17.163 | דבר | , |
| 3:17.164 | לא |  |
| 3:17.165 | בדבר |  |
| 3:17.166 | ולא |  |
| 3:17.167 | על |  |
| 3:17.168 | דבר |  |
| 3:17.169 | כמו |  |
| 3:17.170 | שנאמר |  |
| 3:17.171 | אני |  |
| 3:17.172 | ה | ' |
| 3:17.173 | עשה |  |
| 3:17.174 | כל |  |
| 3:17.175 | נוטה |  |
| 3:17.176 | שמים |  |
| 3:17.177 | לבדי |  |
| 3:17.178 | רוקע |  |
| 3:17.179 | הארץ |  |
| 3:17.180 | מי |  |
| 3:17.181 | אתי | . |
| 3:17.182 | ואמר |  |
| 3:17.183 | איוב |  |
| 3:17.184 | נטה |  |
| 3:17.185 | צפן |  |
| 3:17.186 | על |  |
| 3:17.187 | תהו |  |
| 3:17.188 | תולה |  |
| 3:17.189 | ארץ |  |
| 3:17.190 | על |  |
| 3:17.191 | בלי |  |
| 3:17.192 | מה |  |
| 3:17.193 | וכו | '. |
| 3:17.194 | הוא |  |
| 3:17.195 | הקדמון |  |
| 3:17.196 | הראשון |  |
| 3:17.197 | אשר |  |
| 3:17.198 | אין |  |
| 3:17.199 | ראשית |  |
| 3:17.200 | לראשיתו | . |
| 3:17.201 | ואין |  |
| 3:17.202 | תכלית |  |
| 3:17.203 | לקדמותו |  |
| 3:17.204 | כמו |  |
| 3:17.205 | שאמר |  |
| 3:17.206 | הכתוב |  |
| 3:17.207 | אני |  |
| 3:17.208 | ראשון |  |
| 3:17.209 | ואני |  |
| 3:17.210 | אחרון |  |
| 3:17.211 | ואומר |  |
| 3:17.212 | אני |  |
| 3:17.213 | ה | ' |
| 3:17.214 | ראשון |  |
| 3:17.215 | ואת |  |
| 3:17.216 | אחרונים |  |
| 3:17.217 | אני |  |
| 3:17.218 | הוא |  |
| 3:17.219 | עכ"ל | . |
| 3:17.220 | גם |  |
| 3:17.221 | בתיקוני |  |
| 3:17.222 | התפילות |  |
| 3:17.223 | שתקנו |  |
| 3:17.224 | אנשי |  |
| 3:17.225 | כנה"ג |  |
| 3:17.226 | נאמר |  |
| 3:17.227 | אמת |  |
| 3:17.228 | אתה |  |
| 3:17.229 | הוא |  |
| 3:17.230 | ראשון |  |
| 3:17.231 | ואתה |  |
| 3:17.232 | הוא |  |
| 3:17.233 | אחרון |  |
| 3:17.234 | וכו | '. |


## Milchamot ha-Shem 3:18

Hebrew (Wikisource): יח) ועוד ראיה שאינה קבלה ממשה שהרי מצינו מחלוקת בין המקובלים באיכותה ועצמותה זה אומר בכה וזה אומר בכה, וזו ראיה גמורה שאינה הלכה למשה מסיני כמו שעלה על דעתך, שהרי כתב הרמב"ם ז"ל בהל' ממרים פ"א שדברי קבלה אין בהם מחלוקת כלל. וכל דבר שתמצא בו מחלוקת בידוע שאינו קבלה ממשה רבי' ע"כ, וכיון שכל ראיות שאנו מביאים הם מדברי רז"ל בעלי המשנה והתלמוד והמדרשים האמתיים והגאונים והפוסקים שהם בודאי בעלי הקבלה האמתית ממשה רבי' ע"ה, וכל האומר היפך דבריהם הוא מין ואפיקורוס, ואם כן איך תוכל לומר מי נתן לך לדרוש ולחקור אחר דבריהם של בעלי הקבלה החדשה? וכבר העיד אחד מיוחד מחכמי ימי הבינים בארץ התימן הוא מהרי"י צאהרי בדרשות בפרשת לך לך, וז"ל ואני באתי להגיד לך מעט מזער דברי הקבלה חדשים מקרוב באו ע"כ. ועדותו זאת מועלת לנו לחזק דברינו, כי חדשות המה לא שערום רבותינו קדמונינו. ופליטת ספריהם יעידון יגידון האמת הלא המה מצויים בידינו במדרש הגדול לר' דוד בן עמרם איש עדן כמו שכ' מהר"י בשירי בפרשת בראשית ובתחלת אלה הדברים. וכ"כ בספר סגולת ישראל, וספר נור אל צאלם לר' נתנאל בן ישע ומדרש החפץ לר' יחיא הרופא זלה"ה.

| id | word | punctuation after |
|---|---|---|
| 3:18.1 | יח | ) |
| 3:18.2 | ועוד |  |
| 3:18.3 | ראיה |  |
| 3:18.4 | שאינה |  |
| 3:18.5 | קבלה |  |
| 3:18.6 | ממשה |  |
| 3:18.7 | שהרי |  |
| 3:18.8 | מצינו |  |
| 3:18.9 | מחלוקת |  |
| 3:18.10 | בין |  |
| 3:18.11 | המקובלים |  |
| 3:18.12 | באיכותה |  |
| 3:18.13 | ועצמותה |  |
| 3:18.14 | זה |  |
| 3:18.15 | אומר |  |
| 3:18.16 | בכה |  |
| 3:18.17 | וזה |  |
| 3:18.18 | אומר |  |
| 3:18.19 | בכה | , |
| 3:18.20 | וזו |  |
| 3:18.21 | ראיה |  |
| 3:18.22 | גמורה |  |
| 3:18.23 | שאינה |  |
| 3:18.24 | הלכה |  |
| 3:18.25 | למשה |  |
| 3:18.26 | מסיני |  |
| 3:18.27 | כמו |  |
| 3:18.28 | שעלה |  |
| 3:18.29 | על |  |
| 3:18.30 | דעתך | , |
| 3:18.31 | שהרי |  |
| 3:18.32 | כתב |  |
| 3:18.33 | הרמב"ם |  |
| 3:18.34 | ז"ל |  |
| 3:18.35 | בהל | ' |
| 3:18.36 | ממרים |  |
| 3:18.37 | פ"א |  |
| 3:18.38 | שדברי |  |
| 3:18.39 | קבלה |  |
| 3:18.40 | אין |  |
| 3:18.41 | בהם |  |
| 3:18.42 | מחלוקת |  |
| 3:18.43 | כלל | . |
| 3:18.44 | וכל |  |
| 3:18.45 | דבר |  |
| 3:18.46 | שתמצא |  |
| 3:18.47 | בו |  |
| 3:18.48 | מחלוקת |  |
| 3:18.49 | בידוע |  |
| 3:18.50 | שאינו |  |
| 3:18.51 | קבלה |  |
| 3:18.52 | ממשה |  |
| 3:18.53 | רבי | ' |
| 3:18.54 | ע"כ | , |
| 3:18.55 | וכיון |  |
| 3:18.56 | שכל |  |
| 3:18.57 | ראיות |  |
| 3:18.58 | שאנו |  |
| 3:18.59 | מביאים |  |
| 3:18.60 | הם |  |
| 3:18.61 | מדברי |  |
| 3:18.62 | רז"ל |  |
| 3:18.63 | בעלי |  |
| 3:18.64 | המשנה |  |
| 3:18.65 | והתלמוד |  |
| 3:18.66 | והמדרשים |  |
| 3:18.67 | האמתיים |  |
| 3:18.68 | והגאונים |  |
| 3:18.69 | והפוסקים |  |
| 3:18.70 | שהם |  |
| 3:18.71 | בודאי |  |
| 3:18.72 | בעלי |  |
| 3:18.73 | הקבלה |  |
| 3:18.74 | האמתית |  |
| 3:18.75 | ממשה |  |
| 3:18.76 | רבי | ' |
| 3:18.77 | ע"ה | , |
| 3:18.78 | וכל |  |
| 3:18.79 | האומר |  |
| 3:18.80 | היפך |  |
| 3:18.81 | דבריהם |  |
| 3:18.82 | הוא |  |
| 3:18.83 | מין |  |
| 3:18.84 | ואפיקורוס | , |
| 3:18.85 | ואם |  |
| 3:18.86 | כן |  |
| 3:18.87 | איך |  |
| 3:18.88 | תוכל |  |
| 3:18.89 | לומר |  |
| 3:18.90 | מי |  |
| 3:18.91 | נתן |  |
| 3:18.92 | לך |  |
| 3:18.93 | לדרוש |  |
| 3:18.94 | ולחקור |  |
| 3:18.95 | אחר |  |
| 3:18.96 | דבריהם |  |
| 3:18.97 | של |  |
| 3:18.98 | בעלי |  |
| 3:18.99 | הקבלה |  |
| 3:18.100 | החדשה | ? |
| 3:18.101 | וכבר |  |
| 3:18.102 | העיד |  |
| 3:18.103 | אחד |  |
| 3:18.104 | מיוחד |  |
| 3:18.105 | מחכמי |  |
| 3:18.106 | ימי |  |
| 3:18.107 | הבינים |  |
| 3:18.108 | בארץ |  |
| 3:18.109 | התימן |  |
| 3:18.110 | הוא |  |
| 3:18.111 | מהרי"י |  |
| 3:18.112 | צאהרי |  |
| 3:18.113 | בדרשות |  |
| 3:18.114 | בפרשת |  |
| 3:18.115 | לך |  |
| 3:18.116 | לך | , |
| 3:18.117 | וז"ל |  |
| 3:18.118 | ואני |  |
| 3:18.119 | באתי |  |
| 3:18.120 | להגיד |  |
| 3:18.121 | לך |  |
| 3:18.122 | מעט |  |
| 3:18.123 | מזער |  |
| 3:18.124 | דברי |  |
| 3:18.125 | הקבלה |  |
| 3:18.126 | חדשים |  |
| 3:18.127 | מקרוב |  |
| 3:18.128 | באו |  |
| 3:18.129 | ע"כ | . |
| 3:18.130 | ועדותו |  |
| 3:18.131 | זאת |  |
| 3:18.132 | מועלת |  |
| 3:18.133 | לנו |  |
| 3:18.134 | לחזק |  |
| 3:18.135 | דברינו | , |
| 3:18.136 | כי |  |
| 3:18.137 | חדשות |  |
| 3:18.138 | המה |  |
| 3:18.139 | לא |  |
| 3:18.140 | שערום |  |
| 3:18.141 | רבותינו |  |
| 3:18.142 | קדמונינו | . |
| 3:18.143 | ופליטת |  |
| 3:18.144 | ספריהם |  |
| 3:18.145 | יעידון |  |
| 3:18.146 | יגידון |  |
| 3:18.147 | האמת |  |
| 3:18.148 | הלא |  |
| 3:18.149 | המה |  |
| 3:18.150 | מצויים |  |
| 3:18.151 | בידינו |  |
| 3:18.152 | במדרש |  |
| 3:18.153 | הגדול |  |
| 3:18.154 | לר | ' |
| 3:18.155 | דוד |  |
| 3:18.156 | בן |  |
| 3:18.157 | עמרם |  |
| 3:18.158 | איש |  |
| 3:18.159 | עדן |  |
| 3:18.160 | כמו |  |
| 3:18.161 | שכ | ' |
| 3:18.162 | מהר"י |  |
| 3:18.163 | בשירי |  |
| 3:18.164 | בפרשת |  |
| 3:18.165 | בראשית |  |
| 3:18.166 | ובתחלת |  |
| 3:18.167 | אלה |  |
| 3:18.168 | הדברים | . |
| 3:18.169 | וכ"כ |  |
| 3:18.170 | בספר |  |
| 3:18.171 | סגולת |  |
| 3:18.172 | ישראל | , |
| 3:18.173 | וספר |  |
| 3:18.174 | נור |  |
| 3:18.175 | אל |  |
| 3:18.176 | צאלם |  |
| 3:18.177 | לר | ' |
| 3:18.178 | נתנאל |  |
| 3:18.179 | בן |  |
| 3:18.180 | ישע |  |
| 3:18.181 | ומדרש |  |
| 3:18.182 | החפץ |  |
| 3:18.183 | לר | ' |
| 3:18.184 | יחיא |  |
| 3:18.185 | הרופא |  |
| 3:18.186 | זלה"ה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

