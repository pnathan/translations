# Translation pack: milhamot-hashem perek-6-17 (Milchamot ha-Shem 6:26–6:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-17`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- The words of a verse quoted in the plural (Job 32:16, עמדו): plural, no
  subject added.
- ראב"ע: "Rabbi Eleazar ben Azariah" where a Talmudic passage is cited;
  "Ibn Ezra" where the author cites the commentator; a note says which.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.
- הרא"ש: "the Rosh"; ר"ת: "Rabbeinu Tam"; הגאון alone: "the Gaon", unidentified.
  (Acronym authorities keep the acronym, by the ruling above.)

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


## Milchamot ha-Shem 6:26

Hebrew (Wikisource): אוי לאותה בושה ואוי לאותה כלמה, כי ידע שור קונהו וחמור אבוס בעליו, וישראל לא ידע את אלהיו, ואל איזה צורה ופרצוף מהפרצופים שהמציא מחבר הזהר יתפלל, ויטש אלוה עשהו הראשון וצור לכל, ויבקש אלהים חדשים מקרוב באו אשר אף נגוים עובדי אלילים לא ידעום. ולפי האמור פה חלילה חלילה לחשוב על שום תנא או אמורא כרז"ל אשר ידענו רוב חכמתם וחסידותם וקדושתם שיאמרו דברים כאלה על המשנה והתלמוד אשר הם בעצמם בלו ימיהם בעסק הקדוש הזה, הוא שקלא וטריא דתלמודא, על אבן שתיה, היא המשנה אשר עליה הושת התלמוד הקדוש אשר עיני כל מורי הוראות עליו, ומפיו נבעו כל נהרי נחלי פסקי הלכות, מימות רבותינו הגאונים ועד עתה, וכל משיבי תשובות ממעין המשנה והתלמוד ישאבון.

| id | word | punctuation after |
|---|---|---|
| 6:26.1 | אוי |  |
| 6:26.2 | לאותה |  |
| 6:26.3 | בושה |  |
| 6:26.4 | ואוי |  |
| 6:26.5 | לאותה |  |
| 6:26.6 | כלמה | , |
| 6:26.7 | כי |  |
| 6:26.8 | ידע |  |
| 6:26.9 | שור |  |
| 6:26.10 | קונהו |  |
| 6:26.11 | וחמור |  |
| 6:26.12 | אבוס |  |
| 6:26.13 | בעליו | , |
| 6:26.14 | וישראל |  |
| 6:26.15 | לא |  |
| 6:26.16 | ידע |  |
| 6:26.17 | את |  |
| 6:26.18 | אלהיו | , |
| 6:26.19 | ואל |  |
| 6:26.20 | איזה |  |
| 6:26.21 | צורה |  |
| 6:26.22 | ופרצוף |  |
| 6:26.23 | מהפרצופים |  |
| 6:26.24 | שהמציא |  |
| 6:26.25 | מחבר |  |
| 6:26.26 | הזהר |  |
| 6:26.27 | יתפלל | , |
| 6:26.28 | ויטש |  |
| 6:26.29 | אלוה |  |
| 6:26.30 | עשהו |  |
| 6:26.31 | הראשון |  |
| 6:26.32 | וצור |  |
| 6:26.33 | לכל | , |
| 6:26.34 | ויבקש |  |
| 6:26.35 | אלהים |  |
| 6:26.36 | חדשים |  |
| 6:26.37 | מקרוב |  |
| 6:26.38 | באו |  |
| 6:26.39 | אשר |  |
| 6:26.40 | אף |  |
| 6:26.41 | נגוים |  |
| 6:26.42 | עובדי |  |
| 6:26.43 | אלילים |  |
| 6:26.44 | לא |  |
| 6:26.45 | ידעום | . |
| 6:26.46 | ולפי |  |
| 6:26.47 | האמור |  |
| 6:26.48 | פה |  |
| 6:26.49 | חלילה |  |
| 6:26.50 | חלילה |  |
| 6:26.51 | לחשוב |  |
| 6:26.52 | על |  |
| 6:26.53 | שום |  |
| 6:26.54 | תנא |  |
| 6:26.55 | או |  |
| 6:26.56 | אמורא |  |
| 6:26.57 | כרז"ל |  |
| 6:26.58 | אשר |  |
| 6:26.59 | ידענו |  |
| 6:26.60 | רוב |  |
| 6:26.61 | חכמתם |  |
| 6:26.62 | וחסידותם |  |
| 6:26.63 | וקדושתם |  |
| 6:26.64 | שיאמרו |  |
| 6:26.65 | דברים |  |
| 6:26.66 | כאלה |  |
| 6:26.67 | על |  |
| 6:26.68 | המשנה |  |
| 6:26.69 | והתלמוד |  |
| 6:26.70 | אשר |  |
| 6:26.71 | הם |  |
| 6:26.72 | בעצמם |  |
| 6:26.73 | בלו |  |
| 6:26.74 | ימיהם |  |
| 6:26.75 | בעסק |  |
| 6:26.76 | הקדוש |  |
| 6:26.77 | הזה | , |
| 6:26.78 | הוא |  |
| 6:26.79 | שקלא |  |
| 6:26.80 | וטריא |  |
| 6:26.81 | דתלמודא | , |
| 6:26.82 | על |  |
| 6:26.83 | אבן |  |
| 6:26.84 | שתיה | , |
| 6:26.85 | היא |  |
| 6:26.86 | המשנה |  |
| 6:26.87 | אשר |  |
| 6:26.88 | עליה |  |
| 6:26.89 | הושת |  |
| 6:26.90 | התלמוד |  |
| 6:26.91 | הקדוש |  |
| 6:26.92 | אשר |  |
| 6:26.93 | עיני |  |
| 6:26.94 | כל |  |
| 6:26.95 | מורי |  |
| 6:26.96 | הוראות |  |
| 6:26.97 | עליו | , |
| 6:26.98 | ומפיו |  |
| 6:26.99 | נבעו |  |
| 6:26.100 | כל |  |
| 6:26.101 | נהרי |  |
| 6:26.102 | נחלי |  |
| 6:26.103 | פסקי |  |
| 6:26.104 | הלכות | , |
| 6:26.105 | מימות |  |
| 6:26.106 | רבותינו |  |
| 6:26.107 | הגאונים |  |
| 6:26.108 | ועד |  |
| 6:26.109 | עתה | , |
| 6:26.110 | וכל |  |
| 6:26.111 | משיבי |  |
| 6:26.112 | תשובות |  |
| 6:26.113 | ממעין |  |
| 6:26.114 | המשנה |  |
| 6:26.115 | והתלמוד |  |
| 6:26.116 | ישאבון | . |


