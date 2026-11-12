# Translation pack: milhamot-hashem perek-10-06 (Milchamot ha-Shem 10:7–10:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-06`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ע"ה** → "peace be upon him" (honorific).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
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
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **והוא שלא** → "provided that".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- קורא **נעים / נעלם** → "pleasant reader".
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
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **העתקתו / תרגום**: his own Hebrew translations (10:7–10:9; "[תרגום הרב
  המחבר: …]") are separate segments, each from its own Hebrew; a difference
  in sense in `apparatus`.
- **המיחדים / אלתוחיד**: "כת המיחדים" = "אהל אלתוחיד" → "the upholders of the
  unity".
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


## Milchamot ha-Shem 10:7

Hebrew (Wikisource): העתק לשון ערבי בסי צ״ו ל' הרס״ג. העתקתו ללשון הקודש. וכמה רחוק ההקש שיהיה זה השטן מלאך, לפי שמוסכם אצלנו כת המיחדים שהבורא ית' ברא מלאכיו קרובים, לפי שגלוי לפניו שלא ימרו את דברו, כמו שבחר בנביאו ששלח, בידעו שלא ימרוהו, והקנאה בצדיקים להרע מהעבירות הגדולות, ומי שחשב שזה המשטין לאיוב היה מלאך, הנה שם המלאכים ממרים, היפך המוסכם אצלינו, ותדע יעזרך השי״ת שיש לנפש ג' כחות, ההשכלה, והקנאה, היא אשר בה יתעורר וישתדל להנקם מאויביו ולמלא מאויו במותרות והתאוה, היא אשר בה ישתוקק הבעל חי למזון ולמשגל ושאר הענינים, ותדע עם זה כי צדקת המשכיל הוא שתהינה שתי המדות הללו רצוני לומר הקנאה והתאוה נכנעים לשכל, ואם יגבר אחד מהם על השכל הנה זה עול וחטאת, והנה מי שחשב שאפשר למלאכים לקטרג שהוא אחד ממיני הקנאה, יתחייב אצלו ליחס להם פעולות התאות, והם האכילה ושתיה והמשגל ושאר הדברים הפחותים ומזולזלים וטנופים, שכבר חסך השי״ת מלאכיו מהם, בידעו שלא ימרוהו, ואי אפשר להיות היפך ממה שקדם בידיעתו ית' שאלו יהיה הענין היפך מה שהנחנו, כבר היה גלוי לפניו ית' שכן יהיה, וכבר הגיעני מקצת החיצונים שחלק בזה עד שחשב שזה השטן מלאך, ושהקנאה והרכילות אפשר להיות בפועל מן המלאכים כמו שפי' במעשה איוב וקנאת השטן לו, וכאשר ראה שאמונתו בזה לא תכלול האכילה והשתיה במלאכים קיים בדעתו ג״כ אפשרותו בהם, שוב ראה עוד שכמו שאלו העניינים אפשר היותם בהםמ יחוייב בהם התאוה והמשגל וקיימם בהם ג״כ, ואמר בענין ויראו בני האלהים את בנות האדם כי טבת הנה, הם המלאכים הבאים על בנות האדם בנפלם מן השמים והיו ככל האדם, והביא דמיון ראיה לדבריו ממה שנא' אני אמרתי אלהים אתם ובני עליון כלכם אכן כאדם תמותון וכאחד השרים תפלו ! ולא השאיר מדה פחותה ומגונה עד שיחסה להם ! כל זה הולך על פי שטתו, והרכיב כל אלה הדעות הזרות והנפסדות ממעוט הבנתו בפירוש הדברים בלה״ק, והנני מבאר טעותו בכל מה שהניח בע״ה, דע יעזרך השם, שמה שהביא לזה המתחיל המגיס דעתו הוא וכיוצא בו לומר שהשטן הוא מלאך, הוא אמרו ית' לשטן הנה כל אשר לו בידיך רק אליו אל תשלח ידך, וגם אמרו אך את נפשו שמור, וגם אמרו ויך את איוב בשחין רע, וחשב שכל זה מיוחס אל השטן, והנני מבאר ביאור מספיק שבו יסורו אלו הספקות הג' אשר דמה ואומר כי הרצון והחפץ בלשון בני ישראל יקרא יד, כמו שאמר דוד לאשה הקובלת לפניו היד יואב איתך בכל זאת ענינו חפץ ורצון יואב וכו' ועל פי זה יהיה אמרו כאן הנה כל אשר לו בידיך בחפצך ורצונך וכו' ובסור ג' ספיקות הללו אומר שטעות זה המפרש בענין ויראו בני האלהים את בנות האדם יותר זר ממה שאמר בענין איוב, והוא שלא ידע כלל ממה שמשתמשין ישראל בלשון בענין עיני העדה וההמוּן, שדרכם לומר מעשה הגדולים, כך ומעשה המון העם כך, וכולם אנשים מבני אדם, ואמנם הבדילם לגדוללתם ומעלתם, ועל פי זה יהיה אמרו ויראו בני האלהים אנשים ג״כ, אבל הבדילם מן ההמון לחשיבותם ומעלתם, וזה ביאור יפה לשלול מהקרובים לשו״ת הפחיתות והזרות, ומה שיחזק זה הפירוש אמרו אחר זה וירא ה' כי רבה רעת האדם, ולא גינה המלאכים. וגם אמרו בסמוך אמחה את האדם ולא ענש המלאכים עמהם ע״כ.

