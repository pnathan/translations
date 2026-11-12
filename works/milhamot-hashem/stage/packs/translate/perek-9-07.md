# Translation pack: milhamot-hashem perek-9-07 (Milchamot ha-Shem 9:9–9:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-07`.

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
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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

- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- הסבה **הראשונה** → "the First Cause".
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
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- לית מחשבה **תפיסה** ביה → "no thought grasps Him".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- אשם **תלוי** → "suspended guilt offering".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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


## Milchamot ha-Shem 9:9

Hebrew (Wikisource): קכו) הוא שהנביא ישעיה מוכיח את ישראל הוי גוי חוטא וכו' עזבו את ה' לא על ציצית ותפילין וסוכה ולולב הוכיחם וקרא להם גוי חוטא עם כבד עון כי בודאי היו אוכלים מצה ומרור ועושים שאר מצות אך היו עושים אותם אותם לשם הבעלים והעשתרות אשר בדו להם נביאי השקר תכניתם ולצלמם ודמותם ואמרו להם שנשמתם הוא חלק אלוה ממעל יעשו כ"כ צורות גשמיות עץ ואבן וכסף וזהב, כדי לעבוד לאלוה המוחש במעשה, והלמשיך על ידו את השפע מהצורות, אשר בשמים ממעל שבדו להם נביאיהם וצוו אותם לעבדם! אף כן עשו בעלי הקבלה החדשה שקבלוה מיד המלך העכו"ם, ויאמרו כי שקר נחלו אבותינו לעבוד ולהתפלל ולקרא לה' אלהים אמת שהוא הסבה הראשונה, בעבור דלית מחשבה תפיסה ביה (שאינו יכלל לציירו בדמיונו)! היפך דברי הנביא שאמר ה' עזי ומעזי ומנוסי, אליך גוים ידרושו מאפסי ארץ ויאמרו אך שקר נחלו אבותינו, ויודו לדת האמת, לא שישראל יורו להם לעבוד את האלילים! כאשר עשו המקובלים החדשים על פי הזהר הנמסר להם מידי המלך העכו"ם שאמר שאין ראוי לעבוד אלא לסבות הנאצלים ונמשכים ממנו לפי דעתם, כאילן המשלח פארותיו הילך והילך, שהם מושגים להם על פי רעיונותיהם, ותלו כל מצותיה של תורה כאלו הגופים הנבראים וקראום גופא דמלכא יש מצות תלוים בידא דמלכא, וברגלוי, ושאר אברים שבגוף, ויש מצות אשר לדעתם הם ביסוד. ערית מלכא דלא אריך למחזי בטעם כסוי הדם של חיה ועוף כמו שכתב הרב מקור חיים בטעם כסוי דמם! ועל כל אלה הענינים יבכו הכהנים משרתי ה' בין האולם והמזבח ויאמרו חוסה ה' על עמיך. והט לבם לדעת ולהשכיל אמתתך. ולדעת את דרכיך אשר צוית את משה עבדיך ללמד לישראל עלילותיך ולעבדך בלבב שלם. ולא לעבוד לגופים נבראים שמכנים אותם בשם אורות אלהיים בעבור דבר אלהי המעורב ומשותף בהם שהיא נשמתם כמו שחשבו המקובלים החדשים.

