# Translation pack: milhamot-hashem perek-8-21 (Milchamot ha-Shem 8:25–8:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-21`.

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
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
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
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
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
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **ווי** ליה מאן דלא → "woe to him who does not".
- תהו **ובהו** → "formlessness and void".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 8:25

Hebrew (Wikisource): קיח) ואם נפשך לומר לפי דבריהם שהאלהים העליון הוא נשמה של הפרצופים הנבראים ולנשמתם אנו עובדים ומתפללים, כי הוא מתפשט וממלא את הספירות וכמו שכתב הפילסוף המסית מחבר הזהר דף הנז' לעיל אנא הוא לבר מכלא ולגיו מכלא, ולכל סטרא, ולעילא מכלא, ולתתא מכלא, ולית אלה אחרא עילא ותתא ומכל סטרא, ומלגיו דעסר ספירן דמנהון כלא ובהון תליא כלא ואנת בכל ספירה בארכה ורחבה עילא ותתא. ובין כל ספירה וספירה, ואנת הוא דמקריב לקודשא בריך הוא ושכיניתיה בכל ספירה וספירה ובכל ענפין וכו'. ובדברים אלו לקח את לב קוראיו, וחשבו לתקן את הסתירה ודברי מינות הגלויה ומפורסמת שהאלהים העליון שאין כמהו בכל האלהים. איננו רק מקשר ומחבר ומזויג את הספירות הנקראים קודשא בריך הוא ושכינתיה אשר נעבדם ונשתחוה להם לבדם. אבל הוא לא נעבדנו כנזכר לעיל. והנה עם מה שהקדמתי מדברי רז"ל אין קיום ועמידה לדברים אלו כלל שכיון שקיימו וקבלו שיש לנשמה ציור ובנין הגוף ברמ"ח איברים ושס"ה גידים. גם נשמת הנשמה יהיה לה איברים וגידים. וכן נשמת הנשמה עד אין קץ יהיה לה גם היא ציור ובנין הגוף שהיא נשמה לו. אלא שהיא דקה וקלושה יותר, וכמו שכתב עז לאלהים בבית קדש הקדשים פ"ג דף ט' ע"א שבהכרח גמור יש לנשמה ציור ובנין הגוף ברמ"ח איברים ושס"ה גידים בצורת אדם, ומתפשטת כפי איבריה. ואם תמצא אבר מיותר אינה מתפשטת בו, וכן ג"כ נשמה לנשמה צריך שיהיה לה בחינת הפרצוף וכו' ע"ש. ועוד שהרי ביארו שהפרצופים כולם נמשכים ומתהוים מעצמותו ית'. גם גופם וגם נשמתם ומלבושם כמו שכתב לעיל בשם שושן סודות. וכמבואר בהדיא בעץ חיים להרח"ו שער עיגולים ויושר ובשער ההקדמות. וכיון שהן נבראים אף שהם מעצמותו לפי דבריהם והוא שוכן בתוכם. איך ניחדם ונעבדם עם בוראם. ובהדיא מזהיר רשב"י ע"ה באמרו כל המשתף וכו' שנאמר בלתי לה' לבדו בלי שתוף וצירוף שום נברא בעולם, וכמו שכתב הרמב"ם ז"ל בהל' ע"ז עיקר הציווי בע"ז וכו'. הובא לעיל סי' נ"ו. וראיה לזה שלא לצרף עמו שום נברא אע"פ שהוא שוכן בו, שהרי צונו הקב"ה על ידי משה רבי' ע"ה לעשות לו משכן וארון וכפרת וכרובים ושאר כלי המשכן שנאמר ועשו לי מקדש ושכנתי בתוכם, וכאשר הוקם המשכן נר' כבודו ית' בו, שנאמר וכבוד ה' מלא וכו' ובבא משה אל אהל מועד, וישמע את הקול מדבר אליו וכו' ואף על פי כן אין אנו משתפים ומיחדים את המשכן וכליו או המקדש וכליו עם כבודו ית' השוכן בתוכו לומר שהמשכן או המקדש אלהינו, או נאמר הארון וכפרת וכרובים אלהינו, הואיל והוא אִוָם למושב לו, ולא נשתחוה ונכרע ונתפלל רק לה' אלהינו השוכן בתוכם, בלי צירוף ושיתוף המקום שהוא שוכן בו ולא ניחדם עמו ונאמר כולא חד.