## Milchamot ha-Shem 6:27

Hebrew (Wikisource): עב) וקל וחומר הדברים לחשוב על התנא הקדוש רשב"י ע"ה שיכנה את המשנה בשם צלמא דמחתיה אבנא, ובשם אבן משכית, וקליפה, וכיוצא שהרי בירושלמי מבואר בהדיא דרשב"י סבורא ליה שקריאת המשנה חשובה אצלו יותר מן המקרא. ושקול הוא כקריאת שמע בעונתה כמו שנת' לעיל סי' ס"ו. ואיככה נחשוב על התנא הקדוש הזה שידבר סרה על המשנה לומר שכל טמא ואסור ופסול הנז' במשנה מסטרא דיצר הרע. וניחא ליה בהפקירא לזלזל באיסורין שבה, ושעל זה נענש משה רבי' ע"ה להיקבר בארץ טמאה על שלמדנו את המשנה. ושהיה לו למשה ללמדינו תחתיה סודות ורמזים הנוגעים ופוגמים בכבוד אבינו מלכינו וביסודי תורתינו הקדושה ושרשיה להטותינו מעיקרי אמונתינו הטהורה ביחוד השי"ת שהוא הסבה הראשונה. וליסד לנו דת ותורה חדשה להאמין בסבות רבות המשתלשלים זה מזה. וכל אחת נקראת בשם הוי'ה ואדנות ואלהים והקב"ה. והמה שאמרו והיה העולם לפי דבריו, וייחס העבודה והתפלה וקרבנות שהיו ישראל מקריבים בזמן שבית המקדש קיים לאלוה הנברא קצר האפים (ז"א) והוא הסבה האחרונה שבפרצופי האצילות היפך תורתינו הקדושה שבכתב ושבעל פה. ואע"פ שאמר שנגלה לו עתיקא קדישא כמבואר בזהר בראשית דף כ"ב או אליהו ורעיא מהימנא כמבואר בזהר בכמה מקומות. כבר הוזהרנו בתורה שלא נאמין לו ולא נשמע אליו, כמו שנ' כי יקום בקרבך נביא או חלם חלום וכו' כי נביא שקר הוא, ואף אם יעמיד לנו חמה ולבנה כיהושע בגבעון ועמק אילון לא נאמין לו כמו שאמרו רז"ל שגמרא ומביאו הר' הסמ"ג בהקדמת העשין והרמב"ם בחבור ובהקדמת סדר זרעים ובמורה פרק כ"ד משלישי באורך. והטעם כי מנסה ה' אלהינו אותנו כמו שנאמר כי מנסה ה' אלהיכם אתכם לדעת הישכם אוהבים את ה' אלהיכם וגו' ואין נביא רשאי לחדש דבר מעתה, ואף אם בת קול תכריז מן השמים לאמר שמעו בקולו, לא נשמע לו לסור מן הדרך אשר צוה ה' אלהינו אותנו ולזלזל בשום עיקר מעקרי התורה או בשום מצוה ממצותיה, כל שכן זה המסית בלי שום אות ומופות ועוקר ד' עיקרים מיסודי תורתינו הקדושה שכתב הרמב"ם ז"ל כמו שנת' לעיל סי' מ"ח, והנה כפי הנחת מחבר הזהר שכללות הפרצופים הם אחד, איך נפרד העתיק מאריך אנפין ושאר פרצופים ובא לבדו לבית מדרשו של רשב"י ושאל אותו ואמר לו שמעון שמעון מאן הוא דאמר ויאמר אלהים נעשה אדם מאן הוא האי אלהים וכו' (זוהר בראשית כ״ב:ח׳) וכן בכמה מקומות בזהר שנזכר סבא דסבין שהיה בא לגבי רבי אלעזר איך נפרד ובא לבדו הלא הכל אחד? וכבר קדמוני בטענה זו הראב"ע ז"ל על מאמיני השלוש בריש פרשת וירא וז"ל הנה קצת אמרו כי השם שלשה אנשים הוא אחד והוא שלשה ולא יתפרדו, והנה שכחו ויבאו שני המלאכים סדמה ע"כ. ביאור דבריו כי מאמיני השלוש סוברים שהשם ית' הוא שלשה, אב ובן ורוח הקדש והשלשה הם אחד ולכך אמר הכתוב וירא אליו ה', והנה שלשה אנשים, ואלה השלשה הם השם הנז' בפסוק שקודם זה ולא יתפרדו, וטען עליהם הראב"ע האי מצינו שנפרדו שנ' ויבאו שני המלאכים סדומה? גם אנחנו נאמר להמסית מחבר הזהר שאומר וכלא יד ולא מתפרשין לעלמין. איך נפרד עתיק שהוא סבא דסבין ובא לבדו לבית מדרשו של רשב"י? גם מה שאמר שאימא היא שאמרה לאבא נעשה אדם, וגברה ידה ועשאתו על כרחו של אבא, וכשחטא אדם גירשו מגן עדן וגירשה עמו, ואיך נפרדה אימא מאבא, והלא כלא חד ולא מתפרשין? והנה אם היה סבא דסבין נראה לרשב"י ולר' אלעזר בנו כמה פעמים כמו שמבואר בזהר, למה בכה רשב"י ואמר שפחה של בית אבא נזדמן לה מלאך שלשה פעמים. אני אפי' לא פעם אחת? (מעילה דף פ"ד). וגם למה לא הגידו לנו חכמים שרשב"י היה נביא גדול כ"כ יותר משאר נביאים, שהרי היה נגלה לו עתיקא קדישא והוא עומד על עמדו בתוך החבירים בבית המדרש? ולמה אמרו חכמים שחגי זכריה ומלאכי סוף נביאים, ומשם ואילך פסקה נבואה ישראל, הא אשכחן רשב"י ניבא אחר החרבן, והגיד כמה עניינים באלהות, ושהעבודה לבן האלהים קצר האפים מה שלא הודיעו שאר נביאים מימות משה עד מלאכי, ולא נתגלה זה עד האלף הששי?

