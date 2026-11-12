# Review pack: milhamot-hashem perek-2-16 (Milchamot ha-Shem 2:32–2:32)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-16`.

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


### Translation under review (2:32)

English: And behold, my beloved, like my brother of blessed memory, the ruler of the land of Israel, the Nimshal, for the going far toward the place of the Name in the grasp of the secret is the shortness of the one who grasps and the depth of what is grasped, and the multitude of propositions, which are preventing causes in every grasp of wisdom, and especially in the wisdom of Kabbalah, in several respects, and over all of them is the sin, which obstructs and separates between the intellect and the intelligible, etc., and indeed, even though the author of Pardes Rimonim, because of the greatness of his holiness and his piety, is very careful in many places, and warns greatly against a strange thought that is of division, and limit, and dominion of similitudes, etc., we too shall say: would that a spirit from on high be poured out upon us, that we might grasp and understand the matters without inclining from the truth and from the roots of the religion, etc. And by these, take this rule in your hand: all things that are forbidden in the books of the Kabbalists, which they mention without dispute on the matter and without inquiry, and their plain sense will be understood, and the joining of the things: bind them on a white tablet, and let them be frontlets between your eyes. And even when things are read before you that appear astonishing in themselves, or from the side that it seems, God forbid, that there is division, number, corporeality and change, you shall accept them in the faith of an artisan, and say in your heart that your mind is short, and your knowledge in the depth of understanding, without your doubting, God forbid, one of the roots of the religion that are stated in Yigdal and in several places, for even those fundamental ones, it is not possible to descend to their roots by human reason, etc. And now go and learn that it is forbidden to innovate and to understand a matter from what is written in the book, lest he not aim at the truth, however, to explain verses according to what the Kabbalists said, that the premise is agreed; only that on this he explains the verse, or says that there is an allusion in it, by means of gematria, or the initial and final letters of words. It is good and fine that he only say it as a possibility, without deciding the matter, etc., only let him not waste his time on them, and whatever the LORD has put in his heart, and it is correctly tasted without strain, he should not cast it behind his back, but rather let him write it down for memory, and let it be in his eyes as one who finds great spoil, for it is a sign for his soul, a portion of God from on high, which has nourishment and access in a high place, and a holy one will say to him: and now, behold, I have revealed to you the secrets of my heart. And you, the king, in your reflection and your knowledge, holy and broad, tell me, please, your ways, and I will set aside my opinion before your opinion. Yair Chayyim Bakhrakh.

