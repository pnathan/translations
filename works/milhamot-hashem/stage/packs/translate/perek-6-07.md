# Translation pack: milhamot-hashem perek-6-07 (Milchamot ha-Shem 6:12–6:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-07`.

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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
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

- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- הסבה **הראשונה** → "the First Cause".
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
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **תא חזי** → "come and see".
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ח"א / ח"ב** → "vol. 1" / "vol. 2".
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


## Milchamot ha-Shem 6:12

Hebrew (Wikisource): ועל ידי הרבנים הפתאים שקבלוהו לתומם באמונת אומן שהוא חבור רשב"י והחזיקו בו על שמו ולא שמו אל לבם לחקור ולדרוש מבטן מי יצא, ומי התיר לדרוש ולחקור במופלא ומכוסא, ושהוא סותר אמונתינו הטהורה שקבלנו מאבותינו ומרז"ל באל אחד מיוחד מכל שאר האחדים, וזה לסבת שעירב בו הרבה דברים מדברי רז"ל, ומזהיר מאד בקיום המצות של תורה והעסק בתורה שבכתב ותורתו החדשה שהמציאה חלף המשנה והתלמוד שהם מעורבים מטוב ורע, וקורא לתורתו תורה שבע"פ שכולה טוב ורק שיהיו לשם אלהים אחרים, קב"ה ושכינתיה עלאין (אבא ואימא) וקב"ה ושכינתיה התאין (זעיר ונוקביה) ולא שנה במצות התורה רק בקצת פרטים שכיוון לרמוז אמונתו באלהים רבים, ולשום עיקר העבודה לבן שהוא זעיר אנפין, וברוב המצות מכוינים בהם לקשט הנשים שכינתא עלאה ושכינתא תתאה, כדי למצוא חן בעיני בעליהן לעוררם לתשמיש (כמו שכ' הר' משנת חסידים בתיקון חצות וז"ל וכיצד יעשה ישב על הארץ אצל המזוזה שבפתח שרומזת להח' ורגליו יחפים ממנעל, ויתן אפר מקלה וכו' ועל ידי זה מתחמם הזעיר אנפין ומתאוה לה וכו') זה להוליד נשמות למלאכים. וזה להוליד נשמות חדשות לישראל (זהר יתרו דף פ"ט ע"כ) שלא כרצון חכמים בבראשית רבה (הבאתיו לעיל סי' ל"א כל שיש לו תולדות וכו') ויחס עיקר העבודה וקיום המצות והברכות לבן כנז"ל. ועוד יתבאר לקמן. ועל ידי רעיונותיהם העולים על רוחם חשבו לפרש הכתובים שנאמרו על דרך המליצה והרחבת הלשון או בדרך חיבה כמו שאמרו רז"ל לא זז מחבבה וכו' כמשמען והנחתן הראשונה יד ה', עיני ה', יעשן אף ה' עלה עשן באפו, בצלמנו כדמותנו, בנים אתם לה' אלהיכם, אחי ורעי, ויראו את אלוהי ישראל, עורה למה תישן ה', הנה לא ינום ולא יישן שומר ישראל, וכיוצא בהן, יפרשו אותן על ידי רעיונות אלו כהנחתן הראשונה, היפך מאמר רז"ל המתרגם פסוק כצורתו הרי זה בדאי. וכמו שכ' הרס"ג בהאמונות והדעות מאמר האחדות. והרמב"ם במורה ח"א וזולתם.

| id | word | punctuation after |
|---|---|---|
| 6:12.1 | ועל |  |
| 6:12.2 | ידי |  |
| 6:12.3 | הרבנים |  |
| 6:12.4 | הפתאים |  |
| 6:12.5 | שקבלוהו |  |
| 6:12.6 | לתומם |  |
| 6:12.7 | באמונת |  |
| 6:12.8 | אומן |  |
| 6:12.9 | שהוא |  |
| 6:12.10 | חבור |  |
| 6:12.11 | רשב"י |  |
| 6:12.12 | והחזיקו |  |
| 6:12.13 | בו |  |
| 6:12.14 | על |  |
| 6:12.15 | שמו |  |
| 6:12.16 | ולא |  |
| 6:12.17 | שמו |  |
| 6:12.18 | אל |  |
| 6:12.19 | לבם |  |
| 6:12.20 | לחקור |  |
| 6:12.21 | ולדרוש |  |
| 6:12.22 | מבטן |  |
| 6:12.23 | מי |  |
| 6:12.24 | יצא | , |
| 6:12.25 | ומי |  |
| 6:12.26 | התיר |  |
| 6:12.27 | לדרוש |  |
| 6:12.28 | ולחקור |  |
| 6:12.29 | במופלא |  |
| 6:12.30 | ומכוסא | , |
| 6:12.31 | ושהוא |  |
| 6:12.32 | סותר |  |
| 6:12.33 | אמונתינו |  |
| 6:12.34 | הטהורה |  |
| 6:12.35 | שקבלנו |  |
| 6:12.36 | מאבותינו |  |
| 6:12.37 | ומרז"ל |  |
| 6:12.38 | באל |  |
| 6:12.39 | אחד |  |
| 6:12.40 | מיוחד |  |
| 6:12.41 | מכל |  |
| 6:12.42 | שאר |  |
| 6:12.43 | האחדים | , |
| 6:12.44 | וזה |  |
| 6:12.45 | לסבת |  |
| 6:12.46 | שעירב |  |
| 6:12.47 | בו |  |
| 6:12.48 | הרבה |  |
| 6:12.49 | דברים |  |
| 6:12.50 | מדברי |  |
| 6:12.51 | רז"ל | , |
| 6:12.52 | ומזהיר |  |
| 6:12.53 | מאד |  |
| 6:12.54 | בקיום |  |
| 6:12.55 | המצות |  |
| 6:12.56 | של |  |
| 6:12.57 | תורה |  |
| 6:12.58 | והעסק |  |
| 6:12.59 | בתורה |  |
| 6:12.60 | שבכתב |  |
| 6:12.61 | ותורתו |  |
| 6:12.62 | החדשה |  |
| 6:12.63 | שהמציאה |  |
| 6:12.64 | חלף |  |
| 6:12.65 | המשנה |  |
| 6:12.66 | והתלמוד |  |
| 6:12.67 | שהם |  |
| 6:12.68 | מעורבים |  |
| 6:12.69 | מטוב |  |
| 6:12.70 | ורע | , |
| 6:12.71 | וקורא |  |
| 6:12.72 | לתורתו |  |
| 6:12.73 | תורה |  |
| 6:12.74 | שבע"פ |  |
| 6:12.75 | שכולה |  |
| 6:12.76 | טוב |  |
| 6:12.77 | ורק |  |
| 6:12.78 | שיהיו |  |
| 6:12.79 | לשם |  |
| 6:12.80 | אלהים |  |
| 6:12.81 | אחרים | , |
| 6:12.82 | קב"ה |  |
| 6:12.83 | ושכינתיה |  |
| 6:12.84 | עלאין |  |
| 6:12.85 | אבא |  |
| 6:12.86 | ואימא | ) |
| 6:12.87 | וקב"ה |  |
| 6:12.88 | ושכינתיה |  |
| 6:12.89 | התאין |  |
| 6:12.90 | זעיר |  |
| 6:12.91 | ונוקביה | ) |
| 6:12.92 | ולא |  |
| 6:12.93 | שנה |  |
| 6:12.94 | במצות |  |
| 6:12.95 | התורה |  |
| 6:12.96 | רק |  |
| 6:12.97 | בקצת |  |
| 6:12.98 | פרטים |  |
| 6:12.99 | שכיוון |  |
| 6:12.100 | לרמוז |  |
| 6:12.101 | אמונתו |  |
| 6:12.102 | באלהים |  |
| 6:12.103 | רבים | , |
| 6:12.104 | ולשום |  |
| 6:12.105 | עיקר |  |
| 6:12.106 | העבודה |  |
| 6:12.107 | לבן |  |
| 6:12.108 | שהוא |  |
| 6:12.109 | זעיר |  |
| 6:12.110 | אנפין | , |
| 6:12.111 | וברוב |  |
| 6:12.112 | המצות |  |
| 6:12.113 | מכוינים |  |
| 6:12.114 | בהם |  |
| 6:12.115 | לקשט |  |
| 6:12.116 | הנשים |  |
| 6:12.117 | שכינתא |  |
| 6:12.118 | עלאה |  |
| 6:12.119 | ושכינתא |  |
| 6:12.120 | תתאה | , |
| 6:12.121 | כדי |  |
| 6:12.122 | למצוא |  |
| 6:12.123 | חן |  |
| 6:12.124 | בעיני |  |
| 6:12.125 | בעליהן |  |
| 6:12.126 | לעוררם |  |
| 6:12.127 | לתשמיש |  |
| 6:12.128 | כמו |  |
| 6:12.129 | שכ | ' |
| 6:12.130 | הר | ' |
| 6:12.131 | משנת |  |
| 6:12.132 | חסידים |  |
| 6:12.133 | בתיקון |  |
| 6:12.134 | חצות |  |
| 6:12.135 | וז"ל |  |
| 6:12.136 | וכיצד |  |
| 6:12.137 | יעשה |  |
| 6:12.138 | ישב |  |
| 6:12.139 | על |  |
| 6:12.140 | הארץ |  |
| 6:12.141 | אצל |  |
| 6:12.142 | המזוזה |  |
| 6:12.143 | שבפתח |  |
| 6:12.144 | שרומזת |  |
| 6:12.145 | להח | ' |
| 6:12.146 | ורגליו |  |
| 6:12.147 | יחפים |  |
| 6:12.148 | ממנעל | , |
| 6:12.149 | ויתן |  |
| 6:12.150 | אפר |  |
| 6:12.151 | מקלה |  |
| 6:12.152 | וכו | ' |
| 6:12.153 | ועל |  |
| 6:12.154 | ידי |  |
| 6:12.155 | זה |  |
| 6:12.156 | מתחמם |  |
| 6:12.157 | הזעיר |  |
| 6:12.158 | אנפין |  |
| 6:12.159 | ומתאוה |  |
| 6:12.160 | לה |  |
| 6:12.161 | וכו | ') |
| 6:12.162 | זה |  |
| 6:12.163 | להוליד |  |
| 6:12.164 | נשמות |  |
| 6:12.165 | למלאכים | . |
| 6:12.166 | וזה |  |
| 6:12.167 | להוליד |  |
| 6:12.168 | נשמות |  |
| 6:12.169 | חדשות |  |
| 6:12.170 | לישראל |  |
| 6:12.171 | זהר |  |
| 6:12.172 | יתרו |  |
| 6:12.173 | דף |  |
| 6:12.174 | פ"ט |  |
| 6:12.175 | ע"כ | ) |
| 6:12.176 | שלא |  |
| 6:12.177 | כרצון |  |
| 6:12.178 | חכמים |  |
| 6:12.179 | בבראשית |  |
| 6:12.180 | רבה |  |
| 6:12.181 | הבאתיו |  |
| 6:12.182 | לעיל |  |
| 6:12.183 | סי | ' |
| 6:12.184 | ל"א |  |
| 6:12.185 | כל |  |
| 6:12.186 | שיש |  |
| 6:12.187 | לו |  |
| 6:12.188 | תולדות |  |
| 6:12.189 | וכו | ') |
| 6:12.190 | ויחס |  |
| 6:12.191 | עיקר |  |
| 6:12.192 | העבודה |  |
| 6:12.193 | וקיום |  |
| 6:12.194 | המצות |  |
| 6:12.195 | והברכות |  |
| 6:12.196 | לבן |  |
| 6:12.197 | כנז"ל | . |
| 6:12.198 | ועוד |  |
| 6:12.199 | יתבאר |  |
| 6:12.200 | לקמן | . |
| 6:12.201 | ועל |  |
| 6:12.202 | ידי |  |
| 6:12.203 | רעיונותיהם |  |
| 6:12.204 | העולים |  |
| 6:12.205 | על |  |
| 6:12.206 | רוחם |  |
| 6:12.207 | חשבו |  |
| 6:12.208 | לפרש |  |
| 6:12.209 | הכתובים |  |
| 6:12.210 | שנאמרו |  |
| 6:12.211 | על |  |
| 6:12.212 | דרך |  |
| 6:12.213 | המליצה |  |
| 6:12.214 | והרחבת |  |
| 6:12.215 | הלשון |  |
| 6:12.216 | או |  |
| 6:12.217 | בדרך |  |
| 6:12.218 | חיבה |  |
| 6:12.219 | כמו |  |
| 6:12.220 | שאמרו |  |
| 6:12.221 | רז"ל |  |
| 6:12.222 | לא |  |
| 6:12.223 | זז |  |
| 6:12.224 | מחבבה |  |
| 6:12.225 | וכו | ' |
| 6:12.226 | כמשמען |  |
| 6:12.227 | והנחתן |  |
| 6:12.228 | הראשונה |  |
| 6:12.229 | יד |  |
| 6:12.230 | ה | ', |
| 6:12.231 | עיני |  |
| 6:12.232 | ה | ', |
| 6:12.233 | יעשן |  |
| 6:12.234 | אף |  |
| 6:12.235 | ה | ' |
| 6:12.236 | עלה |  |
| 6:12.237 | עשן |  |
| 6:12.238 | באפו | , |
| 6:12.239 | בצלמנו |  |
| 6:12.240 | כדמותנו | , |
| 6:12.241 | בנים |  |
| 6:12.242 | אתם |  |
| 6:12.243 | לה | ' |
| 6:12.244 | אלהיכם | , |
| 6:12.245 | אחי |  |
| 6:12.246 | ורעי | , |
| 6:12.247 | ויראו |  |
| 6:12.248 | את |  |
| 6:12.249 | אלוהי |  |
| 6:12.250 | ישראל | , |
| 6:12.251 | עורה |  |
| 6:12.252 | למה |  |
| 6:12.253 | תישן |  |
| 6:12.254 | ה | ', |
| 6:12.255 | הנה |  |
| 6:12.256 | לא |  |
| 6:12.257 | ינום |  |
| 6:12.258 | ולא |  |
| 6:12.259 | יישן |  |
| 6:12.260 | שומר |  |
| 6:12.261 | ישראל | , |
| 6:12.262 | וכיוצא |  |
| 6:12.263 | בהן | , |
| 6:12.264 | יפרשו |  |
| 6:12.265 | אותן |  |
| 6:12.266 | על |  |
| 6:12.267 | ידי |  |
| 6:12.268 | רעיונות |  |
| 6:12.269 | אלו |  |
| 6:12.270 | כהנחתן |  |
| 6:12.271 | הראשונה | , |
| 6:12.272 | היפך |  |
| 6:12.273 | מאמר |  |
| 6:12.274 | רז"ל |  |
| 6:12.275 | המתרגם |  |
| 6:12.276 | פסוק |  |
| 6:12.277 | כצורתו |  |
| 6:12.278 | הרי |  |
| 6:12.279 | זה |  |
| 6:12.280 | בדאי | . |
| 6:12.281 | וכמו |  |
| 6:12.282 | שכ | ' |
| 6:12.283 | הרס"ג |  |
| 6:12.284 | בהאמונות |  |
| 6:12.285 | והדעות |  |
| 6:12.286 | מאמר |  |
| 6:12.287 | האחדות | . |
| 6:12.288 | והרמב"ם |  |
| 6:12.289 | במורה |  |
| 6:12.290 | ח"א |  |
| 6:12.291 | וזולתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

