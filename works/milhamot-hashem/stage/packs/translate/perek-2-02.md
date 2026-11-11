# Translation pack: milhamot-hashem perek-2-02 (Milchamot ha-Shem 2:5–2:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-02`.

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
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- דמיון: "analogy" for an argument, "likeness" for resemblance.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 2:5

Hebrew (Wikisource): ואיך נניח כל האלוהות אשר למעלה מעולם דאצילות ונעבוד דוקא לזעיר ונוקביה שבעולם האצילות שכבר קדמו לו כמה עתיקין ואריכי אנפין ואו"א וזעיר ונוקביה בעולמות העליונים אין מספר הקרובים לאין סוף יותר?

| id | word | punctuation after |
|---|---|---|
| 2:5.1 | ואיך |  |
| 2:5.2 | נניח |  |
| 2:5.3 | כל |  |
| 2:5.4 | האלוהות |  |
| 2:5.5 | אשר |  |
| 2:5.6 | למעלה |  |
| 2:5.7 | מעולם |  |
| 2:5.8 | דאצילות |  |
| 2:5.9 | ונעבוד |  |
| 2:5.10 | דוקא |  |
| 2:5.11 | לזעיר |  |
| 2:5.12 | ונוקביה |  |
| 2:5.13 | שבעולם |  |
| 2:5.14 | האצילות |  |
| 2:5.15 | שכבר |  |
| 2:5.16 | קדמו |  |
| 2:5.17 | לו |  |
| 2:5.18 | כמה |  |
| 2:5.19 | עתיקין |  |
| 2:5.20 | ואריכי |  |
| 2:5.21 | אנפין |  |
| 2:5.22 | ואו"א |  |
| 2:5.23 | וזעיר |  |
| 2:5.24 | ונוקביה |  |
| 2:5.25 | בעולמות |  |
| 2:5.26 | העליונים |  |
| 2:5.27 | אין |  |
| 2:5.28 | מספר |  |
| 2:5.29 | הקרובים |  |
| 2:5.30 | לאין |  |
| 2:5.31 | סוף |  |
| 2:5.32 | יותר | ? |


## Milchamot ha-Shem 2:6

Hebrew (Wikisource): ורבותינו הראשונים אמרו שאין ראוי לעבוד ולהודות ולגדל ולפאר רק לאלוה הראשון לכל הנמצאים מבלעדיו! ולמה נניח האין סוף העיקרי המצטמצם למעלה מכולם עם כל הפרצופים הקרובים אליו, ונעבוד לזה הרחוק ממנו מאוד שהוא פרצוף אחרון שבעולם האצילות? את כל אלה הביננו אדונינו נלכה נרדפה את ה' יוצרינו יוצר בראשית. ואל תצוק בנו דבר בליעל ח"ו. כי מי אסף חן ושכל טוב בחפניו השכל וידוע את ה' בחקירת העניינים האלהיים בעיקרי האמונה כמוך. מי יעלה בהר ה' ומי יקום במקום קדשו כמוך, כי בער אנכי מאיש ולא בינת אדם לי. וצור יודע תעלומות לב, בוחן לבבות כי לא נתתי שינה לעיני ונגזלה שנתי מעיני עד אמצא פתרונים והבא לטהר מסייעין אותו, להמציא לו מורה דרך הקדש בשכל ומדע ככלכל ודרדע, ובין חכמי האמת חכם ואומן, כאיתן והימן, שאולי הרבנים החסרים מן החכמה הזאת לנו משאות שוא ומדוחים ח"ו. ואתה אדני לך נגלו תעלומותיה וצפונך תמלא בטנם של תלמידיך, ישבע ויתענגו מטובך, ועלינו לשלם שכר פעולתך. ואתה אל תמנע טוב מבעליו. והנני משביעך באלהי השמים האר פניך אלינו לעות את יעף דבר דבור על אפניו והיינו עבדים לאדוני. ושכר רב טובך תבענה שפתינו, ואז לא נבוש לא בעה"ז ולא בעה"ב ונשובה נרעה בשושנים מסוד חכמים ונבונים.

