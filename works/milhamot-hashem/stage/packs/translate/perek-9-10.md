# Translation pack: milhamot-hashem perek-9-10 (Milchamot ha-Shem 9:12–9:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-10`.

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
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
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

- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
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
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- ואין **צריך** לומר → "needless to say".
- **רוע** לב → "wickedness of heart".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
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
- **עכ"ל** → "Thus far his words." (outside the marks).
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


## Milchamot ha-Shem 9:12

Hebrew (Wikisource): קכט) מי לא ישכיל ויבין לומר שהשאלה הנכרית ששאל השואל את החסיד היתה זרה מאד כעין שאלת פלימו לר', ולמה זה יחשוב על החסיד ההוא מחשבות זרות ונכריות כי הוא מזלזל ח"ו בהעסק בדיני הקרבנות וטומאה וטהרה ותרומה וכיוצא? אלא ודאי רוע לב יש בכאן ושנאת חנם קננה בלבם על חכמי האמת המתכוננים במעשה השי"ת בעולמות שברא כדי להכירו מתוך מעשיו כמו שאמר דוד המלך ע"ה כי אראה שמיך וכו' מה רבו מעשיך ה', ונאמר שאו מרום עיניכם וראו מי ברא אלה לא שיראה אותם כשור וכחמור ח"ו אלא שיתבונן בתקופתן ומהלכיהן במסלותם, ומתוכן יכיר גדולת יוצרם וחכמתו כי נפלאת היא עד מאד, כמו שאמר דוד נפלאים מעשיך וכו'. וכל אותן הראיות שמביא הר"ר מנחם ישאם רוח, ואין כאן מקום להאריך להשיב על חטאעתיו בתרמית האמת המעוור את הפתי שהרי ביאר בעל חובת הלבבות דבריו שם בהקדמה שבתחלה ילמוד האדם ויאמין בכל מצות התורה על ידי חכמי התורה והסמך ויסמוך עד דברי קבלתם, ואחר כך ישיב אל דעתו וישמש בשכלו אחר שיעמוד עליו מצד הקבלה שיהא כוללת את מצות התורה ושרשיהן ופרקיהן ותחקור עליו בשכלך עליו עד שיתברר לך אמתתן וידחה השקר, כדכתיב וידעת היום והשבות אל לבבך וכו' ויצא זה מן הכלל דהיינו יחוד השם ללמד על הכלל כולו, כמו שאמרו רז"ל דבר שהיה בכלל ויוצא מן הכלל ללמד לא ללמד על עצמו יצא אלא ללמד על הכלל יצא עכ"ל. אלמלא שהרב בעל החובות מזהיר שכל מצות התורה צריך ללמוד שרשיהן וחילוקיהן כמו שקבלו חכמים ז"ל, ואח"כ יתבונן בהם עד שיתברר לו אמתתן ככל מה דאפשר, ואשר לא יוכל להשיג כפרה אדומה ושעטנז ובשר בחלב וכיוצא בהם מהחקים יסמוך על הקבלה שכל מצות התורה ניתנו לתועלתינו לטוב לנו כל הימים, אע"פ שלא ידענו מעמם ותכלית הטוב הנמשך מהם אע"פ שהשטן מקטרג עליהן, ואומות העולם משיבין עליהן אזליגן בתר רובא שכיון שנגלה לנו התכלית הטוב ברוב המצות כמו שאמרו רז"ל גם באלה החקים אשר לא נדע תכליתן נסמוך על נותן התורה, כמו שאמרו רז"ל חקה חקקתי גזרה גזרתי אין לך רשות להרהר בהן. קל) וצרעת ממארת ומכה טריה פשתה ורבתה בעמנו היא שנאת החכמה והמדע אשר בה ידע אברהם אבינו והכיר את בוראו טרם הגלותו עליו כמו שהעירו רז"ל במדרשיהם שהיה חוקר ומתבונן במעשי השי"ת בשמים ובארץ ובכל הנמצא בהן עד שהשיג שישי לעולם מנהיג ואז נגלה אליו הקב"ה ואמר לו אני בעל הבירה. והמקובלים החדשים שונאי החכמה לא ידעו להפריש בין פילסופיא טמאה של הטעאות ושבושים לפילסופיא הבנויה על השכל הישר אשר חנן אלהים את האדם והיתה שם פילסופיא בפיהם לשמצה לגנות בה את החכמים היודעים חהבין ולהשכיל במעשה השי"ת ולדעת אותו ידיעה ברורה ולתת מופתים ברורים על מציאותו ית' ויחודו והשגחתו הנפלאה על כל ברואיו מקרני רמים ועד ביצי כנים. ויתמהו איש אל רעהו תמיהה קיימת אצלם על דברי החסיד הר' בחיי בעל חובת הלבבות על שקנה מן החכמה אבל זולתם עובדים זולתו. ותמיהתם הלזו שבורה כשבר נבל יוצרים ותמיהתם ותלונתם אל חיקם תשוב איך הם מתלונים על החסיד השלם בעל חובת הלבבות וקבלו דברי הזהר הבא להם מיד העכו"ם שאמר האלהים סבות רבות כנזכר לעיל פעמים רבות ושהיראה שצונו השי"ת היא מאשתו אבל העבודה היא לזעיר אנפין ולא לסבות שקדמו לו היש עבודת זולתו יותר מזה? ואלו נתנו לבם בדעת ותבונה לדברי הרס"ג והרמב"ם וכו' לא עבדו זולתו.

