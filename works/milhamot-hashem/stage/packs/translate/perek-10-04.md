# Translation pack: milhamot-hashem perek-10-04 (Milchamot ha-Shem 10:5–10:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-04`.

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
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **האר"י** → "Ari" (authority, transliterated; note).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
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

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
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
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- תורה **שבעל** פה → "the Oral Torah".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מרגל** לרגל → "from festival to festival", not "foot".
- **דרך ארץ** → "worldly occupation".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- **מילדי** נכרים → "from the children of strangers".
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
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
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


## Milchamot ha-Shem 10:5

Hebrew (Wikisource): קמג) ולו עיני רז"ל אנשי כנה"ג ראו את כל הנעשה בדת היהדות על פי הספר המזוייף מעבודת הכחות הנעלמים כי אז קרעו את בגדיהם ולבשו שק ואפר, וצעקו בחוצות ובשווקים וברחובות, צעקה גדולה ומרה כמרדכי בשושן הבירה, הזאת היא חקת התורה? להמציא בתורה רמזים לעבודת האלילים? ולימודי ידיעת האלילים האלה קראו להם סתרי תורה, מלשון סתירה שהם סותרים ועוקרים יסודותיה, והורסים פנותיה כמו שדרשו רז"ל (ילקוט פרשת שלח) אמר ר' יצחק דבר זה מסורת בידינו מאבותינו מרגלים על שם מעשיהם נקראו, ולא עלתה בידינו אלא אחד, סתור בן מיכאל סתור שסתר דבריו של הקב"ה, בן מיכאל, שעשה עצמו מך, אמר ר' יוחנן אף אנו נאמר נחבי בן ופסי, שהחביא דבריו של הקב"ה ופסע עד מדותיו של הקב"ה. כדרך הזאת עשה הפילסוף המסית מחבר הזהר והנמשכים אחריו לסתור ולקלקל את אמונותינו באל אחד מיוחד מכל שאר האחדים, ולפסוע על מדותיו של הקב"ה ולעשותן אלוהות במה שקדם, כי לא הספיק לו לחשוב באלהות עילה אחת, והוצרך להמציא ברעיונותיו עילות רבות שבה נוטל רשות מזה וזה נעזר במעשיהו על ידי אחר, כי חלש הוא ולא יוכל לעשות לבדו כמו שנתבאר לעיל, ובחר לו לעבוד את האליל קצר האפים אשר לפי דעתו לא נשתתף בשום מעשה מכל מעשה בראשית, כי הוא נברא ועדיין היה חסר תיקון ולא נשלם תיקוני עד ימי אברהם אבינו כמו שנתבאר לעיל בשם חיי שלום משם האר"י ולא די זה אלא שהוא מבזה ומהתל בשלש עשרה מדות שהתורה נדרשת בהם שקבלם משה בסיני, ועל פ' חי"ג מדות הללו נבנית המשנה והתלמוד הקדוש מימות משה ועד רבי אשי ורבינא מיסדי התלמוד כמו שכתב הרמב"ם ז"ל בהקדמת סדר ורעים ומחבד הזהר מלעיג עליהן וקורא עליהם וימררו את חייהם בעבודה קשה דא קושיא, בחמר דא קל וחומר וכו' (זהר בראשית כ"ז) והוציא לעז ודבה רעה על משה רבינו שקבל תורה שבעל פה מפי מטטרי"ן ולא מפי הקב"ה, ולכך היא מעורבת מטוב ורע, וקורא לה אבן משכית, ובה עסקו בן עזאי ובן זומא ואלישע אחר ונסתכנו! וגם קרא על המשנה תחת שלש רגזה ארץ וכו', ודימה שקלא וטריא דרבנן במשנה ובתלמוד ככלבין דצוחין הב הב, וקורא לה גם כן קליפה (זהר פרשת תצא ותיקוני זהר חדש) ועוד הרבה דברים להבזות את לימודי המשנה והתלמוד ולהפחית כבודם ולזלזל בהם, ולתת כבוד ויקר ללימוד תורתו החדשה שהיא מילדי נכרים, ולהוקיר את הבעלים והעשתרות שהמציא ברעיונותיו להטעות את ישראל, ולשבח את סחורתו אשר חשב להנחיל אותנו תחת המשנה והתלמוד ולרשת את מקומם. ועל זה אמרו רז"ל בסנהדרין פ' חלק (סנהדרין דף צ"ט) כי דבר ה' בזה, זה האומר אין תורה מן השמים וכו' ואפילו אמר כל התורה כולה מן השמים חוץ מדקדוק זה, חוץ מקך וחומר זה מגזרה שוה זו, זהו כי דבר ה' בזה וכו' רבי נתן אומר זה שאינו משגיח על המשנה וכו'. תנא דבי ר' ישמעאל כי דבר ה' בזה זה המבזה דיבור שנאמר למשה אנכי ולא יהיה לך ע"כ! וכל הני מילי דאמור רבנן איתנהו בפילסוף מחבר הזהר.

| id | word | punctuation after |
|---|---|---|
| 10:5.1 | קמג | ) |
| 10:5.2 | ולו |  |
| 10:5.3 | עיני |  |
| 10:5.4 | רז"ל |  |
| 10:5.5 | אנשי |  |
| 10:5.6 | כנה"ג |  |
| 10:5.7 | ראו |  |
| 10:5.8 | את |  |
| 10:5.9 | כל |  |
| 10:5.10 | הנעשה |  |
| 10:5.11 | בדת |  |
| 10:5.12 | היהדות |  |
| 10:5.13 | על |  |
| 10:5.14 | פי |  |
| 10:5.15 | הספר |  |
| 10:5.16 | המזוייף |  |
| 10:5.17 | מעבודת |  |
| 10:5.18 | הכחות |  |
| 10:5.19 | הנעלמים |  |
| 10:5.20 | כי |  |
| 10:5.21 | אז |  |
| 10:5.22 | קרעו |  |
| 10:5.23 | את |  |
| 10:5.24 | בגדיהם |  |
| 10:5.25 | ולבשו |  |
| 10:5.26 | שק |  |
| 10:5.27 | ואפר | , |
| 10:5.28 | וצעקו |  |
| 10:5.29 | בחוצות |  |
| 10:5.30 | ובשווקים |  |
| 10:5.31 | וברחובות | , |
| 10:5.32 | צעקה |  |
| 10:5.33 | גדולה |  |
| 10:5.34 | ומרה |  |
| 10:5.35 | כמרדכי |  |
| 10:5.36 | בשושן |  |
| 10:5.37 | הבירה | , |
| 10:5.38 | הזאת |  |
| 10:5.39 | היא |  |
| 10:5.40 | חקת |  |
| 10:5.41 | התורה | ? |
| 10:5.42 | להמציא |  |
| 10:5.43 | בתורה |  |
| 10:5.44 | רמזים |  |
| 10:5.45 | לעבודת |  |
| 10:5.46 | האלילים | ? |
| 10:5.47 | ולימודי |  |
| 10:5.48 | ידיעת |  |
| 10:5.49 | האלילים |  |
| 10:5.50 | האלה |  |
| 10:5.51 | קראו |  |
| 10:5.52 | להם |  |
| 10:5.53 | סתרי |  |
| 10:5.54 | תורה | , |
| 10:5.55 | מלשון |  |
| 10:5.56 | סתירה |  |
| 10:5.57 | שהם |  |
| 10:5.58 | סותרים |  |
| 10:5.59 | ועוקרים |  |
| 10:5.60 | יסודותיה | , |
| 10:5.61 | והורסים |  |
| 10:5.62 | פנותיה |  |
| 10:5.63 | כמו |  |
| 10:5.64 | שדרשו |  |
| 10:5.65 | רז"ל |  |
| 10:5.66 | ילקוט |  |
| 10:5.67 | פרשת |  |
| 10:5.68 | שלח | ) |
| 10:5.69 | אמר |  |
| 10:5.70 | ר | ' |
| 10:5.71 | יצחק |  |
| 10:5.72 | דבר |  |
| 10:5.73 | זה |  |
| 10:5.74 | מסורת |  |
| 10:5.75 | בידינו |  |
| 10:5.76 | מאבותינו |  |
| 10:5.77 | מרגלים |  |
| 10:5.78 | על |  |
| 10:5.79 | שם |  |
| 10:5.80 | מעשיהם |  |
| 10:5.81 | נקראו | , |
| 10:5.82 | ולא |  |
| 10:5.83 | עלתה |  |
| 10:5.84 | בידינו |  |
| 10:5.85 | אלא |  |
| 10:5.86 | אחד | , |
| 10:5.87 | סתור |  |
| 10:5.88 | בן |  |
| 10:5.89 | מיכאל |  |
| 10:5.90 | סתור |  |
| 10:5.91 | שסתר |  |
| 10:5.92 | דבריו |  |
| 10:5.93 | של |  |
| 10:5.94 | הקב"ה | , |
| 10:5.95 | בן |  |
| 10:5.96 | מיכאל | , |
| 10:5.97 | שעשה |  |
| 10:5.98 | עצמו |  |
| 10:5.99 | מך | , |
| 10:5.100 | אמר |  |
| 10:5.101 | ר | ' |
| 10:5.102 | יוחנן |  |
| 10:5.103 | אף |  |
| 10:5.104 | אנו |  |
| 10:5.105 | נאמר |  |
| 10:5.106 | נחבי |  |
| 10:5.107 | בן |  |
| 10:5.108 | ופסי | , |
| 10:5.109 | שהחביא |  |
| 10:5.110 | דבריו |  |
| 10:5.111 | של |  |
| 10:5.112 | הקב"ה |  |
| 10:5.113 | ופסע |  |
| 10:5.114 | עד |  |
| 10:5.115 | מדותיו |  |
| 10:5.116 | של |  |
| 10:5.117 | הקב"ה | . |
| 10:5.118 | כדרך |  |
| 10:5.119 | הזאת |  |
| 10:5.120 | עשה |  |
| 10:5.121 | הפילסוף |  |
| 10:5.122 | המסית |  |
| 10:5.123 | מחבר |  |
| 10:5.124 | הזהר |  |
| 10:5.125 | והנמשכים |  |
| 10:5.126 | אחריו |  |
| 10:5.127 | לסתור |  |
| 10:5.128 | ולקלקל |  |
| 10:5.129 | את |  |
| 10:5.130 | אמונותינו |  |
| 10:5.131 | באל |  |
| 10:5.132 | אחד |  |
| 10:5.133 | מיוחד |  |
| 10:5.134 | מכל |  |
| 10:5.135 | שאר |  |
| 10:5.136 | האחדים | , |
| 10:5.137 | ולפסוע |  |
| 10:5.138 | על |  |
| 10:5.139 | מדותיו |  |
| 10:5.140 | של |  |
| 10:5.141 | הקב"ה |  |
| 10:5.142 | ולעשותן |  |
| 10:5.143 | אלוהות |  |
| 10:5.144 | במה |  |
| 10:5.145 | שקדם | , |
| 10:5.146 | כי |  |
| 10:5.147 | לא |  |
| 10:5.148 | הספיק |  |
| 10:5.149 | לו |  |
| 10:5.150 | לחשוב |  |
| 10:5.151 | באלהות |  |
| 10:5.152 | עילה |  |
| 10:5.153 | אחת | , |
| 10:5.154 | והוצרך |  |
| 10:5.155 | להמציא |  |
| 10:5.156 | ברעיונותיו |  |
| 10:5.157 | עילות |  |
| 10:5.158 | רבות |  |
| 10:5.159 | שבה |  |
| 10:5.160 | נוטל |  |
| 10:5.161 | רשות |  |
| 10:5.162 | מזה |  |
| 10:5.163 | וזה |  |
| 10:5.164 | נעזר |  |
| 10:5.165 | במעשיהו |  |
| 10:5.166 | על |  |
| 10:5.167 | ידי |  |
| 10:5.168 | אחר | , |
| 10:5.169 | כי |  |
| 10:5.170 | חלש |  |
| 10:5.171 | הוא |  |
| 10:5.172 | ולא |  |
| 10:5.173 | יוכל |  |
| 10:5.174 | לעשות |  |
| 10:5.175 | לבדו |  |
| 10:5.176 | כמו |  |
| 10:5.177 | שנתבאר |  |
| 10:5.178 | לעיל | , |
| 10:5.179 | ובחר |  |
| 10:5.180 | לו |  |
| 10:5.181 | לעבוד |  |
| 10:5.182 | את |  |
| 10:5.183 | האליל |  |
| 10:5.184 | קצר |  |
| 10:5.185 | האפים |  |
| 10:5.186 | אשר |  |
| 10:5.187 | לפי |  |
| 10:5.188 | דעתו |  |
| 10:5.189 | לא |  |
| 10:5.190 | נשתתף |  |
| 10:5.191 | בשום |  |
| 10:5.192 | מעשה |  |
| 10:5.193 | מכל |  |
| 10:5.194 | מעשה |  |
| 10:5.195 | בראשית | , |
| 10:5.196 | כי |  |
| 10:5.197 | הוא |  |
| 10:5.198 | נברא |  |
| 10:5.199 | ועדיין |  |
| 10:5.200 | היה |  |
| 10:5.201 | חסר |  |
| 10:5.202 | תיקון |  |
| 10:5.203 | ולא |  |
| 10:5.204 | נשלם |  |
| 10:5.205 | תיקוני |  |
| 10:5.206 | עד |  |
| 10:5.207 | ימי |  |
| 10:5.208 | אברהם |  |
| 10:5.209 | אבינו |  |
| 10:5.210 | כמו |  |
| 10:5.211 | שנתבאר |  |
| 10:5.212 | לעיל |  |
| 10:5.213 | בשם |  |
| 10:5.214 | חיי |  |
| 10:5.215 | שלום |  |
| 10:5.216 | משם |  |
| 10:5.217 | האר"י |  |
| 10:5.218 | ולא |  |
| 10:5.219 | די |  |
| 10:5.220 | זה |  |
| 10:5.221 | אלא |  |
| 10:5.222 | שהוא |  |
| 10:5.223 | מבזה |  |
| 10:5.224 | ומהתל |  |
| 10:5.225 | בשלש |  |
| 10:5.226 | עשרה |  |
| 10:5.227 | מדות |  |
| 10:5.228 | שהתורה |  |
| 10:5.229 | נדרשת |  |
| 10:5.230 | בהם |  |
| 10:5.231 | שקבלם |  |
| 10:5.232 | משה |  |
| 10:5.233 | בסיני | , |
| 10:5.234 | ועל |  |
| 10:5.235 | פ | ' |
| 10:5.236 | חי"ג |  |
| 10:5.237 | מדות |  |
| 10:5.238 | הללו |  |
| 10:5.239 | נבנית |  |
| 10:5.240 | המשנה |  |
| 10:5.241 | והתלמוד |  |
| 10:5.242 | הקדוש |  |
| 10:5.243 | מימות |  |
| 10:5.244 | משה |  |
| 10:5.245 | ועד |  |
| 10:5.246 | רבי |  |
| 10:5.247 | אשי |  |
| 10:5.248 | ורבינא |  |
| 10:5.249 | מיסדי |  |
| 10:5.250 | התלמוד |  |
| 10:5.251 | כמו |  |
| 10:5.252 | שכתב |  |
| 10:5.253 | הרמב"ם |  |
| 10:5.254 | ז"ל |  |
| 10:5.255 | בהקדמת |  |
| 10:5.256 | סדר |  |
| 10:5.257 | ורעים |  |
| 10:5.258 | ומחבד |  |
| 10:5.259 | הזהר |  |
| 10:5.260 | מלעיג |  |
| 10:5.261 | עליהן |  |
| 10:5.262 | וקורא |  |
| 10:5.263 | עליהם |  |
| 10:5.264 | וימררו |  |
| 10:5.265 | את |  |
| 10:5.266 | חייהם |  |
| 10:5.267 | בעבודה |  |
| 10:5.268 | קשה |  |
| 10:5.269 | דא |  |
| 10:5.270 | קושיא | , |
| 10:5.271 | בחמר |  |
| 10:5.272 | דא |  |
| 10:5.273 | קל |  |
| 10:5.274 | וחומר |  |
| 10:5.275 | וכו | ' |
| 10:5.276 | זהר |  |
| 10:5.277 | בראשית |  |
| 10:5.278 | כ"ז | ) |
| 10:5.279 | והוציא |  |
| 10:5.280 | לעז |  |
| 10:5.281 | ודבה |  |
| 10:5.282 | רעה |  |
| 10:5.283 | על |  |
| 10:5.284 | משה |  |
| 10:5.285 | רבינו |  |
| 10:5.286 | שקבל |  |
| 10:5.287 | תורה |  |
| 10:5.288 | שבעל |  |
| 10:5.289 | פה |  |
| 10:5.290 | מפי |  |
| 10:5.291 | מטטרי"ן |  |
| 10:5.292 | ולא |  |
| 10:5.293 | מפי |  |
| 10:5.294 | הקב"ה | , |
| 10:5.295 | ולכך |  |
| 10:5.296 | היא |  |
| 10:5.297 | מעורבת |  |
| 10:5.298 | מטוב |  |
| 10:5.299 | ורע | , |
| 10:5.300 | וקורא |  |
| 10:5.301 | לה |  |
| 10:5.302 | אבן |  |
| 10:5.303 | משכית | , |
| 10:5.304 | ובה |  |
| 10:5.305 | עסקו |  |
| 10:5.306 | בן |  |
| 10:5.307 | עזאי |  |
| 10:5.308 | ובן |  |
| 10:5.309 | זומא |  |
| 10:5.310 | ואלישע |  |
| 10:5.311 | אחר |  |
| 10:5.312 | ונסתכנו | ! |
| 10:5.313 | וגם |  |
| 10:5.314 | קרא |  |
| 10:5.315 | על |  |
| 10:5.316 | המשנה |  |
| 10:5.317 | תחת |  |
| 10:5.318 | שלש |  |
| 10:5.319 | רגזה |  |
| 10:5.320 | ארץ |  |
| 10:5.321 | וכו | ', |
| 10:5.322 | ודימה |  |
| 10:5.323 | שקלא |  |
| 10:5.324 | וטריא |  |
| 10:5.325 | דרבנן |  |
| 10:5.326 | במשנה |  |
| 10:5.327 | ובתלמוד |  |
| 10:5.328 | ככלבין |  |
| 10:5.329 | דצוחין |  |
| 10:5.330 | הב |  |
| 10:5.331 | הב | , |
| 10:5.332 | וקורא |  |
| 10:5.333 | לה |  |
| 10:5.334 | גם |  |
| 10:5.335 | כן |  |
| 10:5.336 | קליפה |  |
| 10:5.337 | זהר |  |
| 10:5.338 | פרשת |  |
| 10:5.339 | תצא |  |
| 10:5.340 | ותיקוני |  |
| 10:5.341 | זהר |  |
| 10:5.342 | חדש | ) |
| 10:5.343 | ועוד |  |
| 10:5.344 | הרבה |  |
| 10:5.345 | דברים |  |
| 10:5.346 | להבזות |  |
| 10:5.347 | את |  |
| 10:5.348 | לימודי |  |
| 10:5.349 | המשנה |  |
| 10:5.350 | והתלמוד |  |
| 10:5.351 | ולהפחית |  |
| 10:5.352 | כבודם |  |
| 10:5.353 | ולזלזל |  |
| 10:5.354 | בהם | , |
| 10:5.355 | ולתת |  |
| 10:5.356 | כבוד |  |
| 10:5.357 | ויקר |  |
| 10:5.358 | ללימוד |  |
| 10:5.359 | תורתו |  |
| 10:5.360 | החדשה |  |
| 10:5.361 | שהיא |  |
| 10:5.362 | מילדי |  |
| 10:5.363 | נכרים | , |
| 10:5.364 | ולהוקיר |  |
| 10:5.365 | את |  |
| 10:5.366 | הבעלים |  |
| 10:5.367 | והעשתרות |  |
| 10:5.368 | שהמציא |  |
| 10:5.369 | ברעיונותיו |  |
| 10:5.370 | להטעות |  |
| 10:5.371 | את |  |
| 10:5.372 | ישראל | , |
| 10:5.373 | ולשבח |  |
| 10:5.374 | את |  |
| 10:5.375 | סחורתו |  |
| 10:5.376 | אשר |  |
| 10:5.377 | חשב |  |
| 10:5.378 | להנחיל |  |
| 10:5.379 | אותנו |  |
| 10:5.380 | תחת |  |
| 10:5.381 | המשנה |  |
| 10:5.382 | והתלמוד |  |
| 10:5.383 | ולרשת |  |
| 10:5.384 | את |  |
| 10:5.385 | מקומם | . |
| 10:5.386 | ועל |  |
| 10:5.387 | זה |  |
| 10:5.388 | אמרו |  |
| 10:5.389 | רז"ל |  |
| 10:5.390 | בסנהדרין |  |
| 10:5.391 | פ | ' |
| 10:5.392 | חלק |  |
| 10:5.393 | סנהדרין |  |
| 10:5.394 | דף |  |
| 10:5.395 | צ"ט | ) |
| 10:5.396 | כי |  |
| 10:5.397 | דבר |  |
| 10:5.398 | ה | ' |
| 10:5.399 | בזה | , |
| 10:5.400 | זה |  |
| 10:5.401 | האומר |  |
| 10:5.402 | אין |  |
| 10:5.403 | תורה |  |
| 10:5.404 | מן |  |
| 10:5.405 | השמים |  |
| 10:5.406 | וכו | ' |
| 10:5.407 | ואפילו |  |
| 10:5.408 | אמר |  |
| 10:5.409 | כל |  |
| 10:5.410 | התורה |  |
| 10:5.411 | כולה |  |
| 10:5.412 | מן |  |
| 10:5.413 | השמים |  |
| 10:5.414 | חוץ |  |
| 10:5.415 | מדקדוק |  |
| 10:5.416 | זה | , |
| 10:5.417 | חוץ |  |
| 10:5.418 | מקך |  |
| 10:5.419 | וחומר |  |
| 10:5.420 | זה |  |
| 10:5.421 | מגזרה |  |
| 10:5.422 | שוה |  |
| 10:5.423 | זו | , |
| 10:5.424 | זהו |  |
| 10:5.425 | כי |  |
| 10:5.426 | דבר |  |
| 10:5.427 | ה | ' |
| 10:5.428 | בזה |  |
| 10:5.429 | וכו | ' |
| 10:5.430 | רבי |  |
| 10:5.431 | נתן |  |
| 10:5.432 | אומר |  |
| 10:5.433 | זה |  |
| 10:5.434 | שאינו |  |
| 10:5.435 | משגיח |  |
| 10:5.436 | על |  |
| 10:5.437 | המשנה |  |
| 10:5.438 | וכו | '. |
| 10:5.439 | תנא |  |
| 10:5.440 | דבי |  |
| 10:5.441 | ר | ' |
| 10:5.442 | ישמעאל |  |
| 10:5.443 | כי |  |
| 10:5.444 | דבר |  |
| 10:5.445 | ה | ' |
| 10:5.446 | בזה |  |
| 10:5.447 | זה |  |
| 10:5.448 | המבזה |  |
| 10:5.449 | דיבור |  |
| 10:5.450 | שנאמר |  |
| 10:5.451 | למשה |  |
| 10:5.452 | אנכי |  |
| 10:5.453 | ולא |  |
| 10:5.454 | יהיה |  |
| 10:5.455 | לך |  |
| 10:5.456 | ע"כ | ! |
| 10:5.457 | וכל |  |
| 10:5.458 | הני |  |
| 10:5.459 | מילי |  |
| 10:5.460 | דאמור |  |
| 10:5.461 | רבנן |  |
| 10:5.462 | איתנהו |  |
| 10:5.463 | בפילסוף |  |
| 10:5.464 | מחבר |  |
| 10:5.465 | הזהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

