# Translation pack: milhamot-hashem perek-9-02 (Milchamot ha-Shem 9:2–9:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-02`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
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
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
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
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
- **רוע** לב → "wickedness of heart".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
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


## Milchamot ha-Shem 9:2

Hebrew (Wikisource): קכא) כתב הרמב"ם ז"ל בפרק ח' מהלכות יסודי התורה, משה רבינו לא האמינו בו ישראל מפני האותות שעשה שהמאמין על פי האותות יש בלבו דופי שאפשא שיעשה האות בלאט וכשוף וכו' ובמה האמינו במעמד הר סיני. שעינינו ראו ולא זר ואזנינו שמעו ולא אחר. האש ובקולות והלפידים. והוא נגש אל הערפל והקול מדבר אליו,ואנו שומעים, משה משה לך אמור להם כך וכך, וכן הוא אומר פנים בפנים דבר ה' עמכם. ונאמר לא את אבותינו כרת ה' את הברית הזאת. ומנין שמעמד הר סיני לבדו היא הראיה לנבואתו שהיא אמת שאין בו דופי, שנאמר הנה אנכי בא אליך בעב הענן בעבור ישמע העם בדברי עמך וגם בך יאמינו לעולם, מכלל שקודם דבר זה לא האמינו בו נאמנות שהיא עומדת לעולם אלא נאמנות שיש אחריה הרהור ומחשבה, נמצאו אלו ששולח אליהם, הם עדים על נבואתו שהיא אמת ואינו צריך לעשות להן אות שהן והוא אחד בדבר, כשני עדים שראו דבר אחד ביחד, שכל אחד מהן עד לחבירו שהוא אומר אמת ואין אחד מהן צריך להביא ראיה לחבירו, כך משה רבינו ע"ה כל ישראל עדים לו מעמד הר סיני, ואינו צריך לעשות להם אות לפיכך אם עמד נביא ועשהאותות ומופתים גדולים ובקש להכחיש נבואתו של משה רבינו (ואפילו בפרט אחד או להוסיף או לגרוע בשום מצוה קלה או חמורה) אין שומעין לו ואנו יודעין ביחוד שאותן האותות בלאט וכשוף הן. לפי שנבואת משה רבינו אינה על פי האותות כדי שנערוך אותות זה לאותות זה. אלא בעינינו ראינוה ובאזנינו שמענוה כמו שממע הוא, הא למה הדבר דומה לעדים שהעידו על דבר שראה הוא בעיניו שאינו כמו שראה שאינו שומע להן. אלא יודע בודאי שהן עידי שקר. לפיכך אמרה תורה שאם בא האות והמופת לא תשמע אל דברי הנביא ההוא. שהרי זה בא להכחיש מה שראית בעיניך עכ"ל. וכן כתב הסמ"ג בהקדמת העשין וז"ל טכאשר רצה אדון הכל ליתן תורה, הטה שמים ושמי השמים על הר סיני בקולות ולפידים נפלאים ונוראים, וקרא למשה בחירו ואמר לו הנה אנכי בא אליך בעב הענן בעבור ישמע העם בדברי עמך וגם בך יאמינו לעולם, ולמה הוצרך לזה והלא כבר כתוב על הים ויאמינו בה' ובמשה עבדו? אלא כך אמר הקב"ה למשה אני רוצה שיאמינו בך ישראל במה שקדם שאתה נביא על ידי אותות ומופתים שעשית, אבל תורה שאני רוצה ליתן איני רוצה שיאמינו בך ישראל על ידי אות ומופת. אלא על ידי שישמעו באזניהם שאני מדבר עמך ולכן עשה הקב"ה את הדבר הזה, שכשיהיו ישראל בגולה. אם יאמר להם עכו"ם או ישמעאל לעזוב תורתם ולעבוד אלהים אחרים ויתן להם אות ומופת יכולין ישראל לענות לו, אפילו תתעשה הרבה אותות ומופתים כמשה בן עמרם לא נאמין לך בחילוף תורתינו, אלא אם כן נשמע באזנינו שידבר עמך השם כמו שדיבר עם משה, כי אפילו משה בחירו לא רצה שנאמינהו בתורה אלא על ידי ששמענו שדיבר עמו. ולאחר שנתן את התורה אמר אתם ראיתם כי מן השמים דברתי עמכם. כלומר אין לכם להאמין עוד לשום אדם שיאמר שהחלפתי תורה אם לא תראו שאדבר עמכם מן השמים להחליפה, ולפי שאמרנו שאין לנו תורה לשום עכו"ם או ישמעאל באות או מופת שיעשה להחליף תורתינו, פי' בתורה היאך יש יכולת לעשות אות או מופת. שנאמר כי יקום בקרבך נביא או חולם וכו' כי מנסה ה' אלהיכם אתכם, הרי פירש לך שזה האות והמות אינו מכח ע"ז שלו שלא הקב"ה עושה זה האות לנסות בו את ישראל, ואחר פטירת משה רבינו עמדו נביאים לישראל ודברו בחידות ומשלים, ועל ידי כך אמרו הגוים כי הנביאים נבאו בתורה חדשה שלהם לכך שלח הקב"ה לישראל ביד מלאכי שהוא אחרון לכל הנביאים ואמר שני דברים בסוף דבריו לחתום כל הנבואות הנה אנכי שלח לכן את אליה הנביא וכו' כלומר הנני אחרון לכל הנביאים. וממני עד שיבוא אליהו לא יעמוד נביא, והנני אומר לכם בשם הקב"ה זכרו תורת משה עבדי אשר צויתי אותו בחורב אל כל ישראל חקים ומשפטים כלומר אל יעלה על לבבכם שהנביאים שעמדו לפני נתנבאו בחילוף תורת משה, כי כולם לחזק תורת משה באו עכ"ל הסמ"ג קרוב לדברי הרמב"ם ז"ל. וכן דרשו רז"ל בפסוק שובי שובי השולמית הובא במדרש הגדול ויש פרשת וישלח ע"ש.

