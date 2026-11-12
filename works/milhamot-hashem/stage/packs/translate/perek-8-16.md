# Translation pack: milhamot-hashem perek-8-16 (Milchamot ha-Shem 8:19–8:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-16`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **מין / מינות** → "heretic / heresy".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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
- **תא חזי** → "come and see".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
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


## Milchamot ha-Shem 8:19

Hebrew (Wikisource): ונחזור לעניינינו ובהדי דתפסוח אשתמיט ביניתא מן מזייה, היינו שתפסו לאותם האנשים הטועים ומטעים אחרים לדרוש בתורה דרשות של דופי כאלה לעבוד אלוה אחר קצר אפים מבלעדי ה' אלהינו, ולירא מאשתו (זהר וירא דף קי"ב). ואחזו בזקנם שיחזרו בהם מלפרש ולדרוש דרשות זרות כאלה בתורה הקדושה, דרשות ופירושים המביאים לידי מינות ואמונת אלהים רבים. ונשמט שער מזקנם באחזם בה, כי הדרת פנים זקן ונפחת כבודם על ידי זה, ונגרע הודם וחדרם. ולכן רמו קלא ואזל קליהו ארבע מאה פרסי, שצעקו חמס על הודם והדרם וכבודם כי אזל ועבר והלך קול זלזולם ארבע מאה פרסי, באורך ורוחב ארץ ישראל שהיא ארבע מאה פרסי, ואמרי היכי נעביד ליה, דלמא מרחמי עליה מן שמיא, שחששו שמא בזכות התורה שעסקו בה אע"פ שמרמזים בה דברי מינות ירחמו עליהם מן השמים, כמו שמצינו באלישע אחר שעמדה לו זכות תורה בקצת, ואמרו בשמים לא מידן נידייניה דעסק בתורה וכו'. אמר להו נביא שדויה בדורא דאברא וכו' שהבטיחם הנביא זכריה ואמר להם אל תחושו לקול צעקתם וצוה אותם לשקע ולשכח מלב בני ישראל כל אותן הדרשות והרמזים של מינות כדודא דאברא ששוקע בים, כמ"ד צללו כעופרת במים אדירים של תורת אמת המקובלת ממשה רבינו ע"ה להקים בה בתי מדרשות בכל עיר ובכל פלך, ולהעמיד תלמידים הרבה כדאמרינן בפ' קמא דאבות, כי אנשי כנה"ג אמרו שלשה דברים היו מתונים בדין והעמידו תלמידים הרבה, ועשו סייג לתורה שלא לדרוש בה כל איש ככל העולה על רוחו, אלא על פי הכללים והיסודות האמתיים שנמסרו למשה רבינו ע"ה בסיני! וזוהי מליצת חפיוה לפומא באברא שמנעו וסתמו את פי החכמים בעיניהם מלדרוש בתורה הזיות והבלים המביאים לידי מינות, והשבעות רוחות של מתים לבוא להגיד להם נסתרות באברא, אשר אם תקים עליו לא ישמע קול, כן נעשו דורשי דרשות של דופי והזיות כחרשים! ובזה נתבטל יצרא דע"ז, ומן השמים אסתייעא מלתא, ועלתה בידם לשכח שמות האלילים אשר ידמו התרים אחרי לבבם לחשוב רבוי באלוהות, ושיש כמה עלות וסבות המשתלשלים זה מזה, ונוטלים רשות זה מזה לברוא את אשר יחפוץ לברוא אותו כאימא שנטלה רשות מאבא לברוא אדם הראשון, וכל אחד מאלו הצורות נקרא בשם ידוע, ונקרא בשם הוי'ה ואדנות ושאר השמות המיוחדים לשי"ת, ושהעבודה והתפלה להם לבדם אעפ"י שהם נבראים, אבל האלוה העליון איננו ראוי לשום שבח ותהלה לרוב מעלתו והעלמו, ואין לו דמות וצורה ואין לו אשה ולא שם מוגבל! כל מחשבות מינות הללו השכיחום רבותינו אנשי כנסת הגדולה ז"ל.

