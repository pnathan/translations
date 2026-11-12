# Translation pack: milhamot-hashem perek-7-05 (Milchamot ha-Shem 7:6–7:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-05`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- אורות **דקים** → "subtle lights".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **קורין** אל → "call upon".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- אין לנו **להחמיר** → "we may not be stringent".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **אחי** → "my brother" (second person kept).
- **חסר / חטר**: "(חסר)" → "(missing)", a lacuna sign; "(חטר)" is a variant of
  it, recorded.

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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
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


## Milchamot ha-Shem 7:6

Hebrew (Wikisource): פג) עוד שם במדרש על פסוק ירא ה' את בני ומלך, מהו מלך, המליכהו עליך ועם שונים אל תתערב. עם אלו שאומרים יש אלוה שני אל תתערב וכו'. אמר ר' אחא כעס הקב"ה על שלמה כשאמר הפסוק הזה. אמר לו דבר של קדוש שמי היית אומר בלשון נוטריקון עם שונים אל תתערב מיד חזר ופירש את הדבר. יש אחד ואין שני גם בן ואח אין לו. אין לו אח ולא בן. אלא שמע ישראל ה' אלהינו ה' אחד עכ"ל ופי' הר' עץ יוסף. יש אלוה שני. כי המינים אומרים איך יתכן פועל טוב ורע יוצא ממקור אחד? לכן הם מאמינים בשניות וכו'. ובילקוט קהלת רמז תתקע"א. ובקהלת רבותי יש אחד ואין שני. אחד זה הקב"ה דכתיב ביה שמע ישראל ה' אלהינו ה' אחד. ואין שני גם בן ואח אין לו. אח מהיכן יש לו? אלא שחיבב את ישראל וקרא אותם בנים ואחים. ואין קץ לכל עמלו שעמל בששת ימי בראשית (פי' עמל הנז' כאן איננו לשון טורח וידיעה. שאין יגיעה וטורח לפני מי שאמר והיה העולם. אלא הוא לשון פעולה, שכן בלשון ישמעאל קורין לפעולה עמל. וכן תפרש לשון הפיטן מתקן הסליחות דאשמורות שאמר חוסה על עמלך כמו פעלך) ולמי אני עמל ומחסר את נפשי מטובה, לא להדבק בו וללכת בדרכיו? וכל מי שאינו עושה כן. גם זה הבל וענין רע הוא עכ"ל לעניינינו.

