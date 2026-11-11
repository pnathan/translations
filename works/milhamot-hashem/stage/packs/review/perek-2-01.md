# Review pack: milhamot-hashem perek-2-01 (Milchamot ha-Shem 2:1–2:4)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-01`.

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
- ר"א: "Rabbi Eleazar" only where the paragraph names him, else "R. A." with
  a note.
- ולא אמרינן כלא חד (1:7) and באגרונו (1:12) are for the owner: the main
  English stands as translated; the reviewer's alternative is in `readings`
  / `apparatus`.

## Rulings from wave 1 (chapter 1), by the editor

- סתו"מ: always "Torah scrolls, tefillin and mezuzot", with a note.
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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
- עה"ת: "on the Torah"; "a book on the Torah" where it follows a book's name
  and the work is unidentified. Do not guess Etz Chayyim.
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
- והוא שלא: "provided that" (fixed).
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


## Milchamot ha-Shem 2:1

Hebrew (Wikisource): מלחמות ה'

| id | word | punctuation after |
|---|---|---|
| 2:1.1 | מלחמות |  |
| 2:1.2 | ה | ' |


### Translation under review (2:1)

English: The Wars of the LORD

Literal: The-Wars-of the-LORD

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:1.1 | The Wars of |  |  |  |
| 2:1.2 | the LORD |  |  | Tetragrammaton written ה׳ (the pack splits it into ה and the punctuation '); read as 'the LORD' (decisions) |

Phrases:
- 2:1.1 2:1.2 | The Wars of the LORD | The-Wars-of the-LORD

Grammar notes:
- מלחמות is a construct plural, 'wars of'; the Name that follows is the Tetragrammaton, read 'the LORD'.
- The title is taken from Num 21:14, 'the book of the Wars of the LORD', and is given in English as the decisions give it.

Translator's note: Title of the book (decisions: Names and titles).

Commentary rulings:
- apparatus: Num 21:14: 'the book of the Wars of the LORD', the source of the title.

## Milchamot ha-Shem 2:2

Hebrew (Wikisource): שאול שאל אחד מן התלמידים וז"ל יודיענו מ"ו העבודה והתפלה שאנו מתפללים למי שייכא. לעילת העילות ב"ו או לנאצל אשר נאצל ממנו (ז"א) כאשר פירש לנו באוזניכם מ"ו סאלם ן' דוד שנסע מכאן לארץ הקדש ע"פ האמונה האמתית כמפורש בס' מצרף האמונה ובס' עה"ת שהעובד לאין סוף עובד ללא אלהי אמת וללא אלהי משפט כדי שנדע את מי נעבוד ותשובתך מהרה אדוני ע"כ.

| id | word | punctuation after |
|---|---|---|
| 2:2.1 | שאול |  |
| 2:2.2 | שאל |  |
| 2:2.3 | אחד |  |
| 2:2.4 | מן |  |
| 2:2.5 | התלמידים |  |
| 2:2.6 | וז"ל |  |
| 2:2.7 | יודיענו |  |
| 2:2.8 | מ"ו |  |
| 2:2.9 | העבודה |  |
| 2:2.10 | והתפלה |  |
| 2:2.11 | שאנו |  |
| 2:2.12 | מתפללים |  |
| 2:2.13 | למי |  |
| 2:2.14 | שייכא | . |
| 2:2.15 | לעילת |  |
| 2:2.16 | העילות |  |
| 2:2.17 | ב"ו |  |
| 2:2.18 | או |  |
| 2:2.19 | לנאצל |  |
| 2:2.20 | אשר |  |
| 2:2.21 | נאצל |  |
| 2:2.22 | ממנו |  |
| 2:2.23 | ז"א | ) |
| 2:2.24 | כאשר |  |
| 2:2.25 | פירש |  |
| 2:2.26 | לנו |  |
| 2:2.27 | באוזניכם |  |
| 2:2.28 | מ"ו |  |
| 2:2.29 | סאלם |  |
| 2:2.30 | ן | ' |
| 2:2.31 | דוד |  |
| 2:2.32 | שנסע |  |
| 2:2.33 | מכאן |  |
| 2:2.34 | לארץ |  |
| 2:2.35 | הקדש |  |
| 2:2.36 | ע"פ |  |
| 2:2.37 | האמונה |  |
| 2:2.38 | האמתית |  |
| 2:2.39 | כמפורש |  |
| 2:2.40 | בס | ' |
| 2:2.41 | מצרף |  |
| 2:2.42 | האמונה |  |
| 2:2.43 | ובס | ' |
| 2:2.44 | עה"ת |  |
| 2:2.45 | שהעובד |  |
| 2:2.46 | לאין |  |
| 2:2.47 | סוף |  |
| 2:2.48 | עובד |  |
| 2:2.49 | ללא |  |
| 2:2.50 | אלהי |  |
| 2:2.51 | אמת |  |
| 2:2.52 | וללא |  |
| 2:2.53 | אלהי |  |
| 2:2.54 | משפט |  |
| 2:2.55 | כדי |  |
| 2:2.56 | שנדע |  |
| 2:2.57 | את |  |
| 2:2.58 | מי |  |
| 2:2.59 | נעבוד |  |
| 2:2.60 | ותשובתך |  |
| 2:2.61 | מהרה |  |
| 2:2.62 | אדוני |  |
| 2:2.63 | ע"כ | . |


### Translation under review (2:2)

English: Question: one of the students asked, and these are his words: "Let our teacher inform us: the worship and the prayer that we pray, to whom it pertains. To the Cause of causes, blessed be He, or to the emanated being that was emanated from Him (Ze'ir Anpin), as explained to us in your hearing by our teacher Salem son of David, who travelled from here to the Holy Land, according to the true faith, as explained in the book Matzref ha-Emunah and in a book on the Torah, that whoever worships Ein Sof worships without a God of truth and without a God of judgment, so that we may know whom we shall worship. And your reply, speedily, my lord. Thus far."

Literal: asking asked one of the-students and-these-are-his-words let-inform-us our-teacher the-worship and-the-prayer that-we pray to-whom it-pertains to-the-Cause-of causes blessed-be-He or to-the-emanated-being that was-emanated from-Him Ze'ir-Anpin as explained to-us in-your-hearing our-teacher Salem son of-David who-travelled from-here to-the-land of-the-Holy according-to the-faith the-true as-explained in-the-book Matzref ha-Emunah and-in-a-book on-the-Torah that-whoever-worships to-Ein Sof worships without a-God-of truth and-without a-God-of judgment so-that we-may-know - whom we-shall-worship and-your-reply speedily my-lord thus-far

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:2.1 | asking |  |  | infinitive absolute with שאל; the heading 'Question' (decisions) |
| 2:2.2 | asked |  |  |  |
| 2:2.3 | one |  |  |  |
| 2:2.4 | of |  |  |  |
| 2:2.5 | the students |  |  |  |
| 2:2.6 | and these are his words |  |  | the quotation formula וז״ל, 'and these are his words'; the quotation closes at ע״כ (63) |
| 2:2.7 | let inform us |  |  |  |
| 2:2.8 | our teacher |  |  | מ״ו read as מורנו, 'our teacher' (the same at 28); the decisions give no ruling on this abbreviation |
| 2:2.9 | the worship |  |  |  |
| 2:2.10 | and the prayer |  |  |  |
| 2:2.11 | that we |  |  |  |
| 2:2.12 | pray |  |  |  |
| 2:2.13 | to whom |  |  |  |
| 2:2.14 | it pertains |  |  |  |
| 2:2.15 | to the Cause of |  |  |  |
| 2:2.16 | causes |  |  |  |
| 2:2.17 | blessed be He |  |  | ב״ו read as ברוך הוא, 'blessed be He'; the decisions give no ruling on this abbreviation |
| 2:2.18 | or |  |  |  |
| 2:2.19 | to the emanated being |  |  |  |
| 2:2.20 | that |  |  |  |
| 2:2.21 | was emanated |  |  |  |
| 2:2.22 | from Him |  |  |  |
| 2:2.23 | Ze'ir Anpin |  |  | ז״א: Ze'ir Anpin, literally 'small face'; a partzuf of the Kabbalists. The parentheses are the writer's and stay |
| 2:2.24 | as |  |  |  |
| 2:2.25 | explained |  |  |  |
| 2:2.26 | to us |  |  |  |
| 2:2.27 | in your hearing |  |  |  |
| 2:2.28 | our teacher |  |  | מ״ו: 'our teacher', as at 2:2.8 |
| 2:2.29 | Salem |  |  |  |
| 2:2.30 | son |  |  |  |
| 2:2.31 | of David |  |  |  |
| 2:2.32 | who travelled |  |  |  |
| 2:2.33 | from here |  |  |  |
| 2:2.34 | to the land |  |  |  |
| 2:2.35 | of the Holy |  |  |  |
| 2:2.36 | according to |  |  | ע״פ: על פי, 'according to' |
| 2:2.37 | the faith |  |  |  |
| 2:2.38 | the true |  |  |  |
| 2:2.39 | as explained |  |  |  |
| 2:2.40 | in the book |  |  |  |
| 2:2.41 | Matzref |  |  | Matzref ha-Emunah, literally 'the refining of faith' |
| 2:2.42 | ha-Emunah |  |  |  |
| 2:2.43 | and in a book |  |  |  |
| 2:2.44 | on the Torah |  |  | עה״ת read as על התורה, 'on the Torah'; see apparatus |
| 2:2.45 | that whoever worships |  |  |  |
| 2:2.46 | to Ein |  |  | Ein Sof, literally 'without end' |
| 2:2.47 | Sof |  |  |  |
| 2:2.48 | worships |  |  |  |
| 2:2.49 | without |  |  |  |
| 2:2.50 | a God of |  |  |  |
| 2:2.51 | truth |  |  |  |
| 2:2.52 | and without |  |  |  |
| 2:2.53 | a God of |  |  |  |
| 2:2.54 | judgment |  |  |  |
| 2:2.55 | so that |  |  |  |
| 2:2.56 | we may know |  |  |  |
| 2:2.57 | - |  |  | את: the object marker, with no English counterpart |
| 2:2.58 | whom |  |  |  |
| 2:2.59 | we shall worship |  |  |  |
| 2:2.60 | and your reply |  |  |  |
| 2:2.61 | speedily |  |  |  |
| 2:2.62 | my lord |  |  |  |
| 2:2.63 | thus far |  |  | ע״כ: 'thus far'; closes the quotation |

