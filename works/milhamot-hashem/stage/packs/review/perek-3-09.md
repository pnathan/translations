# Review pack: milhamot-hashem perek-3-09 (Milchamot ha-Shem 3:14–3:15)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-09`.

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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **רמ"א** → "Rema" (authority, transliterated; note).
- **רדב"ז** → "Radbaz" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- **מהר"ם** אלשקר → "Maharam al-Shakar" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **מרן** → "Maran" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אלקבולי** → "Alkaboli", Alkabetz in the note; one spelling across the work.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
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
## Terms

- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
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
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **מין / מינות** → "heretic / heresy".
- כופר **בעיקר** → "denies the fundamental principle".
- ימי **הבינים** → "the Middle Ages".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ג"כ** → "also".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- אין לנו **להחמיר** → "we may not be stringent".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **סכלות** והוללות → "folly and madness".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- קורא **נעים / נעלם** → "pleasant reader".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 3:14

Hebrew (Wikisource): יד) ראה והבן ידידנו קורא נעלם איך כמה"ר תם הנז"ל צוח ככרוכיא על העוסקים בקבלה החדשה, כי פרצו גדריה של תורה, וערו ערו יסודותיה ושרשיה ביחוד השי"ת כמו שיתבאר להלן, ותחת אשר חשבו להתקרב לשי"ת תעו מן הדרך ובאו והאמין באלהים רבים אחרים ולעבדם! וכדברי ר' תם ן' יחיא כ' גם כן מהר"י אלבו בס' העיקרים פרק כ"ח משני. ועוד האריך בספר רביד הזהב והביא ג"כ דברי הר' הגדול מהרש"ל בתשובותיו ובחבור ים של שלמה על המתיהרים בחידושים שלומדים בקבלה, וכתב דמאחר שנחתם התלמוד אין לנו להחמיר נגד הגמרא והוי כמו מינות כמבואר בחק יעקב סי' תפ"ט ס"ק י"א באורך. וכל זה הוא לפי ההנחה והיחוס שייחסו הזהר לרשב"י התנא, אך אחרי הבקורת שבקרו וחקרו חכמי לב אחריו מבטן מי יצאו הדברים הנזכרים בו ואימתי יצא החבור הנזכר ומיד מי יצא. ונתבאר באר היטב כי הוא חיבור אחד האחרונים ולא יצאו אותן הרעיונים מרשב"י התנא ע"ה כלל, כי נמצא בו כמה דברים בעיקרי אמונת תורתינו הקדושה היפך דברי רשב"י בתלמוד ובמדרשי רז"ל המפורסמים בישראל, ודאי שאין לסמוך עליו בשום דבר הנוגע למנהגי ישראל שנהגו בהם קדמונינו. ולא יפה עשו חכמי ימי הבינים ששינו עלינו כמה דברים בסידורי התפלות על פי הזהר והמקובלים, כי שינויים ההם יש מהם מביאים לידי מינות להאמין ברשויות הרבה ולעבוד אלהים אחרים כמו שיתבאר להלן.

| id | word | punctuation after |
|---|---|---|
| 3:14.1 | יד | ) |
| 3:14.2 | ראה |  |
| 3:14.3 | והבן |  |
| 3:14.4 | ידידנו |  |
| 3:14.5 | קורא |  |
| 3:14.6 | נעלם |  |
| 3:14.7 | איך |  |
| 3:14.8 | כמה"ר |  |
| 3:14.9 | תם |  |
| 3:14.10 | הנז"ל |  |
| 3:14.11 | צוח |  |
| 3:14.12 | ככרוכיא |  |
| 3:14.13 | על |  |
| 3:14.14 | העוסקים |  |
| 3:14.15 | בקבלה |  |
| 3:14.16 | החדשה | , |
| 3:14.17 | כי |  |
| 3:14.18 | פרצו |  |
| 3:14.19 | גדריה |  |
| 3:14.20 | של |  |
| 3:14.21 | תורה | , |
| 3:14.22 | וערו |  |
| 3:14.23 | ערו |  |
| 3:14.24 | יסודותיה |  |
| 3:14.25 | ושרשיה |  |
| 3:14.26 | ביחוד |  |
| 3:14.27 | השי"ת |  |
| 3:14.28 | כמו |  |
| 3:14.29 | שיתבאר |  |
| 3:14.30 | להלן | , |
| 3:14.31 | ותחת |  |
| 3:14.32 | אשר |  |
| 3:14.33 | חשבו |  |
| 3:14.34 | להתקרב |  |
| 3:14.35 | לשי"ת |  |
| 3:14.36 | תעו |  |
| 3:14.37 | מן |  |
| 3:14.38 | הדרך |  |
| 3:14.39 | ובאו |  |
| 3:14.40 | והאמין |  |
| 3:14.41 | באלהים |  |
| 3:14.42 | רבים |  |
| 3:14.43 | אחרים |  |
| 3:14.44 | ולעבדם | ! |
| 3:14.45 | וכדברי |  |
| 3:14.46 | ר | ' |
| 3:14.47 | תם |  |
| 3:14.48 | ן | ' |
| 3:14.49 | יחיא |  |
| 3:14.50 | כ | ' |
| 3:14.51 | גם |  |
| 3:14.52 | כן |  |
| 3:14.53 | מהר"י |  |
| 3:14.54 | אלבו |  |
| 3:14.55 | בס | ' |
| 3:14.56 | העיקרים |  |
| 3:14.57 | פרק |  |
| 3:14.58 | כ"ח |  |
| 3:14.59 | משני | . |
| 3:14.60 | ועוד |  |
| 3:14.61 | האריך |  |
| 3:14.62 | בספר |  |
| 3:14.63 | רביד |  |
| 3:14.64 | הזהב |  |
| 3:14.65 | והביא |  |
| 3:14.66 | ג"כ |  |
| 3:14.67 | דברי |  |
| 3:14.68 | הר | ' |
| 3:14.69 | הגדול |  |
| 3:14.70 | מהרש"ל |  |
| 3:14.71 | בתשובותיו |  |
| 3:14.72 | ובחבור |  |
| 3:14.73 | ים |  |
| 3:14.74 | של |  |
| 3:14.75 | שלמה |  |
| 3:14.76 | על |  |
| 3:14.77 | המתיהרים |  |
| 3:14.78 | בחידושים |  |
| 3:14.79 | שלומדים |  |
| 3:14.80 | בקבלה | , |
| 3:14.81 | וכתב |  |
| 3:14.82 | דמאחר |  |
| 3:14.83 | שנחתם |  |
| 3:14.84 | התלמוד |  |
| 3:14.85 | אין |  |
| 3:14.86 | לנו |  |
| 3:14.87 | להחמיר |  |
| 3:14.88 | נגד |  |
| 3:14.89 | הגמרא |  |
| 3:14.90 | והוי |  |
| 3:14.91 | כמו |  |
| 3:14.92 | מינות |  |
| 3:14.93 | כמבואר |  |
| 3:14.94 | בחק |  |
| 3:14.95 | יעקב |  |
| 3:14.96 | סי | ' |
| 3:14.97 | תפ"ט |  |
| 3:14.98 | ס"ק |  |
| 3:14.99 | י"א |  |
| 3:14.100 | באורך | . |
| 3:14.101 | וכל |  |
| 3:14.102 | זה |  |
| 3:14.103 | הוא |  |
| 3:14.104 | לפי |  |
| 3:14.105 | ההנחה |  |
| 3:14.106 | והיחוס |  |
| 3:14.107 | שייחסו |  |
| 3:14.108 | הזהר |  |
| 3:14.109 | לרשב"י |  |
| 3:14.110 | התנא | , |
| 3:14.111 | אך |  |
| 3:14.112 | אחרי |  |
| 3:14.113 | הבקורת |  |
| 3:14.114 | שבקרו |  |
| 3:14.115 | וחקרו |  |
| 3:14.116 | חכמי |  |
| 3:14.117 | לב |  |
| 3:14.118 | אחריו |  |
| 3:14.119 | מבטן |  |
| 3:14.120 | מי |  |
| 3:14.121 | יצאו |  |
| 3:14.122 | הדברים |  |
| 3:14.123 | הנזכרים |  |
| 3:14.124 | בו |  |
| 3:14.125 | ואימתי |  |
| 3:14.126 | יצא |  |
| 3:14.127 | החבור |  |
| 3:14.128 | הנזכר |  |
| 3:14.129 | ומיד |  |
| 3:14.130 | מי |  |
| 3:14.131 | יצא | . |
| 3:14.132 | ונתבאר |  |
| 3:14.133 | באר |  |
| 3:14.134 | היטב |  |
| 3:14.135 | כי |  |
| 3:14.136 | הוא |  |
| 3:14.137 | חיבור |  |
| 3:14.138 | אחד |  |
| 3:14.139 | האחרונים |  |
| 3:14.140 | ולא |  |
| 3:14.141 | יצאו |  |
| 3:14.142 | אותן |  |
| 3:14.143 | הרעיונים |  |
| 3:14.144 | מרשב"י |  |
| 3:14.145 | התנא |  |
| 3:14.146 | ע"ה |  |
| 3:14.147 | כלל | , |
| 3:14.148 | כי |  |
| 3:14.149 | נמצא |  |
| 3:14.150 | בו |  |
| 3:14.151 | כמה |  |
| 3:14.152 | דברים |  |
| 3:14.153 | בעיקרי |  |
| 3:14.154 | אמונת |  |
| 3:14.155 | תורתינו |  |
| 3:14.156 | הקדושה |  |
| 3:14.157 | היפך |  |
| 3:14.158 | דברי |  |
| 3:14.159 | רשב"י |  |
| 3:14.160 | בתלמוד |  |
| 3:14.161 | ובמדרשי |  |
| 3:14.162 | רז"ל |  |
| 3:14.163 | המפורסמים |  |
| 3:14.164 | בישראל | , |
| 3:14.165 | ודאי |  |
| 3:14.166 | שאין |  |
| 3:14.167 | לסמוך |  |
| 3:14.168 | עליו |  |
| 3:14.169 | בשום |  |
| 3:14.170 | דבר |  |
| 3:14.171 | הנוגע |  |
| 3:14.172 | למנהגי |  |
| 3:14.173 | ישראל |  |
| 3:14.174 | שנהגו |  |
| 3:14.175 | בהם |  |
| 3:14.176 | קדמונינו | . |
| 3:14.177 | ולא |  |
| 3:14.178 | יפה |  |
| 3:14.179 | עשו |  |
| 3:14.180 | חכמי |  |
| 3:14.181 | ימי |  |
| 3:14.182 | הבינים |  |
| 3:14.183 | ששינו |  |
| 3:14.184 | עלינו |  |
| 3:14.185 | כמה |  |
| 3:14.186 | דברים |  |
| 3:14.187 | בסידורי |  |
| 3:14.188 | התפלות |  |
| 3:14.189 | על |  |
| 3:14.190 | פי |  |
| 3:14.191 | הזהר |  |
| 3:14.192 | והמקובלים | , |
| 3:14.193 | כי |  |
| 3:14.194 | שינויים |  |
| 3:14.195 | ההם |  |
| 3:14.196 | יש |  |
| 3:14.197 | מהם |  |
| 3:14.198 | מביאים |  |
| 3:14.199 | לידי |  |
| 3:14.200 | מינות |  |
| 3:14.201 | להאמין |  |
| 3:14.202 | ברשויות |  |
| 3:14.203 | הרבה |  |
| 3:14.204 | ולעבוד |  |
| 3:14.205 | אלהים |  |
| 3:14.206 | אחרים |  |
| 3:14.207 | כמו |  |
| 3:14.208 | שיתבאר |  |
| 3:14.209 | להלן | . |


