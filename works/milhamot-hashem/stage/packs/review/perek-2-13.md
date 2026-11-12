# Review pack: milhamot-hashem perek-2-13 (Milchamot ha-Shem 2:29–2:29)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-13`.

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


### Translation under review (2:29)

English: I have gone on at length about all this, to set out for you the words of Kisse Eliyahu, but I have already mentioned it above. And there is proof from this that what you ask me is not from the motive of your request for the truth, as you said, but rather to seek some pretext. And who am I that you should hold me to be a Kabbalist, for I am brutish and do not know. And all my foolish talk is not even like the chirping of birds, for the birds are fed without distress, and on them there lies neither the yoke of kingship nor the yoke of the way of the land, and chirping can be done with intent, but I, under so much burden of the times and the cares of livelihood, have my mind and my thoughts confused. And furthermore, I have not received from a wise Kabbalist mouth to mouth. Would that the LORD relieve me of the troubles of time and of livelihood, and provide me a great Kabbalist who would set me on the paths of uprightness, and that I might merit to be among the sons of the ascent. As for what I answered you the second time, it is the words of Kisse Eliyahu that you turned a blind eye to. And as for what you wrote, that you do not ascribe corporeality to the sefirot, But did you not write: "If so, there would be a force in a body"? And this time you say that light is a subtle body. And I already wrote to you that what they called them by the name light is only in order to shock the ear, so that you ascribe corporeality one time and make excuses another. And why do all these doubts come into your mind, except because you ascribe corporeality? As for what you wrote: "Who is it that governs all these hidden worlds?" The One who brings them into being is the one who keeps them alive. And as for what you wrote, that Ze'ir Anpin is found to be the god of the lower ones. Have I not written to you that he is nothing but a vessel for the acts of the Creator? And as for what you wrote, that some of them said "all the great work of the LORD," from which it follows that there is a small one? Those who say so have already explained that they meant Metatron, whose name is like his master's, as they said in the Talmud: "And to Moses He said, 'Come up to the LORD'; it should have said 'to Me'. But this is Metatron, whose name is like his master's." and so it is in the Yalkut. And as for what you wrote, that our true rabbis said, it is cited in Menorat ha-Ma'or that at the time of the giving of the Torah the Holy One, blessed be He, tore 7 heavens, etc. And in Midrash Rabbah they said: "I am the LORD your God; I am the first and I am the last. I am the first, for I have no father, and I am the last, for I have no son; and besides Me there is no God, for I have no brother." Who is it that says he has a father, or a son, or a brother? but you spread slander against the sages who say so, and they have already explained their words, of blessed memory, that all the epithets and the attributes are in the sefirot, to shock the ear. Go now and see what Mahari Gikatilia wrote in the introduction to Sha'arei Orah, and the divine R. M. Ricanti, and R. Y. Chayyim, and these are his words: "In sum, wherever you see things that are not fitting to say of the Creator, may He be blessed, such as measure of stature and the like, it is all said of the sefirot; and when you see them speak in the manner of praise and thanksgiving, it is all said of the Creator, may He be blessed, who is within them and outside them, for nothing limits Him; therefore one may not say of the Creator, may He be blessed, not right and not left, not face and not back, etc." See there. And as for what you wrote, that the opinion of the Kabbalists is that the sefirot descend in a chain, by nature like a son from the father, Far be it that they said so. On the contrary, they said: "Far be it to believe or to bring to mind that the sefirot are part of Ein Sof that came out of it and descended from cause to effect, for it is a criminal sin. For Ein Sof is not divided into parts, God forbid, and it receives no addition and no diminution, but its existence is constant always, without any change whatever, and He originated them, an absolute origination."

