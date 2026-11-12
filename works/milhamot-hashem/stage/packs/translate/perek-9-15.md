# Translation pack: milhamot-hashem perek-9-15 (Milchamot ha-Shem 9:18–9:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-15`.

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
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
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
- **רז"ל** → "our Rabbis of blessed memory".
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
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
- הסבה **הראשונה** → "the First Cause".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **עלה / עלול** → "cause / effect".
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
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- ואין **צריך** לומר → "needless to say".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
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
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
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


## Milchamot ha-Shem 9:18

Hebrew (Wikisource): קלו) וגם רבותינו בעלי הקבלה האמתית במשנה ובתלמוד, למה לא גלו לנו אמונה זו שהיא נשמת התורה (לפי בריהם) אשר בעבורה נזכה לחיי עד חיים הנצחיים לעולם שכולו טוב לעולם שכולו ארוך. שהעבודה להעלולים הנמשכים ממנו ית' שהם קב"ה ושכינתיה? ולמה הניחו את כל המון ישראל וחכמיהם העוסקים במשנה ובתלמוד בטעותם לעבוד ללא אלהי אמת שהיא הסבה הראשונה בלי חבור בלי פירוד, באופן שתפלתם אינה רצויה (לפי דעתם) ולא גילו להם שהעבודה הרצויה היא דוקא לאלוה האחרון קצר אפים (ז"א) כדי לזכותם שתקובל תפלתם ועבודתם לרצון על ידי שיחברו את כל הנתחים אשר ערבו המקובלים החדשים זכרים ונקבות. ובפרט קצר האפים עם שתי נשיו לאה ורחל, אשר האחת היא למעלה אחורי הדעת והיא הנקראת לאה, והשניה שהיא למטה אחורי התפארת, והיא הנקראת רחל. והיא בחינת תפלה של יד (פרי עץ חיים שער התפלין פ"ב)? ולמה אמרו רז"ל במסכת דבחים לשם ששה דברים הזבח נזבח, לשם הזבח, לשם הזובח, לשם השם, לשם אישים, לשם ריח. לשם ניחוח, ולא תקנו לכהנים לומר לשם יחוד קב"ה ושכינתיה וכו' על ידי ההוא טמיר ונעלם שזהו העיקר הגדול בדברי המקובלים ליחד הזכרים עם נקיבותיהם? ואיך קראו רז"ל לההלכות האמורות במשנה ובתלמוד כתרה של תורה כדאמרינן במגלה דף כ"ב ודאשתמש בתנא חלף. תני ר"ל זה המשתמש במי ששונה הלכות כתרה של תורה ע"כ. כי לפי דברי הזהר והמקובלים החדשים הוא תבנא דאוריתא. כמוץ ותבן שזורין ובוררין ומפרישין אותו להשליכו חוצה! ולא די זה אלא שצריך העוסק בהם לשוב מחטאתו ולעזוב אותם ולעסוק בנסתרות הזהר וסודותיו. ואך אז יזכה ליאור באור החיים. את כל אלה הרעות עשה הפילסוף המסית מחבר הזהר להבזות את המשנה והתלמוד הקדוש המקובל ממשה איש מפי איש ולחלל את כבדן כמו שנתבאר לעיל.

