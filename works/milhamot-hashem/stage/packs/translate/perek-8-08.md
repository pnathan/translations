# Translation pack: milhamot-hashem perek-8-08 (Milchamot ha-Shem 8:10–8:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-08`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **רז"ל** → "our Rabbis of blessed memory".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **ריב"ש** → "Ribash" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
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
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 8:10

Hebrew (Wikisource): קו) ושרידי משכילי עמינו אשר הרגישו בזרות הקבלה החדשה כרבי' תם ן' יחיא, ורבי פרץ, ורבי שמשון מקינון, והרישב"ש והרש"ל הרחיקו את עצמם מללמדה כמו שאמר הריב"ש איני תוקע את עצמי וכו', כי הרגיש בזרותה מדרך תורתינו הקדושה והטהורה מכל סיג וחלאה, אך לא דבר כנגדה רק כלאחר יד באמרו וכי יש אלהות לספירות? והמשל שאמר לו דון יוסף ן' שושן לא ערב לו, שידע באמת שהוא כיחש לו, אך לא חפץ להעציבו לומר שקר דברת. ולפי' אמר לו והנה טוב מאד אם הוא כדבריך. כי ידע כי רק דחיה בעלמא קא מדחי ליה, כי גם אמרם כי לנשמתו אתה עובד דחייה בעלמא היא. והמה באמת עובדים להם ומקבלים אלהותם עליהם. ומיחדים את הכלים עם נשמותיהם כמו שכתב שושן סודות והרש"ב בפרשת בהר דף ק"ט ומשנת חסידים סוף מסכת העשייה ועז לאלהים שהגוף והנשמה חד אינון. וכ"כ היכל הברכה פרשת ויקרא. ואלו היתה מניעת הריב"ש מטעם שלא קיבלה מחכם מקובל היה לו ללמדה ולקבלה מפי דון יוסף ן' שושן! אלא ודאי דברים בגו! שראה בחכמתו שזו היא יציאה מדת תורתינו הקדושה לעבוד אלהים אחרים, מעשה ידי מחשבות אדם ורעיונותיו ודמיונותיו אשר עליהם באה האזהרה בתורה ולא תתורו אחרי לבבכם ואחרי עיניכם ובאה הקבלה אחרי לבבכם זו מינות, ואחרי עיניכם זו זנות, כי לבא ועינא תרי סרסורי דחטאה נינהו, ומן הידוע שההולך אחר דעתו להמציא לו רעיונות בתורה. שלא על דרך רז"ל שקבלו תורה שבעל פה שהיא המשנה והתלמוד כמפורסם ודאי יבוא לידי מינות (כדגרסינן במדרש הגדול פרשת שמיני יין ושכר אל תשת. רבי חייא פתח פקודי ה' ישרים משמחי לב, וכתיב יתהלך במישרים. משמחי לב, וכתיב ויין ישמח לבב אנוש. מה היין הזה משמח את הלב. כך דברי תורה העוסק בהן ומחשב בהן יותר מדאי סוף שהוא יוצא למינות ע"כ וכבר הבאתיו בהקדמה). ועל זה רמז הריב"ש במלות קצרות וכי יש אלהות לספירות? ולכך אמר שהמפרש את הקבלה וכו' יהיו לו לבדו, ואיני תוקע את עצמי וכו' וכן עשה רבינו תם ן' יחיא ומהרש"ל המובאים דבריהם בספר רביד הזהב ור"י אלבו בספר העיקרים.

| id | word | punctuation after |
|---|---|---|
| 8:10.1 | קו | ) |
| 8:10.2 | ושרידי |  |
| 8:10.3 | משכילי |  |
| 8:10.4 | עמינו |  |
| 8:10.5 | אשר |  |
| 8:10.6 | הרגישו |  |
| 8:10.7 | בזרות |  |
| 8:10.8 | הקבלה |  |
| 8:10.9 | החדשה |  |
| 8:10.10 | כרבי | ' |
| 8:10.11 | תם |  |
| 8:10.12 | ן | ' |
| 8:10.13 | יחיא | , |
| 8:10.14 | ורבי |  |
| 8:10.15 | פרץ | , |
| 8:10.16 | ורבי |  |
| 8:10.17 | שמשון |  |
| 8:10.18 | מקינון | , |
| 8:10.19 | והרישב"ש |  |
| 8:10.20 | והרש"ל |  |
| 8:10.21 | הרחיקו |  |
| 8:10.22 | את |  |
| 8:10.23 | עצמם |  |
| 8:10.24 | מללמדה |  |
| 8:10.25 | כמו |  |
| 8:10.26 | שאמר |  |
| 8:10.27 | הריב"ש |  |
| 8:10.28 | איני |  |
| 8:10.29 | תוקע |  |
| 8:10.30 | את |  |
| 8:10.31 | עצמי |  |
| 8:10.32 | וכו | ', |
| 8:10.33 | כי |  |
| 8:10.34 | הרגיש |  |
| 8:10.35 | בזרותה |  |
| 8:10.36 | מדרך |  |
| 8:10.37 | תורתינו |  |
| 8:10.38 | הקדושה |  |
| 8:10.39 | והטהורה |  |
| 8:10.40 | מכל |  |
| 8:10.41 | סיג |  |
| 8:10.42 | וחלאה | , |
| 8:10.43 | אך |  |
| 8:10.44 | לא |  |
| 8:10.45 | דבר |  |
| 8:10.46 | כנגדה |  |
| 8:10.47 | רק |  |
| 8:10.48 | כלאחר |  |
| 8:10.49 | יד |  |
| 8:10.50 | באמרו |  |
| 8:10.51 | וכי |  |
| 8:10.52 | יש |  |
| 8:10.53 | אלהות |  |
| 8:10.54 | לספירות | ? |
| 8:10.55 | והמשל |  |
| 8:10.56 | שאמר |  |
| 8:10.57 | לו |  |
| 8:10.58 | דון |  |
| 8:10.59 | יוסף |  |
| 8:10.60 | ן | ' |
| 8:10.61 | שושן |  |
| 8:10.62 | לא |  |
| 8:10.63 | ערב |  |
| 8:10.64 | לו | , |
| 8:10.65 | שידע |  |
| 8:10.66 | באמת |  |
| 8:10.67 | שהוא |  |
| 8:10.68 | כיחש |  |
| 8:10.69 | לו | , |
| 8:10.70 | אך |  |
| 8:10.71 | לא |  |
| 8:10.72 | חפץ |  |
| 8:10.73 | להעציבו |  |
| 8:10.74 | לומר |  |
| 8:10.75 | שקר |  |
| 8:10.76 | דברת | . |
| 8:10.77 | ולפי | ' |
| 8:10.78 | אמר |  |
| 8:10.79 | לו |  |
| 8:10.80 | והנה |  |
| 8:10.81 | טוב |  |
| 8:10.82 | מאד |  |
| 8:10.83 | אם |  |
| 8:10.84 | הוא |  |
| 8:10.85 | כדבריך | . |
| 8:10.86 | כי |  |
| 8:10.87 | ידע |  |
| 8:10.88 | כי |  |
| 8:10.89 | רק |  |
| 8:10.90 | דחיה |  |
| 8:10.91 | בעלמא |  |
| 8:10.92 | קא |  |
| 8:10.93 | מדחי |  |
| 8:10.94 | ליה | , |
| 8:10.95 | כי |  |
| 8:10.96 | גם |  |
| 8:10.97 | אמרם |  |
| 8:10.98 | כי |  |
| 8:10.99 | לנשמתו |  |
| 8:10.100 | אתה |  |
| 8:10.101 | עובד |  |
| 8:10.102 | דחייה |  |
| 8:10.103 | בעלמא |  |
| 8:10.104 | היא | . |
| 8:10.105 | והמה |  |
| 8:10.106 | באמת |  |
| 8:10.107 | עובדים |  |
| 8:10.108 | להם |  |
| 8:10.109 | ומקבלים |  |
| 8:10.110 | אלהותם |  |
| 8:10.111 | עליהם | . |
| 8:10.112 | ומיחדים |  |
| 8:10.113 | את |  |
| 8:10.114 | הכלים |  |
| 8:10.115 | עם |  |
| 8:10.116 | נשמותיהם |  |
| 8:10.117 | כמו |  |
| 8:10.118 | שכתב |  |
| 8:10.119 | שושן |  |
| 8:10.120 | סודות |  |
| 8:10.121 | והרש"ב |  |
| 8:10.122 | בפרשת |  |
| 8:10.123 | בהר |  |
| 8:10.124 | דף |  |
| 8:10.125 | ק"ט |  |
| 8:10.126 | ומשנת |  |
| 8:10.127 | חסידים |  |
| 8:10.128 | סוף |  |
| 8:10.129 | מסכת |  |
| 8:10.130 | העשייה |  |
| 8:10.131 | ועז |  |
| 8:10.132 | לאלהים |  |
| 8:10.133 | שהגוף |  |
| 8:10.134 | והנשמה |  |
| 8:10.135 | חד |  |
| 8:10.136 | אינון | . |
| 8:10.137 | וכ"כ |  |
| 8:10.138 | היכל |  |
| 8:10.139 | הברכה |  |
| 8:10.140 | פרשת |  |
| 8:10.141 | ויקרא | . |
| 8:10.142 | ואלו |  |
| 8:10.143 | היתה |  |
| 8:10.144 | מניעת |  |
| 8:10.145 | הריב"ש |  |
| 8:10.146 | מטעם |  |
| 8:10.147 | שלא |  |
| 8:10.148 | קיבלה |  |
| 8:10.149 | מחכם |  |
| 8:10.150 | מקובל |  |
| 8:10.151 | היה |  |
| 8:10.152 | לו |  |
| 8:10.153 | ללמדה |  |
| 8:10.154 | ולקבלה |  |
| 8:10.155 | מפי |  |
| 8:10.156 | דון |  |
| 8:10.157 | יוסף |  |
| 8:10.158 | ן | ' |
| 8:10.159 | שושן | ! |
| 8:10.160 | אלא |  |
| 8:10.161 | ודאי |  |
| 8:10.162 | דברים |  |
| 8:10.163 | בגו | ! |
| 8:10.164 | שראה |  |
| 8:10.165 | בחכמתו |  |
| 8:10.166 | שזו |  |
| 8:10.167 | היא |  |
| 8:10.168 | יציאה |  |
| 8:10.169 | מדת |  |
| 8:10.170 | תורתינו |  |
| 8:10.171 | הקדושה |  |
| 8:10.172 | לעבוד |  |
| 8:10.173 | אלהים |  |
| 8:10.174 | אחרים | , |
| 8:10.175 | מעשה |  |
| 8:10.176 | ידי |  |
| 8:10.177 | מחשבות |  |
| 8:10.178 | אדם |  |
| 8:10.179 | ורעיונותיו |  |
| 8:10.180 | ודמיונותיו |  |
| 8:10.181 | אשר |  |
| 8:10.182 | עליהם |  |
| 8:10.183 | באה |  |
| 8:10.184 | האזהרה |  |
| 8:10.185 | בתורה |  |
| 8:10.186 | ולא |  |
| 8:10.187 | תתורו |  |
| 8:10.188 | אחרי |  |
| 8:10.189 | לבבכם |  |
| 8:10.190 | ואחרי |  |
| 8:10.191 | עיניכם |  |
| 8:10.192 | ובאה |  |
| 8:10.193 | הקבלה |  |
| 8:10.194 | אחרי |  |
| 8:10.195 | לבבכם |  |
| 8:10.196 | זו |  |
| 8:10.197 | מינות | , |
| 8:10.198 | ואחרי |  |
| 8:10.199 | עיניכם |  |
| 8:10.200 | זו |  |
| 8:10.201 | זנות | , |
| 8:10.202 | כי |  |
| 8:10.203 | לבא |  |
| 8:10.204 | ועינא |  |
| 8:10.205 | תרי |  |
| 8:10.206 | סרסורי |  |
| 8:10.207 | דחטאה |  |
| 8:10.208 | נינהו | , |
| 8:10.209 | ומן |  |
| 8:10.210 | הידוע |  |
| 8:10.211 | שההולך |  |
| 8:10.212 | אחר |  |
| 8:10.213 | דעתו |  |
| 8:10.214 | להמציא |  |
| 8:10.215 | לו |  |
| 8:10.216 | רעיונות |  |
| 8:10.217 | בתורה | . |
| 8:10.218 | שלא |  |
| 8:10.219 | על |  |
| 8:10.220 | דרך |  |
| 8:10.221 | רז"ל |  |
| 8:10.222 | שקבלו |  |
| 8:10.223 | תורה |  |
| 8:10.224 | שבעל |  |
| 8:10.225 | פה |  |
| 8:10.226 | שהיא |  |
| 8:10.227 | המשנה |  |
| 8:10.228 | והתלמוד |  |
| 8:10.229 | כמפורסם |  |
| 8:10.230 | ודאי |  |
| 8:10.231 | יבוא |  |
| 8:10.232 | לידי |  |
| 8:10.233 | מינות |  |
| 8:10.234 | כדגרסינן |  |
| 8:10.235 | במדרש |  |
| 8:10.236 | הגדול |  |
| 8:10.237 | פרשת |  |
| 8:10.238 | שמיני |  |
| 8:10.239 | יין |  |
| 8:10.240 | ושכר |  |
| 8:10.241 | אל |  |
| 8:10.242 | תשת | . |
| 8:10.243 | רבי |  |
| 8:10.244 | חייא |  |
| 8:10.245 | פתח |  |
| 8:10.246 | פקודי |  |
| 8:10.247 | ה | ' |
| 8:10.248 | ישרים |  |
| 8:10.249 | משמחי |  |
| 8:10.250 | לב | , |
| 8:10.251 | וכתיב |  |
| 8:10.252 | יתהלך |  |
| 8:10.253 | במישרים | . |
| 8:10.254 | משמחי |  |
| 8:10.255 | לב | , |
| 8:10.256 | וכתיב |  |
| 8:10.257 | ויין |  |
| 8:10.258 | ישמח |  |
| 8:10.259 | לבב |  |
| 8:10.260 | אנוש | . |
| 8:10.261 | מה |  |
| 8:10.262 | היין |  |
| 8:10.263 | הזה |  |
| 8:10.264 | משמח |  |
| 8:10.265 | את |  |
| 8:10.266 | הלב | . |
| 8:10.267 | כך |  |
| 8:10.268 | דברי |  |
| 8:10.269 | תורה |  |
| 8:10.270 | העוסק |  |
| 8:10.271 | בהן |  |
| 8:10.272 | ומחשב |  |
| 8:10.273 | בהן |  |
| 8:10.274 | יותר |  |
| 8:10.275 | מדאי |  |
| 8:10.276 | סוף |  |
| 8:10.277 | שהוא |  |
| 8:10.278 | יוצא |  |
| 8:10.279 | למינות |  |
| 8:10.280 | ע"כ |  |
| 8:10.281 | וכבר |  |
| 8:10.282 | הבאתיו |  |
| 8:10.283 | בהקדמה | ). |
| 8:10.284 | ועל |  |
| 8:10.285 | זה |  |
| 8:10.286 | רמז |  |
| 8:10.287 | הריב"ש |  |
| 8:10.288 | במלות |  |
| 8:10.289 | קצרות |  |
| 8:10.290 | וכי |  |
| 8:10.291 | יש |  |
| 8:10.292 | אלהות |  |
| 8:10.293 | לספירות | ? |
| 8:10.294 | ולכך |  |
| 8:10.295 | אמר |  |
| 8:10.296 | שהמפרש |  |
| 8:10.297 | את |  |
| 8:10.298 | הקבלה |  |
| 8:10.299 | וכו | ' |
| 8:10.300 | יהיו |  |
| 8:10.301 | לו |  |
| 8:10.302 | לבדו | , |
| 8:10.303 | ואיני |  |
| 8:10.304 | תוקע |  |
| 8:10.305 | את |  |
| 8:10.306 | עצמי |  |
| 8:10.307 | וכו | ' |
| 8:10.308 | וכן |  |
| 8:10.309 | עשה |  |
| 8:10.310 | רבינו |  |
| 8:10.311 | תם |  |
| 8:10.312 | ן | ' |
| 8:10.313 | יחיא |  |
| 8:10.314 | ומהרש"ל |  |
| 8:10.315 | המובאים |  |
| 8:10.316 | דבריהם |  |
| 8:10.317 | בספר |  |
| 8:10.318 | רביד |  |
| 8:10.319 | הזהב |  |
| 8:10.320 | ור"י |  |
| 8:10.321 | אלבו |  |
| 8:10.322 | בספר |  |
| 8:10.323 | העיקרים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

