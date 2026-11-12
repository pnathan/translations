# Translation pack: milhamot-hashem perek-7-08 (Milchamot ha-Shem 7:11–7:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-08`.

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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
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

- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- ימי **הבינים** → "the Middle Ages".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אין לנו **להחמיר** → "we may not be stringent".
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
- מי **אתי** (ketiv): translated, the qere noted.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
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


## Milchamot ha-Shem 7:11

Hebrew (Wikisource): מעתה ראה והתבונן ושים אל לבך בשינוי המנהג הקדום אפי' להחמיר כהני גווני שהבאתי, וכן שלא לומר ומשרתיו שואלים וכו' כי אם במוסף שבת ולא בשאר ימים מטעם האמור במשנת חסידים כולם מביאים לידי מינות כמו שביארתי ושפתי הרב הגדול מהרש"ל ור' תם בן יחיא ברור מללו, שהמשנה ועושה על פי הקבלה החדשה וטעמיה הוא מביא לידי מינות.

| id | word | punctuation after |
|---|---|---|
| 7:11.1 | מעתה |  |
| 7:11.2 | ראה |  |
| 7:11.3 | והתבונן |  |
| 7:11.4 | ושים |  |
| 7:11.5 | אל |  |
| 7:11.6 | לבך |  |
| 7:11.7 | בשינוי |  |
| 7:11.8 | המנהג |  |
| 7:11.9 | הקדום |  |
| 7:11.10 | אפי | ' |
| 7:11.11 | להחמיר |  |
| 7:11.12 | כהני |  |
| 7:11.13 | גווני |  |
| 7:11.14 | שהבאתי | , |
| 7:11.15 | וכן |  |
| 7:11.16 | שלא |  |
| 7:11.17 | לומר |  |
| 7:11.18 | ומשרתיו |  |
| 7:11.19 | שואלים |  |
| 7:11.20 | וכו | ' |
| 7:11.21 | כי |  |
| 7:11.22 | אם |  |
| 7:11.23 | במוסף |  |
| 7:11.24 | שבת |  |
| 7:11.25 | ולא |  |
| 7:11.26 | בשאר |  |
| 7:11.27 | ימים |  |
| 7:11.28 | מטעם |  |
| 7:11.29 | האמור |  |
| 7:11.30 | במשנת |  |
| 7:11.31 | חסידים |  |
| 7:11.32 | כולם |  |
| 7:11.33 | מביאים |  |
| 7:11.34 | לידי |  |
| 7:11.35 | מינות |  |
| 7:11.36 | כמו |  |
| 7:11.37 | שביארתי |  |
| 7:11.38 | ושפתי |  |
| 7:11.39 | הרב |  |
| 7:11.40 | הגדול |  |
| 7:11.41 | מהרש"ל |  |
| 7:11.42 | ור | ' |
| 7:11.43 | תם |  |
| 7:11.44 | בן |  |
| 7:11.45 | יחיא |  |
| 7:11.46 | ברור |  |
| 7:11.47 | מללו | , |
| 7:11.48 | שהמשנה |  |
| 7:11.49 | ועושה |  |
| 7:11.50 | על |  |
| 7:11.51 | פי |  |
| 7:11.52 | הקבלה |  |
| 7:11.53 | החדשה |  |
| 7:11.54 | וטעמיה |  |
| 7:11.55 | הוא |  |
| 7:11.56 | מביא |  |
| 7:11.57 | לידי |  |
| 7:11.58 | מינות | . |


## Milchamot ha-Shem 7:12

