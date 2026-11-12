# Translation pack: milhamot-hashem perek-7-14 (Milchamot ha-Shem 7:19–7:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-14`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **מין / מינות** → "heretic / heresy".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
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


## Milchamot ha-Shem 7:19

Hebrew (Wikisource): צא) ומזה תבין ותשכיל גודל השגיאה והטעות שטעו ושגגו רבותינו בעלי הקבלה החדשה הנכריה. שמלבד שמנו באלהות כמה מדות ומדרגות זו למעלה מזו במקום ובזמן. וגם במעלה עד שאמרו שזה נוטל רשות מזה, והאלהים הנקבה נוטלת רשות מאלהים הזכר, לפי דעתם כמבואר בזהר בראשית דף כ"ב. ויחסו כל עבודותינו וכל תפלותינו וברכותינו למדרגה התחתונה מכולם, וקראו לו בשמו של הקב"ה ואמרו שהוא ה' אלהינו ואנחנו עמו וכו' ואותו נעבוד וכינו אותו בשם "זעיר אנפין", שענינו בלשון הכתוב קצר אפים. והנה כבר ידעת ונתבאר לך מכתבי הקדש גנות ופחיתות המדה המגונה הזאת, ואיך יתארו את האלהים אשר יעבדו אותו במדה המגונה הזאת, ועליהם אמר אליהו בן ברכאל הבוזי במליצתו, האמור לממך בליעל רשע אל נדיבים. ופי' רש"י ז"ל וכי הגון וראוי לומר למלך בליעל? ולבאר רשע אל נדיבים? אשר לא נשא פני שרים, כלומר למלך עולם אשר אין משוא פנים לפניו עכ"ל. וכן פירש רבי אברהם פריצול ז"ל האמור למלך וכו' אשר לא נשא פני שרים, האם ראוי לומר למלך הממונה על הנהגה כזאת שיהיה בליעל? ואם ראוי לומר רשע למי שהוא נדיב מנדיבי עמים ונודע גזעו וטובו? ואיך ראוי ליחס עול למלך צדיק וישר כזה שהוא ה' אלהינו אשר לא נשא פני שרים בענין הדין לכבדם ולשאת את פניהם, ולא נכר לפניו השוע והעשיר בהיותו טוען בדין עם העני לחוס על כבודו, בעבור כי מעשה ידיו כולם, ואין הבדל ביניהם מצד יצירתם לפני הקב"ה עכ"ל. ודון מינה ואוקי הכא, כיון שהם אומרים שזעיר אנפין הוא ה' אלהינו איך יקראו לו בשם קצר אפים אשר יעשה אולת? כי לא יאות לתאר את מלכו של עולם אלא בתארי שבח ועילוי, לא תארי אולת וכסילות כתאר קצר אפים (זעיר אנפין) שהוא עושה אולת, כמו שאי אפשר לומר למלכו של עולם בליעל כמאמר אליהו. ומזה תבין ותשכיל כי לא קבלה אמתית היא. רק נכריה היא. ולא מרז"ל בעלי המשנה והתלמוד יצאו הדברים האלה, חלילה לכל בעל אמונת תורת משה רע"ה וקבלת רבותינו ז"ל להאמין בדברים כאלה. כי כבר נתברר אצלינו גדולת חכמינו וטוב שכלם ועומק עיונם בחכמה, והם השיגו את האמת השגה גמורה ואמתית, ואמרו כל המשתף שם שמים ודבר אחר נעקר מן העולם.

