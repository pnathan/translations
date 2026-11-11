# Translation pack: milhamot-hashem perek-4-12 (Milchamot ha-Shem 4:20–4:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-12`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
- ר"א: "Rabbi Eleazar" only where the paragraph names him, else "R. A." with
  a note.
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
- מרגל לרגל: "from festival to festival" (Num 20 ff. sense), not "foot".
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
- האלהי before a name: "the divine", with a note; never dropped.
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
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".

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


## Milchamot ha-Shem 4:20

Hebrew (Wikisource): מה מאד יחרד לבו ויתר ממקומו בשמעו את מערכת האלהות הרבים, אשר פרו ורבו בישראל מתחלת אלף הששי אצל קצת המחברים, ומדור לדור מזמן הגז"ל הלכה האמונה הזאת הלך וגדול מאד להאמין בעילות רבות זו למעלה מזו, וברצות אחת מהעילות לברוא דבר הוא נמלך ונוטל רשות מן העילה שלמעלה ממנה כמבואר בזהר בהדיא (בראשית כ"ב) שכל אחת מן העלות נוטלת רשות מן העילה שעליה, מלכות מזעיר, וזעיר מאימא, ואימא מאבא, ואבא מאריך, ואריך מעתיק, ועתיק מאדם קדמון. כי הוא הראש לפרצופי האצילות והוא הוא לבדו שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאינו צריך ליטול רשות מאדם קדמאה שלמעלה ממנו. ובכל מעשה בראשית היה המלך הבורא אבא. וקורא לו בזהר מלכה עלאה, ואימא אומנא ובעת בריאת אדם הראשון לא חפץ אבא לבראתו, כיון שהוא עתיד לחטוא, ואימא השיבתו הואיל וחטאו תלוי בי, כדכתיב, ובן כסיל תוגת אמו, לא איכפת לך בזה כמבואר שם ובמקדש מלך.

