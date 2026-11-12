# Translation pack: milhamot-hashem perek-1-03 (Milchamot ha-Shem 1:4–1:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-03`.

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
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
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
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **רש"י** → "Rashi" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **רבי'** before the lawgiver's name → "Moses our teacher".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הכוזרי / כוזרי** → "the Kuzari" (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **ממלכי** → "of the kings of".
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **בא"י** → "in the Land of Israel".

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
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- עלת **העלות** → "the Cause of causes".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
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
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- **זווג** → "coupling".
- **מין / מינות** → "heretic / heresy".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **הינ"ץ** → "H.N.Ts." (note: perhaps ה + ינ"ץ, "may his Rock protect him", of
  the Imam); no person named.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **גיזומיו** → "exaggerations" of claims, else "trimmings"; the other in a note.
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **סכלות** והוללות → "folly and madness".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
- בית **הספר** (non-technical 1931 Hebrew) → "school".
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
- בחלק **שפתיו** → "with the smoothness of his lips".
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **רחום / וחנון** (Exod 34:6) → "the LORD, the LORD, God merciful and
  gracious, slow to anger".
- **כדנה / אלהיא / יאבדו** (Jer 10:11) → "Thus shall you say to them: the gods
  that did not make the heavens and the earth shall perish from the earth and
  from under these heavens."
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **העתקתו / תרגום**: his own Hebrew translations (10:7–10:9; "[תרגום הרב
  המחבר: …]") are separate segments, each from its own Hebrew; a difference
  in sense in `apparatus`.
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
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


## Milchamot ha-Shem 1:4

