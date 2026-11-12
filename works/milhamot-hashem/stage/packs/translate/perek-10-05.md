# Translation pack: milhamot-hashem perek-10-05 (Milchamot ha-Shem 10:6–10:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-05`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
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
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
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
- אלהים **אחרים** → "other gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **גבורה** as a name of God → "the Power".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 10:6

Hebrew (Wikisource): קמד) סוף דבר תנה לבך קורא נעים, והטה אזנך לתושיה ולתבונה, ושים נא כבוד לה' ותן לו תודה על אשר העיר והאיר ברוח קדשו את חוזינו בכתבי הקדש, וגם אחרי כן החלה רוח ה' לפעם את חכמינו בעלי המשנה ושני התלמודים ומדרשים האמתיים והביטו בעיני שכלם הטהורות מראש ומקדם את העת הרעה, אשר עוד היתה רחוקה מהם כמה מאות שנים והדליקו את האבוקות הגדולות להראות לנו את האמונה הטהורה האחדות השי"ת, בטרם יבוא אפל חשך וצלמות ספר הזהר על ידי המלך העכו"ם העומם את יפעת אמונותינו הטהורה, על ידי אמונתו התפילה והבדויה אשר בשקר הוליד אותה לסלף דברי צדיקים ולדבר עתק עד צדיקו של עולם, כי המשיך הקב"ה מאורו אלהים אחרים ומינה אותם תחתיו, וצוה לעבדם ולהתפלל אליהם. על הראשונים שעשו צלמי זהב וכסף ועץ ואבן אמר הנביא ירמיה ע"ה ובאחת יבערו ויכסלו מוסר הבלים עץ הוא! ועל אנשי דורנו זה ההולכים בעקבות נביא השקר מחבר הזהר אמרה תורה כי יקום בקרבך נביא או חולם חלום וכו' נלכה ונעבדה אלהים אחרים אשר לא ידעתם לא מצד הקבלה שהעתיקו רז"ל ולא מצד המופת כמו שפירש ר' עובדיה ספורנו (דברים י"ג) וז"ל שלא בא מופת אפילו על מציאותם, כי אמנם הנודעים בחוש או במופת נודע מהם שהם פועלים צמיד בענין אחד באופן שנודע שהם פועלים טבעיים לא רצוניים, ובכן לא יועיל להתפלל אליהם ולא לעבדם לא תשמע אל דברי הנביא ההוא לא תפנה אל דבריו לראות אם יש בקצתם, כי אין ספק שכל דבריו קשר וכו' עכ"ל. אשר יאמר שכל שלא בא מופת על מציאותם כצורות ופרצופים שהמציא מחבר הזהר בנבואצו שתלה ברשב"י התנא הכל שקר! שהרי אפי' הנודעים בחוש או במופת פועלים טבעיים הם ולא רצוניים, עליהם אמר שמואל הנביא לישראל ולא תסורו כי אחרי התהו אשר לא יועילו ולא יצילו כי תהו המה! כי כל הפרצופים והצורות והעגולים שהמציא בדמיונו תהו והבל המה, ואין להם מציאות נרגש בחוש וגם אין מופת על מציאותם בה' אלהינו שיש כמה מופתים על מציאותו ית' כמפורש בתורה ובנביאים ובכתובים ובדברי רז"ל! כי אמנם כל דבריו ספורי נפלאות ונמנעות אשר ישמטו את המשנה והתלמוד, וישכינו לעפר כבודם, אף כי בת מלך היא המשנה ועליה עמד משה רבינו בהר ארבעים יום ומ' לילה עד שלמד אותה מפי הגבורה וירד לישראל! ונביא השקר מחבר הזהר בשנאתו הגדולה למשנה ולתלמוד, הרחיה עז בנפשו לבשות את המשנה ולהוריד לארץ את כבודה כדי לשכחה ולעקור אותה ואת חקותיה מישראל בתתו אותה לשמצה! אף כמעט כל קוראיו לא הבינו את מזמותיו, כי חשבו את הכל לתורה אחת ומשפט אחד, ולא הרפו את ידם מלעסוק בזה, וגם מזה לא הניחו ידם! ישראל לא ידע, עמי לא התבונן, שכל מטרתו ומגמתו של הפילסוף מחבר הזהר לתת לנו תורה שבעל פה חדשה תמורת המשנה והתלמוד! מהלל שם ה' החונן לאדם דעת ותבונה, להבין ולהשכיל קשט אמרי אמת, ולנפץ חזון פריצים אשר ינשאו להעמיד חזון ונכשלו והכשילו רבים להאמין בשוא נתעה כי ה' אלהינו איננו אחד ומיוחד, על ידי לשון רמיה ושפת חלקות, עתה יראו כל חכמי לב וישכילו מתורת ה' תמימה שבכתב ושבעל פה המשנה והתלמוד עם מדרשי רז"ל האמתיים ועם פירושו למשניות והקדמותיו עם שאר חכמינו זכי התבונה. ונפשו תשבע להם חלות טהורות משוחות בשטן האמונה הטהורה כנתינתה מסיני. ועל ידה יזכה לחזות בנועם ה' ולבקר בהיכלו, עד כה עזרנו השי"ת לגלות פני הבלות מחשבות הפילסוף נביא השקר המתעה מחבר הזהר, אשר זמם לדבר סרה על ה' ועל תורתו שבעל פה פיו ידבר תהפוכות לבחור אלהים חדשים שאר בדא אותם בדמיונו, ונתלה באילן גדול רשב"י התנא הקדוש, למען תפוש את בית ישראל בלבם! ועל האמונה הטהורה המשנה והתלמוד ישליך שקוצים לשומה שפחה חרופה ולהתעמר בה! וירומם את תורתו החדשה בלשון רמיה ושפת חלקות לכל פתי ותועה רוח! ברוך החונן לאדם דעת להשכיל להטיב, לכו ונלכה באור ה' היא תורתו הקדושה והתמימה והטהורה, אשר תבונינו ותלמדינו דעת את ה' אלהים אמת, ועל ידי זה נזכה לחזות בנועם ה' בעולם הזה ובעולם הבא אכי"ר.

| id | word | punctuation after |
|---|---|---|
| 10:6.1 | קמד | ) |
| 10:6.2 | סוף |  |
| 10:6.3 | דבר |  |
| 10:6.4 | תנה |  |
| 10:6.5 | לבך |  |
| 10:6.6 | קורא |  |
| 10:6.7 | נעים | , |
| 10:6.8 | והטה |  |
| 10:6.9 | אזנך |  |
| 10:6.10 | לתושיה |  |
| 10:6.11 | ולתבונה | , |
| 10:6.12 | ושים |  |
| 10:6.13 | נא |  |
| 10:6.14 | כבוד |  |
| 10:6.15 | לה | ' |
| 10:6.16 | ותן |  |
| 10:6.17 | לו |  |
| 10:6.18 | תודה |  |
| 10:6.19 | על |  |
| 10:6.20 | אשר |  |
| 10:6.21 | העיר |  |
| 10:6.22 | והאיר |  |
| 10:6.23 | ברוח |  |
| 10:6.24 | קדשו |  |
| 10:6.25 | את |  |
| 10:6.26 | חוזינו |  |
| 10:6.27 | בכתבי |  |
| 10:6.28 | הקדש | , |
| 10:6.29 | וגם |  |
| 10:6.30 | אחרי |  |
| 10:6.31 | כן |  |
| 10:6.32 | החלה |  |
| 10:6.33 | רוח |  |
| 10:6.34 | ה | ' |
| 10:6.35 | לפעם |  |
| 10:6.36 | את |  |
| 10:6.37 | חכמינו |  |
| 10:6.38 | בעלי |  |
| 10:6.39 | המשנה |  |
| 10:6.40 | ושני |  |
| 10:6.41 | התלמודים |  |
| 10:6.42 | ומדרשים |  |
| 10:6.43 | האמתיים |  |
| 10:6.44 | והביטו |  |
| 10:6.45 | בעיני |  |
| 10:6.46 | שכלם |  |
| 10:6.47 | הטהורות |  |
| 10:6.48 | מראש |  |
| 10:6.49 | ומקדם |  |
| 10:6.50 | את |  |
| 10:6.51 | העת |  |
| 10:6.52 | הרעה | , |
| 10:6.53 | אשר |  |
| 10:6.54 | עוד |  |
| 10:6.55 | היתה |  |
| 10:6.56 | רחוקה |  |
| 10:6.57 | מהם |  |
| 10:6.58 | כמה |  |
| 10:6.59 | מאות |  |
| 10:6.60 | שנים |  |
| 10:6.61 | והדליקו |  |
| 10:6.62 | את |  |
| 10:6.63 | האבוקות |  |
| 10:6.64 | הגדולות |  |
| 10:6.65 | להראות |  |
| 10:6.66 | לנו |  |
| 10:6.67 | את |  |
| 10:6.68 | האמונה |  |
| 10:6.69 | הטהורה |  |
| 10:6.70 | האחדות |  |
| 10:6.71 | השי"ת | , |
| 10:6.72 | בטרם |  |
| 10:6.73 | יבוא |  |
| 10:6.74 | אפל |  |
| 10:6.75 | חשך |  |
| 10:6.76 | וצלמות |  |
| 10:6.77 | ספר |  |
| 10:6.78 | הזהר |  |
| 10:6.79 | על |  |
| 10:6.80 | ידי |  |
| 10:6.81 | המלך |  |
| 10:6.82 | העכו"ם |  |
| 10:6.83 | העומם |  |
| 10:6.84 | את |  |
| 10:6.85 | יפעת |  |
| 10:6.86 | אמונותינו |  |
| 10:6.87 | הטהורה | , |
| 10:6.88 | על |  |
| 10:6.89 | ידי |  |
| 10:6.90 | אמונתו |  |
| 10:6.91 | התפילה |  |
| 10:6.92 | והבדויה |  |
| 10:6.93 | אשר |  |
| 10:6.94 | בשקר |  |
| 10:6.95 | הוליד |  |
| 10:6.96 | אותה |  |
| 10:6.97 | לסלף |  |
| 10:6.98 | דברי |  |
| 10:6.99 | צדיקים |  |
| 10:6.100 | ולדבר |  |
| 10:6.101 | עתק |  |
| 10:6.102 | עד |  |
| 10:6.103 | צדיקו |  |
| 10:6.104 | של |  |
| 10:6.105 | עולם | , |
| 10:6.106 | כי |  |
| 10:6.107 | המשיך |  |
| 10:6.108 | הקב"ה |  |
| 10:6.109 | מאורו |  |
| 10:6.110 | אלהים |  |
| 10:6.111 | אחרים |  |
| 10:6.112 | ומינה |  |
| 10:6.113 | אותם |  |
| 10:6.114 | תחתיו | , |
| 10:6.115 | וצוה |  |
| 10:6.116 | לעבדם |  |
| 10:6.117 | ולהתפלל |  |
| 10:6.118 | אליהם | . |
| 10:6.119 | על |  |
| 10:6.120 | הראשונים |  |
| 10:6.121 | שעשו |  |
| 10:6.122 | צלמי |  |
| 10:6.123 | זהב |  |
| 10:6.124 | וכסף |  |
| 10:6.125 | ועץ |  |
| 10:6.126 | ואבן |  |
| 10:6.127 | אמר |  |
| 10:6.128 | הנביא |  |
| 10:6.129 | ירמיה |  |
| 10:6.130 | ע"ה |  |
| 10:6.131 | ובאחת |  |
| 10:6.132 | יבערו |  |
| 10:6.133 | ויכסלו |  |
| 10:6.134 | מוסר |  |
| 10:6.135 | הבלים |  |
| 10:6.136 | עץ |  |
| 10:6.137 | הוא | ! |
| 10:6.138 | ועל |  |
| 10:6.139 | אנשי |  |
| 10:6.140 | דורנו |  |
| 10:6.141 | זה |  |
| 10:6.142 | ההולכים |  |
| 10:6.143 | בעקבות |  |
| 10:6.144 | נביא |  |
| 10:6.145 | השקר |  |
| 10:6.146 | מחבר |  |
| 10:6.147 | הזהר |  |
| 10:6.148 | אמרה |  |
| 10:6.149 | תורה |  |
| 10:6.150 | כי |  |
| 10:6.151 | יקום |  |
| 10:6.152 | בקרבך |  |
| 10:6.153 | נביא |  |
| 10:6.154 | או |  |
| 10:6.155 | חולם |  |
| 10:6.156 | חלום |  |
| 10:6.157 | וכו | ' |
| 10:6.158 | נלכה |  |
| 10:6.159 | ונעבדה |  |
| 10:6.160 | אלהים |  |
| 10:6.161 | אחרים |  |
| 10:6.162 | אשר |  |
| 10:6.163 | לא |  |
| 10:6.164 | ידעתם |  |
| 10:6.165 | לא |  |
| 10:6.166 | מצד |  |
| 10:6.167 | הקבלה |  |
| 10:6.168 | שהעתיקו |  |
| 10:6.169 | רז"ל |  |
| 10:6.170 | ולא |  |
| 10:6.171 | מצד |  |
| 10:6.172 | המופת |  |
| 10:6.173 | כמו |  |
| 10:6.174 | שפירש |  |
| 10:6.175 | ר | ' |
| 10:6.176 | עובדיה |  |
| 10:6.177 | ספורנו |  |
| 10:6.178 | דברים |  |
| 10:6.179 | י"ג | ) |
| 10:6.180 | וז"ל |  |
| 10:6.181 | שלא |  |
| 10:6.182 | בא |  |
| 10:6.183 | מופת |  |
| 10:6.184 | אפילו |  |
| 10:6.185 | על |  |
| 10:6.186 | מציאותם | , |
| 10:6.187 | כי |  |
| 10:6.188 | אמנם |  |
| 10:6.189 | הנודעים |  |
| 10:6.190 | בחוש |  |
| 10:6.191 | או |  |
| 10:6.192 | במופת |  |
| 10:6.193 | נודע |  |
| 10:6.194 | מהם |  |
| 10:6.195 | שהם |  |
| 10:6.196 | פועלים |  |
| 10:6.197 | צמיד |  |
| 10:6.198 | בענין |  |
| 10:6.199 | אחד |  |
| 10:6.200 | באופן |  |
| 10:6.201 | שנודע |  |
| 10:6.202 | שהם |  |
| 10:6.203 | פועלים |  |
| 10:6.204 | טבעיים |  |
| 10:6.205 | לא |  |
| 10:6.206 | רצוניים | , |
| 10:6.207 | ובכן |  |
| 10:6.208 | לא |  |
| 10:6.209 | יועיל |  |
| 10:6.210 | להתפלל |  |
| 10:6.211 | אליהם |  |
| 10:6.212 | ולא |  |
| 10:6.213 | לעבדם |  |
| 10:6.214 | לא |  |
| 10:6.215 | תשמע |  |
| 10:6.216 | אל |  |
| 10:6.217 | דברי |  |
| 10:6.218 | הנביא |  |
| 10:6.219 | ההוא |  |
| 10:6.220 | לא |  |
| 10:6.221 | תפנה |  |
| 10:6.222 | אל |  |
| 10:6.223 | דבריו |  |
| 10:6.224 | לראות |  |
| 10:6.225 | אם |  |
| 10:6.226 | יש |  |
| 10:6.227 | בקצתם | , |
| 10:6.228 | כי |  |
| 10:6.229 | אין |  |
| 10:6.230 | ספק |  |
| 10:6.231 | שכל |  |
| 10:6.232 | דבריו |  |
| 10:6.233 | קשר |  |
| 10:6.234 | וכו | ' |
| 10:6.235 | עכ"ל | . |
| 10:6.236 | אשר |  |
| 10:6.237 | יאמר |  |
| 10:6.238 | שכל |  |
| 10:6.239 | שלא |  |
| 10:6.240 | בא |  |
| 10:6.241 | מופת |  |
| 10:6.242 | על |  |
| 10:6.243 | מציאותם |  |
| 10:6.244 | כצורות |  |
| 10:6.245 | ופרצופים |  |
| 10:6.246 | שהמציא |  |
| 10:6.247 | מחבר |  |
| 10:6.248 | הזהר |  |
| 10:6.249 | בנבואצו |  |
| 10:6.250 | שתלה |  |
| 10:6.251 | ברשב"י |  |
| 10:6.252 | התנא |  |
| 10:6.253 | הכל |  |
| 10:6.254 | שקר | ! |
| 10:6.255 | שהרי |  |
| 10:6.256 | אפי | ' |
| 10:6.257 | הנודעים |  |
| 10:6.258 | בחוש |  |
| 10:6.259 | או |  |
| 10:6.260 | במופת |  |
| 10:6.261 | פועלים |  |
| 10:6.262 | טבעיים |  |
| 10:6.263 | הם |  |
| 10:6.264 | ולא |  |
| 10:6.265 | רצוניים | , |
| 10:6.266 | עליהם |  |
| 10:6.267 | אמר |  |
| 10:6.268 | שמואל |  |
| 10:6.269 | הנביא |  |
| 10:6.270 | לישראל |  |
| 10:6.271 | ולא |  |
| 10:6.272 | תסורו |  |
| 10:6.273 | כי |  |
| 10:6.274 | אחרי |  |
| 10:6.275 | התהו |  |
| 10:6.276 | אשר |  |
| 10:6.277 | לא |  |
| 10:6.278 | יועילו |  |
| 10:6.279 | ולא |  |
| 10:6.280 | יצילו |  |
| 10:6.281 | כי |  |
| 10:6.282 | תהו |  |
| 10:6.283 | המה | ! |
| 10:6.284 | כי |  |
| 10:6.285 | כל |  |
| 10:6.286 | הפרצופים |  |
| 10:6.287 | והצורות |  |
| 10:6.288 | והעגולים |  |
| 10:6.289 | שהמציא |  |
| 10:6.290 | בדמיונו |  |
| 10:6.291 | תהו |  |
| 10:6.292 | והבל |  |
| 10:6.293 | המה | , |
| 10:6.294 | ואין |  |
| 10:6.295 | להם |  |
| 10:6.296 | מציאות |  |
| 10:6.297 | נרגש |  |
| 10:6.298 | בחוש |  |
| 10:6.299 | וגם |  |
| 10:6.300 | אין |  |
| 10:6.301 | מופת |  |
| 10:6.302 | על |  |
| 10:6.303 | מציאותם |  |
| 10:6.304 | בה | ' |
| 10:6.305 | אלהינו |  |
| 10:6.306 | שיש |  |
| 10:6.307 | כמה |  |
| 10:6.308 | מופתים |  |
| 10:6.309 | על |  |
| 10:6.310 | מציאותו |  |
| 10:6.311 | ית | ' |
| 10:6.312 | כמפורש |  |
| 10:6.313 | בתורה |  |
| 10:6.314 | ובנביאים |  |
| 10:6.315 | ובכתובים |  |
| 10:6.316 | ובדברי |  |
| 10:6.317 | רז"ל | ! |
| 10:6.318 | כי |  |
| 10:6.319 | אמנם |  |
| 10:6.320 | כל |  |
| 10:6.321 | דבריו |  |
| 10:6.322 | ספורי |  |
| 10:6.323 | נפלאות |  |
| 10:6.324 | ונמנעות |  |
| 10:6.325 | אשר |  |
| 10:6.326 | ישמטו |  |
| 10:6.327 | את |  |
| 10:6.328 | המשנה |  |
| 10:6.329 | והתלמוד | , |
| 10:6.330 | וישכינו |  |
| 10:6.331 | לעפר |  |
| 10:6.332 | כבודם | , |
| 10:6.333 | אף |  |
| 10:6.334 | כי |  |
| 10:6.335 | בת |  |
| 10:6.336 | מלך |  |
| 10:6.337 | היא |  |
| 10:6.338 | המשנה |  |
| 10:6.339 | ועליה |  |
| 10:6.340 | עמד |  |
| 10:6.341 | משה |  |
| 10:6.342 | רבינו |  |
| 10:6.343 | בהר |  |
| 10:6.344 | ארבעים |  |
| 10:6.345 | יום |  |
| 10:6.346 | ומ | ' |
| 10:6.347 | לילה |  |
| 10:6.348 | עד |  |
| 10:6.349 | שלמד |  |
| 10:6.350 | אותה |  |
| 10:6.351 | מפי |  |
| 10:6.352 | הגבורה |  |
| 10:6.353 | וירד |  |
| 10:6.354 | לישראל | ! |
| 10:6.355 | ונביא |  |
| 10:6.356 | השקר |  |
| 10:6.357 | מחבר |  |
| 10:6.358 | הזהר |  |
| 10:6.359 | בשנאתו |  |
| 10:6.360 | הגדולה |  |
| 10:6.361 | למשנה |  |
| 10:6.362 | ולתלמוד | , |
| 10:6.363 | הרחיה |  |
| 10:6.364 | עז |  |
| 10:6.365 | בנפשו |  |
| 10:6.366 | לבשות |  |
| 10:6.367 | את |  |
| 10:6.368 | המשנה |  |
| 10:6.369 | ולהוריד |  |
| 10:6.370 | לארץ |  |
| 10:6.371 | את |  |
| 10:6.372 | כבודה |  |
| 10:6.373 | כדי |  |
| 10:6.374 | לשכחה |  |
| 10:6.375 | ולעקור |  |
| 10:6.376 | אותה |  |
| 10:6.377 | ואת |  |
| 10:6.378 | חקותיה |  |
| 10:6.379 | מישראל |  |
| 10:6.380 | בתתו |  |
| 10:6.381 | אותה |  |
| 10:6.382 | לשמצה | ! |
| 10:6.383 | אף |  |
| 10:6.384 | כמעט |  |
| 10:6.385 | כל |  |
| 10:6.386 | קוראיו |  |
| 10:6.387 | לא |  |
| 10:6.388 | הבינו |  |
| 10:6.389 | את |  |
| 10:6.390 | מזמותיו | , |
| 10:6.391 | כי |  |
| 10:6.392 | חשבו |  |
| 10:6.393 | את |  |
| 10:6.394 | הכל |  |
| 10:6.395 | לתורה |  |
| 10:6.396 | אחת |  |
| 10:6.397 | ומשפט |  |
| 10:6.398 | אחד | , |
| 10:6.399 | ולא |  |
| 10:6.400 | הרפו |  |
| 10:6.401 | את |  |
| 10:6.402 | ידם |  |
| 10:6.403 | מלעסוק |  |
| 10:6.404 | בזה | , |
| 10:6.405 | וגם |  |
| 10:6.406 | מזה |  |
| 10:6.407 | לא |  |
| 10:6.408 | הניחו |  |
| 10:6.409 | ידם | ! |
| 10:6.410 | ישראל |  |
| 10:6.411 | לא |  |
| 10:6.412 | ידע | , |
| 10:6.413 | עמי |  |
| 10:6.414 | לא |  |
| 10:6.415 | התבונן | , |
| 10:6.416 | שכל |  |
| 10:6.417 | מטרתו |  |
| 10:6.418 | ומגמתו |  |
| 10:6.419 | של |  |
| 10:6.420 | הפילסוף |  |
| 10:6.421 | מחבר |  |
| 10:6.422 | הזהר |  |
| 10:6.423 | לתת |  |
| 10:6.424 | לנו |  |
| 10:6.425 | תורה |  |
| 10:6.426 | שבעל |  |
| 10:6.427 | פה |  |
| 10:6.428 | חדשה |  |
| 10:6.429 | תמורת |  |
| 10:6.430 | המשנה |  |
| 10:6.431 | והתלמוד | ! |
| 10:6.432 | מהלל |  |
| 10:6.433 | שם |  |
| 10:6.434 | ה | ' |
| 10:6.435 | החונן |  |
| 10:6.436 | לאדם |  |
| 10:6.437 | דעת |  |
| 10:6.438 | ותבונה | , |
| 10:6.439 | להבין |  |
| 10:6.440 | ולהשכיל |  |
| 10:6.441 | קשט |  |
| 10:6.442 | אמרי |  |
| 10:6.443 | אמת | , |
| 10:6.444 | ולנפץ |  |
| 10:6.445 | חזון |  |
| 10:6.446 | פריצים |  |
| 10:6.447 | אשר |  |
| 10:6.448 | ינשאו |  |
| 10:6.449 | להעמיד |  |
| 10:6.450 | חזון |  |
| 10:6.451 | ונכשלו |  |
| 10:6.452 | והכשילו |  |
| 10:6.453 | רבים |  |
| 10:6.454 | להאמין |  |
| 10:6.455 | בשוא |  |
| 10:6.456 | נתעה |  |
| 10:6.457 | כי |  |
| 10:6.458 | ה | ' |
| 10:6.459 | אלהינו |  |
| 10:6.460 | איננו |  |
| 10:6.461 | אחד |  |
| 10:6.462 | ומיוחד | , |
| 10:6.463 | על |  |
| 10:6.464 | ידי |  |
| 10:6.465 | לשון |  |
| 10:6.466 | רמיה |  |
| 10:6.467 | ושפת |  |
| 10:6.468 | חלקות | , |
| 10:6.469 | עתה |  |
| 10:6.470 | יראו |  |
| 10:6.471 | כל |  |
| 10:6.472 | חכמי |  |
| 10:6.473 | לב |  |
| 10:6.474 | וישכילו |  |
| 10:6.475 | מתורת |  |
| 10:6.476 | ה | ' |
| 10:6.477 | תמימה |  |
| 10:6.478 | שבכתב |  |
| 10:6.479 | ושבעל |  |
| 10:6.480 | פה |  |
| 10:6.481 | המשנה |  |
| 10:6.482 | והתלמוד |  |
| 10:6.483 | עם |  |
| 10:6.484 | מדרשי |  |
| 10:6.485 | רז"ל |  |
| 10:6.486 | האמתיים |  |
| 10:6.487 | ועם |  |
| 10:6.488 | פירושו |  |
| 10:6.489 | למשניות |  |
| 10:6.490 | והקדמותיו |  |
| 10:6.491 | עם |  |
| 10:6.492 | שאר |  |
| 10:6.493 | חכמינו |  |
| 10:6.494 | זכי |  |
| 10:6.495 | התבונה | . |
| 10:6.496 | ונפשו |  |
| 10:6.497 | תשבע |  |
| 10:6.498 | להם |  |
| 10:6.499 | חלות |  |
| 10:6.500 | טהורות |  |
| 10:6.501 | משוחות |  |
| 10:6.502 | בשטן |  |
| 10:6.503 | האמונה |  |
| 10:6.504 | הטהורה |  |
| 10:6.505 | כנתינתה |  |
| 10:6.506 | מסיני | . |
| 10:6.507 | ועל |  |
| 10:6.508 | ידה |  |
| 10:6.509 | יזכה |  |
| 10:6.510 | לחזות |  |
| 10:6.511 | בנועם |  |
| 10:6.512 | ה | ' |
| 10:6.513 | ולבקר |  |
| 10:6.514 | בהיכלו | , |
| 10:6.515 | עד |  |
| 10:6.516 | כה |  |
| 10:6.517 | עזרנו |  |
| 10:6.518 | השי"ת |  |
| 10:6.519 | לגלות |  |
| 10:6.520 | פני |  |
| 10:6.521 | הבלות |  |
| 10:6.522 | מחשבות |  |
| 10:6.523 | הפילסוף |  |
| 10:6.524 | נביא |  |
| 10:6.525 | השקר |  |
| 10:6.526 | המתעה |  |
| 10:6.527 | מחבר |  |
| 10:6.528 | הזהר | , |
| 10:6.529 | אשר |  |
| 10:6.530 | זמם |  |
| 10:6.531 | לדבר |  |
| 10:6.532 | סרה |  |
| 10:6.533 | על |  |
| 10:6.534 | ה | ' |
| 10:6.535 | ועל |  |
| 10:6.536 | תורתו |  |
| 10:6.537 | שבעל |  |
| 10:6.538 | פה |  |
| 10:6.539 | פיו |  |
| 10:6.540 | ידבר |  |
| 10:6.541 | תהפוכות |  |
| 10:6.542 | לבחור |  |
| 10:6.543 | אלהים |  |
| 10:6.544 | חדשים |  |
| 10:6.545 | שאר |  |
| 10:6.546 | בדא |  |
| 10:6.547 | אותם |  |
| 10:6.548 | בדמיונו | , |
| 10:6.549 | ונתלה |  |
| 10:6.550 | באילן |  |
| 10:6.551 | גדול |  |
| 10:6.552 | רשב"י |  |
| 10:6.553 | התנא |  |
| 10:6.554 | הקדוש | , |
| 10:6.555 | למען |  |
| 10:6.556 | תפוש |  |
| 10:6.557 | את |  |
| 10:6.558 | בית |  |
| 10:6.559 | ישראל |  |
| 10:6.560 | בלבם | ! |
| 10:6.561 | ועל |  |
| 10:6.562 | האמונה |  |
| 10:6.563 | הטהורה |  |
| 10:6.564 | המשנה |  |
| 10:6.565 | והתלמוד |  |
| 10:6.566 | ישליך |  |
| 10:6.567 | שקוצים |  |
| 10:6.568 | לשומה |  |
| 10:6.569 | שפחה |  |
| 10:6.570 | חרופה |  |
| 10:6.571 | ולהתעמר |  |
| 10:6.572 | בה | ! |
| 10:6.573 | וירומם |  |
| 10:6.574 | את |  |
| 10:6.575 | תורתו |  |
| 10:6.576 | החדשה |  |
| 10:6.577 | בלשון |  |
| 10:6.578 | רמיה |  |
| 10:6.579 | ושפת |  |
| 10:6.580 | חלקות |  |
| 10:6.581 | לכל |  |
| 10:6.582 | פתי |  |
| 10:6.583 | ותועה |  |
| 10:6.584 | רוח | ! |
| 10:6.585 | ברוך |  |
| 10:6.586 | החונן |  |
| 10:6.587 | לאדם |  |
| 10:6.588 | דעת |  |
| 10:6.589 | להשכיל |  |
| 10:6.590 | להטיב | , |
| 10:6.591 | לכו |  |
| 10:6.592 | ונלכה |  |
| 10:6.593 | באור |  |
| 10:6.594 | ה | ' |
| 10:6.595 | היא |  |
| 10:6.596 | תורתו |  |
| 10:6.597 | הקדושה |  |
| 10:6.598 | והתמימה |  |
| 10:6.599 | והטהורה | , |
| 10:6.600 | אשר |  |
| 10:6.601 | תבונינו |  |
| 10:6.602 | ותלמדינו |  |
| 10:6.603 | דעת |  |
| 10:6.604 | את |  |
| 10:6.605 | ה | ' |
| 10:6.606 | אלהים |  |
| 10:6.607 | אמת | , |
| 10:6.608 | ועל |  |
| 10:6.609 | ידי |  |
| 10:6.610 | זה |  |
| 10:6.611 | נזכה |  |
| 10:6.612 | לחזות |  |
| 10:6.613 | בנועם |  |
| 10:6.614 | ה | ' |
| 10:6.615 | בעולם |  |
| 10:6.616 | הזה |  |
| 10:6.617 | ובעולם |  |
| 10:6.618 | הבא |  |
| 10:6.619 | אכי"ר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

