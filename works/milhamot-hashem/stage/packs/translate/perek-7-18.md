# Translation pack: milhamot-hashem perek-7-18 (Milchamot ha-Shem 7:23–7:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-18`.

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
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
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
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **תוה"ק** → "the holy Torah".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ידידי** → "my friend" (second person kept).
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
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **כדנה / אלהיא / יאבדו** (Jer 10:11) → "Thus shall you say to them: the gods
  that did not make the heavens and the earth shall perish from the earth and
  from under these heavens."
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


## Milchamot ha-Shem 7:23

Hebrew (Wikisource): צה) וכל אלה השגיאות נמשכו ונשתרבבו בישראל כשיני עוג מלך הבשן מהאנשים הפתאים אוהבי הנפלאות והנמנעות שנפתו למלך הנכרי שרימא אותם והלכו אחר דעותיו לפרש פסוקי התורה על פי רעיונותיהם העולים על רוחם שלא על דרך רז"ל בעלי התלמוד שהם בעלי הקבלה האמתית. וחדשו מנהגים ופסקי הלכות על פי הקבלה הנכרית החדשה, ונדפסו בסדורי התפלות. עד שנעשית כתורה חדשה מתנגדת לתוה"ק שבע"פ. המבוארת בתלמוד ובפוסקים הראשונים בענין האלהות ויחוד השם ית'. אמנם כאשר נשים לב לעיין יפה בדברי רז"ל ובדברי הזהר נוכל להוכיח בצדק. כי לא מפי רשב"י יצאו הדברים הנזכרים בזהר, שהרי בבראשית רבה פר' כ"ו אמרו, וז"ל ויראו בני האלהים. רשב"י קרא להו בנו דייניא. רשב"י הוה מקלל לכל מאן דקרי להו בני אלהיא תני רשב"י כל פרצה שאינה מן הגדולים אינה פרצה. כומרים גנבו אלהיא מאן מומי ביה או מאן מקריב ליה. ולמה קורא אותן בני האלהים ר' חנינה ור' שמעון בן לקיש תרויחון אמרין שהרבו ימים בלא צער ובלא יסורין וכו' ע"כ. ראה ידידי כמה הוא מרחיק רשב"י ושונא ומקלל את האומר שהם מלאכים? ובזהר מחבב סברא זו וז"ל (בראשית כ"ג) דבר אחר נעשה אדם בצלמינו כדמותינו, הא אוקמוה חבריא על מלאכי השרת דאינון אמרי להאי קרא, א"ל בתר דהוו ידעין מה דהוה, ומה דעתיד להוי ואינון הוו יודעין דעתיד למחטי אמאי בעו למעבד ליה, ולא עוד אלא דעז"א ועזא"ל הוו מקטרגי עליה בזמנא דאמר שכינתא לקב"ה נעשה אדם, אמרו מה אדם ותדעהו, מה את בעי למברי אדם. ותדעהו דעתיד למחטי באתתא דיליה דאיהי חשך, דאור איהו דכורא, וחשך נוקבא שמאלא דבריאה, בההוא זמנא אמר שכינתא לון, בהאי דאתון מקטרגין אתון עתידין למנפל. כדכתיב ויראו בני האלהים את בנות האדם כי טבת הנה. חשקו בהון וטעו בהון. ואפיל להון שכינתא מקדושה דילהון. אמרו חבריא רבי רבי אי הכי עזא ועזאל לא הוו משקרי במלוליהו. דודאי בנוקבא עתיד אדם למחטי. אמר להו הכי אמרה שכינתא. אתון דאזדמנתן לקטרגא קדמי יתיר מחילי מרומא. אי אתון חזיתון שפירין מאדם בעובדיכון יאות לכון לקטרגא עליה, אבל איהו עתיד למחטי באתתא חדא. אתון בנשי סגיאי חבתון יתיר מבני נשא דכתיב ויראו בני האלהים את בנות האדם כי טבת הנה וכו'. גם בדף נ"ה חזר לשנותה בשם ר' יצחק. ובדף נ"ח חזר ושלשה בשם ר"י ור"ח. הרי שמחבר הזהר מחבב סברא זו ולפיכך אמרה ושנאה ושילשה לרוב חבתה. והרי בבראשית רבה מבורא שרשב"י שונא סברא זו עד שהיה מקלל את האומר שהם בני אלהיא. אלא ודאי לא מפי רשב"י יצאו דברים אלו האמורים בזהר. ובשקר יחסוהו לרשב"י.

