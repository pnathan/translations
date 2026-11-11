# Translation pack: milhamot-hashem perek-8-05 (Milchamot ha-Shem 8:6–8:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-05`.

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
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
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


## Milchamot ha-Shem 8:6

Hebrew (Wikisource): ושוב אחר כן בגליון הנזכר העתיקו דבר הרב בעל ספר הברית המחזק את לבנו, ומאמץ את רוחינו בל נירא ונפחד מכבוד ה' ומהדר גאונו כי נגשימהו, כי הבקי בחכמת הניתוח, אורו מעליו יפשיט. ולא יחוס על כבוד קונו. ולא יעצרהו הגשם (אשר לדעת הראב"ד ובעל העיקרים מי שיאמין שהשי"ת בעל תואר כתואר םני המלך כגוף וגויה מפני שהוא טועה בעיונו אין לקרותו מין (עיקרים מאמר א' פ' ב') ואע"פ שאינו כן לפי האמת רק שהוא טועה בפי' הכתובים) כלומד הקבלה החדשה אע"פ שיצייר במחשבתו ובפיו ובספר יחוק כל משיגי הגוף ומקריו, וייחסם לשי"ת. הנה אחר כך ישוב ויפשיטם ויסירם. היו כלא היו, וחזר הענין לכמו שהיה, כלעומת שבא כן ילך. ולא הרויח כלום, רק הפסיד ועבר על הכתוב ועל דברי רז"ל ולא חס על כבוד קונו.

