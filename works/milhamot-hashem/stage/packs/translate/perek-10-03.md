# Translation pack: milhamot-hashem perek-10-03 (Milchamot ha-Shem 10:3–10:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-03`.

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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 10:3

Hebrew (Wikisource): אוי! כי על ידי הספר המתעה הזהר והתקונים דמינו לעמים עובדי האלילים, (הודיים ומצרים וכשדים פרס ומדי ושאר אמונות הבנוים על למודים נסתרים וסודות זרים אשר יסודתן בכח המדמה כאשר יראה המעיין בפסר דברי עולם לחחכם קלמן סולמאן). אשר לא יצאו מתוך ההפכה, וריב להם עם הדעת הישרה המקובלת מפי סופרים וספרים, ספרי רז"ל, המורים לאדם דעת ותבונה בתורתנו הקדושה. בל נטעה ונשגה ברעיונות העולים על לב ריק, לעבור על מאמר השי"ת ולא תתורו אחרי לבבכם זו מינות! ונתקיים בנו על ידי הזהר והתקונים מאמר רז"ל בסנהדרין דף ל"ט ובילקוט מלכים רמז רכ"ו, ריב"ל רמי כתיב וכמשפטי הגוים אשר סביבותיכם עשיתם, וכתיב וכמשפטי הגוים אשר סביבותיכם לא עשיתם? ומשני כמתוקנים שבהם (המאמינים באחדות הבורא) לא עשיתם, כמקולקלים שבהם הנז' לעיל עשיתם! כן אירע לנו באמונת האחדות והעבודה, כמתוקנים שבהם לא עשינו כמקולקלים שבהם עשינו לעבוד אל קצא אפים (ז"א) וליחדו עם האלילה זוגתו.

