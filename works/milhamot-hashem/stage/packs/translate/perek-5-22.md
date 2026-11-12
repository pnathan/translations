# Translation pack: milhamot-hashem perek-5-22 (Milchamot ha-Shem 5:31–5:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-22`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
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
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
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
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **השפע** → "the flow".
- **עלה / עלול** → "cause / effect".
- תורה **שבעל** פה → "the Oral Torah".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ת"ל** → "Scripture says".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **והוא שלא** → "provided that".
- **קורין** אל → "call upon".
- אחד ... **ואחד** of two kinds of case → "alike".
- **אלפין** → "alefs".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 5:31

Hebrew (Wikisource): נז) יבין המבין וישכיל היטב בדבר רבותינו ז"ל הנז' פה כי רעיונות הפילסוף מחבר הזהר והמקובלים הנמשכים אחריו מכוונים מאד למחשבות דורו של אנוש שקראו לעצבים בשמו של הקב"ה והתפללו אליהם, והמקובלים גם הם עשו כן, כי קראו לכל אחד מן הפרצופים בשמו של הקב"ה ואמרו לעבדן ולהשתחות להם ולקרוא אליהם בעת צרה וכל התפלות והברכות שתקנו אנשי כנה"ג לשבח ולפאר לשי"ת בהן. את הכל הסבו לפרצוף ז"א בצירוף שאר פרצופים הללו אע"פ שהן נבראים. וחשבו לתקן הכפירה הנגלית ונראית באמונה זו באמרם שהם עובדים לנשמתם. ושכחו את האמונה האמתית המפורסמת ופשוטה בכל כתבי הקדש תורה שבכתב ותורה שבעל פה. כי אלהינו ב"ה אינו גוף ולא כח בגוף. והם המציאו ועשו לאלהינו כמה גופים ומלבושים מתוארים בשרד ובמחוגה אורך ורוחב להם. ונתנו להם נשמות מהאלהים העליון. והשאירו עיקרו לעילא לעילא מבלי להתפלל אליו, כי לא השאירו לו שם. ותלו כל מצותיה של תורה וספוריה המלמדים לאדם דעת בגופא דמלכא (ולא בנשמתו) מנהון בידוי מנהון ברגלוי וכו' כמו שכ' מהר"י צאהרי בשם הזהר, וסותרים ג"כ דבריהם שכתבתי לעיל סי' ט"ל בן שושן סודות ומשנת חסידים והרמ"ז דבין הגוף של ספירות בין הנשמה בין המלבוש הכל אלהות גמור. כי כולם נמשכים ונהוים מעצמות האין סוף כהדין קמצא דלבושיה מניה וביה. ומשתלשלים זה מזה כבן מן האב. ולפיכך קורין לז"א בן לאבא ואימא. והוא השולט על כל הברואים וזנם ומפרנסם במצות אביו ואמו כדאיתא בזהר בלק דף קצ"א (ח"ג קצא, א). והובא לעיל סי' כ"ו. ושם מבואר שהעבודה והברכות שייכים (לפי דעתם) לז"א דהכי פקידו אבא ואימא דכלא יפלחון ליה וכמו שכ' ג"כ ס' הברית וכס"א ויושר לבב ונחלת יוסף. הנוטע אזן הלא ישמע. ויוצר עין לאדם ולב להבין ולהשכיל שכן היא מחשבת דורו של אנוש וטעותם אשר חשבו את הצורות ואת הפרצופים אשר בדו מלבם לאלוהות וקראו להם בשמותיו של הקב"ה ועבדום להיות אמצעיים להמשיך השפע מהאלהים העליון. ואם תאמר שאני מעשי דורו של אנוש שעשו פרצופים וצורות מעשה ידי אדם וקראום בשמו של הקב"ה ועבדום. אבל המקובלים לא עשו שום צורה בםועל רק חשבו הפרצופים והצורות שהן מעשה השם לאלוהות? הלא גם דורו של אנוש כן עלה בדעתם בתחלה כמו שכ' הרמב"ם ז"ל בפ"א מהל' ע"ש. והרי כבר אמרו רז"ל יכול לא יעשה אבל יעלה בדעתו. ת"ל לא יהיה לך אלהים אחרים מכל מקום. וכן הוא אומר ופן תשא עיניך השמימה וראית את השמש ואת הירח ואת הכוכבים כל צבא השמים ונדחת. הכל בכלל, שלא יעלה בדעתו אלהות וממשלה לשום נברא או נאצל. כי כולם בשם צבא השמים שהזהיר הכתוב מלעבדן. ואף אם קורא להן בשמותיו של הקב"ה כמו שעשו דורו של אנוש אין בהן ממש, כמו שמבואר בילקוט ומדרש הגדול וברש"י שהבאתי לעיל, סי' נ"ו, כי לא ברא הקב"ה דבר מבריותיו להיות נעבד, כי אם להנהיג את העולם כחפצו ורצונו, והמה כעבדים העומדים לפני המלך לעשות רצונו שנ' אלף אלפין ישמשוניה ורבוא רבוון קדמוהי יקומון, לא להיות אלוהות ח"ו ולא לעבדן. כי אין אלוה מבלעדי ה' ואין צור כאלהינו. והוא קרוב לכל קוראיו באמת, ואף אם ברגע א', יקראו אליו כל יצוריו ויגיד כל אחד מהם, ועונה אותו וממלא משאלו שלכל אחד ואחד אם הוא ראוי לפי מעשיו, כי עונותיו של אדם מעכבים מלעשות בקשתו, שנ' לתת לאיש כדרכיו, כפרי מעלליו, ונאמר כי אם עונותיכם היו מבדילים ביניכם לבין אלהיכם וכמו שכ' הרמב"ם בהל' תשובה.

