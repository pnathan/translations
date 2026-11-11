# Review pack: milhamot-hashem perek-2-05 (Milchamot ha-Shem 2:10–2:11)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-05`.

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


### Translation under review (2:10)

English: In the Yerushalmi, Shabbat, chapter Bemah Ishah, our Sages of blessed memory said: "'And the appearance of the fourth is like a son of the gods.' Rabbi Reuven said: At that hour an angel came down and slapped that wicked one on his mouth, and said to him: 'Is your word right, and is there a son for him?' He returned and said: 'Blessed be the God of Shadrach, Meishach and Abed-nego, who sent His angel and saved His servants.' 'His son' is not written here, but 'his angel.' Thus far." That is to say, Nebuchadnezzar the wicked thought that the angels are effects, descending in a chain of descent from him by nature, like a son from the father, according to the words of the Kabbalists concerning the sefirot of all the worlds, and likewise the souls; and for this the angel slapped him. And he returned and said: 'who sent His angel, etc.' Here, then, it is explicit from the words of Midrash Rabbah and the Yerushalmi that this is contrary to the words of the Kabbalists. And all that exists, the Holy One, blessed be He, brought into being in His kindness, something from nothing (yesh me-ayin), from the first form down to a small gnat; and not that they are drawn from His essence, as the son comes into being by the power of the seed drawn from his father.

