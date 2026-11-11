# Translation pack: milhamot-hashem perek-4-02 (Milchamot ha-Shem 4:2–4:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-02`.

## Rulings for this unit

Extracted from `works/milhamot-hashem/decisions.md` for the words of this unit; the rulings below are the whole of what you need. Read the full file only if a question is not answered here.

# Milchamot ha-Shem: decisions for every unit

> Owner answered Q4 and Q7 on 2026-11-10; Q1-3, Q5-6 stand at the defaults (see "Owner's answers and defaults" at the end).

Editorial decisions set before translation. A reviewer may question one in a
`question` patch; the project owner decides. The rulings below come from the
term's actual uses in `stage/corpus.json` (references are chapter:paragraph).

General rule for technical terms: the English names the thing the words name
and goes no further. Kabbalistic technical names are transliterated, because
every English equivalent in use takes a side. Either it carries the
Kabbalists' defence ("configuration", "emanation", "the Infinite", "Divine
Presence") or it carries Kafich's charge against them ("god", "idol"). Kafich
makes his charge in his own words ("אליל", "אלוהות", "אל זר"), and the English
translates those words where he uses them.

Transliteration: the project's plain scheme, as in Avot and Genesis (Shetach,
Chayei Sarah, Lekh Lekha): no diacritics; ח = ch, כ = kh, צ = tz, ק = k.
The author's name follows it: Yichya Kafich (owner's ruling, 2026-11-10; the
conventional "Yihya Kafich" is named once in the edition's front matter).

## Output format for this work (prose)

Read this whole file before translating. A segment is a paragraph of prose.
The pack has no morphemes, lemmas, lexicon or commentary, so:

- `tokens`: one entry per word: `id`, `gloss`, and a `note` where a learner
  needs one. No `sense`, no `lex`.
- `commentary`: `[]`. `context`: only for a realia a reader needs (a Yemenite
  custom, a book, a person); say what is established; cite no modern work.
- `grammar`: 2 to 6 notes per paragraph; in a long paragraph, notes on the
  hardest sentences, not on all. `literal`: a word-for-word line.
- `phrases` divide every token of the paragraph, in order. In a long
  paragraph make a phrase a clause, not a word, so that the file stays
  checkable; each phrase's `english` must occur in the paragraph's `english`.
- Quoted passages (Bible, Talmud, Zohar, Judeo-Arabic) follow the Quotations
  section; Bible verses are translated from the Hebrew of the verse as the
  author quotes it and the reference goes in `apparatus`.
- Every Hebrew abbreviation is expanded in `gloss` (the reading line carries
  the full word).

## Names and titles

- The book: מלחמות ה׳ → *The Wars of the LORD* (Num 21:14, "ספר מלחמת ה'",
  the form of the title page, 1:11). Transliterated title: *Milchamot ha-Shem*.
  In 1:5 the author explains the name: "וקראתי שם הספר מלחמות ה' על יחודו" →
  "I named the book The Wars of the LORD, for His unity". (See open question 1.)
- The author's signature (1:9) "הצעיר יחיה בן סלי' אלקאפח יצ״ו" → "The humble
  Yichya son of Sali' al-Kafich, may his Rock keep him and give him life".
  Keep the abbreviated patronymic as it is written (open question 6). הקדמת
  המחבר: "The Author's Introduction". הקדמת המו"ל / המו"ל: "The Publisher's
  Introduction" / "The publisher": 1:10–1:13 are in the publisher's voice,
  not the author's.
