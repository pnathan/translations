# Translation pack: milhamot-hashem perek-4-07 (Milchamot ha-Shem 4:11–4:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-07`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
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
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- עבודה **אחרת** → "other worship" (note first in a unit).
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **כוונה** (in prayer) → "intention".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- הגשת **מנחה** (argued from Gen 43 / 1 Sam 10) → "gift"; "meal offering" only
  in a sacrificial context.
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ח"א / ח"ב** → "vol. 1" / "vol. 2".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 4:11

Hebrew (Wikisource): כו) עוד בזהר בראשית דף ל"ד (ח"א לד, א) כתיב ה' מלך ה' מלך ה' ימלוך לעולם ועד ה' מלך לעילא. ה' מלך: באמצעיתא. ה' ימלוך לתתא ע"כ. ופי' המקדש מלך בשם הכוונות. ה' מלך הוא אריך אנפין. ה' מלך אבא ואימא. ה' ימלוך זעיר ונוקביה ע"כ. ובפרשת בשלח דף ס"ד (ח"ב סז, א) על פסוק היש ה' בקרבנו אם אין מקשה בזהר וכי טפשין הוו ישראל וכו' אלא בעו למנדע בין עתיקא סתימא דאיקרי אין ובין זעיר אנפין דאיקרי ה', ועל דא לא כתיב היש ה' בקרבנו אם לא כמה דכתיב הילך בתורתי אם לא. אמרו ישראל אי האיי נשאל בגוונא חד. ואי האיי נשאל בגוונא אחרינא. ועל דא מיד ויבא עמלק ופי' מהר"י לופיס שכוונת הזהר היא שאבותינו בקשו לדעת מי הוא המנהיגם ועושה להם כל הנסים הללו אם הקב"ה הנקרא ז"א והויה או אין הנקרא עתיק, ועשו החקירה הזאת כדי לעבדו באופן הראוי לו, אם זעיר אנפין באופן אחר, ואם עתיק באופן אחר לפי שיש הפרש בין עבודה לעבודה, ובין כוונה לכוונה, והיו בספק זה עד ששמעו אנכי ה' אלהיך, ואז ידעו שהוא ז"א עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:11.1 | כו | ) |
| 4:11.2 | עוד |  |
| 4:11.3 | בזהר |  |
| 4:11.4 | בראשית |  |
| 4:11.5 | דף |  |
| 4:11.6 | ל"ד |  |
| 4:11.7 | ח"א |  |
| 4:11.8 | לד | , |
| 4:11.9 | א | ) |
| 4:11.10 | כתיב |  |
| 4:11.11 | ה | ' |
| 4:11.12 | מלך |  |
| 4:11.13 | ה | ' |
| 4:11.14 | מלך |  |
| 4:11.15 | ה | ' |
| 4:11.16 | ימלוך |  |
| 4:11.17 | לעולם |  |
| 4:11.18 | ועד |  |
| 4:11.19 | ה | ' |
| 4:11.20 | מלך |  |
| 4:11.21 | לעילא | . |
| 4:11.22 | ה | ' |
| 4:11.23 | מלך | : |
| 4:11.24 | באמצעיתא | . |
| 4:11.25 | ה | ' |
| 4:11.26 | ימלוך |  |
| 4:11.27 | לתתא |  |
| 4:11.28 | ע"כ | . |
| 4:11.29 | ופי | ' |
| 4:11.30 | המקדש |  |
| 4:11.31 | מלך |  |
| 4:11.32 | בשם |  |
| 4:11.33 | הכוונות | . |
| 4:11.34 | ה | ' |
| 4:11.35 | מלך |  |
| 4:11.36 | הוא |  |
| 4:11.37 | אריך |  |
| 4:11.38 | אנפין | . |
| 4:11.39 | ה | ' |
| 4:11.40 | מלך |  |
| 4:11.41 | אבא |  |
| 4:11.42 | ואימא | . |
| 4:11.43 | ה | ' |
| 4:11.44 | ימלוך |  |
| 4:11.45 | זעיר |  |
| 4:11.46 | ונוקביה |  |
| 4:11.47 | ע"כ | . |
| 4:11.48 | ובפרשת |  |
| 4:11.49 | בשלח |  |
| 4:11.50 | דף |  |
| 4:11.51 | ס"ד |  |
| 4:11.52 | ח"ב |  |
| 4:11.53 | סז | , |
| 4:11.54 | א | ) |
| 4:11.55 | על |  |
| 4:11.56 | פסוק |  |
| 4:11.57 | היש |  |
| 4:11.58 | ה | ' |
| 4:11.59 | בקרבנו |  |
| 4:11.60 | אם |  |
| 4:11.61 | אין |  |
| 4:11.62 | מקשה |  |
| 4:11.63 | בזהר |  |
| 4:11.64 | וכי |  |
| 4:11.65 | טפשין |  |
| 4:11.66 | הוו |  |
| 4:11.67 | ישראל |  |
| 4:11.68 | וכו | ' |
| 4:11.69 | אלא |  |
| 4:11.70 | בעו |  |
| 4:11.71 | למנדע |  |
| 4:11.72 | בין |  |
| 4:11.73 | עתיקא |  |
| 4:11.74 | סתימא |  |
| 4:11.75 | דאיקרי |  |
| 4:11.76 | אין |  |
| 4:11.77 | ובין |  |
| 4:11.78 | זעיר |  |
| 4:11.79 | אנפין |  |
| 4:11.80 | דאיקרי |  |
| 4:11.81 | ה | ', |
| 4:11.82 | ועל |  |
| 4:11.83 | דא |  |
| 4:11.84 | לא |  |
| 4:11.85 | כתיב |  |
| 4:11.86 | היש |  |
| 4:11.87 | ה | ' |
| 4:11.88 | בקרבנו |  |
| 4:11.89 | אם |  |
| 4:11.90 | לא |  |
| 4:11.91 | כמה |  |
| 4:11.92 | דכתיב |  |
| 4:11.93 | הילך |  |
| 4:11.94 | בתורתי |  |
| 4:11.95 | אם |  |
| 4:11.96 | לא | . |
| 4:11.97 | אמרו |  |
| 4:11.98 | ישראל |  |
| 4:11.99 | אי |  |
| 4:11.100 | האיי |  |
| 4:11.101 | נשאל |  |
| 4:11.102 | בגוונא |  |
| 4:11.103 | חד | . |
| 4:11.104 | ואי |  |
| 4:11.105 | האיי |  |
| 4:11.106 | נשאל |  |
| 4:11.107 | בגוונא |  |
| 4:11.108 | אחרינא | . |
| 4:11.109 | ועל |  |
| 4:11.110 | דא |  |
| 4:11.111 | מיד |  |
| 4:11.112 | ויבא |  |
| 4:11.113 | עמלק |  |
| 4:11.114 | ופי | ' |
| 4:11.115 | מהר"י |  |
| 4:11.116 | לופיס |  |
| 4:11.117 | שכוונת |  |
| 4:11.118 | הזהר |  |
| 4:11.119 | היא |  |
| 4:11.120 | שאבותינו |  |
| 4:11.121 | בקשו |  |
| 4:11.122 | לדעת |  |
| 4:11.123 | מי |  |
| 4:11.124 | הוא |  |
| 4:11.125 | המנהיגם |  |
| 4:11.126 | ועושה |  |
| 4:11.127 | להם |  |
| 4:11.128 | כל |  |
| 4:11.129 | הנסים |  |
| 4:11.130 | הללו |  |
| 4:11.131 | אם |  |
| 4:11.132 | הקב"ה |  |
| 4:11.133 | הנקרא |  |
| 4:11.134 | ז"א |  |
| 4:11.135 | והויה |  |
| 4:11.136 | או |  |
| 4:11.137 | אין |  |
| 4:11.138 | הנקרא |  |
| 4:11.139 | עתיק | , |
| 4:11.140 | ועשו |  |
| 4:11.141 | החקירה |  |
| 4:11.142 | הזאת |  |
| 4:11.143 | כדי |  |
| 4:11.144 | לעבדו |  |
| 4:11.145 | באופן |  |
| 4:11.146 | הראוי |  |
| 4:11.147 | לו | , |
| 4:11.148 | אם |  |
| 4:11.149 | זעיר |  |
| 4:11.150 | אנפין |  |
| 4:11.151 | באופן |  |
| 4:11.152 | אחר | , |
| 4:11.153 | ואם |  |
| 4:11.154 | עתיק |  |
| 4:11.155 | באופן |  |
| 4:11.156 | אחר |  |
| 4:11.157 | לפי |  |
| 4:11.158 | שיש |  |
| 4:11.159 | הפרש |  |
| 4:11.160 | בין |  |
| 4:11.161 | עבודה |  |
| 4:11.162 | לעבודה | , |
| 4:11.163 | ובין |  |
| 4:11.164 | כוונה |  |
| 4:11.165 | לכוונה | , |
| 4:11.166 | והיו |  |
| 4:11.167 | בספק |  |
| 4:11.168 | זה |  |
| 4:11.169 | עד |  |
| 4:11.170 | ששמעו |  |
| 4:11.171 | אנכי |  |
| 4:11.172 | ה | ' |
| 4:11.173 | אלהיך | , |
| 4:11.174 | ואז |  |
| 4:11.175 | ידעו |  |
| 4:11.176 | שהוא |  |
| 4:11.177 | ז"א |  |
| 4:11.178 | עכ"ל | . |


