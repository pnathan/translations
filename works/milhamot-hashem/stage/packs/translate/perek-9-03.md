# Translation pack: milhamot-hashem perek-9-03 (Milchamot ha-Shem 9:3–9:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-03`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **תוה"ק** → "the holy Torah".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- בחלק **שפתיו** → "with the smoothness of his lips".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **סנהדרין** → "Sanhedrin 99"; סנהדרין רי"ף → "Sanhedrin, [in] the Rif".
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


## Milchamot ha-Shem 9:3

Hebrew (Wikisource): קכב) ומדברי שני המאורות הגדולים גדולי חכמי ישראל, הרמב"ם והסמ"ג מודעא רבה לאוריתא שלא לשמוע לדברי המלך העכו"ם אשר הביא לנו את הספר החתום בשם רשב"י הזהר המסית ומטעה אותנו לחשוב באלהינו כמה עלות וסבות ולהחליף ולהמיר את כבוד אלהינו אשר אין ראשית לראשיתו באל אחר קצר אפים (זעיר אנפין) אשר לפי דבריו ניתן לו הכח והממשלה מן הסבות אשר למעלה ממנו, ושהם צוו ופקדו את כל באי עולם לעבדו ולאהבה אותו באמרם כי הוא אלהינו ואותו נעבוד, ומבת זוגו נירא. חלילה לנו להאמין כי התנא רשב"י ע"ה יאמר כדברים האלה, ושהוא עשה את החבור הזה בתור תורה שבעל פה חדשה, ויצוה עלינו לשוב בתשובה שלא לעסוק במשנה ובתלמוד. אלא בתורתו החדשה הנותנת דופי במשנה ובתלמוד לקרותה שפחה וקליפה, וסלע אחרא, ואבן משכית, ושהעוסקים בה נקראים עוסקים בתורה שלא לשמה, ודומין לכלבין דצוחין הב הב, ויורשים גיהנם האומרת הב, הב, כמו שנתבאר לעיל. חלילה לכל איש ישראל תמים עם ה' אלהיו ותורתו שבכתב ושבעל פה המסורה וחקוקה במשנה ובתלמוד, להאמין בדברים כאלה למלך העכו"ם שמסר לנו את הספר המסית אותנו לעבוד אלהים אחרים ולמשוך את לבבינו לאהבה אותם בספור מעשיות נוראות ונפלאות הנזכרים בספר הזהר אשר לא בעינינו ראינו ולא אבותינו ספרו לנו, כי עתיק יומין נראה ונגלה לרשב"י ונתן לו רשות לגלות פנים בתורה שלא כהלכה המקובלת, שהקב"ה אחד מיוחד מכל האחדים. חלילה להתנא רשב"י ע"ה וחביריו מעשות זאת להדיחנו מאחרי ה' אלהינו בחלק שפתיו. וברוב לקחו אשר ליקט מדברי רז"ל ועירבם בתוך מליצותיו להטותינו מדרך הקדש ומעבודת אבינו מלכינו יוצר בראשית. כאשה זרה ונכריה, אשר עליה אמר שלמה המלך ע"ה לשמרך מאשה זרה מנכריה אמריה החליקה וכו'. הטתו ברוב לקחה בחלק שפתיה תדיחנו וכו', להאמין ולחשוב באלהינו גופות של אור. אשר מהם ובהם תלוים כל מצות התורה ואם לא הגופות האלה אין תורה ואין מצוה, וכל מעשה מצותינו וברכותינו ותפלותינו ושבחים שאנו משבחים ומפארים לשי"ת הוא רק לגופים האלו הנבראים, דהיינו החמשה פרצופים הנחלקים עוד לשנים עשר פרצופים וביחוד לקצר אפים ובת זוגו (זו"ן) ושהוא המלך המהולל בכל התשבחות, כי כן צוו אביו ואמו לעבדו ולהללו ולשבחו לפארו ולקרותו בשם יהוה. חלילה לכל רבותינו התנאים והאמוראים מראשם ועד סופם אשר ידענו בבירור את אמונתם הטהורה ואת חכמתם חסידותם וזהירותם. לחשוב עליהם שהאמינו בפילסופיא הקדומה הטמאה המתועבת המלאה הזיות והבלים להכניסה בדת תוה"ק ולמשחה בששר, ולעשות לה ציונים ורעיונות זרות הרחוקים מן השכל הישר כדי להכניסה בתורת אלהינו תורת משה רבינו הנקייה מכל דעה זרה ונפסדת, ולהעמיס בה רבוי באלהות, בלי שום אות ומופת. רק בספורים של שקר שנגלה אליו עתיקא קדישא ונשמת משה ואליהו, ונשמות תנאים ואמוראים וגאונים שעדין לא היו ולא נבראו בימי התנא הקדוש רשב"י ז"ל. ובודאי שספורים כאלה הם ספורים בדוים שלא היו ולא נבראו, ואיך נשמע ונקבל על ידי ספורים בדוים להמיר את ה' אלהינו ארך אפים ורב חסד באל אחר קצר אפים בתבנית אדם אוכל לחם העצבים? כדגרסינן בסנהדרין פרק אלו הן הנחנקין אמר רבי אבהו אמר רבי יוחנן בכל אם יאמר לך נביא (הכדוק לך רמב"ם פ"ז מיסוה"ת) עבור על דברי תורה (לפי שעה) שמע לו חוץ מע"ז, שאפי' מעמיד לך חמה באמצע רקיע אל תשמע לו, תניא רבי יוסי הגלילי אומר ראה עד היכן הגיע סוף עובדי אלילים שנתן להם ממשלה, אפי' אם מעמידים חמה באמצע רקיע אל תשמע להן מפני מה? כי מנסה ה' אלהיכם אתכם. אמר רבי עקיבה ח"ו שהקב"ה מעמיד חמה ולבנה לעוברי רצונו. הא אינו מדבר אלא במי שהיה נביא אמת, וחזר להיות נביא שקר ע"כ. והן הן דברי הרמב"ם והסמ"ג שאפי' יעמוד נביא ויעשה אותות ומופתים בשמים ובארץ כמשה בן עמרם ויאמר להחליף תורתינו, או לבטל או להוסיף שום מצוה לא נשמע לו.

