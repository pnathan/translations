# Translation pack: milhamot-hashem perek-8-20 (Milchamot ha-Shem 8:24–8:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-20`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- השם **הנכבד** → "the glorious Name".
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
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 8:24

Hebrew (Wikisource): קיז) וכלל הדבר כי הפילסוף המסית מחבר הזהר שם רעיונותיו ולבו בשאלת מה למעלה ומה למטה וכו' והפקיר איסור זה הנזכר במשנה, ולא הניח אלא מה למעלה מכיפת עיגול הא"ס אשר לדעתו נצתמצם למעלה סביב, ואע"פ שתמצא אותו אומר דאית אתר דלא קימא לשאלה ביה, אין זה אלא לגנוב את העין כאלו הוא מאמין באיסור זה, והראיה שהרי דבר במהות האלוה וצמצומו שנעשה ככדור עגול חלול, וכיון שדבר באין סוף שהוא למעלה מן הכל, איך יהיה אסור לשאול ולדבר במה שלמטה ממנו תוך חללו? וכבר גילה דעתו בהדיא, שכל אסור ופסול וטמא הנזכרים במשנה הם מסטרא דיצר הרע, ולכך נהג בזה הפקר ודבר במהות יוצר הכל ואיך ציורו ואופן חנייתו, ועשה לו גופים ופרצופים קבועים שונים זה מזה בכמותם ואיכותם ומקומם, זה למעלה וזה למטה, וזה ארוך וזה קצר כמו שיראה המעיין בו ובעץ חיים ופריו לכהרח"ו, ושבש ובלבל אמונת ישראל הטהורה המקובלת בידם, והשוה אותה לאמונת העכו"ם. ורבים מעמנו נמשכו אחריו לתומם בהאמנים שמהתנא הקדוש רשב"י יצאו הדברים הכבושים אשר היו קבורים, ולא שמו אל לבם איך אפשר לקבל תורה חדשה מיד עכו"ם, ואיך נקברה בארץ בארגז, ומי קברה? ולמה נתגלית בימי רשב"י וחזרה ונקברה ונשכח זכרה מכל היהודים? אשר על כן חלילה לכל משכיל בתורת ה' שבכתב ושבעל פה להאמין שדברים כאלה יצאו מפי רשב"י ע"ה ולא משום תנא או אמורא, כי אדרבא מצינו ראינו שהיו מרחיקים אלו הדעות הרעות ומתנגדים להם, וכמו שכבר הבאתי קצת מאמרים שמהם נלמד התנגדותם לאלו הדעות החדשות הנזכרים בקבלה החדשה, דממה נפשך הואיל ויש לפרצופים דין אלהות כמו שכתב הרש"ב בגליון פרשת בהר דף ק"ט והרב משנת חסידים סוף מסכת העשייה, ובית קדש הקדשים דף נ"ו ע"ב! נמצאנו מאמינים ועובדים לאלוהות רבים נבראים שכל אחד מהם הוא סבה למה שאחריו כמו שכתב לעיל בשם זוהר בראשית דף כ"ב שהוא קורא לכל אחד מהם בשם הוי'ה ואדנות ואלהים וקב"ה ושאר שמות המיוחדים, עד נדר שכשישמע האיש הישראלי את חבירו מזכיר את השם הנכבד הויה או אלהים או הקב"ה, צריך השומע לשאול איזה הוא הקב"ה שאתה מזכיר, אם עתיק, או אריך, או אבא, או אימא, או זעיר, או נוקבא דיליה, כמו שעשה הוא עצמו מחבר הזהר בפסוק ויאמר אלהים נעשה אדם, ואמר מאן אלהים, כדלעיל סי' כ'! כי האלהים העליון לא תוכל לומר (לפי דבריהם) כי הוא אין לו שם ולא נקודה, ולא תוכל לכנותו בשם מן השמות הקדושים כמו שכתבו המקובלים החדשים, ונמצאנו מאמינים באלהים רבים נבראים ח"ו.

