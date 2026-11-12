# Review pack: milhamot-hashem perek-3-01 (Milchamot ha-Shem 3:1–3:2)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-01`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

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
- סבא **דסבין** → "the Elder of Elders".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
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
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- אחד ... **ואחד** of two kinds of case → "alike".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **סכלות** והוללות → "folly and madness".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 3:1

Hebrew (Wikisource): ב) וכ"ש שלא נאמין לו באמרו שנגלה אליו אליהו או סבא דסבין, שהוא עתיקא קדישא הנקרא כתר עליון, דודאי נביא שקר הוא ובמיתתו בחנק. כמבואה בדברי הרמב"ם בפי' המשנה בהקדמתו לסדר זרעים. ובגמ' פסחים (לד) גרס' אמתר הסח הדעת ר' יוחנן אמר פסול טומאה, שאם יבוא אליהו ויטהרנה ויאמר שלא נגעה טומאה בהן שומעין לו, ור"ל אמר פסול הגוף הוי, שאם יבוא אליהו ויטהרנה אין שומעין לו עכ"ל, ופי' רשי ז"ל פיסול הגוף מעלה היא בקדשים היא שנפסלין בכך, ושם פיסול הוא לעצמו, ואפי' אם יבוא אליהו ויאמר שלא נטמאו כל רותן הימים שהסיח דעתו מהן אין שומעין לו ע"כ, אלמלא כל הבא לשנות ולהתיר את האסור, אפי' איסור דרבנן אין שומעין לו.

| id | word | punctuation after |
|---|---|---|
| 3:1.1 | ב | ) |
| 3:1.2 | וכ"ש |  |
| 3:1.3 | שלא |  |
| 3:1.4 | נאמין |  |
| 3:1.5 | לו |  |
| 3:1.6 | באמרו |  |
| 3:1.7 | שנגלה |  |
| 3:1.8 | אליו |  |
| 3:1.9 | אליהו |  |
| 3:1.10 | או |  |
| 3:1.11 | סבא |  |
| 3:1.12 | דסבין | , |
| 3:1.13 | שהוא |  |
| 3:1.14 | עתיקא |  |
| 3:1.15 | קדישא |  |
| 3:1.16 | הנקרא |  |
| 3:1.17 | כתר |  |
| 3:1.18 | עליון | , |
| 3:1.19 | דודאי |  |
| 3:1.20 | נביא |  |
| 3:1.21 | שקר |  |
| 3:1.22 | הוא |  |
| 3:1.23 | ובמיתתו |  |
| 3:1.24 | בחנק | . |
| 3:1.25 | כמבואה |  |
| 3:1.26 | בדברי |  |
| 3:1.27 | הרמב"ם |  |
| 3:1.28 | בפי | ' |
| 3:1.29 | המשנה |  |
| 3:1.30 | בהקדמתו |  |
| 3:1.31 | לסדר |  |
| 3:1.32 | זרעים | . |
| 3:1.33 | ובגמ | ' |
| 3:1.34 | פסחים |  |
| 3:1.35 | לד | ) |
| 3:1.36 | גרס | ' |
| 3:1.37 | אמתר |  |
| 3:1.38 | הסח |  |
| 3:1.39 | הדעת |  |
| 3:1.40 | ר | ' |
| 3:1.41 | יוחנן |  |
| 3:1.42 | אמר |  |
| 3:1.43 | פסול |  |
| 3:1.44 | טומאה | , |
| 3:1.45 | שאם |  |
| 3:1.46 | יבוא |  |
| 3:1.47 | אליהו |  |
| 3:1.48 | ויטהרנה |  |
| 3:1.49 | ויאמר |  |
| 3:1.50 | שלא |  |
| 3:1.51 | נגעה |  |
| 3:1.52 | טומאה |  |
| 3:1.53 | בהן |  |
| 3:1.54 | שומעין |  |
| 3:1.55 | לו | , |
| 3:1.56 | ור"ל |  |
| 3:1.57 | אמר |  |
| 3:1.58 | פסול |  |
| 3:1.59 | הגוף |  |
| 3:1.60 | הוי | , |
| 3:1.61 | שאם |  |
| 3:1.62 | יבוא |  |
| 3:1.63 | אליהו |  |
| 3:1.64 | ויטהרנה |  |
| 3:1.65 | אין |  |
| 3:1.66 | שומעין |  |
| 3:1.67 | לו |  |
| 3:1.68 | עכ"ל | , |
| 3:1.69 | ופי | ' |
| 3:1.70 | רשי |  |
| 3:1.71 | ז"ל |  |
| 3:1.72 | פיסול |  |
| 3:1.73 | הגוף |  |
| 3:1.74 | מעלה |  |
| 3:1.75 | היא |  |
| 3:1.76 | בקדשים |  |
| 3:1.77 | היא |  |
| 3:1.78 | שנפסלין |  |
| 3:1.79 | בכך | , |
| 3:1.80 | ושם |  |
| 3:1.81 | פיסול |  |
| 3:1.82 | הוא |  |
| 3:1.83 | לעצמו | , |
| 3:1.84 | ואפי | ' |
| 3:1.85 | אם |  |
| 3:1.86 | יבוא |  |
| 3:1.87 | אליהו |  |
| 3:1.88 | ויאמר |  |
| 3:1.89 | שלא |  |
| 3:1.90 | נטמאו |  |
| 3:1.91 | כל |  |
| 3:1.92 | רותן |  |
| 3:1.93 | הימים |  |
| 3:1.94 | שהסיח |  |
| 3:1.95 | דעתו |  |
| 3:1.96 | מהן |  |
| 3:1.97 | אין |  |
| 3:1.98 | שומעין |  |
| 3:1.99 | לו |  |
| 3:1.100 | ע"כ | , |
| 3:1.101 | אלמלא |  |
| 3:1.102 | כל |  |
| 3:1.103 | הבא |  |
| 3:1.104 | לשנות |  |
| 3:1.105 | ולהתיר |  |
| 3:1.106 | את |  |
| 3:1.107 | האסור | , |
| 3:1.108 | אפי | ' |
| 3:1.109 | איסור |  |
| 3:1.110 | דרבנן |  |
| 3:1.111 | אין |  |
| 3:1.112 | שומעין |  |
| 3:1.113 | לו | . |


### Translation under review (3:1)

