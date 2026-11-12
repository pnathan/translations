# Translation pack: milhamot-hashem perek-5-21 (Milchamot ha-Shem 5:30–5:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-21`.

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
- **ב"ה** → "blessed be He".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
## Terms

- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- כופר **בעיקר** → "denies the fundamental principle".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ת"ל** → "Scripture says".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **ווי** ליה מאן דלא → "woe to him who does not".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 5:30

Hebrew (Wikisource): נו) וז"ל הרמב"ם בפרק ב' מהל' ע"ז עיקר הציווי בע"ש שלא לעבוד אחד מכל הנבראים הנבראים מהם. ואעפ"י שהעובד יודע שה' הוא האלהים והוא עובד הנברא הזה על דרך שעבד אנוש ואנשי דורו תחלה הרי זה עובד ע"ז. ובריש פרק א' כ' בימי אנוש טעו בני אדם טעות גדולה וכו' אמרו הואיל והאל ברא כוכבים וגלגלים אלו להנהיג את העולם ונתנן במרום וחלק להם כבוד. והן שמשים שמשמשים לפניו. ראוים הם לשבחם ולפארם ולחלוק להם כבוד. וזה הוא רצון האל ב"ה לגדל ולכבד מי שגדלו וכבדו. כמו שהמלך רוצה לכבד עבדיו והעומדים לפניו וזהו כבודו של מלך. עכ"ל ובילקוט פרשת יתרו לא יהיה לך אלהים אחרים ר' יוסי אומר אלהים אחרים למה נאמר. שלא ליתן פתחון פה לאומות העולם לומר אלו נקראו בשמו כבר היה בהן צורך. הרי נקראו בשמו ואין בהן צורך. אימתי נקראו בשמו בימי אנוש בן שת שנ' אז הוחל לקרא בשם ה'. החל לקרוא ע"ז בשם המקום באותה שעה עלה אוקינוס והציף שלישו של עולם וכו'. ובמדרש הגדול פר' בראשית אנוש, שהיתה דעתו אנושה ונבערה, והתחיל בע"ז בעולם שה' אז הוחל לקרא בשפ ה'. החל לקרוא ע"ז בשם המקום באותה שעה עלה אוקיינוס וכו'. ובמדרש חפץ לר' זכריה הרופא פר' יתרו ובילקוט לא תעשה לך פסל יכול לא תעשה גלומה (ובילקוט גלופה) אבל יעשה אטומה ת"ל כל צמונה. יכול מטע, ת"ל לא תטע לך אשרה. יכול יעשה של עץ ת"ל כל עץ. יכול יעשה של אבן ת"ל ואבן משכית. יכול של זהב ת"ל ואלהי זהב. יכול יעשה של נחשת בדיל ברזל ועופרת ת"ל אלהי מסכה לא תעשה לך. יכול וכו' ומסיים יכול לא יעשה, אבל יעלה בדעתו ת"ל לא יהיה לך אלהים אחרים על פני ע"כ. וכן ב' הרמב"ם ז"ל בפ"א מהל' יסודי התורה וז"ל וכל המעלה בדעתו שיש שם אלוה אחר חוץ מזה עובר בלא תעשה שנ' לא יהיה לך אלהים אחרים על פני וכפר בעיקר שזהו העיקר הגדול שהכל תלוי בו ע"כ. מבואר מכל זה בהדיא שאע"פ שאינו עושה שום מעשה רק שמאמין בדעתו שיש אלוה אחר למטה ממנו מהלאהים העליון ה"ז מין.