| id | word | punctuation after |
|---|---|---|
| 9:9.1 | קכו | ) |
| 9:9.2 | הוא |  |
| 9:9.3 | שהנביא |  |
| 9:9.4 | ישעיה |  |
| 9:9.5 | מוכיח |  |
| 9:9.6 | את |  |
| 9:9.7 | ישראל |  |
| 9:9.8 | הוי |  |
| 9:9.9 | גוי |  |
| 9:9.10 | חוטא |  |
| 9:9.11 | וכו | ' |
| 9:9.12 | עזבו |  |
| 9:9.13 | את |  |
| 9:9.14 | ה | ' |
| 9:9.15 | לא |  |
| 9:9.16 | על |  |
| 9:9.17 | ציצית |  |
| 9:9.18 | ותפילין |  |
| 9:9.19 | וסוכה |  |
| 9:9.20 | ולולב |  |
| 9:9.21 | הוכיחם |  |
| 9:9.22 | וקרא |  |
| 9:9.23 | להם |  |
| 9:9.24 | גוי |  |
| 9:9.25 | חוטא |  |
| 9:9.26 | עם |  |
| 9:9.27 | כבד |  |
| 9:9.28 | עון |  |
| 9:9.29 | כי |  |
| 9:9.30 | בודאי |  |
| 9:9.31 | היו |  |
| 9:9.32 | אוכלים |  |
| 9:9.33 | מצה |  |
| 9:9.34 | ומרור |  |
| 9:9.35 | ועושים |  |
| 9:9.36 | שאר |  |
| 9:9.37 | מצות |  |
| 9:9.38 | אך |  |
| 9:9.39 | היו |  |
| 9:9.40 | עושים |  |
| 9:9.41 | אותם |  |
| 9:9.42 | אותם |  |
| 9:9.43 | לשם |  |
| 9:9.44 | הבעלים |  |
| 9:9.45 | והעשתרות |  |
| 9:9.46 | אשר |  |
| 9:9.47 | בדו |  |
| 9:9.48 | להם |  |
| 9:9.49 | נביאי |  |
| 9:9.50 | השקר |  |
| 9:9.51 | תכניתם |  |
| 9:9.52 | ולצלמם |  |
| 9:9.53 | ודמותם |  |
| 9:9.54 | ואמרו |  |
| 9:9.55 | להם |  |
| 9:9.56 | שנשמתם |  |
| 9:9.57 | הוא |  |
| 9:9.58 | חלק |  |
| 9:9.59 | אלוה |  |
| 9:9.60 | ממעל |  |
| 9:9.61 | יעשו |  |
| 9:9.62 | כ"כ |  |
| 9:9.63 | צורות |  |
| 9:9.64 | גשמיות |  |
| 9:9.65 | עץ |  |
| 9:9.66 | ואבן |  |
| 9:9.67 | וכסף |  |
| 9:9.68 | וזהב | , |
| 9:9.69 | כדי |  |
| 9:9.70 | לעבוד |  |
| 9:9.71 | לאלוה |  |
| 9:9.72 | המוחש |  |
| 9:9.73 | במעשה | , |
| 9:9.74 | והלמשיך |  |
| 9:9.75 | על |  |
| 9:9.76 | ידו |  |
| 9:9.77 | את |  |
| 9:9.78 | השפע |  |
| 9:9.79 | מהצורות | , |
| 9:9.80 | אשר |  |
| 9:9.81 | בשמים |  |
| 9:9.82 | ממעל |  |
| 9:9.83 | שבדו |  |
| 9:9.84 | להם |  |
| 9:9.85 | נביאיהם |  |
| 9:9.86 | וצוו |  |
| 9:9.87 | אותם |  |
| 9:9.88 | לעבדם | ! |
| 9:9.89 | אף |  |
| 9:9.90 | כן |  |
| 9:9.91 | עשו |  |
| 9:9.92 | בעלי |  |
| 9:9.93 | הקבלה |  |
| 9:9.94 | החדשה |  |
| 9:9.95 | שקבלוה |  |
| 9:9.96 | מיד |  |
| 9:9.97 | המלך |  |
| 9:9.98 | העכו"ם | , |
| 9:9.99 | ויאמרו |  |
| 9:9.100 | כי |  |
| 9:9.101 | שקר |  |
| 9:9.102 | נחלו |  |
| 9:9.103 | אבותינו |  |
| 9:9.104 | לעבוד |  |
| 9:9.105 | ולהתפלל |  |
| 9:9.106 | ולקרא |  |
| 9:9.107 | לה | ' |
| 9:9.108 | אלהים |  |
| 9:9.109 | אמת |  |
| 9:9.110 | שהוא |  |
| 9:9.111 | הסבה |  |
| 9:9.112 | הראשונה | , |
| 9:9.113 | בעבור |  |
| 9:9.114 | דלית |  |
| 9:9.115 | מחשבה |  |
| 9:9.116 | תפיסה |  |
| 9:9.117 | ביה |  |
| 9:9.118 | שאינו |  |
| 9:9.119 | יכלל |  |
| 9:9.120 | לציירו |  |
| 9:9.121 | בדמיונו | )! |
| 9:9.122 | היפך |  |
| 9:9.123 | דברי |  |
| 9:9.124 | הנביא |  |
| 9:9.125 | שאמר |  |
| 9:9.126 | ה | ' |
| 9:9.127 | עזי |  |
| 9:9.128 | ומעזי |  |
| 9:9.129 | ומנוסי | , |
| 9:9.130 | אליך |  |
| 9:9.131 | גוים |  |
| 9:9.132 | ידרושו |  |
| 9:9.133 | מאפסי |  |
| 9:9.134 | ארץ |  |
| 9:9.135 | ויאמרו |  |
| 9:9.136 | אך |  |
| 9:9.137 | שקר |  |
| 9:9.138 | נחלו |  |
| 9:9.139 | אבותינו | , |
| 9:9.140 | ויודו |  |
| 9:9.141 | לדת |  |
| 9:9.142 | האמת | , |
| 9:9.143 | לא |  |
| 9:9.144 | שישראל |  |
| 9:9.145 | יורו |  |
| 9:9.146 | להם |  |
| 9:9.147 | לעבוד |  |
| 9:9.148 | את |  |
| 9:9.149 | האלילים | ! |
| 9:9.150 | כאשר |  |
| 9:9.151 | עשו |  |
| 9:9.152 | המקובלים |  |
| 9:9.153 | החדשים |  |
| 9:9.154 | על |  |
| 9:9.155 | פי |  |
| 9:9.156 | הזהר |  |
| 9:9.157 | הנמסר |  |
| 9:9.158 | להם |  |
| 9:9.159 | מידי |  |
| 9:9.160 | המלך |  |
| 9:9.161 | העכו"ם |  |
| 9:9.162 | שאמר |  |
| 9:9.163 | שאין |  |
| 9:9.164 | ראוי |  |
| 9:9.165 | לעבוד |  |
| 9:9.166 | אלא |  |
| 9:9.167 | לסבות |  |
| 9:9.168 | הנאצלים |  |
| 9:9.169 | ונמשכים |  |
| 9:9.170 | ממנו |  |
| 9:9.171 | לפי |  |
| 9:9.172 | דעתם | , |
| 9:9.173 | כאילן |  |
| 9:9.174 | המשלח |  |
| 9:9.175 | פארותיו |  |
| 9:9.176 | הילך |  |
| 9:9.177 | והילך | , |
| 9:9.178 | שהם |  |
| 9:9.179 | מושגים |  |
| 9:9.180 | להם |  |
| 9:9.181 | על |  |
| 9:9.182 | פי |  |
| 9:9.183 | רעיונותיהם | , |
| 9:9.184 | ותלו |  |
| 9:9.185 | כל |  |
| 9:9.186 | מצותיה |  |
| 9:9.187 | של |  |
| 9:9.188 | תורה |  |
| 9:9.189 | כאלו |  |
| 9:9.190 | הגופים |  |
| 9:9.191 | הנבראים |  |
| 9:9.192 | וקראום |  |
| 9:9.193 | גופא |  |
| 9:9.194 | דמלכא |  |
| 9:9.195 | יש |  |
| 9:9.196 | מצות |  |
| 9:9.197 | תלוים |  |
| 9:9.198 | בידא |  |
| 9:9.199 | דמלכא | , |
| 9:9.200 | וברגלוי | , |
| 9:9.201 | ושאר |  |
| 9:9.202 | אברים |  |
| 9:9.203 | שבגוף | , |
| 9:9.204 | ויש |  |
| 9:9.205 | מצות |  |
| 9:9.206 | אשר |  |
| 9:9.207 | לדעתם |  |
| 9:9.208 | הם |  |
| 9:9.209 | ביסוד | . |
| 9:9.210 | ערית |  |
| 9:9.211 | מלכא |  |
| 9:9.212 | דלא |  |
| 9:9.213 | אריך |  |
| 9:9.214 | למחזי |  |
| 9:9.215 | בטעם |  |
| 9:9.216 | כסוי |  |
| 9:9.217 | הדם |  |
| 9:9.218 | של |  |
| 9:9.219 | חיה |  |
| 9:9.220 | ועוף |  |
| 9:9.221 | כמו |  |
| 9:9.222 | שכתב |  |
| 9:9.223 | הרב |  |
| 9:9.224 | מקור |  |
| 9:9.225 | חיים |  |
| 9:9.226 | בטעם |  |
| 9:9.227 | כסוי |  |
| 9:9.228 | דמם | ! |
| 9:9.229 | ועל |  |
| 9:9.230 | כל |  |
| 9:9.231 | אלה |  |
| 9:9.232 | הענינים |  |
| 9:9.233 | יבכו |  |
| 9:9.234 | הכהנים |  |
| 9:9.235 | משרתי |  |
| 9:9.236 | ה | ' |
| 9:9.237 | בין |  |
| 9:9.238 | האולם |  |
| 9:9.239 | והמזבח |  |
| 9:9.240 | ויאמרו |  |
| 9:9.241 | חוסה |  |
| 9:9.242 | ה | ' |
| 9:9.243 | על |  |
| 9:9.244 | עמיך | . |
| 9:9.245 | והט |  |
| 9:9.246 | לבם |  |
| 9:9.247 | לדעת |  |
| 9:9.248 | ולהשכיל |  |
| 9:9.249 | אמתתך | . |
| 9:9.250 | ולדעת |  |
| 9:9.251 | את |  |
| 9:9.252 | דרכיך |  |
| 9:9.253 | אשר |  |
| 9:9.254 | צוית |  |
| 9:9.255 | את |  |
| 9:9.256 | משה |  |
| 9:9.257 | עבדיך |  |
| 9:9.258 | ללמד |  |
| 9:9.259 | לישראל |  |
| 9:9.260 | עלילותיך |  |
| 9:9.261 | ולעבדך |  |
| 9:9.262 | בלבב |  |
| 9:9.263 | שלם | . |
| 9:9.264 | ולא |  |
| 9:9.265 | לעבוד |  |
| 9:9.266 | לגופים |  |
| 9:9.267 | נבראים |  |
| 9:9.268 | שמכנים |  |
| 9:9.269 | אותם |  |
| 9:9.270 | בשם |  |
| 9:9.271 | אורות |  |
| 9:9.272 | אלהיים |  |
| 9:9.273 | בעבור |  |
| 9:9.274 | דבר |  |
| 9:9.275 | אלהי |  |
| 9:9.276 | המעורב |  |
| 9:9.277 | ומשותף |  |
| 9:9.278 | בהם |  |
| 9:9.279 | שהיא |  |
| 9:9.280 | נשמתם |  |
| 9:9.281 | כמו |  |
| 9:9.282 | שחשבו |  |
| 9:9.283 | המקובלים |  |
| 9:9.284 | החדשים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

