# Translation pack: milhamot-hashem perek-4-20 (Milchamot ha-Shem 4:32–4:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-20`.

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
- **קצר אפים**: the author's Hebrew calque of זעיר אנפין, set against ארך
  אפים in Exod 34:6. Write "quick to anger", and render ארך אפים "slow to
  anger" everywhere, so that the pair shows in English. "באל זר קצר אפים
  (ז״א)" (1:4) → "a strange god, quick to anger (Ze'ir Anpin)"; "הנקרא (קצר
  אפים) זעיר אנפין" (3:3) → "called (quick to anger) Ze'ir Anpin". Add a note
  each time: Aramaic "small face"; Prov 14:17. "קצא אפים" (6:23) is translated
  as קצר, with the spelling recorded in `apparatus`.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
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


## Milchamot ha-Shem 4:32

Hebrew (Wikisource): לח) הנה גם מזה נתבאר לך ידידי שגגת רבותינו המקובלים החדשים שציירו ותיארו את השי"ת בכמה ציורים ותארים שונים. עגולים ויושר בכמה פרצופים שונים זה מזה, זה בתוך זה, וזה למעלה מזה, למרות קבלת רז"ל שקבלו שאסור לחשוב ולהרהר בדברים אלו שהם בכלל מה למעלה מה למטה מה לפנים ומה לאחור. ואמרו עוד המקובלים החדשים שבתחלה היה השי"ת ממלא את חלל כל העולם. ואחר כך צמצם את עצמו ונצטמק אל הצדדים סביב, ונשתנה מכמות שהיה כדי להמציא מקום וחלל לכל העולם. וזה לשון מהרח"ו בספר עץ חיים שער עיגולים ויושר ענף ב' והנה אחר הצמצום אשר אז נשאר מקום החלל ואויר פנוי וריקני באמצעות האור שלאין סוף ממש, הנה כבר היה מקום לשיוכלו להיות שם הנאצלים והנבראים והיצורים והנעשים, ואז המשיך מן אור אין סוף קו אחד ישר מן האור העגול שלו מלמעלה למטה ומשתלשל ויורד תוך החלל ההוא, וראש העליון שלהקו נמשך מן הא"ס עצמו ונוגע בו, אמנם סיום הקו ההוא למטה בסופו אינו נוגע באור האין סוף למטה, ובמקום החלל ההוא האציל וברא ויצר ועשה כל העולמות כולם, וקו זה כעין צנור דק אחד אשר דרך בו מתפשט ונמשך מימי האור העליון של אין סוף אל כל העולמות אשר במקום האויר והחלל ההוא, ונבאר עתה ענין חקירת המקובלים לדעת איך יש ראש וסוף בספירות, אמנם בהיות הקו ההוא ראשו נוגע באור אין סוף מצד העליון, וסופו אינו נמשך למטה עד מקום אין סוף הסובב תחת העולמות ואינו דבוק בו, לכן אז יצדק בו ראש וסוף, כי אם היו דרך שני קצוות מקבלים שפע האין סוף היו שני הקצוות בחינת ראשים שווים זה לזה ולא היה אז בחינת מעלה ומטה, וכן אם היה האין סוף נמשך מכל צדי המקום החלל ההוא, לא היה לא מעלה ולא מטה לא פנים ולא אחור. לא מזרח ולא מערב לא צפון ולא דרום, אך בהיות אור האין סוף נמשך דרך קו אחד וצנור דק בלבד יצדק מעלה ומטה פנים ואחור מזרח ומערב (דבריו אלה לא בהשכל דבשלמא ראש וסוף שפיר אך פנים ואחור מזרח ומערב איך יצדק?) והנה בהיות אור האין סוף נמשך בבחינת קו ישר תוך החלל, לא נמשך ונתפשט למטה תיכף, אלא היה מתפשט לאט לאט. רצוני לומר כי בתחלה התחיל קו האור להתפשט שם, ותכף בתחלת התפשטותו בסוד קו נתפשט ונמשך ונעשה כעין גלגל עגול מסביב. והעגול הזה היה בלתי דבוק אל האין סוף הסובב עליו מכל צדדיו, שאם יתדבק בו יחזור הדבר לכמות שהיה, ויתבטל באור אין סוף ולא יתראה כחו כלל, ויהיה הכל אור אין סוף לבד כבראשונה, (מדבריו אלה יש להקשות על דבריו שיזכיר להלן כי דרך הקו הזה יורד אור פנימי שהוא נשמת הספירות ואור מקיף שהוא מלבוש לספירות. למה אינו מתערב ומתדבק בגוף הספירות ויחזור עצם הספירות עם מלבושן ונשמתן כולם גולם אחד ויתבטל הכלי ויתערב באור פנימי ואור חיצון?) לכן העגול הזה סמוך אל האין סוף ובלתי מתדבק בו. וכל עיקר התקשרות ודביקות הנאצל ההוא עם אין סוף המאציל הוא ע"י הקו ההוא אשר דרך בו יורד ונמשך אור מן האין סוף המשפיע בעיגול ההוא. והאין סוף סובב ומקיף עליו מכל צדדיו ורחוק ממנו בהשואה אחת מכל צדדיו. כי גם הוא בבחינת עיגול סביב עליו. כי הוא מוכרח שהארת א"ס בנאצלים תהיה דרך הקו ההוא לבד. כי אם היה האור ההוא נמשך להם גם מכל סביבותיהם היו הנאצלים בבחינת המאציל עצמו בלתי גבול וקצבה, ולא עוד אלא אפילו הקו ההוא הוא דק מאד ולא עבה כדי שיהיה האור ההוא הנמשך אל הנאצלים במדה וקצבה אשר לסבה זאת נקראים הנאצלים עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה ומספר קצוב. והנה העיגול הראשון הזה היותר דבוק עם האין סוף הוא הנקרא כתר דאדם קדמון. ואחר כך נתפשט הקו הזה ונמשך מעט וחזר להתעגל ונעשה עיגול שני תוך עיגול ראשון. וזה נקרא עיגול החכמה דאדם קדמון, עוד נתפשט יותר למטה וחזר להתעגל ונעשה עיגול שלישי תוך עיגול שני הזה ונקרא עיגול בינה דאדם קדמון, ועל דרך זה היה הולך ומתעגל עד עיגול העשירי הנקרא עיגול מלכות דאדם קדמון וכו'. אמנם מבואר ופשוט הוא שכמה מיני עולמות נאצלו ונבראו ונוצרו ונעשו אלף אלפים ורבוא רבבות, וכולם כאחד תוך המקום החלל הנז"ל ואין דבר חוצה לו. וכל עולם ועולם יש בו עשר ספירות פרטיות. וכל ספירה וספירה שבכל עולם ועולם כלולה מעשר ספירות פרטי פרטות. ועתה נבאר בחינה שנית שיש בעשר ספירות, היא בחינת היושר בצורת אדם העליון. הנה דרך הקו הנז' לעיל המתפשט מלמעלה למטה אשר ממנו מתפשטים העיגולים, גם הקו ההוא מתפשט ביושר ממעלה למטה מראש גג העליון שלעיגול העליון עד למטה מתחתית סיום כל העיגולים ממש מלמעלה למטה כלול מעשר ספירות בסוד צלם אדם ישר בעל קומה זקופה כלול מרמ"ח איברים וכו'. והנה בחינה זו השנית נקראת צלם אלהים ועליה רמז הכתוב באמרו ויברא אלהים את האדם בצלמו בצלם אלהים ברא אותו וכו' ושכמעט רוב דברי הזהר והתיקונים מתעסקים בבחינה הזאת השנית שהיא בחינת היושר עכ"ל. ובענף ג' וד' כפל עוד הדברים שאין העשר ספירות דעולם האצילות הם ראשונים וגבוהים מכל מה שנאצלו. כי כמה עולמות אבי"ע קדמו להם. ולרוב העלמם לא שלחו בהם יד להזכירם בזהר ובתיקונים אלא ברמז נפלא. ושם ביאר עוד עד היכן מגיעים רגלי אדם קדמון דיושר. ועד היכן מגיעים רגלי עתיק יומין, ורגלי אריך אנפין, ורגלי אבא ואימא, ורגלי נוקבא דיליה. ושאבא ואימא קצרי קומה, ואין שיעור קומתם רק מהגרון דאריך אנפין עד הטבור שלאריך אנפין, וכדברים האלה מבואר ג"כ בשער ההקדמות ובספר כסא אליהו ומקדש מלך פרשת בראשית עיין שם ושם.

