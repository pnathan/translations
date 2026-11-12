# Translation pack: milhamot-hashem perek-4-09 (Milchamot ha-Shem 4:14–4:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-09`.

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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **ריבה / לרבות** → "include", never "multiply".
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **אחי** → "my brother" (second person kept).
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
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ח"א / ח"ב** → "vol. 1" / "vol. 2".
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


## Milchamot ha-Shem 4:14

Hebrew (Wikisource): כז) והוא שדעת הזהר והמקובלים הנמשכים אחריו שכל עבודתינו ותפלותינו לזעיר אנפין דוקא, וז"ל הזהר וירא דף קיב (ח"א קיב, א) תא חזי לה' אלהיך תירא לא כתיב אלא את ה' אלהיך תירא, מאי את דא דרגא קדמאה (פירש בגליון מלכות רמוזה בתיבת את ה' אלהיך תירא) אחר דחלא דקב"ה, ובגין כך כתיב תירא, דתמן בעי בר נש לדחלא קמי מריה בגין דאיהו בי דינא, ואותו תעבוד דא דרגא עלאה דקימא על האיי דרגא תתאה ולא מתפרשין לעלמין, את ואותו דא בדא דבקין ולא אתפרשן, מאי ואותו דא אתר ברית קדישא אות לעלמין, דהא פולחנא לא שריא באת ולא איהו למיפלח אלא למדחל, אבל פולחנא איהו לעילא, ובגין כך ואותו תעבוד (פי' אבל העבודה היא לזעיר אנפין כמו שנ' ואותו תעבוד ותייחד עמיה שכנתיה ועיין בס' המצות בתפלת אליהו הנביא מפורש באורך) ע"כ לשון הזהר עם פי' הרש"ב בגליון. וכונת הזהר לדרוש מדלא כתיב לה' אלהיך תירא ותעבוד, וכתיב את לרבות, ריבה מדת המלכות שהיא נוקבא דזעיר אנפין ליראה, ומדכתיב ואותו תעבוד מיעט אותה מעבודה דדוקא אותו תעבוד ולא אותה (דהיינו מלכות) הרי ריבה אותה ליראה ומיעט אותה מעבודה שהעבודה לזעיר אנפין דוקא.

