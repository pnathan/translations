# Translation pack: milhamot-hashem perek-9-09 (Milchamot ha-Shem 9:11–9:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-09`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **מרן** → "Maran" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- כופר **בעיקר** → "denies the fundamental principle".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- **תא חזי** → "come and see".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **ע"א / ע"ב** → "a / b".
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 9:11

Hebrew (Wikisource): קכח) עוד זאת רעה חולה, שכל הלומדים בזהר ושאר ספרי קבלה חדשה נקל בעיניהם לעולל עלילות ולשם תהלה במאורי ישראל מורי צדק המאירים עיני ישראל בכל מקצועות התורה אשר יעשה אותם האדם וחי בהם, ובפרט בעיקרי תורתינו איך נאמין ונקיים מצותיה! וביחוד העיקר והעמוד שהכל תלוי בו, והוא יחוד השי"ת לא שנאמר בפינו אחד, ובלבבינו רבים, כמו שאומרים המקובלים כי האלהים העליון דלא שייך ביה עבודה המקשר ומחבר ומזויג את האלהים הרבים התחתונים אשר אותם נעבוד להיותם אחד! אוי לעינים שכך רואות, ואוי לאזנים שכך שומעות, ודברים כאלה הן שאסור לשמען ואסור להרהר בהן אפילו בבית הכסא, כמו שנאמר ולא תתורו אחרי לבבכם וכו' כמו שפירשו רז"ל. ומאמיני הקבלה החדשה הזו אשר נחלוה מיד המלך העכו"ם, ישמיו תהלה במאורי ישראל ומלאכי ה', הרס"ג ובעל חובת הלבבות ורבי יהודה הלוי והרמב"ם ובעל העיקרים כאשר יראה המעיין בספר דרך חיים לרבי מנחם דילונאזנו איך תלה בבעל חה"ל בוקי סריקי מה שלא ניתכוון להם החסיד בעל חובת הלבבות! עוד הוסיף לנטות ידו לבזות את החסיד על מה שכתב בהקדמה שחכם אחד נשאל שאלה נכרית מענייני גירושין, והחכם ההוא השיבו אתה האיש השואל על מה שלא יזיקנו אם לא ידענו הידעת כל מה שאתה חייב לדעתו מן המצות אשר אינך ראשי להתעלם מהן עד שנפנית לחשוב בשאלות נכריות אשר לא תקנה בהם מעלה יתירה בתורתך ואמונתך וכו'. והרחיב הר"מ לונזאנו את פיו לדבר תועה על מלאך ה' צבאות החסיד בעל חובת הלבבות, ולא זכר שר מאי דגרסינן במנחות שרבי גזר על פלימו שיקבל עליו שמתא או יגלה על ששאל לו שאלה נכרית לענין תפלין, דהכי גרסינן בפ' הקומץ, בעא מניה פלימו מר' מי שיש לו שני ראשים באיזה מהם מניח תפילין אמר ליה או קום גלי או קביל עלך שמתא, ואלמלא חס הקב"ה על כבוד צורבא מרבנן פלימו ואתא ההוא נוברא ואמר אתיליד לי ינוקא דאית ליה תרין רישין כמא בעינא למיתב לכהן וכו'! ודברים של הר"מ לונזאנו הי"ל לדבר על ר' שהשיב על שאלת פלימו או קום גלי או קביל עלך שמתא? אלמלא דכל השואל שאלות זרות מאד מנדין אותו, שגורם לבייש את החכם בחנם, וכ"כ מרן בכ"מ סוף פ"ו מהל' ת"ת שהשואל דבר שאי אפשר הוא בכלל מזלזל בד"ת או המבזה את החכם.

