# Translation pack: milhamot-hashem perek-2-16 (Milchamot ha-Shem 2:32–2:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-16`.

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
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **ז"ל** → "of blessed memory" (honorific).
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **רש"י** → "Rashi" (authority, transliterated; note).
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
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

- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- אל **אחר** → "another god".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **אחי** → "my brother" (second person kept).
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
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
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


## Milchamot ha-Shem 2:32

Hebrew (Wikisource): והנה אהובי כאחי ז"ל ממשל ארץ ישראל הנמשל, כי המהלך הרב אל מקום השם בהשגת הסוד הוא קוצר המשיג ועומק המושג, ורבוי ההצעות שהם סבות מונעות בכל השגת חכמה, ובפרט בחכמת הקבלה מכמה פנים, ועל כולם החטא שהוא מוסך ומבדיל בין השכל והמושכל וכו' והנה אף כי בעל פרדס רמונים לגודל קדושתו וחסידותו נזהר מאד בכמה מקומות ומזהיר מאד על מחשבה זרה שלפירוד וגבול וממשל משלים וכו' ואף אנו נאמר מי יתן ויערה עלינו רוח ממרום שנשכיל ונבין הדברים בלי נטייה מן האמת ומשרשי הדת וכו'. ובאלו נקוט האיי כללא בידך שכל דברים האסורים בספרי המקובלים שזכרו בלי מחלוקת בדבר ובלי חקירה, ויובנו פשטן של דברים וחיבור הדברים קשרם על לוח לבן ויהיו לטטפות בין עיניך. ואף כי יקרא לפניך דברים הנראים תמוהים מצד עצמן או מצד שנראה חלילה שיש פירוד ומספר וגשמות ושינוי, תקבלם באמונת אומן, ותאמר בלבבך כי קצרה דעתך וידיעתך בעומק הבנה מבלי שתסתפק ח"ו באחד משרשי הדת האמורים ביגדל ובכמה דוכתי, שהרי גם אותם העיקרי אי אפשר לירד לשרשם מצד שכל אנושי וכו'. ומעתה צא ולמד שאסור לחדש ולהבין דבר מתוך הנכתב בספר פן לא יכוין אל האמת, אמנם לבאר פסוקים על פי מה שאמרו המקובלים שהקדמה מוסכמת רק שעל זה מבאר הפ' או אומר שנרמז בו על פי הגימטריא או ראשי וסופי תיבות. טוב ויפה הוא שרק יאמרנו בדרך אפשר לא להחליט הדבר וכו' רק לא יכלה זמנו בהם, ואשר נתן ה' בלבבו והוא מוטעם נכון בלי דוחק אל ישליך אחרי גיוו רק יכתבנו לזכרון ויהיה בעיניו במוצא שלל רב, כי הוא סימן לנשמתו חלק אלוה ממעל שיש לו יניקה ומבוא במקום גבוה, וקדוש יאמר לו, ועתה הנה גליתי לך מצפוני לבי. ואתה המלך בעיונך ודעתך הקדושה והרחבה הודיעני נא את דרכיך, ואבטל דעתי מפני דעתך. יאיר חיים בכרך.

