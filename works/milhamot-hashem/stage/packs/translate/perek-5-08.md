# Translation pack: milhamot-hashem perek-5-08 (Milchamot ha-Shem 5:14–5:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-08`.

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
- **רבי'** before the lawgiver's name → "Moses our teacher".
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
- **בא"י** → "in the Land of Israel".

## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- אל **אחר** → "another god".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **ריבה / לרבות** → "include", never "multiply".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- **תא חזי** → "come and see".
- מי **אתי** (ketiv): translated, the qere noted.
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


## Milchamot ha-Shem 5:14

Hebrew (Wikisource): מו) ועל יסוד רעוע זה בנה מחבר הזהר מעשה שבדה מלבו ואמר: כשחלה רשב"י וקרבו ימיו למות, שבחר לדון דינו לפני המלך העליון שהוא עתיקא קדישא, ולא לפני זעיר אנפין הנקרא בית דין כמבואר בזהר חדש, ומובא בסוף הזהר בתוספות המובאים שם וז"ל: כשחלה רשב"י עאלו קמיה רבי פנחס ור' חייא ור' אבא. אמרו ליה מאן קיימא דעלמא שכיב. אמר לון לאו בי דינא דלעילא (זעיר אנפין) מעיינין בדינאי דהא אנא חזי דלית אנא מתיהיב למלכא ולדיינא דלעילא. דאנא לאו כשאר אינשי. אלא האי דינא דילי קב"ה דייניה ולאו בדיניה. והיינו דאמר דוד שפטני אלהים וריבה ריבי. וגם שלמה אמר לעשות משפט עבדו. הוא בלחודוי ולא אחרא. דהא תנן כשאדם שכיב, בי דינא דלעילא מסתכלין בדינוי, אית מנהון נוטין לכף זכות דאחזיין זכותיה דבר נש, ואית מנהון נוטין לכף חובה דאחזיין חוביה דבר נש. ולא אפיק אנש מדיניה כמה דהוה בעי. אבל מאן דדאין קמי מלכא עלאה דשליט על כלא הוא טב. ולא יכיל איניש למהוי בההוא דינא בר טב. מאי טעמא דהא תנן כל מכילוי דמלכא עלאה נטו לזכותא תדיר והוא כוליה צד דמהימנותא. ובידוי לשבקא חובין וחטאין הה"ד כי עמך הסליחה למען תיורא, עמך הסליחה ולא עם אחר, ובגין כך בעינא קמיה דהוא דאין דינאי. ואנא איעול תליסר בני לעלמא דאתי וכו'. הרי בהדיא שסובר שרשב"י בחר לדון קמי מלכא עלאה שהוא עתיק הנקרא אצלם רעוא דרעוין שהוא ותרן וסלחן וידון לו לאדם כחפצו מפני שאינו מקפיד על מעשה התחתונים בין טוב ובין רע. כיעקב כן עשו אצלו. ואין אדם נידון ויוצא מלפניו אלא זכאי. מה שאין כן זעיר אנפין הנקרא ב"ד אינו דן את הנידון לפניו כחפצו של נידון רק כפי ראות הדיין מאן דזכי לדינא לדינא, ומאן דזכי לרחמי רחמי, ואין לך אמונת הרבה רשויות בשמים יותר מזה.

