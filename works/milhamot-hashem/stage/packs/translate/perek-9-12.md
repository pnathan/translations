# Translation pack: milhamot-hashem perek-9-12 (Milchamot ha-Shem 9:14–9:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-12`.

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
- **ז"ל** → "of blessed memory" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **ריב"ש** → "Ribash" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
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
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **עלה / עלול** → "cause / effect".
- **זווג** → "coupling".
- **מין / מינות** → "heretic / heresy".
- כופר **בעיקר** → "denies the fundamental principle".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **והוא שלא** → "provided that".
- ואין **צריך** לומר → "needless to say".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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
- **צנומות** דקות שדופות קדים → "withered, thin, blasted by the east wind".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **תעמוד / רעך** (Lev 19:16) → "stand by the blood of your fellow" (no
  "idly").
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 9:14

Hebrew (Wikisource): קלב) וכיון שהעיקר והיסוד המוסד אצלינו שהקב"ה אינו צריך לבריותיו. ואדרבא בריותיו הם הצריכים לו ליתן להם שפע וחיות כמאמר הכתוב ואתה מחיה את כולם. וכמו שכתב הרמב"ם בעיקר הראשון שמהשי"ת קיום כל מציאות הנמצאים ואם נצייר בלבינו מניעת מציאותו, בטלה מציאות כל נמצא, ואלו נעלה על לבינו ביטול הנמצאים כולם. לא בטלה מציאותו בשבילם ולא חסרה כלל, כי הבורא יתברך אין לו צורך במציאותו לזולתו וכו'. ואחרי הדברים והאמת האלה, מי פתי יסור להאמין לדברי הר"מ נבאי אשר הרבה דברים לקלקל עלינו את הכתובים ודברי רז"ל המדברים בפה מלא ובלשון צח שאין הקב"ה צריך לבריותיו להשלים חסרונו ח"ו. והרחיב את פיו לדבר גדולות וגבוהות, בראיות צנומות דקות ויצא עתק מפיו, על הריב"ש ז"ל, ושיבש עלינו פירוש הכתובים ודברי רז"ל ברוב דברים אשר כל משכיל ידע ויכיר שאין בהם ממש, והאמת עד לעצמו שכל עבודותינו ומעשה המצות שצונו הקב"ה בתורה אינם רק לצרף וללבן אותנו, ונהיה טהורים ונקיים במדותינו! ולהרחיק ממנו הדעות והאמונות הרעות ונהיה קרובים אליו יתברך כמאמר רז"ל לא ניתנו המצות אלא לצרוף בהן את הבריות, שלא יאמר אדם אי אפשר וכו'. וכן אמרו ז"ל וכי מה איכפת לו להקב"ה בין שוחט מן הצואר לשוחט מן העורף, הוי אומר לא ניתנו המצות אלא לצרוף בהן את הבריות שנאמר אמרת ה' צרופה! לא לזווגו הוא צריך, ולא לקשט לו את כלתו! כי הוא אחד ומיוחד מכל האחדים שבעולם. וכלל הדברים פשוט הוא שהתורה ניתנה לתועלת מקבליה ומקיימיה לא לתועלת נותנה, שנאמר לטוב לנו כל הימים.

