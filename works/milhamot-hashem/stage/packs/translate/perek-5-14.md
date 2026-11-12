# Translation pack: milhamot-hashem perek-5-14 (Milchamot ha-Shem 5:21–5:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-14`.

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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **בא"י** → "in the Land of Israel".

## Terms

- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
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
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דרך ארץ** → "worldly occupation".
- **אורה** זו תורה (Esth 8:16) → "'orah' means Torah".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 5:21

Hebrew (Wikisource): מט) ומעתה אשובה ארעה בשושנים דברי רז"ל בגמרא ובמדרשים המפורסמים שממקום קדוש יצאו ואין בהם שום ספק, אשמור דבריהם היקרים מפז ויהלום באחדותו יתברך אחדות גמור בלי חיבור וצירוף ושיתוף שום נברא או נאצל וידיעת דרכיו ית' על פי התורה והמצוה אשר בעבודה בחר באבות שנאמר (בראשית י"ח י"ט): "כי ידעתיו למען אשר יצוה את בניו ואת ביתו אחריו ושמרו דרך השם" וכו' וכן עשה אברהם אבינו וציוה לבניו יצחק וישמעאל אמונת היחוד שהוא יתברך אחד אמת, ואין אחד אמת זולתו, ושהוא קדמון לכל נברא, ראשון ואין ראשית לראשיתו!

| id | word | punctuation after |
|---|---|---|
| 5:21.1 | מט | ) |
| 5:21.2 | ומעתה |  |
| 5:21.3 | אשובה |  |
| 5:21.4 | ארעה |  |
| 5:21.5 | בשושנים |  |
| 5:21.6 | דברי |  |
| 5:21.7 | רז"ל |  |
| 5:21.8 | בגמרא |  |
| 5:21.9 | ובמדרשים |  |
| 5:21.10 | המפורסמים |  |
| 5:21.11 | שממקום |  |
| 5:21.12 | קדוש |  |
| 5:21.13 | יצאו |  |
| 5:21.14 | ואין |  |
| 5:21.15 | בהם |  |
| 5:21.16 | שום |  |
| 5:21.17 | ספק | , |
| 5:21.18 | אשמור |  |
| 5:21.19 | דבריהם |  |
| 5:21.20 | היקרים |  |
| 5:21.21 | מפז |  |
| 5:21.22 | ויהלום |  |
| 5:21.23 | באחדותו |  |
| 5:21.24 | יתברך |  |
| 5:21.25 | אחדות |  |
| 5:21.26 | גמור |  |
| 5:21.27 | בלי |  |
| 5:21.28 | חיבור |  |
| 5:21.29 | וצירוף |  |
| 5:21.30 | ושיתוף |  |
| 5:21.31 | שום |  |
| 5:21.32 | נברא |  |
| 5:21.33 | או |  |
| 5:21.34 | נאצל |  |
| 5:21.35 | וידיעת |  |
| 5:21.36 | דרכיו |  |
| 5:21.37 | ית | ' |
| 5:21.38 | על |  |
| 5:21.39 | פי |  |
| 5:21.40 | התורה |  |
| 5:21.41 | והמצוה |  |
| 5:21.42 | אשר |  |
| 5:21.43 | בעבודה |  |
| 5:21.44 | בחר |  |
| 5:21.45 | באבות |  |
| 5:21.46 | שנאמר |  |
| 5:21.47 | בראשית |  |
| 5:21.48 | י"ח |  |
| 5:21.49 | י"ט | ): |
| 5:21.50 | כי |  |
| 5:21.51 | ידעתיו |  |
| 5:21.52 | למען |  |
| 5:21.53 | אשר |  |
| 5:21.54 | יצוה |  |
| 5:21.55 | את |  |
| 5:21.56 | בניו |  |
| 5:21.57 | ואת |  |
| 5:21.58 | ביתו |  |
| 5:21.59 | אחריו |  |
| 5:21.60 | ושמרו |  |
| 5:21.61 | דרך |  |
| 5:21.62 | השם | " |
| 5:21.63 | וכו | ' |
| 5:21.64 | וכן |  |
| 5:21.65 | עשה |  |
| 5:21.66 | אברהם |  |
| 5:21.67 | אבינו |  |
| 5:21.68 | וציוה |  |
| 5:21.69 | לבניו |  |
| 5:21.70 | יצחק |  |
| 5:21.71 | וישמעאל |  |
| 5:21.72 | אמונת |  |
| 5:21.73 | היחוד |  |
| 5:21.74 | שהוא |  |
| 5:21.75 | יתברך |  |
| 5:21.76 | אחד |  |
| 5:21.77 | אמת | , |
| 5:21.78 | ואין |  |
| 5:21.79 | אחד |  |
| 5:21.80 | אמת |  |
| 5:21.81 | זולתו | , |
| 5:21.82 | ושהוא |  |
| 5:21.83 | קדמון |  |
| 5:21.84 | לכל |  |
| 5:21.85 | נברא | , |
| 5:21.86 | ראשון |  |
| 5:21.87 | ואין |  |
| 5:21.88 | ראשית |  |
| 5:21.89 | לראשיתו | ! |


