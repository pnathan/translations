# Translation pack: milhamot-hashem perek-5-12 (Milchamot ha-Shem 5:19–5:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-12`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
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
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
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
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- אורות **דקים** → "subtle lights".
- גוף **דק** → "a subtle body".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 5:19

Hebrew (Wikisource): הנה נתבאר לך יפה טעם מה שהרחיק האר"י לומר אין לו דמות הגוף ואינו גוף. כי אעפ"י שהם אומרים שאינו גוף גס, אם כי אור דק יש לו אבל דמות הגוף, ולא נתנו אל לבם כי גם האש והאור הם גופים דקים. והם גם כן אמרו שכל עצם פרצופי האצילות שהם עיקר האלהות אשר עליהם שייך תפלה ועבודה וקריאה בעת צרה, נשמה גוף ומלבוש. הספירה עצמה גוף, ואור פנימי נשמה, ואור המקיף מלבוש, ובכמה מקומות בזהר אומר גופא דמלכא. על כן שקר הוא אצלם לומר שאין לו דמות הגוף ואינו גוף. שהרי יש לו דמות הגוף ולא יאות לפרצופים אלא שבחים אמתיים השייכים להם כפי קבלתם החדשה שהיא היפך קבלת רז"ל בעלי המשנה והתלמוד והגאונים ופוסקים הראשונים, וכמו שאני עתיד להביא דבריהם לקמן. גם מאנו לומר ראשון ואין ראשית לראשיתו שהרי בקבלתם יש כמה ראשיות לראשיתו של זעיר אנפין שהעבודה והתפלה וכל השבחים שייכים לו כמו שכתבתי לעיל מדברי גדולי המקובלים אשר כל עמודי הקבלה החדשה נשענים עליהם בפירוש דברי הזהר. ולפי המבואר בדבריהם הרי כמה עילות קדמו לזעיר אנפין, והם אין סוף, ואדם קדמאה, ואדם קדמון, ועתיק, ואריך אנפין, ואבא ואימא. ולפיכך מיאנו לשבח לזעיר אנפין שהוא האלוה השליט בכל הברואים והוא הנקרא בשם הוי"ה בשבח זה שהוא ראשון ואין ראשית לראשיתו, לפי שאינו מסכים לקבלתם. ובזהר פרשת תרומה דף קע"ח ע"א מזהיר לומר שירות ותשבחות דאית בהו קבלה אמתית. וזו הקבלה הקדומה שלנו שהוא יתברך ראשון ואין ראשית לראשיתו אינה קבלה אמתית לפי דעתם ולפיכך מיאנו לומר ארבעה בתים הראשונים מפני שאין בהם קבלה אמתית. אך אם יאמר יגדל אלהים חי וישתבח הנו אדון עולם ולכל נוצר יורה גדולתו ומלכותו וכו', שפיר דמי לדידהו בדילוג ארבעה בתים ראשונים, דמהנו אדון עולם ואילך כולהו שייכי בזעיר אנפין והוי שירות ותשבחות דאית בהו קבלה אמתית לפי דעתם. ועל כן לא חרה להאר"י רק על ד' בתים הראשונים לבד שאינם מסכימים לקבלתו.

