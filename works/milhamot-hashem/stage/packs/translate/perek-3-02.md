# Translation pack: milhamot-hashem perek-3-02 (Milchamot ha-Shem 3:3–3:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-02`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רוקח** → "Roke'ach" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הכוזרי / כוזרי** → "the Kuzari" (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **אלצלאם**: ונור אלצלאם (a misspelt title) as printed ("Nur al-Tzalam"), the
  likelier (Nur al-Zulam, Me'or ha-Afelah) in the note and `apparatus`.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **תימן** → "Yemen".
- **בא"י** → "in the Land of Israel".

## Terms

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
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- תורה **שבעל** פה → "the Oral Torah".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **תוה"ק** → "the holy Torah".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **מורא** → "fear".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **בחבלי** השוא … כעבות העגלה → "with cords of falsehood … as with
  cart-ropes".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 3:3

Hebrew (Wikisource): ד) ולא יהא כח רבותינו בעלי הקבלה החדשה גדול מכח ב"ד הגדול היושבים בלשכת הגזית ולא השגתם יתירה מהשגת סנהדרי גדולה! ועם כל זה אמרה תורה שאם טעו והורו להתיר את האסור, או שגגו בע"ז וכו'. למדנו מן התורה ומרבותינו במסכת הוריות, שאפשר לב"ד הגדול לטעות ולהורות התר באיזה פרט בע"ז. כי האדם קורץ מחומר בין גדול בין קטן, ועלול הוא לטעות, ולכן צותה תורה כי יפלא ממך דבר למשפט וכו' וקמת ועלית אל המקום אשר יבחר ה' אלהיך כדי למעט הטעות. איך אינו נמנע שלא יטעו אפי' ב"ד הגדול. שהרי נאמר בתורה ונעלם דבר מעיני הקהל. והיה אם מעיני העדה נעשתה לשגגה. ולפי קוצר דעתינו בראותינו ההפרש הגדול שיש בין רבותינו הראשונים ואיך היתה אמונתם בענין האלוהות ויחוד השי"ת על פי תוה"ק כמבואר בר' בחיי בעל חובת הלבבות ור' יהודה הלוי בספר הכוזרי ובתפלותיו ור' סעדיה גאון בס' האמונות והדעות והרמב"ם בחיבורו ובפי' המשנה ובס' מורה הנבוכים. והרוקח בשער הייחוד, וס' מצות הגדול ורבי יוסף אלבו בספר העיקרים, ור' מאיר אלדבי בספר שבילי אמונה. וזולתם רבים וכן שלימים. ואיך הוא עתה בספרי רבותינו האחרונים ע"פ הקבלה החדשה הנמשכת מתחלת אלף השישי בחבלי השוא עד שנעשית בזמנינו כעבות העגלה והאמינו באלהים רבים. ושהעיקר בהם אשר אליו כל עבודתינו וברכותינו הוא הפרצוף האחרון שבאצילות הנקרא (קצר אפים) זעיר אנפין.