| id | word | punctuation after |
|---|---|---|
| 8:6.1 | ושוב |  |
| 8:6.2 | אחר |  |
| 8:6.3 | כן |  |
| 8:6.4 | בגליון |  |
| 8:6.5 | הנזכר |  |
| 8:6.6 | העתיקו |  |
| 8:6.7 | דבר |  |
| 8:6.8 | הרב |  |
| 8:6.9 | בעל |  |
| 8:6.10 | ספר |  |
| 8:6.11 | הברית |  |
| 8:6.12 | המחזק |  |
| 8:6.13 | את |  |
| 8:6.14 | לבנו | , |
| 8:6.15 | ומאמץ |  |
| 8:6.16 | את |  |
| 8:6.17 | רוחינו |  |
| 8:6.18 | בל |  |
| 8:6.19 | נירא |  |
| 8:6.20 | ונפחד |  |
| 8:6.21 | מכבוד |  |
| 8:6.22 | ה | ' |
| 8:6.23 | ומהדר |  |
| 8:6.24 | גאונו |  |
| 8:6.25 | כי |  |
| 8:6.26 | נגשימהו | , |
| 8:6.27 | כי |  |
| 8:6.28 | הבקי |  |
| 8:6.29 | בחכמת |  |
| 8:6.30 | הניתוח | , |
| 8:6.31 | אורו |  |
| 8:6.32 | מעליו |  |
| 8:6.33 | יפשיט | . |
| 8:6.34 | ולא |  |
| 8:6.35 | יחוס |  |
| 8:6.36 | על |  |
| 8:6.37 | כבוד |  |
| 8:6.38 | קונו | . |
| 8:6.39 | ולא |  |
| 8:6.40 | יעצרהו |  |
| 8:6.41 | הגשם |  |
| 8:6.42 | אשר |  |
| 8:6.43 | לדעת |  |
| 8:6.44 | הראב"ד |  |
| 8:6.45 | ובעל |  |
| 8:6.46 | העיקרים |  |
| 8:6.47 | מי |  |
| 8:6.48 | שיאמין |  |
| 8:6.49 | שהשי"ת |  |
| 8:6.50 | בעל |  |
| 8:6.51 | תואר |  |
| 8:6.52 | כתואר |  |
| 8:6.53 | םני |  |
| 8:6.54 | המלך |  |
| 8:6.55 | כגוף |  |
| 8:6.56 | וגויה |  |
| 8:6.57 | מפני |  |
| 8:6.58 | שהוא |  |
| 8:6.59 | טועה |  |
| 8:6.60 | בעיונו |  |
| 8:6.61 | אין |  |
| 8:6.62 | לקרותו |  |
| 8:6.63 | מין |  |
| 8:6.64 | עיקרים |  |
| 8:6.65 | מאמר |  |
| 8:6.66 | א | ' |
| 8:6.67 | פ | ' |
| 8:6.68 | ב | ') |
| 8:6.69 | ואע"פ |  |
| 8:6.70 | שאינו |  |
| 8:6.71 | כן |  |
| 8:6.72 | לפי |  |
| 8:6.73 | האמת |  |
| 8:6.74 | רק |  |
| 8:6.75 | שהוא |  |
| 8:6.76 | טועה |  |
| 8:6.77 | בפי | ' |
| 8:6.78 | הכתובים | ) |
| 8:6.79 | כלומד |  |
| 8:6.80 | הקבלה |  |
| 8:6.81 | החדשה |  |
| 8:6.82 | אע"פ |  |
| 8:6.83 | שיצייר |  |
| 8:6.84 | במחשבתו |  |
| 8:6.85 | ובפיו |  |
| 8:6.86 | ובספר |  |
| 8:6.87 | יחוק |  |
| 8:6.88 | כל |  |
| 8:6.89 | משיגי |  |
| 8:6.90 | הגוף |  |
| 8:6.91 | ומקריו | , |
| 8:6.92 | וייחסם |  |
| 8:6.93 | לשי"ת | . |
| 8:6.94 | הנה |  |
| 8:6.95 | אחר |  |
| 8:6.96 | כך |  |
| 8:6.97 | ישוב |  |
| 8:6.98 | ויפשיטם |  |
| 8:6.99 | ויסירם | . |
| 8:6.100 | היו |  |
| 8:6.101 | כלא |  |
| 8:6.102 | היו | , |
| 8:6.103 | וחזר |  |
| 8:6.104 | הענין |  |
| 8:6.105 | לכמו |  |
| 8:6.106 | שהיה | , |
| 8:6.107 | כלעומת |  |
| 8:6.108 | שבא |  |
| 8:6.109 | כן |  |
| 8:6.110 | ילך | . |
| 8:6.111 | ולא |  |
| 8:6.112 | הרויח |  |
| 8:6.113 | כלום | , |
| 8:6.114 | רק |  |
| 8:6.115 | הפסיד |  |
| 8:6.116 | ועבר |  |
| 8:6.117 | על |  |
| 8:6.118 | הכתוב |  |
| 8:6.119 | ועל |  |
| 8:6.120 | דברי |  |
| 8:6.121 | רז"ל |  |
| 8:6.122 | ולא |  |
| 8:6.123 | חס |  |
| 8:6.124 | על |  |
| 8:6.125 | כבוד |  |
| 8:6.126 | קונו | . |


## Milchamot ha-Shem 8:7

