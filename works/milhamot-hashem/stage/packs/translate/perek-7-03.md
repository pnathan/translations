# Translation pack: milhamot-hashem perek-7-03 (Milchamot ha-Shem 7:4–7:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-03`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
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

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
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
- הסבה **הראשונה** → "the First Cause".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
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
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ריבה / לרבות** → "include", never "multiply".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- דברי **תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אחד ... **ואחד** of two kinds of case → "alike".
- ואין **צריך** לומר → "needless to say".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
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


## Milchamot ha-Shem 7:4

Hebrew (Wikisource): פא) זכור נזכור את אמונת רז"ל בעלי המשנה והתלמוד הגאונים ז"ל שהיא האמונה הטהורה שהאמינו בה קדמונינו, כי ה' אלהינו אחד, אין עוד מלבדו ואין כיחודו, ותשוח עלינו נפשינו בראותינו עתה ספרי הקבלה החדשה כמה אלוהות נתחדשו להם, אשר לא שערום אבותינו, והמה בכתובים גם בסידורי התפלות כדי לכוון תפלתם אליהם, וכן באשר הברכות, ולא די להם המון האלוהות שמכונים אליהם כמו שכבר ביארנו, אלא שאיוו לאלוהות רבים, אשר אין לקרות בהם בשם, שהרי לפי דעתם גם אדם קדמון ואדם קדמאה אין לכוון תפלתו אליהם, כי הם נעלמים מאד, וכל שכן פרצופים שלמעלה מעולם האצילות אלפים ורבבות אין מספר, עולמות שלא דברו בהם, וכל עולם ועולם כלול מאבי"ע, ומי יודע אם יש בכל עולם ועולם כל הפרצופים שהזכירו באבי"ע זה התחתון כי לא מצינו זה להם בפירוש לפי שהעלימו ולא דברו בהם, ורק זה מצינו להר' מקדש מלך בראשי' דך ט"ו דבור המתחיל ורמז, כי החלק המקיף וסמוך לטהירו נקרא אדם קדמאה סתימאה, להבדיל בינו ובין אדם קדמון, ואדם קדמאה סתימאה דא נעשה עשרה כדורים זה בתוך זה, ובכל אחד מכדורים אלו נבנו כמה עולמות אין מספר וכו'. ונמצא לפי זה מערכת האלוהות הרבים אשר כל אחד ואחד עלה וסבה למה שלמטה ממנו, ראשון לכלם האין סוף המקיף את כל העולמות, וכולם בתוך חלל שלו, ואחריו אדם קדמאה סתימה, ואחריו אדם קדמון, הוא שאמר לפי דבריהם ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין צריך ליטול רשות מאדם קדמאה ולא מאין סוף, ואחריו עתיק יומין, ואחריו אריך אנפין והוא אשר ברא שמים וארץ בכח אבא ואימא, ואחריו אבא, הוא שאמר יהי אור, ויקוו המים, ויהי רקיע, ושאר אמירות דמעשה בראשית. ואימא אומן העושה על פיו ככל אשר יצוה עליה. חוץ ממאמר נעשה אדם. שאימא היא שאמרה לאבא נעשה אדם בצלמינו. ואבא לא חפץ בבריאתו. לפי שגלוי לפניו שעתיד לחטוא כמו שכתבתי לעיל סי' כ'. ואחר אבא ואימא זעיר אנפין ונוקביה. ואל זעיר ונוקביה לפי דבריהם כל פיות פונים לעבדו. וליראה את מלכתו היא נוקבא דיליה שריבה אותה מחבר הזהר ליראה מאת ה' אלהיך תירא. את לרבותה, ליראה ממנה ולא לעבדה (וד"א כמו שדרש ר' עקיבה לרבות ת"ח). ונתנו המקובלים החדשים טעם למה לא הזכירו בזהר ובתקונים את הפרצופים העליונים כי אם ברמז, שהוא לרוב העלמם. והוא בעצמו הטעם למה אין מתפללים אליהם, ולא קורים אליהם בעת צרה מטעם רוב העלמם. ולפי דבריהם לא שייכא עבודה ותפלה אלא לפרצוף שכבר נתעבה ונתגלה וכמו שכתב בספר הברית שהוא מושג לבית יעקב. כלומר שכבר נתעבה ונתגלה כמו שכתב הרח"ו והובא לעיל סי' מ"ז. ועל הטעם הזה יפלא מאד הלא ה' אלהינו שהוא הסבה הראשונה אשר עבדו אבותינו ואבות אבותינו נעלם הוא מאד עד שאפי' המלאכים וחיות הקדש שואלים איה מקום כבודו כמו שאמרו רז"ל ועם כל זה אמר הכתוב כי מי גוי גדול אשר לו אלהים קרובים אליו כה' אלהינו בכל קראנו אליו. כמפה לאזן כמו שנת' לעיל:

