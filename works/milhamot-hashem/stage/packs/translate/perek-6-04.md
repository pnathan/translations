# Translation pack: milhamot-hashem perek-6-04 (Milchamot ha-Shem 6:8–6:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-04`.

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
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **רש"י** → "Rashi" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הכוזרי / כוזרי** → "the Kuzari" (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- מאמיני **השלוש** → "the believers in the Trinity".
- **בא"י** → "in the Land of Israel".

## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
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
- **מין / מינות** → "heretic / heresy".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- תורה **שבעל** פה → "the Oral Torah".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **כ"ש / ק"ו** → "all the more so".
- **סתו"מ** → "Torah scrolls, tefillin and mezuzot" (note).
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- ואין **צריך** לומר → "needless to say".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- מי **אתי** (ketiv): translated, the qere noted.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **העתקתו / תרגום**: his own Hebrew translations (10:7–10:9; "[תרגום הרב
  המחבר: …]") are separate segments, each from its own Hebrew; a difference
  in sense in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 6:8

Hebrew (Wikisource): והנה מצות האחדות ר"ל להאמין שהו ית' אחד פשוט אמתי הנמשך אחר שלילת המקרים, והתארים, ושלילת הגשמות, ואמונת חיוב מציאותו בבחינת עצמו, ושלילת השינוי ממנו ית' וכיוצא בהם מהאמונות האלהיות הבלתי מבוארת לא בתורה שבכתב ולא בתורה שבעל פה, שהמאמין בהם שלא כדת אינו בכלל בעלי האמונה השלימה, ולא אותו ית' יעבוד כמו שכ' הרב בעל חובת הלבבות לא יוכל לעבוד עילת העילות זולת הנביא או החכם השלום, על כן המצוה עלינו לחקור על מצות אלה מפי סופרים ומספרים האלהיים למען לא יטעה באמונתו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:8.1 | והנה |  |
| 6:8.2 | מצות |  |
| 6:8.3 | האחדות |  |
| 6:8.4 | ר"ל |  |
| 6:8.5 | להאמין |  |
| 6:8.6 | שהו |  |
| 6:8.7 | ית | ' |
| 6:8.8 | אחד |  |
| 6:8.9 | פשוט |  |
| 6:8.10 | אמתי |  |
| 6:8.11 | הנמשך |  |
| 6:8.12 | אחר |  |
| 6:8.13 | שלילת |  |
| 6:8.14 | המקרים | , |
| 6:8.15 | והתארים | , |
| 6:8.16 | ושלילת |  |
| 6:8.17 | הגשמות | , |
| 6:8.18 | ואמונת |  |
| 6:8.19 | חיוב |  |
| 6:8.20 | מציאותו |  |
| 6:8.21 | בבחינת |  |
| 6:8.22 | עצמו | , |
| 6:8.23 | ושלילת |  |
| 6:8.24 | השינוי |  |
| 6:8.25 | ממנו |  |
| 6:8.26 | ית | ' |
| 6:8.27 | וכיוצא |  |
| 6:8.28 | בהם |  |
| 6:8.29 | מהאמונות |  |
| 6:8.30 | האלהיות |  |
| 6:8.31 | הבלתי |  |
| 6:8.32 | מבוארת |  |
| 6:8.33 | לא |  |
| 6:8.34 | בתורה |  |
| 6:8.35 | שבכתב |  |
| 6:8.36 | ולא |  |
| 6:8.37 | בתורה |  |
| 6:8.38 | שבעל |  |
| 6:8.39 | פה | , |
| 6:8.40 | שהמאמין |  |
| 6:8.41 | בהם |  |
| 6:8.42 | שלא |  |
| 6:8.43 | כדת |  |
| 6:8.44 | אינו |  |
| 6:8.45 | בכלל |  |
| 6:8.46 | בעלי |  |
| 6:8.47 | האמונה |  |
| 6:8.48 | השלימה | , |
| 6:8.49 | ולא |  |
| 6:8.50 | אותו |  |
| 6:8.51 | ית | ' |
| 6:8.52 | יעבוד |  |
| 6:8.53 | כמו |  |
| 6:8.54 | שכ | ' |
| 6:8.55 | הרב |  |
| 6:8.56 | בעל |  |
| 6:8.57 | חובת |  |
| 6:8.58 | הלבבות |  |
| 6:8.59 | לא |  |
| 6:8.60 | יוכל |  |
| 6:8.61 | לעבוד |  |
| 6:8.62 | עילת |  |
| 6:8.63 | העילות |  |
| 6:8.64 | זולת |  |
| 6:8.65 | הנביא |  |
| 6:8.66 | או |  |
| 6:8.67 | החכם |  |
| 6:8.68 | השלום | , |
| 6:8.69 | על |  |
| 6:8.70 | כן |  |
| 6:8.71 | המצוה |  |
| 6:8.72 | עלינו |  |
| 6:8.73 | לחקור |  |
| 6:8.74 | על |  |
| 6:8.75 | מצות |  |
| 6:8.76 | אלה |  |
| 6:8.77 | מפי |  |
| 6:8.78 | סופרים |  |
| 6:8.79 | ומספרים |  |
| 6:8.80 | האלהיים |  |
| 6:8.81 | למען |  |
| 6:8.82 | לא |  |
| 6:8.83 | יטעה |  |
| 6:8.84 | באמונתו |  |
| 6:8.85 | עכ"ל | . |


## Milchamot ha-Shem 6:9