### Translation under review (3:14)

English: §14. See and understand, our friend the hidden reader, how our honored Rabbi Tam, the aforementioned, cried out like a herald against those who engage in the new Kabbalah, for they breached the fences of the Torah, and laid bare, laid bare, its foundations and its roots in the unity of the Name, may He be blessed, as will be explained below. Instead of thinking to draw near to the Name, may He be blessed, they strayed from the way and came to believe in many other gods, and to worship them! As the words of Rabbi Tam ben Yachya, Mahari Albo also wrote in the Ikkarim, chapter 28, of the second. And further he expounded at length in the book Ravid ha-Zahav, and also brought the words of the great Rabbi Maharshal, in his responses and in the work Yam shel Shlomo, against those who boast of the novelties they study in Kabbalah. And he wrote that since the Talmud was sealed, there is no need for us to be stringent against the Gemara, and it is like heresy, as explained at length in Chok Ya'akov, section 489, subsection 11. All this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai, the Tanna, But after the criticism which the wise of heart examined and investigated after him, from whose womb the things mentioned in it came out, and when the work mentioned came out, and from whose hand it came out. And it was shown clearly and well that it is a single work of the latter ones, and those ideas did not come out of Rabbi Shimon bar Yochai, the Tanna, peace be upon him, at all. for in it are found many things in the fundamentals of the faith of our holy Torah contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory, famous in Israel, certainly one must not rely on it in any matter touching the customs of Israel that our forebears practised. And the sages of the Middle Ages did not do well, who changed upon us several things in the prayer books, according to the Zohar and the Kabbalists, for some of those changes lead to heresy, to believe in many powers and to worship other gods, as will be explained below.

