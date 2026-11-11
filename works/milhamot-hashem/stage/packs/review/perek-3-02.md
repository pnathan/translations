# Review pack: milhamot-hashem perek-3-02 (Milchamot ha-Shem 3:3–3:4)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-02`.

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
- A Biblical quotation whose wording differs from the verse follows the
  author's wording; the difference goes in `apparatus`.
- `literal` on a long paragraph: the glosses in Hebrew order, hyphens joining
  the parts of one word. No reasoning in it; reasons go in `note`.
- `context` may carry a short note on a story the author cites (the Pardes
  story), without a reference unless the author gives one.

## Rulings from the pilot reviews, by the editor

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- ולא אמרינן כלא חד (1:7) and באגרונו (1:12) are for the owner: the main
  English stands as translated; the reviewer's alternative is in `readings`
  / `apparatus`.

## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
- A title the printed text spells wrongly (ונור אלצלאם): translate the printed
  spelling ("Nur al-Tzalam"); the likelier title (Nur al-Zulam, Me'or ha-Afelah)
  goes in the token note and `apparatus`, not in the English.
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
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 3:3

Hebrew (Wikisource): ד) ולא יהא כח רבותינו בעלי הקבלה החדשה גדול מכח ב"ד הגדול היושבים בלשכת הגזית ולא השגתם יתירה מהשגת סנהדרי גדולה! ועם כל זה אמרה תורה שאם טעו והורו להתיר את האסור, או שגגו בע"ז וכו'. למדנו מן התורה ומרבותינו במסכת הוריות, שאפשר לב"ד הגדול לטעות ולהורות התר באיזה פרט בע"ז. כי האדם קורץ מחומר בין גדול בין קטן, ועלול הוא לטעות, ולכן צותה תורה כי יפלא ממך דבר למשפט וכו' וקמת ועלית אל המקום אשר יבחר ה' אלהיך כדי למעט הטעות. איך אינו נמנע שלא יטעו אפי' ב"ד הגדול. שהרי נאמר בתורה ונעלם דבר מעיני הקהל. והיה אם מעיני העדה נעשתה לשגגה. ולפי קוצר דעתינו בראותינו ההפרש הגדול שיש בין רבותינו הראשונים ואיך היתה אמונתם בענין האלוהות ויחוד השי"ת על פי תוה"ק כמבואר בר' בחיי בעל חובת הלבבות ור' יהודה הלוי בספר הכוזרי ובתפלותיו ור' סעדיה גאון בס' האמונות והדעות והרמב"ם בחיבורו ובפי' המשנה ובס' מורה הנבוכים. והרוקח בשער הייחוד, וס' מצות הגדול ורבי יוסף אלבו בספר העיקרים, ור' מאיר אלדבי בספר שבילי אמונה. וזולתם רבים וכן שלימים. ואיך הוא עתה בספרי רבותינו האחרונים ע"פ הקבלה החדשה הנמשכת מתחלת אלף השישי בחבלי השוא עד שנעשית בזמנינו כעבות העגלה והאמינו באלהים רבים. ושהעיקר בהם אשר אליו כל עבודתינו וברכותינו הוא הפרצוף האחרון שבאצילות הנקרא (קצר אפים) זעיר אנפין.

