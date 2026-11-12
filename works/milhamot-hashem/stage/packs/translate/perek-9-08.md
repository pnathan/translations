# Translation pack: milhamot-hashem perek-9-08 (Milchamot ha-Shem 9:10–9:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-08`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- בעל **העיקרים** → "the author of the Ikkarim".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **הנוצרים / הנצרים** → "the Christians".
- מאמיני **השלוש** → "the believers in the Trinity".
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **עכ"ל** → "Thus far his words." (outside the marks).
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **אבותיהם** for "אבתיכם": follow his wording; the verse in `apparatus`.
- כעיר **פרוצה** אין חומה → "like a city breached, without a wall".
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **אללה / הדה / אלגהת** (any Judeo-Arabic passage: Saadia on Job; Choter ben
  Shelomo; a Muslim polemic; the proverb in 2:13): from the Arabic as it
  stands, never transliterated in the English. Ambiguous spelling: the evident
  word (note), the other in `apparatus`; אלי' (alif maqsura) is ilā "to"; a
  scribal letter change (ר for ד) noted, not corrected; an unexplained word
  (שברא): "[unclear]", the letters in `apparatus`. A rendering from Saadia's
  translation (אלבלד שברא): as given, Saadia's translation named in
  `apparatus`; his אמנתא / אמונתנו pairs: the Arabic as written, his Hebrew
  gloss in a note. Glosses inside the Arabic ("אתתלית (תלת קשרי דמהימנותא)")
  stay in parentheses, translated as his. (Q4.)
- **אצהאר / גילוי**: ספר אצהאר אלחק = גילוי האמת; the title as each place writes
  it ("Iẓhar al-Chaqq", "Gilui ha-Emet"), a note that they are one book; its
  author unnamed.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 9:10

Hebrew (Wikisource): קכז) לבי לחוקקי ישראל המתנבים בעם להבין ולהשכיל בתורת ה' בחכמה ותבונה. ובדעת ישרה על פי קבלת רז"ל בעלי המשנה והתלמוד והפוסקים הדולים ומשקים לכל צמא ממקור מים חיים, משה קבל תורה מסיני ומסרה ליהושע וכו' והם מלאו אחר ה' ותורתו הקדושה שלא לפנות אל האלילים, איך בכל מצותיה של תורה האריכו הרחיבו לבאר חילוקי דיניהם בכמה פרטים, וחומרות יתירות ממה שנזכר בתלמוד! ובדיני יחוד השי"ת שהוא העיקר הגדול, העלימו עיניהם ועזבוהו ולא דברו מאומה אפי' בפרט אחד מכל דיני יחודו ית'! ולמה זה ועל מה זה הניחוהו כעיר פרוצה אין חומה, והוא נוהג בכל זמן ובכל מקום בין בארץ בין בחוצא לארץ, בין בפני הבית בין שלא בפני הבית, כאלו נשתכחה מישראל תורת יחוד השי"ת היוצאת מדברי התלמוד בבלי וירושלמי ומדרשי רז"ל, אחרי מות ר' בחיי בעל חובת הלבבות ורבי יהודה הלוי והרס"ג והרמב"ם והסמ"ג ובעל העיקרים וזולתם! עד שבעבור זה נפלו קצת אחרונים בפח הנסיון אשר נסה ה' אלהינו אותנו על ידי המלך העכו"ם בשלחו ספר הזהר לחכמי טוליטולא, והמה שמחו בו שמחה גדולה, כי חפצו להוציא ישן תורת המשנה והתלמוד, מפני חדש תורת הזהר המבזה את המשנה והתלמוד כמבואר לעיל, ולא נתנו אל לבם למנוע מלעבוד אלהים חדשים, מקרבו באו, אשר הזהירנו ה' אלהינו לבלתי עבוד זולתו, אפי' נקראהו בשמו כמעשה דורו של אנוש. וחכמי טוליטולא לא עשו כן אדרבא שלחו מתנות רבות למלך העכו"ם בעבורו והגדילו והחשיבו דברי הפילסוף מחבר הזהר המתנבא בשקר בשם רשב"י שנגלה לו עתיק, יותר ממ"ח נביאים ושבע נביאות שנתנבאו להם לישראל, שמ"ח נביאים ושבע נביאות לא פיחתו ולא הותירו על מה שכתוב בתורה אפי' אות אחת חוץ ממקרא מגלה, ומחבר הזהר הוסיף אמונת אלהים אחרים חדשים אשר לא שערום אבותיהם. ואמר שמי שלא יעסוק בתורתו החדשה לית ליה רשו למיעל קמי מריה. ובספר אצהאר אלחק לאחד מחכמי הערביים בטענתו נגד אמונת השלוש שהיא אמונת הנצרים כתב בדף רפ"א ורפ"ב וז"ל. ועלמא אליהוד מן ודן מוסי' אלי' זמאננא הדא לא יעתרפון אתתלית (תלת קשרי דמהימנותא) וצאהר אן דאת אללה וצפאתה אלכמאליה קדימא גיה מתגיירה אזלא' ואבדא'. פלו כאן אלתתלית חקא', לכאן אלואגב עלי' מוסי' עליה אסלאם ואנביא בני אסראיל אן יביינוה חק אלביאן, פאלעגב כל אלעגב אן תכון אלשריעה אלמוסויה אלתי כאנת ואנבת אלאטאעה בגמיע אלאנביא אלי' עהד עיסי' כאלה עם ביאן הדה אלעקידה אלתי הי מדאר אלנגאה (היא מרכז ההצלה והוא מאמר הזהר דלית ליה רשו למיעל קמי מריה) עלי' זעם אהל אלתתלית ולא ימכן נגאת אחדכדונהא נבייא כאן או גיר נבי, ולא יביין מוסי' ע"ס. ולא נבי מן אלאנביא אלאסראלייה הדה אלעקידה בביאן ואצח בחית תפהם מנה הדה אלעקידה צראחה, ולא יבקש פיהא שך ויביין מוסי' אלאחכאם אלתי הי ענד מקרס אהל אלתתלית צעיפה נאקצה גדא באלתשריח אלאתם. ויכררהא מרה בעד אולי וכרה בעד אכרא, ויאכד ען מחאפצהתא תאכידא בלינא, ויוגב אלקתל למן לם יעתקדהא עכ"ל. (העתק ועיין בסוף הספר) וטענתו זאת בעצמה אמיתית צודקת על הקבלה החדשה, ומדבריו מוכח שלא פשטה עדיין אמונת הקבלה החדשה בישראל.

