# Translation pack: milhamot-hashem perek-5-18 (Milchamot ha-Shem 5:27–5:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-18`.

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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **הרמב"ן / רמב"ן** → "Ramban" (authority, transliterated; note first in a
  unit).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הכוזרי / כוזרי** → "the Kuzari" (book title; meaning in a note).
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

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- קורא **נעים / נעלם** → "pleasant reader".
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
- בחלק **שפתיו** → "with the smoothness of his lips".
- מי **אתי** (ketiv): translated, the qere noted.
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 5:27

Hebrew (Wikisource): נג) תנה לבך קורא נעים בדברי הר' הלזה בתמונת היחוד האמתי. והרמב"ן הביא דבריו בקצור באגרתו לרבני צרפת. היא אמונתינו הטהורה מכל רעיונות מדומות מזויפות. היא הנובעת ממקור נאמן מרז"ל שהביא הר' הנז' מבראשית רבה, ותורת כהנים, וספר יצירה, וספרי וגמרא, כולם נכוחים למבין. כל דבריהם מכוונים ומתאימים לדברי הרמב"ם ז"ל בחיבורו הגדול היד. ובפי' המשנה ובמורה, והר' סעדיה גאון בס' האמונות והדעות ובתרגומו הערבי למקרא ור' יהודה הלוי בכוזרי ובתחנותיו ושיריו, ור' בחיי בחובת הלבבות, ושבילי אמונה, וס' מצות הגדול, שהוא ית' אחד מיוחד שאין אחדות כמהו ואיננו כשאר האחדים, אלא אחדות שאין כמוה. ודמיון כל הברואים אין בבורא. וכל צלם העשוים אין בעושה, ושאין לו שום פרצוף מכל הפרצופים שמדמים בעלי הקבלה החדשה. ושאסור לאדם להרהר ולצייר בדעתו היאך הוא, והיאך הנייתו. כמאמר רז"ל ולא תתורו אחרי לבבכם זו מינות ואין לך מינות גדולה מזו לחשוב לבורא ית' כמה פרצופים חלוקים בשמותן ובאיכותן ובפעולותיהם. ותמיד מזדווגים, ועושים תולדות זה מוליד נשמות למלאכים, וזה מוליד נשמות לישראל (זהר יתרו דף פ"ט) וסטרא אחרא סמא"ל בספירותיו הנקרא אל אחד ואדם בליעל מוליד נשמות לגוים. ואחר שסיפר לנו כמה פרצופים שונים הנקראים בשמותיו של הקב"ה. מה יתן ומה יוסיף מה שהוא משקר בלשונו אחר כך ואומר וכלא חד. והוא כבר דיבר בפיו מה שכבר חישב בלבו. וגם הוציא בשפתיו וכתב בקולמוסו שהם רבים. וכי אתי דיבור לחודיה ומבטל מחשבה ודבור מעשה שכתב בספר? אכן האמת הגמור המקובל אצל אומותינו מדור דור. כי באמת ובתמים אין שיעור וקצב לאלהינו יתברך שמו ולא אורך ולא רוחב, ולא דמות ולא צורת גולם. וכל שכן שאין לו אשה, שתהיה לו לעזר ולהועיל, שהוא ית' אינו צריך לעזר. כי לגדולתו אין חקר ולא סוף, והוא ראשון ואין לו ראשית לראשיתו.

