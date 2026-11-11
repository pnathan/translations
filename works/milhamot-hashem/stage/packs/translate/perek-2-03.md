# Translation pack: milhamot-hashem perek-2-03 (Milchamot ha-Shem 2:8–2:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-03`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
- ריש מלין אמר echoing Dan 7:1: by the verse's sense ("he stated the sum of the
  matter"); the verse in `apparatus`.
- A variant spelling of a word the text does not gloss (ברותא for ברייתא):
  transliterate it ("berota"), with a note.
- An Aramaic word found once and unclear (דחלות): the best reading by its
  grammar, the alternatives in `readings`.
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
- דברי תימה: "astonishing words" (words that astonish).
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 2:8

Hebrew (Wikisource): תשובה מהשואל. איותה נפשי לשמוע אמרות מוזקקות כראי מוצקות, משמי מרום בוערות. כי אני מכל דעה אביון, שואף לצל הקיקיון. ותחת אשר ענותך תרבני וממים רבים כמר מר לי תשקני. עשיתני כמלונה במקשה, ונפשי אנושה בתהלות מרמה (כברכת בלעם הרשע) דחלות. ובדברי נרגן סיימת ולאיש פרוע מוסר ודוב שכול דמית. ובאמת לא זוהי דרכה של תורה. תורתנו הקדושה דרכיה דרכי נועם. ואהבת לרעך כמוך. מה דעלך סני לחברך לא תעביד. כן אמר הלל נשיא ישראל! ריש מלין אמר, מה שאני פערתי פי לבזות חכמי ישראל והוצאתי עליהם מלין אשר לא כדת. חלילה לי לבזות שום אדם אפילו גוי, כל שכן ישראל. ולא ידעתי מקום בדברי לזה שאתה טוען שיש בהם בזיון. ואם להקשות עליהם מדברי רז"ל ולגלות מקום הסתירה, אין בזה שום איסור. וגם בתלמוד מקשים אפילו על גדולים מהם ומביאים ראיה לסתור דבריהם ואומרים תיובתא דר' פלוני. הא דר' פלוני בדותא היא (גירסת הערוך ברותא) ובסנהדרין ובברכות אמרינן: הרואה את הקשת צריך לכרוע דכתיב כמראה הקשת אשר יהיה בענן וכו' ואראה ואפול על פני. לייטי עלה במערבא וכו' מגדף בה ר' אבהו וכו'. והנה במכתבך זה תרעומת מצאתי, תשובה לא מצאתי. וכי אני הוא שקראתים לאלו הפרצופים אלוהות? הלא המקובלים הם עצמם שאמרו כן. וצא למד מכוונת הברכות והתפלה מי הוא אלהינו ומי הוא אלהי אבותינו. ובזהר שואל מאן אלהים דא כך וכך. ועוד שהרי המקובלים ביארו שכמה עולמות אלפי אלפים ורבי רבבות נאצלו ונבראו ונוצרו ונעשו, רק שלא שלחו בהם יד לבארם לרוב העלמם. כי הם אורות דקים ביותר. וכל עולם ועולם כלול מעשר ספירות. וכל ספירה וספירה כלולה כל אחת מעשר פרטי פרטות כמו שאמרנו בעולם דאצילות. וכולן בבחינת עיגולים, זה בתוך זה וזה לפנים מזה. וכן בבחינת ספירות דיושר שהם בציור אדם בעל קומה זקופה ברמ"ח איברים וכו' כל אלו העולמות כל א' מהם כולל שתי בחינות עיגולים ויושר בכלל ובפרט. וכל דברי המקובלים וביאוריהם, אינם רק בעולם דאצילות שכבר נתעבה ונתגלה יותר, ובו ביארו כל ענייני הפרצופים הכוללים שהם חמשה: אריך אנפין, ואבא ואימא, וזעיר ונוקביה. ואחר שכבר נתברר לך מדבריהם שאומרים שנתעבו ונתגלו שהם קורים את הפרצופים הנז' אלוהות באמרם אלהינו ואלהי אבותינו, בתפלה ובברכות, ומכוונים למדה ידועה בכל תיבה. לא נשאר עלי תלונתך לא בטענת גשמות ולא טענת אלוהות, כי אם עליהם, כי הם אמרו והם אמרו. וקושיא חזרה למקומה, איך נניח האין סוף עם כל הספירות אשר בעולמות העליונים עם אריך אנפין ואבא ואימא, ונעבוד לז"א ונאמר שהוא אלהינו ואנחנו עמו ועבדיו הוא עשנו ולו אנחנו. ובמכתבך זה (שאמרת לפי שידוע שהמידות הפנימיות מהם דין ומהם רחמים וכו') כבר הודית הודאה גמורה שכן היא דעתכם שאין שום עבודה ולא תפילה שייכא בעלת על כל העלות, כי אם לז"א, וכמו שכ' הרב מקדש מלך וכס"א ומהר"י לופיס בשער האמונה האמתית. ולכן צריכים אנחנו ממך ליישב ולתרץ לנו איך תסכים דעת המקובלים הלזו שכל עבודתינו ותפלותינו לז"א ונשמתו, עם תורתינו הקדושה שבכתב ושבעל פה. ואשר חרה אפך עלי ואמרת שאני מגשים בספירות, גם זה לא ידעתי לו מקום בדברי, לפי שלא אמרתי דבר שלא מצאתי בספריהם. והם אמרו שעולם האצילות יש בו גוף ונשמה ומלבוש. ולא נעלם ממך כי האור והאויר והאש גופים דקים הם וקלים. ועוד שהרי ביארו דבריהם הם עצמם ואמרו כי העולמות העליונים והפרצופים שלמעלה מאריך אנפין לא דברו בהם, רק מא"א ולמטה שכבר נתעבו ונתגלו יותר. הרי לך בהדיא מדבריהם שהספירות שמדברים בהם נתעבו ונתגלו יותר משל מעלה מהם ומודים הם בעצמם שהספירות שמדברים בהם נתעבו.

