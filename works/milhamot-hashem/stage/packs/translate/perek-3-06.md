# Translation pack: milhamot-hashem perek-3-06 (Milchamot ha-Shem 3:10–3:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-06`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ע"ה** → "peace be upon him" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רדב"ז** → "Radbaz" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **המשרף** (pun on Metzaref ha-Emunah) → "the Mesharef", a note on the pun.
## Terms

- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- הסבה **הראשונה** → "the First Cause".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **מורא** → "fear".
- **מרגל** לרגל → "from festival to festival", not "foot".
- **והוא שלא** → "provided that".
- **קורין** אל → "call upon".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **כמתלהמים** → "ravers" (Prov 26:18 in `apparatus`).
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
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


## Milchamot ha-Shem 3:10

Hebrew (Wikisource): י) ומעתה בעזר השי"ת נשוב לעיקר שאילתנו ששאלנו לך להודיענו על פי שיטת המקובלים, למי אנו קורין בתפלותנו ומברכים כל ברכותנו ועבודתנו וממי אנו מבקשים סליחה וכפרה. והיתה תשובתך בדברי נרגן כמתלהמים טענו חטים והודה לו בשעורים והכחשת את הידוע ומפורסם בספריהם לכל, כנשבע על האיש שהוא אשה, ועל האשה שהיא איש, ועל עמוד של שיש שהוא זהב. וחרה אפך על ספר מצרף האמונה, והלעגת עליו ובזית אותו, וקראת אותו ספר המשרף וכו' ובזה גלית לכל שאתה הוא המבזה את החכמים, ועל שאלותינו כיון שזעיר אנפין לפי דבריהם הוא אלוהינו, והוא הזן אותנו ומפרנסינו, ואותו נעבוד, אם כן כל העולמות שלמעלה מעולם האצילות, מי זן אותם ומפרנסם? אם זורק הוא שפעם ופרנסתם מלמטה למעלה, והם עובדים ומשתחוים לו, או שמא אלהיהם הזנם ומפרנסם, הוא עלה הראשונה, ולו הם עובדים, ואנחנו עובדים לזעיר אנפין הזנינו ומפרנסינו? ואם כן אלהי העליונים אינו אלהי התחתונים ח"ו? וזהו מקום הסתירה אשר בקשנו ממך להודיענו ולהורותנו, ועל זה באה תשובתך שאין הנדון דומה לראיה וכי מביאין ראיה לסתור הלכה למשה מסיני? ושחכמי הקבלה אומרים כך קיבלו ממשה.

