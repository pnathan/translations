# Translation pack: milhamot-hashem perek-6-06 (Milchamot ha-Shem 6:11–6:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-06`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רמ"א** → "Rema" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **ממלכי** → "of the kings of".
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
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- עלת **העלות** → "the Cause of causes".
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
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
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- אל **אחר** → "another god".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
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
- **ג"כ** → "also".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **אורה** זו תורה (Esth 8:16) → "'orah' means Torah".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 6:11

Hebrew (Wikisource): סה) עוד כ' בעל הגדרים בשורש יחוד וז"ל יחוד נגזר מן יחיד שמשו בו להורות על האחדות האמתי אשר אין בו הרכבה כלל. אשר לא יבינהו רק החכם הרגיל בשלילת הגשם ומשיגיו. ואיש בער לא ידע שלילה. וכסיל לא יבינהו רק החכם והנה האחדות האמתי היא מצוה כתובה בתורה. שמע ישראל ה' אלהינו ה' אחד. ומהמבואר שאין לפסוק שום מצוה מן התורה שבכתב לבד כי אם מן התורה שבעל פה שהיא ביאור המצות! והנה מענין מצות האחדות לא מצאנו שדברי חכמים ז"ל בביאורה. ובלי ספק שהיו מוסרין ביאורו איש לאיש בקבלה. כי לא יתכן שיעזבו ראשית המצוה ויסוד האמונה האלהית בלי ביאור! אך עתה שהביאור ההוא נשכח מאתנו עלינו לקחת תורה מפי הפילסופים (התורניים) למען לא נתעה בלי דעת להאמין בשניות ובגשמות (כהמקובלים החדשים) ונעבוד ללא אלהי אמת. והוא הדבר שאמר בעל החובות ז"ל לא יוכל לעבוד עלת העלות אלא נביא הדור או החכם המובהק במה שקנהו מן החכמה. וזולתם עובדים זולתו. ובאמת דבר החכם כי כמה דברים מורכבים שנים יתוארו בשם אחד. על דרך ההסכמה ע"כ. ביאור דבריו כגון שם כירים שהוא מקום שפיתת שתי קדרות, ונקרא בדברי רבותינו בשם אחד כירה. וכן רחים שהוא שם השכב העליון רכב, ונקראו שניהם בשם אחד טחון. בחורים טחון נשאו ובערבי רחא אע"פ שהם זוג. וכן מספרים נקרא בערבי מקץ. וכן כל כיוצא בהם כדתנן בפרק י"ג ממסכת כלים קלגריפון, ומכחול, ומכתב, וזומא ליסטרא, אע"פ שהם שנים מחוברים ומורכבים נקראים בשם אחד על דרך ההסכמה וההעברה. ולא כן היא אחדותו ית' שהוא אחד אמיתי שאין בו שום הרכבה וצירוף, שלא כדעת המקובלים החדשים שאומרים שיש לו כמה פרצופים שנאצלו ממנו. וכל אחד נקרא בשם בפני עצמו, והם מחברים ומשתפים אותם בפיהם ובדעתם עם האלוה המאצילם ואומרים וכלא חד. ועובדים אותם יחד גם הגוף גם הנשמה, ומשקרים לומר כי את נשמתו הם עובדים, וזה לשון עוז לאלהים בבית קדש הקדשים פ' י"ח דף נ"ו ע"ב ואם תאמר שאהיה עובד לנשמה בלא גופא ח"ו, חלילה, אלא כבר אמרנו לעיל שזה הפרצוף דמלכא קדישא ושכינתיה הוא אדוק בנשמה. לפי שהוא עתיקא קדישא דכל קדישין. והוא ונשמתא חד אינון כמא דאמרינן לעיל הוא ושמיה חד אינון ולא מתפרש מנהון לעלמין. ואם יעבוד לנשמה בלי מלכא קדישא יהיה פירוד בין נשמה לגופא. וכן מבואר מדברי גליון הזהר ומשנת חסידים הובא לעיל סי' ט"ל. ועל מה שכ' בעל הגדרים שלא יתכן שיעזבו רז"ל ראשית המצוה ויסודה וכו' כבר העירותי על דבריו לעיל סי' ס"ד שאדרבה נהפוך הוא שכל דבר המפורש בכתובים והוא פשוט ומפורסם בישראל לא ידברו בו חכמי המשנה. ובעלי התלמוד ג"כ לא שמו מגמתם רק לבאר ולישא וליתן בדבריה ולבררה וללבנה, ולהודיע דינים ושאלות שנתחדשו בימיהם, וגזרות ותקנות שתיקנו, אבל דברים הפשוטים אין צורך להזכירם כלל. ודעות ורעיונות הקבלה החדשה עוד לא נודעו ולא נשמעו בימיהם בין היהודים כלל, וזרים המה לאמונת תורתינו הקדושה. כי רק מתחלת האלף הששי בסוף המאה הראשונה החזיקו ברעיונות האלה קצת מחברים שנמשכו אחר דעות הפילסוף הרמאי שדבק בדעת השונים והשלישי הוא מחבר ספר הזהר והתקונים. וכוון בו להטעות את ישראל ולהדיחם מעל ה' אלהינו. ותלה רעיונותיו ברשב"י התנא וחביריו, בפרט בשני דברים, מידי דעל רשב"י תני במאמרו הקצר כל המשתף שם שמים ודבר אחר נעקר מן העולם. והשני בביזוי המשנה שקלא וטריא דבתלמודא שהם עיקר ויסוד תורה שבעל פה שהוא מבזה אותן כמו שיתבאר להלן (דלרשב"י שתורתו אומנותו היא כקריאת שמע בעונתה ירוש' שבת פ"א). וטמן את הספר הנז' בארון וקברו בקרקע כדי לישנו כאלו הוא ישן נושן מימי רשב"י. ואמר למלך ממלכי מזרח לחפור שם בקרקע למצוא מטמון של כסף וזהב, וצוה המלך הנז' לחפור על עסקי ממון ונמצא שם ארון ובו ס' הזהר וכו' ושלח אותו לטוליטולא וכו' ומשם נתפשטה הקבלה החדשה בישראל כמו שכ' הר' חיד"א בספרו הנחמד שם הגדולים במערכת ספרים ערך זהר.

