# Translation pack: milhamot-hashem perek-4-17 (Milchamot ha-Shem 4:28–4:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-17`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- **מערב** beside Ashkenaz, Sefarad, Yemen → "the Maghreb" (note: "the West").
## Terms

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
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- עלת **העלות** → "the Cause of causes".
- הסבה **הראשונה** → "the First Cause".
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- עבודה **אחרת** → "other worship" (note first in a unit).
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **רחמנא ליצלן** → "the Merciful One save us".
- **ידידי** → "my friend" (second person kept).
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 4:28

Hebrew (Wikisource): לה) ומזה יתבאר לך ידידי אמתת אמונתינו הטהורה המזוקקה על פי רבותינו ז"ל מקבלי השמועה ממשה רבי' ע"ה שהשם ית' שהוא הסבה הראשונה לבדו ברא את הכל בלי סיוע שום נברא דלא כמו שדרש הפילוסוף מחבר הזהר שהבאתי לעיל סי' כ' וכ"א וכ"ד כי לפי דבריו העלות הרבות שהמציא מסייעות זו את זו. וכל אחת נוטלת רשות מסבה שהיא למעלה ממנה. ויש מקום לבעל הדין לומר עתיק ממתח בדרום ואריך אנפין בצפון ואבא במזרח ואימא במערב וזעיר ואשתו בזויות ואדם קדמון הגדול שבהם ממדד באמצע. ונגד מה שאמרו רז"ל בנוהג שבעולם מלך ב"ו מתקלס במדינה וגדולי המדינה מתקלסין עמו לפי שהם נושאים עמו במשאו וכו'. כן עשה הפילוסוף נביא השקר מחבר הזהר ושאר מקובלים בכבוד אבינו מלכינו שכל הפרצופים שהמציאו ברעיוניהם הם מתקלסים עמו! אדם קדמון הגדול שבהם מתקלס במאמר ראו עתה כי אני אני הוא אני אמית ואחיה וכו' שאין עילה אחרת עליו שיטול רשות ממנה! ואבא מתקלס במאמר יהי אור ויקוו המים כדלעיל! ואימא מתקלסת במאמר נעשה אדם! ועתיק מתקלס במאמר ה' פעלך בקרב שנים וחייהו. דלעתיקא קדישא אתמר כמבואר באדרת נשא! וזעיר מתקלס במאמר אנכי ה' אלהיך אשר הוצאתיך וכו'. כמו שכ' ספר הברית (יובא לקמן סי' ט"ז) והרב נחלת יוסף דף ס"א וס"ב! ומלכות מתקלסת במאמר אלה הם האלוהים המכים את מצרים שהמלכות היתה מכה בהם בסיועת הבינה כמו שכ' הר' נחלת יוסף ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:28.1 | לה | ) |
| 4:28.2 | ומזה |  |
| 4:28.3 | יתבאר |  |
| 4:28.4 | לך |  |
| 4:28.5 | ידידי |  |
| 4:28.6 | אמתת |  |
| 4:28.7 | אמונתינו |  |
| 4:28.8 | הטהורה |  |
| 4:28.9 | המזוקקה |  |
| 4:28.10 | על |  |
| 4:28.11 | פי |  |
| 4:28.12 | רבותינו |  |
| 4:28.13 | ז"ל |  |
| 4:28.14 | מקבלי |  |
| 4:28.15 | השמועה |  |
| 4:28.16 | ממשה |  |
| 4:28.17 | רבי | ' |
| 4:28.18 | ע"ה |  |
| 4:28.19 | שהשם |  |
| 4:28.20 | ית | ' |
| 4:28.21 | שהוא |  |
| 4:28.22 | הסבה |  |
| 4:28.23 | הראשונה |  |
| 4:28.24 | לבדו |  |
| 4:28.25 | ברא |  |
| 4:28.26 | את |  |
| 4:28.27 | הכל |  |
| 4:28.28 | בלי |  |
| 4:28.29 | סיוע |  |
| 4:28.30 | שום |  |
| 4:28.31 | נברא |  |
| 4:28.32 | דלא |  |
| 4:28.33 | כמו |  |
| 4:28.34 | שדרש |  |
| 4:28.35 | הפילוסוף |  |
| 4:28.36 | מחבר |  |
| 4:28.37 | הזהר |  |
| 4:28.38 | שהבאתי |  |
| 4:28.39 | לעיל |  |
| 4:28.40 | סי | ' |
| 4:28.41 | כ | ' |
| 4:28.42 | וכ"א |  |
| 4:28.43 | וכ"ד |  |
| 4:28.44 | כי |  |
| 4:28.45 | לפי |  |
| 4:28.46 | דבריו |  |
| 4:28.47 | העלות |  |
| 4:28.48 | הרבות |  |
| 4:28.49 | שהמציא |  |
| 4:28.50 | מסייעות |  |
| 4:28.51 | זו |  |
| 4:28.52 | את |  |
| 4:28.53 | זו | . |
| 4:28.54 | וכל |  |
| 4:28.55 | אחת |  |
| 4:28.56 | נוטלת |  |
| 4:28.57 | רשות |  |
| 4:28.58 | מסבה |  |
| 4:28.59 | שהיא |  |
| 4:28.60 | למעלה |  |
| 4:28.61 | ממנה | . |
| 4:28.62 | ויש |  |
| 4:28.63 | מקום |  |
| 4:28.64 | לבעל |  |
| 4:28.65 | הדין |  |
| 4:28.66 | לומר |  |
| 4:28.67 | עתיק |  |
| 4:28.68 | ממתח |  |
| 4:28.69 | בדרום |  |
| 4:28.70 | ואריך |  |
| 4:28.71 | אנפין |  |
| 4:28.72 | בצפון |  |
| 4:28.73 | ואבא |  |
| 4:28.74 | במזרח |  |
| 4:28.75 | ואימא |  |
| 4:28.76 | במערב |  |
| 4:28.77 | וזעיר |  |
| 4:28.78 | ואשתו |  |
| 4:28.79 | בזויות |  |
| 4:28.80 | ואדם |  |
| 4:28.81 | קדמון |  |
| 4:28.82 | הגדול |  |
| 4:28.83 | שבהם |  |
| 4:28.84 | ממדד |  |
| 4:28.85 | באמצע | . |
| 4:28.86 | ונגד |  |
| 4:28.87 | מה |  |
| 4:28.88 | שאמרו |  |
| 4:28.89 | רז"ל |  |
| 4:28.90 | בנוהג |  |
| 4:28.91 | שבעולם |  |
| 4:28.92 | מלך |  |
| 4:28.93 | ב"ו |  |
| 4:28.94 | מתקלס |  |
| 4:28.95 | במדינה |  |
| 4:28.96 | וגדולי |  |
| 4:28.97 | המדינה |  |
| 4:28.98 | מתקלסין |  |
| 4:28.99 | עמו |  |
| 4:28.100 | לפי |  |
| 4:28.101 | שהם |  |
| 4:28.102 | נושאים |  |
| 4:28.103 | עמו |  |
| 4:28.104 | במשאו |  |
| 4:28.105 | וכו | '. |
| 4:28.106 | כן |  |
| 4:28.107 | עשה |  |
| 4:28.108 | הפילוסוף |  |
| 4:28.109 | נביא |  |
| 4:28.110 | השקר |  |
| 4:28.111 | מחבר |  |
| 4:28.112 | הזהר |  |
| 4:28.113 | ושאר |  |
| 4:28.114 | מקובלים |  |
| 4:28.115 | בכבוד |  |
| 4:28.116 | אבינו |  |
| 4:28.117 | מלכינו |  |
| 4:28.118 | שכל |  |
| 4:28.119 | הפרצופים |  |
| 4:28.120 | שהמציאו |  |
| 4:28.121 | ברעיוניהם |  |
| 4:28.122 | הם |  |
| 4:28.123 | מתקלסים |  |
| 4:28.124 | עמו | ! |
| 4:28.125 | אדם |  |
| 4:28.126 | קדמון |  |
| 4:28.127 | הגדול |  |
| 4:28.128 | שבהם |  |
| 4:28.129 | מתקלס |  |
| 4:28.130 | במאמר |  |
| 4:28.131 | ראו |  |
| 4:28.132 | עתה |  |
| 4:28.133 | כי |  |
| 4:28.134 | אני |  |
| 4:28.135 | אני |  |
| 4:28.136 | הוא |  |
| 4:28.137 | אני |  |
| 4:28.138 | אמית |  |
| 4:28.139 | ואחיה |  |
| 4:28.140 | וכו | ' |
| 4:28.141 | שאין |  |
| 4:28.142 | עילה |  |
| 4:28.143 | אחרת |  |
| 4:28.144 | עליו |  |
| 4:28.145 | שיטול |  |
| 4:28.146 | רשות |  |
| 4:28.147 | ממנה | ! |
| 4:28.148 | ואבא |  |
| 4:28.149 | מתקלס |  |
| 4:28.150 | במאמר |  |
| 4:28.151 | יהי |  |
| 4:28.152 | אור |  |
| 4:28.153 | ויקוו |  |
| 4:28.154 | המים |  |
| 4:28.155 | כדלעיל | ! |
| 4:28.156 | ואימא |  |
| 4:28.157 | מתקלסת |  |
| 4:28.158 | במאמר |  |
| 4:28.159 | נעשה |  |
| 4:28.160 | אדם | ! |
| 4:28.161 | ועתיק |  |
| 4:28.162 | מתקלס |  |
| 4:28.163 | במאמר |  |
| 4:28.164 | ה | ' |
| 4:28.165 | פעלך |  |
| 4:28.166 | בקרב |  |
| 4:28.167 | שנים |  |
| 4:28.168 | וחייהו | . |
| 4:28.169 | דלעתיקא |  |
| 4:28.170 | קדישא |  |
| 4:28.171 | אתמר |  |
| 4:28.172 | כמבואר |  |
| 4:28.173 | באדרת |  |
| 4:28.174 | נשא | ! |
| 4:28.175 | וזעיר |  |
| 4:28.176 | מתקלס |  |
| 4:28.177 | במאמר |  |
| 4:28.178 | אנכי |  |
| 4:28.179 | ה | ' |
| 4:28.180 | אלהיך |  |
| 4:28.181 | אשר |  |
| 4:28.182 | הוצאתיך |  |
| 4:28.183 | וכו | '. |
| 4:28.184 | כמו |  |
| 4:28.185 | שכ | ' |
| 4:28.186 | ספר |  |
| 4:28.187 | הברית |  |
| 4:28.188 | יובא |  |
| 4:28.189 | לקמן |  |
| 4:28.190 | סי | ' |
| 4:28.191 | ט"ז | ) |
| 4:28.192 | והרב |  |
| 4:28.193 | נחלת |  |
| 4:28.194 | יוסף |  |
| 4:28.195 | דף |  |
| 4:28.196 | ס"א |  |
| 4:28.197 | וס"ב | ! |
| 4:28.198 | ומלכות |  |
| 4:28.199 | מתקלסת |  |
| 4:28.200 | במאמר |  |
| 4:28.201 | אלה |  |
| 4:28.202 | הם |  |
| 4:28.203 | האלוהים |  |
| 4:28.204 | המכים |  |
| 4:28.205 | את |  |
| 4:28.206 | מצרים |  |
| 4:28.207 | שהמלכות |  |
| 4:28.208 | היתה |  |
| 4:28.209 | מכה |  |
| 4:28.210 | בהם |  |
| 4:28.211 | בסיועת |  |
| 4:28.212 | הבינה |  |
| 4:28.213 | כמו |  |
| 4:28.214 | שכ | ' |
| 4:28.215 | הר | ' |
| 4:28.216 | נחלת |  |
| 4:28.217 | יוסף |  |
| 4:28.218 | ע"ש | . |