| id | word | punctuation after |
|---|---|---|
| 2:6.1 | ורבותינו |  |
| 2:6.2 | הראשונים |  |
| 2:6.3 | אמרו |  |
| 2:6.4 | שאין |  |
| 2:6.5 | ראוי |  |
| 2:6.6 | לעבוד |  |
| 2:6.7 | ולהודות |  |
| 2:6.8 | ולגדל |  |
| 2:6.9 | ולפאר |  |
| 2:6.10 | רק |  |
| 2:6.11 | לאלוה |  |
| 2:6.12 | הראשון |  |
| 2:6.13 | לכל |  |
| 2:6.14 | הנמצאים |  |
| 2:6.15 | מבלעדיו | ! |
| 2:6.16 | ולמה |  |
| 2:6.17 | נניח |  |
| 2:6.18 | האין |  |
| 2:6.19 | סוף |  |
| 2:6.20 | העיקרי |  |
| 2:6.21 | המצטמצם |  |
| 2:6.22 | למעלה |  |
| 2:6.23 | מכולם |  |
| 2:6.24 | עם |  |
| 2:6.25 | כל |  |
| 2:6.26 | הפרצופים |  |
| 2:6.27 | הקרובים |  |
| 2:6.28 | אליו | , |
| 2:6.29 | ונעבוד |  |
| 2:6.30 | לזה |  |
| 2:6.31 | הרחוק |  |
| 2:6.32 | ממנו |  |
| 2:6.33 | מאוד |  |
| 2:6.34 | שהוא |  |
| 2:6.35 | פרצוף |  |
| 2:6.36 | אחרון |  |
| 2:6.37 | שבעולם |  |
| 2:6.38 | האצילות | ? |
| 2:6.39 | את |  |
| 2:6.40 | כל |  |
| 2:6.41 | אלה |  |
| 2:6.42 | הביננו |  |
| 2:6.43 | אדונינו |  |
| 2:6.44 | נלכה |  |
| 2:6.45 | נרדפה |  |
| 2:6.46 | את |  |
| 2:6.47 | ה | ' |
| 2:6.48 | יוצרינו |  |
| 2:6.49 | יוצר |  |
| 2:6.50 | בראשית | . |
| 2:6.51 | ואל |  |
| 2:6.52 | תצוק |  |
| 2:6.53 | בנו |  |
| 2:6.54 | דבר |  |
| 2:6.55 | בליעל |  |
| 2:6.56 | ח"ו | . |
| 2:6.57 | כי |  |
| 2:6.58 | מי |  |
| 2:6.59 | אסף |  |
| 2:6.60 | חן |  |
| 2:6.61 | ושכל |  |
| 2:6.62 | טוב |  |
| 2:6.63 | בחפניו |  |
| 2:6.64 | השכל |  |
| 2:6.65 | וידוע |  |
| 2:6.66 | את |  |
| 2:6.67 | ה | ' |
| 2:6.68 | בחקירת |  |
| 2:6.69 | העניינים |  |
| 2:6.70 | האלהיים |  |
| 2:6.71 | בעיקרי |  |
| 2:6.72 | האמונה |  |
| 2:6.73 | כמוך | . |
| 2:6.74 | מי |  |
| 2:6.75 | יעלה |  |
| 2:6.76 | בהר |  |
| 2:6.77 | ה | ' |
| 2:6.78 | ומי |  |
| 2:6.79 | יקום |  |
| 2:6.80 | במקום |  |
| 2:6.81 | קדשו |  |
| 2:6.82 | כמוך | , |
| 2:6.83 | כי |  |
| 2:6.84 | בער |  |
| 2:6.85 | אנכי |  |
| 2:6.86 | מאיש |  |
| 2:6.87 | ולא |  |
| 2:6.88 | בינת |  |
| 2:6.89 | אדם |  |
| 2:6.90 | לי | . |
| 2:6.91 | וצור |  |
| 2:6.92 | יודע |  |
| 2:6.93 | תעלומות |  |
| 2:6.94 | לב | , |
| 2:6.95 | בוחן |  |
| 2:6.96 | לבבות |  |
| 2:6.97 | כי |  |
| 2:6.98 | לא |  |
| 2:6.99 | נתתי |  |
| 2:6.100 | שינה |  |
| 2:6.101 | לעיני |  |
| 2:6.102 | ונגזלה |  |
| 2:6.103 | שנתי |  |
| 2:6.104 | מעיני |  |
| 2:6.105 | עד |  |
| 2:6.106 | אמצא |  |
| 2:6.107 | פתרונים |  |
| 2:6.108 | והבא |  |
| 2:6.109 | לטהר |  |
| 2:6.110 | מסייעין |  |
| 2:6.111 | אותו | , |
| 2:6.112 | להמציא |  |
| 2:6.113 | לו |  |
| 2:6.114 | מורה |  |
| 2:6.115 | דרך |  |
| 2:6.116 | הקדש |  |
| 2:6.117 | בשכל |  |
| 2:6.118 | ומדע |  |
| 2:6.119 | ככלכל |  |
| 2:6.120 | ודרדע | , |
| 2:6.121 | ובין |  |
| 2:6.122 | חכמי |  |
| 2:6.123 | האמת |  |
| 2:6.124 | חכם |  |
| 2:6.125 | ואומן | , |
| 2:6.126 | כאיתן |  |
| 2:6.127 | והימן | , |
| 2:6.128 | שאולי |  |
| 2:6.129 | הרבנים |  |
| 2:6.130 | החסרים |  |
| 2:6.131 | מן |  |
| 2:6.132 | החכמה |  |
| 2:6.133 | הזאת |  |
| 2:6.134 | לנו |  |
| 2:6.135 | משאות |  |
| 2:6.136 | שוא |  |
| 2:6.137 | ומדוחים |  |
| 2:6.138 | ח"ו | . |
| 2:6.139 | ואתה |  |
| 2:6.140 | אדני |  |
| 2:6.141 | לך |  |
| 2:6.142 | נגלו |  |
| 2:6.143 | תעלומותיה |  |
| 2:6.144 | וצפונך |  |
| 2:6.145 | תמלא |  |
| 2:6.146 | בטנם |  |
| 2:6.147 | של |  |
| 2:6.148 | תלמידיך | , |
| 2:6.149 | ישבע |  |
| 2:6.150 | ויתענגו |  |
| 2:6.151 | מטובך | , |
| 2:6.152 | ועלינו |  |
| 2:6.153 | לשלם |  |
| 2:6.154 | שכר |  |
| 2:6.155 | פעולתך | . |
| 2:6.156 | ואתה |  |
| 2:6.157 | אל |  |
| 2:6.158 | תמנע |  |
| 2:6.159 | טוב |  |
| 2:6.160 | מבעליו | . |
| 2:6.161 | והנני |  |
| 2:6.162 | משביעך |  |
| 2:6.163 | באלהי |  |
| 2:6.164 | השמים |  |
| 2:6.165 | האר |  |
| 2:6.166 | פניך |  |
| 2:6.167 | אלינו |  |
| 2:6.168 | לעות |  |
| 2:6.169 | את |  |
| 2:6.170 | יעף |  |
| 2:6.171 | דבר |  |
| 2:6.172 | דבור |  |
| 2:6.173 | על |  |
| 2:6.174 | אפניו |  |
| 2:6.175 | והיינו |  |
| 2:6.176 | עבדים |  |
| 2:6.177 | לאדוני | . |
| 2:6.178 | ושכר |  |
| 2:6.179 | רב |  |
| 2:6.180 | טובך |  |
| 2:6.181 | תבענה |  |
| 2:6.182 | שפתינו | , |
| 2:6.183 | ואז |  |
| 2:6.184 | לא |  |
| 2:6.185 | נבוש |  |
| 2:6.186 | לא |  |
| 2:6.187 | בעה"ז |  |
| 2:6.188 | ולא |  |
| 2:6.189 | בעה"ב |  |
| 2:6.190 | ונשובה |  |
| 2:6.191 | נרעה |  |
| 2:6.192 | בשושנים |  |
| 2:6.193 | מסוד |  |
| 2:6.194 | חכמים |  |
| 2:6.195 | ונבונים | . |


