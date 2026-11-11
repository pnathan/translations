# Translation pack: milhamot-hashem perek-5-19 (Milchamot ha-Shem 5:28–5:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-19`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
- **מדה / מדות / מידות**: context. Do not render it "sefirot" even where that
  is the referent.
  - The hermeneutical rules: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי"
    (1:6) → "the thirteen rules of Rabbi Ishmael and the thirty-two rules of
    Rabbi Yose the Galilean"; "שלש עשרה מדות שהתורה נדרשת בהם" (10:5) → "the
    thirteen rules by which the Torah is expounded". Add a note: literally
    "measures".
  - Exod 34:6–7: "שלש עשרה מדות של רחמים" (9:8) → "the thirteen attributes of
    mercy". These are a different thirteen from the rules. The author uses
    both, so never let one rendering cover the other.
  - Kabbalistic and philosophical attributes, God's or the sefirot's: "מדת
    הדין" → "the attribute of judgment"; "המדות העליונות" (2:11) → "the
    upper attributes"; "המידות הפנימיות" (2:7) → "the inner attributes";
    "מדת המלכות" → "the attribute of Malkhut"; "מלובש במידותיו" (5:15) →
    "clothed in His attributes"; "ממדותיו" (7:17, Exod 34:6) → "of His
    attributes".
  - Measure: "עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה" (4:32) → "ten
    measures and ten sefirot, because they have measure and limit"; "לא מקום
    ולא מידה" (7:13) → "neither place nor measure"; "במדה שמדדתם" →
    "with the measure with which you measured".
  - A quality or trait: "ממדת הענוה" (4:8) → "of the quality of humility";
    "המדה המגונה הזאת" (7:19) → "this disgraceful trait"; "אין לך מדה גדולה
    מזו" (6:14, Talmud) → "there is no greater virtue than this".
  - A proper name: מסכת מדות → "tractate Middot".
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
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted", with
  "tzimtzum" in a note at the first occurrence in a unit. "אחר הצמצום" →
  "after the contraction".
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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
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
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- והוא שלא: "provided that" (fixed).
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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


## Milchamot ha-Shem 5:28

