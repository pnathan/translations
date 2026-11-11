# Translation pack: milhamot-hashem perek-5-11 (Milchamot ha-Shem 5:18–5:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-11`.

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


## Milchamot ha-Shem 5:18

Hebrew (Wikisource): הנה די מזה באר שהם עקרו ארבעה עיקרים הראשונים שביגדל אשר ביארתי לך בקונדריס שקדם אליך, ואל דבריהם אלה רמז ר' תם ן' יחיא באמרו כי תצא תקלה מזה לפרוץ גדרי התורה וחומותיה ויבא לידי מינות, והוא מה שנמנע הר' האר"י ושאר מקובלים מלומר יגדל, כי חשבו הקבלה החדשה לקבלה אמתית, ואת הקבלה אמתית של רז"ל הפשוטה בכל ישראל כמו שכתב הסמ"ג, והיא המבוארת ופשוטה גם כן בכתובים ובדברי רז"ל, חשבוה לטפלה וכוזבת. ואשר הביאם לידי מחשבה זו לעשות העיקר טפל והטפל עיקר הוא בעבור שראו חכמי ישראל כהרס"ג והרמב"ם ור' יהודה הלוי ור' בחיי בעל חובת הלבבות ושבילי אמונה וכיוצא בהם שהביאו ראיות חזקות על אמתת יחוד השי"ת בדברי הפילוסופים כדי לסתום את פי דוברי השקר הכופרים הטוענים נגד תורתינו הקדושה. ולכן חשבו אותה שהיא לקוחה מדברי הפילוסופים ואינה הקבלה האמיתית. ואת הרעיונות החדשים שהמציא הפילוסוף מחבר הזהר לפרש הכתובים שהם על דרך הרחבת לשון או מושאלים או משותפים כמו שכתב הרס"ג בספר האמונות והדעות והרמב"ם במורה נבוכים חלק א'. והוא נטה לפרשם כהנחתן הראשונה הפשוטה. יד ה', עיני ה', אף ה', אזני ה'. וכן: בנים אתם, אחי ורעי, בדמותי כצלמי, וכיוצא בהם, כהנחתן ומשמעותן הראשונה. וייחס רעיונות האלה בשקר לרשב"י התנא ושאר תנאים ואמוראים שלא היו בימי רשב"י אלא אחריו. ולכן חשבוה בטעות לקבלה אמתית. ולפיכך מיאנו לומר יגדל אלהים חי, כי האמינו וחשבו ליסוד מוסד אצלם כי העבודה והתפלה לאלוה קצר אפים (זעיר אנפין) וקודם הבריאה עוד לא נמצא הזעיר אנפין כמו שכתבתי לעיל בשם ספר הברית וכמבואר גם כן בעץ חיים וכסא אליהו ועז לאלהים. וגם אחר הבריאה עד אברהם אבינו עדיין לא נתקן הזעיר אנפין והיו כל אותן הדורות עד אברהם אבינו ניזונים ומקבלים שפעם מאבא ואימא כמו שכתב מה"ר יחיא הכהן בספרו חיי שלום פרשת חיי שרה בשם ליקוטי תורה פרשת לך לך. ולכך לא חפץ האר"י לומר יגדל אלהים חי כי איך יאמר נמצא ואין עת אל מציאותו, והרי יש עת אל מציאותו דהיינו מהבריאה ואילך ולא קודם לכן, או מאברהם ואילך שמהעת ההיא נשלם ונגמר תיקונו, כי קודם אברהם שעדיין לא נגמר תיקון הזעיר אנפין עדיין מקבלים שפעם מאבא ואימא ואין ראוי לעבוד לאלוה חסר התיקון, ולפי זה הרי יש עת אל מציאותו, דהיינו מהבריאה ואילך או מאברהם ואילך, והוה ליה שבח של שקר אם יאמר נמצא ואין עת אל מציאותו, שהרי קודם זה עדיין לא נמצא שהרי עדיין חסר תיקון ואין ראוי לעבודה. גם לא ערב לו לומר אחד ואין יחוד כיחודו, שהרי גם זה שקר אצלם כמו שנתבאר לעיל ממשל הבית שיש בו כמה חדרים ומקומות ועצים ואבנים וכל אבן חלוקה ומקובצת מכמה צרורות ועפר וסיד, והבנאי הוא שחיברם בחכמתו ועשאם דבר אחד והוא נקרא בשם אחד בית או כותל, נמצא שייחוד השי"ת אצלם כשאר האחדים שהם מחוברים מכמה חלקים והם נקראים בשם אחד, ואיך יאמר אחד ואין יחוד כיחודו! גם נמנע מלומר נעלם שהרי ביארו שיש כמה עולמות נאצלו ונבראו ונוצרו ונעשו למעלה מעולם האצילות, אכן לרוב העלמם לא שלחו בהם יד לבארם ולדבר בהם. וגם באדם קדמון כתבו שלא דיברו בו אלא ברמז קצת, והכלל אצלם שהפרצופים שלמעלה מעולם האצילות לרוב העלמם שהם זכים ודקים ונעלמים מאד לא שייכא תפילה וקריאה אליהם, ורק ביארו הפרצופים שכבר נתעבו ונראו קצת. ובהם דוקא שייך עבודה ותפלה וקרבנות. וזהו הטעם דלא שייך עבודה ותפלה באין סוף ובאדם קדמאה וכו' שבעולמות שלמעלה, ולא באדם קדמון דאצילות, וגם לא בשאר פרצופים שבאצילות לרוב העלמם ודקותם ורק שייך עבודה דוקא בזעיר אנפין פרצוף האחרון שבהם שכבר נתעבה ונתגלה יותר ונעשה בחינה המושגת כמו שכתב ספר הברית וכמו שכתב הרח"ו בעץ חיים ושער עיגולים ויושר ושער סדר אבי"ע ובשער השתלשלות העשר ספירות, ובכסא אליהו שער יראת ה' ועיין לקמן סי' פ"ד מה שכתבתי בשם משנת חסידים כי על כן רצו להשמיט נעלם ואין לו סוף לאחדותו שאיננו שבח אמתי לפי דעתם. גם הרחיקו מלומר אין לו דמות הגוף ואינו גוף. שהרי לפי דעתם יש לו דמות הגוף ברמ"ח איברים ושס"ה גידים כמבואר בכסא אליהו שער ראש חדש דף ל"ב ע"ב כי החמשה פרצופים כל אחד מהם פרצוף גמור בגילוי חתוך איברים רמ"ח איברים ושס"ה גידיים והביא ראיות לזה מדברי הזהר, וכן כתב הרח"ו בספר עץ חיים שער עיגולים ויושר כל הקו המתפשט מאין סוף בבחינת יושר הוא בציור אדם ישר כלול מרמ"ח איברים בסוד צלם אלהים. ועליה רמז הכתוב ויברא אלהים את האדם בצלמו בצלם אלהים ברא אותו, וכל דברי הזהר והתקונים בבחינת דיושר ע"ש, וכן כתב ספר הברית בחלק דברי אמת (מאמר א' פרק ז' וח') וכן כתב הראב"ד בפרק ד' מהלכות תשובה על מה שכתב הרמב"ם: "האומר שיש שם ריבון אחד אבל שהוא גוף ובעל תמונה", כתב הראב"ד: "ולמה קרא לזה מין וכמה גדולים וטובים ממנו הלכו בזו המחשבה לפי מה שראו במקראות" וכו'.