Phrases:
- 2:2.1 2:2.2 2:2.3 2:2.4 2:2.5 | Question: one of the students asked, | asking asked one of the-students
- 2:2.6 | and these are his words: | and-these-are-his-words
- 2:2.7 2:2.8 | "Let our teacher inform us: | let-inform-us our-teacher
- 2:2.9 2:2.10 2:2.11 2:2.12 2:2.13 2:2.14 | the worship and the prayer that we pray, to whom it pertains. | the-worship and-the-prayer that-we pray to-whom it-pertains
- 2:2.15 2:2.16 2:2.17 | To the Cause of causes, blessed be He, | to-the-Cause-of causes blessed-be-He
- 2:2.18 2:2.19 2:2.20 2:2.21 2:2.22 2:2.23 | or to the emanated being that was emanated from Him (Ze'ir Anpin), | or to-the-emanated-being that was-emanated from-Him Ze'ir-Anpin
- 2:2.24 2:2.25 2:2.26 2:2.27 | as explained to us in your hearing | as explained to-us in-your-hearing
- 2:2.28 2:2.29 2:2.30 2:2.31 2:2.32 2:2.33 | by our teacher Salem son of David, who travelled from here | our-teacher Salem son of-David who-travelled from-here
- 2:2.34 2:2.35 | to the Holy Land, | to-the-land of-the-Holy
- 2:2.36 2:2.37 2:2.38 2:2.39 2:2.40 2:2.41 2:2.42 2:2.43 2:2.44 | according to the true faith, as explained in the book Matzref ha-Emunah and in a book on the Torah, | according-to the-faith the-true as-explained in-the-book Matzref ha-Emunah and-in-a-book on-the-Torah
- 2:2.45 2:2.46 2:2.47 2:2.48 | that whoever worships Ein Sof worships | that-whoever-worships to-Ein Sof worships
- 2:2.49 2:2.50 2:2.51 | without a God of truth | without a-God-of truth
- 2:2.52 2:2.53 2:2.54 | and without a God of judgment, | and-without a-God-of judgment
- 2:2.55 2:2.56 2:2.57 2:2.58 2:2.59 | so that we may know whom we shall worship. | so-that we-may-know - whom we-shall-worship
- 2:2.60 2:2.61 2:2.62 2:2.63 | And your reply, speedily, my lord. Thus far." | and-your-reply speedily my-lord thus-far

Grammar notes:
- שאול שאל is an infinitive absolute before a finite verb: 'asking, he asked'. The heading of the correspondence is 'Question' (decisions).
- וז״ל introduces the quotation, which runs to ע״כ at the end; the question is one long quoted sentence-chain.
- The question offers two objects of worship joined by או, 'or': the Cause of causes, or the emanated being. The או sets them as alternatives.
- שהעובד לאין סוף עובד ללא אלהי אמת: the participle שהעובד is the subject, 'whoever worships Ein Sof', and the verb repeats as 'worships'.
- כדי שנדע את מי נעבוד: את is the object marker before מי, 'whom'; the purpose clause closes the question.

Translator's note: The quotation runs from וז״ל to ע״כ. The parentheses around ז״א are the writer's and stay in parentheses inside the quotation.

Commentary rulings:
- apparatus: עה״ת (2:2.44) is read as על התורה, 'on the Torah'. The abbreviation may also stand for the title Etz ha-Chayyim; the decisions give no ruling.

## Milchamot ha-Shem 2:3

Hebrew (Wikisource): תשובה. תשובתך מבוארת בספר כסא אליהו מדף י"ג עד דף י"ח ישמע חכם ויוסף לקח. והמקובל בידינו שהעבודה לאל אדון אחד להמשיך השפע ממנו אל עשר ספירות כמו שכתב לחם שלמה. עכ"ל המשיב.

| id | word | punctuation after |
|---|---|---|
| 2:3.1 | תשובה | . |
| 2:3.2 | תשובתך |  |
| 2:3.3 | מבוארת |  |
| 2:3.4 | בספר |  |
| 2:3.5 | כסא |  |
| 2:3.6 | אליהו |  |
| 2:3.7 | מדף |  |
| 2:3.8 | י"ג |  |
| 2:3.9 | עד |  |
| 2:3.10 | דף |  |
| 2:3.11 | י"ח |  |
| 2:3.12 | ישמע |  |
| 2:3.13 | חכם |  |
| 2:3.14 | ויוסף |  |
| 2:3.15 | לקח | . |
| 2:3.16 | והמקובל |  |
| 2:3.17 | בידינו |  |
| 2:3.18 | שהעבודה |  |
| 2:3.19 | לאל |  |
| 2:3.20 | אדון |  |
| 2:3.21 | אחד |  |
| 2:3.22 | להמשיך |  |
| 2:3.23 | השפע |  |
| 2:3.24 | ממנו |  |
| 2:3.25 | אל |  |
| 2:3.26 | עשר |  |
| 2:3.27 | ספירות |  |
| 2:3.28 | כמו |  |
| 2:3.29 | שכתב |  |
| 2:3.30 | לחם |  |
| 2:3.31 | שלמה | . |
| 2:3.32 | עכ"ל |  |
| 2:3.33 | המשיב | . |


### Translation under review (2:3)

English: Reply. "Your reply is explained in the book Kisse Eliyahu, from folio 13 to folio 18: 'A wise man will hear and add to learning.' And what we have received is that worship is to God, one Lord, in order to draw the flow from Him to the ten sefirot, as Lechem Shlomo wrote. Thus far the words of the respondent."

Literal: Reply your-reply is-explained in-the-book Kisse Eliyahu from-folio 13 to folio 18 will-hear a-wise-man and-will-add to-learning and-what-is-received by-us that-worship to-God lord one in-order-to-draw the-flow from-Him to ten sefirot as wrote Lechem Shlomo thus-far the-respondent

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:3.1 | Reply |  |  | the heading of the reply, 'Reply' (decisions) |
| 2:3.2 | your reply |  |  |  |
| 2:3.3 | is explained |  |  |  |
| 2:3.4 | in the book |  |  |  |
| 2:3.5 | Kisse |  |  | Kisse Eliyahu, literally 'the throne of Elijah'; first occurrence in the unit, so the meaning is given here |
| 2:3.6 | Eliyahu |  |  |  |
| 2:3.7 | from folio |  |  |  |
| 2:3.8 | 13 |  |  |  |
| 2:3.9 | to |  |  |  |
| 2:3.10 | folio |  |  |  |
| 2:3.11 | 18 |  |  |  |
| 2:3.12 | will hear |  |  | Prov 1:5, quoted; the yiqtol is read as the verse's future (see readings) |
| 2:3.13 | a wise man |  |  |  |
| 2:3.14 | and will add |  |  |  |
| 2:3.15 | to learning |  |  |  |
| 2:3.16 | and what is received |  |  | המקובל: the participle of קבל, 'to receive'; the same root as Kabbalah, which the author claims for the rabbinic chain |
| 2:3.17 | by us |  |  |  |
| 2:3.18 | that worship |  |  |  |
| 2:3.19 | to God |  |  |  |
| 2:3.20 | lord |  |  |  |
| 2:3.21 | one |  |  |  |
| 2:3.22 | in order to draw |  |  | להמשיך: infinitive of purpose, 'in order to draw' |
| 2:3.23 | the flow |  |  |  |
| 2:3.24 | from Him |  |  |  |
| 2:3.25 | to |  |  |  |
| 2:3.26 | ten |  |  |  |
| 2:3.27 | sefirot |  |  |  |
| 2:3.28 | as |  |  |  |
| 2:3.29 | wrote |  |  |  |
| 2:3.30 | Lechem |  |  | Lechem Shlomo, literally 'the bread of Solomon'; a book title; the decisions give no ruling on it |
| 2:3.31 | Shlomo |  |  |  |
| 2:3.32 | thus far |  |  | עכ״ל: closes the quotation |
| 2:3.33 | the respondent |  |  | המשיב: 'the respondent' |

Phrases:
- 2:3.1 | Reply. | Reply
- 2:3.2 2:3.3 2:3.4 2:3.5 2:3.6 2:3.7 2:3.8 2:3.9 2:3.10 2:3.11 | "Your reply is explained in the book Kisse Eliyahu, from folio 13 to folio 18: | your-reply is-explained in-the-book Kisse Eliyahu from-folio 13 to folio 18
- 2:3.12 2:3.13 2:3.14 2:3.15 | 'A wise man will hear and add to learning.' | will-hear a-wise-man and-will-add to-learning
- 2:3.16 2:3.17 2:3.18 2:3.19 2:3.20 2:3.21 | And what we have received is that worship is to God, one Lord, | and-what-is-received by-us that-worship to-God lord one
- 2:3.22 2:3.23 2:3.24 2:3.25 2:3.26 2:3.27 | in order to draw the flow from Him to the ten sefirot, | in-order-to-draw the-flow from-Him to ten sefirot
- 2:3.28 2:3.29 2:3.30 2:3.31 | as Lechem Shlomo wrote. | as wrote Lechem Shlomo
- 2:3.32 2:3.33 | Thus far the words of the respondent." | thus-far the-respondent

Grammar notes:
- The reply is the respondent's letter, quoted in double marks from 'Your reply' to the closing 'Thus far'; the verse inside it is in single marks.
- ישמע חכם ויוסף לקח is Prov 1:5, quoted without a formula. The yiqtol ישמע is read as the verse's future, 'will hear'.
- והמקובל בידינו שהעבודה לאל אדון אחד: the participle phrase is the subject of 'is that worship is to one Lord'; the English puts it first.
- להמשיך השפע ממנו אל עשר ספירות: the infinitive gives the purpose of the worship, to draw the flow from Him to the ten sefirot.
- כמו שכתב לחם שלמה: 'as Lechem Shlomo wrote'; a book cited by its title, with no reference given.

Translator's note: The respondent's answer to the question in 2:2; its closing formula עכ״ל closes the quotation.

Commentary rulings:
- apparatus: Prov 1:5 (2:3.12–15): 'A wise man will hear and add to learning'.

## Milchamot ha-Shem 2:4

Hebrew (Wikisource): חזר השואל והשיב. תשובתך ידידי באה סתומה וחתומה ואין בה אלא מורה מקום, "לעיין בכסא אליהו מדף י"ג עד דף י"ח ישמע חכם ויוסף לקח והמקובל בידינו שהעבודה לאל אדון אחד" וכו'. ולא הבנתי דעתך וכוונתך בזה, אם זה המקובל בידך הוא כדברי כסא אליהו או היפך דבריו. כי כאשר עיינתי בכסא אליהו מקום שרשמת לי, לא מצאתי מבוקשי, אדרבא שם בדף ג' מדמה אחדות הבורא לשאר מיני האחדים, ששם כתב: כמו שזה הבית בכללותו נקרא בשם בית אחד. ואם תכנס בו תמצא כמה וכמה חדרים גדולים וקטנים וכמה מקומות זולתם. וכל א' ואחד מתארים לו תואר שם אחד. גם הבנין של עצים שיש בתוכו כמה תארים יש בו. גם תמצא עוד שזה הכותל קדש שנבנה היה חלקים כל אבן ואבן היא חלוקה מצרורות העפר לעצמו והסיד לעצמו כו'. ואחר שבא הבנאי וחבר כל חלקים אלו בחכמתו זה בתוך זה וזה אצל זה ונעשו כותל אחד וכו', נמצא אז הכל חתיכה אחת וחלק אחד. וכן הענין אצלנו וכו'. ושם בדף כ"ה וכ"ט מדמה אותו לגוף האדם המחובר מעצמות וגידים ובשר וראש ועין ואזן וחוטם ופה וידים ורגלים והכל נקרא בשם אדם ראובן או שמעון. ובדף כ"ה וכ"ו וכ"ז כתב: הכלל העולה כי הסבה הראשונה הנקרא בפי כל המקובלים אין סוף, הוא אשר האציל וברא ויצר ועשה והוא מסתתר תוך זעיר אנפין ואזי זעיר אנפין הוא השולט על כל הברואים ומנהיגם וזנם ומפרנסם בכח אין סוף שבו. ולכן הוא אלהינו ואנחנו עמו כי נשמותינו הם חלקו ואותו נעבוד והוא אלהי אבותינו שבידו הנהגת העולמות בשכר ועונש אבל האין סוף עם שאר הפרצופים שלמעלה מזעיר ונוקביה אם יסדרו בני אדם תפלתם להם בפרטות מבלי שיתפללו לזעיר אנפין גם אם יכוונו לנשמה המסתתרת בהם יהיו צועקים ואינם נענים. ואדרבה יהיו המתפללים אליהם נענשים כי זהו רצון הסבה הראשונה שיהיה הוא המוציא והמביא שפעו לתחתונים ואין עוד מלבדו ע"ש היטב. ודברי הר' כסא אליהו עומדים נגד דברי הרמב"ם ז"ל בפי' המשנה ובחבור ובמורה נבוכים ונגד דברי החסיד בעל חובת הלבבות בשער היחוד והרס"ג בס' האמונות והדעות והרוקח שכתבו שאחדותו ית' איננו כאחד הזוג ולא כאחד המין ולא כאיש האחד שנחלק לאחדים רבים. ולא אחד כמו הגוף הפשוט וכו' שמקבל החלוק לאין סוף. אבל השי"ת הוא אחד באחדות שאין כמותה. ג"כ כתב כסא אליהו שהאין סוף הוא נשמת עתיק, וארוך אנפין, ואבא ואימא, וזעיר ונוקביה, ורבותי ז"ל הנז"ל כתבו שאינו גוף ולא כח בגוף! ולפי דבריו הוא כח בגוף. ג"כ כתב שאריך אנפין אבא ואימא קדמו לזעיר אנפין שהוא אלהינו (לפי דבריו) והוא נקרא בן לאבא ואימא וכו'. ורז"ל אמרו שהוא קדמון לכל הנמצאים ושכל נמצא זולתו בלתי קדמון וכו'. ג"כ כתב שאין סוף ב"ה לא שייכא שום עבודה בו ולא שום קריאה כלל. ולא שייך בו שום שם לא הוי"ה ולא אדנות וכו' רק לז"א ונוקביה שבעולם האצילות בלבד ולא לא"א ואבא ואימא וזעיר ונוקביה שיש בעולמות הנאצלים למעלה מעולם דאצילות כמבואר שם ד' מ"ט סוף עמוד א' ובדף נ"ג....

| id | word | punctuation after |
|---|---|---|
| 2:4.1 | חזר |  |
| 2:4.2 | השואל |  |
| 2:4.3 | והשיב | . |
| 2:4.4 | תשובתך |  |
| 2:4.5 | ידידי |  |
| 2:4.6 | באה |  |
| 2:4.7 | סתומה |  |
| 2:4.8 | וחתומה |  |
| 2:4.9 | ואין |  |
| 2:4.10 | בה |  |
| 2:4.11 | אלא |  |
| 2:4.12 | מורה |  |
| 2:4.13 | מקום | , |
| 2:4.14 | לעיין |  |
| 2:4.15 | בכסא |  |
| 2:4.16 | אליהו |  |
| 2:4.17 | מדף |  |
| 2:4.18 | י"ג |  |
| 2:4.19 | עד |  |
| 2:4.20 | דף |  |
| 2:4.21 | י"ח |  |
| 2:4.22 | ישמע |  |
| 2:4.23 | חכם |  |
| 2:4.24 | ויוסף |  |
| 2:4.25 | לקח |  |
| 2:4.26 | והמקובל |  |
| 2:4.27 | בידינו |  |
| 2:4.28 | שהעבודה |  |
| 2:4.29 | לאל |  |
| 2:4.30 | אדון |  |
| 2:4.31 | אחד | " |
| 2:4.32 | וכו | '. |
| 2:4.33 | ולא |  |
| 2:4.34 | הבנתי |  |
| 2:4.35 | דעתך |  |
| 2:4.36 | וכוונתך |  |
| 2:4.37 | בזה | , |
| 2:4.38 | אם |  |
| 2:4.39 | זה |  |
| 2:4.40 | המקובל |  |
| 2:4.41 | בידך |  |
| 2:4.42 | הוא |  |
| 2:4.43 | כדברי |  |
| 2:4.44 | כסא |  |
| 2:4.45 | אליהו |  |
| 2:4.46 | או |  |
| 2:4.47 | היפך |  |
| 2:4.48 | דבריו | . |
| 2:4.49 | כי |  |
| 2:4.50 | כאשר |  |
| 2:4.51 | עיינתי |  |
| 2:4.52 | בכסא |  |
| 2:4.53 | אליהו |  |
| 2:4.54 | מקום |  |
| 2:4.55 | שרשמת |  |
| 2:4.56 | לי | , |
| 2:4.57 | לא |  |
| 2:4.58 | מצאתי |  |
| 2:4.59 | מבוקשי | , |
| 2:4.60 | אדרבא |  |
| 2:4.61 | שם |  |
| 2:4.62 | בדף |  |
| 2:4.63 | ג | ' |
| 2:4.64 | מדמה |  |
| 2:4.65 | אחדות |  |
| 2:4.66 | הבורא |  |
| 2:4.67 | לשאר |  |
| 2:4.68 | מיני |  |
| 2:4.69 | האחדים | , |
| 2:4.70 | ששם |  |
| 2:4.71 | כתב | : |
| 2:4.72 | כמו |  |
| 2:4.73 | שזה |  |
| 2:4.74 | הבית |  |
| 2:4.75 | בכללותו |  |
| 2:4.76 | נקרא |  |
| 2:4.77 | בשם |  |
| 2:4.78 | בית |  |
| 2:4.79 | אחד | . |
| 2:4.80 | ואם |  |
| 2:4.81 | תכנס |  |
| 2:4.82 | בו |  |
| 2:4.83 | תמצא |  |
| 2:4.84 | כמה |  |
| 2:4.85 | וכמה |  |
| 2:4.86 | חדרים |  |
| 2:4.87 | גדולים |  |
| 2:4.88 | וקטנים |  |
| 2:4.89 | וכמה |  |
| 2:4.90 | מקומות |  |
| 2:4.91 | זולתם | . |
| 2:4.92 | וכל |  |
| 2:4.93 | א | ' |
| 2:4.94 | ואחד |  |
| 2:4.95 | מתארים |  |
| 2:4.96 | לו |  |
| 2:4.97 | תואר |  |
| 2:4.98 | שם |  |
| 2:4.99 | אחד | . |
| 2:4.100 | גם |  |
| 2:4.101 | הבנין |  |
| 2:4.102 | של |  |
| 2:4.103 | עצים |  |
| 2:4.104 | שיש |  |
| 2:4.105 | בתוכו |  |
| 2:4.106 | כמה |  |
| 2:4.107 | תארים |  |
| 2:4.108 | יש |  |
| 2:4.109 | בו | . |
| 2:4.110 | גם |  |
| 2:4.111 | תמצא |  |
| 2:4.112 | עוד |  |
| 2:4.113 | שזה |  |
| 2:4.114 | הכותל |  |
| 2:4.115 | קדש |  |
| 2:4.116 | שנבנה |  |
| 2:4.117 | היה |  |
| 2:4.118 | חלקים |  |
| 2:4.119 | כל |  |
| 2:4.120 | אבן |  |
| 2:4.121 | ואבן |  |
| 2:4.122 | היא |  |
| 2:4.123 | חלוקה |  |
| 2:4.124 | מצרורות |  |
| 2:4.125 | העפר |  |
| 2:4.126 | לעצמו |  |
| 2:4.127 | והסיד |  |
| 2:4.128 | לעצמו |  |
| 2:4.129 | כו | '. |
| 2:4.130 | ואחר |  |
| 2:4.131 | שבא |  |
| 2:4.132 | הבנאי |  |
| 2:4.133 | וחבר |  |
| 2:4.134 | כל |  |
| 2:4.135 | חלקים |  |
| 2:4.136 | אלו |  |
| 2:4.137 | בחכמתו |  |
| 2:4.138 | זה |  |
| 2:4.139 | בתוך |  |
| 2:4.140 | זה |  |
| 2:4.141 | וזה |  |
| 2:4.142 | אצל |  |
| 2:4.143 | זה |  |
| 2:4.144 | ונעשו |  |
| 2:4.145 | כותל |  |
| 2:4.146 | אחד |  |
| 2:4.147 | וכו | ', |
| 2:4.148 | נמצא |  |
| 2:4.149 | אז |  |
| 2:4.150 | הכל |  |
| 2:4.151 | חתיכה |  |
| 2:4.152 | אחת |  |
| 2:4.153 | וחלק |  |
| 2:4.154 | אחד | . |
| 2:4.155 | וכן |  |
| 2:4.156 | הענין |  |
| 2:4.157 | אצלנו |  |
| 2:4.158 | וכו | '. |
| 2:4.159 | ושם |  |
| 2:4.160 | בדף |  |
| 2:4.161 | כ"ה |  |
| 2:4.162 | וכ"ט |  |
| 2:4.163 | מדמה |  |
| 2:4.164 | אותו |  |
| 2:4.165 | לגוף |  |
| 2:4.166 | האדם |  |
| 2:4.167 | המחובר |  |
| 2:4.168 | מעצמות |  |
| 2:4.169 | וגידים |  |
| 2:4.170 | ובשר |  |
| 2:4.171 | וראש |  |
| 2:4.172 | ועין |  |
| 2:4.173 | ואזן |  |
| 2:4.174 | וחוטם |  |
| 2:4.175 | ופה |  |
| 2:4.176 | וידים |  |
| 2:4.177 | ורגלים |  |
| 2:4.178 | והכל |  |
| 2:4.179 | נקרא |  |
| 2:4.180 | בשם |  |
| 2:4.181 | אדם |  |
| 2:4.182 | ראובן |  |
| 2:4.183 | או |  |
| 2:4.184 | שמעון | . |
| 2:4.185 | ובדף |  |
| 2:4.186 | כ"ה |  |
| 2:4.187 | וכ"ו |  |
| 2:4.188 | וכ"ז |  |
| 2:4.189 | כתב | : |
| 2:4.190 | הכלל |  |
| 2:4.191 | העולה |  |
| 2:4.192 | כי |  |
| 2:4.193 | הסבה |  |
| 2:4.194 | הראשונה |  |
| 2:4.195 | הנקרא |  |
| 2:4.196 | בפי |  |
| 2:4.197 | כל |  |
| 2:4.198 | המקובלים |  |
| 2:4.199 | אין |  |
| 2:4.200 | סוף | , |
| 2:4.201 | הוא |  |
| 2:4.202 | אשר |  |
| 2:4.203 | האציל |  |
| 2:4.204 | וברא |  |
| 2:4.205 | ויצר |  |
| 2:4.206 | ועשה |  |
| 2:4.207 | והוא |  |
| 2:4.208 | מסתתר |  |
| 2:4.209 | תוך |  |
| 2:4.210 | זעיר |  |
| 2:4.211 | אנפין |  |
| 2:4.212 | ואזי |  |
| 2:4.213 | זעיר |  |
| 2:4.214 | אנפין |  |
| 2:4.215 | הוא |  |
| 2:4.216 | השולט |  |
| 2:4.217 | על |  |
| 2:4.218 | כל |  |
| 2:4.219 | הברואים |  |
| 2:4.220 | ומנהיגם |  |
| 2:4.221 | וזנם |  |
| 2:4.222 | ומפרנסם |  |
| 2:4.223 | בכח |  |
| 2:4.224 | אין |  |
| 2:4.225 | סוף |  |
| 2:4.226 | שבו | . |
| 2:4.227 | ולכן |  |
| 2:4.228 | הוא |  |
| 2:4.229 | אלהינו |  |
| 2:4.230 | ואנחנו |  |
| 2:4.231 | עמו |  |
| 2:4.232 | כי |  |
| 2:4.233 | נשמותינו |  |
| 2:4.234 | הם |  |
| 2:4.235 | חלקו |  |
| 2:4.236 | ואותו |  |
| 2:4.237 | נעבוד |  |
| 2:4.238 | והוא |  |
| 2:4.239 | אלהי |  |
| 2:4.240 | אבותינו |  |
| 2:4.241 | שבידו |  |
| 2:4.242 | הנהגת |  |
| 2:4.243 | העולמות |  |
| 2:4.244 | בשכר |  |
| 2:4.245 | ועונש |  |
| 2:4.246 | אבל |  |
| 2:4.247 | האין |  |
| 2:4.248 | סוף |  |
| 2:4.249 | עם |  |
| 2:4.250 | שאר |  |
| 2:4.251 | הפרצופים |  |
| 2:4.252 | שלמעלה |  |
| 2:4.253 | מזעיר |  |
| 2:4.254 | ונוקביה |  |
| 2:4.255 | אם |  |
| 2:4.256 | יסדרו |  |
| 2:4.257 | בני |  |
| 2:4.258 | אדם |  |
| 2:4.259 | תפלתם |  |
| 2:4.260 | להם |  |
| 2:4.261 | בפרטות |  |
| 2:4.262 | מבלי |  |
| 2:4.263 | שיתפללו |  |
| 2:4.264 | לזעיר |  |
| 2:4.265 | אנפין |  |
| 2:4.266 | גם |  |
| 2:4.267 | אם |  |
| 2:4.268 | יכוונו |  |
| 2:4.269 | לנשמה |  |
| 2:4.270 | המסתתרת |  |
| 2:4.271 | בהם |  |
| 2:4.272 | יהיו |  |
| 2:4.273 | צועקים |  |
| 2:4.274 | ואינם |  |
| 2:4.275 | נענים | . |
| 2:4.276 | ואדרבה |  |
| 2:4.277 | יהיו |  |
| 2:4.278 | המתפללים |  |
| 2:4.279 | אליהם |  |
| 2:4.280 | נענשים |  |
| 2:4.281 | כי |  |
| 2:4.282 | זהו |  |
| 2:4.283 | רצון |  |
| 2:4.284 | הסבה |  |
| 2:4.285 | הראשונה |  |
| 2:4.286 | שיהיה |  |
| 2:4.287 | הוא |  |
| 2:4.288 | המוציא |  |
| 2:4.289 | והמביא |  |
| 2:4.290 | שפעו |  |
| 2:4.291 | לתחתונים |  |
| 2:4.292 | ואין |  |
| 2:4.293 | עוד |  |
| 2:4.294 | מלבדו |  |
| 2:4.295 | ע"ש |  |
| 2:4.296 | היטב | . |
| 2:4.297 | ודברי |  |
| 2:4.298 | הר | ' |
| 2:4.299 | כסא |  |
| 2:4.300 | אליהו |  |
| 2:4.301 | עומדים |  |
| 2:4.302 | נגד |  |
| 2:4.303 | דברי |  |
| 2:4.304 | הרמב"ם |  |
| 2:4.305 | ז"ל |  |
| 2:4.306 | בפי | ' |
| 2:4.307 | המשנה |  |
| 2:4.308 | ובחבור |  |
| 2:4.309 | ובמורה |  |
| 2:4.310 | נבוכים |  |
| 2:4.311 | ונגד |  |
| 2:4.312 | דברי |  |
| 2:4.313 | החסיד |  |
| 2:4.314 | בעל |  |
| 2:4.315 | חובת |  |
| 2:4.316 | הלבבות |  |
| 2:4.317 | בשער |  |
| 2:4.318 | היחוד |  |
| 2:4.319 | והרס"ג |  |
| 2:4.320 | בס | ' |
| 2:4.321 | האמונות |  |
| 2:4.322 | והדעות |  |
| 2:4.323 | והרוקח |  |
| 2:4.324 | שכתבו |  |
| 2:4.325 | שאחדותו |  |
| 2:4.326 | ית | ' |
| 2:4.327 | איננו |  |
| 2:4.328 | כאחד |  |
| 2:4.329 | הזוג |  |
| 2:4.330 | ולא |  |
| 2:4.331 | כאחד |  |
| 2:4.332 | המין |  |
| 2:4.333 | ולא |  |
| 2:4.334 | כאיש |  |
| 2:4.335 | האחד |  |
| 2:4.336 | שנחלק |  |
| 2:4.337 | לאחדים |  |
| 2:4.338 | רבים | . |
| 2:4.339 | ולא |  |
| 2:4.340 | אחד |  |
| 2:4.341 | כמו |  |
| 2:4.342 | הגוף |  |
| 2:4.343 | הפשוט |  |
| 2:4.344 | וכו | ' |
| 2:4.345 | שמקבל |  |
| 2:4.346 | החלוק |  |
| 2:4.347 | לאין |  |
| 2:4.348 | סוף | . |
| 2:4.349 | אבל |  |
| 2:4.350 | השי"ת |  |
| 2:4.351 | הוא |  |
| 2:4.352 | אחד |  |
| 2:4.353 | באחדות |  |
| 2:4.354 | שאין |  |
| 2:4.355 | כמותה | . |
| 2:4.356 | ג"כ |  |
| 2:4.357 | כתב |  |
| 2:4.358 | כסא |  |
| 2:4.359 | אליהו |  |
| 2:4.360 | שהאין |  |
| 2:4.361 | סוף |  |
| 2:4.362 | הוא |  |
| 2:4.363 | נשמת |  |
| 2:4.364 | עתיק | , |
| 2:4.365 | וארוך |  |
| 2:4.366 | אנפין | , |
| 2:4.367 | ואבא |  |
| 2:4.368 | ואימא | , |
| 2:4.369 | וזעיר |  |
| 2:4.370 | ונוקביה | , |
| 2:4.371 | ורבותי |  |
| 2:4.372 | ז"ל |  |
| 2:4.373 | הנז"ל |  |
| 2:4.374 | כתבו |  |
| 2:4.375 | שאינו |  |
| 2:4.376 | גוף |  |
| 2:4.377 | ולא |  |
| 2:4.378 | כח |  |
| 2:4.379 | בגוף | ! |
| 2:4.380 | ולפי |  |
| 2:4.381 | דבריו |  |
| 2:4.382 | הוא |  |
| 2:4.383 | כח |  |
| 2:4.384 | בגוף | . |
| 2:4.385 | ג"כ |  |
| 2:4.386 | כתב |  |
| 2:4.387 | שאריך |  |
| 2:4.388 | אנפין |  |
| 2:4.389 | אבא |  |
| 2:4.390 | ואימא |  |
| 2:4.391 | קדמו |  |
| 2:4.392 | לזעיר |  |
| 2:4.393 | אנפין |  |
| 2:4.394 | שהוא |  |
| 2:4.395 | אלהינו |  |
| 2:4.396 | לפי |  |
| 2:4.397 | דבריו | ) |
| 2:4.398 | והוא |  |
| 2:4.399 | נקרא |  |
| 2:4.400 | בן |  |
| 2:4.401 | לאבא |  |
| 2:4.402 | ואימא |  |
| 2:4.403 | וכו | '. |
| 2:4.404 | ורז"ל |  |
| 2:4.405 | אמרו |  |
| 2:4.406 | שהוא |  |
| 2:4.407 | קדמון |  |
| 2:4.408 | לכל |  |
| 2:4.409 | הנמצאים |  |
| 2:4.410 | ושכל |  |
| 2:4.411 | נמצא |  |
| 2:4.412 | זולתו |  |
| 2:4.413 | בלתי |  |
| 2:4.414 | קדמון |  |
| 2:4.415 | וכו | '. |
| 2:4.416 | ג"כ |  |
| 2:4.417 | כתב |  |
| 2:4.418 | שאין |  |
| 2:4.419 | סוף |  |
| 2:4.420 | ב"ה |  |
| 2:4.421 | לא |  |
| 2:4.422 | שייכא |  |
| 2:4.423 | שום |  |
| 2:4.424 | עבודה |  |
| 2:4.425 | בו |  |
| 2:4.426 | ולא |  |
| 2:4.427 | שום |  |
| 2:4.428 | קריאה |  |
| 2:4.429 | כלל | . |
| 2:4.430 | ולא |  |
| 2:4.431 | שייך |  |
| 2:4.432 | בו |  |
| 2:4.433 | שום |  |
| 2:4.434 | שם |  |
| 2:4.435 | לא |  |
| 2:4.436 | הוי"ה |  |
| 2:4.437 | ולא |  |
| 2:4.438 | אדנות |  |
| 2:4.439 | וכו | ' |
| 2:4.440 | רק |  |
| 2:4.441 | לז"א |  |
| 2:4.442 | ונוקביה |  |
| 2:4.443 | שבעולם |  |
| 2:4.444 | האצילות |  |
| 2:4.445 | בלבד |  |
| 2:4.446 | ולא |  |
| 2:4.447 | לא"א |  |
| 2:4.448 | ואבא |  |
| 2:4.449 | ואימא |  |
| 2:4.450 | וזעיר |  |
| 2:4.451 | ונוקביה |  |
| 2:4.452 | שיש |  |
| 2:4.453 | בעולמות |  |
| 2:4.454 | הנאצלים |  |
| 2:4.455 | למעלה |  |
| 2:4.456 | מעולם |  |
| 2:4.457 | דאצילות |  |
| 2:4.458 | כמבואר |  |
| 2:4.459 | שם |  |
| 2:4.460 | ד | ' |
| 2:4.461 | מ"ט |  |
| 2:4.462 | סוף |  |
| 2:4.463 | עמוד |  |
| 2:4.464 | א | ' |
| 2:4.465 | ובדף |  |
| 2:4.466 | נ"ג | .... |


