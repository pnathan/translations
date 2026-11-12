# Translation pack: milhamot-hashem perek-4-06 (Milchamot ha-Shem 4:10–4:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-06`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
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
- הסבה **הראשונה** → "the First Cause".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
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
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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
- ריש **מלין** אמר (Dan 7:1) → "he stated the sum of the matter".
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


## Milchamot ha-Shem 4:10

Hebrew (Wikisource): עניך הרואות איך רז"ל הרחיקו ליחס לשי"ת שהוא צריך לברואיו ונעזר בהם, או ליחס למלאכים העושים שליחותו שום בריאה, אלא הכל שב אל השי"ת, לא כמו שעשה הפילסוף מחבר הזהר שקרא לאבא בורא, ולאימא שם אותה לאומן העושה רצון בעל הבית. כמו שנתבאר לעיל, ולא יוחס הבריאה לסבה הראשונה רק לאבא, או או לחד פירוש לאריך בכח אבא כאלו הבורא חלש והם מסתייעים זה בזה בבריאה: ולפי דבריך שהמקובלים קיבלו כן ממשה רבי' ע"ה היה להם לע"ב זקנים להעתיק לתלמי המלך את הכתובים כמה שהם ולפרש לו כי בראשית שהוא אמר אבא ברא את אלהים שהיא בינה, וכן נעשה אדם היה להם לבאר שעל שנים נאמר, ולמה קלקלו את הכתובים בחנם, והחשיבו עצה במלין בלי דעת תורה שהיא הלכה למשה מסיני? ראוי היה להם להניח את הכתובים כמו שהם, ולפרש לו משמעות הכתובים בצורתן שהיא הלכה למשה לפי דבריך, והיו דבריהם נכונים ומתקבלים אצלו כפי דעתו באמונת אלוהות רבים? ולפרש לו ג"כ כי בראשית שם הוא כמו שפי' בזהר שהוא אבא והוא אלהים שאמר יהי אור ויהי רקיע, ויקוו המים, וגם יפרשו לו את השמים ואת הארץ. שמות אלוהות הן, כמו שפי' הפילוסוף בזהר שהם זעיר ונוקביה. ורז"ל הרחיקו הדעות האלה תכלית ההרחקה כדגרסינן בחגיגה פרק אין דורשין שאל ר' ישמעאל את ר' עקיבה כשהיו מהלכין בדרך. אמר לו אתה ששמשת את נחום איש גם זו כ"ב שנה שהיה דורש כל אתין שבתורה, את השמים ואת הארץ מאי דרש בהן אמר לו אלו נאמר שמים וארץ הייתי אומר שמותיה של הקב"ה הן. עכשיו שנ' את השמים ואת הארץ. שמים שמים ממש ארץ ארץ ממש. ובבראשית רבה הענין מתוקן יותר וז"ל ר' ישמעאל שאל את ר' עקיבה אמר לו בשביל ששמשת את נחום איש גם זו כ"ב שנה. אכין ורקין מעוטין. אתין וגמין ריבויין. הדין את דכתיב הכא מהו, אמר לו אלו נאמר בראשית ברא אלהים שמים וארץ היינו אומרים שמים וארץ אלוהות הן ע"כ. ועיין מהרש"א ועץ יוסף ותבין שזהו ממש דעת הזהר כי שמים וארץ הנז' כאן הם זעיר ונוקביה שהוא לפי דעתם ה' אלהינו וכמו שיתבאר לקמן.