- רשב"י → "Rabbi Shimon bar Yochai", always written out. רשב"י התנא →
  "Rabbi Shimon bar Yochai the Tanna". The author separates the Tanna from
  the writer who "named himself" after him (1:4: "אשר כינה את עצמו בשם רשב״י
  התנא"). Keep exactly the qualifier he gives. Do not add "pseudo-" or
  "the so-called", and do not settle in the English which of the two he
  means.
- מחבר הזהר → "the author of the Zohar". Never name him (no "Moses de León"):
  the text names no one. His epithets are translated as they stand and stay
  fixed: הפילסוף/הפילוסוף → "the philosopher"; נביא השקר → "the false
  prophet"; המסית → "the enticer"; המדיח → "the seducer". The last three are
  legal categories of Deut 13 and 18, and the author builds a legal case on
  them ("ובמיתתו בחנק", 3:1). Do not soften them to "deceiver" or "charlatan".
- ה׳ (the Tetragrammaton as written): "the LORD", in the author's prose and
  in verses. Where it is cited as a name ("זעיר אנפין דאיקרי ה'", 4:11), write
  "who is called 'the LORD'". Not ה׳ as Tetragrammaton: "מאת ה' אליעזר בן
  יהודה" (1:2) → "from Mr. Eliezer Ben-Yehuda"; ה׳ as a numeral (5) or a page
  letter → the number. The pack splits it into the token ה plus the
  punctuation ', so read the token in context.
- הקב"ה / קב"ה, and קודשא בריך הוא in the Aramaic → "the Holy One, blessed be
  He". As a name under discussion: "the name 'the Holy One, blessed be He'"
  ("שם הקב"ה המורגל בפינו", 4:17).
- Honorifics: ז"ל → "of blessed memory". ע"ה → "peace be upon him". רז"ל →
  "our Rabbis of blessed memory". חז"ל → "our Sages of blessed memory". זלה"ה
  → "may his memory be for the life of the world to come". זצוק"ל → "may the
  memory of the righteous and holy be for a blessing". ת"ו (of Jerusalem) →
  "may it be rebuilt and established". יר"ה (of the Imam, 8:11) → "may his
  majesty be exalted". Translate every one, including the honorifics the
  author gives to opponents. Add none.
- Authorities: keep the acronyms that English uses as names: Rashi, the
  Rambam, the Ramban, the Ari, the Ribash, the Rashba. Write out the rest:
  רס"ג/הרס"ג → "Rabbi Saadia Gaon"; הרח"ו/מהרח"ו → "Rabbi Chayyim Vital";
  בעל חובת הלבבות → "the author of Chovot ha-Levavot". הר' המורה / רבינו, where
  they mean the Rambam → "the Master, the Guide" / "our Master". Translate
  what is written, and say in a note who is meant.
- Book titles: transliterate them (Kisse Eliyahu, Yosher Levav, Etz Chayyim,
  Mikdash Melekh, Emunot ve-De'ot, Moreh Nevukhim, the Kuzari, Sefer ha-Berit,
  Tikkunei Zohar, Zohar Chadash, Idra of Naso), with the meaning in a note at
  the first occurrence in a unit. Abbreviations are written out (כס"א →
  Kisse Eliyahu). Bible books take their English names. Tractates are
  transliterated.
- Peoples and places: תימן → "Yemen"; צנעא → "Sanaa". ממשלת תוגרמה →
  "the government of Turkey", with a note that the Hebrew name is Togarmah.
  הישמעאלים → "the Ishmaelites" (the note says "Muslims"); המושלמנים →
  "the Muslims". הנוצרים → "the Christians". השונים → "the Dualists" (an
  Arabic-style calque, *al-thanawiyya*). האמגושים → "the Magians".
  מאמיני השלוש → "the believers in the Trinity".

## Terms

- **אין סוף / א"ס**: fixed, "Ein Sof", as a name with no article: "האין סוף
  העיקרי" → "the principal Ein Sof"; "כמה מיני אין סוף" (5:2) → "several
  kinds of Ein Sof"; "א"ס ב"ה" → "Ein Sof, blessed be He". Not "the
  Infinite": that makes it a philosophical attribute the author accepts,
  where in the book it is "the First Cause, which they call Ein Sof"
  ("הסבה הראשונה שהוא נקרא אצלם א"ס", 2:11). Give the literal "without end"
  in a note at the first occurrence in a unit.
- **פרצוף / פרצופים**: "partzuf / partzufim" where it means the Kabbalistic
  entities: "פרצוף אחרון שבעולם האצילות" (2:6) → "the last partzuf of the
  World of Emanation"; "חמשה פרצופים" → "five partzufim". In the generic
  sense, paired with צורה or תמונה, it is "face": "אלילי תמונות ופרצופים
  בעלי מפלצות" (1:4) → "idols of images and faces of monstrous form"; "שאין
  לעבוד שום פרצוף וצורה" (2:11) → "that no face or form is to be worshipped".
  Where the author plays on both senses, add a note. Never "configuration" or
  "countenance-structure": that is the defenders' word for them.
- **The names of the partzufim**: transliterated, with the literal sense in a
  note at the first occurrence in a unit. Ze'ir Anpin (Small Face; ז"א →
  "Ze'ir Anpin"; זעיר alone → "Ze'ir"); Arikh Anpin (Long Face); Abba
  (Father); Imma (Mother); Nukva, נוקביה → "his Nukva" (Female); Atik and
  עתיקא קדישא → "Atika Kadisha" (Holy Ancient One); עתיקא סתימא → "Atika
  Setima" (Hidden Ancient One). עתיק יומין: as a partzuf, "Atik Yomin"; in
  Dan 7:9, "the Ancient of Days". The descriptive epithets are translated:
  סבא דסבין → "the Elder of Elders"; עלת העלות → "the Cause of causes";
  הסבה הראשונה → "the First Cause"; כתר עליון → "Keter Elyon (the Supreme
  Crown)". Where the author argues from the meaning of a name ("שאין לי
  אב", "בן לאבא ואימא"), say so in the note and do not translate the name in
  the English.
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
- **Other recurrent terms**, fixed: השפע → "the flow"; בחינה → "aspect";
  השתלשלות → "the chain of descent"; עלה / עלול → "cause / effect"; מחויב
  המציאות → "the Necessary Existent"; כלים → "vessels"; אורות דקים → "subtle
  lights"; גוף דק → "a subtle body"; כוונה / כוונות (in prayer) →
  "intention / intentions"; זווג / זווגים → "coupling / couplings" (literal:
  neither the Kabbalists' "union" nor a coarser word); השגחה → "providence";
  תפיסה ("לית מחשבה תפיסה ביה") → "no thought grasps Him"; מין / מינות →
  "heretic / heresy"; אפיקורוס → "apikoros"; כופר בעיקר → "denies the
  fundamental principle"; ספר חיצון → "an external book" (m. Sanh. 10:1 in a
  note); תורה שבעל פה → "the Oral Torah"; ימי הבינים → "the Middle Ages";
  "אין לו חלק לעולם הבא" → "has no share in the world to come".
- **Ambiguous abbreviations.** Read each one in context, and name the
  expansion in the token `note`:
  - ר"ל: "that is to say" (רוצה לומר); "Resh Lakish" ("אמר ר"ל", 3:1, 6:33);
    "the Merciful One save us" (רחמנא ליצלן, at a sentence end, 8:22).
  - ע"כ: "thus far" after a quotation; "therefore" (על כן); "perforce"
    (על כרחך); "and upon all" (ועל כל, 1:12 "וע"כ עמו ישראל").
  - ע"ש: "see there" (עיין שם); "on account of" (על שם, 5:1 "ע"ש
    התלבשותו").
  - ח"ו → "God forbid"; וכו' → "etc." (always kept: it marks a truncated
    quotation); ג"כ → "also"; כ"ש / ק"ו → "all the more so"; ת"ל → "Scripture
    says"; א"כ → "if so".

## Quotations

- **Scope.** Everything quoted is translated into English from the words the
  author gives: Zohar Aramaic, the Bavli and Yerushalmi, Targum, Biblical
  Aramaic (Jer 10:11 and Dan 3:25 in 1:4), Midrash, the medieval authors, and
  Judeo-Arabic. No Hebrew, Aramaic or Arabic stays in `english` (the check
  refuses Hebrew letters).
- **Marking.** An explicit quotation (introduced by וז"ל, כתב, אמרו, גרסינן,
  שנאמר, כדכתיב, or closed by עכ"ל / ע"כ) goes in double quotation marks. A
  quotation inside it goes in single quotation marks. There are no
  italics and no bracketed labels in `english`: the field is plain text.
  The language goes in `grammar` ("Zohar Aramaic", "Bavli Aramaic",
  "Judeo-Arabic"), as in Avot. The word-for-word sense goes in the segment's
  `literal` and the phrases' `literal`. Do not put a literal beside it in
  `english`.
- **Formulas of the quotation** are translated, because they carry the
  quotation's boundaries: וז"ל → "and these are his words:"; וז"ל הזהר →
  "and these are the words of the Zohar:"; עכ"ל → "Thus far his words."; ע"כ
  after a quotation → "Thus far."; "ע"כ לשון הזהר עם פי' הרש"ב בגליון" (4:14)
  → "Thus far the wording of the Zohar, with the explanation of the RaShB in
  the margin."