| id | word | punctuation after |
|---|---|---|
| 5:14.1 | מו | ) |
| 5:14.2 | ועל |  |
| 5:14.3 | יסוד |  |
| 5:14.4 | רעוע |  |
| 5:14.5 | זה |  |
| 5:14.6 | בנה |  |
| 5:14.7 | מחבר |  |
| 5:14.8 | הזהר |  |
| 5:14.9 | מעשה |  |
| 5:14.10 | שבדה |  |
| 5:14.11 | מלבו |  |
| 5:14.12 | ואמר | : |
| 5:14.13 | כשחלה |  |
| 5:14.14 | רשב"י |  |
| 5:14.15 | וקרבו |  |
| 5:14.16 | ימיו |  |
| 5:14.17 | למות | , |
| 5:14.18 | שבחר |  |
| 5:14.19 | לדון |  |
| 5:14.20 | דינו |  |
| 5:14.21 | לפני |  |
| 5:14.22 | המלך |  |
| 5:14.23 | העליון |  |
| 5:14.24 | שהוא |  |
| 5:14.25 | עתיקא |  |
| 5:14.26 | קדישא | , |
| 5:14.27 | ולא |  |
| 5:14.28 | לפני |  |
| 5:14.29 | זעיר |  |
| 5:14.30 | אנפין |  |
| 5:14.31 | הנקרא |  |
| 5:14.32 | בית |  |
| 5:14.33 | דין |  |
| 5:14.34 | כמבואר |  |
| 5:14.35 | בזהר |  |
| 5:14.36 | חדש | , |
| 5:14.37 | ומובא |  |
| 5:14.38 | בסוף |  |
| 5:14.39 | הזהר |  |
| 5:14.40 | בתוספות |  |
| 5:14.41 | המובאים |  |
| 5:14.42 | שם |  |
| 5:14.43 | וז"ל | : |
| 5:14.44 | כשחלה |  |
| 5:14.45 | רשב"י |  |
| 5:14.46 | עאלו |  |
| 5:14.47 | קמיה |  |
| 5:14.48 | רבי |  |
| 5:14.49 | פנחס |  |
| 5:14.50 | ור | ' |
| 5:14.51 | חייא |  |
| 5:14.52 | ור | ' |
| 5:14.53 | אבא | . |
| 5:14.54 | אמרו |  |
| 5:14.55 | ליה |  |
| 5:14.56 | מאן |  |
| 5:14.57 | קיימא |  |
| 5:14.58 | דעלמא |  |
| 5:14.59 | שכיב | . |
| 5:14.60 | אמר |  |
| 5:14.61 | לון |  |
| 5:14.62 | לאו |  |
| 5:14.63 | בי |  |
| 5:14.64 | דינא |  |
| 5:14.65 | דלעילא |  |
| 5:14.66 | זעיר |  |
| 5:14.67 | אנפין | ) |
| 5:14.68 | מעיינין |  |
| 5:14.69 | בדינאי |  |
| 5:14.70 | דהא |  |
| 5:14.71 | אנא |  |
| 5:14.72 | חזי |  |
| 5:14.73 | דלית |  |
| 5:14.74 | אנא |  |
| 5:14.75 | מתיהיב |  |
| 5:14.76 | למלכא |  |
| 5:14.77 | ולדיינא |  |
| 5:14.78 | דלעילא | . |
| 5:14.79 | דאנא |  |
| 5:14.80 | לאו |  |
| 5:14.81 | כשאר |  |
| 5:14.82 | אינשי | . |
| 5:14.83 | אלא |  |
| 5:14.84 | האי |  |
| 5:14.85 | דינא |  |
| 5:14.86 | דילי |  |
| 5:14.87 | קב"ה |  |
| 5:14.88 | דייניה |  |
| 5:14.89 | ולאו |  |
| 5:14.90 | בדיניה | . |
| 5:14.91 | והיינו |  |
| 5:14.92 | דאמר |  |
| 5:14.93 | דוד |  |
| 5:14.94 | שפטני |  |
| 5:14.95 | אלהים |  |
| 5:14.96 | וריבה |  |
| 5:14.97 | ריבי | . |
| 5:14.98 | וגם |  |
| 5:14.99 | שלמה |  |
| 5:14.100 | אמר |  |
| 5:14.101 | לעשות |  |
| 5:14.102 | משפט |  |
| 5:14.103 | עבדו | . |
| 5:14.104 | הוא |  |
| 5:14.105 | בלחודוי |  |
| 5:14.106 | ולא |  |
| 5:14.107 | אחרא | . |
| 5:14.108 | דהא |  |
| 5:14.109 | תנן |  |
| 5:14.110 | כשאדם |  |
| 5:14.111 | שכיב | , |
| 5:14.112 | בי |  |
| 5:14.113 | דינא |  |
| 5:14.114 | דלעילא |  |
| 5:14.115 | מסתכלין |  |
| 5:14.116 | בדינוי | , |
| 5:14.117 | אית |  |
| 5:14.118 | מנהון |  |
| 5:14.119 | נוטין |  |
| 5:14.120 | לכף |  |
| 5:14.121 | זכות |  |
| 5:14.122 | דאחזיין |  |
| 5:14.123 | זכותיה |  |
| 5:14.124 | דבר |  |
| 5:14.125 | נש | , |
| 5:14.126 | ואית |  |
| 5:14.127 | מנהון |  |
| 5:14.128 | נוטין |  |
| 5:14.129 | לכף |  |
| 5:14.130 | חובה |  |
| 5:14.131 | דאחזיין |  |
| 5:14.132 | חוביה |  |
| 5:14.133 | דבר |  |
| 5:14.134 | נש | . |
| 5:14.135 | ולא |  |
| 5:14.136 | אפיק |  |
| 5:14.137 | אנש |  |
| 5:14.138 | מדיניה |  |
| 5:14.139 | כמה |  |
| 5:14.140 | דהוה |  |
| 5:14.141 | בעי | . |
| 5:14.142 | אבל |  |
| 5:14.143 | מאן |  |
| 5:14.144 | דדאין |  |
| 5:14.145 | קמי |  |
| 5:14.146 | מלכא |  |
| 5:14.147 | עלאה |  |
| 5:14.148 | דשליט |  |
| 5:14.149 | על |  |
| 5:14.150 | כלא |  |
| 5:14.151 | הוא |  |
| 5:14.152 | טב | . |
| 5:14.153 | ולא |  |
| 5:14.154 | יכיל |  |
| 5:14.155 | איניש |  |
| 5:14.156 | למהוי |  |
| 5:14.157 | בההוא |  |
| 5:14.158 | דינא |  |
| 5:14.159 | בר |  |
| 5:14.160 | טב | . |
| 5:14.161 | מאי |  |
| 5:14.162 | טעמא |  |
| 5:14.163 | דהא |  |
| 5:14.164 | תנן |  |
| 5:14.165 | כל |  |
| 5:14.166 | מכילוי |  |
| 5:14.167 | דמלכא |  |
| 5:14.168 | עלאה |  |
| 5:14.169 | נטו |  |
| 5:14.170 | לזכותא |  |
| 5:14.171 | תדיר |  |
| 5:14.172 | והוא |  |
| 5:14.173 | כוליה |  |
| 5:14.174 | צד |  |
| 5:14.175 | דמהימנותא | . |
| 5:14.176 | ובידוי |  |
| 5:14.177 | לשבקא |  |
| 5:14.178 | חובין |  |
| 5:14.179 | וחטאין |  |
| 5:14.180 | הה"ד |  |
| 5:14.181 | כי |  |
| 5:14.182 | עמך |  |
| 5:14.183 | הסליחה |  |
| 5:14.184 | למען |  |
| 5:14.185 | תיורא | , |
| 5:14.186 | עמך |  |
| 5:14.187 | הסליחה |  |
| 5:14.188 | ולא |  |
| 5:14.189 | עם |  |
| 5:14.190 | אחר | , |
| 5:14.191 | ובגין |  |
| 5:14.192 | כך |  |
| 5:14.193 | בעינא |  |
| 5:14.194 | קמיה |  |
| 5:14.195 | דהוא |  |
| 5:14.196 | דאין |  |
| 5:14.197 | דינאי | . |
| 5:14.198 | ואנא |  |
| 5:14.199 | איעול |  |
| 5:14.200 | תליסר |  |
| 5:14.201 | בני |  |
| 5:14.202 | לעלמא |  |
| 5:14.203 | דאתי |  |
| 5:14.204 | וכו | '. |
| 5:14.205 | הרי |  |
| 5:14.206 | בהדיא |  |
| 5:14.207 | שסובר |  |
| 5:14.208 | שרשב"י |  |
| 5:14.209 | בחר |  |
| 5:14.210 | לדון |  |
| 5:14.211 | קמי |  |
| 5:14.212 | מלכא |  |
| 5:14.213 | עלאה |  |
| 5:14.214 | שהוא |  |
| 5:14.215 | עתיק |  |
| 5:14.216 | הנקרא |  |
| 5:14.217 | אצלם |  |
| 5:14.218 | רעוא |  |
| 5:14.219 | דרעוין |  |
| 5:14.220 | שהוא |  |
| 5:14.221 | ותרן |  |
| 5:14.222 | וסלחן |  |
| 5:14.223 | וידון |  |
| 5:14.224 | לו |  |
| 5:14.225 | לאדם |  |
| 5:14.226 | כחפצו |  |
| 5:14.227 | מפני |  |
| 5:14.228 | שאינו |  |
| 5:14.229 | מקפיד |  |
| 5:14.230 | על |  |
| 5:14.231 | מעשה |  |
| 5:14.232 | התחתונים |  |
| 5:14.233 | בין |  |
| 5:14.234 | טוב |  |
| 5:14.235 | ובין |  |
| 5:14.236 | רע | . |
| 5:14.237 | כיעקב |  |
| 5:14.238 | כן |  |
| 5:14.239 | עשו |  |
| 5:14.240 | אצלו | . |
| 5:14.241 | ואין |  |
| 5:14.242 | אדם |  |
| 5:14.243 | נידון |  |
| 5:14.244 | ויוצא |  |
| 5:14.245 | מלפניו |  |
| 5:14.246 | אלא |  |
| 5:14.247 | זכאי | . |
| 5:14.248 | מה |  |
| 5:14.249 | שאין |  |
| 5:14.250 | כן |  |
| 5:14.251 | זעיר |  |
| 5:14.252 | אנפין |  |
| 5:14.253 | הנקרא |  |
| 5:14.254 | ב"ד |  |
| 5:14.255 | אינו |  |
| 5:14.256 | דן |  |
| 5:14.257 | את |  |
| 5:14.258 | הנידון |  |
| 5:14.259 | לפניו |  |
| 5:14.260 | כחפצו |  |
| 5:14.261 | של |  |
| 5:14.262 | נידון |  |
| 5:14.263 | רק |  |
| 5:14.264 | כפי |  |
| 5:14.265 | ראות |  |
| 5:14.266 | הדיין |  |
| 5:14.267 | מאן |  |
| 5:14.268 | דזכי |  |
| 5:14.269 | לדינא |  |
| 5:14.270 | לדינא | , |
| 5:14.271 | ומאן |  |
| 5:14.272 | דזכי |  |
| 5:14.273 | לרחמי |  |
| 5:14.274 | רחמי | , |
| 5:14.275 | ואין |  |
| 5:14.276 | לך |  |
| 5:14.277 | אמונת |  |
| 5:14.278 | הרבה |  |
| 5:14.279 | רשויות |  |
| 5:14.280 | בשמים |  |
| 5:14.281 | יותר |  |
| 5:14.282 | מזה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