| id | word | punctuation after |
|---|---|---|
| 3:3.1 | ד | ) |
| 3:3.2 | ולא |  |
| 3:3.3 | יהא |  |
| 3:3.4 | כח |  |
| 3:3.5 | רבותינו |  |
| 3:3.6 | בעלי |  |
| 3:3.7 | הקבלה |  |
| 3:3.8 | החדשה |  |
| 3:3.9 | גדול |  |
| 3:3.10 | מכח |  |
| 3:3.11 | ב"ד |  |
| 3:3.12 | הגדול |  |
| 3:3.13 | היושבים |  |
| 3:3.14 | בלשכת |  |
| 3:3.15 | הגזית |  |
| 3:3.16 | ולא |  |
| 3:3.17 | השגתם |  |
| 3:3.18 | יתירה |  |
| 3:3.19 | מהשגת |  |
| 3:3.20 | סנהדרי |  |
| 3:3.21 | גדולה | ! |
| 3:3.22 | ועם |  |
| 3:3.23 | כל |  |
| 3:3.24 | זה |  |
| 3:3.25 | אמרה |  |
| 3:3.26 | תורה |  |
| 3:3.27 | שאם |  |
| 3:3.28 | טעו |  |
| 3:3.29 | והורו |  |
| 3:3.30 | להתיר |  |
| 3:3.31 | את |  |
| 3:3.32 | האסור | , |
| 3:3.33 | או |  |
| 3:3.34 | שגגו |  |
| 3:3.35 | בע"ז |  |
| 3:3.36 | וכו | '. |
| 3:3.37 | למדנו |  |
| 3:3.38 | מן |  |
| 3:3.39 | התורה |  |
| 3:3.40 | ומרבותינו |  |
| 3:3.41 | במסכת |  |
| 3:3.42 | הוריות | , |
| 3:3.43 | שאפשר |  |
| 3:3.44 | לב"ד |  |
| 3:3.45 | הגדול |  |
| 3:3.46 | לטעות |  |
| 3:3.47 | ולהורות |  |
| 3:3.48 | התר |  |
| 3:3.49 | באיזה |  |
| 3:3.50 | פרט |  |
| 3:3.51 | בע"ז | . |
| 3:3.52 | כי |  |
| 3:3.53 | האדם |  |
| 3:3.54 | קורץ |  |
| 3:3.55 | מחומר |  |
| 3:3.56 | בין |  |
| 3:3.57 | גדול |  |
| 3:3.58 | בין |  |
| 3:3.59 | קטן | , |
| 3:3.60 | ועלול |  |
| 3:3.61 | הוא |  |
| 3:3.62 | לטעות | , |
| 3:3.63 | ולכן |  |
| 3:3.64 | צותה |  |
| 3:3.65 | תורה |  |
| 3:3.66 | כי |  |
| 3:3.67 | יפלא |  |
| 3:3.68 | ממך |  |
| 3:3.69 | דבר |  |
| 3:3.70 | למשפט |  |
| 3:3.71 | וכו | ' |
| 3:3.72 | וקמת |  |
| 3:3.73 | ועלית |  |
| 3:3.74 | אל |  |
| 3:3.75 | המקום |  |
| 3:3.76 | אשר |  |
| 3:3.77 | יבחר |  |
| 3:3.78 | ה | ' |
| 3:3.79 | אלהיך |  |
| 3:3.80 | כדי |  |
| 3:3.81 | למעט |  |
| 3:3.82 | הטעות | . |
| 3:3.83 | איך |  |
| 3:3.84 | אינו |  |
| 3:3.85 | נמנע |  |
| 3:3.86 | שלא |  |
| 3:3.87 | יטעו |  |
| 3:3.88 | אפי | ' |
| 3:3.89 | ב"ד |  |
| 3:3.90 | הגדול | . |
| 3:3.91 | שהרי |  |
| 3:3.92 | נאמר |  |
| 3:3.93 | בתורה |  |
| 3:3.94 | ונעלם |  |
| 3:3.95 | דבר |  |
| 3:3.96 | מעיני |  |
| 3:3.97 | הקהל | . |
| 3:3.98 | והיה |  |
| 3:3.99 | אם |  |
| 3:3.100 | מעיני |  |
| 3:3.101 | העדה |  |
| 3:3.102 | נעשתה |  |
| 3:3.103 | לשגגה | . |
| 3:3.104 | ולפי |  |
| 3:3.105 | קוצר |  |
| 3:3.106 | דעתינו |  |
| 3:3.107 | בראותינו |  |
| 3:3.108 | ההפרש |  |
| 3:3.109 | הגדול |  |
| 3:3.110 | שיש |  |
| 3:3.111 | בין |  |
| 3:3.112 | רבותינו |  |
| 3:3.113 | הראשונים |  |
| 3:3.114 | ואיך |  |
| 3:3.115 | היתה |  |
| 3:3.116 | אמונתם |  |
| 3:3.117 | בענין |  |
| 3:3.118 | האלוהות |  |
| 3:3.119 | ויחוד |  |
| 3:3.120 | השי"ת |  |
| 3:3.121 | על |  |
| 3:3.122 | פי |  |
| 3:3.123 | תוה"ק |  |
| 3:3.124 | כמבואר |  |
| 3:3.125 | בר | ' |
| 3:3.126 | בחיי |  |
| 3:3.127 | בעל |  |
| 3:3.128 | חובת |  |
| 3:3.129 | הלבבות |  |
| 3:3.130 | ור | ' |
| 3:3.131 | יהודה |  |
| 3:3.132 | הלוי |  |
| 3:3.133 | בספר |  |
| 3:3.134 | הכוזרי |  |
| 3:3.135 | ובתפלותיו |  |
| 3:3.136 | ור | ' |
| 3:3.137 | סעדיה |  |
| 3:3.138 | גאון |  |
| 3:3.139 | בס | ' |
| 3:3.140 | האמונות |  |
| 3:3.141 | והדעות |  |
| 3:3.142 | והרמב"ם |  |
| 3:3.143 | בחיבורו |  |
| 3:3.144 | ובפי | ' |
| 3:3.145 | המשנה |  |
| 3:3.146 | ובס | ' |
| 3:3.147 | מורה |  |
| 3:3.148 | הנבוכים | . |
| 3:3.149 | והרוקח |  |
| 3:3.150 | בשער |  |
| 3:3.151 | הייחוד | , |
| 3:3.152 | וס | ' |
| 3:3.153 | מצות |  |
| 3:3.154 | הגדול |  |
| 3:3.155 | ורבי |  |
| 3:3.156 | יוסף |  |
| 3:3.157 | אלבו |  |
| 3:3.158 | בספר |  |
| 3:3.159 | העיקרים | , |
| 3:3.160 | ור | ' |
| 3:3.161 | מאיר |  |
| 3:3.162 | אלדבי |  |
| 3:3.163 | בספר |  |
| 3:3.164 | שבילי |  |
| 3:3.165 | אמונה | . |
| 3:3.166 | וזולתם |  |
| 3:3.167 | רבים |  |
| 3:3.168 | וכן |  |
| 3:3.169 | שלימים | . |
| 3:3.170 | ואיך |  |
| 3:3.171 | הוא |  |
| 3:3.172 | עתה |  |
| 3:3.173 | בספרי |  |
| 3:3.174 | רבותינו |  |
| 3:3.175 | האחרונים |  |
| 3:3.176 | ע"פ |  |
| 3:3.177 | הקבלה |  |
| 3:3.178 | החדשה |  |
| 3:3.179 | הנמשכת |  |
| 3:3.180 | מתחלת |  |
| 3:3.181 | אלף |  |
| 3:3.182 | השישי |  |
| 3:3.183 | בחבלי |  |
| 3:3.184 | השוא |  |
| 3:3.185 | עד |  |
| 3:3.186 | שנעשית |  |
| 3:3.187 | בזמנינו |  |
| 3:3.188 | כעבות |  |
| 3:3.189 | העגלה |  |
| 3:3.190 | והאמינו |  |
| 3:3.191 | באלהים |  |
| 3:3.192 | רבים | . |
| 3:3.193 | ושהעיקר |  |
| 3:3.194 | בהם |  |
| 3:3.195 | אשר |  |
| 3:3.196 | אליו |  |
| 3:3.197 | כל |  |
| 3:3.198 | עבודתינו |  |
| 3:3.199 | וברכותינו |  |
| 3:3.200 | הוא |  |
| 3:3.201 | הפרצוף |  |
| 3:3.202 | האחרון |  |
| 3:3.203 | שבאצילות |  |
| 3:3.204 | הנקרא |  |
| 3:3.205 | קצר |  |
| 3:3.206 | אפים | ) |
| 3:3.207 | זעיר |  |
| 3:3.208 | אנפין | . |


### Translation under review (3:3)

English: 4) The authority of our rabbis, the masters of the new Kabbalah, shall not be greater than the authority of the Great Beit Din who sat in the Chamber of Hewn Stone, and their grasp is no greater than the grasp of the Great Sanhedrin! And yet the Torah said that if they erred and ruled to permit the forbidden, or sinned inadvertently in idolatry, etc. We have learned from the Torah and from our rabbis, in the tractate Horayot, that the Great Beit Din can err and rule a permission in some particular case of idolatry. For man is formed of clay, whether great or small, and is prone to err; therefore the Torah commanded: "If a matter is too hard for you in judgment, etc. Then arise and go up to the place that the LORD your God will choose," in order to lessen the error. How is it not unavoidable that even the Great Beit Din should err. For it is said in the Torah, "and the matter was hidden from the eyes of the assembly. And it shall be, if from the eyes of the congregation it was done in error." And according to the shortness of our understanding, when we see the great difference between our early rabbis, and how was their faith in the matter of divinity and the unity of the Name, may He be blessed, according to the holy Torah, as explained by Rabbeinu Bachya, the author of Chovot ha-Levavot, and Rabbi Yehuda Halevi in the Kuzari and in his prayers, and Rabbi Saadia Gaon in the Emunot ve-De'ot, and the Rambam in his work, and in his commentary on the Mishnah, and in the Moreh Nevukhim. And the Roke'ach in the Gate of Unity, and the Sefer Mitzvot Gadol, and Rabbi Yosef Albo in the Sefer ha-Ikkarim, and Rabbi Meir Aldabi in the Shvilei Emunah. And many others, and likewise the complete ones. And how is it now in the books of our later rabbis, according to the new Kabbalah, which runs from the start of the sixth millennium with cords of falsehood until it became in our time as with cart-ropes, and they believed in many gods. And that the principal point in them, to which all our worship and blessings are directed, is the last partzuf in Emanation, called (quick to anger) Ze'ir Anpin.

