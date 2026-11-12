# Review pack: milhamot-hashem perek-2-02 (Milchamot ha-Shem 2:5–2:7)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-02`.

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
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **רחמים**: "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes
  of mercy", a different thirteen from the hermeneutical rules.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
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
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- אורות **דקים** → "subtle lights".
- לית מחשבה **תפיסה** ביה → "no thought grasps Him".
- כופר **בעיקר** → "denies the fundamental principle".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **א"כ** → "if so".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין** אל → "call upon".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ווי** ליה מאן דלא → "woe to him who does not".
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 2:5

Hebrew (Wikisource): ואיך נניח כל האלוהות אשר למעלה מעולם דאצילות ונעבוד דוקא לזעיר ונוקביה שבעולם האצילות שכבר קדמו לו כמה עתיקין ואריכי אנפין ואו"א וזעיר ונוקביה בעולמות העליונים אין מספר הקרובים לאין סוף יותר?

| id | word | punctuation after |
|---|---|---|
| 2:5.1 | ואיך |  |
| 2:5.2 | נניח |  |
| 2:5.3 | כל |  |
| 2:5.4 | האלוהות |  |
| 2:5.5 | אשר |  |
| 2:5.6 | למעלה |  |
| 2:5.7 | מעולם |  |
| 2:5.8 | דאצילות |  |
| 2:5.9 | ונעבוד |  |
| 2:5.10 | דוקא |  |
| 2:5.11 | לזעיר |  |
| 2:5.12 | ונוקביה |  |
| 2:5.13 | שבעולם |  |
| 2:5.14 | האצילות |  |
| 2:5.15 | שכבר |  |
| 2:5.16 | קדמו |  |
| 2:5.17 | לו |  |
| 2:5.18 | כמה |  |
| 2:5.19 | עתיקין |  |
| 2:5.20 | ואריכי |  |
| 2:5.21 | אנפין |  |
| 2:5.22 | ואו"א |  |
| 2:5.23 | וזעיר |  |
| 2:5.24 | ונוקביה |  |
| 2:5.25 | בעולמות |  |
| 2:5.26 | העליונים |  |
| 2:5.27 | אין |  |
| 2:5.28 | מספר |  |
| 2:5.29 | הקרובים |  |
| 2:5.30 | לאין |  |
| 2:5.31 | סוף |  |
| 2:5.32 | יותר | ? |


### Translation under review (2:5)

English: And how can we leave all the divinities that are above the World of Emanation, and worship specifically Ze'ir and his Nukva, who are in the World of Emanation, before whom have already preceded several Atikin and Arikhei Anpin, and Abba and Imma, and Ze'ir and his Nukva in the upper worlds; there is no number of those closer to Ein Sof, and more?

Literal: and-how shall-we-leave all the-divinities that above the-World of-Emanation and-worship specifically Ze'ir and-his-Nukva who-are-in-the World-of-Emanation who-already preceded him several Atikin and-Arikhei Anpin and-Abba-and-Imma and-Ze'ir and-his-Nukva in-the-worlds upper without number the-close-ones to-Ein Sof more

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:5.1 | and how |  |  |  |
| 2:5.2 | shall we leave |  |  |  |
| 2:5.3 | all |  |  |  |
| 2:5.4 | the divinities |  |  |  |
| 2:5.5 | that |  |  |  |
| 2:5.6 | above |  |  |  |
| 2:5.7 | the World |  |  |  |
| 2:5.8 | of Emanation |  |  |  |
| 2:5.9 | and worship |  |  |  |
| 2:5.10 | specifically |  |  |  |
| 2:5.11 | Ze'ir |  |  | Ze'ir Anpin, 'Small Face' (Aramaic); here the bare name Ze'ir. |
| 2:5.12 | and his Nukva |  |  | Nukva: the female partzuf paired with Ze'ir; 'his Nukva' carries the Hebrew possessive suffix. |
| 2:5.13 | who are in the |  |  |  |
| 2:5.14 | World of Emanation |  |  |  |
| 2:5.15 | who already |  |  |  |
| 2:5.16 | preceded |  |  |  |
| 2:5.17 | him |  |  |  |
| 2:5.18 | several |  |  |  |
| 2:5.19 | Atikin |  |  | Atik, plural Atikin: partzufim named for the Ancient One. Atika Kadisha is the Holy Ancient One. |
| 2:5.20 | and Arikhei |  |  | Arikh Anpin, plural Arikhei Anpin: 'Long Face'. |
| 2:5.21 | Anpin |  |  |  |
| 2:5.22 | and Abba and Imma |  |  | Abba (Father) and Imma (Mother), written with the abbreviation ואו"א. |
| 2:5.23 | and Ze'ir |  |  |  |
| 2:5.24 | and his Nukva |  |  |  |
| 2:5.25 | in the worlds |  |  |  |
| 2:5.26 | upper |  |  |  |
| 2:5.27 | without |  |  |  |
| 2:5.28 | number |  |  |  |
| 2:5.29 | the close ones |  |  |  |
| 2:5.30 | to Ein |  |  | Ein Sof, literally 'without end'. |
| 2:5.31 | Sof |  |  |  |
| 2:5.32 | more |  |  |  |

Phrases:
- 2:5.1 | And how | and-how
- 2:5.2 | can we leave | shall-we-leave
- 2:5.3 2:5.4 2:5.5 2:5.6 2:5.7 2:5.8 | all the divinities that are above the World of Emanation, | all the-divinities that above the-World of-Emanation
- 2:5.9 2:5.10 | and worship specifically | and-worship specifically
- 2:5.11 2:5.12 | Ze'ir and his Nukva, | Ze'ir and-his-Nukva
- 2:5.13 2:5.14 | who are in the World of Emanation, | who-are-in-the World-of-Emanation
- 2:5.15 2:5.16 2:5.17 | before whom have already preceded | who-already preceded him
- 2:5.18 2:5.19 2:5.20 2:5.21 | several Atikin and Arikhei Anpin, | several Atikin and-Arikhei Anpin
- 2:5.22 | and Abba and Imma, | and-Abba-and-Imma
- 2:5.23 2:5.24 | and Ze'ir and his Nukva | and-Ze'ir and-his-Nukva
- 2:5.25 2:5.26 | in the upper worlds; | in-the-worlds upper
- 2:5.27 2:5.28 2:5.29 2:5.30 2:5.31 | there is no number of those closer to Ein Sof, | without number the-close-ones to-Ein Sof
- 2:5.32 | and more? | more

Grammar notes:
- נניח and ונעבוד are first-person plural cohortatives in a question ('how shall we leave ... and worship'), so the English reads 'can we leave'.
- קדמו לו כמה עתיקין: the verb comes before its subject (verb-subject order), and לו means 'before him'.
- The last clause, אין מספר הקרובים לאין סוף יותר?, is elliptic ('there is no number of those closer to Ein Sof, and more?'); the question mark is the Hebrew's.

Translator's note: The questioner's letter, in his voice. Ze'ir and his Nukva are named as the partzufim the question asks us to worship.

Commentary rulings:
- apparatus: Wikisource base text; not yet collated with the 1931 scan (decisions, Q2).
- apparatus: ואו"א is one abbreviated word for Abba and Imma.

## Milchamot ha-Shem 2:6

Hebrew (Wikisource): ורבותינו הראשונים אמרו שאין ראוי לעבוד ולהודות ולגדל ולפאר רק לאלוה הראשון לכל הנמצאים מבלעדיו! ולמה נניח האין סוף העיקרי המצטמצם למעלה מכולם עם כל הפרצופים הקרובים אליו, ונעבוד לזה הרחוק ממנו מאוד שהוא פרצוף אחרון שבעולם האצילות? את כל אלה הביננו אדונינו נלכה נרדפה את ה' יוצרינו יוצר בראשית. ואל תצוק בנו דבר בליעל ח"ו. כי מי אסף חן ושכל טוב בחפניו השכל וידוע את ה' בחקירת העניינים האלהיים בעיקרי האמונה כמוך. מי יעלה בהר ה' ומי יקום במקום קדשו כמוך, כי בער אנכי מאיש ולא בינת אדם לי. וצור יודע תעלומות לב, בוחן לבבות כי לא נתתי שינה לעיני ונגזלה שנתי מעיני עד אמצא פתרונים והבא לטהר מסייעין אותו, להמציא לו מורה דרך הקדש בשכל ומדע ככלכל ודרדע, ובין חכמי האמת חכם ואומן, כאיתן והימן, שאולי הרבנים החסרים מן החכמה הזאת לנו משאות שוא ומדוחים ח"ו. ואתה אדני לך נגלו תעלומותיה וצפונך תמלא בטנם של תלמידיך, ישבע ויתענגו מטובך, ועלינו לשלם שכר פעולתך. ואתה אל תמנע טוב מבעליו. והנני משביעך באלהי השמים האר פניך אלינו לעות את יעף דבר דבור על אפניו והיינו עבדים לאדוני. ושכר רב טובך תבענה שפתינו, ואז לא נבוש לא בעה"ז ולא בעה"ב ונשובה נרעה בשושנים מסוד חכמים ונבונים.

