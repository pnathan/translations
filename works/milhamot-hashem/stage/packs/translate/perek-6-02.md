# Translation pack: milhamot-hashem perek-6-02 (Milchamot ha-Shem 6:2–6:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-02`.

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
- הוי"ה / הוי'ה / הויה (the Name as a name: Kabbalistic usage, and the
  author's own in 3:17) → "YHVH". Never "being". Examples: "ד' מילוי הויה ע"ב
  ס"ג מ"ה ב"ן" (2:33) → "the four fillings of YHVH, 72, 63, 45, 52" (in the
  note: these are the gematria names AB, SaG, MaH, BaN). "שהוא קורא לכל אחד
  מהם בשם הוי'ה ואדנות ואלהים" (8:24) → "who calls each of them by the names
  YHVH, Adonai and Elohim". אדנות as a name → "Adonai", with a note that the
  word is literally "lordship". עשר הויות (2:33) → "ten havayot", with a note.
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
- **אצילות** and its root: fixed and translated, so that the verbs the
  author strings together stay linked. אצילות → "emanation"; עולם האצילות /
  דאצילות → "the World of Emanation"; נאצל → "emanated"; הנאצלים →
  "the emanated beings"; המאציל → "the Emanator". The four worlds:
  "Emanation, Creation, Formation, Making" (אבי"ע). "נאצלו ונבראו ונוצרו
  ונעשו" → "were emanated, created, formed and made". Keep נאצל and נברא
  distinct everywhere: the author argues that Ze'ir Anpin is "נברא ולא בורא"
  → "created, not creator".
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
- **קבלה / המקובלים / בעלי הקבלה**: context.
  - The lore: "Kabbalah" (חכמת הקבלה → "the wisdom of Kabbalah"; הקבלה החדשה
    → "the new Kabbalah", always "new", never "later" or "Lurianic").
  - The people: המקובלים → "the Kabbalists"; המקובלים החדשים → "the new
    Kabbalists"; "רבותינו המקובלים" / "רבותינו בעלי הקבלה החדשה" → "our
    rabbis the Kabbalists" / "our rabbis, the masters of the new Kabbalah".
    Keep the courtesy as written.
  - The author's claim on the word for the rabbinic chain: "בעלי הקבלה
    האמתית אשר קבלו ממרע"ה" (2:34) → "the masters of the true tradition, who
    received from Moses our teacher". Add a note each time: the Hebrew is the
    word for Kabbalah, and he claims it for the Mishnah and Talmud. "אינה
    קבלה ממשה" (3:18) → "it is not a tradition from Moses".
  - The ordinary participle: "וכמקובל אצל כל איש ישראל" (1:11) → "as is
    accepted by every Jew"; "המקובל בידינו" (2:3) → "what we have received".
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
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- יש מאין / יש מיש: translate, "something from nothing" / "something from
  something"; no transliteration.
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
- דרך ארץ: "worldly occupation", as in the project's Avot.
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 6:2

Hebrew (Wikisource): סא) וכ"כ הסמ"ג בעשיין א' מצות עשה להאמין ולשמוע היא הקבלה שהוא ית' בשמים ובארץ ובארבע רוחות העולם שנ' שמע ישראל ה' אלהינו ה' אחד, שמע כמו תקבל וכן אתה תשמע השמים וגו' וזה האחד אינו כשאר האחדים, המלך הוא אחד בארצו, אבל אינו אחד מכל צדדים, כי יש בארצו אנשים שהם כמהו ובשאר ארצות יש מלכים כמהו, ואם מלאך מן השמים ירד לארץ. הנה המלאך ההוא אחד בארץ. אבל אינו אחד מכל צדדים כי יש בשמים מלאכים כמותו. אבל ה' אלהינו אינו כך. הוא אחד מכל הצדדים. ורב סעדיה כ' להשיב לכופרים שאומרים שהם שנים או שלשה. אם לא יוכל האחד לעשות מעשה אלא בעזר חבירו האי שניהם חלשים. ואם אחד יכול להכריח את חבירו, הרי שניהם מוכרחים. ואם שניהם רשאים לעשות על אחד מה שירצה, ובקש אחד להמית אדם אחד והשני רוצה להחיותו. ראוי שיהיה האיש ההוא מת חי מת חי לאלתר כהרף עין. ואם יוכל האחד להסתיר דבר מחבירו הרי שניהם נלאים. ועוד כתב לאלה שפוקדים על נעשה אדם בצלמנו (היא דעת מחבר הזהר והמקובלים כדלעיל סי' ל"ח) כי דרך שפת לשון הקדש שהגדולים מכנים עצמם בלשון רבים אע"פ שהם יחידים, כמו אולי אוכל נכה בו. ובדניאל נאמר קדם מלכא ובמנוח נעצרה נא אותך ע"כ ועוד האריך רבי סעדיה גאון מאד, ואין צורך להאריך כל כך, כי כל ישראל תקועים באמונה חזקה, כי יוצר הכל אחד מיוחד כמו שנ' ה' אחד. ונאמר ראו עתה כי אני הוא ואין אלהים עמדי. ובמקרא יש הרבה מקראות כאלה עכ"ל הרב הסמ"ג.

| id | word | punctuation after |
|---|---|---|
| 6:2.1 | סא | ) |
| 6:2.2 | וכ"כ |  |
| 6:2.3 | הסמ"ג |  |
| 6:2.4 | בעשיין |  |
| 6:2.5 | א | ' |
| 6:2.6 | מצות |  |
| 6:2.7 | עשה |  |
| 6:2.8 | להאמין |  |
| 6:2.9 | ולשמוע |  |
| 6:2.10 | היא |  |
| 6:2.11 | הקבלה |  |
| 6:2.12 | שהוא |  |
| 6:2.13 | ית | ' |
| 6:2.14 | בשמים |  |
| 6:2.15 | ובארץ |  |
| 6:2.16 | ובארבע |  |
| 6:2.17 | רוחות |  |
| 6:2.18 | העולם |  |
| 6:2.19 | שנ | ' |
| 6:2.20 | שמע |  |
| 6:2.21 | ישראל |  |
| 6:2.22 | ה | ' |
| 6:2.23 | אלהינו |  |
| 6:2.24 | ה | ' |
| 6:2.25 | אחד | , |
| 6:2.26 | שמע |  |
| 6:2.27 | כמו |  |
| 6:2.28 | תקבל |  |
| 6:2.29 | וכן |  |
| 6:2.30 | אתה |  |
| 6:2.31 | תשמע |  |
| 6:2.32 | השמים |  |
| 6:2.33 | וגו | ' |
| 6:2.34 | וזה |  |
| 6:2.35 | האחד |  |
| 6:2.36 | אינו |  |
| 6:2.37 | כשאר |  |
| 6:2.38 | האחדים | , |
| 6:2.39 | המלך |  |
| 6:2.40 | הוא |  |
| 6:2.41 | אחד |  |
| 6:2.42 | בארצו | , |
| 6:2.43 | אבל |  |
| 6:2.44 | אינו |  |
| 6:2.45 | אחד |  |
| 6:2.46 | מכל |  |
| 6:2.47 | צדדים | , |
| 6:2.48 | כי |  |
| 6:2.49 | יש |  |
| 6:2.50 | בארצו |  |
| 6:2.51 | אנשים |  |
| 6:2.52 | שהם |  |
| 6:2.53 | כמהו |  |
| 6:2.54 | ובשאר |  |
| 6:2.55 | ארצות |  |
| 6:2.56 | יש |  |
| 6:2.57 | מלכים |  |
| 6:2.58 | כמהו | , |
| 6:2.59 | ואם |  |
| 6:2.60 | מלאך |  |
| 6:2.61 | מן |  |
| 6:2.62 | השמים |  |
| 6:2.63 | ירד |  |
| 6:2.64 | לארץ | . |
| 6:2.65 | הנה |  |
| 6:2.66 | המלאך |  |
| 6:2.67 | ההוא |  |
| 6:2.68 | אחד |  |
| 6:2.69 | בארץ | . |
| 6:2.70 | אבל |  |
| 6:2.71 | אינו |  |
| 6:2.72 | אחד |  |
| 6:2.73 | מכל |  |
| 6:2.74 | צדדים |  |
| 6:2.75 | כי |  |
| 6:2.76 | יש |  |
| 6:2.77 | בשמים |  |
| 6:2.78 | מלאכים |  |
| 6:2.79 | כמותו | . |
| 6:2.80 | אבל |  |
| 6:2.81 | ה | ' |
| 6:2.82 | אלהינו |  |
| 6:2.83 | אינו |  |
| 6:2.84 | כך | . |
| 6:2.85 | הוא |  |
| 6:2.86 | אחד |  |
| 6:2.87 | מכל |  |
| 6:2.88 | הצדדים | . |
| 6:2.89 | ורב |  |
| 6:2.90 | סעדיה |  |
| 6:2.91 | כ | ' |
| 6:2.92 | להשיב |  |
| 6:2.93 | לכופרים |  |
| 6:2.94 | שאומרים |  |
| 6:2.95 | שהם |  |
| 6:2.96 | שנים |  |
| 6:2.97 | או |  |
| 6:2.98 | שלשה | . |
| 6:2.99 | אם |  |
| 6:2.100 | לא |  |
| 6:2.101 | יוכל |  |
| 6:2.102 | האחד |  |
| 6:2.103 | לעשות |  |
| 6:2.104 | מעשה |  |
| 6:2.105 | אלא |  |
| 6:2.106 | בעזר |  |
| 6:2.107 | חבירו |  |
| 6:2.108 | האי |  |
| 6:2.109 | שניהם |  |
| 6:2.110 | חלשים | . |
| 6:2.111 | ואם |  |
| 6:2.112 | אחד |  |
| 6:2.113 | יכול |  |
| 6:2.114 | להכריח |  |
| 6:2.115 | את |  |
| 6:2.116 | חבירו | , |
| 6:2.117 | הרי |  |
| 6:2.118 | שניהם |  |
| 6:2.119 | מוכרחים | . |
| 6:2.120 | ואם |  |
| 6:2.121 | שניהם |  |
| 6:2.122 | רשאים |  |
| 6:2.123 | לעשות |  |
| 6:2.124 | על |  |
| 6:2.125 | אחד |  |
| 6:2.126 | מה |  |
| 6:2.127 | שירצה | , |
| 6:2.128 | ובקש |  |
| 6:2.129 | אחד |  |
| 6:2.130 | להמית |  |
| 6:2.131 | אדם |  |
| 6:2.132 | אחד |  |
| 6:2.133 | והשני |  |
| 6:2.134 | רוצה |  |
| 6:2.135 | להחיותו | . |
| 6:2.136 | ראוי |  |
| 6:2.137 | שיהיה |  |
| 6:2.138 | האיש |  |
| 6:2.139 | ההוא |  |
| 6:2.140 | מת |  |
| 6:2.141 | חי |  |
| 6:2.142 | מת |  |
| 6:2.143 | חי |  |
| 6:2.144 | לאלתר |  |
| 6:2.145 | כהרף |  |
| 6:2.146 | עין | . |
| 6:2.147 | ואם |  |
| 6:2.148 | יוכל |  |
| 6:2.149 | האחד |  |
| 6:2.150 | להסתיר |  |
| 6:2.151 | דבר |  |
| 6:2.152 | מחבירו |  |
| 6:2.153 | הרי |  |
| 6:2.154 | שניהם |  |
| 6:2.155 | נלאים | . |
| 6:2.156 | ועוד |  |
| 6:2.157 | כתב |  |
| 6:2.158 | לאלה |  |
| 6:2.159 | שפוקדים |  |
| 6:2.160 | על |  |
| 6:2.161 | נעשה |  |
| 6:2.162 | אדם |  |
| 6:2.163 | בצלמנו |  |
| 6:2.164 | היא |  |
| 6:2.165 | דעת |  |
| 6:2.166 | מחבר |  |
| 6:2.167 | הזהר |  |
| 6:2.168 | והמקובלים |  |
| 6:2.169 | כדלעיל |  |
| 6:2.170 | סי | ' |
| 6:2.171 | ל"ח | ) |
| 6:2.172 | כי |  |
| 6:2.173 | דרך |  |
| 6:2.174 | שפת |  |
| 6:2.175 | לשון |  |
| 6:2.176 | הקדש |  |
| 6:2.177 | שהגדולים |  |
| 6:2.178 | מכנים |  |
| 6:2.179 | עצמם |  |
| 6:2.180 | בלשון |  |
| 6:2.181 | רבים |  |
| 6:2.182 | אע"פ |  |
| 6:2.183 | שהם |  |
| 6:2.184 | יחידים | , |
| 6:2.185 | כמו |  |
| 6:2.186 | אולי |  |
| 6:2.187 | אוכל |  |
| 6:2.188 | נכה |  |
| 6:2.189 | בו | . |
| 6:2.190 | ובדניאל |  |
| 6:2.191 | נאמר |  |
| 6:2.192 | קדם |  |
| 6:2.193 | מלכא |  |
| 6:2.194 | ובמנוח |  |
| 6:2.195 | נעצרה |  |
| 6:2.196 | נא |  |
| 6:2.197 | אותך |  |
| 6:2.198 | ע"כ |  |
| 6:2.199 | ועוד |  |
| 6:2.200 | האריך |  |
| 6:2.201 | רבי |  |
| 6:2.202 | סעדיה |  |
| 6:2.203 | גאון |  |
| 6:2.204 | מאד | , |
| 6:2.205 | ואין |  |
| 6:2.206 | צורך |  |
| 6:2.207 | להאריך |  |
| 6:2.208 | כל |  |
| 6:2.209 | כך | , |
| 6:2.210 | כי |  |
| 6:2.211 | כל |  |
| 6:2.212 | ישראל |  |
| 6:2.213 | תקועים |  |
| 6:2.214 | באמונה |  |
| 6:2.215 | חזקה | , |
| 6:2.216 | כי |  |
| 6:2.217 | יוצר |  |
| 6:2.218 | הכל |  |
| 6:2.219 | אחד |  |
| 6:2.220 | מיוחד |  |
| 6:2.221 | כמו |  |
| 6:2.222 | שנ | ' |
| 6:2.223 | ה | ' |
| 6:2.224 | אחד | . |
| 6:2.225 | ונאמר |  |
| 6:2.226 | ראו |  |
| 6:2.227 | עתה |  |
| 6:2.228 | כי |  |
| 6:2.229 | אני |  |
| 6:2.230 | הוא |  |
| 6:2.231 | ואין |  |
| 6:2.232 | אלהים |  |
| 6:2.233 | עמדי | . |
| 6:2.234 | ובמקרא |  |
| 6:2.235 | יש |  |
| 6:2.236 | הרבה |  |
| 6:2.237 | מקראות |  |
| 6:2.238 | כאלה |  |
| 6:2.239 | עכ"ל |  |
| 6:2.240 | הרב |  |
| 6:2.241 | הסמ"ג | . |