| id | word | punctuation after |
|---|---|---|
| 3:3.1 | ד | ) |
| 3:3.2 | ולא |  |
| 3:3.3 | יהא |  |
| 3:3.4 | כח |  |
| 3:3.5 | רבותינו |  |
| 3:3.6 | בעלי |  |
| 3:3.7 | הקבלה |  |
| 3:3.8 | החדשה |  |
| 3:3.9 | גדול |  |
| 3:3.10 | מכח |  |
| 3:3.11 | ב"ד |  |
| 3:3.12 | הגדול |  |
| 3:3.13 | היושבים |  |
| 3:3.14 | בלשכת |  |
| 3:3.15 | הגזית |  |
| 3:3.16 | ולא |  |
| 3:3.17 | השגתם |  |
| 3:3.18 | יתירה |  |
| 3:3.19 | מהשגת |  |
| 3:3.20 | סנהדרי |  |
| 3:3.21 | גדולה | ! |
| 3:3.22 | ועם |  |
| 3:3.23 | כל |  |
| 3:3.24 | זה |  |
| 3:3.25 | אמרה |  |
| 3:3.26 | תורה |  |
| 3:3.27 | שאם |  |
| 3:3.28 | טעו |  |
| 3:3.29 | והורו |  |
| 3:3.30 | להתיר |  |
| 3:3.31 | את |  |
| 3:3.32 | האסור | , |
| 3:3.33 | או |  |
| 3:3.34 | שגגו |  |
| 3:3.35 | בע"ז |  |
| 3:3.36 | וכו | '. |
| 3:3.37 | למדנו |  |
| 3:3.38 | מן |  |
| 3:3.39 | התורה |  |
| 3:3.40 | ומרבותינו |  |
| 3:3.41 | במסכת |  |
| 3:3.42 | הוריות | , |
| 3:3.43 | שאפשר |  |
| 3:3.44 | לב"ד |  |
| 3:3.45 | הגדול |  |
| 3:3.46 | לטעות |  |
| 3:3.47 | ולהורות |  |
| 3:3.48 | התר |  |
| 3:3.49 | באיזה |  |
| 3:3.50 | פרט |  |
| 3:3.51 | בע"ז | . |
| 3:3.52 | כי |  |
| 3:3.53 | האדם |  |
| 3:3.54 | קורץ |  |
| 3:3.55 | מחומר |  |
| 3:3.56 | בין |  |
| 3:3.57 | גדול |  |
| 3:3.58 | בין |  |
| 3:3.59 | קטן | , |
| 3:3.60 | ועלול |  |
| 3:3.61 | הוא |  |
| 3:3.62 | לטעות | , |
| 3:3.63 | ולכן |  |
| 3:3.64 | צותה |  |
| 3:3.65 | תורה |  |
| 3:3.66 | כי |  |
| 3:3.67 | יפלא |  |
| 3:3.68 | ממך |  |
| 3:3.69 | דבר |  |
| 3:3.70 | למשפט |  |
| 3:3.71 | וכו | ' |
| 3:3.72 | וקמת |  |
| 3:3.73 | ועלית |  |
| 3:3.74 | אל |  |
| 3:3.75 | המקום |  |
| 3:3.76 | אשר |  |
| 3:3.77 | יבחר |  |
| 3:3.78 | ה | ' |
| 3:3.79 | אלהיך |  |
| 3:3.80 | כדי |  |
| 3:3.81 | למעט |  |
| 3:3.82 | הטעות | . |
| 3:3.83 | איך |  |
| 3:3.84 | אינו |  |
| 3:3.85 | נמנע |  |
| 3:3.86 | שלא |  |
| 3:3.87 | יטעו |  |
| 3:3.88 | אפי | ' |
| 3:3.89 | ב"ד |  |
| 3:3.90 | הגדול | . |
| 3:3.91 | שהרי |  |
| 3:3.92 | נאמר |  |
| 3:3.93 | בתורה |  |
| 3:3.94 | ונעלם |  |
| 3:3.95 | דבר |  |
| 3:3.96 | מעיני |  |
| 3:3.97 | הקהל | . |
| 3:3.98 | והיה |  |
| 3:3.99 | אם |  |
| 3:3.100 | מעיני |  |
| 3:3.101 | העדה |  |
| 3:3.102 | נעשתה |  |
| 3:3.103 | לשגגה | . |
| 3:3.104 | ולפי |  |
| 3:3.105 | קוצר |  |
| 3:3.106 | דעתינו |  |
| 3:3.107 | בראותינו |  |
| 3:3.108 | ההפרש |  |
| 3:3.109 | הגדול |  |
| 3:3.110 | שיש |  |
| 3:3.111 | בין |  |
| 3:3.112 | רבותינו |  |
| 3:3.113 | הראשונים |  |
| 3:3.114 | ואיך |  |
| 3:3.115 | היתה |  |
| 3:3.116 | אמונתם |  |
| 3:3.117 | בענין |  |
| 3:3.118 | האלוהות |  |
| 3:3.119 | ויחוד |  |
| 3:3.120 | השי"ת |  |
| 3:3.121 | על |  |
| 3:3.122 | פי |  |
| 3:3.123 | תוה"ק |  |
| 3:3.124 | כמבואר |  |
| 3:3.125 | בר | ' |
| 3:3.126 | בחיי |  |
| 3:3.127 | בעל |  |
| 3:3.128 | חובת |  |
| 3:3.129 | הלבבות |  |
| 3:3.130 | ור | ' |
| 3:3.131 | יהודה |  |
| 3:3.132 | הלוי |  |
| 3:3.133 | בספר |  |
| 3:3.134 | הכוזרי |  |
| 3:3.135 | ובתפלותיו |  |
| 3:3.136 | ור | ' |
| 3:3.137 | סעדיה |  |
| 3:3.138 | גאון |  |
| 3:3.139 | בס | ' |
| 3:3.140 | האמונות |  |
| 3:3.141 | והדעות |  |
| 3:3.142 | והרמב"ם |  |
| 3:3.143 | בחיבורו |  |
| 3:3.144 | ובפי | ' |
| 3:3.145 | המשנה |  |
| 3:3.146 | ובס | ' |
| 3:3.147 | מורה |  |
| 3:3.148 | הנבוכים | . |
| 3:3.149 | והרוקח |  |
| 3:3.150 | בשער |  |
| 3:3.151 | הייחוד | , |
| 3:3.152 | וס | ' |
| 3:3.153 | מצות |  |
| 3:3.154 | הגדול |  |
| 3:3.155 | ורבי |  |
| 3:3.156 | יוסף |  |
| 3:3.157 | אלבו |  |
| 3:3.158 | בספר |  |
| 3:3.159 | העיקרים | , |
| 3:3.160 | ור | ' |
| 3:3.161 | מאיר |  |
| 3:3.162 | אלדבי |  |
| 3:3.163 | בספר |  |
| 3:3.164 | שבילי |  |
| 3:3.165 | אמונה | . |
| 3:3.166 | וזולתם |  |
| 3:3.167 | רבים |  |
| 3:3.168 | וכן |  |
| 3:3.169 | שלימים | . |
| 3:3.170 | ואיך |  |
| 3:3.171 | הוא |  |
| 3:3.172 | עתה |  |
| 3:3.173 | בספרי |  |
| 3:3.174 | רבותינו |  |
| 3:3.175 | האחרונים |  |
| 3:3.176 | ע"פ |  |
| 3:3.177 | הקבלה |  |
| 3:3.178 | החדשה |  |
| 3:3.179 | הנמשכת |  |
| 3:3.180 | מתחלת |  |
| 3:3.181 | אלף |  |
| 3:3.182 | השישי |  |
| 3:3.183 | בחבלי |  |
| 3:3.184 | השוא |  |
| 3:3.185 | עד |  |
| 3:3.186 | שנעשית |  |
| 3:3.187 | בזמנינו |  |
| 3:3.188 | כעבות |  |
| 3:3.189 | העגלה |  |
| 3:3.190 | והאמינו |  |
| 3:3.191 | באלהים |  |
| 3:3.192 | רבים | . |
| 3:3.193 | ושהעיקר |  |
| 3:3.194 | בהם |  |
| 3:3.195 | אשר |  |
| 3:3.196 | אליו |  |
| 3:3.197 | כל |  |
| 3:3.198 | עבודתינו |  |
| 3:3.199 | וברכותינו |  |
| 3:3.200 | הוא |  |
| 3:3.201 | הפרצוף |  |
| 3:3.202 | האחרון |  |
| 3:3.203 | שבאצילות |  |
| 3:3.204 | הנקרא |  |
| 3:3.205 | קצר |  |
| 3:3.206 | אפים | ) |
| 3:3.207 | זעיר |  |
| 3:3.208 | אנפין | . |