| id | word | punctuation after |
|---|---|---|
| 2:6.1 | ורבותינו |  |
| 2:6.2 | הראשונים |  |
| 2:6.3 | אמרו |  |
| 2:6.4 | שאין |  |
| 2:6.5 | ראוי |  |
| 2:6.6 | לעבוד |  |
| 2:6.7 | ולהודות |  |
| 2:6.8 | ולגדל |  |
| 2:6.9 | ולפאר |  |
| 2:6.10 | רק |  |
| 2:6.11 | לאלוה |  |
| 2:6.12 | הראשון |  |
| 2:6.13 | לכל |  |
| 2:6.14 | הנמצאים |  |
| 2:6.15 | מבלעדיו | ! |
| 2:6.16 | ולמה |  |
| 2:6.17 | נניח |  |
| 2:6.18 | האין |  |
| 2:6.19 | סוף |  |
| 2:6.20 | העיקרי |  |
| 2:6.21 | המצטמצם |  |
| 2:6.22 | למעלה |  |
| 2:6.23 | מכולם |  |
| 2:6.24 | עם |  |
| 2:6.25 | כל |  |
| 2:6.26 | הפרצופים |  |
| 2:6.27 | הקרובים |  |
| 2:6.28 | אליו | , |
| 2:6.29 | ונעבוד |  |
| 2:6.30 | לזה |  |
| 2:6.31 | הרחוק |  |
| 2:6.32 | ממנו |  |
| 2:6.33 | מאוד |  |
| 2:6.34 | שהוא |  |
| 2:6.35 | פרצוף |  |
| 2:6.36 | אחרון |  |
| 2:6.37 | שבעולם |  |
| 2:6.38 | האצילות | ? |
| 2:6.39 | את |  |
| 2:6.40 | כל |  |
| 2:6.41 | אלה |  |
| 2:6.42 | הביננו |  |
| 2:6.43 | אדונינו |  |
| 2:6.44 | נלכה |  |
| 2:6.45 | נרדפה |  |
| 2:6.46 | את |  |
| 2:6.47 | ה | ' |
| 2:6.48 | יוצרינו |  |
| 2:6.49 | יוצר |  |
| 2:6.50 | בראשית | . |
| 2:6.51 | ואל |  |
| 2:6.52 | תצוק |  |
| 2:6.53 | בנו |  |
| 2:6.54 | דבר |  |
| 2:6.55 | בליעל |  |
| 2:6.56 | ח"ו | . |
| 2:6.57 | כי |  |
| 2:6.58 | מי |  |
| 2:6.59 | אסף |  |
| 2:6.60 | חן |  |
| 2:6.61 | ושכל |  |
| 2:6.62 | טוב |  |
| 2:6.63 | בחפניו |  |
| 2:6.64 | השכל |  |
| 2:6.65 | וידוע |  |
| 2:6.66 | את |  |
| 2:6.67 | ה | ' |
| 2:6.68 | בחקירת |  |
| 2:6.69 | העניינים |  |
| 2:6.70 | האלהיים |  |
| 2:6.71 | בעיקרי |  |
| 2:6.72 | האמונה |  |
| 2:6.73 | כמוך | . |
| 2:6.74 | מי |  |
| 2:6.75 | יעלה |  |
| 2:6.76 | בהר |  |
| 2:6.77 | ה | ' |
| 2:6.78 | ומי |  |
| 2:6.79 | יקום |  |
| 2:6.80 | במקום |  |
| 2:6.81 | קדשו |  |
| 2:6.82 | כמוך | , |
| 2:6.83 | כי |  |
| 2:6.84 | בער |  |
| 2:6.85 | אנכי |  |
| 2:6.86 | מאיש |  |
| 2:6.87 | ולא |  |
| 2:6.88 | בינת |  |
| 2:6.89 | אדם |  |
| 2:6.90 | לי | . |
| 2:6.91 | וצור |  |
| 2:6.92 | יודע |  |
| 2:6.93 | תעלומות |  |
| 2:6.94 | לב | , |
| 2:6.95 | בוחן |  |
| 2:6.96 | לבבות |  |
| 2:6.97 | כי |  |
| 2:6.98 | לא |  |
| 2:6.99 | נתתי |  |
| 2:6.100 | שינה |  |
| 2:6.101 | לעיני |  |
| 2:6.102 | ונגזלה |  |
| 2:6.103 | שנתי |  |
| 2:6.104 | מעיני |  |
| 2:6.105 | עד |  |
| 2:6.106 | אמצא |  |
| 2:6.107 | פתרונים |  |
| 2:6.108 | והבא |  |
| 2:6.109 | לטהר |  |
| 2:6.110 | מסייעין |  |
| 2:6.111 | אותו | , |
| 2:6.112 | להמציא |  |
| 2:6.113 | לו |  |
| 2:6.114 | מורה |  |
| 2:6.115 | דרך |  |
| 2:6.116 | הקדש |  |
| 2:6.117 | בשכל |  |
| 2:6.118 | ומדע |  |
| 2:6.119 | ככלכל |  |
| 2:6.120 | ודרדע | , |
| 2:6.121 | ובין |  |
| 2:6.122 | חכמי |  |
| 2:6.123 | האמת |  |
| 2:6.124 | חכם |  |
| 2:6.125 | ואומן | , |
| 2:6.126 | כאיתן |  |
| 2:6.127 | והימן | , |
| 2:6.128 | שאולי |  |
| 2:6.129 | הרבנים |  |
| 2:6.130 | החסרים |  |
| 2:6.131 | מן |  |
| 2:6.132 | החכמה |  |
| 2:6.133 | הזאת |  |
| 2:6.134 | לנו |  |
| 2:6.135 | משאות |  |
| 2:6.136 | שוא |  |
| 2:6.137 | ומדוחים |  |
| 2:6.138 | ח"ו | . |
| 2:6.139 | ואתה |  |
| 2:6.140 | אדני |  |
| 2:6.141 | לך |  |
| 2:6.142 | נגלו |  |
| 2:6.143 | תעלומותיה |  |
| 2:6.144 | וצפונך |  |
| 2:6.145 | תמלא |  |
| 2:6.146 | בטנם |  |
| 2:6.147 | של |  |
| 2:6.148 | תלמידיך | , |
| 2:6.149 | ישבע |  |
| 2:6.150 | ויתענגו |  |
| 2:6.151 | מטובך | , |
| 2:6.152 | ועלינו |  |
| 2:6.153 | לשלם |  |
| 2:6.154 | שכר |  |
| 2:6.155 | פעולתך | . |
| 2:6.156 | ואתה |  |
| 2:6.157 | אל |  |
| 2:6.158 | תמנע |  |
| 2:6.159 | טוב |  |
| 2:6.160 | מבעליו | . |
| 2:6.161 | והנני |  |
| 2:6.162 | משביעך |  |
| 2:6.163 | באלהי |  |
| 2:6.164 | השמים |  |
| 2:6.165 | האר |  |
| 2:6.166 | פניך |  |
| 2:6.167 | אלינו |  |
| 2:6.168 | לעות |  |
| 2:6.169 | את |  |
| 2:6.170 | יעף |  |
| 2:6.171 | דבר |  |
| 2:6.172 | דבור |  |
| 2:6.173 | על |  |
| 2:6.174 | אפניו |  |
| 2:6.175 | והיינו |  |
| 2:6.176 | עבדים |  |
| 2:6.177 | לאדוני | . |
| 2:6.178 | ושכר |  |
| 2:6.179 | רב |  |
| 2:6.180 | טובך |  |
| 2:6.181 | תבענה |  |
| 2:6.182 | שפתינו | , |
| 2:6.183 | ואז |  |
| 2:6.184 | לא |  |
| 2:6.185 | נבוש |  |
| 2:6.186 | לא |  |
| 2:6.187 | בעה"ז |  |
| 2:6.188 | ולא |  |
| 2:6.189 | בעה"ב |  |
| 2:6.190 | ונשובה |  |
| 2:6.191 | נרעה |  |
| 2:6.192 | בשושנים |  |
| 2:6.193 | מסוד |  |
| 2:6.194 | חכמים |  |
| 2:6.195 | ונבונים | . |


### Translation under review (2:6)

English: And our early rabbis said that it is not fitting to worship, to give thanks, to magnify, and to glorify, except to the first God of all that exists apart from Him! And why should we leave the principal Ein Sof, that contracted above all, with all the partzufim close to it, and worship that which is very far from it, which is the last partzuf of the World of Emanation? All this, make us understand, our master, let us go and pursue the LORD our Maker the Creator of the beginning. And do not press upon us a word of Belial, God forbid. For who has gathered grace and good understanding in his palms, the intelligent one who knows the LORD through investigating divine matters in the principles of the faith, like you. Who will ascend to the mountain of the LORD, and who will stand in His holy place, like you, For I am more brutish than any man, and I have no understanding of a man. And the Rock knows the secrets of the heart, the examiner of hearts, for I did not give sleep to my eyes, and my sleep was taken from my eyes until I find solutions and whoever comes to purify himself, they help him, to provide for him a guide to the way of holiness by intellect and knowledge, like Khalkhal and Darda, and among the sages of truth a sage and an artisan, like Ethan and Heman, perhaps the rabbis who lack this wisdom lay upon us vain burdens and rejected ones, God forbid. And you, my lord, to you its secrets are revealed, and your treasure will fill the bellies of your disciples, they will be sated and delight in your goodness, and upon us to pay the wages of your work. And you, do not withhold good from its owners. And here I adjure you by the God of heaven: make Your face shine upon us, to answer the weary a word spoken in its proper setting, and we shall be servants to our master. And our lips will ask a great reward for your goodness, then we shall not be ashamed, neither in this world nor in the world to come, and we shall return and graze among the lilies from the company of sages and discerning ones.

Literal: and-our-rabbis the-early said that-it-is-not fitting to-worship and-to-give-thanks and-to-magnify and-to-glorify except to-God the-first of-all the-existing-ones apart-from-Him and-why should-we-leave the-Ein Sof the-principal that-contracted above all-of-them with all the-partzufim the-close to-it and-worship this-one the-far from-it very which-is a-partzuf last of-the World-of-Emanation (object-marker) all these make-us-understand our-master let-us-go let-us-pursue (object-marker) the-LORD our-Maker Creator of-the-beginning and-do-not press upon-us a-word-of Belial God-forbid for who has-gathered grace and-understanding good in-his-palms the-intelligent-one and-who-knows (object-marker) the-LORD by-investigating the-matters divine in-the-principles-of the-faith like-you who will-ascend to-the-mountain-of the-LORD and-who will-stand in His-holy-place like-you for brutish I-am than-any-man and-not understanding-of a-man I-have and-the-Rock knows secrets-of the-heart examiner-of hearts for not I-gave sleep to-my-eyes and-was-taken my-sleep from-my-eyes until I-find solutions and-whoever-comes to-purify-himself they-help him to-provide for-him a-guide to-the-way-of holiness by-intellect and-knowledge like-Khalkhal and-Darda and-among the-sages-of truth a-sage and-an-artisan like-Ethan and-Heman perhaps the-rabbis who-lack of the-wisdom this for-us burdens vain rejected God-forbid and-you my-lord to-you are-revealed its-secrets and-your-treasure will-fill their-bellies of your-disciples will-be-sated and-will-delight in-your-goodness and-upon-us to-pay the-wage-of your-work and-you do-not withhold good from-its-owners and-here-I adjure-you by-the-God-of heaven make-shine Your-face upon-us to-answer (object-marker) the-weary a-word spoken in its-proper-setting and-we-shall-be servants to-our-master and-reward great for-your-goodness will-ask our-lips and-then not we-shall-be-ashamed not in-this-world and-not in-the-world-to-come and-we-shall-return and-graze among-the-lilies from-the-company-of sages and-discerning-ones

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:6.1 | and our rabbis |  |  |  |
| 2:6.2 | the early |  |  |  |
| 2:6.3 | said |  |  |  |
| 2:6.4 | that it is not |  |  |  |
| 2:6.5 | fitting |  |  |  |
| 2:6.6 | to worship |  |  |  |
| 2:6.7 | and to give thanks |  |  |  |
| 2:6.8 | and to magnify |  |  |  |
| 2:6.9 | and to glorify |  |  |  |
| 2:6.10 | except |  |  |  |
| 2:6.11 | to God |  |  |  |
| 2:6.12 | the first |  |  |  |
| 2:6.13 | of all |  |  |  |
| 2:6.14 | the existing ones |  |  |  |
| 2:6.15 | apart from Him |  |  |  |
| 2:6.16 | and why |  |  |  |
| 2:6.17 | should we leave |  |  |  |
| 2:6.18 | the Ein |  |  |  |
| 2:6.19 | Sof |  |  |  |
| 2:6.20 | the principal |  |  |  |
| 2:6.21 | that contracted |  |  | צמצם: 'contracted'. The tzimtzum (contraction) is the Kabbalists' doctrine; the English keeps the word fixed. |
| 2:6.22 | above |  |  |  |
| 2:6.23 | all of them |  |  |  |
| 2:6.24 | with |  |  |  |
| 2:6.25 | all |  |  |  |
| 2:6.26 | the partzufim |  |  |  |
| 2:6.27 | the close |  |  |  |
| 2:6.28 | to it |  |  |  |
| 2:6.29 | and worship |  |  |  |
| 2:6.30 | this one |  |  |  |
| 2:6.31 | the far |  |  |  |
| 2:6.32 | from it |  |  |  |
| 2:6.33 | very |  |  |  |
| 2:6.34 | which is |  |  |  |
| 2:6.35 | a partzuf |  |  | Partzuf (plural partzufim): a Kabbalistic entity; in the generic sense it means 'face'. |
| 2:6.36 | last |  |  |  |
| 2:6.37 | of the |  |  |  |
| 2:6.38 | World of Emanation |  |  |  |
| 2:6.39 | (object marker) |  |  |  |
| 2:6.40 | all |  |  |  |
| 2:6.41 | these |  |  |  |
| 2:6.42 | make us understand |  |  |  |
| 2:6.43 | our master |  |  |  |
| 2:6.44 | let us go |  |  |  |
| 2:6.45 | let us pursue |  |  |  |
| 2:6.46 | (object marker) |  |  |  |
| 2:6.47 | the LORD |  |  |  |
| 2:6.48 | our Maker |  |  |  |
| 2:6.49 | Creator |  |  |  |
| 2:6.50 | of the beginning |  |  | Literally 'creator of the beginning'. The Hebrew word is also the first word of Genesis. |
| 2:6.51 | and do not |  |  |  |
| 2:6.52 | press |  |  |  |
| 2:6.53 | upon us |  |  |  |
| 2:6.54 | a word of |  |  |  |
| 2:6.55 | Belial |  |  | Belial: a Biblical word for worthlessness, here 'a word of Belial'. |
| 2:6.56 | God forbid |  |  |  |
| 2:6.57 | for |  |  |  |
| 2:6.58 | who |  |  |  |
| 2:6.59 | has gathered |  |  |  |
| 2:6.60 | grace |  |  |  |
| 2:6.61 | and understanding |  |  |  |
| 2:6.62 | good |  |  |  |
| 2:6.63 | in his palms |  |  |  |
| 2:6.64 | the intelligent one |  |  |  |
| 2:6.65 | and who knows |  |  |  |
| 2:6.66 | (object marker) |  |  |  |
| 2:6.67 | the LORD |  |  |  |
| 2:6.68 | by investigating |  |  |  |
| 2:6.69 | the matters |  |  |  |
| 2:6.70 | divine |  |  |  |
| 2:6.71 | in the principles of |  |  |  |
| 2:6.72 | the faith |  |  |  |
| 2:6.73 | like you |  |  |  |
| 2:6.74 | who |  |  |  |
| 2:6.75 | will ascend |  |  |  |
| 2:6.76 | to the mountain of |  |  |  |
| 2:6.77 | the LORD |  |  |  |
| 2:6.78 | and who |  |  |  |
| 2:6.79 | will stand |  |  |  |
| 2:6.80 | in |  |  |  |
| 2:6.81 | His holy place |  |  |  |
| 2:6.82 | like you |  |  |  |
| 2:6.83 | for |  |  |  |
| 2:6.84 | brutish |  |  |  |
| 2:6.85 | I am |  |  |  |
| 2:6.86 | than any man |  |  |  |
| 2:6.87 | and not |  |  |  |
| 2:6.88 | understanding of |  |  |  |
| 2:6.89 | a man |  |  |  |
| 2:6.90 | I have |  |  |  |
| 2:6.91 | and the Rock |  |  |  |
| 2:6.92 | knows |  |  |  |
| 2:6.93 | secrets of |  |  |  |
| 2:6.94 | the heart |  |  |  |
| 2:6.95 | examiner of |  |  |  |
| 2:6.96 | hearts |  |  |  |
| 2:6.97 | for |  |  |  |
| 2:6.98 | not |  |  |  |
| 2:6.99 | I gave |  |  |  |
| 2:6.100 | sleep |  |  |  |
| 2:6.101 | to my eyes |  |  |  |
| 2:6.102 | and was taken |  |  |  |
| 2:6.103 | my sleep |  |  |  |
| 2:6.104 | from my eyes |  |  |  |
| 2:6.105 | until |  |  |  |
| 2:6.106 | I find |  |  |  |
| 2:6.107 | solutions |  |  |  |
| 2:6.108 | and whoever comes |  |  |  |
| 2:6.109 | to purify himself |  |  |  |
| 2:6.110 | they help |  |  |  |
| 2:6.111 | him |  |  |  |
| 2:6.112 | to provide |  |  |  |
| 2:6.113 | for him |  |  |  |
| 2:6.114 | a guide |  |  |  |
| 2:6.115 | to the way of |  |  |  |
| 2:6.116 | holiness |  |  |  |
| 2:6.117 | by intellect |  |  |  |
| 2:6.118 | and knowledge |  |  |  |
| 2:6.119 | like Khalkhal |  |  | Khalkhal: one of the four wise men named in 1 Kings 5:11. |
| 2:6.120 | and Darda |  |  |  |
| 2:6.121 | and among |  |  |  |
| 2:6.122 | the sages of |  |  |  |
| 2:6.123 | truth |  |  |  |
| 2:6.124 | a sage |  |  |  |
| 2:6.125 | and an artisan |  |  |  |
| 2:6.126 | like Ethan |  |  | Ethan: one of the four wise men named in 1 Kings 5:11. |
| 2:6.127 | and Heman |  |  | Heman: one of the four wise men named in 1 Kings 5:11. |
| 2:6.128 | perhaps |  |  |  |
| 2:6.129 | the rabbis |  |  |  |
| 2:6.130 | who lack |  |  |  |
| 2:6.131 | of |  |  |  |
| 2:6.132 | the wisdom |  |  |  |
| 2:6.133 | this |  |  |  |
| 2:6.134 | for us |  |  |  |
| 2:6.135 | burdens |  |  |  |
| 2:6.136 | vain |  |  |  |
| 2:6.137 | rejected |  |  |  |
| 2:6.138 | God forbid |  |  |  |
| 2:6.139 | and you |  |  |  |
| 2:6.140 | my lord |  |  |  |
| 2:6.141 | to you |  |  |  |
| 2:6.142 | are revealed |  |  |  |
| 2:6.143 | its secrets |  |  |  |
| 2:6.144 | and your treasure |  |  |  |
| 2:6.145 | will fill |  |  |  |
| 2:6.146 | their bellies |  |  |  |
| 2:6.147 | of |  |  |  |
| 2:6.148 | your disciples |  |  |  |
| 2:6.149 | will be sated |  |  |  |
| 2:6.150 | and will delight |  |  |  |
| 2:6.151 | in your goodness |  |  |  |
| 2:6.152 | and upon us |  |  |  |
| 2:6.153 | to pay |  |  |  |
| 2:6.154 | the wage of |  |  |  |
| 2:6.155 | your work |  |  |  |
| 2:6.156 | and you |  |  |  |
| 2:6.157 | do not |  |  |  |
| 2:6.158 | withhold |  |  |  |
| 2:6.159 | good |  |  |  |
| 2:6.160 | from its owners |  |  |  |
| 2:6.161 | and here I |  |  |  |
| 2:6.162 | adjure you |  |  |  |
| 2:6.163 | by the God of |  |  |  |
| 2:6.164 | heaven |  |  |  |
| 2:6.165 | make shine |  |  |  |
| 2:6.166 | Your face |  |  |  |
| 2:6.167 | upon us |  |  |  |
| 2:6.168 | to answer |  |  |  |
| 2:6.169 | (object marker) |  |  |  |
| 2:6.170 | the weary |  |  |  |
| 2:6.171 | a word |  |  |  |
| 2:6.172 | spoken |  |  |  |
| 2:6.173 | in |  |  |  |
| 2:6.174 | its proper setting |  |  | Literally 'its wheels' (אפניו, written as in the source). The idiom means 'its proper setting'. |
| 2:6.175 | and we shall be |  |  |  |
| 2:6.176 | servants |  |  |  |
| 2:6.177 | to our master |  |  |  |
| 2:6.178 | and reward |  |  |  |
| 2:6.179 | great |  |  |  |
| 2:6.180 | for your goodness |  |  |  |
| 2:6.181 | will ask |  |  |  |
| 2:6.182 | our lips |  |  |  |
| 2:6.183 | and then |  |  |  |
| 2:6.184 | not |  |  |  |
| 2:6.185 | we shall be ashamed |  |  |  |
| 2:6.186 | not |  |  |  |
| 2:6.187 | in this world |  |  |  |
| 2:6.188 | and not |  |  |  |
| 2:6.189 | in the world to come |  |  |  |
| 2:6.190 | and we shall return |  |  |  |
| 2:6.191 | and graze |  |  |  |
| 2:6.192 | among the lilies |  |  | Song 2:16: the one who grazes among the lilies. |
| 2:6.193 | from the company of |  |  |  |
| 2:6.194 | sages |  |  |  |
| 2:6.195 | and discerning ones |  |  |  |

