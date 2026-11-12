# Translation pack: milhamot-hashem perek-3-05 (Milchamot ha-Shem 3:9–3:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-05`.

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
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ראב"ע** → "Rabbi Eleazar ben Azariah" in a Talmudic passage, "Ibn Ezra" for
  the commentator (note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בבחירתא** → "in the Bechirta" (note: "chosen").
## Terms

- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **מחויב** המציאות → "the Necessary Existent".
- **מין / מינות** → "heretic / heresy".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **רוע** לב → "wickedness of heart".
- **קטיגורין** → "accuser" (transliteration only in the note).
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **רחמנא ליצלן** → "the Merciful One save us".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 3:9

Hebrew (Wikisource): ט) ואתה ידידנו לא דייך שלא עשית אחת מכל אלה, אלא שהאמנת לדבריהם וקבלת שקריהם, בעד אחד ודיין אחד, ושלא בפני בעל דין, ובלא חקירה ובלא דרישה, ובלא בדיקה, ונהנית בדבריהם, וגם סייעת לחזק שקריהם ונרשם בדעתך כאמת לאמתו, עד שכתבת שאיני יכול להכחיש עוד וכו'. ולא שמת אל לבך שאין מקבלין עדות אלא בפני בעל דין, ושאתה מחויב לדרוש ולחקור אותם בשבע חקירות ובדרישות ובדיקות כמצווה עליך. כל שכן שעשו במחשך מעשיהם, והחליקו לשונם אליך בסתר דבריהם, למצוא עון לשנוא והית מחויב לשתקם בנזיפה. אין זה כי אם רוע לב שאתה נהנה ומשתוקק ברכילותם ומלשינותם. רחמנא ליצלן מהאי דעתא! ואפשר שגם אתה מוסיף משלך מלשינות שקר כדי להתכבד בקלון חבריך. ובפרק ערבי פסחים (קי"ה) גרסינן אמר רב ששת משום ראב"ע כל המספר לשון הרע, והמעיד עדות שקר בחברו ראוי להשליכו לכלבים שנאמר לכלב תשליכון אותו, וכתיב בהדיא לא תשא שמע שוא וקרי ביה לא תשיא, ועוד כמה דרשות מצינו לרז"ל בעונש המספר לשון הרע. ובמחזור ויטרי מסכת אבות על הא דתנן על תדון את חבירך עד שתגיע חמקומו כ' עש שתבא לידך אותה עבירה ותנצל ותשמט ממנה כמו שמצינו בירבעם בן נבט, שבשעה שנשא שלמה את בת פרעה הכניסה לו כל מיני זמר שבעולם ואמרה לו כך עושין לע"ז פלונית פעם אחת הלך לישן והעריב היום הלכה וקראה לחרשי העץ ולבונים, העושים במלאכת חרש וחושב, ועשו על מטתו כעין רקיע ובו חמה ולבנה וכוכבים ומזלות. ובבקר העיר משנתו לפתוח דלתות ההיכל, שהיו מפתחות ההיכל תחת מראשותיו ראה שהרקיע בהיר בכוכבים וצבא, כסבור עדיין לא פנה לילה, חזר לישן. וישן עוד ארבע שעות, עד שהבין והכיר מעצמו שאין הלילה ארוך כל כך. ועל אותה שעה שנינו עדות על תמיד של שחר שקרב בארבע שעות, בבחירתא. ובטרם הקיץ משנתו הלך ירבעם בן נבט על פתח ביתו, ויעבר קול, רשע, עד מתי תישן ותבטל תמיד של שחר, מה לך נרדם, קום קרא אל אלהיך! יצתה בת קול ואמרה לו רשע בן רשע עתיד אתה לבטל מישראל כמה וכמה קרבנות במזיד. ואתה מחייבו על השוגג כמזיד. ועל האונס כרצון! והיינו דכתיב "כדבר אפרים רתת, נשא הוא בישראל, ויאשם בבעל וימות" בדבר ירבעם שהוא משבט אפרים, רתת קטיגורין על שלמה, הוא ירבעם נשיא בישראל ומלך, ובת קול השיבתו שיאשם בבעל וימות עכ"ל. ואתה הגדלת והוספת לעשות רע יותר מירבעם בן נבט. כי אחרי אשר הבאתם אותנו אל בית הכלא במלשינות שקר אצל ממשלתינו הרוממה! הזעקת את כל העם לצאת אל בית עלמין, וצוית להם לעפר בעפר על ראשיהם על לא דבר. לא על מצוה קלה שבטלנו ולא על מצוה חמורה ח"ו כי אם על אשר הלכנו כמנהגי קדמונינו וקבענו למודינו במשנה ובתלמוד ורמב"ם וש"ע ללמוד וללמד, לשמור ולעשות. ולא למדנו עוד בספר מזויף ספר הזהר כמנהג שלכם.

