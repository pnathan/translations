# Translation pack: milhamot-hashem perek-4-10 (Milchamot ha-Shem 4:16–4:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-10`.

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
- **ז"ל** → "of blessed memory" (honorific).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- הסבה **הראשונה** → "the First Cause".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **ריבה / לרבות** → "include", never "multiply".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 4:16

Hebrew (Wikisource): וגם על רבי עקיבה שהוא רבו של רשב"י וגם הוא הגדול מד' שנכנסו לפרדס והגיע למקום אבני שיש טהור, ונכנס בשלום ויצא בשלום ייפלא למה ריבה תלמידי חכמים שהם בשר ודם והשוה מוראם למורא שמים, ורבי' ג"כ קבעה במשנה מסכת אבות פ' ד' בשם ר"א היה לו לרבות אשתו של זעיר אנפין שהוא ה' אלהינו לפי דבריהם? או יותר היה לו לרבות אבותיו של זעיר אנפין שהם אבא ואמא?

| id | word | punctuation after |
|---|---|---|
| 4:16.1 | וגם |  |
| 4:16.2 | על |  |
| 4:16.3 | רבי |  |
| 4:16.4 | עקיבה |  |
| 4:16.5 | שהוא |  |
| 4:16.6 | רבו |  |
| 4:16.7 | של |  |
| 4:16.8 | רשב"י |  |
| 4:16.9 | וגם |  |
| 4:16.10 | הוא |  |
| 4:16.11 | הגדול |  |
| 4:16.12 | מד | ' |
| 4:16.13 | שנכנסו |  |
| 4:16.14 | לפרדס |  |
| 4:16.15 | והגיע |  |
| 4:16.16 | למקום |  |
| 4:16.17 | אבני |  |
| 4:16.18 | שיש |  |
| 4:16.19 | טהור | , |
| 4:16.20 | ונכנס |  |
| 4:16.21 | בשלום |  |
| 4:16.22 | ויצא |  |
| 4:16.23 | בשלום |  |
| 4:16.24 | ייפלא |  |
| 4:16.25 | למה |  |
| 4:16.26 | ריבה |  |
| 4:16.27 | תלמידי |  |
| 4:16.28 | חכמים |  |
| 4:16.29 | שהם |  |
| 4:16.30 | בשר |  |
| 4:16.31 | ודם |  |
| 4:16.32 | והשוה |  |
| 4:16.33 | מוראם |  |
| 4:16.34 | למורא |  |
| 4:16.35 | שמים | , |
| 4:16.36 | ורבי | ' |
| 4:16.37 | ג"כ |  |
| 4:16.38 | קבעה |  |
| 4:16.39 | במשנה |  |
| 4:16.40 | מסכת |  |
| 4:16.41 | אבות |  |
| 4:16.42 | פ | ' |
| 4:16.43 | ד | ' |
| 4:16.44 | בשם |  |
| 4:16.45 | ר"א |  |
| 4:16.46 | היה |  |
| 4:16.47 | לו |  |
| 4:16.48 | לרבות |  |
| 4:16.49 | אשתו |  |
| 4:16.50 | של |  |
| 4:16.51 | זעיר |  |
| 4:16.52 | אנפין |  |
| 4:16.53 | שהוא |  |
| 4:16.54 | ה | ' |
| 4:16.55 | אלהינו |  |
| 4:16.56 | לפי |  |
| 4:16.57 | דבריהם | ? |
| 4:16.58 | או |  |
| 4:16.59 | יותר |  |
| 4:16.60 | היה |  |
| 4:16.61 | לו |  |
| 4:16.62 | לרבות |  |
| 4:16.63 | אבותיו |  |
| 4:16.64 | של |  |
| 4:16.65 | זעיר |  |
| 4:16.66 | אנפין |  |
| 4:16.67 | שהם |  |
| 4:16.68 | אבא |  |
| 4:16.69 | ואמא | ? |


## Milchamot ha-Shem 4:17

