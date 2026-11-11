# Translation pack: milhamot-hashem perek-10-02 (Milchamot ha-Shem 10:2–10:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-02`.

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
- ולא אמרינן כלא חד (1:7) and באגרונו (1:12) are for the owner: the main
  English stands as translated; the reviewer's alternative is in `readings`
  / `apparatus`.

## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
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


## Milchamot ha-Shem 10:2

Hebrew (Wikisource): קמא) על כן יאמר כל מי שנגע יראת השי"ת ירחיק עצמו כמטחוי קשת מכל המנהגים ודינים מחודשים על פי הקבלה החדשה השקרית אשר מקרוב באה ונתפשטה בארץ התימן כעדות מהר"י צאחרי בספר צידה לדרך, לפי שכולם עוקרים וחורסים פנות תורה ה"ק ועמודיה יתפלצון. כי בכל ברכותיהם ותפלותיהם כשמזכירים את השם הנכבד והנורא הם מכוונים לאל אחר קצר אפים (ז"א) בשתוף שאר הפרצופים שלמעלה ממנו כמבואר בספריהם במקומות אין מספר. ולא מהמנהגים שנתחדשו על פי הקבלה הנכרית החדשה בלבד צריך להזהר. אלא גם מכל הדברים שהחזירונו רז"ל. והא לך קצת מהם ספר תורה ותפלין ומזוזה שנכתבו על ידי ישראל האדוק בקבלה הלזו פסולים כי השמות שבהם נתקדשו לשם אל אחר קצר אפים (ז"א) והרי כולן כשמות האמורות במיכה שהן חול כמו שכ' הרמב"ם בסוף פרק ששי מהלכות יסודי התורה וקרוב לומר שדין הוא שישרפו כמו שכתבו הפוסקים שס"ת ותפלין ומזוזה שכתבן מין ישרפו לא יאכלו בשר משחיטת מקובל כי בעת שמזכיר את השם בברכתו על השחיטה דעתו על יראתו (ומטעם זה קבעו קדמונינו הי"ג עיקרים בתוך ביאורי וחידושי הלבות שחיטה, שכל הבא ליטול קבלה לשחוט לעצמו או לאחרים צריך להיות מאמין בהן ויודען היטב, ואם אינו יודען וכ"ש אם אינו מאמין בהם כמו המקובלים החדשים שאין מאמינים בארבעה עיקרים מהן לא יתנו לו רשות לשחוט וכל שכן לקדש ולגרש, וגיטו וקידושו אינן כלום, ונמצא מרבה ממזרים בישראל). אם שליח צבור אדוק בקבלה החדשה ובקי בה לא יענה אחריו אמן וקדוש וברוך ה' המבורך. שכל כוונתו לאל אחר זעיר אנפין. אם מקדש אשה או מגרשה לא יהיו העדים ממאמיני קבלה החדשה כי הקדושין והגט בטלים מפני שהם פסולי עדות דאוריתא, ונמצאת אשת איש נשאת לאחר על ידי גט בטל. וכבר הודעתיך בסי' פ"ג בשם משנת חסידים דבשעת אמירת עשר מכות בליל פסח הם מנסכים ושופכים מכוס יינם לתוך כלי שבור להקליפה הנקראת ארור כדי לתקן המלכות מדם נדתה, ונקתה ונזרעה זרע וכו' ואם כן יינם יין נסך ואסור בהנאה. ועוד הרבה עניינים המביאים לידי מינות ועבודה זרה. וצריך להזהר מהם וממאכלם ומשתיהם פו יטמאו נשמתו באיסורי מאכלות.

