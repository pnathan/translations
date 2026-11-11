# Translation pack: milhamot-hashem perek-2-05 (Milchamot ha-Shem 2:10–2:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-05`.

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
- ה׳ (the Tetragrammaton as written): "the LORD", in the author's prose and
  in verses. Where it is cited as a name ("זעיר אנפין דאיקרי ה'", 4:11), write
  "who is called 'the LORD'". Not ה׳ as Tetragrammaton: "מאת ה' אליעזר בן
  יהודה" (1:2) → "from Mr. Eliezer Ben-Yehuda"; ה׳ as a numeral (5) or a page
  letter → the number. The pack splits it into the token ה plus the
  punctuation ', so read the token in context.
- הקב"ה / קב"ה, and קודשא בריך הוא in the Aramaic → "the Holy One, blessed be
  He". As a name under discussion: "the name 'the Holy One, blessed be He'"
  ("שם הקב"ה המורגל בפינו", 4:17).
- השי"ת → "the Name, may He be blessed". ית׳ / יתברך → "may He be blessed",
  placed where the Hebrew places it ("אחדותו ית'" → "His oneness, may He be
  blessed"). ב"ה → "blessed be He". השם הנכבד → "the glorious Name". שם שמים
  → "the name of Heaven".
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
- מ"ו = "our teacher" (מורנו) in the formula יודיענו מורנו; ב"ו = "blessed be
  He" only before הוא elsewhere ב"ו = בשר ודם ("flesh and blood").
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- ר"ל with a named authority as subject: "means to say".
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
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".

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


## Milchamot ha-Shem 2:10

Hebrew (Wikisource): ובירושלמי שבת פ' במה אשה אמרו חז"ל וריויה די רביעאה דמי לבר אלהין. אמר ר' ראובן: באותה שעה ירד מלאך וסטרו לההוא רשיעא על פיו ואמר ליה תקין מילך ובר אית ליה? חזר ואמר ברוך אלההון די שדרך מישך ועביד נגו די שלה מלאכיה ושיזיב לעבדוהי, בריה אין כתיב כאן אלא מלאכיה, ע"כ. ור"ל שנבוכדנאצר הרשע חשב שהמלאכים הם עלולים ומשתלשלים ממנו בטבע כבן מן האב כדברי המקובלים בספירות של כל העולמות, וכן הנשמות, ולזה סטרו המלאך. וחזר ואמר די שלח מלאכיה וכו'. הרי לך מפורש יוצא מדברי מדרש רבה והירושלמי היפך דברי המקובלים. ושכל הנמצאים המציאם הקב"ה בחסדו יש מאין מהצורה ראשונה עד יתוש קטן. ולא שהם נמשכים מעצמותו כמו שהבן מתהוה מכח הזרע הנמשך מאביו.

| id | word | punctuation after |
|---|---|---|
| 2:10.1 | ובירושלמי |  |
| 2:10.2 | שבת |  |
| 2:10.3 | פ | ' |
| 2:10.4 | במה |  |
| 2:10.5 | אשה |  |
| 2:10.6 | אמרו |  |
| 2:10.7 | חז"ל |  |
| 2:10.8 | וריויה |  |
| 2:10.9 | די |  |
| 2:10.10 | רביעאה |  |
| 2:10.11 | דמי |  |
| 2:10.12 | לבר |  |
| 2:10.13 | אלהין | . |
| 2:10.14 | אמר |  |
| 2:10.15 | ר | ' |
| 2:10.16 | ראובן | : |
| 2:10.17 | באותה |  |
| 2:10.18 | שעה |  |
| 2:10.19 | ירד |  |
| 2:10.20 | מלאך |  |
| 2:10.21 | וסטרו |  |
| 2:10.22 | לההוא |  |
| 2:10.23 | רשיעא |  |
| 2:10.24 | על |  |
| 2:10.25 | פיו |  |
| 2:10.26 | ואמר |  |
| 2:10.27 | ליה |  |
| 2:10.28 | תקין |  |
| 2:10.29 | מילך |  |
| 2:10.30 | ובר |  |
| 2:10.31 | אית |  |
| 2:10.32 | ליה | ? |
| 2:10.33 | חזר |  |
| 2:10.34 | ואמר |  |
| 2:10.35 | ברוך |  |
| 2:10.36 | אלההון |  |
| 2:10.37 | די |  |
| 2:10.38 | שדרך |  |
| 2:10.39 | מישך |  |
| 2:10.40 | ועביד |  |
| 2:10.41 | נגו |  |
| 2:10.42 | די |  |
| 2:10.43 | שלה |  |
| 2:10.44 | מלאכיה |  |
| 2:10.45 | ושיזיב |  |
| 2:10.46 | לעבדוהי | , |
| 2:10.47 | בריה |  |
| 2:10.48 | אין |  |
| 2:10.49 | כתיב |  |
| 2:10.50 | כאן |  |
| 2:10.51 | אלא |  |
| 2:10.52 | מלאכיה | , |
| 2:10.53 | ע"כ | . |
| 2:10.54 | ור"ל |  |
| 2:10.55 | שנבוכדנאצר |  |
| 2:10.56 | הרשע |  |
| 2:10.57 | חשב |  |
| 2:10.58 | שהמלאכים |  |
| 2:10.59 | הם |  |
| 2:10.60 | עלולים |  |
| 2:10.61 | ומשתלשלים |  |
| 2:10.62 | ממנו |  |
| 2:10.63 | בטבע |  |
| 2:10.64 | כבן |  |
| 2:10.65 | מן |  |
| 2:10.66 | האב |  |
| 2:10.67 | כדברי |  |
| 2:10.68 | המקובלים |  |
| 2:10.69 | בספירות |  |
| 2:10.70 | של |  |
| 2:10.71 | כל |  |
| 2:10.72 | העולמות | , |
| 2:10.73 | וכן |  |
| 2:10.74 | הנשמות | , |
| 2:10.75 | ולזה |  |
| 2:10.76 | סטרו |  |
| 2:10.77 | המלאך | . |
| 2:10.78 | וחזר |  |
| 2:10.79 | ואמר |  |
| 2:10.80 | די |  |
| 2:10.81 | שלח |  |
| 2:10.82 | מלאכיה |  |
| 2:10.83 | וכו | '. |
| 2:10.84 | הרי |  |
| 2:10.85 | לך |  |
| 2:10.86 | מפורש |  |
| 2:10.87 | יוצא |  |
| 2:10.88 | מדברי |  |
| 2:10.89 | מדרש |  |
| 2:10.90 | רבה |  |
| 2:10.91 | והירושלמי |  |
| 2:10.92 | היפך |  |
| 2:10.93 | דברי |  |
| 2:10.94 | המקובלים | . |
| 2:10.95 | ושכל |  |
| 2:10.96 | הנמצאים |  |
| 2:10.97 | המציאם |  |
| 2:10.98 | הקב"ה |  |
| 2:10.99 | בחסדו |  |
| 2:10.100 | יש |  |
| 2:10.101 | מאין |  |
| 2:10.102 | מהצורה |  |
| 2:10.103 | ראשונה |  |
| 2:10.104 | עד |  |
| 2:10.105 | יתוש |  |
| 2:10.106 | קטן | . |
| 2:10.107 | ולא |  |
| 2:10.108 | שהם |  |
| 2:10.109 | נמשכים |  |
| 2:10.110 | מעצמותו |  |
| 2:10.111 | כמו |  |
| 2:10.112 | שהבן |  |
| 2:10.113 | מתהוה |  |
| 2:10.114 | מכח |  |
| 2:10.115 | הזרע |  |
| 2:10.116 | הנמשך |  |
| 2:10.117 | מאביו | . |


## Milchamot ha-Shem 2:11

Hebrew (Wikisource): המורם מכל זה שאין לעבוד שום פרצוף וצורה מכל הנבראים הנמצאים בעולם העליון והשפל שהמציאם השי"ת בחסדו. כי לו לבדו ראוי לעבוד ולהתפלל, ועל זה באה המצוה והאזהרה בסיני: אנכי ה' אלהיך לא יהיה לך אלהים אחרים, כמו שכ' הרמב"ם ז"ל בפ' שני מהל' ע"ז. ובפ' ג' מהל' תשובה כתב וז"ל: ואלו שאין להם חלק לעולם הבא אלא נכרתין ואובדין על גודל רשעם וחטאתם לעולם ולעולמי עולמים: המינין והאפיקורסין וכו'. ובדין שאחריו כ': חמשה הן הנקראין מינים. האומר שאין שם אלוה ואין לעולם מנהיג, והאומר שיש שם מנהיג אבל הם שנים או יותר, והאומר שיש שם רבון אחד אלא שהוא גוף ובעל תמונה, והאומר שאינו לבדו הראשון וצור לכל. וכן העובד לזולתו כדי להיות מליץ בינו ובין רבון העולמים. כל אחד מחמשה אלו מין ע"כ. וכתב הרב לחם משנה על חלוקה ד' שכתב רבינו והאומר שאינו לבדו הראשון וכו', כ' הראב"ד, כאותו שאמר אלהיכם צייר גדול הוא אלא שמצא לו סמנים גדולים תהו ובהו וחשך ומים ורוח ובהם עשה מה שעשה, עכ"ל הראב"ד. ופי' הרב לח"מ שכוונתו להוסיף על דברי רבי' שאעפ"י שמודה שלא קדמה שום סבה לאלהינו שבשבילה נהיה, אלא שאומר שברא העולם יש מיש, הרי זה מין ע"ש. ובנדון דידן אמונת הקבלה החדשה כפי המבואר בספריהם לפי הבנתינו והבנת רבים וגדולים חכמי לב מרבותינו הפורשים ממנה ראשונים ואחרונים כהריב"ש וחו"י המובא בפתחי תשובה ושאר אחרונים נקבצו באו בה תלת וארבע מלי מהני דאמור רבנן, דהמאמין כן אין לו חלק לעה"ב. חדא רבוי האלוהות, א"ס ואדם קדמאה ואדם קדמון ועתיק ואריך אנפין ואבא ואימא וזעיר ונוקביה! שנית, שהם בעלי גוף דק, דהיינו אור, והא"ס נשמה לגופות הללו. שלישית, שהעבודה אינה לסבה הראשונה שהוא נקרא אצלם א"ס, רק לז"א שהוא סבה אחרונה שבאלו הסבות. רביעית, שהוא אמצעי הממשיך השפע מאלו המדות העליונות שהם עתיק ואריך אנפין ואבא ואימא, וקורין אותו בן לאו"א, והוא ג"כ אב לכתר דבריאה כפי השתלשלות העולמות לפי דעתם רחמנא ליצלן.

| id | word | punctuation after |
|---|---|---|
| 2:11.1 | המורם |  |
| 2:11.2 | מכל |  |
| 2:11.3 | זה |  |
| 2:11.4 | שאין |  |
| 2:11.5 | לעבוד |  |
| 2:11.6 | שום |  |
| 2:11.7 | פרצוף |  |
| 2:11.8 | וצורה |  |
| 2:11.9 | מכל |  |
| 2:11.10 | הנבראים |  |
| 2:11.11 | הנמצאים |  |
| 2:11.12 | בעולם |  |
| 2:11.13 | העליון |  |
| 2:11.14 | והשפל |  |
| 2:11.15 | שהמציאם |  |
| 2:11.16 | השי"ת |  |
| 2:11.17 | בחסדו | . |
| 2:11.18 | כי |  |
| 2:11.19 | לו |  |
| 2:11.20 | לבדו |  |
| 2:11.21 | ראוי |  |
| 2:11.22 | לעבוד |  |
| 2:11.23 | ולהתפלל | , |
| 2:11.24 | ועל |  |
| 2:11.25 | זה |  |
| 2:11.26 | באה |  |
| 2:11.27 | המצוה |  |
| 2:11.28 | והאזהרה |  |
| 2:11.29 | בסיני | : |
| 2:11.30 | אנכי |  |
| 2:11.31 | ה | ' |
| 2:11.32 | אלהיך |  |
| 2:11.33 | לא |  |
| 2:11.34 | יהיה |  |
| 2:11.35 | לך |  |
| 2:11.36 | אלהים |  |
| 2:11.37 | אחרים | , |
| 2:11.38 | כמו |  |
| 2:11.39 | שכ | ' |
| 2:11.40 | הרמב"ם |  |
| 2:11.41 | ז"ל |  |
| 2:11.42 | בפ | ' |
| 2:11.43 | שני |  |
| 2:11.44 | מהל | ' |
| 2:11.45 | ע"ז | . |
| 2:11.46 | ובפ | ' |
| 2:11.47 | ג | ' |
| 2:11.48 | מהל | ' |
| 2:11.49 | תשובה |  |
| 2:11.50 | כתב |  |
| 2:11.51 | וז"ל | : |
| 2:11.52 | ואלו |  |
| 2:11.53 | שאין |  |
| 2:11.54 | להם |  |
| 2:11.55 | חלק |  |
| 2:11.56 | לעולם |  |
| 2:11.57 | הבא |  |
| 2:11.58 | אלא |  |
| 2:11.59 | נכרתין |  |
| 2:11.60 | ואובדין |  |
| 2:11.61 | על |  |
| 2:11.62 | גודל |  |
| 2:11.63 | רשעם |  |
| 2:11.64 | וחטאתם |  |
| 2:11.65 | לעולם |  |
| 2:11.66 | ולעולמי |  |
| 2:11.67 | עולמים | : |
| 2:11.68 | המינין |  |
| 2:11.69 | והאפיקורסין |  |
| 2:11.70 | וכו | '. |
| 2:11.71 | ובדין |  |
| 2:11.72 | שאחריו |  |
| 2:11.73 | כ | ': |
| 2:11.74 | חמשה |  |
| 2:11.75 | הן |  |
| 2:11.76 | הנקראין |  |
| 2:11.77 | מינים | . |
| 2:11.78 | האומר |  |
| 2:11.79 | שאין |  |
| 2:11.80 | שם |  |
| 2:11.81 | אלוה |  |
| 2:11.82 | ואין |  |
| 2:11.83 | לעולם |  |
| 2:11.84 | מנהיג | , |
| 2:11.85 | והאומר |  |
| 2:11.86 | שיש |  |
| 2:11.87 | שם |  |
| 2:11.88 | מנהיג |  |
| 2:11.89 | אבל |  |
| 2:11.90 | הם |  |
| 2:11.91 | שנים |  |
| 2:11.92 | או |  |
| 2:11.93 | יותר | , |
| 2:11.94 | והאומר |  |
| 2:11.95 | שיש |  |
| 2:11.96 | שם |  |
| 2:11.97 | רבון |  |
| 2:11.98 | אחד |  |
| 2:11.99 | אלא |  |
| 2:11.100 | שהוא |  |
| 2:11.101 | גוף |  |
| 2:11.102 | ובעל |  |
| 2:11.103 | תמונה | , |
| 2:11.104 | והאומר |  |
| 2:11.105 | שאינו |  |
| 2:11.106 | לבדו |  |
| 2:11.107 | הראשון |  |
| 2:11.108 | וצור |  |
| 2:11.109 | לכל | . |
| 2:11.110 | וכן |  |
| 2:11.111 | העובד |  |
| 2:11.112 | לזולתו |  |
| 2:11.113 | כדי |  |
| 2:11.114 | להיות |  |
| 2:11.115 | מליץ |  |
| 2:11.116 | בינו |  |
| 2:11.117 | ובין |  |
| 2:11.118 | רבון |  |
| 2:11.119 | העולמים | . |
| 2:11.120 | כל |  |
| 2:11.121 | אחד |  |
| 2:11.122 | מחמשה |  |
| 2:11.123 | אלו |  |
| 2:11.124 | מין |  |
| 2:11.125 | ע"כ | . |
| 2:11.126 | וכתב |  |
| 2:11.127 | הרב |  |
| 2:11.128 | לחם |  |
| 2:11.129 | משנה |  |
| 2:11.130 | על |  |
| 2:11.131 | חלוקה |  |
| 2:11.132 | ד | ' |
| 2:11.133 | שכתב |  |
| 2:11.134 | רבינו |  |
| 2:11.135 | והאומר |  |
| 2:11.136 | שאינו |  |
| 2:11.137 | לבדו |  |
| 2:11.138 | הראשון |  |
| 2:11.139 | וכו | ', |
| 2:11.140 | כ | ' |
| 2:11.141 | הראב"ד | , |
| 2:11.142 | כאותו |  |
| 2:11.143 | שאמר |  |
| 2:11.144 | אלהיכם |  |
| 2:11.145 | צייר |  |
| 2:11.146 | גדול |  |
| 2:11.147 | הוא |  |
| 2:11.148 | אלא |  |
| 2:11.149 | שמצא |  |
| 2:11.150 | לו |  |
| 2:11.151 | סמנים |  |
| 2:11.152 | גדולים |  |
| 2:11.153 | תהו |  |
| 2:11.154 | ובהו |  |
| 2:11.155 | וחשך |  |
| 2:11.156 | ומים |  |
| 2:11.157 | ורוח |  |
| 2:11.158 | ובהם |  |
| 2:11.159 | עשה |  |
| 2:11.160 | מה |  |
| 2:11.161 | שעשה | , |
| 2:11.162 | עכ"ל |  |
| 2:11.163 | הראב"ד | . |
| 2:11.164 | ופי | ' |
| 2:11.165 | הרב |  |
| 2:11.166 | לח"מ |  |
| 2:11.167 | שכוונתו |  |
| 2:11.168 | להוסיף |  |
| 2:11.169 | על |  |
| 2:11.170 | דברי |  |
| 2:11.171 | רבי | ' |
| 2:11.172 | שאעפ"י |  |
| 2:11.173 | שמודה |  |
| 2:11.174 | שלא |  |
| 2:11.175 | קדמה |  |
| 2:11.176 | שום |  |
| 2:11.177 | סבה |  |
| 2:11.178 | לאלהינו |  |
| 2:11.179 | שבשבילה |  |
| 2:11.180 | נהיה | , |
| 2:11.181 | אלא |  |
| 2:11.182 | שאומר |  |
| 2:11.183 | שברא |  |
| 2:11.184 | העולם |  |
| 2:11.185 | יש |  |
| 2:11.186 | מיש | , |
| 2:11.187 | הרי |  |
| 2:11.188 | זה |  |
| 2:11.189 | מין |  |
| 2:11.190 | ע"ש | . |
| 2:11.191 | ובנדון |  |
| 2:11.192 | דידן |  |
| 2:11.193 | אמונת |  |
| 2:11.194 | הקבלה |  |
| 2:11.195 | החדשה |  |
| 2:11.196 | כפי |  |
| 2:11.197 | המבואר |  |
| 2:11.198 | בספריהם |  |
| 2:11.199 | לפי |  |
| 2:11.200 | הבנתינו |  |
| 2:11.201 | והבנת |  |
| 2:11.202 | רבים |  |
| 2:11.203 | וגדולים |  |
| 2:11.204 | חכמי |  |
| 2:11.205 | לב |  |
| 2:11.206 | מרבותינו |  |
| 2:11.207 | הפורשים |  |
| 2:11.208 | ממנה |  |
| 2:11.209 | ראשונים |  |
| 2:11.210 | ואחרונים |  |
| 2:11.211 | כהריב"ש |  |
| 2:11.212 | וחו"י |  |
| 2:11.213 | המובא |  |
| 2:11.214 | בפתחי |  |
| 2:11.215 | תשובה |  |
| 2:11.216 | ושאר |  |
| 2:11.217 | אחרונים |  |
| 2:11.218 | נקבצו |  |
| 2:11.219 | באו |  |
| 2:11.220 | בה |  |
| 2:11.221 | תלת |  |
| 2:11.222 | וארבע |  |
| 2:11.223 | מלי |  |
| 2:11.224 | מהני |  |
| 2:11.225 | דאמור |  |
| 2:11.226 | רבנן | , |
| 2:11.227 | דהמאמין |  |
| 2:11.228 | כן |  |
| 2:11.229 | אין |  |
| 2:11.230 | לו |  |
| 2:11.231 | חלק |  |
| 2:11.232 | לעה"ב | . |
| 2:11.233 | חדא |  |
| 2:11.234 | רבוי |  |
| 2:11.235 | האלוהות | , |
| 2:11.236 | א"ס |  |
| 2:11.237 | ואדם |  |
| 2:11.238 | קדמאה |  |
| 2:11.239 | ואדם |  |
| 2:11.240 | קדמון |  |
| 2:11.241 | ועתיק |  |
| 2:11.242 | ואריך |  |
| 2:11.243 | אנפין |  |
| 2:11.244 | ואבא |  |
| 2:11.245 | ואימא |  |
| 2:11.246 | וזעיר |  |
| 2:11.247 | ונוקביה | ! |
| 2:11.248 | שנית | , |
| 2:11.249 | שהם |  |
| 2:11.250 | בעלי |  |
| 2:11.251 | גוף |  |
| 2:11.252 | דק | , |
| 2:11.253 | דהיינו |  |
| 2:11.254 | אור | , |
| 2:11.255 | והא"ס |  |
| 2:11.256 | נשמה |  |
| 2:11.257 | לגופות |  |
| 2:11.258 | הללו | . |
| 2:11.259 | שלישית | , |
| 2:11.260 | שהעבודה |  |
| 2:11.261 | אינה |  |
| 2:11.262 | לסבה |  |
| 2:11.263 | הראשונה |  |
| 2:11.264 | שהוא |  |
| 2:11.265 | נקרא |  |
| 2:11.266 | אצלם |  |
| 2:11.267 | א"ס | , |
| 2:11.268 | רק |  |
| 2:11.269 | לז"א |  |
| 2:11.270 | שהוא |  |
| 2:11.271 | סבה |  |
| 2:11.272 | אחרונה |  |
| 2:11.273 | שבאלו |  |
| 2:11.274 | הסבות | . |
| 2:11.275 | רביעית | , |
| 2:11.276 | שהוא |  |
| 2:11.277 | אמצעי |  |
| 2:11.278 | הממשיך |  |
| 2:11.279 | השפע |  |
| 2:11.280 | מאלו |  |
| 2:11.281 | המדות |  |
| 2:11.282 | העליונות |  |
| 2:11.283 | שהם |  |
| 2:11.284 | עתיק |  |
| 2:11.285 | ואריך |  |
| 2:11.286 | אנפין |  |
| 2:11.287 | ואבא |  |
| 2:11.288 | ואימא | , |
| 2:11.289 | וקורין |  |
| 2:11.290 | אותו |  |
| 2:11.291 | בן |  |
| 2:11.292 | לאו"א | , |
| 2:11.293 | והוא |  |
| 2:11.294 | ג"כ |  |
| 2:11.295 | אב |  |
| 2:11.296 | לכתר |  |
| 2:11.297 | דבריאה |  |
| 2:11.298 | כפי |  |
| 2:11.299 | השתלשלות |  |
| 2:11.300 | העולמות |  |
| 2:11.301 | לפי |  |
| 2:11.302 | דעתם |  |
| 2:11.303 | רחמנא |  |
| 2:11.304 | ליצלן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

