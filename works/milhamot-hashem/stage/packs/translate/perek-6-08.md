# Translation pack: milhamot-hashem perek-6-08 (Milchamot ha-Shem 6:13–6:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-08`.

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
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **רמ"א** → "Rema" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
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
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אורות **דקים** → "subtle lights".
- כופר **בעיקר** → "denies the fundamental principle".
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **דרך ארץ** → "worldly occupation".
- **קורין** אל → "call upon".
- אחד ... **ואחד** of two kinds of case → "alike".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ידידי** → "my friend" (second person kept).
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- תהו **ובהו** → "formlessness and void".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 6:13

Hebrew (Wikisource): סו) הן עתה אערוך לפניך ידידי הקורא, השם נגד פניו כבוד השי"ת וכבוד תורה שבעל פה שהיא המשנה והתלמוד אשר כל דורש את דבר ה' בתורת אלהים יפנה ואעתיק פה קצת מאמרים מרז"ל להודיעך איך הם מחבבים ומוקירים את המשנה ואת התלמוד. וכי הם עיקר תורה שבעל פה שקבל משה מסיני, ולא שינה מכל אשר צוהו ה' אלהינו כאשר העיד עליו השי"ת באמרו בכל ביתי נאמן הוא. ואחרי כן אערוך לעיניך הקורא את דברי נביא השקר והפילסוף הרמאי מחבר ס' הזהר איך הוא מבזה את המשנה והתלמוד והוציא לעז וחשד למשה רבי' ע"ה שמפי עצמו אמרה, ושעל זה נקנסה עליו מיתה ולהקבר בארץ נכריה חוץ לארץ, ומן המפורסם וידוע לכל שלפי רוב חבתה של משנה בעיני רבותינו ז"ל בעלי התלמוד הם מדקדקים בלשונה כדרך שעשו במקראי קדש כדגרסינן בעירובין וסוכה מאי שנא מבוי דתני תקנתא ומאי שנא סוכה דתני פסולה וכו'. לפי שרז"ל דבר ברור הוא אצלם שהמשנה היא בעצמה ההלכות שקבל משה בסיני ועליה עמד בהר ארבעים וארבעים לילה ללמוד כללותיה ופרטותיה אשר בהם נזכה לחיי העוה"ב כמקובל בכל האומה הישראלית וכמו שאמרו רז"ל כל השונה הלכות בכל יום מובטח לו שהוא בן העולם הבא. ואמרו עוד שאין לו להקב"ה בעולמו אלא ארבע אמות של הלכה. ור' אמי ור' אסי לא הוו מצלו בבי כנשתא אלא ביני עמודי היכי דגרסי. ובהוריות דף י"ג במעשה דר' מאיר ור' נתן שנתקנאו ברבן שמעון בן גמליאל שתיקן שכשיכנס הנשיא כל העם עומדים מפניו, וכשאב ב"ד נכנס עושין לו שתי שורות אחת מכאן ואחת מכאן, וכשחכם נכנס אחד ועומד ואחד יושב עד שישב במקומו, ובקשו להעביר את רשב"ג מנשיאותו על ידי שישאלו ממנו לשנות להם מסכת עוקצין דלא הוה גמיר לה, וכי לא ידע נימרו ליה מי ימלל גבורות ה' ישמיע כל תהלתו, למי נאה למלל גבורות ה', למי שיכול להשמיע כל תהלתו וכו', אלמא דקרו רבנן למשנה תהלות ה', ובמדרש תנחומא פרשת וירא גרסינן אמר רב יהודה בן שלום בקש משה שתהא המשנה אף היא בכתב וצפה הקב"ה שאומות העולם עתידן לתרגם את התורה ולהיות קורין אותה יונית והן אומרין אנו הן של ישראל. אמר לו הקב"ה למשה אכתב לו רובי תורתי ואם כן כמו זר נחשבו, וכל כך למה מפני שהמשנה היא מסתורין של הקב"ה. ואין הקב"ה מוסר מסתורין שלו אלא לצדיקים שנ' סוד ה' ליריאיו. ועיין מה שכ' בזה מהר"י מדובנא בספרו הנחמד אהל יעקב בחתימת ספר במדבר וינעם לך. ראה כמה היה חביבה עליהם המשנה וקראו לה מסתורין של הקב"ה שאינו מוסרה אלא לצדיקים שנ' סוד ה' ליריאיו. וכן במדרש שמות רבה פרשה מ"ז אמרו כתב לך את הדברים האלה וגו' הדא הוא דכתיב אכתב לו רובי תורתי כמו זר נחשבו בשעה שנגלה הקב"ה בסיני ליתן תורה לישראל אמרה למשה על הסדר מקרא, משנה, תלמוד, ואגדה, שנאמר וידבר אלהים את כל הדברים האלה, אפי' מה שתלמיד שואל לרב אמר הקב"ה למשה באותה שעה ומאחר שלמד אותה מפי הקב"ה אמר לו למדה לישראל, אמר לפניו רבונו של עולם אכתוב אותה להם, אמר איני מבקש ליתנה להם בכתב, מפני שגלוי לפני שאומות העולם עתידים לשלוט בהם וליטול אותה מהם ויהיו בזוים באומות (פי' שיאמרו אומות העולם שניתנה להם) אלא המקרא אני נותן להם בכתב והמשנה והתלמוד והאגדה בעל פה וכו' ומסיים כי על פי הדברים האלה זו המשנה והתלמוד שהם מבדילים בין ישראל לבין האומות ע"כ. וכן כתבו קדמונינו הר' חטר ור' זכריה הרופא בפי' הי"ג עיקרים בעיקר השמיני וז"ל ואנמי' כאן על פה לפי שידע הקב"ה והודיע למשה שאומות העולם עתידין ליקח התורה על ישראל. וכיון שלקחו תורה שבכתב לא יוכלו ליקח תורה שבעל פה מפני שהיא בסופרים ולא בספרים. ואלו לקחו תורה שבכתב עם תורה שבעל פה היו אומרים שהם ישראל בעדות נאמנה ע"כ. ראה ידידי הקורא והבן דברי רז"ל שאמרו כי על המשנה והתלמוד שהם יסוד כל תורה שבעל פה כרת השי"ת ברית עם ישראל. כי המה חיי רוחנו לעולם הבא והם רובה של תורה כדאיתא בגיטין פרק הנזקין. ובפרק קמא דברכות גרסינן ואמר ר' לוי בר חמא אמר ריש לקיש מאי דכתיב ואתנה לך את לוחות האבן והתורה והמצוה אשר כתבתי להורותם, לוחות אלו עשרת הדברות, והתורה זו מקרא. והמצוה זו משנה, אשר כתבתי אלו נביאים וכתובים, להורותם זו גמרא. מלמד שכולם ניתנו למשה מסיני ע"כ. ומביאו רבינו הרמב"ם בהקדמת ספר היד עיין שם.

| id | word | punctuation after |
|---|---|---|
| 6:13.1 | סו | ) |
| 6:13.2 | הן |  |
| 6:13.3 | עתה |  |
| 6:13.4 | אערוך |  |
| 6:13.5 | לפניך |  |
| 6:13.6 | ידידי |  |
| 6:13.7 | הקורא | , |
| 6:13.8 | השם |  |
| 6:13.9 | נגד |  |
| 6:13.10 | פניו |  |
| 6:13.11 | כבוד |  |
| 6:13.12 | השי"ת |  |
| 6:13.13 | וכבוד |  |
| 6:13.14 | תורה |  |
| 6:13.15 | שבעל |  |
| 6:13.16 | פה |  |
| 6:13.17 | שהיא |  |
| 6:13.18 | המשנה |  |
| 6:13.19 | והתלמוד |  |
| 6:13.20 | אשר |  |
| 6:13.21 | כל |  |
| 6:13.22 | דורש |  |
| 6:13.23 | את |  |
| 6:13.24 | דבר |  |
| 6:13.25 | ה | ' |
| 6:13.26 | בתורת |  |
| 6:13.27 | אלהים |  |
| 6:13.28 | יפנה |  |
| 6:13.29 | ואעתיק |  |
| 6:13.30 | פה |  |
| 6:13.31 | קצת |  |
| 6:13.32 | מאמרים |  |
| 6:13.33 | מרז"ל |  |
| 6:13.34 | להודיעך |  |
| 6:13.35 | איך |  |
| 6:13.36 | הם |  |
| 6:13.37 | מחבבים |  |
| 6:13.38 | ומוקירים |  |
| 6:13.39 | את |  |
| 6:13.40 | המשנה |  |
| 6:13.41 | ואת |  |
| 6:13.42 | התלמוד | . |
| 6:13.43 | וכי |  |
| 6:13.44 | הם |  |
| 6:13.45 | עיקר |  |
| 6:13.46 | תורה |  |
| 6:13.47 | שבעל |  |
| 6:13.48 | פה |  |
| 6:13.49 | שקבל |  |
| 6:13.50 | משה |  |
| 6:13.51 | מסיני | , |
| 6:13.52 | ולא |  |
| 6:13.53 | שינה |  |
| 6:13.54 | מכל |  |
| 6:13.55 | אשר |  |
| 6:13.56 | צוהו |  |
| 6:13.57 | ה | ' |
| 6:13.58 | אלהינו |  |
| 6:13.59 | כאשר |  |
| 6:13.60 | העיד |  |
| 6:13.61 | עליו |  |
| 6:13.62 | השי"ת |  |
| 6:13.63 | באמרו |  |
| 6:13.64 | בכל |  |
| 6:13.65 | ביתי |  |
| 6:13.66 | נאמן |  |
| 6:13.67 | הוא | . |
| 6:13.68 | ואחרי |  |
| 6:13.69 | כן |  |
| 6:13.70 | אערוך |  |
| 6:13.71 | לעיניך |  |
| 6:13.72 | הקורא |  |
| 6:13.73 | את |  |
| 6:13.74 | דברי |  |
| 6:13.75 | נביא |  |
| 6:13.76 | השקר |  |
| 6:13.77 | והפילסוף |  |
| 6:13.78 | הרמאי |  |
| 6:13.79 | מחבר |  |
| 6:13.80 | ס | ' |
| 6:13.81 | הזהר |  |
| 6:13.82 | איך |  |
| 6:13.83 | הוא |  |
| 6:13.84 | מבזה |  |
| 6:13.85 | את |  |
| 6:13.86 | המשנה |  |
| 6:13.87 | והתלמוד |  |
| 6:13.88 | והוציא |  |
| 6:13.89 | לעז |  |
| 6:13.90 | וחשד |  |
| 6:13.91 | למשה |  |
| 6:13.92 | רבי | ' |
| 6:13.93 | ע"ה |  |
| 6:13.94 | שמפי |  |
| 6:13.95 | עצמו |  |
| 6:13.96 | אמרה | , |
| 6:13.97 | ושעל |  |
| 6:13.98 | זה |  |
| 6:13.99 | נקנסה |  |
| 6:13.100 | עליו |  |
| 6:13.101 | מיתה |  |
| 6:13.102 | ולהקבר |  |
| 6:13.103 | בארץ |  |
| 6:13.104 | נכריה |  |
| 6:13.105 | חוץ |  |
| 6:13.106 | לארץ | , |
| 6:13.107 | ומן |  |
| 6:13.108 | המפורסם |  |
| 6:13.109 | וידוע |  |
| 6:13.110 | לכל |  |
| 6:13.111 | שלפי |  |
| 6:13.112 | רוב |  |
| 6:13.113 | חבתה |  |
| 6:13.114 | של |  |
| 6:13.115 | משנה |  |
| 6:13.116 | בעיני |  |
| 6:13.117 | רבותינו |  |
| 6:13.118 | ז"ל |  |
| 6:13.119 | בעלי |  |
| 6:13.120 | התלמוד |  |
| 6:13.121 | הם |  |
| 6:13.122 | מדקדקים |  |
| 6:13.123 | בלשונה |  |
| 6:13.124 | כדרך |  |
| 6:13.125 | שעשו |  |
| 6:13.126 | במקראי |  |
| 6:13.127 | קדש |  |
| 6:13.128 | כדגרסינן |  |
| 6:13.129 | בעירובין |  |
| 6:13.130 | וסוכה |  |
| 6:13.131 | מאי |  |
| 6:13.132 | שנא |  |
| 6:13.133 | מבוי |  |
| 6:13.134 | דתני |  |
| 6:13.135 | תקנתא |  |
| 6:13.136 | ומאי |  |
| 6:13.137 | שנא |  |
| 6:13.138 | סוכה |  |
| 6:13.139 | דתני |  |
| 6:13.140 | פסולה |  |
| 6:13.141 | וכו | '. |
| 6:13.142 | לפי |  |
| 6:13.143 | שרז"ל |  |
| 6:13.144 | דבר |  |
| 6:13.145 | ברור |  |
| 6:13.146 | הוא |  |
| 6:13.147 | אצלם |  |
| 6:13.148 | שהמשנה |  |
| 6:13.149 | היא |  |
| 6:13.150 | בעצמה |  |
| 6:13.151 | ההלכות |  |
| 6:13.152 | שקבל |  |
| 6:13.153 | משה |  |
| 6:13.154 | בסיני |  |
| 6:13.155 | ועליה |  |
| 6:13.156 | עמד |  |
| 6:13.157 | בהר |  |
| 6:13.158 | ארבעים |  |
| 6:13.159 | וארבעים |  |
| 6:13.160 | לילה |  |
| 6:13.161 | ללמוד |  |
| 6:13.162 | כללותיה |  |
| 6:13.163 | ופרטותיה |  |
| 6:13.164 | אשר |  |
| 6:13.165 | בהם |  |
| 6:13.166 | נזכה |  |
| 6:13.167 | לחיי |  |
| 6:13.168 | העוה"ב |  |
| 6:13.169 | כמקובל |  |
| 6:13.170 | בכל |  |
| 6:13.171 | האומה |  |
| 6:13.172 | הישראלית |  |
| 6:13.173 | וכמו |  |
| 6:13.174 | שאמרו |  |
| 6:13.175 | רז"ל |  |
| 6:13.176 | כל |  |
| 6:13.177 | השונה |  |
| 6:13.178 | הלכות |  |
| 6:13.179 | בכל |  |
| 6:13.180 | יום |  |
| 6:13.181 | מובטח |  |
| 6:13.182 | לו |  |
| 6:13.183 | שהוא |  |
| 6:13.184 | בן |  |
| 6:13.185 | העולם |  |
| 6:13.186 | הבא | . |
| 6:13.187 | ואמרו |  |
| 6:13.188 | עוד |  |
| 6:13.189 | שאין |  |
| 6:13.190 | לו |  |
| 6:13.191 | להקב"ה |  |
| 6:13.192 | בעולמו |  |
| 6:13.193 | אלא |  |
| 6:13.194 | ארבע |  |
| 6:13.195 | אמות |  |
| 6:13.196 | של |  |
| 6:13.197 | הלכה | . |
| 6:13.198 | ור | ' |
| 6:13.199 | אמי |  |
| 6:13.200 | ור | ' |
| 6:13.201 | אסי |  |
| 6:13.202 | לא |  |
| 6:13.203 | הוו |  |
| 6:13.204 | מצלו |  |
| 6:13.205 | בבי |  |
| 6:13.206 | כנשתא |  |
| 6:13.207 | אלא |  |
| 6:13.208 | ביני |  |
| 6:13.209 | עמודי |  |
| 6:13.210 | היכי |  |
| 6:13.211 | דגרסי | . |
| 6:13.212 | ובהוריות |  |
| 6:13.213 | דף |  |
| 6:13.214 | י"ג |  |
| 6:13.215 | במעשה |  |
| 6:13.216 | דר | ' |
| 6:13.217 | מאיר |  |
| 6:13.218 | ור | ' |
| 6:13.219 | נתן |  |
| 6:13.220 | שנתקנאו |  |
| 6:13.221 | ברבן |  |
| 6:13.222 | שמעון |  |
| 6:13.223 | בן |  |
| 6:13.224 | גמליאל |  |
| 6:13.225 | שתיקן |  |
| 6:13.226 | שכשיכנס |  |
| 6:13.227 | הנשיא |  |
| 6:13.228 | כל |  |
| 6:13.229 | העם |  |
| 6:13.230 | עומדים |  |
| 6:13.231 | מפניו | , |
| 6:13.232 | וכשאב |  |
| 6:13.233 | ב"ד |  |
| 6:13.234 | נכנס |  |
| 6:13.235 | עושין |  |
| 6:13.236 | לו |  |
| 6:13.237 | שתי |  |
| 6:13.238 | שורות |  |
| 6:13.239 | אחת |  |
| 6:13.240 | מכאן |  |
| 6:13.241 | ואחת |  |
| 6:13.242 | מכאן | , |
| 6:13.243 | וכשחכם |  |
| 6:13.244 | נכנס |  |
| 6:13.245 | אחד |  |
| 6:13.246 | ועומד |  |
| 6:13.247 | ואחד |  |
| 6:13.248 | יושב |  |
| 6:13.249 | עד |  |
| 6:13.250 | שישב |  |
| 6:13.251 | במקומו | , |
| 6:13.252 | ובקשו |  |
| 6:13.253 | להעביר |  |
| 6:13.254 | את |  |
| 6:13.255 | רשב"ג |  |
| 6:13.256 | מנשיאותו |  |
| 6:13.257 | על |  |
| 6:13.258 | ידי |  |
| 6:13.259 | שישאלו |  |
| 6:13.260 | ממנו |  |
| 6:13.261 | לשנות |  |
| 6:13.262 | להם |  |
| 6:13.263 | מסכת |  |
| 6:13.264 | עוקצין |  |
| 6:13.265 | דלא |  |
| 6:13.266 | הוה |  |
| 6:13.267 | גמיר |  |
| 6:13.268 | לה | , |
| 6:13.269 | וכי |  |
| 6:13.270 | לא |  |
| 6:13.271 | ידע |  |
| 6:13.272 | נימרו |  |
| 6:13.273 | ליה |  |
| 6:13.274 | מי |  |
| 6:13.275 | ימלל |  |
| 6:13.276 | גבורות |  |
| 6:13.277 | ה | ' |
| 6:13.278 | ישמיע |  |
| 6:13.279 | כל |  |
| 6:13.280 | תהלתו | , |
| 6:13.281 | למי |  |
| 6:13.282 | נאה |  |
| 6:13.283 | למלל |  |
| 6:13.284 | גבורות |  |
| 6:13.285 | ה | ', |
| 6:13.286 | למי |  |
| 6:13.287 | שיכול |  |
| 6:13.288 | להשמיע |  |
| 6:13.289 | כל |  |
| 6:13.290 | תהלתו |  |
| 6:13.291 | וכו | ', |
| 6:13.292 | אלמא |  |
| 6:13.293 | דקרו |  |
| 6:13.294 | רבנן |  |
| 6:13.295 | למשנה |  |
| 6:13.296 | תהלות |  |
| 6:13.297 | ה | ', |
| 6:13.298 | ובמדרש |  |
| 6:13.299 | תנחומא |  |
| 6:13.300 | פרשת |  |
| 6:13.301 | וירא |  |
| 6:13.302 | גרסינן |  |
| 6:13.303 | אמר |  |
| 6:13.304 | רב |  |
| 6:13.305 | יהודה |  |
| 6:13.306 | בן |  |
| 6:13.307 | שלום |  |
| 6:13.308 | בקש |  |
| 6:13.309 | משה |  |
| 6:13.310 | שתהא |  |
| 6:13.311 | המשנה |  |
| 6:13.312 | אף |  |
| 6:13.313 | היא |  |
| 6:13.314 | בכתב |  |
| 6:13.315 | וצפה |  |
| 6:13.316 | הקב"ה |  |
| 6:13.317 | שאומות |  |
| 6:13.318 | העולם |  |
| 6:13.319 | עתידן |  |
| 6:13.320 | לתרגם |  |
| 6:13.321 | את |  |
| 6:13.322 | התורה |  |
| 6:13.323 | ולהיות |  |
| 6:13.324 | קורין |  |
| 6:13.325 | אותה |  |
| 6:13.326 | יונית |  |
| 6:13.327 | והן |  |
| 6:13.328 | אומרין |  |
| 6:13.329 | אנו |  |
| 6:13.330 | הן |  |
| 6:13.331 | של |  |
| 6:13.332 | ישראל | . |
| 6:13.333 | אמר |  |
| 6:13.334 | לו |  |
| 6:13.335 | הקב"ה |  |
| 6:13.336 | למשה |  |
| 6:13.337 | אכתב |  |
| 6:13.338 | לו |  |
| 6:13.339 | רובי |  |
| 6:13.340 | תורתי |  |
| 6:13.341 | ואם |  |
| 6:13.342 | כן |  |
| 6:13.343 | כמו |  |
| 6:13.344 | זר |  |
| 6:13.345 | נחשבו | , |
| 6:13.346 | וכל |  |
| 6:13.347 | כך |  |
| 6:13.348 | למה |  |
| 6:13.349 | מפני |  |
| 6:13.350 | שהמשנה |  |
| 6:13.351 | היא |  |
| 6:13.352 | מסתורין |  |
| 6:13.353 | של |  |
| 6:13.354 | הקב"ה | . |
| 6:13.355 | ואין |  |
| 6:13.356 | הקב"ה |  |
| 6:13.357 | מוסר |  |
| 6:13.358 | מסתורין |  |
| 6:13.359 | שלו |  |
| 6:13.360 | אלא |  |
| 6:13.361 | לצדיקים |  |
| 6:13.362 | שנ | ' |
| 6:13.363 | סוד |  |
| 6:13.364 | ה | ' |
| 6:13.365 | ליריאיו | . |
| 6:13.366 | ועיין |  |
| 6:13.367 | מה |  |
| 6:13.368 | שכ | ' |
| 6:13.369 | בזה |  |
| 6:13.370 | מהר"י |  |
| 6:13.371 | מדובנא |  |
| 6:13.372 | בספרו |  |
| 6:13.373 | הנחמד |  |
| 6:13.374 | אהל |  |
| 6:13.375 | יעקב |  |
| 6:13.376 | בחתימת |  |
| 6:13.377 | ספר |  |
| 6:13.378 | במדבר |  |
| 6:13.379 | וינעם |  |
| 6:13.380 | לך | . |
| 6:13.381 | ראה |  |
| 6:13.382 | כמה |  |
| 6:13.383 | היה |  |
| 6:13.384 | חביבה |  |
| 6:13.385 | עליהם |  |
| 6:13.386 | המשנה |  |
| 6:13.387 | וקראו |  |
| 6:13.388 | לה |  |
| 6:13.389 | מסתורין |  |
| 6:13.390 | של |  |
| 6:13.391 | הקב"ה |  |
| 6:13.392 | שאינו |  |
| 6:13.393 | מוסרה |  |
| 6:13.394 | אלא |  |
| 6:13.395 | לצדיקים |  |
| 6:13.396 | שנ | ' |
| 6:13.397 | סוד |  |
| 6:13.398 | ה | ' |
| 6:13.399 | ליריאיו | . |
| 6:13.400 | וכן |  |
| 6:13.401 | במדרש |  |
| 6:13.402 | שמות |  |
| 6:13.403 | רבה |  |
| 6:13.404 | פרשה |  |
| 6:13.405 | מ"ז |  |
| 6:13.406 | אמרו |  |
| 6:13.407 | כתב |  |
| 6:13.408 | לך |  |
| 6:13.409 | את |  |
| 6:13.410 | הדברים |  |
| 6:13.411 | האלה |  |
| 6:13.412 | וגו | ' |
| 6:13.413 | הדא |  |
| 6:13.414 | הוא |  |
| 6:13.415 | דכתיב |  |
| 6:13.416 | אכתב |  |
| 6:13.417 | לו |  |
| 6:13.418 | רובי |  |
| 6:13.419 | תורתי |  |
| 6:13.420 | כמו |  |
| 6:13.421 | זר |  |
| 6:13.422 | נחשבו |  |
| 6:13.423 | בשעה |  |
| 6:13.424 | שנגלה |  |
| 6:13.425 | הקב"ה |  |
| 6:13.426 | בסיני |  |
| 6:13.427 | ליתן |  |
| 6:13.428 | תורה |  |
| 6:13.429 | לישראל |  |
| 6:13.430 | אמרה |  |
| 6:13.431 | למשה |  |
| 6:13.432 | על |  |
| 6:13.433 | הסדר |  |
| 6:13.434 | מקרא | , |
| 6:13.435 | משנה | , |
| 6:13.436 | תלמוד | , |
| 6:13.437 | ואגדה | , |
| 6:13.438 | שנאמר |  |
| 6:13.439 | וידבר |  |
| 6:13.440 | אלהים |  |
| 6:13.441 | את |  |
| 6:13.442 | כל |  |
| 6:13.443 | הדברים |  |
| 6:13.444 | האלה | , |
| 6:13.445 | אפי | ' |
| 6:13.446 | מה |  |
| 6:13.447 | שתלמיד |  |
| 6:13.448 | שואל |  |
| 6:13.449 | לרב |  |
| 6:13.450 | אמר |  |
| 6:13.451 | הקב"ה |  |
| 6:13.452 | למשה |  |
| 6:13.453 | באותה |  |
| 6:13.454 | שעה |  |
| 6:13.455 | ומאחר |  |
| 6:13.456 | שלמד |  |
| 6:13.457 | אותה |  |
| 6:13.458 | מפי |  |
| 6:13.459 | הקב"ה |  |
| 6:13.460 | אמר |  |
| 6:13.461 | לו |  |
| 6:13.462 | למדה |  |
| 6:13.463 | לישראל | , |
| 6:13.464 | אמר |  |
| 6:13.465 | לפניו |  |
| 6:13.466 | רבונו |  |
| 6:13.467 | של |  |
| 6:13.468 | עולם |  |
| 6:13.469 | אכתוב |  |
| 6:13.470 | אותה |  |
| 6:13.471 | להם | , |
| 6:13.472 | אמר |  |
| 6:13.473 | איני |  |
| 6:13.474 | מבקש |  |
| 6:13.475 | ליתנה |  |
| 6:13.476 | להם |  |
| 6:13.477 | בכתב | , |
| 6:13.478 | מפני |  |
| 6:13.479 | שגלוי |  |
| 6:13.480 | לפני |  |
| 6:13.481 | שאומות |  |
| 6:13.482 | העולם |  |
| 6:13.483 | עתידים |  |
| 6:13.484 | לשלוט |  |
| 6:13.485 | בהם |  |
| 6:13.486 | וליטול |  |
| 6:13.487 | אותה |  |
| 6:13.488 | מהם |  |
| 6:13.489 | ויהיו |  |
| 6:13.490 | בזוים |  |
| 6:13.491 | באומות |  |
| 6:13.492 | פי | ' |
| 6:13.493 | שיאמרו |  |
| 6:13.494 | אומות |  |
| 6:13.495 | העולם |  |
| 6:13.496 | שניתנה |  |
| 6:13.497 | להם | ) |
| 6:13.498 | אלא |  |
| 6:13.499 | המקרא |  |
| 6:13.500 | אני |  |
| 6:13.501 | נותן |  |
| 6:13.502 | להם |  |
| 6:13.503 | בכתב |  |
| 6:13.504 | והמשנה |  |
| 6:13.505 | והתלמוד |  |
| 6:13.506 | והאגדה |  |
| 6:13.507 | בעל |  |
| 6:13.508 | פה |  |
| 6:13.509 | וכו | ' |
| 6:13.510 | ומסיים |  |
| 6:13.511 | כי |  |
| 6:13.512 | על |  |
| 6:13.513 | פי |  |
| 6:13.514 | הדברים |  |
| 6:13.515 | האלה |  |
| 6:13.516 | זו |  |
| 6:13.517 | המשנה |  |
| 6:13.518 | והתלמוד |  |
| 6:13.519 | שהם |  |
| 6:13.520 | מבדילים |  |
| 6:13.521 | בין |  |
| 6:13.522 | ישראל |  |
| 6:13.523 | לבין |  |
| 6:13.524 | האומות |  |
| 6:13.525 | ע"כ | . |
| 6:13.526 | וכן |  |
| 6:13.527 | כתבו |  |
| 6:13.528 | קדמונינו |  |
| 6:13.529 | הר | ' |
| 6:13.530 | חטר |  |
| 6:13.531 | ור | ' |
| 6:13.532 | זכריה |  |
| 6:13.533 | הרופא |  |
| 6:13.534 | בפי | ' |
| 6:13.535 | הי"ג |  |
| 6:13.536 | עיקרים |  |
| 6:13.537 | בעיקר |  |
| 6:13.538 | השמיני |  |
| 6:13.539 | וז"ל |  |
| 6:13.540 | ואנמי | ' |
| 6:13.541 | כאן |  |
| 6:13.542 | על |  |
| 6:13.543 | פה |  |
| 6:13.544 | לפי |  |
| 6:13.545 | שידע |  |
| 6:13.546 | הקב"ה |  |
| 6:13.547 | והודיע |  |
| 6:13.548 | למשה |  |
| 6:13.549 | שאומות |  |
| 6:13.550 | העולם |  |
| 6:13.551 | עתידין |  |
| 6:13.552 | ליקח |  |
| 6:13.553 | התורה |  |
| 6:13.554 | על |  |
| 6:13.555 | ישראל | . |
| 6:13.556 | וכיון |  |
| 6:13.557 | שלקחו |  |
| 6:13.558 | תורה |  |
| 6:13.559 | שבכתב |  |
| 6:13.560 | לא |  |
| 6:13.561 | יוכלו |  |
| 6:13.562 | ליקח |  |
| 6:13.563 | תורה |  |
| 6:13.564 | שבעל |  |
| 6:13.565 | פה |  |
| 6:13.566 | מפני |  |
| 6:13.567 | שהיא |  |
| 6:13.568 | בסופרים |  |
| 6:13.569 | ולא |  |
| 6:13.570 | בספרים | . |
| 6:13.571 | ואלו |  |
| 6:13.572 | לקחו |  |
| 6:13.573 | תורה |  |
| 6:13.574 | שבכתב |  |
| 6:13.575 | עם |  |
| 6:13.576 | תורה |  |
| 6:13.577 | שבעל |  |
| 6:13.578 | פה |  |
| 6:13.579 | היו |  |
| 6:13.580 | אומרים |  |
| 6:13.581 | שהם |  |
| 6:13.582 | ישראל |  |
| 6:13.583 | בעדות |  |
| 6:13.584 | נאמנה |  |
| 6:13.585 | ע"כ | . |
| 6:13.586 | ראה |  |
| 6:13.587 | ידידי |  |
| 6:13.588 | הקורא |  |
| 6:13.589 | והבן |  |
| 6:13.590 | דברי |  |
| 6:13.591 | רז"ל |  |
| 6:13.592 | שאמרו |  |
| 6:13.593 | כי |  |
| 6:13.594 | על |  |
| 6:13.595 | המשנה |  |
| 6:13.596 | והתלמוד |  |
| 6:13.597 | שהם |  |
| 6:13.598 | יסוד |  |
| 6:13.599 | כל |  |
| 6:13.600 | תורה |  |
| 6:13.601 | שבעל |  |
| 6:13.602 | פה |  |
| 6:13.603 | כרת |  |
| 6:13.604 | השי"ת |  |
| 6:13.605 | ברית |  |
| 6:13.606 | עם |  |
| 6:13.607 | ישראל | . |
| 6:13.608 | כי |  |
| 6:13.609 | המה |  |
| 6:13.610 | חיי |  |
| 6:13.611 | רוחנו |  |
| 6:13.612 | לעולם |  |
| 6:13.613 | הבא |  |
| 6:13.614 | והם |  |
| 6:13.615 | רובה |  |
| 6:13.616 | של |  |
| 6:13.617 | תורה |  |
| 6:13.618 | כדאיתא |  |
| 6:13.619 | בגיטין |  |
| 6:13.620 | פרק |  |
| 6:13.621 | הנזקין | . |
| 6:13.622 | ובפרק |  |
| 6:13.623 | קמא |  |
| 6:13.624 | דברכות |  |
| 6:13.625 | גרסינן |  |
| 6:13.626 | ואמר |  |
| 6:13.627 | ר | ' |
| 6:13.628 | לוי |  |
| 6:13.629 | בר |  |
| 6:13.630 | חמא |  |
| 6:13.631 | אמר |  |
| 6:13.632 | ריש |  |
| 6:13.633 | לקיש |  |
| 6:13.634 | מאי |  |
| 6:13.635 | דכתיב |  |
| 6:13.636 | ואתנה |  |
| 6:13.637 | לך |  |
| 6:13.638 | את |  |
| 6:13.639 | לוחות |  |
| 6:13.640 | האבן |  |
| 6:13.641 | והתורה |  |
| 6:13.642 | והמצוה |  |
| 6:13.643 | אשר |  |
| 6:13.644 | כתבתי |  |
| 6:13.645 | להורותם | , |
| 6:13.646 | לוחות |  |
| 6:13.647 | אלו |  |
| 6:13.648 | עשרת |  |
| 6:13.649 | הדברות | , |
| 6:13.650 | והתורה |  |
| 6:13.651 | זו |  |
| 6:13.652 | מקרא | . |
| 6:13.653 | והמצוה |  |
| 6:13.654 | זו |  |
| 6:13.655 | משנה | , |
| 6:13.656 | אשר |  |
| 6:13.657 | כתבתי |  |
| 6:13.658 | אלו |  |
| 6:13.659 | נביאים |  |
| 6:13.660 | וכתובים | , |
| 6:13.661 | להורותם |  |
| 6:13.662 | זו |  |
| 6:13.663 | גמרא | . |
| 6:13.664 | מלמד |  |
| 6:13.665 | שכולם |  |
| 6:13.666 | ניתנו |  |
| 6:13.667 | למשה |  |
| 6:13.668 | מסיני |  |
| 6:13.669 | ע"כ | . |
| 6:13.670 | ומביאו |  |
| 6:13.671 | רבינו |  |
| 6:13.672 | הרמב"ם |  |
| 6:13.673 | בהקדמת |  |
| 6:13.674 | ספר |  |
| 6:13.675 | היד |  |
| 6:13.676 | עיין |  |
| 6:13.677 | שם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