English: §2. And all the more so we should not believe him when he says that Elijah was revealed to him, or the Elder of Elders, who is Atika Kadisha, which is called Keter Elyon (the Supreme Crown), for surely he is a false prophet, and he is put to death by strangling. As is shown in the words of the Rambam in his commentary on the Mishnah, in his introduction to the Order of Zera'im. And in the Gemara, Pesachim 34, it reads: "Amtar, the mind was distracted. Rabbi Yochanan said: a disqualification by impurity, that if Elijah should come and declare it clean and say that impurity had not touched them, we listen to him, and Resh Lakish said: a disqualification of the body, is that if Elijah should come and declare it clean, we do not listen to him." Thus far his words. And Rashi, of blessed memory, explains: "bodily disqualification is a merit in sacred offerings, in that they are disqualified by it; and the disqualification there is in itself; even if Elijah should come and say that all their [unclear] days were not defiled, since his mind was distracted from them, we do not listen to him." Thus far. Indeed, everyone who comes to change and to permit the forbidden, even a rabbinic prohibition, we do not listen to him.

Literal: section 2 and all the more so that not we believe to him in his saying that he was revealed to him Elijah or the Elder of Elders who is Atika Kadisha which is called Keter Elyon surely prophet false he and by his death by strangling as is shown in the words of the Rambam in his commentary on the Mishnah in his introduction to the Order Zera'im and in the Gemara Pesachim 34 it reads Amtar was distracted the mind Rabbi Yochanan said disqualification by impurity that if should come Elijah and declare it clean and say that not touched impurity them we listen to him and Resh Lakish said disqualification of the body is that if should come Elijah and declare it clean we do not listen to him thus far his words and explains Rashi of blessed memory disqualification of the body a merit is in sacred offerings that they are disqualified by it and there disqualification is in itself and even if should come Elijah and say that not were defiled all their [unclear] days since he distracted his mind from them we do not listen to him thus far indeed (literally: were it not) everyone who comes to change and to permit (object marker) the forbidden even a prohibition rabbinic we do not listen to him

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:1.1 | section 2 |  |  | the label ב is the numeral 2 (section 2) |
| 3:1.2 | and all the more so |  |  | כ"ש = כל שכן, 'all the more so' |
| 3:1.3 | that not |  |  |  |
| 3:1.4 | we believe |  |  |  |
| 3:1.5 | to him |  |  |  |
| 3:1.6 | in his saying |  |  |  |
| 3:1.7 | that he was revealed |  |  |  |
| 3:1.8 | to him |  |  |  |
| 3:1.9 | Elijah |  |  |  |
| 3:1.10 | or |  |  |  |
| 3:1.11 | the Elder |  |  | Aramaic סבא, 'elder'; the title is the author's, not a name |
| 3:1.12 | of Elders |  |  |  |
| 3:1.13 | who is |  |  |  |
| 3:1.14 | Atika |  |  | Aramaic 'holy ancient one'; the Hebrew title is kept |
| 3:1.15 | Kadisha |  |  |  |
| 3:1.16 | which is called |  |  |  |
| 3:1.17 | Keter |  |  | כתר עליון, 'Supreme Crown' |
| 3:1.18 | Elyon |  |  |  |
| 3:1.19 | surely |  |  |  |
| 3:1.20 | prophet |  |  |  |
| 3:1.21 | false |  |  |  |
| 3:1.22 | he |  |  |  |
| 3:1.23 | and by his death |  |  |  |
| 3:1.24 | by strangling |  |  |  |
| 3:1.25 | as is shown |  |  |  |
| 3:1.26 | in the words of |  |  |  |
| 3:1.27 | the Rambam |  |  |  |
| 3:1.28 | in his commentary |  |  |  |
| 3:1.29 | on the Mishnah |  |  |  |
| 3:1.30 | in his introduction |  |  |  |
| 3:1.31 | to the Order |  |  |  |
| 3:1.32 | Zera'im |  |  |  |
| 3:1.33 | and in the Gemara |  |  |  |
| 3:1.34 | Pesachim |  |  |  |
| 3:1.35 | 34 |  |  |  |
| 3:1.36 | it reads |  |  |  |
| 3:1.37 | Amtar |  |  | amtar: Aramaic word, sense uncertain in this place; see apparatus |
| 3:1.38 | was distracted |  |  |  |
| 3:1.39 | the mind |  |  |  |
| 3:1.40 | Rabbi |  |  |  |
| 3:1.41 | Yochanan |  |  |  |
| 3:1.42 | said |  |  |  |
| 3:1.43 | disqualification |  |  |  |
| 3:1.44 | by impurity |  |  |  |
| 3:1.45 | that if |  |  |  |
| 3:1.46 | should come |  |  |  |
| 3:1.47 | Elijah |  |  |  |
| 3:1.48 | and declare it clean |  |  |  |
| 3:1.49 | and say |  |  |  |
| 3:1.50 | that not |  |  |  |
| 3:1.51 | touched |  |  |  |
| 3:1.52 | impurity |  |  |  |
| 3:1.53 | them |  |  |  |
| 3:1.54 | we listen |  |  |  |
| 3:1.55 | to him |  |  |  |
| 3:1.56 | and Resh Lakish |  |  | ר"ל = Resh Lakish |
| 3:1.57 | said |  |  |  |
| 3:1.58 | disqualification |  |  |  |
| 3:1.59 | of the body |  |  |  |
| 3:1.60 | is |  |  |  |
| 3:1.61 | that if |  |  |  |
| 3:1.62 | should come |  |  |  |
| 3:1.63 | Elijah |  |  |  |
| 3:1.64 | and declare it clean |  |  |  |
| 3:1.65 | we do not |  |  |  |
| 3:1.66 | listen |  |  |  |
| 3:1.67 | to him |  |  |  |
| 3:1.68 | thus far his words |  |  |  |
| 3:1.69 | and explains |  |  |  |
| 3:1.70 | Rashi |  |  |  |
| 3:1.71 | of blessed memory |  |  |  |
| 3:1.72 | disqualification |  |  |  |
| 3:1.73 | of the body |  |  |  |
| 3:1.74 | a merit |  |  |  |
| 3:1.75 | is |  |  |  |
| 3:1.76 | in sacred offerings |  |  |  |
| 3:1.77 | that |  |  |  |
| 3:1.78 | they are disqualified |  |  |  |
| 3:1.79 | by it |  |  |  |
| 3:1.80 | and there |  |  |  |
| 3:1.81 | disqualification |  |  |  |
| 3:1.82 | is |  |  |  |
| 3:1.83 | in itself |  |  |  |
| 3:1.84 | and even |  |  |  |
| 3:1.85 | if |  |  |  |
| 3:1.86 | should come |  |  |  |
| 3:1.87 | Elijah |  |  |  |
| 3:1.88 | and say |  |  |  |
| 3:1.89 | that not |  |  |  |
| 3:1.90 | were defiled |  |  |  |
| 3:1.91 | all |  |  |  |
| 3:1.92 | their [unclear] |  |  | word not clear to me in this sense; see apparatus |
| 3:1.93 | days |  |  |  |
| 3:1.94 | since he distracted |  |  |  |
| 3:1.95 | his mind |  |  |  |
| 3:1.96 | from them |  |  |  |
| 3:1.97 | we do not |  |  |  |
| 3:1.98 | listen |  |  |  |
| 3:1.99 | to him |  |  |  |
| 3:1.100 | thus far |  |  |  |
| 3:1.101 | indeed (literally: were it not) |  |  | literally 'were it not'; taken as the opening of the author's own conclusion |
| 3:1.102 | everyone |  |  |  |
| 3:1.103 | who comes |  |  |  |
| 3:1.104 | to change |  |  |  |
| 3:1.105 | and to permit |  |  |  |
| 3:1.106 | (object marker) |  |  | את: object marker, no English counterpart |
| 3:1.107 | the forbidden |  |  |  |
| 3:1.108 | even |  |  |  |
| 3:1.109 | a prohibition |  |  |  |
| 3:1.110 | rabbinic |  |  |  |
| 3:1.111 | we do not |  |  |  |
| 3:1.112 | listen |  |  |  |
| 3:1.113 | to him |  |  |  |

