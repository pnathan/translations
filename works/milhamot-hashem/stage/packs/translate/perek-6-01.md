# Translation pack: milhamot-hashem perek-6-01 (Milchamot ha-Shem 6:1–6:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-01`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ב"ה** → "blessed be He".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- עלת **העלות** → "the Cause of causes".
- הסבה **הראשונה** → "the First Cause".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
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
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
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
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- ואין **צריך** לומר → "needless to say".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 6:1

Hebrew (Wikisource): ס) והנה מכל זה תלמוד לומר דודאי שהדברים שנאמרו בזהר בענייני האלהות לא מפי רשב"י ע"ה נאמרו ולא יצאו דברים כאלה מפי שום תנא או אמורא מרז"ל, ורק תלו אותן ברשב"י וחביריו בשקר, שהרי הוא בעצמו מכריז ואומר כל המשתף שם שמים ודבר אחר נעקר מן העולם! חלילה, חלילה להאמין שרשב"י התנא או שום אחר מרבותינו התנאים או האמוראים יחשוב או יאמר כדברים האלה לעבוד אלוה אחר קצר אפים ולשתפו במחשבה עם ה' אלהינו ארך אפים שהוא הסבה הראשונה שהוא בתור נשמה לפי דעתם לפרצופים האלה ועובדים ומתפללים לפרצופים הנז' ואומרים כי לנשמתם שהוא חלק האין סוף שבהם הם עובדים, ומיחדים הפרצופים הזכרים עם נקיבותיהם הנרמזים בשם הוי"ה, יוד חכמה שהוא אבא, ה"א ראשונה בינה אימא, וי"ו תפארת זעיר, ה"א אחרונה מלכות, ובשמע ישראל הוי"ה ראשונה אבא, אלהינו אמא, הוי"ה אחרונה תפארת זעיר, אחד מלכות, לחברם ולקשרם יחד להיות הארבעה אחד (זהר ואתחנן דף רס"ג ע"א), כי עיקר היחוד אצלם לחבר ולקשר במחשבתו עיניינים כשהם מצויירים במחשבתו גופי אור נפרדים והוא דוחק את מחשבתו לשום אותם יחד, וגם בפיו אומר לשם קב"ה ושכינתיה וכו' ליחדא שם יו"ד וה"א בו"ו וה"א וכו'. וכוונתו על אבא ואימא וזעיר ונוקביה הנרמזים לפי דעתם בשם הוי״ה, ואריך ועתיק הנרמזים בקוצו של יו"ד ועיין לעיל ס' מ"ב. ולפי דבריהם למה לא הזכירו רז"ל במשנה ובגמרא במסכת זבחים על הא דתנן לשם ששה דברים הזבח נזבח, לשם זבח, לשם זובח, לשם השם, לשם אישים, לשם ריח, לשם ניחח, והחטאת והאשם לשם חטא, ופי' בגמרא לשם ניחח, נחת רוח להקב"ה שאמר ונעשה רצונו, ולא קחשיב לשם יחוד קב"ה ושכינתיה, אלא ודאי שדעת רז"ל שאלהינו ב"ה יחיד ומיוחד מכל שאר האחדים, ואינו צריך לקשרו ולאחדו מחלקים רבים כמו שחשבו המקובלים החדשים ולא לשתפו עם פרצופים חלוקים ושונים זה משה! חלילה לנו מהעלות על לבבינו דברים זרים לתורתינו הקדושה כאלה שלא צונו הקב"ה בתורה לאחדו ולשתפו עם שום נברא, רק לדעת ולהאמין כי הוא אחד שנאמר וידעת היום והשבות אל לבביך כי ה' הוא האלהים וכו' אין עוד, ועיין לקמן סי' קל"ו.