- **Glosses inside a quotation.** Parentheses inside a quotation ("(פירש
  בגליון …)", "(פי' …)", "(לאלהי ישראל)", "(לפי דבריהם)") stay in parentheses
  where they stand, and are translated: "(the margin explains: …)",
  "(explanation: …)". Do not fold them into the quoted text, and do not move
  them out.
- **Zohar formulas**, fixed: זכאה חולקיה / חולקיהון → "happy is his / their
  portion"; זכאה חולקיהון דישראל → "happy is the portion of Israel"; ווי ליה
  מאן דלא … → "woe to him who does not …"; דמאן דעביד כך וכך → "of one who
  does such-and-such"; תא חזי → "come and see"; רזא דמלה → "the secret of the
  matter"; פתח ר' שמעון → "Rabbi Shimon opened". In 9:4 and 9:16 the author
  quotes them as the Zohar's sales talk ("משבח את סחורתו"), so put them in
  quotation marks. The irony is in his framing, and it is not to be added
  to the formula. Variant spellings in the witness ("הולקיה" 9:4, "דכאה"
  9:16) are translated as the evident word and recorded in `apparatus`. Do
  not correct them silently.
- **Exegesis that turns on a Hebrew word.** Where an exegesis turns on a
  Hebrew word (4:14 את / ואותו; 4:11 אם אין → "Ayin"), keep the word
  transliterated in the English ("it is not written 'to the LORD your God' but
  'et the LORD your God'. What is 'et'?") and explain it in `grammar`.
- **Abbreviated or paraphrased quotation.** Translate only what is on the
  page. Do not supply what וכו' omits or restore a passage from its source. A
  report without a quotation formula ("כתב ש…", "כמו שכתב…", "ופי' מהר"י
  לופיס שכוונת…") is indirect speech, without quotation marks. Where the
  author changes a verse or a source ("אשר לא שערום אבותיהם", 9:10, against
  Deut 32:17 "אבתיכם"), the English follows his wording and `apparatus` gives
  the source.
- **Bible verses**, whether quoted or embedded as allusion (*melitzah*), are
  translated from the Hebrew of the verse as it stands in the text, in plain
  modern English and not King James. An embedded allusion takes no quotation
  marks. `apparatus` gives the verse for every verse the reviewer can verify.
  The verse must be recognised wherever it occurs, because the irony lies in
  where it is applied:
  - "בחלק שפתיו" (1:4, 9:16; Prov 7:21, the adulteress) → "with the
    smoothness of his lips";
  - "כמתלהלה היורה זיקים חצים ומות" (9:4; Prov 26:18) → "like a madman who
    shoots firebrands, arrows and death";
  - "ראיות צנומות דקות שדופות קדים" (7:26; Gen 41:23) → "proofs withered, thin,
    blasted by the east wind";
  - "בחבלי השוא … כעבות העגלה" (3:3; Isa 5:18) → "with cords of falsehood …
    as with cart-ropes";
  - "כעיר פרוצה אין חומה" (9:10; Prov 25:28) → "like a city breached, without
    a wall";
  - "להוציא ישן … מפני חדש" (9:10; Lev 26:10) → "to clear out the old before the
    new";
  - "מילדי נכרים" (9:16; Isa 2:6) → "from the children of strangers";
  - "הבל וריק" (Isa 30:7) → "vanity and emptiness".
  The verses the argument rests on keep one rendering in every unit:
  - Deut 6:4: "Hear, O Israel: the LORD our God, the LORD is one." (Four units
    in order, which the Zohar reading in 6:1 needs.)
  - Deut 32:39: "See now that I, I am He, and there is no god with Me."
  - Isa 44:6: "I am the first and I am the last, and besides Me there is no
    God."
  - Exod 20:2–3: "I am the LORD your God … You shall have no other gods
    before Me."
  - Exod 22:19: "He who sacrifices to the gods shall be destroyed, save to the
    LORD alone."
  - Deut 6:13: "The LORD your God you shall fear, and Him you shall worship."
  - Exod 34:6: "the LORD, the LORD, God merciful and gracious, slow to
    anger".
  - Jer 10:11 (Aramaic): "Thus shall you say to them: the gods that did not
    make the heavens and the earth shall perish from the earth and from under
    these heavens."
  Where the text has the ketiv ("מי אתי", Isa 44:24), translate the ketiv and
  note the qere.
- **Talmudic and Aramaic idioms in the author's own Hebrew** ("דמעיילי פילא
  בקופא דמחטא" 1:5, "ואפריין נמטיה" 1:2, "הני מילי דמרפסן איגרא" 2:12, "בהא
  סלקינן ובהא נחתינן" 6:30, "בנדון דידן") are translated by their sense:
  "who bring an elephant through the eye of a needle", "and we acknowledge
  our debt to him", "words that split roofs", "this is where we come out",
  "in our case". The literal sense goes in the phrase's `literal`. Where the
  author glosses an Aramaic saying himself ("תעלא בעידניה סגיד ליה, ר"ל: אדם
  פחות…", 2:16), translate both, and do not merge them.
- **The medieval authors** (the Rambam, Saadia, Halevi, Bahya, the SeMaG,
  Albo, the Roke'ach, Shevilei Emunah) are translated from the Hebrew the
  author quotes. Do not translate them from the Arabic originals, from the
  standard Hebrew editions, or from published English translations (Pines,
  Hyamson, Rosenblatt, Hirschfeld). Where his Hebrew differs from the known
  text, the English follows his Hebrew. A difference may be noted in
  `apparatus` only if it is certain.
- **Judeo-Arabic** (7:24 Saadia on Job; 5:3 Choter ben Shelomo; 9:10 a Muslim
  polemic against the Trinity, "ספר אצהאר אלחק", called "גילוי האמת" in 10:9;
  the one-line proverb in 2:13, "ואן כאן פי בוצע ולא יחסן יער"). It is
  translated from the Arabic as it stands. The Hebrew or Aramaic glosses the
  author puts inside the Arabic ("אתתלית (תלת קשרי דמהימנותא)", 9:10, which
  glosses the Arabic "trinity" with a Zohar phrase; "(העתק ועיין בסוף
  הספר)") stay in parentheses and are translated as his. The title is
  transliterated as each place writes it ("Iẓhar al-Chaqq", "Gilui ha-Emet"),
  with a note that the two are the same book. Its author is not named.
  The author's own Hebrew translations (10:7–10:9, and the bracketed "[תרגום
  הרב המחבר: …]" in 5:3) are separate segments: each is translated from its
  Hebrew, and neither is copied onto the other. Where the two differ in sense,
  record it in `apparatus`. "כת המיחדים" (10:7) = Arabic "אהל אלתוחיד" (7:24)
  → "the upholders of the unity", in both. (See open question 4.)
- **References** are given exactly as the author gives them, where he gives
  them, in English form: דף → "folio"; ע"א / ע"ב → "a / b"; ח"א / ח"ב / ח"ג
  → "vol. 1 / 2 / 3"; פ' / פרשת → the parashah, transliterated ("Zohar,
  Balak, folio 191b"); "זהר בראשית כ״ב" → "Zohar, Bereshit 22", with no
  "folio" where he writes none; סי' → "section"; שער / ענף / מאמר → "Gate /
  Branch / Treatise"; tractate and folio → "Sanhedrin 99". Numbers are
  written in Arabic numerals. Do not add a reference he lacks, complete a
  partial one, convert it to a modern edition, or correct it: "בשלח דף ס"ד
  (ח"ב סז, א)" (4:11) stays "Be-shallach, folio 64 (vol. 2, 67a)". Never import
  the references the 2015 PDF adds. For Zohar and Talmud passages without a
  reference, nothing is added, even in `apparatus`. Bible verses are the
  only exception, and they go in `apparatus`.
- **Section numbers.** The author's numbered sections run through the whole
  book (ב) at 3:1 to קמד) at 10:6). The opening label → "§2." etc. "לעיל
  סי' מ"ב" / "לקמן סי' נ"ח" (this book) → "above, §42" / "below, §58".
  Irregular numbers are translated as written and recorded: "קיה)" (8:22)
  → "§115", noting the letters; the gap from קכט) to קלא) (9:12–9:13) is
  noted, not repaired.

## Style

- Plain, exact, modern English. Keep the author's register by translating the
  words at their own strength: אליל "idol", נביא השקר "false prophet", הבל
  וריק "vanity and emptiness", סכלות והוללות "folly and madness". Add no
  sarcasm of your own: no scare quotes, "so-called", "sic", or added "!".
  Keep every "!" and every rhetorical question the Hebrew has. Keep his
  disclaimers, which carry the irony: "(לפי דבריו)" → "(according to his
  words)"; "(לפי דעתם)" → "(in their opinion)"; "רחמנא ליצלן" → "the
  Merciful One save us".
