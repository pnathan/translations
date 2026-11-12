# Translation pack: milhamot-hashem perek-8-15 (Milchamot ha-Shem 8:18–8:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-15`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3).
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
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
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

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- הסבה **הראשונה** → "the First Cause".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **מין / מינות** → "heretic / heresy".
- תורה **שבעל** פה → "the Oral Torah".
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
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


## Milchamot ha-Shem 8:18

Hebrew (Wikisource): קיב) וכשעלו מן הגולה עמדו אנשי כנסת הגדולה והרבו תפלה ותחנונים לפני הקב"ה וצעקו בקול גדול בייא בייא היינו דאחרביה לבי מקדשא וכו' ובקשו מהשם ית' לסייעם לעקור יצרא דע"ז מישראל שהוא ניהו דאחרביה לבי מקדשא וקלייה להיכלא וכו' ונמסר בידם. היינו שניתן להם סיוע מן השמים וקבעו לישראל בתי מדרשות הרבה. ומינו בהם אנשים חכמים ונבונים ללמד לעם דרכי תורתינו הקדושה, ולהרחיקם מהאמונות הזרות הרחוקות מדרכי התורה. ונפק כי גוריא דנורא, שיצאה וסרה מהם ההתלהבות והחשק והתשוקה לרמזי האמונות הכוזבות שהיו מעמיסים בבית קדש הקדשים. היא תורתינו הקדושה שהיא לנו קדש קדשים המלמדת לאדם דעת ובינה והשכל לדעת את ה' ולאהבה אותו וליראה ממנו, ולעבדו לבדו, לפי שמתחלה היו ממציאים להאמונות הטפלות רמזים וסמוכות דלות ורעות. ותולים אותם בדרכים שונים הרחוקים מקבלת רז"ל שקבלו ממשה רבינו ע"ה דרכי האמונה האמתית שמחוייב כל בעל דת ישראל להאמין כי הבורא ית' שמו הוא אחד ומיוחד מכל שאר האחדים. הוא ראשון והוא אחרון שאיננו נופל תחת הזמן והמקום. ושלא ישיגוהו משיגי הגוף. ושהוא קדמון ונצחי בלתי בעל תכלית. הן מצד הקודם, הן מצד המתאחר כאשר קרא הנביא ישעיה בשם ה' אני ראשון ואני אחרון. וכן תרגם אונקלוס ואמרתי חי אנכי לעולם. קיים אנא לעלמין בלשון רבים ולא תרגם לעלם בלשון יחיד, להורות שהוא ית' קיים תמיד על ענין ואופן אחד בלי שום שינוי ורבוי כלל חלילה. קודם שנברא העולם ואחר שיכלה אם יחפוץ השי"ת (עיין במורה פרק י"ז לשני והקדמה כוללת לרבי נתן אדלר בעל נתינה לגר). והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי. שר"ל אני הוא אשר הייתי ואני אשר אהיה עד בלתי שמים. תמיד אני קיים בתואר אחד בלי שום שינוי. ושהוא ית' הסבה הראשונה ולו לבדו הוא שראוי לעבוד ולהתפלל, ולהודיע גדולתו, ולעשות מצותיו ושלא יעשה כן לשום נברא מכל הנבראים גשמי או רוחני. אלו הן הדעות הרמות המזוקקות צרופות וטהורות. כראי מוצק חזקות. לא כדעת המקובלים החדשים, אשר בעונות הרבים עמודי הדת ירופפו ואדני תורה שבעל פה יתפלצון על ידי הקבלה המזוייפת. לעבוד ולהתפלל לסבות נבראות לקצר אפים (זעיר אנפין) אף על פי שהוא נברא כמבואר בזהר ומפרשיו ובספר הברית ויושר לבב ונחלת יוסף ומאמר האמונה האמתית למהר"י לופיס וזולתם.