| id | word | punctuation after |
|---|---|---|
| 9:18.1 | קלו | ) |
| 9:18.2 | וגם |  |
| 9:18.3 | רבותינו |  |
| 9:18.4 | בעלי |  |
| 9:18.5 | הקבלה |  |
| 9:18.6 | האמתית |  |
| 9:18.7 | במשנה |  |
| 9:18.8 | ובתלמוד | , |
| 9:18.9 | למה |  |
| 9:18.10 | לא |  |
| 9:18.11 | גלו |  |
| 9:18.12 | לנו |  |
| 9:18.13 | אמונה |  |
| 9:18.14 | זו |  |
| 9:18.15 | שהיא |  |
| 9:18.16 | נשמת |  |
| 9:18.17 | התורה |  |
| 9:18.18 | לפי |  |
| 9:18.19 | בריהם | ) |
| 9:18.20 | אשר |  |
| 9:18.21 | בעבורה |  |
| 9:18.22 | נזכה |  |
| 9:18.23 | לחיי |  |
| 9:18.24 | עד |  |
| 9:18.25 | חיים |  |
| 9:18.26 | הנצחיים |  |
| 9:18.27 | לעולם |  |
| 9:18.28 | שכולו |  |
| 9:18.29 | טוב |  |
| 9:18.30 | לעולם |  |
| 9:18.31 | שכולו |  |
| 9:18.32 | ארוך | . |
| 9:18.33 | שהעבודה |  |
| 9:18.34 | להעלולים |  |
| 9:18.35 | הנמשכים |  |
| 9:18.36 | ממנו |  |
| 9:18.37 | ית | ' |
| 9:18.38 | שהם |  |
| 9:18.39 | קב"ה |  |
| 9:18.40 | ושכינתיה | ? |
| 9:18.41 | ולמה |  |
| 9:18.42 | הניחו |  |
| 9:18.43 | את |  |
| 9:18.44 | כל |  |
| 9:18.45 | המון |  |
| 9:18.46 | ישראל |  |
| 9:18.47 | וחכמיהם |  |
| 9:18.48 | העוסקים |  |
| 9:18.49 | במשנה |  |
| 9:18.50 | ובתלמוד |  |
| 9:18.51 | בטעותם |  |
| 9:18.52 | לעבוד |  |
| 9:18.53 | ללא |  |
| 9:18.54 | אלהי |  |
| 9:18.55 | אמת |  |
| 9:18.56 | שהיא |  |
| 9:18.57 | הסבה |  |
| 9:18.58 | הראשונה |  |
| 9:18.59 | בלי |  |
| 9:18.60 | חבור |  |
| 9:18.61 | בלי |  |
| 9:18.62 | פירוד | , |
| 9:18.63 | באופן |  |
| 9:18.64 | שתפלתם |  |
| 9:18.65 | אינה |  |
| 9:18.66 | רצויה |  |
| 9:18.67 | לפי |  |
| 9:18.68 | דעתם | ) |
| 9:18.69 | ולא |  |
| 9:18.70 | גילו |  |
| 9:18.71 | להם |  |
| 9:18.72 | שהעבודה |  |
| 9:18.73 | הרצויה |  |
| 9:18.74 | היא |  |
| 9:18.75 | דוקא |  |
| 9:18.76 | לאלוה |  |
| 9:18.77 | האחרון |  |
| 9:18.78 | קצר |  |
| 9:18.79 | אפים |  |
| 9:18.80 | ז"א | ) |
| 9:18.81 | כדי |  |
| 9:18.82 | לזכותם |  |
| 9:18.83 | שתקובל |  |
| 9:18.84 | תפלתם |  |
| 9:18.85 | ועבודתם |  |
| 9:18.86 | לרצון |  |
| 9:18.87 | על |  |
| 9:18.88 | ידי |  |
| 9:18.89 | שיחברו |  |
| 9:18.90 | את |  |
| 9:18.91 | כל |  |
| 9:18.92 | הנתחים |  |
| 9:18.93 | אשר |  |
| 9:18.94 | ערבו |  |
| 9:18.95 | המקובלים |  |
| 9:18.96 | החדשים |  |
| 9:18.97 | זכרים |  |
| 9:18.98 | ונקבות | . |
| 9:18.99 | ובפרט |  |
| 9:18.100 | קצר |  |
| 9:18.101 | האפים |  |
| 9:18.102 | עם |  |
| 9:18.103 | שתי |  |
| 9:18.104 | נשיו |  |
| 9:18.105 | לאה |  |
| 9:18.106 | ורחל | , |
| 9:18.107 | אשר |  |
| 9:18.108 | האחת |  |
| 9:18.109 | היא |  |
| 9:18.110 | למעלה |  |
| 9:18.111 | אחורי |  |
| 9:18.112 | הדעת |  |
| 9:18.113 | והיא |  |
| 9:18.114 | הנקראת |  |
| 9:18.115 | לאה | , |
| 9:18.116 | והשניה |  |
| 9:18.117 | שהיא |  |
| 9:18.118 | למטה |  |
| 9:18.119 | אחורי |  |
| 9:18.120 | התפארת | , |
| 9:18.121 | והיא |  |
| 9:18.122 | הנקראת |  |
| 9:18.123 | רחל | . |
| 9:18.124 | והיא |  |
| 9:18.125 | בחינת |  |
| 9:18.126 | תפלה |  |
| 9:18.127 | של |  |
| 9:18.128 | יד |  |
| 9:18.129 | פרי |  |
| 9:18.130 | עץ |  |
| 9:18.131 | חיים |  |
| 9:18.132 | שער |  |
| 9:18.133 | התפלין |  |
| 9:18.134 | פ"ב | )? |
| 9:18.135 | ולמה |  |
| 9:18.136 | אמרו |  |
| 9:18.137 | רז"ל |  |
| 9:18.138 | במסכת |  |
| 9:18.139 | דבחים |  |
| 9:18.140 | לשם |  |
| 9:18.141 | ששה |  |
| 9:18.142 | דברים |  |
| 9:18.143 | הזבח |  |
| 9:18.144 | נזבח | , |
| 9:18.145 | לשם |  |
| 9:18.146 | הזבח | , |
| 9:18.147 | לשם |  |
| 9:18.148 | הזובח | , |
| 9:18.149 | לשם |  |
| 9:18.150 | השם | , |
| 9:18.151 | לשם |  |
| 9:18.152 | אישים | , |
| 9:18.153 | לשם |  |
| 9:18.154 | ריח | . |
| 9:18.155 | לשם |  |
| 9:18.156 | ניחוח | , |
| 9:18.157 | ולא |  |
| 9:18.158 | תקנו |  |
| 9:18.159 | לכהנים |  |
| 9:18.160 | לומר |  |
| 9:18.161 | לשם |  |
| 9:18.162 | יחוד |  |
| 9:18.163 | קב"ה |  |
| 9:18.164 | ושכינתיה |  |
| 9:18.165 | וכו | ' |
| 9:18.166 | על |  |
| 9:18.167 | ידי |  |
| 9:18.168 | ההוא |  |
| 9:18.169 | טמיר |  |
| 9:18.170 | ונעלם |  |
| 9:18.171 | שזהו |  |
| 9:18.172 | העיקר |  |
| 9:18.173 | הגדול |  |
| 9:18.174 | בדברי |  |
| 9:18.175 | המקובלים |  |
| 9:18.176 | ליחד |  |
| 9:18.177 | הזכרים |  |
| 9:18.178 | עם |  |
| 9:18.179 | נקיבותיהם | ? |
| 9:18.180 | ואיך |  |
| 9:18.181 | קראו |  |
| 9:18.182 | רז"ל |  |
| 9:18.183 | לההלכות |  |
| 9:18.184 | האמורות |  |
| 9:18.185 | במשנה |  |
| 9:18.186 | ובתלמוד |  |
| 9:18.187 | כתרה |  |
| 9:18.188 | של |  |
| 9:18.189 | תורה |  |
| 9:18.190 | כדאמרינן |  |
| 9:18.191 | במגלה |  |
| 9:18.192 | דף |  |
| 9:18.193 | כ"ב |  |
| 9:18.194 | ודאשתמש |  |
| 9:18.195 | בתנא |  |
| 9:18.196 | חלף | . |
| 9:18.197 | תני |  |
| 9:18.198 | ר"ל |  |
| 9:18.199 | זה |  |
| 9:18.200 | המשתמש |  |
| 9:18.201 | במי |  |
| 9:18.202 | ששונה |  |
| 9:18.203 | הלכות |  |
| 9:18.204 | כתרה |  |
| 9:18.205 | של |  |
| 9:18.206 | תורה |  |
| 9:18.207 | ע"כ | . |
| 9:18.208 | כי |  |
| 9:18.209 | לפי |  |
| 9:18.210 | דברי |  |
| 9:18.211 | הזהר |  |
| 9:18.212 | והמקובלים |  |
| 9:18.213 | החדשים |  |
| 9:18.214 | הוא |  |
| 9:18.215 | תבנא |  |
| 9:18.216 | דאוריתא | . |
| 9:18.217 | כמוץ |  |
| 9:18.218 | ותבן |  |
| 9:18.219 | שזורין |  |
| 9:18.220 | ובוררין |  |
| 9:18.221 | ומפרישין |  |
| 9:18.222 | אותו |  |
| 9:18.223 | להשליכו |  |
| 9:18.224 | חוצה | ! |
| 9:18.225 | ולא |  |
| 9:18.226 | די |  |
| 9:18.227 | זה |  |
| 9:18.228 | אלא |  |
| 9:18.229 | שצריך |  |
| 9:18.230 | העוסק |  |
| 9:18.231 | בהם |  |
| 9:18.232 | לשוב |  |
| 9:18.233 | מחטאתו |  |
| 9:18.234 | ולעזוב |  |
| 9:18.235 | אותם |  |
| 9:18.236 | ולעסוק |  |
| 9:18.237 | בנסתרות |  |
| 9:18.238 | הזהר |  |
| 9:18.239 | וסודותיו | . |
| 9:18.240 | ואך |  |
| 9:18.241 | אז |  |
| 9:18.242 | יזכה |  |
| 9:18.243 | ליאור |  |
| 9:18.244 | באור |  |
| 9:18.245 | החיים | . |
| 9:18.246 | את |  |
| 9:18.247 | כל |  |
| 9:18.248 | אלה |  |
| 9:18.249 | הרעות |  |
| 9:18.250 | עשה |  |
| 9:18.251 | הפילסוף |  |
| 9:18.252 | המסית |  |
| 9:18.253 | מחבר |  |
| 9:18.254 | הזהר |  |
| 9:18.255 | להבזות |  |
| 9:18.256 | את |  |
| 9:18.257 | המשנה |  |
| 9:18.258 | והתלמוד |  |
| 9:18.259 | הקדוש |  |
| 9:18.260 | המקובל |  |
| 9:18.261 | ממשה |  |
| 9:18.262 | איש |  |
| 9:18.263 | מפי |  |
| 9:18.264 | איש |  |
| 9:18.265 | ולחלל |  |
| 9:18.266 | את |  |
| 9:18.267 | כבדן |  |
| 9:18.268 | כמו |  |
| 9:18.269 | שנתבאר |  |
| 9:18.270 | לעיל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

