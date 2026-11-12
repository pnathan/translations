# Translation pack: milhamot-hashem perek-8-14 (Milchamot ha-Shem 8:17–8:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-14`.

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
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
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
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **השפע** → "the flow".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **זווג** → "coupling".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **כ"ש / ק"ו** → "all the more so".
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **רוע** לב → "wickedness of heart".
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


## Milchamot ha-Shem 8:17

Hebrew (Wikisource): קיא) ואלה הדעות הן הנה הדעות הרעות שהיו לאבותינו בימי השופטים ובימי ירבעם בן נבט ושאר מלכי ישראל וקצת מלכי יהודה בזמן בית ראשון אשר עליהם נתנבאו כל הנביאים, וקראו בגרון הוי החוקקים חקקי און, ומכתבי עמל כתבו, והיה להם דבר ה' צו לצו לצו קו לקו קו זעיר שם זעיר וכמאמר רז"ל אין צו אלא ע"ז שנאמר כי הואיל הלך אחרי צו וכו' וכן קו לקו הם הפרצופים בעלי שלשה קוים המוזכרים בדברי המקובלים כמו שכתב מהרח"ו בעץ חיים שער עגולים ויושר, שספירות היושר הם בציור שלשה קוים, כמראה האדם בעל ראש, וזרועות ושוקים וגוף ורגלים וכו', וכן זעיר שם זעיר שם הוא זעיר אנפין, (קצר אפים) ונאמר ויעזבו את ה' אלהיהם ויעבדו את הבעלים ואת העשתרות, הבעלים הם האלוהות הזכרים, אבא וזעיר אנפין, ונקראו בעלים על שם הזווגים והבעילות שמזדווגין עם הנקיבות, כמו שאמרו רז"ל בירושלמי שבת פרק אמר ר' עקיבא ובמסכת אבות ע"ז פרק כל הצלמים וז"ל אמר ר' הונא רבי חמא בשם רב, הבעל ראש גוייה היה לו ונאפון הוא, וישימו להם בעל ברית לאלהים, עכ"ל הירושלמי הרי מבואר בהדיא מדברי הירוש' שנק' הצלם הזכר מעל על שם שיש לו ראש גויה והוא בועל ומנאף, וידוע הוא כי האלילים הקדמונים עשו להם צלמים מכוונים נגד הצורות והפרצופים שחשבו שיש בשמים, כדי להמשיך השפע מהצורות העליונות להם, כי הם עשוים בצלמם ודמותם, וגם כן נקראו בעלים על שם שעובדיהם שמו אותם אדונים להם כמו שכתב הרד"ק וכלי יקר. ואת העשתרות הם נקבות האלוהות שהם אימא ומלכית, כי עשתרות צאנך הם נקבות הצאן כמו שכתבו המפרשים, וכתן נאמר ואת העשתרות בלי סמיכות מלת הצאן, והם נקיבות האלוהות הנזכרים. וגם בין הכשרים היו האמונות האלו פושטים ולכן אמר נבוכדנאצר הרשע בעת שראה מלאך ה' בכבשן האש שירד להציל לחנניה מיאל ועזריה מן האור, וריויה די רביעאה דמי לבר אלהין, לפי שחשב שהזעיר אנפין הנקרא בן לאבא ואימא ירד להצילם מן האש, כדגרסינן בירושלמי שבת פרק במה אשה, וריוה די רביעאה דמי לבא אלהין אמר רבי ראובן באותה שעה ירד מלאך וסטרו לההוא רשיעא על פיו ואמר ליה תקין מילך, ובר אית ליה? חזר ואמר בריך אלההון די שדרך מישך ועביד נגו, די שלח מלאכיה ושיזיב לעבדוהי, בריה אין כתיב כאן אלא מלאכיה, ע"כ. ומזה יתברר לך שהדעות האלו של מחבר הזהר והמקובלים החדשים, זרים המה לדעת תורתינו הקדושה, ומהעמים אשר השאיר ה' לנסות בם את ישראל באו לנו, ונטו ישראל אחריהם בימי שפוט השופטים, וגם מן הכשדים שגלו ביניהם בעת חרבן בית ראשון.