| id | word | punctuation after |
|---|---|---|
| 10:3.1 | אוי | ! |
| 10:3.2 | כי |  |
| 10:3.3 | על |  |
| 10:3.4 | ידי |  |
| 10:3.5 | הספר |  |
| 10:3.6 | המתעה |  |
| 10:3.7 | הזהר |  |
| 10:3.8 | והתקונים |  |
| 10:3.9 | דמינו |  |
| 10:3.10 | לעמים |  |
| 10:3.11 | עובדי |  |
| 10:3.12 | האלילים | , |
| 10:3.13 | הודיים |  |
| 10:3.14 | ומצרים |  |
| 10:3.15 | וכשדים |  |
| 10:3.16 | פרס |  |
| 10:3.17 | ומדי |  |
| 10:3.18 | ושאר |  |
| 10:3.19 | אמונות |  |
| 10:3.20 | הבנוים |  |
| 10:3.21 | על |  |
| 10:3.22 | למודים |  |
| 10:3.23 | נסתרים |  |
| 10:3.24 | וסודות |  |
| 10:3.25 | זרים |  |
| 10:3.26 | אשר |  |
| 10:3.27 | יסודתן |  |
| 10:3.28 | בכח |  |
| 10:3.29 | המדמה |  |
| 10:3.30 | כאשר |  |
| 10:3.31 | יראה |  |
| 10:3.32 | המעיין |  |
| 10:3.33 | בפסר |  |
| 10:3.34 | דברי |  |
| 10:3.35 | עולם |  |
| 10:3.36 | לחחכם |  |
| 10:3.37 | קלמן |  |
| 10:3.38 | סולמאן | ). |
| 10:3.39 | אשר |  |
| 10:3.40 | לא |  |
| 10:3.41 | יצאו |  |
| 10:3.42 | מתוך |  |
| 10:3.43 | ההפכה | , |
| 10:3.44 | וריב |  |
| 10:3.45 | להם |  |
| 10:3.46 | עם |  |
| 10:3.47 | הדעת |  |
| 10:3.48 | הישרה |  |
| 10:3.49 | המקובלת |  |
| 10:3.50 | מפי |  |
| 10:3.51 | סופרים |  |
| 10:3.52 | וספרים | , |
| 10:3.53 | ספרי |  |
| 10:3.54 | רז"ל | , |
| 10:3.55 | המורים |  |
| 10:3.56 | לאדם |  |
| 10:3.57 | דעת |  |
| 10:3.58 | ותבונה |  |
| 10:3.59 | בתורתנו |  |
| 10:3.60 | הקדושה | . |
| 10:3.61 | בל |  |
| 10:3.62 | נטעה |  |
| 10:3.63 | ונשגה |  |
| 10:3.64 | ברעיונות |  |
| 10:3.65 | העולים |  |
| 10:3.66 | על |  |
| 10:3.67 | לב |  |
| 10:3.68 | ריק | , |
| 10:3.69 | לעבור |  |
| 10:3.70 | על |  |
| 10:3.71 | מאמר |  |
| 10:3.72 | השי"ת |  |
| 10:3.73 | ולא |  |
| 10:3.74 | תתורו |  |
| 10:3.75 | אחרי |  |
| 10:3.76 | לבבכם |  |
| 10:3.77 | זו |  |
| 10:3.78 | מינות | ! |
| 10:3.79 | ונתקיים |  |
| 10:3.80 | בנו |  |
| 10:3.81 | על |  |
| 10:3.82 | ידי |  |
| 10:3.83 | הזהר |  |
| 10:3.84 | והתקונים |  |
| 10:3.85 | מאמר |  |
| 10:3.86 | רז"ל |  |
| 10:3.87 | בסנהדרין |  |
| 10:3.88 | דף |  |
| 10:3.89 | ל"ט |  |
| 10:3.90 | ובילקוט |  |
| 10:3.91 | מלכים |  |
| 10:3.92 | רמז |  |
| 10:3.93 | רכ"ו | , |
| 10:3.94 | ריב"ל |  |
| 10:3.95 | רמי |  |
| 10:3.96 | כתיב |  |
| 10:3.97 | וכמשפטי |  |
| 10:3.98 | הגוים |  |
| 10:3.99 | אשר |  |
| 10:3.100 | סביבותיכם |  |
| 10:3.101 | עשיתם | , |
| 10:3.102 | וכתיב |  |
| 10:3.103 | וכמשפטי |  |
| 10:3.104 | הגוים |  |
| 10:3.105 | אשר |  |
| 10:3.106 | סביבותיכם |  |
| 10:3.107 | לא |  |
| 10:3.108 | עשיתם | ? |
| 10:3.109 | ומשני |  |
| 10:3.110 | כמתוקנים |  |
| 10:3.111 | שבהם |  |
| 10:3.112 | המאמינים |  |
| 10:3.113 | באחדות |  |
| 10:3.114 | הבורא | ) |
| 10:3.115 | לא |  |
| 10:3.116 | עשיתם | , |
| 10:3.117 | כמקולקלים |  |
| 10:3.118 | שבהם |  |
| 10:3.119 | הנז | ' |
| 10:3.120 | לעיל |  |
| 10:3.121 | עשיתם | ! |
| 10:3.122 | כן |  |
| 10:3.123 | אירע |  |
| 10:3.124 | לנו |  |
| 10:3.125 | באמונת |  |
| 10:3.126 | האחדות |  |
| 10:3.127 | והעבודה | , |
| 10:3.128 | כמתוקנים |  |
| 10:3.129 | שבהם |  |
| 10:3.130 | לא |  |
| 10:3.131 | עשינו |  |
| 10:3.132 | כמקולקלים |  |
| 10:3.133 | שבהם |  |
| 10:3.134 | עשינו |  |
| 10:3.135 | לעבוד |  |
| 10:3.136 | אל |  |
| 10:3.137 | קצא |  |
| 10:3.138 | אפים |  |
| 10:3.139 | ז"א | ) |
| 10:3.140 | וליחדו |  |
| 10:3.141 | עם |  |
| 10:3.142 | האלילה |  |
| 10:3.143 | זוגתו | . |


## Milchamot ha-Shem 10:4