Literal: §14 see and understand our friend the reader hidden how as the honored Rabbi Tam the aforementioned cried out like a herald against those who engage in Kabbalah the new for breached the fences of the Torah and laid bare laid bare its foundations and its roots in the unity the Name, may He be blessed as will be explained below and instead of thought to draw near to the Name, may He be blessed strayed from the way and came and believed in gods many other and to worship them and as the words of Rabbi Tam ben Yachya wrote also likewise Mahari Albo in the book Ikkarim chapter 28 of the second and further expounded at length in the book Ravid ha-Zahav and brought also the words the Rabbi great Maharshal in his responses and in the work Yam of Shlomo against those who boast of novelties which they study in Kabbalah and he wrote that since was sealed the Talmud there is no for us to be stringent against the Gemara and it is like heresy as explained in Chok Ya'akov section 489 subsection 11 at length and all this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai the Tanna but after the criticism which they examined and investigated the wise of heart after him from the womb of whose came out the things mentioned in it and when came out the work mentioned and from the hand of whom came out and it was shown clearly well that it a work single of the latter ones and not came out those ideas from Rabbi Shimon bar Yochai the Tanna peace be upon him at all for is found in it many things in the fundamentals of the faith of our Torah holy contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory famous in Israel certainly one must not rely on it in any matter touching the customs of Israel which they practised in our forebears and not well did the sages of the days of the Middle Ages who changed upon us several things in the prayer books the prayers according to the word of the Zohar and the Kabbalists for changes those there are of them that lead to heresy to believe in powers many and to worship gods other as will be explained below

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:14.1 | §14 |  |  |  |
| 3:14.2 | see |  |  |  |
| 3:14.3 | and understand |  |  |  |
| 3:14.4 | our friend |  |  |  |
| 3:14.5 | the reader |  |  |  |
| 3:14.6 | hidden |  |  |  |
| 3:14.7 | how |  |  |  |
| 3:14.8 | as the honored Rabbi |  |  | כמוהר"ר: an honorific; the author gives it to Rabbi Tam. |
| 3:14.9 | Tam |  |  |  |
| 3:14.10 | the aforementioned |  |  |  |
| 3:14.11 | cried out |  |  |  |
| 3:14.12 | like a herald |  |  | ככרוכיא: a herald or town crier (karukya). |
| 3:14.13 | against |  |  |  |
| 3:14.14 | those who engage |  |  |  |
| 3:14.15 | in Kabbalah |  |  |  |
| 3:14.16 | the new |  |  |  |
| 3:14.17 | for |  |  |  |
| 3:14.18 | breached |  |  |  |
| 3:14.19 | the fences |  |  |  |
| 3:14.20 | of |  |  |  |
| 3:14.21 | the Torah |  |  |  |
| 3:14.22 | and laid bare |  |  | וערו ערו: a doubled verb; the English keeps it doubled. |
| 3:14.23 | laid bare |  |  |  |
| 3:14.24 | its foundations |  |  |  |
| 3:14.25 | and its roots |  |  |  |
| 3:14.26 | in the unity |  |  |  |
| 3:14.27 | the Name, may He be blessed |  |  | השי"ת: the Name, may He be blessed (fixed form). |
| 3:14.28 | as |  |  |  |
| 3:14.29 | will be explained |  |  |  |
| 3:14.30 | below |  |  |  |
| 3:14.31 | and instead |  |  |  |
| 3:14.32 | of |  |  |  |
| 3:14.33 | thought |  |  |  |
| 3:14.34 | to draw near |  |  |  |
| 3:14.35 | to the Name, may He be blessed |  |  |  |
| 3:14.36 | strayed |  |  |  |
| 3:14.37 | from |  |  |  |
| 3:14.38 | the way |  |  |  |
| 3:14.39 | and came |  |  |  |
| 3:14.40 | and believed |  |  |  |
| 3:14.41 | in gods |  |  |  |
| 3:14.42 | many |  |  |  |
| 3:14.43 | other |  |  |  |
| 3:14.44 | and to worship them |  |  |  |
| 3:14.45 | and as the words of |  |  |  |
| 3:14.46 | Rabbi |  |  |  |
| 3:14.47 | Tam |  |  |  |
| 3:14.48 | ben |  |  |  |
| 3:14.49 | Yachya |  |  |  |
| 3:14.50 | wrote |  |  |  |
| 3:14.51 | also |  |  |  |
| 3:14.52 | likewise |  |  |  |
| 3:14.53 | Mahari |  |  | מהר"י אלבו: Morenu ha-Rav Rabbi Y. Albo, the author of the Ikkarim; the text names him only by the acronym. |
| 3:14.54 | Albo |  |  |  |
| 3:14.55 | in the book |  |  |  |
| 3:14.56 | Ikkarim |  |  |  |
| 3:14.57 | chapter |  |  |  |
| 3:14.58 | 28 |  |  |  |
| 3:14.59 | of the second |  |  | משני: the second treatise of the Ikkarim. |
| 3:14.60 | and further |  |  |  |
| 3:14.61 | expounded at length |  |  |  |
| 3:14.62 | in the book |  |  |  |
| 3:14.63 | Ravid |  |  |  |
| 3:14.64 | ha-Zahav |  |  |  |
| 3:14.65 | and brought |  |  |  |
| 3:14.66 | also |  |  |  |
| 3:14.67 | the words |  |  |  |
| 3:14.68 | the Rabbi |  |  |  |
| 3:14.69 | great |  |  |  |
| 3:14.70 | Maharshal |  |  |  |
| 3:14.71 | in his responses |  |  |  |
| 3:14.72 | and in the work |  |  |  |
| 3:14.73 | Yam |  |  |  |
| 3:14.74 | of |  |  |  |
| 3:14.75 | Shlomo |  |  |  |
| 3:14.76 | against |  |  |  |
| 3:14.77 | those who boast |  |  |  |
| 3:14.78 | of novelties |  |  |  |
| 3:14.79 | which they study |  |  |  |
| 3:14.80 | in Kabbalah |  |  |  |
| 3:14.81 | and he wrote |  |  |  |
| 3:14.82 | that since |  |  |  |
| 3:14.83 | was sealed |  |  |  |
| 3:14.84 | the Talmud |  |  |  |
| 3:14.85 | there is no |  |  |  |
| 3:14.86 | for us |  |  |  |
| 3:14.87 | to be stringent |  |  |  |
| 3:14.88 | against |  |  |  |
| 3:14.89 | the Gemara |  |  |  |
| 3:14.90 | and it is |  |  |  |
| 3:14.91 | like |  |  |  |
| 3:14.92 | heresy |  |  |  |
| 3:14.93 | as explained |  |  |  |
| 3:14.94 | in Chok |  |  |  |
| 3:14.95 | Ya'akov |  |  |  |
| 3:14.96 | section |  |  |  |
| 3:14.97 | 489 |  |  | תפ"ט: the letters tav, pe, tet, which make 489. |
| 3:14.98 | subsection |  |  |  |
| 3:14.99 | 11 |  |  | י"א: the subsection number, 11. |
| 3:14.100 | at length |  |  |  |
| 3:14.101 | and all |  |  |  |
| 3:14.102 | this |  |  |  |
| 3:14.103 | is |  |  |  |
| 3:14.104 | according to |  |  |  |
| 3:14.105 | the assumption |  |  |  |
| 3:14.106 | and the ascription |  |  |  |
| 3:14.107 | which they ascribed |  |  |  |
| 3:14.108 | the Zohar |  |  |  |
| 3:14.109 | to Rabbi Shimon bar Yochai |  |  |  |
| 3:14.110 | the Tanna |  |  |  |
| 3:14.111 | but |  |  |  |
| 3:14.112 | after |  |  |  |
| 3:14.113 | the criticism |  |  |  |
| 3:14.114 | which they examined |  |  |  |
| 3:14.115 | and investigated |  |  |  |
| 3:14.116 | the wise |  |  |  |
| 3:14.117 | of heart |  |  |  |
| 3:14.118 | after him |  |  |  |
| 3:14.119 | from the womb of |  |  |  |
| 3:14.120 | whose |  |  |  |
| 3:14.121 | came out |  |  |  |
| 3:14.122 | the things |  |  |  |
| 3:14.123 | mentioned |  |  |  |
| 3:14.124 | in it |  |  |  |
| 3:14.125 | and when |  |  |  |
| 3:14.126 | came out |  |  |  |
| 3:14.127 | the work |  |  |  |
| 3:14.128 | mentioned |  |  |  |
| 3:14.129 | and from the hand of |  |  |  |
| 3:14.130 | whom |  |  |  |
| 3:14.131 | came out |  |  |  |
| 3:14.132 | and it was shown |  |  |  |
| 3:14.133 | clearly |  |  |  |
| 3:14.134 | well |  |  |  |
| 3:14.135 | that |  |  |  |
| 3:14.136 | it |  |  |  |
| 3:14.137 | a work |  |  |  |
| 3:14.138 | single |  |  |  |
| 3:14.139 | of the latter ones |  |  |  |
| 3:14.140 | and not |  |  |  |
| 3:14.141 | came out |  |  |  |
| 3:14.142 | those |  |  |  |
| 3:14.143 | ideas |  |  |  |
| 3:14.144 | from Rabbi Shimon bar Yochai |  |  |  |
| 3:14.145 | the Tanna |  |  |  |
| 3:14.146 | peace be upon him |  |  | ע"ה: peace be upon him. |
| 3:14.147 | at all |  |  |  |
| 3:14.148 | for |  |  |  |
| 3:14.149 | is found |  |  |  |
| 3:14.150 | in it |  |  |  |
| 3:14.151 | many |  |  |  |
| 3:14.152 | things |  |  |  |
| 3:14.153 | in the fundamentals of |  |  |  |
| 3:14.154 | the faith of |  |  |  |
| 3:14.155 | our Torah |  |  |  |
| 3:14.156 | holy |  |  |  |
| 3:14.157 | contrary to |  |  |  |
| 3:14.158 | the words of |  |  |  |
| 3:14.159 | Rabbi Shimon bar Yochai |  |  |  |
| 3:14.160 | in the Talmud |  |  |  |
| 3:14.161 | and in the Midrashim of |  |  |  |
| 3:14.162 | our Rabbis of blessed memory |  |  |  |
| 3:14.163 | famous |  |  |  |
| 3:14.164 | in Israel |  |  |  |
| 3:14.165 | certainly |  |  |  |
| 3:14.166 | one must not |  |  |  |
| 3:14.167 | rely |  |  |  |
| 3:14.168 | on it |  |  |  |
| 3:14.169 | in any |  |  |  |
| 3:14.170 | matter |  |  |  |
| 3:14.171 | touching |  |  |  |
| 3:14.172 | the customs of |  |  |  |
| 3:14.173 | Israel |  |  |  |
| 3:14.174 | which they practised |  |  |  |
| 3:14.175 | in |  |  |  |
| 3:14.176 | our forebears |  |  |  |
| 3:14.177 | and not |  |  |  |
| 3:14.178 | well |  |  |  |
| 3:14.179 | did |  |  |  |
| 3:14.180 | the sages of |  |  |  |
| 3:14.181 | the days of |  |  |  |
| 3:14.182 | the Middle Ages |  |  | ימי הבינים: the Middle Ages, the author's period term. |
| 3:14.183 | who changed |  |  |  |
| 3:14.184 | upon us |  |  |  |
| 3:14.185 | several |  |  |  |
| 3:14.186 | things |  |  |  |
| 3:14.187 | in the prayer books |  |  |  |
| 3:14.188 | the prayers |  |  |  |
| 3:14.189 | according to |  |  |  |
| 3:14.190 | the word of |  |  |  |
| 3:14.191 | the Zohar |  |  |  |
| 3:14.192 | and the Kabbalists |  |  |  |
| 3:14.193 | for |  |  |  |
| 3:14.194 | changes |  |  |  |
| 3:14.195 | those |  |  |  |
| 3:14.196 | there are |  |  |  |
| 3:14.197 | of them |  |  |  |
| 3:14.198 | that lead |  |  |  |
| 3:14.199 | to |  |  |  |
| 3:14.200 | heresy |  |  |  |
| 3:14.201 | to believe |  |  |  |
| 3:14.202 | in powers |  |  |  |
| 3:14.203 | many |  |  |  |
| 3:14.204 | and to worship |  |  |  |
| 3:14.205 | gods |  |  |  |
| 3:14.206 | other |  |  |  |
| 3:14.207 | as |  |  |  |
| 3:14.208 | will be explained |  |  |  |
| 3:14.209 | below |  |  |  |

