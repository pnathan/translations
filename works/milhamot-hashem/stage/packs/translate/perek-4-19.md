# Translation pack: milhamot-hashem perek-4-19 (Milchamot ha-Shem 4:31–4:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-19`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **הנוצרים / הנצרים** → "the Christians".
- **בא"י** → "in the Land of Israel".

## Terms

- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
- **ווי** ליה מאן דלא → "woe to him who does not".
- **תא חזי** → "come and see".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 4:31

Hebrew (Wikisource): לז) ומעתה תן דעתך והתבונן מאד בדברי ר"ת בן יחיא המובא בס' רביד הזהב והעתקיו לעיל סי' י"ג שאמר שהמשנה והתלמוד היא הקבלה האמתית והמוסכמת מכל צד שרוצה לומר שאין בה שום ספק. מה שאין כן ספרי הקבלה החדשה יש לספק בהן כמה ספקות. וכבר יצאו עליהם עסיקין כידוע לכל מבין בקי בספרות ישראל. שאין לסמוך עליהם בשום דין מדיני התורה היפך התלמוד והפוסקים כמו שכתבו בעלי הכללים בדרכי הפוסקים. כל שכן וקל וחומר ביחוד השי"ת ועבודתו אשר על זה הוזהרנו בתורה ובנביאים ובכתובים וכדברי רז"ל בכמה אזהרות כמו שנתבאר לעיל. וכדתנן למה קדמה שמע לוהיה אם שמוע וכו' וכמו שכ' הרמב"ם ז"ל בפ"א מהל' קריאת שמע שמקדימין לקרות פרשת שמע מפני שיש בה ציווי על יחוד השם ואהבתו ותלמודו שהוא העיקר הגדול שהכל תלוי בו ע"כ. וכבר כתבתי לעיל סי' ל"ב שרז"ל הקפידו מאד על כל מלתא דנפיק מינה להאמין בשתי רשויות ואמרו שמשתקין אותו. ומה יתן ומה יוסיף עוד אמרו וכלא הד. אחר שכבר אמר וחשב לנו כמה עילות שנוטלין רשות כל אחד מן העילה שלמעלה ממנו. כאלו נצטוינו לומר אחד בפינו. אבל בליבנו הם הרבה אלוהות. וזה לשון הרמב"ם במורה פרק נ' לראשון, דע אתה המעיין במאמרי זה, כי האמונה אינה הענין הנאמר בפה אבל הוא הענים המצוייר בנפש, כשיאמינו בו שהוא כן כמו שצוייר, ואם יספיק לך מן הדעות האמתיות או הנחשבות אמתתם אצלך כשתספרם במאמר מבלתי שתציירם ותאמין בהם. כל שכן כשתבקש מהם האמת. הוא קל מאוד, כמו שתמצא הרבה מן הפתאין ישמרו אמונות לא יציירו להם ענין בשום פנים. אבל אם מלאך לבך לעלות לזאת המדרגה העליונה שהיא מדגרת העיון. ושיתאמת לך שהשם ית' אחד האחדות האמתית עד שלא תמצא לו הרכבה כלל. ושאין לחשוב שום מחשבת חלוק בשום פנים. דע שאין לו ית' תואר עצמי בשום פנים ולא בשום ענין. וכמו שנמנע היותו גשם כן נמנע היותו בעל תואר עצמי. אמנם מי שהאמין שהוא אחד בעל תארים רבים כבר אמר שהוא אחד במלתו. והאמין שהוא רבים במחשבתו וזה כמאמר הנצרים הוא אחד. אבל הוא שלשה והשלשה אחד. כן מאמר האומר הוא אחד, אבל הוא בעל תארים רבים, והוא ותאריו אחד. עם הסתלק הגשמות, והאמנת הפשטות הגמורה כאלו כוונתנו וחיפושנו איך נאמר, לא איך נאמין, ואין האמונה אלא ארת הציור. כי האמונה היא האמונה במה שיצוייר בשכל. ואם יהיה עם זאת האמונה שאי אפשר חלוף זאת האמונה בשום פנים. ולא ימצא בשכל מקום דחייה לאמונה ההיא ולא לשער אפשרות חלופה תהיה אמתית. וכשתפשוט מעליך התאוות והמנהגים, ותהיה בעל תבונה, ותתבונן מה שאומר אותו באלה הפרקים הבאים בהרחקת התארים יתאמת לך מה שאמרנו בהכרח. ותהיה אז ממי שיצייר יחוד השי"ת. לא ממי שיאמר אותו בפיו ולא יצייר לו עניין. ויהיה מכת הנאמר עליהם קרוב אתה בפיהם ורחוק מכליותיהם. אבל צריך שיהיה האדם מכת מי שיצייר האמת וישיגהו אע"פ שלא ידבר בו. כמו שצוו החשובים ונאמר להם אמרו בלבבכם על משכבכם ודומו סלה עכ"ל הרב המורה ז"ל.

