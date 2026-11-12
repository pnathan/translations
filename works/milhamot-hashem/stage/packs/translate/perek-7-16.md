# Translation pack: milhamot-hashem perek-7-16 (Milchamot ha-Shem 7:21–7:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-16`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ראב"ע** → "Rabbi Eleazar ben Azariah" in a Talmudic passage, "Ibn Ezra" for
  the commentator (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **זווג** → "coupling".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
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
- **כ"ש / ק"ו** → "all the more so".
- **תוה"ק** → "the holy Torah".
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **אחי** → "my brother" (second person kept).
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
- **תא חזי** → "come and see".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 7:21

Hebrew (Wikisource): צג) מכל זה למדנו גנות המספר בעניינו המשגל ונבלותו ועונשו החמור אף אם מספר על אדם פחות ומזולזל. או על עצמו איך הוא נזקק לאשתו. ואימתי הוא נזקק לה. וכיצד נזקק לה ואיך הוא מחבקה ומנשקה, כ"ש וק"ו אם מספר ומגלה קלון מלך בשר ודם אימתי משמש ונזקק לאשתו ואיך הוא מחבק לה ומנשקה, בודאי שהמספר ומגלה דברים כאלה, גנאי גדול הוא למלך, וכתוב וערות מלכא לא אריך נא למחזי, ופי' רש"י והראב"ע ז"ל וערות מלכא בזיון המלך וגנותו, לא אריך לנא למחזי, אין הגון לנו לראות ע"כ. ותא חזי מאי סלקא ביה בחם אבי כנען על ידי שראה ערות אביו וקלונו והגיד לשני אחיו שם ויפת, והדברים קל וחומר, ומה זה שספר לאחיו גנות אביו כך, המספר במלכו של עולם עאכ"ו. ובילקוט רמז ס"א, ויתגל בתוך אהלה, אהלה כתיב, בתוך אהלה של אשת נח, כשיצא נח מן התבה הכישו הארי ושברו, ובא לו לשמש מטתו ונתפזר זרעו, וירא חם אבי כנען ויגד, אמר להון ואגיד להון. אמר לון, אדם הראשון היו לו שני בנים ועמד אחד והרג את חבירו, וזה יש לו שלשה בנים, והוא מבקש לעשות ארבעה? מה טעם העבד יוצא בשן ועין, מן הכא וירא ויגד, ויקח שם ויפת את השמלה, שם התחיל במצוה, ובא יפת ונשמע לו לפי' זכה שם לטלית ויפת לפלוניא, ממשמע שנ' וילכו אחורנית איני יודע שערות אביהם לא ראו, אלא מלמד נתנו ידיהם על פניהם והיו מהלכין לאחריהם ונהגו בו כבוד כמורא האב על הבן, אמר הקב"ה לשם, אתה כסית את ערות אביך. חייך שאני פורעה לבניך, שנ' באדין גובריא אליך כפיתו בסרבליהון, אמר הקב"ה ליפת אתה כסית ערות אביך חייך שאני פורע לך שנ' ביום ההוא אתן לגוג מקום קבר בישראל. אמר הקב"ה להם אתה בזית עריתא דאבוך, חייך שאני פורע לך שנ' כן ינהג מלך אשור את שבי מצרים ואת גלות כוש נערים וזקנים ערום ויחף וחשופי שת ערות מצרים עכ"ל. ראה תראה ותדין ק"ו כמו שכתבתי לעיל, ובודאי שהרואה בזיון המלך זגגותו ואינו מוחה, וכ"ש המספר בגנות המלך מלכו של עולם. ומיחס לו עניני תשמיש וזיווג ולידה, עאכ"ו שהוא נעקר מן העה"ז ומן העה"ב. ואין לו חלק באלהי ישראל. כי איככה ניחס לאלוהינו אשר אותו נעבוד פילגש ואיברי משגל. יסוד דהינו גיד הערוה וביצי זכר לבשל הזרע. ונקב אחוריים להוציא הפסולת להיות ניזון ממנה אל אחר? הלא נבוש ונכלם לספר עניינים כאלה בבשר ודם כי אם ברמז לעת הצורך כגון שיש מחלה באיברי המשגל או בפי הטבעת, ועם כל זה אין מזכירים אותם אלא ברמז וכינוי ובהצנע ולא בפרסום אבל שלא לצורך הם מלהזכיר כלל. כי זהו גבלות הפה וגנות גדולה שלמדה אותנו תוה"ק להשמר ממנו כמאמר רז"ל בפ"ק דפסחים אמר ר' יהושע בן לוי לעולם אל יוציא אדם דבר מגונה מפיו שהרי עיקם הכ' שמנה אותיות ולא הוציא דבר מגונה מפיו שנ' ומן הבהמה אשר איננה טהורה רב אשי אמר תשע שנ' כי יהיה בך איש אשר לא יהיה טהור מקרה לילה. רבינא אמר עשר וי"ו דטהור. רב אמר שש עשרה שנ' כי אמר מקרה הוא בלתי טהור הוא כי לא טהור. תנא דבי ר' ישמעאל לעולם יספר אדם בלשון נקייה שהרי בזב וכו'. ואומר ותבחר לשון ערומים. ואומר וזעת שפתי ברור מללו וכו'. ובכל התורה ובדברי הנביאים ורז"ל בתלמוד ובמדרשים לא מצינו שכינו להשפעת השי"ת בשם זכר ונקבה וזיווג אלא בלשון נתינה או פתיחה ויתן לך האלהים. נותן לך כח לעשות חיל. יפתח ה' לך את אוצרו הטוב נותן לחם לכל בשר. פותח ידיך ומשביע לכל חי רצון. ואתה מחיה את כלם וכיוצא בלשונות אלו אן מספר. ובדברי רז"ל יושב הקב"ה וזן את העולם כו'. זן ומפרניס לכל. עורך שולחן, נותן טרף וכיוצא. ולמה זה נעזוב מליצת ספרי הקדש שדברו בלשון נקייה. ונשתמש בשלון מגונה סכרים ונקבות רבות וזווגים והטפת טיפי הזרע ממבושי הזכר לרחמה של נקבה וקליטה וקפאון הזרע ברחמה כמפורסם כל זה בספרי הקבלה החדשה הנכריה. ולא די כל זה אלא שהמסית מחבר הזהר העיז פניו לכנות המבושים אשר חשב באדון העולמים בשם בית אלהים. ושער השמים, כדאיתא בזהר ויצא דף ק"ן ע"ב.

