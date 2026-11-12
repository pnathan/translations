# Translation pack: milhamot-hashem perek-7-02 (Milchamot ha-Shem 7:3–7:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-02`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted"; "tzimtzum" in a
  note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
- גוף **דק** → "a subtle body".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **ידידי** → "my friend" (second person kept).
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
- הבל **וריק** → "vanity and emptiness".
- מי **אתי** (ketiv): translated, the qere noted.
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 7:3

Hebrew (Wikisource): התבונן נא ידידי בדברי הרמב"ם ושבילי אמונה וחידושי גאונים, ואת כל ונוכחת כי יחוד אלהינו ברוך הוא אינו דומה לשאר אחדים שהם נחלקים לכמה מחלקות, והיא היא קבלתינו האמתית שקבלנו מאבותינו דור אחר דור עד משה משה רבי' ועד אברהם אבינו, וכמו שכבר ביארתי לעיל והבאתי לשון המדרש רבה ומדרש הגדול וכמה כתובים ומאמרי רז"ל בגמרא ושאר מדרשים אמתיים מוכיחים כן, כי הקב"ה עומד באחדות אחת לא שנה ולא ישתנה, לא כדעת המקובלים החדשים שנמשכו אחר הדעה החיצונית אשר השיאם מחבר הזוהר שאמרו שהסבה הראשונה הנקרא אצלם אין סוף נחלק למחלקות רבות, ונשתנה שינוים רבים, וחלק הגדול ממנו נשאר אחר הצמצום כדמות דור עגול וחלל באמצעיתו גדול מאוד, אשר בתוך החלל הזה כל העולמות העליונים והתחתונים, מעולמות שלמעלה מעולם האצילות עם עולם האצילות ועם העולמות שלמטה ממנו, וממנו נמשך קו אחד דק כעין צינור אל תוך החלל הזה אשר ממנו נתהוו כל העולמות וממנו נספעים, וז"ל הרח"ו בעץ חיים שער עגולים ויושר ענף ב'. והנה אחר הצמצום אשר נשאר מקום החלל ואויר פנוי וריקן באמצע האור של אין סוף ממש הנה כבר היה מקום לשיוכלו להיות שם הנאצלים והנבראים והיצורים והנעשים, ואז המשיך מן האור האין סוף קו אחד דק ישר מן האור עגול שלו מלמעלה למטה, ומשתלשל ויורד למטה תוך החלל ההוא וראש העליון של הקו נמשך מן האין סוף עצמו ונוגע בו, אמנם סיום הקו הזה למטה בסופו אינו נוגע באור האין סוף, ודרך הקו הזה נמשך ונתפשט אור אין סוף למטה ובמקום החלל ההוא האציל וברא ויצר ועשה כל העולמות כולם, וקו זה כעין צנור דק אשר דרך בו מתפשט ונמשך מימי האור העליון של אין סוף אל כל העולמות אשר במקום האויר הוהחלל ההוא, ונבאר עתה ענין חקירת המקובלים לדעת איך יש ראש, תוך וסוף וכו' עיין לעיל סי' ל"ח שכבר הבאתיו שם.