Hebrew (Wikisource): כט) ובספר הברית חלק א' מאמר כ' פרק ט"ו כתב בעל משנת חסידים בספר יושר לבב, כי שם הקב"ה המורגל בפינו. ושבשם הוי'ה אנו קורים אותו, הוא פרצוף זעיר אנפין שהוא הקב"ה, שנשמתו המסתתרת בו על ידי הפרצופים שלפנים ממנו, הוא הסבה הראשונה ממש, ולה אנו עובדים בעבודתו, הכלל העולה כי הסבה הראשונה הנקרא בפי כל המקובלים אין סוף, הוא אשר ברא הכל יש מאין והוא מקיף את היש ההוא כולו מבחוץ, ובמובחר שבו והוא הבחור שבאבות הנקרא זעיר אנפין נסתר חדר לפנים מחדר כנשמה בתוך הגוף ומחייהו. ולכך המבחר הזה והוא הזעיר אנפין השולט על כל הברואים ומנהיגם וזנם ומפרנסם, והוא אלהינו ואנחנו עמו, כי נשמותינו הם חלקו, והוא הוא המהלל בכל התשבחות הנאמרים בתורתנו הקדושה אשר הנחילנו ועל ידה כל הנסתרות גלה לנו עכ"ל. ועוד האריך בראיות מהזהר והתיקונים שהעבודה לזעיר אנפין שהוא אמצע המקשר את הכל. ועוד כתב בשם האר"י בלקוטי תנ"ך על הפסוק בטחו בה' עדי עד כי ביה ה' צור עולמים וז"ל. נמצא כשיכוין האדם בזעיר אנפין לבד מספיק, כי שם נמצא אריך אנפין ואבא ואמא מתלבשין זה תוך זה תוך זעיר אנפין עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:17.1 | כט | ) |
| 4:17.2 | ובספר |  |
| 4:17.3 | הברית |  |
| 4:17.4 | חלק |  |
| 4:17.5 | א | ' |
| 4:17.6 | מאמר |  |
| 4:17.7 | כ | ' |
| 4:17.8 | פרק |  |
| 4:17.9 | ט"ו |  |
| 4:17.10 | כתב |  |
| 4:17.11 | בעל |  |
| 4:17.12 | משנת |  |
| 4:17.13 | חסידים |  |
| 4:17.14 | בספר |  |
| 4:17.15 | יושר |  |
| 4:17.16 | לבב | , |
| 4:17.17 | כי |  |
| 4:17.18 | שם |  |
| 4:17.19 | הקב"ה |  |
| 4:17.20 | המורגל |  |
| 4:17.21 | בפינו | . |
| 4:17.22 | ושבשם |  |
| 4:17.23 | הוי'ה |  |
| 4:17.24 | אנו |  |
| 4:17.25 | קורים |  |
| 4:17.26 | אותו | , |
| 4:17.27 | הוא |  |
| 4:17.28 | פרצוף |  |
| 4:17.29 | זעיר |  |
| 4:17.30 | אנפין |  |
| 4:17.31 | שהוא |  |
| 4:17.32 | הקב"ה | , |
| 4:17.33 | שנשמתו |  |
| 4:17.34 | המסתתרת |  |
| 4:17.35 | בו |  |
| 4:17.36 | על |  |
| 4:17.37 | ידי |  |
| 4:17.38 | הפרצופים |  |
| 4:17.39 | שלפנים |  |
| 4:17.40 | ממנו | , |
| 4:17.41 | הוא |  |
| 4:17.42 | הסבה |  |
| 4:17.43 | הראשונה |  |
| 4:17.44 | ממש | , |
| 4:17.45 | ולה |  |
| 4:17.46 | אנו |  |
| 4:17.47 | עובדים |  |
| 4:17.48 | בעבודתו | , |
| 4:17.49 | הכלל |  |
| 4:17.50 | העולה |  |
| 4:17.51 | כי |  |
| 4:17.52 | הסבה |  |
| 4:17.53 | הראשונה |  |
| 4:17.54 | הנקרא |  |
| 4:17.55 | בפי |  |
| 4:17.56 | כל |  |
| 4:17.57 | המקובלים |  |
| 4:17.58 | אין |  |
| 4:17.59 | סוף | , |
| 4:17.60 | הוא |  |
| 4:17.61 | אשר |  |
| 4:17.62 | ברא |  |
| 4:17.63 | הכל |  |
| 4:17.64 | יש |  |
| 4:17.65 | מאין |  |
| 4:17.66 | והוא |  |
| 4:17.67 | מקיף |  |
| 4:17.68 | את |  |
| 4:17.69 | היש |  |
| 4:17.70 | ההוא |  |
| 4:17.71 | כולו |  |
| 4:17.72 | מבחוץ | , |
| 4:17.73 | ובמובחר |  |
| 4:17.74 | שבו |  |
| 4:17.75 | והוא |  |
| 4:17.76 | הבחור |  |
| 4:17.77 | שבאבות |  |
| 4:17.78 | הנקרא |  |
| 4:17.79 | זעיר |  |
| 4:17.80 | אנפין |  |
| 4:17.81 | נסתר |  |
| 4:17.82 | חדר |  |
| 4:17.83 | לפנים |  |
| 4:17.84 | מחדר |  |
| 4:17.85 | כנשמה |  |
| 4:17.86 | בתוך |  |
| 4:17.87 | הגוף |  |
| 4:17.88 | ומחייהו | . |
| 4:17.89 | ולכך |  |
| 4:17.90 | המבחר |  |
| 4:17.91 | הזה |  |
| 4:17.92 | והוא |  |
| 4:17.93 | הזעיר |  |
| 4:17.94 | אנפין |  |
| 4:17.95 | השולט |  |
| 4:17.96 | על |  |
| 4:17.97 | כל |  |
| 4:17.98 | הברואים |  |
| 4:17.99 | ומנהיגם |  |
| 4:17.100 | וזנם |  |
| 4:17.101 | ומפרנסם | , |
| 4:17.102 | והוא |  |
| 4:17.103 | אלהינו |  |
| 4:17.104 | ואנחנו |  |
| 4:17.105 | עמו | , |
| 4:17.106 | כי |  |
| 4:17.107 | נשמותינו |  |
| 4:17.108 | הם |  |
| 4:17.109 | חלקו | , |
| 4:17.110 | והוא |  |
| 4:17.111 | הוא |  |
| 4:17.112 | המהלל |  |
| 4:17.113 | בכל |  |
| 4:17.114 | התשבחות |  |
| 4:17.115 | הנאמרים |  |
| 4:17.116 | בתורתנו |  |
| 4:17.117 | הקדושה |  |
| 4:17.118 | אשר |  |
| 4:17.119 | הנחילנו |  |
| 4:17.120 | ועל |  |
| 4:17.121 | ידה |  |
| 4:17.122 | כל |  |
| 4:17.123 | הנסתרות |  |
| 4:17.124 | גלה |  |
| 4:17.125 | לנו |  |
| 4:17.126 | עכ"ל | . |
| 4:17.127 | ועוד |  |
| 4:17.128 | האריך |  |
| 4:17.129 | בראיות |  |
| 4:17.130 | מהזהר |  |
| 4:17.131 | והתיקונים |  |
| 4:17.132 | שהעבודה |  |
| 4:17.133 | לזעיר |  |
| 4:17.134 | אנפין |  |
| 4:17.135 | שהוא |  |
| 4:17.136 | אמצע |  |
| 4:17.137 | המקשר |  |
| 4:17.138 | את |  |
| 4:17.139 | הכל | . |
| 4:17.140 | ועוד |  |
| 4:17.141 | כתב |  |
| 4:17.142 | בשם |  |
| 4:17.143 | האר"י |  |
| 4:17.144 | בלקוטי |  |
| 4:17.145 | תנ"ך |  |
| 4:17.146 | על |  |
| 4:17.147 | הפסוק |  |
| 4:17.148 | בטחו |  |
| 4:17.149 | בה | ' |
| 4:17.150 | עדי |  |
| 4:17.151 | עד |  |
| 4:17.152 | כי |  |
| 4:17.153 | ביה |  |
| 4:17.154 | ה | ' |
| 4:17.155 | צור |  |
| 4:17.156 | עולמים |  |
| 4:17.157 | וז"ל | . |
| 4:17.158 | נמצא |  |
| 4:17.159 | כשיכוין |  |
| 4:17.160 | האדם |  |
| 4:17.161 | בזעיר |  |
| 4:17.162 | אנפין |  |
| 4:17.163 | לבד |  |
| 4:17.164 | מספיק | , |
| 4:17.165 | כי |  |
| 4:17.166 | שם |  |
| 4:17.167 | נמצא |  |
| 4:17.168 | אריך |  |
| 4:17.169 | אנפין |  |
| 4:17.170 | ואבא |  |
| 4:17.171 | ואמא |  |
| 4:17.172 | מתלבשין |  |
| 4:17.173 | זה |  |
| 4:17.174 | תוך |  |
| 4:17.175 | זה |  |
| 4:17.176 | תוך |  |
| 4:17.177 | זעיר |  |
| 4:17.178 | אנפין |  |
| 4:17.179 | עכ"ל | . |


