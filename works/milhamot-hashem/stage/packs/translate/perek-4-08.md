# Translation pack: milhamot-hashem perek-4-08 (Milchamot ha-Shem 4:13–4:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-08`.

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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
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
- באצילות: "in Emanation" (the noun the Kabbalists use), without "World of".
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


## Milchamot ha-Shem 4:13

Hebrew (Wikisource): מכל מאמרי הזהר ומפרשיו הנז' לעיל נתבאר שהם קורין לכל אחד מן הפרצופים שבאצילות הוי'ה ואדנות ואלהים והקב"ה. ובחרו לעבוד לפרצוף האחרון שהוא זעיר אנפין. ואמרו שהאין סוף עם כל הפרצופים הנאצלים ממנו לא שייכא שהו שום עבודה ולא קריאה, ושהמתפלל אליהם אינו נענה לרוב רוממותם, וכ"ש פרצופים שבעולמות שהם למעלה מעולם האצילות דלא שייכא בהו שום עבודה לרוב העלמם ורק בזעיר אנפין לבדו שייך עבודה ותפלה, וקריאה אליו בעת צרה, כי הוא עמוד האמצעי [13] המקשר כל הכחות שלמעלה ושלמטה, שכן גידלוהו אבא ואימא והשליטוהו על כל הברואים, וצוו לעבדו ולברכו, והוא לדעתם ה' אלהינו. והוא מבואר בזהר פרשת בלק דף קצ"א ע"ב (ח"ג קצא, ב) וז"ל עם פי' מקדש מלך על פסוק מונע בר יקבוהו לאום, ורזא דמלה כתיב ברזא דלעילא מה שמו ומה שם בנו כי תדע, ההוא שם ידיעה ה' צבאות שמו (פי' אבא) שם בנו ישראל שמו (פי' זעיר אנפין) דכתיב בני בכורי ישראל, והאי ישראל כל מפתחאן דמהימנותא בית תלין. ואיהו משתכח ואמר ה' אמר אלי בני אתה (פי' ז"א אומר שאבא הנקרא ה' צבאות אצר אלי בני אתה) והכי הוא ודאי. דהא אבא ואימא אעטרו ליה ובריכו ליה בכמה ברכאן. ואמרו ופקידו לכלא נשקו בר פן יאנף. נשקו להאי בר. כביכול שולטנו יהב ליה על כלא דכלא יפלחון דזכי לרחמי לרחמי. כל ברכאן דעילא ותתא לההוא בר סלקין ומתעטרן. ומאי דמנע ברכן מהאי בר יפרשון חטאוי קמי מלכא קדישא (ס"א אימא קדישא) לאם ממש (פירש בינה) ע"כ. הרי בהדיא שהזהר קרי לאבא ה' צבאות, ולזעיר אנפין קרי ליה בן לאבא ואימא. והוא הנקרא בשם הוי'ה בכמה מקומות בזהר. ושאבא ואימא נתנו לו הכח והממשלה בכל הנבראים וצוו לעבדו, ושכל ברכותינו ותפלותינו אליו דוקא יעלו להיות עטרה לז"א, ולא לאבא ואימא, ולא לאריך אנפין, ולא לעתיק, ולא לאדם קדמון אשר קראוהו עלת כל העלות (כדלעיל סי כ"א) וגם לא לאדם קדמאה שכל העולמות שלמעלה מעולם האצילות נתהוו בתוך עיגולי (מקדש מלך בראשית קט"ו ד"ה ורמז) וכ"ש להאין סוף שהוא רחוק וגבוה מעל גבוה למעלה מן הכל. שעליו אמרו דלא שייך ביה שום עבודה ותפלה וברכה, כי רק בז"א שייך ותלוי ההשגחה בתחתונים ובמעשיהם, ולא באלוהות העליונים. והוא זעיר אנפין לבדו המשלם שכר טוב לצדיקים ונפרע מן הרשעים. אבל כל הפרצופי שלמעלה עם האין סוף שהוא למעלה מן הכל אין משגיחים ומפקידין בכל מעשה התחתונים בין טוב לרע, וקצור קצרה ידם מפדות, ואין בהם כח להציל את הקורא אליהם, בעת צרה, עזב ה' את השמים ואת הארץ ח"ו, והפקידם ביד קצר האפים ז"א. שהוא נברא ולא בורא כמו שיתבאר עוד לקמן מדברי הזהר והמקובלים. ורז"ל דרשוהו לפסוק מי עלה שמים וירד באופן נאות הובא בילקוט משלי רמז תתקס"ב ובמדרש הגדול ריש פרשת שמות וז"ל מי עלה שמים זה קב"ה עלה אלהים בתרועה וירד. וירד ה' על הר סיני, מי אסף רוח בחפניו, אשר בידו נפש כל חי. מי צרר מים בשמלה, צורר מים בעביו, מי הקים כל אפסי ארץ ה' ממית ומחיה, מה שמו. צור שמו, שדי שמו, ה' צבאות שמו, ומה שם בנו, בני בכורי ישראל, והענין פשוט, שהוא על ישראל שנקראו בנים למקום, לפי שעל ידי ישראל נתפרסם אלהותו ית' בעולם, לא שיהא שם לז"א הנקרא אצל הפילוסוף בזהר בן והנה עתה נשלים דבר שהתחלנו בו.

| id | word | punctuation after |
|---|---|---|
| 4:13.1 | מכל |  |
| 4:13.2 | מאמרי |  |
| 4:13.3 | הזהר |  |
| 4:13.4 | ומפרשיו |  |
| 4:13.5 | הנז | ' |
| 4:13.6 | לעיל |  |
| 4:13.7 | נתבאר |  |
| 4:13.8 | שהם |  |
| 4:13.9 | קורין |  |
| 4:13.10 | לכל |  |
| 4:13.11 | אחד |  |
| 4:13.12 | מן |  |
| 4:13.13 | הפרצופים |  |
| 4:13.14 | שבאצילות |  |
| 4:13.15 | הוי'ה |  |
| 4:13.16 | ואדנות |  |
| 4:13.17 | ואלהים |  |
| 4:13.18 | והקב"ה | . |
| 4:13.19 | ובחרו |  |
| 4:13.20 | לעבוד |  |
| 4:13.21 | לפרצוף |  |
| 4:13.22 | האחרון |  |
| 4:13.23 | שהוא |  |
| 4:13.24 | זעיר |  |
| 4:13.25 | אנפין | . |
| 4:13.26 | ואמרו |  |
| 4:13.27 | שהאין |  |
| 4:13.28 | סוף |  |
| 4:13.29 | עם |  |
| 4:13.30 | כל |  |
| 4:13.31 | הפרצופים |  |
| 4:13.32 | הנאצלים |  |
| 4:13.33 | ממנו |  |
| 4:13.34 | לא |  |
| 4:13.35 | שייכא |  |
| 4:13.36 | שהו |  |
| 4:13.37 | שום |  |
| 4:13.38 | עבודה |  |
| 4:13.39 | ולא |  |
| 4:13.40 | קריאה | , |
| 4:13.41 | ושהמתפלל |  |
| 4:13.42 | אליהם |  |
| 4:13.43 | אינו |  |
| 4:13.44 | נענה |  |
| 4:13.45 | לרוב |  |
| 4:13.46 | רוממותם | , |
| 4:13.47 | וכ"ש |  |
| 4:13.48 | פרצופים |  |
| 4:13.49 | שבעולמות |  |
| 4:13.50 | שהם |  |
| 4:13.51 | למעלה |  |
| 4:13.52 | מעולם |  |
| 4:13.53 | האצילות |  |
| 4:13.54 | דלא |  |
| 4:13.55 | שייכא |  |
| 4:13.56 | בהו |  |
| 4:13.57 | שום |  |
| 4:13.58 | עבודה |  |
| 4:13.59 | לרוב |  |
| 4:13.60 | העלמם |  |
| 4:13.61 | ורק |  |
| 4:13.62 | בזעיר |  |
| 4:13.63 | אנפין |  |
| 4:13.64 | לבדו |  |
| 4:13.65 | שייך |  |
| 4:13.66 | עבודה |  |
| 4:13.67 | ותפלה | , |
| 4:13.68 | וקריאה |  |
| 4:13.69 | אליו |  |
| 4:13.70 | בעת |  |
| 4:13.71 | צרה | , |
| 4:13.72 | כי |  |
| 4:13.73 | הוא |  |
| 4:13.74 | עמוד |  |
| 4:13.75 | האמצעי | [13] |
| 4:13.76 | המקשר |  |
| 4:13.77 | כל |  |
| 4:13.78 | הכחות |  |
| 4:13.79 | שלמעלה |  |
| 4:13.80 | ושלמטה | , |
| 4:13.81 | שכן |  |
| 4:13.82 | גידלוהו |  |
| 4:13.83 | אבא |  |
| 4:13.84 | ואימא |  |
| 4:13.85 | והשליטוהו |  |
| 4:13.86 | על |  |
| 4:13.87 | כל |  |
| 4:13.88 | הברואים | , |
| 4:13.89 | וצוו |  |
| 4:13.90 | לעבדו |  |
| 4:13.91 | ולברכו | , |
| 4:13.92 | והוא |  |
| 4:13.93 | לדעתם |  |
| 4:13.94 | ה | ' |
| 4:13.95 | אלהינו | . |
| 4:13.96 | והוא |  |
| 4:13.97 | מבואר |  |
| 4:13.98 | בזהר |  |
| 4:13.99 | פרשת |  |
| 4:13.100 | בלק |  |
| 4:13.101 | דף |  |
| 4:13.102 | קצ"א |  |
| 4:13.103 | ע"ב |  |
| 4:13.104 | ח"ג |  |
| 4:13.105 | קצא | , |
| 4:13.106 | ב | ) |
| 4:13.107 | וז"ל |  |
| 4:13.108 | עם |  |
| 4:13.109 | פי | ' |
| 4:13.110 | מקדש |  |
| 4:13.111 | מלך |  |
| 4:13.112 | על |  |
| 4:13.113 | פסוק |  |
| 4:13.114 | מונע |  |
| 4:13.115 | בר |  |
| 4:13.116 | יקבוהו |  |
| 4:13.117 | לאום | , |
| 4:13.118 | ורזא |  |
| 4:13.119 | דמלה |  |
| 4:13.120 | כתיב |  |
| 4:13.121 | ברזא |  |
| 4:13.122 | דלעילא |  |
| 4:13.123 | מה |  |
| 4:13.124 | שמו |  |
| 4:13.125 | ומה |  |
| 4:13.126 | שם |  |
| 4:13.127 | בנו |  |
| 4:13.128 | כי |  |
| 4:13.129 | תדע | , |
| 4:13.130 | ההוא |  |
| 4:13.131 | שם |  |
| 4:13.132 | ידיעה |  |
| 4:13.133 | ה | ' |
| 4:13.134 | צבאות |  |
| 4:13.135 | שמו |  |
| 4:13.136 | פי | ' |
| 4:13.137 | אבא | ) |
| 4:13.138 | שם |  |
| 4:13.139 | בנו |  |
| 4:13.140 | ישראל |  |
| 4:13.141 | שמו |  |
| 4:13.142 | פי | ' |
| 4:13.143 | זעיר |  |
| 4:13.144 | אנפין | ) |
| 4:13.145 | דכתיב |  |
| 4:13.146 | בני |  |
| 4:13.147 | בכורי |  |
| 4:13.148 | ישראל | , |
| 4:13.149 | והאי |  |
| 4:13.150 | ישראל |  |
| 4:13.151 | כל |  |
| 4:13.152 | מפתחאן |  |
| 4:13.153 | דמהימנותא |  |
| 4:13.154 | בית |  |
| 4:13.155 | תלין | . |
| 4:13.156 | ואיהו |  |
| 4:13.157 | משתכח |  |
| 4:13.158 | ואמר |  |
| 4:13.159 | ה | ' |
| 4:13.160 | אמר |  |
| 4:13.161 | אלי |  |
| 4:13.162 | בני |  |
| 4:13.163 | אתה |  |
| 4:13.164 | פי | ' |
| 4:13.165 | ז"א |  |
| 4:13.166 | אומר |  |
| 4:13.167 | שאבא |  |
| 4:13.168 | הנקרא |  |
| 4:13.169 | ה | ' |
| 4:13.170 | צבאות |  |
| 4:13.171 | אצר |  |
| 4:13.172 | אלי |  |
| 4:13.173 | בני |  |
| 4:13.174 | אתה | ) |
| 4:13.175 | והכי |  |
| 4:13.176 | הוא |  |
| 4:13.177 | ודאי | . |
| 4:13.178 | דהא |  |
| 4:13.179 | אבא |  |
| 4:13.180 | ואימא |  |
| 4:13.181 | אעטרו |  |
| 4:13.182 | ליה |  |
| 4:13.183 | ובריכו |  |
| 4:13.184 | ליה |  |
| 4:13.185 | בכמה |  |
| 4:13.186 | ברכאן | . |
| 4:13.187 | ואמרו |  |
| 4:13.188 | ופקידו |  |
| 4:13.189 | לכלא |  |
| 4:13.190 | נשקו |  |
| 4:13.191 | בר |  |
| 4:13.192 | פן |  |
| 4:13.193 | יאנף | . |
| 4:13.194 | נשקו |  |
| 4:13.195 | להאי |  |
| 4:13.196 | בר | . |
| 4:13.197 | כביכול |  |
| 4:13.198 | שולטנו |  |
| 4:13.199 | יהב |  |
| 4:13.200 | ליה |  |
| 4:13.201 | על |  |
| 4:13.202 | כלא |  |
| 4:13.203 | דכלא |  |
| 4:13.204 | יפלחון |  |
| 4:13.205 | דזכי |  |
| 4:13.206 | לרחמי |  |
| 4:13.207 | לרחמי | . |
| 4:13.208 | כל |  |
| 4:13.209 | ברכאן |  |
| 4:13.210 | דעילא |  |
| 4:13.211 | ותתא |  |
| 4:13.212 | לההוא |  |
| 4:13.213 | בר |  |
| 4:13.214 | סלקין |  |
| 4:13.215 | ומתעטרן | . |
| 4:13.216 | ומאי |  |
| 4:13.217 | דמנע |  |
| 4:13.218 | ברכן |  |
| 4:13.219 | מהאי |  |
| 4:13.220 | בר |  |
| 4:13.221 | יפרשון |  |
| 4:13.222 | חטאוי |  |
| 4:13.223 | קמי |  |
| 4:13.224 | מלכא |  |
| 4:13.225 | קדישא |  |
| 4:13.226 | ס"א |  |
| 4:13.227 | אימא |  |
| 4:13.228 | קדישא | ) |
| 4:13.229 | לאם |  |
| 4:13.230 | ממש |  |
| 4:13.231 | פירש |  |
| 4:13.232 | בינה | ) |
| 4:13.233 | ע"כ | . |
| 4:13.234 | הרי |  |
| 4:13.235 | בהדיא |  |
| 4:13.236 | שהזהר |  |
| 4:13.237 | קרי |  |
| 4:13.238 | לאבא |  |
| 4:13.239 | ה | ' |
| 4:13.240 | צבאות | , |
| 4:13.241 | ולזעיר |  |
| 4:13.242 | אנפין |  |
| 4:13.243 | קרי |  |
| 4:13.244 | ליה |  |
| 4:13.245 | בן |  |
| 4:13.246 | לאבא |  |
| 4:13.247 | ואימא | . |
| 4:13.248 | והוא |  |
| 4:13.249 | הנקרא |  |
| 4:13.250 | בשם |  |
| 4:13.251 | הוי'ה |  |
| 4:13.252 | בכמה |  |
| 4:13.253 | מקומות |  |
| 4:13.254 | בזהר | . |
| 4:13.255 | ושאבא |  |
| 4:13.256 | ואימא |  |
| 4:13.257 | נתנו |  |
| 4:13.258 | לו |  |
| 4:13.259 | הכח |  |
| 4:13.260 | והממשלה |  |
| 4:13.261 | בכל |  |
| 4:13.262 | הנבראים |  |
| 4:13.263 | וצוו |  |
| 4:13.264 | לעבדו | , |
| 4:13.265 | ושכל |  |
| 4:13.266 | ברכותינו |  |
| 4:13.267 | ותפלותינו |  |
| 4:13.268 | אליו |  |
| 4:13.269 | דוקא |  |
| 4:13.270 | יעלו |  |
| 4:13.271 | להיות |  |
| 4:13.272 | עטרה |  |
| 4:13.273 | לז"א | , |
| 4:13.274 | ולא |  |
| 4:13.275 | לאבא |  |
| 4:13.276 | ואימא | , |
| 4:13.277 | ולא |  |
| 4:13.278 | לאריך |  |
| 4:13.279 | אנפין | , |
| 4:13.280 | ולא |  |
| 4:13.281 | לעתיק | , |
| 4:13.282 | ולא |  |
| 4:13.283 | לאדם |  |
| 4:13.284 | קדמון |  |
| 4:13.285 | אשר |  |
| 4:13.286 | קראוהו |  |
| 4:13.287 | עלת |  |
| 4:13.288 | כל |  |
| 4:13.289 | העלות |  |
| 4:13.290 | כדלעיל |  |
| 4:13.291 | סי |  |
| 4:13.292 | כ"א | ) |
| 4:13.293 | וגם |  |
| 4:13.294 | לא |  |
| 4:13.295 | לאדם |  |
| 4:13.296 | קדמאה |  |
| 4:13.297 | שכל |  |
| 4:13.298 | העולמות |  |
| 4:13.299 | שלמעלה |  |
| 4:13.300 | מעולם |  |
| 4:13.301 | האצילות |  |
| 4:13.302 | נתהוו |  |
| 4:13.303 | בתוך |  |
| 4:13.304 | עיגולי |  |
| 4:13.305 | מקדש |  |
| 4:13.306 | מלך |  |
| 4:13.307 | בראשית |  |
| 4:13.308 | קט"ו |  |
| 4:13.309 | ד"ה |  |
| 4:13.310 | ורמז | ) |
| 4:13.311 | וכ"ש |  |
| 4:13.312 | להאין |  |
| 4:13.313 | סוף |  |
| 4:13.314 | שהוא |  |
| 4:13.315 | רחוק |  |
| 4:13.316 | וגבוה |  |
| 4:13.317 | מעל |  |
| 4:13.318 | גבוה |  |
| 4:13.319 | למעלה |  |
| 4:13.320 | מן |  |
| 4:13.321 | הכל | . |
| 4:13.322 | שעליו |  |
| 4:13.323 | אמרו |  |
| 4:13.324 | דלא |  |
| 4:13.325 | שייך |  |
| 4:13.326 | ביה |  |
| 4:13.327 | שום |  |
| 4:13.328 | עבודה |  |
| 4:13.329 | ותפלה |  |
| 4:13.330 | וברכה | , |
| 4:13.331 | כי |  |
| 4:13.332 | רק |  |
| 4:13.333 | בז"א |  |
| 4:13.334 | שייך |  |
| 4:13.335 | ותלוי |  |
| 4:13.336 | ההשגחה |  |
| 4:13.337 | בתחתונים |  |
| 4:13.338 | ובמעשיהם | , |
| 4:13.339 | ולא |  |
| 4:13.340 | באלוהות |  |
| 4:13.341 | העליונים | . |
| 4:13.342 | והוא |  |
| 4:13.343 | זעיר |  |
| 4:13.344 | אנפין |  |
| 4:13.345 | לבדו |  |
| 4:13.346 | המשלם |  |
| 4:13.347 | שכר |  |
| 4:13.348 | טוב |  |
| 4:13.349 | לצדיקים |  |
| 4:13.350 | ונפרע |  |
| 4:13.351 | מן |  |
| 4:13.352 | הרשעים | . |
| 4:13.353 | אבל |  |
| 4:13.354 | כל |  |
| 4:13.355 | הפרצופי |  |
| 4:13.356 | שלמעלה |  |
| 4:13.357 | עם |  |
| 4:13.358 | האין |  |
| 4:13.359 | סוף |  |
| 4:13.360 | שהוא |  |
| 4:13.361 | למעלה |  |
| 4:13.362 | מן |  |
| 4:13.363 | הכל |  |
| 4:13.364 | אין |  |
| 4:13.365 | משגיחים |  |
| 4:13.366 | ומפקידין |  |
| 4:13.367 | בכל |  |
| 4:13.368 | מעשה |  |
| 4:13.369 | התחתונים |  |
| 4:13.370 | בין |  |
| 4:13.371 | טוב |  |
| 4:13.372 | לרע | , |
| 4:13.373 | וקצור |  |
| 4:13.374 | קצרה |  |
| 4:13.375 | ידם |  |
| 4:13.376 | מפדות | , |
| 4:13.377 | ואין |  |
| 4:13.378 | בהם |  |
| 4:13.379 | כח |  |
| 4:13.380 | להציל |  |
| 4:13.381 | את |  |
| 4:13.382 | הקורא |  |
| 4:13.383 | אליהם | , |
| 4:13.384 | בעת |  |
| 4:13.385 | צרה | , |
| 4:13.386 | עזב |  |
| 4:13.387 | ה | ' |
| 4:13.388 | את |  |
| 4:13.389 | השמים |  |
| 4:13.390 | ואת |  |
| 4:13.391 | הארץ |  |
| 4:13.392 | ח"ו | , |
| 4:13.393 | והפקידם |  |
| 4:13.394 | ביד |  |
| 4:13.395 | קצר |  |
| 4:13.396 | האפים |  |
| 4:13.397 | ז"א | . |
| 4:13.398 | שהוא |  |
| 4:13.399 | נברא |  |
| 4:13.400 | ולא |  |
| 4:13.401 | בורא |  |
| 4:13.402 | כמו |  |
| 4:13.403 | שיתבאר |  |
| 4:13.404 | עוד |  |
| 4:13.405 | לקמן |  |
| 4:13.406 | מדברי |  |
| 4:13.407 | הזהר |  |
| 4:13.408 | והמקובלים | . |
| 4:13.409 | ורז"ל |  |
| 4:13.410 | דרשוהו |  |
| 4:13.411 | לפסוק |  |
| 4:13.412 | מי |  |
| 4:13.413 | עלה |  |
| 4:13.414 | שמים |  |
| 4:13.415 | וירד |  |
| 4:13.416 | באופן |  |
| 4:13.417 | נאות |  |
| 4:13.418 | הובא |  |
| 4:13.419 | בילקוט |  |
| 4:13.420 | משלי |  |
| 4:13.421 | רמז |  |
| 4:13.422 | תתקס"ב |  |
| 4:13.423 | ובמדרש |  |
| 4:13.424 | הגדול |  |
| 4:13.425 | ריש |  |
| 4:13.426 | פרשת |  |
| 4:13.427 | שמות |  |
| 4:13.428 | וז"ל |  |
| 4:13.429 | מי |  |
| 4:13.430 | עלה |  |
| 4:13.431 | שמים |  |
| 4:13.432 | זה |  |
| 4:13.433 | קב"ה |  |
| 4:13.434 | עלה |  |
| 4:13.435 | אלהים |  |
| 4:13.436 | בתרועה |  |
| 4:13.437 | וירד | . |
| 4:13.438 | וירד |  |
| 4:13.439 | ה | ' |
| 4:13.440 | על |  |
| 4:13.441 | הר |  |
| 4:13.442 | סיני | , |
| 4:13.443 | מי |  |
| 4:13.444 | אסף |  |
| 4:13.445 | רוח |  |
| 4:13.446 | בחפניו | , |
| 4:13.447 | אשר |  |
| 4:13.448 | בידו |  |
| 4:13.449 | נפש |  |
| 4:13.450 | כל |  |
| 4:13.451 | חי | . |
| 4:13.452 | מי |  |
| 4:13.453 | צרר |  |
| 4:13.454 | מים |  |
| 4:13.455 | בשמלה | , |
| 4:13.456 | צורר |  |
| 4:13.457 | מים |  |
| 4:13.458 | בעביו | , |
| 4:13.459 | מי |  |
| 4:13.460 | הקים |  |
| 4:13.461 | כל |  |
| 4:13.462 | אפסי |  |
| 4:13.463 | ארץ |  |
| 4:13.464 | ה | ' |
| 4:13.465 | ממית |  |
| 4:13.466 | ומחיה | , |
| 4:13.467 | מה |  |
| 4:13.468 | שמו | . |
| 4:13.469 | צור |  |
| 4:13.470 | שמו | , |
| 4:13.471 | שדי |  |
| 4:13.472 | שמו | , |
| 4:13.473 | ה | ' |
| 4:13.474 | צבאות |  |
| 4:13.475 | שמו | , |
| 4:13.476 | ומה |  |
| 4:13.477 | שם |  |
| 4:13.478 | בנו | , |
| 4:13.479 | בני |  |
| 4:13.480 | בכורי |  |
| 4:13.481 | ישראל | , |
| 4:13.482 | והענין |  |
| 4:13.483 | פשוט | , |
| 4:13.484 | שהוא |  |
| 4:13.485 | על |  |
| 4:13.486 | ישראל |  |
| 4:13.487 | שנקראו |  |
| 4:13.488 | בנים |  |
| 4:13.489 | למקום | , |
| 4:13.490 | לפי |  |
| 4:13.491 | שעל |  |
| 4:13.492 | ידי |  |
| 4:13.493 | ישראל |  |
| 4:13.494 | נתפרסם |  |
| 4:13.495 | אלהותו |  |
| 4:13.496 | ית | ' |
| 4:13.497 | בעולם | , |
| 4:13.498 | לא |  |
| 4:13.499 | שיהא |  |
| 4:13.500 | שם |  |
| 4:13.501 | לז"א |  |
| 4:13.502 | הנקרא |  |
| 4:13.503 | אצל |  |
| 4:13.504 | הפילוסוף |  |
| 4:13.505 | בזהר |  |
| 4:13.506 | בן |  |
| 4:13.507 | והנה |  |
| 4:13.508 | עתה |  |
| 4:13.509 | נשלים |  |
| 4:13.510 | דבר |  |
| 4:13.511 | שהתחלנו |  |
| 4:13.512 | בו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