Literal: 4) and-not will-be authority of-our-rabbis masters-of the-Kabbalah the-new greater than-the-authority-of Beit-Din the-great who-sit in-the-chamber-of the-hewn-stone and-not their-grasp greater than-the-grasp-of Sanhedrin the-great! and-yet all this said the-Torah that-if erred and-ruled to-permit the forbidden or sinned-inadvertently in-idolatry etc. we-learned from the-Torah and-from our-rabbis in the-tractate Horayot that-possible for-the-Beit-Din the-great to-err and-to-rule permission in-which particular-case in-idolatry. for man is-formed of-clay whether great or small and-prone he to-err therefore commanded the-Torah that is-too-hard for-you a-matter for-judgment etc. and-arise and-go-up to the-place which will-choose the-LORD your-God in-order to-lessen the-error how is-it-not avoidable that-not they-would-err even Beit-Din the-great for-behold it-is-said in-the-Torah and-was-hidden a-matter from-the-eyes-of the-assembly and-it-shall-be if from-the-eyes-of the-congregation was-done in-error and-according-to the-shortness-of our-understanding when-we-see the-difference the-great that-there-is between our-rabbis the-early and-how was their-faith in-the-matter-of the-divinity and-the-unity the-Name-may-He-be-blessed according-to the-mouth-of the-holy-Torah as-explained by-Rabbeinu Bachya the-author-of Chovot ha-Levavot and-Rabbi Yehuda Halevi in-the-book the-Kuzari and-in-his-prayers and-Rabbi Saadia Gaon in-the-book the-Emunot and-the-De'ot and-the-Rambam in-his-work and-in-the-commentary of-the-Mishnah and-in-the-book Moreh Nevukhim and-the-Roke'ach in-the-gate of-unity and-the-book-of Mitzvot the-great and-Rabbi Yosef Albo in-the-book the-Ikkarim and-Rabbi Meir Aldabi in-the-book Shvilei Emunah and-others many and-so-too the-complete-ones and-how it now in-the-books-of our-rabbis the-later according-to the-Kabbalah the-new drawn-out from-the-start-of the-thousand the-sixth with-cords-of falsehood until it-became in-our-time as-with cart-ropes and-they-believed in-gods many and-that the-principal in-them to-which all our-worship and-our-blessings is the-partzuf the-last in-Emanation called quick to-anger Ze'ir Anpin

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:3.1 | 4) |  |  | printed list marker ד) (the fourth item); kept as 4) |
| 3:3.2 | and not |  |  |  |
| 3:3.3 | will be |  |  |  |
| 3:3.4 | authority |  |  |  |
| 3:3.5 | our rabbis |  |  |  |
| 3:3.6 | masters of |  |  |  |
| 3:3.7 | the Kabbalah |  |  |  |
| 3:3.8 | the new |  |  |  |
| 3:3.9 | greater |  |  |  |
| 3:3.10 | than the authority of |  |  |  |
| 3:3.11 | Beit Din (court) |  |  | Beit Din, the court of law (ב"ד) |
| 3:3.12 | the great |  |  |  |
| 3:3.13 | who sit |  |  |  |
| 3:3.14 | in the chamber of |  |  |  |
| 3:3.15 | the hewn stone |  |  | lishkat ha-gazit: the Chamber of Hewn Stone; see context |
| 3:3.16 | and not |  |  |  |
| 3:3.17 | their grasp |  |  |  |
| 3:3.18 | greater |  |  |  |
| 3:3.19 | than the grasp of |  |  |  |
| 3:3.20 | Sanhedrin |  |  | Sanhedrin Gedolah: the Great Sanhedrin |
| 3:3.21 | the Great (gedolah) |  |  |  |
| 3:3.22 | and yet |  |  |  |
| 3:3.23 | all |  |  |  |
| 3:3.24 | this |  |  |  |
| 3:3.25 | said |  |  |  |
| 3:3.26 | the Torah |  |  |  |
| 3:3.27 | that if |  |  |  |
| 3:3.28 | they erred |  |  |  |
| 3:3.29 | and ruled |  |  |  |
| 3:3.30 | to permit |  |  |  |
| 3:3.31 | the |  |  |  |
| 3:3.32 | forbidden |  |  |  |
| 3:3.33 | or |  |  |  |
| 3:3.34 | sinned inadvertently |  |  |  |
| 3:3.35 | in idolatry (avodah zarah) |  |  | avodah zarah, literally 'strange worship'; rendered idolatry |
| 3:3.36 | etc. |  |  | וכו' marks a truncated quotation; 'etc.' is kept |
| 3:3.37 | we learned |  |  |  |
| 3:3.38 | from |  |  |  |
| 3:3.39 | the Torah |  |  |  |
| 3:3.40 | and from our rabbis |  |  |  |
| 3:3.41 | in the tractate |  |  |  |
| 3:3.42 | Horayot |  |  | Mishnah tractate Horayot, on erroneous rulings of a court |
| 3:3.43 | that it is possible |  |  |  |
| 3:3.44 | for the Beit Din |  |  |  |
| 3:3.45 | the great |  |  |  |
| 3:3.46 | to err |  |  |  |
| 3:3.47 | and to rule |  |  |  |
| 3:3.48 | permission |  |  |  |
| 3:3.49 | in which |  |  |  |
| 3:3.50 | particular case |  |  |  |
| 3:3.51 | in idolatry |  |  |  |
| 3:3.52 | for |  |  |  |
| 3:3.53 | man |  |  |  |
| 3:3.54 | is formed |  |  |  |
| 3:3.55 | of clay |  |  |  |
| 3:3.56 | whether |  |  |  |
| 3:3.57 | great |  |  |  |
| 3:3.58 | or |  |  |  |
| 3:3.59 | small |  |  |  |
| 3:3.60 | and prone |  |  |  |
| 3:3.61 | he |  |  |  |
| 3:3.62 | to err |  |  |  |
| 3:3.63 | therefore |  |  |  |
| 3:3.64 | commanded |  |  |  |
| 3:3.65 | the Torah |  |  |  |
| 3:3.66 | that |  |  |  |
| 3:3.67 | is too hard |  |  |  |
| 3:3.68 | for you |  |  |  |
| 3:3.69 | a matter |  |  |  |
| 3:3.70 | for judgment |  |  |  |
| 3:3.71 | etc. |  |  | וכו' truncates the verse; only the words on the page are given |
| 3:3.72 | and you shall arise |  |  |  |
| 3:3.73 | and go up |  |  |  |
| 3:3.74 | to |  |  |  |
| 3:3.75 | the place |  |  |  |
| 3:3.76 | which |  |  |  |
| 3:3.77 | will choose |  |  |  |
| 3:3.78 | the LORD |  |  | ה׳, the Tetragrammaton as written (the ' is punctuation) |
| 3:3.79 | your God |  |  |  |
| 3:3.80 | in order |  |  |  |
| 3:3.81 | to lessen |  |  |  |
| 3:3.82 | the error |  |  |  |
| 3:3.83 | how |  |  |  |
| 3:3.84 | is it not |  |  |  |
| 3:3.85 | avoidable |  |  |  |
| 3:3.86 | that not |  |  |  |
| 3:3.87 | they would err |  |  |  |
| 3:3.88 | even |  |  | אפי' = אפילו |
| 3:3.89 | Beit Din |  |  |  |
| 3:3.90 | the great |  |  |  |
| 3:3.91 | for behold |  |  |  |
| 3:3.92 | it is said |  |  |  |
| 3:3.93 | in the Torah |  |  |  |
| 3:3.94 | and was hidden |  |  |  |
| 3:3.95 | a matter |  |  |  |
| 3:3.96 | from the eyes of |  |  |  |
| 3:3.97 | the assembly |  |  |  |
| 3:3.98 | and it shall be |  |  |  |
| 3:3.99 | if |  |  |  |
| 3:3.100 | from the eyes of |  |  |  |
| 3:3.101 | the congregation |  |  |  |
| 3:3.102 | was done |  |  |  |
| 3:3.103 | in error |  |  |  |
| 3:3.104 | and according to |  |  |  |
| 3:3.105 | the shortness |  |  |  |
| 3:3.106 | of our understanding |  |  |  |
| 3:3.107 | when we see |  |  |  |
| 3:3.108 | the difference |  |  |  |
| 3:3.109 | great |  |  |  |
| 3:3.110 | that there is |  |  |  |
| 3:3.111 | between |  |  |  |
| 3:3.112 | our rabbis |  |  |  |
| 3:3.113 | the early ones |  |  |  |
| 3:3.114 | and how |  |  |  |
| 3:3.115 | was |  |  |  |
| 3:3.116 | their faith |  |  |  |
| 3:3.117 | in the matter of |  |  |  |
| 3:3.118 | the divinity |  |  |  |
| 3:3.119 | and the unity |  |  |  |
| 3:3.120 | the Name, may He be blessed |  |  | השי"ת: the Name, may He be blessed (fixed form) |
| 3:3.121 | according to |  |  |  |
| 3:3.122 | the mouth of |  |  |  |
| 3:3.123 | the holy Torah |  |  | תוה"ק: our holy Torah |
| 3:3.124 | as explained |  |  |  |
| 3:3.125 | Rabbeinu (our rabbi) |  |  | ב'ר = Rabbeinu, fixed form 'Rabbeinu Bachya' |
| 3:3.126 | Bachya |  |  |  |
| 3:3.127 | the author of |  |  |  |
| 3:3.128 | Chovot |  |  |  |
| 3:3.129 | ha-Levavot |  |  | Chovot ha-Levavot, 'Duties of the Heart', the book of Rabbeinu Bachya |
| 3:3.130 | and Rabbi |  |  |  |
| 3:3.131 | Yehuda |  |  |  |
| 3:3.132 | Halevi |  |  |  |
| 3:3.133 | in the book |  |  |  |
| 3:3.134 | the Kuzari |  |  | Sefer ha-Kuzari, Halevi's book, 'The Kuzari' |
| 3:3.135 | and in his prayers |  |  |  |
| 3:3.136 | and Rabbi |  |  |  |
| 3:3.137 | Saadia |  |  |  |
| 3:3.138 | Gaon |  |  |  |
| 3:3.139 | in the book |  |  |  |
| 3:3.140 | the Emunot |  |  | Emunot ve-De'ot, 'Beliefs and Opinions', Saadia Gaon's book |
| 3:3.141 | and the De'ot |  |  |  |
| 3:3.142 | and the Rambam |  |  | the Rambam (Maimonides); the work is not named at this point |
| 3:3.143 | in his work |  |  |  |
| 3:3.144 | and in the commentary |  |  |  |
| 3:3.145 | of the Mishnah |  |  | the Rambam's commentary on the Mishnah |
| 3:3.146 | and in the book |  |  |  |
| 3:3.147 | Moreh |  |  |  |
| 3:3.148 | Nevukhim |  |  | Moreh Nevukhim, the Guide for the Perplexed |
| 3:3.149 | and the Roke'ach |  |  | Roke'ach: fixed form |
| 3:3.150 | in the gate |  |  |  |
| 3:3.151 | of unity |  |  | Shaar ha-Yichud, 'Gate of Unity', a part of his book |
| 3:3.152 | and the book of |  |  |  |
| 3:3.153 | Mitzvot |  |  |  |
| 3:3.154 | the great |  |  | Sefer Mitzvot Gadol, the Semag |
| 3:3.155 | and Rabbi |  |  |  |
| 3:3.156 | Yosef |  |  |  |
| 3:3.157 | Albo |  |  |  |
| 3:3.158 | in the book |  |  |  |
| 3:3.159 | the Ikkarim |  |  | Sefer ha-Ikkarim, 'Principles', Albo's book |
| 3:3.160 | and Rabbi |  |  |  |
| 3:3.161 | Meir |  |  |  |
| 3:3.162 | Aldabi |  |  |  |
| 3:3.163 | in the book |  |  |  |
| 3:3.164 | Shvilei |  |  | Shvilei Emunah, 'Paths of Faith', the book of Meir Aldabi |
| 3:3.165 | Emunah |  |  |  |
| 3:3.166 | and others |  |  |  |
| 3:3.167 | many |  |  |  |
| 3:3.168 | and so too |  |  |  |
| 3:3.169 | the complete ones |  |  | שלימים, 'the complete ones'; the author does not name them |
| 3:3.170 | and how |  |  |  |
| 3:3.171 | it |  |  |  |
| 3:3.172 | now |  |  |  |
| 3:3.173 | in the books of |  |  |  |
| 3:3.174 | our rabbis |  |  |  |
| 3:3.175 | the later ones |  |  |  |
| 3:3.176 | according to |  |  | ע"פ = על פי, 'according to' |
| 3:3.177 | the Kabbalah |  |  |  |
| 3:3.178 | the new |  |  |  |
| 3:3.179 | drawn out |  |  |  |
| 3:3.180 | from the start of |  |  |  |
| 3:3.181 | the thousand |  |  |  |
| 3:3.182 | sixth |  |  | the sixth thousand (sixth millennium) |
| 3:3.183 | with cords of |  |  |  |
| 3:3.184 | falsehood |  |  | Isa 5:18, embedded allusion (no quotation marks) |
| 3:3.185 | until |  |  |  |
| 3:3.186 | it became |  |  |  |
| 3:3.187 | in our time |  |  |  |
| 3:3.188 | as with |  |  |  |
| 3:3.189 | cart-ropes |  |  | Isa 5:18, embedded allusion (no quotation marks) |
| 3:3.190 | and they believed |  |  |  |
| 3:3.191 | in gods |  |  | אלהים רבים: 'gods' in the author's own charge |
| 3:3.192 | many |  |  |  |
| 3:3.193 | and that the principal |  |  |  |
| 3:3.194 | in them |  |  |  |
| 3:3.195 | which |  |  |  |
| 3:3.196 | to it |  |  |  |
| 3:3.197 | all |  |  |  |
| 3:3.198 | our worship |  |  |  |
| 3:3.199 | and our blessings |  |  |  |
| 3:3.200 | is |  |  |  |
| 3:3.201 | the partzuf |  |  |  |
| 3:3.202 | the last |  |  |  |
| 3:3.203 | in Emanation |  |  | באצילות: 'in Emanation', without 'World of' |
| 3:3.204 | called |  |  |  |
| 3:3.205 | quick (to anger) |  |  | קצר אפים: the author's calque of Ze'ir Anpin, set against 'slow to anger' (ארך אפים) in Exod 34:6; cf. Prov 14:17 |
| 3:3.206 | to anger |  |  |  |
| 3:3.207 | Ze'ir |  |  | Aramaic 'small' (ze'ir) |
| 3:3.208 | Anpin |  |  | Aramaic 'face' (anpin) |

Phrases:
- 3:3.1 | 4) | 4)
- 3:3.2 3:3.3 3:3.4 3:3.5 3:3.6 3:3.7 3:3.8 3:3.9 3:3.10 | The authority of our rabbis, the masters of the new Kabbalah, shall not be greater than the authority of | and-not will-be authority of-our-rabbis masters-of the-Kabbalah the-new greater than-the-authority-of
- 3:3.11 3:3.12 3:3.13 3:3.14 3:3.15 | the Great Beit Din who sat in the Chamber of Hewn Stone, | Beit-Din the-great who-sit in-the-chamber-of the-hewn-stone
- 3:3.16 3:3.17 3:3.18 3:3.19 3:3.20 3:3.21 | and their grasp is no greater than the grasp of the Great Sanhedrin! | and-not their-grasp greater than-the-grasp-of Sanhedrin the-great!
- 3:3.22 3:3.23 3:3.24 3:3.25 3:3.26 3:3.27 3:3.28 3:3.29 3:3.30 3:3.31 3:3.32 3:3.33 3:3.34 3:3.35 3:3.36 | And yet the Torah said that if they erred and ruled to permit the forbidden, or sinned inadvertently in idolatry, etc. | and-yet all this said the-Torah that-if erred and-ruled to-permit the forbidden or sinned-inadvertently in-idolatry etc.
- 3:3.37 3:3.38 3:3.39 3:3.40 3:3.41 3:3.42 3:3.43 3:3.44 3:3.45 3:3.46 3:3.47 3:3.48 3:3.49 3:3.50 3:3.51 | We have learned from the Torah and from our rabbis, in the tractate Horayot, that the Great Beit Din can err and rule a permission in some particular case of idolatry. | we-learned from the-Torah and-from our-rabbis in the-tractate Horayot that-possible for-the-Beit-Din the-great to-err and-to-rule permission in-which particular-case in-idolatry.
- 3:3.52 3:3.53 3:3.54 3:3.55 3:3.56 3:3.57 3:3.58 3:3.59 3:3.60 3:3.61 3:3.62 3:3.63 3:3.64 3:3.65 | For man is formed of clay, whether great or small, and is prone to err; therefore the Torah commanded: | for man is-formed of-clay whether great or small and-prone he to-err therefore commanded the-Torah
- 3:3.66 3:3.67 3:3.68 3:3.69 3:3.70 3:3.71 | "If a matter is too hard for you in judgment, etc. | that is-too-hard for-you a-matter for-judgment etc.
- 3:3.72 3:3.73 3:3.74 3:3.75 3:3.76 3:3.77 3:3.78 3:3.79 | Then arise and go up to the place that the LORD your God will choose," | and-arise and-go-up to the-place which will-choose the-LORD your-God
- 3:3.80 3:3.81 3:3.82 | in order to lessen the error. | in-order to-lessen the-error
- 3:3.83 3:3.84 3:3.85 3:3.86 3:3.87 3:3.88 3:3.89 3:3.90 | How is it not unavoidable that even the Great Beit Din should err. | how is-it-not avoidable that-not they-would-err even Beit-Din the-great
- 3:3.91 3:3.92 3:3.93 3:3.94 3:3.95 3:3.96 3:3.97 | For it is said in the Torah, "and the matter was hidden from the eyes of the assembly. | for-behold it-is-said in-the-Torah and-was-hidden a-matter from-the-eyes-of the-assembly
- 3:3.98 3:3.99 3:3.100 3:3.101 3:3.102 3:3.103 | And it shall be, if from the eyes of the congregation it was done in error." | and-it-shall-be if from-the-eyes-of the-congregation was-done in-error
- 3:3.104 3:3.105 3:3.106 3:3.107 3:3.108 3:3.109 3:3.110 3:3.111 3:3.112 3:3.113 | And according to the shortness of our understanding, when we see the great difference between our early rabbis, | and-according-to the-shortness-of our-understanding when-we-see the-difference the-great that-there-is between our-rabbis the-early
- 3:3.114 3:3.115 3:3.116 3:3.117 3:3.118 3:3.119 3:3.120 | and how was their faith in the matter of divinity and the unity of the Name, may He be blessed, | and-how was their-faith in-the-matter-of the-divinity and-the-unity the-Name-may-He-be-blessed
- 3:3.121 3:3.122 3:3.123 3:3.124 3:3.125 3:3.126 3:3.127 3:3.128 3:3.129 | according to the holy Torah, as explained by Rabbeinu Bachya, the author of Chovot ha-Levavot, | according-to the-mouth-of the-holy-Torah as-explained by-Rabbeinu Bachya the-author-of Chovot ha-Levavot
- 3:3.130 3:3.131 3:3.132 3:3.133 3:3.134 3:3.135 | and Rabbi Yehuda Halevi in the Kuzari and in his prayers, | and-Rabbi Yehuda Halevi in-the-book the-Kuzari and-in-his-prayers
- 3:3.136 3:3.137 3:3.138 3:3.139 3:3.140 3:3.141 | and Rabbi Saadia Gaon in the Emunot ve-De'ot, | and-Rabbi Saadia Gaon in-the-book the-Emunot and-the-De'ot
- 3:3.142 3:3.143 3:3.144 3:3.145 | and the Rambam in his work, and in his commentary on the Mishnah, | and-the-Rambam in-his-work and-in-the-commentary of-the-Mishnah
- 3:3.146 3:3.147 3:3.148 | and in the Moreh Nevukhim. | and-in-the-book Moreh Nevukhim
- 3:3.149 3:3.150 3:3.151 | And the Roke'ach in the Gate of Unity, | and-the-Roke'ach in-the-gate of-unity
- 3:3.152 3:3.153 3:3.154 | and the Sefer Mitzvot Gadol, | and-the-book-of Mitzvot the-great
- 3:3.155 3:3.156 3:3.157 3:3.158 3:3.159 | and Rabbi Yosef Albo in the Sefer ha-Ikkarim, | and-Rabbi Yosef Albo in-the-book the-Ikkarim
- 3:3.160 3:3.161 3:3.162 3:3.163 3:3.164 3:3.165 | and Rabbi Meir Aldabi in the Shvilei Emunah. | and-Rabbi Meir Aldabi in-the-book Shvilei Emunah
- 3:3.166 3:3.167 3:3.168 3:3.169 | And many others, and likewise the complete ones. | and-others many and-so-too the-complete-ones
- 3:3.170 3:3.171 3:3.172 3:3.173 3:3.174 3:3.175 3:3.176 3:3.177 3:3.178 | And how is it now in the books of our later rabbis, according to the new Kabbalah, | and-how it now in-the-books-of our-rabbis the-later according-to the-Kabbalah the-new
- 3:3.179 3:3.180 3:3.181 3:3.182 3:3.183 3:3.184 | which runs from the start of the sixth millennium with cords of falsehood | drawn-out from-the-start-of the-thousand the-sixth with-cords-of falsehood
- 3:3.185 3:3.186 3:3.187 3:3.188 3:3.189 3:3.190 3:3.191 3:3.192 | until it became in our time as with cart-ropes, and they believed in many gods. | until it-became in-our-time as-with cart-ropes and-they-believed in-gods many
- 3:3.193 3:3.194 3:3.195 3:3.196 3:3.197 3:3.198 3:3.199 3:3.200 3:3.201 3:3.202 3:3.203 3:3.204 3:3.205 3:3.206 3:3.207 3:3.208 | And that the principal point in them, to which all our worship and blessings are directed, is the last partzuf in Emanation, called (quick to anger) Ze'ir Anpin. | and-that the-principal in-them to-which all our-worship and-our-blessings is the-partzuf the-last in-Emanation called quick to-anger Ze'ir Anpin