| id | word | punctuation after |
|---|---|---|
| 5:31.1 | נז | ) |
| 5:31.2 | יבין |  |
| 5:31.3 | המבין |  |
| 5:31.4 | וישכיל |  |
| 5:31.5 | היטב |  |
| 5:31.6 | בדבר |  |
| 5:31.7 | רבותינו |  |
| 5:31.8 | ז"ל |  |
| 5:31.9 | הנז | ' |
| 5:31.10 | פה |  |
| 5:31.11 | כי |  |
| 5:31.12 | רעיונות |  |
| 5:31.13 | הפילסוף |  |
| 5:31.14 | מחבר |  |
| 5:31.15 | הזהר |  |
| 5:31.16 | והמקובלים |  |
| 5:31.17 | הנמשכים |  |
| 5:31.18 | אחריו |  |
| 5:31.19 | מכוונים |  |
| 5:31.20 | מאד |  |
| 5:31.21 | למחשבות |  |
| 5:31.22 | דורו |  |
| 5:31.23 | של |  |
| 5:31.24 | אנוש |  |
| 5:31.25 | שקראו |  |
| 5:31.26 | לעצבים |  |
| 5:31.27 | בשמו |  |
| 5:31.28 | של |  |
| 5:31.29 | הקב"ה |  |
| 5:31.30 | והתפללו |  |
| 5:31.31 | אליהם | , |
| 5:31.32 | והמקובלים |  |
| 5:31.33 | גם |  |
| 5:31.34 | הם |  |
| 5:31.35 | עשו |  |
| 5:31.36 | כן | , |
| 5:31.37 | כי |  |
| 5:31.38 | קראו |  |
| 5:31.39 | לכל |  |
| 5:31.40 | אחד |  |
| 5:31.41 | מן |  |
| 5:31.42 | הפרצופים |  |
| 5:31.43 | בשמו |  |
| 5:31.44 | של |  |
| 5:31.45 | הקב"ה |  |
| 5:31.46 | ואמרו |  |
| 5:31.47 | לעבדן |  |
| 5:31.48 | ולהשתחות |  |
| 5:31.49 | להם |  |
| 5:31.50 | ולקרוא |  |
| 5:31.51 | אליהם |  |
| 5:31.52 | בעת |  |
| 5:31.53 | צרה |  |
| 5:31.54 | וכל |  |
| 5:31.55 | התפלות |  |
| 5:31.56 | והברכות |  |
| 5:31.57 | שתקנו |  |
| 5:31.58 | אנשי |  |
| 5:31.59 | כנה"ג |  |
| 5:31.60 | לשבח |  |
| 5:31.61 | ולפאר |  |
| 5:31.62 | לשי"ת |  |
| 5:31.63 | בהן | . |
| 5:31.64 | את |  |
| 5:31.65 | הכל |  |
| 5:31.66 | הסבו |  |
| 5:31.67 | לפרצוף |  |
| 5:31.68 | ז"א |  |
| 5:31.69 | בצירוף |  |
| 5:31.70 | שאר |  |
| 5:31.71 | פרצופים |  |
| 5:31.72 | הללו |  |
| 5:31.73 | אע"פ |  |
| 5:31.74 | שהן |  |
| 5:31.75 | נבראים | . |
| 5:31.76 | וחשבו |  |
| 5:31.77 | לתקן |  |
| 5:31.78 | הכפירה |  |
| 5:31.79 | הנגלית |  |
| 5:31.80 | ונראית |  |
| 5:31.81 | באמונה |  |
| 5:31.82 | זו |  |
| 5:31.83 | באמרם |  |
| 5:31.84 | שהם |  |
| 5:31.85 | עובדים |  |
| 5:31.86 | לנשמתם | . |
| 5:31.87 | ושכחו |  |
| 5:31.88 | את |  |
| 5:31.89 | האמונה |  |
| 5:31.90 | האמתית |  |
| 5:31.91 | המפורסמת |  |
| 5:31.92 | ופשוטה |  |
| 5:31.93 | בכל |  |
| 5:31.94 | כתבי |  |
| 5:31.95 | הקדש |  |
| 5:31.96 | תורה |  |
| 5:31.97 | שבכתב |  |
| 5:31.98 | ותורה |  |
| 5:31.99 | שבעל |  |
| 5:31.100 | פה | . |
| 5:31.101 | כי |  |
| 5:31.102 | אלהינו |  |
| 5:31.103 | ב"ה |  |
| 5:31.104 | אינו |  |
| 5:31.105 | גוף |  |
| 5:31.106 | ולא |  |
| 5:31.107 | כח |  |
| 5:31.108 | בגוף | . |
| 5:31.109 | והם |  |
| 5:31.110 | המציאו |  |
| 5:31.111 | ועשו |  |
| 5:31.112 | לאלהינו |  |
| 5:31.113 | כמה |  |
| 5:31.114 | גופים |  |
| 5:31.115 | ומלבושים |  |
| 5:31.116 | מתוארים |  |
| 5:31.117 | בשרד |  |
| 5:31.118 | ובמחוגה |  |
| 5:31.119 | אורך |  |
| 5:31.120 | ורוחב |  |
| 5:31.121 | להם | . |
| 5:31.122 | ונתנו |  |
| 5:31.123 | להם |  |
| 5:31.124 | נשמות |  |
| 5:31.125 | מהאלהים |  |
| 5:31.126 | העליון | . |
| 5:31.127 | והשאירו |  |
| 5:31.128 | עיקרו |  |
| 5:31.129 | לעילא |  |
| 5:31.130 | לעילא |  |
| 5:31.131 | מבלי |  |
| 5:31.132 | להתפלל |  |
| 5:31.133 | אליו | , |
| 5:31.134 | כי |  |
| 5:31.135 | לא |  |
| 5:31.136 | השאירו |  |
| 5:31.137 | לו |  |
| 5:31.138 | שם | . |
| 5:31.139 | ותלו |  |
| 5:31.140 | כל |  |
| 5:31.141 | מצותיה |  |
| 5:31.142 | של |  |
| 5:31.143 | תורה |  |
| 5:31.144 | וספוריה |  |
| 5:31.145 | המלמדים |  |
| 5:31.146 | לאדם |  |
| 5:31.147 | דעת |  |
| 5:31.148 | בגופא |  |
| 5:31.149 | דמלכא |  |
| 5:31.150 | ולא |  |
| 5:31.151 | בנשמתו | ) |
| 5:31.152 | מנהון |  |
| 5:31.153 | בידוי |  |
| 5:31.154 | מנהון |  |
| 5:31.155 | ברגלוי |  |
| 5:31.156 | וכו | ' |
| 5:31.157 | כמו |  |
| 5:31.158 | שכ | ' |
| 5:31.159 | מהר"י |  |
| 5:31.160 | צאהרי |  |
| 5:31.161 | בשם |  |
| 5:31.162 | הזהר | , |
| 5:31.163 | וסותרים |  |
| 5:31.164 | ג"כ |  |
| 5:31.165 | דבריהם |  |
| 5:31.166 | שכתבתי |  |
| 5:31.167 | לעיל |  |
| 5:31.168 | סי | ' |
| 5:31.169 | ט"ל |  |
| 5:31.170 | בן |  |
| 5:31.171 | שושן |  |
| 5:31.172 | סודות |  |
| 5:31.173 | ומשנת |  |
| 5:31.174 | חסידים |  |
| 5:31.175 | והרמ"ז |  |
| 5:31.176 | דבין |  |
| 5:31.177 | הגוף |  |
| 5:31.178 | של |  |
| 5:31.179 | ספירות |  |
| 5:31.180 | בין |  |
| 5:31.181 | הנשמה |  |
| 5:31.182 | בין |  |
| 5:31.183 | המלבוש |  |
| 5:31.184 | הכל |  |
| 5:31.185 | אלהות |  |
| 5:31.186 | גמור | . |
| 5:31.187 | כי |  |
| 5:31.188 | כולם |  |
| 5:31.189 | נמשכים |  |
| 5:31.190 | ונהוים |  |
| 5:31.191 | מעצמות |  |
| 5:31.192 | האין |  |
| 5:31.193 | סוף |  |
| 5:31.194 | כהדין |  |
| 5:31.195 | קמצא |  |
| 5:31.196 | דלבושיה |  |
| 5:31.197 | מניה |  |
| 5:31.198 | וביה | . |
| 5:31.199 | ומשתלשלים |  |
| 5:31.200 | זה |  |
| 5:31.201 | מזה |  |
| 5:31.202 | כבן |  |
| 5:31.203 | מן |  |
| 5:31.204 | האב | . |
| 5:31.205 | ולפיכך |  |
| 5:31.206 | קורין |  |
| 5:31.207 | לז"א |  |
| 5:31.208 | בן |  |
| 5:31.209 | לאבא |  |
| 5:31.210 | ואימא | . |
| 5:31.211 | והוא |  |
| 5:31.212 | השולט |  |
| 5:31.213 | על |  |
| 5:31.214 | כל |  |
| 5:31.215 | הברואים |  |
| 5:31.216 | וזנם |  |
| 5:31.217 | ומפרנסם |  |
| 5:31.218 | במצות |  |
| 5:31.219 | אביו |  |
| 5:31.220 | ואמו |  |
| 5:31.221 | כדאיתא |  |
| 5:31.222 | בזהר |  |
| 5:31.223 | בלק |  |
| 5:31.224 | דף |  |
| 5:31.225 | קצ"א |  |
| 5:31.226 | ח"ג |  |
| 5:31.227 | קצא | , |
| 5:31.228 | א | ). |
| 5:31.229 | והובא |  |
| 5:31.230 | לעיל |  |
| 5:31.231 | סי | ' |
| 5:31.232 | כ"ו | . |
| 5:31.233 | ושם |  |
| 5:31.234 | מבואר |  |
| 5:31.235 | שהעבודה |  |
| 5:31.236 | והברכות |  |
| 5:31.237 | שייכים |  |
| 5:31.238 | לפי |  |
| 5:31.239 | דעתם | ) |
| 5:31.240 | לז"א |  |
| 5:31.241 | דהכי |  |
| 5:31.242 | פקידו |  |
| 5:31.243 | אבא |  |
| 5:31.244 | ואימא |  |
| 5:31.245 | דכלא |  |
| 5:31.246 | יפלחון |  |
| 5:31.247 | ליה |  |
| 5:31.248 | וכמו |  |
| 5:31.249 | שכ | ' |
| 5:31.250 | ג"כ |  |
| 5:31.251 | ס | ' |
| 5:31.252 | הברית |  |
| 5:31.253 | וכס"א |  |
| 5:31.254 | ויושר |  |
| 5:31.255 | לבב |  |
| 5:31.256 | ונחלת |  |
| 5:31.257 | יוסף | . |
| 5:31.258 | הנוטע |  |
| 5:31.259 | אזן |  |
| 5:31.260 | הלא |  |
| 5:31.261 | ישמע | . |
| 5:31.262 | ויוצר |  |
| 5:31.263 | עין |  |
| 5:31.264 | לאדם |  |
| 5:31.265 | ולב |  |
| 5:31.266 | להבין |  |
| 5:31.267 | ולהשכיל |  |
| 5:31.268 | שכן |  |
| 5:31.269 | היא |  |
| 5:31.270 | מחשבת |  |
| 5:31.271 | דורו |  |
| 5:31.272 | של |  |
| 5:31.273 | אנוש |  |
| 5:31.274 | וטעותם |  |
| 5:31.275 | אשר |  |
| 5:31.276 | חשבו |  |
| 5:31.277 | את |  |
| 5:31.278 | הצורות |  |
| 5:31.279 | ואת |  |
| 5:31.280 | הפרצופים |  |
| 5:31.281 | אשר |  |
| 5:31.282 | בדו |  |
| 5:31.283 | מלבם |  |
| 5:31.284 | לאלוהות |  |
| 5:31.285 | וקראו |  |
| 5:31.286 | להם |  |
| 5:31.287 | בשמותיו |  |
| 5:31.288 | של |  |
| 5:31.289 | הקב"ה |  |
| 5:31.290 | ועבדום |  |
| 5:31.291 | להיות |  |
| 5:31.292 | אמצעיים |  |
| 5:31.293 | להמשיך |  |
| 5:31.294 | השפע |  |
| 5:31.295 | מהאלהים |  |
| 5:31.296 | העליון | . |
| 5:31.297 | ואם |  |
| 5:31.298 | תאמר |  |
| 5:31.299 | שאני |  |
| 5:31.300 | מעשי |  |
| 5:31.301 | דורו |  |
| 5:31.302 | של |  |
| 5:31.303 | אנוש |  |
| 5:31.304 | שעשו |  |
| 5:31.305 | פרצופים |  |
| 5:31.306 | וצורות |  |
| 5:31.307 | מעשה |  |
| 5:31.308 | ידי |  |
| 5:31.309 | אדם |  |
| 5:31.310 | וקראום |  |
| 5:31.311 | בשמו |  |
| 5:31.312 | של |  |
| 5:31.313 | הקב"ה |  |
| 5:31.314 | ועבדום | . |
| 5:31.315 | אבל |  |
| 5:31.316 | המקובלים |  |
| 5:31.317 | לא |  |
| 5:31.318 | עשו |  |
| 5:31.319 | שום |  |
| 5:31.320 | צורה |  |
| 5:31.321 | בםועל |  |
| 5:31.322 | רק |  |
| 5:31.323 | חשבו |  |
| 5:31.324 | הפרצופים |  |
| 5:31.325 | והצורות |  |
| 5:31.326 | שהן |  |
| 5:31.327 | מעשה |  |
| 5:31.328 | השם |  |
| 5:31.329 | לאלוהות | ? |
| 5:31.330 | הלא |  |
| 5:31.331 | גם |  |
| 5:31.332 | דורו |  |
| 5:31.333 | של |  |
| 5:31.334 | אנוש |  |
| 5:31.335 | כן |  |
| 5:31.336 | עלה |  |
| 5:31.337 | בדעתם |  |
| 5:31.338 | בתחלה |  |
| 5:31.339 | כמו |  |
| 5:31.340 | שכ | ' |
| 5:31.341 | הרמב"ם |  |
| 5:31.342 | ז"ל |  |
| 5:31.343 | בפ"א |  |
| 5:31.344 | מהל | ' |
| 5:31.345 | ע"ש | . |
| 5:31.346 | והרי |  |
| 5:31.347 | כבר |  |
| 5:31.348 | אמרו |  |
| 5:31.349 | רז"ל |  |
| 5:31.350 | יכול |  |
| 5:31.351 | לא |  |
| 5:31.352 | יעשה |  |
| 5:31.353 | אבל |  |
| 5:31.354 | יעלה |  |
| 5:31.355 | בדעתו | . |
| 5:31.356 | ת"ל |  |
| 5:31.357 | לא |  |
| 5:31.358 | יהיה |  |
| 5:31.359 | לך |  |
| 5:31.360 | אלהים |  |
| 5:31.361 | אחרים |  |
| 5:31.362 | מכל |  |
| 5:31.363 | מקום | . |
| 5:31.364 | וכן |  |
| 5:31.365 | הוא |  |
| 5:31.366 | אומר |  |
| 5:31.367 | ופן |  |
| 5:31.368 | תשא |  |
| 5:31.369 | עיניך |  |
| 5:31.370 | השמימה |  |
| 5:31.371 | וראית |  |
| 5:31.372 | את |  |
| 5:31.373 | השמש |  |
| 5:31.374 | ואת |  |
| 5:31.375 | הירח |  |
| 5:31.376 | ואת |  |
| 5:31.377 | הכוכבים |  |
| 5:31.378 | כל |  |
| 5:31.379 | צבא |  |
| 5:31.380 | השמים |  |
| 5:31.381 | ונדחת | . |
| 5:31.382 | הכל |  |
| 5:31.383 | בכלל | , |
| 5:31.384 | שלא |  |
| 5:31.385 | יעלה |  |
| 5:31.386 | בדעתו |  |
| 5:31.387 | אלהות |  |
| 5:31.388 | וממשלה |  |
| 5:31.389 | לשום |  |
| 5:31.390 | נברא |  |
| 5:31.391 | או |  |
| 5:31.392 | נאצל | . |
| 5:31.393 | כי |  |
| 5:31.394 | כולם |  |
| 5:31.395 | בשם |  |
| 5:31.396 | צבא |  |
| 5:31.397 | השמים |  |
| 5:31.398 | שהזהיר |  |
| 5:31.399 | הכתוב |  |
| 5:31.400 | מלעבדן | . |
| 5:31.401 | ואף |  |
| 5:31.402 | אם |  |
| 5:31.403 | קורא |  |
| 5:31.404 | להן |  |
| 5:31.405 | בשמותיו |  |
| 5:31.406 | של |  |
| 5:31.407 | הקב"ה |  |
| 5:31.408 | כמו |  |
| 5:31.409 | שעשו |  |
| 5:31.410 | דורו |  |
| 5:31.411 | של |  |
| 5:31.412 | אנוש |  |
| 5:31.413 | אין |  |
| 5:31.414 | בהן |  |
| 5:31.415 | ממש | , |
| 5:31.416 | כמו |  |
| 5:31.417 | שמבואר |  |
| 5:31.418 | בילקוט |  |
| 5:31.419 | ומדרש |  |
| 5:31.420 | הגדול |  |
| 5:31.421 | וברש"י |  |
| 5:31.422 | שהבאתי |  |
| 5:31.423 | לעיל | , |
| 5:31.424 | סי | ' |
| 5:31.425 | נ"ו | , |
| 5:31.426 | כי |  |
| 5:31.427 | לא |  |
| 5:31.428 | ברא |  |
| 5:31.429 | הקב"ה |  |
| 5:31.430 | דבר |  |
| 5:31.431 | מבריותיו |  |
| 5:31.432 | להיות |  |
| 5:31.433 | נעבד | , |
| 5:31.434 | כי |  |
| 5:31.435 | אם |  |
| 5:31.436 | להנהיג |  |
| 5:31.437 | את |  |
| 5:31.438 | העולם |  |
| 5:31.439 | כחפצו |  |
| 5:31.440 | ורצונו | , |
| 5:31.441 | והמה |  |
| 5:31.442 | כעבדים |  |
| 5:31.443 | העומדים |  |
| 5:31.444 | לפני |  |
| 5:31.445 | המלך |  |
| 5:31.446 | לעשות |  |
| 5:31.447 | רצונו |  |
| 5:31.448 | שנ | ' |
| 5:31.449 | אלף |  |
| 5:31.450 | אלפין |  |
| 5:31.451 | ישמשוניה |  |
| 5:31.452 | ורבוא |  |
| 5:31.453 | רבוון |  |
| 5:31.454 | קדמוהי |  |
| 5:31.455 | יקומון | , |
| 5:31.456 | לא |  |
| 5:31.457 | להיות |  |
| 5:31.458 | אלוהות |  |
| 5:31.459 | ח"ו |  |
| 5:31.460 | ולא |  |
| 5:31.461 | לעבדן | . |
| 5:31.462 | כי |  |
| 5:31.463 | אין |  |
| 5:31.464 | אלוה |  |
| 5:31.465 | מבלעדי |  |
| 5:31.466 | ה | ' |
| 5:31.467 | ואין |  |
| 5:31.468 | צור |  |
| 5:31.469 | כאלהינו | . |
| 5:31.470 | והוא |  |
| 5:31.471 | קרוב |  |
| 5:31.472 | לכל |  |
| 5:31.473 | קוראיו |  |
| 5:31.474 | באמת | , |
| 5:31.475 | ואף |  |
| 5:31.476 | אם |  |
| 5:31.477 | ברגע |  |
| 5:31.478 | א | ', |
| 5:31.479 | יקראו |  |
| 5:31.480 | אליו |  |
| 5:31.481 | כל |  |
| 5:31.482 | יצוריו |  |
| 5:31.483 | ויגיד |  |
| 5:31.484 | כל |  |
| 5:31.485 | אחד |  |
| 5:31.486 | מהם | , |
| 5:31.487 | ועונה |  |
| 5:31.488 | אותו |  |
| 5:31.489 | וממלא |  |
| 5:31.490 | משאלו |  |
| 5:31.491 | שלכל |  |
| 5:31.492 | אחד |  |
| 5:31.493 | ואחד |  |
| 5:31.494 | אם |  |
| 5:31.495 | הוא |  |
| 5:31.496 | ראוי |  |
| 5:31.497 | לפי |  |
| 5:31.498 | מעשיו | , |
| 5:31.499 | כי |  |
| 5:31.500 | עונותיו |  |
| 5:31.501 | של |  |
| 5:31.502 | אדם |  |
| 5:31.503 | מעכבים |  |
| 5:31.504 | מלעשות |  |
| 5:31.505 | בקשתו | , |
| 5:31.506 | שנ | ' |
| 5:31.507 | לתת |  |
| 5:31.508 | לאיש |  |
| 5:31.509 | כדרכיו | , |
| 5:31.510 | כפרי |  |
| 5:31.511 | מעלליו | , |
| 5:31.512 | ונאמר |  |
| 5:31.513 | כי |  |
| 5:31.514 | אם |  |
| 5:31.515 | עונותיכם |  |
| 5:31.516 | היו |  |
| 5:31.517 | מבדילים |  |
| 5:31.518 | ביניכם |  |
| 5:31.519 | לבין |  |
| 5:31.520 | אלהיכם |  |
| 5:31.521 | וכמו |  |
| 5:31.522 | שכ | ' |
| 5:31.523 | הרמב"ם |  |
| 5:31.524 | בהל | ' |
| 5:31.525 | תשובה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