Literal: I have gone on at length about all this to set out for you the words of Kisse Eliyahu but I have taken up already mentioned it above and there is proof from this that what you ask me is not from the motive of your request for the truth as you said but rather to seek some pretext and what am I that you hold me as a Kabbalist for I am brutish and do not know and all my foolish talk even like the chirping of birds is not for the birds are fed without distress and there is not on them neither the yoke of kingship nor the yoke of the way of the land and it is possible chirping with intent but I so much burden of the times and the worry of livelihood confuse my mind and my thoughts and furthermore that not I have received from a wise man Kabbalist from mouth to mouth and would that will give and relieve me the LORD of the troubles of time and of livelihood and provide me a great Kabbalist who would set me on the paths of uprightness and I will merit to be from the sons of ascent as for what I answered to you time second is the words of Kisse Eliyahu that you turned a blind eye to them and as for what you wrote that you are not ascribing corporeality to the sefirot but did not you write if so there would be a force in a body and at this time this you say that light a body subtle and already I wrote to you that what they called them by the name light is not but in order to shock (object marker) the ear at one time you ascribe corporeality and another time you make excuses and why come up in your mind all these doubts except because you ascribe corporeality as for what you wrote who is it that governs all these worlds hidden the one who brings into being is keeps alive them and as for what you wrote that is found Ze'ir Anpin god of the lower ones have I not written to you that he is not but a vessel for the acts of the Creator and as for what you wrote that some of them said (object marker) all the work of the LORD great from which it follows that there is a small one those who say so already have explained they their intent was on Metatron whose name is like the name of his master as they said in the Talmud and to Moses he said come up to the LORD to Me it should have said it but this is Metatron whose name is like his master's name and so it is in the Yalkut and as for what you wrote that our rabbis the true ones said it is cited in Menorat ha-Ma'or that at the time of the giving of the Torah tore the Holy One, blessed be He 7 heavens etc. and in Midrash Rabbah they said I am the LORD your God I the first and I last I am first for I have no father and I am last for I have no son and besides Me there is no God for I have no brother who is it that says he has to him a father or a son or a brother but you spread slander against the sages who say so and they already have explained their words of blessed memory that all the epithets and the attributes are in the sefirot to shock (object marker) the ear go now and see what wrote Mahari Gikatilia in the introduction to Sha'arei Orah and the divine R. M. Ricanti and R. Y. Chayyim and these are his words in sum the end of the matter all wherever you see things that are not fitting to say of the Creator may He be blessed such as measure of stature and the like of them all is said of the sefirot and when you see them speak in the manner of praise and thanksgiving all is said of the Creator may He be blessed who is within and outside them for there is not anything that limits Him therefore so one may not say of the Creator may He be blessed not right and not left not face and not back etc. see there and as for what you wrote that the opinion of the Kabbalists that the sefirot descend in a chain by nature like a son from the father far be it that they said so on the contrary they said far be it to believe or to bring to mind the heart that the sefirot are part of Ein Sof that came out from it and descended from cause to effect for it is a sin criminal that Ein Sof is not divided into parts God forbid and it does not receive addition and not diminution but its existence is constant always without any change whatever and He originated them origination absolute

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:29.1 | I have gone on at length |  |  |  |
| 2:29.2 | about all |  |  |  |
| 2:29.3 | this |  |  |  |
| 2:29.4 | to set out |  |  |  |
| 2:29.5 | for you |  |  |  |
| 2:29.6 | the words of |  |  |  |
| 2:29.7 | Kisse Eliyahu |  |  | Kisse Eliyahu: 'the Throne of Elijah', the literal sense of the title; the abbreviation כס"א is written out. |
| 2:29.8 | but |  |  |  |
| 2:29.9 | I have taken up |  |  | The root of 'Kabbalah' (received tradition). Here 'I have taken up' the matter, that is, dealt with it. |
| 2:29.10 | already |  |  |  |
| 2:29.11 | mentioned it |  |  |  |
| 2:29.12 | above |  |  |  |
| 2:29.13 | and there is |  |  |  |
| 2:29.14 | proof |  |  |  |
| 2:29.15 | from this |  |  |  |
| 2:29.16 | that what |  |  |  |
| 2:29.17 | you |  |  |  |
| 2:29.18 | ask me |  |  |  |
| 2:29.19 | is not |  |  |  |
| 2:29.20 | from the motive of |  |  |  |
| 2:29.21 | your request |  |  |  |
| 2:29.22 | for the truth |  |  |  |
| 2:29.23 | as |  |  |  |
| 2:29.24 | you said |  |  |  |
| 2:29.25 | but |  |  |  |
| 2:29.26 | rather |  |  |  |
| 2:29.27 | to seek |  |  |  |
| 2:29.28 | some |  |  |  |
| 2:29.29 | pretext |  |  |  |
| 2:29.30 | and what |  |  |  |
| 2:29.31 | am I |  |  |  |
| 2:29.32 | that you |  |  |  |
| 2:29.33 | hold |  |  |  |
| 2:29.34 | me |  |  |  |
| 2:29.35 | as a Kabbalist |  |  |  |
| 2:29.36 | for |  |  |  |
| 2:29.37 | I |  |  |  |
| 2:29.38 | am brutish |  |  | Ps 73:22 ('I am brutish'); see apparatus. |
| 2:29.39 | and |  |  |  |
| 2:29.40 | do not know |  |  |  |
| 2:29.41 | and all |  |  |  |
| 2:29.42 | my foolish talk |  |  |  |
| 2:29.43 | even |  |  |  |
| 2:29.44 | like the chirping |  |  |  |
| 2:29.45 | of birds |  |  |  |
| 2:29.46 | is not |  |  |  |
| 2:29.47 | for |  |  |  |
| 2:29.48 | the birds |  |  |  |
| 2:29.49 | are fed |  |  |  |
| 2:29.50 | without |  |  |  |
| 2:29.51 | distress |  |  |  |
| 2:29.52 | and there is not |  |  |  |
| 2:29.53 | on them |  |  |  |
| 2:29.54 | neither |  |  |  |
| 2:29.55 | the yoke |  |  |  |
| 2:29.56 | of kingship |  |  |  |
| 2:29.57 | nor |  |  |  |
| 2:29.58 | the yoke |  |  |  |
| 2:29.59 | of the way |  |  | Literally 'the way of the land': the ordinary business of earning a living. |
| 2:29.60 | of the land |  |  |  |
| 2:29.61 | and it is possible |  |  |  |
| 2:29.62 | chirping |  |  |  |
| 2:29.63 | with intent |  |  |  |
| 2:29.64 | but |  |  |  |
| 2:29.65 | I |  |  |  |
| 2:29.66 | so much |  |  |  |
| 2:29.67 | burden |  |  |  |
| 2:29.68 | of the times |  |  |  |
| 2:29.69 | and the worry |  |  |  |
| 2:29.70 | of livelihood |  |  |  |
| 2:29.71 | confuse |  |  |  |
| 2:29.72 | my mind |  |  |  |
| 2:29.73 | and my thoughts |  |  |  |
| 2:29.74 | and furthermore |  |  |  |
| 2:29.75 | that not |  |  |  |
| 2:29.76 | I have received |  |  | The root of 'Kabbalah'. Here 'I have received': the writer has not had the tradition from a master. |
| 2:29.77 | from a wise man |  |  |  |
| 2:29.78 | Kabbalist |  |  |  |
| 2:29.79 | from mouth |  |  |  |
| 2:29.80 | to |  |  |  |
| 2:29.81 | mouth |  |  |  |
| 2:29.82 | and would that |  |  |  |
| 2:29.83 | will give |  |  |  |
| 2:29.84 | and relieve |  |  |  |
| 2:29.85 | me |  |  |  |
| 2:29.86 | the LORD |  |  |  |
| 2:29.87 | of the troubles |  |  |  |
| 2:29.88 | of time |  |  |  |
| 2:29.89 | and of livelihood |  |  |  |
| 2:29.90 | and provide |  |  |  |
| 2:29.91 | me |  |  |  |
| 2:29.92 | a great |  |  |  |
| 2:29.93 | Kabbalist |  |  | A Kabbalist: one who holds the received tradition. |
| 2:29.94 | who would set me |  |  |  |
| 2:29.95 | on the paths |  |  |  |
| 2:29.96 | of uprightness |  |  |  |
| 2:29.97 | and I will merit |  |  |  |
| 2:29.98 | to be |  |  |  |
| 2:29.99 | from the sons |  |  |  |
| 2:29.100 | of ascent |  |  | Literally 'sons of ascent': the idiom for those who rise to a higher stage of learning. |
| 2:29.101 | as for |  |  |  |
| 2:29.102 | what I answered |  |  |  |
| 2:29.103 | to you |  |  |  |
| 2:29.104 | time |  |  |  |
| 2:29.105 | second |  |  |  |
| 2:29.106 | is |  |  |  |
| 2:29.107 | the words of |  |  |  |
| 2:29.108 | Kisse Eliyahu |  |  |  |
| 2:29.109 | that you turned |  |  |  |
| 2:29.110 | a blind eye |  |  |  |
| 2:29.111 | to them |  |  |  |
| 2:29.112 | and as for |  |  |  |
| 2:29.113 | what you wrote |  |  |  |
| 2:29.114 | that you are not |  |  |  |
| 2:29.115 | ascribing corporeality |  |  |  |
| 2:29.116 | to the sefirot |  |  |  |
| 2:29.117 | but did not |  |  |  |
| 2:29.118 | you write |  |  |  |
| 2:29.119 | if so |  |  | 'if so' (אם כן). |
| 2:29.120 | there would be |  |  |  |
| 2:29.121 | a force |  |  |  |
| 2:29.122 | in a body |  |  |  |
| 2:29.123 | and at this time |  |  |  |
| 2:29.124 | this |  |  |  |
| 2:29.125 | you |  |  |  |
| 2:29.126 | say |  |  |  |
| 2:29.127 | that |  |  |  |
| 2:29.128 | light |  |  |  |
| 2:29.129 | a body |  |  |  |
| 2:29.130 | subtle |  |  |  |
| 2:29.131 | and already |  |  |  |
| 2:29.132 | I wrote |  |  |  |
| 2:29.133 | to you |  |  |  |
| 2:29.134 | that |  |  |  |
| 2:29.135 | what |  |  |  |
| 2:29.136 | they called |  |  |  |
| 2:29.137 | them |  |  |  |
| 2:29.138 | by the name |  |  |  |
| 2:29.139 | light |  |  | 'light': the name the Kabbalists give the sefirot; the writer argues that it is used as a name, not as a thing. |
| 2:29.140 | is not |  |  |  |
| 2:29.141 | but |  |  |  |
| 2:29.142 | in order |  |  | 'in order': the purpose of the name, as the writer gives it. |
| 2:29.143 | to shock |  |  |  |
| 2:29.144 | (object marker) |  |  |  |
| 2:29.145 | the ear |  |  |  |
| 2:29.146 | at one time |  |  | 'at one time' and 'another time' (פעם ... ופעם): the writer's point that the reader alternates between two errors. |
| 2:29.147 | you ascribe corporeality |  |  |  |
| 2:29.148 | and another time |  |  |  |
| 2:29.149 | you make excuses |  |  |  |
| 2:29.150 | and why |  |  |  |
| 2:29.151 | come up |  |  |  |
| 2:29.152 | in your mind |  |  |  |
| 2:29.153 | all |  |  |  |
| 2:29.154 | these |  |  |  |
| 2:29.155 | doubts |  |  |  |
| 2:29.156 | except |  |  |  |
| 2:29.157 | because |  |  |  |
| 2:29.158 | you |  |  |  |
| 2:29.159 | ascribe corporeality |  |  |  |
| 2:29.160 | as for |  |  |  |
| 2:29.161 | what you wrote |  |  |  |
| 2:29.162 | who |  |  |  |
| 2:29.163 | is it |  |  |  |
| 2:29.164 | that governs |  |  |  |
| 2:29.165 | all |  |  |  |
| 2:29.166 | these |  |  |  |
| 2:29.167 | worlds |  |  |  |
| 2:29.168 | hidden |  |  |  |
| 2:29.169 | the one who brings into being |  |  |  |
| 2:29.170 | is |  |  |  |
| 2:29.171 | keeps alive |  |  |  |
| 2:29.172 | them |  |  |  |
| 2:29.173 | and as for |  |  |  |
| 2:29.174 | what you wrote |  |  |  |
| 2:29.175 | that is found |  |  |  |
| 2:29.176 | Ze'ir Anpin |  |  | Ze'ir Anpin: Aramaic 'small face', the name of a partzuf (ז"א = Ze'ir Anpin). |
| 2:29.177 | god of |  |  | 'god of': the questioner's charge, in his words. The Hebrew is 'elohei' (god of the lower ones). |
| 2:29.178 | the lower ones |  |  |  |
| 2:29.179 | have I not |  |  |  |
| 2:29.180 | written |  |  |  |
| 2:29.181 | to you |  |  |  |
| 2:29.182 | that he is not |  |  |  |
| 2:29.183 | but |  |  |  |
| 2:29.184 | a vessel |  |  |  |
| 2:29.185 | for the acts |  |  |  |
| 2:29.186 | of the Creator |  |  |  |
| 2:29.187 | and as for |  |  |  |
| 2:29.188 | what you wrote |  |  |  |
| 2:29.189 | that some of them |  |  |  |
| 2:29.190 | said |  |  |  |
| 2:29.191 | (object marker) |  |  |  |
| 2:29.192 | all |  |  |  |
| 2:29.193 | the work of |  |  |  |
| 2:29.194 | the LORD |  |  | the LORD: the Tetragrammaton in the quoted verse, Deut 11:7. |
| 2:29.195 | great |  |  |  |
| 2:29.196 | from which it follows |  |  |  |
| 2:29.197 | that there is |  |  |  |
| 2:29.198 | a small one |  |  |  |
| 2:29.199 | those who say |  |  |  |
| 2:29.200 | so |  |  |  |
| 2:29.201 | already |  |  |  |
| 2:29.202 | have explained |  |  |  |
| 2:29.203 | they |  |  |  |
| 2:29.204 | their intent was |  |  |  |
| 2:29.205 | on |  |  |  |
| 2:29.206 | Metatron |  |  | Metatron: the name as the Talmud gives it. |
| 2:29.207 | whose name |  |  |  |
| 2:29.208 | is like the name of |  |  |  |
| 2:29.209 | his master |  |  |  |
| 2:29.210 | as |  |  |  |
| 2:29.211 | they said |  |  |  |
| 2:29.212 | in the Talmud |  |  |  |
| 2:29.213 | and to |  |  |  |
| 2:29.214 | Moses |  |  |  |
| 2:29.215 | he said |  |  |  |
| 2:29.216 | come up |  |  |  |
| 2:29.217 | to |  |  |  |
| 2:29.218 | the LORD |  |  | the LORD: in Exod 24:1. |
| 2:29.219 | to Me |  |  | 'to Me' (אלי): the Talmud's point that the verse should have said 'to Me' and does not. |
| 2:29.220 | it should have said |  |  |  |
| 2:29.221 | it |  |  |  |
| 2:29.222 | but |  |  |  |
| 2:29.223 | this is |  |  |  |
| 2:29.224 | Metatron |  |  |  |
| 2:29.225 | whose name |  |  |  |
| 2:29.226 | is like |  |  |  |
| 2:29.227 | his master's name |  |  |  |
| 2:29.228 | and so it is |  |  |  |
| 2:29.229 | in the Yalkut |  |  | Yalkut: a Midrash anthology. The title is transliterated; its sense is 'compilation'. |
| 2:29.230 | and as for |  |  |  |
| 2:29.231 | what you wrote |  |  |  |
| 2:29.232 | that our rabbis |  |  |  |
| 2:29.233 | the true ones |  |  |  |
| 2:29.234 | said |  |  |  |
| 2:29.235 | it is cited |  |  |  |
| 2:29.236 | in Menorat |  |  |  |
| 2:29.237 | ha-Ma'or |  |  | Menorat ha-Ma'or: 'the candelabrum of the light'. The title is transliterated; the meaning is given here. |
| 2:29.238 | that at the time |  |  |  |
| 2:29.239 | of the giving |  |  |  |
| 2:29.240 | of the Torah |  |  |  |
| 2:29.241 | tore |  |  |  |
| 2:29.242 | the Holy One, blessed be He |  |  | the Holy One, blessed be He (הקב"ה). |
| 2:29.243 | 7 |  |  |  |
| 2:29.244 | heavens |  |  |  |
| 2:29.245 | etc. |  |  |  |
| 2:29.246 | and in |  |  |  |
| 2:29.247 | Midrash Rabbah |  |  |  |
| 2:29.248 | they said |  |  |  |
| 2:29.249 | I am |  |  |  |
| 2:29.250 | the LORD |  |  | the LORD: Exod 20:2 ('I am the LORD your God'), as quoted in the Midrash. |
| 2:29.251 | your God |  |  |  |
| 2:29.252 | I |  |  |  |
| 2:29.253 | the first |  |  |  |
| 2:29.254 | and I |  |  |  |
| 2:29.255 | last |  |  |  |
| 2:29.256 | I am |  |  |  |
| 2:29.257 | first |  |  |  |
| 2:29.258 | for |  |  |  |
| 2:29.259 | I have |  |  |  |
| 2:29.260 | no father |  |  |  |
| 2:29.261 | and I am |  |  |  |
| 2:29.262 | last |  |  |  |
| 2:29.263 | for |  |  |  |
| 2:29.264 | I have |  |  |  |
| 2:29.265 | no son |  |  |  |
| 2:29.266 | and besides Me |  |  |  |
| 2:29.267 | there is |  |  |  |
| 2:29.268 | no God |  |  | God: Isa 44:6 ('besides Me there is no God'), in the Midrash quotation. |
| 2:29.269 | for |  |  |  |
| 2:29.270 | I have |  |  |  |
| 2:29.271 | no brother |  |  |  |
| 2:29.272 | who |  |  |  |
| 2:29.273 | is it |  |  |  |
| 2:29.274 | that says |  |  |  |
| 2:29.275 | he has |  |  |  |
| 2:29.276 | to him |  |  |  |
| 2:29.277 | a father |  |  |  |
| 2:29.278 | or |  |  |  |
| 2:29.279 | a son |  |  | The printed word is 'בם' ('in them'). Read as 'son' (בן) from the sense of the list; see apparatus. |
| 2:29.280 | or |  |  |  |
| 2:29.281 | a brother |  |  |  |
| 2:29.282 | but |  |  |  |
| 2:29.283 | you |  |  |  |
| 2:29.284 | spread |  |  |  |
| 2:29.285 | slander |  |  |  |
| 2:29.286 | against |  |  |  |
| 2:29.287 | the sages |  |  |  |
| 2:29.288 | who |  |  |  |
| 2:29.289 | say |  |  |  |
| 2:29.290 | so |  |  |  |
| 2:29.291 | and they |  |  |  |
| 2:29.292 | already |  |  |  |
| 2:29.293 | have explained |  |  |  |
| 2:29.294 | their words |  |  |  |
| 2:29.295 | of blessed memory |  |  | 'of blessed memory' (ז"ל). |
| 2:29.296 | that all |  |  |  |
| 2:29.297 | the epithets |  |  |  |
| 2:29.298 | and the attributes |  |  |  |
| 2:29.299 | are |  |  |  |
| 2:29.300 | in the sefirot |  |  |  |
| 2:29.301 | to shock |  |  | Literally 'to break the ear': an idiom for shocking the hearer, used in the same sense here. |
| 2:29.302 | (object marker) |  |  |  |
| 2:29.303 | the ear |  |  |  |
| 2:29.304 | go |  |  |  |
| 2:29.305 | now |  |  |  |
| 2:29.306 | and see |  |  |  |
| 2:29.307 | what |  |  |  |
| 2:29.308 | wrote |  |  |  |
| 2:29.309 | Mahari |  |  | Mahari: 'our teacher, the Rav Rabbi Y.' (morenu ha-Rav Rabbi Y.). The first name is not given in the text. |
| 2:29.310 | Gikatilia |  |  |  |
| 2:29.311 | in the introduction |  |  |  |
| 2:29.312 | to |  |  |  |
| 2:29.313 | Sha'arei Orah |  |  | Sha'arei Orah: 'Gates of Light'. The title is transliterated. |
| 2:29.314 | and the divine |  |  | 'the divine': an honorific before the name, not a claim about the man. |
| 2:29.315 | R. M. |  |  | R. M.: the initials are not expanded in the text. The surname is Ricanti (Hebrew spelling). |
| 2:29.316 | Ricanti |  |  |  |
| 2:29.317 | and R. Y. |  |  | R. Y.: the initials are not expanded in the text. |
| 2:29.318 | Chayyim |  |  |  |
| 2:29.319 | and these are his words |  |  | 'and these are his words': the formula that opens the quotation. |
| 2:29.320 | in sum |  |  | 'In sum' (סוף דבר: literally 'the end of the matter'). |
| 2:29.321 | the end of the matter |  |  |  |
| 2:29.322 | all |  |  |  |
| 2:29.323 | wherever |  |  |  |
| 2:29.324 | you see |  |  |  |
| 2:29.325 | things |  |  |  |
| 2:29.326 | that are not |  |  |  |
| 2:29.327 | fitting |  |  |  |
| 2:29.328 | to say |  |  |  |
| 2:29.329 | of the Creator |  |  |  |
| 2:29.330 | may He be blessed |  |  | may He be blessed (ית'). |
| 2:29.331 | such as |  |  |  |
| 2:29.332 | measure |  |  |  |
| 2:29.333 | of stature |  |  |  |
| 2:29.334 | and the like |  |  |  |
| 2:29.335 | of them |  |  |  |
| 2:29.336 | all |  |  |  |
| 2:29.337 | is said |  |  |  |
| 2:29.338 | of |  |  |  |
| 2:29.339 | the sefirot |  |  |  |
| 2:29.340 | and when you see |  |  |  |
| 2:29.341 | them speak |  |  |  |
| 2:29.342 | in the manner |  |  |  |
| 2:29.343 | of praise |  |  |  |
| 2:29.344 | and thanksgiving |  |  |  |
| 2:29.345 | all |  |  |  |
| 2:29.346 | is said |  |  |  |
| 2:29.347 | of |  |  |  |
| 2:29.348 | the Creator |  |  |  |
| 2:29.349 | may He be blessed |  |  | may He be blessed (ית'). |
| 2:29.350 | who is within |  |  |  |
| 2:29.351 | and outside |  |  |  |
| 2:29.352 | them |  |  |  |
| 2:29.353 | for |  |  |  |
| 2:29.354 | there is not |  |  |  |
| 2:29.355 | anything |  |  |  |
| 2:29.356 | that limits Him |  |  |  |
| 2:29.357 | therefore |  |  |  |
| 2:29.358 | so |  |  |  |
| 2:29.359 | one may not |  |  |  |
| 2:29.360 | say |  |  |  |
| 2:29.361 | of the Creator |  |  |  |
| 2:29.362 | may He be blessed |  |  | may He be blessed (ית'). |
| 2:29.363 | not |  |  |  |
| 2:29.364 | right |  |  |  |
| 2:29.365 | and not |  |  |  |
| 2:29.366 | left |  |  |  |
| 2:29.367 | not |  |  |  |
| 2:29.368 | face |  |  |  |
| 2:29.369 | and not |  |  |  |
| 2:29.370 | back |  |  |  |
| 2:29.371 | etc. |  |  | 'etc.': the quotation is cut short in the text. |
| 2:29.372 | see there |  |  | 'see there' (עיין שם). |
| 2:29.373 | and as for |  |  |  |
| 2:29.374 | what you wrote |  |  |  |
| 2:29.375 | that the opinion |  |  |  |
| 2:29.376 | of the Kabbalists |  |  |  |
| 2:29.377 | that the sefirot |  |  |  |
| 2:29.378 | descend in a chain |  |  | descend in a chain: the chain of descent (השתלשלות), here as the verb. |
| 2:29.379 | by nature |  |  |  |
| 2:29.380 | like a son |  |  |  |
| 2:29.381 | from |  |  |  |
| 2:29.382 | the father |  |  |  |
| 2:29.383 | far be it |  |  |  |
| 2:29.384 | that they |  |  |  |
| 2:29.385 | said |  |  |  |
| 2:29.386 | so |  |  |  |
| 2:29.387 | on the contrary |  |  |  |
| 2:29.388 | they said |  |  |  |
| 2:29.389 | far be it |  |  |  |
| 2:29.390 | to believe |  |  |  |
| 2:29.391 | or to bring |  |  |  |
| 2:29.392 | to mind |  |  |  |
| 2:29.393 | the heart |  |  |  |
| 2:29.394 | that the sefirot |  |  |  |
| 2:29.395 | are |  |  |  |
| 2:29.396 | part |  |  |  |
| 2:29.397 | of |  |  |  |
| 2:29.398 | Ein Sof |  |  | Ein Sof: literally 'without end'. It is used as a name with no article. |
| 2:29.399 | that came out |  |  |  |
| 2:29.400 | from it |  |  |  |
| 2:29.401 | and descended |  |  | 'descended': the verb of the chain of descent. |
| 2:29.402 | from cause |  |  | 'cause' (עילה). |
| 2:29.403 | to effect |  |  | 'effect' (עלול). |
| 2:29.404 | for |  |  |  |
| 2:29.405 | it |  |  |  |
| 2:29.406 | is a sin |  |  | 'criminal' (פלילי): a sin for which one is judged, a legal term. |
| 2:29.407 | criminal |  |  |  |
| 2:29.408 | that Ein |  |  |  |
| 2:29.409 | Sof |  |  |  |
| 2:29.410 | is not |  |  |  |
| 2:29.411 | divided |  |  |  |
| 2:29.412 | into parts |  |  |  |
| 2:29.413 | God forbid |  |  | 'God forbid' (ח"ו). |
| 2:29.414 | and it does not |  |  |  |
| 2:29.415 | receive |  |  |  |
| 2:29.416 | addition |  |  |  |
| 2:29.417 | and not |  |  |  |
| 2:29.418 | diminution |  |  |  |
| 2:29.419 | but |  |  |  |
| 2:29.420 | its existence |  |  |  |
| 2:29.421 | is constant |  |  |  |
| 2:29.422 | always |  |  |  |
| 2:29.423 | without |  |  |  |
| 2:29.424 | any |  |  |  |
| 2:29.425 | change |  |  |  |
| 2:29.426 | whatever |  |  |  |
| 2:29.427 | and He |  |  |  |
| 2:29.428 | originated them |  |  | 'originated': renewed, created anew. The writer's word for the origin of the sefirot. |
| 2:29.429 | origination |  |  |  |
| 2:29.430 | absolute |  |  |  |

Phrases:
- 2:29.1 2:29.2 2:29.3 2:29.4 2:29.5 2:29.6 2:29.7 | I have gone on at length about all this, to set out for you the words of Kisse Eliyahu, | I have gone on at length about all this to set out for you the words of Kisse Eliyahu
- 2:29.8 2:29.9 2:29.10 2:29.11 2:29.12 | but I have already mentioned it above. | but I have taken up already mentioned it above
- 2:29.13 2:29.14 2:29.15 | And there is proof from this | and there is proof from this
- 2:29.16 2:29.17 2:29.18 2:29.19 2:29.20 2:29.21 2:29.22 2:29.23 2:29.24 2:29.25 2:29.26 2:29.27 2:29.28 2:29.29 | that what you ask me is not from the motive of your request for the truth, as you said, but rather to seek some pretext. | that what you ask me is not from the motive of your request for the truth as you said but rather to seek some pretext
- 2:29.30 2:29.31 2:29.32 2:29.33 2:29.34 2:29.35 2:29.36 2:29.37 2:29.38 2:29.39 2:29.40 | And who am I that you should hold me to be a Kabbalist, for I am brutish and do not know. | and what am I that you hold me as a Kabbalist for I am brutish and do not know
- 2:29.41 2:29.42 2:29.43 2:29.44 2:29.45 2:29.46 | And all my foolish talk is not even like the chirping of birds, | and all my foolish talk even like the chirping of birds is not
- 2:29.47 2:29.48 2:29.49 2:29.50 2:29.51 2:29.52 2:29.53 2:29.54 2:29.55 2:29.56 2:29.57 2:29.58 2:29.59 2:29.60 | for the birds are fed without distress, and on them there lies neither the yoke of kingship nor the yoke of the way of the land, | for the birds are fed without distress and there is not on them neither the yoke of kingship nor the yoke of the way of the land
- 2:29.61 2:29.62 2:29.63 2:29.64 2:29.65 2:29.66 2:29.67 2:29.68 2:29.69 2:29.70 2:29.71 2:29.72 2:29.73 | and chirping can be done with intent, but I, under so much burden of the times and the cares of livelihood, have my mind and my thoughts confused. | and it is possible chirping with intent but I so much burden of the times and the worry of livelihood confuse my mind and my thoughts
- 2:29.74 2:29.75 2:29.76 2:29.77 2:29.78 2:29.79 2:29.80 2:29.81 | And furthermore, I have not received from a wise Kabbalist mouth to mouth. | and furthermore that not I have received from a wise man Kabbalist from mouth to mouth
- 2:29.82 2:29.83 2:29.84 2:29.85 2:29.86 2:29.87 2:29.88 2:29.89 | Would that the LORD relieve me of the troubles of time and of livelihood, | and would that will give and relieve me the LORD of the troubles of time and of livelihood
- 2:29.90 2:29.91 2:29.92 2:29.93 2:29.94 2:29.95 2:29.96 | and provide me a great Kabbalist who would set me on the paths of uprightness, | and provide me a great Kabbalist who would set me on the paths of uprightness
- 2:29.97 2:29.98 2:29.99 2:29.100 | and that I might merit to be among the sons of the ascent. | and I will merit to be from the sons of ascent
- 2:29.101 2:29.102 2:29.103 2:29.104 2:29.105 2:29.106 2:29.107 2:29.108 2:29.109 2:29.110 2:29.111 | As for what I answered you the second time, it is the words of Kisse Eliyahu that you turned a blind eye to. | as for what I answered to you time second is the words of Kisse Eliyahu that you turned a blind eye to them
- 2:29.112 2:29.113 2:29.114 2:29.115 2:29.116 | And as for what you wrote, that you do not ascribe corporeality to the sefirot, | and as for what you wrote that you are not ascribing corporeality to the sefirot
- 2:29.117 2:29.118 2:29.119 2:29.120 2:29.121 2:29.122 | But did you not write: "If so, there would be a force in a body"? | but did not you write if so there would be a force in a body
- 2:29.123 2:29.124 2:29.125 2:29.126 2:29.127 2:29.128 2:29.129 2:29.130 | And this time you say that light is a subtle body. | and at this time this you say that light a body subtle
- 2:29.131 2:29.132 2:29.133 2:29.134 2:29.135 2:29.136 2:29.137 2:29.138 2:29.139 2:29.140 2:29.141 2:29.142 2:29.143 2:29.144 2:29.145 2:29.146 2:29.147 2:29.148 2:29.149 | And I already wrote to you that what they called them by the name light is only in order to shock the ear, so that you ascribe corporeality one time and make excuses another. | and already I wrote to you that what they called them by the name light is not but in order to shock (object marker) the ear at one time you ascribe corporeality and another time you make excuses
- 2:29.150 2:29.151 2:29.152 2:29.153 2:29.154 2:29.155 2:29.156 2:29.157 2:29.158 2:29.159 | And why do all these doubts come into your mind, except because you ascribe corporeality? | and why come up in your mind all these doubts except because you ascribe corporeality
- 2:29.160 2:29.161 2:29.162 2:29.163 2:29.164 2:29.165 2:29.166 2:29.167 2:29.168 | As for what you wrote: "Who is it that governs all these hidden worlds?" | as for what you wrote who is it that governs all these worlds hidden
- 2:29.169 2:29.170 2:29.171 2:29.172 | The One who brings them into being is the one who keeps them alive. | the one who brings into being is keeps alive them
- 2:29.173 2:29.174 2:29.175 2:29.176 2:29.177 2:29.178 | And as for what you wrote, that Ze'ir Anpin is found to be the god of the lower ones. | and as for what you wrote that is found Ze'ir Anpin god of the lower ones
- 2:29.179 2:29.180 2:29.181 2:29.182 2:29.183 2:29.184 2:29.185 2:29.186 | Have I not written to you that he is nothing but a vessel for the acts of the Creator? | have I not written to you that he is not but a vessel for the acts of the Creator
- 2:29.187 2:29.188 2:29.189 2:29.190 2:29.191 2:29.192 2:29.193 2:29.194 2:29.195 2:29.196 2:29.197 2:29.198 | And as for what you wrote, that some of them said "all the great work of the LORD," from which it follows that there is a small one? | and as for what you wrote that some of them said (object marker) all the work of the LORD great from which it follows that there is a small one
- 2:29.199 2:29.200 2:29.201 2:29.202 2:29.203 2:29.204 2:29.205 2:29.206 2:29.207 2:29.208 2:29.209 | Those who say so have already explained that they meant Metatron, whose name is like his master's, | those who say so already have explained they their intent was on Metatron whose name is like the name of his master
- 2:29.210 2:29.211 2:29.212 2:29.213 2:29.214 2:29.215 2:29.216 2:29.217 2:29.218 2:29.219 2:29.220 2:29.221 2:29.222 2:29.223 2:29.224 2:29.225 2:29.226 2:29.227 | as they said in the Talmud: "And to Moses He said, 'Come up to the LORD'; it should have said 'to Me'. But this is Metatron, whose name is like his master's." | as they said in the Talmud and to Moses he said come up to the LORD to Me it should have said it but this is Metatron whose name is like his master's name
- 2:29.228 2:29.229 | and so it is in the Yalkut. | and so it is in the Yalkut
- 2:29.230 2:29.231 2:29.232 2:29.233 2:29.234 | And as for what you wrote, that our true rabbis said, | and as for what you wrote that our rabbis the true ones said
- 2:29.235 2:29.236 2:29.237 2:29.238 2:29.239 2:29.240 2:29.241 2:29.242 2:29.243 2:29.244 2:29.245 | it is cited in Menorat ha-Ma'or that at the time of the giving of the Torah the Holy One, blessed be He, tore 7 heavens, etc. | it is cited in Menorat ha-Ma'or that at the time of the giving of the Torah tore the Holy One, blessed be He 7 heavens etc.
- 2:29.246 2:29.247 2:29.248 2:29.249 2:29.250 2:29.251 2:29.252 2:29.253 2:29.254 2:29.255 | And in Midrash Rabbah they said: "I am the LORD your God; I am the first and I am the last. | and in Midrash Rabbah they said I am the LORD your God I the first and I last
- 2:29.256 2:29.257 2:29.258 2:29.259 2:29.260 2:29.261 2:29.262 2:29.263 2:29.264 2:29.265 2:29.266 2:29.267 2:29.268 2:29.269 2:29.270 2:29.271 | I am the first, for I have no father, and I am the last, for I have no son; and besides Me there is no God, for I have no brother." | I am first for I have no father and I am last for I have no son and besides Me there is no God for I have no brother
- 2:29.272 2:29.273 2:29.274 2:29.275 2:29.276 2:29.277 2:29.278 2:29.279 2:29.280 2:29.281 | Who is it that says he has a father, or a son, or a brother? | who is it that says he has to him a father or a son or a brother
- 2:29.282 2:29.283 2:29.284 2:29.285 2:29.286 2:29.287 2:29.288 2:29.289 2:29.290 | but you spread slander against the sages who say so, | but you spread slander against the sages who say so
- 2:29.291 2:29.292 2:29.293 2:29.294 2:29.295 | and they have already explained their words, of blessed memory, | and they already have explained their words of blessed memory
- 2:29.296 2:29.297 2:29.298 2:29.299 2:29.300 2:29.301 2:29.302 2:29.303 | that all the epithets and the attributes are in the sefirot, to shock the ear. | that all the epithets and the attributes are in the sefirot to shock (object marker) the ear
- 2:29.304 2:29.305 2:29.306 2:29.307 2:29.308 2:29.309 2:29.310 2:29.311 2:29.312 2:29.313 | Go now and see what Mahari Gikatilia wrote in the introduction to Sha'arei Orah, | go now and see what wrote Mahari Gikatilia in the introduction to Sha'arei Orah
- 2:29.314 2:29.315 2:29.316 2:29.317 2:29.318 2:29.319 | and the divine R. M. Ricanti, and R. Y. Chayyim, and these are his words: | and the divine R. M. Ricanti and R. Y. Chayyim and these are his words
- 2:29.320 2:29.321 2:29.322 2:29.323 2:29.324 2:29.325 2:29.326 2:29.327 2:29.328 2:29.329 2:29.330 2:29.331 2:29.332 2:29.333 2:29.334 2:29.335 2:29.336 2:29.337 2:29.338 2:29.339 | "In sum, wherever you see things that are not fitting to say of the Creator, may He be blessed, such as measure of stature and the like, it is all said of the sefirot; | in sum the end of the matter all wherever you see things that are not fitting to say of the Creator may He be blessed such as measure of stature and the like of them all is said of the sefirot
- 2:29.340 2:29.341 2:29.342 2:29.343 2:29.344 2:29.345 2:29.346 2:29.347 2:29.348 2:29.349 2:29.350 2:29.351 2:29.352 | and when you see them speak in the manner of praise and thanksgiving, it is all said of the Creator, may He be blessed, who is within them and outside them, | and when you see them speak in the manner of praise and thanksgiving all is said of the Creator may He be blessed who is within and outside them
- 2:29.353 2:29.354 2:29.355 2:29.356 | for nothing limits Him; | for there is not anything that limits Him
- 2:29.357 2:29.358 2:29.359 2:29.360 2:29.361 2:29.362 2:29.363 2:29.364 2:29.365 2:29.366 2:29.367 2:29.368 2:29.369 2:29.370 2:29.371 | therefore one may not say of the Creator, may He be blessed, not right and not left, not face and not back, etc." | therefore so one may not say of the Creator may He be blessed not right and not left not face and not back etc.
- 2:29.372 | See there. | see there
- 2:29.373 2:29.374 2:29.375 2:29.376 2:29.377 2:29.378 2:29.379 2:29.380 2:29.381 2:29.382 | And as for what you wrote, that the opinion of the Kabbalists is that the sefirot descend in a chain, by nature like a son from the father, | and as for what you wrote that the opinion of the Kabbalists that the sefirot descend in a chain by nature like a son from the father
- 2:29.383 2:29.384 2:29.385 2:29.386 | Far be it that they said so. | far be it that they said so
- 2:29.387 2:29.388 | On the contrary, they said: | on the contrary they said
- 2:29.389 2:29.390 2:29.391 2:29.392 2:29.393 2:29.394 2:29.395 2:29.396 2:29.397 2:29.398 | "Far be it to believe or to bring to mind that the sefirot are part of Ein Sof | far be it to believe or to bring to mind the heart that the sefirot are part of Ein Sof
- 2:29.399 2:29.400 2:29.401 2:29.402 2:29.403 2:29.404 2:29.405 2:29.406 2:29.407 | that came out of it and descended from cause to effect, for it is a criminal sin. | that came out from it and descended from cause to effect for it is a sin criminal
- 2:29.408 2:29.409 2:29.410 2:29.411 2:29.412 2:29.413 | For Ein Sof is not divided into parts, God forbid, | that Ein Sof is not divided into parts God forbid
- 2:29.414 2:29.415 2:29.416 2:29.417 2:29.418 2:29.419 2:29.420 2:29.421 2:29.422 2:29.423 2:29.424 2:29.425 2:29.426 | and it receives no addition and no diminution, but its existence is constant always, without any change whatever, | and it does not receive addition and not diminution but its existence is constant always without any change whatever
- 2:29.427 2:29.428 2:29.429 2:29.430 | and He originated them, an absolute origination." | and He originated them origination absolute

Grammar notes:
- The paragraph is the respondent's reply to his questioner, in the respondent's own voice. 'You' is the questioner, and the English keeps the second person.
- Each point of the questioner's letter is answered in turn, opening with 'as for what you wrote' (2:29.101, 112, 160, 173, 187, 230, 373). The questions with 'but did you not' and 'and why' keep their force as questions.
- Clauses with 'that' after 'you wrote' are reported speech and take no quotation marks. 'You wrote: if so ...' and 'you wrote: who is it ...' quote the questioner's own words, so they are in quotation marks.
- The Talmud passage turns on the word 'to Me' (אלי): the verse should have said 'to Me', and the Talmud draws its conclusion from the difference. The English gives the Hebrew point in single quotation marks inside the double ones.
- The verb 'to ascribe corporeality' (להגשים) and its forms run through the paragraph. It is the error the writer charges against the questioner, and it is rendered the same way each time.

Translator's note: The respondent's reply (2:29) in the respondent's voice. Honorifics, book titles and abbreviations are rendered by the decisions for this work; see the notes at the first occurrence of each.

Commentary rulings:
- apparatus: Witness: Wikisource base text. The 1931 scan is not yet collated, so the text is translated as read and this is flagged here (decisions, Q2).
- apparatus: Bible verses: Exod 20:2 ('I am the LORD your God', 2:29.249-251, in the Midrash quotation); Isa 44:6 ('I am the first and I am the last, and besides Me there is no God', 2:29.252-271); Exod 24:1 ('And to Moses He said: Come up to the LORD', 2:29.213-218, in the Talmud quotation); Deut 11:7 ('all the great work of the LORD', 2:29.191-194); Ps 73:22 ('I am brutish and do not know', 2:29.38-40); Prov 4:11 ('the paths of uprightness', 2:29.95-96).
- apparatus: Variant: the printed 'בם' (2:29.279), 'in them'. Read as 'בן' (son), the likelier word in 'a father or a son or a brother'. The printed word is kept in the Hebrew; the English follows the sense.
- apparatus: The Talmud and Midrash passages carry no reference in the text, and none is added.
- apparatus: The quotation 'In sum ...' ends at 2:29.371 and is followed by 'See there' (2:29.372). The quotation 'Far be it ...' closes at 2:29.430.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

