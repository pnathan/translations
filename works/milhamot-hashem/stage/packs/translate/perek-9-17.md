# Translation pack: milhamot-hashem perek-9-17 (Milchamot ha-Shem 9:20–9:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-17`.

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
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
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
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **אפיקורוס / אפיקורסים** → "apikoros / apikorsim".
- כופר **בעיקר** → "denies the fundamental principle".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
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
- **תא חזי** → "come and see".
- מי **אתי** (ketiv): translated, the qere noted.
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 9:20

Hebrew (Wikisource): קלח) את כל אלה עשתה יד הפילסוף המסית מחבר הזהר להלביש את קוראיו גאה וגאון אף שאינם יודעים מכל דיני תורתינו הקדושה מאומה, ואף מכל מה שלומדים בו שכן הוא הבטיח להם שכר רב בלמדם אותו אף בלי הבין, ולכן הם חושבים עצמם בני עלייה, ומבזים המה את העוסקים במשנה ובתלמוד שהם כתרה של תורה כעדות רז"ל במס' מגלה והבאתי לעיל סי' קל"ו, ובעיניהם המשנה והתלמוד חשובים לתבן וקש, והעוסקים בהם כבהמות יער האוכלות חציר ותבן, וכעופות כדאיתא בזהר פרשת תצא הובא לעיל סי' ס"ח, וכן חשב מהרח"ו בשבחיו דף י"ב, ולא ידעו ולא יבינו כי בחשכה יתהלכו, ולא ידעו להזהר באיסורים דאוריתא ודרבנן ונקל להם לדבר על ריעיהם לה"ר ורכילות, ולמסור גופם וממונם לאנסים ולהעיד עליהם עדות שקר אף בעש"ג כי הם מינים ואפיקורסים על שפורשים מדרכי הקבלה החדשה אשר רבו בה הבעלים והעשתרות והם האלוהות הזכרים והנקבות, אשר כל עבודת עובדיהם לתקנם ולקשט הנקבות כדי למצוא חן בעיני האלוהות הזכרים שיזדווגו להן כדי להוריד שפע רוחניותם גם לעולם התחתון, כדרשת רז"ל על והנה עמד על היאור לעיל סי' ע"ה, שהאלהים האחרים אלמלא עובדיהם שומרים אותם ומשגיחים עליהם בכל צרכיהם ומחזקים בדקיהם וקלקולם, אזי כרעו ונפלו וחלפו ועברו ובטלו, וכזה היא דעת המקובלים החדשים שכל מעשה מצותיהם עושים לשם האלילים לתקנם כמבואר בעץ חיים שער נקודים סוף פ"ו שקודם אדה"ר היו זעיר ונוקביה אחור באחור וכו' וכשנולד אדה"ר, על ידי מעשיו הטובים החזירם פנים בפנים (ודבריו אלה היפך זהר בראשית כ"ב) אבל אבא ואימא לא הוצרכו למעשיו כי על ידי עצמן חזרו פנים בפנים, (ובזהר שם אמר שנתגרשה אימא על ידי מעשיו). ובס' חיי שלום בשם ס' הלקוטים כתב כי עד אברהם אבינו עדיין לא נתקן הזעיר אנפין וכו' ע"ש, הכלל העולה כי האלוה קצר האפים, שעובדים אותו המקובלים החדשים הוא צריך תיקון בכל עת וזמן עם שתי נשיו רחל ולאה כאשר יראה המעיין בסדר תיקון חצות בספר משנת חסידים ובס' פרי עץ חיים לקשטם לו כדי שיתחמם הזעיר ויתאוה להם (משנת חסידים תיקון חצות פ"ד משנה ג'). אבל אלהינו בהו"ש אינו צריך לבריותיו כלל, רק הם צריכים לו, והוא שומר את יראיו וחסידיו מכל רע, כמבואר בכמה מקומות בכתובים ובדברי רז"ל, ואין צורך להאריך בהם כי די בעיקר שכתב הרמב"ם וכמ"ש רז"ל והנה ה' נצב עליו לשמרו וכו'.