## Milchamot ha-Shem 4:29

Hebrew (Wikisource): ראה ידידי איך המקובלים מקלסים לכל גדולי הפרצופים בקילוסו של מלך מלכי המלכים כאלו הם נושאים במשאי עמו היפך דברי רז"ל שאמרו אתה הוא האלהים לבדך וכו' כדלעיל! דרחמנא לישיזבן מאמונות זרות כאלה! וראיתי להר' מלמד התלמידים בפרשת יתרו שכ' שהמלאכים נבראו קודם בריאת שמים וארץ והביא דבריו בעל מנורת המאור בפ' צ"ג, ומסתברא דנפשיה כתב כן, ואומדנא בעלמא היא, ומסתייע לדייק לה מברכת ברוך שאמר, וראיתו דיוקו ליתיה דאנן מעיקרא לא גרסינן ברוך עושה בראשית כמו שכתב מהרי'ץ בסי' עץ חיים, ואף לפי גרסתם אינו מוכרח דלא מצינו שהמלאכים נקראו ראשית, ולדברי רז"ל בבראשית רבה שומעין דמלתיהו מסתברא טפי, ועוד שדבריהם דברי קבלה, ואין להאריך במילי דאגדתא כ"כ במילי דלא נפקא מינה לדינא, כי כל מטרתינו פה לברר אמתת יחד הבורא ית' על פי תורתינו הקדושה ועל פי קבלת רבותינו ומעתיקי השמועה לדעת מה יעשה ישראל באומרו בכל יום ערב ובוקר שמע ישראל וכו' וה' אלהים יעזר לי וכו'.