### Translation under review (2:4)

English: The questioner replied again. Your reply, my friend, came closed and sealed, and in it there is nothing but an indication of a place: "to study in Kisse Eliyahu, from folio 13 to folio 18, 'A wise man will hear and add to learning,' and what we have received is that worship is to God, one Lord, etc." I did not understand your opinion and your intention in this, whether this, what you have received, is in accordance with the words of Kisse Eliyahu or the opposite of his words. For when I studied in Kisse Eliyahu the place you marked for me, I did not find what I sought; on the contrary, there, on folio 3, he compares the oneness of the Creator to the other kinds of oneness, where he wrote: "Just as this house as a whole is called by the name one house. And if you enter it, you will find many rooms, large and small, and many other places. And each and every one of them ascribes to it one name. Also, a wooden building that has several attributes within it has them. Also, you will find further that this holy wall, which was built, was in parts; each and every stone was separate, of clods of earth by itself and of lime by itself, etc. And after the builder came and joined all these parts with his wisdom, this within that, and this beside that, and they became one wall, etc., Thus, then, all is one piece and one part." And so is the matter with us, etc. And there, on folios 25 and 29, he compares it to the human body, joined of bones and sinews and flesh, head and eye and ear and nose and mouth and hands and feet, and the whole is called by the name man, Reuben or Simeon. And on folios 25, 26 and 27 he wrote: "The sum of the matter is that the First Cause, which all the Kabbalists call Ein Sof, is He who emanated, created, formed and made, and He is hidden within Ze'ir Anpin, and then Ze'ir Anpin is the one who rules over all creatures, leads them, feeds them and sustains them, by the power of Ein Sof within him. Therefore He is our God, and we are His people, for our souls are His portion, and Him we shall worship, and He is the God of our fathers, in whose hand is the conduct of the worlds, for reward and punishment. But Ein Sof, with the other partzufim above Ze'ir and his Nukva, if human beings order their prayers to them individually, without praying to Ze'ir Anpin, even if they intend the soul hidden in them, they will cry out and not be answered. And on the contrary, those who pray to them will be punished, for it is the will of the First Cause that He be the one who brings out and brings in His flow to the lower beings, and there is no other besides Him. See there carefully." And the words of Rabbi Kisse Eliyahu stand against the words of the Rambam, of blessed memory, in his commentary on the Mishnah, in his composition, and in Moreh Nevukhim, and against the words of the pious one, the author of Chovot ha-Levavot, in the Gate of Unity, and Rabbi Saadia Gaon in the book Emunot ve-De'ot, and the Rokeach, who wrote that His oneness, may He be blessed, is not like the oneness of a pair, nor like the oneness of a kind, nor like the one man who is divided into many ones. and not one like a simple body, etc., which receives division to Ein Sof. But the Name, may He be blessed, is one with a oneness that has no like. Also, Kisse Eliyahu wrote that Ein Sof is the soul of Atik, and Arikh Anpin, and Abba and Imma, and Ze'ir and his Nukva, and my teachers, of blessed memory, the aforesaid, wrote that He is not a body, nor a force in a body! And according to his words He is a force in a body. Also he wrote that Arikh Anpin, Abba and Imma preceded Ze'ir Anpin who is our God (according to his words), and he is called a son of Abba and Imma, etc. And our Rabbis of blessed memory said that He is eternal before all that exists, and that every existent other than Him is not eternal, etc. Also he wrote concerning Ein Sof, blessed be He, that no worship pertains to Him, and no invocation at all. And no name applies to Him: not YHVH, nor Adonai, etc. but only to Ze'ir Anpin and his Nukva in the World of Emanation alone, and not to Arikh Anpin, Abba and Imma, and Ze'ir and his Nukva that are in the emanated worlds above the World of Emanation, as explained there, 4, 49, the end of side a, and on folio 53....

