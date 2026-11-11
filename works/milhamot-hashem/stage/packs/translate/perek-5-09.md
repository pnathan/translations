# Translation pack: milhamot-hashem perek-5-09 (Milchamot ha-Shem 5:15–5:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-09`.

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
- **יחוד / ייחוד** (both spellings, one rendering):
  - The doctrine: fixed, "unity". "ייחוד השי"ת" → "the unity of the Name,
    may He be blessed"; "אמונת היחוד" → "the belief in the unity"; "שער
    היחוד" → "the Gate of Unity". Not "monotheism" and not "oneness".
    אחדות → "oneness", kept distinct.
  - The Kabbalistic act: "unification". "לשם יחוד קב"ה ושכינתיה" (6:21) →
    "for the sake of the unification of the Holy One, blessed be He, and His
    Shekhinah"; יחודים → "unifications".
  - The verb: ליחד with God as object, in the rabbinic sense → "to declare
    one": "הזהירנו ליחדו" (5:26) → "commanded us to declare Him one".
    "מיחדים את השי"ת" → "declare the Name, may He be blessed, one". With the
    Kabbalistic entities as object → "to unify": "ליחד זוגתו עמו" (7:5) →
    "to unify his consort with him".
  - The author plays on the two senses ("עיקר היחוד אצלם לחבר ולקשר", 6:1;
    "סוד היחוד הוא ליחד את השם בכל עשר ספירות", 7:25). There, keep "unity" for
    the noun and "unify" for the verb, so that the shared stem shows, and
    explain the play in a note.
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
- מעמד: "assembly" or "occasion" by context; "the Standing at Sinai" where the
  author names that event.
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


## Milchamot ha-Shem 5:15

