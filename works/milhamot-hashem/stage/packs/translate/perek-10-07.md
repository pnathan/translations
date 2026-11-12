# Translation pack: milhamot-hashem perek-10-07 (Milchamot ha-Shem 10:8–10:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-07`.

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
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **השונים** → "the Dualists".
- **האמגושים** → "the Magians".
- **מערב** beside Ashkenaz, Sefarad, Yemen → "the Maghreb" (note: "the West").
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
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
- אל **אחר** → "another god".
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 10:8

Hebrew (Wikisource): העתקת לשון ר' חטר בסי' מ״ב מערבי ללשון הקדש היסוד השני יחודו ית' וזהו לידע שזה האחד שהוא סבת הכל, אחד לא כאחד הסוג, ולא כאחד המין, ולא כאחד המורכב אשר יחולק לאחדים רבים, ולא כגוף הפשוט שהוא אחד במנין (כמו המים שהוא יסוד פשוט) (חטר) ומתחלק לאין קץ. אבל הוא ית' אחד אחדות שאין כמוה אחדות בשום פנים, וזה היסוד השני הוא הנלמד מאמרו שמע ישראל ה' אלהינו ה' אחד! ביאור היסוד הזה הואיל וכבר ביאר רבי' מצויאותו ית'. הוצרך לבאר אחדותו. כי אלו השני יסודות הם עיקר תורת משה רבינו ע״ה. והמה ראשית כל תורה אלהית, וסוף כל תורה נימוסית. כי התורה האלהית מתחלתאמן האלהיות וגומרת בגשמיות. והתורה הנימוסית הפך זה. ביאור זה שהאדם אם ידע מציאות כל הנמצאות מראש ועד סוף רצוני לומר ממעלה מעלה עד מטה מטה. כל עוד שירד בידיעתו מטה נתמעטה האחדות ונתוסף הרבוי עד שמגיע לעולם השפל שבלי ספק לא ימד ולא יספר סוגי מיניו כל שכן מיני אישי כל מין ומין וכו'. אבל אם בא לעלות בנמצאות ממטה למעלה, הנה כל מה שיוסיף לעלות ימעט הרבוי ותתחזק האחדות עד האחד המיוחד שממנו נמצאו כל האחדים שלכל הנמצאות. והוא ית' אחד מיוחד. שלוש אמונות יש, השונים. והאמגושים. והנצרים. והנני מבאר כל אחת מהם. ומי שיעיען בספריהם יתבאר לו זה. השונים אומרים שיש לעולם שני פועלים אחד מהם האור. והוא פועל הטוב בנמצאים. ולא יוכל לעשות רע. והפועל הב' החשך והוא פועל הרע בעולם ולא יוכל לעשות טוב, וחמה חלוקים. יש מהם אומרים שהאור והחשך חיים. ויש אומרים שהם מתים. אמנם האמגוש הנם אומרים שיש לעולם פועל חכם. גבור מנהיג ואינו גוף. והרחיקו ממנו מעשות כל רע. וקראוהו יזדאן. ויחסו כל הרע והצער ההוים בעולם לשטן. וקראוהו אַהְרְמִן בלשונם והמה חלוקים יש מהם אומרים שהוא קדמון ואינו גוף. ומהם אומרים שהוא מחודש וגוף. ונחלקו ג״כ בסברת חדושו ואמרול שאירע ליזדאן מחשבה ואמר בלבו איך יהיה עניינו. ומי זה יחלוק עלי (יריב עמי) ואמרו שהשטן נהיה מזאת המחשבה. אחר כן השלים עם יזדאן ונתפשרו יחד עד גמירא. וזכרו עוד דברים שאין כאן מקום לזכרם. והטוב אצל בעלי שתי האמונות ר״ל השונים והאמגוש. הוא מה שתתאוה לו הנפש. והרע הוא מה שתתעבהו הנפש. אמנם הנצרים אומרים שהשי״ת ג' אקאנים כפי האמת (אפשר שהוא כמו ספירות) אקנום האב והוא עצם הבורא אצלם (אריך) ואקנום הבן והוא אלעלם (חכמה) ואקנום רוח הקדש, והוא החיים (בינה אלהים חיים). והשלשה אחד, והם כת מישראל. וראייתם שהם ג' ונחשבים אחד מאמרו ית' שמע ישראל ה' אלהינו אחד. אמרו הרי אלו שלשה שמות לשלשה נקראים בשם. אמנם השלשה חשובים אחד. ולפי הנראה שהם מאמינים עצם אחד ושני תארים קדומים. ואחר שזכרתי לך אמונות החיצונים החולקים בייחוד כדי שתזהר מהם אשוב לבאר עניין היסוד הזה וכו' עד כאן לשון רבי חטר ז״ל

