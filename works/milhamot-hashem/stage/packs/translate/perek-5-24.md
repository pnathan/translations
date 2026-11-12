# Translation pack: milhamot-hashem perek-5-24 (Milchamot ha-Shem 5:33–5:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-24.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-24`.

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
- **רז"ל** → "our Rabbis of blessed memory".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
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
- **מערב** beside Ashkenaz, Sefarad, Yemen → "the Maghreb" (note: "the West").
## Terms

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
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- לית מחשבה **תפיסה** ביה → "no thought grasps Him".
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- קורא **נעים / נעלם** → "pleasant reader".
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
- מי **אתי** (ketiv): translated, the qere noted.
- **כדנה / אלהיא / יאבדו** (Jer 10:11) → "Thus shall you say to them: the gods
  that did not make the heavens and the earth shall perish from the earth and
  from under these heavens."
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


## Milchamot ha-Shem 5:33

Hebrew (Wikisource): נט) ותו גרסינן במנחות פרק הרי עלי עשרון (מנחות דף ק"י) (דף קי.) אמר ר"א בר יצחק ואמרי לה אמר רב יהודה אמר רב מצור ועד קרתגני מכירין את ישראל ואת אביהם שבשמים, מצור כלפי מערב. ומקרתגני כלפי מזרח אין מכירין לא את ישראל ולא את אביהם שבשמים. אתיביה רב שימי רב חייא לרב. כי ממזרח שמש ועד מבואו גדול שמי בגוים ובכל מקום וכו' אמר ליה שימי את? דקרו ליה (לאלהי ישראל) אלהא דאלהיא ע"כ ופי' רש"י מכירים את ישראל מכירים בכבודן הראשון ואין מכירים את עולם ואת אביהם שבשמים שמאמינים בהקב"ה ואינם עובדים ע"ז. וכתב הר' מהרש"א נקט מתחלה ישראל דעל ידי ישראל למטה נתפרסם אלהותו ית' ויכלתו בעולם התחתון ושהוא אלהי ישראל ומשגיח על הארץ דהיינו אביהם שבשמים אבל מצור כלפי מערב כיון שאין מכירים את ישראל. אין מכירים את אביהם שבשמים שהוא משגיח בעולם השפל. אבל הוא אלהא דאלהא כדמסיק שהוא אלהי האלהים. ולאותן אלוהות שתחתיו נתן כח וממשלה למטה ואינו משגיח בעולם השפל התחתון עכ"ל. הט לבך קורא נעים לתושיה ולתבונה בדברי רז"ל בעלי התלמוד ותפקחנה עיניך ונהרת ונוכחת לדעת כי דעת המקובלים שאמרו שהאין סוף שהוא לעילא לעילא לית מחשבה תפיסה ביה ואין להתפלל ולקרא אליו רק לזעיר אנפין דשולטנו יהיב ליה על כלא. ושהקורא ומתפלל לאין סוף שהוא אלהים עליון לא יענה! היא בעצמה אמונת אוה"ע מצור כלפי מערב ומקרתגני כלפי מזרח האומרים שאבינו שבשמים המגביהי לשבת איננו משגיח בתחתונים. אבל הוא אלהי האלהים שהם חמשה פרצופים ושגם פרצופים העליונים עצמן נתנו הממשלה והכח לאלוה קצר אפים (זעיר אנפין) דאעטרו ליה ויהבו ליה שולטנו על כלא (זהר בלק דף קצ"א) כי רק לזעיר אנפין נתכנו עלילות. אבל האלהים העליונים לא נתכנו להם עלילות (זהר האזינו דף רצ"ב) כמו שנת' לעיל מסי' כ"ט והלאה. הוא שהנביא ישעיה מוכיח את ישראל במאמרו ית' לאמר איזה ספר וכו' הקצור קצרה ידי מפדות ואם אין בי כח להציל וכו' עד שאצטרך להעזר בפרצופים שהם למטה ממני. או שאמנה אותם במקומי? כי הוא ית' הסבה הראשונה לא צוה לעבוד לשום נברא כלל, ואינו כח בגוף כדי לעבוד לנשמתו כמו שחשבו המקובלים החדשים.

| id | word | punctuation after |
|---|---|---|
| 5:33.1 | נט | ) |
| 5:33.2 | ותו |  |
| 5:33.3 | גרסינן |  |
| 5:33.4 | במנחות |  |
| 5:33.5 | פרק |  |
| 5:33.6 | הרי |  |
| 5:33.7 | עלי |  |
| 5:33.8 | עשרון |  |
| 5:33.9 | מנחות |  |
| 5:33.10 | דף |  |
| 5:33.11 | ק"י | ) |
| 5:33.12 | דף |  |
| 5:33.13 | קי | .) |
| 5:33.14 | אמר |  |
| 5:33.15 | ר"א |  |
| 5:33.16 | בר |  |
| 5:33.17 | יצחק |  |
| 5:33.18 | ואמרי |  |
| 5:33.19 | לה |  |
| 5:33.20 | אמר |  |
| 5:33.21 | רב |  |
| 5:33.22 | יהודה |  |
| 5:33.23 | אמר |  |
| 5:33.24 | רב |  |
| 5:33.25 | מצור |  |
| 5:33.26 | ועד |  |
| 5:33.27 | קרתגני |  |
| 5:33.28 | מכירין |  |
| 5:33.29 | את |  |
| 5:33.30 | ישראל |  |
| 5:33.31 | ואת |  |
| 5:33.32 | אביהם |  |
| 5:33.33 | שבשמים | , |
| 5:33.34 | מצור |  |
| 5:33.35 | כלפי |  |
| 5:33.36 | מערב | . |
| 5:33.37 | ומקרתגני |  |
| 5:33.38 | כלפי |  |
| 5:33.39 | מזרח |  |
| 5:33.40 | אין |  |
| 5:33.41 | מכירין |  |
| 5:33.42 | לא |  |
| 5:33.43 | את |  |
| 5:33.44 | ישראל |  |
| 5:33.45 | ולא |  |
| 5:33.46 | את |  |
| 5:33.47 | אביהם |  |
| 5:33.48 | שבשמים | . |
| 5:33.49 | אתיביה |  |
| 5:33.50 | רב |  |
| 5:33.51 | שימי |  |
| 5:33.52 | רב |  |
| 5:33.53 | חייא |  |
| 5:33.54 | לרב | . |
| 5:33.55 | כי |  |
| 5:33.56 | ממזרח |  |
| 5:33.57 | שמש |  |
| 5:33.58 | ועד |  |
| 5:33.59 | מבואו |  |
| 5:33.60 | גדול |  |
| 5:33.61 | שמי |  |
| 5:33.62 | בגוים |  |
| 5:33.63 | ובכל |  |
| 5:33.64 | מקום |  |
| 5:33.65 | וכו | ' |
| 5:33.66 | אמר |  |
| 5:33.67 | ליה |  |
| 5:33.68 | שימי |  |
| 5:33.69 | את | ? |
| 5:33.70 | דקרו |  |
| 5:33.71 | ליה |  |
| 5:33.72 | לאלהי |  |
| 5:33.73 | ישראל | ) |
| 5:33.74 | אלהא |  |
| 5:33.75 | דאלהיא |  |
| 5:33.76 | ע"כ |  |
| 5:33.77 | ופי | ' |
| 5:33.78 | רש"י |  |
| 5:33.79 | מכירים |  |
| 5:33.80 | את |  |
| 5:33.81 | ישראל |  |
| 5:33.82 | מכירים |  |
| 5:33.83 | בכבודן |  |
| 5:33.84 | הראשון |  |
| 5:33.85 | ואין |  |
| 5:33.86 | מכירים |  |
| 5:33.87 | את |  |
| 5:33.88 | עולם |  |
| 5:33.89 | ואת |  |
| 5:33.90 | אביהם |  |
| 5:33.91 | שבשמים |  |
| 5:33.92 | שמאמינים |  |
| 5:33.93 | בהקב"ה |  |
| 5:33.94 | ואינם |  |
| 5:33.95 | עובדים |  |
| 5:33.96 | ע"ז | . |
| 5:33.97 | וכתב |  |
| 5:33.98 | הר | ' |
| 5:33.99 | מהרש"א |  |
| 5:33.100 | נקט |  |
| 5:33.101 | מתחלה |  |
| 5:33.102 | ישראל |  |
| 5:33.103 | דעל |  |
| 5:33.104 | ידי |  |
| 5:33.105 | ישראל |  |
| 5:33.106 | למטה |  |
| 5:33.107 | נתפרסם |  |
| 5:33.108 | אלהותו |  |
| 5:33.109 | ית | ' |
| 5:33.110 | ויכלתו |  |
| 5:33.111 | בעולם |  |
| 5:33.112 | התחתון |  |
| 5:33.113 | ושהוא |  |
| 5:33.114 | אלהי |  |
| 5:33.115 | ישראל |  |
| 5:33.116 | ומשגיח |  |
| 5:33.117 | על |  |
| 5:33.118 | הארץ |  |
| 5:33.119 | דהיינו |  |
| 5:33.120 | אביהם |  |
| 5:33.121 | שבשמים |  |
| 5:33.122 | אבל |  |
| 5:33.123 | מצור |  |
| 5:33.124 | כלפי |  |
| 5:33.125 | מערב |  |
| 5:33.126 | כיון |  |
| 5:33.127 | שאין |  |
| 5:33.128 | מכירים |  |
| 5:33.129 | את |  |
| 5:33.130 | ישראל | . |
| 5:33.131 | אין |  |
| 5:33.132 | מכירים |  |
| 5:33.133 | את |  |
| 5:33.134 | אביהם |  |
| 5:33.135 | שבשמים |  |
| 5:33.136 | שהוא |  |
| 5:33.137 | משגיח |  |
| 5:33.138 | בעולם |  |
| 5:33.139 | השפל | . |
| 5:33.140 | אבל |  |
| 5:33.141 | הוא |  |
| 5:33.142 | אלהא |  |
| 5:33.143 | דאלהא |  |
| 5:33.144 | כדמסיק |  |
| 5:33.145 | שהוא |  |
| 5:33.146 | אלהי |  |
| 5:33.147 | האלהים | . |
| 5:33.148 | ולאותן |  |
| 5:33.149 | אלוהות |  |
| 5:33.150 | שתחתיו |  |
| 5:33.151 | נתן |  |
| 5:33.152 | כח |  |
| 5:33.153 | וממשלה |  |
| 5:33.154 | למטה |  |
| 5:33.155 | ואינו |  |
| 5:33.156 | משגיח |  |
| 5:33.157 | בעולם |  |
| 5:33.158 | השפל |  |
| 5:33.159 | התחתון |  |
| 5:33.160 | עכ"ל | . |
| 5:33.161 | הט |  |
| 5:33.162 | לבך |  |
| 5:33.163 | קורא |  |
| 5:33.164 | נעים |  |
| 5:33.165 | לתושיה |  |
| 5:33.166 | ולתבונה |  |
| 5:33.167 | בדברי |  |
| 5:33.168 | רז"ל |  |
| 5:33.169 | בעלי |  |
| 5:33.170 | התלמוד |  |
| 5:33.171 | ותפקחנה |  |
| 5:33.172 | עיניך |  |
| 5:33.173 | ונהרת |  |
| 5:33.174 | ונוכחת |  |
| 5:33.175 | לדעת |  |
| 5:33.176 | כי |  |
| 5:33.177 | דעת |  |
| 5:33.178 | המקובלים |  |
| 5:33.179 | שאמרו |  |
| 5:33.180 | שהאין |  |
| 5:33.181 | סוף |  |
| 5:33.182 | שהוא |  |
| 5:33.183 | לעילא |  |
| 5:33.184 | לעילא |  |
| 5:33.185 | לית |  |
| 5:33.186 | מחשבה |  |
| 5:33.187 | תפיסה |  |
| 5:33.188 | ביה |  |
| 5:33.189 | ואין |  |
| 5:33.190 | להתפלל |  |
| 5:33.191 | ולקרא |  |
| 5:33.192 | אליו |  |
| 5:33.193 | רק |  |
| 5:33.194 | לזעיר |  |
| 5:33.195 | אנפין |  |
| 5:33.196 | דשולטנו |  |
| 5:33.197 | יהיב |  |
| 5:33.198 | ליה |  |
| 5:33.199 | על |  |
| 5:33.200 | כלא | . |
| 5:33.201 | ושהקורא |  |
| 5:33.202 | ומתפלל |  |
| 5:33.203 | לאין |  |
| 5:33.204 | סוף |  |
| 5:33.205 | שהוא |  |
| 5:33.206 | אלהים |  |
| 5:33.207 | עליון |  |
| 5:33.208 | לא |  |
| 5:33.209 | יענה | ! |
| 5:33.210 | היא |  |
| 5:33.211 | בעצמה |  |
| 5:33.212 | אמונת |  |
| 5:33.213 | אוה"ע |  |
| 5:33.214 | מצור |  |
| 5:33.215 | כלפי |  |
| 5:33.216 | מערב |  |
| 5:33.217 | ומקרתגני |  |
| 5:33.218 | כלפי |  |
| 5:33.219 | מזרח |  |
| 5:33.220 | האומרים |  |
| 5:33.221 | שאבינו |  |
| 5:33.222 | שבשמים |  |
| 5:33.223 | המגביהי |  |
| 5:33.224 | לשבת |  |
| 5:33.225 | איננו |  |
| 5:33.226 | משגיח |  |
| 5:33.227 | בתחתונים | . |
| 5:33.228 | אבל |  |
| 5:33.229 | הוא |  |
| 5:33.230 | אלהי |  |
| 5:33.231 | האלהים |  |
| 5:33.232 | שהם |  |
| 5:33.233 | חמשה |  |
| 5:33.234 | פרצופים |  |
| 5:33.235 | ושגם |  |
| 5:33.236 | פרצופים |  |
| 5:33.237 | העליונים |  |
| 5:33.238 | עצמן |  |
| 5:33.239 | נתנו |  |
| 5:33.240 | הממשלה |  |
| 5:33.241 | והכח |  |
| 5:33.242 | לאלוה |  |
| 5:33.243 | קצר |  |
| 5:33.244 | אפים |  |
| 5:33.245 | זעיר |  |
| 5:33.246 | אנפין | ) |
| 5:33.247 | דאעטרו |  |
| 5:33.248 | ליה |  |
| 5:33.249 | ויהבו |  |
| 5:33.250 | ליה |  |
| 5:33.251 | שולטנו |  |
| 5:33.252 | על |  |
| 5:33.253 | כלא |  |
| 5:33.254 | זהר |  |
| 5:33.255 | בלק |  |
| 5:33.256 | דף |  |
| 5:33.257 | קצ"א | ) |
| 5:33.258 | כי |  |
| 5:33.259 | רק |  |
| 5:33.260 | לזעיר |  |
| 5:33.261 | אנפין |  |
| 5:33.262 | נתכנו |  |
| 5:33.263 | עלילות | . |
| 5:33.264 | אבל |  |
| 5:33.265 | האלהים |  |
| 5:33.266 | העליונים |  |
| 5:33.267 | לא |  |
| 5:33.268 | נתכנו |  |
| 5:33.269 | להם |  |
| 5:33.270 | עלילות |  |
| 5:33.271 | זהר |  |
| 5:33.272 | האזינו |  |
| 5:33.273 | דף |  |
| 5:33.274 | רצ"ב | ) |
| 5:33.275 | כמו |  |
| 5:33.276 | שנת | ' |
| 5:33.277 | לעיל |  |
| 5:33.278 | מסי | ' |
| 5:33.279 | כ"ט |  |
| 5:33.280 | והלאה | . |
| 5:33.281 | הוא |  |
| 5:33.282 | שהנביא |  |
| 5:33.283 | ישעיה |  |
| 5:33.284 | מוכיח |  |
| 5:33.285 | את |  |
| 5:33.286 | ישראל |  |
| 5:33.287 | במאמרו |  |
| 5:33.288 | ית | ' |
| 5:33.289 | לאמר |  |
| 5:33.290 | איזה |  |
| 5:33.291 | ספר |  |
| 5:33.292 | וכו | ' |
| 5:33.293 | הקצור |  |
| 5:33.294 | קצרה |  |
| 5:33.295 | ידי |  |
| 5:33.296 | מפדות |  |
| 5:33.297 | ואם |  |
| 5:33.298 | אין |  |
| 5:33.299 | בי |  |
| 5:33.300 | כח |  |
| 5:33.301 | להציל |  |
| 5:33.302 | וכו | ' |
| 5:33.303 | עד |  |
| 5:33.304 | שאצטרך |  |
| 5:33.305 | להעזר |  |
| 5:33.306 | בפרצופים |  |
| 5:33.307 | שהם |  |
| 5:33.308 | למטה |  |
| 5:33.309 | ממני | . |
| 5:33.310 | או |  |
| 5:33.311 | שאמנה |  |
| 5:33.312 | אותם |  |
| 5:33.313 | במקומי | ? |
| 5:33.314 | כי |  |
| 5:33.315 | הוא |  |
| 5:33.316 | ית | ' |
| 5:33.317 | הסבה |  |
| 5:33.318 | הראשונה |  |
| 5:33.319 | לא |  |
| 5:33.320 | צוה |  |
| 5:33.321 | לעבוד |  |
| 5:33.322 | לשום |  |
| 5:33.323 | נברא |  |
| 5:33.324 | כלל | , |
| 5:33.325 | ואינו |  |
| 5:33.326 | כח |  |
| 5:33.327 | בגוף |  |
| 5:33.328 | כדי |  |
| 5:33.329 | לעבוד |  |
| 5:33.330 | לנשמתו |  |
| 5:33.331 | כמו |  |
| 5:33.332 | שחשבו |  |
| 5:33.333 | המקובלים |  |
| 5:33.334 | החדשים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