| id | word | punctuation after |
|---|---|---|
| 4:14.1 | כז | ) |
| 4:14.2 | והוא |  |
| 4:14.3 | שדעת |  |
| 4:14.4 | הזהר |  |
| 4:14.5 | והמקובלים |  |
| 4:14.6 | הנמשכים |  |
| 4:14.7 | אחריו |  |
| 4:14.8 | שכל |  |
| 4:14.9 | עבודתינו |  |
| 4:14.10 | ותפלותינו |  |
| 4:14.11 | לזעיר |  |
| 4:14.12 | אנפין |  |
| 4:14.13 | דוקא | , |
| 4:14.14 | וז"ל |  |
| 4:14.15 | הזהר |  |
| 4:14.16 | וירא |  |
| 4:14.17 | דף |  |
| 4:14.18 | קיב |  |
| 4:14.19 | ח"א |  |
| 4:14.20 | קיב | , |
| 4:14.21 | א | ) |
| 4:14.22 | תא |  |
| 4:14.23 | חזי |  |
| 4:14.24 | לה | ' |
| 4:14.25 | אלהיך |  |
| 4:14.26 | תירא |  |
| 4:14.27 | לא |  |
| 4:14.28 | כתיב |  |
| 4:14.29 | אלא |  |
| 4:14.30 | את |  |
| 4:14.31 | ה | ' |
| 4:14.32 | אלהיך |  |
| 4:14.33 | תירא | , |
| 4:14.34 | מאי |  |
| 4:14.35 | את |  |
| 4:14.36 | דא |  |
| 4:14.37 | דרגא |  |
| 4:14.38 | קדמאה |  |
| 4:14.39 | פירש |  |
| 4:14.40 | בגליון |  |
| 4:14.41 | מלכות |  |
| 4:14.42 | רמוזה |  |
| 4:14.43 | בתיבת |  |
| 4:14.44 | את |  |
| 4:14.45 | ה | ' |
| 4:14.46 | אלהיך |  |
| 4:14.47 | תירא | ) |
| 4:14.48 | אחר |  |
| 4:14.49 | דחלא |  |
| 4:14.50 | דקב"ה | , |
| 4:14.51 | ובגין |  |
| 4:14.52 | כך |  |
| 4:14.53 | כתיב |  |
| 4:14.54 | תירא | , |
| 4:14.55 | דתמן |  |
| 4:14.56 | בעי |  |
| 4:14.57 | בר |  |
| 4:14.58 | נש |  |
| 4:14.59 | לדחלא |  |
| 4:14.60 | קמי |  |
| 4:14.61 | מריה |  |
| 4:14.62 | בגין |  |
| 4:14.63 | דאיהו |  |
| 4:14.64 | בי |  |
| 4:14.65 | דינא | , |
| 4:14.66 | ואותו |  |
| 4:14.67 | תעבוד |  |
| 4:14.68 | דא |  |
| 4:14.69 | דרגא |  |
| 4:14.70 | עלאה |  |
| 4:14.71 | דקימא |  |
| 4:14.72 | על |  |
| 4:14.73 | האיי |  |
| 4:14.74 | דרגא |  |
| 4:14.75 | תתאה |  |
| 4:14.76 | ולא |  |
| 4:14.77 | מתפרשין |  |
| 4:14.78 | לעלמין | , |
| 4:14.79 | את |  |
| 4:14.80 | ואותו |  |
| 4:14.81 | דא |  |
| 4:14.82 | בדא |  |
| 4:14.83 | דבקין |  |
| 4:14.84 | ולא |  |
| 4:14.85 | אתפרשן | , |
| 4:14.86 | מאי |  |
| 4:14.87 | ואותו |  |
| 4:14.88 | דא |  |
| 4:14.89 | אתר |  |
| 4:14.90 | ברית |  |
| 4:14.91 | קדישא |  |
| 4:14.92 | אות |  |
| 4:14.93 | לעלמין | , |
| 4:14.94 | דהא |  |
| 4:14.95 | פולחנא |  |
| 4:14.96 | לא |  |
| 4:14.97 | שריא |  |
| 4:14.98 | באת |  |
| 4:14.99 | ולא |  |
| 4:14.100 | איהו |  |
| 4:14.101 | למיפלח |  |
| 4:14.102 | אלא |  |
| 4:14.103 | למדחל | , |
| 4:14.104 | אבל |  |
| 4:14.105 | פולחנא |  |
| 4:14.106 | איהו |  |
| 4:14.107 | לעילא | , |
| 4:14.108 | ובגין |  |
| 4:14.109 | כך |  |
| 4:14.110 | ואותו |  |
| 4:14.111 | תעבוד |  |
| 4:14.112 | פי | ' |
| 4:14.113 | אבל |  |
| 4:14.114 | העבודה |  |
| 4:14.115 | היא |  |
| 4:14.116 | לזעיר |  |
| 4:14.117 | אנפין |  |
| 4:14.118 | כמו |  |
| 4:14.119 | שנ | ' |
| 4:14.120 | ואותו |  |
| 4:14.121 | תעבוד |  |
| 4:14.122 | ותייחד |  |
| 4:14.123 | עמיה |  |
| 4:14.124 | שכנתיה |  |
| 4:14.125 | ועיין |  |
| 4:14.126 | בס | ' |
| 4:14.127 | המצות |  |
| 4:14.128 | בתפלת |  |
| 4:14.129 | אליהו |  |
| 4:14.130 | הנביא |  |
| 4:14.131 | מפורש |  |
| 4:14.132 | באורך | ) |
| 4:14.133 | ע"כ |  |
| 4:14.134 | לשון |  |
| 4:14.135 | הזהר |  |
| 4:14.136 | עם |  |
| 4:14.137 | פי | ' |
| 4:14.138 | הרש"ב |  |
| 4:14.139 | בגליון | . |
| 4:14.140 | וכונת |  |
| 4:14.141 | הזהר |  |
| 4:14.142 | לדרוש |  |
| 4:14.143 | מדלא |  |
| 4:14.144 | כתיב |  |
| 4:14.145 | לה | ' |
| 4:14.146 | אלהיך |  |
| 4:14.147 | תירא |  |
| 4:14.148 | ותעבוד | , |
| 4:14.149 | וכתיב |  |
| 4:14.150 | את |  |
| 4:14.151 | לרבות | , |
| 4:14.152 | ריבה |  |
| 4:14.153 | מדת |  |
| 4:14.154 | המלכות |  |
| 4:14.155 | שהיא |  |
| 4:14.156 | נוקבא |  |
| 4:14.157 | דזעיר |  |
| 4:14.158 | אנפין |  |
| 4:14.159 | ליראה | , |
| 4:14.160 | ומדכתיב |  |
| 4:14.161 | ואותו |  |
| 4:14.162 | תעבוד |  |
| 4:14.163 | מיעט |  |
| 4:14.164 | אותה |  |
| 4:14.165 | מעבודה |  |
| 4:14.166 | דדוקא |  |
| 4:14.167 | אותו |  |
| 4:14.168 | תעבוד |  |
| 4:14.169 | ולא |  |
| 4:14.170 | אותה |  |
| 4:14.171 | דהיינו |  |
| 4:14.172 | מלכות | ) |
| 4:14.173 | הרי |  |
| 4:14.174 | ריבה |  |
| 4:14.175 | אותה |  |
| 4:14.176 | ליראה |  |
| 4:14.177 | ומיעט |  |
| 4:14.178 | אותה |  |
| 4:14.179 | מעבודה |  |
| 4:14.180 | שהעבודה |  |
| 4:14.181 | לזעיר |  |
| 4:14.182 | אנפין |  |
| 4:14.183 | דוקא | . |