Phrases:
- 3:14.1 3:14.2 3:14.3 3:14.4 3:14.5 3:14.6 3:14.7 3:14.8 | §14. See and understand, our friend the hidden reader, how our honored Rabbi | §14 see and understand our friend the reader hidden how as the honored Rabbi
- 3:14.9 3:14.10 | Tam, the aforementioned, | Tam the aforementioned
- 3:14.11 3:14.12 3:14.13 3:14.14 3:14.15 3:14.16 | cried out like a herald against those who engage in the new Kabbalah, | cried out like a herald against those who engage in Kabbalah the new
- 3:14.17 3:14.18 3:14.19 3:14.20 3:14.21 | for they breached the fences of the Torah, | for breached the fences of the Torah
- 3:14.22 3:14.23 3:14.24 3:14.25 | and laid bare, laid bare, its foundations and its roots | and laid bare laid bare its foundations and its roots
- 3:14.26 3:14.27 3:14.28 3:14.29 3:14.30 | in the unity of the Name, may He be blessed, as will be explained below. | in the unity the Name, may He be blessed as will be explained below
- 3:14.31 3:14.32 | Instead of | and instead of
- 3:14.33 3:14.34 3:14.35 3:14.36 3:14.37 3:14.38 | thinking to draw near to the Name, may He be blessed, they strayed from the way | thought to draw near to the Name, may He be blessed strayed from the way
- 3:14.39 3:14.40 3:14.41 3:14.42 3:14.43 | and came to believe in many other gods, | and came and believed in gods many other
- 3:14.44 | and to worship them! | and to worship them
- 3:14.45 3:14.46 3:14.47 3:14.48 3:14.49 3:14.50 3:14.51 3:14.52 3:14.53 3:14.54 | As the words of Rabbi Tam ben Yachya, Mahari Albo also wrote | and as the words of Rabbi Tam ben Yachya wrote also likewise Mahari Albo
- 3:14.55 3:14.56 3:14.57 3:14.58 3:14.59 | in the Ikkarim, chapter 28, of the second. | in the book Ikkarim chapter 28 of the second
- 3:14.60 3:14.61 3:14.62 3:14.63 3:14.64 | And further he expounded at length in the book Ravid ha-Zahav, | and further expounded at length in the book Ravid ha-Zahav
- 3:14.65 3:14.66 3:14.67 3:14.68 3:14.69 3:14.70 | and also brought the words of the great Rabbi Maharshal, | and brought also the words the Rabbi great Maharshal
- 3:14.71 3:14.72 3:14.73 3:14.74 3:14.75 | in his responses and in the work Yam shel Shlomo, | in his responses and in the work Yam of Shlomo
- 3:14.76 3:14.77 3:14.78 3:14.79 3:14.80 | against those who boast of the novelties they study in Kabbalah. | against those who boast of novelties which they study in Kabbalah
- 3:14.81 3:14.82 3:14.83 3:14.84 3:14.85 3:14.86 3:14.87 3:14.88 3:14.89 3:14.90 | And he wrote that since the Talmud was sealed, there is no need for us to be stringent against the Gemara, and it is | and he wrote that since was sealed the Talmud there is no for us to be stringent against the Gemara and it is
- 3:14.91 3:14.92 3:14.93 3:14.94 3:14.95 3:14.96 3:14.97 3:14.98 3:14.99 3:14.100 | like heresy, as explained at length in Chok Ya'akov, section 489, subsection 11. | like heresy as explained in Chok Ya'akov section 489 subsection 11 at length
- 3:14.101 3:14.102 3:14.103 3:14.104 3:14.105 3:14.106 3:14.107 3:14.108 3:14.109 3:14.110 | All this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai, the Tanna, | and all this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai the Tanna
- 3:14.111 3:14.112 3:14.113 3:14.114 3:14.115 3:14.116 3:14.117 3:14.118 3:14.119 3:14.120 | But after the criticism which the wise of heart examined and investigated after him, from whose womb | but after the criticism which they examined and investigated the wise of heart after him from the womb of whose
- 3:14.121 3:14.122 3:14.123 3:14.124 3:14.125 3:14.126 3:14.127 3:14.128 3:14.129 3:14.130 3:14.131 | the things mentioned in it came out, and when the work mentioned came out, and from whose hand it came out. | came out the things mentioned in it and when came out the work mentioned and from the hand of whom came out
- 3:14.132 3:14.133 3:14.134 3:14.135 3:14.136 | And it was shown clearly and well that it | and it was shown clearly well that it
- 3:14.137 3:14.138 3:14.139 | is a single work of the latter ones, | a work single of the latter ones
- 3:14.140 3:14.141 3:14.142 3:14.143 3:14.144 3:14.145 3:14.146 3:14.147 | and those ideas did not come out of Rabbi Shimon bar Yochai, the Tanna, peace be upon him, at all. | and not came out those ideas from Rabbi Shimon bar Yochai the Tanna peace be upon him at all
- 3:14.148 3:14.149 3:14.150 3:14.151 3:14.152 3:14.153 3:14.154 3:14.155 3:14.156 3:14.157 3:14.158 3:14.159 3:14.160 3:14.161 3:14.162 3:14.163 3:14.164 | for in it are found many things in the fundamentals of the faith of our holy Torah contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory, famous in Israel, | for is found in it many things in the fundamentals of the faith of our Torah holy contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory famous in Israel
- 3:14.165 3:14.166 3:14.167 3:14.168 3:14.169 3:14.170 3:14.171 3:14.172 3:14.173 3:14.174 3:14.175 3:14.176 | certainly one must not rely on it in any matter touching the customs of Israel that our forebears practised. | certainly one must not rely on it in any matter touching the customs of Israel which they practised in our forebears
- 3:14.177 3:14.178 3:14.179 3:14.180 3:14.181 3:14.182 3:14.183 3:14.184 3:14.185 3:14.186 | And the sages of the Middle Ages did not do well, who changed upon us several things | and not well did the sages of the days of the Middle Ages who changed upon us several things
- 3:14.187 3:14.188 3:14.189 3:14.190 3:14.191 | in the prayer books, according to the Zohar | in the prayer books the prayers according to the word of the Zohar
- 3:14.192 3:14.193 3:14.194 3:14.195 3:14.196 3:14.197 3:14.198 3:14.199 3:14.200 | and the Kabbalists, for some of those changes lead to heresy, | and the Kabbalists for changes those there are of them that lead to heresy
- 3:14.201 3:14.202 3:14.203 3:14.204 3:14.205 3:14.206 3:14.207 3:14.208 3:14.209 | to believe in many powers and to worship other gods, as will be explained below. | to believe in powers many and to worship gods other as will be explained below