| id | word | punctuation after |
|---|---|---|
| 8:17.1 | קיא | ) |
| 8:17.2 | ואלה |  |
| 8:17.3 | הדעות |  |
| 8:17.4 | הן |  |
| 8:17.5 | הנה |  |
| 8:17.6 | הדעות |  |
| 8:17.7 | הרעות |  |
| 8:17.8 | שהיו |  |
| 8:17.9 | לאבותינו |  |
| 8:17.10 | בימי |  |
| 8:17.11 | השופטים |  |
| 8:17.12 | ובימי |  |
| 8:17.13 | ירבעם |  |
| 8:17.14 | בן |  |
| 8:17.15 | נבט |  |
| 8:17.16 | ושאר |  |
| 8:17.17 | מלכי |  |
| 8:17.18 | ישראל |  |
| 8:17.19 | וקצת |  |
| 8:17.20 | מלכי |  |
| 8:17.21 | יהודה |  |
| 8:17.22 | בזמן |  |
| 8:17.23 | בית |  |
| 8:17.24 | ראשון |  |
| 8:17.25 | אשר |  |
| 8:17.26 | עליהם |  |
| 8:17.27 | נתנבאו |  |
| 8:17.28 | כל |  |
| 8:17.29 | הנביאים | , |
| 8:17.30 | וקראו |  |
| 8:17.31 | בגרון |  |
| 8:17.32 | הוי |  |
| 8:17.33 | החוקקים |  |
| 8:17.34 | חקקי |  |
| 8:17.35 | און | , |
| 8:17.36 | ומכתבי |  |
| 8:17.37 | עמל |  |
| 8:17.38 | כתבו | , |
| 8:17.39 | והיה |  |
| 8:17.40 | להם |  |
| 8:17.41 | דבר |  |
| 8:17.42 | ה | ' |
| 8:17.43 | צו |  |
| 8:17.44 | לצו |  |
| 8:17.45 | לצו |  |
| 8:17.46 | קו |  |
| 8:17.47 | לקו |  |
| 8:17.48 | קו |  |
| 8:17.49 | זעיר |  |
| 8:17.50 | שם |  |
| 8:17.51 | זעיר |  |
| 8:17.52 | וכמאמר |  |
| 8:17.53 | רז"ל |  |
| 8:17.54 | אין |  |
| 8:17.55 | צו |  |
| 8:17.56 | אלא |  |
| 8:17.57 | ע"ז |  |
| 8:17.58 | שנאמר |  |
| 8:17.59 | כי |  |
| 8:17.60 | הואיל |  |
| 8:17.61 | הלך |  |
| 8:17.62 | אחרי |  |
| 8:17.63 | צו |  |
| 8:17.64 | וכו | ' |
| 8:17.65 | וכן |  |
| 8:17.66 | קו |  |
| 8:17.67 | לקו |  |
| 8:17.68 | הם |  |
| 8:17.69 | הפרצופים |  |
| 8:17.70 | בעלי |  |
| 8:17.71 | שלשה |  |
| 8:17.72 | קוים |  |
| 8:17.73 | המוזכרים |  |
| 8:17.74 | בדברי |  |
| 8:17.75 | המקובלים |  |
| 8:17.76 | כמו |  |
| 8:17.77 | שכתב |  |
| 8:17.78 | מהרח"ו |  |
| 8:17.79 | בעץ |  |
| 8:17.80 | חיים |  |
| 8:17.81 | שער |  |
| 8:17.82 | עגולים |  |
| 8:17.83 | ויושר | , |
| 8:17.84 | שספירות |  |
| 8:17.85 | היושר |  |
| 8:17.86 | הם |  |
| 8:17.87 | בציור |  |
| 8:17.88 | שלשה |  |
| 8:17.89 | קוים | , |
| 8:17.90 | כמראה |  |
| 8:17.91 | האדם |  |
| 8:17.92 | בעל |  |
| 8:17.93 | ראש | , |
| 8:17.94 | וזרועות |  |
| 8:17.95 | ושוקים |  |
| 8:17.96 | וגוף |  |
| 8:17.97 | ורגלים |  |
| 8:17.98 | וכו | ', |
| 8:17.99 | וכן |  |
| 8:17.100 | זעיר |  |
| 8:17.101 | שם |  |
| 8:17.102 | זעיר |  |
| 8:17.103 | שם |  |
| 8:17.104 | הוא |  |
| 8:17.105 | זעיר |  |
| 8:17.106 | אנפין | , |
| 8:17.107 | קצר |  |
| 8:17.108 | אפים | ) |
| 8:17.109 | ונאמר |  |
| 8:17.110 | ויעזבו |  |
| 8:17.111 | את |  |
| 8:17.112 | ה | ' |
| 8:17.113 | אלהיהם |  |
| 8:17.114 | ויעבדו |  |
| 8:17.115 | את |  |
| 8:17.116 | הבעלים |  |
| 8:17.117 | ואת |  |
| 8:17.118 | העשתרות | , |
| 8:17.119 | הבעלים |  |
| 8:17.120 | הם |  |
| 8:17.121 | האלוהות |  |
| 8:17.122 | הזכרים | , |
| 8:17.123 | אבא |  |
| 8:17.124 | וזעיר |  |
| 8:17.125 | אנפין | , |
| 8:17.126 | ונקראו |  |
| 8:17.127 | בעלים |  |
| 8:17.128 | על |  |
| 8:17.129 | שם |  |
| 8:17.130 | הזווגים |  |
| 8:17.131 | והבעילות |  |
| 8:17.132 | שמזדווגין |  |
| 8:17.133 | עם |  |
| 8:17.134 | הנקיבות | , |
| 8:17.135 | כמו |  |
| 8:17.136 | שאמרו |  |
| 8:17.137 | רז"ל |  |
| 8:17.138 | בירושלמי |  |
| 8:17.139 | שבת |  |
| 8:17.140 | פרק |  |
| 8:17.141 | אמר |  |
| 8:17.142 | ר | ' |
| 8:17.143 | עקיבא |  |
| 8:17.144 | ובמסכת |  |
| 8:17.145 | אבות |  |
| 8:17.146 | ע"ז |  |
| 8:17.147 | פרק |  |
| 8:17.148 | כל |  |
| 8:17.149 | הצלמים |  |
| 8:17.150 | וז"ל |  |
| 8:17.151 | אמר |  |
| 8:17.152 | ר | ' |
| 8:17.153 | הונא |  |
| 8:17.154 | רבי |  |
| 8:17.155 | חמא |  |
| 8:17.156 | בשם |  |
| 8:17.157 | רב | , |
| 8:17.158 | הבעל |  |
| 8:17.159 | ראש |  |
| 8:17.160 | גוייה |  |
| 8:17.161 | היה |  |
| 8:17.162 | לו |  |
| 8:17.163 | ונאפון |  |
| 8:17.164 | הוא | , |
| 8:17.165 | וישימו |  |
| 8:17.166 | להם |  |
| 8:17.167 | בעל |  |
| 8:17.168 | ברית |  |
| 8:17.169 | לאלהים | , |
| 8:17.170 | עכ"ל |  |
| 8:17.171 | הירושלמי |  |
| 8:17.172 | הרי |  |
| 8:17.173 | מבואר |  |
| 8:17.174 | בהדיא |  |
| 8:17.175 | מדברי |  |
| 8:17.176 | הירוש | ' |
| 8:17.177 | שנק | ' |
| 8:17.178 | הצלם |  |
| 8:17.179 | הזכר |  |
| 8:17.180 | מעל |  |
| 8:17.181 | על |  |
| 8:17.182 | שם |  |
| 8:17.183 | שיש |  |
| 8:17.184 | לו |  |
| 8:17.185 | ראש |  |
| 8:17.186 | גויה |  |
| 8:17.187 | והוא |  |
| 8:17.188 | בועל |  |
| 8:17.189 | ומנאף | , |
| 8:17.190 | וידוע |  |
| 8:17.191 | הוא |  |
| 8:17.192 | כי |  |
| 8:17.193 | האלילים |  |
| 8:17.194 | הקדמונים |  |
| 8:17.195 | עשו |  |
| 8:17.196 | להם |  |
| 8:17.197 | צלמים |  |
| 8:17.198 | מכוונים |  |
| 8:17.199 | נגד |  |
| 8:17.200 | הצורות |  |
| 8:17.201 | והפרצופים |  |
| 8:17.202 | שחשבו |  |
| 8:17.203 | שיש |  |
| 8:17.204 | בשמים | , |
| 8:17.205 | כדי |  |
| 8:17.206 | להמשיך |  |
| 8:17.207 | השפע |  |
| 8:17.208 | מהצורות |  |
| 8:17.209 | העליונות |  |
| 8:17.210 | להם | , |
| 8:17.211 | כי |  |
| 8:17.212 | הם |  |
| 8:17.213 | עשוים |  |
| 8:17.214 | בצלמם |  |
| 8:17.215 | ודמותם | , |
| 8:17.216 | וגם |  |
| 8:17.217 | כן |  |
| 8:17.218 | נקראו |  |
| 8:17.219 | בעלים |  |
| 8:17.220 | על |  |
| 8:17.221 | שם |  |
| 8:17.222 | שעובדיהם |  |
| 8:17.223 | שמו |  |
| 8:17.224 | אותם |  |
| 8:17.225 | אדונים |  |
| 8:17.226 | להם |  |
| 8:17.227 | כמו |  |
| 8:17.228 | שכתב |  |
| 8:17.229 | הרד"ק |  |
| 8:17.230 | וכלי |  |
| 8:17.231 | יקר | . |
| 8:17.232 | ואת |  |
| 8:17.233 | העשתרות |  |
| 8:17.234 | הם |  |
| 8:17.235 | נקבות |  |
| 8:17.236 | האלוהות |  |
| 8:17.237 | שהם |  |
| 8:17.238 | אימא |  |
| 8:17.239 | ומלכית | , |
| 8:17.240 | כי |  |
| 8:17.241 | עשתרות |  |
| 8:17.242 | צאנך |  |
| 8:17.243 | הם |  |
| 8:17.244 | נקבות |  |
| 8:17.245 | הצאן |  |
| 8:17.246 | כמו |  |
| 8:17.247 | שכתבו |  |
| 8:17.248 | המפרשים | , |
| 8:17.249 | וכתן |  |
| 8:17.250 | נאמר |  |
| 8:17.251 | ואת |  |
| 8:17.252 | העשתרות |  |
| 8:17.253 | בלי |  |
| 8:17.254 | סמיכות |  |
| 8:17.255 | מלת |  |
| 8:17.256 | הצאן | , |
| 8:17.257 | והם |  |
| 8:17.258 | נקיבות |  |
| 8:17.259 | האלוהות |  |
| 8:17.260 | הנזכרים | . |
| 8:17.261 | וגם |  |
| 8:17.262 | בין |  |
| 8:17.263 | הכשרים |  |
| 8:17.264 | היו |  |
| 8:17.265 | האמונות |  |
| 8:17.266 | האלו |  |
| 8:17.267 | פושטים |  |
| 8:17.268 | ולכן |  |
| 8:17.269 | אמר |  |
| 8:17.270 | נבוכדנאצר |  |
| 8:17.271 | הרשע |  |
| 8:17.272 | בעת |  |
| 8:17.273 | שראה |  |
| 8:17.274 | מלאך |  |
| 8:17.275 | ה | ' |
| 8:17.276 | בכבשן |  |
| 8:17.277 | האש |  |
| 8:17.278 | שירד |  |
| 8:17.279 | להציל |  |
| 8:17.280 | לחנניה |  |
| 8:17.281 | מיאל |  |
| 8:17.282 | ועזריה |  |
| 8:17.283 | מן |  |
| 8:17.284 | האור | , |
| 8:17.285 | וריויה |  |
| 8:17.286 | די |  |
| 8:17.287 | רביעאה |  |
| 8:17.288 | דמי |  |
| 8:17.289 | לבר |  |
| 8:17.290 | אלהין | , |
| 8:17.291 | לפי |  |
| 8:17.292 | שחשב |  |
| 8:17.293 | שהזעיר |  |
| 8:17.294 | אנפין |  |
| 8:17.295 | הנקרא |  |
| 8:17.296 | בן |  |
| 8:17.297 | לאבא |  |
| 8:17.298 | ואימא |  |
| 8:17.299 | ירד |  |
| 8:17.300 | להצילם |  |
| 8:17.301 | מן |  |
| 8:17.302 | האש | , |
| 8:17.303 | כדגרסינן |  |
| 8:17.304 | בירושלמי |  |
| 8:17.305 | שבת |  |
| 8:17.306 | פרק |  |
| 8:17.307 | במה |  |
| 8:17.308 | אשה | , |
| 8:17.309 | וריוה |  |
| 8:17.310 | די |  |
| 8:17.311 | רביעאה |  |
| 8:17.312 | דמי |  |
| 8:17.313 | לבא |  |
| 8:17.314 | אלהין |  |
| 8:17.315 | אמר |  |
| 8:17.316 | רבי |  |
| 8:17.317 | ראובן |  |
| 8:17.318 | באותה |  |
| 8:17.319 | שעה |  |
| 8:17.320 | ירד |  |
| 8:17.321 | מלאך |  |
| 8:17.322 | וסטרו |  |
| 8:17.323 | לההוא |  |
| 8:17.324 | רשיעא |  |
| 8:17.325 | על |  |
| 8:17.326 | פיו |  |
| 8:17.327 | ואמר |  |
| 8:17.328 | ליה |  |
| 8:17.329 | תקין |  |
| 8:17.330 | מילך | , |
| 8:17.331 | ובר |  |
| 8:17.332 | אית |  |
| 8:17.333 | ליה | ? |
| 8:17.334 | חזר |  |
| 8:17.335 | ואמר |  |
| 8:17.336 | בריך |  |
| 8:17.337 | אלההון |  |
| 8:17.338 | די |  |
| 8:17.339 | שדרך |  |
| 8:17.340 | מישך |  |
| 8:17.341 | ועביד |  |
| 8:17.342 | נגו | , |
| 8:17.343 | די |  |
| 8:17.344 | שלח |  |
| 8:17.345 | מלאכיה |  |
| 8:17.346 | ושיזיב |  |
| 8:17.347 | לעבדוהי | , |
| 8:17.348 | בריה |  |
| 8:17.349 | אין |  |
| 8:17.350 | כתיב |  |
| 8:17.351 | כאן |  |
| 8:17.352 | אלא |  |
| 8:17.353 | מלאכיה | , |
| 8:17.354 | ע"כ | . |
| 8:17.355 | ומזה |  |
| 8:17.356 | יתברר |  |
| 8:17.357 | לך |  |
| 8:17.358 | שהדעות |  |
| 8:17.359 | האלו |  |
| 8:17.360 | של |  |
| 8:17.361 | מחבר |  |
| 8:17.362 | הזהר |  |
| 8:17.363 | והמקובלים |  |
| 8:17.364 | החדשים | , |
| 8:17.365 | זרים |  |
| 8:17.366 | המה |  |
| 8:17.367 | לדעת |  |
| 8:17.368 | תורתינו |  |
| 8:17.369 | הקדושה | , |
| 8:17.370 | ומהעמים |  |
| 8:17.371 | אשר |  |
| 8:17.372 | השאיר |  |
| 8:17.373 | ה | ' |
| 8:17.374 | לנסות |  |
| 8:17.375 | בם |  |
| 8:17.376 | את |  |
| 8:17.377 | ישראל |  |
| 8:17.378 | באו |  |
| 8:17.379 | לנו | , |
| 8:17.380 | ונטו |  |
| 8:17.381 | ישראל |  |
| 8:17.382 | אחריהם |  |
| 8:17.383 | בימי |  |
| 8:17.384 | שפוט |  |
| 8:17.385 | השופטים | , |
| 8:17.386 | וגם |  |
| 8:17.387 | מן |  |
| 8:17.388 | הכשדים |  |
| 8:17.389 | שגלו |  |
| 8:17.390 | ביניהם |  |
| 8:17.391 | בעת |  |
| 8:17.392 | חרבן |  |
| 8:17.393 | בית |  |
| 8:17.394 | ראשון | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

