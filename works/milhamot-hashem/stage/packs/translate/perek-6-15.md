# Translation pack: milhamot-hashem perek-6-15 (Milchamot ha-Shem 6:24–6:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-15`.

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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
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

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- הסבה **הראשונה** → "the First Cause".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **גבורה** as a name of God → "the Power".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 6:24

Hebrew (Wikisource): עא) הוא שמשה רבי' ע"ה מוכיח את ישראל בס' משנה תורה, מפי הגבורה ויטש אלוה עשהו וינבל צור ישועתו, ויטש אלוה עשהו הוא הסבה הראשונה, וינבל צור ישועתו, שחשבוהו כעלה, נבל שאינו מושיע את הקורא אליו, ובמדרש הגדול פרשת האזינו, ויטש אלוה עשהו כענין שנאמר כי שתים רעות עשה עמי אותי עזבו מקור מים חים וכו' אמר להם המקום במדה שמדדתם אני מודד לכם שנאמר עזבתי את ביתי נטשתי את נחלתי, ואומר ויטש משכן שילה וימאס באוהל יוסף, ואומר כי נטשת עמך בית יעקב, יקניאוהו בזרים מלמד שעשו להם דברים של זרות, שנ' וגם מעכה אם אסא המלך הסירה מגבירה אשר עשתה לאשרה מפלצת, מאי מפלצת, אמר רב יהודה מפליא ליצנותא, אמר רב יוסף זכרות עשתה לה וכו'. ומחבר ספר הזהר עשה גם הוא מפלצות לבעלים ולעשתרות הם הפרצופים שהמציא בדמיונותיו לעבדם כדאיתא בזהר בראשית דף כ"ו ע"א וז"ל ויצמח ה' אלהים, אבא ואמא, כל עץ נחמד דא צדיק (היא יסוד דז"א) וטוב למאכל דא עמודא דאמצעיתא וכו'. ובפ' ויצא דף ק"ן ע"ב הוא קורא ליסוד דמלכות ויסוד דזעיר שער השמים וז"ל ויירא ויאמר מה נורא המקום הזה מלא דא לתרין סטרין איהו מה נורא המקום הזה חד ההוא מקום דקאמר בקדמייתא (פי' הר' ביבי נגה היינו מלכות) וחד על את קיימא (היינו יסוד) דלא בעיא לאתבטלא ואע"ג דאינון תרי סטרי חד הוא. אמר אין זה כי אם בית אלהים לאשתמשא ביה ולמיעבד ביה פירין ולארקא ביה הרכאן מכל שייפי גופא, דהאי הוא תרעא דכל גופא הח"ד וזה שער השמים, דא תרעא דגופה, ודאי תרעא איהו לארקא ברכאן לתתא דכתיב אין זה כי אם בית אלהים. ועל דא ויירא ויאמר מה נורא המקום הזה ובני נשא לא משגחן ביקרא דביה למהוי ביה שלים לעילא ותתא וכו', ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:24.1 | עא | ) |
| 6:24.2 | הוא |  |
| 6:24.3 | שמשה |  |
| 6:24.4 | רבי | ' |
| 6:24.5 | ע"ה |  |
| 6:24.6 | מוכיח |  |
| 6:24.7 | את |  |
| 6:24.8 | ישראל |  |
| 6:24.9 | בס | ' |
| 6:24.10 | משנה |  |
| 6:24.11 | תורה | , |
| 6:24.12 | מפי |  |
| 6:24.13 | הגבורה |  |
| 6:24.14 | ויטש |  |
| 6:24.15 | אלוה |  |
| 6:24.16 | עשהו |  |
| 6:24.17 | וינבל |  |
| 6:24.18 | צור |  |
| 6:24.19 | ישועתו | , |
| 6:24.20 | ויטש |  |
| 6:24.21 | אלוה |  |
| 6:24.22 | עשהו |  |
| 6:24.23 | הוא |  |
| 6:24.24 | הסבה |  |
| 6:24.25 | הראשונה | , |
| 6:24.26 | וינבל |  |
| 6:24.27 | צור |  |
| 6:24.28 | ישועתו | , |
| 6:24.29 | שחשבוהו |  |
| 6:24.30 | כעלה | , |
| 6:24.31 | נבל |  |
| 6:24.32 | שאינו |  |
| 6:24.33 | מושיע |  |
| 6:24.34 | את |  |
| 6:24.35 | הקורא |  |
| 6:24.36 | אליו | , |
| 6:24.37 | ובמדרש |  |
| 6:24.38 | הגדול |  |
| 6:24.39 | פרשת |  |
| 6:24.40 | האזינו | , |
| 6:24.41 | ויטש |  |
| 6:24.42 | אלוה |  |
| 6:24.43 | עשהו |  |
| 6:24.44 | כענין |  |
| 6:24.45 | שנאמר |  |
| 6:24.46 | כי |  |
| 6:24.47 | שתים |  |
| 6:24.48 | רעות |  |
| 6:24.49 | עשה |  |
| 6:24.50 | עמי |  |
| 6:24.51 | אותי |  |
| 6:24.52 | עזבו |  |
| 6:24.53 | מקור |  |
| 6:24.54 | מים |  |
| 6:24.55 | חים |  |
| 6:24.56 | וכו | ' |
| 6:24.57 | אמר |  |
| 6:24.58 | להם |  |
| 6:24.59 | המקום |  |
| 6:24.60 | במדה |  |
| 6:24.61 | שמדדתם |  |
| 6:24.62 | אני |  |
| 6:24.63 | מודד |  |
| 6:24.64 | לכם |  |
| 6:24.65 | שנאמר |  |
| 6:24.66 | עזבתי |  |
| 6:24.67 | את |  |
| 6:24.68 | ביתי |  |
| 6:24.69 | נטשתי |  |
| 6:24.70 | את |  |
| 6:24.71 | נחלתי | , |
| 6:24.72 | ואומר |  |
| 6:24.73 | ויטש |  |
| 6:24.74 | משכן |  |
| 6:24.75 | שילה |  |
| 6:24.76 | וימאס |  |
| 6:24.77 | באוהל |  |
| 6:24.78 | יוסף | , |
| 6:24.79 | ואומר |  |
| 6:24.80 | כי |  |
| 6:24.81 | נטשת |  |
| 6:24.82 | עמך |  |
| 6:24.83 | בית |  |
| 6:24.84 | יעקב | , |
| 6:24.85 | יקניאוהו |  |
| 6:24.86 | בזרים |  |
| 6:24.87 | מלמד |  |
| 6:24.88 | שעשו |  |
| 6:24.89 | להם |  |
| 6:24.90 | דברים |  |
| 6:24.91 | של |  |
| 6:24.92 | זרות | , |
| 6:24.93 | שנ | ' |
| 6:24.94 | וגם |  |
| 6:24.95 | מעכה |  |
| 6:24.96 | אם |  |
| 6:24.97 | אסא |  |
| 6:24.98 | המלך |  |
| 6:24.99 | הסירה |  |
| 6:24.100 | מגבירה |  |
| 6:24.101 | אשר |  |
| 6:24.102 | עשתה |  |
| 6:24.103 | לאשרה |  |
| 6:24.104 | מפלצת | , |
| 6:24.105 | מאי |  |
| 6:24.106 | מפלצת | , |
| 6:24.107 | אמר |  |
| 6:24.108 | רב |  |
| 6:24.109 | יהודה |  |
| 6:24.110 | מפליא |  |
| 6:24.111 | ליצנותא | , |
| 6:24.112 | אמר |  |
| 6:24.113 | רב |  |
| 6:24.114 | יוסף |  |
| 6:24.115 | זכרות |  |
| 6:24.116 | עשתה |  |
| 6:24.117 | לה |  |
| 6:24.118 | וכו | '. |
| 6:24.119 | ומחבר |  |
| 6:24.120 | ספר |  |
| 6:24.121 | הזהר |  |
| 6:24.122 | עשה |  |
| 6:24.123 | גם |  |
| 6:24.124 | הוא |  |
| 6:24.125 | מפלצות |  |
| 6:24.126 | לבעלים |  |
| 6:24.127 | ולעשתרות |  |
| 6:24.128 | הם |  |
| 6:24.129 | הפרצופים |  |
| 6:24.130 | שהמציא |  |
| 6:24.131 | בדמיונותיו |  |
| 6:24.132 | לעבדם |  |
| 6:24.133 | כדאיתא |  |
| 6:24.134 | בזהר |  |
| 6:24.135 | בראשית |  |
| 6:24.136 | דף |  |
| 6:24.137 | כ"ו |  |
| 6:24.138 | ע"א |  |
| 6:24.139 | וז"ל |  |
| 6:24.140 | ויצמח |  |
| 6:24.141 | ה | ' |
| 6:24.142 | אלהים | , |
| 6:24.143 | אבא |  |
| 6:24.144 | ואמא | , |
| 6:24.145 | כל |  |
| 6:24.146 | עץ |  |
| 6:24.147 | נחמד |  |
| 6:24.148 | דא |  |
| 6:24.149 | צדיק |  |
| 6:24.150 | היא |  |
| 6:24.151 | יסוד |  |
| 6:24.152 | דז"א | ) |
| 6:24.153 | וטוב |  |
| 6:24.154 | למאכל |  |
| 6:24.155 | דא |  |
| 6:24.156 | עמודא |  |
| 6:24.157 | דאמצעיתא |  |
| 6:24.158 | וכו | '. |
| 6:24.159 | ובפ | ' |
| 6:24.160 | ויצא |  |
| 6:24.161 | דף |  |
| 6:24.162 | ק"ן |  |
| 6:24.163 | ע"ב |  |
| 6:24.164 | הוא |  |
| 6:24.165 | קורא |  |
| 6:24.166 | ליסוד |  |
| 6:24.167 | דמלכות |  |
| 6:24.168 | ויסוד |  |
| 6:24.169 | דזעיר |  |
| 6:24.170 | שער |  |
| 6:24.171 | השמים |  |
| 6:24.172 | וז"ל |  |
| 6:24.173 | ויירא |  |
| 6:24.174 | ויאמר |  |
| 6:24.175 | מה |  |
| 6:24.176 | נורא |  |
| 6:24.177 | המקום |  |
| 6:24.178 | הזה |  |
| 6:24.179 | מלא |  |
| 6:24.180 | דא |  |
| 6:24.181 | לתרין |  |
| 6:24.182 | סטרין |  |
| 6:24.183 | איהו |  |
| 6:24.184 | מה |  |
| 6:24.185 | נורא |  |
| 6:24.186 | המקום |  |
| 6:24.187 | הזה |  |
| 6:24.188 | חד |  |
| 6:24.189 | ההוא |  |
| 6:24.190 | מקום |  |
| 6:24.191 | דקאמר |  |
| 6:24.192 | בקדמייתא |  |
| 6:24.193 | פי | ' |
| 6:24.194 | הר | ' |
| 6:24.195 | ביבי |  |
| 6:24.196 | נגה |  |
| 6:24.197 | היינו |  |
| 6:24.198 | מלכות | ) |
| 6:24.199 | וחד |  |
| 6:24.200 | על |  |
| 6:24.201 | את |  |
| 6:24.202 | קיימא |  |
| 6:24.203 | היינו |  |
| 6:24.204 | יסוד | ) |
| 6:24.205 | דלא |  |
| 6:24.206 | בעיא |  |
| 6:24.207 | לאתבטלא |  |
| 6:24.208 | ואע"ג |  |
| 6:24.209 | דאינון |  |
| 6:24.210 | תרי |  |
| 6:24.211 | סטרי |  |
| 6:24.212 | חד |  |
| 6:24.213 | הוא | . |
| 6:24.214 | אמר |  |
| 6:24.215 | אין |  |
| 6:24.216 | זה |  |
| 6:24.217 | כי |  |
| 6:24.218 | אם |  |
| 6:24.219 | בית |  |
| 6:24.220 | אלהים |  |
| 6:24.221 | לאשתמשא |  |
| 6:24.222 | ביה |  |
| 6:24.223 | ולמיעבד |  |
| 6:24.224 | ביה |  |
| 6:24.225 | פירין |  |
| 6:24.226 | ולארקא |  |
| 6:24.227 | ביה |  |
| 6:24.228 | הרכאן |  |
| 6:24.229 | מכל |  |
| 6:24.230 | שייפי |  |
| 6:24.231 | גופא | , |
| 6:24.232 | דהאי |  |
| 6:24.233 | הוא |  |
| 6:24.234 | תרעא |  |
| 6:24.235 | דכל |  |
| 6:24.236 | גופא |  |
| 6:24.237 | הח"ד |  |
| 6:24.238 | וזה |  |
| 6:24.239 | שער |  |
| 6:24.240 | השמים | , |
| 6:24.241 | דא |  |
| 6:24.242 | תרעא |  |
| 6:24.243 | דגופה | , |
| 6:24.244 | ודאי |  |
| 6:24.245 | תרעא |  |
| 6:24.246 | איהו |  |
| 6:24.247 | לארקא |  |
| 6:24.248 | ברכאן |  |
| 6:24.249 | לתתא |  |
| 6:24.250 | דכתיב |  |
| 6:24.251 | אין |  |
| 6:24.252 | זה |  |
| 6:24.253 | כי |  |
| 6:24.254 | אם |  |
| 6:24.255 | בית |  |
| 6:24.256 | אלהים | . |
| 6:24.257 | ועל |  |
| 6:24.258 | דא |  |
| 6:24.259 | ויירא |  |
| 6:24.260 | ויאמר |  |
| 6:24.261 | מה |  |
| 6:24.262 | נורא |  |
| 6:24.263 | המקום |  |
| 6:24.264 | הזה |  |
| 6:24.265 | ובני |  |
| 6:24.266 | נשא |  |
| 6:24.267 | לא |  |
| 6:24.268 | משגחן |  |
| 6:24.269 | ביקרא |  |
| 6:24.270 | דביה |  |
| 6:24.271 | למהוי |  |
| 6:24.272 | ביה |  |
| 6:24.273 | שלים |  |
| 6:24.274 | לעילא |  |
| 6:24.275 | ותתא |  |
| 6:24.276 | וכו | ', |
| 6:24.277 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