## Milchamot ha-Shem 2:7

Hebrew (Wikisource): תשובה שנית ראיתי את כל הרשום בכתב האלוף, אילן שענפיו מרובים, רומו כגובה ארזים, אחר המחילה על השהייה כי טרדת הזמן נטויה, וכל עכבה לטובה. בהשקיפי על המכתב והנה פערת פיך בהזכירך אלוהות בחשבך שהוא כנגדינו. ונחנו מה כי תלינו עלינו ובזיונך אינו מגיע אלינו כי אם אל חכמי ישראל אשר קטנם עבה ממתניך וכל ישראל נחים אליהם ובזכותם העולם עומד. אי עליך כי הוצאת עליהם דברים אשר לא כדת, בתומה ותחתיה ערמה. אבל כבר נתבאר דינך בפי כל בעלי חכמה. ומה שכתבת על כסא אליהו כי בדף ג' מדמה אחדותו ית' לשאר אחדים. כבר אתה הולך חשכים, הלא תחילת דבריו דף עשרים שכל דמיונות אלו על אורות קדושים אשר אין להם תפיסה וצורה ודמיון, לא בשכל ולא בדעת כלל ועיקר, אלא כדי שיוכל השכל לקבל אותם הורשינו לדמות. ובדף ד' ע"ב כתב שכל הכנויים והתארים הם בדרך דמיון בספירות ואודות קדושים הנאצלים, אבל בו ית' אין בו שום דמיון וצורה ותפישה במחשבה כלל ועיקר. ומה שכתבת איך נניח שאר הפרצופים ונעבוד לזעיר אנפין עם כח שבו א"ס ב"ה. כבר ביאר הר' כס"א בדף כ"ה ע"ב שבקראינו לז"א אנחנו קורין לכל הפרצופים כאחד עם נשמה המסתתרת בהם. ומה שכתבת למה נניח הפרצופים הקרובים לא"ס ונקרא לז"א. לפי שידוע שהמידות הפנימיות יש מהם דין ויש מהם רחמים, וזה כלול מכולם, ולזה רצה המאציל לפעול פעולותיו על ידי כלי זה. ועד שאתה מקשה על דמיונות אודות עליונים שאין לך רשות להרהר בהן במכל שכן ללגלג, תקשה על ידך הגופנית שהיא עושה כל המלאכות, והראש שלמעלה אינו עושה כלום. ג"כ כח הדבור שהוא פארך על כל בעל חי למה לא ניתן למעלה מעיניך. אלא כך גזרה חכמתו ית'. ואם יש לך עצה הוגנת תלך אצל בוראך ותיעצנו. ומה שכתבת שדברי כס"א נגד הרמב"ם וכו' שכתבו שאינו כח בגוף, ולפי דברי כס"א הוא כח בגוף. הלא כל המקובלים פה אחד עונים ואומרים שכל המגשים בספירות כמוך הרי זה כופר ואין לו חלק באלהי ישראל. ומה שכינו אותם אורות דקים אינו אלא לשכך את האזן! יש לנו לשאול על מה שכתב הרמב"ם: יכוין לבו כנגד השכינה ויתפלל. ובמורה כ' שהשכינה אור נברא. ויותר מזה מה שכתב על כבוד ה' שהוא נברא. אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחווים לנברא. ויותר על זה מה שכתב שכל צורה שהיו רואים הנביאים ומדברים עמה היא אור נברא. א"כ אין לנו שום נבואה ע"י הבורא, אלא כולם מצד הנבראים. ואם נאמר שהיה הנברא דבק בבורא, א"כ לדעתך הוא כח בגוף! ותשובתך מהרה תצמח, ס' מאנו דעת אותי נאם ה'.

