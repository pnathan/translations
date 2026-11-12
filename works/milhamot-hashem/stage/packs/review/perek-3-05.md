# Review pack: milhamot-hashem perek-3-05 (Milchamot ha-Shem 3:9–3:9)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-05`.

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


### Translation under review (3:9)

English: 9) And you, our friend, it was not enough for you that you did not do any one of all these but you believed their words and accepted their lies on the testimony of one witness and one judge and not in the presence of the litigant without inquiry, without investigation, without examination and you enjoyed their words and you also helped to strengthen their lies and it was fixed in your mind as the truth itself until you wrote that I am no longer able to deny it, etc. And you did not take to heart that testimony is accepted only in the presence of the litigant and that you are obliged to inquire into them with seven inquiries, investigations and examinations, as you are commanded All the more so, since they did their deeds in the dark and smoothed their tongues to you in secret with their words to find a fault to hate and you were obliged to silence them with a rebuke. This is nothing but a wicked heart: that you enjoy and long for their gossip and their informing. the Merciful One save us from this opinion! And it is possible that you too add of your own false informing, in order to honor yourself at the disgrace of your friends. And in the chapter Arvei Pesachim (115) we read: "Rav Sheshet said in the name of R. A.: whoever tells evil talk and whoever bears false witness against his fellow deserves to be thrown to the dogs, as it is said: 'to the dog you shall throw it' and it is written explicitly: 'you shall not carry a false report'; and it is read in it: 'you shall not cause to bear'" And we find several more homilies of our Sages of blessed memory on the punishment of one who tells evil talk. "And in Mahzor Vitry, tractate Avot, on what we learned: 'do not judge your fellow until you reach his place' as he writes there: that the same transgression may come into your hand and you may be saved and escape from it as we find with Jeroboam son of Nebat that when Solomon married the daughter of Pharaoh, she brought in to him every kind of song in the world and said to him: 'So they do for such-and-such idolatry' One time, he went to sleep, and the day grew late she went and called the wood-carvers and the builders, who work at carving and weaving and they made over his bed the likeness of a firmament, and in it the sun, the moon, the stars and the constellations In the morning he woke from his sleep to open the doors of the Temple, for the keys of the Temple were under his pillow he saw that the firmament was bright with stars and the host thinking that the night had not yet turned he went back to sleep And he slept four more hours, until he understood and recognised for himself that the night was not so long And concerning that hour we learned a testimony about the morning daily offering, which is offered at four hours, in the Bechirta. Before he woke from his sleep, Jeroboam son of Nebat went to the door of his house and a voice went out: 'Wicked one how long will you sleep and neglect the morning daily offering What is it to you, sleeper? Arise, call upon your God!' A bat kol came out and said to him: 'Wicked son of a wicked man, you are destined to abolish from Israel many, many offerings wilfully. And you hold him liable for the inadvertent as for the wilful. And for the forced one as for the voluntary!' And this is what is written: 'When Ephraim spoke there was trembling; he was exalted in Israel, but he incurred guilt through Baal and died' This is said of Jeroboam, who was of the tribe of Ephraim, the trembling was the accuser against Solomon, he, Jeroboam, was prince in Israel and king and the bat kol answered him that he would become guilty through Baal and die." Thus far his words. And you have magnified and added to doing evil more than Jeroboam son of Nebat. For after you brought us to the prison by false informing before our exalted government! You called out all the people to go out to the cemetery, and commanded them to dust themselves with dust upon their heads for no reason. Not for a light commandment that we neglected, and not for a weighty commandment, God forbid, but because we walked in the custom of our forefathers and set our learning in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch to learn and to teach, to keep and to do. And we did not study further from the forged book, the book the Zohar, as is your custom.

