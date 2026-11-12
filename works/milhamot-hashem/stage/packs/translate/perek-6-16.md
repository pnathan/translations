# Translation pack: milhamot-hashem perek-6-16 (Milchamot ha-Shem 6:25–6:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-16`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
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
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
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
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **ווי** ליה מאן דלא → "woe to him who does not".
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


## Milchamot ha-Shem 6:25

Hebrew (Wikisource): הכר נא קורא נעים עד היכן הגיעה דעתו ורעיונותיו של מחבר הנזכר לקרוא למפלצות שהם אברי המשגל של הפרצופים שהמציא ברעיונותיו בשם בית אלהים ובשם שער השמים, היש דברים של זרות יותר מזה? עוד בילקוט ובמדרש הגדול ידבחו לשדים לא אלוה, אלו עובדין לחמה וללבנה ולכוכבים ולמזלות ודברים שהן צורך העולם והנייה לעולם בהם לא הית הקנאה כפולה, אלא הן עובדים לדבר שאין אומות העולם מכירים אותם, חדשים מקרוב באו, שכל זמן שהיה אחד מן האומות רואה אותו אומר זה צלם יהודי ע"כ ורש"י בהומש מביא כל זה. וכ' הרלב"ג ז"ל יזבחו לשדים, לא דיים שעבדו לכוכבים ומזלות, אלא שעבדו לדמיונות כוזבים. והנה ידוע ומפורסם כי האומות הראשונים בימי אנוש טעו לעבוד לכוכבים ולמזלות ולשמש ולירח ושאר צבא השמים כמו שאמרו רז"ל במדרשיהם רבה וילקוט ומדרש בגדול והרמב"ם בתחלת הלכות ע"ז. ואחריהם קמו כתות אחרות שהאמינו בשני פועלים, פועל הטוב ופועל הרע, ואחריהם פשטה אמונת ישוע הנצרי מאמיני השלשה, אב ובן ורוח הקדש. ואחריהם עמדה הקבלה המזוייפת מאמיני עשר ספירות שלא כדעת בעל ס' יצירה, שפירשו העשר ספירות שהזכיר כהעולה על רוחם, ועשו אותם ברעיונותים פרצופים רבים השונים זה משה, זה ארוך וזה קצר וזה בינוני, והכל אחד לפי דעתם. והנה לפי דבריהם כללותם הם אחד, הנה האחד מחמשה פרצופים הכוללים אינו אחד שלם, רק חמישית אחד, כי רק החמשה פרצופים הכלוללים הם אחד. וכל פרצוף הוא רק חלק אחד מחמשה מהאלהים השלם, ואם כן איך נכווין בתפלתינו רק אל חלק אחד מחמשה שהוא קצר האפים (זעיר אנפין) לדעת רוב המקובלים, ולדעת הר' עז לאלהים שהעבודה והתפלות שאנו מתפללים אינו לזעיר אנפין אלא למלכא קדישא דכל קדישין. ועל זאת יתפלא כל מבין איך המקובלים בעצמם ההולכים בעקבות הזהר לא הסכימו לדעת אחת למי יתפללו, כי זה מתפלל וקורא לז"א ופילגשיו, וזה מתפלל אל מלכא קדישא דכל קדישין וזוגתו.