| id | word | punctuation after |
|---|---|---|
| 6:11.1 | סה | ) |
| 6:11.2 | עוד |  |
| 6:11.3 | כ | ' |
| 6:11.4 | בעל |  |
| 6:11.5 | הגדרים |  |
| 6:11.6 | בשורש |  |
| 6:11.7 | יחוד |  |
| 6:11.8 | וז"ל |  |
| 6:11.9 | יחוד |  |
| 6:11.10 | נגזר |  |
| 6:11.11 | מן |  |
| 6:11.12 | יחיד |  |
| 6:11.13 | שמשו |  |
| 6:11.14 | בו |  |
| 6:11.15 | להורות |  |
| 6:11.16 | על |  |
| 6:11.17 | האחדות |  |
| 6:11.18 | האמתי |  |
| 6:11.19 | אשר |  |
| 6:11.20 | אין |  |
| 6:11.21 | בו |  |
| 6:11.22 | הרכבה |  |
| 6:11.23 | כלל | . |
| 6:11.24 | אשר |  |
| 6:11.25 | לא |  |
| 6:11.26 | יבינהו |  |
| 6:11.27 | רק |  |
| 6:11.28 | החכם |  |
| 6:11.29 | הרגיל |  |
| 6:11.30 | בשלילת |  |
| 6:11.31 | הגשם |  |
| 6:11.32 | ומשיגיו | . |
| 6:11.33 | ואיש |  |
| 6:11.34 | בער |  |
| 6:11.35 | לא |  |
| 6:11.36 | ידע |  |
| 6:11.37 | שלילה | . |
| 6:11.38 | וכסיל |  |
| 6:11.39 | לא |  |
| 6:11.40 | יבינהו |  |
| 6:11.41 | רק |  |
| 6:11.42 | החכם |  |
| 6:11.43 | והנה |  |
| 6:11.44 | האחדות |  |
| 6:11.45 | האמתי |  |
| 6:11.46 | היא |  |
| 6:11.47 | מצוה |  |
| 6:11.48 | כתובה |  |
| 6:11.49 | בתורה | . |
| 6:11.50 | שמע |  |
| 6:11.51 | ישראל |  |
| 6:11.52 | ה | ' |
| 6:11.53 | אלהינו |  |
| 6:11.54 | ה | ' |
| 6:11.55 | אחד | . |
| 6:11.56 | ומהמבואר |  |
| 6:11.57 | שאין |  |
| 6:11.58 | לפסוק |  |
| 6:11.59 | שום |  |
| 6:11.60 | מצוה |  |
| 6:11.61 | מן |  |
| 6:11.62 | התורה |  |
| 6:11.63 | שבכתב |  |
| 6:11.64 | לבד |  |
| 6:11.65 | כי |  |
| 6:11.66 | אם |  |
| 6:11.67 | מן |  |
| 6:11.68 | התורה |  |
| 6:11.69 | שבעל |  |
| 6:11.70 | פה |  |
| 6:11.71 | שהיא |  |
| 6:11.72 | ביאור |  |
| 6:11.73 | המצות | ! |
| 6:11.74 | והנה |  |
| 6:11.75 | מענין |  |
| 6:11.76 | מצות |  |
| 6:11.77 | האחדות |  |
| 6:11.78 | לא |  |
| 6:11.79 | מצאנו |  |
| 6:11.80 | שדברי |  |
| 6:11.81 | חכמים |  |
| 6:11.82 | ז"ל |  |
| 6:11.83 | בביאורה | . |
| 6:11.84 | ובלי |  |
| 6:11.85 | ספק |  |
| 6:11.86 | שהיו |  |
| 6:11.87 | מוסרין |  |
| 6:11.88 | ביאורו |  |
| 6:11.89 | איש |  |
| 6:11.90 | לאיש |  |
| 6:11.91 | בקבלה | . |
| 6:11.92 | כי |  |
| 6:11.93 | לא |  |
| 6:11.94 | יתכן |  |
| 6:11.95 | שיעזבו |  |
| 6:11.96 | ראשית |  |
| 6:11.97 | המצוה |  |
| 6:11.98 | ויסוד |  |
| 6:11.99 | האמונה |  |
| 6:11.100 | האלהית |  |
| 6:11.101 | בלי |  |
| 6:11.102 | ביאור | ! |
| 6:11.103 | אך |  |
| 6:11.104 | עתה |  |
| 6:11.105 | שהביאור |  |
| 6:11.106 | ההוא |  |
| 6:11.107 | נשכח |  |
| 6:11.108 | מאתנו |  |
| 6:11.109 | עלינו |  |
| 6:11.110 | לקחת |  |
| 6:11.111 | תורה |  |
| 6:11.112 | מפי |  |
| 6:11.113 | הפילסופים |  |
| 6:11.114 | התורניים | ) |
| 6:11.115 | למען |  |
| 6:11.116 | לא |  |
| 6:11.117 | נתעה |  |
| 6:11.118 | בלי |  |
| 6:11.119 | דעת |  |
| 6:11.120 | להאמין |  |
| 6:11.121 | בשניות |  |
| 6:11.122 | ובגשמות |  |
| 6:11.123 | כהמקובלים |  |
| 6:11.124 | החדשים | ) |
| 6:11.125 | ונעבוד |  |
| 6:11.126 | ללא |  |
| 6:11.127 | אלהי |  |
| 6:11.128 | אמת | . |
| 6:11.129 | והוא |  |
| 6:11.130 | הדבר |  |
| 6:11.131 | שאמר |  |
| 6:11.132 | בעל |  |
| 6:11.133 | החובות |  |
| 6:11.134 | ז"ל |  |
| 6:11.135 | לא |  |
| 6:11.136 | יוכל |  |
| 6:11.137 | לעבוד |  |
| 6:11.138 | עלת |  |
| 6:11.139 | העלות |  |
| 6:11.140 | אלא |  |
| 6:11.141 | נביא |  |
| 6:11.142 | הדור |  |
| 6:11.143 | או |  |
| 6:11.144 | החכם |  |
| 6:11.145 | המובהק |  |
| 6:11.146 | במה |  |
| 6:11.147 | שקנהו |  |
| 6:11.148 | מן |  |
| 6:11.149 | החכמה | . |
| 6:11.150 | וזולתם |  |
| 6:11.151 | עובדים |  |
| 6:11.152 | זולתו | . |
| 6:11.153 | ובאמת |  |
| 6:11.154 | דבר |  |
| 6:11.155 | החכם |  |
| 6:11.156 | כי |  |
| 6:11.157 | כמה |  |
| 6:11.158 | דברים |  |
| 6:11.159 | מורכבים |  |
| 6:11.160 | שנים |  |
| 6:11.161 | יתוארו |  |
| 6:11.162 | בשם |  |
| 6:11.163 | אחד | . |
| 6:11.164 | על |  |
| 6:11.165 | דרך |  |
| 6:11.166 | ההסכמה |  |
| 6:11.167 | ע"כ | . |
| 6:11.168 | ביאור |  |
| 6:11.169 | דבריו |  |
| 6:11.170 | כגון |  |
| 6:11.171 | שם |  |
| 6:11.172 | כירים |  |
| 6:11.173 | שהוא |  |
| 6:11.174 | מקום |  |
| 6:11.175 | שפיתת |  |
| 6:11.176 | שתי |  |
| 6:11.177 | קדרות | , |
| 6:11.178 | ונקרא |  |
| 6:11.179 | בדברי |  |
| 6:11.180 | רבותינו |  |
| 6:11.181 | בשם |  |
| 6:11.182 | אחד |  |
| 6:11.183 | כירה | . |
| 6:11.184 | וכן |  |
| 6:11.185 | רחים |  |
| 6:11.186 | שהוא |  |
| 6:11.187 | שם |  |
| 6:11.188 | השכב |  |
| 6:11.189 | העליון |  |
| 6:11.190 | רכב | , |
| 6:11.191 | ונקראו |  |
| 6:11.192 | שניהם |  |
| 6:11.193 | בשם |  |
| 6:11.194 | אחד |  |
| 6:11.195 | טחון | . |
| 6:11.196 | בחורים |  |
| 6:11.197 | טחון |  |
| 6:11.198 | נשאו |  |
| 6:11.199 | ובערבי |  |
| 6:11.200 | רחא |  |
| 6:11.201 | אע"פ |  |
| 6:11.202 | שהם |  |
| 6:11.203 | זוג | . |
| 6:11.204 | וכן |  |
| 6:11.205 | מספרים |  |
| 6:11.206 | נקרא |  |
| 6:11.207 | בערבי |  |
| 6:11.208 | מקץ | . |
| 6:11.209 | וכן |  |
| 6:11.210 | כל |  |
| 6:11.211 | כיוצא |  |
| 6:11.212 | בהם |  |
| 6:11.213 | כדתנן |  |
| 6:11.214 | בפרק |  |
| 6:11.215 | י"ג |  |
| 6:11.216 | ממסכת |  |
| 6:11.217 | כלים |  |
| 6:11.218 | קלגריפון | , |
| 6:11.219 | ומכחול | , |
| 6:11.220 | ומכתב | , |
| 6:11.221 | וזומא |  |
| 6:11.222 | ליסטרא | , |
| 6:11.223 | אע"פ |  |
| 6:11.224 | שהם |  |
| 6:11.225 | שנים |  |
| 6:11.226 | מחוברים |  |
| 6:11.227 | ומורכבים |  |
| 6:11.228 | נקראים |  |
| 6:11.229 | בשם |  |
| 6:11.230 | אחד |  |
| 6:11.231 | על |  |
| 6:11.232 | דרך |  |
| 6:11.233 | ההסכמה |  |
| 6:11.234 | וההעברה | . |
| 6:11.235 | ולא |  |
| 6:11.236 | כן |  |
| 6:11.237 | היא |  |
| 6:11.238 | אחדותו |  |
| 6:11.239 | ית | ' |
| 6:11.240 | שהוא |  |
| 6:11.241 | אחד |  |
| 6:11.242 | אמיתי |  |
| 6:11.243 | שאין |  |
| 6:11.244 | בו |  |
| 6:11.245 | שום |  |
| 6:11.246 | הרכבה |  |
| 6:11.247 | וצירוף | , |
| 6:11.248 | שלא |  |
| 6:11.249 | כדעת |  |
| 6:11.250 | המקובלים |  |
| 6:11.251 | החדשים |  |
| 6:11.252 | שאומרים |  |
| 6:11.253 | שיש |  |
| 6:11.254 | לו |  |
| 6:11.255 | כמה |  |
| 6:11.256 | פרצופים |  |
| 6:11.257 | שנאצלו |  |
| 6:11.258 | ממנו | . |
| 6:11.259 | וכל |  |
| 6:11.260 | אחד |  |
| 6:11.261 | נקרא |  |
| 6:11.262 | בשם |  |
| 6:11.263 | בפני |  |
| 6:11.264 | עצמו | , |
| 6:11.265 | והם |  |
| 6:11.266 | מחברים |  |
| 6:11.267 | ומשתפים |  |
| 6:11.268 | אותם |  |
| 6:11.269 | בפיהם |  |
| 6:11.270 | ובדעתם |  |
| 6:11.271 | עם |  |
| 6:11.272 | האלוה |  |
| 6:11.273 | המאצילם |  |
| 6:11.274 | ואומרים |  |
| 6:11.275 | וכלא |  |
| 6:11.276 | חד | . |
| 6:11.277 | ועובדים |  |
| 6:11.278 | אותם |  |
| 6:11.279 | יחד |  |
| 6:11.280 | גם |  |
| 6:11.281 | הגוף |  |
| 6:11.282 | גם |  |
| 6:11.283 | הנשמה | , |
| 6:11.284 | ומשקרים |  |
| 6:11.285 | לומר |  |
| 6:11.286 | כי |  |
| 6:11.287 | את |  |
| 6:11.288 | נשמתו |  |
| 6:11.289 | הם |  |
| 6:11.290 | עובדים | , |
| 6:11.291 | וזה |  |
| 6:11.292 | לשון |  |
| 6:11.293 | עוז |  |
| 6:11.294 | לאלהים |  |
| 6:11.295 | בבית |  |
| 6:11.296 | קדש |  |
| 6:11.297 | הקדשים |  |
| 6:11.298 | פ | ' |
| 6:11.299 | י"ח |  |
| 6:11.300 | דף |  |
| 6:11.301 | נ"ו |  |
| 6:11.302 | ע"ב |  |
| 6:11.303 | ואם |  |
| 6:11.304 | תאמר |  |
| 6:11.305 | שאהיה |  |
| 6:11.306 | עובד |  |
| 6:11.307 | לנשמה |  |
| 6:11.308 | בלא |  |
| 6:11.309 | גופא |  |
| 6:11.310 | ח"ו | , |
| 6:11.311 | חלילה | , |
| 6:11.312 | אלא |  |
| 6:11.313 | כבר |  |
| 6:11.314 | אמרנו |  |
| 6:11.315 | לעיל |  |
| 6:11.316 | שזה |  |
| 6:11.317 | הפרצוף |  |
| 6:11.318 | דמלכא |  |
| 6:11.319 | קדישא |  |
| 6:11.320 | ושכינתיה |  |
| 6:11.321 | הוא |  |
| 6:11.322 | אדוק |  |
| 6:11.323 | בנשמה | . |
| 6:11.324 | לפי |  |
| 6:11.325 | שהוא |  |
| 6:11.326 | עתיקא |  |
| 6:11.327 | קדישא |  |
| 6:11.328 | דכל |  |
| 6:11.329 | קדישין | . |
| 6:11.330 | והוא |  |
| 6:11.331 | ונשמתא |  |
| 6:11.332 | חד |  |
| 6:11.333 | אינון |  |
| 6:11.334 | כמא |  |
| 6:11.335 | דאמרינן |  |
| 6:11.336 | לעיל |  |
| 6:11.337 | הוא |  |
| 6:11.338 | ושמיה |  |
| 6:11.339 | חד |  |
| 6:11.340 | אינון |  |
| 6:11.341 | ולא |  |
| 6:11.342 | מתפרש |  |
| 6:11.343 | מנהון |  |
| 6:11.344 | לעלמין | . |
| 6:11.345 | ואם |  |
| 6:11.346 | יעבוד |  |
| 6:11.347 | לנשמה |  |
| 6:11.348 | בלי |  |
| 6:11.349 | מלכא |  |
| 6:11.350 | קדישא |  |
| 6:11.351 | יהיה |  |
| 6:11.352 | פירוד |  |
| 6:11.353 | בין |  |
| 6:11.354 | נשמה |  |
| 6:11.355 | לגופא | . |
| 6:11.356 | וכן |  |
| 6:11.357 | מבואר |  |
| 6:11.358 | מדברי |  |
| 6:11.359 | גליון |  |
| 6:11.360 | הזהר |  |
| 6:11.361 | ומשנת |  |
| 6:11.362 | חסידים |  |
| 6:11.363 | הובא |  |
| 6:11.364 | לעיל |  |
| 6:11.365 | סי | ' |
| 6:11.366 | ט"ל | . |
| 6:11.367 | ועל |  |
| 6:11.368 | מה |  |
| 6:11.369 | שכ | ' |
| 6:11.370 | בעל |  |
| 6:11.371 | הגדרים |  |
| 6:11.372 | שלא |  |
| 6:11.373 | יתכן |  |
| 6:11.374 | שיעזבו |  |
| 6:11.375 | רז"ל |  |
| 6:11.376 | ראשית |  |
| 6:11.377 | המצוה |  |
| 6:11.378 | ויסודה |  |
| 6:11.379 | וכו | ' |
| 6:11.380 | כבר |  |
| 6:11.381 | העירותי |  |
| 6:11.382 | על |  |
| 6:11.383 | דבריו |  |
| 6:11.384 | לעיל |  |
| 6:11.385 | סי | ' |
| 6:11.386 | ס"ד |  |
| 6:11.387 | שאדרבה |  |
| 6:11.388 | נהפוך |  |
| 6:11.389 | הוא |  |
| 6:11.390 | שכל |  |
| 6:11.391 | דבר |  |
| 6:11.392 | המפורש |  |
| 6:11.393 | בכתובים |  |
| 6:11.394 | והוא |  |
| 6:11.395 | פשוט |  |
| 6:11.396 | ומפורסם |  |
| 6:11.397 | בישראל |  |
| 6:11.398 | לא |  |
| 6:11.399 | ידברו |  |
| 6:11.400 | בו |  |
| 6:11.401 | חכמי |  |
| 6:11.402 | המשנה | . |
| 6:11.403 | ובעלי |  |
| 6:11.404 | התלמוד |  |
| 6:11.405 | ג"כ |  |
| 6:11.406 | לא |  |
| 6:11.407 | שמו |  |
| 6:11.408 | מגמתם |  |
| 6:11.409 | רק |  |
| 6:11.410 | לבאר |  |
| 6:11.411 | ולישא |  |
| 6:11.412 | וליתן |  |
| 6:11.413 | בדבריה |  |
| 6:11.414 | ולבררה |  |
| 6:11.415 | וללבנה | , |
| 6:11.416 | ולהודיע |  |
| 6:11.417 | דינים |  |
| 6:11.418 | ושאלות |  |
| 6:11.419 | שנתחדשו |  |
| 6:11.420 | בימיהם | , |
| 6:11.421 | וגזרות |  |
| 6:11.422 | ותקנות |  |
| 6:11.423 | שתיקנו | , |
| 6:11.424 | אבל |  |
| 6:11.425 | דברים |  |
| 6:11.426 | הפשוטים |  |
| 6:11.427 | אין |  |
| 6:11.428 | צורך |  |
| 6:11.429 | להזכירם |  |
| 6:11.430 | כלל | . |
| 6:11.431 | ודעות |  |
| 6:11.432 | ורעיונות |  |
| 6:11.433 | הקבלה |  |
| 6:11.434 | החדשה |  |
| 6:11.435 | עוד |  |
| 6:11.436 | לא |  |
| 6:11.437 | נודעו |  |
| 6:11.438 | ולא |  |
| 6:11.439 | נשמעו |  |
| 6:11.440 | בימיהם |  |
| 6:11.441 | בין |  |
| 6:11.442 | היהודים |  |
| 6:11.443 | כלל | , |
| 6:11.444 | וזרים |  |
| 6:11.445 | המה |  |
| 6:11.446 | לאמונת |  |
| 6:11.447 | תורתינו |  |
| 6:11.448 | הקדושה | . |
| 6:11.449 | כי |  |
| 6:11.450 | רק |  |
| 6:11.451 | מתחלת |  |
| 6:11.452 | האלף |  |
| 6:11.453 | הששי |  |
| 6:11.454 | בסוף |  |
| 6:11.455 | המאה |  |
| 6:11.456 | הראשונה |  |
| 6:11.457 | החזיקו |  |
| 6:11.458 | ברעיונות |  |
| 6:11.459 | האלה |  |
| 6:11.460 | קצת |  |
| 6:11.461 | מחברים |  |
| 6:11.462 | שנמשכו |  |
| 6:11.463 | אחר |  |
| 6:11.464 | דעות |  |
| 6:11.465 | הפילסוף |  |
| 6:11.466 | הרמאי |  |
| 6:11.467 | שדבק |  |
| 6:11.468 | בדעת |  |
| 6:11.469 | השונים |  |
| 6:11.470 | והשלישי |  |
| 6:11.471 | הוא |  |
| 6:11.472 | מחבר |  |
| 6:11.473 | ספר |  |
| 6:11.474 | הזהר |  |
| 6:11.475 | והתקונים | . |
| 6:11.476 | וכוון |  |
| 6:11.477 | בו |  |
| 6:11.478 | להטעות |  |
| 6:11.479 | את |  |
| 6:11.480 | ישראל |  |
| 6:11.481 | ולהדיחם |  |
| 6:11.482 | מעל |  |
| 6:11.483 | ה | ' |
| 6:11.484 | אלהינו | . |
| 6:11.485 | ותלה |  |
| 6:11.486 | רעיונותיו |  |
| 6:11.487 | ברשב"י |  |
| 6:11.488 | התנא |  |
| 6:11.489 | וחביריו | , |
| 6:11.490 | בפרט |  |
| 6:11.491 | בשני |  |
| 6:11.492 | דברים | , |
| 6:11.493 | מידי |  |
| 6:11.494 | דעל |  |
| 6:11.495 | רשב"י |  |
| 6:11.496 | תני |  |
| 6:11.497 | במאמרו |  |
| 6:11.498 | הקצר |  |
| 6:11.499 | כל |  |
| 6:11.500 | המשתף |  |
| 6:11.501 | שם |  |
| 6:11.502 | שמים |  |
| 6:11.503 | ודבר |  |
| 6:11.504 | אחר |  |
| 6:11.505 | נעקר |  |
| 6:11.506 | מן |  |
| 6:11.507 | העולם | . |
| 6:11.508 | והשני |  |
| 6:11.509 | בביזוי |  |
| 6:11.510 | המשנה |  |
| 6:11.511 | שקלא |  |
| 6:11.512 | וטריא |  |
| 6:11.513 | דבתלמודא |  |
| 6:11.514 | שהם |  |
| 6:11.515 | עיקר |  |
| 6:11.516 | ויסוד |  |
| 6:11.517 | תורה |  |
| 6:11.518 | שבעל |  |
| 6:11.519 | פה |  |
| 6:11.520 | שהוא |  |
| 6:11.521 | מבזה |  |
| 6:11.522 | אותן |  |
| 6:11.523 | כמו |  |
| 6:11.524 | שיתבאר |  |
| 6:11.525 | להלן |  |
| 6:11.526 | דלרשב"י |  |
| 6:11.527 | שתורתו |  |
| 6:11.528 | אומנותו |  |
| 6:11.529 | היא |  |
| 6:11.530 | כקריאת |  |
| 6:11.531 | שמע |  |
| 6:11.532 | בעונתה |  |
| 6:11.533 | ירוש | ' |
| 6:11.534 | שבת |  |
| 6:11.535 | פ"א | ). |
| 6:11.536 | וטמן |  |
| 6:11.537 | את |  |
| 6:11.538 | הספר |  |
| 6:11.539 | הנז | ' |
| 6:11.540 | בארון |  |
| 6:11.541 | וקברו |  |
| 6:11.542 | בקרקע |  |
| 6:11.543 | כדי |  |
| 6:11.544 | לישנו |  |
| 6:11.545 | כאלו |  |
| 6:11.546 | הוא |  |
| 6:11.547 | ישן |  |
| 6:11.548 | נושן |  |
| 6:11.549 | מימי |  |
| 6:11.550 | רשב"י | . |
| 6:11.551 | ואמר |  |
| 6:11.552 | למלך |  |
| 6:11.553 | ממלכי |  |
| 6:11.554 | מזרח |  |
| 6:11.555 | לחפור |  |
| 6:11.556 | שם |  |
| 6:11.557 | בקרקע |  |
| 6:11.558 | למצוא |  |
| 6:11.559 | מטמון |  |
| 6:11.560 | של |  |
| 6:11.561 | כסף |  |
| 6:11.562 | וזהב | , |
| 6:11.563 | וצוה |  |
| 6:11.564 | המלך |  |
| 6:11.565 | הנז | ' |
| 6:11.566 | לחפור |  |
| 6:11.567 | על |  |
| 6:11.568 | עסקי |  |
| 6:11.569 | ממון |  |
| 6:11.570 | ונמצא |  |
| 6:11.571 | שם |  |
| 6:11.572 | ארון |  |
| 6:11.573 | ובו |  |
| 6:11.574 | ס | ' |
| 6:11.575 | הזהר |  |
| 6:11.576 | וכו | ' |
| 6:11.577 | ושלח |  |
| 6:11.578 | אותו |  |
| 6:11.579 | לטוליטולא |  |
| 6:11.580 | וכו | ' |
| 6:11.581 | ומשם |  |
| 6:11.582 | נתפשטה |  |
| 6:11.583 | הקבלה |  |
| 6:11.584 | החדשה |  |
| 6:11.585 | בישראל |  |
| 6:11.586 | כמו |  |
| 6:11.587 | שכ | ' |
| 6:11.588 | הר | ' |
| 6:11.589 | חיד"א |  |
| 6:11.590 | בספרו |  |
| 6:11.591 | הנחמד |  |
| 6:11.592 | שם |  |
| 6:11.593 | הגדולים |  |
| 6:11.594 | במערכת |  |
| 6:11.595 | ספרים |  |
| 6:11.596 | ערך |  |
| 6:11.597 | זהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

