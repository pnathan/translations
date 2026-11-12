# Translation pack: milhamot-hashem perek-6-02 (Milchamot ha-Shem 6:2–6:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-02`.

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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
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
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **אחי** → "my brother" (second person kept).
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 6:2

Hebrew (Wikisource): סא) וכ"כ הסמ"ג בעשיין א' מצות עשה להאמין ולשמוע היא הקבלה שהוא ית' בשמים ובארץ ובארבע רוחות העולם שנ' שמע ישראל ה' אלהינו ה' אחד, שמע כמו תקבל וכן אתה תשמע השמים וגו' וזה האחד אינו כשאר האחדים, המלך הוא אחד בארצו, אבל אינו אחד מכל צדדים, כי יש בארצו אנשים שהם כמהו ובשאר ארצות יש מלכים כמהו, ואם מלאך מן השמים ירד לארץ. הנה המלאך ההוא אחד בארץ. אבל אינו אחד מכל צדדים כי יש בשמים מלאכים כמותו. אבל ה' אלהינו אינו כך. הוא אחד מכל הצדדים. ורב סעדיה כ' להשיב לכופרים שאומרים שהם שנים או שלשה. אם לא יוכל האחד לעשות מעשה אלא בעזר חבירו האי שניהם חלשים. ואם אחד יכול להכריח את חבירו, הרי שניהם מוכרחים. ואם שניהם רשאים לעשות על אחד מה שירצה, ובקש אחד להמית אדם אחד והשני רוצה להחיותו. ראוי שיהיה האיש ההוא מת חי מת חי לאלתר כהרף עין. ואם יוכל האחד להסתיר דבר מחבירו הרי שניהם נלאים. ועוד כתב לאלה שפוקדים על נעשה אדם בצלמנו (היא דעת מחבר הזהר והמקובלים כדלעיל סי' ל"ח) כי דרך שפת לשון הקדש שהגדולים מכנים עצמם בלשון רבים אע"פ שהם יחידים, כמו אולי אוכל נכה בו. ובדניאל נאמר קדם מלכא ובמנוח נעצרה נא אותך ע"כ ועוד האריך רבי סעדיה גאון מאד, ואין צורך להאריך כל כך, כי כל ישראל תקועים באמונה חזקה, כי יוצר הכל אחד מיוחד כמו שנ' ה' אחד. ונאמר ראו עתה כי אני הוא ואין אלהים עמדי. ובמקרא יש הרבה מקראות כאלה עכ"ל הרב הסמ"ג.

| id | word | punctuation after |
|---|---|---|
| 6:2.1 | סא | ) |
| 6:2.2 | וכ"כ |  |
| 6:2.3 | הסמ"ג |  |
| 6:2.4 | בעשיין |  |
| 6:2.5 | א | ' |
| 6:2.6 | מצות |  |
| 6:2.7 | עשה |  |
| 6:2.8 | להאמין |  |
| 6:2.9 | ולשמוע |  |
| 6:2.10 | היא |  |
| 6:2.11 | הקבלה |  |
| 6:2.12 | שהוא |  |
| 6:2.13 | ית | ' |
| 6:2.14 | בשמים |  |
| 6:2.15 | ובארץ |  |
| 6:2.16 | ובארבע |  |
| 6:2.17 | רוחות |  |
| 6:2.18 | העולם |  |
| 6:2.19 | שנ | ' |
| 6:2.20 | שמע |  |
| 6:2.21 | ישראל |  |
| 6:2.22 | ה | ' |
| 6:2.23 | אלהינו |  |
| 6:2.24 | ה | ' |
| 6:2.25 | אחד | , |
| 6:2.26 | שמע |  |
| 6:2.27 | כמו |  |
| 6:2.28 | תקבל |  |
| 6:2.29 | וכן |  |
| 6:2.30 | אתה |  |
| 6:2.31 | תשמע |  |
| 6:2.32 | השמים |  |
| 6:2.33 | וגו | ' |
| 6:2.34 | וזה |  |
| 6:2.35 | האחד |  |
| 6:2.36 | אינו |  |
| 6:2.37 | כשאר |  |
| 6:2.38 | האחדים | , |
| 6:2.39 | המלך |  |
| 6:2.40 | הוא |  |
| 6:2.41 | אחד |  |
| 6:2.42 | בארצו | , |
| 6:2.43 | אבל |  |
| 6:2.44 | אינו |  |
| 6:2.45 | אחד |  |
| 6:2.46 | מכל |  |
| 6:2.47 | צדדים | , |
| 6:2.48 | כי |  |
| 6:2.49 | יש |  |
| 6:2.50 | בארצו |  |
| 6:2.51 | אנשים |  |
| 6:2.52 | שהם |  |
| 6:2.53 | כמהו |  |
| 6:2.54 | ובשאר |  |
| 6:2.55 | ארצות |  |
| 6:2.56 | יש |  |
| 6:2.57 | מלכים |  |
| 6:2.58 | כמהו | , |
| 6:2.59 | ואם |  |
| 6:2.60 | מלאך |  |
| 6:2.61 | מן |  |
| 6:2.62 | השמים |  |
| 6:2.63 | ירד |  |
| 6:2.64 | לארץ | . |
| 6:2.65 | הנה |  |
| 6:2.66 | המלאך |  |
| 6:2.67 | ההוא |  |
| 6:2.68 | אחד |  |
| 6:2.69 | בארץ | . |
| 6:2.70 | אבל |  |
| 6:2.71 | אינו |  |
| 6:2.72 | אחד |  |
| 6:2.73 | מכל |  |
| 6:2.74 | צדדים |  |
| 6:2.75 | כי |  |
| 6:2.76 | יש |  |
| 6:2.77 | בשמים |  |
| 6:2.78 | מלאכים |  |
| 6:2.79 | כמותו | . |
| 6:2.80 | אבל |  |
| 6:2.81 | ה | ' |
| 6:2.82 | אלהינו |  |
| 6:2.83 | אינו |  |
| 6:2.84 | כך | . |
| 6:2.85 | הוא |  |
| 6:2.86 | אחד |  |
| 6:2.87 | מכל |  |
| 6:2.88 | הצדדים | . |
| 6:2.89 | ורב |  |
| 6:2.90 | סעדיה |  |
| 6:2.91 | כ | ' |
| 6:2.92 | להשיב |  |
| 6:2.93 | לכופרים |  |
| 6:2.94 | שאומרים |  |
| 6:2.95 | שהם |  |
| 6:2.96 | שנים |  |
| 6:2.97 | או |  |
| 6:2.98 | שלשה | . |
| 6:2.99 | אם |  |
| 6:2.100 | לא |  |
| 6:2.101 | יוכל |  |
| 6:2.102 | האחד |  |
| 6:2.103 | לעשות |  |
| 6:2.104 | מעשה |  |
| 6:2.105 | אלא |  |
| 6:2.106 | בעזר |  |
| 6:2.107 | חבירו |  |
| 6:2.108 | האי |  |
| 6:2.109 | שניהם |  |
| 6:2.110 | חלשים | . |
| 6:2.111 | ואם |  |
| 6:2.112 | אחד |  |
| 6:2.113 | יכול |  |
| 6:2.114 | להכריח |  |
| 6:2.115 | את |  |
| 6:2.116 | חבירו | , |
| 6:2.117 | הרי |  |
| 6:2.118 | שניהם |  |
| 6:2.119 | מוכרחים | . |
| 6:2.120 | ואם |  |
| 6:2.121 | שניהם |  |
| 6:2.122 | רשאים |  |
| 6:2.123 | לעשות |  |
| 6:2.124 | על |  |
| 6:2.125 | אחד |  |
| 6:2.126 | מה |  |
| 6:2.127 | שירצה | , |
| 6:2.128 | ובקש |  |
| 6:2.129 | אחד |  |
| 6:2.130 | להמית |  |
| 6:2.131 | אדם |  |
| 6:2.132 | אחד |  |
| 6:2.133 | והשני |  |
| 6:2.134 | רוצה |  |
| 6:2.135 | להחיותו | . |
| 6:2.136 | ראוי |  |
| 6:2.137 | שיהיה |  |
| 6:2.138 | האיש |  |
| 6:2.139 | ההוא |  |
| 6:2.140 | מת |  |
| 6:2.141 | חי |  |
| 6:2.142 | מת |  |
| 6:2.143 | חי |  |
| 6:2.144 | לאלתר |  |
| 6:2.145 | כהרף |  |
| 6:2.146 | עין | . |
| 6:2.147 | ואם |  |
| 6:2.148 | יוכל |  |
| 6:2.149 | האחד |  |
| 6:2.150 | להסתיר |  |
| 6:2.151 | דבר |  |
| 6:2.152 | מחבירו |  |
| 6:2.153 | הרי |  |
| 6:2.154 | שניהם |  |
| 6:2.155 | נלאים | . |
| 6:2.156 | ועוד |  |
| 6:2.157 | כתב |  |
| 6:2.158 | לאלה |  |
| 6:2.159 | שפוקדים |  |
| 6:2.160 | על |  |
| 6:2.161 | נעשה |  |
| 6:2.162 | אדם |  |
| 6:2.163 | בצלמנו |  |
| 6:2.164 | היא |  |
| 6:2.165 | דעת |  |
| 6:2.166 | מחבר |  |
| 6:2.167 | הזהר |  |
| 6:2.168 | והמקובלים |  |
| 6:2.169 | כדלעיל |  |
| 6:2.170 | סי | ' |
| 6:2.171 | ל"ח | ) |
| 6:2.172 | כי |  |
| 6:2.173 | דרך |  |
| 6:2.174 | שפת |  |
| 6:2.175 | לשון |  |
| 6:2.176 | הקדש |  |
| 6:2.177 | שהגדולים |  |
| 6:2.178 | מכנים |  |
| 6:2.179 | עצמם |  |
| 6:2.180 | בלשון |  |
| 6:2.181 | רבים |  |
| 6:2.182 | אע"פ |  |
| 6:2.183 | שהם |  |
| 6:2.184 | יחידים | , |
| 6:2.185 | כמו |  |
| 6:2.186 | אולי |  |
| 6:2.187 | אוכל |  |
| 6:2.188 | נכה |  |
| 6:2.189 | בו | . |
| 6:2.190 | ובדניאל |  |
| 6:2.191 | נאמר |  |
| 6:2.192 | קדם |  |
| 6:2.193 | מלכא |  |
| 6:2.194 | ובמנוח |  |
| 6:2.195 | נעצרה |  |
| 6:2.196 | נא |  |
| 6:2.197 | אותך |  |
| 6:2.198 | ע"כ |  |
| 6:2.199 | ועוד |  |
| 6:2.200 | האריך |  |
| 6:2.201 | רבי |  |
| 6:2.202 | סעדיה |  |
| 6:2.203 | גאון |  |
| 6:2.204 | מאד | , |
| 6:2.205 | ואין |  |
| 6:2.206 | צורך |  |
| 6:2.207 | להאריך |  |
| 6:2.208 | כל |  |
| 6:2.209 | כך | , |
| 6:2.210 | כי |  |
| 6:2.211 | כל |  |
| 6:2.212 | ישראל |  |
| 6:2.213 | תקועים |  |
| 6:2.214 | באמונה |  |
| 6:2.215 | חזקה | , |
| 6:2.216 | כי |  |
| 6:2.217 | יוצר |  |
| 6:2.218 | הכל |  |
| 6:2.219 | אחד |  |
| 6:2.220 | מיוחד |  |
| 6:2.221 | כמו |  |
| 6:2.222 | שנ | ' |
| 6:2.223 | ה | ' |
| 6:2.224 | אחד | . |
| 6:2.225 | ונאמר |  |
| 6:2.226 | ראו |  |
| 6:2.227 | עתה |  |
| 6:2.228 | כי |  |
| 6:2.229 | אני |  |
| 6:2.230 | הוא |  |
| 6:2.231 | ואין |  |
| 6:2.232 | אלהים |  |
| 6:2.233 | עמדי | . |
| 6:2.234 | ובמקרא |  |
| 6:2.235 | יש |  |
| 6:2.236 | הרבה |  |
| 6:2.237 | מקראות |  |
| 6:2.238 | כאלה |  |
| 6:2.239 | עכ"ל |  |
| 6:2.240 | הרב |  |
| 6:2.241 | הסמ"ג | . |


