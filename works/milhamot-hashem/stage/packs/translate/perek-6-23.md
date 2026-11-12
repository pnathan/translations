# Translation pack: milhamot-hashem perek-6-23 (Milchamot ha-Shem 6:34–6:36)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-23.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-23`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **רז"ל** → "our Rabbis of blessed memory".
- **האר"י** → "Ari" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
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
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **תוה"ק** → "the holy Torah".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 6:34

Hebrew (Wikisource): עו) ועוד האריך בעל שבילי אמונה בנתיב הנז' לעיל להוכיח בראיות את קבלת תוה"ק שהשם יתברך הוא אחד ומיוחד מכל שאר האחרים ושאין בכל שאר האחדים אחד אמתי אלא הוא ית' ושאינו נחלק לכמה חלקים אשר בהתחברם יחד נקראים אחד. וזה שלא כדעת מחבר הזהר והמקובלים המאמינים בו שסוברים שכלל הפרצופים דוקא הם אחד, ואף על פי שהם אומרים שהאין סוף הוא אחדות אחת, הרי הם סותרים דברי עצמן, שהם אומרים שהוא מתחלק ומתפשט ומתלבש בכמה פרצופים, וגם אמרו שהאין סוף בבריאה אינו אחד כמו באצילות.

| id | word | punctuation after |
|---|---|---|
| 6:34.1 | עו | ) |
| 6:34.2 | ועוד |  |
| 6:34.3 | האריך |  |
| 6:34.4 | בעל |  |
| 6:34.5 | שבילי |  |
| 6:34.6 | אמונה |  |
| 6:34.7 | בנתיב |  |
| 6:34.8 | הנז | ' |
| 6:34.9 | לעיל |  |
| 6:34.10 | להוכיח |  |
| 6:34.11 | בראיות |  |
| 6:34.12 | את |  |
| 6:34.13 | קבלת |  |
| 6:34.14 | תוה"ק |  |
| 6:34.15 | שהשם |  |
| 6:34.16 | יתברך |  |
| 6:34.17 | הוא |  |
| 6:34.18 | אחד |  |
| 6:34.19 | ומיוחד |  |
| 6:34.20 | מכל |  |
| 6:34.21 | שאר |  |
| 6:34.22 | האחרים |  |
| 6:34.23 | ושאין |  |
| 6:34.24 | בכל |  |
| 6:34.25 | שאר |  |
| 6:34.26 | האחדים |  |
| 6:34.27 | אחד |  |
| 6:34.28 | אמתי |  |
| 6:34.29 | אלא |  |
| 6:34.30 | הוא |  |
| 6:34.31 | ית | ' |
| 6:34.32 | ושאינו |  |
| 6:34.33 | נחלק |  |
| 6:34.34 | לכמה |  |
| 6:34.35 | חלקים |  |
| 6:34.36 | אשר |  |
| 6:34.37 | בהתחברם |  |
| 6:34.38 | יחד |  |
| 6:34.39 | נקראים |  |
| 6:34.40 | אחד | . |
| 6:34.41 | וזה |  |
| 6:34.42 | שלא |  |
| 6:34.43 | כדעת |  |
| 6:34.44 | מחבר |  |
| 6:34.45 | הזהר |  |
| 6:34.46 | והמקובלים |  |
| 6:34.47 | המאמינים |  |
| 6:34.48 | בו |  |
| 6:34.49 | שסוברים |  |
| 6:34.50 | שכלל |  |
| 6:34.51 | הפרצופים |  |
| 6:34.52 | דוקא |  |
| 6:34.53 | הם |  |
| 6:34.54 | אחד | , |
| 6:34.55 | ואף |  |
| 6:34.56 | על |  |
| 6:34.57 | פי |  |
| 6:34.58 | שהם |  |
| 6:34.59 | אומרים |  |
| 6:34.60 | שהאין |  |
| 6:34.61 | סוף |  |
| 6:34.62 | הוא |  |
| 6:34.63 | אחדות |  |
| 6:34.64 | אחת | , |
| 6:34.65 | הרי |  |
| 6:34.66 | הם |  |
| 6:34.67 | סותרים |  |
| 6:34.68 | דברי |  |
| 6:34.69 | עצמן | , |
| 6:34.70 | שהם |  |
| 6:34.71 | אומרים |  |
| 6:34.72 | שהוא |  |
| 6:34.73 | מתחלק |  |
| 6:34.74 | ומתפשט |  |
| 6:34.75 | ומתלבש |  |
| 6:34.76 | בכמה |  |
| 6:34.77 | פרצופים | , |
| 6:34.78 | וגם |  |
| 6:34.79 | אמרו |  |
| 6:34.80 | שהאין |  |
| 6:34.81 | סוף |  |
| 6:34.82 | בבריאה |  |
| 6:34.83 | אינו |  |
| 6:34.84 | אחד |  |
| 6:34.85 | כמו |  |
| 6:34.86 | באצילות | . |


