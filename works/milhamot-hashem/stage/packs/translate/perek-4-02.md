# Translation pack: milhamot-hashem perek-4-02 (Milchamot ha-Shem 4:2–4:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-02`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
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

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- סבא **דסבין** → "the Elder of Elders".
- עלת **העלות** → "the Cause of causes".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- ריש **מלין** אמר (Dan 7:1) → "he stated the sum of the matter".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
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


## Milchamot ha-Shem 4:2

Hebrew (Wikisource): מבואר מזה, כי הקב"ה הנקרא עתיק נגלה לר' שמעון בבית מדרשו ובדמות זקן מופלג, ולכך קרא לו ר' שמעון סבא דסבין רוצה לומר זקן הזקנים כדלעיל בשם דרך אמת ומקדש מלך, והקב"ה זה הנקרא עתיק הוא שהרשה לר' שמעון לגלות פנים בתורה, שהקב"ה הנקרא אמא היא שאמרה להקב"ה הנקרא אבא נעשה אדם בצלמינו בדמותינו וכו'. הרי נתבאר לך שהוא קורא לעתיק קב"ה וכן לאבא, וקורא לאמא אלהים, וגם קורא לאבא אלהים וגם קב"ה, ובכל מעשה בראשית אבא אומר לאמא יהי כך וכך והיא עושה, ובמאמר נעשה אדם, אימא אמרה לאבא נעשה אדם ולא הסכים אבא בבריאת אדם. ואמרה לו אמא, מאי נפקא לך מניה, אם יחטא לי יחטא ולא לך שנ' ובן כסיל תוגת אמו, וכשחטא אדם בעץ הדעת גירש אותו ואותה שנאמר ובפשעיכם שולחה אמכם.

