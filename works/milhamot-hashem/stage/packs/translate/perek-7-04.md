# Translation pack: milhamot-hashem perek-7-04 (Milchamot ha-Shem 7:5–7:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-04`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
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
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- אל **אחר** → "another god".
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
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 7:5

Hebrew (Wikisource): פב) ובמדרש רבה פרשת ואתחנן. שמע ישאל ה' אלהינו ה' אחד, זהו שאמר הכתוב מי לי בשמים ועמך לא חפצתי בארץ, רב אמר ב' רקיעים הם שמים ושמי השמים. ר' אלעזר אמר שבעה רקיעים הם שמים ושמי השמים. רקיע, שחקים, מעון, זבול, ערפל, וכולו פתח הקב"ה לישראל להודיען שאין אלוה אחר אלא הוא. אמרה כנסת ישראל לפני הקב"ה. מי לי בשמים חוץ מכבודך. כשם שאין לי בשמים אלא אתה, כך לא חפצתי אחר בארץ. אלא אני נכנסת בכל יום לבתי כנסיות ומעידה עליך שאין אלוה אחר אלא אתה ואומרת שמע ישראל ה' אלהינו ה' אחד. הנך רואה ידידי דברי רז"ל הנז' במדרש עומדים נגד הקבלה החדשה. שהרי בשמע ישראל שהוא קבלת עול מלכות שמים, המקובלים החדשים משתפים בו אבא ואימא וזעיר ונוקביה כמבואר בזהר ואתחנן. הובא לעיל סי' מ"ב, מה שלא עלה על דעת רז"ל שייעשה כן בישראל וגדולה מזו. אמרו בעלי הקבלה החדשה, שהמתפלל לאין סוף הסבה הראשונה אין תפלתו תפלה, אלא צריך להתפלל דוקא לז"א ליחד זוגתו עמו וגם לשתף במחשבה אבא ואימא ואריך אנפין ועתיק יומין כמו שכתבתי לעיל ב' כסא אליהו ויושר לבב. שהמיחד תפלתו לז"א בלי צירוף ושתוף פרצופים הנז' אינו נענה מהרה, וזה סוד מה תצעק אלי כדלעיל סי' ל'. עוד אמרו במדרש רבה פרשה הנז"ל. וז"ל שמע ישראל ה' אלהינו ה' אחד, רבנן אמרי אמר הקב"ה בני כל מה שבראתי בראתי זוגות, שמים וארץ זוגות, חמה ולבנה זוגות, אדם וחוה זוגות, העולם הזה והעה"ב זוגות, אבל כבודי אחד ומיוחד בעולם, מנין ממה שקרינו בענין שמע ישראל ה' אלהינו ה' אחד עכ"ל. גם דברי המדרש אלו עומדים נגד הקבלה החדשה שאמרה שזעיר ונוקביה הוא אלוהינו ואנחנו עמו וצאן מריעתו. ואותו נעבוד במצות אבא ואימא, והוא משתבח בעצמו ואומר ה' אמר אלי בני אתה וכו' כדלעיל סי' כ"ו. שהרי במדרש אמרו כי ה' אלהינו אחד ומיוחד מכל האחדים שבעולם כנז' לעיל, ואין לו עזר כנגדו לפי שהוא יכול ואינו צריך לעזר זולתו.