| id | word | punctuation after |
|---|---|---|
| 4:20.1 | מה |  |
| 4:20.2 | מאד |  |
| 4:20.3 | יחרד |  |
| 4:20.4 | לבו |  |
| 4:20.5 | ויתר |  |
| 4:20.6 | ממקומו |  |
| 4:20.7 | בשמעו |  |
| 4:20.8 | את |  |
| 4:20.9 | מערכת |  |
| 4:20.10 | האלהות |  |
| 4:20.11 | הרבים | , |
| 4:20.12 | אשר |  |
| 4:20.13 | פרו |  |
| 4:20.14 | ורבו |  |
| 4:20.15 | בישראל |  |
| 4:20.16 | מתחלת |  |
| 4:20.17 | אלף |  |
| 4:20.18 | הששי |  |
| 4:20.19 | אצל |  |
| 4:20.20 | קצת |  |
| 4:20.21 | המחברים | , |
| 4:20.22 | ומדור |  |
| 4:20.23 | לדור |  |
| 4:20.24 | מזמן |  |
| 4:20.25 | הגז"ל |  |
| 4:20.26 | הלכה |  |
| 4:20.27 | האמונה |  |
| 4:20.28 | הזאת |  |
| 4:20.29 | הלך |  |
| 4:20.30 | וגדול |  |
| 4:20.31 | מאד |  |
| 4:20.32 | להאמין |  |
| 4:20.33 | בעילות |  |
| 4:20.34 | רבות |  |
| 4:20.35 | זו |  |
| 4:20.36 | למעלה |  |
| 4:20.37 | מזו | , |
| 4:20.38 | וברצות |  |
| 4:20.39 | אחת |  |
| 4:20.40 | מהעילות |  |
| 4:20.41 | לברוא |  |
| 4:20.42 | דבר |  |
| 4:20.43 | הוא |  |
| 4:20.44 | נמלך |  |
| 4:20.45 | ונוטל |  |
| 4:20.46 | רשות |  |
| 4:20.47 | מן |  |
| 4:20.48 | העילה |  |
| 4:20.49 | שלמעלה |  |
| 4:20.50 | ממנה |  |
| 4:20.51 | כמבואר |  |
| 4:20.52 | בזהר |  |
| 4:20.53 | בהדיא |  |
| 4:20.54 | בראשית |  |
| 4:20.55 | כ"ב | ) |
| 4:20.56 | שכל |  |
| 4:20.57 | אחת |  |
| 4:20.58 | מן |  |
| 4:20.59 | העלות |  |
| 4:20.60 | נוטלת |  |
| 4:20.61 | רשות |  |
| 4:20.62 | מן |  |
| 4:20.63 | העילה |  |
| 4:20.64 | שעליה | , |
| 4:20.65 | מלכות |  |
| 4:20.66 | מזעיר | , |
| 4:20.67 | וזעיר |  |
| 4:20.68 | מאימא | , |
| 4:20.69 | ואימא |  |
| 4:20.70 | מאבא | , |
| 4:20.71 | ואבא |  |
| 4:20.72 | מאריך | , |
| 4:20.73 | ואריך |  |
| 4:20.74 | מעתיק | , |
| 4:20.75 | ועתיק |  |
| 4:20.76 | מאדם |  |
| 4:20.77 | קדמון | . |
| 4:20.78 | כי |  |
| 4:20.79 | הוא |  |
| 4:20.80 | הראש |  |
| 4:20.81 | לפרצופי |  |
| 4:20.82 | האצילות |  |
| 4:20.83 | והוא |  |
| 4:20.84 | הוא |  |
| 4:20.85 | לבדו |  |
| 4:20.86 | שאמר |  |
| 4:20.87 | ראו |  |
| 4:20.88 | עתה |  |
| 4:20.89 | כי |  |
| 4:20.90 | אני |  |
| 4:20.91 | אני |  |
| 4:20.92 | הוא |  |
| 4:20.93 | ואין |  |
| 4:20.94 | אלהים |  |
| 4:20.95 | עמדי | , |
| 4:20.96 | שאינו |  |
| 4:20.97 | צריך |  |
| 4:20.98 | ליטול |  |
| 4:20.99 | רשות |  |
| 4:20.100 | מאדם |  |
| 4:20.101 | קדמאה |  |
| 4:20.102 | שלמעלה |  |
| 4:20.103 | ממנו | . |
| 4:20.104 | ובכל |  |
| 4:20.105 | מעשה |  |
| 4:20.106 | בראשית |  |
| 4:20.107 | היה |  |
| 4:20.108 | המלך |  |
| 4:20.109 | הבורא |  |
| 4:20.110 | אבא | . |
| 4:20.111 | וקורא |  |
| 4:20.112 | לו |  |
| 4:20.113 | בזהר |  |
| 4:20.114 | מלכה |  |
| 4:20.115 | עלאה | , |
| 4:20.116 | ואימא |  |
| 4:20.117 | אומנא |  |
| 4:20.118 | ובעת |  |
| 4:20.119 | בריאת |  |
| 4:20.120 | אדם |  |
| 4:20.121 | הראשון |  |
| 4:20.122 | לא |  |
| 4:20.123 | חפץ |  |
| 4:20.124 | אבא |  |
| 4:20.125 | לבראתו | , |
| 4:20.126 | כיון |  |
| 4:20.127 | שהוא |  |
| 4:20.128 | עתיד |  |
| 4:20.129 | לחטוא | , |
| 4:20.130 | ואימא |  |
| 4:20.131 | השיבתו |  |
| 4:20.132 | הואיל |  |
| 4:20.133 | וחטאו |  |
| 4:20.134 | תלוי |  |
| 4:20.135 | בי | , |
| 4:20.136 | כדכתיב | , |
| 4:20.137 | ובן |  |
| 4:20.138 | כסיל |  |
| 4:20.139 | תוגת |  |
| 4:20.140 | אמו | , |
| 4:20.141 | לא |  |
| 4:20.142 | איכפת |  |
| 4:20.143 | לך |  |
| 4:20.144 | בזה |  |
| 4:20.145 | כמבואר |  |
| 4:20.146 | שם |  |
| 4:20.147 | ובמקדש |  |
| 4:20.148 | מלך | . |


## Milchamot ha-Shem 4:21