| id | word | punctuation after |
|---|---|---|
| 7:3.1 | התבונן |  |
| 7:3.2 | נא |  |
| 7:3.3 | ידידי |  |
| 7:3.4 | בדברי |  |
| 7:3.5 | הרמב"ם |  |
| 7:3.6 | ושבילי |  |
| 7:3.7 | אמונה |  |
| 7:3.8 | וחידושי |  |
| 7:3.9 | גאונים | , |
| 7:3.10 | ואת |  |
| 7:3.11 | כל |  |
| 7:3.12 | ונוכחת |  |
| 7:3.13 | כי |  |
| 7:3.14 | יחוד |  |
| 7:3.15 | אלהינו |  |
| 7:3.16 | ברוך |  |
| 7:3.17 | הוא |  |
| 7:3.18 | אינו |  |
| 7:3.19 | דומה |  |
| 7:3.20 | לשאר |  |
| 7:3.21 | אחדים |  |
| 7:3.22 | שהם |  |
| 7:3.23 | נחלקים |  |
| 7:3.24 | לכמה |  |
| 7:3.25 | מחלקות | , |
| 7:3.26 | והיא |  |
| 7:3.27 | היא |  |
| 7:3.28 | קבלתינו |  |
| 7:3.29 | האמתית |  |
| 7:3.30 | שקבלנו |  |
| 7:3.31 | מאבותינו |  |
| 7:3.32 | דור |  |
| 7:3.33 | אחר |  |
| 7:3.34 | דור |  |
| 7:3.35 | עד |  |
| 7:3.36 | משה |  |
| 7:3.37 | משה |  |
| 7:3.38 | רבי | ' |
| 7:3.39 | ועד |  |
| 7:3.40 | אברהם |  |
| 7:3.41 | אבינו | , |
| 7:3.42 | וכמו |  |
| 7:3.43 | שכבר |  |
| 7:3.44 | ביארתי |  |
| 7:3.45 | לעיל |  |
| 7:3.46 | והבאתי |  |
| 7:3.47 | לשון |  |
| 7:3.48 | המדרש |  |
| 7:3.49 | רבה |  |
| 7:3.50 | ומדרש |  |
| 7:3.51 | הגדול |  |
| 7:3.52 | וכמה |  |
| 7:3.53 | כתובים |  |
| 7:3.54 | ומאמרי |  |
| 7:3.55 | רז"ל |  |
| 7:3.56 | בגמרא |  |
| 7:3.57 | ושאר |  |
| 7:3.58 | מדרשים |  |
| 7:3.59 | אמתיים |  |
| 7:3.60 | מוכיחים |  |
| 7:3.61 | כן | , |
| 7:3.62 | כי |  |
| 7:3.63 | הקב"ה |  |
| 7:3.64 | עומד |  |
| 7:3.65 | באחדות |  |
| 7:3.66 | אחת |  |
| 7:3.67 | לא |  |
| 7:3.68 | שנה |  |
| 7:3.69 | ולא |  |
| 7:3.70 | ישתנה | , |
| 7:3.71 | לא |  |
| 7:3.72 | כדעת |  |
| 7:3.73 | המקובלים |  |
| 7:3.74 | החדשים |  |
| 7:3.75 | שנמשכו |  |
| 7:3.76 | אחר |  |
| 7:3.77 | הדעה |  |
| 7:3.78 | החיצונית |  |
| 7:3.79 | אשר |  |
| 7:3.80 | השיאם |  |
| 7:3.81 | מחבר |  |
| 7:3.82 | הזוהר |  |
| 7:3.83 | שאמרו |  |
| 7:3.84 | שהסבה |  |
| 7:3.85 | הראשונה |  |
| 7:3.86 | הנקרא |  |
| 7:3.87 | אצלם |  |
| 7:3.88 | אין |  |
| 7:3.89 | סוף |  |
| 7:3.90 | נחלק |  |
| 7:3.91 | למחלקות |  |
| 7:3.92 | רבות | , |
| 7:3.93 | ונשתנה |  |
| 7:3.94 | שינוים |  |
| 7:3.95 | רבים | , |
| 7:3.96 | וחלק |  |
| 7:3.97 | הגדול |  |
| 7:3.98 | ממנו |  |
| 7:3.99 | נשאר |  |
| 7:3.100 | אחר |  |
| 7:3.101 | הצמצום |  |
| 7:3.102 | כדמות |  |
| 7:3.103 | דור |  |
| 7:3.104 | עגול |  |
| 7:3.105 | וחלל |  |
| 7:3.106 | באמצעיתו |  |
| 7:3.107 | גדול |  |
| 7:3.108 | מאוד | , |
| 7:3.109 | אשר |  |
| 7:3.110 | בתוך |  |
| 7:3.111 | החלל |  |
| 7:3.112 | הזה |  |
| 7:3.113 | כל |  |
| 7:3.114 | העולמות |  |
| 7:3.115 | העליונים |  |
| 7:3.116 | והתחתונים | , |
| 7:3.117 | מעולמות |  |
| 7:3.118 | שלמעלה |  |
| 7:3.119 | מעולם |  |
| 7:3.120 | האצילות |  |
| 7:3.121 | עם |  |
| 7:3.122 | עולם |  |
| 7:3.123 | האצילות |  |
| 7:3.124 | ועם |  |
| 7:3.125 | העולמות |  |
| 7:3.126 | שלמטה |  |
| 7:3.127 | ממנו | , |
| 7:3.128 | וממנו |  |
| 7:3.129 | נמשך |  |
| 7:3.130 | קו |  |
| 7:3.131 | אחד |  |
| 7:3.132 | דק |  |
| 7:3.133 | כעין |  |
| 7:3.134 | צינור |  |
| 7:3.135 | אל |  |
| 7:3.136 | תוך |  |
| 7:3.137 | החלל |  |
| 7:3.138 | הזה |  |
| 7:3.139 | אשר |  |
| 7:3.140 | ממנו |  |
| 7:3.141 | נתהוו |  |
| 7:3.142 | כל |  |
| 7:3.143 | העולמות |  |
| 7:3.144 | וממנו |  |
| 7:3.145 | נספעים | , |
| 7:3.146 | וז"ל |  |
| 7:3.147 | הרח"ו |  |
| 7:3.148 | בעץ |  |
| 7:3.149 | חיים |  |
| 7:3.150 | שער |  |
| 7:3.151 | עגולים |  |
| 7:3.152 | ויושר |  |
| 7:3.153 | ענף |  |
| 7:3.154 | ב | '. |
| 7:3.155 | והנה |  |
| 7:3.156 | אחר |  |
| 7:3.157 | הצמצום |  |
| 7:3.158 | אשר |  |
| 7:3.159 | נשאר |  |
| 7:3.160 | מקום |  |
| 7:3.161 | החלל |  |
| 7:3.162 | ואויר |  |
| 7:3.163 | פנוי |  |
| 7:3.164 | וריקן |  |
| 7:3.165 | באמצע |  |
| 7:3.166 | האור |  |
| 7:3.167 | של |  |
| 7:3.168 | אין |  |
| 7:3.169 | סוף |  |
| 7:3.170 | ממש |  |
| 7:3.171 | הנה |  |
| 7:3.172 | כבר |  |
| 7:3.173 | היה |  |
| 7:3.174 | מקום |  |
| 7:3.175 | לשיוכלו |  |
| 7:3.176 | להיות |  |
| 7:3.177 | שם |  |
| 7:3.178 | הנאצלים |  |
| 7:3.179 | והנבראים |  |
| 7:3.180 | והיצורים |  |
| 7:3.181 | והנעשים | , |
| 7:3.182 | ואז |  |
| 7:3.183 | המשיך |  |
| 7:3.184 | מן |  |
| 7:3.185 | האור |  |
| 7:3.186 | האין |  |
| 7:3.187 | סוף |  |
| 7:3.188 | קו |  |
| 7:3.189 | אחד |  |
| 7:3.190 | דק |  |
| 7:3.191 | ישר |  |
| 7:3.192 | מן |  |
| 7:3.193 | האור |  |
| 7:3.194 | עגול |  |
| 7:3.195 | שלו |  |
| 7:3.196 | מלמעלה |  |
| 7:3.197 | למטה | , |
| 7:3.198 | ומשתלשל |  |
| 7:3.199 | ויורד |  |
| 7:3.200 | למטה |  |
| 7:3.201 | תוך |  |
| 7:3.202 | החלל |  |
| 7:3.203 | ההוא |  |
| 7:3.204 | וראש |  |
| 7:3.205 | העליון |  |
| 7:3.206 | של |  |
| 7:3.207 | הקו |  |
| 7:3.208 | נמשך |  |
| 7:3.209 | מן |  |
| 7:3.210 | האין |  |
| 7:3.211 | סוף |  |
| 7:3.212 | עצמו |  |
| 7:3.213 | ונוגע |  |
| 7:3.214 | בו | , |
| 7:3.215 | אמנם |  |
| 7:3.216 | סיום |  |
| 7:3.217 | הקו |  |
| 7:3.218 | הזה |  |
| 7:3.219 | למטה |  |
| 7:3.220 | בסופו |  |
| 7:3.221 | אינו |  |
| 7:3.222 | נוגע |  |
| 7:3.223 | באור |  |
| 7:3.224 | האין |  |
| 7:3.225 | סוף | , |
| 7:3.226 | ודרך |  |
| 7:3.227 | הקו |  |
| 7:3.228 | הזה |  |
| 7:3.229 | נמשך |  |
| 7:3.230 | ונתפשט |  |
| 7:3.231 | אור |  |
| 7:3.232 | אין |  |
| 7:3.233 | סוף |  |
| 7:3.234 | למטה |  |
| 7:3.235 | ובמקום |  |
| 7:3.236 | החלל |  |
| 7:3.237 | ההוא |  |
| 7:3.238 | האציל |  |
| 7:3.239 | וברא |  |
| 7:3.240 | ויצר |  |
| 7:3.241 | ועשה |  |
| 7:3.242 | כל |  |
| 7:3.243 | העולמות |  |
| 7:3.244 | כולם | , |
| 7:3.245 | וקו |  |
| 7:3.246 | זה |  |
| 7:3.247 | כעין |  |
| 7:3.248 | צנור |  |
| 7:3.249 | דק |  |
| 7:3.250 | אשר |  |
| 7:3.251 | דרך |  |
| 7:3.252 | בו |  |
| 7:3.253 | מתפשט |  |
| 7:3.254 | ונמשך |  |
| 7:3.255 | מימי |  |
| 7:3.256 | האור |  |
| 7:3.257 | העליון |  |
| 7:3.258 | של |  |
| 7:3.259 | אין |  |
| 7:3.260 | סוף |  |
| 7:3.261 | אל |  |
| 7:3.262 | כל |  |
| 7:3.263 | העולמות |  |
| 7:3.264 | אשר |  |
| 7:3.265 | במקום |  |
| 7:3.266 | האויר |  |
| 7:3.267 | הוהחלל |  |
| 7:3.268 | ההוא | , |
| 7:3.269 | ונבאר |  |
| 7:3.270 | עתה |  |
| 7:3.271 | ענין |  |
| 7:3.272 | חקירת |  |
| 7:3.273 | המקובלים |  |
| 7:3.274 | לדעת |  |
| 7:3.275 | איך |  |
| 7:3.276 | יש |  |
| 7:3.277 | ראש | , |
| 7:3.278 | תוך |  |
| 7:3.279 | וסוף |  |
| 7:3.280 | וכו | ' |
| 7:3.281 | עיין |  |
| 7:3.282 | לעיל |  |
| 7:3.283 | סי | ' |
| 7:3.284 | ל"ח |  |
| 7:3.285 | שכבר |  |
| 7:3.286 | הבאתיו |  |
| 7:3.287 | שם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

