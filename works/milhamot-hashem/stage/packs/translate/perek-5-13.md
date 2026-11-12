# Translation pack: milhamot-hashem perek-5-13 (Milchamot ha-Shem 5:20–5:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-13`.

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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ראב"ד** → "Ravad" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **אדרא / אידרא** of Naso → "Idra of Naso" (book title; meaning in a note).
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
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
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **תא חזי** → "come and see".
- מי **אתי** (ketiv): translated, the qere noted.
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 5:20

Hebrew (Wikisource): אמנם היה האר"י לפעמים אומר במקומו אדון עולם אשר מלך מפני שלא נזכר בו שום שם משמות הקודש. ושייך כל אותו השבח באלוה הראשון הנקרא אין סוף אשר לא יעבדוהו משום דלית ליה שם ידיע ולא שום כינוי. וגם נזכר באדון עולם והוא יהיה בתפארה, שלפי דעתם קצתו מתלבש בתפארת, לכן לא נמנע האר"י מלאמרו. אבל ביגדל שנזכר בו שם אלהים חי שהם מיחסים אותו לפרצופי אבא ואמא וגם לזעיר אנפין אשר אותם יעבודו שיש להם פרצוף שלם בדמות אדם ובצלמו. ואינם ראשונים שכבר קדמו להם סיבות, ואינם נעלמים, לכן הקפיד מאד על ארבעה בתים הראשונים שהם שקר לפי דעתם. הלא אזן מלים תבחן כמה גדלה שגיאת רבותינו בעלי הקבלה החדשה לעקור ארבעה שרשים מיסודי תורתינו הקדושה, ובהדיא אמר הרב יושר לבב בדף ג' ע"ב: "דע את אלהי אביך כולל החמשה פרצופים וכו' ועבדהו לזעיר אנפין אע"פ שהוא נברא כי לנשמתו אתה עובד בזה כי בלעדה אין כל נמצא, משום הכי לא אמר דע אלהי אביך אלא את אלהי אביך" וכו'. ודברי יושר לבב שאמר שהוא נברא מבוארים באדרא רבה דף קל"ח ע"ב וז"ל: "ה' פעלך דא זעיר אנפין וכו' חייהו למאן, חייהו לזעיר אנפין דכל נהירו דיליה מאינון שנים קדמוניות", וכו' ועוד הרבה ראיות בלשונות הזהר שהוא נברא, ועם כל זה מצוה לעבוד אותו. ואילו שמו אל לבם דברי רז"ל בבראשית רבה ובילקוט שהבאתי לעיל סי' ל"א כל שיש לו תולדות וכו' לא עבדו לנברא. גם הטילו חסרון בידיעת השי"ת בעתיד כמו שכתבתי לעיל בשם האר"י בספר ארבע מאות שקל כסף, כי עשו ידיעת השי"ת בעתיד ידיעת ספק כידיעת האצטגנים וכן כתב הראב"ד בהשגות פ"ז מהלכות תשובה. וז"ל: "אעפ"י שאין תשובה נצחת על זה טוב לסמוך לו קצת תשובה ואומר אם היו צדקת האדם ורשעתו תלוים בגזרת הבורא ית' היתה לנו השאלה קשה עד מאד, ועכשיו שהבורא הסיר זאת הממשלה מידו ומסרה בידי האדם עצמו אין ידיעתו גזרה, אבל היא כידיעת האצטגנינים שיודעים מכח אחר מה יהיה דרכיו של זה, והדבר ידוע שכל מקרה האדם מסרם הבורא בכל המזלות אלא שנתן בו השכל להיות מחזיקו לצאת מתחת המזל והוא הכח הנתון באדם", עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:20.1 | אמנם |  |
| 5:20.2 | היה |  |
| 5:20.3 | האר"י |  |
| 5:20.4 | לפעמים |  |
| 5:20.5 | אומר |  |
| 5:20.6 | במקומו |  |
| 5:20.7 | אדון |  |
| 5:20.8 | עולם |  |
| 5:20.9 | אשר |  |
| 5:20.10 | מלך |  |
| 5:20.11 | מפני |  |
| 5:20.12 | שלא |  |
| 5:20.13 | נזכר |  |
| 5:20.14 | בו |  |
| 5:20.15 | שום |  |
| 5:20.16 | שם |  |
| 5:20.17 | משמות |  |
| 5:20.18 | הקודש | . |
| 5:20.19 | ושייך |  |
| 5:20.20 | כל |  |
| 5:20.21 | אותו |  |
| 5:20.22 | השבח |  |
| 5:20.23 | באלוה |  |
| 5:20.24 | הראשון |  |
| 5:20.25 | הנקרא |  |
| 5:20.26 | אין |  |
| 5:20.27 | סוף |  |
| 5:20.28 | אשר |  |
| 5:20.29 | לא |  |
| 5:20.30 | יעבדוהו |  |
| 5:20.31 | משום |  |
| 5:20.32 | דלית |  |
| 5:20.33 | ליה |  |
| 5:20.34 | שם |  |
| 5:20.35 | ידיע |  |
| 5:20.36 | ולא |  |
| 5:20.37 | שום |  |
| 5:20.38 | כינוי | . |
| 5:20.39 | וגם |  |
| 5:20.40 | נזכר |  |
| 5:20.41 | באדון |  |
| 5:20.42 | עולם |  |
| 5:20.43 | והוא |  |
| 5:20.44 | יהיה |  |
| 5:20.45 | בתפארה | , |
| 5:20.46 | שלפי |  |
| 5:20.47 | דעתם |  |
| 5:20.48 | קצתו |  |
| 5:20.49 | מתלבש |  |
| 5:20.50 | בתפארת | , |
| 5:20.51 | לכן |  |
| 5:20.52 | לא |  |
| 5:20.53 | נמנע |  |
| 5:20.54 | האר"י |  |
| 5:20.55 | מלאמרו | . |
| 5:20.56 | אבל |  |
| 5:20.57 | ביגדל |  |
| 5:20.58 | שנזכר |  |
| 5:20.59 | בו |  |
| 5:20.60 | שם |  |
| 5:20.61 | אלהים |  |
| 5:20.62 | חי |  |
| 5:20.63 | שהם |  |
| 5:20.64 | מיחסים |  |
| 5:20.65 | אותו |  |
| 5:20.66 | לפרצופי |  |
| 5:20.67 | אבא |  |
| 5:20.68 | ואמא |  |
| 5:20.69 | וגם |  |
| 5:20.70 | לזעיר |  |
| 5:20.71 | אנפין |  |
| 5:20.72 | אשר |  |
| 5:20.73 | אותם |  |
| 5:20.74 | יעבודו |  |
| 5:20.75 | שיש |  |
| 5:20.76 | להם |  |
| 5:20.77 | פרצוף |  |
| 5:20.78 | שלם |  |
| 5:20.79 | בדמות |  |
| 5:20.80 | אדם |  |
| 5:20.81 | ובצלמו | . |
| 5:20.82 | ואינם |  |
| 5:20.83 | ראשונים |  |
| 5:20.84 | שכבר |  |
| 5:20.85 | קדמו |  |
| 5:20.86 | להם |  |
| 5:20.87 | סיבות | , |
| 5:20.88 | ואינם |  |
| 5:20.89 | נעלמים | , |
| 5:20.90 | לכן |  |
| 5:20.91 | הקפיד |  |
| 5:20.92 | מאד |  |
| 5:20.93 | על |  |
| 5:20.94 | ארבעה |  |
| 5:20.95 | בתים |  |
| 5:20.96 | הראשונים |  |
| 5:20.97 | שהם |  |
| 5:20.98 | שקר |  |
| 5:20.99 | לפי |  |
| 5:20.100 | דעתם | . |
| 5:20.101 | הלא |  |
| 5:20.102 | אזן |  |
| 5:20.103 | מלים |  |
| 5:20.104 | תבחן |  |
| 5:20.105 | כמה |  |
| 5:20.106 | גדלה |  |
| 5:20.107 | שגיאת |  |
| 5:20.108 | רבותינו |  |
| 5:20.109 | בעלי |  |
| 5:20.110 | הקבלה |  |
| 5:20.111 | החדשה |  |
| 5:20.112 | לעקור |  |
| 5:20.113 | ארבעה |  |
| 5:20.114 | שרשים |  |
| 5:20.115 | מיסודי |  |
| 5:20.116 | תורתינו |  |
| 5:20.117 | הקדושה | , |
| 5:20.118 | ובהדיא |  |
| 5:20.119 | אמר |  |
| 5:20.120 | הרב |  |
| 5:20.121 | יושר |  |
| 5:20.122 | לבב |  |
| 5:20.123 | בדף |  |
| 5:20.124 | ג | ' |
| 5:20.125 | ע"ב | : |
| 5:20.126 | דע |  |
| 5:20.127 | את |  |
| 5:20.128 | אלהי |  |
| 5:20.129 | אביך |  |
| 5:20.130 | כולל |  |
| 5:20.131 | החמשה |  |
| 5:20.132 | פרצופים |  |
| 5:20.133 | וכו | ' |
| 5:20.134 | ועבדהו |  |
| 5:20.135 | לזעיר |  |
| 5:20.136 | אנפין |  |
| 5:20.137 | אע"פ |  |
| 5:20.138 | שהוא |  |
| 5:20.139 | נברא |  |
| 5:20.140 | כי |  |
| 5:20.141 | לנשמתו |  |
| 5:20.142 | אתה |  |
| 5:20.143 | עובד |  |
| 5:20.144 | בזה |  |
| 5:20.145 | כי |  |
| 5:20.146 | בלעדה |  |
| 5:20.147 | אין |  |
| 5:20.148 | כל |  |
| 5:20.149 | נמצא | , |
| 5:20.150 | משום |  |
| 5:20.151 | הכי |  |
| 5:20.152 | לא |  |
| 5:20.153 | אמר |  |
| 5:20.154 | דע |  |
| 5:20.155 | אלהי |  |
| 5:20.156 | אביך |  |
| 5:20.157 | אלא |  |
| 5:20.158 | את |  |
| 5:20.159 | אלהי |  |
| 5:20.160 | אביך | " |
| 5:20.161 | וכו | '. |
| 5:20.162 | ודברי |  |
| 5:20.163 | יושר |  |
| 5:20.164 | לבב |  |
| 5:20.165 | שאמר |  |
| 5:20.166 | שהוא |  |
| 5:20.167 | נברא |  |
| 5:20.168 | מבוארים |  |
| 5:20.169 | באדרא |  |
| 5:20.170 | רבה |  |
| 5:20.171 | דף |  |
| 5:20.172 | קל"ח |  |
| 5:20.173 | ע"ב |  |
| 5:20.174 | וז"ל | : |
| 5:20.175 | ה | ' |
| 5:20.176 | פעלך |  |
| 5:20.177 | דא |  |
| 5:20.178 | זעיר |  |
| 5:20.179 | אנפין |  |
| 5:20.180 | וכו | ' |
| 5:20.181 | חייהו |  |
| 5:20.182 | למאן | , |
| 5:20.183 | חייהו |  |
| 5:20.184 | לזעיר |  |
| 5:20.185 | אנפין |  |
| 5:20.186 | דכל |  |
| 5:20.187 | נהירו |  |
| 5:20.188 | דיליה |  |
| 5:20.189 | מאינון |  |
| 5:20.190 | שנים |  |
| 5:20.191 | קדמוניות | ", |
| 5:20.192 | וכו | ' |
| 5:20.193 | ועוד |  |
| 5:20.194 | הרבה |  |
| 5:20.195 | ראיות |  |
| 5:20.196 | בלשונות |  |
| 5:20.197 | הזהר |  |
| 5:20.198 | שהוא |  |
| 5:20.199 | נברא | , |
| 5:20.200 | ועם |  |
| 5:20.201 | כל |  |
| 5:20.202 | זה |  |
| 5:20.203 | מצוה |  |
| 5:20.204 | לעבוד |  |
| 5:20.205 | אותו | . |
| 5:20.206 | ואילו |  |
| 5:20.207 | שמו |  |
| 5:20.208 | אל |  |
| 5:20.209 | לבם |  |
| 5:20.210 | דברי |  |
| 5:20.211 | רז"ל |  |
| 5:20.212 | בבראשית |  |
| 5:20.213 | רבה |  |
| 5:20.214 | ובילקוט |  |
| 5:20.215 | שהבאתי |  |
| 5:20.216 | לעיל |  |
| 5:20.217 | סי | ' |
| 5:20.218 | ל"א |  |
| 5:20.219 | כל |  |
| 5:20.220 | שיש |  |
| 5:20.221 | לו |  |
| 5:20.222 | תולדות |  |
| 5:20.223 | וכו | ' |
| 5:20.224 | לא |  |
| 5:20.225 | עבדו |  |
| 5:20.226 | לנברא | . |
| 5:20.227 | גם |  |
| 5:20.228 | הטילו |  |
| 5:20.229 | חסרון |  |
| 5:20.230 | בידיעת |  |
| 5:20.231 | השי"ת |  |
| 5:20.232 | בעתיד |  |
| 5:20.233 | כמו |  |
| 5:20.234 | שכתבתי |  |
| 5:20.235 | לעיל |  |
| 5:20.236 | בשם |  |
| 5:20.237 | האר"י |  |
| 5:20.238 | בספר |  |
| 5:20.239 | ארבע |  |
| 5:20.240 | מאות |  |
| 5:20.241 | שקל |  |
| 5:20.242 | כסף | , |
| 5:20.243 | כי |  |
| 5:20.244 | עשו |  |
| 5:20.245 | ידיעת |  |
| 5:20.246 | השי"ת |  |
| 5:20.247 | בעתיד |  |
| 5:20.248 | ידיעת |  |
| 5:20.249 | ספק |  |
| 5:20.250 | כידיעת |  |
| 5:20.251 | האצטגנים |  |
| 5:20.252 | וכן |  |
| 5:20.253 | כתב |  |
| 5:20.254 | הראב"ד |  |
| 5:20.255 | בהשגות |  |
| 5:20.256 | פ"ז |  |
| 5:20.257 | מהלכות |  |
| 5:20.258 | תשובה | . |
| 5:20.259 | וז"ל | : |
| 5:20.260 | אעפ"י |  |
| 5:20.261 | שאין |  |
| 5:20.262 | תשובה |  |
| 5:20.263 | נצחת |  |
| 5:20.264 | על |  |
| 5:20.265 | זה |  |
| 5:20.266 | טוב |  |
| 5:20.267 | לסמוך |  |
| 5:20.268 | לו |  |
| 5:20.269 | קצת |  |
| 5:20.270 | תשובה |  |
| 5:20.271 | ואומר |  |
| 5:20.272 | אם |  |
| 5:20.273 | היו |  |
| 5:20.274 | צדקת |  |
| 5:20.275 | האדם |  |
| 5:20.276 | ורשעתו |  |
| 5:20.277 | תלוים |  |
| 5:20.278 | בגזרת |  |
| 5:20.279 | הבורא |  |
| 5:20.280 | ית | ' |
| 5:20.281 | היתה |  |
| 5:20.282 | לנו |  |
| 5:20.283 | השאלה |  |
| 5:20.284 | קשה |  |
| 5:20.285 | עד |  |
| 5:20.286 | מאד | , |
| 5:20.287 | ועכשיו |  |
| 5:20.288 | שהבורא |  |
| 5:20.289 | הסיר |  |
| 5:20.290 | זאת |  |
| 5:20.291 | הממשלה |  |
| 5:20.292 | מידו |  |
| 5:20.293 | ומסרה |  |
| 5:20.294 | בידי |  |
| 5:20.295 | האדם |  |
| 5:20.296 | עצמו |  |
| 5:20.297 | אין |  |
| 5:20.298 | ידיעתו |  |
| 5:20.299 | גזרה | , |
| 5:20.300 | אבל |  |
| 5:20.301 | היא |  |
| 5:20.302 | כידיעת |  |
| 5:20.303 | האצטגנינים |  |
| 5:20.304 | שיודעים |  |
| 5:20.305 | מכח |  |
| 5:20.306 | אחר |  |
| 5:20.307 | מה |  |
| 5:20.308 | יהיה |  |
| 5:20.309 | דרכיו |  |
| 5:20.310 | של |  |
| 5:20.311 | זה | , |
| 5:20.312 | והדבר |  |
| 5:20.313 | ידוע |  |
| 5:20.314 | שכל |  |
| 5:20.315 | מקרה |  |
| 5:20.316 | האדם |  |
| 5:20.317 | מסרם |  |
| 5:20.318 | הבורא |  |
| 5:20.319 | בכל |  |
| 5:20.320 | המזלות |  |
| 5:20.321 | אלא |  |
| 5:20.322 | שנתן |  |
| 5:20.323 | בו |  |
| 5:20.324 | השכל |  |
| 5:20.325 | להיות |  |
| 5:20.326 | מחזיקו |  |
| 5:20.327 | לצאת |  |
| 5:20.328 | מתחת |  |
| 5:20.329 | המזל |  |
| 5:20.330 | והוא |  |
| 5:20.331 | הכח |  |
| 5:20.332 | הנתון |  |
| 5:20.333 | באדם | ", |
| 5:20.334 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

