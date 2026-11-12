# Translation pack: milhamot-hashem perek-4-21 (Milchamot ha-Shem 4:33–4:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-21`.

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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
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

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
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
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- עלת **העלות** → "the Cause of causes".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
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
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השתלשלות** → "the chain of descent".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- גוף **דק** → "a subtle body".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
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
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
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


## Milchamot ha-Shem 4:33

Hebrew (Wikisource): לט) הנה מכל זה מבואר באר היטב, כי גוף הספירות נמשכו ונשתלשלו דרך הקו הדק מעצמות האין סוף, גם העגולים גם היושר, ונמצא לפי זה כי גוף הספירות עם האור הפנימי שלהן שהוא נשמתן של ספירות עם האור המקיף שהוא מלבושן הכל מעצם אחד מן האין סוף מתפשט ונמשך, וכבר בא דבר זה מפורש בהדיא בשושן סודות בסדר התפלה וגם בסדר שבעת ימי הפסח וז"ל דע לך כי עשר ספירות אינם דבר נברא אלא ענין מתפשט מעצם הבורא ית' ואינם נבדלים ממנו והוא תמיד נמצא בהם כהדין קמצא דלבושיה מניה וכו' ע"כ, וכ"כ בגליון הזהר בשם הרמ"ז פר' בהר דף ק"ט שכלי האצילות יש להם בחינת אלהות, וכ"כ הר' משנת חסידים בסוף מסכת העשייה וז"ל והאצילות כולו בין בבחינת אורות בין בבחינת כלים ואפי' לבושיו הם כולם אלהות גמור (וכן כתב הר' עוז לאלהים בבית קדש הקדשים שהגוף והנשמה חד אינון וכן כתב בעל היכל הברכה פרשת ויקרא באוצר החיים דף ז' עמוד ד') אבל בריאה יצירה עשייה מהרוח שלהם ולמטה וגם הרוח בכלל אינם אלהות גמור עכ"ל משנת חסידים. וכן מתבאר מלשון הזהר שהבאתי לעיל סי' כ"א שהוא קורא לכל א' מן הפרצופים עלת העלות לפי שהוא עילה לעילות שתחתיו אבל אדם קדמון נקרא עילת כל העילות לפי שהוא עילה ראשונה לכל הפרצופים כלם כדלעיל. ומכאן תשובה לדבריך שהכחשת ואמרת ח"ו שהמקובלים אמרו כן, אדרבה אמרו שאין להאמין ולהעלות על לב שהספירות הם חלק מהאין סוף ב"ה שיצא ונשתלשל מעילה לעלול וכו' שדבריך אלה סותרים דברי הזהר ומהרח"ו ושושן סודות והרמ"ז, ומשנת חסידים שביארו איכות השתלשלותם מעצם האין סוף ושהם אלוהות גמור, והרי אתה מודה שדבריהם שאמרו כן הוא עוון פלילי. וכבר הבאתי לעיל דברי הזהר שקורא לאבא ואמא ה' צבאות, ולזעיר אנפין בן לאבא ואימא. ושם אמרו שכל עבודתנו וכל תפלותנו לזעיר אנפין זה שהוא בן לאבא ואימא. וכל המקובלים מסכימים כן. כמו שכתבתי לעיל בשם מקדש מלך והרש"ב ויושר לבב וספר הברית ומהרח"ו בעץ חיים והאר"י בס' הליקוטים וכסא אליהו ומצרף האמונה ונחלת יוסף. כולהו בשטה חדא קיימי שכל השבחים וכל הברכות הם לזעיר אנפין דוקא. ולא לאין סוף ולא לשאר פרצופים שלמעלה מזעיר בעולם האצילות ולא לפרצופים שלמטה ממנו בעולמות בריאה יצירה עשייה, אע"פ שהניחו ואמרו שבכל אחד מהשלשה עולמות הנז' יש בו כל הפרצופים שבעולם האצילות ומחומר אחד קורצו מעצם האין סוף נמשכו ונשתלשלו דרך הקו היוצא ממנו ויורד למטה כמו שנתבאר לעיל רק שאינם אלוהות גמור להתפלל אליהם ולקרוא להם בעת צרה כאשר יראה המעיין בדבריהם. ועל הדברים האלה והאמונות הנז' אמרו רז"ל בסנהדרין דך ס"א ע"ב מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים, מה לי קרובים ומה לי רחוקים? אלא הכי קאמר מטיבתן של קרובים אתה למד מה טיבן של רחוקים ע"כ. ופי' רש"י ז"ל לפיכך אני אומר לך, הסתכל בע"ז הקרובה לך ותראה שאין בהם ממש, ומהן אתה למד טיבן של רחוקים עכ"ל. הראו לנו רז"ל דרך בחינת צלמי ע"ז שמטיבם של קרובים נדע מה טיבם של רחוקים ומזה נקיש ונדון על כל פרצופי הקבלה החדשה האומרת שבכל עולם מארבעה עולמות נמצאים הפרצופים הנזכרים, ועכ"ז אין העבודה אלא לזעיר אנפין דאצילות. שכשם שהקרובים אלינו בעולם העשייה והיצירה והבריאה אין לעבדם שאין בהם ממש ואינם אלהות, כן זעיר אנפין דאצילות הרחוק אין בו ממש, ואין אלוה מבלעדי אלהינו. ומאחר שביררנו שהסכמת כל המקובלים החדשים שכל העבודה והשבחים והברכות לזעיר אנפין שבאצילות הנה כל טענותיך שטענת עלינו וכל הכחשותיך כולם דברי הבל ורעות רוח כאומר על האיש שהוא אשה ועל העמוד של שיש שהוא של זהב וכו'. ובתשובה זו שהשבת עשית עצמך כגולם כמאמר התנא בפ' בתרא דאבות ז' דברים בחכם וחילופיהן בגולם, ומכללן החכם שואל כהלכה ומשיב כענין ובגולם בחילוף זה. וכן אתה עשית, כי השואל שאל אותך כהלכה את מי נעבוד על פי הקבלה החדשה את עתיק, או אריך אנפין, או אבא ואימא וכו'. ובאה תשובתך שלא כענין כי קבלתך וכו' כאילו שאל אותך על קבלתך. וגם הרבית דברים והארכת להביא רמזים וגימטריות מדברי חוות יאיר שלא הבין דברי אביו. והזכיר רמזים שיוכל כל אדם לעשות כמותן מבלי קבלה כמו שעשה ר' שמואל שחיבר ספר כח ה' שהוא מלא צירופים וגימטריות ורמזים בלילה אחת מדעתו מבלי קבלה, רק להראות כחו ויכלתו לעשות צירופים וגימטריות ורמזים בנקל ככתוב בספר א"נ. ובשלמא רמזים וצירופים כל שאינו נוגע בכבוד המב"ה לעבוד אלהים אחרים ולזלזל במעלת המשנה והתלמוד ומדרשי רז"ל המספרים שבחי תורתינו הקדושה דליכא מידי דלא רמיזה באוריתא. ובפרט זה הארכת לציין הרמזים שכ' בעל מנורת המאור למקומותם, ואת המבוקש ממך לא ענית, ולא עלה על דל שפתיך כמדת הגולם שהזכיר התנא ע"ה שאינו יודע להשיב כענין המבוקש ממנו.

