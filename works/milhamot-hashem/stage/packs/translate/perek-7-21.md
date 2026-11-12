# Translation pack: milhamot-hashem perek-7-21 (Milchamot ha-Shem 7:26–7:26)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-21`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
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
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרמב"ן / רמב"ן** → "Ramban" (authority, transliterated; note first in a
  unit).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ראב"ד** → "Ravad" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בא"י** → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **עלה / עלול** → "cause / effect".
- אורות **דקים** → "subtle lights".
- **זווג** → "coupling".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- בת **קול** → "bat kol" (note: "daughter of a voice").
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין** אל → "call upon".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **צנומות** דקות שדופות קדים → "withered, thin, blasted by the east wind".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **בנדון** דידן → "in our case".
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 7:26

Hebrew (Wikisource): צח) והנה ידוע ומפורסם הוא לכל בקי בספרות ישראל. שחבור הזהר ושאר ספרי קבלה הנכריה החדשה לא נתפשטו בישראל כי מתחלת האלף השישי, ולא נודע בבירור מהיכן צמחו, כמו שאביא להלן דעות המאמתים אותם. ולא ידעו לברר הדבר, אך הקרוב אל הדעת הוא מה שכתב הרב חיד"א בס' שם הגדולים מערכת ספרים אות זי"ן וכבר הבאתי דבריו לעיל. וחכמי ירושלם באגדותיהם מסכימים לדבריו יע"ש. כי מיד המלך הנכרי זכו לאמונה. ובגליון הבא מירושלם ת"ו האריכו מאד להעתיק בקצור את דברי האחרונים הטחים תפל על האמונות הנכריות לרפאם ולא נרפאו וכל מגמת חכמי ירושת"ו בכל מה שהעתיקו בגליון מדברי האחרונים. הוא לשלול הגשמות מהספירות והצורות והפרצופים הנזכרים בזהר ושאר ספרי קבלה! ואם הצילונו מן הנשמות כלום הצילונו מעבודת אל נר בין גשמי בין רוחני? כי מודים הם שהם סבות נבראות, וכל אחד סבה ועלה למה שלמטה ממנה זכרים ונקבות, כמבואר בזהר וספרים הנמשכים אחריו, רק שהם כוחות רוחניות של אור נברא ואותם יעבודו, ולהם יקראו בעת צרתם ולהם ישבחו ויודו ויהללו בעת ישועתם. ומקבלים עליהם לאלוהות כאלו הוא מותר לנו מדין תורתינו הקדושה לעבוד כח נברא, רק שלא נחשוב אותם גשמים אלא רוחנים. כי נסלק מהם ונשלול הנשמות, ונחשוב בדמיונינו שהם אורות. (ולא ידעו שגם האור והאויר הם גופים דקים סי' מ"ז) ומביאים ראיות צנומות דקות שדופות קדים לשלול הגשמות מהם. ולחשוב את המשפיע זכר ואת מקבל השפע נקבה, ואת הזווגים שלא כפשטן. רק קבלת השפע קרו ליה זווג! ועם כל זה לא הביאו לנו פתרונים לשמות איברים מגונים כיסוד הוא גיד הערוה, וביעין דדכורא, ויסוד דנוקבא ונקב האוחרים. למאי חזו באור הזה המופשט מן הגשמות? ואם התורה הקדושה והנביאים נשתמשו בלשונות המורות על הגשמות כיד ואצבעות. כלום השתמשו ליחס לשי"ת חושים הפחותים באכילה ושתייה ומשוש אף בחלק המעולה והחשוב שבהם. כ"ש בחלק המגונה והפחות שבחוש המשוש שהוא המשגל. כמו שכתב הרב המורה בפרק מ"ז מראשון. והיותר רע ומר הוא מה שישבו סבות רבות באלהות כולן ברואות. והמה אצלם עיקר האלהות. ויחברו ספרים בסדר מערכותן כעורך ומסדר את החלות על השלחן. ויקראו לכל אחד מן הפרצופים בשם הוי"ה ואדנות ואלהים צבאות, ויקבלו עליהן אלהותן כמו שכתב ספר הברית וכסא אליהו ונחלת יוסף לעבוד דוקא לאלוה קצר אפים (זעיר אנפין) ולא לאלוה העליון הבורא את כל הפרצופים והצורות. וכתב עוד ספר הברית שהעובד לאלוה העליון על כולם לא יאבא קצר האפים הנקרא אצלם בשם הוי"ה והקב"ה לסלוח לו, כי אז יעשן אפו וקנאתו, באיש ההוא וכו'. כי רק לו לבדו ניתנה הממשלה מהאלהים העליונים ממנו שהסכימו להמליכו, למרות דבר ה' אשר דבר על יד נביאו וכבודי לאחר לא אתן. ודבריהם אלו מבוארים בזהר בלק דף קצ"א עם פירוש מקדש מלך. ובפרט בספר עז לאלהים שהיא משל על זה ממלך ב"ו שמלך שגם אביו ואימו צריכים לכבדו, וזה המשל מתנגד למשל שהביאו רז"ל בשמות רבה על אנכי ה' אלהיך. ובספר שלשלת הקבלה להראב"ד כתב משם ס' יוחסין וז"ל בשנת כמו חמשת אלפים וחמשים שנה נמצאו כתות אנשים שאמרו כי דברי הזהר אשר הם בלשון ירושלמי הם דברי רשב"י. אבל אשר הם בלשון הקדש אינן דבריו, ויש אומרים שהרמב"ן מצאו בארץ ישראל והביא אותו לקטלוניאה, ואחר לארגון. ונפל ליד רבי משה דיליואן. ויש אומרים כי ר' משה דיליואן הנז' היה חכם גדול ועושה אלו הפירושים משכלו, ולמען יקח מהם מחיר רב מהחכמים היה כותבם ותולה אותם ברשב"י וחבורתו, ע"כ בספר יוחסין ומאריך עוד יותר עכ"ל שלשלת הקבלה. וכבר הבאתי לעיל סי' צ"ד דברי הרב חיד"א, בס' שם הגדולים וגליון שהובא מירושלם ת"י על ידי מלשינות הקשר אשר הרה וילד ארך הלשון. והנה שכחו דין פשוט בשועיורה דעה סי' רפ"א שספר תורה הנמצא ביד גוי דלד' הרי"ף והרא"ש והטור יגנז ואין קורין בו אע"פ שאין בו שום חסרון או יתרון משאר ספרי תורה, ולדעת הרמב"ם ז"ל כיון שאין ידוע מי כתבו ואין בו שום שינוי חזקתו שישראל כתבו וקורין בו, הא אם נמצא בו שום שינוי אפילו בחסר ויתר דבר פשוט הוא דלכ"ע אין קורין בו ובנדון דידן של ספר הזהר יש לדון בו קל וחומר, ומה אם זה שאין בו שום קלקול לאמונתינו יגנז? ספר המבזה את קבלתינו המשנה והתלמוד ומסית ומדיח אותנו לעבוד כחות נבראים. קצר אפים עם הוריו ויולדיו, על אחת כמה וכמה שאין להאמין לו להמיר את אלהינו ברוך הוא וב"ש ולסתור פנות תורתינו הקדושה. ומדבריהם מודאע רבה לאוריתא שהתנאי שאמר הקמב"ן שא"א להבין דברי הקבלה מתוך הספר רק מפי חכם מקובל לאזן מקבל, לא נתקיים מעולם תנאי זה, שהרי הם בעצמן מודים שאחרי מות רשב"י ובנו ר"א וכל הדור ההוא נגנז הזהר ונשכחה ואבדה הקבלה הלזו ולא נשאר אדם יודע בה עד שנתגלה ספר הזהר באלף השישי ע"י המלך העכו"ם. וממנו נתפשטה הקבלה ולא מפי חכם מקובל אשר קבל מפה לאזן. כי כל המקובלים הרבים החדשים שבעולם לא שאבו רק מהזהר וממנו נתפשטה, ולפי דברי הרמב"ן זהו הגורם שכולם שגגו וטעו ועקרו יסודי תורתינו הקדושה והאמינו באלהים רבים. ובאו לידי שתוף שם שמים ודבר אחד בעבור מיעוט הבנתם. אף לפי האמת כיון שגילה דעתו בביאור שהוא מבזה את המשנה והתלמוד ומדמה שקלא וטריא דרבנן המתאבקים בתלמוד לכלבין דצוחין הב, הב, אין מקום לכל זה (עיין תקוני זהר חדש). ולפי דברי הרב שם הגדולים והגליון המובא מירושת"ו שר' נחוניא בן הקנה היה ראש המקובלים אשר גילו את הקבלה החדשה ולמדו אותה בגלוי! עד שתיקן בכניסתו לבית המדרש להתפלל שלא אכשל בדבר הלכה, שלא אומר על טמא טהור וכו' כדתנן בברכות פ' תפלת השחר. היה לו להתפלל שלא תבוא תקלה גדולה כזו על ידו להאמין ברשויות רבות ולעבוד אל אחר קצר אפים (זעיר אנפין)? לפי שעונש המחטיא אחרים להאמין ברשויות רבות ולעבוד אלהי נכר יותר גדול מן הנכשל להתיר את האסור או לאסור את המותר?

