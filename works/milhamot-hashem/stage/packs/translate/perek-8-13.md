# Translation pack: milhamot-hashem perek-8-13 (Milchamot ha-Shem 8:16–8:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-13`.

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
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
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
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- בעל **העיקרים** → "the author of the Ikkarim".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
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
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- אל **אחר** → "another god".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
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
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

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


## Milchamot ha-Shem 8:16

Hebrew (Wikisource): קי) ולעניות דעתי יראה בפי' המאמר בהקדים היסוד שיסדו לנו רז"ל במאמרם על ולא תתורו אחרי לבבכם ואחרי עיניכם. אחרי לבבכם זו מונות, ואחרי עיניכם זו זנות. כי לבא ועינא תרי סרסורי דחטאה נינהו לחכמים בעיניהם ההולכים אחרי דעתם שלא על דרך רז"ל בפשטי הכתובים ודרשות ויסודות המקובלים בידם איש מפי איש עד משה רבינו ע"ה, וסיקלו המסילה לפנינו שלא לנטות ימין ושמאל מן הדרל הישרה אשר דרכו בה רבותינו בעלתי התלמוד והמדרשים האמתיים ואחריהם דרכו בזאת הדרך גם הרס"ג ובעל חובת הלבבות ורבי יהודה הלוי והרמב"ם ובעל העיקרים וס' מצות גדול, ורבים מחכמינו אנשים חכמים ונבונים. וידועים לנו בחכמתם ובינתם. עד אשר נתגלה ספר הזהר במרמת המלך העכו"ם שאמר שמצאו בעפר ושלחו לחכמי טוליטולה בעת אשר נתמעטה חכמת חכמינו ובינת נבונינו הסתתרה, וחשבו הנמנעות לנמצאות. וחידות רבותינו בעלי התלמוד והמדרשים ומשליהם שהם כתפוחי זהב במשכיות כסף האמינו בפירושן ככל העולה על רוחם. על פי הטעאת הזהר אשר יהסוהו לרשב"י בשקר. ובלי מדע ובינה והשכל האמינו לו לחשוב קבלתו המזויפת לקבלה אמתית. ויעשוה לגבירה ובת מלך, ואת הקבלה האמתית המשנה והתלמוד שמו לשפחה חרופה. ויהרפוה ויכניעוה ויכריעוה תחת ממשלתה לרדותה בחזק יד ובפרך. ולקבוע מנהגים ודינים חדשים שלא כרצונה. והתירו אז פרושים את הדבר לחשוב באלהות דברים אשר לא כן על ה' אלהיהם לעשות לו דמות ותבנית ותמונה! ושלשלו מהאלהים העליון כמה פרצופים. וקראו להם בשם הוי'ה ואדנות ושאר השמות המיוחדים כאשר עשו דורו של אנוש שקראו לעצבים בשם ה' ונסמכו על משענת קנה רצוץ המיוחס לרשב"י בשקר שהמציא שני מיני אלוהות הנמשכים ומשתלשלים מהסבה הראשונה. האחד מהם מערכת אלהים קדושים פועלים הטוב והשני מערכת אלהים טמאים סטרא אחרא מסאבא פועלים הרע. וישם את הכל הקדושים והטמאים בסתרו של עולם. ואת ה' אלהים אמת האל הגדול הגבור והנורא ארך אפים ורב חסד וכו' סלקו למעלה, ואמרו שלא יאות לשבחו ולפארו ולהתפלל אליו כי הוא מסולק ומרומם מכל ברכה ותהלה. ולא יענה את הקורא ומשוע אליו, כמו שטעו דור אנוש אשר היו כעורים. מגמת פניהם אל המלך ותעו מן הדרך וחשבו כי לחפצם נגעו, ובא להם זה מחסרון חכמה ומדע. כמו שאמרו רז"ל המובא לעיל, והרמב"ם ז"ל בהל' ע"ז שזהו טעותם של דור אנוש שאמרו שלערך יקרתו ומעלתו של אלהינו חלק ממשלה לזולתו ( ובס' עז לאלהים כתב בבית קדש הקדשים שנקרא מלכא קדישא. קודשא בריך הוא על שם שהוא מתברך ממה שלמעלה ממנו מה שאין כן האין סוף אין לו ממי להתברך על כן לא יצדק בו שם קודשא בריך הוא) ואינו משגיח בתחתונים, ולא יפנה אל שבחי ותהלות המשבחים ומפארים אותו, ושגם לא נקרא לו שם כלל, וכל שבח ותהלה ותפלה יחסו, אל אלוה קצר אפים (זעיר אנפין) ואמרו שהוא המלך המהולל בכל התשבוחות, כמו שכתב בביאור הרב ספר הברית והרב יושר לבב ומהר"י לופיס ונחלת יוסף ורבים מהמקובלים החדשים שנמשכו אחר דעת המסית מחבר הזהר.