Phrases:
- 3:1.1 3:1.2 | §2. And all the more so | section 2 and all the more so
- 3:1.3 3:1.4 3:1.5 3:1.6 3:1.7 3:1.8 3:1.9 | we should not believe him when he says that Elijah was revealed to him, | that not we believe to him in his saying that he was revealed to him Elijah
- 3:1.10 3:1.11 3:1.12 | or the Elder of Elders, | or the Elder of Elders
- 3:1.13 3:1.14 3:1.15 3:1.16 | who is Atika Kadisha, which is called | who is Atika Kadisha which is called
- 3:1.17 3:1.18 | Keter Elyon (the Supreme Crown), | Keter Elyon
- 3:1.19 3:1.20 3:1.21 3:1.22 | for surely he is a false prophet, | surely prophet false he
- 3:1.23 3:1.24 | and he is put to death by strangling. | and by his death by strangling
- 3:1.25 3:1.26 | As is shown in the words of | as is shown in the words of
- 3:1.27 | the Rambam | the Rambam
- 3:1.28 3:1.29 3:1.30 3:1.31 3:1.32 | in his commentary on the Mishnah, in his introduction to the Order of Zera'im. | in his commentary on the Mishnah in his introduction to the Order Zera'im
- 3:1.33 3:1.34 3:1.35 3:1.36 | And in the Gemara, Pesachim 34, it reads: | and in the Gemara Pesachim 34 it reads
- 3:1.37 3:1.38 3:1.39 | "Amtar, the mind was distracted. | Amtar was distracted the mind
- 3:1.40 3:1.41 3:1.42 | Rabbi Yochanan said: | Rabbi Yochanan said
- 3:1.43 3:1.44 | a disqualification by impurity, | disqualification by impurity
- 3:1.45 3:1.46 3:1.47 3:1.48 3:1.49 3:1.50 3:1.51 3:1.52 3:1.53 3:1.54 3:1.55 | that if Elijah should come and declare it clean and say that impurity had not touched them, we listen to him, | that if should come Elijah and declare it clean and say that not touched impurity them we listen to him
- 3:1.56 3:1.57 | and Resh Lakish said: | and Resh Lakish said
- 3:1.58 3:1.59 | a disqualification of the body, | disqualification of the body
- 3:1.60 3:1.61 3:1.62 3:1.63 3:1.64 3:1.65 3:1.66 3:1.67 3:1.68 | is that if Elijah should come and declare it clean, we do not listen to him." Thus far his words. | is that if should come Elijah and declare it clean we do not listen to him thus far his words
- 3:1.69 3:1.70 3:1.71 | And Rashi, of blessed memory, explains: | and explains Rashi of blessed memory
- 3:1.72 3:1.73 3:1.74 3:1.75 3:1.76 3:1.77 3:1.78 3:1.79 | "bodily disqualification is a merit in sacred offerings, in that they are disqualified by it; | disqualification of the body a merit is in sacred offerings that they are disqualified by it
- 3:1.80 3:1.81 3:1.82 3:1.83 | and the disqualification there is in itself; | and there disqualification is in itself
- 3:1.84 3:1.85 3:1.86 3:1.87 3:1.88 3:1.89 3:1.90 3:1.91 3:1.92 3:1.93 3:1.94 3:1.95 3:1.96 3:1.97 3:1.98 3:1.99 3:1.100 | even if Elijah should come and say that all their [unclear] days were not defiled, since his mind was distracted from them, we do not listen to him." Thus far. | and even if should come Elijah and say that not were defiled all their [unclear] days since he distracted his mind from them we do not listen to him thus far
- 3:1.101 3:1.102 3:1.103 3:1.104 3:1.105 3:1.106 3:1.107 3:1.108 3:1.109 3:1.110 3:1.111 3:1.112 3:1.113 | Indeed, everyone who comes to change and to permit the forbidden, even a rabbinic prohibition, we do not listen to him. | indeed (literally: were it not) everyone who comes to change and to permit (object marker) the forbidden even a prohibition rabbinic we do not listen to him

Grammar notes:
- שלא נאמין לו: 'that we should not believe him', a clause with שלא and the imperfect after כ"ש, 'all the more so'.
- סבא דסבין is Aramaic: 'the Elder of Elders', with the Aramaic relative ד and the title given as a name-like epithet.
- The Talmud passage is quoted with 'גרס', 'it reads'; the closing 'עכ"ל' ends it, and 'ע"כ' ends Rashi's words after it.
- ויטהרנה: the verb with the object suffix -נה; the English 'declare it clean' takes the suffix as 'it'.

Translator's note: The paragraph has three parts: the refusal to believe the claimed revelation, a Talmud passage with Resh Lakish's view, and Rashi's explanation with the conclusion. The Hebrew is the Wikisource base text.

Commentary rulings:
- apparatus: Base text: Wikisource edition, translated as base text and flagged until the 1931 scan is collated (decisions, Q2).
- apparatus: Rambam: his commentary on the Mishnah, introduction to Zera'im, as the author gives it (no chapter or page).
- apparatus: Pesachim (לד): the author's '34' as written; the quotation as he gives it, not checked against the Talmud text.
- apparatus: amtar: Aramaic word of uncertain sense here; the English reads it as the word the Gemara uses to introduce the case, and the alternative is to leave it untranslated.
- apparatus: רותן: word not established in this sense; translated 'their [unclear]' and marked in the token notes.

## Milchamot ha-Shem 3:2