## Milchamot ha-Shem 4:15

Hebrew (Wikisource): כח) ומכאן יש להוכיח שאינה קבלה ממשה רבינו כמו שעלה על דעתך. שאם כן שהיא קבלה ממשה. למה פירש שמעון העמסוני או נחמיה העמסוני מלדרוש האתין שבתורה כשהגיע לאת ה' אלהיך תירא כדגר' בפ' כל שעה, תניא שמעון העמסוני ואמרי לה נחמיה העמסוני היה דורש כל אתין שבתורה, כיון שהגיע לאת ה' אלהיך תירא פירש אמרו לו תלמידיו ר' כל אתין שדרשת מה תהא עליהן. אמר להם כשם שקיבלתי שכר על הדרישה כך מקבל אני שכר על הפרישה, עד שבא ר' עקיבה ודרש את ה' אלהיך תירא לרבות תלמידי חכמים, ופי' רש"י ז"ל כיון שהגיע לאת ה' אלהיך תירא, אמר מה ארבה עמו לירא, פירש מכולן וחזר בו מכל רבויין שדרש, עד שבא ר' עקיבה ודרש את לרבות ת"ח שיהא מורא רבך כמורא שמים ע"כ לשון הגמ' ורש"י ז"ל. ואם כדבריך שהיא הלכה למשה מסיני. מה לו לשמעון העמסוני לפרוש? וכי לא ידע לומר לרבות אשתו של זעיר אנפין שהוא ה' אלהינו לפי דבריהם? וכמו שדרשו רז"ל בכתובות תניא כבד את אביך ואת אמך, את אביך זו אשת אביך ואת אמך זו בעל אמך, וא"ו יתירה לרבות אחיך הגדול ע"כ.

