# Translation pack: milhamot-hashem perek-10-03 (Milchamot ha-Shem 10:3–10:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-03`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **השונים** → "the Dualists".
- **בא"י** → "in the Land of Israel".

## Terms

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
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
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
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
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
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **השפע** → "the flow".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 10:3

Hebrew (Wikisource): אוי! כי על ידי הספר המתעה הזהר והתקונים דמינו לעמים עובדי האלילים, (הודיים ומצרים וכשדים פרס ומדי ושאר אמונות הבנוים על למודים נסתרים וסודות זרים אשר יסודתן בכח המדמה כאשר יראה המעיין בפסר דברי עולם לחחכם קלמן סולמאן). אשר לא יצאו מתוך ההפכה, וריב להם עם הדעת הישרה המקובלת מפי סופרים וספרים, ספרי רז"ל, המורים לאדם דעת ותבונה בתורתנו הקדושה. בל נטעה ונשגה ברעיונות העולים על לב ריק, לעבור על מאמר השי"ת ולא תתורו אחרי לבבכם זו מינות! ונתקיים בנו על ידי הזהר והתקונים מאמר רז"ל בסנהדרין דף ל"ט ובילקוט מלכים רמז רכ"ו, ריב"ל רמי כתיב וכמשפטי הגוים אשר סביבותיכם עשיתם, וכתיב וכמשפטי הגוים אשר סביבותיכם לא עשיתם? ומשני כמתוקנים שבהם (המאמינים באחדות הבורא) לא עשיתם, כמקולקלים שבהם הנז' לעיל עשיתם! כן אירע לנו באמונת האחדות והעבודה, כמתוקנים שבהם לא עשינו כמקולקלים שבהם עשינו לעבוד אל קצא אפים (ז"א) וליחדו עם האלילה זוגתו.

