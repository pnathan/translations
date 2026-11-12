# Translation pack: milhamot-hashem perek-9-14 (Milchamot ha-Shem 9:17–9:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-14`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בא"י** → "in the Land of Israel".

## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
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
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 9:17

Hebrew (Wikisource): קלה) רב לכם נשיאי ישראל מכל דברי מינות האמורים בזהר בכמה מקומות, ובפרט בפרשת בהר דף ק"כ שהאלהים העליון שאם ראוי לעבדו (לפי דבריה) ברא והמציא את האלהים התחתונים הרבים, והוא ממלא אותם בחללם וחוצה להם, והוא המאחד ומקשר את נקיבותיהם עם הזכרים לקרא להם בשם קב"ה ושכינתיה, ואותם נעבוד, ונירא את נקיבותיהם, ודוקא לפרצוף אחרון שבהם הנקרא קצר אפים (ז"א) עם מלכת השמים המזדווגת עמו, משום דלא שייך עבודה ותפלה רק לפרצופים שכבר נתגשמו ונתעבו קצת כמבואר במה שקדם מדברי כסא אליהו יושר לבב וספר הברית ועץ חיים להרח"ו שער טנת"ה סוף פרק ב' ובשער הנקודים פ"א שכל מה שהאו יורד למטה הוא מתעבה יותר וניכר ונרגש ומתגלה יותר, אבל באין סוף הוא הסבה הראשונה ואפי' בפרצופים ראשונים (שהמציאו המקובלים החדשים לאלהינו) כאדם קדמאה ואדם קדמון ועתיק ואריך שעדיין לא נתעבו ונתגשמו כל כך לא שייכא עבודה וקריאה, כי אין המה, ולא יאות לקרוא להם בשם אלהים כמפורש יוצא מספר עץ חיים להרח"ו שער הנקודים פ"ז ובהג"ה שם. שזהו האלהים העליון עשה את האדם ישר (והמה בקשו חשבונות רבים) כי האדם קדמון היה בו בחינת יושר וזהו עשה את האדם ישר (הגהה א"מ ןקרא המאציל אלהים בדרך שאלה. כי הרי אפי' באריך אנפין אין שם אלהים) ע"כ. הרי לך בהדיא דפרצופים העליונים לא שייכא בהו אלהות וקריאה כלל, כי לא אלהים המה ורק להעלולים התחתונים נקרא ונעבוד, והמה אלהינו (לפי דעתם) כמו שכתב הרב יושר לבב וספר הברית וכס"א ונח"י ומצרף האמונה! ולפי דבריהם שאומרים דלית רשותא לבר נש למיעל קמי מריה אלא בידיעת אמונה זו, איך אמרו רז"ל רצה הקב"ה לזכות את ישראל לפיכך הרבה להם תורה ומצות שנ' ה' חפץ למען צדקו יגדיל תורה ויאדיר! ולמה זה לא רצה הקב"ה לזכות אותנו ולהודענו על ידי נביאנו הרבים שעמדו להם לישראל ממשה רבינו ע"ה עד שהנביא האחרון מלאכי שאמר זכרו תורת משה עבדי וכו'. את אמונה זו שהעבודה היא להעלולים התחתונים. ושהעובד לאלהים עליון לא יענהו ואדרבא ענוש יענש ומחה ה' את שמו כמו שכתב ספר הברית?