| id | word | punctuation after |
|---|---|---|
| 10:7.1 | העתק |  |
| 10:7.2 | לשון |  |
| 10:7.3 | ערבי |  |
| 10:7.4 | בסי |  |
| 10:7.5 | צ״ו |  |
| 10:7.6 | ל | ' |
| 10:7.7 | הרס״ג | . |
| 10:7.8 | העתקתו |  |
| 10:7.9 | ללשון |  |
| 10:7.10 | הקודש | . |
| 10:7.11 | וכמה |  |
| 10:7.12 | רחוק |  |
| 10:7.13 | ההקש |  |
| 10:7.14 | שיהיה |  |
| 10:7.15 | זה |  |
| 10:7.16 | השטן |  |
| 10:7.17 | מלאך | , |
| 10:7.18 | לפי |  |
| 10:7.19 | שמוסכם |  |
| 10:7.20 | אצלנו |  |
| 10:7.21 | כת |  |
| 10:7.22 | המיחדים |  |
| 10:7.23 | שהבורא |  |
| 10:7.24 | ית | ' |
| 10:7.25 | ברא |  |
| 10:7.26 | מלאכיו |  |
| 10:7.27 | קרובים | , |
| 10:7.28 | לפי |  |
| 10:7.29 | שגלוי |  |
| 10:7.30 | לפניו |  |
| 10:7.31 | שלא |  |
| 10:7.32 | ימרו |  |
| 10:7.33 | את |  |
| 10:7.34 | דברו | , |
| 10:7.35 | כמו |  |
| 10:7.36 | שבחר |  |
| 10:7.37 | בנביאו |  |
| 10:7.38 | ששלח | , |
| 10:7.39 | בידעו |  |
| 10:7.40 | שלא |  |
| 10:7.41 | ימרוהו | , |
| 10:7.42 | והקנאה |  |
| 10:7.43 | בצדיקים |  |
| 10:7.44 | להרע |  |
| 10:7.45 | מהעבירות |  |
| 10:7.46 | הגדולות | , |
| 10:7.47 | ומי |  |
| 10:7.48 | שחשב |  |
| 10:7.49 | שזה |  |
| 10:7.50 | המשטין |  |
| 10:7.51 | לאיוב |  |
| 10:7.52 | היה |  |
| 10:7.53 | מלאך | , |
| 10:7.54 | הנה |  |
| 10:7.55 | שם |  |
| 10:7.56 | המלאכים |  |
| 10:7.57 | ממרים | , |
| 10:7.58 | היפך |  |
| 10:7.59 | המוסכם |  |
| 10:7.60 | אצלינו | , |
| 10:7.61 | ותדע |  |
| 10:7.62 | יעזרך |  |
| 10:7.63 | השי״ת |  |
| 10:7.64 | שיש |  |
| 10:7.65 | לנפש |  |
| 10:7.66 | ג | ' |
| 10:7.67 | כחות | , |
| 10:7.68 | ההשכלה | , |
| 10:7.69 | והקנאה | , |
| 10:7.70 | היא |  |
| 10:7.71 | אשר |  |
| 10:7.72 | בה |  |
| 10:7.73 | יתעורר |  |
| 10:7.74 | וישתדל |  |
| 10:7.75 | להנקם |  |
| 10:7.76 | מאויביו |  |
| 10:7.77 | ולמלא |  |
| 10:7.78 | מאויו |  |
| 10:7.79 | במותרות |  |
| 10:7.80 | והתאוה | , |
| 10:7.81 | היא |  |
| 10:7.82 | אשר |  |
| 10:7.83 | בה |  |
| 10:7.84 | ישתוקק |  |
| 10:7.85 | הבעל |  |
| 10:7.86 | חי |  |
| 10:7.87 | למזון |  |
| 10:7.88 | ולמשגל |  |
| 10:7.89 | ושאר |  |
| 10:7.90 | הענינים | , |
| 10:7.91 | ותדע |  |
| 10:7.92 | עם |  |
| 10:7.93 | זה |  |
| 10:7.94 | כי |  |
| 10:7.95 | צדקת |  |
| 10:7.96 | המשכיל |  |
| 10:7.97 | הוא |  |
| 10:7.98 | שתהינה |  |
| 10:7.99 | שתי |  |
| 10:7.100 | המדות |  |
| 10:7.101 | הללו |  |
| 10:7.102 | רצוני |  |
| 10:7.103 | לומר |  |
| 10:7.104 | הקנאה |  |
| 10:7.105 | והתאוה |  |
| 10:7.106 | נכנעים |  |
| 10:7.107 | לשכל | , |
| 10:7.108 | ואם |  |
| 10:7.109 | יגבר |  |
| 10:7.110 | אחד |  |
| 10:7.111 | מהם |  |
| 10:7.112 | על |  |
| 10:7.113 | השכל |  |
| 10:7.114 | הנה |  |
| 10:7.115 | זה |  |
| 10:7.116 | עול |  |
| 10:7.117 | וחטאת | , |
| 10:7.118 | והנה |  |
| 10:7.119 | מי |  |
| 10:7.120 | שחשב |  |
| 10:7.121 | שאפשר |  |
| 10:7.122 | למלאכים |  |
| 10:7.123 | לקטרג |  |
| 10:7.124 | שהוא |  |
| 10:7.125 | אחד |  |
| 10:7.126 | ממיני |  |
| 10:7.127 | הקנאה | , |
| 10:7.128 | יתחייב |  |
| 10:7.129 | אצלו |  |
| 10:7.130 | ליחס |  |
| 10:7.131 | להם |  |
| 10:7.132 | פעולות |  |
| 10:7.133 | התאות | , |
| 10:7.134 | והם |  |
| 10:7.135 | האכילה |  |
| 10:7.136 | ושתיה |  |
| 10:7.137 | והמשגל |  |
| 10:7.138 | ושאר |  |
| 10:7.139 | הדברים |  |
| 10:7.140 | הפחותים |  |
| 10:7.141 | ומזולזלים |  |
| 10:7.142 | וטנופים | , |
| 10:7.143 | שכבר |  |
| 10:7.144 | חסך |  |
| 10:7.145 | השי״ת |  |
| 10:7.146 | מלאכיו |  |
| 10:7.147 | מהם | , |
| 10:7.148 | בידעו |  |
| 10:7.149 | שלא |  |
| 10:7.150 | ימרוהו | , |
| 10:7.151 | ואי |  |
| 10:7.152 | אפשר |  |
| 10:7.153 | להיות |  |
| 10:7.154 | היפך |  |
| 10:7.155 | ממה |  |
| 10:7.156 | שקדם |  |
| 10:7.157 | בידיעתו |  |
| 10:7.158 | ית | ' |
| 10:7.159 | שאלו |  |
| 10:7.160 | יהיה |  |
| 10:7.161 | הענין |  |
| 10:7.162 | היפך |  |
| 10:7.163 | מה |  |
| 10:7.164 | שהנחנו | , |
| 10:7.165 | כבר |  |
| 10:7.166 | היה |  |
| 10:7.167 | גלוי |  |
| 10:7.168 | לפניו |  |
| 10:7.169 | ית | ' |
| 10:7.170 | שכן |  |
| 10:7.171 | יהיה | , |
| 10:7.172 | וכבר |  |
| 10:7.173 | הגיעני |  |
| 10:7.174 | מקצת |  |
| 10:7.175 | החיצונים |  |
| 10:7.176 | שחלק |  |
| 10:7.177 | בזה |  |
| 10:7.178 | עד |  |
| 10:7.179 | שחשב |  |
| 10:7.180 | שזה |  |
| 10:7.181 | השטן |  |
| 10:7.182 | מלאך | , |
| 10:7.183 | ושהקנאה |  |
| 10:7.184 | והרכילות |  |
| 10:7.185 | אפשר |  |
| 10:7.186 | להיות |  |
| 10:7.187 | בפועל |  |
| 10:7.188 | מן |  |
| 10:7.189 | המלאכים |  |
| 10:7.190 | כמו |  |
| 10:7.191 | שפי | ' |
| 10:7.192 | במעשה |  |
| 10:7.193 | איוב |  |
| 10:7.194 | וקנאת |  |
| 10:7.195 | השטן |  |
| 10:7.196 | לו | , |
| 10:7.197 | וכאשר |  |
| 10:7.198 | ראה |  |
| 10:7.199 | שאמונתו |  |
| 10:7.200 | בזה |  |
| 10:7.201 | לא |  |
| 10:7.202 | תכלול |  |
| 10:7.203 | האכילה |  |
| 10:7.204 | והשתיה |  |
| 10:7.205 | במלאכים |  |
| 10:7.206 | קיים |  |
| 10:7.207 | בדעתו |  |
| 10:7.208 | ג״כ |  |
| 10:7.209 | אפשרותו |  |
| 10:7.210 | בהם | , |
| 10:7.211 | שוב |  |
| 10:7.212 | ראה |  |
| 10:7.213 | עוד |  |
| 10:7.214 | שכמו |  |
| 10:7.215 | שאלו |  |
| 10:7.216 | העניינים |  |
| 10:7.217 | אפשר |  |
| 10:7.218 | היותם |  |
| 10:7.219 | בהםמ |  |
| 10:7.220 | יחוייב |  |
| 10:7.221 | בהם |  |
| 10:7.222 | התאוה |  |
| 10:7.223 | והמשגל |  |
| 10:7.224 | וקיימם |  |
| 10:7.225 | בהם |  |
| 10:7.226 | ג״כ | , |
| 10:7.227 | ואמר |  |
| 10:7.228 | בענין |  |
| 10:7.229 | ויראו |  |
| 10:7.230 | בני |  |
| 10:7.231 | האלהים |  |
| 10:7.232 | את |  |
| 10:7.233 | בנות |  |
| 10:7.234 | האדם |  |
| 10:7.235 | כי |  |
| 10:7.236 | טבת |  |
| 10:7.237 | הנה | , |
| 10:7.238 | הם |  |
| 10:7.239 | המלאכים |  |
| 10:7.240 | הבאים |  |
| 10:7.241 | על |  |
| 10:7.242 | בנות |  |
| 10:7.243 | האדם |  |
| 10:7.244 | בנפלם |  |
| 10:7.245 | מן |  |
| 10:7.246 | השמים |  |
| 10:7.247 | והיו |  |
| 10:7.248 | ככל |  |
| 10:7.249 | האדם | , |
| 10:7.250 | והביא |  |
| 10:7.251 | דמיון |  |
| 10:7.252 | ראיה |  |
| 10:7.253 | לדבריו |  |
| 10:7.254 | ממה |  |
| 10:7.255 | שנא | ' |
| 10:7.256 | אני |  |
| 10:7.257 | אמרתי |  |
| 10:7.258 | אלהים |  |
| 10:7.259 | אתם |  |
| 10:7.260 | ובני |  |
| 10:7.261 | עליון |  |
| 10:7.262 | כלכם |  |
| 10:7.263 | אכן |  |
| 10:7.264 | כאדם |  |
| 10:7.265 | תמותון |  |
| 10:7.266 | וכאחד |  |
| 10:7.267 | השרים |  |
| 10:7.268 | תפלו | ! |
| 10:7.269 | ולא |  |
| 10:7.270 | השאיר |  |
| 10:7.271 | מדה |  |
| 10:7.272 | פחותה |  |
| 10:7.273 | ומגונה |  |
| 10:7.274 | עד |  |
| 10:7.275 | שיחסה |  |
| 10:7.276 | להם | ! |
| 10:7.277 | כל |  |
| 10:7.278 | זה |  |
| 10:7.279 | הולך |  |
| 10:7.280 | על |  |
| 10:7.281 | פי |  |
| 10:7.282 | שטתו | , |
| 10:7.283 | והרכיב |  |
| 10:7.284 | כל |  |
| 10:7.285 | אלה |  |
| 10:7.286 | הדעות |  |
| 10:7.287 | הזרות |  |
| 10:7.288 | והנפסדות |  |
| 10:7.289 | ממעוט |  |
| 10:7.290 | הבנתו |  |
| 10:7.291 | בפירוש |  |
| 10:7.292 | הדברים |  |
| 10:7.293 | בלה״ק | , |
| 10:7.294 | והנני |  |
| 10:7.295 | מבאר |  |
| 10:7.296 | טעותו |  |
| 10:7.297 | בכל |  |
| 10:7.298 | מה |  |
| 10:7.299 | שהניח |  |
| 10:7.300 | בע״ה | , |
| 10:7.301 | דע |  |
| 10:7.302 | יעזרך |  |
| 10:7.303 | השם | , |
| 10:7.304 | שמה |  |
| 10:7.305 | שהביא |  |
| 10:7.306 | לזה |  |
| 10:7.307 | המתחיל |  |
| 10:7.308 | המגיס |  |
| 10:7.309 | דעתו |  |
| 10:7.310 | הוא |  |
| 10:7.311 | וכיוצא |  |
| 10:7.312 | בו |  |
| 10:7.313 | לומר |  |
| 10:7.314 | שהשטן |  |
| 10:7.315 | הוא |  |
| 10:7.316 | מלאך | , |
| 10:7.317 | הוא |  |
| 10:7.318 | אמרו |  |
| 10:7.319 | ית | ' |
| 10:7.320 | לשטן |  |
| 10:7.321 | הנה |  |
| 10:7.322 | כל |  |
| 10:7.323 | אשר |  |
| 10:7.324 | לו |  |
| 10:7.325 | בידיך |  |
| 10:7.326 | רק |  |
| 10:7.327 | אליו |  |
| 10:7.328 | אל |  |
| 10:7.329 | תשלח |  |
| 10:7.330 | ידך | , |
| 10:7.331 | וגם |  |
| 10:7.332 | אמרו |  |
| 10:7.333 | אך |  |
| 10:7.334 | את |  |
| 10:7.335 | נפשו |  |
| 10:7.336 | שמור | , |
| 10:7.337 | וגם |  |
| 10:7.338 | אמרו |  |
| 10:7.339 | ויך |  |
| 10:7.340 | את |  |
| 10:7.341 | איוב |  |
| 10:7.342 | בשחין |  |
| 10:7.343 | רע | , |
| 10:7.344 | וחשב |  |
| 10:7.345 | שכל |  |
| 10:7.346 | זה |  |
| 10:7.347 | מיוחס |  |
| 10:7.348 | אל |  |
| 10:7.349 | השטן | , |
| 10:7.350 | והנני |  |
| 10:7.351 | מבאר |  |
| 10:7.352 | ביאור |  |
| 10:7.353 | מספיק |  |
| 10:7.354 | שבו |  |
| 10:7.355 | יסורו |  |
| 10:7.356 | אלו |  |
| 10:7.357 | הספקות |  |
| 10:7.358 | הג | ' |
| 10:7.359 | אשר |  |
| 10:7.360 | דמה |  |
| 10:7.361 | ואומר |  |
| 10:7.362 | כי |  |
| 10:7.363 | הרצון |  |
| 10:7.364 | והחפץ |  |
| 10:7.365 | בלשון |  |
| 10:7.366 | בני |  |
| 10:7.367 | ישראל |  |
| 10:7.368 | יקרא |  |
| 10:7.369 | יד | , |
| 10:7.370 | כמו |  |
| 10:7.371 | שאמר |  |
| 10:7.372 | דוד |  |
| 10:7.373 | לאשה |  |
| 10:7.374 | הקובלת |  |
| 10:7.375 | לפניו |  |
| 10:7.376 | היד |  |
| 10:7.377 | יואב |  |
| 10:7.378 | איתך |  |
| 10:7.379 | בכל |  |
| 10:7.380 | זאת |  |
| 10:7.381 | ענינו |  |
| 10:7.382 | חפץ |  |
| 10:7.383 | ורצון |  |
| 10:7.384 | יואב |  |
| 10:7.385 | וכו | ' |
| 10:7.386 | ועל |  |
| 10:7.387 | פי |  |
| 10:7.388 | זה |  |
| 10:7.389 | יהיה |  |
| 10:7.390 | אמרו |  |
| 10:7.391 | כאן |  |
| 10:7.392 | הנה |  |
| 10:7.393 | כל |  |
| 10:7.394 | אשר |  |
| 10:7.395 | לו |  |
| 10:7.396 | בידיך |  |
| 10:7.397 | בחפצך |  |
| 10:7.398 | ורצונך |  |
| 10:7.399 | וכו | ' |
| 10:7.400 | ובסור |  |
| 10:7.401 | ג | ' |
| 10:7.402 | ספיקות |  |
| 10:7.403 | הללו |  |
| 10:7.404 | אומר |  |
| 10:7.405 | שטעות |  |
| 10:7.406 | זה |  |
| 10:7.407 | המפרש |  |
| 10:7.408 | בענין |  |
| 10:7.409 | ויראו |  |
| 10:7.410 | בני |  |
| 10:7.411 | האלהים |  |
| 10:7.412 | את |  |
| 10:7.413 | בנות |  |
| 10:7.414 | האדם |  |
| 10:7.415 | יותר |  |
| 10:7.416 | זר |  |
| 10:7.417 | ממה |  |
| 10:7.418 | שאמר |  |
| 10:7.419 | בענין |  |
| 10:7.420 | איוב | , |
| 10:7.421 | והוא |  |
| 10:7.422 | שלא |  |
| 10:7.423 | ידע |  |
| 10:7.424 | כלל |  |
| 10:7.425 | ממה |  |
| 10:7.426 | שמשתמשין |  |
| 10:7.427 | ישראל |  |
| 10:7.428 | בלשון |  |
| 10:7.429 | בענין |  |
| 10:7.430 | עיני |  |
| 10:7.431 | העדה |  |
| 10:7.432 | וההמוּן | , |
| 10:7.433 | שדרכם |  |
| 10:7.434 | לומר |  |
| 10:7.435 | מעשה |  |
| 10:7.436 | הגדולים | , |
| 10:7.437 | כך |  |
| 10:7.438 | ומעשה |  |
| 10:7.439 | המון |  |
| 10:7.440 | העם |  |
| 10:7.441 | כך | , |
| 10:7.442 | וכולם |  |
| 10:7.443 | אנשים |  |
| 10:7.444 | מבני |  |
| 10:7.445 | אדם | , |
| 10:7.446 | ואמנם |  |
| 10:7.447 | הבדילם |  |
| 10:7.448 | לגדוללתם |  |
| 10:7.449 | ומעלתם | , |
| 10:7.450 | ועל |  |
| 10:7.451 | פי |  |
| 10:7.452 | זה |  |
| 10:7.453 | יהיה |  |
| 10:7.454 | אמרו |  |
| 10:7.455 | ויראו |  |
| 10:7.456 | בני |  |
| 10:7.457 | האלהים |  |
| 10:7.458 | אנשים |  |
| 10:7.459 | ג״כ | , |
| 10:7.460 | אבל |  |
| 10:7.461 | הבדילם |  |
| 10:7.462 | מן |  |
| 10:7.463 | ההמון |  |
| 10:7.464 | לחשיבותם |  |
| 10:7.465 | ומעלתם | , |
| 10:7.466 | וזה |  |
| 10:7.467 | ביאור |  |
| 10:7.468 | יפה |  |
| 10:7.469 | לשלול |  |
| 10:7.470 | מהקרובים |  |
| 10:7.471 | לשו״ת |  |
| 10:7.472 | הפחיתות |  |
| 10:7.473 | והזרות | , |
| 10:7.474 | ומה |  |
| 10:7.475 | שיחזק |  |
| 10:7.476 | זה |  |
| 10:7.477 | הפירוש |  |
| 10:7.478 | אמרו |  |
| 10:7.479 | אחר |  |
| 10:7.480 | זה |  |
| 10:7.481 | וירא |  |
| 10:7.482 | ה | ' |
| 10:7.483 | כי |  |
| 10:7.484 | רבה |  |
| 10:7.485 | רעת |  |
| 10:7.486 | האדם | , |
| 10:7.487 | ולא |  |
| 10:7.488 | גינה |  |
| 10:7.489 | המלאכים | . |
| 10:7.490 | וגם |  |
| 10:7.491 | אמרו |  |
| 10:7.492 | בסמוך |  |
| 10:7.493 | אמחה |  |
| 10:7.494 | את |  |
| 10:7.495 | האדם |  |
| 10:7.496 | ולא |  |
| 10:7.497 | ענש |  |
| 10:7.498 | המלאכים |  |
| 10:7.499 | עמהם |  |
| 10:7.500 | ע״כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