## Milchamot ha-Shem 5:22

Hebrew (Wikisource): וזה לשון מדרש הגדול: תרח אבי אברהם היה עובד ע"ז והיה עושה צלמים וכו' וכשמלך נמרוד הרשע באותו זמן ועשה עצמו אלוה וכו' והיו כל האומות באים ומשתחוים לו וכורעין לפניו עד שזרח אורו של עולם אברהם אבינו, וכשנולד אברהם אבינו אמרו לו אצטגניניו קוסמיו וחרטומיו: היום נולד נער שהוא עתיד לירש את העולם ועתיד לאבד את המלכיות ולהוציאך ממלכותך, אם רצונך שלח אצל אביו ופדהו ממנו בבית מלא כסף וזהב ועקרהו מן העולם, והיה תרח יושב ביניהם, אמר להם דומין דבריכם לפרד שאומרים לו טול כור שעורים ונחתוך את ראשך. מה יעשה בשעורים אחר שיחתכו את ראשו? כך זה כשייהרג בנו מה יעשה בדמיו? אמרו לו מתוך דבריך ניכר שנולד לך בן זה, אמר להם לא אכחד מכם בן נולד לי וכבר מת, אמרו לו לא על מת אנו מספרים אלא על חי ואנו רואים מזלו ברקיע, עכשיו לך והבא אותו, מיד הלך תרח ולקח את אברהם ואת מנקתו והטמינם במערה והיו שם שלוש שנים. אחר כך יצא מן המערה ובא אצל אביו והיה לבו משוטט ביצירת העולם ובמאורות להשתחוות להן ולעבדן ולידע איזה מהן הוא אלוה, ראה הלבנה שאורה מבהיק בלילה מסוף העולם ועד סופו ופמליא שלה מרובה, אמר זה אלוה. עברה כל הלילה. בשחרית כיון שזרחה חמה וחשך מאור הלבנה ותשש כחה, אמר זה אלוה. לערב חשכה ויצאה הלבנה והכוכבים והמזלות, אמר בודאי יש אדון לאלו וכו' כיון שנעשה לו נס ויצא בשלום, באותה שעה כיון אברהם אבינו דעתו לשמים ואמר, יתברך שמו של הקב"ה שחמה ולבנה כוכבים ומזלות אינם משתחוים אלא לו. ושמא תאמר רב היה לאברהם אלא זימן לו הקב"ה שתי כליותיו כשני מעינות והיו מוציאין לו חכמה שנאמר הן אמת חפצת בטוחות ובסתום חכמה תודיעני" וכו'.