| id | word | punctuation after |
|---|---|---|
| 4:31.1 | לז | ) |
| 4:31.2 | ומעתה |  |
| 4:31.3 | תן |  |
| 4:31.4 | דעתך |  |
| 4:31.5 | והתבונן |  |
| 4:31.6 | מאד |  |
| 4:31.7 | בדברי |  |
| 4:31.8 | ר"ת |  |
| 4:31.9 | בן |  |
| 4:31.10 | יחיא |  |
| 4:31.11 | המובא |  |
| 4:31.12 | בס | ' |
| 4:31.13 | רביד |  |
| 4:31.14 | הזהב |  |
| 4:31.15 | והעתקיו |  |
| 4:31.16 | לעיל |  |
| 4:31.17 | סי | ' |
| 4:31.18 | י"ג |  |
| 4:31.19 | שאמר |  |
| 4:31.20 | שהמשנה |  |
| 4:31.21 | והתלמוד |  |
| 4:31.22 | היא |  |
| 4:31.23 | הקבלה |  |
| 4:31.24 | האמתית |  |
| 4:31.25 | והמוסכמת |  |
| 4:31.26 | מכל |  |
| 4:31.27 | צד |  |
| 4:31.28 | שרוצה |  |
| 4:31.29 | לומר |  |
| 4:31.30 | שאין |  |
| 4:31.31 | בה |  |
| 4:31.32 | שום |  |
| 4:31.33 | ספק | . |
| 4:31.34 | מה |  |
| 4:31.35 | שאין |  |
| 4:31.36 | כן |  |
| 4:31.37 | ספרי |  |
| 4:31.38 | הקבלה |  |
| 4:31.39 | החדשה |  |
| 4:31.40 | יש |  |
| 4:31.41 | לספק |  |
| 4:31.42 | בהן |  |
| 4:31.43 | כמה |  |
| 4:31.44 | ספקות | . |
| 4:31.45 | וכבר |  |
| 4:31.46 | יצאו |  |
| 4:31.47 | עליהם |  |
| 4:31.48 | עסיקין |  |
| 4:31.49 | כידוע |  |
| 4:31.50 | לכל |  |
| 4:31.51 | מבין |  |
| 4:31.52 | בקי |  |
| 4:31.53 | בספרות |  |
| 4:31.54 | ישראל | . |
| 4:31.55 | שאין |  |
| 4:31.56 | לסמוך |  |
| 4:31.57 | עליהם |  |
| 4:31.58 | בשום |  |
| 4:31.59 | דין |  |
| 4:31.60 | מדיני |  |
| 4:31.61 | התורה |  |
| 4:31.62 | היפך |  |
| 4:31.63 | התלמוד |  |
| 4:31.64 | והפוסקים |  |
| 4:31.65 | כמו |  |
| 4:31.66 | שכתבו |  |
| 4:31.67 | בעלי |  |
| 4:31.68 | הכללים |  |
| 4:31.69 | בדרכי |  |
| 4:31.70 | הפוסקים | . |
| 4:31.71 | כל |  |
| 4:31.72 | שכן |  |
| 4:31.73 | וקל |  |
| 4:31.74 | וחומר |  |
| 4:31.75 | ביחוד |  |
| 4:31.76 | השי"ת |  |
| 4:31.77 | ועבודתו |  |
| 4:31.78 | אשר |  |
| 4:31.79 | על |  |
| 4:31.80 | זה |  |
| 4:31.81 | הוזהרנו |  |
| 4:31.82 | בתורה |  |
| 4:31.83 | ובנביאים |  |
| 4:31.84 | ובכתובים |  |
| 4:31.85 | וכדברי |  |
| 4:31.86 | רז"ל |  |
| 4:31.87 | בכמה |  |
| 4:31.88 | אזהרות |  |
| 4:31.89 | כמו |  |
| 4:31.90 | שנתבאר |  |
| 4:31.91 | לעיל | . |
| 4:31.92 | וכדתנן |  |
| 4:31.93 | למה |  |
| 4:31.94 | קדמה |  |
| 4:31.95 | שמע |  |
| 4:31.96 | לוהיה |  |
| 4:31.97 | אם |  |
| 4:31.98 | שמוע |  |
| 4:31.99 | וכו | ' |
| 4:31.100 | וכמו |  |
| 4:31.101 | שכ | ' |
| 4:31.102 | הרמב"ם |  |
| 4:31.103 | ז"ל |  |
| 4:31.104 | בפ"א |  |
| 4:31.105 | מהל | ' |
| 4:31.106 | קריאת |  |
| 4:31.107 | שמע |  |
| 4:31.108 | שמקדימין |  |
| 4:31.109 | לקרות |  |
| 4:31.110 | פרשת |  |
| 4:31.111 | שמע |  |
| 4:31.112 | מפני |  |
| 4:31.113 | שיש |  |
| 4:31.114 | בה |  |
| 4:31.115 | ציווי |  |
| 4:31.116 | על |  |
| 4:31.117 | יחוד |  |
| 4:31.118 | השם |  |
| 4:31.119 | ואהבתו |  |
| 4:31.120 | ותלמודו |  |
| 4:31.121 | שהוא |  |
| 4:31.122 | העיקר |  |
| 4:31.123 | הגדול |  |
| 4:31.124 | שהכל |  |
| 4:31.125 | תלוי |  |
| 4:31.126 | בו |  |
| 4:31.127 | ע"כ | . |
| 4:31.128 | וכבר |  |
| 4:31.129 | כתבתי |  |
| 4:31.130 | לעיל |  |
| 4:31.131 | סי | ' |
| 4:31.132 | ל"ב |  |
| 4:31.133 | שרז"ל |  |
| 4:31.134 | הקפידו |  |
| 4:31.135 | מאד |  |
| 4:31.136 | על |  |
| 4:31.137 | כל |  |
| 4:31.138 | מלתא |  |
| 4:31.139 | דנפיק |  |
| 4:31.140 | מינה |  |
| 4:31.141 | להאמין |  |
| 4:31.142 | בשתי |  |
| 4:31.143 | רשויות |  |
| 4:31.144 | ואמרו |  |
| 4:31.145 | שמשתקין |  |
| 4:31.146 | אותו | . |
| 4:31.147 | ומה |  |
| 4:31.148 | יתן |  |
| 4:31.149 | ומה |  |
| 4:31.150 | יוסיף |  |
| 4:31.151 | עוד |  |
| 4:31.152 | אמרו |  |
| 4:31.153 | וכלא |  |
| 4:31.154 | הד | . |
| 4:31.155 | אחר |  |
| 4:31.156 | שכבר |  |
| 4:31.157 | אמר |  |
| 4:31.158 | וחשב |  |
| 4:31.159 | לנו |  |
| 4:31.160 | כמה |  |
| 4:31.161 | עילות |  |
| 4:31.162 | שנוטלין |  |
| 4:31.163 | רשות |  |
| 4:31.164 | כל |  |
| 4:31.165 | אחד |  |
| 4:31.166 | מן |  |
| 4:31.167 | העילה |  |
| 4:31.168 | שלמעלה |  |
| 4:31.169 | ממנו | . |
| 4:31.170 | כאלו |  |
| 4:31.171 | נצטוינו |  |
| 4:31.172 | לומר |  |
| 4:31.173 | אחד |  |
| 4:31.174 | בפינו | . |
| 4:31.175 | אבל |  |
| 4:31.176 | בליבנו |  |
| 4:31.177 | הם |  |
| 4:31.178 | הרבה |  |
| 4:31.179 | אלוהות | . |
| 4:31.180 | וזה |  |
| 4:31.181 | לשון |  |
| 4:31.182 | הרמב"ם |  |
| 4:31.183 | במורה |  |
| 4:31.184 | פרק |  |
| 4:31.185 | נ | ' |
| 4:31.186 | לראשון | , |
| 4:31.187 | דע |  |
| 4:31.188 | אתה |  |
| 4:31.189 | המעיין |  |
| 4:31.190 | במאמרי |  |
| 4:31.191 | זה | , |
| 4:31.192 | כי |  |
| 4:31.193 | האמונה |  |
| 4:31.194 | אינה |  |
| 4:31.195 | הענין |  |
| 4:31.196 | הנאמר |  |
| 4:31.197 | בפה |  |
| 4:31.198 | אבל |  |
| 4:31.199 | הוא |  |
| 4:31.200 | הענים |  |
| 4:31.201 | המצוייר |  |
| 4:31.202 | בנפש | , |
| 4:31.203 | כשיאמינו |  |
| 4:31.204 | בו |  |
| 4:31.205 | שהוא |  |
| 4:31.206 | כן |  |
| 4:31.207 | כמו |  |
| 4:31.208 | שצוייר | , |
| 4:31.209 | ואם |  |
| 4:31.210 | יספיק |  |
| 4:31.211 | לך |  |
| 4:31.212 | מן |  |
| 4:31.213 | הדעות |  |
| 4:31.214 | האמתיות |  |
| 4:31.215 | או |  |
| 4:31.216 | הנחשבות |  |
| 4:31.217 | אמתתם |  |
| 4:31.218 | אצלך |  |
| 4:31.219 | כשתספרם |  |
| 4:31.220 | במאמר |  |
| 4:31.221 | מבלתי |  |
| 4:31.222 | שתציירם |  |
| 4:31.223 | ותאמין |  |
| 4:31.224 | בהם | . |
| 4:31.225 | כל |  |
| 4:31.226 | שכן |  |
| 4:31.227 | כשתבקש |  |
| 4:31.228 | מהם |  |
| 4:31.229 | האמת | . |
| 4:31.230 | הוא |  |
| 4:31.231 | קל |  |
| 4:31.232 | מאוד | , |
| 4:31.233 | כמו |  |
| 4:31.234 | שתמצא |  |
| 4:31.235 | הרבה |  |
| 4:31.236 | מן |  |
| 4:31.237 | הפתאין |  |
| 4:31.238 | ישמרו |  |
| 4:31.239 | אמונות |  |
| 4:31.240 | לא |  |
| 4:31.241 | יציירו |  |
| 4:31.242 | להם |  |
| 4:31.243 | ענין |  |
| 4:31.244 | בשום |  |
| 4:31.245 | פנים | . |
| 4:31.246 | אבל |  |
| 4:31.247 | אם |  |
| 4:31.248 | מלאך |  |
| 4:31.249 | לבך |  |
| 4:31.250 | לעלות |  |
| 4:31.251 | לזאת |  |
| 4:31.252 | המדרגה |  |
| 4:31.253 | העליונה |  |
| 4:31.254 | שהיא |  |
| 4:31.255 | מדגרת |  |
| 4:31.256 | העיון | . |
| 4:31.257 | ושיתאמת |  |
| 4:31.258 | לך |  |
| 4:31.259 | שהשם |  |
| 4:31.260 | ית | ' |
| 4:31.261 | אחד |  |
| 4:31.262 | האחדות |  |
| 4:31.263 | האמתית |  |
| 4:31.264 | עד |  |
| 4:31.265 | שלא |  |
| 4:31.266 | תמצא |  |
| 4:31.267 | לו |  |
| 4:31.268 | הרכבה |  |
| 4:31.269 | כלל | . |
| 4:31.270 | ושאין |  |
| 4:31.271 | לחשוב |  |
| 4:31.272 | שום |  |
| 4:31.273 | מחשבת |  |
| 4:31.274 | חלוק |  |
| 4:31.275 | בשום |  |
| 4:31.276 | פנים | . |
| 4:31.277 | דע |  |
| 4:31.278 | שאין |  |
| 4:31.279 | לו |  |
| 4:31.280 | ית | ' |
| 4:31.281 | תואר |  |
| 4:31.282 | עצמי |  |
| 4:31.283 | בשום |  |
| 4:31.284 | פנים |  |
| 4:31.285 | ולא |  |
| 4:31.286 | בשום |  |
| 4:31.287 | ענין | . |
| 4:31.288 | וכמו |  |
| 4:31.289 | שנמנע |  |
| 4:31.290 | היותו |  |
| 4:31.291 | גשם |  |
| 4:31.292 | כן |  |
| 4:31.293 | נמנע |  |
| 4:31.294 | היותו |  |
| 4:31.295 | בעל |  |
| 4:31.296 | תואר |  |
| 4:31.297 | עצמי | . |
| 4:31.298 | אמנם |  |
| 4:31.299 | מי |  |
| 4:31.300 | שהאמין |  |
| 4:31.301 | שהוא |  |
| 4:31.302 | אחד |  |
| 4:31.303 | בעל |  |
| 4:31.304 | תארים |  |
| 4:31.305 | רבים |  |
| 4:31.306 | כבר |  |
| 4:31.307 | אמר |  |
| 4:31.308 | שהוא |  |
| 4:31.309 | אחד |  |
| 4:31.310 | במלתו | . |
| 4:31.311 | והאמין |  |
| 4:31.312 | שהוא |  |
| 4:31.313 | רבים |  |
| 4:31.314 | במחשבתו |  |
| 4:31.315 | וזה |  |
| 4:31.316 | כמאמר |  |
| 4:31.317 | הנצרים |  |
| 4:31.318 | הוא |  |
| 4:31.319 | אחד | . |
| 4:31.320 | אבל |  |
| 4:31.321 | הוא |  |
| 4:31.322 | שלשה |  |
| 4:31.323 | והשלשה |  |
| 4:31.324 | אחד | . |
| 4:31.325 | כן |  |
| 4:31.326 | מאמר |  |
| 4:31.327 | האומר |  |
| 4:31.328 | הוא |  |
| 4:31.329 | אחד | , |
| 4:31.330 | אבל |  |
| 4:31.331 | הוא |  |
| 4:31.332 | בעל |  |
| 4:31.333 | תארים |  |
| 4:31.334 | רבים | , |
| 4:31.335 | והוא |  |
| 4:31.336 | ותאריו |  |
| 4:31.337 | אחד | . |
| 4:31.338 | עם |  |
| 4:31.339 | הסתלק |  |
| 4:31.340 | הגשמות | , |
| 4:31.341 | והאמנת |  |
| 4:31.342 | הפשטות |  |
| 4:31.343 | הגמורה |  |
| 4:31.344 | כאלו |  |
| 4:31.345 | כוונתנו |  |
| 4:31.346 | וחיפושנו |  |
| 4:31.347 | איך |  |
| 4:31.348 | נאמר | , |
| 4:31.349 | לא |  |
| 4:31.350 | איך |  |
| 4:31.351 | נאמין | , |
| 4:31.352 | ואין |  |
| 4:31.353 | האמונה |  |
| 4:31.354 | אלא |  |
| 4:31.355 | ארת |  |
| 4:31.356 | הציור | . |
| 4:31.357 | כי |  |
| 4:31.358 | האמונה |  |
| 4:31.359 | היא |  |
| 4:31.360 | האמונה |  |
| 4:31.361 | במה |  |
| 4:31.362 | שיצוייר |  |
| 4:31.363 | בשכל | . |
| 4:31.364 | ואם |  |
| 4:31.365 | יהיה |  |
| 4:31.366 | עם |  |
| 4:31.367 | זאת |  |
| 4:31.368 | האמונה |  |
| 4:31.369 | שאי |  |
| 4:31.370 | אפשר |  |
| 4:31.371 | חלוף |  |
| 4:31.372 | זאת |  |
| 4:31.373 | האמונה |  |
| 4:31.374 | בשום |  |
| 4:31.375 | פנים | . |
| 4:31.376 | ולא |  |
| 4:31.377 | ימצא |  |
| 4:31.378 | בשכל |  |
| 4:31.379 | מקום |  |
| 4:31.380 | דחייה |  |
| 4:31.381 | לאמונה |  |
| 4:31.382 | ההיא |  |
| 4:31.383 | ולא |  |
| 4:31.384 | לשער |  |
| 4:31.385 | אפשרות |  |
| 4:31.386 | חלופה |  |
| 4:31.387 | תהיה |  |
| 4:31.388 | אמתית | . |
| 4:31.389 | וכשתפשוט |  |
| 4:31.390 | מעליך |  |
| 4:31.391 | התאוות |  |
| 4:31.392 | והמנהגים | , |
| 4:31.393 | ותהיה |  |
| 4:31.394 | בעל |  |
| 4:31.395 | תבונה | , |
| 4:31.396 | ותתבונן |  |
| 4:31.397 | מה |  |
| 4:31.398 | שאומר |  |
| 4:31.399 | אותו |  |
| 4:31.400 | באלה |  |
| 4:31.401 | הפרקים |  |
| 4:31.402 | הבאים |  |
| 4:31.403 | בהרחקת |  |
| 4:31.404 | התארים |  |
| 4:31.405 | יתאמת |  |
| 4:31.406 | לך |  |
| 4:31.407 | מה |  |
| 4:31.408 | שאמרנו |  |
| 4:31.409 | בהכרח | . |
| 4:31.410 | ותהיה |  |
| 4:31.411 | אז |  |
| 4:31.412 | ממי |  |
| 4:31.413 | שיצייר |  |
| 4:31.414 | יחוד |  |
| 4:31.415 | השי"ת | . |
| 4:31.416 | לא |  |
| 4:31.417 | ממי |  |
| 4:31.418 | שיאמר |  |
| 4:31.419 | אותו |  |
| 4:31.420 | בפיו |  |
| 4:31.421 | ולא |  |
| 4:31.422 | יצייר |  |
| 4:31.423 | לו |  |
| 4:31.424 | עניין | . |
| 4:31.425 | ויהיה |  |
| 4:31.426 | מכת |  |
| 4:31.427 | הנאמר |  |
| 4:31.428 | עליהם |  |
| 4:31.429 | קרוב |  |
| 4:31.430 | אתה |  |
| 4:31.431 | בפיהם |  |
| 4:31.432 | ורחוק |  |
| 4:31.433 | מכליותיהם | . |
| 4:31.434 | אבל |  |
| 4:31.435 | צריך |  |
| 4:31.436 | שיהיה |  |
| 4:31.437 | האדם |  |
| 4:31.438 | מכת |  |
| 4:31.439 | מי |  |
| 4:31.440 | שיצייר |  |
| 4:31.441 | האמת |  |
| 4:31.442 | וישיגהו |  |
| 4:31.443 | אע"פ |  |
| 4:31.444 | שלא |  |
| 4:31.445 | ידבר |  |
| 4:31.446 | בו | . |
| 4:31.447 | כמו |  |
| 4:31.448 | שצוו |  |
| 4:31.449 | החשובים |  |
| 4:31.450 | ונאמר |  |
| 4:31.451 | להם |  |
| 4:31.452 | אמרו |  |
| 4:31.453 | בלבבכם |  |
| 4:31.454 | על |  |
| 4:31.455 | משכבכם |  |
| 4:31.456 | ודומו |  |
| 4:31.457 | סלה |  |
| 4:31.458 | עכ"ל |  |
| 4:31.459 | הרב |  |
| 4:31.460 | המורה |  |
| 4:31.461 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

