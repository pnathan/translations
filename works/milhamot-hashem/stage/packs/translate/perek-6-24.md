# Translation pack: milhamot-hashem perek-6-24 (Milchamot ha-Shem 6:37–6:37)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-24.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-24`.

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
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **סכלות** והוללות → "folly and madness".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 6:37

Hebrew (Wikisource): וכדי שלא יטעה הפתי הסכל מסכלותו ופתיותיו ומיעוט הבנתו ויחשוב שהשי"ת הוא על דרך הנראה ממליצת הספר שיש לו גשמות, הזהירנו בתורה שנשמר מאד מלחשוב שיש לבורא ית' צורה ולא דמיון. והוא אמרו ונשמרתם מאד לנפשותיכם כי לא ראיתם כל תמונה. והנביאים הזהירונו להמנע מלדמותו בשכל, ומלהמשילו לרעיון כאמרו ואל מי תדמיון אל, ומה דמות תערכו לו וכתוב ואל מי תדמיוני ואשוה. וכל זה צריך בעל אמונתינו לידע, ואז יוכל להבדיל בין הבורא וזולתו, ויהיה מיחד את השי"ת יחוד גמור כי מי שאינו יודע דרכי ההבדל אינו יכול ליחדו ית' יחוד אמתי. אמנם כאשר ידענו על מתכונתו אז ותקרב לאמתת בוראו ותתקרב נפשו להשגת המציאה. ותתעלה בהשגתה על מה כונתה כשאין אדם זוכה לידע אמתת בוראו אפי' יהיה מיחד אותו בלבו הרי הוא כאלו יש מחיצה בינו ובין בוראו ית'. כי זה הוא דבר אשר יגיע לו הפתי או הקטן. אשר איננו יודע דרכי האמונה ועניינה. ואין אמתתה קבועה בלבו. אבל כל עוד שמתחיל לחקור לידע אמתתו ית'. אזי המחיצה הולכת ומתמעטת. והנפש הולכת ומתקרבת אל בוראה יתעלה. ותתעלה נפשו ושכלו בתענו ג השגתו ית'. כי בהיות הנפש מתעסקת בדברים שכליים אז היא נמשכת אחריהם. לפי שהיא גם כן ממקור השכל נצמאת ולכן היא משתוקקת ונכספה להדמות לבוראה ית' ואינה יכולה להבדל ממנו עד שנמצאת דבוקה בו. ועל כיוצא בענין זה אמר ית' ובו תדבק. וגם על זה אמר הכתוב קרוב ה' לכל קוראיו. כי מאחר שהבורא ית' אינו גוף ולא כח בגוף אי אפשר היותו קרוב או רחוק מן הנפש. אלא רחוק מצד המחיצה המבדלת שהיא הסכלות כשאין לנפש ידיעה באמתתו ית'. וקרוב לנפש מצד הידיעה וההשגה. ולכן כל בעל נפש צריך לתור ולבקש על דרך ידיעת אמיתת מציאות הבורא ית'. כי הוא עיקר ויסוד גדול לאמונה. ועל דבר זה צוה אותנו בתורה ואמר וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:37.1 | וכדי |  |
| 6:37.2 | שלא |  |
| 6:37.3 | יטעה |  |
| 6:37.4 | הפתי |  |
| 6:37.5 | הסכל |  |
| 6:37.6 | מסכלותו |  |
| 6:37.7 | ופתיותיו |  |
| 6:37.8 | ומיעוט |  |
| 6:37.9 | הבנתו |  |
| 6:37.10 | ויחשוב |  |
| 6:37.11 | שהשי"ת |  |
| 6:37.12 | הוא |  |
| 6:37.13 | על |  |
| 6:37.14 | דרך |  |
| 6:37.15 | הנראה |  |
| 6:37.16 | ממליצת |  |
| 6:37.17 | הספר |  |
| 6:37.18 | שיש |  |
| 6:37.19 | לו |  |
| 6:37.20 | גשמות | , |
| 6:37.21 | הזהירנו |  |
| 6:37.22 | בתורה |  |
| 6:37.23 | שנשמר |  |
| 6:37.24 | מאד |  |
| 6:37.25 | מלחשוב |  |
| 6:37.26 | שיש |  |
| 6:37.27 | לבורא |  |
| 6:37.28 | ית | ' |
| 6:37.29 | צורה |  |
| 6:37.30 | ולא |  |
| 6:37.31 | דמיון | . |
| 6:37.32 | והוא |  |
| 6:37.33 | אמרו |  |
| 6:37.34 | ונשמרתם |  |
| 6:37.35 | מאד |  |
| 6:37.36 | לנפשותיכם |  |
| 6:37.37 | כי |  |
| 6:37.38 | לא |  |
| 6:37.39 | ראיתם |  |
| 6:37.40 | כל |  |
| 6:37.41 | תמונה | . |
| 6:37.42 | והנביאים |  |
| 6:37.43 | הזהירונו |  |
| 6:37.44 | להמנע |  |
| 6:37.45 | מלדמותו |  |
| 6:37.46 | בשכל | , |
| 6:37.47 | ומלהמשילו |  |
| 6:37.48 | לרעיון |  |
| 6:37.49 | כאמרו |  |
| 6:37.50 | ואל |  |
| 6:37.51 | מי |  |
| 6:37.52 | תדמיון |  |
| 6:37.53 | אל | , |
| 6:37.54 | ומה |  |
| 6:37.55 | דמות |  |
| 6:37.56 | תערכו |  |
| 6:37.57 | לו |  |
| 6:37.58 | וכתוב |  |
| 6:37.59 | ואל |  |
| 6:37.60 | מי |  |
| 6:37.61 | תדמיוני |  |
| 6:37.62 | ואשוה | . |
| 6:37.63 | וכל |  |
| 6:37.64 | זה |  |
| 6:37.65 | צריך |  |
| 6:37.66 | בעל |  |
| 6:37.67 | אמונתינו |  |
| 6:37.68 | לידע | , |
| 6:37.69 | ואז |  |
| 6:37.70 | יוכל |  |
| 6:37.71 | להבדיל |  |
| 6:37.72 | בין |  |
| 6:37.73 | הבורא |  |
| 6:37.74 | וזולתו | , |
| 6:37.75 | ויהיה |  |
| 6:37.76 | מיחד |  |
| 6:37.77 | את |  |
| 6:37.78 | השי"ת |  |
| 6:37.79 | יחוד |  |
| 6:37.80 | גמור |  |
| 6:37.81 | כי |  |
| 6:37.82 | מי |  |
| 6:37.83 | שאינו |  |
| 6:37.84 | יודע |  |
| 6:37.85 | דרכי |  |
| 6:37.86 | ההבדל |  |
| 6:37.87 | אינו |  |
| 6:37.88 | יכול |  |
| 6:37.89 | ליחדו |  |
| 6:37.90 | ית | ' |
| 6:37.91 | יחוד |  |
| 6:37.92 | אמתי | . |
| 6:37.93 | אמנם |  |
| 6:37.94 | כאשר |  |
| 6:37.95 | ידענו |  |
| 6:37.96 | על |  |
| 6:37.97 | מתכונתו |  |
| 6:37.98 | אז |  |
| 6:37.99 | ותקרב |  |
| 6:37.100 | לאמתת |  |
| 6:37.101 | בוראו |  |
| 6:37.102 | ותתקרב |  |
| 6:37.103 | נפשו |  |
| 6:37.104 | להשגת |  |
| 6:37.105 | המציאה | . |
| 6:37.106 | ותתעלה |  |
| 6:37.107 | בהשגתה |  |
| 6:37.108 | על |  |
| 6:37.109 | מה |  |
| 6:37.110 | כונתה |  |
| 6:37.111 | כשאין |  |
| 6:37.112 | אדם |  |
| 6:37.113 | זוכה |  |
| 6:37.114 | לידע |  |
| 6:37.115 | אמתת |  |
| 6:37.116 | בוראו |  |
| 6:37.117 | אפי | ' |
| 6:37.118 | יהיה |  |
| 6:37.119 | מיחד |  |
| 6:37.120 | אותו |  |
| 6:37.121 | בלבו |  |
| 6:37.122 | הרי |  |
| 6:37.123 | הוא |  |
| 6:37.124 | כאלו |  |
| 6:37.125 | יש |  |
| 6:37.126 | מחיצה |  |
| 6:37.127 | בינו |  |
| 6:37.128 | ובין |  |
| 6:37.129 | בוראו |  |
| 6:37.130 | ית | '. |
| 6:37.131 | כי |  |
| 6:37.132 | זה |  |
| 6:37.133 | הוא |  |
| 6:37.134 | דבר |  |
| 6:37.135 | אשר |  |
| 6:37.136 | יגיע |  |
| 6:37.137 | לו |  |
| 6:37.138 | הפתי |  |
| 6:37.139 | או |  |
| 6:37.140 | הקטן | . |
| 6:37.141 | אשר |  |
| 6:37.142 | איננו |  |
| 6:37.143 | יודע |  |
| 6:37.144 | דרכי |  |
| 6:37.145 | האמונה |  |
| 6:37.146 | ועניינה | . |
| 6:37.147 | ואין |  |
| 6:37.148 | אמתתה |  |
| 6:37.149 | קבועה |  |
| 6:37.150 | בלבו | . |
| 6:37.151 | אבל |  |
| 6:37.152 | כל |  |
| 6:37.153 | עוד |  |
| 6:37.154 | שמתחיל |  |
| 6:37.155 | לחקור |  |
| 6:37.156 | לידע |  |
| 6:37.157 | אמתתו |  |
| 6:37.158 | ית | '. |
| 6:37.159 | אזי |  |
| 6:37.160 | המחיצה |  |
| 6:37.161 | הולכת |  |
| 6:37.162 | ומתמעטת | . |
| 6:37.163 | והנפש |  |
| 6:37.164 | הולכת |  |
| 6:37.165 | ומתקרבת |  |
| 6:37.166 | אל |  |
| 6:37.167 | בוראה |  |
| 6:37.168 | יתעלה | . |
| 6:37.169 | ותתעלה |  |
| 6:37.170 | נפשו |  |
| 6:37.171 | ושכלו |  |
| 6:37.172 | בתענו |  |
| 6:37.173 | ג |  |
| 6:37.174 | השגתו |  |
| 6:37.175 | ית | '. |
| 6:37.176 | כי |  |
| 6:37.177 | בהיות |  |
| 6:37.178 | הנפש |  |
| 6:37.179 | מתעסקת |  |
| 6:37.180 | בדברים |  |
| 6:37.181 | שכליים |  |
| 6:37.182 | אז |  |
| 6:37.183 | היא |  |
| 6:37.184 | נמשכת |  |
| 6:37.185 | אחריהם | . |
| 6:37.186 | לפי |  |
| 6:37.187 | שהיא |  |
| 6:37.188 | גם |  |
| 6:37.189 | כן |  |
| 6:37.190 | ממקור |  |
| 6:37.191 | השכל |  |
| 6:37.192 | נצמאת |  |
| 6:37.193 | ולכן |  |
| 6:37.194 | היא |  |
| 6:37.195 | משתוקקת |  |
| 6:37.196 | ונכספה |  |
| 6:37.197 | להדמות |  |
| 6:37.198 | לבוראה |  |
| 6:37.199 | ית | ' |
| 6:37.200 | ואינה |  |
| 6:37.201 | יכולה |  |
| 6:37.202 | להבדל |  |
| 6:37.203 | ממנו |  |
| 6:37.204 | עד |  |
| 6:37.205 | שנמצאת |  |
| 6:37.206 | דבוקה |  |
| 6:37.207 | בו | . |
| 6:37.208 | ועל |  |
| 6:37.209 | כיוצא |  |
| 6:37.210 | בענין |  |
| 6:37.211 | זה |  |
| 6:37.212 | אמר |  |
| 6:37.213 | ית | ' |
| 6:37.214 | ובו |  |
| 6:37.215 | תדבק | . |
| 6:37.216 | וגם |  |
| 6:37.217 | על |  |
| 6:37.218 | זה |  |
| 6:37.219 | אמר |  |
| 6:37.220 | הכתוב |  |
| 6:37.221 | קרוב |  |
| 6:37.222 | ה | ' |
| 6:37.223 | לכל |  |
| 6:37.224 | קוראיו | . |
| 6:37.225 | כי |  |
| 6:37.226 | מאחר |  |
| 6:37.227 | שהבורא |  |
| 6:37.228 | ית | ' |
| 6:37.229 | אינו |  |
| 6:37.230 | גוף |  |
| 6:37.231 | ולא |  |
| 6:37.232 | כח |  |
| 6:37.233 | בגוף |  |
| 6:37.234 | אי |  |
| 6:37.235 | אפשר |  |
| 6:37.236 | היותו |  |
| 6:37.237 | קרוב |  |
| 6:37.238 | או |  |
| 6:37.239 | רחוק |  |
| 6:37.240 | מן |  |
| 6:37.241 | הנפש | . |
| 6:37.242 | אלא |  |
| 6:37.243 | רחוק |  |
| 6:37.244 | מצד |  |
| 6:37.245 | המחיצה |  |
| 6:37.246 | המבדלת |  |
| 6:37.247 | שהיא |  |
| 6:37.248 | הסכלות |  |
| 6:37.249 | כשאין |  |
| 6:37.250 | לנפש |  |
| 6:37.251 | ידיעה |  |
| 6:37.252 | באמתתו |  |
| 6:37.253 | ית | '. |
| 6:37.254 | וקרוב |  |
| 6:37.255 | לנפש |  |
| 6:37.256 | מצד |  |
| 6:37.257 | הידיעה |  |
| 6:37.258 | וההשגה | . |
| 6:37.259 | ולכן |  |
| 6:37.260 | כל |  |
| 6:37.261 | בעל |  |
| 6:37.262 | נפש |  |
| 6:37.263 | צריך |  |
| 6:37.264 | לתור |  |
| 6:37.265 | ולבקש |  |
| 6:37.266 | על |  |
| 6:37.267 | דרך |  |
| 6:37.268 | ידיעת |  |
| 6:37.269 | אמיתת |  |
| 6:37.270 | מציאות |  |
| 6:37.271 | הבורא |  |
| 6:37.272 | ית | '. |
| 6:37.273 | כי |  |
| 6:37.274 | הוא |  |
| 6:37.275 | עיקר |  |
| 6:37.276 | ויסוד |  |
| 6:37.277 | גדול |  |
| 6:37.278 | לאמונה | . |
| 6:37.279 | ועל |  |
| 6:37.280 | דבר |  |
| 6:37.281 | זה |  |
| 6:37.282 | צוה |  |
| 6:37.283 | אותנו |  |
| 6:37.284 | בתורה |  |
| 6:37.285 | ואמר |  |
| 6:37.286 | וידעת |  |
| 6:37.287 | היום |  |
| 6:37.288 | והשבות |  |
| 6:37.289 | אל |  |
| 6:37.290 | לבבך |  |
| 6:37.291 | כי |  |
| 6:37.292 | ה | ' |
| 6:37.293 | הוא |  |
| 6:37.294 | האלהים |  |
| 6:37.295 | בשמים |  |
| 6:37.296 | ממעל |  |
| 6:37.297 | ועל |  |
| 6:37.298 | הארץ |  |
| 6:37.299 | מתחת |  |
| 6:37.300 | אין |  |
| 6:37.301 | עוד |  |
| 6:37.302 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