Phrases:
- 2:6.1 2:6.2 2:6.3 | And our early rabbis said | and-our-rabbis the-early said
- 2:6.4 2:6.5 | that it is not fitting | that-it-is-not fitting
- 2:6.6 2:6.7 2:6.8 2:6.9 | to worship, to give thanks, to magnify, and to glorify, | to-worship and-to-give-thanks and-to-magnify and-to-glorify
- 2:6.10 2:6.11 2:6.12 | except to the first God | except to-God the-first
- 2:6.13 2:6.14 2:6.15 | of all that exists apart from Him! | of-all the-existing-ones apart-from-Him
- 2:6.16 2:6.17 | And why should we leave | and-why should-we-leave
- 2:6.18 2:6.19 2:6.20 | the principal Ein Sof, | the-Ein Sof the-principal
- 2:6.21 2:6.22 2:6.23 | that contracted above all, | that-contracted above all-of-them
- 2:6.24 2:6.25 2:6.26 2:6.27 2:6.28 | with all the partzufim close to it, | with all the-partzufim the-close to-it
- 2:6.29 2:6.30 2:6.31 2:6.32 2:6.33 | and worship that which is very far from it, | and-worship this-one the-far from-it very
- 2:6.34 2:6.35 2:6.36 2:6.37 2:6.38 | which is the last partzuf of the World of Emanation? | which-is a-partzuf last of-the World-of-Emanation
- 2:6.39 2:6.40 2:6.41 2:6.42 | All this, make us understand, | (object-marker) all these make-us-understand
- 2:6.43 | our master, | our-master
- 2:6.44 2:6.45 | let us go and pursue | let-us-go let-us-pursue
- 2:6.46 2:6.47 2:6.48 | the LORD our Maker | (object-marker) the-LORD our-Maker
- 2:6.49 2:6.50 | the Creator of the beginning. | Creator of-the-beginning
- 2:6.51 2:6.52 2:6.53 2:6.54 2:6.55 | And do not press upon us a word of Belial, | and-do-not press upon-us a-word-of Belial
- 2:6.56 | God forbid. | God-forbid
- 2:6.57 2:6.58 2:6.59 2:6.60 | For who has gathered grace | for who has-gathered grace
- 2:6.61 2:6.62 2:6.63 | and good understanding in his palms, | and-understanding good in-his-palms
- 2:6.64 2:6.65 | the intelligent one who knows | the-intelligent-one and-who-knows
- 2:6.66 2:6.67 2:6.68 2:6.69 2:6.70 | the LORD through investigating divine matters | (object-marker) the-LORD by-investigating the-matters divine
- 2:6.71 2:6.72 2:6.73 | in the principles of the faith, like you. | in-the-principles-of the-faith like-you
- 2:6.74 2:6.75 2:6.76 2:6.77 | Who will ascend to the mountain of the LORD, | who will-ascend to-the-mountain-of the-LORD
- 2:6.78 2:6.79 2:6.80 2:6.81 2:6.82 | and who will stand in His holy place, like you, | and-who will-stand in His-holy-place like-you
- 2:6.83 2:6.84 2:6.85 2:6.86 | For I am more brutish than any man, | for brutish I-am than-any-man
- 2:6.87 2:6.88 2:6.89 2:6.90 | and I have no understanding of a man. | and-not understanding-of a-man I-have
- 2:6.91 2:6.92 2:6.93 2:6.94 | And the Rock knows the secrets of the heart, | and-the-Rock knows secrets-of the-heart
- 2:6.95 2:6.96 | the examiner of hearts, | examiner-of hearts
- 2:6.97 2:6.98 2:6.99 2:6.100 | for I did not give sleep | for not I-gave sleep
- 2:6.101 2:6.102 2:6.103 2:6.104 | to my eyes, and my sleep was taken from my eyes | to-my-eyes and-was-taken my-sleep from-my-eyes
- 2:6.105 2:6.106 2:6.107 | until I find solutions | until I-find solutions
- 2:6.108 2:6.109 2:6.110 2:6.111 | and whoever comes to purify himself, they help him, | and-whoever-comes to-purify-himself they-help him
- 2:6.112 2:6.113 2:6.114 2:6.115 2:6.116 | to provide for him a guide to the way of holiness | to-provide for-him a-guide to-the-way-of holiness
- 2:6.117 2:6.118 | by intellect and knowledge, | by-intellect and-knowledge
- 2:6.119 2:6.120 | like Khalkhal and Darda, | like-Khalkhal and-Darda
- 2:6.121 2:6.122 2:6.123 | and among the sages of truth | and-among the-sages-of truth
- 2:6.124 2:6.125 | a sage and an artisan, | a-sage and-an-artisan
- 2:6.126 2:6.127 | like Ethan and Heman, | like-Ethan and-Heman
- 2:6.128 2:6.129 2:6.130 2:6.131 2:6.132 2:6.133 | perhaps the rabbis who lack this wisdom | perhaps the-rabbis who-lack of the-wisdom this
- 2:6.134 2:6.135 2:6.136 2:6.137 | lay upon us vain burdens and rejected ones, | for-us burdens vain rejected
- 2:6.138 | God forbid. | God-forbid
- 2:6.139 2:6.140 | And you, my lord, | and-you my-lord
- 2:6.141 2:6.142 2:6.143 | to you its secrets are revealed, | to-you are-revealed its-secrets
- 2:6.144 2:6.145 2:6.146 2:6.147 2:6.148 | and your treasure will fill the bellies of your disciples, | and-your-treasure will-fill their-bellies of your-disciples
- 2:6.149 2:6.150 2:6.151 | they will be sated and delight in your goodness, | will-be-sated and-will-delight in-your-goodness
- 2:6.152 2:6.153 2:6.154 2:6.155 | and upon us to pay the wages of your work. | and-upon-us to-pay the-wage-of your-work
- 2:6.156 2:6.157 2:6.158 2:6.159 2:6.160 | And you, do not withhold good from its owners. | and-you do-not withhold good from-its-owners
- 2:6.161 2:6.162 2:6.163 2:6.164 | And here I adjure you by the God of heaven: | and-here-I adjure-you by-the-God-of heaven
- 2:6.165 2:6.166 2:6.167 | make Your face shine upon us, | make-shine Your-face upon-us
- 2:6.168 2:6.169 2:6.170 | to answer the weary | to-answer (object-marker) the-weary
- 2:6.171 2:6.172 2:6.173 2:6.174 | a word spoken in its proper setting, | a-word spoken in its-proper-setting
- 2:6.175 2:6.176 2:6.177 | and we shall be servants to our master. | and-we-shall-be servants to-our-master
- 2:6.178 2:6.179 2:6.180 2:6.181 2:6.182 | And our lips will ask a great reward for your goodness, | and-reward great for-your-goodness will-ask our-lips
- 2:6.183 2:6.184 2:6.185 2:6.186 2:6.187 2:6.188 2:6.189 | then we shall not be ashamed, neither in this world nor in the world to come, | and-then not we-shall-be-ashamed not in-this-world and-not in-the-world-to-come
- 2:6.190 2:6.191 2:6.192 2:6.193 2:6.194 2:6.195 | and we shall return and graze among the lilies from the company of sages and discerning ones. | and-we-shall-return and-graze among-the-lilies from-the-company-of sages and-discerning-ones