| id | word | punctuation after |
|---|---|---|
| 7:6.1 | פג | ) |
| 7:6.2 | עוד |  |
| 7:6.3 | שם |  |
| 7:6.4 | במדרש |  |
| 7:6.5 | על |  |
| 7:6.6 | פסוק |  |
| 7:6.7 | ירא |  |
| 7:6.8 | ה | ' |
| 7:6.9 | את |  |
| 7:6.10 | בני |  |
| 7:6.11 | ומלך | , |
| 7:6.12 | מהו |  |
| 7:6.13 | מלך | , |
| 7:6.14 | המליכהו |  |
| 7:6.15 | עליך |  |
| 7:6.16 | ועם |  |
| 7:6.17 | שונים |  |
| 7:6.18 | אל |  |
| 7:6.19 | תתערב | . |
| 7:6.20 | עם |  |
| 7:6.21 | אלו |  |
| 7:6.22 | שאומרים |  |
| 7:6.23 | יש |  |
| 7:6.24 | אלוה |  |
| 7:6.25 | שני |  |
| 7:6.26 | אל |  |
| 7:6.27 | תתערב |  |
| 7:6.28 | וכו | '. |
| 7:6.29 | אמר |  |
| 7:6.30 | ר | ' |
| 7:6.31 | אחא |  |
| 7:6.32 | כעס |  |
| 7:6.33 | הקב"ה |  |
| 7:6.34 | על |  |
| 7:6.35 | שלמה |  |
| 7:6.36 | כשאמר |  |
| 7:6.37 | הפסוק |  |
| 7:6.38 | הזה | . |
| 7:6.39 | אמר |  |
| 7:6.40 | לו |  |
| 7:6.41 | דבר |  |
| 7:6.42 | של |  |
| 7:6.43 | קדוש |  |
| 7:6.44 | שמי |  |
| 7:6.45 | היית |  |
| 7:6.46 | אומר |  |
| 7:6.47 | בלשון |  |
| 7:6.48 | נוטריקון |  |
| 7:6.49 | עם |  |
| 7:6.50 | שונים |  |
| 7:6.51 | אל |  |
| 7:6.52 | תתערב |  |
| 7:6.53 | מיד |  |
| 7:6.54 | חזר |  |
| 7:6.55 | ופירש |  |
| 7:6.56 | את |  |
| 7:6.57 | הדבר | . |
| 7:6.58 | יש |  |
| 7:6.59 | אחד |  |
| 7:6.60 | ואין |  |
| 7:6.61 | שני |  |
| 7:6.62 | גם |  |
| 7:6.63 | בן |  |
| 7:6.64 | ואח |  |
| 7:6.65 | אין |  |
| 7:6.66 | לו | . |
| 7:6.67 | אין |  |
| 7:6.68 | לו |  |
| 7:6.69 | אח |  |
| 7:6.70 | ולא |  |
| 7:6.71 | בן | . |
| 7:6.72 | אלא |  |
| 7:6.73 | שמע |  |
| 7:6.74 | ישראל |  |
| 7:6.75 | ה | ' |
| 7:6.76 | אלהינו |  |
| 7:6.77 | ה | ' |
| 7:6.78 | אחד |  |
| 7:6.79 | עכ"ל |  |
| 7:6.80 | ופי | ' |
| 7:6.81 | הר | ' |
| 7:6.82 | עץ |  |
| 7:6.83 | יוסף | . |
| 7:6.84 | יש |  |
| 7:6.85 | אלוה |  |
| 7:6.86 | שני | . |
| 7:6.87 | כי |  |
| 7:6.88 | המינים |  |
| 7:6.89 | אומרים |  |
| 7:6.90 | איך |  |
| 7:6.91 | יתכן |  |
| 7:6.92 | פועל |  |
| 7:6.93 | טוב |  |
| 7:6.94 | ורע |  |
| 7:6.95 | יוצא |  |
| 7:6.96 | ממקור |  |
| 7:6.97 | אחד | ? |
| 7:6.98 | לכן |  |
| 7:6.99 | הם |  |
| 7:6.100 | מאמינים |  |
| 7:6.101 | בשניות |  |
| 7:6.102 | וכו | '. |
| 7:6.103 | ובילקוט |  |
| 7:6.104 | קהלת |  |
| 7:6.105 | רמז |  |
| 7:6.106 | תתקע"א | . |
| 7:6.107 | ובקהלת |  |
| 7:6.108 | רבותי |  |
| 7:6.109 | יש |  |
| 7:6.110 | אחד |  |
| 7:6.111 | ואין |  |
| 7:6.112 | שני | . |
| 7:6.113 | אחד |  |
| 7:6.114 | זה |  |
| 7:6.115 | הקב"ה |  |
| 7:6.116 | דכתיב |  |
| 7:6.117 | ביה |  |
| 7:6.118 | שמע |  |
| 7:6.119 | ישראל |  |
| 7:6.120 | ה | ' |
| 7:6.121 | אלהינו |  |
| 7:6.122 | ה | ' |
| 7:6.123 | אחד | . |
| 7:6.124 | ואין |  |
| 7:6.125 | שני |  |
| 7:6.126 | גם |  |
| 7:6.127 | בן |  |
| 7:6.128 | ואח |  |
| 7:6.129 | אין |  |
| 7:6.130 | לו | . |
| 7:6.131 | אח |  |
| 7:6.132 | מהיכן |  |
| 7:6.133 | יש |  |
| 7:6.134 | לו | ? |
| 7:6.135 | אלא |  |
| 7:6.136 | שחיבב |  |
| 7:6.137 | את |  |
| 7:6.138 | ישראל |  |
| 7:6.139 | וקרא |  |
| 7:6.140 | אותם |  |
| 7:6.141 | בנים |  |
| 7:6.142 | ואחים | . |
| 7:6.143 | ואין |  |
| 7:6.144 | קץ |  |
| 7:6.145 | לכל |  |
| 7:6.146 | עמלו |  |
| 7:6.147 | שעמל |  |
| 7:6.148 | בששת |  |
| 7:6.149 | ימי |  |
| 7:6.150 | בראשית |  |
| 7:6.151 | פי | ' |
| 7:6.152 | עמל |  |
| 7:6.153 | הנז | ' |
| 7:6.154 | כאן |  |
| 7:6.155 | איננו |  |
| 7:6.156 | לשון |  |
| 7:6.157 | טורח |  |
| 7:6.158 | וידיעה | . |
| 7:6.159 | שאין |  |
| 7:6.160 | יגיעה |  |
| 7:6.161 | וטורח |  |
| 7:6.162 | לפני |  |
| 7:6.163 | מי |  |
| 7:6.164 | שאמר |  |
| 7:6.165 | והיה |  |
| 7:6.166 | העולם | . |
| 7:6.167 | אלא |  |
| 7:6.168 | הוא |  |
| 7:6.169 | לשון |  |
| 7:6.170 | פעולה | , |
| 7:6.171 | שכן |  |
| 7:6.172 | בלשון |  |
| 7:6.173 | ישמעאל |  |
| 7:6.174 | קורין |  |
| 7:6.175 | לפעולה |  |
| 7:6.176 | עמל | . |
| 7:6.177 | וכן |  |
| 7:6.178 | תפרש |  |
| 7:6.179 | לשון |  |
| 7:6.180 | הפיטן |  |
| 7:6.181 | מתקן |  |
| 7:6.182 | הסליחות |  |
| 7:6.183 | דאשמורות |  |
| 7:6.184 | שאמר |  |
| 7:6.185 | חוסה |  |
| 7:6.186 | על |  |
| 7:6.187 | עמלך |  |
| 7:6.188 | כמו |  |
| 7:6.189 | פעלך | ) |
| 7:6.190 | ולמי |  |
| 7:6.191 | אני |  |
| 7:6.192 | עמל |  |
| 7:6.193 | ומחסר |  |
| 7:6.194 | את |  |
| 7:6.195 | נפשי |  |
| 7:6.196 | מטובה | , |
| 7:6.197 | לא |  |
| 7:6.198 | להדבק |  |
| 7:6.199 | בו |  |
| 7:6.200 | וללכת |  |
| 7:6.201 | בדרכיו | ? |
| 7:6.202 | וכל |  |
| 7:6.203 | מי |  |
| 7:6.204 | שאינו |  |
| 7:6.205 | עושה |  |
| 7:6.206 | כן | . |
| 7:6.207 | גם |  |
| 7:6.208 | זה |  |
| 7:6.209 | הבל |  |
| 7:6.210 | וענין |  |
| 7:6.211 | רע |  |
| 7:6.212 | הוא |  |
| 7:6.213 | עכ"ל |  |
| 7:6.214 | לעניינינו | . |