- Keep the forms of address and the voices. "ידידי" → "my friend"; "קורא
  נעים" → "pleasant reader"; "אחי" → "my brother". Second-person address is
  kept. Chapter 2 is a correspondence (שאול שאל → "Question"; תשובה →
  "Reply"; חזר השואל והשיב → "The questioner replied again"; המשיב → "the
  respondent"). Translate each letter in its writer's voice, and do not
  identify the parties beyond what the text says.
- Long Hebrew periods may be divided into English sentences, but no clause is
  dropped, no connective is added that changes the logic, and the order of
  the argument is kept.
- Numbers and dates: letter numerals → Arabic numerals. A year without the
  thousands → the full year ("תר"ע" → "5670"; "התר"ן לפרט גדול" → "5650 of
  the full count"). The civil year goes in `apparatus`. Where gematria is the
  point, give the number in the English and the letters and words in a note.
- Readings from the witnesses: square brackets in the base text
  ("[מתחלת באלוהיות]") are the transcriber's corrections. Translate the
  corrected reading and record both. "(חסר)" → "(missing)": it is a lacuna
  sign, not text. "(חטר)" in 10:8 is a variant of it, to be recorded. Rule 3
  applies to every Wikisource/PDF difference that bears on the English.
- Book-internal Hebrew of 1931 that is not technical (בית הספר, הסתדרות,
  מכתבי עתים) takes its ordinary modern sense: "school", "organisation",
  "newspapers".

## Rulings from the pilot (perek-1-05, 2-07, 4-10), by the editor

- Ambiguous Judeo-Arabic spelling: translate the evident word, say so in
  `note`, and give the alternative in `apparatus`.
- A direct quotation with no introductory formula takes quotation marks when
  the source is certain; when it is not, no marks and a note.
- Book titles keep their conventional English spelling (Kisse Eliyahu, Sefer
  ha-Kavanot, Machberet ha-Kodesh); the scheme (kh for a כ without dagesh)
  applies to names and terms the project has no fixed form for.
- Fixed forms: רבינו בחיי = "Rabbeinu Bachya"; בעל העיקרים = "the author of
  the Ikkarim"; סמ"ג = "the Semag (Sefer Mitzvot Gadol)"; ר"א = "Rabbi Eleazar"
  only where the paragraph names him, else "R. A." with a note.
- שם meaning "there" or "the name": choose by context; record the other in
  `apparatus`. העתיק: "rendered" (translated), with "copied" in a token note
  when the sense could be copying. פנים in שלשה פנים: "aspects", with the
  literal "faces" in a token note.
- A Biblical quotation whose wording differs from the verse follows the
  author's wording; the difference goes in `apparatus`.
- `literal` on a long paragraph: the glosses in Hebrew order, hyphens joining
  the parts of one word. No reasoning in it; reasons go in `note`.
- `context` may carry a short note on a story the author cites (the Pardes
  story), without a reference unless the author gives one.

## Rulings from the pilot reviews, by the editor

- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- ולא אמרינן כלא חד (1:7) and באגרונו (1:12) are for the owner: the main
  English stands as translated; the reviewer's alternative is in `readings`
  / `apparatus`.

## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
- כלא חד / כולא חד: the same English at every occurrence (1:5-1:7); the
  owner's ruling on the sense is pending; until then, the alternative is in
  `apparatus`.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
- מ"ו = "our teacher" (מורנו) in the formula יודיענו מורנו; ב"ו = "blessed be
  He" only before הוא elsewhere ב"ו = בשר ודם ("flesh and blood").
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- פ' + numeral after a book title: "chapter" when the numeral exceeds the
  parashot of that book (Midrash Rabbah 143 is a section/chapter number as
  the author gives it); otherwise "parashah". Note the choice.
- אם ... או after a full stop: translate as the author punctuates it only if
  the English has a main clause; otherwise join it to the previous sentence
  with a colon or render it as a direct alternative question.
- ריש מלין אמר echoing Dan 7:1: by the verse's sense ("he stated the sum of the
  matter"); the verse in `apparatus`.
- A variant spelling of a word the text does not gloss (ברותא for ברייתא):
  transliterate it ("berota"), with a note.
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
- הרב before a book title (הרב לחם משנה): "the Rav Lechem Mishneh", never drop
  הרב. (Applies with the "author of" ruling: use "the Rabbi, the author of
  [title]" only where the Hebrew says בעל / מחבר.)
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- דרך ארץ: "worldly occupation", as in the project's Avot.
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- ב"ע: "in Aramaic" (בארמית) where it is set against בערבי; the other readings in
  `readings`.
- Names of works: "Guide for the Perplexed" is rendered as the Moreh
  Nevukhim; ממלכי as "of the kings of".
- Closing formulas (וחתם דבריו, עכ"ל, ע"כ) stand outside the closing quotation
  mark; opening formulas (וז"ל, כ' וז"ל) stand outside the opening one.
- A conditional whose main clause follows a Hebrew full stop (ואם ראיתי. הם
  מגלים): join with a comma and say so in a token note.
- Nested quotations: double marks for the outer quotation, single marks for a
  quotation inside it, double again inside that.
- Printed list markers (א), ב)): keep as the author's own, "1)", "2)"; a token
  note.
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

## Translator's checklist (the pilot reviews found these faults most often)

Before running the check, read your English once against the Hebrew:
1. Every Hebrew word has its English; none is added (no "who received it",
   no parenthetical explanation). A full stop in the Hebrew is a full stop in
   the English, and every sentence has a main clause.
2. Recognise Bible verses, even in pieces (Isa 62:6, Job 37:18, Gen 43:14,
   Lev 26:44); translate them from the Hebrew, check each verb's form, and
   list the verse in `apparatus`.
3. Use the fixed renderings of this file for every term, name and abbreviation.
4. A passive participle is passive; a participle is not a perfect.
5. Capitalise only what the file capitalises. Begin each sentence with a
   capital.

## Owner's answers and defaults

Answered by the owner (2026-11-10):

- Q4: Judeo-Arabic and Aramaic are Sonnet's, not Haiku's. `models.json`
  lists the units Sonnet translates: the four Judeo-Arabic units (5:3, 7:24,
  9:10, 2:13) and the units with at least six Aramaic marker words (15 of 179;
  a heuristic, and the Sonnet review covers every unit). The `lang` tags
  stay "heb".
- Q7: the project's transliteration scheme (above).

Not yet answered; the draft's proposal stands as the default and can be
changed before the review stage (a change to 1, 3 or 5 is a find-and-patch):

- Q1 title: "The Wars of the LORD" (Num 21:14).
- Q2 Wikisource additions: translated as base text, flagged in `apparatus`
  until the 1931 scan is collated.
- Q3 השי"ת: "the Name, may He be blessed".
- Q5 קבלה האמתית: "the true tradition", with a note.
- Q6 author's name: patronymic kept as written, "Sali'".


## Milchamot ha-Shem 4:2

Hebrew (Wikisource): מבואר מזה, כי הקב"ה הנקרא עתיק נגלה לר' שמעון בבית מדרשו ובדמות זקן מופלג, ולכך קרא לו ר' שמעון סבא דסבין רוצה לומר זקן הזקנים כדלעיל בשם דרך אמת ומקדש מלך, והקב"ה זה הנקרא עתיק הוא שהרשה לר' שמעון לגלות פנים בתורה, שהקב"ה הנקרא אמא היא שאמרה להקב"ה הנקרא אבא נעשה אדם בצלמינו בדמותינו וכו'. הרי נתבאר לך שהוא קורא לעתיק קב"ה וכן לאבא, וקורא לאמא אלהים, וגם קורא לאבא אלהים וגם קב"ה, ובכל מעשה בראשית אבא אומר לאמא יהי כך וכך והיא עושה, ובמאמר נעשה אדם, אימא אמרה לאבא נעשה אדם ולא הסכים אבא בבריאת אדם. ואמרה לו אמא, מאי נפקא לך מניה, אם יחטא לי יחטא ולא לך שנ' ובן כסיל תוגת אמו, וכשחטא אדם בעץ הדעת גירש אותו ואותה שנאמר ובפשעיכם שולחה אמכם.

| id | word | punctuation after |
|---|---|---|
| 4:2.1 | מבואר |  |
| 4:2.2 | מזה | , |
| 4:2.3 | כי |  |
| 4:2.4 | הקב"ה |  |
| 4:2.5 | הנקרא |  |
| 4:2.6 | עתיק |  |
| 4:2.7 | נגלה |  |
| 4:2.8 | לר | ' |
| 4:2.9 | שמעון |  |
| 4:2.10 | בבית |  |
| 4:2.11 | מדרשו |  |
| 4:2.12 | ובדמות |  |
| 4:2.13 | זקן |  |
| 4:2.14 | מופלג | , |
| 4:2.15 | ולכך |  |
| 4:2.16 | קרא |  |
| 4:2.17 | לו |  |
| 4:2.18 | ר | ' |
| 4:2.19 | שמעון |  |
| 4:2.20 | סבא |  |
| 4:2.21 | דסבין |  |
| 4:2.22 | רוצה |  |
| 4:2.23 | לומר |  |
| 4:2.24 | זקן |  |
| 4:2.25 | הזקנים |  |
| 4:2.26 | כדלעיל |  |
| 4:2.27 | בשם |  |
| 4:2.28 | דרך |  |
| 4:2.29 | אמת |  |
| 4:2.30 | ומקדש |  |
| 4:2.31 | מלך | , |
| 4:2.32 | והקב"ה |  |
| 4:2.33 | זה |  |
| 4:2.34 | הנקרא |  |
| 4:2.35 | עתיק |  |
| 4:2.36 | הוא |  |
| 4:2.37 | שהרשה |  |
| 4:2.38 | לר | ' |
| 4:2.39 | שמעון |  |
| 4:2.40 | לגלות |  |
| 4:2.41 | פנים |  |
| 4:2.42 | בתורה | , |
| 4:2.43 | שהקב"ה |  |
| 4:2.44 | הנקרא |  |
| 4:2.45 | אמא |  |
| 4:2.46 | היא |  |
| 4:2.47 | שאמרה |  |
| 4:2.48 | להקב"ה |  |
| 4:2.49 | הנקרא |  |
| 4:2.50 | אבא |  |
| 4:2.51 | נעשה |  |
| 4:2.52 | אדם |  |
| 4:2.53 | בצלמינו |  |
| 4:2.54 | בדמותינו |  |
| 4:2.55 | וכו | '. |
| 4:2.56 | הרי |  |
| 4:2.57 | נתבאר |  |
| 4:2.58 | לך |  |
| 4:2.59 | שהוא |  |
| 4:2.60 | קורא |  |
| 4:2.61 | לעתיק |  |
| 4:2.62 | קב"ה |  |
| 4:2.63 | וכן |  |
| 4:2.64 | לאבא | , |
| 4:2.65 | וקורא |  |
| 4:2.66 | לאמא |  |
| 4:2.67 | אלהים | , |
| 4:2.68 | וגם |  |
| 4:2.69 | קורא |  |
| 4:2.70 | לאבא |  |
| 4:2.71 | אלהים |  |
| 4:2.72 | וגם |  |
| 4:2.73 | קב"ה | , |
| 4:2.74 | ובכל |  |
| 4:2.75 | מעשה |  |
| 4:2.76 | בראשית |  |
| 4:2.77 | אבא |  |
| 4:2.78 | אומר |  |
| 4:2.79 | לאמא |  |
| 4:2.80 | יהי |  |
| 4:2.81 | כך |  |
| 4:2.82 | וכך |  |
| 4:2.83 | והיא |  |
| 4:2.84 | עושה | , |
| 4:2.85 | ובמאמר |  |
| 4:2.86 | נעשה |  |
| 4:2.87 | אדם | , |
| 4:2.88 | אימא |  |
| 4:2.89 | אמרה |  |
| 4:2.90 | לאבא |  |
| 4:2.91 | נעשה |  |
| 4:2.92 | אדם |  |
| 4:2.93 | ולא |  |
| 4:2.94 | הסכים |  |
| 4:2.95 | אבא |  |
| 4:2.96 | בבריאת |  |
| 4:2.97 | אדם | . |
| 4:2.98 | ואמרה |  |
| 4:2.99 | לו |  |
| 4:2.100 | אמא | , |
| 4:2.101 | מאי |  |
| 4:2.102 | נפקא |  |
| 4:2.103 | לך |  |
| 4:2.104 | מניה | , |
| 4:2.105 | אם |  |
| 4:2.106 | יחטא |  |
| 4:2.107 | לי |  |
| 4:2.108 | יחטא |  |
| 4:2.109 | ולא |  |
| 4:2.110 | לך |  |
| 4:2.111 | שנ | ' |
| 4:2.112 | ובן |  |
| 4:2.113 | כסיל |  |
| 4:2.114 | תוגת |  |
| 4:2.115 | אמו | , |
| 4:2.116 | וכשחטא |  |
| 4:2.117 | אדם |  |
| 4:2.118 | בעץ |  |
| 4:2.119 | הדעת |  |
| 4:2.120 | גירש |  |
| 4:2.121 | אותו |  |
| 4:2.122 | ואותה |  |
| 4:2.123 | שנאמר |  |
| 4:2.124 | ובפשעיכם |  |
| 4:2.125 | שולחה |  |
| 4:2.126 | אמכם | . |


## Milchamot ha-Shem 4:3

Hebrew (Wikisource): ובעץ חיים להרח"ו שער הנקודים פ"ו כ' שכאשר נאצלו שבע תחתונות עדין לא היה אדם הראשון התחתון נברא בעולם ויצאו זעיר ונוקביה אחור באחור מפני פחד החיצונים שלא יינקו, כי אם היו עומדים פנים בפנים היה מקום לקליפות להתאחז במקום אחיזתן וכאשר נברא אדם הראשון ועשה מצות מעשיות החזירן פנים בפנים על ידי מצות ומעשים טובים וכו' עיין שם שכל דבריו הם היפך דברי הזהר פה, שאדרבה אדם הראשון קלקל בחטאו ונתגרשה אימא בעבורו, ומזה תבין שכל אחד מהקובלים בודה לו דברים מלבו כחפצו.

| id | word | punctuation after |
|---|---|---|
| 4:3.1 | ובעץ |  |
| 4:3.2 | חיים |  |
| 4:3.3 | להרח"ו |  |
| 4:3.4 | שער |  |
| 4:3.5 | הנקודים |  |
| 4:3.6 | פ"ו |  |
| 4:3.7 | כ | ' |
| 4:3.8 | שכאשר |  |
| 4:3.9 | נאצלו |  |
| 4:3.10 | שבע |  |
| 4:3.11 | תחתונות |  |
| 4:3.12 | עדין |  |
| 4:3.13 | לא |  |
| 4:3.14 | היה |  |
| 4:3.15 | אדם |  |
| 4:3.16 | הראשון |  |
| 4:3.17 | התחתון |  |
| 4:3.18 | נברא |  |
| 4:3.19 | בעולם |  |
| 4:3.20 | ויצאו |  |
| 4:3.21 | זעיר |  |
| 4:3.22 | ונוקביה |  |
| 4:3.23 | אחור |  |
| 4:3.24 | באחור |  |
| 4:3.25 | מפני |  |
| 4:3.26 | פחד |  |
| 4:3.27 | החיצונים |  |
| 4:3.28 | שלא |  |
| 4:3.29 | יינקו | , |
| 4:3.30 | כי |  |
| 4:3.31 | אם |  |
| 4:3.32 | היו |  |
| 4:3.33 | עומדים |  |
| 4:3.34 | פנים |  |
| 4:3.35 | בפנים |  |
| 4:3.36 | היה |  |
| 4:3.37 | מקום |  |
| 4:3.38 | לקליפות |  |
| 4:3.39 | להתאחז |  |
| 4:3.40 | במקום |  |
| 4:3.41 | אחיזתן |  |
| 4:3.42 | וכאשר |  |
| 4:3.43 | נברא |  |
| 4:3.44 | אדם |  |
| 4:3.45 | הראשון |  |
| 4:3.46 | ועשה |  |
| 4:3.47 | מצות |  |
| 4:3.48 | מעשיות |  |
| 4:3.49 | החזירן |  |
| 4:3.50 | פנים |  |
| 4:3.51 | בפנים |  |
| 4:3.52 | על |  |
| 4:3.53 | ידי |  |
| 4:3.54 | מצות |  |
| 4:3.55 | ומעשים |  |
| 4:3.56 | טובים |  |
| 4:3.57 | וכו | ' |
| 4:3.58 | עיין |  |
| 4:3.59 | שם |  |
| 4:3.60 | שכל |  |
| 4:3.61 | דבריו |  |
| 4:3.62 | הם |  |
| 4:3.63 | היפך |  |
| 4:3.64 | דברי |  |
| 4:3.65 | הזהר |  |
| 4:3.66 | פה | , |
| 4:3.67 | שאדרבה |  |
| 4:3.68 | אדם |  |
| 4:3.69 | הראשון |  |
| 4:3.70 | קלקל |  |
| 4:3.71 | בחטאו |  |
| 4:3.72 | ונתגרשה |  |
| 4:3.73 | אימא |  |
| 4:3.74 | בעבורו | , |
| 4:3.75 | ומזה |  |
| 4:3.76 | תבין |  |
| 4:3.77 | שכל |  |
| 4:3.78 | אחד |  |
| 4:3.79 | מהקובלים |  |
| 4:3.80 | בודה |  |
| 4:3.81 | לו |  |
| 4:3.82 | דברים |  |
| 4:3.83 | מלבו |  |
| 4:3.84 | כחפצו | . |


## Milchamot ha-Shem 4:4

Hebrew (Wikisource): כא) עוד שם בזהר ד' כ"ב ע"ב פתח עוד ר' שמעון ואמר ראו עתה כי אני אני הוא ואין אלהים עמדי, אמר, חבריא, שמעון מלין עתיקין דבעינא לגלאה בתר דאתיהיב רשו עלאה למימר, מאן ניהו דאמר ראו עתה כי אני אני הוא, אלא דא הוא עלת על כל עלאין ההוא דאתקרי עלת העלות, עלת מאלין עלות דלא יעביד חד מאלין עלות שום עובדא עד דנטיל רשות מההוא דעילוה. כמא דאוקימנא לעילא בנעשה אדם, נעשה ודאי על תרין אתמר, דאמר דא לההוא דלעילא מניה נעשה, ולא עביד מדעם עד דנטיל רשותא ואמירה מההוא דלעילא מניה וההוא דלעילא מניה לא עביד מדעם עד דנטיל עצה מחבריה, אבל ההוא דאתקרי עלת כל עלות דלית לעילא מניה, ולא לתתא שוה ליה, כמא דאתמר ואל מי תדמיוני ואשוה יאמר קדוש, אמר ראו עתה כי אני אני הוא ואין אלהים עמדי דנטיל עצה מניה כגוונא דההוא דאמר ויאמר אלהים נעשה אדם, ועיין עץ חיים שער עתיק פ' שני, ופי' הרב מקדש מלך בשם האר"י בספר הלקוטים, עלת על כל עלאין אדם קדמון נקרא עלת על כל עלות, ההוא דאתקרי עלת העלות, עלת מאלין עלות, פי' כשאנו אומרים עלת העלות שייך בכל פרצוף שנקרא כך לפי שהוא עלה לעלות שתחתיו, אבל כשאנו אומרים עלת כל העלות הוא אדם קדמון, שהוא עלה ראשונה כל הפרצופים כולם ע"ש. הרי נתבאר לך מדברי הפילוסוף מחבר הזהר ומפרשיו מקדש מלך בשם האר"י ודרך אמת ושביבי נגה שהאלוה שנתן רשות לרשב"י לדרוש בדברים שאסור להעלותן על לב לחשוב רבוי באלהות, הוא הקב"ה הנקרא עתיק, ולא זה הוא שאמר לישראל ראו עתה כי אני אני הוא וכו'. כי הקב"ה שאמר ראו עתה כי אני אני הוא, הוא אדם קדמון שהוא עלה לראשונה לכל הפרצופים, ואין לו ממי ליטול רשות. והקב"ה שאמר יהי אור, ויהי רקיע, ויקוו המים וכו' הוא אבא, וזה שאמר נעשה אדם בצלמינו היא אמא שאמרה לאבא נעשה אדם, ולא הסכים הקב"ה שהוא אבא בבריאתו, ואמרה לו אימא מאי איכפת לך, אם יחטא לי יחטא ולא לך שנ' ובן כסיל תוגת אמו ולא תוגת אביו, ואז נברא שלא לרצון אבא כמו שפי' שביבי נגה, שאלו הסכים אבא בבריאתו לא גירשה עמו כשחטא, אלא ודאי לא הסכים אבא לזה כמו שאמר בזהר ולא בעא לאשתתפא וכו' ולכך גירשו מגן עדן וגירשה עמו.