## Milchamot ha-Shem 4:18

Hebrew (Wikisource): גם בספר עץ חיים להרח"ו שער הכללים פ' י"א כתב שמשה רבינו אמר לישראל בכניסתן לארץ ואתם הדבקים בה' אלהיכם חיים כולכם היום, כי ה' אלהיכם הוא זעיר אנפין ונוקביה, וישראל הנכנסים לארץ היו וכו' עי"ש. גם בספר הכוונות במנהגי האר"י כתב כי כל פעם שתוציא מפיך שם הוי'ה תכוין לזעיר אנפין, ותכוין בו כך יה"ו זעיר אנפין וה"א אחרונה שבו היא מלכות. תכוין שנעשה שם שלם על אדני באופן וכו' ע"ש. וכן בספר מחברת הקודש בסדר מוסף שבת כ' כי מלאכים המוני מעלה יתנו כתר לזעיר אנפין שהוא ה' אלהינו עיין שם דף מ' ע"א. גם המקדש מלך ד' י"ב ד"ה אלהא רברבא מבואר דלא נקרא אלהא רברבא אלא בזעיר אנפין, ובחכמה נקרא חכם וכו' והמכוין לאין סוף מחמת רוממותו שאין לו שם ולא נקודה שתגבילהו אין תפלתו תפלה, רק יכוין אליו בהיותו מלובש במדותיו וכו' ועיין לקמן סי' ק' בשם נחלת יוסף.