| id | word | punctuation after |
|---|---|---|
| 4:2.1 | מבואר |  |
| 4:2.2 | מזה | , |
| 4:2.3 | כי |  |
| 4:2.4 | הקב"ה |  |
| 4:2.5 | הנקרא |  |
| 4:2.6 | עתיק |  |
| 4:2.7 | נגלה |  |
| 4:2.8 | לר | ' |
| 4:2.9 | שמעון |  |
| 4:2.10 | בבית |  |
| 4:2.11 | מדרשו |  |
| 4:2.12 | ובדמות |  |
| 4:2.13 | זקן |  |
| 4:2.14 | מופלג | , |
| 4:2.15 | ולכך |  |
| 4:2.16 | קרא |  |
| 4:2.17 | לו |  |
| 4:2.18 | ר | ' |
| 4:2.19 | שמעון |  |
| 4:2.20 | סבא |  |
| 4:2.21 | דסבין |  |
| 4:2.22 | רוצה |  |
| 4:2.23 | לומר |  |
| 4:2.24 | זקן |  |
| 4:2.25 | הזקנים |  |
| 4:2.26 | כדלעיל |  |
| 4:2.27 | בשם |  |
| 4:2.28 | דרך |  |
| 4:2.29 | אמת |  |
| 4:2.30 | ומקדש |  |
| 4:2.31 | מלך | , |
| 4:2.32 | והקב"ה |  |
| 4:2.33 | זה |  |
| 4:2.34 | הנקרא |  |
| 4:2.35 | עתיק |  |
| 4:2.36 | הוא |  |
| 4:2.37 | שהרשה |  |
| 4:2.38 | לר | ' |
| 4:2.39 | שמעון |  |
| 4:2.40 | לגלות |  |
| 4:2.41 | פנים |  |
| 4:2.42 | בתורה | , |
| 4:2.43 | שהקב"ה |  |
| 4:2.44 | הנקרא |  |
| 4:2.45 | אמא |  |
| 4:2.46 | היא |  |
| 4:2.47 | שאמרה |  |
| 4:2.48 | להקב"ה |  |
| 4:2.49 | הנקרא |  |
| 4:2.50 | אבא |  |
| 4:2.51 | נעשה |  |
| 4:2.52 | אדם |  |
| 4:2.53 | בצלמינו |  |
| 4:2.54 | בדמותינו |  |
| 4:2.55 | וכו | '. |
| 4:2.56 | הרי |  |
| 4:2.57 | נתבאר |  |
| 4:2.58 | לך |  |
| 4:2.59 | שהוא |  |
| 4:2.60 | קורא |  |
| 4:2.61 | לעתיק |  |
| 4:2.62 | קב"ה |  |
| 4:2.63 | וכן |  |
| 4:2.64 | לאבא | , |
| 4:2.65 | וקורא |  |
| 4:2.66 | לאמא |  |
| 4:2.67 | אלהים | , |
| 4:2.68 | וגם |  |
| 4:2.69 | קורא |  |
| 4:2.70 | לאבא |  |
| 4:2.71 | אלהים |  |
| 4:2.72 | וגם |  |
| 4:2.73 | קב"ה | , |
| 4:2.74 | ובכל |  |
| 4:2.75 | מעשה |  |
| 4:2.76 | בראשית |  |
| 4:2.77 | אבא |  |
| 4:2.78 | אומר |  |
| 4:2.79 | לאמא |  |
| 4:2.80 | יהי |  |
| 4:2.81 | כך |  |
| 4:2.82 | וכך |  |
| 4:2.83 | והיא |  |
| 4:2.84 | עושה | , |
| 4:2.85 | ובמאמר |  |
| 4:2.86 | נעשה |  |
| 4:2.87 | אדם | , |
| 4:2.88 | אימא |  |
| 4:2.89 | אמרה |  |
| 4:2.90 | לאבא |  |
| 4:2.91 | נעשה |  |
| 4:2.92 | אדם |  |
| 4:2.93 | ולא |  |
| 4:2.94 | הסכים |  |
| 4:2.95 | אבא |  |
| 4:2.96 | בבריאת |  |
| 4:2.97 | אדם | . |
| 4:2.98 | ואמרה |  |
| 4:2.99 | לו |  |
| 4:2.100 | אמא | , |
| 4:2.101 | מאי |  |
| 4:2.102 | נפקא |  |
| 4:2.103 | לך |  |
| 4:2.104 | מניה | , |
| 4:2.105 | אם |  |
| 4:2.106 | יחטא |  |
| 4:2.107 | לי |  |
| 4:2.108 | יחטא |  |
| 4:2.109 | ולא |  |
| 4:2.110 | לך |  |
| 4:2.111 | שנ | ' |
| 4:2.112 | ובן |  |
| 4:2.113 | כסיל |  |
| 4:2.114 | תוגת |  |
| 4:2.115 | אמו | , |
| 4:2.116 | וכשחטא |  |
| 4:2.117 | אדם |  |
| 4:2.118 | בעץ |  |
| 4:2.119 | הדעת |  |
| 4:2.120 | גירש |  |
| 4:2.121 | אותו |  |
| 4:2.122 | ואותה |  |
| 4:2.123 | שנאמר |  |
| 4:2.124 | ובפשעיכם |  |
| 4:2.125 | שולחה |  |
| 4:2.126 | אמכם | . |


## Milchamot ha-Shem 4:3

Hebrew (Wikisource): ובעץ חיים להרח"ו שער הנקודים פ"ו כ' שכאשר נאצלו שבע תחתונות עדין לא היה אדם הראשון התחתון נברא בעולם ויצאו זעיר ונוקביה אחור באחור מפני פחד החיצונים שלא יינקו, כי אם היו עומדים פנים בפנים היה מקום לקליפות להתאחז במקום אחיזתן וכאשר נברא אדם הראשון ועשה מצות מעשיות החזירן פנים בפנים על ידי מצות ומעשים טובים וכו' עיין שם שכל דבריו הם היפך דברי הזהר פה, שאדרבה אדם הראשון קלקל בחטאו ונתגרשה אימא בעבורו, ומזה תבין שכל אחד מהקובלים בודה לו דברים מלבו כחפצו.

