# Translation pack: milhamot-hashem perek-4-11 (Milchamot ha-Shem 4:19–4:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-11`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **מרן** → "Maran" (authority, transliterated; note).
- **רוקח** → "Roke'ach" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **א"כ** → "if so".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 4:19

Hebrew (Wikisource): ל) הנה מכל זה נתבאר לך באר היטב כי הזהר וכל מפרשיו מקדש מלך וכסא אליהו ור' עמנואל בעל משנת חסידים, ויושר לבב, והרש"ב, ומצרף האמונה וספר הכוונות, ומהרח"ו בעץ חיים, ומחברת הקדש, וספר הברית, והאר"י בספר הלקוטים ונחלת יוסף, וזולתם, כולם כאחד עונים ואומרים, דבאין סוף שהוא הסבה הראשונה לא שייך שום עבודה וקריאה. רק לזעיר אנפין שהוא סבה אחרונה שבכל הנאצלים שהוא לפי דעתם המקשר את כל הפרצופים, והוא הזן ומפרנס לכל הברואים ולו לבדו נעבוד ואליו נקרא בעת צרה, והוא יעננו ביום קראינו. אך הר' כסא אליהו מצריך להזכיר עמו בהתפללו אליו גם שאר הפרצופים שלמעלה ממנו, שאם אינו משתפם עמו אינו נענה מהרה אלא א"כ מזכיר שאר פרצופים עמו. וז"ל בדף כ"ו אלא מיבעי ליה ליחדם על הדרך שאמרנו להתפלל לזעיר אנפין. ובתפלתם להזכיר שמות הפרצופים הגבוהים ממנו כי צריכים להם, ואם לא יעשו כן, ולא יחדו תפלתם אלא לזעיר אנפין לא יהיו נענים מהרה וכדאמרן וזהו צירוף מ"ה תצע"ק אל"י, שהם אותיות מאצל העתיק וכו', ומסיים שם דגם אם לא היו מתפללים אלא לזעיר אנפין לא היו פוגמין אלא שלא היו נענים כל כך בקל ע"כ. וכן כתב הרב יושר לבב בית ב' חדר ג' פ' ז' ע"ש. הנה מה מאד ישתומם הקורא הנלבב הדבק בתורת משה רבינו ע"ה שבכתב ושבע"פ ובקי בה דהיינו במשנה ובתלמוד ומדרשי רז"ל, ועיניו יחזו דברי הגאונים, החסיד בעל חובת הלבבות, ורבי יהודה הלוי, ורבי סעדיה גאון, והרמב"ם בחיבוריו, ור' אליעזר מגרמיזא ברוקח וספר מצות גדול, וספר מצות קטן ובעל העקרים וכיוצא המדברים בענין יחוד השי"ת על אמתתו כפי קבלת רז"ל מעתיקי השמועה.

