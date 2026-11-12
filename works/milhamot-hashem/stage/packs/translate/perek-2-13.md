# Translation pack: milhamot-hashem perek-2-13 (Milchamot ha-Shem 2:29–2:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-13`.

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
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **ריקנטי**: ר"מ ריקנטי → "Rabbi Menachem Rikanti" (the surname fixes the
  initial).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- עלת **העלות** → "the Cause of causes".
- אל **אחר** → "another god".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **עלה / עלול** → "cause / effect".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- גוף **דק** → "a subtle body".
- **כוונה** (in prayer) → "intention".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **א"כ** → "if so".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **לשבר** את האזן → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **אורה** זו תורה (Esth 8:16) → "'orah' means Torah".
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
- **תא חזי** → "come and see".
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
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


## Milchamot ha-Shem 2:29

Hebrew (Wikisource): והארכתי כל זה לישב לך דברי כס"א, אבל קבלתי כבר זכרתיה למעלה. ויש הוכחה מכאן, שמה שאתה שואלני אינו מטעם בקשתך האמת כמו שאמרת, כי אם לבקש איזה עלילה. ומה אני שאתה מחזיק אותי למקובל, כי אני בער ולא אדע. וכל תפלותי אפי' כצפצוף העופות אינם, כי העופות ניזונים שלא בצער ואין עליהם לא עול מלכות ולא עול דרך ארץ, ואפשר צפצוף בכוונה אבל אני כמה עול הזמן וטרדת המזונות מבלבלים דעתי ומחשבותי. ועוד שלא קבלתי מחכם מקובל מפה אל פה. ומי יתן ויניח לי ה' מטרדות הזמן והמזונות ויזמין לי רב מקובל יניחני במעגלי יושר, ואזכה להיות מבני עליה. ומה שהשבתי לך פעם שניה הוא דברי כס"א שהעלמת עין מהם, ומה שכתבת שאינך מגשים בספירות, והלא כתבת א"כ יהיה כח בגוף. ובפעם הזאת אתה אומר כי האור גוף דק. וכבר כתבתי לך כי מה שכינו אותן בשם אור אינו אלא כדי לשבר את האזן פעם תגשים ופעם תתנצל ולמה עולים בדעתך כל אלו הספקות אלא מפני שאתה מגשים. ומה שכתבת מי הוא המנהיג כל אלו העולמות הנעלמים, הממציאם הוא מחיה אותם. ומה שכתבת שנמצא הז"א אלהי התחתונים. הלא כתבתי לך שאינו אלא כלי לפעולות הבורא. ומה שכתבת שמקצתם אמרו את כל מעשה ה' הגדול מכלל דאיכא קטן? האומרים כן כבר ביארו הם שכוונתם על מטטרון ששמו כשם רבו, כמו שאמרו בתלמוד ואל משה אמר עלה אל ה', אלי מבעי ליה אלא זה מטטרון ששמו כשם רבו וכ"ה בילקוט. ומה שכתבת שרבותינו האמתיים אמרו, הובא במנורת המאור, שבשעת מתן תורה קרע הקב"ה ז' רקיעים וכו'. ובמדרש רבה אמרו אנכי ה' אלהיך, אני ראשון ואני אחרון. אני ראשון שאין לי אב ואני אחרון שאין לי בן ומבלעדי אין אלהים שאין לי אח. מי הוא האומר שיש לו אב או בם או אח? אלא אתה מוציא דבה על החכמים שהם אומרים כן, והם כבר ביארו דבריהם ז"ל, שכל הכנויים והתארים הם בספירות לשבר את האזן. לך נא וראה מה שכתב מהר"י גיקטליא בהקדמת שערי אורה, והאלהי ר"מ ריקנטי, ור"י חיים וז"ל, סוף דבר כל מקום שתראה דברים שאין ראוי לאמרם בבורא ית' כגון שיעור קומה והדומים להם הכל נאמר על הספירות וכשתראה שידברו דרך שבח והודיה, הכל נאמר על הבורא ית' שבתוכן ובחוצה להן כי אין דבר שיגבילהו ועל כן אין לומר בבורא ית' לא ימין ולא שמאל לא פנים ולא אחור וכו' עי"ש. ומה שכתבת שדעת המקובלים שהספירות משתלשלים בטבע כבן מן האב, חלילה שהם אמרו כן, אדרבה אמרו חלילה להאמין ולהעלות על לב שהספירות הם חלק מהאין סוף שיצאו ממנו ונשתלשל מעילה לעלול כי הוא עון פלילי. שהאין סוף אינו מתחלק לחלקים ח"ו, ואינו מקבל תוספת ולא מגרעת אלא מציאותו קיים תמיד בלי שום שינוי כלל, והוא חידשם חידוש גמור.