Grammar notes:
- אמרו שאין ראוי is indirect speech ('said that it is not fitting'), so the segment has no quotation marks; the four infinitives that follow are its objects.
- ולמה נניח ... ונעבוד ...? is a rhetorical question in first-person plural cohortatives; the Hebrew closes it with '?'.
- מי יעלה בהר ה' and כי בער אנכי מאיש are embedded verses (Ps 24:3; Prov 30:2) with no quotation marks; their imperfects keep the verses' future sense.
- הביננו אדונינו נלכה נרדפה: an imperative with a vocative, then a cohortative pair ('let us go and pursue').
- ישבע ויתענגו are imperfects after תמלא and read as a result: 'they will be sated and delight'.

Translator's note: The questioner's letter, in his voice, addressed to the respondent as 'my lord'. Biblical phrases are embedded and unmarked.

Commentary rulings:
- apparatus: Ps 24:3 (embedded): מי יעלה בהר ה' ומי יקום במקום קדשו.
- apparatus: Prov 30:2 (embedded): כי בער אנכי מאיש ולא בינת אדם לי.
- apparatus: Ps 44:22 (embedded): יודע תעלומות לב. Prov 17:3: ובֹחֵן לִבּוֹת, 'examiner of hearts'.
- apparatus: Ps 17:14 (embedded): וצפונך תמלא בטנם.
- apparatus: Prov 3:27: אל תמנע טוב מבעליו.
- apparatus: Ps 80:4 (Hebrew numbering, embedded): האר פניך.
- apparatus: Isa 50:4 (embedded): לעות את יעף דבר.
- apparatus: Prov 25:11 (embedded): דבר דבור על אופניו; the source writes אפניו.
- apparatus: Song 2:16 (embedded): הָרֹעֶה בַּשּׁוֹשַׁנִּים.
- apparatus: 1 Kings 5:11 (Hebrew): the wise men Ethan, Heman, Khalkhal and Darda.
- apparatus: Wikisource base text; not yet collated with the 1931 scan (Q2).

## Milchamot ha-Shem 2:7

Hebrew (Wikisource): תשובה שנית ראיתי את כל הרשום בכתב האלוף, אילן שענפיו מרובים, רומו כגובה ארזים, אחר המחילה על השהייה כי טרדת הזמן נטויה, וכל עכבה לטובה. בהשקיפי על המכתב והנה פערת פיך בהזכירך אלוהות בחשבך שהוא כנגדינו. ונחנו מה כי תלינו עלינו ובזיונך אינו מגיע אלינו כי אם אל חכמי ישראל אשר קטנם עבה ממתניך וכל ישראל נחים אליהם ובזכותם העולם עומד. אי עליך כי הוצאת עליהם דברים אשר לא כדת, בתומה ותחתיה ערמה. אבל כבר נתבאר דינך בפי כל בעלי חכמה. ומה שכתבת על כסא אליהו כי בדף ג' מדמה אחדותו ית' לשאר אחדים. כבר אתה הולך חשכים, הלא תחילת דבריו דף עשרים שכל דמיונות אלו על אורות קדושים אשר אין להם תפיסה וצורה ודמיון, לא בשכל ולא בדעת כלל ועיקר, אלא כדי שיוכל השכל לקבל אותם הורשינו לדמות. ובדף ד' ע"ב כתב שכל הכנויים והתארים הם בדרך דמיון בספירות ואודות קדושים הנאצלים, אבל בו ית' אין בו שום דמיון וצורה ותפישה במחשבה כלל ועיקר. ומה שכתבת איך נניח שאר הפרצופים ונעבוד לזעיר אנפין עם כח שבו א"ס ב"ה. כבר ביאר הר' כס"א בדף כ"ה ע"ב שבקראינו לז"א אנחנו קורין לכל הפרצופים כאחד עם נשמה המסתתרת בהם. ומה שכתבת למה נניח הפרצופים הקרובים לא"ס ונקרא לז"א. לפי שידוע שהמידות הפנימיות יש מהם דין ויש מהם רחמים, וזה כלול מכולם, ולזה רצה המאציל לפעול פעולותיו על ידי כלי זה. ועד שאתה מקשה על דמיונות אודות עליונים שאין לך רשות להרהר בהן במכל שכן ללגלג, תקשה על ידך הגופנית שהיא עושה כל המלאכות, והראש שלמעלה אינו עושה כלום. ג"כ כח הדבור שהוא פארך על כל בעל חי למה לא ניתן למעלה מעיניך. אלא כך גזרה חכמתו ית'. ואם יש לך עצה הוגנת תלך אצל בוראך ותיעצנו. ומה שכתבת שדברי כס"א נגד הרמב"ם וכו' שכתבו שאינו כח בגוף, ולפי דברי כס"א הוא כח בגוף. הלא כל המקובלים פה אחד עונים ואומרים שכל המגשים בספירות כמוך הרי זה כופר ואין לו חלק באלהי ישראל. ומה שכינו אותם אורות דקים אינו אלא לשכך את האזן! יש לנו לשאול על מה שכתב הרמב"ם: יכוין לבו כנגד השכינה ויתפלל. ובמורה כ' שהשכינה אור נברא. ויותר מזה מה שכתב על כבוד ה' שהוא נברא. אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחווים לנברא. ויותר על זה מה שכתב שכל צורה שהיו רואים הנביאים ומדברים עמה היא אור נברא. א"כ אין לנו שום נבואה ע"י הבורא, אלא כולם מצד הנבראים. ואם נאמר שהיה הנברא דבק בבורא, א"כ לדעתך הוא כח בגוף! ותשובתך מהרה תצמח, ס' מאנו דעת אותי נאם ה'.