Hebrew (Wikisource): סד) ודבריו אלה שכתב שמצות האחדות לא ביארוה רבותינו ז"ל נפלאים בעיני עד מאוד מכמה פנים, כי הנה ידוע הוא ומפורסם שדרך רז"ל בעלי המשנה והתלמוד שלא להזכיר רק דבר שאפשר להסתפק בו אם הוא אסור או מותר, אבל דברי הפשוט איסורו או התירו לא יזכיר התנא כלל, דזיל קרי בי רב הוא, וכן כתב הרמב"ם ז"ל בפי' המשנה פרק ז' דברכות וז"ל ודע כי המשנה לא תדבר דבר באיסור דבר המפורסם וידוע איסורו או בהתר דבר שידוע התירו, אך בדבר שיש לספק בו אם הוא אסור או מותר, תבאר המשנה לאסור דבר שיעלה על הדעת שהוא מותר, או בהתר דבר שיעלה על הדעת שהוא אסור, אבל דברים הידועים לא תדבר בהם כלל, ולכן מקשים בגמרא תמיד פשיטא? מאי קמ"ל? אם היה דבר ידוע ומפורסם עכ"ל. וכבר נתבאר לעיל בדברי הר' הסמ"ג כי אחדותו ית' פשוטה ותקועה בלב כל ישראל, וכל שכן בימי רז"ל בעלי המשנה והתלמוד, והשנית שהיא מבוארת בתורה נביאים וכתובים במקומות אין מספר, שמע ישראל ה' אלהינו ה' אחד, וידעת היום וכו' כי ה' הוא האלהים בשמים ועל הארץ מתחת אין עוד, ואמר אין עוד מלבדו ראו עתה כי אני אני הוא ואין אלהים עמדי, ה' בדד ינחנו ואין עמו אל נכר, נוטה שמים לבדי, רוקע הארץ מיאתי, מי הקדימני, ועוד הרבה כתובים מדברים באחדותו ית' וקדמותו, ומפורסם ג"כ באומה הישראלית כנז"ל ס' ס"א, לכן לא הוצרכו רז"ל לבאר היחוד האמתי, להיותו מן המפורסמות, גם שלילת הגשמות מבוארת בכתובים כי לא ראיתם כל תמונה ותמונה אינכם רואים, ואל מי תדמיוני ואשוה, ואל מי תדמיון אל, והרבה כתובים כיוצא בהם, והשינוי ג"כ מבואר שלילתו מהשי"ת בכתוב הוא אמרו ית' אני ה' לא שניתי, וכן שאר דברים המביאים לענינים הנזכרים, ועם כל זה לא שתקו רז"ל בגמרא בבלי וירושלמי ומדרש רבה ותנחומא ושאר מדרשים הידועים ומפורסמים לרז"ל מלהזכיר תשובת טענות ושאלות שאירעו וקרו להם מן החיצונים הכופרים הטוענים נגד אמונותיו ויש בהם די להורותינו הדרך הנכונה ביחוד השי"ת שהוא כמו שביארו ר' סעדיה גאון והרמב"ם בפי' המשנה ובחבור ובמורה ור' בחיי בחובת הלבבות, והרקח, ושבילי אמונה, ור' יהודה הלוי בכוזרי. וגם הרחיקו רז"ל מלהרבות בתארי השי"ת. ומכל ענין היוצא ממנו משמעות שתי רשויות או יותר כאמרם האומר שמע שמע או מודים מודים וכיוצא משתקים אותו וכמו שנת' לעיל סי' ל"ב. וכ"ש שנמצא את כל אלה על פי הקבלה האמתית ביחוד השי"ת אם נשכיל ונעמיק חרק בתרגומים המסורים לנו מרז"ל תרגום אונקלוס הגר ע"ה שתרגם את התורה מפי ר' אליעזר ור' יהושע, ותרגום יונתן בן עוזיאל הגדול שבתלמיד הלל שתרגם את הנביאים כמו שקבל מהלל. ומצאנו ראינו שהם משרישים האמונה האמתית בלב העם ומתקנים תרגום הכתובים שלא יפול ספק בקבלה האמתית כמו שכ' הרמב"ם והרס"ג. ועיין בהקדמה כוללת להר"ר נתן אדלר בעל נתינה לגר. ועל כן ארז"ל המתרגם פסוק כצורתו הרי זה בדאי. לפי שדרך הכתובים לדבר בדרך הרחבת הלשון, או בלשון מושאל, או משותף, או בלשון קצר, או במוקדם ומאוחר. צריך המתרגם לתרגמו כמו שקבלו חכמים, ואלו הן סודות התורה לפרש מקראות הסתומים כפי האמת שקבלו חכמים. ולכן אמרו רז"ל כי כשתרגם יונתן בן עוזיאל את הנביאים יצאה בת קול ואמרה מי גלה רזי לבני אדם. ועליו יאמר רב יוסף בכמה מקומות אלמלא תרגומא דהאיי קרא לא ידענא מאי קאמר. ובמדרש הגדול פ' משפטים ויראו את אלהי ישראל. אמר ר' אליעזר כל המתרגם פסוק כצורתו ה"ז בדאי וכל המוסיף בו הרי זה מחרף ומגדף כגון ויראו את אלהי ישראל. וחזו ית אלהא דישראל ה"ז בדאי. שהקב"ה רואה ואינו נראה. ואם תרגם וחזו ית יקר שכינת אלהא דישראל הרי זה מחרף ומגדף. שהוא עושה כאן שלשה, יקר, ושכינה ואל עכ"ל. וקאי לי טובא שהרי יונתן בן עוזיאל תרגם כי את המלך ה' צבאות ראו עיני. ארי ית יקר שכינת מלך עלמיא ה' צבאות חזאה עיני והרי"ף בסוף מסכת מגלה הביא תוספתא זו ופי' ר' נסים המתרגם פסוק כצורתו וחזו ית' מלאכא דאלהא דישראל, ורחות כרסי יקריה ה"ז מגדף שהשוה כבוד עבד לכבוד רבו ע"כ, והערוך ערך תרגם פי' הרי מחרף ומגדף שעושה הכבוד מלאך, אלא יתרגם וחזו ית יקר אלהא דישראל וכו'. וכנגד אמונת השלוש שנתחדשה בזמן רז"ל בעלי המשנה מתלמידי ישוע הנצרי אמרו רז"ל במדרש רבה פרשת יתרו אנכי ה' אלהיך אשר הוציאתיך וגו' אמר הקב"ה מלך בשר ודם מולך יש לו אב או אח, או בן, אני איני כן, אני ראשון שאין לי אב, ואני אחרון שאין לי בן, ומבלעדי אין אלהים שאין לי אח ע"כ. והמקובלים החדשים נמשכו אחר דעת בעלי השלוש המאמינים באב ובן ורוח הקודש כמובן מהמשל המבואר בס' עז לאלהים בבית קודש הקדשים דף נ"ט ע"ב וז"ל ולפי שמלכא קדישא דכל קדישין לא מתפרש מניה לעלמין והוא מאיר בזעיר נאפין דהוא לב האצילות לכן הוא ג"כ יקרא בוצינא דקרדינותא שמאיר בלב שהוא זעיר אנפין דאצילות. נמצא שמה שנקרא זה השם הוא על שם נשמתו ולא על שם מלבושו כשאר השמות. והוא הנקרא מלך לבני ישראל כאשר נתבאר בפרשת ואלה המלכים. והוא מלכינו ואנחנו עבדיו. ואליו תפלותינו וכוונותינו. ואפי' אבא ואימא מוזהרים לכבד את הבן הזה וליראה ממנו, לפי שהוא המלך השליט על כלא. ואפי' שיצא הבן הקדוש הזה מאבא ואימא יתבאר לקטן שהוא העיקר, אבל בכל זאת אמשול לך משל אם ימלוך א' ויהיה לו אב ואם. למי צריכים לעבוד? בודאי למלך. ואפי' אביו ואמו צריכים לו לפי שהוא מלך. על זה אמר נשקו בר פן יאנף. ואמר בזהר נשקו ידין להאי ברא קדישא. כי שולטנו אתיהיב ליה על כלא וכו' הרי שאפי' אבא ואימא צריכים לו עכ"ל. הרי בהדיא שהוא זותר לדברי רז"ל במדרש שמות רבה הנז' לעיל, וירושלמי בפ' דיני ממונות, דגרסינן התם מהו חותמו של הקב"ה רב בבי בשם ר ראובן אמת. מהו אמת. אמר ר' בון שהוא אלהים חיים ומלך עולם. אמר ריש לקיש אלף רישא דאלף בית מים באמצעיתא, תיו בסופא. לומר אני ה' ראשון שלא קבלתי מאחר. ומבלעדי אין אלהים שאין לי שותף. ואת אחרונים אני הוא שאיני עתיד למוסרה לאחר עכ"ל הירוש'.