| id | word | punctuation after |
|---|---|---|
| 9:17.1 | קלה | ) |
| 9:17.2 | רב |  |
| 9:17.3 | לכם |  |
| 9:17.4 | נשיאי |  |
| 9:17.5 | ישראל |  |
| 9:17.6 | מכל |  |
| 9:17.7 | דברי |  |
| 9:17.8 | מינות |  |
| 9:17.9 | האמורים |  |
| 9:17.10 | בזהר |  |
| 9:17.11 | בכמה |  |
| 9:17.12 | מקומות | , |
| 9:17.13 | ובפרט |  |
| 9:17.14 | בפרשת |  |
| 9:17.15 | בהר |  |
| 9:17.16 | דף |  |
| 9:17.17 | ק"כ |  |
| 9:17.18 | שהאלהים |  |
| 9:17.19 | העליון |  |
| 9:17.20 | שאם |  |
| 9:17.21 | ראוי |  |
| 9:17.22 | לעבדו |  |
| 9:17.23 | לפי |  |
| 9:17.24 | דבריה | ) |
| 9:17.25 | ברא |  |
| 9:17.26 | והמציא |  |
| 9:17.27 | את |  |
| 9:17.28 | האלהים |  |
| 9:17.29 | התחתונים |  |
| 9:17.30 | הרבים | , |
| 9:17.31 | והוא |  |
| 9:17.32 | ממלא |  |
| 9:17.33 | אותם |  |
| 9:17.34 | בחללם |  |
| 9:17.35 | וחוצה |  |
| 9:17.36 | להם | , |
| 9:17.37 | והוא |  |
| 9:17.38 | המאחד |  |
| 9:17.39 | ומקשר |  |
| 9:17.40 | את |  |
| 9:17.41 | נקיבותיהם |  |
| 9:17.42 | עם |  |
| 9:17.43 | הזכרים |  |
| 9:17.44 | לקרא |  |
| 9:17.45 | להם |  |
| 9:17.46 | בשם |  |
| 9:17.47 | קב"ה |  |
| 9:17.48 | ושכינתיה | , |
| 9:17.49 | ואותם |  |
| 9:17.50 | נעבוד | , |
| 9:17.51 | ונירא |  |
| 9:17.52 | את |  |
| 9:17.53 | נקיבותיהם | , |
| 9:17.54 | ודוקא |  |
| 9:17.55 | לפרצוף |  |
| 9:17.56 | אחרון |  |
| 9:17.57 | שבהם |  |
| 9:17.58 | הנקרא |  |
| 9:17.59 | קצר |  |
| 9:17.60 | אפים |  |
| 9:17.61 | ז"א | ) |
| 9:17.62 | עם |  |
| 9:17.63 | מלכת |  |
| 9:17.64 | השמים |  |
| 9:17.65 | המזדווגת |  |
| 9:17.66 | עמו | , |
| 9:17.67 | משום |  |
| 9:17.68 | דלא |  |
| 9:17.69 | שייך |  |
| 9:17.70 | עבודה |  |
| 9:17.71 | ותפלה |  |
| 9:17.72 | רק |  |
| 9:17.73 | לפרצופים |  |
| 9:17.74 | שכבר |  |
| 9:17.75 | נתגשמו |  |
| 9:17.76 | ונתעבו |  |
| 9:17.77 | קצת |  |
| 9:17.78 | כמבואר |  |
| 9:17.79 | במה |  |
| 9:17.80 | שקדם |  |
| 9:17.81 | מדברי |  |
| 9:17.82 | כסא |  |
| 9:17.83 | אליהו |  |
| 9:17.84 | יושר |  |
| 9:17.85 | לבב |  |
| 9:17.86 | וספר |  |
| 9:17.87 | הברית |  |
| 9:17.88 | ועץ |  |
| 9:17.89 | חיים |  |
| 9:17.90 | להרח"ו |  |
| 9:17.91 | שער |  |
| 9:17.92 | טנת"ה |  |
| 9:17.93 | סוף |  |
| 9:17.94 | פרק |  |
| 9:17.95 | ב | ' |
| 9:17.96 | ובשער |  |
| 9:17.97 | הנקודים |  |
| 9:17.98 | פ"א |  |
| 9:17.99 | שכל |  |
| 9:17.100 | מה |  |
| 9:17.101 | שהאו |  |
| 9:17.102 | יורד |  |
| 9:17.103 | למטה |  |
| 9:17.104 | הוא |  |
| 9:17.105 | מתעבה |  |
| 9:17.106 | יותר |  |
| 9:17.107 | וניכר |  |
| 9:17.108 | ונרגש |  |
| 9:17.109 | ומתגלה |  |
| 9:17.110 | יותר | , |
| 9:17.111 | אבל |  |
| 9:17.112 | באין |  |
| 9:17.113 | סוף |  |
| 9:17.114 | הוא |  |
| 9:17.115 | הסבה |  |
| 9:17.116 | הראשונה |  |
| 9:17.117 | ואפי | ' |
| 9:17.118 | בפרצופים |  |
| 9:17.119 | ראשונים |  |
| 9:17.120 | שהמציאו |  |
| 9:17.121 | המקובלים |  |
| 9:17.122 | החדשים |  |
| 9:17.123 | לאלהינו | ) |
| 9:17.124 | כאדם |  |
| 9:17.125 | קדמאה |  |
| 9:17.126 | ואדם |  |
| 9:17.127 | קדמון |  |
| 9:17.128 | ועתיק |  |
| 9:17.129 | ואריך |  |
| 9:17.130 | שעדיין |  |
| 9:17.131 | לא |  |
| 9:17.132 | נתעבו |  |
| 9:17.133 | ונתגשמו |  |
| 9:17.134 | כל |  |
| 9:17.135 | כך |  |
| 9:17.136 | לא |  |
| 9:17.137 | שייכא |  |
| 9:17.138 | עבודה |  |
| 9:17.139 | וקריאה | , |
| 9:17.140 | כי |  |
| 9:17.141 | אין |  |
| 9:17.142 | המה | , |
| 9:17.143 | ולא |  |
| 9:17.144 | יאות |  |
| 9:17.145 | לקרוא |  |
| 9:17.146 | להם |  |
| 9:17.147 | בשם |  |
| 9:17.148 | אלהים |  |
| 9:17.149 | כמפורש |  |
| 9:17.150 | יוצא |  |
| 9:17.151 | מספר |  |
| 9:17.152 | עץ |  |
| 9:17.153 | חיים |  |
| 9:17.154 | להרח"ו |  |
| 9:17.155 | שער |  |
| 9:17.156 | הנקודים |  |
| 9:17.157 | פ"ז |  |
| 9:17.158 | ובהג"ה |  |
| 9:17.159 | שם | . |
| 9:17.160 | שזהו |  |
| 9:17.161 | האלהים |  |
| 9:17.162 | העליון |  |
| 9:17.163 | עשה |  |
| 9:17.164 | את |  |
| 9:17.165 | האדם |  |
| 9:17.166 | ישר |  |
| 9:17.167 | והמה |  |
| 9:17.168 | בקשו |  |
| 9:17.169 | חשבונות |  |
| 9:17.170 | רבים | ) |
| 9:17.171 | כי |  |
| 9:17.172 | האדם |  |
| 9:17.173 | קדמון |  |
| 9:17.174 | היה |  |
| 9:17.175 | בו |  |
| 9:17.176 | בחינת |  |
| 9:17.177 | יושר |  |
| 9:17.178 | וזהו |  |
| 9:17.179 | עשה |  |
| 9:17.180 | את |  |
| 9:17.181 | האדם |  |
| 9:17.182 | ישר |  |
| 9:17.183 | הגהה |  |
| 9:17.184 | א"מ |  |
| 9:17.185 | ןקרא |  |
| 9:17.186 | המאציל |  |
| 9:17.187 | אלהים |  |
| 9:17.188 | בדרך |  |
| 9:17.189 | שאלה | . |
| 9:17.190 | כי |  |
| 9:17.191 | הרי |  |
| 9:17.192 | אפי | ' |
| 9:17.193 | באריך |  |
| 9:17.194 | אנפין |  |
| 9:17.195 | אין |  |
| 9:17.196 | שם |  |
| 9:17.197 | אלהים | ) |
| 9:17.198 | ע"כ | . |
| 9:17.199 | הרי |  |
| 9:17.200 | לך |  |
| 9:17.201 | בהדיא |  |
| 9:17.202 | דפרצופים |  |
| 9:17.203 | העליונים |  |
| 9:17.204 | לא |  |
| 9:17.205 | שייכא |  |
| 9:17.206 | בהו |  |
| 9:17.207 | אלהות |  |
| 9:17.208 | וקריאה |  |
| 9:17.209 | כלל | , |
| 9:17.210 | כי |  |
| 9:17.211 | לא |  |
| 9:17.212 | אלהים |  |
| 9:17.213 | המה |  |
| 9:17.214 | ורק |  |
| 9:17.215 | להעלולים |  |
| 9:17.216 | התחתונים |  |
| 9:17.217 | נקרא |  |
| 9:17.218 | ונעבוד | , |
| 9:17.219 | והמה |  |
| 9:17.220 | אלהינו |  |
| 9:17.221 | לפי |  |
| 9:17.222 | דעתם | ) |
| 9:17.223 | כמו |  |
| 9:17.224 | שכתב |  |
| 9:17.225 | הרב |  |
| 9:17.226 | יושר |  |
| 9:17.227 | לבב |  |
| 9:17.228 | וספר |  |
| 9:17.229 | הברית |  |
| 9:17.230 | וכס"א |  |
| 9:17.231 | ונח"י |  |
| 9:17.232 | ומצרף |  |
| 9:17.233 | האמונה | ! |
| 9:17.234 | ולפי |  |
| 9:17.235 | דבריהם |  |
| 9:17.236 | שאומרים |  |
| 9:17.237 | דלית |  |
| 9:17.238 | רשותא |  |
| 9:17.239 | לבר |  |
| 9:17.240 | נש |  |
| 9:17.241 | למיעל |  |
| 9:17.242 | קמי |  |
| 9:17.243 | מריה |  |
| 9:17.244 | אלא |  |
| 9:17.245 | בידיעת |  |
| 9:17.246 | אמונה |  |
| 9:17.247 | זו | , |
| 9:17.248 | איך |  |
| 9:17.249 | אמרו |  |
| 9:17.250 | רז"ל |  |
| 9:17.251 | רצה |  |
| 9:17.252 | הקב"ה |  |
| 9:17.253 | לזכות |  |
| 9:17.254 | את |  |
| 9:17.255 | ישראל |  |
| 9:17.256 | לפיכך |  |
| 9:17.257 | הרבה |  |
| 9:17.258 | להם |  |
| 9:17.259 | תורה |  |
| 9:17.260 | ומצות |  |
| 9:17.261 | שנ | ' |
| 9:17.262 | ה | ' |
| 9:17.263 | חפץ |  |
| 9:17.264 | למען |  |
| 9:17.265 | צדקו |  |
| 9:17.266 | יגדיל |  |
| 9:17.267 | תורה |  |
| 9:17.268 | ויאדיר | ! |
| 9:17.269 | ולמה |  |
| 9:17.270 | זה |  |
| 9:17.271 | לא |  |
| 9:17.272 | רצה |  |
| 9:17.273 | הקב"ה |  |
| 9:17.274 | לזכות |  |
| 9:17.275 | אותנו |  |
| 9:17.276 | ולהודענו |  |
| 9:17.277 | על |  |
| 9:17.278 | ידי |  |
| 9:17.279 | נביאנו |  |
| 9:17.280 | הרבים |  |
| 9:17.281 | שעמדו |  |
| 9:17.282 | להם |  |
| 9:17.283 | לישראל |  |
| 9:17.284 | ממשה |  |
| 9:17.285 | רבינו |  |
| 9:17.286 | ע"ה |  |
| 9:17.287 | עד |  |
| 9:17.288 | שהנביא |  |
| 9:17.289 | האחרון |  |
| 9:17.290 | מלאכי |  |
| 9:17.291 | שאמר |  |
| 9:17.292 | זכרו |  |
| 9:17.293 | תורת |  |
| 9:17.294 | משה |  |
| 9:17.295 | עבדי |  |
| 9:17.296 | וכו | '. |
| 9:17.297 | את |  |
| 9:17.298 | אמונה |  |
| 9:17.299 | זו |  |
| 9:17.300 | שהעבודה |  |
| 9:17.301 | היא |  |
| 9:17.302 | להעלולים |  |
| 9:17.303 | התחתונים | . |
| 9:17.304 | ושהעובד |  |
| 9:17.305 | לאלהים |  |
| 9:17.306 | עליון |  |
| 9:17.307 | לא |  |
| 9:17.308 | יענהו |  |
| 9:17.309 | ואדרבא |  |
| 9:17.310 | ענוש |  |
| 9:17.311 | יענש |  |
| 9:17.312 | ומחה |  |
| 9:17.313 | ה | ' |
| 9:17.314 | את |  |
| 9:17.315 | שמו |  |
| 9:17.316 | כמו |  |
| 9:17.317 | שכתב |  |
| 9:17.318 | ספר |  |
| 9:17.319 | הברית | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

