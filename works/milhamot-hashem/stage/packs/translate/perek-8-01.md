# Translation pack: milhamot-hashem perek-8-01 (Milchamot ha-Shem 8:1–8:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-01`.

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
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **אדרא / אידרא** of Naso → "Idra of Naso" (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 8:1

Hebrew (Wikisource): ק) וראיתי להרב נחלת יוסף לר' שמואל איש עדן שהביא קצור אמונת הקבלה החדשה, וטרח לישב קצת סתירות במשלים שאינם מתקבלים על השכל הישר. ואין בהם כדאי לאוהב את יחיד השי"ת על דרך האמת אשר הורונו רז"ל על פי תורתינו הקדושה, כמו שיראה המעיין שם במשל שהביא שם! וגם הוא סובר שכל עבודתינו ותפלותינו וברכותינו הם לאלוה קצר אפים (זעיר אנפין) והוא אלהינו ואנחנו עמו ונחלתו, והוא המלך הקדוש המהולל בכל התושבחות. השולט על כל הברואים וזנם ומפרנסם ומנהיגם. ושזעיר אנפין הוא שירד על הר סיני להתחבר עם המלכות שנקראת אש. שכבר קדמה וירדה על הר סיני. ואמר אנכי ה' אלהיך אשר הוציאתיך מארץ מצרים. היא אימא עילאה בינה. כי היא הוציאתם ממצרים על די המלכות שנתנה לה גבורותיה, ומלכות היא המכה במצרים בסיועת הבינה, וכן בסיני על ידי הבור בינה עם תפארת ומלכות נאמר אנכי ה' אלהיך. והתורה יצאה מחכמה עלאה עם כללות הבינה. והכל בכח אין סוף ב"ה וכו' ע"ש דף ס"א וס"ב. ודבריו ג"ד מכיינים לדברי מהר"י לופיס לעיל סי' כ"ו שזעיר אנפין הוא שנגלה על אבותינו בסיני ואמר אנכי ה' אלהיך.

