# Translation pack: milhamot-hashem perek-4-14 (Milchamot ha-Shem 4:23–4:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-14`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **העיקרים** → "the author of the Ikkarim".
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- מכתבי **עתים** (non-technical 1931 Hebrew) → "newspapers".
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
- **תא חזי** → "come and see".
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


## Milchamot ha-Shem 4:23

Hebrew (Wikisource): לב) והנה ידוע ומפורסם כמה הקפידו רז"ל מעתיקי השמועה ממשה רבי' ע"ה על האומר מודים מודים שתי פעמים או שמע שמע שתי פעמים שמשתקים אותנו כדתנן בברכות פ' אין עומדין. ובמגלה פרק הקורא האומר מודים מודים משתקים אותו. ובברכות דף ל"ג (ברכות לג, א) גר' אמר רבי זירא כל האומר שמע שמע כאומר מודים מודים דמי.ופי' כל המפרשים דטעמא משום דמחזי כשתי רשויות וכמקבל עליו שני אלוהות. ובברכות דף י"ד (דף יד.) ההוא דנחית קמיה דרבא שמעיה דקאמר אמת אמת תרי זמני. אמר רבא כל אמת אמת תפסיה להאי ע"כ וכ' הכונה בעין יע' כשם שאין לומר אחד אחד שהוא רבוי ולא אחדות. כך אין לומר אמת אמת כי האמת אינו אלא אחד. והוא מאמר רבא רהטא של אמת תפסתו לזה. כלומר הוא סובר לחזק דבריו באמתות השי"ת בהיותו מוסיף ואינו אלא גורע ע"כ. ובסוכה פרק החליל אהא דתנן אנו ליה וליה עינינו גרסינן איני והאמר ר' זירא כל האומר שמע שמע כאלו אומר מודים מודים? (פי' רש"י כאומר מודים מודים דתנן משתקים אותו משום דמחזי כשתי רשויות) ומתרץ אלא הכי אמרי אנו ליה מודים. ועינינו ליה מיחלות ע"כ. אלמא דרבנן קפדי אבל מלתא דמשתמע מינה שיש שתי רשויות. וכן פסקו כל הפוסקים דהאומר שמע שמע או מודים מודים משתקים אותו משום דמחזי כשתי רשויות. ומזה תדון קל וחומר (כפי הכלל המסור בידינו אדם דן ק"ו מעצמו) לנדון הוזכר בזהר (בראשית ל"ד ע"א) וז"ל כתיב ה' מלך ה' מלך ה' ימלוך לעולם ועד. ה' מלך לעילא. ה' מלך באמצעיתא. ה' ימלוך לתתא ופירש המקדש מלך בשם הכוונות ה' מלך הוא אריך. ה' מלך הוא אבא ואימא. ה' ימלוך זעיר ונוקביה ע"כ. אלמא דבאמרו ה' מלך הוא ממליך לאריך אנפין. ובאמרו ה' מלך הוא ממליך לאבא ואימא. ובאמרו ה' ימלוך לעולם ועד ממליך לזעיר ונוקביה. היש אמונת רשויות רבות ומינות גדולה מזו? שכיון שהמליך שלשה מלכים באמרו ה' מלך וכו' איך ישקר לומר אבינו מלכנו אין לנו מלך אלא אתה. והלא כבר המליך שלשה מלכים? ובאיזה מהם הוא בוחר עתה להמליכו. ולהסיר עטרת מלוכה מהאחרים? ובזה וכיוצא בו צדקו דברי הר' בעל העיקרים מאמר שני סוף פ' כ"ח ששם הזהיר מאד שלא ללמוד בספר הזהר ושאר ספרי קבלה וז"ל וכלל אומר לך השמר לך ושמור נפשך מאד פן תנקש אחריהם ותלכד במצודתם. כי הם עוזבים ארחות יושר ללכת בדרכי חשך כי לא ידעו ולא יבינו בחשכה חתהלכו המתעסקים בקבלה מדעת עצמם מזולת קבלה מפי חכם מקובל עכ"ל. ועל כן כל אוהב את השי"ת ודבק בתורתו שבכתב ושבעל פה אשר העתיקו רז"ל בעלי המשנה והתלמוד. ירחיק עצמו מהקבלה החדשה, ואל ידאג למטעמיה פן יפסיד אמונתו הטהורה המזוקקה שבעתים על פי תנאים, ואמוראים ופוסקים.

| id | word | punctuation after |
|---|---|---|
| 4:23.1 | לב | ) |
| 4:23.2 | והנה |  |
| 4:23.3 | ידוע |  |
| 4:23.4 | ומפורסם |  |
| 4:23.5 | כמה |  |
| 4:23.6 | הקפידו |  |
| 4:23.7 | רז"ל |  |
| 4:23.8 | מעתיקי |  |
| 4:23.9 | השמועה |  |
| 4:23.10 | ממשה |  |
| 4:23.11 | רבי | ' |
| 4:23.12 | ע"ה |  |
| 4:23.13 | על |  |
| 4:23.14 | האומר |  |
| 4:23.15 | מודים |  |
| 4:23.16 | מודים |  |
| 4:23.17 | שתי |  |
| 4:23.18 | פעמים |  |
| 4:23.19 | או |  |
| 4:23.20 | שמע |  |
| 4:23.21 | שמע |  |
| 4:23.22 | שתי |  |
| 4:23.23 | פעמים |  |
| 4:23.24 | שמשתקים |  |
| 4:23.25 | אותנו |  |
| 4:23.26 | כדתנן |  |
| 4:23.27 | בברכות |  |
| 4:23.28 | פ | ' |
| 4:23.29 | אין |  |
| 4:23.30 | עומדין | . |
| 4:23.31 | ובמגלה |  |
| 4:23.32 | פרק |  |
| 4:23.33 | הקורא |  |
| 4:23.34 | האומר |  |
| 4:23.35 | מודים |  |
| 4:23.36 | מודים |  |
| 4:23.37 | משתקים |  |
| 4:23.38 | אותו | . |
| 4:23.39 | ובברכות |  |
| 4:23.40 | דף |  |
| 4:23.41 | ל"ג |  |
| 4:23.42 | ברכות |  |
| 4:23.43 | לג | , |
| 4:23.44 | א | ) |
| 4:23.45 | גר | ' |
| 4:23.46 | אמר |  |
| 4:23.47 | רבי |  |
| 4:23.48 | זירא |  |
| 4:23.49 | כל |  |
| 4:23.50 | האומר |  |
| 4:23.51 | שמע |  |
| 4:23.52 | שמע |  |
| 4:23.53 | כאומר |  |
| 4:23.54 | מודים |  |
| 4:23.55 | מודים |  |
| 4:23.56 | דמי.ופי | ' |
| 4:23.57 | כל |  |
| 4:23.58 | המפרשים |  |
| 4:23.59 | דטעמא |  |
| 4:23.60 | משום |  |
| 4:23.61 | דמחזי |  |
| 4:23.62 | כשתי |  |
| 4:23.63 | רשויות |  |
| 4:23.64 | וכמקבל |  |
| 4:23.65 | עליו |  |
| 4:23.66 | שני |  |
| 4:23.67 | אלוהות | . |
| 4:23.68 | ובברכות |  |
| 4:23.69 | דף |  |
| 4:23.70 | י"ד |  |
| 4:23.71 | דף |  |
| 4:23.72 | יד | .) |
| 4:23.73 | ההוא |  |
| 4:23.74 | דנחית |  |
| 4:23.75 | קמיה |  |
| 4:23.76 | דרבא |  |
| 4:23.77 | שמעיה |  |
| 4:23.78 | דקאמר |  |
| 4:23.79 | אמת |  |
| 4:23.80 | אמת |  |
| 4:23.81 | תרי |  |
| 4:23.82 | זמני | . |
| 4:23.83 | אמר |  |
| 4:23.84 | רבא |  |
| 4:23.85 | כל |  |
| 4:23.86 | אמת |  |
| 4:23.87 | אמת |  |
| 4:23.88 | תפסיה |  |
| 4:23.89 | להאי |  |
| 4:23.90 | ע"כ |  |
| 4:23.91 | וכ | ' |
| 4:23.92 | הכונה |  |
| 4:23.93 | בעין |  |
| 4:23.94 | יע | ' |
| 4:23.95 | כשם |  |
| 4:23.96 | שאין |  |
| 4:23.97 | לומר |  |
| 4:23.98 | אחד |  |
| 4:23.99 | אחד |  |
| 4:23.100 | שהוא |  |
| 4:23.101 | רבוי |  |
| 4:23.102 | ולא |  |
| 4:23.103 | אחדות | . |
| 4:23.104 | כך |  |
| 4:23.105 | אין |  |
| 4:23.106 | לומר |  |
| 4:23.107 | אמת |  |
| 4:23.108 | אמת |  |
| 4:23.109 | כי |  |
| 4:23.110 | האמת |  |
| 4:23.111 | אינו |  |
| 4:23.112 | אלא |  |
| 4:23.113 | אחד | . |
| 4:23.114 | והוא |  |
| 4:23.115 | מאמר |  |
| 4:23.116 | רבא |  |
| 4:23.117 | רהטא |  |
| 4:23.118 | של |  |
| 4:23.119 | אמת |  |
| 4:23.120 | תפסתו |  |
| 4:23.121 | לזה | . |
| 4:23.122 | כלומר |  |
| 4:23.123 | הוא |  |
| 4:23.124 | סובר |  |
| 4:23.125 | לחזק |  |
| 4:23.126 | דבריו |  |
| 4:23.127 | באמתות |  |
| 4:23.128 | השי"ת |  |
| 4:23.129 | בהיותו |  |
| 4:23.130 | מוסיף |  |
| 4:23.131 | ואינו |  |
| 4:23.132 | אלא |  |
| 4:23.133 | גורע |  |
| 4:23.134 | ע"כ | . |
| 4:23.135 | ובסוכה |  |
| 4:23.136 | פרק |  |
| 4:23.137 | החליל |  |
| 4:23.138 | אהא |  |
| 4:23.139 | דתנן |  |
| 4:23.140 | אנו |  |
| 4:23.141 | ליה |  |
| 4:23.142 | וליה |  |
| 4:23.143 | עינינו |  |
| 4:23.144 | גרסינן |  |
| 4:23.145 | איני |  |
| 4:23.146 | והאמר |  |
| 4:23.147 | ר | ' |
| 4:23.148 | זירא |  |
| 4:23.149 | כל |  |
| 4:23.150 | האומר |  |
| 4:23.151 | שמע |  |
| 4:23.152 | שמע |  |
| 4:23.153 | כאלו |  |
| 4:23.154 | אומר |  |
| 4:23.155 | מודים |  |
| 4:23.156 | מודים | ? |
| 4:23.157 | פי | ' |
| 4:23.158 | רש"י |  |
| 4:23.159 | כאומר |  |
| 4:23.160 | מודים |  |
| 4:23.161 | מודים |  |
| 4:23.162 | דתנן |  |
| 4:23.163 | משתקים |  |
| 4:23.164 | אותו |  |
| 4:23.165 | משום |  |
| 4:23.166 | דמחזי |  |
| 4:23.167 | כשתי |  |
| 4:23.168 | רשויות | ) |
| 4:23.169 | ומתרץ |  |
| 4:23.170 | אלא |  |
| 4:23.171 | הכי |  |
| 4:23.172 | אמרי |  |
| 4:23.173 | אנו |  |
| 4:23.174 | ליה |  |
| 4:23.175 | מודים | . |
| 4:23.176 | ועינינו |  |
| 4:23.177 | ליה |  |
| 4:23.178 | מיחלות |  |
| 4:23.179 | ע"כ | . |
| 4:23.180 | אלמא |  |
| 4:23.181 | דרבנן |  |
| 4:23.182 | קפדי |  |
| 4:23.183 | אבל |  |
| 4:23.184 | מלתא |  |
| 4:23.185 | דמשתמע |  |
| 4:23.186 | מינה |  |
| 4:23.187 | שיש |  |
| 4:23.188 | שתי |  |
| 4:23.189 | רשויות | . |
| 4:23.190 | וכן |  |
| 4:23.191 | פסקו |  |
| 4:23.192 | כל |  |
| 4:23.193 | הפוסקים |  |
| 4:23.194 | דהאומר |  |
| 4:23.195 | שמע |  |
| 4:23.196 | שמע |  |
| 4:23.197 | או |  |
| 4:23.198 | מודים |  |
| 4:23.199 | מודים |  |
| 4:23.200 | משתקים |  |
| 4:23.201 | אותו |  |
| 4:23.202 | משום |  |
| 4:23.203 | דמחזי |  |
| 4:23.204 | כשתי |  |
| 4:23.205 | רשויות | . |
| 4:23.206 | ומזה |  |
| 4:23.207 | תדון |  |
| 4:23.208 | קל |  |
| 4:23.209 | וחומר |  |
| 4:23.210 | כפי |  |
| 4:23.211 | הכלל |  |
| 4:23.212 | המסור |  |
| 4:23.213 | בידינו |  |
| 4:23.214 | אדם |  |
| 4:23.215 | דן |  |
| 4:23.216 | ק"ו |  |
| 4:23.217 | מעצמו | ) |
| 4:23.218 | לנדון |  |
| 4:23.219 | הוזכר |  |
| 4:23.220 | בזהר |  |
| 4:23.221 | בראשית |  |
| 4:23.222 | ל"ד |  |
| 4:23.223 | ע"א | ) |
| 4:23.224 | וז"ל |  |
| 4:23.225 | כתיב |  |
| 4:23.226 | ה | ' |
| 4:23.227 | מלך |  |
| 4:23.228 | ה | ' |
| 4:23.229 | מלך |  |
| 4:23.230 | ה | ' |
| 4:23.231 | ימלוך |  |
| 4:23.232 | לעולם |  |
| 4:23.233 | ועד | . |
| 4:23.234 | ה | ' |
| 4:23.235 | מלך |  |
| 4:23.236 | לעילא | . |
| 4:23.237 | ה | ' |
| 4:23.238 | מלך |  |
| 4:23.239 | באמצעיתא | . |
| 4:23.240 | ה | ' |
| 4:23.241 | ימלוך |  |
| 4:23.242 | לתתא |  |
| 4:23.243 | ופירש |  |
| 4:23.244 | המקדש |  |
| 4:23.245 | מלך |  |
| 4:23.246 | בשם |  |
| 4:23.247 | הכוונות |  |
| 4:23.248 | ה | ' |
| 4:23.249 | מלך |  |
| 4:23.250 | הוא |  |
| 4:23.251 | אריך | . |
| 4:23.252 | ה | ' |
| 4:23.253 | מלך |  |
| 4:23.254 | הוא |  |
| 4:23.255 | אבא |  |
| 4:23.256 | ואימא | . |
| 4:23.257 | ה | ' |
| 4:23.258 | ימלוך |  |
| 4:23.259 | זעיר |  |
| 4:23.260 | ונוקביה |  |
| 4:23.261 | ע"כ | . |
| 4:23.262 | אלמא |  |
| 4:23.263 | דבאמרו |  |
| 4:23.264 | ה | ' |
| 4:23.265 | מלך |  |
| 4:23.266 | הוא |  |
| 4:23.267 | ממליך |  |
| 4:23.268 | לאריך |  |
| 4:23.269 | אנפין | . |
| 4:23.270 | ובאמרו |  |
| 4:23.271 | ה | ' |
| 4:23.272 | מלך |  |
| 4:23.273 | הוא |  |
| 4:23.274 | ממליך |  |
| 4:23.275 | לאבא |  |
| 4:23.276 | ואימא | . |
| 4:23.277 | ובאמרו |  |
| 4:23.278 | ה | ' |
| 4:23.279 | ימלוך |  |
| 4:23.280 | לעולם |  |
| 4:23.281 | ועד |  |
| 4:23.282 | ממליך |  |
| 4:23.283 | לזעיר |  |
| 4:23.284 | ונוקביה | . |
| 4:23.285 | היש |  |
| 4:23.286 | אמונת |  |
| 4:23.287 | רשויות |  |
| 4:23.288 | רבות |  |
| 4:23.289 | ומינות |  |
| 4:23.290 | גדולה |  |
| 4:23.291 | מזו | ? |
| 4:23.292 | שכיון |  |
| 4:23.293 | שהמליך |  |
| 4:23.294 | שלשה |  |
| 4:23.295 | מלכים |  |
| 4:23.296 | באמרו |  |
| 4:23.297 | ה | ' |
| 4:23.298 | מלך |  |
| 4:23.299 | וכו | ' |
| 4:23.300 | איך |  |
| 4:23.301 | ישקר |  |
| 4:23.302 | לומר |  |
| 4:23.303 | אבינו |  |
| 4:23.304 | מלכנו |  |
| 4:23.305 | אין |  |
| 4:23.306 | לנו |  |
| 4:23.307 | מלך |  |
| 4:23.308 | אלא |  |
| 4:23.309 | אתה | . |
| 4:23.310 | והלא |  |
| 4:23.311 | כבר |  |
| 4:23.312 | המליך |  |
| 4:23.313 | שלשה |  |
| 4:23.314 | מלכים | ? |
| 4:23.315 | ובאיזה |  |
| 4:23.316 | מהם |  |
| 4:23.317 | הוא |  |
| 4:23.318 | בוחר |  |
| 4:23.319 | עתה |  |
| 4:23.320 | להמליכו | . |
| 4:23.321 | ולהסיר |  |
| 4:23.322 | עטרת |  |
| 4:23.323 | מלוכה |  |
| 4:23.324 | מהאחרים | ? |
| 4:23.325 | ובזה |  |
| 4:23.326 | וכיוצא |  |
| 4:23.327 | בו |  |
| 4:23.328 | צדקו |  |
| 4:23.329 | דברי |  |
| 4:23.330 | הר | ' |
| 4:23.331 | בעל |  |
| 4:23.332 | העיקרים |  |
| 4:23.333 | מאמר |  |
| 4:23.334 | שני |  |
| 4:23.335 | סוף |  |
| 4:23.336 | פ | ' |
| 4:23.337 | כ"ח |  |
| 4:23.338 | ששם |  |
| 4:23.339 | הזהיר |  |
| 4:23.340 | מאד |  |
| 4:23.341 | שלא |  |
| 4:23.342 | ללמוד |  |
| 4:23.343 | בספר |  |
| 4:23.344 | הזהר |  |
| 4:23.345 | ושאר |  |
| 4:23.346 | ספרי |  |
| 4:23.347 | קבלה |  |
| 4:23.348 | וז"ל |  |
| 4:23.349 | וכלל |  |
| 4:23.350 | אומר |  |
| 4:23.351 | לך |  |
| 4:23.352 | השמר |  |
| 4:23.353 | לך |  |
| 4:23.354 | ושמור |  |
| 4:23.355 | נפשך |  |
| 4:23.356 | מאד |  |
| 4:23.357 | פן |  |
| 4:23.358 | תנקש |  |
| 4:23.359 | אחריהם |  |
| 4:23.360 | ותלכד |  |
| 4:23.361 | במצודתם | . |
| 4:23.362 | כי |  |
| 4:23.363 | הם |  |
| 4:23.364 | עוזבים |  |
| 4:23.365 | ארחות |  |
| 4:23.366 | יושר |  |
| 4:23.367 | ללכת |  |
| 4:23.368 | בדרכי |  |
| 4:23.369 | חשך |  |
| 4:23.370 | כי |  |
| 4:23.371 | לא |  |
| 4:23.372 | ידעו |  |
| 4:23.373 | ולא |  |
| 4:23.374 | יבינו |  |
| 4:23.375 | בחשכה |  |
| 4:23.376 | חתהלכו |  |
| 4:23.377 | המתעסקים |  |
| 4:23.378 | בקבלה |  |
| 4:23.379 | מדעת |  |
| 4:23.380 | עצמם |  |
| 4:23.381 | מזולת |  |
| 4:23.382 | קבלה |  |
| 4:23.383 | מפי |  |
| 4:23.384 | חכם |  |
| 4:23.385 | מקובל |  |
| 4:23.386 | עכ"ל | . |
| 4:23.387 | ועל |  |
| 4:23.388 | כן |  |
| 4:23.389 | כל |  |
| 4:23.390 | אוהב |  |
| 4:23.391 | את |  |
| 4:23.392 | השי"ת |  |
| 4:23.393 | ודבק |  |
| 4:23.394 | בתורתו |  |
| 4:23.395 | שבכתב |  |
| 4:23.396 | ושבעל |  |
| 4:23.397 | פה |  |
| 4:23.398 | אשר |  |
| 4:23.399 | העתיקו |  |
| 4:23.400 | רז"ל |  |
| 4:23.401 | בעלי |  |
| 4:23.402 | המשנה |  |
| 4:23.403 | והתלמוד | . |
| 4:23.404 | ירחיק |  |
| 4:23.405 | עצמו |  |
| 4:23.406 | מהקבלה |  |
| 4:23.407 | החדשה | , |
| 4:23.408 | ואל |  |
| 4:23.409 | ידאג |  |
| 4:23.410 | למטעמיה |  |
| 4:23.411 | פן |  |
| 4:23.412 | יפסיד |  |
| 4:23.413 | אמונתו |  |
| 4:23.414 | הטהורה |  |
| 4:23.415 | המזוקקה |  |
| 4:23.416 | שבעתים |  |
| 4:23.417 | על |  |
| 4:23.418 | פי |  |
| 4:23.419 | תנאים | , |
| 4:23.420 | ואמוראים |  |
| 4:23.421 | ופוסקים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