## Milchamot ha-Shem 7:7

Hebrew (Wikisource): ומדבריהם למדנו דפשיטא להו לרבנן כי יחוד השם, ית' בשמע ישראל הוא מבלי לכלול עמו אבא ואימא וזעיר ונוקביה, ודלא כמו שאמר בזהר ואתחנן המובא לעיל סי' מ"ב ועיין שם בשם ר' חטר שהמכוין כן דעת חיצונית היא ולא כאלה חלק יעקוב, וכשתתבונן בדברי הקבלה החדשה תמצא שלא זו בלבד שהם מאמינים במציאות אלהים רבים קדושים כי אף גם זאת שהם מאמינים במציאות אלהים אחרים טמאים בעלי עשר ספירות, ונקראים ספירות דקליפה. ופרצופים רבים להם כמו בקדושה כדכתבינן לעיל. ונקראים אצלם סטרא אחרא מסאבא כמבואר כל זה בספריהם והם הפועלים את הרע בעולם. ולפעמים שולטים גם כן על האלהים הקדושים אשר יעבדו אותם המקובלים. ומשתעבדים בהם בחזקת היד כמו שכתב בעל משנת חסידים מסכת ליל פסח פרק ז' משנה ג' וז"ל ואלו לא גאל הקב"ה (אריך אנפין) את אבותינו (אבא ואימא) ממצרים (אדם בליעל) האחוז בגרון דאריך אנפין ונוטל השפע ומונעו מרדת לזעיר אנפין והיה זעיר אנפין נשאר בגלות וצער גדול כולד המעובר במעי אמו. עיין שם גם במשנה ד' שם, והנה כבר הבאתי ס' י"ג וי"ד מקצת מדברי הר' הגדול מהרש"ל בתשובותיו ובס' ים של שלמה הרבה מהמדקדקים דקדוקי עניות להניח תפילין של יד מיושב. דע אהובי כי חדשים מקרוב באו ורוצים להיות מכת המקובלים וממדרשי הנעלמים ומחלושי הראות. לא יביטו באור הזהר ולא ידעו מצאו ומובאו וכוונתו. אלא שכך מצאו בספרי הרשב"י ודע שכל רבותי הקדושים ששמשו גאוני עולם לא נהגו כך אלא כדברי התלמוד והפוסקים. ועתה פקח עיניך וראה אם יש בזה ממש אי אפשר שלא הזכירו שום מחבר. או גם לא אשתמיט רמז בתלמוד בבלי או ירושלמי ספרא וספרי ושאלתות ופסיקתא. ומי שלא ידע להשיג סודה על נכון יבוא לקצץ בנטיעות. לכך אהובי אל תלך בדרך אתם. ואין לך עסק בנסתרות המתיהרי' בחידושים. כאלו הם יודעים ומבינים רזי תורה וספוניה. והלואי שידעו הגלוים. אבל דע שכל המשנה ידו על התחתונה. וידינו על העליונה עכ"ל. עוד כ' רביד הזהב בשמו וז"ל והמורה היפך התלמוד חייב מיתה. וכמו שכ' בס' צח ואדום בשם רש"ל בתשו' החדשות וז"ל שכל המחמיר שלא לאכול בשר אחד גבינה בלא גריעותא דאית ליה בשינוהי. כי המחמיר חייב מיתה וכו' ובס' ים שלמה כ' דמאחר שנחתם התלמוד אין להחמיר נגד התלמוד והוי כמו מינות ע"ש.