Literal: 9) and you our friend not enough for you that not you did one of all these but that you believed to their words and you accepted their lies on the testimony of one and a judge one and not in the presence of the litigant of the case and without inquiry and without investigation and without examination and you enjoyed in their words and also you helped to strengthen their lies and was fixed in your mind as truth to its truth until that you wrote that I am able to deny any longer etc. and not you put to your heart that there is not they accept testimony but only in the presence of a litigant of the case and that you are obliged to inquire and to investigate them with seven inquiries and with investigations and with examinations as commanded of you all the more so that they did in the dark their deeds and smoothed their tongue to you in secret with their words to find a fault to hate and you were obliged to silence them with a rebuke it is not this that except wickedness of heart that you enjoy and long for their gossip and their informing the Merciful One save us from this opinion and it is possible that also you add of your own informing false in order to honor yourself at the disgrace of your friends and in the chapter Arvei Pesachim (folio 115) we read said Rav Sheshet in the name of R. A. whoever tells evil talk and whoever testifies testimony false against his fellow is fit to cast him to the dogs as it is said to the dog you shall throw it and it is written explicitly not you shall carry report false and it is read in it not you shall cause to bear and further several homilies we find our Sages of blessed memory on the punishment of the one who tells evil talk and in Mahzor Vitry tractate Avot on that which we learned on (al: do not) judge - your fellow until you reach his place as he writes there that may come into your hand that same transgression and be saved and escape from it as we find with Jeroboam son of Nebat that when he married Solomon - daughter of Pharaoh she brought in to him every kind of song in the world and she said to him thus they do for idolatry such-and-such once time he went to sleep and became late the day she went and called the wood-carvers of wood and the builders who do the craft of carving and weaving and they made on his bed the likeness of a firmament and in it the sun and the moon and the stars and the constellations and in the morning he woke from his sleep to open the doors of the Temple for there were keys of the Temple under his pillow he saw that the firmament was bright with stars and the host thinking still not had turned night he returned to sleep and he slept more four hours until he understood and recognised for himself that there is not the night long so thus and concerning that hour we learned testimony about the daily offering of morning which is offered at four hours in the Bechirta and before he woke from his sleep went Jeroboam son of Nebat to the door of his house and went out a voice wicked one until when you sleep and neglect the daily offering of morning what to you sleeper arise call upon your God came out bat (a heavenly) voice and said to him wicked son of wicked destined you to nullify from Israel many and many offerings wilfully and you hold him liable for the inadvertent as the wilful and for the forced as the voluntary and this is what is written as the speaking of Ephraim trembling he was exalted he in Israel and became guilty through Baal and died in the matter of Jeroboam who is of the tribe of Ephraim rattat (trembled) katigorin (accuser) against Solomon he Jeroboam prince in Israel and king and a bat kol answered him that he would become guilty through Baal and die thus far his words and you have magnified and added to do evil more than Jeroboam son of Nebat for after that you brought us to the house of the prison by false informing before our government exalted you called out - all the people to go out to the house of the cemetery and you commanded them to dust themselves with dust on their heads for no thing not for a commandment light that we neglected and not for a commandment weighty God forbid but only for that we walked in the custom of our forefathers and we set our learning in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch to learn and to teach to keep and to do and not we studied further from the book forged the book the Zohar as is the custom of you

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:9.1 | 9) |  |  | printed list marker "9)", kept as the author's own |
| 3:9.2 | and you |  |  |  |
| 3:9.3 | our friend |  |  | ידידנו: "our friend", the ironic form of address, kept |
| 3:9.4 | not |  |  |  |
| 3:9.5 | enough for you |  |  | לא דייך: "it was not enough for you"; the idiom governs the clause to 3:9.10 |
| 3:9.6 | that not |  |  |  |
| 3:9.7 | you did |  |  |  |
| 3:9.8 | one |  |  |  |
| 3:9.9 | of all |  |  |  |
| 3:9.10 | these |  |  |  |
| 3:9.11 | but |  |  |  |
| 3:9.12 | that you believed |  |  |  |
| 3:9.13 | to their words |  |  |  |
| 3:9.14 | and you accepted |  |  |  |
| 3:9.15 | their lies |  |  |  |
| 3:9.16 | on the testimony of |  |  |  |
| 3:9.17 | one |  |  |  |
| 3:9.18 | and a judge |  |  |  |
| 3:9.19 | one |  |  |  |
| 3:9.20 | and not |  |  |  |
| 3:9.21 | in the presence of |  |  |  |
| 3:9.22 | the litigant |  |  |  |
| 3:9.23 | of the case |  |  |  |
| 3:9.24 | and without |  |  |  |
| 3:9.25 | inquiry |  |  |  |
| 3:9.26 | and without |  |  |  |
| 3:9.27 | investigation |  |  |  |
| 3:9.28 | and without |  |  |  |
| 3:9.29 | examination |  |  |  |
| 3:9.30 | and you enjoyed |  |  |  |
| 3:9.31 | in their words |  |  |  |
| 3:9.32 | and also |  |  |  |
| 3:9.33 | you helped |  |  |  |
| 3:9.34 | to strengthen |  |  |  |
| 3:9.35 | their lies |  |  |  |
| 3:9.36 | and was fixed |  |  |  |
| 3:9.37 | in your mind |  |  |  |
| 3:9.38 | as truth |  |  |  |
| 3:9.39 | to its truth |  |  |  |
| 3:9.40 | until |  |  |  |
| 3:9.41 | that you wrote |  |  |  |
| 3:9.42 | that I |  |  |  |
| 3:9.43 | am able |  |  |  |
| 3:9.44 | to deny |  |  |  |
| 3:9.45 | any longer |  |  |  |
| 3:9.46 | etc. |  |  | וכו': marks a truncated quotation; kept as "etc." |
| 3:9.47 | and not |  |  |  |
| 3:9.48 | you put |  |  |  |
| 3:9.49 | to |  |  |  |
| 3:9.50 | your heart |  |  |  |
| 3:9.51 | that there is not |  |  |  |
| 3:9.52 | they accept |  |  |  |
| 3:9.53 | testimony |  |  |  |
| 3:9.54 | but only |  |  |  |
| 3:9.55 | in the presence of |  |  |  |
| 3:9.56 | a litigant |  |  |  |
| 3:9.57 | of the case |  |  |  |
| 3:9.58 | and that you |  |  |  |
| 3:9.59 | are obliged |  |  |  |
| 3:9.60 | to inquire |  |  |  |
| 3:9.61 | and to investigate |  |  |  |
| 3:9.62 | them |  |  |  |
| 3:9.63 | with seven |  |  |  |
| 3:9.64 | inquiries |  |  |  |
| 3:9.65 | and with investigations |  |  |  |
| 3:9.66 | and with examinations |  |  |  |
| 3:9.67 | as commanded |  |  |  |
| 3:9.68 | of you |  |  |  |
| 3:9.69 | all |  |  |  |
| 3:9.70 | the more so |  |  |  |
| 3:9.71 | that they did |  |  |  |
| 3:9.72 | in the dark |  |  |  |
| 3:9.73 | their deeds |  |  |  |
| 3:9.74 | and smoothed |  |  |  |
| 3:9.75 | their tongue |  |  |  |
| 3:9.76 | to you |  |  |  |
| 3:9.77 | in secret |  |  |  |
| 3:9.78 | with their words |  |  |  |
| 3:9.79 | to find |  |  |  |
| 3:9.80 | a fault |  |  |  |
| 3:9.81 | to hate |  |  |  |
| 3:9.82 | and you were |  |  | printed "והית" for "והיית" (and you were); the defective spelling is kept as printed |
| 3:9.83 | obliged |  |  |  |
| 3:9.84 | to silence them |  |  |  |
| 3:9.85 | with a rebuke |  |  |  |
| 3:9.86 | it is not |  |  |  |
| 3:9.87 | this |  |  |  |
| 3:9.88 | that |  |  |  |
| 3:9.89 | except |  |  |  |
| 3:9.90 | wickedness |  |  |  |
| 3:9.91 | of heart |  |  |  |
| 3:9.92 | that you |  |  |  |
| 3:9.93 | enjoy |  |  |  |
| 3:9.94 | and long for |  |  |  |
| 3:9.95 | their gossip |  |  |  |
| 3:9.96 | and their informing |  |  |  |
| 3:9.97 | the Merciful One |  |  |  |
| 3:9.98 | save us |  |  |  |
| 3:9.99 | from this |  |  |  |
| 3:9.100 | opinion |  |  |  |
| 3:9.101 | and it is possible |  |  |  |
| 3:9.102 | that also |  |  |  |
| 3:9.103 | you |  |  |  |
| 3:9.104 | add |  |  |  |
| 3:9.105 | of your own |  |  |  |
| 3:9.106 | informing |  |  |  |
| 3:9.107 | false |  |  |  |
| 3:9.108 | in order |  |  |  |
| 3:9.109 | to honor yourself |  |  |  |
| 3:9.110 | at the disgrace |  |  |  |
| 3:9.111 | of your friends |  |  |  |
| 3:9.112 | and in the chapter |  |  |  |
| 3:9.113 | Arvei |  |  |  |
| 3:9.114 | Pesachim |  |  |  |
| 3:9.115 | (folio 115) |  |  | folio 115 of Pesachim, the reference as the author gives it (קי"ה) |
| 3:9.116 | we read |  |  |  |
| 3:9.117 | said |  |  |  |
| 3:9.118 | Rav |  |  |  |
| 3:9.119 | Sheshet |  |  |  |
| 3:9.120 | in the name of |  |  |  |
| 3:9.121 | R. A. |  |  | ראב"ע: not expanded in the printed text; read as "R. A."; the full name is uncertain |
| 3:9.122 | whoever |  |  |  |
| 3:9.123 | tells |  |  |  |
| 3:9.124 | evil |  |  |  |
| 3:9.125 | talk |  |  |  |
| 3:9.126 | and whoever testifies |  |  |  |
| 3:9.127 | testimony |  |  |  |
| 3:9.128 | false |  |  |  |
| 3:9.129 | against his fellow |  |  |  |
| 3:9.130 | is fit |  |  |  |
| 3:9.131 | to cast him |  |  |  |
| 3:9.132 | to the dogs |  |  |  |
| 3:9.133 | as it is said |  |  |  |
| 3:9.134 | to the dog |  |  |  |
| 3:9.135 | you shall throw |  |  |  |
| 3:9.136 | it |  |  |  |
| 3:9.137 | and it is written |  |  |  |
| 3:9.138 | explicitly |  |  |  |
| 3:9.139 | not |  |  |  |
| 3:9.140 | you shall carry |  |  |  |
| 3:9.141 | report |  |  |  |
| 3:9.142 | false |  |  |  |
| 3:9.143 | and it is read |  |  |  |
| 3:9.144 | in it |  |  |  |
| 3:9.145 | not |  |  |  |
| 3:9.146 | you shall cause to bear |  |  | תשיא: the reading the Talmud gives ("וקרי ביה"), set against the written לא תשא of Exod 23:1 |
| 3:9.147 | and further |  |  |  |
| 3:9.148 | several |  |  |  |
| 3:9.149 | homilies |  |  |  |
| 3:9.150 | we find |  |  |  |
| 3:9.151 | our Sages of blessed memory |  |  | לרז"ל: "our Sages of blessed memory" |
| 3:9.152 | on the punishment |  |  |  |
| 3:9.153 | of the one who tells |  |  |  |
| 3:9.154 | evil |  |  |  |
| 3:9.155 | talk |  |  |  |
| 3:9.156 | and in Mahzor |  |  |  |
| 3:9.157 | Vitry |  |  |  |
| 3:9.158 | tractate |  |  |  |
| 3:9.159 | Avot |  |  |  |
| 3:9.160 | on |  |  |  |
| 3:9.161 | that |  |  |  |
| 3:9.162 | which we learned |  |  |  |
| 3:9.163 | on (al: do not) |  |  | על: the Mishnah's "al" (אל), the negative of "do not judge" |
| 3:9.164 | judge |  |  |  |
| 3:9.165 | - |  |  |  |
| 3:9.166 | your fellow |  |  |  |
| 3:9.167 | until |  |  |  |
| 3:9.168 | you reach |  |  |  |
| 3:9.169 | his place |  |  |  |
| 3:9.170 | as he writes |  |  | כ' ע"ש: "as he writes there"; the expansion is uncertain |
| 3:9.171 | there |  |  |  |
| 3:9.172 | that may come |  |  |  |
| 3:9.173 | into your hand |  |  |  |
| 3:9.174 | that same |  |  |  |
| 3:9.175 | transgression |  |  |  |
| 3:9.176 | and be saved |  |  |  |
| 3:9.177 | and escape |  |  |  |
| 3:9.178 | from it |  |  |  |
| 3:9.179 | as |  |  |  |
| 3:9.180 | we find |  |  |  |
| 3:9.181 | with Jeroboam |  |  |  |
| 3:9.182 | son of |  |  |  |
| 3:9.183 | Nebat |  |  |  |
| 3:9.184 | that when |  |  |  |
| 3:9.185 | he married |  |  |  |
| 3:9.186 | Solomon |  |  |  |
| 3:9.187 | - |  |  |  |
| 3:9.188 | daughter |  |  |  |
| 3:9.189 | of Pharaoh |  |  |  |
| 3:9.190 | she brought in |  |  |  |
| 3:9.191 | to him |  |  |  |
| 3:9.192 | every |  |  |  |
| 3:9.193 | kind of |  |  |  |
| 3:9.194 | song |  |  |  |
| 3:9.195 | in the world |  |  |  |
| 3:9.196 | and she said |  |  |  |
| 3:9.197 | to him |  |  |  |
| 3:9.198 | thus |  |  |  |
| 3:9.199 | they do |  |  |  |
| 3:9.200 | for idolatry |  |  | לע"ז: abbreviation of עבודה זרה, "idolatry" |
| 3:9.201 | such-and-such |  |  |  |
| 3:9.202 | once |  |  |  |
| 3:9.203 | time |  |  |  |
| 3:9.204 | he went |  |  |  |
| 3:9.205 | to sleep |  |  |  |
| 3:9.206 | and became late |  |  |  |
| 3:9.207 | the day |  |  |  |
| 3:9.208 | she went |  |  |  |
| 3:9.209 | and called |  |  |  |
| 3:9.210 | the wood-carvers |  |  |  |
| 3:9.211 | of wood |  |  |  |
| 3:9.212 | and the builders |  |  |  |
| 3:9.213 | who do |  |  |  |
| 3:9.214 | the craft |  |  |  |
| 3:9.215 | of carving |  |  |  |
| 3:9.216 | and weaving |  |  |  |
| 3:9.217 | and they made |  |  |  |
| 3:9.218 | on |  |  |  |
| 3:9.219 | his bed |  |  |  |
| 3:9.220 | the likeness of |  |  |  |
| 3:9.221 | a firmament |  |  |  |
| 3:9.222 | and in it |  |  |  |
| 3:9.223 | the sun |  |  |  |
| 3:9.224 | and the moon |  |  |  |
| 3:9.225 | and the stars |  |  |  |
| 3:9.226 | and the constellations |  |  |  |
| 3:9.227 | and in the morning |  |  |  |
| 3:9.228 | he woke |  |  |  |
| 3:9.229 | from his sleep |  |  |  |
| 3:9.230 | to open |  |  |  |
| 3:9.231 | the doors |  |  |  |
| 3:9.232 | of the Temple |  |  |  |
| 3:9.233 | for there were |  |  |  |
| 3:9.234 | keys |  |  |  |
| 3:9.235 | of the Temple |  |  |  |
| 3:9.236 | under |  |  |  |
| 3:9.237 | his pillow |  |  |  |
| 3:9.238 | he saw |  |  |  |
| 3:9.239 | that the firmament |  |  |  |
| 3:9.240 | was bright |  |  |  |
| 3:9.241 | with stars |  |  |  |
| 3:9.242 | and the host |  |  |  |
| 3:9.243 | thinking |  |  |  |
| 3:9.244 | still |  |  |  |
| 3:9.245 | not |  |  |  |
| 3:9.246 | had turned |  |  |  |
| 3:9.247 | night |  |  |  |
| 3:9.248 | he returned |  |  |  |
| 3:9.249 | to sleep |  |  |  |
| 3:9.250 | and he slept |  |  |  |
| 3:9.251 | more |  |  |  |
| 3:9.252 | four |  |  |  |
| 3:9.253 | hours |  |  |  |
| 3:9.254 | until |  |  |  |
| 3:9.255 | he understood |  |  |  |
| 3:9.256 | and recognised |  |  |  |
| 3:9.257 | for himself |  |  |  |
| 3:9.258 | that there is not |  |  |  |
| 3:9.259 | the night |  |  |  |
| 3:9.260 | long |  |  |  |
| 3:9.261 | so |  |  |  |
| 3:9.262 | thus |  |  |  |
| 3:9.263 | and concerning |  |  |  |
| 3:9.264 | that |  |  |  |
| 3:9.265 | hour |  |  |  |
| 3:9.266 | we learned |  |  |  |
| 3:9.267 | testimony |  |  |  |
| 3:9.268 | about |  |  |  |
| 3:9.269 | the daily offering |  |  |  |
| 3:9.270 | of |  |  |  |
| 3:9.271 | morning |  |  |  |
| 3:9.272 | which is offered |  |  |  |
| 3:9.273 | at four |  |  |  |
| 3:9.274 | hours |  |  |  |
| 3:9.275 | in the Bechirta |  |  | בבחירתא: transliterated; the source is not named in the pack |
| 3:9.276 | and before |  |  |  |
| 3:9.277 | he woke |  |  |  |
| 3:9.278 | from his sleep |  |  |  |
| 3:9.279 | went |  |  |  |
| 3:9.280 | Jeroboam |  |  |  |
| 3:9.281 | son of |  |  |  |
| 3:9.282 | Nebat |  |  |  |
| 3:9.283 | to |  |  |  |
| 3:9.284 | the door |  |  |  |
| 3:9.285 | of his house |  |  |  |
| 3:9.286 | and went out |  |  |  |
| 3:9.287 | a voice |  |  |  |
| 3:9.288 | wicked one |  |  |  |
| 3:9.289 | until |  |  |  |
| 3:9.290 | when |  |  |  |
| 3:9.291 | you sleep |  |  |  |
| 3:9.292 | and neglect |  |  |  |
| 3:9.293 | the daily offering |  |  |  |
| 3:9.294 | of |  |  |  |
| 3:9.295 | morning |  |  |  |
| 3:9.296 | what |  |  |  |
| 3:9.297 | to you |  |  |  |
| 3:9.298 | sleeper |  |  |  |
| 3:9.299 | arise |  |  |  |
| 3:9.300 | call |  |  |  |
| 3:9.301 | upon |  |  |  |
| 3:9.302 | your God |  |  | אלהיך: God of Israel, in the voice's words |
| 3:9.303 | came out |  |  |  |
| 3:9.304 | bat (a heavenly) |  |  |  |
| 3:9.305 | voice |  |  |  |
| 3:9.306 | and said |  |  |  |
| 3:9.307 | to him |  |  |  |
| 3:9.308 | wicked |  |  |  |
| 3:9.309 | son of |  |  |  |
| 3:9.310 | wicked |  |  |  |
| 3:9.311 | destined |  |  |  |
| 3:9.312 | you |  |  |  |
| 3:9.313 | to nullify |  |  |  |
| 3:9.314 | from Israel |  |  |  |
| 3:9.315 | many |  |  |  |
| 3:9.316 | and many |  |  |  |
| 3:9.317 | offerings |  |  |  |
| 3:9.318 | wilfully |  |  |  |
| 3:9.319 | and you |  |  |  |
| 3:9.320 | hold him liable |  |  | מחייבו: the pointing does not fix the subject and object; both are kept as in the Hebrew |
| 3:9.321 | for |  |  |  |
| 3:9.322 | the inadvertent |  |  |  |
| 3:9.323 | as the wilful |  |  |  |
| 3:9.324 | and for |  |  |  |
| 3:9.325 | the forced |  |  |  |
| 3:9.326 | as the voluntary |  |  |  |
| 3:9.327 | and this is |  |  |  |
| 3:9.328 | what is written |  |  |  |
| 3:9.329 | as the speaking of |  |  |  |
| 3:9.330 | Ephraim |  |  |  |
| 3:9.331 | trembling |  |  |  |
| 3:9.332 | he was exalted |  |  |  |
| 3:9.333 | he |  |  |  |
| 3:9.334 | in Israel |  |  |  |
| 3:9.335 | and became guilty |  |  |  |
| 3:9.336 | through Baal |  |  |  |
| 3:9.337 | and died |  |  |  |
| 3:9.338 | in the matter of |  |  |  |
| 3:9.339 | Jeroboam |  |  |  |
| 3:9.340 | who is |  |  |  |
| 3:9.341 | of the tribe of |  |  |  |
| 3:9.342 | Ephraim |  |  |  |
| 3:9.343 | rattat (trembled) |  |  | רתת: "trembled", the verse's word (Hos 13:1), read here by the author as the name of the accuser |
| 3:9.344 | katigorin (accuser) |  |  | קטיגורין: an accuser (the Talmudic term) |
| 3:9.345 | against |  |  |  |
| 3:9.346 | Solomon |  |  |  |
| 3:9.347 | he |  |  |  |
| 3:9.348 | Jeroboam |  |  |  |
| 3:9.349 | prince |  |  |  |
| 3:9.350 | in Israel |  |  |  |
| 3:9.351 | and king |  |  |  |
| 3:9.352 | and a bat |  |  |  |
| 3:9.353 | kol |  |  |  |
| 3:9.354 | answered him |  |  |  |
| 3:9.355 | that he would become guilty |  |  |  |
| 3:9.356 | through Baal |  |  |  |
| 3:9.357 | and die |  |  |  |
| 3:9.358 | thus far his words |  |  |  |
| 3:9.359 | and you |  |  |  |
| 3:9.360 | have magnified |  |  |  |
| 3:9.361 | and added |  |  |  |
| 3:9.362 | to do |  |  |  |
| 3:9.363 | evil |  |  |  |
| 3:9.364 | more |  |  |  |
| 3:9.365 | than Jeroboam |  |  |  |
| 3:9.366 | son of |  |  |  |
| 3:9.367 | Nebat |  |  |  |
| 3:9.368 | for |  |  |  |
| 3:9.369 | after |  |  |  |
| 3:9.370 | that |  |  |  |
| 3:9.371 | you brought |  |  |  |
| 3:9.372 | us |  |  |  |
| 3:9.373 | to |  |  |  |
| 3:9.374 | the house of |  |  |  |
| 3:9.375 | the prison |  |  |  |
| 3:9.376 | by |  |  |  |
| 3:9.377 | false informing |  |  |  |
| 3:9.378 | before |  |  |  |
| 3:9.379 | our government |  |  |  |
| 3:9.380 | exalted |  |  |  |
| 3:9.381 | you called out |  |  |  |
| 3:9.382 | - |  |  |  |
| 3:9.383 | all |  |  |  |
| 3:9.384 | the people |  |  |  |
| 3:9.385 | to go out |  |  |  |
| 3:9.386 | to |  |  |  |
| 3:9.387 | the house of |  |  |  |
| 3:9.388 | the cemetery |  |  |  |
| 3:9.389 | and you commanded |  |  |  |
| 3:9.390 | them |  |  |  |
| 3:9.391 | to dust themselves |  |  |  |
| 3:9.392 | with dust |  |  |  |
| 3:9.393 | on |  |  |  |
| 3:9.394 | their heads |  |  |  |
| 3:9.395 | for |  |  |  |
| 3:9.396 | no |  |  |  |
| 3:9.397 | thing |  |  |  |
| 3:9.398 | not |  |  |  |
| 3:9.399 | for |  |  |  |
| 3:9.400 | a commandment |  |  |  |
| 3:9.401 | light |  |  |  |
| 3:9.402 | that we neglected |  |  |  |
| 3:9.403 | and not |  |  |  |
| 3:9.404 | for |  |  |  |
| 3:9.405 | a commandment |  |  |  |
| 3:9.406 | weighty |  |  |  |
| 3:9.407 | God forbid |  |  | ח"ו: "God forbid" |
| 3:9.408 | but |  |  |  |
| 3:9.409 | only |  |  |  |
| 3:9.410 | for |  |  |  |
| 3:9.411 | that |  |  |  |
| 3:9.412 | we walked |  |  |  |
| 3:9.413 | in the custom of |  |  |  |
| 3:9.414 | our forefathers |  |  |  |
| 3:9.415 | and we set |  |  |  |
| 3:9.416 | our learning |  |  |  |
| 3:9.417 | in the Mishnah |  |  |  |
| 3:9.418 | and in the Talmud |  |  |  |
| 3:9.419 | and the Rambam |  |  | ורמב"ם: the Rambam |
| 3:9.420 | and Shulchan Aruch |  |  | וש"ע: the Shulchan Aruch |
| 3:9.421 | to learn |  |  |  |
| 3:9.422 | and to teach |  |  |  |
| 3:9.423 | to keep |  |  |  |
| 3:9.424 | and to do |  |  |  |
| 3:9.425 | and not |  |  |  |
| 3:9.426 | we studied |  |  |  |
| 3:9.427 | further |  |  |  |
| 3:9.428 | from the book |  |  |  |
| 3:9.429 | forged |  |  | מזויף: "forged", the author's epithet for the Zohar |
| 3:9.430 | the book |  |  |  |
| 3:9.431 | the Zohar |  |  | הזהר: "the book the Zohar", the title kept as the author writes it |
| 3:9.432 | as is the custom |  |  |  |
| 3:9.433 | of you |  |  |  |

Phrases:
- 3:9.1 | 9) | 9)
- 3:9.2 3:9.3 3:9.4 3:9.5 3:9.6 3:9.7 3:9.8 3:9.9 3:9.10 | And you, our friend, it was not enough for you that you did not do any one of all these | and you our friend not enough for you that not you did one of all these
- 3:9.11 3:9.12 3:9.13 3:9.14 3:9.15 | but you believed their words and accepted their lies | but that you believed to their words and you accepted their lies
- 3:9.16 3:9.17 3:9.18 3:9.19 | on the testimony of one witness and one judge | on the testimony of one and a judge one
- 3:9.20 3:9.21 3:9.22 3:9.23 | and not in the presence of the litigant | and not in the presence of the litigant of the case
- 3:9.24 3:9.25 3:9.26 3:9.27 3:9.28 3:9.29 | without inquiry, without investigation, without examination | and without inquiry and without investigation and without examination
- 3:9.30 3:9.31 | and you enjoyed their words | and you enjoyed in their words
- 3:9.32 3:9.33 3:9.34 3:9.35 | and you also helped to strengthen their lies | and also you helped to strengthen their lies
- 3:9.36 3:9.37 3:9.38 3:9.39 | and it was fixed in your mind as the truth itself | and was fixed in your mind as truth to its truth
- 3:9.40 3:9.41 3:9.42 3:9.43 3:9.44 3:9.45 3:9.46 | until you wrote that I am no longer able to deny it, etc. | until that you wrote that I am able to deny any longer etc.
- 3:9.47 3:9.48 3:9.49 3:9.50 3:9.51 3:9.52 3:9.53 3:9.54 3:9.55 3:9.56 3:9.57 | And you did not take to heart that testimony is accepted only in the presence of the litigant | and not you put to your heart that there is not they accept testimony but only in the presence of a litigant of the case
- 3:9.58 3:9.59 3:9.60 3:9.61 3:9.62 3:9.63 3:9.64 3:9.65 3:9.66 3:9.67 3:9.68 | and that you are obliged to inquire into them with seven inquiries, investigations and examinations, as you are commanded | and that you are obliged to inquire and to investigate them with seven inquiries and with investigations and with examinations as commanded of you
- 3:9.69 3:9.70 3:9.71 3:9.72 3:9.73 | All the more so, since they did their deeds in the dark | all the more so that they did in the dark their deeds
- 3:9.74 3:9.75 3:9.76 3:9.77 3:9.78 | and smoothed their tongues to you in secret with their words | and smoothed their tongue to you in secret with their words
- 3:9.79 3:9.80 3:9.81 | to find a fault to hate | to find a fault to hate
- 3:9.82 3:9.83 3:9.84 3:9.85 | and you were obliged to silence them with a rebuke. | and you were obliged to silence them with a rebuke
- 3:9.86 3:9.87 3:9.88 3:9.89 3:9.90 3:9.91 3:9.92 3:9.93 3:9.94 3:9.95 3:9.96 | This is nothing but a wicked heart: that you enjoy and long for their gossip and their informing. | it is not this that except wickedness of heart that you enjoy and long for their gossip and their informing
- 3:9.97 3:9.98 3:9.99 3:9.100 | the Merciful One save us from this opinion! | the Merciful One save us from this opinion
- 3:9.101 3:9.102 3:9.103 3:9.104 3:9.105 3:9.106 3:9.107 3:9.108 3:9.109 3:9.110 3:9.111 | And it is possible that you too add of your own false informing, in order to honor yourself at the disgrace of your friends. | and it is possible that also you add of your own informing false in order to honor yourself at the disgrace of your friends
- 3:9.112 3:9.113 3:9.114 3:9.115 3:9.116 | And in the chapter Arvei Pesachim (115) we read: | and in the chapter Arvei Pesachim (folio 115) we read
- 3:9.117 3:9.118 3:9.119 3:9.120 3:9.121 3:9.122 3:9.123 3:9.124 3:9.125 | "Rav Sheshet said in the name of R. A.: whoever tells evil talk | said Rav Sheshet in the name of R. A. whoever tells evil talk
- 3:9.126 3:9.127 3:9.128 3:9.129 3:9.130 3:9.131 3:9.132 3:9.133 | and whoever bears false witness against his fellow deserves to be thrown to the dogs, as it is said: | and whoever testifies testimony false against his fellow is fit to cast him to the dogs as it is said
- 3:9.134 3:9.135 3:9.136 | 'to the dog you shall throw it' | to the dog you shall throw it
- 3:9.137 3:9.138 3:9.139 3:9.140 3:9.141 3:9.142 3:9.143 3:9.144 3:9.145 3:9.146 | and it is written explicitly: 'you shall not carry a false report'; and it is read in it: 'you shall not cause to bear'" | and it is written explicitly not you shall carry report false and it is read in it not you shall cause to bear
- 3:9.147 3:9.148 3:9.149 3:9.150 3:9.151 3:9.152 3:9.153 3:9.154 3:9.155 | And we find several more homilies of our Sages of blessed memory on the punishment of one who tells evil talk. | and further several homilies we find our Sages of blessed memory on the punishment of the one who tells evil talk
- 3:9.156 3:9.157 3:9.158 3:9.159 3:9.160 3:9.161 3:9.162 | "And in Mahzor Vitry, tractate Avot, on what we learned: | and in Mahzor Vitry tractate Avot on that which we learned
- 3:9.163 3:9.164 3:9.165 3:9.166 3:9.167 3:9.168 3:9.169 | 'do not judge your fellow until you reach his place' | on (al: do not) judge - your fellow until you reach his place
- 3:9.170 3:9.171 3:9.172 3:9.173 3:9.174 3:9.175 | as he writes there: that the same transgression may come into your hand | as he writes there that may come into your hand that same transgression
- 3:9.176 3:9.177 3:9.178 | and you may be saved and escape from it | and be saved and escape from it
- 3:9.179 3:9.180 3:9.181 3:9.182 3:9.183 | as we find with Jeroboam son of Nebat | as we find with Jeroboam son of Nebat
- 3:9.184 3:9.185 3:9.186 3:9.187 3:9.188 3:9.189 | that when Solomon married the daughter of Pharaoh, | that when he married Solomon - daughter of Pharaoh
- 3:9.190 3:9.191 3:9.192 3:9.193 3:9.194 3:9.195 | she brought in to him every kind of song in the world | she brought in to him every kind of song in the world
- 3:9.196 3:9.197 3:9.198 3:9.199 3:9.200 3:9.201 | and said to him: 'So they do for such-and-such idolatry' | and she said to him thus they do for idolatry such-and-such
- 3:9.202 3:9.203 3:9.204 3:9.205 3:9.206 3:9.207 | One time, he went to sleep, and the day grew late | once time he went to sleep and became late the day
- 3:9.208 3:9.209 3:9.210 3:9.211 3:9.212 3:9.213 3:9.214 3:9.215 3:9.216 | she went and called the wood-carvers and the builders, who work at carving and weaving | she went and called the wood-carvers of wood and the builders who do the craft of carving and weaving
- 3:9.217 3:9.218 3:9.219 3:9.220 3:9.221 3:9.222 3:9.223 3:9.224 3:9.225 3:9.226 | and they made over his bed the likeness of a firmament, and in it the sun, the moon, the stars and the constellations | and they made on his bed the likeness of a firmament and in it the sun and the moon and the stars and the constellations
- 3:9.227 3:9.228 3:9.229 3:9.230 3:9.231 3:9.232 | In the morning he woke from his sleep to open the doors of the Temple, | and in the morning he woke from his sleep to open the doors of the Temple
- 3:9.233 3:9.234 3:9.235 3:9.236 3:9.237 | for the keys of the Temple were under his pillow | for there were keys of the Temple under his pillow
- 3:9.238 3:9.239 3:9.240 3:9.241 3:9.242 | he saw that the firmament was bright with stars and the host | he saw that the firmament was bright with stars and the host
- 3:9.243 3:9.244 3:9.245 3:9.246 3:9.247 | thinking that the night had not yet turned | thinking still not had turned night
- 3:9.248 3:9.249 | he went back to sleep | he returned to sleep
- 3:9.250 3:9.251 3:9.252 3:9.253 | And he slept four more hours, | and he slept more four hours
- 3:9.254 3:9.255 3:9.256 3:9.257 3:9.258 3:9.259 3:9.260 3:9.261 3:9.262 | until he understood and recognised for himself that the night was not so long | until he understood and recognised for himself that there is not the night long so thus
- 3:9.263 3:9.264 3:9.265 3:9.266 3:9.267 | And concerning that hour we learned a testimony | and concerning that hour we learned testimony
- 3:9.268 3:9.269 3:9.270 3:9.271 3:9.272 3:9.273 3:9.274 3:9.275 | about the morning daily offering, which is offered at four hours, in the Bechirta. | about the daily offering of morning which is offered at four hours in the Bechirta
- 3:9.276 3:9.277 3:9.278 3:9.279 3:9.280 3:9.281 3:9.282 | Before he woke from his sleep, Jeroboam son of Nebat went | and before he woke from his sleep went Jeroboam son of Nebat
- 3:9.283 3:9.284 3:9.285 | to the door of his house | to the door of his house
- 3:9.286 3:9.287 3:9.288 | and a voice went out: 'Wicked one | and went out a voice wicked one
- 3:9.289 3:9.290 3:9.291 3:9.292 3:9.293 3:9.294 3:9.295 | how long will you sleep and neglect the morning daily offering | until when you sleep and neglect the daily offering of morning
- 3:9.296 3:9.297 3:9.298 | What is it to you, sleeper? | what to you sleeper
- 3:9.299 3:9.300 3:9.301 3:9.302 | Arise, call upon your God!' | arise call upon your God
- 3:9.303 3:9.304 3:9.305 3:9.306 3:9.307 | A bat kol came out and said to him: | came out bat (a heavenly) voice and said to him
- 3:9.308 3:9.309 3:9.310 3:9.311 3:9.312 3:9.313 3:9.314 3:9.315 3:9.316 3:9.317 3:9.318 | 'Wicked son of a wicked man, you are destined to abolish from Israel many, many offerings wilfully. | wicked son of wicked destined you to nullify from Israel many and many offerings wilfully
- 3:9.319 3:9.320 3:9.321 3:9.322 3:9.323 | And you hold him liable for the inadvertent as for the wilful. | and you hold him liable for the inadvertent as the wilful
- 3:9.324 3:9.325 3:9.326 | And for the forced one as for the voluntary!' | and for the forced as the voluntary
- 3:9.327 3:9.328 3:9.329 3:9.330 | And this is what is written: 'When Ephraim spoke | and this is what is written as the speaking of Ephraim
- 3:9.331 3:9.332 3:9.333 3:9.334 3:9.335 3:9.336 3:9.337 | there was trembling; he was exalted in Israel, but he incurred guilt through Baal and died' | trembling he was exalted he in Israel and became guilty through Baal and died
- 3:9.338 3:9.339 3:9.340 3:9.341 3:9.342 | This is said of Jeroboam, who was of the tribe of Ephraim, | in the matter of Jeroboam who is of the tribe of Ephraim
- 3:9.343 3:9.344 3:9.345 3:9.346 | the trembling was the accuser against Solomon, | rattat (trembled) katigorin (accuser) against Solomon
- 3:9.347 3:9.348 3:9.349 3:9.350 3:9.351 | he, Jeroboam, was prince in Israel and king | he Jeroboam prince in Israel and king
- 3:9.352 3:9.353 3:9.354 3:9.355 3:9.356 3:9.357 3:9.358 | and the bat kol answered him that he would become guilty through Baal and die." Thus far his words. | and a bat kol answered him that he would become guilty through Baal and die thus far his words
- 3:9.359 3:9.360 3:9.361 3:9.362 3:9.363 3:9.364 3:9.365 3:9.366 3:9.367 | And you have magnified and added to doing evil more than Jeroboam son of Nebat. | and you have magnified and added to do evil more than Jeroboam son of Nebat
- 3:9.368 3:9.369 3:9.370 3:9.371 3:9.372 3:9.373 3:9.374 3:9.375 | For after you brought us to the prison | for after that you brought us to the house of the prison
- 3:9.376 3:9.377 3:9.378 3:9.379 3:9.380 | by false informing before our exalted government! | by false informing before our government exalted
- 3:9.381 3:9.382 3:9.383 3:9.384 3:9.385 3:9.386 3:9.387 3:9.388 | You called out all the people to go out to the cemetery, | you called out - all the people to go out to the house of the cemetery
- 3:9.389 3:9.390 3:9.391 3:9.392 3:9.393 3:9.394 | and commanded them to dust themselves with dust upon their heads | and you commanded them to dust themselves with dust on their heads
- 3:9.395 3:9.396 3:9.397 | for no reason. | for no thing
- 3:9.398 3:9.399 3:9.400 3:9.401 3:9.402 | Not for a light commandment that we neglected, | not for a commandment light that we neglected
- 3:9.403 3:9.404 3:9.405 3:9.406 3:9.407 | and not for a weighty commandment, God forbid, | and not for a commandment weighty God forbid
- 3:9.408 3:9.409 3:9.410 3:9.411 3:9.412 | but because we walked | but only for that we walked
- 3:9.413 3:9.414 3:9.415 3:9.416 | in the custom of our forefathers and set our learning | in the custom of our forefathers and we set our learning
- 3:9.417 3:9.418 3:9.419 3:9.420 | in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch | in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch
- 3:9.421 3:9.422 3:9.423 3:9.424 | to learn and to teach, to keep and to do. | to learn and to teach to keep and to do
- 3:9.425 3:9.426 3:9.427 3:9.428 3:9.429 3:9.430 3:9.431 3:9.432 3:9.433 | And we did not study further from the forged book, the book the Zohar, as is your custom. | and not we studied further from the book forged the book the Zohar as is the custom of you