| id | word | punctuation after |
|---|---|---|
| 9:12.1 | קכט | ) |
| 9:12.2 | מי |  |
| 9:12.3 | לא |  |
| 9:12.4 | ישכיל |  |
| 9:12.5 | ויבין |  |
| 9:12.6 | לומר |  |
| 9:12.7 | שהשאלה |  |
| 9:12.8 | הנכרית |  |
| 9:12.9 | ששאל |  |
| 9:12.10 | השואל |  |
| 9:12.11 | את |  |
| 9:12.12 | החסיד |  |
| 9:12.13 | היתה |  |
| 9:12.14 | זרה |  |
| 9:12.15 | מאד |  |
| 9:12.16 | כעין |  |
| 9:12.17 | שאלת |  |
| 9:12.18 | פלימו |  |
| 9:12.19 | לר | ', |
| 9:12.20 | ולמה |  |
| 9:12.21 | זה |  |
| 9:12.22 | יחשוב |  |
| 9:12.23 | על |  |
| 9:12.24 | החסיד |  |
| 9:12.25 | ההוא |  |
| 9:12.26 | מחשבות |  |
| 9:12.27 | זרות |  |
| 9:12.28 | ונכריות |  |
| 9:12.29 | כי |  |
| 9:12.30 | הוא |  |
| 9:12.31 | מזלזל |  |
| 9:12.32 | ח"ו |  |
| 9:12.33 | בהעסק |  |
| 9:12.34 | בדיני |  |
| 9:12.35 | הקרבנות |  |
| 9:12.36 | וטומאה |  |
| 9:12.37 | וטהרה |  |
| 9:12.38 | ותרומה |  |
| 9:12.39 | וכיוצא | ? |
| 9:12.40 | אלא |  |
| 9:12.41 | ודאי |  |
| 9:12.42 | רוע |  |
| 9:12.43 | לב |  |
| 9:12.44 | יש |  |
| 9:12.45 | בכאן |  |
| 9:12.46 | ושנאת |  |
| 9:12.47 | חנם |  |
| 9:12.48 | קננה |  |
| 9:12.49 | בלבם |  |
| 9:12.50 | על |  |
| 9:12.51 | חכמי |  |
| 9:12.52 | האמת |  |
| 9:12.53 | המתכוננים |  |
| 9:12.54 | במעשה |  |
| 9:12.55 | השי"ת |  |
| 9:12.56 | בעולמות |  |
| 9:12.57 | שברא |  |
| 9:12.58 | כדי |  |
| 9:12.59 | להכירו |  |
| 9:12.60 | מתוך |  |
| 9:12.61 | מעשיו |  |
| 9:12.62 | כמו |  |
| 9:12.63 | שאמר |  |
| 9:12.64 | דוד |  |
| 9:12.65 | המלך |  |
| 9:12.66 | ע"ה |  |
| 9:12.67 | כי |  |
| 9:12.68 | אראה |  |
| 9:12.69 | שמיך |  |
| 9:12.70 | וכו | ' |
| 9:12.71 | מה |  |
| 9:12.72 | רבו |  |
| 9:12.73 | מעשיך |  |
| 9:12.74 | ה | ', |
| 9:12.75 | ונאמר |  |
| 9:12.76 | שאו |  |
| 9:12.77 | מרום |  |
| 9:12.78 | עיניכם |  |
| 9:12.79 | וראו |  |
| 9:12.80 | מי |  |
| 9:12.81 | ברא |  |
| 9:12.82 | אלה |  |
| 9:12.83 | לא |  |
| 9:12.84 | שיראה |  |
| 9:12.85 | אותם |  |
| 9:12.86 | כשור |  |
| 9:12.87 | וכחמור |  |
| 9:12.88 | ח"ו |  |
| 9:12.89 | אלא |  |
| 9:12.90 | שיתבונן |  |
| 9:12.91 | בתקופתן |  |
| 9:12.92 | ומהלכיהן |  |
| 9:12.93 | במסלותם | , |
| 9:12.94 | ומתוכן |  |
| 9:12.95 | יכיר |  |
| 9:12.96 | גדולת |  |
| 9:12.97 | יוצרם |  |
| 9:12.98 | וחכמתו |  |
| 9:12.99 | כי |  |
| 9:12.100 | נפלאת |  |
| 9:12.101 | היא |  |
| 9:12.102 | עד |  |
| 9:12.103 | מאד | , |
| 9:12.104 | כמו |  |
| 9:12.105 | שאמר |  |
| 9:12.106 | דוד |  |
| 9:12.107 | נפלאים |  |
| 9:12.108 | מעשיך |  |
| 9:12.109 | וכו | '. |
| 9:12.110 | וכל |  |
| 9:12.111 | אותן |  |
| 9:12.112 | הראיות |  |
| 9:12.113 | שמביא |  |
| 9:12.114 | הר"ר |  |
| 9:12.115 | מנחם |  |
| 9:12.116 | ישאם |  |
| 9:12.117 | רוח | , |
| 9:12.118 | ואין |  |
| 9:12.119 | כאן |  |
| 9:12.120 | מקום |  |
| 9:12.121 | להאריך |  |
| 9:12.122 | להשיב |  |
| 9:12.123 | על |  |
| 9:12.124 | חטאעתיו |  |
| 9:12.125 | בתרמית |  |
| 9:12.126 | האמת |  |
| 9:12.127 | המעוור |  |
| 9:12.128 | את |  |
| 9:12.129 | הפתי |  |
| 9:12.130 | שהרי |  |
| 9:12.131 | ביאר |  |
| 9:12.132 | בעל |  |
| 9:12.133 | חובת |  |
| 9:12.134 | הלבבות |  |
| 9:12.135 | דבריו |  |
| 9:12.136 | שם |  |
| 9:12.137 | בהקדמה |  |
| 9:12.138 | שבתחלה |  |
| 9:12.139 | ילמוד |  |
| 9:12.140 | האדם |  |
| 9:12.141 | ויאמין |  |
| 9:12.142 | בכל |  |
| 9:12.143 | מצות |  |
| 9:12.144 | התורה |  |
| 9:12.145 | על |  |
| 9:12.146 | ידי |  |
| 9:12.147 | חכמי |  |
| 9:12.148 | התורה |  |
| 9:12.149 | והסמך |  |
| 9:12.150 | ויסמוך |  |
| 9:12.151 | עד |  |
| 9:12.152 | דברי |  |
| 9:12.153 | קבלתם | , |
| 9:12.154 | ואחר |  |
| 9:12.155 | כך |  |
| 9:12.156 | ישיב |  |
| 9:12.157 | אל |  |
| 9:12.158 | דעתו |  |
| 9:12.159 | וישמש |  |
| 9:12.160 | בשכלו |  |
| 9:12.161 | אחר |  |
| 9:12.162 | שיעמוד |  |
| 9:12.163 | עליו |  |
| 9:12.164 | מצד |  |
| 9:12.165 | הקבלה |  |
| 9:12.166 | שיהא |  |
| 9:12.167 | כוללת |  |
| 9:12.168 | את |  |
| 9:12.169 | מצות |  |
| 9:12.170 | התורה |  |
| 9:12.171 | ושרשיהן |  |
| 9:12.172 | ופרקיהן |  |
| 9:12.173 | ותחקור |  |
| 9:12.174 | עליו |  |
| 9:12.175 | בשכלך |  |
| 9:12.176 | עליו |  |
| 9:12.177 | עד |  |
| 9:12.178 | שיתברר |  |
| 9:12.179 | לך |  |
| 9:12.180 | אמתתן |  |
| 9:12.181 | וידחה |  |
| 9:12.182 | השקר | , |
| 9:12.183 | כדכתיב |  |
| 9:12.184 | וידעת |  |
| 9:12.185 | היום |  |
| 9:12.186 | והשבות |  |
| 9:12.187 | אל |  |
| 9:12.188 | לבבך |  |
| 9:12.189 | וכו | ' |
| 9:12.190 | ויצא |  |
| 9:12.191 | זה |  |
| 9:12.192 | מן |  |
| 9:12.193 | הכלל |  |
| 9:12.194 | דהיינו |  |
| 9:12.195 | יחוד |  |
| 9:12.196 | השם |  |
| 9:12.197 | ללמד |  |
| 9:12.198 | על |  |
| 9:12.199 | הכלל |  |
| 9:12.200 | כולו | , |
| 9:12.201 | כמו |  |
| 9:12.202 | שאמרו |  |
| 9:12.203 | רז"ל |  |
| 9:12.204 | דבר |  |
| 9:12.205 | שהיה |  |
| 9:12.206 | בכלל |  |
| 9:12.207 | ויוצא |  |
| 9:12.208 | מן |  |
| 9:12.209 | הכלל |  |
| 9:12.210 | ללמד |  |
| 9:12.211 | לא |  |
| 9:12.212 | ללמד |  |
| 9:12.213 | על |  |
| 9:12.214 | עצמו |  |
| 9:12.215 | יצא |  |
| 9:12.216 | אלא |  |
| 9:12.217 | ללמד |  |
| 9:12.218 | על |  |
| 9:12.219 | הכלל |  |
| 9:12.220 | יצא |  |
| 9:12.221 | עכ"ל | . |
| 9:12.222 | אלמלא |  |
| 9:12.223 | שהרב |  |
| 9:12.224 | בעל |  |
| 9:12.225 | החובות |  |
| 9:12.226 | מזהיר |  |
| 9:12.227 | שכל |  |
| 9:12.228 | מצות |  |
| 9:12.229 | התורה |  |
| 9:12.230 | צריך |  |
| 9:12.231 | ללמוד |  |
| 9:12.232 | שרשיהן |  |
| 9:12.233 | וחילוקיהן |  |
| 9:12.234 | כמו |  |
| 9:12.235 | שקבלו |  |
| 9:12.236 | חכמים |  |
| 9:12.237 | ז"ל | , |
| 9:12.238 | ואח"כ |  |
| 9:12.239 | יתבונן |  |
| 9:12.240 | בהם |  |
| 9:12.241 | עד |  |
| 9:12.242 | שיתברר |  |
| 9:12.243 | לו |  |
| 9:12.244 | אמתתן |  |
| 9:12.245 | ככל |  |
| 9:12.246 | מה |  |
| 9:12.247 | דאפשר | , |
| 9:12.248 | ואשר |  |
| 9:12.249 | לא |  |
| 9:12.250 | יוכל |  |
| 9:12.251 | להשיג |  |
| 9:12.252 | כפרה |  |
| 9:12.253 | אדומה |  |
| 9:12.254 | ושעטנז |  |
| 9:12.255 | ובשר |  |
| 9:12.256 | בחלב |  |
| 9:12.257 | וכיוצא |  |
| 9:12.258 | בהם |  |
| 9:12.259 | מהחקים |  |
| 9:12.260 | יסמוך |  |
| 9:12.261 | על |  |
| 9:12.262 | הקבלה |  |
| 9:12.263 | שכל |  |
| 9:12.264 | מצות |  |
| 9:12.265 | התורה |  |
| 9:12.266 | ניתנו |  |
| 9:12.267 | לתועלתינו |  |
| 9:12.268 | לטוב |  |
| 9:12.269 | לנו |  |
| 9:12.270 | כל |  |
| 9:12.271 | הימים | , |
| 9:12.272 | אע"פ |  |
| 9:12.273 | שלא |  |
| 9:12.274 | ידענו |  |
| 9:12.275 | מעמם |  |
| 9:12.276 | ותכלית |  |
| 9:12.277 | הטוב |  |
| 9:12.278 | הנמשך |  |
| 9:12.279 | מהם |  |
| 9:12.280 | אע"פ |  |
| 9:12.281 | שהשטן |  |
| 9:12.282 | מקטרג |  |
| 9:12.283 | עליהן | , |
| 9:12.284 | ואומות |  |
| 9:12.285 | העולם |  |
| 9:12.286 | משיבין |  |
| 9:12.287 | עליהן |  |
| 9:12.288 | אזליגן |  |
| 9:12.289 | בתר |  |
| 9:12.290 | רובא |  |
| 9:12.291 | שכיון |  |
| 9:12.292 | שנגלה |  |
| 9:12.293 | לנו |  |
| 9:12.294 | התכלית |  |
| 9:12.295 | הטוב |  |
| 9:12.296 | ברוב |  |
| 9:12.297 | המצות |  |
| 9:12.298 | כמו |  |
| 9:12.299 | שאמרו |  |
| 9:12.300 | רז"ל |  |
| 9:12.301 | גם |  |
| 9:12.302 | באלה |  |
| 9:12.303 | החקים |  |
| 9:12.304 | אשר |  |
| 9:12.305 | לא |  |
| 9:12.306 | נדע |  |
| 9:12.307 | תכליתן |  |
| 9:12.308 | נסמוך |  |
| 9:12.309 | על |  |
| 9:12.310 | נותן |  |
| 9:12.311 | התורה | , |
| 9:12.312 | כמו |  |
| 9:12.313 | שאמרו |  |
| 9:12.314 | רז"ל |  |
| 9:12.315 | חקה |  |
| 9:12.316 | חקקתי |  |
| 9:12.317 | גזרה |  |
| 9:12.318 | גזרתי |  |
| 9:12.319 | אין |  |
| 9:12.320 | לך |  |
| 9:12.321 | רשות |  |
| 9:12.322 | להרהר |  |
| 9:12.323 | בהן | . |
| 9:12.324 | קל | ) |
| 9:12.325 | וצרעת |  |
| 9:12.326 | ממארת |  |
| 9:12.327 | ומכה |  |
| 9:12.328 | טריה |  |
| 9:12.329 | פשתה |  |
| 9:12.330 | ורבתה |  |
| 9:12.331 | בעמנו |  |
| 9:12.332 | היא |  |
| 9:12.333 | שנאת |  |
| 9:12.334 | החכמה |  |
| 9:12.335 | והמדע |  |
| 9:12.336 | אשר |  |
| 9:12.337 | בה |  |
| 9:12.338 | ידע |  |
| 9:12.339 | אברהם |  |
| 9:12.340 | אבינו |  |
| 9:12.341 | והכיר |  |
| 9:12.342 | את |  |
| 9:12.343 | בוראו |  |
| 9:12.344 | טרם |  |
| 9:12.345 | הגלותו |  |
| 9:12.346 | עליו |  |
| 9:12.347 | כמו |  |
| 9:12.348 | שהעירו |  |
| 9:12.349 | רז"ל |  |
| 9:12.350 | במדרשיהם |  |
| 9:12.351 | שהיה |  |
| 9:12.352 | חוקר |  |
| 9:12.353 | ומתבונן |  |
| 9:12.354 | במעשי |  |
| 9:12.355 | השי"ת |  |
| 9:12.356 | בשמים |  |
| 9:12.357 | ובארץ |  |
| 9:12.358 | ובכל |  |
| 9:12.359 | הנמצא |  |
| 9:12.360 | בהן |  |
| 9:12.361 | עד |  |
| 9:12.362 | שהשיג |  |
| 9:12.363 | שישי |  |
| 9:12.364 | לעולם |  |
| 9:12.365 | מנהיג |  |
| 9:12.366 | ואז |  |
| 9:12.367 | נגלה |  |
| 9:12.368 | אליו |  |
| 9:12.369 | הקב"ה |  |
| 9:12.370 | ואמר |  |
| 9:12.371 | לו |  |
| 9:12.372 | אני |  |
| 9:12.373 | בעל |  |
| 9:12.374 | הבירה | . |
| 9:12.375 | והמקובלים |  |
| 9:12.376 | החדשים |  |
| 9:12.377 | שונאי |  |
| 9:12.378 | החכמה |  |
| 9:12.379 | לא |  |
| 9:12.380 | ידעו |  |
| 9:12.381 | להפריש |  |
| 9:12.382 | בין |  |
| 9:12.383 | פילסופיא |  |
| 9:12.384 | טמאה |  |
| 9:12.385 | של |  |
| 9:12.386 | הטעאות |  |
| 9:12.387 | ושבושים |  |
| 9:12.388 | לפילסופיא |  |
| 9:12.389 | הבנויה |  |
| 9:12.390 | על |  |
| 9:12.391 | השכל |  |
| 9:12.392 | הישר |  |
| 9:12.393 | אשר |  |
| 9:12.394 | חנן |  |
| 9:12.395 | אלהים |  |
| 9:12.396 | את |  |
| 9:12.397 | האדם |  |
| 9:12.398 | והיתה |  |
| 9:12.399 | שם |  |
| 9:12.400 | פילסופיא |  |
| 9:12.401 | בפיהם |  |
| 9:12.402 | לשמצה |  |
| 9:12.403 | לגנות |  |
| 9:12.404 | בה |  |
| 9:12.405 | את |  |
| 9:12.406 | החכמים |  |
| 9:12.407 | היודעים |  |
| 9:12.408 | חהבין |  |
| 9:12.409 | ולהשכיל |  |
| 9:12.410 | במעשה |  |
| 9:12.411 | השי"ת |  |
| 9:12.412 | ולדעת |  |
| 9:12.413 | אותו |  |
| 9:12.414 | ידיעה |  |
| 9:12.415 | ברורה |  |
| 9:12.416 | ולתת |  |
| 9:12.417 | מופתים |  |
| 9:12.418 | ברורים |  |
| 9:12.419 | על |  |
| 9:12.420 | מציאותו |  |
| 9:12.421 | ית | ' |
| 9:12.422 | ויחודו |  |
| 9:12.423 | והשגחתו |  |
| 9:12.424 | הנפלאה |  |
| 9:12.425 | על |  |
| 9:12.426 | כל |  |
| 9:12.427 | ברואיו |  |
| 9:12.428 | מקרני |  |
| 9:12.429 | רמים |  |
| 9:12.430 | ועד |  |
| 9:12.431 | ביצי |  |
| 9:12.432 | כנים | . |
| 9:12.433 | ויתמהו |  |
| 9:12.434 | איש |  |
| 9:12.435 | אל |  |
| 9:12.436 | רעהו |  |
| 9:12.437 | תמיהה |  |
| 9:12.438 | קיימת |  |
| 9:12.439 | אצלם |  |
| 9:12.440 | על |  |
| 9:12.441 | דברי |  |
| 9:12.442 | החסיד |  |
| 9:12.443 | הר | ' |
| 9:12.444 | בחיי |  |
| 9:12.445 | בעל |  |
| 9:12.446 | חובת |  |
| 9:12.447 | הלבבות |  |
| 9:12.448 | על |  |
| 9:12.449 | שקנה |  |
| 9:12.450 | מן |  |
| 9:12.451 | החכמה |  |
| 9:12.452 | אבל |  |
| 9:12.453 | זולתם |  |
| 9:12.454 | עובדים |  |
| 9:12.455 | זולתו | . |
| 9:12.456 | ותמיהתם |  |
| 9:12.457 | הלזו |  |
| 9:12.458 | שבורה |  |
| 9:12.459 | כשבר |  |
| 9:12.460 | נבל |  |
| 9:12.461 | יוצרים |  |
| 9:12.462 | ותמיהתם |  |
| 9:12.463 | ותלונתם |  |
| 9:12.464 | אל |  |
| 9:12.465 | חיקם |  |
| 9:12.466 | תשוב |  |
| 9:12.467 | איך |  |
| 9:12.468 | הם |  |
| 9:12.469 | מתלונים |  |
| 9:12.470 | על |  |
| 9:12.471 | החסיד |  |
| 9:12.472 | השלם |  |
| 9:12.473 | בעל |  |
| 9:12.474 | חובת |  |
| 9:12.475 | הלבבות |  |
| 9:12.476 | וקבלו |  |
| 9:12.477 | דברי |  |
| 9:12.478 | הזהר |  |
| 9:12.479 | הבא |  |
| 9:12.480 | להם |  |
| 9:12.481 | מיד |  |
| 9:12.482 | העכו"ם |  |
| 9:12.483 | שאמר |  |
| 9:12.484 | האלהים |  |
| 9:12.485 | סבות |  |
| 9:12.486 | רבות |  |
| 9:12.487 | כנזכר |  |
| 9:12.488 | לעיל |  |
| 9:12.489 | פעמים |  |
| 9:12.490 | רבות |  |
| 9:12.491 | ושהיראה |  |
| 9:12.492 | שצונו |  |
| 9:12.493 | השי"ת |  |
| 9:12.494 | היא |  |
| 9:12.495 | מאשתו |  |
| 9:12.496 | אבל |  |
| 9:12.497 | העבודה |  |
| 9:12.498 | היא |  |
| 9:12.499 | לזעיר |  |
| 9:12.500 | אנפין |  |
| 9:12.501 | ולא |  |
| 9:12.502 | לסבות |  |
| 9:12.503 | שקדמו |  |
| 9:12.504 | לו |  |
| 9:12.505 | היש |  |
| 9:12.506 | עבודת |  |
| 9:12.507 | זולתו |  |
| 9:12.508 | יותר |  |
| 9:12.509 | מזה | ? |
| 9:12.510 | ואלו |  |
| 9:12.511 | נתנו |  |
| 9:12.512 | לבם |  |
| 9:12.513 | בדעת |  |
| 9:12.514 | ותבונה |  |
| 9:12.515 | לדברי |  |
| 9:12.516 | הרס"ג |  |
| 9:12.517 | והרמב"ם |  |
| 9:12.518 | וכו | ' |
| 9:12.519 | לא |  |
| 9:12.520 | עבדו |  |
| 9:12.521 | זולתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