| id | word | punctuation after |
|---|---|---|
| 8:1.1 | ק | ) |
| 8:1.2 | וראיתי |  |
| 8:1.3 | להרב |  |
| 8:1.4 | נחלת |  |
| 8:1.5 | יוסף |  |
| 8:1.6 | לר | ' |
| 8:1.7 | שמואל |  |
| 8:1.8 | איש |  |
| 8:1.9 | עדן |  |
| 8:1.10 | שהביא |  |
| 8:1.11 | קצור |  |
| 8:1.12 | אמונת |  |
| 8:1.13 | הקבלה |  |
| 8:1.14 | החדשה | , |
| 8:1.15 | וטרח |  |
| 8:1.16 | לישב |  |
| 8:1.17 | קצת |  |
| 8:1.18 | סתירות |  |
| 8:1.19 | במשלים |  |
| 8:1.20 | שאינם |  |
| 8:1.21 | מתקבלים |  |
| 8:1.22 | על |  |
| 8:1.23 | השכל |  |
| 8:1.24 | הישר | . |
| 8:1.25 | ואין |  |
| 8:1.26 | בהם |  |
| 8:1.27 | כדאי |  |
| 8:1.28 | לאוהב |  |
| 8:1.29 | את |  |
| 8:1.30 | יחיד |  |
| 8:1.31 | השי"ת |  |
| 8:1.32 | על |  |
| 8:1.33 | דרך |  |
| 8:1.34 | האמת |  |
| 8:1.35 | אשר |  |
| 8:1.36 | הורונו |  |
| 8:1.37 | רז"ל |  |
| 8:1.38 | על |  |
| 8:1.39 | פי |  |
| 8:1.40 | תורתינו |  |
| 8:1.41 | הקדושה | , |
| 8:1.42 | כמו |  |
| 8:1.43 | שיראה |  |
| 8:1.44 | המעיין |  |
| 8:1.45 | שם |  |
| 8:1.46 | במשל |  |
| 8:1.47 | שהביא |  |
| 8:1.48 | שם | ! |
| 8:1.49 | וגם |  |
| 8:1.50 | הוא |  |
| 8:1.51 | סובר |  |
| 8:1.52 | שכל |  |
| 8:1.53 | עבודתינו |  |
| 8:1.54 | ותפלותינו |  |
| 8:1.55 | וברכותינו |  |
| 8:1.56 | הם |  |
| 8:1.57 | לאלוה |  |
| 8:1.58 | קצר |  |
| 8:1.59 | אפים |  |
| 8:1.60 | זעיר |  |
| 8:1.61 | אנפין | ) |
| 8:1.62 | והוא |  |
| 8:1.63 | אלהינו |  |
| 8:1.64 | ואנחנו |  |
| 8:1.65 | עמו |  |
| 8:1.66 | ונחלתו | , |
| 8:1.67 | והוא |  |
| 8:1.68 | המלך |  |
| 8:1.69 | הקדוש |  |
| 8:1.70 | המהולל |  |
| 8:1.71 | בכל |  |
| 8:1.72 | התושבחות | . |
| 8:1.73 | השולט |  |
| 8:1.74 | על |  |
| 8:1.75 | כל |  |
| 8:1.76 | הברואים |  |
| 8:1.77 | וזנם |  |
| 8:1.78 | ומפרנסם |  |
| 8:1.79 | ומנהיגם | . |
| 8:1.80 | ושזעיר |  |
| 8:1.81 | אנפין |  |
| 8:1.82 | הוא |  |
| 8:1.83 | שירד |  |
| 8:1.84 | על |  |
| 8:1.85 | הר |  |
| 8:1.86 | סיני |  |
| 8:1.87 | להתחבר |  |
| 8:1.88 | עם |  |
| 8:1.89 | המלכות |  |
| 8:1.90 | שנקראת |  |
| 8:1.91 | אש | . |
| 8:1.92 | שכבר |  |
| 8:1.93 | קדמה |  |
| 8:1.94 | וירדה |  |
| 8:1.95 | על |  |
| 8:1.96 | הר |  |
| 8:1.97 | סיני | . |
| 8:1.98 | ואמר |  |
| 8:1.99 | אנכי |  |
| 8:1.100 | ה | ' |
| 8:1.101 | אלהיך |  |
| 8:1.102 | אשר |  |
| 8:1.103 | הוציאתיך |  |
| 8:1.104 | מארץ |  |
| 8:1.105 | מצרים | . |
| 8:1.106 | היא |  |
| 8:1.107 | אימא |  |
| 8:1.108 | עילאה |  |
| 8:1.109 | בינה | . |
| 8:1.110 | כי |  |
| 8:1.111 | היא |  |
| 8:1.112 | הוציאתם |  |
| 8:1.113 | ממצרים |  |
| 8:1.114 | על |  |
| 8:1.115 | די |  |
| 8:1.116 | המלכות |  |
| 8:1.117 | שנתנה |  |
| 8:1.118 | לה |  |
| 8:1.119 | גבורותיה | , |
| 8:1.120 | ומלכות |  |
| 8:1.121 | היא |  |
| 8:1.122 | המכה |  |
| 8:1.123 | במצרים |  |
| 8:1.124 | בסיועת |  |
| 8:1.125 | הבינה | , |
| 8:1.126 | וכן |  |
| 8:1.127 | בסיני |  |
| 8:1.128 | על |  |
| 8:1.129 | ידי |  |
| 8:1.130 | הבור |  |
| 8:1.131 | בינה |  |
| 8:1.132 | עם |  |
| 8:1.133 | תפארת |  |
| 8:1.134 | ומלכות |  |
| 8:1.135 | נאמר |  |
| 8:1.136 | אנכי |  |
| 8:1.137 | ה | ' |
| 8:1.138 | אלהיך | . |
| 8:1.139 | והתורה |  |
| 8:1.140 | יצאה |  |
| 8:1.141 | מחכמה |  |
| 8:1.142 | עלאה |  |
| 8:1.143 | עם |  |
| 8:1.144 | כללות |  |
| 8:1.145 | הבינה | . |
| 8:1.146 | והכל |  |
| 8:1.147 | בכח |  |
| 8:1.148 | אין |  |
| 8:1.149 | סוף |  |
| 8:1.150 | ב"ה |  |
| 8:1.151 | וכו | ' |
| 8:1.152 | ע"ש |  |
| 8:1.153 | דף |  |
| 8:1.154 | ס"א |  |
| 8:1.155 | וס"ב | . |
| 8:1.156 | ודבריו |  |
| 8:1.157 | ג"ד |  |
| 8:1.158 | מכיינים |  |
| 8:1.159 | לדברי |  |
| 8:1.160 | מהר"י |  |
| 8:1.161 | לופיס |  |
| 8:1.162 | לעיל |  |
| 8:1.163 | סי | ' |
| 8:1.164 | כ"ו |  |
| 8:1.165 | שזעיר |  |
| 8:1.166 | אנפין |  |
| 8:1.167 | הוא |  |
| 8:1.168 | שנגלה |  |
| 8:1.169 | על |  |
| 8:1.170 | אבותינו |  |
| 8:1.171 | בסיני |  |
| 8:1.172 | ואמר |  |
| 8:1.173 | אנכי |  |
| 8:1.174 | ה | ' |
| 8:1.175 | אלהיך | . |