Grammar notes:
- ולא יהא כח ... גדול מכח: a negated future (יהא) with a comparison; the English keeps the repeated 'the authority of' so that the comparison is plain.
- איך אינו נמנע שלא יטעו: a rhetorical question in form with a full stop in the Hebrew; the English keeps it as a question of argument ('How is it not unavoidable...').
- Three Torah passages are cited: Deut 17:8 (quoted, with וכו' as the author's truncation sign), Lev 4:13 and Num 15:24 (introduced by שהרי נאמר בתורה and continued by והיה).
- The long period ולפי קוצר דעתינו ... is divided at 'when we see' and 'how was their faith'; no clause is dropped and the order of the list of authors is kept.

Translator's note: Plain sentences follow the Hebrew full stops. The Torah quotations take double quotation marks; the author's וכו' is kept as 'etc.'. The list of authors is elliptical in the Hebrew after the full stop at הנבוכים, and the English keeps that ellipsis.

Commentary rulings:
- apparatus: Base text: Wikisource; not yet collated with the 1931 scan (open question 2), so the English is the base text as it stands.
- apparatus: Printed list marker ד) (the fourth item) opens the paragraph; the next paragraph opens ה) (the fifth item).
- apparatus: Deut 17:8, quoted with 'etc.' marking the author's omission (וכו').
- apparatus: Lev 4:13, quoted: ונעלם דבר מעיני הקהל.
- apparatus: Num 15:24, quoted: והיה אם מעיני העדה נעשתה לשגגה.
- apparatus: Isa 5:18, embedded allusion (no quotation marks): בחבלי השוא and כעבות העגלה.
- apparatus: Prov 14:17, for the author's קצר אפים, set against ארך אפים in Exod 34:6.