| id | word | punctuation after |
|---|---|---|
| 7:19.1 | צא | ) |
| 7:19.2 | ומזה |  |
| 7:19.3 | תבין |  |
| 7:19.4 | ותשכיל |  |
| 7:19.5 | גודל |  |
| 7:19.6 | השגיאה |  |
| 7:19.7 | והטעות |  |
| 7:19.8 | שטעו |  |
| 7:19.9 | ושגגו |  |
| 7:19.10 | רבותינו |  |
| 7:19.11 | בעלי |  |
| 7:19.12 | הקבלה |  |
| 7:19.13 | החדשה |  |
| 7:19.14 | הנכריה | . |
| 7:19.15 | שמלבד |  |
| 7:19.16 | שמנו |  |
| 7:19.17 | באלהות |  |
| 7:19.18 | כמה |  |
| 7:19.19 | מדות |  |
| 7:19.20 | ומדרגות |  |
| 7:19.21 | זו |  |
| 7:19.22 | למעלה |  |
| 7:19.23 | מזו |  |
| 7:19.24 | במקום |  |
| 7:19.25 | ובזמן | . |
| 7:19.26 | וגם |  |
| 7:19.27 | במעלה |  |
| 7:19.28 | עד |  |
| 7:19.29 | שאמרו |  |
| 7:19.30 | שזה |  |
| 7:19.31 | נוטל |  |
| 7:19.32 | רשות |  |
| 7:19.33 | מזה | , |
| 7:19.34 | והאלהים |  |
| 7:19.35 | הנקבה |  |
| 7:19.36 | נוטלת |  |
| 7:19.37 | רשות |  |
| 7:19.38 | מאלהים |  |
| 7:19.39 | הזכר | , |
| 7:19.40 | לפי |  |
| 7:19.41 | דעתם |  |
| 7:19.42 | כמבואר |  |
| 7:19.43 | בזהר |  |
| 7:19.44 | בראשית |  |
| 7:19.45 | דף |  |
| 7:19.46 | כ"ב | . |
| 7:19.47 | ויחסו |  |
| 7:19.48 | כל |  |
| 7:19.49 | עבודותינו |  |
| 7:19.50 | וכל |  |
| 7:19.51 | תפלותינו |  |
| 7:19.52 | וברכותינו |  |
| 7:19.53 | למדרגה |  |
| 7:19.54 | התחתונה |  |
| 7:19.55 | מכולם | , |
| 7:19.56 | וקראו |  |
| 7:19.57 | לו |  |
| 7:19.58 | בשמו |  |
| 7:19.59 | של |  |
| 7:19.60 | הקב"ה |  |
| 7:19.61 | ואמרו |  |
| 7:19.62 | שהוא |  |
| 7:19.63 | ה | ' |
| 7:19.64 | אלהינו |  |
| 7:19.65 | ואנחנו |  |
| 7:19.66 | עמו |  |
| 7:19.67 | וכו | ' |
| 7:19.68 | ואותו |  |
| 7:19.69 | נעבוד |  |
| 7:19.70 | וכינו |  |
| 7:19.71 | אותו |  |
| 7:19.72 | בשם |  |
| 7:19.73 | זעיר |  |
| 7:19.74 | אנפין | ", |
| 7:19.75 | שענינו |  |
| 7:19.76 | בלשון |  |
| 7:19.77 | הכתוב |  |
| 7:19.78 | קצר |  |
| 7:19.79 | אפים | . |
| 7:19.80 | והנה |  |
| 7:19.81 | כבר |  |
| 7:19.82 | ידעת |  |
| 7:19.83 | ונתבאר |  |
| 7:19.84 | לך |  |
| 7:19.85 | מכתבי |  |
| 7:19.86 | הקדש |  |
| 7:19.87 | גנות |  |
| 7:19.88 | ופחיתות |  |
| 7:19.89 | המדה |  |
| 7:19.90 | המגונה |  |
| 7:19.91 | הזאת | , |
| 7:19.92 | ואיך |  |
| 7:19.93 | יתארו |  |
| 7:19.94 | את |  |
| 7:19.95 | האלהים |  |
| 7:19.96 | אשר |  |
| 7:19.97 | יעבדו |  |
| 7:19.98 | אותו |  |
| 7:19.99 | במדה |  |
| 7:19.100 | המגונה |  |
| 7:19.101 | הזאת | , |
| 7:19.102 | ועליהם |  |
| 7:19.103 | אמר |  |
| 7:19.104 | אליהו |  |
| 7:19.105 | בן |  |
| 7:19.106 | ברכאל |  |
| 7:19.107 | הבוזי |  |
| 7:19.108 | במליצתו | , |
| 7:19.109 | האמור |  |
| 7:19.110 | לממך |  |
| 7:19.111 | בליעל |  |
| 7:19.112 | רשע |  |
| 7:19.113 | אל |  |
| 7:19.114 | נדיבים | . |
| 7:19.115 | ופי | ' |
| 7:19.116 | רש"י |  |
| 7:19.117 | ז"ל |  |
| 7:19.118 | וכי |  |
| 7:19.119 | הגון |  |
| 7:19.120 | וראוי |  |
| 7:19.121 | לומר |  |
| 7:19.122 | למלך |  |
| 7:19.123 | בליעל | ? |
| 7:19.124 | ולבאר |  |
| 7:19.125 | רשע |  |
| 7:19.126 | אל |  |
| 7:19.127 | נדיבים | ? |
| 7:19.128 | אשר |  |
| 7:19.129 | לא |  |
| 7:19.130 | נשא |  |
| 7:19.131 | פני |  |
| 7:19.132 | שרים | , |
| 7:19.133 | כלומר |  |
| 7:19.134 | למלך |  |
| 7:19.135 | עולם |  |
| 7:19.136 | אשר |  |
| 7:19.137 | אין |  |
| 7:19.138 | משוא |  |
| 7:19.139 | פנים |  |
| 7:19.140 | לפניו |  |
| 7:19.141 | עכ"ל | . |
| 7:19.142 | וכן |  |
| 7:19.143 | פירש |  |
| 7:19.144 | רבי |  |
| 7:19.145 | אברהם |  |
| 7:19.146 | פריצול |  |
| 7:19.147 | ז"ל |  |
| 7:19.148 | האמור |  |
| 7:19.149 | למלך |  |
| 7:19.150 | וכו | ' |
| 7:19.151 | אשר |  |
| 7:19.152 | לא |  |
| 7:19.153 | נשא |  |
| 7:19.154 | פני |  |
| 7:19.155 | שרים | , |
| 7:19.156 | האם |  |
| 7:19.157 | ראוי |  |
| 7:19.158 | לומר |  |
| 7:19.159 | למלך |  |
| 7:19.160 | הממונה |  |
| 7:19.161 | על |  |
| 7:19.162 | הנהגה |  |
| 7:19.163 | כזאת |  |
| 7:19.164 | שיהיה |  |
| 7:19.165 | בליעל | ? |
| 7:19.166 | ואם |  |
| 7:19.167 | ראוי |  |
| 7:19.168 | לומר |  |
| 7:19.169 | רשע |  |
| 7:19.170 | למי |  |
| 7:19.171 | שהוא |  |
| 7:19.172 | נדיב |  |
| 7:19.173 | מנדיבי |  |
| 7:19.174 | עמים |  |
| 7:19.175 | ונודע |  |
| 7:19.176 | גזעו |  |
| 7:19.177 | וטובו | ? |
| 7:19.178 | ואיך |  |
| 7:19.179 | ראוי |  |
| 7:19.180 | ליחס |  |
| 7:19.181 | עול |  |
| 7:19.182 | למלך |  |
| 7:19.183 | צדיק |  |
| 7:19.184 | וישר |  |
| 7:19.185 | כזה |  |
| 7:19.186 | שהוא |  |
| 7:19.187 | ה | ' |
| 7:19.188 | אלהינו |  |
| 7:19.189 | אשר |  |
| 7:19.190 | לא |  |
| 7:19.191 | נשא |  |
| 7:19.192 | פני |  |
| 7:19.193 | שרים |  |
| 7:19.194 | בענין |  |
| 7:19.195 | הדין |  |
| 7:19.196 | לכבדם |  |
| 7:19.197 | ולשאת |  |
| 7:19.198 | את |  |
| 7:19.199 | פניהם | , |
| 7:19.200 | ולא |  |
| 7:19.201 | נכר |  |
| 7:19.202 | לפניו |  |
| 7:19.203 | השוע |  |
| 7:19.204 | והעשיר |  |
| 7:19.205 | בהיותו |  |
| 7:19.206 | טוען |  |
| 7:19.207 | בדין |  |
| 7:19.208 | עם |  |
| 7:19.209 | העני |  |
| 7:19.210 | לחוס |  |
| 7:19.211 | על |  |
| 7:19.212 | כבודו | , |
| 7:19.213 | בעבור |  |
| 7:19.214 | כי |  |
| 7:19.215 | מעשה |  |
| 7:19.216 | ידיו |  |
| 7:19.217 | כולם | , |
| 7:19.218 | ואין |  |
| 7:19.219 | הבדל |  |
| 7:19.220 | ביניהם |  |
| 7:19.221 | מצד |  |
| 7:19.222 | יצירתם |  |
| 7:19.223 | לפני |  |
| 7:19.224 | הקב"ה |  |
| 7:19.225 | עכ"ל | . |
| 7:19.226 | ודון |  |
| 7:19.227 | מינה |  |
| 7:19.228 | ואוקי |  |
| 7:19.229 | הכא | , |
| 7:19.230 | כיון |  |
| 7:19.231 | שהם |  |
| 7:19.232 | אומרים |  |
| 7:19.233 | שזעיר |  |
| 7:19.234 | אנפין |  |
| 7:19.235 | הוא |  |
| 7:19.236 | ה | ' |
| 7:19.237 | אלהינו |  |
| 7:19.238 | איך |  |
| 7:19.239 | יקראו |  |
| 7:19.240 | לו |  |
| 7:19.241 | בשם |  |
| 7:19.242 | קצר |  |
| 7:19.243 | אפים |  |
| 7:19.244 | אשר |  |
| 7:19.245 | יעשה |  |
| 7:19.246 | אולת | ? |
| 7:19.247 | כי |  |
| 7:19.248 | לא |  |
| 7:19.249 | יאות |  |
| 7:19.250 | לתאר |  |
| 7:19.251 | את |  |
| 7:19.252 | מלכו |  |
| 7:19.253 | של |  |
| 7:19.254 | עולם |  |
| 7:19.255 | אלא |  |
| 7:19.256 | בתארי |  |
| 7:19.257 | שבח |  |
| 7:19.258 | ועילוי | , |
| 7:19.259 | לא |  |
| 7:19.260 | תארי |  |
| 7:19.261 | אולת |  |
| 7:19.262 | וכסילות |  |
| 7:19.263 | כתאר |  |
| 7:19.264 | קצר |  |
| 7:19.265 | אפים |  |
| 7:19.266 | זעיר |  |
| 7:19.267 | אנפין | ) |
| 7:19.268 | שהוא |  |
| 7:19.269 | עושה |  |
| 7:19.270 | אולת | , |
| 7:19.271 | כמו |  |
| 7:19.272 | שאי |  |
| 7:19.273 | אפשר |  |
| 7:19.274 | לומר |  |
| 7:19.275 | למלכו |  |
| 7:19.276 | של |  |
| 7:19.277 | עולם |  |
| 7:19.278 | בליעל |  |
| 7:19.279 | כמאמר |  |
| 7:19.280 | אליהו | . |
| 7:19.281 | ומזה |  |
| 7:19.282 | תבין |  |
| 7:19.283 | ותשכיל |  |
| 7:19.284 | כי |  |
| 7:19.285 | לא |  |
| 7:19.286 | קבלה |  |
| 7:19.287 | אמתית |  |
| 7:19.288 | היא | . |
| 7:19.289 | רק |  |
| 7:19.290 | נכריה |  |
| 7:19.291 | היא | . |
| 7:19.292 | ולא |  |
| 7:19.293 | מרז"ל |  |
| 7:19.294 | בעלי |  |
| 7:19.295 | המשנה |  |
| 7:19.296 | והתלמוד |  |
| 7:19.297 | יצאו |  |
| 7:19.298 | הדברים |  |
| 7:19.299 | האלה | , |
| 7:19.300 | חלילה |  |
| 7:19.301 | לכל |  |
| 7:19.302 | בעל |  |
| 7:19.303 | אמונת |  |
| 7:19.304 | תורת |  |
| 7:19.305 | משה |  |
| 7:19.306 | רע"ה |  |
| 7:19.307 | וקבלת |  |
| 7:19.308 | רבותינו |  |
| 7:19.309 | ז"ל |  |
| 7:19.310 | להאמין |  |
| 7:19.311 | בדברים |  |
| 7:19.312 | כאלה | . |
| 7:19.313 | כי |  |
| 7:19.314 | כבר |  |
| 7:19.315 | נתברר |  |
| 7:19.316 | אצלינו |  |
| 7:19.317 | גדולת |  |
| 7:19.318 | חכמינו |  |
| 7:19.319 | וטוב |  |
| 7:19.320 | שכלם |  |
| 7:19.321 | ועומק |  |
| 7:19.322 | עיונם |  |
| 7:19.323 | בחכמה | , |
| 7:19.324 | והם |  |
| 7:19.325 | השיגו |  |
| 7:19.326 | את |  |
| 7:19.327 | האמת |  |
| 7:19.328 | השגה |  |
| 7:19.329 | גמורה |  |
| 7:19.330 | ואמתית | , |
| 7:19.331 | ואמרו |  |
| 7:19.332 | כל |  |
| 7:19.333 | המשתף |  |
| 7:19.334 | שם |  |
| 7:19.335 | שמים |  |
| 7:19.336 | ודבר |  |
| 7:19.337 | אחר |  |
| 7:19.338 | נעקר |  |
| 7:19.339 | מן |  |
| 7:19.340 | העולם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