Literal: and-behold my-beloved like-my-brother of-blessed-memory the-ruler-of the-land-of Israel the-Nimshal for the-going far toward the-place-of the-Name in-the-grasp-of the-secret is the-shortness-of the-one-who-grasps and-the-depth-of what-is-grasped and-the-multitude-of the-propositions which-are causes preventing in-every grasp-of wisdom and-especially in-the-wisdom-of Kabbalah in-several respects and-over all-of-them the-sin which-is obstructing and-separating between the-intellect and-the-intelligible etc. and-indeed even though the-author-of Pardes Rimonim because-of-the-greatness-of his-holiness and-his-piety is-careful very in-many places and-warns greatly against a-thought strange that-is-of-division and-limit and-dominion of-similitudes etc. and-even we shall-say who will-give and-be-poured-out upon-us a-spirit from-on-high that-we-might-grasp and-understand the-matters without inclining from the-truth and-from-the-roots-of the-religion etc. and-by-these take this rule in-your-hand all things that-are-forbidden in-the-books-of the-Kabbalists which-they-mention without dispute on-the-matter and-without inquiry will-be-understood their-plain-sense of the-things and-the-joining-of the-things bind-them on a-tablet white and-let-them-be frontlets between your-eyes and-even when are-read before-you things that-appear astonishing from-the-side-of themselves or from-the-side-of that-it-seems God-forbid that-there-is division and-number and-corporeality and-change you-shall-accept-them in-the-faith-of an-artisan and-say in-your-heart that is-short your-mind and-your-knowledge in-the-depth-of understanding without your-doubting God-forbid one-of the-roots-of the-religion that-are-stated in-Yigdal and-in-several places for even those the-fundamental-ones it-is-not possible to-descend to-their-roots from-the-side-of reason human etc. and-now go and-learn that-it-is-forbidden to-innovate and-to-understand a-matter from-within what-is-written in-the-book lest not aim at the-truth however to-explain verses by the-mouth-of what they-said the-Kabbalists that-the-premise is-agreed only that-on this he-explains the-verse or says that-there-is-an-allusion in-it by means-of gematria or the-initials-of and-the-finals-of words it-is-good and-fine it-is that-only he-say-it as-a-way a-possibility without deciding the-matter etc. only not let-him-waste his-time on-them and-whatever has-put the-LORD in-his-heart and-it is-correctly-tasted rightly without strain he-should-not cast-it behind his-back but-rather let-him-write-it for-memory and-let-it-be in-his-eyes as-one-who-finds spoil great for it is-a-sign for-his-soul a-portion of-God from-on-high which-has to-it nourishment and-access in-a-place high and-a-holy-one will-say to-him and-now behold I-have-revealed to-you the-secrets-of my-heart and-you the-king in-your-reflection and-your-knowledge holy and-broad tell-me please (object-marker) your-ways and-I-will-set-aside my-opinion before your-opinion Yair Chayyim Bakhrakh

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:32.1 | and behold |  |  | והנה: 'and behold'; at 2:32.45 the same word opens the next clause, 'and indeed'. |
| 2:32.2 | my beloved |  |  |  |
| 2:32.3 | like my brother |  |  |  |
| 2:32.4 | of blessed memory |  |  |  |
| 2:32.5 | the ruler of |  |  | ממשל: a noun 'rule' or the word mashal (parable). The epithet plays on mashal and nimshal. Read as 'ruler' here; the parable reading is in readings. |
| 2:32.6 | the land of |  |  |  |
| 2:32.7 | Israel |  |  |  |
| 2:32.8 | the Nimshal |  |  | הנמשל: the referent of a parable (nimshal). The epithet pairs it with ממשל. |
| 2:32.9 | for |  |  |  |
| 2:32.10 | the going |  |  |  |
| 2:32.11 | far |  |  | הרב: 'much, many', here 'far'. |
| 2:32.12 | toward |  |  |  |
| 2:32.13 | the place of |  |  |  |
| 2:32.14 | the Name |  |  |  |
| 2:32.15 | in the grasp of |  |  |  |
| 2:32.16 | the secret |  |  |  |
| 2:32.17 | is |  |  |  |
| 2:32.18 | the shortness of |  |  |  |
| 2:32.19 | the one who grasps |  |  |  |
| 2:32.20 | and the depth of |  |  |  |
| 2:32.21 | what is grasped |  |  |  |
| 2:32.22 | and the multitude of |  |  |  |
| 2:32.23 | the propositions |  |  | הצעות: propositions or premises put forward for the mind. |
| 2:32.24 | which are |  |  |  |
| 2:32.25 | causes |  |  | סבות: causes. The following adjective, מונעות, is 'preventing'. |
| 2:32.26 | preventing |  |  |  |
| 2:32.27 | in every |  |  |  |
| 2:32.28 | grasp of |  |  |  |
| 2:32.29 | wisdom |  |  |  |
| 2:32.30 | and especially |  |  |  |
| 2:32.31 | in the wisdom of |  |  |  |
| 2:32.32 | Kabbalah |  |  |  |
| 2:32.33 | in several |  |  |  |
| 2:32.34 | respects |  |  | פנים: literally 'faces'; here 'respects' (aspects). |
| 2:32.35 | and over |  |  |  |
| 2:32.36 | all of them |  |  |  |
| 2:32.37 | the sin |  |  |  |
| 2:32.38 | which is |  |  |  |
| 2:32.39 | obstructing |  |  | מוסך: obstructs, stands in the way. |
| 2:32.40 | and separating |  |  |  |
| 2:32.41 | between |  |  |  |
| 2:32.42 | the intellect |  |  |  |
| 2:32.43 | and the intelligible |  |  |  |
| 2:32.44 | etc. |  |  | וכו': 'etc.' marks an omitted continuation in the author's own text. |
| 2:32.45 | and indeed |  |  |  |
| 2:32.46 | even |  |  |  |
| 2:32.47 | though |  |  |  |
| 2:32.48 | the author of |  |  |  |
| 2:32.49 | Pardes |  |  | פרדס רמונים: Pardes Rimonim, 'Pomegranate Orchard', the title of a Kabbalistic book. בעל: 'the author of'. The text does not name him. |
| 2:32.50 | Rimonim |  |  |  |
| 2:32.51 | because of the greatness of |  |  |  |
| 2:32.52 | his holiness |  |  |  |
| 2:32.53 | and his piety |  |  |  |
| 2:32.54 | is careful |  |  |  |
| 2:32.55 | very |  |  |  |
| 2:32.56 | in many |  |  |  |
| 2:32.57 | places |  |  |  |
| 2:32.58 | and warns |  |  |  |
| 2:32.59 | greatly |  |  |  |
| 2:32.60 | against |  |  |  |
| 2:32.61 | a thought |  |  |  |
| 2:32.62 | strange |  |  | זרה: 'strange, foreign'. |
| 2:32.63 | that is of division |  |  | פירוד: division, separation. |
| 2:32.64 | and limit |  |  |  |
| 2:32.65 | and dominion |  |  |  |
| 2:32.66 | of similitudes |  |  |  |
| 2:32.67 | etc. |  |  | וכו': 'etc.' as at 2:32.44. |
| 2:32.68 | and even |  |  |  |
| 2:32.69 | we |  |  |  |
| 2:32.70 | shall say |  |  |  |
| 2:32.71 | who |  |  | מי יתן: an optative idiom, literally 'who will give', rendered 'would that'. |
| 2:32.72 | will give |  |  |  |
| 2:32.73 | and be poured out |  |  | ויערה: 'and be poured out', a yiqtol with wish sense. Allusion to Isa 32:15 (apparatus). |
| 2:32.74 | upon us |  |  |  |
| 2:32.75 | a spirit |  |  |  |
| 2:32.76 | from on high |  |  |  |
| 2:32.77 | that we might grasp |  |  |  |
| 2:32.78 | and understand |  |  |  |
| 2:32.79 | the matters |  |  |  |
| 2:32.80 | without |  |  |  |
| 2:32.81 | inclining |  |  |  |
| 2:32.82 | from |  |  |  |
| 2:32.83 | the truth |  |  |  |
| 2:32.84 | and from the roots of |  |  |  |
| 2:32.85 | the religion |  |  |  |
| 2:32.86 | etc. |  |  | וכו': 'etc.', and the full stop follows in the Hebrew. |
| 2:32.87 | and by these |  |  |  |
| 2:32.88 | take |  |  |  |
| 2:32.89 | this |  |  | האיי: the spelling of האי, 'this'. |
| 2:32.90 | rule |  |  | כללא: the Aramaic word for 'rule' (kelala). |
| 2:32.91 | in your hand |  |  |  |
| 2:32.92 | all |  |  | שכל: the Aramaic word for 'all' (kol). |
| 2:32.93 | things |  |  |  |
| 2:32.94 | that are forbidden |  |  |  |
| 2:32.95 | in the books of |  |  |  |
| 2:32.96 | the Kabbalists |  |  |  |
| 2:32.97 | which they mention |  |  |  |
| 2:32.98 | without |  |  |  |
| 2:32.99 | dispute |  |  |  |
| 2:32.100 | on the matter |  |  |  |
| 2:32.101 | and without |  |  |  |
| 2:32.102 | inquiry |  |  |  |
| 2:32.103 | will be understood |  |  | ויובנו: passive, 'will be understood'. |
| 2:32.104 | their plain sense |  |  |  |
| 2:32.105 | of |  |  |  |
| 2:32.106 | the things |  |  |  |
| 2:32.107 | and the joining of |  |  |  |
| 2:32.108 | the things |  |  |  |
| 2:32.109 | bind them |  |  | קשרם: imperative, 'bind them'. |
| 2:32.110 | on |  |  |  |
| 2:32.111 | a tablet |  |  |  |
| 2:32.112 | white |  |  | לבן: literally 'white'. Prov 3:3 reads 'the tablet of your heart' (see apparatus); translated as written. |
| 2:32.113 | and let them be |  |  | ויהיו: 'and let them be', a jussive. |
| 2:32.114 | frontlets |  |  |  |
| 2:32.115 | between |  |  |  |
| 2:32.116 | your eyes |  |  |  |
| 2:32.117 | and even |  |  |  |
| 2:32.118 | when |  |  |  |
| 2:32.119 | are read |  |  | יקרא: 'is read', passive. |
| 2:32.120 | before you |  |  |  |
| 2:32.121 | things |  |  |  |
| 2:32.122 | that appear |  |  |  |
| 2:32.123 | astonishing |  |  |  |
| 2:32.124 | from the side of |  |  |  |
| 2:32.125 | themselves |  |  |  |
| 2:32.126 | or |  |  |  |
| 2:32.127 | from the side of |  |  |  |
| 2:32.128 | that it seems |  |  |  |
| 2:32.129 | God forbid |  |  | חלילה: an oath formula, 'far be it'; rendered 'God forbid' as at 2:32.148. |
| 2:32.130 | that there is |  |  |  |
| 2:32.131 | division |  |  |  |
| 2:32.132 | and number |  |  |  |
| 2:32.133 | and corporeality |  |  |  |
| 2:32.134 | and change |  |  |  |
| 2:32.135 | you shall accept them |  |  |  |
| 2:32.136 | in the faith of |  |  | באמונת אומן: literally 'with the faith of an artisan'; here, accepting on trust without proof. |
| 2:32.137 | an artisan |  |  |  |
| 2:32.138 | and say |  |  |  |
| 2:32.139 | in your heart |  |  |  |
| 2:32.140 | that |  |  |  |
| 2:32.141 | is short |  |  |  |
| 2:32.142 | your mind |  |  |  |
| 2:32.143 | and your knowledge |  |  |  |
| 2:32.144 | in the depth of |  |  |  |
| 2:32.145 | understanding |  |  |  |
| 2:32.146 | without |  |  |  |
| 2:32.147 | your doubting |  |  |  |
| 2:32.148 | God forbid |  |  |  |
| 2:32.149 | one of |  |  |  |
| 2:32.150 | the roots of |  |  |  |
| 2:32.151 | the religion |  |  |  |
| 2:32.152 | that are stated |  |  |  |
| 2:32.153 | in Yigdal |  |  | ביגדל: Yigdal, the hymn that lists the principles of faith. |
| 2:32.154 | and in several |  |  |  |
| 2:32.155 | places |  |  |  |
| 2:32.156 | for |  |  |  |
| 2:32.157 | even |  |  |  |
| 2:32.158 | those |  |  |  |
| 2:32.159 | the fundamental ones |  |  | העיקרי: 'the fundamental ones', the principles. |
| 2:32.160 | it is not |  |  |  |
| 2:32.161 | possible |  |  |  |
| 2:32.162 | to descend |  |  |  |
| 2:32.163 | to their roots |  |  |  |
| 2:32.164 | from the side of |  |  |  |
| 2:32.165 | reason |  |  |  |
| 2:32.166 | human |  |  |  |
| 2:32.167 | etc. |  |  |  |
| 2:32.168 | and now |  |  |  |
| 2:32.169 | go |  |  |  |
| 2:32.170 | and learn |  |  |  |
| 2:32.171 | that it is forbidden |  |  |  |
| 2:32.172 | to innovate |  |  |  |
| 2:32.173 | and to understand |  |  |  |
| 2:32.174 | a matter |  |  |  |
| 2:32.175 | from within |  |  |  |
| 2:32.176 | what is written |  |  | הנכתב: passive participle, 'what is written'. |
| 2:32.177 | in the book |  |  |  |
| 2:32.178 | lest |  |  |  |
| 2:32.179 | not |  |  |  |
| 2:32.180 | aim |  |  |  |
| 2:32.181 | at |  |  |  |
| 2:32.182 | the truth |  |  |  |
| 2:32.183 | however |  |  |  |
| 2:32.184 | to explain |  |  |  |
| 2:32.185 | verses |  |  |  |
| 2:32.186 | by |  |  |  |
| 2:32.187 | the mouth of |  |  | פי: literally 'mouth of'; with על, the idiom 'according to'. |
| 2:32.188 | what |  |  |  |
| 2:32.189 | they said |  |  |  |
| 2:32.190 | the Kabbalists |  |  |  |
| 2:32.191 | that the premise |  |  |  |
| 2:32.192 | is agreed |  |  |  |
| 2:32.193 | only |  |  |  |
| 2:32.194 | that on |  |  |  |
| 2:32.195 | this |  |  |  |
| 2:32.196 | he explains |  |  |  |
| 2:32.197 | the verse |  |  | הפ': the verse (Hebrew פסוק). The author does not name the verse. |
| 2:32.198 | or |  |  |  |
| 2:32.199 | says |  |  |  |
| 2:32.200 | that there is an allusion |  |  | שנרמז: 'is alluded to', from remez. |
| 2:32.201 | in it |  |  |  |
| 2:32.202 | by |  |  |  |
| 2:32.203 | means of |  |  | על פי: 'by means of' with gematria. |
| 2:32.204 | gematria |  |  |  |
| 2:32.205 | or |  |  |  |
| 2:32.206 | the initials of |  |  | ראשי: 'the initials of'. |
| 2:32.207 | and the finals of |  |  |  |
| 2:32.208 | words |  |  | תיבות: words. |
| 2:32.209 | it is good |  |  |  |
| 2:32.210 | and fine |  |  |  |
| 2:32.211 | it is |  |  |  |
| 2:32.212 | that only |  |  |  |
| 2:32.213 | he say it |  |  |  |
| 2:32.214 | as a way |  |  |  |
| 2:32.215 | a possibility |  |  | אפשר: 'a possibility'. |
| 2:32.216 | without |  |  |  |
| 2:32.217 | deciding |  |  |  |
| 2:32.218 | the matter |  |  |  |
| 2:32.219 | etc. |  |  |  |
| 2:32.220 | only |  |  |  |
| 2:32.221 | not |  |  |  |
| 2:32.222 | let him waste |  |  | יכלה: 'waste, use up'. |
| 2:32.223 | his time |  |  |  |
| 2:32.224 | on them |  |  |  |
| 2:32.225 | and whatever |  |  |  |
| 2:32.226 | has put |  |  |  |
| 2:32.227 | the LORD |  |  | ה': the Tetragrammaton as written, here the LORD. |
| 2:32.228 | in his heart |  |  |  |
| 2:32.229 | and it |  |  |  |
| 2:32.230 | is correctly tasted |  |  |  |
| 2:32.231 | rightly |  |  |  |
| 2:32.232 | without |  |  |  |
| 2:32.233 | strain |  |  |  |
| 2:32.234 | he should not |  |  |  |
| 2:32.235 | cast it |  |  |  |
| 2:32.236 | behind |  |  |  |
| 2:32.237 | his back |  |  | גיוו: literally 'his back' or 'his inner part'. 'Behind his back' is the idiom for ignoring something. See apparatus. |
| 2:32.238 | but rather |  |  |  |
| 2:32.239 | let him write it |  |  |  |
| 2:32.240 | for memory |  |  |  |
| 2:32.241 | and let it be |  |  |  |
| 2:32.242 | in his eyes |  |  |  |
| 2:32.243 | as one who finds |  |  | במוצא: 'as one who finds'. Allusion to Ps 119:162 (apparatus). |
| 2:32.244 | spoil |  |  |  |
| 2:32.245 | great |  |  |  |
| 2:32.246 | for |  |  |  |
| 2:32.247 | it |  |  |  |
| 2:32.248 | is a sign |  |  |  |
| 2:32.249 | for his soul |  |  |  |
| 2:32.250 | a portion |  |  |  |
| 2:32.251 | of God |  |  | אלוה: 'God' (singular). Allusion to Job 31:2 (apparatus). |
| 2:32.252 | from on high |  |  |  |
| 2:32.253 | which has |  |  |  |
| 2:32.254 | to it |  |  |  |
| 2:32.255 | nourishment |  |  |  |
| 2:32.256 | and access |  |  |  |
| 2:32.257 | in a place |  |  |  |
| 2:32.258 | high |  |  |  |
| 2:32.259 | and a holy one |  |  |  |
| 2:32.260 | will say |  |  |  |
| 2:32.261 | to him |  |  |  |
| 2:32.262 | and now |  |  |  |
| 2:32.263 | behold |  |  |  |
| 2:32.264 | I have revealed |  |  |  |
| 2:32.265 | to you |  |  |  |
| 2:32.266 | the secrets of |  |  |  |
| 2:32.267 | my heart |  |  |  |
| 2:32.268 | and you |  |  |  |
| 2:32.269 | the king |  |  |  |
| 2:32.270 | in your reflection |  |  |  |
| 2:32.271 | and your knowledge |  |  |  |
| 2:32.272 | holy |  |  |  |
| 2:32.273 | and broad |  |  |  |
| 2:32.274 | tell me |  |  |  |
| 2:32.275 | please |  |  |  |
| 2:32.276 | (object marker) |  |  | את: the object marker, with no English counterpart. |
| 2:32.277 | your ways |  |  |  |
| 2:32.278 | and I will set aside |  |  |  |
| 2:32.279 | my opinion |  |  |  |
| 2:32.280 | before |  |  |  |
| 2:32.281 | your opinion |  |  |  |
| 2:32.282 | Yair |  |  |  |
| 2:32.283 | Chayyim |  |  |  |
| 2:32.284 | Bakhrakh |  |  | בכרך: the surname is unvocalized in the text; the English is a transliteration and its vowels are a guess. |