## Milchamot ha-Shem 3:4

Hebrew (Wikisource): ה) ואמונה זו החדשה רחוקה מאד מדרך תורתינו הק' שבכתב ושבעל פה. ולכן פירשנו עצמינו ממנה ודבקנו בדברי רבותינו ז"ל בעלי המשנה והתלמוד בבלי וירושלמי והפוסקים ומדרשי רבותינו המפורסמים, מדרש רבה ותנחומא וילקוט שמעוני ומדרש הגדול המצוי בתימן בכת"י, ונור אלצלאם וכיוצא בהם. גם שארית חבורי קדמונינו חכמי תימן מצינו וראינו שהם אדוקים בדברי הגדולים הנזכרים לעיל ובדעותיהם. וכראותכם שפירשנו מן הקבלה החדשה חרה אפכם בנו, וברמות רוחא טפלתם עלינו שקר זדים. ושאלנו אתכם להורותינו הדרך איך יסכימו דברי רבותינו המקובלים החדשים עם דברי רבותינו הראשונים הנזכרים לעיל, ועם דברי התנאים והאמוראים. והיתה תשובתכם אלינו בדרך הטעאה וקינטורים.

| id | word | punctuation after |
|---|---|---|
| 3:4.1 | ה | ) |
| 3:4.2 | ואמונה |  |
| 3:4.3 | זו |  |
| 3:4.4 | החדשה |  |
| 3:4.5 | רחוקה |  |
| 3:4.6 | מאד |  |
| 3:4.7 | מדרך |  |
| 3:4.8 | תורתינו |  |
| 3:4.9 | הק | ' |
| 3:4.10 | שבכתב |  |
| 3:4.11 | ושבעל |  |
| 3:4.12 | פה | . |
| 3:4.13 | ולכן |  |
| 3:4.14 | פירשנו |  |
| 3:4.15 | עצמינו |  |
| 3:4.16 | ממנה |  |
| 3:4.17 | ודבקנו |  |
| 3:4.18 | בדברי |  |
| 3:4.19 | רבותינו |  |
| 3:4.20 | ז"ל |  |
| 3:4.21 | בעלי |  |
| 3:4.22 | המשנה |  |
| 3:4.23 | והתלמוד |  |
| 3:4.24 | בבלי |  |
| 3:4.25 | וירושלמי |  |
| 3:4.26 | והפוסקים |  |
| 3:4.27 | ומדרשי |  |
| 3:4.28 | רבותינו |  |
| 3:4.29 | המפורסמים | , |
| 3:4.30 | מדרש |  |
| 3:4.31 | רבה |  |
| 3:4.32 | ותנחומא |  |
| 3:4.33 | וילקוט |  |
| 3:4.34 | שמעוני |  |
| 3:4.35 | ומדרש |  |
| 3:4.36 | הגדול |  |
| 3:4.37 | המצוי |  |
| 3:4.38 | בתימן |  |
| 3:4.39 | בכת"י | , |
| 3:4.40 | ונור |  |
| 3:4.41 | אלצלאם |  |
| 3:4.42 | וכיוצא |  |
| 3:4.43 | בהם | . |
| 3:4.44 | גם |  |
| 3:4.45 | שארית |  |
| 3:4.46 | חבורי |  |
| 3:4.47 | קדמונינו |  |
| 3:4.48 | חכמי |  |
| 3:4.49 | תימן |  |
| 3:4.50 | מצינו |  |
| 3:4.51 | וראינו |  |
| 3:4.52 | שהם |  |
| 3:4.53 | אדוקים |  |
| 3:4.54 | בדברי |  |
| 3:4.55 | הגדולים |  |
| 3:4.56 | הנזכרים |  |
| 3:4.57 | לעיל |  |
| 3:4.58 | ובדעותיהם | . |
| 3:4.59 | וכראותכם |  |
| 3:4.60 | שפירשנו |  |
| 3:4.61 | מן |  |
| 3:4.62 | הקבלה |  |
| 3:4.63 | החדשה |  |
| 3:4.64 | חרה |  |
| 3:4.65 | אפכם |  |
| 3:4.66 | בנו | , |
| 3:4.67 | וברמות |  |
| 3:4.68 | רוחא |  |
| 3:4.69 | טפלתם |  |
| 3:4.70 | עלינו |  |
| 3:4.71 | שקר |  |
| 3:4.72 | זדים | . |
| 3:4.73 | ושאלנו |  |
| 3:4.74 | אתכם |  |
| 3:4.75 | להורותינו |  |
| 3:4.76 | הדרך |  |
| 3:4.77 | איך |  |
| 3:4.78 | יסכימו |  |
| 3:4.79 | דברי |  |
| 3:4.80 | רבותינו |  |
| 3:4.81 | המקובלים |  |
| 3:4.82 | החדשים |  |
| 3:4.83 | עם |  |
| 3:4.84 | דברי |  |
| 3:4.85 | רבותינו |  |
| 3:4.86 | הראשונים |  |
| 3:4.87 | הנזכרים |  |
| 3:4.88 | לעיל | , |
| 3:4.89 | ועם |  |
| 3:4.90 | דברי |  |
| 3:4.91 | התנאים |  |
| 3:4.92 | והאמוראים | . |
| 3:4.93 | והיתה |  |
| 3:4.94 | תשובתכם |  |
| 3:4.95 | אלינו |  |
| 3:4.96 | בדרך |  |
| 3:4.97 | הטעאה |  |
| 3:4.98 | וקינטורים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