| id | word | punctuation after |
|---|---|---|
| 7:5.1 | פב | ) |
| 7:5.2 | ובמדרש |  |
| 7:5.3 | רבה |  |
| 7:5.4 | פרשת |  |
| 7:5.5 | ואתחנן | . |
| 7:5.6 | שמע |  |
| 7:5.7 | ישאל |  |
| 7:5.8 | ה | ' |
| 7:5.9 | אלהינו |  |
| 7:5.10 | ה | ' |
| 7:5.11 | אחד | , |
| 7:5.12 | זהו |  |
| 7:5.13 | שאמר |  |
| 7:5.14 | הכתוב |  |
| 7:5.15 | מי |  |
| 7:5.16 | לי |  |
| 7:5.17 | בשמים |  |
| 7:5.18 | ועמך |  |
| 7:5.19 | לא |  |
| 7:5.20 | חפצתי |  |
| 7:5.21 | בארץ | , |
| 7:5.22 | רב |  |
| 7:5.23 | אמר |  |
| 7:5.24 | ב | ' |
| 7:5.25 | רקיעים |  |
| 7:5.26 | הם |  |
| 7:5.27 | שמים |  |
| 7:5.28 | ושמי |  |
| 7:5.29 | השמים | . |
| 7:5.30 | ר | ' |
| 7:5.31 | אלעזר |  |
| 7:5.32 | אמר |  |
| 7:5.33 | שבעה |  |
| 7:5.34 | רקיעים |  |
| 7:5.35 | הם |  |
| 7:5.36 | שמים |  |
| 7:5.37 | ושמי |  |
| 7:5.38 | השמים | . |
| 7:5.39 | רקיע | , |
| 7:5.40 | שחקים | , |
| 7:5.41 | מעון | , |
| 7:5.42 | זבול | , |
| 7:5.43 | ערפל | , |
| 7:5.44 | וכולו |  |
| 7:5.45 | פתח |  |
| 7:5.46 | הקב"ה |  |
| 7:5.47 | לישראל |  |
| 7:5.48 | להודיען |  |
| 7:5.49 | שאין |  |
| 7:5.50 | אלוה |  |
| 7:5.51 | אחר |  |
| 7:5.52 | אלא |  |
| 7:5.53 | הוא | . |
| 7:5.54 | אמרה |  |
| 7:5.55 | כנסת |  |
| 7:5.56 | ישראל |  |
| 7:5.57 | לפני |  |
| 7:5.58 | הקב"ה | . |
| 7:5.59 | מי |  |
| 7:5.60 | לי |  |
| 7:5.61 | בשמים |  |
| 7:5.62 | חוץ |  |
| 7:5.63 | מכבודך | . |
| 7:5.64 | כשם |  |
| 7:5.65 | שאין |  |
| 7:5.66 | לי |  |
| 7:5.67 | בשמים |  |
| 7:5.68 | אלא |  |
| 7:5.69 | אתה | , |
| 7:5.70 | כך |  |
| 7:5.71 | לא |  |
| 7:5.72 | חפצתי |  |
| 7:5.73 | אחר |  |
| 7:5.74 | בארץ | . |
| 7:5.75 | אלא |  |
| 7:5.76 | אני |  |
| 7:5.77 | נכנסת |  |
| 7:5.78 | בכל |  |
| 7:5.79 | יום |  |
| 7:5.80 | לבתי |  |
| 7:5.81 | כנסיות |  |
| 7:5.82 | ומעידה |  |
| 7:5.83 | עליך |  |
| 7:5.84 | שאין |  |
| 7:5.85 | אלוה |  |
| 7:5.86 | אחר |  |
| 7:5.87 | אלא |  |
| 7:5.88 | אתה |  |
| 7:5.89 | ואומרת |  |
| 7:5.90 | שמע |  |
| 7:5.91 | ישראל |  |
| 7:5.92 | ה | ' |
| 7:5.93 | אלהינו |  |
| 7:5.94 | ה | ' |
| 7:5.95 | אחד | . |
| 7:5.96 | הנך |  |
| 7:5.97 | רואה |  |
| 7:5.98 | ידידי |  |
| 7:5.99 | דברי |  |
| 7:5.100 | רז"ל |  |
| 7:5.101 | הנז | ' |
| 7:5.102 | במדרש |  |
| 7:5.103 | עומדים |  |
| 7:5.104 | נגד |  |
| 7:5.105 | הקבלה |  |
| 7:5.106 | החדשה | . |
| 7:5.107 | שהרי |  |
| 7:5.108 | בשמע |  |
| 7:5.109 | ישראל |  |
| 7:5.110 | שהוא |  |
| 7:5.111 | קבלת |  |
| 7:5.112 | עול |  |
| 7:5.113 | מלכות |  |
| 7:5.114 | שמים | , |
| 7:5.115 | המקובלים |  |
| 7:5.116 | החדשים |  |
| 7:5.117 | משתפים |  |
| 7:5.118 | בו |  |
| 7:5.119 | אבא |  |
| 7:5.120 | ואימא |  |
| 7:5.121 | וזעיר |  |
| 7:5.122 | ונוקביה |  |
| 7:5.123 | כמבואר |  |
| 7:5.124 | בזהר |  |
| 7:5.125 | ואתחנן | . |
| 7:5.126 | הובא |  |
| 7:5.127 | לעיל |  |
| 7:5.128 | סי | ' |
| 7:5.129 | מ"ב | , |
| 7:5.130 | מה |  |
| 7:5.131 | שלא |  |
| 7:5.132 | עלה |  |
| 7:5.133 | על |  |
| 7:5.134 | דעת |  |
| 7:5.135 | רז"ל |  |
| 7:5.136 | שייעשה |  |
| 7:5.137 | כן |  |
| 7:5.138 | בישראל |  |
| 7:5.139 | וגדולה |  |
| 7:5.140 | מזו | . |
| 7:5.141 | אמרו |  |
| 7:5.142 | בעלי |  |
| 7:5.143 | הקבלה |  |
| 7:5.144 | החדשה | , |
| 7:5.145 | שהמתפלל |  |
| 7:5.146 | לאין |  |
| 7:5.147 | סוף |  |
| 7:5.148 | הסבה |  |
| 7:5.149 | הראשונה |  |
| 7:5.150 | אין |  |
| 7:5.151 | תפלתו |  |
| 7:5.152 | תפלה | , |
| 7:5.153 | אלא |  |
| 7:5.154 | צריך |  |
| 7:5.155 | להתפלל |  |
| 7:5.156 | דוקא |  |
| 7:5.157 | לז"א |  |
| 7:5.158 | ליחד |  |
| 7:5.159 | זוגתו |  |
| 7:5.160 | עמו |  |
| 7:5.161 | וגם |  |
| 7:5.162 | לשתף |  |
| 7:5.163 | במחשבה |  |
| 7:5.164 | אבא |  |
| 7:5.165 | ואימא |  |
| 7:5.166 | ואריך |  |
| 7:5.167 | אנפין |  |
| 7:5.168 | ועתיק |  |
| 7:5.169 | יומין |  |
| 7:5.170 | כמו |  |
| 7:5.171 | שכתבתי |  |
| 7:5.172 | לעיל |  |
| 7:5.173 | ב | ' |
| 7:5.174 | כסא |  |
| 7:5.175 | אליהו |  |
| 7:5.176 | ויושר |  |
| 7:5.177 | לבב | . |
| 7:5.178 | שהמיחד |  |
| 7:5.179 | תפלתו |  |
| 7:5.180 | לז"א |  |
| 7:5.181 | בלי |  |
| 7:5.182 | צירוף |  |
| 7:5.183 | ושתוף |  |
| 7:5.184 | פרצופים |  |
| 7:5.185 | הנז | ' |
| 7:5.186 | אינו |  |
| 7:5.187 | נענה |  |
| 7:5.188 | מהרה | , |
| 7:5.189 | וזה |  |
| 7:5.190 | סוד |  |
| 7:5.191 | מה |  |
| 7:5.192 | תצעק |  |
| 7:5.193 | אלי |  |
| 7:5.194 | כדלעיל |  |
| 7:5.195 | סי | ' |
| 7:5.196 | ל | '. |
| 7:5.197 | עוד |  |
| 7:5.198 | אמרו |  |
| 7:5.199 | במדרש |  |
| 7:5.200 | רבה |  |
| 7:5.201 | פרשה |  |
| 7:5.202 | הנז"ל | . |
| 7:5.203 | וז"ל |  |
| 7:5.204 | שמע |  |
| 7:5.205 | ישראל |  |
| 7:5.206 | ה | ' |
| 7:5.207 | אלהינו |  |
| 7:5.208 | ה | ' |
| 7:5.209 | אחד | , |
| 7:5.210 | רבנן |  |
| 7:5.211 | אמרי |  |
| 7:5.212 | אמר |  |
| 7:5.213 | הקב"ה |  |
| 7:5.214 | בני |  |
| 7:5.215 | כל |  |
| 7:5.216 | מה |  |
| 7:5.217 | שבראתי |  |
| 7:5.218 | בראתי |  |
| 7:5.219 | זוגות | , |
| 7:5.220 | שמים |  |
| 7:5.221 | וארץ |  |
| 7:5.222 | זוגות | , |
| 7:5.223 | חמה |  |
| 7:5.224 | ולבנה |  |
| 7:5.225 | זוגות | , |
| 7:5.226 | אדם |  |
| 7:5.227 | וחוה |  |
| 7:5.228 | זוגות | , |
| 7:5.229 | העולם |  |
| 7:5.230 | הזה |  |
| 7:5.231 | והעה"ב |  |
| 7:5.232 | זוגות | , |
| 7:5.233 | אבל |  |
| 7:5.234 | כבודי |  |
| 7:5.235 | אחד |  |
| 7:5.236 | ומיוחד |  |
| 7:5.237 | בעולם | , |
| 7:5.238 | מנין |  |
| 7:5.239 | ממה |  |
| 7:5.240 | שקרינו |  |
| 7:5.241 | בענין |  |
| 7:5.242 | שמע |  |
| 7:5.243 | ישראל |  |
| 7:5.244 | ה | ' |
| 7:5.245 | אלהינו |  |
| 7:5.246 | ה | ' |
| 7:5.247 | אחד |  |
| 7:5.248 | עכ"ל | . |
| 7:5.249 | גם |  |
| 7:5.250 | דברי |  |
| 7:5.251 | המדרש |  |
| 7:5.252 | אלו |  |
| 7:5.253 | עומדים |  |
| 7:5.254 | נגד |  |
| 7:5.255 | הקבלה |  |
| 7:5.256 | החדשה |  |
| 7:5.257 | שאמרה |  |
| 7:5.258 | שזעיר |  |
| 7:5.259 | ונוקביה |  |
| 7:5.260 | הוא |  |
| 7:5.261 | אלוהינו |  |
| 7:5.262 | ואנחנו |  |
| 7:5.263 | עמו |  |
| 7:5.264 | וצאן |  |
| 7:5.265 | מריעתו | . |
| 7:5.266 | ואותו |  |
| 7:5.267 | נעבוד |  |
| 7:5.268 | במצות |  |
| 7:5.269 | אבא |  |
| 7:5.270 | ואימא | , |
| 7:5.271 | והוא |  |
| 7:5.272 | משתבח |  |
| 7:5.273 | בעצמו |  |
| 7:5.274 | ואומר |  |
| 7:5.275 | ה | ' |
| 7:5.276 | אמר |  |
| 7:5.277 | אלי |  |
| 7:5.278 | בני |  |
| 7:5.279 | אתה |  |
| 7:5.280 | וכו | ' |
| 7:5.281 | כדלעיל |  |
| 7:5.282 | סי | ' |
| 7:5.283 | כ"ו | . |
| 7:5.284 | שהרי |  |
| 7:5.285 | במדרש |  |
| 7:5.286 | אמרו |  |
| 7:5.287 | כי |  |
| 7:5.288 | ה | ' |
| 7:5.289 | אלהינו |  |
| 7:5.290 | אחד |  |
| 7:5.291 | ומיוחד |  |
| 7:5.292 | מכל |  |
| 7:5.293 | האחדים |  |
| 7:5.294 | שבעולם |  |
| 7:5.295 | כנז | ' |
| 7:5.296 | לעיל | , |
| 7:5.297 | ואין |  |
| 7:5.298 | לו |  |
| 7:5.299 | עזר |  |
| 7:5.300 | כנגדו |  |
| 7:5.301 | לפי |  |
| 7:5.302 | שהוא |  |
| 7:5.303 | יכול |  |
| 7:5.304 | ואינו |  |
| 7:5.305 | צריך |  |
| 7:5.306 | לעזר |  |
| 7:5.307 | זולתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