## Milchamot ha-Shem 4:12

Hebrew (Wikisource): ובמנחות פרק הרי עלי עשרון גרסינן תנא אמר ר' שמעון בן עזאי בא וראה מה כתיב בפרשת קרבנות שלא נאמר בהן לא אל ולא אלהים אלא לה' שלא ליתן פתחון פה לבעל הדין לחלוק, ופי' רש"י לבעל הדין לומר רשויות הן, והדבר מוכיח שזה ששמו כך צוה להקריב לו מנחה, וזה ששמו כך צוה להקריב לו פרים ואלים ע"כ. הא למדת, שזוהי דעת מחבר הזהר שאלוהות מחולקים הם, ובקשו אבותינו (לפי דבריו) לידע בחקירתם איזה הוא האלוה המנהיגם, אם עתיק הנקרא אין עבודתו באופן אחר, ואם זעיר אנפין עבודתו אחרת, ולא נתברר להם ספק זה עד מתן תורה שאז ידעו (לפי שיטתו) שהוא זעיר אנפין. ויתבאר זה כמו שכ' הרמב"ם פרק ג' מה"ל ע"ז וז"ל עבודות הרבה קבעו עובדו ע"ז לכל צלם וצלם עבודה מיוחדת לו. אבל אנחנו מאמיני האחדות שאין כמוה בכל שאר האחדים אין מקום לדברים אלו כדברי רז"ל במנחות הנז"ל.