Phrases:
- 2:32.1 2:32.2 2:32.3 2:32.4 2:32.5 2:32.6 2:32.7 2:32.8 | And behold, my beloved, like my brother of blessed memory, the ruler of the land of Israel, the Nimshal, | and-behold my-beloved like-my-brother of-blessed-memory the-ruler-of the-land-of Israel the-Nimshal
- 2:32.9 2:32.10 2:32.11 2:32.12 2:32.13 2:32.14 2:32.15 2:32.16 2:32.17 | for the going far toward the place of the Name in the grasp of the secret is | for the-going far toward the-place-of the-Name in-the-grasp-of the-secret is
- 2:32.18 2:32.19 2:32.20 2:32.21 | the shortness of the one who grasps and the depth of what is grasped, | the-shortness-of the-one-who-grasps and-the-depth-of what-is-grasped
- 2:32.22 2:32.23 2:32.24 2:32.25 2:32.26 2:32.27 2:32.28 2:32.29 | and the multitude of propositions, which are preventing causes in every grasp of wisdom, | and-the-multitude-of the-propositions which-are causes preventing in-every grasp-of wisdom
- 2:32.30 2:32.31 2:32.32 2:32.33 2:32.34 | and especially in the wisdom of Kabbalah, in several respects, | and-especially in-the-wisdom-of Kabbalah in-several respects
- 2:32.35 2:32.36 2:32.37 2:32.38 2:32.39 2:32.40 2:32.41 2:32.42 2:32.43 2:32.44 | and over all of them is the sin, which obstructs and separates between the intellect and the intelligible, etc., | and-over all-of-them the-sin which-is obstructing and-separating between the-intellect and-the-intelligible etc.
- 2:32.45 2:32.46 2:32.47 2:32.48 2:32.49 2:32.50 2:32.51 2:32.52 2:32.53 2:32.54 2:32.55 2:32.56 2:32.57 2:32.58 2:32.59 2:32.60 2:32.61 2:32.62 2:32.63 2:32.64 2:32.65 2:32.66 2:32.67 | and indeed, even though the author of Pardes Rimonim, because of the greatness of his holiness and his piety, is very careful in many places, and warns greatly against a strange thought that is of division, and limit, and dominion of similitudes, etc., | and-indeed even though the-author-of Pardes Rimonim because-of-the-greatness-of his-holiness and-his-piety is-careful very in-many places and-warns greatly against a-thought strange that-is-of-division and-limit and-dominion of-similitudes etc.
- 2:32.68 2:32.69 2:32.70 2:32.71 2:32.72 2:32.73 2:32.74 2:32.75 2:32.76 | we too shall say: would that a spirit from on high be poured out upon us, | and-even we shall-say who will-give and-be-poured-out upon-us a-spirit from-on-high
- 2:32.77 2:32.78 2:32.79 2:32.80 2:32.81 2:32.82 2:32.83 2:32.84 2:32.85 2:32.86 | that we might grasp and understand the matters without inclining from the truth and from the roots of the religion, etc. | that-we-might-grasp and-understand the-matters without inclining from the-truth and-from-the-roots-of the-religion etc.
- 2:32.87 2:32.88 2:32.89 2:32.90 2:32.91 2:32.92 2:32.93 | And by these, take this rule in your hand: all things | and-by-these take this rule in-your-hand all things
- 2:32.94 2:32.95 2:32.96 2:32.97 2:32.98 2:32.99 2:32.100 2:32.101 2:32.102 | that are forbidden in the books of the Kabbalists, which they mention without dispute on the matter and without inquiry, | that-are-forbidden in-the-books-of the-Kabbalists which-they-mention without dispute on-the-matter and-without inquiry
- 2:32.103 2:32.104 2:32.105 2:32.106 2:32.107 2:32.108 | and their plain sense will be understood, and the joining of the things: | will-be-understood their-plain-sense of the-things and-the-joining-of the-things
- 2:32.109 2:32.110 2:32.111 2:32.112 2:32.113 2:32.114 2:32.115 2:32.116 | bind them on a white tablet, and let them be frontlets between your eyes. | bind-them on a-tablet white and-let-them-be frontlets between your-eyes
- 2:32.117 2:32.118 2:32.119 2:32.120 2:32.121 2:32.122 2:32.123 2:32.124 2:32.125 2:32.126 2:32.127 2:32.128 2:32.129 2:32.130 2:32.131 2:32.132 2:32.133 2:32.134 | And even when things are read before you that appear astonishing in themselves, or from the side that it seems, God forbid, that there is division, number, corporeality and change, | and-even when are-read before-you things that-appear astonishing from-the-side-of themselves or from-the-side-of that-it-seems God-forbid that-there-is division and-number and-corporeality and-change
- 2:32.135 2:32.136 2:32.137 2:32.138 2:32.139 | you shall accept them in the faith of an artisan, and say in your heart | you-shall-accept-them in-the-faith-of an-artisan and-say in-your-heart
- 2:32.140 2:32.141 2:32.142 2:32.143 2:32.144 2:32.145 2:32.146 2:32.147 | that your mind is short, and your knowledge in the depth of understanding, without your doubting, | that is-short your-mind and-your-knowledge in-the-depth-of understanding without your-doubting
- 2:32.148 2:32.149 2:32.150 2:32.151 2:32.152 2:32.153 2:32.154 2:32.155 | God forbid, one of the roots of the religion that are stated in Yigdal and in several places, | God-forbid one-of the-roots-of the-religion that-are-stated in-Yigdal and-in-several places
- 2:32.156 2:32.157 2:32.158 2:32.159 2:32.160 2:32.161 2:32.162 2:32.163 2:32.164 2:32.165 2:32.166 2:32.167 | for even those fundamental ones, it is not possible to descend to their roots by human reason, etc. | for even those the-fundamental-ones it-is-not possible to-descend to-their-roots from-the-side-of reason human etc.
- 2:32.168 2:32.169 2:32.170 2:32.171 2:32.172 2:32.173 2:32.174 | And now go and learn that it is forbidden to innovate and to understand a matter | and-now go and-learn that-it-is-forbidden to-innovate and-to-understand a-matter
- 2:32.175 2:32.176 2:32.177 2:32.178 2:32.179 2:32.180 2:32.181 2:32.182 | from what is written in the book, lest he not aim at the truth, | from-within what-is-written in-the-book lest not aim at the-truth
- 2:32.183 2:32.184 2:32.185 2:32.186 2:32.187 2:32.188 2:32.189 2:32.190 | however, to explain verses according to what the Kabbalists said, | however to-explain verses by the-mouth-of what they-said the-Kabbalists
- 2:32.191 2:32.192 2:32.193 2:32.194 2:32.195 2:32.196 2:32.197 | that the premise is agreed; only that on this he explains the verse, | that-the-premise is-agreed only that-on this he-explains the-verse
- 2:32.198 2:32.199 2:32.200 2:32.201 2:32.202 2:32.203 2:32.204 2:32.205 2:32.206 2:32.207 2:32.208 | or says that there is an allusion in it, by means of gematria, or the initial and final letters of words. | or says that-there-is-an-allusion in-it by means-of gematria or the-initials-of and-the-finals-of words
- 2:32.209 2:32.210 2:32.211 2:32.212 2:32.213 2:32.214 2:32.215 2:32.216 2:32.217 2:32.218 2:32.219 | It is good and fine that he only say it as a possibility, without deciding the matter, etc., | it-is-good and-fine it-is that-only he-say-it as-a-way a-possibility without deciding the-matter etc.
- 2:32.220 2:32.221 2:32.222 2:32.223 2:32.224 | only let him not waste his time on them, | only not let-him-waste his-time on-them
- 2:32.225 2:32.226 2:32.227 2:32.228 2:32.229 2:32.230 2:32.231 2:32.232 2:32.233 | and whatever the LORD has put in his heart, and it is correctly tasted without strain, | and-whatever has-put the-LORD in-his-heart and-it is-correctly-tasted rightly without strain
- 2:32.234 2:32.235 2:32.236 2:32.237 | he should not cast it behind his back, | he-should-not cast-it behind his-back
- 2:32.238 2:32.239 2:32.240 2:32.241 2:32.242 2:32.243 2:32.244 2:32.245 | but rather let him write it down for memory, and let it be in his eyes as one who finds great spoil, | but-rather let-him-write-it for-memory and-let-it-be in-his-eyes as-one-who-finds spoil great
- 2:32.246 2:32.247 2:32.248 2:32.249 2:32.250 2:32.251 2:32.252 2:32.253 2:32.254 2:32.255 2:32.256 2:32.257 2:32.258 | for it is a sign for his soul, a portion of God from on high, which has nourishment and access in a high place, | for it is-a-sign for-his-soul a-portion of-God from-on-high which-has to-it nourishment and-access in-a-place high
- 2:32.259 2:32.260 2:32.261 2:32.262 2:32.263 2:32.264 2:32.265 | and a holy one will say to him: and now, behold, I have revealed to you | and-a-holy-one will-say to-him and-now behold I-have-revealed to-you
- 2:32.266 2:32.267 | the secrets of my heart. | the-secrets-of my-heart
- 2:32.268 2:32.269 2:32.270 2:32.271 2:32.272 2:32.273 2:32.274 2:32.275 2:32.276 2:32.277 | And you, the king, in your reflection and your knowledge, holy and broad, tell me, please, your ways, | and-you the-king in-your-reflection and-your-knowledge holy and-broad tell-me please (object-marker) your-ways
- 2:32.278 2:32.279 2:32.280 2:32.281 | and I will set aside my opinion before your opinion. | and-I-will-set-aside my-opinion before your-opinion
- 2:32.282 2:32.283 2:32.284 | Yair Chayyim Bakhrakh. | Yair Chayyim Bakhrakh

