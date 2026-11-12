# Translation pack: milhamot-hashem perek-6-17 (Milchamot ha-Shem 6:26–6:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-17`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ראב"ע** → "Rabbi Eleazar ben Azariah" in a Talmudic passage, "Ibn Ezra" for
  the commentator (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- מאמיני **השלוש** → "the believers in the Trinity".
- **בא"י** → "in the Land of Israel".

## Terms

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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- סבא **דסבין** → "the Elder of Elders".
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
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
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
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 6:26

Hebrew (Wikisource): אוי לאותה בושה ואוי לאותה כלמה, כי ידע שור קונהו וחמור אבוס בעליו, וישראל לא ידע את אלהיו, ואל איזה צורה ופרצוף מהפרצופים שהמציא מחבר הזהר יתפלל, ויטש אלוה עשהו הראשון וצור לכל, ויבקש אלהים חדשים מקרוב באו אשר אף נגוים עובדי אלילים לא ידעום. ולפי האמור פה חלילה חלילה לחשוב על שום תנא או אמורא כרז"ל אשר ידענו רוב חכמתם וחסידותם וקדושתם שיאמרו דברים כאלה על המשנה והתלמוד אשר הם בעצמם בלו ימיהם בעסק הקדוש הזה, הוא שקלא וטריא דתלמודא, על אבן שתיה, היא המשנה אשר עליה הושת התלמוד הקדוש אשר עיני כל מורי הוראות עליו, ומפיו נבעו כל נהרי נחלי פסקי הלכות, מימות רבותינו הגאונים ועד עתה, וכל משיבי תשובות ממעין המשנה והתלמוד ישאבון.

| id | word | punctuation after |
|---|---|---|
| 6:26.1 | אוי |  |
| 6:26.2 | לאותה |  |
| 6:26.3 | בושה |  |
| 6:26.4 | ואוי |  |
| 6:26.5 | לאותה |  |
| 6:26.6 | כלמה | , |
| 6:26.7 | כי |  |
| 6:26.8 | ידע |  |
| 6:26.9 | שור |  |
| 6:26.10 | קונהו |  |
| 6:26.11 | וחמור |  |
| 6:26.12 | אבוס |  |
| 6:26.13 | בעליו | , |
| 6:26.14 | וישראל |  |
| 6:26.15 | לא |  |
| 6:26.16 | ידע |  |
| 6:26.17 | את |  |
| 6:26.18 | אלהיו | , |
| 6:26.19 | ואל |  |
| 6:26.20 | איזה |  |
| 6:26.21 | צורה |  |
| 6:26.22 | ופרצוף |  |
| 6:26.23 | מהפרצופים |  |
| 6:26.24 | שהמציא |  |
| 6:26.25 | מחבר |  |
| 6:26.26 | הזהר |  |
| 6:26.27 | יתפלל | , |
| 6:26.28 | ויטש |  |
| 6:26.29 | אלוה |  |
| 6:26.30 | עשהו |  |
| 6:26.31 | הראשון |  |
| 6:26.32 | וצור |  |
| 6:26.33 | לכל | , |
| 6:26.34 | ויבקש |  |
| 6:26.35 | אלהים |  |
| 6:26.36 | חדשים |  |
| 6:26.37 | מקרוב |  |
| 6:26.38 | באו |  |
| 6:26.39 | אשר |  |
| 6:26.40 | אף |  |
| 6:26.41 | נגוים |  |
| 6:26.42 | עובדי |  |
| 6:26.43 | אלילים |  |
| 6:26.44 | לא |  |
| 6:26.45 | ידעום | . |
| 6:26.46 | ולפי |  |
| 6:26.47 | האמור |  |
| 6:26.48 | פה |  |
| 6:26.49 | חלילה |  |
| 6:26.50 | חלילה |  |
| 6:26.51 | לחשוב |  |
| 6:26.52 | על |  |
| 6:26.53 | שום |  |
| 6:26.54 | תנא |  |
| 6:26.55 | או |  |
| 6:26.56 | אמורא |  |
| 6:26.57 | כרז"ל |  |
| 6:26.58 | אשר |  |
| 6:26.59 | ידענו |  |
| 6:26.60 | רוב |  |
| 6:26.61 | חכמתם |  |
| 6:26.62 | וחסידותם |  |
| 6:26.63 | וקדושתם |  |
| 6:26.64 | שיאמרו |  |
| 6:26.65 | דברים |  |
| 6:26.66 | כאלה |  |
| 6:26.67 | על |  |
| 6:26.68 | המשנה |  |
| 6:26.69 | והתלמוד |  |
| 6:26.70 | אשר |  |
| 6:26.71 | הם |  |
| 6:26.72 | בעצמם |  |
| 6:26.73 | בלו |  |
| 6:26.74 | ימיהם |  |
| 6:26.75 | בעסק |  |
| 6:26.76 | הקדוש |  |
| 6:26.77 | הזה | , |
| 6:26.78 | הוא |  |
| 6:26.79 | שקלא |  |
| 6:26.80 | וטריא |  |
| 6:26.81 | דתלמודא | , |
| 6:26.82 | על |  |
| 6:26.83 | אבן |  |
| 6:26.84 | שתיה | , |
| 6:26.85 | היא |  |
| 6:26.86 | המשנה |  |
| 6:26.87 | אשר |  |
| 6:26.88 | עליה |  |
| 6:26.89 | הושת |  |
| 6:26.90 | התלמוד |  |
| 6:26.91 | הקדוש |  |
| 6:26.92 | אשר |  |
| 6:26.93 | עיני |  |
| 6:26.94 | כל |  |
| 6:26.95 | מורי |  |
| 6:26.96 | הוראות |  |
| 6:26.97 | עליו | , |
| 6:26.98 | ומפיו |  |
| 6:26.99 | נבעו |  |
| 6:26.100 | כל |  |
| 6:26.101 | נהרי |  |
| 6:26.102 | נחלי |  |
| 6:26.103 | פסקי |  |
| 6:26.104 | הלכות | , |
| 6:26.105 | מימות |  |
| 6:26.106 | רבותינו |  |
| 6:26.107 | הגאונים |  |
| 6:26.108 | ועד |  |
| 6:26.109 | עתה | , |
| 6:26.110 | וכל |  |
| 6:26.111 | משיבי |  |
| 6:26.112 | תשובות |  |
| 6:26.113 | ממעין |  |
| 6:26.114 | המשנה |  |
| 6:26.115 | והתלמוד |  |
| 6:26.116 | ישאבון | . |


## Milchamot ha-Shem 6:27

Hebrew (Wikisource): עב) וקל וחומר הדברים לחשוב על התנא הקדוש רשב"י ע"ה שיכנה את המשנה בשם צלמא דמחתיה אבנא, ובשם אבן משכית, וקליפה, וכיוצא שהרי בירושלמי מבואר בהדיא דרשב"י סבורא ליה שקריאת המשנה חשובה אצלו יותר מן המקרא. ושקול הוא כקריאת שמע בעונתה כמו שנת' לעיל סי' ס"ו. ואיככה נחשוב על התנא הקדוש הזה שידבר סרה על המשנה לומר שכל טמא ואסור ופסול הנז' במשנה מסטרא דיצר הרע. וניחא ליה בהפקירא לזלזל באיסורין שבה, ושעל זה נענש משה רבי' ע"ה להיקבר בארץ טמאה על שלמדנו את המשנה. ושהיה לו למשה ללמדינו תחתיה סודות ורמזים הנוגעים ופוגמים בכבוד אבינו מלכינו וביסודי תורתינו הקדושה ושרשיה להטותינו מעיקרי אמונתינו הטהורה ביחוד השי"ת שהוא הסבה הראשונה. וליסד לנו דת ותורה חדשה להאמין בסבות רבות המשתלשלים זה מזה. וכל אחת נקראת בשם הוי'ה ואדנות ואלהים והקב"ה. והמה שאמרו והיה העולם לפי דבריו, וייחס העבודה והתפלה וקרבנות שהיו ישראל מקריבים בזמן שבית המקדש קיים לאלוה הנברא קצר האפים (ז"א) והוא הסבה האחרונה שבפרצופי האצילות היפך תורתינו הקדושה שבכתב ושבעל פה. ואע"פ שאמר שנגלה לו עתיקא קדישא כמבואר בזהר בראשית דף כ"ב או אליהו ורעיא מהימנא כמבואר בזהר בכמה מקומות. כבר הוזהרנו בתורה שלא נאמין לו ולא נשמע אליו, כמו שנ' כי יקום בקרבך נביא או חלם חלום וכו' כי נביא שקר הוא, ואף אם יעמיד לנו חמה ולבנה כיהושע בגבעון ועמק אילון לא נאמין לו כמו שאמרו רז"ל שגמרא ומביאו הר' הסמ"ג בהקדמת העשין והרמב"ם בחבור ובהקדמת סדר זרעים ובמורה פרק כ"ד משלישי באורך. והטעם כי מנסה ה' אלהינו אותנו כמו שנאמר כי מנסה ה' אלהיכם אתכם לדעת הישכם אוהבים את ה' אלהיכם וגו' ואין נביא רשאי לחדש דבר מעתה, ואף אם בת קול תכריז מן השמים לאמר שמעו בקולו, לא נשמע לו לסור מן הדרך אשר צוה ה' אלהינו אותנו ולזלזל בשום עיקר מעקרי התורה או בשום מצוה ממצותיה, כל שכן זה המסית בלי שום אות ומופות ועוקר ד' עיקרים מיסודי תורתינו הקדושה שכתב הרמב"ם ז"ל כמו שנת' לעיל סי' מ"ח, והנה כפי הנחת מחבר הזהר שכללות הפרצופים הם אחד, איך נפרד העתיק מאריך אנפין ושאר פרצופים ובא לבדו לבית מדרשו של רשב"י ושאל אותו ואמר לו שמעון שמעון מאן הוא דאמר ויאמר אלהים נעשה אדם מאן הוא האי אלהים וכו' (זוהר בראשית כ״ב:ח׳) וכן בכמה מקומות בזהר שנזכר סבא דסבין שהיה בא לגבי רבי אלעזר איך נפרד ובא לבדו הלא הכל אחד? וכבר קדמוני בטענה זו הראב"ע ז"ל על מאמיני השלוש בריש פרשת וירא וז"ל הנה קצת אמרו כי השם שלשה אנשים הוא אחד והוא שלשה ולא יתפרדו, והנה שכחו ויבאו שני המלאכים סדמה ע"כ. ביאור דבריו כי מאמיני השלוש סוברים שהשם ית' הוא שלשה, אב ובן ורוח הקדש והשלשה הם אחד ולכך אמר הכתוב וירא אליו ה', והנה שלשה אנשים, ואלה השלשה הם השם הנז' בפסוק שקודם זה ולא יתפרדו, וטען עליהם הראב"ע האי מצינו שנפרדו שנ' ויבאו שני המלאכים סדומה? גם אנחנו נאמר להמסית מחבר הזהר שאומר וכלא יד ולא מתפרשין לעלמין. איך נפרד עתיק שהוא סבא דסבין ובא לבדו לבית מדרשו של רשב"י? גם מה שאמר שאימא היא שאמרה לאבא נעשה אדם, וגברה ידה ועשאתו על כרחו של אבא, וכשחטא אדם גירשו מגן עדן וגירשה עמו, ואיך נפרדה אימא מאבא, והלא כלא חד ולא מתפרשין? והנה אם היה סבא דסבין נראה לרשב"י ולר' אלעזר בנו כמה פעמים כמו שמבואר בזהר, למה בכה רשב"י ואמר שפחה של בית אבא נזדמן לה מלאך שלשה פעמים. אני אפי' לא פעם אחת? (מעילה דף פ"ד). וגם למה לא הגידו לנו חכמים שרשב"י היה נביא גדול כ"כ יותר משאר נביאים, שהרי היה נגלה לו עתיקא קדישא והוא עומד על עמדו בתוך החבירים בבית המדרש? ולמה אמרו חכמים שחגי זכריה ומלאכי סוף נביאים, ומשם ואילך פסקה נבואה ישראל, הא אשכחן רשב"י ניבא אחר החרבן, והגיד כמה עניינים באלהות, ושהעבודה לבן האלהים קצר האפים מה שלא הודיעו שאר נביאים מימות משה עד מלאכי, ולא נתגלה זה עד האלף הששי?