| id | word | punctuation after |
|---|---|---|
| 5:19.1 | הנה |  |
| 5:19.2 | נתבאר |  |
| 5:19.3 | לך |  |
| 5:19.4 | יפה |  |
| 5:19.5 | טעם |  |
| 5:19.6 | מה |  |
| 5:19.7 | שהרחיק |  |
| 5:19.8 | האר"י |  |
| 5:19.9 | לומר |  |
| 5:19.10 | אין |  |
| 5:19.11 | לו |  |
| 5:19.12 | דמות |  |
| 5:19.13 | הגוף |  |
| 5:19.14 | ואינו |  |
| 5:19.15 | גוף | . |
| 5:19.16 | כי |  |
| 5:19.17 | אעפ"י |  |
| 5:19.18 | שהם |  |
| 5:19.19 | אומרים |  |
| 5:19.20 | שאינו |  |
| 5:19.21 | גוף |  |
| 5:19.22 | גס | , |
| 5:19.23 | אם |  |
| 5:19.24 | כי |  |
| 5:19.25 | אור |  |
| 5:19.26 | דק |  |
| 5:19.27 | יש |  |
| 5:19.28 | לו |  |
| 5:19.29 | אבל |  |
| 5:19.30 | דמות |  |
| 5:19.31 | הגוף | , |
| 5:19.32 | ולא |  |
| 5:19.33 | נתנו |  |
| 5:19.34 | אל |  |
| 5:19.35 | לבם |  |
| 5:19.36 | כי |  |
| 5:19.37 | גם |  |
| 5:19.38 | האש |  |
| 5:19.39 | והאור |  |
| 5:19.40 | הם |  |
| 5:19.41 | גופים |  |
| 5:19.42 | דקים | . |
| 5:19.43 | והם |  |
| 5:19.44 | גם |  |
| 5:19.45 | כן |  |
| 5:19.46 | אמרו |  |
| 5:19.47 | שכל |  |
| 5:19.48 | עצם |  |
| 5:19.49 | פרצופי |  |
| 5:19.50 | האצילות |  |
| 5:19.51 | שהם |  |
| 5:19.52 | עיקר |  |
| 5:19.53 | האלהות |  |
| 5:19.54 | אשר |  |
| 5:19.55 | עליהם |  |
| 5:19.56 | שייך |  |
| 5:19.57 | תפלה |  |
| 5:19.58 | ועבודה |  |
| 5:19.59 | וקריאה |  |
| 5:19.60 | בעת |  |
| 5:19.61 | צרה | , |
| 5:19.62 | נשמה |  |
| 5:19.63 | גוף |  |
| 5:19.64 | ומלבוש | . |
| 5:19.65 | הספירה |  |
| 5:19.66 | עצמה |  |
| 5:19.67 | גוף | , |
| 5:19.68 | ואור |  |
| 5:19.69 | פנימי |  |
| 5:19.70 | נשמה | , |
| 5:19.71 | ואור |  |
| 5:19.72 | המקיף |  |
| 5:19.73 | מלבוש | , |
| 5:19.74 | ובכמה |  |
| 5:19.75 | מקומות |  |
| 5:19.76 | בזהר |  |
| 5:19.77 | אומר |  |
| 5:19.78 | גופא |  |
| 5:19.79 | דמלכא | . |
| 5:19.80 | על |  |
| 5:19.81 | כן |  |
| 5:19.82 | שקר |  |
| 5:19.83 | הוא |  |
| 5:19.84 | אצלם |  |
| 5:19.85 | לומר |  |
| 5:19.86 | שאין |  |
| 5:19.87 | לו |  |
| 5:19.88 | דמות |  |
| 5:19.89 | הגוף |  |
| 5:19.90 | ואינו |  |
| 5:19.91 | גוף | . |
| 5:19.92 | שהרי |  |
| 5:19.93 | יש |  |
| 5:19.94 | לו |  |
| 5:19.95 | דמות |  |
| 5:19.96 | הגוף |  |
| 5:19.97 | ולא |  |
| 5:19.98 | יאות |  |
| 5:19.99 | לפרצופים |  |
| 5:19.100 | אלא |  |
| 5:19.101 | שבחים |  |
| 5:19.102 | אמתיים |  |
| 5:19.103 | השייכים |  |
| 5:19.104 | להם |  |
| 5:19.105 | כפי |  |
| 5:19.106 | קבלתם |  |
| 5:19.107 | החדשה |  |
| 5:19.108 | שהיא |  |
| 5:19.109 | היפך |  |
| 5:19.110 | קבלת |  |
| 5:19.111 | רז"ל |  |
| 5:19.112 | בעלי |  |
| 5:19.113 | המשנה |  |
| 5:19.114 | והתלמוד |  |
| 5:19.115 | והגאונים |  |
| 5:19.116 | ופוסקים |  |
| 5:19.117 | הראשונים | , |
| 5:19.118 | וכמו |  |
| 5:19.119 | שאני |  |
| 5:19.120 | עתיד |  |
| 5:19.121 | להביא |  |
| 5:19.122 | דבריהם |  |
| 5:19.123 | לקמן | . |
| 5:19.124 | גם |  |
| 5:19.125 | מאנו |  |
| 5:19.126 | לומר |  |
| 5:19.127 | ראשון |  |
| 5:19.128 | ואין |  |
| 5:19.129 | ראשית |  |
| 5:19.130 | לראשיתו |  |
| 5:19.131 | שהרי |  |
| 5:19.132 | בקבלתם |  |
| 5:19.133 | יש |  |
| 5:19.134 | כמה |  |
| 5:19.135 | ראשיות |  |
| 5:19.136 | לראשיתו |  |
| 5:19.137 | של |  |
| 5:19.138 | זעיר |  |
| 5:19.139 | אנפין |  |
| 5:19.140 | שהעבודה |  |
| 5:19.141 | והתפלה |  |
| 5:19.142 | וכל |  |
| 5:19.143 | השבחים |  |
| 5:19.144 | שייכים |  |
| 5:19.145 | לו |  |
| 5:19.146 | כמו |  |
| 5:19.147 | שכתבתי |  |
| 5:19.148 | לעיל |  |
| 5:19.149 | מדברי |  |
| 5:19.150 | גדולי |  |
| 5:19.151 | המקובלים |  |
| 5:19.152 | אשר |  |
| 5:19.153 | כל |  |
| 5:19.154 | עמודי |  |
| 5:19.155 | הקבלה |  |
| 5:19.156 | החדשה |  |
| 5:19.157 | נשענים |  |
| 5:19.158 | עליהם |  |
| 5:19.159 | בפירוש |  |
| 5:19.160 | דברי |  |
| 5:19.161 | הזהר | . |
| 5:19.162 | ולפי |  |
| 5:19.163 | המבואר |  |
| 5:19.164 | בדבריהם |  |
| 5:19.165 | הרי |  |
| 5:19.166 | כמה |  |
| 5:19.167 | עילות |  |
| 5:19.168 | קדמו |  |
| 5:19.169 | לזעיר |  |
| 5:19.170 | אנפין | , |
| 5:19.171 | והם |  |
| 5:19.172 | אין |  |
| 5:19.173 | סוף | , |
| 5:19.174 | ואדם |  |
| 5:19.175 | קדמאה | , |
| 5:19.176 | ואדם |  |
| 5:19.177 | קדמון | , |
| 5:19.178 | ועתיק | , |
| 5:19.179 | ואריך |  |
| 5:19.180 | אנפין | , |
| 5:19.181 | ואבא |  |
| 5:19.182 | ואימא | . |
| 5:19.183 | ולפיכך |  |
| 5:19.184 | מיאנו |  |
| 5:19.185 | לשבח |  |
| 5:19.186 | לזעיר |  |
| 5:19.187 | אנפין |  |
| 5:19.188 | שהוא |  |
| 5:19.189 | האלוה |  |
| 5:19.190 | השליט |  |
| 5:19.191 | בכל |  |
| 5:19.192 | הברואים |  |
| 5:19.193 | והוא |  |
| 5:19.194 | הנקרא |  |
| 5:19.195 | בשם |  |
| 5:19.196 | הוי"ה |  |
| 5:19.197 | בשבח |  |
| 5:19.198 | זה |  |
| 5:19.199 | שהוא |  |
| 5:19.200 | ראשון |  |
| 5:19.201 | ואין |  |
| 5:19.202 | ראשית |  |
| 5:19.203 | לראשיתו | , |
| 5:19.204 | לפי |  |
| 5:19.205 | שאינו |  |
| 5:19.206 | מסכים |  |
| 5:19.207 | לקבלתם | . |
| 5:19.208 | ובזהר |  |
| 5:19.209 | פרשת |  |
| 5:19.210 | תרומה |  |
| 5:19.211 | דף |  |
| 5:19.212 | קע"ח |  |
| 5:19.213 | ע"א |  |
| 5:19.214 | מזהיר |  |
| 5:19.215 | לומר |  |
| 5:19.216 | שירות |  |
| 5:19.217 | ותשבחות |  |
| 5:19.218 | דאית |  |
| 5:19.219 | בהו |  |
| 5:19.220 | קבלה |  |
| 5:19.221 | אמתית | . |
| 5:19.222 | וזו |  |
| 5:19.223 | הקבלה |  |
| 5:19.224 | הקדומה |  |
| 5:19.225 | שלנו |  |
| 5:19.226 | שהוא |  |
| 5:19.227 | יתברך |  |
| 5:19.228 | ראשון |  |
| 5:19.229 | ואין |  |
| 5:19.230 | ראשית |  |
| 5:19.231 | לראשיתו |  |
| 5:19.232 | אינה |  |
| 5:19.233 | קבלה |  |
| 5:19.234 | אמתית |  |
| 5:19.235 | לפי |  |
| 5:19.236 | דעתם |  |
| 5:19.237 | ולפיכך |  |
| 5:19.238 | מיאנו |  |
| 5:19.239 | לומר |  |
| 5:19.240 | ארבעה |  |
| 5:19.241 | בתים |  |
| 5:19.242 | הראשונים |  |
| 5:19.243 | מפני |  |
| 5:19.244 | שאין |  |
| 5:19.245 | בהם |  |
| 5:19.246 | קבלה |  |
| 5:19.247 | אמתית | . |
| 5:19.248 | אך |  |
| 5:19.249 | אם |  |
| 5:19.250 | יאמר |  |
| 5:19.251 | יגדל |  |
| 5:19.252 | אלהים |  |
| 5:19.253 | חי |  |
| 5:19.254 | וישתבח |  |
| 5:19.255 | הנו |  |
| 5:19.256 | אדון |  |
| 5:19.257 | עולם |  |
| 5:19.258 | ולכל |  |
| 5:19.259 | נוצר |  |
| 5:19.260 | יורה |  |
| 5:19.261 | גדולתו |  |
| 5:19.262 | ומלכותו |  |
| 5:19.263 | וכו | ', |
| 5:19.264 | שפיר |  |
| 5:19.265 | דמי |  |
| 5:19.266 | לדידהו |  |
| 5:19.267 | בדילוג |  |
| 5:19.268 | ארבעה |  |
| 5:19.269 | בתים |  |
| 5:19.270 | ראשונים | , |
| 5:19.271 | דמהנו |  |
| 5:19.272 | אדון |  |
| 5:19.273 | עולם |  |
| 5:19.274 | ואילך |  |
| 5:19.275 | כולהו |  |
| 5:19.276 | שייכי |  |
| 5:19.277 | בזעיר |  |
| 5:19.278 | אנפין |  |
| 5:19.279 | והוי |  |
| 5:19.280 | שירות |  |
| 5:19.281 | ותשבחות |  |
| 5:19.282 | דאית |  |
| 5:19.283 | בהו |  |
| 5:19.284 | קבלה |  |
| 5:19.285 | אמתית |  |
| 5:19.286 | לפי |  |
| 5:19.287 | דעתם | . |
| 5:19.288 | ועל |  |
| 5:19.289 | כן |  |
| 5:19.290 | לא |  |
| 5:19.291 | חרה |  |
| 5:19.292 | להאר"י |  |
| 5:19.293 | רק |  |
| 5:19.294 | על |  |
| 5:19.295 | ד | ' |
| 5:19.296 | בתים |  |
| 5:19.297 | הראשונים |  |
| 5:19.298 | לבד |  |
| 5:19.299 | שאינם |  |
| 5:19.300 | מסכימים |  |
| 5:19.301 | לקבלתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