| id | word | punctuation after |
|---|---|---|
| 8:16.1 | קי | ) |
| 8:16.2 | ולעניות |  |
| 8:16.3 | דעתי |  |
| 8:16.4 | יראה |  |
| 8:16.5 | בפי | ' |
| 8:16.6 | המאמר |  |
| 8:16.7 | בהקדים |  |
| 8:16.8 | היסוד |  |
| 8:16.9 | שיסדו |  |
| 8:16.10 | לנו |  |
| 8:16.11 | רז"ל |  |
| 8:16.12 | במאמרם |  |
| 8:16.13 | על |  |
| 8:16.14 | ולא |  |
| 8:16.15 | תתורו |  |
| 8:16.16 | אחרי |  |
| 8:16.17 | לבבכם |  |
| 8:16.18 | ואחרי |  |
| 8:16.19 | עיניכם | . |
| 8:16.20 | אחרי |  |
| 8:16.21 | לבבכם |  |
| 8:16.22 | זו |  |
| 8:16.23 | מונות | , |
| 8:16.24 | ואחרי |  |
| 8:16.25 | עיניכם |  |
| 8:16.26 | זו |  |
| 8:16.27 | זנות | . |
| 8:16.28 | כי |  |
| 8:16.29 | לבא |  |
| 8:16.30 | ועינא |  |
| 8:16.31 | תרי |  |
| 8:16.32 | סרסורי |  |
| 8:16.33 | דחטאה |  |
| 8:16.34 | נינהו |  |
| 8:16.35 | לחכמים |  |
| 8:16.36 | בעיניהם |  |
| 8:16.37 | ההולכים |  |
| 8:16.38 | אחרי |  |
| 8:16.39 | דעתם |  |
| 8:16.40 | שלא |  |
| 8:16.41 | על |  |
| 8:16.42 | דרך |  |
| 8:16.43 | רז"ל |  |
| 8:16.44 | בפשטי |  |
| 8:16.45 | הכתובים |  |
| 8:16.46 | ודרשות |  |
| 8:16.47 | ויסודות |  |
| 8:16.48 | המקובלים |  |
| 8:16.49 | בידם |  |
| 8:16.50 | איש |  |
| 8:16.51 | מפי |  |
| 8:16.52 | איש |  |
| 8:16.53 | עד |  |
| 8:16.54 | משה |  |
| 8:16.55 | רבינו |  |
| 8:16.56 | ע"ה | , |
| 8:16.57 | וסיקלו |  |
| 8:16.58 | המסילה |  |
| 8:16.59 | לפנינו |  |
| 8:16.60 | שלא |  |
| 8:16.61 | לנטות |  |
| 8:16.62 | ימין |  |
| 8:16.63 | ושמאל |  |
| 8:16.64 | מן |  |
| 8:16.65 | הדרל |  |
| 8:16.66 | הישרה |  |
| 8:16.67 | אשר |  |
| 8:16.68 | דרכו |  |
| 8:16.69 | בה |  |
| 8:16.70 | רבותינו |  |
| 8:16.71 | בעלתי |  |
| 8:16.72 | התלמוד |  |
| 8:16.73 | והמדרשים |  |
| 8:16.74 | האמתיים |  |
| 8:16.75 | ואחריהם |  |
| 8:16.76 | דרכו |  |
| 8:16.77 | בזאת |  |
| 8:16.78 | הדרך |  |
| 8:16.79 | גם |  |
| 8:16.80 | הרס"ג |  |
| 8:16.81 | ובעל |  |
| 8:16.82 | חובת |  |
| 8:16.83 | הלבבות |  |
| 8:16.84 | ורבי |  |
| 8:16.85 | יהודה |  |
| 8:16.86 | הלוי |  |
| 8:16.87 | והרמב"ם |  |
| 8:16.88 | ובעל |  |
| 8:16.89 | העיקרים |  |
| 8:16.90 | וס | ' |
| 8:16.91 | מצות |  |
| 8:16.92 | גדול | , |
| 8:16.93 | ורבים |  |
| 8:16.94 | מחכמינו |  |
| 8:16.95 | אנשים |  |
| 8:16.96 | חכמים |  |
| 8:16.97 | ונבונים | . |
| 8:16.98 | וידועים |  |
| 8:16.99 | לנו |  |
| 8:16.100 | בחכמתם |  |
| 8:16.101 | ובינתם | . |
| 8:16.102 | עד |  |
| 8:16.103 | אשר |  |
| 8:16.104 | נתגלה |  |
| 8:16.105 | ספר |  |
| 8:16.106 | הזהר |  |
| 8:16.107 | במרמת |  |
| 8:16.108 | המלך |  |
| 8:16.109 | העכו"ם |  |
| 8:16.110 | שאמר |  |
| 8:16.111 | שמצאו |  |
| 8:16.112 | בעפר |  |
| 8:16.113 | ושלחו |  |
| 8:16.114 | לחכמי |  |
| 8:16.115 | טוליטולה |  |
| 8:16.116 | בעת |  |
| 8:16.117 | אשר |  |
| 8:16.118 | נתמעטה |  |
| 8:16.119 | חכמת |  |
| 8:16.120 | חכמינו |  |
| 8:16.121 | ובינת |  |
| 8:16.122 | נבונינו |  |
| 8:16.123 | הסתתרה | , |
| 8:16.124 | וחשבו |  |
| 8:16.125 | הנמנעות |  |
| 8:16.126 | לנמצאות | . |
| 8:16.127 | וחידות |  |
| 8:16.128 | רבותינו |  |
| 8:16.129 | בעלי |  |
| 8:16.130 | התלמוד |  |
| 8:16.131 | והמדרשים |  |
| 8:16.132 | ומשליהם |  |
| 8:16.133 | שהם |  |
| 8:16.134 | כתפוחי |  |
| 8:16.135 | זהב |  |
| 8:16.136 | במשכיות |  |
| 8:16.137 | כסף |  |
| 8:16.138 | האמינו |  |
| 8:16.139 | בפירושן |  |
| 8:16.140 | ככל |  |
| 8:16.141 | העולה |  |
| 8:16.142 | על |  |
| 8:16.143 | רוחם | . |
| 8:16.144 | על |  |
| 8:16.145 | פי |  |
| 8:16.146 | הטעאת |  |
| 8:16.147 | הזהר |  |
| 8:16.148 | אשר |  |
| 8:16.149 | יהסוהו |  |
| 8:16.150 | לרשב"י |  |
| 8:16.151 | בשקר | . |
| 8:16.152 | ובלי |  |
| 8:16.153 | מדע |  |
| 8:16.154 | ובינה |  |
| 8:16.155 | והשכל |  |
| 8:16.156 | האמינו |  |
| 8:16.157 | לו |  |
| 8:16.158 | לחשוב |  |
| 8:16.159 | קבלתו |  |
| 8:16.160 | המזויפת |  |
| 8:16.161 | לקבלה |  |
| 8:16.162 | אמתית | . |
| 8:16.163 | ויעשוה |  |
| 8:16.164 | לגבירה |  |
| 8:16.165 | ובת |  |
| 8:16.166 | מלך | , |
| 8:16.167 | ואת |  |
| 8:16.168 | הקבלה |  |
| 8:16.169 | האמתית |  |
| 8:16.170 | המשנה |  |
| 8:16.171 | והתלמוד |  |
| 8:16.172 | שמו |  |
| 8:16.173 | לשפחה |  |
| 8:16.174 | חרופה | . |
| 8:16.175 | ויהרפוה |  |
| 8:16.176 | ויכניעוה |  |
| 8:16.177 | ויכריעוה |  |
| 8:16.178 | תחת |  |
| 8:16.179 | ממשלתה |  |
| 8:16.180 | לרדותה |  |
| 8:16.181 | בחזק |  |
| 8:16.182 | יד |  |
| 8:16.183 | ובפרך | . |
| 8:16.184 | ולקבוע |  |
| 8:16.185 | מנהגים |  |
| 8:16.186 | ודינים |  |
| 8:16.187 | חדשים |  |
| 8:16.188 | שלא |  |
| 8:16.189 | כרצונה | . |
| 8:16.190 | והתירו |  |
| 8:16.191 | אז |  |
| 8:16.192 | פרושים |  |
| 8:16.193 | את |  |
| 8:16.194 | הדבר |  |
| 8:16.195 | לחשוב |  |
| 8:16.196 | באלהות |  |
| 8:16.197 | דברים |  |
| 8:16.198 | אשר |  |
| 8:16.199 | לא |  |
| 8:16.200 | כן |  |
| 8:16.201 | על |  |
| 8:16.202 | ה | ' |
| 8:16.203 | אלהיהם |  |
| 8:16.204 | לעשות |  |
| 8:16.205 | לו |  |
| 8:16.206 | דמות |  |
| 8:16.207 | ותבנית |  |
| 8:16.208 | ותמונה | ! |
| 8:16.209 | ושלשלו |  |
| 8:16.210 | מהאלהים |  |
| 8:16.211 | העליון |  |
| 8:16.212 | כמה |  |
| 8:16.213 | פרצופים | . |
| 8:16.214 | וקראו |  |
| 8:16.215 | להם |  |
| 8:16.216 | בשם |  |
| 8:16.217 | הוי'ה |  |
| 8:16.218 | ואדנות |  |
| 8:16.219 | ושאר |  |
| 8:16.220 | השמות |  |
| 8:16.221 | המיוחדים |  |
| 8:16.222 | כאשר |  |
| 8:16.223 | עשו |  |
| 8:16.224 | דורו |  |
| 8:16.225 | של |  |
| 8:16.226 | אנוש |  |
| 8:16.227 | שקראו |  |
| 8:16.228 | לעצבים |  |
| 8:16.229 | בשם |  |
| 8:16.230 | ה | ' |
| 8:16.231 | ונסמכו |  |
| 8:16.232 | על |  |
| 8:16.233 | משענת |  |
| 8:16.234 | קנה |  |
| 8:16.235 | רצוץ |  |
| 8:16.236 | המיוחס |  |
| 8:16.237 | לרשב"י |  |
| 8:16.238 | בשקר |  |
| 8:16.239 | שהמציא |  |
| 8:16.240 | שני |  |
| 8:16.241 | מיני |  |
| 8:16.242 | אלוהות |  |
| 8:16.243 | הנמשכים |  |
| 8:16.244 | ומשתלשלים |  |
| 8:16.245 | מהסבה |  |
| 8:16.246 | הראשונה | . |
| 8:16.247 | האחד |  |
| 8:16.248 | מהם |  |
| 8:16.249 | מערכת |  |
| 8:16.250 | אלהים |  |
| 8:16.251 | קדושים |  |
| 8:16.252 | פועלים |  |
| 8:16.253 | הטוב |  |
| 8:16.254 | והשני |  |
| 8:16.255 | מערכת |  |
| 8:16.256 | אלהים |  |
| 8:16.257 | טמאים |  |
| 8:16.258 | סטרא |  |
| 8:16.259 | אחרא |  |
| 8:16.260 | מסאבא |  |
| 8:16.261 | פועלים |  |
| 8:16.262 | הרע | . |
| 8:16.263 | וישם |  |
| 8:16.264 | את |  |
| 8:16.265 | הכל |  |
| 8:16.266 | הקדושים |  |
| 8:16.267 | והטמאים |  |
| 8:16.268 | בסתרו |  |
| 8:16.269 | של |  |
| 8:16.270 | עולם | . |
| 8:16.271 | ואת |  |
| 8:16.272 | ה | ' |
| 8:16.273 | אלהים |  |
| 8:16.274 | אמת |  |
| 8:16.275 | האל |  |
| 8:16.276 | הגדול |  |
| 8:16.277 | הגבור |  |
| 8:16.278 | והנורא |  |
| 8:16.279 | ארך |  |
| 8:16.280 | אפים |  |
| 8:16.281 | ורב |  |
| 8:16.282 | חסד |  |
| 8:16.283 | וכו | ' |
| 8:16.284 | סלקו |  |
| 8:16.285 | למעלה | , |
| 8:16.286 | ואמרו |  |
| 8:16.287 | שלא |  |
| 8:16.288 | יאות |  |
| 8:16.289 | לשבחו |  |
| 8:16.290 | ולפארו |  |
| 8:16.291 | ולהתפלל |  |
| 8:16.292 | אליו |  |
| 8:16.293 | כי |  |
| 8:16.294 | הוא |  |
| 8:16.295 | מסולק |  |
| 8:16.296 | ומרומם |  |
| 8:16.297 | מכל |  |
| 8:16.298 | ברכה |  |
| 8:16.299 | ותהלה | . |
| 8:16.300 | ולא |  |
| 8:16.301 | יענה |  |
| 8:16.302 | את |  |
| 8:16.303 | הקורא |  |
| 8:16.304 | ומשוע |  |
| 8:16.305 | אליו | , |
| 8:16.306 | כמו |  |
| 8:16.307 | שטעו |  |
| 8:16.308 | דור |  |
| 8:16.309 | אנוש |  |
| 8:16.310 | אשר |  |
| 8:16.311 | היו |  |
| 8:16.312 | כעורים | . |
| 8:16.313 | מגמת |  |
| 8:16.314 | פניהם |  |
| 8:16.315 | אל |  |
| 8:16.316 | המלך |  |
| 8:16.317 | ותעו |  |
| 8:16.318 | מן |  |
| 8:16.319 | הדרך |  |
| 8:16.320 | וחשבו |  |
| 8:16.321 | כי |  |
| 8:16.322 | לחפצם |  |
| 8:16.323 | נגעו | , |
| 8:16.324 | ובא |  |
| 8:16.325 | להם |  |
| 8:16.326 | זה |  |
| 8:16.327 | מחסרון |  |
| 8:16.328 | חכמה |  |
| 8:16.329 | ומדע | . |
| 8:16.330 | כמו |  |
| 8:16.331 | שאמרו |  |
| 8:16.332 | רז"ל |  |
| 8:16.333 | המובא |  |
| 8:16.334 | לעיל | , |
| 8:16.335 | והרמב"ם |  |
| 8:16.336 | ז"ל |  |
| 8:16.337 | בהל | ' |
| 8:16.338 | ע"ז |  |
| 8:16.339 | שזהו |  |
| 8:16.340 | טעותם |  |
| 8:16.341 | של |  |
| 8:16.342 | דור |  |
| 8:16.343 | אנוש |  |
| 8:16.344 | שאמרו |  |
| 8:16.345 | שלערך |  |
| 8:16.346 | יקרתו |  |
| 8:16.347 | ומעלתו |  |
| 8:16.348 | של |  |
| 8:16.349 | אלהינו |  |
| 8:16.350 | חלק |  |
| 8:16.351 | ממשלה |  |
| 8:16.352 | לזולתו | ( |
| 8:16.353 | ובס | ' |
| 8:16.354 | עז |  |
| 8:16.355 | לאלהים |  |
| 8:16.356 | כתב |  |
| 8:16.357 | בבית |  |
| 8:16.358 | קדש |  |
| 8:16.359 | הקדשים |  |
| 8:16.360 | שנקרא |  |
| 8:16.361 | מלכא |  |
| 8:16.362 | קדישא | . |
| 8:16.363 | קודשא |  |
| 8:16.364 | בריך |  |
| 8:16.365 | הוא |  |
| 8:16.366 | על |  |
| 8:16.367 | שם |  |
| 8:16.368 | שהוא |  |
| 8:16.369 | מתברך |  |
| 8:16.370 | ממה |  |
| 8:16.371 | שלמעלה |  |
| 8:16.372 | ממנו |  |
| 8:16.373 | מה |  |
| 8:16.374 | שאין |  |
| 8:16.375 | כן |  |
| 8:16.376 | האין |  |
| 8:16.377 | סוף |  |
| 8:16.378 | אין |  |
| 8:16.379 | לו |  |
| 8:16.380 | ממי |  |
| 8:16.381 | להתברך |  |
| 8:16.382 | על |  |
| 8:16.383 | כן |  |
| 8:16.384 | לא |  |
| 8:16.385 | יצדק |  |
| 8:16.386 | בו |  |
| 8:16.387 | שם |  |
| 8:16.388 | קודשא |  |
| 8:16.389 | בריך |  |
| 8:16.390 | הוא | ) |
| 8:16.391 | ואינו |  |
| 8:16.392 | משגיח |  |
| 8:16.393 | בתחתונים | , |
| 8:16.394 | ולא |  |
| 8:16.395 | יפנה |  |
| 8:16.396 | אל |  |
| 8:16.397 | שבחי |  |
| 8:16.398 | ותהלות |  |
| 8:16.399 | המשבחים |  |
| 8:16.400 | ומפארים |  |
| 8:16.401 | אותו | , |
| 8:16.402 | ושגם |  |
| 8:16.403 | לא |  |
| 8:16.404 | נקרא |  |
| 8:16.405 | לו |  |
| 8:16.406 | שם |  |
| 8:16.407 | כלל | , |
| 8:16.408 | וכל |  |
| 8:16.409 | שבח |  |
| 8:16.410 | ותהלה |  |
| 8:16.411 | ותפלה |  |
| 8:16.412 | יחסו | , |
| 8:16.413 | אל |  |
| 8:16.414 | אלוה |  |
| 8:16.415 | קצר |  |
| 8:16.416 | אפים |  |
| 8:16.417 | זעיר |  |
| 8:16.418 | אנפין | ) |
| 8:16.419 | ואמרו |  |
| 8:16.420 | שהוא |  |
| 8:16.421 | המלך |  |
| 8:16.422 | המהולל |  |
| 8:16.423 | בכל |  |
| 8:16.424 | התשבוחות | , |
| 8:16.425 | כמו |  |
| 8:16.426 | שכתב |  |
| 8:16.427 | בביאור |  |
| 8:16.428 | הרב |  |
| 8:16.429 | ספר |  |
| 8:16.430 | הברית |  |
| 8:16.431 | והרב |  |
| 8:16.432 | יושר |  |
| 8:16.433 | לבב |  |
| 8:16.434 | ומהר"י |  |
| 8:16.435 | לופיס |  |
| 8:16.436 | ונחלת |  |
| 8:16.437 | יוסף |  |
| 8:16.438 | ורבים |  |
| 8:16.439 | מהמקובלים |  |
| 8:16.440 | החדשים |  |
| 8:16.441 | שנמשכו |  |
| 8:16.442 | אחר |  |
| 8:16.443 | דעת |  |
| 8:16.444 | המסית |  |
| 8:16.445 | מחבר |  |
| 8:16.446 | הזהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