| id | word | punctuation after |
|---|---|---|
| 4:29.1 | ראה |  |
| 4:29.2 | ידידי |  |
| 4:29.3 | איך |  |
| 4:29.4 | המקובלים |  |
| 4:29.5 | מקלסים |  |
| 4:29.6 | לכל |  |
| 4:29.7 | גדולי |  |
| 4:29.8 | הפרצופים |  |
| 4:29.9 | בקילוסו |  |
| 4:29.10 | של |  |
| 4:29.11 | מלך |  |
| 4:29.12 | מלכי |  |
| 4:29.13 | המלכים |  |
| 4:29.14 | כאלו |  |
| 4:29.15 | הם |  |
| 4:29.16 | נושאים |  |
| 4:29.17 | במשאי |  |
| 4:29.18 | עמו |  |
| 4:29.19 | היפך |  |
| 4:29.20 | דברי |  |
| 4:29.21 | רז"ל |  |
| 4:29.22 | שאמרו |  |
| 4:29.23 | אתה |  |
| 4:29.24 | הוא |  |
| 4:29.25 | האלהים |  |
| 4:29.26 | לבדך |  |
| 4:29.27 | וכו | ' |
| 4:29.28 | כדלעיל | ! |
| 4:29.29 | דרחמנא |  |
| 4:29.30 | לישיזבן |  |
| 4:29.31 | מאמונות |  |
| 4:29.32 | זרות |  |
| 4:29.33 | כאלה | ! |
| 4:29.34 | וראיתי |  |
| 4:29.35 | להר | ' |
| 4:29.36 | מלמד |  |
| 4:29.37 | התלמידים |  |
| 4:29.38 | בפרשת |  |
| 4:29.39 | יתרו |  |
| 4:29.40 | שכ | ' |
| 4:29.41 | שהמלאכים |  |
| 4:29.42 | נבראו |  |
| 4:29.43 | קודם |  |
| 4:29.44 | בריאת |  |
| 4:29.45 | שמים |  |
| 4:29.46 | וארץ |  |
| 4:29.47 | והביא |  |
| 4:29.48 | דבריו |  |
| 4:29.49 | בעל |  |
| 4:29.50 | מנורת |  |
| 4:29.51 | המאור |  |
| 4:29.52 | בפ | ' |
| 4:29.53 | צ"ג | , |
| 4:29.54 | ומסתברא |  |
| 4:29.55 | דנפשיה |  |
| 4:29.56 | כתב |  |
| 4:29.57 | כן | , |
| 4:29.58 | ואומדנא |  |
| 4:29.59 | בעלמא |  |
| 4:29.60 | היא | , |
| 4:29.61 | ומסתייע |  |
| 4:29.62 | לדייק |  |
| 4:29.63 | לה |  |
| 4:29.64 | מברכת |  |
| 4:29.65 | ברוך |  |
| 4:29.66 | שאמר | , |
| 4:29.67 | וראיתו |  |
| 4:29.68 | דיוקו |  |
| 4:29.69 | ליתיה |  |
| 4:29.70 | דאנן |  |
| 4:29.71 | מעיקרא |  |
| 4:29.72 | לא |  |
| 4:29.73 | גרסינן |  |
| 4:29.74 | ברוך |  |
| 4:29.75 | עושה |  |
| 4:29.76 | בראשית |  |
| 4:29.77 | כמו |  |
| 4:29.78 | שכתב |  |
| 4:29.79 | מהרי'ץ |  |
| 4:29.80 | בסי | ' |
| 4:29.81 | עץ |  |
| 4:29.82 | חיים | , |
| 4:29.83 | ואף |  |
| 4:29.84 | לפי |  |
| 4:29.85 | גרסתם |  |
| 4:29.86 | אינו |  |
| 4:29.87 | מוכרח |  |
| 4:29.88 | דלא |  |
| 4:29.89 | מצינו |  |
| 4:29.90 | שהמלאכים |  |
| 4:29.91 | נקראו |  |
| 4:29.92 | ראשית | , |
| 4:29.93 | ולדברי |  |
| 4:29.94 | רז"ל |  |
| 4:29.95 | בבראשית |  |
| 4:29.96 | רבה |  |
| 4:29.97 | שומעין |  |
| 4:29.98 | דמלתיהו |  |
| 4:29.99 | מסתברא |  |
| 4:29.100 | טפי | , |
| 4:29.101 | ועוד |  |
| 4:29.102 | שדבריהם |  |
| 4:29.103 | דברי |  |
| 4:29.104 | קבלה | , |
| 4:29.105 | ואין |  |
| 4:29.106 | להאריך |  |
| 4:29.107 | במילי |  |
| 4:29.108 | דאגדתא |  |
| 4:29.109 | כ"כ |  |
| 4:29.110 | במילי |  |
| 4:29.111 | דלא |  |
| 4:29.112 | נפקא |  |
| 4:29.113 | מינה |  |
| 4:29.114 | לדינא | , |
| 4:29.115 | כי |  |
| 4:29.116 | כל |  |
| 4:29.117 | מטרתינו |  |
| 4:29.118 | פה |  |
| 4:29.119 | לברר |  |
| 4:29.120 | אמתת |  |
| 4:29.121 | יחד |  |
| 4:29.122 | הבורא |  |
| 4:29.123 | ית | ' |
| 4:29.124 | על |  |
| 4:29.125 | פי |  |
| 4:29.126 | תורתינו |  |
| 4:29.127 | הקדושה |  |
| 4:29.128 | ועל |  |
| 4:29.129 | פי |  |
| 4:29.130 | קבלת |  |
| 4:29.131 | רבותינו |  |
| 4:29.132 | ומעתיקי |  |
| 4:29.133 | השמועה |  |
| 4:29.134 | לדעת |  |
| 4:29.135 | מה |  |
| 4:29.136 | יעשה |  |
| 4:29.137 | ישראל |  |
| 4:29.138 | באומרו |  |
| 4:29.139 | בכל |  |
| 4:29.140 | יום |  |
| 4:29.141 | ערב |  |
| 4:29.142 | ובוקר |  |
| 4:29.143 | שמע |  |
| 4:29.144 | ישראל |  |
| 4:29.145 | וכו | ' |
| 4:29.146 | וה | ' |
| 4:29.147 | אלהים |  |
| 4:29.148 | יעזר |  |
| 4:29.149 | לי |  |
| 4:29.150 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