| id | word | punctuation after |
|---|---|---|
| 4:33.1 | לט | ) |
| 4:33.2 | הנה |  |
| 4:33.3 | מכל |  |
| 4:33.4 | זה |  |
| 4:33.5 | מבואר |  |
| 4:33.6 | באר |  |
| 4:33.7 | היטב | , |
| 4:33.8 | כי |  |
| 4:33.9 | גוף |  |
| 4:33.10 | הספירות |  |
| 4:33.11 | נמשכו |  |
| 4:33.12 | ונשתלשלו |  |
| 4:33.13 | דרך |  |
| 4:33.14 | הקו |  |
| 4:33.15 | הדק |  |
| 4:33.16 | מעצמות |  |
| 4:33.17 | האין |  |
| 4:33.18 | סוף | , |
| 4:33.19 | גם |  |
| 4:33.20 | העגולים |  |
| 4:33.21 | גם |  |
| 4:33.22 | היושר | , |
| 4:33.23 | ונמצא |  |
| 4:33.24 | לפי |  |
| 4:33.25 | זה |  |
| 4:33.26 | כי |  |
| 4:33.27 | גוף |  |
| 4:33.28 | הספירות |  |
| 4:33.29 | עם |  |
| 4:33.30 | האור |  |
| 4:33.31 | הפנימי |  |
| 4:33.32 | שלהן |  |
| 4:33.33 | שהוא |  |
| 4:33.34 | נשמתן |  |
| 4:33.35 | של |  |
| 4:33.36 | ספירות |  |
| 4:33.37 | עם |  |
| 4:33.38 | האור |  |
| 4:33.39 | המקיף |  |
| 4:33.40 | שהוא |  |
| 4:33.41 | מלבושן |  |
| 4:33.42 | הכל |  |
| 4:33.43 | מעצם |  |
| 4:33.44 | אחד |  |
| 4:33.45 | מן |  |
| 4:33.46 | האין |  |
| 4:33.47 | סוף |  |
| 4:33.48 | מתפשט |  |
| 4:33.49 | ונמשך | , |
| 4:33.50 | וכבר |  |
| 4:33.51 | בא |  |
| 4:33.52 | דבר |  |
| 4:33.53 | זה |  |
| 4:33.54 | מפורש |  |
| 4:33.55 | בהדיא |  |
| 4:33.56 | בשושן |  |
| 4:33.57 | סודות |  |
| 4:33.58 | בסדר |  |
| 4:33.59 | התפלה |  |
| 4:33.60 | וגם |  |
| 4:33.61 | בסדר |  |
| 4:33.62 | שבעת |  |
| 4:33.63 | ימי |  |
| 4:33.64 | הפסח |  |
| 4:33.65 | וז"ל |  |
| 4:33.66 | דע |  |
| 4:33.67 | לך |  |
| 4:33.68 | כי |  |
| 4:33.69 | עשר |  |
| 4:33.70 | ספירות |  |
| 4:33.71 | אינם |  |
| 4:33.72 | דבר |  |
| 4:33.73 | נברא |  |
| 4:33.74 | אלא |  |
| 4:33.75 | ענין |  |
| 4:33.76 | מתפשט |  |
| 4:33.77 | מעצם |  |
| 4:33.78 | הבורא |  |
| 4:33.79 | ית | ' |
| 4:33.80 | ואינם |  |
| 4:33.81 | נבדלים |  |
| 4:33.82 | ממנו |  |
| 4:33.83 | והוא |  |
| 4:33.84 | תמיד |  |
| 4:33.85 | נמצא |  |
| 4:33.86 | בהם |  |
| 4:33.87 | כהדין |  |
| 4:33.88 | קמצא |  |
| 4:33.89 | דלבושיה |  |
| 4:33.90 | מניה |  |
| 4:33.91 | וכו | ' |
| 4:33.92 | ע"כ | , |
| 4:33.93 | וכ"כ |  |
| 4:33.94 | בגליון |  |
| 4:33.95 | הזהר |  |
| 4:33.96 | בשם |  |
| 4:33.97 | הרמ"ז |  |
| 4:33.98 | פר | ' |
| 4:33.99 | בהר |  |
| 4:33.100 | דף |  |
| 4:33.101 | ק"ט |  |
| 4:33.102 | שכלי |  |
| 4:33.103 | האצילות |  |
| 4:33.104 | יש |  |
| 4:33.105 | להם |  |
| 4:33.106 | בחינת |  |
| 4:33.107 | אלהות | , |
| 4:33.108 | וכ"כ |  |
| 4:33.109 | הר | ' |
| 4:33.110 | משנת |  |
| 4:33.111 | חסידים |  |
| 4:33.112 | בסוף |  |
| 4:33.113 | מסכת |  |
| 4:33.114 | העשייה |  |
| 4:33.115 | וז"ל |  |
| 4:33.116 | והאצילות |  |
| 4:33.117 | כולו |  |
| 4:33.118 | בין |  |
| 4:33.119 | בבחינת |  |
| 4:33.120 | אורות |  |
| 4:33.121 | בין |  |
| 4:33.122 | בבחינת |  |
| 4:33.123 | כלים |  |
| 4:33.124 | ואפי | ' |
| 4:33.125 | לבושיו |  |
| 4:33.126 | הם |  |
| 4:33.127 | כולם |  |
| 4:33.128 | אלהות |  |
| 4:33.129 | גמור |  |
| 4:33.130 | וכן |  |
| 4:33.131 | כתב |  |
| 4:33.132 | הר | ' |
| 4:33.133 | עוז |  |
| 4:33.134 | לאלהים |  |
| 4:33.135 | בבית |  |
| 4:33.136 | קדש |  |
| 4:33.137 | הקדשים |  |
| 4:33.138 | שהגוף |  |
| 4:33.139 | והנשמה |  |
| 4:33.140 | חד |  |
| 4:33.141 | אינון |  |
| 4:33.142 | וכן |  |
| 4:33.143 | כתב |  |
| 4:33.144 | בעל |  |
| 4:33.145 | היכל |  |
| 4:33.146 | הברכה |  |
| 4:33.147 | פרשת |  |
| 4:33.148 | ויקרא |  |
| 4:33.149 | באוצר |  |
| 4:33.150 | החיים |  |
| 4:33.151 | דף |  |
| 4:33.152 | ז | ' |
| 4:33.153 | עמוד |  |
| 4:33.154 | ד | ') |
| 4:33.155 | אבל |  |
| 4:33.156 | בריאה |  |
| 4:33.157 | יצירה |  |
| 4:33.158 | עשייה |  |
| 4:33.159 | מהרוח |  |
| 4:33.160 | שלהם |  |
| 4:33.161 | ולמטה |  |
| 4:33.162 | וגם |  |
| 4:33.163 | הרוח |  |
| 4:33.164 | בכלל |  |
| 4:33.165 | אינם |  |
| 4:33.166 | אלהות |  |
| 4:33.167 | גמור |  |
| 4:33.168 | עכ"ל |  |
| 4:33.169 | משנת |  |
| 4:33.170 | חסידים | . |
| 4:33.171 | וכן |  |
| 4:33.172 | מתבאר |  |
| 4:33.173 | מלשון |  |
| 4:33.174 | הזהר |  |
| 4:33.175 | שהבאתי |  |
| 4:33.176 | לעיל |  |
| 4:33.177 | סי | ' |
| 4:33.178 | כ"א |  |
| 4:33.179 | שהוא |  |
| 4:33.180 | קורא |  |
| 4:33.181 | לכל |  |
| 4:33.182 | א | ' |
| 4:33.183 | מן |  |
| 4:33.184 | הפרצופים |  |
| 4:33.185 | עלת |  |
| 4:33.186 | העלות |  |
| 4:33.187 | לפי |  |
| 4:33.188 | שהוא |  |
| 4:33.189 | עילה |  |
| 4:33.190 | לעילות |  |
| 4:33.191 | שתחתיו |  |
| 4:33.192 | אבל |  |
| 4:33.193 | אדם |  |
| 4:33.194 | קדמון |  |
| 4:33.195 | נקרא |  |
| 4:33.196 | עילת |  |
| 4:33.197 | כל |  |
| 4:33.198 | העילות |  |
| 4:33.199 | לפי |  |
| 4:33.200 | שהוא |  |
| 4:33.201 | עילה |  |
| 4:33.202 | ראשונה |  |
| 4:33.203 | לכל |  |
| 4:33.204 | הפרצופים |  |
| 4:33.205 | כלם |  |
| 4:33.206 | כדלעיל | . |
| 4:33.207 | ומכאן |  |
| 4:33.208 | תשובה |  |
| 4:33.209 | לדבריך |  |
| 4:33.210 | שהכחשת |  |
| 4:33.211 | ואמרת |  |
| 4:33.212 | ח"ו |  |
| 4:33.213 | שהמקובלים |  |
| 4:33.214 | אמרו |  |
| 4:33.215 | כן | , |
| 4:33.216 | אדרבה |  |
| 4:33.217 | אמרו |  |
| 4:33.218 | שאין |  |
| 4:33.219 | להאמין |  |
| 4:33.220 | ולהעלות |  |
| 4:33.221 | על |  |
| 4:33.222 | לב |  |
| 4:33.223 | שהספירות |  |
| 4:33.224 | הם |  |
| 4:33.225 | חלק |  |
| 4:33.226 | מהאין |  |
| 4:33.227 | סוף |  |
| 4:33.228 | ב"ה |  |
| 4:33.229 | שיצא |  |
| 4:33.230 | ונשתלשל |  |
| 4:33.231 | מעילה |  |
| 4:33.232 | לעלול |  |
| 4:33.233 | וכו | ' |
| 4:33.234 | שדבריך |  |
| 4:33.235 | אלה |  |
| 4:33.236 | סותרים |  |
| 4:33.237 | דברי |  |
| 4:33.238 | הזהר |  |
| 4:33.239 | ומהרח"ו |  |
| 4:33.240 | ושושן |  |
| 4:33.241 | סודות |  |
| 4:33.242 | והרמ"ז | , |
| 4:33.243 | ומשנת |  |
| 4:33.244 | חסידים |  |
| 4:33.245 | שביארו |  |
| 4:33.246 | איכות |  |
| 4:33.247 | השתלשלותם |  |
| 4:33.248 | מעצם |  |
| 4:33.249 | האין |  |
| 4:33.250 | סוף |  |
| 4:33.251 | ושהם |  |
| 4:33.252 | אלוהות |  |
| 4:33.253 | גמור | , |
| 4:33.254 | והרי |  |
| 4:33.255 | אתה |  |
| 4:33.256 | מודה |  |
| 4:33.257 | שדבריהם |  |
| 4:33.258 | שאמרו |  |
| 4:33.259 | כן |  |
| 4:33.260 | הוא |  |
| 4:33.261 | עוון |  |
| 4:33.262 | פלילי | . |
| 4:33.263 | וכבר |  |
| 4:33.264 | הבאתי |  |
| 4:33.265 | לעיל |  |
| 4:33.266 | דברי |  |
| 4:33.267 | הזהר |  |
| 4:33.268 | שקורא |  |
| 4:33.269 | לאבא |  |
| 4:33.270 | ואמא |  |
| 4:33.271 | ה | ' |
| 4:33.272 | צבאות | , |
| 4:33.273 | ולזעיר |  |
| 4:33.274 | אנפין |  |
| 4:33.275 | בן |  |
| 4:33.276 | לאבא |  |
| 4:33.277 | ואימא | . |
| 4:33.278 | ושם |  |
| 4:33.279 | אמרו |  |
| 4:33.280 | שכל |  |
| 4:33.281 | עבודתנו |  |
| 4:33.282 | וכל |  |
| 4:33.283 | תפלותנו |  |
| 4:33.284 | לזעיר |  |
| 4:33.285 | אנפין |  |
| 4:33.286 | זה |  |
| 4:33.287 | שהוא |  |
| 4:33.288 | בן |  |
| 4:33.289 | לאבא |  |
| 4:33.290 | ואימא | . |
| 4:33.291 | וכל |  |
| 4:33.292 | המקובלים |  |
| 4:33.293 | מסכימים |  |
| 4:33.294 | כן | . |
| 4:33.295 | כמו |  |
| 4:33.296 | שכתבתי |  |
| 4:33.297 | לעיל |  |
| 4:33.298 | בשם |  |
| 4:33.299 | מקדש |  |
| 4:33.300 | מלך |  |
| 4:33.301 | והרש"ב |  |
| 4:33.302 | ויושר |  |
| 4:33.303 | לבב |  |
| 4:33.304 | וספר |  |
| 4:33.305 | הברית |  |
| 4:33.306 | ומהרח"ו |  |
| 4:33.307 | בעץ |  |
| 4:33.308 | חיים |  |
| 4:33.309 | והאר"י |  |
| 4:33.310 | בס | ' |
| 4:33.311 | הליקוטים |  |
| 4:33.312 | וכסא |  |
| 4:33.313 | אליהו |  |
| 4:33.314 | ומצרף |  |
| 4:33.315 | האמונה |  |
| 4:33.316 | ונחלת |  |
| 4:33.317 | יוסף | . |
| 4:33.318 | כולהו |  |
| 4:33.319 | בשטה |  |
| 4:33.320 | חדא |  |
| 4:33.321 | קיימי |  |
| 4:33.322 | שכל |  |
| 4:33.323 | השבחים |  |
| 4:33.324 | וכל |  |
| 4:33.325 | הברכות |  |
| 4:33.326 | הם |  |
| 4:33.327 | לזעיר |  |
| 4:33.328 | אנפין |  |
| 4:33.329 | דוקא | . |
| 4:33.330 | ולא |  |
| 4:33.331 | לאין |  |
| 4:33.332 | סוף |  |
| 4:33.333 | ולא |  |
| 4:33.334 | לשאר |  |
| 4:33.335 | פרצופים |  |
| 4:33.336 | שלמעלה |  |
| 4:33.337 | מזעיר |  |
| 4:33.338 | בעולם |  |
| 4:33.339 | האצילות |  |
| 4:33.340 | ולא |  |
| 4:33.341 | לפרצופים |  |
| 4:33.342 | שלמטה |  |
| 4:33.343 | ממנו |  |
| 4:33.344 | בעולמות |  |
| 4:33.345 | בריאה |  |
| 4:33.346 | יצירה |  |
| 4:33.347 | עשייה | , |
| 4:33.348 | אע"פ |  |
| 4:33.349 | שהניחו |  |
| 4:33.350 | ואמרו |  |
| 4:33.351 | שבכל |  |
| 4:33.352 | אחד |  |
| 4:33.353 | מהשלשה |  |
| 4:33.354 | עולמות |  |
| 4:33.355 | הנז | ' |
| 4:33.356 | יש |  |
| 4:33.357 | בו |  |
| 4:33.358 | כל |  |
| 4:33.359 | הפרצופים |  |
| 4:33.360 | שבעולם |  |
| 4:33.361 | האצילות |  |
| 4:33.362 | ומחומר |  |
| 4:33.363 | אחד |  |
| 4:33.364 | קורצו |  |
| 4:33.365 | מעצם |  |
| 4:33.366 | האין |  |
| 4:33.367 | סוף |  |
| 4:33.368 | נמשכו |  |
| 4:33.369 | ונשתלשלו |  |
| 4:33.370 | דרך |  |
| 4:33.371 | הקו |  |
| 4:33.372 | היוצא |  |
| 4:33.373 | ממנו |  |
| 4:33.374 | ויורד |  |
| 4:33.375 | למטה |  |
| 4:33.376 | כמו |  |
| 4:33.377 | שנתבאר |  |
| 4:33.378 | לעיל |  |
| 4:33.379 | רק |  |
| 4:33.380 | שאינם |  |
| 4:33.381 | אלוהות |  |
| 4:33.382 | גמור |  |
| 4:33.383 | להתפלל |  |
| 4:33.384 | אליהם |  |
| 4:33.385 | ולקרוא |  |
| 4:33.386 | להם |  |
| 4:33.387 | בעת |  |
| 4:33.388 | צרה |  |
| 4:33.389 | כאשר |  |
| 4:33.390 | יראה |  |
| 4:33.391 | המעיין |  |
| 4:33.392 | בדבריהם | . |
| 4:33.393 | ועל |  |
| 4:33.394 | הדברים |  |
| 4:33.395 | האלה |  |
| 4:33.396 | והאמונות |  |
| 4:33.397 | הנז | ' |
| 4:33.398 | אמרו |  |
| 4:33.399 | רז"ל |  |
| 4:33.400 | בסנהדרין |  |
| 4:33.401 | דך |  |
| 4:33.402 | ס"א |  |
| 4:33.403 | ע"ב |  |
| 4:33.404 | מאלהי |  |
| 4:33.405 | העמים |  |
| 4:33.406 | אשר |  |
| 4:33.407 | סביבותיכם |  |
| 4:33.408 | הקרובים |  |
| 4:33.409 | אליך |  |
| 4:33.410 | או |  |
| 4:33.411 | הרחוקים | , |
| 4:33.412 | מה |  |
| 4:33.413 | לי |  |
| 4:33.414 | קרובים |  |
| 4:33.415 | ומה |  |
| 4:33.416 | לי |  |
| 4:33.417 | רחוקים | ? |
| 4:33.418 | אלא |  |
| 4:33.419 | הכי |  |
| 4:33.420 | קאמר |  |
| 4:33.421 | מטיבתן |  |
| 4:33.422 | של |  |
| 4:33.423 | קרובים |  |
| 4:33.424 | אתה |  |
| 4:33.425 | למד |  |
| 4:33.426 | מה |  |
| 4:33.427 | טיבן |  |
| 4:33.428 | של |  |
| 4:33.429 | רחוקים |  |
| 4:33.430 | ע"כ | . |
| 4:33.431 | ופי | ' |
| 4:33.432 | רש"י |  |
| 4:33.433 | ז"ל |  |
| 4:33.434 | לפיכך |  |
| 4:33.435 | אני |  |
| 4:33.436 | אומר |  |
| 4:33.437 | לך | , |
| 4:33.438 | הסתכל |  |
| 4:33.439 | בע"ז |  |
| 4:33.440 | הקרובה |  |
| 4:33.441 | לך |  |
| 4:33.442 | ותראה |  |
| 4:33.443 | שאין |  |
| 4:33.444 | בהם |  |
| 4:33.445 | ממש | , |
| 4:33.446 | ומהן |  |
| 4:33.447 | אתה |  |
| 4:33.448 | למד |  |
| 4:33.449 | טיבן |  |
| 4:33.450 | של |  |
| 4:33.451 | רחוקים |  |
| 4:33.452 | עכ"ל | . |
| 4:33.453 | הראו |  |
| 4:33.454 | לנו |  |
| 4:33.455 | רז"ל |  |
| 4:33.456 | דרך |  |
| 4:33.457 | בחינת |  |
| 4:33.458 | צלמי |  |
| 4:33.459 | ע"ז |  |
| 4:33.460 | שמטיבם |  |
| 4:33.461 | של |  |
| 4:33.462 | קרובים |  |
| 4:33.463 | נדע |  |
| 4:33.464 | מה |  |
| 4:33.465 | טיבם |  |
| 4:33.466 | של |  |
| 4:33.467 | רחוקים |  |
| 4:33.468 | ומזה |  |
| 4:33.469 | נקיש |  |
| 4:33.470 | ונדון |  |
| 4:33.471 | על |  |
| 4:33.472 | כל |  |
| 4:33.473 | פרצופי |  |
| 4:33.474 | הקבלה |  |
| 4:33.475 | החדשה |  |
| 4:33.476 | האומרת |  |
| 4:33.477 | שבכל |  |
| 4:33.478 | עולם |  |
| 4:33.479 | מארבעה |  |
| 4:33.480 | עולמות |  |
| 4:33.481 | נמצאים |  |
| 4:33.482 | הפרצופים |  |
| 4:33.483 | הנזכרים | , |
| 4:33.484 | ועכ"ז |  |
| 4:33.485 | אין |  |
| 4:33.486 | העבודה |  |
| 4:33.487 | אלא |  |
| 4:33.488 | לזעיר |  |
| 4:33.489 | אנפין |  |
| 4:33.490 | דאצילות | . |
| 4:33.491 | שכשם |  |
| 4:33.492 | שהקרובים |  |
| 4:33.493 | אלינו |  |
| 4:33.494 | בעולם |  |
| 4:33.495 | העשייה |  |
| 4:33.496 | והיצירה |  |
| 4:33.497 | והבריאה |  |
| 4:33.498 | אין |  |
| 4:33.499 | לעבדם |  |
| 4:33.500 | שאין |  |
| 4:33.501 | בהם |  |
| 4:33.502 | ממש |  |
| 4:33.503 | ואינם |  |
| 4:33.504 | אלהות | , |
| 4:33.505 | כן |  |
| 4:33.506 | זעיר |  |
| 4:33.507 | אנפין |  |
| 4:33.508 | דאצילות |  |
| 4:33.509 | הרחוק |  |
| 4:33.510 | אין |  |
| 4:33.511 | בו |  |
| 4:33.512 | ממש | , |
| 4:33.513 | ואין |  |
| 4:33.514 | אלוה |  |
| 4:33.515 | מבלעדי |  |
| 4:33.516 | אלהינו | . |
| 4:33.517 | ומאחר |  |
| 4:33.518 | שביררנו |  |
| 4:33.519 | שהסכמת |  |
| 4:33.520 | כל |  |
| 4:33.521 | המקובלים |  |
| 4:33.522 | החדשים |  |
| 4:33.523 | שכל |  |
| 4:33.524 | העבודה |  |
| 4:33.525 | והשבחים |  |
| 4:33.526 | והברכות |  |
| 4:33.527 | לזעיר |  |
| 4:33.528 | אנפין |  |
| 4:33.529 | שבאצילות |  |
| 4:33.530 | הנה |  |
| 4:33.531 | כל |  |
| 4:33.532 | טענותיך |  |
| 4:33.533 | שטענת |  |
| 4:33.534 | עלינו |  |
| 4:33.535 | וכל |  |
| 4:33.536 | הכחשותיך |  |
| 4:33.537 | כולם |  |
| 4:33.538 | דברי |  |
| 4:33.539 | הבל |  |
| 4:33.540 | ורעות |  |
| 4:33.541 | רוח |  |
| 4:33.542 | כאומר |  |
| 4:33.543 | על |  |
| 4:33.544 | האיש |  |
| 4:33.545 | שהוא |  |
| 4:33.546 | אשה |  |
| 4:33.547 | ועל |  |
| 4:33.548 | העמוד |  |
| 4:33.549 | של |  |
| 4:33.550 | שיש |  |
| 4:33.551 | שהוא |  |
| 4:33.552 | של |  |
| 4:33.553 | זהב |  |
| 4:33.554 | וכו | '. |
| 4:33.555 | ובתשובה |  |
| 4:33.556 | זו |  |
| 4:33.557 | שהשבת |  |
| 4:33.558 | עשית |  |
| 4:33.559 | עצמך |  |
| 4:33.560 | כגולם |  |
| 4:33.561 | כמאמר |  |
| 4:33.562 | התנא |  |
| 4:33.563 | בפ | ' |
| 4:33.564 | בתרא |  |
| 4:33.565 | דאבות |  |
| 4:33.566 | ז | ' |
| 4:33.567 | דברים |  |
| 4:33.568 | בחכם |  |
| 4:33.569 | וחילופיהן |  |
| 4:33.570 | בגולם | , |
| 4:33.571 | ומכללן |  |
| 4:33.572 | החכם |  |
| 4:33.573 | שואל |  |
| 4:33.574 | כהלכה |  |
| 4:33.575 | ומשיב |  |
| 4:33.576 | כענין |  |
| 4:33.577 | ובגולם |  |
| 4:33.578 | בחילוף |  |
| 4:33.579 | זה | . |
| 4:33.580 | וכן |  |
| 4:33.581 | אתה |  |
| 4:33.582 | עשית | , |
| 4:33.583 | כי |  |
| 4:33.584 | השואל |  |
| 4:33.585 | שאל |  |
| 4:33.586 | אותך |  |
| 4:33.587 | כהלכה |  |
| 4:33.588 | את |  |
| 4:33.589 | מי |  |
| 4:33.590 | נעבוד |  |
| 4:33.591 | על |  |
| 4:33.592 | פי |  |
| 4:33.593 | הקבלה |  |
| 4:33.594 | החדשה |  |
| 4:33.595 | את |  |
| 4:33.596 | עתיק | , |
| 4:33.597 | או |  |
| 4:33.598 | אריך |  |
| 4:33.599 | אנפין | , |
| 4:33.600 | או |  |
| 4:33.601 | אבא |  |
| 4:33.602 | ואימא |  |
| 4:33.603 | וכו | '. |
| 4:33.604 | ובאה |  |
| 4:33.605 | תשובתך |  |
| 4:33.606 | שלא |  |
| 4:33.607 | כענין |  |
| 4:33.608 | כי |  |
| 4:33.609 | קבלתך |  |
| 4:33.610 | וכו | ' |
| 4:33.611 | כאילו |  |
| 4:33.612 | שאל |  |
| 4:33.613 | אותך |  |
| 4:33.614 | על |  |
| 4:33.615 | קבלתך | . |
| 4:33.616 | וגם |  |
| 4:33.617 | הרבית |  |
| 4:33.618 | דברים |  |
| 4:33.619 | והארכת |  |
| 4:33.620 | להביא |  |
| 4:33.621 | רמזים |  |
| 4:33.622 | וגימטריות |  |
| 4:33.623 | מדברי |  |
| 4:33.624 | חוות |  |
| 4:33.625 | יאיר |  |
| 4:33.626 | שלא |  |
| 4:33.627 | הבין |  |
| 4:33.628 | דברי |  |
| 4:33.629 | אביו | . |
| 4:33.630 | והזכיר |  |
| 4:33.631 | רמזים |  |
| 4:33.632 | שיוכל |  |
| 4:33.633 | כל |  |
| 4:33.634 | אדם |  |
| 4:33.635 | לעשות |  |
| 4:33.636 | כמותן |  |
| 4:33.637 | מבלי |  |
| 4:33.638 | קבלה |  |
| 4:33.639 | כמו |  |
| 4:33.640 | שעשה |  |
| 4:33.641 | ר | ' |
| 4:33.642 | שמואל |  |
| 4:33.643 | שחיבר |  |
| 4:33.644 | ספר |  |
| 4:33.645 | כח |  |
| 4:33.646 | ה | ' |
| 4:33.647 | שהוא |  |
| 4:33.648 | מלא |  |
| 4:33.649 | צירופים |  |
| 4:33.650 | וגימטריות |  |
| 4:33.651 | ורמזים |  |
| 4:33.652 | בלילה |  |
| 4:33.653 | אחת |  |
| 4:33.654 | מדעתו |  |
| 4:33.655 | מבלי |  |
| 4:33.656 | קבלה | , |
| 4:33.657 | רק |  |
| 4:33.658 | להראות |  |
| 4:33.659 | כחו |  |
| 4:33.660 | ויכלתו |  |
| 4:33.661 | לעשות |  |
| 4:33.662 | צירופים |  |
| 4:33.663 | וגימטריות |  |
| 4:33.664 | ורמזים |  |
| 4:33.665 | בנקל |  |
| 4:33.666 | ככתוב |  |
| 4:33.667 | בספר |  |
| 4:33.668 | א"נ | . |
| 4:33.669 | ובשלמא |  |
| 4:33.670 | רמזים |  |
| 4:33.671 | וצירופים |  |
| 4:33.672 | כל |  |
| 4:33.673 | שאינו |  |
| 4:33.674 | נוגע |  |
| 4:33.675 | בכבוד |  |
| 4:33.676 | המב"ה |  |
| 4:33.677 | לעבוד |  |
| 4:33.678 | אלהים |  |
| 4:33.679 | אחרים |  |
| 4:33.680 | ולזלזל |  |
| 4:33.681 | במעלת |  |
| 4:33.682 | המשנה |  |
| 4:33.683 | והתלמוד |  |
| 4:33.684 | ומדרשי |  |
| 4:33.685 | רז"ל |  |
| 4:33.686 | המספרים |  |
| 4:33.687 | שבחי |  |
| 4:33.688 | תורתינו |  |
| 4:33.689 | הקדושה |  |
| 4:33.690 | דליכא |  |
| 4:33.691 | מידי |  |
| 4:33.692 | דלא |  |
| 4:33.693 | רמיזה |  |
| 4:33.694 | באוריתא | . |
| 4:33.695 | ובפרט |  |
| 4:33.696 | זה |  |
| 4:33.697 | הארכת |  |
| 4:33.698 | לציין |  |
| 4:33.699 | הרמזים |  |
| 4:33.700 | שכ | ' |
| 4:33.701 | בעל |  |
| 4:33.702 | מנורת |  |
| 4:33.703 | המאור |  |
| 4:33.704 | למקומותם | , |
| 4:33.705 | ואת |  |
| 4:33.706 | המבוקש |  |
| 4:33.707 | ממך |  |
| 4:33.708 | לא |  |
| 4:33.709 | ענית | , |
| 4:33.710 | ולא |  |
| 4:33.711 | עלה |  |
| 4:33.712 | על |  |
| 4:33.713 | דל |  |
| 4:33.714 | שפתיך |  |
| 4:33.715 | כמדת |  |
| 4:33.716 | הגולם |  |
| 4:33.717 | שהזכיר |  |
| 4:33.718 | התנא |  |
| 4:33.719 | ע"ה |  |
| 4:33.720 | שאינו |  |
| 4:33.721 | יודע |  |
| 4:33.722 | להשיב |  |
| 4:33.723 | כענין |  |
| 4:33.724 | המבוקש |  |
| 4:33.725 | ממנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

