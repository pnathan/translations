# Translation pack: milhamot-hashem perek-4-18 (Milchamot ha-Shem 4:30–4:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-18`.

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
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
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

- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- תורה **שבעל** פה → "the Oral Torah".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **והוא שלא** → "provided that".
- **קורין** אל → "call upon".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 4:30

Hebrew (Wikisource): לו) כבר ידעת ידידי הקורא מאמר רז"ל בפ"ק דשבת ובחגיגה פ' אין דורשין, אמר רב יהודה אמר רב ברם זכור אותו האיש לטוב וחנניה בן חזקיה שמו שאלמלא הוא נגנז ספר יחזקאל שהיו דבריו סותרין דברי תורה, מה עשה העלו לו שלש מאות גרבי שמן וישב בעליה ודרשן ע"כ, ופי' רש"י ז"ל סותרין דברי תורה כגון נבלה וטרפה לא יאכלו הכוהנים הא ישראל יאכלו? וכגון וכן תעשה בשבעה בחדש, היכן נרמז קרבן זה בתורה ודרשן כדדרשינן להני קראי במנחות דמשום דאישתרי מליקה לכהנים בחטאת העוף הוצרך להזהירן על מליקת חולין שהוא נבלה עכ"ל. ודון מינה ואוקי באתרין ומה ספר יחזקאל שכבר הוחזק לנביא אמת והחיה מתים בבקעת דורא וספרו מכלל ספרי הקדש ומצילין אותו מן הדליקה בשבת. ועם כל זה בקשו חכמים לגנזו מפני שהיה נראה כסותר דברי תורה באיסור לאו? כל שכן קל וחומר ס' הזהר שנראה וניכר בו סתירה לתורה שבכתב ושבעל פה באיסור ע"ז שהיא בסקילה וכרת. והוא ככופר בכל התורה כולה. ולא עוד אלא שהוא מבזה את המשנה והתלמוד ומכנה אותם בשם קליפה וסלע אחרא וכו' כמו שיתבאר לפנינו סי' ס"ז וס"ח. על אחת כמה וכמה שאסור לקרות בו כלל וכלל. כיון שכבר נתברר ונתפרסם זיופו לחכמים ונבונים ושדבריו בענין האלוהות סותרים דברי תורה ודברי סופרים המסורים לנו מרבותינו מעתיקי השמועה והגאונים והפוסקים שקבלו מהם. ועוד גרסינן בפ' במה מדליקין אמר רב יהודה בריה דרב יהודה בר שילת משמיה דרב ביקשו חכמים לגנוז ספר קהלת מפני שדבריו סותרים זה את זה וכו' באבות דרבי נתן בתחלה היו קורין משלי בלבד. ושיר השירים וקהלת היו גנוזים עד שבאו אנשי כנסת הגדולה ופירשום וקבעום בין הכתובים ע"כ הנה גם מזה תדין קל וחומר. ומה אם שלמה המלך ע"ה שנ' בו ויחכם מכל האדם. ובנה בית המקדש ושרתה בו שכינה ונראה אליו ה' כמבואר בכתוב. ואמרינן נמי בעירובין בשעה שתיקן עירובין ונטילת ידים. יצאה בת קול ואמרה בני אם חכם לבך ישמח לבי גם אני. ואומר חכם בני ושמח לבי ואשיבה חורפי דבר, ואמרינן נמי התם ואזן וחקר תיקן משלים הרבה, אמר ר"א בתחלה היתה תורה דומה לכפיפה שאין לה אזנים עד שבא שלמה ועשה לה אזנים! ועם כל השבח הזה לא נשאו לו חכמים פנים ובקשו לגנוז ספריו על שהיו דברים נראים סותרים זה את זה, עד שפירשום אנשי כנסת הגדולה! ספרים המבוארים שדבריהם סותרים זה את זה וגם סותרים דברי תורה ועוקרים יסודותיה. וגם לא נודע בבירור מי חברן. ומבטן מי יצאו הדברים הנאמרים בהם. ומיד בן נכר באו לידינו. רק שאותם אנשים שבאו לידם יחסו אותו לר' שמעון בן יוחאי התנא. ויש ידים מוכיחות דבשקר נתיחסו לרשב"י עאכ"ו שאסור לקרות בהן כלל.