| id | word | punctuation after |
|---|---|---|
| 9:2.1 | קכא | ) |
| 9:2.2 | כתב |  |
| 9:2.3 | הרמב"ם |  |
| 9:2.4 | ז"ל |  |
| 9:2.5 | בפרק |  |
| 9:2.6 | ח | ' |
| 9:2.7 | מהלכות |  |
| 9:2.8 | יסודי |  |
| 9:2.9 | התורה | , |
| 9:2.10 | משה |  |
| 9:2.11 | רבינו |  |
| 9:2.12 | לא |  |
| 9:2.13 | האמינו |  |
| 9:2.14 | בו |  |
| 9:2.15 | ישראל |  |
| 9:2.16 | מפני |  |
| 9:2.17 | האותות |  |
| 9:2.18 | שעשה |  |
| 9:2.19 | שהמאמין |  |
| 9:2.20 | על |  |
| 9:2.21 | פי |  |
| 9:2.22 | האותות |  |
| 9:2.23 | יש |  |
| 9:2.24 | בלבו |  |
| 9:2.25 | דופי |  |
| 9:2.26 | שאפשא |  |
| 9:2.27 | שיעשה |  |
| 9:2.28 | האות |  |
| 9:2.29 | בלאט |  |
| 9:2.30 | וכשוף |  |
| 9:2.31 | וכו | ' |
| 9:2.32 | ובמה |  |
| 9:2.33 | האמינו |  |
| 9:2.34 | במעמד |  |
| 9:2.35 | הר |  |
| 9:2.36 | סיני | . |
| 9:2.37 | שעינינו |  |
| 9:2.38 | ראו |  |
| 9:2.39 | ולא |  |
| 9:2.40 | זר |  |
| 9:2.41 | ואזנינו |  |
| 9:2.42 | שמעו |  |
| 9:2.43 | ולא |  |
| 9:2.44 | אחר | . |
| 9:2.45 | האש |  |
| 9:2.46 | ובקולות |  |
| 9:2.47 | והלפידים | . |
| 9:2.48 | והוא |  |
| 9:2.49 | נגש |  |
| 9:2.50 | אל |  |
| 9:2.51 | הערפל |  |
| 9:2.52 | והקול |  |
| 9:2.53 | מדבר |  |
| 9:2.54 | אליו,ואנו |  |
| 9:2.55 | שומעים | , |
| 9:2.56 | משה |  |
| 9:2.57 | משה |  |
| 9:2.58 | לך |  |
| 9:2.59 | אמור |  |
| 9:2.60 | להם |  |
| 9:2.61 | כך |  |
| 9:2.62 | וכך | , |
| 9:2.63 | וכן |  |
| 9:2.64 | הוא |  |
| 9:2.65 | אומר |  |
| 9:2.66 | פנים |  |
| 9:2.67 | בפנים |  |
| 9:2.68 | דבר |  |
| 9:2.69 | ה | ' |
| 9:2.70 | עמכם | . |
| 9:2.71 | ונאמר |  |
| 9:2.72 | לא |  |
| 9:2.73 | את |  |
| 9:2.74 | אבותינו |  |
| 9:2.75 | כרת |  |
| 9:2.76 | ה | ' |
| 9:2.77 | את |  |
| 9:2.78 | הברית |  |
| 9:2.79 | הזאת | . |
| 9:2.80 | ומנין |  |
| 9:2.81 | שמעמד |  |
| 9:2.82 | הר |  |
| 9:2.83 | סיני |  |
| 9:2.84 | לבדו |  |
| 9:2.85 | היא |  |
| 9:2.86 | הראיה |  |
| 9:2.87 | לנבואתו |  |
| 9:2.88 | שהיא |  |
| 9:2.89 | אמת |  |
| 9:2.90 | שאין |  |
| 9:2.91 | בו |  |
| 9:2.92 | דופי | , |
| 9:2.93 | שנאמר |  |
| 9:2.94 | הנה |  |
| 9:2.95 | אנכי |  |
| 9:2.96 | בא |  |
| 9:2.97 | אליך |  |
| 9:2.98 | בעב |  |
| 9:2.99 | הענן |  |
| 9:2.100 | בעבור |  |
| 9:2.101 | ישמע |  |
| 9:2.102 | העם |  |
| 9:2.103 | בדברי |  |
| 9:2.104 | עמך |  |
| 9:2.105 | וגם |  |
| 9:2.106 | בך |  |
| 9:2.107 | יאמינו |  |
| 9:2.108 | לעולם | , |
| 9:2.109 | מכלל |  |
| 9:2.110 | שקודם |  |
| 9:2.111 | דבר |  |
| 9:2.112 | זה |  |
| 9:2.113 | לא |  |
| 9:2.114 | האמינו |  |
| 9:2.115 | בו |  |
| 9:2.116 | נאמנות |  |
| 9:2.117 | שהיא |  |
| 9:2.118 | עומדת |  |
| 9:2.119 | לעולם |  |
| 9:2.120 | אלא |  |
| 9:2.121 | נאמנות |  |
| 9:2.122 | שיש |  |
| 9:2.123 | אחריה |  |
| 9:2.124 | הרהור |  |
| 9:2.125 | ומחשבה | , |
| 9:2.126 | נמצאו |  |
| 9:2.127 | אלו |  |
| 9:2.128 | ששולח |  |
| 9:2.129 | אליהם | , |
| 9:2.130 | הם |  |
| 9:2.131 | עדים |  |
| 9:2.132 | על |  |
| 9:2.133 | נבואתו |  |
| 9:2.134 | שהיא |  |
| 9:2.135 | אמת |  |
| 9:2.136 | ואינו |  |
| 9:2.137 | צריך |  |
| 9:2.138 | לעשות |  |
| 9:2.139 | להן |  |
| 9:2.140 | אות |  |
| 9:2.141 | שהן |  |
| 9:2.142 | והוא |  |
| 9:2.143 | אחד |  |
| 9:2.144 | בדבר | , |
| 9:2.145 | כשני |  |
| 9:2.146 | עדים |  |
| 9:2.147 | שראו |  |
| 9:2.148 | דבר |  |
| 9:2.149 | אחד |  |
| 9:2.150 | ביחד | , |
| 9:2.151 | שכל |  |
| 9:2.152 | אחד |  |
| 9:2.153 | מהן |  |
| 9:2.154 | עד |  |
| 9:2.155 | לחבירו |  |
| 9:2.156 | שהוא |  |
| 9:2.157 | אומר |  |
| 9:2.158 | אמת |  |
| 9:2.159 | ואין |  |
| 9:2.160 | אחד |  |
| 9:2.161 | מהן |  |
| 9:2.162 | צריך |  |
| 9:2.163 | להביא |  |
| 9:2.164 | ראיה |  |
| 9:2.165 | לחבירו | , |
| 9:2.166 | כך |  |
| 9:2.167 | משה |  |
| 9:2.168 | רבינו |  |
| 9:2.169 | ע"ה |  |
| 9:2.170 | כל |  |
| 9:2.171 | ישראל |  |
| 9:2.172 | עדים |  |
| 9:2.173 | לו |  |
| 9:2.174 | מעמד |  |
| 9:2.175 | הר |  |
| 9:2.176 | סיני | , |
| 9:2.177 | ואינו |  |
| 9:2.178 | צריך |  |
| 9:2.179 | לעשות |  |
| 9:2.180 | להם |  |
| 9:2.181 | אות |  |
| 9:2.182 | לפיכך |  |
| 9:2.183 | אם |  |
| 9:2.184 | עמד |  |
| 9:2.185 | נביא |  |
| 9:2.186 | ועשהאותות |  |
| 9:2.187 | ומופתים |  |
| 9:2.188 | גדולים |  |
| 9:2.189 | ובקש |  |
| 9:2.190 | להכחיש |  |
| 9:2.191 | נבואתו |  |
| 9:2.192 | של |  |
| 9:2.193 | משה |  |
| 9:2.194 | רבינו |  |
| 9:2.195 | ואפילו |  |
| 9:2.196 | בפרט |  |
| 9:2.197 | אחד |  |
| 9:2.198 | או |  |
| 9:2.199 | להוסיף |  |
| 9:2.200 | או |  |
| 9:2.201 | לגרוע |  |
| 9:2.202 | בשום |  |
| 9:2.203 | מצוה |  |
| 9:2.204 | קלה |  |
| 9:2.205 | או |  |
| 9:2.206 | חמורה | ) |
| 9:2.207 | אין |  |
| 9:2.208 | שומעין |  |
| 9:2.209 | לו |  |
| 9:2.210 | ואנו |  |
| 9:2.211 | יודעין |  |
| 9:2.212 | ביחוד |  |
| 9:2.213 | שאותן |  |
| 9:2.214 | האותות |  |
| 9:2.215 | בלאט |  |
| 9:2.216 | וכשוף |  |
| 9:2.217 | הן | . |
| 9:2.218 | לפי |  |
| 9:2.219 | שנבואת |  |
| 9:2.220 | משה |  |
| 9:2.221 | רבינו |  |
| 9:2.222 | אינה |  |
| 9:2.223 | על |  |
| 9:2.224 | פי |  |
| 9:2.225 | האותות |  |
| 9:2.226 | כדי |  |
| 9:2.227 | שנערוך |  |
| 9:2.228 | אותות |  |
| 9:2.229 | זה |  |
| 9:2.230 | לאותות |  |
| 9:2.231 | זה | . |
| 9:2.232 | אלא |  |
| 9:2.233 | בעינינו |  |
| 9:2.234 | ראינוה |  |
| 9:2.235 | ובאזנינו |  |
| 9:2.236 | שמענוה |  |
| 9:2.237 | כמו |  |
| 9:2.238 | שממע |  |
| 9:2.239 | הוא | , |
| 9:2.240 | הא |  |
| 9:2.241 | למה |  |
| 9:2.242 | הדבר |  |
| 9:2.243 | דומה |  |
| 9:2.244 | לעדים |  |
| 9:2.245 | שהעידו |  |
| 9:2.246 | על |  |
| 9:2.247 | דבר |  |
| 9:2.248 | שראה |  |
| 9:2.249 | הוא |  |
| 9:2.250 | בעיניו |  |
| 9:2.251 | שאינו |  |
| 9:2.252 | כמו |  |
| 9:2.253 | שראה |  |
| 9:2.254 | שאינו |  |
| 9:2.255 | שומע |  |
| 9:2.256 | להן | . |
| 9:2.257 | אלא |  |
| 9:2.258 | יודע |  |
| 9:2.259 | בודאי |  |
| 9:2.260 | שהן |  |
| 9:2.261 | עידי |  |
| 9:2.262 | שקר | . |
| 9:2.263 | לפיכך |  |
| 9:2.264 | אמרה |  |
| 9:2.265 | תורה |  |
| 9:2.266 | שאם |  |
| 9:2.267 | בא |  |
| 9:2.268 | האות |  |
| 9:2.269 | והמופת |  |
| 9:2.270 | לא |  |
| 9:2.271 | תשמע |  |
| 9:2.272 | אל |  |
| 9:2.273 | דברי |  |
| 9:2.274 | הנביא |  |
| 9:2.275 | ההוא | . |
| 9:2.276 | שהרי |  |
| 9:2.277 | זה |  |
| 9:2.278 | בא |  |
| 9:2.279 | להכחיש |  |
| 9:2.280 | מה |  |
| 9:2.281 | שראית |  |
| 9:2.282 | בעיניך |  |
| 9:2.283 | עכ"ל | . |
| 9:2.284 | וכן |  |
| 9:2.285 | כתב |  |
| 9:2.286 | הסמ"ג |  |
| 9:2.287 | בהקדמת |  |
| 9:2.288 | העשין |  |
| 9:2.289 | וז"ל |  |
| 9:2.290 | טכאשר |  |
| 9:2.291 | רצה |  |
| 9:2.292 | אדון |  |
| 9:2.293 | הכל |  |
| 9:2.294 | ליתן |  |
| 9:2.295 | תורה | , |
| 9:2.296 | הטה |  |
| 9:2.297 | שמים |  |
| 9:2.298 | ושמי |  |
| 9:2.299 | השמים |  |
| 9:2.300 | על |  |
| 9:2.301 | הר |  |
| 9:2.302 | סיני |  |
| 9:2.303 | בקולות |  |
| 9:2.304 | ולפידים |  |
| 9:2.305 | נפלאים |  |
| 9:2.306 | ונוראים | , |
| 9:2.307 | וקרא |  |
| 9:2.308 | למשה |  |
| 9:2.309 | בחירו |  |
| 9:2.310 | ואמר |  |
| 9:2.311 | לו |  |
| 9:2.312 | הנה |  |
| 9:2.313 | אנכי |  |
| 9:2.314 | בא |  |
| 9:2.315 | אליך |  |
| 9:2.316 | בעב |  |
| 9:2.317 | הענן |  |
| 9:2.318 | בעבור |  |
| 9:2.319 | ישמע |  |
| 9:2.320 | העם |  |
| 9:2.321 | בדברי |  |
| 9:2.322 | עמך |  |
| 9:2.323 | וגם |  |
| 9:2.324 | בך |  |
| 9:2.325 | יאמינו |  |
| 9:2.326 | לעולם | , |
| 9:2.327 | ולמה |  |
| 9:2.328 | הוצרך |  |
| 9:2.329 | לזה |  |
| 9:2.330 | והלא |  |
| 9:2.331 | כבר |  |
| 9:2.332 | כתוב |  |
| 9:2.333 | על |  |
| 9:2.334 | הים |  |
| 9:2.335 | ויאמינו |  |
| 9:2.336 | בה | ' |
| 9:2.337 | ובמשה |  |
| 9:2.338 | עבדו | ? |
| 9:2.339 | אלא |  |
| 9:2.340 | כך |  |
| 9:2.341 | אמר |  |
| 9:2.342 | הקב"ה |  |
| 9:2.343 | למשה |  |
| 9:2.344 | אני |  |
| 9:2.345 | רוצה |  |
| 9:2.346 | שיאמינו |  |
| 9:2.347 | בך |  |
| 9:2.348 | ישראל |  |
| 9:2.349 | במה |  |
| 9:2.350 | שקדם |  |
| 9:2.351 | שאתה |  |
| 9:2.352 | נביא |  |
| 9:2.353 | על |  |
| 9:2.354 | ידי |  |
| 9:2.355 | אותות |  |
| 9:2.356 | ומופתים |  |
| 9:2.357 | שעשית | , |
| 9:2.358 | אבל |  |
| 9:2.359 | תורה |  |
| 9:2.360 | שאני |  |
| 9:2.361 | רוצה |  |
| 9:2.362 | ליתן |  |
| 9:2.363 | איני |  |
| 9:2.364 | רוצה |  |
| 9:2.365 | שיאמינו |  |
| 9:2.366 | בך |  |
| 9:2.367 | ישראל |  |
| 9:2.368 | על |  |
| 9:2.369 | ידי |  |
| 9:2.370 | אות |  |
| 9:2.371 | ומופת | . |
| 9:2.372 | אלא |  |
| 9:2.373 | על |  |
| 9:2.374 | ידי |  |
| 9:2.375 | שישמעו |  |
| 9:2.376 | באזניהם |  |
| 9:2.377 | שאני |  |
| 9:2.378 | מדבר |  |
| 9:2.379 | עמך |  |
| 9:2.380 | ולכן |  |
| 9:2.381 | עשה |  |
| 9:2.382 | הקב"ה |  |
| 9:2.383 | את |  |
| 9:2.384 | הדבר |  |
| 9:2.385 | הזה | , |
| 9:2.386 | שכשיהיו |  |
| 9:2.387 | ישראל |  |
| 9:2.388 | בגולה | . |
| 9:2.389 | אם |  |
| 9:2.390 | יאמר |  |
| 9:2.391 | להם |  |
| 9:2.392 | עכו"ם |  |
| 9:2.393 | או |  |
| 9:2.394 | ישמעאל |  |
| 9:2.395 | לעזוב |  |
| 9:2.396 | תורתם |  |
| 9:2.397 | ולעבוד |  |
| 9:2.398 | אלהים |  |
| 9:2.399 | אחרים |  |
| 9:2.400 | ויתן |  |
| 9:2.401 | להם |  |
| 9:2.402 | אות |  |
| 9:2.403 | ומופת |  |
| 9:2.404 | יכולין |  |
| 9:2.405 | ישראל |  |
| 9:2.406 | לענות |  |
| 9:2.407 | לו | , |
| 9:2.408 | אפילו |  |
| 9:2.409 | תתעשה |  |
| 9:2.410 | הרבה |  |
| 9:2.411 | אותות |  |
| 9:2.412 | ומופתים |  |
| 9:2.413 | כמשה |  |
| 9:2.414 | בן |  |
| 9:2.415 | עמרם |  |
| 9:2.416 | לא |  |
| 9:2.417 | נאמין |  |
| 9:2.418 | לך |  |
| 9:2.419 | בחילוף |  |
| 9:2.420 | תורתינו | , |
| 9:2.421 | אלא |  |
| 9:2.422 | אם |  |
| 9:2.423 | כן |  |
| 9:2.424 | נשמע |  |
| 9:2.425 | באזנינו |  |
| 9:2.426 | שידבר |  |
| 9:2.427 | עמך |  |
| 9:2.428 | השם |  |
| 9:2.429 | כמו |  |
| 9:2.430 | שדיבר |  |
| 9:2.431 | עם |  |
| 9:2.432 | משה | , |
| 9:2.433 | כי |  |
| 9:2.434 | אפילו |  |
| 9:2.435 | משה |  |
| 9:2.436 | בחירו |  |
| 9:2.437 | לא |  |
| 9:2.438 | רצה |  |
| 9:2.439 | שנאמינהו |  |
| 9:2.440 | בתורה |  |
| 9:2.441 | אלא |  |
| 9:2.442 | על |  |
| 9:2.443 | ידי |  |
| 9:2.444 | ששמענו |  |
| 9:2.445 | שדיבר |  |
| 9:2.446 | עמו | . |
| 9:2.447 | ולאחר |  |
| 9:2.448 | שנתן |  |
| 9:2.449 | את |  |
| 9:2.450 | התורה |  |
| 9:2.451 | אמר |  |
| 9:2.452 | אתם |  |
| 9:2.453 | ראיתם |  |
| 9:2.454 | כי |  |
| 9:2.455 | מן |  |
| 9:2.456 | השמים |  |
| 9:2.457 | דברתי |  |
| 9:2.458 | עמכם | . |
| 9:2.459 | כלומר |  |
| 9:2.460 | אין |  |
| 9:2.461 | לכם |  |
| 9:2.462 | להאמין |  |
| 9:2.463 | עוד |  |
| 9:2.464 | לשום |  |
| 9:2.465 | אדם |  |
| 9:2.466 | שיאמר |  |
| 9:2.467 | שהחלפתי |  |
| 9:2.468 | תורה |  |
| 9:2.469 | אם |  |
| 9:2.470 | לא |  |
| 9:2.471 | תראו |  |
| 9:2.472 | שאדבר |  |
| 9:2.473 | עמכם |  |
| 9:2.474 | מן |  |
| 9:2.475 | השמים |  |
| 9:2.476 | להחליפה | , |
| 9:2.477 | ולפי |  |
| 9:2.478 | שאמרנו |  |
| 9:2.479 | שאין |  |
| 9:2.480 | לנו |  |
| 9:2.481 | תורה |  |
| 9:2.482 | לשום |  |
| 9:2.483 | עכו"ם |  |
| 9:2.484 | או |  |
| 9:2.485 | ישמעאל |  |
| 9:2.486 | באות |  |
| 9:2.487 | או |  |
| 9:2.488 | מופת |  |
| 9:2.489 | שיעשה |  |
| 9:2.490 | להחליף |  |
| 9:2.491 | תורתינו | , |
| 9:2.492 | פי | ' |
| 9:2.493 | בתורה |  |
| 9:2.494 | היאך |  |
| 9:2.495 | יש |  |
| 9:2.496 | יכולת |  |
| 9:2.497 | לעשות |  |
| 9:2.498 | אות |  |
| 9:2.499 | או |  |
| 9:2.500 | מופת | . |
| 9:2.501 | שנאמר |  |
| 9:2.502 | כי |  |
| 9:2.503 | יקום |  |
| 9:2.504 | בקרבך |  |
| 9:2.505 | נביא |  |
| 9:2.506 | או |  |
| 9:2.507 | חולם |  |
| 9:2.508 | וכו | ' |
| 9:2.509 | כי |  |
| 9:2.510 | מנסה |  |
| 9:2.511 | ה | ' |
| 9:2.512 | אלהיכם |  |
| 9:2.513 | אתכם | , |
| 9:2.514 | הרי |  |
| 9:2.515 | פירש |  |
| 9:2.516 | לך |  |
| 9:2.517 | שזה |  |
| 9:2.518 | האות |  |
| 9:2.519 | והמות |  |
| 9:2.520 | אינו |  |
| 9:2.521 | מכח |  |
| 9:2.522 | ע"ז |  |
| 9:2.523 | שלו |  |
| 9:2.524 | שלא |  |
| 9:2.525 | הקב"ה |  |
| 9:2.526 | עושה |  |
| 9:2.527 | זה |  |
| 9:2.528 | האות |  |
| 9:2.529 | לנסות |  |
| 9:2.530 | בו |  |
| 9:2.531 | את |  |
| 9:2.532 | ישראל | , |
| 9:2.533 | ואחר |  |
| 9:2.534 | פטירת |  |
| 9:2.535 | משה |  |
| 9:2.536 | רבינו |  |
| 9:2.537 | עמדו |  |
| 9:2.538 | נביאים |  |
| 9:2.539 | לישראל |  |
| 9:2.540 | ודברו |  |
| 9:2.541 | בחידות |  |
| 9:2.542 | ומשלים | , |
| 9:2.543 | ועל |  |
| 9:2.544 | ידי |  |
| 9:2.545 | כך |  |
| 9:2.546 | אמרו |  |
| 9:2.547 | הגוים |  |
| 9:2.548 | כי |  |
| 9:2.549 | הנביאים |  |
| 9:2.550 | נבאו |  |
| 9:2.551 | בתורה |  |
| 9:2.552 | חדשה |  |
| 9:2.553 | שלהם |  |
| 9:2.554 | לכך |  |
| 9:2.555 | שלח |  |
| 9:2.556 | הקב"ה |  |
| 9:2.557 | לישראל |  |
| 9:2.558 | ביד |  |
| 9:2.559 | מלאכי |  |
| 9:2.560 | שהוא |  |
| 9:2.561 | אחרון |  |
| 9:2.562 | לכל |  |
| 9:2.563 | הנביאים |  |
| 9:2.564 | ואמר |  |
| 9:2.565 | שני |  |
| 9:2.566 | דברים |  |
| 9:2.567 | בסוף |  |
| 9:2.568 | דבריו |  |
| 9:2.569 | לחתום |  |
| 9:2.570 | כל |  |
| 9:2.571 | הנבואות |  |
| 9:2.572 | הנה |  |
| 9:2.573 | אנכי |  |
| 9:2.574 | שלח |  |
| 9:2.575 | לכן |  |
| 9:2.576 | את |  |
| 9:2.577 | אליה |  |
| 9:2.578 | הנביא |  |
| 9:2.579 | וכו | ' |
| 9:2.580 | כלומר |  |
| 9:2.581 | הנני |  |
| 9:2.582 | אחרון |  |
| 9:2.583 | לכל |  |
| 9:2.584 | הנביאים | . |
| 9:2.585 | וממני |  |
| 9:2.586 | עד |  |
| 9:2.587 | שיבוא |  |
| 9:2.588 | אליהו |  |
| 9:2.589 | לא |  |
| 9:2.590 | יעמוד |  |
| 9:2.591 | נביא | , |
| 9:2.592 | והנני |  |
| 9:2.593 | אומר |  |
| 9:2.594 | לכם |  |
| 9:2.595 | בשם |  |
| 9:2.596 | הקב"ה |  |
| 9:2.597 | זכרו |  |
| 9:2.598 | תורת |  |
| 9:2.599 | משה |  |
| 9:2.600 | עבדי |  |
| 9:2.601 | אשר |  |
| 9:2.602 | צויתי |  |
| 9:2.603 | אותו |  |
| 9:2.604 | בחורב |  |
| 9:2.605 | אל |  |
| 9:2.606 | כל |  |
| 9:2.607 | ישראל |  |
| 9:2.608 | חקים |  |
| 9:2.609 | ומשפטים |  |
| 9:2.610 | כלומר |  |
| 9:2.611 | אל |  |
| 9:2.612 | יעלה |  |
| 9:2.613 | על |  |
| 9:2.614 | לבבכם |  |
| 9:2.615 | שהנביאים |  |
| 9:2.616 | שעמדו |  |
| 9:2.617 | לפני |  |
| 9:2.618 | נתנבאו |  |
| 9:2.619 | בחילוף |  |
| 9:2.620 | תורת |  |
| 9:2.621 | משה | , |
| 9:2.622 | כי |  |
| 9:2.623 | כולם |  |
| 9:2.624 | לחזק |  |
| 9:2.625 | תורת |  |
| 9:2.626 | משה |  |
| 9:2.627 | באו |  |
| 9:2.628 | עכ"ל |  |
| 9:2.629 | הסמ"ג |  |
| 9:2.630 | קרוב |  |
| 9:2.631 | לדברי |  |
| 9:2.632 | הרמב"ם |  |
| 9:2.633 | ז"ל | . |
| 9:2.634 | וכן |  |
| 9:2.635 | דרשו |  |
| 9:2.636 | רז"ל |  |
| 9:2.637 | בפסוק |  |
| 9:2.638 | שובי |  |
| 9:2.639 | שובי |  |
| 9:2.640 | השולמית |  |
| 9:2.641 | הובא |  |
| 9:2.642 | במדרש |  |
| 9:2.643 | הגדול |  |
| 9:2.644 | ויש |  |
| 9:2.645 | פרשת |  |
| 9:2.646 | וישלח |  |
| 9:2.647 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

