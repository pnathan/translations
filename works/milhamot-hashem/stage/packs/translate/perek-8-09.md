# Translation pack: milhamot-hashem perek-8-09 (Milchamot ha-Shem 8:11–8:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-09`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **האלהי** before a name → "the divine" (note); never dropped.
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **תימן** → "Yemen".
- **תוגרמה**: ממשלת תוגרמה → "the government of Turkey" (note: Togarmah).
- **המושלמנים** → "the Muslims".
## Terms

- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- עבודה **אחרת** → "other worship" (note first in a unit).
- אל **אחר** → "another god".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **תשבורת** → "mathematics".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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


## Milchamot ha-Shem 8:11

Hebrew (Wikisource): ומ"מ אפשר שלא ירדו לעמקה לדעת את מרורות הפתנים אשר בקרבה, רק עשו כאשר עשינו אנחנו שנים רבות שפירשנו ממנה מפני שקרא עליה ערער. ולא דברנו אודותיה כמצווה עלינו להודיע את חקי האלהים ואת תורותיו ולמי ראוי לעבוד! ולכן חרה אף השי"ת בנו, ויענישינו על מריינו, שהיינו שומעים דברים המביאים לידי טעות, ועל שלא מחינו בכבוד קונינו. הסית השי"ת בנו את ארך הלשון אשר כחזקתו גבה לבו עד להשחית. וישיאו לבו ויצרו הרע להתגולל להוציא דבה רעה על לומדי תורה לשמה על אמתתה במשנה ובתלמוד. שהם מינים גמורים ושבבית שני תיקן רבן גמליאל על ידי שמואל הקטן ברכת המינים, ושכל התלמידים הלומדים משנה והרמב"ם יצאו לתרבות רעה על שאין לומדים זהר ותיקונים! ומיום ליום יביע אומר מלשינות ושקר לפני שרי הממשלה, שהמה אלה מכתבי עמל. ורצונם למרוד בממשלה הרוממה ולחסות תחת ממשלה אנגליז, ולהמריד את יושבי הארץ לחסות בממשלה אחרת נכריה ועל ידי רכילותו זאת בקש הואלי (המושל) לנתשם מעל אדמתם אל ארץ אחרת, ויתן אותנו בבית הסהר יותר מחדש ימים כשלשים איש, ויפל פחד ורעדה בלב העם מהדר גאונו ועריצותו! ושוב המציא במזימותיו רכילות אחרת במכתבים אל האימאם יחיא אלמתוכל יר"ה. שקצת מן היהודים רוצים להתנשא על המושלמנים וללבוש בגדי כבוד, ולרכוב על סוסים. ומבקשים בתי ספר ללמוד שפה תורכית ושאר לשונות כדי להתגדל על האיסלאם. ואלמלא אני מלמדם ומדריכם בדרכי כניעת היהדות ואיך הם מחוייבים לקבל עליהם עול האיסלאם. כי עתה פרצו ועלו על האיסלאם. ועל ידי זה זכה ליתן לו רשיון מן האימאם יר"ה לרדות בעם ה' אלה ולעשות כחפצו, אך לימודי שפת תוגרמה וידיעת גלילות הארץ וגבולותיה וחשבון ותשבורת וכיוצא לא אסר אותן האימאם יר"ה, כי הוא איש משכיל ונבון, ובמכתבו כתב שאין למנוע הלימודים האלה מנערי בני ישראל. ובעל הלשון הרחיב עוד את פיו לפני אחד השופטים של האימאם בפנינו ג"כ להלשין עלינו. כי אנחנו עוזבי דת הישראלית, ולומדים בספרי פילסופי יון, ומכחישים אמונת מציאות השדים. ובאמת לא נמצא אצלינו שום ספר מספרי יון, ולא זולתם מספרי הפילסופים החיצונים כלל! ורק זדון לבו השיאו להלשין עלינו שקר, והרבה שנאת חנם במחנה העברים על לא דבר! ויוסף עוד לחטוא לשלח מכתבים אל כל כפרי ועיירות התימן להוציא ולפרסם את הרבה הרעה הזאת עלינו בשם מינים! וערום יערים לעשות נוסחי דבות רעות לאמר, הן כבר עשינו משפט ברשעים האלה ונביא אותם אל בית הכלא, ועליכם המצוה למהר במכתבי קובלנא אל האימאם יר"ה, כי אלה מקלקלי דתינו ומפסידי אמונתינו, מהרו, חושו, חושו, להגדיל המדורה עליהם! ולא היה כותבם בכתב ידו, רק מוסרם לילדים מקרוביו כדי להעתיקם בכתב ידם, ושולחם לכל ערי התימן שלא בכתב ידו, כדי שיאמר לא פעלתי און.