Literal: in the Yerushalmi Shabbat chapter Bemah Ishah said our Sages of blessed memory the appearance of the fourth is like a son of gods said Rabbi Reuven in that hour came down an angel and slapped that wicked one on his mouth and said to him right his word and a son is there for him he returned and said Blessed their God of Shadrach Meishach and Abed Nego who sent his angel and saved his servants his son not written here but his angel thus far that is to say that Nebuchadnezzar the wicked thought that the angels they are effects and descend in a chain of descent from him by nature like a son from the father according to the words of the Kabbalists concerning the sefirot of all the worlds and likewise the souls and for this slapped him the angel and returned and said who sent his angel etc. here for you explicit comes out from the words of Midrash Rabbah and the Yerushalmi contrary to the words of the Kabbalists and all what exists brought into being the Holy One, blessed be He in His kindness something from nothing from the form first down to a gnat small and not that they are drawn from His essence as that the son comes into being by the power of the seed drawn from his father

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:10.1 | in the Yerushalmi |  |  |  |
| 2:10.2 | Shabbat |  |  |  |
| 2:10.3 | chapter |  |  | פ' = פרק: chapter. |
| 2:10.4 | Bemah |  |  |  |
| 2:10.5 | Ishah |  |  |  |
| 2:10.6 | said |  |  |  |
| 2:10.7 | our Sages of blessed memory |  |  | Hazal: 'our Sages of blessed memory'. The Yerushalmi passage that follows is quoted (Shabbat, chapter Bemah Ishah), up to 2:10.53. |
| 2:10.8 | the appearance |  |  | Aramaic: the appearance of the fourth figure in Dan 3:25 (וריויה, 'and the appearance'). |
| 2:10.9 | of |  |  |  |
| 2:10.10 | the fourth |  |  |  |
| 2:10.11 | is like |  |  |  |
| 2:10.12 | a son |  |  |  |
| 2:10.13 | of gods |  |  | Aramaic bar elahin: 'son of the gods' (Dan 3:25). |
| 2:10.14 | said |  |  |  |
| 2:10.15 | Rabbi |  |  | ר' = Rabbi. |
| 2:10.16 | Reuven |  |  |  |
| 2:10.17 | in that |  |  |  |
| 2:10.18 | hour |  |  |  |
| 2:10.19 | came down |  |  |  |
| 2:10.20 | an angel |  |  |  |
| 2:10.21 | and slapped |  |  | סטר: to slap. |
| 2:10.22 | that |  |  |  |
| 2:10.23 | wicked one |  |  | רשיעא: the wicked one, Nebuchadnezzar, named at 2:10.55. |
| 2:10.24 | on |  |  |  |
| 2:10.25 | his mouth |  |  |  |
| 2:10.26 | and said |  |  |  |
| 2:10.27 | to him |  |  |  |
| 2:10.28 | right |  |  | תקין: right, correct (Aramaic). |
| 2:10.29 | his word |  |  |  |
| 2:10.30 | and a son |  |  |  |
| 2:10.31 | is there |  |  |  |
| 2:10.32 | for him |  |  |  |
| 2:10.33 | he returned |  |  |  |
| 2:10.34 | and said |  |  |  |
| 2:10.35 | Blessed |  |  | The blessing is Dan 3:28 (Aramaic), in the Yerushalmi's wording. |
| 2:10.36 | their God |  |  |  |
| 2:10.37 | of |  |  |  |
| 2:10.38 | Shadrach |  |  |  |
| 2:10.39 | Meishach |  |  |  |
| 2:10.40 | and Abed |  |  |  |
| 2:10.41 | Nego |  |  |  |
| 2:10.42 | who |  |  |  |
| 2:10.43 | sent |  |  |  |
| 2:10.44 | his angel |  |  |  |
| 2:10.45 | and saved |  |  |  |
| 2:10.46 | his servants |  |  |  |
| 2:10.47 | his son |  |  | בריה: 'son'. The Yerushalmi says the verse writes 'angel' here, not 'son'. |
| 2:10.48 | not |  |  |  |
| 2:10.49 | written |  |  |  |
| 2:10.50 | here |  |  |  |
| 2:10.51 | but |  |  |  |
| 2:10.52 | his angel |  |  |  |
| 2:10.53 | thus far |  |  | ע"כ: 'thus far', closing the Yerushalmi quotation. |
| 2:10.54 | that is to say |  |  | ור"ל: 'that is to say' (רוצה לומר). From here the author speaks in his own voice. |
| 2:10.55 | that Nebuchadnezzar |  |  | Nebuchadnezzar, king of Babylon, of Dan 3. |
| 2:10.56 | the wicked |  |  |  |
| 2:10.57 | thought |  |  |  |
| 2:10.58 | that the angels |  |  |  |
| 2:10.59 | they |  |  |  |
| 2:10.60 | are effects |  |  | עלולים: effects, the things caused (עלה / עלול, cause and effect). |
| 2:10.61 | and descend in a chain of descent |  |  | משתלשלים: descend in a chain, from השתלשלות: the chain of descent. |
| 2:10.62 | from him |  |  |  |
| 2:10.63 | by nature |  |  |  |
| 2:10.64 | like a son |  |  |  |
| 2:10.65 | from |  |  |  |
| 2:10.66 | the father |  |  |  |
| 2:10.67 | according to the words of |  |  |  |
| 2:10.68 | the Kabbalists |  |  | המקובלים: the Kabbalists (the author's term). |
| 2:10.69 | concerning the sefirot |  |  | ספירות: sefirot. |
| 2:10.70 | of |  |  |  |
| 2:10.71 | all |  |  |  |
| 2:10.72 | the worlds |  |  |  |
| 2:10.73 | and likewise |  |  |  |
| 2:10.74 | the souls |  |  |  |
| 2:10.75 | and for this |  |  |  |
| 2:10.76 | slapped him |  |  |  |
| 2:10.77 | the angel |  |  |  |
| 2:10.78 | and returned |  |  |  |
| 2:10.79 | and said |  |  |  |
| 2:10.80 | who |  |  |  |
| 2:10.81 | sent |  |  |  |
| 2:10.82 | his angel |  |  |  |
| 2:10.83 | etc. |  |  | וכו': 'etc.'. The quotation is cut short. |
| 2:10.84 | here |  |  |  |
| 2:10.85 | for you |  |  |  |
| 2:10.86 | explicit |  |  |  |
| 2:10.87 | comes out |  |  |  |
| 2:10.88 | from the words of |  |  |  |
| 2:10.89 | Midrash |  |  |  |
| 2:10.90 | Rabbah |  |  |  |
| 2:10.91 | and the Yerushalmi |  |  |  |
| 2:10.92 | contrary to |  |  |  |
| 2:10.93 | the words of |  |  |  |
| 2:10.94 | the Kabbalists |  |  |  |
| 2:10.95 | and all |  |  |  |
| 2:10.96 | what exists |  |  |  |
| 2:10.97 | brought into being |  |  |  |
| 2:10.98 | the Holy One, blessed be He |  |  |  |
| 2:10.99 | in His kindness |  |  |  |
| 2:10.100 | something |  |  | יש מאין: 'something from nothing' (creation ex nihilo). Decisions.md has no fixed form. |
| 2:10.101 | from nothing |  |  |  |
| 2:10.102 | from the form |  |  |  |
| 2:10.103 | first |  |  |  |
| 2:10.104 | down to |  |  |  |
| 2:10.105 | a gnat |  |  |  |
| 2:10.106 | small |  |  |  |
| 2:10.107 | and not |  |  |  |
| 2:10.108 | that they |  |  |  |
| 2:10.109 | are drawn |  |  |  |
| 2:10.110 | from His essence |  |  |  |
| 2:10.111 | as |  |  |  |
| 2:10.112 | that the son |  |  |  |
| 2:10.113 | comes into being |  |  |  |
| 2:10.114 | by the power of |  |  |  |
| 2:10.115 | the seed |  |  |  |
| 2:10.116 | drawn |  |  |  |
| 2:10.117 | from his father |  |  |  |

Phrases:
- 2:10.1 2:10.2 2:10.3 2:10.4 2:10.5 2:10.6 2:10.7 | In the Yerushalmi, Shabbat, chapter Bemah Ishah, our Sages of blessed memory said: | in the Yerushalmi Shabbat chapter Bemah Ishah said our Sages of blessed memory
- 2:10.8 2:10.9 2:10.10 2:10.11 2:10.12 2:10.13 | "'And the appearance of the fourth is like a son of the gods.' | the appearance of the fourth is like a son of gods
- 2:10.14 2:10.15 2:10.16 2:10.17 2:10.18 2:10.19 2:10.20 2:10.21 2:10.22 2:10.23 2:10.24 2:10.25 | Rabbi Reuven said: At that hour an angel came down and slapped that wicked one on his mouth, | said Rabbi Reuven in that hour came down an angel and slapped that wicked one on his mouth
- 2:10.26 2:10.27 2:10.28 2:10.29 2:10.30 2:10.31 2:10.32 | and said to him: 'Is your word right, and is there a son for him?' | and said to him right his word and a son is there for him
- 2:10.33 2:10.34 | He returned and said: | he returned and said
- 2:10.35 2:10.36 2:10.37 2:10.38 2:10.39 2:10.40 2:10.41 2:10.42 2:10.43 2:10.44 2:10.45 2:10.46 | 'Blessed be the God of Shadrach, Meishach and Abed-nego, who sent His angel and saved His servants.' | Blessed their God of Shadrach Meishach and Abed Nego who sent his angel and saved his servants
- 2:10.47 2:10.48 2:10.49 2:10.50 2:10.51 2:10.52 2:10.53 | 'His son' is not written here, but 'his angel.' Thus far." | his son not written here but his angel thus far
- 2:10.54 2:10.55 2:10.56 2:10.57 2:10.58 2:10.59 2:10.60 2:10.61 2:10.62 2:10.63 2:10.64 2:10.65 2:10.66 2:10.67 2:10.68 2:10.69 2:10.70 2:10.71 2:10.72 2:10.73 2:10.74 2:10.75 2:10.76 2:10.77 | That is to say, Nebuchadnezzar the wicked thought that the angels are effects, descending in a chain of descent from him by nature, like a son from the father, according to the words of the Kabbalists concerning the sefirot of all the worlds, and likewise the souls; and for this the angel slapped him. | that is to say that Nebuchadnezzar the wicked thought that the angels they are effects and descend in a chain of descent from him by nature like a son from the father according to the words of the Kabbalists concerning the sefirot of all the worlds and likewise the souls and for this slapped him the angel
- 2:10.78 2:10.79 2:10.80 2:10.81 2:10.82 2:10.83 | And he returned and said: 'who sent His angel, etc.' | and returned and said who sent his angel etc.
- 2:10.84 2:10.85 2:10.86 2:10.87 2:10.88 2:10.89 2:10.90 2:10.91 2:10.92 2:10.93 2:10.94 | Here, then, it is explicit from the words of Midrash Rabbah and the Yerushalmi that this is contrary to the words of the Kabbalists. | here for you explicit comes out from the words of Midrash Rabbah and the Yerushalmi contrary to the words of the Kabbalists
- 2:10.95 2:10.96 2:10.97 2:10.98 2:10.99 2:10.100 2:10.101 2:10.102 2:10.103 2:10.104 2:10.105 2:10.106 2:10.107 2:10.108 2:10.109 2:10.110 2:10.111 2:10.112 2:10.113 2:10.114 2:10.115 2:10.116 2:10.117 | And all that exists, the Holy One, blessed be He, brought into being in His kindness, something from nothing (yesh me-ayin), from the first form down to a small gnat; and not that they are drawn from His essence, as the son comes into being by the power of the seed drawn from his father. | and all what exists brought into being the Holy One, blessed be He in His kindness something from nothing from the form first down to a gnat small and not that they are drawn from His essence as that the son comes into being by the power of the seed drawn from his father

Grammar notes:
- The Yerushalmi is quoted as the author's source, and its Aramaic is translated as he gives it. Dan 3:25 and 3:28 are quoted inside it, in single quotation marks.
- 'ע"כ' (2:10.53) closes the Yerushalmi quotation. From 2:10.54 the author argues in his own voice, and his argument turns on the words 'בר' (son) and 'מלאכיה' (his angel).
- 'ובר אית ליה?' is an Aramaic question, 'and is there a son for him?'. 'סטרו' (2:10.21, 2:10.76) is a verb of the angel's action, 'slapped him'.
- The Hebrew of the second clause 'שהם נמשכים' (2:10.108-109) is a negation, 'not that they are drawn': the author rejects the Kabbalists' reading of the chain.

Translator's note: The Yerushalmi passage is the author's source. The argument of the paragraph is set by 2:10.54 onward, where the author draws his conclusion from it against the Kabbalists.

Commentary rulings:
- apparatus: Dan 3:25 (Aramaic): the Yerushalmi's wording 'וריויה די רביעאה דמי לבר אלהין' is kept; 'son of the gods' is the verse's phrase.
- apparatus: Dan 3:28 (Aramaic): the blessing of the three men, with 'בריה' contrasted with 'מלאכיה' as the Yerushalmi gives them.
- apparatus: Yerushalmi: tractate and chapter as the author gives them (Shabbat, 'Bemah Ishah'); no folio or halakhah number is added.
- apparatus: Q2 (Wikisource additions): the base text is translated; no word is flagged until the 1931 scan is collated.

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


### Translation under review (2:11)

English: The conclusion from all this is that no face or form is to be worshipped of all the created things that exist in the upper and the lower world, which the Name, may He be blessed, brought into being in His kindness. For to Him alone it is fitting to worship and to pray, and for this the commandment and the warning came at Sinai: "I am the LORD your God … You shall have no other gods before Me." as the Rambam of blessed memory wrote in chapter two of the Laws of Idolatry. and in chapter three of the Laws of Repentance he wrote, and these are his words: "And these are those who have no share in the world to come, but are cut off and perish for the greatness of their wickedness and their sins, forever and ever: the heretics and the Epicureans, etc. And in the law that follows it he wrote: Five are those that are called heretics. The one who says that there is no God, and no ruler of the world, and the one who says that there is a ruler, but they are two or more, and the one who says that there is one Lord, but he is a body and possessed of a form, and the one who says that he is not alone the first and the rock of all. Likewise one who worships another in order to be an intercessor between him and the Lord of the worlds. Each of these five is a heretic. Thus far." And the Lechem Mishneh wrote on the fourth division that our Master wrote, 'and the one who says that He is not alone the first, etc.' The Ravad wrote: "like one who said, your God is a great painter, but he found for himself great pigments, tohu and bohu, and darkness and water and wind, and with them made what he made." Thus far, the words of the Ravad. And the Lechem Mishneh explains that his intention is to add to the words of our Master that although he concedes that no cause preceded our God, for whose sake He came to be, but says that He created the world from something (yesh me-yesh), then this one is a heretic; see there. In our case, the belief of the new Kabbalah, as explained in their books, according to our understanding and the understanding of many great men, wise of heart, among our Rabbis who separated from it, early and late, like the Ribash and Chavat Yair, cited in Pitchei Teshuvah and the other later authorities, gathered and came against it, three or four of these words that our Rabbis said: "Whoever believes thus has no share in the world to come." First, a multiplicity of divinities: Ein Sof, and Adam Kadma'ah, and Adam Kadmon, and Atik, and Arikh Anpin, and Abba and Imma, and Ze'ir and his Nukva! Second, that they have a subtle body, namely light, and Ein Sof is a soul to these bodies. Third, that the worship is not to the First Cause, which they call Ein Sof, but to Ze'ir Anpin, which is the last cause among these causes. Fourth, that he is the intermediary drawing down the flow from those upper attributes, which are Atik, Arikh Anpin, Abba and Imma, and they call him son of Abba and Imma, and he is also father to the Keter of Beriah, according to the chain of descent of the worlds, in their opinion, the Merciful One save us.

Literal: the conclusion from this that not to worship any face and form of all the created things that exist in the world the upper and the lower that brought into being the Name, may He be blessed in His kindness for to Him alone fitting to worship and to pray and on this came the commandment and the warning at Sinai I am the LORD your God not there shall be for you gods other as wrote the Rambam of blessed memory in chapter two of the Laws of Idolatry and in chapter three of the Laws of Repentance wrote and these are his words and these are who have no for them share in the world to come but are cut off and perish for the greatness of their wickedness and their sins forever and to ages the heretics and the Epicureans etc. and in the law that follows it wrote five are that are called heretics the one who says that there is no any God and no for the world ruler the one who says that there is a ruler but they are two or more the one who says that there is a Lord one but that he is a body and possessed of a form the one who says that he is not alone the first and rock of all and likewise the one who worships another in order to be an intercessor between him and the Lord of the worlds each one of the five these a heretic thus far and wrote the Rav Lechem Mishneh on the division the fourth that he wrote our Master the one who says that he is not alone the first etc. wrote the Ravad like the one who said your God a painter great he is but who found for himself pigments great tohu and bohu and darkness and water and wind and with them he made what that he made thus far the Ravad and explains the Rav the Lechem Mishneh that his intention is to add to the words of our Master that although he concedes that no preceded any cause our God for whose sake he came to be but says that He created the world something from something then this one is a heretic see there and in our case the belief of the Kabbalah the new according to as explained in their books according to our understanding and the understanding of many great wise of heart among our Rabbis who separated from it early and late the Ribash Chavat Yair cited in Pitchei Teshuvah and the other later authorities gathered came against it three and four words of these that said our Rabbis whoever believes thus has no to him a share in the world to come first a multiplicity of divinities Ein Sof and Adam Kadma'ah and Adam Kadmon and Atik and Arikh Anpin and Abba and Imma and Ze'ir and his Nukva second that they are possessors of a body subtle namely light and Ein Sof a soul to these bodies these third that the worship is not to the cause first which is called by them Ein Sof but to Ze'ir Anpin which is a cause last among these causes fourth that he is intermediary drawing down the flow from those attributes upper which are Atik and Arikh Anpin and Abba and Imma and they call him son of Abba and Imma and he also father to the Keter of Beriah according to the chain of descent of the worlds according to their opinion the Merciful One save us

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:11.1 | the conclusion |  |  | המורם: the conclusion, what is inferred from what precedes. |
| 2:11.2 | from |  |  |  |
| 2:11.3 | this |  |  |  |
| 2:11.4 | that not |  |  |  |
| 2:11.5 | to worship |  |  |  |
| 2:11.6 | any |  |  |  |
| 2:11.7 | face |  |  | פרצוף: 'face', in the generic sense (not the Kabbalistic partzuf). |
| 2:11.8 | and form |  |  | צורה: form. |
| 2:11.9 | of all |  |  |  |
| 2:11.10 | the created things |  |  |  |
| 2:11.11 | that exist |  |  |  |
| 2:11.12 | in the world |  |  |  |
| 2:11.13 | the upper |  |  |  |
| 2:11.14 | and the lower |  |  |  |
| 2:11.15 | that brought into being |  |  |  |
| 2:11.16 | the Name, may He be blessed |  |  | השי"ת = 'the Name, may He be blessed'. |
| 2:11.17 | in His kindness |  |  |  |
| 2:11.18 | for |  |  |  |
| 2:11.19 | to Him |  |  |  |
| 2:11.20 | alone |  |  |  |
| 2:11.21 | fitting |  |  |  |
| 2:11.22 | to worship |  |  |  |
| 2:11.23 | and to pray |  |  |  |
| 2:11.24 | and on |  |  |  |
| 2:11.25 | this |  |  |  |
| 2:11.26 | came |  |  |  |
| 2:11.27 | the commandment |  |  |  |
| 2:11.28 | and the warning |  |  |  |
| 2:11.29 | at Sinai |  |  |  |
| 2:11.30 | I am |  |  |  |
| 2:11.31 | the LORD |  |  | ה' = the Tetragrammaton, read as 'the LORD' in the verse. |
| 2:11.32 | your God |  |  |  |
| 2:11.33 | not |  |  |  |
| 2:11.34 | there shall be |  |  |  |
| 2:11.35 | for you |  |  |  |
| 2:11.36 | gods |  |  |  |
| 2:11.37 | other |  |  |  |
| 2:11.38 | as |  |  |  |
| 2:11.39 | wrote |  |  | שכ' = שכתב: 'wrote'. |
| 2:11.40 | the Rambam |  |  | הרמב"ם: the Rambam (Moses Maimonides). |
| 2:11.41 | of blessed memory |  |  | ז"ל: 'of blessed memory'. |
| 2:11.42 | in chapter |  |  |  |
| 2:11.43 | two |  |  |  |
| 2:11.44 | of the Laws |  |  |  |
| 2:11.45 | of Idolatry |  |  | ע"ז: Avodah Zarah, 'strange worship'. Here the Laws of Idolatry of the Mishneh Torah. |
| 2:11.46 | and in chapter |  |  |  |
| 2:11.47 | three |  |  |  |
| 2:11.48 | of the Laws |  |  |  |
| 2:11.49 | of Repentance |  |  | Hilchot Teshuvah: the Laws of Repentance of the Mishneh Torah. |
| 2:11.50 | wrote |  |  |  |
| 2:11.51 | and these are his words |  |  | וז"ל: 'and these are his words'. |
| 2:11.52 | and these are |  |  |  |
| 2:11.53 | who have no |  |  |  |
| 2:11.54 | for them |  |  |  |
| 2:11.55 | share |  |  |  |
| 2:11.56 | in the world |  |  |  |
| 2:11.57 | to come |  |  |  |
| 2:11.58 | but |  |  |  |
| 2:11.59 | are cut off |  |  |  |
| 2:11.60 | and perish |  |  |  |
| 2:11.61 | for |  |  |  |
| 2:11.62 | the greatness of |  |  |  |
| 2:11.63 | their wickedness |  |  |  |
| 2:11.64 | and their sins |  |  |  |
| 2:11.65 | forever |  |  |  |
| 2:11.66 | and to |  |  |  |
| 2:11.67 | ages |  |  |  |
| 2:11.68 | the heretics |  |  | המינין: the heretics (singular min). |
| 2:11.69 | and the Epicureans |  |  | האפיקורסין: the apikorsim (singular apikoros). |
| 2:11.70 | etc. |  |  | וכו': 'etc.'. The quotation is cut short. |
| 2:11.71 | and in the law |  |  |  |
| 2:11.72 | that follows it |  |  |  |
| 2:11.73 | wrote |  |  | כ': 'wrote' (כתב). |
| 2:11.74 | five |  |  |  |
| 2:11.75 | are |  |  |  |
| 2:11.76 | that are called |  |  |  |
| 2:11.77 | heretics |  |  |  |
| 2:11.78 | the one who says |  |  |  |
| 2:11.79 | that there is no |  |  |  |
| 2:11.80 | any |  |  |  |
| 2:11.81 | God |  |  |  |
| 2:11.82 | and no |  |  |  |
| 2:11.83 | for the world |  |  |  |
| 2:11.84 | ruler |  |  | מנהיג: ruler, governor. |
| 2:11.85 | the one who says |  |  |  |
| 2:11.86 | that there is |  |  |  |
| 2:11.87 | a |  |  |  |
| 2:11.88 | ruler |  |  |  |
| 2:11.89 | but |  |  |  |
| 2:11.90 | they are |  |  |  |
| 2:11.91 | two |  |  |  |
| 2:11.92 | or |  |  |  |
| 2:11.93 | more |  |  |  |
| 2:11.94 | the one who says |  |  |  |
| 2:11.95 | that there is |  |  |  |
| 2:11.96 | a |  |  |  |
| 2:11.97 | Lord |  |  |  |
| 2:11.98 | one |  |  |  |
| 2:11.99 | but |  |  |  |
| 2:11.100 | that he is |  |  |  |
| 2:11.101 | a body |  |  |  |
| 2:11.102 | and possessed of |  |  |  |
| 2:11.103 | a form |  |  | תמונה: a visible form. |
| 2:11.104 | the one who says |  |  |  |
| 2:11.105 | that he is not |  |  |  |
| 2:11.106 | alone |  |  |  |
| 2:11.107 | the first |  |  |  |
| 2:11.108 | and rock |  |  |  |
| 2:11.109 | of all |  |  |  |
| 2:11.110 | and likewise |  |  |  |
| 2:11.111 | the one who worships |  |  |  |
| 2:11.112 | another |  |  |  |
| 2:11.113 | in order |  |  |  |
| 2:11.114 | to be |  |  |  |
| 2:11.115 | an intercessor |  |  | מליץ: intercessor. |
| 2:11.116 | between him |  |  |  |
| 2:11.117 | and |  |  |  |
| 2:11.118 | the Lord |  |  |  |
| 2:11.119 | of the worlds |  |  |  |
| 2:11.120 | each |  |  |  |
| 2:11.121 | one |  |  |  |
| 2:11.122 | of the five |  |  |  |
| 2:11.123 | these |  |  |  |
| 2:11.124 | a heretic |  |  |  |
| 2:11.125 | thus far |  |  | ע"כ: 'thus far', closing the Rambam's quotation. |
| 2:11.126 | and wrote |  |  |  |
| 2:11.127 | the Rav |  |  | הרב: 'the Rav', the honorific of the commentator. The Lechem Mishneh is a commentary on the Mishneh Torah. |
| 2:11.128 | Lechem |  |  |  |
| 2:11.129 | Mishneh |  |  |  |
| 2:11.130 | on |  |  |  |
| 2:11.131 | the division |  |  |  |
| 2:11.132 | the fourth |  |  |  |
| 2:11.133 | that he wrote |  |  |  |
| 2:11.134 | our Master |  |  | רבינו: 'our Master', the Rambam. |
| 2:11.135 | the one who says |  |  |  |
| 2:11.136 | that he is not |  |  |  |
| 2:11.137 | alone |  |  |  |
| 2:11.138 | the first |  |  |  |
| 2:11.139 | etc. |  |  | וכו': 'etc.'. |
| 2:11.140 | wrote |  |  |  |
| 2:11.141 | the Ravad |  |  | הראב"ד: the Ravad (Rabad), whose criticisms of the Mishneh Torah the Lechem Mishneh cites. |
| 2:11.142 | like the one |  |  |  |
| 2:11.143 | who said |  |  |  |
| 2:11.144 | your God |  |  |  |
| 2:11.145 | a painter |  |  |  |
| 2:11.146 | great |  |  |  |
| 2:11.147 | he is |  |  |  |
| 2:11.148 | but |  |  |  |
| 2:11.149 | who found |  |  |  |
| 2:11.150 | for himself |  |  |  |
| 2:11.151 | pigments |  |  |  |
| 2:11.152 | great |  |  |  |
| 2:11.153 | tohu |  |  |  |
| 2:11.154 | and bohu |  |  |  |
| 2:11.155 | and darkness |  |  |  |
| 2:11.156 | and water |  |  |  |
| 2:11.157 | and wind |  |  |  |
| 2:11.158 | and with them |  |  |  |
| 2:11.159 | he made |  |  |  |
| 2:11.160 | what |  |  |  |
| 2:11.161 | that he made |  |  |  |
| 2:11.162 | thus far |  |  | עכ"ל: 'thus far, the words of', closing the Ravad's quotation. |
| 2:11.163 | the Ravad |  |  |  |
| 2:11.164 | and explains |  |  |  |
| 2:11.165 | the Rav |  |  |  |
| 2:11.166 | the Lechem Mishneh |  |  | לח"מ: the Lechem Mishneh. |
| 2:11.167 | that his intention |  |  |  |
| 2:11.168 | is to add |  |  |  |
| 2:11.169 | to |  |  |  |
| 2:11.170 | the words of |  |  |  |
| 2:11.171 | our Master |  |  | רבי': 'our Master' (the Rambam). |
| 2:11.172 | that although |  |  |  |
| 2:11.173 | he concedes |  |  |  |
| 2:11.174 | that no |  |  |  |
| 2:11.175 | preceded |  |  |  |
| 2:11.176 | any |  |  |  |
| 2:11.177 | cause |  |  |  |
| 2:11.178 | our God |  |  |  |
| 2:11.179 | for whose sake |  |  |  |
| 2:11.180 | he came to be |  |  |  |
| 2:11.181 | but |  |  |  |
| 2:11.182 | says |  |  |  |
| 2:11.183 | that He created |  |  |  |
| 2:11.184 | the world |  |  |  |
| 2:11.185 | something |  |  |  |
| 2:11.186 | from something |  |  |  |
| 2:11.187 | then |  |  |  |
| 2:11.188 | this one |  |  |  |
| 2:11.189 | is a heretic |  |  |  |
| 2:11.190 | see there |  |  | ע"ש: 'see there' (עיין שם). |
| 2:11.191 | and in |  |  |  |
| 2:11.192 | our case |  |  | דידן: 'ours'; 'in our case'. |
| 2:11.193 | the belief of |  |  |  |
| 2:11.194 | the Kabbalah |  |  | הקבלה: the Kabbalah. The author's term is 'the new Kabbalah'. |
| 2:11.195 | the new |  |  |  |
| 2:11.196 | according to |  |  |  |
| 2:11.197 | as explained |  |  |  |
| 2:11.198 | in their books |  |  |  |
| 2:11.199 | according to |  |  |  |
| 2:11.200 | our understanding |  |  |  |
| 2:11.201 | and the understanding of |  |  |  |
| 2:11.202 | many |  |  |  |
| 2:11.203 | great |  |  |  |
| 2:11.204 | wise of |  |  |  |
| 2:11.205 | heart |  |  | לב: 'heart'; 'wise of heart'. |
| 2:11.206 | among our Rabbis |  |  |  |
| 2:11.207 | who separated |  |  |  |
| 2:11.208 | from it |  |  |  |
| 2:11.209 | early |  |  |  |
| 2:11.210 | and late |  |  |  |
| 2:11.211 | the Ribash |  |  | הריב"ש: the Ribash. |
| 2:11.212 | Chavat Yair |  |  | חו"י: read as Chavat Yair. The expansion is not certain. |
| 2:11.213 | cited |  |  |  |
| 2:11.214 | in Pitchei |  |  | פתחי תשובה: Pitchei Teshuvah, a commentary on the Shulchan Arukh. |
| 2:11.215 | Teshuvah |  |  |  |
| 2:11.216 | and the other |  |  |  |
| 2:11.217 | later authorities |  |  |  |
| 2:11.218 | gathered |  |  |  |
| 2:11.219 | came |  |  |  |
| 2:11.220 | against it |  |  |  |
| 2:11.221 | three |  |  |  |
| 2:11.222 | and four |  |  |  |
| 2:11.223 | words |  |  |  |
| 2:11.224 | of these |  |  |  |
| 2:11.225 | that said |  |  |  |
| 2:11.226 | our Rabbis |  |  |  |
| 2:11.227 | whoever believes |  |  |  |
| 2:11.228 | thus |  |  |  |
| 2:11.229 | has no |  |  |  |
| 2:11.230 | to him |  |  |  |
| 2:11.231 | a share |  |  |  |
| 2:11.232 | in the world to come |  |  | לעה"ב: le-olam ha-ba, 'the world to come'. |
| 2:11.233 | first |  |  |  |
| 2:11.234 | a multiplicity |  |  |  |
| 2:11.235 | of divinities |  |  |  |
| 2:11.236 | Ein Sof |  |  | א"ס: Ein Sof, literally 'without end'. This is the first occurrence in the unit. |
| 2:11.237 | and Adam |  |  |  |
| 2:11.238 | Kadma'ah |  |  | קדמאה: Adam Kadma'ah, a distinct entity from Adam Kadmon. |
| 2:11.239 | and Adam |  |  |  |
| 2:11.240 | Kadmon |  |  | קדמון: Adam Kadmon, 'Primordial Man'. |
| 2:11.241 | and Atik |  |  | עתיק: Atik, a partzuf. |
| 2:11.242 | and Arikh |  |  | אריך אנפין: Arikh Anpin, 'Long Face'. |
| 2:11.243 | Anpin |  |  |  |
| 2:11.244 | and Abba |  |  | אבא: Abba, 'Father'. |
| 2:11.245 | and Imma |  |  | אימא: Imma, 'Mother'. |
| 2:11.246 | and Ze'ir |  |  | זעיר: Ze'ir Anpin, 'Small Face'. |
| 2:11.247 | and his Nukva |  |  | נוקביה: his Nukva, 'female'. |
| 2:11.248 | second |  |  |  |
| 2:11.249 | that they are |  |  |  |
| 2:11.250 | possessors of |  |  |  |
| 2:11.251 | a body |  |  |  |
| 2:11.252 | subtle |  |  |  |
| 2:11.253 | namely |  |  |  |
| 2:11.254 | light |  |  |  |
| 2:11.255 | and Ein Sof |  |  |  |
| 2:11.256 | a soul |  |  | נשמה: soul. |
| 2:11.257 | to these bodies |  |  |  |
| 2:11.258 | these |  |  |  |
| 2:11.259 | third |  |  |  |
| 2:11.260 | that the worship |  |  |  |
| 2:11.261 | is not |  |  |  |
| 2:11.262 | to the cause |  |  |  |
| 2:11.263 | first |  |  |  |
| 2:11.264 | which |  |  |  |
| 2:11.265 | is called |  |  |  |
| 2:11.266 | by them |  |  |  |
| 2:11.267 | Ein Sof |  |  |  |
| 2:11.268 | but |  |  |  |
| 2:11.269 | to Ze'ir Anpin |  |  | לז"א: Ze'ir Anpin, 'small face'. |
| 2:11.270 | which is |  |  |  |
| 2:11.271 | a cause |  |  |  |
| 2:11.272 | last |  |  |  |
| 2:11.273 | among these |  |  |  |
| 2:11.274 | causes |  |  |  |
| 2:11.275 | fourth |  |  |  |
| 2:11.276 | that he is |  |  |  |
| 2:11.277 | intermediary |  |  |  |
| 2:11.278 | drawing down |  |  |  |
| 2:11.279 | the flow |  |  |  |
| 2:11.280 | from those |  |  |  |
| 2:11.281 | attributes |  |  |  |
| 2:11.282 | upper |  |  |  |
| 2:11.283 | which are |  |  |  |
| 2:11.284 | Atik |  |  |  |
| 2:11.285 | and Arikh |  |  |  |
| 2:11.286 | Anpin |  |  |  |
| 2:11.287 | and Abba |  |  |  |
| 2:11.288 | and Imma |  |  |  |
| 2:11.289 | and they call |  |  |  |
| 2:11.290 | him |  |  |  |
| 2:11.291 | son of |  |  |  |
| 2:11.292 | Abba and Imma |  |  | לאו"א: 'Abba and Imma'. The author argues from the name, 'son of Abba and Imma'. |
| 2:11.293 | and he |  |  |  |
| 2:11.294 | also |  |  |  |
| 2:11.295 | father |  |  |  |
| 2:11.296 | to the Keter |  |  | כתר: Keter, the crown. |
| 2:11.297 | of Beriah |  |  | דבריאה: of Beriah, the World of Creation. |
| 2:11.298 | according to |  |  |  |
| 2:11.299 | the chain of descent |  |  |  |
| 2:11.300 | of the worlds |  |  |  |
| 2:11.301 | according to |  |  |  |
| 2:11.302 | their opinion |  |  |  |
| 2:11.303 | the Merciful One |  |  | רחמנא ליצלן: 'the Merciful One save us', the formula kept, with its irony. |
| 2:11.304 | save us |  |  |  |

Phrases:
- 2:11.1 2:11.2 2:11.3 2:11.4 2:11.5 2:11.6 2:11.7 2:11.8 | The conclusion from all this is that no face or form is to be worshipped | the conclusion from this that not to worship any face and form
- 2:11.9 2:11.10 2:11.11 2:11.12 2:11.13 2:11.14 2:11.15 2:11.16 2:11.17 | of all the created things that exist in the upper and the lower world, which the Name, may He be blessed, brought into being in His kindness. | of all the created things that exist in the world the upper and the lower that brought into being the Name, may He be blessed in His kindness
- 2:11.18 2:11.19 2:11.20 2:11.21 2:11.22 2:11.23 | For to Him alone it is fitting to worship and to pray, | for to Him alone fitting to worship and to pray
- 2:11.24 2:11.25 2:11.26 2:11.27 2:11.28 2:11.29 | and for this the commandment and the warning came at Sinai: | and on this came the commandment and the warning at Sinai
- 2:11.30 2:11.31 2:11.32 2:11.33 2:11.34 2:11.35 2:11.36 2:11.37 | "I am the LORD your God … You shall have no other gods before Me." | I am the LORD your God not there shall be for you gods other
- 2:11.38 2:11.39 2:11.40 2:11.41 2:11.42 2:11.43 2:11.44 2:11.45 | as the Rambam of blessed memory wrote in chapter two of the Laws of Idolatry. | as wrote the Rambam of blessed memory in chapter two of the Laws of Idolatry
- 2:11.46 2:11.47 2:11.48 2:11.49 2:11.50 2:11.51 | and in chapter three of the Laws of Repentance he wrote, and these are his words: | and in chapter three of the Laws of Repentance wrote and these are his words
- 2:11.52 2:11.53 2:11.54 2:11.55 2:11.56 2:11.57 2:11.58 2:11.59 2:11.60 2:11.61 2:11.62 2:11.63 2:11.64 2:11.65 2:11.66 2:11.67 | "And these are those who have no share in the world to come, but are cut off and perish for the greatness of their wickedness and their sins, forever and ever: | and these are who have no for them share in the world to come but are cut off and perish for the greatness of their wickedness and their sins forever and to ages
- 2:11.68 2:11.69 2:11.70 | the heretics and the Epicureans, etc. | the heretics and the Epicureans etc.
- 2:11.71 2:11.72 2:11.73 2:11.74 2:11.75 2:11.76 2:11.77 | And in the law that follows it he wrote: Five are those that are called heretics. | and in the law that follows it wrote five are that are called heretics
- 2:11.78 2:11.79 2:11.80 2:11.81 2:11.82 2:11.83 2:11.84 | The one who says that there is no God, and no ruler of the world, | the one who says that there is no any God and no for the world ruler
- 2:11.85 2:11.86 2:11.87 2:11.88 2:11.89 2:11.90 2:11.91 2:11.92 2:11.93 | and the one who says that there is a ruler, but they are two or more, | the one who says that there is a ruler but they are two or more
- 2:11.94 2:11.95 2:11.96 2:11.97 2:11.98 2:11.99 2:11.100 2:11.101 2:11.102 2:11.103 | and the one who says that there is one Lord, but he is a body and possessed of a form, | the one who says that there is a Lord one but that he is a body and possessed of a form
- 2:11.104 2:11.105 2:11.106 2:11.107 2:11.108 2:11.109 | and the one who says that he is not alone the first and the rock of all. | the one who says that he is not alone the first and rock of all
- 2:11.110 2:11.111 2:11.112 2:11.113 2:11.114 2:11.115 2:11.116 2:11.117 2:11.118 2:11.119 2:11.120 2:11.121 2:11.122 2:11.123 2:11.124 2:11.125 | Likewise one who worships another in order to be an intercessor between him and the Lord of the worlds. Each of these five is a heretic. Thus far." | and likewise the one who worships another in order to be an intercessor between him and the Lord of the worlds each one of the five these a heretic thus far
- 2:11.126 2:11.127 2:11.128 2:11.129 2:11.130 2:11.131 2:11.132 2:11.133 2:11.134 2:11.135 2:11.136 2:11.137 2:11.138 2:11.139 | And the Lechem Mishneh wrote on the fourth division that our Master wrote, 'and the one who says that He is not alone the first, etc.' | and wrote the Rav Lechem Mishneh on the division the fourth that he wrote our Master the one who says that he is not alone the first etc.
- 2:11.140 2:11.141 2:11.142 2:11.143 2:11.144 2:11.145 2:11.146 2:11.147 2:11.148 2:11.149 2:11.150 2:11.151 2:11.152 2:11.153 2:11.154 2:11.155 2:11.156 2:11.157 2:11.158 2:11.159 2:11.160 2:11.161 2:11.162 2:11.163 | The Ravad wrote: "like one who said, your God is a great painter, but he found for himself great pigments, tohu and bohu, and darkness and water and wind, and with them made what he made." Thus far, the words of the Ravad. | wrote the Ravad like the one who said your God a painter great he is but who found for himself pigments great tohu and bohu and darkness and water and wind and with them he made what that he made thus far the Ravad
- 2:11.164 2:11.165 2:11.166 2:11.167 2:11.168 2:11.169 2:11.170 2:11.171 | And the Lechem Mishneh explains that his intention is to add to the words of our Master | and explains the Rav the Lechem Mishneh that his intention is to add to the words of our Master
- 2:11.172 2:11.173 2:11.174 2:11.175 2:11.176 2:11.177 2:11.178 2:11.179 2:11.180 2:11.181 2:11.182 2:11.183 2:11.184 2:11.185 2:11.186 2:11.187 2:11.188 2:11.189 2:11.190 | that although he concedes that no cause preceded our God, for whose sake He came to be, but says that He created the world from something (yesh me-yesh), then this one is a heretic; see there. | that although he concedes that no preceded any cause our God for whose sake he came to be but says that He created the world something from something then this one is a heretic see there
- 2:11.191 2:11.192 2:11.193 2:11.194 2:11.195 2:11.196 2:11.197 2:11.198 2:11.199 | In our case, the belief of the new Kabbalah, as explained in their books, according to | and in our case the belief of the Kabbalah the new according to as explained in their books according to
- 2:11.200 2:11.201 2:11.202 2:11.203 2:11.204 2:11.205 2:11.206 | our understanding and the understanding of many great men, wise of heart, among our Rabbis | our understanding and the understanding of many great wise of heart among our Rabbis
- 2:11.207 2:11.208 2:11.209 2:11.210 2:11.211 2:11.212 2:11.213 2:11.214 2:11.215 2:11.216 2:11.217 | who separated from it, early and late, like the Ribash and Chavat Yair, cited in Pitchei Teshuvah and the other later authorities, | who separated from it early and late the Ribash Chavat Yair cited in Pitchei Teshuvah and the other later authorities
- 2:11.218 2:11.219 2:11.220 2:11.221 2:11.222 2:11.223 2:11.224 2:11.225 2:11.226 2:11.227 2:11.228 2:11.229 2:11.230 2:11.231 2:11.232 | gathered and came against it, three or four of these words that our Rabbis said: "Whoever believes thus has no share in the world to come." | gathered came against it three and four words of these that said our Rabbis whoever believes thus has no to him a share in the world to come
- 2:11.233 2:11.234 2:11.235 2:11.236 2:11.237 2:11.238 2:11.239 2:11.240 2:11.241 2:11.242 2:11.243 2:11.244 2:11.245 2:11.246 2:11.247 | First, a multiplicity of divinities: Ein Sof, and Adam Kadma'ah, and Adam Kadmon, and Atik, and Arikh Anpin, and Abba and Imma, and Ze'ir and his Nukva! | first a multiplicity of divinities Ein Sof and Adam Kadma'ah and Adam Kadmon and Atik and Arikh Anpin and Abba and Imma and Ze'ir and his Nukva
- 2:11.248 2:11.249 2:11.250 2:11.251 2:11.252 2:11.253 2:11.254 2:11.255 2:11.256 2:11.257 2:11.258 | Second, that they have a subtle body, namely light, and Ein Sof is a soul to these bodies. | second that they are possessors of a body subtle namely light and Ein Sof a soul to these bodies these
- 2:11.259 2:11.260 2:11.261 2:11.262 2:11.263 2:11.264 2:11.265 2:11.266 2:11.267 2:11.268 2:11.269 2:11.270 2:11.271 2:11.272 2:11.273 2:11.274 | Third, that the worship is not to the First Cause, which they call Ein Sof, but to Ze'ir Anpin, which is the last cause among these causes. | third that the worship is not to the cause first which is called by them Ein Sof but to Ze'ir Anpin which is a cause last among these causes
- 2:11.275 2:11.276 2:11.277 2:11.278 2:11.279 2:11.280 2:11.281 2:11.282 2:11.283 2:11.284 2:11.285 2:11.286 2:11.287 2:11.288 2:11.289 2:11.290 2:11.291 2:11.292 | Fourth, that he is the intermediary drawing down the flow from those upper attributes, which are Atik, Arikh Anpin, Abba and Imma, and they call him son of Abba and Imma, | fourth that he is intermediary drawing down the flow from those attributes upper which are Atik and Arikh Anpin and Abba and Imma and they call him son of Abba and Imma
- 2:11.293 2:11.294 2:11.295 2:11.296 2:11.297 2:11.298 2:11.299 2:11.300 2:11.301 2:11.302 2:11.303 2:11.304 | and he is also father to the Keter of Beriah, according to the chain of descent of the worlds, in their opinion, the Merciful One save us. | and he also father to the Keter of Beriah according to the chain of descent of the worlds according to their opinion the Merciful One save us

Grammar notes:
- The paragraph is the Rambam's quotation, framed by the author's argument. The Rambam's words run from 2:11.51 ('וז"ל') to 2:11.125 ('ע"כ'), and the author's gloss on them starts at 2:11.126.
- 'כ'' (2:11.39, 2:11.73, 2:11.140) is 'wrote', an abbreviation of כתב. 'הרב' (2:11.127, 2:11.165) is an honorific of the commentator, not a title.
- 'שאין לעבוד' (2:11.4-5) is 'that one must not worship': 'לעבוד' is the verb 'to worship', used of God throughout the paragraph.
- The numbers 'חדא', 'שנית', 'שלישית', 'רביעית' (2:11.233, 248, 259, 275) introduce four objections. Each is followed by a clause with 'שהם' or 'שהעבודה'.

Translator's note: The paragraph is the author's conclusion from the Rambam's law, followed by his four objections to the new Kabbalah, in his own words and in the words of the Lechem Mishneh and the Ravad.

Commentary rulings:
- apparatus: Exod 20:2-3: quoted to 'אחרים' (2:11.30-37). 'before Me' is the fixed rendering of Exod 20:3 in decisions.md; the author's quotation has no 'before Me'.
- apparatus: Rambam, Laws of Idolatry ch. 2 and Laws of Repentance ch. 3 (Mishneh Torah), as the author gives them; no halakhah number is added.
- apparatus: The Lechem Mishneh and the Ravad are quoted as the author gives them; 'וכו'' is kept as 'etc.'.
- apparatus: No ruling in decisions.md covers יש מאין or יש מיש. The English gives the transliterated phrase with its sense.
- apparatus: חו"י: the expansion Chavat Yair is a reading, not a certainty. The abbreviation is kept as 'Chavat Yair' in the English.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