| id | word | punctuation after |
|---|---|---|
| 8:25.1 | קיח | ) |
| 8:25.2 | ואם |  |
| 8:25.3 | נפשך |  |
| 8:25.4 | לומר |  |
| 8:25.5 | לפי |  |
| 8:25.6 | דבריהם |  |
| 8:25.7 | שהאלהים |  |
| 8:25.8 | העליון |  |
| 8:25.9 | הוא |  |
| 8:25.10 | נשמה |  |
| 8:25.11 | של |  |
| 8:25.12 | הפרצופים |  |
| 8:25.13 | הנבראים |  |
| 8:25.14 | ולנשמתם |  |
| 8:25.15 | אנו |  |
| 8:25.16 | עובדים |  |
| 8:25.17 | ומתפללים | , |
| 8:25.18 | כי |  |
| 8:25.19 | הוא |  |
| 8:25.20 | מתפשט |  |
| 8:25.21 | וממלא |  |
| 8:25.22 | את |  |
| 8:25.23 | הספירות |  |
| 8:25.24 | וכמו |  |
| 8:25.25 | שכתב |  |
| 8:25.26 | הפילסוף |  |
| 8:25.27 | המסית |  |
| 8:25.28 | מחבר |  |
| 8:25.29 | הזהר |  |
| 8:25.30 | דף |  |
| 8:25.31 | הנז | ' |
| 8:25.32 | לעיל |  |
| 8:25.33 | אנא |  |
| 8:25.34 | הוא |  |
| 8:25.35 | לבר |  |
| 8:25.36 | מכלא |  |
| 8:25.37 | ולגיו |  |
| 8:25.38 | מכלא | , |
| 8:25.39 | ולכל |  |
| 8:25.40 | סטרא | , |
| 8:25.41 | ולעילא |  |
| 8:25.42 | מכלא | , |
| 8:25.43 | ולתתא |  |
| 8:25.44 | מכלא | , |
| 8:25.45 | ולית |  |
| 8:25.46 | אלה |  |
| 8:25.47 | אחרא |  |
| 8:25.48 | עילא |  |
| 8:25.49 | ותתא |  |
| 8:25.50 | ומכל |  |
| 8:25.51 | סטרא | , |
| 8:25.52 | ומלגיו |  |
| 8:25.53 | דעסר |  |
| 8:25.54 | ספירן |  |
| 8:25.55 | דמנהון |  |
| 8:25.56 | כלא |  |
| 8:25.57 | ובהון |  |
| 8:25.58 | תליא |  |
| 8:25.59 | כלא |  |
| 8:25.60 | ואנת |  |
| 8:25.61 | בכל |  |
| 8:25.62 | ספירה |  |
| 8:25.63 | בארכה |  |
| 8:25.64 | ורחבה |  |
| 8:25.65 | עילא |  |
| 8:25.66 | ותתא | . |
| 8:25.67 | ובין |  |
| 8:25.68 | כל |  |
| 8:25.69 | ספירה |  |
| 8:25.70 | וספירה | , |
| 8:25.71 | ואנת |  |
| 8:25.72 | הוא |  |
| 8:25.73 | דמקריב |  |
| 8:25.74 | לקודשא |  |
| 8:25.75 | בריך |  |
| 8:25.76 | הוא |  |
| 8:25.77 | ושכיניתיה |  |
| 8:25.78 | בכל |  |
| 8:25.79 | ספירה |  |
| 8:25.80 | וספירה |  |
| 8:25.81 | ובכל |  |
| 8:25.82 | ענפין |  |
| 8:25.83 | וכו | '. |
| 8:25.84 | ובדברים |  |
| 8:25.85 | אלו |  |
| 8:25.86 | לקח |  |
| 8:25.87 | את |  |
| 8:25.88 | לב |  |
| 8:25.89 | קוראיו | , |
| 8:25.90 | וחשבו |  |
| 8:25.91 | לתקן |  |
| 8:25.92 | את |  |
| 8:25.93 | הסתירה |  |
| 8:25.94 | ודברי |  |
| 8:25.95 | מינות |  |
| 8:25.96 | הגלויה |  |
| 8:25.97 | ומפורסמת |  |
| 8:25.98 | שהאלהים |  |
| 8:25.99 | העליון |  |
| 8:25.100 | שאין |  |
| 8:25.101 | כמהו |  |
| 8:25.102 | בכל |  |
| 8:25.103 | האלהים | . |
| 8:25.104 | איננו |  |
| 8:25.105 | רק |  |
| 8:25.106 | מקשר |  |
| 8:25.107 | ומחבר |  |
| 8:25.108 | ומזויג |  |
| 8:25.109 | את |  |
| 8:25.110 | הספירות |  |
| 8:25.111 | הנקראים |  |
| 8:25.112 | קודשא |  |
| 8:25.113 | בריך |  |
| 8:25.114 | הוא |  |
| 8:25.115 | ושכינתיה |  |
| 8:25.116 | אשר |  |
| 8:25.117 | נעבדם |  |
| 8:25.118 | ונשתחוה |  |
| 8:25.119 | להם |  |
| 8:25.120 | לבדם | . |
| 8:25.121 | אבל |  |
| 8:25.122 | הוא |  |
| 8:25.123 | לא |  |
| 8:25.124 | נעבדנו |  |
| 8:25.125 | כנזכר |  |
| 8:25.126 | לעיל | . |
| 8:25.127 | והנה |  |
| 8:25.128 | עם |  |
| 8:25.129 | מה |  |
| 8:25.130 | שהקדמתי |  |
| 8:25.131 | מדברי |  |
| 8:25.132 | רז"ל |  |
| 8:25.133 | אין |  |
| 8:25.134 | קיום |  |
| 8:25.135 | ועמידה |  |
| 8:25.136 | לדברים |  |
| 8:25.137 | אלו |  |
| 8:25.138 | כלל |  |
| 8:25.139 | שכיון |  |
| 8:25.140 | שקיימו |  |
| 8:25.141 | וקבלו |  |
| 8:25.142 | שיש |  |
| 8:25.143 | לנשמה |  |
| 8:25.144 | ציור |  |
| 8:25.145 | ובנין |  |
| 8:25.146 | הגוף |  |
| 8:25.147 | ברמ"ח |  |
| 8:25.148 | איברים |  |
| 8:25.149 | ושס"ה |  |
| 8:25.150 | גידים | . |
| 8:25.151 | גם |  |
| 8:25.152 | נשמת |  |
| 8:25.153 | הנשמה |  |
| 8:25.154 | יהיה |  |
| 8:25.155 | לה |  |
| 8:25.156 | איברים |  |
| 8:25.157 | וגידים | . |
| 8:25.158 | וכן |  |
| 8:25.159 | נשמת |  |
| 8:25.160 | הנשמה |  |
| 8:25.161 | עד |  |
| 8:25.162 | אין |  |
| 8:25.163 | קץ |  |
| 8:25.164 | יהיה |  |
| 8:25.165 | לה |  |
| 8:25.166 | גם |  |
| 8:25.167 | היא |  |
| 8:25.168 | ציור |  |
| 8:25.169 | ובנין |  |
| 8:25.170 | הגוף |  |
| 8:25.171 | שהיא |  |
| 8:25.172 | נשמה |  |
| 8:25.173 | לו | . |
| 8:25.174 | אלא |  |
| 8:25.175 | שהיא |  |
| 8:25.176 | דקה |  |
| 8:25.177 | וקלושה |  |
| 8:25.178 | יותר | , |
| 8:25.179 | וכמו |  |
| 8:25.180 | שכתב |  |
| 8:25.181 | עז |  |
| 8:25.182 | לאלהים |  |
| 8:25.183 | בבית |  |
| 8:25.184 | קדש |  |
| 8:25.185 | הקדשים |  |
| 8:25.186 | פ"ג |  |
| 8:25.187 | דף |  |
| 8:25.188 | ט | ' |
| 8:25.189 | ע"א |  |
| 8:25.190 | שבהכרח |  |
| 8:25.191 | גמור |  |
| 8:25.192 | יש |  |
| 8:25.193 | לנשמה |  |
| 8:25.194 | ציור |  |
| 8:25.195 | ובנין |  |
| 8:25.196 | הגוף |  |
| 8:25.197 | ברמ"ח |  |
| 8:25.198 | איברים |  |
| 8:25.199 | ושס"ה |  |
| 8:25.200 | גידים |  |
| 8:25.201 | בצורת |  |
| 8:25.202 | אדם | , |
| 8:25.203 | ומתפשטת |  |
| 8:25.204 | כפי |  |
| 8:25.205 | איבריה | . |
| 8:25.206 | ואם |  |
| 8:25.207 | תמצא |  |
| 8:25.208 | אבר |  |
| 8:25.209 | מיותר |  |
| 8:25.210 | אינה |  |
| 8:25.211 | מתפשטת |  |
| 8:25.212 | בו | , |
| 8:25.213 | וכן |  |
| 8:25.214 | ג"כ |  |
| 8:25.215 | נשמה |  |
| 8:25.216 | לנשמה |  |
| 8:25.217 | צריך |  |
| 8:25.218 | שיהיה |  |
| 8:25.219 | לה |  |
| 8:25.220 | בחינת |  |
| 8:25.221 | הפרצוף |  |
| 8:25.222 | וכו | ' |
| 8:25.223 | ע"ש | . |
| 8:25.224 | ועוד |  |
| 8:25.225 | שהרי |  |
| 8:25.226 | ביארו |  |
| 8:25.227 | שהפרצופים |  |
| 8:25.228 | כולם |  |
| 8:25.229 | נמשכים |  |
| 8:25.230 | ומתהוים |  |
| 8:25.231 | מעצמותו |  |
| 8:25.232 | ית | '. |
| 8:25.233 | גם |  |
| 8:25.234 | גופם |  |
| 8:25.235 | וגם |  |
| 8:25.236 | נשמתם |  |
| 8:25.237 | ומלבושם |  |
| 8:25.238 | כמו |  |
| 8:25.239 | שכתב |  |
| 8:25.240 | לעיל |  |
| 8:25.241 | בשם |  |
| 8:25.242 | שושן |  |
| 8:25.243 | סודות | . |
| 8:25.244 | וכמבואר |  |
| 8:25.245 | בהדיא |  |
| 8:25.246 | בעץ |  |
| 8:25.247 | חיים |  |
| 8:25.248 | להרח"ו |  |
| 8:25.249 | שער |  |
| 8:25.250 | עיגולים |  |
| 8:25.251 | ויושר |  |
| 8:25.252 | ובשער |  |
| 8:25.253 | ההקדמות | . |
| 8:25.254 | וכיון |  |
| 8:25.255 | שהן |  |
| 8:25.256 | נבראים |  |
| 8:25.257 | אף |  |
| 8:25.258 | שהם |  |
| 8:25.259 | מעצמותו |  |
| 8:25.260 | לפי |  |
| 8:25.261 | דבריהם |  |
| 8:25.262 | והוא |  |
| 8:25.263 | שוכן |  |
| 8:25.264 | בתוכם | . |
| 8:25.265 | איך |  |
| 8:25.266 | ניחדם |  |
| 8:25.267 | ונעבדם |  |
| 8:25.268 | עם |  |
| 8:25.269 | בוראם | . |
| 8:25.270 | ובהדיא |  |
| 8:25.271 | מזהיר |  |
| 8:25.272 | רשב"י |  |
| 8:25.273 | ע"ה |  |
| 8:25.274 | באמרו |  |
| 8:25.275 | כל |  |
| 8:25.276 | המשתף |  |
| 8:25.277 | וכו | ' |
| 8:25.278 | שנאמר |  |
| 8:25.279 | בלתי |  |
| 8:25.280 | לה | ' |
| 8:25.281 | לבדו |  |
| 8:25.282 | בלי |  |
| 8:25.283 | שתוף |  |
| 8:25.284 | וצירוף |  |
| 8:25.285 | שום |  |
| 8:25.286 | נברא |  |
| 8:25.287 | בעולם | , |
| 8:25.288 | וכמו |  |
| 8:25.289 | שכתב |  |
| 8:25.290 | הרמב"ם |  |
| 8:25.291 | ז"ל |  |
| 8:25.292 | בהל | ' |
| 8:25.293 | ע"ז |  |
| 8:25.294 | עיקר |  |
| 8:25.295 | הציווי |  |
| 8:25.296 | בע"ז |  |
| 8:25.297 | וכו | '. |
| 8:25.298 | הובא |  |
| 8:25.299 | לעיל |  |
| 8:25.300 | סי | ' |
| 8:25.301 | נ"ו | . |
| 8:25.302 | וראיה |  |
| 8:25.303 | לזה |  |
| 8:25.304 | שלא |  |
| 8:25.305 | לצרף |  |
| 8:25.306 | עמו |  |
| 8:25.307 | שום |  |
| 8:25.308 | נברא |  |
| 8:25.309 | אע"פ |  |
| 8:25.310 | שהוא |  |
| 8:25.311 | שוכן |  |
| 8:25.312 | בו | , |
| 8:25.313 | שהרי |  |
| 8:25.314 | צונו |  |
| 8:25.315 | הקב"ה |  |
| 8:25.316 | על |  |
| 8:25.317 | ידי |  |
| 8:25.318 | משה |  |
| 8:25.319 | רבי | ' |
| 8:25.320 | ע"ה |  |
| 8:25.321 | לעשות |  |
| 8:25.322 | לו |  |
| 8:25.323 | משכן |  |
| 8:25.324 | וארון |  |
| 8:25.325 | וכפרת |  |
| 8:25.326 | וכרובים |  |
| 8:25.327 | ושאר |  |
| 8:25.328 | כלי |  |
| 8:25.329 | המשכן |  |
| 8:25.330 | שנאמר |  |
| 8:25.331 | ועשו |  |
| 8:25.332 | לי |  |
| 8:25.333 | מקדש |  |
| 8:25.334 | ושכנתי |  |
| 8:25.335 | בתוכם | , |
| 8:25.336 | וכאשר |  |
| 8:25.337 | הוקם |  |
| 8:25.338 | המשכן |  |
| 8:25.339 | נר | ' |
| 8:25.340 | כבודו |  |
| 8:25.341 | ית | ' |
| 8:25.342 | בו | , |
| 8:25.343 | שנאמר |  |
| 8:25.344 | וכבוד |  |
| 8:25.345 | ה | ' |
| 8:25.346 | מלא |  |
| 8:25.347 | וכו | ' |
| 8:25.348 | ובבא |  |
| 8:25.349 | משה |  |
| 8:25.350 | אל |  |
| 8:25.351 | אהל |  |
| 8:25.352 | מועד | , |
| 8:25.353 | וישמע |  |
| 8:25.354 | את |  |
| 8:25.355 | הקול |  |
| 8:25.356 | מדבר |  |
| 8:25.357 | אליו |  |
| 8:25.358 | וכו | ' |
| 8:25.359 | ואף |  |
| 8:25.360 | על |  |
| 8:25.361 | פי |  |
| 8:25.362 | כן |  |
| 8:25.363 | אין |  |
| 8:25.364 | אנו |  |
| 8:25.365 | משתפים |  |
| 8:25.366 | ומיחדים |  |
| 8:25.367 | את |  |
| 8:25.368 | המשכן |  |
| 8:25.369 | וכליו |  |
| 8:25.370 | או |  |
| 8:25.371 | המקדש |  |
| 8:25.372 | וכליו |  |
| 8:25.373 | עם |  |
| 8:25.374 | כבודו |  |
| 8:25.375 | ית | ' |
| 8:25.376 | השוכן |  |
| 8:25.377 | בתוכו |  |
| 8:25.378 | לומר |  |
| 8:25.379 | שהמשכן |  |
| 8:25.380 | או |  |
| 8:25.381 | המקדש |  |
| 8:25.382 | אלהינו | , |
| 8:25.383 | או |  |
| 8:25.384 | נאמר |  |
| 8:25.385 | הארון |  |
| 8:25.386 | וכפרת |  |
| 8:25.387 | וכרובים |  |
| 8:25.388 | אלהינו | , |
| 8:25.389 | הואיל |  |
| 8:25.390 | והוא |  |
| 8:25.391 | אִוָם |  |
| 8:25.392 | למושב |  |
| 8:25.393 | לו | , |
| 8:25.394 | ולא |  |
| 8:25.395 | נשתחוה |  |
| 8:25.396 | ונכרע |  |
| 8:25.397 | ונתפלל |  |
| 8:25.398 | רק |  |
| 8:25.399 | לה | ' |
| 8:25.400 | אלהינו |  |
| 8:25.401 | השוכן |  |
| 8:25.402 | בתוכם | , |
| 8:25.403 | בלי |  |
| 8:25.404 | צירוף |  |
| 8:25.405 | ושיתוף |  |
| 8:25.406 | המקום |  |
| 8:25.407 | שהוא |  |
| 8:25.408 | שוכן |  |
| 8:25.409 | בו |  |
| 8:25.410 | ולא |  |
| 8:25.411 | ניחדם |  |
| 8:25.412 | עמו |  |
| 8:25.413 | ונאמר |  |
| 8:25.414 | כולא |  |
| 8:25.415 | חד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