Literal: again the-questioner replied your-reply my-friend came closed and-sealed and-there-is-not in-it but an-indication of-a-place to-study in-Kisse Eliyahu from-folio 13 to folio 18 will-hear a-wise-man and-will-add to-learning and-what-is-received by-us that-worship to-God lord one etc. and-not I-understood your-opinion and-your-intention in-this whether this what-is-received by-you is in-accordance-with-the-words-of Kisse Eliyahu or the-opposite-of his-words for when I-studied in-Kisse Eliyahu the-place you-marked for-me not I-found what-I-sought on-the-contrary there on-folio 3 he-compares the-oneness of-the-Creator to-the-other kinds of-oneness where-there he-wrote just-as this the-house as-a-whole is-called by-the-name house one and-if you-enter in-it you-will-find some and-some rooms large and-small and-some places other-than-them and-each one and-one they-ascribe to-it a-name-attribute name one also the-building of wood that-has within-it several attributes has in-it also you-will-find further that-this the-wall holy which-was-built was in-parts each stone and-stone it was-divided of-clods the-earth by-itself and-of-lime by-itself etc. and-after came the-builder and-joined all parts these with-his-wisdom this within that and-this beside that and-became a-wall one etc. it-is-found then all a-piece one and-a-part one and-so the-matter with-us etc. and-there on-folio 25 and-29 he-compares it to-the-body of-man joined of-bones and-sinews and-flesh and-head and-eye and-ear and-nose and-mouth and-hands and-feet and-all is-called by-the-name man Reuben or Simeon and-on-folio 25 and-26 and-27 he-wrote the-sum that-emerges that the-cause the-first which-is-called by-the-mouth-of all the-Kabbalists without end he who emanated and-created and-formed and-made and-he is-hidden within Ze'ir Anpin and-then Ze'ir Anpin he who-rules over all creatures and-leads-them and-feeds-them and-sustains-them by-the-power Ein Sof within-him therefore He our-God and-we his-people for our-souls they His-portion and-Him we-shall-worship and-he God-of our-fathers in-whose-hand the-conduct the-worlds for-reward and-punishment but the-Ein Sof with the-rest the-partzufim that-are-above Ze'ir and-his-Nukva if they-arrange human beings their-prayer to-them individually without praying to-Ze'ir Anpin even if they-intend to-the-soul hidden in-them they-will-be crying-out and-not answered and-on-the-contrary will-be those-who-pray to-them punished for this-is the-will of-the-cause the-first that he the-one-who-brings-out and-brings-in His-flow to-the-lower-beings and-there-is-not other besides-Him see-there carefully and-the-words-of Rabbi Kisse Eliyahu stand against the-words-of the-Rambam of-blessed-memory in-his-commentary on-the-Mishnah and-in-his-composition and-in-Moreh Nevukhim and-against the-words-of the-pious-one the-author-of Chovot ha-Levavot in-the-gate-of Unity and-Rabbi-Saadia-Gaon in-the-book Emunot and-De'ot and-the-Rokeach who-wrote that-His-oneness may-He-be-blessed is-not like-the-oneness-of the-pair nor like-the-oneness-of the-kind nor like-the-man the-one who-is-divided into-ones many and-not one like the-body simple etc. which-receives the-division to-Ein Sof but the-Name,-may-He-be-blessed He one with-an-oneness that-has-no like-it also wrote Kisse Eliyahu that-the-Ein Sof is the-soul-of Atik and-Arikh Anpin and-Abba and-Imma and-Ze'ir and-his-Nukva and-my-teachers of-blessed-memory the-aforesaid wrote that-He-is-not a-body nor a-force in-a-body and-according-to his-words He a-force in-a-body also he-wrote that-Arikh Anpin Abba and-Imma preceded Ze'ir Anpin who is-our-God according-to his-words and-he is-called son of-Abba and-Imma etc. and-our-Rabbis-of-blessed-memory said that-He eternal before-all the-existents and-that-every existent other-than-Him not eternal etc. also he-wrote concerning-Ein Sof blessed-be-He not pertains any worship to-Him and-not any calling at-all and-not applies to-Him any name not YHVH nor Adonai etc. but-only to-Ze'ir-Anpin and-his-Nukva in-the-world of-Emanation alone and-not to-Arikh-Anpin and-Abba and-Imma and-Ze'ir and-his-Nukva that-are in-the-worlds the-emanated above the-world-of Emanation as-explained there 4 49 end-of side a and-on-folio 53

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:4.1 | again |  |  |  |
| 2:4.2 | the questioner |  |  |  |
| 2:4.3 | replied |  |  |  |
| 2:4.4 | your reply |  |  |  |
| 2:4.5 | my friend |  |  |  |
| 2:4.6 | came |  |  |  |
| 2:4.7 | closed |  |  |  |
| 2:4.8 | and sealed |  |  |  |
| 2:4.9 | and there is not |  |  |  |
| 2:4.10 | in it |  |  |  |
| 2:4.11 | but |  |  |  |
| 2:4.12 | an indication |  |  |  |
| 2:4.13 | of a place |  |  |  |
| 2:4.14 | to study |  |  |  |
| 2:4.15 | in Kisse |  |  |  |
| 2:4.16 | Eliyahu |  |  |  |
| 2:4.17 | from folio |  |  |  |
| 2:4.18 | 13 |  |  |  |
| 2:4.19 | to |  |  |  |
| 2:4.20 | folio |  |  |  |
| 2:4.21 | 18 |  |  |  |
| 2:4.22 | will hear |  |  |  |
| 2:4.23 | a wise man |  |  |  |
| 2:4.24 | and will add |  |  |  |
| 2:4.25 | to learning |  |  |  |
| 2:4.26 | and what is received |  |  |  |
| 2:4.27 | by us |  |  |  |
| 2:4.28 | that worship |  |  |  |
| 2:4.29 | to God |  |  |  |
| 2:4.30 | lord |  |  |  |
| 2:4.31 | one |  |  |  |
| 2:4.32 | etc. |  |  | וכו׳ = etc.: the quotation is cut short and 'etc.' is kept (decisions) |
| 2:4.33 | and not |  |  |  |
| 2:4.34 | I understood |  |  |  |
| 2:4.35 | your opinion |  |  |  |
| 2:4.36 | and your intention |  |  |  |
| 2:4.37 | in this |  |  |  |
| 2:4.38 | whether |  |  |  |
| 2:4.39 | this |  |  |  |
| 2:4.40 | what is received |  |  |  |
| 2:4.41 | by you |  |  |  |
| 2:4.42 | is |  |  |  |
| 2:4.43 | in accordance with the words of |  |  |  |
| 2:4.44 | Kisse |  |  |  |
| 2:4.45 | Eliyahu |  |  |  |
| 2:4.46 | or |  |  |  |
| 2:4.47 | the opposite of |  |  |  |
| 2:4.48 | his words |  |  |  |
| 2:4.49 | for |  |  |  |
| 2:4.50 | when |  |  |  |
| 2:4.51 | I studied |  |  |  |
| 2:4.52 | in Kisse |  |  |  |
| 2:4.53 | Eliyahu |  |  |  |
| 2:4.54 | the place |  |  |  |
| 2:4.55 | you marked |  |  |  |
| 2:4.56 | for me |  |  |  |
| 2:4.57 | not |  |  |  |
| 2:4.58 | I found |  |  |  |
| 2:4.59 | what I sought |  |  |  |
| 2:4.60 | on the contrary |  |  |  |
| 2:4.61 | there |  |  | שם read as 'there', the place in the book; see apparatus |
| 2:4.62 | on folio |  |  |  |
| 2:4.63 | 3 |  |  | ג׳ is folio 3 |
| 2:4.64 | he compares |  |  |  |
| 2:4.65 | the oneness |  |  | אחדות (oneness) is kept distinct from יחוד (unity) |
| 2:4.66 | of the Creator |  |  |  |
| 2:4.67 | to the other |  |  |  |
| 2:4.68 | kinds |  |  |  |
| 2:4.69 | of oneness |  |  |  |
| 2:4.70 | where there |  |  |  |
| 2:4.71 | he wrote |  |  |  |
| 2:4.72 | just as |  |  |  |
| 2:4.73 | this |  |  |  |
| 2:4.74 | the house |  |  |  |
| 2:4.75 | as a whole |  |  |  |
| 2:4.76 | is called |  |  |  |
| 2:4.77 | by the name |  |  |  |
| 2:4.78 | house |  |  |  |
| 2:4.79 | one |  |  |  |
| 2:4.80 | and if |  |  |  |
| 2:4.81 | you enter |  |  |  |
| 2:4.82 | in it |  |  |  |
| 2:4.83 | you will find |  |  |  |
| 2:4.84 | some |  |  |  |
| 2:4.85 | and some |  |  |  |
| 2:4.86 | rooms |  |  |  |
| 2:4.87 | large |  |  |  |
| 2:4.88 | and small |  |  |  |
| 2:4.89 | and some |  |  |  |
| 2:4.90 | places |  |  |  |
| 2:4.91 | other than them |  |  |  |
| 2:4.92 | and each |  |  |  |
| 2:4.93 | one |  |  | א׳ is written with an apostrophe; in כל א׳ ואחד it is 'one', rendered 'each and every one' |
| 2:4.94 | and one |  |  |  |
| 2:4.95 | they ascribe |  |  |  |
| 2:4.96 | to it |  |  |  |
| 2:4.97 | a name-attribute |  |  |  |
| 2:4.98 | name |  |  |  |
| 2:4.99 | one |  |  |  |
| 2:4.100 | also |  |  |  |
| 2:4.101 | the building |  |  |  |
| 2:4.102 | of |  |  |  |
| 2:4.103 | wood |  |  |  |
| 2:4.104 | that has |  |  |  |
| 2:4.105 | within it |  |  |  |
| 2:4.106 | several |  |  |  |
| 2:4.107 | attributes |  |  |  |
| 2:4.108 | has |  |  |  |
| 2:4.109 | in it |  |  |  |
| 2:4.110 | also |  |  |  |
| 2:4.111 | you will find |  |  |  |
| 2:4.112 | further |  |  |  |
| 2:4.113 | that this |  |  |  |
| 2:4.114 | the wall |  |  |  |
| 2:4.115 | holy |  |  |  |
| 2:4.116 | which was built |  |  |  |
| 2:4.117 | was |  |  |  |
| 2:4.118 | in parts |  |  |  |
| 2:4.119 | each |  |  |  |
| 2:4.120 | stone |  |  |  |
| 2:4.121 | and stone |  |  |  |
| 2:4.122 | it |  |  |  |
| 2:4.123 | was divided |  |  |  |
| 2:4.124 | of clods |  |  |  |
| 2:4.125 | the earth |  |  |  |
| 2:4.126 | by itself |  |  |  |
| 2:4.127 | and of lime |  |  |  |
| 2:4.128 | by itself |  |  |  |
| 2:4.129 | etc. |  |  |  |
| 2:4.130 | and after |  |  |  |
| 2:4.131 | came |  |  |  |
| 2:4.132 | the builder |  |  |  |
| 2:4.133 | and joined |  |  |  |
| 2:4.134 | all |  |  |  |
| 2:4.135 | parts |  |  |  |
| 2:4.136 | these |  |  |  |
| 2:4.137 | with his wisdom |  |  |  |
| 2:4.138 | this |  |  |  |
| 2:4.139 | within |  |  |  |
| 2:4.140 | that |  |  |  |
| 2:4.141 | and this |  |  |  |
| 2:4.142 | beside |  |  |  |
| 2:4.143 | that |  |  |  |
| 2:4.144 | and became |  |  |  |
| 2:4.145 | a wall |  |  |  |
| 2:4.146 | one |  |  |  |
| 2:4.147 | etc. |  |  |  |
| 2:4.148 | it is found |  |  |  |
| 2:4.149 | then |  |  |  |
| 2:4.150 | all |  |  |  |
| 2:4.151 | a piece |  |  |  |
| 2:4.152 | one |  |  |  |
| 2:4.153 | and a part |  |  |  |
| 2:4.154 | one |  |  |  |
| 2:4.155 | and so |  |  |  |
| 2:4.156 | the matter |  |  |  |
| 2:4.157 | with us |  |  |  |
| 2:4.158 | etc. |  |  |  |
| 2:4.159 | and there |  |  |  |
| 2:4.160 | on folio |  |  |  |
| 2:4.161 | 25 |  |  |  |
| 2:4.162 | and 29 |  |  |  |
| 2:4.163 | he compares |  |  |  |
| 2:4.164 | it |  |  |  |
| 2:4.165 | to the body |  |  |  |
| 2:4.166 | of man |  |  |  |
| 2:4.167 | joined |  |  |  |
| 2:4.168 | of bones |  |  |  |
| 2:4.169 | and sinews |  |  |  |
| 2:4.170 | and flesh |  |  |  |
| 2:4.171 | and head |  |  |  |
| 2:4.172 | and eye |  |  |  |
| 2:4.173 | and ear |  |  |  |
| 2:4.174 | and nose |  |  |  |
| 2:4.175 | and mouth |  |  |  |
| 2:4.176 | and hands |  |  |  |
| 2:4.177 | and feet |  |  |  |
| 2:4.178 | and all |  |  |  |
| 2:4.179 | is called |  |  |  |
| 2:4.180 | by the name |  |  |  |
| 2:4.181 | man |  |  |  |
| 2:4.182 | Reuben |  |  |  |
| 2:4.183 | or |  |  |  |
| 2:4.184 | Simeon |  |  |  |
| 2:4.185 | and on folio |  |  |  |
| 2:4.186 | 25 |  |  |  |
| 2:4.187 | and 26 |  |  |  |
| 2:4.188 | and 27 |  |  |  |
| 2:4.189 | he wrote |  |  |  |
| 2:4.190 | the sum |  |  |  |
| 2:4.191 | that emerges |  |  |  |
| 2:4.192 | that |  |  |  |
| 2:4.193 | the cause |  |  |  |
| 2:4.194 | the first |  |  |  |
| 2:4.195 | which is called |  |  |  |
| 2:4.196 | by the mouth of |  |  |  |
| 2:4.197 | all |  |  |  |
| 2:4.198 | the Kabbalists |  |  |  |
| 2:4.199 | without |  |  |  |
| 2:4.200 | end |  |  |  |
| 2:4.201 | he |  |  |  |
| 2:4.202 | who |  |  |  |
| 2:4.203 | emanated |  |  | האציל, the verb of אצילות (emanation), kept distinct from ברא (create) |
| 2:4.204 | and created |  |  |  |
| 2:4.205 | and formed |  |  |  |
| 2:4.206 | and made |  |  |  |
| 2:4.207 | and he |  |  |  |
| 2:4.208 | is hidden |  |  |  |
| 2:4.209 | within |  |  |  |
| 2:4.210 | Ze'ir |  |  | זעיר: 'small', from Ze'ir Anpin, literally 'small face' |
| 2:4.211 | Anpin |  |  |  |
| 2:4.212 | and then |  |  |  |
| 2:4.213 | Ze'ir |  |  |  |
| 2:4.214 | Anpin |  |  |  |
| 2:4.215 | he |  |  |  |
| 2:4.216 | who rules |  |  |  |
| 2:4.217 | over |  |  |  |
| 2:4.218 | all |  |  |  |
| 2:4.219 | creatures |  |  |  |
| 2:4.220 | and leads them |  |  |  |
| 2:4.221 | and feeds them |  |  |  |
| 2:4.222 | and sustains them |  |  |  |
| 2:4.223 | by the power |  |  |  |
| 2:4.224 | Ein |  |  |  |
| 2:4.225 | Sof |  |  |  |
| 2:4.226 | within him |  |  |  |
| 2:4.227 | therefore |  |  |  |
| 2:4.228 | He |  |  |  |
| 2:4.229 | our God |  |  |  |
| 2:4.230 | and we |  |  |  |
| 2:4.231 | his people |  |  |  |
| 2:4.232 | for |  |  |  |
| 2:4.233 | our souls |  |  |  |
| 2:4.234 | they |  |  |  |
| 2:4.235 | His portion |  |  |  |
| 2:4.236 | and Him |  |  |  |
| 2:4.237 | we shall worship |  |  |  |
| 2:4.238 | and he |  |  |  |
| 2:4.239 | God of |  |  |  |
| 2:4.240 | our fathers |  |  |  |
| 2:4.241 | in whose hand |  |  |  |
| 2:4.242 | the conduct |  |  |  |
| 2:4.243 | the worlds |  |  |  |
| 2:4.244 | for reward |  |  |  |
| 2:4.245 | and punishment |  |  |  |
| 2:4.246 | but |  |  |  |
| 2:4.247 | the Ein |  |  |  |
| 2:4.248 | Sof |  |  |  |
| 2:4.249 | with |  |  |  |
| 2:4.250 | the rest |  |  |  |
| 2:4.251 | the partzufim |  |  |  |
| 2:4.252 | that are above |  |  |  |
| 2:4.253 | Ze'ir |  |  |  |
| 2:4.254 | and his Nukva |  |  | ונוקביה: his Nukva, literally 'female' |
| 2:4.255 | if |  |  |  |
| 2:4.256 | they arrange |  |  |  |
| 2:4.257 | human |  |  |  |
| 2:4.258 | beings |  |  |  |
| 2:4.259 | their prayer |  |  |  |
| 2:4.260 | to them |  |  |  |
| 2:4.261 | individually |  |  |  |
| 2:4.262 | without |  |  |  |
| 2:4.263 | praying |  |  |  |
| 2:4.264 | to Ze'ir |  |  |  |
| 2:4.265 | Anpin |  |  |  |
| 2:4.266 | even |  |  |  |
| 2:4.267 | if |  |  |  |
| 2:4.268 | they intend |  |  |  |
| 2:4.269 | to the soul |  |  |  |
| 2:4.270 | hidden |  |  |  |
| 2:4.271 | in them |  |  |  |
| 2:4.272 | they will be |  |  |  |
| 2:4.273 | crying out |  |  |  |
| 2:4.274 | and not |  |  |  |
| 2:4.275 | answered |  |  |  |
| 2:4.276 | and on the contrary |  |  |  |
| 2:4.277 | will be |  |  |  |
| 2:4.278 | those who pray |  |  |  |
| 2:4.279 | to them |  |  |  |
| 2:4.280 | punished |  |  |  |
| 2:4.281 | for |  |  |  |
| 2:4.282 | this is |  |  |  |
| 2:4.283 | the will |  |  |  |
| 2:4.284 | of the cause |  |  |  |
| 2:4.285 | the first |  |  |  |
| 2:4.286 | that |  |  |  |
| 2:4.287 | he |  |  |  |
| 2:4.288 | the one who brings out |  |  |  |
| 2:4.289 | and brings in |  |  |  |
| 2:4.290 | His flow |  |  |  |
| 2:4.291 | to the lower beings |  |  |  |
| 2:4.292 | and there is not |  |  |  |
| 2:4.293 | other |  |  |  |
| 2:4.294 | besides Him |  |  |  |
| 2:4.295 | see there |  |  |  |
| 2:4.296 | carefully |  |  |  |
| 2:4.297 | and the words of |  |  |  |
| 2:4.298 | Rabbi |  |  | הר׳ = הרב, 'Rabbi', here used of the author of Kisse Eliyahu; the decisions give no ruling on this usage |
| 2:4.299 | Kisse |  |  |  |
| 2:4.300 | Eliyahu |  |  |  |
| 2:4.301 | stand |  |  |  |
| 2:4.302 | against |  |  |  |
| 2:4.303 | the words of |  |  |  |
| 2:4.304 | the Rambam |  |  | הרמב״ם: the Rambam, of blessed memory |
| 2:4.305 | of blessed memory |  |  |  |
| 2:4.306 | in his commentary |  |  |  |
| 2:4.307 | on the Mishnah |  |  |  |
| 2:4.308 | and in his composition |  |  | ובחבור: 'in his composition'; the text does not name the work |
| 2:4.309 | and in Moreh |  |  | מורה נבוכים: Moreh Nevukhim, literally 'Guide of the Perplexed' |
| 2:4.310 | Nevukhim |  |  |  |
| 2:4.311 | and against |  |  |  |
| 2:4.312 | the words of |  |  |  |
| 2:4.313 | the pious one |  |  |  |
| 2:4.314 | the author of |  |  |  |
| 2:4.315 | Chovot |  |  | חובת הלבבות: Chovot ha-Levavot, literally 'Duties of the Hearts' |
| 2:4.316 | ha-Levavot |  |  |  |
| 2:4.317 | in the gate of |  |  |  |
| 2:4.318 | Unity |  |  |  |
| 2:4.319 | and Rabbi Saadia Gaon |  |  |  |
| 2:4.320 | in the book |  |  |  |
| 2:4.321 | Emunot |  |  | אמונות ודעות: Emunot ve-De'ot, literally 'Beliefs and Opinions' |
| 2:4.322 | and De'ot |  |  |  |
| 2:4.323 | and the Rokeach |  |  |  |
| 2:4.324 | who wrote |  |  |  |
| 2:4.325 | that His oneness |  |  |  |
| 2:4.326 | may He be blessed |  |  | ית׳ = יתברך, 'may He be blessed' |
| 2:4.327 | is not |  |  |  |
| 2:4.328 | like the oneness of |  |  |  |
| 2:4.329 | the pair |  |  |  |
| 2:4.330 | nor |  |  |  |
| 2:4.331 | like the oneness of |  |  |  |
| 2:4.332 | the kind |  |  |  |
| 2:4.333 | nor |  |  |  |
| 2:4.334 | like the man |  |  |  |
| 2:4.335 | the one |  |  |  |
| 2:4.336 | who is divided |  |  |  |
| 2:4.337 | into ones |  |  |  |
| 2:4.338 | many |  |  |  |
| 2:4.339 | and not |  |  |  |
| 2:4.340 | one |  |  |  |
| 2:4.341 | like |  |  |  |
| 2:4.342 | the body |  |  |  |
| 2:4.343 | simple |  |  |  |
| 2:4.344 | etc. |  |  |  |
| 2:4.345 | which receives |  |  |  |
| 2:4.346 | the division |  |  |  |
| 2:4.347 | to Ein |  |  |  |
| 2:4.348 | Sof |  |  |  |
| 2:4.349 | but |  |  |  |
| 2:4.350 | the Name, may He be blessed |  |  | השי״ת: 'the Name, may He be blessed' |
| 2:4.351 | He |  |  |  |
| 2:4.352 | one |  |  |  |
| 2:4.353 | with an oneness |  |  |  |
| 2:4.354 | that has no |  |  |  |
| 2:4.355 | like it |  |  |  |
| 2:4.356 | also |  |  |  |
| 2:4.357 | wrote |  |  |  |
| 2:4.358 | Kisse |  |  |  |
| 2:4.359 | Eliyahu |  |  |  |
| 2:4.360 | that the Ein |  |  |  |
| 2:4.361 | Sof |  |  |  |
| 2:4.362 | is |  |  |  |
| 2:4.363 | the soul of |  |  |  |
| 2:4.364 | Atik |  |  | עתיק: Atik, a partzuf |
| 2:4.365 | and Arikh |  |  | ארוך: Arikh, 'long'; Arikh Anpin, 'long face' |
| 2:4.366 | Anpin |  |  |  |
| 2:4.367 | and Abba |  |  |  |
| 2:4.368 | and Imma |  |  |  |
| 2:4.369 | and Ze'ir |  |  |  |
| 2:4.370 | and his Nukva |  |  |  |
| 2:4.371 | and my teachers |  |  |  |
| 2:4.372 | of blessed memory |  |  |  |
| 2:4.373 | the aforesaid |  |  |  |
| 2:4.374 | wrote |  |  |  |
| 2:4.375 | that He is not |  |  |  |
| 2:4.376 | a body |  |  |  |
| 2:4.377 | nor |  |  |  |
| 2:4.378 | a force |  |  |  |
| 2:4.379 | in a body |  |  |  |
| 2:4.380 | and according to |  |  |  |
| 2:4.381 | his words |  |  |  |
| 2:4.382 | He |  |  |  |
| 2:4.383 | a force |  |  |  |
| 2:4.384 | in a body |  |  |  |
| 2:4.385 | also |  |  |  |
| 2:4.386 | he wrote |  |  |  |
| 2:4.387 | that Arikh |  |  |  |
| 2:4.388 | Anpin |  |  |  |
| 2:4.389 | Abba |  |  |  |
| 2:4.390 | and Imma |  |  |  |
| 2:4.391 | preceded |  |  |  |
| 2:4.392 | Ze'ir |  |  |  |
| 2:4.393 | Anpin |  |  |  |
| 2:4.394 | who |  |  |  |
| 2:4.395 | is our God |  |  |  |
| 2:4.396 | according to |  |  |  |
| 2:4.397 | his words |  |  |  |
| 2:4.398 | and he |  |  |  |
| 2:4.399 | is called |  |  |  |
| 2:4.400 | son |  |  |  |
| 2:4.401 | of Abba |  |  |  |
| 2:4.402 | and Imma |  |  |  |
| 2:4.403 | etc. |  |  |  |
| 2:4.404 | and our Rabbis of blessed memory |  |  |  |
| 2:4.405 | said |  |  |  |
| 2:4.406 | that He |  |  |  |
| 2:4.407 | eternal |  |  | קדמון: 'eternal'; the decisions give the note 'without beginning' |
| 2:4.408 | before all |  |  |  |
| 2:4.409 | the existents |  |  |  |
| 2:4.410 | and that every |  |  |  |
| 2:4.411 | existent |  |  |  |
| 2:4.412 | other than Him |  |  |  |
| 2:4.413 | not |  |  |  |
| 2:4.414 | eternal |  |  |  |
| 2:4.415 | etc. |  |  |  |
| 2:4.416 | also |  |  |  |
| 2:4.417 | he wrote |  |  |  |
| 2:4.418 | concerning Ein |  |  |  |
| 2:4.419 | Sof |  |  |  |
| 2:4.420 | blessed be He |  |  | ב״ה: 'blessed be He' |
| 2:4.421 | not |  |  |  |
| 2:4.422 | pertains |  |  |  |
| 2:4.423 | any |  |  |  |
| 2:4.424 | worship |  |  |  |
| 2:4.425 | to Him |  |  |  |
| 2:4.426 | and not |  |  |  |
| 2:4.427 | any |  |  |  |
| 2:4.428 | calling |  |  |  |
| 2:4.429 | at all |  |  |  |
| 2:4.430 | and not |  |  |  |
| 2:4.431 | applies |  |  |  |
| 2:4.432 | to Him |  |  |  |
| 2:4.433 | any |  |  |  |
| 2:4.434 | name |  |  |  |
| 2:4.435 | not |  |  |  |
| 2:4.436 | YHVH |  |  | הוי״ה: the Name YHVH, as a name |
| 2:4.437 | nor |  |  |  |
| 2:4.438 | Adonai |  |  | אדנות: literally 'lordship' |
| 2:4.439 | etc. |  |  |  |
| 2:4.440 | but only |  |  |  |
| 2:4.441 | to Ze'ir Anpin |  |  | לז״א: to Ze'ir Anpin |
| 2:4.442 | and his Nukva |  |  |  |
| 2:4.443 | in the world |  |  |  |
| 2:4.444 | of Emanation |  |  |  |
| 2:4.445 | alone |  |  |  |
| 2:4.446 | and not |  |  |  |
| 2:4.447 | to Arikh Anpin |  |  | לא״א: to Arikh Anpin (א״א) |
| 2:4.448 | and Abba |  |  |  |
| 2:4.449 | and Imma |  |  |  |
| 2:4.450 | and Ze'ir |  |  |  |
| 2:4.451 | and his Nukva |  |  |  |
| 2:4.452 | that are |  |  |  |
| 2:4.453 | in the worlds |  |  |  |
| 2:4.454 | the emanated |  |  |  |
| 2:4.455 | above |  |  |  |
| 2:4.456 | the world of |  |  |  |
| 2:4.457 | Emanation |  |  |  |
| 2:4.458 | as explained |  |  |  |
| 2:4.459 | there |  |  |  |
| 2:4.460 | 4 |  |  | ד׳: 4, the folio |
| 2:4.461 | 49 |  |  | מ״ט: 49 |
| 2:4.462 | end of |  |  |  |
| 2:4.463 | side |  |  |  |
| 2:4.464 | a |  |  |  |
| 2:4.465 | and on folio |  |  |  |
| 2:4.466 | 53 |  |  | נ״ג: 53 |

