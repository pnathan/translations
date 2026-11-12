# Translation pack: milhamot-hashem perek-2-08 (Milchamot ha-Shem 2:16–2:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-08`.

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
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
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
- **רז"ל** → "our Rabbis of blessed memory".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **העיקרים** → "the author of the Ikkarim".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בא"י** → "in the Land of Israel".

## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
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
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **עלה / עלול** → "cause / effect".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ת"ל** → "Scripture says".
- **ביהכנ"ס** → "synagogue".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **אנדרטא** → "statue".
- אשם **תלוי** → "suspended guilt offering".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין** אל → "call upon".
- אחד ... **ואחד** of two kinds of case → "alike".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ווי** ליה מאן דלא → "woe to him who does not".
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **ע"א / ע"ב** → "a / b".
- **מיתות**: פ' ד' מיתות (פ' + a chapter title) → "Four Death Penalties".
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
- **סנהדרין** → "Sanhedrin 99"; סנהדרין רי"ף → "Sanhedrin, [in] the Rif".
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


## Milchamot ha-Shem 2:16

Hebrew (Wikisource): ומה שכתבת: ויותר מזה מה שכ' על כבוד ה' שהוא אור נברא אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחוים לנברא? לא לכבוד הנראה משתחוים, רק למי שהשרה כבודו שם. והכבוד הנראה שם לאות וסימן להם, כי השי"ת איווה לאוהל מועד למושב לו, כמו שאנחנו משתחוים ומתפללים לשי"ת נכח המקדש וירושלם, כי בחר ה' בציון אוה למושב לו. וכבר ביאר הרמב"ם בפ' ס"ד מראשון, כי כבוד ה', פעמים רוצים בו האור הנברא אשר ישכינהו הק' במקום, כמו: וישכון כבוד ה' על הר סיני ויכסהו וכו', ופעמים רוצים בו עצמותו ית' ואמתתו, כמו "הראיני נא את כבודך", ובא המענה: "כי לא יראני האדם וחי", מורה כי הכבוד הנאמר פה עצמו ית', ואמר "כבודך" להגדיל וכו'. וכיוצא בזה מצינו בדברי רז"ל ששם אלהים משותף לדיינים ולמלאכים ולשי"ת, גם שם שכינה משותף למה שהזכרתי למעלה. וראיתי להעירך פה שאין איסור להשתחוות לאדם כל שאין דעתו לקבלו באלוה וגם אינו נעבד, רק נכנע לו ומשפיל עצמו לפניו, או דרך ההודאה על טובה שעשה לו. שהרי אברהם אבינו השתחוה לבני חת על החסד שעשו עמו ליתן לו אחזת קבר. וכן תרגם רבי סעדיה גאון: וישתחו לעם הארץ לבני חת, וסגד לאהל אלבלד שברא. ויהושוע השתחווה למלאך. ונתן הנביא ובת שבע השתחוו לדוד. והרבה מצינו בכתובים שהשתחוו לנברא. וגרסינן בסנהדרין פ' ארבע מיתות (סנהדרין ס"א ע"ב), אתמר: העובד ע"ז מאהבה ומיראה, אביי אמר חייב, דהא פלחה. רבא אמר פטור, אי קבליה עליה באלוה אין, אי לא לא. ואמר אביי מנא אמינא לה, דתנן העובד ע"ז אחד העובד ואחד המזבח ואחד המקטיר ואחד המנסך ואחד המשתחוה ואחד המקבלו עליו באלוה והאומר לו אלי אתה. מאי לאו אחד העובד מאהבה או מיראה? ורבא אמר לך, לא, אלא כדמתרץ ר' ירמיה (פירש רש"י: אחד העובד כל עבודה שהיא דרכה בכך, ואחד המזבח וכו' שלא כדרכה, ולעולם רישא וסיפא לשם אלהות). אמר אביי: מנא אמינא לה? דתניא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך. יכול אפי' לעבד כהמן? ת"ל ולא תעבדם. והא המן מיראה הוא נעבד! ורבא? כהמן ולא כהמן. כהמן, דאיהו גופיה ע"ז; ולא כהמן, דאלו המן מיראה, והכא לאו מיראה. ואמר אביי: מנא אמינא לה? דתניא: כהן משיח בע"ז. ר' אומר בשגגת מעשה, וחכמים אומרים בהעלם דבר, ושווין דבשעירה כיחיד, ושווין שאין מביא אשם תלוי. האי שגגת מעשה היכי דמי? אי קסבר ביהכנ"ס והשתחוה לו, הרי לבו לשמים. אלא דחזא אנדרטא והשתחוה לו (פירוש, שראה דמות המלך שהיו רגילים לעשות כדמות המלך והרואה אותו משתחוה לו לכבוד המלך, ופעמים שעובדין אותו, וזה ראה אחד שנעבד והשתחוה לו), אי קבליה עליה באלוה מזיד הוא, ואי לא קבליה עליה באלוה לאו כלום הוא (פירש רש"י: אלא לכבוד המלך השתחוה לו, לאו כלום הוא דהא לא איכוון לע"ז). אלא לאו מאהבה ומיראה? ורבא אמר לך: באומר מותר. אומר מותר היינו העלם דבר? באומר מותר לגמרי, העלם דבר קיום מקצת וביטול מקצת, עכ"ל הגמ' עם קצת פי' רש"י. ודבר נלמד, דלכל נברא אסור להשתחוות, זולת האדם, והוא שלא נעבד כהמן שעשה עצמו ע"ז, כדאמירן במגלה (יט), דאי לאו הכי לא היה מרדכי נמנע מלהשתחות לו, כדקתני בבריתא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך, והוא שלא יהא נעבד כהמן שעשה עצמו אלוה, רק משתחוה לו לאות הכבוד ולא לאלהות, ולכך השתחוה יהושע למלאך הואיל ונגלה לו בדמות אדם, כמ"ש: והנה איש עומד לנגדו, ולכן השתחוה לו לאות הכבוד ולא לעבדו. ועיין בס' העיקרים מאמר ב' פ' כ"ח שהאריך בביאור הכתובים שקורין לכבוד הנראה לנביאים בשם הוי"ה, וגם לארון, כמו שגאמר: "ויהי בנסע הארון וכו' קומה ה', ובנחה יאמר שובה ה'", ויהושע קרא לארון "אדון כל הארץ" ע"ש. ולדעתו שמה שהשתחוה יהושע למלאך, לפי שההשתחויה ההיא היתה לכבוד השי"ת, ולהיות זה שלוחו וכו' ע"ש. ורז"ל אמרו תעלא בעידניה סגיד ליה, ר"ל: אדם פחות ובזוי כשועל, והשעה עומדת לו באיזה משרה שבידו, סגיד ליה. וכן הוא נהוג אצלנו ואצל הערביים להשתחוות ולנשק אדם גדול המעלה, ואין בזה שום איסור, שאינו מקבלו עליו באלוה, רק לאות כבוד הוא משתחוה לו ומנשק את רגליו.

| id | word | punctuation after |
|---|---|---|
| 2:16.1 | ומה |  |
| 2:16.2 | שכתבת | : |
| 2:16.3 | ויותר |  |
| 2:16.4 | מזה |  |
| 2:16.5 | מה |  |
| 2:16.6 | שכ | ' |
| 2:16.7 | על |  |
| 2:16.8 | כבוד |  |
| 2:16.9 | ה | ' |
| 2:16.10 | שהוא |  |
| 2:16.11 | אור |  |
| 2:16.12 | נברא |  |
| 2:16.13 | אם |  |
| 2:16.14 | כן |  |
| 2:16.15 | כיצד |  |
| 2:16.16 | היו |  |
| 2:16.17 | ישראל |  |
| 2:16.18 | יוצאים |  |
| 2:16.19 | פתח |  |
| 2:16.20 | אהליהם |  |
| 2:16.21 | ומשתחוים |  |
| 2:16.22 | לנברא | ? |
| 2:16.23 | לא |  |
| 2:16.24 | לכבוד |  |
| 2:16.25 | הנראה |  |
| 2:16.26 | משתחוים | , |
| 2:16.27 | רק |  |
| 2:16.28 | למי |  |
| 2:16.29 | שהשרה |  |
| 2:16.30 | כבודו |  |
| 2:16.31 | שם | . |
| 2:16.32 | והכבוד |  |
| 2:16.33 | הנראה |  |
| 2:16.34 | שם |  |
| 2:16.35 | לאות |  |
| 2:16.36 | וסימן |  |
| 2:16.37 | להם | , |
| 2:16.38 | כי |  |
| 2:16.39 | השי"ת |  |
| 2:16.40 | איווה |  |
| 2:16.41 | לאוהל |  |
| 2:16.42 | מועד |  |
| 2:16.43 | למושב |  |
| 2:16.44 | לו | , |
| 2:16.45 | כמו |  |
| 2:16.46 | שאנחנו |  |
| 2:16.47 | משתחוים |  |
| 2:16.48 | ומתפללים |  |
| 2:16.49 | לשי"ת |  |
| 2:16.50 | נכח |  |
| 2:16.51 | המקדש |  |
| 2:16.52 | וירושלם | , |
| 2:16.53 | כי |  |
| 2:16.54 | בחר |  |
| 2:16.55 | ה | ' |
| 2:16.56 | בציון |  |
| 2:16.57 | אוה |  |
| 2:16.58 | למושב |  |
| 2:16.59 | לו | . |
| 2:16.60 | וכבר |  |
| 2:16.61 | ביאר |  |
| 2:16.62 | הרמב"ם |  |
| 2:16.63 | בפ | ' |
| 2:16.64 | ס"ד |  |
| 2:16.65 | מראשון | , |
| 2:16.66 | כי |  |
| 2:16.67 | כבוד |  |
| 2:16.68 | ה | ', |
| 2:16.69 | פעמים |  |
| 2:16.70 | רוצים |  |
| 2:16.71 | בו |  |
| 2:16.72 | האור |  |
| 2:16.73 | הנברא |  |
| 2:16.74 | אשר |  |
| 2:16.75 | ישכינהו |  |
| 2:16.76 | הק | ' |
| 2:16.77 | במקום | , |
| 2:16.78 | כמו | : |
| 2:16.79 | וישכון |  |
| 2:16.80 | כבוד |  |
| 2:16.81 | ה | ' |
| 2:16.82 | על |  |
| 2:16.83 | הר |  |
| 2:16.84 | סיני |  |
| 2:16.85 | ויכסהו |  |
| 2:16.86 | וכו | ', |
| 2:16.87 | ופעמים |  |
| 2:16.88 | רוצים |  |
| 2:16.89 | בו |  |
| 2:16.90 | עצמותו |  |
| 2:16.91 | ית | ' |
| 2:16.92 | ואמתתו | , |
| 2:16.93 | כמו |  |
| 2:16.94 | הראיני |  |
| 2:16.95 | נא |  |
| 2:16.96 | את |  |
| 2:16.97 | כבודך | ", |
| 2:16.98 | ובא |  |
| 2:16.99 | המענה | : |
| 2:16.100 | כי |  |
| 2:16.101 | לא |  |
| 2:16.102 | יראני |  |
| 2:16.103 | האדם |  |
| 2:16.104 | וחי | ", |
| 2:16.105 | מורה |  |
| 2:16.106 | כי |  |
| 2:16.107 | הכבוד |  |
| 2:16.108 | הנאמר |  |
| 2:16.109 | פה |  |
| 2:16.110 | עצמו |  |
| 2:16.111 | ית | ', |
| 2:16.112 | ואמר |  |
| 2:16.113 | כבודך | " |
| 2:16.114 | להגדיל |  |
| 2:16.115 | וכו | '. |
| 2:16.116 | וכיוצא |  |
| 2:16.117 | בזה |  |
| 2:16.118 | מצינו |  |
| 2:16.119 | בדברי |  |
| 2:16.120 | רז"ל |  |
| 2:16.121 | ששם |  |
| 2:16.122 | אלהים |  |
| 2:16.123 | משותף |  |
| 2:16.124 | לדיינים |  |
| 2:16.125 | ולמלאכים |  |
| 2:16.126 | ולשי"ת | , |
| 2:16.127 | גם |  |
| 2:16.128 | שם |  |
| 2:16.129 | שכינה |  |
| 2:16.130 | משותף |  |
| 2:16.131 | למה |  |
| 2:16.132 | שהזכרתי |  |
| 2:16.133 | למעלה | . |
| 2:16.134 | וראיתי |  |
| 2:16.135 | להעירך |  |
| 2:16.136 | פה |  |
| 2:16.137 | שאין |  |
| 2:16.138 | איסור |  |
| 2:16.139 | להשתחוות |  |
| 2:16.140 | לאדם |  |
| 2:16.141 | כל |  |
| 2:16.142 | שאין |  |
| 2:16.143 | דעתו |  |
| 2:16.144 | לקבלו |  |
| 2:16.145 | באלוה |  |
| 2:16.146 | וגם |  |
| 2:16.147 | אינו |  |
| 2:16.148 | נעבד | , |
| 2:16.149 | רק |  |
| 2:16.150 | נכנע |  |
| 2:16.151 | לו |  |
| 2:16.152 | ומשפיל |  |
| 2:16.153 | עצמו |  |
| 2:16.154 | לפניו | , |
| 2:16.155 | או |  |
| 2:16.156 | דרך |  |
| 2:16.157 | ההודאה |  |
| 2:16.158 | על |  |
| 2:16.159 | טובה |  |
| 2:16.160 | שעשה |  |
| 2:16.161 | לו | . |
| 2:16.162 | שהרי |  |
| 2:16.163 | אברהם |  |
| 2:16.164 | אבינו |  |
| 2:16.165 | השתחוה |  |
| 2:16.166 | לבני |  |
| 2:16.167 | חת |  |
| 2:16.168 | על |  |
| 2:16.169 | החסד |  |
| 2:16.170 | שעשו |  |
| 2:16.171 | עמו |  |
| 2:16.172 | ליתן |  |
| 2:16.173 | לו |  |
| 2:16.174 | אחזת |  |
| 2:16.175 | קבר | . |
| 2:16.176 | וכן |  |
| 2:16.177 | תרגם |  |
| 2:16.178 | רבי |  |
| 2:16.179 | סעדיה |  |
| 2:16.180 | גאון | : |
| 2:16.181 | וישתחו |  |
| 2:16.182 | לעם |  |
| 2:16.183 | הארץ |  |
| 2:16.184 | לבני |  |
| 2:16.185 | חת | , |
| 2:16.186 | וסגד |  |
| 2:16.187 | לאהל |  |
| 2:16.188 | אלבלד |  |
| 2:16.189 | שברא | . |
| 2:16.190 | ויהושוע |  |
| 2:16.191 | השתחווה |  |
| 2:16.192 | למלאך | . |
| 2:16.193 | ונתן |  |
| 2:16.194 | הנביא |  |
| 2:16.195 | ובת |  |
| 2:16.196 | שבע |  |
| 2:16.197 | השתחוו |  |
| 2:16.198 | לדוד | . |
| 2:16.199 | והרבה |  |
| 2:16.200 | מצינו |  |
| 2:16.201 | בכתובים |  |
| 2:16.202 | שהשתחוו |  |
| 2:16.203 | לנברא | . |
| 2:16.204 | וגרסינן |  |
| 2:16.205 | בסנהדרין |  |
| 2:16.206 | פ | ' |
| 2:16.207 | ארבע |  |
| 2:16.208 | מיתות |  |
| 2:16.209 | סנהדרין |  |
| 2:16.210 | ס"א |  |
| 2:16.211 | ע"ב | ), |
| 2:16.212 | אתמר | : |
| 2:16.213 | העובד |  |
| 2:16.214 | ע"ז |  |
| 2:16.215 | מאהבה |  |
| 2:16.216 | ומיראה | , |
| 2:16.217 | אביי |  |
| 2:16.218 | אמר |  |
| 2:16.219 | חייב | , |
| 2:16.220 | דהא |  |
| 2:16.221 | פלחה | . |
| 2:16.222 | רבא |  |
| 2:16.223 | אמר |  |
| 2:16.224 | פטור | , |
| 2:16.225 | אי |  |
| 2:16.226 | קבליה |  |
| 2:16.227 | עליה |  |
| 2:16.228 | באלוה |  |
| 2:16.229 | אין | , |
| 2:16.230 | אי |  |
| 2:16.231 | לא |  |
| 2:16.232 | לא | . |
| 2:16.233 | ואמר |  |
| 2:16.234 | אביי |  |
| 2:16.235 | מנא |  |
| 2:16.236 | אמינא |  |
| 2:16.237 | לה | , |
| 2:16.238 | דתנן |  |
| 2:16.239 | העובד |  |
| 2:16.240 | ע"ז |  |
| 2:16.241 | אחד |  |
| 2:16.242 | העובד |  |
| 2:16.243 | ואחד |  |
| 2:16.244 | המזבח |  |
| 2:16.245 | ואחד |  |
| 2:16.246 | המקטיר |  |
| 2:16.247 | ואחד |  |
| 2:16.248 | המנסך |  |
| 2:16.249 | ואחד |  |
| 2:16.250 | המשתחוה |  |
| 2:16.251 | ואחד |  |
| 2:16.252 | המקבלו |  |
| 2:16.253 | עליו |  |
| 2:16.254 | באלוה |  |
| 2:16.255 | והאומר |  |
| 2:16.256 | לו |  |
| 2:16.257 | אלי |  |
| 2:16.258 | אתה | . |
| 2:16.259 | מאי |  |
| 2:16.260 | לאו |  |
| 2:16.261 | אחד |  |
| 2:16.262 | העובד |  |
| 2:16.263 | מאהבה |  |
| 2:16.264 | או |  |
| 2:16.265 | מיראה | ? |
| 2:16.266 | ורבא |  |
| 2:16.267 | אמר |  |
| 2:16.268 | לך | , |
| 2:16.269 | לא | , |
| 2:16.270 | אלא |  |
| 2:16.271 | כדמתרץ |  |
| 2:16.272 | ר | ' |
| 2:16.273 | ירמיה |  |
| 2:16.274 | פירש |  |
| 2:16.275 | רש"י | : |
| 2:16.276 | אחד |  |
| 2:16.277 | העובד |  |
| 2:16.278 | כל |  |
| 2:16.279 | עבודה |  |
| 2:16.280 | שהיא |  |
| 2:16.281 | דרכה |  |
| 2:16.282 | בכך | , |
| 2:16.283 | ואחד |  |
| 2:16.284 | המזבח |  |
| 2:16.285 | וכו | ' |
| 2:16.286 | שלא |  |
| 2:16.287 | כדרכה | , |
| 2:16.288 | ולעולם |  |
| 2:16.289 | רישא |  |
| 2:16.290 | וסיפא |  |
| 2:16.291 | לשם |  |
| 2:16.292 | אלהות | ). |
| 2:16.293 | אמר |  |
| 2:16.294 | אביי | : |
| 2:16.295 | מנא |  |
| 2:16.296 | אמינא |  |
| 2:16.297 | לה | ? |
| 2:16.298 | דתניא | : |
| 2:16.299 | לא |  |
| 2:16.300 | תשתחוה |  |
| 2:16.301 | להם | , |
| 2:16.302 | להם |  |
| 2:16.303 | אי |  |
| 2:16.304 | אתה |  |
| 2:16.305 | משתחוה | , |
| 2:16.306 | אבל |  |
| 2:16.307 | אתה |  |
| 2:16.308 | משתחוה |  |
| 2:16.309 | לאדם |  |
| 2:16.310 | כמותך | . |
| 2:16.311 | יכול |  |
| 2:16.312 | אפי | ' |
| 2:16.313 | לעבד |  |
| 2:16.314 | כהמן | ? |
| 2:16.315 | ת"ל |  |
| 2:16.316 | ולא |  |
| 2:16.317 | תעבדם | . |
| 2:16.318 | והא |  |
| 2:16.319 | המן |  |
| 2:16.320 | מיראה |  |
| 2:16.321 | הוא |  |
| 2:16.322 | נעבד | ! |
| 2:16.323 | ורבא | ? |
| 2:16.324 | כהמן |  |
| 2:16.325 | ולא |  |
| 2:16.326 | כהמן | . |
| 2:16.327 | כהמן | , |
| 2:16.328 | דאיהו |  |
| 2:16.329 | גופיה |  |
| 2:16.330 | ע"ז | ; |
| 2:16.331 | ולא |  |
| 2:16.332 | כהמן | , |
| 2:16.333 | דאלו |  |
| 2:16.334 | המן |  |
| 2:16.335 | מיראה | , |
| 2:16.336 | והכא |  |
| 2:16.337 | לאו |  |
| 2:16.338 | מיראה | . |
| 2:16.339 | ואמר |  |
| 2:16.340 | אביי | : |
| 2:16.341 | מנא |  |
| 2:16.342 | אמינא |  |
| 2:16.343 | לה | ? |
| 2:16.344 | דתניא | : |
| 2:16.345 | כהן |  |
| 2:16.346 | משיח |  |
| 2:16.347 | בע"ז | . |
| 2:16.348 | ר | ' |
| 2:16.349 | אומר |  |
| 2:16.350 | בשגגת |  |
| 2:16.351 | מעשה | , |
| 2:16.352 | וחכמים |  |
| 2:16.353 | אומרים |  |
| 2:16.354 | בהעלם |  |
| 2:16.355 | דבר | , |
| 2:16.356 | ושווין |  |
| 2:16.357 | דבשעירה |  |
| 2:16.358 | כיחיד | , |
| 2:16.359 | ושווין |  |
| 2:16.360 | שאין |  |
| 2:16.361 | מביא |  |
| 2:16.362 | אשם |  |
| 2:16.363 | תלוי | . |
| 2:16.364 | האי |  |
| 2:16.365 | שגגת |  |
| 2:16.366 | מעשה |  |
| 2:16.367 | היכי |  |
| 2:16.368 | דמי | ? |
| 2:16.369 | אי |  |
| 2:16.370 | קסבר |  |
| 2:16.371 | ביהכנ"ס |  |
| 2:16.372 | והשתחוה |  |
| 2:16.373 | לו | , |
| 2:16.374 | הרי |  |
| 2:16.375 | לבו |  |
| 2:16.376 | לשמים | . |
| 2:16.377 | אלא |  |
| 2:16.378 | דחזא |  |
| 2:16.379 | אנדרטא |  |
| 2:16.380 | והשתחוה |  |
| 2:16.381 | לו |  |
| 2:16.382 | פירוש | , |
| 2:16.383 | שראה |  |
| 2:16.384 | דמות |  |
| 2:16.385 | המלך |  |
| 2:16.386 | שהיו |  |
| 2:16.387 | רגילים |  |
| 2:16.388 | לעשות |  |
| 2:16.389 | כדמות |  |
| 2:16.390 | המלך |  |
| 2:16.391 | והרואה |  |
| 2:16.392 | אותו |  |
| 2:16.393 | משתחוה |  |
| 2:16.394 | לו |  |
| 2:16.395 | לכבוד |  |
| 2:16.396 | המלך | , |
| 2:16.397 | ופעמים |  |
| 2:16.398 | שעובדין |  |
| 2:16.399 | אותו | , |
| 2:16.400 | וזה |  |
| 2:16.401 | ראה |  |
| 2:16.402 | אחד |  |
| 2:16.403 | שנעבד |  |
| 2:16.404 | והשתחוה |  |
| 2:16.405 | לו | ), |
| 2:16.406 | אי |  |
| 2:16.407 | קבליה |  |
| 2:16.408 | עליה |  |
| 2:16.409 | באלוה |  |
| 2:16.410 | מזיד |  |
| 2:16.411 | הוא | , |
| 2:16.412 | ואי |  |
| 2:16.413 | לא |  |
| 2:16.414 | קבליה |  |
| 2:16.415 | עליה |  |
| 2:16.416 | באלוה |  |
| 2:16.417 | לאו |  |
| 2:16.418 | כלום |  |
| 2:16.419 | הוא |  |
| 2:16.420 | פירש |  |
| 2:16.421 | רש"י | : |
| 2:16.422 | אלא |  |
| 2:16.423 | לכבוד |  |
| 2:16.424 | המלך |  |
| 2:16.425 | השתחוה |  |
| 2:16.426 | לו | , |
| 2:16.427 | לאו |  |
| 2:16.428 | כלום |  |
| 2:16.429 | הוא |  |
| 2:16.430 | דהא |  |
| 2:16.431 | לא |  |
| 2:16.432 | איכוון |  |
| 2:16.433 | לע"ז | ). |
| 2:16.434 | אלא |  |
| 2:16.435 | לאו |  |
| 2:16.436 | מאהבה |  |
| 2:16.437 | ומיראה | ? |
| 2:16.438 | ורבא |  |
| 2:16.439 | אמר |  |
| 2:16.440 | לך | : |
| 2:16.441 | באומר |  |
| 2:16.442 | מותר | . |
| 2:16.443 | אומר |  |
| 2:16.444 | מותר |  |
| 2:16.445 | היינו |  |
| 2:16.446 | העלם |  |
| 2:16.447 | דבר | ? |
| 2:16.448 | באומר |  |
| 2:16.449 | מותר |  |
| 2:16.450 | לגמרי | , |
| 2:16.451 | העלם |  |
| 2:16.452 | דבר |  |
| 2:16.453 | קיום |  |
| 2:16.454 | מקצת |  |
| 2:16.455 | וביטול |  |
| 2:16.456 | מקצת | , |
| 2:16.457 | עכ"ל |  |
| 2:16.458 | הגמ | ' |
| 2:16.459 | עם |  |
| 2:16.460 | קצת |  |
| 2:16.461 | פי | ' |
| 2:16.462 | רש"י | . |
| 2:16.463 | ודבר |  |
| 2:16.464 | נלמד | , |
| 2:16.465 | דלכל |  |
| 2:16.466 | נברא |  |
| 2:16.467 | אסור |  |
| 2:16.468 | להשתחוות | , |
| 2:16.469 | זולת |  |
| 2:16.470 | האדם | , |
| 2:16.471 | והוא |  |
| 2:16.472 | שלא |  |
| 2:16.473 | נעבד |  |
| 2:16.474 | כהמן |  |
| 2:16.475 | שעשה |  |
| 2:16.476 | עצמו |  |
| 2:16.477 | ע"ז | , |
| 2:16.478 | כדאמירן |  |
| 2:16.479 | במגלה |  |
| 2:16.480 | יט | ), |
| 2:16.481 | דאי |  |
| 2:16.482 | לאו |  |
| 2:16.483 | הכי |  |
| 2:16.484 | לא |  |
| 2:16.485 | היה |  |
| 2:16.486 | מרדכי |  |
| 2:16.487 | נמנע |  |
| 2:16.488 | מלהשתחות |  |
| 2:16.489 | לו | , |
| 2:16.490 | כדקתני |  |
| 2:16.491 | בבריתא | : |
| 2:16.492 | לא |  |
| 2:16.493 | תשתחוה |  |
| 2:16.494 | להם | , |
| 2:16.495 | להם |  |
| 2:16.496 | אי |  |
| 2:16.497 | אתה |  |
| 2:16.498 | משתחוה | , |
| 2:16.499 | אבל |  |
| 2:16.500 | אתה |  |
| 2:16.501 | משתחוה |  |
| 2:16.502 | לאדם |  |
| 2:16.503 | כמותך | , |
| 2:16.504 | והוא |  |
| 2:16.505 | שלא |  |
| 2:16.506 | יהא |  |
| 2:16.507 | נעבד |  |
| 2:16.508 | כהמן |  |
| 2:16.509 | שעשה |  |
| 2:16.510 | עצמו |  |
| 2:16.511 | אלוה | , |
| 2:16.512 | רק |  |
| 2:16.513 | משתחוה |  |
| 2:16.514 | לו |  |
| 2:16.515 | לאות |  |
| 2:16.516 | הכבוד |  |
| 2:16.517 | ולא |  |
| 2:16.518 | לאלהות | , |
| 2:16.519 | ולכך |  |
| 2:16.520 | השתחוה |  |
| 2:16.521 | יהושע |  |
| 2:16.522 | למלאך |  |
| 2:16.523 | הואיל |  |
| 2:16.524 | ונגלה |  |
| 2:16.525 | לו |  |
| 2:16.526 | בדמות |  |
| 2:16.527 | אדם | , |
| 2:16.528 | כמ"ש | : |
| 2:16.529 | והנה |  |
| 2:16.530 | איש |  |
| 2:16.531 | עומד |  |
| 2:16.532 | לנגדו | , |
| 2:16.533 | ולכן |  |
| 2:16.534 | השתחוה |  |
| 2:16.535 | לו |  |
| 2:16.536 | לאות |  |
| 2:16.537 | הכבוד |  |
| 2:16.538 | ולא |  |
| 2:16.539 | לעבדו | . |
| 2:16.540 | ועיין |  |
| 2:16.541 | בס | ' |
| 2:16.542 | העיקרים |  |
| 2:16.543 | מאמר |  |
| 2:16.544 | ב | ' |
| 2:16.545 | פ | ' |
| 2:16.546 | כ"ח |  |
| 2:16.547 | שהאריך |  |
| 2:16.548 | בביאור |  |
| 2:16.549 | הכתובים |  |
| 2:16.550 | שקורין |  |
| 2:16.551 | לכבוד |  |
| 2:16.552 | הנראה |  |
| 2:16.553 | לנביאים |  |
| 2:16.554 | בשם |  |
| 2:16.555 | הוי"ה | , |
| 2:16.556 | וגם |  |
| 2:16.557 | לארון | , |
| 2:16.558 | כמו |  |
| 2:16.559 | שגאמר | : |
| 2:16.560 | ויהי |  |
| 2:16.561 | בנסע |  |
| 2:16.562 | הארון |  |
| 2:16.563 | וכו | ' |
| 2:16.564 | קומה |  |
| 2:16.565 | ה | ', |
| 2:16.566 | ובנחה |  |
| 2:16.567 | יאמר |  |
| 2:16.568 | שובה |  |
| 2:16.569 | ה | '", |
| 2:16.570 | ויהושע |  |
| 2:16.571 | קרא |  |
| 2:16.572 | לארון |  |
| 2:16.573 | אדון |  |
| 2:16.574 | כל |  |
| 2:16.575 | הארץ | " |
| 2:16.576 | ע"ש | . |
| 2:16.577 | ולדעתו |  |
| 2:16.578 | שמה |  |
| 2:16.579 | שהשתחוה |  |
| 2:16.580 | יהושע |  |
| 2:16.581 | למלאך | , |
| 2:16.582 | לפי |  |
| 2:16.583 | שההשתחויה |  |
| 2:16.584 | ההיא |  |
| 2:16.585 | היתה |  |
| 2:16.586 | לכבוד |  |
| 2:16.587 | השי"ת | , |
| 2:16.588 | ולהיות |  |
| 2:16.589 | זה |  |
| 2:16.590 | שלוחו |  |
| 2:16.591 | וכו | ' |
| 2:16.592 | ע"ש | . |
| 2:16.593 | ורז"ל |  |
| 2:16.594 | אמרו |  |
| 2:16.595 | תעלא |  |
| 2:16.596 | בעידניה |  |
| 2:16.597 | סגיד |  |
| 2:16.598 | ליה | , |
| 2:16.599 | ר"ל | : |
| 2:16.600 | אדם |  |
| 2:16.601 | פחות |  |
| 2:16.602 | ובזוי |  |
| 2:16.603 | כשועל | , |
| 2:16.604 | והשעה |  |
| 2:16.605 | עומדת |  |
| 2:16.606 | לו |  |
| 2:16.607 | באיזה |  |
| 2:16.608 | משרה |  |
| 2:16.609 | שבידו | , |
| 2:16.610 | סגיד |  |
| 2:16.611 | ליה | . |
| 2:16.612 | וכן |  |
| 2:16.613 | הוא |  |
| 2:16.614 | נהוג |  |
| 2:16.615 | אצלנו |  |
| 2:16.616 | ואצל |  |
| 2:16.617 | הערביים |  |
| 2:16.618 | להשתחוות |  |
| 2:16.619 | ולנשק |  |
| 2:16.620 | אדם |  |
| 2:16.621 | גדול |  |
| 2:16.622 | המעלה | , |
| 2:16.623 | ואין |  |
| 2:16.624 | בזה |  |
| 2:16.625 | שום |  |
| 2:16.626 | איסור | , |
| 2:16.627 | שאינו |  |
| 2:16.628 | מקבלו |  |
| 2:16.629 | עליו |  |
| 2:16.630 | באלוה | , |
| 2:16.631 | רק |  |
| 2:16.632 | לאות |  |
| 2:16.633 | כבוד |  |
| 2:16.634 | הוא |  |
| 2:16.635 | משתחוה |  |
| 2:16.636 | לו |  |
| 2:16.637 | ומנשק |  |
| 2:16.638 | את |  |
| 2:16.639 | רגליו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