Hebrew (Wikisource): ג) ובאמת שהחלת דברי פיך סכלות והוללות רעה שאמרת איך מלאני לבי לדרוש ולחקור וכו' ולמה לא? אדרבא איפכא מסתברא (כיון שאומרים לעבוד לז"א שהוא נברא יושר לבב דף ג' ע"ב ועוז לאלוהים). חוב גמור עלינו ועל כל אחד ואחד בישראל לדעת את ה' שהוא בורא ולא נברא ולעבדו ככתוב וידעת היום והשבות אל לבבך. וידעת כי ה' אלהיך הוא האלהים וכו' ואומר דע את אלהי אביך ועבדהו. והתירו חכמים לאפרושי מאיסורא אפי' בבית הכסא ובבית המרחץ שאסור להרהר בהם בדברי תורה. ולאפרושי מאיסורא שרי. ועל תמיהתך איך יטעו? אין זה פלא. ולא עדיפי כת המקובלים האחרונים אשר הלכו לתומם אחרי צו של הפילוסוף מחבר הזהר שתלאו ברשב"י ור' אלעזר בנו וחביריו מסנהדרי גדולה היושבים בלשכת הגזית שעליהם אמר הכתוב כי יפלא ממך דבר למשפט וכו' וקמת ועלית אל המקום אשר יבחר ה' אלוהיך וכו' ועשית על פי הדבר אשר יגידו לך וכו'. ועם כל זה אמרה תורה שאם שגגו ב"ד הגדול בהוראה ועשו העם על פיהם מביאים קרבן על שגגם. אם בע"ז שגגו והורו מביאים פר לעולה ושעיר לחטאת מכל שבט ושבט. והוא האמור בפרשת שלח והיה אם מעיני העדה נעשתה לשגגה, ואם בשאר כריתות שגגו והורו מביא כל שבט ושבט פר לחטאת והוא האמור בפרשת ויקרא. ואם כל עדת ישראל ישגו ונעלם דבר מעיני הקהל.

| id | word | punctuation after |
|---|---|---|
| 3:2.1 | ג | ) |
| 3:2.2 | ובאמת |  |
| 3:2.3 | שהחלת |  |
| 3:2.4 | דברי |  |
| 3:2.5 | פיך |  |
| 3:2.6 | סכלות |  |
| 3:2.7 | והוללות |  |
| 3:2.8 | רעה |  |
| 3:2.9 | שאמרת |  |
| 3:2.10 | איך |  |
| 3:2.11 | מלאני |  |
| 3:2.12 | לבי |  |
| 3:2.13 | לדרוש |  |
| 3:2.14 | ולחקור |  |
| 3:2.15 | וכו | ' |
| 3:2.16 | ולמה |  |
| 3:2.17 | לא | ? |
| 3:2.18 | אדרבא |  |
| 3:2.19 | איפכא |  |
| 3:2.20 | מסתברא |  |
| 3:2.21 | כיון |  |
| 3:2.22 | שאומרים |  |
| 3:2.23 | לעבוד |  |
| 3:2.24 | לז"א |  |
| 3:2.25 | שהוא |  |
| 3:2.26 | נברא |  |
| 3:2.27 | יושר |  |
| 3:2.28 | לבב |  |
| 3:2.29 | דף |  |
| 3:2.30 | ג | ' |
| 3:2.31 | ע"ב |  |
| 3:2.32 | ועוז |  |
| 3:2.33 | לאלוהים | ). |
| 3:2.34 | חוב |  |
| 3:2.35 | גמור |  |
| 3:2.36 | עלינו |  |
| 3:2.37 | ועל |  |
| 3:2.38 | כל |  |
| 3:2.39 | אחד |  |
| 3:2.40 | ואחד |  |
| 3:2.41 | בישראל |  |
| 3:2.42 | לדעת |  |
| 3:2.43 | את |  |
| 3:2.44 | ה | ' |
| 3:2.45 | שהוא |  |
| 3:2.46 | בורא |  |
| 3:2.47 | ולא |  |
| 3:2.48 | נברא |  |
| 3:2.49 | ולעבדו |  |
| 3:2.50 | ככתוב |  |
| 3:2.51 | וידעת |  |
| 3:2.52 | היום |  |
| 3:2.53 | והשבות |  |
| 3:2.54 | אל |  |
| 3:2.55 | לבבך | . |
| 3:2.56 | וידעת |  |
| 3:2.57 | כי |  |
| 3:2.58 | ה | ' |
| 3:2.59 | אלהיך |  |
| 3:2.60 | הוא |  |
| 3:2.61 | האלהים |  |
| 3:2.62 | וכו | ' |
| 3:2.63 | ואומר |  |
| 3:2.64 | דע |  |
| 3:2.65 | את |  |
| 3:2.66 | אלהי |  |
| 3:2.67 | אביך |  |
| 3:2.68 | ועבדהו | . |
| 3:2.69 | והתירו |  |
| 3:2.70 | חכמים |  |
| 3:2.71 | לאפרושי |  |
| 3:2.72 | מאיסורא |  |
| 3:2.73 | אפי | ' |
| 3:2.74 | בבית |  |
| 3:2.75 | הכסא |  |
| 3:2.76 | ובבית |  |
| 3:2.77 | המרחץ |  |
| 3:2.78 | שאסור |  |
| 3:2.79 | להרהר |  |
| 3:2.80 | בהם |  |
| 3:2.81 | בדברי |  |
| 3:2.82 | תורה | . |
| 3:2.83 | ולאפרושי |  |
| 3:2.84 | מאיסורא |  |
| 3:2.85 | שרי | . |
| 3:2.86 | ועל |  |
| 3:2.87 | תמיהתך |  |
| 3:2.88 | איך |  |
| 3:2.89 | יטעו | ? |
| 3:2.90 | אין |  |
| 3:2.91 | זה |  |
| 3:2.92 | פלא | . |
| 3:2.93 | ולא |  |
| 3:2.94 | עדיפי |  |
| 3:2.95 | כת |  |
| 3:2.96 | המקובלים |  |
| 3:2.97 | האחרונים |  |
| 3:2.98 | אשר |  |
| 3:2.99 | הלכו |  |
| 3:2.100 | לתומם |  |
| 3:2.101 | אחרי |  |
| 3:2.102 | צו |  |
| 3:2.103 | של |  |
| 3:2.104 | הפילוסוף |  |
| 3:2.105 | מחבר |  |
| 3:2.106 | הזהר |  |
| 3:2.107 | שתלאו |  |
| 3:2.108 | ברשב"י |  |
| 3:2.109 | ור | ' |
| 3:2.110 | אלעזר |  |
| 3:2.111 | בנו |  |
| 3:2.112 | וחביריו |  |
| 3:2.113 | מסנהדרי |  |
| 3:2.114 | גדולה |  |
| 3:2.115 | היושבים |  |
| 3:2.116 | בלשכת |  |
| 3:2.117 | הגזית |  |
| 3:2.118 | שעליהם |  |
| 3:2.119 | אמר |  |
| 3:2.120 | הכתוב |  |
| 3:2.121 | כי |  |
| 3:2.122 | יפלא |  |
| 3:2.123 | ממך |  |
| 3:2.124 | דבר |  |
| 3:2.125 | למשפט |  |
| 3:2.126 | וכו | ' |
| 3:2.127 | וקמת |  |
| 3:2.128 | ועלית |  |
| 3:2.129 | אל |  |
| 3:2.130 | המקום |  |
| 3:2.131 | אשר |  |
| 3:2.132 | יבחר |  |
| 3:2.133 | ה | ' |
| 3:2.134 | אלוהיך |  |
| 3:2.135 | וכו | ' |
| 3:2.136 | ועשית |  |
| 3:2.137 | על |  |
| 3:2.138 | פי |  |
| 3:2.139 | הדבר |  |
| 3:2.140 | אשר |  |
| 3:2.141 | יגידו |  |
| 3:2.142 | לך |  |
| 3:2.143 | וכו | '. |
| 3:2.144 | ועם |  |
| 3:2.145 | כל |  |
| 3:2.146 | זה |  |
| 3:2.147 | אמרה |  |
| 3:2.148 | תורה |  |
| 3:2.149 | שאם |  |
| 3:2.150 | שגגו |  |
| 3:2.151 | ב"ד |  |
| 3:2.152 | הגדול |  |
| 3:2.153 | בהוראה |  |
| 3:2.154 | ועשו |  |
| 3:2.155 | העם |  |
| 3:2.156 | על |  |
| 3:2.157 | פיהם |  |
| 3:2.158 | מביאים |  |
| 3:2.159 | קרבן |  |
| 3:2.160 | על |  |
| 3:2.161 | שגגם | . |
| 3:2.162 | אם |  |
| 3:2.163 | בע"ז |  |
| 3:2.164 | שגגו |  |
| 3:2.165 | והורו |  |
| 3:2.166 | מביאים |  |
| 3:2.167 | פר |  |
| 3:2.168 | לעולה |  |
| 3:2.169 | ושעיר |  |
| 3:2.170 | לחטאת |  |
| 3:2.171 | מכל |  |
| 3:2.172 | שבט |  |
| 3:2.173 | ושבט | . |
| 3:2.174 | והוא |  |
| 3:2.175 | האמור |  |
| 3:2.176 | בפרשת |  |
| 3:2.177 | שלח |  |
| 3:2.178 | והיה |  |
| 3:2.179 | אם |  |
| 3:2.180 | מעיני |  |
| 3:2.181 | העדה |  |
| 3:2.182 | נעשתה |  |
| 3:2.183 | לשגגה | , |
| 3:2.184 | ואם |  |
| 3:2.185 | בשאר |  |
| 3:2.186 | כריתות |  |
| 3:2.187 | שגגו |  |
| 3:2.188 | והורו |  |
| 3:2.189 | מביא |  |
| 3:2.190 | כל |  |
| 3:2.191 | שבט |  |
| 3:2.192 | ושבט |  |
| 3:2.193 | פר |  |
| 3:2.194 | לחטאת |  |
| 3:2.195 | והוא |  |
| 3:2.196 | האמור |  |
| 3:2.197 | בפרשת |  |
| 3:2.198 | ויקרא | . |
| 3:2.199 | ואם |  |
| 3:2.200 | כל |  |
| 3:2.201 | עדת |  |
| 3:2.202 | ישראל |  |
| 3:2.203 | ישגו |  |
| 3:2.204 | ונעלם |  |
| 3:2.205 | דבר |  |
| 3:2.206 | מעיני |  |
| 3:2.207 | הקהל | . |