| id | word | punctuation after |
|---|---|---|
| 4:3.1 | ובעץ |  |
| 4:3.2 | חיים |  |
| 4:3.3 | להרח"ו |  |
| 4:3.4 | שער |  |
| 4:3.5 | הנקודים |  |
| 4:3.6 | פ"ו |  |
| 4:3.7 | כ | ' |
| 4:3.8 | שכאשר |  |
| 4:3.9 | נאצלו |  |
| 4:3.10 | שבע |  |
| 4:3.11 | תחתונות |  |
| 4:3.12 | עדין |  |
| 4:3.13 | לא |  |
| 4:3.14 | היה |  |
| 4:3.15 | אדם |  |
| 4:3.16 | הראשון |  |
| 4:3.17 | התחתון |  |
| 4:3.18 | נברא |  |
| 4:3.19 | בעולם |  |
| 4:3.20 | ויצאו |  |
| 4:3.21 | זעיר |  |
| 4:3.22 | ונוקביה |  |
| 4:3.23 | אחור |  |
| 4:3.24 | באחור |  |
| 4:3.25 | מפני |  |
| 4:3.26 | פחד |  |
| 4:3.27 | החיצונים |  |
| 4:3.28 | שלא |  |
| 4:3.29 | יינקו | , |
| 4:3.30 | כי |  |
| 4:3.31 | אם |  |
| 4:3.32 | היו |  |
| 4:3.33 | עומדים |  |
| 4:3.34 | פנים |  |
| 4:3.35 | בפנים |  |
| 4:3.36 | היה |  |
| 4:3.37 | מקום |  |
| 4:3.38 | לקליפות |  |
| 4:3.39 | להתאחז |  |
| 4:3.40 | במקום |  |
| 4:3.41 | אחיזתן |  |
| 4:3.42 | וכאשר |  |
| 4:3.43 | נברא |  |
| 4:3.44 | אדם |  |
| 4:3.45 | הראשון |  |
| 4:3.46 | ועשה |  |
| 4:3.47 | מצות |  |
| 4:3.48 | מעשיות |  |
| 4:3.49 | החזירן |  |
| 4:3.50 | פנים |  |
| 4:3.51 | בפנים |  |
| 4:3.52 | על |  |
| 4:3.53 | ידי |  |
| 4:3.54 | מצות |  |
| 4:3.55 | ומעשים |  |
| 4:3.56 | טובים |  |
| 4:3.57 | וכו | ' |
| 4:3.58 | עיין |  |
| 4:3.59 | שם |  |
| 4:3.60 | שכל |  |
| 4:3.61 | דבריו |  |
| 4:3.62 | הם |  |
| 4:3.63 | היפך |  |
| 4:3.64 | דברי |  |
| 4:3.65 | הזהר |  |
| 4:3.66 | פה | , |
| 4:3.67 | שאדרבה |  |
| 4:3.68 | אדם |  |
| 4:3.69 | הראשון |  |
| 4:3.70 | קלקל |  |
| 4:3.71 | בחטאו |  |
| 4:3.72 | ונתגרשה |  |
| 4:3.73 | אימא |  |
| 4:3.74 | בעבורו | , |
| 4:3.75 | ומזה |  |
| 4:3.76 | תבין |  |
| 4:3.77 | שכל |  |
| 4:3.78 | אחד |  |
| 4:3.79 | מהקובלים |  |
| 4:3.80 | בודה |  |
| 4:3.81 | לו |  |
| 4:3.82 | דברים |  |
| 4:3.83 | מלבו |  |
| 4:3.84 | כחפצו | . |


## Milchamot ha-Shem 4:4