## Milchamot ha-Shem 3:4

Hebrew (Wikisource): ה) ואמונה זו החדשה רחוקה מאד מדרך תורתינו הק' שבכתב ושבעל פה. ולכן פירשנו עצמינו ממנה ודבקנו בדברי רבותינו ז"ל בעלי המשנה והתלמוד בבלי וירושלמי והפוסקים ומדרשי רבותינו המפורסמים, מדרש רבה ותנחומא וילקוט שמעוני ומדרש הגדול המצוי בתימן בכת"י, ונור אלצלאם וכיוצא בהם. גם שארית חבורי קדמונינו חכמי תימן מצינו וראינו שהם אדוקים בדברי הגדולים הנזכרים לעיל ובדעותיהם. וכראותכם שפירשנו מן הקבלה החדשה חרה אפכם בנו, וברמות רוחא טפלתם עלינו שקר זדים. ושאלנו אתכם להורותינו הדרך איך יסכימו דברי רבותינו המקובלים החדשים עם דברי רבותינו הראשונים הנזכרים לעיל, ועם דברי התנאים והאמוראים. והיתה תשובתכם אלינו בדרך הטעאה וקינטורים.

| id | word | punctuation after |
|---|---|---|
| 3:4.1 | ה | ) |
| 3:4.2 | ואמונה |  |
| 3:4.3 | זו |  |
| 3:4.4 | החדשה |  |
| 3:4.5 | רחוקה |  |
| 3:4.6 | מאד |  |
| 3:4.7 | מדרך |  |
| 3:4.8 | תורתינו |  |
| 3:4.9 | הק | ' |
| 3:4.10 | שבכתב |  |
| 3:4.11 | ושבעל |  |
| 3:4.12 | פה | . |
| 3:4.13 | ולכן |  |
| 3:4.14 | פירשנו |  |
| 3:4.15 | עצמינו |  |
| 3:4.16 | ממנה |  |
| 3:4.17 | ודבקנו |  |
| 3:4.18 | בדברי |  |
| 3:4.19 | רבותינו |  |
| 3:4.20 | ז"ל |  |
| 3:4.21 | בעלי |  |
| 3:4.22 | המשנה |  |
| 3:4.23 | והתלמוד |  |
| 3:4.24 | בבלי |  |
| 3:4.25 | וירושלמי |  |
| 3:4.26 | והפוסקים |  |
| 3:4.27 | ומדרשי |  |
| 3:4.28 | רבותינו |  |
| 3:4.29 | המפורסמים | , |
| 3:4.30 | מדרש |  |
| 3:4.31 | רבה |  |
| 3:4.32 | ותנחומא |  |
| 3:4.33 | וילקוט |  |
| 3:4.34 | שמעוני |  |
| 3:4.35 | ומדרש |  |
| 3:4.36 | הגדול |  |
| 3:4.37 | המצוי |  |
| 3:4.38 | בתימן |  |
| 3:4.39 | בכת"י | , |
| 3:4.40 | ונור |  |
| 3:4.41 | אלצלאם |  |
| 3:4.42 | וכיוצא |  |
| 3:4.43 | בהם | . |
| 3:4.44 | גם |  |
| 3:4.45 | שארית |  |
| 3:4.46 | חבורי |  |
| 3:4.47 | קדמונינו |  |
| 3:4.48 | חכמי |  |
| 3:4.49 | תימן |  |
| 3:4.50 | מצינו |  |
| 3:4.51 | וראינו |  |
| 3:4.52 | שהם |  |
| 3:4.53 | אדוקים |  |
| 3:4.54 | בדברי |  |
| 3:4.55 | הגדולים |  |
| 3:4.56 | הנזכרים |  |
| 3:4.57 | לעיל |  |
| 3:4.58 | ובדעותיהם | . |
| 3:4.59 | וכראותכם |  |
| 3:4.60 | שפירשנו |  |
| 3:4.61 | מן |  |
| 3:4.62 | הקבלה |  |
| 3:4.63 | החדשה |  |
| 3:4.64 | חרה |  |
| 3:4.65 | אפכם |  |
| 3:4.66 | בנו | , |
| 3:4.67 | וברמות |  |
| 3:4.68 | רוחא |  |
| 3:4.69 | טפלתם |  |
| 3:4.70 | עלינו |  |
| 3:4.71 | שקר |  |
| 3:4.72 | זדים | . |
| 3:4.73 | ושאלנו |  |
| 3:4.74 | אתכם |  |
| 3:4.75 | להורותינו |  |
| 3:4.76 | הדרך |  |
| 3:4.77 | איך |  |
| 3:4.78 | יסכימו |  |
| 3:4.79 | דברי |  |
| 3:4.80 | רבותינו |  |
| 3:4.81 | המקובלים |  |
| 3:4.82 | החדשים |  |
| 3:4.83 | עם |  |
| 3:4.84 | דברי |  |
| 3:4.85 | רבותינו |  |
| 3:4.86 | הראשונים |  |
| 3:4.87 | הנזכרים |  |
| 3:4.88 | לעיל | , |
| 3:4.89 | ועם |  |
| 3:4.90 | דברי |  |
| 3:4.91 | התנאים |  |
| 3:4.92 | והאמוראים | . |
| 3:4.93 | והיתה |  |
| 3:4.94 | תשובתכם |  |
| 3:4.95 | אלינו |  |
| 3:4.96 | בדרך |  |
| 3:4.97 | הטעאה |  |
| 3:4.98 | וקינטורים | . |


