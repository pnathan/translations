# Translation pack: milhamot-hashem perek-6-19 (Milchamot ha-Shem 6:29–6:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-19`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- עלת **העלות** → "the Cause of causes".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **רחמנא ליצלן** → "the Merciful One save us".
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


## Milchamot ha-Shem 6:29

Hebrew (Wikisource): עד) ראה יראה כל משכיל מעם ישראל האוהב את השי"ת ומאמין בתורתו הנתונה לנו מורשה בכתב ובעל פה כמה טרחו רז"ל בעלי המשנה והתלמוד לבאר הכתובים המדברים באלהות בלשון רבים, שמלבד שתשובת המתעקש בצדו שהוא ית' אחד מיוחד מכל שאר האחדים, אלא שהודיענו רז"ל שעניננו שמורה על כל מיני קריבות אע"פ שהוא שוכן בגבהי מרומים הרחוקים עד מאוד, וכן הוא קדוש בכל מיני קדושות, לא כדעת מחבר הזהר שאמר נעשה ודאי על תרין אתמר, ומספר שם שאימא אמרה לאבא נעשה אדם וכו', וכשחטא אדם הראשון נתגרשה עמו ונפרדה מאבא והרי אינו חד! רחמנא ליצלן מאמונת מינות כזאת וכיוצא בה, הוא שהנביא מוכיח את ישראל כי מספר עריך היו אלהיך ישראל שנחשבו כאלהות כמה מדרגות וצורות שונים המשתלשלים זה מזה בנים ובנות, גם בני בנים בעלי נשים ופילגשים, ויעבדו את הסבות שכבר נתעבו והוגשמו, ולא שתו אל לבם מאמר רז"ל בשבת פרק אמר ר' עקיבה כדגרסינן התם אמר ר' יהושע ב"ל, כשעלה משה למרום אמרו מלאכי השרת לפני הקב"ה רבונו של"ע מה לילוד אשה בינינו וכו' שוב מה כתיב בה כבד את אביך ואת אמך אב ואם יש לכם? הרי בהדיא שמשה רבינו נצח את המלאכים גם בטענה זו שמתמיה עליהם וכי יש לכם אב ואם? כי כולם משרתיו עושי רצונו ואין נולדים מאב ואם. והנה מזה תדין קל וחומר, ומה אם המלאכים אין להם אב ואם, כל שכן וקל וחומר להעלות על לב ולחשוב באלוהינו ית' אב ואם, ובן ובת כמו שחשב המסית נביא השקר מחבר הזהר שכיון להסיתנו ולחדש לנו אמונה הדומה בתארה לאמונת העכו"ם ולזלזל באמונת האחדות המפורשת בתורה בנביאים ובכתובים ובדברי רז"ל כמו שנתבאר לעיל.