Phrases:
- 2:4.1 2:4.2 2:4.3 | The questioner replied again. | again the-questioner replied
- 2:4.4 2:4.5 2:4.6 2:4.7 2:4.8 2:4.9 2:4.10 2:4.11 2:4.12 2:4.13 | Your reply, my friend, came closed and sealed, and in it there is nothing but an indication of a place: | your-reply my-friend came closed and-sealed and-there-is-not in-it but an-indication of-a-place
- 2:4.14 2:4.15 2:4.16 2:4.17 2:4.18 2:4.19 2:4.20 2:4.21 | "to study in Kisse Eliyahu, from folio 13 to folio 18, | to-study in-Kisse Eliyahu from-folio 13 to folio 18
- 2:4.22 2:4.23 2:4.24 2:4.25 | 'A wise man will hear and add to learning,' | will-hear a-wise-man and-will-add to-learning
- 2:4.26 2:4.27 2:4.28 2:4.29 2:4.30 2:4.31 | and what we have received is that worship is to God, one Lord, | and-what-is-received by-us that-worship to-God lord one
- 2:4.32 | etc." | etc.
- 2:4.33 2:4.34 2:4.35 2:4.36 2:4.37 | I did not understand your opinion and your intention in this, | and-not I-understood your-opinion and-your-intention in-this
- 2:4.38 2:4.39 2:4.40 2:4.41 2:4.42 2:4.43 2:4.44 2:4.45 2:4.46 2:4.47 2:4.48 | whether this, what you have received, is in accordance with the words of Kisse Eliyahu or the opposite of his words. | whether this what-is-received by-you is in-accordance-with-the-words-of Kisse Eliyahu or the-opposite-of his-words
- 2:4.49 2:4.50 2:4.51 2:4.52 2:4.53 2:4.54 2:4.55 2:4.56 | For when I studied in Kisse Eliyahu the place you marked for me, | for when I-studied in-Kisse Eliyahu the-place you-marked for-me
- 2:4.57 2:4.58 2:4.59 2:4.60 | I did not find what I sought; on the contrary, | not I-found what-I-sought on-the-contrary
- 2:4.61 2:4.62 2:4.63 2:4.64 | there, on folio 3, he compares | there on-folio 3 he-compares
- 2:4.65 2:4.66 2:4.67 2:4.68 2:4.69 2:4.70 2:4.71 | the oneness of the Creator to the other kinds of oneness, where he wrote: | the-oneness of-the-Creator to-the-other kinds of-oneness where-there he-wrote
- 2:4.72 2:4.73 2:4.74 2:4.75 2:4.76 2:4.77 2:4.78 2:4.79 | "Just as this house as a whole is called by the name one house. | just-as this the-house as-a-whole is-called by-the-name house one
- 2:4.80 2:4.81 2:4.82 2:4.83 2:4.84 2:4.85 2:4.86 2:4.87 2:4.88 2:4.89 2:4.90 2:4.91 | And if you enter it, you will find many rooms, large and small, and many other places. | and-if you-enter in-it you-will-find some and-some rooms large and-small and-some places other-than-them
- 2:4.92 2:4.93 2:4.94 2:4.95 2:4.96 2:4.97 2:4.98 2:4.99 | And each and every one of them ascribes to it one name. | and-each one and-one they-ascribe to-it a-name-attribute name one
- 2:4.100 2:4.101 2:4.102 2:4.103 2:4.104 2:4.105 2:4.106 2:4.107 2:4.108 2:4.109 | Also, a wooden building that has several attributes within it has them. | also the-building of wood that-has within-it several attributes has in-it
- 2:4.110 2:4.111 2:4.112 2:4.113 2:4.114 2:4.115 2:4.116 2:4.117 2:4.118 | Also, you will find further that this holy wall, which was built, was in parts; | also you-will-find further that-this the-wall holy which-was-built was in-parts
- 2:4.119 2:4.120 2:4.121 2:4.122 2:4.123 2:4.124 2:4.125 2:4.126 2:4.127 2:4.128 2:4.129 | each and every stone was separate, of clods of earth by itself and of lime by itself, etc. | each stone and-stone it was-divided of-clods the-earth by-itself and-of-lime by-itself etc.
- 2:4.130 2:4.131 2:4.132 2:4.133 2:4.134 2:4.135 2:4.136 2:4.137 | And after the builder came and joined all these parts with his wisdom, | and-after came the-builder and-joined all parts these with-his-wisdom
- 2:4.138 2:4.139 2:4.140 2:4.141 2:4.142 2:4.143 2:4.144 2:4.145 2:4.146 | this within that, and this beside that, and they became one wall, | this within that and-this beside that and-became a-wall one
- 2:4.147 | etc., | etc.
- 2:4.148 2:4.149 2:4.150 2:4.151 2:4.152 2:4.153 2:4.154 | Thus, then, all is one piece and one part." | it-is-found then all a-piece one and-a-part one
- 2:4.155 2:4.156 2:4.157 2:4.158 | And so is the matter with us, etc. | and-so the-matter with-us etc.
- 2:4.159 2:4.160 2:4.161 2:4.162 2:4.163 2:4.164 2:4.165 2:4.166 | And there, on folios 25 and 29, he compares it to the human body, | and-there on-folio 25 and-29 he-compares it to-the-body of-man
- 2:4.167 2:4.168 2:4.169 2:4.170 2:4.171 2:4.172 2:4.173 2:4.174 2:4.175 2:4.176 2:4.177 | joined of bones and sinews and flesh, head and eye and ear and nose and mouth and hands and feet, | joined of-bones and-sinews and-flesh and-head and-eye and-ear and-nose and-mouth and-hands and-feet
- 2:4.178 2:4.179 2:4.180 2:4.181 2:4.182 2:4.183 2:4.184 | and the whole is called by the name man, Reuben or Simeon. | and-all is-called by-the-name man Reuben or Simeon
- 2:4.185 2:4.186 2:4.187 2:4.188 2:4.189 | And on folios 25, 26 and 27 he wrote: | and-on-folio 25 and-26 and-27 he-wrote
- 2:4.190 2:4.191 2:4.192 2:4.193 2:4.194 2:4.195 2:4.196 2:4.197 2:4.198 2:4.199 2:4.200 | "The sum of the matter is that the First Cause, which all the Kabbalists call Ein Sof, | the-sum that-emerges that the-cause the-first which-is-called by-the-mouth-of all the-Kabbalists without end
- 2:4.201 2:4.202 2:4.203 2:4.204 2:4.205 2:4.206 | is He who emanated, created, formed and made, | he who emanated and-created and-formed and-made
- 2:4.207 2:4.208 2:4.209 2:4.210 2:4.211 | and He is hidden within Ze'ir Anpin, | and-he is-hidden within Ze'ir Anpin
- 2:4.212 2:4.213 2:4.214 2:4.215 2:4.216 | and then Ze'ir Anpin is the one who rules | and-then Ze'ir Anpin he who-rules
- 2:4.217 2:4.218 2:4.219 2:4.220 2:4.221 2:4.222 2:4.223 2:4.224 2:4.225 2:4.226 | over all creatures, leads them, feeds them and sustains them, by the power of Ein Sof within him. | over all creatures and-leads-them and-feeds-them and-sustains-them by-the-power Ein Sof within-him
- 2:4.227 2:4.228 2:4.229 2:4.230 2:4.231 | Therefore He is our God, and we are His people, | therefore He our-God and-we his-people
- 2:4.232 2:4.233 2:4.234 2:4.235 2:4.236 2:4.237 | for our souls are His portion, and Him we shall worship, | for our-souls they His-portion and-Him we-shall-worship
- 2:4.238 2:4.239 2:4.240 2:4.241 2:4.242 2:4.243 | and He is the God of our fathers, in whose hand is the conduct of the worlds, | and-he God-of our-fathers in-whose-hand the-conduct the-worlds
- 2:4.244 2:4.245 | for reward and punishment. | for-reward and-punishment
- 2:4.246 2:4.247 2:4.248 2:4.249 2:4.250 2:4.251 2:4.252 2:4.253 2:4.254 | But Ein Sof, with the other partzufim above Ze'ir and his Nukva, | but the-Ein Sof with the-rest the-partzufim that-are-above Ze'ir and-his-Nukva
- 2:4.255 2:4.256 2:4.257 2:4.258 2:4.259 2:4.260 2:4.261 2:4.262 | if human beings order their prayers to them individually, without | if they-arrange human beings their-prayer to-them individually without
- 2:4.263 2:4.264 2:4.265 | praying to Ze'ir Anpin, | praying to-Ze'ir Anpin
- 2:4.266 2:4.267 2:4.268 2:4.269 2:4.270 2:4.271 2:4.272 2:4.273 2:4.274 2:4.275 | even if they intend the soul hidden in them, they will cry out and not be answered. | even if they-intend to-the-soul hidden in-them they-will-be crying-out and-not answered
- 2:4.276 2:4.277 2:4.278 2:4.279 2:4.280 | And on the contrary, those who pray to them will be punished, | and-on-the-contrary will-be those-who-pray to-them punished
- 2:4.281 2:4.282 2:4.283 2:4.284 2:4.285 2:4.286 2:4.287 2:4.288 2:4.289 2:4.290 2:4.291 | for it is the will of the First Cause that He be the one who brings out and brings in His flow to the lower beings, | for this-is the-will of-the-cause the-first that he the-one-who-brings-out and-brings-in His-flow to-the-lower-beings
- 2:4.292 2:4.293 2:4.294 | and there is no other besides Him. | and-there-is-not other besides-Him
- 2:4.295 2:4.296 | See there carefully." | see-there carefully
- 2:4.297 2:4.298 2:4.299 2:4.300 2:4.301 2:4.302 | And the words of Rabbi Kisse Eliyahu stand against | and-the-words-of Rabbi Kisse Eliyahu stand against
- 2:4.303 2:4.304 2:4.305 2:4.306 2:4.307 2:4.308 2:4.309 2:4.310 | the words of the Rambam, of blessed memory, in his commentary on the Mishnah, in his composition, and in Moreh Nevukhim, | the-words-of the-Rambam of-blessed-memory in-his-commentary on-the-Mishnah and-in-his-composition and-in-Moreh Nevukhim
- 2:4.311 2:4.312 2:4.313 2:4.314 2:4.315 2:4.316 2:4.317 2:4.318 | and against the words of the pious one, the author of Chovot ha-Levavot, in the Gate of Unity, | and-against the-words-of the-pious-one the-author-of Chovot ha-Levavot in-the-gate-of Unity
- 2:4.319 2:4.320 2:4.321 2:4.322 2:4.323 | and Rabbi Saadia Gaon in the book Emunot ve-De'ot, and the Rokeach, | and-Rabbi-Saadia-Gaon in-the-book Emunot and-De'ot and-the-Rokeach
- 2:4.324 2:4.325 2:4.326 2:4.327 2:4.328 2:4.329 | who wrote that His oneness, may He be blessed, is not like the oneness of a pair, | who-wrote that-His-oneness may-He-be-blessed is-not like-the-oneness-of the-pair
- 2:4.330 2:4.331 2:4.332 2:4.333 2:4.334 2:4.335 2:4.336 2:4.337 2:4.338 | nor like the oneness of a kind, nor like the one man who is divided into many ones. | nor like-the-oneness-of the-kind nor like-the-man the-one who-is-divided into-ones many
- 2:4.339 2:4.340 2:4.341 2:4.342 2:4.343 2:4.344 2:4.345 2:4.346 2:4.347 2:4.348 | and not one like a simple body, etc., which receives division to Ein Sof. | and-not one like the-body simple etc. which-receives the-division to-Ein Sof
- 2:4.349 2:4.350 2:4.351 2:4.352 2:4.353 2:4.354 2:4.355 | But the Name, may He be blessed, is one with a oneness that has no like. | but the-Name,-may-He-be-blessed He one with-an-oneness that-has-no like-it
- 2:4.356 2:4.357 2:4.358 2:4.359 2:4.360 2:4.361 | Also, Kisse Eliyahu wrote that Ein Sof | also wrote Kisse Eliyahu that-the-Ein Sof
- 2:4.362 2:4.363 2:4.364 2:4.365 2:4.366 2:4.367 2:4.368 | is the soul of Atik, and Arikh Anpin, and Abba and Imma, | is the-soul-of Atik and-Arikh Anpin and-Abba and-Imma
- 2:4.369 2:4.370 | and Ze'ir and his Nukva, | and-Ze'ir and-his-Nukva
- 2:4.371 2:4.372 2:4.373 2:4.374 | and my teachers, of blessed memory, the aforesaid, wrote | and-my-teachers of-blessed-memory the-aforesaid wrote
- 2:4.375 2:4.376 2:4.377 2:4.378 2:4.379 | that He is not a body, nor a force in a body! | that-He-is-not a-body nor a-force in-a-body
- 2:4.380 2:4.381 2:4.382 2:4.383 2:4.384 | And according to his words He is a force in a body. | and-according-to his-words He a-force in-a-body
- 2:4.385 2:4.386 2:4.387 2:4.388 2:4.389 2:4.390 2:4.391 2:4.392 2:4.393 | Also he wrote that Arikh Anpin, Abba and Imma preceded Ze'ir Anpin | also he-wrote that-Arikh Anpin Abba and-Imma preceded Ze'ir Anpin
- 2:4.394 2:4.395 2:4.396 2:4.397 | who is our God (according to his words), | who is-our-God according-to his-words
- 2:4.398 2:4.399 2:4.400 2:4.401 2:4.402 2:4.403 | and he is called a son of Abba and Imma, etc. | and-he is-called son of-Abba and-Imma etc.
- 2:4.404 2:4.405 2:4.406 | And our Rabbis of blessed memory said that He | and-our-Rabbis-of-blessed-memory said that-He
- 2:4.407 2:4.408 2:4.409 2:4.410 2:4.411 2:4.412 2:4.413 2:4.414 2:4.415 | is eternal before all that exists, and that every existent other than Him is not eternal, etc. | eternal before-all the-existents and-that-every existent other-than-Him not eternal etc.
- 2:4.416 2:4.417 2:4.418 2:4.419 2:4.420 | Also he wrote concerning Ein Sof, blessed be He, | also he-wrote concerning-Ein Sof blessed-be-He
- 2:4.421 2:4.422 2:4.423 2:4.424 2:4.425 2:4.426 2:4.427 2:4.428 2:4.429 | that no worship pertains to Him, and no invocation at all. | not pertains any worship to-Him and-not any calling at-all
- 2:4.430 2:4.431 2:4.432 2:4.433 2:4.434 2:4.435 2:4.436 | And no name applies to Him: not YHVH, | and-not applies to-Him any name not YHVH
- 2:4.437 2:4.438 2:4.439 | nor Adonai, etc. | nor Adonai etc.
- 2:4.440 2:4.441 2:4.442 2:4.443 2:4.444 2:4.445 | but only to Ze'ir Anpin and his Nukva in the World of Emanation alone, | but-only to-Ze'ir-Anpin and-his-Nukva in-the-world of-Emanation alone
- 2:4.446 2:4.447 2:4.448 2:4.449 2:4.450 2:4.451 | and not to Arikh Anpin, Abba and Imma, and Ze'ir and his Nukva | and-not to-Arikh-Anpin and-Abba and-Imma and-Ze'ir and-his-Nukva
- 2:4.452 2:4.453 2:4.454 2:4.455 2:4.456 2:4.457 2:4.458 | that are in the emanated worlds above the World of Emanation, as explained | that-are in-the-worlds the-emanated above the-world-of Emanation as-explained
- 2:4.459 2:4.460 2:4.461 2:4.462 2:4.463 2:4.464 2:4.465 2:4.466 | there, 4, 49, the end of side a, and on folio 53.... | there 4 49 end-of side a and-on-folio 53