### Translation under review (3:4)

English: 5) And this new faith is very far from the way of our holy Torah, written and oral. Therefore we separated ourselves from it, and we clung to the words of our rabbis of blessed memory, the masters of the Mishnah and the Talmud, Bavli and Yerushalmi, and the decisors, and the midrashim of our famous rabbis, Midrash Rabbah and Tanchuma and Yalkut Shimoni, and the great Midrash found in Yemen in manuscript, and Nur al-Salam, and the like. Also, the remnant of the books of our forebears, the sages of Yemen, we find and see that they adhere to the words of the great ones mentioned above and to their opinions. And as you see that we separated from the new Kabbalah, your anger burned against us, and in haughtiness of spirit you ascribed to us arrogant falsehood. And we asked you to teach us the way: how would the words of our rabbis, the new Kabbalists, agree with the words of our early rabbis mentioned above, and with the words of the Tannaim and the Amoraim. And your answer to us was in the manner of deception and quarrels.

Literal: 5) and-the-faith this the-new far very from-the-way-of our-Torah the-holy that-written and-that-oral therefore we-separated ourselves from-it and-we-clung to-the-words-of our-rabbis of-blessed-memory the-masters-of the-Mishnah and-the-Talmud Bavli and-Yerushalmi and-the-decisors and-the-midrashim-of our-rabbis the-famous Midrash Rabbah and-Tanchuma and-Yalkut Shimoni and-Midrash the-great found in-Yemen in-manuscript and-Nur al-Salam and-the-like of-them also the-remnant-of the-books-of our-forebears the-sages-of Yemen we-find and-we-see that-they adhere to-the-words-of the-great-ones mentioned above and-to-their-opinions and-as-you-see that-we-separated from the-Kabbalah the-new burned your-anger against-us and-in haughtiness-of spirit you-ascribed to-us falsehood arrogant and-we-asked of-you to-teach-us the-way how would-agree the-words-of our-rabbis the-Kabbalists the-new with the-words-of our-rabbis the-early mentioned above and-with the-words-of the-Tannaim and-the-Amoraim and-was your-answer to-us in-the-manner-of deception and-quarrels

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:4.1 | 5) |  |  | printed list marker ה) (the fifth item); kept as 5) |
| 3:4.2 | and the faith |  |  |  |
| 3:4.3 | this |  |  |  |
| 3:4.4 | new |  |  |  |
| 3:4.5 | far |  |  |  |
| 3:4.6 | very |  |  |  |
| 3:4.7 | from the way of |  |  |  |
| 3:4.8 | our Torah |  |  |  |
| 3:4.9 | the holy |  |  | הק' = הקדושה, 'holy' |
| 3:4.10 | that is written |  |  |  |
| 3:4.11 | and that is oral |  |  |  |
| 3:4.12 | oral (Torah) |  |  |  |
| 3:4.13 | therefore |  |  |  |
| 3:4.14 | we separated |  |  |  |
| 3:4.15 | ourselves |  |  |  |
| 3:4.16 | from it |  |  |  |
| 3:4.17 | and we clung |  |  |  |
| 3:4.18 | to the words of |  |  |  |
| 3:4.19 | our rabbis |  |  |  |
| 3:4.20 | of blessed memory |  |  | ז"ל: 'of blessed memory' |
| 3:4.21 | the masters of |  |  |  |
| 3:4.22 | the Mishnah |  |  |  |
| 3:4.23 | and the Talmud |  |  |  |
| 3:4.24 | Bavli |  |  | the Babylonian Talmud |
| 3:4.25 | and Yerushalmi |  |  | the Jerusalem Talmud |
| 3:4.26 | and the decisors |  |  | הפוסקים: the halakhic decisors |
| 3:4.27 | and the midrashim of |  |  |  |
| 3:4.28 | our rabbis |  |  |  |
| 3:4.29 | the famous |  |  |  |
| 3:4.30 | Midrash |  |  |  |
| 3:4.31 | Rabbah |  |  |  |
| 3:4.32 | and Tanchuma |  |  |  |
| 3:4.33 | and Yalkut |  |  |  |
| 3:4.34 | Shimoni |  |  |  |
| 3:4.35 | and the Midrash |  |  |  |
| 3:4.36 | the great |  |  | Midrash ha-Gadol |
| 3:4.37 | found |  |  |  |
| 3:4.38 | in Yemen |  |  |  |
| 3:4.39 | in manuscript |  |  | בכת"י = ketav yad, manuscript |
| 3:4.40 | and Nur |  |  | Nur al-Salam: the title as the author gives it (Arabic, 'light of peace') |
| 3:4.41 | al-Salam |  |  |  |
| 3:4.42 | and the like |  |  |  |
| 3:4.43 | of them |  |  |  |
| 3:4.44 | also |  |  |  |
| 3:4.45 | the remnant of |  |  |  |
| 3:4.46 | the books of |  |  |  |
| 3:4.47 | our forebears |  |  |  |
| 3:4.48 | the sages of |  |  |  |
| 3:4.49 | Yemen |  |  |  |
| 3:4.50 | we find |  |  |  |
| 3:4.51 | and we see |  |  |  |
| 3:4.52 | that they |  |  |  |
| 3:4.53 | adhere |  |  |  |
| 3:4.54 | to the words of |  |  |  |
| 3:4.55 | the great ones |  |  |  |
| 3:4.56 | mentioned |  |  |  |
| 3:4.57 | above |  |  |  |
| 3:4.58 | and to their opinions |  |  |  |
| 3:4.59 | and as you see |  |  |  |
| 3:4.60 | that we separated |  |  |  |
| 3:4.61 | from |  |  |  |
| 3:4.62 | the Kabbalah |  |  |  |
| 3:4.63 | the new |  |  |  |
| 3:4.64 | burned |  |  |  |
| 3:4.65 | your anger |  |  |  |
| 3:4.66 | against us |  |  |  |
| 3:4.67 | and in |  |  |  |
| 3:4.68 | haughtiness of spirit |  |  | Aramaic idiom ברמות רוחא, 'in haughtiness of spirit' |
| 3:4.69 | you ascribed |  |  |  |
| 3:4.70 | to us |  |  |  |
| 3:4.71 | falsehood |  |  |  |
| 3:4.72 | arrogant |  |  | זדים: 'arrogant' |
| 3:4.73 | and we asked |  |  |  |
| 3:4.74 | of you |  |  |  |
| 3:4.75 | to teach us |  |  |  |
| 3:4.76 | the way |  |  |  |
| 3:4.77 | how |  |  |  |
| 3:4.78 | would agree |  |  |  |
| 3:4.79 | the words of |  |  |  |
| 3:4.80 | our rabbis |  |  |  |
| 3:4.81 | the Kabbalists |  |  |  |
| 3:4.82 | the new |  |  |  |
| 3:4.83 | with |  |  |  |
| 3:4.84 | the words of |  |  |  |
| 3:4.85 | our rabbis |  |  |  |
| 3:4.86 | the early ones |  |  |  |
| 3:4.87 | mentioned |  |  |  |
| 3:4.88 | above |  |  |  |
| 3:4.89 | and with |  |  |  |
| 3:4.90 | the words of |  |  |  |
| 3:4.91 | the Tannaim |  |  |  |
| 3:4.92 | and the Amoraim |  |  |  |
| 3:4.93 | and was |  |  |  |
| 3:4.94 | your answer |  |  |  |
| 3:4.95 | to us |  |  |  |
| 3:4.96 | in the manner of |  |  |  |
| 3:4.97 | deception |  |  | הטעאה: 'deception' |
| 3:4.98 | and quarrels |  |  |  |

