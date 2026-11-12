# Translation pack: milhamot-hashem perek-8-02 (Milchamot ha-Shem 8:3–8:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-02`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

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
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
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
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ת"ל** → "Scripture says".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 8:3

Hebrew (Wikisource): קא) ולתקן הכפירה והמינות הגלויה באמונה זו אמרו, כי עבודתינו שאנו עובדים לו אע"פ שהוא נברא כמו שכתב הרב יושר לבב כי לנשמתו אנו עובדים. ואע"פ שאנו מאמינים כי אלהינו אינו גוף ולא כח בגוף! והנה ידוע מדבריהם כי הנשמה שלכל נברא יש לה דמות וצורת הגוף בכל חלקי איבריו נשמת הזכר זכר ונשמת הנקבה נקבה, כמו שכתב עז לאלהים בבית קדש הקדשים פ"ג דף ט' ע"א שבהכרח גמור יש לנשמה פרצוף ברמ"ח איברים ושס"ה גידים בצורת אדם. ומתפשטת בצורת איבריה, ואם תמצא אבר מיותר אינה מתפשטת בו, וכן ג"כ נשמה לנשמה צריך שיהיה לה צורת הפרצוף וכו' עיין שם. כי על כן צריך לכפול לדידהו ה' אלהיכם אמת כדי להשלים רמ"ח איברי הנשמה שבזכר, ואם לא יכפול תשאר נשמתו חסרה (וזהו למרות דברי רז"ל שאמרו האומר מלה מלה בקריאת שמע וכופלה משתקין אותו משום דמחזי כשתי רשויות, ובכפילת ה' אלהיכם אמת כל שכן דמיחזי כשתי רשויות). ובמעמד הר סיני לפי דבריהם קדמה אשתו של זעיר אנפין שהיא המלכות לרדת על הר סיני. ואחר כך הלך הזעיר אנפין אחריה כמנוח אחרי אשתו. ונמצא הוא וזוגתו כולם ירדו להר סיני. שהרי הוא ירד להתחבר עמה כמו שכתב נחלת יוסף. וגם אבא ואימא גם הם נמצאו שם בהר סיני. שהרי אימא היא שאמרה אשר הוציאתיך. כי היא הוציאתם ממצרים, והמלכות היא חכמה במצרים בסיועת הבינה שהיא אימא, והתורה יצאה מחכמה שהוא אבא, ונמצא לפי זה שבמעמד הר סיני ירדו אבא ואימא וזעיר נוקביה על הר סיני. וזעיר אנפין הוא שאמר אנכי ה' אלהיך. ואם כן הוא יקשה מאד מאמר רז"ל בשמות רבה וז"ל אנכי ה' אלהיך. מלך בשר ודם שמלך יש לו אב או אח, או בן. אמר הקב"ה אני איני כן אני ראשון שאין לי אב. ואני אחרון שאין לי בן. ומבלעדי אין אלהים שאין לי אח ע"כ. הלא לפי הקבלה החדשה אבא ואימא ומלכות כולם נגלו עם קצר אפים (זעיר אנפין) בסיני והתיצבו שם עמו, והוא המדבר אנכי ה' אלהיך. איך אמר שאין לי אב ואם הרי הם נצבים עמו? ואיך אמר שאין לי אח והרי זוגתו קדמה לו לירד על הר סיני. והוא ירד אחריה להתחבר לה? השקר הוא מדבר. ולמה זה ישקר עלינו והם עומדים עמו? האם להתל בנו נגלה עלינו? חלילה לה' אלהינו מן השקר. כי הוא ה' אלהים אמת. ותורת אמת הנחילנו, ובמצות תורתינו הקדושה הרחיקנו מאמונות זרות, וכוזבות שטעו באן האומות עובדי האלילים זולתינו.

| id | word | punctuation after |
|---|---|---|
| 8:3.1 | קא | ) |
| 8:3.2 | ולתקן |  |
| 8:3.3 | הכפירה |  |
| 8:3.4 | והמינות |  |
| 8:3.5 | הגלויה |  |
| 8:3.6 | באמונה |  |
| 8:3.7 | זו |  |
| 8:3.8 | אמרו | , |
| 8:3.9 | כי |  |
| 8:3.10 | עבודתינו |  |
| 8:3.11 | שאנו |  |
| 8:3.12 | עובדים |  |
| 8:3.13 | לו |  |
| 8:3.14 | אע"פ |  |
| 8:3.15 | שהוא |  |
| 8:3.16 | נברא |  |
| 8:3.17 | כמו |  |
| 8:3.18 | שכתב |  |
| 8:3.19 | הרב |  |
| 8:3.20 | יושר |  |
| 8:3.21 | לבב |  |
| 8:3.22 | כי |  |
| 8:3.23 | לנשמתו |  |
| 8:3.24 | אנו |  |
| 8:3.25 | עובדים | . |
| 8:3.26 | ואע"פ |  |
| 8:3.27 | שאנו |  |
| 8:3.28 | מאמינים |  |
| 8:3.29 | כי |  |
| 8:3.30 | אלהינו |  |
| 8:3.31 | אינו |  |
| 8:3.32 | גוף |  |
| 8:3.33 | ולא |  |
| 8:3.34 | כח |  |
| 8:3.35 | בגוף | ! |
| 8:3.36 | והנה |  |
| 8:3.37 | ידוע |  |
| 8:3.38 | מדבריהם |  |
| 8:3.39 | כי |  |
| 8:3.40 | הנשמה |  |
| 8:3.41 | שלכל |  |
| 8:3.42 | נברא |  |
| 8:3.43 | יש |  |
| 8:3.44 | לה |  |
| 8:3.45 | דמות |  |
| 8:3.46 | וצורת |  |
| 8:3.47 | הגוף |  |
| 8:3.48 | בכל |  |
| 8:3.49 | חלקי |  |
| 8:3.50 | איבריו |  |
| 8:3.51 | נשמת |  |
| 8:3.52 | הזכר |  |
| 8:3.53 | זכר |  |
| 8:3.54 | ונשמת |  |
| 8:3.55 | הנקבה |  |
| 8:3.56 | נקבה | , |
| 8:3.57 | כמו |  |
| 8:3.58 | שכתב |  |
| 8:3.59 | עז |  |
| 8:3.60 | לאלהים |  |
| 8:3.61 | בבית |  |
| 8:3.62 | קדש |  |
| 8:3.63 | הקדשים |  |
| 8:3.64 | פ"ג |  |
| 8:3.65 | דף |  |
| 8:3.66 | ט | ' |
| 8:3.67 | ע"א |  |
| 8:3.68 | שבהכרח |  |
| 8:3.69 | גמור |  |
| 8:3.70 | יש |  |
| 8:3.71 | לנשמה |  |
| 8:3.72 | פרצוף |  |
| 8:3.73 | ברמ"ח |  |
| 8:3.74 | איברים |  |
| 8:3.75 | ושס"ה |  |
| 8:3.76 | גידים |  |
| 8:3.77 | בצורת |  |
| 8:3.78 | אדם | . |
| 8:3.79 | ומתפשטת |  |
| 8:3.80 | בצורת |  |
| 8:3.81 | איבריה | , |
| 8:3.82 | ואם |  |
| 8:3.83 | תמצא |  |
| 8:3.84 | אבר |  |
| 8:3.85 | מיותר |  |
| 8:3.86 | אינה |  |
| 8:3.87 | מתפשטת |  |
| 8:3.88 | בו | , |
| 8:3.89 | וכן |  |
| 8:3.90 | ג"כ |  |
| 8:3.91 | נשמה |  |
| 8:3.92 | לנשמה |  |
| 8:3.93 | צריך |  |
| 8:3.94 | שיהיה |  |
| 8:3.95 | לה |  |
| 8:3.96 | צורת |  |
| 8:3.97 | הפרצוף |  |
| 8:3.98 | וכו | ' |
| 8:3.99 | עיין |  |
| 8:3.100 | שם | . |
| 8:3.101 | כי |  |
| 8:3.102 | על |  |
| 8:3.103 | כן |  |
| 8:3.104 | צריך |  |
| 8:3.105 | לכפול |  |
| 8:3.106 | לדידהו |  |
| 8:3.107 | ה | ' |
| 8:3.108 | אלהיכם |  |
| 8:3.109 | אמת |  |
| 8:3.110 | כדי |  |
| 8:3.111 | להשלים |  |
| 8:3.112 | רמ"ח |  |
| 8:3.113 | איברי |  |
| 8:3.114 | הנשמה |  |
| 8:3.115 | שבזכר | , |
| 8:3.116 | ואם |  |
| 8:3.117 | לא |  |
| 8:3.118 | יכפול |  |
| 8:3.119 | תשאר |  |
| 8:3.120 | נשמתו |  |
| 8:3.121 | חסרה |  |
| 8:3.122 | וזהו |  |
| 8:3.123 | למרות |  |
| 8:3.124 | דברי |  |
| 8:3.125 | רז"ל |  |
| 8:3.126 | שאמרו |  |
| 8:3.127 | האומר |  |
| 8:3.128 | מלה |  |
| 8:3.129 | מלה |  |
| 8:3.130 | בקריאת |  |
| 8:3.131 | שמע |  |
| 8:3.132 | וכופלה |  |
| 8:3.133 | משתקין |  |
| 8:3.134 | אותו |  |
| 8:3.135 | משום |  |
| 8:3.136 | דמחזי |  |
| 8:3.137 | כשתי |  |
| 8:3.138 | רשויות | , |
| 8:3.139 | ובכפילת |  |
| 8:3.140 | ה | ' |
| 8:3.141 | אלהיכם |  |
| 8:3.142 | אמת |  |
| 8:3.143 | כל |  |
| 8:3.144 | שכן |  |
| 8:3.145 | דמיחזי |  |
| 8:3.146 | כשתי |  |
| 8:3.147 | רשויות | ). |
| 8:3.148 | ובמעמד |  |
| 8:3.149 | הר |  |
| 8:3.150 | סיני |  |
| 8:3.151 | לפי |  |
| 8:3.152 | דבריהם |  |
| 8:3.153 | קדמה |  |
| 8:3.154 | אשתו |  |
| 8:3.155 | של |  |
| 8:3.156 | זעיר |  |
| 8:3.157 | אנפין |  |
| 8:3.158 | שהיא |  |
| 8:3.159 | המלכות |  |
| 8:3.160 | לרדת |  |
| 8:3.161 | על |  |
| 8:3.162 | הר |  |
| 8:3.163 | סיני | . |
| 8:3.164 | ואחר |  |
| 8:3.165 | כך |  |
| 8:3.166 | הלך |  |
| 8:3.167 | הזעיר |  |
| 8:3.168 | אנפין |  |
| 8:3.169 | אחריה |  |
| 8:3.170 | כמנוח |  |
| 8:3.171 | אחרי |  |
| 8:3.172 | אשתו | . |
| 8:3.173 | ונמצא |  |
| 8:3.174 | הוא |  |
| 8:3.175 | וזוגתו |  |
| 8:3.176 | כולם |  |
| 8:3.177 | ירדו |  |
| 8:3.178 | להר |  |
| 8:3.179 | סיני | . |
| 8:3.180 | שהרי |  |
| 8:3.181 | הוא |  |
| 8:3.182 | ירד |  |
| 8:3.183 | להתחבר |  |
| 8:3.184 | עמה |  |
| 8:3.185 | כמו |  |
| 8:3.186 | שכתב |  |
| 8:3.187 | נחלת |  |
| 8:3.188 | יוסף | . |
| 8:3.189 | וגם |  |
| 8:3.190 | אבא |  |
| 8:3.191 | ואימא |  |
| 8:3.192 | גם |  |
| 8:3.193 | הם |  |
| 8:3.194 | נמצאו |  |
| 8:3.195 | שם |  |
| 8:3.196 | בהר |  |
| 8:3.197 | סיני | . |
| 8:3.198 | שהרי |  |
| 8:3.199 | אימא |  |
| 8:3.200 | היא |  |
| 8:3.201 | שאמרה |  |
| 8:3.202 | אשר |  |
| 8:3.203 | הוציאתיך | . |
| 8:3.204 | כי |  |
| 8:3.205 | היא |  |
| 8:3.206 | הוציאתם |  |
| 8:3.207 | ממצרים | , |
| 8:3.208 | והמלכות |  |
| 8:3.209 | היא |  |
| 8:3.210 | חכמה |  |
| 8:3.211 | במצרים |  |
| 8:3.212 | בסיועת |  |
| 8:3.213 | הבינה |  |
| 8:3.214 | שהיא |  |
| 8:3.215 | אימא | , |
| 8:3.216 | והתורה |  |
| 8:3.217 | יצאה |  |
| 8:3.218 | מחכמה |  |
| 8:3.219 | שהוא |  |
| 8:3.220 | אבא | , |
| 8:3.221 | ונמצא |  |
| 8:3.222 | לפי |  |
| 8:3.223 | זה |  |
| 8:3.224 | שבמעמד |  |
| 8:3.225 | הר |  |
| 8:3.226 | סיני |  |
| 8:3.227 | ירדו |  |
| 8:3.228 | אבא |  |
| 8:3.229 | ואימא |  |
| 8:3.230 | וזעיר |  |
| 8:3.231 | נוקביה |  |
| 8:3.232 | על |  |
| 8:3.233 | הר |  |
| 8:3.234 | סיני | . |
| 8:3.235 | וזעיר |  |
| 8:3.236 | אנפין |  |
| 8:3.237 | הוא |  |
| 8:3.238 | שאמר |  |
| 8:3.239 | אנכי |  |
| 8:3.240 | ה | ' |
| 8:3.241 | אלהיך | . |
| 8:3.242 | ואם |  |
| 8:3.243 | כן |  |
| 8:3.244 | הוא |  |
| 8:3.245 | יקשה |  |
| 8:3.246 | מאד |  |
| 8:3.247 | מאמר |  |
| 8:3.248 | רז"ל |  |
| 8:3.249 | בשמות |  |
| 8:3.250 | רבה |  |
| 8:3.251 | וז"ל |  |
| 8:3.252 | אנכי |  |
| 8:3.253 | ה | ' |
| 8:3.254 | אלהיך | . |
| 8:3.255 | מלך |  |
| 8:3.256 | בשר |  |
| 8:3.257 | ודם |  |
| 8:3.258 | שמלך |  |
| 8:3.259 | יש |  |
| 8:3.260 | לו |  |
| 8:3.261 | אב |  |
| 8:3.262 | או |  |
| 8:3.263 | אח | , |
| 8:3.264 | או |  |
| 8:3.265 | בן | . |
| 8:3.266 | אמר |  |
| 8:3.267 | הקב"ה |  |
| 8:3.268 | אני |  |
| 8:3.269 | איני |  |
| 8:3.270 | כן |  |
| 8:3.271 | אני |  |
| 8:3.272 | ראשון |  |
| 8:3.273 | שאין |  |
| 8:3.274 | לי |  |
| 8:3.275 | אב | . |
| 8:3.276 | ואני |  |
| 8:3.277 | אחרון |  |
| 8:3.278 | שאין |  |
| 8:3.279 | לי |  |
| 8:3.280 | בן | . |
| 8:3.281 | ומבלעדי |  |
| 8:3.282 | אין |  |
| 8:3.283 | אלהים |  |
| 8:3.284 | שאין |  |
| 8:3.285 | לי |  |
| 8:3.286 | אח |  |
| 8:3.287 | ע"כ | . |
| 8:3.288 | הלא |  |
| 8:3.289 | לפי |  |
| 8:3.290 | הקבלה |  |
| 8:3.291 | החדשה |  |
| 8:3.292 | אבא |  |
| 8:3.293 | ואימא |  |
| 8:3.294 | ומלכות |  |
| 8:3.295 | כולם |  |
| 8:3.296 | נגלו |  |
| 8:3.297 | עם |  |
| 8:3.298 | קצר |  |
| 8:3.299 | אפים |  |
| 8:3.300 | זעיר |  |
| 8:3.301 | אנפין | ) |
| 8:3.302 | בסיני |  |
| 8:3.303 | והתיצבו |  |
| 8:3.304 | שם |  |
| 8:3.305 | עמו | , |
| 8:3.306 | והוא |  |
| 8:3.307 | המדבר |  |
| 8:3.308 | אנכי |  |
| 8:3.309 | ה | ' |
| 8:3.310 | אלהיך | . |
| 8:3.311 | איך |  |
| 8:3.312 | אמר |  |
| 8:3.313 | שאין |  |
| 8:3.314 | לי |  |
| 8:3.315 | אב |  |
| 8:3.316 | ואם |  |
| 8:3.317 | הרי |  |
| 8:3.318 | הם |  |
| 8:3.319 | נצבים |  |
| 8:3.320 | עמו | ? |
| 8:3.321 | ואיך |  |
| 8:3.322 | אמר |  |
| 8:3.323 | שאין |  |
| 8:3.324 | לי |  |
| 8:3.325 | אח |  |
| 8:3.326 | והרי |  |
| 8:3.327 | זוגתו |  |
| 8:3.328 | קדמה |  |
| 8:3.329 | לו |  |
| 8:3.330 | לירד |  |
| 8:3.331 | על |  |
| 8:3.332 | הר |  |
| 8:3.333 | סיני | . |
| 8:3.334 | והוא |  |
| 8:3.335 | ירד |  |
| 8:3.336 | אחריה |  |
| 8:3.337 | להתחבר |  |
| 8:3.338 | לה | ? |
| 8:3.339 | השקר |  |
| 8:3.340 | הוא |  |
| 8:3.341 | מדבר | . |
| 8:3.342 | ולמה |  |
| 8:3.343 | זה |  |
| 8:3.344 | ישקר |  |
| 8:3.345 | עלינו |  |
| 8:3.346 | והם |  |
| 8:3.347 | עומדים |  |
| 8:3.348 | עמו | ? |
| 8:3.349 | האם |  |
| 8:3.350 | להתל |  |
| 8:3.351 | בנו |  |
| 8:3.352 | נגלה |  |
| 8:3.353 | עלינו | ? |
| 8:3.354 | חלילה |  |
| 8:3.355 | לה | ' |
| 8:3.356 | אלהינו |  |
| 8:3.357 | מן |  |
| 8:3.358 | השקר | . |
| 8:3.359 | כי |  |
| 8:3.360 | הוא |  |
| 8:3.361 | ה | ' |
| 8:3.362 | אלהים |  |
| 8:3.363 | אמת | . |
| 8:3.364 | ותורת |  |
| 8:3.365 | אמת |  |
| 8:3.366 | הנחילנו | , |
| 8:3.367 | ובמצות |  |
| 8:3.368 | תורתינו |  |
| 8:3.369 | הקדושה |  |
| 8:3.370 | הרחיקנו |  |
| 8:3.371 | מאמונות |  |
| 8:3.372 | זרות | , |
| 8:3.373 | וכוזבות |  |
| 8:3.374 | שטעו |  |
| 8:3.375 | באן |  |
| 8:3.376 | האומות |  |
| 8:3.377 | עובדי |  |
| 8:3.378 | האלילים |  |
| 8:3.379 | זולתינו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