Grammar notes:
- ראה והבן ... איך ... צוח: 'see and understand how ... cried out'; the two imperatives open one long period, which runs to the end of the paragraph.
- ותחת אשר חשבו להתקרב ... תעו: 'instead of thinking to draw near, they strayed'; the author sets the intention against the act.
- כ' גם כן מהר"י אלבו: 'wrote also Mahari Albo'; the verb comes first in the Hebrew and the English places it after the name.
- וכתב דמאחר ... והוי כמו מינות: an indirect report with no quotation marks.
- וכל זה הוא לפי ההנחה והיחוס: 'all this is according to the assumption and the ascription'; the ascription to Rabbi Shimon bar Yochai is what is disputed.

Translator's note: The paragraph is a single period with a report of two authorities at its end. The Wikisource base text is used; the 1931 scan is not collated.

Commentary rulings:
- apparatus: Allusion, not quotation: 'אלהים אחרים' echoes Exod 20:3 ('other gods'); the English follows the author's wording.
- apparatus: Base text is Wikisource; a general line until the 1931 scan is collated (open question 2).

## Milchamot ha-Shem 3:15

Hebrew (Wikisource): טו) גם הביא עוד הר' רביד הזהב דברי מרן בב"י דחייש לדברי הזהר במקום שלא נזכר בגמרא היפך מדבריו, וכתב עליו הרמ"א ז"ל שאע"פ שלא נזכר בגמרא היפך דבריו, אין לזה מדברי הפוסקים, ואין לחוש לדברי הזהר ע"ש. עוד הביא דברי הרדב"ז ז"ל שכתב שכל דבר שנזכר בגמרא או באחד מן הפוסקים, או בעלי ההלכות אפי' שיהיה היפך ממה שכתוב בספרי הקבלה אני מורה בו בין להקל בין להחמיר, ולא אחוש למה שכתוב באותן הספרים. ולעצמי אם הוא חומרא אני נוהג אותו ואם קולא לא אחוש לה. ועוד כתב בשם הרא"ם ושאר פוסקים שאין לחוש לדברי המקובלים נגד שום פוסק ע"ש. אמור מעתה שמדברי הפוסקים הנזכרים בתשובת הרב רבי הזהב והם מרן ז"ל ומור"ם והרדב"ז וכנה"ג ומהר"י הלוי כ' מהר"ש הלוי ור' אליהו הלוי, והר"מ אלקבולי ור"ת בן יחיא והר' חק יעקב ומהרש"ל ומהר"ם אלשקר שאומרים שאין לחוש לדברי הזהר והמקובלים שהוא נגד המשנה והתלמוד ואפילו נגד פוסק יחידי. הנה מזה הוכיח לעצמך שכתבת שהם הלכה למשה מסיני, שדבריך בדותא וסכלות גמורה.

