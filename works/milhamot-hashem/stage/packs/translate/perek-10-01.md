# Translation pack: milhamot-hashem perek-10-01 (Milchamot ha-Shem 10:1–10:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-01`.

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
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **תוה"ק** → "the holy Torah".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
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
- **תא חזי** → "come and see".
- **אבותיהם** for "אבתיכם": follow his wording; the verse in `apparatus`.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **אצהאר / גילוי**: ספר אצהאר אלחק = גילוי האמת; the title as each place writes
  it ("Iẓhar al-Chaqq", "Gilui ha-Emet"), a note that they are one book; its
  author unnamed.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 10:1

Hebrew (Wikisource): קמ) לפי האמור עד כה צריך אדם ליזהר מכל מנהגים וקולות וחומרות, ושינוי נוסחאות שחידשו האחרונים על פי הזהר והמקובלים הנמשכים אחריו, כי רובם שהוא ככולם כנוים על פי טעמים המביאים לידי מינות ורבוי באלהות כמו שנתבאר לעיל סי' פ"ג וכבדברי הרש"ל ז"ל, והן רבים עמי הארץ המחזיקים במנהגים המשונים על ידי רבנים שקדמינו ששינו מנהגי אבותיהם המיוסדים על טהרת תוה"ק שבכתב ושבעל פה, והחליפום על פי ספרי הקבלה החדשה שאין לה בית אב, ובשקר גמור יחסוה להתנא הקדוש רשב"י ע"ה, וחס ליה לפה קדוש רשב"י שיצאו מפיו דברים שכאלה ולומר שנגלה לו עתיקא קדישא והרשהו לחפות דברים אשר לא כן על ה' אלהינו מבלי לחוס על כבוד קונו, וליחס בו שינוי ורבוי איברים וידים ורגלים וגילוי מבושים בזכרים ונקבות וזיווגיהם, חלילה חלילה להאמין כן על התנא הקדוש הלזה. ועל שאר תנאים ואמוראים לחשוב אף בהרהור שום שינוי ורבוי כאל ית'. וכ"ש זכרים ונקבות שונים זה מזה, יש זכר ונקבה בגוף אחד. ויש זכר ונקבה בפרצופים חלוקים איש ואשתו וגיד ערוה, וביצי זכר לפרצופים הזכרים ומקור ועיבור ולידה והנקה לנקבות מדדי אדם וגם מדדי בהמה, לא כאלה חלק יעקב להמיר את כבודם האל המושיעם עושה גדולות במסרים וכו', כי אחת דבר אלהים שתים זו שמענו אנכי ולא יהיה לך. ואותו לבדו נעבוד ונירא מהדר גאונו ולא נמירהו בשום נברא או נאצל, יהיה מי שיהיה כמאמר הכתוב אל תמר בו, ופי' רז"ל אל תמירני בו ואל תמירהו בי (רבה פרשת משפטים). ובסנהדרין פ' א' דיני ממונות אמר ליה ההוא מינא לרב אידית כתיב ואל משה אמר עלה אל ה', עלה אלי מבעי ליה אמר ליה זהו מטטרו"ן ששמו כשם רבו דכ' כי שמי בקרבו, אמר אי הכי נפלחו ליה אמר ליה כתוב אל תמר בו, אל תמירי בו וכו' ועיין רש"י ומהרש"א.