Grammar notes:
- The paragraph is the questioner's reply. The respondent's words and Kisse Eliyahu are quoted in double marks; a verse inside them is in single marks.
- ולא הבנתי דעתך וכוונתך בזה, אם זה ... או היפך דבריו: an indirect question, 'whether ... or ...', with אם introducing the first alternative.
- אדרבא שם: 'on the contrary' reverses what was expected; the book says the opposite of what the questioner sought.
- Indirect reports introduced by כתב ש- or כתבו ש- (for example 'Also he wrote that ...') carry no quotation marks; only the direct formula כתב: opens a quotation.
- ואין עוד מלבדו is an embedded verse (Deut 4:35, 4:39), translated from its Hebrew as 'there is no other besides Him'.

Translator's note: The questioner's second letter (חזר השואל והשיב). The Wikisource text ends mid-sentence ('....') and the English keeps the break.

Commentary rulings:
- apparatus: Prov 1:5 (2:4.22–25): 'ישמע חכם ויוסף לקח', rendered as the verse's future.
- apparatus: Deut 4:35 and 4:39 (2:4.292–294): 'ואין עוד מלבדו'.
- apparatus: שם (2:4.61) is read as 'there', the place in the book; the other sense, 'the name', does not fit the sentence.
- apparatus: The author writes ארוך (2:4.365) and שאריך (2:4.387) for Arikh; both are rendered Arikh Anpin.
- apparatus: ד׳ (2:4.460) is read as 4, the folio; the numeral is kept because the author writes דף in full nearby.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