| id | word | punctuation after |
|---|---|---|
| 5:30.1 | נו | ) |
| 5:30.2 | וז"ל |  |
| 5:30.3 | הרמב"ם |  |
| 5:30.4 | בפרק |  |
| 5:30.5 | ב | ' |
| 5:30.6 | מהל | ' |
| 5:30.7 | ע"ז |  |
| 5:30.8 | עיקר |  |
| 5:30.9 | הציווי |  |
| 5:30.10 | בע"ש |  |
| 5:30.11 | שלא |  |
| 5:30.12 | לעבוד |  |
| 5:30.13 | אחד |  |
| 5:30.14 | מכל |  |
| 5:30.15 | הנבראים |  |
| 5:30.16 | הנבראים |  |
| 5:30.17 | מהם | . |
| 5:30.18 | ואעפ"י |  |
| 5:30.19 | שהעובד |  |
| 5:30.20 | יודע |  |
| 5:30.21 | שה | ' |
| 5:30.22 | הוא |  |
| 5:30.23 | האלהים |  |
| 5:30.24 | והוא |  |
| 5:30.25 | עובד |  |
| 5:30.26 | הנברא |  |
| 5:30.27 | הזה |  |
| 5:30.28 | על |  |
| 5:30.29 | דרך |  |
| 5:30.30 | שעבד |  |
| 5:30.31 | אנוש |  |
| 5:30.32 | ואנשי |  |
| 5:30.33 | דורו |  |
| 5:30.34 | תחלה |  |
| 5:30.35 | הרי |  |
| 5:30.36 | זה |  |
| 5:30.37 | עובד |  |
| 5:30.38 | ע"ז | . |
| 5:30.39 | ובריש |  |
| 5:30.40 | פרק |  |
| 5:30.41 | א | ' |
| 5:30.42 | כ | ' |
| 5:30.43 | בימי |  |
| 5:30.44 | אנוש |  |
| 5:30.45 | טעו |  |
| 5:30.46 | בני |  |
| 5:30.47 | אדם |  |
| 5:30.48 | טעות |  |
| 5:30.49 | גדולה |  |
| 5:30.50 | וכו | ' |
| 5:30.51 | אמרו |  |
| 5:30.52 | הואיל |  |
| 5:30.53 | והאל |  |
| 5:30.54 | ברא |  |
| 5:30.55 | כוכבים |  |
| 5:30.56 | וגלגלים |  |
| 5:30.57 | אלו |  |
| 5:30.58 | להנהיג |  |
| 5:30.59 | את |  |
| 5:30.60 | העולם |  |
| 5:30.61 | ונתנן |  |
| 5:30.62 | במרום |  |
| 5:30.63 | וחלק |  |
| 5:30.64 | להם |  |
| 5:30.65 | כבוד | . |
| 5:30.66 | והן |  |
| 5:30.67 | שמשים |  |
| 5:30.68 | שמשמשים |  |
| 5:30.69 | לפניו | . |
| 5:30.70 | ראוים |  |
| 5:30.71 | הם |  |
| 5:30.72 | לשבחם |  |
| 5:30.73 | ולפארם |  |
| 5:30.74 | ולחלוק |  |
| 5:30.75 | להם |  |
| 5:30.76 | כבוד | . |
| 5:30.77 | וזה |  |
| 5:30.78 | הוא |  |
| 5:30.79 | רצון |  |
| 5:30.80 | האל |  |
| 5:30.81 | ב"ה |  |
| 5:30.82 | לגדל |  |
| 5:30.83 | ולכבד |  |
| 5:30.84 | מי |  |
| 5:30.85 | שגדלו |  |
| 5:30.86 | וכבדו | . |
| 5:30.87 | כמו |  |
| 5:30.88 | שהמלך |  |
| 5:30.89 | רוצה |  |
| 5:30.90 | לכבד |  |
| 5:30.91 | עבדיו |  |
| 5:30.92 | והעומדים |  |
| 5:30.93 | לפניו |  |
| 5:30.94 | וזהו |  |
| 5:30.95 | כבודו |  |
| 5:30.96 | של |  |
| 5:30.97 | מלך | . |
| 5:30.98 | עכ"ל |  |
| 5:30.99 | ובילקוט |  |
| 5:30.100 | פרשת |  |
| 5:30.101 | יתרו |  |
| 5:30.102 | לא |  |
| 5:30.103 | יהיה |  |
| 5:30.104 | לך |  |
| 5:30.105 | אלהים |  |
| 5:30.106 | אחרים |  |
| 5:30.107 | ר | ' |
| 5:30.108 | יוסי |  |
| 5:30.109 | אומר |  |
| 5:30.110 | אלהים |  |
| 5:30.111 | אחרים |  |
| 5:30.112 | למה |  |
| 5:30.113 | נאמר | . |
| 5:30.114 | שלא |  |
| 5:30.115 | ליתן |  |
| 5:30.116 | פתחון |  |
| 5:30.117 | פה |  |
| 5:30.118 | לאומות |  |
| 5:30.119 | העולם |  |
| 5:30.120 | לומר |  |
| 5:30.121 | אלו |  |
| 5:30.122 | נקראו |  |
| 5:30.123 | בשמו |  |
| 5:30.124 | כבר |  |
| 5:30.125 | היה |  |
| 5:30.126 | בהן |  |
| 5:30.127 | צורך | . |
| 5:30.128 | הרי |  |
| 5:30.129 | נקראו |  |
| 5:30.130 | בשמו |  |
| 5:30.131 | ואין |  |
| 5:30.132 | בהן |  |
| 5:30.133 | צורך | . |
| 5:30.134 | אימתי |  |
| 5:30.135 | נקראו |  |
| 5:30.136 | בשמו |  |
| 5:30.137 | בימי |  |
| 5:30.138 | אנוש |  |
| 5:30.139 | בן |  |
| 5:30.140 | שת |  |
| 5:30.141 | שנ | ' |
| 5:30.142 | אז |  |
| 5:30.143 | הוחל |  |
| 5:30.144 | לקרא |  |
| 5:30.145 | בשם |  |
| 5:30.146 | ה | '. |
| 5:30.147 | החל |  |
| 5:30.148 | לקרוא |  |
| 5:30.149 | ע"ז |  |
| 5:30.150 | בשם |  |
| 5:30.151 | המקום |  |
| 5:30.152 | באותה |  |
| 5:30.153 | שעה |  |
| 5:30.154 | עלה |  |
| 5:30.155 | אוקינוס |  |
| 5:30.156 | והציף |  |
| 5:30.157 | שלישו |  |
| 5:30.158 | של |  |
| 5:30.159 | עולם |  |
| 5:30.160 | וכו | '. |
| 5:30.161 | ובמדרש |  |
| 5:30.162 | הגדול |  |
| 5:30.163 | פר | ' |
| 5:30.164 | בראשית |  |
| 5:30.165 | אנוש | , |
| 5:30.166 | שהיתה |  |
| 5:30.167 | דעתו |  |
| 5:30.168 | אנושה |  |
| 5:30.169 | ונבערה | , |
| 5:30.170 | והתחיל |  |
| 5:30.171 | בע"ז |  |
| 5:30.172 | בעולם |  |
| 5:30.173 | שה | ' |
| 5:30.174 | אז |  |
| 5:30.175 | הוחל |  |
| 5:30.176 | לקרא |  |
| 5:30.177 | בשפ |  |
| 5:30.178 | ה | '. |
| 5:30.179 | החל |  |
| 5:30.180 | לקרוא |  |
| 5:30.181 | ע"ז |  |
| 5:30.182 | בשם |  |
| 5:30.183 | המקום |  |
| 5:30.184 | באותה |  |
| 5:30.185 | שעה |  |
| 5:30.186 | עלה |  |
| 5:30.187 | אוקיינוס |  |
| 5:30.188 | וכו | '. |
| 5:30.189 | ובמדרש |  |
| 5:30.190 | חפץ |  |
| 5:30.191 | לר | ' |
| 5:30.192 | זכריה |  |
| 5:30.193 | הרופא |  |
| 5:30.194 | פר | ' |
| 5:30.195 | יתרו |  |
| 5:30.196 | ובילקוט |  |
| 5:30.197 | לא |  |
| 5:30.198 | תעשה |  |
| 5:30.199 | לך |  |
| 5:30.200 | פסל |  |
| 5:30.201 | יכול |  |
| 5:30.202 | לא |  |
| 5:30.203 | תעשה |  |
| 5:30.204 | גלומה |  |
| 5:30.205 | ובילקוט |  |
| 5:30.206 | גלופה | ) |
| 5:30.207 | אבל |  |
| 5:30.208 | יעשה |  |
| 5:30.209 | אטומה |  |
| 5:30.210 | ת"ל |  |
| 5:30.211 | כל |  |
| 5:30.212 | צמונה | . |
| 5:30.213 | יכול |  |
| 5:30.214 | מטע | , |
| 5:30.215 | ת"ל |  |
| 5:30.216 | לא |  |
| 5:30.217 | תטע |  |
| 5:30.218 | לך |  |
| 5:30.219 | אשרה | . |
| 5:30.220 | יכול |  |
| 5:30.221 | יעשה |  |
| 5:30.222 | של |  |
| 5:30.223 | עץ |  |
| 5:30.224 | ת"ל |  |
| 5:30.225 | כל |  |
| 5:30.226 | עץ | . |
| 5:30.227 | יכול |  |
| 5:30.228 | יעשה |  |
| 5:30.229 | של |  |
| 5:30.230 | אבן |  |
| 5:30.231 | ת"ל |  |
| 5:30.232 | ואבן |  |
| 5:30.233 | משכית | . |
| 5:30.234 | יכול |  |
| 5:30.235 | של |  |
| 5:30.236 | זהב |  |
| 5:30.237 | ת"ל |  |
| 5:30.238 | ואלהי |  |
| 5:30.239 | זהב | . |
| 5:30.240 | יכול |  |
| 5:30.241 | יעשה |  |
| 5:30.242 | של |  |
| 5:30.243 | נחשת |  |
| 5:30.244 | בדיל |  |
| 5:30.245 | ברזל |  |
| 5:30.246 | ועופרת |  |
| 5:30.247 | ת"ל |  |
| 5:30.248 | אלהי |  |
| 5:30.249 | מסכה |  |
| 5:30.250 | לא |  |
| 5:30.251 | תעשה |  |
| 5:30.252 | לך | . |
| 5:30.253 | יכול |  |
| 5:30.254 | וכו | ' |
| 5:30.255 | ומסיים |  |
| 5:30.256 | יכול |  |
| 5:30.257 | לא |  |
| 5:30.258 | יעשה | , |
| 5:30.259 | אבל |  |
| 5:30.260 | יעלה |  |
| 5:30.261 | בדעתו |  |
| 5:30.262 | ת"ל |  |
| 5:30.263 | לא |  |
| 5:30.264 | יהיה |  |
| 5:30.265 | לך |  |
| 5:30.266 | אלהים |  |
| 5:30.267 | אחרים |  |
| 5:30.268 | על |  |
| 5:30.269 | פני |  |
| 5:30.270 | ע"כ | . |
| 5:30.271 | וכן |  |
| 5:30.272 | ב | ' |
| 5:30.273 | הרמב"ם |  |
| 5:30.274 | ז"ל |  |
| 5:30.275 | בפ"א |  |
| 5:30.276 | מהל | ' |
| 5:30.277 | יסודי |  |
| 5:30.278 | התורה |  |
| 5:30.279 | וז"ל |  |
| 5:30.280 | וכל |  |
| 5:30.281 | המעלה |  |
| 5:30.282 | בדעתו |  |
| 5:30.283 | שיש |  |
| 5:30.284 | שם |  |
| 5:30.285 | אלוה |  |
| 5:30.286 | אחר |  |
| 5:30.287 | חוץ |  |
| 5:30.288 | מזה |  |
| 5:30.289 | עובר |  |
| 5:30.290 | בלא |  |
| 5:30.291 | תעשה |  |
| 5:30.292 | שנ | ' |
| 5:30.293 | לא |  |
| 5:30.294 | יהיה |  |
| 5:30.295 | לך |  |
| 5:30.296 | אלהים |  |
| 5:30.297 | אחרים |  |
| 5:30.298 | על |  |
| 5:30.299 | פני |  |
| 5:30.300 | וכפר |  |
| 5:30.301 | בעיקר |  |
| 5:30.302 | שזהו |  |
| 5:30.303 | העיקר |  |
| 5:30.304 | הגדול |  |
| 5:30.305 | שהכל |  |
| 5:30.306 | תלוי |  |
| 5:30.307 | בו |  |
| 5:30.308 | ע"כ | . |
| 5:30.309 | מבואר |  |
| 5:30.310 | מכל |  |
| 5:30.311 | זה |  |
| 5:30.312 | בהדיא |  |
| 5:30.313 | שאע"פ |  |
| 5:30.314 | שאינו |  |
| 5:30.315 | עושה |  |
| 5:30.316 | שום |  |
| 5:30.317 | מעשה |  |
| 5:30.318 | רק |  |
| 5:30.319 | שמאמין |  |
| 5:30.320 | בדעתו |  |
| 5:30.321 | שיש |  |
| 5:30.322 | אלוה |  |
| 5:30.323 | אחר |  |
| 5:30.324 | למטה |  |
| 5:30.325 | ממנו |  |
| 5:30.326 | מהלאהים |  |
| 5:30.327 | העליון |  |
| 5:30.328 | ה"ז |  |
| 5:30.329 | מין | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