| id | word | punctuation after |
|---|---|---|
| 3:10.1 | י | ) |
| 3:10.2 | ומעתה |  |
| 3:10.3 | בעזר |  |
| 3:10.4 | השי"ת |  |
| 3:10.5 | נשוב |  |
| 3:10.6 | לעיקר |  |
| 3:10.7 | שאילתנו |  |
| 3:10.8 | ששאלנו |  |
| 3:10.9 | לך |  |
| 3:10.10 | להודיענו |  |
| 3:10.11 | על |  |
| 3:10.12 | פי |  |
| 3:10.13 | שיטת |  |
| 3:10.14 | המקובלים | , |
| 3:10.15 | למי |  |
| 3:10.16 | אנו |  |
| 3:10.17 | קורין |  |
| 3:10.18 | בתפלותנו |  |
| 3:10.19 | ומברכים |  |
| 3:10.20 | כל |  |
| 3:10.21 | ברכותנו |  |
| 3:10.22 | ועבודתנו |  |
| 3:10.23 | וממי |  |
| 3:10.24 | אנו |  |
| 3:10.25 | מבקשים |  |
| 3:10.26 | סליחה |  |
| 3:10.27 | וכפרה | . |
| 3:10.28 | והיתה |  |
| 3:10.29 | תשובתך |  |
| 3:10.30 | בדברי |  |
| 3:10.31 | נרגן |  |
| 3:10.32 | כמתלהמים |  |
| 3:10.33 | טענו |  |
| 3:10.34 | חטים |  |
| 3:10.35 | והודה |  |
| 3:10.36 | לו |  |
| 3:10.37 | בשעורים |  |
| 3:10.38 | והכחשת |  |
| 3:10.39 | את |  |
| 3:10.40 | הידוע |  |
| 3:10.41 | ומפורסם |  |
| 3:10.42 | בספריהם |  |
| 3:10.43 | לכל | , |
| 3:10.44 | כנשבע |  |
| 3:10.45 | על |  |
| 3:10.46 | האיש |  |
| 3:10.47 | שהוא |  |
| 3:10.48 | אשה | , |
| 3:10.49 | ועל |  |
| 3:10.50 | האשה |  |
| 3:10.51 | שהיא |  |
| 3:10.52 | איש | , |
| 3:10.53 | ועל |  |
| 3:10.54 | עמוד |  |
| 3:10.55 | של |  |
| 3:10.56 | שיש |  |
| 3:10.57 | שהוא |  |
| 3:10.58 | זהב | . |
| 3:10.59 | וחרה |  |
| 3:10.60 | אפך |  |
| 3:10.61 | על |  |
| 3:10.62 | ספר |  |
| 3:10.63 | מצרף |  |
| 3:10.64 | האמונה | , |
| 3:10.65 | והלעגת |  |
| 3:10.66 | עליו |  |
| 3:10.67 | ובזית |  |
| 3:10.68 | אותו | , |
| 3:10.69 | וקראת |  |
| 3:10.70 | אותו |  |
| 3:10.71 | ספר |  |
| 3:10.72 | המשרף |  |
| 3:10.73 | וכו | ' |
| 3:10.74 | ובזה |  |
| 3:10.75 | גלית |  |
| 3:10.76 | לכל |  |
| 3:10.77 | שאתה |  |
| 3:10.78 | הוא |  |
| 3:10.79 | המבזה |  |
| 3:10.80 | את |  |
| 3:10.81 | החכמים | , |
| 3:10.82 | ועל |  |
| 3:10.83 | שאלותינו |  |
| 3:10.84 | כיון |  |
| 3:10.85 | שזעיר |  |
| 3:10.86 | אנפין |  |
| 3:10.87 | לפי |  |
| 3:10.88 | דבריהם |  |
| 3:10.89 | הוא |  |
| 3:10.90 | אלוהינו | , |
| 3:10.91 | והוא |  |
| 3:10.92 | הזן |  |
| 3:10.93 | אותנו |  |
| 3:10.94 | ומפרנסינו | , |
| 3:10.95 | ואותו |  |
| 3:10.96 | נעבוד | , |
| 3:10.97 | אם |  |
| 3:10.98 | כן |  |
| 3:10.99 | כל |  |
| 3:10.100 | העולמות |  |
| 3:10.101 | שלמעלה |  |
| 3:10.102 | מעולם |  |
| 3:10.103 | האצילות | , |
| 3:10.104 | מי |  |
| 3:10.105 | זן |  |
| 3:10.106 | אותם |  |
| 3:10.107 | ומפרנסם | ? |
| 3:10.108 | אם |  |
| 3:10.109 | זורק |  |
| 3:10.110 | הוא |  |
| 3:10.111 | שפעם |  |
| 3:10.112 | ופרנסתם |  |
| 3:10.113 | מלמטה |  |
| 3:10.114 | למעלה | , |
| 3:10.115 | והם |  |
| 3:10.116 | עובדים |  |
| 3:10.117 | ומשתחוים |  |
| 3:10.118 | לו | , |
| 3:10.119 | או |  |
| 3:10.120 | שמא |  |
| 3:10.121 | אלהיהם |  |
| 3:10.122 | הזנם |  |
| 3:10.123 | ומפרנסם | , |
| 3:10.124 | הוא |  |
| 3:10.125 | עלה |  |
| 3:10.126 | הראשונה | , |
| 3:10.127 | ולו |  |
| 3:10.128 | הם |  |
| 3:10.129 | עובדים | , |
| 3:10.130 | ואנחנו |  |
| 3:10.131 | עובדים |  |
| 3:10.132 | לזעיר |  |
| 3:10.133 | אנפין |  |
| 3:10.134 | הזנינו |  |
| 3:10.135 | ומפרנסינו | ? |
| 3:10.136 | ואם |  |
| 3:10.137 | כן |  |
| 3:10.138 | אלהי |  |
| 3:10.139 | העליונים |  |
| 3:10.140 | אינו |  |
| 3:10.141 | אלהי |  |
| 3:10.142 | התחתונים |  |
| 3:10.143 | ח"ו | ? |
| 3:10.144 | וזהו |  |
| 3:10.145 | מקום |  |
| 3:10.146 | הסתירה |  |
| 3:10.147 | אשר |  |
| 3:10.148 | בקשנו |  |
| 3:10.149 | ממך |  |
| 3:10.150 | להודיענו |  |
| 3:10.151 | ולהורותנו | , |
| 3:10.152 | ועל |  |
| 3:10.153 | זה |  |
| 3:10.154 | באה |  |
| 3:10.155 | תשובתך |  |
| 3:10.156 | שאין |  |
| 3:10.157 | הנדון |  |
| 3:10.158 | דומה |  |
| 3:10.159 | לראיה |  |
| 3:10.160 | וכי |  |
| 3:10.161 | מביאין |  |
| 3:10.162 | ראיה |  |
| 3:10.163 | לסתור |  |
| 3:10.164 | הלכה |  |
| 3:10.165 | למשה |  |
| 3:10.166 | מסיני | ? |
| 3:10.167 | ושחכמי |  |
| 3:10.168 | הקבלה |  |
| 3:10.169 | אומרים |  |
| 3:10.170 | כך |  |
| 3:10.171 | קיבלו |  |
| 3:10.172 | ממשה | . |