| id | word | punctuation after |
|---|---|---|
| 7:4.1 | פא | ) |
| 7:4.2 | זכור |  |
| 7:4.3 | נזכור |  |
| 7:4.4 | את |  |
| 7:4.5 | אמונת |  |
| 7:4.6 | רז"ל |  |
| 7:4.7 | בעלי |  |
| 7:4.8 | המשנה |  |
| 7:4.9 | והתלמוד |  |
| 7:4.10 | הגאונים |  |
| 7:4.11 | ז"ל |  |
| 7:4.12 | שהיא |  |
| 7:4.13 | האמונה |  |
| 7:4.14 | הטהורה |  |
| 7:4.15 | שהאמינו |  |
| 7:4.16 | בה |  |
| 7:4.17 | קדמונינו | , |
| 7:4.18 | כי |  |
| 7:4.19 | ה | ' |
| 7:4.20 | אלהינו |  |
| 7:4.21 | אחד | , |
| 7:4.22 | אין |  |
| 7:4.23 | עוד |  |
| 7:4.24 | מלבדו |  |
| 7:4.25 | ואין |  |
| 7:4.26 | כיחודו | , |
| 7:4.27 | ותשוח |  |
| 7:4.28 | עלינו |  |
| 7:4.29 | נפשינו |  |
| 7:4.30 | בראותינו |  |
| 7:4.31 | עתה |  |
| 7:4.32 | ספרי |  |
| 7:4.33 | הקבלה |  |
| 7:4.34 | החדשה |  |
| 7:4.35 | כמה |  |
| 7:4.36 | אלוהות |  |
| 7:4.37 | נתחדשו |  |
| 7:4.38 | להם | , |
| 7:4.39 | אשר |  |
| 7:4.40 | לא |  |
| 7:4.41 | שערום |  |
| 7:4.42 | אבותינו | , |
| 7:4.43 | והמה |  |
| 7:4.44 | בכתובים |  |
| 7:4.45 | גם |  |
| 7:4.46 | בסידורי |  |
| 7:4.47 | התפלות |  |
| 7:4.48 | כדי |  |
| 7:4.49 | לכוון |  |
| 7:4.50 | תפלתם |  |
| 7:4.51 | אליהם | , |
| 7:4.52 | וכן |  |
| 7:4.53 | באשר |  |
| 7:4.54 | הברכות | , |
| 7:4.55 | ולא |  |
| 7:4.56 | די |  |
| 7:4.57 | להם |  |
| 7:4.58 | המון |  |
| 7:4.59 | האלוהות |  |
| 7:4.60 | שמכונים |  |
| 7:4.61 | אליהם |  |
| 7:4.62 | כמו |  |
| 7:4.63 | שכבר |  |
| 7:4.64 | ביארנו | , |
| 7:4.65 | אלא |  |
| 7:4.66 | שאיוו |  |
| 7:4.67 | לאלוהות |  |
| 7:4.68 | רבים | , |
| 7:4.69 | אשר |  |
| 7:4.70 | אין |  |
| 7:4.71 | לקרות |  |
| 7:4.72 | בהם |  |
| 7:4.73 | בשם | , |
| 7:4.74 | שהרי |  |
| 7:4.75 | לפי |  |
| 7:4.76 | דעתם |  |
| 7:4.77 | גם |  |
| 7:4.78 | אדם |  |
| 7:4.79 | קדמון |  |
| 7:4.80 | ואדם |  |
| 7:4.81 | קדמאה |  |
| 7:4.82 | אין |  |
| 7:4.83 | לכוון |  |
| 7:4.84 | תפלתו |  |
| 7:4.85 | אליהם | , |
| 7:4.86 | כי |  |
| 7:4.87 | הם |  |
| 7:4.88 | נעלמים |  |
| 7:4.89 | מאד | , |
| 7:4.90 | וכל |  |
| 7:4.91 | שכן |  |
| 7:4.92 | פרצופים |  |
| 7:4.93 | שלמעלה |  |
| 7:4.94 | מעולם |  |
| 7:4.95 | האצילות |  |
| 7:4.96 | אלפים |  |
| 7:4.97 | ורבבות |  |
| 7:4.98 | אין |  |
| 7:4.99 | מספר | , |
| 7:4.100 | עולמות |  |
| 7:4.101 | שלא |  |
| 7:4.102 | דברו |  |
| 7:4.103 | בהם | , |
| 7:4.104 | וכל |  |
| 7:4.105 | עולם |  |
| 7:4.106 | ועולם |  |
| 7:4.107 | כלול |  |
| 7:4.108 | מאבי"ע | , |
| 7:4.109 | ומי |  |
| 7:4.110 | יודע |  |
| 7:4.111 | אם |  |
| 7:4.112 | יש |  |
| 7:4.113 | בכל |  |
| 7:4.114 | עולם |  |
| 7:4.115 | ועולם |  |
| 7:4.116 | כל |  |
| 7:4.117 | הפרצופים |  |
| 7:4.118 | שהזכירו |  |
| 7:4.119 | באבי"ע |  |
| 7:4.120 | זה |  |
| 7:4.121 | התחתון |  |
| 7:4.122 | כי |  |
| 7:4.123 | לא |  |
| 7:4.124 | מצינו |  |
| 7:4.125 | זה |  |
| 7:4.126 | להם |  |
| 7:4.127 | בפירוש |  |
| 7:4.128 | לפי |  |
| 7:4.129 | שהעלימו |  |
| 7:4.130 | ולא |  |
| 7:4.131 | דברו |  |
| 7:4.132 | בהם | , |
| 7:4.133 | ורק |  |
| 7:4.134 | זה |  |
| 7:4.135 | מצינו |  |
| 7:4.136 | להר | ' |
| 7:4.137 | מקדש |  |
| 7:4.138 | מלך |  |
| 7:4.139 | בראשי | ' |
| 7:4.140 | דך |  |
| 7:4.141 | ט"ו |  |
| 7:4.142 | דבור |  |
| 7:4.143 | המתחיל |  |
| 7:4.144 | ורמז | , |
| 7:4.145 | כי |  |
| 7:4.146 | החלק |  |
| 7:4.147 | המקיף |  |
| 7:4.148 | וסמוך |  |
| 7:4.149 | לטהירו |  |
| 7:4.150 | נקרא |  |
| 7:4.151 | אדם |  |
| 7:4.152 | קדמאה |  |
| 7:4.153 | סתימאה | , |
| 7:4.154 | להבדיל |  |
| 7:4.155 | בינו |  |
| 7:4.156 | ובין |  |
| 7:4.157 | אדם |  |
| 7:4.158 | קדמון | , |
| 7:4.159 | ואדם |  |
| 7:4.160 | קדמאה |  |
| 7:4.161 | סתימאה |  |
| 7:4.162 | דא |  |
| 7:4.163 | נעשה |  |
| 7:4.164 | עשרה |  |
| 7:4.165 | כדורים |  |
| 7:4.166 | זה |  |
| 7:4.167 | בתוך |  |
| 7:4.168 | זה | , |
| 7:4.169 | ובכל |  |
| 7:4.170 | אחד |  |
| 7:4.171 | מכדורים |  |
| 7:4.172 | אלו |  |
| 7:4.173 | נבנו |  |
| 7:4.174 | כמה |  |
| 7:4.175 | עולמות |  |
| 7:4.176 | אין |  |
| 7:4.177 | מספר |  |
| 7:4.178 | וכו | '. |
| 7:4.179 | ונמצא |  |
| 7:4.180 | לפי |  |
| 7:4.181 | זה |  |
| 7:4.182 | מערכת |  |
| 7:4.183 | האלוהות |  |
| 7:4.184 | הרבים |  |
| 7:4.185 | אשר |  |
| 7:4.186 | כל |  |
| 7:4.187 | אחד |  |
| 7:4.188 | ואחד |  |
| 7:4.189 | עלה |  |
| 7:4.190 | וסבה |  |
| 7:4.191 | למה |  |
| 7:4.192 | שלמטה |  |
| 7:4.193 | ממנו | , |
| 7:4.194 | ראשון |  |
| 7:4.195 | לכלם |  |
| 7:4.196 | האין |  |
| 7:4.197 | סוף |  |
| 7:4.198 | המקיף |  |
| 7:4.199 | את |  |
| 7:4.200 | כל |  |
| 7:4.201 | העולמות | , |
| 7:4.202 | וכולם |  |
| 7:4.203 | בתוך |  |
| 7:4.204 | חלל |  |
| 7:4.205 | שלו | , |
| 7:4.206 | ואחריו |  |
| 7:4.207 | אדם |  |
| 7:4.208 | קדמאה |  |
| 7:4.209 | סתימה | , |
| 7:4.210 | ואחריו |  |
| 7:4.211 | אדם |  |
| 7:4.212 | קדמון | , |
| 7:4.213 | הוא |  |
| 7:4.214 | שאמר |  |
| 7:4.215 | לפי |  |
| 7:4.216 | דבריהם |  |
| 7:4.217 | ראו |  |
| 7:4.218 | עתה |  |
| 7:4.219 | כי |  |
| 7:4.220 | אני |  |
| 7:4.221 | אני |  |
| 7:4.222 | הוא |  |
| 7:4.223 | ואין |  |
| 7:4.224 | אלהים |  |
| 7:4.225 | עמדי | , |
| 7:4.226 | שאין |  |
| 7:4.227 | צריך |  |
| 7:4.228 | ליטול |  |
| 7:4.229 | רשות |  |
| 7:4.230 | מאדם |  |
| 7:4.231 | קדמאה |  |
| 7:4.232 | ולא |  |
| 7:4.233 | מאין |  |
| 7:4.234 | סוף | , |
| 7:4.235 | ואחריו |  |
| 7:4.236 | עתיק |  |
| 7:4.237 | יומין | , |
| 7:4.238 | ואחריו |  |
| 7:4.239 | אריך |  |
| 7:4.240 | אנפין |  |
| 7:4.241 | והוא |  |
| 7:4.242 | אשר |  |
| 7:4.243 | ברא |  |
| 7:4.244 | שמים |  |
| 7:4.245 | וארץ |  |
| 7:4.246 | בכח |  |
| 7:4.247 | אבא |  |
| 7:4.248 | ואימא | , |
| 7:4.249 | ואחריו |  |
| 7:4.250 | אבא | , |
| 7:4.251 | הוא |  |
| 7:4.252 | שאמר |  |
| 7:4.253 | יהי |  |
| 7:4.254 | אור | , |
| 7:4.255 | ויקוו |  |
| 7:4.256 | המים | , |
| 7:4.257 | ויהי |  |
| 7:4.258 | רקיע | , |
| 7:4.259 | ושאר |  |
| 7:4.260 | אמירות |  |
| 7:4.261 | דמעשה |  |
| 7:4.262 | בראשית | . |
| 7:4.263 | ואימא |  |
| 7:4.264 | אומן |  |
| 7:4.265 | העושה |  |
| 7:4.266 | על |  |
| 7:4.267 | פיו |  |
| 7:4.268 | ככל |  |
| 7:4.269 | אשר |  |
| 7:4.270 | יצוה |  |
| 7:4.271 | עליה | . |
| 7:4.272 | חוץ |  |
| 7:4.273 | ממאמר |  |
| 7:4.274 | נעשה |  |
| 7:4.275 | אדם | . |
| 7:4.276 | שאימא |  |
| 7:4.277 | היא |  |
| 7:4.278 | שאמרה |  |
| 7:4.279 | לאבא |  |
| 7:4.280 | נעשה |  |
| 7:4.281 | אדם |  |
| 7:4.282 | בצלמינו | . |
| 7:4.283 | ואבא |  |
| 7:4.284 | לא |  |
| 7:4.285 | חפץ |  |
| 7:4.286 | בבריאתו | . |
| 7:4.287 | לפי |  |
| 7:4.288 | שגלוי |  |
| 7:4.289 | לפניו |  |
| 7:4.290 | שעתיד |  |
| 7:4.291 | לחטוא |  |
| 7:4.292 | כמו |  |
| 7:4.293 | שכתבתי |  |
| 7:4.294 | לעיל |  |
| 7:4.295 | סי | ' |
| 7:4.296 | כ | '. |
| 7:4.297 | ואחר |  |
| 7:4.298 | אבא |  |
| 7:4.299 | ואימא |  |
| 7:4.300 | זעיר |  |
| 7:4.301 | אנפין |  |
| 7:4.302 | ונוקביה | . |
| 7:4.303 | ואל |  |
| 7:4.304 | זעיר |  |
| 7:4.305 | ונוקביה |  |
| 7:4.306 | לפי |  |
| 7:4.307 | דבריהם |  |
| 7:4.308 | כל |  |
| 7:4.309 | פיות |  |
| 7:4.310 | פונים |  |
| 7:4.311 | לעבדו | . |
| 7:4.312 | וליראה |  |
| 7:4.313 | את |  |
| 7:4.314 | מלכתו |  |
| 7:4.315 | היא |  |
| 7:4.316 | נוקבא |  |
| 7:4.317 | דיליה |  |
| 7:4.318 | שריבה |  |
| 7:4.319 | אותה |  |
| 7:4.320 | מחבר |  |
| 7:4.321 | הזהר |  |
| 7:4.322 | ליראה |  |
| 7:4.323 | מאת |  |
| 7:4.324 | ה | ' |
| 7:4.325 | אלהיך |  |
| 7:4.326 | תירא | . |
| 7:4.327 | את |  |
| 7:4.328 | לרבותה | , |
| 7:4.329 | ליראה |  |
| 7:4.330 | ממנה |  |
| 7:4.331 | ולא |  |
| 7:4.332 | לעבדה |  |
| 7:4.333 | וד"א |  |
| 7:4.334 | כמו |  |
| 7:4.335 | שדרש |  |
| 7:4.336 | ר | ' |
| 7:4.337 | עקיבה |  |
| 7:4.338 | לרבות |  |
| 7:4.339 | ת"ח | ). |
| 7:4.340 | ונתנו |  |
| 7:4.341 | המקובלים |  |
| 7:4.342 | החדשים |  |
| 7:4.343 | טעם |  |
| 7:4.344 | למה |  |
| 7:4.345 | לא |  |
| 7:4.346 | הזכירו |  |
| 7:4.347 | בזהר |  |
| 7:4.348 | ובתקונים |  |
| 7:4.349 | את |  |
| 7:4.350 | הפרצופים |  |
| 7:4.351 | העליונים |  |
| 7:4.352 | כי |  |
| 7:4.353 | אם |  |
| 7:4.354 | ברמז | , |
| 7:4.355 | שהוא |  |
| 7:4.356 | לרוב |  |
| 7:4.357 | העלמם | . |
| 7:4.358 | והוא |  |
| 7:4.359 | בעצמו |  |
| 7:4.360 | הטעם |  |
| 7:4.361 | למה |  |
| 7:4.362 | אין |  |
| 7:4.363 | מתפללים |  |
| 7:4.364 | אליהם | , |
| 7:4.365 | ולא |  |
| 7:4.366 | קורים |  |
| 7:4.367 | אליהם |  |
| 7:4.368 | בעת |  |
| 7:4.369 | צרה |  |
| 7:4.370 | מטעם |  |
| 7:4.371 | רוב |  |
| 7:4.372 | העלמם | . |
| 7:4.373 | ולפי |  |
| 7:4.374 | דבריהם |  |
| 7:4.375 | לא |  |
| 7:4.376 | שייכא |  |
| 7:4.377 | עבודה |  |
| 7:4.378 | ותפלה |  |
| 7:4.379 | אלא |  |
| 7:4.380 | לפרצוף |  |
| 7:4.381 | שכבר |  |
| 7:4.382 | נתעבה |  |
| 7:4.383 | ונתגלה |  |
| 7:4.384 | וכמו |  |
| 7:4.385 | שכתב |  |
| 7:4.386 | בספר |  |
| 7:4.387 | הברית |  |
| 7:4.388 | שהוא |  |
| 7:4.389 | מושג |  |
| 7:4.390 | לבית |  |
| 7:4.391 | יעקב | . |
| 7:4.392 | כלומר |  |
| 7:4.393 | שכבר |  |
| 7:4.394 | נתעבה |  |
| 7:4.395 | ונתגלה |  |
| 7:4.396 | כמו |  |
| 7:4.397 | שכתב |  |
| 7:4.398 | הרח"ו |  |
| 7:4.399 | והובא |  |
| 7:4.400 | לעיל |  |
| 7:4.401 | סי | ' |
| 7:4.402 | מ"ז | . |
| 7:4.403 | ועל |  |
| 7:4.404 | הטעם |  |
| 7:4.405 | הזה |  |
| 7:4.406 | יפלא |  |
| 7:4.407 | מאד |  |
| 7:4.408 | הלא |  |
| 7:4.409 | ה | ' |
| 7:4.410 | אלהינו |  |
| 7:4.411 | שהוא |  |
| 7:4.412 | הסבה |  |
| 7:4.413 | הראשונה |  |
| 7:4.414 | אשר |  |
| 7:4.415 | עבדו |  |
| 7:4.416 | אבותינו |  |
| 7:4.417 | ואבות |  |
| 7:4.418 | אבותינו |  |
| 7:4.419 | נעלם |  |
| 7:4.420 | הוא |  |
| 7:4.421 | מאד |  |
| 7:4.422 | עד |  |
| 7:4.423 | שאפי | ' |
| 7:4.424 | המלאכים |  |
| 7:4.425 | וחיות |  |
| 7:4.426 | הקדש |  |
| 7:4.427 | שואלים |  |
| 7:4.428 | איה |  |
| 7:4.429 | מקום |  |
| 7:4.430 | כבודו |  |
| 7:4.431 | כמו |  |
| 7:4.432 | שאמרו |  |
| 7:4.433 | רז"ל |  |
| 7:4.434 | ועם |  |
| 7:4.435 | כל |  |
| 7:4.436 | זה |  |
| 7:4.437 | אמר |  |
| 7:4.438 | הכתוב |  |
| 7:4.439 | כי |  |
| 7:4.440 | מי |  |
| 7:4.441 | גוי |  |
| 7:4.442 | גדול |  |
| 7:4.443 | אשר |  |
| 7:4.444 | לו |  |
| 7:4.445 | אלהים |  |
| 7:4.446 | קרובים |  |
| 7:4.447 | אליו |  |
| 7:4.448 | כה | ' |
| 7:4.449 | אלהינו |  |
| 7:4.450 | בכל |  |
| 7:4.451 | קראנו |  |
| 7:4.452 | אליו | . |
| 7:4.453 | כמפה |  |
| 7:4.454 | לאזן |  |
| 7:4.455 | כמו |  |
| 7:4.456 | שנת | ' |
| 7:4.457 | לעיל | : |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