Hebrew (Wikisource): נד) וזוהי אמונתינו וקבלתינו העיקרית שקבלנו מאבותינו דור אחר דור עד משה רבי' ועד אברהם אבינו. והיא שוה ומתאמת לקבלת הישמעאלים המקובלת אצלם ממחוקקם שקיבל מישראל כנז' לעיל. כי ממקור אחד מאברהם אבינו וכו' כנז"ל. וחיבה יתירה נודעת לנו. שקירבנו הקב"ה לפני הר סיני, ונגלה כבודו לנו שם, ונתן בקולו קול עוז לכל קהלינו לאמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים על פני. והנחילנו תורה ומצות וחקים. ואין להביא ראיה מדורות ימי הבינים שהאמינו בקבלה החדשה לפי שהיו לומדים ס' הזהר בגרסא בעלמא מבלי הבין כידוע. וקיימו בעצמם לגרוס איניש לבד, אבל והדר ליסבר לא קיימו, ולא העמיקו חקר בדברי הספר הנזכר ולא בשאר ספרי הקבלה החדשה להבין שהם הורסים פנות תורתינו הקדושה ועיקריה ויסודותיה שביארם הרמב"ם בפי' המשנה. ולפיכך הם אומרים יגדל אלהים חי בכל ליל שבת קודש ועוסקים בגירסת ס' הזהר ולא ידעו שהם סותרים זה את זה כמו שכ' רבי' תם ן' יחיא ז"ל. ואלו ידעו והבינו שהוא כן היו מרחיקים עצמם ממנה תכלית ההרחקה. כי עין רואה ואזן שומעת שהם מגשימים אפי' את האין סוף שהוא (לפי דבריהם) לעילא לעילא, ומציירים אותו ככדור עיגול וחלול באמצעיתו פנוי וריקם המתהוה על ידי הצמצום הנעשה בו לצדדין שוה בהשואה א'. כדי להמציא מקום שיוכלו להיות שם במקום הריקן כל הנאצלים והנבראים והנוצרים והנעשים. ואז נמשך מאור האין סוף קו דק מן האור העגול של האין סוף מלמעלה למטה ונעשו ספירות העיגולים מאותו הקו הנמשך ממנו עגול בתוך עגול. וגם נשתלשל ונמשך הקו ההוא ביושר תוך החלל ונעשו ספירות דיושר. אשר קראום מדות וכלים והם עיקר האלהות, וגם נשמתם שהוא האור הפנימי גם מלבושם שהוא האור המקיף הכל נמשך ונשתלשל מעצם האין סוף המקיף את הכל, ממנו גוף, ממנו נשמה, ממנו מלבוש. כלם נמשכים מעצמות האור ההוא כמו שכתבתי לעיל סי' ט"ל בשם שושן סודות ומשנת חסידים שגם הכלים אלהות גמור כי מחומר אחד קורצו. והסבות הללו נמשכים ונעשה כל אחד סבה למה שלמטה ממנה. זה סבה לזה וזה סבה לזה. הסבה הראשונה מפרצופי היושר ופרצופים הכוללים אדם קדמאה, ואחריו אדם קדמון, ואחריו עתיקא קדישא, ואחריו אריך אנפין, וממנו נחיו אבא ואימא, ומהם נהיו זעיר ונוקביה, והוא הנקרא בן לאבא ואימא. כל הפרצופים הללו נשתלשלו זה מזה, ונקרא כל אחד מהם עילה וסבה לפרצוף העלול ממנו אחריו. ואבא ואימא אעטרו לז"א בתרין עטרין והשליטוהו על כל הברואים וצוו (לפי דבריהם) לעבדו ליראה את אשתו, שריבה אותה הכתוב ליראה באמרו את ה' אלהיך תירא, את לרבות נוקבא דיליה ליראה אותה כמו שנתבאר לעיל סי' כ"ו כ"ז כ"ח, מה שלא עלה על דעת שמעון העמסוני ואמרי לה נחמיה העמסוני ור' עקיבה. ויחסו העבודה לזעיר אנפין דוקא כמו שנ' ואותו תעבוד והוא נקרא בפיהם יהוה אלהינו. והוא הסבה והפרצוף האחרון שבפרצופים האצילות. ואמרו שהוא הוא אלהינו ואנחנו עמו ועבדיו כמבואר בכסא אליהו וס' הברית ויושר לבב דף כ"א ול"א ול"ב. וענין השתלשלותם בעץ חיים שער עיגולים ויושר וכסא אלהיו.