Hebrew (Wikisource): קמב) מי יגלה עפר מעיניכם רבותינו אנשי כנסת הגדולה אשר שפכתם שיח בתחנה ובקשה וצעקה לפני אבינו שבשמים בקול גדול על יצרא דע"ז, להעביר גלולים ובעלים ועשתרות מעם ישראל ועלתה בידכם! אך לא לנצח כי קבלתכם האמתית המשנה והתלמוד הקדוש ופסקי הגאונים הבאים אחריכם שם אותה המתעה מחבר הזהר וההולכים בעקבותיו לשפחה וקליפה ומוץ ותבן וסלע אחרא מסטרא דיצרא בישא, ואת לימוד עבודת הבעלים והעשתרות והאשרות שהם פרצופי וצורות האלוהות הזכרים והנקבות הרבים שם לגבירה, ויעבדום בקדושה ובטהרת הגוף ורתת ופחד, עם טומאת השכל ובלבול המחשבה עד שלא יוכל להבין מה שמוציא מפיו כאשר יראה המעיין בספרי הכוונות הרבים השונים. ועליהם אמר דוד המלך ע"ה (לפי פי' דר"ר עובדיה ספורנו והרב זמרת יה) וז"ל הרב ספורנו במזמור י"ו שמרני אל כי חסיתי בך, ולא בהבלי החושבים לפעול פלאות בשמות מלאכים וזולתם (החמשה פרצופים או הי"ב פרצופים) אמרת לה, את נפשי המשכלת, ה' אתה, אתה אדון לעמך בלי אמצעי טובתי, את נפשי המשכלת, טובה, שלי, בל עליך, אין שום דבר למעלה ממך! לקדושים אשר בארץ המה, ואמרת גם כן לאותם אשר בארץ בחזקת קדושים, ואדירי כל חפצי בם שהם אדירים גם כן. ונחשבים בישראל אשר כל חפצי בם ירבו עצבותם, אמרת שירבו עצבותם וצרותם, כי אמנם בתכליתם. אחר מהרו מיהרו לפנות אל אל אחד מבלעדי אלהים אבל אל מלאכים וזולתם (ירמוז אל פרצופי האצילות) בל אסיך נסכיהם מדם לא אקרב לעשות הנסכים שהם עושים להוריד הכתות, יותר ממה שהייתי מנסך את הדם שהוא מתועב לנסכים ובל אשא את שמותם אותם שמות שהם חושבים לפעול בהם פלאות לא אשא על שפתי ולא אזכירם כי אמנם אתה ה' לבדך מנת חלקי בבחורתי, וכוסי בירושת אבות, אתה תומיך גורלי אתה תתמכני בגורל שעלה לי להיות במספר עבדיך ותסמכני שלא אטה ממנו, כי אמנם חבלים נפלו לי בנעימים, הגיעו לחלקח במושכלות שהם אותם נעימים שהשגתי אף נחלת קבלת אבות שפרה עלי שמצאתיה מסכמת לאמת המושכל ובכן אברך את ה' אשר יעצני ללכת בדרכיו וכו' ועיין זמרת יה שהאריך עוד יותר מזה. גם במזמור כ"ו היה דוד המלך ע"ה, מתחנן ומתפלל לשי"ת לאמר שפטני ה' היודע כליות ולב, כי אני בתמי הלכתי ולא נטיתי מעיקרי האמונה, ובה' לבדו בטחתי בלי שום שתוף כח אחר מכל הנבראים גשמי או רוחני, בחנני ה' ונסיני וצרפה כליותי ולבי שהם כלי המחשבה אם אטה מני הדרך עיקרי הדת, כי חסדך לנגד עיני ששמעת תפלתי בבקשי מלפניך הורני ה' דרכך אהלך באמתיך, ועשית בקשתי ונתת לי נפש משכלת, ועל ידי זה התהלכתי באמתיך ולא ישבתי עם מתי שוא התרים אחר לבם להמציא ברעיונותם סבות אחרות וכחות ופרצופים אלהיים נעלמים מבלעדיך ליחדם ולקשרם ולתקנם להיות אחד אע"פ שהם רבים במספר, ולהיותם אמצעיים למשוך בכחם השפע מהאלהים הגנוז בכתר, ולא מהאלהים העליון המקיף את כל החלל של כל העולמות (לפי שעתם) וזהו שאמר ועם נעלמים הדורשים מכוחות נעלמים מבלעדי ה' לא אבוא, כי שנאתי קהל מרעים המתהללים באלילים ועם אלה רשעים לא אשב. אבל ארחץ בנקיון כפי מכל אמונות זרות וכוזבות ואסובבה את מזבחך ה'.

| id | word | punctuation after |
|---|---|---|
| 10:4.1 | קמב | ) |
| 10:4.2 | מי |  |
| 10:4.3 | יגלה |  |
| 10:4.4 | עפר |  |
| 10:4.5 | מעיניכם |  |
| 10:4.6 | רבותינו |  |
| 10:4.7 | אנשי |  |
| 10:4.8 | כנסת |  |
| 10:4.9 | הגדולה |  |
| 10:4.10 | אשר |  |
| 10:4.11 | שפכתם |  |
| 10:4.12 | שיח |  |
| 10:4.13 | בתחנה |  |
| 10:4.14 | ובקשה |  |
| 10:4.15 | וצעקה |  |
| 10:4.16 | לפני |  |
| 10:4.17 | אבינו |  |
| 10:4.18 | שבשמים |  |
| 10:4.19 | בקול |  |
| 10:4.20 | גדול |  |
| 10:4.21 | על |  |
| 10:4.22 | יצרא |  |
| 10:4.23 | דע"ז | , |
| 10:4.24 | להעביר |  |
| 10:4.25 | גלולים |  |
| 10:4.26 | ובעלים |  |
| 10:4.27 | ועשתרות |  |
| 10:4.28 | מעם |  |
| 10:4.29 | ישראל |  |
| 10:4.30 | ועלתה |  |
| 10:4.31 | בידכם | ! |
| 10:4.32 | אך |  |
| 10:4.33 | לא |  |
| 10:4.34 | לנצח |  |
| 10:4.35 | כי |  |
| 10:4.36 | קבלתכם |  |
| 10:4.37 | האמתית |  |
| 10:4.38 | המשנה |  |
| 10:4.39 | והתלמוד |  |
| 10:4.40 | הקדוש |  |
| 10:4.41 | ופסקי |  |
| 10:4.42 | הגאונים |  |
| 10:4.43 | הבאים |  |
| 10:4.44 | אחריכם |  |
| 10:4.45 | שם |  |
| 10:4.46 | אותה |  |
| 10:4.47 | המתעה |  |
| 10:4.48 | מחבר |  |
| 10:4.49 | הזהר |  |
| 10:4.50 | וההולכים |  |
| 10:4.51 | בעקבותיו |  |
| 10:4.52 | לשפחה |  |
| 10:4.53 | וקליפה |  |
| 10:4.54 | ומוץ |  |
| 10:4.55 | ותבן |  |
| 10:4.56 | וסלע |  |
| 10:4.57 | אחרא |  |
| 10:4.58 | מסטרא |  |
| 10:4.59 | דיצרא |  |
| 10:4.60 | בישא | , |
| 10:4.61 | ואת |  |
| 10:4.62 | לימוד |  |
| 10:4.63 | עבודת |  |
| 10:4.64 | הבעלים |  |
| 10:4.65 | והעשתרות |  |
| 10:4.66 | והאשרות |  |
| 10:4.67 | שהם |  |
| 10:4.68 | פרצופי |  |
| 10:4.69 | וצורות |  |
| 10:4.70 | האלוהות |  |
| 10:4.71 | הזכרים |  |
| 10:4.72 | והנקבות |  |
| 10:4.73 | הרבים |  |
| 10:4.74 | שם |  |
| 10:4.75 | לגבירה | , |
| 10:4.76 | ויעבדום |  |
| 10:4.77 | בקדושה |  |
| 10:4.78 | ובטהרת |  |
| 10:4.79 | הגוף |  |
| 10:4.80 | ורתת |  |
| 10:4.81 | ופחד | , |
| 10:4.82 | עם |  |
| 10:4.83 | טומאת |  |
| 10:4.84 | השכל |  |
| 10:4.85 | ובלבול |  |
| 10:4.86 | המחשבה |  |
| 10:4.87 | עד |  |
| 10:4.88 | שלא |  |
| 10:4.89 | יוכל |  |
| 10:4.90 | להבין |  |
| 10:4.91 | מה |  |
| 10:4.92 | שמוציא |  |
| 10:4.93 | מפיו |  |
| 10:4.94 | כאשר |  |
| 10:4.95 | יראה |  |
| 10:4.96 | המעיין |  |
| 10:4.97 | בספרי |  |
| 10:4.98 | הכוונות |  |
| 10:4.99 | הרבים |  |
| 10:4.100 | השונים | . |
| 10:4.101 | ועליהם |  |
| 10:4.102 | אמר |  |
| 10:4.103 | דוד |  |
| 10:4.104 | המלך |  |
| 10:4.105 | ע"ה |  |
| 10:4.106 | לפי |  |
| 10:4.107 | פי | ' |
| 10:4.108 | דר"ר |  |
| 10:4.109 | עובדיה |  |
| 10:4.110 | ספורנו |  |
| 10:4.111 | והרב |  |
| 10:4.112 | זמרת |  |
| 10:4.113 | יה | ) |
| 10:4.114 | וז"ל |  |
| 10:4.115 | הרב |  |
| 10:4.116 | ספורנו |  |
| 10:4.117 | במזמור |  |
| 10:4.118 | י"ו |  |
| 10:4.119 | שמרני |  |
| 10:4.120 | אל |  |
| 10:4.121 | כי |  |
| 10:4.122 | חסיתי |  |
| 10:4.123 | בך | , |
| 10:4.124 | ולא |  |
| 10:4.125 | בהבלי |  |
| 10:4.126 | החושבים |  |
| 10:4.127 | לפעול |  |
| 10:4.128 | פלאות |  |
| 10:4.129 | בשמות |  |
| 10:4.130 | מלאכים |  |
| 10:4.131 | וזולתם |  |
| 10:4.132 | החמשה |  |
| 10:4.133 | פרצופים |  |
| 10:4.134 | או |  |
| 10:4.135 | הי"ב |  |
| 10:4.136 | פרצופים | ) |
| 10:4.137 | אמרת |  |
| 10:4.138 | לה | , |
| 10:4.139 | את |  |
| 10:4.140 | נפשי |  |
| 10:4.141 | המשכלת | , |
| 10:4.142 | ה | ' |
| 10:4.143 | אתה | , |
| 10:4.144 | אתה |  |
| 10:4.145 | אדון |  |
| 10:4.146 | לעמך |  |
| 10:4.147 | בלי |  |
| 10:4.148 | אמצעי |  |
| 10:4.149 | טובתי | , |
| 10:4.150 | את |  |
| 10:4.151 | נפשי |  |
| 10:4.152 | המשכלת | , |
| 10:4.153 | טובה | , |
| 10:4.154 | שלי | , |
| 10:4.155 | בל |  |
| 10:4.156 | עליך | , |
| 10:4.157 | אין |  |
| 10:4.158 | שום |  |
| 10:4.159 | דבר |  |
| 10:4.160 | למעלה |  |
| 10:4.161 | ממך | ! |
| 10:4.162 | לקדושים |  |
| 10:4.163 | אשר |  |
| 10:4.164 | בארץ |  |
| 10:4.165 | המה | , |
| 10:4.166 | ואמרת |  |
| 10:4.167 | גם |  |
| 10:4.168 | כן |  |
| 10:4.169 | לאותם |  |
| 10:4.170 | אשר |  |
| 10:4.171 | בארץ |  |
| 10:4.172 | בחזקת |  |
| 10:4.173 | קדושים | , |
| 10:4.174 | ואדירי |  |
| 10:4.175 | כל |  |
| 10:4.176 | חפצי |  |
| 10:4.177 | בם |  |
| 10:4.178 | שהם |  |
| 10:4.179 | אדירים |  |
| 10:4.180 | גם |  |
| 10:4.181 | כן | . |
| 10:4.182 | ונחשבים |  |
| 10:4.183 | בישראל |  |
| 10:4.184 | אשר |  |
| 10:4.185 | כל |  |
| 10:4.186 | חפצי |  |
| 10:4.187 | בם |  |
| 10:4.188 | ירבו |  |
| 10:4.189 | עצבותם | , |
| 10:4.190 | אמרת |  |
| 10:4.191 | שירבו |  |
| 10:4.192 | עצבותם |  |
| 10:4.193 | וצרותם | , |
| 10:4.194 | כי |  |
| 10:4.195 | אמנם |  |
| 10:4.196 | בתכליתם | . |
| 10:4.197 | אחר |  |
| 10:4.198 | מהרו |  |
| 10:4.199 | מיהרו |  |
| 10:4.200 | לפנות |  |
| 10:4.201 | אל |  |
| 10:4.202 | אל |  |
| 10:4.203 | אחד |  |
| 10:4.204 | מבלעדי |  |
| 10:4.205 | אלהים |  |
| 10:4.206 | אבל |  |
| 10:4.207 | אל |  |
| 10:4.208 | מלאכים |  |
| 10:4.209 | וזולתם |  |
| 10:4.210 | ירמוז |  |
| 10:4.211 | אל |  |
| 10:4.212 | פרצופי |  |
| 10:4.213 | האצילות | ) |
| 10:4.214 | בל |  |
| 10:4.215 | אסיך |  |
| 10:4.216 | נסכיהם |  |
| 10:4.217 | מדם |  |
| 10:4.218 | לא |  |
| 10:4.219 | אקרב |  |
| 10:4.220 | לעשות |  |
| 10:4.221 | הנסכים |  |
| 10:4.222 | שהם |  |
| 10:4.223 | עושים |  |
| 10:4.224 | להוריד |  |
| 10:4.225 | הכתות | , |
| 10:4.226 | יותר |  |
| 10:4.227 | ממה |  |
| 10:4.228 | שהייתי |  |
| 10:4.229 | מנסך |  |
| 10:4.230 | את |  |
| 10:4.231 | הדם |  |
| 10:4.232 | שהוא |  |
| 10:4.233 | מתועב |  |
| 10:4.234 | לנסכים |  |
| 10:4.235 | ובל |  |
| 10:4.236 | אשא |  |
| 10:4.237 | את |  |
| 10:4.238 | שמותם |  |
| 10:4.239 | אותם |  |
| 10:4.240 | שמות |  |
| 10:4.241 | שהם |  |
| 10:4.242 | חושבים |  |
| 10:4.243 | לפעול |  |
| 10:4.244 | בהם |  |
| 10:4.245 | פלאות |  |
| 10:4.246 | לא |  |
| 10:4.247 | אשא |  |
| 10:4.248 | על |  |
| 10:4.249 | שפתי |  |
| 10:4.250 | ולא |  |
| 10:4.251 | אזכירם |  |
| 10:4.252 | כי |  |
| 10:4.253 | אמנם |  |
| 10:4.254 | אתה |  |
| 10:4.255 | ה | ' |
| 10:4.256 | לבדך |  |
| 10:4.257 | מנת |  |
| 10:4.258 | חלקי |  |
| 10:4.259 | בבחורתי | , |
| 10:4.260 | וכוסי |  |
| 10:4.261 | בירושת |  |
| 10:4.262 | אבות | , |
| 10:4.263 | אתה |  |
| 10:4.264 | תומיך |  |
| 10:4.265 | גורלי |  |
| 10:4.266 | אתה |  |
| 10:4.267 | תתמכני |  |
| 10:4.268 | בגורל |  |
| 10:4.269 | שעלה |  |
| 10:4.270 | לי |  |
| 10:4.271 | להיות |  |
| 10:4.272 | במספר |  |
| 10:4.273 | עבדיך |  |
| 10:4.274 | ותסמכני |  |
| 10:4.275 | שלא |  |
| 10:4.276 | אטה |  |
| 10:4.277 | ממנו | , |
| 10:4.278 | כי |  |
| 10:4.279 | אמנם |  |
| 10:4.280 | חבלים |  |
| 10:4.281 | נפלו |  |
| 10:4.282 | לי |  |
| 10:4.283 | בנעימים | , |
| 10:4.284 | הגיעו |  |
| 10:4.285 | לחלקח |  |
| 10:4.286 | במושכלות |  |
| 10:4.287 | שהם |  |
| 10:4.288 | אותם |  |
| 10:4.289 | נעימים |  |
| 10:4.290 | שהשגתי |  |
| 10:4.291 | אף |  |
| 10:4.292 | נחלת |  |
| 10:4.293 | קבלת |  |
| 10:4.294 | אבות |  |
| 10:4.295 | שפרה |  |
| 10:4.296 | עלי |  |
| 10:4.297 | שמצאתיה |  |
| 10:4.298 | מסכמת |  |
| 10:4.299 | לאמת |  |
| 10:4.300 | המושכל |  |
| 10:4.301 | ובכן |  |
| 10:4.302 | אברך |  |
| 10:4.303 | את |  |
| 10:4.304 | ה | ' |
| 10:4.305 | אשר |  |
| 10:4.306 | יעצני |  |
| 10:4.307 | ללכת |  |
| 10:4.308 | בדרכיו |  |
| 10:4.309 | וכו | ' |
| 10:4.310 | ועיין |  |
| 10:4.311 | זמרת |  |
| 10:4.312 | יה |  |
| 10:4.313 | שהאריך |  |
| 10:4.314 | עוד |  |
| 10:4.315 | יותר |  |
| 10:4.316 | מזה | . |
| 10:4.317 | גם |  |
| 10:4.318 | במזמור |  |
| 10:4.319 | כ"ו |  |
| 10:4.320 | היה |  |
| 10:4.321 | דוד |  |
| 10:4.322 | המלך |  |
| 10:4.323 | ע"ה | , |
| 10:4.324 | מתחנן |  |
| 10:4.325 | ומתפלל |  |
| 10:4.326 | לשי"ת |  |
| 10:4.327 | לאמר |  |
| 10:4.328 | שפטני |  |
| 10:4.329 | ה | ' |
| 10:4.330 | היודע |  |
| 10:4.331 | כליות |  |
| 10:4.332 | ולב | , |
| 10:4.333 | כי |  |
| 10:4.334 | אני |  |
| 10:4.335 | בתמי |  |
| 10:4.336 | הלכתי |  |
| 10:4.337 | ולא |  |
| 10:4.338 | נטיתי |  |
| 10:4.339 | מעיקרי |  |
| 10:4.340 | האמונה | , |
| 10:4.341 | ובה | ' |
| 10:4.342 | לבדו |  |
| 10:4.343 | בטחתי |  |
| 10:4.344 | בלי |  |
| 10:4.345 | שום |  |
| 10:4.346 | שתוף |  |
| 10:4.347 | כח |  |
| 10:4.348 | אחר |  |
| 10:4.349 | מכל |  |
| 10:4.350 | הנבראים |  |
| 10:4.351 | גשמי |  |
| 10:4.352 | או |  |
| 10:4.353 | רוחני | , |
| 10:4.354 | בחנני |  |
| 10:4.355 | ה | ' |
| 10:4.356 | ונסיני |  |
| 10:4.357 | וצרפה |  |
| 10:4.358 | כליותי |  |
| 10:4.359 | ולבי |  |
| 10:4.360 | שהם |  |
| 10:4.361 | כלי |  |
| 10:4.362 | המחשבה |  |
| 10:4.363 | אם |  |
| 10:4.364 | אטה |  |
| 10:4.365 | מני |  |
| 10:4.366 | הדרך |  |
| 10:4.367 | עיקרי |  |
| 10:4.368 | הדת | , |
| 10:4.369 | כי |  |
| 10:4.370 | חסדך |  |
| 10:4.371 | לנגד |  |
| 10:4.372 | עיני |  |
| 10:4.373 | ששמעת |  |
| 10:4.374 | תפלתי |  |
| 10:4.375 | בבקשי |  |
| 10:4.376 | מלפניך |  |
| 10:4.377 | הורני |  |
| 10:4.378 | ה | ' |
| 10:4.379 | דרכך |  |
| 10:4.380 | אהלך |  |
| 10:4.381 | באמתיך | , |
| 10:4.382 | ועשית |  |
| 10:4.383 | בקשתי |  |
| 10:4.384 | ונתת |  |
| 10:4.385 | לי |  |
| 10:4.386 | נפש |  |
| 10:4.387 | משכלת | , |
| 10:4.388 | ועל |  |
| 10:4.389 | ידי |  |
| 10:4.390 | זה |  |
| 10:4.391 | התהלכתי |  |
| 10:4.392 | באמתיך |  |
| 10:4.393 | ולא |  |
| 10:4.394 | ישבתי |  |
| 10:4.395 | עם |  |
| 10:4.396 | מתי |  |
| 10:4.397 | שוא |  |
| 10:4.398 | התרים |  |
| 10:4.399 | אחר |  |
| 10:4.400 | לבם |  |
| 10:4.401 | להמציא |  |
| 10:4.402 | ברעיונותם |  |
| 10:4.403 | סבות |  |
| 10:4.404 | אחרות |  |
| 10:4.405 | וכחות |  |
| 10:4.406 | ופרצופים |  |
| 10:4.407 | אלהיים |  |
| 10:4.408 | נעלמים |  |
| 10:4.409 | מבלעדיך |  |
| 10:4.410 | ליחדם |  |
| 10:4.411 | ולקשרם |  |
| 10:4.412 | ולתקנם |  |
| 10:4.413 | להיות |  |
| 10:4.414 | אחד |  |
| 10:4.415 | אע"פ |  |
| 10:4.416 | שהם |  |
| 10:4.417 | רבים |  |
| 10:4.418 | במספר | , |
| 10:4.419 | ולהיותם |  |
| 10:4.420 | אמצעיים |  |
| 10:4.421 | למשוך |  |
| 10:4.422 | בכחם |  |
| 10:4.423 | השפע |  |
| 10:4.424 | מהאלהים |  |
| 10:4.425 | הגנוז |  |
| 10:4.426 | בכתר | , |
| 10:4.427 | ולא |  |
| 10:4.428 | מהאלהים |  |
| 10:4.429 | העליון |  |
| 10:4.430 | המקיף |  |
| 10:4.431 | את |  |
| 10:4.432 | כל |  |
| 10:4.433 | החלל |  |
| 10:4.434 | של |  |
| 10:4.435 | כל |  |
| 10:4.436 | העולמות |  |
| 10:4.437 | לפי |  |
| 10:4.438 | שעתם | ) |
| 10:4.439 | וזהו |  |
| 10:4.440 | שאמר |  |
| 10:4.441 | ועם |  |
| 10:4.442 | נעלמים |  |
| 10:4.443 | הדורשים |  |
| 10:4.444 | מכוחות |  |
| 10:4.445 | נעלמים |  |
| 10:4.446 | מבלעדי |  |
| 10:4.447 | ה | ' |
| 10:4.448 | לא |  |
| 10:4.449 | אבוא | , |
| 10:4.450 | כי |  |
| 10:4.451 | שנאתי |  |
| 10:4.452 | קהל |  |
| 10:4.453 | מרעים |  |
| 10:4.454 | המתהללים |  |
| 10:4.455 | באלילים |  |
| 10:4.456 | ועם |  |
| 10:4.457 | אלה |  |
| 10:4.458 | רשעים |  |
| 10:4.459 | לא |  |
| 10:4.460 | אשב | . |
| 10:4.461 | אבל |  |
| 10:4.462 | ארחץ |  |
| 10:4.463 | בנקיון |  |
| 10:4.464 | כפי |  |
| 10:4.465 | מכל |  |
| 10:4.466 | אמונות |  |
| 10:4.467 | זרות |  |
| 10:4.468 | וכוזבות |  |
| 10:4.469 | ואסובבה |  |
| 10:4.470 | את |  |
| 10:4.471 | מזבחך |  |
| 10:4.472 | ה | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

