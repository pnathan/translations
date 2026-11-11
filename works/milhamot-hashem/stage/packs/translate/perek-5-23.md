# Translation pack: milhamot-hashem perek-5-23 (Milchamot ha-Shem 5:32–5:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-23.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-23`.

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

- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
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


## Milchamot ha-Shem 5:32

Hebrew (Wikisource): נח) ובסנהדרין פ' ד' מיתות גרסינן אמר ר' יהודה אלמלא וי"ו שבהעלוך נתחייבו שונאיהם של ישראל כלייה. אמר לו רשב"י והלא כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתי לה' לבדו אלא מה ת"ל אשר העלוך שאיוו לאלוהות הרבה ע"כ, ומביאו מעשה רקח בפרק ב' מהל' ע"ז. ובפ' לולב וערבה אמתניתין דהיו אומרין ליה ולך מזבח פריך תלמודא והתניא כל המשתף שם שמים ודבר אחד נעקר מן העולם? אלא הכי קאמרי וכו' וכ' הרב מהרש"א בפ' ד' מיתות דודאי לכ"ע גרע טפי כופר בעיקר וחייב כלייה, אלא דרשב"י קאמר שכל המשתף שם שמים ודבר אחר דהיינו המאמין נמי בדבר אחר חייב כלייה ונעקר מן העולם, אלא מה ת"ל אשר העלוך דודאי האמינו בנמצא וכו' ולכך לא נעקרו אלא שאיוו לאלוהות הרבה ואמרו שהאחד הנמצא נתן כח לאלוהות וכוחות רבים עזב ה' את הארץ. ואלו הכוחות והאלוהות יהיו מנהיגים בארץ. וכן כ' הרמב"ם שבה דעת רבים מעובדי ע"ז שעובדים אותה. שאומרים שניתנה ההנהגה בידה. ועל דרך זה יש לכוון מה שאמר משה ראה אתה אומר אלי העל את העם הזה. שבא משה להתנצל בעדם. שאמר משה להקב"ה שהרי אתה אומר אלי העל את העם הזה. וכוונתך שאתה מעמיד אותי למנהיג להם, ותלית העלייה בי, כיוצא בזה כיוונו כלל לע"ז ולשתף שם שמים ודבר אחר עכ"ל וכל זה דוחק גדול. וכ' מעשה רקח דאפשר דשיתוף דקאמר היינו לכל האלוהות עם הקב"ה ע"ש ופשוט הוא. מכל זה נתבאר לך ידידי האוהב את תורת האמת ודברי דודים מעתיקיה ומקבליה ממשה רבי' איש מפי איש, שזה שאמרו המקובלים החדשים כי אלהינו בהו"ש בהיותו בבחינתו הפשוטה לא שייך בו שום עבודה ותפלה וקריאה כלל, כי קצור קצרה ידו ח"ו מפדות ואין בו כח להציל את הקורא ומשוע אליו ורק בהיותו משתתף בפרצופים הנז' וביחוד פרצוף זעיר אנפין אז שייך בו עבודה וברכה. וצריך דוקא לאחדו ולשתפו עם נוקבא דילה ועם שאר החמשה פרצופים, ואז בהיותו משתתף עמהם אז המה עונים את כל הקורא אליהם, וזהו ממש שיתוף שם שמים ודבר אחר, כי לבדו אינו יכול לעשות מאומה, ורק בשיתופו עמהם הוא יכול לפעול את פעולותיו, ואין לך שיתוף שם שמים ודבר אחר יותר מזה. ומאי דעל רשב"י סני בקבלה האמיתית בתלמוד, עשה רשב"י המשקר מחבר ס' הזהר בקבלה המזויפת, רחמנא ליצלן מאמונות זרות כאלה, והאמונה הברורה הטהורה מכל סיג וחלאה היא להאמין באמונה שלימה תמימה כי לו לבדו ית' הכח והממשלה בכל הנבראים שהמציא בחסדו ולו ית' נכוין בכל עבודותינו ותפלותינו כמצווה עלינו בתורה בלתי לה' לבדו שהוא הסבה הראשונה בלי צרוף ושיתוף שום כח אחר לא עתיק ונוקביה, ולא אריך ונוקביה, ולא אבא ואימא, ולא זעיר ונוקביה, כי כל הנמצאים צריכים לו והוא אינו צריך להם, וכן כתב הראב"ע בפי' החומש בלתי לה' לבדו ולא יתערב אחר עמו ע"כ, ושמעתי קצת טפשים אמרו דלא מיקרי משתף שם שמים ודבר אחר אלא דוקא כשמוציא בפיו דומיא דמתניתין דקאמרי ליה ולך מזבח, שעליה הקשו בגמרא והא משתף שם שמים ודבר אחר וכו'. ואין זה אמת דאפי' לכוון במחשבתו בהזכירו את השם בברכה או בתפלה לשום נברא או נאצל אף אם מכיון לשתף כח הניתן בו מאת השי"ת אסור, והרי זה עובד ע"ז, דמחשבת ע"ז הקב"ה מצרפה למעשה כדגרסינן בקידושין ד' ל"ט ובחולין קמ"ב ואלא הא דכתיב למען תפוש את בית ישראל בלבם ההוא כמחשבת ע"ז הוא דכתיב דאפי' חרהור הוי כמעשה. ובנ"ד כיון שמכוין במחשבתו לזעיר אנפין הרי זה כעובד ע"ז. וקרינא ביה מה שכ' הרמב"ם בפרק ג' לראשון קרוב אתה בפיהם ורחוק מכליותיהם.