| id | word | punctuation after |
|---|---|---|
| 6:29.1 | עד | ) |
| 6:29.2 | ראה |  |
| 6:29.3 | יראה |  |
| 6:29.4 | כל |  |
| 6:29.5 | משכיל |  |
| 6:29.6 | מעם |  |
| 6:29.7 | ישראל |  |
| 6:29.8 | האוהב |  |
| 6:29.9 | את |  |
| 6:29.10 | השי"ת |  |
| 6:29.11 | ומאמין |  |
| 6:29.12 | בתורתו |  |
| 6:29.13 | הנתונה |  |
| 6:29.14 | לנו |  |
| 6:29.15 | מורשה |  |
| 6:29.16 | בכתב |  |
| 6:29.17 | ובעל |  |
| 6:29.18 | פה |  |
| 6:29.19 | כמה |  |
| 6:29.20 | טרחו |  |
| 6:29.21 | רז"ל |  |
| 6:29.22 | בעלי |  |
| 6:29.23 | המשנה |  |
| 6:29.24 | והתלמוד |  |
| 6:29.25 | לבאר |  |
| 6:29.26 | הכתובים |  |
| 6:29.27 | המדברים |  |
| 6:29.28 | באלהות |  |
| 6:29.29 | בלשון |  |
| 6:29.30 | רבים | , |
| 6:29.31 | שמלבד |  |
| 6:29.32 | שתשובת |  |
| 6:29.33 | המתעקש |  |
| 6:29.34 | בצדו |  |
| 6:29.35 | שהוא |  |
| 6:29.36 | ית | ' |
| 6:29.37 | אחד |  |
| 6:29.38 | מיוחד |  |
| 6:29.39 | מכל |  |
| 6:29.40 | שאר |  |
| 6:29.41 | האחדים | , |
| 6:29.42 | אלא |  |
| 6:29.43 | שהודיענו |  |
| 6:29.44 | רז"ל |  |
| 6:29.45 | שעניננו |  |
| 6:29.46 | שמורה |  |
| 6:29.47 | על |  |
| 6:29.48 | כל |  |
| 6:29.49 | מיני |  |
| 6:29.50 | קריבות |  |
| 6:29.51 | אע"פ |  |
| 6:29.52 | שהוא |  |
| 6:29.53 | שוכן |  |
| 6:29.54 | בגבהי |  |
| 6:29.55 | מרומים |  |
| 6:29.56 | הרחוקים |  |
| 6:29.57 | עד |  |
| 6:29.58 | מאוד | , |
| 6:29.59 | וכן |  |
| 6:29.60 | הוא |  |
| 6:29.61 | קדוש |  |
| 6:29.62 | בכל |  |
| 6:29.63 | מיני |  |
| 6:29.64 | קדושות | , |
| 6:29.65 | לא |  |
| 6:29.66 | כדעת |  |
| 6:29.67 | מחבר |  |
| 6:29.68 | הזהר |  |
| 6:29.69 | שאמר |  |
| 6:29.70 | נעשה |  |
| 6:29.71 | ודאי |  |
| 6:29.72 | על |  |
| 6:29.73 | תרין |  |
| 6:29.74 | אתמר | , |
| 6:29.75 | ומספר |  |
| 6:29.76 | שם |  |
| 6:29.77 | שאימא |  |
| 6:29.78 | אמרה |  |
| 6:29.79 | לאבא |  |
| 6:29.80 | נעשה |  |
| 6:29.81 | אדם |  |
| 6:29.82 | וכו | ', |
| 6:29.83 | וכשחטא |  |
| 6:29.84 | אדם |  |
| 6:29.85 | הראשון |  |
| 6:29.86 | נתגרשה |  |
| 6:29.87 | עמו |  |
| 6:29.88 | ונפרדה |  |
| 6:29.89 | מאבא |  |
| 6:29.90 | והרי |  |
| 6:29.91 | אינו |  |
| 6:29.92 | חד | ! |
| 6:29.93 | רחמנא |  |
| 6:29.94 | ליצלן |  |
| 6:29.95 | מאמונת |  |
| 6:29.96 | מינות |  |
| 6:29.97 | כזאת |  |
| 6:29.98 | וכיוצא |  |
| 6:29.99 | בה | , |
| 6:29.100 | הוא |  |
| 6:29.101 | שהנביא |  |
| 6:29.102 | מוכיח |  |
| 6:29.103 | את |  |
| 6:29.104 | ישראל |  |
| 6:29.105 | כי |  |
| 6:29.106 | מספר |  |
| 6:29.107 | עריך |  |
| 6:29.108 | היו |  |
| 6:29.109 | אלהיך |  |
| 6:29.110 | ישראל |  |
| 6:29.111 | שנחשבו |  |
| 6:29.112 | כאלהות |  |
| 6:29.113 | כמה |  |
| 6:29.114 | מדרגות |  |
| 6:29.115 | וצורות |  |
| 6:29.116 | שונים |  |
| 6:29.117 | המשתלשלים |  |
| 6:29.118 | זה |  |
| 6:29.119 | מזה |  |
| 6:29.120 | בנים |  |
| 6:29.121 | ובנות | , |
| 6:29.122 | גם |  |
| 6:29.123 | בני |  |
| 6:29.124 | בנים |  |
| 6:29.125 | בעלי |  |
| 6:29.126 | נשים |  |
| 6:29.127 | ופילגשים | , |
| 6:29.128 | ויעבדו |  |
| 6:29.129 | את |  |
| 6:29.130 | הסבות |  |
| 6:29.131 | שכבר |  |
| 6:29.132 | נתעבו |  |
| 6:29.133 | והוגשמו | , |
| 6:29.134 | ולא |  |
| 6:29.135 | שתו |  |
| 6:29.136 | אל |  |
| 6:29.137 | לבם |  |
| 6:29.138 | מאמר |  |
| 6:29.139 | רז"ל |  |
| 6:29.140 | בשבת |  |
| 6:29.141 | פרק |  |
| 6:29.142 | אמר |  |
| 6:29.143 | ר | ' |
| 6:29.144 | עקיבה |  |
| 6:29.145 | כדגרסינן |  |
| 6:29.146 | התם |  |
| 6:29.147 | אמר |  |
| 6:29.148 | ר | ' |
| 6:29.149 | יהושע |  |
| 6:29.150 | ב"ל | , |
| 6:29.151 | כשעלה |  |
| 6:29.152 | משה |  |
| 6:29.153 | למרום |  |
| 6:29.154 | אמרו |  |
| 6:29.155 | מלאכי |  |
| 6:29.156 | השרת |  |
| 6:29.157 | לפני |  |
| 6:29.158 | הקב"ה |  |
| 6:29.159 | רבונו |  |
| 6:29.160 | של"ע |  |
| 6:29.161 | מה |  |
| 6:29.162 | לילוד |  |
| 6:29.163 | אשה |  |
| 6:29.164 | בינינו |  |
| 6:29.165 | וכו | ' |
| 6:29.166 | שוב |  |
| 6:29.167 | מה |  |
| 6:29.168 | כתיב |  |
| 6:29.169 | בה |  |
| 6:29.170 | כבד |  |
| 6:29.171 | את |  |
| 6:29.172 | אביך |  |
| 6:29.173 | ואת |  |
| 6:29.174 | אמך |  |
| 6:29.175 | אב |  |
| 6:29.176 | ואם |  |
| 6:29.177 | יש |  |
| 6:29.178 | לכם | ? |
| 6:29.179 | הרי |  |
| 6:29.180 | בהדיא |  |
| 6:29.181 | שמשה |  |
| 6:29.182 | רבינו |  |
| 6:29.183 | נצח |  |
| 6:29.184 | את |  |
| 6:29.185 | המלאכים |  |
| 6:29.186 | גם |  |
| 6:29.187 | בטענה |  |
| 6:29.188 | זו |  |
| 6:29.189 | שמתמיה |  |
| 6:29.190 | עליהם |  |
| 6:29.191 | וכי |  |
| 6:29.192 | יש |  |
| 6:29.193 | לכם |  |
| 6:29.194 | אב |  |
| 6:29.195 | ואם | ? |
| 6:29.196 | כי |  |
| 6:29.197 | כולם |  |
| 6:29.198 | משרתיו |  |
| 6:29.199 | עושי |  |
| 6:29.200 | רצונו |  |
| 6:29.201 | ואין |  |
| 6:29.202 | נולדים |  |
| 6:29.203 | מאב |  |
| 6:29.204 | ואם | . |
| 6:29.205 | והנה |  |
| 6:29.206 | מזה |  |
| 6:29.207 | תדין |  |
| 6:29.208 | קל |  |
| 6:29.209 | וחומר | , |
| 6:29.210 | ומה |  |
| 6:29.211 | אם |  |
| 6:29.212 | המלאכים |  |
| 6:29.213 | אין |  |
| 6:29.214 | להם |  |
| 6:29.215 | אב |  |
| 6:29.216 | ואם | , |
| 6:29.217 | כל |  |
| 6:29.218 | שכן |  |
| 6:29.219 | וקל |  |
| 6:29.220 | וחומר |  |
| 6:29.221 | להעלות |  |
| 6:29.222 | על |  |
| 6:29.223 | לב |  |
| 6:29.224 | ולחשוב |  |
| 6:29.225 | באלוהינו |  |
| 6:29.226 | ית | ' |
| 6:29.227 | אב |  |
| 6:29.228 | ואם | , |
| 6:29.229 | ובן |  |
| 6:29.230 | ובת |  |
| 6:29.231 | כמו |  |
| 6:29.232 | שחשב |  |
| 6:29.233 | המסית |  |
| 6:29.234 | נביא |  |
| 6:29.235 | השקר |  |
| 6:29.236 | מחבר |  |
| 6:29.237 | הזהר |  |
| 6:29.238 | שכיון |  |
| 6:29.239 | להסיתנו |  |
| 6:29.240 | ולחדש |  |
| 6:29.241 | לנו |  |
| 6:29.242 | אמונה |  |
| 6:29.243 | הדומה |  |
| 6:29.244 | בתארה |  |
| 6:29.245 | לאמונת |  |
| 6:29.246 | העכו"ם |  |
| 6:29.247 | ולזלזל |  |
| 6:29.248 | באמונת |  |
| 6:29.249 | האחדות |  |
| 6:29.250 | המפורשת |  |
| 6:29.251 | בתורה |  |
| 6:29.252 | בנביאים |  |
| 6:29.253 | ובכתובים |  |
| 6:29.254 | ובדברי |  |
| 6:29.255 | רז"ל |  |
| 6:29.256 | כמו |  |
| 6:29.257 | שנתבאר |  |
| 6:29.258 | לעיל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