| id | word | punctuation after |
|---|---|---|
| 8:18.1 | קיב | ) |
| 8:18.2 | וכשעלו |  |
| 8:18.3 | מן |  |
| 8:18.4 | הגולה |  |
| 8:18.5 | עמדו |  |
| 8:18.6 | אנשי |  |
| 8:18.7 | כנסת |  |
| 8:18.8 | הגדולה |  |
| 8:18.9 | והרבו |  |
| 8:18.10 | תפלה |  |
| 8:18.11 | ותחנונים |  |
| 8:18.12 | לפני |  |
| 8:18.13 | הקב"ה |  |
| 8:18.14 | וצעקו |  |
| 8:18.15 | בקול |  |
| 8:18.16 | גדול |  |
| 8:18.17 | בייא |  |
| 8:18.18 | בייא |  |
| 8:18.19 | היינו |  |
| 8:18.20 | דאחרביה |  |
| 8:18.21 | לבי |  |
| 8:18.22 | מקדשא |  |
| 8:18.23 | וכו | ' |
| 8:18.24 | ובקשו |  |
| 8:18.25 | מהשם |  |
| 8:18.26 | ית | ' |
| 8:18.27 | לסייעם |  |
| 8:18.28 | לעקור |  |
| 8:18.29 | יצרא |  |
| 8:18.30 | דע"ז |  |
| 8:18.31 | מישראל |  |
| 8:18.32 | שהוא |  |
| 8:18.33 | ניהו |  |
| 8:18.34 | דאחרביה |  |
| 8:18.35 | לבי |  |
| 8:18.36 | מקדשא |  |
| 8:18.37 | וקלייה |  |
| 8:18.38 | להיכלא |  |
| 8:18.39 | וכו | ' |
| 8:18.40 | ונמסר |  |
| 8:18.41 | בידם | . |
| 8:18.42 | היינו |  |
| 8:18.43 | שניתן |  |
| 8:18.44 | להם |  |
| 8:18.45 | סיוע |  |
| 8:18.46 | מן |  |
| 8:18.47 | השמים |  |
| 8:18.48 | וקבעו |  |
| 8:18.49 | לישראל |  |
| 8:18.50 | בתי |  |
| 8:18.51 | מדרשות |  |
| 8:18.52 | הרבה | . |
| 8:18.53 | ומינו |  |
| 8:18.54 | בהם |  |
| 8:18.55 | אנשים |  |
| 8:18.56 | חכמים |  |
| 8:18.57 | ונבונים |  |
| 8:18.58 | ללמד |  |
| 8:18.59 | לעם |  |
| 8:18.60 | דרכי |  |
| 8:18.61 | תורתינו |  |
| 8:18.62 | הקדושה | , |
| 8:18.63 | ולהרחיקם |  |
| 8:18.64 | מהאמונות |  |
| 8:18.65 | הזרות |  |
| 8:18.66 | הרחוקות |  |
| 8:18.67 | מדרכי |  |
| 8:18.68 | התורה | . |
| 8:18.69 | ונפק |  |
| 8:18.70 | כי |  |
| 8:18.71 | גוריא |  |
| 8:18.72 | דנורא | , |
| 8:18.73 | שיצאה |  |
| 8:18.74 | וסרה |  |
| 8:18.75 | מהם |  |
| 8:18.76 | ההתלהבות |  |
| 8:18.77 | והחשק |  |
| 8:18.78 | והתשוקה |  |
| 8:18.79 | לרמזי |  |
| 8:18.80 | האמונות |  |
| 8:18.81 | הכוזבות |  |
| 8:18.82 | שהיו |  |
| 8:18.83 | מעמיסים |  |
| 8:18.84 | בבית |  |
| 8:18.85 | קדש |  |
| 8:18.86 | הקדשים | . |
| 8:18.87 | היא |  |
| 8:18.88 | תורתינו |  |
| 8:18.89 | הקדושה |  |
| 8:18.90 | שהיא |  |
| 8:18.91 | לנו |  |
| 8:18.92 | קדש |  |
| 8:18.93 | קדשים |  |
| 8:18.94 | המלמדת |  |
| 8:18.95 | לאדם |  |
| 8:18.96 | דעת |  |
| 8:18.97 | ובינה |  |
| 8:18.98 | והשכל |  |
| 8:18.99 | לדעת |  |
| 8:18.100 | את |  |
| 8:18.101 | ה | ' |
| 8:18.102 | ולאהבה |  |
| 8:18.103 | אותו |  |
| 8:18.104 | וליראה |  |
| 8:18.105 | ממנו | , |
| 8:18.106 | ולעבדו |  |
| 8:18.107 | לבדו | , |
| 8:18.108 | לפי |  |
| 8:18.109 | שמתחלה |  |
| 8:18.110 | היו |  |
| 8:18.111 | ממציאים |  |
| 8:18.112 | להאמונות |  |
| 8:18.113 | הטפלות |  |
| 8:18.114 | רמזים |  |
| 8:18.115 | וסמוכות |  |
| 8:18.116 | דלות |  |
| 8:18.117 | ורעות | . |
| 8:18.118 | ותולים |  |
| 8:18.119 | אותם |  |
| 8:18.120 | בדרכים |  |
| 8:18.121 | שונים |  |
| 8:18.122 | הרחוקים |  |
| 8:18.123 | מקבלת |  |
| 8:18.124 | רז"ל |  |
| 8:18.125 | שקבלו |  |
| 8:18.126 | ממשה |  |
| 8:18.127 | רבינו |  |
| 8:18.128 | ע"ה |  |
| 8:18.129 | דרכי |  |
| 8:18.130 | האמונה |  |
| 8:18.131 | האמתית |  |
| 8:18.132 | שמחוייב |  |
| 8:18.133 | כל |  |
| 8:18.134 | בעל |  |
| 8:18.135 | דת |  |
| 8:18.136 | ישראל |  |
| 8:18.137 | להאמין |  |
| 8:18.138 | כי |  |
| 8:18.139 | הבורא |  |
| 8:18.140 | ית | ' |
| 8:18.141 | שמו |  |
| 8:18.142 | הוא |  |
| 8:18.143 | אחד |  |
| 8:18.144 | ומיוחד |  |
| 8:18.145 | מכל |  |
| 8:18.146 | שאר |  |
| 8:18.147 | האחדים | . |
| 8:18.148 | הוא |  |
| 8:18.149 | ראשון |  |
| 8:18.150 | והוא |  |
| 8:18.151 | אחרון |  |
| 8:18.152 | שאיננו |  |
| 8:18.153 | נופל |  |
| 8:18.154 | תחת |  |
| 8:18.155 | הזמן |  |
| 8:18.156 | והמקום | . |
| 8:18.157 | ושלא |  |
| 8:18.158 | ישיגוהו |  |
| 8:18.159 | משיגי |  |
| 8:18.160 | הגוף | . |
| 8:18.161 | ושהוא |  |
| 8:18.162 | קדמון |  |
| 8:18.163 | ונצחי |  |
| 8:18.164 | בלתי |  |
| 8:18.165 | בעל |  |
| 8:18.166 | תכלית | . |
| 8:18.167 | הן |  |
| 8:18.168 | מצד |  |
| 8:18.169 | הקודם | , |
| 8:18.170 | הן |  |
| 8:18.171 | מצד |  |
| 8:18.172 | המתאחר |  |
| 8:18.173 | כאשר |  |
| 8:18.174 | קרא |  |
| 8:18.175 | הנביא |  |
| 8:18.176 | ישעיה |  |
| 8:18.177 | בשם |  |
| 8:18.178 | ה | ' |
| 8:18.179 | אני |  |
| 8:18.180 | ראשון |  |
| 8:18.181 | ואני |  |
| 8:18.182 | אחרון | . |
| 8:18.183 | וכן |  |
| 8:18.184 | תרגם |  |
| 8:18.185 | אונקלוס |  |
| 8:18.186 | ואמרתי |  |
| 8:18.187 | חי |  |
| 8:18.188 | אנכי |  |
| 8:18.189 | לעולם | . |
| 8:18.190 | קיים |  |
| 8:18.191 | אנא |  |
| 8:18.192 | לעלמין |  |
| 8:18.193 | בלשון |  |
| 8:18.194 | רבים |  |
| 8:18.195 | ולא |  |
| 8:18.196 | תרגם |  |
| 8:18.197 | לעלם |  |
| 8:18.198 | בלשון |  |
| 8:18.199 | יחיד | , |
| 8:18.200 | להורות |  |
| 8:18.201 | שהוא |  |
| 8:18.202 | ית | ' |
| 8:18.203 | קיים |  |
| 8:18.204 | תמיד |  |
| 8:18.205 | על |  |
| 8:18.206 | ענין |  |
| 8:18.207 | ואופן |  |
| 8:18.208 | אחד |  |
| 8:18.209 | בלי |  |
| 8:18.210 | שום |  |
| 8:18.211 | שינוי |  |
| 8:18.212 | ורבוי |  |
| 8:18.213 | כלל |  |
| 8:18.214 | חלילה | . |
| 8:18.215 | קודם |  |
| 8:18.216 | שנברא |  |
| 8:18.217 | העולם |  |
| 8:18.218 | ואחר |  |
| 8:18.219 | שיכלה |  |
| 8:18.220 | אם |  |
| 8:18.221 | יחפוץ |  |
| 8:18.222 | השי"ת |  |
| 8:18.223 | עיין |  |
| 8:18.224 | במורה |  |
| 8:18.225 | פרק |  |
| 8:18.226 | י"ז |  |
| 8:18.227 | לשני |  |
| 8:18.228 | והקדמה |  |
| 8:18.229 | כוללת |  |
| 8:18.230 | לרבי |  |
| 8:18.231 | נתן |  |
| 8:18.232 | אדלר |  |
| 8:18.233 | בעל |  |
| 8:18.234 | נתינה |  |
| 8:18.235 | לגר | ). |
| 8:18.236 | והוא |  |
| 8:18.237 | שאמר |  |
| 8:18.238 | ראו |  |
| 8:18.239 | עתה |  |
| 8:18.240 | כי |  |
| 8:18.241 | אני |  |
| 8:18.242 | אני |  |
| 8:18.243 | הוא |  |
| 8:18.244 | ואין |  |
| 8:18.245 | אלהים |  |
| 8:18.246 | עמדי | . |
| 8:18.247 | שר"ל |  |
| 8:18.248 | אני |  |
| 8:18.249 | הוא |  |
| 8:18.250 | אשר |  |
| 8:18.251 | הייתי |  |
| 8:18.252 | ואני |  |
| 8:18.253 | אשר |  |
| 8:18.254 | אהיה |  |
| 8:18.255 | עד |  |
| 8:18.256 | בלתי |  |
| 8:18.257 | שמים | . |
| 8:18.258 | תמיד |  |
| 8:18.259 | אני |  |
| 8:18.260 | קיים |  |
| 8:18.261 | בתואר |  |
| 8:18.262 | אחד |  |
| 8:18.263 | בלי |  |
| 8:18.264 | שום |  |
| 8:18.265 | שינוי | . |
| 8:18.266 | ושהוא |  |
| 8:18.267 | ית | ' |
| 8:18.268 | הסבה |  |
| 8:18.269 | הראשונה |  |
| 8:18.270 | ולו |  |
| 8:18.271 | לבדו |  |
| 8:18.272 | הוא |  |
| 8:18.273 | שראוי |  |
| 8:18.274 | לעבוד |  |
| 8:18.275 | ולהתפלל | , |
| 8:18.276 | ולהודיע |  |
| 8:18.277 | גדולתו | , |
| 8:18.278 | ולעשות |  |
| 8:18.279 | מצותיו |  |
| 8:18.280 | ושלא |  |
| 8:18.281 | יעשה |  |
| 8:18.282 | כן |  |
| 8:18.283 | לשום |  |
| 8:18.284 | נברא |  |
| 8:18.285 | מכל |  |
| 8:18.286 | הנבראים |  |
| 8:18.287 | גשמי |  |
| 8:18.288 | או |  |
| 8:18.289 | רוחני | . |
| 8:18.290 | אלו |  |
| 8:18.291 | הן |  |
| 8:18.292 | הדעות |  |
| 8:18.293 | הרמות |  |
| 8:18.294 | המזוקקות |  |
| 8:18.295 | צרופות |  |
| 8:18.296 | וטהורות | . |
| 8:18.297 | כראי |  |
| 8:18.298 | מוצק |  |
| 8:18.299 | חזקות | . |
| 8:18.300 | לא |  |
| 8:18.301 | כדעת |  |
| 8:18.302 | המקובלים |  |
| 8:18.303 | החדשים | , |
| 8:18.304 | אשר |  |
| 8:18.305 | בעונות |  |
| 8:18.306 | הרבים |  |
| 8:18.307 | עמודי |  |
| 8:18.308 | הדת |  |
| 8:18.309 | ירופפו |  |
| 8:18.310 | ואדני |  |
| 8:18.311 | תורה |  |
| 8:18.312 | שבעל |  |
| 8:18.313 | פה |  |
| 8:18.314 | יתפלצון |  |
| 8:18.315 | על |  |
| 8:18.316 | ידי |  |
| 8:18.317 | הקבלה |  |
| 8:18.318 | המזוייפת | . |
| 8:18.319 | לעבוד |  |
| 8:18.320 | ולהתפלל |  |
| 8:18.321 | לסבות |  |
| 8:18.322 | נבראות |  |
| 8:18.323 | לקצר |  |
| 8:18.324 | אפים |  |
| 8:18.325 | זעיר |  |
| 8:18.326 | אנפין | ) |
| 8:18.327 | אף |  |
| 8:18.328 | על |  |
| 8:18.329 | פי |  |
| 8:18.330 | שהוא |  |
| 8:18.331 | נברא |  |
| 8:18.332 | כמבואר |  |
| 8:18.333 | בזהר |  |
| 8:18.334 | ומפרשיו |  |
| 8:18.335 | ובספר |  |
| 8:18.336 | הברית |  |
| 8:18.337 | ויושר |  |
| 8:18.338 | לבב |  |
| 8:18.339 | ונחלת |  |
| 8:18.340 | יוסף |  |
| 8:18.341 | ומאמר |  |
| 8:18.342 | האמונה |  |
| 8:18.343 | האמתית |  |
| 8:18.344 | למהר"י |  |
| 8:18.345 | לופיס |  |
| 8:18.346 | וזולתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

