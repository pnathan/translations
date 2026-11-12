# Translation pack: milhamot-hashem perek-3-10 (Milchamot ha-Shem 3:16–3:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-10`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
## Terms

- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- אל **אחר** → "another god".
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
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- ואין **צריך** לומר → "needless to say".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
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


## Milchamot ha-Shem 3:16

Hebrew (Wikisource): טז) ואשר טענת עלי ואמרת מי נתן לך רשות לדרוש ולחקור אחר דבריהם, תמה אני מאד על דבריך אלה? וכי נתנה תורה לקרות בה כעוף המצפצף או כבהמה גועה ואינו יודע מה מדבר? והלא העושה כן נקרא רשע ערום כדגרסינן בסוטה (דף כב.) עולא אמר זה שקרא ושנה ולא שמש ת"ח, ר' אליעזר אמר ה"ז עם הארץ, ר' שמואל בר נחמני אמר הרי זה בור, ר' ינאי אמר ה"ז כותי, ר' אחא בר יעקב אמר ה"ז אמגוש, אמר רב נחמן בר יצחק מסתברא כרב אחא בר יעקב דאמרי אינשי רטין מגושא ולא ידע מה רטין, תני תנא ולא ידע מה תני ע"כ. ולפי דבריך עסקי הגוף ומכתבי מסחר שהם מעסקי חיי הגופות עדיפי מעסק התורה שהם חיי הנפש, אשר בהם נזכה לחיי העולם הבא ונקביל בהם פני אלהינו מלכנו? שהרי אם יביאו לך שני מכתבים מעסקי מסחר שהם סותרים זה את זה, אתה נותן לב בהם להבין כדי שלא תפסיד ואינך מקבלם באמונת אומן! ובתורה שהיא חיי הנפש תקבל את הסותר אמונתך שאמונת אומן אם טוב ואם רע בלי הבנה? הלא עיקר עסק התורה לדעת ולהבין מצותיה וחקותיה! ודוד המלך ע"ה אמר מתחנן לפני הקב"ה הבינני ואלמדה מצותיך. ואמר עוד טוב טעם ודעת למדני כי במצותיך האמנתי וקל וחומר הוא, ומה בשאר מצות של תורה צריך לב להבין שלא יטעה להתיר את האסור, או אפילו לאסור את המותר וכמאמר רז"ל הוי זהיר בתלמוד ששגגת תלמוד עולה זדון? במצות ידיעת השי"ת ויחודו על פי האמת והצדק שצונו השי"ת בתורה כמו שכתוב וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד. ודוד המלך ע"ה צוה את בנו לאמר דע את אלהי אביך ועבדהו. והנביא ירמיהו מכריז ואמור אל יתהלל חכם בחכמתו, ואל יתהלל הגבור בגבורתו, אל יתהלל עשיר בעשרו, כי אם בזאת יתהלל המתהלל השכל וידוע אותי כי אני ה' עושה חסדו וכו'. על אחת כמה וכמה שאנחנו צריכים לידע את מי נעבוד.

