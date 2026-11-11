# Translation pack: milhamot-hashem perek-7-03 (Milchamot ha-Shem 7:4–7:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-03`.

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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 7:4

Hebrew (Wikisource): פא) זכור נזכור את אמונת רז"ל בעלי המשנה והתלמוד הגאונים ז"ל שהיא האמונה הטהורה שהאמינו בה קדמונינו, כי ה' אלהינו אחד, אין עוד מלבדו ואין כיחודו, ותשוח עלינו נפשינו בראותינו עתה ספרי הקבלה החדשה כמה אלוהות נתחדשו להם, אשר לא שערום אבותינו, והמה בכתובים גם בסידורי התפלות כדי לכוון תפלתם אליהם, וכן באשר הברכות, ולא די להם המון האלוהות שמכונים אליהם כמו שכבר ביארנו, אלא שאיוו לאלוהות רבים, אשר אין לקרות בהם בשם, שהרי לפי דעתם גם אדם קדמון ואדם קדמאה אין לכוון תפלתו אליהם, כי הם נעלמים מאד, וכל שכן פרצופים שלמעלה מעולם האצילות אלפים ורבבות אין מספר, עולמות שלא דברו בהם, וכל עולם ועולם כלול מאבי"ע, ומי יודע אם יש בכל עולם ועולם כל הפרצופים שהזכירו באבי"ע זה התחתון כי לא מצינו זה להם בפירוש לפי שהעלימו ולא דברו בהם, ורק זה מצינו להר' מקדש מלך בראשי' דך ט"ו דבור המתחיל ורמז, כי החלק המקיף וסמוך לטהירו נקרא אדם קדמאה סתימאה, להבדיל בינו ובין אדם קדמון, ואדם קדמאה סתימאה דא נעשה עשרה כדורים זה בתוך זה, ובכל אחד מכדורים אלו נבנו כמה עולמות אין מספר וכו'. ונמצא לפי זה מערכת האלוהות הרבים אשר כל אחד ואחד עלה וסבה למה שלמטה ממנו, ראשון לכלם האין סוף המקיף את כל העולמות, וכולם בתוך חלל שלו, ואחריו אדם קדמאה סתימה, ואחריו אדם קדמון, הוא שאמר לפי דבריהם ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין צריך ליטול רשות מאדם קדמאה ולא מאין סוף, ואחריו עתיק יומין, ואחריו אריך אנפין והוא אשר ברא שמים וארץ בכח אבא ואימא, ואחריו אבא, הוא שאמר יהי אור, ויקוו המים, ויהי רקיע, ושאר אמירות דמעשה בראשית. ואימא אומן העושה על פיו ככל אשר יצוה עליה. חוץ ממאמר נעשה אדם. שאימא היא שאמרה לאבא נעשה אדם בצלמינו. ואבא לא חפץ בבריאתו. לפי שגלוי לפניו שעתיד לחטוא כמו שכתבתי לעיל סי' כ'. ואחר אבא ואימא זעיר אנפין ונוקביה. ואל זעיר ונוקביה לפי דבריהם כל פיות פונים לעבדו. וליראה את מלכתו היא נוקבא דיליה שריבה אותה מחבר הזהר ליראה מאת ה' אלהיך תירא. את לרבותה, ליראה ממנה ולא לעבדה (וד"א כמו שדרש ר' עקיבה לרבות ת"ח). ונתנו המקובלים החדשים טעם למה לא הזכירו בזהר ובתקונים את הפרצופים העליונים כי אם ברמז, שהוא לרוב העלמם. והוא בעצמו הטעם למה אין מתפללים אליהם, ולא קורים אליהם בעת צרה מטעם רוב העלמם. ולפי דבריהם לא שייכא עבודה ותפלה אלא לפרצוף שכבר נתעבה ונתגלה וכמו שכתב בספר הברית שהוא מושג לבית יעקב. כלומר שכבר נתעבה ונתגלה כמו שכתב הרח"ו והובא לעיל סי' מ"ז. ועל הטעם הזה יפלא מאד הלא ה' אלהינו שהוא הסבה הראשונה אשר עבדו אבותינו ואבות אבותינו נעלם הוא מאד עד שאפי' המלאכים וחיות הקדש שואלים איה מקום כבודו כמו שאמרו רז"ל ועם כל זה אמר הכתוב כי מי גוי גדול אשר לו אלהים קרובים אליו כה' אלהינו בכל קראנו אליו. כמפה לאזן כמו שנת' לעיל:

| id | word | punctuation after |
|---|---|---|
| 7:4.1 | פא | ) |
| 7:4.2 | זכור |  |
| 7:4.3 | נזכור |  |
| 7:4.4 | את |  |
| 7:4.5 | אמונת |  |
| 7:4.6 | רז"ל |  |
| 7:4.7 | בעלי |  |
| 7:4.8 | המשנה |  |
| 7:4.9 | והתלמוד |  |
| 7:4.10 | הגאונים |  |
| 7:4.11 | ז"ל |  |
| 7:4.12 | שהיא |  |
| 7:4.13 | האמונה |  |
| 7:4.14 | הטהורה |  |
| 7:4.15 | שהאמינו |  |
| 7:4.16 | בה |  |
| 7:4.17 | קדמונינו | , |
| 7:4.18 | כי |  |
| 7:4.19 | ה | ' |
| 7:4.20 | אלהינו |  |
| 7:4.21 | אחד | , |
| 7:4.22 | אין |  |
| 7:4.23 | עוד |  |
| 7:4.24 | מלבדו |  |
| 7:4.25 | ואין |  |
| 7:4.26 | כיחודו | , |
| 7:4.27 | ותשוח |  |
| 7:4.28 | עלינו |  |
| 7:4.29 | נפשינו |  |
| 7:4.30 | בראותינו |  |
| 7:4.31 | עתה |  |
| 7:4.32 | ספרי |  |
| 7:4.33 | הקבלה |  |
| 7:4.34 | החדשה |  |
| 7:4.35 | כמה |  |
| 7:4.36 | אלוהות |  |
| 7:4.37 | נתחדשו |  |
| 7:4.38 | להם | , |
| 7:4.39 | אשר |  |
| 7:4.40 | לא |  |
| 7:4.41 | שערום |  |
| 7:4.42 | אבותינו | , |
| 7:4.43 | והמה |  |
| 7:4.44 | בכתובים |  |
| 7:4.45 | גם |  |
| 7:4.46 | בסידורי |  |
| 7:4.47 | התפלות |  |
| 7:4.48 | כדי |  |
| 7:4.49 | לכוון |  |
| 7:4.50 | תפלתם |  |
| 7:4.51 | אליהם | , |
| 7:4.52 | וכן |  |
| 7:4.53 | באשר |  |
| 7:4.54 | הברכות | , |
| 7:4.55 | ולא |  |
| 7:4.56 | די |  |
| 7:4.57 | להם |  |
| 7:4.58 | המון |  |
| 7:4.59 | האלוהות |  |
| 7:4.60 | שמכונים |  |
| 7:4.61 | אליהם |  |
| 7:4.62 | כמו |  |
| 7:4.63 | שכבר |  |
| 7:4.64 | ביארנו | , |
| 7:4.65 | אלא |  |
| 7:4.66 | שאיוו |  |
| 7:4.67 | לאלוהות |  |
| 7:4.68 | רבים | , |
| 7:4.69 | אשר |  |
| 7:4.70 | אין |  |
| 7:4.71 | לקרות |  |
| 7:4.72 | בהם |  |
| 7:4.73 | בשם | , |
| 7:4.74 | שהרי |  |
| 7:4.75 | לפי |  |
| 7:4.76 | דעתם |  |
| 7:4.77 | גם |  |
| 7:4.78 | אדם |  |
| 7:4.79 | קדמון |  |
| 7:4.80 | ואדם |  |
| 7:4.81 | קדמאה |  |
| 7:4.82 | אין |  |
| 7:4.83 | לכוון |  |
| 7:4.84 | תפלתו |  |
| 7:4.85 | אליהם | , |
| 7:4.86 | כי |  |
| 7:4.87 | הם |  |
| 7:4.88 | נעלמים |  |
| 7:4.89 | מאד | , |
| 7:4.90 | וכל |  |
| 7:4.91 | שכן |  |
| 7:4.92 | פרצופים |  |
| 7:4.93 | שלמעלה |  |
| 7:4.94 | מעולם |  |
| 7:4.95 | האצילות |  |
| 7:4.96 | אלפים |  |
| 7:4.97 | ורבבות |  |
| 7:4.98 | אין |  |
| 7:4.99 | מספר | , |
| 7:4.100 | עולמות |  |
| 7:4.101 | שלא |  |
| 7:4.102 | דברו |  |
| 7:4.103 | בהם | , |
| 7:4.104 | וכל |  |
| 7:4.105 | עולם |  |
| 7:4.106 | ועולם |  |
| 7:4.107 | כלול |  |
| 7:4.108 | מאבי"ע | , |
| 7:4.109 | ומי |  |
| 7:4.110 | יודע |  |
| 7:4.111 | אם |  |
| 7:4.112 | יש |  |
| 7:4.113 | בכל |  |
| 7:4.114 | עולם |  |
| 7:4.115 | ועולם |  |
| 7:4.116 | כל |  |
| 7:4.117 | הפרצופים |  |
| 7:4.118 | שהזכירו |  |
| 7:4.119 | באבי"ע |  |
| 7:4.120 | זה |  |
| 7:4.121 | התחתון |  |
| 7:4.122 | כי |  |
| 7:4.123 | לא |  |
| 7:4.124 | מצינו |  |
| 7:4.125 | זה |  |
| 7:4.126 | להם |  |
| 7:4.127 | בפירוש |  |
| 7:4.128 | לפי |  |
| 7:4.129 | שהעלימו |  |
| 7:4.130 | ולא |  |
| 7:4.131 | דברו |  |
| 7:4.132 | בהם | , |
| 7:4.133 | ורק |  |
| 7:4.134 | זה |  |
| 7:4.135 | מצינו |  |
| 7:4.136 | להר | ' |
| 7:4.137 | מקדש |  |
| 7:4.138 | מלך |  |
| 7:4.139 | בראשי | ' |
| 7:4.140 | דך |  |
| 7:4.141 | ט"ו |  |
| 7:4.142 | דבור |  |
| 7:4.143 | המתחיל |  |
| 7:4.144 | ורמז | , |
| 7:4.145 | כי |  |
| 7:4.146 | החלק |  |
| 7:4.147 | המקיף |  |
| 7:4.148 | וסמוך |  |
| 7:4.149 | לטהירו |  |
| 7:4.150 | נקרא |  |
| 7:4.151 | אדם |  |
| 7:4.152 | קדמאה |  |
| 7:4.153 | סתימאה | , |
| 7:4.154 | להבדיל |  |
| 7:4.155 | בינו |  |
| 7:4.156 | ובין |  |
| 7:4.157 | אדם |  |
| 7:4.158 | קדמון | , |
| 7:4.159 | ואדם |  |
| 7:4.160 | קדמאה |  |
| 7:4.161 | סתימאה |  |
| 7:4.162 | דא |  |
| 7:4.163 | נעשה |  |
| 7:4.164 | עשרה |  |
| 7:4.165 | כדורים |  |
| 7:4.166 | זה |  |
| 7:4.167 | בתוך |  |
| 7:4.168 | זה | , |
| 7:4.169 | ובכל |  |
| 7:4.170 | אחד |  |
| 7:4.171 | מכדורים |  |
| 7:4.172 | אלו |  |
| 7:4.173 | נבנו |  |
| 7:4.174 | כמה |  |
| 7:4.175 | עולמות |  |
| 7:4.176 | אין |  |
| 7:4.177 | מספר |  |
| 7:4.178 | וכו | '. |
| 7:4.179 | ונמצא |  |
| 7:4.180 | לפי |  |
| 7:4.181 | זה |  |
| 7:4.182 | מערכת |  |
| 7:4.183 | האלוהות |  |
| 7:4.184 | הרבים |  |
| 7:4.185 | אשר |  |
| 7:4.186 | כל |  |
| 7:4.187 | אחד |  |
| 7:4.188 | ואחד |  |
| 7:4.189 | עלה |  |
| 7:4.190 | וסבה |  |
| 7:4.191 | למה |  |
| 7:4.192 | שלמטה |  |
| 7:4.193 | ממנו | , |
| 7:4.194 | ראשון |  |
| 7:4.195 | לכלם |  |
| 7:4.196 | האין |  |
| 7:4.197 | סוף |  |
| 7:4.198 | המקיף |  |
| 7:4.199 | את |  |
| 7:4.200 | כל |  |
| 7:4.201 | העולמות | , |
| 7:4.202 | וכולם |  |
| 7:4.203 | בתוך |  |
| 7:4.204 | חלל |  |
| 7:4.205 | שלו | , |
| 7:4.206 | ואחריו |  |
| 7:4.207 | אדם |  |
| 7:4.208 | קדמאה |  |
| 7:4.209 | סתימה | , |
| 7:4.210 | ואחריו |  |
| 7:4.211 | אדם |  |
| 7:4.212 | קדמון | , |
| 7:4.213 | הוא |  |
| 7:4.214 | שאמר |  |
| 7:4.215 | לפי |  |
| 7:4.216 | דבריהם |  |
| 7:4.217 | ראו |  |
| 7:4.218 | עתה |  |
| 7:4.219 | כי |  |
| 7:4.220 | אני |  |
| 7:4.221 | אני |  |
| 7:4.222 | הוא |  |
| 7:4.223 | ואין |  |
| 7:4.224 | אלהים |  |
| 7:4.225 | עמדי | , |
| 7:4.226 | שאין |  |
| 7:4.227 | צריך |  |
| 7:4.228 | ליטול |  |
| 7:4.229 | רשות |  |
| 7:4.230 | מאדם |  |
| 7:4.231 | קדמאה |  |
| 7:4.232 | ולא |  |
| 7:4.233 | מאין |  |
| 7:4.234 | סוף | , |
| 7:4.235 | ואחריו |  |
| 7:4.236 | עתיק |  |
| 7:4.237 | יומין | , |
| 7:4.238 | ואחריו |  |
| 7:4.239 | אריך |  |
| 7:4.240 | אנפין |  |
| 7:4.241 | והוא |  |
| 7:4.242 | אשר |  |
| 7:4.243 | ברא |  |
| 7:4.244 | שמים |  |
| 7:4.245 | וארץ |  |
| 7:4.246 | בכח |  |
| 7:4.247 | אבא |  |
| 7:4.248 | ואימא | , |
| 7:4.249 | ואחריו |  |
| 7:4.250 | אבא | , |
| 7:4.251 | הוא |  |
| 7:4.252 | שאמר |  |
| 7:4.253 | יהי |  |
| 7:4.254 | אור | , |
| 7:4.255 | ויקוו |  |
| 7:4.256 | המים | , |
| 7:4.257 | ויהי |  |
| 7:4.258 | רקיע | , |
| 7:4.259 | ושאר |  |
| 7:4.260 | אמירות |  |
| 7:4.261 | דמעשה |  |
| 7:4.262 | בראשית | . |
| 7:4.263 | ואימא |  |
| 7:4.264 | אומן |  |
| 7:4.265 | העושה |  |
| 7:4.266 | על |  |
| 7:4.267 | פיו |  |
| 7:4.268 | ככל |  |
| 7:4.269 | אשר |  |
| 7:4.270 | יצוה |  |
| 7:4.271 | עליה | . |
| 7:4.272 | חוץ |  |
| 7:4.273 | ממאמר |  |
| 7:4.274 | נעשה |  |
| 7:4.275 | אדם | . |
| 7:4.276 | שאימא |  |
| 7:4.277 | היא |  |
| 7:4.278 | שאמרה |  |
| 7:4.279 | לאבא |  |
| 7:4.280 | נעשה |  |
| 7:4.281 | אדם |  |
| 7:4.282 | בצלמינו | . |
| 7:4.283 | ואבא |  |
| 7:4.284 | לא |  |
| 7:4.285 | חפץ |  |
| 7:4.286 | בבריאתו | . |
| 7:4.287 | לפי |  |
| 7:4.288 | שגלוי |  |
| 7:4.289 | לפניו |  |
| 7:4.290 | שעתיד |  |
| 7:4.291 | לחטוא |  |
| 7:4.292 | כמו |  |
| 7:4.293 | שכתבתי |  |
| 7:4.294 | לעיל |  |
| 7:4.295 | סי | ' |
| 7:4.296 | כ | '. |
| 7:4.297 | ואחר |  |
| 7:4.298 | אבא |  |
| 7:4.299 | ואימא |  |
| 7:4.300 | זעיר |  |
| 7:4.301 | אנפין |  |
| 7:4.302 | ונוקביה | . |
| 7:4.303 | ואל |  |
| 7:4.304 | זעיר |  |
| 7:4.305 | ונוקביה |  |
| 7:4.306 | לפי |  |
| 7:4.307 | דבריהם |  |
| 7:4.308 | כל |  |
| 7:4.309 | פיות |  |
| 7:4.310 | פונים |  |
| 7:4.311 | לעבדו | . |
| 7:4.312 | וליראה |  |
| 7:4.313 | את |  |
| 7:4.314 | מלכתו |  |
| 7:4.315 | היא |  |
| 7:4.316 | נוקבא |  |
| 7:4.317 | דיליה |  |
| 7:4.318 | שריבה |  |
| 7:4.319 | אותה |  |
| 7:4.320 | מחבר |  |
| 7:4.321 | הזהר |  |
| 7:4.322 | ליראה |  |
| 7:4.323 | מאת |  |
| 7:4.324 | ה | ' |
| 7:4.325 | אלהיך |  |
| 7:4.326 | תירא | . |
| 7:4.327 | את |  |
| 7:4.328 | לרבותה | , |
| 7:4.329 | ליראה |  |
| 7:4.330 | ממנה |  |
| 7:4.331 | ולא |  |
| 7:4.332 | לעבדה |  |
| 7:4.333 | וד"א |  |
| 7:4.334 | כמו |  |
| 7:4.335 | שדרש |  |
| 7:4.336 | ר | ' |
| 7:4.337 | עקיבה |  |
| 7:4.338 | לרבות |  |
| 7:4.339 | ת"ח | ). |
| 7:4.340 | ונתנו |  |
| 7:4.341 | המקובלים |  |
| 7:4.342 | החדשים |  |
| 7:4.343 | טעם |  |
| 7:4.344 | למה |  |
| 7:4.345 | לא |  |
| 7:4.346 | הזכירו |  |
| 7:4.347 | בזהר |  |
| 7:4.348 | ובתקונים |  |
| 7:4.349 | את |  |
| 7:4.350 | הפרצופים |  |
| 7:4.351 | העליונים |  |
| 7:4.352 | כי |  |
| 7:4.353 | אם |  |
| 7:4.354 | ברמז | , |
| 7:4.355 | שהוא |  |
| 7:4.356 | לרוב |  |
| 7:4.357 | העלמם | . |
| 7:4.358 | והוא |  |
| 7:4.359 | בעצמו |  |
| 7:4.360 | הטעם |  |
| 7:4.361 | למה |  |
| 7:4.362 | אין |  |
| 7:4.363 | מתפללים |  |
| 7:4.364 | אליהם | , |
| 7:4.365 | ולא |  |
| 7:4.366 | קורים |  |
| 7:4.367 | אליהם |  |
| 7:4.368 | בעת |  |
| 7:4.369 | צרה |  |
| 7:4.370 | מטעם |  |
| 7:4.371 | רוב |  |
| 7:4.372 | העלמם | . |
| 7:4.373 | ולפי |  |
| 7:4.374 | דבריהם |  |
| 7:4.375 | לא |  |
| 7:4.376 | שייכא |  |
| 7:4.377 | עבודה |  |
| 7:4.378 | ותפלה |  |
| 7:4.379 | אלא |  |
| 7:4.380 | לפרצוף |  |
| 7:4.381 | שכבר |  |
| 7:4.382 | נתעבה |  |
| 7:4.383 | ונתגלה |  |
| 7:4.384 | וכמו |  |
| 7:4.385 | שכתב |  |
| 7:4.386 | בספר |  |
| 7:4.387 | הברית |  |
| 7:4.388 | שהוא |  |
| 7:4.389 | מושג |  |
| 7:4.390 | לבית |  |
| 7:4.391 | יעקב | . |
| 7:4.392 | כלומר |  |
| 7:4.393 | שכבר |  |
| 7:4.394 | נתעבה |  |
| 7:4.395 | ונתגלה |  |
| 7:4.396 | כמו |  |
| 7:4.397 | שכתב |  |
| 7:4.398 | הרח"ו |  |
| 7:4.399 | והובא |  |
| 7:4.400 | לעיל |  |
| 7:4.401 | סי | ' |
| 7:4.402 | מ"ז | . |
| 7:4.403 | ועל |  |
| 7:4.404 | הטעם |  |
| 7:4.405 | הזה |  |
| 7:4.406 | יפלא |  |
| 7:4.407 | מאד |  |
| 7:4.408 | הלא |  |
| 7:4.409 | ה | ' |
| 7:4.410 | אלהינו |  |
| 7:4.411 | שהוא |  |
| 7:4.412 | הסבה |  |
| 7:4.413 | הראשונה |  |
| 7:4.414 | אשר |  |
| 7:4.415 | עבדו |  |
| 7:4.416 | אבותינו |  |
| 7:4.417 | ואבות |  |
| 7:4.418 | אבותינו |  |
| 7:4.419 | נעלם |  |
| 7:4.420 | הוא |  |
| 7:4.421 | מאד |  |
| 7:4.422 | עד |  |
| 7:4.423 | שאפי | ' |
| 7:4.424 | המלאכים |  |
| 7:4.425 | וחיות |  |
| 7:4.426 | הקדש |  |
| 7:4.427 | שואלים |  |
| 7:4.428 | איה |  |
| 7:4.429 | מקום |  |
| 7:4.430 | כבודו |  |
| 7:4.431 | כמו |  |
| 7:4.432 | שאמרו |  |
| 7:4.433 | רז"ל |  |
| 7:4.434 | ועם |  |
| 7:4.435 | כל |  |
| 7:4.436 | זה |  |
| 7:4.437 | אמר |  |
| 7:4.438 | הכתוב |  |
| 7:4.439 | כי |  |
| 7:4.440 | מי |  |
| 7:4.441 | גוי |  |
| 7:4.442 | גדול |  |
| 7:4.443 | אשר |  |
| 7:4.444 | לו |  |
| 7:4.445 | אלהים |  |
| 7:4.446 | קרובים |  |
| 7:4.447 | אליו |  |
| 7:4.448 | כה | ' |
| 7:4.449 | אלהינו |  |
| 7:4.450 | בכל |  |
| 7:4.451 | קראנו |  |
| 7:4.452 | אליו | . |
| 7:4.453 | כמפה |  |
| 7:4.454 | לאזן |  |
| 7:4.455 | כמו |  |
| 7:4.456 | שנת | ' |
| 7:4.457 | לעיל | : |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