| id | word | punctuation after |
|---|---|---|
| 5:18.1 | הנה |  |
| 5:18.2 | די |  |
| 5:18.3 | מזה |  |
| 5:18.4 | באר |  |
| 5:18.5 | שהם |  |
| 5:18.6 | עקרו |  |
| 5:18.7 | ארבעה |  |
| 5:18.8 | עיקרים |  |
| 5:18.9 | הראשונים |  |
| 5:18.10 | שביגדל |  |
| 5:18.11 | אשר |  |
| 5:18.12 | ביארתי |  |
| 5:18.13 | לך |  |
| 5:18.14 | בקונדריס |  |
| 5:18.15 | שקדם |  |
| 5:18.16 | אליך | , |
| 5:18.17 | ואל |  |
| 5:18.18 | דבריהם |  |
| 5:18.19 | אלה |  |
| 5:18.20 | רמז |  |
| 5:18.21 | ר | ' |
| 5:18.22 | תם |  |
| 5:18.23 | ן | ' |
| 5:18.24 | יחיא |  |
| 5:18.25 | באמרו |  |
| 5:18.26 | כי |  |
| 5:18.27 | תצא |  |
| 5:18.28 | תקלה |  |
| 5:18.29 | מזה |  |
| 5:18.30 | לפרוץ |  |
| 5:18.31 | גדרי |  |
| 5:18.32 | התורה |  |
| 5:18.33 | וחומותיה |  |
| 5:18.34 | ויבא |  |
| 5:18.35 | לידי |  |
| 5:18.36 | מינות | , |
| 5:18.37 | והוא |  |
| 5:18.38 | מה |  |
| 5:18.39 | שנמנע |  |
| 5:18.40 | הר | ' |
| 5:18.41 | האר"י |  |
| 5:18.42 | ושאר |  |
| 5:18.43 | מקובלים |  |
| 5:18.44 | מלומר |  |
| 5:18.45 | יגדל | , |
| 5:18.46 | כי |  |
| 5:18.47 | חשבו |  |
| 5:18.48 | הקבלה |  |
| 5:18.49 | החדשה |  |
| 5:18.50 | לקבלה |  |
| 5:18.51 | אמתית | , |
| 5:18.52 | ואת |  |
| 5:18.53 | הקבלה |  |
| 5:18.54 | אמתית |  |
| 5:18.55 | של |  |
| 5:18.56 | רז"ל |  |
| 5:18.57 | הפשוטה |  |
| 5:18.58 | בכל |  |
| 5:18.59 | ישראל |  |
| 5:18.60 | כמו |  |
| 5:18.61 | שכתב |  |
| 5:18.62 | הסמ"ג | , |
| 5:18.63 | והיא |  |
| 5:18.64 | המבוארת |  |
| 5:18.65 | ופשוטה |  |
| 5:18.66 | גם |  |
| 5:18.67 | כן |  |
| 5:18.68 | בכתובים |  |
| 5:18.69 | ובדברי |  |
| 5:18.70 | רז"ל | , |
| 5:18.71 | חשבוה |  |
| 5:18.72 | לטפלה |  |
| 5:18.73 | וכוזבת | . |
| 5:18.74 | ואשר |  |
| 5:18.75 | הביאם |  |
| 5:18.76 | לידי |  |
| 5:18.77 | מחשבה |  |
| 5:18.78 | זו |  |
| 5:18.79 | לעשות |  |
| 5:18.80 | העיקר |  |
| 5:18.81 | טפל |  |
| 5:18.82 | והטפל |  |
| 5:18.83 | עיקר |  |
| 5:18.84 | הוא |  |
| 5:18.85 | בעבור |  |
| 5:18.86 | שראו |  |
| 5:18.87 | חכמי |  |
| 5:18.88 | ישראל |  |
| 5:18.89 | כהרס"ג |  |
| 5:18.90 | והרמב"ם |  |
| 5:18.91 | ור | ' |
| 5:18.92 | יהודה |  |
| 5:18.93 | הלוי |  |
| 5:18.94 | ור | ' |
| 5:18.95 | בחיי |  |
| 5:18.96 | בעל |  |
| 5:18.97 | חובת |  |
| 5:18.98 | הלבבות |  |
| 5:18.99 | ושבילי |  |
| 5:18.100 | אמונה |  |
| 5:18.101 | וכיוצא |  |
| 5:18.102 | בהם |  |
| 5:18.103 | שהביאו |  |
| 5:18.104 | ראיות |  |
| 5:18.105 | חזקות |  |
| 5:18.106 | על |  |
| 5:18.107 | אמתת |  |
| 5:18.108 | יחוד |  |
| 5:18.109 | השי"ת |  |
| 5:18.110 | בדברי |  |
| 5:18.111 | הפילוסופים |  |
| 5:18.112 | כדי |  |
| 5:18.113 | לסתום |  |
| 5:18.114 | את |  |
| 5:18.115 | פי |  |
| 5:18.116 | דוברי |  |
| 5:18.117 | השקר |  |
| 5:18.118 | הכופרים |  |
| 5:18.119 | הטוענים |  |
| 5:18.120 | נגד |  |
| 5:18.121 | תורתינו |  |
| 5:18.122 | הקדושה | . |
| 5:18.123 | ולכן |  |
| 5:18.124 | חשבו |  |
| 5:18.125 | אותה |  |
| 5:18.126 | שהיא |  |
| 5:18.127 | לקוחה |  |
| 5:18.128 | מדברי |  |
| 5:18.129 | הפילוסופים |  |
| 5:18.130 | ואינה |  |
| 5:18.131 | הקבלה |  |
| 5:18.132 | האמיתית | . |
| 5:18.133 | ואת |  |
| 5:18.134 | הרעיונות |  |
| 5:18.135 | החדשים |  |
| 5:18.136 | שהמציא |  |
| 5:18.137 | הפילוסוף |  |
| 5:18.138 | מחבר |  |
| 5:18.139 | הזהר |  |
| 5:18.140 | לפרש |  |
| 5:18.141 | הכתובים |  |
| 5:18.142 | שהם |  |
| 5:18.143 | על |  |
| 5:18.144 | דרך |  |
| 5:18.145 | הרחבת |  |
| 5:18.146 | לשון |  |
| 5:18.147 | או |  |
| 5:18.148 | מושאלים |  |
| 5:18.149 | או |  |
| 5:18.150 | משותפים |  |
| 5:18.151 | כמו |  |
| 5:18.152 | שכתב |  |
| 5:18.153 | הרס"ג |  |
| 5:18.154 | בספר |  |
| 5:18.155 | האמונות |  |
| 5:18.156 | והדעות |  |
| 5:18.157 | והרמב"ם |  |
| 5:18.158 | במורה |  |
| 5:18.159 | נבוכים |  |
| 5:18.160 | חלק |  |
| 5:18.161 | א | '. |
| 5:18.162 | והוא |  |
| 5:18.163 | נטה |  |
| 5:18.164 | לפרשם |  |
| 5:18.165 | כהנחתן |  |
| 5:18.166 | הראשונה |  |
| 5:18.167 | הפשוטה | . |
| 5:18.168 | יד |  |
| 5:18.169 | ה | ', |
| 5:18.170 | עיני |  |
| 5:18.171 | ה | ', |
| 5:18.172 | אף |  |
| 5:18.173 | ה | ', |
| 5:18.174 | אזני |  |
| 5:18.175 | ה | '. |
| 5:18.176 | וכן | : |
| 5:18.177 | בנים |  |
| 5:18.178 | אתם | , |
| 5:18.179 | אחי |  |
| 5:18.180 | ורעי | , |
| 5:18.181 | בדמותי |  |
| 5:18.182 | כצלמי | , |
| 5:18.183 | וכיוצא |  |
| 5:18.184 | בהם | , |
| 5:18.185 | כהנחתן |  |
| 5:18.186 | ומשמעותן |  |
| 5:18.187 | הראשונה | . |
| 5:18.188 | וייחס |  |
| 5:18.189 | רעיונות |  |
| 5:18.190 | האלה |  |
| 5:18.191 | בשקר |  |
| 5:18.192 | לרשב"י |  |
| 5:18.193 | התנא |  |
| 5:18.194 | ושאר |  |
| 5:18.195 | תנאים |  |
| 5:18.196 | ואמוראים |  |
| 5:18.197 | שלא |  |
| 5:18.198 | היו |  |
| 5:18.199 | בימי |  |
| 5:18.200 | רשב"י |  |
| 5:18.201 | אלא |  |
| 5:18.202 | אחריו | . |
| 5:18.203 | ולכן |  |
| 5:18.204 | חשבוה |  |
| 5:18.205 | בטעות |  |
| 5:18.206 | לקבלה |  |
| 5:18.207 | אמתית | . |
| 5:18.208 | ולפיכך |  |
| 5:18.209 | מיאנו |  |
| 5:18.210 | לומר |  |
| 5:18.211 | יגדל |  |
| 5:18.212 | אלהים |  |
| 5:18.213 | חי | , |
| 5:18.214 | כי |  |
| 5:18.215 | האמינו |  |
| 5:18.216 | וחשבו |  |
| 5:18.217 | ליסוד |  |
| 5:18.218 | מוסד |  |
| 5:18.219 | אצלם |  |
| 5:18.220 | כי |  |
| 5:18.221 | העבודה |  |
| 5:18.222 | והתפלה |  |
| 5:18.223 | לאלוה |  |
| 5:18.224 | קצר |  |
| 5:18.225 | אפים |  |
| 5:18.226 | זעיר |  |
| 5:18.227 | אנפין | ) |
| 5:18.228 | וקודם |  |
| 5:18.229 | הבריאה |  |
| 5:18.230 | עוד |  |
| 5:18.231 | לא |  |
| 5:18.232 | נמצא |  |
| 5:18.233 | הזעיר |  |
| 5:18.234 | אנפין |  |
| 5:18.235 | כמו |  |
| 5:18.236 | שכתבתי |  |
| 5:18.237 | לעיל |  |
| 5:18.238 | בשם |  |
| 5:18.239 | ספר |  |
| 5:18.240 | הברית |  |
| 5:18.241 | וכמבואר |  |
| 5:18.242 | גם |  |
| 5:18.243 | כן |  |
| 5:18.244 | בעץ |  |
| 5:18.245 | חיים |  |
| 5:18.246 | וכסא |  |
| 5:18.247 | אליהו |  |
| 5:18.248 | ועז |  |
| 5:18.249 | לאלהים | . |
| 5:18.250 | וגם |  |
| 5:18.251 | אחר |  |
| 5:18.252 | הבריאה |  |
| 5:18.253 | עד |  |
| 5:18.254 | אברהם |  |
| 5:18.255 | אבינו |  |
| 5:18.256 | עדיין |  |
| 5:18.257 | לא |  |
| 5:18.258 | נתקן |  |
| 5:18.259 | הזעיר |  |
| 5:18.260 | אנפין |  |
| 5:18.261 | והיו |  |
| 5:18.262 | כל |  |
| 5:18.263 | אותן |  |
| 5:18.264 | הדורות |  |
| 5:18.265 | עד |  |
| 5:18.266 | אברהם |  |
| 5:18.267 | אבינו |  |
| 5:18.268 | ניזונים |  |
| 5:18.269 | ומקבלים |  |
| 5:18.270 | שפעם |  |
| 5:18.271 | מאבא |  |
| 5:18.272 | ואימא |  |
| 5:18.273 | כמו |  |
| 5:18.274 | שכתב |  |
| 5:18.275 | מה"ר |  |
| 5:18.276 | יחיא |  |
| 5:18.277 | הכהן |  |
| 5:18.278 | בספרו |  |
| 5:18.279 | חיי |  |
| 5:18.280 | שלום |  |
| 5:18.281 | פרשת |  |
| 5:18.282 | חיי |  |
| 5:18.283 | שרה |  |
| 5:18.284 | בשם |  |
| 5:18.285 | ליקוטי |  |
| 5:18.286 | תורה |  |
| 5:18.287 | פרשת |  |
| 5:18.288 | לך |  |
| 5:18.289 | לך | . |
| 5:18.290 | ולכך |  |
| 5:18.291 | לא |  |
| 5:18.292 | חפץ |  |
| 5:18.293 | האר"י |  |
| 5:18.294 | לומר |  |
| 5:18.295 | יגדל |  |
| 5:18.296 | אלהים |  |
| 5:18.297 | חי |  |
| 5:18.298 | כי |  |
| 5:18.299 | איך |  |
| 5:18.300 | יאמר |  |
| 5:18.301 | נמצא |  |
| 5:18.302 | ואין |  |
| 5:18.303 | עת |  |
| 5:18.304 | אל |  |
| 5:18.305 | מציאותו | , |
| 5:18.306 | והרי |  |
| 5:18.307 | יש |  |
| 5:18.308 | עת |  |
| 5:18.309 | אל |  |
| 5:18.310 | מציאותו |  |
| 5:18.311 | דהיינו |  |
| 5:18.312 | מהבריאה |  |
| 5:18.313 | ואילך |  |
| 5:18.314 | ולא |  |
| 5:18.315 | קודם |  |
| 5:18.316 | לכן | , |
| 5:18.317 | או |  |
| 5:18.318 | מאברהם |  |
| 5:18.319 | ואילך |  |
| 5:18.320 | שמהעת |  |
| 5:18.321 | ההיא |  |
| 5:18.322 | נשלם |  |
| 5:18.323 | ונגמר |  |
| 5:18.324 | תיקונו | , |
| 5:18.325 | כי |  |
| 5:18.326 | קודם |  |
| 5:18.327 | אברהם |  |
| 5:18.328 | שעדיין |  |
| 5:18.329 | לא |  |
| 5:18.330 | נגמר |  |
| 5:18.331 | תיקון |  |
| 5:18.332 | הזעיר |  |
| 5:18.333 | אנפין |  |
| 5:18.334 | עדיין |  |
| 5:18.335 | מקבלים |  |
| 5:18.336 | שפעם |  |
| 5:18.337 | מאבא |  |
| 5:18.338 | ואימא |  |
| 5:18.339 | ואין |  |
| 5:18.340 | ראוי |  |
| 5:18.341 | לעבוד |  |
| 5:18.342 | לאלוה |  |
| 5:18.343 | חסר |  |
| 5:18.344 | התיקון | , |
| 5:18.345 | ולפי |  |
| 5:18.346 | זה |  |
| 5:18.347 | הרי |  |
| 5:18.348 | יש |  |
| 5:18.349 | עת |  |
| 5:18.350 | אל |  |
| 5:18.351 | מציאותו | , |
| 5:18.352 | דהיינו |  |
| 5:18.353 | מהבריאה |  |
| 5:18.354 | ואילך |  |
| 5:18.355 | או |  |
| 5:18.356 | מאברהם |  |
| 5:18.357 | ואילך | , |
| 5:18.358 | והוה |  |
| 5:18.359 | ליה |  |
| 5:18.360 | שבח |  |
| 5:18.361 | של |  |
| 5:18.362 | שקר |  |
| 5:18.363 | אם |  |
| 5:18.364 | יאמר |  |
| 5:18.365 | נמצא |  |
| 5:18.366 | ואין |  |
| 5:18.367 | עת |  |
| 5:18.368 | אל |  |
| 5:18.369 | מציאותו | , |
| 5:18.370 | שהרי |  |
| 5:18.371 | קודם |  |
| 5:18.372 | זה |  |
| 5:18.373 | עדיין |  |
| 5:18.374 | לא |  |
| 5:18.375 | נמצא |  |
| 5:18.376 | שהרי |  |
| 5:18.377 | עדיין |  |
| 5:18.378 | חסר |  |
| 5:18.379 | תיקון |  |
| 5:18.380 | ואין |  |
| 5:18.381 | ראוי |  |
| 5:18.382 | לעבודה | . |
| 5:18.383 | גם |  |
| 5:18.384 | לא |  |
| 5:18.385 | ערב |  |
| 5:18.386 | לו |  |
| 5:18.387 | לומר |  |
| 5:18.388 | אחד |  |
| 5:18.389 | ואין |  |
| 5:18.390 | יחוד |  |
| 5:18.391 | כיחודו | , |
| 5:18.392 | שהרי |  |
| 5:18.393 | גם |  |
| 5:18.394 | זה |  |
| 5:18.395 | שקר |  |
| 5:18.396 | אצלם |  |
| 5:18.397 | כמו |  |
| 5:18.398 | שנתבאר |  |
| 5:18.399 | לעיל |  |
| 5:18.400 | ממשל |  |
| 5:18.401 | הבית |  |
| 5:18.402 | שיש |  |
| 5:18.403 | בו |  |
| 5:18.404 | כמה |  |
| 5:18.405 | חדרים |  |
| 5:18.406 | ומקומות |  |
| 5:18.407 | ועצים |  |
| 5:18.408 | ואבנים |  |
| 5:18.409 | וכל |  |
| 5:18.410 | אבן |  |
| 5:18.411 | חלוקה |  |
| 5:18.412 | ומקובצת |  |
| 5:18.413 | מכמה |  |
| 5:18.414 | צרורות |  |
| 5:18.415 | ועפר |  |
| 5:18.416 | וסיד | , |
| 5:18.417 | והבנאי |  |
| 5:18.418 | הוא |  |
| 5:18.419 | שחיברם |  |
| 5:18.420 | בחכמתו |  |
| 5:18.421 | ועשאם |  |
| 5:18.422 | דבר |  |
| 5:18.423 | אחד |  |
| 5:18.424 | והוא |  |
| 5:18.425 | נקרא |  |
| 5:18.426 | בשם |  |
| 5:18.427 | אחד |  |
| 5:18.428 | בית |  |
| 5:18.429 | או |  |
| 5:18.430 | כותל | , |
| 5:18.431 | נמצא |  |
| 5:18.432 | שייחוד |  |
| 5:18.433 | השי"ת |  |
| 5:18.434 | אצלם |  |
| 5:18.435 | כשאר |  |
| 5:18.436 | האחדים |  |
| 5:18.437 | שהם |  |
| 5:18.438 | מחוברים |  |
| 5:18.439 | מכמה |  |
| 5:18.440 | חלקים |  |
| 5:18.441 | והם |  |
| 5:18.442 | נקראים |  |
| 5:18.443 | בשם |  |
| 5:18.444 | אחד | , |
| 5:18.445 | ואיך |  |
| 5:18.446 | יאמר |  |
| 5:18.447 | אחד |  |
| 5:18.448 | ואין |  |
| 5:18.449 | יחוד |  |
| 5:18.450 | כיחודו | ! |
| 5:18.451 | גם |  |
| 5:18.452 | נמנע |  |
| 5:18.453 | מלומר |  |
| 5:18.454 | נעלם |  |
| 5:18.455 | שהרי |  |
| 5:18.456 | ביארו |  |
| 5:18.457 | שיש |  |
| 5:18.458 | כמה |  |
| 5:18.459 | עולמות |  |
| 5:18.460 | נאצלו |  |
| 5:18.461 | ונבראו |  |
| 5:18.462 | ונוצרו |  |
| 5:18.463 | ונעשו |  |
| 5:18.464 | למעלה |  |
| 5:18.465 | מעולם |  |
| 5:18.466 | האצילות | , |
| 5:18.467 | אכן |  |
| 5:18.468 | לרוב |  |
| 5:18.469 | העלמם |  |
| 5:18.470 | לא |  |
| 5:18.471 | שלחו |  |
| 5:18.472 | בהם |  |
| 5:18.473 | יד |  |
| 5:18.474 | לבארם |  |
| 5:18.475 | ולדבר |  |
| 5:18.476 | בהם | . |
| 5:18.477 | וגם |  |
| 5:18.478 | באדם |  |
| 5:18.479 | קדמון |  |
| 5:18.480 | כתבו |  |
| 5:18.481 | שלא |  |
| 5:18.482 | דיברו |  |
| 5:18.483 | בו |  |
| 5:18.484 | אלא |  |
| 5:18.485 | ברמז |  |
| 5:18.486 | קצת | , |
| 5:18.487 | והכלל |  |
| 5:18.488 | אצלם |  |
| 5:18.489 | שהפרצופים |  |
| 5:18.490 | שלמעלה |  |
| 5:18.491 | מעולם |  |
| 5:18.492 | האצילות |  |
| 5:18.493 | לרוב |  |
| 5:18.494 | העלמם |  |
| 5:18.495 | שהם |  |
| 5:18.496 | זכים |  |
| 5:18.497 | ודקים |  |
| 5:18.498 | ונעלמים |  |
| 5:18.499 | מאד |  |
| 5:18.500 | לא |  |
| 5:18.501 | שייכא |  |
| 5:18.502 | תפילה |  |
| 5:18.503 | וקריאה |  |
| 5:18.504 | אליהם | , |
| 5:18.505 | ורק |  |
| 5:18.506 | ביארו |  |
| 5:18.507 | הפרצופים |  |
| 5:18.508 | שכבר |  |
| 5:18.509 | נתעבו |  |
| 5:18.510 | ונראו |  |
| 5:18.511 | קצת | . |
| 5:18.512 | ובהם |  |
| 5:18.513 | דוקא |  |
| 5:18.514 | שייך |  |
| 5:18.515 | עבודה |  |
| 5:18.516 | ותפלה |  |
| 5:18.517 | וקרבנות | . |
| 5:18.518 | וזהו |  |
| 5:18.519 | הטעם |  |
| 5:18.520 | דלא |  |
| 5:18.521 | שייך |  |
| 5:18.522 | עבודה |  |
| 5:18.523 | ותפלה |  |
| 5:18.524 | באין |  |
| 5:18.525 | סוף |  |
| 5:18.526 | ובאדם |  |
| 5:18.527 | קדמאה |  |
| 5:18.528 | וכו | ' |
| 5:18.529 | שבעולמות |  |
| 5:18.530 | שלמעלה | , |
| 5:18.531 | ולא |  |
| 5:18.532 | באדם |  |
| 5:18.533 | קדמון |  |
| 5:18.534 | דאצילות | , |
| 5:18.535 | וגם |  |
| 5:18.536 | לא |  |
| 5:18.537 | בשאר |  |
| 5:18.538 | פרצופים |  |
| 5:18.539 | שבאצילות |  |
| 5:18.540 | לרוב |  |
| 5:18.541 | העלמם |  |
| 5:18.542 | ודקותם |  |
| 5:18.543 | ורק |  |
| 5:18.544 | שייך |  |
| 5:18.545 | עבודה |  |
| 5:18.546 | דוקא |  |
| 5:18.547 | בזעיר |  |
| 5:18.548 | אנפין |  |
| 5:18.549 | פרצוף |  |
| 5:18.550 | האחרון |  |
| 5:18.551 | שבהם |  |
| 5:18.552 | שכבר |  |
| 5:18.553 | נתעבה |  |
| 5:18.554 | ונתגלה |  |
| 5:18.555 | יותר |  |
| 5:18.556 | ונעשה |  |
| 5:18.557 | בחינה |  |
| 5:18.558 | המושגת |  |
| 5:18.559 | כמו |  |
| 5:18.560 | שכתב |  |
| 5:18.561 | ספר |  |
| 5:18.562 | הברית |  |
| 5:18.563 | וכמו |  |
| 5:18.564 | שכתב |  |
| 5:18.565 | הרח"ו |  |
| 5:18.566 | בעץ |  |
| 5:18.567 | חיים |  |
| 5:18.568 | ושער |  |
| 5:18.569 | עיגולים |  |
| 5:18.570 | ויושר |  |
| 5:18.571 | ושער |  |
| 5:18.572 | סדר |  |
| 5:18.573 | אבי"ע |  |
| 5:18.574 | ובשער |  |
| 5:18.575 | השתלשלות |  |
| 5:18.576 | העשר |  |
| 5:18.577 | ספירות | , |
| 5:18.578 | ובכסא |  |
| 5:18.579 | אליהו |  |
| 5:18.580 | שער |  |
| 5:18.581 | יראת |  |
| 5:18.582 | ה | ' |
| 5:18.583 | ועיין |  |
| 5:18.584 | לקמן |  |
| 5:18.585 | סי | ' |
| 5:18.586 | פ"ד |  |
| 5:18.587 | מה |  |
| 5:18.588 | שכתבתי |  |
| 5:18.589 | בשם |  |
| 5:18.590 | משנת |  |
| 5:18.591 | חסידים |  |
| 5:18.592 | כי |  |
| 5:18.593 | על |  |
| 5:18.594 | כן |  |
| 5:18.595 | רצו |  |
| 5:18.596 | להשמיט |  |
| 5:18.597 | נעלם |  |
| 5:18.598 | ואין |  |
| 5:18.599 | לו |  |
| 5:18.600 | סוף |  |
| 5:18.601 | לאחדותו |  |
| 5:18.602 | שאיננו |  |
| 5:18.603 | שבח |  |
| 5:18.604 | אמתי |  |
| 5:18.605 | לפי |  |
| 5:18.606 | דעתם | . |
| 5:18.607 | גם |  |
| 5:18.608 | הרחיקו |  |
| 5:18.609 | מלומר |  |
| 5:18.610 | אין |  |
| 5:18.611 | לו |  |
| 5:18.612 | דמות |  |
| 5:18.613 | הגוף |  |
| 5:18.614 | ואינו |  |
| 5:18.615 | גוף | . |
| 5:18.616 | שהרי |  |
| 5:18.617 | לפי |  |
| 5:18.618 | דעתם |  |
| 5:18.619 | יש |  |
| 5:18.620 | לו |  |
| 5:18.621 | דמות |  |
| 5:18.622 | הגוף |  |
| 5:18.623 | ברמ"ח |  |
| 5:18.624 | איברים |  |
| 5:18.625 | ושס"ה |  |
| 5:18.626 | גידים |  |
| 5:18.627 | כמבואר |  |
| 5:18.628 | בכסא |  |
| 5:18.629 | אליהו |  |
| 5:18.630 | שער |  |
| 5:18.631 | ראש |  |
| 5:18.632 | חדש |  |
| 5:18.633 | דף |  |
| 5:18.634 | ל"ב |  |
| 5:18.635 | ע"ב |  |
| 5:18.636 | כי |  |
| 5:18.637 | החמשה |  |
| 5:18.638 | פרצופים |  |
| 5:18.639 | כל |  |
| 5:18.640 | אחד |  |
| 5:18.641 | מהם |  |
| 5:18.642 | פרצוף |  |
| 5:18.643 | גמור |  |
| 5:18.644 | בגילוי |  |
| 5:18.645 | חתוך |  |
| 5:18.646 | איברים |  |
| 5:18.647 | רמ"ח |  |
| 5:18.648 | איברים |  |
| 5:18.649 | ושס"ה |  |
| 5:18.650 | גידיים |  |
| 5:18.651 | והביא |  |
| 5:18.652 | ראיות |  |
| 5:18.653 | לזה |  |
| 5:18.654 | מדברי |  |
| 5:18.655 | הזהר | , |
| 5:18.656 | וכן |  |
| 5:18.657 | כתב |  |
| 5:18.658 | הרח"ו |  |
| 5:18.659 | בספר |  |
| 5:18.660 | עץ |  |
| 5:18.661 | חיים |  |
| 5:18.662 | שער |  |
| 5:18.663 | עיגולים |  |
| 5:18.664 | ויושר |  |
| 5:18.665 | כל |  |
| 5:18.666 | הקו |  |
| 5:18.667 | המתפשט |  |
| 5:18.668 | מאין |  |
| 5:18.669 | סוף |  |
| 5:18.670 | בבחינת |  |
| 5:18.671 | יושר |  |
| 5:18.672 | הוא |  |
| 5:18.673 | בציור |  |
| 5:18.674 | אדם |  |
| 5:18.675 | ישר |  |
| 5:18.676 | כלול |  |
| 5:18.677 | מרמ"ח |  |
| 5:18.678 | איברים |  |
| 5:18.679 | בסוד |  |
| 5:18.680 | צלם |  |
| 5:18.681 | אלהים | . |
| 5:18.682 | ועליה |  |
| 5:18.683 | רמז |  |
| 5:18.684 | הכתוב |  |
| 5:18.685 | ויברא |  |
| 5:18.686 | אלהים |  |
| 5:18.687 | את |  |
| 5:18.688 | האדם |  |
| 5:18.689 | בצלמו |  |
| 5:18.690 | בצלם |  |
| 5:18.691 | אלהים |  |
| 5:18.692 | ברא |  |
| 5:18.693 | אותו | , |
| 5:18.694 | וכל |  |
| 5:18.695 | דברי |  |
| 5:18.696 | הזהר |  |
| 5:18.697 | והתקונים |  |
| 5:18.698 | בבחינת |  |
| 5:18.699 | דיושר |  |
| 5:18.700 | ע"ש | , |
| 5:18.701 | וכן |  |
| 5:18.702 | כתב |  |
| 5:18.703 | ספר |  |
| 5:18.704 | הברית |  |
| 5:18.705 | בחלק |  |
| 5:18.706 | דברי |  |
| 5:18.707 | אמת |  |
| 5:18.708 | מאמר |  |
| 5:18.709 | א | ' |
| 5:18.710 | פרק |  |
| 5:18.711 | ז | ' |
| 5:18.712 | וח | ') |
| 5:18.713 | וכן |  |
| 5:18.714 | כתב |  |
| 5:18.715 | הראב"ד |  |
| 5:18.716 | בפרק |  |
| 5:18.717 | ד | ' |
| 5:18.718 | מהלכות |  |
| 5:18.719 | תשובה |  |
| 5:18.720 | על |  |
| 5:18.721 | מה |  |
| 5:18.722 | שכתב |  |
| 5:18.723 | הרמב"ם | : |
| 5:18.724 | האומר |  |
| 5:18.725 | שיש |  |
| 5:18.726 | שם |  |
| 5:18.727 | ריבון |  |
| 5:18.728 | אחד |  |
| 5:18.729 | אבל |  |
| 5:18.730 | שהוא |  |
| 5:18.731 | גוף |  |
| 5:18.732 | ובעל |  |
| 5:18.733 | תמונה | ", |
| 5:18.734 | כתב |  |
| 5:18.735 | הראב"ד | : |
| 5:18.736 | ולמה |  |
| 5:18.737 | קרא |  |
| 5:18.738 | לזה |  |
| 5:18.739 | מין |  |
| 5:18.740 | וכמה |  |
| 5:18.741 | גדולים |  |
| 5:18.742 | וטובים |  |
| 5:18.743 | ממנו |  |
| 5:18.744 | הלכו |  |
| 5:18.745 | בזו |  |
| 5:18.746 | המחשבה |  |
| 5:18.747 | לפי |  |
| 5:18.748 | מה |  |
| 5:18.749 | שראו |  |
| 5:18.750 | במקראות | " |
| 5:18.751 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