## Milchamot ha-Shem 6:3

Hebrew (Wikisource): סב) הבט נא אחי בעין חודרת בדברי הר' הסמ"ג ז"ל איך הביא דברי הרס"ג לחזק דבריו. ובעיניו הדבר פשוט מאוד, שכן היא אמונת אחדותו ית' אחדות שאין כמוה בשאר כל האחדים. שלא כדברי הקבלה החדשה שחשבו שאחדותו ית' פחותה גם משאר האחדים, כי הזוג האחד הם שנים בלבד, ואצל המקובלים החדשים האלהות כוללת כמה פרצופים כמבואר לעיל, חמשה פרצופים כוללים. והחמשה נפרטים עוד לשנים עשר פרצופים בעולם האצילות, חוץ מן הפרצופים שבעולמות העליונים אין מספר שלא דברו בהם. וחוץ מהפרצופים שבבריאה ויצירה ועשייה, ואלה הפרצופים כל אחד מהם נקרא בשם הוי'ה ואדנות ואלהים ובשם הקב"ה ומדרגותיהם חלוקים זה מזה, זה למעלה מזה, וזה למטה ממנו. ואין רשות לפרצוף לברוא או להחיות או להמית אלא בנטילת רשות מאותו פרצוף שלמעלה ממנו, כשרים הממונים מתחת יד המלך שהמלך נותן רשות למשנהו. והמשנה לשליש והשליש לשרי אלפים ושרי אלפים למאות וכו'. וכל אחד יעשה חפצו בדבר שנתמנה לו, אבל דבר גדול לא יעשה מדעת עצמו אלא בנטילת רשות מהממונה שלמעלה ממנו כמבואר לעיל סי' כ"א בשם זוהר בראשית וכמו שפירשו וביארו דבריו המפרשים, ושהמלך העליון מכולם אדם קדמון. והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי אני אמית ואחיה, לפי שאדם קדמון זה הוא תחלת פרצופי האצילות, ולית לעילה מניה דימחי ביה. ויש לתמוה על זה למה אינו צריך ליטול רשות מאדם קדמאה שהוא למעלה ממנו אשר כל העולמות העליונים שאין להם מספר נתונים בחלל עיגוליו כמו שכתב המקדש מלך דף ט"ו ד"ה ורמז, או מהאין סוף שהוא האלוה העליון מעל כולם?

