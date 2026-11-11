# Translation pack: milhamot-hashem perek-4-21 (Milchamot ha-Shem 4:33–4:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-21`.

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
- באצילות: "in Emanation" (the noun the Kabbalists use), without "World of".
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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


## Milchamot ha-Shem 4:33

Hebrew (Wikisource): לט) הנה מכל זה מבואר באר היטב, כי גוף הספירות נמשכו ונשתלשלו דרך הקו הדק מעצמות האין סוף, גם העגולים גם היושר, ונמצא לפי זה כי גוף הספירות עם האור הפנימי שלהן שהוא נשמתן של ספירות עם האור המקיף שהוא מלבושן הכל מעצם אחד מן האין סוף מתפשט ונמשך, וכבר בא דבר זה מפורש בהדיא בשושן סודות בסדר התפלה וגם בסדר שבעת ימי הפסח וז"ל דע לך כי עשר ספירות אינם דבר נברא אלא ענין מתפשט מעצם הבורא ית' ואינם נבדלים ממנו והוא תמיד נמצא בהם כהדין קמצא דלבושיה מניה וכו' ע"כ, וכ"כ בגליון הזהר בשם הרמ"ז פר' בהר דף ק"ט שכלי האצילות יש להם בחינת אלהות, וכ"כ הר' משנת חסידים בסוף מסכת העשייה וז"ל והאצילות כולו בין בבחינת אורות בין בבחינת כלים ואפי' לבושיו הם כולם אלהות גמור (וכן כתב הר' עוז לאלהים בבית קדש הקדשים שהגוף והנשמה חד אינון וכן כתב בעל היכל הברכה פרשת ויקרא באוצר החיים דף ז' עמוד ד') אבל בריאה יצירה עשייה מהרוח שלהם ולמטה וגם הרוח בכלל אינם אלהות גמור עכ"ל משנת חסידים. וכן מתבאר מלשון הזהר שהבאתי לעיל סי' כ"א שהוא קורא לכל א' מן הפרצופים עלת העלות לפי שהוא עילה לעילות שתחתיו אבל אדם קדמון נקרא עילת כל העילות לפי שהוא עילה ראשונה לכל הפרצופים כלם כדלעיל. ומכאן תשובה לדבריך שהכחשת ואמרת ח"ו שהמקובלים אמרו כן, אדרבה אמרו שאין להאמין ולהעלות על לב שהספירות הם חלק מהאין סוף ב"ה שיצא ונשתלשל מעילה לעלול וכו' שדבריך אלה סותרים דברי הזהר ומהרח"ו ושושן סודות והרמ"ז, ומשנת חסידים שביארו איכות השתלשלותם מעצם האין סוף ושהם אלוהות גמור, והרי אתה מודה שדבריהם שאמרו כן הוא עוון פלילי. וכבר הבאתי לעיל דברי הזהר שקורא לאבא ואמא ה' צבאות, ולזעיר אנפין בן לאבא ואימא. ושם אמרו שכל עבודתנו וכל תפלותנו לזעיר אנפין זה שהוא בן לאבא ואימא. וכל המקובלים מסכימים כן. כמו שכתבתי לעיל בשם מקדש מלך והרש"ב ויושר לבב וספר הברית ומהרח"ו בעץ חיים והאר"י בס' הליקוטים וכסא אליהו ומצרף האמונה ונחלת יוסף. כולהו בשטה חדא קיימי שכל השבחים וכל הברכות הם לזעיר אנפין דוקא. ולא לאין סוף ולא לשאר פרצופים שלמעלה מזעיר בעולם האצילות ולא לפרצופים שלמטה ממנו בעולמות בריאה יצירה עשייה, אע"פ שהניחו ואמרו שבכל אחד מהשלשה עולמות הנז' יש בו כל הפרצופים שבעולם האצילות ומחומר אחד קורצו מעצם האין סוף נמשכו ונשתלשלו דרך הקו היוצא ממנו ויורד למטה כמו שנתבאר לעיל רק שאינם אלוהות גמור להתפלל אליהם ולקרוא להם בעת צרה כאשר יראה המעיין בדבריהם. ועל הדברים האלה והאמונות הנז' אמרו רז"ל בסנהדרין דך ס"א ע"ב מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים, מה לי קרובים ומה לי רחוקים? אלא הכי קאמר מטיבתן של קרובים אתה למד מה טיבן של רחוקים ע"כ. ופי' רש"י ז"ל לפיכך אני אומר לך, הסתכל בע"ז הקרובה לך ותראה שאין בהם ממש, ומהן אתה למד טיבן של רחוקים עכ"ל. הראו לנו רז"ל דרך בחינת צלמי ע"ז שמטיבם של קרובים נדע מה טיבם של רחוקים ומזה נקיש ונדון על כל פרצופי הקבלה החדשה האומרת שבכל עולם מארבעה עולמות נמצאים הפרצופים הנזכרים, ועכ"ז אין העבודה אלא לזעיר אנפין דאצילות. שכשם שהקרובים אלינו בעולם העשייה והיצירה והבריאה אין לעבדם שאין בהם ממש ואינם אלהות, כן זעיר אנפין דאצילות הרחוק אין בו ממש, ואין אלוה מבלעדי אלהינו. ומאחר שביררנו שהסכמת כל המקובלים החדשים שכל העבודה והשבחים והברכות לזעיר אנפין שבאצילות הנה כל טענותיך שטענת עלינו וכל הכחשותיך כולם דברי הבל ורעות רוח כאומר על האיש שהוא אשה ועל העמוד של שיש שהוא של זהב וכו'. ובתשובה זו שהשבת עשית עצמך כגולם כמאמר התנא בפ' בתרא דאבות ז' דברים בחכם וחילופיהן בגולם, ומכללן החכם שואל כהלכה ומשיב כענין ובגולם בחילוף זה. וכן אתה עשית, כי השואל שאל אותך כהלכה את מי נעבוד על פי הקבלה החדשה את עתיק, או אריך אנפין, או אבא ואימא וכו'. ובאה תשובתך שלא כענין כי קבלתך וכו' כאילו שאל אותך על קבלתך. וגם הרבית דברים והארכת להביא רמזים וגימטריות מדברי חוות יאיר שלא הבין דברי אביו. והזכיר רמזים שיוכל כל אדם לעשות כמותן מבלי קבלה כמו שעשה ר' שמואל שחיבר ספר כח ה' שהוא מלא צירופים וגימטריות ורמזים בלילה אחת מדעתו מבלי קבלה, רק להראות כחו ויכלתו לעשות צירופים וגימטריות ורמזים בנקל ככתוב בספר א"נ. ובשלמא רמזים וצירופים כל שאינו נוגע בכבוד המב"ה לעבוד אלהים אחרים ולזלזל במעלת המשנה והתלמוד ומדרשי רז"ל המספרים שבחי תורתינו הקדושה דליכא מידי דלא רמיזה באוריתא. ובפרט זה הארכת לציין הרמזים שכ' בעל מנורת המאור למקומותם, ואת המבוקש ממך לא ענית, ולא עלה על דל שפתיך כמדת הגולם שהזכיר התנא ע"ה שאינו יודע להשיב כענין המבוקש ממנו.