Hebrew (Wikisource): מז) והנה שכח משנה שלמה ששנינו באבות (ד' כ"ב): "לידע ולהודיע ולהודע, שהוא אל, הוא הבורא, הוא היוצר, הוא המבין, הוא הדיין, הוא עד, הוא בעל דין, והוא עתיד לדון, ברוך הוא שאין לפניו לא עולה לא שכחה, לא משוא פנים ולא מקח שוחד", וארז"ל (ירושלמי שקלים פ"ה ה"א, מ"ח ע"ד): "כל האומר הקב"ה ותרן יותרון מעוהי, אלא מאריך רוחיה וגבי דיליה", ואמרו עוד (שם) "על וסביביו נשערה מאוד" שהקב"ה מדקדק עם חסידיו כחוט השערה, ואמרו עוד (שבועות ל"ט ע"א): "ונקה לא ינקה, מנקה הוא לשבים ואינו מנקה לשאינן שבים", וכן תרגם אונקלוס מפי ר' אליעזר ור' יהושע (שמות ל"ד ז'): "סלח לדיתיבין לאוריתיה, ולדלא תיבין לא מזכי" וכו'. ואל כל אלה האמונות הזרות לתורתינו הקדושה ירמזון דברי ר' תם ן' יחיא המובא בספר רביד הזהב שהבאתי לעיל סי' י"ג שכתב שהעוסקים בקבלה החדשה הם הורסים פינות התורה ועמודיה יתפלצון וכו'. וגם ירמזון אל מה שכתוב בכוונת האר"י בספר מחברת הקדש דף כ"ח ע"ב, ובספר חמדת ימים סוף פרק ז' מתקוני שבת שלא לאמר יגדל אלהים חי וישתבח וכו' מפני שסוברים כי שום עבודה ותפלה ושבח לא שייכי באלהינו ב"ה שהוא הסיבה הראשונה בהיותו מופשט מחומר וצורה, ונפתו בזה אחר עובדי האלילים שאמרו שהוא יתברך נעדר ומרומם ומסולק מכל ברכה ותהלה וכל תורה ועבודה ותפלה לא שייכי ביה. וגם בהתלבשו בספירות שבעולמות העליונים אין מספר למעלה מעולם האצילות (לפי דבריהם) שהם דקים וקלושים ואינם מושגים. וגם באדם קדמון דאצילות ועתיק יומין ואריך אנפין ואבא ואימא לא שייכי שום עבודה וקריאה ותפלה לרוב דקותם והעלמם. ורק בפרצוף אחרון דאצילות שהוא זעיר אנפין ונוקביה שכבר נתעבו קצת שייך בהו עבודה ותפלה. ודברים אלו מתבארים מדברי הרב בעל ספר הברית מאמר הנזכר לעיל והרב כסא אליהו והרח"ו בספר עץ חיים והרב בעל המצרף ונחלת יוסף וזולתם. וזה לשון ספר הברית: "ובזאת תדע אחי שאמונתינו אינה כאמונת הפילוסופים (האלהיים) והישמעאלים בענין היחוד והבורא. כי הם אין איתם יודע את השם הנכבד והנורא הוי'ה ב"ה. רק מאמינים במחוייב המציאות סתם בבחינתו הפשוטה כאשר בתחלה קודם הבריאה (ערום בלי לבוש) וזהו בחינתו הבלתי מושג אל זולתו בשום רמז ורמיזה. ובבחינה הזאת הטיבו אשר דיברו הפילוסופים (הכופרים) שלא שייך אליו שום עבודה ותפלה על כי הוא בבחינתו זאת נעלה ומרומם מכל ברכה ותהלה ובטלה כל תורה ומצוה. כמו שכתב ר' מאיר גבאי. והמתבונן בזה היטב יתבאר לו בלי שום ספק שאין מקום לכל זה בלתי אצילות הספירות. לא כן עם אלהי אברהם שמאמינים במחוייב המציאות מלובש במידותיו הוא סוד הוי"ה ב"ה שהוא הבחינה שאחר הבריאה, בחינה המושגת לבית יעקב שנגלה למשה בסנה והוציאנו מארץ מצרים ונתן תורת אמת לאבותינו דור המדבר פנים בפנים, שייך אליו העבודה והקרבנות והתפלה וכל המצוות המדובר בתוה"ק. ועליו הזהיר על (זעיר אנפין) משה ואמר ליראה את השם הנכבד והנורא הזה ה' את אלהיך" עכ"ל, וכן כתב מהרח"ו בעץ חיים שער טנת"א סוף פרק ב' שהאור היוצא מתוך אדם קדמון הוא זך מאוד, ועל ידי ירידתו והתרחקותו הוא מתעבה עיבוי אחר עיבוי, כיצד האור הנמשך ויוצא דרך האוזן הוא זך מאוד וכאשר נמשך האור הזה עד החוטם ויוצא קצת דרך שם הוא מתעבה וקונה איזה עביות וגסות מחמת הריחוק שנתרחק ונמשך יותר למטה בצאתו עד הפה ויוצא קצתו דרך שם מתעבה יותר בצאתו משם וכו' וכן כתב בשער הנקודים פרק א' שכל מה שהאור יורד למטה הוא ניכר ונרגש ומתגלה שם עכ"ל. וזהו טעמם במה שאמרו דלא שייך עבודה ותפלה רק לאלוה הנרגש ויש בו עובי וגסות קצת שתוכל לציירו בדמות תבניתו על ידי שנתעבה ונתגשם. אבל האין סוף או אדם קדמאה ואדם קדמון ואפילו עתיק ואריך שעדיין לא נתעבו ונתגשמו כל כך, לא שייך בהן עבודה, כי עדיין לא נתעבו ולא נתגשמו שתוכל לציירם, ורק העבודה והיראה לאלוה קצר האפים (זעיר אנפין) והוא לבדו לפי דעתם אלהינו על ידי שנתגשם ונגלה למשה בסנה ובסיני לאבותינו. והנה שכחו מקרא מלא כי מלאך הוא שנגלה למשה בסנה שנאמר (שמות ג' ב'): "וירא מלאך ה' אליו בלבת אש מתוך הסנה". ובמעמד הר סיני הלא נאמר (דברים ד' ט"ו): "כי לא ראיתם כל תמונה ביום דבר ה' עמכם", וכו' לא ראו דמות המושגת ברמ"ח איברים ושס"ה גידים כדבריהם. וקרוב לדבריהם כתב הרב עז לאלהים בבית קדש הקדשים דף כ"ו ע"ד וזה לשונו: "מי שמאמין שהאין סוף הוא עצם האלוהות כאילו כופר בעיקר. ועליהם אמר הנביא (ישעיהו כ"ט ט"ו): 'הוי המעמיקים מה' לסתיר עצה'. ועוד שעצם האלוהות הוא שם הוי"ה, וכל דבר שיש לו שם ואות יש בו שיעור קומה וכו'. ושם הדבר הוא גבולו שכן שם האדם לא שייך שיקרא בשם אדם אלא מראשו ועד רגליו, הרי שהשם הוא גבולו. לכן כל מה שתמצא משמות דשייכי באדם שייכים אליו יתברך, ואינו כסברת הטועים שאומרים שעצם האלוהות הוא אין סוף"! וחולק הוא על מהרח"ו ושאר מקובלים שסוברים שהתפארת שהוא זעיר אנפין הוא אלהינו, וסבירא ליה שהעבודה והתפלה למלכא קדישא דכל קדישין שהוא מאיר בלב זעיר אנפין כמו שכתב בפרק י"ט דף נ"ט ע"ד, וזה לשונו בדף כ"ו סוף ע"ד: "בר מן דין ובר מן דין אם איתא שהתפארת הוא עצם האלהות למה נעבוד את התפארת שהוא מקבל מאבא ואמא, נעבוד לאריך אנפין או לכתר שהוא יותר למעלה? ואם האין סוף הוא עצם האלוהות, איך אנו אומרים ברוך אתה ה', ממי יתברך ויביא לו שפע"? עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:15.1 | מז | ) |
| 5:15.2 | והנה |  |
| 5:15.3 | שכח |  |
| 5:15.4 | משנה |  |
| 5:15.5 | שלמה |  |
| 5:15.6 | ששנינו |  |
| 5:15.7 | באבות |  |
| 5:15.8 | ד | ' |
| 5:15.9 | כ"ב | ): |
| 5:15.10 | לידע |  |
| 5:15.11 | ולהודיע |  |
| 5:15.12 | ולהודע | , |
| 5:15.13 | שהוא |  |
| 5:15.14 | אל | , |
| 5:15.15 | הוא |  |
| 5:15.16 | הבורא | , |
| 5:15.17 | הוא |  |
| 5:15.18 | היוצר | , |
| 5:15.19 | הוא |  |
| 5:15.20 | המבין | , |
| 5:15.21 | הוא |  |
| 5:15.22 | הדיין | , |
| 5:15.23 | הוא |  |
| 5:15.24 | עד | , |
| 5:15.25 | הוא |  |
| 5:15.26 | בעל |  |
| 5:15.27 | דין | , |
| 5:15.28 | והוא |  |
| 5:15.29 | עתיד |  |
| 5:15.30 | לדון | , |
| 5:15.31 | ברוך |  |
| 5:15.32 | הוא |  |
| 5:15.33 | שאין |  |
| 5:15.34 | לפניו |  |
| 5:15.35 | לא |  |
| 5:15.36 | עולה |  |
| 5:15.37 | לא |  |
| 5:15.38 | שכחה | , |
| 5:15.39 | לא |  |
| 5:15.40 | משוא |  |
| 5:15.41 | פנים |  |
| 5:15.42 | ולא |  |
| 5:15.43 | מקח |  |
| 5:15.44 | שוחד | ", |
| 5:15.45 | וארז"ל |  |
| 5:15.46 | ירושלמי |  |
| 5:15.47 | שקלים |  |
| 5:15.48 | פ"ה |  |
| 5:15.49 | ה"א | , |
| 5:15.50 | מ"ח |  |
| 5:15.51 | ע"ד | ): |
| 5:15.52 | כל |  |
| 5:15.53 | האומר |  |
| 5:15.54 | הקב"ה |  |
| 5:15.55 | ותרן |  |
| 5:15.56 | יותרון |  |
| 5:15.57 | מעוהי | , |
| 5:15.58 | אלא |  |
| 5:15.59 | מאריך |  |
| 5:15.60 | רוחיה |  |
| 5:15.61 | וגבי |  |
| 5:15.62 | דיליה | ", |
| 5:15.63 | ואמרו |  |
| 5:15.64 | עוד |  |
| 5:15.65 | שם | ) |
| 5:15.66 | על |  |
| 5:15.67 | וסביביו |  |
| 5:15.68 | נשערה |  |
| 5:15.69 | מאוד | " |
| 5:15.70 | שהקב"ה |  |
| 5:15.71 | מדקדק |  |
| 5:15.72 | עם |  |
| 5:15.73 | חסידיו |  |
| 5:15.74 | כחוט |  |
| 5:15.75 | השערה | , |
| 5:15.76 | ואמרו |  |
| 5:15.77 | עוד |  |
| 5:15.78 | שבועות |  |
| 5:15.79 | ל"ט |  |
| 5:15.80 | ע"א | ): |
| 5:15.81 | ונקה |  |
| 5:15.82 | לא |  |
| 5:15.83 | ינקה | , |
| 5:15.84 | מנקה |  |
| 5:15.85 | הוא |  |
| 5:15.86 | לשבים |  |
| 5:15.87 | ואינו |  |
| 5:15.88 | מנקה |  |
| 5:15.89 | לשאינן |  |
| 5:15.90 | שבים | ", |
| 5:15.91 | וכן |  |
| 5:15.92 | תרגם |  |
| 5:15.93 | אונקלוס |  |
| 5:15.94 | מפי |  |
| 5:15.95 | ר | ' |
| 5:15.96 | אליעזר |  |
| 5:15.97 | ור | ' |
| 5:15.98 | יהושע |  |
| 5:15.99 | שמות |  |
| 5:15.100 | ל"ד |  |
| 5:15.101 | ז | '): |
| 5:15.102 | סלח |  |
| 5:15.103 | לדיתיבין |  |
| 5:15.104 | לאוריתיה | , |
| 5:15.105 | ולדלא |  |
| 5:15.106 | תיבין |  |
| 5:15.107 | לא |  |
| 5:15.108 | מזכי | " |
| 5:15.109 | וכו | '. |
| 5:15.110 | ואל |  |
| 5:15.111 | כל |  |
| 5:15.112 | אלה |  |
| 5:15.113 | האמונות |  |
| 5:15.114 | הזרות |  |
| 5:15.115 | לתורתינו |  |
| 5:15.116 | הקדושה |  |
| 5:15.117 | ירמזון |  |
| 5:15.118 | דברי |  |
| 5:15.119 | ר | ' |
| 5:15.120 | תם |  |
| 5:15.121 | ן | ' |
| 5:15.122 | יחיא |  |
| 5:15.123 | המובא |  |
| 5:15.124 | בספר |  |
| 5:15.125 | רביד |  |
| 5:15.126 | הזהב |  |
| 5:15.127 | שהבאתי |  |
| 5:15.128 | לעיל |  |
| 5:15.129 | סי | ' |
| 5:15.130 | י"ג |  |
| 5:15.131 | שכתב |  |
| 5:15.132 | שהעוסקים |  |
| 5:15.133 | בקבלה |  |
| 5:15.134 | החדשה |  |
| 5:15.135 | הם |  |
| 5:15.136 | הורסים |  |
| 5:15.137 | פינות |  |
| 5:15.138 | התורה |  |
| 5:15.139 | ועמודיה |  |
| 5:15.140 | יתפלצון |  |
| 5:15.141 | וכו | '. |
| 5:15.142 | וגם |  |
| 5:15.143 | ירמזון |  |
| 5:15.144 | אל |  |
| 5:15.145 | מה |  |
| 5:15.146 | שכתוב |  |
| 5:15.147 | בכוונת |  |
| 5:15.148 | האר"י |  |
| 5:15.149 | בספר |  |
| 5:15.150 | מחברת |  |
| 5:15.151 | הקדש |  |
| 5:15.152 | דף |  |
| 5:15.153 | כ"ח |  |
| 5:15.154 | ע"ב | , |
| 5:15.155 | ובספר |  |
| 5:15.156 | חמדת |  |
| 5:15.157 | ימים |  |
| 5:15.158 | סוף |  |
| 5:15.159 | פרק |  |
| 5:15.160 | ז | ' |
| 5:15.161 | מתקוני |  |
| 5:15.162 | שבת |  |
| 5:15.163 | שלא |  |
| 5:15.164 | לאמר |  |
| 5:15.165 | יגדל |  |
| 5:15.166 | אלהים |  |
| 5:15.167 | חי |  |
| 5:15.168 | וישתבח |  |
| 5:15.169 | וכו | ' |
| 5:15.170 | מפני |  |
| 5:15.171 | שסוברים |  |
| 5:15.172 | כי |  |
| 5:15.173 | שום |  |
| 5:15.174 | עבודה |  |
| 5:15.175 | ותפלה |  |
| 5:15.176 | ושבח |  |
| 5:15.177 | לא |  |
| 5:15.178 | שייכי |  |
| 5:15.179 | באלהינו |  |
| 5:15.180 | ב"ה |  |
| 5:15.181 | שהוא |  |
| 5:15.182 | הסיבה |  |
| 5:15.183 | הראשונה |  |
| 5:15.184 | בהיותו |  |
| 5:15.185 | מופשט |  |
| 5:15.186 | מחומר |  |
| 5:15.187 | וצורה | , |
| 5:15.188 | ונפתו |  |
| 5:15.189 | בזה |  |
| 5:15.190 | אחר |  |
| 5:15.191 | עובדי |  |
| 5:15.192 | האלילים |  |
| 5:15.193 | שאמרו |  |
| 5:15.194 | שהוא |  |
| 5:15.195 | יתברך |  |
| 5:15.196 | נעדר |  |
| 5:15.197 | ומרומם |  |
| 5:15.198 | ומסולק |  |
| 5:15.199 | מכל |  |
| 5:15.200 | ברכה |  |
| 5:15.201 | ותהלה |  |
| 5:15.202 | וכל |  |
| 5:15.203 | תורה |  |
| 5:15.204 | ועבודה |  |
| 5:15.205 | ותפלה |  |
| 5:15.206 | לא |  |
| 5:15.207 | שייכי |  |
| 5:15.208 | ביה | . |
| 5:15.209 | וגם |  |
| 5:15.210 | בהתלבשו |  |
| 5:15.211 | בספירות |  |
| 5:15.212 | שבעולמות |  |
| 5:15.213 | העליונים |  |
| 5:15.214 | אין |  |
| 5:15.215 | מספר |  |
| 5:15.216 | למעלה |  |
| 5:15.217 | מעולם |  |
| 5:15.218 | האצילות |  |
| 5:15.219 | לפי |  |
| 5:15.220 | דבריהם | ) |
| 5:15.221 | שהם |  |
| 5:15.222 | דקים |  |
| 5:15.223 | וקלושים |  |
| 5:15.224 | ואינם |  |
| 5:15.225 | מושגים | . |
| 5:15.226 | וגם |  |
| 5:15.227 | באדם |  |
| 5:15.228 | קדמון |  |
| 5:15.229 | דאצילות |  |
| 5:15.230 | ועתיק |  |
| 5:15.231 | יומין |  |
| 5:15.232 | ואריך |  |
| 5:15.233 | אנפין |  |
| 5:15.234 | ואבא |  |
| 5:15.235 | ואימא |  |
| 5:15.236 | לא |  |
| 5:15.237 | שייכי |  |
| 5:15.238 | שום |  |
| 5:15.239 | עבודה |  |
| 5:15.240 | וקריאה |  |
| 5:15.241 | ותפלה |  |
| 5:15.242 | לרוב |  |
| 5:15.243 | דקותם |  |
| 5:15.244 | והעלמם | . |
| 5:15.245 | ורק |  |
| 5:15.246 | בפרצוף |  |
| 5:15.247 | אחרון |  |
| 5:15.248 | דאצילות |  |
| 5:15.249 | שהוא |  |
| 5:15.250 | זעיר |  |
| 5:15.251 | אנפין |  |
| 5:15.252 | ונוקביה |  |
| 5:15.253 | שכבר |  |
| 5:15.254 | נתעבו |  |
| 5:15.255 | קצת |  |
| 5:15.256 | שייך |  |
| 5:15.257 | בהו |  |
| 5:15.258 | עבודה |  |
| 5:15.259 | ותפלה | . |
| 5:15.260 | ודברים |  |
| 5:15.261 | אלו |  |
| 5:15.262 | מתבארים |  |
| 5:15.263 | מדברי |  |
| 5:15.264 | הרב |  |
| 5:15.265 | בעל |  |
| 5:15.266 | ספר |  |
| 5:15.267 | הברית |  |
| 5:15.268 | מאמר |  |
| 5:15.269 | הנזכר |  |
| 5:15.270 | לעיל |  |
| 5:15.271 | והרב |  |
| 5:15.272 | כסא |  |
| 5:15.273 | אליהו |  |
| 5:15.274 | והרח"ו |  |
| 5:15.275 | בספר |  |
| 5:15.276 | עץ |  |
| 5:15.277 | חיים |  |
| 5:15.278 | והרב |  |
| 5:15.279 | בעל |  |
| 5:15.280 | המצרף |  |
| 5:15.281 | ונחלת |  |
| 5:15.282 | יוסף |  |
| 5:15.283 | וזולתם | . |
| 5:15.284 | וזה |  |
| 5:15.285 | לשון |  |
| 5:15.286 | ספר |  |
| 5:15.287 | הברית | : |
| 5:15.288 | ובזאת |  |
| 5:15.289 | תדע |  |
| 5:15.290 | אחי |  |
| 5:15.291 | שאמונתינו |  |
| 5:15.292 | אינה |  |
| 5:15.293 | כאמונת |  |
| 5:15.294 | הפילוסופים |  |
| 5:15.295 | האלהיים | ) |
| 5:15.296 | והישמעאלים |  |
| 5:15.297 | בענין |  |
| 5:15.298 | היחוד |  |
| 5:15.299 | והבורא | . |
| 5:15.300 | כי |  |
| 5:15.301 | הם |  |
| 5:15.302 | אין |  |
| 5:15.303 | איתם |  |
| 5:15.304 | יודע |  |
| 5:15.305 | את |  |
| 5:15.306 | השם |  |
| 5:15.307 | הנכבד |  |
| 5:15.308 | והנורא |  |
| 5:15.309 | הוי'ה |  |
| 5:15.310 | ב"ה | . |
| 5:15.311 | רק |  |
| 5:15.312 | מאמינים |  |
| 5:15.313 | במחוייב |  |
| 5:15.314 | המציאות |  |
| 5:15.315 | סתם |  |
| 5:15.316 | בבחינתו |  |
| 5:15.317 | הפשוטה |  |
| 5:15.318 | כאשר |  |
| 5:15.319 | בתחלה |  |
| 5:15.320 | קודם |  |
| 5:15.321 | הבריאה |  |
| 5:15.322 | ערום |  |
| 5:15.323 | בלי |  |
| 5:15.324 | לבוש | ) |
| 5:15.325 | וזהו |  |
| 5:15.326 | בחינתו |  |
| 5:15.327 | הבלתי |  |
| 5:15.328 | מושג |  |
| 5:15.329 | אל |  |
| 5:15.330 | זולתו |  |
| 5:15.331 | בשום |  |
| 5:15.332 | רמז |  |
| 5:15.333 | ורמיזה | . |
| 5:15.334 | ובבחינה |  |
| 5:15.335 | הזאת |  |
| 5:15.336 | הטיבו |  |
| 5:15.337 | אשר |  |
| 5:15.338 | דיברו |  |
| 5:15.339 | הפילוסופים |  |
| 5:15.340 | הכופרים | ) |
| 5:15.341 | שלא |  |
| 5:15.342 | שייך |  |
| 5:15.343 | אליו |  |
| 5:15.344 | שום |  |
| 5:15.345 | עבודה |  |
| 5:15.346 | ותפלה |  |
| 5:15.347 | על |  |
| 5:15.348 | כי |  |
| 5:15.349 | הוא |  |
| 5:15.350 | בבחינתו |  |
| 5:15.351 | זאת |  |
| 5:15.352 | נעלה |  |
| 5:15.353 | ומרומם |  |
| 5:15.354 | מכל |  |
| 5:15.355 | ברכה |  |
| 5:15.356 | ותהלה |  |
| 5:15.357 | ובטלה |  |
| 5:15.358 | כל |  |
| 5:15.359 | תורה |  |
| 5:15.360 | ומצוה | . |
| 5:15.361 | כמו |  |
| 5:15.362 | שכתב |  |
| 5:15.363 | ר | ' |
| 5:15.364 | מאיר |  |
| 5:15.365 | גבאי | . |
| 5:15.366 | והמתבונן |  |
| 5:15.367 | בזה |  |
| 5:15.368 | היטב |  |
| 5:15.369 | יתבאר |  |
| 5:15.370 | לו |  |
| 5:15.371 | בלי |  |
| 5:15.372 | שום |  |
| 5:15.373 | ספק |  |
| 5:15.374 | שאין |  |
| 5:15.375 | מקום |  |
| 5:15.376 | לכל |  |
| 5:15.377 | זה |  |
| 5:15.378 | בלתי |  |
| 5:15.379 | אצילות |  |
| 5:15.380 | הספירות | . |
| 5:15.381 | לא |  |
| 5:15.382 | כן |  |
| 5:15.383 | עם |  |
| 5:15.384 | אלהי |  |
| 5:15.385 | אברהם |  |
| 5:15.386 | שמאמינים |  |
| 5:15.387 | במחוייב |  |
| 5:15.388 | המציאות |  |
| 5:15.389 | מלובש |  |
| 5:15.390 | במידותיו |  |
| 5:15.391 | הוא |  |
| 5:15.392 | סוד |  |
| 5:15.393 | הוי"ה |  |
| 5:15.394 | ב"ה |  |
| 5:15.395 | שהוא |  |
| 5:15.396 | הבחינה |  |
| 5:15.397 | שאחר |  |
| 5:15.398 | הבריאה | , |
| 5:15.399 | בחינה |  |
| 5:15.400 | המושגת |  |
| 5:15.401 | לבית |  |
| 5:15.402 | יעקב |  |
| 5:15.403 | שנגלה |  |
| 5:15.404 | למשה |  |
| 5:15.405 | בסנה |  |
| 5:15.406 | והוציאנו |  |
| 5:15.407 | מארץ |  |
| 5:15.408 | מצרים |  |
| 5:15.409 | ונתן |  |
| 5:15.410 | תורת |  |
| 5:15.411 | אמת |  |
| 5:15.412 | לאבותינו |  |
| 5:15.413 | דור |  |
| 5:15.414 | המדבר |  |
| 5:15.415 | פנים |  |
| 5:15.416 | בפנים | , |
| 5:15.417 | שייך |  |
| 5:15.418 | אליו |  |
| 5:15.419 | העבודה |  |
| 5:15.420 | והקרבנות |  |
| 5:15.421 | והתפלה |  |
| 5:15.422 | וכל |  |
| 5:15.423 | המצוות |  |
| 5:15.424 | המדובר |  |
| 5:15.425 | בתוה"ק | . |
| 5:15.426 | ועליו |  |
| 5:15.427 | הזהיר |  |
| 5:15.428 | על |  |
| 5:15.429 | זעיר |  |
| 5:15.430 | אנפין | ) |
| 5:15.431 | משה |  |
| 5:15.432 | ואמר |  |
| 5:15.433 | ליראה |  |
| 5:15.434 | את |  |
| 5:15.435 | השם |  |
| 5:15.436 | הנכבד |  |
| 5:15.437 | והנורא |  |
| 5:15.438 | הזה |  |
| 5:15.439 | ה | ' |
| 5:15.440 | את |  |
| 5:15.441 | אלהיך | " |
| 5:15.442 | עכ"ל | , |
| 5:15.443 | וכן |  |
| 5:15.444 | כתב |  |
| 5:15.445 | מהרח"ו |  |
| 5:15.446 | בעץ |  |
| 5:15.447 | חיים |  |
| 5:15.448 | שער |  |
| 5:15.449 | טנת"א |  |
| 5:15.450 | סוף |  |
| 5:15.451 | פרק |  |
| 5:15.452 | ב | ' |
| 5:15.453 | שהאור |  |
| 5:15.454 | היוצא |  |
| 5:15.455 | מתוך |  |
| 5:15.456 | אדם |  |
| 5:15.457 | קדמון |  |
| 5:15.458 | הוא |  |
| 5:15.459 | זך |  |
| 5:15.460 | מאוד | , |
| 5:15.461 | ועל |  |
| 5:15.462 | ידי |  |
| 5:15.463 | ירידתו |  |
| 5:15.464 | והתרחקותו |  |
| 5:15.465 | הוא |  |
| 5:15.466 | מתעבה |  |
| 5:15.467 | עיבוי |  |
| 5:15.468 | אחר |  |
| 5:15.469 | עיבוי | , |
| 5:15.470 | כיצד |  |
| 5:15.471 | האור |  |
| 5:15.472 | הנמשך |  |
| 5:15.473 | ויוצא |  |
| 5:15.474 | דרך |  |
| 5:15.475 | האוזן |  |
| 5:15.476 | הוא |  |
| 5:15.477 | זך |  |
| 5:15.478 | מאוד |  |
| 5:15.479 | וכאשר |  |
| 5:15.480 | נמשך |  |
| 5:15.481 | האור |  |
| 5:15.482 | הזה |  |
| 5:15.483 | עד |  |
| 5:15.484 | החוטם |  |
| 5:15.485 | ויוצא |  |
| 5:15.486 | קצת |  |
| 5:15.487 | דרך |  |
| 5:15.488 | שם |  |
| 5:15.489 | הוא |  |
| 5:15.490 | מתעבה |  |
| 5:15.491 | וקונה |  |
| 5:15.492 | איזה |  |
| 5:15.493 | עביות |  |
| 5:15.494 | וגסות |  |
| 5:15.495 | מחמת |  |
| 5:15.496 | הריחוק |  |
| 5:15.497 | שנתרחק |  |
| 5:15.498 | ונמשך |  |
| 5:15.499 | יותר |  |
| 5:15.500 | למטה |  |
| 5:15.501 | בצאתו |  |
| 5:15.502 | עד |  |
| 5:15.503 | הפה |  |
| 5:15.504 | ויוצא |  |
| 5:15.505 | קצתו |  |
| 5:15.506 | דרך |  |
| 5:15.507 | שם |  |
| 5:15.508 | מתעבה |  |
| 5:15.509 | יותר |  |
| 5:15.510 | בצאתו |  |
| 5:15.511 | משם |  |
| 5:15.512 | וכו | ' |
| 5:15.513 | וכן |  |
| 5:15.514 | כתב |  |
| 5:15.515 | בשער |  |
| 5:15.516 | הנקודים |  |
| 5:15.517 | פרק |  |
| 5:15.518 | א | ' |
| 5:15.519 | שכל |  |
| 5:15.520 | מה |  |
| 5:15.521 | שהאור |  |
| 5:15.522 | יורד |  |
| 5:15.523 | למטה |  |
| 5:15.524 | הוא |  |
| 5:15.525 | ניכר |  |
| 5:15.526 | ונרגש |  |
| 5:15.527 | ומתגלה |  |
| 5:15.528 | שם |  |
| 5:15.529 | עכ"ל | . |
| 5:15.530 | וזהו |  |
| 5:15.531 | טעמם |  |
| 5:15.532 | במה |  |
| 5:15.533 | שאמרו |  |
| 5:15.534 | דלא |  |
| 5:15.535 | שייך |  |
| 5:15.536 | עבודה |  |
| 5:15.537 | ותפלה |  |
| 5:15.538 | רק |  |
| 5:15.539 | לאלוה |  |
| 5:15.540 | הנרגש |  |
| 5:15.541 | ויש |  |
| 5:15.542 | בו |  |
| 5:15.543 | עובי |  |
| 5:15.544 | וגסות |  |
| 5:15.545 | קצת |  |
| 5:15.546 | שתוכל |  |
| 5:15.547 | לציירו |  |
| 5:15.548 | בדמות |  |
| 5:15.549 | תבניתו |  |
| 5:15.550 | על |  |
| 5:15.551 | ידי |  |
| 5:15.552 | שנתעבה |  |
| 5:15.553 | ונתגשם | . |
| 5:15.554 | אבל |  |
| 5:15.555 | האין |  |
| 5:15.556 | סוף |  |
| 5:15.557 | או |  |
| 5:15.558 | אדם |  |
| 5:15.559 | קדמאה |  |
| 5:15.560 | ואדם |  |
| 5:15.561 | קדמון |  |
| 5:15.562 | ואפילו |  |
| 5:15.563 | עתיק |  |
| 5:15.564 | ואריך |  |
| 5:15.565 | שעדיין |  |
| 5:15.566 | לא |  |
| 5:15.567 | נתעבו |  |
| 5:15.568 | ונתגשמו |  |
| 5:15.569 | כל |  |
| 5:15.570 | כך | , |
| 5:15.571 | לא |  |
| 5:15.572 | שייך |  |
| 5:15.573 | בהן |  |
| 5:15.574 | עבודה | , |
| 5:15.575 | כי |  |
| 5:15.576 | עדיין |  |
| 5:15.577 | לא |  |
| 5:15.578 | נתעבו |  |
| 5:15.579 | ולא |  |
| 5:15.580 | נתגשמו |  |
| 5:15.581 | שתוכל |  |
| 5:15.582 | לציירם | , |
| 5:15.583 | ורק |  |
| 5:15.584 | העבודה |  |
| 5:15.585 | והיראה |  |
| 5:15.586 | לאלוה |  |
| 5:15.587 | קצר |  |
| 5:15.588 | האפים |  |
| 5:15.589 | זעיר |  |
| 5:15.590 | אנפין | ) |
| 5:15.591 | והוא |  |
| 5:15.592 | לבדו |  |
| 5:15.593 | לפי |  |
| 5:15.594 | דעתם |  |
| 5:15.595 | אלהינו |  |
| 5:15.596 | על |  |
| 5:15.597 | ידי |  |
| 5:15.598 | שנתגשם |  |
| 5:15.599 | ונגלה |  |
| 5:15.600 | למשה |  |
| 5:15.601 | בסנה |  |
| 5:15.602 | ובסיני |  |
| 5:15.603 | לאבותינו | . |
| 5:15.604 | והנה |  |
| 5:15.605 | שכחו |  |
| 5:15.606 | מקרא |  |
| 5:15.607 | מלא |  |
| 5:15.608 | כי |  |
| 5:15.609 | מלאך |  |
| 5:15.610 | הוא |  |
| 5:15.611 | שנגלה |  |
| 5:15.612 | למשה |  |
| 5:15.613 | בסנה |  |
| 5:15.614 | שנאמר |  |
| 5:15.615 | שמות |  |
| 5:15.616 | ג | ' |
| 5:15.617 | ב | '): |
| 5:15.618 | וירא |  |
| 5:15.619 | מלאך |  |
| 5:15.620 | ה | ' |
| 5:15.621 | אליו |  |
| 5:15.622 | בלבת |  |
| 5:15.623 | אש |  |
| 5:15.624 | מתוך |  |
| 5:15.625 | הסנה | ". |
| 5:15.626 | ובמעמד |  |
| 5:15.627 | הר |  |
| 5:15.628 | סיני |  |
| 5:15.629 | הלא |  |
| 5:15.630 | נאמר |  |
| 5:15.631 | דברים |  |
| 5:15.632 | ד | ' |
| 5:15.633 | ט"ו | ): |
| 5:15.634 | כי |  |
| 5:15.635 | לא |  |
| 5:15.636 | ראיתם |  |
| 5:15.637 | כל |  |
| 5:15.638 | תמונה |  |
| 5:15.639 | ביום |  |
| 5:15.640 | דבר |  |
| 5:15.641 | ה | ' |
| 5:15.642 | עמכם | ", |
| 5:15.643 | וכו | ' |
| 5:15.644 | לא |  |
| 5:15.645 | ראו |  |
| 5:15.646 | דמות |  |
| 5:15.647 | המושגת |  |
| 5:15.648 | ברמ"ח |  |
| 5:15.649 | איברים |  |
| 5:15.650 | ושס"ה |  |
| 5:15.651 | גידים |  |
| 5:15.652 | כדבריהם | . |
| 5:15.653 | וקרוב |  |
| 5:15.654 | לדבריהם |  |
| 5:15.655 | כתב |  |
| 5:15.656 | הרב |  |
| 5:15.657 | עז |  |
| 5:15.658 | לאלהים |  |
| 5:15.659 | בבית |  |
| 5:15.660 | קדש |  |
| 5:15.661 | הקדשים |  |
| 5:15.662 | דף |  |
| 5:15.663 | כ"ו |  |
| 5:15.664 | ע"ד |  |
| 5:15.665 | וזה |  |
| 5:15.666 | לשונו | : |
| 5:15.667 | מי |  |
| 5:15.668 | שמאמין |  |
| 5:15.669 | שהאין |  |
| 5:15.670 | סוף |  |
| 5:15.671 | הוא |  |
| 5:15.672 | עצם |  |
| 5:15.673 | האלוהות |  |
| 5:15.674 | כאילו |  |
| 5:15.675 | כופר |  |
| 5:15.676 | בעיקר | . |
| 5:15.677 | ועליהם |  |
| 5:15.678 | אמר |  |
| 5:15.679 | הנביא |  |
| 5:15.680 | ישעיהו |  |
| 5:15.681 | כ"ט |  |
| 5:15.682 | ט"ו | ): |
| 5:15.683 | הוי |  |
| 5:15.684 | המעמיקים |  |
| 5:15.685 | מה | ' |
| 5:15.686 | לסתיר |  |
| 5:15.687 | עצה | '. |
| 5:15.688 | ועוד |  |
| 5:15.689 | שעצם |  |
| 5:15.690 | האלוהות |  |
| 5:15.691 | הוא |  |
| 5:15.692 | שם |  |
| 5:15.693 | הוי"ה | , |
| 5:15.694 | וכל |  |
| 5:15.695 | דבר |  |
| 5:15.696 | שיש |  |
| 5:15.697 | לו |  |
| 5:15.698 | שם |  |
| 5:15.699 | ואות |  |
| 5:15.700 | יש |  |
| 5:15.701 | בו |  |
| 5:15.702 | שיעור |  |
| 5:15.703 | קומה |  |
| 5:15.704 | וכו | '. |
| 5:15.705 | ושם |  |
| 5:15.706 | הדבר |  |
| 5:15.707 | הוא |  |
| 5:15.708 | גבולו |  |
| 5:15.709 | שכן |  |
| 5:15.710 | שם |  |
| 5:15.711 | האדם |  |
| 5:15.712 | לא |  |
| 5:15.713 | שייך |  |
| 5:15.714 | שיקרא |  |
| 5:15.715 | בשם |  |
| 5:15.716 | אדם |  |
| 5:15.717 | אלא |  |
| 5:15.718 | מראשו |  |
| 5:15.719 | ועד |  |
| 5:15.720 | רגליו | , |
| 5:15.721 | הרי |  |
| 5:15.722 | שהשם |  |
| 5:15.723 | הוא |  |
| 5:15.724 | גבולו | . |
| 5:15.725 | לכן |  |
| 5:15.726 | כל |  |
| 5:15.727 | מה |  |
| 5:15.728 | שתמצא |  |
| 5:15.729 | משמות |  |
| 5:15.730 | דשייכי |  |
| 5:15.731 | באדם |  |
| 5:15.732 | שייכים |  |
| 5:15.733 | אליו |  |
| 5:15.734 | יתברך | , |
| 5:15.735 | ואינו |  |
| 5:15.736 | כסברת |  |
| 5:15.737 | הטועים |  |
| 5:15.738 | שאומרים |  |
| 5:15.739 | שעצם |  |
| 5:15.740 | האלוהות |  |
| 5:15.741 | הוא |  |
| 5:15.742 | אין |  |
| 5:15.743 | סוף | "! |
| 5:15.744 | וחולק |  |
| 5:15.745 | הוא |  |
| 5:15.746 | על |  |
| 5:15.747 | מהרח"ו |  |
| 5:15.748 | ושאר |  |
| 5:15.749 | מקובלים |  |
| 5:15.750 | שסוברים |  |
| 5:15.751 | שהתפארת |  |
| 5:15.752 | שהוא |  |
| 5:15.753 | זעיר |  |
| 5:15.754 | אנפין |  |
| 5:15.755 | הוא |  |
| 5:15.756 | אלהינו | , |
| 5:15.757 | וסבירא |  |
| 5:15.758 | ליה |  |
| 5:15.759 | שהעבודה |  |
| 5:15.760 | והתפלה |  |
| 5:15.761 | למלכא |  |
| 5:15.762 | קדישא |  |
| 5:15.763 | דכל |  |
| 5:15.764 | קדישין |  |
| 5:15.765 | שהוא |  |
| 5:15.766 | מאיר |  |
| 5:15.767 | בלב |  |
| 5:15.768 | זעיר |  |
| 5:15.769 | אנפין |  |
| 5:15.770 | כמו |  |
| 5:15.771 | שכתב |  |
| 5:15.772 | בפרק |  |
| 5:15.773 | י"ט |  |
| 5:15.774 | דף |  |
| 5:15.775 | נ"ט |  |
| 5:15.776 | ע"ד | , |
| 5:15.777 | וזה |  |
| 5:15.778 | לשונו |  |
| 5:15.779 | בדף |  |
| 5:15.780 | כ"ו |  |
| 5:15.781 | סוף |  |
| 5:15.782 | ע"ד | : |
| 5:15.783 | בר |  |
| 5:15.784 | מן |  |
| 5:15.785 | דין |  |
| 5:15.786 | ובר |  |
| 5:15.787 | מן |  |
| 5:15.788 | דין |  |
| 5:15.789 | אם |  |
| 5:15.790 | איתא |  |
| 5:15.791 | שהתפארת |  |
| 5:15.792 | הוא |  |
| 5:15.793 | עצם |  |
| 5:15.794 | האלהות |  |
| 5:15.795 | למה |  |
| 5:15.796 | נעבוד |  |
| 5:15.797 | את |  |
| 5:15.798 | התפארת |  |
| 5:15.799 | שהוא |  |
| 5:15.800 | מקבל |  |
| 5:15.801 | מאבא |  |
| 5:15.802 | ואמא | , |
| 5:15.803 | נעבוד |  |
| 5:15.804 | לאריך |  |
| 5:15.805 | אנפין |  |
| 5:15.806 | או |  |
| 5:15.807 | לכתר |  |
| 5:15.808 | שהוא |  |
| 5:15.809 | יותר |  |
| 5:15.810 | למעלה | ? |
| 5:15.811 | ואם |  |
| 5:15.812 | האין |  |
| 5:15.813 | סוף |  |
| 5:15.814 | הוא |  |
| 5:15.815 | עצם |  |
| 5:15.816 | האלוהות | , |
| 5:15.817 | איך |  |
| 5:15.818 | אנו |  |
| 5:15.819 | אומרים |  |
| 5:15.820 | ברוך |  |
| 5:15.821 | אתה |  |
| 5:15.822 | ה | ', |
| 5:15.823 | ממי |  |
| 5:15.824 | יתברך |  |
| 5:15.825 | ויביא |  |
| 5:15.826 | לו |  |
| 5:15.827 | שפע | "? |
| 5:15.828 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