| id | word | punctuation after |
|---|---|---|
| 2:7.1 | תשובה |  |
| 2:7.2 | שנית |  |
| 2:7.3 | ראיתי |  |
| 2:7.4 | את |  |
| 2:7.5 | כל |  |
| 2:7.6 | הרשום |  |
| 2:7.7 | בכתב |  |
| 2:7.8 | האלוף | , |
| 2:7.9 | אילן |  |
| 2:7.10 | שענפיו |  |
| 2:7.11 | מרובים | , |
| 2:7.12 | רומו |  |
| 2:7.13 | כגובה |  |
| 2:7.14 | ארזים | , |
| 2:7.15 | אחר |  |
| 2:7.16 | המחילה |  |
| 2:7.17 | על |  |
| 2:7.18 | השהייה |  |
| 2:7.19 | כי |  |
| 2:7.20 | טרדת |  |
| 2:7.21 | הזמן |  |
| 2:7.22 | נטויה | , |
| 2:7.23 | וכל |  |
| 2:7.24 | עכבה |  |
| 2:7.25 | לטובה | . |
| 2:7.26 | בהשקיפי |  |
| 2:7.27 | על |  |
| 2:7.28 | המכתב |  |
| 2:7.29 | והנה |  |
| 2:7.30 | פערת |  |
| 2:7.31 | פיך |  |
| 2:7.32 | בהזכירך |  |
| 2:7.33 | אלוהות |  |
| 2:7.34 | בחשבך |  |
| 2:7.35 | שהוא |  |
| 2:7.36 | כנגדינו | . |
| 2:7.37 | ונחנו |  |
| 2:7.38 | מה |  |
| 2:7.39 | כי |  |
| 2:7.40 | תלינו |  |
| 2:7.41 | עלינו |  |
| 2:7.42 | ובזיונך |  |
| 2:7.43 | אינו |  |
| 2:7.44 | מגיע |  |
| 2:7.45 | אלינו |  |
| 2:7.46 | כי |  |
| 2:7.47 | אם |  |
| 2:7.48 | אל |  |
| 2:7.49 | חכמי |  |
| 2:7.50 | ישראל |  |
| 2:7.51 | אשר |  |
| 2:7.52 | קטנם |  |
| 2:7.53 | עבה |  |
| 2:7.54 | ממתניך |  |
| 2:7.55 | וכל |  |
| 2:7.56 | ישראל |  |
| 2:7.57 | נחים |  |
| 2:7.58 | אליהם |  |
| 2:7.59 | ובזכותם |  |
| 2:7.60 | העולם |  |
| 2:7.61 | עומד | . |
| 2:7.62 | אי |  |
| 2:7.63 | עליך |  |
| 2:7.64 | כי |  |
| 2:7.65 | הוצאת |  |
| 2:7.66 | עליהם |  |
| 2:7.67 | דברים |  |
| 2:7.68 | אשר |  |
| 2:7.69 | לא |  |
| 2:7.70 | כדת | , |
| 2:7.71 | בתומה |  |
| 2:7.72 | ותחתיה |  |
| 2:7.73 | ערמה | . |
| 2:7.74 | אבל |  |
| 2:7.75 | כבר |  |
| 2:7.76 | נתבאר |  |
| 2:7.77 | דינך |  |
| 2:7.78 | בפי |  |
| 2:7.79 | כל |  |
| 2:7.80 | בעלי |  |
| 2:7.81 | חכמה | . |
| 2:7.82 | ומה |  |
| 2:7.83 | שכתבת |  |
| 2:7.84 | על |  |
| 2:7.85 | כסא |  |
| 2:7.86 | אליהו |  |
| 2:7.87 | כי |  |
| 2:7.88 | בדף |  |
| 2:7.89 | ג | ' |
| 2:7.90 | מדמה |  |
| 2:7.91 | אחדותו |  |
| 2:7.92 | ית | ' |
| 2:7.93 | לשאר |  |
| 2:7.94 | אחדים | . |
| 2:7.95 | כבר |  |
| 2:7.96 | אתה |  |
| 2:7.97 | הולך |  |
| 2:7.98 | חשכים | , |
| 2:7.99 | הלא |  |
| 2:7.100 | תחילת |  |
| 2:7.101 | דבריו |  |
| 2:7.102 | דף |  |
| 2:7.103 | עשרים |  |
| 2:7.104 | שכל |  |
| 2:7.105 | דמיונות |  |
| 2:7.106 | אלו |  |
| 2:7.107 | על |  |
| 2:7.108 | אורות |  |
| 2:7.109 | קדושים |  |
| 2:7.110 | אשר |  |
| 2:7.111 | אין |  |
| 2:7.112 | להם |  |
| 2:7.113 | תפיסה |  |
| 2:7.114 | וצורה |  |
| 2:7.115 | ודמיון | , |
| 2:7.116 | לא |  |
| 2:7.117 | בשכל |  |
| 2:7.118 | ולא |  |
| 2:7.119 | בדעת |  |
| 2:7.120 | כלל |  |
| 2:7.121 | ועיקר | , |
| 2:7.122 | אלא |  |
| 2:7.123 | כדי |  |
| 2:7.124 | שיוכל |  |
| 2:7.125 | השכל |  |
| 2:7.126 | לקבל |  |
| 2:7.127 | אותם |  |
| 2:7.128 | הורשינו |  |
| 2:7.129 | לדמות | . |
| 2:7.130 | ובדף |  |
| 2:7.131 | ד | ' |
| 2:7.132 | ע"ב |  |
| 2:7.133 | כתב |  |
| 2:7.134 | שכל |  |
| 2:7.135 | הכנויים |  |
| 2:7.136 | והתארים |  |
| 2:7.137 | הם |  |
| 2:7.138 | בדרך |  |
| 2:7.139 | דמיון |  |
| 2:7.140 | בספירות |  |
| 2:7.141 | ואודות |  |
| 2:7.142 | קדושים |  |
| 2:7.143 | הנאצלים | , |
| 2:7.144 | אבל |  |
| 2:7.145 | בו |  |
| 2:7.146 | ית | ' |
| 2:7.147 | אין |  |
| 2:7.148 | בו |  |
| 2:7.149 | שום |  |
| 2:7.150 | דמיון |  |
| 2:7.151 | וצורה |  |
| 2:7.152 | ותפישה |  |
| 2:7.153 | במחשבה |  |
| 2:7.154 | כלל |  |
| 2:7.155 | ועיקר | . |
| 2:7.156 | ומה |  |
| 2:7.157 | שכתבת |  |
| 2:7.158 | איך |  |
| 2:7.159 | נניח |  |
| 2:7.160 | שאר |  |
| 2:7.161 | הפרצופים |  |
| 2:7.162 | ונעבוד |  |
| 2:7.163 | לזעיר |  |
| 2:7.164 | אנפין |  |
| 2:7.165 | עם |  |
| 2:7.166 | כח |  |
| 2:7.167 | שבו |  |
| 2:7.168 | א"ס |  |
| 2:7.169 | ב"ה | . |
| 2:7.170 | כבר |  |
| 2:7.171 | ביאר |  |
| 2:7.172 | הר | ' |
| 2:7.173 | כס"א |  |
| 2:7.174 | בדף |  |
| 2:7.175 | כ"ה |  |
| 2:7.176 | ע"ב |  |
| 2:7.177 | שבקראינו |  |
| 2:7.178 | לז"א |  |
| 2:7.179 | אנחנו |  |
| 2:7.180 | קורין |  |
| 2:7.181 | לכל |  |
| 2:7.182 | הפרצופים |  |
| 2:7.183 | כאחד |  |
| 2:7.184 | עם |  |
| 2:7.185 | נשמה |  |
| 2:7.186 | המסתתרת |  |
| 2:7.187 | בהם | . |
| 2:7.188 | ומה |  |
| 2:7.189 | שכתבת |  |
| 2:7.190 | למה |  |
| 2:7.191 | נניח |  |
| 2:7.192 | הפרצופים |  |
| 2:7.193 | הקרובים |  |
| 2:7.194 | לא"ס |  |
| 2:7.195 | ונקרא |  |
| 2:7.196 | לז"א | . |
| 2:7.197 | לפי |  |
| 2:7.198 | שידוע |  |
| 2:7.199 | שהמידות |  |
| 2:7.200 | הפנימיות |  |
| 2:7.201 | יש |  |
| 2:7.202 | מהם |  |
| 2:7.203 | דין |  |
| 2:7.204 | ויש |  |
| 2:7.205 | מהם |  |
| 2:7.206 | רחמים | , |
| 2:7.207 | וזה |  |
| 2:7.208 | כלול |  |
| 2:7.209 | מכולם | , |
| 2:7.210 | ולזה |  |
| 2:7.211 | רצה |  |
| 2:7.212 | המאציל |  |
| 2:7.213 | לפעול |  |
| 2:7.214 | פעולותיו |  |
| 2:7.215 | על |  |
| 2:7.216 | ידי |  |
| 2:7.217 | כלי |  |
| 2:7.218 | זה | . |
| 2:7.219 | ועד |  |
| 2:7.220 | שאתה |  |
| 2:7.221 | מקשה |  |
| 2:7.222 | על |  |
| 2:7.223 | דמיונות |  |
| 2:7.224 | אודות |  |
| 2:7.225 | עליונים |  |
| 2:7.226 | שאין |  |
| 2:7.227 | לך |  |
| 2:7.228 | רשות |  |
| 2:7.229 | להרהר |  |
| 2:7.230 | בהן |  |
| 2:7.231 | במכל |  |
| 2:7.232 | שכן |  |
| 2:7.233 | ללגלג | , |
| 2:7.234 | תקשה |  |
| 2:7.235 | על |  |
| 2:7.236 | ידך |  |
| 2:7.237 | הגופנית |  |
| 2:7.238 | שהיא |  |
| 2:7.239 | עושה |  |
| 2:7.240 | כל |  |
| 2:7.241 | המלאכות | , |
| 2:7.242 | והראש |  |
| 2:7.243 | שלמעלה |  |
| 2:7.244 | אינו |  |
| 2:7.245 | עושה |  |
| 2:7.246 | כלום | . |
| 2:7.247 | ג"כ |  |
| 2:7.248 | כח |  |
| 2:7.249 | הדבור |  |
| 2:7.250 | שהוא |  |
| 2:7.251 | פארך |  |
| 2:7.252 | על |  |
| 2:7.253 | כל |  |
| 2:7.254 | בעל |  |
| 2:7.255 | חי |  |
| 2:7.256 | למה |  |
| 2:7.257 | לא |  |
| 2:7.258 | ניתן |  |
| 2:7.259 | למעלה |  |
| 2:7.260 | מעיניך | . |
| 2:7.261 | אלא |  |
| 2:7.262 | כך |  |
| 2:7.263 | גזרה |  |
| 2:7.264 | חכמתו |  |
| 2:7.265 | ית | '. |
| 2:7.266 | ואם |  |
| 2:7.267 | יש |  |
| 2:7.268 | לך |  |
| 2:7.269 | עצה |  |
| 2:7.270 | הוגנת |  |
| 2:7.271 | תלך |  |
| 2:7.272 | אצל |  |
| 2:7.273 | בוראך |  |
| 2:7.274 | ותיעצנו | . |
| 2:7.275 | ומה |  |
| 2:7.276 | שכתבת |  |
| 2:7.277 | שדברי |  |
| 2:7.278 | כס"א |  |
| 2:7.279 | נגד |  |
| 2:7.280 | הרמב"ם |  |
| 2:7.281 | וכו | ' |
| 2:7.282 | שכתבו |  |
| 2:7.283 | שאינו |  |
| 2:7.284 | כח |  |
| 2:7.285 | בגוף | , |
| 2:7.286 | ולפי |  |
| 2:7.287 | דברי |  |
| 2:7.288 | כס"א |  |
| 2:7.289 | הוא |  |
| 2:7.290 | כח |  |
| 2:7.291 | בגוף | . |
| 2:7.292 | הלא |  |
| 2:7.293 | כל |  |
| 2:7.294 | המקובלים |  |
| 2:7.295 | פה |  |
| 2:7.296 | אחד |  |
| 2:7.297 | עונים |  |
| 2:7.298 | ואומרים |  |
| 2:7.299 | שכל |  |
| 2:7.300 | המגשים |  |
| 2:7.301 | בספירות |  |
| 2:7.302 | כמוך |  |
| 2:7.303 | הרי |  |
| 2:7.304 | זה |  |
| 2:7.305 | כופר |  |
| 2:7.306 | ואין |  |
| 2:7.307 | לו |  |
| 2:7.308 | חלק |  |
| 2:7.309 | באלהי |  |
| 2:7.310 | ישראל | . |
| 2:7.311 | ומה |  |
| 2:7.312 | שכינו |  |
| 2:7.313 | אותם |  |
| 2:7.314 | אורות |  |
| 2:7.315 | דקים |  |
| 2:7.316 | אינו |  |
| 2:7.317 | אלא |  |
| 2:7.318 | לשכך |  |
| 2:7.319 | את |  |
| 2:7.320 | האזן | ! |
| 2:7.321 | יש |  |
| 2:7.322 | לנו |  |
| 2:7.323 | לשאול |  |
| 2:7.324 | על |  |
| 2:7.325 | מה |  |
| 2:7.326 | שכתב |  |
| 2:7.327 | הרמב"ם | : |
| 2:7.328 | יכוין |  |
| 2:7.329 | לבו |  |
| 2:7.330 | כנגד |  |
| 2:7.331 | השכינה |  |
| 2:7.332 | ויתפלל | . |
| 2:7.333 | ובמורה |  |
| 2:7.334 | כ | ' |
| 2:7.335 | שהשכינה |  |
| 2:7.336 | אור |  |
| 2:7.337 | נברא | . |
| 2:7.338 | ויותר |  |
| 2:7.339 | מזה |  |
| 2:7.340 | מה |  |
| 2:7.341 | שכתב |  |
| 2:7.342 | על |  |
| 2:7.343 | כבוד |  |
| 2:7.344 | ה | ' |
| 2:7.345 | שהוא |  |
| 2:7.346 | נברא | . |
| 2:7.347 | אם |  |
| 2:7.348 | כן |  |
| 2:7.349 | כיצד |  |
| 2:7.350 | היו |  |
| 2:7.351 | ישראל |  |
| 2:7.352 | יוצאים |  |
| 2:7.353 | פתח |  |
| 2:7.354 | אהליהם |  |
| 2:7.355 | ומשתחווים |  |
| 2:7.356 | לנברא | . |
| 2:7.357 | ויותר |  |
| 2:7.358 | על |  |
| 2:7.359 | זה |  |
| 2:7.360 | מה |  |
| 2:7.361 | שכתב |  |
| 2:7.362 | שכל |  |
| 2:7.363 | צורה |  |
| 2:7.364 | שהיו |  |
| 2:7.365 | רואים |  |
| 2:7.366 | הנביאים |  |
| 2:7.367 | ומדברים |  |
| 2:7.368 | עמה |  |
| 2:7.369 | היא |  |
| 2:7.370 | אור |  |
| 2:7.371 | נברא | . |
| 2:7.372 | א"כ |  |
| 2:7.373 | אין |  |
| 2:7.374 | לנו |  |
| 2:7.375 | שום |  |
| 2:7.376 | נבואה |  |
| 2:7.377 | ע"י |  |
| 2:7.378 | הבורא | , |
| 2:7.379 | אלא |  |
| 2:7.380 | כולם |  |
| 2:7.381 | מצד |  |
| 2:7.382 | הנבראים | . |
| 2:7.383 | ואם |  |
| 2:7.384 | נאמר |  |
| 2:7.385 | שהיה |  |
| 2:7.386 | הנברא |  |
| 2:7.387 | דבק |  |
| 2:7.388 | בבורא | , |
| 2:7.389 | א"כ |  |
| 2:7.390 | לדעתך |  |
| 2:7.391 | הוא |  |
| 2:7.392 | כח |  |
| 2:7.393 | בגוף | ! |
| 2:7.394 | ותשובתך |  |
| 2:7.395 | מהרה |  |
| 2:7.396 | תצמח | , |
| 2:7.397 | ס | ' |
| 2:7.398 | מאנו |  |
| 2:7.399 | דעת |  |
| 2:7.400 | אותי |  |
| 2:7.401 | נאם |  |
| 2:7.402 | ה | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

