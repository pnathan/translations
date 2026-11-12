# Translation pack: milhamot-hashem perek-5-07 (Milchamot ha-Shem 5:13–5:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-07`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **אדרא / אידרא** of Naso → "Idra of Naso" (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **השגחה** → "providence".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
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
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 5:13

Hebrew (Wikisource): מה) בין תבין ותשכיל ידידנו הקורא בשום לב איך שגו רבותינו המקובלים החדשים להטיל דופי בידיעת צופה ומביט עד סוף כל הדורות דבר שיהיה קודם שיהיה והשוו ידיעתו יתברך לידיעתם היפך דברי קדמונינו וחכמינו אנשי כנסת הגדולה מתקני התפילות והברכות, וגם היפך דברי התנא שאמר (אבות ג' ט"ו): "הכל צפוי (והרשות נתונה)" - עבר והוה ועתיד. והיפך דברי רבינו סעדיה גאון בספר האמונות והדעות והרמב"ם בספר המדע ובפירוש המשנה ובמורה. ולפי דבריהם ידיעתו יתברך מסופקת, לפי שאפשר לו לאדם לבחור בזולת מה שידע השי"ת הואיל והוא בעל בחירה! ומה שהביאם לזה הוא מפני השוואתם ידיעתו יתברך לידיעתם. ולא שמו אל לבם דברי הנביא שאמר (ישעיה נ"ה ח') "כי לא מחשבותי" וכו'. ולא די זה אלא שסלקו מהסיבה הראשונה ההשגחה במעשה התחתונים ואמרו איכה ידע אל בעתיד ואין דעה בעליון (עתיקא קדישא), ולפי דבריהם אין הפרש למעלה בין צדיק לרשע כמו שכתב האר"י, שכמו יעקב כן עשו אין הפרש ביניהם כטוב כחוטא. ורק בזעיר אנפין תלוי ההשגחה, והוא הדן ומשלם לאיש איש כפעלו וכמעלליו ישיב לו. ודבר זה מבואר באדרא פרשת האזינו דף רצ"ב דבזעיר אנפין כתיב כי אל דעות ה' ולו נתכנו עלילות בגין דאיהו בתרי גווני לו נתכנו עלילות. אבל לעתיקא קדישא סתימאה לא נתכנו. מאי טעמא נתכנו להאי בגין דירית תרין חולקין, וכתיב עם חסיד תתחסד וכו' ע"כ. ופירש בגיליון בשם הרש"ב הוא כמו שאמר בפרשת בלק "שולטנו יהיב ליה על כלא דכלא יפלחון ליה בגין דאעטרו ליה בדינא ורחמי. מאן דזכי לדינא לדינא ומאן דזכי לרחמי רחמי". מבואר מזה כי למעלה בנאצלים העליונים, לפי דעת הפילוסוף מחבר הזהר והמקובלים, אין מקפידים על מעשה התחתונים, כל שכן בסיבה הראשונה. ורק זעיר אנפין לבדו לו נתכנו עלילות, אבל למעלה לא נתכנו.

