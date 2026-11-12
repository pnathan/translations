# Translation pack: milhamot-hashem perek-8-22 (Milchamot ha-Shem 8:26–8:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-22`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
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
- **האלהי** before a name → "the divine" (note); never dropped.
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רוקח** → "Roke'ach" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
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
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **רחמנא ליצלן** → "the Merciful One save us".
- קורא **נעים / נעלם** → "pleasant reader".
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


## Milchamot ha-Shem 8:26

Hebrew (Wikisource): ואמר השי"ת וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד, ואמר על ידי הנביא ירמיה, אם יסתר איש במסתרים ואני לא אראנו נאם ה' הלא את השמים ואת הארץ אני מלא נאם ה' ונאמר מלא כל הארץ כבודו, ורבותינו אנשי כנה"ג בתיקוני התפלה אמרו כבודו והודו מלא העולם וכו'! האים על פי הדברים האלה נוכל לומר כיון שהוא ית' מלא הכוכבים ומסלותיהם, והשמש והירח בכל חלקיהם בפנים ובחוץ ובארץ ובימים ובנהרות ובגיאות ובמערות ובחוחים בפנים בפנים ובחוץ! הנוכל לומר על שום כוכב מהם הוא אלהינו, או על השמש והירח והארץ אשר הוא מלא אותם המה אלהינו ונצרפם יחד עם כבודו ית' השוכן בהם? לא! ולא! ואיך נומר על אבא ואימא קודשא בריך הוא ושכינתה דלעילא, ועל זעיר אנפין ונוקביה קודשא בריך הוא ושכינתיה דלתתא, ונאמר הוא ה' אלהינו והוא המלך המהולל בכל התושבחות, ובעבור מה נקרא להצורות והפרצופים שהמציא המסית מחבר הזהר בשם הוי'ה ואדנות ושאר השמות. יענה המקובל החדש בעבור שהאלהים העליון ממלא את הפרצופים בפנים ובחוץ ובין כל פרצוף ופרצוף! אף אתה אמור לו הלא הקב"ה מכריז על ידי נביאו הלא את השמים ואת הארץ אני מלא נאם ה' ועם כל זה הוא מזהיר אותנו בפירוש בתורה שלא נעבוד שום נברא זולתו ולא נשתף עמו שום דבר אחר לא מן העליונים ולא מן התחתונים וכמאמר רשב"י הנז' לעיל וכמו שנת' לעיל סי' נ"ח וסי' ס"ד.