| id | word | punctuation after |
|---|---|---|
| 9:11.1 | קכח | ) |
| 9:11.2 | עוד |  |
| 9:11.3 | זאת |  |
| 9:11.4 | רעה |  |
| 9:11.5 | חולה | , |
| 9:11.6 | שכל |  |
| 9:11.7 | הלומדים |  |
| 9:11.8 | בזהר |  |
| 9:11.9 | ושאר |  |
| 9:11.10 | ספרי |  |
| 9:11.11 | קבלה |  |
| 9:11.12 | חדשה |  |
| 9:11.13 | נקל |  |
| 9:11.14 | בעיניהם |  |
| 9:11.15 | לעולל |  |
| 9:11.16 | עלילות |  |
| 9:11.17 | ולשם |  |
| 9:11.18 | תהלה |  |
| 9:11.19 | במאורי |  |
| 9:11.20 | ישראל |  |
| 9:11.21 | מורי |  |
| 9:11.22 | צדק |  |
| 9:11.23 | המאירים |  |
| 9:11.24 | עיני |  |
| 9:11.25 | ישראל |  |
| 9:11.26 | בכל |  |
| 9:11.27 | מקצועות |  |
| 9:11.28 | התורה |  |
| 9:11.29 | אשר |  |
| 9:11.30 | יעשה |  |
| 9:11.31 | אותם |  |
| 9:11.32 | האדם |  |
| 9:11.33 | וחי |  |
| 9:11.34 | בהם | , |
| 9:11.35 | ובפרט |  |
| 9:11.36 | בעיקרי |  |
| 9:11.37 | תורתינו |  |
| 9:11.38 | איך |  |
| 9:11.39 | נאמין |  |
| 9:11.40 | ונקיים |  |
| 9:11.41 | מצותיה | ! |
| 9:11.42 | וביחוד |  |
| 9:11.43 | העיקר |  |
| 9:11.44 | והעמוד |  |
| 9:11.45 | שהכל |  |
| 9:11.46 | תלוי |  |
| 9:11.47 | בו | , |
| 9:11.48 | והוא |  |
| 9:11.49 | יחוד |  |
| 9:11.50 | השי"ת |  |
| 9:11.51 | לא |  |
| 9:11.52 | שנאמר |  |
| 9:11.53 | בפינו |  |
| 9:11.54 | אחד | , |
| 9:11.55 | ובלבבינו |  |
| 9:11.56 | רבים | , |
| 9:11.57 | כמו |  |
| 9:11.58 | שאומרים |  |
| 9:11.59 | המקובלים |  |
| 9:11.60 | כי |  |
| 9:11.61 | האלהים |  |
| 9:11.62 | העליון |  |
| 9:11.63 | דלא |  |
| 9:11.64 | שייך |  |
| 9:11.65 | ביה |  |
| 9:11.66 | עבודה |  |
| 9:11.67 | המקשר |  |
| 9:11.68 | ומחבר |  |
| 9:11.69 | ומזויג |  |
| 9:11.70 | את |  |
| 9:11.71 | האלהים |  |
| 9:11.72 | הרבים |  |
| 9:11.73 | התחתונים |  |
| 9:11.74 | אשר |  |
| 9:11.75 | אותם |  |
| 9:11.76 | נעבוד |  |
| 9:11.77 | להיותם |  |
| 9:11.78 | אחד | ! |
| 9:11.79 | אוי |  |
| 9:11.80 | לעינים |  |
| 9:11.81 | שכך |  |
| 9:11.82 | רואות | , |
| 9:11.83 | ואוי |  |
| 9:11.84 | לאזנים |  |
| 9:11.85 | שכך |  |
| 9:11.86 | שומעות | , |
| 9:11.87 | ודברים |  |
| 9:11.88 | כאלה |  |
| 9:11.89 | הן |  |
| 9:11.90 | שאסור |  |
| 9:11.91 | לשמען |  |
| 9:11.92 | ואסור |  |
| 9:11.93 | להרהר |  |
| 9:11.94 | בהן |  |
| 9:11.95 | אפילו |  |
| 9:11.96 | בבית |  |
| 9:11.97 | הכסא | , |
| 9:11.98 | כמו |  |
| 9:11.99 | שנאמר |  |
| 9:11.100 | ולא |  |
| 9:11.101 | תתורו |  |
| 9:11.102 | אחרי |  |
| 9:11.103 | לבבכם |  |
| 9:11.104 | וכו | ' |
| 9:11.105 | כמו |  |
| 9:11.106 | שפירשו |  |
| 9:11.107 | רז"ל | . |
| 9:11.108 | ומאמיני |  |
| 9:11.109 | הקבלה |  |
| 9:11.110 | החדשה |  |
| 9:11.111 | הזו |  |
| 9:11.112 | אשר |  |
| 9:11.113 | נחלוה |  |
| 9:11.114 | מיד |  |
| 9:11.115 | המלך |  |
| 9:11.116 | העכו"ם | , |
| 9:11.117 | ישמיו |  |
| 9:11.118 | תהלה |  |
| 9:11.119 | במאורי |  |
| 9:11.120 | ישראל |  |
| 9:11.121 | ומלאכי |  |
| 9:11.122 | ה | ', |
| 9:11.123 | הרס"ג |  |
| 9:11.124 | ובעל |  |
| 9:11.125 | חובת |  |
| 9:11.126 | הלבבות |  |
| 9:11.127 | ורבי |  |
| 9:11.128 | יהודה |  |
| 9:11.129 | הלוי |  |
| 9:11.130 | והרמב"ם |  |
| 9:11.131 | ובעל |  |
| 9:11.132 | העיקרים |  |
| 9:11.133 | כאשר |  |
| 9:11.134 | יראה |  |
| 9:11.135 | המעיין |  |
| 9:11.136 | בספר |  |
| 9:11.137 | דרך |  |
| 9:11.138 | חיים |  |
| 9:11.139 | לרבי |  |
| 9:11.140 | מנחם |  |
| 9:11.141 | דילונאזנו |  |
| 9:11.142 | איך |  |
| 9:11.143 | תלה |  |
| 9:11.144 | בבעל |  |
| 9:11.145 | חה"ל |  |
| 9:11.146 | בוקי |  |
| 9:11.147 | סריקי |  |
| 9:11.148 | מה |  |
| 9:11.149 | שלא |  |
| 9:11.150 | ניתכוון |  |
| 9:11.151 | להם |  |
| 9:11.152 | החסיד |  |
| 9:11.153 | בעל |  |
| 9:11.154 | חובת |  |
| 9:11.155 | הלבבות | ! |
| 9:11.156 | עוד |  |
| 9:11.157 | הוסיף |  |
| 9:11.158 | לנטות |  |
| 9:11.159 | ידו |  |
| 9:11.160 | לבזות |  |
| 9:11.161 | את |  |
| 9:11.162 | החסיד |  |
| 9:11.163 | על |  |
| 9:11.164 | מה |  |
| 9:11.165 | שכתב |  |
| 9:11.166 | בהקדמה |  |
| 9:11.167 | שחכם |  |
| 9:11.168 | אחד |  |
| 9:11.169 | נשאל |  |
| 9:11.170 | שאלה |  |
| 9:11.171 | נכרית |  |
| 9:11.172 | מענייני |  |
| 9:11.173 | גירושין | , |
| 9:11.174 | והחכם |  |
| 9:11.175 | ההוא |  |
| 9:11.176 | השיבו |  |
| 9:11.177 | אתה |  |
| 9:11.178 | האיש |  |
| 9:11.179 | השואל |  |
| 9:11.180 | על |  |
| 9:11.181 | מה |  |
| 9:11.182 | שלא |  |
| 9:11.183 | יזיקנו |  |
| 9:11.184 | אם |  |
| 9:11.185 | לא |  |
| 9:11.186 | ידענו |  |
| 9:11.187 | הידעת |  |
| 9:11.188 | כל |  |
| 9:11.189 | מה |  |
| 9:11.190 | שאתה |  |
| 9:11.191 | חייב |  |
| 9:11.192 | לדעתו |  |
| 9:11.193 | מן |  |
| 9:11.194 | המצות |  |
| 9:11.195 | אשר |  |
| 9:11.196 | אינך |  |
| 9:11.197 | ראשי |  |
| 9:11.198 | להתעלם |  |
| 9:11.199 | מהן |  |
| 9:11.200 | עד |  |
| 9:11.201 | שנפנית |  |
| 9:11.202 | לחשוב |  |
| 9:11.203 | בשאלות |  |
| 9:11.204 | נכריות |  |
| 9:11.205 | אשר |  |
| 9:11.206 | לא |  |
| 9:11.207 | תקנה |  |
| 9:11.208 | בהם |  |
| 9:11.209 | מעלה |  |
| 9:11.210 | יתירה |  |
| 9:11.211 | בתורתך |  |
| 9:11.212 | ואמונתך |  |
| 9:11.213 | וכו | '. |
| 9:11.214 | והרחיב |  |
| 9:11.215 | הר"מ |  |
| 9:11.216 | לונזאנו |  |
| 9:11.217 | את |  |
| 9:11.218 | פיו |  |
| 9:11.219 | לדבר |  |
| 9:11.220 | תועה |  |
| 9:11.221 | על |  |
| 9:11.222 | מלאך |  |
| 9:11.223 | ה | ' |
| 9:11.224 | צבאות |  |
| 9:11.225 | החסיד |  |
| 9:11.226 | בעל |  |
| 9:11.227 | חובת |  |
| 9:11.228 | הלבבות | , |
| 9:11.229 | ולא |  |
| 9:11.230 | זכר |  |
| 9:11.231 | שר |  |
| 9:11.232 | מאי |  |
| 9:11.233 | דגרסינן |  |
| 9:11.234 | במנחות |  |
| 9:11.235 | שרבי |  |
| 9:11.236 | גזר |  |
| 9:11.237 | על |  |
| 9:11.238 | פלימו |  |
| 9:11.239 | שיקבל |  |
| 9:11.240 | עליו |  |
| 9:11.241 | שמתא |  |
| 9:11.242 | או |  |
| 9:11.243 | יגלה |  |
| 9:11.244 | על |  |
| 9:11.245 | ששאל |  |
| 9:11.246 | לו |  |
| 9:11.247 | שאלה |  |
| 9:11.248 | נכרית |  |
| 9:11.249 | לענין |  |
| 9:11.250 | תפלין | , |
| 9:11.251 | דהכי |  |
| 9:11.252 | גרסינן |  |
| 9:11.253 | בפ | ' |
| 9:11.254 | הקומץ | , |
| 9:11.255 | בעא |  |
| 9:11.256 | מניה |  |
| 9:11.257 | פלימו |  |
| 9:11.258 | מר | ' |
| 9:11.259 | מי |  |
| 9:11.260 | שיש |  |
| 9:11.261 | לו |  |
| 9:11.262 | שני |  |
| 9:11.263 | ראשים |  |
| 9:11.264 | באיזה |  |
| 9:11.265 | מהם |  |
| 9:11.266 | מניח |  |
| 9:11.267 | תפילין |  |
| 9:11.268 | אמר |  |
| 9:11.269 | ליה |  |
| 9:11.270 | או |  |
| 9:11.271 | קום |  |
| 9:11.272 | גלי |  |
| 9:11.273 | או |  |
| 9:11.274 | קביל |  |
| 9:11.275 | עלך |  |
| 9:11.276 | שמתא | , |
| 9:11.277 | ואלמלא |  |
| 9:11.278 | חס |  |
| 9:11.279 | הקב"ה |  |
| 9:11.280 | על |  |
| 9:11.281 | כבוד |  |
| 9:11.282 | צורבא |  |
| 9:11.283 | מרבנן |  |
| 9:11.284 | פלימו |  |
| 9:11.285 | ואתא |  |
| 9:11.286 | ההוא |  |
| 9:11.287 | נוברא |  |
| 9:11.288 | ואמר |  |
| 9:11.289 | אתיליד |  |
| 9:11.290 | לי |  |
| 9:11.291 | ינוקא |  |
| 9:11.292 | דאית |  |
| 9:11.293 | ליה |  |
| 9:11.294 | תרין |  |
| 9:11.295 | רישין |  |
| 9:11.296 | כמא |  |
| 9:11.297 | בעינא |  |
| 9:11.298 | למיתב |  |
| 9:11.299 | לכהן |  |
| 9:11.300 | וכו | '! |
| 9:11.301 | ודברים |  |
| 9:11.302 | של |  |
| 9:11.303 | הר"מ |  |
| 9:11.304 | לונזאנו |  |
| 9:11.305 | הי"ל |  |
| 9:11.306 | לדבר |  |
| 9:11.307 | על |  |
| 9:11.308 | ר | ' |
| 9:11.309 | שהשיב |  |
| 9:11.310 | על |  |
| 9:11.311 | שאלת |  |
| 9:11.312 | פלימו |  |
| 9:11.313 | או |  |
| 9:11.314 | קום |  |
| 9:11.315 | גלי |  |
| 9:11.316 | או |  |
| 9:11.317 | קביל |  |
| 9:11.318 | עלך |  |
| 9:11.319 | שמתא | ? |
| 9:11.320 | אלמלא |  |
| 9:11.321 | דכל |  |
| 9:11.322 | השואל |  |
| 9:11.323 | שאלות |  |
| 9:11.324 | זרות |  |
| 9:11.325 | מאד |  |
| 9:11.326 | מנדין |  |
| 9:11.327 | אותו | , |
| 9:11.328 | שגורם |  |
| 9:11.329 | לבייש |  |
| 9:11.330 | את |  |
| 9:11.331 | החכם |  |
| 9:11.332 | בחנם | , |
| 9:11.333 | וכ"כ |  |
| 9:11.334 | מרן |  |
| 9:11.335 | בכ"מ |  |
| 9:11.336 | סוף |  |
| 9:11.337 | פ"ו |  |
| 9:11.338 | מהל | ' |
| 9:11.339 | ת"ת |  |
| 9:11.340 | שהשואל |  |
| 9:11.341 | דבר |  |
| 9:11.342 | שאי |  |
| 9:11.343 | אפשר |  |
| 9:11.344 | הוא |  |
| 9:11.345 | בכלל |  |
| 9:11.346 | מזלזל |  |
| 9:11.347 | בד"ת |  |
| 9:11.348 | או |  |
| 9:11.349 | המבזה |  |
| 9:11.350 | את |  |
| 9:11.351 | החכם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