| id | word | punctuation after |
|---|---|---|
| 2:32.1 | והנה |  |
| 2:32.2 | אהובי |  |
| 2:32.3 | כאחי |  |
| 2:32.4 | ז"ל |  |
| 2:32.5 | ממשל |  |
| 2:32.6 | ארץ |  |
| 2:32.7 | ישראל |  |
| 2:32.8 | הנמשל | , |
| 2:32.9 | כי |  |
| 2:32.10 | המהלך |  |
| 2:32.11 | הרב |  |
| 2:32.12 | אל |  |
| 2:32.13 | מקום |  |
| 2:32.14 | השם |  |
| 2:32.15 | בהשגת |  |
| 2:32.16 | הסוד |  |
| 2:32.17 | הוא |  |
| 2:32.18 | קוצר |  |
| 2:32.19 | המשיג |  |
| 2:32.20 | ועומק |  |
| 2:32.21 | המושג | , |
| 2:32.22 | ורבוי |  |
| 2:32.23 | ההצעות |  |
| 2:32.24 | שהם |  |
| 2:32.25 | סבות |  |
| 2:32.26 | מונעות |  |
| 2:32.27 | בכל |  |
| 2:32.28 | השגת |  |
| 2:32.29 | חכמה | , |
| 2:32.30 | ובפרט |  |
| 2:32.31 | בחכמת |  |
| 2:32.32 | הקבלה |  |
| 2:32.33 | מכמה |  |
| 2:32.34 | פנים | , |
| 2:32.35 | ועל |  |
| 2:32.36 | כולם |  |
| 2:32.37 | החטא |  |
| 2:32.38 | שהוא |  |
| 2:32.39 | מוסך |  |
| 2:32.40 | ומבדיל |  |
| 2:32.41 | בין |  |
| 2:32.42 | השכל |  |
| 2:32.43 | והמושכל |  |
| 2:32.44 | וכו | ' |
| 2:32.45 | והנה |  |
| 2:32.46 | אף |  |
| 2:32.47 | כי |  |
| 2:32.48 | בעל |  |
| 2:32.49 | פרדס |  |
| 2:32.50 | רמונים |  |
| 2:32.51 | לגודל |  |
| 2:32.52 | קדושתו |  |
| 2:32.53 | וחסידותו |  |
| 2:32.54 | נזהר |  |
| 2:32.55 | מאד |  |
| 2:32.56 | בכמה |  |
| 2:32.57 | מקומות |  |
| 2:32.58 | ומזהיר |  |
| 2:32.59 | מאד |  |
| 2:32.60 | על |  |
| 2:32.61 | מחשבה |  |
| 2:32.62 | זרה |  |
| 2:32.63 | שלפירוד |  |
| 2:32.64 | וגבול |  |
| 2:32.65 | וממשל |  |
| 2:32.66 | משלים |  |
| 2:32.67 | וכו | ' |
| 2:32.68 | ואף |  |
| 2:32.69 | אנו |  |
| 2:32.70 | נאמר |  |
| 2:32.71 | מי |  |
| 2:32.72 | יתן |  |
| 2:32.73 | ויערה |  |
| 2:32.74 | עלינו |  |
| 2:32.75 | רוח |  |
| 2:32.76 | ממרום |  |
| 2:32.77 | שנשכיל |  |
| 2:32.78 | ונבין |  |
| 2:32.79 | הדברים |  |
| 2:32.80 | בלי |  |
| 2:32.81 | נטייה |  |
| 2:32.82 | מן |  |
| 2:32.83 | האמת |  |
| 2:32.84 | ומשרשי |  |
| 2:32.85 | הדת |  |
| 2:32.86 | וכו | '. |
| 2:32.87 | ובאלו |  |
| 2:32.88 | נקוט |  |
| 2:32.89 | האיי |  |
| 2:32.90 | כללא |  |
| 2:32.91 | בידך |  |
| 2:32.92 | שכל |  |
| 2:32.93 | דברים |  |
| 2:32.94 | האסורים |  |
| 2:32.95 | בספרי |  |
| 2:32.96 | המקובלים |  |
| 2:32.97 | שזכרו |  |
| 2:32.98 | בלי |  |
| 2:32.99 | מחלוקת |  |
| 2:32.100 | בדבר |  |
| 2:32.101 | ובלי |  |
| 2:32.102 | חקירה | , |
| 2:32.103 | ויובנו |  |
| 2:32.104 | פשטן |  |
| 2:32.105 | של |  |
| 2:32.106 | דברים |  |
| 2:32.107 | וחיבור |  |
| 2:32.108 | הדברים |  |
| 2:32.109 | קשרם |  |
| 2:32.110 | על |  |
| 2:32.111 | לוח |  |
| 2:32.112 | לבן |  |
| 2:32.113 | ויהיו |  |
| 2:32.114 | לטטפות |  |
| 2:32.115 | בין |  |
| 2:32.116 | עיניך | . |
| 2:32.117 | ואף |  |
| 2:32.118 | כי |  |
| 2:32.119 | יקרא |  |
| 2:32.120 | לפניך |  |
| 2:32.121 | דברים |  |
| 2:32.122 | הנראים |  |
| 2:32.123 | תמוהים |  |
| 2:32.124 | מצד |  |
| 2:32.125 | עצמן |  |
| 2:32.126 | או |  |
| 2:32.127 | מצד |  |
| 2:32.128 | שנראה |  |
| 2:32.129 | חלילה |  |
| 2:32.130 | שיש |  |
| 2:32.131 | פירוד |  |
| 2:32.132 | ומספר |  |
| 2:32.133 | וגשמות |  |
| 2:32.134 | ושינוי | , |
| 2:32.135 | תקבלם |  |
| 2:32.136 | באמונת |  |
| 2:32.137 | אומן | , |
| 2:32.138 | ותאמר |  |
| 2:32.139 | בלבבך |  |
| 2:32.140 | כי |  |
| 2:32.141 | קצרה |  |
| 2:32.142 | דעתך |  |
| 2:32.143 | וידיעתך |  |
| 2:32.144 | בעומק |  |
| 2:32.145 | הבנה |  |
| 2:32.146 | מבלי |  |
| 2:32.147 | שתסתפק |  |
| 2:32.148 | ח"ו |  |
| 2:32.149 | באחד |  |
| 2:32.150 | משרשי |  |
| 2:32.151 | הדת |  |
| 2:32.152 | האמורים |  |
| 2:32.153 | ביגדל |  |
| 2:32.154 | ובכמה |  |
| 2:32.155 | דוכתי | , |
| 2:32.156 | שהרי |  |
| 2:32.157 | גם |  |
| 2:32.158 | אותם |  |
| 2:32.159 | העיקרי |  |
| 2:32.160 | אי |  |
| 2:32.161 | אפשר |  |
| 2:32.162 | לירד |  |
| 2:32.163 | לשרשם |  |
| 2:32.164 | מצד |  |
| 2:32.165 | שכל |  |
| 2:32.166 | אנושי |  |
| 2:32.167 | וכו | '. |
| 2:32.168 | ומעתה |  |
| 2:32.169 | צא |  |
| 2:32.170 | ולמד |  |
| 2:32.171 | שאסור |  |
| 2:32.172 | לחדש |  |
| 2:32.173 | ולהבין |  |
| 2:32.174 | דבר |  |
| 2:32.175 | מתוך |  |
| 2:32.176 | הנכתב |  |
| 2:32.177 | בספר |  |
| 2:32.178 | פן |  |
| 2:32.179 | לא |  |
| 2:32.180 | יכוין |  |
| 2:32.181 | אל |  |
| 2:32.182 | האמת | , |
| 2:32.183 | אמנם |  |
| 2:32.184 | לבאר |  |
| 2:32.185 | פסוקים |  |
| 2:32.186 | על |  |
| 2:32.187 | פי |  |
| 2:32.188 | מה |  |
| 2:32.189 | שאמרו |  |
| 2:32.190 | המקובלים |  |
| 2:32.191 | שהקדמה |  |
| 2:32.192 | מוסכמת |  |
| 2:32.193 | רק |  |
| 2:32.194 | שעל |  |
| 2:32.195 | זה |  |
| 2:32.196 | מבאר |  |
| 2:32.197 | הפ | ' |
| 2:32.198 | או |  |
| 2:32.199 | אומר |  |
| 2:32.200 | שנרמז |  |
| 2:32.201 | בו |  |
| 2:32.202 | על |  |
| 2:32.203 | פי |  |
| 2:32.204 | הגימטריא |  |
| 2:32.205 | או |  |
| 2:32.206 | ראשי |  |
| 2:32.207 | וסופי |  |
| 2:32.208 | תיבות | . |
| 2:32.209 | טוב |  |
| 2:32.210 | ויפה |  |
| 2:32.211 | הוא |  |
| 2:32.212 | שרק |  |
| 2:32.213 | יאמרנו |  |
| 2:32.214 | בדרך |  |
| 2:32.215 | אפשר |  |
| 2:32.216 | לא |  |
| 2:32.217 | להחליט |  |
| 2:32.218 | הדבר |  |
| 2:32.219 | וכו | ' |
| 2:32.220 | רק |  |
| 2:32.221 | לא |  |
| 2:32.222 | יכלה |  |
| 2:32.223 | זמנו |  |
| 2:32.224 | בהם | , |
| 2:32.225 | ואשר |  |
| 2:32.226 | נתן |  |
| 2:32.227 | ה | ' |
| 2:32.228 | בלבבו |  |
| 2:32.229 | והוא |  |
| 2:32.230 | מוטעם |  |
| 2:32.231 | נכון |  |
| 2:32.232 | בלי |  |
| 2:32.233 | דוחק |  |
| 2:32.234 | אל |  |
| 2:32.235 | ישליך |  |
| 2:32.236 | אחרי |  |
| 2:32.237 | גיוו |  |
| 2:32.238 | רק |  |
| 2:32.239 | יכתבנו |  |
| 2:32.240 | לזכרון |  |
| 2:32.241 | ויהיה |  |
| 2:32.242 | בעיניו |  |
| 2:32.243 | במוצא |  |
| 2:32.244 | שלל |  |
| 2:32.245 | רב | , |
| 2:32.246 | כי |  |
| 2:32.247 | הוא |  |
| 2:32.248 | סימן |  |
| 2:32.249 | לנשמתו |  |
| 2:32.250 | חלק |  |
| 2:32.251 | אלוה |  |
| 2:32.252 | ממעל |  |
| 2:32.253 | שיש |  |
| 2:32.254 | לו |  |
| 2:32.255 | יניקה |  |
| 2:32.256 | ומבוא |  |
| 2:32.257 | במקום |  |
| 2:32.258 | גבוה | , |
| 2:32.259 | וקדוש |  |
| 2:32.260 | יאמר |  |
| 2:32.261 | לו | , |
| 2:32.262 | ועתה |  |
| 2:32.263 | הנה |  |
| 2:32.264 | גליתי |  |
| 2:32.265 | לך |  |
| 2:32.266 | מצפוני |  |
| 2:32.267 | לבי | . |
| 2:32.268 | ואתה |  |
| 2:32.269 | המלך |  |
| 2:32.270 | בעיונך |  |
| 2:32.271 | ודעתך |  |
| 2:32.272 | הקדושה |  |
| 2:32.273 | והרחבה |  |
| 2:32.274 | הודיעני |  |
| 2:32.275 | נא |  |
| 2:32.276 | את |  |
| 2:32.277 | דרכיך | , |
| 2:32.278 | ואבטל |  |
| 2:32.279 | דעתי |  |
| 2:32.280 | מפני |  |
| 2:32.281 | דעתך | . |
| 2:32.282 | יאיר |  |
| 2:32.283 | חיים |  |
| 2:32.284 | בכרך | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