| id | word | punctuation after |
|---|---|---|
| 8:26.1 | ואמר |  |
| 8:26.2 | השי"ת |  |
| 8:26.3 | וידעת |  |
| 8:26.4 | היום |  |
| 8:26.5 | והשבות |  |
| 8:26.6 | אל |  |
| 8:26.7 | לבבך |  |
| 8:26.8 | כי |  |
| 8:26.9 | ה | ' |
| 8:26.10 | הוא |  |
| 8:26.11 | האלהים |  |
| 8:26.12 | בשמים |  |
| 8:26.13 | ממעל |  |
| 8:26.14 | ועל |  |
| 8:26.15 | הארץ |  |
| 8:26.16 | מתחת |  |
| 8:26.17 | אין |  |
| 8:26.18 | עוד | , |
| 8:26.19 | ואמר |  |
| 8:26.20 | על |  |
| 8:26.21 | ידי |  |
| 8:26.22 | הנביא |  |
| 8:26.23 | ירמיה | , |
| 8:26.24 | אם |  |
| 8:26.25 | יסתר |  |
| 8:26.26 | איש |  |
| 8:26.27 | במסתרים |  |
| 8:26.28 | ואני |  |
| 8:26.29 | לא |  |
| 8:26.30 | אראנו |  |
| 8:26.31 | נאם |  |
| 8:26.32 | ה | ' |
| 8:26.33 | הלא |  |
| 8:26.34 | את |  |
| 8:26.35 | השמים |  |
| 8:26.36 | ואת |  |
| 8:26.37 | הארץ |  |
| 8:26.38 | אני |  |
| 8:26.39 | מלא |  |
| 8:26.40 | נאם |  |
| 8:26.41 | ה | ' |
| 8:26.42 | ונאמר |  |
| 8:26.43 | מלא |  |
| 8:26.44 | כל |  |
| 8:26.45 | הארץ |  |
| 8:26.46 | כבודו | , |
| 8:26.47 | ורבותינו |  |
| 8:26.48 | אנשי |  |
| 8:26.49 | כנה"ג |  |
| 8:26.50 | בתיקוני |  |
| 8:26.51 | התפלה |  |
| 8:26.52 | אמרו |  |
| 8:26.53 | כבודו |  |
| 8:26.54 | והודו |  |
| 8:26.55 | מלא |  |
| 8:26.56 | העולם |  |
| 8:26.57 | וכו | '! |
| 8:26.58 | האים |  |
| 8:26.59 | על |  |
| 8:26.60 | פי |  |
| 8:26.61 | הדברים |  |
| 8:26.62 | האלה |  |
| 8:26.63 | נוכל |  |
| 8:26.64 | לומר |  |
| 8:26.65 | כיון |  |
| 8:26.66 | שהוא |  |
| 8:26.67 | ית | ' |
| 8:26.68 | מלא |  |
| 8:26.69 | הכוכבים |  |
| 8:26.70 | ומסלותיהם | , |
| 8:26.71 | והשמש |  |
| 8:26.72 | והירח |  |
| 8:26.73 | בכל |  |
| 8:26.74 | חלקיהם |  |
| 8:26.75 | בפנים |  |
| 8:26.76 | ובחוץ |  |
| 8:26.77 | ובארץ |  |
| 8:26.78 | ובימים |  |
| 8:26.79 | ובנהרות |  |
| 8:26.80 | ובגיאות |  |
| 8:26.81 | ובמערות |  |
| 8:26.82 | ובחוחים |  |
| 8:26.83 | בפנים |  |
| 8:26.84 | בפנים |  |
| 8:26.85 | ובחוץ | ! |
| 8:26.86 | הנוכל |  |
| 8:26.87 | לומר |  |
| 8:26.88 | על |  |
| 8:26.89 | שום |  |
| 8:26.90 | כוכב |  |
| 8:26.91 | מהם |  |
| 8:26.92 | הוא |  |
| 8:26.93 | אלהינו | , |
| 8:26.94 | או |  |
| 8:26.95 | על |  |
| 8:26.96 | השמש |  |
| 8:26.97 | והירח |  |
| 8:26.98 | והארץ |  |
| 8:26.99 | אשר |  |
| 8:26.100 | הוא |  |
| 8:26.101 | מלא |  |
| 8:26.102 | אותם |  |
| 8:26.103 | המה |  |
| 8:26.104 | אלהינו |  |
| 8:26.105 | ונצרפם |  |
| 8:26.106 | יחד |  |
| 8:26.107 | עם |  |
| 8:26.108 | כבודו |  |
| 8:26.109 | ית | ' |
| 8:26.110 | השוכן |  |
| 8:26.111 | בהם | ? |
| 8:26.112 | לא | ! |
| 8:26.113 | ולא | ! |
| 8:26.114 | ואיך |  |
| 8:26.115 | נומר |  |
| 8:26.116 | על |  |
| 8:26.117 | אבא |  |
| 8:26.118 | ואימא |  |
| 8:26.119 | קודשא |  |
| 8:26.120 | בריך |  |
| 8:26.121 | הוא |  |
| 8:26.122 | ושכינתה |  |
| 8:26.123 | דלעילא | , |
| 8:26.124 | ועל |  |
| 8:26.125 | זעיר |  |
| 8:26.126 | אנפין |  |
| 8:26.127 | ונוקביה |  |
| 8:26.128 | קודשא |  |
| 8:26.129 | בריך |  |
| 8:26.130 | הוא |  |
| 8:26.131 | ושכינתיה |  |
| 8:26.132 | דלתתא | , |
| 8:26.133 | ונאמר |  |
| 8:26.134 | הוא |  |
| 8:26.135 | ה | ' |
| 8:26.136 | אלהינו |  |
| 8:26.137 | והוא |  |
| 8:26.138 | המלך |  |
| 8:26.139 | המהולל |  |
| 8:26.140 | בכל |  |
| 8:26.141 | התושבחות | , |
| 8:26.142 | ובעבור |  |
| 8:26.143 | מה |  |
| 8:26.144 | נקרא |  |
| 8:26.145 | להצורות |  |
| 8:26.146 | והפרצופים |  |
| 8:26.147 | שהמציא |  |
| 8:26.148 | המסית |  |
| 8:26.149 | מחבר |  |
| 8:26.150 | הזהר |  |
| 8:26.151 | בשם |  |
| 8:26.152 | הוי'ה |  |
| 8:26.153 | ואדנות |  |
| 8:26.154 | ושאר |  |
| 8:26.155 | השמות | . |
| 8:26.156 | יענה |  |
| 8:26.157 | המקובל |  |
| 8:26.158 | החדש |  |
| 8:26.159 | בעבור |  |
| 8:26.160 | שהאלהים |  |
| 8:26.161 | העליון |  |
| 8:26.162 | ממלא |  |
| 8:26.163 | את |  |
| 8:26.164 | הפרצופים |  |
| 8:26.165 | בפנים |  |
| 8:26.166 | ובחוץ |  |
| 8:26.167 | ובין |  |
| 8:26.168 | כל |  |
| 8:26.169 | פרצוף |  |
| 8:26.170 | ופרצוף | ! |
| 8:26.171 | אף |  |
| 8:26.172 | אתה |  |
| 8:26.173 | אמור |  |
| 8:26.174 | לו |  |
| 8:26.175 | הלא |  |
| 8:26.176 | הקב"ה |  |
| 8:26.177 | מכריז |  |
| 8:26.178 | על |  |
| 8:26.179 | ידי |  |
| 8:26.180 | נביאו |  |
| 8:26.181 | הלא |  |
| 8:26.182 | את |  |
| 8:26.183 | השמים |  |
| 8:26.184 | ואת |  |
| 8:26.185 | הארץ |  |
| 8:26.186 | אני |  |
| 8:26.187 | מלא |  |
| 8:26.188 | נאם |  |
| 8:26.189 | ה | ' |
| 8:26.190 | ועם |  |
| 8:26.191 | כל |  |
| 8:26.192 | זה |  |
| 8:26.193 | הוא |  |
| 8:26.194 | מזהיר |  |
| 8:26.195 | אותנו |  |
| 8:26.196 | בפירוש |  |
| 8:26.197 | בתורה |  |
| 8:26.198 | שלא |  |
| 8:26.199 | נעבוד |  |
| 8:26.200 | שום |  |
| 8:26.201 | נברא |  |
| 8:26.202 | זולתו |  |
| 8:26.203 | ולא |  |
| 8:26.204 | נשתף |  |
| 8:26.205 | עמו |  |
| 8:26.206 | שום |  |
| 8:26.207 | דבר |  |
| 8:26.208 | אחר |  |
| 8:26.209 | לא |  |
| 8:26.210 | מן |  |
| 8:26.211 | העליונים |  |
| 8:26.212 | ולא |  |
| 8:26.213 | מן |  |
| 8:26.214 | התחתונים |  |
| 8:26.215 | וכמאמר |  |
| 8:26.216 | רשב"י |  |
| 8:26.217 | הנז | ' |
| 8:26.218 | לעיל |  |
| 8:26.219 | וכמו |  |
| 8:26.220 | שנת | ' |
| 8:26.221 | לעיל |  |
| 8:26.222 | סי | ' |
| 8:26.223 | נ"ח |  |
| 8:26.224 | וסי | ' |
| 8:26.225 | ס"ד | . |