| id | word | punctuation after |
|---|---|---|
| 4:18.1 | גם |  |
| 4:18.2 | בספר |  |
| 4:18.3 | עץ |  |
| 4:18.4 | חיים |  |
| 4:18.5 | להרח"ו |  |
| 4:18.6 | שער |  |
| 4:18.7 | הכללים |  |
| 4:18.8 | פ | ' |
| 4:18.9 | י"א |  |
| 4:18.10 | כתב |  |
| 4:18.11 | שמשה |  |
| 4:18.12 | רבינו |  |
| 4:18.13 | אמר |  |
| 4:18.14 | לישראל |  |
| 4:18.15 | בכניסתן |  |
| 4:18.16 | לארץ |  |
| 4:18.17 | ואתם |  |
| 4:18.18 | הדבקים |  |
| 4:18.19 | בה | ' |
| 4:18.20 | אלהיכם |  |
| 4:18.21 | חיים |  |
| 4:18.22 | כולכם |  |
| 4:18.23 | היום | , |
| 4:18.24 | כי |  |
| 4:18.25 | ה | ' |
| 4:18.26 | אלהיכם |  |
| 4:18.27 | הוא |  |
| 4:18.28 | זעיר |  |
| 4:18.29 | אנפין |  |
| 4:18.30 | ונוקביה | , |
| 4:18.31 | וישראל |  |
| 4:18.32 | הנכנסים |  |
| 4:18.33 | לארץ |  |
| 4:18.34 | היו |  |
| 4:18.35 | וכו | ' |
| 4:18.36 | עי"ש | . |
| 4:18.37 | גם |  |
| 4:18.38 | בספר |  |
| 4:18.39 | הכוונות |  |
| 4:18.40 | במנהגי |  |
| 4:18.41 | האר"י |  |
| 4:18.42 | כתב |  |
| 4:18.43 | כי |  |
| 4:18.44 | כל |  |
| 4:18.45 | פעם |  |
| 4:18.46 | שתוציא |  |
| 4:18.47 | מפיך |  |
| 4:18.48 | שם |  |
| 4:18.49 | הוי'ה |  |
| 4:18.50 | תכוין |  |
| 4:18.51 | לזעיר |  |
| 4:18.52 | אנפין | , |
| 4:18.53 | ותכוין |  |
| 4:18.54 | בו |  |
| 4:18.55 | כך |  |
| 4:18.56 | יה"ו |  |
| 4:18.57 | זעיר |  |
| 4:18.58 | אנפין |  |
| 4:18.59 | וה"א |  |
| 4:18.60 | אחרונה |  |
| 4:18.61 | שבו |  |
| 4:18.62 | היא |  |
| 4:18.63 | מלכות | . |
| 4:18.64 | תכוין |  |
| 4:18.65 | שנעשה |  |
| 4:18.66 | שם |  |
| 4:18.67 | שלם |  |
| 4:18.68 | על |  |
| 4:18.69 | אדני |  |
| 4:18.70 | באופן |  |
| 4:18.71 | וכו | ' |
| 4:18.72 | ע"ש | . |
| 4:18.73 | וכן |  |
| 4:18.74 | בספר |  |
| 4:18.75 | מחברת |  |
| 4:18.76 | הקודש |  |
| 4:18.77 | בסדר |  |
| 4:18.78 | מוסף |  |
| 4:18.79 | שבת |  |
| 4:18.80 | כ | ' |
| 4:18.81 | כי |  |
| 4:18.82 | מלאכים |  |
| 4:18.83 | המוני |  |
| 4:18.84 | מעלה |  |
| 4:18.85 | יתנו |  |
| 4:18.86 | כתר |  |
| 4:18.87 | לזעיר |  |
| 4:18.88 | אנפין |  |
| 4:18.89 | שהוא |  |
| 4:18.90 | ה | ' |
| 4:18.91 | אלהינו |  |
| 4:18.92 | עיין |  |
| 4:18.93 | שם |  |
| 4:18.94 | דף |  |
| 4:18.95 | מ | ' |
| 4:18.96 | ע"א | . |
| 4:18.97 | גם |  |
| 4:18.98 | המקדש |  |
| 4:18.99 | מלך |  |
| 4:18.100 | ד | ' |
| 4:18.101 | י"ב |  |
| 4:18.102 | ד"ה |  |
| 4:18.103 | אלהא |  |
| 4:18.104 | רברבא |  |
| 4:18.105 | מבואר |  |
| 4:18.106 | דלא |  |
| 4:18.107 | נקרא |  |
| 4:18.108 | אלהא |  |
| 4:18.109 | רברבא |  |
| 4:18.110 | אלא |  |
| 4:18.111 | בזעיר |  |
| 4:18.112 | אנפין | , |
| 4:18.113 | ובחכמה |  |
| 4:18.114 | נקרא |  |
| 4:18.115 | חכם |  |
| 4:18.116 | וכו | ' |
| 4:18.117 | והמכוין |  |
| 4:18.118 | לאין |  |
| 4:18.119 | סוף |  |
| 4:18.120 | מחמת |  |
| 4:18.121 | רוממותו |  |
| 4:18.122 | שאין |  |
| 4:18.123 | לו |  |
| 4:18.124 | שם |  |
| 4:18.125 | ולא |  |
| 4:18.126 | נקודה |  |
| 4:18.127 | שתגבילהו |  |
| 4:18.128 | אין |  |
| 4:18.129 | תפלתו |  |
| 4:18.130 | תפלה | , |
| 4:18.131 | רק |  |
| 4:18.132 | יכוין |  |
| 4:18.133 | אליו |  |
| 4:18.134 | בהיותו |  |
| 4:18.135 | מלובש |  |
| 4:18.136 | במדותיו |  |
| 4:18.137 | וכו | ' |
| 4:18.138 | ועיין |  |
| 4:18.139 | לקמן |  |
| 4:18.140 | סי | ' |
| 4:18.141 | ק | ' |
| 4:18.142 | בשם |  |
| 4:18.143 | נחלת |  |
| 4:18.144 | יוסף | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