| id | word | punctuation after |
|---|---|---|
| 3:9.1 | ט | ) |
| 3:9.2 | ואתה |  |
| 3:9.3 | ידידנו |  |
| 3:9.4 | לא |  |
| 3:9.5 | דייך |  |
| 3:9.6 | שלא |  |
| 3:9.7 | עשית |  |
| 3:9.8 | אחת |  |
| 3:9.9 | מכל |  |
| 3:9.10 | אלה | , |
| 3:9.11 | אלא |  |
| 3:9.12 | שהאמנת |  |
| 3:9.13 | לדבריהם |  |
| 3:9.14 | וקבלת |  |
| 3:9.15 | שקריהם | , |
| 3:9.16 | בעד |  |
| 3:9.17 | אחד |  |
| 3:9.18 | ודיין |  |
| 3:9.19 | אחד | , |
| 3:9.20 | ושלא |  |
| 3:9.21 | בפני |  |
| 3:9.22 | בעל |  |
| 3:9.23 | דין | , |
| 3:9.24 | ובלא |  |
| 3:9.25 | חקירה |  |
| 3:9.26 | ובלא |  |
| 3:9.27 | דרישה | , |
| 3:9.28 | ובלא |  |
| 3:9.29 | בדיקה | , |
| 3:9.30 | ונהנית |  |
| 3:9.31 | בדבריהם | , |
| 3:9.32 | וגם |  |
| 3:9.33 | סייעת |  |
| 3:9.34 | לחזק |  |
| 3:9.35 | שקריהם |  |
| 3:9.36 | ונרשם |  |
| 3:9.37 | בדעתך |  |
| 3:9.38 | כאמת |  |
| 3:9.39 | לאמתו | , |
| 3:9.40 | עד |  |
| 3:9.41 | שכתבת |  |
| 3:9.42 | שאיני |  |
| 3:9.43 | יכול |  |
| 3:9.44 | להכחיש |  |
| 3:9.45 | עוד |  |
| 3:9.46 | וכו | '. |
| 3:9.47 | ולא |  |
| 3:9.48 | שמת |  |
| 3:9.49 | אל |  |
| 3:9.50 | לבך |  |
| 3:9.51 | שאין |  |
| 3:9.52 | מקבלין |  |
| 3:9.53 | עדות |  |
| 3:9.54 | אלא |  |
| 3:9.55 | בפני |  |
| 3:9.56 | בעל |  |
| 3:9.57 | דין | , |
| 3:9.58 | ושאתה |  |
| 3:9.59 | מחויב |  |
| 3:9.60 | לדרוש |  |
| 3:9.61 | ולחקור |  |
| 3:9.62 | אותם |  |
| 3:9.63 | בשבע |  |
| 3:9.64 | חקירות |  |
| 3:9.65 | ובדרישות |  |
| 3:9.66 | ובדיקות |  |
| 3:9.67 | כמצווה |  |
| 3:9.68 | עליך | . |
| 3:9.69 | כל |  |
| 3:9.70 | שכן |  |
| 3:9.71 | שעשו |  |
| 3:9.72 | במחשך |  |
| 3:9.73 | מעשיהם | , |
| 3:9.74 | והחליקו |  |
| 3:9.75 | לשונם |  |
| 3:9.76 | אליך |  |
| 3:9.77 | בסתר |  |
| 3:9.78 | דבריהם | , |
| 3:9.79 | למצוא |  |
| 3:9.80 | עון |  |
| 3:9.81 | לשנוא |  |
| 3:9.82 | והית |  |
| 3:9.83 | מחויב |  |
| 3:9.84 | לשתקם |  |
| 3:9.85 | בנזיפה | . |
| 3:9.86 | אין |  |
| 3:9.87 | זה |  |
| 3:9.88 | כי |  |
| 3:9.89 | אם |  |
| 3:9.90 | רוע |  |
| 3:9.91 | לב |  |
| 3:9.92 | שאתה |  |
| 3:9.93 | נהנה |  |
| 3:9.94 | ומשתוקק |  |
| 3:9.95 | ברכילותם |  |
| 3:9.96 | ומלשינותם | . |
| 3:9.97 | רחמנא |  |
| 3:9.98 | ליצלן |  |
| 3:9.99 | מהאי |  |
| 3:9.100 | דעתא | ! |
| 3:9.101 | ואפשר |  |
| 3:9.102 | שגם |  |
| 3:9.103 | אתה |  |
| 3:9.104 | מוסיף |  |
| 3:9.105 | משלך |  |
| 3:9.106 | מלשינות |  |
| 3:9.107 | שקר |  |
| 3:9.108 | כדי |  |
| 3:9.109 | להתכבד |  |
| 3:9.110 | בקלון |  |
| 3:9.111 | חבריך | . |
| 3:9.112 | ובפרק |  |
| 3:9.113 | ערבי |  |
| 3:9.114 | פסחים |  |
| 3:9.115 | קי"ה | ) |
| 3:9.116 | גרסינן |  |
| 3:9.117 | אמר |  |
| 3:9.118 | רב |  |
| 3:9.119 | ששת |  |
| 3:9.120 | משום |  |
| 3:9.121 | ראב"ע |  |
| 3:9.122 | כל |  |
| 3:9.123 | המספר |  |
| 3:9.124 | לשון |  |
| 3:9.125 | הרע | , |
| 3:9.126 | והמעיד |  |
| 3:9.127 | עדות |  |
| 3:9.128 | שקר |  |
| 3:9.129 | בחברו |  |
| 3:9.130 | ראוי |  |
| 3:9.131 | להשליכו |  |
| 3:9.132 | לכלבים |  |
| 3:9.133 | שנאמר |  |
| 3:9.134 | לכלב |  |
| 3:9.135 | תשליכון |  |
| 3:9.136 | אותו | , |
| 3:9.137 | וכתיב |  |
| 3:9.138 | בהדיא |  |
| 3:9.139 | לא |  |
| 3:9.140 | תשא |  |
| 3:9.141 | שמע |  |
| 3:9.142 | שוא |  |
| 3:9.143 | וקרי |  |
| 3:9.144 | ביה |  |
| 3:9.145 | לא |  |
| 3:9.146 | תשיא | , |
| 3:9.147 | ועוד |  |
| 3:9.148 | כמה |  |
| 3:9.149 | דרשות |  |
| 3:9.150 | מצינו |  |
| 3:9.151 | לרז"ל |  |
| 3:9.152 | בעונש |  |
| 3:9.153 | המספר |  |
| 3:9.154 | לשון |  |
| 3:9.155 | הרע | . |
| 3:9.156 | ובמחזור |  |
| 3:9.157 | ויטרי |  |
| 3:9.158 | מסכת |  |
| 3:9.159 | אבות |  |
| 3:9.160 | על |  |
| 3:9.161 | הא |  |
| 3:9.162 | דתנן |  |
| 3:9.163 | על |  |
| 3:9.164 | תדון |  |
| 3:9.165 | את |  |
| 3:9.166 | חבירך |  |
| 3:9.167 | עד |  |
| 3:9.168 | שתגיע |  |
| 3:9.169 | חמקומו |  |
| 3:9.170 | כ | ' |
| 3:9.171 | עש |  |
| 3:9.172 | שתבא |  |
| 3:9.173 | לידך |  |
| 3:9.174 | אותה |  |
| 3:9.175 | עבירה |  |
| 3:9.176 | ותנצל |  |
| 3:9.177 | ותשמט |  |
| 3:9.178 | ממנה |  |
| 3:9.179 | כמו |  |
| 3:9.180 | שמצינו |  |
| 3:9.181 | בירבעם |  |
| 3:9.182 | בן |  |
| 3:9.183 | נבט | , |
| 3:9.184 | שבשעה |  |
| 3:9.185 | שנשא |  |
| 3:9.186 | שלמה |  |
| 3:9.187 | את |  |
| 3:9.188 | בת |  |
| 3:9.189 | פרעה |  |
| 3:9.190 | הכניסה |  |
| 3:9.191 | לו |  |
| 3:9.192 | כל |  |
| 3:9.193 | מיני |  |
| 3:9.194 | זמר |  |
| 3:9.195 | שבעולם |  |
| 3:9.196 | ואמרה |  |
| 3:9.197 | לו |  |
| 3:9.198 | כך |  |
| 3:9.199 | עושין |  |
| 3:9.200 | לע"ז |  |
| 3:9.201 | פלונית |  |
| 3:9.202 | פעם |  |
| 3:9.203 | אחת |  |
| 3:9.204 | הלך |  |
| 3:9.205 | לישן |  |
| 3:9.206 | והעריב |  |
| 3:9.207 | היום |  |
| 3:9.208 | הלכה |  |
| 3:9.209 | וקראה |  |
| 3:9.210 | לחרשי |  |
| 3:9.211 | העץ |  |
| 3:9.212 | ולבונים | , |
| 3:9.213 | העושים |  |
| 3:9.214 | במלאכת |  |
| 3:9.215 | חרש |  |
| 3:9.216 | וחושב | , |
| 3:9.217 | ועשו |  |
| 3:9.218 | על |  |
| 3:9.219 | מטתו |  |
| 3:9.220 | כעין |  |
| 3:9.221 | רקיע |  |
| 3:9.222 | ובו |  |
| 3:9.223 | חמה |  |
| 3:9.224 | ולבנה |  |
| 3:9.225 | וכוכבים |  |
| 3:9.226 | ומזלות | . |
| 3:9.227 | ובבקר |  |
| 3:9.228 | העיר |  |
| 3:9.229 | משנתו |  |
| 3:9.230 | לפתוח |  |
| 3:9.231 | דלתות |  |
| 3:9.232 | ההיכל | , |
| 3:9.233 | שהיו |  |
| 3:9.234 | מפתחות |  |
| 3:9.235 | ההיכל |  |
| 3:9.236 | תחת |  |
| 3:9.237 | מראשותיו |  |
| 3:9.238 | ראה |  |
| 3:9.239 | שהרקיע |  |
| 3:9.240 | בהיר |  |
| 3:9.241 | בכוכבים |  |
| 3:9.242 | וצבא | , |
| 3:9.243 | כסבור |  |
| 3:9.244 | עדיין |  |
| 3:9.245 | לא |  |
| 3:9.246 | פנה |  |
| 3:9.247 | לילה | , |
| 3:9.248 | חזר |  |
| 3:9.249 | לישן | . |
| 3:9.250 | וישן |  |
| 3:9.251 | עוד |  |
| 3:9.252 | ארבע |  |
| 3:9.253 | שעות | , |
| 3:9.254 | עד |  |
| 3:9.255 | שהבין |  |
| 3:9.256 | והכיר |  |
| 3:9.257 | מעצמו |  |
| 3:9.258 | שאין |  |
| 3:9.259 | הלילה |  |
| 3:9.260 | ארוך |  |
| 3:9.261 | כל |  |
| 3:9.262 | כך | . |
| 3:9.263 | ועל |  |
| 3:9.264 | אותה |  |
| 3:9.265 | שעה |  |
| 3:9.266 | שנינו |  |
| 3:9.267 | עדות |  |
| 3:9.268 | על |  |
| 3:9.269 | תמיד |  |
| 3:9.270 | של |  |
| 3:9.271 | שחר |  |
| 3:9.272 | שקרב |  |
| 3:9.273 | בארבע |  |
| 3:9.274 | שעות | , |
| 3:9.275 | בבחירתא | . |
| 3:9.276 | ובטרם |  |
| 3:9.277 | הקיץ |  |
| 3:9.278 | משנתו |  |
| 3:9.279 | הלך |  |
| 3:9.280 | ירבעם |  |
| 3:9.281 | בן |  |
| 3:9.282 | נבט |  |
| 3:9.283 | על |  |
| 3:9.284 | פתח |  |
| 3:9.285 | ביתו | , |
| 3:9.286 | ויעבר |  |
| 3:9.287 | קול | , |
| 3:9.288 | רשע | , |
| 3:9.289 | עד |  |
| 3:9.290 | מתי |  |
| 3:9.291 | תישן |  |
| 3:9.292 | ותבטל |  |
| 3:9.293 | תמיד |  |
| 3:9.294 | של |  |
| 3:9.295 | שחר | , |
| 3:9.296 | מה |  |
| 3:9.297 | לך |  |
| 3:9.298 | נרדם | , |
| 3:9.299 | קום |  |
| 3:9.300 | קרא |  |
| 3:9.301 | אל |  |
| 3:9.302 | אלהיך | ! |
| 3:9.303 | יצתה |  |
| 3:9.304 | בת |  |
| 3:9.305 | קול |  |
| 3:9.306 | ואמרה |  |
| 3:9.307 | לו |  |
| 3:9.308 | רשע |  |
| 3:9.309 | בן |  |
| 3:9.310 | רשע |  |
| 3:9.311 | עתיד |  |
| 3:9.312 | אתה |  |
| 3:9.313 | לבטל |  |
| 3:9.314 | מישראל |  |
| 3:9.315 | כמה |  |
| 3:9.316 | וכמה |  |
| 3:9.317 | קרבנות |  |
| 3:9.318 | במזיד | . |
| 3:9.319 | ואתה |  |
| 3:9.320 | מחייבו |  |
| 3:9.321 | על |  |
| 3:9.322 | השוגג |  |
| 3:9.323 | כמזיד | . |
| 3:9.324 | ועל |  |
| 3:9.325 | האונס |  |
| 3:9.326 | כרצון | ! |
| 3:9.327 | והיינו |  |
| 3:9.328 | דכתיב |  |
| 3:9.329 | כדבר |  |
| 3:9.330 | אפרים |  |
| 3:9.331 | רתת | , |
| 3:9.332 | נשא |  |
| 3:9.333 | הוא |  |
| 3:9.334 | בישראל | , |
| 3:9.335 | ויאשם |  |
| 3:9.336 | בבעל |  |
| 3:9.337 | וימות | " |
| 3:9.338 | בדבר |  |
| 3:9.339 | ירבעם |  |
| 3:9.340 | שהוא |  |
| 3:9.341 | משבט |  |
| 3:9.342 | אפרים | , |
| 3:9.343 | רתת |  |
| 3:9.344 | קטיגורין |  |
| 3:9.345 | על |  |
| 3:9.346 | שלמה | , |
| 3:9.347 | הוא |  |
| 3:9.348 | ירבעם |  |
| 3:9.349 | נשיא |  |
| 3:9.350 | בישראל |  |
| 3:9.351 | ומלך | , |
| 3:9.352 | ובת |  |
| 3:9.353 | קול |  |
| 3:9.354 | השיבתו |  |
| 3:9.355 | שיאשם |  |
| 3:9.356 | בבעל |  |
| 3:9.357 | וימות |  |
| 3:9.358 | עכ"ל | . |
| 3:9.359 | ואתה |  |
| 3:9.360 | הגדלת |  |
| 3:9.361 | והוספת |  |
| 3:9.362 | לעשות |  |
| 3:9.363 | רע |  |
| 3:9.364 | יותר |  |
| 3:9.365 | מירבעם |  |
| 3:9.366 | בן |  |
| 3:9.367 | נבט | . |
| 3:9.368 | כי |  |
| 3:9.369 | אחרי |  |
| 3:9.370 | אשר |  |
| 3:9.371 | הבאתם |  |
| 3:9.372 | אותנו |  |
| 3:9.373 | אל |  |
| 3:9.374 | בית |  |
| 3:9.375 | הכלא |  |
| 3:9.376 | במלשינות |  |
| 3:9.377 | שקר |  |
| 3:9.378 | אצל |  |
| 3:9.379 | ממשלתינו |  |
| 3:9.380 | הרוממה | ! |
| 3:9.381 | הזעקת |  |
| 3:9.382 | את |  |
| 3:9.383 | כל |  |
| 3:9.384 | העם |  |
| 3:9.385 | לצאת |  |
| 3:9.386 | אל |  |
| 3:9.387 | בית |  |
| 3:9.388 | עלמין | , |
| 3:9.389 | וצוית |  |
| 3:9.390 | להם |  |
| 3:9.391 | לעפר |  |
| 3:9.392 | בעפר |  |
| 3:9.393 | על |  |
| 3:9.394 | ראשיהם |  |
| 3:9.395 | על |  |
| 3:9.396 | לא |  |
| 3:9.397 | דבר | . |
| 3:9.398 | לא |  |
| 3:9.399 | על |  |
| 3:9.400 | מצוה |  |
| 3:9.401 | קלה |  |
| 3:9.402 | שבטלנו |  |
| 3:9.403 | ולא |  |
| 3:9.404 | על |  |
| 3:9.405 | מצוה |  |
| 3:9.406 | חמורה |  |
| 3:9.407 | ח"ו |  |
| 3:9.408 | כי |  |
| 3:9.409 | אם |  |
| 3:9.410 | על |  |
| 3:9.411 | אשר |  |
| 3:9.412 | הלכנו |  |
| 3:9.413 | כמנהגי |  |
| 3:9.414 | קדמונינו |  |
| 3:9.415 | וקבענו |  |
| 3:9.416 | למודינו |  |
| 3:9.417 | במשנה |  |
| 3:9.418 | ובתלמוד |  |
| 3:9.419 | ורמב"ם |  |
| 3:9.420 | וש"ע |  |
| 3:9.421 | ללמוד |  |
| 3:9.422 | וללמד | , |
| 3:9.423 | לשמור |  |
| 3:9.424 | ולעשות | . |
| 3:9.425 | ולא |  |
| 3:9.426 | למדנו |  |
| 3:9.427 | עוד |  |
| 3:9.428 | בספר |  |
| 3:9.429 | מזויף |  |
| 3:9.430 | ספר |  |
| 3:9.431 | הזהר |  |
| 3:9.432 | כמנהג |  |
| 3:9.433 | שלכם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