## Milchamot ha-Shem 8:27

Hebrew (Wikisource): קיט) ואם תאמר שאני שאר הנבראים התחתונים והשפלים שהם רחוקים מהבורא שאין להם דין אלהות! אבל הקרובים לאלוה העליון שהוא הסבה העליונה יש להם דין אלהות! והלא כבר נאמר האלהי מקרוב אני נאם ה' ולא אלהי מרחוק, יאמר שהוא ית' מושל בעליונים ובתחתונים, וכולם נסקרים בסקירה אחת באין הפרש והבדל ביניהם כלל, ונאמר וכבודי לאחר לא אתן, שלא מסר עולמו לשום נברא לעשות בעולם כחפצו, רק הוא אמר ויהי אור צוה ויעמוד, ולא ישנו את תפקידתם אשר תוה עליהם, וכמו שאמרו רז"ל בפ' קמא דפסחים קרייה רחמנא לנהורא ופקדיה אמצותא דיממא, וקרייה לחשוא ופקדיה אמצותא דליליא, אבל אי לשום אחד מכל הנבראים דין אלהות כלל! ואע"פ שהם מלאים מכבודו, ובין שהם קרובים ובין שהם רחוקים, שהרי און שפל ורחוק כמונו יושבי חלד, ועם כל זה אמר הכתוב מי גוי גדול אשר לו אלהים קרובים אליו כה' אלהינו! שהוא קרוב בכל מיני קריבות כמו שאמרו רז"ל שאדם עומד אחורי העמוד בבית הכנסת ומתפלל בלחש, והקב"ה מאזין. כפה לאזן כדלעיל, וכן אמר נעים זמירות ישראל מי כה' אלהינו המגביהי לשבת, המשפילי לראות בשמים ובארץ, כקרוב כרחוק, הכל גלוי לפניו, והכל עבדיו לעשות רצונו וחפצו ובידו לשנוהן כחפצו ורצונו, ואין להם דין אלהות כלל, כי לעולם לא נתן ולא יתן כבודו לאחרים, שנאמר וכבודי לאחר לא אתן. והמופת הגדול לנו יציאת מצרים וקריעת ים סוף וכל המופתים שנעשו שמה. ולא מחשבותינו מחשבותיו ולא דרכינו דרכיו. כי לא נוכל להשיג מהותו ואיכותו ואופן חנייתו כמו שנתבאר לעיל סי' נ"ב ונ"ג בשם הרוקח.