Grammar notes:
- One long period runs from the opening to 'etc.' (2:32.44). Its subject is 'the going far toward the place of the Name in the grasp of the secret'; its predicate is 'is the shortness of the one who grasps and the depth of what is grasped'. The list that follows (the propositions, the sin) hangs on the same clause.
- 'Even when' (2:32.117-118, ואף כי) and 'even though' (2:32.46-47, והנה אף כי) open concessive clauses. The main clause of the first is at 2:32.135-139, and of the second at 2:32.68-70 'we too shall say'. The English keeps the Hebrew's order, and has no full stop where the Hebrew has none (2:32.44, 2:32.67).
- 'Would that' is the optative 'who will give' (2:32.71-72). 'And be poured out' (2:32.73) is the wish, and 'that we might grasp' (2:32.77) is its purpose clause, with the verb in the form שנשכיל.
- Imperatives and jussives: 'take' (2:32.88), 'go and learn' (2:32.169-170), 'let them be' (2:32.113), 'let him write it' (2:32.239). 'He should not cast' (2:32.234-235) is a negative command with אל.
- Passive forms, each rendered passive: 'will be understood' (2:32.103), 'are read' (2:32.119), 'is correctly tasted' (2:32.230), 'what is written' (2:32.176), 'is agreed' (2:32.192).

