# Translation pack: milhamot-hashem perek-4-16 (Milchamot ha-Shem 4:26–4:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-16`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **מרן** → "Maran" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- אל **אחר** → "another god".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **תשבורת** → "mathematics".
- ואין **צריך** לומר → "needless to say".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- **רחמנא ליצלן** → "the Merciful One save us".
- קורא **נעים / נעלם** → "pleasant reader".
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
- מי **אתי** (ketiv): translated, the qere noted.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **אללה / הדה / אלגהת** (any Judeo-Arabic passage: Saadia on Job; Choter ben
  Shelomo; a Muslim polemic; the proverb in 2:13): from the Arabic as it
  stands, never transliterated in the English. Ambiguous spelling: the evident
  word (note), the other in `apparatus`; אלי' (alif maqsura) is ilā "to"; a
  scribal letter change (ר for ד) noted, not corrected; an unexplained word
  (שברא): "[unclear]", the letters in `apparatus`. A rendering from Saadia's
  translation (אלבלד שברא): as given, Saadia's translation named in
  `apparatus`; his אמנתא / אמונתנו pairs: the Arabic as written, his Hebrew
  gloss in a note. Glosses inside the Arabic ("אתתלית (תלת קשרי דמהימנותא)")
  stay in parentheses, translated as his. (Q4.)
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


## Milchamot ha-Shem 4:26

Hebrew (Wikisource): לד) הרח"ו בספר עץ חיים הורה היתר לעצמו לדבר ולדרוש במה לפנים ולבאר בה ראשי פרקים. וחשב ליתן טעם למה נברא העולם בזמן זה ולא קודם לכן, שבזמן הקדום נתעסק הקב"ה בבריאת עולמות של מעלה, ולא הספיק לגמרן אלא עד שהגיע זמן בריאת העה"ז. ולא מצא לו הבורא ח"ו עת וזמן פנוי לברוא את העה"ז, כי כל משך הזמן הקדום היה ח"ו טרוד בהתעסקו בברית העולמות העליונים לפי שבהכרח צריך עדן ועדנין לבריאת העולם ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:26.1 | לד | ) |
| 4:26.2 | הרח"ו |  |
| 4:26.3 | בספר |  |
| 4:26.4 | עץ |  |
| 4:26.5 | חיים |  |
| 4:26.6 | הורה |  |
| 4:26.7 | היתר |  |
| 4:26.8 | לעצמו |  |
| 4:26.9 | לדבר |  |
| 4:26.10 | ולדרוש |  |
| 4:26.11 | במה |  |
| 4:26.12 | לפנים |  |
| 4:26.13 | ולבאר |  |
| 4:26.14 | בה |  |
| 4:26.15 | ראשי |  |
| 4:26.16 | פרקים | . |
| 4:26.17 | וחשב |  |
| 4:26.18 | ליתן |  |
| 4:26.19 | טעם |  |
| 4:26.20 | למה |  |
| 4:26.21 | נברא |  |
| 4:26.22 | העולם |  |
| 4:26.23 | בזמן |  |
| 4:26.24 | זה |  |
| 4:26.25 | ולא |  |
| 4:26.26 | קודם |  |
| 4:26.27 | לכן | , |
| 4:26.28 | שבזמן |  |
| 4:26.29 | הקדום |  |
| 4:26.30 | נתעסק |  |
| 4:26.31 | הקב"ה |  |
| 4:26.32 | בבריאת |  |
| 4:26.33 | עולמות |  |
| 4:26.34 | של |  |
| 4:26.35 | מעלה | , |
| 4:26.36 | ולא |  |
| 4:26.37 | הספיק |  |
| 4:26.38 | לגמרן |  |
| 4:26.39 | אלא |  |
| 4:26.40 | עד |  |
| 4:26.41 | שהגיע |  |
| 4:26.42 | זמן |  |
| 4:26.43 | בריאת |  |
| 4:26.44 | העה"ז | . |
| 4:26.45 | ולא |  |
| 4:26.46 | מצא |  |
| 4:26.47 | לו |  |
| 4:26.48 | הבורא |  |
| 4:26.49 | ח"ו |  |
| 4:26.50 | עת |  |
| 4:26.51 | וזמן |  |
| 4:26.52 | פנוי |  |
| 4:26.53 | לברוא |  |
| 4:26.54 | את |  |
| 4:26.55 | העה"ז | , |
| 4:26.56 | כי |  |
| 4:26.57 | כל |  |
| 4:26.58 | משך |  |
| 4:26.59 | הזמן |  |
| 4:26.60 | הקדום |  |
| 4:26.61 | היה |  |
| 4:26.62 | ח"ו |  |
| 4:26.63 | טרוד |  |
| 4:26.64 | בהתעסקו |  |
| 4:26.65 | בברית |  |
| 4:26.66 | העולמות |  |
| 4:26.67 | העליונים |  |
| 4:26.68 | לפי |  |
| 4:26.69 | שבהכרח |  |
| 4:26.70 | צריך |  |
| 4:26.71 | עדן |  |
| 4:26.72 | ועדנין |  |
| 4:26.73 | לבריאת |  |
| 4:26.74 | העולם |  |
| 4:26.75 | ע"ש | . |