| id | word | punctuation after |
|---|---|---|
| 4:33.1 | לט | ) |
| 4:33.2 | הנה |  |
| 4:33.3 | מכל |  |
| 4:33.4 | זה |  |
| 4:33.5 | מבואר |  |
| 4:33.6 | באר |  |
| 4:33.7 | היטב | , |
| 4:33.8 | כי |  |
| 4:33.9 | גוף |  |
| 4:33.10 | הספירות |  |
| 4:33.11 | נמשכו |  |
| 4:33.12 | ונשתלשלו |  |
| 4:33.13 | דרך |  |
| 4:33.14 | הקו |  |
| 4:33.15 | הדק |  |
| 4:33.16 | מעצמות |  |
| 4:33.17 | האין |  |
| 4:33.18 | סוף | , |
| 4:33.19 | גם |  |
| 4:33.20 | העגולים |  |
| 4:33.21 | גם |  |
| 4:33.22 | היושר | , |
| 4:33.23 | ונמצא |  |
| 4:33.24 | לפי |  |
| 4:33.25 | זה |  |
| 4:33.26 | כי |  |
| 4:33.27 | גוף |  |
| 4:33.28 | הספירות |  |
| 4:33.29 | עם |  |
| 4:33.30 | האור |  |
| 4:33.31 | הפנימי |  |
| 4:33.32 | שלהן |  |
| 4:33.33 | שהוא |  |
| 4:33.34 | נשמתן |  |
| 4:33.35 | של |  |
| 4:33.36 | ספירות |  |
| 4:33.37 | עם |  |
| 4:33.38 | האור |  |
| 4:33.39 | המקיף |  |
| 4:33.40 | שהוא |  |
| 4:33.41 | מלבושן |  |
| 4:33.42 | הכל |  |
| 4:33.43 | מעצם |  |
| 4:33.44 | אחד |  |
| 4:33.45 | מן |  |
| 4:33.46 | האין |  |
| 4:33.47 | סוף |  |
| 4:33.48 | מתפשט |  |
| 4:33.49 | ונמשך | , |
| 4:33.50 | וכבר |  |
| 4:33.51 | בא |  |
| 4:33.52 | דבר |  |
| 4:33.53 | זה |  |
| 4:33.54 | מפורש |  |
| 4:33.55 | בהדיא |  |
| 4:33.56 | בשושן |  |
| 4:33.57 | סודות |  |
| 4:33.58 | בסדר |  |
| 4:33.59 | התפלה |  |
| 4:33.60 | וגם |  |
| 4:33.61 | בסדר |  |
| 4:33.62 | שבעת |  |
| 4:33.63 | ימי |  |
| 4:33.64 | הפסח |  |
| 4:33.65 | וז"ל |  |
| 4:33.66 | דע |  |
| 4:33.67 | לך |  |
| 4:33.68 | כי |  |
| 4:33.69 | עשר |  |
| 4:33.70 | ספירות |  |
| 4:33.71 | אינם |  |
| 4:33.72 | דבר |  |
| 4:33.73 | נברא |  |
| 4:33.74 | אלא |  |
| 4:33.75 | ענין |  |
| 4:33.76 | מתפשט |  |
| 4:33.77 | מעצם |  |
| 4:33.78 | הבורא |  |
| 4:33.79 | ית | ' |
| 4:33.80 | ואינם |  |
| 4:33.81 | נבדלים |  |
| 4:33.82 | ממנו |  |
| 4:33.83 | והוא |  |
| 4:33.84 | תמיד |  |
| 4:33.85 | נמצא |  |
| 4:33.86 | בהם |  |
| 4:33.87 | כהדין |  |
| 4:33.88 | קמצא |  |
| 4:33.89 | דלבושיה |  |
| 4:33.90 | מניה |  |
| 4:33.91 | וכו | ' |
| 4:33.92 | ע"כ | , |
| 4:33.93 | וכ"כ |  |
| 4:33.94 | בגליון |  |
| 4:33.95 | הזהר |  |
| 4:33.96 | בשם |  |
| 4:33.97 | הרמ"ז |  |
| 4:33.98 | פר | ' |
| 4:33.99 | בהר |  |
| 4:33.100 | דף |  |
| 4:33.101 | ק"ט |  |
| 4:33.102 | שכלי |  |
| 4:33.103 | האצילות |  |
| 4:33.104 | יש |  |
| 4:33.105 | להם |  |
| 4:33.106 | בחינת |  |
| 4:33.107 | אלהות | , |
| 4:33.108 | וכ"כ |  |
| 4:33.109 | הר | ' |
| 4:33.110 | משנת |  |
| 4:33.111 | חסידים |  |
| 4:33.112 | בסוף |  |
| 4:33.113 | מסכת |  |
| 4:33.114 | העשייה |  |
| 4:33.115 | וז"ל |  |
| 4:33.116 | והאצילות |  |
| 4:33.117 | כולו |  |
| 4:33.118 | בין |  |
| 4:33.119 | בבחינת |  |
| 4:33.120 | אורות |  |
| 4:33.121 | בין |  |
| 4:33.122 | בבחינת |  |
| 4:33.123 | כלים |  |
| 4:33.124 | ואפי | ' |
| 4:33.125 | לבושיו |  |
| 4:33.126 | הם |  |
| 4:33.127 | כולם |  |
| 4:33.128 | אלהות |  |
| 4:33.129 | גמור |  |
| 4:33.130 | וכן |  |
| 4:33.131 | כתב |  |
| 4:33.132 | הר | ' |
| 4:33.133 | עוז |  |
| 4:33.134 | לאלהים |  |
| 4:33.135 | בבית |  |
| 4:33.136 | קדש |  |
| 4:33.137 | הקדשים |  |
| 4:33.138 | שהגוף |  |
| 4:33.139 | והנשמה |  |
| 4:33.140 | חד |  |
| 4:33.141 | אינון |  |
| 4:33.142 | וכן |  |
| 4:33.143 | כתב |  |
| 4:33.144 | בעל |  |
| 4:33.145 | היכל |  |
| 4:33.146 | הברכה |  |
| 4:33.147 | פרשת |  |
| 4:33.148 | ויקרא |  |
| 4:33.149 | באוצר |  |
| 4:33.150 | החיים |  |
| 4:33.151 | דף |  |
| 4:33.152 | ז | ' |
| 4:33.153 | עמוד |  |
| 4:33.154 | ד | ') |
| 4:33.155 | אבל |  |
| 4:33.156 | בריאה |  |
| 4:33.157 | יצירה |  |
| 4:33.158 | עשייה |  |
| 4:33.159 | מהרוח |  |
| 4:33.160 | שלהם |  |
| 4:33.161 | ולמטה |  |
| 4:33.162 | וגם |  |
| 4:33.163 | הרוח |  |
| 4:33.164 | בכלל |  |
| 4:33.165 | אינם |  |
| 4:33.166 | אלהות |  |
| 4:33.167 | גמור |  |
| 4:33.168 | עכ"ל |  |
| 4:33.169 | משנת |  |
| 4:33.170 | חסידים | . |
| 4:33.171 | וכן |  |
| 4:33.172 | מתבאר |  |
| 4:33.173 | מלשון |  |
| 4:33.174 | הזהר |  |
| 4:33.175 | שהבאתי |  |
| 4:33.176 | לעיל |  |
| 4:33.177 | סי | ' |
| 4:33.178 | כ"א |  |
| 4:33.179 | שהוא |  |
| 4:33.180 | קורא |  |
| 4:33.181 | לכל |  |
| 4:33.182 | א | ' |
| 4:33.183 | מן |  |
| 4:33.184 | הפרצופים |  |
| 4:33.185 | עלת |  |
| 4:33.186 | העלות |  |
| 4:33.187 | לפי |  |
| 4:33.188 | שהוא |  |
| 4:33.189 | עילה |  |
| 4:33.190 | לעילות |  |
| 4:33.191 | שתחתיו |  |
| 4:33.192 | אבל |  |
| 4:33.193 | אדם |  |
| 4:33.194 | קדמון |  |
| 4:33.195 | נקרא |  |
| 4:33.196 | עילת |  |
| 4:33.197 | כל |  |
| 4:33.198 | העילות |  |
| 4:33.199 | לפי |  |
| 4:33.200 | שהוא |  |
| 4:33.201 | עילה |  |
| 4:33.202 | ראשונה |  |
| 4:33.203 | לכל |  |
| 4:33.204 | הפרצופים |  |
| 4:33.205 | כלם |  |
| 4:33.206 | כדלעיל | . |
| 4:33.207 | ומכאן |  |
| 4:33.208 | תשובה |  |
| 4:33.209 | לדבריך |  |
| 4:33.210 | שהכחשת |  |
| 4:33.211 | ואמרת |  |
| 4:33.212 | ח"ו |  |
| 4:33.213 | שהמקובלים |  |
| 4:33.214 | אמרו |  |
| 4:33.215 | כן | , |
| 4:33.216 | אדרבה |  |
| 4:33.217 | אמרו |  |
| 4:33.218 | שאין |  |
| 4:33.219 | להאמין |  |
| 4:33.220 | ולהעלות |  |
| 4:33.221 | על |  |
| 4:33.222 | לב |  |
| 4:33.223 | שהספירות |  |
| 4:33.224 | הם |  |
| 4:33.225 | חלק |  |
| 4:33.226 | מהאין |  |
| 4:33.227 | סוף |  |
| 4:33.228 | ב"ה |  |
| 4:33.229 | שיצא |  |
| 4:33.230 | ונשתלשל |  |
| 4:33.231 | מעילה |  |
| 4:33.232 | לעלול |  |
| 4:33.233 | וכו | ' |
| 4:33.234 | שדבריך |  |
| 4:33.235 | אלה |  |
| 4:33.236 | סותרים |  |
| 4:33.237 | דברי |  |
| 4:33.238 | הזהר |  |
| 4:33.239 | ומהרח"ו |  |
| 4:33.240 | ושושן |  |
| 4:33.241 | סודות |  |
| 4:33.242 | והרמ"ז | , |
| 4:33.243 | ומשנת |  |
| 4:33.244 | חסידים |  |
| 4:33.245 | שביארו |  |
| 4:33.246 | איכות |  |
| 4:33.247 | השתלשלותם |  |
| 4:33.248 | מעצם |  |
| 4:33.249 | האין |  |
| 4:33.250 | סוף |  |
| 4:33.251 | ושהם |  |
| 4:33.252 | אלוהות |  |
| 4:33.253 | גמור | , |
| 4:33.254 | והרי |  |
| 4:33.255 | אתה |  |
| 4:33.256 | מודה |  |
| 4:33.257 | שדבריהם |  |
| 4:33.258 | שאמרו |  |
| 4:33.259 | כן |  |
| 4:33.260 | הוא |  |
| 4:33.261 | עוון |  |
| 4:33.262 | פלילי | . |
| 4:33.263 | וכבר |  |
| 4:33.264 | הבאתי |  |
| 4:33.265 | לעיל |  |
| 4:33.266 | דברי |  |
| 4:33.267 | הזהר |  |
| 4:33.268 | שקורא |  |
| 4:33.269 | לאבא |  |
| 4:33.270 | ואמא |  |
| 4:33.271 | ה | ' |
| 4:33.272 | צבאות | , |
| 4:33.273 | ולזעיר |  |
| 4:33.274 | אנפין |  |
| 4:33.275 | בן |  |
| 4:33.276 | לאבא |  |
| 4:33.277 | ואימא | . |
| 4:33.278 | ושם |  |
| 4:33.279 | אמרו |  |
| 4:33.280 | שכל |  |
| 4:33.281 | עבודתנו |  |
| 4:33.282 | וכל |  |
| 4:33.283 | תפלותנו |  |
| 4:33.284 | לזעיר |  |
| 4:33.285 | אנפין |  |
| 4:33.286 | זה |  |
| 4:33.287 | שהוא |  |
| 4:33.288 | בן |  |
| 4:33.289 | לאבא |  |
| 4:33.290 | ואימא | . |
| 4:33.291 | וכל |  |
| 4:33.292 | המקובלים |  |
| 4:33.293 | מסכימים |  |
| 4:33.294 | כן | . |
| 4:33.295 | כמו |  |
| 4:33.296 | שכתבתי |  |
| 4:33.297 | לעיל |  |
| 4:33.298 | בשם |  |
| 4:33.299 | מקדש |  |
| 4:33.300 | מלך |  |
| 4:33.301 | והרש"ב |  |
| 4:33.302 | ויושר |  |
| 4:33.303 | לבב |  |
| 4:33.304 | וספר |  |
| 4:33.305 | הברית |  |
| 4:33.306 | ומהרח"ו |  |
| 4:33.307 | בעץ |  |
| 4:33.308 | חיים |  |
| 4:33.309 | והאר"י |  |
| 4:33.310 | בס | ' |
| 4:33.311 | הליקוטים |  |
| 4:33.312 | וכסא |  |
| 4:33.313 | אליהו |  |
| 4:33.314 | ומצרף |  |
| 4:33.315 | האמונה |  |
| 4:33.316 | ונחלת |  |
| 4:33.317 | יוסף | . |
| 4:33.318 | כולהו |  |
| 4:33.319 | בשטה |  |
| 4:33.320 | חדא |  |
| 4:33.321 | קיימי |  |
| 4:33.322 | שכל |  |
| 4:33.323 | השבחים |  |
| 4:33.324 | וכל |  |
| 4:33.325 | הברכות |  |
| 4:33.326 | הם |  |
| 4:33.327 | לזעיר |  |
| 4:33.328 | אנפין |  |
| 4:33.329 | דוקא | . |
| 4:33.330 | ולא |  |
| 4:33.331 | לאין |  |
| 4:33.332 | סוף |  |
| 4:33.333 | ולא |  |
| 4:33.334 | לשאר |  |
| 4:33.335 | פרצופים |  |
| 4:33.336 | שלמעלה |  |
| 4:33.337 | מזעיר |  |
| 4:33.338 | בעולם |  |
| 4:33.339 | האצילות |  |
| 4:33.340 | ולא |  |
| 4:33.341 | לפרצופים |  |
| 4:33.342 | שלמטה |  |
| 4:33.343 | ממנו |  |
| 4:33.344 | בעולמות |  |
| 4:33.345 | בריאה |  |
| 4:33.346 | יצירה |  |
| 4:33.347 | עשייה | , |
| 4:33.348 | אע"פ |  |
| 4:33.349 | שהניחו |  |
| 4:33.350 | ואמרו |  |
| 4:33.351 | שבכל |  |
| 4:33.352 | אחד |  |
| 4:33.353 | מהשלשה |  |
| 4:33.354 | עולמות |  |
| 4:33.355 | הנז | ' |
| 4:33.356 | יש |  |
| 4:33.357 | בו |  |
| 4:33.358 | כל |  |
| 4:33.359 | הפרצופים |  |
| 4:33.360 | שבעולם |  |
| 4:33.361 | האצילות |  |
| 4:33.362 | ומחומר |  |
| 4:33.363 | אחד |  |
| 4:33.364 | קורצו |  |
| 4:33.365 | מעצם |  |
| 4:33.366 | האין |  |
| 4:33.367 | סוף |  |
| 4:33.368 | נמשכו |  |
| 4:33.369 | ונשתלשלו |  |
| 4:33.370 | דרך |  |
| 4:33.371 | הקו |  |
| 4:33.372 | היוצא |  |
| 4:33.373 | ממנו |  |
| 4:33.374 | ויורד |  |
| 4:33.375 | למטה |  |
| 4:33.376 | כמו |  |
| 4:33.377 | שנתבאר |  |
| 4:33.378 | לעיל |  |
| 4:33.379 | רק |  |
| 4:33.380 | שאינם |  |
| 4:33.381 | אלוהות |  |
| 4:33.382 | גמור |  |
| 4:33.383 | להתפלל |  |
| 4:33.384 | אליהם |  |
| 4:33.385 | ולקרוא |  |
| 4:33.386 | להם |  |
| 4:33.387 | בעת |  |
| 4:33.388 | צרה |  |
| 4:33.389 | כאשר |  |
| 4:33.390 | יראה |  |
| 4:33.391 | המעיין |  |
| 4:33.392 | בדבריהם | . |
| 4:33.393 | ועל |  |
| 4:33.394 | הדברים |  |
| 4:33.395 | האלה |  |
| 4:33.396 | והאמונות |  |
| 4:33.397 | הנז | ' |
| 4:33.398 | אמרו |  |
| 4:33.399 | רז"ל |  |
| 4:33.400 | בסנהדרין |  |
| 4:33.401 | דך |  |
| 4:33.402 | ס"א |  |
| 4:33.403 | ע"ב |  |
| 4:33.404 | מאלהי |  |
| 4:33.405 | העמים |  |
| 4:33.406 | אשר |  |
| 4:33.407 | סביבותיכם |  |
| 4:33.408 | הקרובים |  |
| 4:33.409 | אליך |  |
| 4:33.410 | או |  |
| 4:33.411 | הרחוקים | , |
| 4:33.412 | מה |  |
| 4:33.413 | לי |  |
| 4:33.414 | קרובים |  |
| 4:33.415 | ומה |  |
| 4:33.416 | לי |  |
| 4:33.417 | רחוקים | ? |
| 4:33.418 | אלא |  |
| 4:33.419 | הכי |  |
| 4:33.420 | קאמר |  |
| 4:33.421 | מטיבתן |  |
| 4:33.422 | של |  |
| 4:33.423 | קרובים |  |
| 4:33.424 | אתה |  |
| 4:33.425 | למד |  |
| 4:33.426 | מה |  |
| 4:33.427 | טיבן |  |
| 4:33.428 | של |  |
| 4:33.429 | רחוקים |  |
| 4:33.430 | ע"כ | . |
| 4:33.431 | ופי | ' |
| 4:33.432 | רש"י |  |
| 4:33.433 | ז"ל |  |
| 4:33.434 | לפיכך |  |
| 4:33.435 | אני |  |
| 4:33.436 | אומר |  |
| 4:33.437 | לך | , |
| 4:33.438 | הסתכל |  |
| 4:33.439 | בע"ז |  |
| 4:33.440 | הקרובה |  |
| 4:33.441 | לך |  |
| 4:33.442 | ותראה |  |
| 4:33.443 | שאין |  |
| 4:33.444 | בהם |  |
| 4:33.445 | ממש | , |
| 4:33.446 | ומהן |  |
| 4:33.447 | אתה |  |
| 4:33.448 | למד |  |
| 4:33.449 | טיבן |  |
| 4:33.450 | של |  |
| 4:33.451 | רחוקים |  |
| 4:33.452 | עכ"ל | . |
| 4:33.453 | הראו |  |
| 4:33.454 | לנו |  |
| 4:33.455 | רז"ל |  |
| 4:33.456 | דרך |  |
| 4:33.457 | בחינת |  |
| 4:33.458 | צלמי |  |
| 4:33.459 | ע"ז |  |
| 4:33.460 | שמטיבם |  |
| 4:33.461 | של |  |
| 4:33.462 | קרובים |  |
| 4:33.463 | נדע |  |
| 4:33.464 | מה |  |
| 4:33.465 | טיבם |  |
| 4:33.466 | של |  |
| 4:33.467 | רחוקים |  |
| 4:33.468 | ומזה |  |
| 4:33.469 | נקיש |  |
| 4:33.470 | ונדון |  |
| 4:33.471 | על |  |
| 4:33.472 | כל |  |
| 4:33.473 | פרצופי |  |
| 4:33.474 | הקבלה |  |
| 4:33.475 | החדשה |  |
| 4:33.476 | האומרת |  |
| 4:33.477 | שבכל |  |
| 4:33.478 | עולם |  |
| 4:33.479 | מארבעה |  |
| 4:33.480 | עולמות |  |
| 4:33.481 | נמצאים |  |
| 4:33.482 | הפרצופים |  |
| 4:33.483 | הנזכרים | , |
| 4:33.484 | ועכ"ז |  |
| 4:33.485 | אין |  |
| 4:33.486 | העבודה |  |
| 4:33.487 | אלא |  |
| 4:33.488 | לזעיר |  |
| 4:33.489 | אנפין |  |
| 4:33.490 | דאצילות | . |
| 4:33.491 | שכשם |  |
| 4:33.492 | שהקרובים |  |
| 4:33.493 | אלינו |  |
| 4:33.494 | בעולם |  |
| 4:33.495 | העשייה |  |
| 4:33.496 | והיצירה |  |
| 4:33.497 | והבריאה |  |
| 4:33.498 | אין |  |
| 4:33.499 | לעבדם |  |
| 4:33.500 | שאין |  |
| 4:33.501 | בהם |  |
| 4:33.502 | ממש |  |
| 4:33.503 | ואינם |  |
| 4:33.504 | אלהות | , |
| 4:33.505 | כן |  |
| 4:33.506 | זעיר |  |
| 4:33.507 | אנפין |  |
| 4:33.508 | דאצילות |  |
| 4:33.509 | הרחוק |  |
| 4:33.510 | אין |  |
| 4:33.511 | בו |  |
| 4:33.512 | ממש | , |
| 4:33.513 | ואין |  |
| 4:33.514 | אלוה |  |
| 4:33.515 | מבלעדי |  |
| 4:33.516 | אלהינו | . |
| 4:33.517 | ומאחר |  |
| 4:33.518 | שביררנו |  |
| 4:33.519 | שהסכמת |  |
| 4:33.520 | כל |  |
| 4:33.521 | המקובלים |  |
| 4:33.522 | החדשים |  |
| 4:33.523 | שכל |  |
| 4:33.524 | העבודה |  |
| 4:33.525 | והשבחים |  |
| 4:33.526 | והברכות |  |
| 4:33.527 | לזעיר |  |
| 4:33.528 | אנפין |  |
| 4:33.529 | שבאצילות |  |
| 4:33.530 | הנה |  |
| 4:33.531 | כל |  |
| 4:33.532 | טענותיך |  |
| 4:33.533 | שטענת |  |
| 4:33.534 | עלינו |  |
| 4:33.535 | וכל |  |
| 4:33.536 | הכחשותיך |  |
| 4:33.537 | כולם |  |
| 4:33.538 | דברי |  |
| 4:33.539 | הבל |  |
| 4:33.540 | ורעות |  |
| 4:33.541 | רוח |  |
| 4:33.542 | כאומר |  |
| 4:33.543 | על |  |
| 4:33.544 | האיש |  |
| 4:33.545 | שהוא |  |
| 4:33.546 | אשה |  |
| 4:33.547 | ועל |  |
| 4:33.548 | העמוד |  |
| 4:33.549 | של |  |
| 4:33.550 | שיש |  |
| 4:33.551 | שהוא |  |
| 4:33.552 | של |  |
| 4:33.553 | זהב |  |
| 4:33.554 | וכו | '. |
| 4:33.555 | ובתשובה |  |
| 4:33.556 | זו |  |
| 4:33.557 | שהשבת |  |
| 4:33.558 | עשית |  |
| 4:33.559 | עצמך |  |
| 4:33.560 | כגולם |  |
| 4:33.561 | כמאמר |  |
| 4:33.562 | התנא |  |
| 4:33.563 | בפ | ' |
| 4:33.564 | בתרא |  |
| 4:33.565 | דאבות |  |
| 4:33.566 | ז | ' |
| 4:33.567 | דברים |  |
| 4:33.568 | בחכם |  |
| 4:33.569 | וחילופיהן |  |
| 4:33.570 | בגולם | , |
| 4:33.571 | ומכללן |  |
| 4:33.572 | החכם |  |
| 4:33.573 | שואל |  |
| 4:33.574 | כהלכה |  |
| 4:33.575 | ומשיב |  |
| 4:33.576 | כענין |  |
| 4:33.577 | ובגולם |  |
| 4:33.578 | בחילוף |  |
| 4:33.579 | זה | . |
| 4:33.580 | וכן |  |
| 4:33.581 | אתה |  |
| 4:33.582 | עשית | , |
| 4:33.583 | כי |  |
| 4:33.584 | השואל |  |
| 4:33.585 | שאל |  |
| 4:33.586 | אותך |  |
| 4:33.587 | כהלכה |  |
| 4:33.588 | את |  |
| 4:33.589 | מי |  |
| 4:33.590 | נעבוד |  |
| 4:33.591 | על |  |
| 4:33.592 | פי |  |
| 4:33.593 | הקבלה |  |
| 4:33.594 | החדשה |  |
| 4:33.595 | את |  |
| 4:33.596 | עתיק | , |
| 4:33.597 | או |  |
| 4:33.598 | אריך |  |
| 4:33.599 | אנפין | , |
| 4:33.600 | או |  |
| 4:33.601 | אבא |  |
| 4:33.602 | ואימא |  |
| 4:33.603 | וכו | '. |
| 4:33.604 | ובאה |  |
| 4:33.605 | תשובתך |  |
| 4:33.606 | שלא |  |
| 4:33.607 | כענין |  |
| 4:33.608 | כי |  |
| 4:33.609 | קבלתך |  |
| 4:33.610 | וכו | ' |
| 4:33.611 | כאילו |  |
| 4:33.612 | שאל |  |
| 4:33.613 | אותך |  |
| 4:33.614 | על |  |
| 4:33.615 | קבלתך | . |
| 4:33.616 | וגם |  |
| 4:33.617 | הרבית |  |
| 4:33.618 | דברים |  |
| 4:33.619 | והארכת |  |
| 4:33.620 | להביא |  |
| 4:33.621 | רמזים |  |
| 4:33.622 | וגימטריות |  |
| 4:33.623 | מדברי |  |
| 4:33.624 | חוות |  |
| 4:33.625 | יאיר |  |
| 4:33.626 | שלא |  |
| 4:33.627 | הבין |  |
| 4:33.628 | דברי |  |
| 4:33.629 | אביו | . |
| 4:33.630 | והזכיר |  |
| 4:33.631 | רמזים |  |
| 4:33.632 | שיוכל |  |
| 4:33.633 | כל |  |
| 4:33.634 | אדם |  |
| 4:33.635 | לעשות |  |
| 4:33.636 | כמותן |  |
| 4:33.637 | מבלי |  |
| 4:33.638 | קבלה |  |
| 4:33.639 | כמו |  |
| 4:33.640 | שעשה |  |
| 4:33.641 | ר | ' |
| 4:33.642 | שמואל |  |
| 4:33.643 | שחיבר |  |
| 4:33.644 | ספר |  |
| 4:33.645 | כח |  |
| 4:33.646 | ה | ' |
| 4:33.647 | שהוא |  |
| 4:33.648 | מלא |  |
| 4:33.649 | צירופים |  |
| 4:33.650 | וגימטריות |  |
| 4:33.651 | ורמזים |  |
| 4:33.652 | בלילה |  |
| 4:33.653 | אחת |  |
| 4:33.654 | מדעתו |  |
| 4:33.655 | מבלי |  |
| 4:33.656 | קבלה | , |
| 4:33.657 | רק |  |
| 4:33.658 | להראות |  |
| 4:33.659 | כחו |  |
| 4:33.660 | ויכלתו |  |
| 4:33.661 | לעשות |  |
| 4:33.662 | צירופים |  |
| 4:33.663 | וגימטריות |  |
| 4:33.664 | ורמזים |  |
| 4:33.665 | בנקל |  |
| 4:33.666 | ככתוב |  |
| 4:33.667 | בספר |  |
| 4:33.668 | א"נ | . |
| 4:33.669 | ובשלמא |  |
| 4:33.670 | רמזים |  |
| 4:33.671 | וצירופים |  |
| 4:33.672 | כל |  |
| 4:33.673 | שאינו |  |
| 4:33.674 | נוגע |  |
| 4:33.675 | בכבוד |  |
| 4:33.676 | המב"ה |  |
| 4:33.677 | לעבוד |  |
| 4:33.678 | אלהים |  |
| 4:33.679 | אחרים |  |
| 4:33.680 | ולזלזל |  |
| 4:33.681 | במעלת |  |
| 4:33.682 | המשנה |  |
| 4:33.683 | והתלמוד |  |
| 4:33.684 | ומדרשי |  |
| 4:33.685 | רז"ל |  |
| 4:33.686 | המספרים |  |
| 4:33.687 | שבחי |  |
| 4:33.688 | תורתינו |  |
| 4:33.689 | הקדושה |  |
| 4:33.690 | דליכא |  |
| 4:33.691 | מידי |  |
| 4:33.692 | דלא |  |
| 4:33.693 | רמיזה |  |
| 4:33.694 | באוריתא | . |
| 4:33.695 | ובפרט |  |
| 4:33.696 | זה |  |
| 4:33.697 | הארכת |  |
| 4:33.698 | לציין |  |
| 4:33.699 | הרמזים |  |
| 4:33.700 | שכ | ' |
| 4:33.701 | בעל |  |
| 4:33.702 | מנורת |  |
| 4:33.703 | המאור |  |
| 4:33.704 | למקומותם | , |
| 4:33.705 | ואת |  |
| 4:33.706 | המבוקש |  |
| 4:33.707 | ממך |  |
| 4:33.708 | לא |  |
| 4:33.709 | ענית | , |
| 4:33.710 | ולא |  |
| 4:33.711 | עלה |  |
| 4:33.712 | על |  |
| 4:33.713 | דל |  |
| 4:33.714 | שפתיך |  |
| 4:33.715 | כמדת |  |
| 4:33.716 | הגולם |  |
| 4:33.717 | שהזכיר |  |
| 4:33.718 | התנא |  |
| 4:33.719 | ע"ה |  |
| 4:33.720 | שאינו |  |
| 4:33.721 | יודע |  |
| 4:33.722 | להשיב |  |
| 4:33.723 | כענין |  |
| 4:33.724 | המבוקש |  |
| 4:33.725 | ממנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

