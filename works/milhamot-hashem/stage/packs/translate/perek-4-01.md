# Translation pack: milhamot-hashem perek-4-01 (Milchamot ha-Shem 4:1–4:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-01`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
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
- **בא"י** → "in the Land of Israel".

## Terms

- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- סבא **דסבין** → "the Elder of Elders".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **רזא דמלה** → "the secret of the matter".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 4:1

Hebrew (Wikisource): כ) ומעתה הט אזנך ושמע תשובות תורת משה אמת. לכבוד ה' אלהים אמת. מה שכתבת והשבת על דברי להכחיש שהם קורים לחמשה פרצופים אלוהות ואמרת ח"ו שהם אומרים כן כסות עינים היה לך שלא ראית דברי הזהר בראשית כ"ב וז"ל ויאמר אלהים נעשה אדם בצלמינו כדמותינו. סוד ה' ליריאיו. פתח ההוא סבא דסבין ואמר שמעון שמעון מאן הוא דאמר ויאמר אלהים. מאי ניהו האי אלהים! אדהכי פרח ההוא סבא דסבין ולא חמא ליה. וכמא דשמע ר' שמעון דהוה קרי ליה שמעון שמעון ולא רבי שמעון. אמר לחברוי ודאי האי הוא קב"ה דאתאמר ביה ועתיק יומין יתיב (פירוש הר' דרך אמת זקן הזקנים והוא כתר עליון. והמקדש מלך כ' סבא דסבין הוא עתיק כמו שאמר רשב"י לקמן כי אבא ואמא נקראו זקנים ועתיק ואריך נקראין סבא דסבין וקימא ליה לרשב"י שהוא עתיק ולא אריך) הא כען איהו שעתא למפתח בהאי רזא דודאי הכא אית רזא דלא אתיהיב רשו לאתגליא. פתח ואמר למלכא דהוה ליה כמה בניינין למבני. וההוא אומנא לא הוה עבד מדעם אלא מרשו דמלכא כמה דאתאמר ואהיה אצלו אמון. מלכא ודאי איהו חכמה עילאה לעילא. ועמודא דאמצעיתא מלכא לתתא, אלהים אומנא לעילא, ודא אימא עלאה אלהים אומנא לתתא ודא שכנתא ואתתא לית לה רשו למעבד מדעם בלא רשו בעלה, וכל בניינין דהוו באורח אצילותה הוה אמר אבא באמירה לגבי אימא יהא כדין וכדין ומיד הוה, כמא דאתמר ויאמר אלהים יהי אור, ויהי אור, ויאמר הוה אמר יהי אור מארי דבניינא איהו אמר, ואומנא עביד מיד, והכי כל בנינין באורח אצילותא הוה אמר יהי רקיע, יהי מארת, וכלא אתעביד מיד, כד מטא לעלמא דפירודא דאיהו עולם הנבדלים, אמר אומנא למארי בניינא, נעשה אדם בצלמינו כדמותינו, אמר מארי בניינא ודאי טב הוא למעבד ליה אבל עתיד הוא למחטי קמך בגין דאיהו בן כסיל הה"ד בן חכם ישמח אב ובן כסיל תוגת אמו, איהי אמרה, כיון דחוביה תליא באימא ולא באבא, אנא בעינא למברי ליה בדיוקנא דילי, הה"ד ויברא אלהים את האדם בצלמו, ולא בעא לאשתתפא ביה אבא, בזמנא דחב מה כתיב ובפשעיכם שולחה אמכם. אמר מלכא לאימא, ולא אמרית לך דעתיד למחטי, בההוא זמנא תריך ליה ותריך אימא עמיה, ובגין דא כתיב בן חכם ישמח אב דא אדם דאיהו באורח אצילותא, ובן כסיל דא אדם דבריאה עכ"ל. ופירש הר' שביבי נגה דמהכא משמע שהקב"ה בעצמו כביכול לא הסכים בבריאה אדם רק השכינה התחננה להיות אדם בגורלה ע"כ, ופי' הר' מקדש מלך בן חכם דא אדם דאצילות, הוא זעיר דאצילות, אדם דבריאה הוא אדם הראשון וע"כ. ולעיל מזה כ' הר' מקדש מלך בדבור המתחיל אמר אומנא פי' אימא אמרה לאבא נעשה אדם, וזהו ויאמר אלהים שהיא אימא אמרה לאבא נעשה אדם מה שאין כן כל ויאמר דמעשה בראשית שרוצה לומר ויאמר אבא לאלהים שהיא אימא (בינה) כי אבא אומר ובינה עושה ועכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:1.1 | כ | ) |
| 4:1.2 | ומעתה |  |
| 4:1.3 | הט |  |
| 4:1.4 | אזנך |  |
| 4:1.5 | ושמע |  |
| 4:1.6 | תשובות |  |
| 4:1.7 | תורת |  |
| 4:1.8 | משה |  |
| 4:1.9 | אמת | . |
| 4:1.10 | לכבוד |  |
| 4:1.11 | ה | ' |
| 4:1.12 | אלהים |  |
| 4:1.13 | אמת | . |
| 4:1.14 | מה |  |
| 4:1.15 | שכתבת |  |
| 4:1.16 | והשבת |  |
| 4:1.17 | על |  |
| 4:1.18 | דברי |  |
| 4:1.19 | להכחיש |  |
| 4:1.20 | שהם |  |
| 4:1.21 | קורים |  |
| 4:1.22 | לחמשה |  |
| 4:1.23 | פרצופים |  |
| 4:1.24 | אלוהות |  |
| 4:1.25 | ואמרת |  |
| 4:1.26 | ח"ו |  |
| 4:1.27 | שהם |  |
| 4:1.28 | אומרים |  |
| 4:1.29 | כן |  |
| 4:1.30 | כסות |  |
| 4:1.31 | עינים |  |
| 4:1.32 | היה |  |
| 4:1.33 | לך |  |
| 4:1.34 | שלא |  |
| 4:1.35 | ראית |  |
| 4:1.36 | דברי |  |
| 4:1.37 | הזהר |  |
| 4:1.38 | בראשית |  |
| 4:1.39 | כ"ב |  |
| 4:1.40 | וז"ל |  |
| 4:1.41 | ויאמר |  |
| 4:1.42 | אלהים |  |
| 4:1.43 | נעשה |  |
| 4:1.44 | אדם |  |
| 4:1.45 | בצלמינו |  |
| 4:1.46 | כדמותינו | . |
| 4:1.47 | סוד |  |
| 4:1.48 | ה | ' |
| 4:1.49 | ליריאיו | . |
| 4:1.50 | פתח |  |
| 4:1.51 | ההוא |  |
| 4:1.52 | סבא |  |
| 4:1.53 | דסבין |  |
| 4:1.54 | ואמר |  |
| 4:1.55 | שמעון |  |
| 4:1.56 | שמעון |  |
| 4:1.57 | מאן |  |
| 4:1.58 | הוא |  |
| 4:1.59 | דאמר |  |
| 4:1.60 | ויאמר |  |
| 4:1.61 | אלהים | . |
| 4:1.62 | מאי |  |
| 4:1.63 | ניהו |  |
| 4:1.64 | האי |  |
| 4:1.65 | אלהים | ! |
| 4:1.66 | אדהכי |  |
| 4:1.67 | פרח |  |
| 4:1.68 | ההוא |  |
| 4:1.69 | סבא |  |
| 4:1.70 | דסבין |  |
| 4:1.71 | ולא |  |
| 4:1.72 | חמא |  |
| 4:1.73 | ליה | . |
| 4:1.74 | וכמא |  |
| 4:1.75 | דשמע |  |
| 4:1.76 | ר | ' |
| 4:1.77 | שמעון |  |
| 4:1.78 | דהוה |  |
| 4:1.79 | קרי |  |
| 4:1.80 | ליה |  |
| 4:1.81 | שמעון |  |
| 4:1.82 | שמעון |  |
| 4:1.83 | ולא |  |
| 4:1.84 | רבי |  |
| 4:1.85 | שמעון | . |
| 4:1.86 | אמר |  |
| 4:1.87 | לחברוי |  |
| 4:1.88 | ודאי |  |
| 4:1.89 | האי |  |
| 4:1.90 | הוא |  |
| 4:1.91 | קב"ה |  |
| 4:1.92 | דאתאמר |  |
| 4:1.93 | ביה |  |
| 4:1.94 | ועתיק |  |
| 4:1.95 | יומין |  |
| 4:1.96 | יתיב |  |
| 4:1.97 | פירוש |  |
| 4:1.98 | הר | ' |
| 4:1.99 | דרך |  |
| 4:1.100 | אמת |  |
| 4:1.101 | זקן |  |
| 4:1.102 | הזקנים |  |
| 4:1.103 | והוא |  |
| 4:1.104 | כתר |  |
| 4:1.105 | עליון | . |
| 4:1.106 | והמקדש |  |
| 4:1.107 | מלך |  |
| 4:1.108 | כ | ' |
| 4:1.109 | סבא |  |
| 4:1.110 | דסבין |  |
| 4:1.111 | הוא |  |
| 4:1.112 | עתיק |  |
| 4:1.113 | כמו |  |
| 4:1.114 | שאמר |  |
| 4:1.115 | רשב"י |  |
| 4:1.116 | לקמן |  |
| 4:1.117 | כי |  |
| 4:1.118 | אבא |  |
| 4:1.119 | ואמא |  |
| 4:1.120 | נקראו |  |
| 4:1.121 | זקנים |  |
| 4:1.122 | ועתיק |  |
| 4:1.123 | ואריך |  |
| 4:1.124 | נקראין |  |
| 4:1.125 | סבא |  |
| 4:1.126 | דסבין |  |
| 4:1.127 | וקימא |  |
| 4:1.128 | ליה |  |
| 4:1.129 | לרשב"י |  |
| 4:1.130 | שהוא |  |
| 4:1.131 | עתיק |  |
| 4:1.132 | ולא |  |
| 4:1.133 | אריך | ) |
| 4:1.134 | הא |  |
| 4:1.135 | כען |  |
| 4:1.136 | איהו |  |
| 4:1.137 | שעתא |  |
| 4:1.138 | למפתח |  |
| 4:1.139 | בהאי |  |
| 4:1.140 | רזא |  |
| 4:1.141 | דודאי |  |
| 4:1.142 | הכא |  |
| 4:1.143 | אית |  |
| 4:1.144 | רזא |  |
| 4:1.145 | דלא |  |
| 4:1.146 | אתיהיב |  |
| 4:1.147 | רשו |  |
| 4:1.148 | לאתגליא | . |
| 4:1.149 | פתח |  |
| 4:1.150 | ואמר |  |
| 4:1.151 | למלכא |  |
| 4:1.152 | דהוה |  |
| 4:1.153 | ליה |  |
| 4:1.154 | כמה |  |
| 4:1.155 | בניינין |  |
| 4:1.156 | למבני | . |
| 4:1.157 | וההוא |  |
| 4:1.158 | אומנא |  |
| 4:1.159 | לא |  |
| 4:1.160 | הוה |  |
| 4:1.161 | עבד |  |
| 4:1.162 | מדעם |  |
| 4:1.163 | אלא |  |
| 4:1.164 | מרשו |  |
| 4:1.165 | דמלכא |  |
| 4:1.166 | כמה |  |
| 4:1.167 | דאתאמר |  |
| 4:1.168 | ואהיה |  |
| 4:1.169 | אצלו |  |
| 4:1.170 | אמון | . |
| 4:1.171 | מלכא |  |
| 4:1.172 | ודאי |  |
| 4:1.173 | איהו |  |
| 4:1.174 | חכמה |  |
| 4:1.175 | עילאה |  |
| 4:1.176 | לעילא | . |
| 4:1.177 | ועמודא |  |
| 4:1.178 | דאמצעיתא |  |
| 4:1.179 | מלכא |  |
| 4:1.180 | לתתא | , |
| 4:1.181 | אלהים |  |
| 4:1.182 | אומנא |  |
| 4:1.183 | לעילא | , |
| 4:1.184 | ודא |  |
| 4:1.185 | אימא |  |
| 4:1.186 | עלאה |  |
| 4:1.187 | אלהים |  |
| 4:1.188 | אומנא |  |
| 4:1.189 | לתתא |  |
| 4:1.190 | ודא |  |
| 4:1.191 | שכנתא |  |
| 4:1.192 | ואתתא |  |
| 4:1.193 | לית |  |
| 4:1.194 | לה |  |
| 4:1.195 | רשו |  |
| 4:1.196 | למעבד |  |
| 4:1.197 | מדעם |  |
| 4:1.198 | בלא |  |
| 4:1.199 | רשו |  |
| 4:1.200 | בעלה | , |
| 4:1.201 | וכל |  |
| 4:1.202 | בניינין |  |
| 4:1.203 | דהוו |  |
| 4:1.204 | באורח |  |
| 4:1.205 | אצילותה |  |
| 4:1.206 | הוה |  |
| 4:1.207 | אמר |  |
| 4:1.208 | אבא |  |
| 4:1.209 | באמירה |  |
| 4:1.210 | לגבי |  |
| 4:1.211 | אימא |  |
| 4:1.212 | יהא |  |
| 4:1.213 | כדין |  |
| 4:1.214 | וכדין |  |
| 4:1.215 | ומיד |  |
| 4:1.216 | הוה | , |
| 4:1.217 | כמא |  |
| 4:1.218 | דאתמר |  |
| 4:1.219 | ויאמר |  |
| 4:1.220 | אלהים |  |
| 4:1.221 | יהי |  |
| 4:1.222 | אור | , |
| 4:1.223 | ויהי |  |
| 4:1.224 | אור | , |
| 4:1.225 | ויאמר |  |
| 4:1.226 | הוה |  |
| 4:1.227 | אמר |  |
| 4:1.228 | יהי |  |
| 4:1.229 | אור |  |
| 4:1.230 | מארי |  |
| 4:1.231 | דבניינא |  |
| 4:1.232 | איהו |  |
| 4:1.233 | אמר | , |
| 4:1.234 | ואומנא |  |
| 4:1.235 | עביד |  |
| 4:1.236 | מיד | , |
| 4:1.237 | והכי |  |
| 4:1.238 | כל |  |
| 4:1.239 | בנינין |  |
| 4:1.240 | באורח |  |
| 4:1.241 | אצילותא |  |
| 4:1.242 | הוה |  |
| 4:1.243 | אמר |  |
| 4:1.244 | יהי |  |
| 4:1.245 | רקיע | , |
| 4:1.246 | יהי |  |
| 4:1.247 | מארת | , |
| 4:1.248 | וכלא |  |
| 4:1.249 | אתעביד |  |
| 4:1.250 | מיד | , |
| 4:1.251 | כד |  |
| 4:1.252 | מטא |  |
| 4:1.253 | לעלמא |  |
| 4:1.254 | דפירודא |  |
| 4:1.255 | דאיהו |  |
| 4:1.256 | עולם |  |
| 4:1.257 | הנבדלים | , |
| 4:1.258 | אמר |  |
| 4:1.259 | אומנא |  |
| 4:1.260 | למארי |  |
| 4:1.261 | בניינא | , |
| 4:1.262 | נעשה |  |
| 4:1.263 | אדם |  |
| 4:1.264 | בצלמינו |  |
| 4:1.265 | כדמותינו | , |
| 4:1.266 | אמר |  |
| 4:1.267 | מארי |  |
| 4:1.268 | בניינא |  |
| 4:1.269 | ודאי |  |
| 4:1.270 | טב |  |
| 4:1.271 | הוא |  |
| 4:1.272 | למעבד |  |
| 4:1.273 | ליה |  |
| 4:1.274 | אבל |  |
| 4:1.275 | עתיד |  |
| 4:1.276 | הוא |  |
| 4:1.277 | למחטי |  |
| 4:1.278 | קמך |  |
| 4:1.279 | בגין |  |
| 4:1.280 | דאיהו |  |
| 4:1.281 | בן |  |
| 4:1.282 | כסיל |  |
| 4:1.283 | הה"ד |  |
| 4:1.284 | בן |  |
| 4:1.285 | חכם |  |
| 4:1.286 | ישמח |  |
| 4:1.287 | אב |  |
| 4:1.288 | ובן |  |
| 4:1.289 | כסיל |  |
| 4:1.290 | תוגת |  |
| 4:1.291 | אמו | , |
| 4:1.292 | איהי |  |
| 4:1.293 | אמרה | , |
| 4:1.294 | כיון |  |
| 4:1.295 | דחוביה |  |
| 4:1.296 | תליא |  |
| 4:1.297 | באימא |  |
| 4:1.298 | ולא |  |
| 4:1.299 | באבא | , |
| 4:1.300 | אנא |  |
| 4:1.301 | בעינא |  |
| 4:1.302 | למברי |  |
| 4:1.303 | ליה |  |
| 4:1.304 | בדיוקנא |  |
| 4:1.305 | דילי | , |
| 4:1.306 | הה"ד |  |
| 4:1.307 | ויברא |  |
| 4:1.308 | אלהים |  |
| 4:1.309 | את |  |
| 4:1.310 | האדם |  |
| 4:1.311 | בצלמו | , |
| 4:1.312 | ולא |  |
| 4:1.313 | בעא |  |
| 4:1.314 | לאשתתפא |  |
| 4:1.315 | ביה |  |
| 4:1.316 | אבא | , |
| 4:1.317 | בזמנא |  |
| 4:1.318 | דחב |  |
| 4:1.319 | מה |  |
| 4:1.320 | כתיב |  |
| 4:1.321 | ובפשעיכם |  |
| 4:1.322 | שולחה |  |
| 4:1.323 | אמכם | . |
| 4:1.324 | אמר |  |
| 4:1.325 | מלכא |  |
| 4:1.326 | לאימא | , |
| 4:1.327 | ולא |  |
| 4:1.328 | אמרית |  |
| 4:1.329 | לך |  |
| 4:1.330 | דעתיד |  |
| 4:1.331 | למחטי | , |
| 4:1.332 | בההוא |  |
| 4:1.333 | זמנא |  |
| 4:1.334 | תריך |  |
| 4:1.335 | ליה |  |
| 4:1.336 | ותריך |  |
| 4:1.337 | אימא |  |
| 4:1.338 | עמיה | , |
| 4:1.339 | ובגין |  |
| 4:1.340 | דא |  |
| 4:1.341 | כתיב |  |
| 4:1.342 | בן |  |
| 4:1.343 | חכם |  |
| 4:1.344 | ישמח |  |
| 4:1.345 | אב |  |
| 4:1.346 | דא |  |
| 4:1.347 | אדם |  |
| 4:1.348 | דאיהו |  |
| 4:1.349 | באורח |  |
| 4:1.350 | אצילותא | , |
| 4:1.351 | ובן |  |
| 4:1.352 | כסיל |  |
| 4:1.353 | דא |  |
| 4:1.354 | אדם |  |
| 4:1.355 | דבריאה |  |
| 4:1.356 | עכ"ל | . |
| 4:1.357 | ופירש |  |
| 4:1.358 | הר | ' |
| 4:1.359 | שביבי |  |
| 4:1.360 | נגה |  |
| 4:1.361 | דמהכא |  |
| 4:1.362 | משמע |  |
| 4:1.363 | שהקב"ה |  |
| 4:1.364 | בעצמו |  |
| 4:1.365 | כביכול |  |
| 4:1.366 | לא |  |
| 4:1.367 | הסכים |  |
| 4:1.368 | בבריאה |  |
| 4:1.369 | אדם |  |
| 4:1.370 | רק |  |
| 4:1.371 | השכינה |  |
| 4:1.372 | התחננה |  |
| 4:1.373 | להיות |  |
| 4:1.374 | אדם |  |
| 4:1.375 | בגורלה |  |
| 4:1.376 | ע"כ | , |
| 4:1.377 | ופי | ' |
| 4:1.378 | הר | ' |
| 4:1.379 | מקדש |  |
| 4:1.380 | מלך |  |
| 4:1.381 | בן |  |
| 4:1.382 | חכם |  |
| 4:1.383 | דא |  |
| 4:1.384 | אדם |  |
| 4:1.385 | דאצילות | , |
| 4:1.386 | הוא |  |
| 4:1.387 | זעיר |  |
| 4:1.388 | דאצילות | , |
| 4:1.389 | אדם |  |
| 4:1.390 | דבריאה |  |
| 4:1.391 | הוא |  |
| 4:1.392 | אדם |  |
| 4:1.393 | הראשון |  |
| 4:1.394 | וע"כ | . |
| 4:1.395 | ולעיל |  |
| 4:1.396 | מזה |  |
| 4:1.397 | כ | ' |
| 4:1.398 | הר | ' |
| 4:1.399 | מקדש |  |
| 4:1.400 | מלך |  |
| 4:1.401 | בדבור |  |
| 4:1.402 | המתחיל |  |
| 4:1.403 | אמר |  |
| 4:1.404 | אומנא |  |
| 4:1.405 | פי | ' |
| 4:1.406 | אימא |  |
| 4:1.407 | אמרה |  |
| 4:1.408 | לאבא |  |
| 4:1.409 | נעשה |  |
| 4:1.410 | אדם | , |
| 4:1.411 | וזהו |  |
| 4:1.412 | ויאמר |  |
| 4:1.413 | אלהים |  |
| 4:1.414 | שהיא |  |
| 4:1.415 | אימא |  |
| 4:1.416 | אמרה |  |
| 4:1.417 | לאבא |  |
| 4:1.418 | נעשה |  |
| 4:1.419 | אדם |  |
| 4:1.420 | מה |  |
| 4:1.421 | שאין |  |
| 4:1.422 | כן |  |
| 4:1.423 | כל |  |
| 4:1.424 | ויאמר |  |
| 4:1.425 | דמעשה |  |
| 4:1.426 | בראשית |  |
| 4:1.427 | שרוצה |  |
| 4:1.428 | לומר |  |
| 4:1.429 | ויאמר |  |
| 4:1.430 | אבא |  |
| 4:1.431 | לאלהים |  |
| 4:1.432 | שהיא |  |
| 4:1.433 | אימא |  |
| 4:1.434 | בינה | ) |
| 4:1.435 | כי |  |
| 4:1.436 | אבא |  |
| 4:1.437 | אומר |  |
| 4:1.438 | ובינה |  |
| 4:1.439 | עושה |  |
| 4:1.440 | ועכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