| id | word | punctuation after |
|---|---|---|
| 5:13.1 | מה | ) |
| 5:13.2 | בין |  |
| 5:13.3 | תבין |  |
| 5:13.4 | ותשכיל |  |
| 5:13.5 | ידידנו |  |
| 5:13.6 | הקורא |  |
| 5:13.7 | בשום |  |
| 5:13.8 | לב |  |
| 5:13.9 | איך |  |
| 5:13.10 | שגו |  |
| 5:13.11 | רבותינו |  |
| 5:13.12 | המקובלים |  |
| 5:13.13 | החדשים |  |
| 5:13.14 | להטיל |  |
| 5:13.15 | דופי |  |
| 5:13.16 | בידיעת |  |
| 5:13.17 | צופה |  |
| 5:13.18 | ומביט |  |
| 5:13.19 | עד |  |
| 5:13.20 | סוף |  |
| 5:13.21 | כל |  |
| 5:13.22 | הדורות |  |
| 5:13.23 | דבר |  |
| 5:13.24 | שיהיה |  |
| 5:13.25 | קודם |  |
| 5:13.26 | שיהיה |  |
| 5:13.27 | והשוו |  |
| 5:13.28 | ידיעתו |  |
| 5:13.29 | יתברך |  |
| 5:13.30 | לידיעתם |  |
| 5:13.31 | היפך |  |
| 5:13.32 | דברי |  |
| 5:13.33 | קדמונינו |  |
| 5:13.34 | וחכמינו |  |
| 5:13.35 | אנשי |  |
| 5:13.36 | כנסת |  |
| 5:13.37 | הגדולה |  |
| 5:13.38 | מתקני |  |
| 5:13.39 | התפילות |  |
| 5:13.40 | והברכות | , |
| 5:13.41 | וגם |  |
| 5:13.42 | היפך |  |
| 5:13.43 | דברי |  |
| 5:13.44 | התנא |  |
| 5:13.45 | שאמר |  |
| 5:13.46 | אבות |  |
| 5:13.47 | ג | ' |
| 5:13.48 | ט"ו | ): |
| 5:13.49 | הכל |  |
| 5:13.50 | צפוי |  |
| 5:13.51 | והרשות |  |
| 5:13.52 | נתונה | )"- |
| 5:13.53 | עבר |  |
| 5:13.54 | והוה |  |
| 5:13.55 | ועתיד | . |
| 5:13.56 | והיפך |  |
| 5:13.57 | דברי |  |
| 5:13.58 | רבינו |  |
| 5:13.59 | סעדיה |  |
| 5:13.60 | גאון |  |
| 5:13.61 | בספר |  |
| 5:13.62 | האמונות |  |
| 5:13.63 | והדעות |  |
| 5:13.64 | והרמב"ם |  |
| 5:13.65 | בספר |  |
| 5:13.66 | המדע |  |
| 5:13.67 | ובפירוש |  |
| 5:13.68 | המשנה |  |
| 5:13.69 | ובמורה | . |
| 5:13.70 | ולפי |  |
| 5:13.71 | דבריהם |  |
| 5:13.72 | ידיעתו |  |
| 5:13.73 | יתברך |  |
| 5:13.74 | מסופקת | , |
| 5:13.75 | לפי |  |
| 5:13.76 | שאפשר |  |
| 5:13.77 | לו |  |
| 5:13.78 | לאדם |  |
| 5:13.79 | לבחור |  |
| 5:13.80 | בזולת |  |
| 5:13.81 | מה |  |
| 5:13.82 | שידע |  |
| 5:13.83 | השי"ת |  |
| 5:13.84 | הואיל |  |
| 5:13.85 | והוא |  |
| 5:13.86 | בעל |  |
| 5:13.87 | בחירה | ! |
| 5:13.88 | ומה |  |
| 5:13.89 | שהביאם |  |
| 5:13.90 | לזה |  |
| 5:13.91 | הוא |  |
| 5:13.92 | מפני |  |
| 5:13.93 | השוואתם |  |
| 5:13.94 | ידיעתו |  |
| 5:13.95 | יתברך |  |
| 5:13.96 | לידיעתם | . |
| 5:13.97 | ולא |  |
| 5:13.98 | שמו |  |
| 5:13.99 | אל |  |
| 5:13.100 | לבם |  |
| 5:13.101 | דברי |  |
| 5:13.102 | הנביא |  |
| 5:13.103 | שאמר |  |
| 5:13.104 | ישעיה |  |
| 5:13.105 | נ"ה |  |
| 5:13.106 | ח | ') |
| 5:13.107 | כי |  |
| 5:13.108 | לא |  |
| 5:13.109 | מחשבותי | " |
| 5:13.110 | וכו | '. |
| 5:13.111 | ולא |  |
| 5:13.112 | די |  |
| 5:13.113 | זה |  |
| 5:13.114 | אלא |  |
| 5:13.115 | שסלקו |  |
| 5:13.116 | מהסיבה |  |
| 5:13.117 | הראשונה |  |
| 5:13.118 | ההשגחה |  |
| 5:13.119 | במעשה |  |
| 5:13.120 | התחתונים |  |
| 5:13.121 | ואמרו |  |
| 5:13.122 | איכה |  |
| 5:13.123 | ידע |  |
| 5:13.124 | אל |  |
| 5:13.125 | בעתיד |  |
| 5:13.126 | ואין |  |
| 5:13.127 | דעה |  |
| 5:13.128 | בעליון |  |
| 5:13.129 | עתיקא |  |
| 5:13.130 | קדישא | ), |
| 5:13.131 | ולפי |  |
| 5:13.132 | דבריהם |  |
| 5:13.133 | אין |  |
| 5:13.134 | הפרש |  |
| 5:13.135 | למעלה |  |
| 5:13.136 | בין |  |
| 5:13.137 | צדיק |  |
| 5:13.138 | לרשע |  |
| 5:13.139 | כמו |  |
| 5:13.140 | שכתב |  |
| 5:13.141 | האר"י | , |
| 5:13.142 | שכמו |  |
| 5:13.143 | יעקב |  |
| 5:13.144 | כן |  |
| 5:13.145 | עשו |  |
| 5:13.146 | אין |  |
| 5:13.147 | הפרש |  |
| 5:13.148 | ביניהם |  |
| 5:13.149 | כטוב |  |
| 5:13.150 | כחוטא | . |
| 5:13.151 | ורק |  |
| 5:13.152 | בזעיר |  |
| 5:13.153 | אנפין |  |
| 5:13.154 | תלוי |  |
| 5:13.155 | ההשגחה | , |
| 5:13.156 | והוא |  |
| 5:13.157 | הדן |  |
| 5:13.158 | ומשלם |  |
| 5:13.159 | לאיש |  |
| 5:13.160 | איש |  |
| 5:13.161 | כפעלו |  |
| 5:13.162 | וכמעלליו |  |
| 5:13.163 | ישיב |  |
| 5:13.164 | לו | . |
| 5:13.165 | ודבר |  |
| 5:13.166 | זה |  |
| 5:13.167 | מבואר |  |
| 5:13.168 | באדרא |  |
| 5:13.169 | פרשת |  |
| 5:13.170 | האזינו |  |
| 5:13.171 | דף |  |
| 5:13.172 | רצ"ב |  |
| 5:13.173 | דבזעיר |  |
| 5:13.174 | אנפין |  |
| 5:13.175 | כתיב |  |
| 5:13.176 | כי |  |
| 5:13.177 | אל |  |
| 5:13.178 | דעות |  |
| 5:13.179 | ה | ' |
| 5:13.180 | ולו |  |
| 5:13.181 | נתכנו |  |
| 5:13.182 | עלילות |  |
| 5:13.183 | בגין |  |
| 5:13.184 | דאיהו |  |
| 5:13.185 | בתרי |  |
| 5:13.186 | גווני |  |
| 5:13.187 | לו |  |
| 5:13.188 | נתכנו |  |
| 5:13.189 | עלילות | . |
| 5:13.190 | אבל |  |
| 5:13.191 | לעתיקא |  |
| 5:13.192 | קדישא |  |
| 5:13.193 | סתימאה |  |
| 5:13.194 | לא |  |
| 5:13.195 | נתכנו | . |
| 5:13.196 | מאי |  |
| 5:13.197 | טעמא |  |
| 5:13.198 | נתכנו |  |
| 5:13.199 | להאי |  |
| 5:13.200 | בגין |  |
| 5:13.201 | דירית |  |
| 5:13.202 | תרין |  |
| 5:13.203 | חולקין | , |
| 5:13.204 | וכתיב |  |
| 5:13.205 | עם |  |
| 5:13.206 | חסיד |  |
| 5:13.207 | תתחסד |  |
| 5:13.208 | וכו | ' |
| 5:13.209 | ע"כ | . |
| 5:13.210 | ופירש |  |
| 5:13.211 | בגיליון |  |
| 5:13.212 | בשם |  |
| 5:13.213 | הרש"ב |  |
| 5:13.214 | הוא |  |
| 5:13.215 | כמו |  |
| 5:13.216 | שאמר |  |
| 5:13.217 | בפרשת |  |
| 5:13.218 | בלק |  |
| 5:13.219 | שולטנו |  |
| 5:13.220 | יהיב |  |
| 5:13.221 | ליה |  |
| 5:13.222 | על |  |
| 5:13.223 | כלא |  |
| 5:13.224 | דכלא |  |
| 5:13.225 | יפלחון |  |
| 5:13.226 | ליה |  |
| 5:13.227 | בגין |  |
| 5:13.228 | דאעטרו |  |
| 5:13.229 | ליה |  |
| 5:13.230 | בדינא |  |
| 5:13.231 | ורחמי | . |
| 5:13.232 | מאן |  |
| 5:13.233 | דזכי |  |
| 5:13.234 | לדינא |  |
| 5:13.235 | לדינא |  |
| 5:13.236 | ומאן |  |
| 5:13.237 | דזכי |  |
| 5:13.238 | לרחמי |  |
| 5:13.239 | רחמי | ". |
| 5:13.240 | מבואר |  |
| 5:13.241 | מזה |  |
| 5:13.242 | כי |  |
| 5:13.243 | למעלה |  |
| 5:13.244 | בנאצלים |  |
| 5:13.245 | העליונים | , |
| 5:13.246 | לפי |  |
| 5:13.247 | דעת |  |
| 5:13.248 | הפילוסוף |  |
| 5:13.249 | מחבר |  |
| 5:13.250 | הזהר |  |
| 5:13.251 | והמקובלים | , |
| 5:13.252 | אין |  |
| 5:13.253 | מקפידים |  |
| 5:13.254 | על |  |
| 5:13.255 | מעשה |  |
| 5:13.256 | התחתונים | , |
| 5:13.257 | כל |  |
| 5:13.258 | שכן |  |
| 5:13.259 | בסיבה |  |
| 5:13.260 | הראשונה | . |
| 5:13.261 | ורק |  |
| 5:13.262 | זעיר |  |
| 5:13.263 | אנפין |  |
| 5:13.264 | לבדו |  |
| 5:13.265 | לו |  |
| 5:13.266 | נתכנו |  |
| 5:13.267 | עלילות | , |
| 5:13.268 | אבל |  |
| 5:13.269 | למעלה |  |
| 5:13.270 | לא |  |
| 5:13.271 | נתכנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