| id | word | punctuation after |
|---|---|---|
| 2:29.1 | והארכתי |  |
| 2:29.2 | כל |  |
| 2:29.3 | זה |  |
| 2:29.4 | לישב |  |
| 2:29.5 | לך |  |
| 2:29.6 | דברי |  |
| 2:29.7 | כס"א | , |
| 2:29.8 | אבל |  |
| 2:29.9 | קבלתי |  |
| 2:29.10 | כבר |  |
| 2:29.11 | זכרתיה |  |
| 2:29.12 | למעלה | . |
| 2:29.13 | ויש |  |
| 2:29.14 | הוכחה |  |
| 2:29.15 | מכאן | , |
| 2:29.16 | שמה |  |
| 2:29.17 | שאתה |  |
| 2:29.18 | שואלני |  |
| 2:29.19 | אינו |  |
| 2:29.20 | מטעם |  |
| 2:29.21 | בקשתך |  |
| 2:29.22 | האמת |  |
| 2:29.23 | כמו |  |
| 2:29.24 | שאמרת | , |
| 2:29.25 | כי |  |
| 2:29.26 | אם |  |
| 2:29.27 | לבקש |  |
| 2:29.28 | איזה |  |
| 2:29.29 | עלילה | . |
| 2:29.30 | ומה |  |
| 2:29.31 | אני |  |
| 2:29.32 | שאתה |  |
| 2:29.33 | מחזיק |  |
| 2:29.34 | אותי |  |
| 2:29.35 | למקובל | , |
| 2:29.36 | כי |  |
| 2:29.37 | אני |  |
| 2:29.38 | בער |  |
| 2:29.39 | ולא |  |
| 2:29.40 | אדע | . |
| 2:29.41 | וכל |  |
| 2:29.42 | תפלותי |  |
| 2:29.43 | אפי | ' |
| 2:29.44 | כצפצוף |  |
| 2:29.45 | העופות |  |
| 2:29.46 | אינם | , |
| 2:29.47 | כי |  |
| 2:29.48 | העופות |  |
| 2:29.49 | ניזונים |  |
| 2:29.50 | שלא |  |
| 2:29.51 | בצער |  |
| 2:29.52 | ואין |  |
| 2:29.53 | עליהם |  |
| 2:29.54 | לא |  |
| 2:29.55 | עול |  |
| 2:29.56 | מלכות |  |
| 2:29.57 | ולא |  |
| 2:29.58 | עול |  |
| 2:29.59 | דרך |  |
| 2:29.60 | ארץ | , |
| 2:29.61 | ואפשר |  |
| 2:29.62 | צפצוף |  |
| 2:29.63 | בכוונה |  |
| 2:29.64 | אבל |  |
| 2:29.65 | אני |  |
| 2:29.66 | כמה |  |
| 2:29.67 | עול |  |
| 2:29.68 | הזמן |  |
| 2:29.69 | וטרדת |  |
| 2:29.70 | המזונות |  |
| 2:29.71 | מבלבלים |  |
| 2:29.72 | דעתי |  |
| 2:29.73 | ומחשבותי | . |
| 2:29.74 | ועוד |  |
| 2:29.75 | שלא |  |
| 2:29.76 | קבלתי |  |
| 2:29.77 | מחכם |  |
| 2:29.78 | מקובל |  |
| 2:29.79 | מפה |  |
| 2:29.80 | אל |  |
| 2:29.81 | פה | . |
| 2:29.82 | ומי |  |
| 2:29.83 | יתן |  |
| 2:29.84 | ויניח |  |
| 2:29.85 | לי |  |
| 2:29.86 | ה | ' |
| 2:29.87 | מטרדות |  |
| 2:29.88 | הזמן |  |
| 2:29.89 | והמזונות |  |
| 2:29.90 | ויזמין |  |
| 2:29.91 | לי |  |
| 2:29.92 | רב |  |
| 2:29.93 | מקובל |  |
| 2:29.94 | יניחני |  |
| 2:29.95 | במעגלי |  |
| 2:29.96 | יושר | , |
| 2:29.97 | ואזכה |  |
| 2:29.98 | להיות |  |
| 2:29.99 | מבני |  |
| 2:29.100 | עליה | . |
| 2:29.101 | ומה |  |
| 2:29.102 | שהשבתי |  |
| 2:29.103 | לך |  |
| 2:29.104 | פעם |  |
| 2:29.105 | שניה |  |
| 2:29.106 | הוא |  |
| 2:29.107 | דברי |  |
| 2:29.108 | כס"א |  |
| 2:29.109 | שהעלמת |  |
| 2:29.110 | עין |  |
| 2:29.111 | מהם | , |
| 2:29.112 | ומה |  |
| 2:29.113 | שכתבת |  |
| 2:29.114 | שאינך |  |
| 2:29.115 | מגשים |  |
| 2:29.116 | בספירות | , |
| 2:29.117 | והלא |  |
| 2:29.118 | כתבת |  |
| 2:29.119 | א"כ |  |
| 2:29.120 | יהיה |  |
| 2:29.121 | כח |  |
| 2:29.122 | בגוף | . |
| 2:29.123 | ובפעם |  |
| 2:29.124 | הזאת |  |
| 2:29.125 | אתה |  |
| 2:29.126 | אומר |  |
| 2:29.127 | כי |  |
| 2:29.128 | האור |  |
| 2:29.129 | גוף |  |
| 2:29.130 | דק | . |
| 2:29.131 | וכבר |  |
| 2:29.132 | כתבתי |  |
| 2:29.133 | לך |  |
| 2:29.134 | כי |  |
| 2:29.135 | מה |  |
| 2:29.136 | שכינו |  |
| 2:29.137 | אותן |  |
| 2:29.138 | בשם |  |
| 2:29.139 | אור |  |
| 2:29.140 | אינו |  |
| 2:29.141 | אלא |  |
| 2:29.142 | כדי |  |
| 2:29.143 | לשבר |  |
| 2:29.144 | את |  |
| 2:29.145 | האזן |  |
| 2:29.146 | פעם |  |
| 2:29.147 | תגשים |  |
| 2:29.148 | ופעם |  |
| 2:29.149 | תתנצל |  |
| 2:29.150 | ולמה |  |
| 2:29.151 | עולים |  |
| 2:29.152 | בדעתך |  |
| 2:29.153 | כל |  |
| 2:29.154 | אלו |  |
| 2:29.155 | הספקות |  |
| 2:29.156 | אלא |  |
| 2:29.157 | מפני |  |
| 2:29.158 | שאתה |  |
| 2:29.159 | מגשים | . |
| 2:29.160 | ומה |  |
| 2:29.161 | שכתבת |  |
| 2:29.162 | מי |  |
| 2:29.163 | הוא |  |
| 2:29.164 | המנהיג |  |
| 2:29.165 | כל |  |
| 2:29.166 | אלו |  |
| 2:29.167 | העולמות |  |
| 2:29.168 | הנעלמים | , |
| 2:29.169 | הממציאם |  |
| 2:29.170 | הוא |  |
| 2:29.171 | מחיה |  |
| 2:29.172 | אותם | . |
| 2:29.173 | ומה |  |
| 2:29.174 | שכתבת |  |
| 2:29.175 | שנמצא |  |
| 2:29.176 | הז"א |  |
| 2:29.177 | אלהי |  |
| 2:29.178 | התחתונים | . |
| 2:29.179 | הלא |  |
| 2:29.180 | כתבתי |  |
| 2:29.181 | לך |  |
| 2:29.182 | שאינו |  |
| 2:29.183 | אלא |  |
| 2:29.184 | כלי |  |
| 2:29.185 | לפעולות |  |
| 2:29.186 | הבורא | . |
| 2:29.187 | ומה |  |
| 2:29.188 | שכתבת |  |
| 2:29.189 | שמקצתם |  |
| 2:29.190 | אמרו |  |
| 2:29.191 | את |  |
| 2:29.192 | כל |  |
| 2:29.193 | מעשה |  |
| 2:29.194 | ה | ' |
| 2:29.195 | הגדול |  |
| 2:29.196 | מכלל |  |
| 2:29.197 | דאיכא |  |
| 2:29.198 | קטן | ? |
| 2:29.199 | האומרים |  |
| 2:29.200 | כן |  |
| 2:29.201 | כבר |  |
| 2:29.202 | ביארו |  |
| 2:29.203 | הם |  |
| 2:29.204 | שכוונתם |  |
| 2:29.205 | על |  |
| 2:29.206 | מטטרון |  |
| 2:29.207 | ששמו |  |
| 2:29.208 | כשם |  |
| 2:29.209 | רבו | , |
| 2:29.210 | כמו |  |
| 2:29.211 | שאמרו |  |
| 2:29.212 | בתלמוד |  |
| 2:29.213 | ואל |  |
| 2:29.214 | משה |  |
| 2:29.215 | אמר |  |
| 2:29.216 | עלה |  |
| 2:29.217 | אל |  |
| 2:29.218 | ה | ', |
| 2:29.219 | אלי |  |
| 2:29.220 | מבעי |  |
| 2:29.221 | ליה |  |
| 2:29.222 | אלא |  |
| 2:29.223 | זה |  |
| 2:29.224 | מטטרון |  |
| 2:29.225 | ששמו |  |
| 2:29.226 | כשם |  |
| 2:29.227 | רבו |  |
| 2:29.228 | וכ"ה |  |
| 2:29.229 | בילקוט | . |
| 2:29.230 | ומה |  |
| 2:29.231 | שכתבת |  |
| 2:29.232 | שרבותינו |  |
| 2:29.233 | האמתיים |  |
| 2:29.234 | אמרו | , |
| 2:29.235 | הובא |  |
| 2:29.236 | במנורת |  |
| 2:29.237 | המאור | , |
| 2:29.238 | שבשעת |  |
| 2:29.239 | מתן |  |
| 2:29.240 | תורה |  |
| 2:29.241 | קרע |  |
| 2:29.242 | הקב"ה |  |
| 2:29.243 | ז | ' |
| 2:29.244 | רקיעים |  |
| 2:29.245 | וכו | '. |
| 2:29.246 | ובמדרש |  |
| 2:29.247 | רבה |  |
| 2:29.248 | אמרו |  |
| 2:29.249 | אנכי |  |
| 2:29.250 | ה | ' |
| 2:29.251 | אלהיך | , |
| 2:29.252 | אני |  |
| 2:29.253 | ראשון |  |
| 2:29.254 | ואני |  |
| 2:29.255 | אחרון | . |
| 2:29.256 | אני |  |
| 2:29.257 | ראשון |  |
| 2:29.258 | שאין |  |
| 2:29.259 | לי |  |
| 2:29.260 | אב |  |
| 2:29.261 | ואני |  |
| 2:29.262 | אחרון |  |
| 2:29.263 | שאין |  |
| 2:29.264 | לי |  |
| 2:29.265 | בן |  |
| 2:29.266 | ומבלעדי |  |
| 2:29.267 | אין |  |
| 2:29.268 | אלהים |  |
| 2:29.269 | שאין |  |
| 2:29.270 | לי |  |
| 2:29.271 | אח | . |
| 2:29.272 | מי |  |
| 2:29.273 | הוא |  |
| 2:29.274 | האומר |  |
| 2:29.275 | שיש |  |
| 2:29.276 | לו |  |
| 2:29.277 | אב |  |
| 2:29.278 | או |  |
| 2:29.279 | בם |  |
| 2:29.280 | או |  |
| 2:29.281 | אח | ? |
| 2:29.282 | אלא |  |
| 2:29.283 | אתה |  |
| 2:29.284 | מוציא |  |
| 2:29.285 | דבה |  |
| 2:29.286 | על |  |
| 2:29.287 | החכמים |  |
| 2:29.288 | שהם |  |
| 2:29.289 | אומרים |  |
| 2:29.290 | כן | , |
| 2:29.291 | והם |  |
| 2:29.292 | כבר |  |
| 2:29.293 | ביארו |  |
| 2:29.294 | דבריהם |  |
| 2:29.295 | ז"ל | , |
| 2:29.296 | שכל |  |
| 2:29.297 | הכנויים |  |
| 2:29.298 | והתארים |  |
| 2:29.299 | הם |  |
| 2:29.300 | בספירות |  |
| 2:29.301 | לשבר |  |
| 2:29.302 | את |  |
| 2:29.303 | האזן | . |
| 2:29.304 | לך |  |
| 2:29.305 | נא |  |
| 2:29.306 | וראה |  |
| 2:29.307 | מה |  |
| 2:29.308 | שכתב |  |
| 2:29.309 | מהר"י |  |
| 2:29.310 | גיקטליא |  |
| 2:29.311 | בהקדמת |  |
| 2:29.312 | שערי |  |
| 2:29.313 | אורה | , |
| 2:29.314 | והאלהי |  |
| 2:29.315 | ר"מ |  |
| 2:29.316 | ריקנטי | , |
| 2:29.317 | ור"י |  |
| 2:29.318 | חיים |  |
| 2:29.319 | וז"ל | , |
| 2:29.320 | סוף |  |
| 2:29.321 | דבר |  |
| 2:29.322 | כל |  |
| 2:29.323 | מקום |  |
| 2:29.324 | שתראה |  |
| 2:29.325 | דברים |  |
| 2:29.326 | שאין |  |
| 2:29.327 | ראוי |  |
| 2:29.328 | לאמרם |  |
| 2:29.329 | בבורא |  |
| 2:29.330 | ית | ' |
| 2:29.331 | כגון |  |
| 2:29.332 | שיעור |  |
| 2:29.333 | קומה |  |
| 2:29.334 | והדומים |  |
| 2:29.335 | להם |  |
| 2:29.336 | הכל |  |
| 2:29.337 | נאמר |  |
| 2:29.338 | על |  |
| 2:29.339 | הספירות |  |
| 2:29.340 | וכשתראה |  |
| 2:29.341 | שידברו |  |
| 2:29.342 | דרך |  |
| 2:29.343 | שבח |  |
| 2:29.344 | והודיה | , |
| 2:29.345 | הכל |  |
| 2:29.346 | נאמר |  |
| 2:29.347 | על |  |
| 2:29.348 | הבורא |  |
| 2:29.349 | ית | ' |
| 2:29.350 | שבתוכן |  |
| 2:29.351 | ובחוצה |  |
| 2:29.352 | להן |  |
| 2:29.353 | כי |  |
| 2:29.354 | אין |  |
| 2:29.355 | דבר |  |
| 2:29.356 | שיגבילהו |  |
| 2:29.357 | ועל |  |
| 2:29.358 | כן |  |
| 2:29.359 | אין |  |
| 2:29.360 | לומר |  |
| 2:29.361 | בבורא |  |
| 2:29.362 | ית | ' |
| 2:29.363 | לא |  |
| 2:29.364 | ימין |  |
| 2:29.365 | ולא |  |
| 2:29.366 | שמאל |  |
| 2:29.367 | לא |  |
| 2:29.368 | פנים |  |
| 2:29.369 | ולא |  |
| 2:29.370 | אחור |  |
| 2:29.371 | וכו | ' |
| 2:29.372 | עי"ש | . |
| 2:29.373 | ומה |  |
| 2:29.374 | שכתבת |  |
| 2:29.375 | שדעת |  |
| 2:29.376 | המקובלים |  |
| 2:29.377 | שהספירות |  |
| 2:29.378 | משתלשלים |  |
| 2:29.379 | בטבע |  |
| 2:29.380 | כבן |  |
| 2:29.381 | מן |  |
| 2:29.382 | האב | , |
| 2:29.383 | חלילה |  |
| 2:29.384 | שהם |  |
| 2:29.385 | אמרו |  |
| 2:29.386 | כן | , |
| 2:29.387 | אדרבה |  |
| 2:29.388 | אמרו |  |
| 2:29.389 | חלילה |  |
| 2:29.390 | להאמין |  |
| 2:29.391 | ולהעלות |  |
| 2:29.392 | על |  |
| 2:29.393 | לב |  |
| 2:29.394 | שהספירות |  |
| 2:29.395 | הם |  |
| 2:29.396 | חלק |  |
| 2:29.397 | מהאין |  |
| 2:29.398 | סוף |  |
| 2:29.399 | שיצאו |  |
| 2:29.400 | ממנו |  |
| 2:29.401 | ונשתלשל |  |
| 2:29.402 | מעילה |  |
| 2:29.403 | לעלול |  |
| 2:29.404 | כי |  |
| 2:29.405 | הוא |  |
| 2:29.406 | עון |  |
| 2:29.407 | פלילי | . |
| 2:29.408 | שהאין |  |
| 2:29.409 | סוף |  |
| 2:29.410 | אינו |  |
| 2:29.411 | מתחלק |  |
| 2:29.412 | לחלקים |  |
| 2:29.413 | ח"ו | , |
| 2:29.414 | ואינו |  |
| 2:29.415 | מקבל |  |
| 2:29.416 | תוספת |  |
| 2:29.417 | ולא |  |
| 2:29.418 | מגרעת |  |
| 2:29.419 | אלא |  |
| 2:29.420 | מציאותו |  |
| 2:29.421 | קיים |  |
| 2:29.422 | תמיד |  |
| 2:29.423 | בלי |  |
| 2:29.424 | שום |  |
| 2:29.425 | שינוי |  |
| 2:29.426 | כלל | , |
| 2:29.427 | והוא |  |
| 2:29.428 | חידשם |  |
| 2:29.429 | חידוש |  |
| 2:29.430 | גמור | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