### Translation under review (3:2)

English: §3. And indeed, that you have begun the words of your mouth, folly and madness, evil that you said, "How my heart urged me to inquire and to examine, etc." And why not? On the contrary, the reverse is more reasonable (since they say to worship Ze'ir Anpin, who is created, upright of heart, folio 3b, and strength to God). It is a complete obligation upon us and upon each and every one of Israel to know the LORD, that He is Creator and not created, and to worship Him, as it is written: "And you shall know this day and take it to your heart." "And you shall know that the LORD your God is God, etc." And it says: "Know the God of your father and worship Him." The Sages permitted abstaining from the forbidden, even in the privy and in the bathhouse, where it is forbidden to meditate on them in words of Torah. but to keep away from the forbidden is permitted. As to your astonishment, how could they err? There is no wonder in this. Nor is the sect of the later Kabbalists any better, those who went in their simplicity after the command of the philosopher, the author of the Zohar who attributed it to Rabbi Shimon bar Yochai, and Rabbi Eleazar his son, and their companions of the Great Sanhedrin who sat in the chamber of the Hewn Stone about whom Scripture said: "If a matter is too hard for you in judgment, etc." "And you shall rise and go up to the place that the LORD your God will choose, etc." "And you shall act according to the word that they tell you, etc." And yet, for all this, the Torah said that if the Great Court erred in a ruling and the people acted on their word, they bring an offering for their error. If they erred in idolatry and ruled, they bring a bull as a burnt offering and a goat as a sin offering, from each tribe. This is what is stated in the parashah of Shelach: "And it shall be, if from the eyes of the congregation it was done in error." And if in the other offences punishable by karet they erred and ruled, each tribe brings a bull for a sin offering. This is what is stated in the parashah of Vayikra. And if "all the congregation of Israel errs and the matter is hidden from the eyes of the assembly."

Literal: section 3 and indeed that you have begun words of your mouth folly and madness evil that you said how urged me my heart to inquire and to examine etc. and why not on the contrary the reverse is more reasonable since they say to worship Ze'ir Anpin who is created upright of heart folio 3 b and strength to God obligation complete upon us and upon every one and one in Israel to know (object marker) the LORD that he is Creator and not created and to worship him as it is written and you shall know this day and take to heart to your heart and you shall know that the LORD your God he God etc. and it says know (object marker) the God of your father and worship him and they permitted the Sages to keep away from the forbidden even in the privy and in the bathhouse for it is forbidden to meditate on them in the words of Torah but to keep away from the forbidden it is permitted as to your astonishment how could they err there is no this wonder nor are better sect the Kabbalists the later ones who went in their simplicity after the command of the philosopher the author of the Zohar who attributed to Rabbi Shimon bar Yochai and Rabbi Eleazar his son and their companions of the Sanhedrin the Great who sat in the chamber of the Hewn Stone about whom said Scripture if is too hard for you a matter for judgment etc. and you shall rise and go up to the place which will choose the LORD your God etc. and you shall do according to the mouth of the word which they tell you etc. and yet all this said the Torah that if they erred Beit Din (the court) the Great in a ruling and acted the people on their word they bring an offering for their error if in idolatry (avodah zarah) they erred and ruled they bring a bull for a burnt offering and a goat for a sin offering from each tribe (repeated) this is what is stated in the parashah of Shelach and it shall be if from the eyes of the congregation it was done in error and if in the other offences punishable by karet they erred and ruled he brings every tribe and each bull for a sin offering this is what is stated in the parashah of Vayikra and if all the congregation of Israel errs and is hidden matter from the eyes of the assembly

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:2.1 | section 3 |  |  |  |
| 3:2.2 | and indeed |  |  |  |
| 3:2.3 | that you have begun |  |  |  |
| 3:2.4 | words |  |  |  |
| 3:2.5 | of your mouth |  |  |  |
| 3:2.6 | folly |  |  |  |
| 3:2.7 | and madness |  |  |  |
| 3:2.8 | evil |  |  |  |
| 3:2.9 | that you said |  |  |  |
| 3:2.10 | how |  |  |  |
| 3:2.11 | urged me |  |  |  |
| 3:2.12 | my heart |  |  |  |
| 3:2.13 | to inquire |  |  |  |
| 3:2.14 | and to examine |  |  |  |
| 3:2.15 | etc. |  |  |  |
| 3:2.16 | and why |  |  |  |
| 3:2.17 | not |  |  |  |
| 3:2.18 | on the contrary |  |  |  |
| 3:2.19 | the reverse |  |  |  |
| 3:2.20 | is more reasonable |  |  |  |
| 3:2.21 | since |  |  |  |
| 3:2.22 | they say |  |  |  |
| 3:2.23 | to worship |  |  |  |
| 3:2.24 | Ze'ir Anpin |  |  | לז"א = לזעיר אנפין, 'to Ze'ir Anpin' (the Small Face, a Kabbalistic partzuf); first occurrence in this unit |
| 3:2.25 | who is |  |  |  |
| 3:2.26 | created |  |  |  |
| 3:2.27 | upright |  |  |  |
| 3:2.28 | of heart |  |  |  |
| 3:2.29 | folio |  |  |  |
| 3:2.30 | 3 |  |  |  |
| 3:2.31 | b |  |  | ע"ב = side b of the folio |
| 3:2.32 | and strength |  |  |  |
| 3:2.33 | to God |  |  | the author's spelling אלוהים for אלהים; 'to God' |
| 3:2.34 | obligation |  |  |  |
| 3:2.35 | complete |  |  |  |
| 3:2.36 | upon us |  |  |  |
| 3:2.37 | and upon |  |  |  |
| 3:2.38 | every |  |  |  |
| 3:2.39 | one |  |  |  |
| 3:2.40 | and one |  |  |  |
| 3:2.41 | in Israel |  |  |  |
| 3:2.42 | to know |  |  |  |
| 3:2.43 | (object marker) |  |  |  |
| 3:2.44 | the LORD |  |  |  |
| 3:2.45 | that he is |  |  |  |
| 3:2.46 | Creator |  |  |  |
| 3:2.47 | and not |  |  |  |
| 3:2.48 | created |  |  |  |
| 3:2.49 | and to worship him |  |  |  |
| 3:2.50 | as it is written |  |  |  |
| 3:2.51 | and you shall know |  |  |  |
| 3:2.52 | this day |  |  |  |
| 3:2.53 | and take to heart |  |  |  |
| 3:2.54 | to |  |  |  |
| 3:2.55 | your heart |  |  |  |
| 3:2.56 | and you shall know |  |  |  |
| 3:2.57 | that |  |  |  |
| 3:2.58 | the LORD |  |  |  |
| 3:2.59 | your God |  |  |  |
| 3:2.60 | he |  |  |  |
| 3:2.61 | God |  |  |  |
| 3:2.62 | etc. |  |  |  |
| 3:2.63 | and it says |  |  |  |
| 3:2.64 | know |  |  |  |
| 3:2.65 | (object marker) |  |  |  |
| 3:2.66 | the God of |  |  |  |
| 3:2.67 | your father |  |  |  |
| 3:2.68 | and worship him |  |  |  |
| 3:2.69 | and they permitted |  |  |  |
| 3:2.70 | the Sages |  |  |  |
| 3:2.71 | to keep away |  |  |  |
| 3:2.72 | from the forbidden |  |  |  |
| 3:2.73 | even |  |  |  |
| 3:2.74 | in the |  |  |  |
| 3:2.75 | privy |  |  |  |
| 3:2.76 | and in the |  |  |  |
| 3:2.77 | bathhouse |  |  |  |
| 3:2.78 | for it is forbidden |  |  |  |
| 3:2.79 | to meditate |  |  |  |
| 3:2.80 | on them |  |  |  |
| 3:2.81 | in the words of |  |  |  |
| 3:2.82 | Torah |  |  |  |
| 3:2.83 | but to keep away |  |  |  |
| 3:2.84 | from the forbidden |  |  |  |
| 3:2.85 | it is permitted |  |  |  |
| 3:2.86 | as to |  |  |  |
| 3:2.87 | your astonishment |  |  |  |
| 3:2.88 | how |  |  |  |
| 3:2.89 | could they err |  |  |  |
| 3:2.90 | there is no |  |  |  |
| 3:2.91 | this |  |  |  |
| 3:2.92 | wonder |  |  |  |
| 3:2.93 | nor |  |  |  |
| 3:2.94 | are better |  |  |  |
| 3:2.95 | sect |  |  | כת, 'sect', a group |
| 3:2.96 | the Kabbalists |  |  |  |
| 3:2.97 | the later ones |  |  | האחרונים, 'the later ones' |
| 3:2.98 | who |  |  |  |
| 3:2.99 | went |  |  |  |
| 3:2.100 | in their simplicity |  |  |  |
| 3:2.101 | after |  |  |  |
| 3:2.102 | the command |  |  |  |
| 3:2.103 | of |  |  |  |
| 3:2.104 | the philosopher |  |  | הפילוסוף, the author's epithet for the author of the Zohar (decisions) |
| 3:2.105 | the author of |  |  |  |
| 3:2.106 | the Zohar |  |  |  |
| 3:2.107 | who attributed |  |  |  |
| 3:2.108 | to Rabbi Shimon bar Yochai |  |  | שתלאו, literally 'hung on'; the verb means to attribute |
| 3:2.109 | and Rabbi |  |  |  |
| 3:2.110 | Eleazar |  |  |  |
| 3:2.111 | his son |  |  |  |
| 3:2.112 | and their companions |  |  |  |
| 3:2.113 | of the Sanhedrin |  |  |  |
| 3:2.114 | the Great |  |  |  |
| 3:2.115 | who sat |  |  |  |
| 3:2.116 | in the chamber of |  |  |  |
| 3:2.117 | the Hewn Stone |  |  |  |
| 3:2.118 | about whom |  |  |  |
| 3:2.119 | said |  |  |  |
| 3:2.120 | Scripture |  |  |  |
| 3:2.121 | if |  |  |  |
| 3:2.122 | is too hard |  |  |  |
| 3:2.123 | for you |  |  |  |
| 3:2.124 | a matter |  |  |  |
| 3:2.125 | for judgment |  |  |  |
| 3:2.126 | etc. |  |  |  |
| 3:2.127 | and you shall rise |  |  |  |
| 3:2.128 | and go up |  |  |  |
| 3:2.129 | to |  |  |  |
| 3:2.130 | the place |  |  |  |
| 3:2.131 | which |  |  |  |
| 3:2.132 | will choose |  |  |  |
| 3:2.133 | the LORD |  |  |  |
| 3:2.134 | your God |  |  |  |
| 3:2.135 | etc. |  |  |  |
| 3:2.136 | and you shall do |  |  |  |
| 3:2.137 | according to |  |  |  |
| 3:2.138 | the mouth of |  |  |  |
| 3:2.139 | the word |  |  |  |
| 3:2.140 | which |  |  |  |
| 3:2.141 | they tell |  |  |  |
| 3:2.142 | you |  |  |  |
| 3:2.143 | etc. |  |  |  |
| 3:2.144 | and yet |  |  |  |
| 3:2.145 | all |  |  |  |
| 3:2.146 | this |  |  |  |
| 3:2.147 | said |  |  |  |
| 3:2.148 | the Torah |  |  |  |
| 3:2.149 | that if |  |  |  |
| 3:2.150 | they erred |  |  |  |
| 3:2.151 | Beit Din (the court) |  |  | ב"ד הגדול = בית דין הגדול, 'the Great Court' |
| 3:2.152 | the Great |  |  |  |
| 3:2.153 | in a ruling |  |  |  |
| 3:2.154 | and acted |  |  |  |
| 3:2.155 | the people |  |  |  |
| 3:2.156 | on |  |  |  |
| 3:2.157 | their word |  |  |  |
| 3:2.158 | they bring |  |  |  |
| 3:2.159 | an offering |  |  |  |
| 3:2.160 | for |  |  |  |
| 3:2.161 | their error |  |  |  |
| 3:2.162 | if |  |  |  |
| 3:2.163 | in idolatry (avodah zarah) |  |  | ע"ז = עבודה זרה, 'idolatry' (literally 'strange worship') |
| 3:2.164 | they erred |  |  |  |
| 3:2.165 | and ruled |  |  |  |
| 3:2.166 | they bring |  |  |  |
| 3:2.167 | a bull |  |  |  |
| 3:2.168 | for a burnt offering |  |  |  |
| 3:2.169 | and a goat |  |  |  |
| 3:2.170 | for a sin offering |  |  |  |
| 3:2.171 | from |  |  |  |
| 3:2.172 | each |  |  |  |
| 3:2.173 | tribe (repeated) |  |  |  |
| 3:2.174 | this is |  |  |  |
| 3:2.175 | what is stated |  |  |  |
| 3:2.176 | in the parashah of |  |  |  |
| 3:2.177 | Shelach |  |  | Shelach: the parashah of Numbers 13-15 |
| 3:2.178 | and it shall be |  |  |  |
| 3:2.179 | if |  |  |  |
| 3:2.180 | from the eyes of |  |  |  |
| 3:2.181 | the congregation |  |  |  |
| 3:2.182 | it was done |  |  |  |
| 3:2.183 | in error |  |  |  |
| 3:2.184 | and if |  |  |  |
| 3:2.185 | in the other |  |  |  |
| 3:2.186 | offences punishable by karet |  |  | כריתות: offences punished by karet |
| 3:2.187 | they erred |  |  |  |
| 3:2.188 | and ruled |  |  |  |
| 3:2.189 | he brings |  |  |  |
| 3:2.190 | every |  |  |  |
| 3:2.191 | tribe |  |  |  |
| 3:2.192 | and each |  |  |  |
| 3:2.193 | bull |  |  |  |
| 3:2.194 | for a sin offering |  |  |  |
| 3:2.195 | this is |  |  |  |
| 3:2.196 | what is stated |  |  |  |
| 3:2.197 | in the parashah of |  |  |  |
| 3:2.198 | Vayikra |  |  | ויקרא: the parashah of Vayikra (Leviticus) |
| 3:2.199 | and if |  |  |  |
| 3:2.200 | all |  |  |  |
| 3:2.201 | the congregation of |  |  |  |
| 3:2.202 | Israel |  |  |  |
| 3:2.203 | errs |  |  |  |
| 3:2.204 | and is hidden |  |  |  |
| 3:2.205 | matter |  |  |  |
| 3:2.206 | from the eyes of |  |  |  |
| 3:2.207 | the assembly |  |  |  |

Phrases:
- 3:2.1 3:2.2 | §3. And indeed, | section 3 and indeed
- 3:2.3 3:2.4 3:2.5 3:2.6 3:2.7 3:2.8 3:2.9 | that you have begun the words of your mouth, folly and madness, evil that you said, | that you have begun words of your mouth folly and madness evil that you said
- 3:2.10 3:2.11 3:2.12 3:2.13 3:2.14 | "How my heart urged me to inquire and to examine, | how urged me my heart to inquire and to examine
- 3:2.15 | etc." | etc.
- 3:2.16 3:2.17 | And why not? | and why not
- 3:2.18 3:2.19 3:2.20 | On the contrary, the reverse is more reasonable | on the contrary the reverse is more reasonable
- 3:2.21 3:2.22 3:2.23 3:2.24 3:2.25 3:2.26 3:2.27 3:2.28 3:2.29 3:2.30 3:2.31 3:2.32 3:2.33 | (since they say to worship Ze'ir Anpin, who is created, upright of heart, folio 3b, and strength to God). | since they say to worship Ze'ir Anpin who is created upright of heart folio 3 b and strength to God
- 3:2.34 3:2.35 3:2.36 | It is a complete obligation upon us | obligation complete upon us
- 3:2.37 3:2.38 3:2.39 3:2.40 3:2.41 | and upon each and every one of Israel | and upon every one and one in Israel
- 3:2.42 3:2.43 3:2.44 3:2.45 3:2.46 3:2.47 3:2.48 3:2.49 3:2.50 | to know the LORD, that He is Creator and not created, and to worship Him, as it is written: | to know (object marker) the LORD that he is Creator and not created and to worship him as it is written
- 3:2.51 3:2.52 3:2.53 3:2.54 3:2.55 | "And you shall know this day and take it to your heart." | and you shall know this day and take to heart to your heart
- 3:2.56 3:2.57 3:2.58 3:2.59 3:2.60 3:2.61 3:2.62 | "And you shall know that the LORD your God is God, etc." | and you shall know that the LORD your God he God etc.
- 3:2.63 3:2.64 3:2.65 3:2.66 3:2.67 3:2.68 | And it says: "Know the God of your father and worship Him." | and it says know (object marker) the God of your father and worship him
- 3:2.69 3:2.70 3:2.71 3:2.72 | The Sages permitted abstaining from the forbidden, | and they permitted the Sages to keep away from the forbidden
- 3:2.73 3:2.74 3:2.75 3:2.76 3:2.77 | even in the privy and in the bathhouse, | even in the privy and in the bathhouse
- 3:2.78 3:2.79 3:2.80 3:2.81 3:2.82 | where it is forbidden to meditate on them in words of Torah. | for it is forbidden to meditate on them in the words of Torah
- 3:2.83 3:2.84 3:2.85 | but to keep away from the forbidden is permitted. | but to keep away from the forbidden it is permitted
- 3:2.86 3:2.87 3:2.88 3:2.89 | As to your astonishment, how could they err? | as to your astonishment how could they err
- 3:2.90 3:2.91 3:2.92 | There is no wonder in this. | there is no this wonder
- 3:2.93 3:2.94 3:2.95 3:2.96 3:2.97 3:2.98 3:2.99 3:2.100 | Nor is the sect of the later Kabbalists any better, those who went in their simplicity | nor are better sect the Kabbalists the later ones who went in their simplicity
- 3:2.101 3:2.102 3:2.103 3:2.104 3:2.105 3:2.106 | after the command of the philosopher, the author of the Zohar | after the command of the philosopher the author of the Zohar
- 3:2.107 3:2.108 3:2.109 3:2.110 3:2.111 3:2.112 3:2.113 3:2.114 | who attributed it to Rabbi Shimon bar Yochai, and Rabbi Eleazar his son, and their companions of the Great Sanhedrin | who attributed to Rabbi Shimon bar Yochai and Rabbi Eleazar his son and their companions of the Sanhedrin the Great
- 3:2.115 3:2.116 3:2.117 | who sat in the chamber of the Hewn Stone | who sat in the chamber of the Hewn Stone
- 3:2.118 3:2.119 3:2.120 3:2.121 3:2.122 3:2.123 3:2.124 3:2.125 3:2.126 | about whom Scripture said: "If a matter is too hard for you in judgment, etc." | about whom said Scripture if is too hard for you a matter for judgment etc.
- 3:2.127 3:2.128 3:2.129 3:2.130 3:2.131 3:2.132 3:2.133 3:2.134 3:2.135 | "And you shall rise and go up to the place that the LORD your God will choose, etc." | and you shall rise and go up to the place which will choose the LORD your God etc.
- 3:2.136 3:2.137 3:2.138 3:2.139 3:2.140 3:2.141 3:2.142 3:2.143 | "And you shall act according to the word that they tell you, etc." | and you shall do according to the mouth of the word which they tell you etc.
- 3:2.144 3:2.145 3:2.146 3:2.147 3:2.148 | And yet, for all this, the Torah said | and yet all this said the Torah
- 3:2.149 3:2.150 3:2.151 3:2.152 3:2.153 3:2.154 3:2.155 3:2.156 3:2.157 3:2.158 3:2.159 3:2.160 3:2.161 | that if the Great Court erred in a ruling and the people acted on their word, they bring an offering for their error. | that if they erred Beit Din (the court) the Great in a ruling and acted the people on their word they bring an offering for their error
- 3:2.162 3:2.163 3:2.164 3:2.165 3:2.166 3:2.167 3:2.168 3:2.169 3:2.170 3:2.171 3:2.172 3:2.173 | If they erred in idolatry and ruled, they bring a bull as a burnt offering and a goat as a sin offering, from each tribe. | if in idolatry (avodah zarah) they erred and ruled they bring a bull for a burnt offering and a goat for a sin offering from each tribe (repeated)
- 3:2.174 3:2.175 3:2.176 3:2.177 3:2.178 3:2.179 3:2.180 3:2.181 3:2.182 3:2.183 | This is what is stated in the parashah of Shelach: "And it shall be, if from the eyes of the congregation it was done in error." | this is what is stated in the parashah of Shelach and it shall be if from the eyes of the congregation it was done in error
- 3:2.184 3:2.185 3:2.186 3:2.187 3:2.188 3:2.189 3:2.190 3:2.191 3:2.192 3:2.193 3:2.194 3:2.195 3:2.196 3:2.197 3:2.198 | And if in the other offences punishable by karet they erred and ruled, each tribe brings a bull for a sin offering. This is what is stated in the parashah of Vayikra. | and if in the other offences punishable by karet they erred and ruled he brings every tribe and each bull for a sin offering this is what is stated in the parashah of Vayikra
- 3:2.199 3:2.200 3:2.201 3:2.202 3:2.203 3:2.204 3:2.205 3:2.206 3:2.207 | And if "all the congregation of Israel errs and the matter is hidden from the eyes of the assembly." | and if all the congregation of Israel errs and is hidden matter from the eyes of the assembly

Grammar notes:
- שהחלת דברי פיך: a perfect with the second-person ending, 'that you have begun the words of your mouth'; the reader's words are cited as 'שאמרת'.
- איך מלאני לבי לדרוש ולחקור: the infinitives לדרוש and לחקור follow 'urged me'; the reader's question is closed by 'וכו'' and the author's reply begins with 'ולמה לא'.
- ועוז לאלוהים sits in the author's parenthesis and gives no formula, so it is not marked as quotation; the parenthesis is kept as written.
- ואומר: 'and it says' introduces a verse; 'ועבדהו' is the verb 'worship' with the object suffix -הו, 'worship him'.
- The conditional clauses (אם שגגו ... מביאים; ואם שגגו ... מביא) are built on perfects of 'to err' with present-tense apodosis, as the verses have them.

Commentary rulings:
- apparatus: Deut 4:39: 'וידעת היום והשבות אל לבבך', quoted with 'ככתוב'.
- apparatus: Deut 7:9: 'וידעת כי ה' אלהיך הוא האלהים', quoted with 'וכו''.
- apparatus: 1 Chr 28:9: 'דע את אלהי אביך ועבדהו', quoted with 'ואומר'.
- apparatus: ועוז לאלוהים in the parenthesis: possibly Ps 68:35; not verified, and no formula marks it.
- apparatus: Deut 17:8: 'כי יפלא ממך דבר למשפט' and 'וקמת ועלית אל המקום אשר יבחר ה' אלהיך'. Deut 17:10: 'ועשית על פי הדבר אשר יגידו לך'. The author writes 'אלוהיך' for 'אלהיך'; the English follows the verse.
- apparatus: Num 15:24: 'והיה אם מעיני העדה נעשתה לשגגה'.
- apparatus: Lev 4:13: 'ואם כל עדת ישראל ישגו ונעלם דבר מעיני הקהל'.
- apparatus: The author's 'folio 3b' is given as he writes it; no Zohar text is quoted here.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

