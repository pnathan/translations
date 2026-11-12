# Translation pack: milhamot-hashem perek-9-06 (Milchamot ha-Shem 9:7–9:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-06`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **חז"ל** → "our Sages of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
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

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **רחמים**: "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes
  of mercy", a different thirteen from the hermeneutical rules.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- מי **אתי** (ketiv): translated, the qere noted.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 9:7

Hebrew (Wikisource): קכה) וכל הדברים האלה קמו עמדו להאמין בהם בלי קולות וברקים ולפידים וענן וקול שופר ואפי' אות ומופת כמו שאמר הכתוב בנביא השקר ונתן אליך אות או מופת לא שמענו ולא ראינו ולא בת קול מן השמים קראה באזנינו רק דיו שחור על נייר לבן סיפורי פלאות בדויים משם עתיקא דקישא שנראה לבית מדרשו של רשב"י המשקר (ולא לרשב"י תנא דידן ח"ו על כבודו) ושאל לו על פסוק ויאמר אלהים נעשה אדם וכו' כדאיתא בזוהר בראשית כ"ב. היעלה על לב ישראל לשאול על אחד מעשרה מאמרות שבהן נברא העולם מי האומר מאמר זה ומי האומר מאמר זה או על דעתו לומר כי האומר ראו עתה כי אני אני הוא ואין אלהים עמדי ולא זהו שאמר כל העשרה מאמרות? וגם הרבה ספורים משם אליהו ורעיא מהימנא ואחרים מהתנאים והאמוראים והאחרונים אשר כל משכיל בתורה התמימה ודברי נביאינו וחז"ל ישום וישרוק ועמד משתאה ומשתומם איך נוציא ישן אשר שמענו במעמד הר סיני הנבחר ויקם עדות ביעקב ותורה שם בישראל להודיעם לבניהם בקולות וברקים ולפידים וענן כבד וקול ה' בכח מדבר מתוך האש אשר לא עשה כן לכל גוי ובקול עוז ותעצומות דבר עמנו פנים בפנים אנכי ה' אלהיך אשר הוציאתיך וכו' ונאמין שאין לי אב ולא אח ולא בן ראו עתה כי אני אני הוא ואין אלהים עמדי וכו' ונאמין למלך העכו"ם שהביא לנו תורה שבעל פה חדשה האומרת כי האלהים ברא מעצמותו חמשת פרצופין והשליטם על מעשה בארית והם השליטו הזעיר אנפין על כל הברואים וצוו את הכל לעבדו והוא אלהינו ואנחנו עמו ונחלתו?

