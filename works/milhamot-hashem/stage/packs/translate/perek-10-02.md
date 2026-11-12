# Translation pack: milhamot-hashem perek-10-02 (Milchamot ha-Shem 10:2–10:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-02`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- השם **הנכבד** → "the glorious Name".
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **תימן** → "Yemen".
- **בא"י** → "in the Land of Israel".

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
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **תא חזי** → "come and see".
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


## Milchamot ha-Shem 10:2

Hebrew (Wikisource): קמא) על כן יאמר כל מי שנגע יראת השי"ת ירחיק עצמו כמטחוי קשת מכל המנהגים ודינים מחודשים על פי הקבלה החדשה השקרית אשר מקרוב באה ונתפשטה בארץ התימן כעדות מהר"י צאחרי בספר צידה לדרך, לפי שכולם עוקרים וחורסים פנות תורה ה"ק ועמודיה יתפלצון. כי בכל ברכותיהם ותפלותיהם כשמזכירים את השם הנכבד והנורא הם מכוונים לאל אחר קצר אפים (ז"א) בשתוף שאר הפרצופים שלמעלה ממנו כמבואר בספריהם במקומות אין מספר. ולא מהמנהגים שנתחדשו על פי הקבלה הנכרית החדשה בלבד צריך להזהר. אלא גם מכל הדברים שהחזירונו רז"ל. והא לך קצת מהם ספר תורה ותפלין ומזוזה שנכתבו על ידי ישראל האדוק בקבלה הלזו פסולים כי השמות שבהם נתקדשו לשם אל אחר קצר אפים (ז"א) והרי כולן כשמות האמורות במיכה שהן חול כמו שכ' הרמב"ם בסוף פרק ששי מהלכות יסודי התורה וקרוב לומר שדין הוא שישרפו כמו שכתבו הפוסקים שס"ת ותפלין ומזוזה שכתבן מין ישרפו לא יאכלו בשר משחיטת מקובל כי בעת שמזכיר את השם בברכתו על השחיטה דעתו על יראתו (ומטעם זה קבעו קדמונינו הי"ג עיקרים בתוך ביאורי וחידושי הלבות שחיטה, שכל הבא ליטול קבלה לשחוט לעצמו או לאחרים צריך להיות מאמין בהן ויודען היטב, ואם אינו יודען וכ"ש אם אינו מאמין בהם כמו המקובלים החדשים שאין מאמינים בארבעה עיקרים מהן לא יתנו לו רשות לשחוט וכל שכן לקדש ולגרש, וגיטו וקידושו אינן כלום, ונמצא מרבה ממזרים בישראל). אם שליח צבור אדוק בקבלה החדשה ובקי בה לא יענה אחריו אמן וקדוש וברוך ה' המבורך. שכל כוונתו לאל אחר זעיר אנפין. אם מקדש אשה או מגרשה לא יהיו העדים ממאמיני קבלה החדשה כי הקדושין והגט בטלים מפני שהם פסולי עדות דאוריתא, ונמצאת אשת איש נשאת לאחר על ידי גט בטל. וכבר הודעתיך בסי' פ"ג בשם משנת חסידים דבשעת אמירת עשר מכות בליל פסח הם מנסכים ושופכים מכוס יינם לתוך כלי שבור להקליפה הנקראת ארור כדי לתקן המלכות מדם נדתה, ונקתה ונזרעה זרע וכו' ואם כן יינם יין נסך ואסור בהנאה. ועוד הרבה עניינים המביאים לידי מינות ועבודה זרה. וצריך להזהר מהם וממאכלם ומשתיהם פו יטמאו נשמתו באיסורי מאכלות.