| id | word | punctuation after |
|---|---|---|
| 5:22.1 | וזה |  |
| 5:22.2 | לשון |  |
| 5:22.3 | מדרש |  |
| 5:22.4 | הגדול | : |
| 5:22.5 | תרח |  |
| 5:22.6 | אבי |  |
| 5:22.7 | אברהם |  |
| 5:22.8 | היה |  |
| 5:22.9 | עובד |  |
| 5:22.10 | ע"ז |  |
| 5:22.11 | והיה |  |
| 5:22.12 | עושה |  |
| 5:22.13 | צלמים |  |
| 5:22.14 | וכו | ' |
| 5:22.15 | וכשמלך |  |
| 5:22.16 | נמרוד |  |
| 5:22.17 | הרשע |  |
| 5:22.18 | באותו |  |
| 5:22.19 | זמן |  |
| 5:22.20 | ועשה |  |
| 5:22.21 | עצמו |  |
| 5:22.22 | אלוה |  |
| 5:22.23 | וכו | ' |
| 5:22.24 | והיו |  |
| 5:22.25 | כל |  |
| 5:22.26 | האומות |  |
| 5:22.27 | באים |  |
| 5:22.28 | ומשתחוים |  |
| 5:22.29 | לו |  |
| 5:22.30 | וכורעין |  |
| 5:22.31 | לפניו |  |
| 5:22.32 | עד |  |
| 5:22.33 | שזרח |  |
| 5:22.34 | אורו |  |
| 5:22.35 | של |  |
| 5:22.36 | עולם |  |
| 5:22.37 | אברהם |  |
| 5:22.38 | אבינו | , |
| 5:22.39 | וכשנולד |  |
| 5:22.40 | אברהם |  |
| 5:22.41 | אבינו |  |
| 5:22.42 | אמרו |  |
| 5:22.43 | לו |  |
| 5:22.44 | אצטגניניו |  |
| 5:22.45 | קוסמיו |  |
| 5:22.46 | וחרטומיו | : |
| 5:22.47 | היום |  |
| 5:22.48 | נולד |  |
| 5:22.49 | נער |  |
| 5:22.50 | שהוא |  |
| 5:22.51 | עתיד |  |
| 5:22.52 | לירש |  |
| 5:22.53 | את |  |
| 5:22.54 | העולם |  |
| 5:22.55 | ועתיד |  |
| 5:22.56 | לאבד |  |
| 5:22.57 | את |  |
| 5:22.58 | המלכיות |  |
| 5:22.59 | ולהוציאך |  |
| 5:22.60 | ממלכותך | , |
| 5:22.61 | אם |  |
| 5:22.62 | רצונך |  |
| 5:22.63 | שלח |  |
| 5:22.64 | אצל |  |
| 5:22.65 | אביו |  |
| 5:22.66 | ופדהו |  |
| 5:22.67 | ממנו |  |
| 5:22.68 | בבית |  |
| 5:22.69 | מלא |  |
| 5:22.70 | כסף |  |
| 5:22.71 | וזהב |  |
| 5:22.72 | ועקרהו |  |
| 5:22.73 | מן |  |
| 5:22.74 | העולם | , |
| 5:22.75 | והיה |  |
| 5:22.76 | תרח |  |
| 5:22.77 | יושב |  |
| 5:22.78 | ביניהם | , |
| 5:22.79 | אמר |  |
| 5:22.80 | להם |  |
| 5:22.81 | דומין |  |
| 5:22.82 | דבריכם |  |
| 5:22.83 | לפרד |  |
| 5:22.84 | שאומרים |  |
| 5:22.85 | לו |  |
| 5:22.86 | טול |  |
| 5:22.87 | כור |  |
| 5:22.88 | שעורים |  |
| 5:22.89 | ונחתוך |  |
| 5:22.90 | את |  |
| 5:22.91 | ראשך | . |
| 5:22.92 | מה |  |
| 5:22.93 | יעשה |  |
| 5:22.94 | בשעורים |  |
| 5:22.95 | אחר |  |
| 5:22.96 | שיחתכו |  |
| 5:22.97 | את |  |
| 5:22.98 | ראשו | ? |
| 5:22.99 | כך |  |
| 5:22.100 | זה |  |
| 5:22.101 | כשייהרג |  |
| 5:22.102 | בנו |  |
| 5:22.103 | מה |  |
| 5:22.104 | יעשה |  |
| 5:22.105 | בדמיו | ? |
| 5:22.106 | אמרו |  |
| 5:22.107 | לו |  |
| 5:22.108 | מתוך |  |
| 5:22.109 | דבריך |  |
| 5:22.110 | ניכר |  |
| 5:22.111 | שנולד |  |
| 5:22.112 | לך |  |
| 5:22.113 | בן |  |
| 5:22.114 | זה | , |
| 5:22.115 | אמר |  |
| 5:22.116 | להם |  |
| 5:22.117 | לא |  |
| 5:22.118 | אכחד |  |
| 5:22.119 | מכם |  |
| 5:22.120 | בן |  |
| 5:22.121 | נולד |  |
| 5:22.122 | לי |  |
| 5:22.123 | וכבר |  |
| 5:22.124 | מת | , |
| 5:22.125 | אמרו |  |
| 5:22.126 | לו |  |
| 5:22.127 | לא |  |
| 5:22.128 | על |  |
| 5:22.129 | מת |  |
| 5:22.130 | אנו |  |
| 5:22.131 | מספרים |  |
| 5:22.132 | אלא |  |
| 5:22.133 | על |  |
| 5:22.134 | חי |  |
| 5:22.135 | ואנו |  |
| 5:22.136 | רואים |  |
| 5:22.137 | מזלו |  |
| 5:22.138 | ברקיע | , |
| 5:22.139 | עכשיו |  |
| 5:22.140 | לך |  |
| 5:22.141 | והבא |  |
| 5:22.142 | אותו | , |
| 5:22.143 | מיד |  |
| 5:22.144 | הלך |  |
| 5:22.145 | תרח |  |
| 5:22.146 | ולקח |  |
| 5:22.147 | את |  |
| 5:22.148 | אברהם |  |
| 5:22.149 | ואת |  |
| 5:22.150 | מנקתו |  |
| 5:22.151 | והטמינם |  |
| 5:22.152 | במערה |  |
| 5:22.153 | והיו |  |
| 5:22.154 | שם |  |
| 5:22.155 | שלוש |  |
| 5:22.156 | שנים | . |
| 5:22.157 | אחר |  |
| 5:22.158 | כך |  |
| 5:22.159 | יצא |  |
| 5:22.160 | מן |  |
| 5:22.161 | המערה |  |
| 5:22.162 | ובא |  |
| 5:22.163 | אצל |  |
| 5:22.164 | אביו |  |
| 5:22.165 | והיה |  |
| 5:22.166 | לבו |  |
| 5:22.167 | משוטט |  |
| 5:22.168 | ביצירת |  |
| 5:22.169 | העולם |  |
| 5:22.170 | ובמאורות |  |
| 5:22.171 | להשתחוות |  |
| 5:22.172 | להן |  |
| 5:22.173 | ולעבדן |  |
| 5:22.174 | ולידע |  |
| 5:22.175 | איזה |  |
| 5:22.176 | מהן |  |
| 5:22.177 | הוא |  |
| 5:22.178 | אלוה | , |
| 5:22.179 | ראה |  |
| 5:22.180 | הלבנה |  |
| 5:22.181 | שאורה |  |
| 5:22.182 | מבהיק |  |
| 5:22.183 | בלילה |  |
| 5:22.184 | מסוף |  |
| 5:22.185 | העולם |  |
| 5:22.186 | ועד |  |
| 5:22.187 | סופו |  |
| 5:22.188 | ופמליא |  |
| 5:22.189 | שלה |  |
| 5:22.190 | מרובה | , |
| 5:22.191 | אמר |  |
| 5:22.192 | זה |  |
| 5:22.193 | אלוה | . |
| 5:22.194 | עברה |  |
| 5:22.195 | כל |  |
| 5:22.196 | הלילה | . |
| 5:22.197 | בשחרית |  |
| 5:22.198 | כיון |  |
| 5:22.199 | שזרחה |  |
| 5:22.200 | חמה |  |
| 5:22.201 | וחשך |  |
| 5:22.202 | מאור |  |
| 5:22.203 | הלבנה |  |
| 5:22.204 | ותשש |  |
| 5:22.205 | כחה | , |
| 5:22.206 | אמר |  |
| 5:22.207 | זה |  |
| 5:22.208 | אלוה | . |
| 5:22.209 | לערב |  |
| 5:22.210 | חשכה |  |
| 5:22.211 | ויצאה |  |
| 5:22.212 | הלבנה |  |
| 5:22.213 | והכוכבים |  |
| 5:22.214 | והמזלות | , |
| 5:22.215 | אמר |  |
| 5:22.216 | בודאי |  |
| 5:22.217 | יש |  |
| 5:22.218 | אדון |  |
| 5:22.219 | לאלו |  |
| 5:22.220 | וכו | ' |
| 5:22.221 | כיון |  |
| 5:22.222 | שנעשה |  |
| 5:22.223 | לו |  |
| 5:22.224 | נס |  |
| 5:22.225 | ויצא |  |
| 5:22.226 | בשלום | , |
| 5:22.227 | באותה |  |
| 5:22.228 | שעה |  |
| 5:22.229 | כיון |  |
| 5:22.230 | אברהם |  |
| 5:22.231 | אבינו |  |
| 5:22.232 | דעתו |  |
| 5:22.233 | לשמים |  |
| 5:22.234 | ואמר | , |
| 5:22.235 | יתברך |  |
| 5:22.236 | שמו |  |
| 5:22.237 | של |  |
| 5:22.238 | הקב"ה |  |
| 5:22.239 | שחמה |  |
| 5:22.240 | ולבנה |  |
| 5:22.241 | כוכבים |  |
| 5:22.242 | ומזלות |  |
| 5:22.243 | אינם |  |
| 5:22.244 | משתחוים |  |
| 5:22.245 | אלא |  |
| 5:22.246 | לו | . |
| 5:22.247 | ושמא |  |
| 5:22.248 | תאמר |  |
| 5:22.249 | רב |  |
| 5:22.250 | היה |  |
| 5:22.251 | לאברהם |  |
| 5:22.252 | אלא |  |
| 5:22.253 | זימן |  |
| 5:22.254 | לו |  |
| 5:22.255 | הקב"ה |  |
| 5:22.256 | שתי |  |
| 5:22.257 | כליותיו |  |
| 5:22.258 | כשני |  |
| 5:22.259 | מעינות |  |
| 5:22.260 | והיו |  |
| 5:22.261 | מוציאין |  |
| 5:22.262 | לו |  |
| 5:22.263 | חכמה |  |
| 5:22.264 | שנאמר |  |
| 5:22.265 | הן |  |
| 5:22.266 | אמת |  |
| 5:22.267 | חפצת |  |
| 5:22.268 | בטוחות |  |
| 5:22.269 | ובסתום |  |
| 5:22.270 | חכמה |  |
| 5:22.271 | תודיעני | " |
| 5:22.272 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

