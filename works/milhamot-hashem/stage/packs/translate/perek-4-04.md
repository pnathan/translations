# Translation pack: milhamot-hashem perek-4-04 (Milchamot ha-Shem 4:7–4:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-04`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
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
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
## Terms

- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- אל **אחר** → "another god".
- **מין / מינות** → "heretic / heresy".
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
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- מי **אתי** (ketiv): translated, the qere noted.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ח"א / ח"ב** → "vol. 1" / "vol. 2".
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


## Milchamot ha-Shem 4:7

Hebrew (Wikisource): כד) עוד איתא בזהר בראשית ט"ו בראשית ברא אלהים את השמים, בהאי בראשית ברא ההוא סתימא דלא אתידע להאי היכלא וכו' ופי' הרב מקדש מלך בהאי ראשית שהוא אבא, ברא ההוא סתימא, הוא אריך דלא אתידע, הוא עתיק, להיכלא דא היא אימא הנקראת אלהים ע"כ. ועוד פי' שם פי' אחר בראשית ברא אלהים את השמים ואת הארץ, ר"ל בכח אבא ברא אלהים שהיא בינה את השמים ואת הארץ שהוא זעיר אנפין ונוקביה ע"כ. הרי מבואר מזה דבין לפירוש קמא בין לפי בתרא הוא קורא לאבא בראשית, אלא דלפירוש קמא קראו אריך ועתיק את אלהים שהיא בינה דהיינו אמא. ולפירוש בתרא בכח אבא ברא אלהים שהיא בינה את השמים ואת הארץ שהם זעיר ונוקביה (מקדש מלך ד"ה אי בהאי). ובפרשת בא דף ל"ז (ח"ב לז, א) את השמים דא ימינא דקב"ה ואת הארץ דא שמאלא, דסטא קב"ה ימיניה וברא את שמיא, וסטא שמאליה וברא ית ארעה. ופי' בגליון הרח"ו החסדים בונים זעיר, והגבורות בונים הנוקבא, הדא הוא דכתיב אף ידי יסדה ארץ וימיני טפחה שמים קורא אני אליהם יעמדו יחדיו, מהו יעמדו יחדיו, סלקא דעתך שמיא וארעא. לאו הכי אלא ימינא ושמאלא דאיהו את ואת וכו'. ופי' המקדש מלך, אף ידי יסדה ארץ, הבינה אומרת כן, סלקא דעתך שמיא וארעא, פי' דאי אפשר לאמר עליהם יעמדו יחדו דמשמע תמיד הם עומדים יחדיו, וא"א לומר כן, שהרי הנוקבא פעמים נפרדת ממנו בגלות ע"כ, ועיין לקטן ס"י כ"ז היפך זה דלא מתפרשין לעלמין. הרואה יראה ויכיר ויבין איך הפילוסוף מחבר הזהר מעקם עלינו את נדרך בדרשת הכתובים שלא כרצון חכמים כמבואר להלן בסי' דלקמן.

| id | word | punctuation after |
|---|---|---|
| 4:7.1 | כד | ) |
| 4:7.2 | עוד |  |
| 4:7.3 | איתא |  |
| 4:7.4 | בזהר |  |
| 4:7.5 | בראשית |  |
| 4:7.6 | ט"ו |  |
| 4:7.7 | בראשית |  |
| 4:7.8 | ברא |  |
| 4:7.9 | אלהים |  |
| 4:7.10 | את |  |
| 4:7.11 | השמים | , |
| 4:7.12 | בהאי |  |
| 4:7.13 | בראשית |  |
| 4:7.14 | ברא |  |
| 4:7.15 | ההוא |  |
| 4:7.16 | סתימא |  |
| 4:7.17 | דלא |  |
| 4:7.18 | אתידע |  |
| 4:7.19 | להאי |  |
| 4:7.20 | היכלא |  |
| 4:7.21 | וכו | ' |
| 4:7.22 | ופי | ' |
| 4:7.23 | הרב |  |
| 4:7.24 | מקדש |  |
| 4:7.25 | מלך |  |
| 4:7.26 | בהאי |  |
| 4:7.27 | ראשית |  |
| 4:7.28 | שהוא |  |
| 4:7.29 | אבא | , |
| 4:7.30 | ברא |  |
| 4:7.31 | ההוא |  |
| 4:7.32 | סתימא | , |
| 4:7.33 | הוא |  |
| 4:7.34 | אריך |  |
| 4:7.35 | דלא |  |
| 4:7.36 | אתידע | , |
| 4:7.37 | הוא |  |
| 4:7.38 | עתיק | , |
| 4:7.39 | להיכלא |  |
| 4:7.40 | דא |  |
| 4:7.41 | היא |  |
| 4:7.42 | אימא |  |
| 4:7.43 | הנקראת |  |
| 4:7.44 | אלהים |  |
| 4:7.45 | ע"כ | . |
| 4:7.46 | ועוד |  |
| 4:7.47 | פי | ' |
| 4:7.48 | שם |  |
| 4:7.49 | פי | ' |
| 4:7.50 | אחר |  |
| 4:7.51 | בראשית |  |
| 4:7.52 | ברא |  |
| 4:7.53 | אלהים |  |
| 4:7.54 | את |  |
| 4:7.55 | השמים |  |
| 4:7.56 | ואת |  |
| 4:7.57 | הארץ | , |
| 4:7.58 | ר"ל |  |
| 4:7.59 | בכח |  |
| 4:7.60 | אבא |  |
| 4:7.61 | ברא |  |
| 4:7.62 | אלהים |  |
| 4:7.63 | שהיא |  |
| 4:7.64 | בינה |  |
| 4:7.65 | את |  |
| 4:7.66 | השמים |  |
| 4:7.67 | ואת |  |
| 4:7.68 | הארץ |  |
| 4:7.69 | שהוא |  |
| 4:7.70 | זעיר |  |
| 4:7.71 | אנפין |  |
| 4:7.72 | ונוקביה |  |
| 4:7.73 | ע"כ | . |
| 4:7.74 | הרי |  |
| 4:7.75 | מבואר |  |
| 4:7.76 | מזה |  |
| 4:7.77 | דבין |  |
| 4:7.78 | לפירוש |  |
| 4:7.79 | קמא |  |
| 4:7.80 | בין |  |
| 4:7.81 | לפי |  |
| 4:7.82 | בתרא |  |
| 4:7.83 | הוא |  |
| 4:7.84 | קורא |  |
| 4:7.85 | לאבא |  |
| 4:7.86 | בראשית | , |
| 4:7.87 | אלא |  |
| 4:7.88 | דלפירוש |  |
| 4:7.89 | קמא |  |
| 4:7.90 | קראו |  |
| 4:7.91 | אריך |  |
| 4:7.92 | ועתיק |  |
| 4:7.93 | את |  |
| 4:7.94 | אלהים |  |
| 4:7.95 | שהיא |  |
| 4:7.96 | בינה |  |
| 4:7.97 | דהיינו |  |
| 4:7.98 | אמא | . |
| 4:7.99 | ולפירוש |  |
| 4:7.100 | בתרא |  |
| 4:7.101 | בכח |  |
| 4:7.102 | אבא |  |
| 4:7.103 | ברא |  |
| 4:7.104 | אלהים |  |
| 4:7.105 | שהיא |  |
| 4:7.106 | בינה |  |
| 4:7.107 | את |  |
| 4:7.108 | השמים |  |
| 4:7.109 | ואת |  |
| 4:7.110 | הארץ |  |
| 4:7.111 | שהם |  |
| 4:7.112 | זעיר |  |
| 4:7.113 | ונוקביה |  |
| 4:7.114 | מקדש |  |
| 4:7.115 | מלך |  |
| 4:7.116 | ד"ה |  |
| 4:7.117 | אי |  |
| 4:7.118 | בהאי | ). |
| 4:7.119 | ובפרשת |  |
| 4:7.120 | בא |  |
| 4:7.121 | דף |  |
| 4:7.122 | ל"ז |  |
| 4:7.123 | ח"ב |  |
| 4:7.124 | לז | , |
| 4:7.125 | א | ) |
| 4:7.126 | את |  |
| 4:7.127 | השמים |  |
| 4:7.128 | דא |  |
| 4:7.129 | ימינא |  |
| 4:7.130 | דקב"ה |  |
| 4:7.131 | ואת |  |
| 4:7.132 | הארץ |  |
| 4:7.133 | דא |  |
| 4:7.134 | שמאלא | , |
| 4:7.135 | דסטא |  |
| 4:7.136 | קב"ה |  |
| 4:7.137 | ימיניה |  |
| 4:7.138 | וברא |  |
| 4:7.139 | את |  |
| 4:7.140 | שמיא | , |
| 4:7.141 | וסטא |  |
| 4:7.142 | שמאליה |  |
| 4:7.143 | וברא |  |
| 4:7.144 | ית |  |
| 4:7.145 | ארעה | . |
| 4:7.146 | ופי | ' |
| 4:7.147 | בגליון |  |
| 4:7.148 | הרח"ו |  |
| 4:7.149 | החסדים |  |
| 4:7.150 | בונים |  |
| 4:7.151 | זעיר | , |
| 4:7.152 | והגבורות |  |
| 4:7.153 | בונים |  |
| 4:7.154 | הנוקבא | , |
| 4:7.155 | הדא |  |
| 4:7.156 | הוא |  |
| 4:7.157 | דכתיב |  |
| 4:7.158 | אף |  |
| 4:7.159 | ידי |  |
| 4:7.160 | יסדה |  |
| 4:7.161 | ארץ |  |
| 4:7.162 | וימיני |  |
| 4:7.163 | טפחה |  |
| 4:7.164 | שמים |  |
| 4:7.165 | קורא |  |
| 4:7.166 | אני |  |
| 4:7.167 | אליהם |  |
| 4:7.168 | יעמדו |  |
| 4:7.169 | יחדיו | , |
| 4:7.170 | מהו |  |
| 4:7.171 | יעמדו |  |
| 4:7.172 | יחדיו | , |
| 4:7.173 | סלקא |  |
| 4:7.174 | דעתך |  |
| 4:7.175 | שמיא |  |
| 4:7.176 | וארעא | . |
| 4:7.177 | לאו |  |
| 4:7.178 | הכי |  |
| 4:7.179 | אלא |  |
| 4:7.180 | ימינא |  |
| 4:7.181 | ושמאלא |  |
| 4:7.182 | דאיהו |  |
| 4:7.183 | את |  |
| 4:7.184 | ואת |  |
| 4:7.185 | וכו | '. |
| 4:7.186 | ופי | ' |
| 4:7.187 | המקדש |  |
| 4:7.188 | מלך | , |
| 4:7.189 | אף |  |
| 4:7.190 | ידי |  |
| 4:7.191 | יסדה |  |
| 4:7.192 | ארץ | , |
| 4:7.193 | הבינה |  |
| 4:7.194 | אומרת |  |
| 4:7.195 | כן | , |
| 4:7.196 | סלקא |  |
| 4:7.197 | דעתך |  |
| 4:7.198 | שמיא |  |
| 4:7.199 | וארעא | , |
| 4:7.200 | פי | ' |
| 4:7.201 | דאי |  |
| 4:7.202 | אפשר |  |
| 4:7.203 | לאמר |  |
| 4:7.204 | עליהם |  |
| 4:7.205 | יעמדו |  |
| 4:7.206 | יחדו |  |
| 4:7.207 | דמשמע |  |
| 4:7.208 | תמיד |  |
| 4:7.209 | הם |  |
| 4:7.210 | עומדים |  |
| 4:7.211 | יחדיו | , |
| 4:7.212 | וא"א |  |
| 4:7.213 | לומר |  |
| 4:7.214 | כן | , |
| 4:7.215 | שהרי |  |
| 4:7.216 | הנוקבא |  |
| 4:7.217 | פעמים |  |
| 4:7.218 | נפרדת |  |
| 4:7.219 | ממנו |  |
| 4:7.220 | בגלות |  |
| 4:7.221 | ע"כ | , |
| 4:7.222 | ועיין |  |
| 4:7.223 | לקטן |  |
| 4:7.224 | ס"י |  |
| 4:7.225 | כ"ז |  |
| 4:7.226 | היפך |  |
| 4:7.227 | זה |  |
| 4:7.228 | דלא |  |
| 4:7.229 | מתפרשין |  |
| 4:7.230 | לעלמין | . |
| 4:7.231 | הרואה |  |
| 4:7.232 | יראה |  |
| 4:7.233 | ויכיר |  |
| 4:7.234 | ויבין |  |
| 4:7.235 | איך |  |
| 4:7.236 | הפילוסוף |  |
| 4:7.237 | מחבר |  |
| 4:7.238 | הזהר |  |
| 4:7.239 | מעקם |  |
| 4:7.240 | עלינו |  |
| 4:7.241 | את |  |
| 4:7.242 | נדרך |  |
| 4:7.243 | בדרשת |  |
| 4:7.244 | הכתובים |  |
| 4:7.245 | שלא |  |
| 4:7.246 | כרצון |  |
| 4:7.247 | חכמים |  |
| 4:7.248 | כמבואר |  |
| 4:7.249 | להלן |  |
| 4:7.250 | בסי | ' |
| 4:7.251 | דלקמן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