| id | word | punctuation after |
|---|---|---|
| 10:3.1 | אוי | ! |
| 10:3.2 | כי |  |
| 10:3.3 | על |  |
| 10:3.4 | ידי |  |
| 10:3.5 | הספר |  |
| 10:3.6 | המתעה |  |
| 10:3.7 | הזהר |  |
| 10:3.8 | והתקונים |  |
| 10:3.9 | דמינו |  |
| 10:3.10 | לעמים |  |
| 10:3.11 | עובדי |  |
| 10:3.12 | האלילים | , |
| 10:3.13 | הודיים |  |
| 10:3.14 | ומצרים |  |
| 10:3.15 | וכשדים |  |
| 10:3.16 | פרס |  |
| 10:3.17 | ומדי |  |
| 10:3.18 | ושאר |  |
| 10:3.19 | אמונות |  |
| 10:3.20 | הבנוים |  |
| 10:3.21 | על |  |
| 10:3.22 | למודים |  |
| 10:3.23 | נסתרים |  |
| 10:3.24 | וסודות |  |
| 10:3.25 | זרים |  |
| 10:3.26 | אשר |  |
| 10:3.27 | יסודתן |  |
| 10:3.28 | בכח |  |
| 10:3.29 | המדמה |  |
| 10:3.30 | כאשר |  |
| 10:3.31 | יראה |  |
| 10:3.32 | המעיין |  |
| 10:3.33 | בפסר |  |
| 10:3.34 | דברי |  |
| 10:3.35 | עולם |  |
| 10:3.36 | לחחכם |  |
| 10:3.37 | קלמן |  |
| 10:3.38 | סולמאן | ). |
| 10:3.39 | אשר |  |
| 10:3.40 | לא |  |
| 10:3.41 | יצאו |  |
| 10:3.42 | מתוך |  |
| 10:3.43 | ההפכה | , |
| 10:3.44 | וריב |  |
| 10:3.45 | להם |  |
| 10:3.46 | עם |  |
| 10:3.47 | הדעת |  |
| 10:3.48 | הישרה |  |
| 10:3.49 | המקובלת |  |
| 10:3.50 | מפי |  |
| 10:3.51 | סופרים |  |
| 10:3.52 | וספרים | , |
| 10:3.53 | ספרי |  |
| 10:3.54 | רז"ל | , |
| 10:3.55 | המורים |  |
| 10:3.56 | לאדם |  |
| 10:3.57 | דעת |  |
| 10:3.58 | ותבונה |  |
| 10:3.59 | בתורתנו |  |
| 10:3.60 | הקדושה | . |
| 10:3.61 | בל |  |
| 10:3.62 | נטעה |  |
| 10:3.63 | ונשגה |  |
| 10:3.64 | ברעיונות |  |
| 10:3.65 | העולים |  |
| 10:3.66 | על |  |
| 10:3.67 | לב |  |
| 10:3.68 | ריק | , |
| 10:3.69 | לעבור |  |
| 10:3.70 | על |  |
| 10:3.71 | מאמר |  |
| 10:3.72 | השי"ת |  |
| 10:3.73 | ולא |  |
| 10:3.74 | תתורו |  |
| 10:3.75 | אחרי |  |
| 10:3.76 | לבבכם |  |
| 10:3.77 | זו |  |
| 10:3.78 | מינות | ! |
| 10:3.79 | ונתקיים |  |
| 10:3.80 | בנו |  |
| 10:3.81 | על |  |
| 10:3.82 | ידי |  |
| 10:3.83 | הזהר |  |
| 10:3.84 | והתקונים |  |
| 10:3.85 | מאמר |  |
| 10:3.86 | רז"ל |  |
| 10:3.87 | בסנהדרין |  |
| 10:3.88 | דף |  |
| 10:3.89 | ל"ט |  |
| 10:3.90 | ובילקוט |  |
| 10:3.91 | מלכים |  |
| 10:3.92 | רמז |  |
| 10:3.93 | רכ"ו | , |
| 10:3.94 | ריב"ל |  |
| 10:3.95 | רמי |  |
| 10:3.96 | כתיב |  |
| 10:3.97 | וכמשפטי |  |
| 10:3.98 | הגוים |  |
| 10:3.99 | אשר |  |
| 10:3.100 | סביבותיכם |  |
| 10:3.101 | עשיתם | , |
| 10:3.102 | וכתיב |  |
| 10:3.103 | וכמשפטי |  |
| 10:3.104 | הגוים |  |
| 10:3.105 | אשר |  |
| 10:3.106 | סביבותיכם |  |
| 10:3.107 | לא |  |
| 10:3.108 | עשיתם | ? |
| 10:3.109 | ומשני |  |
| 10:3.110 | כמתוקנים |  |
| 10:3.111 | שבהם |  |
| 10:3.112 | המאמינים |  |
| 10:3.113 | באחדות |  |
| 10:3.114 | הבורא | ) |
| 10:3.115 | לא |  |
| 10:3.116 | עשיתם | , |
| 10:3.117 | כמקולקלים |  |
| 10:3.118 | שבהם |  |
| 10:3.119 | הנז | ' |
| 10:3.120 | לעיל |  |
| 10:3.121 | עשיתם | ! |
| 10:3.122 | כן |  |
| 10:3.123 | אירע |  |
| 10:3.124 | לנו |  |
| 10:3.125 | באמונת |  |
| 10:3.126 | האחדות |  |
| 10:3.127 | והעבודה | , |
| 10:3.128 | כמתוקנים |  |
| 10:3.129 | שבהם |  |
| 10:3.130 | לא |  |
| 10:3.131 | עשינו |  |
| 10:3.132 | כמקולקלים |  |
| 10:3.133 | שבהם |  |
| 10:3.134 | עשינו |  |
| 10:3.135 | לעבוד |  |
| 10:3.136 | אל |  |
| 10:3.137 | קצא |  |
| 10:3.138 | אפים |  |
| 10:3.139 | ז"א | ) |
| 10:3.140 | וליחדו |  |
| 10:3.141 | עם |  |
| 10:3.142 | האלילה |  |
| 10:3.143 | זוגתו | . |


## Milchamot ha-Shem 10:4

