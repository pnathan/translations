# Translation pack: milhamot-hashem perek-8-07 (Milchamot ha-Shem 8:9–8:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-07`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **חז"ל** → "our Sages of blessed memory".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **ממלכי** → "of the kings of".
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **השונים** → "the Dualists".
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
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
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אחד ... **ואחד** of two kinds of case → "alike".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- **אבותיהם** for "אבתיכם": follow his wording; the verse in `apparatus`.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 8:9

Hebrew (Wikisource): קה) והדעות הנפסדות הנזכרים עם דעת השונים האומרים שני אלוהות יש אחד פועל הטוב ואחד פועל הרע. אסף בחפניו המסית מחבר הזהר ויקרא להן סתרי תורה. ופועל הטוב אצלו זעיר ונוקביה עם הפרצופים שלמעלה ממנו. ופועל הרע שמאל ונוקביה עם הפרצופים שלהם. וקרא לו שם אדם בליעל. ושם מקומם אחורי פרצופים דקדושה. וקורא להם שם סטרא אחרא, סטרא דקליפה, אל אחר בעל עשר ספירות, ונתן לפועל הרע דין קדימה על פועל הטוב בהיותו משבירת הכלים כידוע, ותלה דבריו בדברי חז"ל שאמרו בבראשית רבה שהיה הקב"ה בונה עולמות ומחריבן, וחיפה דעות זרות אלו בערפלי טוהר וענני חסידות של הבל, וקיים ואימת מציאות אל אחר פועל הרע, היפך דברי רז"ל. יוחס את כל הדעות הנפסדות להתנא הקדוש רשב"י ובנו וחביריו אשר ידענו צדקתם ושלימותם וחכמתם הרבה. ורבותינו האחרונים נמשכו אחריו בעצימת עינים וחשבו אותן שהן דברי קבלה אמתית, ומתחלת אלף השישי אחר שנתגלו על ידי מלך ממלכי הגוים עכו"ם פרו וישרצו הדעות הנפסדות. ובהבלי אדם בעבותות שקר נמשכו אחריהן הרבה מחכמי ישראל מבלי שים לב מאין יצאו הדברים. וינהו אחריהם. וימאסו במשנת אבותיהם אשר קיבל משה רבינו ע"ה, ועמד בהר סיני ארבעים יום וארבעים לילה לחם לא אכל ומים לא שתה כעדות תורתינו הקדושה, ויחללו את המשנה ואת התלמוד הקדוש. ויקראו להם תבנא דאוריתא וקליפה, ומוץ דילה דאתאמר בהון מוץ ותבן פטורים מן המעשר וכו' (טהרת הקדש דף מ"א ב' זהר בראשית) ושפחה כי תירש גברתה, וישליכום אחרי גום! ויצפו לגאולה על ידי עסקם בזהר, ולא ידעו כי היא בעצמה האמונה התפילה הלזו היא הגורמת אריכות גלותינו ושעבודינו בין הגוים כמו שאמר הנביא יחזקאל בשם השי"ת ואומר להם (פי' במצרים על ידי קהת ועמרם) איש שקוצי עיניו השליכו ובגלולי מצרים אל תטמאו ולא שמעו וכו' וכמו שדרשו רז"ל בסוף פרק ערבי פסחים על בזר עמים קרבות יחפצו, מי גרם להם לישראל שיתפזרו בין אוה"ע קריבות שהיו חפצים בהם, שדרשו אותו מלשון נקבה והתחברות להם בדיעותיהם הנפסדות, והוא על דרך מאמר הכתוב ביהושפט מלך יהודה (ד"ה ב' כ) בהתחברך עם אחזיהו פרץ ה' את מעשיך! כי עם ישראל לבדד ישכון, ובגוים לא יתחשב ויתערב באמונתו הטהורה, ולא יתגאל באמונות כוזבות מזוייפות, לפי שאמונת התורה ודעותיה הישרים תנצרנו אם משכיל בה בעינים פקוחות ושכל ישר על דרך שקבלו רז"ל במשנה ובתלמוד בבלי וירושלמי ומדרשי רז"ל הקדמונים.