| id | word | punctuation after |
|---|---|---|
| 8:27.1 | קיט | ) |
| 8:27.2 | ואם |  |
| 8:27.3 | תאמר |  |
| 8:27.4 | שאני |  |
| 8:27.5 | שאר |  |
| 8:27.6 | הנבראים |  |
| 8:27.7 | התחתונים |  |
| 8:27.8 | והשפלים |  |
| 8:27.9 | שהם |  |
| 8:27.10 | רחוקים |  |
| 8:27.11 | מהבורא |  |
| 8:27.12 | שאין |  |
| 8:27.13 | להם |  |
| 8:27.14 | דין |  |
| 8:27.15 | אלהות | ! |
| 8:27.16 | אבל |  |
| 8:27.17 | הקרובים |  |
| 8:27.18 | לאלוה |  |
| 8:27.19 | העליון |  |
| 8:27.20 | שהוא |  |
| 8:27.21 | הסבה |  |
| 8:27.22 | העליונה |  |
| 8:27.23 | יש |  |
| 8:27.24 | להם |  |
| 8:27.25 | דין |  |
| 8:27.26 | אלהות | ! |
| 8:27.27 | והלא |  |
| 8:27.28 | כבר |  |
| 8:27.29 | נאמר |  |
| 8:27.30 | האלהי |  |
| 8:27.31 | מקרוב |  |
| 8:27.32 | אני |  |
| 8:27.33 | נאם |  |
| 8:27.34 | ה | ' |
| 8:27.35 | ולא |  |
| 8:27.36 | אלהי |  |
| 8:27.37 | מרחוק | , |
| 8:27.38 | יאמר |  |
| 8:27.39 | שהוא |  |
| 8:27.40 | ית | ' |
| 8:27.41 | מושל |  |
| 8:27.42 | בעליונים |  |
| 8:27.43 | ובתחתונים | , |
| 8:27.44 | וכולם |  |
| 8:27.45 | נסקרים |  |
| 8:27.46 | בסקירה |  |
| 8:27.47 | אחת |  |
| 8:27.48 | באין |  |
| 8:27.49 | הפרש |  |
| 8:27.50 | והבדל |  |
| 8:27.51 | ביניהם |  |
| 8:27.52 | כלל | , |
| 8:27.53 | ונאמר |  |
| 8:27.54 | וכבודי |  |
| 8:27.55 | לאחר |  |
| 8:27.56 | לא |  |
| 8:27.57 | אתן | , |
| 8:27.58 | שלא |  |
| 8:27.59 | מסר |  |
| 8:27.60 | עולמו |  |
| 8:27.61 | לשום |  |
| 8:27.62 | נברא |  |
| 8:27.63 | לעשות |  |
| 8:27.64 | בעולם |  |
| 8:27.65 | כחפצו | , |
| 8:27.66 | רק |  |
| 8:27.67 | הוא |  |
| 8:27.68 | אמר |  |
| 8:27.69 | ויהי |  |
| 8:27.70 | אור |  |
| 8:27.71 | צוה |  |
| 8:27.72 | ויעמוד | , |
| 8:27.73 | ולא |  |
| 8:27.74 | ישנו |  |
| 8:27.75 | את |  |
| 8:27.76 | תפקידתם |  |
| 8:27.77 | אשר |  |
| 8:27.78 | תוה |  |
| 8:27.79 | עליהם | , |
| 8:27.80 | וכמו |  |
| 8:27.81 | שאמרו |  |
| 8:27.82 | רז"ל |  |
| 8:27.83 | בפ | ' |
| 8:27.84 | קמא |  |
| 8:27.85 | דפסחים |  |
| 8:27.86 | קרייה |  |
| 8:27.87 | רחמנא |  |
| 8:27.88 | לנהורא |  |
| 8:27.89 | ופקדיה |  |
| 8:27.90 | אמצותא |  |
| 8:27.91 | דיממא | , |
| 8:27.92 | וקרייה |  |
| 8:27.93 | לחשוא |  |
| 8:27.94 | ופקדיה |  |
| 8:27.95 | אמצותא |  |
| 8:27.96 | דליליא | , |
| 8:27.97 | אבל |  |
| 8:27.98 | אי |  |
| 8:27.99 | לשום |  |
| 8:27.100 | אחד |  |
| 8:27.101 | מכל |  |
| 8:27.102 | הנבראים |  |
| 8:27.103 | דין |  |
| 8:27.104 | אלהות |  |
| 8:27.105 | כלל | ! |
| 8:27.106 | ואע"פ |  |
| 8:27.107 | שהם |  |
| 8:27.108 | מלאים |  |
| 8:27.109 | מכבודו | , |
| 8:27.110 | ובין |  |
| 8:27.111 | שהם |  |
| 8:27.112 | קרובים |  |
| 8:27.113 | ובין |  |
| 8:27.114 | שהם |  |
| 8:27.115 | רחוקים | , |
| 8:27.116 | שהרי |  |
| 8:27.117 | און |  |
| 8:27.118 | שפל |  |
| 8:27.119 | ורחוק |  |
| 8:27.120 | כמונו |  |
| 8:27.121 | יושבי |  |
| 8:27.122 | חלד | , |
| 8:27.123 | ועם |  |
| 8:27.124 | כל |  |
| 8:27.125 | זה |  |
| 8:27.126 | אמר |  |
| 8:27.127 | הכתוב |  |
| 8:27.128 | מי |  |
| 8:27.129 | גוי |  |
| 8:27.130 | גדול |  |
| 8:27.131 | אשר |  |
| 8:27.132 | לו |  |
| 8:27.133 | אלהים |  |
| 8:27.134 | קרובים |  |
| 8:27.135 | אליו |  |
| 8:27.136 | כה | ' |
| 8:27.137 | אלהינו | ! |
| 8:27.138 | שהוא |  |
| 8:27.139 | קרוב |  |
| 8:27.140 | בכל |  |
| 8:27.141 | מיני |  |
| 8:27.142 | קריבות |  |
| 8:27.143 | כמו |  |
| 8:27.144 | שאמרו |  |
| 8:27.145 | רז"ל |  |
| 8:27.146 | שאדם |  |
| 8:27.147 | עומד |  |
| 8:27.148 | אחורי |  |
| 8:27.149 | העמוד |  |
| 8:27.150 | בבית |  |
| 8:27.151 | הכנסת |  |
| 8:27.152 | ומתפלל |  |
| 8:27.153 | בלחש | , |
| 8:27.154 | והקב"ה |  |
| 8:27.155 | מאזין | . |
| 8:27.156 | כפה |  |
| 8:27.157 | לאזן |  |
| 8:27.158 | כדלעיל | , |
| 8:27.159 | וכן |  |
| 8:27.160 | אמר |  |
| 8:27.161 | נעים |  |
| 8:27.162 | זמירות |  |
| 8:27.163 | ישראל |  |
| 8:27.164 | מי |  |
| 8:27.165 | כה | ' |
| 8:27.166 | אלהינו |  |
| 8:27.167 | המגביהי |  |
| 8:27.168 | לשבת | , |
| 8:27.169 | המשפילי |  |
| 8:27.170 | לראות |  |
| 8:27.171 | בשמים |  |
| 8:27.172 | ובארץ | , |
| 8:27.173 | כקרוב |  |
| 8:27.174 | כרחוק | , |
| 8:27.175 | הכל |  |
| 8:27.176 | גלוי |  |
| 8:27.177 | לפניו | , |
| 8:27.178 | והכל |  |
| 8:27.179 | עבדיו |  |
| 8:27.180 | לעשות |  |
| 8:27.181 | רצונו |  |
| 8:27.182 | וחפצו |  |
| 8:27.183 | ובידו |  |
| 8:27.184 | לשנוהן |  |
| 8:27.185 | כחפצו |  |
| 8:27.186 | ורצונו | , |
| 8:27.187 | ואין |  |
| 8:27.188 | להם |  |
| 8:27.189 | דין |  |
| 8:27.190 | אלהות |  |
| 8:27.191 | כלל | , |
| 8:27.192 | כי |  |
| 8:27.193 | לעולם |  |
| 8:27.194 | לא |  |
| 8:27.195 | נתן |  |
| 8:27.196 | ולא |  |
| 8:27.197 | יתן |  |
| 8:27.198 | כבודו |  |
| 8:27.199 | לאחרים | , |
| 8:27.200 | שנאמר |  |
| 8:27.201 | וכבודי |  |
| 8:27.202 | לאחר |  |
| 8:27.203 | לא |  |
| 8:27.204 | אתן | . |
| 8:27.205 | והמופת |  |
| 8:27.206 | הגדול |  |
| 8:27.207 | לנו |  |
| 8:27.208 | יציאת |  |
| 8:27.209 | מצרים |  |
| 8:27.210 | וקריעת |  |
| 8:27.211 | ים |  |
| 8:27.212 | סוף |  |
| 8:27.213 | וכל |  |
| 8:27.214 | המופתים |  |
| 8:27.215 | שנעשו |  |
| 8:27.216 | שמה | . |
| 8:27.217 | ולא |  |
| 8:27.218 | מחשבותינו |  |
| 8:27.219 | מחשבותיו |  |
| 8:27.220 | ולא |  |
| 8:27.221 | דרכינו |  |
| 8:27.222 | דרכיו | . |
| 8:27.223 | כי |  |
| 8:27.224 | לא |  |
| 8:27.225 | נוכל |  |
| 8:27.226 | להשיג |  |
| 8:27.227 | מהותו |  |
| 8:27.228 | ואיכותו |  |
| 8:27.229 | ואופן |  |
| 8:27.230 | חנייתו |  |
| 8:27.231 | כמו |  |
| 8:27.232 | שנתבאר |  |
| 8:27.233 | לעיל |  |
| 8:27.234 | סי | ' |
| 8:27.235 | נ"ב |  |
| 8:27.236 | ונ"ג |  |
| 8:27.237 | בשם |  |
| 8:27.238 | הרוקח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