| id | word | punctuation after |
|---|---|---|
| 5:28.1 | נד | ) |
| 5:28.2 | וזוהי |  |
| 5:28.3 | אמונתינו |  |
| 5:28.4 | וקבלתינו |  |
| 5:28.5 | העיקרית |  |
| 5:28.6 | שקבלנו |  |
| 5:28.7 | מאבותינו |  |
| 5:28.8 | דור |  |
| 5:28.9 | אחר |  |
| 5:28.10 | דור |  |
| 5:28.11 | עד |  |
| 5:28.12 | משה |  |
| 5:28.13 | רבי | ' |
| 5:28.14 | ועד |  |
| 5:28.15 | אברהם |  |
| 5:28.16 | אבינו | . |
| 5:28.17 | והיא |  |
| 5:28.18 | שוה |  |
| 5:28.19 | ומתאמת |  |
| 5:28.20 | לקבלת |  |
| 5:28.21 | הישמעאלים |  |
| 5:28.22 | המקובלת |  |
| 5:28.23 | אצלם |  |
| 5:28.24 | ממחוקקם |  |
| 5:28.25 | שקיבל |  |
| 5:28.26 | מישראל |  |
| 5:28.27 | כנז | ' |
| 5:28.28 | לעיל | . |
| 5:28.29 | כי |  |
| 5:28.30 | ממקור |  |
| 5:28.31 | אחד |  |
| 5:28.32 | מאברהם |  |
| 5:28.33 | אבינו |  |
| 5:28.34 | וכו | ' |
| 5:28.35 | כנז"ל | . |
| 5:28.36 | וחיבה |  |
| 5:28.37 | יתירה |  |
| 5:28.38 | נודעת |  |
| 5:28.39 | לנו | . |
| 5:28.40 | שקירבנו |  |
| 5:28.41 | הקב"ה |  |
| 5:28.42 | לפני |  |
| 5:28.43 | הר |  |
| 5:28.44 | סיני | , |
| 5:28.45 | ונגלה |  |
| 5:28.46 | כבודו |  |
| 5:28.47 | לנו |  |
| 5:28.48 | שם | , |
| 5:28.49 | ונתן |  |
| 5:28.50 | בקולו |  |
| 5:28.51 | קול |  |
| 5:28.52 | עוז |  |
| 5:28.53 | לכל |  |
| 5:28.54 | קהלינו |  |
| 5:28.55 | לאמר |  |
| 5:28.56 | אנכי |  |
| 5:28.57 | ה | ' |
| 5:28.58 | אלהיך | , |
| 5:28.59 | לא |  |
| 5:28.60 | יהיה |  |
| 5:28.61 | לך |  |
| 5:28.62 | אלהים |  |
| 5:28.63 | אחרים |  |
| 5:28.64 | על |  |
| 5:28.65 | פני | . |
| 5:28.66 | והנחילנו |  |
| 5:28.67 | תורה |  |
| 5:28.68 | ומצות |  |
| 5:28.69 | וחקים | . |
| 5:28.70 | ואין |  |
| 5:28.71 | להביא |  |
| 5:28.72 | ראיה |  |
| 5:28.73 | מדורות |  |
| 5:28.74 | ימי |  |
| 5:28.75 | הבינים |  |
| 5:28.76 | שהאמינו |  |
| 5:28.77 | בקבלה |  |
| 5:28.78 | החדשה |  |
| 5:28.79 | לפי |  |
| 5:28.80 | שהיו |  |
| 5:28.81 | לומדים |  |
| 5:28.82 | ס | ' |
| 5:28.83 | הזהר |  |
| 5:28.84 | בגרסא |  |
| 5:28.85 | בעלמא |  |
| 5:28.86 | מבלי |  |
| 5:28.87 | הבין |  |
| 5:28.88 | כידוע | . |
| 5:28.89 | וקיימו |  |
| 5:28.90 | בעצמם |  |
| 5:28.91 | לגרוס |  |
| 5:28.92 | איניש |  |
| 5:28.93 | לבד | , |
| 5:28.94 | אבל |  |
| 5:28.95 | והדר |  |
| 5:28.96 | ליסבר |  |
| 5:28.97 | לא |  |
| 5:28.98 | קיימו | , |
| 5:28.99 | ולא |  |
| 5:28.100 | העמיקו |  |
| 5:28.101 | חקר |  |
| 5:28.102 | בדברי |  |
| 5:28.103 | הספר |  |
| 5:28.104 | הנזכר |  |
| 5:28.105 | ולא |  |
| 5:28.106 | בשאר |  |
| 5:28.107 | ספרי |  |
| 5:28.108 | הקבלה |  |
| 5:28.109 | החדשה |  |
| 5:28.110 | להבין |  |
| 5:28.111 | שהם |  |
| 5:28.112 | הורסים |  |
| 5:28.113 | פנות |  |
| 5:28.114 | תורתינו |  |
| 5:28.115 | הקדושה |  |
| 5:28.116 | ועיקריה |  |
| 5:28.117 | ויסודותיה |  |
| 5:28.118 | שביארם |  |
| 5:28.119 | הרמב"ם |  |
| 5:28.120 | בפי | ' |
| 5:28.121 | המשנה | . |
| 5:28.122 | ולפיכך |  |
| 5:28.123 | הם |  |
| 5:28.124 | אומרים |  |
| 5:28.125 | יגדל |  |
| 5:28.126 | אלהים |  |
| 5:28.127 | חי |  |
| 5:28.128 | בכל |  |
| 5:28.129 | ליל |  |
| 5:28.130 | שבת |  |
| 5:28.131 | קודש |  |
| 5:28.132 | ועוסקים |  |
| 5:28.133 | בגירסת |  |
| 5:28.134 | ס | ' |
| 5:28.135 | הזהר |  |
| 5:28.136 | ולא |  |
| 5:28.137 | ידעו |  |
| 5:28.138 | שהם |  |
| 5:28.139 | סותרים |  |
| 5:28.140 | זה |  |
| 5:28.141 | את |  |
| 5:28.142 | זה |  |
| 5:28.143 | כמו |  |
| 5:28.144 | שכ | ' |
| 5:28.145 | רבי | ' |
| 5:28.146 | תם |  |
| 5:28.147 | ן | ' |
| 5:28.148 | יחיא |  |
| 5:28.149 | ז"ל | . |
| 5:28.150 | ואלו |  |
| 5:28.151 | ידעו |  |
| 5:28.152 | והבינו |  |
| 5:28.153 | שהוא |  |
| 5:28.154 | כן |  |
| 5:28.155 | היו |  |
| 5:28.156 | מרחיקים |  |
| 5:28.157 | עצמם |  |
| 5:28.158 | ממנה |  |
| 5:28.159 | תכלית |  |
| 5:28.160 | ההרחקה | . |
| 5:28.161 | כי |  |
| 5:28.162 | עין |  |
| 5:28.163 | רואה |  |
| 5:28.164 | ואזן |  |
| 5:28.165 | שומעת |  |
| 5:28.166 | שהם |  |
| 5:28.167 | מגשימים |  |
| 5:28.168 | אפי | ' |
| 5:28.169 | את |  |
| 5:28.170 | האין |  |
| 5:28.171 | סוף |  |
| 5:28.172 | שהוא |  |
| 5:28.173 | לפי |  |
| 5:28.174 | דבריהם | ) |
| 5:28.175 | לעילא |  |
| 5:28.176 | לעילא | , |
| 5:28.177 | ומציירים |  |
| 5:28.178 | אותו |  |
| 5:28.179 | ככדור |  |
| 5:28.180 | עיגול |  |
| 5:28.181 | וחלול |  |
| 5:28.182 | באמצעיתו |  |
| 5:28.183 | פנוי |  |
| 5:28.184 | וריקם |  |
| 5:28.185 | המתהוה |  |
| 5:28.186 | על |  |
| 5:28.187 | ידי |  |
| 5:28.188 | הצמצום |  |
| 5:28.189 | הנעשה |  |
| 5:28.190 | בו |  |
| 5:28.191 | לצדדין |  |
| 5:28.192 | שוה |  |
| 5:28.193 | בהשואה |  |
| 5:28.194 | א | '. |
| 5:28.195 | כדי |  |
| 5:28.196 | להמציא |  |
| 5:28.197 | מקום |  |
| 5:28.198 | שיוכלו |  |
| 5:28.199 | להיות |  |
| 5:28.200 | שם |  |
| 5:28.201 | במקום |  |
| 5:28.202 | הריקן |  |
| 5:28.203 | כל |  |
| 5:28.204 | הנאצלים |  |
| 5:28.205 | והנבראים |  |
| 5:28.206 | והנוצרים |  |
| 5:28.207 | והנעשים | . |
| 5:28.208 | ואז |  |
| 5:28.209 | נמשך |  |
| 5:28.210 | מאור |  |
| 5:28.211 | האין |  |
| 5:28.212 | סוף |  |
| 5:28.213 | קו |  |
| 5:28.214 | דק |  |
| 5:28.215 | מן |  |
| 5:28.216 | האור |  |
| 5:28.217 | העגול |  |
| 5:28.218 | של |  |
| 5:28.219 | האין |  |
| 5:28.220 | סוף |  |
| 5:28.221 | מלמעלה |  |
| 5:28.222 | למטה |  |
| 5:28.223 | ונעשו |  |
| 5:28.224 | ספירות |  |
| 5:28.225 | העיגולים |  |
| 5:28.226 | מאותו |  |
| 5:28.227 | הקו |  |
| 5:28.228 | הנמשך |  |
| 5:28.229 | ממנו |  |
| 5:28.230 | עגול |  |
| 5:28.231 | בתוך |  |
| 5:28.232 | עגול | . |
| 5:28.233 | וגם |  |
| 5:28.234 | נשתלשל |  |
| 5:28.235 | ונמשך |  |
| 5:28.236 | הקו |  |
| 5:28.237 | ההוא |  |
| 5:28.238 | ביושר |  |
| 5:28.239 | תוך |  |
| 5:28.240 | החלל |  |
| 5:28.241 | ונעשו |  |
| 5:28.242 | ספירות |  |
| 5:28.243 | דיושר | . |
| 5:28.244 | אשר |  |
| 5:28.245 | קראום |  |
| 5:28.246 | מדות |  |
| 5:28.247 | וכלים |  |
| 5:28.248 | והם |  |
| 5:28.249 | עיקר |  |
| 5:28.250 | האלהות | , |
| 5:28.251 | וגם |  |
| 5:28.252 | נשמתם |  |
| 5:28.253 | שהוא |  |
| 5:28.254 | האור |  |
| 5:28.255 | הפנימי |  |
| 5:28.256 | גם |  |
| 5:28.257 | מלבושם |  |
| 5:28.258 | שהוא |  |
| 5:28.259 | האור |  |
| 5:28.260 | המקיף |  |
| 5:28.261 | הכל |  |
| 5:28.262 | נמשך |  |
| 5:28.263 | ונשתלשל |  |
| 5:28.264 | מעצם |  |
| 5:28.265 | האין |  |
| 5:28.266 | סוף |  |
| 5:28.267 | המקיף |  |
| 5:28.268 | את |  |
| 5:28.269 | הכל | , |
| 5:28.270 | ממנו |  |
| 5:28.271 | גוף | , |
| 5:28.272 | ממנו |  |
| 5:28.273 | נשמה | , |
| 5:28.274 | ממנו |  |
| 5:28.275 | מלבוש | . |
| 5:28.276 | כלם |  |
| 5:28.277 | נמשכים |  |
| 5:28.278 | מעצמות |  |
| 5:28.279 | האור |  |
| 5:28.280 | ההוא |  |
| 5:28.281 | כמו |  |
| 5:28.282 | שכתבתי |  |
| 5:28.283 | לעיל |  |
| 5:28.284 | סי | ' |
| 5:28.285 | ט"ל |  |
| 5:28.286 | בשם |  |
| 5:28.287 | שושן |  |
| 5:28.288 | סודות |  |
| 5:28.289 | ומשנת |  |
| 5:28.290 | חסידים |  |
| 5:28.291 | שגם |  |
| 5:28.292 | הכלים |  |
| 5:28.293 | אלהות |  |
| 5:28.294 | גמור |  |
| 5:28.295 | כי |  |
| 5:28.296 | מחומר |  |
| 5:28.297 | אחד |  |
| 5:28.298 | קורצו | . |
| 5:28.299 | והסבות |  |
| 5:28.300 | הללו |  |
| 5:28.301 | נמשכים |  |
| 5:28.302 | ונעשה |  |
| 5:28.303 | כל |  |
| 5:28.304 | אחד |  |
| 5:28.305 | סבה |  |
| 5:28.306 | למה |  |
| 5:28.307 | שלמטה |  |
| 5:28.308 | ממנה | . |
| 5:28.309 | זה |  |
| 5:28.310 | סבה |  |
| 5:28.311 | לזה |  |
| 5:28.312 | וזה |  |
| 5:28.313 | סבה |  |
| 5:28.314 | לזה | . |
| 5:28.315 | הסבה |  |
| 5:28.316 | הראשונה |  |
| 5:28.317 | מפרצופי |  |
| 5:28.318 | היושר |  |
| 5:28.319 | ופרצופים |  |
| 5:28.320 | הכוללים |  |
| 5:28.321 | אדם |  |
| 5:28.322 | קדמאה | , |
| 5:28.323 | ואחריו |  |
| 5:28.324 | אדם |  |
| 5:28.325 | קדמון | , |
| 5:28.326 | ואחריו |  |
| 5:28.327 | עתיקא |  |
| 5:28.328 | קדישא | , |
| 5:28.329 | ואחריו |  |
| 5:28.330 | אריך |  |
| 5:28.331 | אנפין | , |
| 5:28.332 | וממנו |  |
| 5:28.333 | נחיו |  |
| 5:28.334 | אבא |  |
| 5:28.335 | ואימא | , |
| 5:28.336 | ומהם |  |
| 5:28.337 | נהיו |  |
| 5:28.338 | זעיר |  |
| 5:28.339 | ונוקביה | , |
| 5:28.340 | והוא |  |
| 5:28.341 | הנקרא |  |
| 5:28.342 | בן |  |
| 5:28.343 | לאבא |  |
| 5:28.344 | ואימא | . |
| 5:28.345 | כל |  |
| 5:28.346 | הפרצופים |  |
| 5:28.347 | הללו |  |
| 5:28.348 | נשתלשלו |  |
| 5:28.349 | זה |  |
| 5:28.350 | מזה | , |
| 5:28.351 | ונקרא |  |
| 5:28.352 | כל |  |
| 5:28.353 | אחד |  |
| 5:28.354 | מהם |  |
| 5:28.355 | עילה |  |
| 5:28.356 | וסבה |  |
| 5:28.357 | לפרצוף |  |
| 5:28.358 | העלול |  |
| 5:28.359 | ממנו |  |
| 5:28.360 | אחריו | . |
| 5:28.361 | ואבא |  |
| 5:28.362 | ואימא |  |
| 5:28.363 | אעטרו |  |
| 5:28.364 | לז"א |  |
| 5:28.365 | בתרין |  |
| 5:28.366 | עטרין |  |
| 5:28.367 | והשליטוהו |  |
| 5:28.368 | על |  |
| 5:28.369 | כל |  |
| 5:28.370 | הברואים |  |
| 5:28.371 | וצוו |  |
| 5:28.372 | לפי |  |
| 5:28.373 | דבריהם | ) |
| 5:28.374 | לעבדו |  |
| 5:28.375 | ליראה |  |
| 5:28.376 | את |  |
| 5:28.377 | אשתו | , |
| 5:28.378 | שריבה |  |
| 5:28.379 | אותה |  |
| 5:28.380 | הכתוב |  |
| 5:28.381 | ליראה |  |
| 5:28.382 | באמרו |  |
| 5:28.383 | את |  |
| 5:28.384 | ה | ' |
| 5:28.385 | אלהיך |  |
| 5:28.386 | תירא | , |
| 5:28.387 | את |  |
| 5:28.388 | לרבות |  |
| 5:28.389 | נוקבא |  |
| 5:28.390 | דיליה |  |
| 5:28.391 | ליראה |  |
| 5:28.392 | אותה |  |
| 5:28.393 | כמו |  |
| 5:28.394 | שנתבאר |  |
| 5:28.395 | לעיל |  |
| 5:28.396 | סי | ' |
| 5:28.397 | כ"ו |  |
| 5:28.398 | כ"ז |  |
| 5:28.399 | כ"ח | , |
| 5:28.400 | מה |  |
| 5:28.401 | שלא |  |
| 5:28.402 | עלה |  |
| 5:28.403 | על |  |
| 5:28.404 | דעת |  |
| 5:28.405 | שמעון |  |
| 5:28.406 | העמסוני |  |
| 5:28.407 | ואמרי |  |
| 5:28.408 | לה |  |
| 5:28.409 | נחמיה |  |
| 5:28.410 | העמסוני |  |
| 5:28.411 | ור | ' |
| 5:28.412 | עקיבה | . |
| 5:28.413 | ויחסו |  |
| 5:28.414 | העבודה |  |
| 5:28.415 | לזעיר |  |
| 5:28.416 | אנפין |  |
| 5:28.417 | דוקא |  |
| 5:28.418 | כמו |  |
| 5:28.419 | שנ | ' |
| 5:28.420 | ואותו |  |
| 5:28.421 | תעבוד |  |
| 5:28.422 | והוא |  |
| 5:28.423 | נקרא |  |
| 5:28.424 | בפיהם |  |
| 5:28.425 | יהוה |  |
| 5:28.426 | אלהינו | . |
| 5:28.427 | והוא |  |
| 5:28.428 | הסבה |  |
| 5:28.429 | והפרצוף |  |
| 5:28.430 | האחרון |  |
| 5:28.431 | שבפרצופים |  |
| 5:28.432 | האצילות | . |
| 5:28.433 | ואמרו |  |
| 5:28.434 | שהוא |  |
| 5:28.435 | הוא |  |
| 5:28.436 | אלהינו |  |
| 5:28.437 | ואנחנו |  |
| 5:28.438 | עמו |  |
| 5:28.439 | ועבדיו |  |
| 5:28.440 | כמבואר |  |
| 5:28.441 | בכסא |  |
| 5:28.442 | אליהו |  |
| 5:28.443 | וס | ' |
| 5:28.444 | הברית |  |
| 5:28.445 | ויושר |  |
| 5:28.446 | לבב |  |
| 5:28.447 | דף |  |
| 5:28.448 | כ"א |  |
| 5:28.449 | ול"א |  |
| 5:28.450 | ול"ב | . |
| 5:28.451 | וענין |  |
| 5:28.452 | השתלשלותם |  |
| 5:28.453 | בעץ |  |
| 5:28.454 | חיים |  |
| 5:28.455 | שער |  |
| 5:28.456 | עיגולים |  |
| 5:28.457 | ויושר |  |
| 5:28.458 | וכסא |  |
| 5:28.459 | אלהיו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