| id | word | punctuation after |
|---|---|---|
| 7:23.1 | צה | ) |
| 7:23.2 | וכל |  |
| 7:23.3 | אלה |  |
| 7:23.4 | השגיאות |  |
| 7:23.5 | נמשכו |  |
| 7:23.6 | ונשתרבבו |  |
| 7:23.7 | בישראל |  |
| 7:23.8 | כשיני |  |
| 7:23.9 | עוג |  |
| 7:23.10 | מלך |  |
| 7:23.11 | הבשן |  |
| 7:23.12 | מהאנשים |  |
| 7:23.13 | הפתאים |  |
| 7:23.14 | אוהבי |  |
| 7:23.15 | הנפלאות |  |
| 7:23.16 | והנמנעות |  |
| 7:23.17 | שנפתו |  |
| 7:23.18 | למלך |  |
| 7:23.19 | הנכרי |  |
| 7:23.20 | שרימא |  |
| 7:23.21 | אותם |  |
| 7:23.22 | והלכו |  |
| 7:23.23 | אחר |  |
| 7:23.24 | דעותיו |  |
| 7:23.25 | לפרש |  |
| 7:23.26 | פסוקי |  |
| 7:23.27 | התורה |  |
| 7:23.28 | על |  |
| 7:23.29 | פי |  |
| 7:23.30 | רעיונותיהם |  |
| 7:23.31 | העולים |  |
| 7:23.32 | על |  |
| 7:23.33 | רוחם |  |
| 7:23.34 | שלא |  |
| 7:23.35 | על |  |
| 7:23.36 | דרך |  |
| 7:23.37 | רז"ל |  |
| 7:23.38 | בעלי |  |
| 7:23.39 | התלמוד |  |
| 7:23.40 | שהם |  |
| 7:23.41 | בעלי |  |
| 7:23.42 | הקבלה |  |
| 7:23.43 | האמתית | . |
| 7:23.44 | וחדשו |  |
| 7:23.45 | מנהגים |  |
| 7:23.46 | ופסקי |  |
| 7:23.47 | הלכות |  |
| 7:23.48 | על |  |
| 7:23.49 | פי |  |
| 7:23.50 | הקבלה |  |
| 7:23.51 | הנכרית |  |
| 7:23.52 | החדשה | , |
| 7:23.53 | ונדפסו |  |
| 7:23.54 | בסדורי |  |
| 7:23.55 | התפלות | . |
| 7:23.56 | עד |  |
| 7:23.57 | שנעשית |  |
| 7:23.58 | כתורה |  |
| 7:23.59 | חדשה |  |
| 7:23.60 | מתנגדת |  |
| 7:23.61 | לתוה"ק |  |
| 7:23.62 | שבע"פ | . |
| 7:23.63 | המבוארת |  |
| 7:23.64 | בתלמוד |  |
| 7:23.65 | ובפוסקים |  |
| 7:23.66 | הראשונים |  |
| 7:23.67 | בענין |  |
| 7:23.68 | האלהות |  |
| 7:23.69 | ויחוד |  |
| 7:23.70 | השם |  |
| 7:23.71 | ית | '. |
| 7:23.72 | אמנם |  |
| 7:23.73 | כאשר |  |
| 7:23.74 | נשים |  |
| 7:23.75 | לב |  |
| 7:23.76 | לעיין |  |
| 7:23.77 | יפה |  |
| 7:23.78 | בדברי |  |
| 7:23.79 | רז"ל |  |
| 7:23.80 | ובדברי |  |
| 7:23.81 | הזהר |  |
| 7:23.82 | נוכל |  |
| 7:23.83 | להוכיח |  |
| 7:23.84 | בצדק | . |
| 7:23.85 | כי |  |
| 7:23.86 | לא |  |
| 7:23.87 | מפי |  |
| 7:23.88 | רשב"י |  |
| 7:23.89 | יצאו |  |
| 7:23.90 | הדברים |  |
| 7:23.91 | הנזכרים |  |
| 7:23.92 | בזהר | , |
| 7:23.93 | שהרי |  |
| 7:23.94 | בבראשית |  |
| 7:23.95 | רבה |  |
| 7:23.96 | פר | ' |
| 7:23.97 | כ"ו |  |
| 7:23.98 | אמרו | , |
| 7:23.99 | וז"ל |  |
| 7:23.100 | ויראו |  |
| 7:23.101 | בני |  |
| 7:23.102 | האלהים | . |
| 7:23.103 | רשב"י |  |
| 7:23.104 | קרא |  |
| 7:23.105 | להו |  |
| 7:23.106 | בנו |  |
| 7:23.107 | דייניא | . |
| 7:23.108 | רשב"י |  |
| 7:23.109 | הוה |  |
| 7:23.110 | מקלל |  |
| 7:23.111 | לכל |  |
| 7:23.112 | מאן |  |
| 7:23.113 | דקרי |  |
| 7:23.114 | להו |  |
| 7:23.115 | בני |  |
| 7:23.116 | אלהיא |  |
| 7:23.117 | תני |  |
| 7:23.118 | רשב"י |  |
| 7:23.119 | כל |  |
| 7:23.120 | פרצה |  |
| 7:23.121 | שאינה |  |
| 7:23.122 | מן |  |
| 7:23.123 | הגדולים |  |
| 7:23.124 | אינה |  |
| 7:23.125 | פרצה | . |
| 7:23.126 | כומרים |  |
| 7:23.127 | גנבו |  |
| 7:23.128 | אלהיא |  |
| 7:23.129 | מאן |  |
| 7:23.130 | מומי |  |
| 7:23.131 | ביה |  |
| 7:23.132 | או |  |
| 7:23.133 | מאן |  |
| 7:23.134 | מקריב |  |
| 7:23.135 | ליה | . |
| 7:23.136 | ולמה |  |
| 7:23.137 | קורא |  |
| 7:23.138 | אותן |  |
| 7:23.139 | בני |  |
| 7:23.140 | האלהים |  |
| 7:23.141 | ר | ' |
| 7:23.142 | חנינה |  |
| 7:23.143 | ור | ' |
| 7:23.144 | שמעון |  |
| 7:23.145 | בן |  |
| 7:23.146 | לקיש |  |
| 7:23.147 | תרויחון |  |
| 7:23.148 | אמרין |  |
| 7:23.149 | שהרבו |  |
| 7:23.150 | ימים |  |
| 7:23.151 | בלא |  |
| 7:23.152 | צער |  |
| 7:23.153 | ובלא |  |
| 7:23.154 | יסורין |  |
| 7:23.155 | וכו | ' |
| 7:23.156 | ע"כ | . |
| 7:23.157 | ראה |  |
| 7:23.158 | ידידי |  |
| 7:23.159 | כמה |  |
| 7:23.160 | הוא |  |
| 7:23.161 | מרחיק |  |
| 7:23.162 | רשב"י |  |
| 7:23.163 | ושונא |  |
| 7:23.164 | ומקלל |  |
| 7:23.165 | את |  |
| 7:23.166 | האומר |  |
| 7:23.167 | שהם |  |
| 7:23.168 | מלאכים | ? |
| 7:23.169 | ובזהר |  |
| 7:23.170 | מחבב |  |
| 7:23.171 | סברא |  |
| 7:23.172 | זו |  |
| 7:23.173 | וז"ל |  |
| 7:23.174 | בראשית |  |
| 7:23.175 | כ"ג | ) |
| 7:23.176 | דבר |  |
| 7:23.177 | אחר |  |
| 7:23.178 | נעשה |  |
| 7:23.179 | אדם |  |
| 7:23.180 | בצלמינו |  |
| 7:23.181 | כדמותינו | , |
| 7:23.182 | הא |  |
| 7:23.183 | אוקמוה |  |
| 7:23.184 | חבריא |  |
| 7:23.185 | על |  |
| 7:23.186 | מלאכי |  |
| 7:23.187 | השרת |  |
| 7:23.188 | דאינון |  |
| 7:23.189 | אמרי |  |
| 7:23.190 | להאי |  |
| 7:23.191 | קרא | , |
| 7:23.192 | א"ל |  |
| 7:23.193 | בתר |  |
| 7:23.194 | דהוו |  |
| 7:23.195 | ידעין |  |
| 7:23.196 | מה |  |
| 7:23.197 | דהוה | , |
| 7:23.198 | ומה |  |
| 7:23.199 | דעתיד |  |
| 7:23.200 | להוי |  |
| 7:23.201 | ואינון |  |
| 7:23.202 | הוו |  |
| 7:23.203 | יודעין |  |
| 7:23.204 | דעתיד |  |
| 7:23.205 | למחטי |  |
| 7:23.206 | אמאי |  |
| 7:23.207 | בעו |  |
| 7:23.208 | למעבד |  |
| 7:23.209 | ליה | , |
| 7:23.210 | ולא |  |
| 7:23.211 | עוד |  |
| 7:23.212 | אלא |  |
| 7:23.213 | דעז"א |  |
| 7:23.214 | ועזא"ל |  |
| 7:23.215 | הוו |  |
| 7:23.216 | מקטרגי |  |
| 7:23.217 | עליה |  |
| 7:23.218 | בזמנא |  |
| 7:23.219 | דאמר |  |
| 7:23.220 | שכינתא |  |
| 7:23.221 | לקב"ה |  |
| 7:23.222 | נעשה |  |
| 7:23.223 | אדם | , |
| 7:23.224 | אמרו |  |
| 7:23.225 | מה |  |
| 7:23.226 | אדם |  |
| 7:23.227 | ותדעהו | , |
| 7:23.228 | מה |  |
| 7:23.229 | את |  |
| 7:23.230 | בעי |  |
| 7:23.231 | למברי |  |
| 7:23.232 | אדם | . |
| 7:23.233 | ותדעהו |  |
| 7:23.234 | דעתיד |  |
| 7:23.235 | למחטי |  |
| 7:23.236 | באתתא |  |
| 7:23.237 | דיליה |  |
| 7:23.238 | דאיהי |  |
| 7:23.239 | חשך | , |
| 7:23.240 | דאור |  |
| 7:23.241 | איהו |  |
| 7:23.242 | דכורא | , |
| 7:23.243 | וחשך |  |
| 7:23.244 | נוקבא |  |
| 7:23.245 | שמאלא |  |
| 7:23.246 | דבריאה | , |
| 7:23.247 | בההוא |  |
| 7:23.248 | זמנא |  |
| 7:23.249 | אמר |  |
| 7:23.250 | שכינתא |  |
| 7:23.251 | לון | , |
| 7:23.252 | בהאי |  |
| 7:23.253 | דאתון |  |
| 7:23.254 | מקטרגין |  |
| 7:23.255 | אתון |  |
| 7:23.256 | עתידין |  |
| 7:23.257 | למנפל | . |
| 7:23.258 | כדכתיב |  |
| 7:23.259 | ויראו |  |
| 7:23.260 | בני |  |
| 7:23.261 | האלהים |  |
| 7:23.262 | את |  |
| 7:23.263 | בנות |  |
| 7:23.264 | האדם |  |
| 7:23.265 | כי |  |
| 7:23.266 | טבת |  |
| 7:23.267 | הנה | . |
| 7:23.268 | חשקו |  |
| 7:23.269 | בהון |  |
| 7:23.270 | וטעו |  |
| 7:23.271 | בהון | . |
| 7:23.272 | ואפיל |  |
| 7:23.273 | להון |  |
| 7:23.274 | שכינתא |  |
| 7:23.275 | מקדושה |  |
| 7:23.276 | דילהון | . |
| 7:23.277 | אמרו |  |
| 7:23.278 | חבריא |  |
| 7:23.279 | רבי |  |
| 7:23.280 | רבי |  |
| 7:23.281 | אי |  |
| 7:23.282 | הכי |  |
| 7:23.283 | עזא |  |
| 7:23.284 | ועזאל |  |
| 7:23.285 | לא |  |
| 7:23.286 | הוו |  |
| 7:23.287 | משקרי |  |
| 7:23.288 | במלוליהו | . |
| 7:23.289 | דודאי |  |
| 7:23.290 | בנוקבא |  |
| 7:23.291 | עתיד |  |
| 7:23.292 | אדם |  |
| 7:23.293 | למחטי | . |
| 7:23.294 | אמר |  |
| 7:23.295 | להו |  |
| 7:23.296 | הכי |  |
| 7:23.297 | אמרה |  |
| 7:23.298 | שכינתא | . |
| 7:23.299 | אתון |  |
| 7:23.300 | דאזדמנתן |  |
| 7:23.301 | לקטרגא |  |
| 7:23.302 | קדמי |  |
| 7:23.303 | יתיר |  |
| 7:23.304 | מחילי |  |
| 7:23.305 | מרומא | . |
| 7:23.306 | אי |  |
| 7:23.307 | אתון |  |
| 7:23.308 | חזיתון |  |
| 7:23.309 | שפירין |  |
| 7:23.310 | מאדם |  |
| 7:23.311 | בעובדיכון |  |
| 7:23.312 | יאות |  |
| 7:23.313 | לכון |  |
| 7:23.314 | לקטרגא |  |
| 7:23.315 | עליה | , |
| 7:23.316 | אבל |  |
| 7:23.317 | איהו |  |
| 7:23.318 | עתיד |  |
| 7:23.319 | למחטי |  |
| 7:23.320 | באתתא |  |
| 7:23.321 | חדא | . |
| 7:23.322 | אתון |  |
| 7:23.323 | בנשי |  |
| 7:23.324 | סגיאי |  |
| 7:23.325 | חבתון |  |
| 7:23.326 | יתיר |  |
| 7:23.327 | מבני |  |
| 7:23.328 | נשא |  |
| 7:23.329 | דכתיב |  |
| 7:23.330 | ויראו |  |
| 7:23.331 | בני |  |
| 7:23.332 | האלהים |  |
| 7:23.333 | את |  |
| 7:23.334 | בנות |  |
| 7:23.335 | האדם |  |
| 7:23.336 | כי |  |
| 7:23.337 | טבת |  |
| 7:23.338 | הנה |  |
| 7:23.339 | וכו | '. |
| 7:23.340 | גם |  |
| 7:23.341 | בדף |  |
| 7:23.342 | נ"ה |  |
| 7:23.343 | חזר |  |
| 7:23.344 | לשנותה |  |
| 7:23.345 | בשם |  |
| 7:23.346 | ר | ' |
| 7:23.347 | יצחק | . |
| 7:23.348 | ובדף |  |
| 7:23.349 | נ"ח |  |
| 7:23.350 | חזר |  |
| 7:23.351 | ושלשה |  |
| 7:23.352 | בשם |  |
| 7:23.353 | ר"י |  |
| 7:23.354 | ור"ח | . |
| 7:23.355 | הרי |  |
| 7:23.356 | שמחבר |  |
| 7:23.357 | הזהר |  |
| 7:23.358 | מחבב |  |
| 7:23.359 | סברא |  |
| 7:23.360 | זו |  |
| 7:23.361 | ולפיכך |  |
| 7:23.362 | אמרה |  |
| 7:23.363 | ושנאה |  |
| 7:23.364 | ושילשה |  |
| 7:23.365 | לרוב |  |
| 7:23.366 | חבתה | . |
| 7:23.367 | והרי |  |
| 7:23.368 | בבראשית |  |
| 7:23.369 | רבה |  |
| 7:23.370 | מבורא |  |
| 7:23.371 | שרשב"י |  |
| 7:23.372 | שונא |  |
| 7:23.373 | סברא |  |
| 7:23.374 | זו |  |
| 7:23.375 | עד |  |
| 7:23.376 | שהיה |  |
| 7:23.377 | מקלל |  |
| 7:23.378 | את |  |
| 7:23.379 | האומר |  |
| 7:23.380 | שהם |  |
| 7:23.381 | בני |  |
| 7:23.382 | אלהיא | . |
| 7:23.383 | אלא |  |
| 7:23.384 | ודאי |  |
| 7:23.385 | לא |  |
| 7:23.386 | מפי |  |
| 7:23.387 | רשב"י |  |
| 7:23.388 | יצאו |  |
| 7:23.389 | דברים |  |
| 7:23.390 | אלו |  |
| 7:23.391 | האמורים |  |
| 7:23.392 | בזהר | . |
| 7:23.393 | ובשקר |  |
| 7:23.394 | יחסוהו |  |
| 7:23.395 | לרשב"י | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

