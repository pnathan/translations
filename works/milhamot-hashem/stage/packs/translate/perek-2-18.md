# Translation pack: milhamot-hashem perek-2-18 (Milchamot ha-Shem 2:34–2:34)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-18`.

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
- שם **שמים** → "the name of Heaven".
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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **תוה"ק** → "the holy Torah".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **דרך ארץ** → "worldly occupation".
- **רוע** לב → "wickedness of heart".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 2:34

Hebrew (Wikisource): תשובה אשאלה מאלהי ינחני במעגלי יושר ויצילנו משפת שקר, מלשון רמיה. הוריני ה' דרכיך אהלך באמתך. יחד לבבי ליראה שמך ושלח אורך ואמתך המה ינחוני להשיב דבר דבור על אפניו. גל מעלי חרפה ובוז כי עדותיך נצרתי. גם כי ישבו שרים בי נדברו, עבדך ישיח בחוקיך. אז לא אבוש כאשר לא אסור מן הדרך אשר למדונו רז"ל בעלי המשנה והתלמוד והמדרשים האמתיים והפוסקים ההולכים בעקבותיהם. אשר כל בית ישראל נשענים עליהם. כי המה בעלי הקבלה האמתית אשר קבלו ממרע"ה כללותיה ופרטותיה של תורה אשר עליהם עמד בסיני ארבעים יום וארבעים לילה ומסרן ליהושע ויהושע לזקנים וכו' כדאיתא במס' אבות. א) וחוב קדוש מוטל עלינו ללכת בעקבותיהם בכל מידי דנפקא מינה לדינא בכל ענייני איסור והתר וטמא וטהור ופסול וכשר. וכ"ש בענייני יחוד השי"ת שהוא העיקר הגדול ואחריו קבלת כל המצות וכיצד עשייתן כמאמר התנא ע"ה למה קדמה פרשת שמע לוהיה אם שמוע וכו'! וכל מחבר ספר אשר יבוא להוסיף או לגרוע מדבריהם וכ"ש להטותינו מאמונת תוה"ק כפי קבלת רבותינו בעלי המשנה והתלמוד. לא נשמע לו. ואף אם יעש אותות ומופתים בשמים ובארץ לא נסור מאחרי ה' אלהינו על פי קבלת רז"ל בעלי המשנה והתלמוד והפוסקים הבאים אחריהם ההולכים בעקבותיהם, וגם אם מן השמים תקרא בת קול שמעו אליו לבור מקבלתם לא משגחינן בה כמאמר ר' יהושע אין משגיחין בבת קול כי לא בשמים היא.