## Milchamot ha-Shem 6:35

Hebrew (Wikisource): ויותר קשה מאד מזה שהם כללים ומשתפים ביחודו את החמשה פרצופים שהם אומרים שהם נבראים, כמו שמפורש בזהר פרשת נשא על זעיר אנפין, ששם דרש על פסוק ה' פעלך בקרב שנים חייהו, שה' הנז' בפסוק הוא מדבר בעתיק יומין, ואומר לו ה' שהוא עתיק יומין, פעלך שהוא זעיר אנפין בקרב שנים חייהו וכו', וכן מבואר בהדיא בספר יושר לבב שאמר ועבדהו לזעיר אנפין אע"פ שהוא נברא, כי לנשמתו אתה עובד כמו שנתבאר לעיל סוף סי' מ"ח.

| id | word | punctuation after |
|---|---|---|
| 6:35.1 | ויותר |  |
| 6:35.2 | קשה |  |
| 6:35.3 | מאד |  |
| 6:35.4 | מזה |  |
| 6:35.5 | שהם |  |
| 6:35.6 | כללים |  |
| 6:35.7 | ומשתפים |  |
| 6:35.8 | ביחודו |  |
| 6:35.9 | את |  |
| 6:35.10 | החמשה |  |
| 6:35.11 | פרצופים |  |
| 6:35.12 | שהם |  |
| 6:35.13 | אומרים |  |
| 6:35.14 | שהם |  |
| 6:35.15 | נבראים | , |
| 6:35.16 | כמו |  |
| 6:35.17 | שמפורש |  |
| 6:35.18 | בזהר |  |
| 6:35.19 | פרשת |  |
| 6:35.20 | נשא |  |
| 6:35.21 | על |  |
| 6:35.22 | זעיר |  |
| 6:35.23 | אנפין | , |
| 6:35.24 | ששם |  |
| 6:35.25 | דרש |  |
| 6:35.26 | על |  |
| 6:35.27 | פסוק |  |
| 6:35.28 | ה | ' |
| 6:35.29 | פעלך |  |
| 6:35.30 | בקרב |  |
| 6:35.31 | שנים |  |
| 6:35.32 | חייהו | , |
| 6:35.33 | שה | ' |
| 6:35.34 | הנז | ' |
| 6:35.35 | בפסוק |  |
| 6:35.36 | הוא |  |
| 6:35.37 | מדבר |  |
| 6:35.38 | בעתיק |  |
| 6:35.39 | יומין | , |
| 6:35.40 | ואומר |  |
| 6:35.41 | לו |  |
| 6:35.42 | ה | ' |
| 6:35.43 | שהוא |  |
| 6:35.44 | עתיק |  |
| 6:35.45 | יומין | , |
| 6:35.46 | פעלך |  |
| 6:35.47 | שהוא |  |
| 6:35.48 | זעיר |  |
| 6:35.49 | אנפין |  |
| 6:35.50 | בקרב |  |
| 6:35.51 | שנים |  |
| 6:35.52 | חייהו |  |
| 6:35.53 | וכו | ', |
| 6:35.54 | וכן |  |
| 6:35.55 | מבואר |  |
| 6:35.56 | בהדיא |  |
| 6:35.57 | בספר |  |
| 6:35.58 | יושר |  |
| 6:35.59 | לבב |  |
| 6:35.60 | שאמר |  |
| 6:35.61 | ועבדהו |  |
| 6:35.62 | לזעיר |  |
| 6:35.63 | אנפין |  |
| 6:35.64 | אע"פ |  |
| 6:35.65 | שהוא |  |
| 6:35.66 | נברא | , |
| 6:35.67 | כי |  |
| 6:35.68 | לנשמתו |  |
| 6:35.69 | אתה |  |
| 6:35.70 | עובד |  |
| 6:35.71 | כמו |  |
| 6:35.72 | שנתבאר |  |
| 6:35.73 | לעיל |  |
| 6:35.74 | סוף |  |
| 6:35.75 | סי | ' |
| 6:35.76 | מ"ח | . |


## Milchamot ha-Shem 6:36