## Milchamot ha-Shem 4:27

Hebrew (Wikisource): ונעלם מעיני הרח"ו מה שכתב הרמב"ם ז"ל במורה סוף פרק ט"ו משני וז"ל כי אלו תאמר על דרך משל כי הבורא ברא עולמות רבות קודם זה העולם כפי מנין מלוא רוחב גלגל העליון חרדל, וכל עולם ועולם מהן עמד שנים רבות בלי מספר כמלוא רוחב גלגל העליון (חרדל), היה זה כשנעריכהו אל מציאותו יתברך שאין לו תכלית כמו אלו, אמרת כי האל ברא אתמול העולם, כי כשנקיים התחלת מציאות אחר העדר (ג"א דאפיסה) אין הפרש בין שתשים זה קודם מאות אלפים שנים או מזמן קרוב מאוד ע"כ. ומזה תבין חסרון חכמה מהרח"ו שחשב להחכם מבלי חכמה. ובספר עוז לאלהים בבית קדש הקדשים פ' כ"א ד' ס"ז ע"ד העיז פניו להשיב וז"ל ואם לא תתרצה בתשובת הבחירה נשיב לך שמעת שנתהוה מלכא קדישא ברא העולמות, ובזה לא תוכל להקשות למה בראם עכשו ולא קודם, כי מעת שנתהוה בראם עכ"ל. ונמצא לפי דבריו שאין הבורא יתברך גדל מאדם הראשון אלא בחמשה ימים בלבד. ולדבריו ג"כ איך יצוייר מה שדרשו ז"ל על פסוק ואהיה אצלו אמון, שהתורה קדמה לבריאת העולם תתעק"ד דורות, כי לדבריו היתה התורה קודם התהוות הבורא, רחמנא לצלן ממחשבות זרות כאלה. וז"ל מהר'י צהרי בשם ר' סעדיה גאון פרשה בראשית הענין השביעי בא להודיענו, שאעפ"י שתראו השמים גדולי הערך במדתם כפי מה שנתבאר במופת לבעלי חכמת התשבורת בשמים ובארץ, אל תחשבו שהוצרכו זמן רב לבריאתם, אלא קורא אני להם יעמדו יחדו, ורוצה לומר בראשית בזמן קטן שבזמנים בלי עמל ובלי טורח ובלי יגיעה, וכן הוא אומר לא ייעף ולא ייגע אין חקר לתבונתו עכ"ל. ובבראשית רבה פ' א' וג' ובילקוט בראשית אמז ה' וישעיה שכ"ב, ותילים תתל"ו שאלו רז"ל אימתי נבראו המלאכים. ר' יוחנן אמר בשני נבראו הדה הוא דכ' המקרה במים עליותיו וכו' וכתיב בתריה עושה מלאכיו רוחות. ר' חנינה אמר בחמשי נבראו המלאכים הדה הא דכתיב ועוף יעופף על הארץ וכתיב ובשתים יעופף ר' לוליאני בשם ר' יצחק אמר בין על דעתיה דר' יוחנן בין על דעתיה דר' חנינה, הכל מודים שלא נברא ביום א' כלום. שלא יאמרו מיכאל היה מותח בדרומה של רקיע וגבריאל בצפונו והקב"ה ממדד באמצע אלא אנכי ה' עושה כל נוטה שמים לבדי רוקע הארץ מיאתי. מי אתי כתיב, מי היה שותף עמי בבריאתו של עולם ע"כ. וכן כ' רש"י ז"ל בפ' יו"ט של ר"ה שלא נבראו המלאכים עד יום ב' וכר' יוחנן. וכ"כ הסמ"ג בהקדמת עשין וז"ל כשברא הקב"ה את העולם בראו ביום הראשון והמלאכים נבראו ביום השני שנאמר המקרה במים עליותיו וכתיב בתריה עושה מלאכיו רוחות ע"כ ואתיא כר' יוחנן. ובילקוט תילים ויהי ערב ויהי בקר יום אחד יום יחידו של עולם שלא היה בעולם אלא הוא, הוי אומר אתה הוא ה' האלהים לבדך ע"כ. ד"א כי גדול אתה ועושה נפלאות בנוהג שבעולם מלך בשר ודם מתקלס במדינה וגדולי המדינה מתקלסין עמו לפי שנושאין עמו במשאו. אבל הקב"ה אינו כן אלא הוא לבדו ברא את העולם הוא לבדו מתקלס בעולם, הוא לבדו מתהדר בעולמו, אמר ר' תנחומא כי גדול אתה ועושה נפלאות למה? כי אתה אלהים לבדך אתה לבדך בראת את העולם ע"כ.