| id | word | punctuation after |
|---|---|---|
| 9:7.1 | קכה | ) |
| 9:7.2 | וכל |  |
| 9:7.3 | הדברים |  |
| 9:7.4 | האלה |  |
| 9:7.5 | קמו |  |
| 9:7.6 | עמדו |  |
| 9:7.7 | להאמין |  |
| 9:7.8 | בהם |  |
| 9:7.9 | בלי |  |
| 9:7.10 | קולות |  |
| 9:7.11 | וברקים |  |
| 9:7.12 | ולפידים |  |
| 9:7.13 | וענן |  |
| 9:7.14 | וקול |  |
| 9:7.15 | שופר |  |
| 9:7.16 | ואפי | ' |
| 9:7.17 | אות |  |
| 9:7.18 | ומופת |  |
| 9:7.19 | כמו |  |
| 9:7.20 | שאמר |  |
| 9:7.21 | הכתוב |  |
| 9:7.22 | בנביא |  |
| 9:7.23 | השקר |  |
| 9:7.24 | ונתן |  |
| 9:7.25 | אליך |  |
| 9:7.26 | אות |  |
| 9:7.27 | או |  |
| 9:7.28 | מופת |  |
| 9:7.29 | לא |  |
| 9:7.30 | שמענו |  |
| 9:7.31 | ולא |  |
| 9:7.32 | ראינו |  |
| 9:7.33 | ולא |  |
| 9:7.34 | בת |  |
| 9:7.35 | קול |  |
| 9:7.36 | מן |  |
| 9:7.37 | השמים |  |
| 9:7.38 | קראה |  |
| 9:7.39 | באזנינו |  |
| 9:7.40 | רק |  |
| 9:7.41 | דיו |  |
| 9:7.42 | שחור |  |
| 9:7.43 | על |  |
| 9:7.44 | נייר |  |
| 9:7.45 | לבן |  |
| 9:7.46 | סיפורי |  |
| 9:7.47 | פלאות |  |
| 9:7.48 | בדויים |  |
| 9:7.49 | משם |  |
| 9:7.50 | עתיקא |  |
| 9:7.51 | דקישא |  |
| 9:7.52 | שנראה |  |
| 9:7.53 | לבית |  |
| 9:7.54 | מדרשו |  |
| 9:7.55 | של |  |
| 9:7.56 | רשב"י |  |
| 9:7.57 | המשקר |  |
| 9:7.58 | ולא |  |
| 9:7.59 | לרשב"י |  |
| 9:7.60 | תנא |  |
| 9:7.61 | דידן |  |
| 9:7.62 | ח"ו |  |
| 9:7.63 | על |  |
| 9:7.64 | כבודו | ) |
| 9:7.65 | ושאל |  |
| 9:7.66 | לו |  |
| 9:7.67 | על |  |
| 9:7.68 | פסוק |  |
| 9:7.69 | ויאמר |  |
| 9:7.70 | אלהים |  |
| 9:7.71 | נעשה |  |
| 9:7.72 | אדם |  |
| 9:7.73 | וכו | ' |
| 9:7.74 | כדאיתא |  |
| 9:7.75 | בזוהר |  |
| 9:7.76 | בראשית |  |
| 9:7.77 | כ"ב | . |
| 9:7.78 | היעלה |  |
| 9:7.79 | על |  |
| 9:7.80 | לב |  |
| 9:7.81 | ישראל |  |
| 9:7.82 | לשאול |  |
| 9:7.83 | על |  |
| 9:7.84 | אחד |  |
| 9:7.85 | מעשרה |  |
| 9:7.86 | מאמרות |  |
| 9:7.87 | שבהן |  |
| 9:7.88 | נברא |  |
| 9:7.89 | העולם |  |
| 9:7.90 | מי |  |
| 9:7.91 | האומר |  |
| 9:7.92 | מאמר |  |
| 9:7.93 | זה |  |
| 9:7.94 | ומי |  |
| 9:7.95 | האומר |  |
| 9:7.96 | מאמר |  |
| 9:7.97 | זה |  |
| 9:7.98 | או |  |
| 9:7.99 | על |  |
| 9:7.100 | דעתו |  |
| 9:7.101 | לומר |  |
| 9:7.102 | כי |  |
| 9:7.103 | האומר |  |
| 9:7.104 | ראו |  |
| 9:7.105 | עתה |  |
| 9:7.106 | כי |  |
| 9:7.107 | אני |  |
| 9:7.108 | אני |  |
| 9:7.109 | הוא |  |
| 9:7.110 | ואין |  |
| 9:7.111 | אלהים |  |
| 9:7.112 | עמדי |  |
| 9:7.113 | ולא |  |
| 9:7.114 | זהו |  |
| 9:7.115 | שאמר |  |
| 9:7.116 | כל |  |
| 9:7.117 | העשרה |  |
| 9:7.118 | מאמרות | ? |
| 9:7.119 | וגם |  |
| 9:7.120 | הרבה |  |
| 9:7.121 | ספורים |  |
| 9:7.122 | משם |  |
| 9:7.123 | אליהו |  |
| 9:7.124 | ורעיא |  |
| 9:7.125 | מהימנא |  |
| 9:7.126 | ואחרים |  |
| 9:7.127 | מהתנאים |  |
| 9:7.128 | והאמוראים |  |
| 9:7.129 | והאחרונים |  |
| 9:7.130 | אשר |  |
| 9:7.131 | כל |  |
| 9:7.132 | משכיל |  |
| 9:7.133 | בתורה |  |
| 9:7.134 | התמימה |  |
| 9:7.135 | ודברי |  |
| 9:7.136 | נביאינו |  |
| 9:7.137 | וחז"ל |  |
| 9:7.138 | ישום |  |
| 9:7.139 | וישרוק |  |
| 9:7.140 | ועמד |  |
| 9:7.141 | משתאה |  |
| 9:7.142 | ומשתומם |  |
| 9:7.143 | איך |  |
| 9:7.144 | נוציא |  |
| 9:7.145 | ישן |  |
| 9:7.146 | אשר |  |
| 9:7.147 | שמענו |  |
| 9:7.148 | במעמד |  |
| 9:7.149 | הר |  |
| 9:7.150 | סיני |  |
| 9:7.151 | הנבחר |  |
| 9:7.152 | ויקם |  |
| 9:7.153 | עדות |  |
| 9:7.154 | ביעקב |  |
| 9:7.155 | ותורה |  |
| 9:7.156 | שם |  |
| 9:7.157 | בישראל |  |
| 9:7.158 | להודיעם |  |
| 9:7.159 | לבניהם |  |
| 9:7.160 | בקולות |  |
| 9:7.161 | וברקים |  |
| 9:7.162 | ולפידים |  |
| 9:7.163 | וענן |  |
| 9:7.164 | כבד |  |
| 9:7.165 | וקול |  |
| 9:7.166 | ה | ' |
| 9:7.167 | בכח |  |
| 9:7.168 | מדבר |  |
| 9:7.169 | מתוך |  |
| 9:7.170 | האש |  |
| 9:7.171 | אשר |  |
| 9:7.172 | לא |  |
| 9:7.173 | עשה |  |
| 9:7.174 | כן |  |
| 9:7.175 | לכל |  |
| 9:7.176 | גוי |  |
| 9:7.177 | ובקול |  |
| 9:7.178 | עוז |  |
| 9:7.179 | ותעצומות |  |
| 9:7.180 | דבר |  |
| 9:7.181 | עמנו |  |
| 9:7.182 | פנים |  |
| 9:7.183 | בפנים |  |
| 9:7.184 | אנכי |  |
| 9:7.185 | ה | ' |
| 9:7.186 | אלהיך |  |
| 9:7.187 | אשר |  |
| 9:7.188 | הוציאתיך |  |
| 9:7.189 | וכו | ' |
| 9:7.190 | ונאמין |  |
| 9:7.191 | שאין |  |
| 9:7.192 | לי |  |
| 9:7.193 | אב |  |
| 9:7.194 | ולא |  |
| 9:7.195 | אח |  |
| 9:7.196 | ולא |  |
| 9:7.197 | בן |  |
| 9:7.198 | ראו |  |
| 9:7.199 | עתה |  |
| 9:7.200 | כי |  |
| 9:7.201 | אני |  |
| 9:7.202 | אני |  |
| 9:7.203 | הוא |  |
| 9:7.204 | ואין |  |
| 9:7.205 | אלהים |  |
| 9:7.206 | עמדי |  |
| 9:7.207 | וכו | ' |
| 9:7.208 | ונאמין |  |
| 9:7.209 | למלך |  |
| 9:7.210 | העכו"ם |  |
| 9:7.211 | שהביא |  |
| 9:7.212 | לנו |  |
| 9:7.213 | תורה |  |
| 9:7.214 | שבעל |  |
| 9:7.215 | פה |  |
| 9:7.216 | חדשה |  |
| 9:7.217 | האומרת |  |
| 9:7.218 | כי |  |
| 9:7.219 | האלהים |  |
| 9:7.220 | ברא |  |
| 9:7.221 | מעצמותו |  |
| 9:7.222 | חמשת |  |
| 9:7.223 | פרצופין |  |
| 9:7.224 | והשליטם |  |
| 9:7.225 | על |  |
| 9:7.226 | מעשה |  |
| 9:7.227 | בארית |  |
| 9:7.228 | והם |  |
| 9:7.229 | השליטו |  |
| 9:7.230 | הזעיר |  |
| 9:7.231 | אנפין |  |
| 9:7.232 | על |  |
| 9:7.233 | כל |  |
| 9:7.234 | הברואים |  |
| 9:7.235 | וצוו |  |
| 9:7.236 | את |  |
| 9:7.237 | הכל |  |
| 9:7.238 | לעבדו |  |
| 9:7.239 | והוא |  |
| 9:7.240 | אלהינו |  |
| 9:7.241 | ואנחנו |  |
| 9:7.242 | עמו |  |
| 9:7.243 | ונחלתו | ? |


