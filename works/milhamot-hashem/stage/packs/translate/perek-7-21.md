# Translation pack: milhamot-hashem perek-7-21 (Milchamot ha-Shem 7:26–7:26)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-21`.

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
- רשב"י → "Rabbi Shimon bar Yochai", always written out. רשב"י התנא →
  "Rabbi Shimon bar Yochai the Tanna". The author separates the Tanna from
  the writer who "named himself" after him (1:4: "אשר כינה את עצמו בשם רשב״י
  התנא"). Keep exactly the qualifier he gives. Do not add "pseudo-" or
  "the so-called", and do not settle in the English which of the two he
  means.
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
- **שיתוף / שתוף / לשתף** (there is no שותפות in the text): context.
  - Associating another with God: fixed, "associate / association". "כל
    המשתף שם שמים ודבר אחר נעקר מן העולם" (Sanh. 63a; 1:7, 5:29, 5:32) →
    "whoever associates the name of Heaven with something else is uprooted
    from the world"; "בשיתוף שאר הפרצופים" (5:1) → "in association with the
    other partzufim"; "בלי חיבור וצירוף ושיתוף" (5:21) → "without
    composition, conjunction or association".
  - The Maimonidean term for equivocal words: "בשתוף השם בלבד" (5:10) → "by
    equivocation of the term only"; "מושאלים או משותפים" (5:18) →
    "borrowed or equivocal". Never "association" here. The note gives
    Guide I.
  - צירוף: "conjunction" when used of God; "permutation" in letter mysticism
    ("י"ב צירופי הוי'ה", 2:33 → "the twelve permutations of YHVH").
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
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
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


## Milchamot ha-Shem 7:26

Hebrew (Wikisource): צח) והנה ידוע ומפורסם הוא לכל בקי בספרות ישראל. שחבור הזהר ושאר ספרי קבלה הנכריה החדשה לא נתפשטו בישראל כי מתחלת האלף השישי, ולא נודע בבירור מהיכן צמחו, כמו שאביא להלן דעות המאמתים אותם. ולא ידעו לברר הדבר, אך הקרוב אל הדעת הוא מה שכתב הרב חיד"א בס' שם הגדולים מערכת ספרים אות זי"ן וכבר הבאתי דבריו לעיל. וחכמי ירושלם באגדותיהם מסכימים לדבריו יע"ש. כי מיד המלך הנכרי זכו לאמונה. ובגליון הבא מירושלם ת"ו האריכו מאד להעתיק בקצור את דברי האחרונים הטחים תפל על האמונות הנכריות לרפאם ולא נרפאו וכל מגמת חכמי ירושת"ו בכל מה שהעתיקו בגליון מדברי האחרונים. הוא לשלול הגשמות מהספירות והצורות והפרצופים הנזכרים בזהר ושאר ספרי קבלה! ואם הצילונו מן הנשמות כלום הצילונו מעבודת אל נר בין גשמי בין רוחני? כי מודים הם שהם סבות נבראות, וכל אחד סבה ועלה למה שלמטה ממנה זכרים ונקבות, כמבואר בזהר וספרים הנמשכים אחריו, רק שהם כוחות רוחניות של אור נברא ואותם יעבודו, ולהם יקראו בעת צרתם ולהם ישבחו ויודו ויהללו בעת ישועתם. ומקבלים עליהם לאלוהות כאלו הוא מותר לנו מדין תורתינו הקדושה לעבוד כח נברא, רק שלא נחשוב אותם גשמים אלא רוחנים. כי נסלק מהם ונשלול הנשמות, ונחשוב בדמיונינו שהם אורות. (ולא ידעו שגם האור והאויר הם גופים דקים סי' מ"ז) ומביאים ראיות צנומות דקות שדופות קדים לשלול הגשמות מהם. ולחשוב את המשפיע זכר ואת מקבל השפע נקבה, ואת הזווגים שלא כפשטן. רק קבלת השפע קרו ליה זווג! ועם כל זה לא הביאו לנו פתרונים לשמות איברים מגונים כיסוד הוא גיד הערוה, וביעין דדכורא, ויסוד דנוקבא ונקב האוחרים. למאי חזו באור הזה המופשט מן הגשמות? ואם התורה הקדושה והנביאים נשתמשו בלשונות המורות על הגשמות כיד ואצבעות. כלום השתמשו ליחס לשי"ת חושים הפחותים באכילה ושתייה ומשוש אף בחלק המעולה והחשוב שבהם. כ"ש בחלק המגונה והפחות שבחוש המשוש שהוא המשגל. כמו שכתב הרב המורה בפרק מ"ז מראשון. והיותר רע ומר הוא מה שישבו סבות רבות באלהות כולן ברואות. והמה אצלם עיקר האלהות. ויחברו ספרים בסדר מערכותן כעורך ומסדר את החלות על השלחן. ויקראו לכל אחד מן הפרצופים בשם הוי"ה ואדנות ואלהים צבאות, ויקבלו עליהן אלהותן כמו שכתב ספר הברית וכסא אליהו ונחלת יוסף לעבוד דוקא לאלוה קצר אפים (זעיר אנפין) ולא לאלוה העליון הבורא את כל הפרצופים והצורות. וכתב עוד ספר הברית שהעובד לאלוה העליון על כולם לא יאבא קצר האפים הנקרא אצלם בשם הוי"ה והקב"ה לסלוח לו, כי אז יעשן אפו וקנאתו, באיש ההוא וכו'. כי רק לו לבדו ניתנה הממשלה מהאלהים העליונים ממנו שהסכימו להמליכו, למרות דבר ה' אשר דבר על יד נביאו וכבודי לאחר לא אתן. ודבריהם אלו מבוארים בזהר בלק דף קצ"א עם פירוש מקדש מלך. ובפרט בספר עז לאלהים שהיא משל על זה ממלך ב"ו שמלך שגם אביו ואימו צריכים לכבדו, וזה המשל מתנגד למשל שהביאו רז"ל בשמות רבה על אנכי ה' אלהיך. ובספר שלשלת הקבלה להראב"ד כתב משם ס' יוחסין וז"ל בשנת כמו חמשת אלפים וחמשים שנה נמצאו כתות אנשים שאמרו כי דברי הזהר אשר הם בלשון ירושלמי הם דברי רשב"י. אבל אשר הם בלשון הקדש אינן דבריו, ויש אומרים שהרמב"ן מצאו בארץ ישראל והביא אותו לקטלוניאה, ואחר לארגון. ונפל ליד רבי משה דיליואן. ויש אומרים כי ר' משה דיליואן הנז' היה חכם גדול ועושה אלו הפירושים משכלו, ולמען יקח מהם מחיר רב מהחכמים היה כותבם ותולה אותם ברשב"י וחבורתו, ע"כ בספר יוחסין ומאריך עוד יותר עכ"ל שלשלת הקבלה. וכבר הבאתי לעיל סי' צ"ד דברי הרב חיד"א, בס' שם הגדולים וגליון שהובא מירושלם ת"י על ידי מלשינות הקשר אשר הרה וילד ארך הלשון. והנה שכחו דין פשוט בשועיורה דעה סי' רפ"א שספר תורה הנמצא ביד גוי דלד' הרי"ף והרא"ש והטור יגנז ואין קורין בו אע"פ שאין בו שום חסרון או יתרון משאר ספרי תורה, ולדעת הרמב"ם ז"ל כיון שאין ידוע מי כתבו ואין בו שום שינוי חזקתו שישראל כתבו וקורין בו, הא אם נמצא בו שום שינוי אפילו בחסר ויתר דבר פשוט הוא דלכ"ע אין קורין בו ובנדון דידן של ספר הזהר יש לדון בו קל וחומר, ומה אם זה שאין בו שום קלקול לאמונתינו יגנז? ספר המבזה את קבלתינו המשנה והתלמוד ומסית ומדיח אותנו לעבוד כחות נבראים. קצר אפים עם הוריו ויולדיו, על אחת כמה וכמה שאין להאמין לו להמיר את אלהינו ברוך הוא וב"ש ולסתור פנות תורתינו הקדושה. ומדבריהם מודאע רבה לאוריתא שהתנאי שאמר הקמב"ן שא"א להבין דברי הקבלה מתוך הספר רק מפי חכם מקובל לאזן מקבל, לא נתקיים מעולם תנאי זה, שהרי הם בעצמן מודים שאחרי מות רשב"י ובנו ר"א וכל הדור ההוא נגנז הזהר ונשכחה ואבדה הקבלה הלזו ולא נשאר אדם יודע בה עד שנתגלה ספר הזהר באלף השישי ע"י המלך העכו"ם. וממנו נתפשטה הקבלה ולא מפי חכם מקובל אשר קבל מפה לאזן. כי כל המקובלים הרבים החדשים שבעולם לא שאבו רק מהזהר וממנו נתפשטה, ולפי דברי הרמב"ן זהו הגורם שכולם שגגו וטעו ועקרו יסודי תורתינו הקדושה והאמינו באלהים רבים. ובאו לידי שתוף שם שמים ודבר אחד בעבור מיעוט הבנתם. אף לפי האמת כיון שגילה דעתו בביאור שהוא מבזה את המשנה והתלמוד ומדמה שקלא וטריא דרבנן המתאבקים בתלמוד לכלבין דצוחין הב, הב, אין מקום לכל זה (עיין תקוני זהר חדש). ולפי דברי הרב שם הגדולים והגליון המובא מירושת"ו שר' נחוניא בן הקנה היה ראש המקובלים אשר גילו את הקבלה החדשה ולמדו אותה בגלוי! עד שתיקן בכניסתו לבית המדרש להתפלל שלא אכשל בדבר הלכה, שלא אומר על טמא טהור וכו' כדתנן בברכות פ' תפלת השחר. היה לו להתפלל שלא תבוא תקלה גדולה כזו על ידו להאמין ברשויות רבות ולעבוד אל אחר קצר אפים (זעיר אנפין)? לפי שעונש המחטיא אחרים להאמין ברשויות רבות ולעבוד אלהי נכר יותר גדול מן הנכשל להתיר את האסור או לאסור את המותר?