| id | word | punctuation after |
|---|---|---|
| 4:10.1 | עניך |  |
| 4:10.2 | הרואות |  |
| 4:10.3 | איך |  |
| 4:10.4 | רז"ל |  |
| 4:10.5 | הרחיקו |  |
| 4:10.6 | ליחס |  |
| 4:10.7 | לשי"ת |  |
| 4:10.8 | שהוא |  |
| 4:10.9 | צריך |  |
| 4:10.10 | לברואיו |  |
| 4:10.11 | ונעזר |  |
| 4:10.12 | בהם | , |
| 4:10.13 | או |  |
| 4:10.14 | ליחס |  |
| 4:10.15 | למלאכים |  |
| 4:10.16 | העושים |  |
| 4:10.17 | שליחותו |  |
| 4:10.18 | שום |  |
| 4:10.19 | בריאה | , |
| 4:10.20 | אלא |  |
| 4:10.21 | הכל |  |
| 4:10.22 | שב |  |
| 4:10.23 | אל |  |
| 4:10.24 | השי"ת | , |
| 4:10.25 | לא |  |
| 4:10.26 | כמו |  |
| 4:10.27 | שעשה |  |
| 4:10.28 | הפילסוף |  |
| 4:10.29 | מחבר |  |
| 4:10.30 | הזהר |  |
| 4:10.31 | שקרא |  |
| 4:10.32 | לאבא |  |
| 4:10.33 | בורא | , |
| 4:10.34 | ולאימא |  |
| 4:10.35 | שם |  |
| 4:10.36 | אותה |  |
| 4:10.37 | לאומן |  |
| 4:10.38 | העושה |  |
| 4:10.39 | רצון |  |
| 4:10.40 | בעל |  |
| 4:10.41 | הבית | . |
| 4:10.42 | כמו |  |
| 4:10.43 | שנתבאר |  |
| 4:10.44 | לעיל | , |
| 4:10.45 | ולא |  |
| 4:10.46 | יוחס |  |
| 4:10.47 | הבריאה |  |
| 4:10.48 | לסבה |  |
| 4:10.49 | הראשונה |  |
| 4:10.50 | רק |  |
| 4:10.51 | לאבא | , |
| 4:10.52 | או |  |
| 4:10.53 | או |  |
| 4:10.54 | לחד |  |
| 4:10.55 | פירוש |  |
| 4:10.56 | לאריך |  |
| 4:10.57 | בכח |  |
| 4:10.58 | אבא |  |
| 4:10.59 | כאלו |  |
| 4:10.60 | הבורא |  |
| 4:10.61 | חלש |  |
| 4:10.62 | והם |  |
| 4:10.63 | מסתייעים |  |
| 4:10.64 | זה |  |
| 4:10.65 | בזה |  |
| 4:10.66 | בבריאה | : |
| 4:10.67 | ולפי |  |
| 4:10.68 | דבריך |  |
| 4:10.69 | שהמקובלים |  |
| 4:10.70 | קיבלו |  |
| 4:10.71 | כן |  |
| 4:10.72 | ממשה |  |
| 4:10.73 | רבי | ' |
| 4:10.74 | ע"ה |  |
| 4:10.75 | היה |  |
| 4:10.76 | להם |  |
| 4:10.77 | לע"ב |  |
| 4:10.78 | זקנים |  |
| 4:10.79 | להעתיק |  |
| 4:10.80 | לתלמי |  |
| 4:10.81 | המלך |  |
| 4:10.82 | את |  |
| 4:10.83 | הכתובים |  |
| 4:10.84 | כמה |  |
| 4:10.85 | שהם |  |
| 4:10.86 | ולפרש |  |
| 4:10.87 | לו |  |
| 4:10.88 | כי |  |
| 4:10.89 | בראשית |  |
| 4:10.90 | שהוא |  |
| 4:10.91 | אמר |  |
| 4:10.92 | אבא |  |
| 4:10.93 | ברא |  |
| 4:10.94 | את |  |
| 4:10.95 | אלהים |  |
| 4:10.96 | שהיא |  |
| 4:10.97 | בינה | , |
| 4:10.98 | וכן |  |
| 4:10.99 | נעשה |  |
| 4:10.100 | אדם |  |
| 4:10.101 | היה |  |
| 4:10.102 | להם |  |
| 4:10.103 | לבאר |  |
| 4:10.104 | שעל |  |
| 4:10.105 | שנים |  |
| 4:10.106 | נאמר | , |
| 4:10.107 | ולמה |  |
| 4:10.108 | קלקלו |  |
| 4:10.109 | את |  |
| 4:10.110 | הכתובים |  |
| 4:10.111 | בחנם | , |
| 4:10.112 | והחשיבו |  |
| 4:10.113 | עצה |  |
| 4:10.114 | במלין |  |
| 4:10.115 | בלי |  |
| 4:10.116 | דעת |  |
| 4:10.117 | תורה |  |
| 4:10.118 | שהיא |  |
| 4:10.119 | הלכה |  |
| 4:10.120 | למשה |  |
| 4:10.121 | מסיני | ? |
| 4:10.122 | ראוי |  |
| 4:10.123 | היה |  |
| 4:10.124 | להם |  |
| 4:10.125 | להניח |  |
| 4:10.126 | את |  |
| 4:10.127 | הכתובים |  |
| 4:10.128 | כמו |  |
| 4:10.129 | שהם | , |
| 4:10.130 | ולפרש |  |
| 4:10.131 | לו |  |
| 4:10.132 | משמעות |  |
| 4:10.133 | הכתובים |  |
| 4:10.134 | בצורתן |  |
| 4:10.135 | שהיא |  |
| 4:10.136 | הלכה |  |
| 4:10.137 | למשה |  |
| 4:10.138 | לפי |  |
| 4:10.139 | דבריך | , |
| 4:10.140 | והיו |  |
| 4:10.141 | דבריהם |  |
| 4:10.142 | נכונים |  |
| 4:10.143 | ומתקבלים |  |
| 4:10.144 | אצלו |  |
| 4:10.145 | כפי |  |
| 4:10.146 | דעתו |  |
| 4:10.147 | באמונת |  |
| 4:10.148 | אלוהות |  |
| 4:10.149 | רבים | ? |
| 4:10.150 | ולפרש |  |
| 4:10.151 | לו |  |
| 4:10.152 | ג"כ |  |
| 4:10.153 | כי |  |
| 4:10.154 | בראשית |  |
| 4:10.155 | שם |  |
| 4:10.156 | הוא |  |
| 4:10.157 | כמו |  |
| 4:10.158 | שפי | ' |
| 4:10.159 | בזהר |  |
| 4:10.160 | שהוא |  |
| 4:10.161 | אבא |  |
| 4:10.162 | והוא |  |
| 4:10.163 | אלהים |  |
| 4:10.164 | שאמר |  |
| 4:10.165 | יהי |  |
| 4:10.166 | אור |  |
| 4:10.167 | ויהי |  |
| 4:10.168 | רקיע | , |
| 4:10.169 | ויקוו |  |
| 4:10.170 | המים | , |
| 4:10.171 | וגם |  |
| 4:10.172 | יפרשו |  |
| 4:10.173 | לו |  |
| 4:10.174 | את |  |
| 4:10.175 | השמים |  |
| 4:10.176 | ואת |  |
| 4:10.177 | הארץ | . |
| 4:10.178 | שמות |  |
| 4:10.179 | אלוהות |  |
| 4:10.180 | הן | , |
| 4:10.181 | כמו |  |
| 4:10.182 | שפי | ' |
| 4:10.183 | הפילוסוף |  |
| 4:10.184 | בזהר |  |
| 4:10.185 | שהם |  |
| 4:10.186 | זעיר |  |
| 4:10.187 | ונוקביה | . |
| 4:10.188 | ורז"ל |  |
| 4:10.189 | הרחיקו |  |
| 4:10.190 | הדעות |  |
| 4:10.191 | האלה |  |
| 4:10.192 | תכלית |  |
| 4:10.193 | ההרחקה |  |
| 4:10.194 | כדגרסינן |  |
| 4:10.195 | בחגיגה |  |
| 4:10.196 | פרק |  |
| 4:10.197 | אין |  |
| 4:10.198 | דורשין |  |
| 4:10.199 | שאל |  |
| 4:10.200 | ר | ' |
| 4:10.201 | ישמעאל |  |
| 4:10.202 | את |  |
| 4:10.203 | ר | ' |
| 4:10.204 | עקיבה |  |
| 4:10.205 | כשהיו |  |
| 4:10.206 | מהלכין |  |
| 4:10.207 | בדרך | . |
| 4:10.208 | אמר |  |
| 4:10.209 | לו |  |
| 4:10.210 | אתה |  |
| 4:10.211 | ששמשת |  |
| 4:10.212 | את |  |
| 4:10.213 | נחום |  |
| 4:10.214 | איש |  |
| 4:10.215 | גם |  |
| 4:10.216 | זו |  |
| 4:10.217 | כ"ב |  |
| 4:10.218 | שנה |  |
| 4:10.219 | שהיה |  |
| 4:10.220 | דורש |  |
| 4:10.221 | כל |  |
| 4:10.222 | אתין |  |
| 4:10.223 | שבתורה | , |
| 4:10.224 | את |  |
| 4:10.225 | השמים |  |
| 4:10.226 | ואת |  |
| 4:10.227 | הארץ |  |
| 4:10.228 | מאי |  |
| 4:10.229 | דרש |  |
| 4:10.230 | בהן |  |
| 4:10.231 | אמר |  |
| 4:10.232 | לו |  |
| 4:10.233 | אלו |  |
| 4:10.234 | נאמר |  |
| 4:10.235 | שמים |  |
| 4:10.236 | וארץ |  |
| 4:10.237 | הייתי |  |
| 4:10.238 | אומר |  |
| 4:10.239 | שמותיה |  |
| 4:10.240 | של |  |
| 4:10.241 | הקב"ה |  |
| 4:10.242 | הן | . |
| 4:10.243 | עכשיו |  |
| 4:10.244 | שנ | ' |
| 4:10.245 | את |  |
| 4:10.246 | השמים |  |
| 4:10.247 | ואת |  |
| 4:10.248 | הארץ | . |
| 4:10.249 | שמים |  |
| 4:10.250 | שמים |  |
| 4:10.251 | ממש |  |
| 4:10.252 | ארץ |  |
| 4:10.253 | ארץ |  |
| 4:10.254 | ממש | . |
| 4:10.255 | ובבראשית |  |
| 4:10.256 | רבה |  |
| 4:10.257 | הענין |  |
| 4:10.258 | מתוקן |  |
| 4:10.259 | יותר |  |
| 4:10.260 | וז"ל |  |
| 4:10.261 | ר | ' |
| 4:10.262 | ישמעאל |  |
| 4:10.263 | שאל |  |
| 4:10.264 | את |  |
| 4:10.265 | ר | ' |
| 4:10.266 | עקיבה |  |
| 4:10.267 | אמר |  |
| 4:10.268 | לו |  |
| 4:10.269 | בשביל |  |
| 4:10.270 | ששמשת |  |
| 4:10.271 | את |  |
| 4:10.272 | נחום |  |
| 4:10.273 | איש |  |
| 4:10.274 | גם |  |
| 4:10.275 | זו |  |
| 4:10.276 | כ"ב |  |
| 4:10.277 | שנה | . |
| 4:10.278 | אכין |  |
| 4:10.279 | ורקין |  |
| 4:10.280 | מעוטין | . |
| 4:10.281 | אתין |  |
| 4:10.282 | וגמין |  |
| 4:10.283 | ריבויין | . |
| 4:10.284 | הדין |  |
| 4:10.285 | את |  |
| 4:10.286 | דכתיב |  |
| 4:10.287 | הכא |  |
| 4:10.288 | מהו | , |
| 4:10.289 | אמר |  |
| 4:10.290 | לו |  |
| 4:10.291 | אלו |  |
| 4:10.292 | נאמר |  |
| 4:10.293 | בראשית |  |
| 4:10.294 | ברא |  |
| 4:10.295 | אלהים |  |
| 4:10.296 | שמים |  |
| 4:10.297 | וארץ |  |
| 4:10.298 | היינו |  |
| 4:10.299 | אומרים |  |
| 4:10.300 | שמים |  |
| 4:10.301 | וארץ |  |
| 4:10.302 | אלוהות |  |
| 4:10.303 | הן |  |
| 4:10.304 | ע"כ | . |
| 4:10.305 | ועיין |  |
| 4:10.306 | מהרש"א |  |
| 4:10.307 | ועץ |  |
| 4:10.308 | יוסף |  |
| 4:10.309 | ותבין |  |
| 4:10.310 | שזהו |  |
| 4:10.311 | ממש |  |
| 4:10.312 | דעת |  |
| 4:10.313 | הזהר |  |
| 4:10.314 | כי |  |
| 4:10.315 | שמים |  |
| 4:10.316 | וארץ |  |
| 4:10.317 | הנז | ' |
| 4:10.318 | כאן |  |
| 4:10.319 | הם |  |
| 4:10.320 | זעיר |  |
| 4:10.321 | ונוקביה |  |
| 4:10.322 | שהוא |  |
| 4:10.323 | לפי |  |
| 4:10.324 | דעתם |  |
| 4:10.325 | ה | ' |
| 4:10.326 | אלהינו |  |
| 4:10.327 | וכמו |  |
| 4:10.328 | שיתבאר |  |
| 4:10.329 | לקמן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