## Milchamot ha-Shem 6:3

Hebrew (Wikisource): סב) הבט נא אחי בעין חודרת בדברי הר' הסמ"ג ז"ל איך הביא דברי הרס"ג לחזק דבריו. ובעיניו הדבר פשוט מאוד, שכן היא אמונת אחדותו ית' אחדות שאין כמוה בשאר כל האחדים. שלא כדברי הקבלה החדשה שחשבו שאחדותו ית' פחותה גם משאר האחדים, כי הזוג האחד הם שנים בלבד, ואצל המקובלים החדשים האלהות כוללת כמה פרצופים כמבואר לעיל, חמשה פרצופים כוללים. והחמשה נפרטים עוד לשנים עשר פרצופים בעולם האצילות, חוץ מן הפרצופים שבעולמות העליונים אין מספר שלא דברו בהם. וחוץ מהפרצופים שבבריאה ויצירה ועשייה, ואלה הפרצופים כל אחד מהם נקרא בשם הוי'ה ואדנות ואלהים ובשם הקב"ה ומדרגותיהם חלוקים זה מזה, זה למעלה מזה, וזה למטה ממנו. ואין רשות לפרצוף לברוא או להחיות או להמית אלא בנטילת רשות מאותו פרצוף שלמעלה ממנו, כשרים הממונים מתחת יד המלך שהמלך נותן רשות למשנהו. והמשנה לשליש והשליש לשרי אלפים ושרי אלפים למאות וכו'. וכל אחד יעשה חפצו בדבר שנתמנה לו, אבל דבר גדול לא יעשה מדעת עצמו אלא בנטילת רשות מהממונה שלמעלה ממנו כמבואר לעיל סי' כ"א בשם זוהר בראשית וכמו שפירשו וביארו דבריו המפרשים, ושהמלך העליון מכולם אדם קדמון. והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי אני אמית ואחיה, לפי שאדם קדמון זה הוא תחלת פרצופי האצילות, ולית לעילה מניה דימחי ביה. ויש לתמוה על זה למה אינו צריך ליטול רשות מאדם קדמאה שהוא למעלה ממנו אשר כל העולמות העליונים שאין להם מספר נתונים בחלל עיגוליו כמו שכתב המקדש מלך דף ט"ו ד"ה ורמז, או מהאין סוף שהוא האלוה העליון מעל כולם?