| id | word | punctuation after |
|---|---|---|
| 6:27.1 | עב | ) |
| 6:27.2 | וקל |  |
| 6:27.3 | וחומר |  |
| 6:27.4 | הדברים |  |
| 6:27.5 | לחשוב |  |
| 6:27.6 | על |  |
| 6:27.7 | התנא |  |
| 6:27.8 | הקדוש |  |
| 6:27.9 | רשב"י |  |
| 6:27.10 | ע"ה |  |
| 6:27.11 | שיכנה |  |
| 6:27.12 | את |  |
| 6:27.13 | המשנה |  |
| 6:27.14 | בשם |  |
| 6:27.15 | צלמא |  |
| 6:27.16 | דמחתיה |  |
| 6:27.17 | אבנא | , |
| 6:27.18 | ובשם |  |
| 6:27.19 | אבן |  |
| 6:27.20 | משכית | , |
| 6:27.21 | וקליפה | , |
| 6:27.22 | וכיוצא |  |
| 6:27.23 | שהרי |  |
| 6:27.24 | בירושלמי |  |
| 6:27.25 | מבואר |  |
| 6:27.26 | בהדיא |  |
| 6:27.27 | דרשב"י |  |
| 6:27.28 | סבורא |  |
| 6:27.29 | ליה |  |
| 6:27.30 | שקריאת |  |
| 6:27.31 | המשנה |  |
| 6:27.32 | חשובה |  |
| 6:27.33 | אצלו |  |
| 6:27.34 | יותר |  |
| 6:27.35 | מן |  |
| 6:27.36 | המקרא | . |
| 6:27.37 | ושקול |  |
| 6:27.38 | הוא |  |
| 6:27.39 | כקריאת |  |
| 6:27.40 | שמע |  |
| 6:27.41 | בעונתה |  |
| 6:27.42 | כמו |  |
| 6:27.43 | שנת | ' |
| 6:27.44 | לעיל |  |
| 6:27.45 | סי | ' |
| 6:27.46 | ס"ו | . |
| 6:27.47 | ואיככה |  |
| 6:27.48 | נחשוב |  |
| 6:27.49 | על |  |
| 6:27.50 | התנא |  |
| 6:27.51 | הקדוש |  |
| 6:27.52 | הזה |  |
| 6:27.53 | שידבר |  |
| 6:27.54 | סרה |  |
| 6:27.55 | על |  |
| 6:27.56 | המשנה |  |
| 6:27.57 | לומר |  |
| 6:27.58 | שכל |  |
| 6:27.59 | טמא |  |
| 6:27.60 | ואסור |  |
| 6:27.61 | ופסול |  |
| 6:27.62 | הנז | ' |
| 6:27.63 | במשנה |  |
| 6:27.64 | מסטרא |  |
| 6:27.65 | דיצר |  |
| 6:27.66 | הרע | . |
| 6:27.67 | וניחא |  |
| 6:27.68 | ליה |  |
| 6:27.69 | בהפקירא |  |
| 6:27.70 | לזלזל |  |
| 6:27.71 | באיסורין |  |
| 6:27.72 | שבה | , |
| 6:27.73 | ושעל |  |
| 6:27.74 | זה |  |
| 6:27.75 | נענש |  |
| 6:27.76 | משה |  |
| 6:27.77 | רבי | ' |
| 6:27.78 | ע"ה |  |
| 6:27.79 | להיקבר |  |
| 6:27.80 | בארץ |  |
| 6:27.81 | טמאה |  |
| 6:27.82 | על |  |
| 6:27.83 | שלמדנו |  |
| 6:27.84 | את |  |
| 6:27.85 | המשנה | . |
| 6:27.86 | ושהיה |  |
| 6:27.87 | לו |  |
| 6:27.88 | למשה |  |
| 6:27.89 | ללמדינו |  |
| 6:27.90 | תחתיה |  |
| 6:27.91 | סודות |  |
| 6:27.92 | ורמזים |  |
| 6:27.93 | הנוגעים |  |
| 6:27.94 | ופוגמים |  |
| 6:27.95 | בכבוד |  |
| 6:27.96 | אבינו |  |
| 6:27.97 | מלכינו |  |
| 6:27.98 | וביסודי |  |
| 6:27.99 | תורתינו |  |
| 6:27.100 | הקדושה |  |
| 6:27.101 | ושרשיה |  |
| 6:27.102 | להטותינו |  |
| 6:27.103 | מעיקרי |  |
| 6:27.104 | אמונתינו |  |
| 6:27.105 | הטהורה |  |
| 6:27.106 | ביחוד |  |
| 6:27.107 | השי"ת |  |
| 6:27.108 | שהוא |  |
| 6:27.109 | הסבה |  |
| 6:27.110 | הראשונה | . |
| 6:27.111 | וליסד |  |
| 6:27.112 | לנו |  |
| 6:27.113 | דת |  |
| 6:27.114 | ותורה |  |
| 6:27.115 | חדשה |  |
| 6:27.116 | להאמין |  |
| 6:27.117 | בסבות |  |
| 6:27.118 | רבות |  |
| 6:27.119 | המשתלשלים |  |
| 6:27.120 | זה |  |
| 6:27.121 | מזה | . |
| 6:27.122 | וכל |  |
| 6:27.123 | אחת |  |
| 6:27.124 | נקראת |  |
| 6:27.125 | בשם |  |
| 6:27.126 | הוי'ה |  |
| 6:27.127 | ואדנות |  |
| 6:27.128 | ואלהים |  |
| 6:27.129 | והקב"ה | . |
| 6:27.130 | והמה |  |
| 6:27.131 | שאמרו |  |
| 6:27.132 | והיה |  |
| 6:27.133 | העולם |  |
| 6:27.134 | לפי |  |
| 6:27.135 | דבריו | , |
| 6:27.136 | וייחס |  |
| 6:27.137 | העבודה |  |
| 6:27.138 | והתפלה |  |
| 6:27.139 | וקרבנות |  |
| 6:27.140 | שהיו |  |
| 6:27.141 | ישראל |  |
| 6:27.142 | מקריבים |  |
| 6:27.143 | בזמן |  |
| 6:27.144 | שבית |  |
| 6:27.145 | המקדש |  |
| 6:27.146 | קיים |  |
| 6:27.147 | לאלוה |  |
| 6:27.148 | הנברא |  |
| 6:27.149 | קצר |  |
| 6:27.150 | האפים |  |
| 6:27.151 | ז"א | ) |
| 6:27.152 | והוא |  |
| 6:27.153 | הסבה |  |
| 6:27.154 | האחרונה |  |
| 6:27.155 | שבפרצופי |  |
| 6:27.156 | האצילות |  |
| 6:27.157 | היפך |  |
| 6:27.158 | תורתינו |  |
| 6:27.159 | הקדושה |  |
| 6:27.160 | שבכתב |  |
| 6:27.161 | ושבעל |  |
| 6:27.162 | פה | . |
| 6:27.163 | ואע"פ |  |
| 6:27.164 | שאמר |  |
| 6:27.165 | שנגלה |  |
| 6:27.166 | לו |  |
| 6:27.167 | עתיקא |  |
| 6:27.168 | קדישא |  |
| 6:27.169 | כמבואר |  |
| 6:27.170 | בזהר |  |
| 6:27.171 | בראשית |  |
| 6:27.172 | דף |  |
| 6:27.173 | כ"ב |  |
| 6:27.174 | או |  |
| 6:27.175 | אליהו |  |
| 6:27.176 | ורעיא |  |
| 6:27.177 | מהימנא |  |
| 6:27.178 | כמבואר |  |
| 6:27.179 | בזהר |  |
| 6:27.180 | בכמה |  |
| 6:27.181 | מקומות | . |
| 6:27.182 | כבר |  |
| 6:27.183 | הוזהרנו |  |
| 6:27.184 | בתורה |  |
| 6:27.185 | שלא |  |
| 6:27.186 | נאמין |  |
| 6:27.187 | לו |  |
| 6:27.188 | ולא |  |
| 6:27.189 | נשמע |  |
| 6:27.190 | אליו | , |
| 6:27.191 | כמו |  |
| 6:27.192 | שנ | ' |
| 6:27.193 | כי |  |
| 6:27.194 | יקום |  |
| 6:27.195 | בקרבך |  |
| 6:27.196 | נביא |  |
| 6:27.197 | או |  |
| 6:27.198 | חלם |  |
| 6:27.199 | חלום |  |
| 6:27.200 | וכו | ' |
| 6:27.201 | כי |  |
| 6:27.202 | נביא |  |
| 6:27.203 | שקר |  |
| 6:27.204 | הוא | , |
| 6:27.205 | ואף |  |
| 6:27.206 | אם |  |
| 6:27.207 | יעמיד |  |
| 6:27.208 | לנו |  |
| 6:27.209 | חמה |  |
| 6:27.210 | ולבנה |  |
| 6:27.211 | כיהושע |  |
| 6:27.212 | בגבעון |  |
| 6:27.213 | ועמק |  |
| 6:27.214 | אילון |  |
| 6:27.215 | לא |  |
| 6:27.216 | נאמין |  |
| 6:27.217 | לו |  |
| 6:27.218 | כמו |  |
| 6:27.219 | שאמרו |  |
| 6:27.220 | רז"ל |  |
| 6:27.221 | שגמרא |  |
| 6:27.222 | ומביאו |  |
| 6:27.223 | הר | ' |
| 6:27.224 | הסמ"ג |  |
| 6:27.225 | בהקדמת |  |
| 6:27.226 | העשין |  |
| 6:27.227 | והרמב"ם |  |
| 6:27.228 | בחבור |  |
| 6:27.229 | ובהקדמת |  |
| 6:27.230 | סדר |  |
| 6:27.231 | זרעים |  |
| 6:27.232 | ובמורה |  |
| 6:27.233 | פרק |  |
| 6:27.234 | כ"ד |  |
| 6:27.235 | משלישי |  |
| 6:27.236 | באורך | . |
| 6:27.237 | והטעם |  |
| 6:27.238 | כי |  |
| 6:27.239 | מנסה |  |
| 6:27.240 | ה | ' |
| 6:27.241 | אלהינו |  |
| 6:27.242 | אותנו |  |
| 6:27.243 | כמו |  |
| 6:27.244 | שנאמר |  |
| 6:27.245 | כי |  |
| 6:27.246 | מנסה |  |
| 6:27.247 | ה | ' |
| 6:27.248 | אלהיכם |  |
| 6:27.249 | אתכם |  |
| 6:27.250 | לדעת |  |
| 6:27.251 | הישכם |  |
| 6:27.252 | אוהבים |  |
| 6:27.253 | את |  |
| 6:27.254 | ה | ' |
| 6:27.255 | אלהיכם |  |
| 6:27.256 | וגו | ' |
| 6:27.257 | ואין |  |
| 6:27.258 | נביא |  |
| 6:27.259 | רשאי |  |
| 6:27.260 | לחדש |  |
| 6:27.261 | דבר |  |
| 6:27.262 | מעתה | , |
| 6:27.263 | ואף |  |
| 6:27.264 | אם |  |
| 6:27.265 | בת |  |
| 6:27.266 | קול |  |
| 6:27.267 | תכריז |  |
| 6:27.268 | מן |  |
| 6:27.269 | השמים |  |
| 6:27.270 | לאמר |  |
| 6:27.271 | שמעו |  |
| 6:27.272 | בקולו | , |
| 6:27.273 | לא |  |
| 6:27.274 | נשמע |  |
| 6:27.275 | לו |  |
| 6:27.276 | לסור |  |
| 6:27.277 | מן |  |
| 6:27.278 | הדרך |  |
| 6:27.279 | אשר |  |
| 6:27.280 | צוה |  |
| 6:27.281 | ה | ' |
| 6:27.282 | אלהינו |  |
| 6:27.283 | אותנו |  |
| 6:27.284 | ולזלזל |  |
| 6:27.285 | בשום |  |
| 6:27.286 | עיקר |  |
| 6:27.287 | מעקרי |  |
| 6:27.288 | התורה |  |
| 6:27.289 | או |  |
| 6:27.290 | בשום |  |
| 6:27.291 | מצוה |  |
| 6:27.292 | ממצותיה | , |
| 6:27.293 | כל |  |
| 6:27.294 | שכן |  |
| 6:27.295 | זה |  |
| 6:27.296 | המסית |  |
| 6:27.297 | בלי |  |
| 6:27.298 | שום |  |
| 6:27.299 | אות |  |
| 6:27.300 | ומופות |  |
| 6:27.301 | ועוקר |  |
| 6:27.302 | ד | ' |
| 6:27.303 | עיקרים |  |
| 6:27.304 | מיסודי |  |
| 6:27.305 | תורתינו |  |
| 6:27.306 | הקדושה |  |
| 6:27.307 | שכתב |  |
| 6:27.308 | הרמב"ם |  |
| 6:27.309 | ז"ל |  |
| 6:27.310 | כמו |  |
| 6:27.311 | שנת | ' |
| 6:27.312 | לעיל |  |
| 6:27.313 | סי | ' |
| 6:27.314 | מ"ח | , |
| 6:27.315 | והנה |  |
| 6:27.316 | כפי |  |
| 6:27.317 | הנחת |  |
| 6:27.318 | מחבר |  |
| 6:27.319 | הזהר |  |
| 6:27.320 | שכללות |  |
| 6:27.321 | הפרצופים |  |
| 6:27.322 | הם |  |
| 6:27.323 | אחד | , |
| 6:27.324 | איך |  |
| 6:27.325 | נפרד |  |
| 6:27.326 | העתיק |  |
| 6:27.327 | מאריך |  |
| 6:27.328 | אנפין |  |
| 6:27.329 | ושאר |  |
| 6:27.330 | פרצופים |  |
| 6:27.331 | ובא |  |
| 6:27.332 | לבדו |  |
| 6:27.333 | לבית |  |
| 6:27.334 | מדרשו |  |
| 6:27.335 | של |  |
| 6:27.336 | רשב"י |  |
| 6:27.337 | ושאל |  |
| 6:27.338 | אותו |  |
| 6:27.339 | ואמר |  |
| 6:27.340 | לו |  |
| 6:27.341 | שמעון |  |
| 6:27.342 | שמעון |  |
| 6:27.343 | מאן |  |
| 6:27.344 | הוא |  |
| 6:27.345 | דאמר |  |
| 6:27.346 | ויאמר |  |
| 6:27.347 | אלהים |  |
| 6:27.348 | נעשה |  |
| 6:27.349 | אדם |  |
| 6:27.350 | מאן |  |
| 6:27.351 | הוא |  |
| 6:27.352 | האי |  |
| 6:27.353 | אלהים |  |
| 6:27.354 | וכו | ' |
| 6:27.355 | זוהר |  |
| 6:27.356 | בראשית |  |
| 6:27.357 | כ״ב:ח׳ | ) |
| 6:27.358 | וכן |  |
| 6:27.359 | בכמה |  |
| 6:27.360 | מקומות |  |
| 6:27.361 | בזהר |  |
| 6:27.362 | שנזכר |  |
| 6:27.363 | סבא |  |
| 6:27.364 | דסבין |  |
| 6:27.365 | שהיה |  |
| 6:27.366 | בא |  |
| 6:27.367 | לגבי |  |
| 6:27.368 | רבי |  |
| 6:27.369 | אלעזר |  |
| 6:27.370 | איך |  |
| 6:27.371 | נפרד |  |
| 6:27.372 | ובא |  |
| 6:27.373 | לבדו |  |
| 6:27.374 | הלא |  |
| 6:27.375 | הכל |  |
| 6:27.376 | אחד | ? |
| 6:27.377 | וכבר |  |
| 6:27.378 | קדמוני |  |
| 6:27.379 | בטענה |  |
| 6:27.380 | זו |  |
| 6:27.381 | הראב"ע |  |
| 6:27.382 | ז"ל |  |
| 6:27.383 | על |  |
| 6:27.384 | מאמיני |  |
| 6:27.385 | השלוש |  |
| 6:27.386 | בריש |  |
| 6:27.387 | פרשת |  |
| 6:27.388 | וירא |  |
| 6:27.389 | וז"ל |  |
| 6:27.390 | הנה |  |
| 6:27.391 | קצת |  |
| 6:27.392 | אמרו |  |
| 6:27.393 | כי |  |
| 6:27.394 | השם |  |
| 6:27.395 | שלשה |  |
| 6:27.396 | אנשים |  |
| 6:27.397 | הוא |  |
| 6:27.398 | אחד |  |
| 6:27.399 | והוא |  |
| 6:27.400 | שלשה |  |
| 6:27.401 | ולא |  |
| 6:27.402 | יתפרדו | , |
| 6:27.403 | והנה |  |
| 6:27.404 | שכחו |  |
| 6:27.405 | ויבאו |  |
| 6:27.406 | שני |  |
| 6:27.407 | המלאכים |  |
| 6:27.408 | סדמה |  |
| 6:27.409 | ע"כ | . |
| 6:27.410 | ביאור |  |
| 6:27.411 | דבריו |  |
| 6:27.412 | כי |  |
| 6:27.413 | מאמיני |  |
| 6:27.414 | השלוש |  |
| 6:27.415 | סוברים |  |
| 6:27.416 | שהשם |  |
| 6:27.417 | ית | ' |
| 6:27.418 | הוא |  |
| 6:27.419 | שלשה | , |
| 6:27.420 | אב |  |
| 6:27.421 | ובן |  |
| 6:27.422 | ורוח |  |
| 6:27.423 | הקדש |  |
| 6:27.424 | והשלשה |  |
| 6:27.425 | הם |  |
| 6:27.426 | אחד |  |
| 6:27.427 | ולכך |  |
| 6:27.428 | אמר |  |
| 6:27.429 | הכתוב |  |
| 6:27.430 | וירא |  |
| 6:27.431 | אליו |  |
| 6:27.432 | ה | ', |
| 6:27.433 | והנה |  |
| 6:27.434 | שלשה |  |
| 6:27.435 | אנשים | , |
| 6:27.436 | ואלה |  |
| 6:27.437 | השלשה |  |
| 6:27.438 | הם |  |
| 6:27.439 | השם |  |
| 6:27.440 | הנז | ' |
| 6:27.441 | בפסוק |  |
| 6:27.442 | שקודם |  |
| 6:27.443 | זה |  |
| 6:27.444 | ולא |  |
| 6:27.445 | יתפרדו | , |
| 6:27.446 | וטען |  |
| 6:27.447 | עליהם |  |
| 6:27.448 | הראב"ע |  |
| 6:27.449 | האי |  |
| 6:27.450 | מצינו |  |
| 6:27.451 | שנפרדו |  |
| 6:27.452 | שנ | ' |
| 6:27.453 | ויבאו |  |
| 6:27.454 | שני |  |
| 6:27.455 | המלאכים |  |
| 6:27.456 | סדומה | ? |
| 6:27.457 | גם |  |
| 6:27.458 | אנחנו |  |
| 6:27.459 | נאמר |  |
| 6:27.460 | להמסית |  |
| 6:27.461 | מחבר |  |
| 6:27.462 | הזהר |  |
| 6:27.463 | שאומר |  |
| 6:27.464 | וכלא |  |
| 6:27.465 | יד |  |
| 6:27.466 | ולא |  |
| 6:27.467 | מתפרשין |  |
| 6:27.468 | לעלמין | . |
| 6:27.469 | איך |  |
| 6:27.470 | נפרד |  |
| 6:27.471 | עתיק |  |
| 6:27.472 | שהוא |  |
| 6:27.473 | סבא |  |
| 6:27.474 | דסבין |  |
| 6:27.475 | ובא |  |
| 6:27.476 | לבדו |  |
| 6:27.477 | לבית |  |
| 6:27.478 | מדרשו |  |
| 6:27.479 | של |  |
| 6:27.480 | רשב"י | ? |
| 6:27.481 | גם |  |
| 6:27.482 | מה |  |
| 6:27.483 | שאמר |  |
| 6:27.484 | שאימא |  |
| 6:27.485 | היא |  |
| 6:27.486 | שאמרה |  |
| 6:27.487 | לאבא |  |
| 6:27.488 | נעשה |  |
| 6:27.489 | אדם | , |
| 6:27.490 | וגברה |  |
| 6:27.491 | ידה |  |
| 6:27.492 | ועשאתו |  |
| 6:27.493 | על |  |
| 6:27.494 | כרחו |  |
| 6:27.495 | של |  |
| 6:27.496 | אבא | , |
| 6:27.497 | וכשחטא |  |
| 6:27.498 | אדם |  |
| 6:27.499 | גירשו |  |
| 6:27.500 | מגן |  |
| 6:27.501 | עדן |  |
| 6:27.502 | וגירשה |  |
| 6:27.503 | עמו | , |
| 6:27.504 | ואיך |  |
| 6:27.505 | נפרדה |  |
| 6:27.506 | אימא |  |
| 6:27.507 | מאבא | , |
| 6:27.508 | והלא |  |
| 6:27.509 | כלא |  |
| 6:27.510 | חד |  |
| 6:27.511 | ולא |  |
| 6:27.512 | מתפרשין | ? |
| 6:27.513 | והנה |  |
| 6:27.514 | אם |  |
| 6:27.515 | היה |  |
| 6:27.516 | סבא |  |
| 6:27.517 | דסבין |  |
| 6:27.518 | נראה |  |
| 6:27.519 | לרשב"י |  |
| 6:27.520 | ולר | ' |
| 6:27.521 | אלעזר |  |
| 6:27.522 | בנו |  |
| 6:27.523 | כמה |  |
| 6:27.524 | פעמים |  |
| 6:27.525 | כמו |  |
| 6:27.526 | שמבואר |  |
| 6:27.527 | בזהר | , |
| 6:27.528 | למה |  |
| 6:27.529 | בכה |  |
| 6:27.530 | רשב"י |  |
| 6:27.531 | ואמר |  |
| 6:27.532 | שפחה |  |
| 6:27.533 | של |  |
| 6:27.534 | בית |  |
| 6:27.535 | אבא |  |
| 6:27.536 | נזדמן |  |
| 6:27.537 | לה |  |
| 6:27.538 | מלאך |  |
| 6:27.539 | שלשה |  |
| 6:27.540 | פעמים | . |
| 6:27.541 | אני |  |
| 6:27.542 | אפי | ' |
| 6:27.543 | לא |  |
| 6:27.544 | פעם |  |
| 6:27.545 | אחת | ? |
| 6:27.546 | מעילה |  |
| 6:27.547 | דף |  |
| 6:27.548 | פ"ד | ). |
| 6:27.549 | וגם |  |
| 6:27.550 | למה |  |
| 6:27.551 | לא |  |
| 6:27.552 | הגידו |  |
| 6:27.553 | לנו |  |
| 6:27.554 | חכמים |  |
| 6:27.555 | שרשב"י |  |
| 6:27.556 | היה |  |
| 6:27.557 | נביא |  |
| 6:27.558 | גדול |  |
| 6:27.559 | כ"כ |  |
| 6:27.560 | יותר |  |
| 6:27.561 | משאר |  |
| 6:27.562 | נביאים | , |
| 6:27.563 | שהרי |  |
| 6:27.564 | היה |  |
| 6:27.565 | נגלה |  |
| 6:27.566 | לו |  |
| 6:27.567 | עתיקא |  |
| 6:27.568 | קדישא |  |
| 6:27.569 | והוא |  |
| 6:27.570 | עומד |  |
| 6:27.571 | על |  |
| 6:27.572 | עמדו |  |
| 6:27.573 | בתוך |  |
| 6:27.574 | החבירים |  |
| 6:27.575 | בבית |  |
| 6:27.576 | המדרש | ? |
| 6:27.577 | ולמה |  |
| 6:27.578 | אמרו |  |
| 6:27.579 | חכמים |  |
| 6:27.580 | שחגי |  |
| 6:27.581 | זכריה |  |
| 6:27.582 | ומלאכי |  |
| 6:27.583 | סוף |  |
| 6:27.584 | נביאים | , |
| 6:27.585 | ומשם |  |
| 6:27.586 | ואילך |  |
| 6:27.587 | פסקה |  |
| 6:27.588 | נבואה |  |
| 6:27.589 | ישראל | , |
| 6:27.590 | הא |  |
| 6:27.591 | אשכחן |  |
| 6:27.592 | רשב"י |  |
| 6:27.593 | ניבא |  |
| 6:27.594 | אחר |  |
| 6:27.595 | החרבן | , |
| 6:27.596 | והגיד |  |
| 6:27.597 | כמה |  |
| 6:27.598 | עניינים |  |
| 6:27.599 | באלהות | , |
| 6:27.600 | ושהעבודה |  |
| 6:27.601 | לבן |  |
| 6:27.602 | האלהים |  |
| 6:27.603 | קצר |  |
| 6:27.604 | האפים |  |
| 6:27.605 | מה |  |
| 6:27.606 | שלא |  |
| 6:27.607 | הודיעו |  |
| 6:27.608 | שאר |  |
| 6:27.609 | נביאים |  |
| 6:27.610 | מימות |  |
| 6:27.611 | משה |  |
| 6:27.612 | עד |  |
| 6:27.613 | מלאכי | , |
| 6:27.614 | ולא |  |
| 6:27.615 | נתגלה |  |
| 6:27.616 | זה |  |
| 6:27.617 | עד |  |
| 6:27.618 | האלף |  |
| 6:27.619 | הששי | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