| id | word | punctuation after |
|---|---|---|
| 6:27.1 | עב | ) |
| 6:27.2 | וקל |  |
| 6:27.3 | וחומר |  |
| 6:27.4 | הדברים |  |
| 6:27.5 | לחשוב |  |
| 6:27.6 | על |  |
| 6:27.7 | התנא |  |
| 6:27.8 | הקדוש |  |
| 6:27.9 | רשב"י |  |
| 6:27.10 | ע"ה |  |
| 6:27.11 | שיכנה |  |
| 6:27.12 | את |  |
| 6:27.13 | המשנה |  |
| 6:27.14 | בשם |  |
| 6:27.15 | צלמא |  |
| 6:27.16 | דמחתיה |  |
| 6:27.17 | אבנא | , |
| 6:27.18 | ובשם |  |
| 6:27.19 | אבן |  |
| 6:27.20 | משכית | , |
| 6:27.21 | וקליפה | , |
| 6:27.22 | וכיוצא |  |
| 6:27.23 | שהרי |  |
| 6:27.24 | בירושלמי |  |
| 6:27.25 | מבואר |  |
| 6:27.26 | בהדיא |  |
| 6:27.27 | דרשב"י |  |
| 6:27.28 | סבורא |  |
| 6:27.29 | ליה |  |
| 6:27.30 | שקריאת |  |
| 6:27.31 | המשנה |  |
| 6:27.32 | חשובה |  |
| 6:27.33 | אצלו |  |
| 6:27.34 | יותר |  |
| 6:27.35 | מן |  |
| 6:27.36 | המקרא | . |
| 6:27.37 | ושקול |  |
| 6:27.38 | הוא |  |
| 6:27.39 | כקריאת |  |
| 6:27.40 | שמע |  |
| 6:27.41 | בעונתה |  |
| 6:27.42 | כמו |  |
| 6:27.43 | שנת | ' |
| 6:27.44 | לעיל |  |
| 6:27.45 | סי | ' |
| 6:27.46 | ס"ו | . |
| 6:27.47 | ואיככה |  |
| 6:27.48 | נחשוב |  |
| 6:27.49 | על |  |
| 6:27.50 | התנא |  |
| 6:27.51 | הקדוש |  |
| 6:27.52 | הזה |  |
| 6:27.53 | שידבר |  |
| 6:27.54 | סרה |  |
| 6:27.55 | על |  |
| 6:27.56 | המשנה |  |
| 6:27.57 | לומר |  |
| 6:27.58 | שכל |  |
| 6:27.59 | טמא |  |
| 6:27.60 | ואסור |  |
| 6:27.61 | ופסול |  |
| 6:27.62 | הנז | ' |
| 6:27.63 | במשנה |  |
| 6:27.64 | מסטרא |  |
| 6:27.65 | דיצר |  |
| 6:27.66 | הרע | . |
| 6:27.67 | וניחא |  |
| 6:27.68 | ליה |  |
| 6:27.69 | בהפקירא |  |
| 6:27.70 | לזלזל |  |
| 6:27.71 | באיסורין |  |
| 6:27.72 | שבה | , |
| 6:27.73 | ושעל |  |
| 6:27.74 | זה |  |
| 6:27.75 | נענש |  |
| 6:27.76 | משה |  |
| 6:27.77 | רבי | ' |
| 6:27.78 | ע"ה |  |
| 6:27.79 | להיקבר |  |
| 6:27.80 | בארץ |  |
| 6:27.81 | טמאה |  |
| 6:27.82 | על |  |
| 6:27.83 | שלמדנו |  |
| 6:27.84 | את |  |
| 6:27.85 | המשנה | . |
| 6:27.86 | ושהיה |  |
| 6:27.87 | לו |  |
| 6:27.88 | למשה |  |
| 6:27.89 | ללמדינו |  |
| 6:27.90 | תחתיה |  |
| 6:27.91 | סודות |  |
| 6:27.92 | ורמזים |  |
| 6:27.93 | הנוגעים |  |
| 6:27.94 | ופוגמים |  |
| 6:27.95 | בכבוד |  |
| 6:27.96 | אבינו |  |
| 6:27.97 | מלכינו |  |
| 6:27.98 | וביסודי |  |
| 6:27.99 | תורתינו |  |
| 6:27.100 | הקדושה |  |
| 6:27.101 | ושרשיה |  |
| 6:27.102 | להטותינו |  |
| 6:27.103 | מעיקרי |  |
| 6:27.104 | אמונתינו |  |
| 6:27.105 | הטהורה |  |
| 6:27.106 | ביחוד |  |
| 6:27.107 | השי"ת |  |
| 6:27.108 | שהוא |  |
| 6:27.109 | הסבה |  |
| 6:27.110 | הראשונה | . |
| 6:27.111 | וליסד |  |
| 6:27.112 | לנו |  |
| 6:27.113 | דת |  |
| 6:27.114 | ותורה |  |
| 6:27.115 | חדשה |  |
| 6:27.116 | להאמין |  |
| 6:27.117 | בסבות |  |
| 6:27.118 | רבות |  |
| 6:27.119 | המשתלשלים |  |
| 6:27.120 | זה |  |
| 6:27.121 | מזה | . |
| 6:27.122 | וכל |  |
| 6:27.123 | אחת |  |
| 6:27.124 | נקראת |  |
| 6:27.125 | בשם |  |
| 6:27.126 | הוי'ה |  |
| 6:27.127 | ואדנות |  |
| 6:27.128 | ואלהים |  |
| 6:27.129 | והקב"ה | . |
| 6:27.130 | והמה |  |
| 6:27.131 | שאמרו |  |
| 6:27.132 | והיה |  |
| 6:27.133 | העולם |  |
| 6:27.134 | לפי |  |
| 6:27.135 | דבריו | , |
| 6:27.136 | וייחס |  |
| 6:27.137 | העבודה |  |
| 6:27.138 | והתפלה |  |
| 6:27.139 | וקרבנות |  |
| 6:27.140 | שהיו |  |
| 6:27.141 | ישראל |  |
| 6:27.142 | מקריבים |  |
| 6:27.143 | בזמן |  |
| 6:27.144 | שבית |  |
| 6:27.145 | המקדש |  |
| 6:27.146 | קיים |  |
| 6:27.147 | לאלוה |  |
| 6:27.148 | הנברא |  |
| 6:27.149 | קצר |  |
| 6:27.150 | האפים |  |
| 6:27.151 | ז"א | ) |
| 6:27.152 | והוא |  |
| 6:27.153 | הסבה |  |
| 6:27.154 | האחרונה |  |
| 6:27.155 | שבפרצופי |  |
| 6:27.156 | האצילות |  |
| 6:27.157 | היפך |  |
| 6:27.158 | תורתינו |  |
| 6:27.159 | הקדושה |  |
| 6:27.160 | שבכתב |  |
| 6:27.161 | ושבעל |  |
| 6:27.162 | פה | . |
| 6:27.163 | ואע"פ |  |
| 6:27.164 | שאמר |  |
| 6:27.165 | שנגלה |  |
| 6:27.166 | לו |  |
| 6:27.167 | עתיקא |  |
| 6:27.168 | קדישא |  |
| 6:27.169 | כמבואר |  |
| 6:27.170 | בזהר |  |
| 6:27.171 | בראשית |  |
| 6:27.172 | דף |  |
| 6:27.173 | כ"ב |  |
| 6:27.174 | או |  |
| 6:27.175 | אליהו |  |
| 6:27.176 | ורעיא |  |
| 6:27.177 | מהימנא |  |
| 6:27.178 | כמבואר |  |
| 6:27.179 | בזהר |  |
| 6:27.180 | בכמה |  |
| 6:27.181 | מקומות | . |
| 6:27.182 | כבר |  |
| 6:27.183 | הוזהרנו |  |
| 6:27.184 | בתורה |  |
| 6:27.185 | שלא |  |
| 6:27.186 | נאמין |  |
| 6:27.187 | לו |  |
| 6:27.188 | ולא |  |
| 6:27.189 | נשמע |  |
| 6:27.190 | אליו | , |
| 6:27.191 | כמו |  |
| 6:27.192 | שנ | ' |
| 6:27.193 | כי |  |
| 6:27.194 | יקום |  |
| 6:27.195 | בקרבך |  |
| 6:27.196 | נביא |  |
| 6:27.197 | או |  |
| 6:27.198 | חלם |  |
| 6:27.199 | חלום |  |
| 6:27.200 | וכו | ' |
| 6:27.201 | כי |  |
| 6:27.202 | נביא |  |
| 6:27.203 | שקר |  |
| 6:27.204 | הוא | , |
| 6:27.205 | ואף |  |
| 6:27.206 | אם |  |
| 6:27.207 | יעמיד |  |
| 6:27.208 | לנו |  |
| 6:27.209 | חמה |  |
| 6:27.210 | ולבנה |  |
| 6:27.211 | כיהושע |  |
| 6:27.212 | בגבעון |  |
| 6:27.213 | ועמק |  |
| 6:27.214 | אילון |  |
| 6:27.215 | לא |  |
| 6:27.216 | נאמין |  |
| 6:27.217 | לו |  |
| 6:27.218 | כמו |  |
| 6:27.219 | שאמרו |  |
| 6:27.220 | רז"ל |  |
| 6:27.221 | שגמרא |  |
| 6:27.222 | ומביאו |  |
| 6:27.223 | הר | ' |
| 6:27.224 | הסמ"ג |  |
| 6:27.225 | בהקדמת |  |
| 6:27.226 | העשין |  |
| 6:27.227 | והרמב"ם |  |
| 6:27.228 | בחבור |  |
| 6:27.229 | ובהקדמת |  |
| 6:27.230 | סדר |  |
| 6:27.231 | זרעים |  |
| 6:27.232 | ובמורה |  |
| 6:27.233 | פרק |  |
| 6:27.234 | כ"ד |  |
| 6:27.235 | משלישי |  |
| 6:27.236 | באורך | . |
| 6:27.237 | והטעם |  |
| 6:27.238 | כי |  |
| 6:27.239 | מנסה |  |
| 6:27.240 | ה | ' |
| 6:27.241 | אלהינו |  |
| 6:27.242 | אותנו |  |
| 6:27.243 | כמו |  |
| 6:27.244 | שנאמר |  |
| 6:27.245 | כי |  |
| 6:27.246 | מנסה |  |
| 6:27.247 | ה | ' |
| 6:27.248 | אלהיכם |  |
| 6:27.249 | אתכם |  |
| 6:27.250 | לדעת |  |
| 6:27.251 | הישכם |  |
| 6:27.252 | אוהבים |  |
| 6:27.253 | את |  |
| 6:27.254 | ה | ' |
| 6:27.255 | אלהיכם |  |
| 6:27.256 | וגו | ' |
| 6:27.257 | ואין |  |
| 6:27.258 | נביא |  |
| 6:27.259 | רשאי |  |
| 6:27.260 | לחדש |  |
| 6:27.261 | דבר |  |
| 6:27.262 | מעתה | , |
| 6:27.263 | ואף |  |
| 6:27.264 | אם |  |
| 6:27.265 | בת |  |
| 6:27.266 | קול |  |
| 6:27.267 | תכריז |  |
| 6:27.268 | מן |  |
| 6:27.269 | השמים |  |
| 6:27.270 | לאמר |  |
| 6:27.271 | שמעו |  |
| 6:27.272 | בקולו | , |
| 6:27.273 | לא |  |
| 6:27.274 | נשמע |  |
| 6:27.275 | לו |  |
| 6:27.276 | לסור |  |
| 6:27.277 | מן |  |
| 6:27.278 | הדרך |  |
| 6:27.279 | אשר |  |
| 6:27.280 | צוה |  |
| 6:27.281 | ה | ' |
| 6:27.282 | אלהינו |  |
| 6:27.283 | אותנו |  |
| 6:27.284 | ולזלזל |  |
| 6:27.285 | בשום |  |
| 6:27.286 | עיקר |  |
| 6:27.287 | מעקרי |  |
| 6:27.288 | התורה |  |
| 6:27.289 | או |  |
| 6:27.290 | בשום |  |
| 6:27.291 | מצוה |  |
| 6:27.292 | ממצותיה | , |
| 6:27.293 | כל |  |
| 6:27.294 | שכן |  |
| 6:27.295 | זה |  |
| 6:27.296 | המסית |  |
| 6:27.297 | בלי |  |
| 6:27.298 | שום |  |
| 6:27.299 | אות |  |
| 6:27.300 | ומופות |  |
| 6:27.301 | ועוקר |  |
| 6:27.302 | ד | ' |
| 6:27.303 | עיקרים |  |
| 6:27.304 | מיסודי |  |
| 6:27.305 | תורתינו |  |
| 6:27.306 | הקדושה |  |
| 6:27.307 | שכתב |  |
| 6:27.308 | הרמב"ם |  |
| 6:27.309 | ז"ל |  |
| 6:27.310 | כמו |  |
| 6:27.311 | שנת | ' |
| 6:27.312 | לעיל |  |
| 6:27.313 | סי | ' |
| 6:27.314 | מ"ח | , |
| 6:27.315 | והנה |  |
| 6:27.316 | כפי |  |
| 6:27.317 | הנחת |  |
| 6:27.318 | מחבר |  |
| 6:27.319 | הזהר |  |
| 6:27.320 | שכללות |  |
| 6:27.321 | הפרצופים |  |
| 6:27.322 | הם |  |
| 6:27.323 | אחד | , |
| 6:27.324 | איך |  |
| 6:27.325 | נפרד |  |
| 6:27.326 | העתיק |  |
| 6:27.327 | מאריך |  |
| 6:27.328 | אנפין |  |
| 6:27.329 | ושאר |  |
| 6:27.330 | פרצופים |  |
| 6:27.331 | ובא |  |
| 6:27.332 | לבדו |  |
| 6:27.333 | לבית |  |
| 6:27.334 | מדרשו |  |
| 6:27.335 | של |  |
| 6:27.336 | רשב"י |  |
| 6:27.337 | ושאל |  |
| 6:27.338 | אותו |  |
| 6:27.339 | ואמר |  |
| 6:27.340 | לו |  |
| 6:27.341 | שמעון |  |
| 6:27.342 | שמעון |  |
| 6:27.343 | מאן |  |
| 6:27.344 | הוא |  |
| 6:27.345 | דאמר |  |
| 6:27.346 | ויאמר |  |
| 6:27.347 | אלהים |  |
| 6:27.348 | נעשה |  |
| 6:27.349 | אדם |  |
| 6:27.350 | מאן |  |
| 6:27.351 | הוא |  |
| 6:27.352 | האי |  |
| 6:27.353 | אלהים |  |
| 6:27.354 | וכו | ' |
| 6:27.355 | זוהר |  |
| 6:27.356 | בראשית |  |
| 6:27.357 | כ״ב:ח׳ | ) |
| 6:27.358 | וכן |  |
| 6:27.359 | בכמה |  |
| 6:27.360 | מקומות |  |
| 6:27.361 | בזהר |  |
| 6:27.362 | שנזכר |  |
| 6:27.363 | סבא |  |
| 6:27.364 | דסבין |  |
| 6:27.365 | שהיה |  |
| 6:27.366 | בא |  |
| 6:27.367 | לגבי |  |
| 6:27.368 | רבי |  |
| 6:27.369 | אלעזר |  |
| 6:27.370 | איך |  |
| 6:27.371 | נפרד |  |
| 6:27.372 | ובא |  |
| 6:27.373 | לבדו |  |
| 6:27.374 | הלא |  |
| 6:27.375 | הכל |  |
| 6:27.376 | אחד | ? |
| 6:27.377 | וכבר |  |
| 6:27.378 | קדמוני |  |
| 6:27.379 | בטענה |  |
| 6:27.380 | זו |  |
| 6:27.381 | הראב"ע |  |
| 6:27.382 | ז"ל |  |
| 6:27.383 | על |  |
| 6:27.384 | מאמיני |  |
| 6:27.385 | השלוש |  |
| 6:27.386 | בריש |  |
| 6:27.387 | פרשת |  |
| 6:27.388 | וירא |  |
| 6:27.389 | וז"ל |  |
| 6:27.390 | הנה |  |
| 6:27.391 | קצת |  |
| 6:27.392 | אמרו |  |
| 6:27.393 | כי |  |
| 6:27.394 | השם |  |
| 6:27.395 | שלשה |  |
| 6:27.396 | אנשים |  |
| 6:27.397 | הוא |  |
| 6:27.398 | אחד |  |
| 6:27.399 | והוא |  |
| 6:27.400 | שלשה |  |
| 6:27.401 | ולא |  |
| 6:27.402 | יתפרדו | , |
| 6:27.403 | והנה |  |
| 6:27.404 | שכחו |  |
| 6:27.405 | ויבאו |  |
| 6:27.406 | שני |  |
| 6:27.407 | המלאכים |  |
| 6:27.408 | סדמה |  |
| 6:27.409 | ע"כ | . |
| 6:27.410 | ביאור |  |
| 6:27.411 | דבריו |  |
| 6:27.412 | כי |  |
| 6:27.413 | מאמיני |  |
| 6:27.414 | השלוש |  |
| 6:27.415 | סוברים |  |
| 6:27.416 | שהשם |  |
| 6:27.417 | ית | ' |
| 6:27.418 | הוא |  |
| 6:27.419 | שלשה | , |
| 6:27.420 | אב |  |
| 6:27.421 | ובן |  |
| 6:27.422 | ורוח |  |
| 6:27.423 | הקדש |  |
| 6:27.424 | והשלשה |  |
| 6:27.425 | הם |  |
| 6:27.426 | אחד |  |
| 6:27.427 | ולכך |  |
| 6:27.428 | אמר |  |
| 6:27.429 | הכתוב |  |
| 6:27.430 | וירא |  |
| 6:27.431 | אליו |  |
| 6:27.432 | ה | ', |
| 6:27.433 | והנה |  |
| 6:27.434 | שלשה |  |
| 6:27.435 | אנשים | , |
| 6:27.436 | ואלה |  |
| 6:27.437 | השלשה |  |
| 6:27.438 | הם |  |
| 6:27.439 | השם |  |
| 6:27.440 | הנז | ' |
| 6:27.441 | בפסוק |  |
| 6:27.442 | שקודם |  |
| 6:27.443 | זה |  |
| 6:27.444 | ולא |  |
| 6:27.445 | יתפרדו | , |
| 6:27.446 | וטען |  |
| 6:27.447 | עליהם |  |
| 6:27.448 | הראב"ע |  |
| 6:27.449 | האי |  |
| 6:27.450 | מצינו |  |
| 6:27.451 | שנפרדו |  |
| 6:27.452 | שנ | ' |
| 6:27.453 | ויבאו |  |
| 6:27.454 | שני |  |
| 6:27.455 | המלאכים |  |
| 6:27.456 | סדומה | ? |
| 6:27.457 | גם |  |
| 6:27.458 | אנחנו |  |
| 6:27.459 | נאמר |  |
| 6:27.460 | להמסית |  |
| 6:27.461 | מחבר |  |
| 6:27.462 | הזהר |  |
| 6:27.463 | שאומר |  |
| 6:27.464 | וכלא |  |
| 6:27.465 | יד |  |
| 6:27.466 | ולא |  |
| 6:27.467 | מתפרשין |  |
| 6:27.468 | לעלמין | . |
| 6:27.469 | איך |  |
| 6:27.470 | נפרד |  |
| 6:27.471 | עתיק |  |
| 6:27.472 | שהוא |  |
| 6:27.473 | סבא |  |
| 6:27.474 | דסבין |  |
| 6:27.475 | ובא |  |
| 6:27.476 | לבדו |  |
| 6:27.477 | לבית |  |
| 6:27.478 | מדרשו |  |
| 6:27.479 | של |  |
| 6:27.480 | רשב"י | ? |
| 6:27.481 | גם |  |
| 6:27.482 | מה |  |
| 6:27.483 | שאמר |  |
| 6:27.484 | שאימא |  |
| 6:27.485 | היא |  |
| 6:27.486 | שאמרה |  |
| 6:27.487 | לאבא |  |
| 6:27.488 | נעשה |  |
| 6:27.489 | אדם | , |
| 6:27.490 | וגברה |  |
| 6:27.491 | ידה |  |
| 6:27.492 | ועשאתו |  |
| 6:27.493 | על |  |
| 6:27.494 | כרחו |  |
| 6:27.495 | של |  |
| 6:27.496 | אבא | , |
| 6:27.497 | וכשחטא |  |
| 6:27.498 | אדם |  |
| 6:27.499 | גירשו |  |
| 6:27.500 | מגן |  |
| 6:27.501 | עדן |  |
| 6:27.502 | וגירשה |  |
| 6:27.503 | עמו | , |
| 6:27.504 | ואיך |  |
| 6:27.505 | נפרדה |  |
| 6:27.506 | אימא |  |
| 6:27.507 | מאבא | , |
| 6:27.508 | והלא |  |
| 6:27.509 | כלא |  |
| 6:27.510 | חד |  |
| 6:27.511 | ולא |  |
| 6:27.512 | מתפרשין | ? |
| 6:27.513 | והנה |  |
| 6:27.514 | אם |  |
| 6:27.515 | היה |  |
| 6:27.516 | סבא |  |
| 6:27.517 | דסבין |  |
| 6:27.518 | נראה |  |
| 6:27.519 | לרשב"י |  |
| 6:27.520 | ולר | ' |
| 6:27.521 | אלעזר |  |
| 6:27.522 | בנו |  |
| 6:27.523 | כמה |  |
| 6:27.524 | פעמים |  |
| 6:27.525 | כמו |  |
| 6:27.526 | שמבואר |  |
| 6:27.527 | בזהר | , |
| 6:27.528 | למה |  |
| 6:27.529 | בכה |  |
| 6:27.530 | רשב"י |  |
| 6:27.531 | ואמר |  |
| 6:27.532 | שפחה |  |
| 6:27.533 | של |  |
| 6:27.534 | בית |  |
| 6:27.535 | אבא |  |
| 6:27.536 | נזדמן |  |
| 6:27.537 | לה |  |
| 6:27.538 | מלאך |  |
| 6:27.539 | שלשה |  |
| 6:27.540 | פעמים | . |
| 6:27.541 | אני |  |
| 6:27.542 | אפי | ' |
| 6:27.543 | לא |  |
| 6:27.544 | פעם |  |
| 6:27.545 | אחת | ? |
| 6:27.546 | מעילה |  |
| 6:27.547 | דף |  |
| 6:27.548 | פ"ד | ). |
| 6:27.549 | וגם |  |
| 6:27.550 | למה |  |
| 6:27.551 | לא |  |
| 6:27.552 | הגידו |  |
| 6:27.553 | לנו |  |
| 6:27.554 | חכמים |  |
| 6:27.555 | שרשב"י |  |
| 6:27.556 | היה |  |
| 6:27.557 | נביא |  |
| 6:27.558 | גדול |  |
| 6:27.559 | כ"כ |  |
| 6:27.560 | יותר |  |
| 6:27.561 | משאר |  |
| 6:27.562 | נביאים | , |
| 6:27.563 | שהרי |  |
| 6:27.564 | היה |  |
| 6:27.565 | נגלה |  |
| 6:27.566 | לו |  |
| 6:27.567 | עתיקא |  |
| 6:27.568 | קדישא |  |
| 6:27.569 | והוא |  |
| 6:27.570 | עומד |  |
| 6:27.571 | על |  |
| 6:27.572 | עמדו |  |
| 6:27.573 | בתוך |  |
| 6:27.574 | החבירים |  |
| 6:27.575 | בבית |  |
| 6:27.576 | המדרש | ? |
| 6:27.577 | ולמה |  |
| 6:27.578 | אמרו |  |
| 6:27.579 | חכמים |  |
| 6:27.580 | שחגי |  |
| 6:27.581 | זכריה |  |
| 6:27.582 | ומלאכי |  |
| 6:27.583 | סוף |  |
| 6:27.584 | נביאים | , |
| 6:27.585 | ומשם |  |
| 6:27.586 | ואילך |  |
| 6:27.587 | פסקה |  |
| 6:27.588 | נבואה |  |
| 6:27.589 | ישראל | , |
| 6:27.590 | הא |  |
| 6:27.591 | אשכחן |  |
| 6:27.592 | רשב"י |  |
| 6:27.593 | ניבא |  |
| 6:27.594 | אחר |  |
| 6:27.595 | החרבן | , |
| 6:27.596 | והגיד |  |
| 6:27.597 | כמה |  |
| 6:27.598 | עניינים |  |
| 6:27.599 | באלהות | , |
| 6:27.600 | ושהעבודה |  |
| 6:27.601 | לבן |  |
| 6:27.602 | האלהים |  |
| 6:27.603 | קצר |  |
| 6:27.604 | האפים |  |
| 6:27.605 | מה |  |
| 6:27.606 | שלא |  |
| 6:27.607 | הודיעו |  |
| 6:27.608 | שאר |  |
| 6:27.609 | נביאים |  |
| 6:27.610 | מימות |  |
| 6:27.611 | משה |  |
| 6:27.612 | עד |  |
| 6:27.613 | מלאכי | , |
| 6:27.614 | ולא |  |
| 6:27.615 | נתגלה |  |
| 6:27.616 | זה |  |
| 6:27.617 | עד |  |
| 6:27.618 | האלף |  |
| 6:27.619 | הששי | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