| id | word | punctuation after |
|---|---|---|
| 8:24.1 | קיז | ) |
| 8:24.2 | וכלל |  |
| 8:24.3 | הדבר |  |
| 8:24.4 | כי |  |
| 8:24.5 | הפילסוף |  |
| 8:24.6 | המסית |  |
| 8:24.7 | מחבר |  |
| 8:24.8 | הזהר |  |
| 8:24.9 | שם |  |
| 8:24.10 | רעיונותיו |  |
| 8:24.11 | ולבו |  |
| 8:24.12 | בשאלת |  |
| 8:24.13 | מה |  |
| 8:24.14 | למעלה |  |
| 8:24.15 | ומה |  |
| 8:24.16 | למטה |  |
| 8:24.17 | וכו | ' |
| 8:24.18 | והפקיר |  |
| 8:24.19 | איסור |  |
| 8:24.20 | זה |  |
| 8:24.21 | הנזכר |  |
| 8:24.22 | במשנה | , |
| 8:24.23 | ולא |  |
| 8:24.24 | הניח |  |
| 8:24.25 | אלא |  |
| 8:24.26 | מה |  |
| 8:24.27 | למעלה |  |
| 8:24.28 | מכיפת |  |
| 8:24.29 | עיגול |  |
| 8:24.30 | הא"ס |  |
| 8:24.31 | אשר |  |
| 8:24.32 | לדעתו |  |
| 8:24.33 | נצתמצם |  |
| 8:24.34 | למעלה |  |
| 8:24.35 | סביב | , |
| 8:24.36 | ואע"פ |  |
| 8:24.37 | שתמצא |  |
| 8:24.38 | אותו |  |
| 8:24.39 | אומר |  |
| 8:24.40 | דאית |  |
| 8:24.41 | אתר |  |
| 8:24.42 | דלא |  |
| 8:24.43 | קימא |  |
| 8:24.44 | לשאלה |  |
| 8:24.45 | ביה | , |
| 8:24.46 | אין |  |
| 8:24.47 | זה |  |
| 8:24.48 | אלא |  |
| 8:24.49 | לגנוב |  |
| 8:24.50 | את |  |
| 8:24.51 | העין |  |
| 8:24.52 | כאלו |  |
| 8:24.53 | הוא |  |
| 8:24.54 | מאמין |  |
| 8:24.55 | באיסור |  |
| 8:24.56 | זה | , |
| 8:24.57 | והראיה |  |
| 8:24.58 | שהרי |  |
| 8:24.59 | דבר |  |
| 8:24.60 | במהות |  |
| 8:24.61 | האלוה |  |
| 8:24.62 | וצמצומו |  |
| 8:24.63 | שנעשה |  |
| 8:24.64 | ככדור |  |
| 8:24.65 | עגול |  |
| 8:24.66 | חלול | , |
| 8:24.67 | וכיון |  |
| 8:24.68 | שדבר |  |
| 8:24.69 | באין |  |
| 8:24.70 | סוף |  |
| 8:24.71 | שהוא |  |
| 8:24.72 | למעלה |  |
| 8:24.73 | מן |  |
| 8:24.74 | הכל | , |
| 8:24.75 | איך |  |
| 8:24.76 | יהיה |  |
| 8:24.77 | אסור |  |
| 8:24.78 | לשאול |  |
| 8:24.79 | ולדבר |  |
| 8:24.80 | במה |  |
| 8:24.81 | שלמטה |  |
| 8:24.82 | ממנו |  |
| 8:24.83 | תוך |  |
| 8:24.84 | חללו | ? |
| 8:24.85 | וכבר |  |
| 8:24.86 | גילה |  |
| 8:24.87 | דעתו |  |
| 8:24.88 | בהדיא | , |
| 8:24.89 | שכל |  |
| 8:24.90 | אסור |  |
| 8:24.91 | ופסול |  |
| 8:24.92 | וטמא |  |
| 8:24.93 | הנזכרים |  |
| 8:24.94 | במשנה |  |
| 8:24.95 | הם |  |
| 8:24.96 | מסטרא |  |
| 8:24.97 | דיצר |  |
| 8:24.98 | הרע | , |
| 8:24.99 | ולכך |  |
| 8:24.100 | נהג |  |
| 8:24.101 | בזה |  |
| 8:24.102 | הפקר |  |
| 8:24.103 | ודבר |  |
| 8:24.104 | במהות |  |
| 8:24.105 | יוצר |  |
| 8:24.106 | הכל |  |
| 8:24.107 | ואיך |  |
| 8:24.108 | ציורו |  |
| 8:24.109 | ואופן |  |
| 8:24.110 | חנייתו | , |
| 8:24.111 | ועשה |  |
| 8:24.112 | לו |  |
| 8:24.113 | גופים |  |
| 8:24.114 | ופרצופים |  |
| 8:24.115 | קבועים |  |
| 8:24.116 | שונים |  |
| 8:24.117 | זה |  |
| 8:24.118 | מזה |  |
| 8:24.119 | בכמותם |  |
| 8:24.120 | ואיכותם |  |
| 8:24.121 | ומקומם | , |
| 8:24.122 | זה |  |
| 8:24.123 | למעלה |  |
| 8:24.124 | וזה |  |
| 8:24.125 | למטה | , |
| 8:24.126 | וזה |  |
| 8:24.127 | ארוך |  |
| 8:24.128 | וזה |  |
| 8:24.129 | קצר |  |
| 8:24.130 | כמו |  |
| 8:24.131 | שיראה |  |
| 8:24.132 | המעיין |  |
| 8:24.133 | בו |  |
| 8:24.134 | ובעץ |  |
| 8:24.135 | חיים |  |
| 8:24.136 | ופריו |  |
| 8:24.137 | לכהרח"ו | , |
| 8:24.138 | ושבש |  |
| 8:24.139 | ובלבל |  |
| 8:24.140 | אמונת |  |
| 8:24.141 | ישראל |  |
| 8:24.142 | הטהורה |  |
| 8:24.143 | המקובלת |  |
| 8:24.144 | בידם | , |
| 8:24.145 | והשוה |  |
| 8:24.146 | אותה |  |
| 8:24.147 | לאמונת |  |
| 8:24.148 | העכו"ם | . |
| 8:24.149 | ורבים |  |
| 8:24.150 | מעמנו |  |
| 8:24.151 | נמשכו |  |
| 8:24.152 | אחריו |  |
| 8:24.153 | לתומם |  |
| 8:24.154 | בהאמנים |  |
| 8:24.155 | שמהתנא |  |
| 8:24.156 | הקדוש |  |
| 8:24.157 | רשב"י |  |
| 8:24.158 | יצאו |  |
| 8:24.159 | הדברים |  |
| 8:24.160 | הכבושים |  |
| 8:24.161 | אשר |  |
| 8:24.162 | היו |  |
| 8:24.163 | קבורים | , |
| 8:24.164 | ולא |  |
| 8:24.165 | שמו |  |
| 8:24.166 | אל |  |
| 8:24.167 | לבם |  |
| 8:24.168 | איך |  |
| 8:24.169 | אפשר |  |
| 8:24.170 | לקבל |  |
| 8:24.171 | תורה |  |
| 8:24.172 | חדשה |  |
| 8:24.173 | מיד |  |
| 8:24.174 | עכו"ם | , |
| 8:24.175 | ואיך |  |
| 8:24.176 | נקברה |  |
| 8:24.177 | בארץ |  |
| 8:24.178 | בארגז | , |
| 8:24.179 | ומי |  |
| 8:24.180 | קברה | ? |
| 8:24.181 | ולמה |  |
| 8:24.182 | נתגלית |  |
| 8:24.183 | בימי |  |
| 8:24.184 | רשב"י |  |
| 8:24.185 | וחזרה |  |
| 8:24.186 | ונקברה |  |
| 8:24.187 | ונשכח |  |
| 8:24.188 | זכרה |  |
| 8:24.189 | מכל |  |
| 8:24.190 | היהודים | ? |
| 8:24.191 | אשר |  |
| 8:24.192 | על |  |
| 8:24.193 | כן |  |
| 8:24.194 | חלילה |  |
| 8:24.195 | לכל |  |
| 8:24.196 | משכיל |  |
| 8:24.197 | בתורת |  |
| 8:24.198 | ה | ' |
| 8:24.199 | שבכתב |  |
| 8:24.200 | ושבעל |  |
| 8:24.201 | פה |  |
| 8:24.202 | להאמין |  |
| 8:24.203 | שדברים |  |
| 8:24.204 | כאלה |  |
| 8:24.205 | יצאו |  |
| 8:24.206 | מפי |  |
| 8:24.207 | רשב"י |  |
| 8:24.208 | ע"ה |  |
| 8:24.209 | ולא |  |
| 8:24.210 | משום |  |
| 8:24.211 | תנא |  |
| 8:24.212 | או |  |
| 8:24.213 | אמורא | , |
| 8:24.214 | כי |  |
| 8:24.215 | אדרבא |  |
| 8:24.216 | מצינו |  |
| 8:24.217 | ראינו |  |
| 8:24.218 | שהיו |  |
| 8:24.219 | מרחיקים |  |
| 8:24.220 | אלו |  |
| 8:24.221 | הדעות |  |
| 8:24.222 | הרעות |  |
| 8:24.223 | ומתנגדים |  |
| 8:24.224 | להם | , |
| 8:24.225 | וכמו |  |
| 8:24.226 | שכבר |  |
| 8:24.227 | הבאתי |  |
| 8:24.228 | קצת |  |
| 8:24.229 | מאמרים |  |
| 8:24.230 | שמהם |  |
| 8:24.231 | נלמד |  |
| 8:24.232 | התנגדותם |  |
| 8:24.233 | לאלו |  |
| 8:24.234 | הדעות |  |
| 8:24.235 | החדשות |  |
| 8:24.236 | הנזכרים |  |
| 8:24.237 | בקבלה |  |
| 8:24.238 | החדשה | , |
| 8:24.239 | דממה |  |
| 8:24.240 | נפשך |  |
| 8:24.241 | הואיל |  |
| 8:24.242 | ויש |  |
| 8:24.243 | לפרצופים |  |
| 8:24.244 | דין |  |
| 8:24.245 | אלהות |  |
| 8:24.246 | כמו |  |
| 8:24.247 | שכתב |  |
| 8:24.248 | הרש"ב |  |
| 8:24.249 | בגליון |  |
| 8:24.250 | פרשת |  |
| 8:24.251 | בהר |  |
| 8:24.252 | דף |  |
| 8:24.253 | ק"ט |  |
| 8:24.254 | והרב |  |
| 8:24.255 | משנת |  |
| 8:24.256 | חסידים |  |
| 8:24.257 | סוף |  |
| 8:24.258 | מסכת |  |
| 8:24.259 | העשייה | , |
| 8:24.260 | ובית |  |
| 8:24.261 | קדש |  |
| 8:24.262 | הקדשים |  |
| 8:24.263 | דף |  |
| 8:24.264 | נ"ו |  |
| 8:24.265 | ע"ב | ! |
| 8:24.266 | נמצאנו |  |
| 8:24.267 | מאמינים |  |
| 8:24.268 | ועובדים |  |
| 8:24.269 | לאלוהות |  |
| 8:24.270 | רבים |  |
| 8:24.271 | נבראים |  |
| 8:24.272 | שכל |  |
| 8:24.273 | אחד |  |
| 8:24.274 | מהם |  |
| 8:24.275 | הוא |  |
| 8:24.276 | סבה |  |
| 8:24.277 | למה |  |
| 8:24.278 | שאחריו |  |
| 8:24.279 | כמו |  |
| 8:24.280 | שכתב |  |
| 8:24.281 | לעיל |  |
| 8:24.282 | בשם |  |
| 8:24.283 | זוהר |  |
| 8:24.284 | בראשית |  |
| 8:24.285 | דף |  |
| 8:24.286 | כ"ב |  |
| 8:24.287 | שהוא |  |
| 8:24.288 | קורא |  |
| 8:24.289 | לכל |  |
| 8:24.290 | אחד |  |
| 8:24.291 | מהם |  |
| 8:24.292 | בשם |  |
| 8:24.293 | הוי'ה |  |
| 8:24.294 | ואדנות |  |
| 8:24.295 | ואלהים |  |
| 8:24.296 | וקב"ה |  |
| 8:24.297 | ושאר |  |
| 8:24.298 | שמות |  |
| 8:24.299 | המיוחדים | , |
| 8:24.300 | עד |  |
| 8:24.301 | נדר |  |
| 8:24.302 | שכשישמע |  |
| 8:24.303 | האיש |  |
| 8:24.304 | הישראלי |  |
| 8:24.305 | את |  |
| 8:24.306 | חבירו |  |
| 8:24.307 | מזכיר |  |
| 8:24.308 | את |  |
| 8:24.309 | השם |  |
| 8:24.310 | הנכבד |  |
| 8:24.311 | הויה |  |
| 8:24.312 | או |  |
| 8:24.313 | אלהים |  |
| 8:24.314 | או |  |
| 8:24.315 | הקב"ה | , |
| 8:24.316 | צריך |  |
| 8:24.317 | השומע |  |
| 8:24.318 | לשאול |  |
| 8:24.319 | איזה |  |
| 8:24.320 | הוא |  |
| 8:24.321 | הקב"ה |  |
| 8:24.322 | שאתה |  |
| 8:24.323 | מזכיר | , |
| 8:24.324 | אם |  |
| 8:24.325 | עתיק | , |
| 8:24.326 | או |  |
| 8:24.327 | אריך | , |
| 8:24.328 | או |  |
| 8:24.329 | אבא | , |
| 8:24.330 | או |  |
| 8:24.331 | אימא | , |
| 8:24.332 | או |  |
| 8:24.333 | זעיר | , |
| 8:24.334 | או |  |
| 8:24.335 | נוקבא |  |
| 8:24.336 | דיליה | , |
| 8:24.337 | כמו |  |
| 8:24.338 | שעשה |  |
| 8:24.339 | הוא |  |
| 8:24.340 | עצמו |  |
| 8:24.341 | מחבר |  |
| 8:24.342 | הזהר |  |
| 8:24.343 | בפסוק |  |
| 8:24.344 | ויאמר |  |
| 8:24.345 | אלהים |  |
| 8:24.346 | נעשה |  |
| 8:24.347 | אדם | , |
| 8:24.348 | ואמר |  |
| 8:24.349 | מאן |  |
| 8:24.350 | אלהים | , |
| 8:24.351 | כדלעיל |  |
| 8:24.352 | סי | ' |
| 8:24.353 | כ | '! |
| 8:24.354 | כי |  |
| 8:24.355 | האלהים |  |
| 8:24.356 | העליון |  |
| 8:24.357 | לא |  |
| 8:24.358 | תוכל |  |
| 8:24.359 | לומר |  |
| 8:24.360 | לפי |  |
| 8:24.361 | דבריהם | ) |
| 8:24.362 | כי |  |
| 8:24.363 | הוא |  |
| 8:24.364 | אין |  |
| 8:24.365 | לו |  |
| 8:24.366 | שם |  |
| 8:24.367 | ולא |  |
| 8:24.368 | נקודה | , |
| 8:24.369 | ולא |  |
| 8:24.370 | תוכל |  |
| 8:24.371 | לכנותו |  |
| 8:24.372 | בשם |  |
| 8:24.373 | מן |  |
| 8:24.374 | השמות |  |
| 8:24.375 | הקדושים |  |
| 8:24.376 | כמו |  |
| 8:24.377 | שכתבו |  |
| 8:24.378 | המקובלים |  |
| 8:24.379 | החדשים | , |
| 8:24.380 | ונמצאנו |  |
| 8:24.381 | מאמינים |  |
| 8:24.382 | באלהים |  |
| 8:24.383 | רבים |  |
| 8:24.384 | נבראים |  |
| 8:24.385 | ח"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

