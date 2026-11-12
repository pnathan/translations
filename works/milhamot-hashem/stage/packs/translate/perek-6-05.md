# Translation pack: milhamot-hashem perek-6-05 (Milchamot ha-Shem 6:10–6:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-05`.

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
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- בעל **העיקרים** → "the author of the Ikkarim".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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
- מאמיני **השלוש** → "the believers in the Trinity".
## Terms

- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
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
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
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
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
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


## Milchamot ha-Shem 6:10

Hebrew (Wikisource): ובשיר השירים רבה פר' א' איתה, אמר ריש לקיש ולמה הוא אמת. אלף בראש האותיות, מ"ם באצמע, תיו בסופן. לומר אני ראשון ואני אחרון ומבלעדי אין אלהים. אני ראשון שלא קיבלתי מלכותי מאחר, ואני אחרון שאיני מוסרה לאחר שאינו בעולם. ומבלעדי אין אלהים שאין לי שני ע"כ. ופי' הר' עץ יוסף ששלשה שרשים מיוחדים באלהותו ית' הא' בקדמותו, והשני אחדותו, והג' נצחיותו ע"כ. הרי נתבאר לך מזה שהירושלמי ושמות רבה ושיר השירים רבה כולם מסכימים לעיקרים שכתב הרמב"ם באחדותו ית' וקדמותו. ודברי הזהר הומקובלים החדשים סותרים העיקרים ההם. לפי שנמשכו אחר מאמיני השלוש לעבוד את הבן כמו שנתבאר. ורז"ל הפליגו בהרחקת בעלי אמונה זו עד שאסרו לישא וליתן עמהם כדגרסינן בע"ז פרק אין מעמידין לא ישא ויתן עם המינים ואין מתרפאין מהן אפי' לחיי שעה. ומעשה בבן דמא בן אחותו של ר' ישמעאל שהכישו נחש ובא יעקב איש כפר סכניא לרפאתו ולא הניחו ר' ישמעאל. אמר לו ישמעאל אחי, הנח לי ואתרפא ממנו. ואני אביא לך מקרא מן התורה שהוא מותר. ולא הספיק לגמור את הדבר עד שיצאה נשמתו ומת. קרי עליה ר' ישמעאל אשריך בן דמא שגופך טהור ונשמתך יצאה בטהרה, ולא עברת על דברי חביריך שהיו אומרים ופורץ גדר ישכנו נחש, ומסיים שם שאני מינות משאר איסורין של תורה. משום דמשכא ואתי למשך בתריהו ע"כ. טהכי נמי אמרו בירוש' שבת פרק ח' שרצים. בר ברוה דר' יהושע בן לוי הוה ליה בלי את בר נש ולחש ליה מן שמיה דישו פנדירא ואינשם כד נפיק אמר ליה מאן לחשת ליה, אמר ליה מילת פלן. אמר ליה נוח הוה ליה אלו הוה מאית. והות לי' כן. בשגגה שיוצא מלפני השליט. עוד שם מעשה בר' אלעזר בן דמא שנשכו נחש ובא יעקב איש כפר סמא (הוא יעקב איש כפר סכניא הנז' בבבלי שהוא מתלמידי יש"ו) משם של יש"ו פנדריא לרפותו ולא הניח לו ר' ישמעאל. אמר לו אני מביא לך ראיה שירפאני. לא הספיק להביא ראיה עד שמת בן דמא. אמר לו ר' ישמעאל אשריך בן דמא שיצאת בשלום מן העולם ולא פרצת גדירן של חכמים דכ' ופורץ ישכנו נחש וכו' ומאי הוה ליה מימר אשר יעשה אותם האדם וחי בהם ע"כ. מכל זה אתה למד כמה הפליגו והרחיקו שלא לישא וליתן עם מאמיני השלשה ואין מתרפאין מהן אפי' במקום סכנה. ואע"פ שהם אומרים שהשלשה אחד, ולפי שכל ישראל בזמן רז"ל אדוקים ודבוקים ביחוד השי"ת באמת לא הוצרכו במשנה ובתלמוד לבאר דבר זה ולומר שמי שחושב עלות נאצלים מהבורא ואומר שהם אחד עם מאצילם שהוא עובד ע"ז. כי זו היא אמונת המשיחים. ועליה אמר ר' שמעון ב"י כל המשתף שם שמים ודבר אחר נעקר מן העולם. וכבר אמרו הנצרים עובדי ע"ז הם. ויום ראשון יום אידם הוא. לפיכך אסור לשאת ולתת עמהן בארץ ישראל יום ה' ויום ו' שבכל שבת ושבת ואין צריך לומר יום ראשון עצמו שהוא אסור בכל מקום. וכן בכל אידיהן (הרמב"ם פ"ט מע"ז) ולכן כ' הסמ"ג שאין צריך להאריך בזה מפני שכל ישראל אדוקים באמונת היחוד.

| id | word | punctuation after |
|---|---|---|
| 6:10.1 | ובשיר |  |
| 6:10.2 | השירים |  |
| 6:10.3 | רבה |  |
| 6:10.4 | פר | ' |
| 6:10.5 | א | ' |
| 6:10.6 | איתה | , |
| 6:10.7 | אמר |  |
| 6:10.8 | ריש |  |
| 6:10.9 | לקיש |  |
| 6:10.10 | ולמה |  |
| 6:10.11 | הוא |  |
| 6:10.12 | אמת | . |
| 6:10.13 | אלף |  |
| 6:10.14 | בראש |  |
| 6:10.15 | האותיות | , |
| 6:10.16 | מ"ם |  |
| 6:10.17 | באצמע | , |
| 6:10.18 | תיו |  |
| 6:10.19 | בסופן | . |
| 6:10.20 | לומר |  |
| 6:10.21 | אני |  |
| 6:10.22 | ראשון |  |
| 6:10.23 | ואני |  |
| 6:10.24 | אחרון |  |
| 6:10.25 | ומבלעדי |  |
| 6:10.26 | אין |  |
| 6:10.27 | אלהים | . |
| 6:10.28 | אני |  |
| 6:10.29 | ראשון |  |
| 6:10.30 | שלא |  |
| 6:10.31 | קיבלתי |  |
| 6:10.32 | מלכותי |  |
| 6:10.33 | מאחר | , |
| 6:10.34 | ואני |  |
| 6:10.35 | אחרון |  |
| 6:10.36 | שאיני |  |
| 6:10.37 | מוסרה |  |
| 6:10.38 | לאחר |  |
| 6:10.39 | שאינו |  |
| 6:10.40 | בעולם | . |
| 6:10.41 | ומבלעדי |  |
| 6:10.42 | אין |  |
| 6:10.43 | אלהים |  |
| 6:10.44 | שאין |  |
| 6:10.45 | לי |  |
| 6:10.46 | שני |  |
| 6:10.47 | ע"כ | . |
| 6:10.48 | ופי | ' |
| 6:10.49 | הר | ' |
| 6:10.50 | עץ |  |
| 6:10.51 | יוסף |  |
| 6:10.52 | ששלשה |  |
| 6:10.53 | שרשים |  |
| 6:10.54 | מיוחדים |  |
| 6:10.55 | באלהותו |  |
| 6:10.56 | ית | ' |
| 6:10.57 | הא | ' |
| 6:10.58 | בקדמותו | , |
| 6:10.59 | והשני |  |
| 6:10.60 | אחדותו | , |
| 6:10.61 | והג | ' |
| 6:10.62 | נצחיותו |  |
| 6:10.63 | ע"כ | . |
| 6:10.64 | הרי |  |
| 6:10.65 | נתבאר |  |
| 6:10.66 | לך |  |
| 6:10.67 | מזה |  |
| 6:10.68 | שהירושלמי |  |
| 6:10.69 | ושמות |  |
| 6:10.70 | רבה |  |
| 6:10.71 | ושיר |  |
| 6:10.72 | השירים |  |
| 6:10.73 | רבה |  |
| 6:10.74 | כולם |  |
| 6:10.75 | מסכימים |  |
| 6:10.76 | לעיקרים |  |
| 6:10.77 | שכתב |  |
| 6:10.78 | הרמב"ם |  |
| 6:10.79 | באחדותו |  |
| 6:10.80 | ית | ' |
| 6:10.81 | וקדמותו | . |
| 6:10.82 | ודברי |  |
| 6:10.83 | הזהר |  |
| 6:10.84 | הומקובלים |  |
| 6:10.85 | החדשים |  |
| 6:10.86 | סותרים |  |
| 6:10.87 | העיקרים |  |
| 6:10.88 | ההם | . |
| 6:10.89 | לפי |  |
| 6:10.90 | שנמשכו |  |
| 6:10.91 | אחר |  |
| 6:10.92 | מאמיני |  |
| 6:10.93 | השלוש |  |
| 6:10.94 | לעבוד |  |
| 6:10.95 | את |  |
| 6:10.96 | הבן |  |
| 6:10.97 | כמו |  |
| 6:10.98 | שנתבאר | . |
| 6:10.99 | ורז"ל |  |
| 6:10.100 | הפליגו |  |
| 6:10.101 | בהרחקת |  |
| 6:10.102 | בעלי |  |
| 6:10.103 | אמונה |  |
| 6:10.104 | זו |  |
| 6:10.105 | עד |  |
| 6:10.106 | שאסרו |  |
| 6:10.107 | לישא |  |
| 6:10.108 | וליתן |  |
| 6:10.109 | עמהם |  |
| 6:10.110 | כדגרסינן |  |
| 6:10.111 | בע"ז |  |
| 6:10.112 | פרק |  |
| 6:10.113 | אין |  |
| 6:10.114 | מעמידין |  |
| 6:10.115 | לא |  |
| 6:10.116 | ישא |  |
| 6:10.117 | ויתן |  |
| 6:10.118 | עם |  |
| 6:10.119 | המינים |  |
| 6:10.120 | ואין |  |
| 6:10.121 | מתרפאין |  |
| 6:10.122 | מהן |  |
| 6:10.123 | אפי | ' |
| 6:10.124 | לחיי |  |
| 6:10.125 | שעה | . |
| 6:10.126 | ומעשה |  |
| 6:10.127 | בבן |  |
| 6:10.128 | דמא |  |
| 6:10.129 | בן |  |
| 6:10.130 | אחותו |  |
| 6:10.131 | של |  |
| 6:10.132 | ר | ' |
| 6:10.133 | ישמעאל |  |
| 6:10.134 | שהכישו |  |
| 6:10.135 | נחש |  |
| 6:10.136 | ובא |  |
| 6:10.137 | יעקב |  |
| 6:10.138 | איש |  |
| 6:10.139 | כפר |  |
| 6:10.140 | סכניא |  |
| 6:10.141 | לרפאתו |  |
| 6:10.142 | ולא |  |
| 6:10.143 | הניחו |  |
| 6:10.144 | ר | ' |
| 6:10.145 | ישמעאל | . |
| 6:10.146 | אמר |  |
| 6:10.147 | לו |  |
| 6:10.148 | ישמעאל |  |
| 6:10.149 | אחי | , |
| 6:10.150 | הנח |  |
| 6:10.151 | לי |  |
| 6:10.152 | ואתרפא |  |
| 6:10.153 | ממנו | . |
| 6:10.154 | ואני |  |
| 6:10.155 | אביא |  |
| 6:10.156 | לך |  |
| 6:10.157 | מקרא |  |
| 6:10.158 | מן |  |
| 6:10.159 | התורה |  |
| 6:10.160 | שהוא |  |
| 6:10.161 | מותר | . |
| 6:10.162 | ולא |  |
| 6:10.163 | הספיק |  |
| 6:10.164 | לגמור |  |
| 6:10.165 | את |  |
| 6:10.166 | הדבר |  |
| 6:10.167 | עד |  |
| 6:10.168 | שיצאה |  |
| 6:10.169 | נשמתו |  |
| 6:10.170 | ומת | . |
| 6:10.171 | קרי |  |
| 6:10.172 | עליה |  |
| 6:10.173 | ר | ' |
| 6:10.174 | ישמעאל |  |
| 6:10.175 | אשריך |  |
| 6:10.176 | בן |  |
| 6:10.177 | דמא |  |
| 6:10.178 | שגופך |  |
| 6:10.179 | טהור |  |
| 6:10.180 | ונשמתך |  |
| 6:10.181 | יצאה |  |
| 6:10.182 | בטהרה | , |
| 6:10.183 | ולא |  |
| 6:10.184 | עברת |  |
| 6:10.185 | על |  |
| 6:10.186 | דברי |  |
| 6:10.187 | חביריך |  |
| 6:10.188 | שהיו |  |
| 6:10.189 | אומרים |  |
| 6:10.190 | ופורץ |  |
| 6:10.191 | גדר |  |
| 6:10.192 | ישכנו |  |
| 6:10.193 | נחש | , |
| 6:10.194 | ומסיים |  |
| 6:10.195 | שם |  |
| 6:10.196 | שאני |  |
| 6:10.197 | מינות |  |
| 6:10.198 | משאר |  |
| 6:10.199 | איסורין |  |
| 6:10.200 | של |  |
| 6:10.201 | תורה | . |
| 6:10.202 | משום |  |
| 6:10.203 | דמשכא |  |
| 6:10.204 | ואתי |  |
| 6:10.205 | למשך |  |
| 6:10.206 | בתריהו |  |
| 6:10.207 | ע"כ | . |
| 6:10.208 | טהכי |  |
| 6:10.209 | נמי |  |
| 6:10.210 | אמרו |  |
| 6:10.211 | בירוש | ' |
| 6:10.212 | שבת |  |
| 6:10.213 | פרק |  |
| 6:10.214 | ח | ' |
| 6:10.215 | שרצים | . |
| 6:10.216 | בר |  |
| 6:10.217 | ברוה |  |
| 6:10.218 | דר | ' |
| 6:10.219 | יהושע |  |
| 6:10.220 | בן |  |
| 6:10.221 | לוי |  |
| 6:10.222 | הוה |  |
| 6:10.223 | ליה |  |
| 6:10.224 | בלי |  |
| 6:10.225 | את |  |
| 6:10.226 | בר |  |
| 6:10.227 | נש |  |
| 6:10.228 | ולחש |  |
| 6:10.229 | ליה |  |
| 6:10.230 | מן |  |
| 6:10.231 | שמיה |  |
| 6:10.232 | דישו |  |
| 6:10.233 | פנדירא |  |
| 6:10.234 | ואינשם |  |
| 6:10.235 | כד |  |
| 6:10.236 | נפיק |  |
| 6:10.237 | אמר |  |
| 6:10.238 | ליה |  |
| 6:10.239 | מאן |  |
| 6:10.240 | לחשת |  |
| 6:10.241 | ליה | , |
| 6:10.242 | אמר |  |
| 6:10.243 | ליה |  |
| 6:10.244 | מילת |  |
| 6:10.245 | פלן | . |
| 6:10.246 | אמר |  |
| 6:10.247 | ליה |  |
| 6:10.248 | נוח |  |
| 6:10.249 | הוה |  |
| 6:10.250 | ליה |  |
| 6:10.251 | אלו |  |
| 6:10.252 | הוה |  |
| 6:10.253 | מאית | . |
| 6:10.254 | והות |  |
| 6:10.255 | לי | ' |
| 6:10.256 | כן | . |
| 6:10.257 | בשגגה |  |
| 6:10.258 | שיוצא |  |
| 6:10.259 | מלפני |  |
| 6:10.260 | השליט | . |
| 6:10.261 | עוד |  |
| 6:10.262 | שם |  |
| 6:10.263 | מעשה |  |
| 6:10.264 | בר | ' |
| 6:10.265 | אלעזר |  |
| 6:10.266 | בן |  |
| 6:10.267 | דמא |  |
| 6:10.268 | שנשכו |  |
| 6:10.269 | נחש |  |
| 6:10.270 | ובא |  |
| 6:10.271 | יעקב |  |
| 6:10.272 | איש |  |
| 6:10.273 | כפר |  |
| 6:10.274 | סמא |  |
| 6:10.275 | הוא |  |
| 6:10.276 | יעקב |  |
| 6:10.277 | איש |  |
| 6:10.278 | כפר |  |
| 6:10.279 | סכניא |  |
| 6:10.280 | הנז | ' |
| 6:10.281 | בבבלי |  |
| 6:10.282 | שהוא |  |
| 6:10.283 | מתלמידי |  |
| 6:10.284 | יש"ו | ) |
| 6:10.285 | משם |  |
| 6:10.286 | של |  |
| 6:10.287 | יש"ו |  |
| 6:10.288 | פנדריא |  |
| 6:10.289 | לרפותו |  |
| 6:10.290 | ולא |  |
| 6:10.291 | הניח |  |
| 6:10.292 | לו |  |
| 6:10.293 | ר | ' |
| 6:10.294 | ישמעאל | . |
| 6:10.295 | אמר |  |
| 6:10.296 | לו |  |
| 6:10.297 | אני |  |
| 6:10.298 | מביא |  |
| 6:10.299 | לך |  |
| 6:10.300 | ראיה |  |
| 6:10.301 | שירפאני | . |
| 6:10.302 | לא |  |
| 6:10.303 | הספיק |  |
| 6:10.304 | להביא |  |
| 6:10.305 | ראיה |  |
| 6:10.306 | עד |  |
| 6:10.307 | שמת |  |
| 6:10.308 | בן |  |
| 6:10.309 | דמא | . |
| 6:10.310 | אמר |  |
| 6:10.311 | לו |  |
| 6:10.312 | ר | ' |
| 6:10.313 | ישמעאל |  |
| 6:10.314 | אשריך |  |
| 6:10.315 | בן |  |
| 6:10.316 | דמא |  |
| 6:10.317 | שיצאת |  |
| 6:10.318 | בשלום |  |
| 6:10.319 | מן |  |
| 6:10.320 | העולם |  |
| 6:10.321 | ולא |  |
| 6:10.322 | פרצת |  |
| 6:10.323 | גדירן |  |
| 6:10.324 | של |  |
| 6:10.325 | חכמים |  |
| 6:10.326 | דכ | ' |
| 6:10.327 | ופורץ |  |
| 6:10.328 | ישכנו |  |
| 6:10.329 | נחש |  |
| 6:10.330 | וכו | ' |
| 6:10.331 | ומאי |  |
| 6:10.332 | הוה |  |
| 6:10.333 | ליה |  |
| 6:10.334 | מימר |  |
| 6:10.335 | אשר |  |
| 6:10.336 | יעשה |  |
| 6:10.337 | אותם |  |
| 6:10.338 | האדם |  |
| 6:10.339 | וחי |  |
| 6:10.340 | בהם |  |
| 6:10.341 | ע"כ | . |
| 6:10.342 | מכל |  |
| 6:10.343 | זה |  |
| 6:10.344 | אתה |  |
| 6:10.345 | למד |  |
| 6:10.346 | כמה |  |
| 6:10.347 | הפליגו |  |
| 6:10.348 | והרחיקו |  |
| 6:10.349 | שלא |  |
| 6:10.350 | לישא |  |
| 6:10.351 | וליתן |  |
| 6:10.352 | עם |  |
| 6:10.353 | מאמיני |  |
| 6:10.354 | השלשה |  |
| 6:10.355 | ואין |  |
| 6:10.356 | מתרפאין |  |
| 6:10.357 | מהן |  |
| 6:10.358 | אפי | ' |
| 6:10.359 | במקום |  |
| 6:10.360 | סכנה | . |
| 6:10.361 | ואע"פ |  |
| 6:10.362 | שהם |  |
| 6:10.363 | אומרים |  |
| 6:10.364 | שהשלשה |  |
| 6:10.365 | אחד | , |
| 6:10.366 | ולפי |  |
| 6:10.367 | שכל |  |
| 6:10.368 | ישראל |  |
| 6:10.369 | בזמן |  |
| 6:10.370 | רז"ל |  |
| 6:10.371 | אדוקים |  |
| 6:10.372 | ודבוקים |  |
| 6:10.373 | ביחוד |  |
| 6:10.374 | השי"ת |  |
| 6:10.375 | באמת |  |
| 6:10.376 | לא |  |
| 6:10.377 | הוצרכו |  |
| 6:10.378 | במשנה |  |
| 6:10.379 | ובתלמוד |  |
| 6:10.380 | לבאר |  |
| 6:10.381 | דבר |  |
| 6:10.382 | זה |  |
| 6:10.383 | ולומר |  |
| 6:10.384 | שמי |  |
| 6:10.385 | שחושב |  |
| 6:10.386 | עלות |  |
| 6:10.387 | נאצלים |  |
| 6:10.388 | מהבורא |  |
| 6:10.389 | ואומר |  |
| 6:10.390 | שהם |  |
| 6:10.391 | אחד |  |
| 6:10.392 | עם |  |
| 6:10.393 | מאצילם |  |
| 6:10.394 | שהוא |  |
| 6:10.395 | עובד |  |
| 6:10.396 | ע"ז | . |
| 6:10.397 | כי |  |
| 6:10.398 | זו |  |
| 6:10.399 | היא |  |
| 6:10.400 | אמונת |  |
| 6:10.401 | המשיחים | . |
| 6:10.402 | ועליה |  |
| 6:10.403 | אמר |  |
| 6:10.404 | ר | ' |
| 6:10.405 | שמעון |  |
| 6:10.406 | ב"י |  |
| 6:10.407 | כל |  |
| 6:10.408 | המשתף |  |
| 6:10.409 | שם |  |
| 6:10.410 | שמים |  |
| 6:10.411 | ודבר |  |
| 6:10.412 | אחר |  |
| 6:10.413 | נעקר |  |
| 6:10.414 | מן |  |
| 6:10.415 | העולם | . |
| 6:10.416 | וכבר |  |
| 6:10.417 | אמרו |  |
| 6:10.418 | הנצרים |  |
| 6:10.419 | עובדי |  |
| 6:10.420 | ע"ז |  |
| 6:10.421 | הם | . |
| 6:10.422 | ויום |  |
| 6:10.423 | ראשון |  |
| 6:10.424 | יום |  |
| 6:10.425 | אידם |  |
| 6:10.426 | הוא | . |
| 6:10.427 | לפיכך |  |
| 6:10.428 | אסור |  |
| 6:10.429 | לשאת |  |
| 6:10.430 | ולתת |  |
| 6:10.431 | עמהן |  |
| 6:10.432 | בארץ |  |
| 6:10.433 | ישראל |  |
| 6:10.434 | יום |  |
| 6:10.435 | ה | ' |
| 6:10.436 | ויום |  |
| 6:10.437 | ו | ' |
| 6:10.438 | שבכל |  |
| 6:10.439 | שבת |  |
| 6:10.440 | ושבת |  |
| 6:10.441 | ואין |  |
| 6:10.442 | צריך |  |
| 6:10.443 | לומר |  |
| 6:10.444 | יום |  |
| 6:10.445 | ראשון |  |
| 6:10.446 | עצמו |  |
| 6:10.447 | שהוא |  |
| 6:10.448 | אסור |  |
| 6:10.449 | בכל |  |
| 6:10.450 | מקום | . |
| 6:10.451 | וכן |  |
| 6:10.452 | בכל |  |
| 6:10.453 | אידיהן |  |
| 6:10.454 | הרמב"ם |  |
| 6:10.455 | פ"ט |  |
| 6:10.456 | מע"ז | ) |
| 6:10.457 | ולכן |  |
| 6:10.458 | כ | ' |
| 6:10.459 | הסמ"ג |  |
| 6:10.460 | שאין |  |
| 6:10.461 | צריך |  |
| 6:10.462 | להאריך |  |
| 6:10.463 | בזה |  |
| 6:10.464 | מפני |  |
| 6:10.465 | שכל |  |
| 6:10.466 | ישראל |  |
| 6:10.467 | אדוקים |  |
| 6:10.468 | באמונת |  |
| 6:10.469 | היחוד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