Hebrew (Wikisource): לא) והדבר הזה קשה עד מאד למה אדם קדמון אין צריך ליטול רשות מאדם קדמאה ולא מא"ס שהוא למעלה ממנו? וכבר הבאתי לעיל לשון ספר הברית משם יושר לבב שאמר ששם הקב"ה המורגל בפינו ראי על זעיר אנפין. וכן שם הוי'ה, וכ"כ מצרף האמונה. רק שלא העתקתי דבריו, כיון שאתה מחלל אותו וחושבו למין וכופר מחמת חסרון בקיאותך. ובודאי שטעית בזה כמו שכתבתי לעיל שכן סוברים כל המקובלים. ובבראשית רבה פרשה י"ב הובא ג"כ בילקוט רמז י"ח וז"ל לכל יש תולדות שמים וארץ יש להן תולדות שנ' אלה תולדות השמים והארץ בהבראם. הרים יש להם תולדות שנ' בטרם הרים יולדו. מטר יש לו תולדות שנ' היש למטר אב. טל יש לו תולדות שנ' או מי הוליד אגלי טל. ר' שמעון בן לקיש אמר אלו מרגליות של טל. תני כל מי שיש לו תולדות מת ובולה ונברא ואני בורא וכל מי שאין לו תולדות אינו לא מת ולא בולה. ובורא ואינו נברא ר' עזריה בשם ר' אמר כלפי מעלה הדבר אמור ע"כ. וזה כלל אמתי וברור. ופי' הר' עץ יוסף מת ובולה הם שני דברים האחד ההפסד והשני הבלייה בעודן קיימין. שכל הנמצאים תמיד חסרים והולכים בטבעם לולי שהקב"ה מחדש בטובו בכל יום תמיד מעשה בראשית וכל מי שאין לו תולדות פי' אע"ג דליכא אלא הקב"ה קאמר כל איידי דתני נושא כל . וכמו שפי' ר' עזריה בשם רבי כלפי מעלה הדבר אמור ע"כ. וממוצא דבר אתה למד שכיון שיש לפרצופים הנזכרים בדברי הזהר והמקובלים תולדות ועיבור ויניקה. שקר הוא ומן הנמנע לקרות לאבא או לזולתו מכל הפרצופים בורא. אי נמי אפשר לפרש מי שיש לו תולדות דהיינו שיש סבה להוייתו שהוא אפשרי המציאות. וכל אפשרי המציאות הוא אפשר בו ההפסד אבל אלהינו ברוך הוא. אין סבה אחרת קודמת לו והוא לבדו סבה א' לכל הברואים. כמו שכ' הרמב"ם ז"ל בהל' יסודי התורה ע"ש ובמפרש, וכן כתבו קדמונינו על מה שכתב רבי' בעיקר הרביעי שנאמין שזה האחד הוא קדמון בהחלט וכל נמצא זולתו בלתי קדמון וכו' כתבו הר' ר' חטר ורבי זכריה הרופא וז"ל הקדמון הוא אשר אין ראשית לו והמחודש הוא שבראו והמציאו זולתו. הרי נתברך שהאלוה הוא הראשון בהחלט ע"ש. והכי מסתברא לפרש דברי המדרש רבה כנראה ממאי דקאמר בהרים ובמטר ובטל. וממילא שקר הוא לקרות שום אחד מן הפרצופים אלוה, או בורא, כיון שכל א' יש סבה אחרת קדמה לו. ובמשנת חסידים מס' ליל פסח פרק ז' משנה כ"ג בפי' הגדת ליל פסח כ' וז"ל ויוציאנו ה' אלהינו שהם אבא ואימא וכו' ואלו לא גאל הקב"ה, שהוא אריך אנפין את אבותינו שהם אבא ואימא ממצרים וכו' ובס' הנדרים כ' בערך אדם כי המקובלים קורין להקב"ה אדם קדמון. ואם כן, כל אחד ואחד מן הפרצופים נקרא ה' צבאות והוי'ה ואדנות והקב"ה ושאר השמות והכנויים המיוחדים לשי"ת. כאשר תחזינה עיניך פה וכ"ש אם תלמוד בספריהם בשום שכל. וגם קורין לאבא ואבא אבותינו שנגאלו ממצרים על ידי אריך אנפין מיד פרעה (הוא הסטרא אחרא הנקרא אל אחר) שנאלם אריך אנפין מידו וזה נמשך לפי שטתם שהם מאמינים במציאות אל אחר היפך דברי רז"ל וגם האמינו שהאלהים האחרים המדומים שלטו בהאלוהים הקדושים שהם אבא ואימא וזעיר ונוקביה, וגאלם אריך אנפין מהם.