| id | word | punctuation after |
|---|---|---|
| 9:20.1 | קלח | ) |
| 9:20.2 | את |  |
| 9:20.3 | כל |  |
| 9:20.4 | אלה |  |
| 9:20.5 | עשתה |  |
| 9:20.6 | יד |  |
| 9:20.7 | הפילסוף |  |
| 9:20.8 | המסית |  |
| 9:20.9 | מחבר |  |
| 9:20.10 | הזהר |  |
| 9:20.11 | להלביש |  |
| 9:20.12 | את |  |
| 9:20.13 | קוראיו |  |
| 9:20.14 | גאה |  |
| 9:20.15 | וגאון |  |
| 9:20.16 | אף |  |
| 9:20.17 | שאינם |  |
| 9:20.18 | יודעים |  |
| 9:20.19 | מכל |  |
| 9:20.20 | דיני |  |
| 9:20.21 | תורתינו |  |
| 9:20.22 | הקדושה |  |
| 9:20.23 | מאומה | , |
| 9:20.24 | ואף |  |
| 9:20.25 | מכל |  |
| 9:20.26 | מה |  |
| 9:20.27 | שלומדים |  |
| 9:20.28 | בו |  |
| 9:20.29 | שכן |  |
| 9:20.30 | הוא |  |
| 9:20.31 | הבטיח |  |
| 9:20.32 | להם |  |
| 9:20.33 | שכר |  |
| 9:20.34 | רב |  |
| 9:20.35 | בלמדם |  |
| 9:20.36 | אותו |  |
| 9:20.37 | אף |  |
| 9:20.38 | בלי |  |
| 9:20.39 | הבין | , |
| 9:20.40 | ולכן |  |
| 9:20.41 | הם |  |
| 9:20.42 | חושבים |  |
| 9:20.43 | עצמם |  |
| 9:20.44 | בני |  |
| 9:20.45 | עלייה | , |
| 9:20.46 | ומבזים |  |
| 9:20.47 | המה |  |
| 9:20.48 | את |  |
| 9:20.49 | העוסקים |  |
| 9:20.50 | במשנה |  |
| 9:20.51 | ובתלמוד |  |
| 9:20.52 | שהם |  |
| 9:20.53 | כתרה |  |
| 9:20.54 | של |  |
| 9:20.55 | תורה |  |
| 9:20.56 | כעדות |  |
| 9:20.57 | רז"ל |  |
| 9:20.58 | במס | ' |
| 9:20.59 | מגלה |  |
| 9:20.60 | והבאתי |  |
| 9:20.61 | לעיל |  |
| 9:20.62 | סי | ' |
| 9:20.63 | קל"ו | , |
| 9:20.64 | ובעיניהם |  |
| 9:20.65 | המשנה |  |
| 9:20.66 | והתלמוד |  |
| 9:20.67 | חשובים |  |
| 9:20.68 | לתבן |  |
| 9:20.69 | וקש | , |
| 9:20.70 | והעוסקים |  |
| 9:20.71 | בהם |  |
| 9:20.72 | כבהמות |  |
| 9:20.73 | יער |  |
| 9:20.74 | האוכלות |  |
| 9:20.75 | חציר |  |
| 9:20.76 | ותבן | , |
| 9:20.77 | וכעופות |  |
| 9:20.78 | כדאיתא |  |
| 9:20.79 | בזהר |  |
| 9:20.80 | פרשת |  |
| 9:20.81 | תצא |  |
| 9:20.82 | הובא |  |
| 9:20.83 | לעיל |  |
| 9:20.84 | סי | ' |
| 9:20.85 | ס"ח | , |
| 9:20.86 | וכן |  |
| 9:20.87 | חשב |  |
| 9:20.88 | מהרח"ו |  |
| 9:20.89 | בשבחיו |  |
| 9:20.90 | דף |  |
| 9:20.91 | י"ב | , |
| 9:20.92 | ולא |  |
| 9:20.93 | ידעו |  |
| 9:20.94 | ולא |  |
| 9:20.95 | יבינו |  |
| 9:20.96 | כי |  |
| 9:20.97 | בחשכה |  |
| 9:20.98 | יתהלכו | , |
| 9:20.99 | ולא |  |
| 9:20.100 | ידעו |  |
| 9:20.101 | להזהר |  |
| 9:20.102 | באיסורים |  |
| 9:20.103 | דאוריתא |  |
| 9:20.104 | ודרבנן |  |
| 9:20.105 | ונקל |  |
| 9:20.106 | להם |  |
| 9:20.107 | לדבר |  |
| 9:20.108 | על |  |
| 9:20.109 | ריעיהם |  |
| 9:20.110 | לה"ר |  |
| 9:20.111 | ורכילות | , |
| 9:20.112 | ולמסור |  |
| 9:20.113 | גופם |  |
| 9:20.114 | וממונם |  |
| 9:20.115 | לאנסים |  |
| 9:20.116 | ולהעיד |  |
| 9:20.117 | עליהם |  |
| 9:20.118 | עדות |  |
| 9:20.119 | שקר |  |
| 9:20.120 | אף |  |
| 9:20.121 | בעש"ג |  |
| 9:20.122 | כי |  |
| 9:20.123 | הם |  |
| 9:20.124 | מינים |  |
| 9:20.125 | ואפיקורסים |  |
| 9:20.126 | על |  |
| 9:20.127 | שפורשים |  |
| 9:20.128 | מדרכי |  |
| 9:20.129 | הקבלה |  |
| 9:20.130 | החדשה |  |
| 9:20.131 | אשר |  |
| 9:20.132 | רבו |  |
| 9:20.133 | בה |  |
| 9:20.134 | הבעלים |  |
| 9:20.135 | והעשתרות |  |
| 9:20.136 | והם |  |
| 9:20.137 | האלוהות |  |
| 9:20.138 | הזכרים |  |
| 9:20.139 | והנקבות | , |
| 9:20.140 | אשר |  |
| 9:20.141 | כל |  |
| 9:20.142 | עבודת |  |
| 9:20.143 | עובדיהם |  |
| 9:20.144 | לתקנם |  |
| 9:20.145 | ולקשט |  |
| 9:20.146 | הנקבות |  |
| 9:20.147 | כדי |  |
| 9:20.148 | למצוא |  |
| 9:20.149 | חן |  |
| 9:20.150 | בעיני |  |
| 9:20.151 | האלוהות |  |
| 9:20.152 | הזכרים |  |
| 9:20.153 | שיזדווגו |  |
| 9:20.154 | להן |  |
| 9:20.155 | כדי |  |
| 9:20.156 | להוריד |  |
| 9:20.157 | שפע |  |
| 9:20.158 | רוחניותם |  |
| 9:20.159 | גם |  |
| 9:20.160 | לעולם |  |
| 9:20.161 | התחתון | , |
| 9:20.162 | כדרשת |  |
| 9:20.163 | רז"ל |  |
| 9:20.164 | על |  |
| 9:20.165 | והנה |  |
| 9:20.166 | עמד |  |
| 9:20.167 | על |  |
| 9:20.168 | היאור |  |
| 9:20.169 | לעיל |  |
| 9:20.170 | סי | ' |
| 9:20.171 | ע"ה | , |
| 9:20.172 | שהאלהים |  |
| 9:20.173 | האחרים |  |
| 9:20.174 | אלמלא |  |
| 9:20.175 | עובדיהם |  |
| 9:20.176 | שומרים |  |
| 9:20.177 | אותם |  |
| 9:20.178 | ומשגיחים |  |
| 9:20.179 | עליהם |  |
| 9:20.180 | בכל |  |
| 9:20.181 | צרכיהם |  |
| 9:20.182 | ומחזקים |  |
| 9:20.183 | בדקיהם |  |
| 9:20.184 | וקלקולם | , |
| 9:20.185 | אזי |  |
| 9:20.186 | כרעו |  |
| 9:20.187 | ונפלו |  |
| 9:20.188 | וחלפו |  |
| 9:20.189 | ועברו |  |
| 9:20.190 | ובטלו | , |
| 9:20.191 | וכזה |  |
| 9:20.192 | היא |  |
| 9:20.193 | דעת |  |
| 9:20.194 | המקובלים |  |
| 9:20.195 | החדשים |  |
| 9:20.196 | שכל |  |
| 9:20.197 | מעשה |  |
| 9:20.198 | מצותיהם |  |
| 9:20.199 | עושים |  |
| 9:20.200 | לשם |  |
| 9:20.201 | האלילים |  |
| 9:20.202 | לתקנם |  |
| 9:20.203 | כמבואר |  |
| 9:20.204 | בעץ |  |
| 9:20.205 | חיים |  |
| 9:20.206 | שער |  |
| 9:20.207 | נקודים |  |
| 9:20.208 | סוף |  |
| 9:20.209 | פ"ו |  |
| 9:20.210 | שקודם |  |
| 9:20.211 | אדה"ר |  |
| 9:20.212 | היו |  |
| 9:20.213 | זעיר |  |
| 9:20.214 | ונוקביה |  |
| 9:20.215 | אחור |  |
| 9:20.216 | באחור |  |
| 9:20.217 | וכו | ' |
| 9:20.218 | וכשנולד |  |
| 9:20.219 | אדה"ר | , |
| 9:20.220 | על |  |
| 9:20.221 | ידי |  |
| 9:20.222 | מעשיו |  |
| 9:20.223 | הטובים |  |
| 9:20.224 | החזירם |  |
| 9:20.225 | פנים |  |
| 9:20.226 | בפנים |  |
| 9:20.227 | ודבריו |  |
| 9:20.228 | אלה |  |
| 9:20.229 | היפך |  |
| 9:20.230 | זהר |  |
| 9:20.231 | בראשית |  |
| 9:20.232 | כ"ב | ) |
| 9:20.233 | אבל |  |
| 9:20.234 | אבא |  |
| 9:20.235 | ואימא |  |
| 9:20.236 | לא |  |
| 9:20.237 | הוצרכו |  |
| 9:20.238 | למעשיו |  |
| 9:20.239 | כי |  |
| 9:20.240 | על |  |
| 9:20.241 | ידי |  |
| 9:20.242 | עצמן |  |
| 9:20.243 | חזרו |  |
| 9:20.244 | פנים |  |
| 9:20.245 | בפנים | , |
| 9:20.246 | ובזהר |  |
| 9:20.247 | שם |  |
| 9:20.248 | אמר |  |
| 9:20.249 | שנתגרשה |  |
| 9:20.250 | אימא |  |
| 9:20.251 | על |  |
| 9:20.252 | ידי |  |
| 9:20.253 | מעשיו | ). |
| 9:20.254 | ובס | ' |
| 9:20.255 | חיי |  |
| 9:20.256 | שלום |  |
| 9:20.257 | בשם |  |
| 9:20.258 | ס | ' |
| 9:20.259 | הלקוטים |  |
| 9:20.260 | כתב |  |
| 9:20.261 | כי |  |
| 9:20.262 | עד |  |
| 9:20.263 | אברהם |  |
| 9:20.264 | אבינו |  |
| 9:20.265 | עדיין |  |
| 9:20.266 | לא |  |
| 9:20.267 | נתקן |  |
| 9:20.268 | הזעיר |  |
| 9:20.269 | אנפין |  |
| 9:20.270 | וכו | ' |
| 9:20.271 | ע"ש | , |
| 9:20.272 | הכלל |  |
| 9:20.273 | העולה |  |
| 9:20.274 | כי |  |
| 9:20.275 | האלוה |  |
| 9:20.276 | קצר |  |
| 9:20.277 | האפים | , |
| 9:20.278 | שעובדים |  |
| 9:20.279 | אותו |  |
| 9:20.280 | המקובלים |  |
| 9:20.281 | החדשים |  |
| 9:20.282 | הוא |  |
| 9:20.283 | צריך |  |
| 9:20.284 | תיקון |  |
| 9:20.285 | בכל |  |
| 9:20.286 | עת |  |
| 9:20.287 | וזמן |  |
| 9:20.288 | עם |  |
| 9:20.289 | שתי |  |
| 9:20.290 | נשיו |  |
| 9:20.291 | רחל |  |
| 9:20.292 | ולאה |  |
| 9:20.293 | כאשר |  |
| 9:20.294 | יראה |  |
| 9:20.295 | המעיין |  |
| 9:20.296 | בסדר |  |
| 9:20.297 | תיקון |  |
| 9:20.298 | חצות |  |
| 9:20.299 | בספר |  |
| 9:20.300 | משנת |  |
| 9:20.301 | חסידים |  |
| 9:20.302 | ובס | ' |
| 9:20.303 | פרי |  |
| 9:20.304 | עץ |  |
| 9:20.305 | חיים |  |
| 9:20.306 | לקשטם |  |
| 9:20.307 | לו |  |
| 9:20.308 | כדי |  |
| 9:20.309 | שיתחמם |  |
| 9:20.310 | הזעיר |  |
| 9:20.311 | ויתאוה |  |
| 9:20.312 | להם |  |
| 9:20.313 | משנת |  |
| 9:20.314 | חסידים |  |
| 9:20.315 | תיקון |  |
| 9:20.316 | חצות |  |
| 9:20.317 | פ"ד |  |
| 9:20.318 | משנה |  |
| 9:20.319 | ג | '). |
| 9:20.320 | אבל |  |
| 9:20.321 | אלהינו |  |
| 9:20.322 | בהו"ש |  |
| 9:20.323 | אינו |  |
| 9:20.324 | צריך |  |
| 9:20.325 | לבריותיו |  |
| 9:20.326 | כלל | , |
| 9:20.327 | רק |  |
| 9:20.328 | הם |  |
| 9:20.329 | צריכים |  |
| 9:20.330 | לו | , |
| 9:20.331 | והוא |  |
| 9:20.332 | שומר |  |
| 9:20.333 | את |  |
| 9:20.334 | יראיו |  |
| 9:20.335 | וחסידיו |  |
| 9:20.336 | מכל |  |
| 9:20.337 | רע | , |
| 9:20.338 | כמבואר |  |
| 9:20.339 | בכמה |  |
| 9:20.340 | מקומות |  |
| 9:20.341 | בכתובים |  |
| 9:20.342 | ובדברי |  |
| 9:20.343 | רז"ל | , |
| 9:20.344 | ואין |  |
| 9:20.345 | צורך |  |
| 9:20.346 | להאריך |  |
| 9:20.347 | בהם |  |
| 9:20.348 | כי |  |
| 9:20.349 | די |  |
| 9:20.350 | בעיקר |  |
| 9:20.351 | שכתב |  |
| 9:20.352 | הרמב"ם |  |
| 9:20.353 | וכמ"ש |  |
| 9:20.354 | רז"ל |  |
| 9:20.355 | והנה |  |
| 9:20.356 | ה | ' |
| 9:20.357 | נצב |  |
| 9:20.358 | עליו |  |
| 9:20.359 | לשמרו |  |
| 9:20.360 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

