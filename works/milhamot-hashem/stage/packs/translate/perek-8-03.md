# Translation pack: milhamot-hashem perek-8-03 (Milchamot ha-Shem 8:4–8:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-03`.

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
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

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


## Milchamot ha-Shem 8:4

Hebrew (Wikisource): וסתירת דברי נחלת יוסף שאמר שבכח אין סוף ורוחו דבר בם ומלתו על לשונם. ובזהר בראשית שהבאתי לעיל סי' כ"ו וכ"ז מבואר שאבא ואימא הם שמינו את קצר האפים והשליטוהו על כל הברואים. ולא האין סוף. וגם מבואר שם בזהר והובא לעיל סי' כ' וכ"א שכל העלות הללו כל אחד נוטל רשות מחבירו שהוא האלוה שלמעלה ממנו, ולא מא"ס. והוא קורא לכל אחד מהם בשם הוי'ה ואלהים וקב"ה. עד שיש מקום לשומע את חבירו מזכיר שם הקב"ה המורגל בפי רז"ל ובפי כל ישראל, צריך השומע לשאול על איזה הקב"ה אתה אומר אם על עתיק יומין או אריך אנפין או אבא, או אימא, או זעיר ונוקביה, כי כל אחד מהפרצופים האלו הוא נקרא בשם הקב"ה, והם גופים מחולקים הנוטלים רשות זה מזה, ומסייעים זה את זה. כמו שאמר הרב נחלת יוסף שהמלכות היא המכה במצרים בסיועת הבינה. ואם הפועל בהם הוא אחד ומלתו על לשונם. והוא האין סוף המתלבש בהם, והוא המדבר בהם. איך לא חפץ אבא לברוא אדם ואימא חפצה לבראתו, ואמרה לאבא ומאי אכפת לך אם יחטא לי יחטא ולא לך. כמה דאיתמר ובן כסיל תוגת אמו ולא תוגת אביו? ובראה את האדם שלא כחפצו ורצונו של אבא. וכשחטא אדם הראשון גירשו אבא מג"ע וגירשה עמו, כמבואר בזהר בראשית דף כ"ב, והרי מכאן שכל אחד מהם חלש וצריך לסיוע חבירו, ואיך נעבוד לאלוה חלש הצריך לאלוה אחר לסייעו? וגם לא רוח אחד דבר בם, שהרי מה שעלה על רוח אימא לא עלה על רוח אבא בענין בריאת אדם. ומאי אית לך, למימר, ודאי שדברים כאלה לא יצאו מפי שום תנא או אמורא מרז"ל, וכל שכן רשב"י, התנא, שהרי הוא אומר כל המשתף שם שמים ודבר אחר נעקר מן העולם, (סנהדרין פרק ד' מיתות). ובגליון הבא מיורשת"ו אודות מלשינות שקר שכתב להם ארך הלשון האריכו למעניתם, ובשם קונדריס סוד ישרים מכחיש מציאות זכרים ונקיבות בהאלהים הקדושים אשר יעבדום המקובלים החדשים בתפלתם וברכותיהם, אשר יכוונו אליהם, ושכל האומר שיש לאלהים אשה הוא כופר בעיקר ומקולל, ועונו גדול כעובד ע"ז, וששם שכינה הוא אור נברא שברא האלהים למעלה עם כמה אורות קדושים רוחניים למעלה מן המלאכים ונקראים בשם ספירות וכו' ונקרא האור הזה בשם שכינה, ובאמת הוא אור נברא, ובלשון ערבי נקרא נור, והוא רוחני מאד, ואין לו שום צורה ודמות כלל, ואי אפשר לשום אדם לצייר מהותו בשכלו, כמו שאי אפשר לצייר מהות עצמות הנשמה שבגוף האדם וכו' ואור זה מתיחס לשי"ת שנקרא אור ה' ככתוב בית יעקב לכו ונלכה באור ה' וסיים בדברי הרמב"ם במורה נבוכים שכתב בענין שכינה ע"ש. ודבריו אלה המכחישים, מוכחשים באמת, ועומדים נגד הזהר והרח"ו ושאר מקובלים, ולפי דבריו בחנם טרחו ויגעו המקובלים החדשים והאריכו למעניתם בדברים רבים וספרים שונים לצייר לנו עשתונותיהם בענין האלהות וחזרנו מכל דברינו לפשוט לשלול מדעתינו כל צורה ותבנית ודמות מאותן הספירות שהן אור נברא, וכל אותן הדברים הנאמרים בקבלה החדשה, אינם רק לעבור על דברי רז"ל שאמרו כל שלא חס על כבוד קונו ראוי לו שלא בא לעולם! אתמהא! הלא מרבים העם לצייר ולדמות ולחלק את האור לחלקים ומדרגות רבות, זכרים ונקבות, בפרצופים שונים. וקראו בשמותם עלי אדמות, יש זכר ונקבה כאדם הראשון שנברא דו פרצופין דבוקים יחד דשייך בהו שאלת רז"ל הי מנייהו מסגי ברישא, ויש פרצופים חלוקים לגמרי זכר בפני עצמו ונקבה בפני עצמה כזעיר ונוקביה שאמרו שקדמה אשתו לרדת על הר סיני, ואח"כ ירד הזעיר אנפין אחריה להתחבר עמה כמו שנתבאר לעיל בשם נחלת יוסף והמה עיקר האלהות ולהם היא העבודה כמפורסם.