| id | word | punctuation after |
|---|---|---|
| 6:1.1 | ס | ) |
| 6:1.2 | והנה |  |
| 6:1.3 | מכל |  |
| 6:1.4 | זה |  |
| 6:1.5 | תלמוד |  |
| 6:1.6 | לומר |  |
| 6:1.7 | דודאי |  |
| 6:1.8 | שהדברים |  |
| 6:1.9 | שנאמרו |  |
| 6:1.10 | בזהר |  |
| 6:1.11 | בענייני |  |
| 6:1.12 | האלהות |  |
| 6:1.13 | לא |  |
| 6:1.14 | מפי |  |
| 6:1.15 | רשב"י |  |
| 6:1.16 | ע"ה |  |
| 6:1.17 | נאמרו |  |
| 6:1.18 | ולא |  |
| 6:1.19 | יצאו |  |
| 6:1.20 | דברים |  |
| 6:1.21 | כאלה |  |
| 6:1.22 | מפי |  |
| 6:1.23 | שום |  |
| 6:1.24 | תנא |  |
| 6:1.25 | או |  |
| 6:1.26 | אמורא |  |
| 6:1.27 | מרז"ל | , |
| 6:1.28 | ורק |  |
| 6:1.29 | תלו |  |
| 6:1.30 | אותן |  |
| 6:1.31 | ברשב"י |  |
| 6:1.32 | וחביריו |  |
| 6:1.33 | בשקר | , |
| 6:1.34 | שהרי |  |
| 6:1.35 | הוא |  |
| 6:1.36 | בעצמו |  |
| 6:1.37 | מכריז |  |
| 6:1.38 | ואומר |  |
| 6:1.39 | כל |  |
| 6:1.40 | המשתף |  |
| 6:1.41 | שם |  |
| 6:1.42 | שמים |  |
| 6:1.43 | ודבר |  |
| 6:1.44 | אחר |  |
| 6:1.45 | נעקר |  |
| 6:1.46 | מן |  |
| 6:1.47 | העולם | ! |
| 6:1.48 | חלילה | , |
| 6:1.49 | חלילה |  |
| 6:1.50 | להאמין |  |
| 6:1.51 | שרשב"י |  |
| 6:1.52 | התנא |  |
| 6:1.53 | או |  |
| 6:1.54 | שום |  |
| 6:1.55 | אחר |  |
| 6:1.56 | מרבותינו |  |
| 6:1.57 | התנאים |  |
| 6:1.58 | או |  |
| 6:1.59 | האמוראים |  |
| 6:1.60 | יחשוב |  |
| 6:1.61 | או |  |
| 6:1.62 | יאמר |  |
| 6:1.63 | כדברים |  |
| 6:1.64 | האלה |  |
| 6:1.65 | לעבוד |  |
| 6:1.66 | אלוה |  |
| 6:1.67 | אחר |  |
| 6:1.68 | קצר |  |
| 6:1.69 | אפים |  |
| 6:1.70 | ולשתפו |  |
| 6:1.71 | במחשבה |  |
| 6:1.72 | עם |  |
| 6:1.73 | ה | ' |
| 6:1.74 | אלהינו |  |
| 6:1.75 | ארך |  |
| 6:1.76 | אפים |  |
| 6:1.77 | שהוא |  |
| 6:1.78 | הסבה |  |
| 6:1.79 | הראשונה |  |
| 6:1.80 | שהוא |  |
| 6:1.81 | בתור |  |
| 6:1.82 | נשמה |  |
| 6:1.83 | לפי |  |
| 6:1.84 | דעתם |  |
| 6:1.85 | לפרצופים |  |
| 6:1.86 | האלה |  |
| 6:1.87 | ועובדים |  |
| 6:1.88 | ומתפללים |  |
| 6:1.89 | לפרצופים |  |
| 6:1.90 | הנז | ' |
| 6:1.91 | ואומרים |  |
| 6:1.92 | כי |  |
| 6:1.93 | לנשמתם |  |
| 6:1.94 | שהוא |  |
| 6:1.95 | חלק |  |
| 6:1.96 | האין |  |
| 6:1.97 | סוף |  |
| 6:1.98 | שבהם |  |
| 6:1.99 | הם |  |
| 6:1.100 | עובדים | , |
| 6:1.101 | ומיחדים |  |
| 6:1.102 | הפרצופים |  |
| 6:1.103 | הזכרים |  |
| 6:1.104 | עם |  |
| 6:1.105 | נקיבותיהם |  |
| 6:1.106 | הנרמזים |  |
| 6:1.107 | בשם |  |
| 6:1.108 | הוי"ה | , |
| 6:1.109 | יוד |  |
| 6:1.110 | חכמה |  |
| 6:1.111 | שהוא |  |
| 6:1.112 | אבא | , |
| 6:1.113 | ה"א |  |
| 6:1.114 | ראשונה |  |
| 6:1.115 | בינה |  |
| 6:1.116 | אימא | , |
| 6:1.117 | וי"ו |  |
| 6:1.118 | תפארת |  |
| 6:1.119 | זעיר | , |
| 6:1.120 | ה"א |  |
| 6:1.121 | אחרונה |  |
| 6:1.122 | מלכות | , |
| 6:1.123 | ובשמע |  |
| 6:1.124 | ישראל |  |
| 6:1.125 | הוי"ה |  |
| 6:1.126 | ראשונה |  |
| 6:1.127 | אבא | , |
| 6:1.128 | אלהינו |  |
| 6:1.129 | אמא | , |
| 6:1.130 | הוי"ה |  |
| 6:1.131 | אחרונה |  |
| 6:1.132 | תפארת |  |
| 6:1.133 | זעיר | , |
| 6:1.134 | אחד |  |
| 6:1.135 | מלכות | , |
| 6:1.136 | לחברם |  |
| 6:1.137 | ולקשרם |  |
| 6:1.138 | יחד |  |
| 6:1.139 | להיות |  |
| 6:1.140 | הארבעה |  |
| 6:1.141 | אחד |  |
| 6:1.142 | זהר |  |
| 6:1.143 | ואתחנן |  |
| 6:1.144 | דף |  |
| 6:1.145 | רס"ג |  |
| 6:1.146 | ע"א | ), |
| 6:1.147 | כי |  |
| 6:1.148 | עיקר |  |
| 6:1.149 | היחוד |  |
| 6:1.150 | אצלם |  |
| 6:1.151 | לחבר |  |
| 6:1.152 | ולקשר |  |
| 6:1.153 | במחשבתו |  |
| 6:1.154 | עיניינים |  |
| 6:1.155 | כשהם |  |
| 6:1.156 | מצויירים |  |
| 6:1.157 | במחשבתו |  |
| 6:1.158 | גופי |  |
| 6:1.159 | אור |  |
| 6:1.160 | נפרדים |  |
| 6:1.161 | והוא |  |
| 6:1.162 | דוחק |  |
| 6:1.163 | את |  |
| 6:1.164 | מחשבתו |  |
| 6:1.165 | לשום |  |
| 6:1.166 | אותם |  |
| 6:1.167 | יחד | , |
| 6:1.168 | וגם |  |
| 6:1.169 | בפיו |  |
| 6:1.170 | אומר |  |
| 6:1.171 | לשם |  |
| 6:1.172 | קב"ה |  |
| 6:1.173 | ושכינתיה |  |
| 6:1.174 | וכו | ' |
| 6:1.175 | ליחדא |  |
| 6:1.176 | שם |  |
| 6:1.177 | יו"ד |  |
| 6:1.178 | וה"א |  |
| 6:1.179 | בו"ו |  |
| 6:1.180 | וה"א |  |
| 6:1.181 | וכו | '. |
| 6:1.182 | וכוונתו |  |
| 6:1.183 | על |  |
| 6:1.184 | אבא |  |
| 6:1.185 | ואימא |  |
| 6:1.186 | וזעיר |  |
| 6:1.187 | ונוקביה |  |
| 6:1.188 | הנרמזים |  |
| 6:1.189 | לפי |  |
| 6:1.190 | דעתם |  |
| 6:1.191 | בשם |  |
| 6:1.192 | הוי״ה | , |
| 6:1.193 | ואריך |  |
| 6:1.194 | ועתיק |  |
| 6:1.195 | הנרמזים |  |
| 6:1.196 | בקוצו |  |
| 6:1.197 | של |  |
| 6:1.198 | יו"ד |  |
| 6:1.199 | ועיין |  |
| 6:1.200 | לעיל |  |
| 6:1.201 | ס | ' |
| 6:1.202 | מ"ב | . |
| 6:1.203 | ולפי |  |
| 6:1.204 | דבריהם |  |
| 6:1.205 | למה |  |
| 6:1.206 | לא |  |
| 6:1.207 | הזכירו |  |
| 6:1.208 | רז"ל |  |
| 6:1.209 | במשנה |  |
| 6:1.210 | ובגמרא |  |
| 6:1.211 | במסכת |  |
| 6:1.212 | זבחים |  |
| 6:1.213 | על |  |
| 6:1.214 | הא |  |
| 6:1.215 | דתנן |  |
| 6:1.216 | לשם |  |
| 6:1.217 | ששה |  |
| 6:1.218 | דברים |  |
| 6:1.219 | הזבח |  |
| 6:1.220 | נזבח | , |
| 6:1.221 | לשם |  |
| 6:1.222 | זבח | , |
| 6:1.223 | לשם |  |
| 6:1.224 | זובח | , |
| 6:1.225 | לשם |  |
| 6:1.226 | השם | , |
| 6:1.227 | לשם |  |
| 6:1.228 | אישים | , |
| 6:1.229 | לשם |  |
| 6:1.230 | ריח | , |
| 6:1.231 | לשם |  |
| 6:1.232 | ניחח | , |
| 6:1.233 | והחטאת |  |
| 6:1.234 | והאשם |  |
| 6:1.235 | לשם |  |
| 6:1.236 | חטא | , |
| 6:1.237 | ופי | ' |
| 6:1.238 | בגמרא |  |
| 6:1.239 | לשם |  |
| 6:1.240 | ניחח | , |
| 6:1.241 | נחת |  |
| 6:1.242 | רוח |  |
| 6:1.243 | להקב"ה |  |
| 6:1.244 | שאמר |  |
| 6:1.245 | ונעשה |  |
| 6:1.246 | רצונו | , |
| 6:1.247 | ולא |  |
| 6:1.248 | קחשיב |  |
| 6:1.249 | לשם |  |
| 6:1.250 | יחוד |  |
| 6:1.251 | קב"ה |  |
| 6:1.252 | ושכינתיה | , |
| 6:1.253 | אלא |  |
| 6:1.254 | ודאי |  |
| 6:1.255 | שדעת |  |
| 6:1.256 | רז"ל |  |
| 6:1.257 | שאלהינו |  |
| 6:1.258 | ב"ה |  |
| 6:1.259 | יחיד |  |
| 6:1.260 | ומיוחד |  |
| 6:1.261 | מכל |  |
| 6:1.262 | שאר |  |
| 6:1.263 | האחדים | , |
| 6:1.264 | ואינו |  |
| 6:1.265 | צריך |  |
| 6:1.266 | לקשרו |  |
| 6:1.267 | ולאחדו |  |
| 6:1.268 | מחלקים |  |
| 6:1.269 | רבים |  |
| 6:1.270 | כמו |  |
| 6:1.271 | שחשבו |  |
| 6:1.272 | המקובלים |  |
| 6:1.273 | החדשים |  |
| 6:1.274 | ולא |  |
| 6:1.275 | לשתפו |  |
| 6:1.276 | עם |  |
| 6:1.277 | פרצופים |  |
| 6:1.278 | חלוקים |  |
| 6:1.279 | ושונים |  |
| 6:1.280 | זה |  |
| 6:1.281 | משה | ! |
| 6:1.282 | חלילה |  |
| 6:1.283 | לנו |  |
| 6:1.284 | מהעלות |  |
| 6:1.285 | על |  |
| 6:1.286 | לבבינו |  |
| 6:1.287 | דברים |  |
| 6:1.288 | זרים |  |
| 6:1.289 | לתורתינו |  |
| 6:1.290 | הקדושה |  |
| 6:1.291 | כאלה |  |
| 6:1.292 | שלא |  |
| 6:1.293 | צונו |  |
| 6:1.294 | הקב"ה |  |
| 6:1.295 | בתורה |  |
| 6:1.296 | לאחדו |  |
| 6:1.297 | ולשתפו |  |
| 6:1.298 | עם |  |
| 6:1.299 | שום |  |
| 6:1.300 | נברא | , |
| 6:1.301 | רק |  |
| 6:1.302 | לדעת |  |
| 6:1.303 | ולהאמין |  |
| 6:1.304 | כי |  |
| 6:1.305 | הוא |  |
| 6:1.306 | אחד |  |
| 6:1.307 | שנאמר |  |
| 6:1.308 | וידעת |  |
| 6:1.309 | היום |  |
| 6:1.310 | והשבות |  |
| 6:1.311 | אל |  |
| 6:1.312 | לבביך |  |
| 6:1.313 | כי |  |
| 6:1.314 | ה | ' |
| 6:1.315 | הוא |  |
| 6:1.316 | האלהים |  |
| 6:1.317 | וכו | ' |
| 6:1.318 | אין |  |
| 6:1.319 | עוד | , |
| 6:1.320 | ועיין |  |
| 6:1.321 | לקמן |  |
| 6:1.322 | סי | ' |
| 6:1.323 | קל"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