Phrases:
- 3:4.1 | 5) | 5)
- 3:4.2 3:4.3 3:4.4 3:4.5 3:4.6 3:4.7 3:4.8 3:4.9 3:4.10 3:4.11 3:4.12 | And this new faith is very far from the way of our holy Torah, written and oral. | and-the-faith this the-new far very from-the-way-of our-Torah the-holy that-written and-that-oral
- 3:4.13 3:4.14 3:4.15 3:4.16 | Therefore we separated ourselves from it, | therefore we-separated ourselves from-it
- 3:4.17 3:4.18 3:4.19 3:4.20 3:4.21 | and we clung to the words of our rabbis of blessed memory, the masters of | and-we-clung to-the-words-of our-rabbis of-blessed-memory the-masters-of
- 3:4.22 3:4.23 3:4.24 3:4.25 3:4.26 | the Mishnah and the Talmud, Bavli and Yerushalmi, and the decisors, | the-Mishnah and-the-Talmud Bavli and-Yerushalmi and-the-decisors
- 3:4.27 3:4.28 3:4.29 3:4.30 3:4.31 3:4.32 3:4.33 3:4.34 | and the midrashim of our famous rabbis, Midrash Rabbah and Tanchuma and Yalkut Shimoni, | and-the-midrashim-of our-rabbis the-famous Midrash Rabbah and-Tanchuma and-Yalkut Shimoni
- 3:4.35 3:4.36 3:4.37 3:4.38 3:4.39 3:4.40 3:4.41 3:4.42 3:4.43 | and the great Midrash found in Yemen in manuscript, and Nur al-Salam, and the like. | and-Midrash the-great found in-Yemen in-manuscript and-Nur al-Salam and-the-like of-them
- 3:4.44 3:4.45 3:4.46 3:4.47 3:4.48 3:4.49 3:4.50 3:4.51 3:4.52 3:4.53 | Also, the remnant of the books of our forebears, the sages of Yemen, we find and see that they adhere | also the-remnant-of the-books-of our-forebears the-sages-of Yemen we-find and-we-see that-they adhere
- 3:4.54 3:4.55 3:4.56 3:4.57 3:4.58 | to the words of the great ones mentioned above and to their opinions. | to-the-words-of the-great-ones mentioned above and-to-their-opinions
- 3:4.59 3:4.60 3:4.61 3:4.62 3:4.63 3:4.64 3:4.65 3:4.66 | And as you see that we separated from the new Kabbalah, your anger burned against us, | and-as-you-see that-we-separated from the-Kabbalah the-new burned your-anger against-us
- 3:4.67 3:4.68 3:4.69 3:4.70 3:4.71 3:4.72 | and in haughtiness of spirit you ascribed to us arrogant falsehood. | and-in haughtiness-of spirit you-ascribed to-us falsehood arrogant
- 3:4.73 3:4.74 3:4.75 3:4.76 3:4.77 | And we asked you to teach us the way: how | and-we-asked of-you to-teach-us the-way how
- 3:4.78 3:4.79 3:4.80 3:4.81 3:4.82 3:4.83 3:4.84 3:4.85 3:4.86 3:4.87 3:4.88 | would the words of our rabbis, the new Kabbalists, agree with the words of our early rabbis mentioned above, | would-agree the-words-of our-rabbis the-Kabbalists the-new with the-words-of our-rabbis the-early mentioned above
- 3:4.89 3:4.90 3:4.91 3:4.92 | and with the words of the Tannaim and the Amoraim. | and-with the-words-of the-Tannaim and-the-Amoraim
- 3:4.93 3:4.94 3:4.95 3:4.96 3:4.97 3:4.98 | And your answer to us was in the manner of deception and quarrels. | and-was your-answer to-us in-the-manner-of deception and-quarrels

Grammar notes:
- The paragraph is a chain of first-person plural past acts (פירשנו, ודבקנו, שאלנו); the reply of the other side is in the perfect (והיתה) and is given as a report.
- ושאלנו אתכם ... איך יסכימו: the question is indirect (איך with the imperfect) and carries a full stop; the English keeps it as a report, 'how would ... agree'.
- ברמות רוחא is an Aramaic idiom for arrogance and is translated by its sense; the literal is in the phrase literal.

Translator's note: The list of sources is kept in the order the author gives it. The Midrash ha-Gadol and Nur al-Salam are named only by their description and title in the Hebrew.

Commentary rulings:

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