| id | word | punctuation after |
|---|---|---|
| 6:25.1 | הכר |  |
| 6:25.2 | נא |  |
| 6:25.3 | קורא |  |
| 6:25.4 | נעים |  |
| 6:25.5 | עד |  |
| 6:25.6 | היכן |  |
| 6:25.7 | הגיעה |  |
| 6:25.8 | דעתו |  |
| 6:25.9 | ורעיונותיו |  |
| 6:25.10 | של |  |
| 6:25.11 | מחבר |  |
| 6:25.12 | הנזכר |  |
| 6:25.13 | לקרוא |  |
| 6:25.14 | למפלצות |  |
| 6:25.15 | שהם |  |
| 6:25.16 | אברי |  |
| 6:25.17 | המשגל |  |
| 6:25.18 | של |  |
| 6:25.19 | הפרצופים |  |
| 6:25.20 | שהמציא |  |
| 6:25.21 | ברעיונותיו |  |
| 6:25.22 | בשם |  |
| 6:25.23 | בית |  |
| 6:25.24 | אלהים |  |
| 6:25.25 | ובשם |  |
| 6:25.26 | שער |  |
| 6:25.27 | השמים | , |
| 6:25.28 | היש |  |
| 6:25.29 | דברים |  |
| 6:25.30 | של |  |
| 6:25.31 | זרות |  |
| 6:25.32 | יותר |  |
| 6:25.33 | מזה | ? |
| 6:25.34 | עוד |  |
| 6:25.35 | בילקוט |  |
| 6:25.36 | ובמדרש |  |
| 6:25.37 | הגדול |  |
| 6:25.38 | ידבחו |  |
| 6:25.39 | לשדים |  |
| 6:25.40 | לא |  |
| 6:25.41 | אלוה | , |
| 6:25.42 | אלו |  |
| 6:25.43 | עובדין |  |
| 6:25.44 | לחמה |  |
| 6:25.45 | וללבנה |  |
| 6:25.46 | ולכוכבים |  |
| 6:25.47 | ולמזלות |  |
| 6:25.48 | ודברים |  |
| 6:25.49 | שהן |  |
| 6:25.50 | צורך |  |
| 6:25.51 | העולם |  |
| 6:25.52 | והנייה |  |
| 6:25.53 | לעולם |  |
| 6:25.54 | בהם |  |
| 6:25.55 | לא |  |
| 6:25.56 | הית |  |
| 6:25.57 | הקנאה |  |
| 6:25.58 | כפולה | , |
| 6:25.59 | אלא |  |
| 6:25.60 | הן |  |
| 6:25.61 | עובדים |  |
| 6:25.62 | לדבר |  |
| 6:25.63 | שאין |  |
| 6:25.64 | אומות |  |
| 6:25.65 | העולם |  |
| 6:25.66 | מכירים |  |
| 6:25.67 | אותם | , |
| 6:25.68 | חדשים |  |
| 6:25.69 | מקרוב |  |
| 6:25.70 | באו | , |
| 6:25.71 | שכל |  |
| 6:25.72 | זמן |  |
| 6:25.73 | שהיה |  |
| 6:25.74 | אחד |  |
| 6:25.75 | מן |  |
| 6:25.76 | האומות |  |
| 6:25.77 | רואה |  |
| 6:25.78 | אותו |  |
| 6:25.79 | אומר |  |
| 6:25.80 | זה |  |
| 6:25.81 | צלם |  |
| 6:25.82 | יהודי |  |
| 6:25.83 | ע"כ |  |
| 6:25.84 | ורש"י |  |
| 6:25.85 | בהומש |  |
| 6:25.86 | מביא |  |
| 6:25.87 | כל |  |
| 6:25.88 | זה | . |
| 6:25.89 | וכ | ' |
| 6:25.90 | הרלב"ג |  |
| 6:25.91 | ז"ל |  |
| 6:25.92 | יזבחו |  |
| 6:25.93 | לשדים | , |
| 6:25.94 | לא |  |
| 6:25.95 | דיים |  |
| 6:25.96 | שעבדו |  |
| 6:25.97 | לכוכבים |  |
| 6:25.98 | ומזלות | , |
| 6:25.99 | אלא |  |
| 6:25.100 | שעבדו |  |
| 6:25.101 | לדמיונות |  |
| 6:25.102 | כוזבים | . |
| 6:25.103 | והנה |  |
| 6:25.104 | ידוע |  |
| 6:25.105 | ומפורסם |  |
| 6:25.106 | כי |  |
| 6:25.107 | האומות |  |
| 6:25.108 | הראשונים |  |
| 6:25.109 | בימי |  |
| 6:25.110 | אנוש |  |
| 6:25.111 | טעו |  |
| 6:25.112 | לעבוד |  |
| 6:25.113 | לכוכבים |  |
| 6:25.114 | ולמזלות |  |
| 6:25.115 | ולשמש |  |
| 6:25.116 | ולירח |  |
| 6:25.117 | ושאר |  |
| 6:25.118 | צבא |  |
| 6:25.119 | השמים |  |
| 6:25.120 | כמו |  |
| 6:25.121 | שאמרו |  |
| 6:25.122 | רז"ל |  |
| 6:25.123 | במדרשיהם |  |
| 6:25.124 | רבה |  |
| 6:25.125 | וילקוט |  |
| 6:25.126 | ומדרש |  |
| 6:25.127 | בגדול |  |
| 6:25.128 | והרמב"ם |  |
| 6:25.129 | בתחלת |  |
| 6:25.130 | הלכות |  |
| 6:25.131 | ע"ז | . |
| 6:25.132 | ואחריהם |  |
| 6:25.133 | קמו |  |
| 6:25.134 | כתות |  |
| 6:25.135 | אחרות |  |
| 6:25.136 | שהאמינו |  |
| 6:25.137 | בשני |  |
| 6:25.138 | פועלים | , |
| 6:25.139 | פועל |  |
| 6:25.140 | הטוב |  |
| 6:25.141 | ופועל |  |
| 6:25.142 | הרע | , |
| 6:25.143 | ואחריהם |  |
| 6:25.144 | פשטה |  |
| 6:25.145 | אמונת |  |
| 6:25.146 | ישוע |  |
| 6:25.147 | הנצרי |  |
| 6:25.148 | מאמיני |  |
| 6:25.149 | השלשה | , |
| 6:25.150 | אב |  |
| 6:25.151 | ובן |  |
| 6:25.152 | ורוח |  |
| 6:25.153 | הקדש | . |
| 6:25.154 | ואחריהם |  |
| 6:25.155 | עמדה |  |
| 6:25.156 | הקבלה |  |
| 6:25.157 | המזוייפת |  |
| 6:25.158 | מאמיני |  |
| 6:25.159 | עשר |  |
| 6:25.160 | ספירות |  |
| 6:25.161 | שלא |  |
| 6:25.162 | כדעת |  |
| 6:25.163 | בעל |  |
| 6:25.164 | ס | ' |
| 6:25.165 | יצירה | , |
| 6:25.166 | שפירשו |  |
| 6:25.167 | העשר |  |
| 6:25.168 | ספירות |  |
| 6:25.169 | שהזכיר |  |
| 6:25.170 | כהעולה |  |
| 6:25.171 | על |  |
| 6:25.172 | רוחם | , |
| 6:25.173 | ועשו |  |
| 6:25.174 | אותם |  |
| 6:25.175 | ברעיונותים |  |
| 6:25.176 | פרצופים |  |
| 6:25.177 | רבים |  |
| 6:25.178 | השונים |  |
| 6:25.179 | זה |  |
| 6:25.180 | משה | , |
| 6:25.181 | זה |  |
| 6:25.182 | ארוך |  |
| 6:25.183 | וזה |  |
| 6:25.184 | קצר |  |
| 6:25.185 | וזה |  |
| 6:25.186 | בינוני | , |
| 6:25.187 | והכל |  |
| 6:25.188 | אחד |  |
| 6:25.189 | לפי |  |
| 6:25.190 | דעתם | . |
| 6:25.191 | והנה |  |
| 6:25.192 | לפי |  |
| 6:25.193 | דבריהם |  |
| 6:25.194 | כללותם |  |
| 6:25.195 | הם |  |
| 6:25.196 | אחד | , |
| 6:25.197 | הנה |  |
| 6:25.198 | האחד |  |
| 6:25.199 | מחמשה |  |
| 6:25.200 | פרצופים |  |
| 6:25.201 | הכוללים |  |
| 6:25.202 | אינו |  |
| 6:25.203 | אחד |  |
| 6:25.204 | שלם | , |
| 6:25.205 | רק |  |
| 6:25.206 | חמישית |  |
| 6:25.207 | אחד | , |
| 6:25.208 | כי |  |
| 6:25.209 | רק |  |
| 6:25.210 | החמשה |  |
| 6:25.211 | פרצופים |  |
| 6:25.212 | הכלוללים |  |
| 6:25.213 | הם |  |
| 6:25.214 | אחד | . |
| 6:25.215 | וכל |  |
| 6:25.216 | פרצוף |  |
| 6:25.217 | הוא |  |
| 6:25.218 | רק |  |
| 6:25.219 | חלק |  |
| 6:25.220 | אחד |  |
| 6:25.221 | מחמשה |  |
| 6:25.222 | מהאלהים |  |
| 6:25.223 | השלם | , |
| 6:25.224 | ואם |  |
| 6:25.225 | כן |  |
| 6:25.226 | איך |  |
| 6:25.227 | נכווין |  |
| 6:25.228 | בתפלתינו |  |
| 6:25.229 | רק |  |
| 6:25.230 | אל |  |
| 6:25.231 | חלק |  |
| 6:25.232 | אחד |  |
| 6:25.233 | מחמשה |  |
| 6:25.234 | שהוא |  |
| 6:25.235 | קצר |  |
| 6:25.236 | האפים |  |
| 6:25.237 | זעיר |  |
| 6:25.238 | אנפין | ) |
| 6:25.239 | לדעת |  |
| 6:25.240 | רוב |  |
| 6:25.241 | המקובלים | , |
| 6:25.242 | ולדעת |  |
| 6:25.243 | הר | ' |
| 6:25.244 | עז |  |
| 6:25.245 | לאלהים |  |
| 6:25.246 | שהעבודה |  |
| 6:25.247 | והתפלות |  |
| 6:25.248 | שאנו |  |
| 6:25.249 | מתפללים |  |
| 6:25.250 | אינו |  |
| 6:25.251 | לזעיר |  |
| 6:25.252 | אנפין |  |
| 6:25.253 | אלא |  |
| 6:25.254 | למלכא |  |
| 6:25.255 | קדישא |  |
| 6:25.256 | דכל |  |
| 6:25.257 | קדישין | . |
| 6:25.258 | ועל |  |
| 6:25.259 | זאת |  |
| 6:25.260 | יתפלא |  |
| 6:25.261 | כל |  |
| 6:25.262 | מבין |  |
| 6:25.263 | איך |  |
| 6:25.264 | המקובלים |  |
| 6:25.265 | בעצמם |  |
| 6:25.266 | ההולכים |  |
| 6:25.267 | בעקבות |  |
| 6:25.268 | הזהר |  |
| 6:25.269 | לא |  |
| 6:25.270 | הסכימו |  |
| 6:25.271 | לדעת |  |
| 6:25.272 | אחת |  |
| 6:25.273 | למי |  |
| 6:25.274 | יתפללו | , |
| 6:25.275 | כי |  |
| 6:25.276 | זה |  |
| 6:25.277 | מתפלל |  |
| 6:25.278 | וקורא |  |
| 6:25.279 | לז"א |  |
| 6:25.280 | ופילגשיו | , |
| 6:25.281 | וזה |  |
| 6:25.282 | מתפלל |  |
| 6:25.283 | אל |  |
| 6:25.284 | מלכא |  |
| 6:25.285 | קדישא |  |
| 6:25.286 | דכל |  |
| 6:25.287 | קדישין |  |
| 6:25.288 | וזוגתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