Hebrew (Wikisource): כא) עוד שם בזהר ד' כ"ב ע"ב פתח עוד ר' שמעון ואמר ראו עתה כי אני אני הוא ואין אלהים עמדי, אמר, חבריא, שמעון מלין עתיקין דבעינא לגלאה בתר דאתיהיב רשו עלאה למימר, מאן ניהו דאמר ראו עתה כי אני אני הוא, אלא דא הוא עלת על כל עלאין ההוא דאתקרי עלת העלות, עלת מאלין עלות דלא יעביד חד מאלין עלות שום עובדא עד דנטיל רשות מההוא דעילוה. כמא דאוקימנא לעילא בנעשה אדם, נעשה ודאי על תרין אתמר, דאמר דא לההוא דלעילא מניה נעשה, ולא עביד מדעם עד דנטיל רשותא ואמירה מההוא דלעילא מניה וההוא דלעילא מניה לא עביד מדעם עד דנטיל עצה מחבריה, אבל ההוא דאתקרי עלת כל עלות דלית לעילא מניה, ולא לתתא שוה ליה, כמא דאתמר ואל מי תדמיוני ואשוה יאמר קדוש, אמר ראו עתה כי אני אני הוא ואין אלהים עמדי דנטיל עצה מניה כגוונא דההוא דאמר ויאמר אלהים נעשה אדם, ועיין עץ חיים שער עתיק פ' שני, ופי' הרב מקדש מלך בשם האר"י בספר הלקוטים, עלת על כל עלאין אדם קדמון נקרא עלת על כל עלות, ההוא דאתקרי עלת העלות, עלת מאלין עלות, פי' כשאנו אומרים עלת העלות שייך בכל פרצוף שנקרא כך לפי שהוא עלה לעלות שתחתיו, אבל כשאנו אומרים עלת כל העלות הוא אדם קדמון, שהוא עלה ראשונה כל הפרצופים כולם ע"ש. הרי נתבאר לך מדברי הפילוסוף מחבר הזהר ומפרשיו מקדש מלך בשם האר"י ודרך אמת ושביבי נגה שהאלוה שנתן רשות לרשב"י לדרוש בדברים שאסור להעלותן על לב לחשוב רבוי באלהות, הוא הקב"ה הנקרא עתיק, ולא זה הוא שאמר לישראל ראו עתה כי אני אני הוא וכו'. כי הקב"ה שאמר ראו עתה כי אני אני הוא, הוא אדם קדמון שהוא עלה לראשונה לכל הפרצופים, ואין לו ממי ליטול רשות. והקב"ה שאמר יהי אור, ויהי רקיע, ויקוו המים וכו' הוא אבא, וזה שאמר נעשה אדם בצלמינו היא אמא שאמרה לאבא נעשה אדם, ולא הסכים הקב"ה שהוא אבא בבריאתו, ואמרה לו אימא מאי איכפת לך, אם יחטא לי יחטא ולא לך שנ' ובן כסיל תוגת אמו ולא תוגת אביו, ואז נברא שלא לרצון אבא כמו שפי' שביבי נגה, שאלו הסכים אבא בבריאתו לא גירשה עמו כשחטא, אלא ודאי לא הסכים אבא לזה כמו שאמר בזהר ולא בעא לאשתתפא וכו' ולכך גירשו מגן עדן וגירשה עמו.