| id | word | punctuation after |
|---|---|---|
| 5:27.1 | נג | ) |
| 5:27.2 | תנה |  |
| 5:27.3 | לבך |  |
| 5:27.4 | קורא |  |
| 5:27.5 | נעים |  |
| 5:27.6 | בדברי |  |
| 5:27.7 | הר | ' |
| 5:27.8 | הלזה |  |
| 5:27.9 | בתמונת |  |
| 5:27.10 | היחוד |  |
| 5:27.11 | האמתי | . |
| 5:27.12 | והרמב"ן |  |
| 5:27.13 | הביא |  |
| 5:27.14 | דבריו |  |
| 5:27.15 | בקצור |  |
| 5:27.16 | באגרתו |  |
| 5:27.17 | לרבני |  |
| 5:27.18 | צרפת | . |
| 5:27.19 | היא |  |
| 5:27.20 | אמונתינו |  |
| 5:27.21 | הטהורה |  |
| 5:27.22 | מכל |  |
| 5:27.23 | רעיונות |  |
| 5:27.24 | מדומות |  |
| 5:27.25 | מזויפות | . |
| 5:27.26 | היא |  |
| 5:27.27 | הנובעת |  |
| 5:27.28 | ממקור |  |
| 5:27.29 | נאמן |  |
| 5:27.30 | מרז"ל |  |
| 5:27.31 | שהביא |  |
| 5:27.32 | הר | ' |
| 5:27.33 | הנז | ' |
| 5:27.34 | מבראשית |  |
| 5:27.35 | רבה | , |
| 5:27.36 | ותורת |  |
| 5:27.37 | כהנים | , |
| 5:27.38 | וספר |  |
| 5:27.39 | יצירה | , |
| 5:27.40 | וספרי |  |
| 5:27.41 | וגמרא | , |
| 5:27.42 | כולם |  |
| 5:27.43 | נכוחים |  |
| 5:27.44 | למבין | . |
| 5:27.45 | כל |  |
| 5:27.46 | דבריהם |  |
| 5:27.47 | מכוונים |  |
| 5:27.48 | ומתאימים |  |
| 5:27.49 | לדברי |  |
| 5:27.50 | הרמב"ם |  |
| 5:27.51 | ז"ל |  |
| 5:27.52 | בחיבורו |  |
| 5:27.53 | הגדול |  |
| 5:27.54 | היד | . |
| 5:27.55 | ובפי | ' |
| 5:27.56 | המשנה |  |
| 5:27.57 | ובמורה | , |
| 5:27.58 | והר | ' |
| 5:27.59 | סעדיה |  |
| 5:27.60 | גאון |  |
| 5:27.61 | בס | ' |
| 5:27.62 | האמונות |  |
| 5:27.63 | והדעות |  |
| 5:27.64 | ובתרגומו |  |
| 5:27.65 | הערבי |  |
| 5:27.66 | למקרא |  |
| 5:27.67 | ור | ' |
| 5:27.68 | יהודה |  |
| 5:27.69 | הלוי |  |
| 5:27.70 | בכוזרי |  |
| 5:27.71 | ובתחנותיו |  |
| 5:27.72 | ושיריו | , |
| 5:27.73 | ור | ' |
| 5:27.74 | בחיי |  |
| 5:27.75 | בחובת |  |
| 5:27.76 | הלבבות | , |
| 5:27.77 | ושבילי |  |
| 5:27.78 | אמונה | , |
| 5:27.79 | וס | ' |
| 5:27.80 | מצות |  |
| 5:27.81 | הגדול | , |
| 5:27.82 | שהוא |  |
| 5:27.83 | ית | ' |
| 5:27.84 | אחד |  |
| 5:27.85 | מיוחד |  |
| 5:27.86 | שאין |  |
| 5:27.87 | אחדות |  |
| 5:27.88 | כמהו |  |
| 5:27.89 | ואיננו |  |
| 5:27.90 | כשאר |  |
| 5:27.91 | האחדים | , |
| 5:27.92 | אלא |  |
| 5:27.93 | אחדות |  |
| 5:27.94 | שאין |  |
| 5:27.95 | כמוה | . |
| 5:27.96 | ודמיון |  |
| 5:27.97 | כל |  |
| 5:27.98 | הברואים |  |
| 5:27.99 | אין |  |
| 5:27.100 | בבורא | . |
| 5:27.101 | וכל |  |
| 5:27.102 | צלם |  |
| 5:27.103 | העשוים |  |
| 5:27.104 | אין |  |
| 5:27.105 | בעושה | , |
| 5:27.106 | ושאין |  |
| 5:27.107 | לו |  |
| 5:27.108 | שום |  |
| 5:27.109 | פרצוף |  |
| 5:27.110 | מכל |  |
| 5:27.111 | הפרצופים |  |
| 5:27.112 | שמדמים |  |
| 5:27.113 | בעלי |  |
| 5:27.114 | הקבלה |  |
| 5:27.115 | החדשה | . |
| 5:27.116 | ושאסור |  |
| 5:27.117 | לאדם |  |
| 5:27.118 | להרהר |  |
| 5:27.119 | ולצייר |  |
| 5:27.120 | בדעתו |  |
| 5:27.121 | היאך |  |
| 5:27.122 | הוא | , |
| 5:27.123 | והיאך |  |
| 5:27.124 | הנייתו | . |
| 5:27.125 | כמאמר |  |
| 5:27.126 | רז"ל |  |
| 5:27.127 | ולא |  |
| 5:27.128 | תתורו |  |
| 5:27.129 | אחרי |  |
| 5:27.130 | לבבכם |  |
| 5:27.131 | זו |  |
| 5:27.132 | מינות |  |
| 5:27.133 | ואין |  |
| 5:27.134 | לך |  |
| 5:27.135 | מינות |  |
| 5:27.136 | גדולה |  |
| 5:27.137 | מזו |  |
| 5:27.138 | לחשוב |  |
| 5:27.139 | לבורא |  |
| 5:27.140 | ית | ' |
| 5:27.141 | כמה |  |
| 5:27.142 | פרצופים |  |
| 5:27.143 | חלוקים |  |
| 5:27.144 | בשמותן |  |
| 5:27.145 | ובאיכותן |  |
| 5:27.146 | ובפעולותיהם | . |
| 5:27.147 | ותמיד |  |
| 5:27.148 | מזדווגים | , |
| 5:27.149 | ועושים |  |
| 5:27.150 | תולדות |  |
| 5:27.151 | זה |  |
| 5:27.152 | מוליד |  |
| 5:27.153 | נשמות |  |
| 5:27.154 | למלאכים | , |
| 5:27.155 | וזה |  |
| 5:27.156 | מוליד |  |
| 5:27.157 | נשמות |  |
| 5:27.158 | לישראל |  |
| 5:27.159 | זהר |  |
| 5:27.160 | יתרו |  |
| 5:27.161 | דף |  |
| 5:27.162 | פ"ט | ) |
| 5:27.163 | וסטרא |  |
| 5:27.164 | אחרא |  |
| 5:27.165 | סמא"ל |  |
| 5:27.166 | בספירותיו |  |
| 5:27.167 | הנקרא |  |
| 5:27.168 | אל |  |
| 5:27.169 | אחד |  |
| 5:27.170 | ואדם |  |
| 5:27.171 | בליעל |  |
| 5:27.172 | מוליד |  |
| 5:27.173 | נשמות |  |
| 5:27.174 | לגוים | . |
| 5:27.175 | ואחר |  |
| 5:27.176 | שסיפר |  |
| 5:27.177 | לנו |  |
| 5:27.178 | כמה |  |
| 5:27.179 | פרצופים |  |
| 5:27.180 | שונים |  |
| 5:27.181 | הנקראים |  |
| 5:27.182 | בשמותיו |  |
| 5:27.183 | של |  |
| 5:27.184 | הקב"ה | . |
| 5:27.185 | מה |  |
| 5:27.186 | יתן |  |
| 5:27.187 | ומה |  |
| 5:27.188 | יוסיף |  |
| 5:27.189 | מה |  |
| 5:27.190 | שהוא |  |
| 5:27.191 | משקר |  |
| 5:27.192 | בלשונו |  |
| 5:27.193 | אחר |  |
| 5:27.194 | כך |  |
| 5:27.195 | ואומר |  |
| 5:27.196 | וכלא |  |
| 5:27.197 | חד | . |
| 5:27.198 | והוא |  |
| 5:27.199 | כבר |  |
| 5:27.200 | דיבר |  |
| 5:27.201 | בפיו |  |
| 5:27.202 | מה |  |
| 5:27.203 | שכבר |  |
| 5:27.204 | חישב |  |
| 5:27.205 | בלבו | . |
| 5:27.206 | וגם |  |
| 5:27.207 | הוציא |  |
| 5:27.208 | בשפתיו |  |
| 5:27.209 | וכתב |  |
| 5:27.210 | בקולמוסו |  |
| 5:27.211 | שהם |  |
| 5:27.212 | רבים | . |
| 5:27.213 | וכי |  |
| 5:27.214 | אתי |  |
| 5:27.215 | דיבור |  |
| 5:27.216 | לחודיה |  |
| 5:27.217 | ומבטל |  |
| 5:27.218 | מחשבה |  |
| 5:27.219 | ודבור |  |
| 5:27.220 | מעשה |  |
| 5:27.221 | שכתב |  |
| 5:27.222 | בספר | ? |
| 5:27.223 | אכן |  |
| 5:27.224 | האמת |  |
| 5:27.225 | הגמור |  |
| 5:27.226 | המקובל |  |
| 5:27.227 | אצל |  |
| 5:27.228 | אומותינו |  |
| 5:27.229 | מדור |  |
| 5:27.230 | דור | . |
| 5:27.231 | כי |  |
| 5:27.232 | באמת |  |
| 5:27.233 | ובתמים |  |
| 5:27.234 | אין |  |
| 5:27.235 | שיעור |  |
| 5:27.236 | וקצב |  |
| 5:27.237 | לאלהינו |  |
| 5:27.238 | יתברך |  |
| 5:27.239 | שמו |  |
| 5:27.240 | ולא |  |
| 5:27.241 | אורך |  |
| 5:27.242 | ולא |  |
| 5:27.243 | רוחב | , |
| 5:27.244 | ולא |  |
| 5:27.245 | דמות |  |
| 5:27.246 | ולא |  |
| 5:27.247 | צורת |  |
| 5:27.248 | גולם | . |
| 5:27.249 | וכל |  |
| 5:27.250 | שכן |  |
| 5:27.251 | שאין |  |
| 5:27.252 | לו |  |
| 5:27.253 | אשה | , |
| 5:27.254 | שתהיה |  |
| 5:27.255 | לו |  |
| 5:27.256 | לעזר |  |
| 5:27.257 | ולהועיל | , |
| 5:27.258 | שהוא |  |
| 5:27.259 | ית | ' |
| 5:27.260 | אינו |  |
| 5:27.261 | צריך |  |
| 5:27.262 | לעזר | . |
| 5:27.263 | כי |  |
| 5:27.264 | לגדולתו |  |
| 5:27.265 | אין |  |
| 5:27.266 | חקר |  |
| 5:27.267 | ולא |  |
| 5:27.268 | סוף | , |
| 5:27.269 | והוא |  |
| 5:27.270 | ראשון |  |
| 5:27.271 | ואין |  |
| 5:27.272 | לו |  |
| 5:27.273 | ראשית |  |
| 5:27.274 | לראשיתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