| id | word | punctuation after |
|---|---|---|
| 9:14.1 | קלב | ) |
| 9:14.2 | וכיון |  |
| 9:14.3 | שהעיקר |  |
| 9:14.4 | והיסוד |  |
| 9:14.5 | המוסד |  |
| 9:14.6 | אצלינו |  |
| 9:14.7 | שהקב"ה |  |
| 9:14.8 | אינו |  |
| 9:14.9 | צריך |  |
| 9:14.10 | לבריותיו | . |
| 9:14.11 | ואדרבא |  |
| 9:14.12 | בריותיו |  |
| 9:14.13 | הם |  |
| 9:14.14 | הצריכים |  |
| 9:14.15 | לו |  |
| 9:14.16 | ליתן |  |
| 9:14.17 | להם |  |
| 9:14.18 | שפע |  |
| 9:14.19 | וחיות |  |
| 9:14.20 | כמאמר |  |
| 9:14.21 | הכתוב |  |
| 9:14.22 | ואתה |  |
| 9:14.23 | מחיה |  |
| 9:14.24 | את |  |
| 9:14.25 | כולם | . |
| 9:14.26 | וכמו |  |
| 9:14.27 | שכתב |  |
| 9:14.28 | הרמב"ם |  |
| 9:14.29 | בעיקר |  |
| 9:14.30 | הראשון |  |
| 9:14.31 | שמהשי"ת |  |
| 9:14.32 | קיום |  |
| 9:14.33 | כל |  |
| 9:14.34 | מציאות |  |
| 9:14.35 | הנמצאים |  |
| 9:14.36 | ואם |  |
| 9:14.37 | נצייר |  |
| 9:14.38 | בלבינו |  |
| 9:14.39 | מניעת |  |
| 9:14.40 | מציאותו | , |
| 9:14.41 | בטלה |  |
| 9:14.42 | מציאות |  |
| 9:14.43 | כל |  |
| 9:14.44 | נמצא | , |
| 9:14.45 | ואלו |  |
| 9:14.46 | נעלה |  |
| 9:14.47 | על |  |
| 9:14.48 | לבינו |  |
| 9:14.49 | ביטול |  |
| 9:14.50 | הנמצאים |  |
| 9:14.51 | כולם | . |
| 9:14.52 | לא |  |
| 9:14.53 | בטלה |  |
| 9:14.54 | מציאותו |  |
| 9:14.55 | בשבילם |  |
| 9:14.56 | ולא |  |
| 9:14.57 | חסרה |  |
| 9:14.58 | כלל | , |
| 9:14.59 | כי |  |
| 9:14.60 | הבורא |  |
| 9:14.61 | יתברך |  |
| 9:14.62 | אין |  |
| 9:14.63 | לו |  |
| 9:14.64 | צורך |  |
| 9:14.65 | במציאותו |  |
| 9:14.66 | לזולתו |  |
| 9:14.67 | וכו | '. |
| 9:14.68 | ואחרי |  |
| 9:14.69 | הדברים |  |
| 9:14.70 | והאמת |  |
| 9:14.71 | האלה | , |
| 9:14.72 | מי |  |
| 9:14.73 | פתי |  |
| 9:14.74 | יסור |  |
| 9:14.75 | להאמין |  |
| 9:14.76 | לדברי |  |
| 9:14.77 | הר"מ |  |
| 9:14.78 | נבאי |  |
| 9:14.79 | אשר |  |
| 9:14.80 | הרבה |  |
| 9:14.81 | דברים |  |
| 9:14.82 | לקלקל |  |
| 9:14.83 | עלינו |  |
| 9:14.84 | את |  |
| 9:14.85 | הכתובים |  |
| 9:14.86 | ודברי |  |
| 9:14.87 | רז"ל |  |
| 9:14.88 | המדברים |  |
| 9:14.89 | בפה |  |
| 9:14.90 | מלא |  |
| 9:14.91 | ובלשון |  |
| 9:14.92 | צח |  |
| 9:14.93 | שאין |  |
| 9:14.94 | הקב"ה |  |
| 9:14.95 | צריך |  |
| 9:14.96 | לבריותיו |  |
| 9:14.97 | להשלים |  |
| 9:14.98 | חסרונו |  |
| 9:14.99 | ח"ו | . |
| 9:14.100 | והרחיב |  |
| 9:14.101 | את |  |
| 9:14.102 | פיו |  |
| 9:14.103 | לדבר |  |
| 9:14.104 | גדולות |  |
| 9:14.105 | וגבוהות | , |
| 9:14.106 | בראיות |  |
| 9:14.107 | צנומות |  |
| 9:14.108 | דקות |  |
| 9:14.109 | ויצא |  |
| 9:14.110 | עתק |  |
| 9:14.111 | מפיו | , |
| 9:14.112 | על |  |
| 9:14.113 | הריב"ש |  |
| 9:14.114 | ז"ל | , |
| 9:14.115 | ושיבש |  |
| 9:14.116 | עלינו |  |
| 9:14.117 | פירוש |  |
| 9:14.118 | הכתובים |  |
| 9:14.119 | ודברי |  |
| 9:14.120 | רז"ל |  |
| 9:14.121 | ברוב |  |
| 9:14.122 | דברים |  |
| 9:14.123 | אשר |  |
| 9:14.124 | כל |  |
| 9:14.125 | משכיל |  |
| 9:14.126 | ידע |  |
| 9:14.127 | ויכיר |  |
| 9:14.128 | שאין |  |
| 9:14.129 | בהם |  |
| 9:14.130 | ממש | , |
| 9:14.131 | והאמת |  |
| 9:14.132 | עד |  |
| 9:14.133 | לעצמו |  |
| 9:14.134 | שכל |  |
| 9:14.135 | עבודותינו |  |
| 9:14.136 | ומעשה |  |
| 9:14.137 | המצות |  |
| 9:14.138 | שצונו |  |
| 9:14.139 | הקב"ה |  |
| 9:14.140 | בתורה |  |
| 9:14.141 | אינם |  |
| 9:14.142 | רק |  |
| 9:14.143 | לצרף |  |
| 9:14.144 | וללבן |  |
| 9:14.145 | אותנו | , |
| 9:14.146 | ונהיה |  |
| 9:14.147 | טהורים |  |
| 9:14.148 | ונקיים |  |
| 9:14.149 | במדותינו | ! |
| 9:14.150 | ולהרחיק |  |
| 9:14.151 | ממנו |  |
| 9:14.152 | הדעות |  |
| 9:14.153 | והאמונות |  |
| 9:14.154 | הרעות |  |
| 9:14.155 | ונהיה |  |
| 9:14.156 | קרובים |  |
| 9:14.157 | אליו |  |
| 9:14.158 | יתברך |  |
| 9:14.159 | כמאמר |  |
| 9:14.160 | רז"ל |  |
| 9:14.161 | לא |  |
| 9:14.162 | ניתנו |  |
| 9:14.163 | המצות |  |
| 9:14.164 | אלא |  |
| 9:14.165 | לצרוף |  |
| 9:14.166 | בהן |  |
| 9:14.167 | את |  |
| 9:14.168 | הבריות | , |
| 9:14.169 | שלא |  |
| 9:14.170 | יאמר |  |
| 9:14.171 | אדם |  |
| 9:14.172 | אי |  |
| 9:14.173 | אפשר |  |
| 9:14.174 | וכו | '. |
| 9:14.175 | וכן |  |
| 9:14.176 | אמרו |  |
| 9:14.177 | ז"ל |  |
| 9:14.178 | וכי |  |
| 9:14.179 | מה |  |
| 9:14.180 | איכפת |  |
| 9:14.181 | לו |  |
| 9:14.182 | להקב"ה |  |
| 9:14.183 | בין |  |
| 9:14.184 | שוחט |  |
| 9:14.185 | מן |  |
| 9:14.186 | הצואר |  |
| 9:14.187 | לשוחט |  |
| 9:14.188 | מן |  |
| 9:14.189 | העורף | , |
| 9:14.190 | הוי |  |
| 9:14.191 | אומר |  |
| 9:14.192 | לא |  |
| 9:14.193 | ניתנו |  |
| 9:14.194 | המצות |  |
| 9:14.195 | אלא |  |
| 9:14.196 | לצרוף |  |
| 9:14.197 | בהן |  |
| 9:14.198 | את |  |
| 9:14.199 | הבריות |  |
| 9:14.200 | שנאמר |  |
| 9:14.201 | אמרת |  |
| 9:14.202 | ה | ' |
| 9:14.203 | צרופה | ! |
| 9:14.204 | לא |  |
| 9:14.205 | לזווגו |  |
| 9:14.206 | הוא |  |
| 9:14.207 | צריך | , |
| 9:14.208 | ולא |  |
| 9:14.209 | לקשט |  |
| 9:14.210 | לו |  |
| 9:14.211 | את |  |
| 9:14.212 | כלתו | ! |
| 9:14.213 | כי |  |
| 9:14.214 | הוא |  |
| 9:14.215 | אחד |  |
| 9:14.216 | ומיוחד |  |
| 9:14.217 | מכל |  |
| 9:14.218 | האחדים |  |
| 9:14.219 | שבעולם | . |
| 9:14.220 | וכלל |  |
| 9:14.221 | הדברים |  |
| 9:14.222 | פשוט |  |
| 9:14.223 | הוא |  |
| 9:14.224 | שהתורה |  |
| 9:14.225 | ניתנה |  |
| 9:14.226 | לתועלת |  |
| 9:14.227 | מקבליה |  |
| 9:14.228 | ומקיימיה |  |
| 9:14.229 | לא |  |
| 9:14.230 | לתועלת |  |
| 9:14.231 | נותנה | , |
| 9:14.232 | שנאמר |  |
| 9:14.233 | לטוב |  |
| 9:14.234 | לנו |  |
| 9:14.235 | כל |  |
| 9:14.236 | הימים | . |