Hebrew (Wikisource): עוד האריך בעל שבילי אמונה להוכיח שהוא יתברך עילת כל העילות וסבת כל הסבות, ואין שום דבר מכל הנאצלים והנבראים קודם לו, ואין ראשית לראשיתו (שלא כדעת המקובלים שאמרו הקב"ה הנקרא בשם הוי"ה הוא הזעיר אנפין והוא לפי דבריהם ה' אלהינו) ושאינ וגוף, ולא כח בגוף וכל הענינים אשר יסופר ויתואר בהם הבורא ית' אינם אלא מושאלים ועל זה וכיוצא בו אמרו רז"ל דברה תורה כלשון בני אדם. וכן נאמר בכל ענין הדומה לזה שיש בספר התורה, כמו גמול העוה"ב ועונשיו, כי התורה קצרה בביאורו מפני שסמכה בהם על השכל ורמזה עליהם ברמזים להתעורר עליהם מי שיוכל לחקור ולדרוש אותם ולהבינם כמו שנאמר ומבקשי ה' יבינו כל.

| id | word | punctuation after |
|---|---|---|
| 6:36.1 | עוד |  |
| 6:36.2 | האריך |  |
| 6:36.3 | בעל |  |
| 6:36.4 | שבילי |  |
| 6:36.5 | אמונה |  |
| 6:36.6 | להוכיח |  |
| 6:36.7 | שהוא |  |
| 6:36.8 | יתברך |  |
| 6:36.9 | עילת |  |
| 6:36.10 | כל |  |
| 6:36.11 | העילות |  |
| 6:36.12 | וסבת |  |
| 6:36.13 | כל |  |
| 6:36.14 | הסבות | , |
| 6:36.15 | ואין |  |
| 6:36.16 | שום |  |
| 6:36.17 | דבר |  |
| 6:36.18 | מכל |  |
| 6:36.19 | הנאצלים |  |
| 6:36.20 | והנבראים |  |
| 6:36.21 | קודם |  |
| 6:36.22 | לו | , |
| 6:36.23 | ואין |  |
| 6:36.24 | ראשית |  |
| 6:36.25 | לראשיתו |  |
| 6:36.26 | שלא |  |
| 6:36.27 | כדעת |  |
| 6:36.28 | המקובלים |  |
| 6:36.29 | שאמרו |  |
| 6:36.30 | הקב"ה |  |
| 6:36.31 | הנקרא |  |
| 6:36.32 | בשם |  |
| 6:36.33 | הוי"ה |  |
| 6:36.34 | הוא |  |
| 6:36.35 | הזעיר |  |
| 6:36.36 | אנפין |  |
| 6:36.37 | והוא |  |
| 6:36.38 | לפי |  |
| 6:36.39 | דבריהם |  |
| 6:36.40 | ה | ' |
| 6:36.41 | אלהינו | ) |
| 6:36.42 | ושאינ |  |
| 6:36.43 | וגוף | , |
| 6:36.44 | ולא |  |
| 6:36.45 | כח |  |
| 6:36.46 | בגוף |  |
| 6:36.47 | וכל |  |
| 6:36.48 | הענינים |  |
| 6:36.49 | אשר |  |
| 6:36.50 | יסופר |  |
| 6:36.51 | ויתואר |  |
| 6:36.52 | בהם |  |
| 6:36.53 | הבורא |  |
| 6:36.54 | ית | ' |
| 6:36.55 | אינם |  |
| 6:36.56 | אלא |  |
| 6:36.57 | מושאלים |  |
| 6:36.58 | ועל |  |
| 6:36.59 | זה |  |
| 6:36.60 | וכיוצא |  |
| 6:36.61 | בו |  |
| 6:36.62 | אמרו |  |
| 6:36.63 | רז"ל |  |
| 6:36.64 | דברה |  |
| 6:36.65 | תורה |  |
| 6:36.66 | כלשון |  |
| 6:36.67 | בני |  |
| 6:36.68 | אדם | . |
| 6:36.69 | וכן |  |
| 6:36.70 | נאמר |  |
| 6:36.71 | בכל |  |
| 6:36.72 | ענין |  |
| 6:36.73 | הדומה |  |
| 6:36.74 | לזה |  |
| 6:36.75 | שיש |  |
| 6:36.76 | בספר |  |
| 6:36.77 | התורה | , |
| 6:36.78 | כמו |  |
| 6:36.79 | גמול |  |
| 6:36.80 | העוה"ב |  |
| 6:36.81 | ועונשיו | , |
| 6:36.82 | כי |  |
| 6:36.83 | התורה |  |
| 6:36.84 | קצרה |  |
| 6:36.85 | בביאורו |  |
| 6:36.86 | מפני |  |
| 6:36.87 | שסמכה |  |
| 6:36.88 | בהם |  |
| 6:36.89 | על |  |
| 6:36.90 | השכל |  |
| 6:36.91 | ורמזה |  |
| 6:36.92 | עליהם |  |
| 6:36.93 | ברמזים |  |
| 6:36.94 | להתעורר |  |
| 6:36.95 | עליהם |  |
| 6:36.96 | מי |  |
| 6:36.97 | שיוכל |  |
| 6:36.98 | לחקור |  |
| 6:36.99 | ולדרוש |  |
| 6:36.100 | אותם |  |
| 6:36.101 | ולהבינם |  |
| 6:36.102 | כמו |  |
| 6:36.103 | שנאמר |  |
| 6:36.104 | ומבקשי |  |
| 6:36.105 | ה | ' |
| 6:36.106 | יבינו |  |
| 6:36.107 | כל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