| id | word | punctuation after |
|---|---|---|
| 3:16.1 | טז | ) |
| 3:16.2 | ואשר |  |
| 3:16.3 | טענת |  |
| 3:16.4 | עלי |  |
| 3:16.5 | ואמרת |  |
| 3:16.6 | מי |  |
| 3:16.7 | נתן |  |
| 3:16.8 | לך |  |
| 3:16.9 | רשות |  |
| 3:16.10 | לדרוש |  |
| 3:16.11 | ולחקור |  |
| 3:16.12 | אחר |  |
| 3:16.13 | דבריהם | , |
| 3:16.14 | תמה |  |
| 3:16.15 | אני |  |
| 3:16.16 | מאד |  |
| 3:16.17 | על |  |
| 3:16.18 | דבריך |  |
| 3:16.19 | אלה | ? |
| 3:16.20 | וכי |  |
| 3:16.21 | נתנה |  |
| 3:16.22 | תורה |  |
| 3:16.23 | לקרות |  |
| 3:16.24 | בה |  |
| 3:16.25 | כעוף |  |
| 3:16.26 | המצפצף |  |
| 3:16.27 | או |  |
| 3:16.28 | כבהמה |  |
| 3:16.29 | גועה |  |
| 3:16.30 | ואינו |  |
| 3:16.31 | יודע |  |
| 3:16.32 | מה |  |
| 3:16.33 | מדבר | ? |
| 3:16.34 | והלא |  |
| 3:16.35 | העושה |  |
| 3:16.36 | כן |  |
| 3:16.37 | נקרא |  |
| 3:16.38 | רשע |  |
| 3:16.39 | ערום |  |
| 3:16.40 | כדגרסינן |  |
| 3:16.41 | בסוטה |  |
| 3:16.42 | דף |  |
| 3:16.43 | כב | .) |
| 3:16.44 | עולא |  |
| 3:16.45 | אמר |  |
| 3:16.46 | זה |  |
| 3:16.47 | שקרא |  |
| 3:16.48 | ושנה |  |
| 3:16.49 | ולא |  |
| 3:16.50 | שמש |  |
| 3:16.51 | ת"ח | , |
| 3:16.52 | ר | ' |
| 3:16.53 | אליעזר |  |
| 3:16.54 | אמר |  |
| 3:16.55 | ה"ז |  |
| 3:16.56 | עם |  |
| 3:16.57 | הארץ | , |
| 3:16.58 | ר | ' |
| 3:16.59 | שמואל |  |
| 3:16.60 | בר |  |
| 3:16.61 | נחמני |  |
| 3:16.62 | אמר |  |
| 3:16.63 | הרי |  |
| 3:16.64 | זה |  |
| 3:16.65 | בור | , |
| 3:16.66 | ר | ' |
| 3:16.67 | ינאי |  |
| 3:16.68 | אמר |  |
| 3:16.69 | ה"ז |  |
| 3:16.70 | כותי | , |
| 3:16.71 | ר | ' |
| 3:16.72 | אחא |  |
| 3:16.73 | בר |  |
| 3:16.74 | יעקב |  |
| 3:16.75 | אמר |  |
| 3:16.76 | ה"ז |  |
| 3:16.77 | אמגוש | , |
| 3:16.78 | אמר |  |
| 3:16.79 | רב |  |
| 3:16.80 | נחמן |  |
| 3:16.81 | בר |  |
| 3:16.82 | יצחק |  |
| 3:16.83 | מסתברא |  |
| 3:16.84 | כרב |  |
| 3:16.85 | אחא |  |
| 3:16.86 | בר |  |
| 3:16.87 | יעקב |  |
| 3:16.88 | דאמרי |  |
| 3:16.89 | אינשי |  |
| 3:16.90 | רטין |  |
| 3:16.91 | מגושא |  |
| 3:16.92 | ולא |  |
| 3:16.93 | ידע |  |
| 3:16.94 | מה |  |
| 3:16.95 | רטין | , |
| 3:16.96 | תני |  |
| 3:16.97 | תנא |  |
| 3:16.98 | ולא |  |
| 3:16.99 | ידע |  |
| 3:16.100 | מה |  |
| 3:16.101 | תני |  |
| 3:16.102 | ע"כ | . |
| 3:16.103 | ולפי |  |
| 3:16.104 | דבריך |  |
| 3:16.105 | עסקי |  |
| 3:16.106 | הגוף |  |
| 3:16.107 | ומכתבי |  |
| 3:16.108 | מסחר |  |
| 3:16.109 | שהם |  |
| 3:16.110 | מעסקי |  |
| 3:16.111 | חיי |  |
| 3:16.112 | הגופות |  |
| 3:16.113 | עדיפי |  |
| 3:16.114 | מעסק |  |
| 3:16.115 | התורה |  |
| 3:16.116 | שהם |  |
| 3:16.117 | חיי |  |
| 3:16.118 | הנפש | , |
| 3:16.119 | אשר |  |
| 3:16.120 | בהם |  |
| 3:16.121 | נזכה |  |
| 3:16.122 | לחיי |  |
| 3:16.123 | העולם |  |
| 3:16.124 | הבא |  |
| 3:16.125 | ונקביל |  |
| 3:16.126 | בהם |  |
| 3:16.127 | פני |  |
| 3:16.128 | אלהינו |  |
| 3:16.129 | מלכנו | ? |
| 3:16.130 | שהרי |  |
| 3:16.131 | אם |  |
| 3:16.132 | יביאו |  |
| 3:16.133 | לך |  |
| 3:16.134 | שני |  |
| 3:16.135 | מכתבים |  |
| 3:16.136 | מעסקי |  |
| 3:16.137 | מסחר |  |
| 3:16.138 | שהם |  |
| 3:16.139 | סותרים |  |
| 3:16.140 | זה |  |
| 3:16.141 | את |  |
| 3:16.142 | זה | , |
| 3:16.143 | אתה |  |
| 3:16.144 | נותן |  |
| 3:16.145 | לב |  |
| 3:16.146 | בהם |  |
| 3:16.147 | להבין |  |
| 3:16.148 | כדי |  |
| 3:16.149 | שלא |  |
| 3:16.150 | תפסיד |  |
| 3:16.151 | ואינך |  |
| 3:16.152 | מקבלם |  |
| 3:16.153 | באמונת |  |
| 3:16.154 | אומן | ! |
| 3:16.155 | ובתורה |  |
| 3:16.156 | שהיא |  |
| 3:16.157 | חיי |  |
| 3:16.158 | הנפש |  |
| 3:16.159 | תקבל |  |
| 3:16.160 | את |  |
| 3:16.161 | הסותר |  |
| 3:16.162 | אמונתך |  |
| 3:16.163 | שאמונת |  |
| 3:16.164 | אומן |  |
| 3:16.165 | אם |  |
| 3:16.166 | טוב |  |
| 3:16.167 | ואם |  |
| 3:16.168 | רע |  |
| 3:16.169 | בלי |  |
| 3:16.170 | הבנה | ? |
| 3:16.171 | הלא |  |
| 3:16.172 | עיקר |  |
| 3:16.173 | עסק |  |
| 3:16.174 | התורה |  |
| 3:16.175 | לדעת |  |
| 3:16.176 | ולהבין |  |
| 3:16.177 | מצותיה |  |
| 3:16.178 | וחקותיה | ! |
| 3:16.179 | ודוד |  |
| 3:16.180 | המלך |  |
| 3:16.181 | ע"ה |  |
| 3:16.182 | אמר |  |
| 3:16.183 | מתחנן |  |
| 3:16.184 | לפני |  |
| 3:16.185 | הקב"ה |  |
| 3:16.186 | הבינני |  |
| 3:16.187 | ואלמדה |  |
| 3:16.188 | מצותיך | . |
| 3:16.189 | ואמר |  |
| 3:16.190 | עוד |  |
| 3:16.191 | טוב |  |
| 3:16.192 | טעם |  |
| 3:16.193 | ודעת |  |
| 3:16.194 | למדני |  |
| 3:16.195 | כי |  |
| 3:16.196 | במצותיך |  |
| 3:16.197 | האמנתי |  |
| 3:16.198 | וקל |  |
| 3:16.199 | וחומר |  |
| 3:16.200 | הוא | , |
| 3:16.201 | ומה |  |
| 3:16.202 | בשאר |  |
| 3:16.203 | מצות |  |
| 3:16.204 | של |  |
| 3:16.205 | תורה |  |
| 3:16.206 | צריך |  |
| 3:16.207 | לב |  |
| 3:16.208 | להבין |  |
| 3:16.209 | שלא |  |
| 3:16.210 | יטעה |  |
| 3:16.211 | להתיר |  |
| 3:16.212 | את |  |
| 3:16.213 | האסור | , |
| 3:16.214 | או |  |
| 3:16.215 | אפילו |  |
| 3:16.216 | לאסור |  |
| 3:16.217 | את |  |
| 3:16.218 | המותר |  |
| 3:16.219 | וכמאמר |  |
| 3:16.220 | רז"ל |  |
| 3:16.221 | הוי |  |
| 3:16.222 | זהיר |  |
| 3:16.223 | בתלמוד |  |
| 3:16.224 | ששגגת |  |
| 3:16.225 | תלמוד |  |
| 3:16.226 | עולה |  |
| 3:16.227 | זדון | ? |
| 3:16.228 | במצות |  |
| 3:16.229 | ידיעת |  |
| 3:16.230 | השי"ת |  |
| 3:16.231 | ויחודו |  |
| 3:16.232 | על |  |
| 3:16.233 | פי |  |
| 3:16.234 | האמת |  |
| 3:16.235 | והצדק |  |
| 3:16.236 | שצונו |  |
| 3:16.237 | השי"ת |  |
| 3:16.238 | בתורה |  |
| 3:16.239 | כמו |  |
| 3:16.240 | שכתוב |  |
| 3:16.241 | וידעת |  |
| 3:16.242 | היום |  |
| 3:16.243 | והשבות |  |
| 3:16.244 | אל |  |
| 3:16.245 | לבבך |  |
| 3:16.246 | כי |  |
| 3:16.247 | ה | ' |
| 3:16.248 | הוא |  |
| 3:16.249 | האלהים |  |
| 3:16.250 | בשמים |  |
| 3:16.251 | ממעל |  |
| 3:16.252 | ועל |  |
| 3:16.253 | הארץ |  |
| 3:16.254 | מתחת |  |
| 3:16.255 | אין |  |
| 3:16.256 | עוד | . |
| 3:16.257 | ודוד |  |
| 3:16.258 | המלך |  |
| 3:16.259 | ע"ה |  |
| 3:16.260 | צוה |  |
| 3:16.261 | את |  |
| 3:16.262 | בנו |  |
| 3:16.263 | לאמר |  |
| 3:16.264 | דע |  |
| 3:16.265 | את |  |
| 3:16.266 | אלהי |  |
| 3:16.267 | אביך |  |
| 3:16.268 | ועבדהו | . |
| 3:16.269 | והנביא |  |
| 3:16.270 | ירמיהו |  |
| 3:16.271 | מכריז |  |
| 3:16.272 | ואמור |  |
| 3:16.273 | אל |  |
| 3:16.274 | יתהלל |  |
| 3:16.275 | חכם |  |
| 3:16.276 | בחכמתו | , |
| 3:16.277 | ואל |  |
| 3:16.278 | יתהלל |  |
| 3:16.279 | הגבור |  |
| 3:16.280 | בגבורתו | , |
| 3:16.281 | אל |  |
| 3:16.282 | יתהלל |  |
| 3:16.283 | עשיר |  |
| 3:16.284 | בעשרו | , |
| 3:16.285 | כי |  |
| 3:16.286 | אם |  |
| 3:16.287 | בזאת |  |
| 3:16.288 | יתהלל |  |
| 3:16.289 | המתהלל |  |
| 3:16.290 | השכל |  |
| 3:16.291 | וידוע |  |
| 3:16.292 | אותי |  |
| 3:16.293 | כי |  |
| 3:16.294 | אני |  |
| 3:16.295 | ה | ' |
| 3:16.296 | עושה |  |
| 3:16.297 | חסדו |  |
| 3:16.298 | וכו | '. |
| 3:16.299 | על |  |
| 3:16.300 | אחת |  |
| 3:16.301 | כמה |  |
| 3:16.302 | וכמה |  |
| 3:16.303 | שאנחנו |  |
| 3:16.304 | צריכים |  |
| 3:16.305 | לידע |  |
| 3:16.306 | את |  |
| 3:16.307 | מי |  |
| 3:16.308 | נעבוד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