Hebrew (Wikisource): קג) ולא אמרו הנז' שהמאמין שהוא ית' גוף וגויא וכו' אינו אלא במי שהוא מאמין וחושב בדעתו שהשי"ת שהוא הסבא הראשונה הוא גוף ובעל תמונה, ברם דא עקא וחולי שיש בו סכנה ואין לו רפואה כי הוא מיחס העבודה לסבות נבראות מחודשות, הם האלהים הקדושים שהם החמשה פרצופים, ושאין לעבוד לסבה הראשונה רק לפרצוף קצר האפים (זעיר ונוקביה) ולכלול עמו גם שאר פרצופים שלמעלה ממנו, אבל העובד לסבה הראשונה בבחינתו הפשוטה (שאין לו דמות הגוף) טועה הוא ואין תפלתו תפלה, כי לא תאות העבודה רק לחלק היושר המושג (בציור הגוף). אבל העגולים עם האין סוף המקיף, אין להם שום אלהות כלל, וכל עיקר האלהות הם פרצופי היושר, והם הברואים, ועשרה מאמרות שבהם נברא העולם נתחלקו ביניהם, בראשית שהוא מאמר ראשון נתיחס לעתיק ואריך עם אבא ואימא, כמו שנתבאר לעיל סי' כ"ד דלפי' קמא בכח אבא אריך את אלהים שהיא אימא, ולפירוש בתרא בכה אבא ברא אלהים שהיא בינה את האלוהות שאחריה המכונים בשם שמים וארץ, שהם קצר אפים וזוגתו (זעיר אנפין ונוקביה) אשר העבודה שייכא בהם דווקא. ושאר הח' מאמרות נתיחסו לאבא לבדו! ומאמר עשירי שהוא נעשה אדם, נתיחס לאימא. אבל זעיר ונוקביה אשר אותם ונעבוד לפי דבריהם לא נשתתפו במעשה בראשית בשום מאמר מעשרה מאמרות כאשר יראה המעיין בספר הזהר ומפרשיו. וכבר ביארתי לך לעיל סי' כ"ה איך הרחיקו רז"ל מלחשוב כי שמים וארץ אלוהות הן. והנה הבטחתו של הרב ספר הברית שלא נירא מהגשם, לא הועילה להצילנו מאמונת המון אלוהות רבים הנעזרים זה נזה, כי לא יכול לעשות והוא נצרך לכח אחר, וגם מיעץ אותנו לעבוד את קצר אפים (זעיר אנפין) הקדוש, והוא קורא אותו אלהא רברבא, והוא בן בן בן בנו של האלהים העליון שאמר אני ראשון ואני אחרון ומבלעדי אין אלהים. אני ראשון שאין לי אב, ואני אחרון שאין לי בן. ומבלעדי אין אלהים שאין לי אח, ושעל זה האלהים העליון אמרו בזהר (בחר דף ק"ט) ואנת לית לך גופא. ולית לך איברים ולית לך נוקבא אלא חד בלא שינוי וכו' כמבואר שם ב' רב פעלים (והרב עז לאלהים פרק כ"ה כתב על מאמר רז"ל על פסוק יש אחד ואין שני גם בן ואח אין לו, אבל אחות יש לו כו' ע"ש) וכללות דבריהם כי כל העניינים המורים על האחדות הגמורה בכתובים ובדברי רז"ל לא נאמרו רק על האלהים העליון אשר לא נעבוד אותו ולא נקרא אליו. כי לא יענה ח"ו. כמאמר הנביא מוכיח לישראל הקצור קצרה ידי מפדות ואם אין בי כח הציל וכו'. וכביכול לפי דבריהם הוא משתבח בעצמו בלבד ואינו עושה למבקש ממנו כלום. אבל עיקר העבודה והתפלה היא לקצר אפים. שהוא בגוף ואיברים ובעל אשה הוא. והוא אשר יענה לכל קוראיו ברצון אביו ואמו אשר גידלו אותו ורוממוהו ואעטרו ליה בתרין עטרין עד אשר לו נתכנו עלילות! אבל לעתיקא קדישא לא נתכנו עלילות, והזעיר אנפין הוא אשר נקרא בשם הוי'ה ובשם הקב"ה. ואליו תכון עבודתינו ותפלותינו לפי דבריהם אבל האלוה העליון שהוא יחיד בלי גוף ואיברים ואין לו אשה, לא נקרא אליו ולא נתפלל, והעובד אותו ומתפלל אליו אין תפלתו תפלה ולא יענה אותו. ואדרבא ענוש יענש.

| id | word | punctuation after |
|---|---|---|
| 8:7.1 | קג | ) |
| 8:7.2 | ולא |  |
| 8:7.3 | אמרו |  |
| 8:7.4 | הנז | ' |
| 8:7.5 | שהמאמין |  |
| 8:7.6 | שהוא |  |
| 8:7.7 | ית | ' |
| 8:7.8 | גוף |  |
| 8:7.9 | וגויא |  |
| 8:7.10 | וכו | ' |
| 8:7.11 | אינו |  |
| 8:7.12 | אלא |  |
| 8:7.13 | במי |  |
| 8:7.14 | שהוא |  |
| 8:7.15 | מאמין |  |
| 8:7.16 | וחושב |  |
| 8:7.17 | בדעתו |  |
| 8:7.18 | שהשי"ת |  |
| 8:7.19 | שהוא |  |
| 8:7.20 | הסבא |  |
| 8:7.21 | הראשונה |  |
| 8:7.22 | הוא |  |
| 8:7.23 | גוף |  |
| 8:7.24 | ובעל |  |
| 8:7.25 | תמונה | , |
| 8:7.26 | ברם |  |
| 8:7.27 | דא |  |
| 8:7.28 | עקא |  |
| 8:7.29 | וחולי |  |
| 8:7.30 | שיש |  |
| 8:7.31 | בו |  |
| 8:7.32 | סכנה |  |
| 8:7.33 | ואין |  |
| 8:7.34 | לו |  |
| 8:7.35 | רפואה |  |
| 8:7.36 | כי |  |
| 8:7.37 | הוא |  |
| 8:7.38 | מיחס |  |
| 8:7.39 | העבודה |  |
| 8:7.40 | לסבות |  |
| 8:7.41 | נבראות |  |
| 8:7.42 | מחודשות | , |
| 8:7.43 | הם |  |
| 8:7.44 | האלהים |  |
| 8:7.45 | הקדושים |  |
| 8:7.46 | שהם |  |
| 8:7.47 | החמשה |  |
| 8:7.48 | פרצופים | , |
| 8:7.49 | ושאין |  |
| 8:7.50 | לעבוד |  |
| 8:7.51 | לסבה |  |
| 8:7.52 | הראשונה |  |
| 8:7.53 | רק |  |
| 8:7.54 | לפרצוף |  |
| 8:7.55 | קצר |  |
| 8:7.56 | האפים |  |
| 8:7.57 | זעיר |  |
| 8:7.58 | ונוקביה | ) |
| 8:7.59 | ולכלול |  |
| 8:7.60 | עמו |  |
| 8:7.61 | גם |  |
| 8:7.62 | שאר |  |
| 8:7.63 | פרצופים |  |
| 8:7.64 | שלמעלה |  |
| 8:7.65 | ממנו | , |
| 8:7.66 | אבל |  |
| 8:7.67 | העובד |  |
| 8:7.68 | לסבה |  |
| 8:7.69 | הראשונה |  |
| 8:7.70 | בבחינתו |  |
| 8:7.71 | הפשוטה |  |
| 8:7.72 | שאין |  |
| 8:7.73 | לו |  |
| 8:7.74 | דמות |  |
| 8:7.75 | הגוף | ) |
| 8:7.76 | טועה |  |
| 8:7.77 | הוא |  |
| 8:7.78 | ואין |  |
| 8:7.79 | תפלתו |  |
| 8:7.80 | תפלה | , |
| 8:7.81 | כי |  |
| 8:7.82 | לא |  |
| 8:7.83 | תאות |  |
| 8:7.84 | העבודה |  |
| 8:7.85 | רק |  |
| 8:7.86 | לחלק |  |
| 8:7.87 | היושר |  |
| 8:7.88 | המושג |  |
| 8:7.89 | בציור |  |
| 8:7.90 | הגוף | ). |
| 8:7.91 | אבל |  |
| 8:7.92 | העגולים |  |
| 8:7.93 | עם |  |
| 8:7.94 | האין |  |
| 8:7.95 | סוף |  |
| 8:7.96 | המקיף | , |
| 8:7.97 | אין |  |
| 8:7.98 | להם |  |
| 8:7.99 | שום |  |
| 8:7.100 | אלהות |  |
| 8:7.101 | כלל | , |
| 8:7.102 | וכל |  |
| 8:7.103 | עיקר |  |
| 8:7.104 | האלהות |  |
| 8:7.105 | הם |  |
| 8:7.106 | פרצופי |  |
| 8:7.107 | היושר | , |
| 8:7.108 | והם |  |
| 8:7.109 | הברואים | , |
| 8:7.110 | ועשרה |  |
| 8:7.111 | מאמרות |  |
| 8:7.112 | שבהם |  |
| 8:7.113 | נברא |  |
| 8:7.114 | העולם |  |
| 8:7.115 | נתחלקו |  |
| 8:7.116 | ביניהם | , |
| 8:7.117 | בראשית |  |
| 8:7.118 | שהוא |  |
| 8:7.119 | מאמר |  |
| 8:7.120 | ראשון |  |
| 8:7.121 | נתיחס |  |
| 8:7.122 | לעתיק |  |
| 8:7.123 | ואריך |  |
| 8:7.124 | עם |  |
| 8:7.125 | אבא |  |
| 8:7.126 | ואימא | , |
| 8:7.127 | כמו |  |
| 8:7.128 | שנתבאר |  |
| 8:7.129 | לעיל |  |
| 8:7.130 | סי | ' |
| 8:7.131 | כ"ד |  |
| 8:7.132 | דלפי | ' |
| 8:7.133 | קמא |  |
| 8:7.134 | בכח |  |
| 8:7.135 | אבא |  |
| 8:7.136 | אריך |  |
| 8:7.137 | את |  |
| 8:7.138 | אלהים |  |
| 8:7.139 | שהיא |  |
| 8:7.140 | אימא | , |
| 8:7.141 | ולפירוש |  |
| 8:7.142 | בתרא |  |
| 8:7.143 | בכה |  |
| 8:7.144 | אבא |  |
| 8:7.145 | ברא |  |
| 8:7.146 | אלהים |  |
| 8:7.147 | שהיא |  |
| 8:7.148 | בינה |  |
| 8:7.149 | את |  |
| 8:7.150 | האלוהות |  |
| 8:7.151 | שאחריה |  |
| 8:7.152 | המכונים |  |
| 8:7.153 | בשם |  |
| 8:7.154 | שמים |  |
| 8:7.155 | וארץ | , |
| 8:7.156 | שהם |  |
| 8:7.157 | קצר |  |
| 8:7.158 | אפים |  |
| 8:7.159 | וזוגתו |  |
| 8:7.160 | זעיר |  |
| 8:7.161 | אנפין |  |
| 8:7.162 | ונוקביה | ) |
| 8:7.163 | אשר |  |
| 8:7.164 | העבודה |  |
| 8:7.165 | שייכא |  |
| 8:7.166 | בהם |  |
| 8:7.167 | דווקא | . |
| 8:7.168 | ושאר |  |
| 8:7.169 | הח | ' |
| 8:7.170 | מאמרות |  |
| 8:7.171 | נתיחסו |  |
| 8:7.172 | לאבא |  |
| 8:7.173 | לבדו | ! |
| 8:7.174 | ומאמר |  |
| 8:7.175 | עשירי |  |
| 8:7.176 | שהוא |  |
| 8:7.177 | נעשה |  |
| 8:7.178 | אדם | , |
| 8:7.179 | נתיחס |  |
| 8:7.180 | לאימא | . |
| 8:7.181 | אבל |  |
| 8:7.182 | זעיר |  |
| 8:7.183 | ונוקביה |  |
| 8:7.184 | אשר |  |
| 8:7.185 | אותם |  |
| 8:7.186 | ונעבוד |  |
| 8:7.187 | לפי |  |
| 8:7.188 | דבריהם |  |
| 8:7.189 | לא |  |
| 8:7.190 | נשתתפו |  |
| 8:7.191 | במעשה |  |
| 8:7.192 | בראשית |  |
| 8:7.193 | בשום |  |
| 8:7.194 | מאמר |  |
| 8:7.195 | מעשרה |  |
| 8:7.196 | מאמרות |  |
| 8:7.197 | כאשר |  |
| 8:7.198 | יראה |  |
| 8:7.199 | המעיין |  |
| 8:7.200 | בספר |  |
| 8:7.201 | הזהר |  |
| 8:7.202 | ומפרשיו | . |
| 8:7.203 | וכבר |  |
| 8:7.204 | ביארתי |  |
| 8:7.205 | לך |  |
| 8:7.206 | לעיל |  |
| 8:7.207 | סי | ' |
| 8:7.208 | כ"ה |  |
| 8:7.209 | איך |  |
| 8:7.210 | הרחיקו |  |
| 8:7.211 | רז"ל |  |
| 8:7.212 | מלחשוב |  |
| 8:7.213 | כי |  |
| 8:7.214 | שמים |  |
| 8:7.215 | וארץ |  |
| 8:7.216 | אלוהות |  |
| 8:7.217 | הן | . |
| 8:7.218 | והנה |  |
| 8:7.219 | הבטחתו |  |
| 8:7.220 | של |  |
| 8:7.221 | הרב |  |
| 8:7.222 | ספר |  |
| 8:7.223 | הברית |  |
| 8:7.224 | שלא |  |
| 8:7.225 | נירא |  |
| 8:7.226 | מהגשם | , |
| 8:7.227 | לא |  |
| 8:7.228 | הועילה |  |
| 8:7.229 | להצילנו |  |
| 8:7.230 | מאמונת |  |
| 8:7.231 | המון |  |
| 8:7.232 | אלוהות |  |
| 8:7.233 | רבים |  |
| 8:7.234 | הנעזרים |  |
| 8:7.235 | זה |  |
| 8:7.236 | נזה | , |
| 8:7.237 | כי |  |
| 8:7.238 | לא |  |
| 8:7.239 | יכול |  |
| 8:7.240 | לעשות |  |
| 8:7.241 | והוא |  |
| 8:7.242 | נצרך |  |
| 8:7.243 | לכח |  |
| 8:7.244 | אחר | , |
| 8:7.245 | וגם |  |
| 8:7.246 | מיעץ |  |
| 8:7.247 | אותנו |  |
| 8:7.248 | לעבוד |  |
| 8:7.249 | את |  |
| 8:7.250 | קצר |  |
| 8:7.251 | אפים |  |
| 8:7.252 | זעיר |  |
| 8:7.253 | אנפין | ) |
| 8:7.254 | הקדוש | , |
| 8:7.255 | והוא |  |
| 8:7.256 | קורא |  |
| 8:7.257 | אותו |  |
| 8:7.258 | אלהא |  |
| 8:7.259 | רברבא | , |
| 8:7.260 | והוא |  |
| 8:7.261 | בן |  |
| 8:7.262 | בן |  |
| 8:7.263 | בן |  |
| 8:7.264 | בנו |  |
| 8:7.265 | של |  |
| 8:7.266 | האלהים |  |
| 8:7.267 | העליון |  |
| 8:7.268 | שאמר |  |
| 8:7.269 | אני |  |
| 8:7.270 | ראשון |  |
| 8:7.271 | ואני |  |
| 8:7.272 | אחרון |  |
| 8:7.273 | ומבלעדי |  |
| 8:7.274 | אין |  |
| 8:7.275 | אלהים | . |
| 8:7.276 | אני |  |
| 8:7.277 | ראשון |  |
| 8:7.278 | שאין |  |
| 8:7.279 | לי |  |
| 8:7.280 | אב | , |
| 8:7.281 | ואני |  |
| 8:7.282 | אחרון |  |
| 8:7.283 | שאין |  |
| 8:7.284 | לי |  |
| 8:7.285 | בן | . |
| 8:7.286 | ומבלעדי |  |
| 8:7.287 | אין |  |
| 8:7.288 | אלהים |  |
| 8:7.289 | שאין |  |
| 8:7.290 | לי |  |
| 8:7.291 | אח | , |
| 8:7.292 | ושעל |  |
| 8:7.293 | זה |  |
| 8:7.294 | האלהים |  |
| 8:7.295 | העליון |  |
| 8:7.296 | אמרו |  |
| 8:7.297 | בזהר |  |
| 8:7.298 | בחר |  |
| 8:7.299 | דף |  |
| 8:7.300 | ק"ט | ) |
| 8:7.301 | ואנת |  |
| 8:7.302 | לית |  |
| 8:7.303 | לך |  |
| 8:7.304 | גופא | . |
| 8:7.305 | ולית |  |
| 8:7.306 | לך |  |
| 8:7.307 | איברים |  |
| 8:7.308 | ולית |  |
| 8:7.309 | לך |  |
| 8:7.310 | נוקבא |  |
| 8:7.311 | אלא |  |
| 8:7.312 | חד |  |
| 8:7.313 | בלא |  |
| 8:7.314 | שינוי |  |
| 8:7.315 | וכו | ' |
| 8:7.316 | כמבואר |  |
| 8:7.317 | שם |  |
| 8:7.318 | ב | ' |
| 8:7.319 | רב |  |
| 8:7.320 | פעלים |  |
| 8:7.321 | והרב |  |
| 8:7.322 | עז |  |
| 8:7.323 | לאלהים |  |
| 8:7.324 | פרק |  |
| 8:7.325 | כ"ה |  |
| 8:7.326 | כתב |  |
| 8:7.327 | על |  |
| 8:7.328 | מאמר |  |
| 8:7.329 | רז"ל |  |
| 8:7.330 | על |  |
| 8:7.331 | פסוק |  |
| 8:7.332 | יש |  |
| 8:7.333 | אחד |  |
| 8:7.334 | ואין |  |
| 8:7.335 | שני |  |
| 8:7.336 | גם |  |
| 8:7.337 | בן |  |
| 8:7.338 | ואח |  |
| 8:7.339 | אין |  |
| 8:7.340 | לו | , |
| 8:7.341 | אבל |  |
| 8:7.342 | אחות |  |
| 8:7.343 | יש |  |
| 8:7.344 | לו |  |
| 8:7.345 | כו | ' |
| 8:7.346 | ע"ש | ) |
| 8:7.347 | וכללות |  |
| 8:7.348 | דבריהם |  |
| 8:7.349 | כי |  |
| 8:7.350 | כל |  |
| 8:7.351 | העניינים |  |
| 8:7.352 | המורים |  |
| 8:7.353 | על |  |
| 8:7.354 | האחדות |  |
| 8:7.355 | הגמורה |  |
| 8:7.356 | בכתובים |  |
| 8:7.357 | ובדברי |  |
| 8:7.358 | רז"ל |  |
| 8:7.359 | לא |  |
| 8:7.360 | נאמרו |  |
| 8:7.361 | רק |  |
| 8:7.362 | על |  |
| 8:7.363 | האלהים |  |
| 8:7.364 | העליון |  |
| 8:7.365 | אשר |  |
| 8:7.366 | לא |  |
| 8:7.367 | נעבוד |  |
| 8:7.368 | אותו |  |
| 8:7.369 | ולא |  |
| 8:7.370 | נקרא |  |
| 8:7.371 | אליו | . |
| 8:7.372 | כי |  |
| 8:7.373 | לא |  |
| 8:7.374 | יענה |  |
| 8:7.375 | ח"ו | . |
| 8:7.376 | כמאמר |  |
| 8:7.377 | הנביא |  |
| 8:7.378 | מוכיח |  |
| 8:7.379 | לישראל |  |
| 8:7.380 | הקצור |  |
| 8:7.381 | קצרה |  |
| 8:7.382 | ידי |  |
| 8:7.383 | מפדות |  |
| 8:7.384 | ואם |  |
| 8:7.385 | אין |  |
| 8:7.386 | בי |  |
| 8:7.387 | כח |  |
| 8:7.388 | הציל |  |
| 8:7.389 | וכו | '. |
| 8:7.390 | וכביכול |  |
| 8:7.391 | לפי |  |
| 8:7.392 | דבריהם |  |
| 8:7.393 | הוא |  |
| 8:7.394 | משתבח |  |
| 8:7.395 | בעצמו |  |
| 8:7.396 | בלבד |  |
| 8:7.397 | ואינו |  |
| 8:7.398 | עושה |  |
| 8:7.399 | למבקש |  |
| 8:7.400 | ממנו |  |
| 8:7.401 | כלום | . |
| 8:7.402 | אבל |  |
| 8:7.403 | עיקר |  |
| 8:7.404 | העבודה |  |
| 8:7.405 | והתפלה |  |
| 8:7.406 | היא |  |
| 8:7.407 | לקצר |  |
| 8:7.408 | אפים | . |
| 8:7.409 | שהוא |  |
| 8:7.410 | בגוף |  |
| 8:7.411 | ואיברים |  |
| 8:7.412 | ובעל |  |
| 8:7.413 | אשה |  |
| 8:7.414 | הוא | . |
| 8:7.415 | והוא |  |
| 8:7.416 | אשר |  |
| 8:7.417 | יענה |  |
| 8:7.418 | לכל |  |
| 8:7.419 | קוראיו |  |
| 8:7.420 | ברצון |  |
| 8:7.421 | אביו |  |
| 8:7.422 | ואמו |  |
| 8:7.423 | אשר |  |
| 8:7.424 | גידלו |  |
| 8:7.425 | אותו |  |
| 8:7.426 | ורוממוהו |  |
| 8:7.427 | ואעטרו |  |
| 8:7.428 | ליה |  |
| 8:7.429 | בתרין |  |
| 8:7.430 | עטרין |  |
| 8:7.431 | עד |  |
| 8:7.432 | אשר |  |
| 8:7.433 | לו |  |
| 8:7.434 | נתכנו |  |
| 8:7.435 | עלילות | ! |
| 8:7.436 | אבל |  |
| 8:7.437 | לעתיקא |  |
| 8:7.438 | קדישא |  |
| 8:7.439 | לא |  |
| 8:7.440 | נתכנו |  |
| 8:7.441 | עלילות | , |
| 8:7.442 | והזעיר |  |
| 8:7.443 | אנפין |  |
| 8:7.444 | הוא |  |
| 8:7.445 | אשר |  |
| 8:7.446 | נקרא |  |
| 8:7.447 | בשם |  |
| 8:7.448 | הוי'ה |  |
| 8:7.449 | ובשם |  |
| 8:7.450 | הקב"ה | . |
| 8:7.451 | ואליו |  |
| 8:7.452 | תכון |  |
| 8:7.453 | עבודתינו |  |
| 8:7.454 | ותפלותינו |  |
| 8:7.455 | לפי |  |
| 8:7.456 | דבריהם |  |
| 8:7.457 | אבל |  |
| 8:7.458 | האלוה |  |
| 8:7.459 | העליון |  |
| 8:7.460 | שהוא |  |
| 8:7.461 | יחיד |  |
| 8:7.462 | בלי |  |
| 8:7.463 | גוף |  |
| 8:7.464 | ואיברים |  |
| 8:7.465 | ואין |  |
| 8:7.466 | לו |  |
| 8:7.467 | אשה | , |
| 8:7.468 | לא |  |
| 8:7.469 | נקרא |  |
| 8:7.470 | אליו |  |
| 8:7.471 | ולא |  |
| 8:7.472 | נתפלל | , |
| 8:7.473 | והעובד |  |
| 8:7.474 | אותו |  |
| 8:7.475 | ומתפלל |  |
| 8:7.476 | אליו |  |
| 8:7.477 | אין |  |
| 8:7.478 | תפלתו |  |
| 8:7.479 | תפלה |  |
| 8:7.480 | ולא |  |
| 8:7.481 | יענה |  |
| 8:7.482 | אותו | . |
| 8:7.483 | ואדרבא |  |
| 8:7.484 | ענוש |  |
| 8:7.485 | יענש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