## Milchamot ha-Shem 9:15

Hebrew (Wikisource): קלג) ועל כל אלה הדעות הרעות החיצונות והזרות לנו לתורתינו הקדושה הזהיר הכתוב אותנו באמרו כי יסיתך איחך בן אמך או בנך או בתך או אשת חיקך או רעך אשר כנפשך בסתר לאמר נלכה ונעבדה אלהים אחרים אשר לא ידעת אתה ואבתיך, מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים ממך בסתר! אם יאמר בא ואגלה לך דברים נסתרים וסודות עמוקים לאמר, אלילים יש בשמים ממעל בכמה פרצופים זה למעלה מזה, וכך היא אוכלת וכך היא שותה, שמושכת לה מוחין למזונה וחיותה תמיד מאצל העתיק, ועל ידי מוחין שמושך ויונק פרצוף התחתון מהפרצוף העליון על ידי תפלותינו ובקשותינו ממנו הוא משפיע לנו שפע והארה, ובכך היא מטיבה לנו! וכך היא מריעה למי שלא יעבוד אותה, שכן צוו עליה אבא ואימא דכלא יפלחון ליה לזעיר אנפין הוא ואשתו, ומי שלא יעבוד רק לאלהים עליון ענוש יענש, שכך רצון המלך העליון שנתפלל ונקרא לאלוה קצר אפים (זעור) בצירוף הפרצופים שלמעלה ממנו כנזכר לעיל כמה פעמים בשם ספר הברית וכס"א ונח"י ויושר לבב! לא תאבה לו ולא תשמע אליו וכו'! ובסנהדרין דף ס"א אמרינן לא תאבה לו ולא תשמע אליו, הא אבה ושמע חייב, ואביי מפליג בה בין ניסת מפי עצמו לניסת מפי אחרים, דניסת מפי עצמו מימליך, אבל ניסת מפי אחרים גריד בתריהו, ורבא אמר אידי ואידי בניסת מפי אחרים, הא אמר ליה כך אוכלת, כך שותה, כך מטיבה כך מריעה, הא דלא אמר ליה כך אוכלת וכך שותה וכו'. אמר רבא מנא אמינא לה דכתיב מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים ממך, מה לי קרובים מה לי רחוקים אלא הכי קאמר, מטיבותן של קרובים אצה למד טיבותן של רחוקים ע"כ, ופירש רש"י ז"ל דכתיב בפרשת מסית מאלהי העמים הקרובים אליך או הרחוקים ממך, מה לי קרובים מה לי רחוקים. ולמה הזכיר קרוב ורחוק אלא מפני שדרכו של מסית לדבר על ע"ז רחוקה שאינו מכיר בה, ומוציא מפיו השקר ואומר כך אוכלת כך שותה, לפיכך אני אומר לך הסתכל בע"ז הקרובות אליך ותראה שאין בהם ממש, ומהן אתה למד טיבן של רחוקים ע"כ. ועיין לעיל סי' ט"ל.