| id | word | punctuation after |
|---|---|---|
| 4:4.1 | כא | ) |
| 4:4.2 | עוד |  |
| 4:4.3 | שם |  |
| 4:4.4 | בזהר |  |
| 4:4.5 | ד | ' |
| 4:4.6 | כ"ב |  |
| 4:4.7 | ע"ב |  |
| 4:4.8 | פתח |  |
| 4:4.9 | עוד |  |
| 4:4.10 | ר | ' |
| 4:4.11 | שמעון |  |
| 4:4.12 | ואמר |  |
| 4:4.13 | ראו |  |
| 4:4.14 | עתה |  |
| 4:4.15 | כי |  |
| 4:4.16 | אני |  |
| 4:4.17 | אני |  |
| 4:4.18 | הוא |  |
| 4:4.19 | ואין |  |
| 4:4.20 | אלהים |  |
| 4:4.21 | עמדי | , |
| 4:4.22 | אמר | , |
| 4:4.23 | חבריא | , |
| 4:4.24 | שמעון |  |
| 4:4.25 | מלין |  |
| 4:4.26 | עתיקין |  |
| 4:4.27 | דבעינא |  |
| 4:4.28 | לגלאה |  |
| 4:4.29 | בתר |  |
| 4:4.30 | דאתיהיב |  |
| 4:4.31 | רשו |  |
| 4:4.32 | עלאה |  |
| 4:4.33 | למימר | , |
| 4:4.34 | מאן |  |
| 4:4.35 | ניהו |  |
| 4:4.36 | דאמר |  |
| 4:4.37 | ראו |  |
| 4:4.38 | עתה |  |
| 4:4.39 | כי |  |
| 4:4.40 | אני |  |
| 4:4.41 | אני |  |
| 4:4.42 | הוא | , |
| 4:4.43 | אלא |  |
| 4:4.44 | דא |  |
| 4:4.45 | הוא |  |
| 4:4.46 | עלת |  |
| 4:4.47 | על |  |
| 4:4.48 | כל |  |
| 4:4.49 | עלאין |  |
| 4:4.50 | ההוא |  |
| 4:4.51 | דאתקרי |  |
| 4:4.52 | עלת |  |
| 4:4.53 | העלות | , |
| 4:4.54 | עלת |  |
| 4:4.55 | מאלין |  |
| 4:4.56 | עלות |  |
| 4:4.57 | דלא |  |
| 4:4.58 | יעביד |  |
| 4:4.59 | חד |  |
| 4:4.60 | מאלין |  |
| 4:4.61 | עלות |  |
| 4:4.62 | שום |  |
| 4:4.63 | עובדא |  |
| 4:4.64 | עד |  |
| 4:4.65 | דנטיל |  |
| 4:4.66 | רשות |  |
| 4:4.67 | מההוא |  |
| 4:4.68 | דעילוה | . |
| 4:4.69 | כמא |  |
| 4:4.70 | דאוקימנא |  |
| 4:4.71 | לעילא |  |
| 4:4.72 | בנעשה |  |
| 4:4.73 | אדם | , |
| 4:4.74 | נעשה |  |
| 4:4.75 | ודאי |  |
| 4:4.76 | על |  |
| 4:4.77 | תרין |  |
| 4:4.78 | אתמר | , |
| 4:4.79 | דאמר |  |
| 4:4.80 | דא |  |
| 4:4.81 | לההוא |  |
| 4:4.82 | דלעילא |  |
| 4:4.83 | מניה |  |
| 4:4.84 | נעשה | , |
| 4:4.85 | ולא |  |
| 4:4.86 | עביד |  |
| 4:4.87 | מדעם |  |
| 4:4.88 | עד |  |
| 4:4.89 | דנטיל |  |
| 4:4.90 | רשותא |  |
| 4:4.91 | ואמירה |  |
| 4:4.92 | מההוא |  |
| 4:4.93 | דלעילא |  |
| 4:4.94 | מניה |  |
| 4:4.95 | וההוא |  |
| 4:4.96 | דלעילא |  |
| 4:4.97 | מניה |  |
| 4:4.98 | לא |  |
| 4:4.99 | עביד |  |
| 4:4.100 | מדעם |  |
| 4:4.101 | עד |  |
| 4:4.102 | דנטיל |  |
| 4:4.103 | עצה |  |
| 4:4.104 | מחבריה | , |
| 4:4.105 | אבל |  |
| 4:4.106 | ההוא |  |
| 4:4.107 | דאתקרי |  |
| 4:4.108 | עלת |  |
| 4:4.109 | כל |  |
| 4:4.110 | עלות |  |
| 4:4.111 | דלית |  |
| 4:4.112 | לעילא |  |
| 4:4.113 | מניה | , |
| 4:4.114 | ולא |  |
| 4:4.115 | לתתא |  |
| 4:4.116 | שוה |  |
| 4:4.117 | ליה | , |
| 4:4.118 | כמא |  |
| 4:4.119 | דאתמר |  |
| 4:4.120 | ואל |  |
| 4:4.121 | מי |  |
| 4:4.122 | תדמיוני |  |
| 4:4.123 | ואשוה |  |
| 4:4.124 | יאמר |  |
| 4:4.125 | קדוש | , |
| 4:4.126 | אמר |  |
| 4:4.127 | ראו |  |
| 4:4.128 | עתה |  |
| 4:4.129 | כי |  |
| 4:4.130 | אני |  |
| 4:4.131 | אני |  |
| 4:4.132 | הוא |  |
| 4:4.133 | ואין |  |
| 4:4.134 | אלהים |  |
| 4:4.135 | עמדי |  |
| 4:4.136 | דנטיל |  |
| 4:4.137 | עצה |  |
| 4:4.138 | מניה |  |
| 4:4.139 | כגוונא |  |
| 4:4.140 | דההוא |  |
| 4:4.141 | דאמר |  |
| 4:4.142 | ויאמר |  |
| 4:4.143 | אלהים |  |
| 4:4.144 | נעשה |  |
| 4:4.145 | אדם | , |
| 4:4.146 | ועיין |  |
| 4:4.147 | עץ |  |
| 4:4.148 | חיים |  |
| 4:4.149 | שער |  |
| 4:4.150 | עתיק |  |
| 4:4.151 | פ | ' |
| 4:4.152 | שני | , |
| 4:4.153 | ופי | ' |
| 4:4.154 | הרב |  |
| 4:4.155 | מקדש |  |
| 4:4.156 | מלך |  |
| 4:4.157 | בשם |  |
| 4:4.158 | האר"י |  |
| 4:4.159 | בספר |  |
| 4:4.160 | הלקוטים | , |
| 4:4.161 | עלת |  |
| 4:4.162 | על |  |
| 4:4.163 | כל |  |
| 4:4.164 | עלאין |  |
| 4:4.165 | אדם |  |
| 4:4.166 | קדמון |  |
| 4:4.167 | נקרא |  |
| 4:4.168 | עלת |  |
| 4:4.169 | על |  |
| 4:4.170 | כל |  |
| 4:4.171 | עלות | , |
| 4:4.172 | ההוא |  |
| 4:4.173 | דאתקרי |  |
| 4:4.174 | עלת |  |
| 4:4.175 | העלות | , |
| 4:4.176 | עלת |  |
| 4:4.177 | מאלין |  |
| 4:4.178 | עלות | , |
| 4:4.179 | פי | ' |
| 4:4.180 | כשאנו |  |
| 4:4.181 | אומרים |  |
| 4:4.182 | עלת |  |
| 4:4.183 | העלות |  |
| 4:4.184 | שייך |  |
| 4:4.185 | בכל |  |
| 4:4.186 | פרצוף |  |
| 4:4.187 | שנקרא |  |
| 4:4.188 | כך |  |
| 4:4.189 | לפי |  |
| 4:4.190 | שהוא |  |
| 4:4.191 | עלה |  |
| 4:4.192 | לעלות |  |
| 4:4.193 | שתחתיו | , |
| 4:4.194 | אבל |  |
| 4:4.195 | כשאנו |  |
| 4:4.196 | אומרים |  |
| 4:4.197 | עלת |  |
| 4:4.198 | כל |  |
| 4:4.199 | העלות |  |
| 4:4.200 | הוא |  |
| 4:4.201 | אדם |  |
| 4:4.202 | קדמון | , |
| 4:4.203 | שהוא |  |
| 4:4.204 | עלה |  |
| 4:4.205 | ראשונה |  |
| 4:4.206 | כל |  |
| 4:4.207 | הפרצופים |  |
| 4:4.208 | כולם |  |
| 4:4.209 | ע"ש | . |
| 4:4.210 | הרי |  |
| 4:4.211 | נתבאר |  |
| 4:4.212 | לך |  |
| 4:4.213 | מדברי |  |
| 4:4.214 | הפילוסוף |  |
| 4:4.215 | מחבר |  |
| 4:4.216 | הזהר |  |
| 4:4.217 | ומפרשיו |  |
| 4:4.218 | מקדש |  |
| 4:4.219 | מלך |  |
| 4:4.220 | בשם |  |
| 4:4.221 | האר"י |  |
| 4:4.222 | ודרך |  |
| 4:4.223 | אמת |  |
| 4:4.224 | ושביבי |  |
| 4:4.225 | נגה |  |
| 4:4.226 | שהאלוה |  |
| 4:4.227 | שנתן |  |
| 4:4.228 | רשות |  |
| 4:4.229 | לרשב"י |  |
| 4:4.230 | לדרוש |  |
| 4:4.231 | בדברים |  |
| 4:4.232 | שאסור |  |
| 4:4.233 | להעלותן |  |
| 4:4.234 | על |  |
| 4:4.235 | לב |  |
| 4:4.236 | לחשוב |  |
| 4:4.237 | רבוי |  |
| 4:4.238 | באלהות | , |
| 4:4.239 | הוא |  |
| 4:4.240 | הקב"ה |  |
| 4:4.241 | הנקרא |  |
| 4:4.242 | עתיק | , |
| 4:4.243 | ולא |  |
| 4:4.244 | זה |  |
| 4:4.245 | הוא |  |
| 4:4.246 | שאמר |  |
| 4:4.247 | לישראל |  |
| 4:4.248 | ראו |  |
| 4:4.249 | עתה |  |
| 4:4.250 | כי |  |
| 4:4.251 | אני |  |
| 4:4.252 | אני |  |
| 4:4.253 | הוא |  |
| 4:4.254 | וכו | '. |
| 4:4.255 | כי |  |
| 4:4.256 | הקב"ה |  |
| 4:4.257 | שאמר |  |
| 4:4.258 | ראו |  |
| 4:4.259 | עתה |  |
| 4:4.260 | כי |  |
| 4:4.261 | אני |  |
| 4:4.262 | אני |  |
| 4:4.263 | הוא | , |
| 4:4.264 | הוא |  |
| 4:4.265 | אדם |  |
| 4:4.266 | קדמון |  |
| 4:4.267 | שהוא |  |
| 4:4.268 | עלה |  |
| 4:4.269 | לראשונה |  |
| 4:4.270 | לכל |  |
| 4:4.271 | הפרצופים | , |
| 4:4.272 | ואין |  |
| 4:4.273 | לו |  |
| 4:4.274 | ממי |  |
| 4:4.275 | ליטול |  |
| 4:4.276 | רשות | . |
| 4:4.277 | והקב"ה |  |
| 4:4.278 | שאמר |  |
| 4:4.279 | יהי |  |
| 4:4.280 | אור | , |
| 4:4.281 | ויהי |  |
| 4:4.282 | רקיע | , |
| 4:4.283 | ויקוו |  |
| 4:4.284 | המים |  |
| 4:4.285 | וכו | ' |
| 4:4.286 | הוא |  |
| 4:4.287 | אבא | , |
| 4:4.288 | וזה |  |
| 4:4.289 | שאמר |  |
| 4:4.290 | נעשה |  |
| 4:4.291 | אדם |  |
| 4:4.292 | בצלמינו |  |
| 4:4.293 | היא |  |
| 4:4.294 | אמא |  |
| 4:4.295 | שאמרה |  |
| 4:4.296 | לאבא |  |
| 4:4.297 | נעשה |  |
| 4:4.298 | אדם | , |
| 4:4.299 | ולא |  |
| 4:4.300 | הסכים |  |
| 4:4.301 | הקב"ה |  |
| 4:4.302 | שהוא |  |
| 4:4.303 | אבא |  |
| 4:4.304 | בבריאתו | , |
| 4:4.305 | ואמרה |  |
| 4:4.306 | לו |  |
| 4:4.307 | אימא |  |
| 4:4.308 | מאי |  |
| 4:4.309 | איכפת |  |
| 4:4.310 | לך | , |
| 4:4.311 | אם |  |
| 4:4.312 | יחטא |  |
| 4:4.313 | לי |  |
| 4:4.314 | יחטא |  |
| 4:4.315 | ולא |  |
| 4:4.316 | לך |  |
| 4:4.317 | שנ | ' |
| 4:4.318 | ובן |  |
| 4:4.319 | כסיל |  |
| 4:4.320 | תוגת |  |
| 4:4.321 | אמו |  |
| 4:4.322 | ולא |  |
| 4:4.323 | תוגת |  |
| 4:4.324 | אביו | , |
| 4:4.325 | ואז |  |
| 4:4.326 | נברא |  |
| 4:4.327 | שלא |  |
| 4:4.328 | לרצון |  |
| 4:4.329 | אבא |  |
| 4:4.330 | כמו |  |
| 4:4.331 | שפי | ' |
| 4:4.332 | שביבי |  |
| 4:4.333 | נגה | , |
| 4:4.334 | שאלו |  |
| 4:4.335 | הסכים |  |
| 4:4.336 | אבא |  |
| 4:4.337 | בבריאתו |  |
| 4:4.338 | לא |  |
| 4:4.339 | גירשה |  |
| 4:4.340 | עמו |  |
| 4:4.341 | כשחטא | , |
| 4:4.342 | אלא |  |
| 4:4.343 | ודאי |  |
| 4:4.344 | לא |  |
| 4:4.345 | הסכים |  |
| 4:4.346 | אבא |  |
| 4:4.347 | לזה |  |
| 4:4.348 | כמו |  |
| 4:4.349 | שאמר |  |
| 4:4.350 | בזהר |  |
| 4:4.351 | ולא |  |
| 4:4.352 | בעא |  |
| 4:4.353 | לאשתתפא |  |
| 4:4.354 | וכו | ' |
| 4:4.355 | ולכך |  |
| 4:4.356 | גירשו |  |
| 4:4.357 | מגן |  |
| 4:4.358 | עדן |  |
| 4:4.359 | וגירשה |  |
| 4:4.360 | עמו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

