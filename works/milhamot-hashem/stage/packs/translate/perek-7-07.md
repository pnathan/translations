# Translation pack: milhamot-hashem perek-7-07 (Milchamot ha-Shem 7:10–7:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-07`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **תימן** → "Yemen".
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- אשם **תלוי** → "suspended guilt offering".
- הגשת **מנחה** (argued from Gen 43 / 1 Sam 10) → "gift"; "meal offering" only
  in a sacrificial context.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- קורא **נעים / נעלם** → "pleasant reader".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 7:10

Hebrew (Wikisource): פד) וכן ידוע הוא ומפורסם שבכל ארץ התימן נוהגין לומר סדר קדושה בחזרת התפלה בנוסח שכתב הרמב"ם ז"ל הסדר התפלה שלעולם שליח צבור אומר ברכה שלישית בנוסח שיש בו ומשרתיו שואלים זה לזה איה מקום כבודו ובין בשחרית בין במנחה בין במוסף שבת ור"ח ויום טוב וחול המועד לעולם אומרין איה מקום כבודו והוא מיוסד מרבותינו אנשי כנה"ג מתקני התפלות והברכות, וכמ"ש רז"ל בספרי ומדרש רבה שאפילו מלאכים וחיות הקדש שואלים איה מקום כבודו כמו שכ' לעיל סי' נ"ב בשם הרקח. וחדשים מקרוב באו לערך מאה שנה או יותר מעט שינו הנוסח שלא לומר ומשרתיו שואלים זה לזה איה מקום כבודו רק במוסף שבת בלבד מקפידים על האומרו בשאר ימים ומשתקים אותו. וטעמם כמו שכ' הר' משנת חסידים במסכת ר"ח פ"ג משנה ט' וז"ל ומפני שאין לזעיר ונוקביה. עלייה (במוסף ר"ח) עתה למעלה ממקום אצילותם כמו בשבת. לפיכך אין אומרים ומשרתיו שואלים זה לזה איה מקום כבודו. כי משרתיו יודעים מקומו עכ"ל. הרי בהדיא מבואר בדבריו דסבירא להו לבעלי הקבלה החדשה כי מאמר רבותי' מסדרי התפלות ובעלי התלמוד שהקב"ה נעלם ורואה ואינו נראה עד שאפילו מלאכים וחיות הקדש שואלים איה מקום כבודו, הוא מצד העתקו בזמן מה לעלות למעלה ממקום משכן כבודו. דהיינו במוסף שבת דוקא שאז נעתק ועלה למעלה. ואז כשרואים מקומו המורגל להיות שוכן שם פנוי. אז יפקד מקום מושבו ואז שואלים איה מקום כבודו. אמנם בשאר זמנים שהוא שוכן במקומו כל משרתיו יודעים מקום משכן כבודו ואין צריכים לשאול עליו כמאמר יהונתן בן שאול לדוד מחר חדש ונפקדת כי יפקד מושבך. ודבריהם אלו סותרים העיקר שהזכירו הראשונים האמור ביגדל אלהים חי ומבואר ג"כ בדברי הרמב"ם בפ' א' מיסודי התורה ובמורה. ובס' אהל יעקב פרשת וירא על מאמר אברהם אבינו לשי"ת אל נא תעבור מעל עבדך כתב דבהשקפה ראשונה יפלא מאד על אברהם אבינו איך ערב אל לבו לומר לפני השי"ת אל נא תעבור מעל עבדך. ובגמרא שבת קכ"ז אר"א מדת ב"ו אין קטן יכול לומר לגדול המתן עד שאבוא אצלך, ואלו בהקב"ה כתיב ויאמר אל נא תעבר מעל עבדך. אבל הבנת הענין הוא עם המאומת והמקובל אצלינו מפי קדמונינו ע"ה ובראשם הר' המורה, כי הקב"ה נעלה ומרומם מכל שינוי ותנועה והעתק מקום ובשגם מצינו בפרשת וירא וילך ה' וכו' אין הרצון בו חלילה שהקב"ה נעתק וסר ממקומו ובא להתראות אל מקום מצב מי שיראה אליו. אבל כל התקרבותו ית' והתגלותו לאדם תלוי אך בהתקרבות האדם אליו כמאמר אני לדודי ודודי לי, וזהו שאמר אני לדודי ועלי תשוקתו, כלומר תשוקתו ית' אלינו הוא דבר תלוי בידינו, ואברהם אבינו בעוצם דבקותו בו ית' נגלה אליו, ובהזדמן האורחים אצלו וצידד להכניסם ולטבוח טבח והכן ירא לנפשו שלא יפול מהדבק בשי"ת ושאל ממנו אל נא תעבור מעל עבדך, לבל יסיר הקב"ה עזרתו מאתו ולא ינתק מדבקותו בטרדת הכנסת אורחים וכו' עיין שם.

| id | word | punctuation after |
|---|---|---|
| 7:10.1 | פד | ) |
| 7:10.2 | וכן |  |
| 7:10.3 | ידוע |  |
| 7:10.4 | הוא |  |
| 7:10.5 | ומפורסם |  |
| 7:10.6 | שבכל |  |
| 7:10.7 | ארץ |  |
| 7:10.8 | התימן |  |
| 7:10.9 | נוהגין |  |
| 7:10.10 | לומר |  |
| 7:10.11 | סדר |  |
| 7:10.12 | קדושה |  |
| 7:10.13 | בחזרת |  |
| 7:10.14 | התפלה |  |
| 7:10.15 | בנוסח |  |
| 7:10.16 | שכתב |  |
| 7:10.17 | הרמב"ם |  |
| 7:10.18 | ז"ל |  |
| 7:10.19 | הסדר |  |
| 7:10.20 | התפלה |  |
| 7:10.21 | שלעולם |  |
| 7:10.22 | שליח |  |
| 7:10.23 | צבור |  |
| 7:10.24 | אומר |  |
| 7:10.25 | ברכה |  |
| 7:10.26 | שלישית |  |
| 7:10.27 | בנוסח |  |
| 7:10.28 | שיש |  |
| 7:10.29 | בו |  |
| 7:10.30 | ומשרתיו |  |
| 7:10.31 | שואלים |  |
| 7:10.32 | זה |  |
| 7:10.33 | לזה |  |
| 7:10.34 | איה |  |
| 7:10.35 | מקום |  |
| 7:10.36 | כבודו |  |
| 7:10.37 | ובין |  |
| 7:10.38 | בשחרית |  |
| 7:10.39 | בין |  |
| 7:10.40 | במנחה |  |
| 7:10.41 | בין |  |
| 7:10.42 | במוסף |  |
| 7:10.43 | שבת |  |
| 7:10.44 | ור"ח |  |
| 7:10.45 | ויום |  |
| 7:10.46 | טוב |  |
| 7:10.47 | וחול |  |
| 7:10.48 | המועד |  |
| 7:10.49 | לעולם |  |
| 7:10.50 | אומרין |  |
| 7:10.51 | איה |  |
| 7:10.52 | מקום |  |
| 7:10.53 | כבודו |  |
| 7:10.54 | והוא |  |
| 7:10.55 | מיוסד |  |
| 7:10.56 | מרבותינו |  |
| 7:10.57 | אנשי |  |
| 7:10.58 | כנה"ג |  |
| 7:10.59 | מתקני |  |
| 7:10.60 | התפלות |  |
| 7:10.61 | והברכות | , |
| 7:10.62 | וכמ"ש |  |
| 7:10.63 | רז"ל |  |
| 7:10.64 | בספרי |  |
| 7:10.65 | ומדרש |  |
| 7:10.66 | רבה |  |
| 7:10.67 | שאפילו |  |
| 7:10.68 | מלאכים |  |
| 7:10.69 | וחיות |  |
| 7:10.70 | הקדש |  |
| 7:10.71 | שואלים |  |
| 7:10.72 | איה |  |
| 7:10.73 | מקום |  |
| 7:10.74 | כבודו |  |
| 7:10.75 | כמו |  |
| 7:10.76 | שכ | ' |
| 7:10.77 | לעיל |  |
| 7:10.78 | סי | ' |
| 7:10.79 | נ"ב |  |
| 7:10.80 | בשם |  |
| 7:10.81 | הרקח | . |
| 7:10.82 | וחדשים |  |
| 7:10.83 | מקרוב |  |
| 7:10.84 | באו |  |
| 7:10.85 | לערך |  |
| 7:10.86 | מאה |  |
| 7:10.87 | שנה |  |
| 7:10.88 | או |  |
| 7:10.89 | יותר |  |
| 7:10.90 | מעט |  |
| 7:10.91 | שינו |  |
| 7:10.92 | הנוסח |  |
| 7:10.93 | שלא |  |
| 7:10.94 | לומר |  |
| 7:10.95 | ומשרתיו |  |
| 7:10.96 | שואלים |  |
| 7:10.97 | זה |  |
| 7:10.98 | לזה |  |
| 7:10.99 | איה |  |
| 7:10.100 | מקום |  |
| 7:10.101 | כבודו |  |
| 7:10.102 | רק |  |
| 7:10.103 | במוסף |  |
| 7:10.104 | שבת |  |
| 7:10.105 | בלבד |  |
| 7:10.106 | מקפידים |  |
| 7:10.107 | על |  |
| 7:10.108 | האומרו |  |
| 7:10.109 | בשאר |  |
| 7:10.110 | ימים |  |
| 7:10.111 | ומשתקים |  |
| 7:10.112 | אותו | . |
| 7:10.113 | וטעמם |  |
| 7:10.114 | כמו |  |
| 7:10.115 | שכ | ' |
| 7:10.116 | הר | ' |
| 7:10.117 | משנת |  |
| 7:10.118 | חסידים |  |
| 7:10.119 | במסכת |  |
| 7:10.120 | ר"ח |  |
| 7:10.121 | פ"ג |  |
| 7:10.122 | משנה |  |
| 7:10.123 | ט | ' |
| 7:10.124 | וז"ל |  |
| 7:10.125 | ומפני |  |
| 7:10.126 | שאין |  |
| 7:10.127 | לזעיר |  |
| 7:10.128 | ונוקביה | . |
| 7:10.129 | עלייה |  |
| 7:10.130 | במוסף |  |
| 7:10.131 | ר"ח | ) |
| 7:10.132 | עתה |  |
| 7:10.133 | למעלה |  |
| 7:10.134 | ממקום |  |
| 7:10.135 | אצילותם |  |
| 7:10.136 | כמו |  |
| 7:10.137 | בשבת | . |
| 7:10.138 | לפיכך |  |
| 7:10.139 | אין |  |
| 7:10.140 | אומרים |  |
| 7:10.141 | ומשרתיו |  |
| 7:10.142 | שואלים |  |
| 7:10.143 | זה |  |
| 7:10.144 | לזה |  |
| 7:10.145 | איה |  |
| 7:10.146 | מקום |  |
| 7:10.147 | כבודו | . |
| 7:10.148 | כי |  |
| 7:10.149 | משרתיו |  |
| 7:10.150 | יודעים |  |
| 7:10.151 | מקומו |  |
| 7:10.152 | עכ"ל | . |
| 7:10.153 | הרי |  |
| 7:10.154 | בהדיא |  |
| 7:10.155 | מבואר |  |
| 7:10.156 | בדבריו |  |
| 7:10.157 | דסבירא |  |
| 7:10.158 | להו |  |
| 7:10.159 | לבעלי |  |
| 7:10.160 | הקבלה |  |
| 7:10.161 | החדשה |  |
| 7:10.162 | כי |  |
| 7:10.163 | מאמר |  |
| 7:10.164 | רבותי | ' |
| 7:10.165 | מסדרי |  |
| 7:10.166 | התפלות |  |
| 7:10.167 | ובעלי |  |
| 7:10.168 | התלמוד |  |
| 7:10.169 | שהקב"ה |  |
| 7:10.170 | נעלם |  |
| 7:10.171 | ורואה |  |
| 7:10.172 | ואינו |  |
| 7:10.173 | נראה |  |
| 7:10.174 | עד |  |
| 7:10.175 | שאפילו |  |
| 7:10.176 | מלאכים |  |
| 7:10.177 | וחיות |  |
| 7:10.178 | הקדש |  |
| 7:10.179 | שואלים |  |
| 7:10.180 | איה |  |
| 7:10.181 | מקום |  |
| 7:10.182 | כבודו | , |
| 7:10.183 | הוא |  |
| 7:10.184 | מצד |  |
| 7:10.185 | העתקו |  |
| 7:10.186 | בזמן |  |
| 7:10.187 | מה |  |
| 7:10.188 | לעלות |  |
| 7:10.189 | למעלה |  |
| 7:10.190 | ממקום |  |
| 7:10.191 | משכן |  |
| 7:10.192 | כבודו | . |
| 7:10.193 | דהיינו |  |
| 7:10.194 | במוסף |  |
| 7:10.195 | שבת |  |
| 7:10.196 | דוקא |  |
| 7:10.197 | שאז |  |
| 7:10.198 | נעתק |  |
| 7:10.199 | ועלה |  |
| 7:10.200 | למעלה | . |
| 7:10.201 | ואז |  |
| 7:10.202 | כשרואים |  |
| 7:10.203 | מקומו |  |
| 7:10.204 | המורגל |  |
| 7:10.205 | להיות |  |
| 7:10.206 | שוכן |  |
| 7:10.207 | שם |  |
| 7:10.208 | פנוי | . |
| 7:10.209 | אז |  |
| 7:10.210 | יפקד |  |
| 7:10.211 | מקום |  |
| 7:10.212 | מושבו |  |
| 7:10.213 | ואז |  |
| 7:10.214 | שואלים |  |
| 7:10.215 | איה |  |
| 7:10.216 | מקום |  |
| 7:10.217 | כבודו | . |
| 7:10.218 | אמנם |  |
| 7:10.219 | בשאר |  |
| 7:10.220 | זמנים |  |
| 7:10.221 | שהוא |  |
| 7:10.222 | שוכן |  |
| 7:10.223 | במקומו |  |
| 7:10.224 | כל |  |
| 7:10.225 | משרתיו |  |
| 7:10.226 | יודעים |  |
| 7:10.227 | מקום |  |
| 7:10.228 | משכן |  |
| 7:10.229 | כבודו |  |
| 7:10.230 | ואין |  |
| 7:10.231 | צריכים |  |
| 7:10.232 | לשאול |  |
| 7:10.233 | עליו |  |
| 7:10.234 | כמאמר |  |
| 7:10.235 | יהונתן |  |
| 7:10.236 | בן |  |
| 7:10.237 | שאול |  |
| 7:10.238 | לדוד |  |
| 7:10.239 | מחר |  |
| 7:10.240 | חדש |  |
| 7:10.241 | ונפקדת |  |
| 7:10.242 | כי |  |
| 7:10.243 | יפקד |  |
| 7:10.244 | מושבך | . |
| 7:10.245 | ודבריהם |  |
| 7:10.246 | אלו |  |
| 7:10.247 | סותרים |  |
| 7:10.248 | העיקר |  |
| 7:10.249 | שהזכירו |  |
| 7:10.250 | הראשונים |  |
| 7:10.251 | האמור |  |
| 7:10.252 | ביגדל |  |
| 7:10.253 | אלהים |  |
| 7:10.254 | חי |  |
| 7:10.255 | ומבואר |  |
| 7:10.256 | ג"כ |  |
| 7:10.257 | בדברי |  |
| 7:10.258 | הרמב"ם |  |
| 7:10.259 | בפ | ' |
| 7:10.260 | א | ' |
| 7:10.261 | מיסודי |  |
| 7:10.262 | התורה |  |
| 7:10.263 | ובמורה | . |
| 7:10.264 | ובס | ' |
| 7:10.265 | אהל |  |
| 7:10.266 | יעקב |  |
| 7:10.267 | פרשת |  |
| 7:10.268 | וירא |  |
| 7:10.269 | על |  |
| 7:10.270 | מאמר |  |
| 7:10.271 | אברהם |  |
| 7:10.272 | אבינו |  |
| 7:10.273 | לשי"ת |  |
| 7:10.274 | אל |  |
| 7:10.275 | נא |  |
| 7:10.276 | תעבור |  |
| 7:10.277 | מעל |  |
| 7:10.278 | עבדך |  |
| 7:10.279 | כתב |  |
| 7:10.280 | דבהשקפה |  |
| 7:10.281 | ראשונה |  |
| 7:10.282 | יפלא |  |
| 7:10.283 | מאד |  |
| 7:10.284 | על |  |
| 7:10.285 | אברהם |  |
| 7:10.286 | אבינו |  |
| 7:10.287 | איך |  |
| 7:10.288 | ערב |  |
| 7:10.289 | אל |  |
| 7:10.290 | לבו |  |
| 7:10.291 | לומר |  |
| 7:10.292 | לפני |  |
| 7:10.293 | השי"ת |  |
| 7:10.294 | אל |  |
| 7:10.295 | נא |  |
| 7:10.296 | תעבור |  |
| 7:10.297 | מעל |  |
| 7:10.298 | עבדך | . |
| 7:10.299 | ובגמרא |  |
| 7:10.300 | שבת |  |
| 7:10.301 | קכ"ז |  |
| 7:10.302 | אר"א |  |
| 7:10.303 | מדת |  |
| 7:10.304 | ב"ו |  |
| 7:10.305 | אין |  |
| 7:10.306 | קטן |  |
| 7:10.307 | יכול |  |
| 7:10.308 | לומר |  |
| 7:10.309 | לגדול |  |
| 7:10.310 | המתן |  |
| 7:10.311 | עד |  |
| 7:10.312 | שאבוא |  |
| 7:10.313 | אצלך | , |
| 7:10.314 | ואלו |  |
| 7:10.315 | בהקב"ה |  |
| 7:10.316 | כתיב |  |
| 7:10.317 | ויאמר |  |
| 7:10.318 | אל |  |
| 7:10.319 | נא |  |
| 7:10.320 | תעבר |  |
| 7:10.321 | מעל |  |
| 7:10.322 | עבדך | . |
| 7:10.323 | אבל |  |
| 7:10.324 | הבנת |  |
| 7:10.325 | הענין |  |
| 7:10.326 | הוא |  |
| 7:10.327 | עם |  |
| 7:10.328 | המאומת |  |
| 7:10.329 | והמקובל |  |
| 7:10.330 | אצלינו |  |
| 7:10.331 | מפי |  |
| 7:10.332 | קדמונינו |  |
| 7:10.333 | ע"ה |  |
| 7:10.334 | ובראשם |  |
| 7:10.335 | הר | ' |
| 7:10.336 | המורה | , |
| 7:10.337 | כי |  |
| 7:10.338 | הקב"ה |  |
| 7:10.339 | נעלה |  |
| 7:10.340 | ומרומם |  |
| 7:10.341 | מכל |  |
| 7:10.342 | שינוי |  |
| 7:10.343 | ותנועה |  |
| 7:10.344 | והעתק |  |
| 7:10.345 | מקום |  |
| 7:10.346 | ובשגם |  |
| 7:10.347 | מצינו |  |
| 7:10.348 | בפרשת |  |
| 7:10.349 | וירא |  |
| 7:10.350 | וילך |  |
| 7:10.351 | ה | ' |
| 7:10.352 | וכו | ' |
| 7:10.353 | אין |  |
| 7:10.354 | הרצון |  |
| 7:10.355 | בו |  |
| 7:10.356 | חלילה |  |
| 7:10.357 | שהקב"ה |  |
| 7:10.358 | נעתק |  |
| 7:10.359 | וסר |  |
| 7:10.360 | ממקומו |  |
| 7:10.361 | ובא |  |
| 7:10.362 | להתראות |  |
| 7:10.363 | אל |  |
| 7:10.364 | מקום |  |
| 7:10.365 | מצב |  |
| 7:10.366 | מי |  |
| 7:10.367 | שיראה |  |
| 7:10.368 | אליו | . |
| 7:10.369 | אבל |  |
| 7:10.370 | כל |  |
| 7:10.371 | התקרבותו |  |
| 7:10.372 | ית | ' |
| 7:10.373 | והתגלותו |  |
| 7:10.374 | לאדם |  |
| 7:10.375 | תלוי |  |
| 7:10.376 | אך |  |
| 7:10.377 | בהתקרבות |  |
| 7:10.378 | האדם |  |
| 7:10.379 | אליו |  |
| 7:10.380 | כמאמר |  |
| 7:10.381 | אני |  |
| 7:10.382 | לדודי |  |
| 7:10.383 | ודודי |  |
| 7:10.384 | לי | , |
| 7:10.385 | וזהו |  |
| 7:10.386 | שאמר |  |
| 7:10.387 | אני |  |
| 7:10.388 | לדודי |  |
| 7:10.389 | ועלי |  |
| 7:10.390 | תשוקתו | , |
| 7:10.391 | כלומר |  |
| 7:10.392 | תשוקתו |  |
| 7:10.393 | ית | ' |
| 7:10.394 | אלינו |  |
| 7:10.395 | הוא |  |
| 7:10.396 | דבר |  |
| 7:10.397 | תלוי |  |
| 7:10.398 | בידינו | , |
| 7:10.399 | ואברהם |  |
| 7:10.400 | אבינו |  |
| 7:10.401 | בעוצם |  |
| 7:10.402 | דבקותו |  |
| 7:10.403 | בו |  |
| 7:10.404 | ית | ' |
| 7:10.405 | נגלה |  |
| 7:10.406 | אליו | , |
| 7:10.407 | ובהזדמן |  |
| 7:10.408 | האורחים |  |
| 7:10.409 | אצלו |  |
| 7:10.410 | וצידד |  |
| 7:10.411 | להכניסם |  |
| 7:10.412 | ולטבוח |  |
| 7:10.413 | טבח |  |
| 7:10.414 | והכן |  |
| 7:10.415 | ירא |  |
| 7:10.416 | לנפשו |  |
| 7:10.417 | שלא |  |
| 7:10.418 | יפול |  |
| 7:10.419 | מהדבק |  |
| 7:10.420 | בשי"ת |  |
| 7:10.421 | ושאל |  |
| 7:10.422 | ממנו |  |
| 7:10.423 | אל |  |
| 7:10.424 | נא |  |
| 7:10.425 | תעבור |  |
| 7:10.426 | מעל |  |
| 7:10.427 | עבדך | , |
| 7:10.428 | לבל |  |
| 7:10.429 | יסיר |  |
| 7:10.430 | הקב"ה |  |
| 7:10.431 | עזרתו |  |
| 7:10.432 | מאתו |  |
| 7:10.433 | ולא |  |
| 7:10.434 | ינתק |  |
| 7:10.435 | מדבקותו |  |
| 7:10.436 | בטרדת |  |
| 7:10.437 | הכנסת |  |
| 7:10.438 | אורחים |  |
| 7:10.439 | וכו | ' |
| 7:10.440 | עיין |  |
| 7:10.441 | שם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