| id | word | punctuation after |
|---|---|---|
| 8:9.1 | קה | ) |
| 8:9.2 | והדעות |  |
| 8:9.3 | הנפסדות |  |
| 8:9.4 | הנזכרים |  |
| 8:9.5 | עם |  |
| 8:9.6 | דעת |  |
| 8:9.7 | השונים |  |
| 8:9.8 | האומרים |  |
| 8:9.9 | שני |  |
| 8:9.10 | אלוהות |  |
| 8:9.11 | יש |  |
| 8:9.12 | אחד |  |
| 8:9.13 | פועל |  |
| 8:9.14 | הטוב |  |
| 8:9.15 | ואחד |  |
| 8:9.16 | פועל |  |
| 8:9.17 | הרע | . |
| 8:9.18 | אסף |  |
| 8:9.19 | בחפניו |  |
| 8:9.20 | המסית |  |
| 8:9.21 | מחבר |  |
| 8:9.22 | הזהר |  |
| 8:9.23 | ויקרא |  |
| 8:9.24 | להן |  |
| 8:9.25 | סתרי |  |
| 8:9.26 | תורה | . |
| 8:9.27 | ופועל |  |
| 8:9.28 | הטוב |  |
| 8:9.29 | אצלו |  |
| 8:9.30 | זעיר |  |
| 8:9.31 | ונוקביה |  |
| 8:9.32 | עם |  |
| 8:9.33 | הפרצופים |  |
| 8:9.34 | שלמעלה |  |
| 8:9.35 | ממנו | . |
| 8:9.36 | ופועל |  |
| 8:9.37 | הרע |  |
| 8:9.38 | שמאל |  |
| 8:9.39 | ונוקביה |  |
| 8:9.40 | עם |  |
| 8:9.41 | הפרצופים |  |
| 8:9.42 | שלהם | . |
| 8:9.43 | וקרא |  |
| 8:9.44 | לו |  |
| 8:9.45 | שם |  |
| 8:9.46 | אדם |  |
| 8:9.47 | בליעל | . |
| 8:9.48 | ושם |  |
| 8:9.49 | מקומם |  |
| 8:9.50 | אחורי |  |
| 8:9.51 | פרצופים |  |
| 8:9.52 | דקדושה | . |
| 8:9.53 | וקורא |  |
| 8:9.54 | להם |  |
| 8:9.55 | שם |  |
| 8:9.56 | סטרא |  |
| 8:9.57 | אחרא | , |
| 8:9.58 | סטרא |  |
| 8:9.59 | דקליפה | , |
| 8:9.60 | אל |  |
| 8:9.61 | אחר |  |
| 8:9.62 | בעל |  |
| 8:9.63 | עשר |  |
| 8:9.64 | ספירות | , |
| 8:9.65 | ונתן |  |
| 8:9.66 | לפועל |  |
| 8:9.67 | הרע |  |
| 8:9.68 | דין |  |
| 8:9.69 | קדימה |  |
| 8:9.70 | על |  |
| 8:9.71 | פועל |  |
| 8:9.72 | הטוב |  |
| 8:9.73 | בהיותו |  |
| 8:9.74 | משבירת |  |
| 8:9.75 | הכלים |  |
| 8:9.76 | כידוע | , |
| 8:9.77 | ותלה |  |
| 8:9.78 | דבריו |  |
| 8:9.79 | בדברי |  |
| 8:9.80 | חז"ל |  |
| 8:9.81 | שאמרו |  |
| 8:9.82 | בבראשית |  |
| 8:9.83 | רבה |  |
| 8:9.84 | שהיה |  |
| 8:9.85 | הקב"ה |  |
| 8:9.86 | בונה |  |
| 8:9.87 | עולמות |  |
| 8:9.88 | ומחריבן | , |
| 8:9.89 | וחיפה |  |
| 8:9.90 | דעות |  |
| 8:9.91 | זרות |  |
| 8:9.92 | אלו |  |
| 8:9.93 | בערפלי |  |
| 8:9.94 | טוהר |  |
| 8:9.95 | וענני |  |
| 8:9.96 | חסידות |  |
| 8:9.97 | של |  |
| 8:9.98 | הבל | , |
| 8:9.99 | וקיים |  |
| 8:9.100 | ואימת |  |
| 8:9.101 | מציאות |  |
| 8:9.102 | אל |  |
| 8:9.103 | אחר |  |
| 8:9.104 | פועל |  |
| 8:9.105 | הרע | , |
| 8:9.106 | היפך |  |
| 8:9.107 | דברי |  |
| 8:9.108 | רז"ל | . |
| 8:9.109 | יוחס |  |
| 8:9.110 | את |  |
| 8:9.111 | כל |  |
| 8:9.112 | הדעות |  |
| 8:9.113 | הנפסדות |  |
| 8:9.114 | להתנא |  |
| 8:9.115 | הקדוש |  |
| 8:9.116 | רשב"י |  |
| 8:9.117 | ובנו |  |
| 8:9.118 | וחביריו |  |
| 8:9.119 | אשר |  |
| 8:9.120 | ידענו |  |
| 8:9.121 | צדקתם |  |
| 8:9.122 | ושלימותם |  |
| 8:9.123 | וחכמתם |  |
| 8:9.124 | הרבה | . |
| 8:9.125 | ורבותינו |  |
| 8:9.126 | האחרונים |  |
| 8:9.127 | נמשכו |  |
| 8:9.128 | אחריו |  |
| 8:9.129 | בעצימת |  |
| 8:9.130 | עינים |  |
| 8:9.131 | וחשבו |  |
| 8:9.132 | אותן |  |
| 8:9.133 | שהן |  |
| 8:9.134 | דברי |  |
| 8:9.135 | קבלה |  |
| 8:9.136 | אמתית | , |
| 8:9.137 | ומתחלת |  |
| 8:9.138 | אלף |  |
| 8:9.139 | השישי |  |
| 8:9.140 | אחר |  |
| 8:9.141 | שנתגלו |  |
| 8:9.142 | על |  |
| 8:9.143 | ידי |  |
| 8:9.144 | מלך |  |
| 8:9.145 | ממלכי |  |
| 8:9.146 | הגוים |  |
| 8:9.147 | עכו"ם |  |
| 8:9.148 | פרו |  |
| 8:9.149 | וישרצו |  |
| 8:9.150 | הדעות |  |
| 8:9.151 | הנפסדות | . |
| 8:9.152 | ובהבלי |  |
| 8:9.153 | אדם |  |
| 8:9.154 | בעבותות |  |
| 8:9.155 | שקר |  |
| 8:9.156 | נמשכו |  |
| 8:9.157 | אחריהן |  |
| 8:9.158 | הרבה |  |
| 8:9.159 | מחכמי |  |
| 8:9.160 | ישראל |  |
| 8:9.161 | מבלי |  |
| 8:9.162 | שים |  |
| 8:9.163 | לב |  |
| 8:9.164 | מאין |  |
| 8:9.165 | יצאו |  |
| 8:9.166 | הדברים | . |
| 8:9.167 | וינהו |  |
| 8:9.168 | אחריהם | . |
| 8:9.169 | וימאסו |  |
| 8:9.170 | במשנת |  |
| 8:9.171 | אבותיהם |  |
| 8:9.172 | אשר |  |
| 8:9.173 | קיבל |  |
| 8:9.174 | משה |  |
| 8:9.175 | רבינו |  |
| 8:9.176 | ע"ה | , |
| 8:9.177 | ועמד |  |
| 8:9.178 | בהר |  |
| 8:9.179 | סיני |  |
| 8:9.180 | ארבעים |  |
| 8:9.181 | יום |  |
| 8:9.182 | וארבעים |  |
| 8:9.183 | לילה |  |
| 8:9.184 | לחם |  |
| 8:9.185 | לא |  |
| 8:9.186 | אכל |  |
| 8:9.187 | ומים |  |
| 8:9.188 | לא |  |
| 8:9.189 | שתה |  |
| 8:9.190 | כעדות |  |
| 8:9.191 | תורתינו |  |
| 8:9.192 | הקדושה | , |
| 8:9.193 | ויחללו |  |
| 8:9.194 | את |  |
| 8:9.195 | המשנה |  |
| 8:9.196 | ואת |  |
| 8:9.197 | התלמוד |  |
| 8:9.198 | הקדוש | . |
| 8:9.199 | ויקראו |  |
| 8:9.200 | להם |  |
| 8:9.201 | תבנא |  |
| 8:9.202 | דאוריתא |  |
| 8:9.203 | וקליפה | , |
| 8:9.204 | ומוץ |  |
| 8:9.205 | דילה |  |
| 8:9.206 | דאתאמר |  |
| 8:9.207 | בהון |  |
| 8:9.208 | מוץ |  |
| 8:9.209 | ותבן |  |
| 8:9.210 | פטורים |  |
| 8:9.211 | מן |  |
| 8:9.212 | המעשר |  |
| 8:9.213 | וכו | ' |
| 8:9.214 | טהרת |  |
| 8:9.215 | הקדש |  |
| 8:9.216 | דף |  |
| 8:9.217 | מ"א |  |
| 8:9.218 | ב | ' |
| 8:9.219 | זהר |  |
| 8:9.220 | בראשית | ) |
| 8:9.221 | ושפחה |  |
| 8:9.222 | כי |  |
| 8:9.223 | תירש |  |
| 8:9.224 | גברתה | , |
| 8:9.225 | וישליכום |  |
| 8:9.226 | אחרי |  |
| 8:9.227 | גום | ! |
| 8:9.228 | ויצפו |  |
| 8:9.229 | לגאולה |  |
| 8:9.230 | על |  |
| 8:9.231 | ידי |  |
| 8:9.232 | עסקם |  |
| 8:9.233 | בזהר | , |
| 8:9.234 | ולא |  |
| 8:9.235 | ידעו |  |
| 8:9.236 | כי |  |
| 8:9.237 | היא |  |
| 8:9.238 | בעצמה |  |
| 8:9.239 | האמונה |  |
| 8:9.240 | התפילה |  |
| 8:9.241 | הלזו |  |
| 8:9.242 | היא |  |
| 8:9.243 | הגורמת |  |
| 8:9.244 | אריכות |  |
| 8:9.245 | גלותינו |  |
| 8:9.246 | ושעבודינו |  |
| 8:9.247 | בין |  |
| 8:9.248 | הגוים |  |
| 8:9.249 | כמו |  |
| 8:9.250 | שאמר |  |
| 8:9.251 | הנביא |  |
| 8:9.252 | יחזקאל |  |
| 8:9.253 | בשם |  |
| 8:9.254 | השי"ת |  |
| 8:9.255 | ואומר |  |
| 8:9.256 | להם |  |
| 8:9.257 | פי | ' |
| 8:9.258 | במצרים |  |
| 8:9.259 | על |  |
| 8:9.260 | ידי |  |
| 8:9.261 | קהת |  |
| 8:9.262 | ועמרם | ) |
| 8:9.263 | איש |  |
| 8:9.264 | שקוצי |  |
| 8:9.265 | עיניו |  |
| 8:9.266 | השליכו |  |
| 8:9.267 | ובגלולי |  |
| 8:9.268 | מצרים |  |
| 8:9.269 | אל |  |
| 8:9.270 | תטמאו |  |
| 8:9.271 | ולא |  |
| 8:9.272 | שמעו |  |
| 8:9.273 | וכו | ' |
| 8:9.274 | וכמו |  |
| 8:9.275 | שדרשו |  |
| 8:9.276 | רז"ל |  |
| 8:9.277 | בסוף |  |
| 8:9.278 | פרק |  |
| 8:9.279 | ערבי |  |
| 8:9.280 | פסחים |  |
| 8:9.281 | על |  |
| 8:9.282 | בזר |  |
| 8:9.283 | עמים |  |
| 8:9.284 | קרבות |  |
| 8:9.285 | יחפצו | , |
| 8:9.286 | מי |  |
| 8:9.287 | גרם |  |
| 8:9.288 | להם |  |
| 8:9.289 | לישראל |  |
| 8:9.290 | שיתפזרו |  |
| 8:9.291 | בין |  |
| 8:9.292 | אוה"ע |  |
| 8:9.293 | קריבות |  |
| 8:9.294 | שהיו |  |
| 8:9.295 | חפצים |  |
| 8:9.296 | בהם | , |
| 8:9.297 | שדרשו |  |
| 8:9.298 | אותו |  |
| 8:9.299 | מלשון |  |
| 8:9.300 | נקבה |  |
| 8:9.301 | והתחברות |  |
| 8:9.302 | להם |  |
| 8:9.303 | בדיעותיהם |  |
| 8:9.304 | הנפסדות | , |
| 8:9.305 | והוא |  |
| 8:9.306 | על |  |
| 8:9.307 | דרך |  |
| 8:9.308 | מאמר |  |
| 8:9.309 | הכתוב |  |
| 8:9.310 | ביהושפט |  |
| 8:9.311 | מלך |  |
| 8:9.312 | יהודה |  |
| 8:9.313 | ד"ה |  |
| 8:9.314 | ב | ' |
| 8:9.315 | כ | ) |
| 8:9.316 | בהתחברך |  |
| 8:9.317 | עם |  |
| 8:9.318 | אחזיהו |  |
| 8:9.319 | פרץ |  |
| 8:9.320 | ה | ' |
| 8:9.321 | את |  |
| 8:9.322 | מעשיך | ! |
| 8:9.323 | כי |  |
| 8:9.324 | עם |  |
| 8:9.325 | ישראל |  |
| 8:9.326 | לבדד |  |
| 8:9.327 | ישכון | , |
| 8:9.328 | ובגוים |  |
| 8:9.329 | לא |  |
| 8:9.330 | יתחשב |  |
| 8:9.331 | ויתערב |  |
| 8:9.332 | באמונתו |  |
| 8:9.333 | הטהורה | , |
| 8:9.334 | ולא |  |
| 8:9.335 | יתגאל |  |
| 8:9.336 | באמונות |  |
| 8:9.337 | כוזבות |  |
| 8:9.338 | מזוייפות | , |
| 8:9.339 | לפי |  |
| 8:9.340 | שאמונת |  |
| 8:9.341 | התורה |  |
| 8:9.342 | ודעותיה |  |
| 8:9.343 | הישרים |  |
| 8:9.344 | תנצרנו |  |
| 8:9.345 | אם |  |
| 8:9.346 | משכיל |  |
| 8:9.347 | בה |  |
| 8:9.348 | בעינים |  |
| 8:9.349 | פקוחות |  |
| 8:9.350 | ושכל |  |
| 8:9.351 | ישר |  |
| 8:9.352 | על |  |
| 8:9.353 | דרך |  |
| 8:9.354 | שקבלו |  |
| 8:9.355 | רז"ל |  |
| 8:9.356 | במשנה |  |
| 8:9.357 | ובתלמוד |  |
| 8:9.358 | בבלי |  |
| 8:9.359 | וירושלמי |  |
| 8:9.360 | ומדרשי |  |
| 8:9.361 | רז"ל |  |
| 8:9.362 | הקדמונים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