| id | word | punctuation after |
|---|---|---|
| 7:26.1 | צח | ) |
| 7:26.2 | והנה |  |
| 7:26.3 | ידוע |  |
| 7:26.4 | ומפורסם |  |
| 7:26.5 | הוא |  |
| 7:26.6 | לכל |  |
| 7:26.7 | בקי |  |
| 7:26.8 | בספרות |  |
| 7:26.9 | ישראל | . |
| 7:26.10 | שחבור |  |
| 7:26.11 | הזהר |  |
| 7:26.12 | ושאר |  |
| 7:26.13 | ספרי |  |
| 7:26.14 | קבלה |  |
| 7:26.15 | הנכריה |  |
| 7:26.16 | החדשה |  |
| 7:26.17 | לא |  |
| 7:26.18 | נתפשטו |  |
| 7:26.19 | בישראל |  |
| 7:26.20 | כי |  |
| 7:26.21 | מתחלת |  |
| 7:26.22 | האלף |  |
| 7:26.23 | השישי | , |
| 7:26.24 | ולא |  |
| 7:26.25 | נודע |  |
| 7:26.26 | בבירור |  |
| 7:26.27 | מהיכן |  |
| 7:26.28 | צמחו | , |
| 7:26.29 | כמו |  |
| 7:26.30 | שאביא |  |
| 7:26.31 | להלן |  |
| 7:26.32 | דעות |  |
| 7:26.33 | המאמתים |  |
| 7:26.34 | אותם | . |
| 7:26.35 | ולא |  |
| 7:26.36 | ידעו |  |
| 7:26.37 | לברר |  |
| 7:26.38 | הדבר | , |
| 7:26.39 | אך |  |
| 7:26.40 | הקרוב |  |
| 7:26.41 | אל |  |
| 7:26.42 | הדעת |  |
| 7:26.43 | הוא |  |
| 7:26.44 | מה |  |
| 7:26.45 | שכתב |  |
| 7:26.46 | הרב |  |
| 7:26.47 | חיד"א |  |
| 7:26.48 | בס | ' |
| 7:26.49 | שם |  |
| 7:26.50 | הגדולים |  |
| 7:26.51 | מערכת |  |
| 7:26.52 | ספרים |  |
| 7:26.53 | אות |  |
| 7:26.54 | זי"ן |  |
| 7:26.55 | וכבר |  |
| 7:26.56 | הבאתי |  |
| 7:26.57 | דבריו |  |
| 7:26.58 | לעיל | . |
| 7:26.59 | וחכמי |  |
| 7:26.60 | ירושלם |  |
| 7:26.61 | באגדותיהם |  |
| 7:26.62 | מסכימים |  |
| 7:26.63 | לדבריו |  |
| 7:26.64 | יע"ש | . |
| 7:26.65 | כי |  |
| 7:26.66 | מיד |  |
| 7:26.67 | המלך |  |
| 7:26.68 | הנכרי |  |
| 7:26.69 | זכו |  |
| 7:26.70 | לאמונה | . |
| 7:26.71 | ובגליון |  |
| 7:26.72 | הבא |  |
| 7:26.73 | מירושלם |  |
| 7:26.74 | ת"ו |  |
| 7:26.75 | האריכו |  |
| 7:26.76 | מאד |  |
| 7:26.77 | להעתיק |  |
| 7:26.78 | בקצור |  |
| 7:26.79 | את |  |
| 7:26.80 | דברי |  |
| 7:26.81 | האחרונים |  |
| 7:26.82 | הטחים |  |
| 7:26.83 | תפל |  |
| 7:26.84 | על |  |
| 7:26.85 | האמונות |  |
| 7:26.86 | הנכריות |  |
| 7:26.87 | לרפאם |  |
| 7:26.88 | ולא |  |
| 7:26.89 | נרפאו |  |
| 7:26.90 | וכל |  |
| 7:26.91 | מגמת |  |
| 7:26.92 | חכמי |  |
| 7:26.93 | ירושת"ו |  |
| 7:26.94 | בכל |  |
| 7:26.95 | מה |  |
| 7:26.96 | שהעתיקו |  |
| 7:26.97 | בגליון |  |
| 7:26.98 | מדברי |  |
| 7:26.99 | האחרונים | . |
| 7:26.100 | הוא |  |
| 7:26.101 | לשלול |  |
| 7:26.102 | הגשמות |  |
| 7:26.103 | מהספירות |  |
| 7:26.104 | והצורות |  |
| 7:26.105 | והפרצופים |  |
| 7:26.106 | הנזכרים |  |
| 7:26.107 | בזהר |  |
| 7:26.108 | ושאר |  |
| 7:26.109 | ספרי |  |
| 7:26.110 | קבלה | ! |
| 7:26.111 | ואם |  |
| 7:26.112 | הצילונו |  |
| 7:26.113 | מן |  |
| 7:26.114 | הנשמות |  |
| 7:26.115 | כלום |  |
| 7:26.116 | הצילונו |  |
| 7:26.117 | מעבודת |  |
| 7:26.118 | אל |  |
| 7:26.119 | נר |  |
| 7:26.120 | בין |  |
| 7:26.121 | גשמי |  |
| 7:26.122 | בין |  |
| 7:26.123 | רוחני | ? |
| 7:26.124 | כי |  |
| 7:26.125 | מודים |  |
| 7:26.126 | הם |  |
| 7:26.127 | שהם |  |
| 7:26.128 | סבות |  |
| 7:26.129 | נבראות | , |
| 7:26.130 | וכל |  |
| 7:26.131 | אחד |  |
| 7:26.132 | סבה |  |
| 7:26.133 | ועלה |  |
| 7:26.134 | למה |  |
| 7:26.135 | שלמטה |  |
| 7:26.136 | ממנה |  |
| 7:26.137 | זכרים |  |
| 7:26.138 | ונקבות | , |
| 7:26.139 | כמבואר |  |
| 7:26.140 | בזהר |  |
| 7:26.141 | וספרים |  |
| 7:26.142 | הנמשכים |  |
| 7:26.143 | אחריו | , |
| 7:26.144 | רק |  |
| 7:26.145 | שהם |  |
| 7:26.146 | כוחות |  |
| 7:26.147 | רוחניות |  |
| 7:26.148 | של |  |
| 7:26.149 | אור |  |
| 7:26.150 | נברא |  |
| 7:26.151 | ואותם |  |
| 7:26.152 | יעבודו | , |
| 7:26.153 | ולהם |  |
| 7:26.154 | יקראו |  |
| 7:26.155 | בעת |  |
| 7:26.156 | צרתם |  |
| 7:26.157 | ולהם |  |
| 7:26.158 | ישבחו |  |
| 7:26.159 | ויודו |  |
| 7:26.160 | ויהללו |  |
| 7:26.161 | בעת |  |
| 7:26.162 | ישועתם | . |
| 7:26.163 | ומקבלים |  |
| 7:26.164 | עליהם |  |
| 7:26.165 | לאלוהות |  |
| 7:26.166 | כאלו |  |
| 7:26.167 | הוא |  |
| 7:26.168 | מותר |  |
| 7:26.169 | לנו |  |
| 7:26.170 | מדין |  |
| 7:26.171 | תורתינו |  |
| 7:26.172 | הקדושה |  |
| 7:26.173 | לעבוד |  |
| 7:26.174 | כח |  |
| 7:26.175 | נברא | , |
| 7:26.176 | רק |  |
| 7:26.177 | שלא |  |
| 7:26.178 | נחשוב |  |
| 7:26.179 | אותם |  |
| 7:26.180 | גשמים |  |
| 7:26.181 | אלא |  |
| 7:26.182 | רוחנים | . |
| 7:26.183 | כי |  |
| 7:26.184 | נסלק |  |
| 7:26.185 | מהם |  |
| 7:26.186 | ונשלול |  |
| 7:26.187 | הנשמות | , |
| 7:26.188 | ונחשוב |  |
| 7:26.189 | בדמיונינו |  |
| 7:26.190 | שהם |  |
| 7:26.191 | אורות | . |
| 7:26.192 | ולא |  |
| 7:26.193 | ידעו |  |
| 7:26.194 | שגם |  |
| 7:26.195 | האור |  |
| 7:26.196 | והאויר |  |
| 7:26.197 | הם |  |
| 7:26.198 | גופים |  |
| 7:26.199 | דקים |  |
| 7:26.200 | סי | ' |
| 7:26.201 | מ"ז | ) |
| 7:26.202 | ומביאים |  |
| 7:26.203 | ראיות |  |
| 7:26.204 | צנומות |  |
| 7:26.205 | דקות |  |
| 7:26.206 | שדופות |  |
| 7:26.207 | קדים |  |
| 7:26.208 | לשלול |  |
| 7:26.209 | הגשמות |  |
| 7:26.210 | מהם | . |
| 7:26.211 | ולחשוב |  |
| 7:26.212 | את |  |
| 7:26.213 | המשפיע |  |
| 7:26.214 | זכר |  |
| 7:26.215 | ואת |  |
| 7:26.216 | מקבל |  |
| 7:26.217 | השפע |  |
| 7:26.218 | נקבה | , |
| 7:26.219 | ואת |  |
| 7:26.220 | הזווגים |  |
| 7:26.221 | שלא |  |
| 7:26.222 | כפשטן | . |
| 7:26.223 | רק |  |
| 7:26.224 | קבלת |  |
| 7:26.225 | השפע |  |
| 7:26.226 | קרו |  |
| 7:26.227 | ליה |  |
| 7:26.228 | זווג | ! |
| 7:26.229 | ועם |  |
| 7:26.230 | כל |  |
| 7:26.231 | זה |  |
| 7:26.232 | לא |  |
| 7:26.233 | הביאו |  |
| 7:26.234 | לנו |  |
| 7:26.235 | פתרונים |  |
| 7:26.236 | לשמות |  |
| 7:26.237 | איברים |  |
| 7:26.238 | מגונים |  |
| 7:26.239 | כיסוד |  |
| 7:26.240 | הוא |  |
| 7:26.241 | גיד |  |
| 7:26.242 | הערוה | , |
| 7:26.243 | וביעין |  |
| 7:26.244 | דדכורא | , |
| 7:26.245 | ויסוד |  |
| 7:26.246 | דנוקבא |  |
| 7:26.247 | ונקב |  |
| 7:26.248 | האוחרים | . |
| 7:26.249 | למאי |  |
| 7:26.250 | חזו |  |
| 7:26.251 | באור |  |
| 7:26.252 | הזה |  |
| 7:26.253 | המופשט |  |
| 7:26.254 | מן |  |
| 7:26.255 | הגשמות | ? |
| 7:26.256 | ואם |  |
| 7:26.257 | התורה |  |
| 7:26.258 | הקדושה |  |
| 7:26.259 | והנביאים |  |
| 7:26.260 | נשתמשו |  |
| 7:26.261 | בלשונות |  |
| 7:26.262 | המורות |  |
| 7:26.263 | על |  |
| 7:26.264 | הגשמות |  |
| 7:26.265 | כיד |  |
| 7:26.266 | ואצבעות | . |
| 7:26.267 | כלום |  |
| 7:26.268 | השתמשו |  |
| 7:26.269 | ליחס |  |
| 7:26.270 | לשי"ת |  |
| 7:26.271 | חושים |  |
| 7:26.272 | הפחותים |  |
| 7:26.273 | באכילה |  |
| 7:26.274 | ושתייה |  |
| 7:26.275 | ומשוש |  |
| 7:26.276 | אף |  |
| 7:26.277 | בחלק |  |
| 7:26.278 | המעולה |  |
| 7:26.279 | והחשוב |  |
| 7:26.280 | שבהם | . |
| 7:26.281 | כ"ש |  |
| 7:26.282 | בחלק |  |
| 7:26.283 | המגונה |  |
| 7:26.284 | והפחות |  |
| 7:26.285 | שבחוש |  |
| 7:26.286 | המשוש |  |
| 7:26.287 | שהוא |  |
| 7:26.288 | המשגל | . |
| 7:26.289 | כמו |  |
| 7:26.290 | שכתב |  |
| 7:26.291 | הרב |  |
| 7:26.292 | המורה |  |
| 7:26.293 | בפרק |  |
| 7:26.294 | מ"ז |  |
| 7:26.295 | מראשון | . |
| 7:26.296 | והיותר |  |
| 7:26.297 | רע |  |
| 7:26.298 | ומר |  |
| 7:26.299 | הוא |  |
| 7:26.300 | מה |  |
| 7:26.301 | שישבו |  |
| 7:26.302 | סבות |  |
| 7:26.303 | רבות |  |
| 7:26.304 | באלהות |  |
| 7:26.305 | כולן |  |
| 7:26.306 | ברואות | . |
| 7:26.307 | והמה |  |
| 7:26.308 | אצלם |  |
| 7:26.309 | עיקר |  |
| 7:26.310 | האלהות | . |
| 7:26.311 | ויחברו |  |
| 7:26.312 | ספרים |  |
| 7:26.313 | בסדר |  |
| 7:26.314 | מערכותן |  |
| 7:26.315 | כעורך |  |
| 7:26.316 | ומסדר |  |
| 7:26.317 | את |  |
| 7:26.318 | החלות |  |
| 7:26.319 | על |  |
| 7:26.320 | השלחן | . |
| 7:26.321 | ויקראו |  |
| 7:26.322 | לכל |  |
| 7:26.323 | אחד |  |
| 7:26.324 | מן |  |
| 7:26.325 | הפרצופים |  |
| 7:26.326 | בשם |  |
| 7:26.327 | הוי"ה |  |
| 7:26.328 | ואדנות |  |
| 7:26.329 | ואלהים |  |
| 7:26.330 | צבאות | , |
| 7:26.331 | ויקבלו |  |
| 7:26.332 | עליהן |  |
| 7:26.333 | אלהותן |  |
| 7:26.334 | כמו |  |
| 7:26.335 | שכתב |  |
| 7:26.336 | ספר |  |
| 7:26.337 | הברית |  |
| 7:26.338 | וכסא |  |
| 7:26.339 | אליהו |  |
| 7:26.340 | ונחלת |  |
| 7:26.341 | יוסף |  |
| 7:26.342 | לעבוד |  |
| 7:26.343 | דוקא |  |
| 7:26.344 | לאלוה |  |
| 7:26.345 | קצר |  |
| 7:26.346 | אפים |  |
| 7:26.347 | זעיר |  |
| 7:26.348 | אנפין | ) |
| 7:26.349 | ולא |  |
| 7:26.350 | לאלוה |  |
| 7:26.351 | העליון |  |
| 7:26.352 | הבורא |  |
| 7:26.353 | את |  |
| 7:26.354 | כל |  |
| 7:26.355 | הפרצופים |  |
| 7:26.356 | והצורות | . |
| 7:26.357 | וכתב |  |
| 7:26.358 | עוד |  |
| 7:26.359 | ספר |  |
| 7:26.360 | הברית |  |
| 7:26.361 | שהעובד |  |
| 7:26.362 | לאלוה |  |
| 7:26.363 | העליון |  |
| 7:26.364 | על |  |
| 7:26.365 | כולם |  |
| 7:26.366 | לא |  |
| 7:26.367 | יאבא |  |
| 7:26.368 | קצר |  |
| 7:26.369 | האפים |  |
| 7:26.370 | הנקרא |  |
| 7:26.371 | אצלם |  |
| 7:26.372 | בשם |  |
| 7:26.373 | הוי"ה |  |
| 7:26.374 | והקב"ה |  |
| 7:26.375 | לסלוח |  |
| 7:26.376 | לו | , |
| 7:26.377 | כי |  |
| 7:26.378 | אז |  |
| 7:26.379 | יעשן |  |
| 7:26.380 | אפו |  |
| 7:26.381 | וקנאתו | , |
| 7:26.382 | באיש |  |
| 7:26.383 | ההוא |  |
| 7:26.384 | וכו | '. |
| 7:26.385 | כי |  |
| 7:26.386 | רק |  |
| 7:26.387 | לו |  |
| 7:26.388 | לבדו |  |
| 7:26.389 | ניתנה |  |
| 7:26.390 | הממשלה |  |
| 7:26.391 | מהאלהים |  |
| 7:26.392 | העליונים |  |
| 7:26.393 | ממנו |  |
| 7:26.394 | שהסכימו |  |
| 7:26.395 | להמליכו | , |
| 7:26.396 | למרות |  |
| 7:26.397 | דבר |  |
| 7:26.398 | ה | ' |
| 7:26.399 | אשר |  |
| 7:26.400 | דבר |  |
| 7:26.401 | על |  |
| 7:26.402 | יד |  |
| 7:26.403 | נביאו |  |
| 7:26.404 | וכבודי |  |
| 7:26.405 | לאחר |  |
| 7:26.406 | לא |  |
| 7:26.407 | אתן | . |
| 7:26.408 | ודבריהם |  |
| 7:26.409 | אלו |  |
| 7:26.410 | מבוארים |  |
| 7:26.411 | בזהר |  |
| 7:26.412 | בלק |  |
| 7:26.413 | דף |  |
| 7:26.414 | קצ"א |  |
| 7:26.415 | עם |  |
| 7:26.416 | פירוש |  |
| 7:26.417 | מקדש |  |
| 7:26.418 | מלך | . |
| 7:26.419 | ובפרט |  |
| 7:26.420 | בספר |  |
| 7:26.421 | עז |  |
| 7:26.422 | לאלהים |  |
| 7:26.423 | שהיא |  |
| 7:26.424 | משל |  |
| 7:26.425 | על |  |
| 7:26.426 | זה |  |
| 7:26.427 | ממלך |  |
| 7:26.428 | ב"ו |  |
| 7:26.429 | שמלך |  |
| 7:26.430 | שגם |  |
| 7:26.431 | אביו |  |
| 7:26.432 | ואימו |  |
| 7:26.433 | צריכים |  |
| 7:26.434 | לכבדו | , |
| 7:26.435 | וזה |  |
| 7:26.436 | המשל |  |
| 7:26.437 | מתנגד |  |
| 7:26.438 | למשל |  |
| 7:26.439 | שהביאו |  |
| 7:26.440 | רז"ל |  |
| 7:26.441 | בשמות |  |
| 7:26.442 | רבה |  |
| 7:26.443 | על |  |
| 7:26.444 | אנכי |  |
| 7:26.445 | ה | ' |
| 7:26.446 | אלהיך | . |
| 7:26.447 | ובספר |  |
| 7:26.448 | שלשלת |  |
| 7:26.449 | הקבלה |  |
| 7:26.450 | להראב"ד |  |
| 7:26.451 | כתב |  |
| 7:26.452 | משם |  |
| 7:26.453 | ס | ' |
| 7:26.454 | יוחסין |  |
| 7:26.455 | וז"ל |  |
| 7:26.456 | בשנת |  |
| 7:26.457 | כמו |  |
| 7:26.458 | חמשת |  |
| 7:26.459 | אלפים |  |
| 7:26.460 | וחמשים |  |
| 7:26.461 | שנה |  |
| 7:26.462 | נמצאו |  |
| 7:26.463 | כתות |  |
| 7:26.464 | אנשים |  |
| 7:26.465 | שאמרו |  |
| 7:26.466 | כי |  |
| 7:26.467 | דברי |  |
| 7:26.468 | הזהר |  |
| 7:26.469 | אשר |  |
| 7:26.470 | הם |  |
| 7:26.471 | בלשון |  |
| 7:26.472 | ירושלמי |  |
| 7:26.473 | הם |  |
| 7:26.474 | דברי |  |
| 7:26.475 | רשב"י | . |
| 7:26.476 | אבל |  |
| 7:26.477 | אשר |  |
| 7:26.478 | הם |  |
| 7:26.479 | בלשון |  |
| 7:26.480 | הקדש |  |
| 7:26.481 | אינן |  |
| 7:26.482 | דבריו | , |
| 7:26.483 | ויש |  |
| 7:26.484 | אומרים |  |
| 7:26.485 | שהרמב"ן |  |
| 7:26.486 | מצאו |  |
| 7:26.487 | בארץ |  |
| 7:26.488 | ישראל |  |
| 7:26.489 | והביא |  |
| 7:26.490 | אותו |  |
| 7:26.491 | לקטלוניאה | , |
| 7:26.492 | ואחר |  |
| 7:26.493 | לארגון | . |
| 7:26.494 | ונפל |  |
| 7:26.495 | ליד |  |
| 7:26.496 | רבי |  |
| 7:26.497 | משה |  |
| 7:26.498 | דיליואן | . |
| 7:26.499 | ויש |  |
| 7:26.500 | אומרים |  |
| 7:26.501 | כי |  |
| 7:26.502 | ר | ' |
| 7:26.503 | משה |  |
| 7:26.504 | דיליואן |  |
| 7:26.505 | הנז | ' |
| 7:26.506 | היה |  |
| 7:26.507 | חכם |  |
| 7:26.508 | גדול |  |
| 7:26.509 | ועושה |  |
| 7:26.510 | אלו |  |
| 7:26.511 | הפירושים |  |
| 7:26.512 | משכלו | , |
| 7:26.513 | ולמען |  |
| 7:26.514 | יקח |  |
| 7:26.515 | מהם |  |
| 7:26.516 | מחיר |  |
| 7:26.517 | רב |  |
| 7:26.518 | מהחכמים |  |
| 7:26.519 | היה |  |
| 7:26.520 | כותבם |  |
| 7:26.521 | ותולה |  |
| 7:26.522 | אותם |  |
| 7:26.523 | ברשב"י |  |
| 7:26.524 | וחבורתו | , |
| 7:26.525 | ע"כ |  |
| 7:26.526 | בספר |  |
| 7:26.527 | יוחסין |  |
| 7:26.528 | ומאריך |  |
| 7:26.529 | עוד |  |
| 7:26.530 | יותר |  |
| 7:26.531 | עכ"ל |  |
| 7:26.532 | שלשלת |  |
| 7:26.533 | הקבלה | . |
| 7:26.534 | וכבר |  |
| 7:26.535 | הבאתי |  |
| 7:26.536 | לעיל |  |
| 7:26.537 | סי | ' |
| 7:26.538 | צ"ד |  |
| 7:26.539 | דברי |  |
| 7:26.540 | הרב |  |
| 7:26.541 | חיד"א | , |
| 7:26.542 | בס | ' |
| 7:26.543 | שם |  |
| 7:26.544 | הגדולים |  |
| 7:26.545 | וגליון |  |
| 7:26.546 | שהובא |  |
| 7:26.547 | מירושלם |  |
| 7:26.548 | ת"י |  |
| 7:26.549 | על |  |
| 7:26.550 | ידי |  |
| 7:26.551 | מלשינות |  |
| 7:26.552 | הקשר |  |
| 7:26.553 | אשר |  |
| 7:26.554 | הרה |  |
| 7:26.555 | וילד |  |
| 7:26.556 | ארך |  |
| 7:26.557 | הלשון | . |
| 7:26.558 | והנה |  |
| 7:26.559 | שכחו |  |
| 7:26.560 | דין |  |
| 7:26.561 | פשוט |  |
| 7:26.562 | בשועיורה |  |
| 7:26.563 | דעה |  |
| 7:26.564 | סי | ' |
| 7:26.565 | רפ"א |  |
| 7:26.566 | שספר |  |
| 7:26.567 | תורה |  |
| 7:26.568 | הנמצא |  |
| 7:26.569 | ביד |  |
| 7:26.570 | גוי |  |
| 7:26.571 | דלד | ' |
| 7:26.572 | הרי"ף |  |
| 7:26.573 | והרא"ש |  |
| 7:26.574 | והטור |  |
| 7:26.575 | יגנז |  |
| 7:26.576 | ואין |  |
| 7:26.577 | קורין |  |
| 7:26.578 | בו |  |
| 7:26.579 | אע"פ |  |
| 7:26.580 | שאין |  |
| 7:26.581 | בו |  |
| 7:26.582 | שום |  |
| 7:26.583 | חסרון |  |
| 7:26.584 | או |  |
| 7:26.585 | יתרון |  |
| 7:26.586 | משאר |  |
| 7:26.587 | ספרי |  |
| 7:26.588 | תורה | , |
| 7:26.589 | ולדעת |  |
| 7:26.590 | הרמב"ם |  |
| 7:26.591 | ז"ל |  |
| 7:26.592 | כיון |  |
| 7:26.593 | שאין |  |
| 7:26.594 | ידוע |  |
| 7:26.595 | מי |  |
| 7:26.596 | כתבו |  |
| 7:26.597 | ואין |  |
| 7:26.598 | בו |  |
| 7:26.599 | שום |  |
| 7:26.600 | שינוי |  |
| 7:26.601 | חזקתו |  |
| 7:26.602 | שישראל |  |
| 7:26.603 | כתבו |  |
| 7:26.604 | וקורין |  |
| 7:26.605 | בו | , |
| 7:26.606 | הא |  |
| 7:26.607 | אם |  |
| 7:26.608 | נמצא |  |
| 7:26.609 | בו |  |
| 7:26.610 | שום |  |
| 7:26.611 | שינוי |  |
| 7:26.612 | אפילו |  |
| 7:26.613 | בחסר |  |
| 7:26.614 | ויתר |  |
| 7:26.615 | דבר |  |
| 7:26.616 | פשוט |  |
| 7:26.617 | הוא |  |
| 7:26.618 | דלכ"ע |  |
| 7:26.619 | אין |  |
| 7:26.620 | קורין |  |
| 7:26.621 | בו |  |
| 7:26.622 | ובנדון |  |
| 7:26.623 | דידן |  |
| 7:26.624 | של |  |
| 7:26.625 | ספר |  |
| 7:26.626 | הזהר |  |
| 7:26.627 | יש |  |
| 7:26.628 | לדון |  |
| 7:26.629 | בו |  |
| 7:26.630 | קל |  |
| 7:26.631 | וחומר | , |
| 7:26.632 | ומה |  |
| 7:26.633 | אם |  |
| 7:26.634 | זה |  |
| 7:26.635 | שאין |  |
| 7:26.636 | בו |  |
| 7:26.637 | שום |  |
| 7:26.638 | קלקול |  |
| 7:26.639 | לאמונתינו |  |
| 7:26.640 | יגנז | ? |
| 7:26.641 | ספר |  |
| 7:26.642 | המבזה |  |
| 7:26.643 | את |  |
| 7:26.644 | קבלתינו |  |
| 7:26.645 | המשנה |  |
| 7:26.646 | והתלמוד |  |
| 7:26.647 | ומסית |  |
| 7:26.648 | ומדיח |  |
| 7:26.649 | אותנו |  |
| 7:26.650 | לעבוד |  |
| 7:26.651 | כחות |  |
| 7:26.652 | נבראים | . |
| 7:26.653 | קצר |  |
| 7:26.654 | אפים |  |
| 7:26.655 | עם |  |
| 7:26.656 | הוריו |  |
| 7:26.657 | ויולדיו | , |
| 7:26.658 | על |  |
| 7:26.659 | אחת |  |
| 7:26.660 | כמה |  |
| 7:26.661 | וכמה |  |
| 7:26.662 | שאין |  |
| 7:26.663 | להאמין |  |
| 7:26.664 | לו |  |
| 7:26.665 | להמיר |  |
| 7:26.666 | את |  |
| 7:26.667 | אלהינו |  |
| 7:26.668 | ברוך |  |
| 7:26.669 | הוא |  |
| 7:26.670 | וב"ש |  |
| 7:26.671 | ולסתור |  |
| 7:26.672 | פנות |  |
| 7:26.673 | תורתינו |  |
| 7:26.674 | הקדושה | . |
| 7:26.675 | ומדבריהם |  |
| 7:26.676 | מודאע |  |
| 7:26.677 | רבה |  |
| 7:26.678 | לאוריתא |  |
| 7:26.679 | שהתנאי |  |
| 7:26.680 | שאמר |  |
| 7:26.681 | הקמב"ן |  |
| 7:26.682 | שא"א |  |
| 7:26.683 | להבין |  |
| 7:26.684 | דברי |  |
| 7:26.685 | הקבלה |  |
| 7:26.686 | מתוך |  |
| 7:26.687 | הספר |  |
| 7:26.688 | רק |  |
| 7:26.689 | מפי |  |
| 7:26.690 | חכם |  |
| 7:26.691 | מקובל |  |
| 7:26.692 | לאזן |  |
| 7:26.693 | מקבל | , |
| 7:26.694 | לא |  |
| 7:26.695 | נתקיים |  |
| 7:26.696 | מעולם |  |
| 7:26.697 | תנאי |  |
| 7:26.698 | זה | , |
| 7:26.699 | שהרי |  |
| 7:26.700 | הם |  |
| 7:26.701 | בעצמן |  |
| 7:26.702 | מודים |  |
| 7:26.703 | שאחרי |  |
| 7:26.704 | מות |  |
| 7:26.705 | רשב"י |  |
| 7:26.706 | ובנו |  |
| 7:26.707 | ר"א |  |
| 7:26.708 | וכל |  |
| 7:26.709 | הדור |  |
| 7:26.710 | ההוא |  |
| 7:26.711 | נגנז |  |
| 7:26.712 | הזהר |  |
| 7:26.713 | ונשכחה |  |
| 7:26.714 | ואבדה |  |
| 7:26.715 | הקבלה |  |
| 7:26.716 | הלזו |  |
| 7:26.717 | ולא |  |
| 7:26.718 | נשאר |  |
| 7:26.719 | אדם |  |
| 7:26.720 | יודע |  |
| 7:26.721 | בה |  |
| 7:26.722 | עד |  |
| 7:26.723 | שנתגלה |  |
| 7:26.724 | ספר |  |
| 7:26.725 | הזהר |  |
| 7:26.726 | באלף |  |
| 7:26.727 | השישי |  |
| 7:26.728 | ע"י |  |
| 7:26.729 | המלך |  |
| 7:26.730 | העכו"ם | . |
| 7:26.731 | וממנו |  |
| 7:26.732 | נתפשטה |  |
| 7:26.733 | הקבלה |  |
| 7:26.734 | ולא |  |
| 7:26.735 | מפי |  |
| 7:26.736 | חכם |  |
| 7:26.737 | מקובל |  |
| 7:26.738 | אשר |  |
| 7:26.739 | קבל |  |
| 7:26.740 | מפה |  |
| 7:26.741 | לאזן | . |
| 7:26.742 | כי |  |
| 7:26.743 | כל |  |
| 7:26.744 | המקובלים |  |
| 7:26.745 | הרבים |  |
| 7:26.746 | החדשים |  |
| 7:26.747 | שבעולם |  |
| 7:26.748 | לא |  |
| 7:26.749 | שאבו |  |
| 7:26.750 | רק |  |
| 7:26.751 | מהזהר |  |
| 7:26.752 | וממנו |  |
| 7:26.753 | נתפשטה | , |
| 7:26.754 | ולפי |  |
| 7:26.755 | דברי |  |
| 7:26.756 | הרמב"ן |  |
| 7:26.757 | זהו |  |
| 7:26.758 | הגורם |  |
| 7:26.759 | שכולם |  |
| 7:26.760 | שגגו |  |
| 7:26.761 | וטעו |  |
| 7:26.762 | ועקרו |  |
| 7:26.763 | יסודי |  |
| 7:26.764 | תורתינו |  |
| 7:26.765 | הקדושה |  |
| 7:26.766 | והאמינו |  |
| 7:26.767 | באלהים |  |
| 7:26.768 | רבים | . |
| 7:26.769 | ובאו |  |
| 7:26.770 | לידי |  |
| 7:26.771 | שתוף |  |
| 7:26.772 | שם |  |
| 7:26.773 | שמים |  |
| 7:26.774 | ודבר |  |
| 7:26.775 | אחד |  |
| 7:26.776 | בעבור |  |
| 7:26.777 | מיעוט |  |
| 7:26.778 | הבנתם | . |
| 7:26.779 | אף |  |
| 7:26.780 | לפי |  |
| 7:26.781 | האמת |  |
| 7:26.782 | כיון |  |
| 7:26.783 | שגילה |  |
| 7:26.784 | דעתו |  |
| 7:26.785 | בביאור |  |
| 7:26.786 | שהוא |  |
| 7:26.787 | מבזה |  |
| 7:26.788 | את |  |
| 7:26.789 | המשנה |  |
| 7:26.790 | והתלמוד |  |
| 7:26.791 | ומדמה |  |
| 7:26.792 | שקלא |  |
| 7:26.793 | וטריא |  |
| 7:26.794 | דרבנן |  |
| 7:26.795 | המתאבקים |  |
| 7:26.796 | בתלמוד |  |
| 7:26.797 | לכלבין |  |
| 7:26.798 | דצוחין |  |
| 7:26.799 | הב | , |
| 7:26.800 | הב | , |
| 7:26.801 | אין |  |
| 7:26.802 | מקום |  |
| 7:26.803 | לכל |  |
| 7:26.804 | זה |  |
| 7:26.805 | עיין |  |
| 7:26.806 | תקוני |  |
| 7:26.807 | זהר |  |
| 7:26.808 | חדש | ). |
| 7:26.809 | ולפי |  |
| 7:26.810 | דברי |  |
| 7:26.811 | הרב |  |
| 7:26.812 | שם |  |
| 7:26.813 | הגדולים |  |
| 7:26.814 | והגליון |  |
| 7:26.815 | המובא |  |
| 7:26.816 | מירושת"ו |  |
| 7:26.817 | שר | ' |
| 7:26.818 | נחוניא |  |
| 7:26.819 | בן |  |
| 7:26.820 | הקנה |  |
| 7:26.821 | היה |  |
| 7:26.822 | ראש |  |
| 7:26.823 | המקובלים |  |
| 7:26.824 | אשר |  |
| 7:26.825 | גילו |  |
| 7:26.826 | את |  |
| 7:26.827 | הקבלה |  |
| 7:26.828 | החדשה |  |
| 7:26.829 | ולמדו |  |
| 7:26.830 | אותה |  |
| 7:26.831 | בגלוי | ! |
| 7:26.832 | עד |  |
| 7:26.833 | שתיקן |  |
| 7:26.834 | בכניסתו |  |
| 7:26.835 | לבית |  |
| 7:26.836 | המדרש |  |
| 7:26.837 | להתפלל |  |
| 7:26.838 | שלא |  |
| 7:26.839 | אכשל |  |
| 7:26.840 | בדבר |  |
| 7:26.841 | הלכה | , |
| 7:26.842 | שלא |  |
| 7:26.843 | אומר |  |
| 7:26.844 | על |  |
| 7:26.845 | טמא |  |
| 7:26.846 | טהור |  |
| 7:26.847 | וכו | ' |
| 7:26.848 | כדתנן |  |
| 7:26.849 | בברכות |  |
| 7:26.850 | פ | ' |
| 7:26.851 | תפלת |  |
| 7:26.852 | השחר | . |
| 7:26.853 | היה |  |
| 7:26.854 | לו |  |
| 7:26.855 | להתפלל |  |
| 7:26.856 | שלא |  |
| 7:26.857 | תבוא |  |
| 7:26.858 | תקלה |  |
| 7:26.859 | גדולה |  |
| 7:26.860 | כזו |  |
| 7:26.861 | על |  |
| 7:26.862 | ידו |  |
| 7:26.863 | להאמין |  |
| 7:26.864 | ברשויות |  |
| 7:26.865 | רבות |  |
| 7:26.866 | ולעבוד |  |
| 7:26.867 | אל |  |
| 7:26.868 | אחר |  |
| 7:26.869 | קצר |  |
| 7:26.870 | אפים |  |
| 7:26.871 | זעיר |  |
| 7:26.872 | אנפין | )? |
| 7:26.873 | לפי |  |
| 7:26.874 | שעונש |  |
| 7:26.875 | המחטיא |  |
| 7:26.876 | אחרים |  |
| 7:26.877 | להאמין |  |
| 7:26.878 | ברשויות |  |
| 7:26.879 | רבות |  |
| 7:26.880 | ולעבוד |  |
| 7:26.881 | אלהי |  |
| 7:26.882 | נכר |  |
| 7:26.883 | יותר |  |
| 7:26.884 | גדול |  |
| 7:26.885 | מן |  |
| 7:26.886 | הנכשל |  |
| 7:26.887 | להתיר |  |
| 7:26.888 | את |  |
| 7:26.889 | האסור |  |
| 7:26.890 | או |  |
| 7:26.891 | לאסור |  |
| 7:26.892 | את |  |
| 7:26.893 | המותר | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