| id | word | punctuation after |
|---|---|---|
| 7:21.1 | צג | ) |
| 7:21.2 | מכל |  |
| 7:21.3 | זה |  |
| 7:21.4 | למדנו |  |
| 7:21.5 | גנות |  |
| 7:21.6 | המספר |  |
| 7:21.7 | בעניינו |  |
| 7:21.8 | המשגל |  |
| 7:21.9 | ונבלותו |  |
| 7:21.10 | ועונשו |  |
| 7:21.11 | החמור |  |
| 7:21.12 | אף |  |
| 7:21.13 | אם |  |
| 7:21.14 | מספר |  |
| 7:21.15 | על |  |
| 7:21.16 | אדם |  |
| 7:21.17 | פחות |  |
| 7:21.18 | ומזולזל | . |
| 7:21.19 | או |  |
| 7:21.20 | על |  |
| 7:21.21 | עצמו |  |
| 7:21.22 | איך |  |
| 7:21.23 | הוא |  |
| 7:21.24 | נזקק |  |
| 7:21.25 | לאשתו | . |
| 7:21.26 | ואימתי |  |
| 7:21.27 | הוא |  |
| 7:21.28 | נזקק |  |
| 7:21.29 | לה | . |
| 7:21.30 | וכיצד |  |
| 7:21.31 | נזקק |  |
| 7:21.32 | לה |  |
| 7:21.33 | ואיך |  |
| 7:21.34 | הוא |  |
| 7:21.35 | מחבקה |  |
| 7:21.36 | ומנשקה | , |
| 7:21.37 | כ"ש |  |
| 7:21.38 | וק"ו |  |
| 7:21.39 | אם |  |
| 7:21.40 | מספר |  |
| 7:21.41 | ומגלה |  |
| 7:21.42 | קלון |  |
| 7:21.43 | מלך |  |
| 7:21.44 | בשר |  |
| 7:21.45 | ודם |  |
| 7:21.46 | אימתי |  |
| 7:21.47 | משמש |  |
| 7:21.48 | ונזקק |  |
| 7:21.49 | לאשתו |  |
| 7:21.50 | ואיך |  |
| 7:21.51 | הוא |  |
| 7:21.52 | מחבק |  |
| 7:21.53 | לה |  |
| 7:21.54 | ומנשקה | , |
| 7:21.55 | בודאי |  |
| 7:21.56 | שהמספר |  |
| 7:21.57 | ומגלה |  |
| 7:21.58 | דברים |  |
| 7:21.59 | כאלה | , |
| 7:21.60 | גנאי |  |
| 7:21.61 | גדול |  |
| 7:21.62 | הוא |  |
| 7:21.63 | למלך | , |
| 7:21.64 | וכתוב |  |
| 7:21.65 | וערות |  |
| 7:21.66 | מלכא |  |
| 7:21.67 | לא |  |
| 7:21.68 | אריך |  |
| 7:21.69 | נא |  |
| 7:21.70 | למחזי | , |
| 7:21.71 | ופי | ' |
| 7:21.72 | רש"י |  |
| 7:21.73 | והראב"ע |  |
| 7:21.74 | ז"ל |  |
| 7:21.75 | וערות |  |
| 7:21.76 | מלכא |  |
| 7:21.77 | בזיון |  |
| 7:21.78 | המלך |  |
| 7:21.79 | וגנותו | , |
| 7:21.80 | לא |  |
| 7:21.81 | אריך |  |
| 7:21.82 | לנא |  |
| 7:21.83 | למחזי | , |
| 7:21.84 | אין |  |
| 7:21.85 | הגון |  |
| 7:21.86 | לנו |  |
| 7:21.87 | לראות |  |
| 7:21.88 | ע"כ | . |
| 7:21.89 | ותא |  |
| 7:21.90 | חזי |  |
| 7:21.91 | מאי |  |
| 7:21.92 | סלקא |  |
| 7:21.93 | ביה |  |
| 7:21.94 | בחם |  |
| 7:21.95 | אבי |  |
| 7:21.96 | כנען |  |
| 7:21.97 | על |  |
| 7:21.98 | ידי |  |
| 7:21.99 | שראה |  |
| 7:21.100 | ערות |  |
| 7:21.101 | אביו |  |
| 7:21.102 | וקלונו |  |
| 7:21.103 | והגיד |  |
| 7:21.104 | לשני |  |
| 7:21.105 | אחיו |  |
| 7:21.106 | שם |  |
| 7:21.107 | ויפת | , |
| 7:21.108 | והדברים |  |
| 7:21.109 | קל |  |
| 7:21.110 | וחומר | , |
| 7:21.111 | ומה |  |
| 7:21.112 | זה |  |
| 7:21.113 | שספר |  |
| 7:21.114 | לאחיו |  |
| 7:21.115 | גנות |  |
| 7:21.116 | אביו |  |
| 7:21.117 | כך | , |
| 7:21.118 | המספר |  |
| 7:21.119 | במלכו |  |
| 7:21.120 | של |  |
| 7:21.121 | עולם |  |
| 7:21.122 | עאכ"ו | . |
| 7:21.123 | ובילקוט |  |
| 7:21.124 | רמז |  |
| 7:21.125 | ס"א | , |
| 7:21.126 | ויתגל |  |
| 7:21.127 | בתוך |  |
| 7:21.128 | אהלה | , |
| 7:21.129 | אהלה |  |
| 7:21.130 | כתיב | , |
| 7:21.131 | בתוך |  |
| 7:21.132 | אהלה |  |
| 7:21.133 | של |  |
| 7:21.134 | אשת |  |
| 7:21.135 | נח | , |
| 7:21.136 | כשיצא |  |
| 7:21.137 | נח |  |
| 7:21.138 | מן |  |
| 7:21.139 | התבה |  |
| 7:21.140 | הכישו |  |
| 7:21.141 | הארי |  |
| 7:21.142 | ושברו | , |
| 7:21.143 | ובא |  |
| 7:21.144 | לו |  |
| 7:21.145 | לשמש |  |
| 7:21.146 | מטתו |  |
| 7:21.147 | ונתפזר |  |
| 7:21.148 | זרעו | , |
| 7:21.149 | וירא |  |
| 7:21.150 | חם |  |
| 7:21.151 | אבי |  |
| 7:21.152 | כנען |  |
| 7:21.153 | ויגד | , |
| 7:21.154 | אמר |  |
| 7:21.155 | להון |  |
| 7:21.156 | ואגיד |  |
| 7:21.157 | להון | . |
| 7:21.158 | אמר |  |
| 7:21.159 | לון | , |
| 7:21.160 | אדם |  |
| 7:21.161 | הראשון |  |
| 7:21.162 | היו |  |
| 7:21.163 | לו |  |
| 7:21.164 | שני |  |
| 7:21.165 | בנים |  |
| 7:21.166 | ועמד |  |
| 7:21.167 | אחד |  |
| 7:21.168 | והרג |  |
| 7:21.169 | את |  |
| 7:21.170 | חבירו | , |
| 7:21.171 | וזה |  |
| 7:21.172 | יש |  |
| 7:21.173 | לו |  |
| 7:21.174 | שלשה |  |
| 7:21.175 | בנים | , |
| 7:21.176 | והוא |  |
| 7:21.177 | מבקש |  |
| 7:21.178 | לעשות |  |
| 7:21.179 | ארבעה | ? |
| 7:21.180 | מה |  |
| 7:21.181 | טעם |  |
| 7:21.182 | העבד |  |
| 7:21.183 | יוצא |  |
| 7:21.184 | בשן |  |
| 7:21.185 | ועין | , |
| 7:21.186 | מן |  |
| 7:21.187 | הכא |  |
| 7:21.188 | וירא |  |
| 7:21.189 | ויגד | , |
| 7:21.190 | ויקח |  |
| 7:21.191 | שם |  |
| 7:21.192 | ויפת |  |
| 7:21.193 | את |  |
| 7:21.194 | השמלה | , |
| 7:21.195 | שם |  |
| 7:21.196 | התחיל |  |
| 7:21.197 | במצוה | , |
| 7:21.198 | ובא |  |
| 7:21.199 | יפת |  |
| 7:21.200 | ונשמע |  |
| 7:21.201 | לו |  |
| 7:21.202 | לפי | ' |
| 7:21.203 | זכה |  |
| 7:21.204 | שם |  |
| 7:21.205 | לטלית |  |
| 7:21.206 | ויפת |  |
| 7:21.207 | לפלוניא | , |
| 7:21.208 | ממשמע |  |
| 7:21.209 | שנ | ' |
| 7:21.210 | וילכו |  |
| 7:21.211 | אחורנית |  |
| 7:21.212 | איני |  |
| 7:21.213 | יודע |  |
| 7:21.214 | שערות |  |
| 7:21.215 | אביהם |  |
| 7:21.216 | לא |  |
| 7:21.217 | ראו | , |
| 7:21.218 | אלא |  |
| 7:21.219 | מלמד |  |
| 7:21.220 | נתנו |  |
| 7:21.221 | ידיהם |  |
| 7:21.222 | על |  |
| 7:21.223 | פניהם |  |
| 7:21.224 | והיו |  |
| 7:21.225 | מהלכין |  |
| 7:21.226 | לאחריהם |  |
| 7:21.227 | ונהגו |  |
| 7:21.228 | בו |  |
| 7:21.229 | כבוד |  |
| 7:21.230 | כמורא |  |
| 7:21.231 | האב |  |
| 7:21.232 | על |  |
| 7:21.233 | הבן | , |
| 7:21.234 | אמר |  |
| 7:21.235 | הקב"ה |  |
| 7:21.236 | לשם | , |
| 7:21.237 | אתה |  |
| 7:21.238 | כסית |  |
| 7:21.239 | את |  |
| 7:21.240 | ערות |  |
| 7:21.241 | אביך | . |
| 7:21.242 | חייך |  |
| 7:21.243 | שאני |  |
| 7:21.244 | פורעה |  |
| 7:21.245 | לבניך | , |
| 7:21.246 | שנ | ' |
| 7:21.247 | באדין |  |
| 7:21.248 | גובריא |  |
| 7:21.249 | אליך |  |
| 7:21.250 | כפיתו |  |
| 7:21.251 | בסרבליהון | , |
| 7:21.252 | אמר |  |
| 7:21.253 | הקב"ה |  |
| 7:21.254 | ליפת |  |
| 7:21.255 | אתה |  |
| 7:21.256 | כסית |  |
| 7:21.257 | ערות |  |
| 7:21.258 | אביך |  |
| 7:21.259 | חייך |  |
| 7:21.260 | שאני |  |
| 7:21.261 | פורע |  |
| 7:21.262 | לך |  |
| 7:21.263 | שנ | ' |
| 7:21.264 | ביום |  |
| 7:21.265 | ההוא |  |
| 7:21.266 | אתן |  |
| 7:21.267 | לגוג |  |
| 7:21.268 | מקום |  |
| 7:21.269 | קבר |  |
| 7:21.270 | בישראל | . |
| 7:21.271 | אמר |  |
| 7:21.272 | הקב"ה |  |
| 7:21.273 | להם |  |
| 7:21.274 | אתה |  |
| 7:21.275 | בזית |  |
| 7:21.276 | עריתא |  |
| 7:21.277 | דאבוך | , |
| 7:21.278 | חייך |  |
| 7:21.279 | שאני |  |
| 7:21.280 | פורע |  |
| 7:21.281 | לך |  |
| 7:21.282 | שנ | ' |
| 7:21.283 | כן |  |
| 7:21.284 | ינהג |  |
| 7:21.285 | מלך |  |
| 7:21.286 | אשור |  |
| 7:21.287 | את |  |
| 7:21.288 | שבי |  |
| 7:21.289 | מצרים |  |
| 7:21.290 | ואת |  |
| 7:21.291 | גלות |  |
| 7:21.292 | כוש |  |
| 7:21.293 | נערים |  |
| 7:21.294 | וזקנים |  |
| 7:21.295 | ערום |  |
| 7:21.296 | ויחף |  |
| 7:21.297 | וחשופי |  |
| 7:21.298 | שת |  |
| 7:21.299 | ערות |  |
| 7:21.300 | מצרים |  |
| 7:21.301 | עכ"ל | . |
| 7:21.302 | ראה |  |
| 7:21.303 | תראה |  |
| 7:21.304 | ותדין |  |
| 7:21.305 | ק"ו |  |
| 7:21.306 | כמו |  |
| 7:21.307 | שכתבתי |  |
| 7:21.308 | לעיל | , |
| 7:21.309 | ובודאי |  |
| 7:21.310 | שהרואה |  |
| 7:21.311 | בזיון |  |
| 7:21.312 | המלך |  |
| 7:21.313 | זגגותו |  |
| 7:21.314 | ואינו |  |
| 7:21.315 | מוחה | , |
| 7:21.316 | וכ"ש |  |
| 7:21.317 | המספר |  |
| 7:21.318 | בגנות |  |
| 7:21.319 | המלך |  |
| 7:21.320 | מלכו |  |
| 7:21.321 | של |  |
| 7:21.322 | עולם | . |
| 7:21.323 | ומיחס |  |
| 7:21.324 | לו |  |
| 7:21.325 | עניני |  |
| 7:21.326 | תשמיש |  |
| 7:21.327 | וזיווג |  |
| 7:21.328 | ולידה | , |
| 7:21.329 | עאכ"ו |  |
| 7:21.330 | שהוא |  |
| 7:21.331 | נעקר |  |
| 7:21.332 | מן |  |
| 7:21.333 | העה"ז |  |
| 7:21.334 | ומן |  |
| 7:21.335 | העה"ב | . |
| 7:21.336 | ואין |  |
| 7:21.337 | לו |  |
| 7:21.338 | חלק |  |
| 7:21.339 | באלהי |  |
| 7:21.340 | ישראל | . |
| 7:21.341 | כי |  |
| 7:21.342 | איככה |  |
| 7:21.343 | ניחס |  |
| 7:21.344 | לאלוהינו |  |
| 7:21.345 | אשר |  |
| 7:21.346 | אותו |  |
| 7:21.347 | נעבוד |  |
| 7:21.348 | פילגש |  |
| 7:21.349 | ואיברי |  |
| 7:21.350 | משגל | . |
| 7:21.351 | יסוד |  |
| 7:21.352 | דהינו |  |
| 7:21.353 | גיד |  |
| 7:21.354 | הערוה |  |
| 7:21.355 | וביצי |  |
| 7:21.356 | זכר |  |
| 7:21.357 | לבשל |  |
| 7:21.358 | הזרע | . |
| 7:21.359 | ונקב |  |
| 7:21.360 | אחוריים |  |
| 7:21.361 | להוציא |  |
| 7:21.362 | הפסולת |  |
| 7:21.363 | להיות |  |
| 7:21.364 | ניזון |  |
| 7:21.365 | ממנה |  |
| 7:21.366 | אל |  |
| 7:21.367 | אחר | ? |
| 7:21.368 | הלא |  |
| 7:21.369 | נבוש |  |
| 7:21.370 | ונכלם |  |
| 7:21.371 | לספר |  |
| 7:21.372 | עניינים |  |
| 7:21.373 | כאלה |  |
| 7:21.374 | בבשר |  |
| 7:21.375 | ודם |  |
| 7:21.376 | כי |  |
| 7:21.377 | אם |  |
| 7:21.378 | ברמז |  |
| 7:21.379 | לעת |  |
| 7:21.380 | הצורך |  |
| 7:21.381 | כגון |  |
| 7:21.382 | שיש |  |
| 7:21.383 | מחלה |  |
| 7:21.384 | באיברי |  |
| 7:21.385 | המשגל |  |
| 7:21.386 | או |  |
| 7:21.387 | בפי |  |
| 7:21.388 | הטבעת | , |
| 7:21.389 | ועם |  |
| 7:21.390 | כל |  |
| 7:21.391 | זה |  |
| 7:21.392 | אין |  |
| 7:21.393 | מזכירים |  |
| 7:21.394 | אותם |  |
| 7:21.395 | אלא |  |
| 7:21.396 | ברמז |  |
| 7:21.397 | וכינוי |  |
| 7:21.398 | ובהצנע |  |
| 7:21.399 | ולא |  |
| 7:21.400 | בפרסום |  |
| 7:21.401 | אבל |  |
| 7:21.402 | שלא |  |
| 7:21.403 | לצורך |  |
| 7:21.404 | הם |  |
| 7:21.405 | מלהזכיר |  |
| 7:21.406 | כלל | . |
| 7:21.407 | כי |  |
| 7:21.408 | זהו |  |
| 7:21.409 | גבלות |  |
| 7:21.410 | הפה |  |
| 7:21.411 | וגנות |  |
| 7:21.412 | גדולה |  |
| 7:21.413 | שלמדה |  |
| 7:21.414 | אותנו |  |
| 7:21.415 | תוה"ק |  |
| 7:21.416 | להשמר |  |
| 7:21.417 | ממנו |  |
| 7:21.418 | כמאמר |  |
| 7:21.419 | רז"ל |  |
| 7:21.420 | בפ"ק |  |
| 7:21.421 | דפסחים |  |
| 7:21.422 | אמר |  |
| 7:21.423 | ר | ' |
| 7:21.424 | יהושע |  |
| 7:21.425 | בן |  |
| 7:21.426 | לוי |  |
| 7:21.427 | לעולם |  |
| 7:21.428 | אל |  |
| 7:21.429 | יוציא |  |
| 7:21.430 | אדם |  |
| 7:21.431 | דבר |  |
| 7:21.432 | מגונה |  |
| 7:21.433 | מפיו |  |
| 7:21.434 | שהרי |  |
| 7:21.435 | עיקם |  |
| 7:21.436 | הכ | ' |
| 7:21.437 | שמנה |  |
| 7:21.438 | אותיות |  |
| 7:21.439 | ולא |  |
| 7:21.440 | הוציא |  |
| 7:21.441 | דבר |  |
| 7:21.442 | מגונה |  |
| 7:21.443 | מפיו |  |
| 7:21.444 | שנ | ' |
| 7:21.445 | ומן |  |
| 7:21.446 | הבהמה |  |
| 7:21.447 | אשר |  |
| 7:21.448 | איננה |  |
| 7:21.449 | טהורה |  |
| 7:21.450 | רב |  |
| 7:21.451 | אשי |  |
| 7:21.452 | אמר |  |
| 7:21.453 | תשע |  |
| 7:21.454 | שנ | ' |
| 7:21.455 | כי |  |
| 7:21.456 | יהיה |  |
| 7:21.457 | בך |  |
| 7:21.458 | איש |  |
| 7:21.459 | אשר |  |
| 7:21.460 | לא |  |
| 7:21.461 | יהיה |  |
| 7:21.462 | טהור |  |
| 7:21.463 | מקרה |  |
| 7:21.464 | לילה | . |
| 7:21.465 | רבינא |  |
| 7:21.466 | אמר |  |
| 7:21.467 | עשר |  |
| 7:21.468 | וי"ו |  |
| 7:21.469 | דטהור | . |
| 7:21.470 | רב |  |
| 7:21.471 | אמר |  |
| 7:21.472 | שש |  |
| 7:21.473 | עשרה |  |
| 7:21.474 | שנ | ' |
| 7:21.475 | כי |  |
| 7:21.476 | אמר |  |
| 7:21.477 | מקרה |  |
| 7:21.478 | הוא |  |
| 7:21.479 | בלתי |  |
| 7:21.480 | טהור |  |
| 7:21.481 | הוא |  |
| 7:21.482 | כי |  |
| 7:21.483 | לא |  |
| 7:21.484 | טהור | . |
| 7:21.485 | תנא |  |
| 7:21.486 | דבי |  |
| 7:21.487 | ר | ' |
| 7:21.488 | ישמעאל |  |
| 7:21.489 | לעולם |  |
| 7:21.490 | יספר |  |
| 7:21.491 | אדם |  |
| 7:21.492 | בלשון |  |
| 7:21.493 | נקייה |  |
| 7:21.494 | שהרי |  |
| 7:21.495 | בזב |  |
| 7:21.496 | וכו | '. |
| 7:21.497 | ואומר |  |
| 7:21.498 | ותבחר |  |
| 7:21.499 | לשון |  |
| 7:21.500 | ערומים | . |
| 7:21.501 | ואומר |  |
| 7:21.502 | וזעת |  |
| 7:21.503 | שפתי |  |
| 7:21.504 | ברור |  |
| 7:21.505 | מללו |  |
| 7:21.506 | וכו | '. |
| 7:21.507 | ובכל |  |
| 7:21.508 | התורה |  |
| 7:21.509 | ובדברי |  |
| 7:21.510 | הנביאים |  |
| 7:21.511 | ורז"ל |  |
| 7:21.512 | בתלמוד |  |
| 7:21.513 | ובמדרשים |  |
| 7:21.514 | לא |  |
| 7:21.515 | מצינו |  |
| 7:21.516 | שכינו |  |
| 7:21.517 | להשפעת |  |
| 7:21.518 | השי"ת |  |
| 7:21.519 | בשם |  |
| 7:21.520 | זכר |  |
| 7:21.521 | ונקבה |  |
| 7:21.522 | וזיווג |  |
| 7:21.523 | אלא |  |
| 7:21.524 | בלשון |  |
| 7:21.525 | נתינה |  |
| 7:21.526 | או |  |
| 7:21.527 | פתיחה |  |
| 7:21.528 | ויתן |  |
| 7:21.529 | לך |  |
| 7:21.530 | האלהים | . |
| 7:21.531 | נותן |  |
| 7:21.532 | לך |  |
| 7:21.533 | כח |  |
| 7:21.534 | לעשות |  |
| 7:21.535 | חיל | . |
| 7:21.536 | יפתח |  |
| 7:21.537 | ה | ' |
| 7:21.538 | לך |  |
| 7:21.539 | את |  |
| 7:21.540 | אוצרו |  |
| 7:21.541 | הטוב |  |
| 7:21.542 | נותן |  |
| 7:21.543 | לחם |  |
| 7:21.544 | לכל |  |
| 7:21.545 | בשר | . |
| 7:21.546 | פותח |  |
| 7:21.547 | ידיך |  |
| 7:21.548 | ומשביע |  |
| 7:21.549 | לכל |  |
| 7:21.550 | חי |  |
| 7:21.551 | רצון | . |
| 7:21.552 | ואתה |  |
| 7:21.553 | מחיה |  |
| 7:21.554 | את |  |
| 7:21.555 | כלם |  |
| 7:21.556 | וכיוצא |  |
| 7:21.557 | בלשונות |  |
| 7:21.558 | אלו |  |
| 7:21.559 | אן |  |
| 7:21.560 | מספר | . |
| 7:21.561 | ובדברי |  |
| 7:21.562 | רז"ל |  |
| 7:21.563 | יושב |  |
| 7:21.564 | הקב"ה |  |
| 7:21.565 | וזן |  |
| 7:21.566 | את |  |
| 7:21.567 | העולם |  |
| 7:21.568 | כו | '. |
| 7:21.569 | זן |  |
| 7:21.570 | ומפרניס |  |
| 7:21.571 | לכל | . |
| 7:21.572 | עורך |  |
| 7:21.573 | שולחן | , |
| 7:21.574 | נותן |  |
| 7:21.575 | טרף |  |
| 7:21.576 | וכיוצא | . |
| 7:21.577 | ולמה |  |
| 7:21.578 | זה |  |
| 7:21.579 | נעזוב |  |
| 7:21.580 | מליצת |  |
| 7:21.581 | ספרי |  |
| 7:21.582 | הקדש |  |
| 7:21.583 | שדברו |  |
| 7:21.584 | בלשון |  |
| 7:21.585 | נקייה | . |
| 7:21.586 | ונשתמש |  |
| 7:21.587 | בשלון |  |
| 7:21.588 | מגונה |  |
| 7:21.589 | סכרים |  |
| 7:21.590 | ונקבות |  |
| 7:21.591 | רבות |  |
| 7:21.592 | וזווגים |  |
| 7:21.593 | והטפת |  |
| 7:21.594 | טיפי |  |
| 7:21.595 | הזרע |  |
| 7:21.596 | ממבושי |  |
| 7:21.597 | הזכר |  |
| 7:21.598 | לרחמה |  |
| 7:21.599 | של |  |
| 7:21.600 | נקבה |  |
| 7:21.601 | וקליטה |  |
| 7:21.602 | וקפאון |  |
| 7:21.603 | הזרע |  |
| 7:21.604 | ברחמה |  |
| 7:21.605 | כמפורסם |  |
| 7:21.606 | כל |  |
| 7:21.607 | זה |  |
| 7:21.608 | בספרי |  |
| 7:21.609 | הקבלה |  |
| 7:21.610 | החדשה |  |
| 7:21.611 | הנכריה | . |
| 7:21.612 | ולא |  |
| 7:21.613 | די |  |
| 7:21.614 | כל |  |
| 7:21.615 | זה |  |
| 7:21.616 | אלא |  |
| 7:21.617 | שהמסית |  |
| 7:21.618 | מחבר |  |
| 7:21.619 | הזהר |  |
| 7:21.620 | העיז |  |
| 7:21.621 | פניו |  |
| 7:21.622 | לכנות |  |
| 7:21.623 | המבושים |  |
| 7:21.624 | אשר |  |
| 7:21.625 | חשב |  |
| 7:21.626 | באדון |  |
| 7:21.627 | העולמים |  |
| 7:21.628 | בשם |  |
| 7:21.629 | בית |  |
| 7:21.630 | אלהים | . |
| 7:21.631 | ושער |  |
| 7:21.632 | השמים | , |
| 7:21.633 | כדאיתא |  |
| 7:21.634 | בזהר |  |
| 7:21.635 | ויצא |  |
| 7:21.636 | דף |  |
| 7:21.637 | ק"ן |  |
| 7:21.638 | ע"ב | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