| id | word | punctuation after |
|---|---|---|
| 8:4.1 | וסתירת |  |
| 8:4.2 | דברי |  |
| 8:4.3 | נחלת |  |
| 8:4.4 | יוסף |  |
| 8:4.5 | שאמר |  |
| 8:4.6 | שבכח |  |
| 8:4.7 | אין |  |
| 8:4.8 | סוף |  |
| 8:4.9 | ורוחו |  |
| 8:4.10 | דבר |  |
| 8:4.11 | בם |  |
| 8:4.12 | ומלתו |  |
| 8:4.13 | על |  |
| 8:4.14 | לשונם | . |
| 8:4.15 | ובזהר |  |
| 8:4.16 | בראשית |  |
| 8:4.17 | שהבאתי |  |
| 8:4.18 | לעיל |  |
| 8:4.19 | סי | ' |
| 8:4.20 | כ"ו |  |
| 8:4.21 | וכ"ז |  |
| 8:4.22 | מבואר |  |
| 8:4.23 | שאבא |  |
| 8:4.24 | ואימא |  |
| 8:4.25 | הם |  |
| 8:4.26 | שמינו |  |
| 8:4.27 | את |  |
| 8:4.28 | קצר |  |
| 8:4.29 | האפים |  |
| 8:4.30 | והשליטוהו |  |
| 8:4.31 | על |  |
| 8:4.32 | כל |  |
| 8:4.33 | הברואים | . |
| 8:4.34 | ולא |  |
| 8:4.35 | האין |  |
| 8:4.36 | סוף | . |
| 8:4.37 | וגם |  |
| 8:4.38 | מבואר |  |
| 8:4.39 | שם |  |
| 8:4.40 | בזהר |  |
| 8:4.41 | והובא |  |
| 8:4.42 | לעיל |  |
| 8:4.43 | סי | ' |
| 8:4.44 | כ | ' |
| 8:4.45 | וכ"א |  |
| 8:4.46 | שכל |  |
| 8:4.47 | העלות |  |
| 8:4.48 | הללו |  |
| 8:4.49 | כל |  |
| 8:4.50 | אחד |  |
| 8:4.51 | נוטל |  |
| 8:4.52 | רשות |  |
| 8:4.53 | מחבירו |  |
| 8:4.54 | שהוא |  |
| 8:4.55 | האלוה |  |
| 8:4.56 | שלמעלה |  |
| 8:4.57 | ממנו | , |
| 8:4.58 | ולא |  |
| 8:4.59 | מא"ס | . |
| 8:4.60 | והוא |  |
| 8:4.61 | קורא |  |
| 8:4.62 | לכל |  |
| 8:4.63 | אחד |  |
| 8:4.64 | מהם |  |
| 8:4.65 | בשם |  |
| 8:4.66 | הוי'ה |  |
| 8:4.67 | ואלהים |  |
| 8:4.68 | וקב"ה | . |
| 8:4.69 | עד |  |
| 8:4.70 | שיש |  |
| 8:4.71 | מקום |  |
| 8:4.72 | לשומע |  |
| 8:4.73 | את |  |
| 8:4.74 | חבירו |  |
| 8:4.75 | מזכיר |  |
| 8:4.76 | שם |  |
| 8:4.77 | הקב"ה |  |
| 8:4.78 | המורגל |  |
| 8:4.79 | בפי |  |
| 8:4.80 | רז"ל |  |
| 8:4.81 | ובפי |  |
| 8:4.82 | כל |  |
| 8:4.83 | ישראל | , |
| 8:4.84 | צריך |  |
| 8:4.85 | השומע |  |
| 8:4.86 | לשאול |  |
| 8:4.87 | על |  |
| 8:4.88 | איזה |  |
| 8:4.89 | הקב"ה |  |
| 8:4.90 | אתה |  |
| 8:4.91 | אומר |  |
| 8:4.92 | אם |  |
| 8:4.93 | על |  |
| 8:4.94 | עתיק |  |
| 8:4.95 | יומין |  |
| 8:4.96 | או |  |
| 8:4.97 | אריך |  |
| 8:4.98 | אנפין |  |
| 8:4.99 | או |  |
| 8:4.100 | אבא | , |
| 8:4.101 | או |  |
| 8:4.102 | אימא | , |
| 8:4.103 | או |  |
| 8:4.104 | זעיר |  |
| 8:4.105 | ונוקביה | , |
| 8:4.106 | כי |  |
| 8:4.107 | כל |  |
| 8:4.108 | אחד |  |
| 8:4.109 | מהפרצופים |  |
| 8:4.110 | האלו |  |
| 8:4.111 | הוא |  |
| 8:4.112 | נקרא |  |
| 8:4.113 | בשם |  |
| 8:4.114 | הקב"ה | , |
| 8:4.115 | והם |  |
| 8:4.116 | גופים |  |
| 8:4.117 | מחולקים |  |
| 8:4.118 | הנוטלים |  |
| 8:4.119 | רשות |  |
| 8:4.120 | זה |  |
| 8:4.121 | מזה | , |
| 8:4.122 | ומסייעים |  |
| 8:4.123 | זה |  |
| 8:4.124 | את |  |
| 8:4.125 | זה | . |
| 8:4.126 | כמו |  |
| 8:4.127 | שאמר |  |
| 8:4.128 | הרב |  |
| 8:4.129 | נחלת |  |
| 8:4.130 | יוסף |  |
| 8:4.131 | שהמלכות |  |
| 8:4.132 | היא |  |
| 8:4.133 | המכה |  |
| 8:4.134 | במצרים |  |
| 8:4.135 | בסיועת |  |
| 8:4.136 | הבינה | . |
| 8:4.137 | ואם |  |
| 8:4.138 | הפועל |  |
| 8:4.139 | בהם |  |
| 8:4.140 | הוא |  |
| 8:4.141 | אחד |  |
| 8:4.142 | ומלתו |  |
| 8:4.143 | על |  |
| 8:4.144 | לשונם | . |
| 8:4.145 | והוא |  |
| 8:4.146 | האין |  |
| 8:4.147 | סוף |  |
| 8:4.148 | המתלבש |  |
| 8:4.149 | בהם | , |
| 8:4.150 | והוא |  |
| 8:4.151 | המדבר |  |
| 8:4.152 | בהם | . |
| 8:4.153 | איך |  |
| 8:4.154 | לא |  |
| 8:4.155 | חפץ |  |
| 8:4.156 | אבא |  |
| 8:4.157 | לברוא |  |
| 8:4.158 | אדם |  |
| 8:4.159 | ואימא |  |
| 8:4.160 | חפצה |  |
| 8:4.161 | לבראתו | , |
| 8:4.162 | ואמרה |  |
| 8:4.163 | לאבא |  |
| 8:4.164 | ומאי |  |
| 8:4.165 | אכפת |  |
| 8:4.166 | לך |  |
| 8:4.167 | אם |  |
| 8:4.168 | יחטא |  |
| 8:4.169 | לי |  |
| 8:4.170 | יחטא |  |
| 8:4.171 | ולא |  |
| 8:4.172 | לך | . |
| 8:4.173 | כמה |  |
| 8:4.174 | דאיתמר |  |
| 8:4.175 | ובן |  |
| 8:4.176 | כסיל |  |
| 8:4.177 | תוגת |  |
| 8:4.178 | אמו |  |
| 8:4.179 | ולא |  |
| 8:4.180 | תוגת |  |
| 8:4.181 | אביו | ? |
| 8:4.182 | ובראה |  |
| 8:4.183 | את |  |
| 8:4.184 | האדם |  |
| 8:4.185 | שלא |  |
| 8:4.186 | כחפצו |  |
| 8:4.187 | ורצונו |  |
| 8:4.188 | של |  |
| 8:4.189 | אבא | . |
| 8:4.190 | וכשחטא |  |
| 8:4.191 | אדם |  |
| 8:4.192 | הראשון |  |
| 8:4.193 | גירשו |  |
| 8:4.194 | אבא |  |
| 8:4.195 | מג"ע |  |
| 8:4.196 | וגירשה |  |
| 8:4.197 | עמו | , |
| 8:4.198 | כמבואר |  |
| 8:4.199 | בזהר |  |
| 8:4.200 | בראשית |  |
| 8:4.201 | דף |  |
| 8:4.202 | כ"ב | , |
| 8:4.203 | והרי |  |
| 8:4.204 | מכאן |  |
| 8:4.205 | שכל |  |
| 8:4.206 | אחד |  |
| 8:4.207 | מהם |  |
| 8:4.208 | חלש |  |
| 8:4.209 | וצריך |  |
| 8:4.210 | לסיוע |  |
| 8:4.211 | חבירו | , |
| 8:4.212 | ואיך |  |
| 8:4.213 | נעבוד |  |
| 8:4.214 | לאלוה |  |
| 8:4.215 | חלש |  |
| 8:4.216 | הצריך |  |
| 8:4.217 | לאלוה |  |
| 8:4.218 | אחר |  |
| 8:4.219 | לסייעו | ? |
| 8:4.220 | וגם |  |
| 8:4.221 | לא |  |
| 8:4.222 | רוח |  |
| 8:4.223 | אחד |  |
| 8:4.224 | דבר |  |
| 8:4.225 | בם | , |
| 8:4.226 | שהרי |  |
| 8:4.227 | מה |  |
| 8:4.228 | שעלה |  |
| 8:4.229 | על |  |
| 8:4.230 | רוח |  |
| 8:4.231 | אימא |  |
| 8:4.232 | לא |  |
| 8:4.233 | עלה |  |
| 8:4.234 | על |  |
| 8:4.235 | רוח |  |
| 8:4.236 | אבא |  |
| 8:4.237 | בענין |  |
| 8:4.238 | בריאת |  |
| 8:4.239 | אדם | . |
| 8:4.240 | ומאי |  |
| 8:4.241 | אית |  |
| 8:4.242 | לך | , |
| 8:4.243 | למימר | , |
| 8:4.244 | ודאי |  |
| 8:4.245 | שדברים |  |
| 8:4.246 | כאלה |  |
| 8:4.247 | לא |  |
| 8:4.248 | יצאו |  |
| 8:4.249 | מפי |  |
| 8:4.250 | שום |  |
| 8:4.251 | תנא |  |
| 8:4.252 | או |  |
| 8:4.253 | אמורא |  |
| 8:4.254 | מרז"ל | , |
| 8:4.255 | וכל |  |
| 8:4.256 | שכן |  |
| 8:4.257 | רשב"י | , |
| 8:4.258 | התנא | , |
| 8:4.259 | שהרי |  |
| 8:4.260 | הוא |  |
| 8:4.261 | אומר |  |
| 8:4.262 | כל |  |
| 8:4.263 | המשתף |  |
| 8:4.264 | שם |  |
| 8:4.265 | שמים |  |
| 8:4.266 | ודבר |  |
| 8:4.267 | אחר |  |
| 8:4.268 | נעקר |  |
| 8:4.269 | מן |  |
| 8:4.270 | העולם | , |
| 8:4.271 | סנהדרין |  |
| 8:4.272 | פרק |  |
| 8:4.273 | ד | ' |
| 8:4.274 | מיתות | ). |
| 8:4.275 | ובגליון |  |
| 8:4.276 | הבא |  |
| 8:4.277 | מיורשת"ו |  |
| 8:4.278 | אודות |  |
| 8:4.279 | מלשינות |  |
| 8:4.280 | שקר |  |
| 8:4.281 | שכתב |  |
| 8:4.282 | להם |  |
| 8:4.283 | ארך |  |
| 8:4.284 | הלשון |  |
| 8:4.285 | האריכו |  |
| 8:4.286 | למעניתם | , |
| 8:4.287 | ובשם |  |
| 8:4.288 | קונדריס |  |
| 8:4.289 | סוד |  |
| 8:4.290 | ישרים |  |
| 8:4.291 | מכחיש |  |
| 8:4.292 | מציאות |  |
| 8:4.293 | זכרים |  |
| 8:4.294 | ונקיבות |  |
| 8:4.295 | בהאלהים |  |
| 8:4.296 | הקדושים |  |
| 8:4.297 | אשר |  |
| 8:4.298 | יעבדום |  |
| 8:4.299 | המקובלים |  |
| 8:4.300 | החדשים |  |
| 8:4.301 | בתפלתם |  |
| 8:4.302 | וברכותיהם | , |
| 8:4.303 | אשר |  |
| 8:4.304 | יכוונו |  |
| 8:4.305 | אליהם | , |
| 8:4.306 | ושכל |  |
| 8:4.307 | האומר |  |
| 8:4.308 | שיש |  |
| 8:4.309 | לאלהים |  |
| 8:4.310 | אשה |  |
| 8:4.311 | הוא |  |
| 8:4.312 | כופר |  |
| 8:4.313 | בעיקר |  |
| 8:4.314 | ומקולל | , |
| 8:4.315 | ועונו |  |
| 8:4.316 | גדול |  |
| 8:4.317 | כעובד |  |
| 8:4.318 | ע"ז | , |
| 8:4.319 | וששם |  |
| 8:4.320 | שכינה |  |
| 8:4.321 | הוא |  |
| 8:4.322 | אור |  |
| 8:4.323 | נברא |  |
| 8:4.324 | שברא |  |
| 8:4.325 | האלהים |  |
| 8:4.326 | למעלה |  |
| 8:4.327 | עם |  |
| 8:4.328 | כמה |  |
| 8:4.329 | אורות |  |
| 8:4.330 | קדושים |  |
| 8:4.331 | רוחניים |  |
| 8:4.332 | למעלה |  |
| 8:4.333 | מן |  |
| 8:4.334 | המלאכים |  |
| 8:4.335 | ונקראים |  |
| 8:4.336 | בשם |  |
| 8:4.337 | ספירות |  |
| 8:4.338 | וכו | ' |
| 8:4.339 | ונקרא |  |
| 8:4.340 | האור |  |
| 8:4.341 | הזה |  |
| 8:4.342 | בשם |  |
| 8:4.343 | שכינה | , |
| 8:4.344 | ובאמת |  |
| 8:4.345 | הוא |  |
| 8:4.346 | אור |  |
| 8:4.347 | נברא | , |
| 8:4.348 | ובלשון |  |
| 8:4.349 | ערבי |  |
| 8:4.350 | נקרא |  |
| 8:4.351 | נור | , |
| 8:4.352 | והוא |  |
| 8:4.353 | רוחני |  |
| 8:4.354 | מאד | , |
| 8:4.355 | ואין |  |
| 8:4.356 | לו |  |
| 8:4.357 | שום |  |
| 8:4.358 | צורה |  |
| 8:4.359 | ודמות |  |
| 8:4.360 | כלל | , |
| 8:4.361 | ואי |  |
| 8:4.362 | אפשר |  |
| 8:4.363 | לשום |  |
| 8:4.364 | אדם |  |
| 8:4.365 | לצייר |  |
| 8:4.366 | מהותו |  |
| 8:4.367 | בשכלו | , |
| 8:4.368 | כמו |  |
| 8:4.369 | שאי |  |
| 8:4.370 | אפשר |  |
| 8:4.371 | לצייר |  |
| 8:4.372 | מהות |  |
| 8:4.373 | עצמות |  |
| 8:4.374 | הנשמה |  |
| 8:4.375 | שבגוף |  |
| 8:4.376 | האדם |  |
| 8:4.377 | וכו | ' |
| 8:4.378 | ואור |  |
| 8:4.379 | זה |  |
| 8:4.380 | מתיחס |  |
| 8:4.381 | לשי"ת |  |
| 8:4.382 | שנקרא |  |
| 8:4.383 | אור |  |
| 8:4.384 | ה | ' |
| 8:4.385 | ככתוב |  |
| 8:4.386 | בית |  |
| 8:4.387 | יעקב |  |
| 8:4.388 | לכו |  |
| 8:4.389 | ונלכה |  |
| 8:4.390 | באור |  |
| 8:4.391 | ה | ' |
| 8:4.392 | וסיים |  |
| 8:4.393 | בדברי |  |
| 8:4.394 | הרמב"ם |  |
| 8:4.395 | במורה |  |
| 8:4.396 | נבוכים |  |
| 8:4.397 | שכתב |  |
| 8:4.398 | בענין |  |
| 8:4.399 | שכינה |  |
| 8:4.400 | ע"ש | . |
| 8:4.401 | ודבריו |  |
| 8:4.402 | אלה |  |
| 8:4.403 | המכחישים | , |
| 8:4.404 | מוכחשים |  |
| 8:4.405 | באמת | , |
| 8:4.406 | ועומדים |  |
| 8:4.407 | נגד |  |
| 8:4.408 | הזהר |  |
| 8:4.409 | והרח"ו |  |
| 8:4.410 | ושאר |  |
| 8:4.411 | מקובלים | , |
| 8:4.412 | ולפי |  |
| 8:4.413 | דבריו |  |
| 8:4.414 | בחנם |  |
| 8:4.415 | טרחו |  |
| 8:4.416 | ויגעו |  |
| 8:4.417 | המקובלים |  |
| 8:4.418 | החדשים |  |
| 8:4.419 | והאריכו |  |
| 8:4.420 | למעניתם |  |
| 8:4.421 | בדברים |  |
| 8:4.422 | רבים |  |
| 8:4.423 | וספרים |  |
| 8:4.424 | שונים |  |
| 8:4.425 | לצייר |  |
| 8:4.426 | לנו |  |
| 8:4.427 | עשתונותיהם |  |
| 8:4.428 | בענין |  |
| 8:4.429 | האלהות |  |
| 8:4.430 | וחזרנו |  |
| 8:4.431 | מכל |  |
| 8:4.432 | דברינו |  |
| 8:4.433 | לפשוט |  |
| 8:4.434 | לשלול |  |
| 8:4.435 | מדעתינו |  |
| 8:4.436 | כל |  |
| 8:4.437 | צורה |  |
| 8:4.438 | ותבנית |  |
| 8:4.439 | ודמות |  |
| 8:4.440 | מאותן |  |
| 8:4.441 | הספירות |  |
| 8:4.442 | שהן |  |
| 8:4.443 | אור |  |
| 8:4.444 | נברא | , |
| 8:4.445 | וכל |  |
| 8:4.446 | אותן |  |
| 8:4.447 | הדברים |  |
| 8:4.448 | הנאמרים |  |
| 8:4.449 | בקבלה |  |
| 8:4.450 | החדשה | , |
| 8:4.451 | אינם |  |
| 8:4.452 | רק |  |
| 8:4.453 | לעבור |  |
| 8:4.454 | על |  |
| 8:4.455 | דברי |  |
| 8:4.456 | רז"ל |  |
| 8:4.457 | שאמרו |  |
| 8:4.458 | כל |  |
| 8:4.459 | שלא |  |
| 8:4.460 | חס |  |
| 8:4.461 | על |  |
| 8:4.462 | כבוד |  |
| 8:4.463 | קונו |  |
| 8:4.464 | ראוי |  |
| 8:4.465 | לו |  |
| 8:4.466 | שלא |  |
| 8:4.467 | בא |  |
| 8:4.468 | לעולם | ! |
| 8:4.469 | אתמהא | ! |
| 8:4.470 | הלא |  |
| 8:4.471 | מרבים |  |
| 8:4.472 | העם |  |
| 8:4.473 | לצייר |  |
| 8:4.474 | ולדמות |  |
| 8:4.475 | ולחלק |  |
| 8:4.476 | את |  |
| 8:4.477 | האור |  |
| 8:4.478 | לחלקים |  |
| 8:4.479 | ומדרגות |  |
| 8:4.480 | רבות | , |
| 8:4.481 | זכרים |  |
| 8:4.482 | ונקבות | , |
| 8:4.483 | בפרצופים |  |
| 8:4.484 | שונים | . |
| 8:4.485 | וקראו |  |
| 8:4.486 | בשמותם |  |
| 8:4.487 | עלי |  |
| 8:4.488 | אדמות | , |
| 8:4.489 | יש |  |
| 8:4.490 | זכר |  |
| 8:4.491 | ונקבה |  |
| 8:4.492 | כאדם |  |
| 8:4.493 | הראשון |  |
| 8:4.494 | שנברא |  |
| 8:4.495 | דו |  |
| 8:4.496 | פרצופין |  |
| 8:4.497 | דבוקים |  |
| 8:4.498 | יחד |  |
| 8:4.499 | דשייך |  |
| 8:4.500 | בהו |  |
| 8:4.501 | שאלת |  |
| 8:4.502 | רז"ל |  |
| 8:4.503 | הי |  |
| 8:4.504 | מנייהו |  |
| 8:4.505 | מסגי |  |
| 8:4.506 | ברישא | , |
| 8:4.507 | ויש |  |
| 8:4.508 | פרצופים |  |
| 8:4.509 | חלוקים |  |
| 8:4.510 | לגמרי |  |
| 8:4.511 | זכר |  |
| 8:4.512 | בפני |  |
| 8:4.513 | עצמו |  |
| 8:4.514 | ונקבה |  |
| 8:4.515 | בפני |  |
| 8:4.516 | עצמה |  |
| 8:4.517 | כזעיר |  |
| 8:4.518 | ונוקביה |  |
| 8:4.519 | שאמרו |  |
| 8:4.520 | שקדמה |  |
| 8:4.521 | אשתו |  |
| 8:4.522 | לרדת |  |
| 8:4.523 | על |  |
| 8:4.524 | הר |  |
| 8:4.525 | סיני | , |
| 8:4.526 | ואח"כ |  |
| 8:4.527 | ירד |  |
| 8:4.528 | הזעיר |  |
| 8:4.529 | אנפין |  |
| 8:4.530 | אחריה |  |
| 8:4.531 | להתחבר |  |
| 8:4.532 | עמה |  |
| 8:4.533 | כמו |  |
| 8:4.534 | שנתבאר |  |
| 8:4.535 | לעיל |  |
| 8:4.536 | בשם |  |
| 8:4.537 | נחלת |  |
| 8:4.538 | יוסף |  |
| 8:4.539 | והמה |  |
| 8:4.540 | עיקר |  |
| 8:4.541 | האלהות |  |
| 8:4.542 | ולהם |  |
| 8:4.543 | היא |  |
| 8:4.544 | העבודה |  |
| 8:4.545 | כמפורסם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

