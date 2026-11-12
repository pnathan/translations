# Translation pack: milhamot-hashem perek-8-04 (Milchamot ha-Shem 8:5–8:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-04`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
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
- **הישמעאלים / ישמעאלים** → "the Ishmaelites" (note: Muslims).
- **השונים** → "the Dualists".
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted"; "tzimtzum" in a
  note first in a unit.
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- הסבה **הראשונה** → "the First Cause".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- עבודה **אחרת** → "other worship" (note first in a unit).
- אל **אחר** → "another god".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- **השתלשלות** → "the chain of descent".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
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


## Milchamot ha-Shem 8:5

Hebrew (Wikisource): קב) ולפי דברי הרב סוד ישרים לא הרוחנו מאומה בידיעת ענין הצמצום עם השתלשלות העגולים והפרצופים של יושר זה אחר זה, רק שעברנו על דברי משיח אלהי יעקב שאמר מי כה' אלהינו המגביהי לשבת המשפילי לראות בשמים ובארץ, ואומר אמרו בלבבכם על משכבכם ודומו סלה, והיה די לנו במאמר הנביא לכו ונלכה באור ה' (ולפי האמת האור הנזכר בפסוק הזה הוא אור התורה הקדושה שנאמר כי נר מצוה ותורה אור שהיה מאירה לאדם ומישרת דרכיו לשמור מצותיו וחוקותיו ית', אבל האורות הרוחניים אינם מאירים ואין בכחם ליתן תורה אחרת, כי לא יחליף האל לא ימיר תורתינו באחרת) ולמה זה נבקש חשבונות רבים המביאים לידי מינות לעבוד אל אחר קצר אפים, והאלהים עשה את האדם בשכל ישר, ונתן לו תורת אמת, והראנו את כבודו ואת גדלו ואמר כי לא ראיתם כל תמונה! ומ"מ גם לפי הנחת סוד ישרים שהאור הנברא עם כמה אורות קדושים רוחנים למעלה מהלמאכים ונקראים בשם ספירות ואין להם שום צורה ודמות, מ"מ בכלל צבא השמים שברא אלהינו יחשבו, ואסור לעבדן כיון שהם נבראים, ואין ראוי לעבוד אלא לאלהי קדם שבראם, בין שהם גשמים או רוחנים, וגם לצרפם ולאחדם ולשתפם לבוראם אסור לעבוד בלתי לה' לבדו שהוא הסבה הראשונה כמו שקבלו רז"ל וכמו שנבאר לעיל שכפי האמת הגמור היא שוה לאמונת הישמעאלים שקבלוה ממחוקקם שקבל מישראל כידוע, וזכו בני ישמעאל לזאת האמונה הטהורה ביחוד השי"ת בשכר תשובה שעשה ישמעאל אביהם, כד קימא לך שישמעאל עשה תשובה והאמין בה' שהוא הסבה הראשונה, ונקרא בלשונם אללה, וזה זכרי לדר דר וכו', ובלשון רז"ל הקב"ה או רבש"ע, והוא אלהינו ואנחנו עמו, ולו אנחנו עובדים ומשתחוים ומברכים, כמקובל בידינו מאבות אבותינו ומרז"ל (ואיך יאמר בעל ספר הברית שאין אמונתינו באמונת הישמעאלים) ולחנם נכתבו בדברי המקובלים החדשים כל אותן הציורים וההגשמות והפרצופים ומעלותיהם השונים זה מזה, והוא ב"ה רק בחפץ רצונו ית' מראה מכבודו לנביאים ובמשכן ובבית עולמים כדי לקבוע אמונתו בלבינו ככתוב כי לבעבור גסות אתכם בא האלהים ובעבור תהיה יראתו על פניכם לבלתי תחטאו, ואין כאן דבר ידוע וקבוע באור הנזכר לומר בו זה פרצוף פלוני, וזה פרצוף לפוני אלמוני, וזה קצר, וזה ארוך, וזה מעלתו למעלה, וזה למטה, ועד היכן מגיעים רגליו כמו שחשבו המקובלים החדשים, אלא הכל כחפצו ית' הוא מראה לנו.

| id | word | punctuation after |
|---|---|---|
| 8:5.1 | קב | ) |
| 8:5.2 | ולפי |  |
| 8:5.3 | דברי |  |
| 8:5.4 | הרב |  |
| 8:5.5 | סוד |  |
| 8:5.6 | ישרים |  |
| 8:5.7 | לא |  |
| 8:5.8 | הרוחנו |  |
| 8:5.9 | מאומה |  |
| 8:5.10 | בידיעת |  |
| 8:5.11 | ענין |  |
| 8:5.12 | הצמצום |  |
| 8:5.13 | עם |  |
| 8:5.14 | השתלשלות |  |
| 8:5.15 | העגולים |  |
| 8:5.16 | והפרצופים |  |
| 8:5.17 | של |  |
| 8:5.18 | יושר |  |
| 8:5.19 | זה |  |
| 8:5.20 | אחר |  |
| 8:5.21 | זה | , |
| 8:5.22 | רק |  |
| 8:5.23 | שעברנו |  |
| 8:5.24 | על |  |
| 8:5.25 | דברי |  |
| 8:5.26 | משיח |  |
| 8:5.27 | אלהי |  |
| 8:5.28 | יעקב |  |
| 8:5.29 | שאמר |  |
| 8:5.30 | מי |  |
| 8:5.31 | כה | ' |
| 8:5.32 | אלהינו |  |
| 8:5.33 | המגביהי |  |
| 8:5.34 | לשבת |  |
| 8:5.35 | המשפילי |  |
| 8:5.36 | לראות |  |
| 8:5.37 | בשמים |  |
| 8:5.38 | ובארץ | , |
| 8:5.39 | ואומר |  |
| 8:5.40 | אמרו |  |
| 8:5.41 | בלבבכם |  |
| 8:5.42 | על |  |
| 8:5.43 | משכבכם |  |
| 8:5.44 | ודומו |  |
| 8:5.45 | סלה | , |
| 8:5.46 | והיה |  |
| 8:5.47 | די |  |
| 8:5.48 | לנו |  |
| 8:5.49 | במאמר |  |
| 8:5.50 | הנביא |  |
| 8:5.51 | לכו |  |
| 8:5.52 | ונלכה |  |
| 8:5.53 | באור |  |
| 8:5.54 | ה | ' |
| 8:5.55 | ולפי |  |
| 8:5.56 | האמת |  |
| 8:5.57 | האור |  |
| 8:5.58 | הנזכר |  |
| 8:5.59 | בפסוק |  |
| 8:5.60 | הזה |  |
| 8:5.61 | הוא |  |
| 8:5.62 | אור |  |
| 8:5.63 | התורה |  |
| 8:5.64 | הקדושה |  |
| 8:5.65 | שנאמר |  |
| 8:5.66 | כי |  |
| 8:5.67 | נר |  |
| 8:5.68 | מצוה |  |
| 8:5.69 | ותורה |  |
| 8:5.70 | אור |  |
| 8:5.71 | שהיה |  |
| 8:5.72 | מאירה |  |
| 8:5.73 | לאדם |  |
| 8:5.74 | ומישרת |  |
| 8:5.75 | דרכיו |  |
| 8:5.76 | לשמור |  |
| 8:5.77 | מצותיו |  |
| 8:5.78 | וחוקותיו |  |
| 8:5.79 | ית | ', |
| 8:5.80 | אבל |  |
| 8:5.81 | האורות |  |
| 8:5.82 | הרוחניים |  |
| 8:5.83 | אינם |  |
| 8:5.84 | מאירים |  |
| 8:5.85 | ואין |  |
| 8:5.86 | בכחם |  |
| 8:5.87 | ליתן |  |
| 8:5.88 | תורה |  |
| 8:5.89 | אחרת | , |
| 8:5.90 | כי |  |
| 8:5.91 | לא |  |
| 8:5.92 | יחליף |  |
| 8:5.93 | האל |  |
| 8:5.94 | לא |  |
| 8:5.95 | ימיר |  |
| 8:5.96 | תורתינו |  |
| 8:5.97 | באחרת | ) |
| 8:5.98 | ולמה |  |
| 8:5.99 | זה |  |
| 8:5.100 | נבקש |  |
| 8:5.101 | חשבונות |  |
| 8:5.102 | רבים |  |
| 8:5.103 | המביאים |  |
| 8:5.104 | לידי |  |
| 8:5.105 | מינות |  |
| 8:5.106 | לעבוד |  |
| 8:5.107 | אל |  |
| 8:5.108 | אחר |  |
| 8:5.109 | קצר |  |
| 8:5.110 | אפים | , |
| 8:5.111 | והאלהים |  |
| 8:5.112 | עשה |  |
| 8:5.113 | את |  |
| 8:5.114 | האדם |  |
| 8:5.115 | בשכל |  |
| 8:5.116 | ישר | , |
| 8:5.117 | ונתן |  |
| 8:5.118 | לו |  |
| 8:5.119 | תורת |  |
| 8:5.120 | אמת | , |
| 8:5.121 | והראנו |  |
| 8:5.122 | את |  |
| 8:5.123 | כבודו |  |
| 8:5.124 | ואת |  |
| 8:5.125 | גדלו |  |
| 8:5.126 | ואמר |  |
| 8:5.127 | כי |  |
| 8:5.128 | לא |  |
| 8:5.129 | ראיתם |  |
| 8:5.130 | כל |  |
| 8:5.131 | תמונה | ! |
| 8:5.132 | ומ"מ |  |
| 8:5.133 | גם |  |
| 8:5.134 | לפי |  |
| 8:5.135 | הנחת |  |
| 8:5.136 | סוד |  |
| 8:5.137 | ישרים |  |
| 8:5.138 | שהאור |  |
| 8:5.139 | הנברא |  |
| 8:5.140 | עם |  |
| 8:5.141 | כמה |  |
| 8:5.142 | אורות |  |
| 8:5.143 | קדושים |  |
| 8:5.144 | רוחנים |  |
| 8:5.145 | למעלה |  |
| 8:5.146 | מהלמאכים |  |
| 8:5.147 | ונקראים |  |
| 8:5.148 | בשם |  |
| 8:5.149 | ספירות |  |
| 8:5.150 | ואין |  |
| 8:5.151 | להם |  |
| 8:5.152 | שום |  |
| 8:5.153 | צורה |  |
| 8:5.154 | ודמות | , |
| 8:5.155 | מ"מ |  |
| 8:5.156 | בכלל |  |
| 8:5.157 | צבא |  |
| 8:5.158 | השמים |  |
| 8:5.159 | שברא |  |
| 8:5.160 | אלהינו |  |
| 8:5.161 | יחשבו | , |
| 8:5.162 | ואסור |  |
| 8:5.163 | לעבדן |  |
| 8:5.164 | כיון |  |
| 8:5.165 | שהם |  |
| 8:5.166 | נבראים | , |
| 8:5.167 | ואין |  |
| 8:5.168 | ראוי |  |
| 8:5.169 | לעבוד |  |
| 8:5.170 | אלא |  |
| 8:5.171 | לאלהי |  |
| 8:5.172 | קדם |  |
| 8:5.173 | שבראם | , |
| 8:5.174 | בין |  |
| 8:5.175 | שהם |  |
| 8:5.176 | גשמים |  |
| 8:5.177 | או |  |
| 8:5.178 | רוחנים | , |
| 8:5.179 | וגם |  |
| 8:5.180 | לצרפם |  |
| 8:5.181 | ולאחדם |  |
| 8:5.182 | ולשתפם |  |
| 8:5.183 | לבוראם |  |
| 8:5.184 | אסור |  |
| 8:5.185 | לעבוד |  |
| 8:5.186 | בלתי |  |
| 8:5.187 | לה | ' |
| 8:5.188 | לבדו |  |
| 8:5.189 | שהוא |  |
| 8:5.190 | הסבה |  |
| 8:5.191 | הראשונה |  |
| 8:5.192 | כמו |  |
| 8:5.193 | שקבלו |  |
| 8:5.194 | רז"ל |  |
| 8:5.195 | וכמו |  |
| 8:5.196 | שנבאר |  |
| 8:5.197 | לעיל |  |
| 8:5.198 | שכפי |  |
| 8:5.199 | האמת |  |
| 8:5.200 | הגמור |  |
| 8:5.201 | היא |  |
| 8:5.202 | שוה |  |
| 8:5.203 | לאמונת |  |
| 8:5.204 | הישמעאלים |  |
| 8:5.205 | שקבלוה |  |
| 8:5.206 | ממחוקקם |  |
| 8:5.207 | שקבל |  |
| 8:5.208 | מישראל |  |
| 8:5.209 | כידוע | , |
| 8:5.210 | וזכו |  |
| 8:5.211 | בני |  |
| 8:5.212 | ישמעאל |  |
| 8:5.213 | לזאת |  |
| 8:5.214 | האמונה |  |
| 8:5.215 | הטהורה |  |
| 8:5.216 | ביחוד |  |
| 8:5.217 | השי"ת |  |
| 8:5.218 | בשכר |  |
| 8:5.219 | תשובה |  |
| 8:5.220 | שעשה |  |
| 8:5.221 | ישמעאל |  |
| 8:5.222 | אביהם | , |
| 8:5.223 | כד |  |
| 8:5.224 | קימא |  |
| 8:5.225 | לך |  |
| 8:5.226 | שישמעאל |  |
| 8:5.227 | עשה |  |
| 8:5.228 | תשובה |  |
| 8:5.229 | והאמין |  |
| 8:5.230 | בה | ' |
| 8:5.231 | שהוא |  |
| 8:5.232 | הסבה |  |
| 8:5.233 | הראשונה | , |
| 8:5.234 | ונקרא |  |
| 8:5.235 | בלשונם |  |
| 8:5.236 | אללה | , |
| 8:5.237 | וזה |  |
| 8:5.238 | זכרי |  |
| 8:5.239 | לדר |  |
| 8:5.240 | דר |  |
| 8:5.241 | וכו | ', |
| 8:5.242 | ובלשון |  |
| 8:5.243 | רז"ל |  |
| 8:5.244 | הקב"ה |  |
| 8:5.245 | או |  |
| 8:5.246 | רבש"ע | , |
| 8:5.247 | והוא |  |
| 8:5.248 | אלהינו |  |
| 8:5.249 | ואנחנו |  |
| 8:5.250 | עמו | , |
| 8:5.251 | ולו |  |
| 8:5.252 | אנחנו |  |
| 8:5.253 | עובדים |  |
| 8:5.254 | ומשתחוים |  |
| 8:5.255 | ומברכים | , |
| 8:5.256 | כמקובל |  |
| 8:5.257 | בידינו |  |
| 8:5.258 | מאבות |  |
| 8:5.259 | אבותינו |  |
| 8:5.260 | ומרז"ל |  |
| 8:5.261 | ואיך |  |
| 8:5.262 | יאמר |  |
| 8:5.263 | בעל |  |
| 8:5.264 | ספר |  |
| 8:5.265 | הברית |  |
| 8:5.266 | שאין |  |
| 8:5.267 | אמונתינו |  |
| 8:5.268 | באמונת |  |
| 8:5.269 | הישמעאלים | ) |
| 8:5.270 | ולחנם |  |
| 8:5.271 | נכתבו |  |
| 8:5.272 | בדברי |  |
| 8:5.273 | המקובלים |  |
| 8:5.274 | החדשים |  |
| 8:5.275 | כל |  |
| 8:5.276 | אותן |  |
| 8:5.277 | הציורים |  |
| 8:5.278 | וההגשמות |  |
| 8:5.279 | והפרצופים |  |
| 8:5.280 | ומעלותיהם |  |
| 8:5.281 | השונים |  |
| 8:5.282 | זה |  |
| 8:5.283 | מזה | , |
| 8:5.284 | והוא |  |
| 8:5.285 | ב"ה |  |
| 8:5.286 | רק |  |
| 8:5.287 | בחפץ |  |
| 8:5.288 | רצונו |  |
| 8:5.289 | ית | ' |
| 8:5.290 | מראה |  |
| 8:5.291 | מכבודו |  |
| 8:5.292 | לנביאים |  |
| 8:5.293 | ובמשכן |  |
| 8:5.294 | ובבית |  |
| 8:5.295 | עולמים |  |
| 8:5.296 | כדי |  |
| 8:5.297 | לקבוע |  |
| 8:5.298 | אמונתו |  |
| 8:5.299 | בלבינו |  |
| 8:5.300 | ככתוב |  |
| 8:5.301 | כי |  |
| 8:5.302 | לבעבור |  |
| 8:5.303 | גסות |  |
| 8:5.304 | אתכם |  |
| 8:5.305 | בא |  |
| 8:5.306 | האלהים |  |
| 8:5.307 | ובעבור |  |
| 8:5.308 | תהיה |  |
| 8:5.309 | יראתו |  |
| 8:5.310 | על |  |
| 8:5.311 | פניכם |  |
| 8:5.312 | לבלתי |  |
| 8:5.313 | תחטאו | , |
| 8:5.314 | ואין |  |
| 8:5.315 | כאן |  |
| 8:5.316 | דבר |  |
| 8:5.317 | ידוע |  |
| 8:5.318 | וקבוע |  |
| 8:5.319 | באור |  |
| 8:5.320 | הנזכר |  |
| 8:5.321 | לומר |  |
| 8:5.322 | בו |  |
| 8:5.323 | זה |  |
| 8:5.324 | פרצוף |  |
| 8:5.325 | פלוני | , |
| 8:5.326 | וזה |  |
| 8:5.327 | פרצוף |  |
| 8:5.328 | לפוני |  |
| 8:5.329 | אלמוני | , |
| 8:5.330 | וזה |  |
| 8:5.331 | קצר | , |
| 8:5.332 | וזה |  |
| 8:5.333 | ארוך | , |
| 8:5.334 | וזה |  |
| 8:5.335 | מעלתו |  |
| 8:5.336 | למעלה | , |
| 8:5.337 | וזה |  |
| 8:5.338 | למטה | , |
| 8:5.339 | ועד |  |
| 8:5.340 | היכן |  |
| 8:5.341 | מגיעים |  |
| 8:5.342 | רגליו |  |
| 8:5.343 | כמו |  |
| 8:5.344 | שחשבו |  |
| 8:5.345 | המקובלים |  |
| 8:5.346 | החדשים | , |
| 8:5.347 | אלא |  |
| 8:5.348 | הכל |  |
| 8:5.349 | כחפצו |  |
| 8:5.350 | ית | ' |
| 8:5.351 | הוא |  |
| 8:5.352 | מראה |  |
| 8:5.353 | לנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