| id | word | punctuation after |
|---|---|---|
| 2:34.1 | תשובה |  |
| 2:34.2 | אשאלה |  |
| 2:34.3 | מאלהי |  |
| 2:34.4 | ינחני |  |
| 2:34.5 | במעגלי |  |
| 2:34.6 | יושר |  |
| 2:34.7 | ויצילנו |  |
| 2:34.8 | משפת |  |
| 2:34.9 | שקר | , |
| 2:34.10 | מלשון |  |
| 2:34.11 | רמיה | . |
| 2:34.12 | הוריני |  |
| 2:34.13 | ה | ' |
| 2:34.14 | דרכיך |  |
| 2:34.15 | אהלך |  |
| 2:34.16 | באמתך | . |
| 2:34.17 | יחד |  |
| 2:34.18 | לבבי |  |
| 2:34.19 | ליראה |  |
| 2:34.20 | שמך |  |
| 2:34.21 | ושלח |  |
| 2:34.22 | אורך |  |
| 2:34.23 | ואמתך |  |
| 2:34.24 | המה |  |
| 2:34.25 | ינחוני |  |
| 2:34.26 | להשיב |  |
| 2:34.27 | דבר |  |
| 2:34.28 | דבור |  |
| 2:34.29 | על |  |
| 2:34.30 | אפניו | . |
| 2:34.31 | גל |  |
| 2:34.32 | מעלי |  |
| 2:34.33 | חרפה |  |
| 2:34.34 | ובוז |  |
| 2:34.35 | כי |  |
| 2:34.36 | עדותיך |  |
| 2:34.37 | נצרתי | . |
| 2:34.38 | גם |  |
| 2:34.39 | כי |  |
| 2:34.40 | ישבו |  |
| 2:34.41 | שרים |  |
| 2:34.42 | בי |  |
| 2:34.43 | נדברו | , |
| 2:34.44 | עבדך |  |
| 2:34.45 | ישיח |  |
| 2:34.46 | בחוקיך | . |
| 2:34.47 | אז |  |
| 2:34.48 | לא |  |
| 2:34.49 | אבוש |  |
| 2:34.50 | כאשר |  |
| 2:34.51 | לא |  |
| 2:34.52 | אסור |  |
| 2:34.53 | מן |  |
| 2:34.54 | הדרך |  |
| 2:34.55 | אשר |  |
| 2:34.56 | למדונו |  |
| 2:34.57 | רז"ל |  |
| 2:34.58 | בעלי |  |
| 2:34.59 | המשנה |  |
| 2:34.60 | והתלמוד |  |
| 2:34.61 | והמדרשים |  |
| 2:34.62 | האמתיים |  |
| 2:34.63 | והפוסקים |  |
| 2:34.64 | ההולכים |  |
| 2:34.65 | בעקבותיהם | . |
| 2:34.66 | אשר |  |
| 2:34.67 | כל |  |
| 2:34.68 | בית |  |
| 2:34.69 | ישראל |  |
| 2:34.70 | נשענים |  |
| 2:34.71 | עליהם | . |
| 2:34.72 | כי |  |
| 2:34.73 | המה |  |
| 2:34.74 | בעלי |  |
| 2:34.75 | הקבלה |  |
| 2:34.76 | האמתית |  |
| 2:34.77 | אשר |  |
| 2:34.78 | קבלו |  |
| 2:34.79 | ממרע"ה |  |
| 2:34.80 | כללותיה |  |
| 2:34.81 | ופרטותיה |  |
| 2:34.82 | של |  |
| 2:34.83 | תורה |  |
| 2:34.84 | אשר |  |
| 2:34.85 | עליהם |  |
| 2:34.86 | עמד |  |
| 2:34.87 | בסיני |  |
| 2:34.88 | ארבעים |  |
| 2:34.89 | יום |  |
| 2:34.90 | וארבעים |  |
| 2:34.91 | לילה |  |
| 2:34.92 | ומסרן |  |
| 2:34.93 | ליהושע |  |
| 2:34.94 | ויהושע |  |
| 2:34.95 | לזקנים |  |
| 2:34.96 | וכו | ' |
| 2:34.97 | כדאיתא |  |
| 2:34.98 | במס | ' |
| 2:34.99 | אבות | . |
| 2:34.100 | א | ) |
| 2:34.101 | וחוב |  |
| 2:34.102 | קדוש |  |
| 2:34.103 | מוטל |  |
| 2:34.104 | עלינו |  |
| 2:34.105 | ללכת |  |
| 2:34.106 | בעקבותיהם |  |
| 2:34.107 | בכל |  |
| 2:34.108 | מידי |  |
| 2:34.109 | דנפקא |  |
| 2:34.110 | מינה |  |
| 2:34.111 | לדינא |  |
| 2:34.112 | בכל |  |
| 2:34.113 | ענייני |  |
| 2:34.114 | איסור |  |
| 2:34.115 | והתר |  |
| 2:34.116 | וטמא |  |
| 2:34.117 | וטהור |  |
| 2:34.118 | ופסול |  |
| 2:34.119 | וכשר | . |
| 2:34.120 | וכ"ש |  |
| 2:34.121 | בענייני |  |
| 2:34.122 | יחוד |  |
| 2:34.123 | השי"ת |  |
| 2:34.124 | שהוא |  |
| 2:34.125 | העיקר |  |
| 2:34.126 | הגדול |  |
| 2:34.127 | ואחריו |  |
| 2:34.128 | קבלת |  |
| 2:34.129 | כל |  |
| 2:34.130 | המצות |  |
| 2:34.131 | וכיצד |  |
| 2:34.132 | עשייתן |  |
| 2:34.133 | כמאמר |  |
| 2:34.134 | התנא |  |
| 2:34.135 | ע"ה |  |
| 2:34.136 | למה |  |
| 2:34.137 | קדמה |  |
| 2:34.138 | פרשת |  |
| 2:34.139 | שמע |  |
| 2:34.140 | לוהיה |  |
| 2:34.141 | אם |  |
| 2:34.142 | שמוע |  |
| 2:34.143 | וכו | '! |
| 2:34.144 | וכל |  |
| 2:34.145 | מחבר |  |
| 2:34.146 | ספר |  |
| 2:34.147 | אשר |  |
| 2:34.148 | יבוא |  |
| 2:34.149 | להוסיף |  |
| 2:34.150 | או |  |
| 2:34.151 | לגרוע |  |
| 2:34.152 | מדבריהם |  |
| 2:34.153 | וכ"ש |  |
| 2:34.154 | להטותינו |  |
| 2:34.155 | מאמונת |  |
| 2:34.156 | תוה"ק |  |
| 2:34.157 | כפי |  |
| 2:34.158 | קבלת |  |
| 2:34.159 | רבותינו |  |
| 2:34.160 | בעלי |  |
| 2:34.161 | המשנה |  |
| 2:34.162 | והתלמוד | . |
| 2:34.163 | לא |  |
| 2:34.164 | נשמע |  |
| 2:34.165 | לו | . |
| 2:34.166 | ואף |  |
| 2:34.167 | אם |  |
| 2:34.168 | יעש |  |
| 2:34.169 | אותות |  |
| 2:34.170 | ומופתים |  |
| 2:34.171 | בשמים |  |
| 2:34.172 | ובארץ |  |
| 2:34.173 | לא |  |
| 2:34.174 | נסור |  |
| 2:34.175 | מאחרי |  |
| 2:34.176 | ה | ' |
| 2:34.177 | אלהינו |  |
| 2:34.178 | על |  |
| 2:34.179 | פי |  |
| 2:34.180 | קבלת |  |
| 2:34.181 | רז"ל |  |
| 2:34.182 | בעלי |  |
| 2:34.183 | המשנה |  |
| 2:34.184 | והתלמוד |  |
| 2:34.185 | והפוסקים |  |
| 2:34.186 | הבאים |  |
| 2:34.187 | אחריהם |  |
| 2:34.188 | ההולכים |  |
| 2:34.189 | בעקבותיהם | , |
| 2:34.190 | וגם |  |
| 2:34.191 | אם |  |
| 2:34.192 | מן |  |
| 2:34.193 | השמים |  |
| 2:34.194 | תקרא |  |
| 2:34.195 | בת |  |
| 2:34.196 | קול |  |
| 2:34.197 | שמעו |  |
| 2:34.198 | אליו |  |
| 2:34.199 | לבור |  |
| 2:34.200 | מקבלתם |  |
| 2:34.201 | לא |  |
| 2:34.202 | משגחינן |  |
| 2:34.203 | בה |  |
| 2:34.204 | כמאמר |  |
| 2:34.205 | ר | ' |
| 2:34.206 | יהושע |  |
| 2:34.207 | אין |  |
| 2:34.208 | משגיחין |  |
| 2:34.209 | בבת |  |
| 2:34.210 | קול |  |
| 2:34.211 | כי |  |
| 2:34.212 | לא |  |
| 2:34.213 | בשמים |  |
| 2:34.214 | היא | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