| id | word | punctuation after |
|---|---|---|
| 4:19.1 | ל | ) |
| 4:19.2 | הנה |  |
| 4:19.3 | מכל |  |
| 4:19.4 | זה |  |
| 4:19.5 | נתבאר |  |
| 4:19.6 | לך |  |
| 4:19.7 | באר |  |
| 4:19.8 | היטב |  |
| 4:19.9 | כי |  |
| 4:19.10 | הזהר |  |
| 4:19.11 | וכל |  |
| 4:19.12 | מפרשיו |  |
| 4:19.13 | מקדש |  |
| 4:19.14 | מלך |  |
| 4:19.15 | וכסא |  |
| 4:19.16 | אליהו |  |
| 4:19.17 | ור | ' |
| 4:19.18 | עמנואל |  |
| 4:19.19 | בעל |  |
| 4:19.20 | משנת |  |
| 4:19.21 | חסידים | , |
| 4:19.22 | ויושר |  |
| 4:19.23 | לבב | , |
| 4:19.24 | והרש"ב | , |
| 4:19.25 | ומצרף |  |
| 4:19.26 | האמונה |  |
| 4:19.27 | וספר |  |
| 4:19.28 | הכוונות | , |
| 4:19.29 | ומהרח"ו |  |
| 4:19.30 | בעץ |  |
| 4:19.31 | חיים | , |
| 4:19.32 | ומחברת |  |
| 4:19.33 | הקדש | , |
| 4:19.34 | וספר |  |
| 4:19.35 | הברית | , |
| 4:19.36 | והאר"י |  |
| 4:19.37 | בספר |  |
| 4:19.38 | הלקוטים |  |
| 4:19.39 | ונחלת |  |
| 4:19.40 | יוסף | , |
| 4:19.41 | וזולתם | , |
| 4:19.42 | כולם |  |
| 4:19.43 | כאחד |  |
| 4:19.44 | עונים |  |
| 4:19.45 | ואומרים | , |
| 4:19.46 | דבאין |  |
| 4:19.47 | סוף |  |
| 4:19.48 | שהוא |  |
| 4:19.49 | הסבה |  |
| 4:19.50 | הראשונה |  |
| 4:19.51 | לא |  |
| 4:19.52 | שייך |  |
| 4:19.53 | שום |  |
| 4:19.54 | עבודה |  |
| 4:19.55 | וקריאה | . |
| 4:19.56 | רק |  |
| 4:19.57 | לזעיר |  |
| 4:19.58 | אנפין |  |
| 4:19.59 | שהוא |  |
| 4:19.60 | סבה |  |
| 4:19.61 | אחרונה |  |
| 4:19.62 | שבכל |  |
| 4:19.63 | הנאצלים |  |
| 4:19.64 | שהוא |  |
| 4:19.65 | לפי |  |
| 4:19.66 | דעתם |  |
| 4:19.67 | המקשר |  |
| 4:19.68 | את |  |
| 4:19.69 | כל |  |
| 4:19.70 | הפרצופים | , |
| 4:19.71 | והוא |  |
| 4:19.72 | הזן |  |
| 4:19.73 | ומפרנס |  |
| 4:19.74 | לכל |  |
| 4:19.75 | הברואים |  |
| 4:19.76 | ולו |  |
| 4:19.77 | לבדו |  |
| 4:19.78 | נעבוד |  |
| 4:19.79 | ואליו |  |
| 4:19.80 | נקרא |  |
| 4:19.81 | בעת |  |
| 4:19.82 | צרה | , |
| 4:19.83 | והוא |  |
| 4:19.84 | יעננו |  |
| 4:19.85 | ביום |  |
| 4:19.86 | קראינו | . |
| 4:19.87 | אך |  |
| 4:19.88 | הר | ' |
| 4:19.89 | כסא |  |
| 4:19.90 | אליהו |  |
| 4:19.91 | מצריך |  |
| 4:19.92 | להזכיר |  |
| 4:19.93 | עמו |  |
| 4:19.94 | בהתפללו |  |
| 4:19.95 | אליו |  |
| 4:19.96 | גם |  |
| 4:19.97 | שאר |  |
| 4:19.98 | הפרצופים |  |
| 4:19.99 | שלמעלה |  |
| 4:19.100 | ממנו | , |
| 4:19.101 | שאם |  |
| 4:19.102 | אינו |  |
| 4:19.103 | משתפם |  |
| 4:19.104 | עמו |  |
| 4:19.105 | אינו |  |
| 4:19.106 | נענה |  |
| 4:19.107 | מהרה |  |
| 4:19.108 | אלא |  |
| 4:19.109 | א"כ |  |
| 4:19.110 | מזכיר |  |
| 4:19.111 | שאר |  |
| 4:19.112 | פרצופים |  |
| 4:19.113 | עמו | . |
| 4:19.114 | וז"ל |  |
| 4:19.115 | בדף |  |
| 4:19.116 | כ"ו |  |
| 4:19.117 | אלא |  |
| 4:19.118 | מיבעי |  |
| 4:19.119 | ליה |  |
| 4:19.120 | ליחדם |  |
| 4:19.121 | על |  |
| 4:19.122 | הדרך |  |
| 4:19.123 | שאמרנו |  |
| 4:19.124 | להתפלל |  |
| 4:19.125 | לזעיר |  |
| 4:19.126 | אנפין | . |
| 4:19.127 | ובתפלתם |  |
| 4:19.128 | להזכיר |  |
| 4:19.129 | שמות |  |
| 4:19.130 | הפרצופים |  |
| 4:19.131 | הגבוהים |  |
| 4:19.132 | ממנו |  |
| 4:19.133 | כי |  |
| 4:19.134 | צריכים |  |
| 4:19.135 | להם | , |
| 4:19.136 | ואם |  |
| 4:19.137 | לא |  |
| 4:19.138 | יעשו |  |
| 4:19.139 | כן | , |
| 4:19.140 | ולא |  |
| 4:19.141 | יחדו |  |
| 4:19.142 | תפלתם |  |
| 4:19.143 | אלא |  |
| 4:19.144 | לזעיר |  |
| 4:19.145 | אנפין |  |
| 4:19.146 | לא |  |
| 4:19.147 | יהיו |  |
| 4:19.148 | נענים |  |
| 4:19.149 | מהרה |  |
| 4:19.150 | וכדאמרן |  |
| 4:19.151 | וזהו |  |
| 4:19.152 | צירוף |  |
| 4:19.153 | מ"ה |  |
| 4:19.154 | תצע"ק |  |
| 4:19.155 | אל"י | , |
| 4:19.156 | שהם |  |
| 4:19.157 | אותיות |  |
| 4:19.158 | מאצל |  |
| 4:19.159 | העתיק |  |
| 4:19.160 | וכו | ', |
| 4:19.161 | ומסיים |  |
| 4:19.162 | שם |  |
| 4:19.163 | דגם |  |
| 4:19.164 | אם |  |
| 4:19.165 | לא |  |
| 4:19.166 | היו |  |
| 4:19.167 | מתפללים |  |
| 4:19.168 | אלא |  |
| 4:19.169 | לזעיר |  |
| 4:19.170 | אנפין |  |
| 4:19.171 | לא |  |
| 4:19.172 | היו |  |
| 4:19.173 | פוגמין |  |
| 4:19.174 | אלא |  |
| 4:19.175 | שלא |  |
| 4:19.176 | היו |  |
| 4:19.177 | נענים |  |
| 4:19.178 | כל |  |
| 4:19.179 | כך |  |
| 4:19.180 | בקל |  |
| 4:19.181 | ע"כ | . |
| 4:19.182 | וכן |  |
| 4:19.183 | כתב |  |
| 4:19.184 | הרב |  |
| 4:19.185 | יושר |  |
| 4:19.186 | לבב |  |
| 4:19.187 | בית |  |
| 4:19.188 | ב | ' |
| 4:19.189 | חדר |  |
| 4:19.190 | ג | ' |
| 4:19.191 | פ | ' |
| 4:19.192 | ז | ' |
| 4:19.193 | ע"ש | . |
| 4:19.194 | הנה |  |
| 4:19.195 | מה |  |
| 4:19.196 | מאד |  |
| 4:19.197 | ישתומם |  |
| 4:19.198 | הקורא |  |
| 4:19.199 | הנלבב |  |
| 4:19.200 | הדבק |  |
| 4:19.201 | בתורת |  |
| 4:19.202 | משה |  |
| 4:19.203 | רבינו |  |
| 4:19.204 | ע"ה |  |
| 4:19.205 | שבכתב |  |
| 4:19.206 | ושבע"פ |  |
| 4:19.207 | ובקי |  |
| 4:19.208 | בה |  |
| 4:19.209 | דהיינו |  |
| 4:19.210 | במשנה |  |
| 4:19.211 | ובתלמוד |  |
| 4:19.212 | ומדרשי |  |
| 4:19.213 | רז"ל | , |
| 4:19.214 | ועיניו |  |
| 4:19.215 | יחזו |  |
| 4:19.216 | דברי |  |
| 4:19.217 | הגאונים | , |
| 4:19.218 | החסיד |  |
| 4:19.219 | בעל |  |
| 4:19.220 | חובת |  |
| 4:19.221 | הלבבות | , |
| 4:19.222 | ורבי |  |
| 4:19.223 | יהודה |  |
| 4:19.224 | הלוי | , |
| 4:19.225 | ורבי |  |
| 4:19.226 | סעדיה |  |
| 4:19.227 | גאון | , |
| 4:19.228 | והרמב"ם |  |
| 4:19.229 | בחיבוריו | , |
| 4:19.230 | ור | ' |
| 4:19.231 | אליעזר |  |
| 4:19.232 | מגרמיזא |  |
| 4:19.233 | ברוקח |  |
| 4:19.234 | וספר |  |
| 4:19.235 | מצות |  |
| 4:19.236 | גדול | , |
| 4:19.237 | וספר |  |
| 4:19.238 | מצות |  |
| 4:19.239 | קטן |  |
| 4:19.240 | ובעל |  |
| 4:19.241 | העקרים |  |
| 4:19.242 | וכיוצא |  |
| 4:19.243 | המדברים |  |
| 4:19.244 | בענין |  |
| 4:19.245 | יחוד |  |
| 4:19.246 | השי"ת |  |
| 4:19.247 | על |  |
| 4:19.248 | אמתתו |  |
| 4:19.249 | כפי |  |
| 4:19.250 | קבלת |  |
| 4:19.251 | רז"ל |  |
| 4:19.252 | מעתיקי |  |
| 4:19.253 | השמועה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