Hebrew (Wikisource): ולתקופת השנה השלישית פרצה המלחמה נגד הלמודים הטובים והסדרים הנעימים כמצווה עלינו מרז״ל. ויוציאו המתנגדים דבה רעה לאמר כי על ידי זה יצאו הילדים ח״ו לתרבות רעה, והרבה מהטפשים האמינו לשקריהם והוציאו את ילדיהם מבית הספר. וגם הבנים הסכלים אשר יצאו מבית הספר עוינים את מוריהם כי לעניות דעתם הם מינים ויוצאים מן הכלל אך המשכילים שבהם יודעים ומכירים כי שקר בימין מוציאי הדבה הרעה הלזו, ודאגו מאד על שהוציאום אבותם מבית הספר בסכלותם! וכאשר התעו רבים חזקה יד החשכים ובמלשינות שקר הלשינו על המלמדים אצל הממשלה שהם חפצים להמריד את יושבי הארץ. ושהם כותבי פלסתר על הממשלה אל ממשלת אנגליא לחסות בצל ממשלתה וגם אל הינ״ץ ועבור זה המלשינות הובלו המלמדים אל בית הכלא וגם גורשו מבית הספר. והביאו תחתיהם מלמד אחר אשר לא ידע לא עברית על בריה ולא ערבית כתקנה. ואין עליו ללמדם רק פרשת השבוע עם תרגום אונקלוס ותרגום רס״ג בערבי. והפטרה עם תרגום יונתן ב״ע באין הבין כלום כשאר המלמדים! ולימים מספר בטל בית הספר לגמרי! כל זאת באתנו על אשר לא מחינו בכבוד קונינו בעת שהיינו שומעים מהלומדים בזהר ושאר מקובלים דברים שאסור לשמען ומחרישים להם כמו שאמרו רז״ל דר״א נתפש למינות וכו' שמא דבר מינות וכו' ונתקיים בנו מאמר רז״ל דמודי להון נפיל בידיהון דמתרחץ עליהון דיליה דילהון ע״כ! לכן עזבנו דרכינו בלימוד השפות העברית והערבית בדקדוקן על כורחנו מפני חמת המציק על ידי מלשינות שקר ונפן אל דרך הקדש כבוד אבינו שבשמים בבטול דעות הפילסוף נביא השקר מחבר הזהר אשר אסף בחפניו דעות הפלסופיא הקדומה עובדי האלילים תמונות וכוחות נעלמים אשר בדו מלבם. ומחבר הזהר תפס דעותיהם וחשב להכניסם באמונת תורתינו הקדושה, וקרא את שמו בשם רשב״י התנא בשקר ואמר שנגלה לו הקב״ה הנקרא אצלו עתיקא קדישא וגילה לו רזין טמירין. ותלה נבואתו ורעיונותיו ברשב״י התנא ע״ה. ועשה את חבורו הזהר ואמר שהעליונים הסכימו לחבורו לפי שראה הפילסוף הזה כי המשנה ושני התלמודים ומדרשי רז״ל האמתיים עומדים נגדו באמונת אחדותו ית' לכן חשב לפסלן ובזה ולעג להם ולכל העוסקים בהם, וכינה אותם בשמות של זלזול ופחיתות, כאשר יראה המעיין בפנים! והבטיח שלא יזכו ישראל לגאולה עד אשר ישובו מדרכם מלעסוק במשנה ובתלמוד. רק יעסקו בתורתו החדשה המלאה אלילי תמונות ופרצופים בעלי מפלצות וזווגים שונים. ובתחבולה ומרמה קברו בארון בקרקע עולם להתיישן ויתחיל להרקיב, אחר ימים פרסמו והוציאו ע״י מלך ממלכי העכו״מ (כמ״ש הר' חיד״א בספרו שם הגדולים במערכת ספרים) כאלו הוא חבור קדום מרשב״י התנא. והרבה מרבותינו האחרונים נמשכו אחר הפילסוף הנז' בעצימות עין כי לקח אותם בחלק שפתיו וגיזומיו, ולא שתו לב לכזביו! חלילה לשום בר ישראל להאמין כי רשב״י התנא או זולתו מרז״ל האמינו באמונות כאלה להחליף ולהמיר את ה' אלוהינו אשר הודיע דרכיו למשה לבני ישראל עלילותיו, כי חנון ורחום הוא ארך אפים ורב חסד וכו' באל זר קצר אפים (ז״א) ולצרף ולשתף עמו חמשה פרצופים אשר לא נתברר אפי' מציאותם במופת (ר' עובדיה ספורנו דברים י״ג ג'. ושם כ״ט כ״ה וע״ש כ״ח י״ד) ולקרותם בשם ה' אלהינו, ואת ה' אלהים אמת שנתבררה מציאותו בכמה מופתים כראי מוצק חזקים כמו שכ' הרס״ג בס' האמונות והדעות והר' בעל חובת הלבבות והרמב״ם במורה נבוכים וביד, נטוש ונעזוב ונאמר שאין לו שם ולא כנוי ונעבוד את הפרצופים והצורות אשר לפי דבריו נבראו ונשתלשלו ממנו! וכבר הארכנו בפנים בראיות נכוחות מרז״ל בביטול הדעות הזרות האלו של הפילסוף נביא השקר הנז'! כי הוא מטרת תורתינו הקדושה להרחיקנו מאמונת האלילים בין גשמיים בין רוחניים. ולדעת כי ה' הוא האלהים אין עוד מלבדו. ואין לו לא אב ולא אח ולא בן (רבה פרשת יתרו) ולא קיבל מלכותו מאחר ואינו עתיד למסרה לאחר כמו שביארו רבותינו ז״ל (רבה במדבר שיר השירים פרשה א' ירושלמי סנהדרין פרק דיני ממונות) על פסוק אני ראשון ואני אחרון ומבלעדי אין אלהים ועל מה שאמר נבוכדנצר וריויה די רביעאה דמי לבר אלהין! וראיתי להר' מנחם די לונזבנו בס' דרך החיים שהרחיב פיו לדבר תועה על החסיד הקדוש בעל חובת הלבבות על אמרו בשם הפילסוף שאמר שלא יוכל לעבוד עלת העלות רק נביא הדור בטבעו, או הפילסוף המובהק במה שקנהו מן החכמה אבל זולתם עובדים זולתו מפני שאינם מבינים נמצא אלא מורכב! ועל דברי החסיד האלו התקצף וחשב כי כאן מקום תלונה עליו ותואנה מצא להתעות השומעים ולהבהילם באמרו וכי פלוני ופלוני מחכמינו האחרונים המפורסמים עובדים זולתו? והוא ספרם בלשונו אחד אחד בשמותם. ולא שם אל לבו כי הרבה מהחכמים שַמָנָה הלכו בעצימות עינים והאמינו לדברי מחבר הזהר אשר כינה את עצמו בשם רשב״י התנא כדי לתלות נבואותיו השקריות בשמו של התנא, ודבר סרה על ה' אלהינו שנגלה לו בבית מדרשו עתיקא קדישא וגילה לו סודות נעלמים וכוחות נעלמות (זהר בראשית כ״ב) כי הבורא הוא האליל הנקרא אבא שהיה מצוה לזוגתו האלילה אימא והיא עושה כל אשר יאמר אליה חוץ מבריאת אדם הראשון שהיא אימא שאמרה ועשתה את אדם הראשון שלא ברצון אבא! ובפרשת וירא דף קי״ב וכן בפרשת בלק דף קצ״א מיחס עבודתינו ותפילותינו לאל זר קצר אפים (זעיר אנפין) אשר לא עשה את השמים ואת הארץ. וקורא הוא לעצבים הללו בשם ה' אלהינו. ועל זעיר אנפין קרא ירמיה הנביא ע״ה כדנה תימרון להום אלהיא די שמיא וארקא לא עבדו ייבדו מן ארעא ומן תחות שמיא אלה. שהרי מבואר בזהר בכמה מקומות שהוא נברא ולא בורא (עיין זהר בראשית כ״ב וביושר לבב דף ג' ע״ב) ועכ״ז הוא אומר לעבדו ולברכו ולקבלו עליו באלוה ולירא מזוגתו! והרבה מרבותינו האחרונים שהזכיר בספר דרך החיים בתמיהתו על החסיד בעל חובת הלבבות הלכו בעצימת עין ונמשכו אחריו והאמינו באמונתו וקיבלו האליל קצר אפים (ז״א) עליהם באלוה! וכל זה גרם להם על שלא התבוננו בחכמת יוצר הכל שהוא לחקור ולבחון במעשה בראשית המבוארים בדברי הר' מורה נבוכים ובדברי החסיד בעל חובת הלבבות. ורבינו סעדיה גאון ור״י הלוי בכוזרי כמו שאמר הנביא ע״ה ואת פעל ה' לא יביטו. כי חשבו דבריהם לחכמה יונית ורחקו ממנה הרחק כמטחוי קשת, וגם לא שמו אל לבם דברי אבותינו ז״ל האמתיים במדרשים שלא הכיר אברהם אבינו את בוראו אלא בהתבוננו בפעל ה' וחקר ודרש בהם והביא מופתים מהם עד שהשיג האמת ואז נגלה לו הקב״ה ואמר לו אהבת צדק ותשנא רשע וכו' משל לאחד שהיה מהלך בדרך וראה בירה וכו' (מדרש הגדול ריש פרשת לך לך) ועל כרחנו נודה על האמת שדברי החסיד בעל חובת הלבבות ברור מללו שלא יוכל לעבוד את אלהים אמת אלא וכו' אבל זולתם עובדים זולתו כמאמר הנביא וימים רבים לישראל ללא אלהי אמת, וכמו שעינינו רואות ואזנינו שומעות עד שאפשר שיאמר אדם לחבירו! עַמְּךָ עַמִּי וְאלֹהֶיך לֹא אֱלהָי! ואל אלהי אמת נשפוך שיחנו לנחותינו במעגלי יושר ולהצילנו מכל אמונה זרה אכי״ר.

| id | word | punctuation after |
|---|---|---|
| 1:4.1 | ולתקופת |  |
| 1:4.2 | השנה |  |
| 1:4.3 | השלישית |  |
| 1:4.4 | פרצה |  |
| 1:4.5 | המלחמה |  |
| 1:4.6 | נגד |  |
| 1:4.7 | הלמודים |  |
| 1:4.8 | הטובים |  |
| 1:4.9 | והסדרים |  |
| 1:4.10 | הנעימים |  |
| 1:4.11 | כמצווה |  |
| 1:4.12 | עלינו |  |
| 1:4.13 | מרז״ל | . |
| 1:4.14 | ויוציאו |  |
| 1:4.15 | המתנגדים |  |
| 1:4.16 | דבה |  |
| 1:4.17 | רעה |  |
| 1:4.18 | לאמר |  |
| 1:4.19 | כי |  |
| 1:4.20 | על |  |
| 1:4.21 | ידי |  |
| 1:4.22 | זה |  |
| 1:4.23 | יצאו |  |
| 1:4.24 | הילדים |  |
| 1:4.25 | ח״ו |  |
| 1:4.26 | לתרבות |  |
| 1:4.27 | רעה | , |
| 1:4.28 | והרבה |  |
| 1:4.29 | מהטפשים |  |
| 1:4.30 | האמינו |  |
| 1:4.31 | לשקריהם |  |
| 1:4.32 | והוציאו |  |
| 1:4.33 | את |  |
| 1:4.34 | ילדיהם |  |
| 1:4.35 | מבית |  |
| 1:4.36 | הספר | . |
| 1:4.37 | וגם |  |
| 1:4.38 | הבנים |  |
| 1:4.39 | הסכלים |  |
| 1:4.40 | אשר |  |
| 1:4.41 | יצאו |  |
| 1:4.42 | מבית |  |
| 1:4.43 | הספר |  |
| 1:4.44 | עוינים |  |
| 1:4.45 | את |  |
| 1:4.46 | מוריהם |  |
| 1:4.47 | כי |  |
| 1:4.48 | לעניות |  |
| 1:4.49 | דעתם |  |
| 1:4.50 | הם |  |
| 1:4.51 | מינים |  |
| 1:4.52 | ויוצאים |  |
| 1:4.53 | מן |  |
| 1:4.54 | הכלל |  |
| 1:4.55 | אך |  |
| 1:4.56 | המשכילים |  |
| 1:4.57 | שבהם |  |
| 1:4.58 | יודעים |  |
| 1:4.59 | ומכירים |  |
| 1:4.60 | כי |  |
| 1:4.61 | שקר |  |
| 1:4.62 | בימין |  |
| 1:4.63 | מוציאי |  |
| 1:4.64 | הדבה |  |
| 1:4.65 | הרעה |  |
| 1:4.66 | הלזו | , |
| 1:4.67 | ודאגו |  |
| 1:4.68 | מאד |  |
| 1:4.69 | על |  |
| 1:4.70 | שהוציאום |  |
| 1:4.71 | אבותם |  |
| 1:4.72 | מבית |  |
| 1:4.73 | הספר |  |
| 1:4.74 | בסכלותם | ! |
| 1:4.75 | וכאשר |  |
| 1:4.76 | התעו |  |
| 1:4.77 | רבים |  |
| 1:4.78 | חזקה |  |
| 1:4.79 | יד |  |
| 1:4.80 | החשכים |  |
| 1:4.81 | ובמלשינות |  |
| 1:4.82 | שקר |  |
| 1:4.83 | הלשינו |  |
| 1:4.84 | על |  |
| 1:4.85 | המלמדים |  |
| 1:4.86 | אצל |  |
| 1:4.87 | הממשלה |  |
| 1:4.88 | שהם |  |
| 1:4.89 | חפצים |  |
| 1:4.90 | להמריד |  |
| 1:4.91 | את |  |
| 1:4.92 | יושבי |  |
| 1:4.93 | הארץ | . |
| 1:4.94 | ושהם |  |
| 1:4.95 | כותבי |  |
| 1:4.96 | פלסתר |  |
| 1:4.97 | על |  |
| 1:4.98 | הממשלה |  |
| 1:4.99 | אל |  |
| 1:4.100 | ממשלת |  |
| 1:4.101 | אנגליא |  |
| 1:4.102 | לחסות |  |
| 1:4.103 | בצל |  |
| 1:4.104 | ממשלתה |  |
| 1:4.105 | וגם |  |
| 1:4.106 | אל |  |
| 1:4.107 | הינ״ץ |  |
| 1:4.108 | ועבור |  |
| 1:4.109 | זה |  |
| 1:4.110 | המלשינות |  |
| 1:4.111 | הובלו |  |
| 1:4.112 | המלמדים |  |
| 1:4.113 | אל |  |
| 1:4.114 | בית |  |
| 1:4.115 | הכלא |  |
| 1:4.116 | וגם |  |
| 1:4.117 | גורשו |  |
| 1:4.118 | מבית |  |
| 1:4.119 | הספר | . |
| 1:4.120 | והביאו |  |
| 1:4.121 | תחתיהם |  |
| 1:4.122 | מלמד |  |
| 1:4.123 | אחר |  |
| 1:4.124 | אשר |  |
| 1:4.125 | לא |  |
| 1:4.126 | ידע |  |
| 1:4.127 | לא |  |
| 1:4.128 | עברית |  |
| 1:4.129 | על |  |
| 1:4.130 | בריה |  |
| 1:4.131 | ולא |  |
| 1:4.132 | ערבית |  |
| 1:4.133 | כתקנה | . |
| 1:4.134 | ואין |  |
| 1:4.135 | עליו |  |
| 1:4.136 | ללמדם |  |
| 1:4.137 | רק |  |
| 1:4.138 | פרשת |  |
| 1:4.139 | השבוע |  |
| 1:4.140 | עם |  |
| 1:4.141 | תרגום |  |
| 1:4.142 | אונקלוס |  |
| 1:4.143 | ותרגום |  |
| 1:4.144 | רס״ג |  |
| 1:4.145 | בערבי | . |
| 1:4.146 | והפטרה |  |
| 1:4.147 | עם |  |
| 1:4.148 | תרגום |  |
| 1:4.149 | יונתן |  |
| 1:4.150 | ב״ע |  |
| 1:4.151 | באין |  |
| 1:4.152 | הבין |  |
| 1:4.153 | כלום |  |
| 1:4.154 | כשאר |  |
| 1:4.155 | המלמדים | ! |
| 1:4.156 | ולימים |  |
| 1:4.157 | מספר |  |
| 1:4.158 | בטל |  |
| 1:4.159 | בית |  |
| 1:4.160 | הספר |  |
| 1:4.161 | לגמרי | ! |
| 1:4.162 | כל |  |
| 1:4.163 | זאת |  |
| 1:4.164 | באתנו |  |
| 1:4.165 | על |  |
| 1:4.166 | אשר |  |
| 1:4.167 | לא |  |
| 1:4.168 | מחינו |  |
| 1:4.169 | בכבוד |  |
| 1:4.170 | קונינו |  |
| 1:4.171 | בעת |  |
| 1:4.172 | שהיינו |  |
| 1:4.173 | שומעים |  |
| 1:4.174 | מהלומדים |  |
| 1:4.175 | בזהר |  |
| 1:4.176 | ושאר |  |
| 1:4.177 | מקובלים |  |
| 1:4.178 | דברים |  |
| 1:4.179 | שאסור |  |
| 1:4.180 | לשמען |  |
| 1:4.181 | ומחרישים |  |
| 1:4.182 | להם |  |
| 1:4.183 | כמו |  |
| 1:4.184 | שאמרו |  |
| 1:4.185 | רז״ל |  |
| 1:4.186 | דר״א |  |
| 1:4.187 | נתפש |  |
| 1:4.188 | למינות |  |
| 1:4.189 | וכו | ' |
| 1:4.190 | שמא |  |
| 1:4.191 | דבר |  |
| 1:4.192 | מינות |  |
| 1:4.193 | וכו | ' |
| 1:4.194 | ונתקיים |  |
| 1:4.195 | בנו |  |
| 1:4.196 | מאמר |  |
| 1:4.197 | רז״ל |  |
| 1:4.198 | דמודי |  |
| 1:4.199 | להון |  |
| 1:4.200 | נפיל |  |
| 1:4.201 | בידיהון |  |
| 1:4.202 | דמתרחץ |  |
| 1:4.203 | עליהון |  |
| 1:4.204 | דיליה |  |
| 1:4.205 | דילהון |  |
| 1:4.206 | ע״כ | ! |
| 1:4.207 | לכן |  |
| 1:4.208 | עזבנו |  |
| 1:4.209 | דרכינו |  |
| 1:4.210 | בלימוד |  |
| 1:4.211 | השפות |  |
| 1:4.212 | העברית |  |
| 1:4.213 | והערבית |  |
| 1:4.214 | בדקדוקן |  |
| 1:4.215 | על |  |
| 1:4.216 | כורחנו |  |
| 1:4.217 | מפני |  |
| 1:4.218 | חמת |  |
| 1:4.219 | המציק |  |
| 1:4.220 | על |  |
| 1:4.221 | ידי |  |
| 1:4.222 | מלשינות |  |
| 1:4.223 | שקר |  |
| 1:4.224 | ונפן |  |
| 1:4.225 | אל |  |
| 1:4.226 | דרך |  |
| 1:4.227 | הקדש |  |
| 1:4.228 | כבוד |  |
| 1:4.229 | אבינו |  |
| 1:4.230 | שבשמים |  |
| 1:4.231 | בבטול |  |
| 1:4.232 | דעות |  |
| 1:4.233 | הפילסוף |  |
| 1:4.234 | נביא |  |
| 1:4.235 | השקר |  |
| 1:4.236 | מחבר |  |
| 1:4.237 | הזהר |  |
| 1:4.238 | אשר |  |
| 1:4.239 | אסף |  |
| 1:4.240 | בחפניו |  |
| 1:4.241 | דעות |  |
| 1:4.242 | הפלסופיא |  |
| 1:4.243 | הקדומה |  |
| 1:4.244 | עובדי |  |
| 1:4.245 | האלילים |  |
| 1:4.246 | תמונות |  |
| 1:4.247 | וכוחות |  |
| 1:4.248 | נעלמים |  |
| 1:4.249 | אשר |  |
| 1:4.250 | בדו |  |
| 1:4.251 | מלבם | . |
| 1:4.252 | ומחבר |  |
| 1:4.253 | הזהר |  |
| 1:4.254 | תפס |  |
| 1:4.255 | דעותיהם |  |
| 1:4.256 | וחשב |  |
| 1:4.257 | להכניסם |  |
| 1:4.258 | באמונת |  |
| 1:4.259 | תורתינו |  |
| 1:4.260 | הקדושה | , |
| 1:4.261 | וקרא |  |
| 1:4.262 | את |  |
| 1:4.263 | שמו |  |
| 1:4.264 | בשם |  |
| 1:4.265 | רשב״י |  |
| 1:4.266 | התנא |  |
| 1:4.267 | בשקר |  |
| 1:4.268 | ואמר |  |
| 1:4.269 | שנגלה |  |
| 1:4.270 | לו |  |
| 1:4.271 | הקב״ה |  |
| 1:4.272 | הנקרא |  |
| 1:4.273 | אצלו |  |
| 1:4.274 | עתיקא |  |
| 1:4.275 | קדישא |  |
| 1:4.276 | וגילה |  |
| 1:4.277 | לו |  |
| 1:4.278 | רזין |  |
| 1:4.279 | טמירין | . |
| 1:4.280 | ותלה |  |
| 1:4.281 | נבואתו |  |
| 1:4.282 | ורעיונותיו |  |
| 1:4.283 | ברשב״י |  |
| 1:4.284 | התנא |  |
| 1:4.285 | ע״ה | . |
| 1:4.286 | ועשה |  |
| 1:4.287 | את |  |
| 1:4.288 | חבורו |  |
| 1:4.289 | הזהר |  |
| 1:4.290 | ואמר |  |
| 1:4.291 | שהעליונים |  |
| 1:4.292 | הסכימו |  |
| 1:4.293 | לחבורו |  |
| 1:4.294 | לפי |  |
| 1:4.295 | שראה |  |
| 1:4.296 | הפילסוף |  |
| 1:4.297 | הזה |  |
| 1:4.298 | כי |  |
| 1:4.299 | המשנה |  |
| 1:4.300 | ושני |  |
| 1:4.301 | התלמודים |  |
| 1:4.302 | ומדרשי |  |
| 1:4.303 | רז״ל |  |
| 1:4.304 | האמתיים |  |
| 1:4.305 | עומדים |  |
| 1:4.306 | נגדו |  |
| 1:4.307 | באמונת |  |
| 1:4.308 | אחדותו |  |
| 1:4.309 | ית | ' |
| 1:4.310 | לכן |  |
| 1:4.311 | חשב |  |
| 1:4.312 | לפסלן |  |
| 1:4.313 | ובזה |  |
| 1:4.314 | ולעג |  |
| 1:4.315 | להם |  |
| 1:4.316 | ולכל |  |
| 1:4.317 | העוסקים |  |
| 1:4.318 | בהם | , |
| 1:4.319 | וכינה |  |
| 1:4.320 | אותם |  |
| 1:4.321 | בשמות |  |
| 1:4.322 | של |  |
| 1:4.323 | זלזול |  |
| 1:4.324 | ופחיתות | , |
| 1:4.325 | כאשר |  |
| 1:4.326 | יראה |  |
| 1:4.327 | המעיין |  |
| 1:4.328 | בפנים | ! |
| 1:4.329 | והבטיח |  |
| 1:4.330 | שלא |  |
| 1:4.331 | יזכו |  |
| 1:4.332 | ישראל |  |
| 1:4.333 | לגאולה |  |
| 1:4.334 | עד |  |
| 1:4.335 | אשר |  |
| 1:4.336 | ישובו |  |
| 1:4.337 | מדרכם |  |
| 1:4.338 | מלעסוק |  |
| 1:4.339 | במשנה |  |
| 1:4.340 | ובתלמוד | . |
| 1:4.341 | רק |  |
| 1:4.342 | יעסקו |  |
| 1:4.343 | בתורתו |  |
| 1:4.344 | החדשה |  |
| 1:4.345 | המלאה |  |
| 1:4.346 | אלילי |  |
| 1:4.347 | תמונות |  |
| 1:4.348 | ופרצופים |  |
| 1:4.349 | בעלי |  |
| 1:4.350 | מפלצות |  |
| 1:4.351 | וזווגים |  |
| 1:4.352 | שונים | . |
| 1:4.353 | ובתחבולה |  |
| 1:4.354 | ומרמה |  |
| 1:4.355 | קברו |  |
| 1:4.356 | בארון |  |
| 1:4.357 | בקרקע |  |
| 1:4.358 | עולם |  |
| 1:4.359 | להתיישן |  |
| 1:4.360 | ויתחיל |  |
| 1:4.361 | להרקיב | , |
| 1:4.362 | אחר |  |
| 1:4.363 | ימים |  |
| 1:4.364 | פרסמו |  |
| 1:4.365 | והוציאו |  |
| 1:4.366 | ע״י |  |
| 1:4.367 | מלך |  |
| 1:4.368 | ממלכי |  |
| 1:4.369 | העכו״מ |  |
| 1:4.370 | כמ״ש |  |
| 1:4.371 | הר | ' |
| 1:4.372 | חיד״א |  |
| 1:4.373 | בספרו |  |
| 1:4.374 | שם |  |
| 1:4.375 | הגדולים |  |
| 1:4.376 | במערכת |  |
| 1:4.377 | ספרים | ) |
| 1:4.378 | כאלו |  |
| 1:4.379 | הוא |  |
| 1:4.380 | חבור |  |
| 1:4.381 | קדום |  |
| 1:4.382 | מרשב״י |  |
| 1:4.383 | התנא | . |
| 1:4.384 | והרבה |  |
| 1:4.385 | מרבותינו |  |
| 1:4.386 | האחרונים |  |
| 1:4.387 | נמשכו |  |
| 1:4.388 | אחר |  |
| 1:4.389 | הפילסוף |  |
| 1:4.390 | הנז | ' |
| 1:4.391 | בעצימות |  |
| 1:4.392 | עין |  |
| 1:4.393 | כי |  |
| 1:4.394 | לקח |  |
| 1:4.395 | אותם |  |
| 1:4.396 | בחלק |  |
| 1:4.397 | שפתיו |  |
| 1:4.398 | וגיזומיו | , |
| 1:4.399 | ולא |  |
| 1:4.400 | שתו |  |
| 1:4.401 | לב |  |
| 1:4.402 | לכזביו | ! |
| 1:4.403 | חלילה |  |
| 1:4.404 | לשום |  |
| 1:4.405 | בר |  |
| 1:4.406 | ישראל |  |
| 1:4.407 | להאמין |  |
| 1:4.408 | כי |  |
| 1:4.409 | רשב״י |  |
| 1:4.410 | התנא |  |
| 1:4.411 | או |  |
| 1:4.412 | זולתו |  |
| 1:4.413 | מרז״ל |  |
| 1:4.414 | האמינו |  |
| 1:4.415 | באמונות |  |
| 1:4.416 | כאלה |  |
| 1:4.417 | להחליף |  |
| 1:4.418 | ולהמיר |  |
| 1:4.419 | את |  |
| 1:4.420 | ה | ' |
| 1:4.421 | אלוהינו |  |
| 1:4.422 | אשר |  |
| 1:4.423 | הודיע |  |
| 1:4.424 | דרכיו |  |
| 1:4.425 | למשה |  |
| 1:4.426 | לבני |  |
| 1:4.427 | ישראל |  |
| 1:4.428 | עלילותיו | , |
| 1:4.429 | כי |  |
| 1:4.430 | חנון |  |
| 1:4.431 | ורחום |  |
| 1:4.432 | הוא |  |
| 1:4.433 | ארך |  |
| 1:4.434 | אפים |  |
| 1:4.435 | ורב |  |
| 1:4.436 | חסד |  |
| 1:4.437 | וכו | ' |
| 1:4.438 | באל |  |
| 1:4.439 | זר |  |
| 1:4.440 | קצר |  |
| 1:4.441 | אפים |  |
| 1:4.442 | ז״א | ) |
| 1:4.443 | ולצרף |  |
| 1:4.444 | ולשתף |  |
| 1:4.445 | עמו |  |
| 1:4.446 | חמשה |  |
| 1:4.447 | פרצופים |  |
| 1:4.448 | אשר |  |
| 1:4.449 | לא |  |
| 1:4.450 | נתברר |  |
| 1:4.451 | אפי | ' |
| 1:4.452 | מציאותם |  |
| 1:4.453 | במופת |  |
| 1:4.454 | ר | ' |
| 1:4.455 | עובדיה |  |
| 1:4.456 | ספורנו |  |
| 1:4.457 | דברים |  |
| 1:4.458 | י״ג |  |
| 1:4.459 | ג | '. |
| 1:4.460 | ושם |  |
| 1:4.461 | כ״ט |  |
| 1:4.462 | כ״ה |  |
| 1:4.463 | וע״ש |  |
| 1:4.464 | כ״ח |  |
| 1:4.465 | י״ד | ) |
| 1:4.466 | ולקרותם |  |
| 1:4.467 | בשם |  |
| 1:4.468 | ה | ' |
| 1:4.469 | אלהינו | , |
| 1:4.470 | ואת |  |
| 1:4.471 | ה | ' |
| 1:4.472 | אלהים |  |
| 1:4.473 | אמת |  |
| 1:4.474 | שנתבררה |  |
| 1:4.475 | מציאותו |  |
| 1:4.476 | בכמה |  |
| 1:4.477 | מופתים |  |
| 1:4.478 | כראי |  |
| 1:4.479 | מוצק |  |
| 1:4.480 | חזקים |  |
| 1:4.481 | כמו |  |
| 1:4.482 | שכ | ' |
| 1:4.483 | הרס״ג |  |
| 1:4.484 | בס | ' |
| 1:4.485 | האמונות |  |
| 1:4.486 | והדעות |  |
| 1:4.487 | והר | ' |
| 1:4.488 | בעל |  |
| 1:4.489 | חובת |  |
| 1:4.490 | הלבבות |  |
| 1:4.491 | והרמב״ם |  |
| 1:4.492 | במורה |  |
| 1:4.493 | נבוכים |  |
| 1:4.494 | וביד | , |
| 1:4.495 | נטוש |  |
| 1:4.496 | ונעזוב |  |
| 1:4.497 | ונאמר |  |
| 1:4.498 | שאין |  |
| 1:4.499 | לו |  |
| 1:4.500 | שם |  |
| 1:4.501 | ולא |  |
| 1:4.502 | כנוי |  |
| 1:4.503 | ונעבוד |  |
| 1:4.504 | את |  |
| 1:4.505 | הפרצופים |  |
| 1:4.506 | והצורות |  |
| 1:4.507 | אשר |  |
| 1:4.508 | לפי |  |
| 1:4.509 | דבריו |  |
| 1:4.510 | נבראו |  |
| 1:4.511 | ונשתלשלו |  |
| 1:4.512 | ממנו | ! |
| 1:4.513 | וכבר |  |
| 1:4.514 | הארכנו |  |
| 1:4.515 | בפנים |  |
| 1:4.516 | בראיות |  |
| 1:4.517 | נכוחות |  |
| 1:4.518 | מרז״ל |  |
| 1:4.519 | בביטול |  |
| 1:4.520 | הדעות |  |
| 1:4.521 | הזרות |  |
| 1:4.522 | האלו |  |
| 1:4.523 | של |  |
| 1:4.524 | הפילסוף |  |
| 1:4.525 | נביא |  |
| 1:4.526 | השקר |  |
| 1:4.527 | הנז | '! |
| 1:4.528 | כי |  |
| 1:4.529 | הוא |  |
| 1:4.530 | מטרת |  |
| 1:4.531 | תורתינו |  |
| 1:4.532 | הקדושה |  |
| 1:4.533 | להרחיקנו |  |
| 1:4.534 | מאמונת |  |
| 1:4.535 | האלילים |  |
| 1:4.536 | בין |  |
| 1:4.537 | גשמיים |  |
| 1:4.538 | בין |  |
| 1:4.539 | רוחניים | . |
| 1:4.540 | ולדעת |  |
| 1:4.541 | כי |  |
| 1:4.542 | ה | ' |
| 1:4.543 | הוא |  |
| 1:4.544 | האלהים |  |
| 1:4.545 | אין |  |
| 1:4.546 | עוד |  |
| 1:4.547 | מלבדו | . |
| 1:4.548 | ואין |  |
| 1:4.549 | לו |  |
| 1:4.550 | לא |  |
| 1:4.551 | אב |  |
| 1:4.552 | ולא |  |
| 1:4.553 | אח |  |
| 1:4.554 | ולא |  |
| 1:4.555 | בן |  |
| 1:4.556 | רבה |  |
| 1:4.557 | פרשת |  |
| 1:4.558 | יתרו | ) |
| 1:4.559 | ולא |  |
| 1:4.560 | קיבל |  |
| 1:4.561 | מלכותו |  |
| 1:4.562 | מאחר |  |
| 1:4.563 | ואינו |  |
| 1:4.564 | עתיד |  |
| 1:4.565 | למסרה |  |
| 1:4.566 | לאחר |  |
| 1:4.567 | כמו |  |
| 1:4.568 | שביארו |  |
| 1:4.569 | רבותינו |  |
| 1:4.570 | ז״ל |  |
| 1:4.571 | רבה |  |
| 1:4.572 | במדבר |  |
| 1:4.573 | שיר |  |
| 1:4.574 | השירים |  |
| 1:4.575 | פרשה |  |
| 1:4.576 | א | ' |
| 1:4.577 | ירושלמי |  |
| 1:4.578 | סנהדרין |  |
| 1:4.579 | פרק |  |
| 1:4.580 | דיני |  |
| 1:4.581 | ממונות | ) |
| 1:4.582 | על |  |
| 1:4.583 | פסוק |  |
| 1:4.584 | אני |  |
| 1:4.585 | ראשון |  |
| 1:4.586 | ואני |  |
| 1:4.587 | אחרון |  |
| 1:4.588 | ומבלעדי |  |
| 1:4.589 | אין |  |
| 1:4.590 | אלהים |  |
| 1:4.591 | ועל |  |
| 1:4.592 | מה |  |
| 1:4.593 | שאמר |  |
| 1:4.594 | נבוכדנצר |  |
| 1:4.595 | וריויה |  |
| 1:4.596 | די |  |
| 1:4.597 | רביעאה |  |
| 1:4.598 | דמי |  |
| 1:4.599 | לבר |  |
| 1:4.600 | אלהין | ! |
| 1:4.601 | וראיתי |  |
| 1:4.602 | להר | ' |
| 1:4.603 | מנחם |  |
| 1:4.604 | די |  |
| 1:4.605 | לונזבנו |  |
| 1:4.606 | בס | ' |
| 1:4.607 | דרך |  |
| 1:4.608 | החיים |  |
| 1:4.609 | שהרחיב |  |
| 1:4.610 | פיו |  |
| 1:4.611 | לדבר |  |
| 1:4.612 | תועה |  |
| 1:4.613 | על |  |
| 1:4.614 | החסיד |  |
| 1:4.615 | הקדוש |  |
| 1:4.616 | בעל |  |
| 1:4.617 | חובת |  |
| 1:4.618 | הלבבות |  |
| 1:4.619 | על |  |
| 1:4.620 | אמרו |  |
| 1:4.621 | בשם |  |
| 1:4.622 | הפילסוף |  |
| 1:4.623 | שאמר |  |
| 1:4.624 | שלא |  |
| 1:4.625 | יוכל |  |
| 1:4.626 | לעבוד |  |
| 1:4.627 | עלת |  |
| 1:4.628 | העלות |  |
| 1:4.629 | רק |  |
| 1:4.630 | נביא |  |
| 1:4.631 | הדור |  |
| 1:4.632 | בטבעו | , |
| 1:4.633 | או |  |
| 1:4.634 | הפילסוף |  |
| 1:4.635 | המובהק |  |
| 1:4.636 | במה |  |
| 1:4.637 | שקנהו |  |
| 1:4.638 | מן |  |
| 1:4.639 | החכמה |  |
| 1:4.640 | אבל |  |
| 1:4.641 | זולתם |  |
| 1:4.642 | עובדים |  |
| 1:4.643 | זולתו |  |
| 1:4.644 | מפני |  |
| 1:4.645 | שאינם |  |
| 1:4.646 | מבינים |  |
| 1:4.647 | נמצא |  |
| 1:4.648 | אלא |  |
| 1:4.649 | מורכב | ! |
| 1:4.650 | ועל |  |
| 1:4.651 | דברי |  |
| 1:4.652 | החסיד |  |
| 1:4.653 | האלו |  |
| 1:4.654 | התקצף |  |
| 1:4.655 | וחשב |  |
| 1:4.656 | כי |  |
| 1:4.657 | כאן |  |
| 1:4.658 | מקום |  |
| 1:4.659 | תלונה |  |
| 1:4.660 | עליו |  |
| 1:4.661 | ותואנה |  |
| 1:4.662 | מצא |  |
| 1:4.663 | להתעות |  |
| 1:4.664 | השומעים |  |
| 1:4.665 | ולהבהילם |  |
| 1:4.666 | באמרו |  |
| 1:4.667 | וכי |  |
| 1:4.668 | פלוני |  |
| 1:4.669 | ופלוני |  |
| 1:4.670 | מחכמינו |  |
| 1:4.671 | האחרונים |  |
| 1:4.672 | המפורסמים |  |
| 1:4.673 | עובדים |  |
| 1:4.674 | זולתו | ? |
| 1:4.675 | והוא |  |
| 1:4.676 | ספרם |  |
| 1:4.677 | בלשונו |  |
| 1:4.678 | אחד |  |
| 1:4.679 | אחד |  |
| 1:4.680 | בשמותם | . |
| 1:4.681 | ולא |  |
| 1:4.682 | שם |  |
| 1:4.683 | אל |  |
| 1:4.684 | לבו |  |
| 1:4.685 | כי |  |
| 1:4.686 | הרבה |  |
| 1:4.687 | מהחכמים |  |
| 1:4.688 | שַמָנָה |  |
| 1:4.689 | הלכו |  |
| 1:4.690 | בעצימות |  |
| 1:4.691 | עינים |  |
| 1:4.692 | והאמינו |  |
| 1:4.693 | לדברי |  |
| 1:4.694 | מחבר |  |
| 1:4.695 | הזהר |  |
| 1:4.696 | אשר |  |
| 1:4.697 | כינה |  |
| 1:4.698 | את |  |
| 1:4.699 | עצמו |  |
| 1:4.700 | בשם |  |
| 1:4.701 | רשב״י |  |
| 1:4.702 | התנא |  |
| 1:4.703 | כדי |  |
| 1:4.704 | לתלות |  |
| 1:4.705 | נבואותיו |  |
| 1:4.706 | השקריות |  |
| 1:4.707 | בשמו |  |
| 1:4.708 | של |  |
| 1:4.709 | התנא | , |
| 1:4.710 | ודבר |  |
| 1:4.711 | סרה |  |
| 1:4.712 | על |  |
| 1:4.713 | ה | ' |
| 1:4.714 | אלהינו |  |
| 1:4.715 | שנגלה |  |
| 1:4.716 | לו |  |
| 1:4.717 | בבית |  |
| 1:4.718 | מדרשו |  |
| 1:4.719 | עתיקא |  |
| 1:4.720 | קדישא |  |
| 1:4.721 | וגילה |  |
| 1:4.722 | לו |  |
| 1:4.723 | סודות |  |
| 1:4.724 | נעלמים |  |
| 1:4.725 | וכוחות |  |
| 1:4.726 | נעלמות |  |
| 1:4.727 | זהר |  |
| 1:4.728 | בראשית |  |
| 1:4.729 | כ״ב | ) |
| 1:4.730 | כי |  |
| 1:4.731 | הבורא |  |
| 1:4.732 | הוא |  |
| 1:4.733 | האליל |  |
| 1:4.734 | הנקרא |  |
| 1:4.735 | אבא |  |
| 1:4.736 | שהיה |  |
| 1:4.737 | מצוה |  |
| 1:4.738 | לזוגתו |  |
| 1:4.739 | האלילה |  |
| 1:4.740 | אימא |  |
| 1:4.741 | והיא |  |
| 1:4.742 | עושה |  |
| 1:4.743 | כל |  |
| 1:4.744 | אשר |  |
| 1:4.745 | יאמר |  |
| 1:4.746 | אליה |  |
| 1:4.747 | חוץ |  |
| 1:4.748 | מבריאת |  |
| 1:4.749 | אדם |  |
| 1:4.750 | הראשון |  |
| 1:4.751 | שהיא |  |
| 1:4.752 | אימא |  |
| 1:4.753 | שאמרה |  |
| 1:4.754 | ועשתה |  |
| 1:4.755 | את |  |
| 1:4.756 | אדם |  |
| 1:4.757 | הראשון |  |
| 1:4.758 | שלא |  |
| 1:4.759 | ברצון |  |
| 1:4.760 | אבא | ! |
| 1:4.761 | ובפרשת |  |
| 1:4.762 | וירא |  |
| 1:4.763 | דף |  |
| 1:4.764 | קי״ב |  |
| 1:4.765 | וכן |  |
| 1:4.766 | בפרשת |  |
| 1:4.767 | בלק |  |
| 1:4.768 | דף |  |
| 1:4.769 | קצ״א |  |
| 1:4.770 | מיחס |  |
| 1:4.771 | עבודתינו |  |
| 1:4.772 | ותפילותינו |  |
| 1:4.773 | לאל |  |
| 1:4.774 | זר |  |
| 1:4.775 | קצר |  |
| 1:4.776 | אפים |  |
| 1:4.777 | זעיר |  |
| 1:4.778 | אנפין | ) |
| 1:4.779 | אשר |  |
| 1:4.780 | לא |  |
| 1:4.781 | עשה |  |
| 1:4.782 | את |  |
| 1:4.783 | השמים |  |
| 1:4.784 | ואת |  |
| 1:4.785 | הארץ | . |
| 1:4.786 | וקורא |  |
| 1:4.787 | הוא |  |
| 1:4.788 | לעצבים |  |
| 1:4.789 | הללו |  |
| 1:4.790 | בשם |  |
| 1:4.791 | ה | ' |
| 1:4.792 | אלהינו | . |
| 1:4.793 | ועל |  |
| 1:4.794 | זעיר |  |
| 1:4.795 | אנפין |  |
| 1:4.796 | קרא |  |
| 1:4.797 | ירמיה |  |
| 1:4.798 | הנביא |  |
| 1:4.799 | ע״ה |  |
| 1:4.800 | כדנה |  |
| 1:4.801 | תימרון |  |
| 1:4.802 | להום |  |
| 1:4.803 | אלהיא |  |
| 1:4.804 | די |  |
| 1:4.805 | שמיא |  |
| 1:4.806 | וארקא |  |
| 1:4.807 | לא |  |
| 1:4.808 | עבדו |  |
| 1:4.809 | ייבדו |  |
| 1:4.810 | מן |  |
| 1:4.811 | ארעא |  |
| 1:4.812 | ומן |  |
| 1:4.813 | תחות |  |
| 1:4.814 | שמיא |  |
| 1:4.815 | אלה | . |
| 1:4.816 | שהרי |  |
| 1:4.817 | מבואר |  |
| 1:4.818 | בזהר |  |
| 1:4.819 | בכמה |  |
| 1:4.820 | מקומות |  |
| 1:4.821 | שהוא |  |
| 1:4.822 | נברא |  |
| 1:4.823 | ולא |  |
| 1:4.824 | בורא |  |
| 1:4.825 | עיין |  |
| 1:4.826 | זהר |  |
| 1:4.827 | בראשית |  |
| 1:4.828 | כ״ב |  |
| 1:4.829 | וביושר |  |
| 1:4.830 | לבב |  |
| 1:4.831 | דף |  |
| 1:4.832 | ג | ' |
| 1:4.833 | ע״ב | ) |
| 1:4.834 | ועכ״ז |  |
| 1:4.835 | הוא |  |
| 1:4.836 | אומר |  |
| 1:4.837 | לעבדו |  |
| 1:4.838 | ולברכו |  |
| 1:4.839 | ולקבלו |  |
| 1:4.840 | עליו |  |
| 1:4.841 | באלוה |  |
| 1:4.842 | ולירא |  |
| 1:4.843 | מזוגתו | ! |
| 1:4.844 | והרבה |  |
| 1:4.845 | מרבותינו |  |
| 1:4.846 | האחרונים |  |
| 1:4.847 | שהזכיר |  |
| 1:4.848 | בספר |  |
| 1:4.849 | דרך |  |
| 1:4.850 | החיים |  |
| 1:4.851 | בתמיהתו |  |
| 1:4.852 | על |  |
| 1:4.853 | החסיד |  |
| 1:4.854 | בעל |  |
| 1:4.855 | חובת |  |
| 1:4.856 | הלבבות |  |
| 1:4.857 | הלכו |  |
| 1:4.858 | בעצימת |  |
| 1:4.859 | עין |  |
| 1:4.860 | ונמשכו |  |
| 1:4.861 | אחריו |  |
| 1:4.862 | והאמינו |  |
| 1:4.863 | באמונתו |  |
| 1:4.864 | וקיבלו |  |
| 1:4.865 | האליל |  |
| 1:4.866 | קצר |  |
| 1:4.867 | אפים |  |
| 1:4.868 | ז״א | ) |
| 1:4.869 | עליהם |  |
| 1:4.870 | באלוה | ! |
| 1:4.871 | וכל |  |
| 1:4.872 | זה |  |
| 1:4.873 | גרם |  |
| 1:4.874 | להם |  |
| 1:4.875 | על |  |
| 1:4.876 | שלא |  |
| 1:4.877 | התבוננו |  |
| 1:4.878 | בחכמת |  |
| 1:4.879 | יוצר |  |
| 1:4.880 | הכל |  |
| 1:4.881 | שהוא |  |
| 1:4.882 | לחקור |  |
| 1:4.883 | ולבחון |  |
| 1:4.884 | במעשה |  |
| 1:4.885 | בראשית |  |
| 1:4.886 | המבוארים |  |
| 1:4.887 | בדברי |  |
| 1:4.888 | הר | ' |
| 1:4.889 | מורה |  |
| 1:4.890 | נבוכים |  |
| 1:4.891 | ובדברי |  |
| 1:4.892 | החסיד |  |
| 1:4.893 | בעל |  |
| 1:4.894 | חובת |  |
| 1:4.895 | הלבבות | . |
| 1:4.896 | ורבינו |  |
| 1:4.897 | סעדיה |  |
| 1:4.898 | גאון |  |
| 1:4.899 | ור״י |  |
| 1:4.900 | הלוי |  |
| 1:4.901 | בכוזרי |  |
| 1:4.902 | כמו |  |
| 1:4.903 | שאמר |  |
| 1:4.904 | הנביא |  |
| 1:4.905 | ע״ה |  |
| 1:4.906 | ואת |  |
| 1:4.907 | פעל |  |
| 1:4.908 | ה | ' |
| 1:4.909 | לא |  |
| 1:4.910 | יביטו | . |
| 1:4.911 | כי |  |
| 1:4.912 | חשבו |  |
| 1:4.913 | דבריהם |  |
| 1:4.914 | לחכמה |  |
| 1:4.915 | יונית |  |
| 1:4.916 | ורחקו |  |
| 1:4.917 | ממנה |  |
| 1:4.918 | הרחק |  |
| 1:4.919 | כמטחוי |  |
| 1:4.920 | קשת | , |
| 1:4.921 | וגם |  |
| 1:4.922 | לא |  |
| 1:4.923 | שמו |  |
| 1:4.924 | אל |  |
| 1:4.925 | לבם |  |
| 1:4.926 | דברי |  |
| 1:4.927 | אבותינו |  |
| 1:4.928 | ז״ל |  |
| 1:4.929 | האמתיים |  |
| 1:4.930 | במדרשים |  |
| 1:4.931 | שלא |  |
| 1:4.932 | הכיר |  |
| 1:4.933 | אברהם |  |
| 1:4.934 | אבינו |  |
| 1:4.935 | את |  |
| 1:4.936 | בוראו |  |
| 1:4.937 | אלא |  |
| 1:4.938 | בהתבוננו |  |
| 1:4.939 | בפעל |  |
| 1:4.940 | ה | ' |
| 1:4.941 | וחקר |  |
| 1:4.942 | ודרש |  |
| 1:4.943 | בהם |  |
| 1:4.944 | והביא |  |
| 1:4.945 | מופתים |  |
| 1:4.946 | מהם |  |
| 1:4.947 | עד |  |
| 1:4.948 | שהשיג |  |
| 1:4.949 | האמת |  |
| 1:4.950 | ואז |  |
| 1:4.951 | נגלה |  |
| 1:4.952 | לו |  |
| 1:4.953 | הקב״ה |  |
| 1:4.954 | ואמר |  |
| 1:4.955 | לו |  |
| 1:4.956 | אהבת |  |
| 1:4.957 | צדק |  |
| 1:4.958 | ותשנא |  |
| 1:4.959 | רשע |  |
| 1:4.960 | וכו | ' |
| 1:4.961 | משל |  |
| 1:4.962 | לאחד |  |
| 1:4.963 | שהיה |  |
| 1:4.964 | מהלך |  |
| 1:4.965 | בדרך |  |
| 1:4.966 | וראה |  |
| 1:4.967 | בירה |  |
| 1:4.968 | וכו | ' |
| 1:4.969 | מדרש |  |
| 1:4.970 | הגדול |  |
| 1:4.971 | ריש |  |
| 1:4.972 | פרשת |  |
| 1:4.973 | לך |  |
| 1:4.974 | לך | ) |
| 1:4.975 | ועל |  |
| 1:4.976 | כרחנו |  |
| 1:4.977 | נודה |  |
| 1:4.978 | על |  |
| 1:4.979 | האמת |  |
| 1:4.980 | שדברי |  |
| 1:4.981 | החסיד |  |
| 1:4.982 | בעל |  |
| 1:4.983 | חובת |  |
| 1:4.984 | הלבבות |  |
| 1:4.985 | ברור |  |
| 1:4.986 | מללו |  |
| 1:4.987 | שלא |  |
| 1:4.988 | יוכל |  |
| 1:4.989 | לעבוד |  |
| 1:4.990 | את |  |
| 1:4.991 | אלהים |  |
| 1:4.992 | אמת |  |
| 1:4.993 | אלא |  |
| 1:4.994 | וכו | ' |
| 1:4.995 | אבל |  |
| 1:4.996 | זולתם |  |
| 1:4.997 | עובדים |  |
| 1:4.998 | זולתו |  |
| 1:4.999 | כמאמר |  |
| 1:4.1000 | הנביא |  |
| 1:4.1001 | וימים |  |
| 1:4.1002 | רבים |  |
| 1:4.1003 | לישראל |  |
| 1:4.1004 | ללא |  |
| 1:4.1005 | אלהי |  |
| 1:4.1006 | אמת | , |
| 1:4.1007 | וכמו |  |
| 1:4.1008 | שעינינו |  |
| 1:4.1009 | רואות |  |
| 1:4.1010 | ואזנינו |  |
| 1:4.1011 | שומעות |  |
| 1:4.1012 | עד |  |
| 1:4.1013 | שאפשר |  |
| 1:4.1014 | שיאמר |  |
| 1:4.1015 | אדם |  |
| 1:4.1016 | לחבירו | ! |
| 1:4.1017 | עַמְּךָ |  |
| 1:4.1018 | עַמִּי |  |
| 1:4.1019 | וְאלֹהֶיך |  |
| 1:4.1020 | לֹא |  |
| 1:4.1021 | אֱלהָי | ! |
| 1:4.1022 | ואל |  |
| 1:4.1023 | אלהי |  |
| 1:4.1024 | אמת |  |
| 1:4.1025 | נשפוך |  |
| 1:4.1026 | שיחנו |  |
| 1:4.1027 | לנחותינו |  |
| 1:4.1028 | במעגלי |  |
| 1:4.1029 | יושר |  |
| 1:4.1030 | ולהצילנו |  |
| 1:4.1031 | מכל |  |
| 1:4.1032 | אמונה |  |
| 1:4.1033 | זרה |  |
| 1:4.1034 | אכי״ר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

