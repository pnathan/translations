# Translation pack: milhamot-hashem perek-7-06 (Milchamot ha-Shem 7:8–7:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-06`.

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
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **הגאון** alone → "the Gaon".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
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
- **בא"י** → "in the Land of Israel".

## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **זווג** → "coupling".
- **מין / מינות** → "heretic / heresy".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
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
- אין לנו **להחמיר** → "we may not be stringent".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **תא חזי** → "come and see".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **מיתות**: פ' ד' מיתות (פ' + a chapter title) → "Four Death Penalties".
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


## Milchamot ha-Shem 7:8

Hebrew (Wikisource): שים לבך אתה הקורא לדברי הגאון מהרש"ל ורמיזותיו שאמר שחלושי הראות לא יביטו. ולא ידעו מוצאו ומובאו של ספר הזהר. כי מיד בן נכר יצא ובא לחכמי טוליטולא כמו שכ' הר' חיד"א בספרו שם הגדולים. ומה שכתב וכוונתו, ר"ל שכל כוונתו להטעות עם ישראל להאמין באלהים רבים. ומה שכ' פקח עיניך וראה אם יש בזה ממש וכו' כלומר אם אמת דבריו שכן אמר רשב"י לא אשתמיט שום רמז בתלמוד בבלי וירושלמי וספרא ספרי. ושאלתות ופסיקתא ממה ששמעו מרשב"י. כי חבר היה לחכמי המשנה ואי אפשר שלא ישמעו ממנו דבר מכל הנאמר בזהר ותיקונים. אלא ודאי הכל שקר וכזב. ומי שלא ידע להשיג תורה על נכון בתלמוד ופוסקים יבוא לקצץ בנטיעות להאמין שאלהים רבים וכמו שמסיים בסוף דבריו והוי כמו מינות. וכמאמר שלמה המלך ע"ה פתי יאמין לכל דבר. וערום יבין לאשורו. עוד כתב רביד הזהב ואע"ג דהש"ך כ' דאין דברי המרש"ל מוכרחים, וגם הדמשק אליעזר הביאם הכנה"ג המה עלוי דמה מינות שייך במה שהאדם מחמיר על עצמו וכו' ע"כ לעינינו.

| id | word | punctuation after |
|---|---|---|
| 7:8.1 | שים |  |
| 7:8.2 | לבך |  |
| 7:8.3 | אתה |  |
| 7:8.4 | הקורא |  |
| 7:8.5 | לדברי |  |
| 7:8.6 | הגאון |  |
| 7:8.7 | מהרש"ל |  |
| 7:8.8 | ורמיזותיו |  |
| 7:8.9 | שאמר |  |
| 7:8.10 | שחלושי |  |
| 7:8.11 | הראות |  |
| 7:8.12 | לא |  |
| 7:8.13 | יביטו | . |
| 7:8.14 | ולא |  |
| 7:8.15 | ידעו |  |
| 7:8.16 | מוצאו |  |
| 7:8.17 | ומובאו |  |
| 7:8.18 | של |  |
| 7:8.19 | ספר |  |
| 7:8.20 | הזהר | . |
| 7:8.21 | כי |  |
| 7:8.22 | מיד |  |
| 7:8.23 | בן |  |
| 7:8.24 | נכר |  |
| 7:8.25 | יצא |  |
| 7:8.26 | ובא |  |
| 7:8.27 | לחכמי |  |
| 7:8.28 | טוליטולא |  |
| 7:8.29 | כמו |  |
| 7:8.30 | שכ | ' |
| 7:8.31 | הר | ' |
| 7:8.32 | חיד"א |  |
| 7:8.33 | בספרו |  |
| 7:8.34 | שם |  |
| 7:8.35 | הגדולים | . |
| 7:8.36 | ומה |  |
| 7:8.37 | שכתב |  |
| 7:8.38 | וכוונתו | , |
| 7:8.39 | ר"ל |  |
| 7:8.40 | שכל |  |
| 7:8.41 | כוונתו |  |
| 7:8.42 | להטעות |  |
| 7:8.43 | עם |  |
| 7:8.44 | ישראל |  |
| 7:8.45 | להאמין |  |
| 7:8.46 | באלהים |  |
| 7:8.47 | רבים | . |
| 7:8.48 | ומה |  |
| 7:8.49 | שכ | ' |
| 7:8.50 | פקח |  |
| 7:8.51 | עיניך |  |
| 7:8.52 | וראה |  |
| 7:8.53 | אם |  |
| 7:8.54 | יש |  |
| 7:8.55 | בזה |  |
| 7:8.56 | ממש |  |
| 7:8.57 | וכו | ' |
| 7:8.58 | כלומר |  |
| 7:8.59 | אם |  |
| 7:8.60 | אמת |  |
| 7:8.61 | דבריו |  |
| 7:8.62 | שכן |  |
| 7:8.63 | אמר |  |
| 7:8.64 | רשב"י |  |
| 7:8.65 | לא |  |
| 7:8.66 | אשתמיט |  |
| 7:8.67 | שום |  |
| 7:8.68 | רמז |  |
| 7:8.69 | בתלמוד |  |
| 7:8.70 | בבלי |  |
| 7:8.71 | וירושלמי |  |
| 7:8.72 | וספרא |  |
| 7:8.73 | ספרי | . |
| 7:8.74 | ושאלתות |  |
| 7:8.75 | ופסיקתא |  |
| 7:8.76 | ממה |  |
| 7:8.77 | ששמעו |  |
| 7:8.78 | מרשב"י | . |
| 7:8.79 | כי |  |
| 7:8.80 | חבר |  |
| 7:8.81 | היה |  |
| 7:8.82 | לחכמי |  |
| 7:8.83 | המשנה |  |
| 7:8.84 | ואי |  |
| 7:8.85 | אפשר |  |
| 7:8.86 | שלא |  |
| 7:8.87 | ישמעו |  |
| 7:8.88 | ממנו |  |
| 7:8.89 | דבר |  |
| 7:8.90 | מכל |  |
| 7:8.91 | הנאמר |  |
| 7:8.92 | בזהר |  |
| 7:8.93 | ותיקונים | . |
| 7:8.94 | אלא |  |
| 7:8.95 | ודאי |  |
| 7:8.96 | הכל |  |
| 7:8.97 | שקר |  |
| 7:8.98 | וכזב | . |
| 7:8.99 | ומי |  |
| 7:8.100 | שלא |  |
| 7:8.101 | ידע |  |
| 7:8.102 | להשיג |  |
| 7:8.103 | תורה |  |
| 7:8.104 | על |  |
| 7:8.105 | נכון |  |
| 7:8.106 | בתלמוד |  |
| 7:8.107 | ופוסקים |  |
| 7:8.108 | יבוא |  |
| 7:8.109 | לקצץ |  |
| 7:8.110 | בנטיעות |  |
| 7:8.111 | להאמין |  |
| 7:8.112 | שאלהים |  |
| 7:8.113 | רבים |  |
| 7:8.114 | וכמו |  |
| 7:8.115 | שמסיים |  |
| 7:8.116 | בסוף |  |
| 7:8.117 | דבריו |  |
| 7:8.118 | והוי |  |
| 7:8.119 | כמו |  |
| 7:8.120 | מינות | . |
| 7:8.121 | וכמאמר |  |
| 7:8.122 | שלמה |  |
| 7:8.123 | המלך |  |
| 7:8.124 | ע"ה |  |
| 7:8.125 | פתי |  |
| 7:8.126 | יאמין |  |
| 7:8.127 | לכל |  |
| 7:8.128 | דבר | . |
| 7:8.129 | וערום |  |
| 7:8.130 | יבין |  |
| 7:8.131 | לאשורו | . |
| 7:8.132 | עוד |  |
| 7:8.133 | כתב |  |
| 7:8.134 | רביד |  |
| 7:8.135 | הזהב |  |
| 7:8.136 | ואע"ג |  |
| 7:8.137 | דהש"ך |  |
| 7:8.138 | כ | ' |
| 7:8.139 | דאין |  |
| 7:8.140 | דברי |  |
| 7:8.141 | המרש"ל |  |
| 7:8.142 | מוכרחים | , |
| 7:8.143 | וגם |  |
| 7:8.144 | הדמשק |  |
| 7:8.145 | אליעזר |  |
| 7:8.146 | הביאם |  |
| 7:8.147 | הכנה"ג |  |
| 7:8.148 | המה |  |
| 7:8.149 | עלוי |  |
| 7:8.150 | דמה |  |
| 7:8.151 | מינות |  |
| 7:8.152 | שייך |  |
| 7:8.153 | במה |  |
| 7:8.154 | שהאדם |  |
| 7:8.155 | מחמיר |  |
| 7:8.156 | על |  |
| 7:8.157 | עצמו |  |
| 7:8.158 | וכו | ' |
| 7:8.159 | ע"כ |  |
| 7:8.160 | לעינינו | . |


## Milchamot ha-Shem 7:9

Hebrew (Wikisource): הנה מה מאד יתפלא הקורא הנכון על דברי הש"ך שכתב על דברי הגאון מהרש"ל שאין דבריו מוכרחים. והוא בעצמו לא ירד לסוף דעתו, וכל שכן על תמיהת הדמשק אליעזר איך מלאו לבו לתמוה על אדם גדול כמהרש"ל אשר שט בים התלמוד וספרא וספרי ושאלתות ועיניו ראו ועפעפיו בחנו, כי דרך עצב בדבר הזהר והמקובלים החדשים, והוא לבדו ידע מהיכן מוצאו ומובאו, כי מיד עכו"ם יצא ובא אלינו, להסית ולהדיח אותנו להאמין באלהים רבים ולעבוד אלהים אחרים. סבות רבות מחודשות נבראות המה הצורות והפרצופים שהמציא מחבר הזהר ברעיונותיו, וזהו שכ' רש"ל שלא ידעו מוצאו ומובאו וכוונתו. והוסיף עוד לומר פקח עיניך וראה אם יש ממש בדבריו וממקור נאמן יצאו דברים כאלה רבוי באלהות בפרצופים שונים ומלכות וממשלה להם, לא לשתמיט שום רמז מזה בתלמוד בבלי וירושלמי וספרא וספרי ופסיקתא ושאלתות, אלא ודאי לא מישראל יצאו הדברים האלה וחיצונים הם, ולכך אמר לכן אהובי אל תלך בדרך אתם וכו', ואף מי שדעתו להחמיר על עצמו, הואיל ואין לזאת החומרא שום רמז בתלמוד והוא מקפיד בדבר שלא הקפידו בו חכמים הוי כמו מינות כדתנן בצבועים איני עובר אף בלבנים לא יעבור וכמ"ש התוספת י"ט ויובא לקמן סי כ"ה. אכן עוד תגדל הפליאה על הדמשק אליעזר באמרו מה מינות שייך במה שאדם מחמיר על עצמו, כי ידוע הוא ומפורסם בכל האיסורים והחומרות והמנהגים שחידשו הזהר והמקובלים החדשים הבאים אחריו, כולהו אית בהו טעמא לפי שטתם באמונת אלוהות רבים, והרש"ל ז"ל הרגיש בזרותם כי עיניו יחזו ועפעפיו יבחנו, כי בילדי נכרים עובדי ע"ז ישפיקו, ודרך עוצב יש בטעמי האיסורים והחומרות והמנהגים שחידשו שלא כדרך הגמרא ואית בהו לתא דע"ז. ולפי שהראשונים כר' תם ן' יחיא ומהר"י אלבו והר' מהרש"ל וכיוצא בהם יראו לנפשם מאד פן יפגעו להם הרודפים הקנאים כאיזבל לאליהו, לכן תלו הטעות בלומדים והזהירום ממנה, ולפי האמת הברור כי אותן החומרות מינות גמורה. והא לך דוגמא מה שכתב הרב משנת חסידים מסכת ליל פסח פרק י"א שכשאומר בהגדה דם ואש ותימרות עשן וכן בהזכירו עשר מכות בפרט ובכלל ישפוך מהכוס לתוך כלי שבור הרומז אל הקליפה הנקרא ארור, ועל ידי זה תפסוק המלכות (שהיא דשתנא) מדם נדתה ותהיה ראויה לזווג ונקתה ונזרעה זרע, כי היין הנשאר בה הוא יין המשמח דם טהור המתהפך לחלב אחר כך ע"כ. והחכם עיניו בראשו להשכיל, דכל כי האי גונא מנסך הוא לע"ז שמנסך מכוסו לקליפה הנקראת ארור כדי לתקן המלכות ויינו אסור בהנאה מן התורה שהרי ניסך ממנו לקליפה הנקראת ארור כדי לתקן, והמנסך היא אחת מארבע עבודות שחייב עליהן בין שדרך עבודתה בכך, בין שאין דרך עבודתה בכך כדתנן בסנהדרין פ' ארבע מיתות, (סנהדרין דף ס' ע"ב). וכן מה שמקפידין לחלוץ התפלין קודם מוסף ראש חדש טעמא הוא לפי שהתפילין הם בז"א שלפי דבריהם עליו אמרו רז"ל במס' ברכות שהקב"ה מניח תפילין, ובמוסף ר"ח הוא בא אור גדול מעתיק שהוא האלוה העליון שאינו מניח תפלין, כי התפלין הם שייכים לקצר אפים שהוא זעיר אנפין, וכי האיי גונא מינות הוא להאמין ברשויות הרבה, זה מניח תפילין וזה אין מניח תפילין, והרי אין הקב"ה אחד.

| id | word | punctuation after |
|---|---|---|
| 7:9.1 | הנה |  |
| 7:9.2 | מה |  |
| 7:9.3 | מאד |  |
| 7:9.4 | יתפלא |  |
| 7:9.5 | הקורא |  |
| 7:9.6 | הנכון |  |
| 7:9.7 | על |  |
| 7:9.8 | דברי |  |
| 7:9.9 | הש"ך |  |
| 7:9.10 | שכתב |  |
| 7:9.11 | על |  |
| 7:9.12 | דברי |  |
| 7:9.13 | הגאון |  |
| 7:9.14 | מהרש"ל |  |
| 7:9.15 | שאין |  |
| 7:9.16 | דבריו |  |
| 7:9.17 | מוכרחים | . |
| 7:9.18 | והוא |  |
| 7:9.19 | בעצמו |  |
| 7:9.20 | לא |  |
| 7:9.21 | ירד |  |
| 7:9.22 | לסוף |  |
| 7:9.23 | דעתו | , |
| 7:9.24 | וכל |  |
| 7:9.25 | שכן |  |
| 7:9.26 | על |  |
| 7:9.27 | תמיהת |  |
| 7:9.28 | הדמשק |  |
| 7:9.29 | אליעזר |  |
| 7:9.30 | איך |  |
| 7:9.31 | מלאו |  |
| 7:9.32 | לבו |  |
| 7:9.33 | לתמוה |  |
| 7:9.34 | על |  |
| 7:9.35 | אדם |  |
| 7:9.36 | גדול |  |
| 7:9.37 | כמהרש"ל |  |
| 7:9.38 | אשר |  |
| 7:9.39 | שט |  |
| 7:9.40 | בים |  |
| 7:9.41 | התלמוד |  |
| 7:9.42 | וספרא |  |
| 7:9.43 | וספרי |  |
| 7:9.44 | ושאלתות |  |
| 7:9.45 | ועיניו |  |
| 7:9.46 | ראו |  |
| 7:9.47 | ועפעפיו |  |
| 7:9.48 | בחנו | , |
| 7:9.49 | כי |  |
| 7:9.50 | דרך |  |
| 7:9.51 | עצב |  |
| 7:9.52 | בדבר |  |
| 7:9.53 | הזהר |  |
| 7:9.54 | והמקובלים |  |
| 7:9.55 | החדשים | , |
| 7:9.56 | והוא |  |
| 7:9.57 | לבדו |  |
| 7:9.58 | ידע |  |
| 7:9.59 | מהיכן |  |
| 7:9.60 | מוצאו |  |
| 7:9.61 | ומובאו | , |
| 7:9.62 | כי |  |
| 7:9.63 | מיד |  |
| 7:9.64 | עכו"ם |  |
| 7:9.65 | יצא |  |
| 7:9.66 | ובא |  |
| 7:9.67 | אלינו | , |
| 7:9.68 | להסית |  |
| 7:9.69 | ולהדיח |  |
| 7:9.70 | אותנו |  |
| 7:9.71 | להאמין |  |
| 7:9.72 | באלהים |  |
| 7:9.73 | רבים |  |
| 7:9.74 | ולעבוד |  |
| 7:9.75 | אלהים |  |
| 7:9.76 | אחרים | . |
| 7:9.77 | סבות |  |
| 7:9.78 | רבות |  |
| 7:9.79 | מחודשות |  |
| 7:9.80 | נבראות |  |
| 7:9.81 | המה |  |
| 7:9.82 | הצורות |  |
| 7:9.83 | והפרצופים |  |
| 7:9.84 | שהמציא |  |
| 7:9.85 | מחבר |  |
| 7:9.86 | הזהר |  |
| 7:9.87 | ברעיונותיו | , |
| 7:9.88 | וזהו |  |
| 7:9.89 | שכ | ' |
| 7:9.90 | רש"ל |  |
| 7:9.91 | שלא |  |
| 7:9.92 | ידעו |  |
| 7:9.93 | מוצאו |  |
| 7:9.94 | ומובאו |  |
| 7:9.95 | וכוונתו | . |
| 7:9.96 | והוסיף |  |
| 7:9.97 | עוד |  |
| 7:9.98 | לומר |  |
| 7:9.99 | פקח |  |
| 7:9.100 | עיניך |  |
| 7:9.101 | וראה |  |
| 7:9.102 | אם |  |
| 7:9.103 | יש |  |
| 7:9.104 | ממש |  |
| 7:9.105 | בדבריו |  |
| 7:9.106 | וממקור |  |
| 7:9.107 | נאמן |  |
| 7:9.108 | יצאו |  |
| 7:9.109 | דברים |  |
| 7:9.110 | כאלה |  |
| 7:9.111 | רבוי |  |
| 7:9.112 | באלהות |  |
| 7:9.113 | בפרצופים |  |
| 7:9.114 | שונים |  |
| 7:9.115 | ומלכות |  |
| 7:9.116 | וממשלה |  |
| 7:9.117 | להם | , |
| 7:9.118 | לא |  |
| 7:9.119 | לשתמיט |  |
| 7:9.120 | שום |  |
| 7:9.121 | רמז |  |
| 7:9.122 | מזה |  |
| 7:9.123 | בתלמוד |  |
| 7:9.124 | בבלי |  |
| 7:9.125 | וירושלמי |  |
| 7:9.126 | וספרא |  |
| 7:9.127 | וספרי |  |
| 7:9.128 | ופסיקתא |  |
| 7:9.129 | ושאלתות | , |
| 7:9.130 | אלא |  |
| 7:9.131 | ודאי |  |
| 7:9.132 | לא |  |
| 7:9.133 | מישראל |  |
| 7:9.134 | יצאו |  |
| 7:9.135 | הדברים |  |
| 7:9.136 | האלה |  |
| 7:9.137 | וחיצונים |  |
| 7:9.138 | הם | , |
| 7:9.139 | ולכך |  |
| 7:9.140 | אמר |  |
| 7:9.141 | לכן |  |
| 7:9.142 | אהובי |  |
| 7:9.143 | אל |  |
| 7:9.144 | תלך |  |
| 7:9.145 | בדרך |  |
| 7:9.146 | אתם |  |
| 7:9.147 | וכו | ', |
| 7:9.148 | ואף |  |
| 7:9.149 | מי |  |
| 7:9.150 | שדעתו |  |
| 7:9.151 | להחמיר |  |
| 7:9.152 | על |  |
| 7:9.153 | עצמו | , |
| 7:9.154 | הואיל |  |
| 7:9.155 | ואין |  |
| 7:9.156 | לזאת |  |
| 7:9.157 | החומרא |  |
| 7:9.158 | שום |  |
| 7:9.159 | רמז |  |
| 7:9.160 | בתלמוד |  |
| 7:9.161 | והוא |  |
| 7:9.162 | מקפיד |  |
| 7:9.163 | בדבר |  |
| 7:9.164 | שלא |  |
| 7:9.165 | הקפידו |  |
| 7:9.166 | בו |  |
| 7:9.167 | חכמים |  |
| 7:9.168 | הוי |  |
| 7:9.169 | כמו |  |
| 7:9.170 | מינות |  |
| 7:9.171 | כדתנן |  |
| 7:9.172 | בצבועים |  |
| 7:9.173 | איני |  |
| 7:9.174 | עובר |  |
| 7:9.175 | אף |  |
| 7:9.176 | בלבנים |  |
| 7:9.177 | לא |  |
| 7:9.178 | יעבור |  |
| 7:9.179 | וכמ"ש |  |
| 7:9.180 | התוספת |  |
| 7:9.181 | י"ט |  |
| 7:9.182 | ויובא |  |
| 7:9.183 | לקמן |  |
| 7:9.184 | סי |  |
| 7:9.185 | כ"ה | . |
| 7:9.186 | אכן |  |
| 7:9.187 | עוד |  |
| 7:9.188 | תגדל |  |
| 7:9.189 | הפליאה |  |
| 7:9.190 | על |  |
| 7:9.191 | הדמשק |  |
| 7:9.192 | אליעזר |  |
| 7:9.193 | באמרו |  |
| 7:9.194 | מה |  |
| 7:9.195 | מינות |  |
| 7:9.196 | שייך |  |
| 7:9.197 | במה |  |
| 7:9.198 | שאדם |  |
| 7:9.199 | מחמיר |  |
| 7:9.200 | על |  |
| 7:9.201 | עצמו | , |
| 7:9.202 | כי |  |
| 7:9.203 | ידוע |  |
| 7:9.204 | הוא |  |
| 7:9.205 | ומפורסם |  |
| 7:9.206 | בכל |  |
| 7:9.207 | האיסורים |  |
| 7:9.208 | והחומרות |  |
| 7:9.209 | והמנהגים |  |
| 7:9.210 | שחידשו |  |
| 7:9.211 | הזהר |  |
| 7:9.212 | והמקובלים |  |
| 7:9.213 | החדשים |  |
| 7:9.214 | הבאים |  |
| 7:9.215 | אחריו | , |
| 7:9.216 | כולהו |  |
| 7:9.217 | אית |  |
| 7:9.218 | בהו |  |
| 7:9.219 | טעמא |  |
| 7:9.220 | לפי |  |
| 7:9.221 | שטתם |  |
| 7:9.222 | באמונת |  |
| 7:9.223 | אלוהות |  |
| 7:9.224 | רבים | , |
| 7:9.225 | והרש"ל |  |
| 7:9.226 | ז"ל |  |
| 7:9.227 | הרגיש |  |
| 7:9.228 | בזרותם |  |
| 7:9.229 | כי |  |
| 7:9.230 | עיניו |  |
| 7:9.231 | יחזו |  |
| 7:9.232 | ועפעפיו |  |
| 7:9.233 | יבחנו | , |
| 7:9.234 | כי |  |
| 7:9.235 | בילדי |  |
| 7:9.236 | נכרים |  |
| 7:9.237 | עובדי |  |
| 7:9.238 | ע"ז |  |
| 7:9.239 | ישפיקו | , |
| 7:9.240 | ודרך |  |
| 7:9.241 | עוצב |  |
| 7:9.242 | יש |  |
| 7:9.243 | בטעמי |  |
| 7:9.244 | האיסורים |  |
| 7:9.245 | והחומרות |  |
| 7:9.246 | והמנהגים |  |
| 7:9.247 | שחידשו |  |
| 7:9.248 | שלא |  |
| 7:9.249 | כדרך |  |
| 7:9.250 | הגמרא |  |
| 7:9.251 | ואית |  |
| 7:9.252 | בהו |  |
| 7:9.253 | לתא |  |
| 7:9.254 | דע"ז | . |
| 7:9.255 | ולפי |  |
| 7:9.256 | שהראשונים |  |
| 7:9.257 | כר | ' |
| 7:9.258 | תם |  |
| 7:9.259 | ן | ' |
| 7:9.260 | יחיא |  |
| 7:9.261 | ומהר"י |  |
| 7:9.262 | אלבו |  |
| 7:9.263 | והר | ' |
| 7:9.264 | מהרש"ל |  |
| 7:9.265 | וכיוצא |  |
| 7:9.266 | בהם |  |
| 7:9.267 | יראו |  |
| 7:9.268 | לנפשם |  |
| 7:9.269 | מאד |  |
| 7:9.270 | פן |  |
| 7:9.271 | יפגעו |  |
| 7:9.272 | להם |  |
| 7:9.273 | הרודפים |  |
| 7:9.274 | הקנאים |  |
| 7:9.275 | כאיזבל |  |
| 7:9.276 | לאליהו | , |
| 7:9.277 | לכן |  |
| 7:9.278 | תלו |  |
| 7:9.279 | הטעות |  |
| 7:9.280 | בלומדים |  |
| 7:9.281 | והזהירום |  |
| 7:9.282 | ממנה | , |
| 7:9.283 | ולפי |  |
| 7:9.284 | האמת |  |
| 7:9.285 | הברור |  |
| 7:9.286 | כי |  |
| 7:9.287 | אותן |  |
| 7:9.288 | החומרות |  |
| 7:9.289 | מינות |  |
| 7:9.290 | גמורה | . |
| 7:9.291 | והא |  |
| 7:9.292 | לך |  |
| 7:9.293 | דוגמא |  |
| 7:9.294 | מה |  |
| 7:9.295 | שכתב |  |
| 7:9.296 | הרב |  |
| 7:9.297 | משנת |  |
| 7:9.298 | חסידים |  |
| 7:9.299 | מסכת |  |
| 7:9.300 | ליל |  |
| 7:9.301 | פסח |  |
| 7:9.302 | פרק |  |
| 7:9.303 | י"א |  |
| 7:9.304 | שכשאומר |  |
| 7:9.305 | בהגדה |  |
| 7:9.306 | דם |  |
| 7:9.307 | ואש |  |
| 7:9.308 | ותימרות |  |
| 7:9.309 | עשן |  |
| 7:9.310 | וכן |  |
| 7:9.311 | בהזכירו |  |
| 7:9.312 | עשר |  |
| 7:9.313 | מכות |  |
| 7:9.314 | בפרט |  |
| 7:9.315 | ובכלל |  |
| 7:9.316 | ישפוך |  |
| 7:9.317 | מהכוס |  |
| 7:9.318 | לתוך |  |
| 7:9.319 | כלי |  |
| 7:9.320 | שבור |  |
| 7:9.321 | הרומז |  |
| 7:9.322 | אל |  |
| 7:9.323 | הקליפה |  |
| 7:9.324 | הנקרא |  |
| 7:9.325 | ארור | , |
| 7:9.326 | ועל |  |
| 7:9.327 | ידי |  |
| 7:9.328 | זה |  |
| 7:9.329 | תפסוק |  |
| 7:9.330 | המלכות |  |
| 7:9.331 | שהיא |  |
| 7:9.332 | דשתנא | ) |
| 7:9.333 | מדם |  |
| 7:9.334 | נדתה |  |
| 7:9.335 | ותהיה |  |
| 7:9.336 | ראויה |  |
| 7:9.337 | לזווג |  |
| 7:9.338 | ונקתה |  |
| 7:9.339 | ונזרעה |  |
| 7:9.340 | זרע | , |
| 7:9.341 | כי |  |
| 7:9.342 | היין |  |
| 7:9.343 | הנשאר |  |
| 7:9.344 | בה |  |
| 7:9.345 | הוא |  |
| 7:9.346 | יין |  |
| 7:9.347 | המשמח |  |
| 7:9.348 | דם |  |
| 7:9.349 | טהור |  |
| 7:9.350 | המתהפך |  |
| 7:9.351 | לחלב |  |
| 7:9.352 | אחר |  |
| 7:9.353 | כך |  |
| 7:9.354 | ע"כ | . |
| 7:9.355 | והחכם |  |
| 7:9.356 | עיניו |  |
| 7:9.357 | בראשו |  |
| 7:9.358 | להשכיל | , |
| 7:9.359 | דכל |  |
| 7:9.360 | כי |  |
| 7:9.361 | האי |  |
| 7:9.362 | גונא |  |
| 7:9.363 | מנסך |  |
| 7:9.364 | הוא |  |
| 7:9.365 | לע"ז |  |
| 7:9.366 | שמנסך |  |
| 7:9.367 | מכוסו |  |
| 7:9.368 | לקליפה |  |
| 7:9.369 | הנקראת |  |
| 7:9.370 | ארור |  |
| 7:9.371 | כדי |  |
| 7:9.372 | לתקן |  |
| 7:9.373 | המלכות |  |
| 7:9.374 | ויינו |  |
| 7:9.375 | אסור |  |
| 7:9.376 | בהנאה |  |
| 7:9.377 | מן |  |
| 7:9.378 | התורה |  |
| 7:9.379 | שהרי |  |
| 7:9.380 | ניסך |  |
| 7:9.381 | ממנו |  |
| 7:9.382 | לקליפה |  |
| 7:9.383 | הנקראת |  |
| 7:9.384 | ארור |  |
| 7:9.385 | כדי |  |
| 7:9.386 | לתקן | , |
| 7:9.387 | והמנסך |  |
| 7:9.388 | היא |  |
| 7:9.389 | אחת |  |
| 7:9.390 | מארבע |  |
| 7:9.391 | עבודות |  |
| 7:9.392 | שחייב |  |
| 7:9.393 | עליהן |  |
| 7:9.394 | בין |  |
| 7:9.395 | שדרך |  |
| 7:9.396 | עבודתה |  |
| 7:9.397 | בכך | , |
| 7:9.398 | בין |  |
| 7:9.399 | שאין |  |
| 7:9.400 | דרך |  |
| 7:9.401 | עבודתה |  |
| 7:9.402 | בכך |  |
| 7:9.403 | כדתנן |  |
| 7:9.404 | בסנהדרין |  |
| 7:9.405 | פ | ' |
| 7:9.406 | ארבע |  |
| 7:9.407 | מיתות | , |
| 7:9.408 | סנהדרין |  |
| 7:9.409 | דף |  |
| 7:9.410 | ס | ' |
| 7:9.411 | ע"ב | ). |
| 7:9.412 | וכן |  |
| 7:9.413 | מה |  |
| 7:9.414 | שמקפידין |  |
| 7:9.415 | לחלוץ |  |
| 7:9.416 | התפלין |  |
| 7:9.417 | קודם |  |
| 7:9.418 | מוסף |  |
| 7:9.419 | ראש |  |
| 7:9.420 | חדש |  |
| 7:9.421 | טעמא |  |
| 7:9.422 | הוא |  |
| 7:9.423 | לפי |  |
| 7:9.424 | שהתפילין |  |
| 7:9.425 | הם |  |
| 7:9.426 | בז"א |  |
| 7:9.427 | שלפי |  |
| 7:9.428 | דבריהם |  |
| 7:9.429 | עליו |  |
| 7:9.430 | אמרו |  |
| 7:9.431 | רז"ל |  |
| 7:9.432 | במס | ' |
| 7:9.433 | ברכות |  |
| 7:9.434 | שהקב"ה |  |
| 7:9.435 | מניח |  |
| 7:9.436 | תפילין | , |
| 7:9.437 | ובמוסף |  |
| 7:9.438 | ר"ח |  |
| 7:9.439 | הוא |  |
| 7:9.440 | בא |  |
| 7:9.441 | אור |  |
| 7:9.442 | גדול |  |
| 7:9.443 | מעתיק |  |
| 7:9.444 | שהוא |  |
| 7:9.445 | האלוה |  |
| 7:9.446 | העליון |  |
| 7:9.447 | שאינו |  |
| 7:9.448 | מניח |  |
| 7:9.449 | תפלין | , |
| 7:9.450 | כי |  |
| 7:9.451 | התפלין |  |
| 7:9.452 | הם |  |
| 7:9.453 | שייכים |  |
| 7:9.454 | לקצר |  |
| 7:9.455 | אפים |  |
| 7:9.456 | שהוא |  |
| 7:9.457 | זעיר |  |
| 7:9.458 | אנפין | , |
| 7:9.459 | וכי |  |
| 7:9.460 | האיי |  |
| 7:9.461 | גונא |  |
| 7:9.462 | מינות |  |
| 7:9.463 | הוא |  |
| 7:9.464 | להאמין |  |
| 7:9.465 | ברשויות |  |
| 7:9.466 | הרבה | , |
| 7:9.467 | זה |  |
| 7:9.468 | מניח |  |
| 7:9.469 | תפילין |  |
| 7:9.470 | וזה |  |
| 7:9.471 | אין |  |
| 7:9.472 | מניח |  |
| 7:9.473 | תפילין | , |
| 7:9.474 | והרי |  |
| 7:9.475 | אין |  |
| 7:9.476 | הקב"ה |  |
| 7:9.477 | אחד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