| id | word | punctuation after |
|---|---|---|
| 9:10.1 | קכז | ) |
| 9:10.2 | לבי |  |
| 9:10.3 | לחוקקי |  |
| 9:10.4 | ישראל |  |
| 9:10.5 | המתנבים |  |
| 9:10.6 | בעם |  |
| 9:10.7 | להבין |  |
| 9:10.8 | ולהשכיל |  |
| 9:10.9 | בתורת |  |
| 9:10.10 | ה | ' |
| 9:10.11 | בחכמה |  |
| 9:10.12 | ותבונה | . |
| 9:10.13 | ובדעת |  |
| 9:10.14 | ישרה |  |
| 9:10.15 | על |  |
| 9:10.16 | פי |  |
| 9:10.17 | קבלת |  |
| 9:10.18 | רז"ל |  |
| 9:10.19 | בעלי |  |
| 9:10.20 | המשנה |  |
| 9:10.21 | והתלמוד |  |
| 9:10.22 | והפוסקים |  |
| 9:10.23 | הדולים |  |
| 9:10.24 | ומשקים |  |
| 9:10.25 | לכל |  |
| 9:10.26 | צמא |  |
| 9:10.27 | ממקור |  |
| 9:10.28 | מים |  |
| 9:10.29 | חיים | , |
| 9:10.30 | משה |  |
| 9:10.31 | קבל |  |
| 9:10.32 | תורה |  |
| 9:10.33 | מסיני |  |
| 9:10.34 | ומסרה |  |
| 9:10.35 | ליהושע |  |
| 9:10.36 | וכו | ' |
| 9:10.37 | והם |  |
| 9:10.38 | מלאו |  |
| 9:10.39 | אחר |  |
| 9:10.40 | ה | ' |
| 9:10.41 | ותורתו |  |
| 9:10.42 | הקדושה |  |
| 9:10.43 | שלא |  |
| 9:10.44 | לפנות |  |
| 9:10.45 | אל |  |
| 9:10.46 | האלילים | , |
| 9:10.47 | איך |  |
| 9:10.48 | בכל |  |
| 9:10.49 | מצותיה |  |
| 9:10.50 | של |  |
| 9:10.51 | תורה |  |
| 9:10.52 | האריכו |  |
| 9:10.53 | הרחיבו |  |
| 9:10.54 | לבאר |  |
| 9:10.55 | חילוקי |  |
| 9:10.56 | דיניהם |  |
| 9:10.57 | בכמה |  |
| 9:10.58 | פרטים | , |
| 9:10.59 | וחומרות |  |
| 9:10.60 | יתירות |  |
| 9:10.61 | ממה |  |
| 9:10.62 | שנזכר |  |
| 9:10.63 | בתלמוד | ! |
| 9:10.64 | ובדיני |  |
| 9:10.65 | יחוד |  |
| 9:10.66 | השי"ת |  |
| 9:10.67 | שהוא |  |
| 9:10.68 | העיקר |  |
| 9:10.69 | הגדול | , |
| 9:10.70 | העלימו |  |
| 9:10.71 | עיניהם |  |
| 9:10.72 | ועזבוהו |  |
| 9:10.73 | ולא |  |
| 9:10.74 | דברו |  |
| 9:10.75 | מאומה |  |
| 9:10.76 | אפי | ' |
| 9:10.77 | בפרט |  |
| 9:10.78 | אחד |  |
| 9:10.79 | מכל |  |
| 9:10.80 | דיני |  |
| 9:10.81 | יחודו |  |
| 9:10.82 | ית | '! |
| 9:10.83 | ולמה |  |
| 9:10.84 | זה |  |
| 9:10.85 | ועל |  |
| 9:10.86 | מה |  |
| 9:10.87 | זה |  |
| 9:10.88 | הניחוהו |  |
| 9:10.89 | כעיר |  |
| 9:10.90 | פרוצה |  |
| 9:10.91 | אין |  |
| 9:10.92 | חומה | , |
| 9:10.93 | והוא |  |
| 9:10.94 | נוהג |  |
| 9:10.95 | בכל |  |
| 9:10.96 | זמן |  |
| 9:10.97 | ובכל |  |
| 9:10.98 | מקום |  |
| 9:10.99 | בין |  |
| 9:10.100 | בארץ |  |
| 9:10.101 | בין |  |
| 9:10.102 | בחוצא |  |
| 9:10.103 | לארץ | , |
| 9:10.104 | בין |  |
| 9:10.105 | בפני |  |
| 9:10.106 | הבית |  |
| 9:10.107 | בין |  |
| 9:10.108 | שלא |  |
| 9:10.109 | בפני |  |
| 9:10.110 | הבית | , |
| 9:10.111 | כאלו |  |
| 9:10.112 | נשתכחה |  |
| 9:10.113 | מישראל |  |
| 9:10.114 | תורת |  |
| 9:10.115 | יחוד |  |
| 9:10.116 | השי"ת |  |
| 9:10.117 | היוצאת |  |
| 9:10.118 | מדברי |  |
| 9:10.119 | התלמוד |  |
| 9:10.120 | בבלי |  |
| 9:10.121 | וירושלמי |  |
| 9:10.122 | ומדרשי |  |
| 9:10.123 | רז"ל | , |
| 9:10.124 | אחרי |  |
| 9:10.125 | מות |  |
| 9:10.126 | ר | ' |
| 9:10.127 | בחיי |  |
| 9:10.128 | בעל |  |
| 9:10.129 | חובת |  |
| 9:10.130 | הלבבות |  |
| 9:10.131 | ורבי |  |
| 9:10.132 | יהודה |  |
| 9:10.133 | הלוי |  |
| 9:10.134 | והרס"ג |  |
| 9:10.135 | והרמב"ם |  |
| 9:10.136 | והסמ"ג |  |
| 9:10.137 | ובעל |  |
| 9:10.138 | העיקרים |  |
| 9:10.139 | וזולתם | ! |
| 9:10.140 | עד |  |
| 9:10.141 | שבעבור |  |
| 9:10.142 | זה |  |
| 9:10.143 | נפלו |  |
| 9:10.144 | קצת |  |
| 9:10.145 | אחרונים |  |
| 9:10.146 | בפח |  |
| 9:10.147 | הנסיון |  |
| 9:10.148 | אשר |  |
| 9:10.149 | נסה |  |
| 9:10.150 | ה | ' |
| 9:10.151 | אלהינו |  |
| 9:10.152 | אותנו |  |
| 9:10.153 | על |  |
| 9:10.154 | ידי |  |
| 9:10.155 | המלך |  |
| 9:10.156 | העכו"ם |  |
| 9:10.157 | בשלחו |  |
| 9:10.158 | ספר |  |
| 9:10.159 | הזהר |  |
| 9:10.160 | לחכמי |  |
| 9:10.161 | טוליטולא | , |
| 9:10.162 | והמה |  |
| 9:10.163 | שמחו |  |
| 9:10.164 | בו |  |
| 9:10.165 | שמחה |  |
| 9:10.166 | גדולה | , |
| 9:10.167 | כי |  |
| 9:10.168 | חפצו |  |
| 9:10.169 | להוציא |  |
| 9:10.170 | ישן |  |
| 9:10.171 | תורת |  |
| 9:10.172 | המשנה |  |
| 9:10.173 | והתלמוד | , |
| 9:10.174 | מפני |  |
| 9:10.175 | חדש |  |
| 9:10.176 | תורת |  |
| 9:10.177 | הזהר |  |
| 9:10.178 | המבזה |  |
| 9:10.179 | את |  |
| 9:10.180 | המשנה |  |
| 9:10.181 | והתלמוד |  |
| 9:10.182 | כמבואר |  |
| 9:10.183 | לעיל | , |
| 9:10.184 | ולא |  |
| 9:10.185 | נתנו |  |
| 9:10.186 | אל |  |
| 9:10.187 | לבם |  |
| 9:10.188 | למנוע |  |
| 9:10.189 | מלעבוד |  |
| 9:10.190 | אלהים |  |
| 9:10.191 | חדשים | , |
| 9:10.192 | מקרבו |  |
| 9:10.193 | באו | , |
| 9:10.194 | אשר |  |
| 9:10.195 | הזהירנו |  |
| 9:10.196 | ה | ' |
| 9:10.197 | אלהינו |  |
| 9:10.198 | לבלתי |  |
| 9:10.199 | עבוד |  |
| 9:10.200 | זולתו | , |
| 9:10.201 | אפי | ' |
| 9:10.202 | נקראהו |  |
| 9:10.203 | בשמו |  |
| 9:10.204 | כמעשה |  |
| 9:10.205 | דורו |  |
| 9:10.206 | של |  |
| 9:10.207 | אנוש | . |
| 9:10.208 | וחכמי |  |
| 9:10.209 | טוליטולא |  |
| 9:10.210 | לא |  |
| 9:10.211 | עשו |  |
| 9:10.212 | כן |  |
| 9:10.213 | אדרבא |  |
| 9:10.214 | שלחו |  |
| 9:10.215 | מתנות |  |
| 9:10.216 | רבות |  |
| 9:10.217 | למלך |  |
| 9:10.218 | העכו"ם |  |
| 9:10.219 | בעבורו |  |
| 9:10.220 | והגדילו |  |
| 9:10.221 | והחשיבו |  |
| 9:10.222 | דברי |  |
| 9:10.223 | הפילסוף |  |
| 9:10.224 | מחבר |  |
| 9:10.225 | הזהר |  |
| 9:10.226 | המתנבא |  |
| 9:10.227 | בשקר |  |
| 9:10.228 | בשם |  |
| 9:10.229 | רשב"י |  |
| 9:10.230 | שנגלה |  |
| 9:10.231 | לו |  |
| 9:10.232 | עתיק | , |
| 9:10.233 | יותר |  |
| 9:10.234 | ממ"ח |  |
| 9:10.235 | נביאים |  |
| 9:10.236 | ושבע |  |
| 9:10.237 | נביאות |  |
| 9:10.238 | שנתנבאו |  |
| 9:10.239 | להם |  |
| 9:10.240 | לישראל | , |
| 9:10.241 | שמ"ח |  |
| 9:10.242 | נביאים |  |
| 9:10.243 | ושבע |  |
| 9:10.244 | נביאות |  |
| 9:10.245 | לא |  |
| 9:10.246 | פיחתו |  |
| 9:10.247 | ולא |  |
| 9:10.248 | הותירו |  |
| 9:10.249 | על |  |
| 9:10.250 | מה |  |
| 9:10.251 | שכתוב |  |
| 9:10.252 | בתורה |  |
| 9:10.253 | אפי | ' |
| 9:10.254 | אות |  |
| 9:10.255 | אחת |  |
| 9:10.256 | חוץ |  |
| 9:10.257 | ממקרא |  |
| 9:10.258 | מגלה | , |
| 9:10.259 | ומחבר |  |
| 9:10.260 | הזהר |  |
| 9:10.261 | הוסיף |  |
| 9:10.262 | אמונת |  |
| 9:10.263 | אלהים |  |
| 9:10.264 | אחרים |  |
| 9:10.265 | חדשים |  |
| 9:10.266 | אשר |  |
| 9:10.267 | לא |  |
| 9:10.268 | שערום |  |
| 9:10.269 | אבותיהם | . |
| 9:10.270 | ואמר |  |
| 9:10.271 | שמי |  |
| 9:10.272 | שלא |  |
| 9:10.273 | יעסוק |  |
| 9:10.274 | בתורתו |  |
| 9:10.275 | החדשה |  |
| 9:10.276 | לית |  |
| 9:10.277 | ליה |  |
| 9:10.278 | רשו |  |
| 9:10.279 | למיעל |  |
| 9:10.280 | קמי |  |
| 9:10.281 | מריה | . |
| 9:10.282 | ובספר |  |
| 9:10.283 | אצהאר |  |
| 9:10.284 | אלחק |  |
| 9:10.285 | לאחד |  |
| 9:10.286 | מחכמי |  |
| 9:10.287 | הערביים |  |
| 9:10.288 | בטענתו |  |
| 9:10.289 | נגד |  |
| 9:10.290 | אמונת |  |
| 9:10.291 | השלוש |  |
| 9:10.292 | שהיא |  |
| 9:10.293 | אמונת |  |
| 9:10.294 | הנצרים |  |
| 9:10.295 | כתב |  |
| 9:10.296 | בדף |  |
| 9:10.297 | רפ"א |  |
| 9:10.298 | ורפ"ב |  |
| 9:10.299 | וז"ל | . |
| 9:10.300 | ועלמא |  |
| 9:10.301 | אליהוד |  |
| 9:10.302 | מן |  |
| 9:10.303 | ודן |  |
| 9:10.304 | מוסי | ' |
| 9:10.305 | אלי | ' |
| 9:10.306 | זמאננא |  |
| 9:10.307 | הדא |  |
| 9:10.308 | לא |  |
| 9:10.309 | יעתרפון |  |
| 9:10.310 | אתתלית |  |
| 9:10.311 | תלת |  |
| 9:10.312 | קשרי |  |
| 9:10.313 | דמהימנותא | ) |
| 9:10.314 | וצאהר |  |
| 9:10.315 | אן |  |
| 9:10.316 | דאת |  |
| 9:10.317 | אללה |  |
| 9:10.318 | וצפאתה |  |
| 9:10.319 | אלכמאליה |  |
| 9:10.320 | קדימא |  |
| 9:10.321 | גיה |  |
| 9:10.322 | מתגיירה |  |
| 9:10.323 | אזלא | ' |
| 9:10.324 | ואבדא | '. |
| 9:10.325 | פלו |  |
| 9:10.326 | כאן |  |
| 9:10.327 | אלתתלית |  |
| 9:10.328 | חקא | ', |
| 9:10.329 | לכאן |  |
| 9:10.330 | אלואגב |  |
| 9:10.331 | עלי | ' |
| 9:10.332 | מוסי | ' |
| 9:10.333 | עליה |  |
| 9:10.334 | אסלאם |  |
| 9:10.335 | ואנביא |  |
| 9:10.336 | בני |  |
| 9:10.337 | אסראיל |  |
| 9:10.338 | אן |  |
| 9:10.339 | יביינוה |  |
| 9:10.340 | חק |  |
| 9:10.341 | אלביאן | , |
| 9:10.342 | פאלעגב |  |
| 9:10.343 | כל |  |
| 9:10.344 | אלעגב |  |
| 9:10.345 | אן |  |
| 9:10.346 | תכון |  |
| 9:10.347 | אלשריעה |  |
| 9:10.348 | אלמוסויה |  |
| 9:10.349 | אלתי |  |
| 9:10.350 | כאנת |  |
| 9:10.351 | ואנבת |  |
| 9:10.352 | אלאטאעה |  |
| 9:10.353 | בגמיע |  |
| 9:10.354 | אלאנביא |  |
| 9:10.355 | אלי | ' |
| 9:10.356 | עהד |  |
| 9:10.357 | עיסי | ' |
| 9:10.358 | כאלה |  |
| 9:10.359 | עם |  |
| 9:10.360 | ביאן |  |
| 9:10.361 | הדה |  |
| 9:10.362 | אלעקידה |  |
| 9:10.363 | אלתי |  |
| 9:10.364 | הי |  |
| 9:10.365 | מדאר |  |
| 9:10.366 | אלנגאה |  |
| 9:10.367 | היא |  |
| 9:10.368 | מרכז |  |
| 9:10.369 | ההצלה |  |
| 9:10.370 | והוא |  |
| 9:10.371 | מאמר |  |
| 9:10.372 | הזהר |  |
| 9:10.373 | דלית |  |
| 9:10.374 | ליה |  |
| 9:10.375 | רשו |  |
| 9:10.376 | למיעל |  |
| 9:10.377 | קמי |  |
| 9:10.378 | מריה | ) |
| 9:10.379 | עלי | ' |
| 9:10.380 | זעם |  |
| 9:10.381 | אהל |  |
| 9:10.382 | אלתתלית |  |
| 9:10.383 | ולא |  |
| 9:10.384 | ימכן |  |
| 9:10.385 | נגאת |  |
| 9:10.386 | אחדכדונהא |  |
| 9:10.387 | נבייא |  |
| 9:10.388 | כאן |  |
| 9:10.389 | או |  |
| 9:10.390 | גיר |  |
| 9:10.391 | נבי | , |
| 9:10.392 | ולא |  |
| 9:10.393 | יביין |  |
| 9:10.394 | מוסי | ' |
| 9:10.395 | ע"ס | . |
| 9:10.396 | ולא |  |
| 9:10.397 | נבי |  |
| 9:10.398 | מן |  |
| 9:10.399 | אלאנביא |  |
| 9:10.400 | אלאסראלייה |  |
| 9:10.401 | הדה |  |
| 9:10.402 | אלעקידה |  |
| 9:10.403 | בביאן |  |
| 9:10.404 | ואצח |  |
| 9:10.405 | בחית |  |
| 9:10.406 | תפהם |  |
| 9:10.407 | מנה |  |
| 9:10.408 | הדה |  |
| 9:10.409 | אלעקידה |  |
| 9:10.410 | צראחה | , |
| 9:10.411 | ולא |  |
| 9:10.412 | יבקש |  |
| 9:10.413 | פיהא |  |
| 9:10.414 | שך |  |
| 9:10.415 | ויביין |  |
| 9:10.416 | מוסי | ' |
| 9:10.417 | אלאחכאם |  |
| 9:10.418 | אלתי |  |
| 9:10.419 | הי |  |
| 9:10.420 | ענד |  |
| 9:10.421 | מקרס |  |
| 9:10.422 | אהל |  |
| 9:10.423 | אלתתלית |  |
| 9:10.424 | צעיפה |  |
| 9:10.425 | נאקצה |  |
| 9:10.426 | גדא |  |
| 9:10.427 | באלתשריח |  |
| 9:10.428 | אלאתם | . |
| 9:10.429 | ויכררהא |  |
| 9:10.430 | מרה |  |
| 9:10.431 | בעד |  |
| 9:10.432 | אולי |  |
| 9:10.433 | וכרה |  |
| 9:10.434 | בעד |  |
| 9:10.435 | אכרא | , |
| 9:10.436 | ויאכד |  |
| 9:10.437 | ען |  |
| 9:10.438 | מחאפצהתא |  |
| 9:10.439 | תאכידא |  |
| 9:10.440 | בלינא | , |
| 9:10.441 | ויוגב |  |
| 9:10.442 | אלקתל |  |
| 9:10.443 | למן |  |
| 9:10.444 | לם |  |
| 9:10.445 | יעתקדהא |  |
| 9:10.446 | עכ"ל | . |
| 9:10.447 | העתק |  |
| 9:10.448 | ועיין |  |
| 9:10.449 | בסוף |  |
| 9:10.450 | הספר | ) |
| 9:10.451 | וטענתו |  |
| 9:10.452 | זאת |  |
| 9:10.453 | בעצמה |  |
| 9:10.454 | אמיתית |  |
| 9:10.455 | צודקת |  |
| 9:10.456 | על |  |
| 9:10.457 | הקבלה |  |
| 9:10.458 | החדשה | , |
| 9:10.459 | ומדבריו |  |
| 9:10.460 | מוכח |  |
| 9:10.461 | שלא |  |
| 9:10.462 | פשטה |  |
| 9:10.463 | עדיין |  |
| 9:10.464 | אמונת |  |
| 9:10.465 | הקבלה |  |
| 9:10.466 | החדשה |  |
| 9:10.467 | בישראל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

