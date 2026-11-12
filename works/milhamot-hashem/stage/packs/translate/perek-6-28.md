# Translation pack: milhamot-hashem perek-6-28 (Milchamot ha-Shem 6:42–6:42)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-28.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-28`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
## Terms

- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **תוה"ק** → "the holy Torah".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **דרך ארץ** → "worldly occupation".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
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


## Milchamot ha-Shem 6:42

Hebrew (Wikisource): וכתב מהר"י אלבו בספרו העיקרים מאמר ג' פרק כ"ג. כי מאמר תוה"ק אחרי רבים להטות. ובלבד שיהיו הרוב מן החכמים, לא רוב דעות עמי הארץ, לפי שההמון ועמי הארץ נפתים לדבר שאינו. עד שיבואו להעיד עליו ולומר שהוא כן. ועל כן אין ראוי להביא ראיה מהסכמת ההמון. כי פעמים יסכימו ההמון על חילוף האמת (כנדון דידן) שהרי בימי אחאב ומנשה ודומיהם היו מסכימים כולם לעבוד ע"ז. זולתי הנביאים ותלמידהם. והיה ראוי שיהיו אליהו ומאה איש שהחביא עובדיה המה הרשעים, כי הם המיעוט, והרוב שהם עובדי ע"ז הם הצדיקים! אלא ודאי ראוי שתהיה ההכרעה מסורה לחכמים בלבד, כי החכמה מתת אלהים היא. כי ה' יתן חכמה ומפיו דעת ותבונה. ובזה תהיה תורת ה' תמימה בכל דור ודור, ובלתי חסרה דבר עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:42.1 | וכתב |  |
| 6:42.2 | מהר"י |  |
| 6:42.3 | אלבו |  |
| 6:42.4 | בספרו |  |
| 6:42.5 | העיקרים |  |
| 6:42.6 | מאמר |  |
| 6:42.7 | ג | ' |
| 6:42.8 | פרק |  |
| 6:42.9 | כ"ג | . |
| 6:42.10 | כי |  |
| 6:42.11 | מאמר |  |
| 6:42.12 | תוה"ק |  |
| 6:42.13 | אחרי |  |
| 6:42.14 | רבים |  |
| 6:42.15 | להטות | . |
| 6:42.16 | ובלבד |  |
| 6:42.17 | שיהיו |  |
| 6:42.18 | הרוב |  |
| 6:42.19 | מן |  |
| 6:42.20 | החכמים | , |
| 6:42.21 | לא |  |
| 6:42.22 | רוב |  |
| 6:42.23 | דעות |  |
| 6:42.24 | עמי |  |
| 6:42.25 | הארץ | , |
| 6:42.26 | לפי |  |
| 6:42.27 | שההמון |  |
| 6:42.28 | ועמי |  |
| 6:42.29 | הארץ |  |
| 6:42.30 | נפתים |  |
| 6:42.31 | לדבר |  |
| 6:42.32 | שאינו | . |
| 6:42.33 | עד |  |
| 6:42.34 | שיבואו |  |
| 6:42.35 | להעיד |  |
| 6:42.36 | עליו |  |
| 6:42.37 | ולומר |  |
| 6:42.38 | שהוא |  |
| 6:42.39 | כן | . |
| 6:42.40 | ועל |  |
| 6:42.41 | כן |  |
| 6:42.42 | אין |  |
| 6:42.43 | ראוי |  |
| 6:42.44 | להביא |  |
| 6:42.45 | ראיה |  |
| 6:42.46 | מהסכמת |  |
| 6:42.47 | ההמון | . |
| 6:42.48 | כי |  |
| 6:42.49 | פעמים |  |
| 6:42.50 | יסכימו |  |
| 6:42.51 | ההמון |  |
| 6:42.52 | על |  |
| 6:42.53 | חילוף |  |
| 6:42.54 | האמת |  |
| 6:42.55 | כנדון |  |
| 6:42.56 | דידן | ) |
| 6:42.57 | שהרי |  |
| 6:42.58 | בימי |  |
| 6:42.59 | אחאב |  |
| 6:42.60 | ומנשה |  |
| 6:42.61 | ודומיהם |  |
| 6:42.62 | היו |  |
| 6:42.63 | מסכימים |  |
| 6:42.64 | כולם |  |
| 6:42.65 | לעבוד |  |
| 6:42.66 | ע"ז | . |
| 6:42.67 | זולתי |  |
| 6:42.68 | הנביאים |  |
| 6:42.69 | ותלמידהם | . |
| 6:42.70 | והיה |  |
| 6:42.71 | ראוי |  |
| 6:42.72 | שיהיו |  |
| 6:42.73 | אליהו |  |
| 6:42.74 | ומאה |  |
| 6:42.75 | איש |  |
| 6:42.76 | שהחביא |  |
| 6:42.77 | עובדיה |  |
| 6:42.78 | המה |  |
| 6:42.79 | הרשעים | , |
| 6:42.80 | כי |  |
| 6:42.81 | הם |  |
| 6:42.82 | המיעוט | , |
| 6:42.83 | והרוב |  |
| 6:42.84 | שהם |  |
| 6:42.85 | עובדי |  |
| 6:42.86 | ע"ז |  |
| 6:42.87 | הם |  |
| 6:42.88 | הצדיקים | ! |
| 6:42.89 | אלא |  |
| 6:42.90 | ודאי |  |
| 6:42.91 | ראוי |  |
| 6:42.92 | שתהיה |  |
| 6:42.93 | ההכרעה |  |
| 6:42.94 | מסורה |  |
| 6:42.95 | לחכמים |  |
| 6:42.96 | בלבד | , |
| 6:42.97 | כי |  |
| 6:42.98 | החכמה |  |
| 6:42.99 | מתת |  |
| 6:42.100 | אלהים |  |
| 6:42.101 | היא | . |
| 6:42.102 | כי |  |
| 6:42.103 | ה | ' |
| 6:42.104 | יתן |  |
| 6:42.105 | חכמה |  |
| 6:42.106 | ומפיו |  |
| 6:42.107 | דעת |  |
| 6:42.108 | ותבונה | . |
| 6:42.109 | ובזה |  |
| 6:42.110 | תהיה |  |
| 6:42.111 | תורת |  |
| 6:42.112 | ה | ' |
| 6:42.113 | תמימה |  |
| 6:42.114 | בכל |  |
| 6:42.115 | דור |  |
| 6:42.116 | ודור | , |
| 6:42.117 | ובלתי |  |
| 6:42.118 | חסרה |  |
| 6:42.119 | דבר |  |
| 6:42.120 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