## Milchamot ha-Shem 8:2

Hebrew (Wikisource): והנה כבר הודעתיך בסי' מ' שכל אין סוף שהזכירו המקובלים החדשים הוא החלק הגנוז בכתר הנקרא עתיקא דכל עתיקין בלבד, אבל עדיין יש אין סוף שהוא לעילא לעילא. והוא המקיף את כל אשר בתוכו מכל הנבראים ואין דבר חוצה לו, ורק אל האין סוף אשר בכתר מליהם ירמזון, וממנו הם ממשיכים בהתפללם אל קצר אפים (זעיר אנפין). ואל עתיקא קדישא לפי דבריהם אמר הנביא, ה' שהוא עתיק, פעלך שהוא זעיר אנפין חייהו וכו' כמו שנתבאר לעיל סוף סימן מ"ח ב' האדרא רבה דף קל"ח. והעולה מדברי הרב נח"י הוא הדבר שכבר זכרתי לך פעמים רבות. שדעת המקובלים החדשים שכל עבודתינו ותפלותינו וברכותינו להאלוה קצר אפים, והזעיר אנפין הוא שנגלה על אבותינו בהר סיני ונתן לנו את התורה, ומפוי שמעו אנכי ה' אלהיך, הוא ולא אבא ואימא, ולא אריך אנפין, ולא עתיק יומין, ולא אדם קדמון, וכ"ש שאינו האין סוף שהוא לעילא לעילא.