Hebrew (Wikisource): ולא יפה עשו דורות שעמדו לפנינו בימי הבינים אשר עזבו למוד המשנה והתלמוד ופסקי הלכות וקבעו רוב עסקם בספר החיצון המסית ס' הזהר ולא בעיון, רק כעוף המצפצף, ושינו עלינו כמה דברים בסידורי התפלות ועשיית המצות חוקי האלהים ותורותיו על פי הקבלה החדשה, ועזבו מנהגי אבותם המיוסדים על פי רבותינו בעלי התלמוד והרמב"ם והרס"ג ור"י הלוי וכיוצא, וגרעו מהן והוסיפו עליהן ככל העולה על רוחם, ובעיון ראשון ומושכל ראשון לא ירגיש האדם כי דעת חיצונית, היוצאת מגדר תורה היא, עד אשר יעמיק חקר באותו שינוי ובטעמו הנעשה שלא על פי רבותינו ז"ל, ואז ימצא באותו שינוי ואותה חומרא לתא דאיסורא או חשש ברכה לבטלה וכיוצא בזה! מה שאין כן מנהגי קדמונינו כולם הם מנוקים ומנופים מכל סיג, ומיוסדים על אדני פז על פי התלמוד והפוסקים הראשונים כהרי"פ והרמב"ם וכיוצא. הלא תראה מאי דתנן בפרק הקורא עומד, האומר איני עובר לפני התיבה בצבועין. אף בלבנים לא יעבור. בסנדל איני עובר, אף יחף לא יעבור וכו' אלמא כל דקפיד במלתא שאין דרך ישראל להקפיד חישנין שמא מחשבת שטות עלתה בו והוי מינות כמו שכ' התוספת יו"ט בשם ר"ן. ולכן כל הרוצה להדבק בשם ית' ובתורתו ירחיק עצמו מהמנהגים שנתחדשו בימי הביניים. וידבק במנהגי קדמונינו, על ידי שיתן לב ויעיין יפה בתלמוד ופוסקים הראשונים. וסדורי תפלה העתיקים ובשאר ספרי קדמונינו, ויקבע לימודו במשנה עם פי' הרמב"ם ור"ע וכיוצא ובתלמוד ופוסקים המובהקים וילך לאורם, כמאמר רבותינו ז"ל האי מאן דבעי למהוי חסידא ליקיים מילי דאבות וברכות.