| id | word | punctuation after |
|---|---|---|
| 3:15.1 | טו | ) |
| 3:15.2 | גם |  |
| 3:15.3 | הביא |  |
| 3:15.4 | עוד |  |
| 3:15.5 | הר | ' |
| 3:15.6 | רביד |  |
| 3:15.7 | הזהב |  |
| 3:15.8 | דברי |  |
| 3:15.9 | מרן |  |
| 3:15.10 | בב"י |  |
| 3:15.11 | דחייש |  |
| 3:15.12 | לדברי |  |
| 3:15.13 | הזהר |  |
| 3:15.14 | במקום |  |
| 3:15.15 | שלא |  |
| 3:15.16 | נזכר |  |
| 3:15.17 | בגמרא |  |
| 3:15.18 | היפך |  |
| 3:15.19 | מדבריו | , |
| 3:15.20 | וכתב |  |
| 3:15.21 | עליו |  |
| 3:15.22 | הרמ"א |  |
| 3:15.23 | ז"ל |  |
| 3:15.24 | שאע"פ |  |
| 3:15.25 | שלא |  |
| 3:15.26 | נזכר |  |
| 3:15.27 | בגמרא |  |
| 3:15.28 | היפך |  |
| 3:15.29 | דבריו | , |
| 3:15.30 | אין |  |
| 3:15.31 | לזה |  |
| 3:15.32 | מדברי |  |
| 3:15.33 | הפוסקים | , |
| 3:15.34 | ואין |  |
| 3:15.35 | לחוש |  |
| 3:15.36 | לדברי |  |
| 3:15.37 | הזהר |  |
| 3:15.38 | ע"ש | . |
| 3:15.39 | עוד |  |
| 3:15.40 | הביא |  |
| 3:15.41 | דברי |  |
| 3:15.42 | הרדב"ז |  |
| 3:15.43 | ז"ל |  |
| 3:15.44 | שכתב |  |
| 3:15.45 | שכל |  |
| 3:15.46 | דבר |  |
| 3:15.47 | שנזכר |  |
| 3:15.48 | בגמרא |  |
| 3:15.49 | או |  |
| 3:15.50 | באחד |  |
| 3:15.51 | מן |  |
| 3:15.52 | הפוסקים | , |
| 3:15.53 | או |  |
| 3:15.54 | בעלי |  |
| 3:15.55 | ההלכות |  |
| 3:15.56 | אפי | ' |
| 3:15.57 | שיהיה |  |
| 3:15.58 | היפך |  |
| 3:15.59 | ממה |  |
| 3:15.60 | שכתוב |  |
| 3:15.61 | בספרי |  |
| 3:15.62 | הקבלה |  |
| 3:15.63 | אני |  |
| 3:15.64 | מורה |  |
| 3:15.65 | בו |  |
| 3:15.66 | בין |  |
| 3:15.67 | להקל |  |
| 3:15.68 | בין |  |
| 3:15.69 | להחמיר | , |
| 3:15.70 | ולא |  |
| 3:15.71 | אחוש |  |
| 3:15.72 | למה |  |
| 3:15.73 | שכתוב |  |
| 3:15.74 | באותן |  |
| 3:15.75 | הספרים | . |
| 3:15.76 | ולעצמי |  |
| 3:15.77 | אם |  |
| 3:15.78 | הוא |  |
| 3:15.79 | חומרא |  |
| 3:15.80 | אני |  |
| 3:15.81 | נוהג |  |
| 3:15.82 | אותו |  |
| 3:15.83 | ואם |  |
| 3:15.84 | קולא |  |
| 3:15.85 | לא |  |
| 3:15.86 | אחוש |  |
| 3:15.87 | לה | . |
| 3:15.88 | ועוד |  |
| 3:15.89 | כתב |  |
| 3:15.90 | בשם |  |
| 3:15.91 | הרא"ם |  |
| 3:15.92 | ושאר |  |
| 3:15.93 | פוסקים |  |
| 3:15.94 | שאין |  |
| 3:15.95 | לחוש |  |
| 3:15.96 | לדברי |  |
| 3:15.97 | המקובלים |  |
| 3:15.98 | נגד |  |
| 3:15.99 | שום |  |
| 3:15.100 | פוסק |  |
| 3:15.101 | ע"ש | . |
| 3:15.102 | אמור |  |
| 3:15.103 | מעתה |  |
| 3:15.104 | שמדברי |  |
| 3:15.105 | הפוסקים |  |
| 3:15.106 | הנזכרים |  |
| 3:15.107 | בתשובת |  |
| 3:15.108 | הרב |  |
| 3:15.109 | רבי |  |
| 3:15.110 | הזהב |  |
| 3:15.111 | והם |  |
| 3:15.112 | מרן |  |
| 3:15.113 | ז"ל |  |
| 3:15.114 | ומור"ם |  |
| 3:15.115 | והרדב"ז |  |
| 3:15.116 | וכנה"ג |  |
| 3:15.117 | ומהר"י |  |
| 3:15.118 | הלוי |  |
| 3:15.119 | כ | ' |
| 3:15.120 | מהר"ש |  |
| 3:15.121 | הלוי |  |
| 3:15.122 | ור | ' |
| 3:15.123 | אליהו |  |
| 3:15.124 | הלוי | , |
| 3:15.125 | והר"מ |  |
| 3:15.126 | אלקבולי |  |
| 3:15.127 | ור"ת |  |
| 3:15.128 | בן |  |
| 3:15.129 | יחיא |  |
| 3:15.130 | והר | ' |
| 3:15.131 | חק |  |
| 3:15.132 | יעקב |  |
| 3:15.133 | ומהרש"ל |  |
| 3:15.134 | ומהר"ם |  |
| 3:15.135 | אלשקר |  |
| 3:15.136 | שאומרים |  |
| 3:15.137 | שאין |  |
| 3:15.138 | לחוש |  |
| 3:15.139 | לדברי |  |
| 3:15.140 | הזהר |  |
| 3:15.141 | והמקובלים |  |
| 3:15.142 | שהוא |  |
| 3:15.143 | נגד |  |
| 3:15.144 | המשנה |  |
| 3:15.145 | והתלמוד |  |
| 3:15.146 | ואפילו |  |
| 3:15.147 | נגד |  |
| 3:15.148 | פוסק |  |
| 3:15.149 | יחידי | . |
| 3:15.150 | הנה |  |
| 3:15.151 | מזה |  |
| 3:15.152 | הוכיח |  |
| 3:15.153 | לעצמך |  |
| 3:15.154 | שכתבת |  |
| 3:15.155 | שהם |  |
| 3:15.156 | הלכה |  |
| 3:15.157 | למשה |  |
| 3:15.158 | מסיני | , |
| 3:15.159 | שדבריך |  |
| 3:15.160 | בדותא |  |
| 3:15.161 | וסכלות |  |
| 3:15.162 | גמורה | . |


### Translation under review (3:15)

English: §15. Also he brought further the Rabbi, the author of Ravid ha-Zahav, the words of Maran in Beit Yosef, who was concerned about the words of the Zohar in a place where the Gemara does not mention anything contrary to them. And he wrote about him, the Rema, of blessed memory, that although the Gemara does not mention anything contrary to its words, there is no basis for this in the words of the decisors, and one need not be concerned with the words of the Zohar, see there. Further he brought the words of the Radbaz, of blessed memory, that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot, even if it be contrary to what is written in the books of Kabbalah, I rule on, both to be lenient and to be strict, and I will not be concerned with what is written in those books. As for myself, if it is a stringency, I follow it and if a leniency, I will not be concerned with it. Further he wrote in the name of the Ra'em and other decisors that one need not be concerned about the words of the Kabbalists against any decisor, see there. Say now, that from the decisors mentioned in the responsum of the Rav, Rabbi ha-Zahav, and they are Maran, of blessed memory, and Moreinu ha-Rav Moshe, and the Radbaz, and the Kenesset ha-Gedolah, and Mahari ha-Levi wrote, Maharash ha-Levi and Rabbi Eliyahu ha-Levi, and Rabbi Moshe Alkabetz, and Rabbi Tam ben Yachya, and the Rabbi, the author of Chok Ya'akov, and Maharshal, and Maharam al-Shakar, who say that one need not be concerned with the words of the Zohar and the Kabbalists, when they are against the Mishnah and the Talmud, and even against a single decisor. Behold, from this prove to yourself that you wrote that they are a law of Moses from Sinai, that your words are a fabrication and utter folly.

