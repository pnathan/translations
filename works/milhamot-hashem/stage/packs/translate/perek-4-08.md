# Translation pack: milhamot-hashem perek-4-08 (Milchamot ha-Shem 4:13–4:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-08`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
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
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
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
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **השגחה** → "providence".
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
- **כ"ש / ק"ו** → "all the more so".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין** אל → "call upon".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **רוע** לב → "wickedness of heart".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **רזא דמלה** → "the secret of the matter".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 4:13

Hebrew (Wikisource): מכל מאמרי הזהר ומפרשיו הנז' לעיל נתבאר שהם קורין לכל אחד מן הפרצופים שבאצילות הוי'ה ואדנות ואלהים והקב"ה. ובחרו לעבוד לפרצוף האחרון שהוא זעיר אנפין. ואמרו שהאין סוף עם כל הפרצופים הנאצלים ממנו לא שייכא שהו שום עבודה ולא קריאה, ושהמתפלל אליהם אינו נענה לרוב רוממותם, וכ"ש פרצופים שבעולמות שהם למעלה מעולם האצילות דלא שייכא בהו שום עבודה לרוב העלמם ורק בזעיר אנפין לבדו שייך עבודה ותפלה, וקריאה אליו בעת צרה, כי הוא עמוד האמצעי [13] המקשר כל הכחות שלמעלה ושלמטה, שכן גידלוהו אבא ואימא והשליטוהו על כל הברואים, וצוו לעבדו ולברכו, והוא לדעתם ה' אלהינו. והוא מבואר בזהר פרשת בלק דף קצ"א ע"ב (ח"ג קצא, ב) וז"ל עם פי' מקדש מלך על פסוק מונע בר יקבוהו לאום, ורזא דמלה כתיב ברזא דלעילא מה שמו ומה שם בנו כי תדע, ההוא שם ידיעה ה' צבאות שמו (פי' אבא) שם בנו ישראל שמו (פי' זעיר אנפין) דכתיב בני בכורי ישראל, והאי ישראל כל מפתחאן דמהימנותא בית תלין. ואיהו משתכח ואמר ה' אמר אלי בני אתה (פי' ז"א אומר שאבא הנקרא ה' צבאות אצר אלי בני אתה) והכי הוא ודאי. דהא אבא ואימא אעטרו ליה ובריכו ליה בכמה ברכאן. ואמרו ופקידו לכלא נשקו בר פן יאנף. נשקו להאי בר. כביכול שולטנו יהב ליה על כלא דכלא יפלחון דזכי לרחמי לרחמי. כל ברכאן דעילא ותתא לההוא בר סלקין ומתעטרן. ומאי דמנע ברכן מהאי בר יפרשון חטאוי קמי מלכא קדישא (ס"א אימא קדישא) לאם ממש (פירש בינה) ע"כ. הרי בהדיא שהזהר קרי לאבא ה' צבאות, ולזעיר אנפין קרי ליה בן לאבא ואימא. והוא הנקרא בשם הוי'ה בכמה מקומות בזהר. ושאבא ואימא נתנו לו הכח והממשלה בכל הנבראים וצוו לעבדו, ושכל ברכותינו ותפלותינו אליו דוקא יעלו להיות עטרה לז"א, ולא לאבא ואימא, ולא לאריך אנפין, ולא לעתיק, ולא לאדם קדמון אשר קראוהו עלת כל העלות (כדלעיל סי כ"א) וגם לא לאדם קדמאה שכל העולמות שלמעלה מעולם האצילות נתהוו בתוך עיגולי (מקדש מלך בראשית קט"ו ד"ה ורמז) וכ"ש להאין סוף שהוא רחוק וגבוה מעל גבוה למעלה מן הכל. שעליו אמרו דלא שייך ביה שום עבודה ותפלה וברכה, כי רק בז"א שייך ותלוי ההשגחה בתחתונים ובמעשיהם, ולא באלוהות העליונים. והוא זעיר אנפין לבדו המשלם שכר טוב לצדיקים ונפרע מן הרשעים. אבל כל הפרצופי שלמעלה עם האין סוף שהוא למעלה מן הכל אין משגיחים ומפקידין בכל מעשה התחתונים בין טוב לרע, וקצור קצרה ידם מפדות, ואין בהם כח להציל את הקורא אליהם, בעת צרה, עזב ה' את השמים ואת הארץ ח"ו, והפקידם ביד קצר האפים ז"א. שהוא נברא ולא בורא כמו שיתבאר עוד לקמן מדברי הזהר והמקובלים. ורז"ל דרשוהו לפסוק מי עלה שמים וירד באופן נאות הובא בילקוט משלי רמז תתקס"ב ובמדרש הגדול ריש פרשת שמות וז"ל מי עלה שמים זה קב"ה עלה אלהים בתרועה וירד. וירד ה' על הר סיני, מי אסף רוח בחפניו, אשר בידו נפש כל חי. מי צרר מים בשמלה, צורר מים בעביו, מי הקים כל אפסי ארץ ה' ממית ומחיה, מה שמו. צור שמו, שדי שמו, ה' צבאות שמו, ומה שם בנו, בני בכורי ישראל, והענין פשוט, שהוא על ישראל שנקראו בנים למקום, לפי שעל ידי ישראל נתפרסם אלהותו ית' בעולם, לא שיהא שם לז"א הנקרא אצל הפילוסוף בזהר בן והנה עתה נשלים דבר שהתחלנו בו.

| id | word | punctuation after |
|---|---|---|
| 4:13.1 | מכל |  |
| 4:13.2 | מאמרי |  |
| 4:13.3 | הזהר |  |
| 4:13.4 | ומפרשיו |  |
| 4:13.5 | הנז | ' |
| 4:13.6 | לעיל |  |
| 4:13.7 | נתבאר |  |
| 4:13.8 | שהם |  |
| 4:13.9 | קורין |  |
| 4:13.10 | לכל |  |
| 4:13.11 | אחד |  |
| 4:13.12 | מן |  |
| 4:13.13 | הפרצופים |  |
| 4:13.14 | שבאצילות |  |
| 4:13.15 | הוי'ה |  |
| 4:13.16 | ואדנות |  |
| 4:13.17 | ואלהים |  |
| 4:13.18 | והקב"ה | . |
| 4:13.19 | ובחרו |  |
| 4:13.20 | לעבוד |  |
| 4:13.21 | לפרצוף |  |
| 4:13.22 | האחרון |  |
| 4:13.23 | שהוא |  |
| 4:13.24 | זעיר |  |
| 4:13.25 | אנפין | . |
| 4:13.26 | ואמרו |  |
| 4:13.27 | שהאין |  |
| 4:13.28 | סוף |  |
| 4:13.29 | עם |  |
| 4:13.30 | כל |  |
| 4:13.31 | הפרצופים |  |
| 4:13.32 | הנאצלים |  |
| 4:13.33 | ממנו |  |
| 4:13.34 | לא |  |
| 4:13.35 | שייכא |  |
| 4:13.36 | שהו |  |
| 4:13.37 | שום |  |
| 4:13.38 | עבודה |  |
| 4:13.39 | ולא |  |
| 4:13.40 | קריאה | , |
| 4:13.41 | ושהמתפלל |  |
| 4:13.42 | אליהם |  |
| 4:13.43 | אינו |  |
| 4:13.44 | נענה |  |
| 4:13.45 | לרוב |  |
| 4:13.46 | רוממותם | , |
| 4:13.47 | וכ"ש |  |
| 4:13.48 | פרצופים |  |
| 4:13.49 | שבעולמות |  |
| 4:13.50 | שהם |  |
| 4:13.51 | למעלה |  |
| 4:13.52 | מעולם |  |
| 4:13.53 | האצילות |  |
| 4:13.54 | דלא |  |
| 4:13.55 | שייכא |  |
| 4:13.56 | בהו |  |
| 4:13.57 | שום |  |
| 4:13.58 | עבודה |  |
| 4:13.59 | לרוב |  |
| 4:13.60 | העלמם |  |
| 4:13.61 | ורק |  |
| 4:13.62 | בזעיר |  |
| 4:13.63 | אנפין |  |
| 4:13.64 | לבדו |  |
| 4:13.65 | שייך |  |
| 4:13.66 | עבודה |  |
| 4:13.67 | ותפלה | , |
| 4:13.68 | וקריאה |  |
| 4:13.69 | אליו |  |
| 4:13.70 | בעת |  |
| 4:13.71 | צרה | , |
| 4:13.72 | כי |  |
| 4:13.73 | הוא |  |
| 4:13.74 | עמוד |  |
| 4:13.75 | האמצעי | [13] |
| 4:13.76 | המקשר |  |
| 4:13.77 | כל |  |
| 4:13.78 | הכחות |  |
| 4:13.79 | שלמעלה |  |
| 4:13.80 | ושלמטה | , |
| 4:13.81 | שכן |  |
| 4:13.82 | גידלוהו |  |
| 4:13.83 | אבא |  |
| 4:13.84 | ואימא |  |
| 4:13.85 | והשליטוהו |  |
| 4:13.86 | על |  |
| 4:13.87 | כל |  |
| 4:13.88 | הברואים | , |
| 4:13.89 | וצוו |  |
| 4:13.90 | לעבדו |  |
| 4:13.91 | ולברכו | , |
| 4:13.92 | והוא |  |
| 4:13.93 | לדעתם |  |
| 4:13.94 | ה | ' |
| 4:13.95 | אלהינו | . |
| 4:13.96 | והוא |  |
| 4:13.97 | מבואר |  |
| 4:13.98 | בזהר |  |
| 4:13.99 | פרשת |  |
| 4:13.100 | בלק |  |
| 4:13.101 | דף |  |
| 4:13.102 | קצ"א |  |
| 4:13.103 | ע"ב |  |
| 4:13.104 | ח"ג |  |
| 4:13.105 | קצא | , |
| 4:13.106 | ב | ) |
| 4:13.107 | וז"ל |  |
| 4:13.108 | עם |  |
| 4:13.109 | פי | ' |
| 4:13.110 | מקדש |  |
| 4:13.111 | מלך |  |
| 4:13.112 | על |  |
| 4:13.113 | פסוק |  |
| 4:13.114 | מונע |  |
| 4:13.115 | בר |  |
| 4:13.116 | יקבוהו |  |
| 4:13.117 | לאום | , |
| 4:13.118 | ורזא |  |
| 4:13.119 | דמלה |  |
| 4:13.120 | כתיב |  |
| 4:13.121 | ברזא |  |
| 4:13.122 | דלעילא |  |
| 4:13.123 | מה |  |
| 4:13.124 | שמו |  |
| 4:13.125 | ומה |  |
| 4:13.126 | שם |  |
| 4:13.127 | בנו |  |
| 4:13.128 | כי |  |
| 4:13.129 | תדע | , |
| 4:13.130 | ההוא |  |
| 4:13.131 | שם |  |
| 4:13.132 | ידיעה |  |
| 4:13.133 | ה | ' |
| 4:13.134 | צבאות |  |
| 4:13.135 | שמו |  |
| 4:13.136 | פי | ' |
| 4:13.137 | אבא | ) |
| 4:13.138 | שם |  |
| 4:13.139 | בנו |  |
| 4:13.140 | ישראל |  |
| 4:13.141 | שמו |  |
| 4:13.142 | פי | ' |
| 4:13.143 | זעיר |  |
| 4:13.144 | אנפין | ) |
| 4:13.145 | דכתיב |  |
| 4:13.146 | בני |  |
| 4:13.147 | בכורי |  |
| 4:13.148 | ישראל | , |
| 4:13.149 | והאי |  |
| 4:13.150 | ישראל |  |
| 4:13.151 | כל |  |
| 4:13.152 | מפתחאן |  |
| 4:13.153 | דמהימנותא |  |
| 4:13.154 | בית |  |
| 4:13.155 | תלין | . |
| 4:13.156 | ואיהו |  |
| 4:13.157 | משתכח |  |
| 4:13.158 | ואמר |  |
| 4:13.159 | ה | ' |
| 4:13.160 | אמר |  |
| 4:13.161 | אלי |  |
| 4:13.162 | בני |  |
| 4:13.163 | אתה |  |
| 4:13.164 | פי | ' |
| 4:13.165 | ז"א |  |
| 4:13.166 | אומר |  |
| 4:13.167 | שאבא |  |
| 4:13.168 | הנקרא |  |
| 4:13.169 | ה | ' |
| 4:13.170 | צבאות |  |
| 4:13.171 | אצר |  |
| 4:13.172 | אלי |  |
| 4:13.173 | בני |  |
| 4:13.174 | אתה | ) |
| 4:13.175 | והכי |  |
| 4:13.176 | הוא |  |
| 4:13.177 | ודאי | . |
| 4:13.178 | דהא |  |
| 4:13.179 | אבא |  |
| 4:13.180 | ואימא |  |
| 4:13.181 | אעטרו |  |
| 4:13.182 | ליה |  |
| 4:13.183 | ובריכו |  |
| 4:13.184 | ליה |  |
| 4:13.185 | בכמה |  |
| 4:13.186 | ברכאן | . |
| 4:13.187 | ואמרו |  |
| 4:13.188 | ופקידו |  |
| 4:13.189 | לכלא |  |
| 4:13.190 | נשקו |  |
| 4:13.191 | בר |  |
| 4:13.192 | פן |  |
| 4:13.193 | יאנף | . |
| 4:13.194 | נשקו |  |
| 4:13.195 | להאי |  |
| 4:13.196 | בר | . |
| 4:13.197 | כביכול |  |
| 4:13.198 | שולטנו |  |
| 4:13.199 | יהב |  |
| 4:13.200 | ליה |  |
| 4:13.201 | על |  |
| 4:13.202 | כלא |  |
| 4:13.203 | דכלא |  |
| 4:13.204 | יפלחון |  |
| 4:13.205 | דזכי |  |
| 4:13.206 | לרחמי |  |
| 4:13.207 | לרחמי | . |
| 4:13.208 | כל |  |
| 4:13.209 | ברכאן |  |
| 4:13.210 | דעילא |  |
| 4:13.211 | ותתא |  |
| 4:13.212 | לההוא |  |
| 4:13.213 | בר |  |
| 4:13.214 | סלקין |  |
| 4:13.215 | ומתעטרן | . |
| 4:13.216 | ומאי |  |
| 4:13.217 | דמנע |  |
| 4:13.218 | ברכן |  |
| 4:13.219 | מהאי |  |
| 4:13.220 | בר |  |
| 4:13.221 | יפרשון |  |
| 4:13.222 | חטאוי |  |
| 4:13.223 | קמי |  |
| 4:13.224 | מלכא |  |
| 4:13.225 | קדישא |  |
| 4:13.226 | ס"א |  |
| 4:13.227 | אימא |  |
| 4:13.228 | קדישא | ) |
| 4:13.229 | לאם |  |
| 4:13.230 | ממש |  |
| 4:13.231 | פירש |  |
| 4:13.232 | בינה | ) |
| 4:13.233 | ע"כ | . |
| 4:13.234 | הרי |  |
| 4:13.235 | בהדיא |  |
| 4:13.236 | שהזהר |  |
| 4:13.237 | קרי |  |
| 4:13.238 | לאבא |  |
| 4:13.239 | ה | ' |
| 4:13.240 | צבאות | , |
| 4:13.241 | ולזעיר |  |
| 4:13.242 | אנפין |  |
| 4:13.243 | קרי |  |
| 4:13.244 | ליה |  |
| 4:13.245 | בן |  |
| 4:13.246 | לאבא |  |
| 4:13.247 | ואימא | . |
| 4:13.248 | והוא |  |
| 4:13.249 | הנקרא |  |
| 4:13.250 | בשם |  |
| 4:13.251 | הוי'ה |  |
| 4:13.252 | בכמה |  |
| 4:13.253 | מקומות |  |
| 4:13.254 | בזהר | . |
| 4:13.255 | ושאבא |  |
| 4:13.256 | ואימא |  |
| 4:13.257 | נתנו |  |
| 4:13.258 | לו |  |
| 4:13.259 | הכח |  |
| 4:13.260 | והממשלה |  |
| 4:13.261 | בכל |  |
| 4:13.262 | הנבראים |  |
| 4:13.263 | וצוו |  |
| 4:13.264 | לעבדו | , |
| 4:13.265 | ושכל |  |
| 4:13.266 | ברכותינו |  |
| 4:13.267 | ותפלותינו |  |
| 4:13.268 | אליו |  |
| 4:13.269 | דוקא |  |
| 4:13.270 | יעלו |  |
| 4:13.271 | להיות |  |
| 4:13.272 | עטרה |  |
| 4:13.273 | לז"א | , |
| 4:13.274 | ולא |  |
| 4:13.275 | לאבא |  |
| 4:13.276 | ואימא | , |
| 4:13.277 | ולא |  |
| 4:13.278 | לאריך |  |
| 4:13.279 | אנפין | , |
| 4:13.280 | ולא |  |
| 4:13.281 | לעתיק | , |
| 4:13.282 | ולא |  |
| 4:13.283 | לאדם |  |
| 4:13.284 | קדמון |  |
| 4:13.285 | אשר |  |
| 4:13.286 | קראוהו |  |
| 4:13.287 | עלת |  |
| 4:13.288 | כל |  |
| 4:13.289 | העלות |  |
| 4:13.290 | כדלעיל |  |
| 4:13.291 | סי |  |
| 4:13.292 | כ"א | ) |
| 4:13.293 | וגם |  |
| 4:13.294 | לא |  |
| 4:13.295 | לאדם |  |
| 4:13.296 | קדמאה |  |
| 4:13.297 | שכל |  |
| 4:13.298 | העולמות |  |
| 4:13.299 | שלמעלה |  |
| 4:13.300 | מעולם |  |
| 4:13.301 | האצילות |  |
| 4:13.302 | נתהוו |  |
| 4:13.303 | בתוך |  |
| 4:13.304 | עיגולי |  |
| 4:13.305 | מקדש |  |
| 4:13.306 | מלך |  |
| 4:13.307 | בראשית |  |
| 4:13.308 | קט"ו |  |
| 4:13.309 | ד"ה |  |
| 4:13.310 | ורמז | ) |
| 4:13.311 | וכ"ש |  |
| 4:13.312 | להאין |  |
| 4:13.313 | סוף |  |
| 4:13.314 | שהוא |  |
| 4:13.315 | רחוק |  |
| 4:13.316 | וגבוה |  |
| 4:13.317 | מעל |  |
| 4:13.318 | גבוה |  |
| 4:13.319 | למעלה |  |
| 4:13.320 | מן |  |
| 4:13.321 | הכל | . |
| 4:13.322 | שעליו |  |
| 4:13.323 | אמרו |  |
| 4:13.324 | דלא |  |
| 4:13.325 | שייך |  |
| 4:13.326 | ביה |  |
| 4:13.327 | שום |  |
| 4:13.328 | עבודה |  |
| 4:13.329 | ותפלה |  |
| 4:13.330 | וברכה | , |
| 4:13.331 | כי |  |
| 4:13.332 | רק |  |
| 4:13.333 | בז"א |  |
| 4:13.334 | שייך |  |
| 4:13.335 | ותלוי |  |
| 4:13.336 | ההשגחה |  |
| 4:13.337 | בתחתונים |  |
| 4:13.338 | ובמעשיהם | , |
| 4:13.339 | ולא |  |
| 4:13.340 | באלוהות |  |
| 4:13.341 | העליונים | . |
| 4:13.342 | והוא |  |
| 4:13.343 | זעיר |  |
| 4:13.344 | אנפין |  |
| 4:13.345 | לבדו |  |
| 4:13.346 | המשלם |  |
| 4:13.347 | שכר |  |
| 4:13.348 | טוב |  |
| 4:13.349 | לצדיקים |  |
| 4:13.350 | ונפרע |  |
| 4:13.351 | מן |  |
| 4:13.352 | הרשעים | . |
| 4:13.353 | אבל |  |
| 4:13.354 | כל |  |
| 4:13.355 | הפרצופי |  |
| 4:13.356 | שלמעלה |  |
| 4:13.357 | עם |  |
| 4:13.358 | האין |  |
| 4:13.359 | סוף |  |
| 4:13.360 | שהוא |  |
| 4:13.361 | למעלה |  |
| 4:13.362 | מן |  |
| 4:13.363 | הכל |  |
| 4:13.364 | אין |  |
| 4:13.365 | משגיחים |  |
| 4:13.366 | ומפקידין |  |
| 4:13.367 | בכל |  |
| 4:13.368 | מעשה |  |
| 4:13.369 | התחתונים |  |
| 4:13.370 | בין |  |
| 4:13.371 | טוב |  |
| 4:13.372 | לרע | , |
| 4:13.373 | וקצור |  |
| 4:13.374 | קצרה |  |
| 4:13.375 | ידם |  |
| 4:13.376 | מפדות | , |
| 4:13.377 | ואין |  |
| 4:13.378 | בהם |  |
| 4:13.379 | כח |  |
| 4:13.380 | להציל |  |
| 4:13.381 | את |  |
| 4:13.382 | הקורא |  |
| 4:13.383 | אליהם | , |
| 4:13.384 | בעת |  |
| 4:13.385 | צרה | , |
| 4:13.386 | עזב |  |
| 4:13.387 | ה | ' |
| 4:13.388 | את |  |
| 4:13.389 | השמים |  |
| 4:13.390 | ואת |  |
| 4:13.391 | הארץ |  |
| 4:13.392 | ח"ו | , |
| 4:13.393 | והפקידם |  |
| 4:13.394 | ביד |  |
| 4:13.395 | קצר |  |
| 4:13.396 | האפים |  |
| 4:13.397 | ז"א | . |
| 4:13.398 | שהוא |  |
| 4:13.399 | נברא |  |
| 4:13.400 | ולא |  |
| 4:13.401 | בורא |  |
| 4:13.402 | כמו |  |
| 4:13.403 | שיתבאר |  |
| 4:13.404 | עוד |  |
| 4:13.405 | לקמן |  |
| 4:13.406 | מדברי |  |
| 4:13.407 | הזהר |  |
| 4:13.408 | והמקובלים | . |
| 4:13.409 | ורז"ל |  |
| 4:13.410 | דרשוהו |  |
| 4:13.411 | לפסוק |  |
| 4:13.412 | מי |  |
| 4:13.413 | עלה |  |
| 4:13.414 | שמים |  |
| 4:13.415 | וירד |  |
| 4:13.416 | באופן |  |
| 4:13.417 | נאות |  |
| 4:13.418 | הובא |  |
| 4:13.419 | בילקוט |  |
| 4:13.420 | משלי |  |
| 4:13.421 | רמז |  |
| 4:13.422 | תתקס"ב |  |
| 4:13.423 | ובמדרש |  |
| 4:13.424 | הגדול |  |
| 4:13.425 | ריש |  |
| 4:13.426 | פרשת |  |
| 4:13.427 | שמות |  |
| 4:13.428 | וז"ל |  |
| 4:13.429 | מי |  |
| 4:13.430 | עלה |  |
| 4:13.431 | שמים |  |
| 4:13.432 | זה |  |
| 4:13.433 | קב"ה |  |
| 4:13.434 | עלה |  |
| 4:13.435 | אלהים |  |
| 4:13.436 | בתרועה |  |
| 4:13.437 | וירד | . |
| 4:13.438 | וירד |  |
| 4:13.439 | ה | ' |
| 4:13.440 | על |  |
| 4:13.441 | הר |  |
| 4:13.442 | סיני | , |
| 4:13.443 | מי |  |
| 4:13.444 | אסף |  |
| 4:13.445 | רוח |  |
| 4:13.446 | בחפניו | , |
| 4:13.447 | אשר |  |
| 4:13.448 | בידו |  |
| 4:13.449 | נפש |  |
| 4:13.450 | כל |  |
| 4:13.451 | חי | . |
| 4:13.452 | מי |  |
| 4:13.453 | צרר |  |
| 4:13.454 | מים |  |
| 4:13.455 | בשמלה | , |
| 4:13.456 | צורר |  |
| 4:13.457 | מים |  |
| 4:13.458 | בעביו | , |
| 4:13.459 | מי |  |
| 4:13.460 | הקים |  |
| 4:13.461 | כל |  |
| 4:13.462 | אפסי |  |
| 4:13.463 | ארץ |  |
| 4:13.464 | ה | ' |
| 4:13.465 | ממית |  |
| 4:13.466 | ומחיה | , |
| 4:13.467 | מה |  |
| 4:13.468 | שמו | . |
| 4:13.469 | צור |  |
| 4:13.470 | שמו | , |
| 4:13.471 | שדי |  |
| 4:13.472 | שמו | , |
| 4:13.473 | ה | ' |
| 4:13.474 | צבאות |  |
| 4:13.475 | שמו | , |
| 4:13.476 | ומה |  |
| 4:13.477 | שם |  |
| 4:13.478 | בנו | , |
| 4:13.479 | בני |  |
| 4:13.480 | בכורי |  |
| 4:13.481 | ישראל | , |
| 4:13.482 | והענין |  |
| 4:13.483 | פשוט | , |
| 4:13.484 | שהוא |  |
| 4:13.485 | על |  |
| 4:13.486 | ישראל |  |
| 4:13.487 | שנקראו |  |
| 4:13.488 | בנים |  |
| 4:13.489 | למקום | , |
| 4:13.490 | לפי |  |
| 4:13.491 | שעל |  |
| 4:13.492 | ידי |  |
| 4:13.493 | ישראל |  |
| 4:13.494 | נתפרסם |  |
| 4:13.495 | אלהותו |  |
| 4:13.496 | ית | ' |
| 4:13.497 | בעולם | , |
| 4:13.498 | לא |  |
| 4:13.499 | שיהא |  |
| 4:13.500 | שם |  |
| 4:13.501 | לז"א |  |
| 4:13.502 | הנקרא |  |
| 4:13.503 | אצל |  |
| 4:13.504 | הפילוסוף |  |
| 4:13.505 | בזהר |  |
| 4:13.506 | בן |  |
| 4:13.507 | והנה |  |
| 4:13.508 | עתה |  |
| 4:13.509 | נשלים |  |
| 4:13.510 | דבר |  |
| 4:13.511 | שהתחלנו |  |
| 4:13.512 | בו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