Hebrew (Wikisource): קמב) מי יגלה עפר מעיניכם רבותינו אנשי כנסת הגדולה אשר שפכתם שיח בתחנה ובקשה וצעקה לפני אבינו שבשמים בקול גדול על יצרא דע"ז, להעביר גלולים ובעלים ועשתרות מעם ישראל ועלתה בידכם! אך לא לנצח כי קבלתכם האמתית המשנה והתלמוד הקדוש ופסקי הגאונים הבאים אחריכם שם אותה המתעה מחבר הזהר וההולכים בעקבותיו לשפחה וקליפה ומוץ ותבן וסלע אחרא מסטרא דיצרא בישא, ואת לימוד עבודת הבעלים והעשתרות והאשרות שהם פרצופי וצורות האלוהות הזכרים והנקבות הרבים שם לגבירה, ויעבדום בקדושה ובטהרת הגוף ורתת ופחד, עם טומאת השכל ובלבול המחשבה עד שלא יוכל להבין מה שמוציא מפיו כאשר יראה המעיין בספרי הכוונות הרבים השונים. ועליהם אמר דוד המלך ע"ה (לפי פי' דר"ר עובדיה ספורנו והרב זמרת יה) וז"ל הרב ספורנו במזמור י"ו שמרני אל כי חסיתי בך, ולא בהבלי החושבים לפעול פלאות בשמות מלאכים וזולתם (החמשה פרצופים או הי"ב פרצופים) אמרת לה, את נפשי המשכלת, ה' אתה, אתה אדון לעמך בלי אמצעי טובתי, את נפשי המשכלת, טובה, שלי, בל עליך, אין שום דבר למעלה ממך! לקדושים אשר בארץ המה, ואמרת גם כן לאותם אשר בארץ בחזקת קדושים, ואדירי כל חפצי בם שהם אדירים גם כן. ונחשבים בישראל אשר כל חפצי בם ירבו עצבותם, אמרת שירבו עצבותם וצרותם, כי אמנם בתכליתם. אחר מהרו מיהרו לפנות אל אל אחד מבלעדי אלהים אבל אל מלאכים וזולתם (ירמוז אל פרצופי האצילות) בל אסיך נסכיהם מדם לא אקרב לעשות הנסכים שהם עושים להוריד הכתות, יותר ממה שהייתי מנסך את הדם שהוא מתועב לנסכים ובל אשא את שמותם אותם שמות שהם חושבים לפעול בהם פלאות לא אשא על שפתי ולא אזכירם כי אמנם אתה ה' לבדך מנת חלקי בבחורתי, וכוסי בירושת אבות, אתה תומיך גורלי אתה תתמכני בגורל שעלה לי להיות במספר עבדיך ותסמכני שלא אטה ממנו, כי אמנם חבלים נפלו לי בנעימים, הגיעו לחלקח במושכלות שהם אותם נעימים שהשגתי אף נחלת קבלת אבות שפרה עלי שמצאתיה מסכמת לאמת המושכל ובכן אברך את ה' אשר יעצני ללכת בדרכיו וכו' ועיין זמרת יה שהאריך עוד יותר מזה. גם במזמור כ"ו היה דוד המלך ע"ה, מתחנן ומתפלל לשי"ת לאמר שפטני ה' היודע כליות ולב, כי אני בתמי הלכתי ולא נטיתי מעיקרי האמונה, ובה' לבדו בטחתי בלי שום שתוף כח אחר מכל הנבראים גשמי או רוחני, בחנני ה' ונסיני וצרפה כליותי ולבי שהם כלי המחשבה אם אטה מני הדרך עיקרי הדת, כי חסדך לנגד עיני ששמעת תפלתי בבקשי מלפניך הורני ה' דרכך אהלך באמתיך, ועשית בקשתי ונתת לי נפש משכלת, ועל ידי זה התהלכתי באמתיך ולא ישבתי עם מתי שוא התרים אחר לבם להמציא ברעיונותם סבות אחרות וכחות ופרצופים אלהיים נעלמים מבלעדיך ליחדם ולקשרם ולתקנם להיות אחד אע"פ שהם רבים במספר, ולהיותם אמצעיים למשוך בכחם השפע מהאלהים הגנוז בכתר, ולא מהאלהים העליון המקיף את כל החלל של כל העולמות (לפי שעתם) וזהו שאמר ועם נעלמים הדורשים מכוחות נעלמים מבלעדי ה' לא אבוא, כי שנאתי קהל מרעים המתהללים באלילים ועם אלה רשעים לא אשב. אבל ארחץ בנקיון כפי מכל אמונות זרות וכוזבות ואסובבה את מזבחך ה'.

| id | word | punctuation after |
|---|---|---|
| 10:4.1 | קמב | ) |
| 10:4.2 | מי |  |
| 10:4.3 | יגלה |  |
| 10:4.4 | עפר |  |
| 10:4.5 | מעיניכם |  |
| 10:4.6 | רבותינו |  |
| 10:4.7 | אנשי |  |
| 10:4.8 | כנסת |  |
| 10:4.9 | הגדולה |  |
| 10:4.10 | אשר |  |
| 10:4.11 | שפכתם |  |
| 10:4.12 | שיח |  |
| 10:4.13 | בתחנה |  |
| 10:4.14 | ובקשה |  |
| 10:4.15 | וצעקה |  |
| 10:4.16 | לפני |  |
| 10:4.17 | אבינו |  |
| 10:4.18 | שבשמים |  |
| 10:4.19 | בקול |  |
| 10:4.20 | גדול |  |
| 10:4.21 | על |  |
| 10:4.22 | יצרא |  |
| 10:4.23 | דע"ז | , |
| 10:4.24 | להעביר |  |
| 10:4.25 | גלולים |  |
| 10:4.26 | ובעלים |  |
| 10:4.27 | ועשתרות |  |
| 10:4.28 | מעם |  |
| 10:4.29 | ישראל |  |
| 10:4.30 | ועלתה |  |
| 10:4.31 | בידכם | ! |
| 10:4.32 | אך |  |
| 10:4.33 | לא |  |
| 10:4.34 | לנצח |  |
| 10:4.35 | כי |  |
| 10:4.36 | קבלתכם |  |
| 10:4.37 | האמתית |  |
| 10:4.38 | המשנה |  |
| 10:4.39 | והתלמוד |  |
| 10:4.40 | הקדוש |  |
| 10:4.41 | ופסקי |  |
| 10:4.42 | הגאונים |  |
| 10:4.43 | הבאים |  |
| 10:4.44 | אחריכם |  |
| 10:4.45 | שם |  |
| 10:4.46 | אותה |  |
| 10:4.47 | המתעה |  |
| 10:4.48 | מחבר |  |
| 10:4.49 | הזהר |  |
| 10:4.50 | וההולכים |  |
| 10:4.51 | בעקבותיו |  |
| 10:4.52 | לשפחה |  |
| 10:4.53 | וקליפה |  |
| 10:4.54 | ומוץ |  |
| 10:4.55 | ותבן |  |
| 10:4.56 | וסלע |  |
| 10:4.57 | אחרא |  |
| 10:4.58 | מסטרא |  |
| 10:4.59 | דיצרא |  |
| 10:4.60 | בישא | , |
| 10:4.61 | ואת |  |
| 10:4.62 | לימוד |  |
| 10:4.63 | עבודת |  |
| 10:4.64 | הבעלים |  |
| 10:4.65 | והעשתרות |  |
| 10:4.66 | והאשרות |  |
| 10:4.67 | שהם |  |
| 10:4.68 | פרצופי |  |
| 10:4.69 | וצורות |  |
| 10:4.70 | האלוהות |  |
| 10:4.71 | הזכרים |  |
| 10:4.72 | והנקבות |  |
| 10:4.73 | הרבים |  |
| 10:4.74 | שם |  |
| 10:4.75 | לגבירה | , |
| 10:4.76 | ויעבדום |  |
| 10:4.77 | בקדושה |  |
| 10:4.78 | ובטהרת |  |
| 10:4.79 | הגוף |  |
| 10:4.80 | ורתת |  |
| 10:4.81 | ופחד | , |
| 10:4.82 | עם |  |
| 10:4.83 | טומאת |  |
| 10:4.84 | השכל |  |
| 10:4.85 | ובלבול |  |
| 10:4.86 | המחשבה |  |
| 10:4.87 | עד |  |
| 10:4.88 | שלא |  |
| 10:4.89 | יוכל |  |
| 10:4.90 | להבין |  |
| 10:4.91 | מה |  |
| 10:4.92 | שמוציא |  |
| 10:4.93 | מפיו |  |
| 10:4.94 | כאשר |  |
| 10:4.95 | יראה |  |
| 10:4.96 | המעיין |  |
| 10:4.97 | בספרי |  |
| 10:4.98 | הכוונות |  |
| 10:4.99 | הרבים |  |
| 10:4.100 | השונים | . |
| 10:4.101 | ועליהם |  |
| 10:4.102 | אמר |  |
| 10:4.103 | דוד |  |
| 10:4.104 | המלך |  |
| 10:4.105 | ע"ה |  |
| 10:4.106 | לפי |  |
| 10:4.107 | פי | ' |
| 10:4.108 | דר"ר |  |
| 10:4.109 | עובדיה |  |
| 10:4.110 | ספורנו |  |
| 10:4.111 | והרב |  |
| 10:4.112 | זמרת |  |
| 10:4.113 | יה | ) |
| 10:4.114 | וז"ל |  |
| 10:4.115 | הרב |  |
| 10:4.116 | ספורנו |  |
| 10:4.117 | במזמור |  |
| 10:4.118 | י"ו |  |
| 10:4.119 | שמרני |  |
| 10:4.120 | אל |  |
| 10:4.121 | כי |  |
| 10:4.122 | חסיתי |  |
| 10:4.123 | בך | , |
| 10:4.124 | ולא |  |
| 10:4.125 | בהבלי |  |
| 10:4.126 | החושבים |  |
| 10:4.127 | לפעול |  |
| 10:4.128 | פלאות |  |
| 10:4.129 | בשמות |  |
| 10:4.130 | מלאכים |  |
| 10:4.131 | וזולתם |  |
| 10:4.132 | החמשה |  |
| 10:4.133 | פרצופים |  |
| 10:4.134 | או |  |
| 10:4.135 | הי"ב |  |
| 10:4.136 | פרצופים | ) |
| 10:4.137 | אמרת |  |
| 10:4.138 | לה | , |
| 10:4.139 | את |  |
| 10:4.140 | נפשי |  |
| 10:4.141 | המשכלת | , |
| 10:4.142 | ה | ' |
| 10:4.143 | אתה | , |
| 10:4.144 | אתה |  |
| 10:4.145 | אדון |  |
| 10:4.146 | לעמך |  |
| 10:4.147 | בלי |  |
| 10:4.148 | אמצעי |  |
| 10:4.149 | טובתי | , |
| 10:4.150 | את |  |
| 10:4.151 | נפשי |  |
| 10:4.152 | המשכלת | , |
| 10:4.153 | טובה | , |
| 10:4.154 | שלי | , |
| 10:4.155 | בל |  |
| 10:4.156 | עליך | , |
| 10:4.157 | אין |  |
| 10:4.158 | שום |  |
| 10:4.159 | דבר |  |
| 10:4.160 | למעלה |  |
| 10:4.161 | ממך | ! |
| 10:4.162 | לקדושים |  |
| 10:4.163 | אשר |  |
| 10:4.164 | בארץ |  |
| 10:4.165 | המה | , |
| 10:4.166 | ואמרת |  |
| 10:4.167 | גם |  |
| 10:4.168 | כן |  |
| 10:4.169 | לאותם |  |
| 10:4.170 | אשר |  |
| 10:4.171 | בארץ |  |
| 10:4.172 | בחזקת |  |
| 10:4.173 | קדושים | , |
| 10:4.174 | ואדירי |  |
| 10:4.175 | כל |  |
| 10:4.176 | חפצי |  |
| 10:4.177 | בם |  |
| 10:4.178 | שהם |  |
| 10:4.179 | אדירים |  |
| 10:4.180 | גם |  |
| 10:4.181 | כן | . |
| 10:4.182 | ונחשבים |  |
| 10:4.183 | בישראל |  |
| 10:4.184 | אשר |  |
| 10:4.185 | כל |  |
| 10:4.186 | חפצי |  |
| 10:4.187 | בם |  |
| 10:4.188 | ירבו |  |
| 10:4.189 | עצבותם | , |
| 10:4.190 | אמרת |  |
| 10:4.191 | שירבו |  |
| 10:4.192 | עצבותם |  |
| 10:4.193 | וצרותם | , |
| 10:4.194 | כי |  |
| 10:4.195 | אמנם |  |
| 10:4.196 | בתכליתם | . |
| 10:4.197 | אחר |  |
| 10:4.198 | מהרו |  |
| 10:4.199 | מיהרו |  |
| 10:4.200 | לפנות |  |
| 10:4.201 | אל |  |
| 10:4.202 | אל |  |
| 10:4.203 | אחד |  |
| 10:4.204 | מבלעדי |  |
| 10:4.205 | אלהים |  |
| 10:4.206 | אבל |  |
| 10:4.207 | אל |  |
| 10:4.208 | מלאכים |  |
| 10:4.209 | וזולתם |  |
| 10:4.210 | ירמוז |  |
| 10:4.211 | אל |  |
| 10:4.212 | פרצופי |  |
| 10:4.213 | האצילות | ) |
| 10:4.214 | בל |  |
| 10:4.215 | אסיך |  |
| 10:4.216 | נסכיהם |  |
| 10:4.217 | מדם |  |
| 10:4.218 | לא |  |
| 10:4.219 | אקרב |  |
| 10:4.220 | לעשות |  |
| 10:4.221 | הנסכים |  |
| 10:4.222 | שהם |  |
| 10:4.223 | עושים |  |
| 10:4.224 | להוריד |  |
| 10:4.225 | הכתות | , |
| 10:4.226 | יותר |  |
| 10:4.227 | ממה |  |
| 10:4.228 | שהייתי |  |
| 10:4.229 | מנסך |  |
| 10:4.230 | את |  |
| 10:4.231 | הדם |  |
| 10:4.232 | שהוא |  |
| 10:4.233 | מתועב |  |
| 10:4.234 | לנסכים |  |
| 10:4.235 | ובל |  |
| 10:4.236 | אשא |  |
| 10:4.237 | את |  |
| 10:4.238 | שמותם |  |
| 10:4.239 | אותם |  |
| 10:4.240 | שמות |  |
| 10:4.241 | שהם |  |
| 10:4.242 | חושבים |  |
| 10:4.243 | לפעול |  |
| 10:4.244 | בהם |  |
| 10:4.245 | פלאות |  |
| 10:4.246 | לא |  |
| 10:4.247 | אשא |  |
| 10:4.248 | על |  |
| 10:4.249 | שפתי |  |
| 10:4.250 | ולא |  |
| 10:4.251 | אזכירם |  |
| 10:4.252 | כי |  |
| 10:4.253 | אמנם |  |
| 10:4.254 | אתה |  |
| 10:4.255 | ה | ' |
| 10:4.256 | לבדך |  |
| 10:4.257 | מנת |  |
| 10:4.258 | חלקי |  |
| 10:4.259 | בבחורתי | , |
| 10:4.260 | וכוסי |  |
| 10:4.261 | בירושת |  |
| 10:4.262 | אבות | , |
| 10:4.263 | אתה |  |
| 10:4.264 | תומיך |  |
| 10:4.265 | גורלי |  |
| 10:4.266 | אתה |  |
| 10:4.267 | תתמכני |  |
| 10:4.268 | בגורל |  |
| 10:4.269 | שעלה |  |
| 10:4.270 | לי |  |
| 10:4.271 | להיות |  |
| 10:4.272 | במספר |  |
| 10:4.273 | עבדיך |  |
| 10:4.274 | ותסמכני |  |
| 10:4.275 | שלא |  |
| 10:4.276 | אטה |  |
| 10:4.277 | ממנו | , |
| 10:4.278 | כי |  |
| 10:4.279 | אמנם |  |
| 10:4.280 | חבלים |  |
| 10:4.281 | נפלו |  |
| 10:4.282 | לי |  |
| 10:4.283 | בנעימים | , |
| 10:4.284 | הגיעו |  |
| 10:4.285 | לחלקח |  |
| 10:4.286 | במושכלות |  |
| 10:4.287 | שהם |  |
| 10:4.288 | אותם |  |
| 10:4.289 | נעימים |  |
| 10:4.290 | שהשגתי |  |
| 10:4.291 | אף |  |
| 10:4.292 | נחלת |  |
| 10:4.293 | קבלת |  |
| 10:4.294 | אבות |  |
| 10:4.295 | שפרה |  |
| 10:4.296 | עלי |  |
| 10:4.297 | שמצאתיה |  |
| 10:4.298 | מסכמת |  |
| 10:4.299 | לאמת |  |
| 10:4.300 | המושכל |  |
| 10:4.301 | ובכן |  |
| 10:4.302 | אברך |  |
| 10:4.303 | את |  |
| 10:4.304 | ה | ' |
| 10:4.305 | אשר |  |
| 10:4.306 | יעצני |  |
| 10:4.307 | ללכת |  |
| 10:4.308 | בדרכיו |  |
| 10:4.309 | וכו | ' |
| 10:4.310 | ועיין |  |
| 10:4.311 | זמרת |  |
| 10:4.312 | יה |  |
| 10:4.313 | שהאריך |  |
| 10:4.314 | עוד |  |
| 10:4.315 | יותר |  |
| 10:4.316 | מזה | . |
| 10:4.317 | גם |  |
| 10:4.318 | במזמור |  |
| 10:4.319 | כ"ו |  |
| 10:4.320 | היה |  |
| 10:4.321 | דוד |  |
| 10:4.322 | המלך |  |
| 10:4.323 | ע"ה | , |
| 10:4.324 | מתחנן |  |
| 10:4.325 | ומתפלל |  |
| 10:4.326 | לשי"ת |  |
| 10:4.327 | לאמר |  |
| 10:4.328 | שפטני |  |
| 10:4.329 | ה | ' |
| 10:4.330 | היודע |  |
| 10:4.331 | כליות |  |
| 10:4.332 | ולב | , |
| 10:4.333 | כי |  |
| 10:4.334 | אני |  |
| 10:4.335 | בתמי |  |
| 10:4.336 | הלכתי |  |
| 10:4.337 | ולא |  |
| 10:4.338 | נטיתי |  |
| 10:4.339 | מעיקרי |  |
| 10:4.340 | האמונה | , |
| 10:4.341 | ובה | ' |
| 10:4.342 | לבדו |  |
| 10:4.343 | בטחתי |  |
| 10:4.344 | בלי |  |
| 10:4.345 | שום |  |
| 10:4.346 | שתוף |  |
| 10:4.347 | כח |  |
| 10:4.348 | אחר |  |
| 10:4.349 | מכל |  |
| 10:4.350 | הנבראים |  |
| 10:4.351 | גשמי |  |
| 10:4.352 | או |  |
| 10:4.353 | רוחני | , |
| 10:4.354 | בחנני |  |
| 10:4.355 | ה | ' |
| 10:4.356 | ונסיני |  |
| 10:4.357 | וצרפה |  |
| 10:4.358 | כליותי |  |
| 10:4.359 | ולבי |  |
| 10:4.360 | שהם |  |
| 10:4.361 | כלי |  |
| 10:4.362 | המחשבה |  |
| 10:4.363 | אם |  |
| 10:4.364 | אטה |  |
| 10:4.365 | מני |  |
| 10:4.366 | הדרך |  |
| 10:4.367 | עיקרי |  |
| 10:4.368 | הדת | , |
| 10:4.369 | כי |  |
| 10:4.370 | חסדך |  |
| 10:4.371 | לנגד |  |
| 10:4.372 | עיני |  |
| 10:4.373 | ששמעת |  |
| 10:4.374 | תפלתי |  |
| 10:4.375 | בבקשי |  |
| 10:4.376 | מלפניך |  |
| 10:4.377 | הורני |  |
| 10:4.378 | ה | ' |
| 10:4.379 | דרכך |  |
| 10:4.380 | אהלך |  |
| 10:4.381 | באמתיך | , |
| 10:4.382 | ועשית |  |
| 10:4.383 | בקשתי |  |
| 10:4.384 | ונתת |  |
| 10:4.385 | לי |  |
| 10:4.386 | נפש |  |
| 10:4.387 | משכלת | , |
| 10:4.388 | ועל |  |
| 10:4.389 | ידי |  |
| 10:4.390 | זה |  |
| 10:4.391 | התהלכתי |  |
| 10:4.392 | באמתיך |  |
| 10:4.393 | ולא |  |
| 10:4.394 | ישבתי |  |
| 10:4.395 | עם |  |
| 10:4.396 | מתי |  |
| 10:4.397 | שוא |  |
| 10:4.398 | התרים |  |
| 10:4.399 | אחר |  |
| 10:4.400 | לבם |  |
| 10:4.401 | להמציא |  |
| 10:4.402 | ברעיונותם |  |
| 10:4.403 | סבות |  |
| 10:4.404 | אחרות |  |
| 10:4.405 | וכחות |  |
| 10:4.406 | ופרצופים |  |
| 10:4.407 | אלהיים |  |
| 10:4.408 | נעלמים |  |
| 10:4.409 | מבלעדיך |  |
| 10:4.410 | ליחדם |  |
| 10:4.411 | ולקשרם |  |
| 10:4.412 | ולתקנם |  |
| 10:4.413 | להיות |  |
| 10:4.414 | אחד |  |
| 10:4.415 | אע"פ |  |
| 10:4.416 | שהם |  |
| 10:4.417 | רבים |  |
| 10:4.418 | במספר | , |
| 10:4.419 | ולהיותם |  |
| 10:4.420 | אמצעיים |  |
| 10:4.421 | למשוך |  |
| 10:4.422 | בכחם |  |
| 10:4.423 | השפע |  |
| 10:4.424 | מהאלהים |  |
| 10:4.425 | הגנוז |  |
| 10:4.426 | בכתר | , |
| 10:4.427 | ולא |  |
| 10:4.428 | מהאלהים |  |
| 10:4.429 | העליון |  |
| 10:4.430 | המקיף |  |
| 10:4.431 | את |  |
| 10:4.432 | כל |  |
| 10:4.433 | החלל |  |
| 10:4.434 | של |  |
| 10:4.435 | כל |  |
| 10:4.436 | העולמות |  |
| 10:4.437 | לפי |  |
| 10:4.438 | שעתם | ) |
| 10:4.439 | וזהו |  |
| 10:4.440 | שאמר |  |
| 10:4.441 | ועם |  |
| 10:4.442 | נעלמים |  |
| 10:4.443 | הדורשים |  |
| 10:4.444 | מכוחות |  |
| 10:4.445 | נעלמים |  |
| 10:4.446 | מבלעדי |  |
| 10:4.447 | ה | ' |
| 10:4.448 | לא |  |
| 10:4.449 | אבוא | , |
| 10:4.450 | כי |  |
| 10:4.451 | שנאתי |  |
| 10:4.452 | קהל |  |
| 10:4.453 | מרעים |  |
| 10:4.454 | המתהללים |  |
| 10:4.455 | באלילים |  |
| 10:4.456 | ועם |  |
| 10:4.457 | אלה |  |
| 10:4.458 | רשעים |  |
| 10:4.459 | לא |  |
| 10:4.460 | אשב | . |
| 10:4.461 | אבל |  |
| 10:4.462 | ארחץ |  |
| 10:4.463 | בנקיון |  |
| 10:4.464 | כפי |  |
| 10:4.465 | מכל |  |
| 10:4.466 | אמונות |  |
| 10:4.467 | זרות |  |
| 10:4.468 | וכוזבות |  |
| 10:4.469 | ואסובבה |  |
| 10:4.470 | את |  |
| 10:4.471 | מזבחך |  |
| 10:4.472 | ה | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