| id | word | punctuation after |
|---|---|---|
| 7:26.1 | צח | ) |
| 7:26.2 | והנה |  |
| 7:26.3 | ידוע |  |
| 7:26.4 | ומפורסם |  |
| 7:26.5 | הוא |  |
| 7:26.6 | לכל |  |
| 7:26.7 | בקי |  |
| 7:26.8 | בספרות |  |
| 7:26.9 | ישראל | . |
| 7:26.10 | שחבור |  |
| 7:26.11 | הזהר |  |
| 7:26.12 | ושאר |  |
| 7:26.13 | ספרי |  |
| 7:26.14 | קבלה |  |
| 7:26.15 | הנכריה |  |
| 7:26.16 | החדשה |  |
| 7:26.17 | לא |  |
| 7:26.18 | נתפשטו |  |
| 7:26.19 | בישראל |  |
| 7:26.20 | כי |  |
| 7:26.21 | מתחלת |  |
| 7:26.22 | האלף |  |
| 7:26.23 | השישי | , |
| 7:26.24 | ולא |  |
| 7:26.25 | נודע |  |
| 7:26.26 | בבירור |  |
| 7:26.27 | מהיכן |  |
| 7:26.28 | צמחו | , |
| 7:26.29 | כמו |  |
| 7:26.30 | שאביא |  |
| 7:26.31 | להלן |  |
| 7:26.32 | דעות |  |
| 7:26.33 | המאמתים |  |
| 7:26.34 | אותם | . |
| 7:26.35 | ולא |  |
| 7:26.36 | ידעו |  |
| 7:26.37 | לברר |  |
| 7:26.38 | הדבר | , |
| 7:26.39 | אך |  |
| 7:26.40 | הקרוב |  |
| 7:26.41 | אל |  |
| 7:26.42 | הדעת |  |
| 7:26.43 | הוא |  |
| 7:26.44 | מה |  |
| 7:26.45 | שכתב |  |
| 7:26.46 | הרב |  |
| 7:26.47 | חיד"א |  |
| 7:26.48 | בס | ' |
| 7:26.49 | שם |  |
| 7:26.50 | הגדולים |  |
| 7:26.51 | מערכת |  |
| 7:26.52 | ספרים |  |
| 7:26.53 | אות |  |
| 7:26.54 | זי"ן |  |
| 7:26.55 | וכבר |  |
| 7:26.56 | הבאתי |  |
| 7:26.57 | דבריו |  |
| 7:26.58 | לעיל | . |
| 7:26.59 | וחכמי |  |
| 7:26.60 | ירושלם |  |
| 7:26.61 | באגדותיהם |  |
| 7:26.62 | מסכימים |  |
| 7:26.63 | לדבריו |  |
| 7:26.64 | יע"ש | . |
| 7:26.65 | כי |  |
| 7:26.66 | מיד |  |
| 7:26.67 | המלך |  |
| 7:26.68 | הנכרי |  |
| 7:26.69 | זכו |  |
| 7:26.70 | לאמונה | . |
| 7:26.71 | ובגליון |  |
| 7:26.72 | הבא |  |
| 7:26.73 | מירושלם |  |
| 7:26.74 | ת"ו |  |
| 7:26.75 | האריכו |  |
| 7:26.76 | מאד |  |
| 7:26.77 | להעתיק |  |
| 7:26.78 | בקצור |  |
| 7:26.79 | את |  |
| 7:26.80 | דברי |  |
| 7:26.81 | האחרונים |  |
| 7:26.82 | הטחים |  |
| 7:26.83 | תפל |  |
| 7:26.84 | על |  |
| 7:26.85 | האמונות |  |
| 7:26.86 | הנכריות |  |
| 7:26.87 | לרפאם |  |
| 7:26.88 | ולא |  |
| 7:26.89 | נרפאו |  |
| 7:26.90 | וכל |  |
| 7:26.91 | מגמת |  |
| 7:26.92 | חכמי |  |
| 7:26.93 | ירושת"ו |  |
| 7:26.94 | בכל |  |
| 7:26.95 | מה |  |
| 7:26.96 | שהעתיקו |  |
| 7:26.97 | בגליון |  |
| 7:26.98 | מדברי |  |
| 7:26.99 | האחרונים | . |
| 7:26.100 | הוא |  |
| 7:26.101 | לשלול |  |
| 7:26.102 | הגשמות |  |
| 7:26.103 | מהספירות |  |
| 7:26.104 | והצורות |  |
| 7:26.105 | והפרצופים |  |
| 7:26.106 | הנזכרים |  |
| 7:26.107 | בזהר |  |
| 7:26.108 | ושאר |  |
| 7:26.109 | ספרי |  |
| 7:26.110 | קבלה | ! |
| 7:26.111 | ואם |  |
| 7:26.112 | הצילונו |  |
| 7:26.113 | מן |  |
| 7:26.114 | הנשמות |  |
| 7:26.115 | כלום |  |
| 7:26.116 | הצילונו |  |
| 7:26.117 | מעבודת |  |
| 7:26.118 | אל |  |
| 7:26.119 | נר |  |
| 7:26.120 | בין |  |
| 7:26.121 | גשמי |  |
| 7:26.122 | בין |  |
| 7:26.123 | רוחני | ? |
| 7:26.124 | כי |  |
| 7:26.125 | מודים |  |
| 7:26.126 | הם |  |
| 7:26.127 | שהם |  |
| 7:26.128 | סבות |  |
| 7:26.129 | נבראות | , |
| 7:26.130 | וכל |  |
| 7:26.131 | אחד |  |
| 7:26.132 | סבה |  |
| 7:26.133 | ועלה |  |
| 7:26.134 | למה |  |
| 7:26.135 | שלמטה |  |
| 7:26.136 | ממנה |  |
| 7:26.137 | זכרים |  |
| 7:26.138 | ונקבות | , |
| 7:26.139 | כמבואר |  |
| 7:26.140 | בזהר |  |
| 7:26.141 | וספרים |  |
| 7:26.142 | הנמשכים |  |
| 7:26.143 | אחריו | , |
| 7:26.144 | רק |  |
| 7:26.145 | שהם |  |
| 7:26.146 | כוחות |  |
| 7:26.147 | רוחניות |  |
| 7:26.148 | של |  |
| 7:26.149 | אור |  |
| 7:26.150 | נברא |  |
| 7:26.151 | ואותם |  |
| 7:26.152 | יעבודו | , |
| 7:26.153 | ולהם |  |
| 7:26.154 | יקראו |  |
| 7:26.155 | בעת |  |
| 7:26.156 | צרתם |  |
| 7:26.157 | ולהם |  |
| 7:26.158 | ישבחו |  |
| 7:26.159 | ויודו |  |
| 7:26.160 | ויהללו |  |
| 7:26.161 | בעת |  |
| 7:26.162 | ישועתם | . |
| 7:26.163 | ומקבלים |  |
| 7:26.164 | עליהם |  |
| 7:26.165 | לאלוהות |  |
| 7:26.166 | כאלו |  |
| 7:26.167 | הוא |  |
| 7:26.168 | מותר |  |
| 7:26.169 | לנו |  |
| 7:26.170 | מדין |  |
| 7:26.171 | תורתינו |  |
| 7:26.172 | הקדושה |  |
| 7:26.173 | לעבוד |  |
| 7:26.174 | כח |  |
| 7:26.175 | נברא | , |
| 7:26.176 | רק |  |
| 7:26.177 | שלא |  |
| 7:26.178 | נחשוב |  |
| 7:26.179 | אותם |  |
| 7:26.180 | גשמים |  |
| 7:26.181 | אלא |  |
| 7:26.182 | רוחנים | . |
| 7:26.183 | כי |  |
| 7:26.184 | נסלק |  |
| 7:26.185 | מהם |  |
| 7:26.186 | ונשלול |  |
| 7:26.187 | הנשמות | , |
| 7:26.188 | ונחשוב |  |
| 7:26.189 | בדמיונינו |  |
| 7:26.190 | שהם |  |
| 7:26.191 | אורות | . |
| 7:26.192 | ולא |  |
| 7:26.193 | ידעו |  |
| 7:26.194 | שגם |  |
| 7:26.195 | האור |  |
| 7:26.196 | והאויר |  |
| 7:26.197 | הם |  |
| 7:26.198 | גופים |  |
| 7:26.199 | דקים |  |
| 7:26.200 | סי | ' |
| 7:26.201 | מ"ז | ) |
| 7:26.202 | ומביאים |  |
| 7:26.203 | ראיות |  |
| 7:26.204 | צנומות |  |
| 7:26.205 | דקות |  |
| 7:26.206 | שדופות |  |
| 7:26.207 | קדים |  |
| 7:26.208 | לשלול |  |
| 7:26.209 | הגשמות |  |
| 7:26.210 | מהם | . |
| 7:26.211 | ולחשוב |  |
| 7:26.212 | את |  |
| 7:26.213 | המשפיע |  |
| 7:26.214 | זכר |  |
| 7:26.215 | ואת |  |
| 7:26.216 | מקבל |  |
| 7:26.217 | השפע |  |
| 7:26.218 | נקבה | , |
| 7:26.219 | ואת |  |
| 7:26.220 | הזווגים |  |
| 7:26.221 | שלא |  |
| 7:26.222 | כפשטן | . |
| 7:26.223 | רק |  |
| 7:26.224 | קבלת |  |
| 7:26.225 | השפע |  |
| 7:26.226 | קרו |  |
| 7:26.227 | ליה |  |
| 7:26.228 | זווג | ! |
| 7:26.229 | ועם |  |
| 7:26.230 | כל |  |
| 7:26.231 | זה |  |
| 7:26.232 | לא |  |
| 7:26.233 | הביאו |  |
| 7:26.234 | לנו |  |
| 7:26.235 | פתרונים |  |
| 7:26.236 | לשמות |  |
| 7:26.237 | איברים |  |
| 7:26.238 | מגונים |  |
| 7:26.239 | כיסוד |  |
| 7:26.240 | הוא |  |
| 7:26.241 | גיד |  |
| 7:26.242 | הערוה | , |
| 7:26.243 | וביעין |  |
| 7:26.244 | דדכורא | , |
| 7:26.245 | ויסוד |  |
| 7:26.246 | דנוקבא |  |
| 7:26.247 | ונקב |  |
| 7:26.248 | האוחרים | . |
| 7:26.249 | למאי |  |
| 7:26.250 | חזו |  |
| 7:26.251 | באור |  |
| 7:26.252 | הזה |  |
| 7:26.253 | המופשט |  |
| 7:26.254 | מן |  |
| 7:26.255 | הגשמות | ? |
| 7:26.256 | ואם |  |
| 7:26.257 | התורה |  |
| 7:26.258 | הקדושה |  |
| 7:26.259 | והנביאים |  |
| 7:26.260 | נשתמשו |  |
| 7:26.261 | בלשונות |  |
| 7:26.262 | המורות |  |
| 7:26.263 | על |  |
| 7:26.264 | הגשמות |  |
| 7:26.265 | כיד |  |
| 7:26.266 | ואצבעות | . |
| 7:26.267 | כלום |  |
| 7:26.268 | השתמשו |  |
| 7:26.269 | ליחס |  |
| 7:26.270 | לשי"ת |  |
| 7:26.271 | חושים |  |
| 7:26.272 | הפחותים |  |
| 7:26.273 | באכילה |  |
| 7:26.274 | ושתייה |  |
| 7:26.275 | ומשוש |  |
| 7:26.276 | אף |  |
| 7:26.277 | בחלק |  |
| 7:26.278 | המעולה |  |
| 7:26.279 | והחשוב |  |
| 7:26.280 | שבהם | . |
| 7:26.281 | כ"ש |  |
| 7:26.282 | בחלק |  |
| 7:26.283 | המגונה |  |
| 7:26.284 | והפחות |  |
| 7:26.285 | שבחוש |  |
| 7:26.286 | המשוש |  |
| 7:26.287 | שהוא |  |
| 7:26.288 | המשגל | . |
| 7:26.289 | כמו |  |
| 7:26.290 | שכתב |  |
| 7:26.291 | הרב |  |
| 7:26.292 | המורה |  |
| 7:26.293 | בפרק |  |
| 7:26.294 | מ"ז |  |
| 7:26.295 | מראשון | . |
| 7:26.296 | והיותר |  |
| 7:26.297 | רע |  |
| 7:26.298 | ומר |  |
| 7:26.299 | הוא |  |
| 7:26.300 | מה |  |
| 7:26.301 | שישבו |  |
| 7:26.302 | סבות |  |
| 7:26.303 | רבות |  |
| 7:26.304 | באלהות |  |
| 7:26.305 | כולן |  |
| 7:26.306 | ברואות | . |
| 7:26.307 | והמה |  |
| 7:26.308 | אצלם |  |
| 7:26.309 | עיקר |  |
| 7:26.310 | האלהות | . |
| 7:26.311 | ויחברו |  |
| 7:26.312 | ספרים |  |
| 7:26.313 | בסדר |  |
| 7:26.314 | מערכותן |  |
| 7:26.315 | כעורך |  |
| 7:26.316 | ומסדר |  |
| 7:26.317 | את |  |
| 7:26.318 | החלות |  |
| 7:26.319 | על |  |
| 7:26.320 | השלחן | . |
| 7:26.321 | ויקראו |  |
| 7:26.322 | לכל |  |
| 7:26.323 | אחד |  |
| 7:26.324 | מן |  |
| 7:26.325 | הפרצופים |  |
| 7:26.326 | בשם |  |
| 7:26.327 | הוי"ה |  |
| 7:26.328 | ואדנות |  |
| 7:26.329 | ואלהים |  |
| 7:26.330 | צבאות | , |
| 7:26.331 | ויקבלו |  |
| 7:26.332 | עליהן |  |
| 7:26.333 | אלהותן |  |
| 7:26.334 | כמו |  |
| 7:26.335 | שכתב |  |
| 7:26.336 | ספר |  |
| 7:26.337 | הברית |  |
| 7:26.338 | וכסא |  |
| 7:26.339 | אליהו |  |
| 7:26.340 | ונחלת |  |
| 7:26.341 | יוסף |  |
| 7:26.342 | לעבוד |  |
| 7:26.343 | דוקא |  |
| 7:26.344 | לאלוה |  |
| 7:26.345 | קצר |  |
| 7:26.346 | אפים |  |
| 7:26.347 | זעיר |  |
| 7:26.348 | אנפין | ) |
| 7:26.349 | ולא |  |
| 7:26.350 | לאלוה |  |
| 7:26.351 | העליון |  |
| 7:26.352 | הבורא |  |
| 7:26.353 | את |  |
| 7:26.354 | כל |  |
| 7:26.355 | הפרצופים |  |
| 7:26.356 | והצורות | . |
| 7:26.357 | וכתב |  |
| 7:26.358 | עוד |  |
| 7:26.359 | ספר |  |
| 7:26.360 | הברית |  |
| 7:26.361 | שהעובד |  |
| 7:26.362 | לאלוה |  |
| 7:26.363 | העליון |  |
| 7:26.364 | על |  |
| 7:26.365 | כולם |  |
| 7:26.366 | לא |  |
| 7:26.367 | יאבא |  |
| 7:26.368 | קצר |  |
| 7:26.369 | האפים |  |
| 7:26.370 | הנקרא |  |
| 7:26.371 | אצלם |  |
| 7:26.372 | בשם |  |
| 7:26.373 | הוי"ה |  |
| 7:26.374 | והקב"ה |  |
| 7:26.375 | לסלוח |  |
| 7:26.376 | לו | , |
| 7:26.377 | כי |  |
| 7:26.378 | אז |  |
| 7:26.379 | יעשן |  |
| 7:26.380 | אפו |  |
| 7:26.381 | וקנאתו | , |
| 7:26.382 | באיש |  |
| 7:26.383 | ההוא |  |
| 7:26.384 | וכו | '. |
| 7:26.385 | כי |  |
| 7:26.386 | רק |  |
| 7:26.387 | לו |  |
| 7:26.388 | לבדו |  |
| 7:26.389 | ניתנה |  |
| 7:26.390 | הממשלה |  |
| 7:26.391 | מהאלהים |  |
| 7:26.392 | העליונים |  |
| 7:26.393 | ממנו |  |
| 7:26.394 | שהסכימו |  |
| 7:26.395 | להמליכו | , |
| 7:26.396 | למרות |  |
| 7:26.397 | דבר |  |
| 7:26.398 | ה | ' |
| 7:26.399 | אשר |  |
| 7:26.400 | דבר |  |
| 7:26.401 | על |  |
| 7:26.402 | יד |  |
| 7:26.403 | נביאו |  |
| 7:26.404 | וכבודי |  |
| 7:26.405 | לאחר |  |
| 7:26.406 | לא |  |
| 7:26.407 | אתן | . |
| 7:26.408 | ודבריהם |  |
| 7:26.409 | אלו |  |
| 7:26.410 | מבוארים |  |
| 7:26.411 | בזהר |  |
| 7:26.412 | בלק |  |
| 7:26.413 | דף |  |
| 7:26.414 | קצ"א |  |
| 7:26.415 | עם |  |
| 7:26.416 | פירוש |  |
| 7:26.417 | מקדש |  |
| 7:26.418 | מלך | . |
| 7:26.419 | ובפרט |  |
| 7:26.420 | בספר |  |
| 7:26.421 | עז |  |
| 7:26.422 | לאלהים |  |
| 7:26.423 | שהיא |  |
| 7:26.424 | משל |  |
| 7:26.425 | על |  |
| 7:26.426 | זה |  |
| 7:26.427 | ממלך |  |
| 7:26.428 | ב"ו |  |
| 7:26.429 | שמלך |  |
| 7:26.430 | שגם |  |
| 7:26.431 | אביו |  |
| 7:26.432 | ואימו |  |
| 7:26.433 | צריכים |  |
| 7:26.434 | לכבדו | , |
| 7:26.435 | וזה |  |
| 7:26.436 | המשל |  |
| 7:26.437 | מתנגד |  |
| 7:26.438 | למשל |  |
| 7:26.439 | שהביאו |  |
| 7:26.440 | רז"ל |  |
| 7:26.441 | בשמות |  |
| 7:26.442 | רבה |  |
| 7:26.443 | על |  |
| 7:26.444 | אנכי |  |
| 7:26.445 | ה | ' |
| 7:26.446 | אלהיך | . |
| 7:26.447 | ובספר |  |
| 7:26.448 | שלשלת |  |
| 7:26.449 | הקבלה |  |
| 7:26.450 | להראב"ד |  |
| 7:26.451 | כתב |  |
| 7:26.452 | משם |  |
| 7:26.453 | ס | ' |
| 7:26.454 | יוחסין |  |
| 7:26.455 | וז"ל |  |
| 7:26.456 | בשנת |  |
| 7:26.457 | כמו |  |
| 7:26.458 | חמשת |  |
| 7:26.459 | אלפים |  |
| 7:26.460 | וחמשים |  |
| 7:26.461 | שנה |  |
| 7:26.462 | נמצאו |  |
| 7:26.463 | כתות |  |
| 7:26.464 | אנשים |  |
| 7:26.465 | שאמרו |  |
| 7:26.466 | כי |  |
| 7:26.467 | דברי |  |
| 7:26.468 | הזהר |  |
| 7:26.469 | אשר |  |
| 7:26.470 | הם |  |
| 7:26.471 | בלשון |  |
| 7:26.472 | ירושלמי |  |
| 7:26.473 | הם |  |
| 7:26.474 | דברי |  |
| 7:26.475 | רשב"י | . |
| 7:26.476 | אבל |  |
| 7:26.477 | אשר |  |
| 7:26.478 | הם |  |
| 7:26.479 | בלשון |  |
| 7:26.480 | הקדש |  |
| 7:26.481 | אינן |  |
| 7:26.482 | דבריו | , |
| 7:26.483 | ויש |  |
| 7:26.484 | אומרים |  |
| 7:26.485 | שהרמב"ן |  |
| 7:26.486 | מצאו |  |
| 7:26.487 | בארץ |  |
| 7:26.488 | ישראל |  |
| 7:26.489 | והביא |  |
| 7:26.490 | אותו |  |
| 7:26.491 | לקטלוניאה | , |
| 7:26.492 | ואחר |  |
| 7:26.493 | לארגון | . |
| 7:26.494 | ונפל |  |
| 7:26.495 | ליד |  |
| 7:26.496 | רבי |  |
| 7:26.497 | משה |  |
| 7:26.498 | דיליואן | . |
| 7:26.499 | ויש |  |
| 7:26.500 | אומרים |  |
| 7:26.501 | כי |  |
| 7:26.502 | ר | ' |
| 7:26.503 | משה |  |
| 7:26.504 | דיליואן |  |
| 7:26.505 | הנז | ' |
| 7:26.506 | היה |  |
| 7:26.507 | חכם |  |
| 7:26.508 | גדול |  |
| 7:26.509 | ועושה |  |
| 7:26.510 | אלו |  |
| 7:26.511 | הפירושים |  |
| 7:26.512 | משכלו | , |
| 7:26.513 | ולמען |  |
| 7:26.514 | יקח |  |
| 7:26.515 | מהם |  |
| 7:26.516 | מחיר |  |
| 7:26.517 | רב |  |
| 7:26.518 | מהחכמים |  |
| 7:26.519 | היה |  |
| 7:26.520 | כותבם |  |
| 7:26.521 | ותולה |  |
| 7:26.522 | אותם |  |
| 7:26.523 | ברשב"י |  |
| 7:26.524 | וחבורתו | , |
| 7:26.525 | ע"כ |  |
| 7:26.526 | בספר |  |
| 7:26.527 | יוחסין |  |
| 7:26.528 | ומאריך |  |
| 7:26.529 | עוד |  |
| 7:26.530 | יותר |  |
| 7:26.531 | עכ"ל |  |
| 7:26.532 | שלשלת |  |
| 7:26.533 | הקבלה | . |
| 7:26.534 | וכבר |  |
| 7:26.535 | הבאתי |  |
| 7:26.536 | לעיל |  |
| 7:26.537 | סי | ' |
| 7:26.538 | צ"ד |  |
| 7:26.539 | דברי |  |
| 7:26.540 | הרב |  |
| 7:26.541 | חיד"א | , |
| 7:26.542 | בס | ' |
| 7:26.543 | שם |  |
| 7:26.544 | הגדולים |  |
| 7:26.545 | וגליון |  |
| 7:26.546 | שהובא |  |
| 7:26.547 | מירושלם |  |
| 7:26.548 | ת"י |  |
| 7:26.549 | על |  |
| 7:26.550 | ידי |  |
| 7:26.551 | מלשינות |  |
| 7:26.552 | הקשר |  |
| 7:26.553 | אשר |  |
| 7:26.554 | הרה |  |
| 7:26.555 | וילד |  |
| 7:26.556 | ארך |  |
| 7:26.557 | הלשון | . |
| 7:26.558 | והנה |  |
| 7:26.559 | שכחו |  |
| 7:26.560 | דין |  |
| 7:26.561 | פשוט |  |
| 7:26.562 | בשועיורה |  |
| 7:26.563 | דעה |  |
| 7:26.564 | סי | ' |
| 7:26.565 | רפ"א |  |
| 7:26.566 | שספר |  |
| 7:26.567 | תורה |  |
| 7:26.568 | הנמצא |  |
| 7:26.569 | ביד |  |
| 7:26.570 | גוי |  |
| 7:26.571 | דלד | ' |
| 7:26.572 | הרי"ף |  |
| 7:26.573 | והרא"ש |  |
| 7:26.574 | והטור |  |
| 7:26.575 | יגנז |  |
| 7:26.576 | ואין |  |
| 7:26.577 | קורין |  |
| 7:26.578 | בו |  |
| 7:26.579 | אע"פ |  |
| 7:26.580 | שאין |  |
| 7:26.581 | בו |  |
| 7:26.582 | שום |  |
| 7:26.583 | חסרון |  |
| 7:26.584 | או |  |
| 7:26.585 | יתרון |  |
| 7:26.586 | משאר |  |
| 7:26.587 | ספרי |  |
| 7:26.588 | תורה | , |
| 7:26.589 | ולדעת |  |
| 7:26.590 | הרמב"ם |  |
| 7:26.591 | ז"ל |  |
| 7:26.592 | כיון |  |
| 7:26.593 | שאין |  |
| 7:26.594 | ידוע |  |
| 7:26.595 | מי |  |
| 7:26.596 | כתבו |  |
| 7:26.597 | ואין |  |
| 7:26.598 | בו |  |
| 7:26.599 | שום |  |
| 7:26.600 | שינוי |  |
| 7:26.601 | חזקתו |  |
| 7:26.602 | שישראל |  |
| 7:26.603 | כתבו |  |
| 7:26.604 | וקורין |  |
| 7:26.605 | בו | , |
| 7:26.606 | הא |  |
| 7:26.607 | אם |  |
| 7:26.608 | נמצא |  |
| 7:26.609 | בו |  |
| 7:26.610 | שום |  |
| 7:26.611 | שינוי |  |
| 7:26.612 | אפילו |  |
| 7:26.613 | בחסר |  |
| 7:26.614 | ויתר |  |
| 7:26.615 | דבר |  |
| 7:26.616 | פשוט |  |
| 7:26.617 | הוא |  |
| 7:26.618 | דלכ"ע |  |
| 7:26.619 | אין |  |
| 7:26.620 | קורין |  |
| 7:26.621 | בו |  |
| 7:26.622 | ובנדון |  |
| 7:26.623 | דידן |  |
| 7:26.624 | של |  |
| 7:26.625 | ספר |  |
| 7:26.626 | הזהר |  |
| 7:26.627 | יש |  |
| 7:26.628 | לדון |  |
| 7:26.629 | בו |  |
| 7:26.630 | קל |  |
| 7:26.631 | וחומר | , |
| 7:26.632 | ומה |  |
| 7:26.633 | אם |  |
| 7:26.634 | זה |  |
| 7:26.635 | שאין |  |
| 7:26.636 | בו |  |
| 7:26.637 | שום |  |
| 7:26.638 | קלקול |  |
| 7:26.639 | לאמונתינו |  |
| 7:26.640 | יגנז | ? |
| 7:26.641 | ספר |  |
| 7:26.642 | המבזה |  |
| 7:26.643 | את |  |
| 7:26.644 | קבלתינו |  |
| 7:26.645 | המשנה |  |
| 7:26.646 | והתלמוד |  |
| 7:26.647 | ומסית |  |
| 7:26.648 | ומדיח |  |
| 7:26.649 | אותנו |  |
| 7:26.650 | לעבוד |  |
| 7:26.651 | כחות |  |
| 7:26.652 | נבראים | . |
| 7:26.653 | קצר |  |
| 7:26.654 | אפים |  |
| 7:26.655 | עם |  |
| 7:26.656 | הוריו |  |
| 7:26.657 | ויולדיו | , |
| 7:26.658 | על |  |
| 7:26.659 | אחת |  |
| 7:26.660 | כמה |  |
| 7:26.661 | וכמה |  |
| 7:26.662 | שאין |  |
| 7:26.663 | להאמין |  |
| 7:26.664 | לו |  |
| 7:26.665 | להמיר |  |
| 7:26.666 | את |  |
| 7:26.667 | אלהינו |  |
| 7:26.668 | ברוך |  |
| 7:26.669 | הוא |  |
| 7:26.670 | וב"ש |  |
| 7:26.671 | ולסתור |  |
| 7:26.672 | פנות |  |
| 7:26.673 | תורתינו |  |
| 7:26.674 | הקדושה | . |
| 7:26.675 | ומדבריהם |  |
| 7:26.676 | מודאע |  |
| 7:26.677 | רבה |  |
| 7:26.678 | לאוריתא |  |
| 7:26.679 | שהתנאי |  |
| 7:26.680 | שאמר |  |
| 7:26.681 | הקמב"ן |  |
| 7:26.682 | שא"א |  |
| 7:26.683 | להבין |  |
| 7:26.684 | דברי |  |
| 7:26.685 | הקבלה |  |
| 7:26.686 | מתוך |  |
| 7:26.687 | הספר |  |
| 7:26.688 | רק |  |
| 7:26.689 | מפי |  |
| 7:26.690 | חכם |  |
| 7:26.691 | מקובל |  |
| 7:26.692 | לאזן |  |
| 7:26.693 | מקבל | , |
| 7:26.694 | לא |  |
| 7:26.695 | נתקיים |  |
| 7:26.696 | מעולם |  |
| 7:26.697 | תנאי |  |
| 7:26.698 | זה | , |
| 7:26.699 | שהרי |  |
| 7:26.700 | הם |  |
| 7:26.701 | בעצמן |  |
| 7:26.702 | מודים |  |
| 7:26.703 | שאחרי |  |
| 7:26.704 | מות |  |
| 7:26.705 | רשב"י |  |
| 7:26.706 | ובנו |  |
| 7:26.707 | ר"א |  |
| 7:26.708 | וכל |  |
| 7:26.709 | הדור |  |
| 7:26.710 | ההוא |  |
| 7:26.711 | נגנז |  |
| 7:26.712 | הזהר |  |
| 7:26.713 | ונשכחה |  |
| 7:26.714 | ואבדה |  |
| 7:26.715 | הקבלה |  |
| 7:26.716 | הלזו |  |
| 7:26.717 | ולא |  |
| 7:26.718 | נשאר |  |
| 7:26.719 | אדם |  |
| 7:26.720 | יודע |  |
| 7:26.721 | בה |  |
| 7:26.722 | עד |  |
| 7:26.723 | שנתגלה |  |
| 7:26.724 | ספר |  |
| 7:26.725 | הזהר |  |
| 7:26.726 | באלף |  |
| 7:26.727 | השישי |  |
| 7:26.728 | ע"י |  |
| 7:26.729 | המלך |  |
| 7:26.730 | העכו"ם | . |
| 7:26.731 | וממנו |  |
| 7:26.732 | נתפשטה |  |
| 7:26.733 | הקבלה |  |
| 7:26.734 | ולא |  |
| 7:26.735 | מפי |  |
| 7:26.736 | חכם |  |
| 7:26.737 | מקובל |  |
| 7:26.738 | אשר |  |
| 7:26.739 | קבל |  |
| 7:26.740 | מפה |  |
| 7:26.741 | לאזן | . |
| 7:26.742 | כי |  |
| 7:26.743 | כל |  |
| 7:26.744 | המקובלים |  |
| 7:26.745 | הרבים |  |
| 7:26.746 | החדשים |  |
| 7:26.747 | שבעולם |  |
| 7:26.748 | לא |  |
| 7:26.749 | שאבו |  |
| 7:26.750 | רק |  |
| 7:26.751 | מהזהר |  |
| 7:26.752 | וממנו |  |
| 7:26.753 | נתפשטה | , |
| 7:26.754 | ולפי |  |
| 7:26.755 | דברי |  |
| 7:26.756 | הרמב"ן |  |
| 7:26.757 | זהו |  |
| 7:26.758 | הגורם |  |
| 7:26.759 | שכולם |  |
| 7:26.760 | שגגו |  |
| 7:26.761 | וטעו |  |
| 7:26.762 | ועקרו |  |
| 7:26.763 | יסודי |  |
| 7:26.764 | תורתינו |  |
| 7:26.765 | הקדושה |  |
| 7:26.766 | והאמינו |  |
| 7:26.767 | באלהים |  |
| 7:26.768 | רבים | . |
| 7:26.769 | ובאו |  |
| 7:26.770 | לידי |  |
| 7:26.771 | שתוף |  |
| 7:26.772 | שם |  |
| 7:26.773 | שמים |  |
| 7:26.774 | ודבר |  |
| 7:26.775 | אחד |  |
| 7:26.776 | בעבור |  |
| 7:26.777 | מיעוט |  |
| 7:26.778 | הבנתם | . |
| 7:26.779 | אף |  |
| 7:26.780 | לפי |  |
| 7:26.781 | האמת |  |
| 7:26.782 | כיון |  |
| 7:26.783 | שגילה |  |
| 7:26.784 | דעתו |  |
| 7:26.785 | בביאור |  |
| 7:26.786 | שהוא |  |
| 7:26.787 | מבזה |  |
| 7:26.788 | את |  |
| 7:26.789 | המשנה |  |
| 7:26.790 | והתלמוד |  |
| 7:26.791 | ומדמה |  |
| 7:26.792 | שקלא |  |
| 7:26.793 | וטריא |  |
| 7:26.794 | דרבנן |  |
| 7:26.795 | המתאבקים |  |
| 7:26.796 | בתלמוד |  |
| 7:26.797 | לכלבין |  |
| 7:26.798 | דצוחין |  |
| 7:26.799 | הב | , |
| 7:26.800 | הב | , |
| 7:26.801 | אין |  |
| 7:26.802 | מקום |  |
| 7:26.803 | לכל |  |
| 7:26.804 | זה |  |
| 7:26.805 | עיין |  |
| 7:26.806 | תקוני |  |
| 7:26.807 | זהר |  |
| 7:26.808 | חדש | ). |
| 7:26.809 | ולפי |  |
| 7:26.810 | דברי |  |
| 7:26.811 | הרב |  |
| 7:26.812 | שם |  |
| 7:26.813 | הגדולים |  |
| 7:26.814 | והגליון |  |
| 7:26.815 | המובא |  |
| 7:26.816 | מירושת"ו |  |
| 7:26.817 | שר | ' |
| 7:26.818 | נחוניא |  |
| 7:26.819 | בן |  |
| 7:26.820 | הקנה |  |
| 7:26.821 | היה |  |
| 7:26.822 | ראש |  |
| 7:26.823 | המקובלים |  |
| 7:26.824 | אשר |  |
| 7:26.825 | גילו |  |
| 7:26.826 | את |  |
| 7:26.827 | הקבלה |  |
| 7:26.828 | החדשה |  |
| 7:26.829 | ולמדו |  |
| 7:26.830 | אותה |  |
| 7:26.831 | בגלוי | ! |
| 7:26.832 | עד |  |
| 7:26.833 | שתיקן |  |
| 7:26.834 | בכניסתו |  |
| 7:26.835 | לבית |  |
| 7:26.836 | המדרש |  |
| 7:26.837 | להתפלל |  |
| 7:26.838 | שלא |  |
| 7:26.839 | אכשל |  |
| 7:26.840 | בדבר |  |
| 7:26.841 | הלכה | , |
| 7:26.842 | שלא |  |
| 7:26.843 | אומר |  |
| 7:26.844 | על |  |
| 7:26.845 | טמא |  |
| 7:26.846 | טהור |  |
| 7:26.847 | וכו | ' |
| 7:26.848 | כדתנן |  |
| 7:26.849 | בברכות |  |
| 7:26.850 | פ | ' |
| 7:26.851 | תפלת |  |
| 7:26.852 | השחר | . |
| 7:26.853 | היה |  |
| 7:26.854 | לו |  |
| 7:26.855 | להתפלל |  |
| 7:26.856 | שלא |  |
| 7:26.857 | תבוא |  |
| 7:26.858 | תקלה |  |
| 7:26.859 | גדולה |  |
| 7:26.860 | כזו |  |
| 7:26.861 | על |  |
| 7:26.862 | ידו |  |
| 7:26.863 | להאמין |  |
| 7:26.864 | ברשויות |  |
| 7:26.865 | רבות |  |
| 7:26.866 | ולעבוד |  |
| 7:26.867 | אל |  |
| 7:26.868 | אחר |  |
| 7:26.869 | קצר |  |
| 7:26.870 | אפים |  |
| 7:26.871 | זעיר |  |
| 7:26.872 | אנפין | )? |
| 7:26.873 | לפי |  |
| 7:26.874 | שעונש |  |
| 7:26.875 | המחטיא |  |
| 7:26.876 | אחרים |  |
| 7:26.877 | להאמין |  |
| 7:26.878 | ברשויות |  |
| 7:26.879 | רבות |  |
| 7:26.880 | ולעבוד |  |
| 7:26.881 | אלהי |  |
| 7:26.882 | נכר |  |
| 7:26.883 | יותר |  |
| 7:26.884 | גדול |  |
| 7:26.885 | מן |  |
| 7:26.886 | הנכשל |  |
| 7:26.887 | להתיר |  |
| 7:26.888 | את |  |
| 7:26.889 | האסור |  |
| 7:26.890 | או |  |
| 7:26.891 | לאסור |  |
| 7:26.892 | את |  |
| 7:26.893 | המותר | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