Translator's note: Single prose paragraph in the writer's own voice, addressed to a correspondent ('my beloved'; 'you, the king'). The English follows the Hebrew's order. Where the Hebrew has no full stop (2:32.44, 2:32.67), the English has none. Three choices bear on the words and are recorded in apparatus or readings: ממשל (2:32.5, 'ruler' or the parable), גיוו (2:32.237, 'back' or 'inner part'), and the signature surname (2:32.284), which is unvocalized in the text.

Commentary rulings:
- apparatus: Base text: Wikisource. The 1931 scan is not yet collated, and Wikisource additions are translated as base text and flagged here (Q2).
- apparatus: 'ויערה עלינו רוח ממרום' (2:32.73-76): allusion to Isa 32:15, 'until a spirit be poured out upon us from on high'.
- apparatus: 'לטטפות בין עיניך' (2:32.114-116): allusion to Deut 6:8, 'and they shall be frontlets between your eyes'.
- apparatus: 'קשרם על לוח לבן' (2:32.109-112): allusion to Prov 3:3, which reads 'the tablet of your heart' (לוח לבך). The author writes לבן, 'white', and the English follows his wording.
- apparatus: 'כמוצא שלל רב' (2:32.243-245): allusion to Ps 119:162, 'as one who finds great spoil'.
- apparatus: 'חלק אלוה ממעל' (2:32.250-252): allusion to Job 31:2, 'a portion of God from on high'.
- apparatus: 'ממשל' (2:32.5): read here as 'ruler'. The alternative, 'mashal' (parable), gives 'the parable of the land of Israel, the referent'; it is in readings and is not adopted.
- apparatus: 'גיוו' (2:32.237): 'back' or 'inner part'. The English takes the idiom 'behind his back', to ignore; the literal sense is in the token note.
- apparatus: 'הפ'' (2:32.197): read as 'the verse' by the rule for ס' before a verse. The author does not name the verse.
- apparatus: Signature (2:32.282-284): the surname בכרך is unvocalized. 'Bakhrakh' is a guess at the vowels, not a fixed form. The first two names are as written.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