| id | word | punctuation after |
|---|---|---|
| 8:2.1 | והנה |  |
| 8:2.2 | כבר |  |
| 8:2.3 | הודעתיך |  |
| 8:2.4 | בסי | ' |
| 8:2.5 | מ | ' |
| 8:2.6 | שכל |  |
| 8:2.7 | אין |  |
| 8:2.8 | סוף |  |
| 8:2.9 | שהזכירו |  |
| 8:2.10 | המקובלים |  |
| 8:2.11 | החדשים |  |
| 8:2.12 | הוא |  |
| 8:2.13 | החלק |  |
| 8:2.14 | הגנוז |  |
| 8:2.15 | בכתר |  |
| 8:2.16 | הנקרא |  |
| 8:2.17 | עתיקא |  |
| 8:2.18 | דכל |  |
| 8:2.19 | עתיקין |  |
| 8:2.20 | בלבד | , |
| 8:2.21 | אבל |  |
| 8:2.22 | עדיין |  |
| 8:2.23 | יש |  |
| 8:2.24 | אין |  |
| 8:2.25 | סוף |  |
| 8:2.26 | שהוא |  |
| 8:2.27 | לעילא |  |
| 8:2.28 | לעילא | . |
| 8:2.29 | והוא |  |
| 8:2.30 | המקיף |  |
| 8:2.31 | את |  |
| 8:2.32 | כל |  |
| 8:2.33 | אשר |  |
| 8:2.34 | בתוכו |  |
| 8:2.35 | מכל |  |
| 8:2.36 | הנבראים |  |
| 8:2.37 | ואין |  |
| 8:2.38 | דבר |  |
| 8:2.39 | חוצה |  |
| 8:2.40 | לו | , |
| 8:2.41 | ורק |  |
| 8:2.42 | אל |  |
| 8:2.43 | האין |  |
| 8:2.44 | סוף |  |
| 8:2.45 | אשר |  |
| 8:2.46 | בכתר |  |
| 8:2.47 | מליהם |  |
| 8:2.48 | ירמזון | , |
| 8:2.49 | וממנו |  |
| 8:2.50 | הם |  |
| 8:2.51 | ממשיכים |  |
| 8:2.52 | בהתפללם |  |
| 8:2.53 | אל |  |
| 8:2.54 | קצר |  |
| 8:2.55 | אפים |  |
| 8:2.56 | זעיר |  |
| 8:2.57 | אנפין | ). |
| 8:2.58 | ואל |  |
| 8:2.59 | עתיקא |  |
| 8:2.60 | קדישא |  |
| 8:2.61 | לפי |  |
| 8:2.62 | דבריהם |  |
| 8:2.63 | אמר |  |
| 8:2.64 | הנביא | , |
| 8:2.65 | ה | ' |
| 8:2.66 | שהוא |  |
| 8:2.67 | עתיק | , |
| 8:2.68 | פעלך |  |
| 8:2.69 | שהוא |  |
| 8:2.70 | זעיר |  |
| 8:2.71 | אנפין |  |
| 8:2.72 | חייהו |  |
| 8:2.73 | וכו | ' |
| 8:2.74 | כמו |  |
| 8:2.75 | שנתבאר |  |
| 8:2.76 | לעיל |  |
| 8:2.77 | סוף |  |
| 8:2.78 | סימן |  |
| 8:2.79 | מ"ח |  |
| 8:2.80 | ב | ' |
| 8:2.81 | האדרא |  |
| 8:2.82 | רבה |  |
| 8:2.83 | דף |  |
| 8:2.84 | קל"ח | . |
| 8:2.85 | והעולה |  |
| 8:2.86 | מדברי |  |
| 8:2.87 | הרב |  |
| 8:2.88 | נח"י |  |
| 8:2.89 | הוא |  |
| 8:2.90 | הדבר |  |
| 8:2.91 | שכבר |  |
| 8:2.92 | זכרתי |  |
| 8:2.93 | לך |  |
| 8:2.94 | פעמים |  |
| 8:2.95 | רבות | . |
| 8:2.96 | שדעת |  |
| 8:2.97 | המקובלים |  |
| 8:2.98 | החדשים |  |
| 8:2.99 | שכל |  |
| 8:2.100 | עבודתינו |  |
| 8:2.101 | ותפלותינו |  |
| 8:2.102 | וברכותינו |  |
| 8:2.103 | להאלוה |  |
| 8:2.104 | קצר |  |
| 8:2.105 | אפים | , |
| 8:2.106 | והזעיר |  |
| 8:2.107 | אנפין |  |
| 8:2.108 | הוא |  |
| 8:2.109 | שנגלה |  |
| 8:2.110 | על |  |
| 8:2.111 | אבותינו |  |
| 8:2.112 | בהר |  |
| 8:2.113 | סיני |  |
| 8:2.114 | ונתן |  |
| 8:2.115 | לנו |  |
| 8:2.116 | את |  |
| 8:2.117 | התורה | , |
| 8:2.118 | ומפוי |  |
| 8:2.119 | שמעו |  |
| 8:2.120 | אנכי |  |
| 8:2.121 | ה | ' |
| 8:2.122 | אלהיך | , |
| 8:2.123 | הוא |  |
| 8:2.124 | ולא |  |
| 8:2.125 | אבא |  |
| 8:2.126 | ואימא | , |
| 8:2.127 | ולא |  |
| 8:2.128 | אריך |  |
| 8:2.129 | אנפין | , |
| 8:2.130 | ולא |  |
| 8:2.131 | עתיק |  |
| 8:2.132 | יומין | , |
| 8:2.133 | ולא |  |
| 8:2.134 | אדם |  |
| 8:2.135 | קדמון | , |
| 8:2.136 | וכ"ש |  |
| 8:2.137 | שאינו |  |
| 8:2.138 | האין |  |
| 8:2.139 | סוף |  |
| 8:2.140 | שהוא |  |
| 8:2.141 | לעילא |  |
| 8:2.142 | לעילא | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

