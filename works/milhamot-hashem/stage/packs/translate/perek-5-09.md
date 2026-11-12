# Translation pack: milhamot-hashem perek-5-09 (Milchamot ha-Shem 5:15–5:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-09`.

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
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- השם **הנכבד** → "the glorious Name".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **האלהי** before a name → "the divine" (note); never dropped.
- **רז"ל** → "our Rabbis of blessed memory".
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **הישמעאלים / ישמעאלים** → "the Ishmaelites" (note: Muslims).
- **בא"י** → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
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
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **בחינה** → "aspect".
- **עלה / עלול** → "cause / effect".
- אורות **דקים** → "subtle lights".
- **מין / מינות** → "heretic / heresy".
- כופר **בעיקר** → "denies the fundamental principle".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **תוה"ק** → "the holy Torah".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 5:15

Hebrew (Wikisource): מז) והנה שכח משנה שלמה ששנינו באבות (ד' כ"ב): "לידע ולהודיע ולהודע, שהוא אל, הוא הבורא, הוא היוצר, הוא המבין, הוא הדיין, הוא עד, הוא בעל דין, והוא עתיד לדון, ברוך הוא שאין לפניו לא עולה לא שכחה, לא משוא פנים ולא מקח שוחד", וארז"ל (ירושלמי שקלים פ"ה ה"א, מ"ח ע"ד): "כל האומר הקב"ה ותרן יותרון מעוהי, אלא מאריך רוחיה וגבי דיליה", ואמרו עוד (שם) "על וסביביו נשערה מאוד" שהקב"ה מדקדק עם חסידיו כחוט השערה, ואמרו עוד (שבועות ל"ט ע"א): "ונקה לא ינקה, מנקה הוא לשבים ואינו מנקה לשאינן שבים", וכן תרגם אונקלוס מפי ר' אליעזר ור' יהושע (שמות ל"ד ז'): "סלח לדיתיבין לאוריתיה, ולדלא תיבין לא מזכי" וכו'. ואל כל אלה האמונות הזרות לתורתינו הקדושה ירמזון דברי ר' תם ן' יחיא המובא בספר רביד הזהב שהבאתי לעיל סי' י"ג שכתב שהעוסקים בקבלה החדשה הם הורסים פינות התורה ועמודיה יתפלצון וכו'. וגם ירמזון אל מה שכתוב בכוונת האר"י בספר מחברת הקדש דף כ"ח ע"ב, ובספר חמדת ימים סוף פרק ז' מתקוני שבת שלא לאמר יגדל אלהים חי וישתבח וכו' מפני שסוברים כי שום עבודה ותפלה ושבח לא שייכי באלהינו ב"ה שהוא הסיבה הראשונה בהיותו מופשט מחומר וצורה, ונפתו בזה אחר עובדי האלילים שאמרו שהוא יתברך נעדר ומרומם ומסולק מכל ברכה ותהלה וכל תורה ועבודה ותפלה לא שייכי ביה. וגם בהתלבשו בספירות שבעולמות העליונים אין מספר למעלה מעולם האצילות (לפי דבריהם) שהם דקים וקלושים ואינם מושגים. וגם באדם קדמון דאצילות ועתיק יומין ואריך אנפין ואבא ואימא לא שייכי שום עבודה וקריאה ותפלה לרוב דקותם והעלמם. ורק בפרצוף אחרון דאצילות שהוא זעיר אנפין ונוקביה שכבר נתעבו קצת שייך בהו עבודה ותפלה. ודברים אלו מתבארים מדברי הרב בעל ספר הברית מאמר הנזכר לעיל והרב כסא אליהו והרח"ו בספר עץ חיים והרב בעל המצרף ונחלת יוסף וזולתם. וזה לשון ספר הברית: "ובזאת תדע אחי שאמונתינו אינה כאמונת הפילוסופים (האלהיים) והישמעאלים בענין היחוד והבורא. כי הם אין איתם יודע את השם הנכבד והנורא הוי'ה ב"ה. רק מאמינים במחוייב המציאות סתם בבחינתו הפשוטה כאשר בתחלה קודם הבריאה (ערום בלי לבוש) וזהו בחינתו הבלתי מושג אל זולתו בשום רמז ורמיזה. ובבחינה הזאת הטיבו אשר דיברו הפילוסופים (הכופרים) שלא שייך אליו שום עבודה ותפלה על כי הוא בבחינתו זאת נעלה ומרומם מכל ברכה ותהלה ובטלה כל תורה ומצוה. כמו שכתב ר' מאיר גבאי. והמתבונן בזה היטב יתבאר לו בלי שום ספק שאין מקום לכל זה בלתי אצילות הספירות. לא כן עם אלהי אברהם שמאמינים במחוייב המציאות מלובש במידותיו הוא סוד הוי"ה ב"ה שהוא הבחינה שאחר הבריאה, בחינה המושגת לבית יעקב שנגלה למשה בסנה והוציאנו מארץ מצרים ונתן תורת אמת לאבותינו דור המדבר פנים בפנים, שייך אליו העבודה והקרבנות והתפלה וכל המצוות המדובר בתוה"ק. ועליו הזהיר על (זעיר אנפין) משה ואמר ליראה את השם הנכבד והנורא הזה ה' את אלהיך" עכ"ל, וכן כתב מהרח"ו בעץ חיים שער טנת"א סוף פרק ב' שהאור היוצא מתוך אדם קדמון הוא זך מאוד, ועל ידי ירידתו והתרחקותו הוא מתעבה עיבוי אחר עיבוי, כיצד האור הנמשך ויוצא דרך האוזן הוא זך מאוד וכאשר נמשך האור הזה עד החוטם ויוצא קצת דרך שם הוא מתעבה וקונה איזה עביות וגסות מחמת הריחוק שנתרחק ונמשך יותר למטה בצאתו עד הפה ויוצא קצתו דרך שם מתעבה יותר בצאתו משם וכו' וכן כתב בשער הנקודים פרק א' שכל מה שהאור יורד למטה הוא ניכר ונרגש ומתגלה שם עכ"ל. וזהו טעמם במה שאמרו דלא שייך עבודה ותפלה רק לאלוה הנרגש ויש בו עובי וגסות קצת שתוכל לציירו בדמות תבניתו על ידי שנתעבה ונתגשם. אבל האין סוף או אדם קדמאה ואדם קדמון ואפילו עתיק ואריך שעדיין לא נתעבו ונתגשמו כל כך, לא שייך בהן עבודה, כי עדיין לא נתעבו ולא נתגשמו שתוכל לציירם, ורק העבודה והיראה לאלוה קצר האפים (זעיר אנפין) והוא לבדו לפי דעתם אלהינו על ידי שנתגשם ונגלה למשה בסנה ובסיני לאבותינו. והנה שכחו מקרא מלא כי מלאך הוא שנגלה למשה בסנה שנאמר (שמות ג' ב'): "וירא מלאך ה' אליו בלבת אש מתוך הסנה". ובמעמד הר סיני הלא נאמר (דברים ד' ט"ו): "כי לא ראיתם כל תמונה ביום דבר ה' עמכם", וכו' לא ראו דמות המושגת ברמ"ח איברים ושס"ה גידים כדבריהם. וקרוב לדבריהם כתב הרב עז לאלהים בבית קדש הקדשים דף כ"ו ע"ד וזה לשונו: "מי שמאמין שהאין סוף הוא עצם האלוהות כאילו כופר בעיקר. ועליהם אמר הנביא (ישעיהו כ"ט ט"ו): 'הוי המעמיקים מה' לסתיר עצה'. ועוד שעצם האלוהות הוא שם הוי"ה, וכל דבר שיש לו שם ואות יש בו שיעור קומה וכו'. ושם הדבר הוא גבולו שכן שם האדם לא שייך שיקרא בשם אדם אלא מראשו ועד רגליו, הרי שהשם הוא גבולו. לכן כל מה שתמצא משמות דשייכי באדם שייכים אליו יתברך, ואינו כסברת הטועים שאומרים שעצם האלוהות הוא אין סוף"! וחולק הוא על מהרח"ו ושאר מקובלים שסוברים שהתפארת שהוא זעיר אנפין הוא אלהינו, וסבירא ליה שהעבודה והתפלה למלכא קדישא דכל קדישין שהוא מאיר בלב זעיר אנפין כמו שכתב בפרק י"ט דף נ"ט ע"ד, וזה לשונו בדף כ"ו סוף ע"ד: "בר מן דין ובר מן דין אם איתא שהתפארת הוא עצם האלהות למה נעבוד את התפארת שהוא מקבל מאבא ואמא, נעבוד לאריך אנפין או לכתר שהוא יותר למעלה? ואם האין סוף הוא עצם האלוהות, איך אנו אומרים ברוך אתה ה', ממי יתברך ויביא לו שפע"? עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:15.1 | מז | ) |
| 5:15.2 | והנה |  |
| 5:15.3 | שכח |  |
| 5:15.4 | משנה |  |
| 5:15.5 | שלמה |  |
| 5:15.6 | ששנינו |  |
| 5:15.7 | באבות |  |
| 5:15.8 | ד | ' |
| 5:15.9 | כ"ב | ): |
| 5:15.10 | לידע |  |
| 5:15.11 | ולהודיע |  |
| 5:15.12 | ולהודע | , |
| 5:15.13 | שהוא |  |
| 5:15.14 | אל | , |
| 5:15.15 | הוא |  |
| 5:15.16 | הבורא | , |
| 5:15.17 | הוא |  |
| 5:15.18 | היוצר | , |
| 5:15.19 | הוא |  |
| 5:15.20 | המבין | , |
| 5:15.21 | הוא |  |
| 5:15.22 | הדיין | , |
| 5:15.23 | הוא |  |
| 5:15.24 | עד | , |
| 5:15.25 | הוא |  |
| 5:15.26 | בעל |  |
| 5:15.27 | דין | , |
| 5:15.28 | והוא |  |
| 5:15.29 | עתיד |  |
| 5:15.30 | לדון | , |
| 5:15.31 | ברוך |  |
| 5:15.32 | הוא |  |
| 5:15.33 | שאין |  |
| 5:15.34 | לפניו |  |
| 5:15.35 | לא |  |
| 5:15.36 | עולה |  |
| 5:15.37 | לא |  |
| 5:15.38 | שכחה | , |
| 5:15.39 | לא |  |
| 5:15.40 | משוא |  |
| 5:15.41 | פנים |  |
| 5:15.42 | ולא |  |
| 5:15.43 | מקח |  |
| 5:15.44 | שוחד | ", |
| 5:15.45 | וארז"ל |  |
| 5:15.46 | ירושלמי |  |
| 5:15.47 | שקלים |  |
| 5:15.48 | פ"ה |  |
| 5:15.49 | ה"א | , |
| 5:15.50 | מ"ח |  |
| 5:15.51 | ע"ד | ): |
| 5:15.52 | כל |  |
| 5:15.53 | האומר |  |
| 5:15.54 | הקב"ה |  |
| 5:15.55 | ותרן |  |
| 5:15.56 | יותרון |  |
| 5:15.57 | מעוהי | , |
| 5:15.58 | אלא |  |
| 5:15.59 | מאריך |  |
| 5:15.60 | רוחיה |  |
| 5:15.61 | וגבי |  |
| 5:15.62 | דיליה | ", |
| 5:15.63 | ואמרו |  |
| 5:15.64 | עוד |  |
| 5:15.65 | שם | ) |
| 5:15.66 | על |  |
| 5:15.67 | וסביביו |  |
| 5:15.68 | נשערה |  |
| 5:15.69 | מאוד | " |
| 5:15.70 | שהקב"ה |  |
| 5:15.71 | מדקדק |  |
| 5:15.72 | עם |  |
| 5:15.73 | חסידיו |  |
| 5:15.74 | כחוט |  |
| 5:15.75 | השערה | , |
| 5:15.76 | ואמרו |  |
| 5:15.77 | עוד |  |
| 5:15.78 | שבועות |  |
| 5:15.79 | ל"ט |  |
| 5:15.80 | ע"א | ): |
| 5:15.81 | ונקה |  |
| 5:15.82 | לא |  |
| 5:15.83 | ינקה | , |
| 5:15.84 | מנקה |  |
| 5:15.85 | הוא |  |
| 5:15.86 | לשבים |  |
| 5:15.87 | ואינו |  |
| 5:15.88 | מנקה |  |
| 5:15.89 | לשאינן |  |
| 5:15.90 | שבים | ", |
| 5:15.91 | וכן |  |
| 5:15.92 | תרגם |  |
| 5:15.93 | אונקלוס |  |
| 5:15.94 | מפי |  |
| 5:15.95 | ר | ' |
| 5:15.96 | אליעזר |  |
| 5:15.97 | ור | ' |
| 5:15.98 | יהושע |  |
| 5:15.99 | שמות |  |
| 5:15.100 | ל"ד |  |
| 5:15.101 | ז | '): |
| 5:15.102 | סלח |  |
| 5:15.103 | לדיתיבין |  |
| 5:15.104 | לאוריתיה | , |
| 5:15.105 | ולדלא |  |
| 5:15.106 | תיבין |  |
| 5:15.107 | לא |  |
| 5:15.108 | מזכי | " |
| 5:15.109 | וכו | '. |
| 5:15.110 | ואל |  |
| 5:15.111 | כל |  |
| 5:15.112 | אלה |  |
| 5:15.113 | האמונות |  |
| 5:15.114 | הזרות |  |
| 5:15.115 | לתורתינו |  |
| 5:15.116 | הקדושה |  |
| 5:15.117 | ירמזון |  |
| 5:15.118 | דברי |  |
| 5:15.119 | ר | ' |
| 5:15.120 | תם |  |
| 5:15.121 | ן | ' |
| 5:15.122 | יחיא |  |
| 5:15.123 | המובא |  |
| 5:15.124 | בספר |  |
| 5:15.125 | רביד |  |
| 5:15.126 | הזהב |  |
| 5:15.127 | שהבאתי |  |
| 5:15.128 | לעיל |  |
| 5:15.129 | סי | ' |
| 5:15.130 | י"ג |  |
| 5:15.131 | שכתב |  |
| 5:15.132 | שהעוסקים |  |
| 5:15.133 | בקבלה |  |
| 5:15.134 | החדשה |  |
| 5:15.135 | הם |  |
| 5:15.136 | הורסים |  |
| 5:15.137 | פינות |  |
| 5:15.138 | התורה |  |
| 5:15.139 | ועמודיה |  |
| 5:15.140 | יתפלצון |  |
| 5:15.141 | וכו | '. |
| 5:15.142 | וגם |  |
| 5:15.143 | ירמזון |  |
| 5:15.144 | אל |  |
| 5:15.145 | מה |  |
| 5:15.146 | שכתוב |  |
| 5:15.147 | בכוונת |  |
| 5:15.148 | האר"י |  |
| 5:15.149 | בספר |  |
| 5:15.150 | מחברת |  |
| 5:15.151 | הקדש |  |
| 5:15.152 | דף |  |
| 5:15.153 | כ"ח |  |
| 5:15.154 | ע"ב | , |
| 5:15.155 | ובספר |  |
| 5:15.156 | חמדת |  |
| 5:15.157 | ימים |  |
| 5:15.158 | סוף |  |
| 5:15.159 | פרק |  |
| 5:15.160 | ז | ' |
| 5:15.161 | מתקוני |  |
| 5:15.162 | שבת |  |
| 5:15.163 | שלא |  |
| 5:15.164 | לאמר |  |
| 5:15.165 | יגדל |  |
| 5:15.166 | אלהים |  |
| 5:15.167 | חי |  |
| 5:15.168 | וישתבח |  |
| 5:15.169 | וכו | ' |
| 5:15.170 | מפני |  |
| 5:15.171 | שסוברים |  |
| 5:15.172 | כי |  |
| 5:15.173 | שום |  |
| 5:15.174 | עבודה |  |
| 5:15.175 | ותפלה |  |
| 5:15.176 | ושבח |  |
| 5:15.177 | לא |  |
| 5:15.178 | שייכי |  |
| 5:15.179 | באלהינו |  |
| 5:15.180 | ב"ה |  |
| 5:15.181 | שהוא |  |
| 5:15.182 | הסיבה |  |
| 5:15.183 | הראשונה |  |
| 5:15.184 | בהיותו |  |
| 5:15.185 | מופשט |  |
| 5:15.186 | מחומר |  |
| 5:15.187 | וצורה | , |
| 5:15.188 | ונפתו |  |
| 5:15.189 | בזה |  |
| 5:15.190 | אחר |  |
| 5:15.191 | עובדי |  |
| 5:15.192 | האלילים |  |
| 5:15.193 | שאמרו |  |
| 5:15.194 | שהוא |  |
| 5:15.195 | יתברך |  |
| 5:15.196 | נעדר |  |
| 5:15.197 | ומרומם |  |
| 5:15.198 | ומסולק |  |
| 5:15.199 | מכל |  |
| 5:15.200 | ברכה |  |
| 5:15.201 | ותהלה |  |
| 5:15.202 | וכל |  |
| 5:15.203 | תורה |  |
| 5:15.204 | ועבודה |  |
| 5:15.205 | ותפלה |  |
| 5:15.206 | לא |  |
| 5:15.207 | שייכי |  |
| 5:15.208 | ביה | . |
| 5:15.209 | וגם |  |
| 5:15.210 | בהתלבשו |  |
| 5:15.211 | בספירות |  |
| 5:15.212 | שבעולמות |  |
| 5:15.213 | העליונים |  |
| 5:15.214 | אין |  |
| 5:15.215 | מספר |  |
| 5:15.216 | למעלה |  |
| 5:15.217 | מעולם |  |
| 5:15.218 | האצילות |  |
| 5:15.219 | לפי |  |
| 5:15.220 | דבריהם | ) |
| 5:15.221 | שהם |  |
| 5:15.222 | דקים |  |
| 5:15.223 | וקלושים |  |
| 5:15.224 | ואינם |  |
| 5:15.225 | מושגים | . |
| 5:15.226 | וגם |  |
| 5:15.227 | באדם |  |
| 5:15.228 | קדמון |  |
| 5:15.229 | דאצילות |  |
| 5:15.230 | ועתיק |  |
| 5:15.231 | יומין |  |
| 5:15.232 | ואריך |  |
| 5:15.233 | אנפין |  |
| 5:15.234 | ואבא |  |
| 5:15.235 | ואימא |  |
| 5:15.236 | לא |  |
| 5:15.237 | שייכי |  |
| 5:15.238 | שום |  |
| 5:15.239 | עבודה |  |
| 5:15.240 | וקריאה |  |
| 5:15.241 | ותפלה |  |
| 5:15.242 | לרוב |  |
| 5:15.243 | דקותם |  |
| 5:15.244 | והעלמם | . |
| 5:15.245 | ורק |  |
| 5:15.246 | בפרצוף |  |
| 5:15.247 | אחרון |  |
| 5:15.248 | דאצילות |  |
| 5:15.249 | שהוא |  |
| 5:15.250 | זעיר |  |
| 5:15.251 | אנפין |  |
| 5:15.252 | ונוקביה |  |
| 5:15.253 | שכבר |  |
| 5:15.254 | נתעבו |  |
| 5:15.255 | קצת |  |
| 5:15.256 | שייך |  |
| 5:15.257 | בהו |  |
| 5:15.258 | עבודה |  |
| 5:15.259 | ותפלה | . |
| 5:15.260 | ודברים |  |
| 5:15.261 | אלו |  |
| 5:15.262 | מתבארים |  |
| 5:15.263 | מדברי |  |
| 5:15.264 | הרב |  |
| 5:15.265 | בעל |  |
| 5:15.266 | ספר |  |
| 5:15.267 | הברית |  |
| 5:15.268 | מאמר |  |
| 5:15.269 | הנזכר |  |
| 5:15.270 | לעיל |  |
| 5:15.271 | והרב |  |
| 5:15.272 | כסא |  |
| 5:15.273 | אליהו |  |
| 5:15.274 | והרח"ו |  |
| 5:15.275 | בספר |  |
| 5:15.276 | עץ |  |
| 5:15.277 | חיים |  |
| 5:15.278 | והרב |  |
| 5:15.279 | בעל |  |
| 5:15.280 | המצרף |  |
| 5:15.281 | ונחלת |  |
| 5:15.282 | יוסף |  |
| 5:15.283 | וזולתם | . |
| 5:15.284 | וזה |  |
| 5:15.285 | לשון |  |
| 5:15.286 | ספר |  |
| 5:15.287 | הברית | : |
| 5:15.288 | ובזאת |  |
| 5:15.289 | תדע |  |
| 5:15.290 | אחי |  |
| 5:15.291 | שאמונתינו |  |
| 5:15.292 | אינה |  |
| 5:15.293 | כאמונת |  |
| 5:15.294 | הפילוסופים |  |
| 5:15.295 | האלהיים | ) |
| 5:15.296 | והישמעאלים |  |
| 5:15.297 | בענין |  |
| 5:15.298 | היחוד |  |
| 5:15.299 | והבורא | . |
| 5:15.300 | כי |  |
| 5:15.301 | הם |  |
| 5:15.302 | אין |  |
| 5:15.303 | איתם |  |
| 5:15.304 | יודע |  |
| 5:15.305 | את |  |
| 5:15.306 | השם |  |
| 5:15.307 | הנכבד |  |
| 5:15.308 | והנורא |  |
| 5:15.309 | הוי'ה |  |
| 5:15.310 | ב"ה | . |
| 5:15.311 | רק |  |
| 5:15.312 | מאמינים |  |
| 5:15.313 | במחוייב |  |
| 5:15.314 | המציאות |  |
| 5:15.315 | סתם |  |
| 5:15.316 | בבחינתו |  |
| 5:15.317 | הפשוטה |  |
| 5:15.318 | כאשר |  |
| 5:15.319 | בתחלה |  |
| 5:15.320 | קודם |  |
| 5:15.321 | הבריאה |  |
| 5:15.322 | ערום |  |
| 5:15.323 | בלי |  |
| 5:15.324 | לבוש | ) |
| 5:15.325 | וזהו |  |
| 5:15.326 | בחינתו |  |
| 5:15.327 | הבלתי |  |
| 5:15.328 | מושג |  |
| 5:15.329 | אל |  |
| 5:15.330 | זולתו |  |
| 5:15.331 | בשום |  |
| 5:15.332 | רמז |  |
| 5:15.333 | ורמיזה | . |
| 5:15.334 | ובבחינה |  |
| 5:15.335 | הזאת |  |
| 5:15.336 | הטיבו |  |
| 5:15.337 | אשר |  |
| 5:15.338 | דיברו |  |
| 5:15.339 | הפילוסופים |  |
| 5:15.340 | הכופרים | ) |
| 5:15.341 | שלא |  |
| 5:15.342 | שייך |  |
| 5:15.343 | אליו |  |
| 5:15.344 | שום |  |
| 5:15.345 | עבודה |  |
| 5:15.346 | ותפלה |  |
| 5:15.347 | על |  |
| 5:15.348 | כי |  |
| 5:15.349 | הוא |  |
| 5:15.350 | בבחינתו |  |
| 5:15.351 | זאת |  |
| 5:15.352 | נעלה |  |
| 5:15.353 | ומרומם |  |
| 5:15.354 | מכל |  |
| 5:15.355 | ברכה |  |
| 5:15.356 | ותהלה |  |
| 5:15.357 | ובטלה |  |
| 5:15.358 | כל |  |
| 5:15.359 | תורה |  |
| 5:15.360 | ומצוה | . |
| 5:15.361 | כמו |  |
| 5:15.362 | שכתב |  |
| 5:15.363 | ר | ' |
| 5:15.364 | מאיר |  |
| 5:15.365 | גבאי | . |
| 5:15.366 | והמתבונן |  |
| 5:15.367 | בזה |  |
| 5:15.368 | היטב |  |
| 5:15.369 | יתבאר |  |
| 5:15.370 | לו |  |
| 5:15.371 | בלי |  |
| 5:15.372 | שום |  |
| 5:15.373 | ספק |  |
| 5:15.374 | שאין |  |
| 5:15.375 | מקום |  |
| 5:15.376 | לכל |  |
| 5:15.377 | זה |  |
| 5:15.378 | בלתי |  |
| 5:15.379 | אצילות |  |
| 5:15.380 | הספירות | . |
| 5:15.381 | לא |  |
| 5:15.382 | כן |  |
| 5:15.383 | עם |  |
| 5:15.384 | אלהי |  |
| 5:15.385 | אברהם |  |
| 5:15.386 | שמאמינים |  |
| 5:15.387 | במחוייב |  |
| 5:15.388 | המציאות |  |
| 5:15.389 | מלובש |  |
| 5:15.390 | במידותיו |  |
| 5:15.391 | הוא |  |
| 5:15.392 | סוד |  |
| 5:15.393 | הוי"ה |  |
| 5:15.394 | ב"ה |  |
| 5:15.395 | שהוא |  |
| 5:15.396 | הבחינה |  |
| 5:15.397 | שאחר |  |
| 5:15.398 | הבריאה | , |
| 5:15.399 | בחינה |  |
| 5:15.400 | המושגת |  |
| 5:15.401 | לבית |  |
| 5:15.402 | יעקב |  |
| 5:15.403 | שנגלה |  |
| 5:15.404 | למשה |  |
| 5:15.405 | בסנה |  |
| 5:15.406 | והוציאנו |  |
| 5:15.407 | מארץ |  |
| 5:15.408 | מצרים |  |
| 5:15.409 | ונתן |  |
| 5:15.410 | תורת |  |
| 5:15.411 | אמת |  |
| 5:15.412 | לאבותינו |  |
| 5:15.413 | דור |  |
| 5:15.414 | המדבר |  |
| 5:15.415 | פנים |  |
| 5:15.416 | בפנים | , |
| 5:15.417 | שייך |  |
| 5:15.418 | אליו |  |
| 5:15.419 | העבודה |  |
| 5:15.420 | והקרבנות |  |
| 5:15.421 | והתפלה |  |
| 5:15.422 | וכל |  |
| 5:15.423 | המצוות |  |
| 5:15.424 | המדובר |  |
| 5:15.425 | בתוה"ק | . |
| 5:15.426 | ועליו |  |
| 5:15.427 | הזהיר |  |
| 5:15.428 | על |  |
| 5:15.429 | זעיר |  |
| 5:15.430 | אנפין | ) |
| 5:15.431 | משה |  |
| 5:15.432 | ואמר |  |
| 5:15.433 | ליראה |  |
| 5:15.434 | את |  |
| 5:15.435 | השם |  |
| 5:15.436 | הנכבד |  |
| 5:15.437 | והנורא |  |
| 5:15.438 | הזה |  |
| 5:15.439 | ה | ' |
| 5:15.440 | את |  |
| 5:15.441 | אלהיך | " |
| 5:15.442 | עכ"ל | , |
| 5:15.443 | וכן |  |
| 5:15.444 | כתב |  |
| 5:15.445 | מהרח"ו |  |
| 5:15.446 | בעץ |  |
| 5:15.447 | חיים |  |
| 5:15.448 | שער |  |
| 5:15.449 | טנת"א |  |
| 5:15.450 | סוף |  |
| 5:15.451 | פרק |  |
| 5:15.452 | ב | ' |
| 5:15.453 | שהאור |  |
| 5:15.454 | היוצא |  |
| 5:15.455 | מתוך |  |
| 5:15.456 | אדם |  |
| 5:15.457 | קדמון |  |
| 5:15.458 | הוא |  |
| 5:15.459 | זך |  |
| 5:15.460 | מאוד | , |
| 5:15.461 | ועל |  |
| 5:15.462 | ידי |  |
| 5:15.463 | ירידתו |  |
| 5:15.464 | והתרחקותו |  |
| 5:15.465 | הוא |  |
| 5:15.466 | מתעבה |  |
| 5:15.467 | עיבוי |  |
| 5:15.468 | אחר |  |
| 5:15.469 | עיבוי | , |
| 5:15.470 | כיצד |  |
| 5:15.471 | האור |  |
| 5:15.472 | הנמשך |  |
| 5:15.473 | ויוצא |  |
| 5:15.474 | דרך |  |
| 5:15.475 | האוזן |  |
| 5:15.476 | הוא |  |
| 5:15.477 | זך |  |
| 5:15.478 | מאוד |  |
| 5:15.479 | וכאשר |  |
| 5:15.480 | נמשך |  |
| 5:15.481 | האור |  |
| 5:15.482 | הזה |  |
| 5:15.483 | עד |  |
| 5:15.484 | החוטם |  |
| 5:15.485 | ויוצא |  |
| 5:15.486 | קצת |  |
| 5:15.487 | דרך |  |
| 5:15.488 | שם |  |
| 5:15.489 | הוא |  |
| 5:15.490 | מתעבה |  |
| 5:15.491 | וקונה |  |
| 5:15.492 | איזה |  |
| 5:15.493 | עביות |  |
| 5:15.494 | וגסות |  |
| 5:15.495 | מחמת |  |
| 5:15.496 | הריחוק |  |
| 5:15.497 | שנתרחק |  |
| 5:15.498 | ונמשך |  |
| 5:15.499 | יותר |  |
| 5:15.500 | למטה |  |
| 5:15.501 | בצאתו |  |
| 5:15.502 | עד |  |
| 5:15.503 | הפה |  |
| 5:15.504 | ויוצא |  |
| 5:15.505 | קצתו |  |
| 5:15.506 | דרך |  |
| 5:15.507 | שם |  |
| 5:15.508 | מתעבה |  |
| 5:15.509 | יותר |  |
| 5:15.510 | בצאתו |  |
| 5:15.511 | משם |  |
| 5:15.512 | וכו | ' |
| 5:15.513 | וכן |  |
| 5:15.514 | כתב |  |
| 5:15.515 | בשער |  |
| 5:15.516 | הנקודים |  |
| 5:15.517 | פרק |  |
| 5:15.518 | א | ' |
| 5:15.519 | שכל |  |
| 5:15.520 | מה |  |
| 5:15.521 | שהאור |  |
| 5:15.522 | יורד |  |
| 5:15.523 | למטה |  |
| 5:15.524 | הוא |  |
| 5:15.525 | ניכר |  |
| 5:15.526 | ונרגש |  |
| 5:15.527 | ומתגלה |  |
| 5:15.528 | שם |  |
| 5:15.529 | עכ"ל | . |
| 5:15.530 | וזהו |  |
| 5:15.531 | טעמם |  |
| 5:15.532 | במה |  |
| 5:15.533 | שאמרו |  |
| 5:15.534 | דלא |  |
| 5:15.535 | שייך |  |
| 5:15.536 | עבודה |  |
| 5:15.537 | ותפלה |  |
| 5:15.538 | רק |  |
| 5:15.539 | לאלוה |  |
| 5:15.540 | הנרגש |  |
| 5:15.541 | ויש |  |
| 5:15.542 | בו |  |
| 5:15.543 | עובי |  |
| 5:15.544 | וגסות |  |
| 5:15.545 | קצת |  |
| 5:15.546 | שתוכל |  |
| 5:15.547 | לציירו |  |
| 5:15.548 | בדמות |  |
| 5:15.549 | תבניתו |  |
| 5:15.550 | על |  |
| 5:15.551 | ידי |  |
| 5:15.552 | שנתעבה |  |
| 5:15.553 | ונתגשם | . |
| 5:15.554 | אבל |  |
| 5:15.555 | האין |  |
| 5:15.556 | סוף |  |
| 5:15.557 | או |  |
| 5:15.558 | אדם |  |
| 5:15.559 | קדמאה |  |
| 5:15.560 | ואדם |  |
| 5:15.561 | קדמון |  |
| 5:15.562 | ואפילו |  |
| 5:15.563 | עתיק |  |
| 5:15.564 | ואריך |  |
| 5:15.565 | שעדיין |  |
| 5:15.566 | לא |  |
| 5:15.567 | נתעבו |  |
| 5:15.568 | ונתגשמו |  |
| 5:15.569 | כל |  |
| 5:15.570 | כך | , |
| 5:15.571 | לא |  |
| 5:15.572 | שייך |  |
| 5:15.573 | בהן |  |
| 5:15.574 | עבודה | , |
| 5:15.575 | כי |  |
| 5:15.576 | עדיין |  |
| 5:15.577 | לא |  |
| 5:15.578 | נתעבו |  |
| 5:15.579 | ולא |  |
| 5:15.580 | נתגשמו |  |
| 5:15.581 | שתוכל |  |
| 5:15.582 | לציירם | , |
| 5:15.583 | ורק |  |
| 5:15.584 | העבודה |  |
| 5:15.585 | והיראה |  |
| 5:15.586 | לאלוה |  |
| 5:15.587 | קצר |  |
| 5:15.588 | האפים |  |
| 5:15.589 | זעיר |  |
| 5:15.590 | אנפין | ) |
| 5:15.591 | והוא |  |
| 5:15.592 | לבדו |  |
| 5:15.593 | לפי |  |
| 5:15.594 | דעתם |  |
| 5:15.595 | אלהינו |  |
| 5:15.596 | על |  |
| 5:15.597 | ידי |  |
| 5:15.598 | שנתגשם |  |
| 5:15.599 | ונגלה |  |
| 5:15.600 | למשה |  |
| 5:15.601 | בסנה |  |
| 5:15.602 | ובסיני |  |
| 5:15.603 | לאבותינו | . |
| 5:15.604 | והנה |  |
| 5:15.605 | שכחו |  |
| 5:15.606 | מקרא |  |
| 5:15.607 | מלא |  |
| 5:15.608 | כי |  |
| 5:15.609 | מלאך |  |
| 5:15.610 | הוא |  |
| 5:15.611 | שנגלה |  |
| 5:15.612 | למשה |  |
| 5:15.613 | בסנה |  |
| 5:15.614 | שנאמר |  |
| 5:15.615 | שמות |  |
| 5:15.616 | ג | ' |
| 5:15.617 | ב | '): |
| 5:15.618 | וירא |  |
| 5:15.619 | מלאך |  |
| 5:15.620 | ה | ' |
| 5:15.621 | אליו |  |
| 5:15.622 | בלבת |  |
| 5:15.623 | אש |  |
| 5:15.624 | מתוך |  |
| 5:15.625 | הסנה | ". |
| 5:15.626 | ובמעמד |  |
| 5:15.627 | הר |  |
| 5:15.628 | סיני |  |
| 5:15.629 | הלא |  |
| 5:15.630 | נאמר |  |
| 5:15.631 | דברים |  |
| 5:15.632 | ד | ' |
| 5:15.633 | ט"ו | ): |
| 5:15.634 | כי |  |
| 5:15.635 | לא |  |
| 5:15.636 | ראיתם |  |
| 5:15.637 | כל |  |
| 5:15.638 | תמונה |  |
| 5:15.639 | ביום |  |
| 5:15.640 | דבר |  |
| 5:15.641 | ה | ' |
| 5:15.642 | עמכם | ", |
| 5:15.643 | וכו | ' |
| 5:15.644 | לא |  |
| 5:15.645 | ראו |  |
| 5:15.646 | דמות |  |
| 5:15.647 | המושגת |  |
| 5:15.648 | ברמ"ח |  |
| 5:15.649 | איברים |  |
| 5:15.650 | ושס"ה |  |
| 5:15.651 | גידים |  |
| 5:15.652 | כדבריהם | . |
| 5:15.653 | וקרוב |  |
| 5:15.654 | לדבריהם |  |
| 5:15.655 | כתב |  |
| 5:15.656 | הרב |  |
| 5:15.657 | עז |  |
| 5:15.658 | לאלהים |  |
| 5:15.659 | בבית |  |
| 5:15.660 | קדש |  |
| 5:15.661 | הקדשים |  |
| 5:15.662 | דף |  |
| 5:15.663 | כ"ו |  |
| 5:15.664 | ע"ד |  |
| 5:15.665 | וזה |  |
| 5:15.666 | לשונו | : |
| 5:15.667 | מי |  |
| 5:15.668 | שמאמין |  |
| 5:15.669 | שהאין |  |
| 5:15.670 | סוף |  |
| 5:15.671 | הוא |  |
| 5:15.672 | עצם |  |
| 5:15.673 | האלוהות |  |
| 5:15.674 | כאילו |  |
| 5:15.675 | כופר |  |
| 5:15.676 | בעיקר | . |
| 5:15.677 | ועליהם |  |
| 5:15.678 | אמר |  |
| 5:15.679 | הנביא |  |
| 5:15.680 | ישעיהו |  |
| 5:15.681 | כ"ט |  |
| 5:15.682 | ט"ו | ): |
| 5:15.683 | הוי |  |
| 5:15.684 | המעמיקים |  |
| 5:15.685 | מה | ' |
| 5:15.686 | לסתיר |  |
| 5:15.687 | עצה | '. |
| 5:15.688 | ועוד |  |
| 5:15.689 | שעצם |  |
| 5:15.690 | האלוהות |  |
| 5:15.691 | הוא |  |
| 5:15.692 | שם |  |
| 5:15.693 | הוי"ה | , |
| 5:15.694 | וכל |  |
| 5:15.695 | דבר |  |
| 5:15.696 | שיש |  |
| 5:15.697 | לו |  |
| 5:15.698 | שם |  |
| 5:15.699 | ואות |  |
| 5:15.700 | יש |  |
| 5:15.701 | בו |  |
| 5:15.702 | שיעור |  |
| 5:15.703 | קומה |  |
| 5:15.704 | וכו | '. |
| 5:15.705 | ושם |  |
| 5:15.706 | הדבר |  |
| 5:15.707 | הוא |  |
| 5:15.708 | גבולו |  |
| 5:15.709 | שכן |  |
| 5:15.710 | שם |  |
| 5:15.711 | האדם |  |
| 5:15.712 | לא |  |
| 5:15.713 | שייך |  |
| 5:15.714 | שיקרא |  |
| 5:15.715 | בשם |  |
| 5:15.716 | אדם |  |
| 5:15.717 | אלא |  |
| 5:15.718 | מראשו |  |
| 5:15.719 | ועד |  |
| 5:15.720 | רגליו | , |
| 5:15.721 | הרי |  |
| 5:15.722 | שהשם |  |
| 5:15.723 | הוא |  |
| 5:15.724 | גבולו | . |
| 5:15.725 | לכן |  |
| 5:15.726 | כל |  |
| 5:15.727 | מה |  |
| 5:15.728 | שתמצא |  |
| 5:15.729 | משמות |  |
| 5:15.730 | דשייכי |  |
| 5:15.731 | באדם |  |
| 5:15.732 | שייכים |  |
| 5:15.733 | אליו |  |
| 5:15.734 | יתברך | , |
| 5:15.735 | ואינו |  |
| 5:15.736 | כסברת |  |
| 5:15.737 | הטועים |  |
| 5:15.738 | שאומרים |  |
| 5:15.739 | שעצם |  |
| 5:15.740 | האלוהות |  |
| 5:15.741 | הוא |  |
| 5:15.742 | אין |  |
| 5:15.743 | סוף | "! |
| 5:15.744 | וחולק |  |
| 5:15.745 | הוא |  |
| 5:15.746 | על |  |
| 5:15.747 | מהרח"ו |  |
| 5:15.748 | ושאר |  |
| 5:15.749 | מקובלים |  |
| 5:15.750 | שסוברים |  |
| 5:15.751 | שהתפארת |  |
| 5:15.752 | שהוא |  |
| 5:15.753 | זעיר |  |
| 5:15.754 | אנפין |  |
| 5:15.755 | הוא |  |
| 5:15.756 | אלהינו | , |
| 5:15.757 | וסבירא |  |
| 5:15.758 | ליה |  |
| 5:15.759 | שהעבודה |  |
| 5:15.760 | והתפלה |  |
| 5:15.761 | למלכא |  |
| 5:15.762 | קדישא |  |
| 5:15.763 | דכל |  |
| 5:15.764 | קדישין |  |
| 5:15.765 | שהוא |  |
| 5:15.766 | מאיר |  |
| 5:15.767 | בלב |  |
| 5:15.768 | זעיר |  |
| 5:15.769 | אנפין |  |
| 5:15.770 | כמו |  |
| 5:15.771 | שכתב |  |
| 5:15.772 | בפרק |  |
| 5:15.773 | י"ט |  |
| 5:15.774 | דף |  |
| 5:15.775 | נ"ט |  |
| 5:15.776 | ע"ד | , |
| 5:15.777 | וזה |  |
| 5:15.778 | לשונו |  |
| 5:15.779 | בדף |  |
| 5:15.780 | כ"ו |  |
| 5:15.781 | סוף |  |
| 5:15.782 | ע"ד | : |
| 5:15.783 | בר |  |
| 5:15.784 | מן |  |
| 5:15.785 | דין |  |
| 5:15.786 | ובר |  |
| 5:15.787 | מן |  |
| 5:15.788 | דין |  |
| 5:15.789 | אם |  |
| 5:15.790 | איתא |  |
| 5:15.791 | שהתפארת |  |
| 5:15.792 | הוא |  |
| 5:15.793 | עצם |  |
| 5:15.794 | האלהות |  |
| 5:15.795 | למה |  |
| 5:15.796 | נעבוד |  |
| 5:15.797 | את |  |
| 5:15.798 | התפארת |  |
| 5:15.799 | שהוא |  |
| 5:15.800 | מקבל |  |
| 5:15.801 | מאבא |  |
| 5:15.802 | ואמא | , |
| 5:15.803 | נעבוד |  |
| 5:15.804 | לאריך |  |
| 5:15.805 | אנפין |  |
| 5:15.806 | או |  |
| 5:15.807 | לכתר |  |
| 5:15.808 | שהוא |  |
| 5:15.809 | יותר |  |
| 5:15.810 | למעלה | ? |
| 5:15.811 | ואם |  |
| 5:15.812 | האין |  |
| 5:15.813 | סוף |  |
| 5:15.814 | הוא |  |
| 5:15.815 | עצם |  |
| 5:15.816 | האלוהות | , |
| 5:15.817 | איך |  |
| 5:15.818 | אנו |  |
| 5:15.819 | אומרים |  |
| 5:15.820 | ברוך |  |
| 5:15.821 | אתה |  |
| 5:15.822 | ה | ', |
| 5:15.823 | ממי |  |
| 5:15.824 | יתברך |  |
| 5:15.825 | ויביא |  |
| 5:15.826 | לו |  |
| 5:15.827 | שפע | "? |
| 5:15.828 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