| id | word | punctuation after |
|---|---|---|
| 4:30.1 | לו | ) |
| 4:30.2 | כבר |  |
| 4:30.3 | ידעת |  |
| 4:30.4 | ידידי |  |
| 4:30.5 | הקורא |  |
| 4:30.6 | מאמר |  |
| 4:30.7 | רז"ל |  |
| 4:30.8 | בפ"ק |  |
| 4:30.9 | דשבת |  |
| 4:30.10 | ובחגיגה |  |
| 4:30.11 | פ | ' |
| 4:30.12 | אין |  |
| 4:30.13 | דורשין | , |
| 4:30.14 | אמר |  |
| 4:30.15 | רב |  |
| 4:30.16 | יהודה |  |
| 4:30.17 | אמר |  |
| 4:30.18 | רב |  |
| 4:30.19 | ברם |  |
| 4:30.20 | זכור |  |
| 4:30.21 | אותו |  |
| 4:30.22 | האיש |  |
| 4:30.23 | לטוב |  |
| 4:30.24 | וחנניה |  |
| 4:30.25 | בן |  |
| 4:30.26 | חזקיה |  |
| 4:30.27 | שמו |  |
| 4:30.28 | שאלמלא |  |
| 4:30.29 | הוא |  |
| 4:30.30 | נגנז |  |
| 4:30.31 | ספר |  |
| 4:30.32 | יחזקאל |  |
| 4:30.33 | שהיו |  |
| 4:30.34 | דבריו |  |
| 4:30.35 | סותרין |  |
| 4:30.36 | דברי |  |
| 4:30.37 | תורה | , |
| 4:30.38 | מה |  |
| 4:30.39 | עשה |  |
| 4:30.40 | העלו |  |
| 4:30.41 | לו |  |
| 4:30.42 | שלש |  |
| 4:30.43 | מאות |  |
| 4:30.44 | גרבי |  |
| 4:30.45 | שמן |  |
| 4:30.46 | וישב |  |
| 4:30.47 | בעליה |  |
| 4:30.48 | ודרשן |  |
| 4:30.49 | ע"כ | , |
| 4:30.50 | ופי | ' |
| 4:30.51 | רש"י |  |
| 4:30.52 | ז"ל |  |
| 4:30.53 | סותרין |  |
| 4:30.54 | דברי |  |
| 4:30.55 | תורה |  |
| 4:30.56 | כגון |  |
| 4:30.57 | נבלה |  |
| 4:30.58 | וטרפה |  |
| 4:30.59 | לא |  |
| 4:30.60 | יאכלו |  |
| 4:30.61 | הכוהנים |  |
| 4:30.62 | הא |  |
| 4:30.63 | ישראל |  |
| 4:30.64 | יאכלו | ? |
| 4:30.65 | וכגון |  |
| 4:30.66 | וכן |  |
| 4:30.67 | תעשה |  |
| 4:30.68 | בשבעה |  |
| 4:30.69 | בחדש | , |
| 4:30.70 | היכן |  |
| 4:30.71 | נרמז |  |
| 4:30.72 | קרבן |  |
| 4:30.73 | זה |  |
| 4:30.74 | בתורה |  |
| 4:30.75 | ודרשן |  |
| 4:30.76 | כדדרשינן |  |
| 4:30.77 | להני |  |
| 4:30.78 | קראי |  |
| 4:30.79 | במנחות |  |
| 4:30.80 | דמשום |  |
| 4:30.81 | דאישתרי |  |
| 4:30.82 | מליקה |  |
| 4:30.83 | לכהנים |  |
| 4:30.84 | בחטאת |  |
| 4:30.85 | העוף |  |
| 4:30.86 | הוצרך |  |
| 4:30.87 | להזהירן |  |
| 4:30.88 | על |  |
| 4:30.89 | מליקת |  |
| 4:30.90 | חולין |  |
| 4:30.91 | שהוא |  |
| 4:30.92 | נבלה |  |
| 4:30.93 | עכ"ל | . |
| 4:30.94 | ודון |  |
| 4:30.95 | מינה |  |
| 4:30.96 | ואוקי |  |
| 4:30.97 | באתרין |  |
| 4:30.98 | ומה |  |
| 4:30.99 | ספר |  |
| 4:30.100 | יחזקאל |  |
| 4:30.101 | שכבר |  |
| 4:30.102 | הוחזק |  |
| 4:30.103 | לנביא |  |
| 4:30.104 | אמת |  |
| 4:30.105 | והחיה |  |
| 4:30.106 | מתים |  |
| 4:30.107 | בבקעת |  |
| 4:30.108 | דורא |  |
| 4:30.109 | וספרו |  |
| 4:30.110 | מכלל |  |
| 4:30.111 | ספרי |  |
| 4:30.112 | הקדש |  |
| 4:30.113 | ומצילין |  |
| 4:30.114 | אותו |  |
| 4:30.115 | מן |  |
| 4:30.116 | הדליקה |  |
| 4:30.117 | בשבת | . |
| 4:30.118 | ועם |  |
| 4:30.119 | כל |  |
| 4:30.120 | זה |  |
| 4:30.121 | בקשו |  |
| 4:30.122 | חכמים |  |
| 4:30.123 | לגנזו |  |
| 4:30.124 | מפני |  |
| 4:30.125 | שהיה |  |
| 4:30.126 | נראה |  |
| 4:30.127 | כסותר |  |
| 4:30.128 | דברי |  |
| 4:30.129 | תורה |  |
| 4:30.130 | באיסור |  |
| 4:30.131 | לאו | ? |
| 4:30.132 | כל |  |
| 4:30.133 | שכן |  |
| 4:30.134 | קל |  |
| 4:30.135 | וחומר |  |
| 4:30.136 | ס | ' |
| 4:30.137 | הזהר |  |
| 4:30.138 | שנראה |  |
| 4:30.139 | וניכר |  |
| 4:30.140 | בו |  |
| 4:30.141 | סתירה |  |
| 4:30.142 | לתורה |  |
| 4:30.143 | שבכתב |  |
| 4:30.144 | ושבעל |  |
| 4:30.145 | פה |  |
| 4:30.146 | באיסור |  |
| 4:30.147 | ע"ז |  |
| 4:30.148 | שהיא |  |
| 4:30.149 | בסקילה |  |
| 4:30.150 | וכרת | . |
| 4:30.151 | והוא |  |
| 4:30.152 | ככופר |  |
| 4:30.153 | בכל |  |
| 4:30.154 | התורה |  |
| 4:30.155 | כולה | . |
| 4:30.156 | ולא |  |
| 4:30.157 | עוד |  |
| 4:30.158 | אלא |  |
| 4:30.159 | שהוא |  |
| 4:30.160 | מבזה |  |
| 4:30.161 | את |  |
| 4:30.162 | המשנה |  |
| 4:30.163 | והתלמוד |  |
| 4:30.164 | ומכנה |  |
| 4:30.165 | אותם |  |
| 4:30.166 | בשם |  |
| 4:30.167 | קליפה |  |
| 4:30.168 | וסלע |  |
| 4:30.169 | אחרא |  |
| 4:30.170 | וכו | ' |
| 4:30.171 | כמו |  |
| 4:30.172 | שיתבאר |  |
| 4:30.173 | לפנינו |  |
| 4:30.174 | סי | ' |
| 4:30.175 | ס"ז |  |
| 4:30.176 | וס"ח | . |
| 4:30.177 | על |  |
| 4:30.178 | אחת |  |
| 4:30.179 | כמה |  |
| 4:30.180 | וכמה |  |
| 4:30.181 | שאסור |  |
| 4:30.182 | לקרות |  |
| 4:30.183 | בו |  |
| 4:30.184 | כלל |  |
| 4:30.185 | וכלל | . |
| 4:30.186 | כיון |  |
| 4:30.187 | שכבר |  |
| 4:30.188 | נתברר |  |
| 4:30.189 | ונתפרסם |  |
| 4:30.190 | זיופו |  |
| 4:30.191 | לחכמים |  |
| 4:30.192 | ונבונים |  |
| 4:30.193 | ושדבריו |  |
| 4:30.194 | בענין |  |
| 4:30.195 | האלוהות |  |
| 4:30.196 | סותרים |  |
| 4:30.197 | דברי |  |
| 4:30.198 | תורה |  |
| 4:30.199 | ודברי |  |
| 4:30.200 | סופרים |  |
| 4:30.201 | המסורים |  |
| 4:30.202 | לנו |  |
| 4:30.203 | מרבותינו |  |
| 4:30.204 | מעתיקי |  |
| 4:30.205 | השמועה |  |
| 4:30.206 | והגאונים |  |
| 4:30.207 | והפוסקים |  |
| 4:30.208 | שקבלו |  |
| 4:30.209 | מהם | . |
| 4:30.210 | ועוד |  |
| 4:30.211 | גרסינן |  |
| 4:30.212 | בפ | ' |
| 4:30.213 | במה |  |
| 4:30.214 | מדליקין |  |
| 4:30.215 | אמר |  |
| 4:30.216 | רב |  |
| 4:30.217 | יהודה |  |
| 4:30.218 | בריה |  |
| 4:30.219 | דרב |  |
| 4:30.220 | יהודה |  |
| 4:30.221 | בר |  |
| 4:30.222 | שילת |  |
| 4:30.223 | משמיה |  |
| 4:30.224 | דרב |  |
| 4:30.225 | ביקשו |  |
| 4:30.226 | חכמים |  |
| 4:30.227 | לגנוז |  |
| 4:30.228 | ספר |  |
| 4:30.229 | קהלת |  |
| 4:30.230 | מפני |  |
| 4:30.231 | שדבריו |  |
| 4:30.232 | סותרים |  |
| 4:30.233 | זה |  |
| 4:30.234 | את |  |
| 4:30.235 | זה |  |
| 4:30.236 | וכו | ' |
| 4:30.237 | באבות |  |
| 4:30.238 | דרבי |  |
| 4:30.239 | נתן |  |
| 4:30.240 | בתחלה |  |
| 4:30.241 | היו |  |
| 4:30.242 | קורין |  |
| 4:30.243 | משלי |  |
| 4:30.244 | בלבד | . |
| 4:30.245 | ושיר |  |
| 4:30.246 | השירים |  |
| 4:30.247 | וקהלת |  |
| 4:30.248 | היו |  |
| 4:30.249 | גנוזים |  |
| 4:30.250 | עד |  |
| 4:30.251 | שבאו |  |
| 4:30.252 | אנשי |  |
| 4:30.253 | כנסת |  |
| 4:30.254 | הגדולה |  |
| 4:30.255 | ופירשום |  |
| 4:30.256 | וקבעום |  |
| 4:30.257 | בין |  |
| 4:30.258 | הכתובים |  |
| 4:30.259 | ע"כ |  |
| 4:30.260 | הנה |  |
| 4:30.261 | גם |  |
| 4:30.262 | מזה |  |
| 4:30.263 | תדין |  |
| 4:30.264 | קל |  |
| 4:30.265 | וחומר | . |
| 4:30.266 | ומה |  |
| 4:30.267 | אם |  |
| 4:30.268 | שלמה |  |
| 4:30.269 | המלך |  |
| 4:30.270 | ע"ה |  |
| 4:30.271 | שנ | ' |
| 4:30.272 | בו |  |
| 4:30.273 | ויחכם |  |
| 4:30.274 | מכל |  |
| 4:30.275 | האדם | . |
| 4:30.276 | ובנה |  |
| 4:30.277 | בית |  |
| 4:30.278 | המקדש |  |
| 4:30.279 | ושרתה |  |
| 4:30.280 | בו |  |
| 4:30.281 | שכינה |  |
| 4:30.282 | ונראה |  |
| 4:30.283 | אליו |  |
| 4:30.284 | ה | ' |
| 4:30.285 | כמבואר |  |
| 4:30.286 | בכתוב | . |
| 4:30.287 | ואמרינן |  |
| 4:30.288 | נמי |  |
| 4:30.289 | בעירובין |  |
| 4:30.290 | בשעה |  |
| 4:30.291 | שתיקן |  |
| 4:30.292 | עירובין |  |
| 4:30.293 | ונטילת |  |
| 4:30.294 | ידים | . |
| 4:30.295 | יצאה |  |
| 4:30.296 | בת |  |
| 4:30.297 | קול |  |
| 4:30.298 | ואמרה |  |
| 4:30.299 | בני |  |
| 4:30.300 | אם |  |
| 4:30.301 | חכם |  |
| 4:30.302 | לבך |  |
| 4:30.303 | ישמח |  |
| 4:30.304 | לבי |  |
| 4:30.305 | גם |  |
| 4:30.306 | אני | . |
| 4:30.307 | ואומר |  |
| 4:30.308 | חכם |  |
| 4:30.309 | בני |  |
| 4:30.310 | ושמח |  |
| 4:30.311 | לבי |  |
| 4:30.312 | ואשיבה |  |
| 4:30.313 | חורפי |  |
| 4:30.314 | דבר | , |
| 4:30.315 | ואמרינן |  |
| 4:30.316 | נמי |  |
| 4:30.317 | התם |  |
| 4:30.318 | ואזן |  |
| 4:30.319 | וחקר |  |
| 4:30.320 | תיקן |  |
| 4:30.321 | משלים |  |
| 4:30.322 | הרבה | , |
| 4:30.323 | אמר |  |
| 4:30.324 | ר"א |  |
| 4:30.325 | בתחלה |  |
| 4:30.326 | היתה |  |
| 4:30.327 | תורה |  |
| 4:30.328 | דומה |  |
| 4:30.329 | לכפיפה |  |
| 4:30.330 | שאין |  |
| 4:30.331 | לה |  |
| 4:30.332 | אזנים |  |
| 4:30.333 | עד |  |
| 4:30.334 | שבא |  |
| 4:30.335 | שלמה |  |
| 4:30.336 | ועשה |  |
| 4:30.337 | לה |  |
| 4:30.338 | אזנים | ! |
| 4:30.339 | ועם |  |
| 4:30.340 | כל |  |
| 4:30.341 | השבח |  |
| 4:30.342 | הזה |  |
| 4:30.343 | לא |  |
| 4:30.344 | נשאו |  |
| 4:30.345 | לו |  |
| 4:30.346 | חכמים |  |
| 4:30.347 | פנים |  |
| 4:30.348 | ובקשו |  |
| 4:30.349 | לגנוז |  |
| 4:30.350 | ספריו |  |
| 4:30.351 | על |  |
| 4:30.352 | שהיו |  |
| 4:30.353 | דברים |  |
| 4:30.354 | נראים |  |
| 4:30.355 | סותרים |  |
| 4:30.356 | זה |  |
| 4:30.357 | את |  |
| 4:30.358 | זה | , |
| 4:30.359 | עד |  |
| 4:30.360 | שפירשום |  |
| 4:30.361 | אנשי |  |
| 4:30.362 | כנסת |  |
| 4:30.363 | הגדולה | ! |
| 4:30.364 | ספרים |  |
| 4:30.365 | המבוארים |  |
| 4:30.366 | שדבריהם |  |
| 4:30.367 | סותרים |  |
| 4:30.368 | זה |  |
| 4:30.369 | את |  |
| 4:30.370 | זה |  |
| 4:30.371 | וגם |  |
| 4:30.372 | סותרים |  |
| 4:30.373 | דברי |  |
| 4:30.374 | תורה |  |
| 4:30.375 | ועוקרים |  |
| 4:30.376 | יסודותיה | . |
| 4:30.377 | וגם |  |
| 4:30.378 | לא |  |
| 4:30.379 | נודע |  |
| 4:30.380 | בבירור |  |
| 4:30.381 | מי |  |
| 4:30.382 | חברן | . |
| 4:30.383 | ומבטן |  |
| 4:30.384 | מי |  |
| 4:30.385 | יצאו |  |
| 4:30.386 | הדברים |  |
| 4:30.387 | הנאמרים |  |
| 4:30.388 | בהם | . |
| 4:30.389 | ומיד |  |
| 4:30.390 | בן |  |
| 4:30.391 | נכר |  |
| 4:30.392 | באו |  |
| 4:30.393 | לידינו | . |
| 4:30.394 | רק |  |
| 4:30.395 | שאותם |  |
| 4:30.396 | אנשים |  |
| 4:30.397 | שבאו |  |
| 4:30.398 | לידם |  |
| 4:30.399 | יחסו |  |
| 4:30.400 | אותו |  |
| 4:30.401 | לר | ' |
| 4:30.402 | שמעון |  |
| 4:30.403 | בן |  |
| 4:30.404 | יוחאי |  |
| 4:30.405 | התנא | . |
| 4:30.406 | ויש |  |
| 4:30.407 | ידים |  |
| 4:30.408 | מוכיחות |  |
| 4:30.409 | דבשקר |  |
| 4:30.410 | נתיחסו |  |
| 4:30.411 | לרשב"י |  |
| 4:30.412 | עאכ"ו |  |
| 4:30.413 | שאסור |  |
| 4:30.414 | לקרות |  |
| 4:30.415 | בהן |  |
| 4:30.416 | כלל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

