# Translation pack: milhamot-hashem perek-6-18 (Milchamot ha-Shem 6:28–6:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-18`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **רוע** לב → "wickedness of heart".
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
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 6:28

Hebrew (Wikisource): עג) ובירושלמי פ' הרואה גרסינן המינין שאלו את ר' שמלאי כמה אלוהות בראו את העולם, אמר להם ולי אתם שולאים, לכו ושאלו את אדם הראשון, שנ' כי שאל נא לימים ראשונים וגו' אשר ברא אלהים אדם על הארץ, אין כתיב כאן אלא למן היום אשר ברא אלהים אדם על הארץ, אמרו ליה והא כתיב בראשית ברא אלהים, אמר להו וכי כתיב בראו, אין כתיב אלא ברא, אמר ר' שמלאי כל מקום שפקרו המינין תשובתן בצדן, חזרו ושאלו אותו מה ההן דכתיב נעשה אדם בצלמינו כדמותינו, אמר להם ויבראו אלהים את האדם בצלמם אין כתיב כאן, אלא ויברא אלהים את האדם בצלמו, אמרו לו תלמידיו לאלה דחית בקנה, לנו מה אתה משיב, אמר להם לשעבר אדם נברא מן העפר, וחוה נבראת מן האדם, מאדם ואילך בצלמינו כדמותינו, אי אפשר לאיש בלא אשה, ואי אפשר לאשה בלא איש, אי אפשר לשניהם בלא שכינה, וחזרו ושאלו אותן, מה ההן דכתיב אל אלהים ה', אל אלהים ה' הוא יודע, אמר להם הם יודעים אין כתיב כאן. אלא הוא יודע כתיב, אמרו לו תלמידי הרבי לאלו דחית בקנה, לנו מה אתה משיב, אמר להן שלשתן שם אחד כאיניש דאמר בסיליוס קיסר אנושתום (אדום מושל ושליט). חזרו ושאלו אותו, מהו דכתיב אל אלהים ה' דבר ויקרא ארץ, אמר להן וכי דברו-ויקראו כתוב כאן, אין כתיב אלא דבר ויקראו ארץ, אמרו לו תלמידיו רבי לאלה דחית בקנה, לנו מה אתה משיב אמר להם שלשתן שם אחד הן כאיניש דאמר אומנון בניןן ארכיטקטנן (נ"א ארכי קטונדס ערוך אמן). חזרו ושאלו אותו מהו דכתיב כי אלהים קדושים הוא אמר להם קדושים המה אין כתוב כאן אלא הוא אל קנא הוא, אמרו לו תלמידיו רבי לאלו דחית בקנה. ולנו מה אתה משיב, אמר להם קדוש הוא בכל מיני קדושות, דאמר רבי יודן בשם רבי אחא הקב"ה דרכו בקדושה, דיבורו בקדושה, ישובו בקדושה, אלהים בקדש דרכך, הילוכו בקדושה, הליכות אל מלכי בקדש, מושבו בקדושה אלהים ישב על כסא קדשו, דיבורו בקדושה אלהים דיבר בקדשו, חשיפת דרועו בקדושה, חשף ה' את זרוע קדשו נורא ואדיר בקדושה, מי כמוך נאדר בקדש, חזרו ושאלו אותו מהו ההן דכתיב מי גוי גדול אשר לו אלהים קרובים אליו. אמר להם כה' אלהינו בכל קראינו אליהם אין כתיב כאן, אלא בכל קראינו אליו. אמרו לו תלמידיו רבי לאלו דחית בקנה. לנו מה אתה משיב, אמר להם קרוב הוא בכל מיני קריבות דאמר ר' פנחס בשפ ר' יהודה בר' סימון. ע"ז נראית קרובה ואינה אלא רחוקה. מה טעם ישאוהו על כתף יסבלוהו וגו' סוף דבר אלוהו עמו בבית והוא צועק עד שימות ולא ישמע ולא יושיע מצרתו. אבל הקב"ה נראה רחוק ואין קרוב ממנו, דאמר רבי לוי מהארץ עד לרקיע מהלך חמש מאות שנה. וכן לכל רקיע ורקיע, ואמר רבי חלבו ורבי ברכיה בשם רבי אבא סמוקא אף טלפי החיות מהלך חמש מאות חמש עשרה מנין ישרה, ראה כמה גבוה הוא מעולמו, ואדם נכנס לבית הכנסת ועומד אחורי העמוד ומתפלל בלחישה והקב"ה מאזין את תפלתו שנאמר וחנה היא מדברת אל לבה רק שפתיה נעות והאזין הקב"ה תפלתה. וכן כל בריותיו שנ' תפלה לעני כי יעטף כאדם המשיח באזן חבירו והוא שומע, וכי יש לך אלוה קרוב משה שהוא קרוב לבריותיו מפה לאזן ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:28.1 | עג | ) |
| 6:28.2 | ובירושלמי |  |
| 6:28.3 | פ | ' |
| 6:28.4 | הרואה |  |
| 6:28.5 | גרסינן |  |
| 6:28.6 | המינין |  |
| 6:28.7 | שאלו |  |
| 6:28.8 | את |  |
| 6:28.9 | ר | ' |
| 6:28.10 | שמלאי |  |
| 6:28.11 | כמה |  |
| 6:28.12 | אלוהות |  |
| 6:28.13 | בראו |  |
| 6:28.14 | את |  |
| 6:28.15 | העולם | , |
| 6:28.16 | אמר |  |
| 6:28.17 | להם |  |
| 6:28.18 | ולי |  |
| 6:28.19 | אתם |  |
| 6:28.20 | שולאים | , |
| 6:28.21 | לכו |  |
| 6:28.22 | ושאלו |  |
| 6:28.23 | את |  |
| 6:28.24 | אדם |  |
| 6:28.25 | הראשון | , |
| 6:28.26 | שנ | ' |
| 6:28.27 | כי |  |
| 6:28.28 | שאל |  |
| 6:28.29 | נא |  |
| 6:28.30 | לימים |  |
| 6:28.31 | ראשונים |  |
| 6:28.32 | וגו | ' |
| 6:28.33 | אשר |  |
| 6:28.34 | ברא |  |
| 6:28.35 | אלהים |  |
| 6:28.36 | אדם |  |
| 6:28.37 | על |  |
| 6:28.38 | הארץ | , |
| 6:28.39 | אין |  |
| 6:28.40 | כתיב |  |
| 6:28.41 | כאן |  |
| 6:28.42 | אלא |  |
| 6:28.43 | למן |  |
| 6:28.44 | היום |  |
| 6:28.45 | אשר |  |
| 6:28.46 | ברא |  |
| 6:28.47 | אלהים |  |
| 6:28.48 | אדם |  |
| 6:28.49 | על |  |
| 6:28.50 | הארץ | , |
| 6:28.51 | אמרו |  |
| 6:28.52 | ליה |  |
| 6:28.53 | והא |  |
| 6:28.54 | כתיב |  |
| 6:28.55 | בראשית |  |
| 6:28.56 | ברא |  |
| 6:28.57 | אלהים | , |
| 6:28.58 | אמר |  |
| 6:28.59 | להו |  |
| 6:28.60 | וכי |  |
| 6:28.61 | כתיב |  |
| 6:28.62 | בראו | , |
| 6:28.63 | אין |  |
| 6:28.64 | כתיב |  |
| 6:28.65 | אלא |  |
| 6:28.66 | ברא | , |
| 6:28.67 | אמר |  |
| 6:28.68 | ר | ' |
| 6:28.69 | שמלאי |  |
| 6:28.70 | כל |  |
| 6:28.71 | מקום |  |
| 6:28.72 | שפקרו |  |
| 6:28.73 | המינין |  |
| 6:28.74 | תשובתן |  |
| 6:28.75 | בצדן | , |
| 6:28.76 | חזרו |  |
| 6:28.77 | ושאלו |  |
| 6:28.78 | אותו |  |
| 6:28.79 | מה |  |
| 6:28.80 | ההן |  |
| 6:28.81 | דכתיב |  |
| 6:28.82 | נעשה |  |
| 6:28.83 | אדם |  |
| 6:28.84 | בצלמינו |  |
| 6:28.85 | כדמותינו | , |
| 6:28.86 | אמר |  |
| 6:28.87 | להם |  |
| 6:28.88 | ויבראו |  |
| 6:28.89 | אלהים |  |
| 6:28.90 | את |  |
| 6:28.91 | האדם |  |
| 6:28.92 | בצלמם |  |
| 6:28.93 | אין |  |
| 6:28.94 | כתיב |  |
| 6:28.95 | כאן | , |
| 6:28.96 | אלא |  |
| 6:28.97 | ויברא |  |
| 6:28.98 | אלהים |  |
| 6:28.99 | את |  |
| 6:28.100 | האדם |  |
| 6:28.101 | בצלמו | , |
| 6:28.102 | אמרו |  |
| 6:28.103 | לו |  |
| 6:28.104 | תלמידיו |  |
| 6:28.105 | לאלה |  |
| 6:28.106 | דחית |  |
| 6:28.107 | בקנה | , |
| 6:28.108 | לנו |  |
| 6:28.109 | מה |  |
| 6:28.110 | אתה |  |
| 6:28.111 | משיב | , |
| 6:28.112 | אמר |  |
| 6:28.113 | להם |  |
| 6:28.114 | לשעבר |  |
| 6:28.115 | אדם |  |
| 6:28.116 | נברא |  |
| 6:28.117 | מן |  |
| 6:28.118 | העפר | , |
| 6:28.119 | וחוה |  |
| 6:28.120 | נבראת |  |
| 6:28.121 | מן |  |
| 6:28.122 | האדם | , |
| 6:28.123 | מאדם |  |
| 6:28.124 | ואילך |  |
| 6:28.125 | בצלמינו |  |
| 6:28.126 | כדמותינו | , |
| 6:28.127 | אי |  |
| 6:28.128 | אפשר |  |
| 6:28.129 | לאיש |  |
| 6:28.130 | בלא |  |
| 6:28.131 | אשה | , |
| 6:28.132 | ואי |  |
| 6:28.133 | אפשר |  |
| 6:28.134 | לאשה |  |
| 6:28.135 | בלא |  |
| 6:28.136 | איש | , |
| 6:28.137 | אי |  |
| 6:28.138 | אפשר |  |
| 6:28.139 | לשניהם |  |
| 6:28.140 | בלא |  |
| 6:28.141 | שכינה | , |
| 6:28.142 | וחזרו |  |
| 6:28.143 | ושאלו |  |
| 6:28.144 | אותן | , |
| 6:28.145 | מה |  |
| 6:28.146 | ההן |  |
| 6:28.147 | דכתיב |  |
| 6:28.148 | אל |  |
| 6:28.149 | אלהים |  |
| 6:28.150 | ה | ', |
| 6:28.151 | אל |  |
| 6:28.152 | אלהים |  |
| 6:28.153 | ה | ' |
| 6:28.154 | הוא |  |
| 6:28.155 | יודע | , |
| 6:28.156 | אמר |  |
| 6:28.157 | להם |  |
| 6:28.158 | הם |  |
| 6:28.159 | יודעים |  |
| 6:28.160 | אין |  |
| 6:28.161 | כתיב |  |
| 6:28.162 | כאן | . |
| 6:28.163 | אלא |  |
| 6:28.164 | הוא |  |
| 6:28.165 | יודע |  |
| 6:28.166 | כתיב | , |
| 6:28.167 | אמרו |  |
| 6:28.168 | לו |  |
| 6:28.169 | תלמידי |  |
| 6:28.170 | הרבי |  |
| 6:28.171 | לאלו |  |
| 6:28.172 | דחית |  |
| 6:28.173 | בקנה | , |
| 6:28.174 | לנו |  |
| 6:28.175 | מה |  |
| 6:28.176 | אתה |  |
| 6:28.177 | משיב | , |
| 6:28.178 | אמר |  |
| 6:28.179 | להן |  |
| 6:28.180 | שלשתן |  |
| 6:28.181 | שם |  |
| 6:28.182 | אחד |  |
| 6:28.183 | כאיניש |  |
| 6:28.184 | דאמר |  |
| 6:28.185 | בסיליוס |  |
| 6:28.186 | קיסר |  |
| 6:28.187 | אנושתום |  |
| 6:28.188 | אדום |  |
| 6:28.189 | מושל |  |
| 6:28.190 | ושליט | ). |
| 6:28.191 | חזרו |  |
| 6:28.192 | ושאלו |  |
| 6:28.193 | אותו | , |
| 6:28.194 | מהו |  |
| 6:28.195 | דכתיב |  |
| 6:28.196 | אל |  |
| 6:28.197 | אלהים |  |
| 6:28.198 | ה | ' |
| 6:28.199 | דבר |  |
| 6:28.200 | ויקרא |  |
| 6:28.201 | ארץ | , |
| 6:28.202 | אמר |  |
| 6:28.203 | להן |  |
| 6:28.204 | וכי |  |
| 6:28.205 | דברו-ויקראו |  |
| 6:28.206 | כתוב |  |
| 6:28.207 | כאן | , |
| 6:28.208 | אין |  |
| 6:28.209 | כתיב |  |
| 6:28.210 | אלא |  |
| 6:28.211 | דבר |  |
| 6:28.212 | ויקראו |  |
| 6:28.213 | ארץ | , |
| 6:28.214 | אמרו |  |
| 6:28.215 | לו |  |
| 6:28.216 | תלמידיו |  |
| 6:28.217 | רבי |  |
| 6:28.218 | לאלה |  |
| 6:28.219 | דחית |  |
| 6:28.220 | בקנה | , |
| 6:28.221 | לנו |  |
| 6:28.222 | מה |  |
| 6:28.223 | אתה |  |
| 6:28.224 | משיב |  |
| 6:28.225 | אמר |  |
| 6:28.226 | להם |  |
| 6:28.227 | שלשתן |  |
| 6:28.228 | שם |  |
| 6:28.229 | אחד |  |
| 6:28.230 | הן |  |
| 6:28.231 | כאיניש |  |
| 6:28.232 | דאמר |  |
| 6:28.233 | אומנון |  |
| 6:28.234 | בניןן |  |
| 6:28.235 | ארכיטקטנן |  |
| 6:28.236 | נ"א |  |
| 6:28.237 | ארכי |  |
| 6:28.238 | קטונדס |  |
| 6:28.239 | ערוך |  |
| 6:28.240 | אמן | ). |
| 6:28.241 | חזרו |  |
| 6:28.242 | ושאלו |  |
| 6:28.243 | אותו |  |
| 6:28.244 | מהו |  |
| 6:28.245 | דכתיב |  |
| 6:28.246 | כי |  |
| 6:28.247 | אלהים |  |
| 6:28.248 | קדושים |  |
| 6:28.249 | הוא |  |
| 6:28.250 | אמר |  |
| 6:28.251 | להם |  |
| 6:28.252 | קדושים |  |
| 6:28.253 | המה |  |
| 6:28.254 | אין |  |
| 6:28.255 | כתוב |  |
| 6:28.256 | כאן |  |
| 6:28.257 | אלא |  |
| 6:28.258 | הוא |  |
| 6:28.259 | אל |  |
| 6:28.260 | קנא |  |
| 6:28.261 | הוא | , |
| 6:28.262 | אמרו |  |
| 6:28.263 | לו |  |
| 6:28.264 | תלמידיו |  |
| 6:28.265 | רבי |  |
| 6:28.266 | לאלו |  |
| 6:28.267 | דחית |  |
| 6:28.268 | בקנה | . |
| 6:28.269 | ולנו |  |
| 6:28.270 | מה |  |
| 6:28.271 | אתה |  |
| 6:28.272 | משיב | , |
| 6:28.273 | אמר |  |
| 6:28.274 | להם |  |
| 6:28.275 | קדוש |  |
| 6:28.276 | הוא |  |
| 6:28.277 | בכל |  |
| 6:28.278 | מיני |  |
| 6:28.279 | קדושות | , |
| 6:28.280 | דאמר |  |
| 6:28.281 | רבי |  |
| 6:28.282 | יודן |  |
| 6:28.283 | בשם |  |
| 6:28.284 | רבי |  |
| 6:28.285 | אחא |  |
| 6:28.286 | הקב"ה |  |
| 6:28.287 | דרכו |  |
| 6:28.288 | בקדושה | , |
| 6:28.289 | דיבורו |  |
| 6:28.290 | בקדושה | , |
| 6:28.291 | ישובו |  |
| 6:28.292 | בקדושה | , |
| 6:28.293 | אלהים |  |
| 6:28.294 | בקדש |  |
| 6:28.295 | דרכך | , |
| 6:28.296 | הילוכו |  |
| 6:28.297 | בקדושה | , |
| 6:28.298 | הליכות |  |
| 6:28.299 | אל |  |
| 6:28.300 | מלכי |  |
| 6:28.301 | בקדש | , |
| 6:28.302 | מושבו |  |
| 6:28.303 | בקדושה |  |
| 6:28.304 | אלהים |  |
| 6:28.305 | ישב |  |
| 6:28.306 | על |  |
| 6:28.307 | כסא |  |
| 6:28.308 | קדשו | , |
| 6:28.309 | דיבורו |  |
| 6:28.310 | בקדושה |  |
| 6:28.311 | אלהים |  |
| 6:28.312 | דיבר |  |
| 6:28.313 | בקדשו | , |
| 6:28.314 | חשיפת |  |
| 6:28.315 | דרועו |  |
| 6:28.316 | בקדושה | , |
| 6:28.317 | חשף |  |
| 6:28.318 | ה | ' |
| 6:28.319 | את |  |
| 6:28.320 | זרוע |  |
| 6:28.321 | קדשו |  |
| 6:28.322 | נורא |  |
| 6:28.323 | ואדיר |  |
| 6:28.324 | בקדושה | , |
| 6:28.325 | מי |  |
| 6:28.326 | כמוך |  |
| 6:28.327 | נאדר |  |
| 6:28.328 | בקדש | , |
| 6:28.329 | חזרו |  |
| 6:28.330 | ושאלו |  |
| 6:28.331 | אותו |  |
| 6:28.332 | מהו |  |
| 6:28.333 | ההן |  |
| 6:28.334 | דכתיב |  |
| 6:28.335 | מי |  |
| 6:28.336 | גוי |  |
| 6:28.337 | גדול |  |
| 6:28.338 | אשר |  |
| 6:28.339 | לו |  |
| 6:28.340 | אלהים |  |
| 6:28.341 | קרובים |  |
| 6:28.342 | אליו | . |
| 6:28.343 | אמר |  |
| 6:28.344 | להם |  |
| 6:28.345 | כה | ' |
| 6:28.346 | אלהינו |  |
| 6:28.347 | בכל |  |
| 6:28.348 | קראינו |  |
| 6:28.349 | אליהם |  |
| 6:28.350 | אין |  |
| 6:28.351 | כתיב |  |
| 6:28.352 | כאן | , |
| 6:28.353 | אלא |  |
| 6:28.354 | בכל |  |
| 6:28.355 | קראינו |  |
| 6:28.356 | אליו | . |
| 6:28.357 | אמרו |  |
| 6:28.358 | לו |  |
| 6:28.359 | תלמידיו |  |
| 6:28.360 | רבי |  |
| 6:28.361 | לאלו |  |
| 6:28.362 | דחית |  |
| 6:28.363 | בקנה | . |
| 6:28.364 | לנו |  |
| 6:28.365 | מה |  |
| 6:28.366 | אתה |  |
| 6:28.367 | משיב | , |
| 6:28.368 | אמר |  |
| 6:28.369 | להם |  |
| 6:28.370 | קרוב |  |
| 6:28.371 | הוא |  |
| 6:28.372 | בכל |  |
| 6:28.373 | מיני |  |
| 6:28.374 | קריבות |  |
| 6:28.375 | דאמר |  |
| 6:28.376 | ר | ' |
| 6:28.377 | פנחס |  |
| 6:28.378 | בשפ |  |
| 6:28.379 | ר | ' |
| 6:28.380 | יהודה |  |
| 6:28.381 | בר | ' |
| 6:28.382 | סימון | . |
| 6:28.383 | ע"ז |  |
| 6:28.384 | נראית |  |
| 6:28.385 | קרובה |  |
| 6:28.386 | ואינה |  |
| 6:28.387 | אלא |  |
| 6:28.388 | רחוקה | . |
| 6:28.389 | מה |  |
| 6:28.390 | טעם |  |
| 6:28.391 | ישאוהו |  |
| 6:28.392 | על |  |
| 6:28.393 | כתף |  |
| 6:28.394 | יסבלוהו |  |
| 6:28.395 | וגו | ' |
| 6:28.396 | סוף |  |
| 6:28.397 | דבר |  |
| 6:28.398 | אלוהו |  |
| 6:28.399 | עמו |  |
| 6:28.400 | בבית |  |
| 6:28.401 | והוא |  |
| 6:28.402 | צועק |  |
| 6:28.403 | עד |  |
| 6:28.404 | שימות |  |
| 6:28.405 | ולא |  |
| 6:28.406 | ישמע |  |
| 6:28.407 | ולא |  |
| 6:28.408 | יושיע |  |
| 6:28.409 | מצרתו | . |
| 6:28.410 | אבל |  |
| 6:28.411 | הקב"ה |  |
| 6:28.412 | נראה |  |
| 6:28.413 | רחוק |  |
| 6:28.414 | ואין |  |
| 6:28.415 | קרוב |  |
| 6:28.416 | ממנו | , |
| 6:28.417 | דאמר |  |
| 6:28.418 | רבי |  |
| 6:28.419 | לוי |  |
| 6:28.420 | מהארץ |  |
| 6:28.421 | עד |  |
| 6:28.422 | לרקיע |  |
| 6:28.423 | מהלך |  |
| 6:28.424 | חמש |  |
| 6:28.425 | מאות |  |
| 6:28.426 | שנה | . |
| 6:28.427 | וכן |  |
| 6:28.428 | לכל |  |
| 6:28.429 | רקיע |  |
| 6:28.430 | ורקיע | , |
| 6:28.431 | ואמר |  |
| 6:28.432 | רבי |  |
| 6:28.433 | חלבו |  |
| 6:28.434 | ורבי |  |
| 6:28.435 | ברכיה |  |
| 6:28.436 | בשם |  |
| 6:28.437 | רבי |  |
| 6:28.438 | אבא |  |
| 6:28.439 | סמוקא |  |
| 6:28.440 | אף |  |
| 6:28.441 | טלפי |  |
| 6:28.442 | החיות |  |
| 6:28.443 | מהלך |  |
| 6:28.444 | חמש |  |
| 6:28.445 | מאות |  |
| 6:28.446 | חמש |  |
| 6:28.447 | עשרה |  |
| 6:28.448 | מנין |  |
| 6:28.449 | ישרה | , |
| 6:28.450 | ראה |  |
| 6:28.451 | כמה |  |
| 6:28.452 | גבוה |  |
| 6:28.453 | הוא |  |
| 6:28.454 | מעולמו | , |
| 6:28.455 | ואדם |  |
| 6:28.456 | נכנס |  |
| 6:28.457 | לבית |  |
| 6:28.458 | הכנסת |  |
| 6:28.459 | ועומד |  |
| 6:28.460 | אחורי |  |
| 6:28.461 | העמוד |  |
| 6:28.462 | ומתפלל |  |
| 6:28.463 | בלחישה |  |
| 6:28.464 | והקב"ה |  |
| 6:28.465 | מאזין |  |
| 6:28.466 | את |  |
| 6:28.467 | תפלתו |  |
| 6:28.468 | שנאמר |  |
| 6:28.469 | וחנה |  |
| 6:28.470 | היא |  |
| 6:28.471 | מדברת |  |
| 6:28.472 | אל |  |
| 6:28.473 | לבה |  |
| 6:28.474 | רק |  |
| 6:28.475 | שפתיה |  |
| 6:28.476 | נעות |  |
| 6:28.477 | והאזין |  |
| 6:28.478 | הקב"ה |  |
| 6:28.479 | תפלתה | . |
| 6:28.480 | וכן |  |
| 6:28.481 | כל |  |
| 6:28.482 | בריותיו |  |
| 6:28.483 | שנ | ' |
| 6:28.484 | תפלה |  |
| 6:28.485 | לעני |  |
| 6:28.486 | כי |  |
| 6:28.487 | יעטף |  |
| 6:28.488 | כאדם |  |
| 6:28.489 | המשיח |  |
| 6:28.490 | באזן |  |
| 6:28.491 | חבירו |  |
| 6:28.492 | והוא |  |
| 6:28.493 | שומע | , |
| 6:28.494 | וכי |  |
| 6:28.495 | יש |  |
| 6:28.496 | לך |  |
| 6:28.497 | אלוה |  |
| 6:28.498 | קרוב |  |
| 6:28.499 | משה |  |
| 6:28.500 | שהוא |  |
| 6:28.501 | קרוב |  |
| 6:28.502 | לבריותיו |  |
| 6:28.503 | מפה |  |
| 6:28.504 | לאזן |  |
| 6:28.505 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