| id | word | punctuation after |
|---|---|---|
| 5:32.1 | נח | ) |
| 5:32.2 | ובסנהדרין |  |
| 5:32.3 | פ | ' |
| 5:32.4 | ד | ' |
| 5:32.5 | מיתות |  |
| 5:32.6 | גרסינן |  |
| 5:32.7 | אמר |  |
| 5:32.8 | ר | ' |
| 5:32.9 | יהודה |  |
| 5:32.10 | אלמלא |  |
| 5:32.11 | וי"ו |  |
| 5:32.12 | שבהעלוך |  |
| 5:32.13 | נתחייבו |  |
| 5:32.14 | שונאיהם |  |
| 5:32.15 | של |  |
| 5:32.16 | ישראל |  |
| 5:32.17 | כלייה | . |
| 5:32.18 | אמר |  |
| 5:32.19 | לו |  |
| 5:32.20 | רשב"י |  |
| 5:32.21 | והלא |  |
| 5:32.22 | כל |  |
| 5:32.23 | המשתף |  |
| 5:32.24 | שם |  |
| 5:32.25 | שמים |  |
| 5:32.26 | ודבר |  |
| 5:32.27 | אחר |  |
| 5:32.28 | נעקר |  |
| 5:32.29 | מן |  |
| 5:32.30 | העולם |  |
| 5:32.31 | שנ | ' |
| 5:32.32 | בלתי |  |
| 5:32.33 | לה | ' |
| 5:32.34 | לבדו |  |
| 5:32.35 | אלא |  |
| 5:32.36 | מה |  |
| 5:32.37 | ת"ל |  |
| 5:32.38 | אשר |  |
| 5:32.39 | העלוך |  |
| 5:32.40 | שאיוו |  |
| 5:32.41 | לאלוהות |  |
| 5:32.42 | הרבה |  |
| 5:32.43 | ע"כ | , |
| 5:32.44 | ומביאו |  |
| 5:32.45 | מעשה |  |
| 5:32.46 | רקח |  |
| 5:32.47 | בפרק |  |
| 5:32.48 | ב | ' |
| 5:32.49 | מהל | ' |
| 5:32.50 | ע"ז | . |
| 5:32.51 | ובפ | ' |
| 5:32.52 | לולב |  |
| 5:32.53 | וערבה |  |
| 5:32.54 | אמתניתין |  |
| 5:32.55 | דהיו |  |
| 5:32.56 | אומרין |  |
| 5:32.57 | ליה |  |
| 5:32.58 | ולך |  |
| 5:32.59 | מזבח |  |
| 5:32.60 | פריך |  |
| 5:32.61 | תלמודא |  |
| 5:32.62 | והתניא |  |
| 5:32.63 | כל |  |
| 5:32.64 | המשתף |  |
| 5:32.65 | שם |  |
| 5:32.66 | שמים |  |
| 5:32.67 | ודבר |  |
| 5:32.68 | אחד |  |
| 5:32.69 | נעקר |  |
| 5:32.70 | מן |  |
| 5:32.71 | העולם | ? |
| 5:32.72 | אלא |  |
| 5:32.73 | הכי |  |
| 5:32.74 | קאמרי |  |
| 5:32.75 | וכו | ' |
| 5:32.76 | וכ | ' |
| 5:32.77 | הרב |  |
| 5:32.78 | מהרש"א |  |
| 5:32.79 | בפ | ' |
| 5:32.80 | ד | ' |
| 5:32.81 | מיתות |  |
| 5:32.82 | דודאי |  |
| 5:32.83 | לכ"ע |  |
| 5:32.84 | גרע |  |
| 5:32.85 | טפי |  |
| 5:32.86 | כופר |  |
| 5:32.87 | בעיקר |  |
| 5:32.88 | וחייב |  |
| 5:32.89 | כלייה | , |
| 5:32.90 | אלא |  |
| 5:32.91 | דרשב"י |  |
| 5:32.92 | קאמר |  |
| 5:32.93 | שכל |  |
| 5:32.94 | המשתף |  |
| 5:32.95 | שם |  |
| 5:32.96 | שמים |  |
| 5:32.97 | ודבר |  |
| 5:32.98 | אחר |  |
| 5:32.99 | דהיינו |  |
| 5:32.100 | המאמין |  |
| 5:32.101 | נמי |  |
| 5:32.102 | בדבר |  |
| 5:32.103 | אחר |  |
| 5:32.104 | חייב |  |
| 5:32.105 | כלייה |  |
| 5:32.106 | ונעקר |  |
| 5:32.107 | מן |  |
| 5:32.108 | העולם | , |
| 5:32.109 | אלא |  |
| 5:32.110 | מה |  |
| 5:32.111 | ת"ל |  |
| 5:32.112 | אשר |  |
| 5:32.113 | העלוך |  |
| 5:32.114 | דודאי |  |
| 5:32.115 | האמינו |  |
| 5:32.116 | בנמצא |  |
| 5:32.117 | וכו | ' |
| 5:32.118 | ולכך |  |
| 5:32.119 | לא |  |
| 5:32.120 | נעקרו |  |
| 5:32.121 | אלא |  |
| 5:32.122 | שאיוו |  |
| 5:32.123 | לאלוהות |  |
| 5:32.124 | הרבה |  |
| 5:32.125 | ואמרו |  |
| 5:32.126 | שהאחד |  |
| 5:32.127 | הנמצא |  |
| 5:32.128 | נתן |  |
| 5:32.129 | כח |  |
| 5:32.130 | לאלוהות |  |
| 5:32.131 | וכוחות |  |
| 5:32.132 | רבים |  |
| 5:32.133 | עזב |  |
| 5:32.134 | ה | ' |
| 5:32.135 | את |  |
| 5:32.136 | הארץ | . |
| 5:32.137 | ואלו |  |
| 5:32.138 | הכוחות |  |
| 5:32.139 | והאלוהות |  |
| 5:32.140 | יהיו |  |
| 5:32.141 | מנהיגים |  |
| 5:32.142 | בארץ | . |
| 5:32.143 | וכן |  |
| 5:32.144 | כ | ' |
| 5:32.145 | הרמב"ם |  |
| 5:32.146 | שבה |  |
| 5:32.147 | דעת |  |
| 5:32.148 | רבים |  |
| 5:32.149 | מעובדי |  |
| 5:32.150 | ע"ז |  |
| 5:32.151 | שעובדים |  |
| 5:32.152 | אותה | . |
| 5:32.153 | שאומרים |  |
| 5:32.154 | שניתנה |  |
| 5:32.155 | ההנהגה |  |
| 5:32.156 | בידה | . |
| 5:32.157 | ועל |  |
| 5:32.158 | דרך |  |
| 5:32.159 | זה |  |
| 5:32.160 | יש |  |
| 5:32.161 | לכוון |  |
| 5:32.162 | מה |  |
| 5:32.163 | שאמר |  |
| 5:32.164 | משה |  |
| 5:32.165 | ראה |  |
| 5:32.166 | אתה |  |
| 5:32.167 | אומר |  |
| 5:32.168 | אלי |  |
| 5:32.169 | העל |  |
| 5:32.170 | את |  |
| 5:32.171 | העם |  |
| 5:32.172 | הזה | . |
| 5:32.173 | שבא |  |
| 5:32.174 | משה |  |
| 5:32.175 | להתנצל |  |
| 5:32.176 | בעדם | . |
| 5:32.177 | שאמר |  |
| 5:32.178 | משה |  |
| 5:32.179 | להקב"ה |  |
| 5:32.180 | שהרי |  |
| 5:32.181 | אתה |  |
| 5:32.182 | אומר |  |
| 5:32.183 | אלי |  |
| 5:32.184 | העל |  |
| 5:32.185 | את |  |
| 5:32.186 | העם |  |
| 5:32.187 | הזה | . |
| 5:32.188 | וכוונתך |  |
| 5:32.189 | שאתה |  |
| 5:32.190 | מעמיד |  |
| 5:32.191 | אותי |  |
| 5:32.192 | למנהיג |  |
| 5:32.193 | להם | , |
| 5:32.194 | ותלית |  |
| 5:32.195 | העלייה |  |
| 5:32.196 | בי | , |
| 5:32.197 | כיוצא |  |
| 5:32.198 | בזה |  |
| 5:32.199 | כיוונו |  |
| 5:32.200 | כלל |  |
| 5:32.201 | לע"ז |  |
| 5:32.202 | ולשתף |  |
| 5:32.203 | שם |  |
| 5:32.204 | שמים |  |
| 5:32.205 | ודבר |  |
| 5:32.206 | אחר |  |
| 5:32.207 | עכ"ל |  |
| 5:32.208 | וכל |  |
| 5:32.209 | זה |  |
| 5:32.210 | דוחק |  |
| 5:32.211 | גדול | . |
| 5:32.212 | וכ | ' |
| 5:32.213 | מעשה |  |
| 5:32.214 | רקח |  |
| 5:32.215 | דאפשר |  |
| 5:32.216 | דשיתוף |  |
| 5:32.217 | דקאמר |  |
| 5:32.218 | היינו |  |
| 5:32.219 | לכל |  |
| 5:32.220 | האלוהות |  |
| 5:32.221 | עם |  |
| 5:32.222 | הקב"ה |  |
| 5:32.223 | ע"ש |  |
| 5:32.224 | ופשוט |  |
| 5:32.225 | הוא | . |
| 5:32.226 | מכל |  |
| 5:32.227 | זה |  |
| 5:32.228 | נתבאר |  |
| 5:32.229 | לך |  |
| 5:32.230 | ידידי |  |
| 5:32.231 | האוהב |  |
| 5:32.232 | את |  |
| 5:32.233 | תורת |  |
| 5:32.234 | האמת |  |
| 5:32.235 | ודברי |  |
| 5:32.236 | דודים |  |
| 5:32.237 | מעתיקיה |  |
| 5:32.238 | ומקבליה |  |
| 5:32.239 | ממשה |  |
| 5:32.240 | רבי | ' |
| 5:32.241 | איש |  |
| 5:32.242 | מפי |  |
| 5:32.243 | איש | , |
| 5:32.244 | שזה |  |
| 5:32.245 | שאמרו |  |
| 5:32.246 | המקובלים |  |
| 5:32.247 | החדשים |  |
| 5:32.248 | כי |  |
| 5:32.249 | אלהינו |  |
| 5:32.250 | בהו"ש |  |
| 5:32.251 | בהיותו |  |
| 5:32.252 | בבחינתו |  |
| 5:32.253 | הפשוטה |  |
| 5:32.254 | לא |  |
| 5:32.255 | שייך |  |
| 5:32.256 | בו |  |
| 5:32.257 | שום |  |
| 5:32.258 | עבודה |  |
| 5:32.259 | ותפלה |  |
| 5:32.260 | וקריאה |  |
| 5:32.261 | כלל | , |
| 5:32.262 | כי |  |
| 5:32.263 | קצור |  |
| 5:32.264 | קצרה |  |
| 5:32.265 | ידו |  |
| 5:32.266 | ח"ו |  |
| 5:32.267 | מפדות |  |
| 5:32.268 | ואין |  |
| 5:32.269 | בו |  |
| 5:32.270 | כח |  |
| 5:32.271 | להציל |  |
| 5:32.272 | את |  |
| 5:32.273 | הקורא |  |
| 5:32.274 | ומשוע |  |
| 5:32.275 | אליו |  |
| 5:32.276 | ורק |  |
| 5:32.277 | בהיותו |  |
| 5:32.278 | משתתף |  |
| 5:32.279 | בפרצופים |  |
| 5:32.280 | הנז | ' |
| 5:32.281 | וביחוד |  |
| 5:32.282 | פרצוף |  |
| 5:32.283 | זעיר |  |
| 5:32.284 | אנפין |  |
| 5:32.285 | אז |  |
| 5:32.286 | שייך |  |
| 5:32.287 | בו |  |
| 5:32.288 | עבודה |  |
| 5:32.289 | וברכה | . |
| 5:32.290 | וצריך |  |
| 5:32.291 | דוקא |  |
| 5:32.292 | לאחדו |  |
| 5:32.293 | ולשתפו |  |
| 5:32.294 | עם |  |
| 5:32.295 | נוקבא |  |
| 5:32.296 | דילה |  |
| 5:32.297 | ועם |  |
| 5:32.298 | שאר |  |
| 5:32.299 | החמשה |  |
| 5:32.300 | פרצופים | , |
| 5:32.301 | ואז |  |
| 5:32.302 | בהיותו |  |
| 5:32.303 | משתתף |  |
| 5:32.304 | עמהם |  |
| 5:32.305 | אז |  |
| 5:32.306 | המה |  |
| 5:32.307 | עונים |  |
| 5:32.308 | את |  |
| 5:32.309 | כל |  |
| 5:32.310 | הקורא |  |
| 5:32.311 | אליהם | , |
| 5:32.312 | וזהו |  |
| 5:32.313 | ממש |  |
| 5:32.314 | שיתוף |  |
| 5:32.315 | שם |  |
| 5:32.316 | שמים |  |
| 5:32.317 | ודבר |  |
| 5:32.318 | אחר | , |
| 5:32.319 | כי |  |
| 5:32.320 | לבדו |  |
| 5:32.321 | אינו |  |
| 5:32.322 | יכול |  |
| 5:32.323 | לעשות |  |
| 5:32.324 | מאומה | , |
| 5:32.325 | ורק |  |
| 5:32.326 | בשיתופו |  |
| 5:32.327 | עמהם |  |
| 5:32.328 | הוא |  |
| 5:32.329 | יכול |  |
| 5:32.330 | לפעול |  |
| 5:32.331 | את |  |
| 5:32.332 | פעולותיו | , |
| 5:32.333 | ואין |  |
| 5:32.334 | לך |  |
| 5:32.335 | שיתוף |  |
| 5:32.336 | שם |  |
| 5:32.337 | שמים |  |
| 5:32.338 | ודבר |  |
| 5:32.339 | אחר |  |
| 5:32.340 | יותר |  |
| 5:32.341 | מזה | . |
| 5:32.342 | ומאי |  |
| 5:32.343 | דעל |  |
| 5:32.344 | רשב"י |  |
| 5:32.345 | סני |  |
| 5:32.346 | בקבלה |  |
| 5:32.347 | האמיתית |  |
| 5:32.348 | בתלמוד | , |
| 5:32.349 | עשה |  |
| 5:32.350 | רשב"י |  |
| 5:32.351 | המשקר |  |
| 5:32.352 | מחבר |  |
| 5:32.353 | ס | ' |
| 5:32.354 | הזהר |  |
| 5:32.355 | בקבלה |  |
| 5:32.356 | המזויפת | , |
| 5:32.357 | רחמנא |  |
| 5:32.358 | ליצלן |  |
| 5:32.359 | מאמונות |  |
| 5:32.360 | זרות |  |
| 5:32.361 | כאלה | , |
| 5:32.362 | והאמונה |  |
| 5:32.363 | הברורה |  |
| 5:32.364 | הטהורה |  |
| 5:32.365 | מכל |  |
| 5:32.366 | סיג |  |
| 5:32.367 | וחלאה |  |
| 5:32.368 | היא |  |
| 5:32.369 | להאמין |  |
| 5:32.370 | באמונה |  |
| 5:32.371 | שלימה |  |
| 5:32.372 | תמימה |  |
| 5:32.373 | כי |  |
| 5:32.374 | לו |  |
| 5:32.375 | לבדו |  |
| 5:32.376 | ית | ' |
| 5:32.377 | הכח |  |
| 5:32.378 | והממשלה |  |
| 5:32.379 | בכל |  |
| 5:32.380 | הנבראים |  |
| 5:32.381 | שהמציא |  |
| 5:32.382 | בחסדו |  |
| 5:32.383 | ולו |  |
| 5:32.384 | ית | ' |
| 5:32.385 | נכוין |  |
| 5:32.386 | בכל |  |
| 5:32.387 | עבודותינו |  |
| 5:32.388 | ותפלותינו |  |
| 5:32.389 | כמצווה |  |
| 5:32.390 | עלינו |  |
| 5:32.391 | בתורה |  |
| 5:32.392 | בלתי |  |
| 5:32.393 | לה | ' |
| 5:32.394 | לבדו |  |
| 5:32.395 | שהוא |  |
| 5:32.396 | הסבה |  |
| 5:32.397 | הראשונה |  |
| 5:32.398 | בלי |  |
| 5:32.399 | צרוף |  |
| 5:32.400 | ושיתוף |  |
| 5:32.401 | שום |  |
| 5:32.402 | כח |  |
| 5:32.403 | אחר |  |
| 5:32.404 | לא |  |
| 5:32.405 | עתיק |  |
| 5:32.406 | ונוקביה | , |
| 5:32.407 | ולא |  |
| 5:32.408 | אריך |  |
| 5:32.409 | ונוקביה | , |
| 5:32.410 | ולא |  |
| 5:32.411 | אבא |  |
| 5:32.412 | ואימא | , |
| 5:32.413 | ולא |  |
| 5:32.414 | זעיר |  |
| 5:32.415 | ונוקביה | , |
| 5:32.416 | כי |  |
| 5:32.417 | כל |  |
| 5:32.418 | הנמצאים |  |
| 5:32.419 | צריכים |  |
| 5:32.420 | לו |  |
| 5:32.421 | והוא |  |
| 5:32.422 | אינו |  |
| 5:32.423 | צריך |  |
| 5:32.424 | להם | , |
| 5:32.425 | וכן |  |
| 5:32.426 | כתב |  |
| 5:32.427 | הראב"ע |  |
| 5:32.428 | בפי | ' |
| 5:32.429 | החומש |  |
| 5:32.430 | בלתי |  |
| 5:32.431 | לה | ' |
| 5:32.432 | לבדו |  |
| 5:32.433 | ולא |  |
| 5:32.434 | יתערב |  |
| 5:32.435 | אחר |  |
| 5:32.436 | עמו |  |
| 5:32.437 | ע"כ | , |
| 5:32.438 | ושמעתי |  |
| 5:32.439 | קצת |  |
| 5:32.440 | טפשים |  |
| 5:32.441 | אמרו |  |
| 5:32.442 | דלא |  |
| 5:32.443 | מיקרי |  |
| 5:32.444 | משתף |  |
| 5:32.445 | שם |  |
| 5:32.446 | שמים |  |
| 5:32.447 | ודבר |  |
| 5:32.448 | אחר |  |
| 5:32.449 | אלא |  |
| 5:32.450 | דוקא |  |
| 5:32.451 | כשמוציא |  |
| 5:32.452 | בפיו |  |
| 5:32.453 | דומיא |  |
| 5:32.454 | דמתניתין |  |
| 5:32.455 | דקאמרי |  |
| 5:32.456 | ליה |  |
| 5:32.457 | ולך |  |
| 5:32.458 | מזבח | , |
| 5:32.459 | שעליה |  |
| 5:32.460 | הקשו |  |
| 5:32.461 | בגמרא |  |
| 5:32.462 | והא |  |
| 5:32.463 | משתף |  |
| 5:32.464 | שם |  |
| 5:32.465 | שמים |  |
| 5:32.466 | ודבר |  |
| 5:32.467 | אחר |  |
| 5:32.468 | וכו | '. |
| 5:32.469 | ואין |  |
| 5:32.470 | זה |  |
| 5:32.471 | אמת |  |
| 5:32.472 | דאפי | ' |
| 5:32.473 | לכוון |  |
| 5:32.474 | במחשבתו |  |
| 5:32.475 | בהזכירו |  |
| 5:32.476 | את |  |
| 5:32.477 | השם |  |
| 5:32.478 | בברכה |  |
| 5:32.479 | או |  |
| 5:32.480 | בתפלה |  |
| 5:32.481 | לשום |  |
| 5:32.482 | נברא |  |
| 5:32.483 | או |  |
| 5:32.484 | נאצל |  |
| 5:32.485 | אף |  |
| 5:32.486 | אם |  |
| 5:32.487 | מכיון |  |
| 5:32.488 | לשתף |  |
| 5:32.489 | כח |  |
| 5:32.490 | הניתן |  |
| 5:32.491 | בו |  |
| 5:32.492 | מאת |  |
| 5:32.493 | השי"ת |  |
| 5:32.494 | אסור | , |
| 5:32.495 | והרי |  |
| 5:32.496 | זה |  |
| 5:32.497 | עובד |  |
| 5:32.498 | ע"ז | , |
| 5:32.499 | דמחשבת |  |
| 5:32.500 | ע"ז |  |
| 5:32.501 | הקב"ה |  |
| 5:32.502 | מצרפה |  |
| 5:32.503 | למעשה |  |
| 5:32.504 | כדגרסינן |  |
| 5:32.505 | בקידושין |  |
| 5:32.506 | ד | ' |
| 5:32.507 | ל"ט |  |
| 5:32.508 | ובחולין |  |
| 5:32.509 | קמ"ב |  |
| 5:32.510 | ואלא |  |
| 5:32.511 | הא |  |
| 5:32.512 | דכתיב |  |
| 5:32.513 | למען |  |
| 5:32.514 | תפוש |  |
| 5:32.515 | את |  |
| 5:32.516 | בית |  |
| 5:32.517 | ישראל |  |
| 5:32.518 | בלבם |  |
| 5:32.519 | ההוא |  |
| 5:32.520 | כמחשבת |  |
| 5:32.521 | ע"ז |  |
| 5:32.522 | הוא |  |
| 5:32.523 | דכתיב |  |
| 5:32.524 | דאפי | ' |
| 5:32.525 | חרהור |  |
| 5:32.526 | הוי |  |
| 5:32.527 | כמעשה | . |
| 5:32.528 | ובנ"ד |  |
| 5:32.529 | כיון |  |
| 5:32.530 | שמכוין |  |
| 5:32.531 | במחשבתו |  |
| 5:32.532 | לזעיר |  |
| 5:32.533 | אנפין |  |
| 5:32.534 | הרי |  |
| 5:32.535 | זה |  |
| 5:32.536 | כעובד |  |
| 5:32.537 | ע"ז | . |
| 5:32.538 | וקרינא |  |
| 5:32.539 | ביה |  |
| 5:32.540 | מה |  |
| 5:32.541 | שכ | ' |
| 5:32.542 | הרמב"ם |  |
| 5:32.543 | בפרק |  |
| 5:32.544 | ג | ' |
| 5:32.545 | לראשון |  |
| 5:32.546 | קרוב |  |
| 5:32.547 | אתה |  |
| 5:32.548 | בפיהם |  |
| 5:32.549 | ורחוק |  |
| 5:32.550 | מכליותיהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