| id | word | punctuation after |
|---|---|---|
| 7:7.1 | ומדבריהם |  |
| 7:7.2 | למדנו |  |
| 7:7.3 | דפשיטא |  |
| 7:7.4 | להו |  |
| 7:7.5 | לרבנן |  |
| 7:7.6 | כי |  |
| 7:7.7 | יחוד |  |
| 7:7.8 | השם | , |
| 7:7.9 | ית | ' |
| 7:7.10 | בשמע |  |
| 7:7.11 | ישראל |  |
| 7:7.12 | הוא |  |
| 7:7.13 | מבלי |  |
| 7:7.14 | לכלול |  |
| 7:7.15 | עמו |  |
| 7:7.16 | אבא |  |
| 7:7.17 | ואימא |  |
| 7:7.18 | וזעיר |  |
| 7:7.19 | ונוקביה | , |
| 7:7.20 | ודלא |  |
| 7:7.21 | כמו |  |
| 7:7.22 | שאמר |  |
| 7:7.23 | בזהר |  |
| 7:7.24 | ואתחנן |  |
| 7:7.25 | המובא |  |
| 7:7.26 | לעיל |  |
| 7:7.27 | סי | ' |
| 7:7.28 | מ"ב |  |
| 7:7.29 | ועיין |  |
| 7:7.30 | שם |  |
| 7:7.31 | בשם |  |
| 7:7.32 | ר | ' |
| 7:7.33 | חטר |  |
| 7:7.34 | שהמכוין |  |
| 7:7.35 | כן |  |
| 7:7.36 | דעת |  |
| 7:7.37 | חיצונית |  |
| 7:7.38 | היא |  |
| 7:7.39 | ולא |  |
| 7:7.40 | כאלה |  |
| 7:7.41 | חלק |  |
| 7:7.42 | יעקוב | , |
| 7:7.43 | וכשתתבונן |  |
| 7:7.44 | בדברי |  |
| 7:7.45 | הקבלה |  |
| 7:7.46 | החדשה |  |
| 7:7.47 | תמצא |  |
| 7:7.48 | שלא |  |
| 7:7.49 | זו |  |
| 7:7.50 | בלבד |  |
| 7:7.51 | שהם |  |
| 7:7.52 | מאמינים |  |
| 7:7.53 | במציאות |  |
| 7:7.54 | אלהים |  |
| 7:7.55 | רבים |  |
| 7:7.56 | קדושים |  |
| 7:7.57 | כי |  |
| 7:7.58 | אף |  |
| 7:7.59 | גם |  |
| 7:7.60 | זאת |  |
| 7:7.61 | שהם |  |
| 7:7.62 | מאמינים |  |
| 7:7.63 | במציאות |  |
| 7:7.64 | אלהים |  |
| 7:7.65 | אחרים |  |
| 7:7.66 | טמאים |  |
| 7:7.67 | בעלי |  |
| 7:7.68 | עשר |  |
| 7:7.69 | ספירות | , |
| 7:7.70 | ונקראים |  |
| 7:7.71 | ספירות |  |
| 7:7.72 | דקליפה | . |
| 7:7.73 | ופרצופים |  |
| 7:7.74 | רבים |  |
| 7:7.75 | להם |  |
| 7:7.76 | כמו |  |
| 7:7.77 | בקדושה |  |
| 7:7.78 | כדכתבינן |  |
| 7:7.79 | לעיל | . |
| 7:7.80 | ונקראים |  |
| 7:7.81 | אצלם |  |
| 7:7.82 | סטרא |  |
| 7:7.83 | אחרא |  |
| 7:7.84 | מסאבא |  |
| 7:7.85 | כמבואר |  |
| 7:7.86 | כל |  |
| 7:7.87 | זה |  |
| 7:7.88 | בספריהם |  |
| 7:7.89 | והם |  |
| 7:7.90 | הפועלים |  |
| 7:7.91 | את |  |
| 7:7.92 | הרע |  |
| 7:7.93 | בעולם | . |
| 7:7.94 | ולפעמים |  |
| 7:7.95 | שולטים |  |
| 7:7.96 | גם |  |
| 7:7.97 | כן |  |
| 7:7.98 | על |  |
| 7:7.99 | האלהים |  |
| 7:7.100 | הקדושים |  |
| 7:7.101 | אשר |  |
| 7:7.102 | יעבדו |  |
| 7:7.103 | אותם |  |
| 7:7.104 | המקובלים | . |
| 7:7.105 | ומשתעבדים |  |
| 7:7.106 | בהם |  |
| 7:7.107 | בחזקת |  |
| 7:7.108 | היד |  |
| 7:7.109 | כמו |  |
| 7:7.110 | שכתב |  |
| 7:7.111 | בעל |  |
| 7:7.112 | משנת |  |
| 7:7.113 | חסידים |  |
| 7:7.114 | מסכת |  |
| 7:7.115 | ליל |  |
| 7:7.116 | פסח |  |
| 7:7.117 | פרק |  |
| 7:7.118 | ז | ' |
| 7:7.119 | משנה |  |
| 7:7.120 | ג | ' |
| 7:7.121 | וז"ל |  |
| 7:7.122 | ואלו |  |
| 7:7.123 | לא |  |
| 7:7.124 | גאל |  |
| 7:7.125 | הקב"ה |  |
| 7:7.126 | אריך |  |
| 7:7.127 | אנפין | ) |
| 7:7.128 | את |  |
| 7:7.129 | אבותינו |  |
| 7:7.130 | אבא |  |
| 7:7.131 | ואימא | ) |
| 7:7.132 | ממצרים |  |
| 7:7.133 | אדם |  |
| 7:7.134 | בליעל | ) |
| 7:7.135 | האחוז |  |
| 7:7.136 | בגרון |  |
| 7:7.137 | דאריך |  |
| 7:7.138 | אנפין |  |
| 7:7.139 | ונוטל |  |
| 7:7.140 | השפע |  |
| 7:7.141 | ומונעו |  |
| 7:7.142 | מרדת |  |
| 7:7.143 | לזעיר |  |
| 7:7.144 | אנפין |  |
| 7:7.145 | והיה |  |
| 7:7.146 | זעיר |  |
| 7:7.147 | אנפין |  |
| 7:7.148 | נשאר |  |
| 7:7.149 | בגלות |  |
| 7:7.150 | וצער |  |
| 7:7.151 | גדול |  |
| 7:7.152 | כולד |  |
| 7:7.153 | המעובר |  |
| 7:7.154 | במעי |  |
| 7:7.155 | אמו | . |
| 7:7.156 | עיין |  |
| 7:7.157 | שם |  |
| 7:7.158 | גם |  |
| 7:7.159 | במשנה |  |
| 7:7.160 | ד | ' |
| 7:7.161 | שם | , |
| 7:7.162 | והנה |  |
| 7:7.163 | כבר |  |
| 7:7.164 | הבאתי |  |
| 7:7.165 | ס | ' |
| 7:7.166 | י"ג |  |
| 7:7.167 | וי"ד |  |
| 7:7.168 | מקצת |  |
| 7:7.169 | מדברי |  |
| 7:7.170 | הר | ' |
| 7:7.171 | הגדול |  |
| 7:7.172 | מהרש"ל |  |
| 7:7.173 | בתשובותיו |  |
| 7:7.174 | ובס | ' |
| 7:7.175 | ים |  |
| 7:7.176 | של |  |
| 7:7.177 | שלמה |  |
| 7:7.178 | הרבה |  |
| 7:7.179 | מהמדקדקים |  |
| 7:7.180 | דקדוקי |  |
| 7:7.181 | עניות |  |
| 7:7.182 | להניח |  |
| 7:7.183 | תפילין |  |
| 7:7.184 | של |  |
| 7:7.185 | יד |  |
| 7:7.186 | מיושב | . |
| 7:7.187 | דע |  |
| 7:7.188 | אהובי |  |
| 7:7.189 | כי |  |
| 7:7.190 | חדשים |  |
| 7:7.191 | מקרוב |  |
| 7:7.192 | באו |  |
| 7:7.193 | ורוצים |  |
| 7:7.194 | להיות |  |
| 7:7.195 | מכת |  |
| 7:7.196 | המקובלים |  |
| 7:7.197 | וממדרשי |  |
| 7:7.198 | הנעלמים |  |
| 7:7.199 | ומחלושי |  |
| 7:7.200 | הראות | . |
| 7:7.201 | לא |  |
| 7:7.202 | יביטו |  |
| 7:7.203 | באור |  |
| 7:7.204 | הזהר |  |
| 7:7.205 | ולא |  |
| 7:7.206 | ידעו |  |
| 7:7.207 | מצאו |  |
| 7:7.208 | ומובאו |  |
| 7:7.209 | וכוונתו | . |
| 7:7.210 | אלא |  |
| 7:7.211 | שכך |  |
| 7:7.212 | מצאו |  |
| 7:7.213 | בספרי |  |
| 7:7.214 | הרשב"י |  |
| 7:7.215 | ודע |  |
| 7:7.216 | שכל |  |
| 7:7.217 | רבותי |  |
| 7:7.218 | הקדושים |  |
| 7:7.219 | ששמשו |  |
| 7:7.220 | גאוני |  |
| 7:7.221 | עולם |  |
| 7:7.222 | לא |  |
| 7:7.223 | נהגו |  |
| 7:7.224 | כך |  |
| 7:7.225 | אלא |  |
| 7:7.226 | כדברי |  |
| 7:7.227 | התלמוד |  |
| 7:7.228 | והפוסקים | . |
| 7:7.229 | ועתה |  |
| 7:7.230 | פקח |  |
| 7:7.231 | עיניך |  |
| 7:7.232 | וראה |  |
| 7:7.233 | אם |  |
| 7:7.234 | יש |  |
| 7:7.235 | בזה |  |
| 7:7.236 | ממש |  |
| 7:7.237 | אי |  |
| 7:7.238 | אפשר |  |
| 7:7.239 | שלא |  |
| 7:7.240 | הזכירו |  |
| 7:7.241 | שום |  |
| 7:7.242 | מחבר | . |
| 7:7.243 | או |  |
| 7:7.244 | גם |  |
| 7:7.245 | לא |  |
| 7:7.246 | אשתמיט |  |
| 7:7.247 | רמז |  |
| 7:7.248 | בתלמוד |  |
| 7:7.249 | בבלי |  |
| 7:7.250 | או |  |
| 7:7.251 | ירושלמי |  |
| 7:7.252 | ספרא |  |
| 7:7.253 | וספרי |  |
| 7:7.254 | ושאלתות |  |
| 7:7.255 | ופסיקתא | . |
| 7:7.256 | ומי |  |
| 7:7.257 | שלא |  |
| 7:7.258 | ידע |  |
| 7:7.259 | להשיג |  |
| 7:7.260 | סודה |  |
| 7:7.261 | על |  |
| 7:7.262 | נכון |  |
| 7:7.263 | יבוא |  |
| 7:7.264 | לקצץ |  |
| 7:7.265 | בנטיעות | . |
| 7:7.266 | לכך |  |
| 7:7.267 | אהובי |  |
| 7:7.268 | אל |  |
| 7:7.269 | תלך |  |
| 7:7.270 | בדרך |  |
| 7:7.271 | אתם | . |
| 7:7.272 | ואין |  |
| 7:7.273 | לך |  |
| 7:7.274 | עסק |  |
| 7:7.275 | בנסתרות |  |
| 7:7.276 | המתיהרי | ' |
| 7:7.277 | בחידושים | . |
| 7:7.278 | כאלו |  |
| 7:7.279 | הם |  |
| 7:7.280 | יודעים |  |
| 7:7.281 | ומבינים |  |
| 7:7.282 | רזי |  |
| 7:7.283 | תורה |  |
| 7:7.284 | וספוניה | . |
| 7:7.285 | והלואי |  |
| 7:7.286 | שידעו |  |
| 7:7.287 | הגלוים | . |
| 7:7.288 | אבל |  |
| 7:7.289 | דע |  |
| 7:7.290 | שכל |  |
| 7:7.291 | המשנה |  |
| 7:7.292 | ידו |  |
| 7:7.293 | על |  |
| 7:7.294 | התחתונה | . |
| 7:7.295 | וידינו |  |
| 7:7.296 | על |  |
| 7:7.297 | העליונה |  |
| 7:7.298 | עכ"ל | . |
| 7:7.299 | עוד |  |
| 7:7.300 | כ | ' |
| 7:7.301 | רביד |  |
| 7:7.302 | הזהב |  |
| 7:7.303 | בשמו |  |
| 7:7.304 | וז"ל |  |
| 7:7.305 | והמורה |  |
| 7:7.306 | היפך |  |
| 7:7.307 | התלמוד |  |
| 7:7.308 | חייב |  |
| 7:7.309 | מיתה | . |
| 7:7.310 | וכמו |  |
| 7:7.311 | שכ | ' |
| 7:7.312 | בס | ' |
| 7:7.313 | צח |  |
| 7:7.314 | ואדום |  |
| 7:7.315 | בשם |  |
| 7:7.316 | רש"ל |  |
| 7:7.317 | בתשו | ' |
| 7:7.318 | החדשות |  |
| 7:7.319 | וז"ל |  |
| 7:7.320 | שכל |  |
| 7:7.321 | המחמיר |  |
| 7:7.322 | שלא |  |
| 7:7.323 | לאכול |  |
| 7:7.324 | בשר |  |
| 7:7.325 | אחד |  |
| 7:7.326 | גבינה |  |
| 7:7.327 | בלא |  |
| 7:7.328 | גריעותא |  |
| 7:7.329 | דאית |  |
| 7:7.330 | ליה |  |
| 7:7.331 | בשינוהי | . |
| 7:7.332 | כי |  |
| 7:7.333 | המחמיר |  |
| 7:7.334 | חייב |  |
| 7:7.335 | מיתה |  |
| 7:7.336 | וכו | ' |
| 7:7.337 | ובס | ' |
| 7:7.338 | ים |  |
| 7:7.339 | שלמה |  |
| 7:7.340 | כ | ' |
| 7:7.341 | דמאחר |  |
| 7:7.342 | שנחתם |  |
| 7:7.343 | התלמוד |  |
| 7:7.344 | אין |  |
| 7:7.345 | להחמיר |  |
| 7:7.346 | נגד |  |
| 7:7.347 | התלמוד |  |
| 7:7.348 | והוי |  |
| 7:7.349 | כמו |  |
| 7:7.350 | מינות |  |
| 7:7.351 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

