# Translation pack: milhamot-hashem perek-4-13 (Milchamot ha-Shem 4:22–4:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-13`.

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
- **ב"ה** → "blessed be He".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **רז"ל** → "our Rabbis of blessed memory".
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **ר"ת** → "Rabbeinu Tam".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
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
- מי **אתי** (ketiv): translated, the qere noted.
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
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


## Milchamot ha-Shem 4:22

Hebrew (Wikisource): ובטול רעיונות אלו שבידי המקובלים האחרונים הנמשכים אחר רעיונות מחבר הזהר יתבאר ממה שאמרו רז"ל (שבת פ') אמר ר' עקיבה כשעלה משה לקבל את התורה אמרו מלאכי השרת לפני הקב"ה מה לילוד אשה בינינו? אמר להם לקבל תורה בא וכו', עד אמר לו הקב"ה למשה החזר להם תשובה, אמר לפניו מתירא אני שמא ישרפוני בהבל פיהם, אמר לו אחוז בכסא כבודי והחזר להם תשובה, ופירש הקב"ה מזיו שכינתו עליו ואמר לפניו רבונו של עולם תורה שאתה נותן להם מה כתיב בה, אנכי ה' אלהיך אשר הוצאתיך מארץ מצרים, אמר להם למצרים ירדתם? לפרעה נשתעבדתם? תורה למה לכם? ולפי דברי המקבולים החדשים, הרי אפי' אבא ואמא וזעיר ונוקביה נשתעבדו לפרעה (שהוא הסטרא אחרא) כל שכן המלאכים? ומבואר ג"כ שם במשנת חסידים בהדיא שזעיר אנפין בצער בגלות גדול כולד המעובר במעי אמו! ודבריהם סותרים זה את זה. שהרי אמרו שהזעיר אנפין הוא נותן התורה והוא שאמר אנכי ה' אלהיך אשר הוצאתיך מארץ מצרים. והכא קאמר שהקב"ה שהוא אריך אנפין הוציא את אבותינו שהם אבא ואימא וזעיר היה בצער וגלות? ולא תקשי לך ממאי דאמור רבנן גלו למצרים שכינה עמהם. גלו לבבל שכינה עמהם שנ' עמו אנכי בצרה. שכל זה הוא על דרך מאמר הכ' כי לא אחפוץ במות המת כי אם בשוב רשע מדרכו וחיה. ועוד שהרי כביכול מוכרח לקיים הבטחתו לנו בגלות. ואף גם זאת בהיותם בארץ אויביהם לא מאסתים ולא געלתים לכלותם וכו' והשגחתו ית' דבקה בנו בגלות שלא יעשו אויבינו אתנו כלה! ועוד ית' לפנינו דעת רז"ל שאין מציאות אל אחר בעולם כלל. ולא דברה תורה אלא לפי מחשבת עובדיהם.

