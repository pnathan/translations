# Translation pack: milhamot-hashem perek-5-02 (Milchamot ha-Shem 5:2–5:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-02`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **רשב"א** → "Rashba" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
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

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **בחינה** → "aspect".
- **עלה / עלול** → "cause / effect".
- **כוונה** (in prayer) → "intention".
- לית מחשבה **תפיסה** ביה → "no thought grasps Him".
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
- **א"כ** → "if so".
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **רוע** לב → "wickedness of heart".
- **סכלות** והוללות → "folly and madness".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **תא חזי** → "come and see".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 5:2

Hebrew (Wikisource): מא) ועל התגדלך והתפארך בעצמך, והתרוממת על מרי שמיא לדרוש בהמון כי ידעתו שבתו וקומו של ממ"ה הקב"ה וליטול את השם ולאמר אבל קבלתנו וכו', הנה לא הודעתנו אל איזה א"ס אתה מתפלל, כי כמה מיני א"ס יש, וגם לא בארת ממי קבלת קבלה זו. ומעיקרא לא על קבלתך שאלך השואל, ולמה לך להשיב שלא כענין, היפך דברי הזהר והמקובלים כמ"ש לעיל, וכמ"ש בהדיא בכסא אליהו דף ט"ו ע"ב וז"ל: "אמנם ענין הברכה לו ית"ש אין הכוונה לעצמות אדון יחיד ב"ה חלילה חלילה, כי הוא מרומם מעל כל ברכה" ע"כ, וכ"כ ספר הברית במאמר שני כי הא"ס בבחינתו הפשוטה אשר דיברו הפילוסופים (הכופרים) שלא שייך בו שום עבודה ותפילה, על כי הוא בבחינה זאת נעלה ומרומם מכל ברכה ותהילה, ובטלה כל מצוה וכל תורה ועבודה, עי"ש ועוד יתבאר לפנינו. וכיון שקבלתך זו שהזכרת היא היפך דברי הזהר והמקובלים מקדש מלך וכסא אליהו ומצרף האמונה ויושר לבב ומשנת חסידים והרש"ב ומהרח"ו בעץ חיים וספר הכוונות ומחברת הקודש וספר הברית והאר"י בספר הליקוטים, א"כ הוי שוטה וסכל יקראו לך! על מה אתה נלחם בלומדי תורה ברכילות ומלשינות ולה"ר? ועל מה אתה מבזה אותם? העל שאינם מהגים ומצפצפים בספר הזהר כמוך? רק מחזיקים במשנה ותלמוד ופוסקים? הן על זה פירשו מללמוד בספרי הקבלה החדשה כי רק מיד בני נכר הגיעה אלינו כמבואר בספר שם הגדולים להר"ר חיד"א. ובגיליון הבא מירושלים בשם הרב הנזכר ובשם הרב יש"ר וכמו שיתבאר להלן. ולדעת רבים וגדולים מחכמינו הוי מינות. ולכן קבענו כל לימודנו במשנה ותלמוד ופוסקים, כהוראת רב האי גאון המובא בתשובת הרשב"א ובמבוא התלמוד שאין טוב לישראל כעסק המשנה והתלמוד, זה לנו יותר משלושים שנה. ובזדון לבכם השמעתם בקהל רב כי אנחנו מינים וכופרים על שאין אנו לומדים בזהר ותיקונים. הן עתה נתברר לנו מתשובותיכם כי לא ידעתם מספרי ישראל מאומה, לא במדרשי רז"ל ולא בתלמוד ולא בקבלה החדשה. כמאמר ר' יהודה בן תיבון בפתיחה לשער השני מספר חובת הלבבות בשם מבחר הפנינים: "מי שמתחכם מבלי חכמה הוא כחמור הריחיים, סובב סובב ואינו זז ממקומו". עוד שם: "אל תאמר במה שלא ידעת - ידעתי, פן תחשד במה שידעת, ואמרו החכמים: מי שסכל וחשב שידע סכלותו כפולה", עכ"ל. הוי רועים רעים המטעים את העם בשקריהם, ומחשיבים עצמם לחכמים בעיניהם! חכמים אתם להרע ולחלל את השי"ת ואת תורתו! ובדת אמת אשר לא ידעתם אותה על בריה תסחרו, להיות לכם לקרדום לחפור בה בורות נשברים אשר לא יכילו המים, כמאמר הנביא (מיכה ג' ה'): "כה אמר ה' על הנביאים המתעים את עמי, הנשכים בשניהם וקראו שלום. ואשר לא יתן אל פיהם, וקדשו עליו מלחמה". ומתשובתך ניכר שלא ידעת רק אין סוף אחד המקיף את כל המציאות. ולא ידעת שיש כמה מיני אין סוף בדברי המקובלים, ושאין לכוון בהתפללו לזעיר אנפין רק להמשיך מאין סוף הגנוז בכתר בלבד. אבל האין סוף אשר הוא לעילא לעילא לית מחשבה תפיסה ביה כלל ואין לכוון תפילתו אליו, כמ"ש לעיל בשם הרח"ו בספר ארבע מאות שקל כסף וספר הברית.