| id | word | punctuation after |
|---|---|---|
| 6:3.1 | סב | ) |
| 6:3.2 | הבט |  |
| 6:3.3 | נא |  |
| 6:3.4 | אחי |  |
| 6:3.5 | בעין |  |
| 6:3.6 | חודרת |  |
| 6:3.7 | בדברי |  |
| 6:3.8 | הר | ' |
| 6:3.9 | הסמ"ג |  |
| 6:3.10 | ז"ל |  |
| 6:3.11 | איך |  |
| 6:3.12 | הביא |  |
| 6:3.13 | דברי |  |
| 6:3.14 | הרס"ג |  |
| 6:3.15 | לחזק |  |
| 6:3.16 | דבריו | . |
| 6:3.17 | ובעיניו |  |
| 6:3.18 | הדבר |  |
| 6:3.19 | פשוט |  |
| 6:3.20 | מאוד | , |
| 6:3.21 | שכן |  |
| 6:3.22 | היא |  |
| 6:3.23 | אמונת |  |
| 6:3.24 | אחדותו |  |
| 6:3.25 | ית | ' |
| 6:3.26 | אחדות |  |
| 6:3.27 | שאין |  |
| 6:3.28 | כמוה |  |
| 6:3.29 | בשאר |  |
| 6:3.30 | כל |  |
| 6:3.31 | האחדים | . |
| 6:3.32 | שלא |  |
| 6:3.33 | כדברי |  |
| 6:3.34 | הקבלה |  |
| 6:3.35 | החדשה |  |
| 6:3.36 | שחשבו |  |
| 6:3.37 | שאחדותו |  |
| 6:3.38 | ית | ' |
| 6:3.39 | פחותה |  |
| 6:3.40 | גם |  |
| 6:3.41 | משאר |  |
| 6:3.42 | האחדים | , |
| 6:3.43 | כי |  |
| 6:3.44 | הזוג |  |
| 6:3.45 | האחד |  |
| 6:3.46 | הם |  |
| 6:3.47 | שנים |  |
| 6:3.48 | בלבד | , |
| 6:3.49 | ואצל |  |
| 6:3.50 | המקובלים |  |
| 6:3.51 | החדשים |  |
| 6:3.52 | האלהות |  |
| 6:3.53 | כוללת |  |
| 6:3.54 | כמה |  |
| 6:3.55 | פרצופים |  |
| 6:3.56 | כמבואר |  |
| 6:3.57 | לעיל | , |
| 6:3.58 | חמשה |  |
| 6:3.59 | פרצופים |  |
| 6:3.60 | כוללים | . |
| 6:3.61 | והחמשה |  |
| 6:3.62 | נפרטים |  |
| 6:3.63 | עוד |  |
| 6:3.64 | לשנים |  |
| 6:3.65 | עשר |  |
| 6:3.66 | פרצופים |  |
| 6:3.67 | בעולם |  |
| 6:3.68 | האצילות | , |
| 6:3.69 | חוץ |  |
| 6:3.70 | מן |  |
| 6:3.71 | הפרצופים |  |
| 6:3.72 | שבעולמות |  |
| 6:3.73 | העליונים |  |
| 6:3.74 | אין |  |
| 6:3.75 | מספר |  |
| 6:3.76 | שלא |  |
| 6:3.77 | דברו |  |
| 6:3.78 | בהם | . |
| 6:3.79 | וחוץ |  |
| 6:3.80 | מהפרצופים |  |
| 6:3.81 | שבבריאה |  |
| 6:3.82 | ויצירה |  |
| 6:3.83 | ועשייה | , |
| 6:3.84 | ואלה |  |
| 6:3.85 | הפרצופים |  |
| 6:3.86 | כל |  |
| 6:3.87 | אחד |  |
| 6:3.88 | מהם |  |
| 6:3.89 | נקרא |  |
| 6:3.90 | בשם |  |
| 6:3.91 | הוי'ה |  |
| 6:3.92 | ואדנות |  |
| 6:3.93 | ואלהים |  |
| 6:3.94 | ובשם |  |
| 6:3.95 | הקב"ה |  |
| 6:3.96 | ומדרגותיהם |  |
| 6:3.97 | חלוקים |  |
| 6:3.98 | זה |  |
| 6:3.99 | מזה | , |
| 6:3.100 | זה |  |
| 6:3.101 | למעלה |  |
| 6:3.102 | מזה | , |
| 6:3.103 | וזה |  |
| 6:3.104 | למטה |  |
| 6:3.105 | ממנו | . |
| 6:3.106 | ואין |  |
| 6:3.107 | רשות |  |
| 6:3.108 | לפרצוף |  |
| 6:3.109 | לברוא |  |
| 6:3.110 | או |  |
| 6:3.111 | להחיות |  |
| 6:3.112 | או |  |
| 6:3.113 | להמית |  |
| 6:3.114 | אלא |  |
| 6:3.115 | בנטילת |  |
| 6:3.116 | רשות |  |
| 6:3.117 | מאותו |  |
| 6:3.118 | פרצוף |  |
| 6:3.119 | שלמעלה |  |
| 6:3.120 | ממנו | , |
| 6:3.121 | כשרים |  |
| 6:3.122 | הממונים |  |
| 6:3.123 | מתחת |  |
| 6:3.124 | יד |  |
| 6:3.125 | המלך |  |
| 6:3.126 | שהמלך |  |
| 6:3.127 | נותן |  |
| 6:3.128 | רשות |  |
| 6:3.129 | למשנהו | . |
| 6:3.130 | והמשנה |  |
| 6:3.131 | לשליש |  |
| 6:3.132 | והשליש |  |
| 6:3.133 | לשרי |  |
| 6:3.134 | אלפים |  |
| 6:3.135 | ושרי |  |
| 6:3.136 | אלפים |  |
| 6:3.137 | למאות |  |
| 6:3.138 | וכו | '. |
| 6:3.139 | וכל |  |
| 6:3.140 | אחד |  |
| 6:3.141 | יעשה |  |
| 6:3.142 | חפצו |  |
| 6:3.143 | בדבר |  |
| 6:3.144 | שנתמנה |  |
| 6:3.145 | לו | , |
| 6:3.146 | אבל |  |
| 6:3.147 | דבר |  |
| 6:3.148 | גדול |  |
| 6:3.149 | לא |  |
| 6:3.150 | יעשה |  |
| 6:3.151 | מדעת |  |
| 6:3.152 | עצמו |  |
| 6:3.153 | אלא |  |
| 6:3.154 | בנטילת |  |
| 6:3.155 | רשות |  |
| 6:3.156 | מהממונה |  |
| 6:3.157 | שלמעלה |  |
| 6:3.158 | ממנו |  |
| 6:3.159 | כמבואר |  |
| 6:3.160 | לעיל |  |
| 6:3.161 | סי | ' |
| 6:3.162 | כ"א |  |
| 6:3.163 | בשם |  |
| 6:3.164 | זוהר |  |
| 6:3.165 | בראשית |  |
| 6:3.166 | וכמו |  |
| 6:3.167 | שפירשו |  |
| 6:3.168 | וביארו |  |
| 6:3.169 | דבריו |  |
| 6:3.170 | המפרשים | , |
| 6:3.171 | ושהמלך |  |
| 6:3.172 | העליון |  |
| 6:3.173 | מכולם |  |
| 6:3.174 | אדם |  |
| 6:3.175 | קדמון | . |
| 6:3.176 | והוא |  |
| 6:3.177 | שאמר |  |
| 6:3.178 | ראו |  |
| 6:3.179 | עתה |  |
| 6:3.180 | כי |  |
| 6:3.181 | אני |  |
| 6:3.182 | אני |  |
| 6:3.183 | הוא |  |
| 6:3.184 | ואין |  |
| 6:3.185 | אלהים |  |
| 6:3.186 | עמדי |  |
| 6:3.187 | אני |  |
| 6:3.188 | אמית |  |
| 6:3.189 | ואחיה | , |
| 6:3.190 | לפי |  |
| 6:3.191 | שאדם |  |
| 6:3.192 | קדמון |  |
| 6:3.193 | זה |  |
| 6:3.194 | הוא |  |
| 6:3.195 | תחלת |  |
| 6:3.196 | פרצופי |  |
| 6:3.197 | האצילות | , |
| 6:3.198 | ולית |  |
| 6:3.199 | לעילה |  |
| 6:3.200 | מניה |  |
| 6:3.201 | דימחי |  |
| 6:3.202 | ביה | . |
| 6:3.203 | ויש |  |
| 6:3.204 | לתמוה |  |
| 6:3.205 | על |  |
| 6:3.206 | זה |  |
| 6:3.207 | למה |  |
| 6:3.208 | אינו |  |
| 6:3.209 | צריך |  |
| 6:3.210 | ליטול |  |
| 6:3.211 | רשות |  |
| 6:3.212 | מאדם |  |
| 6:3.213 | קדמאה |  |
| 6:3.214 | שהוא |  |
| 6:3.215 | למעלה |  |
| 6:3.216 | ממנו |  |
| 6:3.217 | אשר |  |
| 6:3.218 | כל |  |
| 6:3.219 | העולמות |  |
| 6:3.220 | העליונים |  |
| 6:3.221 | שאין |  |
| 6:3.222 | להם |  |
| 6:3.223 | מספר |  |
| 6:3.224 | נתונים |  |
| 6:3.225 | בחלל |  |
| 6:3.226 | עיגוליו |  |
| 6:3.227 | כמו |  |
| 6:3.228 | שכתב |  |
| 6:3.229 | המקדש |  |
| 6:3.230 | מלך |  |
| 6:3.231 | דף |  |
| 6:3.232 | ט"ו |  |
| 6:3.233 | ד"ה |  |
| 6:3.234 | ורמז | , |
| 6:3.235 | או |  |
| 6:3.236 | מהאין |  |
| 6:3.237 | סוף |  |
| 6:3.238 | שהוא |  |
| 6:3.239 | האלוה |  |
| 6:3.240 | העליון |  |
| 6:3.241 | מעל |  |
| 6:3.242 | כולם | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