| id | word | punctuation after |
|---|---|---|
| 8:11.1 | ומ"מ |  |
| 8:11.2 | אפשר |  |
| 8:11.3 | שלא |  |
| 8:11.4 | ירדו |  |
| 8:11.5 | לעמקה |  |
| 8:11.6 | לדעת |  |
| 8:11.7 | את |  |
| 8:11.8 | מרורות |  |
| 8:11.9 | הפתנים |  |
| 8:11.10 | אשר |  |
| 8:11.11 | בקרבה | , |
| 8:11.12 | רק |  |
| 8:11.13 | עשו |  |
| 8:11.14 | כאשר |  |
| 8:11.15 | עשינו |  |
| 8:11.16 | אנחנו |  |
| 8:11.17 | שנים |  |
| 8:11.18 | רבות |  |
| 8:11.19 | שפירשנו |  |
| 8:11.20 | ממנה |  |
| 8:11.21 | מפני |  |
| 8:11.22 | שקרא |  |
| 8:11.23 | עליה |  |
| 8:11.24 | ערער | . |
| 8:11.25 | ולא |  |
| 8:11.26 | דברנו |  |
| 8:11.27 | אודותיה |  |
| 8:11.28 | כמצווה |  |
| 8:11.29 | עלינו |  |
| 8:11.30 | להודיע |  |
| 8:11.31 | את |  |
| 8:11.32 | חקי |  |
| 8:11.33 | האלהים |  |
| 8:11.34 | ואת |  |
| 8:11.35 | תורותיו |  |
| 8:11.36 | ולמי |  |
| 8:11.37 | ראוי |  |
| 8:11.38 | לעבוד | ! |
| 8:11.39 | ולכן |  |
| 8:11.40 | חרה |  |
| 8:11.41 | אף |  |
| 8:11.42 | השי"ת |  |
| 8:11.43 | בנו | , |
| 8:11.44 | ויענישינו |  |
| 8:11.45 | על |  |
| 8:11.46 | מריינו | , |
| 8:11.47 | שהיינו |  |
| 8:11.48 | שומעים |  |
| 8:11.49 | דברים |  |
| 8:11.50 | המביאים |  |
| 8:11.51 | לידי |  |
| 8:11.52 | טעות | , |
| 8:11.53 | ועל |  |
| 8:11.54 | שלא |  |
| 8:11.55 | מחינו |  |
| 8:11.56 | בכבוד |  |
| 8:11.57 | קונינו | . |
| 8:11.58 | הסית |  |
| 8:11.59 | השי"ת |  |
| 8:11.60 | בנו |  |
| 8:11.61 | את |  |
| 8:11.62 | ארך |  |
| 8:11.63 | הלשון |  |
| 8:11.64 | אשר |  |
| 8:11.65 | כחזקתו |  |
| 8:11.66 | גבה |  |
| 8:11.67 | לבו |  |
| 8:11.68 | עד |  |
| 8:11.69 | להשחית | . |
| 8:11.70 | וישיאו |  |
| 8:11.71 | לבו |  |
| 8:11.72 | ויצרו |  |
| 8:11.73 | הרע |  |
| 8:11.74 | להתגולל |  |
| 8:11.75 | להוציא |  |
| 8:11.76 | דבה |  |
| 8:11.77 | רעה |  |
| 8:11.78 | על |  |
| 8:11.79 | לומדי |  |
| 8:11.80 | תורה |  |
| 8:11.81 | לשמה |  |
| 8:11.82 | על |  |
| 8:11.83 | אמתתה |  |
| 8:11.84 | במשנה |  |
| 8:11.85 | ובתלמוד | . |
| 8:11.86 | שהם |  |
| 8:11.87 | מינים |  |
| 8:11.88 | גמורים |  |
| 8:11.89 | ושבבית |  |
| 8:11.90 | שני |  |
| 8:11.91 | תיקן |  |
| 8:11.92 | רבן |  |
| 8:11.93 | גמליאל |  |
| 8:11.94 | על |  |
| 8:11.95 | ידי |  |
| 8:11.96 | שמואל |  |
| 8:11.97 | הקטן |  |
| 8:11.98 | ברכת |  |
| 8:11.99 | המינים | , |
| 8:11.100 | ושכל |  |
| 8:11.101 | התלמידים |  |
| 8:11.102 | הלומדים |  |
| 8:11.103 | משנה |  |
| 8:11.104 | והרמב"ם |  |
| 8:11.105 | יצאו |  |
| 8:11.106 | לתרבות |  |
| 8:11.107 | רעה |  |
| 8:11.108 | על |  |
| 8:11.109 | שאין |  |
| 8:11.110 | לומדים |  |
| 8:11.111 | זהר |  |
| 8:11.112 | ותיקונים | ! |
| 8:11.113 | ומיום |  |
| 8:11.114 | ליום |  |
| 8:11.115 | יביע |  |
| 8:11.116 | אומר |  |
| 8:11.117 | מלשינות |  |
| 8:11.118 | ושקר |  |
| 8:11.119 | לפני |  |
| 8:11.120 | שרי |  |
| 8:11.121 | הממשלה | , |
| 8:11.122 | שהמה |  |
| 8:11.123 | אלה |  |
| 8:11.124 | מכתבי |  |
| 8:11.125 | עמל | . |
| 8:11.126 | ורצונם |  |
| 8:11.127 | למרוד |  |
| 8:11.128 | בממשלה |  |
| 8:11.129 | הרוממה |  |
| 8:11.130 | ולחסות |  |
| 8:11.131 | תחת |  |
| 8:11.132 | ממשלה |  |
| 8:11.133 | אנגליז | , |
| 8:11.134 | ולהמריד |  |
| 8:11.135 | את |  |
| 8:11.136 | יושבי |  |
| 8:11.137 | הארץ |  |
| 8:11.138 | לחסות |  |
| 8:11.139 | בממשלה |  |
| 8:11.140 | אחרת |  |
| 8:11.141 | נכריה |  |
| 8:11.142 | ועל |  |
| 8:11.143 | ידי |  |
| 8:11.144 | רכילותו |  |
| 8:11.145 | זאת |  |
| 8:11.146 | בקש |  |
| 8:11.147 | הואלי |  |
| 8:11.148 | המושל | ) |
| 8:11.149 | לנתשם |  |
| 8:11.150 | מעל |  |
| 8:11.151 | אדמתם |  |
| 8:11.152 | אל |  |
| 8:11.153 | ארץ |  |
| 8:11.154 | אחרת | , |
| 8:11.155 | ויתן |  |
| 8:11.156 | אותנו |  |
| 8:11.157 | בבית |  |
| 8:11.158 | הסהר |  |
| 8:11.159 | יותר |  |
| 8:11.160 | מחדש |  |
| 8:11.161 | ימים |  |
| 8:11.162 | כשלשים |  |
| 8:11.163 | איש | , |
| 8:11.164 | ויפל |  |
| 8:11.165 | פחד |  |
| 8:11.166 | ורעדה |  |
| 8:11.167 | בלב |  |
| 8:11.168 | העם |  |
| 8:11.169 | מהדר |  |
| 8:11.170 | גאונו |  |
| 8:11.171 | ועריצותו | ! |
| 8:11.172 | ושוב |  |
| 8:11.173 | המציא |  |
| 8:11.174 | במזימותיו |  |
| 8:11.175 | רכילות |  |
| 8:11.176 | אחרת |  |
| 8:11.177 | במכתבים |  |
| 8:11.178 | אל |  |
| 8:11.179 | האימאם |  |
| 8:11.180 | יחיא |  |
| 8:11.181 | אלמתוכל |  |
| 8:11.182 | יר"ה | . |
| 8:11.183 | שקצת |  |
| 8:11.184 | מן |  |
| 8:11.185 | היהודים |  |
| 8:11.186 | רוצים |  |
| 8:11.187 | להתנשא |  |
| 8:11.188 | על |  |
| 8:11.189 | המושלמנים |  |
| 8:11.190 | וללבוש |  |
| 8:11.191 | בגדי |  |
| 8:11.192 | כבוד | , |
| 8:11.193 | ולרכוב |  |
| 8:11.194 | על |  |
| 8:11.195 | סוסים | . |
| 8:11.196 | ומבקשים |  |
| 8:11.197 | בתי |  |
| 8:11.198 | ספר |  |
| 8:11.199 | ללמוד |  |
| 8:11.200 | שפה |  |
| 8:11.201 | תורכית |  |
| 8:11.202 | ושאר |  |
| 8:11.203 | לשונות |  |
| 8:11.204 | כדי |  |
| 8:11.205 | להתגדל |  |
| 8:11.206 | על |  |
| 8:11.207 | האיסלאם | . |
| 8:11.208 | ואלמלא |  |
| 8:11.209 | אני |  |
| 8:11.210 | מלמדם |  |
| 8:11.211 | ומדריכם |  |
| 8:11.212 | בדרכי |  |
| 8:11.213 | כניעת |  |
| 8:11.214 | היהדות |  |
| 8:11.215 | ואיך |  |
| 8:11.216 | הם |  |
| 8:11.217 | מחוייבים |  |
| 8:11.218 | לקבל |  |
| 8:11.219 | עליהם |  |
| 8:11.220 | עול |  |
| 8:11.221 | האיסלאם | . |
| 8:11.222 | כי |  |
| 8:11.223 | עתה |  |
| 8:11.224 | פרצו |  |
| 8:11.225 | ועלו |  |
| 8:11.226 | על |  |
| 8:11.227 | האיסלאם | . |
| 8:11.228 | ועל |  |
| 8:11.229 | ידי |  |
| 8:11.230 | זה |  |
| 8:11.231 | זכה |  |
| 8:11.232 | ליתן |  |
| 8:11.233 | לו |  |
| 8:11.234 | רשיון |  |
| 8:11.235 | מן |  |
| 8:11.236 | האימאם |  |
| 8:11.237 | יר"ה |  |
| 8:11.238 | לרדות |  |
| 8:11.239 | בעם |  |
| 8:11.240 | ה | ' |
| 8:11.241 | אלה |  |
| 8:11.242 | ולעשות |  |
| 8:11.243 | כחפצו | , |
| 8:11.244 | אך |  |
| 8:11.245 | לימודי |  |
| 8:11.246 | שפת |  |
| 8:11.247 | תוגרמה |  |
| 8:11.248 | וידיעת |  |
| 8:11.249 | גלילות |  |
| 8:11.250 | הארץ |  |
| 8:11.251 | וגבולותיה |  |
| 8:11.252 | וחשבון |  |
| 8:11.253 | ותשבורת |  |
| 8:11.254 | וכיוצא |  |
| 8:11.255 | לא |  |
| 8:11.256 | אסר |  |
| 8:11.257 | אותן |  |
| 8:11.258 | האימאם |  |
| 8:11.259 | יר"ה | , |
| 8:11.260 | כי |  |
| 8:11.261 | הוא |  |
| 8:11.262 | איש |  |
| 8:11.263 | משכיל |  |
| 8:11.264 | ונבון | , |
| 8:11.265 | ובמכתבו |  |
| 8:11.266 | כתב |  |
| 8:11.267 | שאין |  |
| 8:11.268 | למנוע |  |
| 8:11.269 | הלימודים |  |
| 8:11.270 | האלה |  |
| 8:11.271 | מנערי |  |
| 8:11.272 | בני |  |
| 8:11.273 | ישראל | . |
| 8:11.274 | ובעל |  |
| 8:11.275 | הלשון |  |
| 8:11.276 | הרחיב |  |
| 8:11.277 | עוד |  |
| 8:11.278 | את |  |
| 8:11.279 | פיו |  |
| 8:11.280 | לפני |  |
| 8:11.281 | אחד |  |
| 8:11.282 | השופטים |  |
| 8:11.283 | של |  |
| 8:11.284 | האימאם |  |
| 8:11.285 | בפנינו |  |
| 8:11.286 | ג"כ |  |
| 8:11.287 | להלשין |  |
| 8:11.288 | עלינו | . |
| 8:11.289 | כי |  |
| 8:11.290 | אנחנו |  |
| 8:11.291 | עוזבי |  |
| 8:11.292 | דת |  |
| 8:11.293 | הישראלית | , |
| 8:11.294 | ולומדים |  |
| 8:11.295 | בספרי |  |
| 8:11.296 | פילסופי |  |
| 8:11.297 | יון | , |
| 8:11.298 | ומכחישים |  |
| 8:11.299 | אמונת |  |
| 8:11.300 | מציאות |  |
| 8:11.301 | השדים | . |
| 8:11.302 | ובאמת |  |
| 8:11.303 | לא |  |
| 8:11.304 | נמצא |  |
| 8:11.305 | אצלינו |  |
| 8:11.306 | שום |  |
| 8:11.307 | ספר |  |
| 8:11.308 | מספרי |  |
| 8:11.309 | יון | , |
| 8:11.310 | ולא |  |
| 8:11.311 | זולתם |  |
| 8:11.312 | מספרי |  |
| 8:11.313 | הפילסופים |  |
| 8:11.314 | החיצונים |  |
| 8:11.315 | כלל | ! |
| 8:11.316 | ורק |  |
| 8:11.317 | זדון |  |
| 8:11.318 | לבו |  |
| 8:11.319 | השיאו |  |
| 8:11.320 | להלשין |  |
| 8:11.321 | עלינו |  |
| 8:11.322 | שקר | , |
| 8:11.323 | והרבה |  |
| 8:11.324 | שנאת |  |
| 8:11.325 | חנם |  |
| 8:11.326 | במחנה |  |
| 8:11.327 | העברים |  |
| 8:11.328 | על |  |
| 8:11.329 | לא |  |
| 8:11.330 | דבר | ! |
| 8:11.331 | ויוסף |  |
| 8:11.332 | עוד |  |
| 8:11.333 | לחטוא |  |
| 8:11.334 | לשלח |  |
| 8:11.335 | מכתבים |  |
| 8:11.336 | אל |  |
| 8:11.337 | כל |  |
| 8:11.338 | כפרי |  |
| 8:11.339 | ועיירות |  |
| 8:11.340 | התימן |  |
| 8:11.341 | להוציא |  |
| 8:11.342 | ולפרסם |  |
| 8:11.343 | את |  |
| 8:11.344 | הרבה |  |
| 8:11.345 | הרעה |  |
| 8:11.346 | הזאת |  |
| 8:11.347 | עלינו |  |
| 8:11.348 | בשם |  |
| 8:11.349 | מינים | ! |
| 8:11.350 | וערום |  |
| 8:11.351 | יערים |  |
| 8:11.352 | לעשות |  |
| 8:11.353 | נוסחי |  |
| 8:11.354 | דבות |  |
| 8:11.355 | רעות |  |
| 8:11.356 | לאמר | , |
| 8:11.357 | הן |  |
| 8:11.358 | כבר |  |
| 8:11.359 | עשינו |  |
| 8:11.360 | משפט |  |
| 8:11.361 | ברשעים |  |
| 8:11.362 | האלה |  |
| 8:11.363 | ונביא |  |
| 8:11.364 | אותם |  |
| 8:11.365 | אל |  |
| 8:11.366 | בית |  |
| 8:11.367 | הכלא | , |
| 8:11.368 | ועליכם |  |
| 8:11.369 | המצוה |  |
| 8:11.370 | למהר |  |
| 8:11.371 | במכתבי |  |
| 8:11.372 | קובלנא |  |
| 8:11.373 | אל |  |
| 8:11.374 | האימאם |  |
| 8:11.375 | יר"ה | , |
| 8:11.376 | כי |  |
| 8:11.377 | אלה |  |
| 8:11.378 | מקלקלי |  |
| 8:11.379 | דתינו |  |
| 8:11.380 | ומפסידי |  |
| 8:11.381 | אמונתינו | , |
| 8:11.382 | מהרו | , |
| 8:11.383 | חושו | , |
| 8:11.384 | חושו | , |
| 8:11.385 | להגדיל |  |
| 8:11.386 | המדורה |  |
| 8:11.387 | עליהם | ! |
| 8:11.388 | ולא |  |
| 8:11.389 | היה |  |
| 8:11.390 | כותבם |  |
| 8:11.391 | בכתב |  |
| 8:11.392 | ידו | , |
| 8:11.393 | רק |  |
| 8:11.394 | מוסרם |  |
| 8:11.395 | לילדים |  |
| 8:11.396 | מקרוביו |  |
| 8:11.397 | כדי |  |
| 8:11.398 | להעתיקם |  |
| 8:11.399 | בכתב |  |
| 8:11.400 | ידם | , |
| 8:11.401 | ושולחם |  |
| 8:11.402 | לכל |  |
| 8:11.403 | ערי |  |
| 8:11.404 | התימן |  |
| 8:11.405 | שלא |  |
| 8:11.406 | בכתב |  |
| 8:11.407 | ידו | , |
| 8:11.408 | כדי |  |
| 8:11.409 | שיאמר |  |
| 8:11.410 | לא |  |
| 8:11.411 | פעלתי |  |
| 8:11.412 | און | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