Literal: §15 also he brought further the Rabbi, the author of Ravid ha-Zahav the words of Maran in Beit Yosef who was concerned about the words of the Zohar in a place where not mentioned in the Gemara contrary to them and he wrote about him the Rema of blessed memory that although that not mentioned in the Gemara contrary to its words there is no for this in the words of the decisors and one need not be concerned about the words of the Zohar see there further he brought the words of the Radbaz of blessed memory that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot even if it be contrary to what is written in the books of Kabbalah I rule on it both to be lenient and to be strict and not I will be concerned with what is written in those books as for myself if it is a stringency I follow it and if a leniency not I will be concerned with it further he wrote in the name of the Ra'em and other decisors that one need not be concerned about words the Kabbalists against any decisor see there say now that from the words of the decisors mentioned in the responsum of the Rav Rabbi ha-Zahav and they are Maran of blessed memory and Moreinu ha-Rav Moshe and the Radbaz and the Kenesset ha-Gedolah and Mahari ha-Levi wrote Maharash ha-Levi and Rabbi Eliyahu ha-Levi and Rabbi Moshe Alkabetz and Rabbi Tam ben Yachya and the Rabbi, the author of Chok Ya'akov and Maharshal and Maharam al-Shakar who say that one need not be concerned about words the Zohar and the Kabbalists when they are against the Mishnah and the Talmud and even against a decisor single behold from this prove to yourself that you wrote that they are a law of Moses from Sinai that your words are a fabrication and folly utter

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:15.1 | §15 |  |  |  |
| 3:15.2 | also |  |  |  |
| 3:15.3 | he brought |  |  |  |
| 3:15.4 | further |  |  |  |
| 3:15.5 | the Rabbi, the author of |  |  | הר' + a book title: 'the Rabbi, the author of Ravid ha-Zahav'; the title is never a personal name. |
| 3:15.6 | Ravid |  |  |  |
| 3:15.7 | ha-Zahav |  |  |  |
| 3:15.8 | the words of |  |  |  |
| 3:15.9 | Maran |  |  | מרן: the author's honorific for the writer of the Beit Yosef; the text names no one. Literally 'our master'. |
| 3:15.10 | in Beit Yosef |  |  | בב"י: Beit Yosef, the commentary on the Tur. |
| 3:15.11 | who was concerned |  |  |  |
| 3:15.12 | about the words of |  |  |  |
| 3:15.13 | the Zohar |  |  |  |
| 3:15.14 | in a place |  |  |  |
| 3:15.15 | where not |  |  |  |
| 3:15.16 | mentioned |  |  |  |
| 3:15.17 | in the Gemara |  |  |  |
| 3:15.18 | contrary |  |  |  |
| 3:15.19 | to them |  |  |  |
| 3:15.20 | and he wrote |  |  |  |
| 3:15.21 | about him |  |  |  |
| 3:15.22 | the Rema |  |  | הרמ"א: the Rema, Rabbi Moshe Isserles. |
| 3:15.23 | of blessed memory |  |  |  |
| 3:15.24 | that although |  |  |  |
| 3:15.25 | that not |  |  |  |
| 3:15.26 | mentioned |  |  |  |
| 3:15.27 | in the Gemara |  |  |  |
| 3:15.28 | contrary |  |  |  |
| 3:15.29 | to its words |  |  |  |
| 3:15.30 | there is no |  |  |  |
| 3:15.31 | for this |  |  |  |
| 3:15.32 | in the words of |  |  |  |
| 3:15.33 | the decisors |  |  |  |
| 3:15.34 | and one need not |  |  |  |
| 3:15.35 | be concerned |  |  |  |
| 3:15.36 | about the words of |  |  |  |
| 3:15.37 | the Zohar |  |  |  |
| 3:15.38 | see there |  |  |  |
| 3:15.39 | further |  |  |  |
| 3:15.40 | he brought |  |  |  |
| 3:15.41 | the words of |  |  |  |
| 3:15.42 | the Radbaz |  |  | הרדב"ז: the Radbaz, Rabbi David ibn Zimra. |
| 3:15.43 | of blessed memory |  |  |  |
| 3:15.44 | that he wrote |  |  |  |
| 3:15.45 | that every |  |  |  |
| 3:15.46 | matter |  |  |  |
| 3:15.47 | mentioned |  |  |  |
| 3:15.48 | in the Gemara |  |  |  |
| 3:15.49 | or |  |  |  |
| 3:15.50 | in one |  |  |  |
| 3:15.51 | of |  |  |  |
| 3:15.52 | the decisors |  |  |  |
| 3:15.53 | or |  |  |  |
| 3:15.54 | the masters of |  |  |  |
| 3:15.55 | the halakhot |  |  |  |
| 3:15.56 | even |  |  |  |
| 3:15.57 | if it be |  |  |  |
| 3:15.58 | contrary |  |  |  |
| 3:15.59 | to what |  |  |  |
| 3:15.60 | is written |  |  |  |
| 3:15.61 | in the books |  |  |  |
| 3:15.62 | of Kabbalah |  |  |  |
| 3:15.63 | I |  |  |  |
| 3:15.64 | rule |  |  |  |
| 3:15.65 | on it |  |  |  |
| 3:15.66 | both |  |  |  |
| 3:15.67 | to be lenient |  |  |  |
| 3:15.68 | and |  |  |  |
| 3:15.69 | to be strict |  |  |  |
| 3:15.70 | and not |  |  |  |
| 3:15.71 | I will be concerned |  |  |  |
| 3:15.72 | with what |  |  |  |
| 3:15.73 | is written |  |  |  |
| 3:15.74 | in those |  |  |  |
| 3:15.75 | books |  |  |  |
| 3:15.76 | as for myself |  |  |  |
| 3:15.77 | if |  |  |  |
| 3:15.78 | it is |  |  |  |
| 3:15.79 | a stringency |  |  |  |
| 3:15.80 | I |  |  |  |
| 3:15.81 | follow |  |  |  |
| 3:15.82 | it |  |  |  |
| 3:15.83 | and if |  |  |  |
| 3:15.84 | a leniency |  |  |  |
| 3:15.85 | not |  |  |  |
| 3:15.86 | I will be concerned |  |  |  |
| 3:15.87 | with it |  |  |  |
| 3:15.88 | further |  |  |  |
| 3:15.89 | he wrote |  |  |  |
| 3:15.90 | in the name of |  |  |  |
| 3:15.91 | the Ra'em |  |  | הרא"ם: the Ra'em. |
| 3:15.92 | and other |  |  |  |
| 3:15.93 | decisors |  |  |  |
| 3:15.94 | that one need not |  |  |  |
| 3:15.95 | be concerned |  |  |  |
| 3:15.96 | about words |  |  |  |
| 3:15.97 | the Kabbalists |  |  |  |
| 3:15.98 | against |  |  |  |
| 3:15.99 | any |  |  |  |
| 3:15.100 | decisor |  |  |  |
| 3:15.101 | see there |  |  |  |
| 3:15.102 | say |  |  |  |
| 3:15.103 | now |  |  |  |
| 3:15.104 | that from the words of |  |  |  |
| 3:15.105 | the decisors |  |  |  |
| 3:15.106 | mentioned |  |  |  |
| 3:15.107 | in the responsum of |  |  |  |
| 3:15.108 | the Rav |  |  |  |
| 3:15.109 | Rabbi |  |  |  |
| 3:15.110 | ha-Zahav |  |  | הזהב: the Hebrew prints 'הרב רבי הזהב' here; see apparatus. |
| 3:15.111 | and they are |  |  |  |
| 3:15.112 | Maran |  |  |  |
| 3:15.113 | of blessed memory |  |  |  |
| 3:15.114 | and Moreinu ha-Rav Moshe |  |  |  |
| 3:15.115 | and the Radbaz |  |  |  |
| 3:15.116 | and the Kenesset ha-Gedolah |  |  | כנה"ג: the Kenesset ha-Gedolah. |
| 3:15.117 | and Mahari |  |  |  |
| 3:15.118 | ha-Levi |  |  |  |
| 3:15.119 | wrote |  |  |  |
| 3:15.120 | Maharash |  |  |  |
| 3:15.121 | ha-Levi |  |  |  |
| 3:15.122 | and Rabbi |  |  |  |
| 3:15.123 | Eliyahu |  |  |  |
| 3:15.124 | ha-Levi |  |  |  |
| 3:15.125 | and Rabbi Moshe |  |  | הר"מ: read as Rabbi Moshe; the abbreviation is not expanded in the text. |
| 3:15.126 | Alkabetz |  |  |  |
| 3:15.127 | and Rabbi Tam |  |  |  |
| 3:15.128 | ben |  |  |  |
| 3:15.129 | Yachya |  |  |  |
| 3:15.130 | and the Rabbi, the author of |  |  |  |
| 3:15.131 | Chok |  |  |  |
| 3:15.132 | Ya'akov |  |  |  |
| 3:15.133 | and Maharshal |  |  |  |
| 3:15.134 | and Maharam |  |  |  |
| 3:15.135 | al-Shakar |  |  | מהר"ם אלשקר: Maharam al-Shakar, as the author writes him. |
| 3:15.136 | who say |  |  |  |
| 3:15.137 | that one need not |  |  |  |
| 3:15.138 | be concerned |  |  |  |
| 3:15.139 | about words |  |  |  |
| 3:15.140 | the Zohar |  |  |  |
| 3:15.141 | and the Kabbalists |  |  |  |
| 3:15.142 | when they are |  |  |  |
| 3:15.143 | against |  |  |  |
| 3:15.144 | the Mishnah |  |  |  |
| 3:15.145 | and the Talmud |  |  |  |
| 3:15.146 | and even |  |  |  |
| 3:15.147 | against |  |  |  |
| 3:15.148 | a decisor |  |  |  |
| 3:15.149 | single |  |  |  |
| 3:15.150 | behold |  |  |  |
| 3:15.151 | from this |  |  |  |
| 3:15.152 | prove |  |  |  |
| 3:15.153 | to yourself |  |  |  |
| 3:15.154 | that you wrote |  |  |  |
| 3:15.155 | that they are |  |  |  |
| 3:15.156 | a law |  |  |  |
| 3:15.157 | of Moses |  |  |  |
| 3:15.158 | from Sinai |  |  |  |
| 3:15.159 | that your words are |  |  |  |
| 3:15.160 | a fabrication |  |  |  |
| 3:15.161 | and folly |  |  |  |
| 3:15.162 | utter |  |  |  |