| id | word | punctuation after |
|---|---|---|
| 2:7.1 | תשובה |  |
| 2:7.2 | שנית |  |
| 2:7.3 | ראיתי |  |
| 2:7.4 | את |  |
| 2:7.5 | כל |  |
| 2:7.6 | הרשום |  |
| 2:7.7 | בכתב |  |
| 2:7.8 | האלוף | , |
| 2:7.9 | אילן |  |
| 2:7.10 | שענפיו |  |
| 2:7.11 | מרובים | , |
| 2:7.12 | רומו |  |
| 2:7.13 | כגובה |  |
| 2:7.14 | ארזים | , |
| 2:7.15 | אחר |  |
| 2:7.16 | המחילה |  |
| 2:7.17 | על |  |
| 2:7.18 | השהייה |  |
| 2:7.19 | כי |  |
| 2:7.20 | טרדת |  |
| 2:7.21 | הזמן |  |
| 2:7.22 | נטויה | , |
| 2:7.23 | וכל |  |
| 2:7.24 | עכבה |  |
| 2:7.25 | לטובה | . |
| 2:7.26 | בהשקיפי |  |
| 2:7.27 | על |  |
| 2:7.28 | המכתב |  |
| 2:7.29 | והנה |  |
| 2:7.30 | פערת |  |
| 2:7.31 | פיך |  |
| 2:7.32 | בהזכירך |  |
| 2:7.33 | אלוהות |  |
| 2:7.34 | בחשבך |  |
| 2:7.35 | שהוא |  |
| 2:7.36 | כנגדינו | . |
| 2:7.37 | ונחנו |  |
| 2:7.38 | מה |  |
| 2:7.39 | כי |  |
| 2:7.40 | תלינו |  |
| 2:7.41 | עלינו |  |
| 2:7.42 | ובזיונך |  |
| 2:7.43 | אינו |  |
| 2:7.44 | מגיע |  |
| 2:7.45 | אלינו |  |
| 2:7.46 | כי |  |
| 2:7.47 | אם |  |
| 2:7.48 | אל |  |
| 2:7.49 | חכמי |  |
| 2:7.50 | ישראל |  |
| 2:7.51 | אשר |  |
| 2:7.52 | קטנם |  |
| 2:7.53 | עבה |  |
| 2:7.54 | ממתניך |  |
| 2:7.55 | וכל |  |
| 2:7.56 | ישראל |  |
| 2:7.57 | נחים |  |
| 2:7.58 | אליהם |  |
| 2:7.59 | ובזכותם |  |
| 2:7.60 | העולם |  |
| 2:7.61 | עומד | . |
| 2:7.62 | אי |  |
| 2:7.63 | עליך |  |
| 2:7.64 | כי |  |
| 2:7.65 | הוצאת |  |
| 2:7.66 | עליהם |  |
| 2:7.67 | דברים |  |
| 2:7.68 | אשר |  |
| 2:7.69 | לא |  |
| 2:7.70 | כדת | , |
| 2:7.71 | בתומה |  |
| 2:7.72 | ותחתיה |  |
| 2:7.73 | ערמה | . |
| 2:7.74 | אבל |  |
| 2:7.75 | כבר |  |
| 2:7.76 | נתבאר |  |
| 2:7.77 | דינך |  |
| 2:7.78 | בפי |  |
| 2:7.79 | כל |  |
| 2:7.80 | בעלי |  |
| 2:7.81 | חכמה | . |
| 2:7.82 | ומה |  |
| 2:7.83 | שכתבת |  |
| 2:7.84 | על |  |
| 2:7.85 | כסא |  |
| 2:7.86 | אליהו |  |
| 2:7.87 | כי |  |
| 2:7.88 | בדף |  |
| 2:7.89 | ג | ' |
| 2:7.90 | מדמה |  |
| 2:7.91 | אחדותו |  |
| 2:7.92 | ית | ' |
| 2:7.93 | לשאר |  |
| 2:7.94 | אחדים | . |
| 2:7.95 | כבר |  |
| 2:7.96 | אתה |  |
| 2:7.97 | הולך |  |
| 2:7.98 | חשכים | , |
| 2:7.99 | הלא |  |
| 2:7.100 | תחילת |  |
| 2:7.101 | דבריו |  |
| 2:7.102 | דף |  |
| 2:7.103 | עשרים |  |
| 2:7.104 | שכל |  |
| 2:7.105 | דמיונות |  |
| 2:7.106 | אלו |  |
| 2:7.107 | על |  |
| 2:7.108 | אורות |  |
| 2:7.109 | קדושים |  |
| 2:7.110 | אשר |  |
| 2:7.111 | אין |  |
| 2:7.112 | להם |  |
| 2:7.113 | תפיסה |  |
| 2:7.114 | וצורה |  |
| 2:7.115 | ודמיון | , |
| 2:7.116 | לא |  |
| 2:7.117 | בשכל |  |
| 2:7.118 | ולא |  |
| 2:7.119 | בדעת |  |
| 2:7.120 | כלל |  |
| 2:7.121 | ועיקר | , |
| 2:7.122 | אלא |  |
| 2:7.123 | כדי |  |
| 2:7.124 | שיוכל |  |
| 2:7.125 | השכל |  |
| 2:7.126 | לקבל |  |
| 2:7.127 | אותם |  |
| 2:7.128 | הורשינו |  |
| 2:7.129 | לדמות | . |
| 2:7.130 | ובדף |  |
| 2:7.131 | ד | ' |
| 2:7.132 | ע"ב |  |
| 2:7.133 | כתב |  |
| 2:7.134 | שכל |  |
| 2:7.135 | הכנויים |  |
| 2:7.136 | והתארים |  |
| 2:7.137 | הם |  |
| 2:7.138 | בדרך |  |
| 2:7.139 | דמיון |  |
| 2:7.140 | בספירות |  |
| 2:7.141 | ואודות |  |
| 2:7.142 | קדושים |  |
| 2:7.143 | הנאצלים | , |
| 2:7.144 | אבל |  |
| 2:7.145 | בו |  |
| 2:7.146 | ית | ' |
| 2:7.147 | אין |  |
| 2:7.148 | בו |  |
| 2:7.149 | שום |  |
| 2:7.150 | דמיון |  |
| 2:7.151 | וצורה |  |
| 2:7.152 | ותפישה |  |
| 2:7.153 | במחשבה |  |
| 2:7.154 | כלל |  |
| 2:7.155 | ועיקר | . |
| 2:7.156 | ומה |  |
| 2:7.157 | שכתבת |  |
| 2:7.158 | איך |  |
| 2:7.159 | נניח |  |
| 2:7.160 | שאר |  |
| 2:7.161 | הפרצופים |  |
| 2:7.162 | ונעבוד |  |
| 2:7.163 | לזעיר |  |
| 2:7.164 | אנפין |  |
| 2:7.165 | עם |  |
| 2:7.166 | כח |  |
| 2:7.167 | שבו |  |
| 2:7.168 | א"ס |  |
| 2:7.169 | ב"ה | . |
| 2:7.170 | כבר |  |
| 2:7.171 | ביאר |  |
| 2:7.172 | הר | ' |
| 2:7.173 | כס"א |  |
| 2:7.174 | בדף |  |
| 2:7.175 | כ"ה |  |
| 2:7.176 | ע"ב |  |
| 2:7.177 | שבקראינו |  |
| 2:7.178 | לז"א |  |
| 2:7.179 | אנחנו |  |
| 2:7.180 | קורין |  |
| 2:7.181 | לכל |  |
| 2:7.182 | הפרצופים |  |
| 2:7.183 | כאחד |  |
| 2:7.184 | עם |  |
| 2:7.185 | נשמה |  |
| 2:7.186 | המסתתרת |  |
| 2:7.187 | בהם | . |
| 2:7.188 | ומה |  |
| 2:7.189 | שכתבת |  |
| 2:7.190 | למה |  |
| 2:7.191 | נניח |  |
| 2:7.192 | הפרצופים |  |
| 2:7.193 | הקרובים |  |
| 2:7.194 | לא"ס |  |
| 2:7.195 | ונקרא |  |
| 2:7.196 | לז"א | . |
| 2:7.197 | לפי |  |
| 2:7.198 | שידוע |  |
| 2:7.199 | שהמידות |  |
| 2:7.200 | הפנימיות |  |
| 2:7.201 | יש |  |
| 2:7.202 | מהם |  |
| 2:7.203 | דין |  |
| 2:7.204 | ויש |  |
| 2:7.205 | מהם |  |
| 2:7.206 | רחמים | , |
| 2:7.207 | וזה |  |
| 2:7.208 | כלול |  |
| 2:7.209 | מכולם | , |
| 2:7.210 | ולזה |  |
| 2:7.211 | רצה |  |
| 2:7.212 | המאציל |  |
| 2:7.213 | לפעול |  |
| 2:7.214 | פעולותיו |  |
| 2:7.215 | על |  |
| 2:7.216 | ידי |  |
| 2:7.217 | כלי |  |
| 2:7.218 | זה | . |
| 2:7.219 | ועד |  |
| 2:7.220 | שאתה |  |
| 2:7.221 | מקשה |  |
| 2:7.222 | על |  |
| 2:7.223 | דמיונות |  |
| 2:7.224 | אודות |  |
| 2:7.225 | עליונים |  |
| 2:7.226 | שאין |  |
| 2:7.227 | לך |  |
| 2:7.228 | רשות |  |
| 2:7.229 | להרהר |  |
| 2:7.230 | בהן |  |
| 2:7.231 | במכל |  |
| 2:7.232 | שכן |  |
| 2:7.233 | ללגלג | , |
| 2:7.234 | תקשה |  |
| 2:7.235 | על |  |
| 2:7.236 | ידך |  |
| 2:7.237 | הגופנית |  |
| 2:7.238 | שהיא |  |
| 2:7.239 | עושה |  |
| 2:7.240 | כל |  |
| 2:7.241 | המלאכות | , |
| 2:7.242 | והראש |  |
| 2:7.243 | שלמעלה |  |
| 2:7.244 | אינו |  |
| 2:7.245 | עושה |  |
| 2:7.246 | כלום | . |
| 2:7.247 | ג"כ |  |
| 2:7.248 | כח |  |
| 2:7.249 | הדבור |  |
| 2:7.250 | שהוא |  |
| 2:7.251 | פארך |  |
| 2:7.252 | על |  |
| 2:7.253 | כל |  |
| 2:7.254 | בעל |  |
| 2:7.255 | חי |  |
| 2:7.256 | למה |  |
| 2:7.257 | לא |  |
| 2:7.258 | ניתן |  |
| 2:7.259 | למעלה |  |
| 2:7.260 | מעיניך | . |
| 2:7.261 | אלא |  |
| 2:7.262 | כך |  |
| 2:7.263 | גזרה |  |
| 2:7.264 | חכמתו |  |
| 2:7.265 | ית | '. |
| 2:7.266 | ואם |  |
| 2:7.267 | יש |  |
| 2:7.268 | לך |  |
| 2:7.269 | עצה |  |
| 2:7.270 | הוגנת |  |
| 2:7.271 | תלך |  |
| 2:7.272 | אצל |  |
| 2:7.273 | בוראך |  |
| 2:7.274 | ותיעצנו | . |
| 2:7.275 | ומה |  |
| 2:7.276 | שכתבת |  |
| 2:7.277 | שדברי |  |
| 2:7.278 | כס"א |  |
| 2:7.279 | נגד |  |
| 2:7.280 | הרמב"ם |  |
| 2:7.281 | וכו | ' |
| 2:7.282 | שכתבו |  |
| 2:7.283 | שאינו |  |
| 2:7.284 | כח |  |
| 2:7.285 | בגוף | , |
| 2:7.286 | ולפי |  |
| 2:7.287 | דברי |  |
| 2:7.288 | כס"א |  |
| 2:7.289 | הוא |  |
| 2:7.290 | כח |  |
| 2:7.291 | בגוף | . |
| 2:7.292 | הלא |  |
| 2:7.293 | כל |  |
| 2:7.294 | המקובלים |  |
| 2:7.295 | פה |  |
| 2:7.296 | אחד |  |
| 2:7.297 | עונים |  |
| 2:7.298 | ואומרים |  |
| 2:7.299 | שכל |  |
| 2:7.300 | המגשים |  |
| 2:7.301 | בספירות |  |
| 2:7.302 | כמוך |  |
| 2:7.303 | הרי |  |
| 2:7.304 | זה |  |
| 2:7.305 | כופר |  |
| 2:7.306 | ואין |  |
| 2:7.307 | לו |  |
| 2:7.308 | חלק |  |
| 2:7.309 | באלהי |  |
| 2:7.310 | ישראל | . |
| 2:7.311 | ומה |  |
| 2:7.312 | שכינו |  |
| 2:7.313 | אותם |  |
| 2:7.314 | אורות |  |
| 2:7.315 | דקים |  |
| 2:7.316 | אינו |  |
| 2:7.317 | אלא |  |
| 2:7.318 | לשכך |  |
| 2:7.319 | את |  |
| 2:7.320 | האזן | ! |
| 2:7.321 | יש |  |
| 2:7.322 | לנו |  |
| 2:7.323 | לשאול |  |
| 2:7.324 | על |  |
| 2:7.325 | מה |  |
| 2:7.326 | שכתב |  |
| 2:7.327 | הרמב"ם | : |
| 2:7.328 | יכוין |  |
| 2:7.329 | לבו |  |
| 2:7.330 | כנגד |  |
| 2:7.331 | השכינה |  |
| 2:7.332 | ויתפלל | . |
| 2:7.333 | ובמורה |  |
| 2:7.334 | כ | ' |
| 2:7.335 | שהשכינה |  |
| 2:7.336 | אור |  |
| 2:7.337 | נברא | . |
| 2:7.338 | ויותר |  |
| 2:7.339 | מזה |  |
| 2:7.340 | מה |  |
| 2:7.341 | שכתב |  |
| 2:7.342 | על |  |
| 2:7.343 | כבוד |  |
| 2:7.344 | ה | ' |
| 2:7.345 | שהוא |  |
| 2:7.346 | נברא | . |
| 2:7.347 | אם |  |
| 2:7.348 | כן |  |
| 2:7.349 | כיצד |  |
| 2:7.350 | היו |  |
| 2:7.351 | ישראל |  |
| 2:7.352 | יוצאים |  |
| 2:7.353 | פתח |  |
| 2:7.354 | אהליהם |  |
| 2:7.355 | ומשתחווים |  |
| 2:7.356 | לנברא | . |
| 2:7.357 | ויותר |  |
| 2:7.358 | על |  |
| 2:7.359 | זה |  |
| 2:7.360 | מה |  |
| 2:7.361 | שכתב |  |
| 2:7.362 | שכל |  |
| 2:7.363 | צורה |  |
| 2:7.364 | שהיו |  |
| 2:7.365 | רואים |  |
| 2:7.366 | הנביאים |  |
| 2:7.367 | ומדברים |  |
| 2:7.368 | עמה |  |
| 2:7.369 | היא |  |
| 2:7.370 | אור |  |
| 2:7.371 | נברא | . |
| 2:7.372 | א"כ |  |
| 2:7.373 | אין |  |
| 2:7.374 | לנו |  |
| 2:7.375 | שום |  |
| 2:7.376 | נבואה |  |
| 2:7.377 | ע"י |  |
| 2:7.378 | הבורא | , |
| 2:7.379 | אלא |  |
| 2:7.380 | כולם |  |
| 2:7.381 | מצד |  |
| 2:7.382 | הנבראים | . |
| 2:7.383 | ואם |  |
| 2:7.384 | נאמר |  |
| 2:7.385 | שהיה |  |
| 2:7.386 | הנברא |  |
| 2:7.387 | דבק |  |
| 2:7.388 | בבורא | , |
| 2:7.389 | א"כ |  |
| 2:7.390 | לדעתך |  |
| 2:7.391 | הוא |  |
| 2:7.392 | כח |  |
| 2:7.393 | בגוף | ! |
| 2:7.394 | ותשובתך |  |
| 2:7.395 | מהרה |  |
| 2:7.396 | תצמח | , |
| 2:7.397 | ס | ' |
| 2:7.398 | מאנו |  |
| 2:7.399 | דעת |  |
| 2:7.400 | אותי |  |
| 2:7.401 | נאם |  |
| 2:7.402 | ה | '. |


### Translation under review (2:7)

English: Second Reply. I saw all that is written in the letter of the eminent one, a tree whose branches are many, its height like the height of cedars, After the pardon for the delay, for the pressure of the times is outstretched, and every delay is for good. When I looked at the letter, behold, you opened your mouth, by mentioning divinities, thinking that it is against us. And what are we that you lay upon us, and your scorn does not reach us, but only to the sages of Israel, whose least is thicker than your loins, and all Israel rests on them, and by their merit the world stands. Woe to you, for you have brought out against them words that are not according to the law, in innocence, and in its place guile. But your judgment has already been made clear, by the mouth of all the masters of wisdom. And as for what you wrote about Kisse Eliyahu, that on folio 3 he likens His oneness, may He be blessed, to the other unities. You are already walking in darkness. Surely the beginning of his words on folio 20 is that all these likenesses are about holy lights, which have no grasp form and likeness, neither in intellect nor in knowledge at all and in principle, but so that the intellect can receive them, we were permitted to liken. And on folio 4b he wrote that all the epithets and attributes are in the manner of likeness, in the sefirot and concerning holy, the emanated beings. But in Him, may He be blessed, there is not in Him any likeness and form, and grasp in thought at all and in principle. And as for what you wrote, how can we leave the other partzufim and worship Ze'ir Anpin with the power in him, Ein Sof, blessed be He. Rabbi Kisse Eliyahu has already explained on folio 25b, that when we call upon Ze'ir Anpin, we call all the partzufim together, with a soul hidden in them. And as for what you wrote, why do we leave the partzufim closest to Ein Sof and call upon Ze'ir Anpin. Because it is known that the inner attributes some of them judgment, and some of them mercy, and this one includes them all, and therefore the Emanator wanted to act His actions through this vessel. even to the point that you object to the likenesses concerning the upper ones, which you have no right to ponder all the more so to mock, then object to your corporeal hand, which does all the labors, and the head above does not do anything. Also, the power of speech, which is your glory, above every living creature, why was it not given above your eyes. But thus His wisdom decreed, may He be blessed. And if you have proper counsel, go to your Creator and let Him counsel us. And as for what you wrote, that the words of Kisse Eliyahu are against the Rambam, etc., who wrote that it is not a force in a body, and according to the words of Kisse Eliyahu it is a force in a body. Surely all the Kabbalists with one mouth answer and say that whoever ascribes corporeality to the sefirot, like you, is a denier and has no share in the God of Israel. And as for what they called them subtle lights, is only to soothe the ear! We have to ask about what the Rambam wrote: "Let him direct his heart toward the Shekhinah and pray." And in the Guide, 20, that the Shekhinah is a created light. And more than this, what he wrote about the glory of the LORD, that it is created. If so, how did Israel go out to the opening of their tents and bow down to the created thing. And more than this, what he wrote, that every form which the prophets saw and spoke with is a created light. If so, we have no prophecy through the Creator, but all of it from the side of the created ones. And if it is said that the created thing was attached to the Creator, then, in your opinion, it is a force in a body! Your answer will quickly spring up, "They refuse to know Me, says the LORD."