| id | word | punctuation after |
|---|---|---|
| 10:1.1 | קמ | ) |
| 10:1.2 | לפי |  |
| 10:1.3 | האמור |  |
| 10:1.4 | עד |  |
| 10:1.5 | כה |  |
| 10:1.6 | צריך |  |
| 10:1.7 | אדם |  |
| 10:1.8 | ליזהר |  |
| 10:1.9 | מכל |  |
| 10:1.10 | מנהגים |  |
| 10:1.11 | וקולות |  |
| 10:1.12 | וחומרות | , |
| 10:1.13 | ושינוי |  |
| 10:1.14 | נוסחאות |  |
| 10:1.15 | שחידשו |  |
| 10:1.16 | האחרונים |  |
| 10:1.17 | על |  |
| 10:1.18 | פי |  |
| 10:1.19 | הזהר |  |
| 10:1.20 | והמקובלים |  |
| 10:1.21 | הנמשכים |  |
| 10:1.22 | אחריו | , |
| 10:1.23 | כי |  |
| 10:1.24 | רובם |  |
| 10:1.25 | שהוא |  |
| 10:1.26 | ככולם |  |
| 10:1.27 | כנוים |  |
| 10:1.28 | על |  |
| 10:1.29 | פי |  |
| 10:1.30 | טעמים |  |
| 10:1.31 | המביאים |  |
| 10:1.32 | לידי |  |
| 10:1.33 | מינות |  |
| 10:1.34 | ורבוי |  |
| 10:1.35 | באלהות |  |
| 10:1.36 | כמו |  |
| 10:1.37 | שנתבאר |  |
| 10:1.38 | לעיל |  |
| 10:1.39 | סי | ' |
| 10:1.40 | פ"ג |  |
| 10:1.41 | וכבדברי |  |
| 10:1.42 | הרש"ל |  |
| 10:1.43 | ז"ל | , |
| 10:1.44 | והן |  |
| 10:1.45 | רבים |  |
| 10:1.46 | עמי |  |
| 10:1.47 | הארץ |  |
| 10:1.48 | המחזיקים |  |
| 10:1.49 | במנהגים |  |
| 10:1.50 | המשונים |  |
| 10:1.51 | על |  |
| 10:1.52 | ידי |  |
| 10:1.53 | רבנים |  |
| 10:1.54 | שקדמינו |  |
| 10:1.55 | ששינו |  |
| 10:1.56 | מנהגי |  |
| 10:1.57 | אבותיהם |  |
| 10:1.58 | המיוסדים |  |
| 10:1.59 | על |  |
| 10:1.60 | טהרת |  |
| 10:1.61 | תוה"ק |  |
| 10:1.62 | שבכתב |  |
| 10:1.63 | ושבעל |  |
| 10:1.64 | פה | , |
| 10:1.65 | והחליפום |  |
| 10:1.66 | על |  |
| 10:1.67 | פי |  |
| 10:1.68 | ספרי |  |
| 10:1.69 | הקבלה |  |
| 10:1.70 | החדשה |  |
| 10:1.71 | שאין |  |
| 10:1.72 | לה |  |
| 10:1.73 | בית |  |
| 10:1.74 | אב | , |
| 10:1.75 | ובשקר |  |
| 10:1.76 | גמור |  |
| 10:1.77 | יחסוה |  |
| 10:1.78 | להתנא |  |
| 10:1.79 | הקדוש |  |
| 10:1.80 | רשב"י |  |
| 10:1.81 | ע"ה | , |
| 10:1.82 | וחס |  |
| 10:1.83 | ליה |  |
| 10:1.84 | לפה |  |
| 10:1.85 | קדוש |  |
| 10:1.86 | רשב"י |  |
| 10:1.87 | שיצאו |  |
| 10:1.88 | מפיו |  |
| 10:1.89 | דברים |  |
| 10:1.90 | שכאלה |  |
| 10:1.91 | ולומר |  |
| 10:1.92 | שנגלה |  |
| 10:1.93 | לו |  |
| 10:1.94 | עתיקא |  |
| 10:1.95 | קדישא |  |
| 10:1.96 | והרשהו |  |
| 10:1.97 | לחפות |  |
| 10:1.98 | דברים |  |
| 10:1.99 | אשר |  |
| 10:1.100 | לא |  |
| 10:1.101 | כן |  |
| 10:1.102 | על |  |
| 10:1.103 | ה | ' |
| 10:1.104 | אלהינו |  |
| 10:1.105 | מבלי |  |
| 10:1.106 | לחוס |  |
| 10:1.107 | על |  |
| 10:1.108 | כבוד |  |
| 10:1.109 | קונו | , |
| 10:1.110 | וליחס |  |
| 10:1.111 | בו |  |
| 10:1.112 | שינוי |  |
| 10:1.113 | ורבוי |  |
| 10:1.114 | איברים |  |
| 10:1.115 | וידים |  |
| 10:1.116 | ורגלים |  |
| 10:1.117 | וגילוי |  |
| 10:1.118 | מבושים |  |
| 10:1.119 | בזכרים |  |
| 10:1.120 | ונקבות |  |
| 10:1.121 | וזיווגיהם | , |
| 10:1.122 | חלילה |  |
| 10:1.123 | חלילה |  |
| 10:1.124 | להאמין |  |
| 10:1.125 | כן |  |
| 10:1.126 | על |  |
| 10:1.127 | התנא |  |
| 10:1.128 | הקדוש |  |
| 10:1.129 | הלזה | . |
| 10:1.130 | ועל |  |
| 10:1.131 | שאר |  |
| 10:1.132 | תנאים |  |
| 10:1.133 | ואמוראים |  |
| 10:1.134 | לחשוב |  |
| 10:1.135 | אף |  |
| 10:1.136 | בהרהור |  |
| 10:1.137 | שום |  |
| 10:1.138 | שינוי |  |
| 10:1.139 | ורבוי |  |
| 10:1.140 | כאל |  |
| 10:1.141 | ית | '. |
| 10:1.142 | וכ"ש |  |
| 10:1.143 | זכרים |  |
| 10:1.144 | ונקבות |  |
| 10:1.145 | שונים |  |
| 10:1.146 | זה |  |
| 10:1.147 | מזה | , |
| 10:1.148 | יש |  |
| 10:1.149 | זכר |  |
| 10:1.150 | ונקבה |  |
| 10:1.151 | בגוף |  |
| 10:1.152 | אחד | . |
| 10:1.153 | ויש |  |
| 10:1.154 | זכר |  |
| 10:1.155 | ונקבה |  |
| 10:1.156 | בפרצופים |  |
| 10:1.157 | חלוקים |  |
| 10:1.158 | איש |  |
| 10:1.159 | ואשתו |  |
| 10:1.160 | וגיד |  |
| 10:1.161 | ערוה | , |
| 10:1.162 | וביצי |  |
| 10:1.163 | זכר |  |
| 10:1.164 | לפרצופים |  |
| 10:1.165 | הזכרים |  |
| 10:1.166 | ומקור |  |
| 10:1.167 | ועיבור |  |
| 10:1.168 | ולידה |  |
| 10:1.169 | והנקה |  |
| 10:1.170 | לנקבות |  |
| 10:1.171 | מדדי |  |
| 10:1.172 | אדם |  |
| 10:1.173 | וגם |  |
| 10:1.174 | מדדי |  |
| 10:1.175 | בהמה | , |
| 10:1.176 | לא |  |
| 10:1.177 | כאלה |  |
| 10:1.178 | חלק |  |
| 10:1.179 | יעקב |  |
| 10:1.180 | להמיר |  |
| 10:1.181 | את |  |
| 10:1.182 | כבודם |  |
| 10:1.183 | האל |  |
| 10:1.184 | המושיעם |  |
| 10:1.185 | עושה |  |
| 10:1.186 | גדולות |  |
| 10:1.187 | במסרים |  |
| 10:1.188 | וכו | ', |
| 10:1.189 | כי |  |
| 10:1.190 | אחת |  |
| 10:1.191 | דבר |  |
| 10:1.192 | אלהים |  |
| 10:1.193 | שתים |  |
| 10:1.194 | זו |  |
| 10:1.195 | שמענו |  |
| 10:1.196 | אנכי |  |
| 10:1.197 | ולא |  |
| 10:1.198 | יהיה |  |
| 10:1.199 | לך | . |
| 10:1.200 | ואותו |  |
| 10:1.201 | לבדו |  |
| 10:1.202 | נעבוד |  |
| 10:1.203 | ונירא |  |
| 10:1.204 | מהדר |  |
| 10:1.205 | גאונו |  |
| 10:1.206 | ולא |  |
| 10:1.207 | נמירהו |  |
| 10:1.208 | בשום |  |
| 10:1.209 | נברא |  |
| 10:1.210 | או |  |
| 10:1.211 | נאצל | , |
| 10:1.212 | יהיה |  |
| 10:1.213 | מי |  |
| 10:1.214 | שיהיה |  |
| 10:1.215 | כמאמר |  |
| 10:1.216 | הכתוב |  |
| 10:1.217 | אל |  |
| 10:1.218 | תמר |  |
| 10:1.219 | בו | , |
| 10:1.220 | ופי | ' |
| 10:1.221 | רז"ל |  |
| 10:1.222 | אל |  |
| 10:1.223 | תמירני |  |
| 10:1.224 | בו |  |
| 10:1.225 | ואל |  |
| 10:1.226 | תמירהו |  |
| 10:1.227 | בי |  |
| 10:1.228 | רבה |  |
| 10:1.229 | פרשת |  |
| 10:1.230 | משפטים | ). |
| 10:1.231 | ובסנהדרין |  |
| 10:1.232 | פ | ' |
| 10:1.233 | א | ' |
| 10:1.234 | דיני |  |
| 10:1.235 | ממונות |  |
| 10:1.236 | אמר |  |
| 10:1.237 | ליה |  |
| 10:1.238 | ההוא |  |
| 10:1.239 | מינא |  |
| 10:1.240 | לרב |  |
| 10:1.241 | אידית |  |
| 10:1.242 | כתיב |  |
| 10:1.243 | ואל |  |
| 10:1.244 | משה |  |
| 10:1.245 | אמר |  |
| 10:1.246 | עלה |  |
| 10:1.247 | אל |  |
| 10:1.248 | ה | ', |
| 10:1.249 | עלה |  |
| 10:1.250 | אלי |  |
| 10:1.251 | מבעי |  |
| 10:1.252 | ליה |  |
| 10:1.253 | אמר |  |
| 10:1.254 | ליה |  |
| 10:1.255 | זהו |  |
| 10:1.256 | מטטרו"ן |  |
| 10:1.257 | ששמו |  |
| 10:1.258 | כשם |  |
| 10:1.259 | רבו |  |
| 10:1.260 | דכ | ' |
| 10:1.261 | כי |  |
| 10:1.262 | שמי |  |
| 10:1.263 | בקרבו | , |
| 10:1.264 | אמר |  |
| 10:1.265 | אי |  |
| 10:1.266 | הכי |  |
| 10:1.267 | נפלחו |  |
| 10:1.268 | ליה |  |
| 10:1.269 | אמר |  |
| 10:1.270 | ליה |  |
| 10:1.271 | כתוב |  |
| 10:1.272 | אל |  |
| 10:1.273 | תמר |  |
| 10:1.274 | בו | , |
| 10:1.275 | אל |  |
| 10:1.276 | תמירי |  |
| 10:1.277 | בו |  |
| 10:1.278 | וכו | ' |
| 10:1.279 | ועיין |  |
| 10:1.280 | רש"י |  |
| 10:1.281 | ומהרש"א | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