| id | word | punctuation after |
|---|---|---|
| 6:3.1 | סב | ) |
| 6:3.2 | הבט |  |
| 6:3.3 | נא |  |
| 6:3.4 | אחי |  |
| 6:3.5 | בעין |  |
| 6:3.6 | חודרת |  |
| 6:3.7 | בדברי |  |
| 6:3.8 | הר | ' |
| 6:3.9 | הסמ"ג |  |
| 6:3.10 | ז"ל |  |
| 6:3.11 | איך |  |
| 6:3.12 | הביא |  |
| 6:3.13 | דברי |  |
| 6:3.14 | הרס"ג |  |
| 6:3.15 | לחזק |  |
| 6:3.16 | דבריו | . |
| 6:3.17 | ובעיניו |  |
| 6:3.18 | הדבר |  |
| 6:3.19 | פשוט |  |
| 6:3.20 | מאוד | , |
| 6:3.21 | שכן |  |
| 6:3.22 | היא |  |
| 6:3.23 | אמונת |  |
| 6:3.24 | אחדותו |  |
| 6:3.25 | ית | ' |
| 6:3.26 | אחדות |  |
| 6:3.27 | שאין |  |
| 6:3.28 | כמוה |  |
| 6:3.29 | בשאר |  |
| 6:3.30 | כל |  |
| 6:3.31 | האחדים | . |
| 6:3.32 | שלא |  |
| 6:3.33 | כדברי |  |
| 6:3.34 | הקבלה |  |
| 6:3.35 | החדשה |  |
| 6:3.36 | שחשבו |  |
| 6:3.37 | שאחדותו |  |
| 6:3.38 | ית | ' |
| 6:3.39 | פחותה |  |
| 6:3.40 | גם |  |
| 6:3.41 | משאר |  |
| 6:3.42 | האחדים | , |
| 6:3.43 | כי |  |
| 6:3.44 | הזוג |  |
| 6:3.45 | האחד |  |
| 6:3.46 | הם |  |
| 6:3.47 | שנים |  |
| 6:3.48 | בלבד | , |
| 6:3.49 | ואצל |  |
| 6:3.50 | המקובלים |  |
| 6:3.51 | החדשים |  |
| 6:3.52 | האלהות |  |
| 6:3.53 | כוללת |  |
| 6:3.54 | כמה |  |
| 6:3.55 | פרצופים |  |
| 6:3.56 | כמבואר |  |
| 6:3.57 | לעיל | , |
| 6:3.58 | חמשה |  |
| 6:3.59 | פרצופים |  |
| 6:3.60 | כוללים | . |
| 6:3.61 | והחמשה |  |
| 6:3.62 | נפרטים |  |
| 6:3.63 | עוד |  |
| 6:3.64 | לשנים |  |
| 6:3.65 | עשר |  |
| 6:3.66 | פרצופים |  |
| 6:3.67 | בעולם |  |
| 6:3.68 | האצילות | , |
| 6:3.69 | חוץ |  |
| 6:3.70 | מן |  |
| 6:3.71 | הפרצופים |  |
| 6:3.72 | שבעולמות |  |
| 6:3.73 | העליונים |  |
| 6:3.74 | אין |  |
| 6:3.75 | מספר |  |
| 6:3.76 | שלא |  |
| 6:3.77 | דברו |  |
| 6:3.78 | בהם | . |
| 6:3.79 | וחוץ |  |
| 6:3.80 | מהפרצופים |  |
| 6:3.81 | שבבריאה |  |
| 6:3.82 | ויצירה |  |
| 6:3.83 | ועשייה | , |
| 6:3.84 | ואלה |  |
| 6:3.85 | הפרצופים |  |
| 6:3.86 | כל |  |
| 6:3.87 | אחד |  |
| 6:3.88 | מהם |  |
| 6:3.89 | נקרא |  |
| 6:3.90 | בשם |  |
| 6:3.91 | הוי'ה |  |
| 6:3.92 | ואדנות |  |
| 6:3.93 | ואלהים |  |
| 6:3.94 | ובשם |  |
| 6:3.95 | הקב"ה |  |
| 6:3.96 | ומדרגותיהם |  |
| 6:3.97 | חלוקים |  |
| 6:3.98 | זה |  |
| 6:3.99 | מזה | , |
| 6:3.100 | זה |  |
| 6:3.101 | למעלה |  |
| 6:3.102 | מזה | , |
| 6:3.103 | וזה |  |
| 6:3.104 | למטה |  |
| 6:3.105 | ממנו | . |
| 6:3.106 | ואין |  |
| 6:3.107 | רשות |  |
| 6:3.108 | לפרצוף |  |
| 6:3.109 | לברוא |  |
| 6:3.110 | או |  |
| 6:3.111 | להחיות |  |
| 6:3.112 | או |  |
| 6:3.113 | להמית |  |
| 6:3.114 | אלא |  |
| 6:3.115 | בנטילת |  |
| 6:3.116 | רשות |  |
| 6:3.117 | מאותו |  |
| 6:3.118 | פרצוף |  |
| 6:3.119 | שלמעלה |  |
| 6:3.120 | ממנו | , |
| 6:3.121 | כשרים |  |
| 6:3.122 | הממונים |  |
| 6:3.123 | מתחת |  |
| 6:3.124 | יד |  |
| 6:3.125 | המלך |  |
| 6:3.126 | שהמלך |  |
| 6:3.127 | נותן |  |
| 6:3.128 | רשות |  |
| 6:3.129 | למשנהו | . |
| 6:3.130 | והמשנה |  |
| 6:3.131 | לשליש |  |
| 6:3.132 | והשליש |  |
| 6:3.133 | לשרי |  |
| 6:3.134 | אלפים |  |
| 6:3.135 | ושרי |  |
| 6:3.136 | אלפים |  |
| 6:3.137 | למאות |  |
| 6:3.138 | וכו | '. |
| 6:3.139 | וכל |  |
| 6:3.140 | אחד |  |
| 6:3.141 | יעשה |  |
| 6:3.142 | חפצו |  |
| 6:3.143 | בדבר |  |
| 6:3.144 | שנתמנה |  |
| 6:3.145 | לו | , |
| 6:3.146 | אבל |  |
| 6:3.147 | דבר |  |
| 6:3.148 | גדול |  |
| 6:3.149 | לא |  |
| 6:3.150 | יעשה |  |
| 6:3.151 | מדעת |  |
| 6:3.152 | עצמו |  |
| 6:3.153 | אלא |  |
| 6:3.154 | בנטילת |  |
| 6:3.155 | רשות |  |
| 6:3.156 | מהממונה |  |
| 6:3.157 | שלמעלה |  |
| 6:3.158 | ממנו |  |
| 6:3.159 | כמבואר |  |
| 6:3.160 | לעיל |  |
| 6:3.161 | סי | ' |
| 6:3.162 | כ"א |  |
| 6:3.163 | בשם |  |
| 6:3.164 | זוהר |  |
| 6:3.165 | בראשית |  |
| 6:3.166 | וכמו |  |
| 6:3.167 | שפירשו |  |
| 6:3.168 | וביארו |  |
| 6:3.169 | דבריו |  |
| 6:3.170 | המפרשים | , |
| 6:3.171 | ושהמלך |  |
| 6:3.172 | העליון |  |
| 6:3.173 | מכולם |  |
| 6:3.174 | אדם |  |
| 6:3.175 | קדמון | . |
| 6:3.176 | והוא |  |
| 6:3.177 | שאמר |  |
| 6:3.178 | ראו |  |
| 6:3.179 | עתה |  |
| 6:3.180 | כי |  |
| 6:3.181 | אני |  |
| 6:3.182 | אני |  |
| 6:3.183 | הוא |  |
| 6:3.184 | ואין |  |
| 6:3.185 | אלהים |  |
| 6:3.186 | עמדי |  |
| 6:3.187 | אני |  |
| 6:3.188 | אמית |  |
| 6:3.189 | ואחיה | , |
| 6:3.190 | לפי |  |
| 6:3.191 | שאדם |  |
| 6:3.192 | קדמון |  |
| 6:3.193 | זה |  |
| 6:3.194 | הוא |  |
| 6:3.195 | תחלת |  |
| 6:3.196 | פרצופי |  |
| 6:3.197 | האצילות | , |
| 6:3.198 | ולית |  |
| 6:3.199 | לעילה |  |
| 6:3.200 | מניה |  |
| 6:3.201 | דימחי |  |
| 6:3.202 | ביה | . |
| 6:3.203 | ויש |  |
| 6:3.204 | לתמוה |  |
| 6:3.205 | על |  |
| 6:3.206 | זה |  |
| 6:3.207 | למה |  |
| 6:3.208 | אינו |  |
| 6:3.209 | צריך |  |
| 6:3.210 | ליטול |  |
| 6:3.211 | רשות |  |
| 6:3.212 | מאדם |  |
| 6:3.213 | קדמאה |  |
| 6:3.214 | שהוא |  |
| 6:3.215 | למעלה |  |
| 6:3.216 | ממנו |  |
| 6:3.217 | אשר |  |
| 6:3.218 | כל |  |
| 6:3.219 | העולמות |  |
| 6:3.220 | העליונים |  |
| 6:3.221 | שאין |  |
| 6:3.222 | להם |  |
| 6:3.223 | מספר |  |
| 6:3.224 | נתונים |  |
| 6:3.225 | בחלל |  |
| 6:3.226 | עיגוליו |  |
| 6:3.227 | כמו |  |
| 6:3.228 | שכתב |  |
| 6:3.229 | המקדש |  |
| 6:3.230 | מלך |  |
| 6:3.231 | דף |  |
| 6:3.232 | ט"ו |  |
| 6:3.233 | ד"ה |  |
| 6:3.234 | ורמז | , |
| 6:3.235 | או |  |
| 6:3.236 | מהאין |  |
| 6:3.237 | סוף |  |
| 6:3.238 | שהוא |  |
| 6:3.239 | האלוה |  |
| 6:3.240 | העליון |  |
| 6:3.241 | מעל |  |
| 6:3.242 | כולם | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