| id | word | punctuation after |
|---|---|---|
| 5:2.1 | מא | ) |
| 5:2.2 | ועל |  |
| 5:2.3 | התגדלך |  |
| 5:2.4 | והתפארך |  |
| 5:2.5 | בעצמך | , |
| 5:2.6 | והתרוממת |  |
| 5:2.7 | על |  |
| 5:2.8 | מרי |  |
| 5:2.9 | שמיא |  |
| 5:2.10 | לדרוש |  |
| 5:2.11 | בהמון |  |
| 5:2.12 | כי |  |
| 5:2.13 | ידעתו |  |
| 5:2.14 | שבתו |  |
| 5:2.15 | וקומו |  |
| 5:2.16 | של |  |
| 5:2.17 | ממ"ה |  |
| 5:2.18 | הקב"ה |  |
| 5:2.19 | וליטול |  |
| 5:2.20 | את |  |
| 5:2.21 | השם |  |
| 5:2.22 | ולאמר |  |
| 5:2.23 | אבל |  |
| 5:2.24 | קבלתנו |  |
| 5:2.25 | וכו | ', |
| 5:2.26 | הנה |  |
| 5:2.27 | לא |  |
| 5:2.28 | הודעתנו |  |
| 5:2.29 | אל |  |
| 5:2.30 | איזה |  |
| 5:2.31 | א"ס |  |
| 5:2.32 | אתה |  |
| 5:2.33 | מתפלל | , |
| 5:2.34 | כי |  |
| 5:2.35 | כמה |  |
| 5:2.36 | מיני |  |
| 5:2.37 | א"ס |  |
| 5:2.38 | יש | , |
| 5:2.39 | וגם |  |
| 5:2.40 | לא |  |
| 5:2.41 | בארת |  |
| 5:2.42 | ממי |  |
| 5:2.43 | קבלת |  |
| 5:2.44 | קבלה |  |
| 5:2.45 | זו | . |
| 5:2.46 | ומעיקרא |  |
| 5:2.47 | לא |  |
| 5:2.48 | על |  |
| 5:2.49 | קבלתך |  |
| 5:2.50 | שאלך |  |
| 5:2.51 | השואל | , |
| 5:2.52 | ולמה |  |
| 5:2.53 | לך |  |
| 5:2.54 | להשיב |  |
| 5:2.55 | שלא |  |
| 5:2.56 | כענין | , |
| 5:2.57 | היפך |  |
| 5:2.58 | דברי |  |
| 5:2.59 | הזהר |  |
| 5:2.60 | והמקובלים |  |
| 5:2.61 | כמ"ש |  |
| 5:2.62 | לעיל | , |
| 5:2.63 | וכמ"ש |  |
| 5:2.64 | בהדיא |  |
| 5:2.65 | בכסא |  |
| 5:2.66 | אליהו |  |
| 5:2.67 | דף |  |
| 5:2.68 | ט"ו |  |
| 5:2.69 | ע"ב |  |
| 5:2.70 | וז"ל | : |
| 5:2.71 | אמנם |  |
| 5:2.72 | ענין |  |
| 5:2.73 | הברכה |  |
| 5:2.74 | לו |  |
| 5:2.75 | ית"ש |  |
| 5:2.76 | אין |  |
| 5:2.77 | הכוונה |  |
| 5:2.78 | לעצמות |  |
| 5:2.79 | אדון |  |
| 5:2.80 | יחיד |  |
| 5:2.81 | ב"ה |  |
| 5:2.82 | חלילה |  |
| 5:2.83 | חלילה | , |
| 5:2.84 | כי |  |
| 5:2.85 | הוא |  |
| 5:2.86 | מרומם |  |
| 5:2.87 | מעל |  |
| 5:2.88 | כל |  |
| 5:2.89 | ברכה | " |
| 5:2.90 | ע"כ | , |
| 5:2.91 | וכ"כ |  |
| 5:2.92 | ספר |  |
| 5:2.93 | הברית |  |
| 5:2.94 | במאמר |  |
| 5:2.95 | שני |  |
| 5:2.96 | כי |  |
| 5:2.97 | הא"ס |  |
| 5:2.98 | בבחינתו |  |
| 5:2.99 | הפשוטה |  |
| 5:2.100 | אשר |  |
| 5:2.101 | דיברו |  |
| 5:2.102 | הפילוסופים |  |
| 5:2.103 | הכופרים | ) |
| 5:2.104 | שלא |  |
| 5:2.105 | שייך |  |
| 5:2.106 | בו |  |
| 5:2.107 | שום |  |
| 5:2.108 | עבודה |  |
| 5:2.109 | ותפילה | , |
| 5:2.110 | על |  |
| 5:2.111 | כי |  |
| 5:2.112 | הוא |  |
| 5:2.113 | בבחינה |  |
| 5:2.114 | זאת |  |
| 5:2.115 | נעלה |  |
| 5:2.116 | ומרומם |  |
| 5:2.117 | מכל |  |
| 5:2.118 | ברכה |  |
| 5:2.119 | ותהילה | , |
| 5:2.120 | ובטלה |  |
| 5:2.121 | כל |  |
| 5:2.122 | מצוה |  |
| 5:2.123 | וכל |  |
| 5:2.124 | תורה |  |
| 5:2.125 | ועבודה | , |
| 5:2.126 | עי"ש |  |
| 5:2.127 | ועוד |  |
| 5:2.128 | יתבאר |  |
| 5:2.129 | לפנינו | . |
| 5:2.130 | וכיון |  |
| 5:2.131 | שקבלתך |  |
| 5:2.132 | זו |  |
| 5:2.133 | שהזכרת |  |
| 5:2.134 | היא |  |
| 5:2.135 | היפך |  |
| 5:2.136 | דברי |  |
| 5:2.137 | הזהר |  |
| 5:2.138 | והמקובלים |  |
| 5:2.139 | מקדש |  |
| 5:2.140 | מלך |  |
| 5:2.141 | וכסא |  |
| 5:2.142 | אליהו |  |
| 5:2.143 | ומצרף |  |
| 5:2.144 | האמונה |  |
| 5:2.145 | ויושר |  |
| 5:2.146 | לבב |  |
| 5:2.147 | ומשנת |  |
| 5:2.148 | חסידים |  |
| 5:2.149 | והרש"ב |  |
| 5:2.150 | ומהרח"ו |  |
| 5:2.151 | בעץ |  |
| 5:2.152 | חיים |  |
| 5:2.153 | וספר |  |
| 5:2.154 | הכוונות |  |
| 5:2.155 | ומחברת |  |
| 5:2.156 | הקודש |  |
| 5:2.157 | וספר |  |
| 5:2.158 | הברית |  |
| 5:2.159 | והאר"י |  |
| 5:2.160 | בספר |  |
| 5:2.161 | הליקוטים | , |
| 5:2.162 | א"כ |  |
| 5:2.163 | הוי |  |
| 5:2.164 | שוטה |  |
| 5:2.165 | וסכל |  |
| 5:2.166 | יקראו |  |
| 5:2.167 | לך | ! |
| 5:2.168 | על |  |
| 5:2.169 | מה |  |
| 5:2.170 | אתה |  |
| 5:2.171 | נלחם |  |
| 5:2.172 | בלומדי |  |
| 5:2.173 | תורה |  |
| 5:2.174 | ברכילות |  |
| 5:2.175 | ומלשינות |  |
| 5:2.176 | ולה"ר | ? |
| 5:2.177 | ועל |  |
| 5:2.178 | מה |  |
| 5:2.179 | אתה |  |
| 5:2.180 | מבזה |  |
| 5:2.181 | אותם | ? |
| 5:2.182 | העל |  |
| 5:2.183 | שאינם |  |
| 5:2.184 | מהגים |  |
| 5:2.185 | ומצפצפים |  |
| 5:2.186 | בספר |  |
| 5:2.187 | הזהר |  |
| 5:2.188 | כמוך | ? |
| 5:2.189 | רק |  |
| 5:2.190 | מחזיקים |  |
| 5:2.191 | במשנה |  |
| 5:2.192 | ותלמוד |  |
| 5:2.193 | ופוסקים | ? |
| 5:2.194 | הן |  |
| 5:2.195 | על |  |
| 5:2.196 | זה |  |
| 5:2.197 | פירשו |  |
| 5:2.198 | מללמוד |  |
| 5:2.199 | בספרי |  |
| 5:2.200 | הקבלה |  |
| 5:2.201 | החדשה |  |
| 5:2.202 | כי |  |
| 5:2.203 | רק |  |
| 5:2.204 | מיד |  |
| 5:2.205 | בני |  |
| 5:2.206 | נכר |  |
| 5:2.207 | הגיעה |  |
| 5:2.208 | אלינו |  |
| 5:2.209 | כמבואר |  |
| 5:2.210 | בספר |  |
| 5:2.211 | שם |  |
| 5:2.212 | הגדולים |  |
| 5:2.213 | להר"ר |  |
| 5:2.214 | חיד"א | . |
| 5:2.215 | ובגיליון |  |
| 5:2.216 | הבא |  |
| 5:2.217 | מירושלים |  |
| 5:2.218 | בשם |  |
| 5:2.219 | הרב |  |
| 5:2.220 | הנזכר |  |
| 5:2.221 | ובשם |  |
| 5:2.222 | הרב |  |
| 5:2.223 | יש"ר |  |
| 5:2.224 | וכמו |  |
| 5:2.225 | שיתבאר |  |
| 5:2.226 | להלן | . |
| 5:2.227 | ולדעת |  |
| 5:2.228 | רבים |  |
| 5:2.229 | וגדולים |  |
| 5:2.230 | מחכמינו |  |
| 5:2.231 | הוי |  |
| 5:2.232 | מינות | . |
| 5:2.233 | ולכן |  |
| 5:2.234 | קבענו |  |
| 5:2.235 | כל |  |
| 5:2.236 | לימודנו |  |
| 5:2.237 | במשנה |  |
| 5:2.238 | ותלמוד |  |
| 5:2.239 | ופוסקים | , |
| 5:2.240 | כהוראת |  |
| 5:2.241 | רב |  |
| 5:2.242 | האי |  |
| 5:2.243 | גאון |  |
| 5:2.244 | המובא |  |
| 5:2.245 | בתשובת |  |
| 5:2.246 | הרשב"א |  |
| 5:2.247 | ובמבוא |  |
| 5:2.248 | התלמוד |  |
| 5:2.249 | שאין |  |
| 5:2.250 | טוב |  |
| 5:2.251 | לישראל |  |
| 5:2.252 | כעסק |  |
| 5:2.253 | המשנה |  |
| 5:2.254 | והתלמוד | , |
| 5:2.255 | זה |  |
| 5:2.256 | לנו |  |
| 5:2.257 | יותר |  |
| 5:2.258 | משלושים |  |
| 5:2.259 | שנה | . |
| 5:2.260 | ובזדון |  |
| 5:2.261 | לבכם |  |
| 5:2.262 | השמעתם |  |
| 5:2.263 | בקהל |  |
| 5:2.264 | רב |  |
| 5:2.265 | כי |  |
| 5:2.266 | אנחנו |  |
| 5:2.267 | מינים |  |
| 5:2.268 | וכופרים |  |
| 5:2.269 | על |  |
| 5:2.270 | שאין |  |
| 5:2.271 | אנו |  |
| 5:2.272 | לומדים |  |
| 5:2.273 | בזהר |  |
| 5:2.274 | ותיקונים | . |
| 5:2.275 | הן |  |
| 5:2.276 | עתה |  |
| 5:2.277 | נתברר |  |
| 5:2.278 | לנו |  |
| 5:2.279 | מתשובותיכם |  |
| 5:2.280 | כי |  |
| 5:2.281 | לא |  |
| 5:2.282 | ידעתם |  |
| 5:2.283 | מספרי |  |
| 5:2.284 | ישראל |  |
| 5:2.285 | מאומה | , |
| 5:2.286 | לא |  |
| 5:2.287 | במדרשי |  |
| 5:2.288 | רז"ל |  |
| 5:2.289 | ולא |  |
| 5:2.290 | בתלמוד |  |
| 5:2.291 | ולא |  |
| 5:2.292 | בקבלה |  |
| 5:2.293 | החדשה | . |
| 5:2.294 | כמאמר |  |
| 5:2.295 | ר | ' |
| 5:2.296 | יהודה |  |
| 5:2.297 | בן |  |
| 5:2.298 | תיבון |  |
| 5:2.299 | בפתיחה |  |
| 5:2.300 | לשער |  |
| 5:2.301 | השני |  |
| 5:2.302 | מספר |  |
| 5:2.303 | חובת |  |
| 5:2.304 | הלבבות |  |
| 5:2.305 | בשם |  |
| 5:2.306 | מבחר |  |
| 5:2.307 | הפנינים | : |
| 5:2.308 | מי |  |
| 5:2.309 | שמתחכם |  |
| 5:2.310 | מבלי |  |
| 5:2.311 | חכמה |  |
| 5:2.312 | הוא |  |
| 5:2.313 | כחמור |  |
| 5:2.314 | הריחיים | , |
| 5:2.315 | סובב |  |
| 5:2.316 | סובב |  |
| 5:2.317 | ואינו |  |
| 5:2.318 | זז |  |
| 5:2.319 | ממקומו | ". |
| 5:2.320 | עוד |  |
| 5:2.321 | שם | : |
| 5:2.322 | אל |  |
| 5:2.323 | תאמר |  |
| 5:2.324 | במה |  |
| 5:2.325 | שלא |  |
| 5:2.326 | ידעת | - |
| 5:2.327 | ידעתי | , |
| 5:2.328 | פן |  |
| 5:2.329 | תחשד |  |
| 5:2.330 | במה |  |
| 5:2.331 | שידעת | , |
| 5:2.332 | ואמרו |  |
| 5:2.333 | החכמים | : |
| 5:2.334 | מי |  |
| 5:2.335 | שסכל |  |
| 5:2.336 | וחשב |  |
| 5:2.337 | שידע |  |
| 5:2.338 | סכלותו |  |
| 5:2.339 | כפולה | ", |
| 5:2.340 | עכ"ל | . |
| 5:2.341 | הוי |  |
| 5:2.342 | רועים |  |
| 5:2.343 | רעים |  |
| 5:2.344 | המטעים |  |
| 5:2.345 | את |  |
| 5:2.346 | העם |  |
| 5:2.347 | בשקריהם | , |
| 5:2.348 | ומחשיבים |  |
| 5:2.349 | עצמם |  |
| 5:2.350 | לחכמים |  |
| 5:2.351 | בעיניהם | ! |
| 5:2.352 | חכמים |  |
| 5:2.353 | אתם |  |
| 5:2.354 | להרע |  |
| 5:2.355 | ולחלל |  |
| 5:2.356 | את |  |
| 5:2.357 | השי"ת |  |
| 5:2.358 | ואת |  |
| 5:2.359 | תורתו | ! |
| 5:2.360 | ובדת |  |
| 5:2.361 | אמת |  |
| 5:2.362 | אשר |  |
| 5:2.363 | לא |  |
| 5:2.364 | ידעתם |  |
| 5:2.365 | אותה |  |
| 5:2.366 | על |  |
| 5:2.367 | בריה |  |
| 5:2.368 | תסחרו | , |
| 5:2.369 | להיות |  |
| 5:2.370 | לכם |  |
| 5:2.371 | לקרדום |  |
| 5:2.372 | לחפור |  |
| 5:2.373 | בה |  |
| 5:2.374 | בורות |  |
| 5:2.375 | נשברים |  |
| 5:2.376 | אשר |  |
| 5:2.377 | לא |  |
| 5:2.378 | יכילו |  |
| 5:2.379 | המים | , |
| 5:2.380 | כמאמר |  |
| 5:2.381 | הנביא |  |
| 5:2.382 | מיכה |  |
| 5:2.383 | ג | ' |
| 5:2.384 | ה | '): |
| 5:2.385 | כה |  |
| 5:2.386 | אמר |  |
| 5:2.387 | ה | ' |
| 5:2.388 | על |  |
| 5:2.389 | הנביאים |  |
| 5:2.390 | המתעים |  |
| 5:2.391 | את |  |
| 5:2.392 | עמי | , |
| 5:2.393 | הנשכים |  |
| 5:2.394 | בשניהם |  |
| 5:2.395 | וקראו |  |
| 5:2.396 | שלום | . |
| 5:2.397 | ואשר |  |
| 5:2.398 | לא |  |
| 5:2.399 | יתן |  |
| 5:2.400 | אל |  |
| 5:2.401 | פיהם | , |
| 5:2.402 | וקדשו |  |
| 5:2.403 | עליו |  |
| 5:2.404 | מלחמה | ". |
| 5:2.405 | ומתשובתך |  |
| 5:2.406 | ניכר |  |
| 5:2.407 | שלא |  |
| 5:2.408 | ידעת |  |
| 5:2.409 | רק |  |
| 5:2.410 | אין |  |
| 5:2.411 | סוף |  |
| 5:2.412 | אחד |  |
| 5:2.413 | המקיף |  |
| 5:2.414 | את |  |
| 5:2.415 | כל |  |
| 5:2.416 | המציאות | . |
| 5:2.417 | ולא |  |
| 5:2.418 | ידעת |  |
| 5:2.419 | שיש |  |
| 5:2.420 | כמה |  |
| 5:2.421 | מיני |  |
| 5:2.422 | אין |  |
| 5:2.423 | סוף |  |
| 5:2.424 | בדברי |  |
| 5:2.425 | המקובלים | , |
| 5:2.426 | ושאין |  |
| 5:2.427 | לכוון |  |
| 5:2.428 | בהתפללו |  |
| 5:2.429 | לזעיר |  |
| 5:2.430 | אנפין |  |
| 5:2.431 | רק |  |
| 5:2.432 | להמשיך |  |
| 5:2.433 | מאין |  |
| 5:2.434 | סוף |  |
| 5:2.435 | הגנוז |  |
| 5:2.436 | בכתר |  |
| 5:2.437 | בלבד | . |
| 5:2.438 | אבל |  |
| 5:2.439 | האין |  |
| 5:2.440 | סוף |  |
| 5:2.441 | אשר |  |
| 5:2.442 | הוא |  |
| 5:2.443 | לעילא |  |
| 5:2.444 | לעילא |  |
| 5:2.445 | לית |  |
| 5:2.446 | מחשבה |  |
| 5:2.447 | תפיסה |  |
| 5:2.448 | ביה |  |
| 5:2.449 | כלל |  |
| 5:2.450 | ואין |  |
| 5:2.451 | לכוון |  |
| 5:2.452 | תפילתו |  |
| 5:2.453 | אליו | , |
| 5:2.454 | כמ"ש |  |
| 5:2.455 | לעיל |  |
| 5:2.456 | בשם |  |
| 5:2.457 | הרח"ו |  |
| 5:2.458 | בספר |  |
| 5:2.459 | ארבע |  |
| 5:2.460 | מאות |  |
| 5:2.461 | שקל |  |
| 5:2.462 | כסף |  |
| 5:2.463 | וספר |  |
| 5:2.464 | הברית | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