| id | word | punctuation after |
|---|---|---|
| 4:15.1 | כח | ) |
| 4:15.2 | ומכאן |  |
| 4:15.3 | יש |  |
| 4:15.4 | להוכיח |  |
| 4:15.5 | שאינה |  |
| 4:15.6 | קבלה |  |
| 4:15.7 | ממשה |  |
| 4:15.8 | רבינו |  |
| 4:15.9 | כמו |  |
| 4:15.10 | שעלה |  |
| 4:15.11 | על |  |
| 4:15.12 | דעתך | . |
| 4:15.13 | שאם |  |
| 4:15.14 | כן |  |
| 4:15.15 | שהיא |  |
| 4:15.16 | קבלה |  |
| 4:15.17 | ממשה | . |
| 4:15.18 | למה |  |
| 4:15.19 | פירש |  |
| 4:15.20 | שמעון |  |
| 4:15.21 | העמסוני |  |
| 4:15.22 | או |  |
| 4:15.23 | נחמיה |  |
| 4:15.24 | העמסוני |  |
| 4:15.25 | מלדרוש |  |
| 4:15.26 | האתין |  |
| 4:15.27 | שבתורה |  |
| 4:15.28 | כשהגיע |  |
| 4:15.29 | לאת |  |
| 4:15.30 | ה | ' |
| 4:15.31 | אלהיך |  |
| 4:15.32 | תירא |  |
| 4:15.33 | כדגר | ' |
| 4:15.34 | בפ | ' |
| 4:15.35 | כל |  |
| 4:15.36 | שעה | , |
| 4:15.37 | תניא |  |
| 4:15.38 | שמעון |  |
| 4:15.39 | העמסוני |  |
| 4:15.40 | ואמרי |  |
| 4:15.41 | לה |  |
| 4:15.42 | נחמיה |  |
| 4:15.43 | העמסוני |  |
| 4:15.44 | היה |  |
| 4:15.45 | דורש |  |
| 4:15.46 | כל |  |
| 4:15.47 | אתין |  |
| 4:15.48 | שבתורה | , |
| 4:15.49 | כיון |  |
| 4:15.50 | שהגיע |  |
| 4:15.51 | לאת |  |
| 4:15.52 | ה | ' |
| 4:15.53 | אלהיך |  |
| 4:15.54 | תירא |  |
| 4:15.55 | פירש |  |
| 4:15.56 | אמרו |  |
| 4:15.57 | לו |  |
| 4:15.58 | תלמידיו |  |
| 4:15.59 | ר | ' |
| 4:15.60 | כל |  |
| 4:15.61 | אתין |  |
| 4:15.62 | שדרשת |  |
| 4:15.63 | מה |  |
| 4:15.64 | תהא |  |
| 4:15.65 | עליהן | . |
| 4:15.66 | אמר |  |
| 4:15.67 | להם |  |
| 4:15.68 | כשם |  |
| 4:15.69 | שקיבלתי |  |
| 4:15.70 | שכר |  |
| 4:15.71 | על |  |
| 4:15.72 | הדרישה |  |
| 4:15.73 | כך |  |
| 4:15.74 | מקבל |  |
| 4:15.75 | אני |  |
| 4:15.76 | שכר |  |
| 4:15.77 | על |  |
| 4:15.78 | הפרישה | , |
| 4:15.79 | עד |  |
| 4:15.80 | שבא |  |
| 4:15.81 | ר | ' |
| 4:15.82 | עקיבה |  |
| 4:15.83 | ודרש |  |
| 4:15.84 | את |  |
| 4:15.85 | ה | ' |
| 4:15.86 | אלהיך |  |
| 4:15.87 | תירא |  |
| 4:15.88 | לרבות |  |
| 4:15.89 | תלמידי |  |
| 4:15.90 | חכמים | , |
| 4:15.91 | ופי | ' |
| 4:15.92 | רש"י |  |
| 4:15.93 | ז"ל |  |
| 4:15.94 | כיון |  |
| 4:15.95 | שהגיע |  |
| 4:15.96 | לאת |  |
| 4:15.97 | ה | ' |
| 4:15.98 | אלהיך |  |
| 4:15.99 | תירא | , |
| 4:15.100 | אמר |  |
| 4:15.101 | מה |  |
| 4:15.102 | ארבה |  |
| 4:15.103 | עמו |  |
| 4:15.104 | לירא | , |
| 4:15.105 | פירש |  |
| 4:15.106 | מכולן |  |
| 4:15.107 | וחזר |  |
| 4:15.108 | בו |  |
| 4:15.109 | מכל |  |
| 4:15.110 | רבויין |  |
| 4:15.111 | שדרש | , |
| 4:15.112 | עד |  |
| 4:15.113 | שבא |  |
| 4:15.114 | ר | ' |
| 4:15.115 | עקיבה |  |
| 4:15.116 | ודרש |  |
| 4:15.117 | את |  |
| 4:15.118 | לרבות |  |
| 4:15.119 | ת"ח |  |
| 4:15.120 | שיהא |  |
| 4:15.121 | מורא |  |
| 4:15.122 | רבך |  |
| 4:15.123 | כמורא |  |
| 4:15.124 | שמים |  |
| 4:15.125 | ע"כ |  |
| 4:15.126 | לשון |  |
| 4:15.127 | הגמ | ' |
| 4:15.128 | ורש"י |  |
| 4:15.129 | ז"ל | . |
| 4:15.130 | ואם |  |
| 4:15.131 | כדבריך |  |
| 4:15.132 | שהיא |  |
| 4:15.133 | הלכה |  |
| 4:15.134 | למשה |  |
| 4:15.135 | מסיני | . |
| 4:15.136 | מה |  |
| 4:15.137 | לו |  |
| 4:15.138 | לשמעון |  |
| 4:15.139 | העמסוני |  |
| 4:15.140 | לפרוש | ? |
| 4:15.141 | וכי |  |
| 4:15.142 | לא |  |
| 4:15.143 | ידע |  |
| 4:15.144 | לומר |  |
| 4:15.145 | לרבות |  |
| 4:15.146 | אשתו |  |
| 4:15.147 | של |  |
| 4:15.148 | זעיר |  |
| 4:15.149 | אנפין |  |
| 4:15.150 | שהוא |  |
| 4:15.151 | ה | ' |
| 4:15.152 | אלהינו |  |
| 4:15.153 | לפי |  |
| 4:15.154 | דבריהם | ? |
| 4:15.155 | וכמו |  |
| 4:15.156 | שדרשו |  |
| 4:15.157 | רז"ל |  |
| 4:15.158 | בכתובות |  |
| 4:15.159 | תניא |  |
| 4:15.160 | כבד |  |
| 4:15.161 | את |  |
| 4:15.162 | אביך |  |
| 4:15.163 | ואת |  |
| 4:15.164 | אמך | , |
| 4:15.165 | את |  |
| 4:15.166 | אביך |  |
| 4:15.167 | זו |  |
| 4:15.168 | אשת |  |
| 4:15.169 | אביך |  |
| 4:15.170 | ואת |  |
| 4:15.171 | אמך |  |
| 4:15.172 | זו |  |
| 4:15.173 | בעל |  |
| 4:15.174 | אמך | , |
| 4:15.175 | וא"ו |  |
| 4:15.176 | יתירה |  |
| 4:15.177 | לרבות |  |
| 4:15.178 | אחיך |  |
| 4:15.179 | הגדול |  |
| 4:15.180 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