| id | word | punctuation after |
|---|---|---|
| 4:12.1 | ובמנחות |  |
| 4:12.2 | פרק |  |
| 4:12.3 | הרי |  |
| 4:12.4 | עלי |  |
| 4:12.5 | עשרון |  |
| 4:12.6 | גרסינן |  |
| 4:12.7 | תנא |  |
| 4:12.8 | אמר |  |
| 4:12.9 | ר | ' |
| 4:12.10 | שמעון |  |
| 4:12.11 | בן |  |
| 4:12.12 | עזאי |  |
| 4:12.13 | בא |  |
| 4:12.14 | וראה |  |
| 4:12.15 | מה |  |
| 4:12.16 | כתיב |  |
| 4:12.17 | בפרשת |  |
| 4:12.18 | קרבנות |  |
| 4:12.19 | שלא |  |
| 4:12.20 | נאמר |  |
| 4:12.21 | בהן |  |
| 4:12.22 | לא |  |
| 4:12.23 | אל |  |
| 4:12.24 | ולא |  |
| 4:12.25 | אלהים |  |
| 4:12.26 | אלא |  |
| 4:12.27 | לה | ' |
| 4:12.28 | שלא |  |
| 4:12.29 | ליתן |  |
| 4:12.30 | פתחון |  |
| 4:12.31 | פה |  |
| 4:12.32 | לבעל |  |
| 4:12.33 | הדין |  |
| 4:12.34 | לחלוק | , |
| 4:12.35 | ופי | ' |
| 4:12.36 | רש"י |  |
| 4:12.37 | לבעל |  |
| 4:12.38 | הדין |  |
| 4:12.39 | לומר |  |
| 4:12.40 | רשויות |  |
| 4:12.41 | הן | , |
| 4:12.42 | והדבר |  |
| 4:12.43 | מוכיח |  |
| 4:12.44 | שזה |  |
| 4:12.45 | ששמו |  |
| 4:12.46 | כך |  |
| 4:12.47 | צוה |  |
| 4:12.48 | להקריב |  |
| 4:12.49 | לו |  |
| 4:12.50 | מנחה | , |
| 4:12.51 | וזה |  |
| 4:12.52 | ששמו |  |
| 4:12.53 | כך |  |
| 4:12.54 | צוה |  |
| 4:12.55 | להקריב |  |
| 4:12.56 | לו |  |
| 4:12.57 | פרים |  |
| 4:12.58 | ואלים |  |
| 4:12.59 | ע"כ | . |
| 4:12.60 | הא |  |
| 4:12.61 | למדת | , |
| 4:12.62 | שזוהי |  |
| 4:12.63 | דעת |  |
| 4:12.64 | מחבר |  |
| 4:12.65 | הזהר |  |
| 4:12.66 | שאלוהות |  |
| 4:12.67 | מחולקים |  |
| 4:12.68 | הם | , |
| 4:12.69 | ובקשו |  |
| 4:12.70 | אבותינו |  |
| 4:12.71 | לפי |  |
| 4:12.72 | דבריו | ) |
| 4:12.73 | לידע |  |
| 4:12.74 | בחקירתם |  |
| 4:12.75 | איזה |  |
| 4:12.76 | הוא |  |
| 4:12.77 | האלוה |  |
| 4:12.78 | המנהיגם | , |
| 4:12.79 | אם |  |
| 4:12.80 | עתיק |  |
| 4:12.81 | הנקרא |  |
| 4:12.82 | אין |  |
| 4:12.83 | עבודתו |  |
| 4:12.84 | באופן |  |
| 4:12.85 | אחר | , |
| 4:12.86 | ואם |  |
| 4:12.87 | זעיר |  |
| 4:12.88 | אנפין |  |
| 4:12.89 | עבודתו |  |
| 4:12.90 | אחרת | , |
| 4:12.91 | ולא |  |
| 4:12.92 | נתברר |  |
| 4:12.93 | להם |  |
| 4:12.94 | ספק |  |
| 4:12.95 | זה |  |
| 4:12.96 | עד |  |
| 4:12.97 | מתן |  |
| 4:12.98 | תורה |  |
| 4:12.99 | שאז |  |
| 4:12.100 | ידעו |  |
| 4:12.101 | לפי |  |
| 4:12.102 | שיטתו | ) |
| 4:12.103 | שהוא |  |
| 4:12.104 | זעיר |  |
| 4:12.105 | אנפין | . |
| 4:12.106 | ויתבאר |  |
| 4:12.107 | זה |  |
| 4:12.108 | כמו |  |
| 4:12.109 | שכ | ' |
| 4:12.110 | הרמב"ם |  |
| 4:12.111 | פרק |  |
| 4:12.112 | ג | ' |
| 4:12.113 | מה"ל |  |
| 4:12.114 | ע"ז |  |
| 4:12.115 | וז"ל |  |
| 4:12.116 | עבודות |  |
| 4:12.117 | הרבה |  |
| 4:12.118 | קבעו |  |
| 4:12.119 | עובדו |  |
| 4:12.120 | ע"ז |  |
| 4:12.121 | לכל |  |
| 4:12.122 | צלם |  |
| 4:12.123 | וצלם |  |
| 4:12.124 | עבודה |  |
| 4:12.125 | מיוחדת |  |
| 4:12.126 | לו | . |
| 4:12.127 | אבל |  |
| 4:12.128 | אנחנו |  |
| 4:12.129 | מאמיני |  |
| 4:12.130 | האחדות |  |
| 4:12.131 | שאין |  |
| 4:12.132 | כמוה |  |
| 4:12.133 | בכל |  |
| 4:12.134 | שאר |  |
| 4:12.135 | האחדים |  |
| 4:12.136 | אין |  |
| 4:12.137 | מקום |  |
| 4:12.138 | לדברים |  |
| 4:12.139 | אלו |  |
| 4:12.140 | כדברי |  |
| 4:12.141 | רז"ל |  |
| 4:12.142 | במנחות |  |
| 4:12.143 | הנז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