| id | word | punctuation after |
|---|---|---|
| 6:9.1 | סד | ) |
| 6:9.2 | ודבריו |  |
| 6:9.3 | אלה |  |
| 6:9.4 | שכתב |  |
| 6:9.5 | שמצות |  |
| 6:9.6 | האחדות |  |
| 6:9.7 | לא |  |
| 6:9.8 | ביארוה |  |
| 6:9.9 | רבותינו |  |
| 6:9.10 | ז"ל |  |
| 6:9.11 | נפלאים |  |
| 6:9.12 | בעיני |  |
| 6:9.13 | עד |  |
| 6:9.14 | מאוד |  |
| 6:9.15 | מכמה |  |
| 6:9.16 | פנים | , |
| 6:9.17 | כי |  |
| 6:9.18 | הנה |  |
| 6:9.19 | ידוע |  |
| 6:9.20 | הוא |  |
| 6:9.21 | ומפורסם |  |
| 6:9.22 | שדרך |  |
| 6:9.23 | רז"ל |  |
| 6:9.24 | בעלי |  |
| 6:9.25 | המשנה |  |
| 6:9.26 | והתלמוד |  |
| 6:9.27 | שלא |  |
| 6:9.28 | להזכיר |  |
| 6:9.29 | רק |  |
| 6:9.30 | דבר |  |
| 6:9.31 | שאפשר |  |
| 6:9.32 | להסתפק |  |
| 6:9.33 | בו |  |
| 6:9.34 | אם |  |
| 6:9.35 | הוא |  |
| 6:9.36 | אסור |  |
| 6:9.37 | או |  |
| 6:9.38 | מותר | , |
| 6:9.39 | אבל |  |
| 6:9.40 | דברי |  |
| 6:9.41 | הפשוט |  |
| 6:9.42 | איסורו |  |
| 6:9.43 | או |  |
| 6:9.44 | התירו |  |
| 6:9.45 | לא |  |
| 6:9.46 | יזכיר |  |
| 6:9.47 | התנא |  |
| 6:9.48 | כלל | , |
| 6:9.49 | דזיל |  |
| 6:9.50 | קרי |  |
| 6:9.51 | בי |  |
| 6:9.52 | רב |  |
| 6:9.53 | הוא | , |
| 6:9.54 | וכן |  |
| 6:9.55 | כתב |  |
| 6:9.56 | הרמב"ם |  |
| 6:9.57 | ז"ל |  |
| 6:9.58 | בפי | ' |
| 6:9.59 | המשנה |  |
| 6:9.60 | פרק |  |
| 6:9.61 | ז | ' |
| 6:9.62 | דברכות |  |
| 6:9.63 | וז"ל |  |
| 6:9.64 | ודע |  |
| 6:9.65 | כי |  |
| 6:9.66 | המשנה |  |
| 6:9.67 | לא |  |
| 6:9.68 | תדבר |  |
| 6:9.69 | דבר |  |
| 6:9.70 | באיסור |  |
| 6:9.71 | דבר |  |
| 6:9.72 | המפורסם |  |
| 6:9.73 | וידוע |  |
| 6:9.74 | איסורו |  |
| 6:9.75 | או |  |
| 6:9.76 | בהתר |  |
| 6:9.77 | דבר |  |
| 6:9.78 | שידוע |  |
| 6:9.79 | התירו | , |
| 6:9.80 | אך |  |
| 6:9.81 | בדבר |  |
| 6:9.82 | שיש |  |
| 6:9.83 | לספק |  |
| 6:9.84 | בו |  |
| 6:9.85 | אם |  |
| 6:9.86 | הוא |  |
| 6:9.87 | אסור |  |
| 6:9.88 | או |  |
| 6:9.89 | מותר | , |
| 6:9.90 | תבאר |  |
| 6:9.91 | המשנה |  |
| 6:9.92 | לאסור |  |
| 6:9.93 | דבר |  |
| 6:9.94 | שיעלה |  |
| 6:9.95 | על |  |
| 6:9.96 | הדעת |  |
| 6:9.97 | שהוא |  |
| 6:9.98 | מותר | , |
| 6:9.99 | או |  |
| 6:9.100 | בהתר |  |
| 6:9.101 | דבר |  |
| 6:9.102 | שיעלה |  |
| 6:9.103 | על |  |
| 6:9.104 | הדעת |  |
| 6:9.105 | שהוא |  |
| 6:9.106 | אסור | , |
| 6:9.107 | אבל |  |
| 6:9.108 | דברים |  |
| 6:9.109 | הידועים |  |
| 6:9.110 | לא |  |
| 6:9.111 | תדבר |  |
| 6:9.112 | בהם |  |
| 6:9.113 | כלל | , |
| 6:9.114 | ולכן |  |
| 6:9.115 | מקשים |  |
| 6:9.116 | בגמרא |  |
| 6:9.117 | תמיד |  |
| 6:9.118 | פשיטא | ? |
| 6:9.119 | מאי |  |
| 6:9.120 | קמ"ל | ? |
| 6:9.121 | אם |  |
| 6:9.122 | היה |  |
| 6:9.123 | דבר |  |
| 6:9.124 | ידוע |  |
| 6:9.125 | ומפורסם |  |
| 6:9.126 | עכ"ל | . |
| 6:9.127 | וכבר |  |
| 6:9.128 | נתבאר |  |
| 6:9.129 | לעיל |  |
| 6:9.130 | בדברי |  |
| 6:9.131 | הר | ' |
| 6:9.132 | הסמ"ג |  |
| 6:9.133 | כי |  |
| 6:9.134 | אחדותו |  |
| 6:9.135 | ית | ' |
| 6:9.136 | פשוטה |  |
| 6:9.137 | ותקועה |  |
| 6:9.138 | בלב |  |
| 6:9.139 | כל |  |
| 6:9.140 | ישראל | , |
| 6:9.141 | וכל |  |
| 6:9.142 | שכן |  |
| 6:9.143 | בימי |  |
| 6:9.144 | רז"ל |  |
| 6:9.145 | בעלי |  |
| 6:9.146 | המשנה |  |
| 6:9.147 | והתלמוד | , |
| 6:9.148 | והשנית |  |
| 6:9.149 | שהיא |  |
| 6:9.150 | מבוארת |  |
| 6:9.151 | בתורה |  |
| 6:9.152 | נביאים |  |
| 6:9.153 | וכתובים |  |
| 6:9.154 | במקומות |  |
| 6:9.155 | אין |  |
| 6:9.156 | מספר | , |
| 6:9.157 | שמע |  |
| 6:9.158 | ישראל |  |
| 6:9.159 | ה | ' |
| 6:9.160 | אלהינו |  |
| 6:9.161 | ה | ' |
| 6:9.162 | אחד | , |
| 6:9.163 | וידעת |  |
| 6:9.164 | היום |  |
| 6:9.165 | וכו | ' |
| 6:9.166 | כי |  |
| 6:9.167 | ה | ' |
| 6:9.168 | הוא |  |
| 6:9.169 | האלהים |  |
| 6:9.170 | בשמים |  |
| 6:9.171 | ועל |  |
| 6:9.172 | הארץ |  |
| 6:9.173 | מתחת |  |
| 6:9.174 | אין |  |
| 6:9.175 | עוד | , |
| 6:9.176 | ואמר |  |
| 6:9.177 | אין |  |
| 6:9.178 | עוד |  |
| 6:9.179 | מלבדו |  |
| 6:9.180 | ראו |  |
| 6:9.181 | עתה |  |
| 6:9.182 | כי |  |
| 6:9.183 | אני |  |
| 6:9.184 | אני |  |
| 6:9.185 | הוא |  |
| 6:9.186 | ואין |  |
| 6:9.187 | אלהים |  |
| 6:9.188 | עמדי | , |
| 6:9.189 | ה | ' |
| 6:9.190 | בדד |  |
| 6:9.191 | ינחנו |  |
| 6:9.192 | ואין |  |
| 6:9.193 | עמו |  |
| 6:9.194 | אל |  |
| 6:9.195 | נכר | , |
| 6:9.196 | נוטה |  |
| 6:9.197 | שמים |  |
| 6:9.198 | לבדי | , |
| 6:9.199 | רוקע |  |
| 6:9.200 | הארץ |  |
| 6:9.201 | מיאתי | , |
| 6:9.202 | מי |  |
| 6:9.203 | הקדימני | , |
| 6:9.204 | ועוד |  |
| 6:9.205 | הרבה |  |
| 6:9.206 | כתובים |  |
| 6:9.207 | מדברים |  |
| 6:9.208 | באחדותו |  |
| 6:9.209 | ית | ' |
| 6:9.210 | וקדמותו | , |
| 6:9.211 | ומפורסם |  |
| 6:9.212 | ג"כ |  |
| 6:9.213 | באומה |  |
| 6:9.214 | הישראלית |  |
| 6:9.215 | כנז"ל |  |
| 6:9.216 | ס | ' |
| 6:9.217 | ס"א | , |
| 6:9.218 | לכן |  |
| 6:9.219 | לא |  |
| 6:9.220 | הוצרכו |  |
| 6:9.221 | רז"ל |  |
| 6:9.222 | לבאר |  |
| 6:9.223 | היחוד |  |
| 6:9.224 | האמתי | , |
| 6:9.225 | להיותו |  |
| 6:9.226 | מן |  |
| 6:9.227 | המפורסמות | , |
| 6:9.228 | גם |  |
| 6:9.229 | שלילת |  |
| 6:9.230 | הגשמות |  |
| 6:9.231 | מבוארת |  |
| 6:9.232 | בכתובים |  |
| 6:9.233 | כי |  |
| 6:9.234 | לא |  |
| 6:9.235 | ראיתם |  |
| 6:9.236 | כל |  |
| 6:9.237 | תמונה |  |
| 6:9.238 | ותמונה |  |
| 6:9.239 | אינכם |  |
| 6:9.240 | רואים | , |
| 6:9.241 | ואל |  |
| 6:9.242 | מי |  |
| 6:9.243 | תדמיוני |  |
| 6:9.244 | ואשוה | , |
| 6:9.245 | ואל |  |
| 6:9.246 | מי |  |
| 6:9.247 | תדמיון |  |
| 6:9.248 | אל | , |
| 6:9.249 | והרבה |  |
| 6:9.250 | כתובים |  |
| 6:9.251 | כיוצא |  |
| 6:9.252 | בהם | , |
| 6:9.253 | והשינוי |  |
| 6:9.254 | ג"כ |  |
| 6:9.255 | מבואר |  |
| 6:9.256 | שלילתו |  |
| 6:9.257 | מהשי"ת |  |
| 6:9.258 | בכתוב |  |
| 6:9.259 | הוא |  |
| 6:9.260 | אמרו |  |
| 6:9.261 | ית | ' |
| 6:9.262 | אני |  |
| 6:9.263 | ה | ' |
| 6:9.264 | לא |  |
| 6:9.265 | שניתי | , |
| 6:9.266 | וכן |  |
| 6:9.267 | שאר |  |
| 6:9.268 | דברים |  |
| 6:9.269 | המביאים |  |
| 6:9.270 | לענינים |  |
| 6:9.271 | הנזכרים | , |
| 6:9.272 | ועם |  |
| 6:9.273 | כל |  |
| 6:9.274 | זה |  |
| 6:9.275 | לא |  |
| 6:9.276 | שתקו |  |
| 6:9.277 | רז"ל |  |
| 6:9.278 | בגמרא |  |
| 6:9.279 | בבלי |  |
| 6:9.280 | וירושלמי |  |
| 6:9.281 | ומדרש |  |
| 6:9.282 | רבה |  |
| 6:9.283 | ותנחומא |  |
| 6:9.284 | ושאר |  |
| 6:9.285 | מדרשים |  |
| 6:9.286 | הידועים |  |
| 6:9.287 | ומפורסמים |  |
| 6:9.288 | לרז"ל |  |
| 6:9.289 | מלהזכיר |  |
| 6:9.290 | תשובת |  |
| 6:9.291 | טענות |  |
| 6:9.292 | ושאלות |  |
| 6:9.293 | שאירעו |  |
| 6:9.294 | וקרו |  |
| 6:9.295 | להם |  |
| 6:9.296 | מן |  |
| 6:9.297 | החיצונים |  |
| 6:9.298 | הכופרים |  |
| 6:9.299 | הטוענים |  |
| 6:9.300 | נגד |  |
| 6:9.301 | אמונותיו |  |
| 6:9.302 | ויש |  |
| 6:9.303 | בהם |  |
| 6:9.304 | די |  |
| 6:9.305 | להורותינו |  |
| 6:9.306 | הדרך |  |
| 6:9.307 | הנכונה |  |
| 6:9.308 | ביחוד |  |
| 6:9.309 | השי"ת |  |
| 6:9.310 | שהוא |  |
| 6:9.311 | כמו |  |
| 6:9.312 | שביארו |  |
| 6:9.313 | ר | ' |
| 6:9.314 | סעדיה |  |
| 6:9.315 | גאון |  |
| 6:9.316 | והרמב"ם |  |
| 6:9.317 | בפי | ' |
| 6:9.318 | המשנה |  |
| 6:9.319 | ובחבור |  |
| 6:9.320 | ובמורה |  |
| 6:9.321 | ור | ' |
| 6:9.322 | בחיי |  |
| 6:9.323 | בחובת |  |
| 6:9.324 | הלבבות | , |
| 6:9.325 | והרקח | , |
| 6:9.326 | ושבילי |  |
| 6:9.327 | אמונה | , |
| 6:9.328 | ור | ' |
| 6:9.329 | יהודה |  |
| 6:9.330 | הלוי |  |
| 6:9.331 | בכוזרי | . |
| 6:9.332 | וגם |  |
| 6:9.333 | הרחיקו |  |
| 6:9.334 | רז"ל |  |
| 6:9.335 | מלהרבות |  |
| 6:9.336 | בתארי |  |
| 6:9.337 | השי"ת | . |
| 6:9.338 | ומכל |  |
| 6:9.339 | ענין |  |
| 6:9.340 | היוצא |  |
| 6:9.341 | ממנו |  |
| 6:9.342 | משמעות |  |
| 6:9.343 | שתי |  |
| 6:9.344 | רשויות |  |
| 6:9.345 | או |  |
| 6:9.346 | יותר |  |
| 6:9.347 | כאמרם |  |
| 6:9.348 | האומר |  |
| 6:9.349 | שמע |  |
| 6:9.350 | שמע |  |
| 6:9.351 | או |  |
| 6:9.352 | מודים |  |
| 6:9.353 | מודים |  |
| 6:9.354 | וכיוצא |  |
| 6:9.355 | משתקים |  |
| 6:9.356 | אותו |  |
| 6:9.357 | וכמו |  |
| 6:9.358 | שנת | ' |
| 6:9.359 | לעיל |  |
| 6:9.360 | סי | ' |
| 6:9.361 | ל"ב | . |
| 6:9.362 | וכ"ש |  |
| 6:9.363 | שנמצא |  |
| 6:9.364 | את |  |
| 6:9.365 | כל |  |
| 6:9.366 | אלה |  |
| 6:9.367 | על |  |
| 6:9.368 | פי |  |
| 6:9.369 | הקבלה |  |
| 6:9.370 | האמתית |  |
| 6:9.371 | ביחוד |  |
| 6:9.372 | השי"ת |  |
| 6:9.373 | אם |  |
| 6:9.374 | נשכיל |  |
| 6:9.375 | ונעמיק |  |
| 6:9.376 | חרק |  |
| 6:9.377 | בתרגומים |  |
| 6:9.378 | המסורים |  |
| 6:9.379 | לנו |  |
| 6:9.380 | מרז"ל |  |
| 6:9.381 | תרגום |  |
| 6:9.382 | אונקלוס |  |
| 6:9.383 | הגר |  |
| 6:9.384 | ע"ה |  |
| 6:9.385 | שתרגם |  |
| 6:9.386 | את |  |
| 6:9.387 | התורה |  |
| 6:9.388 | מפי |  |
| 6:9.389 | ר | ' |
| 6:9.390 | אליעזר |  |
| 6:9.391 | ור | ' |
| 6:9.392 | יהושע | , |
| 6:9.393 | ותרגום |  |
| 6:9.394 | יונתן |  |
| 6:9.395 | בן |  |
| 6:9.396 | עוזיאל |  |
| 6:9.397 | הגדול |  |
| 6:9.398 | שבתלמיד |  |
| 6:9.399 | הלל |  |
| 6:9.400 | שתרגם |  |
| 6:9.401 | את |  |
| 6:9.402 | הנביאים |  |
| 6:9.403 | כמו |  |
| 6:9.404 | שקבל |  |
| 6:9.405 | מהלל | . |
| 6:9.406 | ומצאנו |  |
| 6:9.407 | ראינו |  |
| 6:9.408 | שהם |  |
| 6:9.409 | משרישים |  |
| 6:9.410 | האמונה |  |
| 6:9.411 | האמתית |  |
| 6:9.412 | בלב |  |
| 6:9.413 | העם |  |
| 6:9.414 | ומתקנים |  |
| 6:9.415 | תרגום |  |
| 6:9.416 | הכתובים |  |
| 6:9.417 | שלא |  |
| 6:9.418 | יפול |  |
| 6:9.419 | ספק |  |
| 6:9.420 | בקבלה |  |
| 6:9.421 | האמתית |  |
| 6:9.422 | כמו |  |
| 6:9.423 | שכ | ' |
| 6:9.424 | הרמב"ם |  |
| 6:9.425 | והרס"ג | . |
| 6:9.426 | ועיין |  |
| 6:9.427 | בהקדמה |  |
| 6:9.428 | כוללת |  |
| 6:9.429 | להר"ר |  |
| 6:9.430 | נתן |  |
| 6:9.431 | אדלר |  |
| 6:9.432 | בעל |  |
| 6:9.433 | נתינה |  |
| 6:9.434 | לגר | . |
| 6:9.435 | ועל |  |
| 6:9.436 | כן |  |
| 6:9.437 | ארז"ל |  |
| 6:9.438 | המתרגם |  |
| 6:9.439 | פסוק |  |
| 6:9.440 | כצורתו |  |
| 6:9.441 | הרי |  |
| 6:9.442 | זה |  |
| 6:9.443 | בדאי | . |
| 6:9.444 | לפי |  |
| 6:9.445 | שדרך |  |
| 6:9.446 | הכתובים |  |
| 6:9.447 | לדבר |  |
| 6:9.448 | בדרך |  |
| 6:9.449 | הרחבת |  |
| 6:9.450 | הלשון | , |
| 6:9.451 | או |  |
| 6:9.452 | בלשון |  |
| 6:9.453 | מושאל | , |
| 6:9.454 | או |  |
| 6:9.455 | משותף | , |
| 6:9.456 | או |  |
| 6:9.457 | בלשון |  |
| 6:9.458 | קצר | , |
| 6:9.459 | או |  |
| 6:9.460 | במוקדם |  |
| 6:9.461 | ומאוחר | . |
| 6:9.462 | צריך |  |
| 6:9.463 | המתרגם |  |
| 6:9.464 | לתרגמו |  |
| 6:9.465 | כמו |  |
| 6:9.466 | שקבלו |  |
| 6:9.467 | חכמים | , |
| 6:9.468 | ואלו |  |
| 6:9.469 | הן |  |
| 6:9.470 | סודות |  |
| 6:9.471 | התורה |  |
| 6:9.472 | לפרש |  |
| 6:9.473 | מקראות |  |
| 6:9.474 | הסתומים |  |
| 6:9.475 | כפי |  |
| 6:9.476 | האמת |  |
| 6:9.477 | שקבלו |  |
| 6:9.478 | חכמים | . |
| 6:9.479 | ולכן |  |
| 6:9.480 | אמרו |  |
| 6:9.481 | רז"ל |  |
| 6:9.482 | כי |  |
| 6:9.483 | כשתרגם |  |
| 6:9.484 | יונתן |  |
| 6:9.485 | בן |  |
| 6:9.486 | עוזיאל |  |
| 6:9.487 | את |  |
| 6:9.488 | הנביאים |  |
| 6:9.489 | יצאה |  |
| 6:9.490 | בת |  |
| 6:9.491 | קול |  |
| 6:9.492 | ואמרה |  |
| 6:9.493 | מי |  |
| 6:9.494 | גלה |  |
| 6:9.495 | רזי |  |
| 6:9.496 | לבני |  |
| 6:9.497 | אדם | . |
| 6:9.498 | ועליו |  |
| 6:9.499 | יאמר |  |
| 6:9.500 | רב |  |
| 6:9.501 | יוסף |  |
| 6:9.502 | בכמה |  |
| 6:9.503 | מקומות |  |
| 6:9.504 | אלמלא |  |
| 6:9.505 | תרגומא |  |
| 6:9.506 | דהאיי |  |
| 6:9.507 | קרא |  |
| 6:9.508 | לא |  |
| 6:9.509 | ידענא |  |
| 6:9.510 | מאי |  |
| 6:9.511 | קאמר | . |
| 6:9.512 | ובמדרש |  |
| 6:9.513 | הגדול |  |
| 6:9.514 | פ | ' |
| 6:9.515 | משפטים |  |
| 6:9.516 | ויראו |  |
| 6:9.517 | את |  |
| 6:9.518 | אלהי |  |
| 6:9.519 | ישראל | . |
| 6:9.520 | אמר |  |
| 6:9.521 | ר | ' |
| 6:9.522 | אליעזר |  |
| 6:9.523 | כל |  |
| 6:9.524 | המתרגם |  |
| 6:9.525 | פסוק |  |
| 6:9.526 | כצורתו |  |
| 6:9.527 | ה"ז |  |
| 6:9.528 | בדאי |  |
| 6:9.529 | וכל |  |
| 6:9.530 | המוסיף |  |
| 6:9.531 | בו |  |
| 6:9.532 | הרי |  |
| 6:9.533 | זה |  |
| 6:9.534 | מחרף |  |
| 6:9.535 | ומגדף |  |
| 6:9.536 | כגון |  |
| 6:9.537 | ויראו |  |
| 6:9.538 | את |  |
| 6:9.539 | אלהי |  |
| 6:9.540 | ישראל | . |
| 6:9.541 | וחזו |  |
| 6:9.542 | ית |  |
| 6:9.543 | אלהא |  |
| 6:9.544 | דישראל |  |
| 6:9.545 | ה"ז |  |
| 6:9.546 | בדאי | . |
| 6:9.547 | שהקב"ה |  |
| 6:9.548 | רואה |  |
| 6:9.549 | ואינו |  |
| 6:9.550 | נראה | . |
| 6:9.551 | ואם |  |
| 6:9.552 | תרגם |  |
| 6:9.553 | וחזו |  |
| 6:9.554 | ית |  |
| 6:9.555 | יקר |  |
| 6:9.556 | שכינת |  |
| 6:9.557 | אלהא |  |
| 6:9.558 | דישראל |  |
| 6:9.559 | הרי |  |
| 6:9.560 | זה |  |
| 6:9.561 | מחרף |  |
| 6:9.562 | ומגדף | . |
| 6:9.563 | שהוא |  |
| 6:9.564 | עושה |  |
| 6:9.565 | כאן |  |
| 6:9.566 | שלשה | , |
| 6:9.567 | יקר | , |
| 6:9.568 | ושכינה |  |
| 6:9.569 | ואל |  |
| 6:9.570 | עכ"ל | . |
| 6:9.571 | וקאי |  |
| 6:9.572 | לי |  |
| 6:9.573 | טובא |  |
| 6:9.574 | שהרי |  |
| 6:9.575 | יונתן |  |
| 6:9.576 | בן |  |
| 6:9.577 | עוזיאל |  |
| 6:9.578 | תרגם |  |
| 6:9.579 | כי |  |
| 6:9.580 | את |  |
| 6:9.581 | המלך |  |
| 6:9.582 | ה | ' |
| 6:9.583 | צבאות |  |
| 6:9.584 | ראו |  |
| 6:9.585 | עיני | . |
| 6:9.586 | ארי |  |
| 6:9.587 | ית |  |
| 6:9.588 | יקר |  |
| 6:9.589 | שכינת |  |
| 6:9.590 | מלך |  |
| 6:9.591 | עלמיא |  |
| 6:9.592 | ה | ' |
| 6:9.593 | צבאות |  |
| 6:9.594 | חזאה |  |
| 6:9.595 | עיני |  |
| 6:9.596 | והרי"ף |  |
| 6:9.597 | בסוף |  |
| 6:9.598 | מסכת |  |
| 6:9.599 | מגלה |  |
| 6:9.600 | הביא |  |
| 6:9.601 | תוספתא |  |
| 6:9.602 | זו |  |
| 6:9.603 | ופי | ' |
| 6:9.604 | ר | ' |
| 6:9.605 | נסים |  |
| 6:9.606 | המתרגם |  |
| 6:9.607 | פסוק |  |
| 6:9.608 | כצורתו |  |
| 6:9.609 | וחזו |  |
| 6:9.610 | ית | ' |
| 6:9.611 | מלאכא |  |
| 6:9.612 | דאלהא |  |
| 6:9.613 | דישראל | , |
| 6:9.614 | ורחות |  |
| 6:9.615 | כרסי |  |
| 6:9.616 | יקריה |  |
| 6:9.617 | ה"ז |  |
| 6:9.618 | מגדף |  |
| 6:9.619 | שהשוה |  |
| 6:9.620 | כבוד |  |
| 6:9.621 | עבד |  |
| 6:9.622 | לכבוד |  |
| 6:9.623 | רבו |  |
| 6:9.624 | ע"כ | , |
| 6:9.625 | והערוך |  |
| 6:9.626 | ערך |  |
| 6:9.627 | תרגם |  |
| 6:9.628 | פי | ' |
| 6:9.629 | הרי |  |
| 6:9.630 | מחרף |  |
| 6:9.631 | ומגדף |  |
| 6:9.632 | שעושה |  |
| 6:9.633 | הכבוד |  |
| 6:9.634 | מלאך | , |
| 6:9.635 | אלא |  |
| 6:9.636 | יתרגם |  |
| 6:9.637 | וחזו |  |
| 6:9.638 | ית |  |
| 6:9.639 | יקר |  |
| 6:9.640 | אלהא |  |
| 6:9.641 | דישראל |  |
| 6:9.642 | וכו | '. |
| 6:9.643 | וכנגד |  |
| 6:9.644 | אמונת |  |
| 6:9.645 | השלוש |  |
| 6:9.646 | שנתחדשה |  |
| 6:9.647 | בזמן |  |
| 6:9.648 | רז"ל |  |
| 6:9.649 | בעלי |  |
| 6:9.650 | המשנה |  |
| 6:9.651 | מתלמידי |  |
| 6:9.652 | ישוע |  |
| 6:9.653 | הנצרי |  |
| 6:9.654 | אמרו |  |
| 6:9.655 | רז"ל |  |
| 6:9.656 | במדרש |  |
| 6:9.657 | רבה |  |
| 6:9.658 | פרשת |  |
| 6:9.659 | יתרו |  |
| 6:9.660 | אנכי |  |
| 6:9.661 | ה | ' |
| 6:9.662 | אלהיך |  |
| 6:9.663 | אשר |  |
| 6:9.664 | הוציאתיך |  |
| 6:9.665 | וגו | ' |
| 6:9.666 | אמר |  |
| 6:9.667 | הקב"ה |  |
| 6:9.668 | מלך |  |
| 6:9.669 | בשר |  |
| 6:9.670 | ודם |  |
| 6:9.671 | מולך |  |
| 6:9.672 | יש |  |
| 6:9.673 | לו |  |
| 6:9.674 | אב |  |
| 6:9.675 | או |  |
| 6:9.676 | אח | , |
| 6:9.677 | או |  |
| 6:9.678 | בן | , |
| 6:9.679 | אני |  |
| 6:9.680 | איני |  |
| 6:9.681 | כן | , |
| 6:9.682 | אני |  |
| 6:9.683 | ראשון |  |
| 6:9.684 | שאין |  |
| 6:9.685 | לי |  |
| 6:9.686 | אב | , |
| 6:9.687 | ואני |  |
| 6:9.688 | אחרון |  |
| 6:9.689 | שאין |  |
| 6:9.690 | לי |  |
| 6:9.691 | בן | , |
| 6:9.692 | ומבלעדי |  |
| 6:9.693 | אין |  |
| 6:9.694 | אלהים |  |
| 6:9.695 | שאין |  |
| 6:9.696 | לי |  |
| 6:9.697 | אח |  |
| 6:9.698 | ע"כ | . |
| 6:9.699 | והמקובלים |  |
| 6:9.700 | החדשים |  |
| 6:9.701 | נמשכו |  |
| 6:9.702 | אחר |  |
| 6:9.703 | דעת |  |
| 6:9.704 | בעלי |  |
| 6:9.705 | השלוש |  |
| 6:9.706 | המאמינים |  |
| 6:9.707 | באב |  |
| 6:9.708 | ובן |  |
| 6:9.709 | ורוח |  |
| 6:9.710 | הקודש |  |
| 6:9.711 | כמובן |  |
| 6:9.712 | מהמשל |  |
| 6:9.713 | המבואר |  |
| 6:9.714 | בס | ' |
| 6:9.715 | עז |  |
| 6:9.716 | לאלהים |  |
| 6:9.717 | בבית |  |
| 6:9.718 | קודש |  |
| 6:9.719 | הקדשים |  |
| 6:9.720 | דף |  |
| 6:9.721 | נ"ט |  |
| 6:9.722 | ע"ב |  |
| 6:9.723 | וז"ל |  |
| 6:9.724 | ולפי |  |
| 6:9.725 | שמלכא |  |
| 6:9.726 | קדישא |  |
| 6:9.727 | דכל |  |
| 6:9.728 | קדישין |  |
| 6:9.729 | לא |  |
| 6:9.730 | מתפרש |  |
| 6:9.731 | מניה |  |
| 6:9.732 | לעלמין |  |
| 6:9.733 | והוא |  |
| 6:9.734 | מאיר |  |
| 6:9.735 | בזעיר |  |
| 6:9.736 | נאפין |  |
| 6:9.737 | דהוא |  |
| 6:9.738 | לב |  |
| 6:9.739 | האצילות |  |
| 6:9.740 | לכן |  |
| 6:9.741 | הוא |  |
| 6:9.742 | ג"כ |  |
| 6:9.743 | יקרא |  |
| 6:9.744 | בוצינא |  |
| 6:9.745 | דקרדינותא |  |
| 6:9.746 | שמאיר |  |
| 6:9.747 | בלב |  |
| 6:9.748 | שהוא |  |
| 6:9.749 | זעיר |  |
| 6:9.750 | אנפין |  |
| 6:9.751 | דאצילות | . |
| 6:9.752 | נמצא |  |
| 6:9.753 | שמה |  |
| 6:9.754 | שנקרא |  |
| 6:9.755 | זה |  |
| 6:9.756 | השם |  |
| 6:9.757 | הוא |  |
| 6:9.758 | על |  |
| 6:9.759 | שם |  |
| 6:9.760 | נשמתו |  |
| 6:9.761 | ולא |  |
| 6:9.762 | על |  |
| 6:9.763 | שם |  |
| 6:9.764 | מלבושו |  |
| 6:9.765 | כשאר |  |
| 6:9.766 | השמות | . |
| 6:9.767 | והוא |  |
| 6:9.768 | הנקרא |  |
| 6:9.769 | מלך |  |
| 6:9.770 | לבני |  |
| 6:9.771 | ישראל |  |
| 6:9.772 | כאשר |  |
| 6:9.773 | נתבאר |  |
| 6:9.774 | בפרשת |  |
| 6:9.775 | ואלה |  |
| 6:9.776 | המלכים | . |
| 6:9.777 | והוא |  |
| 6:9.778 | מלכינו |  |
| 6:9.779 | ואנחנו |  |
| 6:9.780 | עבדיו | . |
| 6:9.781 | ואליו |  |
| 6:9.782 | תפלותינו |  |
| 6:9.783 | וכוונותינו | . |
| 6:9.784 | ואפי | ' |
| 6:9.785 | אבא |  |
| 6:9.786 | ואימא |  |
| 6:9.787 | מוזהרים |  |
| 6:9.788 | לכבד |  |
| 6:9.789 | את |  |
| 6:9.790 | הבן |  |
| 6:9.791 | הזה |  |
| 6:9.792 | וליראה |  |
| 6:9.793 | ממנו | , |
| 6:9.794 | לפי |  |
| 6:9.795 | שהוא |  |
| 6:9.796 | המלך |  |
| 6:9.797 | השליט |  |
| 6:9.798 | על |  |
| 6:9.799 | כלא | . |
| 6:9.800 | ואפי | ' |
| 6:9.801 | שיצא |  |
| 6:9.802 | הבן |  |
| 6:9.803 | הקדוש |  |
| 6:9.804 | הזה |  |
| 6:9.805 | מאבא |  |
| 6:9.806 | ואימא |  |
| 6:9.807 | יתבאר |  |
| 6:9.808 | לקטן |  |
| 6:9.809 | שהוא |  |
| 6:9.810 | העיקר | , |
| 6:9.811 | אבל |  |
| 6:9.812 | בכל |  |
| 6:9.813 | זאת |  |
| 6:9.814 | אמשול |  |
| 6:9.815 | לך |  |
| 6:9.816 | משל |  |
| 6:9.817 | אם |  |
| 6:9.818 | ימלוך |  |
| 6:9.819 | א | ' |
| 6:9.820 | ויהיה |  |
| 6:9.821 | לו |  |
| 6:9.822 | אב |  |
| 6:9.823 | ואם | . |
| 6:9.824 | למי |  |
| 6:9.825 | צריכים |  |
| 6:9.826 | לעבוד | ? |
| 6:9.827 | בודאי |  |
| 6:9.828 | למלך | . |
| 6:9.829 | ואפי | ' |
| 6:9.830 | אביו |  |
| 6:9.831 | ואמו |  |
| 6:9.832 | צריכים |  |
| 6:9.833 | לו |  |
| 6:9.834 | לפי |  |
| 6:9.835 | שהוא |  |
| 6:9.836 | מלך | . |
| 6:9.837 | על |  |
| 6:9.838 | זה |  |
| 6:9.839 | אמר |  |
| 6:9.840 | נשקו |  |
| 6:9.841 | בר |  |
| 6:9.842 | פן |  |
| 6:9.843 | יאנף | . |
| 6:9.844 | ואמר |  |
| 6:9.845 | בזהר |  |
| 6:9.846 | נשקו |  |
| 6:9.847 | ידין |  |
| 6:9.848 | להאי |  |
| 6:9.849 | ברא |  |
| 6:9.850 | קדישא | . |
| 6:9.851 | כי |  |
| 6:9.852 | שולטנו |  |
| 6:9.853 | אתיהיב |  |
| 6:9.854 | ליה |  |
| 6:9.855 | על |  |
| 6:9.856 | כלא |  |
| 6:9.857 | וכו | ' |
| 6:9.858 | הרי |  |
| 6:9.859 | שאפי | ' |
| 6:9.860 | אבא |  |
| 6:9.861 | ואימא |  |
| 6:9.862 | צריכים |  |
| 6:9.863 | לו |  |
| 6:9.864 | עכ"ל | . |
| 6:9.865 | הרי |  |
| 6:9.866 | בהדיא |  |
| 6:9.867 | שהוא |  |
| 6:9.868 | זותר |  |
| 6:9.869 | לדברי |  |
| 6:9.870 | רז"ל |  |
| 6:9.871 | במדרש |  |
| 6:9.872 | שמות |  |
| 6:9.873 | רבה |  |
| 6:9.874 | הנז | ' |
| 6:9.875 | לעיל | , |
| 6:9.876 | וירושלמי |  |
| 6:9.877 | בפ | ' |
| 6:9.878 | דיני |  |
| 6:9.879 | ממונות | , |
| 6:9.880 | דגרסינן |  |
| 6:9.881 | התם |  |
| 6:9.882 | מהו |  |
| 6:9.883 | חותמו |  |
| 6:9.884 | של |  |
| 6:9.885 | הקב"ה |  |
| 6:9.886 | רב |  |
| 6:9.887 | בבי |  |
| 6:9.888 | בשם |  |
| 6:9.889 | ר |  |
| 6:9.890 | ראובן |  |
| 6:9.891 | אמת | . |
| 6:9.892 | מהו |  |
| 6:9.893 | אמת | . |
| 6:9.894 | אמר |  |
| 6:9.895 | ר | ' |
| 6:9.896 | בון |  |
| 6:9.897 | שהוא |  |
| 6:9.898 | אלהים |  |
| 6:9.899 | חיים |  |
| 6:9.900 | ומלך |  |
| 6:9.901 | עולם | . |
| 6:9.902 | אמר |  |
| 6:9.903 | ריש |  |
| 6:9.904 | לקיש |  |
| 6:9.905 | אלף |  |
| 6:9.906 | רישא |  |
| 6:9.907 | דאלף |  |
| 6:9.908 | בית |  |
| 6:9.909 | מים |  |
| 6:9.910 | באמצעיתא | , |
| 6:9.911 | תיו |  |
| 6:9.912 | בסופא | . |
| 6:9.913 | לומר |  |
| 6:9.914 | אני |  |
| 6:9.915 | ה | ' |
| 6:9.916 | ראשון |  |
| 6:9.917 | שלא |  |
| 6:9.918 | קבלתי |  |
| 6:9.919 | מאחר | . |
| 6:9.920 | ומבלעדי |  |
| 6:9.921 | אין |  |
| 6:9.922 | אלהים |  |
| 6:9.923 | שאין |  |
| 6:9.924 | לי |  |
| 6:9.925 | שותף | . |
| 6:9.926 | ואת |  |
| 6:9.927 | אחרונים |  |
| 6:9.928 | אני |  |
| 6:9.929 | הוא |  |
| 6:9.930 | שאיני |  |
| 6:9.931 | עתיד |  |
| 6:9.932 | למוסרה |  |
| 6:9.933 | לאחר |  |
| 6:9.934 | עכ"ל |  |
| 6:9.935 | הירוש | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

