# Translation pack: milhamot-hashem perek-9-18 (Milchamot ha-Shem 9:21–9:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-18`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
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
- **מערב** beside Ashkenaz, Sefarad, Yemen → "the Maghreb" (note: "the West").
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- הסבה **הראשונה** → "the First Cause".
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
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **מין / מינות** → "heretic / heresy".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **תוה"ק** → "the holy Torah".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רחום / וחנון** (Exod 34:6) → "the LORD, the LORD, God merciful and
  gracious, slow to anger".
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


## Milchamot ha-Shem 9:21

Hebrew (Wikisource): קלט) לפיכך כל מי שלבו שלם עם ה' אלהינו שהוא הסבה הראשונה לכל הנמצאים בלי שום שתוף כח אחר, אל ישמע לדברי המסית מחבר הזהר והנמשכים אחריו! כי ה' אלהינו הוא לבדו הנמציא כל הנמצאים ברוב טובו וחסדו, והוא היחיד ומיוחד מכל שאר האחדים, והוא לבדו ראשון שאין ראשית לראשיתו, והוא שהוציאנו ממצרים, והוא שנגלה בסיני ונתן לנו את התורה בקולות וברקים ולפידים ורול שופר, ורעשו כל ההרים ונזדעזע כל העולם לקול נבורתו, ונתן בקולו קול עוז אנכי ה' אלהיך אשר הוציאתיך, לא היה לך אלהים אחרים על פני, וכל עשרה דברים הנאמרים במעמד הנבחר ההוא בפרסום גדול, עד שכל מלכי מזרח ומערב חרדו ותמהו, והלכו אצל בלעם נביאם להודיעם מה זה, כמו שאמרו רז"ל במדרש, ואבותינו כולם ששים רבוא איש רגלי מלבד הזקינים והנשים והטף כולם נצבים בתחתית הההר הר סיני לקבל תורה ומצות, וענו כולם ואמרו כל אשר דבר ה' נעשה ונשמע, ומיראה ופחד גדול מקול ה' בכח, אמרו למשה רבינו דבר אתה עמנו ונשמעה וכו'. וכל זה עשה הקב"ה עמנו כדי להאמין במשה בחירו שתורת אמת נתן לנו, ולא יחליפה ולא יחליפה ולא ימירה וגם אמר את ה' האמרת וכו' וה' האמירך היום להיות לו לעם סגלה, ולא נאמין לשום נביא או קוסם שיפתה אותנו ל עבוד אלוה אחר מבלעדיו, רק אותו לבדו נעבוד שלי שום שתוף דבר אחר עמו, והוא הודיע דרכיו למשה. כי הוא חנון ורחום ארף אפים וכו' ולא קצר אפים ואם כן איך נעזוב את אלהינו ארך אפים ונעבוד לאלוה קצא אפים? וגם אם יעשה נביא השקר אותות ומופתים לעינינו בשמים ובארץ. כמו שנאמר ונתן אליך אות או מופת וכי' ובא האות והמופת אשר דבר אליך לאמר נלכה אחרי אלהים אחרים אשר לא ידעתם (לא בנביאנו ולא מרז"ל) ונעבדם. לא תשמע אל דברי הנביא ההוא וכו' כל שכן וק"ו שלא נשמע לדברי הפילסוף מחבר הזהר, אשר זה מקרוב צמח נן הארץ ולא נודע לו בית אב מי הוא זה מחברו באמת, כשאר ספרי רז"ל הידועים ומפורסמים מחבריהם מיום שחוברו על ידי גדולי ישראל מעתיקי השמועה ועד עתה. וגדול ועצום הוא הק"ו הזה. ומה אם נביא העושה אותות ומופתים לענינינו בשמים ובארץ כמו שפירשו רז"ל ומביאו רש"י ז"ל, אמרה תורה לא תשמע אל דברי הנביא ההוא כי מנסה ה' אלהיכם אתכם? ספר הזהר הבא לנו מיד עכו"ם, ומשמיע גדולות ונוראות משם תנאים ואמוראים ונשמות צדיקים שקדמו לרשב"י כמה דורות. ואמוראים המאוחרים מן רב מאד אחר רשב"י, ואליהו ועתיקא קדישא שנגלה לנו בבית מדרשו (בין חביריו ותלמידיו) וכיוצא בספורים אלו הרשומים בדיו שחור על נייר לבן, אשר לא ראינו בעינינו ולא אבותינו ספרו לנו כל אותן הדברים והמעשיות הרשומים שם! כ"ש שלא נשמע בליו לעבוד שום פרצוף מכל הפרצופים, ובפרט קצר אפים (ז"א) כי אם את אלהינו בהו"ש אשר קבלנו מאבותינו, ותוה"ק ספרה לנו שאין לו דמות הגוף, ואין ראשית לראשיתו, והוא אחד ואין יחוד כיחודו לפי שכל הפרצופים והעגולים כולם יש ראשית לראשיתם, ואין לעבוד רק לצור הכל בלי שום שתוף נברא או נאצל מכל הנבראים והנאצלים, וגם אין ליחד ולקשר עם אחדותו שום אחד מהם, כמאמר רשב"י כל המשתף וכו'.

| id | word | punctuation after |
|---|---|---|
| 9:21.1 | קלט | ) |
| 9:21.2 | לפיכך |  |
| 9:21.3 | כל |  |
| 9:21.4 | מי |  |
| 9:21.5 | שלבו |  |
| 9:21.6 | שלם |  |
| 9:21.7 | עם |  |
| 9:21.8 | ה | ' |
| 9:21.9 | אלהינו |  |
| 9:21.10 | שהוא |  |
| 9:21.11 | הסבה |  |
| 9:21.12 | הראשונה |  |
| 9:21.13 | לכל |  |
| 9:21.14 | הנמצאים |  |
| 9:21.15 | בלי |  |
| 9:21.16 | שום |  |
| 9:21.17 | שתוף |  |
| 9:21.18 | כח |  |
| 9:21.19 | אחר | , |
| 9:21.20 | אל |  |
| 9:21.21 | ישמע |  |
| 9:21.22 | לדברי |  |
| 9:21.23 | המסית |  |
| 9:21.24 | מחבר |  |
| 9:21.25 | הזהר |  |
| 9:21.26 | והנמשכים |  |
| 9:21.27 | אחריו | ! |
| 9:21.28 | כי |  |
| 9:21.29 | ה | ' |
| 9:21.30 | אלהינו |  |
| 9:21.31 | הוא |  |
| 9:21.32 | לבדו |  |
| 9:21.33 | הנמציא |  |
| 9:21.34 | כל |  |
| 9:21.35 | הנמצאים |  |
| 9:21.36 | ברוב |  |
| 9:21.37 | טובו |  |
| 9:21.38 | וחסדו | , |
| 9:21.39 | והוא |  |
| 9:21.40 | היחיד |  |
| 9:21.41 | ומיוחד |  |
| 9:21.42 | מכל |  |
| 9:21.43 | שאר |  |
| 9:21.44 | האחדים | , |
| 9:21.45 | והוא |  |
| 9:21.46 | לבדו |  |
| 9:21.47 | ראשון |  |
| 9:21.48 | שאין |  |
| 9:21.49 | ראשית |  |
| 9:21.50 | לראשיתו | , |
| 9:21.51 | והוא |  |
| 9:21.52 | שהוציאנו |  |
| 9:21.53 | ממצרים | , |
| 9:21.54 | והוא |  |
| 9:21.55 | שנגלה |  |
| 9:21.56 | בסיני |  |
| 9:21.57 | ונתן |  |
| 9:21.58 | לנו |  |
| 9:21.59 | את |  |
| 9:21.60 | התורה |  |
| 9:21.61 | בקולות |  |
| 9:21.62 | וברקים |  |
| 9:21.63 | ולפידים |  |
| 9:21.64 | ורול |  |
| 9:21.65 | שופר | , |
| 9:21.66 | ורעשו |  |
| 9:21.67 | כל |  |
| 9:21.68 | ההרים |  |
| 9:21.69 | ונזדעזע |  |
| 9:21.70 | כל |  |
| 9:21.71 | העולם |  |
| 9:21.72 | לקול |  |
| 9:21.73 | נבורתו | , |
| 9:21.74 | ונתן |  |
| 9:21.75 | בקולו |  |
| 9:21.76 | קול |  |
| 9:21.77 | עוז |  |
| 9:21.78 | אנכי |  |
| 9:21.79 | ה | ' |
| 9:21.80 | אלהיך |  |
| 9:21.81 | אשר |  |
| 9:21.82 | הוציאתיך | , |
| 9:21.83 | לא |  |
| 9:21.84 | היה |  |
| 9:21.85 | לך |  |
| 9:21.86 | אלהים |  |
| 9:21.87 | אחרים |  |
| 9:21.88 | על |  |
| 9:21.89 | פני | , |
| 9:21.90 | וכל |  |
| 9:21.91 | עשרה |  |
| 9:21.92 | דברים |  |
| 9:21.93 | הנאמרים |  |
| 9:21.94 | במעמד |  |
| 9:21.95 | הנבחר |  |
| 9:21.96 | ההוא |  |
| 9:21.97 | בפרסום |  |
| 9:21.98 | גדול | , |
| 9:21.99 | עד |  |
| 9:21.100 | שכל |  |
| 9:21.101 | מלכי |  |
| 9:21.102 | מזרח |  |
| 9:21.103 | ומערב |  |
| 9:21.104 | חרדו |  |
| 9:21.105 | ותמהו | , |
| 9:21.106 | והלכו |  |
| 9:21.107 | אצל |  |
| 9:21.108 | בלעם |  |
| 9:21.109 | נביאם |  |
| 9:21.110 | להודיעם |  |
| 9:21.111 | מה |  |
| 9:21.112 | זה | , |
| 9:21.113 | כמו |  |
| 9:21.114 | שאמרו |  |
| 9:21.115 | רז"ל |  |
| 9:21.116 | במדרש | , |
| 9:21.117 | ואבותינו |  |
| 9:21.118 | כולם |  |
| 9:21.119 | ששים |  |
| 9:21.120 | רבוא |  |
| 9:21.121 | איש |  |
| 9:21.122 | רגלי |  |
| 9:21.123 | מלבד |  |
| 9:21.124 | הזקינים |  |
| 9:21.125 | והנשים |  |
| 9:21.126 | והטף |  |
| 9:21.127 | כולם |  |
| 9:21.128 | נצבים |  |
| 9:21.129 | בתחתית |  |
| 9:21.130 | הההר |  |
| 9:21.131 | הר |  |
| 9:21.132 | סיני |  |
| 9:21.133 | לקבל |  |
| 9:21.134 | תורה |  |
| 9:21.135 | ומצות | , |
| 9:21.136 | וענו |  |
| 9:21.137 | כולם |  |
| 9:21.138 | ואמרו |  |
| 9:21.139 | כל |  |
| 9:21.140 | אשר |  |
| 9:21.141 | דבר |  |
| 9:21.142 | ה | ' |
| 9:21.143 | נעשה |  |
| 9:21.144 | ונשמע | , |
| 9:21.145 | ומיראה |  |
| 9:21.146 | ופחד |  |
| 9:21.147 | גדול |  |
| 9:21.148 | מקול |  |
| 9:21.149 | ה | ' |
| 9:21.150 | בכח | , |
| 9:21.151 | אמרו |  |
| 9:21.152 | למשה |  |
| 9:21.153 | רבינו |  |
| 9:21.154 | דבר |  |
| 9:21.155 | אתה |  |
| 9:21.156 | עמנו |  |
| 9:21.157 | ונשמעה |  |
| 9:21.158 | וכו | '. |
| 9:21.159 | וכל |  |
| 9:21.160 | זה |  |
| 9:21.161 | עשה |  |
| 9:21.162 | הקב"ה |  |
| 9:21.163 | עמנו |  |
| 9:21.164 | כדי |  |
| 9:21.165 | להאמין |  |
| 9:21.166 | במשה |  |
| 9:21.167 | בחירו |  |
| 9:21.168 | שתורת |  |
| 9:21.169 | אמת |  |
| 9:21.170 | נתן |  |
| 9:21.171 | לנו | , |
| 9:21.172 | ולא |  |
| 9:21.173 | יחליפה |  |
| 9:21.174 | ולא |  |
| 9:21.175 | יחליפה |  |
| 9:21.176 | ולא |  |
| 9:21.177 | ימירה |  |
| 9:21.178 | וגם |  |
| 9:21.179 | אמר |  |
| 9:21.180 | את |  |
| 9:21.181 | ה | ' |
| 9:21.182 | האמרת |  |
| 9:21.183 | וכו | ' |
| 9:21.184 | וה | ' |
| 9:21.185 | האמירך |  |
| 9:21.186 | היום |  |
| 9:21.187 | להיות |  |
| 9:21.188 | לו |  |
| 9:21.189 | לעם |  |
| 9:21.190 | סגלה | , |
| 9:21.191 | ולא |  |
| 9:21.192 | נאמין |  |
| 9:21.193 | לשום |  |
| 9:21.194 | נביא |  |
| 9:21.195 | או |  |
| 9:21.196 | קוסם |  |
| 9:21.197 | שיפתה |  |
| 9:21.198 | אותנו |  |
| 9:21.199 | ל |  |
| 9:21.200 | עבוד |  |
| 9:21.201 | אלוה |  |
| 9:21.202 | אחר |  |
| 9:21.203 | מבלעדיו | , |
| 9:21.204 | רק |  |
| 9:21.205 | אותו |  |
| 9:21.206 | לבדו |  |
| 9:21.207 | נעבוד |  |
| 9:21.208 | שלי |  |
| 9:21.209 | שום |  |
| 9:21.210 | שתוף |  |
| 9:21.211 | דבר |  |
| 9:21.212 | אחר |  |
| 9:21.213 | עמו | , |
| 9:21.214 | והוא |  |
| 9:21.215 | הודיע |  |
| 9:21.216 | דרכיו |  |
| 9:21.217 | למשה | . |
| 9:21.218 | כי |  |
| 9:21.219 | הוא |  |
| 9:21.220 | חנון |  |
| 9:21.221 | ורחום |  |
| 9:21.222 | ארף |  |
| 9:21.223 | אפים |  |
| 9:21.224 | וכו | ' |
| 9:21.225 | ולא |  |
| 9:21.226 | קצר |  |
| 9:21.227 | אפים |  |
| 9:21.228 | ואם |  |
| 9:21.229 | כן |  |
| 9:21.230 | איך |  |
| 9:21.231 | נעזוב |  |
| 9:21.232 | את |  |
| 9:21.233 | אלהינו |  |
| 9:21.234 | ארך |  |
| 9:21.235 | אפים |  |
| 9:21.236 | ונעבוד |  |
| 9:21.237 | לאלוה |  |
| 9:21.238 | קצא |  |
| 9:21.239 | אפים | ? |
| 9:21.240 | וגם |  |
| 9:21.241 | אם |  |
| 9:21.242 | יעשה |  |
| 9:21.243 | נביא |  |
| 9:21.244 | השקר |  |
| 9:21.245 | אותות |  |
| 9:21.246 | ומופתים |  |
| 9:21.247 | לעינינו |  |
| 9:21.248 | בשמים |  |
| 9:21.249 | ובארץ | . |
| 9:21.250 | כמו |  |
| 9:21.251 | שנאמר |  |
| 9:21.252 | ונתן |  |
| 9:21.253 | אליך |  |
| 9:21.254 | אות |  |
| 9:21.255 | או |  |
| 9:21.256 | מופת |  |
| 9:21.257 | וכי | ' |
| 9:21.258 | ובא |  |
| 9:21.259 | האות |  |
| 9:21.260 | והמופת |  |
| 9:21.261 | אשר |  |
| 9:21.262 | דבר |  |
| 9:21.263 | אליך |  |
| 9:21.264 | לאמר |  |
| 9:21.265 | נלכה |  |
| 9:21.266 | אחרי |  |
| 9:21.267 | אלהים |  |
| 9:21.268 | אחרים |  |
| 9:21.269 | אשר |  |
| 9:21.270 | לא |  |
| 9:21.271 | ידעתם |  |
| 9:21.272 | לא |  |
| 9:21.273 | בנביאנו |  |
| 9:21.274 | ולא |  |
| 9:21.275 | מרז"ל | ) |
| 9:21.276 | ונעבדם | . |
| 9:21.277 | לא |  |
| 9:21.278 | תשמע |  |
| 9:21.279 | אל |  |
| 9:21.280 | דברי |  |
| 9:21.281 | הנביא |  |
| 9:21.282 | ההוא |  |
| 9:21.283 | וכו | ' |
| 9:21.284 | כל |  |
| 9:21.285 | שכן |  |
| 9:21.286 | וק"ו |  |
| 9:21.287 | שלא |  |
| 9:21.288 | נשמע |  |
| 9:21.289 | לדברי |  |
| 9:21.290 | הפילסוף |  |
| 9:21.291 | מחבר |  |
| 9:21.292 | הזהר | , |
| 9:21.293 | אשר |  |
| 9:21.294 | זה |  |
| 9:21.295 | מקרוב |  |
| 9:21.296 | צמח |  |
| 9:21.297 | נן |  |
| 9:21.298 | הארץ |  |
| 9:21.299 | ולא |  |
| 9:21.300 | נודע |  |
| 9:21.301 | לו |  |
| 9:21.302 | בית |  |
| 9:21.303 | אב |  |
| 9:21.304 | מי |  |
| 9:21.305 | הוא |  |
| 9:21.306 | זה |  |
| 9:21.307 | מחברו |  |
| 9:21.308 | באמת | , |
| 9:21.309 | כשאר |  |
| 9:21.310 | ספרי |  |
| 9:21.311 | רז"ל |  |
| 9:21.312 | הידועים |  |
| 9:21.313 | ומפורסמים |  |
| 9:21.314 | מחבריהם |  |
| 9:21.315 | מיום |  |
| 9:21.316 | שחוברו |  |
| 9:21.317 | על |  |
| 9:21.318 | ידי |  |
| 9:21.319 | גדולי |  |
| 9:21.320 | ישראל |  |
| 9:21.321 | מעתיקי |  |
| 9:21.322 | השמועה |  |
| 9:21.323 | ועד |  |
| 9:21.324 | עתה | . |
| 9:21.325 | וגדול |  |
| 9:21.326 | ועצום |  |
| 9:21.327 | הוא |  |
| 9:21.328 | הק"ו |  |
| 9:21.329 | הזה | . |
| 9:21.330 | ומה |  |
| 9:21.331 | אם |  |
| 9:21.332 | נביא |  |
| 9:21.333 | העושה |  |
| 9:21.334 | אותות |  |
| 9:21.335 | ומופתים |  |
| 9:21.336 | לענינינו |  |
| 9:21.337 | בשמים |  |
| 9:21.338 | ובארץ |  |
| 9:21.339 | כמו |  |
| 9:21.340 | שפירשו |  |
| 9:21.341 | רז"ל |  |
| 9:21.342 | ומביאו |  |
| 9:21.343 | רש"י |  |
| 9:21.344 | ז"ל | , |
| 9:21.345 | אמרה |  |
| 9:21.346 | תורה |  |
| 9:21.347 | לא |  |
| 9:21.348 | תשמע |  |
| 9:21.349 | אל |  |
| 9:21.350 | דברי |  |
| 9:21.351 | הנביא |  |
| 9:21.352 | ההוא |  |
| 9:21.353 | כי |  |
| 9:21.354 | מנסה |  |
| 9:21.355 | ה | ' |
| 9:21.356 | אלהיכם |  |
| 9:21.357 | אתכם | ? |
| 9:21.358 | ספר |  |
| 9:21.359 | הזהר |  |
| 9:21.360 | הבא |  |
| 9:21.361 | לנו |  |
| 9:21.362 | מיד |  |
| 9:21.363 | עכו"ם | , |
| 9:21.364 | ומשמיע |  |
| 9:21.365 | גדולות |  |
| 9:21.366 | ונוראות |  |
| 9:21.367 | משם |  |
| 9:21.368 | תנאים |  |
| 9:21.369 | ואמוראים |  |
| 9:21.370 | ונשמות |  |
| 9:21.371 | צדיקים |  |
| 9:21.372 | שקדמו |  |
| 9:21.373 | לרשב"י |  |
| 9:21.374 | כמה |  |
| 9:21.375 | דורות | . |
| 9:21.376 | ואמוראים |  |
| 9:21.377 | המאוחרים |  |
| 9:21.378 | מן |  |
| 9:21.379 | רב |  |
| 9:21.380 | מאד |  |
| 9:21.381 | אחר |  |
| 9:21.382 | רשב"י | , |
| 9:21.383 | ואליהו |  |
| 9:21.384 | ועתיקא |  |
| 9:21.385 | קדישא |  |
| 9:21.386 | שנגלה |  |
| 9:21.387 | לנו |  |
| 9:21.388 | בבית |  |
| 9:21.389 | מדרשו |  |
| 9:21.390 | בין |  |
| 9:21.391 | חביריו |  |
| 9:21.392 | ותלמידיו | ) |
| 9:21.393 | וכיוצא |  |
| 9:21.394 | בספורים |  |
| 9:21.395 | אלו |  |
| 9:21.396 | הרשומים |  |
| 9:21.397 | בדיו |  |
| 9:21.398 | שחור |  |
| 9:21.399 | על |  |
| 9:21.400 | נייר |  |
| 9:21.401 | לבן | , |
| 9:21.402 | אשר |  |
| 9:21.403 | לא |  |
| 9:21.404 | ראינו |  |
| 9:21.405 | בעינינו |  |
| 9:21.406 | ולא |  |
| 9:21.407 | אבותינו |  |
| 9:21.408 | ספרו |  |
| 9:21.409 | לנו |  |
| 9:21.410 | כל |  |
| 9:21.411 | אותן |  |
| 9:21.412 | הדברים |  |
| 9:21.413 | והמעשיות |  |
| 9:21.414 | הרשומים |  |
| 9:21.415 | שם | ! |
| 9:21.416 | כ"ש |  |
| 9:21.417 | שלא |  |
| 9:21.418 | נשמע |  |
| 9:21.419 | בליו |  |
| 9:21.420 | לעבוד |  |
| 9:21.421 | שום |  |
| 9:21.422 | פרצוף |  |
| 9:21.423 | מכל |  |
| 9:21.424 | הפרצופים | , |
| 9:21.425 | ובפרט |  |
| 9:21.426 | קצר |  |
| 9:21.427 | אפים |  |
| 9:21.428 | ז"א | ) |
| 9:21.429 | כי |  |
| 9:21.430 | אם |  |
| 9:21.431 | את |  |
| 9:21.432 | אלהינו |  |
| 9:21.433 | בהו"ש |  |
| 9:21.434 | אשר |  |
| 9:21.435 | קבלנו |  |
| 9:21.436 | מאבותינו | , |
| 9:21.437 | ותוה"ק |  |
| 9:21.438 | ספרה |  |
| 9:21.439 | לנו |  |
| 9:21.440 | שאין |  |
| 9:21.441 | לו |  |
| 9:21.442 | דמות |  |
| 9:21.443 | הגוף | , |
| 9:21.444 | ואין |  |
| 9:21.445 | ראשית |  |
| 9:21.446 | לראשיתו | , |
| 9:21.447 | והוא |  |
| 9:21.448 | אחד |  |
| 9:21.449 | ואין |  |
| 9:21.450 | יחוד |  |
| 9:21.451 | כיחודו |  |
| 9:21.452 | לפי |  |
| 9:21.453 | שכל |  |
| 9:21.454 | הפרצופים |  |
| 9:21.455 | והעגולים |  |
| 9:21.456 | כולם |  |
| 9:21.457 | יש |  |
| 9:21.458 | ראשית |  |
| 9:21.459 | לראשיתם | , |
| 9:21.460 | ואין |  |
| 9:21.461 | לעבוד |  |
| 9:21.462 | רק |  |
| 9:21.463 | לצור |  |
| 9:21.464 | הכל |  |
| 9:21.465 | בלי |  |
| 9:21.466 | שום |  |
| 9:21.467 | שתוף |  |
| 9:21.468 | נברא |  |
| 9:21.469 | או |  |
| 9:21.470 | נאצל |  |
| 9:21.471 | מכל |  |
| 9:21.472 | הנבראים |  |
| 9:21.473 | והנאצלים | , |
| 9:21.474 | וגם |  |
| 9:21.475 | אין |  |
| 9:21.476 | ליחד |  |
| 9:21.477 | ולקשר |  |
| 9:21.478 | עם |  |
| 9:21.479 | אחדותו |  |
| 9:21.480 | שום |  |
| 9:21.481 | אחד |  |
| 9:21.482 | מהם | , |
| 9:21.483 | כמאמר |  |
| 9:21.484 | רשב"י |  |
| 9:21.485 | כל |  |
| 9:21.486 | המשתף |  |
| 9:21.487 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

