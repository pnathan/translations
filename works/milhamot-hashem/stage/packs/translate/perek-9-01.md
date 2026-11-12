# Translation pack: milhamot-hashem perek-9-01 (Milchamot ha-Shem 9:1–9:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-01`.

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
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרמב"ן / רמב"ן** → "Ramban" (authority, transliterated; note first in a
  unit).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **הגאון** alone → "the Gaon".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- **השונים** → "the Dualists".
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
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
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- עבודה **אחרת** → "other worship" (note first in a unit).
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- אורות **דקים** → "subtle lights".
- גוף **דק** → "a subtle body".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
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
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **סכלות** והוללות → "folly and madness".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **חסר / חטר**: "(חסר)" → "(missing)", a lacuna sign; "(חטר)" is a variant of
  it, recorded.

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
- מי **אתי** (ketiv): translated, the qere noted.
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


## Milchamot ha-Shem 9:1

Hebrew (Wikisource): קך) וכתב רבינו סעדיה גאון בהקדמת ספרו האמונות והדעות. וז"ל והאמונה על ב' דרכים. אמת ושקר, והאמונה האמתית היא שידע הדבר כאשר הוא וכו' והאמונה השקרית הוא שידע הדבר בהפך מה הוא וכו'! והחכם המשובח מי ששם אמיתת הדברים שרש. ומנהיג דעתו עליהם ועם חכמתו בטוח אל המובטח, ונזהר ממה שצריך להזהר ממנו (פירש הרב שביל האמונה כלומר עם שהוא חכם. בחכמתו הוא ירא בכל מה שהתורה האלהית הגוזר עליו כהפך עיונו, והוא בוטח על הדבר הראוי לבטוח וירא ונזהר מהדבר אשר ראוי להזהר ממנו יען שנפשו יודעת מאד כי יש מביאות גם לדברים שלא נפל דעתו עליהם עד אשר ייטיב עיונו עליהם עד שיסורו כל הספיקות מלבו. ויתברר לו אמתתו) והכסיל המגונה מי ששם דעתו לשורש, וחושב כי אמתות הדברים הולכות אחר דעתו, ועם סכלותו בוטח במה שצריך להזהר ממנו, ונזהר מהבטוח בו, וכאשר אמר החכם (משלי ד') חכם ירא וסר מרע, וכסיל מתעבר ובוטח ע"כ. ואחרי הדברים והאמת האלה תבין ותשכיל כי קבלת המסית מחבר הזהר והנמשכים אחריו לתומם היא בעצמה דעת הפילסופים הקדמונים הרחוקים מן התורה והמצוה אשר צונו ה' אלהינו על ידי משה בחירו, וזקינים ונביאים שקיבלו איש מפי איש עד חיבור המשנה והתלמוד ועד סוף הגאונים הרס"ג ז"ל אשר נתן אל לבו להשיב על כתות הפילסופים התועים ההולכים אחרי ההבל, לבל יתערבו דעותיהם באמונת תורתינו הטהורה וז"ל הכת האחת היא האומרת שהבורא יש לו גשמים רוחניים דקים ונקודות דקות מאוד כאבק דק וכו' שם במאמר א' חידוש. והכת האחרת היא האומרת שבורא הגשמים בראם מעצמו (רוצה לומר שהפריש חלק מעצמותו ומאותו החלק ברא את העולם) מצאתי אלה האנשים לא נתכן להם לכחש בעושה, ועם כל זה לא קבל שכלם היות דבר מלא דבר. וכיון שאין דבר כי אם הבורא, האמינו כי ברא הדברים מן עצמו. ואלה יותר כסילים מן הראשונים. וראיתי לגלות אולתם בשלש עשרה פנים וכו' ע"ש בדברי הגאון ועומק תשובותיו בזה, ותבין להשיב על הטעאות מחבר הזהר והמקובלים החדשים שנטו אחריו. כי שלי שום ספק שהיא לקוחה מאלה השתי דעות הזרות. ועיין בפי' הרמב"ן לשיר השירים בפסוק אפררון עשה לו המלך שלמה מעצי הלבנון שגם הוא טעה בזה ונטה לרעת הזרה הזאת. וחשב ליחסה לרבי אלעזר הגדול שדעתו כדעת אפלטון האומר כי שוא הוא שהבורא ימציא דבר מלא דבר. מעצמותו ית' והוא דרך משל כחומר ליוצר וכברזל לנפח, יצייר ממנו מה שירצה מהכלים. כן הבורא יצייר מזה הנאצל שמים וארץ וזולתם אין זה חסרון בחוקו ית' כשלא יברא דבר אלא מדבר ודבריו אלה היפך המקובל בכל האומה הישראלית, שהקב"ה המציא כל עם ישראל שהיא מקובלת בידם. ויאמינו סברא פחותה כזו שהקב"ה המציא כל המציאות מעצמותו! וכבר העיר הרמב"ם ז"ל במורה נבוכים פרק כ"ו משני על דברי רבי עליאזר והניח בקושיא. ולא דן אותו לחולק על הקבלה האמתית שהיא פשוטה בכל ישראל. ולעניות דעתי נראה כוונת ר' אליעזר באמרו מאוד לבושו ר"ל מרוב כוחו ית' ועצמותו וגאותו המכונה בכתוב בשם לבוש שנ' לבשי עזך ציון, עורי לבשי עז. לבש ה' עז! ונגד זה נקרא חסרון הכח והיכולת בשם לבוש שנאמר ה' מלך גאות לבש (תהילים צ״ג:א׳). וכן העוז שהוא החוזק והיכולת נקרא חסרון הכח והיכולת בשם לבוש בשת, ילבשו, שוטוני כלמה, שונאיך ילבשו בשת ילבשו בשת וכלמה! ורצה ר' אליעזר ע"ה לומר שברוב כחו ית' ועצמתו ברא השמים יש מאין המוחלט, והשמים נקראו כסא, שנאמר השמים כסאי, ומהגרמים השמימים הנקראים שלג, ברא את הארץ, וזהו כי לשלג יאמר הוי ארץ! והרצון יקרא אור, כמו שנאמר ואור פני לא יפילון, ר"ל שלא יבטלו דברי ויעברו על מה שרציתי כן נ"ל. וזו ראיה גמורה (ר"ל סברת הרמב"ם שמיחס לר"א) שהקבלה החדשה לקוחה מדברי הפילסופים הקדמונים, וגם מאמונת השונים המאמינים בשני פועלים פועל הטוב ופועל הרע, סטרא דקדושה, וסטרא אחרא מסאבא, עיין עליהם ועל תשובת הגאון הרס"ג ז"ל במאמר חדוש, ובפרק ז' מאמר האחדות מצאתי בדברי הגאון ראיה וחזוק למה שכתבתי לעיל וז"ל והכת האומרת כי קצת האל זב גוף ורוח, יתחייבו עם שני התשובות האלה כל מה שהתחייבו בו האומרים כי הדברים הנבראים הם מהבורא והכת האומרת שגופו ורוחו מחודשים ויש בו רוח אחרת מהבורא תחייב כי גוף מחודש שב להיות אלוה בשיתוף דבר אלהי בו ומקישים אותו להשראת הכבוד על הר סיני ובעץ הסנה ובאהל מועד, וזה מחייב אותם להאמין כי האלה והסנה וההר אלהים ג"כ ויוספו רעה וכו' ע"כ. ומדברי הגאון תבין אמתת דברינו שכתבתי לעיל כי נכונים הדברים על פי תורתינו הקדושה וקבלת רז"ל וכמו שכתבו הפוסקים דכל דאתאמר בדברי גאון אליו שומעין שדבריהם דברי קבלה.

| id | word | punctuation after |
|---|---|---|
| 9:1.1 | קך | ) |
| 9:1.2 | וכתב |  |
| 9:1.3 | רבינו |  |
| 9:1.4 | סעדיה |  |
| 9:1.5 | גאון |  |
| 9:1.6 | בהקדמת |  |
| 9:1.7 | ספרו |  |
| 9:1.8 | האמונות |  |
| 9:1.9 | והדעות | . |
| 9:1.10 | וז"ל |  |
| 9:1.11 | והאמונה |  |
| 9:1.12 | על |  |
| 9:1.13 | ב | ' |
| 9:1.14 | דרכים | . |
| 9:1.15 | אמת |  |
| 9:1.16 | ושקר | , |
| 9:1.17 | והאמונה |  |
| 9:1.18 | האמתית |  |
| 9:1.19 | היא |  |
| 9:1.20 | שידע |  |
| 9:1.21 | הדבר |  |
| 9:1.22 | כאשר |  |
| 9:1.23 | הוא |  |
| 9:1.24 | וכו | ' |
| 9:1.25 | והאמונה |  |
| 9:1.26 | השקרית |  |
| 9:1.27 | הוא |  |
| 9:1.28 | שידע |  |
| 9:1.29 | הדבר |  |
| 9:1.30 | בהפך |  |
| 9:1.31 | מה |  |
| 9:1.32 | הוא |  |
| 9:1.33 | וכו | '! |
| 9:1.34 | והחכם |  |
| 9:1.35 | המשובח |  |
| 9:1.36 | מי |  |
| 9:1.37 | ששם |  |
| 9:1.38 | אמיתת |  |
| 9:1.39 | הדברים |  |
| 9:1.40 | שרש | . |
| 9:1.41 | ומנהיג |  |
| 9:1.42 | דעתו |  |
| 9:1.43 | עליהם |  |
| 9:1.44 | ועם |  |
| 9:1.45 | חכמתו |  |
| 9:1.46 | בטוח |  |
| 9:1.47 | אל |  |
| 9:1.48 | המובטח | , |
| 9:1.49 | ונזהר |  |
| 9:1.50 | ממה |  |
| 9:1.51 | שצריך |  |
| 9:1.52 | להזהר |  |
| 9:1.53 | ממנו |  |
| 9:1.54 | פירש |  |
| 9:1.55 | הרב |  |
| 9:1.56 | שביל |  |
| 9:1.57 | האמונה |  |
| 9:1.58 | כלומר |  |
| 9:1.59 | עם |  |
| 9:1.60 | שהוא |  |
| 9:1.61 | חכם | . |
| 9:1.62 | בחכמתו |  |
| 9:1.63 | הוא |  |
| 9:1.64 | ירא |  |
| 9:1.65 | בכל |  |
| 9:1.66 | מה |  |
| 9:1.67 | שהתורה |  |
| 9:1.68 | האלהית |  |
| 9:1.69 | הגוזר |  |
| 9:1.70 | עליו |  |
| 9:1.71 | כהפך |  |
| 9:1.72 | עיונו | , |
| 9:1.73 | והוא |  |
| 9:1.74 | בוטח |  |
| 9:1.75 | על |  |
| 9:1.76 | הדבר |  |
| 9:1.77 | הראוי |  |
| 9:1.78 | לבטוח |  |
| 9:1.79 | וירא |  |
| 9:1.80 | ונזהר |  |
| 9:1.81 | מהדבר |  |
| 9:1.82 | אשר |  |
| 9:1.83 | ראוי |  |
| 9:1.84 | להזהר |  |
| 9:1.85 | ממנו |  |
| 9:1.86 | יען |  |
| 9:1.87 | שנפשו |  |
| 9:1.88 | יודעת |  |
| 9:1.89 | מאד |  |
| 9:1.90 | כי |  |
| 9:1.91 | יש |  |
| 9:1.92 | מביאות |  |
| 9:1.93 | גם |  |
| 9:1.94 | לדברים |  |
| 9:1.95 | שלא |  |
| 9:1.96 | נפל |  |
| 9:1.97 | דעתו |  |
| 9:1.98 | עליהם |  |
| 9:1.99 | עד |  |
| 9:1.100 | אשר |  |
| 9:1.101 | ייטיב |  |
| 9:1.102 | עיונו |  |
| 9:1.103 | עליהם |  |
| 9:1.104 | עד |  |
| 9:1.105 | שיסורו |  |
| 9:1.106 | כל |  |
| 9:1.107 | הספיקות |  |
| 9:1.108 | מלבו | . |
| 9:1.109 | ויתברר |  |
| 9:1.110 | לו |  |
| 9:1.111 | אמתתו | ) |
| 9:1.112 | והכסיל |  |
| 9:1.113 | המגונה |  |
| 9:1.114 | מי |  |
| 9:1.115 | ששם |  |
| 9:1.116 | דעתו |  |
| 9:1.117 | לשורש | , |
| 9:1.118 | וחושב |  |
| 9:1.119 | כי |  |
| 9:1.120 | אמתות |  |
| 9:1.121 | הדברים |  |
| 9:1.122 | הולכות |  |
| 9:1.123 | אחר |  |
| 9:1.124 | דעתו | , |
| 9:1.125 | ועם |  |
| 9:1.126 | סכלותו |  |
| 9:1.127 | בוטח |  |
| 9:1.128 | במה |  |
| 9:1.129 | שצריך |  |
| 9:1.130 | להזהר |  |
| 9:1.131 | ממנו | , |
| 9:1.132 | ונזהר |  |
| 9:1.133 | מהבטוח |  |
| 9:1.134 | בו | , |
| 9:1.135 | וכאשר |  |
| 9:1.136 | אמר |  |
| 9:1.137 | החכם |  |
| 9:1.138 | משלי |  |
| 9:1.139 | ד | ') |
| 9:1.140 | חכם |  |
| 9:1.141 | ירא |  |
| 9:1.142 | וסר |  |
| 9:1.143 | מרע | , |
| 9:1.144 | וכסיל |  |
| 9:1.145 | מתעבר |  |
| 9:1.146 | ובוטח |  |
| 9:1.147 | ע"כ | . |
| 9:1.148 | ואחרי |  |
| 9:1.149 | הדברים |  |
| 9:1.150 | והאמת |  |
| 9:1.151 | האלה |  |
| 9:1.152 | תבין |  |
| 9:1.153 | ותשכיל |  |
| 9:1.154 | כי |  |
| 9:1.155 | קבלת |  |
| 9:1.156 | המסית |  |
| 9:1.157 | מחבר |  |
| 9:1.158 | הזהר |  |
| 9:1.159 | והנמשכים |  |
| 9:1.160 | אחריו |  |
| 9:1.161 | לתומם |  |
| 9:1.162 | היא |  |
| 9:1.163 | בעצמה |  |
| 9:1.164 | דעת |  |
| 9:1.165 | הפילסופים |  |
| 9:1.166 | הקדמונים |  |
| 9:1.167 | הרחוקים |  |
| 9:1.168 | מן |  |
| 9:1.169 | התורה |  |
| 9:1.170 | והמצוה |  |
| 9:1.171 | אשר |  |
| 9:1.172 | צונו |  |
| 9:1.173 | ה | ' |
| 9:1.174 | אלהינו |  |
| 9:1.175 | על |  |
| 9:1.176 | ידי |  |
| 9:1.177 | משה |  |
| 9:1.178 | בחירו | , |
| 9:1.179 | וזקינים |  |
| 9:1.180 | ונביאים |  |
| 9:1.181 | שקיבלו |  |
| 9:1.182 | איש |  |
| 9:1.183 | מפי |  |
| 9:1.184 | איש |  |
| 9:1.185 | עד |  |
| 9:1.186 | חיבור |  |
| 9:1.187 | המשנה |  |
| 9:1.188 | והתלמוד |  |
| 9:1.189 | ועד |  |
| 9:1.190 | סוף |  |
| 9:1.191 | הגאונים |  |
| 9:1.192 | הרס"ג |  |
| 9:1.193 | ז"ל |  |
| 9:1.194 | אשר |  |
| 9:1.195 | נתן |  |
| 9:1.196 | אל |  |
| 9:1.197 | לבו |  |
| 9:1.198 | להשיב |  |
| 9:1.199 | על |  |
| 9:1.200 | כתות |  |
| 9:1.201 | הפילסופים |  |
| 9:1.202 | התועים |  |
| 9:1.203 | ההולכים |  |
| 9:1.204 | אחרי |  |
| 9:1.205 | ההבל | , |
| 9:1.206 | לבל |  |
| 9:1.207 | יתערבו |  |
| 9:1.208 | דעותיהם |  |
| 9:1.209 | באמונת |  |
| 9:1.210 | תורתינו |  |
| 9:1.211 | הטהורה |  |
| 9:1.212 | וז"ל |  |
| 9:1.213 | הכת |  |
| 9:1.214 | האחת |  |
| 9:1.215 | היא |  |
| 9:1.216 | האומרת |  |
| 9:1.217 | שהבורא |  |
| 9:1.218 | יש |  |
| 9:1.219 | לו |  |
| 9:1.220 | גשמים |  |
| 9:1.221 | רוחניים |  |
| 9:1.222 | דקים |  |
| 9:1.223 | ונקודות |  |
| 9:1.224 | דקות |  |
| 9:1.225 | מאוד |  |
| 9:1.226 | כאבק |  |
| 9:1.227 | דק |  |
| 9:1.228 | וכו | ' |
| 9:1.229 | שם |  |
| 9:1.230 | במאמר |  |
| 9:1.231 | א | ' |
| 9:1.232 | חידוש | . |
| 9:1.233 | והכת |  |
| 9:1.234 | האחרת |  |
| 9:1.235 | היא |  |
| 9:1.236 | האומרת |  |
| 9:1.237 | שבורא |  |
| 9:1.238 | הגשמים |  |
| 9:1.239 | בראם |  |
| 9:1.240 | מעצמו |  |
| 9:1.241 | רוצה |  |
| 9:1.242 | לומר |  |
| 9:1.243 | שהפריש |  |
| 9:1.244 | חלק |  |
| 9:1.245 | מעצמותו |  |
| 9:1.246 | ומאותו |  |
| 9:1.247 | החלק |  |
| 9:1.248 | ברא |  |
| 9:1.249 | את |  |
| 9:1.250 | העולם | ) |
| 9:1.251 | מצאתי |  |
| 9:1.252 | אלה |  |
| 9:1.253 | האנשים |  |
| 9:1.254 | לא |  |
| 9:1.255 | נתכן |  |
| 9:1.256 | להם |  |
| 9:1.257 | לכחש |  |
| 9:1.258 | בעושה | , |
| 9:1.259 | ועם |  |
| 9:1.260 | כל |  |
| 9:1.261 | זה |  |
| 9:1.262 | לא |  |
| 9:1.263 | קבל |  |
| 9:1.264 | שכלם |  |
| 9:1.265 | היות |  |
| 9:1.266 | דבר |  |
| 9:1.267 | מלא |  |
| 9:1.268 | דבר | . |
| 9:1.269 | וכיון |  |
| 9:1.270 | שאין |  |
| 9:1.271 | דבר |  |
| 9:1.272 | כי |  |
| 9:1.273 | אם |  |
| 9:1.274 | הבורא | , |
| 9:1.275 | האמינו |  |
| 9:1.276 | כי |  |
| 9:1.277 | ברא |  |
| 9:1.278 | הדברים |  |
| 9:1.279 | מן |  |
| 9:1.280 | עצמו | . |
| 9:1.281 | ואלה |  |
| 9:1.282 | יותר |  |
| 9:1.283 | כסילים |  |
| 9:1.284 | מן |  |
| 9:1.285 | הראשונים | . |
| 9:1.286 | וראיתי |  |
| 9:1.287 | לגלות |  |
| 9:1.288 | אולתם |  |
| 9:1.289 | בשלש |  |
| 9:1.290 | עשרה |  |
| 9:1.291 | פנים |  |
| 9:1.292 | וכו | ' |
| 9:1.293 | ע"ש |  |
| 9:1.294 | בדברי |  |
| 9:1.295 | הגאון |  |
| 9:1.296 | ועומק |  |
| 9:1.297 | תשובותיו |  |
| 9:1.298 | בזה | , |
| 9:1.299 | ותבין |  |
| 9:1.300 | להשיב |  |
| 9:1.301 | על |  |
| 9:1.302 | הטעאות |  |
| 9:1.303 | מחבר |  |
| 9:1.304 | הזהר |  |
| 9:1.305 | והמקובלים |  |
| 9:1.306 | החדשים |  |
| 9:1.307 | שנטו |  |
| 9:1.308 | אחריו | . |
| 9:1.309 | כי |  |
| 9:1.310 | שלי |  |
| 9:1.311 | שום |  |
| 9:1.312 | ספק |  |
| 9:1.313 | שהיא |  |
| 9:1.314 | לקוחה |  |
| 9:1.315 | מאלה |  |
| 9:1.316 | השתי |  |
| 9:1.317 | דעות |  |
| 9:1.318 | הזרות | . |
| 9:1.319 | ועיין |  |
| 9:1.320 | בפי | ' |
| 9:1.321 | הרמב"ן |  |
| 9:1.322 | לשיר |  |
| 9:1.323 | השירים |  |
| 9:1.324 | בפסוק |  |
| 9:1.325 | אפררון |  |
| 9:1.326 | עשה |  |
| 9:1.327 | לו |  |
| 9:1.328 | המלך |  |
| 9:1.329 | שלמה |  |
| 9:1.330 | מעצי |  |
| 9:1.331 | הלבנון |  |
| 9:1.332 | שגם |  |
| 9:1.333 | הוא |  |
| 9:1.334 | טעה |  |
| 9:1.335 | בזה |  |
| 9:1.336 | ונטה |  |
| 9:1.337 | לרעת |  |
| 9:1.338 | הזרה |  |
| 9:1.339 | הזאת | . |
| 9:1.340 | וחשב |  |
| 9:1.341 | ליחסה |  |
| 9:1.342 | לרבי |  |
| 9:1.343 | אלעזר |  |
| 9:1.344 | הגדול |  |
| 9:1.345 | שדעתו |  |
| 9:1.346 | כדעת |  |
| 9:1.347 | אפלטון |  |
| 9:1.348 | האומר |  |
| 9:1.349 | כי |  |
| 9:1.350 | שוא |  |
| 9:1.351 | הוא |  |
| 9:1.352 | שהבורא |  |
| 9:1.353 | ימציא |  |
| 9:1.354 | דבר |  |
| 9:1.355 | מלא |  |
| 9:1.356 | דבר | . |
| 9:1.357 | מעצמותו |  |
| 9:1.358 | ית | ' |
| 9:1.359 | והוא |  |
| 9:1.360 | דרך |  |
| 9:1.361 | משל |  |
| 9:1.362 | כחומר |  |
| 9:1.363 | ליוצר |  |
| 9:1.364 | וכברזל |  |
| 9:1.365 | לנפח | , |
| 9:1.366 | יצייר |  |
| 9:1.367 | ממנו |  |
| 9:1.368 | מה |  |
| 9:1.369 | שירצה |  |
| 9:1.370 | מהכלים | . |
| 9:1.371 | כן |  |
| 9:1.372 | הבורא |  |
| 9:1.373 | יצייר |  |
| 9:1.374 | מזה |  |
| 9:1.375 | הנאצל |  |
| 9:1.376 | שמים |  |
| 9:1.377 | וארץ |  |
| 9:1.378 | וזולתם |  |
| 9:1.379 | אין |  |
| 9:1.380 | זה |  |
| 9:1.381 | חסרון |  |
| 9:1.382 | בחוקו |  |
| 9:1.383 | ית | ' |
| 9:1.384 | כשלא |  |
| 9:1.385 | יברא |  |
| 9:1.386 | דבר |  |
| 9:1.387 | אלא |  |
| 9:1.388 | מדבר |  |
| 9:1.389 | ודבריו |  |
| 9:1.390 | אלה |  |
| 9:1.391 | היפך |  |
| 9:1.392 | המקובל |  |
| 9:1.393 | בכל |  |
| 9:1.394 | האומה |  |
| 9:1.395 | הישראלית | , |
| 9:1.396 | שהקב"ה |  |
| 9:1.397 | המציא |  |
| 9:1.398 | כל |  |
| 9:1.399 | עם |  |
| 9:1.400 | ישראל |  |
| 9:1.401 | שהיא |  |
| 9:1.402 | מקובלת |  |
| 9:1.403 | בידם | . |
| 9:1.404 | ויאמינו |  |
| 9:1.405 | סברא |  |
| 9:1.406 | פחותה |  |
| 9:1.407 | כזו |  |
| 9:1.408 | שהקב"ה |  |
| 9:1.409 | המציא |  |
| 9:1.410 | כל |  |
| 9:1.411 | המציאות |  |
| 9:1.412 | מעצמותו | ! |
| 9:1.413 | וכבר |  |
| 9:1.414 | העיר |  |
| 9:1.415 | הרמב"ם |  |
| 9:1.416 | ז"ל |  |
| 9:1.417 | במורה |  |
| 9:1.418 | נבוכים |  |
| 9:1.419 | פרק |  |
| 9:1.420 | כ"ו |  |
| 9:1.421 | משני |  |
| 9:1.422 | על |  |
| 9:1.423 | דברי |  |
| 9:1.424 | רבי |  |
| 9:1.425 | עליאזר |  |
| 9:1.426 | והניח |  |
| 9:1.427 | בקושיא | . |
| 9:1.428 | ולא |  |
| 9:1.429 | דן |  |
| 9:1.430 | אותו |  |
| 9:1.431 | לחולק |  |
| 9:1.432 | על |  |
| 9:1.433 | הקבלה |  |
| 9:1.434 | האמתית |  |
| 9:1.435 | שהיא |  |
| 9:1.436 | פשוטה |  |
| 9:1.437 | בכל |  |
| 9:1.438 | ישראל | . |
| 9:1.439 | ולעניות |  |
| 9:1.440 | דעתי |  |
| 9:1.441 | נראה |  |
| 9:1.442 | כוונת |  |
| 9:1.443 | ר | ' |
| 9:1.444 | אליעזר |  |
| 9:1.445 | באמרו |  |
| 9:1.446 | מאוד |  |
| 9:1.447 | לבושו |  |
| 9:1.448 | ר"ל |  |
| 9:1.449 | מרוב |  |
| 9:1.450 | כוחו |  |
| 9:1.451 | ית | ' |
| 9:1.452 | ועצמותו |  |
| 9:1.453 | וגאותו |  |
| 9:1.454 | המכונה |  |
| 9:1.455 | בכתוב |  |
| 9:1.456 | בשם |  |
| 9:1.457 | לבוש |  |
| 9:1.458 | שנ | ' |
| 9:1.459 | לבשי |  |
| 9:1.460 | עזך |  |
| 9:1.461 | ציון | , |
| 9:1.462 | עורי |  |
| 9:1.463 | לבשי |  |
| 9:1.464 | עז | . |
| 9:1.465 | לבש |  |
| 9:1.466 | ה | ' |
| 9:1.467 | עז | ! |
| 9:1.468 | ונגד |  |
| 9:1.469 | זה |  |
| 9:1.470 | נקרא |  |
| 9:1.471 | חסרון |  |
| 9:1.472 | הכח |  |
| 9:1.473 | והיכולת |  |
| 9:1.474 | בשם |  |
| 9:1.475 | לבוש |  |
| 9:1.476 | שנאמר |  |
| 9:1.477 | ה | ' |
| 9:1.478 | מלך |  |
| 9:1.479 | גאות |  |
| 9:1.480 | לבש |  |
| 9:1.481 | תהילים |  |
| 9:1.482 | צ״ג:א׳ | ). |
| 9:1.483 | וכן |  |
| 9:1.484 | העוז |  |
| 9:1.485 | שהוא |  |
| 9:1.486 | החוזק |  |
| 9:1.487 | והיכולת |  |
| 9:1.488 | נקרא |  |
| 9:1.489 | חסרון |  |
| 9:1.490 | הכח |  |
| 9:1.491 | והיכולת |  |
| 9:1.492 | בשם |  |
| 9:1.493 | לבוש |  |
| 9:1.494 | בשת | , |
| 9:1.495 | ילבשו | , |
| 9:1.496 | שוטוני |  |
| 9:1.497 | כלמה | , |
| 9:1.498 | שונאיך |  |
| 9:1.499 | ילבשו |  |
| 9:1.500 | בשת |  |
| 9:1.501 | ילבשו |  |
| 9:1.502 | בשת |  |
| 9:1.503 | וכלמה | ! |
| 9:1.504 | ורצה |  |
| 9:1.505 | ר | ' |
| 9:1.506 | אליעזר |  |
| 9:1.507 | ע"ה |  |
| 9:1.508 | לומר |  |
| 9:1.509 | שברוב |  |
| 9:1.510 | כחו |  |
| 9:1.511 | ית | ' |
| 9:1.512 | ועצמתו |  |
| 9:1.513 | ברא |  |
| 9:1.514 | השמים |  |
| 9:1.515 | יש |  |
| 9:1.516 | מאין |  |
| 9:1.517 | המוחלט | , |
| 9:1.518 | והשמים |  |
| 9:1.519 | נקראו |  |
| 9:1.520 | כסא | , |
| 9:1.521 | שנאמר |  |
| 9:1.522 | השמים |  |
| 9:1.523 | כסאי | , |
| 9:1.524 | ומהגרמים |  |
| 9:1.525 | השמימים |  |
| 9:1.526 | הנקראים |  |
| 9:1.527 | שלג | , |
| 9:1.528 | ברא |  |
| 9:1.529 | את |  |
| 9:1.530 | הארץ | , |
| 9:1.531 | וזהו |  |
| 9:1.532 | כי |  |
| 9:1.533 | לשלג |  |
| 9:1.534 | יאמר |  |
| 9:1.535 | הוי |  |
| 9:1.536 | ארץ | ! |
| 9:1.537 | והרצון |  |
| 9:1.538 | יקרא |  |
| 9:1.539 | אור | , |
| 9:1.540 | כמו |  |
| 9:1.541 | שנאמר |  |
| 9:1.542 | ואור |  |
| 9:1.543 | פני |  |
| 9:1.544 | לא |  |
| 9:1.545 | יפילון | , |
| 9:1.546 | ר"ל |  |
| 9:1.547 | שלא |  |
| 9:1.548 | יבטלו |  |
| 9:1.549 | דברי |  |
| 9:1.550 | ויעברו |  |
| 9:1.551 | על |  |
| 9:1.552 | מה |  |
| 9:1.553 | שרציתי |  |
| 9:1.554 | כן |  |
| 9:1.555 | נ"ל | . |
| 9:1.556 | וזו |  |
| 9:1.557 | ראיה |  |
| 9:1.558 | גמורה |  |
| 9:1.559 | ר"ל |  |
| 9:1.560 | סברת |  |
| 9:1.561 | הרמב"ם |  |
| 9:1.562 | שמיחס |  |
| 9:1.563 | לר"א | ) |
| 9:1.564 | שהקבלה |  |
| 9:1.565 | החדשה |  |
| 9:1.566 | לקוחה |  |
| 9:1.567 | מדברי |  |
| 9:1.568 | הפילסופים |  |
| 9:1.569 | הקדמונים | , |
| 9:1.570 | וגם |  |
| 9:1.571 | מאמונת |  |
| 9:1.572 | השונים |  |
| 9:1.573 | המאמינים |  |
| 9:1.574 | בשני |  |
| 9:1.575 | פועלים |  |
| 9:1.576 | פועל |  |
| 9:1.577 | הטוב |  |
| 9:1.578 | ופועל |  |
| 9:1.579 | הרע | , |
| 9:1.580 | סטרא |  |
| 9:1.581 | דקדושה | , |
| 9:1.582 | וסטרא |  |
| 9:1.583 | אחרא |  |
| 9:1.584 | מסאבא | , |
| 9:1.585 | עיין |  |
| 9:1.586 | עליהם |  |
| 9:1.587 | ועל |  |
| 9:1.588 | תשובת |  |
| 9:1.589 | הגאון |  |
| 9:1.590 | הרס"ג |  |
| 9:1.591 | ז"ל |  |
| 9:1.592 | במאמר |  |
| 9:1.593 | חדוש | , |
| 9:1.594 | ובפרק |  |
| 9:1.595 | ז | ' |
| 9:1.596 | מאמר |  |
| 9:1.597 | האחדות |  |
| 9:1.598 | מצאתי |  |
| 9:1.599 | בדברי |  |
| 9:1.600 | הגאון |  |
| 9:1.601 | ראיה |  |
| 9:1.602 | וחזוק |  |
| 9:1.603 | למה |  |
| 9:1.604 | שכתבתי |  |
| 9:1.605 | לעיל |  |
| 9:1.606 | וז"ל |  |
| 9:1.607 | והכת |  |
| 9:1.608 | האומרת |  |
| 9:1.609 | כי |  |
| 9:1.610 | קצת |  |
| 9:1.611 | האל |  |
| 9:1.612 | זב |  |
| 9:1.613 | גוף |  |
| 9:1.614 | ורוח | , |
| 9:1.615 | יתחייבו |  |
| 9:1.616 | עם |  |
| 9:1.617 | שני |  |
| 9:1.618 | התשובות |  |
| 9:1.619 | האלה |  |
| 9:1.620 | כל |  |
| 9:1.621 | מה |  |
| 9:1.622 | שהתחייבו |  |
| 9:1.623 | בו |  |
| 9:1.624 | האומרים |  |
| 9:1.625 | כי |  |
| 9:1.626 | הדברים |  |
| 9:1.627 | הנבראים |  |
| 9:1.628 | הם |  |
| 9:1.629 | מהבורא |  |
| 9:1.630 | והכת |  |
| 9:1.631 | האומרת |  |
| 9:1.632 | שגופו |  |
| 9:1.633 | ורוחו |  |
| 9:1.634 | מחודשים |  |
| 9:1.635 | ויש |  |
| 9:1.636 | בו |  |
| 9:1.637 | רוח |  |
| 9:1.638 | אחרת |  |
| 9:1.639 | מהבורא |  |
| 9:1.640 | תחייב |  |
| 9:1.641 | כי |  |
| 9:1.642 | גוף |  |
| 9:1.643 | מחודש |  |
| 9:1.644 | שב |  |
| 9:1.645 | להיות |  |
| 9:1.646 | אלוה |  |
| 9:1.647 | בשיתוף |  |
| 9:1.648 | דבר |  |
| 9:1.649 | אלהי |  |
| 9:1.650 | בו |  |
| 9:1.651 | ומקישים |  |
| 9:1.652 | אותו |  |
| 9:1.653 | להשראת |  |
| 9:1.654 | הכבוד |  |
| 9:1.655 | על |  |
| 9:1.656 | הר |  |
| 9:1.657 | סיני |  |
| 9:1.658 | ובעץ |  |
| 9:1.659 | הסנה |  |
| 9:1.660 | ובאהל |  |
| 9:1.661 | מועד | , |
| 9:1.662 | וזה |  |
| 9:1.663 | מחייב |  |
| 9:1.664 | אותם |  |
| 9:1.665 | להאמין |  |
| 9:1.666 | כי |  |
| 9:1.667 | האלה |  |
| 9:1.668 | והסנה |  |
| 9:1.669 | וההר |  |
| 9:1.670 | אלהים |  |
| 9:1.671 | ג"כ |  |
| 9:1.672 | ויוספו |  |
| 9:1.673 | רעה |  |
| 9:1.674 | וכו | ' |
| 9:1.675 | ע"כ | . |
| 9:1.676 | ומדברי |  |
| 9:1.677 | הגאון |  |
| 9:1.678 | תבין |  |
| 9:1.679 | אמתת |  |
| 9:1.680 | דברינו |  |
| 9:1.681 | שכתבתי |  |
| 9:1.682 | לעיל |  |
| 9:1.683 | כי |  |
| 9:1.684 | נכונים |  |
| 9:1.685 | הדברים |  |
| 9:1.686 | על |  |
| 9:1.687 | פי |  |
| 9:1.688 | תורתינו |  |
| 9:1.689 | הקדושה |  |
| 9:1.690 | וקבלת |  |
| 9:1.691 | רז"ל |  |
| 9:1.692 | וכמו |  |
| 9:1.693 | שכתבו |  |
| 9:1.694 | הפוסקים |  |
| 9:1.695 | דכל |  |
| 9:1.696 | דאתאמר |  |
| 9:1.697 | בדברי |  |
| 9:1.698 | גאון |  |
| 9:1.699 | אליו |  |
| 9:1.700 | שומעין |  |
| 9:1.701 | שדבריהם |  |
| 9:1.702 | דברי |  |
| 9:1.703 | קבלה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