| id | word | punctuation after |
|---|---|---|
| 4:22.1 | ובטול |  |
| 4:22.2 | רעיונות |  |
| 4:22.3 | אלו |  |
| 4:22.4 | שבידי |  |
| 4:22.5 | המקובלים |  |
| 4:22.6 | האחרונים |  |
| 4:22.7 | הנמשכים |  |
| 4:22.8 | אחר |  |
| 4:22.9 | רעיונות |  |
| 4:22.10 | מחבר |  |
| 4:22.11 | הזהר |  |
| 4:22.12 | יתבאר |  |
| 4:22.13 | ממה |  |
| 4:22.14 | שאמרו |  |
| 4:22.15 | רז"ל |  |
| 4:22.16 | שבת |  |
| 4:22.17 | פ | ') |
| 4:22.18 | אמר |  |
| 4:22.19 | ר | ' |
| 4:22.20 | עקיבה |  |
| 4:22.21 | כשעלה |  |
| 4:22.22 | משה |  |
| 4:22.23 | לקבל |  |
| 4:22.24 | את |  |
| 4:22.25 | התורה |  |
| 4:22.26 | אמרו |  |
| 4:22.27 | מלאכי |  |
| 4:22.28 | השרת |  |
| 4:22.29 | לפני |  |
| 4:22.30 | הקב"ה |  |
| 4:22.31 | מה |  |
| 4:22.32 | לילוד |  |
| 4:22.33 | אשה |  |
| 4:22.34 | בינינו | ? |
| 4:22.35 | אמר |  |
| 4:22.36 | להם |  |
| 4:22.37 | לקבל |  |
| 4:22.38 | תורה |  |
| 4:22.39 | בא |  |
| 4:22.40 | וכו | ', |
| 4:22.41 | עד |  |
| 4:22.42 | אמר |  |
| 4:22.43 | לו |  |
| 4:22.44 | הקב"ה |  |
| 4:22.45 | למשה |  |
| 4:22.46 | החזר |  |
| 4:22.47 | להם |  |
| 4:22.48 | תשובה | , |
| 4:22.49 | אמר |  |
| 4:22.50 | לפניו |  |
| 4:22.51 | מתירא |  |
| 4:22.52 | אני |  |
| 4:22.53 | שמא |  |
| 4:22.54 | ישרפוני |  |
| 4:22.55 | בהבל |  |
| 4:22.56 | פיהם | , |
| 4:22.57 | אמר |  |
| 4:22.58 | לו |  |
| 4:22.59 | אחוז |  |
| 4:22.60 | בכסא |  |
| 4:22.61 | כבודי |  |
| 4:22.62 | והחזר |  |
| 4:22.63 | להם |  |
| 4:22.64 | תשובה | , |
| 4:22.65 | ופירש |  |
| 4:22.66 | הקב"ה |  |
| 4:22.67 | מזיו |  |
| 4:22.68 | שכינתו |  |
| 4:22.69 | עליו |  |
| 4:22.70 | ואמר |  |
| 4:22.71 | לפניו |  |
| 4:22.72 | רבונו |  |
| 4:22.73 | של |  |
| 4:22.74 | עולם |  |
| 4:22.75 | תורה |  |
| 4:22.76 | שאתה |  |
| 4:22.77 | נותן |  |
| 4:22.78 | להם |  |
| 4:22.79 | מה |  |
| 4:22.80 | כתיב |  |
| 4:22.81 | בה | , |
| 4:22.82 | אנכי |  |
| 4:22.83 | ה | ' |
| 4:22.84 | אלהיך |  |
| 4:22.85 | אשר |  |
| 4:22.86 | הוצאתיך |  |
| 4:22.87 | מארץ |  |
| 4:22.88 | מצרים | , |
| 4:22.89 | אמר |  |
| 4:22.90 | להם |  |
| 4:22.91 | למצרים |  |
| 4:22.92 | ירדתם | ? |
| 4:22.93 | לפרעה |  |
| 4:22.94 | נשתעבדתם | ? |
| 4:22.95 | תורה |  |
| 4:22.96 | למה |  |
| 4:22.97 | לכם | ? |
| 4:22.98 | ולפי |  |
| 4:22.99 | דברי |  |
| 4:22.100 | המקבולים |  |
| 4:22.101 | החדשים | , |
| 4:22.102 | הרי |  |
| 4:22.103 | אפי | ' |
| 4:22.104 | אבא |  |
| 4:22.105 | ואמא |  |
| 4:22.106 | וזעיר |  |
| 4:22.107 | ונוקביה |  |
| 4:22.108 | נשתעבדו |  |
| 4:22.109 | לפרעה |  |
| 4:22.110 | שהוא |  |
| 4:22.111 | הסטרא |  |
| 4:22.112 | אחרא | ) |
| 4:22.113 | כל |  |
| 4:22.114 | שכן |  |
| 4:22.115 | המלאכים | ? |
| 4:22.116 | ומבואר |  |
| 4:22.117 | ג"כ |  |
| 4:22.118 | שם |  |
| 4:22.119 | במשנת |  |
| 4:22.120 | חסידים |  |
| 4:22.121 | בהדיא |  |
| 4:22.122 | שזעיר |  |
| 4:22.123 | אנפין |  |
| 4:22.124 | בצער |  |
| 4:22.125 | בגלות |  |
| 4:22.126 | גדול |  |
| 4:22.127 | כולד |  |
| 4:22.128 | המעובר |  |
| 4:22.129 | במעי |  |
| 4:22.130 | אמו | ! |
| 4:22.131 | ודבריהם |  |
| 4:22.132 | סותרים |  |
| 4:22.133 | זה |  |
| 4:22.134 | את |  |
| 4:22.135 | זה | . |
| 4:22.136 | שהרי |  |
| 4:22.137 | אמרו |  |
| 4:22.138 | שהזעיר |  |
| 4:22.139 | אנפין |  |
| 4:22.140 | הוא |  |
| 4:22.141 | נותן |  |
| 4:22.142 | התורה |  |
| 4:22.143 | והוא |  |
| 4:22.144 | שאמר |  |
| 4:22.145 | אנכי |  |
| 4:22.146 | ה | ' |
| 4:22.147 | אלהיך |  |
| 4:22.148 | אשר |  |
| 4:22.149 | הוצאתיך |  |
| 4:22.150 | מארץ |  |
| 4:22.151 | מצרים | . |
| 4:22.152 | והכא |  |
| 4:22.153 | קאמר |  |
| 4:22.154 | שהקב"ה |  |
| 4:22.155 | שהוא |  |
| 4:22.156 | אריך |  |
| 4:22.157 | אנפין |  |
| 4:22.158 | הוציא |  |
| 4:22.159 | את |  |
| 4:22.160 | אבותינו |  |
| 4:22.161 | שהם |  |
| 4:22.162 | אבא |  |
| 4:22.163 | ואימא |  |
| 4:22.164 | וזעיר |  |
| 4:22.165 | היה |  |
| 4:22.166 | בצער |  |
| 4:22.167 | וגלות | ? |
| 4:22.168 | ולא |  |
| 4:22.169 | תקשי |  |
| 4:22.170 | לך |  |
| 4:22.171 | ממאי |  |
| 4:22.172 | דאמור |  |
| 4:22.173 | רבנן |  |
| 4:22.174 | גלו |  |
| 4:22.175 | למצרים |  |
| 4:22.176 | שכינה |  |
| 4:22.177 | עמהם | . |
| 4:22.178 | גלו |  |
| 4:22.179 | לבבל |  |
| 4:22.180 | שכינה |  |
| 4:22.181 | עמהם |  |
| 4:22.182 | שנ | ' |
| 4:22.183 | עמו |  |
| 4:22.184 | אנכי |  |
| 4:22.185 | בצרה | . |
| 4:22.186 | שכל |  |
| 4:22.187 | זה |  |
| 4:22.188 | הוא |  |
| 4:22.189 | על |  |
| 4:22.190 | דרך |  |
| 4:22.191 | מאמר |  |
| 4:22.192 | הכ | ' |
| 4:22.193 | כי |  |
| 4:22.194 | לא |  |
| 4:22.195 | אחפוץ |  |
| 4:22.196 | במות |  |
| 4:22.197 | המת |  |
| 4:22.198 | כי |  |
| 4:22.199 | אם |  |
| 4:22.200 | בשוב |  |
| 4:22.201 | רשע |  |
| 4:22.202 | מדרכו |  |
| 4:22.203 | וחיה | . |
| 4:22.204 | ועוד |  |
| 4:22.205 | שהרי |  |
| 4:22.206 | כביכול |  |
| 4:22.207 | מוכרח |  |
| 4:22.208 | לקיים |  |
| 4:22.209 | הבטחתו |  |
| 4:22.210 | לנו |  |
| 4:22.211 | בגלות | . |
| 4:22.212 | ואף |  |
| 4:22.213 | גם |  |
| 4:22.214 | זאת |  |
| 4:22.215 | בהיותם |  |
| 4:22.216 | בארץ |  |
| 4:22.217 | אויביהם |  |
| 4:22.218 | לא |  |
| 4:22.219 | מאסתים |  |
| 4:22.220 | ולא |  |
| 4:22.221 | געלתים |  |
| 4:22.222 | לכלותם |  |
| 4:22.223 | וכו | ' |
| 4:22.224 | והשגחתו |  |
| 4:22.225 | ית | ' |
| 4:22.226 | דבקה |  |
| 4:22.227 | בנו |  |
| 4:22.228 | בגלות |  |
| 4:22.229 | שלא |  |
| 4:22.230 | יעשו |  |
| 4:22.231 | אויבינו |  |
| 4:22.232 | אתנו |  |
| 4:22.233 | כלה | ! |
| 4:22.234 | ועוד |  |
| 4:22.235 | ית | ' |
| 4:22.236 | לפנינו |  |
| 4:22.237 | דעת |  |
| 4:22.238 | רז"ל |  |
| 4:22.239 | שאין |  |
| 4:22.240 | מציאות |  |
| 4:22.241 | אל |  |
| 4:22.242 | אחר |  |
| 4:22.243 | בעולם |  |
| 4:22.244 | כלל | . |
| 4:22.245 | ולא |  |
| 4:22.246 | דברה |  |
| 4:22.247 | תורה |  |
| 4:22.248 | אלא |  |
| 4:22.249 | לפי |  |
| 4:22.250 | מחשבת |  |
| 4:22.251 | עובדיהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