| id | word | punctuation after |
|---|---|---|
| 10:2.1 | קמא | ) |
| 10:2.2 | על |  |
| 10:2.3 | כן |  |
| 10:2.4 | יאמר |  |
| 10:2.5 | כל |  |
| 10:2.6 | מי |  |
| 10:2.7 | שנגע |  |
| 10:2.8 | יראת |  |
| 10:2.9 | השי"ת |  |
| 10:2.10 | ירחיק |  |
| 10:2.11 | עצמו |  |
| 10:2.12 | כמטחוי |  |
| 10:2.13 | קשת |  |
| 10:2.14 | מכל |  |
| 10:2.15 | המנהגים |  |
| 10:2.16 | ודינים |  |
| 10:2.17 | מחודשים |  |
| 10:2.18 | על |  |
| 10:2.19 | פי |  |
| 10:2.20 | הקבלה |  |
| 10:2.21 | החדשה |  |
| 10:2.22 | השקרית |  |
| 10:2.23 | אשר |  |
| 10:2.24 | מקרוב |  |
| 10:2.25 | באה |  |
| 10:2.26 | ונתפשטה |  |
| 10:2.27 | בארץ |  |
| 10:2.28 | התימן |  |
| 10:2.29 | כעדות |  |
| 10:2.30 | מהר"י |  |
| 10:2.31 | צאחרי |  |
| 10:2.32 | בספר |  |
| 10:2.33 | צידה |  |
| 10:2.34 | לדרך | , |
| 10:2.35 | לפי |  |
| 10:2.36 | שכולם |  |
| 10:2.37 | עוקרים |  |
| 10:2.38 | וחורסים |  |
| 10:2.39 | פנות |  |
| 10:2.40 | תורה |  |
| 10:2.41 | ה"ק |  |
| 10:2.42 | ועמודיה |  |
| 10:2.43 | יתפלצון | . |
| 10:2.44 | כי |  |
| 10:2.45 | בכל |  |
| 10:2.46 | ברכותיהם |  |
| 10:2.47 | ותפלותיהם |  |
| 10:2.48 | כשמזכירים |  |
| 10:2.49 | את |  |
| 10:2.50 | השם |  |
| 10:2.51 | הנכבד |  |
| 10:2.52 | והנורא |  |
| 10:2.53 | הם |  |
| 10:2.54 | מכוונים |  |
| 10:2.55 | לאל |  |
| 10:2.56 | אחר |  |
| 10:2.57 | קצר |  |
| 10:2.58 | אפים |  |
| 10:2.59 | ז"א | ) |
| 10:2.60 | בשתוף |  |
| 10:2.61 | שאר |  |
| 10:2.62 | הפרצופים |  |
| 10:2.63 | שלמעלה |  |
| 10:2.64 | ממנו |  |
| 10:2.65 | כמבואר |  |
| 10:2.66 | בספריהם |  |
| 10:2.67 | במקומות |  |
| 10:2.68 | אין |  |
| 10:2.69 | מספר | . |
| 10:2.70 | ולא |  |
| 10:2.71 | מהמנהגים |  |
| 10:2.72 | שנתחדשו |  |
| 10:2.73 | על |  |
| 10:2.74 | פי |  |
| 10:2.75 | הקבלה |  |
| 10:2.76 | הנכרית |  |
| 10:2.77 | החדשה |  |
| 10:2.78 | בלבד |  |
| 10:2.79 | צריך |  |
| 10:2.80 | להזהר | . |
| 10:2.81 | אלא |  |
| 10:2.82 | גם |  |
| 10:2.83 | מכל |  |
| 10:2.84 | הדברים |  |
| 10:2.85 | שהחזירונו |  |
| 10:2.86 | רז"ל | . |
| 10:2.87 | והא |  |
| 10:2.88 | לך |  |
| 10:2.89 | קצת |  |
| 10:2.90 | מהם |  |
| 10:2.91 | ספר |  |
| 10:2.92 | תורה |  |
| 10:2.93 | ותפלין |  |
| 10:2.94 | ומזוזה |  |
| 10:2.95 | שנכתבו |  |
| 10:2.96 | על |  |
| 10:2.97 | ידי |  |
| 10:2.98 | ישראל |  |
| 10:2.99 | האדוק |  |
| 10:2.100 | בקבלה |  |
| 10:2.101 | הלזו |  |
| 10:2.102 | פסולים |  |
| 10:2.103 | כי |  |
| 10:2.104 | השמות |  |
| 10:2.105 | שבהם |  |
| 10:2.106 | נתקדשו |  |
| 10:2.107 | לשם |  |
| 10:2.108 | אל |  |
| 10:2.109 | אחר |  |
| 10:2.110 | קצר |  |
| 10:2.111 | אפים |  |
| 10:2.112 | ז"א | ) |
| 10:2.113 | והרי |  |
| 10:2.114 | כולן |  |
| 10:2.115 | כשמות |  |
| 10:2.116 | האמורות |  |
| 10:2.117 | במיכה |  |
| 10:2.118 | שהן |  |
| 10:2.119 | חול |  |
| 10:2.120 | כמו |  |
| 10:2.121 | שכ | ' |
| 10:2.122 | הרמב"ם |  |
| 10:2.123 | בסוף |  |
| 10:2.124 | פרק |  |
| 10:2.125 | ששי |  |
| 10:2.126 | מהלכות |  |
| 10:2.127 | יסודי |  |
| 10:2.128 | התורה |  |
| 10:2.129 | וקרוב |  |
| 10:2.130 | לומר |  |
| 10:2.131 | שדין |  |
| 10:2.132 | הוא |  |
| 10:2.133 | שישרפו |  |
| 10:2.134 | כמו |  |
| 10:2.135 | שכתבו |  |
| 10:2.136 | הפוסקים |  |
| 10:2.137 | שס"ת |  |
| 10:2.138 | ותפלין |  |
| 10:2.139 | ומזוזה |  |
| 10:2.140 | שכתבן |  |
| 10:2.141 | מין |  |
| 10:2.142 | ישרפו |  |
| 10:2.143 | לא |  |
| 10:2.144 | יאכלו |  |
| 10:2.145 | בשר |  |
| 10:2.146 | משחיטת |  |
| 10:2.147 | מקובל |  |
| 10:2.148 | כי |  |
| 10:2.149 | בעת |  |
| 10:2.150 | שמזכיר |  |
| 10:2.151 | את |  |
| 10:2.152 | השם |  |
| 10:2.153 | בברכתו |  |
| 10:2.154 | על |  |
| 10:2.155 | השחיטה |  |
| 10:2.156 | דעתו |  |
| 10:2.157 | על |  |
| 10:2.158 | יראתו |  |
| 10:2.159 | ומטעם |  |
| 10:2.160 | זה |  |
| 10:2.161 | קבעו |  |
| 10:2.162 | קדמונינו |  |
| 10:2.163 | הי"ג |  |
| 10:2.164 | עיקרים |  |
| 10:2.165 | בתוך |  |
| 10:2.166 | ביאורי |  |
| 10:2.167 | וחידושי |  |
| 10:2.168 | הלבות |  |
| 10:2.169 | שחיטה | , |
| 10:2.170 | שכל |  |
| 10:2.171 | הבא |  |
| 10:2.172 | ליטול |  |
| 10:2.173 | קבלה |  |
| 10:2.174 | לשחוט |  |
| 10:2.175 | לעצמו |  |
| 10:2.176 | או |  |
| 10:2.177 | לאחרים |  |
| 10:2.178 | צריך |  |
| 10:2.179 | להיות |  |
| 10:2.180 | מאמין |  |
| 10:2.181 | בהן |  |
| 10:2.182 | ויודען |  |
| 10:2.183 | היטב | , |
| 10:2.184 | ואם |  |
| 10:2.185 | אינו |  |
| 10:2.186 | יודען |  |
| 10:2.187 | וכ"ש |  |
| 10:2.188 | אם |  |
| 10:2.189 | אינו |  |
| 10:2.190 | מאמין |  |
| 10:2.191 | בהם |  |
| 10:2.192 | כמו |  |
| 10:2.193 | המקובלים |  |
| 10:2.194 | החדשים |  |
| 10:2.195 | שאין |  |
| 10:2.196 | מאמינים |  |
| 10:2.197 | בארבעה |  |
| 10:2.198 | עיקרים |  |
| 10:2.199 | מהן |  |
| 10:2.200 | לא |  |
| 10:2.201 | יתנו |  |
| 10:2.202 | לו |  |
| 10:2.203 | רשות |  |
| 10:2.204 | לשחוט |  |
| 10:2.205 | וכל |  |
| 10:2.206 | שכן |  |
| 10:2.207 | לקדש |  |
| 10:2.208 | ולגרש | , |
| 10:2.209 | וגיטו |  |
| 10:2.210 | וקידושו |  |
| 10:2.211 | אינן |  |
| 10:2.212 | כלום | , |
| 10:2.213 | ונמצא |  |
| 10:2.214 | מרבה |  |
| 10:2.215 | ממזרים |  |
| 10:2.216 | בישראל | ). |
| 10:2.217 | אם |  |
| 10:2.218 | שליח |  |
| 10:2.219 | צבור |  |
| 10:2.220 | אדוק |  |
| 10:2.221 | בקבלה |  |
| 10:2.222 | החדשה |  |
| 10:2.223 | ובקי |  |
| 10:2.224 | בה |  |
| 10:2.225 | לא |  |
| 10:2.226 | יענה |  |
| 10:2.227 | אחריו |  |
| 10:2.228 | אמן |  |
| 10:2.229 | וקדוש |  |
| 10:2.230 | וברוך |  |
| 10:2.231 | ה | ' |
| 10:2.232 | המבורך | . |
| 10:2.233 | שכל |  |
| 10:2.234 | כוונתו |  |
| 10:2.235 | לאל |  |
| 10:2.236 | אחר |  |
| 10:2.237 | זעיר |  |
| 10:2.238 | אנפין | . |
| 10:2.239 | אם |  |
| 10:2.240 | מקדש |  |
| 10:2.241 | אשה |  |
| 10:2.242 | או |  |
| 10:2.243 | מגרשה |  |
| 10:2.244 | לא |  |
| 10:2.245 | יהיו |  |
| 10:2.246 | העדים |  |
| 10:2.247 | ממאמיני |  |
| 10:2.248 | קבלה |  |
| 10:2.249 | החדשה |  |
| 10:2.250 | כי |  |
| 10:2.251 | הקדושין |  |
| 10:2.252 | והגט |  |
| 10:2.253 | בטלים |  |
| 10:2.254 | מפני |  |
| 10:2.255 | שהם |  |
| 10:2.256 | פסולי |  |
| 10:2.257 | עדות |  |
| 10:2.258 | דאוריתא | , |
| 10:2.259 | ונמצאת |  |
| 10:2.260 | אשת |  |
| 10:2.261 | איש |  |
| 10:2.262 | נשאת |  |
| 10:2.263 | לאחר |  |
| 10:2.264 | על |  |
| 10:2.265 | ידי |  |
| 10:2.266 | גט |  |
| 10:2.267 | בטל | . |
| 10:2.268 | וכבר |  |
| 10:2.269 | הודעתיך |  |
| 10:2.270 | בסי | ' |
| 10:2.271 | פ"ג |  |
| 10:2.272 | בשם |  |
| 10:2.273 | משנת |  |
| 10:2.274 | חסידים |  |
| 10:2.275 | דבשעת |  |
| 10:2.276 | אמירת |  |
| 10:2.277 | עשר |  |
| 10:2.278 | מכות |  |
| 10:2.279 | בליל |  |
| 10:2.280 | פסח |  |
| 10:2.281 | הם |  |
| 10:2.282 | מנסכים |  |
| 10:2.283 | ושופכים |  |
| 10:2.284 | מכוס |  |
| 10:2.285 | יינם |  |
| 10:2.286 | לתוך |  |
| 10:2.287 | כלי |  |
| 10:2.288 | שבור |  |
| 10:2.289 | להקליפה |  |
| 10:2.290 | הנקראת |  |
| 10:2.291 | ארור |  |
| 10:2.292 | כדי |  |
| 10:2.293 | לתקן |  |
| 10:2.294 | המלכות |  |
| 10:2.295 | מדם |  |
| 10:2.296 | נדתה | , |
| 10:2.297 | ונקתה |  |
| 10:2.298 | ונזרעה |  |
| 10:2.299 | זרע |  |
| 10:2.300 | וכו | ' |
| 10:2.301 | ואם |  |
| 10:2.302 | כן |  |
| 10:2.303 | יינם |  |
| 10:2.304 | יין |  |
| 10:2.305 | נסך |  |
| 10:2.306 | ואסור |  |
| 10:2.307 | בהנאה | . |
| 10:2.308 | ועוד |  |
| 10:2.309 | הרבה |  |
| 10:2.310 | עניינים |  |
| 10:2.311 | המביאים |  |
| 10:2.312 | לידי |  |
| 10:2.313 | מינות |  |
| 10:2.314 | ועבודה |  |
| 10:2.315 | זרה | . |
| 10:2.316 | וצריך |  |
| 10:2.317 | להזהר |  |
| 10:2.318 | מהם |  |
| 10:2.319 | וממאכלם |  |
| 10:2.320 | ומשתיהם |  |
| 10:2.321 | פו |  |
| 10:2.322 | יטמאו |  |
| 10:2.323 | נשמתו |  |
| 10:2.324 | באיסורי |  |
| 10:2.325 | מאכלות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