Grammar notes:
- לא דייך שלא עשית (3:9.5-7): an idiom, "it was not enough for you that you did not do". The irony turns on the reversal: the charge is not only the omission but the belief and acceptance that follow ("but", אלא).
- בעד אחד ודיין אחד, ושלא בפני בעל דין, ובלא חקירה ובלא דרישה, ובלא בדיקה (3:9.16-29): the grounds of acceptance are listed in order, each negative with its own ובלא; the English keeps the repetition.
- The Pesachim passage, after גרסינן, is Talmudic Hebrew with its own formulas (אמר רב ששת, כתיב בהדיא, וקרי ביה). It cites Exod 22:30 and Exod 23:1; the Talmud's reading (לא תשיא) is set against the written form (לא תשא).
- The Mahzor Vitry passage runs from על הא דתנן to עכ"ל (thus far his words). The bat kol's speech sits inside it in single marks, and the verse of Hos 13:1 is its proof. The pronoun in ואתה מחייבו is left open, as the Hebrew leaves it.
- The last part turns from the story to the addressee: ואתה הגדלת (and you have magnified) and כי אחרי אשר (for after) make the charge that his informing brought the author to prison.

Translator's note: One paragraph (Wikisource, 1931 print). The printed list marker 9) is kept; the Hebrew abbreviations are expanded in the token glosses.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated, so the reading is flagged here until it is (Q2 default).
- apparatus: Exod 22:30 (לכלב תשליכון אותו), quoted as written: "to the dog you shall throw it".
- apparatus: Exod 23:1 (לא תשא שמע שוא), quoted as written; the Talmud's reading לא תשיא is given as read (וקרי).
- apparatus: Hos 13:1 (כדבר אפרים רתת נשא הוא בישראל ויאשם בבעל וימות), quoted by the author; translated from the Hebrew.
- apparatus: The Mishnah quoted is named only as tractate Avot; the author writes על for אל. No chapter reference is added.
- apparatus: Pesachim: the author gives folio 115 only.
- apparatus: ראב"ע and כ' ע"ש are not expanded with certainty; the English gives the likeliest sense.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