| id | word | punctuation after |
|---|---|---|
| 4:4.1 | כא | ) |
| 4:4.2 | עוד |  |
| 4:4.3 | שם |  |
| 4:4.4 | בזהר |  |
| 4:4.5 | ד | ' |
| 4:4.6 | כ"ב |  |
| 4:4.7 | ע"ב |  |
| 4:4.8 | פתח |  |
| 4:4.9 | עוד |  |
| 4:4.10 | ר | ' |
| 4:4.11 | שמעון |  |
| 4:4.12 | ואמר |  |
| 4:4.13 | ראו |  |
| 4:4.14 | עתה |  |
| 4:4.15 | כי |  |
| 4:4.16 | אני |  |
| 4:4.17 | אני |  |
| 4:4.18 | הוא |  |
| 4:4.19 | ואין |  |
| 4:4.20 | אלהים |  |
| 4:4.21 | עמדי | , |
| 4:4.22 | אמר | , |
| 4:4.23 | חבריא | , |
| 4:4.24 | שמעון |  |
| 4:4.25 | מלין |  |
| 4:4.26 | עתיקין |  |
| 4:4.27 | דבעינא |  |
| 4:4.28 | לגלאה |  |
| 4:4.29 | בתר |  |
| 4:4.30 | דאתיהיב |  |
| 4:4.31 | רשו |  |
| 4:4.32 | עלאה |  |
| 4:4.33 | למימר | , |
| 4:4.34 | מאן |  |
| 4:4.35 | ניהו |  |
| 4:4.36 | דאמר |  |
| 4:4.37 | ראו |  |
| 4:4.38 | עתה |  |
| 4:4.39 | כי |  |
| 4:4.40 | אני |  |
| 4:4.41 | אני |  |
| 4:4.42 | הוא | , |
| 4:4.43 | אלא |  |
| 4:4.44 | דא |  |
| 4:4.45 | הוא |  |
| 4:4.46 | עלת |  |
| 4:4.47 | על |  |
| 4:4.48 | כל |  |
| 4:4.49 | עלאין |  |
| 4:4.50 | ההוא |  |
| 4:4.51 | דאתקרי |  |
| 4:4.52 | עלת |  |
| 4:4.53 | העלות | , |
| 4:4.54 | עלת |  |
| 4:4.55 | מאלין |  |
| 4:4.56 | עלות |  |
| 4:4.57 | דלא |  |
| 4:4.58 | יעביד |  |
| 4:4.59 | חד |  |
| 4:4.60 | מאלין |  |
| 4:4.61 | עלות |  |
| 4:4.62 | שום |  |
| 4:4.63 | עובדא |  |
| 4:4.64 | עד |  |
| 4:4.65 | דנטיל |  |
| 4:4.66 | רשות |  |
| 4:4.67 | מההוא |  |
| 4:4.68 | דעילוה | . |
| 4:4.69 | כמא |  |
| 4:4.70 | דאוקימנא |  |
| 4:4.71 | לעילא |  |
| 4:4.72 | בנעשה |  |
| 4:4.73 | אדם | , |
| 4:4.74 | נעשה |  |
| 4:4.75 | ודאי |  |
| 4:4.76 | על |  |
| 4:4.77 | תרין |  |
| 4:4.78 | אתמר | , |
| 4:4.79 | דאמר |  |
| 4:4.80 | דא |  |
| 4:4.81 | לההוא |  |
| 4:4.82 | דלעילא |  |
| 4:4.83 | מניה |  |
| 4:4.84 | נעשה | , |
| 4:4.85 | ולא |  |
| 4:4.86 | עביד |  |
| 4:4.87 | מדעם |  |
| 4:4.88 | עד |  |
| 4:4.89 | דנטיל |  |
| 4:4.90 | רשותא |  |
| 4:4.91 | ואמירה |  |
| 4:4.92 | מההוא |  |
| 4:4.93 | דלעילא |  |
| 4:4.94 | מניה |  |
| 4:4.95 | וההוא |  |
| 4:4.96 | דלעילא |  |
| 4:4.97 | מניה |  |
| 4:4.98 | לא |  |
| 4:4.99 | עביד |  |
| 4:4.100 | מדעם |  |
| 4:4.101 | עד |  |
| 4:4.102 | דנטיל |  |
| 4:4.103 | עצה |  |
| 4:4.104 | מחבריה | , |
| 4:4.105 | אבל |  |
| 4:4.106 | ההוא |  |
| 4:4.107 | דאתקרי |  |
| 4:4.108 | עלת |  |
| 4:4.109 | כל |  |
| 4:4.110 | עלות |  |
| 4:4.111 | דלית |  |
| 4:4.112 | לעילא |  |
| 4:4.113 | מניה | , |
| 4:4.114 | ולא |  |
| 4:4.115 | לתתא |  |
| 4:4.116 | שוה |  |
| 4:4.117 | ליה | , |
| 4:4.118 | כמא |  |
| 4:4.119 | דאתמר |  |
| 4:4.120 | ואל |  |
| 4:4.121 | מי |  |
| 4:4.122 | תדמיוני |  |
| 4:4.123 | ואשוה |  |
| 4:4.124 | יאמר |  |
| 4:4.125 | קדוש | , |
| 4:4.126 | אמר |  |
| 4:4.127 | ראו |  |
| 4:4.128 | עתה |  |
| 4:4.129 | כי |  |
| 4:4.130 | אני |  |
| 4:4.131 | אני |  |
| 4:4.132 | הוא |  |
| 4:4.133 | ואין |  |
| 4:4.134 | אלהים |  |
| 4:4.135 | עמדי |  |
| 4:4.136 | דנטיל |  |
| 4:4.137 | עצה |  |
| 4:4.138 | מניה |  |
| 4:4.139 | כגוונא |  |
| 4:4.140 | דההוא |  |
| 4:4.141 | דאמר |  |
| 4:4.142 | ויאמר |  |
| 4:4.143 | אלהים |  |
| 4:4.144 | נעשה |  |
| 4:4.145 | אדם | , |
| 4:4.146 | ועיין |  |
| 4:4.147 | עץ |  |
| 4:4.148 | חיים |  |
| 4:4.149 | שער |  |
| 4:4.150 | עתיק |  |
| 4:4.151 | פ | ' |
| 4:4.152 | שני | , |
| 4:4.153 | ופי | ' |
| 4:4.154 | הרב |  |
| 4:4.155 | מקדש |  |
| 4:4.156 | מלך |  |
| 4:4.157 | בשם |  |
| 4:4.158 | האר"י |  |
| 4:4.159 | בספר |  |
| 4:4.160 | הלקוטים | , |
| 4:4.161 | עלת |  |
| 4:4.162 | על |  |
| 4:4.163 | כל |  |
| 4:4.164 | עלאין |  |
| 4:4.165 | אדם |  |
| 4:4.166 | קדמון |  |
| 4:4.167 | נקרא |  |
| 4:4.168 | עלת |  |
| 4:4.169 | על |  |
| 4:4.170 | כל |  |
| 4:4.171 | עלות | , |
| 4:4.172 | ההוא |  |
| 4:4.173 | דאתקרי |  |
| 4:4.174 | עלת |  |
| 4:4.175 | העלות | , |
| 4:4.176 | עלת |  |
| 4:4.177 | מאלין |  |
| 4:4.178 | עלות | , |
| 4:4.179 | פי | ' |
| 4:4.180 | כשאנו |  |
| 4:4.181 | אומרים |  |
| 4:4.182 | עלת |  |
| 4:4.183 | העלות |  |
| 4:4.184 | שייך |  |
| 4:4.185 | בכל |  |
| 4:4.186 | פרצוף |  |
| 4:4.187 | שנקרא |  |
| 4:4.188 | כך |  |
| 4:4.189 | לפי |  |
| 4:4.190 | שהוא |  |
| 4:4.191 | עלה |  |
| 4:4.192 | לעלות |  |
| 4:4.193 | שתחתיו | , |
| 4:4.194 | אבל |  |
| 4:4.195 | כשאנו |  |
| 4:4.196 | אומרים |  |
| 4:4.197 | עלת |  |
| 4:4.198 | כל |  |
| 4:4.199 | העלות |  |
| 4:4.200 | הוא |  |
| 4:4.201 | אדם |  |
| 4:4.202 | קדמון | , |
| 4:4.203 | שהוא |  |
| 4:4.204 | עלה |  |
| 4:4.205 | ראשונה |  |
| 4:4.206 | כל |  |
| 4:4.207 | הפרצופים |  |
| 4:4.208 | כולם |  |
| 4:4.209 | ע"ש | . |
| 4:4.210 | הרי |  |
| 4:4.211 | נתבאר |  |
| 4:4.212 | לך |  |
| 4:4.213 | מדברי |  |
| 4:4.214 | הפילוסוף |  |
| 4:4.215 | מחבר |  |
| 4:4.216 | הזהר |  |
| 4:4.217 | ומפרשיו |  |
| 4:4.218 | מקדש |  |
| 4:4.219 | מלך |  |
| 4:4.220 | בשם |  |
| 4:4.221 | האר"י |  |
| 4:4.222 | ודרך |  |
| 4:4.223 | אמת |  |
| 4:4.224 | ושביבי |  |
| 4:4.225 | נגה |  |
| 4:4.226 | שהאלוה |  |
| 4:4.227 | שנתן |  |
| 4:4.228 | רשות |  |
| 4:4.229 | לרשב"י |  |
| 4:4.230 | לדרוש |  |
| 4:4.231 | בדברים |  |
| 4:4.232 | שאסור |  |
| 4:4.233 | להעלותן |  |
| 4:4.234 | על |  |
| 4:4.235 | לב |  |
| 4:4.236 | לחשוב |  |
| 4:4.237 | רבוי |  |
| 4:4.238 | באלהות | , |
| 4:4.239 | הוא |  |
| 4:4.240 | הקב"ה |  |
| 4:4.241 | הנקרא |  |
| 4:4.242 | עתיק | , |
| 4:4.243 | ולא |  |
| 4:4.244 | זה |  |
| 4:4.245 | הוא |  |
| 4:4.246 | שאמר |  |
| 4:4.247 | לישראל |  |
| 4:4.248 | ראו |  |
| 4:4.249 | עתה |  |
| 4:4.250 | כי |  |
| 4:4.251 | אני |  |
| 4:4.252 | אני |  |
| 4:4.253 | הוא |  |
| 4:4.254 | וכו | '. |
| 4:4.255 | כי |  |
| 4:4.256 | הקב"ה |  |
| 4:4.257 | שאמר |  |
| 4:4.258 | ראו |  |
| 4:4.259 | עתה |  |
| 4:4.260 | כי |  |
| 4:4.261 | אני |  |
| 4:4.262 | אני |  |
| 4:4.263 | הוא | , |
| 4:4.264 | הוא |  |
| 4:4.265 | אדם |  |
| 4:4.266 | קדמון |  |
| 4:4.267 | שהוא |  |
| 4:4.268 | עלה |  |
| 4:4.269 | לראשונה |  |
| 4:4.270 | לכל |  |
| 4:4.271 | הפרצופים | , |
| 4:4.272 | ואין |  |
| 4:4.273 | לו |  |
| 4:4.274 | ממי |  |
| 4:4.275 | ליטול |  |
| 4:4.276 | רשות | . |
| 4:4.277 | והקב"ה |  |
| 4:4.278 | שאמר |  |
| 4:4.279 | יהי |  |
| 4:4.280 | אור | , |
| 4:4.281 | ויהי |  |
| 4:4.282 | רקיע | , |
| 4:4.283 | ויקוו |  |
| 4:4.284 | המים |  |
| 4:4.285 | וכו | ' |
| 4:4.286 | הוא |  |
| 4:4.287 | אבא | , |
| 4:4.288 | וזה |  |
| 4:4.289 | שאמר |  |
| 4:4.290 | נעשה |  |
| 4:4.291 | אדם |  |
| 4:4.292 | בצלמינו |  |
| 4:4.293 | היא |  |
| 4:4.294 | אמא |  |
| 4:4.295 | שאמרה |  |
| 4:4.296 | לאבא |  |
| 4:4.297 | נעשה |  |
| 4:4.298 | אדם | , |
| 4:4.299 | ולא |  |
| 4:4.300 | הסכים |  |
| 4:4.301 | הקב"ה |  |
| 4:4.302 | שהוא |  |
| 4:4.303 | אבא |  |
| 4:4.304 | בבריאתו | , |
| 4:4.305 | ואמרה |  |
| 4:4.306 | לו |  |
| 4:4.307 | אימא |  |
| 4:4.308 | מאי |  |
| 4:4.309 | איכפת |  |
| 4:4.310 | לך | , |
| 4:4.311 | אם |  |
| 4:4.312 | יחטא |  |
| 4:4.313 | לי |  |
| 4:4.314 | יחטא |  |
| 4:4.315 | ולא |  |
| 4:4.316 | לך |  |
| 4:4.317 | שנ | ' |
| 4:4.318 | ובן |  |
| 4:4.319 | כסיל |  |
| 4:4.320 | תוגת |  |
| 4:4.321 | אמו |  |
| 4:4.322 | ולא |  |
| 4:4.323 | תוגת |  |
| 4:4.324 | אביו | , |
| 4:4.325 | ואז |  |
| 4:4.326 | נברא |  |
| 4:4.327 | שלא |  |
| 4:4.328 | לרצון |  |
| 4:4.329 | אבא |  |
| 4:4.330 | כמו |  |
| 4:4.331 | שפי | ' |
| 4:4.332 | שביבי |  |
| 4:4.333 | נגה | , |
| 4:4.334 | שאלו |  |
| 4:4.335 | הסכים |  |
| 4:4.336 | אבא |  |
| 4:4.337 | בבריאתו |  |
| 4:4.338 | לא |  |
| 4:4.339 | גירשה |  |
| 4:4.340 | עמו |  |
| 4:4.341 | כשחטא | , |
| 4:4.342 | אלא |  |
| 4:4.343 | ודאי |  |
| 4:4.344 | לא |  |
| 4:4.345 | הסכים |  |
| 4:4.346 | אבא |  |
| 4:4.347 | לזה |  |
| 4:4.348 | כמו |  |
| 4:4.349 | שאמר |  |
| 4:4.350 | בזהר |  |
| 4:4.351 | ולא |  |
| 4:4.352 | בעא |  |
| 4:4.353 | לאשתתפא |  |
| 4:4.354 | וכו | ' |
| 4:4.355 | ולכך |  |
| 4:4.356 | גירשו |  |
| 4:4.357 | מגן |  |
| 4:4.358 | עדן |  |
| 4:4.359 | וגירשה |  |
| 4:4.360 | עמו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