## Milchamot ha-Shem 9:8

Hebrew (Wikisource): ואם ח"ו רצונו ית' להמיר כבודו באל אחר קצר אפים ולהשליטו על כל הברואים כמו שאמר הזהר שולטנו אתיהיב ליה וכו' היפך מה שצונו בתחלה תמר בו ואמרו רז"ל אל תמירני בו ואל תמירהו בי (רזה פרשת משפטים) מה לא יט שמים וירד על איזה נר שיבחר וידבר עמנו עוד הפעם לאמר עתה בסוף הימים נלאתי כלכל והמלכתי במקומי את קצר האפים? זאת לא תהיה לא זו הדרך ולא זו העיר ולא יחליף האל ולא ימיר ולא יעשה כן בישראל להמיר את אלהינו אשר הודיע דרכיו למשה כי הוא ארך אפים ורב חסד באל אחר קצר אפים ובעל אשה. ורז"ל אמרו ברית כרותה לשלש עשרה מדות של רחמים שאינן חוזרות ריקם ולכך קבעו בתפלות של פרקים על כל צרה שלא תבוא וגם בימים נוראים להגביר הי"ג מדות של רחמים ואיך תיסק אדעתין ליחסם לזעיר אנפין שהוא קצר אפים הא הוו הני מילי סתרן אחדדי קורים אנו אותו ז"א שעניינו קצר אפים ומשבחים אותו בההיפך שהוא ארך אפים? אתמהא?