Phrases:
- 3:15.1 3:15.2 3:15.3 3:15.4 | §15. Also he brought further | §15 also he brought further
- 3:15.5 3:15.6 3:15.7 | the Rabbi, the author of Ravid ha-Zahav, | the Rabbi, the author of Ravid ha-Zahav
- 3:15.8 3:15.9 3:15.10 3:15.11 | the words of Maran in Beit Yosef, who was concerned | the words of Maran in Beit Yosef who was concerned
- 3:15.12 3:15.13 3:15.14 3:15.15 3:15.16 3:15.17 3:15.18 3:15.19 | about the words of the Zohar in a place where the Gemara does not mention anything contrary to them. | about the words of the Zohar in a place where not mentioned in the Gemara contrary to them
- 3:15.20 3:15.21 3:15.22 3:15.23 | And he wrote about him, the Rema, of blessed memory, | and he wrote about him the Rema of blessed memory
- 3:15.24 3:15.25 3:15.26 3:15.27 3:15.28 3:15.29 | that although the Gemara does not mention anything contrary to its words, | that although that not mentioned in the Gemara contrary to its words
- 3:15.30 3:15.31 3:15.32 3:15.33 3:15.34 3:15.35 3:15.36 3:15.37 3:15.38 | there is no basis for this in the words of the decisors, and one need not be concerned with the words of the Zohar, see there. | there is no for this in the words of the decisors and one need not be concerned about the words of the Zohar see there
- 3:15.39 3:15.40 3:15.41 3:15.42 3:15.43 | Further he brought the words of the Radbaz, of blessed memory, | further he brought the words of the Radbaz of blessed memory
- 3:15.44 3:15.45 3:15.46 3:15.47 3:15.48 3:15.49 3:15.50 3:15.51 3:15.52 3:15.53 3:15.54 3:15.55 3:15.56 3:15.57 3:15.58 3:15.59 3:15.60 3:15.61 3:15.62 3:15.63 3:15.64 3:15.65 3:15.66 3:15.67 3:15.68 3:15.69 3:15.70 3:15.71 3:15.72 3:15.73 3:15.74 3:15.75 | that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot, even if it be contrary to what is written in the books of Kabbalah, I rule on, both to be lenient and to be strict, and I will not be concerned with what is written in those books. | that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot even if it be contrary to what is written in the books of Kabbalah I rule on it both to be lenient and to be strict and not I will be concerned with what is written in those books
- 3:15.76 3:15.77 3:15.78 3:15.79 3:15.80 3:15.81 3:15.82 | As for myself, if it is a stringency, I follow it | as for myself if it is a stringency I follow it
- 3:15.83 3:15.84 3:15.85 3:15.86 3:15.87 | and if a leniency, I will not be concerned with it. | and if a leniency not I will be concerned with it
- 3:15.88 3:15.89 3:15.90 3:15.91 3:15.92 3:15.93 3:15.94 3:15.95 3:15.96 3:15.97 3:15.98 3:15.99 3:15.100 3:15.101 | Further he wrote in the name of the Ra'em and other decisors that one need not be concerned about the words of the Kabbalists against any decisor, see there. | further he wrote in the name of the Ra'em and other decisors that one need not be concerned about words the Kabbalists against any decisor see there
- 3:15.102 3:15.103 3:15.104 3:15.105 3:15.106 3:15.107 3:15.108 3:15.109 3:15.110 | Say now, that from the decisors mentioned in the responsum of the Rav, Rabbi ha-Zahav, | say now that from the words of the decisors mentioned in the responsum of the Rav Rabbi ha-Zahav
- 3:15.111 3:15.112 3:15.113 3:15.114 3:15.115 3:15.116 3:15.117 3:15.118 3:15.119 3:15.120 3:15.121 3:15.122 3:15.123 3:15.124 | and they are Maran, of blessed memory, and Moreinu ha-Rav Moshe, and the Radbaz, and the Kenesset ha-Gedolah, and Mahari ha-Levi wrote, Maharash ha-Levi and Rabbi Eliyahu ha-Levi, | and they are Maran of blessed memory and Moreinu ha-Rav Moshe and the Radbaz and the Kenesset ha-Gedolah and Mahari ha-Levi wrote Maharash ha-Levi and Rabbi Eliyahu ha-Levi
- 3:15.125 3:15.126 3:15.127 3:15.128 3:15.129 | and Rabbi Moshe Alkabetz, and Rabbi Tam ben Yachya, | and Rabbi Moshe Alkabetz and Rabbi Tam ben Yachya
- 3:15.130 3:15.131 3:15.132 3:15.133 3:15.134 3:15.135 | and the Rabbi, the author of Chok Ya'akov, and Maharshal, and Maharam al-Shakar, | and the Rabbi, the author of Chok Ya'akov and Maharshal and Maharam al-Shakar
- 3:15.136 3:15.137 3:15.138 3:15.139 3:15.140 3:15.141 3:15.142 3:15.143 3:15.144 3:15.145 3:15.146 3:15.147 3:15.148 3:15.149 | who say that one need not be concerned with the words of the Zohar and the Kabbalists, when they are against the Mishnah and the Talmud, and even against a single decisor. | who say that one need not be concerned about words the Zohar and the Kabbalists when they are against the Mishnah and the Talmud and even against a decisor single
- 3:15.150 3:15.151 3:15.152 3:15.153 3:15.154 3:15.155 3:15.156 3:15.157 3:15.158 3:15.159 3:15.160 3:15.161 3:15.162 | Behold, from this prove to yourself that you wrote that they are a law of Moses from Sinai, that your words are a fabrication and utter folly. | behold from this prove to yourself that you wrote that they are a law of Moses from Sinai that your words are a fabrication and folly utter

Grammar notes:
- גם הביא עוד ... וכתב עליו ... שכתב שכל דבר: the reporting verbs carry indirect speech throughout the paragraph, so the English has no quotation marks.
- שכתב שכל דבר ... אני מורה בו ... ולעצמי: the Radbaz's words are reported in indirect speech, and the first person ('I rule', 'I follow') is kept.
- הנה מזה הוכיח לעצמך: the closing turn is addressed to a second person, whom the author charges with calling these rulings 'a law of Moses from Sinai'.

Translator's note: The paragraph reports a sequence of decisors and ends with the author's address to his opponent. The base text is Wikisource.

Commentary rulings:
- apparatus: The printed 'הרב רבי הזהב' (3:15.108-110) is read as the author of Ravid ha-Zahav named at 3:15.5-7; the Hebrew is kept, and the reading is not certain.
- apparatus: Base text is Wikisource; a general line until the 1931 scan is collated (open question 2).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