| id | word | punctuation after |
|---|---|---|
| 9:15.1 | קלג | ) |
| 9:15.2 | ועל |  |
| 9:15.3 | כל |  |
| 9:15.4 | אלה |  |
| 9:15.5 | הדעות |  |
| 9:15.6 | הרעות |  |
| 9:15.7 | החיצונות |  |
| 9:15.8 | והזרות |  |
| 9:15.9 | לנו |  |
| 9:15.10 | לתורתינו |  |
| 9:15.11 | הקדושה |  |
| 9:15.12 | הזהיר |  |
| 9:15.13 | הכתוב |  |
| 9:15.14 | אותנו |  |
| 9:15.15 | באמרו |  |
| 9:15.16 | כי |  |
| 9:15.17 | יסיתך |  |
| 9:15.18 | איחך |  |
| 9:15.19 | בן |  |
| 9:15.20 | אמך |  |
| 9:15.21 | או |  |
| 9:15.22 | בנך |  |
| 9:15.23 | או |  |
| 9:15.24 | בתך |  |
| 9:15.25 | או |  |
| 9:15.26 | אשת |  |
| 9:15.27 | חיקך |  |
| 9:15.28 | או |  |
| 9:15.29 | רעך |  |
| 9:15.30 | אשר |  |
| 9:15.31 | כנפשך |  |
| 9:15.32 | בסתר |  |
| 9:15.33 | לאמר |  |
| 9:15.34 | נלכה |  |
| 9:15.35 | ונעבדה |  |
| 9:15.36 | אלהים |  |
| 9:15.37 | אחרים |  |
| 9:15.38 | אשר |  |
| 9:15.39 | לא |  |
| 9:15.40 | ידעת |  |
| 9:15.41 | אתה |  |
| 9:15.42 | ואבתיך | , |
| 9:15.43 | מאלהי |  |
| 9:15.44 | העמים |  |
| 9:15.45 | אשר |  |
| 9:15.46 | סביבותיכם |  |
| 9:15.47 | הקרובים |  |
| 9:15.48 | אליך |  |
| 9:15.49 | או |  |
| 9:15.50 | הרחוקים |  |
| 9:15.51 | ממך |  |
| 9:15.52 | בסתר | ! |
| 9:15.53 | אם |  |
| 9:15.54 | יאמר |  |
| 9:15.55 | בא |  |
| 9:15.56 | ואגלה |  |
| 9:15.57 | לך |  |
| 9:15.58 | דברים |  |
| 9:15.59 | נסתרים |  |
| 9:15.60 | וסודות |  |
| 9:15.61 | עמוקים |  |
| 9:15.62 | לאמר | , |
| 9:15.63 | אלילים |  |
| 9:15.64 | יש |  |
| 9:15.65 | בשמים |  |
| 9:15.66 | ממעל |  |
| 9:15.67 | בכמה |  |
| 9:15.68 | פרצופים |  |
| 9:15.69 | זה |  |
| 9:15.70 | למעלה |  |
| 9:15.71 | מזה | , |
| 9:15.72 | וכך |  |
| 9:15.73 | היא |  |
| 9:15.74 | אוכלת |  |
| 9:15.75 | וכך |  |
| 9:15.76 | היא |  |
| 9:15.77 | שותה | , |
| 9:15.78 | שמושכת |  |
| 9:15.79 | לה |  |
| 9:15.80 | מוחין |  |
| 9:15.81 | למזונה |  |
| 9:15.82 | וחיותה |  |
| 9:15.83 | תמיד |  |
| 9:15.84 | מאצל |  |
| 9:15.85 | העתיק | , |
| 9:15.86 | ועל |  |
| 9:15.87 | ידי |  |
| 9:15.88 | מוחין |  |
| 9:15.89 | שמושך |  |
| 9:15.90 | ויונק |  |
| 9:15.91 | פרצוף |  |
| 9:15.92 | התחתון |  |
| 9:15.93 | מהפרצוף |  |
| 9:15.94 | העליון |  |
| 9:15.95 | על |  |
| 9:15.96 | ידי |  |
| 9:15.97 | תפלותינו |  |
| 9:15.98 | ובקשותינו |  |
| 9:15.99 | ממנו |  |
| 9:15.100 | הוא |  |
| 9:15.101 | משפיע |  |
| 9:15.102 | לנו |  |
| 9:15.103 | שפע |  |
| 9:15.104 | והארה | , |
| 9:15.105 | ובכך |  |
| 9:15.106 | היא |  |
| 9:15.107 | מטיבה |  |
| 9:15.108 | לנו | ! |
| 9:15.109 | וכך |  |
| 9:15.110 | היא |  |
| 9:15.111 | מריעה |  |
| 9:15.112 | למי |  |
| 9:15.113 | שלא |  |
| 9:15.114 | יעבוד |  |
| 9:15.115 | אותה | , |
| 9:15.116 | שכן |  |
| 9:15.117 | צוו |  |
| 9:15.118 | עליה |  |
| 9:15.119 | אבא |  |
| 9:15.120 | ואימא |  |
| 9:15.121 | דכלא |  |
| 9:15.122 | יפלחון |  |
| 9:15.123 | ליה |  |
| 9:15.124 | לזעיר |  |
| 9:15.125 | אנפין |  |
| 9:15.126 | הוא |  |
| 9:15.127 | ואשתו | , |
| 9:15.128 | ומי |  |
| 9:15.129 | שלא |  |
| 9:15.130 | יעבוד |  |
| 9:15.131 | רק |  |
| 9:15.132 | לאלהים |  |
| 9:15.133 | עליון |  |
| 9:15.134 | ענוש |  |
| 9:15.135 | יענש | , |
| 9:15.136 | שכך |  |
| 9:15.137 | רצון |  |
| 9:15.138 | המלך |  |
| 9:15.139 | העליון |  |
| 9:15.140 | שנתפלל |  |
| 9:15.141 | ונקרא |  |
| 9:15.142 | לאלוה |  |
| 9:15.143 | קצר |  |
| 9:15.144 | אפים |  |
| 9:15.145 | זעור | ) |
| 9:15.146 | בצירוף |  |
| 9:15.147 | הפרצופים |  |
| 9:15.148 | שלמעלה |  |
| 9:15.149 | ממנו |  |
| 9:15.150 | כנזכר |  |
| 9:15.151 | לעיל |  |
| 9:15.152 | כמה |  |
| 9:15.153 | פעמים |  |
| 9:15.154 | בשם |  |
| 9:15.155 | ספר |  |
| 9:15.156 | הברית |  |
| 9:15.157 | וכס"א |  |
| 9:15.158 | ונח"י |  |
| 9:15.159 | ויושר |  |
| 9:15.160 | לבב | ! |
| 9:15.161 | לא |  |
| 9:15.162 | תאבה |  |
| 9:15.163 | לו |  |
| 9:15.164 | ולא |  |
| 9:15.165 | תשמע |  |
| 9:15.166 | אליו |  |
| 9:15.167 | וכו | '! |
| 9:15.168 | ובסנהדרין |  |
| 9:15.169 | דף |  |
| 9:15.170 | ס"א |  |
| 9:15.171 | אמרינן |  |
| 9:15.172 | לא |  |
| 9:15.173 | תאבה |  |
| 9:15.174 | לו |  |
| 9:15.175 | ולא |  |
| 9:15.176 | תשמע |  |
| 9:15.177 | אליו | , |
| 9:15.178 | הא |  |
| 9:15.179 | אבה |  |
| 9:15.180 | ושמע |  |
| 9:15.181 | חייב | , |
| 9:15.182 | ואביי |  |
| 9:15.183 | מפליג |  |
| 9:15.184 | בה |  |
| 9:15.185 | בין |  |
| 9:15.186 | ניסת |  |
| 9:15.187 | מפי |  |
| 9:15.188 | עצמו |  |
| 9:15.189 | לניסת |  |
| 9:15.190 | מפי |  |
| 9:15.191 | אחרים | , |
| 9:15.192 | דניסת |  |
| 9:15.193 | מפי |  |
| 9:15.194 | עצמו |  |
| 9:15.195 | מימליך | , |
| 9:15.196 | אבל |  |
| 9:15.197 | ניסת |  |
| 9:15.198 | מפי |  |
| 9:15.199 | אחרים |  |
| 9:15.200 | גריד |  |
| 9:15.201 | בתריהו | , |
| 9:15.202 | ורבא |  |
| 9:15.203 | אמר |  |
| 9:15.204 | אידי |  |
| 9:15.205 | ואידי |  |
| 9:15.206 | בניסת |  |
| 9:15.207 | מפי |  |
| 9:15.208 | אחרים | , |
| 9:15.209 | הא |  |
| 9:15.210 | אמר |  |
| 9:15.211 | ליה |  |
| 9:15.212 | כך |  |
| 9:15.213 | אוכלת | , |
| 9:15.214 | כך |  |
| 9:15.215 | שותה | , |
| 9:15.216 | כך |  |
| 9:15.217 | מטיבה |  |
| 9:15.218 | כך |  |
| 9:15.219 | מריעה | , |
| 9:15.220 | הא |  |
| 9:15.221 | דלא |  |
| 9:15.222 | אמר |  |
| 9:15.223 | ליה |  |
| 9:15.224 | כך |  |
| 9:15.225 | אוכלת |  |
| 9:15.226 | וכך |  |
| 9:15.227 | שותה |  |
| 9:15.228 | וכו | '. |
| 9:15.229 | אמר |  |
| 9:15.230 | רבא |  |
| 9:15.231 | מנא |  |
| 9:15.232 | אמינא |  |
| 9:15.233 | לה |  |
| 9:15.234 | דכתיב |  |
| 9:15.235 | מאלהי |  |
| 9:15.236 | העמים |  |
| 9:15.237 | אשר |  |
| 9:15.238 | סביבותיכם |  |
| 9:15.239 | הקרובים |  |
| 9:15.240 | אליך |  |
| 9:15.241 | או |  |
| 9:15.242 | הרחוקים |  |
| 9:15.243 | ממך | , |
| 9:15.244 | מה |  |
| 9:15.245 | לי |  |
| 9:15.246 | קרובים |  |
| 9:15.247 | מה |  |
| 9:15.248 | לי |  |
| 9:15.249 | רחוקים |  |
| 9:15.250 | אלא |  |
| 9:15.251 | הכי |  |
| 9:15.252 | קאמר | , |
| 9:15.253 | מטיבותן |  |
| 9:15.254 | של |  |
| 9:15.255 | קרובים |  |
| 9:15.256 | אצה |  |
| 9:15.257 | למד |  |
| 9:15.258 | טיבותן |  |
| 9:15.259 | של |  |
| 9:15.260 | רחוקים |  |
| 9:15.261 | ע"כ | , |
| 9:15.262 | ופירש |  |
| 9:15.263 | רש"י |  |
| 9:15.264 | ז"ל |  |
| 9:15.265 | דכתיב |  |
| 9:15.266 | בפרשת |  |
| 9:15.267 | מסית |  |
| 9:15.268 | מאלהי |  |
| 9:15.269 | העמים |  |
| 9:15.270 | הקרובים |  |
| 9:15.271 | אליך |  |
| 9:15.272 | או |  |
| 9:15.273 | הרחוקים |  |
| 9:15.274 | ממך | , |
| 9:15.275 | מה |  |
| 9:15.276 | לי |  |
| 9:15.277 | קרובים |  |
| 9:15.278 | מה |  |
| 9:15.279 | לי |  |
| 9:15.280 | רחוקים | . |
| 9:15.281 | ולמה |  |
| 9:15.282 | הזכיר |  |
| 9:15.283 | קרוב |  |
| 9:15.284 | ורחוק |  |
| 9:15.285 | אלא |  |
| 9:15.286 | מפני |  |
| 9:15.287 | שדרכו |  |
| 9:15.288 | של |  |
| 9:15.289 | מסית |  |
| 9:15.290 | לדבר |  |
| 9:15.291 | על |  |
| 9:15.292 | ע"ז |  |
| 9:15.293 | רחוקה |  |
| 9:15.294 | שאינו |  |
| 9:15.295 | מכיר |  |
| 9:15.296 | בה | , |
| 9:15.297 | ומוציא |  |
| 9:15.298 | מפיו |  |
| 9:15.299 | השקר |  |
| 9:15.300 | ואומר |  |
| 9:15.301 | כך |  |
| 9:15.302 | אוכלת |  |
| 9:15.303 | כך |  |
| 9:15.304 | שותה | , |
| 9:15.305 | לפיכך |  |
| 9:15.306 | אני |  |
| 9:15.307 | אומר |  |
| 9:15.308 | לך |  |
| 9:15.309 | הסתכל |  |
| 9:15.310 | בע"ז |  |
| 9:15.311 | הקרובות |  |
| 9:15.312 | אליך |  |
| 9:15.313 | ותראה |  |
| 9:15.314 | שאין |  |
| 9:15.315 | בהם |  |
| 9:15.316 | ממש | , |
| 9:15.317 | ומהן |  |
| 9:15.318 | אתה |  |
| 9:15.319 | למד |  |
| 9:15.320 | טיבן |  |
| 9:15.321 | של |  |
| 9:15.322 | רחוקים |  |
| 9:15.323 | ע"כ | . |
| 9:15.324 | ועיין |  |
| 9:15.325 | לעיל |  |
| 9:15.326 | סי | ' |
| 9:15.327 | ט"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