| id | word | punctuation after |
|---|---|---|
| 10:8.1 | העתקת |  |
| 10:8.2 | לשון |  |
| 10:8.3 | ר | ' |
| 10:8.4 | חטר |  |
| 10:8.5 | בסי | ' |
| 10:8.6 | מ״ב |  |
| 10:8.7 | מערבי |  |
| 10:8.8 | ללשון |  |
| 10:8.9 | הקדש |  |
| 10:8.10 | היסוד |  |
| 10:8.11 | השני |  |
| 10:8.12 | יחודו |  |
| 10:8.13 | ית | ' |
| 10:8.14 | וזהו |  |
| 10:8.15 | לידע |  |
| 10:8.16 | שזה |  |
| 10:8.17 | האחד |  |
| 10:8.18 | שהוא |  |
| 10:8.19 | סבת |  |
| 10:8.20 | הכל | , |
| 10:8.21 | אחד |  |
| 10:8.22 | לא |  |
| 10:8.23 | כאחד |  |
| 10:8.24 | הסוג | , |
| 10:8.25 | ולא |  |
| 10:8.26 | כאחד |  |
| 10:8.27 | המין | , |
| 10:8.28 | ולא |  |
| 10:8.29 | כאחד |  |
| 10:8.30 | המורכב |  |
| 10:8.31 | אשר |  |
| 10:8.32 | יחולק |  |
| 10:8.33 | לאחדים |  |
| 10:8.34 | רבים | , |
| 10:8.35 | ולא |  |
| 10:8.36 | כגוף |  |
| 10:8.37 | הפשוט |  |
| 10:8.38 | שהוא |  |
| 10:8.39 | אחד |  |
| 10:8.40 | במנין |  |
| 10:8.41 | כמו |  |
| 10:8.42 | המים |  |
| 10:8.43 | שהוא |  |
| 10:8.44 | יסוד |  |
| 10:8.45 | פשוט | ) |
| 10:8.46 | חטר | ) |
| 10:8.47 | ומתחלק |  |
| 10:8.48 | לאין |  |
| 10:8.49 | קץ | . |
| 10:8.50 | אבל |  |
| 10:8.51 | הוא |  |
| 10:8.52 | ית | ' |
| 10:8.53 | אחד |  |
| 10:8.54 | אחדות |  |
| 10:8.55 | שאין |  |
| 10:8.56 | כמוה |  |
| 10:8.57 | אחדות |  |
| 10:8.58 | בשום |  |
| 10:8.59 | פנים | , |
| 10:8.60 | וזה |  |
| 10:8.61 | היסוד |  |
| 10:8.62 | השני |  |
| 10:8.63 | הוא |  |
| 10:8.64 | הנלמד |  |
| 10:8.65 | מאמרו |  |
| 10:8.66 | שמע |  |
| 10:8.67 | ישראל |  |
| 10:8.68 | ה | ' |
| 10:8.69 | אלהינו |  |
| 10:8.70 | ה | ' |
| 10:8.71 | אחד | ! |
| 10:8.72 | ביאור |  |
| 10:8.73 | היסוד |  |
| 10:8.74 | הזה |  |
| 10:8.75 | הואיל |  |
| 10:8.76 | וכבר |  |
| 10:8.77 | ביאר |  |
| 10:8.78 | רבי | ' |
| 10:8.79 | מצויאותו |  |
| 10:8.80 | ית | '. |
| 10:8.81 | הוצרך |  |
| 10:8.82 | לבאר |  |
| 10:8.83 | אחדותו | . |
| 10:8.84 | כי |  |
| 10:8.85 | אלו |  |
| 10:8.86 | השני |  |
| 10:8.87 | יסודות |  |
| 10:8.88 | הם |  |
| 10:8.89 | עיקר |  |
| 10:8.90 | תורת |  |
| 10:8.91 | משה |  |
| 10:8.92 | רבינו |  |
| 10:8.93 | ע״ה | . |
| 10:8.94 | והמה |  |
| 10:8.95 | ראשית |  |
| 10:8.96 | כל |  |
| 10:8.97 | תורה |  |
| 10:8.98 | אלהית | , |
| 10:8.99 | וסוף |  |
| 10:8.100 | כל |  |
| 10:8.101 | תורה |  |
| 10:8.102 | נימוסית | . |
| 10:8.103 | כי |  |
| 10:8.104 | התורה |  |
| 10:8.105 | האלהית |  |
| 10:8.106 | מתחלתאמן |  |
| 10:8.107 | האלהיות |  |
| 10:8.108 | וגומרת |  |
| 10:8.109 | בגשמיות | . |
| 10:8.110 | והתורה |  |
| 10:8.111 | הנימוסית |  |
| 10:8.112 | הפך |  |
| 10:8.113 | זה | . |
| 10:8.114 | ביאור |  |
| 10:8.115 | זה |  |
| 10:8.116 | שהאדם |  |
| 10:8.117 | אם |  |
| 10:8.118 | ידע |  |
| 10:8.119 | מציאות |  |
| 10:8.120 | כל |  |
| 10:8.121 | הנמצאות |  |
| 10:8.122 | מראש |  |
| 10:8.123 | ועד |  |
| 10:8.124 | סוף |  |
| 10:8.125 | רצוני |  |
| 10:8.126 | לומר |  |
| 10:8.127 | ממעלה |  |
| 10:8.128 | מעלה |  |
| 10:8.129 | עד |  |
| 10:8.130 | מטה |  |
| 10:8.131 | מטה | . |
| 10:8.132 | כל |  |
| 10:8.133 | עוד |  |
| 10:8.134 | שירד |  |
| 10:8.135 | בידיעתו |  |
| 10:8.136 | מטה |  |
| 10:8.137 | נתמעטה |  |
| 10:8.138 | האחדות |  |
| 10:8.139 | ונתוסף |  |
| 10:8.140 | הרבוי |  |
| 10:8.141 | עד |  |
| 10:8.142 | שמגיע |  |
| 10:8.143 | לעולם |  |
| 10:8.144 | השפל |  |
| 10:8.145 | שבלי |  |
| 10:8.146 | ספק |  |
| 10:8.147 | לא |  |
| 10:8.148 | ימד |  |
| 10:8.149 | ולא |  |
| 10:8.150 | יספר |  |
| 10:8.151 | סוגי |  |
| 10:8.152 | מיניו |  |
| 10:8.153 | כל |  |
| 10:8.154 | שכן |  |
| 10:8.155 | מיני |  |
| 10:8.156 | אישי |  |
| 10:8.157 | כל |  |
| 10:8.158 | מין |  |
| 10:8.159 | ומין |  |
| 10:8.160 | וכו | '. |
| 10:8.161 | אבל |  |
| 10:8.162 | אם |  |
| 10:8.163 | בא |  |
| 10:8.164 | לעלות |  |
| 10:8.165 | בנמצאות |  |
| 10:8.166 | ממטה |  |
| 10:8.167 | למעלה | , |
| 10:8.168 | הנה |  |
| 10:8.169 | כל |  |
| 10:8.170 | מה |  |
| 10:8.171 | שיוסיף |  |
| 10:8.172 | לעלות |  |
| 10:8.173 | ימעט |  |
| 10:8.174 | הרבוי |  |
| 10:8.175 | ותתחזק |  |
| 10:8.176 | האחדות |  |
| 10:8.177 | עד |  |
| 10:8.178 | האחד |  |
| 10:8.179 | המיוחד |  |
| 10:8.180 | שממנו |  |
| 10:8.181 | נמצאו |  |
| 10:8.182 | כל |  |
| 10:8.183 | האחדים |  |
| 10:8.184 | שלכל |  |
| 10:8.185 | הנמצאות | . |
| 10:8.186 | והוא |  |
| 10:8.187 | ית | ' |
| 10:8.188 | אחד |  |
| 10:8.189 | מיוחד | . |
| 10:8.190 | שלוש |  |
| 10:8.191 | אמונות |  |
| 10:8.192 | יש | , |
| 10:8.193 | השונים | . |
| 10:8.194 | והאמגושים | . |
| 10:8.195 | והנצרים | . |
| 10:8.196 | והנני |  |
| 10:8.197 | מבאר |  |
| 10:8.198 | כל |  |
| 10:8.199 | אחת |  |
| 10:8.200 | מהם | . |
| 10:8.201 | ומי |  |
| 10:8.202 | שיעיען |  |
| 10:8.203 | בספריהם |  |
| 10:8.204 | יתבאר |  |
| 10:8.205 | לו |  |
| 10:8.206 | זה | . |
| 10:8.207 | השונים |  |
| 10:8.208 | אומרים |  |
| 10:8.209 | שיש |  |
| 10:8.210 | לעולם |  |
| 10:8.211 | שני |  |
| 10:8.212 | פועלים |  |
| 10:8.213 | אחד |  |
| 10:8.214 | מהם |  |
| 10:8.215 | האור | . |
| 10:8.216 | והוא |  |
| 10:8.217 | פועל |  |
| 10:8.218 | הטוב |  |
| 10:8.219 | בנמצאים | . |
| 10:8.220 | ולא |  |
| 10:8.221 | יוכל |  |
| 10:8.222 | לעשות |  |
| 10:8.223 | רע | . |
| 10:8.224 | והפועל |  |
| 10:8.225 | הב | ' |
| 10:8.226 | החשך |  |
| 10:8.227 | והוא |  |
| 10:8.228 | פועל |  |
| 10:8.229 | הרע |  |
| 10:8.230 | בעולם |  |
| 10:8.231 | ולא |  |
| 10:8.232 | יוכל |  |
| 10:8.233 | לעשות |  |
| 10:8.234 | טוב | , |
| 10:8.235 | וחמה |  |
| 10:8.236 | חלוקים | . |
| 10:8.237 | יש |  |
| 10:8.238 | מהם |  |
| 10:8.239 | אומרים |  |
| 10:8.240 | שהאור |  |
| 10:8.241 | והחשך |  |
| 10:8.242 | חיים | . |
| 10:8.243 | ויש |  |
| 10:8.244 | אומרים |  |
| 10:8.245 | שהם |  |
| 10:8.246 | מתים | . |
| 10:8.247 | אמנם |  |
| 10:8.248 | האמגוש |  |
| 10:8.249 | הנם |  |
| 10:8.250 | אומרים |  |
| 10:8.251 | שיש |  |
| 10:8.252 | לעולם |  |
| 10:8.253 | פועל |  |
| 10:8.254 | חכם | . |
| 10:8.255 | גבור |  |
| 10:8.256 | מנהיג |  |
| 10:8.257 | ואינו |  |
| 10:8.258 | גוף | . |
| 10:8.259 | והרחיקו |  |
| 10:8.260 | ממנו |  |
| 10:8.261 | מעשות |  |
| 10:8.262 | כל |  |
| 10:8.263 | רע | . |
| 10:8.264 | וקראוהו |  |
| 10:8.265 | יזדאן | . |
| 10:8.266 | ויחסו |  |
| 10:8.267 | כל |  |
| 10:8.268 | הרע |  |
| 10:8.269 | והצער |  |
| 10:8.270 | ההוים |  |
| 10:8.271 | בעולם |  |
| 10:8.272 | לשטן | . |
| 10:8.273 | וקראוהו |  |
| 10:8.274 | אַהְרְמִן |  |
| 10:8.275 | בלשונם |  |
| 10:8.276 | והמה |  |
| 10:8.277 | חלוקים |  |
| 10:8.278 | יש |  |
| 10:8.279 | מהם |  |
| 10:8.280 | אומרים |  |
| 10:8.281 | שהוא |  |
| 10:8.282 | קדמון |  |
| 10:8.283 | ואינו |  |
| 10:8.284 | גוף | . |
| 10:8.285 | ומהם |  |
| 10:8.286 | אומרים |  |
| 10:8.287 | שהוא |  |
| 10:8.288 | מחודש |  |
| 10:8.289 | וגוף | . |
| 10:8.290 | ונחלקו |  |
| 10:8.291 | ג״כ |  |
| 10:8.292 | בסברת |  |
| 10:8.293 | חדושו |  |
| 10:8.294 | ואמרול |  |
| 10:8.295 | שאירע |  |
| 10:8.296 | ליזדאן |  |
| 10:8.297 | מחשבה |  |
| 10:8.298 | ואמר |  |
| 10:8.299 | בלבו |  |
| 10:8.300 | איך |  |
| 10:8.301 | יהיה |  |
| 10:8.302 | עניינו | . |
| 10:8.303 | ומי |  |
| 10:8.304 | זה |  |
| 10:8.305 | יחלוק |  |
| 10:8.306 | עלי |  |
| 10:8.307 | יריב |  |
| 10:8.308 | עמי | ) |
| 10:8.309 | ואמרו |  |
| 10:8.310 | שהשטן |  |
| 10:8.311 | נהיה |  |
| 10:8.312 | מזאת |  |
| 10:8.313 | המחשבה | . |
| 10:8.314 | אחר |  |
| 10:8.315 | כן |  |
| 10:8.316 | השלים |  |
| 10:8.317 | עם |  |
| 10:8.318 | יזדאן |  |
| 10:8.319 | ונתפשרו |  |
| 10:8.320 | יחד |  |
| 10:8.321 | עד |  |
| 10:8.322 | גמירא | . |
| 10:8.323 | וזכרו |  |
| 10:8.324 | עוד |  |
| 10:8.325 | דברים |  |
| 10:8.326 | שאין |  |
| 10:8.327 | כאן |  |
| 10:8.328 | מקום |  |
| 10:8.329 | לזכרם | . |
| 10:8.330 | והטוב |  |
| 10:8.331 | אצל |  |
| 10:8.332 | בעלי |  |
| 10:8.333 | שתי |  |
| 10:8.334 | האמונות |  |
| 10:8.335 | ר״ל |  |
| 10:8.336 | השונים |  |
| 10:8.337 | והאמגוש | . |
| 10:8.338 | הוא |  |
| 10:8.339 | מה |  |
| 10:8.340 | שתתאוה |  |
| 10:8.341 | לו |  |
| 10:8.342 | הנפש | . |
| 10:8.343 | והרע |  |
| 10:8.344 | הוא |  |
| 10:8.345 | מה |  |
| 10:8.346 | שתתעבהו |  |
| 10:8.347 | הנפש | . |
| 10:8.348 | אמנם |  |
| 10:8.349 | הנצרים |  |
| 10:8.350 | אומרים |  |
| 10:8.351 | שהשי״ת |  |
| 10:8.352 | ג | ' |
| 10:8.353 | אקאנים |  |
| 10:8.354 | כפי |  |
| 10:8.355 | האמת |  |
| 10:8.356 | אפשר |  |
| 10:8.357 | שהוא |  |
| 10:8.358 | כמו |  |
| 10:8.359 | ספירות | ) |
| 10:8.360 | אקנום |  |
| 10:8.361 | האב |  |
| 10:8.362 | והוא |  |
| 10:8.363 | עצם |  |
| 10:8.364 | הבורא |  |
| 10:8.365 | אצלם |  |
| 10:8.366 | אריך | ) |
| 10:8.367 | ואקנום |  |
| 10:8.368 | הבן |  |
| 10:8.369 | והוא |  |
| 10:8.370 | אלעלם |  |
| 10:8.371 | חכמה | ) |
| 10:8.372 | ואקנום |  |
| 10:8.373 | רוח |  |
| 10:8.374 | הקדש | , |
| 10:8.375 | והוא |  |
| 10:8.376 | החיים |  |
| 10:8.377 | בינה |  |
| 10:8.378 | אלהים |  |
| 10:8.379 | חיים | ). |
| 10:8.380 | והשלשה |  |
| 10:8.381 | אחד | , |
| 10:8.382 | והם |  |
| 10:8.383 | כת |  |
| 10:8.384 | מישראל | . |
| 10:8.385 | וראייתם |  |
| 10:8.386 | שהם |  |
| 10:8.387 | ג | ' |
| 10:8.388 | ונחשבים |  |
| 10:8.389 | אחד |  |
| 10:8.390 | מאמרו |  |
| 10:8.391 | ית | ' |
| 10:8.392 | שמע |  |
| 10:8.393 | ישראל |  |
| 10:8.394 | ה | ' |
| 10:8.395 | אלהינו |  |
| 10:8.396 | אחד | . |
| 10:8.397 | אמרו |  |
| 10:8.398 | הרי |  |
| 10:8.399 | אלו |  |
| 10:8.400 | שלשה |  |
| 10:8.401 | שמות |  |
| 10:8.402 | לשלשה |  |
| 10:8.403 | נקראים |  |
| 10:8.404 | בשם | . |
| 10:8.405 | אמנם |  |
| 10:8.406 | השלשה |  |
| 10:8.407 | חשובים |  |
| 10:8.408 | אחד | . |
| 10:8.409 | ולפי |  |
| 10:8.410 | הנראה |  |
| 10:8.411 | שהם |  |
| 10:8.412 | מאמינים |  |
| 10:8.413 | עצם |  |
| 10:8.414 | אחד |  |
| 10:8.415 | ושני |  |
| 10:8.416 | תארים |  |
| 10:8.417 | קדומים | . |
| 10:8.418 | ואחר |  |
| 10:8.419 | שזכרתי |  |
| 10:8.420 | לך |  |
| 10:8.421 | אמונות |  |
| 10:8.422 | החיצונים |  |
| 10:8.423 | החולקים |  |
| 10:8.424 | בייחוד |  |
| 10:8.425 | כדי |  |
| 10:8.426 | שתזהר |  |
| 10:8.427 | מהם |  |
| 10:8.428 | אשוב |  |
| 10:8.429 | לבאר |  |
| 10:8.430 | עניין |  |
| 10:8.431 | היסוד |  |
| 10:8.432 | הזה |  |
| 10:8.433 | וכו | ' |
| 10:8.434 | עד |  |
| 10:8.435 | כאן |  |
| 10:8.436 | לשון |  |
| 10:8.437 | רבי |  |
| 10:8.438 | חטר |  |
| 10:8.439 | ז״ל |  |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