| id | word | punctuation after |
|---|---|---|
| 2:8.1 | תשובה |  |
| 2:8.2 | מהשואל | . |
| 2:8.3 | איותה |  |
| 2:8.4 | נפשי |  |
| 2:8.5 | לשמוע |  |
| 2:8.6 | אמרות |  |
| 2:8.7 | מוזקקות |  |
| 2:8.8 | כראי |  |
| 2:8.9 | מוצקות | , |
| 2:8.10 | משמי |  |
| 2:8.11 | מרום |  |
| 2:8.12 | בוערות | . |
| 2:8.13 | כי |  |
| 2:8.14 | אני |  |
| 2:8.15 | מכל |  |
| 2:8.16 | דעה |  |
| 2:8.17 | אביון | , |
| 2:8.18 | שואף |  |
| 2:8.19 | לצל |  |
| 2:8.20 | הקיקיון | . |
| 2:8.21 | ותחת |  |
| 2:8.22 | אשר |  |
| 2:8.23 | ענותך |  |
| 2:8.24 | תרבני |  |
| 2:8.25 | וממים |  |
| 2:8.26 | רבים |  |
| 2:8.27 | כמר |  |
| 2:8.28 | מר |  |
| 2:8.29 | לי |  |
| 2:8.30 | תשקני | . |
| 2:8.31 | עשיתני |  |
| 2:8.32 | כמלונה |  |
| 2:8.33 | במקשה | , |
| 2:8.34 | ונפשי |  |
| 2:8.35 | אנושה |  |
| 2:8.36 | בתהלות |  |
| 2:8.37 | מרמה |  |
| 2:8.38 | כברכת |  |
| 2:8.39 | בלעם |  |
| 2:8.40 | הרשע | ) |
| 2:8.41 | דחלות | . |
| 2:8.42 | ובדברי |  |
| 2:8.43 | נרגן |  |
| 2:8.44 | סיימת |  |
| 2:8.45 | ולאיש |  |
| 2:8.46 | פרוע |  |
| 2:8.47 | מוסר |  |
| 2:8.48 | ודוב |  |
| 2:8.49 | שכול |  |
| 2:8.50 | דמית | . |
| 2:8.51 | ובאמת |  |
| 2:8.52 | לא |  |
| 2:8.53 | זוהי |  |
| 2:8.54 | דרכה |  |
| 2:8.55 | של |  |
| 2:8.56 | תורה | . |
| 2:8.57 | תורתנו |  |
| 2:8.58 | הקדושה |  |
| 2:8.59 | דרכיה |  |
| 2:8.60 | דרכי |  |
| 2:8.61 | נועם | . |
| 2:8.62 | ואהבת |  |
| 2:8.63 | לרעך |  |
| 2:8.64 | כמוך | . |
| 2:8.65 | מה |  |
| 2:8.66 | דעלך |  |
| 2:8.67 | סני |  |
| 2:8.68 | לחברך |  |
| 2:8.69 | לא |  |
| 2:8.70 | תעביד | . |
| 2:8.71 | כן |  |
| 2:8.72 | אמר |  |
| 2:8.73 | הלל |  |
| 2:8.74 | נשיא |  |
| 2:8.75 | ישראל | ! |
| 2:8.76 | ריש |  |
| 2:8.77 | מלין |  |
| 2:8.78 | אמר | , |
| 2:8.79 | מה |  |
| 2:8.80 | שאני |  |
| 2:8.81 | פערתי |  |
| 2:8.82 | פי |  |
| 2:8.83 | לבזות |  |
| 2:8.84 | חכמי |  |
| 2:8.85 | ישראל |  |
| 2:8.86 | והוצאתי |  |
| 2:8.87 | עליהם |  |
| 2:8.88 | מלין |  |
| 2:8.89 | אשר |  |
| 2:8.90 | לא |  |
| 2:8.91 | כדת | . |
| 2:8.92 | חלילה |  |
| 2:8.93 | לי |  |
| 2:8.94 | לבזות |  |
| 2:8.95 | שום |  |
| 2:8.96 | אדם |  |
| 2:8.97 | אפילו |  |
| 2:8.98 | גוי | , |
| 2:8.99 | כל |  |
| 2:8.100 | שכן |  |
| 2:8.101 | ישראל | . |
| 2:8.102 | ולא |  |
| 2:8.103 | ידעתי |  |
| 2:8.104 | מקום |  |
| 2:8.105 | בדברי |  |
| 2:8.106 | לזה |  |
| 2:8.107 | שאתה |  |
| 2:8.108 | טוען |  |
| 2:8.109 | שיש |  |
| 2:8.110 | בהם |  |
| 2:8.111 | בזיון | . |
| 2:8.112 | ואם |  |
| 2:8.113 | להקשות |  |
| 2:8.114 | עליהם |  |
| 2:8.115 | מדברי |  |
| 2:8.116 | רז"ל |  |
| 2:8.117 | ולגלות |  |
| 2:8.118 | מקום |  |
| 2:8.119 | הסתירה | , |
| 2:8.120 | אין |  |
| 2:8.121 | בזה |  |
| 2:8.122 | שום |  |
| 2:8.123 | איסור | . |
| 2:8.124 | וגם |  |
| 2:8.125 | בתלמוד |  |
| 2:8.126 | מקשים |  |
| 2:8.127 | אפילו |  |
| 2:8.128 | על |  |
| 2:8.129 | גדולים |  |
| 2:8.130 | מהם |  |
| 2:8.131 | ומביאים |  |
| 2:8.132 | ראיה |  |
| 2:8.133 | לסתור |  |
| 2:8.134 | דבריהם |  |
| 2:8.135 | ואומרים |  |
| 2:8.136 | תיובתא |  |
| 2:8.137 | דר | ' |
| 2:8.138 | פלוני | . |
| 2:8.139 | הא |  |
| 2:8.140 | דר | ' |
| 2:8.141 | פלוני |  |
| 2:8.142 | בדותא |  |
| 2:8.143 | היא |  |
| 2:8.144 | גירסת |  |
| 2:8.145 | הערוך |  |
| 2:8.146 | ברותא | ) |
| 2:8.147 | ובסנהדרין |  |
| 2:8.148 | ובברכות |  |
| 2:8.149 | אמרינן | : |
| 2:8.150 | הרואה |  |
| 2:8.151 | את |  |
| 2:8.152 | הקשת |  |
| 2:8.153 | צריך |  |
| 2:8.154 | לכרוע |  |
| 2:8.155 | דכתיב |  |
| 2:8.156 | כמראה |  |
| 2:8.157 | הקשת |  |
| 2:8.158 | אשר |  |
| 2:8.159 | יהיה |  |
| 2:8.160 | בענן |  |
| 2:8.161 | וכו | ' |
| 2:8.162 | ואראה |  |
| 2:8.163 | ואפול |  |
| 2:8.164 | על |  |
| 2:8.165 | פני | . |
| 2:8.166 | לייטי |  |
| 2:8.167 | עלה |  |
| 2:8.168 | במערבא |  |
| 2:8.169 | וכו | ' |
| 2:8.170 | מגדף |  |
| 2:8.171 | בה |  |
| 2:8.172 | ר | ' |
| 2:8.173 | אבהו |  |
| 2:8.174 | וכו | '. |
| 2:8.175 | והנה |  |
| 2:8.176 | במכתבך |  |
| 2:8.177 | זה |  |
| 2:8.178 | תרעומת |  |
| 2:8.179 | מצאתי | , |
| 2:8.180 | תשובה |  |
| 2:8.181 | לא |  |
| 2:8.182 | מצאתי | . |
| 2:8.183 | וכי |  |
| 2:8.184 | אני |  |
| 2:8.185 | הוא |  |
| 2:8.186 | שקראתים |  |
| 2:8.187 | לאלו |  |
| 2:8.188 | הפרצופים |  |
| 2:8.189 | אלוהות | ? |
| 2:8.190 | הלא |  |
| 2:8.191 | המקובלים |  |
| 2:8.192 | הם |  |
| 2:8.193 | עצמם |  |
| 2:8.194 | שאמרו |  |
| 2:8.195 | כן | . |
| 2:8.196 | וצא |  |
| 2:8.197 | למד |  |
| 2:8.198 | מכוונת |  |
| 2:8.199 | הברכות |  |
| 2:8.200 | והתפלה |  |
| 2:8.201 | מי |  |
| 2:8.202 | הוא |  |
| 2:8.203 | אלהינו |  |
| 2:8.204 | ומי |  |
| 2:8.205 | הוא |  |
| 2:8.206 | אלהי |  |
| 2:8.207 | אבותינו | . |
| 2:8.208 | ובזהר |  |
| 2:8.209 | שואל |  |
| 2:8.210 | מאן |  |
| 2:8.211 | אלהים |  |
| 2:8.212 | דא |  |
| 2:8.213 | כך |  |
| 2:8.214 | וכך | . |
| 2:8.215 | ועוד |  |
| 2:8.216 | שהרי |  |
| 2:8.217 | המקובלים |  |
| 2:8.218 | ביארו |  |
| 2:8.219 | שכמה |  |
| 2:8.220 | עולמות |  |
| 2:8.221 | אלפי |  |
| 2:8.222 | אלפים |  |
| 2:8.223 | ורבי |  |
| 2:8.224 | רבבות |  |
| 2:8.225 | נאצלו |  |
| 2:8.226 | ונבראו |  |
| 2:8.227 | ונוצרו |  |
| 2:8.228 | ונעשו | , |
| 2:8.229 | רק |  |
| 2:8.230 | שלא |  |
| 2:8.231 | שלחו |  |
| 2:8.232 | בהם |  |
| 2:8.233 | יד |  |
| 2:8.234 | לבארם |  |
| 2:8.235 | לרוב |  |
| 2:8.236 | העלמם | . |
| 2:8.237 | כי |  |
| 2:8.238 | הם |  |
| 2:8.239 | אורות |  |
| 2:8.240 | דקים |  |
| 2:8.241 | ביותר | . |
| 2:8.242 | וכל |  |
| 2:8.243 | עולם |  |
| 2:8.244 | ועולם |  |
| 2:8.245 | כלול |  |
| 2:8.246 | מעשר |  |
| 2:8.247 | ספירות | . |
| 2:8.248 | וכל |  |
| 2:8.249 | ספירה |  |
| 2:8.250 | וספירה |  |
| 2:8.251 | כלולה |  |
| 2:8.252 | כל |  |
| 2:8.253 | אחת |  |
| 2:8.254 | מעשר |  |
| 2:8.255 | פרטי |  |
| 2:8.256 | פרטות |  |
| 2:8.257 | כמו |  |
| 2:8.258 | שאמרנו |  |
| 2:8.259 | בעולם |  |
| 2:8.260 | דאצילות | . |
| 2:8.261 | וכולן |  |
| 2:8.262 | בבחינת |  |
| 2:8.263 | עיגולים | , |
| 2:8.264 | זה |  |
| 2:8.265 | בתוך |  |
| 2:8.266 | זה |  |
| 2:8.267 | וזה |  |
| 2:8.268 | לפנים |  |
| 2:8.269 | מזה | . |
| 2:8.270 | וכן |  |
| 2:8.271 | בבחינת |  |
| 2:8.272 | ספירות |  |
| 2:8.273 | דיושר |  |
| 2:8.274 | שהם |  |
| 2:8.275 | בציור |  |
| 2:8.276 | אדם |  |
| 2:8.277 | בעל |  |
| 2:8.278 | קומה |  |
| 2:8.279 | זקופה |  |
| 2:8.280 | ברמ"ח |  |
| 2:8.281 | איברים |  |
| 2:8.282 | וכו | ' |
| 2:8.283 | כל |  |
| 2:8.284 | אלו |  |
| 2:8.285 | העולמות |  |
| 2:8.286 | כל |  |
| 2:8.287 | א | ' |
| 2:8.288 | מהם |  |
| 2:8.289 | כולל |  |
| 2:8.290 | שתי |  |
| 2:8.291 | בחינות |  |
| 2:8.292 | עיגולים |  |
| 2:8.293 | ויושר |  |
| 2:8.294 | בכלל |  |
| 2:8.295 | ובפרט | . |
| 2:8.296 | וכל |  |
| 2:8.297 | דברי |  |
| 2:8.298 | המקובלים |  |
| 2:8.299 | וביאוריהם | , |
| 2:8.300 | אינם |  |
| 2:8.301 | רק |  |
| 2:8.302 | בעולם |  |
| 2:8.303 | דאצילות |  |
| 2:8.304 | שכבר |  |
| 2:8.305 | נתעבה |  |
| 2:8.306 | ונתגלה |  |
| 2:8.307 | יותר | , |
| 2:8.308 | ובו |  |
| 2:8.309 | ביארו |  |
| 2:8.310 | כל |  |
| 2:8.311 | ענייני |  |
| 2:8.312 | הפרצופים |  |
| 2:8.313 | הכוללים |  |
| 2:8.314 | שהם |  |
| 2:8.315 | חמשה | : |
| 2:8.316 | אריך |  |
| 2:8.317 | אנפין | , |
| 2:8.318 | ואבא |  |
| 2:8.319 | ואימא | , |
| 2:8.320 | וזעיר |  |
| 2:8.321 | ונוקביה | . |
| 2:8.322 | ואחר |  |
| 2:8.323 | שכבר |  |
| 2:8.324 | נתברר |  |
| 2:8.325 | לך |  |
| 2:8.326 | מדבריהם |  |
| 2:8.327 | שאומרים |  |
| 2:8.328 | שנתעבו |  |
| 2:8.329 | ונתגלו |  |
| 2:8.330 | שהם |  |
| 2:8.331 | קורים |  |
| 2:8.332 | את |  |
| 2:8.333 | הפרצופים |  |
| 2:8.334 | הנז | ' |
| 2:8.335 | אלוהות |  |
| 2:8.336 | באמרם |  |
| 2:8.337 | אלהינו |  |
| 2:8.338 | ואלהי |  |
| 2:8.339 | אבותינו | , |
| 2:8.340 | בתפלה |  |
| 2:8.341 | ובברכות | , |
| 2:8.342 | ומכוונים |  |
| 2:8.343 | למדה |  |
| 2:8.344 | ידועה |  |
| 2:8.345 | בכל |  |
| 2:8.346 | תיבה | . |
| 2:8.347 | לא |  |
| 2:8.348 | נשאר |  |
| 2:8.349 | עלי |  |
| 2:8.350 | תלונתך |  |
| 2:8.351 | לא |  |
| 2:8.352 | בטענת |  |
| 2:8.353 | גשמות |  |
| 2:8.354 | ולא |  |
| 2:8.355 | טענת |  |
| 2:8.356 | אלוהות | , |
| 2:8.357 | כי |  |
| 2:8.358 | אם |  |
| 2:8.359 | עליהם | , |
| 2:8.360 | כי |  |
| 2:8.361 | הם |  |
| 2:8.362 | אמרו |  |
| 2:8.363 | והם |  |
| 2:8.364 | אמרו | . |
| 2:8.365 | וקושיא |  |
| 2:8.366 | חזרה |  |
| 2:8.367 | למקומה | , |
| 2:8.368 | איך |  |
| 2:8.369 | נניח |  |
| 2:8.370 | האין |  |
| 2:8.371 | סוף |  |
| 2:8.372 | עם |  |
| 2:8.373 | כל |  |
| 2:8.374 | הספירות |  |
| 2:8.375 | אשר |  |
| 2:8.376 | בעולמות |  |
| 2:8.377 | העליונים |  |
| 2:8.378 | עם |  |
| 2:8.379 | אריך |  |
| 2:8.380 | אנפין |  |
| 2:8.381 | ואבא |  |
| 2:8.382 | ואימא | , |
| 2:8.383 | ונעבוד |  |
| 2:8.384 | לז"א |  |
| 2:8.385 | ונאמר |  |
| 2:8.386 | שהוא |  |
| 2:8.387 | אלהינו |  |
| 2:8.388 | ואנחנו |  |
| 2:8.389 | עמו |  |
| 2:8.390 | ועבדיו |  |
| 2:8.391 | הוא |  |
| 2:8.392 | עשנו |  |
| 2:8.393 | ולו |  |
| 2:8.394 | אנחנו | . |
| 2:8.395 | ובמכתבך |  |
| 2:8.396 | זה |  |
| 2:8.397 | שאמרת |  |
| 2:8.398 | לפי |  |
| 2:8.399 | שידוע |  |
| 2:8.400 | שהמידות |  |
| 2:8.401 | הפנימיות |  |
| 2:8.402 | מהם |  |
| 2:8.403 | דין |  |
| 2:8.404 | ומהם |  |
| 2:8.405 | רחמים |  |
| 2:8.406 | וכו | ') |
| 2:8.407 | כבר |  |
| 2:8.408 | הודית |  |
| 2:8.409 | הודאה |  |
| 2:8.410 | גמורה |  |
| 2:8.411 | שכן |  |
| 2:8.412 | היא |  |
| 2:8.413 | דעתכם |  |
| 2:8.414 | שאין |  |
| 2:8.415 | שום |  |
| 2:8.416 | עבודה |  |
| 2:8.417 | ולא |  |
| 2:8.418 | תפילה |  |
| 2:8.419 | שייכא |  |
| 2:8.420 | בעלת |  |
| 2:8.421 | על |  |
| 2:8.422 | כל |  |
| 2:8.423 | העלות | , |
| 2:8.424 | כי |  |
| 2:8.425 | אם |  |
| 2:8.426 | לז"א | , |
| 2:8.427 | וכמו |  |
| 2:8.428 | שכ | ' |
| 2:8.429 | הרב |  |
| 2:8.430 | מקדש |  |
| 2:8.431 | מלך |  |
| 2:8.432 | וכס"א |  |
| 2:8.433 | ומהר"י |  |
| 2:8.434 | לופיס |  |
| 2:8.435 | בשער |  |
| 2:8.436 | האמונה |  |
| 2:8.437 | האמתית | . |
| 2:8.438 | ולכן |  |
| 2:8.439 | צריכים |  |
| 2:8.440 | אנחנו |  |
| 2:8.441 | ממך |  |
| 2:8.442 | ליישב |  |
| 2:8.443 | ולתרץ |  |
| 2:8.444 | לנו |  |
| 2:8.445 | איך |  |
| 2:8.446 | תסכים |  |
| 2:8.447 | דעת |  |
| 2:8.448 | המקובלים |  |
| 2:8.449 | הלזו |  |
| 2:8.450 | שכל |  |
| 2:8.451 | עבודתינו |  |
| 2:8.452 | ותפלותינו |  |
| 2:8.453 | לז"א |  |
| 2:8.454 | ונשמתו | , |
| 2:8.455 | עם |  |
| 2:8.456 | תורתינו |  |
| 2:8.457 | הקדושה |  |
| 2:8.458 | שבכתב |  |
| 2:8.459 | ושבעל |  |
| 2:8.460 | פה | . |
| 2:8.461 | ואשר |  |
| 2:8.462 | חרה |  |
| 2:8.463 | אפך |  |
| 2:8.464 | עלי |  |
| 2:8.465 | ואמרת |  |
| 2:8.466 | שאני |  |
| 2:8.467 | מגשים |  |
| 2:8.468 | בספירות | , |
| 2:8.469 | גם |  |
| 2:8.470 | זה |  |
| 2:8.471 | לא |  |
| 2:8.472 | ידעתי |  |
| 2:8.473 | לו |  |
| 2:8.474 | מקום |  |
| 2:8.475 | בדברי | , |
| 2:8.476 | לפי |  |
| 2:8.477 | שלא |  |
| 2:8.478 | אמרתי |  |
| 2:8.479 | דבר |  |
| 2:8.480 | שלא |  |
| 2:8.481 | מצאתי |  |
| 2:8.482 | בספריהם | . |
| 2:8.483 | והם |  |
| 2:8.484 | אמרו |  |
| 2:8.485 | שעולם |  |
| 2:8.486 | האצילות |  |
| 2:8.487 | יש |  |
| 2:8.488 | בו |  |
| 2:8.489 | גוף |  |
| 2:8.490 | ונשמה |  |
| 2:8.491 | ומלבוש | . |
| 2:8.492 | ולא |  |
| 2:8.493 | נעלם |  |
| 2:8.494 | ממך |  |
| 2:8.495 | כי |  |
| 2:8.496 | האור |  |
| 2:8.497 | והאויר |  |
| 2:8.498 | והאש |  |
| 2:8.499 | גופים |  |
| 2:8.500 | דקים |  |
| 2:8.501 | הם |  |
| 2:8.502 | וקלים | . |
| 2:8.503 | ועוד |  |
| 2:8.504 | שהרי |  |
| 2:8.505 | ביארו |  |
| 2:8.506 | דבריהם |  |
| 2:8.507 | הם |  |
| 2:8.508 | עצמם |  |
| 2:8.509 | ואמרו |  |
| 2:8.510 | כי |  |
| 2:8.511 | העולמות |  |
| 2:8.512 | העליונים |  |
| 2:8.513 | והפרצופים |  |
| 2:8.514 | שלמעלה |  |
| 2:8.515 | מאריך |  |
| 2:8.516 | אנפין |  |
| 2:8.517 | לא |  |
| 2:8.518 | דברו |  |
| 2:8.519 | בהם | , |
| 2:8.520 | רק |  |
| 2:8.521 | מא"א |  |
| 2:8.522 | ולמטה |  |
| 2:8.523 | שכבר |  |
| 2:8.524 | נתעבו |  |
| 2:8.525 | ונתגלו |  |
| 2:8.526 | יותר | . |
| 2:8.527 | הרי |  |
| 2:8.528 | לך |  |
| 2:8.529 | בהדיא |  |
| 2:8.530 | מדבריהם |  |
| 2:8.531 | שהספירות |  |
| 2:8.532 | שמדברים |  |
| 2:8.533 | בהם |  |
| 2:8.534 | נתעבו |  |
| 2:8.535 | ונתגלו |  |
| 2:8.536 | יותר |  |
| 2:8.537 | משל |  |
| 2:8.538 | מעלה |  |
| 2:8.539 | מהם |  |
| 2:8.540 | ומודים |  |
| 2:8.541 | הם |  |
| 2:8.542 | בעצמם |  |
| 2:8.543 | שהספירות |  |
| 2:8.544 | שמדברים |  |
| 2:8.545 | בהם |  |
| 2:8.546 | נתעבו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