| id | word | punctuation after |
|---|---|---|
| 10:2.1 | קמא | ) |
| 10:2.2 | על |  |
| 10:2.3 | כן |  |
| 10:2.4 | יאמר |  |
| 10:2.5 | כל |  |
| 10:2.6 | מי |  |
| 10:2.7 | שנגע |  |
| 10:2.8 | יראת |  |
| 10:2.9 | השי"ת |  |
| 10:2.10 | ירחיק |  |
| 10:2.11 | עצמו |  |
| 10:2.12 | כמטחוי |  |
| 10:2.13 | קשת |  |
| 10:2.14 | מכל |  |
| 10:2.15 | המנהגים |  |
| 10:2.16 | ודינים |  |
| 10:2.17 | מחודשים |  |
| 10:2.18 | על |  |
| 10:2.19 | פי |  |
| 10:2.20 | הקבלה |  |
| 10:2.21 | החדשה |  |
| 10:2.22 | השקרית |  |
| 10:2.23 | אשר |  |
| 10:2.24 | מקרוב |  |
| 10:2.25 | באה |  |
| 10:2.26 | ונתפשטה |  |
| 10:2.27 | בארץ |  |
| 10:2.28 | התימן |  |
| 10:2.29 | כעדות |  |
| 10:2.30 | מהר"י |  |
| 10:2.31 | צאחרי |  |
| 10:2.32 | בספר |  |
| 10:2.33 | צידה |  |
| 10:2.34 | לדרך | , |
| 10:2.35 | לפי |  |
| 10:2.36 | שכולם |  |
| 10:2.37 | עוקרים |  |
| 10:2.38 | וחורסים |  |
| 10:2.39 | פנות |  |
| 10:2.40 | תורה |  |
| 10:2.41 | ה"ק |  |
| 10:2.42 | ועמודיה |  |
| 10:2.43 | יתפלצון | . |
| 10:2.44 | כי |  |
| 10:2.45 | בכל |  |
| 10:2.46 | ברכותיהם |  |
| 10:2.47 | ותפלותיהם |  |
| 10:2.48 | כשמזכירים |  |
| 10:2.49 | את |  |
| 10:2.50 | השם |  |
| 10:2.51 | הנכבד |  |
| 10:2.52 | והנורא |  |
| 10:2.53 | הם |  |
| 10:2.54 | מכוונים |  |
| 10:2.55 | לאל |  |
| 10:2.56 | אחר |  |
| 10:2.57 | קצר |  |
| 10:2.58 | אפים |  |
| 10:2.59 | ז"א | ) |
| 10:2.60 | בשתוף |  |
| 10:2.61 | שאר |  |
| 10:2.62 | הפרצופים |  |
| 10:2.63 | שלמעלה |  |
| 10:2.64 | ממנו |  |
| 10:2.65 | כמבואר |  |
| 10:2.66 | בספריהם |  |
| 10:2.67 | במקומות |  |
| 10:2.68 | אין |  |
| 10:2.69 | מספר | . |
| 10:2.70 | ולא |  |
| 10:2.71 | מהמנהגים |  |
| 10:2.72 | שנתחדשו |  |
| 10:2.73 | על |  |
| 10:2.74 | פי |  |
| 10:2.75 | הקבלה |  |
| 10:2.76 | הנכרית |  |
| 10:2.77 | החדשה |  |
| 10:2.78 | בלבד |  |
| 10:2.79 | צריך |  |
| 10:2.80 | להזהר | . |
| 10:2.81 | אלא |  |
| 10:2.82 | גם |  |
| 10:2.83 | מכל |  |
| 10:2.84 | הדברים |  |
| 10:2.85 | שהחזירונו |  |
| 10:2.86 | רז"ל | . |
| 10:2.87 | והא |  |
| 10:2.88 | לך |  |
| 10:2.89 | קצת |  |
| 10:2.90 | מהם |  |
| 10:2.91 | ספר |  |
| 10:2.92 | תורה |  |
| 10:2.93 | ותפלין |  |
| 10:2.94 | ומזוזה |  |
| 10:2.95 | שנכתבו |  |
| 10:2.96 | על |  |
| 10:2.97 | ידי |  |
| 10:2.98 | ישראל |  |
| 10:2.99 | האדוק |  |
| 10:2.100 | בקבלה |  |
| 10:2.101 | הלזו |  |
| 10:2.102 | פסולים |  |
| 10:2.103 | כי |  |
| 10:2.104 | השמות |  |
| 10:2.105 | שבהם |  |
| 10:2.106 | נתקדשו |  |
| 10:2.107 | לשם |  |
| 10:2.108 | אל |  |
| 10:2.109 | אחר |  |
| 10:2.110 | קצר |  |
| 10:2.111 | אפים |  |
| 10:2.112 | ז"א | ) |
| 10:2.113 | והרי |  |
| 10:2.114 | כולן |  |
| 10:2.115 | כשמות |  |
| 10:2.116 | האמורות |  |
| 10:2.117 | במיכה |  |
| 10:2.118 | שהן |  |
| 10:2.119 | חול |  |
| 10:2.120 | כמו |  |
| 10:2.121 | שכ | ' |
| 10:2.122 | הרמב"ם |  |
| 10:2.123 | בסוף |  |
| 10:2.124 | פרק |  |
| 10:2.125 | ששי |  |
| 10:2.126 | מהלכות |  |
| 10:2.127 | יסודי |  |
| 10:2.128 | התורה |  |
| 10:2.129 | וקרוב |  |
| 10:2.130 | לומר |  |
| 10:2.131 | שדין |  |
| 10:2.132 | הוא |  |
| 10:2.133 | שישרפו |  |
| 10:2.134 | כמו |  |
| 10:2.135 | שכתבו |  |
| 10:2.136 | הפוסקים |  |
| 10:2.137 | שס"ת |  |
| 10:2.138 | ותפלין |  |
| 10:2.139 | ומזוזה |  |
| 10:2.140 | שכתבן |  |
| 10:2.141 | מין |  |
| 10:2.142 | ישרפו |  |
| 10:2.143 | לא |  |
| 10:2.144 | יאכלו |  |
| 10:2.145 | בשר |  |
| 10:2.146 | משחיטת |  |
| 10:2.147 | מקובל |  |
| 10:2.148 | כי |  |
| 10:2.149 | בעת |  |
| 10:2.150 | שמזכיר |  |
| 10:2.151 | את |  |
| 10:2.152 | השם |  |
| 10:2.153 | בברכתו |  |
| 10:2.154 | על |  |
| 10:2.155 | השחיטה |  |
| 10:2.156 | דעתו |  |
| 10:2.157 | על |  |
| 10:2.158 | יראתו |  |
| 10:2.159 | ומטעם |  |
| 10:2.160 | זה |  |
| 10:2.161 | קבעו |  |
| 10:2.162 | קדמונינו |  |
| 10:2.163 | הי"ג |  |
| 10:2.164 | עיקרים |  |
| 10:2.165 | בתוך |  |
| 10:2.166 | ביאורי |  |
| 10:2.167 | וחידושי |  |
| 10:2.168 | הלבות |  |
| 10:2.169 | שחיטה | , |
| 10:2.170 | שכל |  |
| 10:2.171 | הבא |  |
| 10:2.172 | ליטול |  |
| 10:2.173 | קבלה |  |
| 10:2.174 | לשחוט |  |
| 10:2.175 | לעצמו |  |
| 10:2.176 | או |  |
| 10:2.177 | לאחרים |  |
| 10:2.178 | צריך |  |
| 10:2.179 | להיות |  |
| 10:2.180 | מאמין |  |
| 10:2.181 | בהן |  |
| 10:2.182 | ויודען |  |
| 10:2.183 | היטב | , |
| 10:2.184 | ואם |  |
| 10:2.185 | אינו |  |
| 10:2.186 | יודען |  |
| 10:2.187 | וכ"ש |  |
| 10:2.188 | אם |  |
| 10:2.189 | אינו |  |
| 10:2.190 | מאמין |  |
| 10:2.191 | בהם |  |
| 10:2.192 | כמו |  |
| 10:2.193 | המקובלים |  |
| 10:2.194 | החדשים |  |
| 10:2.195 | שאין |  |
| 10:2.196 | מאמינים |  |
| 10:2.197 | בארבעה |  |
| 10:2.198 | עיקרים |  |
| 10:2.199 | מהן |  |
| 10:2.200 | לא |  |
| 10:2.201 | יתנו |  |
| 10:2.202 | לו |  |
| 10:2.203 | רשות |  |
| 10:2.204 | לשחוט |  |
| 10:2.205 | וכל |  |
| 10:2.206 | שכן |  |
| 10:2.207 | לקדש |  |
| 10:2.208 | ולגרש | , |
| 10:2.209 | וגיטו |  |
| 10:2.210 | וקידושו |  |
| 10:2.211 | אינן |  |
| 10:2.212 | כלום | , |
| 10:2.213 | ונמצא |  |
| 10:2.214 | מרבה |  |
| 10:2.215 | ממזרים |  |
| 10:2.216 | בישראל | ). |
| 10:2.217 | אם |  |
| 10:2.218 | שליח |  |
| 10:2.219 | צבור |  |
| 10:2.220 | אדוק |  |
| 10:2.221 | בקבלה |  |
| 10:2.222 | החדשה |  |
| 10:2.223 | ובקי |  |
| 10:2.224 | בה |  |
| 10:2.225 | לא |  |
| 10:2.226 | יענה |  |
| 10:2.227 | אחריו |  |
| 10:2.228 | אמן |  |
| 10:2.229 | וקדוש |  |
| 10:2.230 | וברוך |  |
| 10:2.231 | ה | ' |
| 10:2.232 | המבורך | . |
| 10:2.233 | שכל |  |
| 10:2.234 | כוונתו |  |
| 10:2.235 | לאל |  |
| 10:2.236 | אחר |  |
| 10:2.237 | זעיר |  |
| 10:2.238 | אנפין | . |
| 10:2.239 | אם |  |
| 10:2.240 | מקדש |  |
| 10:2.241 | אשה |  |
| 10:2.242 | או |  |
| 10:2.243 | מגרשה |  |
| 10:2.244 | לא |  |
| 10:2.245 | יהיו |  |
| 10:2.246 | העדים |  |
| 10:2.247 | ממאמיני |  |
| 10:2.248 | קבלה |  |
| 10:2.249 | החדשה |  |
| 10:2.250 | כי |  |
| 10:2.251 | הקדושין |  |
| 10:2.252 | והגט |  |
| 10:2.253 | בטלים |  |
| 10:2.254 | מפני |  |
| 10:2.255 | שהם |  |
| 10:2.256 | פסולי |  |
| 10:2.257 | עדות |  |
| 10:2.258 | דאוריתא | , |
| 10:2.259 | ונמצאת |  |
| 10:2.260 | אשת |  |
| 10:2.261 | איש |  |
| 10:2.262 | נשאת |  |
| 10:2.263 | לאחר |  |
| 10:2.264 | על |  |
| 10:2.265 | ידי |  |
| 10:2.266 | גט |  |
| 10:2.267 | בטל | . |
| 10:2.268 | וכבר |  |
| 10:2.269 | הודעתיך |  |
| 10:2.270 | בסי | ' |
| 10:2.271 | פ"ג |  |
| 10:2.272 | בשם |  |
| 10:2.273 | משנת |  |
| 10:2.274 | חסידים |  |
| 10:2.275 | דבשעת |  |
| 10:2.276 | אמירת |  |
| 10:2.277 | עשר |  |
| 10:2.278 | מכות |  |
| 10:2.279 | בליל |  |
| 10:2.280 | פסח |  |
| 10:2.281 | הם |  |
| 10:2.282 | מנסכים |  |
| 10:2.283 | ושופכים |  |
| 10:2.284 | מכוס |  |
| 10:2.285 | יינם |  |
| 10:2.286 | לתוך |  |
| 10:2.287 | כלי |  |
| 10:2.288 | שבור |  |
| 10:2.289 | להקליפה |  |
| 10:2.290 | הנקראת |  |
| 10:2.291 | ארור |  |
| 10:2.292 | כדי |  |
| 10:2.293 | לתקן |  |
| 10:2.294 | המלכות |  |
| 10:2.295 | מדם |  |
| 10:2.296 | נדתה | , |
| 10:2.297 | ונקתה |  |
| 10:2.298 | ונזרעה |  |
| 10:2.299 | זרע |  |
| 10:2.300 | וכו | ' |
| 10:2.301 | ואם |  |
| 10:2.302 | כן |  |
| 10:2.303 | יינם |  |
| 10:2.304 | יין |  |
| 10:2.305 | נסך |  |
| 10:2.306 | ואסור |  |
| 10:2.307 | בהנאה | . |
| 10:2.308 | ועוד |  |
| 10:2.309 | הרבה |  |
| 10:2.310 | עניינים |  |
| 10:2.311 | המביאים |  |
| 10:2.312 | לידי |  |
| 10:2.313 | מינות |  |
| 10:2.314 | ועבודה |  |
| 10:2.315 | זרה | . |
| 10:2.316 | וצריך |  |
| 10:2.317 | להזהר |  |
| 10:2.318 | מהם |  |
| 10:2.319 | וממאכלם |  |
| 10:2.320 | ומשתיהם |  |
| 10:2.321 | פו |  |
| 10:2.322 | יטמאו |  |
| 10:2.323 | נשמתו |  |
| 10:2.324 | באיסורי |  |
| 10:2.325 | מאכלות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