| id | word | punctuation after |
|---|---|---|
| 8:19.1 | ונחזור |  |
| 8:19.2 | לעניינינו |  |
| 8:19.3 | ובהדי |  |
| 8:19.4 | דתפסוח |  |
| 8:19.5 | אשתמיט |  |
| 8:19.6 | ביניתא |  |
| 8:19.7 | מן |  |
| 8:19.8 | מזייה | , |
| 8:19.9 | היינו |  |
| 8:19.10 | שתפסו |  |
| 8:19.11 | לאותם |  |
| 8:19.12 | האנשים |  |
| 8:19.13 | הטועים |  |
| 8:19.14 | ומטעים |  |
| 8:19.15 | אחרים |  |
| 8:19.16 | לדרוש |  |
| 8:19.17 | בתורה |  |
| 8:19.18 | דרשות |  |
| 8:19.19 | של |  |
| 8:19.20 | דופי |  |
| 8:19.21 | כאלה |  |
| 8:19.22 | לעבוד |  |
| 8:19.23 | אלוה |  |
| 8:19.24 | אחר |  |
| 8:19.25 | קצר |  |
| 8:19.26 | אפים |  |
| 8:19.27 | מבלעדי |  |
| 8:19.28 | ה | ' |
| 8:19.29 | אלהינו | , |
| 8:19.30 | ולירא |  |
| 8:19.31 | מאשתו |  |
| 8:19.32 | זהר |  |
| 8:19.33 | וירא |  |
| 8:19.34 | דף |  |
| 8:19.35 | קי"ב | ). |
| 8:19.36 | ואחזו |  |
| 8:19.37 | בזקנם |  |
| 8:19.38 | שיחזרו |  |
| 8:19.39 | בהם |  |
| 8:19.40 | מלפרש |  |
| 8:19.41 | ולדרוש |  |
| 8:19.42 | דרשות |  |
| 8:19.43 | זרות |  |
| 8:19.44 | כאלה |  |
| 8:19.45 | בתורה |  |
| 8:19.46 | הקדושה | , |
| 8:19.47 | דרשות |  |
| 8:19.48 | ופירושים |  |
| 8:19.49 | המביאים |  |
| 8:19.50 | לידי |  |
| 8:19.51 | מינות |  |
| 8:19.52 | ואמונת |  |
| 8:19.53 | אלהים |  |
| 8:19.54 | רבים | . |
| 8:19.55 | ונשמט |  |
| 8:19.56 | שער |  |
| 8:19.57 | מזקנם |  |
| 8:19.58 | באחזם |  |
| 8:19.59 | בה | , |
| 8:19.60 | כי |  |
| 8:19.61 | הדרת |  |
| 8:19.62 | פנים |  |
| 8:19.63 | זקן |  |
| 8:19.64 | ונפחת |  |
| 8:19.65 | כבודם |  |
| 8:19.66 | על |  |
| 8:19.67 | ידי |  |
| 8:19.68 | זה | , |
| 8:19.69 | ונגרע |  |
| 8:19.70 | הודם |  |
| 8:19.71 | וחדרם | . |
| 8:19.72 | ולכן |  |
| 8:19.73 | רמו |  |
| 8:19.74 | קלא |  |
| 8:19.75 | ואזל |  |
| 8:19.76 | קליהו |  |
| 8:19.77 | ארבע |  |
| 8:19.78 | מאה |  |
| 8:19.79 | פרסי | , |
| 8:19.80 | שצעקו |  |
| 8:19.81 | חמס |  |
| 8:19.82 | על |  |
| 8:19.83 | הודם |  |
| 8:19.84 | והדרם |  |
| 8:19.85 | וכבודם |  |
| 8:19.86 | כי |  |
| 8:19.87 | אזל |  |
| 8:19.88 | ועבר |  |
| 8:19.89 | והלך |  |
| 8:19.90 | קול |  |
| 8:19.91 | זלזולם |  |
| 8:19.92 | ארבע |  |
| 8:19.93 | מאה |  |
| 8:19.94 | פרסי | , |
| 8:19.95 | באורך |  |
| 8:19.96 | ורוחב |  |
| 8:19.97 | ארץ |  |
| 8:19.98 | ישראל |  |
| 8:19.99 | שהיא |  |
| 8:19.100 | ארבע |  |
| 8:19.101 | מאה |  |
| 8:19.102 | פרסי | , |
| 8:19.103 | ואמרי |  |
| 8:19.104 | היכי |  |
| 8:19.105 | נעביד |  |
| 8:19.106 | ליה | , |
| 8:19.107 | דלמא |  |
| 8:19.108 | מרחמי |  |
| 8:19.109 | עליה |  |
| 8:19.110 | מן |  |
| 8:19.111 | שמיא | , |
| 8:19.112 | שחששו |  |
| 8:19.113 | שמא |  |
| 8:19.114 | בזכות |  |
| 8:19.115 | התורה |  |
| 8:19.116 | שעסקו |  |
| 8:19.117 | בה |  |
| 8:19.118 | אע"פ |  |
| 8:19.119 | שמרמזים |  |
| 8:19.120 | בה |  |
| 8:19.121 | דברי |  |
| 8:19.122 | מינות |  |
| 8:19.123 | ירחמו |  |
| 8:19.124 | עליהם |  |
| 8:19.125 | מן |  |
| 8:19.126 | השמים | , |
| 8:19.127 | כמו |  |
| 8:19.128 | שמצינו |  |
| 8:19.129 | באלישע |  |
| 8:19.130 | אחר |  |
| 8:19.131 | שעמדה |  |
| 8:19.132 | לו |  |
| 8:19.133 | זכות |  |
| 8:19.134 | תורה |  |
| 8:19.135 | בקצת | , |
| 8:19.136 | ואמרו |  |
| 8:19.137 | בשמים |  |
| 8:19.138 | לא |  |
| 8:19.139 | מידן |  |
| 8:19.140 | נידייניה |  |
| 8:19.141 | דעסק |  |
| 8:19.142 | בתורה |  |
| 8:19.143 | וכו | '. |
| 8:19.144 | אמר |  |
| 8:19.145 | להו |  |
| 8:19.146 | נביא |  |
| 8:19.147 | שדויה |  |
| 8:19.148 | בדורא |  |
| 8:19.149 | דאברא |  |
| 8:19.150 | וכו | ' |
| 8:19.151 | שהבטיחם |  |
| 8:19.152 | הנביא |  |
| 8:19.153 | זכריה |  |
| 8:19.154 | ואמר |  |
| 8:19.155 | להם |  |
| 8:19.156 | אל |  |
| 8:19.157 | תחושו |  |
| 8:19.158 | לקול |  |
| 8:19.159 | צעקתם |  |
| 8:19.160 | וצוה |  |
| 8:19.161 | אותם |  |
| 8:19.162 | לשקע |  |
| 8:19.163 | ולשכח |  |
| 8:19.164 | מלב |  |
| 8:19.165 | בני |  |
| 8:19.166 | ישראל |  |
| 8:19.167 | כל |  |
| 8:19.168 | אותן |  |
| 8:19.169 | הדרשות |  |
| 8:19.170 | והרמזים |  |
| 8:19.171 | של |  |
| 8:19.172 | מינות |  |
| 8:19.173 | כדודא |  |
| 8:19.174 | דאברא |  |
| 8:19.175 | ששוקע |  |
| 8:19.176 | בים | , |
| 8:19.177 | כמ"ד |  |
| 8:19.178 | צללו |  |
| 8:19.179 | כעופרת |  |
| 8:19.180 | במים |  |
| 8:19.181 | אדירים |  |
| 8:19.182 | של |  |
| 8:19.183 | תורת |  |
| 8:19.184 | אמת |  |
| 8:19.185 | המקובלת |  |
| 8:19.186 | ממשה |  |
| 8:19.187 | רבינו |  |
| 8:19.188 | ע"ה |  |
| 8:19.189 | להקים |  |
| 8:19.190 | בה |  |
| 8:19.191 | בתי |  |
| 8:19.192 | מדרשות |  |
| 8:19.193 | בכל |  |
| 8:19.194 | עיר |  |
| 8:19.195 | ובכל |  |
| 8:19.196 | פלך | , |
| 8:19.197 | ולהעמיד |  |
| 8:19.198 | תלמידים |  |
| 8:19.199 | הרבה |  |
| 8:19.200 | כדאמרינן |  |
| 8:19.201 | בפ | ' |
| 8:19.202 | קמא |  |
| 8:19.203 | דאבות | , |
| 8:19.204 | כי |  |
| 8:19.205 | אנשי |  |
| 8:19.206 | כנה"ג |  |
| 8:19.207 | אמרו |  |
| 8:19.208 | שלשה |  |
| 8:19.209 | דברים |  |
| 8:19.210 | היו |  |
| 8:19.211 | מתונים |  |
| 8:19.212 | בדין |  |
| 8:19.213 | והעמידו |  |
| 8:19.214 | תלמידים |  |
| 8:19.215 | הרבה | , |
| 8:19.216 | ועשו |  |
| 8:19.217 | סייג |  |
| 8:19.218 | לתורה |  |
| 8:19.219 | שלא |  |
| 8:19.220 | לדרוש |  |
| 8:19.221 | בה |  |
| 8:19.222 | כל |  |
| 8:19.223 | איש |  |
| 8:19.224 | ככל |  |
| 8:19.225 | העולה |  |
| 8:19.226 | על |  |
| 8:19.227 | רוחו | , |
| 8:19.228 | אלא |  |
| 8:19.229 | על |  |
| 8:19.230 | פי |  |
| 8:19.231 | הכללים |  |
| 8:19.232 | והיסודות |  |
| 8:19.233 | האמתיים |  |
| 8:19.234 | שנמסרו |  |
| 8:19.235 | למשה |  |
| 8:19.236 | רבינו |  |
| 8:19.237 | ע"ה |  |
| 8:19.238 | בסיני | ! |
| 8:19.239 | וזוהי |  |
| 8:19.240 | מליצת |  |
| 8:19.241 | חפיוה |  |
| 8:19.242 | לפומא |  |
| 8:19.243 | באברא |  |
| 8:19.244 | שמנעו |  |
| 8:19.245 | וסתמו |  |
| 8:19.246 | את |  |
| 8:19.247 | פי |  |
| 8:19.248 | החכמים |  |
| 8:19.249 | בעיניהם |  |
| 8:19.250 | מלדרוש |  |
| 8:19.251 | בתורה |  |
| 8:19.252 | הזיות |  |
| 8:19.253 | והבלים |  |
| 8:19.254 | המביאים |  |
| 8:19.255 | לידי |  |
| 8:19.256 | מינות | , |
| 8:19.257 | והשבעות |  |
| 8:19.258 | רוחות |  |
| 8:19.259 | של |  |
| 8:19.260 | מתים |  |
| 8:19.261 | לבוא |  |
| 8:19.262 | להגיד |  |
| 8:19.263 | להם |  |
| 8:19.264 | נסתרות |  |
| 8:19.265 | באברא | , |
| 8:19.266 | אשר |  |
| 8:19.267 | אם |  |
| 8:19.268 | תקים |  |
| 8:19.269 | עליו |  |
| 8:19.270 | לא |  |
| 8:19.271 | ישמע |  |
| 8:19.272 | קול | , |
| 8:19.273 | כן |  |
| 8:19.274 | נעשו |  |
| 8:19.275 | דורשי |  |
| 8:19.276 | דרשות |  |
| 8:19.277 | של |  |
| 8:19.278 | דופי |  |
| 8:19.279 | והזיות |  |
| 8:19.280 | כחרשים | ! |
| 8:19.281 | ובזה |  |
| 8:19.282 | נתבטל |  |
| 8:19.283 | יצרא |  |
| 8:19.284 | דע"ז | , |
| 8:19.285 | ומן |  |
| 8:19.286 | השמים |  |
| 8:19.287 | אסתייעא |  |
| 8:19.288 | מלתא | , |
| 8:19.289 | ועלתה |  |
| 8:19.290 | בידם |  |
| 8:19.291 | לשכח |  |
| 8:19.292 | שמות |  |
| 8:19.293 | האלילים |  |
| 8:19.294 | אשר |  |
| 8:19.295 | ידמו |  |
| 8:19.296 | התרים |  |
| 8:19.297 | אחרי |  |
| 8:19.298 | לבבם |  |
| 8:19.299 | לחשוב |  |
| 8:19.300 | רבוי |  |
| 8:19.301 | באלוהות | , |
| 8:19.302 | ושיש |  |
| 8:19.303 | כמה |  |
| 8:19.304 | עלות |  |
| 8:19.305 | וסבות |  |
| 8:19.306 | המשתלשלים |  |
| 8:19.307 | זה |  |
| 8:19.308 | מזה | , |
| 8:19.309 | ונוטלים |  |
| 8:19.310 | רשות |  |
| 8:19.311 | זה |  |
| 8:19.312 | מזה |  |
| 8:19.313 | לברוא |  |
| 8:19.314 | את |  |
| 8:19.315 | אשר |  |
| 8:19.316 | יחפוץ |  |
| 8:19.317 | לברוא |  |
| 8:19.318 | אותו |  |
| 8:19.319 | כאימא |  |
| 8:19.320 | שנטלה |  |
| 8:19.321 | רשות |  |
| 8:19.322 | מאבא |  |
| 8:19.323 | לברוא |  |
| 8:19.324 | אדם |  |
| 8:19.325 | הראשון | , |
| 8:19.326 | וכל |  |
| 8:19.327 | אחד |  |
| 8:19.328 | מאלו |  |
| 8:19.329 | הצורות |  |
| 8:19.330 | נקרא |  |
| 8:19.331 | בשם |  |
| 8:19.332 | ידוע | , |
| 8:19.333 | ונקרא |  |
| 8:19.334 | בשם |  |
| 8:19.335 | הוי'ה |  |
| 8:19.336 | ואדנות |  |
| 8:19.337 | ושאר |  |
| 8:19.338 | השמות |  |
| 8:19.339 | המיוחדים |  |
| 8:19.340 | לשי"ת | , |
| 8:19.341 | ושהעבודה |  |
| 8:19.342 | והתפלה |  |
| 8:19.343 | להם |  |
| 8:19.344 | לבדם |  |
| 8:19.345 | אעפ"י |  |
| 8:19.346 | שהם |  |
| 8:19.347 | נבראים | , |
| 8:19.348 | אבל |  |
| 8:19.349 | האלוה |  |
| 8:19.350 | העליון |  |
| 8:19.351 | איננו |  |
| 8:19.352 | ראוי |  |
| 8:19.353 | לשום |  |
| 8:19.354 | שבח |  |
| 8:19.355 | ותהלה |  |
| 8:19.356 | לרוב |  |
| 8:19.357 | מעלתו |  |
| 8:19.358 | והעלמו | , |
| 8:19.359 | ואין |  |
| 8:19.360 | לו |  |
| 8:19.361 | דמות |  |
| 8:19.362 | וצורה |  |
| 8:19.363 | ואין |  |
| 8:19.364 | לו |  |
| 8:19.365 | אשה |  |
| 8:19.366 | ולא |  |
| 8:19.367 | שם |  |
| 8:19.368 | מוגבל | ! |
| 8:19.369 | כל |  |
| 8:19.370 | מחשבות |  |
| 8:19.371 | מינות |  |
| 8:19.372 | הללו |  |
| 8:19.373 | השכיחום |  |
| 8:19.374 | רבותינו |  |
| 8:19.375 | אנשי |  |
| 8:19.376 | כנסת |  |
| 8:19.377 | הגדולה |  |
| 8:19.378 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

