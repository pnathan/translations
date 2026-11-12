# Translation pack: milhamot-hashem perek-5-15 (Milchamot ha-Shem 5:23–5:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-15`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דרך ארץ** → "worldly occupation".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 5:23

Hebrew (Wikisource): עוד שם בריש פרשת לך לך זהו שאמר הכתוב (תהלים מ"ה ח'): "אהבת צדק ותשנא רשע על כן משחך אלהים אלהיך שמן ששון מחבריך", נאמר הפסוק הזה כנגד אברהם אבינו שאהב את הקב"ה ונתקרב תחת כנפי השכינה ושנא עבודה זרה של אביו. לפי שבבית אביו היו עושין צלמים ונותנין לאברהם למכור בשוק, וכשהיה בא אדם לקנות ממנו צלם מה היה עושה היה נוטל הקורנס ומכה על ראש כל אחד מהם ואומר לזה אתה מבקש או זה, וכיון שהיו רואים שהוא מכה אותן, מיד מניחין אותן והולכין להן. והיה משוטט בדעתו ואומר עד מתי אנו משתחוין למעשה ידינו, אין ראוי לעבוד ולהשתחות אלא לארץ שהיא הוציאה פירות ומחיה אותנו. כיון שראה שהיא צריכה לגשמים ואלמלא השמים נפתחים ומורידים גשמים לא היתה מצמחת כלום, חזר ואמר אין ראוי להשתחוות אלא לשמים. חזר וראה את השמש שהיא מאירה לעולם והיא מגדלת צמחים. אמר אין ראוי להשתחות אלא לה. כיון שראה אותה ששקעה אמר אין זה אלוה. חזר וראה את הירח ואת הכוכבים שהם מאירים בלילה, אמר לאלו ראוי לעבוד ולהשתחות. כיון שעלה עמוד השחר בטלו כולם. אמר אין אלו אלוהות, ואמר אלמלא יש להם מנהיג, מפני מה זה שוקע וזה עולה. משל למה הדבר דומה לאחד שהיה מהלך בדרך וראה בירה גדולה וגבוהה ביותר. ביקש להכנס לתוכה חיזר על כל צדדיה ולא מצא פתח. קרא כמה קולות ואין עונה אותו, נשא עיניו וראה כלי צמר אדומים שטוחים על הגג. אחר כך ראה כלי פשתן לבנים. חשכה - ראה נרות דולקות. אמר אותו האיש, ודאי אדם יש בבירה, ואלו לא היה אדם בבירה למה אלו ניטלין ואלו ניתנין? כיון שראה אותו בעל הבירה שהוא מצטער על כך, הציץ עליו ואמר אני בעל הבירה. כך היה אברהם אבינו וכו' עד שהציץ הקב"ה ואמר לו אני בעל העולם וכו' בבראשית רבה וילקוט שמעוני. ובבראשית רבה פרשה ס"ד: "עקב אשר שמע אברהם בקולי ר' יוחנן ור' חנינא תרויהון אמרין בן מ"ח שנה הכיר את בוראו, ריש לקיש אמר בן שלש שנים הכיר אברהם את בוראו, מנין? עקב שמע אברהם בקול בוראו", וכתב עץ יוסף: "נראה דמודו, דתחלת ההכרה בן שלש שנים ולא נגמרה ההכרה אלא עד מ"ח שנים", עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:23.1 | עוד |  |
| 5:23.2 | שם |  |
| 5:23.3 | בריש |  |
| 5:23.4 | פרשת |  |
| 5:23.5 | לך |  |
| 5:23.6 | לך |  |
| 5:23.7 | זהו |  |
| 5:23.8 | שאמר |  |
| 5:23.9 | הכתוב |  |
| 5:23.10 | תהלים |  |
| 5:23.11 | מ"ה |  |
| 5:23.12 | ח | '): |
| 5:23.13 | אהבת |  |
| 5:23.14 | צדק |  |
| 5:23.15 | ותשנא |  |
| 5:23.16 | רשע |  |
| 5:23.17 | על |  |
| 5:23.18 | כן |  |
| 5:23.19 | משחך |  |
| 5:23.20 | אלהים |  |
| 5:23.21 | אלהיך |  |
| 5:23.22 | שמן |  |
| 5:23.23 | ששון |  |
| 5:23.24 | מחבריך | ", |
| 5:23.25 | נאמר |  |
| 5:23.26 | הפסוק |  |
| 5:23.27 | הזה |  |
| 5:23.28 | כנגד |  |
| 5:23.29 | אברהם |  |
| 5:23.30 | אבינו |  |
| 5:23.31 | שאהב |  |
| 5:23.32 | את |  |
| 5:23.33 | הקב"ה |  |
| 5:23.34 | ונתקרב |  |
| 5:23.35 | תחת |  |
| 5:23.36 | כנפי |  |
| 5:23.37 | השכינה |  |
| 5:23.38 | ושנא |  |
| 5:23.39 | עבודה |  |
| 5:23.40 | זרה |  |
| 5:23.41 | של |  |
| 5:23.42 | אביו | . |
| 5:23.43 | לפי |  |
| 5:23.44 | שבבית |  |
| 5:23.45 | אביו |  |
| 5:23.46 | היו |  |
| 5:23.47 | עושין |  |
| 5:23.48 | צלמים |  |
| 5:23.49 | ונותנין |  |
| 5:23.50 | לאברהם |  |
| 5:23.51 | למכור |  |
| 5:23.52 | בשוק | , |
| 5:23.53 | וכשהיה |  |
| 5:23.54 | בא |  |
| 5:23.55 | אדם |  |
| 5:23.56 | לקנות |  |
| 5:23.57 | ממנו |  |
| 5:23.58 | צלם |  |
| 5:23.59 | מה |  |
| 5:23.60 | היה |  |
| 5:23.61 | עושה |  |
| 5:23.62 | היה |  |
| 5:23.63 | נוטל |  |
| 5:23.64 | הקורנס |  |
| 5:23.65 | ומכה |  |
| 5:23.66 | על |  |
| 5:23.67 | ראש |  |
| 5:23.68 | כל |  |
| 5:23.69 | אחד |  |
| 5:23.70 | מהם |  |
| 5:23.71 | ואומר |  |
| 5:23.72 | לזה |  |
| 5:23.73 | אתה |  |
| 5:23.74 | מבקש |  |
| 5:23.75 | או |  |
| 5:23.76 | זה | , |
| 5:23.77 | וכיון |  |
| 5:23.78 | שהיו |  |
| 5:23.79 | רואים |  |
| 5:23.80 | שהוא |  |
| 5:23.81 | מכה |  |
| 5:23.82 | אותן | , |
| 5:23.83 | מיד |  |
| 5:23.84 | מניחין |  |
| 5:23.85 | אותן |  |
| 5:23.86 | והולכין |  |
| 5:23.87 | להן | . |
| 5:23.88 | והיה |  |
| 5:23.89 | משוטט |  |
| 5:23.90 | בדעתו |  |
| 5:23.91 | ואומר |  |
| 5:23.92 | עד |  |
| 5:23.93 | מתי |  |
| 5:23.94 | אנו |  |
| 5:23.95 | משתחוין |  |
| 5:23.96 | למעשה |  |
| 5:23.97 | ידינו | , |
| 5:23.98 | אין |  |
| 5:23.99 | ראוי |  |
| 5:23.100 | לעבוד |  |
| 5:23.101 | ולהשתחות |  |
| 5:23.102 | אלא |  |
| 5:23.103 | לארץ |  |
| 5:23.104 | שהיא |  |
| 5:23.105 | הוציאה |  |
| 5:23.106 | פירות |  |
| 5:23.107 | ומחיה |  |
| 5:23.108 | אותנו | . |
| 5:23.109 | כיון |  |
| 5:23.110 | שראה |  |
| 5:23.111 | שהיא |  |
| 5:23.112 | צריכה |  |
| 5:23.113 | לגשמים |  |
| 5:23.114 | ואלמלא |  |
| 5:23.115 | השמים |  |
| 5:23.116 | נפתחים |  |
| 5:23.117 | ומורידים |  |
| 5:23.118 | גשמים |  |
| 5:23.119 | לא |  |
| 5:23.120 | היתה |  |
| 5:23.121 | מצמחת |  |
| 5:23.122 | כלום | , |
| 5:23.123 | חזר |  |
| 5:23.124 | ואמר |  |
| 5:23.125 | אין |  |
| 5:23.126 | ראוי |  |
| 5:23.127 | להשתחוות |  |
| 5:23.128 | אלא |  |
| 5:23.129 | לשמים | . |
| 5:23.130 | חזר |  |
| 5:23.131 | וראה |  |
| 5:23.132 | את |  |
| 5:23.133 | השמש |  |
| 5:23.134 | שהיא |  |
| 5:23.135 | מאירה |  |
| 5:23.136 | לעולם |  |
| 5:23.137 | והיא |  |
| 5:23.138 | מגדלת |  |
| 5:23.139 | צמחים | . |
| 5:23.140 | אמר |  |
| 5:23.141 | אין |  |
| 5:23.142 | ראוי |  |
| 5:23.143 | להשתחות |  |
| 5:23.144 | אלא |  |
| 5:23.145 | לה | . |
| 5:23.146 | כיון |  |
| 5:23.147 | שראה |  |
| 5:23.148 | אותה |  |
| 5:23.149 | ששקעה |  |
| 5:23.150 | אמר |  |
| 5:23.151 | אין |  |
| 5:23.152 | זה |  |
| 5:23.153 | אלוה | . |
| 5:23.154 | חזר |  |
| 5:23.155 | וראה |  |
| 5:23.156 | את |  |
| 5:23.157 | הירח |  |
| 5:23.158 | ואת |  |
| 5:23.159 | הכוכבים |  |
| 5:23.160 | שהם |  |
| 5:23.161 | מאירים |  |
| 5:23.162 | בלילה | , |
| 5:23.163 | אמר |  |
| 5:23.164 | לאלו |  |
| 5:23.165 | ראוי |  |
| 5:23.166 | לעבוד |  |
| 5:23.167 | ולהשתחות | . |
| 5:23.168 | כיון |  |
| 5:23.169 | שעלה |  |
| 5:23.170 | עמוד |  |
| 5:23.171 | השחר |  |
| 5:23.172 | בטלו |  |
| 5:23.173 | כולם | . |
| 5:23.174 | אמר |  |
| 5:23.175 | אין |  |
| 5:23.176 | אלו |  |
| 5:23.177 | אלוהות | , |
| 5:23.178 | ואמר |  |
| 5:23.179 | אלמלא |  |
| 5:23.180 | יש |  |
| 5:23.181 | להם |  |
| 5:23.182 | מנהיג | , |
| 5:23.183 | מפני |  |
| 5:23.184 | מה |  |
| 5:23.185 | זה |  |
| 5:23.186 | שוקע |  |
| 5:23.187 | וזה |  |
| 5:23.188 | עולה | . |
| 5:23.189 | משל |  |
| 5:23.190 | למה |  |
| 5:23.191 | הדבר |  |
| 5:23.192 | דומה |  |
| 5:23.193 | לאחד |  |
| 5:23.194 | שהיה |  |
| 5:23.195 | מהלך |  |
| 5:23.196 | בדרך |  |
| 5:23.197 | וראה |  |
| 5:23.198 | בירה |  |
| 5:23.199 | גדולה |  |
| 5:23.200 | וגבוהה |  |
| 5:23.201 | ביותר | . |
| 5:23.202 | ביקש |  |
| 5:23.203 | להכנס |  |
| 5:23.204 | לתוכה |  |
| 5:23.205 | חיזר |  |
| 5:23.206 | על |  |
| 5:23.207 | כל |  |
| 5:23.208 | צדדיה |  |
| 5:23.209 | ולא |  |
| 5:23.210 | מצא |  |
| 5:23.211 | פתח | . |
| 5:23.212 | קרא |  |
| 5:23.213 | כמה |  |
| 5:23.214 | קולות |  |
| 5:23.215 | ואין |  |
| 5:23.216 | עונה |  |
| 5:23.217 | אותו | , |
| 5:23.218 | נשא |  |
| 5:23.219 | עיניו |  |
| 5:23.220 | וראה |  |
| 5:23.221 | כלי |  |
| 5:23.222 | צמר |  |
| 5:23.223 | אדומים |  |
| 5:23.224 | שטוחים |  |
| 5:23.225 | על |  |
| 5:23.226 | הגג | . |
| 5:23.227 | אחר |  |
| 5:23.228 | כך |  |
| 5:23.229 | ראה |  |
| 5:23.230 | כלי |  |
| 5:23.231 | פשתן |  |
| 5:23.232 | לבנים | . |
| 5:23.233 | חשכה | - |
| 5:23.234 | ראה |  |
| 5:23.235 | נרות |  |
| 5:23.236 | דולקות | . |
| 5:23.237 | אמר |  |
| 5:23.238 | אותו |  |
| 5:23.239 | האיש | , |
| 5:23.240 | ודאי |  |
| 5:23.241 | אדם |  |
| 5:23.242 | יש |  |
| 5:23.243 | בבירה | , |
| 5:23.244 | ואלו |  |
| 5:23.245 | לא |  |
| 5:23.246 | היה |  |
| 5:23.247 | אדם |  |
| 5:23.248 | בבירה |  |
| 5:23.249 | למה |  |
| 5:23.250 | אלו |  |
| 5:23.251 | ניטלין |  |
| 5:23.252 | ואלו |  |
| 5:23.253 | ניתנין | ? |
| 5:23.254 | כיון |  |
| 5:23.255 | שראה |  |
| 5:23.256 | אותו |  |
| 5:23.257 | בעל |  |
| 5:23.258 | הבירה |  |
| 5:23.259 | שהוא |  |
| 5:23.260 | מצטער |  |
| 5:23.261 | על |  |
| 5:23.262 | כך | , |
| 5:23.263 | הציץ |  |
| 5:23.264 | עליו |  |
| 5:23.265 | ואמר |  |
| 5:23.266 | אני |  |
| 5:23.267 | בעל |  |
| 5:23.268 | הבירה | . |
| 5:23.269 | כך |  |
| 5:23.270 | היה |  |
| 5:23.271 | אברהם |  |
| 5:23.272 | אבינו |  |
| 5:23.273 | וכו | ' |
| 5:23.274 | עד |  |
| 5:23.275 | שהציץ |  |
| 5:23.276 | הקב"ה |  |
| 5:23.277 | ואמר |  |
| 5:23.278 | לו |  |
| 5:23.279 | אני |  |
| 5:23.280 | בעל |  |
| 5:23.281 | העולם |  |
| 5:23.282 | וכו | ' |
| 5:23.283 | בבראשית |  |
| 5:23.284 | רבה |  |
| 5:23.285 | וילקוט |  |
| 5:23.286 | שמעוני | . |
| 5:23.287 | ובבראשית |  |
| 5:23.288 | רבה |  |
| 5:23.289 | פרשה |  |
| 5:23.290 | ס"ד | : |
| 5:23.291 | עקב |  |
| 5:23.292 | אשר |  |
| 5:23.293 | שמע |  |
| 5:23.294 | אברהם |  |
| 5:23.295 | בקולי |  |
| 5:23.296 | ר | ' |
| 5:23.297 | יוחנן |  |
| 5:23.298 | ור | ' |
| 5:23.299 | חנינא |  |
| 5:23.300 | תרויהון |  |
| 5:23.301 | אמרין |  |
| 5:23.302 | בן |  |
| 5:23.303 | מ"ח |  |
| 5:23.304 | שנה |  |
| 5:23.305 | הכיר |  |
| 5:23.306 | את |  |
| 5:23.307 | בוראו | , |
| 5:23.308 | ריש |  |
| 5:23.309 | לקיש |  |
| 5:23.310 | אמר |  |
| 5:23.311 | בן |  |
| 5:23.312 | שלש |  |
| 5:23.313 | שנים |  |
| 5:23.314 | הכיר |  |
| 5:23.315 | אברהם |  |
| 5:23.316 | את |  |
| 5:23.317 | בוראו | , |
| 5:23.318 | מנין | ? |
| 5:23.319 | עקב |  |
| 5:23.320 | שמע |  |
| 5:23.321 | אברהם |  |
| 5:23.322 | בקול |  |
| 5:23.323 | בוראו | ", |
| 5:23.324 | וכתב |  |
| 5:23.325 | עץ |  |
| 5:23.326 | יוסף | : |
| 5:23.327 | נראה |  |
| 5:23.328 | דמודו | , |
| 5:23.329 | דתחלת |  |
| 5:23.330 | ההכרה |  |
| 5:23.331 | בן |  |
| 5:23.332 | שלש |  |
| 5:23.333 | שנים |  |
| 5:23.334 | ולא |  |
| 5:23.335 | נגמרה |  |
| 5:23.336 | ההכרה |  |
| 5:23.337 | אלא |  |
| 5:23.338 | עד |  |
| 5:23.339 | מ"ח |  |
| 5:23.340 | שנים | ", |
| 5:23.341 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