## Milchamot ha-Shem 3:11

Hebrew (Wikisource): יא) מאד אנו תמהים ומתפלאים על דבריך ועל תשובתך זאת, וכי הלכה למשה מסיני להאמין באלוהות רבים? ולעבוד אלהים אחרים מבלעדי ה' אלהינו הראשון וצור לכל? והוא סבת כל הסבות כולם, ובמרמה את בא בתשובותיך אלה לאנשים עזי לב אשר התפארת עליהם והראית להם דבריך אלה לאמר. ראו כי סתרתי את דבריהם, ודחיתי את ראיותיהם בקש, וחוזר קונדריסך על כל המסובין, ואחר אשר קילסוך ושבחוך שאתה חריף ושנון הבאת אותה לנו. והנה כולה הטעאות כאלו לא ידעת על מה תסוב שאלתינו, ועל טעותך הכוזבת הלזו נשוב עוד, כי ידוע ומפורסם בכללי הפוסקים (כנה"ג ורדב"ז ושה"ז) שכל מקום שהקבלה חולקת על הפסק הלך אחרי הפסק, והכלל הזה מפורסם בישראל. ואם דבריך שהקבלה הלזו היא הלכה למשה מסיני, וכי תדחה הלכה למשה מפני סברת איזה תנא או אמורא, וכ"ש פוסק אחרון (עיין לקמן סי' ט"ו) ועליך להביא ראיה ולהודיענו היכן מצינו זה ששום תנא או אמורא יעיז פניו לחלוק ולהורות היפך הלכה המקובלת ממשה רבי' ע"ה. הלא שוטה או רשע וגס רוח יקרא? ובכמה מקומות בתלמוד מרגלא בפומיהו דרבנן לומר אם הלכה נקבל, ואם לדין יש תשובה.