| id | word | punctuation after |
|---|---|---|
| 4:21.1 | לא | ) |
| 4:21.2 | והדבר |  |
| 4:21.3 | הזה |  |
| 4:21.4 | קשה |  |
| 4:21.5 | עד |  |
| 4:21.6 | מאד |  |
| 4:21.7 | למה |  |
| 4:21.8 | אדם |  |
| 4:21.9 | קדמון |  |
| 4:21.10 | אין |  |
| 4:21.11 | צריך |  |
| 4:21.12 | ליטול |  |
| 4:21.13 | רשות |  |
| 4:21.14 | מאדם |  |
| 4:21.15 | קדמאה |  |
| 4:21.16 | ולא |  |
| 4:21.17 | מא"ס |  |
| 4:21.18 | שהוא |  |
| 4:21.19 | למעלה |  |
| 4:21.20 | ממנו | ? |
| 4:21.21 | וכבר |  |
| 4:21.22 | הבאתי |  |
| 4:21.23 | לעיל |  |
| 4:21.24 | לשון |  |
| 4:21.25 | ספר |  |
| 4:21.26 | הברית |  |
| 4:21.27 | משם |  |
| 4:21.28 | יושר |  |
| 4:21.29 | לבב |  |
| 4:21.30 | שאמר |  |
| 4:21.31 | ששם |  |
| 4:21.32 | הקב"ה |  |
| 4:21.33 | המורגל |  |
| 4:21.34 | בפינו |  |
| 4:21.35 | ראי |  |
| 4:21.36 | על |  |
| 4:21.37 | זעיר |  |
| 4:21.38 | אנפין | . |
| 4:21.39 | וכן |  |
| 4:21.40 | שם |  |
| 4:21.41 | הוי'ה | , |
| 4:21.42 | וכ"כ |  |
| 4:21.43 | מצרף |  |
| 4:21.44 | האמונה | . |
| 4:21.45 | רק |  |
| 4:21.46 | שלא |  |
| 4:21.47 | העתקתי |  |
| 4:21.48 | דבריו | , |
| 4:21.49 | כיון |  |
| 4:21.50 | שאתה |  |
| 4:21.51 | מחלל |  |
| 4:21.52 | אותו |  |
| 4:21.53 | וחושבו |  |
| 4:21.54 | למין |  |
| 4:21.55 | וכופר |  |
| 4:21.56 | מחמת |  |
| 4:21.57 | חסרון |  |
| 4:21.58 | בקיאותך | . |
| 4:21.59 | ובודאי |  |
| 4:21.60 | שטעית |  |
| 4:21.61 | בזה |  |
| 4:21.62 | כמו |  |
| 4:21.63 | שכתבתי |  |
| 4:21.64 | לעיל |  |
| 4:21.65 | שכן |  |
| 4:21.66 | סוברים |  |
| 4:21.67 | כל |  |
| 4:21.68 | המקובלים | . |
| 4:21.69 | ובבראשית |  |
| 4:21.70 | רבה |  |
| 4:21.71 | פרשה |  |
| 4:21.72 | י"ב |  |
| 4:21.73 | הובא |  |
| 4:21.74 | ג"כ |  |
| 4:21.75 | בילקוט |  |
| 4:21.76 | רמז |  |
| 4:21.77 | י"ח |  |
| 4:21.78 | וז"ל |  |
| 4:21.79 | לכל |  |
| 4:21.80 | יש |  |
| 4:21.81 | תולדות |  |
| 4:21.82 | שמים |  |
| 4:21.83 | וארץ |  |
| 4:21.84 | יש |  |
| 4:21.85 | להן |  |
| 4:21.86 | תולדות |  |
| 4:21.87 | שנ | ' |
| 4:21.88 | אלה |  |
| 4:21.89 | תולדות |  |
| 4:21.90 | השמים |  |
| 4:21.91 | והארץ |  |
| 4:21.92 | בהבראם | . |
| 4:21.93 | הרים |  |
| 4:21.94 | יש |  |
| 4:21.95 | להם |  |
| 4:21.96 | תולדות |  |
| 4:21.97 | שנ | ' |
| 4:21.98 | בטרם |  |
| 4:21.99 | הרים |  |
| 4:21.100 | יולדו | . |
| 4:21.101 | מטר |  |
| 4:21.102 | יש |  |
| 4:21.103 | לו |  |
| 4:21.104 | תולדות |  |
| 4:21.105 | שנ | ' |
| 4:21.106 | היש |  |
| 4:21.107 | למטר |  |
| 4:21.108 | אב | . |
| 4:21.109 | טל |  |
| 4:21.110 | יש |  |
| 4:21.111 | לו |  |
| 4:21.112 | תולדות |  |
| 4:21.113 | שנ | ' |
| 4:21.114 | או |  |
| 4:21.115 | מי |  |
| 4:21.116 | הוליד |  |
| 4:21.117 | אגלי |  |
| 4:21.118 | טל | . |
| 4:21.119 | ר | ' |
| 4:21.120 | שמעון |  |
| 4:21.121 | בן |  |
| 4:21.122 | לקיש |  |
| 4:21.123 | אמר |  |
| 4:21.124 | אלו |  |
| 4:21.125 | מרגליות |  |
| 4:21.126 | של |  |
| 4:21.127 | טל | . |
| 4:21.128 | תני |  |
| 4:21.129 | כל |  |
| 4:21.130 | מי |  |
| 4:21.131 | שיש |  |
| 4:21.132 | לו |  |
| 4:21.133 | תולדות |  |
| 4:21.134 | מת |  |
| 4:21.135 | ובולה |  |
| 4:21.136 | ונברא |  |
| 4:21.137 | ואני |  |
| 4:21.138 | בורא |  |
| 4:21.139 | וכל |  |
| 4:21.140 | מי |  |
| 4:21.141 | שאין |  |
| 4:21.142 | לו |  |
| 4:21.143 | תולדות |  |
| 4:21.144 | אינו |  |
| 4:21.145 | לא |  |
| 4:21.146 | מת |  |
| 4:21.147 | ולא |  |
| 4:21.148 | בולה | . |
| 4:21.149 | ובורא |  |
| 4:21.150 | ואינו |  |
| 4:21.151 | נברא |  |
| 4:21.152 | ר | ' |
| 4:21.153 | עזריה |  |
| 4:21.154 | בשם |  |
| 4:21.155 | ר | ' |
| 4:21.156 | אמר |  |
| 4:21.157 | כלפי |  |
| 4:21.158 | מעלה |  |
| 4:21.159 | הדבר |  |
| 4:21.160 | אמור |  |
| 4:21.161 | ע"כ | . |
| 4:21.162 | וזה |  |
| 4:21.163 | כלל |  |
| 4:21.164 | אמתי |  |
| 4:21.165 | וברור | . |
| 4:21.166 | ופי | ' |
| 4:21.167 | הר | ' |
| 4:21.168 | עץ |  |
| 4:21.169 | יוסף |  |
| 4:21.170 | מת |  |
| 4:21.171 | ובולה |  |
| 4:21.172 | הם |  |
| 4:21.173 | שני |  |
| 4:21.174 | דברים |  |
| 4:21.175 | האחד |  |
| 4:21.176 | ההפסד |  |
| 4:21.177 | והשני |  |
| 4:21.178 | הבלייה |  |
| 4:21.179 | בעודן |  |
| 4:21.180 | קיימין | . |
| 4:21.181 | שכל |  |
| 4:21.182 | הנמצאים |  |
| 4:21.183 | תמיד |  |
| 4:21.184 | חסרים |  |
| 4:21.185 | והולכים |  |
| 4:21.186 | בטבעם |  |
| 4:21.187 | לולי |  |
| 4:21.188 | שהקב"ה |  |
| 4:21.189 | מחדש |  |
| 4:21.190 | בטובו |  |
| 4:21.191 | בכל |  |
| 4:21.192 | יום |  |
| 4:21.193 | תמיד |  |
| 4:21.194 | מעשה |  |
| 4:21.195 | בראשית |  |
| 4:21.196 | וכל |  |
| 4:21.197 | מי |  |
| 4:21.198 | שאין |  |
| 4:21.199 | לו |  |
| 4:21.200 | תולדות |  |
| 4:21.201 | פי | ' |
| 4:21.202 | אע"ג |  |
| 4:21.203 | דליכא |  |
| 4:21.204 | אלא |  |
| 4:21.205 | הקב"ה |  |
| 4:21.206 | קאמר |  |
| 4:21.207 | כל |  |
| 4:21.208 | איידי |  |
| 4:21.209 | דתני |  |
| 4:21.210 | נושא |  |
| 4:21.211 | כל | . |
| 4:21.212 | וכמו |  |
| 4:21.213 | שפי | ' |
| 4:21.214 | ר | ' |
| 4:21.215 | עזריה |  |
| 4:21.216 | בשם |  |
| 4:21.217 | רבי |  |
| 4:21.218 | כלפי |  |
| 4:21.219 | מעלה |  |
| 4:21.220 | הדבר |  |
| 4:21.221 | אמור |  |
| 4:21.222 | ע"כ | . |
| 4:21.223 | וממוצא |  |
| 4:21.224 | דבר |  |
| 4:21.225 | אתה |  |
| 4:21.226 | למד |  |
| 4:21.227 | שכיון |  |
| 4:21.228 | שיש |  |
| 4:21.229 | לפרצופים |  |
| 4:21.230 | הנזכרים |  |
| 4:21.231 | בדברי |  |
| 4:21.232 | הזהר |  |
| 4:21.233 | והמקובלים |  |
| 4:21.234 | תולדות |  |
| 4:21.235 | ועיבור |  |
| 4:21.236 | ויניקה | . |
| 4:21.237 | שקר |  |
| 4:21.238 | הוא |  |
| 4:21.239 | ומן |  |
| 4:21.240 | הנמנע |  |
| 4:21.241 | לקרות |  |
| 4:21.242 | לאבא |  |
| 4:21.243 | או |  |
| 4:21.244 | לזולתו |  |
| 4:21.245 | מכל |  |
| 4:21.246 | הפרצופים |  |
| 4:21.247 | בורא | . |
| 4:21.248 | אי |  |
| 4:21.249 | נמי |  |
| 4:21.250 | אפשר |  |
| 4:21.251 | לפרש |  |
| 4:21.252 | מי |  |
| 4:21.253 | שיש |  |
| 4:21.254 | לו |  |
| 4:21.255 | תולדות |  |
| 4:21.256 | דהיינו |  |
| 4:21.257 | שיש |  |
| 4:21.258 | סבה |  |
| 4:21.259 | להוייתו |  |
| 4:21.260 | שהוא |  |
| 4:21.261 | אפשרי |  |
| 4:21.262 | המציאות | . |
| 4:21.263 | וכל |  |
| 4:21.264 | אפשרי |  |
| 4:21.265 | המציאות |  |
| 4:21.266 | הוא |  |
| 4:21.267 | אפשר |  |
| 4:21.268 | בו |  |
| 4:21.269 | ההפסד |  |
| 4:21.270 | אבל |  |
| 4:21.271 | אלהינו |  |
| 4:21.272 | ברוך |  |
| 4:21.273 | הוא | . |
| 4:21.274 | אין |  |
| 4:21.275 | סבה |  |
| 4:21.276 | אחרת |  |
| 4:21.277 | קודמת |  |
| 4:21.278 | לו |  |
| 4:21.279 | והוא |  |
| 4:21.280 | לבדו |  |
| 4:21.281 | סבה |  |
| 4:21.282 | א | ' |
| 4:21.283 | לכל |  |
| 4:21.284 | הברואים | . |
| 4:21.285 | כמו |  |
| 4:21.286 | שכ | ' |
| 4:21.287 | הרמב"ם |  |
| 4:21.288 | ז"ל |  |
| 4:21.289 | בהל | ' |
| 4:21.290 | יסודי |  |
| 4:21.291 | התורה |  |
| 4:21.292 | ע"ש |  |
| 4:21.293 | ובמפרש | , |
| 4:21.294 | וכן |  |
| 4:21.295 | כתבו |  |
| 4:21.296 | קדמונינו |  |
| 4:21.297 | על |  |
| 4:21.298 | מה |  |
| 4:21.299 | שכתב |  |
| 4:21.300 | רבי | ' |
| 4:21.301 | בעיקר |  |
| 4:21.302 | הרביעי |  |
| 4:21.303 | שנאמין |  |
| 4:21.304 | שזה |  |
| 4:21.305 | האחד |  |
| 4:21.306 | הוא |  |
| 4:21.307 | קדמון |  |
| 4:21.308 | בהחלט |  |
| 4:21.309 | וכל |  |
| 4:21.310 | נמצא |  |
| 4:21.311 | זולתו |  |
| 4:21.312 | בלתי |  |
| 4:21.313 | קדמון |  |
| 4:21.314 | וכו | ' |
| 4:21.315 | כתבו |  |
| 4:21.316 | הר | ' |
| 4:21.317 | ר | ' |
| 4:21.318 | חטר |  |
| 4:21.319 | ורבי |  |
| 4:21.320 | זכריה |  |
| 4:21.321 | הרופא |  |
| 4:21.322 | וז"ל |  |
| 4:21.323 | הקדמון |  |
| 4:21.324 | הוא |  |
| 4:21.325 | אשר |  |
| 4:21.326 | אין |  |
| 4:21.327 | ראשית |  |
| 4:21.328 | לו |  |
| 4:21.329 | והמחודש |  |
| 4:21.330 | הוא |  |
| 4:21.331 | שבראו |  |
| 4:21.332 | והמציאו |  |
| 4:21.333 | זולתו | . |
| 4:21.334 | הרי |  |
| 4:21.335 | נתברך |  |
| 4:21.336 | שהאלוה |  |
| 4:21.337 | הוא |  |
| 4:21.338 | הראשון |  |
| 4:21.339 | בהחלט |  |
| 4:21.340 | ע"ש | . |
| 4:21.341 | והכי |  |
| 4:21.342 | מסתברא |  |
| 4:21.343 | לפרש |  |
| 4:21.344 | דברי |  |
| 4:21.345 | המדרש |  |
| 4:21.346 | רבה |  |
| 4:21.347 | כנראה |  |
| 4:21.348 | ממאי |  |
| 4:21.349 | דקאמר |  |
| 4:21.350 | בהרים |  |
| 4:21.351 | ובמטר |  |
| 4:21.352 | ובטל | . |
| 4:21.353 | וממילא |  |
| 4:21.354 | שקר |  |
| 4:21.355 | הוא |  |
| 4:21.356 | לקרות |  |
| 4:21.357 | שום |  |
| 4:21.358 | אחד |  |
| 4:21.359 | מן |  |
| 4:21.360 | הפרצופים |  |
| 4:21.361 | אלוה | , |
| 4:21.362 | או |  |
| 4:21.363 | בורא | , |
| 4:21.364 | כיון |  |
| 4:21.365 | שכל |  |
| 4:21.366 | א | ' |
| 4:21.367 | יש |  |
| 4:21.368 | סבה |  |
| 4:21.369 | אחרת |  |
| 4:21.370 | קדמה |  |
| 4:21.371 | לו | . |
| 4:21.372 | ובמשנת |  |
| 4:21.373 | חסידים |  |
| 4:21.374 | מס | ' |
| 4:21.375 | ליל |  |
| 4:21.376 | פסח |  |
| 4:21.377 | פרק |  |
| 4:21.378 | ז | ' |
| 4:21.379 | משנה |  |
| 4:21.380 | כ"ג |  |
| 4:21.381 | בפי | ' |
| 4:21.382 | הגדת |  |
| 4:21.383 | ליל |  |
| 4:21.384 | פסח |  |
| 4:21.385 | כ | ' |
| 4:21.386 | וז"ל |  |
| 4:21.387 | ויוציאנו |  |
| 4:21.388 | ה | ' |
| 4:21.389 | אלהינו |  |
| 4:21.390 | שהם |  |
| 4:21.391 | אבא |  |
| 4:21.392 | ואימא |  |
| 4:21.393 | וכו | ' |
| 4:21.394 | ואלו |  |
| 4:21.395 | לא |  |
| 4:21.396 | גאל |  |
| 4:21.397 | הקב"ה | , |
| 4:21.398 | שהוא |  |
| 4:21.399 | אריך |  |
| 4:21.400 | אנפין |  |
| 4:21.401 | את |  |
| 4:21.402 | אבותינו |  |
| 4:21.403 | שהם |  |
| 4:21.404 | אבא |  |
| 4:21.405 | ואימא |  |
| 4:21.406 | ממצרים |  |
| 4:21.407 | וכו | ' |
| 4:21.408 | ובס | ' |
| 4:21.409 | הנדרים |  |
| 4:21.410 | כ | ' |
| 4:21.411 | בערך |  |
| 4:21.412 | אדם |  |
| 4:21.413 | כי |  |
| 4:21.414 | המקובלים |  |
| 4:21.415 | קורין |  |
| 4:21.416 | להקב"ה |  |
| 4:21.417 | אדם |  |
| 4:21.418 | קדמון | . |
| 4:21.419 | ואם |  |
| 4:21.420 | כן | , |
| 4:21.421 | כל |  |
| 4:21.422 | אחד |  |
| 4:21.423 | ואחד |  |
| 4:21.424 | מן |  |
| 4:21.425 | הפרצופים |  |
| 4:21.426 | נקרא |  |
| 4:21.427 | ה | ' |
| 4:21.428 | צבאות |  |
| 4:21.429 | והוי'ה |  |
| 4:21.430 | ואדנות |  |
| 4:21.431 | והקב"ה |  |
| 4:21.432 | ושאר |  |
| 4:21.433 | השמות |  |
| 4:21.434 | והכנויים |  |
| 4:21.435 | המיוחדים |  |
| 4:21.436 | לשי"ת | . |
| 4:21.437 | כאשר |  |
| 4:21.438 | תחזינה |  |
| 4:21.439 | עיניך |  |
| 4:21.440 | פה |  |
| 4:21.441 | וכ"ש |  |
| 4:21.442 | אם |  |
| 4:21.443 | תלמוד |  |
| 4:21.444 | בספריהם |  |
| 4:21.445 | בשום |  |
| 4:21.446 | שכל | . |
| 4:21.447 | וגם |  |
| 4:21.448 | קורין |  |
| 4:21.449 | לאבא |  |
| 4:21.450 | ואבא |  |
| 4:21.451 | אבותינו |  |
| 4:21.452 | שנגאלו |  |
| 4:21.453 | ממצרים |  |
| 4:21.454 | על |  |
| 4:21.455 | ידי |  |
| 4:21.456 | אריך |  |
| 4:21.457 | אנפין |  |
| 4:21.458 | מיד |  |
| 4:21.459 | פרעה |  |
| 4:21.460 | הוא |  |
| 4:21.461 | הסטרא |  |
| 4:21.462 | אחרא |  |
| 4:21.463 | הנקרא |  |
| 4:21.464 | אל |  |
| 4:21.465 | אחר | ) |
| 4:21.466 | שנאלם |  |
| 4:21.467 | אריך |  |
| 4:21.468 | אנפין |  |
| 4:21.469 | מידו |  |
| 4:21.470 | וזה |  |
| 4:21.471 | נמשך |  |
| 4:21.472 | לפי |  |
| 4:21.473 | שטתם |  |
| 4:21.474 | שהם |  |
| 4:21.475 | מאמינים |  |
| 4:21.476 | במציאות |  |
| 4:21.477 | אל |  |
| 4:21.478 | אחר |  |
| 4:21.479 | היפך |  |
| 4:21.480 | דברי |  |
| 4:21.481 | רז"ל |  |
| 4:21.482 | וגם |  |
| 4:21.483 | האמינו |  |
| 4:21.484 | שהאלהים |  |
| 4:21.485 | האחרים |  |
| 4:21.486 | המדומים |  |
| 4:21.487 | שלטו |  |
| 4:21.488 | בהאלוהים |  |
| 4:21.489 | הקדושים |  |
| 4:21.490 | שהם |  |
| 4:21.491 | אבא |  |
| 4:21.492 | ואימא |  |
| 4:21.493 | וזעיר |  |
| 4:21.494 | ונוקביה | , |
| 4:21.495 | וגאלם |  |
| 4:21.496 | אריך |  |
| 4:21.497 | אנפין |  |
| 4:21.498 | מהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