| id | word | punctuation after |
|---|---|---|
| 4:32.1 | לח | ) |
| 4:32.2 | הנה |  |
| 4:32.3 | גם |  |
| 4:32.4 | מזה |  |
| 4:32.5 | נתבאר |  |
| 4:32.6 | לך |  |
| 4:32.7 | ידידי |  |
| 4:32.8 | שגגת |  |
| 4:32.9 | רבותינו |  |
| 4:32.10 | המקובלים |  |
| 4:32.11 | החדשים |  |
| 4:32.12 | שציירו |  |
| 4:32.13 | ותיארו |  |
| 4:32.14 | את |  |
| 4:32.15 | השי"ת |  |
| 4:32.16 | בכמה |  |
| 4:32.17 | ציורים |  |
| 4:32.18 | ותארים |  |
| 4:32.19 | שונים | . |
| 4:32.20 | עגולים |  |
| 4:32.21 | ויושר |  |
| 4:32.22 | בכמה |  |
| 4:32.23 | פרצופים |  |
| 4:32.24 | שונים |  |
| 4:32.25 | זה |  |
| 4:32.26 | מזה | , |
| 4:32.27 | זה |  |
| 4:32.28 | בתוך |  |
| 4:32.29 | זה | , |
| 4:32.30 | וזה |  |
| 4:32.31 | למעלה |  |
| 4:32.32 | מזה | , |
| 4:32.33 | למרות |  |
| 4:32.34 | קבלת |  |
| 4:32.35 | רז"ל |  |
| 4:32.36 | שקבלו |  |
| 4:32.37 | שאסור |  |
| 4:32.38 | לחשוב |  |
| 4:32.39 | ולהרהר |  |
| 4:32.40 | בדברים |  |
| 4:32.41 | אלו |  |
| 4:32.42 | שהם |  |
| 4:32.43 | בכלל |  |
| 4:32.44 | מה |  |
| 4:32.45 | למעלה |  |
| 4:32.46 | מה |  |
| 4:32.47 | למטה |  |
| 4:32.48 | מה |  |
| 4:32.49 | לפנים |  |
| 4:32.50 | ומה |  |
| 4:32.51 | לאחור | . |
| 4:32.52 | ואמרו |  |
| 4:32.53 | עוד |  |
| 4:32.54 | המקובלים |  |
| 4:32.55 | החדשים |  |
| 4:32.56 | שבתחלה |  |
| 4:32.57 | היה |  |
| 4:32.58 | השי"ת |  |
| 4:32.59 | ממלא |  |
| 4:32.60 | את |  |
| 4:32.61 | חלל |  |
| 4:32.62 | כל |  |
| 4:32.63 | העולם | . |
| 4:32.64 | ואחר |  |
| 4:32.65 | כך |  |
| 4:32.66 | צמצם |  |
| 4:32.67 | את |  |
| 4:32.68 | עצמו |  |
| 4:32.69 | ונצטמק |  |
| 4:32.70 | אל |  |
| 4:32.71 | הצדדים |  |
| 4:32.72 | סביב | , |
| 4:32.73 | ונשתנה |  |
| 4:32.74 | מכמות |  |
| 4:32.75 | שהיה |  |
| 4:32.76 | כדי |  |
| 4:32.77 | להמציא |  |
| 4:32.78 | מקום |  |
| 4:32.79 | וחלל |  |
| 4:32.80 | לכל |  |
| 4:32.81 | העולם | . |
| 4:32.82 | וזה |  |
| 4:32.83 | לשון |  |
| 4:32.84 | מהרח"ו |  |
| 4:32.85 | בספר |  |
| 4:32.86 | עץ |  |
| 4:32.87 | חיים |  |
| 4:32.88 | שער |  |
| 4:32.89 | עיגולים |  |
| 4:32.90 | ויושר |  |
| 4:32.91 | ענף |  |
| 4:32.92 | ב | ' |
| 4:32.93 | והנה |  |
| 4:32.94 | אחר |  |
| 4:32.95 | הצמצום |  |
| 4:32.96 | אשר |  |
| 4:32.97 | אז |  |
| 4:32.98 | נשאר |  |
| 4:32.99 | מקום |  |
| 4:32.100 | החלל |  |
| 4:32.101 | ואויר |  |
| 4:32.102 | פנוי |  |
| 4:32.103 | וריקני |  |
| 4:32.104 | באמצעות |  |
| 4:32.105 | האור |  |
| 4:32.106 | שלאין |  |
| 4:32.107 | סוף |  |
| 4:32.108 | ממש | , |
| 4:32.109 | הנה |  |
| 4:32.110 | כבר |  |
| 4:32.111 | היה |  |
| 4:32.112 | מקום |  |
| 4:32.113 | לשיוכלו |  |
| 4:32.114 | להיות |  |
| 4:32.115 | שם |  |
| 4:32.116 | הנאצלים |  |
| 4:32.117 | והנבראים |  |
| 4:32.118 | והיצורים |  |
| 4:32.119 | והנעשים | , |
| 4:32.120 | ואז |  |
| 4:32.121 | המשיך |  |
| 4:32.122 | מן |  |
| 4:32.123 | אור |  |
| 4:32.124 | אין |  |
| 4:32.125 | סוף |  |
| 4:32.126 | קו |  |
| 4:32.127 | אחד |  |
| 4:32.128 | ישר |  |
| 4:32.129 | מן |  |
| 4:32.130 | האור |  |
| 4:32.131 | העגול |  |
| 4:32.132 | שלו |  |
| 4:32.133 | מלמעלה |  |
| 4:32.134 | למטה |  |
| 4:32.135 | ומשתלשל |  |
| 4:32.136 | ויורד |  |
| 4:32.137 | תוך |  |
| 4:32.138 | החלל |  |
| 4:32.139 | ההוא | , |
| 4:32.140 | וראש |  |
| 4:32.141 | העליון |  |
| 4:32.142 | שלהקו |  |
| 4:32.143 | נמשך |  |
| 4:32.144 | מן |  |
| 4:32.145 | הא"ס |  |
| 4:32.146 | עצמו |  |
| 4:32.147 | ונוגע |  |
| 4:32.148 | בו | , |
| 4:32.149 | אמנם |  |
| 4:32.150 | סיום |  |
| 4:32.151 | הקו |  |
| 4:32.152 | ההוא |  |
| 4:32.153 | למטה |  |
| 4:32.154 | בסופו |  |
| 4:32.155 | אינו |  |
| 4:32.156 | נוגע |  |
| 4:32.157 | באור |  |
| 4:32.158 | האין |  |
| 4:32.159 | סוף |  |
| 4:32.160 | למטה | , |
| 4:32.161 | ובמקום |  |
| 4:32.162 | החלל |  |
| 4:32.163 | ההוא |  |
| 4:32.164 | האציל |  |
| 4:32.165 | וברא |  |
| 4:32.166 | ויצר |  |
| 4:32.167 | ועשה |  |
| 4:32.168 | כל |  |
| 4:32.169 | העולמות |  |
| 4:32.170 | כולם | , |
| 4:32.171 | וקו |  |
| 4:32.172 | זה |  |
| 4:32.173 | כעין |  |
| 4:32.174 | צנור |  |
| 4:32.175 | דק |  |
| 4:32.176 | אחד |  |
| 4:32.177 | אשר |  |
| 4:32.178 | דרך |  |
| 4:32.179 | בו |  |
| 4:32.180 | מתפשט |  |
| 4:32.181 | ונמשך |  |
| 4:32.182 | מימי |  |
| 4:32.183 | האור |  |
| 4:32.184 | העליון |  |
| 4:32.185 | של |  |
| 4:32.186 | אין |  |
| 4:32.187 | סוף |  |
| 4:32.188 | אל |  |
| 4:32.189 | כל |  |
| 4:32.190 | העולמות |  |
| 4:32.191 | אשר |  |
| 4:32.192 | במקום |  |
| 4:32.193 | האויר |  |
| 4:32.194 | והחלל |  |
| 4:32.195 | ההוא | , |
| 4:32.196 | ונבאר |  |
| 4:32.197 | עתה |  |
| 4:32.198 | ענין |  |
| 4:32.199 | חקירת |  |
| 4:32.200 | המקובלים |  |
| 4:32.201 | לדעת |  |
| 4:32.202 | איך |  |
| 4:32.203 | יש |  |
| 4:32.204 | ראש |  |
| 4:32.205 | וסוף |  |
| 4:32.206 | בספירות | , |
| 4:32.207 | אמנם |  |
| 4:32.208 | בהיות |  |
| 4:32.209 | הקו |  |
| 4:32.210 | ההוא |  |
| 4:32.211 | ראשו |  |
| 4:32.212 | נוגע |  |
| 4:32.213 | באור |  |
| 4:32.214 | אין |  |
| 4:32.215 | סוף |  |
| 4:32.216 | מצד |  |
| 4:32.217 | העליון | , |
| 4:32.218 | וסופו |  |
| 4:32.219 | אינו |  |
| 4:32.220 | נמשך |  |
| 4:32.221 | למטה |  |
| 4:32.222 | עד |  |
| 4:32.223 | מקום |  |
| 4:32.224 | אין |  |
| 4:32.225 | סוף |  |
| 4:32.226 | הסובב |  |
| 4:32.227 | תחת |  |
| 4:32.228 | העולמות |  |
| 4:32.229 | ואינו |  |
| 4:32.230 | דבוק |  |
| 4:32.231 | בו | , |
| 4:32.232 | לכן |  |
| 4:32.233 | אז |  |
| 4:32.234 | יצדק |  |
| 4:32.235 | בו |  |
| 4:32.236 | ראש |  |
| 4:32.237 | וסוף | , |
| 4:32.238 | כי |  |
| 4:32.239 | אם |  |
| 4:32.240 | היו |  |
| 4:32.241 | דרך |  |
| 4:32.242 | שני |  |
| 4:32.243 | קצוות |  |
| 4:32.244 | מקבלים |  |
| 4:32.245 | שפע |  |
| 4:32.246 | האין |  |
| 4:32.247 | סוף |  |
| 4:32.248 | היו |  |
| 4:32.249 | שני |  |
| 4:32.250 | הקצוות |  |
| 4:32.251 | בחינת |  |
| 4:32.252 | ראשים |  |
| 4:32.253 | שווים |  |
| 4:32.254 | זה |  |
| 4:32.255 | לזה |  |
| 4:32.256 | ולא |  |
| 4:32.257 | היה |  |
| 4:32.258 | אז |  |
| 4:32.259 | בחינת |  |
| 4:32.260 | מעלה |  |
| 4:32.261 | ומטה | , |
| 4:32.262 | וכן |  |
| 4:32.263 | אם |  |
| 4:32.264 | היה |  |
| 4:32.265 | האין |  |
| 4:32.266 | סוף |  |
| 4:32.267 | נמשך |  |
| 4:32.268 | מכל |  |
| 4:32.269 | צדי |  |
| 4:32.270 | המקום |  |
| 4:32.271 | החלל |  |
| 4:32.272 | ההוא | , |
| 4:32.273 | לא |  |
| 4:32.274 | היה |  |
| 4:32.275 | לא |  |
| 4:32.276 | מעלה |  |
| 4:32.277 | ולא |  |
| 4:32.278 | מטה |  |
| 4:32.279 | לא |  |
| 4:32.280 | פנים |  |
| 4:32.281 | ולא |  |
| 4:32.282 | אחור | . |
| 4:32.283 | לא |  |
| 4:32.284 | מזרח |  |
| 4:32.285 | ולא |  |
| 4:32.286 | מערב |  |
| 4:32.287 | לא |  |
| 4:32.288 | צפון |  |
| 4:32.289 | ולא |  |
| 4:32.290 | דרום | , |
| 4:32.291 | אך |  |
| 4:32.292 | בהיות |  |
| 4:32.293 | אור |  |
| 4:32.294 | האין |  |
| 4:32.295 | סוף |  |
| 4:32.296 | נמשך |  |
| 4:32.297 | דרך |  |
| 4:32.298 | קו |  |
| 4:32.299 | אחד |  |
| 4:32.300 | וצנור |  |
| 4:32.301 | דק |  |
| 4:32.302 | בלבד |  |
| 4:32.303 | יצדק |  |
| 4:32.304 | מעלה |  |
| 4:32.305 | ומטה |  |
| 4:32.306 | פנים |  |
| 4:32.307 | ואחור |  |
| 4:32.308 | מזרח |  |
| 4:32.309 | ומערב |  |
| 4:32.310 | דבריו |  |
| 4:32.311 | אלה |  |
| 4:32.312 | לא |  |
| 4:32.313 | בהשכל |  |
| 4:32.314 | דבשלמא |  |
| 4:32.315 | ראש |  |
| 4:32.316 | וסוף |  |
| 4:32.317 | שפיר |  |
| 4:32.318 | אך |  |
| 4:32.319 | פנים |  |
| 4:32.320 | ואחור |  |
| 4:32.321 | מזרח |  |
| 4:32.322 | ומערב |  |
| 4:32.323 | איך |  |
| 4:32.324 | יצדק | ?) |
| 4:32.325 | והנה |  |
| 4:32.326 | בהיות |  |
| 4:32.327 | אור |  |
| 4:32.328 | האין |  |
| 4:32.329 | סוף |  |
| 4:32.330 | נמשך |  |
| 4:32.331 | בבחינת |  |
| 4:32.332 | קו |  |
| 4:32.333 | ישר |  |
| 4:32.334 | תוך |  |
| 4:32.335 | החלל | , |
| 4:32.336 | לא |  |
| 4:32.337 | נמשך |  |
| 4:32.338 | ונתפשט |  |
| 4:32.339 | למטה |  |
| 4:32.340 | תיכף | , |
| 4:32.341 | אלא |  |
| 4:32.342 | היה |  |
| 4:32.343 | מתפשט |  |
| 4:32.344 | לאט |  |
| 4:32.345 | לאט | . |
| 4:32.346 | רצוני |  |
| 4:32.347 | לומר |  |
| 4:32.348 | כי |  |
| 4:32.349 | בתחלה |  |
| 4:32.350 | התחיל |  |
| 4:32.351 | קו |  |
| 4:32.352 | האור |  |
| 4:32.353 | להתפשט |  |
| 4:32.354 | שם | , |
| 4:32.355 | ותכף |  |
| 4:32.356 | בתחלת |  |
| 4:32.357 | התפשטותו |  |
| 4:32.358 | בסוד |  |
| 4:32.359 | קו |  |
| 4:32.360 | נתפשט |  |
| 4:32.361 | ונמשך |  |
| 4:32.362 | ונעשה |  |
| 4:32.363 | כעין |  |
| 4:32.364 | גלגל |  |
| 4:32.365 | עגול |  |
| 4:32.366 | מסביב | . |
| 4:32.367 | והעגול |  |
| 4:32.368 | הזה |  |
| 4:32.369 | היה |  |
| 4:32.370 | בלתי |  |
| 4:32.371 | דבוק |  |
| 4:32.372 | אל |  |
| 4:32.373 | האין |  |
| 4:32.374 | סוף |  |
| 4:32.375 | הסובב |  |
| 4:32.376 | עליו |  |
| 4:32.377 | מכל |  |
| 4:32.378 | צדדיו | , |
| 4:32.379 | שאם |  |
| 4:32.380 | יתדבק |  |
| 4:32.381 | בו |  |
| 4:32.382 | יחזור |  |
| 4:32.383 | הדבר |  |
| 4:32.384 | לכמות |  |
| 4:32.385 | שהיה | , |
| 4:32.386 | ויתבטל |  |
| 4:32.387 | באור |  |
| 4:32.388 | אין |  |
| 4:32.389 | סוף |  |
| 4:32.390 | ולא |  |
| 4:32.391 | יתראה |  |
| 4:32.392 | כחו |  |
| 4:32.393 | כלל | , |
| 4:32.394 | ויהיה |  |
| 4:32.395 | הכל |  |
| 4:32.396 | אור |  |
| 4:32.397 | אין |  |
| 4:32.398 | סוף |  |
| 4:32.399 | לבד |  |
| 4:32.400 | כבראשונה | , |
| 4:32.401 | מדבריו |  |
| 4:32.402 | אלה |  |
| 4:32.403 | יש |  |
| 4:32.404 | להקשות |  |
| 4:32.405 | על |  |
| 4:32.406 | דבריו |  |
| 4:32.407 | שיזכיר |  |
| 4:32.408 | להלן |  |
| 4:32.409 | כי |  |
| 4:32.410 | דרך |  |
| 4:32.411 | הקו |  |
| 4:32.412 | הזה |  |
| 4:32.413 | יורד |  |
| 4:32.414 | אור |  |
| 4:32.415 | פנימי |  |
| 4:32.416 | שהוא |  |
| 4:32.417 | נשמת |  |
| 4:32.418 | הספירות |  |
| 4:32.419 | ואור |  |
| 4:32.420 | מקיף |  |
| 4:32.421 | שהוא |  |
| 4:32.422 | מלבוש |  |
| 4:32.423 | לספירות | . |
| 4:32.424 | למה |  |
| 4:32.425 | אינו |  |
| 4:32.426 | מתערב |  |
| 4:32.427 | ומתדבק |  |
| 4:32.428 | בגוף |  |
| 4:32.429 | הספירות |  |
| 4:32.430 | ויחזור |  |
| 4:32.431 | עצם |  |
| 4:32.432 | הספירות |  |
| 4:32.433 | עם |  |
| 4:32.434 | מלבושן |  |
| 4:32.435 | ונשמתן |  |
| 4:32.436 | כולם |  |
| 4:32.437 | גולם |  |
| 4:32.438 | אחד |  |
| 4:32.439 | ויתבטל |  |
| 4:32.440 | הכלי |  |
| 4:32.441 | ויתערב |  |
| 4:32.442 | באור |  |
| 4:32.443 | פנימי |  |
| 4:32.444 | ואור |  |
| 4:32.445 | חיצון | ?) |
| 4:32.446 | לכן |  |
| 4:32.447 | העגול |  |
| 4:32.448 | הזה |  |
| 4:32.449 | סמוך |  |
| 4:32.450 | אל |  |
| 4:32.451 | האין |  |
| 4:32.452 | סוף |  |
| 4:32.453 | ובלתי |  |
| 4:32.454 | מתדבק |  |
| 4:32.455 | בו | . |
| 4:32.456 | וכל |  |
| 4:32.457 | עיקר |  |
| 4:32.458 | התקשרות |  |
| 4:32.459 | ודביקות |  |
| 4:32.460 | הנאצל |  |
| 4:32.461 | ההוא |  |
| 4:32.462 | עם |  |
| 4:32.463 | אין |  |
| 4:32.464 | סוף |  |
| 4:32.465 | המאציל |  |
| 4:32.466 | הוא |  |
| 4:32.467 | ע"י |  |
| 4:32.468 | הקו |  |
| 4:32.469 | ההוא |  |
| 4:32.470 | אשר |  |
| 4:32.471 | דרך |  |
| 4:32.472 | בו |  |
| 4:32.473 | יורד |  |
| 4:32.474 | ונמשך |  |
| 4:32.475 | אור |  |
| 4:32.476 | מן |  |
| 4:32.477 | האין |  |
| 4:32.478 | סוף |  |
| 4:32.479 | המשפיע |  |
| 4:32.480 | בעיגול |  |
| 4:32.481 | ההוא | . |
| 4:32.482 | והאין |  |
| 4:32.483 | סוף |  |
| 4:32.484 | סובב |  |
| 4:32.485 | ומקיף |  |
| 4:32.486 | עליו |  |
| 4:32.487 | מכל |  |
| 4:32.488 | צדדיו |  |
| 4:32.489 | ורחוק |  |
| 4:32.490 | ממנו |  |
| 4:32.491 | בהשואה |  |
| 4:32.492 | אחת |  |
| 4:32.493 | מכל |  |
| 4:32.494 | צדדיו | . |
| 4:32.495 | כי |  |
| 4:32.496 | גם |  |
| 4:32.497 | הוא |  |
| 4:32.498 | בבחינת |  |
| 4:32.499 | עיגול |  |
| 4:32.500 | סביב |  |
| 4:32.501 | עליו | . |
| 4:32.502 | כי |  |
| 4:32.503 | הוא |  |
| 4:32.504 | מוכרח |  |
| 4:32.505 | שהארת |  |
| 4:32.506 | א"ס |  |
| 4:32.507 | בנאצלים |  |
| 4:32.508 | תהיה |  |
| 4:32.509 | דרך |  |
| 4:32.510 | הקו |  |
| 4:32.511 | ההוא |  |
| 4:32.512 | לבד | . |
| 4:32.513 | כי |  |
| 4:32.514 | אם |  |
| 4:32.515 | היה |  |
| 4:32.516 | האור |  |
| 4:32.517 | ההוא |  |
| 4:32.518 | נמשך |  |
| 4:32.519 | להם |  |
| 4:32.520 | גם |  |
| 4:32.521 | מכל |  |
| 4:32.522 | סביבותיהם |  |
| 4:32.523 | היו |  |
| 4:32.524 | הנאצלים |  |
| 4:32.525 | בבחינת |  |
| 4:32.526 | המאציל |  |
| 4:32.527 | עצמו |  |
| 4:32.528 | בלתי |  |
| 4:32.529 | גבול |  |
| 4:32.530 | וקצבה | , |
| 4:32.531 | ולא |  |
| 4:32.532 | עוד |  |
| 4:32.533 | אלא |  |
| 4:32.534 | אפילו |  |
| 4:32.535 | הקו |  |
| 4:32.536 | ההוא |  |
| 4:32.537 | הוא |  |
| 4:32.538 | דק |  |
| 4:32.539 | מאד |  |
| 4:32.540 | ולא |  |
| 4:32.541 | עבה |  |
| 4:32.542 | כדי |  |
| 4:32.543 | שיהיה |  |
| 4:32.544 | האור |  |
| 4:32.545 | ההוא |  |
| 4:32.546 | הנמשך |  |
| 4:32.547 | אל |  |
| 4:32.548 | הנאצלים |  |
| 4:32.549 | במדה |  |
| 4:32.550 | וקצבה |  |
| 4:32.551 | אשר |  |
| 4:32.552 | לסבה |  |
| 4:32.553 | זאת |  |
| 4:32.554 | נקראים |  |
| 4:32.555 | הנאצלים |  |
| 4:32.556 | עשר |  |
| 4:32.557 | מדות |  |
| 4:32.558 | ועשר |  |
| 4:32.559 | ספירות |  |
| 4:32.560 | לפי |  |
| 4:32.561 | שיש |  |
| 4:32.562 | להם |  |
| 4:32.563 | מדה |  |
| 4:32.564 | וקצבה |  |
| 4:32.565 | ומספר |  |
| 4:32.566 | קצוב | . |
| 4:32.567 | והנה |  |
| 4:32.568 | העיגול |  |
| 4:32.569 | הראשון |  |
| 4:32.570 | הזה |  |
| 4:32.571 | היותר |  |
| 4:32.572 | דבוק |  |
| 4:32.573 | עם |  |
| 4:32.574 | האין |  |
| 4:32.575 | סוף |  |
| 4:32.576 | הוא |  |
| 4:32.577 | הנקרא |  |
| 4:32.578 | כתר |  |
| 4:32.579 | דאדם |  |
| 4:32.580 | קדמון | . |
| 4:32.581 | ואחר |  |
| 4:32.582 | כך |  |
| 4:32.583 | נתפשט |  |
| 4:32.584 | הקו |  |
| 4:32.585 | הזה |  |
| 4:32.586 | ונמשך |  |
| 4:32.587 | מעט |  |
| 4:32.588 | וחזר |  |
| 4:32.589 | להתעגל |  |
| 4:32.590 | ונעשה |  |
| 4:32.591 | עיגול |  |
| 4:32.592 | שני |  |
| 4:32.593 | תוך |  |
| 4:32.594 | עיגול |  |
| 4:32.595 | ראשון | . |
| 4:32.596 | וזה |  |
| 4:32.597 | נקרא |  |
| 4:32.598 | עיגול |  |
| 4:32.599 | החכמה |  |
| 4:32.600 | דאדם |  |
| 4:32.601 | קדמון | , |
| 4:32.602 | עוד |  |
| 4:32.603 | נתפשט |  |
| 4:32.604 | יותר |  |
| 4:32.605 | למטה |  |
| 4:32.606 | וחזר |  |
| 4:32.607 | להתעגל |  |
| 4:32.608 | ונעשה |  |
| 4:32.609 | עיגול |  |
| 4:32.610 | שלישי |  |
| 4:32.611 | תוך |  |
| 4:32.612 | עיגול |  |
| 4:32.613 | שני |  |
| 4:32.614 | הזה |  |
| 4:32.615 | ונקרא |  |
| 4:32.616 | עיגול |  |
| 4:32.617 | בינה |  |
| 4:32.618 | דאדם |  |
| 4:32.619 | קדמון | , |
| 4:32.620 | ועל |  |
| 4:32.621 | דרך |  |
| 4:32.622 | זה |  |
| 4:32.623 | היה |  |
| 4:32.624 | הולך |  |
| 4:32.625 | ומתעגל |  |
| 4:32.626 | עד |  |
| 4:32.627 | עיגול |  |
| 4:32.628 | העשירי |  |
| 4:32.629 | הנקרא |  |
| 4:32.630 | עיגול |  |
| 4:32.631 | מלכות |  |
| 4:32.632 | דאדם |  |
| 4:32.633 | קדמון |  |
| 4:32.634 | וכו | '. |
| 4:32.635 | אמנם |  |
| 4:32.636 | מבואר |  |
| 4:32.637 | ופשוט |  |
| 4:32.638 | הוא |  |
| 4:32.639 | שכמה |  |
| 4:32.640 | מיני |  |
| 4:32.641 | עולמות |  |
| 4:32.642 | נאצלו |  |
| 4:32.643 | ונבראו |  |
| 4:32.644 | ונוצרו |  |
| 4:32.645 | ונעשו |  |
| 4:32.646 | אלף |  |
| 4:32.647 | אלפים |  |
| 4:32.648 | ורבוא |  |
| 4:32.649 | רבבות | , |
| 4:32.650 | וכולם |  |
| 4:32.651 | כאחד |  |
| 4:32.652 | תוך |  |
| 4:32.653 | המקום |  |
| 4:32.654 | החלל |  |
| 4:32.655 | הנז"ל |  |
| 4:32.656 | ואין |  |
| 4:32.657 | דבר |  |
| 4:32.658 | חוצה |  |
| 4:32.659 | לו | . |
| 4:32.660 | וכל |  |
| 4:32.661 | עולם |  |
| 4:32.662 | ועולם |  |
| 4:32.663 | יש |  |
| 4:32.664 | בו |  |
| 4:32.665 | עשר |  |
| 4:32.666 | ספירות |  |
| 4:32.667 | פרטיות | . |
| 4:32.668 | וכל |  |
| 4:32.669 | ספירה |  |
| 4:32.670 | וספירה |  |
| 4:32.671 | שבכל |  |
| 4:32.672 | עולם |  |
| 4:32.673 | ועולם |  |
| 4:32.674 | כלולה |  |
| 4:32.675 | מעשר |  |
| 4:32.676 | ספירות |  |
| 4:32.677 | פרטי |  |
| 4:32.678 | פרטות | . |
| 4:32.679 | ועתה |  |
| 4:32.680 | נבאר |  |
| 4:32.681 | בחינה |  |
| 4:32.682 | שנית |  |
| 4:32.683 | שיש |  |
| 4:32.684 | בעשר |  |
| 4:32.685 | ספירות | , |
| 4:32.686 | היא |  |
| 4:32.687 | בחינת |  |
| 4:32.688 | היושר |  |
| 4:32.689 | בצורת |  |
| 4:32.690 | אדם |  |
| 4:32.691 | העליון | . |
| 4:32.692 | הנה |  |
| 4:32.693 | דרך |  |
| 4:32.694 | הקו |  |
| 4:32.695 | הנז | ' |
| 4:32.696 | לעיל |  |
| 4:32.697 | המתפשט |  |
| 4:32.698 | מלמעלה |  |
| 4:32.699 | למטה |  |
| 4:32.700 | אשר |  |
| 4:32.701 | ממנו |  |
| 4:32.702 | מתפשטים |  |
| 4:32.703 | העיגולים | , |
| 4:32.704 | גם |  |
| 4:32.705 | הקו |  |
| 4:32.706 | ההוא |  |
| 4:32.707 | מתפשט |  |
| 4:32.708 | ביושר |  |
| 4:32.709 | ממעלה |  |
| 4:32.710 | למטה |  |
| 4:32.711 | מראש |  |
| 4:32.712 | גג |  |
| 4:32.713 | העליון |  |
| 4:32.714 | שלעיגול |  |
| 4:32.715 | העליון |  |
| 4:32.716 | עד |  |
| 4:32.717 | למטה |  |
| 4:32.718 | מתחתית |  |
| 4:32.719 | סיום |  |
| 4:32.720 | כל |  |
| 4:32.721 | העיגולים |  |
| 4:32.722 | ממש |  |
| 4:32.723 | מלמעלה |  |
| 4:32.724 | למטה |  |
| 4:32.725 | כלול |  |
| 4:32.726 | מעשר |  |
| 4:32.727 | ספירות |  |
| 4:32.728 | בסוד |  |
| 4:32.729 | צלם |  |
| 4:32.730 | אדם |  |
| 4:32.731 | ישר |  |
| 4:32.732 | בעל |  |
| 4:32.733 | קומה |  |
| 4:32.734 | זקופה |  |
| 4:32.735 | כלול |  |
| 4:32.736 | מרמ"ח |  |
| 4:32.737 | איברים |  |
| 4:32.738 | וכו | '. |
| 4:32.739 | והנה |  |
| 4:32.740 | בחינה |  |
| 4:32.741 | זו |  |
| 4:32.742 | השנית |  |
| 4:32.743 | נקראת |  |
| 4:32.744 | צלם |  |
| 4:32.745 | אלהים |  |
| 4:32.746 | ועליה |  |
| 4:32.747 | רמז |  |
| 4:32.748 | הכתוב |  |
| 4:32.749 | באמרו |  |
| 4:32.750 | ויברא |  |
| 4:32.751 | אלהים |  |
| 4:32.752 | את |  |
| 4:32.753 | האדם |  |
| 4:32.754 | בצלמו |  |
| 4:32.755 | בצלם |  |
| 4:32.756 | אלהים |  |
| 4:32.757 | ברא |  |
| 4:32.758 | אותו |  |
| 4:32.759 | וכו | ' |
| 4:32.760 | ושכמעט |  |
| 4:32.761 | רוב |  |
| 4:32.762 | דברי |  |
| 4:32.763 | הזהר |  |
| 4:32.764 | והתיקונים |  |
| 4:32.765 | מתעסקים |  |
| 4:32.766 | בבחינה |  |
| 4:32.767 | הזאת |  |
| 4:32.768 | השנית |  |
| 4:32.769 | שהיא |  |
| 4:32.770 | בחינת |  |
| 4:32.771 | היושר |  |
| 4:32.772 | עכ"ל | . |
| 4:32.773 | ובענף |  |
| 4:32.774 | ג | ' |
| 4:32.775 | וד | ' |
| 4:32.776 | כפל |  |
| 4:32.777 | עוד |  |
| 4:32.778 | הדברים |  |
| 4:32.779 | שאין |  |
| 4:32.780 | העשר |  |
| 4:32.781 | ספירות |  |
| 4:32.782 | דעולם |  |
| 4:32.783 | האצילות |  |
| 4:32.784 | הם |  |
| 4:32.785 | ראשונים |  |
| 4:32.786 | וגבוהים |  |
| 4:32.787 | מכל |  |
| 4:32.788 | מה |  |
| 4:32.789 | שנאצלו | . |
| 4:32.790 | כי |  |
| 4:32.791 | כמה |  |
| 4:32.792 | עולמות |  |
| 4:32.793 | אבי"ע |  |
| 4:32.794 | קדמו |  |
| 4:32.795 | להם | . |
| 4:32.796 | ולרוב |  |
| 4:32.797 | העלמם |  |
| 4:32.798 | לא |  |
| 4:32.799 | שלחו |  |
| 4:32.800 | בהם |  |
| 4:32.801 | יד |  |
| 4:32.802 | להזכירם |  |
| 4:32.803 | בזהר |  |
| 4:32.804 | ובתיקונים |  |
| 4:32.805 | אלא |  |
| 4:32.806 | ברמז |  |
| 4:32.807 | נפלא | . |
| 4:32.808 | ושם |  |
| 4:32.809 | ביאר |  |
| 4:32.810 | עוד |  |
| 4:32.811 | עד |  |
| 4:32.812 | היכן |  |
| 4:32.813 | מגיעים |  |
| 4:32.814 | רגלי |  |
| 4:32.815 | אדם |  |
| 4:32.816 | קדמון |  |
| 4:32.817 | דיושר | . |
| 4:32.818 | ועד |  |
| 4:32.819 | היכן |  |
| 4:32.820 | מגיעים |  |
| 4:32.821 | רגלי |  |
| 4:32.822 | עתיק |  |
| 4:32.823 | יומין | , |
| 4:32.824 | ורגלי |  |
| 4:32.825 | אריך |  |
| 4:32.826 | אנפין | , |
| 4:32.827 | ורגלי |  |
| 4:32.828 | אבא |  |
| 4:32.829 | ואימא | , |
| 4:32.830 | ורגלי |  |
| 4:32.831 | נוקבא |  |
| 4:32.832 | דיליה | . |
| 4:32.833 | ושאבא |  |
| 4:32.834 | ואימא |  |
| 4:32.835 | קצרי |  |
| 4:32.836 | קומה | , |
| 4:32.837 | ואין |  |
| 4:32.838 | שיעור |  |
| 4:32.839 | קומתם |  |
| 4:32.840 | רק |  |
| 4:32.841 | מהגרון |  |
| 4:32.842 | דאריך |  |
| 4:32.843 | אנפין |  |
| 4:32.844 | עד |  |
| 4:32.845 | הטבור |  |
| 4:32.846 | שלאריך |  |
| 4:32.847 | אנפין | , |
| 4:32.848 | וכדברים |  |
| 4:32.849 | האלה |  |
| 4:32.850 | מבואר |  |
| 4:32.851 | ג"כ |  |
| 4:32.852 | בשער |  |
| 4:32.853 | ההקדמות |  |
| 4:32.854 | ובספר |  |
| 4:32.855 | כסא |  |
| 4:32.856 | אליהו |  |
| 4:32.857 | ומקדש |  |
| 4:32.858 | מלך |  |
| 4:32.859 | פרשת |  |
| 4:32.860 | בראשית |  |
| 4:32.861 | עיין |  |
| 4:32.862 | שם |  |
| 4:32.863 | ושם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