| id | word | punctuation after |
|---|---|---|
| 4:27.1 | ונעלם |  |
| 4:27.2 | מעיני |  |
| 4:27.3 | הרח"ו |  |
| 4:27.4 | מה |  |
| 4:27.5 | שכתב |  |
| 4:27.6 | הרמב"ם |  |
| 4:27.7 | ז"ל |  |
| 4:27.8 | במורה |  |
| 4:27.9 | סוף |  |
| 4:27.10 | פרק |  |
| 4:27.11 | ט"ו |  |
| 4:27.12 | משני |  |
| 4:27.13 | וז"ל |  |
| 4:27.14 | כי |  |
| 4:27.15 | אלו |  |
| 4:27.16 | תאמר |  |
| 4:27.17 | על |  |
| 4:27.18 | דרך |  |
| 4:27.19 | משל |  |
| 4:27.20 | כי |  |
| 4:27.21 | הבורא |  |
| 4:27.22 | ברא |  |
| 4:27.23 | עולמות |  |
| 4:27.24 | רבות |  |
| 4:27.25 | קודם |  |
| 4:27.26 | זה |  |
| 4:27.27 | העולם |  |
| 4:27.28 | כפי |  |
| 4:27.29 | מנין |  |
| 4:27.30 | מלוא |  |
| 4:27.31 | רוחב |  |
| 4:27.32 | גלגל |  |
| 4:27.33 | העליון |  |
| 4:27.34 | חרדל | , |
| 4:27.35 | וכל |  |
| 4:27.36 | עולם |  |
| 4:27.37 | ועולם |  |
| 4:27.38 | מהן |  |
| 4:27.39 | עמד |  |
| 4:27.40 | שנים |  |
| 4:27.41 | רבות |  |
| 4:27.42 | בלי |  |
| 4:27.43 | מספר |  |
| 4:27.44 | כמלוא |  |
| 4:27.45 | רוחב |  |
| 4:27.46 | גלגל |  |
| 4:27.47 | העליון |  |
| 4:27.48 | חרדל | ), |
| 4:27.49 | היה |  |
| 4:27.50 | זה |  |
| 4:27.51 | כשנעריכהו |  |
| 4:27.52 | אל |  |
| 4:27.53 | מציאותו |  |
| 4:27.54 | יתברך |  |
| 4:27.55 | שאין |  |
| 4:27.56 | לו |  |
| 4:27.57 | תכלית |  |
| 4:27.58 | כמו |  |
| 4:27.59 | אלו | , |
| 4:27.60 | אמרת |  |
| 4:27.61 | כי |  |
| 4:27.62 | האל |  |
| 4:27.63 | ברא |  |
| 4:27.64 | אתמול |  |
| 4:27.65 | העולם | , |
| 4:27.66 | כי |  |
| 4:27.67 | כשנקיים |  |
| 4:27.68 | התחלת |  |
| 4:27.69 | מציאות |  |
| 4:27.70 | אחר |  |
| 4:27.71 | העדר |  |
| 4:27.72 | ג"א |  |
| 4:27.73 | דאפיסה | ) |
| 4:27.74 | אין |  |
| 4:27.75 | הפרש |  |
| 4:27.76 | בין |  |
| 4:27.77 | שתשים |  |
| 4:27.78 | זה |  |
| 4:27.79 | קודם |  |
| 4:27.80 | מאות |  |
| 4:27.81 | אלפים |  |
| 4:27.82 | שנים |  |
| 4:27.83 | או |  |
| 4:27.84 | מזמן |  |
| 4:27.85 | קרוב |  |
| 4:27.86 | מאוד |  |
| 4:27.87 | ע"כ | . |
| 4:27.88 | ומזה |  |
| 4:27.89 | תבין |  |
| 4:27.90 | חסרון |  |
| 4:27.91 | חכמה |  |
| 4:27.92 | מהרח"ו |  |
| 4:27.93 | שחשב |  |
| 4:27.94 | להחכם |  |
| 4:27.95 | מבלי |  |
| 4:27.96 | חכמה | . |
| 4:27.97 | ובספר |  |
| 4:27.98 | עוז |  |
| 4:27.99 | לאלהים |  |
| 4:27.100 | בבית |  |
| 4:27.101 | קדש |  |
| 4:27.102 | הקדשים |  |
| 4:27.103 | פ | ' |
| 4:27.104 | כ"א |  |
| 4:27.105 | ד | ' |
| 4:27.106 | ס"ז |  |
| 4:27.107 | ע"ד |  |
| 4:27.108 | העיז |  |
| 4:27.109 | פניו |  |
| 4:27.110 | להשיב |  |
| 4:27.111 | וז"ל |  |
| 4:27.112 | ואם |  |
| 4:27.113 | לא |  |
| 4:27.114 | תתרצה |  |
| 4:27.115 | בתשובת |  |
| 4:27.116 | הבחירה |  |
| 4:27.117 | נשיב |  |
| 4:27.118 | לך |  |
| 4:27.119 | שמעת |  |
| 4:27.120 | שנתהוה |  |
| 4:27.121 | מלכא |  |
| 4:27.122 | קדישא |  |
| 4:27.123 | ברא |  |
| 4:27.124 | העולמות | , |
| 4:27.125 | ובזה |  |
| 4:27.126 | לא |  |
| 4:27.127 | תוכל |  |
| 4:27.128 | להקשות |  |
| 4:27.129 | למה |  |
| 4:27.130 | בראם |  |
| 4:27.131 | עכשו |  |
| 4:27.132 | ולא |  |
| 4:27.133 | קודם | , |
| 4:27.134 | כי |  |
| 4:27.135 | מעת |  |
| 4:27.136 | שנתהוה |  |
| 4:27.137 | בראם |  |
| 4:27.138 | עכ"ל | . |
| 4:27.139 | ונמצא |  |
| 4:27.140 | לפי |  |
| 4:27.141 | דבריו |  |
| 4:27.142 | שאין |  |
| 4:27.143 | הבורא |  |
| 4:27.144 | יתברך |  |
| 4:27.145 | גדל |  |
| 4:27.146 | מאדם |  |
| 4:27.147 | הראשון |  |
| 4:27.148 | אלא |  |
| 4:27.149 | בחמשה |  |
| 4:27.150 | ימים |  |
| 4:27.151 | בלבד | . |
| 4:27.152 | ולדבריו |  |
| 4:27.153 | ג"כ |  |
| 4:27.154 | איך |  |
| 4:27.155 | יצוייר |  |
| 4:27.156 | מה |  |
| 4:27.157 | שדרשו |  |
| 4:27.158 | ז"ל |  |
| 4:27.159 | על |  |
| 4:27.160 | פסוק |  |
| 4:27.161 | ואהיה |  |
| 4:27.162 | אצלו |  |
| 4:27.163 | אמון | , |
| 4:27.164 | שהתורה |  |
| 4:27.165 | קדמה |  |
| 4:27.166 | לבריאת |  |
| 4:27.167 | העולם |  |
| 4:27.168 | תתעק"ד |  |
| 4:27.169 | דורות | , |
| 4:27.170 | כי |  |
| 4:27.171 | לדבריו |  |
| 4:27.172 | היתה |  |
| 4:27.173 | התורה |  |
| 4:27.174 | קודם |  |
| 4:27.175 | התהוות |  |
| 4:27.176 | הבורא | , |
| 4:27.177 | רחמנא |  |
| 4:27.178 | לצלן |  |
| 4:27.179 | ממחשבות |  |
| 4:27.180 | זרות |  |
| 4:27.181 | כאלה | . |
| 4:27.182 | וז"ל |  |
| 4:27.183 | מהר'י |  |
| 4:27.184 | צהרי |  |
| 4:27.185 | בשם |  |
| 4:27.186 | ר | ' |
| 4:27.187 | סעדיה |  |
| 4:27.188 | גאון |  |
| 4:27.189 | פרשה |  |
| 4:27.190 | בראשית |  |
| 4:27.191 | הענין |  |
| 4:27.192 | השביעי |  |
| 4:27.193 | בא |  |
| 4:27.194 | להודיענו | , |
| 4:27.195 | שאעפ"י |  |
| 4:27.196 | שתראו |  |
| 4:27.197 | השמים |  |
| 4:27.198 | גדולי |  |
| 4:27.199 | הערך |  |
| 4:27.200 | במדתם |  |
| 4:27.201 | כפי |  |
| 4:27.202 | מה |  |
| 4:27.203 | שנתבאר |  |
| 4:27.204 | במופת |  |
| 4:27.205 | לבעלי |  |
| 4:27.206 | חכמת |  |
| 4:27.207 | התשבורת |  |
| 4:27.208 | בשמים |  |
| 4:27.209 | ובארץ | , |
| 4:27.210 | אל |  |
| 4:27.211 | תחשבו |  |
| 4:27.212 | שהוצרכו |  |
| 4:27.213 | זמן |  |
| 4:27.214 | רב |  |
| 4:27.215 | לבריאתם | , |
| 4:27.216 | אלא |  |
| 4:27.217 | קורא |  |
| 4:27.218 | אני |  |
| 4:27.219 | להם |  |
| 4:27.220 | יעמדו |  |
| 4:27.221 | יחדו | , |
| 4:27.222 | ורוצה |  |
| 4:27.223 | לומר |  |
| 4:27.224 | בראשית |  |
| 4:27.225 | בזמן |  |
| 4:27.226 | קטן |  |
| 4:27.227 | שבזמנים |  |
| 4:27.228 | בלי |  |
| 4:27.229 | עמל |  |
| 4:27.230 | ובלי |  |
| 4:27.231 | טורח |  |
| 4:27.232 | ובלי |  |
| 4:27.233 | יגיעה | , |
| 4:27.234 | וכן |  |
| 4:27.235 | הוא |  |
| 4:27.236 | אומר |  |
| 4:27.237 | לא |  |
| 4:27.238 | ייעף |  |
| 4:27.239 | ולא |  |
| 4:27.240 | ייגע |  |
| 4:27.241 | אין |  |
| 4:27.242 | חקר |  |
| 4:27.243 | לתבונתו |  |
| 4:27.244 | עכ"ל | . |
| 4:27.245 | ובבראשית |  |
| 4:27.246 | רבה |  |
| 4:27.247 | פ | ' |
| 4:27.248 | א | ' |
| 4:27.249 | וג | ' |
| 4:27.250 | ובילקוט |  |
| 4:27.251 | בראשית |  |
| 4:27.252 | אמז |  |
| 4:27.253 | ה | ' |
| 4:27.254 | וישעיה |  |
| 4:27.255 | שכ"ב | , |
| 4:27.256 | ותילים |  |
| 4:27.257 | תתל"ו |  |
| 4:27.258 | שאלו |  |
| 4:27.259 | רז"ל |  |
| 4:27.260 | אימתי |  |
| 4:27.261 | נבראו |  |
| 4:27.262 | המלאכים | . |
| 4:27.263 | ר | ' |
| 4:27.264 | יוחנן |  |
| 4:27.265 | אמר |  |
| 4:27.266 | בשני |  |
| 4:27.267 | נבראו |  |
| 4:27.268 | הדה |  |
| 4:27.269 | הוא |  |
| 4:27.270 | דכ | ' |
| 4:27.271 | המקרה |  |
| 4:27.272 | במים |  |
| 4:27.273 | עליותיו |  |
| 4:27.274 | וכו | ' |
| 4:27.275 | וכתיב |  |
| 4:27.276 | בתריה |  |
| 4:27.277 | עושה |  |
| 4:27.278 | מלאכיו |  |
| 4:27.279 | רוחות | . |
| 4:27.280 | ר | ' |
| 4:27.281 | חנינה |  |
| 4:27.282 | אמר |  |
| 4:27.283 | בחמשי |  |
| 4:27.284 | נבראו |  |
| 4:27.285 | המלאכים |  |
| 4:27.286 | הדה |  |
| 4:27.287 | הא |  |
| 4:27.288 | דכתיב |  |
| 4:27.289 | ועוף |  |
| 4:27.290 | יעופף |  |
| 4:27.291 | על |  |
| 4:27.292 | הארץ |  |
| 4:27.293 | וכתיב |  |
| 4:27.294 | ובשתים |  |
| 4:27.295 | יעופף |  |
| 4:27.296 | ר | ' |
| 4:27.297 | לוליאני |  |
| 4:27.298 | בשם |  |
| 4:27.299 | ר | ' |
| 4:27.300 | יצחק |  |
| 4:27.301 | אמר |  |
| 4:27.302 | בין |  |
| 4:27.303 | על |  |
| 4:27.304 | דעתיה |  |
| 4:27.305 | דר | ' |
| 4:27.306 | יוחנן |  |
| 4:27.307 | בין |  |
| 4:27.308 | על |  |
| 4:27.309 | דעתיה |  |
| 4:27.310 | דר | ' |
| 4:27.311 | חנינה | , |
| 4:27.312 | הכל |  |
| 4:27.313 | מודים |  |
| 4:27.314 | שלא |  |
| 4:27.315 | נברא |  |
| 4:27.316 | ביום |  |
| 4:27.317 | א | ' |
| 4:27.318 | כלום | . |
| 4:27.319 | שלא |  |
| 4:27.320 | יאמרו |  |
| 4:27.321 | מיכאל |  |
| 4:27.322 | היה |  |
| 4:27.323 | מותח |  |
| 4:27.324 | בדרומה |  |
| 4:27.325 | של |  |
| 4:27.326 | רקיע |  |
| 4:27.327 | וגבריאל |  |
| 4:27.328 | בצפונו |  |
| 4:27.329 | והקב"ה |  |
| 4:27.330 | ממדד |  |
| 4:27.331 | באמצע |  |
| 4:27.332 | אלא |  |
| 4:27.333 | אנכי |  |
| 4:27.334 | ה | ' |
| 4:27.335 | עושה |  |
| 4:27.336 | כל |  |
| 4:27.337 | נוטה |  |
| 4:27.338 | שמים |  |
| 4:27.339 | לבדי |  |
| 4:27.340 | רוקע |  |
| 4:27.341 | הארץ |  |
| 4:27.342 | מיאתי | . |
| 4:27.343 | מי |  |
| 4:27.344 | אתי |  |
| 4:27.345 | כתיב | , |
| 4:27.346 | מי |  |
| 4:27.347 | היה |  |
| 4:27.348 | שותף |  |
| 4:27.349 | עמי |  |
| 4:27.350 | בבריאתו |  |
| 4:27.351 | של |  |
| 4:27.352 | עולם |  |
| 4:27.353 | ע"כ | . |
| 4:27.354 | וכן |  |
| 4:27.355 | כ | ' |
| 4:27.356 | רש"י |  |
| 4:27.357 | ז"ל |  |
| 4:27.358 | בפ | ' |
| 4:27.359 | יו"ט |  |
| 4:27.360 | של |  |
| 4:27.361 | ר"ה |  |
| 4:27.362 | שלא |  |
| 4:27.363 | נבראו |  |
| 4:27.364 | המלאכים |  |
| 4:27.365 | עד |  |
| 4:27.366 | יום |  |
| 4:27.367 | ב | ' |
| 4:27.368 | וכר | ' |
| 4:27.369 | יוחנן | . |
| 4:27.370 | וכ"כ |  |
| 4:27.371 | הסמ"ג |  |
| 4:27.372 | בהקדמת |  |
| 4:27.373 | עשין |  |
| 4:27.374 | וז"ל |  |
| 4:27.375 | כשברא |  |
| 4:27.376 | הקב"ה |  |
| 4:27.377 | את |  |
| 4:27.378 | העולם |  |
| 4:27.379 | בראו |  |
| 4:27.380 | ביום |  |
| 4:27.381 | הראשון |  |
| 4:27.382 | והמלאכים |  |
| 4:27.383 | נבראו |  |
| 4:27.384 | ביום |  |
| 4:27.385 | השני |  |
| 4:27.386 | שנאמר |  |
| 4:27.387 | המקרה |  |
| 4:27.388 | במים |  |
| 4:27.389 | עליותיו |  |
| 4:27.390 | וכתיב |  |
| 4:27.391 | בתריה |  |
| 4:27.392 | עושה |  |
| 4:27.393 | מלאכיו |  |
| 4:27.394 | רוחות |  |
| 4:27.395 | ע"כ |  |
| 4:27.396 | ואתיא |  |
| 4:27.397 | כר | ' |
| 4:27.398 | יוחנן | . |
| 4:27.399 | ובילקוט |  |
| 4:27.400 | תילים |  |
| 4:27.401 | ויהי |  |
| 4:27.402 | ערב |  |
| 4:27.403 | ויהי |  |
| 4:27.404 | בקר |  |
| 4:27.405 | יום |  |
| 4:27.406 | אחד |  |
| 4:27.407 | יום |  |
| 4:27.408 | יחידו |  |
| 4:27.409 | של |  |
| 4:27.410 | עולם |  |
| 4:27.411 | שלא |  |
| 4:27.412 | היה |  |
| 4:27.413 | בעולם |  |
| 4:27.414 | אלא |  |
| 4:27.415 | הוא | , |
| 4:27.416 | הוי |  |
| 4:27.417 | אומר |  |
| 4:27.418 | אתה |  |
| 4:27.419 | הוא |  |
| 4:27.420 | ה | ' |
| 4:27.421 | האלהים |  |
| 4:27.422 | לבדך |  |
| 4:27.423 | ע"כ | . |
| 4:27.424 | ד"א |  |
| 4:27.425 | כי |  |
| 4:27.426 | גדול |  |
| 4:27.427 | אתה |  |
| 4:27.428 | ועושה |  |
| 4:27.429 | נפלאות |  |
| 4:27.430 | בנוהג |  |
| 4:27.431 | שבעולם |  |
| 4:27.432 | מלך |  |
| 4:27.433 | בשר |  |
| 4:27.434 | ודם |  |
| 4:27.435 | מתקלס |  |
| 4:27.436 | במדינה |  |
| 4:27.437 | וגדולי |  |
| 4:27.438 | המדינה |  |
| 4:27.439 | מתקלסין |  |
| 4:27.440 | עמו |  |
| 4:27.441 | לפי |  |
| 4:27.442 | שנושאין |  |
| 4:27.443 | עמו |  |
| 4:27.444 | במשאו | . |
| 4:27.445 | אבל |  |
| 4:27.446 | הקב"ה |  |
| 4:27.447 | אינו |  |
| 4:27.448 | כן |  |
| 4:27.449 | אלא |  |
| 4:27.450 | הוא |  |
| 4:27.451 | לבדו |  |
| 4:27.452 | ברא |  |
| 4:27.453 | את |  |
| 4:27.454 | העולם |  |
| 4:27.455 | הוא |  |
| 4:27.456 | לבדו |  |
| 4:27.457 | מתקלס |  |
| 4:27.458 | בעולם | , |
| 4:27.459 | הוא |  |
| 4:27.460 | לבדו |  |
| 4:27.461 | מתהדר |  |
| 4:27.462 | בעולמו | , |
| 4:27.463 | אמר |  |
| 4:27.464 | ר | ' |
| 4:27.465 | תנחומא |  |
| 4:27.466 | כי |  |
| 4:27.467 | גדול |  |
| 4:27.468 | אתה |  |
| 4:27.469 | ועושה |  |
| 4:27.470 | נפלאות |  |
| 4:27.471 | למה | ? |
| 4:27.472 | כי |  |
| 4:27.473 | אתה |  |
| 4:27.474 | אלהים |  |
| 4:27.475 | לבדך |  |
| 4:27.476 | אתה |  |
| 4:27.477 | לבדך |  |
| 4:27.478 | בראת |  |
| 4:27.479 | את |  |
| 4:27.480 | העולם |  |
| 4:27.481 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