| id | word | punctuation after |
|---|---|---|
| 9:8.1 | ואם |  |
| 9:8.2 | ח"ו |  |
| 9:8.3 | רצונו |  |
| 9:8.4 | ית | ' |
| 9:8.5 | להמיר |  |
| 9:8.6 | כבודו |  |
| 9:8.7 | באל |  |
| 9:8.8 | אחר |  |
| 9:8.9 | קצר |  |
| 9:8.10 | אפים |  |
| 9:8.11 | ולהשליטו |  |
| 9:8.12 | על |  |
| 9:8.13 | כל |  |
| 9:8.14 | הברואים |  |
| 9:8.15 | כמו |  |
| 9:8.16 | שאמר |  |
| 9:8.17 | הזהר |  |
| 9:8.18 | שולטנו |  |
| 9:8.19 | אתיהיב |  |
| 9:8.20 | ליה |  |
| 9:8.21 | וכו | ' |
| 9:8.22 | היפך |  |
| 9:8.23 | מה |  |
| 9:8.24 | שצונו |  |
| 9:8.25 | בתחלה |  |
| 9:8.26 | תמר |  |
| 9:8.27 | בו |  |
| 9:8.28 | ואמרו |  |
| 9:8.29 | רז"ל |  |
| 9:8.30 | אל |  |
| 9:8.31 | תמירני |  |
| 9:8.32 | בו |  |
| 9:8.33 | ואל |  |
| 9:8.34 | תמירהו |  |
| 9:8.35 | בי |  |
| 9:8.36 | רזה |  |
| 9:8.37 | פרשת |  |
| 9:8.38 | משפטים | ) |
| 9:8.39 | מה |  |
| 9:8.40 | לא |  |
| 9:8.41 | יט |  |
| 9:8.42 | שמים |  |
| 9:8.43 | וירד |  |
| 9:8.44 | על |  |
| 9:8.45 | איזה |  |
| 9:8.46 | נר |  |
| 9:8.47 | שיבחר |  |
| 9:8.48 | וידבר |  |
| 9:8.49 | עמנו |  |
| 9:8.50 | עוד |  |
| 9:8.51 | הפעם |  |
| 9:8.52 | לאמר |  |
| 9:8.53 | עתה |  |
| 9:8.54 | בסוף |  |
| 9:8.55 | הימים |  |
| 9:8.56 | נלאתי |  |
| 9:8.57 | כלכל |  |
| 9:8.58 | והמלכתי |  |
| 9:8.59 | במקומי |  |
| 9:8.60 | את |  |
| 9:8.61 | קצר |  |
| 9:8.62 | האפים | ? |
| 9:8.63 | זאת |  |
| 9:8.64 | לא |  |
| 9:8.65 | תהיה |  |
| 9:8.66 | לא |  |
| 9:8.67 | זו |  |
| 9:8.68 | הדרך |  |
| 9:8.69 | ולא |  |
| 9:8.70 | זו |  |
| 9:8.71 | העיר |  |
| 9:8.72 | ולא |  |
| 9:8.73 | יחליף |  |
| 9:8.74 | האל |  |
| 9:8.75 | ולא |  |
| 9:8.76 | ימיר |  |
| 9:8.77 | ולא |  |
| 9:8.78 | יעשה |  |
| 9:8.79 | כן |  |
| 9:8.80 | בישראל |  |
| 9:8.81 | להמיר |  |
| 9:8.82 | את |  |
| 9:8.83 | אלהינו |  |
| 9:8.84 | אשר |  |
| 9:8.85 | הודיע |  |
| 9:8.86 | דרכיו |  |
| 9:8.87 | למשה |  |
| 9:8.88 | כי |  |
| 9:8.89 | הוא |  |
| 9:8.90 | ארך |  |
| 9:8.91 | אפים |  |
| 9:8.92 | ורב |  |
| 9:8.93 | חסד |  |
| 9:8.94 | באל |  |
| 9:8.95 | אחר |  |
| 9:8.96 | קצר |  |
| 9:8.97 | אפים |  |
| 9:8.98 | ובעל |  |
| 9:8.99 | אשה | . |
| 9:8.100 | ורז"ל |  |
| 9:8.101 | אמרו |  |
| 9:8.102 | ברית |  |
| 9:8.103 | כרותה |  |
| 9:8.104 | לשלש |  |
| 9:8.105 | עשרה |  |
| 9:8.106 | מדות |  |
| 9:8.107 | של |  |
| 9:8.108 | רחמים |  |
| 9:8.109 | שאינן |  |
| 9:8.110 | חוזרות |  |
| 9:8.111 | ריקם |  |
| 9:8.112 | ולכך |  |
| 9:8.113 | קבעו |  |
| 9:8.114 | בתפלות |  |
| 9:8.115 | של |  |
| 9:8.116 | פרקים |  |
| 9:8.117 | על |  |
| 9:8.118 | כל |  |
| 9:8.119 | צרה |  |
| 9:8.120 | שלא |  |
| 9:8.121 | תבוא |  |
| 9:8.122 | וגם |  |
| 9:8.123 | בימים |  |
| 9:8.124 | נוראים |  |
| 9:8.125 | להגביר |  |
| 9:8.126 | הי"ג |  |
| 9:8.127 | מדות |  |
| 9:8.128 | של |  |
| 9:8.129 | רחמים |  |
| 9:8.130 | ואיך |  |
| 9:8.131 | תיסק |  |
| 9:8.132 | אדעתין |  |
| 9:8.133 | ליחסם |  |
| 9:8.134 | לזעיר |  |
| 9:8.135 | אנפין |  |
| 9:8.136 | שהוא |  |
| 9:8.137 | קצר |  |
| 9:8.138 | אפים |  |
| 9:8.139 | הא |  |
| 9:8.140 | הוו |  |
| 9:8.141 | הני |  |
| 9:8.142 | מילי |  |
| 9:8.143 | סתרן |  |
| 9:8.144 | אחדדי |  |
| 9:8.145 | קורים |  |
| 9:8.146 | אנו |  |
| 9:8.147 | אותו |  |
| 9:8.148 | ז"א |  |
| 9:8.149 | שעניינו |  |
| 9:8.150 | קצר |  |
| 9:8.151 | אפים |  |
| 9:8.152 | ומשבחים |  |
| 9:8.153 | אותו |  |
| 9:8.154 | בההיפך |  |
| 9:8.155 | שהוא |  |
| 9:8.156 | ארך |  |
| 9:8.157 | אפים | ? |
| 9:8.158 | אתמהא | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