Literal: Reply second I-saw (object-marker) all what-is-written in-the-letter-of the-eminent-one a-tree whose-branches are-many its-height like-the-height-of cedars after the-pardon for the-delay for the-pressure-of the-times is-outstretched and-every delay is-for-good when-I-looked at the-letter and-behold you-opened your-mouth by-mentioning divinities thinking that-it-is against-us and-we what that you-lay upon-us and-your-scorn does-not reach to-us but only to the-sages-of Israel who their-least is-thicker from-your-loins and-all Israel rest on-them and-by-their-merit the-world stands woe to-you for you-have-brought-out against-them words which are-not according-to-the-law in-innocence and-in-its-place guile but already was-made-clear your-judgment by-the-mouth-of all the-masters-of wisdom and-as-for what-you-wrote about Kisse Eliyahu that on-folio 3 he-likens His-oneness may-He-be-blessed to-the-other unities already you walk in-darkness surely the-beginning-of his-words folio 20 that-all the-likenesses these are-about lights holy which have-no for-them grasp and-form and-likeness neither in-intellect nor in-knowledge at-all and-in-principle but so-that that-it-can the-intellect receive them we-were-permitted to-liken and-on-folio 4 b he-wrote that-all the-epithets and-the-attributes are in-the-manner-of likeness in-the-sefirot and-concerning holy the-emanated-beings but in-Him may-He-be-blessed there-is-not in-Him any likeness and-form and-grasp in-thought at-all and-in-principle and-as-for what-you-wrote how can-we-leave the-other the-partzufim and-worship Ze'ir Anpin with the-power in-him Ein-Sof blessed-be-He already explained Rabbi Kisse-Eliyahu on-folio 25 b that-when-we-call upon-Ze'ir-Anpin we call all the-partzufim together with a-soul hidden in-them and-as-for what-you-wrote why do-we-leave the-partzufim closest to-Ein-Sof and-call-upon Ze'ir-Anpin because it-is-known that-the-attributes inner there-are of-them judgment and-there-are of-them mercy and-this-one includes all-of-them and-therefore wanted the-Emanator to-act His-actions through the-hand-of a-vessel this even-to-the-point that-you object to the-likenesses concerning the-upper-ones that-you-have-no for-you right to-ponder them all-the-more so to-mock then-object to your-hand the-corporeal which-is does all the-labors and-the-head above does-not do anything also the-power-of speech which-is your-glory above every living creature why not was-it-given above your-eyes but thus decreed His-wisdom may-He-be-blessed and-if you-have for-you counsel proper go to your-Creator and-let-Him-counsel-us and-as-for what-you-wrote that-the-words-of Kisse-Eliyahu are-against the-Rambam etc. who-wrote that-it-is-not a-force in-a-body and-according-to the-words-of Kisse-Eliyahu it-is a-force in-a-body surely all the-Kabbalists with-one mouth answer and-say that-all those-who-ascribe-corporeality to-the-sefirot like-you is such-a-one a-denier and-has-no for-him share in-the-God-of Israel and-as-for what-they-called them lights subtle is-not but to-soothe (object-marker) the-ear we-have for-us to-ask about what he-wrote the-Rambam let-him-direct his-heart toward the-Shekhinah and-pray and-in-the-Guide 20 that-the-Shekhinah-is a-light created and-more than-this what he-wrote about the-glory-of the-LORD that-it-is created if so how were Israel going-out to-the-opening-of their-tents and-bowing to-the-created-thing and-more than this what he-wrote that-every form which seeing the-prophets and-speaking with-it is a-light created if-so there-is-no for-us any prophecy through the-Creator but all-of-it from-the-side-of the-created-ones and-if it-is-said that the-created attached to-the-Creator if-so in-your-opinion it-is a-force in-a-body your-answer quickly will-spring-up (the-verse) they-refuse to-know Me says the-LORD

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:7.1 | Reply |  |  |  |
| 2:7.2 | second |  |  |  |
| 2:7.3 | I saw |  |  |  |
| 2:7.4 | (object marker) |  |  |  |
| 2:7.5 | all |  |  |  |
| 2:7.6 | what is written |  |  |  |
| 2:7.7 | in the letter of |  |  |  |
| 2:7.8 | the eminent one |  |  |  |
| 2:7.9 | a tree |  |  |  |
| 2:7.10 | whose branches |  |  |  |
| 2:7.11 | are many |  |  |  |
| 2:7.12 | its height |  |  |  |
| 2:7.13 | like the height of |  |  |  |
| 2:7.14 | cedars |  |  |  |
| 2:7.15 | after |  |  |  |
| 2:7.16 | the pardon |  |  |  |
| 2:7.17 | for |  |  |  |
| 2:7.18 | the delay |  |  |  |
| 2:7.19 | for |  |  |  |
| 2:7.20 | the pressure of |  |  |  |
| 2:7.21 | the times |  |  |  |
| 2:7.22 | is outstretched |  |  |  |
| 2:7.23 | and every |  |  |  |
| 2:7.24 | delay |  |  |  |
| 2:7.25 | is for good |  |  |  |
| 2:7.26 | when I looked |  |  |  |
| 2:7.27 | at |  |  |  |
| 2:7.28 | the letter |  |  |  |
| 2:7.29 | and behold |  |  |  |
| 2:7.30 | you opened |  |  |  |
| 2:7.31 | your mouth |  |  |  |
| 2:7.32 | by mentioning |  |  |  |
| 2:7.33 | divinities |  |  | Counted here, so 'divinities'; the abstract is 'divinity'. |
| 2:7.34 | thinking |  |  |  |
| 2:7.35 | that it is |  |  |  |
| 2:7.36 | against us |  |  |  |
| 2:7.37 | and we |  |  |  |
| 2:7.38 | what |  |  |  |
| 2:7.39 | that |  |  |  |
| 2:7.40 | you lay |  |  |  |
| 2:7.41 | upon us |  |  |  |
| 2:7.42 | and your scorn |  |  |  |
| 2:7.43 | does not |  |  |  |
| 2:7.44 | reach |  |  |  |
| 2:7.45 | to us |  |  |  |
| 2:7.46 | but |  |  |  |
| 2:7.47 | only |  |  |  |
| 2:7.48 | to |  |  |  |
| 2:7.49 | the sages of |  |  |  |
| 2:7.50 | Israel |  |  |  |
| 2:7.51 | who |  |  |  |
| 2:7.52 | their least |  |  |  |
| 2:7.53 | is thicker |  |  |  |
| 2:7.54 | from your loins |  |  |  |
| 2:7.55 | and all |  |  |  |
| 2:7.56 | Israel |  |  |  |
| 2:7.57 | rest |  |  |  |
| 2:7.58 | on them |  |  |  |
| 2:7.59 | and by their merit |  |  |  |
| 2:7.60 | the world |  |  |  |
| 2:7.61 | stands |  |  |  |
| 2:7.62 | woe |  |  |  |
| 2:7.63 | to you |  |  |  |
| 2:7.64 | for |  |  |  |
| 2:7.65 | you have brought out |  |  |  |
| 2:7.66 | against them |  |  |  |
| 2:7.67 | words |  |  |  |
| 2:7.68 | which |  |  |  |
| 2:7.69 | are not |  |  |  |
| 2:7.70 | according to the law |  |  |  |
| 2:7.71 | in innocence |  |  |  |
| 2:7.72 | and in its place |  |  |  |
| 2:7.73 | guile |  |  |  |
| 2:7.74 | but |  |  |  |
| 2:7.75 | already |  |  |  |
| 2:7.76 | was made clear |  |  |  |
| 2:7.77 | your judgment |  |  |  |
| 2:7.78 | by the mouth of |  |  |  |
| 2:7.79 | all |  |  |  |
| 2:7.80 | the masters of |  |  |  |
| 2:7.81 | wisdom |  |  |  |
| 2:7.82 | and as for |  |  |  |
| 2:7.83 | what you wrote |  |  |  |
| 2:7.84 | about |  |  |  |
| 2:7.85 | Kisse |  |  | Kisse Eliyahu: the book, named in full here; the Hebrew title means 'the throne of Elijah'. First occurrence in the unit. |
| 2:7.86 | Eliyahu |  |  |  |
| 2:7.87 | that |  |  |  |
| 2:7.88 | on folio |  |  |  |
| 2:7.89 | 3 |  |  |  |
| 2:7.90 | he likens |  |  |  |
| 2:7.91 | His oneness |  |  |  |
| 2:7.92 | may He be blessed |  |  |  |
| 2:7.93 | to the other |  |  |  |
| 2:7.94 | unities |  |  |  |
| 2:7.95 | already |  |  |  |
| 2:7.96 | you |  |  |  |
| 2:7.97 | walk |  |  |  |
| 2:7.98 | in darkness |  |  |  |
| 2:7.99 | surely |  |  |  |
| 2:7.100 | the beginning of |  |  |  |
| 2:7.101 | his words |  |  |  |
| 2:7.102 | folio |  |  |  |
| 2:7.103 | 20 |  |  | Folio 20, as the author gives it. |
| 2:7.104 | that all |  |  |  |
| 2:7.105 | the likenesses |  |  |  |
| 2:7.106 | these |  |  |  |
| 2:7.107 | are about |  |  |  |
| 2:7.108 | lights |  |  |  |
| 2:7.109 | holy |  |  |  |
| 2:7.110 | which |  |  |  |
| 2:7.111 | have no |  |  |  |
| 2:7.112 | for them |  |  |  |
| 2:7.113 | grasp |  |  |  |
| 2:7.114 | and form |  |  |  |
| 2:7.115 | and likeness |  |  |  |
| 2:7.116 | neither |  |  |  |
| 2:7.117 | in intellect |  |  |  |
| 2:7.118 | nor |  |  |  |
| 2:7.119 | in knowledge |  |  |  |
| 2:7.120 | at all |  |  |  |
| 2:7.121 | and in principle |  |  |  |
| 2:7.122 | but |  |  |  |
| 2:7.123 | so that |  |  |  |
| 2:7.124 | that it can |  |  |  |
| 2:7.125 | the intellect |  |  |  |
| 2:7.126 | receive |  |  |  |
| 2:7.127 | them |  |  |  |
| 2:7.128 | we were permitted |  |  |  |
| 2:7.129 | to liken |  |  |  |
| 2:7.130 | and on folio |  |  |  |
| 2:7.131 | 4 |  |  | Folio 4 side b (ד' ע"ב): 4b. |
| 2:7.132 | b |  |  |  |
| 2:7.133 | he wrote |  |  |  |
| 2:7.134 | that all |  |  |  |
| 2:7.135 | the epithets |  |  |  |
| 2:7.136 | and the attributes |  |  |  |
| 2:7.137 | are |  |  |  |
| 2:7.138 | in the manner of |  |  |  |
| 2:7.139 | likeness |  |  |  |
| 2:7.140 | in the sefirot |  |  |  |
| 2:7.141 | and concerning |  |  |  |
| 2:7.142 | holy |  |  |  |
| 2:7.143 | the emanated beings |  |  |  |
| 2:7.144 | but |  |  |  |
| 2:7.145 | in Him |  |  |  |
| 2:7.146 | may He be blessed |  |  |  |
| 2:7.147 | there is not |  |  |  |
| 2:7.148 | in Him |  |  |  |
| 2:7.149 | any |  |  |  |
| 2:7.150 | likeness |  |  |  |
| 2:7.151 | and form |  |  |  |
| 2:7.152 | and grasp |  |  |  |
| 2:7.153 | in thought |  |  |  |
| 2:7.154 | at all |  |  |  |
| 2:7.155 | and in principle |  |  |  |
| 2:7.156 | and as for |  |  |  |
| 2:7.157 | what you wrote |  |  |  |
| 2:7.158 | how |  |  |  |
| 2:7.159 | can we leave |  |  |  |
| 2:7.160 | the other |  |  |  |
| 2:7.161 | the partzufim |  |  |  |
| 2:7.162 | and worship |  |  |  |
| 2:7.163 | Ze'ir |  |  | Ze'ir Anpin, 'Small Face' (see 2:5.11). |
| 2:7.164 | Anpin |  |  |  |
| 2:7.165 | with |  |  |  |
| 2:7.166 | the power |  |  |  |
| 2:7.167 | in him |  |  |  |
| 2:7.168 | Ein Sof |  |  |  |
| 2:7.169 | blessed be He |  |  |  |
| 2:7.170 | already |  |  |  |
| 2:7.171 | explained |  |  |  |
| 2:7.172 | Rabbi |  |  | הר' is the title 'Rabbi' before the book's name, as the Hebrew gives it. |
| 2:7.173 | Kisse Eliyahu |  |  | כס"א: Kisse Eliyahu, named in full at 2:7.85. |
| 2:7.174 | on folio |  |  |  |
| 2:7.175 | 25 |  |  | Folio 25 side b (כ"ה ע"ב): 25b. |
| 2:7.176 | b |  |  |  |
| 2:7.177 | that when we call |  |  |  |
| 2:7.178 | upon Ze'ir Anpin |  |  |  |
| 2:7.179 | we |  |  |  |
| 2:7.180 | call |  |  |  |
| 2:7.181 | all |  |  |  |
| 2:7.182 | the partzufim |  |  |  |
| 2:7.183 | together |  |  |  |
| 2:7.184 | with |  |  |  |
| 2:7.185 | a soul |  |  |  |
| 2:7.186 | hidden |  |  |  |
| 2:7.187 | in them |  |  |  |
| 2:7.188 | and as for |  |  |  |
| 2:7.189 | what you wrote |  |  |  |
| 2:7.190 | why |  |  |  |
| 2:7.191 | do we leave |  |  |  |
| 2:7.192 | the partzufim |  |  |  |
| 2:7.193 | closest |  |  |  |
| 2:7.194 | to Ein Sof |  |  |  |
| 2:7.195 | and call upon |  |  |  |
| 2:7.196 | Ze'ir Anpin |  |  |  |
| 2:7.197 | because |  |  |  |
| 2:7.198 | it is known |  |  |  |
| 2:7.199 | that the attributes |  |  |  |
| 2:7.200 | inner |  |  |  |
| 2:7.201 | there are |  |  |  |
| 2:7.202 | of them |  |  |  |
| 2:7.203 | judgment |  |  |  |
| 2:7.204 | and there are |  |  |  |
| 2:7.205 | of them |  |  |  |
| 2:7.206 | mercy |  |  |  |
| 2:7.207 | and this one |  |  |  |
| 2:7.208 | includes |  |  |  |
| 2:7.209 | all of them |  |  |  |
| 2:7.210 | and therefore |  |  |  |
| 2:7.211 | wanted |  |  |  |
| 2:7.212 | the Emanator |  |  |  |
| 2:7.213 | to act |  |  |  |
| 2:7.214 | His actions |  |  |  |
| 2:7.215 | through |  |  |  |
| 2:7.216 | the hand of |  |  |  |
| 2:7.217 | a vessel |  |  |  |
| 2:7.218 | this |  |  |  |
| 2:7.219 | even to the point |  |  |  |
| 2:7.220 | that you |  |  |  |
| 2:7.221 | object |  |  |  |
| 2:7.222 | to |  |  |  |
| 2:7.223 | the likenesses |  |  |  |
| 2:7.224 | concerning |  |  |  |
| 2:7.225 | the upper ones |  |  |  |
| 2:7.226 | that you have no |  |  |  |
| 2:7.227 | for you |  |  |  |
| 2:7.228 | right |  |  |  |
| 2:7.229 | to ponder |  |  |  |
| 2:7.230 | them |  |  |  |
| 2:7.231 | all the more |  |  |  |
| 2:7.232 | so |  |  |  |
| 2:7.233 | to mock |  |  |  |
| 2:7.234 | then object |  |  |  |
| 2:7.235 | to |  |  |  |
| 2:7.236 | your hand |  |  |  |
| 2:7.237 | the corporeal |  |  |  |
| 2:7.238 | which is |  |  |  |
| 2:7.239 | does |  |  |  |
| 2:7.240 | all |  |  |  |
| 2:7.241 | the labors |  |  |  |
| 2:7.242 | and the head |  |  |  |
| 2:7.243 | above |  |  |  |
| 2:7.244 | does not |  |  |  |
| 2:7.245 | do |  |  |  |
| 2:7.246 | anything |  |  |  |
| 2:7.247 | also |  |  |  |
| 2:7.248 | the power of |  |  |  |
| 2:7.249 | speech |  |  |  |
| 2:7.250 | which is |  |  |  |
| 2:7.251 | your glory |  |  |  |
| 2:7.252 | above |  |  |  |
| 2:7.253 | every |  |  |  |
| 2:7.254 | living |  |  |  |
| 2:7.255 | creature |  |  |  |
| 2:7.256 | why |  |  |  |
| 2:7.257 | not |  |  |  |
| 2:7.258 | was it given |  |  |  |
| 2:7.259 | above |  |  |  |
| 2:7.260 | your eyes |  |  |  |
| 2:7.261 | but |  |  |  |
| 2:7.262 | thus |  |  |  |
| 2:7.263 | decreed |  |  |  |
| 2:7.264 | His wisdom |  |  |  |
| 2:7.265 | may He be blessed |  |  |  |
| 2:7.266 | and if |  |  |  |
| 2:7.267 | you have |  |  |  |
| 2:7.268 | for you |  |  |  |
| 2:7.269 | counsel |  |  |  |
| 2:7.270 | proper |  |  |  |
| 2:7.271 | go |  |  |  |
| 2:7.272 | to |  |  |  |
| 2:7.273 | your Creator |  |  |  |
| 2:7.274 | and let Him counsel us |  |  |  |
| 2:7.275 | and as for |  |  |  |
| 2:7.276 | what you wrote |  |  |  |
| 2:7.277 | that the words of |  |  |  |
| 2:7.278 | Kisse Eliyahu |  |  |  |
| 2:7.279 | are against |  |  |  |
| 2:7.280 | the Rambam |  |  |  |
| 2:7.281 | etc. |  |  | וכו': 'etc.', kept as the Hebrew marks a truncated passage. |
| 2:7.282 | who wrote |  |  |  |
| 2:7.283 | that it is not |  |  |  |
| 2:7.284 | a force |  |  |  |
| 2:7.285 | in a body |  |  |  |
| 2:7.286 | and according to |  |  |  |
| 2:7.287 | the words of |  |  |  |
| 2:7.288 | Kisse Eliyahu |  |  |  |
| 2:7.289 | it is |  |  |  |
| 2:7.290 | a force |  |  |  |
| 2:7.291 | in a body |  |  |  |
| 2:7.292 | surely |  |  |  |
| 2:7.293 | all |  |  |  |
| 2:7.294 | the Kabbalists |  |  |  |
| 2:7.295 | with one |  |  |  |
| 2:7.296 | mouth |  |  |  |
| 2:7.297 | answer |  |  |  |
| 2:7.298 | and say |  |  |  |
| 2:7.299 | that all |  |  |  |
| 2:7.300 | those who ascribe corporeality |  |  |  |
| 2:7.301 | to the sefirot |  |  |  |
| 2:7.302 | like you |  |  |  |
| 2:7.303 | is |  |  |  |
| 2:7.304 | such a one |  |  |  |
| 2:7.305 | a denier |  |  |  |
| 2:7.306 | and has no |  |  |  |
| 2:7.307 | for him |  |  |  |
| 2:7.308 | share |  |  |  |
| 2:7.309 | in the God of |  |  |  |
| 2:7.310 | Israel |  |  |  |
| 2:7.311 | and as for |  |  |  |
| 2:7.312 | what they called |  |  |  |
| 2:7.313 | them |  |  |  |
| 2:7.314 | lights |  |  |  |
| 2:7.315 | subtle |  |  |  |
| 2:7.316 | is not |  |  |  |
| 2:7.317 | but |  |  |  |
| 2:7.318 | to soothe |  |  |  |
| 2:7.319 | (object marker) |  |  |  |
| 2:7.320 | the ear |  |  |  |
| 2:7.321 | we have |  |  |  |
| 2:7.322 | for us |  |  |  |
| 2:7.323 | to ask |  |  |  |
| 2:7.324 | about |  |  |  |
| 2:7.325 | what |  |  |  |
| 2:7.326 | he wrote |  |  |  |
| 2:7.327 | the Rambam |  |  |  |
| 2:7.328 | let him direct |  |  |  |
| 2:7.329 | his heart |  |  |  |
| 2:7.330 | toward |  |  |  |
| 2:7.331 | the Shekhinah |  |  |  |
| 2:7.332 | and pray |  |  |  |
| 2:7.333 | and in the Guide |  |  | Moreh Nevukhim (the Guide for the Perplexed); the Hebrew names it only as 'the Guide' (במורה). |
| 2:7.334 | 20 |  |  | Chapter 20, as the author gives it (כ'). |
| 2:7.335 | that the Shekhinah is |  |  |  |
| 2:7.336 | a light |  |  |  |
| 2:7.337 | created |  |  |  |
| 2:7.338 | and more |  |  |  |
| 2:7.339 | than this |  |  |  |
| 2:7.340 | what |  |  |  |
| 2:7.341 | he wrote |  |  |  |
| 2:7.342 | about |  |  |  |
| 2:7.343 | the glory of |  |  |  |
| 2:7.344 | the LORD |  |  |  |
| 2:7.345 | that it is |  |  |  |
| 2:7.346 | created |  |  |  |
| 2:7.347 | if |  |  |  |
| 2:7.348 | so |  |  |  |
| 2:7.349 | how |  |  |  |
| 2:7.350 | were |  |  |  |
| 2:7.351 | Israel |  |  |  |
| 2:7.352 | going out |  |  |  |
| 2:7.353 | to the opening of |  |  |  |
| 2:7.354 | their tents |  |  |  |
| 2:7.355 | and bowing |  |  |  |
| 2:7.356 | to the created thing |  |  |  |
| 2:7.357 | and more |  |  |  |
| 2:7.358 | than |  |  |  |
| 2:7.359 | this |  |  |  |
| 2:7.360 | what |  |  |  |
| 2:7.361 | he wrote |  |  |  |
| 2:7.362 | that every |  |  |  |
| 2:7.363 | form |  |  |  |
| 2:7.364 | which |  |  |  |
| 2:7.365 | seeing |  |  |  |
| 2:7.366 | the prophets |  |  |  |
| 2:7.367 | and speaking |  |  |  |
| 2:7.368 | with it |  |  |  |
| 2:7.369 | is |  |  |  |
| 2:7.370 | a light |  |  |  |
| 2:7.371 | created |  |  |  |
| 2:7.372 | if so |  |  |  |
| 2:7.373 | there is no |  |  |  |
| 2:7.374 | for us |  |  |  |
| 2:7.375 | any |  |  |  |
| 2:7.376 | prophecy |  |  |  |
| 2:7.377 | through |  |  |  |
| 2:7.378 | the Creator |  |  |  |
| 2:7.379 | but |  |  |  |
| 2:7.380 | all of it |  |  |  |
| 2:7.381 | from the side of |  |  |  |
| 2:7.382 | the created ones |  |  |  |
| 2:7.383 | and if |  |  |  |
| 2:7.384 | it is said |  |  |  |
| 2:7.385 | that |  |  |  |
| 2:7.386 | the created |  |  |  |
| 2:7.387 | attached |  |  |  |
| 2:7.388 | to the Creator |  |  |  |
| 2:7.389 | if so |  |  |  |
| 2:7.390 | in your opinion |  |  |  |
| 2:7.391 | it is |  |  |  |
| 2:7.392 | a force |  |  |  |
| 2:7.393 | in a body |  |  |  |
| 2:7.394 | your answer |  |  |  |
| 2:7.395 | quickly |  |  |  |
| 2:7.396 | will spring up |  |  |  |
| 2:7.397 | (the verse) |  |  | ס' before the verse is an abbreviation the Hebrew does not expand; the verse is Jer 9:5. |
| 2:7.398 | they refuse |  |  |  |
| 2:7.399 | to know |  |  |  |
| 2:7.400 | Me |  |  |  |
| 2:7.401 | says |  |  |  |
| 2:7.402 | the LORD |  |  |  |

Phrases:
- 2:7.1 2:7.2 | Second Reply. | Reply second
- 2:7.3 2:7.4 2:7.5 2:7.6 | I saw all that is written | I-saw (object-marker) all what-is-written
- 2:7.7 2:7.8 | in the letter of the eminent one, | in-the-letter-of the-eminent-one
- 2:7.9 2:7.10 2:7.11 | a tree whose branches are many, | a-tree whose-branches are-many
- 2:7.12 2:7.13 2:7.14 | its height like the height of cedars, | its-height like-the-height-of cedars
- 2:7.15 2:7.16 2:7.17 2:7.18 | After the pardon for the delay, | after the-pardon for the-delay
- 2:7.19 2:7.20 2:7.21 2:7.22 | for the pressure of the times is outstretched, | for the-pressure-of the-times is-outstretched
- 2:7.23 2:7.24 2:7.25 | and every delay is for good. | and-every delay is-for-good
- 2:7.26 2:7.27 2:7.28 | When I looked at the letter, | when-I-looked at the-letter
- 2:7.29 2:7.30 2:7.31 | behold, you opened your mouth, | and-behold you-opened your-mouth
- 2:7.32 2:7.33 | by mentioning divinities, | by-mentioning divinities
- 2:7.34 2:7.35 2:7.36 | thinking that it is against us. | thinking that-it-is against-us
- 2:7.37 2:7.38 2:7.39 2:7.40 2:7.41 | And what are we that you lay upon us, | and-we what that you-lay upon-us
- 2:7.42 2:7.43 2:7.44 2:7.45 | and your scorn does not reach us, | and-your-scorn does-not reach to-us
- 2:7.46 2:7.47 2:7.48 2:7.49 2:7.50 | but only to the sages of Israel, | but only to the-sages-of Israel
- 2:7.51 2:7.52 2:7.53 2:7.54 | whose least is thicker than your loins, | who their-least is-thicker from-your-loins
- 2:7.55 2:7.56 2:7.57 2:7.58 | and all Israel rests on them, | and-all Israel rest on-them
- 2:7.59 2:7.60 2:7.61 | and by their merit the world stands. | and-by-their-merit the-world stands
- 2:7.62 2:7.63 | Woe to you, | woe to-you
- 2:7.64 2:7.65 2:7.66 2:7.67 2:7.68 2:7.69 2:7.70 | for you have brought out against them words that are not according to the law, | for you-have-brought-out against-them words which are-not according-to-the-law
- 2:7.71 2:7.72 2:7.73 | in innocence, and in its place guile. | in-innocence and-in-its-place guile
- 2:7.74 2:7.75 2:7.76 2:7.77 | But your judgment has already been made clear, | but already was-made-clear your-judgment
- 2:7.78 2:7.79 2:7.80 2:7.81 | by the mouth of all the masters of wisdom. | by-the-mouth-of all the-masters-of wisdom
- 2:7.82 2:7.83 2:7.84 | And as for what you wrote about | and-as-for what-you-wrote about
- 2:7.85 2:7.86 | Kisse Eliyahu, | Kisse Eliyahu
- 2:7.87 2:7.88 2:7.89 | that on folio 3 | that on-folio 3
- 2:7.90 2:7.91 2:7.92 2:7.93 2:7.94 | he likens His oneness, may He be blessed, to the other unities. | he-likens His-oneness may-He-be-blessed to-the-other unities
- 2:7.95 2:7.96 2:7.97 2:7.98 | You are already walking in darkness. | already you walk in-darkness
- 2:7.99 2:7.100 2:7.101 2:7.102 2:7.103 | Surely the beginning of his words on folio 20 | surely the-beginning-of his-words folio 20
- 2:7.104 2:7.105 2:7.106 2:7.107 | is that all these likenesses are about | that-all the-likenesses these are-about
- 2:7.108 2:7.109 2:7.110 | holy lights, which | lights holy which
- 2:7.111 2:7.112 2:7.113 | have no grasp | have-no for-them grasp
- 2:7.114 2:7.115 | form and likeness, | and-form and-likeness
- 2:7.116 2:7.117 2:7.118 2:7.119 2:7.120 2:7.121 | neither in intellect nor in knowledge at all and in principle, | neither in-intellect nor in-knowledge at-all and-in-principle
- 2:7.122 2:7.123 2:7.124 2:7.125 2:7.126 2:7.127 2:7.128 2:7.129 | but so that the intellect can receive them, we were permitted to liken. | but so-that that-it-can the-intellect receive them we-were-permitted to-liken
- 2:7.130 2:7.131 2:7.132 2:7.133 | And on folio 4b he wrote | and-on-folio 4 b he-wrote
- 2:7.134 2:7.135 2:7.136 2:7.137 2:7.138 | that all the epithets and attributes are in the manner of | that-all the-epithets and-the-attributes are in-the-manner-of
- 2:7.139 2:7.140 2:7.141 | likeness, in the sefirot and concerning | likeness in-the-sefirot and-concerning
- 2:7.142 2:7.143 | holy, the emanated beings. | holy the-emanated-beings
- 2:7.144 2:7.145 2:7.146 | But in Him, may He be blessed, | but in-Him may-He-be-blessed
- 2:7.147 2:7.148 2:7.149 2:7.150 | there is not in Him any likeness | there-is-not in-Him any likeness
- 2:7.151 2:7.152 2:7.153 2:7.154 2:7.155 | and form, and grasp in thought at all and in principle. | and-form and-grasp in-thought at-all and-in-principle
- 2:7.156 2:7.157 | And as for what you wrote, | and-as-for what-you-wrote
- 2:7.158 2:7.159 2:7.160 2:7.161 | how can we leave the other partzufim | how can-we-leave the-other the-partzufim
- 2:7.162 2:7.163 2:7.164 | and worship Ze'ir Anpin | and-worship Ze'ir Anpin
- 2:7.165 2:7.166 2:7.167 2:7.168 2:7.169 | with the power in him, Ein Sof, blessed be He. | with the-power in-him Ein-Sof blessed-be-He
- 2:7.170 2:7.171 2:7.172 2:7.173 | Rabbi Kisse Eliyahu has already explained | already explained Rabbi Kisse-Eliyahu
- 2:7.174 2:7.175 2:7.176 | on folio 25b, | on-folio 25 b
- 2:7.177 2:7.178 | that when we call upon Ze'ir Anpin, | that-when-we-call upon-Ze'ir-Anpin
- 2:7.179 2:7.180 2:7.181 2:7.182 2:7.183 | we call all the partzufim together, | we call all the-partzufim together
- 2:7.184 2:7.185 2:7.186 2:7.187 | with a soul hidden in them. | with a-soul hidden in-them
- 2:7.188 2:7.189 | And as for what you wrote, | and-as-for what-you-wrote
- 2:7.190 2:7.191 2:7.192 | why do we leave the partzufim | why do-we-leave the-partzufim
- 2:7.193 2:7.194 | closest to Ein Sof | closest to-Ein-Sof
- 2:7.195 2:7.196 | and call upon Ze'ir Anpin. | and-call-upon Ze'ir-Anpin
- 2:7.197 2:7.198 | Because it is known | because it-is-known
- 2:7.199 2:7.200 | that the inner attributes | that-the-attributes inner
- 2:7.201 2:7.202 2:7.203 | some of them judgment, | there-are of-them judgment
- 2:7.204 2:7.205 2:7.206 | and some of them mercy, | and-there-are of-them mercy
- 2:7.207 2:7.208 2:7.209 | and this one includes them all, | and-this-one includes all-of-them
- 2:7.210 2:7.211 2:7.212 | and therefore the Emanator wanted | and-therefore wanted the-Emanator
- 2:7.213 2:7.214 | to act His actions | to-act His-actions
- 2:7.215 2:7.216 2:7.217 2:7.218 | through this vessel. | through the-hand-of a-vessel this
- 2:7.219 2:7.220 2:7.221 2:7.222 2:7.223 | even to the point that you object to the likenesses | even-to-the-point that-you object to the-likenesses
- 2:7.224 2:7.225 | concerning the upper ones, | concerning the-upper-ones
- 2:7.226 2:7.227 2:7.228 2:7.229 2:7.230 | which you have no right to ponder | that-you-have-no for-you right to-ponder them
- 2:7.231 2:7.232 2:7.233 | all the more so to mock, | all-the-more so to-mock
- 2:7.234 2:7.235 2:7.236 2:7.237 | then object to your corporeal hand, | then-object to your-hand the-corporeal
- 2:7.238 2:7.239 2:7.240 2:7.241 | which does all the labors, | which-is does all the-labors
- 2:7.242 2:7.243 2:7.244 2:7.245 2:7.246 | and the head above does not do anything. | and-the-head above does-not do anything
- 2:7.247 2:7.248 2:7.249 2:7.250 2:7.251 | Also, the power of speech, which is your glory, | also the-power-of speech which-is your-glory
- 2:7.252 2:7.253 2:7.254 2:7.255 | above every living creature, | above every living creature
- 2:7.256 2:7.257 2:7.258 2:7.259 2:7.260 | why was it not given above your eyes. | why not was-it-given above your-eyes
- 2:7.261 2:7.262 2:7.263 2:7.264 2:7.265 | But thus His wisdom decreed, may He be blessed. | but thus decreed His-wisdom may-He-be-blessed
- 2:7.266 2:7.267 2:7.268 2:7.269 2:7.270 | And if you have proper counsel, | and-if you-have for-you counsel proper
- 2:7.271 2:7.272 2:7.273 | go to your Creator | go to your-Creator
- 2:7.274 | and let Him counsel us. | and-let-Him-counsel-us
- 2:7.275 2:7.276 2:7.277 | And as for what you wrote, that the words of | and-as-for what-you-wrote that-the-words-of
- 2:7.278 2:7.279 2:7.280 2:7.281 | Kisse Eliyahu are against the Rambam, etc., | Kisse-Eliyahu are-against the-Rambam etc.
- 2:7.282 2:7.283 2:7.284 2:7.285 | who wrote that it is not a force in a body, | who-wrote that-it-is-not a-force in-a-body
- 2:7.286 2:7.287 2:7.288 2:7.289 2:7.290 2:7.291 | and according to the words of Kisse Eliyahu it is a force in a body. | and-according-to the-words-of Kisse-Eliyahu it-is a-force in-a-body
- 2:7.292 2:7.293 2:7.294 | Surely all the Kabbalists | surely all the-Kabbalists
- 2:7.295 2:7.296 | with one mouth | with-one mouth
- 2:7.297 2:7.298 | answer and say | answer and-say
- 2:7.299 2:7.300 2:7.301 2:7.302 | that whoever ascribes corporeality to the sefirot, like you, | that-all those-who-ascribe-corporeality to-the-sefirot like-you
- 2:7.303 2:7.304 2:7.305 | is a denier | is such-a-one a-denier
- 2:7.306 2:7.307 2:7.308 2:7.309 2:7.310 | and has no share in the God of Israel. | and-has-no for-him share in-the-God-of Israel
- 2:7.311 2:7.312 2:7.313 | And as for what they called them | and-as-for what-they-called them
- 2:7.314 2:7.315 | subtle lights, | lights subtle
- 2:7.316 2:7.317 2:7.318 | is only to soothe | is-not but to-soothe
- 2:7.319 2:7.320 | the ear! | (object-marker) the-ear
- 2:7.321 2:7.322 2:7.323 | We have to ask | we-have for-us to-ask
- 2:7.324 2:7.325 2:7.326 2:7.327 | about what the Rambam wrote: | about what he-wrote the-Rambam
- 2:7.328 2:7.329 2:7.330 2:7.331 2:7.332 | "Let him direct his heart toward the Shekhinah and pray." | let-him-direct his-heart toward the-Shekhinah and-pray
- 2:7.333 2:7.334 | And in the Guide, 20, | and-in-the-Guide 20
- 2:7.335 2:7.336 2:7.337 | that the Shekhinah is a created light. | that-the-Shekhinah-is a-light created
- 2:7.338 2:7.339 | And more than this, | and-more than-this
- 2:7.340 2:7.341 2:7.342 | what he wrote about | what he-wrote about
- 2:7.343 2:7.344 2:7.345 2:7.346 | the glory of the LORD, that it is created. | the-glory-of the-LORD that-it-is created
- 2:7.347 2:7.348 | If so, | if so
- 2:7.349 2:7.350 2:7.351 2:7.352 | how did Israel go out | how were Israel going-out
- 2:7.353 2:7.354 | to the opening of their tents | to-the-opening-of their-tents
- 2:7.355 2:7.356 | and bow down to the created thing. | and-bowing to-the-created-thing
- 2:7.357 2:7.358 2:7.359 | And more than this, | and-more than this
- 2:7.360 2:7.361 | what he wrote, | what he-wrote
- 2:7.362 2:7.363 | that every form | that-every form
- 2:7.364 2:7.365 2:7.366 2:7.367 2:7.368 | which the prophets saw and spoke with | which seeing the-prophets and-speaking with-it
- 2:7.369 2:7.370 2:7.371 | is a created light. | is a-light created
- 2:7.372 | If so, | if-so
- 2:7.373 2:7.374 2:7.375 2:7.376 | we have no prophecy | there-is-no for-us any prophecy
- 2:7.377 2:7.378 | through the Creator, | through the-Creator
- 2:7.379 2:7.380 2:7.381 2:7.382 | but all of it from the side of the created ones. | but all-of-it from-the-side-of the-created-ones
- 2:7.383 2:7.384 2:7.385 | And if it is said that | and-if it-is-said that
- 2:7.386 2:7.387 2:7.388 | the created thing was attached to the Creator, | the-created attached to-the-Creator
- 2:7.389 2:7.390 2:7.391 2:7.392 2:7.393 | then, in your opinion, it is a force in a body! | if-so in-your-opinion it-is a-force in-a-body
- 2:7.394 2:7.395 2:7.396 | Your answer will quickly spring up, | your-answer quickly will-spring-up
- 2:7.397 2:7.398 2:7.399 2:7.400 2:7.401 2:7.402 | "They refuse to know Me, says the LORD." | (the-verse) they-refuse to-know Me says the-LORD

Grammar notes:
- תשובה שנית is a heading ('Second Reply'): Chapter 2 is a correspondence, and the writer now answers the questioner in the second person (פערת פיך, כנגדינו).
- אילן שענפיו מרובים, רומו כגובה ארזים is an image in apposition to כתב האלוף. Its last phrase alludes to Amos 2:9 and is unmarked.
- Each 'ומה שכתבת' clause answers one point of the questioner's letter, in the order he gave them.
- שכתב הרמב"ם: יכוין לבו כנגד השכינה ויתפלל is an explicit quotation with the formula 'wrote', so it takes double quotation marks; the Hebrew colon is kept.
- ותשובתך מהרה תצמח alludes to Isa 58:8 and takes no marks. ס' מאנו דעת אותי נאם ה' is a quoted verse with the abbreviation ס' before it, so it takes quotation marks.

Translator's note: The respondent's reply, in his voice. The Kabbalistic terms and the book's abbreviations follow the decisions for the work.

Commentary rulings:
- apparatus: Jer 9:5 (closing words, quoted): מאנו דעת אותי נאם ה'.
- apparatus: Isa 58:8 (allusion): ארכתך מהרה תצמח.
- apparatus: Amos 2:9 (allusion): אשר גבהו כגובה ארזים.
- apparatus: Folio numbers (3, 20, 4b, 25b) are the author's own; no reference is added.
- apparatus: Wikisource base text; not yet collated with the 1931 scan (Q2).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