| id | word | punctuation after |
|---|---|---|
| 7:12.1 | ולא |  |
| 7:12.2 | יפה |  |
| 7:12.3 | עשו |  |
| 7:12.4 | דורות |  |
| 7:12.5 | שעמדו |  |
| 7:12.6 | לפנינו |  |
| 7:12.7 | בימי |  |
| 7:12.8 | הבינים |  |
| 7:12.9 | אשר |  |
| 7:12.10 | עזבו |  |
| 7:12.11 | למוד |  |
| 7:12.12 | המשנה |  |
| 7:12.13 | והתלמוד |  |
| 7:12.14 | ופסקי |  |
| 7:12.15 | הלכות |  |
| 7:12.16 | וקבעו |  |
| 7:12.17 | רוב |  |
| 7:12.18 | עסקם |  |
| 7:12.19 | בספר |  |
| 7:12.20 | החיצון |  |
| 7:12.21 | המסית |  |
| 7:12.22 | ס | ' |
| 7:12.23 | הזהר |  |
| 7:12.24 | ולא |  |
| 7:12.25 | בעיון | , |
| 7:12.26 | רק |  |
| 7:12.27 | כעוף |  |
| 7:12.28 | המצפצף | , |
| 7:12.29 | ושינו |  |
| 7:12.30 | עלינו |  |
| 7:12.31 | כמה |  |
| 7:12.32 | דברים |  |
| 7:12.33 | בסידורי |  |
| 7:12.34 | התפלות |  |
| 7:12.35 | ועשיית |  |
| 7:12.36 | המצות |  |
| 7:12.37 | חוקי |  |
| 7:12.38 | האלהים |  |
| 7:12.39 | ותורותיו |  |
| 7:12.40 | על |  |
| 7:12.41 | פי |  |
| 7:12.42 | הקבלה |  |
| 7:12.43 | החדשה | , |
| 7:12.44 | ועזבו |  |
| 7:12.45 | מנהגי |  |
| 7:12.46 | אבותם |  |
| 7:12.47 | המיוסדים |  |
| 7:12.48 | על |  |
| 7:12.49 | פי |  |
| 7:12.50 | רבותינו |  |
| 7:12.51 | בעלי |  |
| 7:12.52 | התלמוד |  |
| 7:12.53 | והרמב"ם |  |
| 7:12.54 | והרס"ג |  |
| 7:12.55 | ור"י |  |
| 7:12.56 | הלוי |  |
| 7:12.57 | וכיוצא | , |
| 7:12.58 | וגרעו |  |
| 7:12.59 | מהן |  |
| 7:12.60 | והוסיפו |  |
| 7:12.61 | עליהן |  |
| 7:12.62 | ככל |  |
| 7:12.63 | העולה |  |
| 7:12.64 | על |  |
| 7:12.65 | רוחם | , |
| 7:12.66 | ובעיון |  |
| 7:12.67 | ראשון |  |
| 7:12.68 | ומושכל |  |
| 7:12.69 | ראשון |  |
| 7:12.70 | לא |  |
| 7:12.71 | ירגיש |  |
| 7:12.72 | האדם |  |
| 7:12.73 | כי |  |
| 7:12.74 | דעת |  |
| 7:12.75 | חיצונית | , |
| 7:12.76 | היוצאת |  |
| 7:12.77 | מגדר |  |
| 7:12.78 | תורה |  |
| 7:12.79 | היא | , |
| 7:12.80 | עד |  |
| 7:12.81 | אשר |  |
| 7:12.82 | יעמיק |  |
| 7:12.83 | חקר |  |
| 7:12.84 | באותו |  |
| 7:12.85 | שינוי |  |
| 7:12.86 | ובטעמו |  |
| 7:12.87 | הנעשה |  |
| 7:12.88 | שלא |  |
| 7:12.89 | על |  |
| 7:12.90 | פי |  |
| 7:12.91 | רבותינו |  |
| 7:12.92 | ז"ל | , |
| 7:12.93 | ואז |  |
| 7:12.94 | ימצא |  |
| 7:12.95 | באותו |  |
| 7:12.96 | שינוי |  |
| 7:12.97 | ואותה |  |
| 7:12.98 | חומרא |  |
| 7:12.99 | לתא |  |
| 7:12.100 | דאיסורא |  |
| 7:12.101 | או |  |
| 7:12.102 | חשש |  |
| 7:12.103 | ברכה |  |
| 7:12.104 | לבטלה |  |
| 7:12.105 | וכיוצא |  |
| 7:12.106 | בזה | ! |
| 7:12.107 | מה |  |
| 7:12.108 | שאין |  |
| 7:12.109 | כן |  |
| 7:12.110 | מנהגי |  |
| 7:12.111 | קדמונינו |  |
| 7:12.112 | כולם |  |
| 7:12.113 | הם |  |
| 7:12.114 | מנוקים |  |
| 7:12.115 | ומנופים |  |
| 7:12.116 | מכל |  |
| 7:12.117 | סיג | , |
| 7:12.118 | ומיוסדים |  |
| 7:12.119 | על |  |
| 7:12.120 | אדני |  |
| 7:12.121 | פז |  |
| 7:12.122 | על |  |
| 7:12.123 | פי |  |
| 7:12.124 | התלמוד |  |
| 7:12.125 | והפוסקים |  |
| 7:12.126 | הראשונים |  |
| 7:12.127 | כהרי"פ |  |
| 7:12.128 | והרמב"ם |  |
| 7:12.129 | וכיוצא | . |
| 7:12.130 | הלא |  |
| 7:12.131 | תראה |  |
| 7:12.132 | מאי |  |
| 7:12.133 | דתנן |  |
| 7:12.134 | בפרק |  |
| 7:12.135 | הקורא |  |
| 7:12.136 | עומד | , |
| 7:12.137 | האומר |  |
| 7:12.138 | איני |  |
| 7:12.139 | עובר |  |
| 7:12.140 | לפני |  |
| 7:12.141 | התיבה |  |
| 7:12.142 | בצבועין | . |
| 7:12.143 | אף |  |
| 7:12.144 | בלבנים |  |
| 7:12.145 | לא |  |
| 7:12.146 | יעבור | . |
| 7:12.147 | בסנדל |  |
| 7:12.148 | איני |  |
| 7:12.149 | עובר | , |
| 7:12.150 | אף |  |
| 7:12.151 | יחף |  |
| 7:12.152 | לא |  |
| 7:12.153 | יעבור |  |
| 7:12.154 | וכו | ' |
| 7:12.155 | אלמא |  |
| 7:12.156 | כל |  |
| 7:12.157 | דקפיד |  |
| 7:12.158 | במלתא |  |
| 7:12.159 | שאין |  |
| 7:12.160 | דרך |  |
| 7:12.161 | ישראל |  |
| 7:12.162 | להקפיד |  |
| 7:12.163 | חישנין |  |
| 7:12.164 | שמא |  |
| 7:12.165 | מחשבת |  |
| 7:12.166 | שטות |  |
| 7:12.167 | עלתה |  |
| 7:12.168 | בו |  |
| 7:12.169 | והוי |  |
| 7:12.170 | מינות |  |
| 7:12.171 | כמו |  |
| 7:12.172 | שכ | ' |
| 7:12.173 | התוספת |  |
| 7:12.174 | יו"ט |  |
| 7:12.175 | בשם |  |
| 7:12.176 | ר"ן | . |
| 7:12.177 | ולכן |  |
| 7:12.178 | כל |  |
| 7:12.179 | הרוצה |  |
| 7:12.180 | להדבק |  |
| 7:12.181 | בשם |  |
| 7:12.182 | ית | ' |
| 7:12.183 | ובתורתו |  |
| 7:12.184 | ירחיק |  |
| 7:12.185 | עצמו |  |
| 7:12.186 | מהמנהגים |  |
| 7:12.187 | שנתחדשו |  |
| 7:12.188 | בימי |  |
| 7:12.189 | הביניים | . |
| 7:12.190 | וידבק |  |
| 7:12.191 | במנהגי |  |
| 7:12.192 | קדמונינו | , |
| 7:12.193 | על |  |
| 7:12.194 | ידי |  |
| 7:12.195 | שיתן |  |
| 7:12.196 | לב |  |
| 7:12.197 | ויעיין |  |
| 7:12.198 | יפה |  |
| 7:12.199 | בתלמוד |  |
| 7:12.200 | ופוסקים |  |
| 7:12.201 | הראשונים | . |
| 7:12.202 | וסדורי |  |
| 7:12.203 | תפלה |  |
| 7:12.204 | העתיקים |  |
| 7:12.205 | ובשאר |  |
| 7:12.206 | ספרי |  |
| 7:12.207 | קדמונינו | , |
| 7:12.208 | ויקבע |  |
| 7:12.209 | לימודו |  |
| 7:12.210 | במשנה |  |
| 7:12.211 | עם |  |
| 7:12.212 | פי | ' |
| 7:12.213 | הרמב"ם |  |
| 7:12.214 | ור"ע |  |
| 7:12.215 | וכיוצא |  |
| 7:12.216 | ובתלמוד |  |
| 7:12.217 | ופוסקים |  |
| 7:12.218 | המובהקים |  |
| 7:12.219 | וילך |  |
| 7:12.220 | לאורם | , |
| 7:12.221 | כמאמר |  |
| 7:12.222 | רבותינו |  |
| 7:12.223 | ז"ל |  |
| 7:12.224 | האי |  |
| 7:12.225 | מאן |  |
| 7:12.226 | דבעי |  |
| 7:12.227 | למהוי |  |
| 7:12.228 | חסידא |  |
| 7:12.229 | ליקיים |  |
| 7:12.230 | מילי |  |
| 7:12.231 | דאבות |  |
| 7:12.232 | וברכות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