| id | word | punctuation after |
|---|---|---|
| 9:3.1 | קכב | ) |
| 9:3.2 | ומדברי |  |
| 9:3.3 | שני |  |
| 9:3.4 | המאורות |  |
| 9:3.5 | הגדולים |  |
| 9:3.6 | גדולי |  |
| 9:3.7 | חכמי |  |
| 9:3.8 | ישראל | , |
| 9:3.9 | הרמב"ם |  |
| 9:3.10 | והסמ"ג |  |
| 9:3.11 | מודעא |  |
| 9:3.12 | רבה |  |
| 9:3.13 | לאוריתא |  |
| 9:3.14 | שלא |  |
| 9:3.15 | לשמוע |  |
| 9:3.16 | לדברי |  |
| 9:3.17 | המלך |  |
| 9:3.18 | העכו"ם |  |
| 9:3.19 | אשר |  |
| 9:3.20 | הביא |  |
| 9:3.21 | לנו |  |
| 9:3.22 | את |  |
| 9:3.23 | הספר |  |
| 9:3.24 | החתום |  |
| 9:3.25 | בשם |  |
| 9:3.26 | רשב"י |  |
| 9:3.27 | הזהר |  |
| 9:3.28 | המסית |  |
| 9:3.29 | ומטעה |  |
| 9:3.30 | אותנו |  |
| 9:3.31 | לחשוב |  |
| 9:3.32 | באלהינו |  |
| 9:3.33 | כמה |  |
| 9:3.34 | עלות |  |
| 9:3.35 | וסבות |  |
| 9:3.36 | ולהחליף |  |
| 9:3.37 | ולהמיר |  |
| 9:3.38 | את |  |
| 9:3.39 | כבוד |  |
| 9:3.40 | אלהינו |  |
| 9:3.41 | אשר |  |
| 9:3.42 | אין |  |
| 9:3.43 | ראשית |  |
| 9:3.44 | לראשיתו |  |
| 9:3.45 | באל |  |
| 9:3.46 | אחר |  |
| 9:3.47 | קצר |  |
| 9:3.48 | אפים |  |
| 9:3.49 | זעיר |  |
| 9:3.50 | אנפין | ) |
| 9:3.51 | אשר |  |
| 9:3.52 | לפי |  |
| 9:3.53 | דבריו |  |
| 9:3.54 | ניתן |  |
| 9:3.55 | לו |  |
| 9:3.56 | הכח |  |
| 9:3.57 | והממשלה |  |
| 9:3.58 | מן |  |
| 9:3.59 | הסבות |  |
| 9:3.60 | אשר |  |
| 9:3.61 | למעלה |  |
| 9:3.62 | ממנו | , |
| 9:3.63 | ושהם |  |
| 9:3.64 | צוו |  |
| 9:3.65 | ופקדו |  |
| 9:3.66 | את |  |
| 9:3.67 | כל |  |
| 9:3.68 | באי |  |
| 9:3.69 | עולם |  |
| 9:3.70 | לעבדו |  |
| 9:3.71 | ולאהבה |  |
| 9:3.72 | אותו |  |
| 9:3.73 | באמרם |  |
| 9:3.74 | כי |  |
| 9:3.75 | הוא |  |
| 9:3.76 | אלהינו |  |
| 9:3.77 | ואותו |  |
| 9:3.78 | נעבוד | , |
| 9:3.79 | ומבת |  |
| 9:3.80 | זוגו |  |
| 9:3.81 | נירא | . |
| 9:3.82 | חלילה |  |
| 9:3.83 | לנו |  |
| 9:3.84 | להאמין |  |
| 9:3.85 | כי |  |
| 9:3.86 | התנא |  |
| 9:3.87 | רשב"י |  |
| 9:3.88 | ע"ה |  |
| 9:3.89 | יאמר |  |
| 9:3.90 | כדברים |  |
| 9:3.91 | האלה | , |
| 9:3.92 | ושהוא |  |
| 9:3.93 | עשה |  |
| 9:3.94 | את |  |
| 9:3.95 | החבור |  |
| 9:3.96 | הזה |  |
| 9:3.97 | בתור |  |
| 9:3.98 | תורה |  |
| 9:3.99 | שבעל |  |
| 9:3.100 | פה |  |
| 9:3.101 | חדשה | , |
| 9:3.102 | ויצוה |  |
| 9:3.103 | עלינו |  |
| 9:3.104 | לשוב |  |
| 9:3.105 | בתשובה |  |
| 9:3.106 | שלא |  |
| 9:3.107 | לעסוק |  |
| 9:3.108 | במשנה |  |
| 9:3.109 | ובתלמוד | . |
| 9:3.110 | אלא |  |
| 9:3.111 | בתורתו |  |
| 9:3.112 | החדשה |  |
| 9:3.113 | הנותנת |  |
| 9:3.114 | דופי |  |
| 9:3.115 | במשנה |  |
| 9:3.116 | ובתלמוד |  |
| 9:3.117 | לקרותה |  |
| 9:3.118 | שפחה |  |
| 9:3.119 | וקליפה | , |
| 9:3.120 | וסלע |  |
| 9:3.121 | אחרא | , |
| 9:3.122 | ואבן |  |
| 9:3.123 | משכית | , |
| 9:3.124 | ושהעוסקים |  |
| 9:3.125 | בה |  |
| 9:3.126 | נקראים |  |
| 9:3.127 | עוסקים |  |
| 9:3.128 | בתורה |  |
| 9:3.129 | שלא |  |
| 9:3.130 | לשמה | , |
| 9:3.131 | ודומין |  |
| 9:3.132 | לכלבין |  |
| 9:3.133 | דצוחין |  |
| 9:3.134 | הב |  |
| 9:3.135 | הב | , |
| 9:3.136 | ויורשים |  |
| 9:3.137 | גיהנם |  |
| 9:3.138 | האומרת |  |
| 9:3.139 | הב | , |
| 9:3.140 | הב | , |
| 9:3.141 | כמו |  |
| 9:3.142 | שנתבאר |  |
| 9:3.143 | לעיל | . |
| 9:3.144 | חלילה |  |
| 9:3.145 | לכל |  |
| 9:3.146 | איש |  |
| 9:3.147 | ישראל |  |
| 9:3.148 | תמים |  |
| 9:3.149 | עם |  |
| 9:3.150 | ה | ' |
| 9:3.151 | אלהיו |  |
| 9:3.152 | ותורתו |  |
| 9:3.153 | שבכתב |  |
| 9:3.154 | ושבעל |  |
| 9:3.155 | פה |  |
| 9:3.156 | המסורה |  |
| 9:3.157 | וחקוקה |  |
| 9:3.158 | במשנה |  |
| 9:3.159 | ובתלמוד | , |
| 9:3.160 | להאמין |  |
| 9:3.161 | בדברים |  |
| 9:3.162 | כאלה |  |
| 9:3.163 | למלך |  |
| 9:3.164 | העכו"ם |  |
| 9:3.165 | שמסר |  |
| 9:3.166 | לנו |  |
| 9:3.167 | את |  |
| 9:3.168 | הספר |  |
| 9:3.169 | המסית |  |
| 9:3.170 | אותנו |  |
| 9:3.171 | לעבוד |  |
| 9:3.172 | אלהים |  |
| 9:3.173 | אחרים |  |
| 9:3.174 | ולמשוך |  |
| 9:3.175 | את |  |
| 9:3.176 | לבבינו |  |
| 9:3.177 | לאהבה |  |
| 9:3.178 | אותם |  |
| 9:3.179 | בספור |  |
| 9:3.180 | מעשיות |  |
| 9:3.181 | נוראות |  |
| 9:3.182 | ונפלאות |  |
| 9:3.183 | הנזכרים |  |
| 9:3.184 | בספר |  |
| 9:3.185 | הזהר |  |
| 9:3.186 | אשר |  |
| 9:3.187 | לא |  |
| 9:3.188 | בעינינו |  |
| 9:3.189 | ראינו |  |
| 9:3.190 | ולא |  |
| 9:3.191 | אבותינו |  |
| 9:3.192 | ספרו |  |
| 9:3.193 | לנו | , |
| 9:3.194 | כי |  |
| 9:3.195 | עתיק |  |
| 9:3.196 | יומין |  |
| 9:3.197 | נראה |  |
| 9:3.198 | ונגלה |  |
| 9:3.199 | לרשב"י |  |
| 9:3.200 | ונתן |  |
| 9:3.201 | לו |  |
| 9:3.202 | רשות |  |
| 9:3.203 | לגלות |  |
| 9:3.204 | פנים |  |
| 9:3.205 | בתורה |  |
| 9:3.206 | שלא |  |
| 9:3.207 | כהלכה |  |
| 9:3.208 | המקובלת | , |
| 9:3.209 | שהקב"ה |  |
| 9:3.210 | אחד |  |
| 9:3.211 | מיוחד |  |
| 9:3.212 | מכל |  |
| 9:3.213 | האחדים | . |
| 9:3.214 | חלילה |  |
| 9:3.215 | להתנא |  |
| 9:3.216 | רשב"י |  |
| 9:3.217 | ע"ה |  |
| 9:3.218 | וחביריו |  |
| 9:3.219 | מעשות |  |
| 9:3.220 | זאת |  |
| 9:3.221 | להדיחנו |  |
| 9:3.222 | מאחרי |  |
| 9:3.223 | ה | ' |
| 9:3.224 | אלהינו |  |
| 9:3.225 | בחלק |  |
| 9:3.226 | שפתיו | . |
| 9:3.227 | וברוב |  |
| 9:3.228 | לקחו |  |
| 9:3.229 | אשר |  |
| 9:3.230 | ליקט |  |
| 9:3.231 | מדברי |  |
| 9:3.232 | רז"ל |  |
| 9:3.233 | ועירבם |  |
| 9:3.234 | בתוך |  |
| 9:3.235 | מליצותיו |  |
| 9:3.236 | להטותינו |  |
| 9:3.237 | מדרך |  |
| 9:3.238 | הקדש |  |
| 9:3.239 | ומעבודת |  |
| 9:3.240 | אבינו |  |
| 9:3.241 | מלכינו |  |
| 9:3.242 | יוצר |  |
| 9:3.243 | בראשית | . |
| 9:3.244 | כאשה |  |
| 9:3.245 | זרה |  |
| 9:3.246 | ונכריה | , |
| 9:3.247 | אשר |  |
| 9:3.248 | עליה |  |
| 9:3.249 | אמר |  |
| 9:3.250 | שלמה |  |
| 9:3.251 | המלך |  |
| 9:3.252 | ע"ה |  |
| 9:3.253 | לשמרך |  |
| 9:3.254 | מאשה |  |
| 9:3.255 | זרה |  |
| 9:3.256 | מנכריה |  |
| 9:3.257 | אמריה |  |
| 9:3.258 | החליקה |  |
| 9:3.259 | וכו | '. |
| 9:3.260 | הטתו |  |
| 9:3.261 | ברוב |  |
| 9:3.262 | לקחה |  |
| 9:3.263 | בחלק |  |
| 9:3.264 | שפתיה |  |
| 9:3.265 | תדיחנו |  |
| 9:3.266 | וכו | ', |
| 9:3.267 | להאמין |  |
| 9:3.268 | ולחשוב |  |
| 9:3.269 | באלהינו |  |
| 9:3.270 | גופות |  |
| 9:3.271 | של |  |
| 9:3.272 | אור | . |
| 9:3.273 | אשר |  |
| 9:3.274 | מהם |  |
| 9:3.275 | ובהם |  |
| 9:3.276 | תלוים |  |
| 9:3.277 | כל |  |
| 9:3.278 | מצות |  |
| 9:3.279 | התורה |  |
| 9:3.280 | ואם |  |
| 9:3.281 | לא |  |
| 9:3.282 | הגופות |  |
| 9:3.283 | האלה |  |
| 9:3.284 | אין |  |
| 9:3.285 | תורה |  |
| 9:3.286 | ואין |  |
| 9:3.287 | מצוה | , |
| 9:3.288 | וכל |  |
| 9:3.289 | מעשה |  |
| 9:3.290 | מצותינו |  |
| 9:3.291 | וברכותינו |  |
| 9:3.292 | ותפלותינו |  |
| 9:3.293 | ושבחים |  |
| 9:3.294 | שאנו |  |
| 9:3.295 | משבחים |  |
| 9:3.296 | ומפארים |  |
| 9:3.297 | לשי"ת |  |
| 9:3.298 | הוא |  |
| 9:3.299 | רק |  |
| 9:3.300 | לגופים |  |
| 9:3.301 | האלו |  |
| 9:3.302 | הנבראים | , |
| 9:3.303 | דהיינו |  |
| 9:3.304 | החמשה |  |
| 9:3.305 | פרצופים |  |
| 9:3.306 | הנחלקים |  |
| 9:3.307 | עוד |  |
| 9:3.308 | לשנים |  |
| 9:3.309 | עשר |  |
| 9:3.310 | פרצופים |  |
| 9:3.311 | וביחוד |  |
| 9:3.312 | לקצר |  |
| 9:3.313 | אפים |  |
| 9:3.314 | ובת |  |
| 9:3.315 | זוגו |  |
| 9:3.316 | זו"ן | ) |
| 9:3.317 | ושהוא |  |
| 9:3.318 | המלך |  |
| 9:3.319 | המהולל |  |
| 9:3.320 | בכל |  |
| 9:3.321 | התשבחות | , |
| 9:3.322 | כי |  |
| 9:3.323 | כן |  |
| 9:3.324 | צוו |  |
| 9:3.325 | אביו |  |
| 9:3.326 | ואמו |  |
| 9:3.327 | לעבדו |  |
| 9:3.328 | ולהללו |  |
| 9:3.329 | ולשבחו |  |
| 9:3.330 | לפארו |  |
| 9:3.331 | ולקרותו |  |
| 9:3.332 | בשם |  |
| 9:3.333 | יהוה | . |
| 9:3.334 | חלילה |  |
| 9:3.335 | לכל |  |
| 9:3.336 | רבותינו |  |
| 9:3.337 | התנאים |  |
| 9:3.338 | והאמוראים |  |
| 9:3.339 | מראשם |  |
| 9:3.340 | ועד |  |
| 9:3.341 | סופם |  |
| 9:3.342 | אשר |  |
| 9:3.343 | ידענו |  |
| 9:3.344 | בבירור |  |
| 9:3.345 | את |  |
| 9:3.346 | אמונתם |  |
| 9:3.347 | הטהורה |  |
| 9:3.348 | ואת |  |
| 9:3.349 | חכמתם |  |
| 9:3.350 | חסידותם |  |
| 9:3.351 | וזהירותם | . |
| 9:3.352 | לחשוב |  |
| 9:3.353 | עליהם |  |
| 9:3.354 | שהאמינו |  |
| 9:3.355 | בפילסופיא |  |
| 9:3.356 | הקדומה |  |
| 9:3.357 | הטמאה |  |
| 9:3.358 | המתועבת |  |
| 9:3.359 | המלאה |  |
| 9:3.360 | הזיות |  |
| 9:3.361 | והבלים |  |
| 9:3.362 | להכניסה |  |
| 9:3.363 | בדת |  |
| 9:3.364 | תוה"ק |  |
| 9:3.365 | ולמשחה |  |
| 9:3.366 | בששר | , |
| 9:3.367 | ולעשות |  |
| 9:3.368 | לה |  |
| 9:3.369 | ציונים |  |
| 9:3.370 | ורעיונות |  |
| 9:3.371 | זרות |  |
| 9:3.372 | הרחוקים |  |
| 9:3.373 | מן |  |
| 9:3.374 | השכל |  |
| 9:3.375 | הישר |  |
| 9:3.376 | כדי |  |
| 9:3.377 | להכניסה |  |
| 9:3.378 | בתורת |  |
| 9:3.379 | אלהינו |  |
| 9:3.380 | תורת |  |
| 9:3.381 | משה |  |
| 9:3.382 | רבינו |  |
| 9:3.383 | הנקייה |  |
| 9:3.384 | מכל |  |
| 9:3.385 | דעה |  |
| 9:3.386 | זרה |  |
| 9:3.387 | ונפסדת | , |
| 9:3.388 | ולהעמיס |  |
| 9:3.389 | בה |  |
| 9:3.390 | רבוי |  |
| 9:3.391 | באלהות | , |
| 9:3.392 | בלי |  |
| 9:3.393 | שום |  |
| 9:3.394 | אות |  |
| 9:3.395 | ומופת | . |
| 9:3.396 | רק |  |
| 9:3.397 | בספורים |  |
| 9:3.398 | של |  |
| 9:3.399 | שקר |  |
| 9:3.400 | שנגלה |  |
| 9:3.401 | אליו |  |
| 9:3.402 | עתיקא |  |
| 9:3.403 | קדישא |  |
| 9:3.404 | ונשמת |  |
| 9:3.405 | משה |  |
| 9:3.406 | ואליהו | , |
| 9:3.407 | ונשמות |  |
| 9:3.408 | תנאים |  |
| 9:3.409 | ואמוראים |  |
| 9:3.410 | וגאונים |  |
| 9:3.411 | שעדין |  |
| 9:3.412 | לא |  |
| 9:3.413 | היו |  |
| 9:3.414 | ולא |  |
| 9:3.415 | נבראו |  |
| 9:3.416 | בימי |  |
| 9:3.417 | התנא |  |
| 9:3.418 | הקדוש |  |
| 9:3.419 | רשב"י |  |
| 9:3.420 | ז"ל | . |
| 9:3.421 | ובודאי |  |
| 9:3.422 | שספורים |  |
| 9:3.423 | כאלה |  |
| 9:3.424 | הם |  |
| 9:3.425 | ספורים |  |
| 9:3.426 | בדוים |  |
| 9:3.427 | שלא |  |
| 9:3.428 | היו |  |
| 9:3.429 | ולא |  |
| 9:3.430 | נבראו | , |
| 9:3.431 | ואיך |  |
| 9:3.432 | נשמע |  |
| 9:3.433 | ונקבל |  |
| 9:3.434 | על |  |
| 9:3.435 | ידי |  |
| 9:3.436 | ספורים |  |
| 9:3.437 | בדוים |  |
| 9:3.438 | להמיר |  |
| 9:3.439 | את |  |
| 9:3.440 | ה | ' |
| 9:3.441 | אלהינו |  |
| 9:3.442 | ארך |  |
| 9:3.443 | אפים |  |
| 9:3.444 | ורב |  |
| 9:3.445 | חסד |  |
| 9:3.446 | באל |  |
| 9:3.447 | אחר |  |
| 9:3.448 | קצר |  |
| 9:3.449 | אפים |  |
| 9:3.450 | בתבנית |  |
| 9:3.451 | אדם |  |
| 9:3.452 | אוכל |  |
| 9:3.453 | לחם |  |
| 9:3.454 | העצבים | ? |
| 9:3.455 | כדגרסינן |  |
| 9:3.456 | בסנהדרין |  |
| 9:3.457 | פרק |  |
| 9:3.458 | אלו |  |
| 9:3.459 | הן |  |
| 9:3.460 | הנחנקין |  |
| 9:3.461 | אמר |  |
| 9:3.462 | רבי |  |
| 9:3.463 | אבהו |  |
| 9:3.464 | אמר |  |
| 9:3.465 | רבי |  |
| 9:3.466 | יוחנן |  |
| 9:3.467 | בכל |  |
| 9:3.468 | אם |  |
| 9:3.469 | יאמר |  |
| 9:3.470 | לך |  |
| 9:3.471 | נביא |  |
| 9:3.472 | הכדוק |  |
| 9:3.473 | לך |  |
| 9:3.474 | רמב"ם |  |
| 9:3.475 | פ"ז |  |
| 9:3.476 | מיסוה"ת | ) |
| 9:3.477 | עבור |  |
| 9:3.478 | על |  |
| 9:3.479 | דברי |  |
| 9:3.480 | תורה |  |
| 9:3.481 | לפי |  |
| 9:3.482 | שעה | ) |
| 9:3.483 | שמע |  |
| 9:3.484 | לו |  |
| 9:3.485 | חוץ |  |
| 9:3.486 | מע"ז | , |
| 9:3.487 | שאפי | ' |
| 9:3.488 | מעמיד |  |
| 9:3.489 | לך |  |
| 9:3.490 | חמה |  |
| 9:3.491 | באמצע |  |
| 9:3.492 | רקיע |  |
| 9:3.493 | אל |  |
| 9:3.494 | תשמע |  |
| 9:3.495 | לו | , |
| 9:3.496 | תניא |  |
| 9:3.497 | רבי |  |
| 9:3.498 | יוסי |  |
| 9:3.499 | הגלילי |  |
| 9:3.500 | אומר |  |
| 9:3.501 | ראה |  |
| 9:3.502 | עד |  |
| 9:3.503 | היכן |  |
| 9:3.504 | הגיע |  |
| 9:3.505 | סוף |  |
| 9:3.506 | עובדי |  |
| 9:3.507 | אלילים |  |
| 9:3.508 | שנתן |  |
| 9:3.509 | להם |  |
| 9:3.510 | ממשלה | , |
| 9:3.511 | אפי | ' |
| 9:3.512 | אם |  |
| 9:3.513 | מעמידים |  |
| 9:3.514 | חמה |  |
| 9:3.515 | באמצע |  |
| 9:3.516 | רקיע |  |
| 9:3.517 | אל |  |
| 9:3.518 | תשמע |  |
| 9:3.519 | להן |  |
| 9:3.520 | מפני |  |
| 9:3.521 | מה | ? |
| 9:3.522 | כי |  |
| 9:3.523 | מנסה |  |
| 9:3.524 | ה | ' |
| 9:3.525 | אלהיכם |  |
| 9:3.526 | אתכם | . |
| 9:3.527 | אמר |  |
| 9:3.528 | רבי |  |
| 9:3.529 | עקיבה |  |
| 9:3.530 | ח"ו |  |
| 9:3.531 | שהקב"ה |  |
| 9:3.532 | מעמיד |  |
| 9:3.533 | חמה |  |
| 9:3.534 | ולבנה |  |
| 9:3.535 | לעוברי |  |
| 9:3.536 | רצונו | . |
| 9:3.537 | הא |  |
| 9:3.538 | אינו |  |
| 9:3.539 | מדבר |  |
| 9:3.540 | אלא |  |
| 9:3.541 | במי |  |
| 9:3.542 | שהיה |  |
| 9:3.543 | נביא |  |
| 9:3.544 | אמת | , |
| 9:3.545 | וחזר |  |
| 9:3.546 | להיות |  |
| 9:3.547 | נביא |  |
| 9:3.548 | שקר |  |
| 9:3.549 | ע"כ | . |
| 9:3.550 | והן |  |
| 9:3.551 | הן |  |
| 9:3.552 | דברי |  |
| 9:3.553 | הרמב"ם |  |
| 9:3.554 | והסמ"ג |  |
| 9:3.555 | שאפי | ' |
| 9:3.556 | יעמוד |  |
| 9:3.557 | נביא |  |
| 9:3.558 | ויעשה |  |
| 9:3.559 | אותות |  |
| 9:3.560 | ומופתים |  |
| 9:3.561 | בשמים |  |
| 9:3.562 | ובארץ |  |
| 9:3.563 | כמשה |  |
| 9:3.564 | בן |  |
| 9:3.565 | עמרם |  |
| 9:3.566 | ויאמר |  |
| 9:3.567 | להחליף |  |
| 9:3.568 | תורתינו | , |
| 9:3.569 | או |  |
| 9:3.570 | לבטל |  |
| 9:3.571 | או |  |
| 9:3.572 | להוסיף |  |
| 9:3.573 | שום |  |
| 9:3.574 | מצוה |  |
| 9:3.575 | לא |  |
| 9:3.576 | נשמע |  |
| 9:3.577 | לו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