| id | word | punctuation after |
|---|---|---|
| 3:11.1 | יא | ) |
| 3:11.2 | מאד |  |
| 3:11.3 | אנו |  |
| 3:11.4 | תמהים |  |
| 3:11.5 | ומתפלאים |  |
| 3:11.6 | על |  |
| 3:11.7 | דבריך |  |
| 3:11.8 | ועל |  |
| 3:11.9 | תשובתך |  |
| 3:11.10 | זאת | , |
| 3:11.11 | וכי |  |
| 3:11.12 | הלכה |  |
| 3:11.13 | למשה |  |
| 3:11.14 | מסיני |  |
| 3:11.15 | להאמין |  |
| 3:11.16 | באלוהות |  |
| 3:11.17 | רבים | ? |
| 3:11.18 | ולעבוד |  |
| 3:11.19 | אלהים |  |
| 3:11.20 | אחרים |  |
| 3:11.21 | מבלעדי |  |
| 3:11.22 | ה | ' |
| 3:11.23 | אלהינו |  |
| 3:11.24 | הראשון |  |
| 3:11.25 | וצור |  |
| 3:11.26 | לכל | ? |
| 3:11.27 | והוא |  |
| 3:11.28 | סבת |  |
| 3:11.29 | כל |  |
| 3:11.30 | הסבות |  |
| 3:11.31 | כולם | , |
| 3:11.32 | ובמרמה |  |
| 3:11.33 | את |  |
| 3:11.34 | בא |  |
| 3:11.35 | בתשובותיך |  |
| 3:11.36 | אלה |  |
| 3:11.37 | לאנשים |  |
| 3:11.38 | עזי |  |
| 3:11.39 | לב |  |
| 3:11.40 | אשר |  |
| 3:11.41 | התפארת |  |
| 3:11.42 | עליהם |  |
| 3:11.43 | והראית |  |
| 3:11.44 | להם |  |
| 3:11.45 | דבריך |  |
| 3:11.46 | אלה |  |
| 3:11.47 | לאמר | . |
| 3:11.48 | ראו |  |
| 3:11.49 | כי |  |
| 3:11.50 | סתרתי |  |
| 3:11.51 | את |  |
| 3:11.52 | דבריהם | , |
| 3:11.53 | ודחיתי |  |
| 3:11.54 | את |  |
| 3:11.55 | ראיותיהם |  |
| 3:11.56 | בקש | , |
| 3:11.57 | וחוזר |  |
| 3:11.58 | קונדריסך |  |
| 3:11.59 | על |  |
| 3:11.60 | כל |  |
| 3:11.61 | המסובין | , |
| 3:11.62 | ואחר |  |
| 3:11.63 | אשר |  |
| 3:11.64 | קילסוך |  |
| 3:11.65 | ושבחוך |  |
| 3:11.66 | שאתה |  |
| 3:11.67 | חריף |  |
| 3:11.68 | ושנון |  |
| 3:11.69 | הבאת |  |
| 3:11.70 | אותה |  |
| 3:11.71 | לנו | . |
| 3:11.72 | והנה |  |
| 3:11.73 | כולה |  |
| 3:11.74 | הטעאות |  |
| 3:11.75 | כאלו |  |
| 3:11.76 | לא |  |
| 3:11.77 | ידעת |  |
| 3:11.78 | על |  |
| 3:11.79 | מה |  |
| 3:11.80 | תסוב |  |
| 3:11.81 | שאלתינו | , |
| 3:11.82 | ועל |  |
| 3:11.83 | טעותך |  |
| 3:11.84 | הכוזבת |  |
| 3:11.85 | הלזו |  |
| 3:11.86 | נשוב |  |
| 3:11.87 | עוד | , |
| 3:11.88 | כי |  |
| 3:11.89 | ידוע |  |
| 3:11.90 | ומפורסם |  |
| 3:11.91 | בכללי |  |
| 3:11.92 | הפוסקים |  |
| 3:11.93 | כנה"ג |  |
| 3:11.94 | ורדב"ז |  |
| 3:11.95 | ושה"ז | ) |
| 3:11.96 | שכל |  |
| 3:11.97 | מקום |  |
| 3:11.98 | שהקבלה |  |
| 3:11.99 | חולקת |  |
| 3:11.100 | על |  |
| 3:11.101 | הפסק |  |
| 3:11.102 | הלך |  |
| 3:11.103 | אחרי |  |
| 3:11.104 | הפסק | , |
| 3:11.105 | והכלל |  |
| 3:11.106 | הזה |  |
| 3:11.107 | מפורסם |  |
| 3:11.108 | בישראל | . |
| 3:11.109 | ואם |  |
| 3:11.110 | דבריך |  |
| 3:11.111 | שהקבלה |  |
| 3:11.112 | הלזו |  |
| 3:11.113 | היא |  |
| 3:11.114 | הלכה |  |
| 3:11.115 | למשה |  |
| 3:11.116 | מסיני | , |
| 3:11.117 | וכי |  |
| 3:11.118 | תדחה |  |
| 3:11.119 | הלכה |  |
| 3:11.120 | למשה |  |
| 3:11.121 | מפני |  |
| 3:11.122 | סברת |  |
| 3:11.123 | איזה |  |
| 3:11.124 | תנא |  |
| 3:11.125 | או |  |
| 3:11.126 | אמורא | , |
| 3:11.127 | וכ"ש |  |
| 3:11.128 | פוסק |  |
| 3:11.129 | אחרון |  |
| 3:11.130 | עיין |  |
| 3:11.131 | לקמן |  |
| 3:11.132 | סי | ' |
| 3:11.133 | ט"ו | ) |
| 3:11.134 | ועליך |  |
| 3:11.135 | להביא |  |
| 3:11.136 | ראיה |  |
| 3:11.137 | ולהודיענו |  |
| 3:11.138 | היכן |  |
| 3:11.139 | מצינו |  |
| 3:11.140 | זה |  |
| 3:11.141 | ששום |  |
| 3:11.142 | תנא |  |
| 3:11.143 | או |  |
| 3:11.144 | אמורא |  |
| 3:11.145 | יעיז |  |
| 3:11.146 | פניו |  |
| 3:11.147 | לחלוק |  |
| 3:11.148 | ולהורות |  |
| 3:11.149 | היפך |  |
| 3:11.150 | הלכה |  |
| 3:11.151 | המקובלת |  |
| 3:11.152 | ממשה |  |
| 3:11.153 | רבי | ' |
| 3:11.154 | ע"ה | . |
| 3:11.155 | הלא |  |
| 3:11.156 | שוטה |  |
| 3:11.157 | או |  |
| 3:11.158 | רשע |  |
| 3:11.159 | וגס |  |
| 3:11.160 | רוח |  |
| 3:11.161 | יקרא | ? |
| 3:11.162 | ובכמה |  |
| 3:11.163 | מקומות |  |
| 3:11.164 | בתלמוד |  |
| 3:11.165 | מרגלא |  |
| 3:11.166 | בפומיהו |  |
| 3:11.167 | דרבנן |  |
| 3:11.168 | לומר |  |
| 3:11.169 | אם |  |
| 3:11.170 | הלכה |  |
| 3:11.171 | נקבל | , |
| 3:11.172 | ואם |  |
| 3:11.173 | לדין |  |
| 3:11.174 | יש |  |
| 3:11.175 | תשובה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

