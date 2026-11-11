# Translation pack: milhamot-hashem perek-6-04 (Milchamot ha-Shem 6:8–6:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-04`.

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
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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

- סתו"מ: always "Torah scrolls, tefillin and mezuzot", with a note.
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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
- מ"ו = "our teacher" (מורנו) in the formula יודיענו מורנו; ב"ו = "blessed be
  He" only before הוא elsewhere ב"ו = בשר ודם ("flesh and blood").
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- ר"ל with a named authority as subject: "means to say".
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
- דמיון: "analogy" for an argument, "likeness" for resemblance.
- דברי תימה: "astonishing words" (words that astonish).
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- משותף in the Maimonidean sense (a term said of God and of creatures by
  equivocation): "equivocal"; "shared" only for a thing held in common.
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
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 6:8

Hebrew (Wikisource): והנה מצות האחדות ר"ל להאמין שהו ית' אחד פשוט אמתי הנמשך אחר שלילת המקרים, והתארים, ושלילת הגשמות, ואמונת חיוב מציאותו בבחינת עצמו, ושלילת השינוי ממנו ית' וכיוצא בהם מהאמונות האלהיות הבלתי מבוארת לא בתורה שבכתב ולא בתורה שבעל פה, שהמאמין בהם שלא כדת אינו בכלל בעלי האמונה השלימה, ולא אותו ית' יעבוד כמו שכ' הרב בעל חובת הלבבות לא יוכל לעבוד עילת העילות זולת הנביא או החכם השלום, על כן המצוה עלינו לחקור על מצות אלה מפי סופרים ומספרים האלהיים למען לא יטעה באמונתו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:8.1 | והנה |  |
| 6:8.2 | מצות |  |
| 6:8.3 | האחדות |  |
| 6:8.4 | ר"ל |  |
| 6:8.5 | להאמין |  |
| 6:8.6 | שהו |  |
| 6:8.7 | ית | ' |
| 6:8.8 | אחד |  |
| 6:8.9 | פשוט |  |
| 6:8.10 | אמתי |  |
| 6:8.11 | הנמשך |  |
| 6:8.12 | אחר |  |
| 6:8.13 | שלילת |  |
| 6:8.14 | המקרים | , |
| 6:8.15 | והתארים | , |
| 6:8.16 | ושלילת |  |
| 6:8.17 | הגשמות | , |
| 6:8.18 | ואמונת |  |
| 6:8.19 | חיוב |  |
| 6:8.20 | מציאותו |  |
| 6:8.21 | בבחינת |  |
| 6:8.22 | עצמו | , |
| 6:8.23 | ושלילת |  |
| 6:8.24 | השינוי |  |
| 6:8.25 | ממנו |  |
| 6:8.26 | ית | ' |
| 6:8.27 | וכיוצא |  |
| 6:8.28 | בהם |  |
| 6:8.29 | מהאמונות |  |
| 6:8.30 | האלהיות |  |
| 6:8.31 | הבלתי |  |
| 6:8.32 | מבוארת |  |
| 6:8.33 | לא |  |
| 6:8.34 | בתורה |  |
| 6:8.35 | שבכתב |  |
| 6:8.36 | ולא |  |
| 6:8.37 | בתורה |  |
| 6:8.38 | שבעל |  |
| 6:8.39 | פה | , |
| 6:8.40 | שהמאמין |  |
| 6:8.41 | בהם |  |
| 6:8.42 | שלא |  |
| 6:8.43 | כדת |  |
| 6:8.44 | אינו |  |
| 6:8.45 | בכלל |  |
| 6:8.46 | בעלי |  |
| 6:8.47 | האמונה |  |
| 6:8.48 | השלימה | , |
| 6:8.49 | ולא |  |
| 6:8.50 | אותו |  |
| 6:8.51 | ית | ' |
| 6:8.52 | יעבוד |  |
| 6:8.53 | כמו |  |
| 6:8.54 | שכ | ' |
| 6:8.55 | הרב |  |
| 6:8.56 | בעל |  |
| 6:8.57 | חובת |  |
| 6:8.58 | הלבבות |  |
| 6:8.59 | לא |  |
| 6:8.60 | יוכל |  |
| 6:8.61 | לעבוד |  |
| 6:8.62 | עילת |  |
| 6:8.63 | העילות |  |
| 6:8.64 | זולת |  |
| 6:8.65 | הנביא |  |
| 6:8.66 | או |  |
| 6:8.67 | החכם |  |
| 6:8.68 | השלום | , |
| 6:8.69 | על |  |
| 6:8.70 | כן |  |
| 6:8.71 | המצוה |  |
| 6:8.72 | עלינו |  |
| 6:8.73 | לחקור |  |
| 6:8.74 | על |  |
| 6:8.75 | מצות |  |
| 6:8.76 | אלה |  |
| 6:8.77 | מפי |  |
| 6:8.78 | סופרים |  |
| 6:8.79 | ומספרים |  |
| 6:8.80 | האלהיים |  |
| 6:8.81 | למען |  |
| 6:8.82 | לא |  |
| 6:8.83 | יטעה |  |
| 6:8.84 | באמונתו |  |
| 6:8.85 | עכ"ל | . |


## Milchamot ha-Shem 6:9

Hebrew (Wikisource): סד) ודבריו אלה שכתב שמצות האחדות לא ביארוה רבותינו ז"ל נפלאים בעיני עד מאוד מכמה פנים, כי הנה ידוע הוא ומפורסם שדרך רז"ל בעלי המשנה והתלמוד שלא להזכיר רק דבר שאפשר להסתפק בו אם הוא אסור או מותר, אבל דברי הפשוט איסורו או התירו לא יזכיר התנא כלל, דזיל קרי בי רב הוא, וכן כתב הרמב"ם ז"ל בפי' המשנה פרק ז' דברכות וז"ל ודע כי המשנה לא תדבר דבר באיסור דבר המפורסם וידוע איסורו או בהתר דבר שידוע התירו, אך בדבר שיש לספק בו אם הוא אסור או מותר, תבאר המשנה לאסור דבר שיעלה על הדעת שהוא מותר, או בהתר דבר שיעלה על הדעת שהוא אסור, אבל דברים הידועים לא תדבר בהם כלל, ולכן מקשים בגמרא תמיד פשיטא? מאי קמ"ל? אם היה דבר ידוע ומפורסם עכ"ל. וכבר נתבאר לעיל בדברי הר' הסמ"ג כי אחדותו ית' פשוטה ותקועה בלב כל ישראל, וכל שכן בימי רז"ל בעלי המשנה והתלמוד, והשנית שהיא מבוארת בתורה נביאים וכתובים במקומות אין מספר, שמע ישראל ה' אלהינו ה' אחד, וידעת היום וכו' כי ה' הוא האלהים בשמים ועל הארץ מתחת אין עוד, ואמר אין עוד מלבדו ראו עתה כי אני אני הוא ואין אלהים עמדי, ה' בדד ינחנו ואין עמו אל נכר, נוטה שמים לבדי, רוקע הארץ מיאתי, מי הקדימני, ועוד הרבה כתובים מדברים באחדותו ית' וקדמותו, ומפורסם ג"כ באומה הישראלית כנז"ל ס' ס"א, לכן לא הוצרכו רז"ל לבאר היחוד האמתי, להיותו מן המפורסמות, גם שלילת הגשמות מבוארת בכתובים כי לא ראיתם כל תמונה ותמונה אינכם רואים, ואל מי תדמיוני ואשוה, ואל מי תדמיון אל, והרבה כתובים כיוצא בהם, והשינוי ג"כ מבואר שלילתו מהשי"ת בכתוב הוא אמרו ית' אני ה' לא שניתי, וכן שאר דברים המביאים לענינים הנזכרים, ועם כל זה לא שתקו רז"ל בגמרא בבלי וירושלמי ומדרש רבה ותנחומא ושאר מדרשים הידועים ומפורסמים לרז"ל מלהזכיר תשובת טענות ושאלות שאירעו וקרו להם מן החיצונים הכופרים הטוענים נגד אמונותיו ויש בהם די להורותינו הדרך הנכונה ביחוד השי"ת שהוא כמו שביארו ר' סעדיה גאון והרמב"ם בפי' המשנה ובחבור ובמורה ור' בחיי בחובת הלבבות, והרקח, ושבילי אמונה, ור' יהודה הלוי בכוזרי. וגם הרחיקו רז"ל מלהרבות בתארי השי"ת. ומכל ענין היוצא ממנו משמעות שתי רשויות או יותר כאמרם האומר שמע שמע או מודים מודים וכיוצא משתקים אותו וכמו שנת' לעיל סי' ל"ב. וכ"ש שנמצא את כל אלה על פי הקבלה האמתית ביחוד השי"ת אם נשכיל ונעמיק חרק בתרגומים המסורים לנו מרז"ל תרגום אונקלוס הגר ע"ה שתרגם את התורה מפי ר' אליעזר ור' יהושע, ותרגום יונתן בן עוזיאל הגדול שבתלמיד הלל שתרגם את הנביאים כמו שקבל מהלל. ומצאנו ראינו שהם משרישים האמונה האמתית בלב העם ומתקנים תרגום הכתובים שלא יפול ספק בקבלה האמתית כמו שכ' הרמב"ם והרס"ג. ועיין בהקדמה כוללת להר"ר נתן אדלר בעל נתינה לגר. ועל כן ארז"ל המתרגם פסוק כצורתו הרי זה בדאי. לפי שדרך הכתובים לדבר בדרך הרחבת הלשון, או בלשון מושאל, או משותף, או בלשון קצר, או במוקדם ומאוחר. צריך המתרגם לתרגמו כמו שקבלו חכמים, ואלו הן סודות התורה לפרש מקראות הסתומים כפי האמת שקבלו חכמים. ולכן אמרו רז"ל כי כשתרגם יונתן בן עוזיאל את הנביאים יצאה בת קול ואמרה מי גלה רזי לבני אדם. ועליו יאמר רב יוסף בכמה מקומות אלמלא תרגומא דהאיי קרא לא ידענא מאי קאמר. ובמדרש הגדול פ' משפטים ויראו את אלהי ישראל. אמר ר' אליעזר כל המתרגם פסוק כצורתו ה"ז בדאי וכל המוסיף בו הרי זה מחרף ומגדף כגון ויראו את אלהי ישראל. וחזו ית אלהא דישראל ה"ז בדאי. שהקב"ה רואה ואינו נראה. ואם תרגם וחזו ית יקר שכינת אלהא דישראל הרי זה מחרף ומגדף. שהוא עושה כאן שלשה, יקר, ושכינה ואל עכ"ל. וקאי לי טובא שהרי יונתן בן עוזיאל תרגם כי את המלך ה' צבאות ראו עיני. ארי ית יקר שכינת מלך עלמיא ה' צבאות חזאה עיני והרי"ף בסוף מסכת מגלה הביא תוספתא זו ופי' ר' נסים המתרגם פסוק כצורתו וחזו ית' מלאכא דאלהא דישראל, ורחות כרסי יקריה ה"ז מגדף שהשוה כבוד עבד לכבוד רבו ע"כ, והערוך ערך תרגם פי' הרי מחרף ומגדף שעושה הכבוד מלאך, אלא יתרגם וחזו ית יקר אלהא דישראל וכו'. וכנגד אמונת השלוש שנתחדשה בזמן רז"ל בעלי המשנה מתלמידי ישוע הנצרי אמרו רז"ל במדרש רבה פרשת יתרו אנכי ה' אלהיך אשר הוציאתיך וגו' אמר הקב"ה מלך בשר ודם מולך יש לו אב או אח, או בן, אני איני כן, אני ראשון שאין לי אב, ואני אחרון שאין לי בן, ומבלעדי אין אלהים שאין לי אח ע"כ. והמקובלים החדשים נמשכו אחר דעת בעלי השלוש המאמינים באב ובן ורוח הקודש כמובן מהמשל המבואר בס' עז לאלהים בבית קודש הקדשים דף נ"ט ע"ב וז"ל ולפי שמלכא קדישא דכל קדישין לא מתפרש מניה לעלמין והוא מאיר בזעיר נאפין דהוא לב האצילות לכן הוא ג"כ יקרא בוצינא דקרדינותא שמאיר בלב שהוא זעיר אנפין דאצילות. נמצא שמה שנקרא זה השם הוא על שם נשמתו ולא על שם מלבושו כשאר השמות. והוא הנקרא מלך לבני ישראל כאשר נתבאר בפרשת ואלה המלכים. והוא מלכינו ואנחנו עבדיו. ואליו תפלותינו וכוונותינו. ואפי' אבא ואימא מוזהרים לכבד את הבן הזה וליראה ממנו, לפי שהוא המלך השליט על כלא. ואפי' שיצא הבן הקדוש הזה מאבא ואימא יתבאר לקטן שהוא העיקר, אבל בכל זאת אמשול לך משל אם ימלוך א' ויהיה לו אב ואם. למי צריכים לעבוד? בודאי למלך. ואפי' אביו ואמו צריכים לו לפי שהוא מלך. על זה אמר נשקו בר פן יאנף. ואמר בזהר נשקו ידין להאי ברא קדישא. כי שולטנו אתיהיב ליה על כלא וכו' הרי שאפי' אבא ואימא צריכים לו עכ"ל. הרי בהדיא שהוא זותר לדברי רז"ל במדרש שמות רבה הנז' לעיל, וירושלמי בפ' דיני ממונות, דגרסינן התם מהו חותמו של הקב"ה רב בבי בשם ר ראובן אמת. מהו אמת. אמר ר' בון שהוא אלהים חיים ומלך עולם. אמר ריש לקיש אלף רישא דאלף בית מים באמצעיתא, תיו בסופא. לומר אני ה' ראשון שלא קבלתי מאחר. ומבלעדי אין אלהים שאין לי שותף. ואת אחרונים אני הוא שאיני עתיד למוסרה לאחר עכ"ל הירוש'.

| id | word | punctuation after |
|---|---|---|
| 6:9.1 | סד | ) |
| 6:9.2 | ודבריו |  |
| 6:9.3 | אלה |  |
| 6:9.4 | שכתב |  |
| 6:9.5 | שמצות |  |
| 6:9.6 | האחדות |  |
| 6:9.7 | לא |  |
| 6:9.8 | ביארוה |  |
| 6:9.9 | רבותינו |  |
| 6:9.10 | ז"ל |  |
| 6:9.11 | נפלאים |  |
| 6:9.12 | בעיני |  |
| 6:9.13 | עד |  |
| 6:9.14 | מאוד |  |
| 6:9.15 | מכמה |  |
| 6:9.16 | פנים | , |
| 6:9.17 | כי |  |
| 6:9.18 | הנה |  |
| 6:9.19 | ידוע |  |
| 6:9.20 | הוא |  |
| 6:9.21 | ומפורסם |  |
| 6:9.22 | שדרך |  |
| 6:9.23 | רז"ל |  |
| 6:9.24 | בעלי |  |
| 6:9.25 | המשנה |  |
| 6:9.26 | והתלמוד |  |
| 6:9.27 | שלא |  |
| 6:9.28 | להזכיר |  |
| 6:9.29 | רק |  |
| 6:9.30 | דבר |  |
| 6:9.31 | שאפשר |  |
| 6:9.32 | להסתפק |  |
| 6:9.33 | בו |  |
| 6:9.34 | אם |  |
| 6:9.35 | הוא |  |
| 6:9.36 | אסור |  |
| 6:9.37 | או |  |
| 6:9.38 | מותר | , |
| 6:9.39 | אבל |  |
| 6:9.40 | דברי |  |
| 6:9.41 | הפשוט |  |
| 6:9.42 | איסורו |  |
| 6:9.43 | או |  |
| 6:9.44 | התירו |  |
| 6:9.45 | לא |  |
| 6:9.46 | יזכיר |  |
| 6:9.47 | התנא |  |
| 6:9.48 | כלל | , |
| 6:9.49 | דזיל |  |
| 6:9.50 | קרי |  |
| 6:9.51 | בי |  |
| 6:9.52 | רב |  |
| 6:9.53 | הוא | , |
| 6:9.54 | וכן |  |
| 6:9.55 | כתב |  |
| 6:9.56 | הרמב"ם |  |
| 6:9.57 | ז"ל |  |
| 6:9.58 | בפי | ' |
| 6:9.59 | המשנה |  |
| 6:9.60 | פרק |  |
| 6:9.61 | ז | ' |
| 6:9.62 | דברכות |  |
| 6:9.63 | וז"ל |  |
| 6:9.64 | ודע |  |
| 6:9.65 | כי |  |
| 6:9.66 | המשנה |  |
| 6:9.67 | לא |  |
| 6:9.68 | תדבר |  |
| 6:9.69 | דבר |  |
| 6:9.70 | באיסור |  |
| 6:9.71 | דבר |  |
| 6:9.72 | המפורסם |  |
| 6:9.73 | וידוע |  |
| 6:9.74 | איסורו |  |
| 6:9.75 | או |  |
| 6:9.76 | בהתר |  |
| 6:9.77 | דבר |  |
| 6:9.78 | שידוע |  |
| 6:9.79 | התירו | , |
| 6:9.80 | אך |  |
| 6:9.81 | בדבר |  |
| 6:9.82 | שיש |  |
| 6:9.83 | לספק |  |
| 6:9.84 | בו |  |
| 6:9.85 | אם |  |
| 6:9.86 | הוא |  |
| 6:9.87 | אסור |  |
| 6:9.88 | או |  |
| 6:9.89 | מותר | , |
| 6:9.90 | תבאר |  |
| 6:9.91 | המשנה |  |
| 6:9.92 | לאסור |  |
| 6:9.93 | דבר |  |
| 6:9.94 | שיעלה |  |
| 6:9.95 | על |  |
| 6:9.96 | הדעת |  |
| 6:9.97 | שהוא |  |
| 6:9.98 | מותר | , |
| 6:9.99 | או |  |
| 6:9.100 | בהתר |  |
| 6:9.101 | דבר |  |
| 6:9.102 | שיעלה |  |
| 6:9.103 | על |  |
| 6:9.104 | הדעת |  |
| 6:9.105 | שהוא |  |
| 6:9.106 | אסור | , |
| 6:9.107 | אבל |  |
| 6:9.108 | דברים |  |
| 6:9.109 | הידועים |  |
| 6:9.110 | לא |  |
| 6:9.111 | תדבר |  |
| 6:9.112 | בהם |  |
| 6:9.113 | כלל | , |
| 6:9.114 | ולכן |  |
| 6:9.115 | מקשים |  |
| 6:9.116 | בגמרא |  |
| 6:9.117 | תמיד |  |
| 6:9.118 | פשיטא | ? |
| 6:9.119 | מאי |  |
| 6:9.120 | קמ"ל | ? |
| 6:9.121 | אם |  |
| 6:9.122 | היה |  |
| 6:9.123 | דבר |  |
| 6:9.124 | ידוע |  |
| 6:9.125 | ומפורסם |  |
| 6:9.126 | עכ"ל | . |
| 6:9.127 | וכבר |  |
| 6:9.128 | נתבאר |  |
| 6:9.129 | לעיל |  |
| 6:9.130 | בדברי |  |
| 6:9.131 | הר | ' |
| 6:9.132 | הסמ"ג |  |
| 6:9.133 | כי |  |
| 6:9.134 | אחדותו |  |
| 6:9.135 | ית | ' |
| 6:9.136 | פשוטה |  |
| 6:9.137 | ותקועה |  |
| 6:9.138 | בלב |  |
| 6:9.139 | כל |  |
| 6:9.140 | ישראל | , |
| 6:9.141 | וכל |  |
| 6:9.142 | שכן |  |
| 6:9.143 | בימי |  |
| 6:9.144 | רז"ל |  |
| 6:9.145 | בעלי |  |
| 6:9.146 | המשנה |  |
| 6:9.147 | והתלמוד | , |
| 6:9.148 | והשנית |  |
| 6:9.149 | שהיא |  |
| 6:9.150 | מבוארת |  |
| 6:9.151 | בתורה |  |
| 6:9.152 | נביאים |  |
| 6:9.153 | וכתובים |  |
| 6:9.154 | במקומות |  |
| 6:9.155 | אין |  |
| 6:9.156 | מספר | , |
| 6:9.157 | שמע |  |
| 6:9.158 | ישראל |  |
| 6:9.159 | ה | ' |
| 6:9.160 | אלהינו |  |
| 6:9.161 | ה | ' |
| 6:9.162 | אחד | , |
| 6:9.163 | וידעת |  |
| 6:9.164 | היום |  |
| 6:9.165 | וכו | ' |
| 6:9.166 | כי |  |
| 6:9.167 | ה | ' |
| 6:9.168 | הוא |  |
| 6:9.169 | האלהים |  |
| 6:9.170 | בשמים |  |
| 6:9.171 | ועל |  |
| 6:9.172 | הארץ |  |
| 6:9.173 | מתחת |  |
| 6:9.174 | אין |  |
| 6:9.175 | עוד | , |
| 6:9.176 | ואמר |  |
| 6:9.177 | אין |  |
| 6:9.178 | עוד |  |
| 6:9.179 | מלבדו |  |
| 6:9.180 | ראו |  |
| 6:9.181 | עתה |  |
| 6:9.182 | כי |  |
| 6:9.183 | אני |  |
| 6:9.184 | אני |  |
| 6:9.185 | הוא |  |
| 6:9.186 | ואין |  |
| 6:9.187 | אלהים |  |
| 6:9.188 | עמדי | , |
| 6:9.189 | ה | ' |
| 6:9.190 | בדד |  |
| 6:9.191 | ינחנו |  |
| 6:9.192 | ואין |  |
| 6:9.193 | עמו |  |
| 6:9.194 | אל |  |
| 6:9.195 | נכר | , |
| 6:9.196 | נוטה |  |
| 6:9.197 | שמים |  |
| 6:9.198 | לבדי | , |
| 6:9.199 | רוקע |  |
| 6:9.200 | הארץ |  |
| 6:9.201 | מיאתי | , |
| 6:9.202 | מי |  |
| 6:9.203 | הקדימני | , |
| 6:9.204 | ועוד |  |
| 6:9.205 | הרבה |  |
| 6:9.206 | כתובים |  |
| 6:9.207 | מדברים |  |
| 6:9.208 | באחדותו |  |
| 6:9.209 | ית | ' |
| 6:9.210 | וקדמותו | , |
| 6:9.211 | ומפורסם |  |
| 6:9.212 | ג"כ |  |
| 6:9.213 | באומה |  |
| 6:9.214 | הישראלית |  |
| 6:9.215 | כנז"ל |  |
| 6:9.216 | ס | ' |
| 6:9.217 | ס"א | , |
| 6:9.218 | לכן |  |
| 6:9.219 | לא |  |
| 6:9.220 | הוצרכו |  |
| 6:9.221 | רז"ל |  |
| 6:9.222 | לבאר |  |
| 6:9.223 | היחוד |  |
| 6:9.224 | האמתי | , |
| 6:9.225 | להיותו |  |
| 6:9.226 | מן |  |
| 6:9.227 | המפורסמות | , |
| 6:9.228 | גם |  |
| 6:9.229 | שלילת |  |
| 6:9.230 | הגשמות |  |
| 6:9.231 | מבוארת |  |
| 6:9.232 | בכתובים |  |
| 6:9.233 | כי |  |
| 6:9.234 | לא |  |
| 6:9.235 | ראיתם |  |
| 6:9.236 | כל |  |
| 6:9.237 | תמונה |  |
| 6:9.238 | ותמונה |  |
| 6:9.239 | אינכם |  |
| 6:9.240 | רואים | , |
| 6:9.241 | ואל |  |
| 6:9.242 | מי |  |
| 6:9.243 | תדמיוני |  |
| 6:9.244 | ואשוה | , |
| 6:9.245 | ואל |  |
| 6:9.246 | מי |  |
| 6:9.247 | תדמיון |  |
| 6:9.248 | אל | , |
| 6:9.249 | והרבה |  |
| 6:9.250 | כתובים |  |
| 6:9.251 | כיוצא |  |
| 6:9.252 | בהם | , |
| 6:9.253 | והשינוי |  |
| 6:9.254 | ג"כ |  |
| 6:9.255 | מבואר |  |
| 6:9.256 | שלילתו |  |
| 6:9.257 | מהשי"ת |  |
| 6:9.258 | בכתוב |  |
| 6:9.259 | הוא |  |
| 6:9.260 | אמרו |  |
| 6:9.261 | ית | ' |
| 6:9.262 | אני |  |
| 6:9.263 | ה | ' |
| 6:9.264 | לא |  |
| 6:9.265 | שניתי | , |
| 6:9.266 | וכן |  |
| 6:9.267 | שאר |  |
| 6:9.268 | דברים |  |
| 6:9.269 | המביאים |  |
| 6:9.270 | לענינים |  |
| 6:9.271 | הנזכרים | , |
| 6:9.272 | ועם |  |
| 6:9.273 | כל |  |
| 6:9.274 | זה |  |
| 6:9.275 | לא |  |
| 6:9.276 | שתקו |  |
| 6:9.277 | רז"ל |  |
| 6:9.278 | בגמרא |  |
| 6:9.279 | בבלי |  |
| 6:9.280 | וירושלמי |  |
| 6:9.281 | ומדרש |  |
| 6:9.282 | רבה |  |
| 6:9.283 | ותנחומא |  |
| 6:9.284 | ושאר |  |
| 6:9.285 | מדרשים |  |
| 6:9.286 | הידועים |  |
| 6:9.287 | ומפורסמים |  |
| 6:9.288 | לרז"ל |  |
| 6:9.289 | מלהזכיר |  |
| 6:9.290 | תשובת |  |
| 6:9.291 | טענות |  |
| 6:9.292 | ושאלות |  |
| 6:9.293 | שאירעו |  |
| 6:9.294 | וקרו |  |
| 6:9.295 | להם |  |
| 6:9.296 | מן |  |
| 6:9.297 | החיצונים |  |
| 6:9.298 | הכופרים |  |
| 6:9.299 | הטוענים |  |
| 6:9.300 | נגד |  |
| 6:9.301 | אמונותיו |  |
| 6:9.302 | ויש |  |
| 6:9.303 | בהם |  |
| 6:9.304 | די |  |
| 6:9.305 | להורותינו |  |
| 6:9.306 | הדרך |  |
| 6:9.307 | הנכונה |  |
| 6:9.308 | ביחוד |  |
| 6:9.309 | השי"ת |  |
| 6:9.310 | שהוא |  |
| 6:9.311 | כמו |  |
| 6:9.312 | שביארו |  |
| 6:9.313 | ר | ' |
| 6:9.314 | סעדיה |  |
| 6:9.315 | גאון |  |
| 6:9.316 | והרמב"ם |  |
| 6:9.317 | בפי | ' |
| 6:9.318 | המשנה |  |
| 6:9.319 | ובחבור |  |
| 6:9.320 | ובמורה |  |
| 6:9.321 | ור | ' |
| 6:9.322 | בחיי |  |
| 6:9.323 | בחובת |  |
| 6:9.324 | הלבבות | , |
| 6:9.325 | והרקח | , |
| 6:9.326 | ושבילי |  |
| 6:9.327 | אמונה | , |
| 6:9.328 | ור | ' |
| 6:9.329 | יהודה |  |
| 6:9.330 | הלוי |  |
| 6:9.331 | בכוזרי | . |
| 6:9.332 | וגם |  |
| 6:9.333 | הרחיקו |  |
| 6:9.334 | רז"ל |  |
| 6:9.335 | מלהרבות |  |
| 6:9.336 | בתארי |  |
| 6:9.337 | השי"ת | . |
| 6:9.338 | ומכל |  |
| 6:9.339 | ענין |  |
| 6:9.340 | היוצא |  |
| 6:9.341 | ממנו |  |
| 6:9.342 | משמעות |  |
| 6:9.343 | שתי |  |
| 6:9.344 | רשויות |  |
| 6:9.345 | או |  |
| 6:9.346 | יותר |  |
| 6:9.347 | כאמרם |  |
| 6:9.348 | האומר |  |
| 6:9.349 | שמע |  |
| 6:9.350 | שמע |  |
| 6:9.351 | או |  |
| 6:9.352 | מודים |  |
| 6:9.353 | מודים |  |
| 6:9.354 | וכיוצא |  |
| 6:9.355 | משתקים |  |
| 6:9.356 | אותו |  |
| 6:9.357 | וכמו |  |
| 6:9.358 | שנת | ' |
| 6:9.359 | לעיל |  |
| 6:9.360 | סי | ' |
| 6:9.361 | ל"ב | . |
| 6:9.362 | וכ"ש |  |
| 6:9.363 | שנמצא |  |
| 6:9.364 | את |  |
| 6:9.365 | כל |  |
| 6:9.366 | אלה |  |
| 6:9.367 | על |  |
| 6:9.368 | פי |  |
| 6:9.369 | הקבלה |  |
| 6:9.370 | האמתית |  |
| 6:9.371 | ביחוד |  |
| 6:9.372 | השי"ת |  |
| 6:9.373 | אם |  |
| 6:9.374 | נשכיל |  |
| 6:9.375 | ונעמיק |  |
| 6:9.376 | חרק |  |
| 6:9.377 | בתרגומים |  |
| 6:9.378 | המסורים |  |
| 6:9.379 | לנו |  |
| 6:9.380 | מרז"ל |  |
| 6:9.381 | תרגום |  |
| 6:9.382 | אונקלוס |  |
| 6:9.383 | הגר |  |
| 6:9.384 | ע"ה |  |
| 6:9.385 | שתרגם |  |
| 6:9.386 | את |  |
| 6:9.387 | התורה |  |
| 6:9.388 | מפי |  |
| 6:9.389 | ר | ' |
| 6:9.390 | אליעזר |  |
| 6:9.391 | ור | ' |
| 6:9.392 | יהושע | , |
| 6:9.393 | ותרגום |  |
| 6:9.394 | יונתן |  |
| 6:9.395 | בן |  |
| 6:9.396 | עוזיאל |  |
| 6:9.397 | הגדול |  |
| 6:9.398 | שבתלמיד |  |
| 6:9.399 | הלל |  |
| 6:9.400 | שתרגם |  |
| 6:9.401 | את |  |
| 6:9.402 | הנביאים |  |
| 6:9.403 | כמו |  |
| 6:9.404 | שקבל |  |
| 6:9.405 | מהלל | . |
| 6:9.406 | ומצאנו |  |
| 6:9.407 | ראינו |  |
| 6:9.408 | שהם |  |
| 6:9.409 | משרישים |  |
| 6:9.410 | האמונה |  |
| 6:9.411 | האמתית |  |
| 6:9.412 | בלב |  |
| 6:9.413 | העם |  |
| 6:9.414 | ומתקנים |  |
| 6:9.415 | תרגום |  |
| 6:9.416 | הכתובים |  |
| 6:9.417 | שלא |  |
| 6:9.418 | יפול |  |
| 6:9.419 | ספק |  |
| 6:9.420 | בקבלה |  |
| 6:9.421 | האמתית |  |
| 6:9.422 | כמו |  |
| 6:9.423 | שכ | ' |
| 6:9.424 | הרמב"ם |  |
| 6:9.425 | והרס"ג | . |
| 6:9.426 | ועיין |  |
| 6:9.427 | בהקדמה |  |
| 6:9.428 | כוללת |  |
| 6:9.429 | להר"ר |  |
| 6:9.430 | נתן |  |
| 6:9.431 | אדלר |  |
| 6:9.432 | בעל |  |
| 6:9.433 | נתינה |  |
| 6:9.434 | לגר | . |
| 6:9.435 | ועל |  |
| 6:9.436 | כן |  |
| 6:9.437 | ארז"ל |  |
| 6:9.438 | המתרגם |  |
| 6:9.439 | פסוק |  |
| 6:9.440 | כצורתו |  |
| 6:9.441 | הרי |  |
| 6:9.442 | זה |  |
| 6:9.443 | בדאי | . |
| 6:9.444 | לפי |  |
| 6:9.445 | שדרך |  |
| 6:9.446 | הכתובים |  |
| 6:9.447 | לדבר |  |
| 6:9.448 | בדרך |  |
| 6:9.449 | הרחבת |  |
| 6:9.450 | הלשון | , |
| 6:9.451 | או |  |
| 6:9.452 | בלשון |  |
| 6:9.453 | מושאל | , |
| 6:9.454 | או |  |
| 6:9.455 | משותף | , |
| 6:9.456 | או |  |
| 6:9.457 | בלשון |  |
| 6:9.458 | קצר | , |
| 6:9.459 | או |  |
| 6:9.460 | במוקדם |  |
| 6:9.461 | ומאוחר | . |
| 6:9.462 | צריך |  |
| 6:9.463 | המתרגם |  |
| 6:9.464 | לתרגמו |  |
| 6:9.465 | כמו |  |
| 6:9.466 | שקבלו |  |
| 6:9.467 | חכמים | , |
| 6:9.468 | ואלו |  |
| 6:9.469 | הן |  |
| 6:9.470 | סודות |  |
| 6:9.471 | התורה |  |
| 6:9.472 | לפרש |  |
| 6:9.473 | מקראות |  |
| 6:9.474 | הסתומים |  |
| 6:9.475 | כפי |  |
| 6:9.476 | האמת |  |
| 6:9.477 | שקבלו |  |
| 6:9.478 | חכמים | . |
| 6:9.479 | ולכן |  |
| 6:9.480 | אמרו |  |
| 6:9.481 | רז"ל |  |
| 6:9.482 | כי |  |
| 6:9.483 | כשתרגם |  |
| 6:9.484 | יונתן |  |
| 6:9.485 | בן |  |
| 6:9.486 | עוזיאל |  |
| 6:9.487 | את |  |
| 6:9.488 | הנביאים |  |
| 6:9.489 | יצאה |  |
| 6:9.490 | בת |  |
| 6:9.491 | קול |  |
| 6:9.492 | ואמרה |  |
| 6:9.493 | מי |  |
| 6:9.494 | גלה |  |
| 6:9.495 | רזי |  |
| 6:9.496 | לבני |  |
| 6:9.497 | אדם | . |
| 6:9.498 | ועליו |  |
| 6:9.499 | יאמר |  |
| 6:9.500 | רב |  |
| 6:9.501 | יוסף |  |
| 6:9.502 | בכמה |  |
| 6:9.503 | מקומות |  |
| 6:9.504 | אלמלא |  |
| 6:9.505 | תרגומא |  |
| 6:9.506 | דהאיי |  |
| 6:9.507 | קרא |  |
| 6:9.508 | לא |  |
| 6:9.509 | ידענא |  |
| 6:9.510 | מאי |  |
| 6:9.511 | קאמר | . |
| 6:9.512 | ובמדרש |  |
| 6:9.513 | הגדול |  |
| 6:9.514 | פ | ' |
| 6:9.515 | משפטים |  |
| 6:9.516 | ויראו |  |
| 6:9.517 | את |  |
| 6:9.518 | אלהי |  |
| 6:9.519 | ישראל | . |
| 6:9.520 | אמר |  |
| 6:9.521 | ר | ' |
| 6:9.522 | אליעזר |  |
| 6:9.523 | כל |  |
| 6:9.524 | המתרגם |  |
| 6:9.525 | פסוק |  |
| 6:9.526 | כצורתו |  |
| 6:9.527 | ה"ז |  |
| 6:9.528 | בדאי |  |
| 6:9.529 | וכל |  |
| 6:9.530 | המוסיף |  |
| 6:9.531 | בו |  |
| 6:9.532 | הרי |  |
| 6:9.533 | זה |  |
| 6:9.534 | מחרף |  |
| 6:9.535 | ומגדף |  |
| 6:9.536 | כגון |  |
| 6:9.537 | ויראו |  |
| 6:9.538 | את |  |
| 6:9.539 | אלהי |  |
| 6:9.540 | ישראל | . |
| 6:9.541 | וחזו |  |
| 6:9.542 | ית |  |
| 6:9.543 | אלהא |  |
| 6:9.544 | דישראל |  |
| 6:9.545 | ה"ז |  |
| 6:9.546 | בדאי | . |
| 6:9.547 | שהקב"ה |  |
| 6:9.548 | רואה |  |
| 6:9.549 | ואינו |  |
| 6:9.550 | נראה | . |
| 6:9.551 | ואם |  |
| 6:9.552 | תרגם |  |
| 6:9.553 | וחזו |  |
| 6:9.554 | ית |  |
| 6:9.555 | יקר |  |
| 6:9.556 | שכינת |  |
| 6:9.557 | אלהא |  |
| 6:9.558 | דישראל |  |
| 6:9.559 | הרי |  |
| 6:9.560 | זה |  |
| 6:9.561 | מחרף |  |
| 6:9.562 | ומגדף | . |
| 6:9.563 | שהוא |  |
| 6:9.564 | עושה |  |
| 6:9.565 | כאן |  |
| 6:9.566 | שלשה | , |
| 6:9.567 | יקר | , |
| 6:9.568 | ושכינה |  |
| 6:9.569 | ואל |  |
| 6:9.570 | עכ"ל | . |
| 6:9.571 | וקאי |  |
| 6:9.572 | לי |  |
| 6:9.573 | טובא |  |
| 6:9.574 | שהרי |  |
| 6:9.575 | יונתן |  |
| 6:9.576 | בן |  |
| 6:9.577 | עוזיאל |  |
| 6:9.578 | תרגם |  |
| 6:9.579 | כי |  |
| 6:9.580 | את |  |
| 6:9.581 | המלך |  |
| 6:9.582 | ה | ' |
| 6:9.583 | צבאות |  |
| 6:9.584 | ראו |  |
| 6:9.585 | עיני | . |
| 6:9.586 | ארי |  |
| 6:9.587 | ית |  |
| 6:9.588 | יקר |  |
| 6:9.589 | שכינת |  |
| 6:9.590 | מלך |  |
| 6:9.591 | עלמיא |  |
| 6:9.592 | ה | ' |
| 6:9.593 | צבאות |  |
| 6:9.594 | חזאה |  |
| 6:9.595 | עיני |  |
| 6:9.596 | והרי"ף |  |
| 6:9.597 | בסוף |  |
| 6:9.598 | מסכת |  |
| 6:9.599 | מגלה |  |
| 6:9.600 | הביא |  |
| 6:9.601 | תוספתא |  |
| 6:9.602 | זו |  |
| 6:9.603 | ופי | ' |
| 6:9.604 | ר | ' |
| 6:9.605 | נסים |  |
| 6:9.606 | המתרגם |  |
| 6:9.607 | פסוק |  |
| 6:9.608 | כצורתו |  |
| 6:9.609 | וחזו |  |
| 6:9.610 | ית | ' |
| 6:9.611 | מלאכא |  |
| 6:9.612 | דאלהא |  |
| 6:9.613 | דישראל | , |
| 6:9.614 | ורחות |  |
| 6:9.615 | כרסי |  |
| 6:9.616 | יקריה |  |
| 6:9.617 | ה"ז |  |
| 6:9.618 | מגדף |  |
| 6:9.619 | שהשוה |  |
| 6:9.620 | כבוד |  |
| 6:9.621 | עבד |  |
| 6:9.622 | לכבוד |  |
| 6:9.623 | רבו |  |
| 6:9.624 | ע"כ | , |
| 6:9.625 | והערוך |  |
| 6:9.626 | ערך |  |
| 6:9.627 | תרגם |  |
| 6:9.628 | פי | ' |
| 6:9.629 | הרי |  |
| 6:9.630 | מחרף |  |
| 6:9.631 | ומגדף |  |
| 6:9.632 | שעושה |  |
| 6:9.633 | הכבוד |  |
| 6:9.634 | מלאך | , |
| 6:9.635 | אלא |  |
| 6:9.636 | יתרגם |  |
| 6:9.637 | וחזו |  |
| 6:9.638 | ית |  |
| 6:9.639 | יקר |  |
| 6:9.640 | אלהא |  |
| 6:9.641 | דישראל |  |
| 6:9.642 | וכו | '. |
| 6:9.643 | וכנגד |  |
| 6:9.644 | אמונת |  |
| 6:9.645 | השלוש |  |
| 6:9.646 | שנתחדשה |  |
| 6:9.647 | בזמן |  |
| 6:9.648 | רז"ל |  |
| 6:9.649 | בעלי |  |
| 6:9.650 | המשנה |  |
| 6:9.651 | מתלמידי |  |
| 6:9.652 | ישוע |  |
| 6:9.653 | הנצרי |  |
| 6:9.654 | אמרו |  |
| 6:9.655 | רז"ל |  |
| 6:9.656 | במדרש |  |
| 6:9.657 | רבה |  |
| 6:9.658 | פרשת |  |
| 6:9.659 | יתרו |  |
| 6:9.660 | אנכי |  |
| 6:9.661 | ה | ' |
| 6:9.662 | אלהיך |  |
| 6:9.663 | אשר |  |
| 6:9.664 | הוציאתיך |  |
| 6:9.665 | וגו | ' |
| 6:9.666 | אמר |  |
| 6:9.667 | הקב"ה |  |
| 6:9.668 | מלך |  |
| 6:9.669 | בשר |  |
| 6:9.670 | ודם |  |
| 6:9.671 | מולך |  |
| 6:9.672 | יש |  |
| 6:9.673 | לו |  |
| 6:9.674 | אב |  |
| 6:9.675 | או |  |
| 6:9.676 | אח | , |
| 6:9.677 | או |  |
| 6:9.678 | בן | , |
| 6:9.679 | אני |  |
| 6:9.680 | איני |  |
| 6:9.681 | כן | , |
| 6:9.682 | אני |  |
| 6:9.683 | ראשון |  |
| 6:9.684 | שאין |  |
| 6:9.685 | לי |  |
| 6:9.686 | אב | , |
| 6:9.687 | ואני |  |
| 6:9.688 | אחרון |  |
| 6:9.689 | שאין |  |
| 6:9.690 | לי |  |
| 6:9.691 | בן | , |
| 6:9.692 | ומבלעדי |  |
| 6:9.693 | אין |  |
| 6:9.694 | אלהים |  |
| 6:9.695 | שאין |  |
| 6:9.696 | לי |  |
| 6:9.697 | אח |  |
| 6:9.698 | ע"כ | . |
| 6:9.699 | והמקובלים |  |
| 6:9.700 | החדשים |  |
| 6:9.701 | נמשכו |  |
| 6:9.702 | אחר |  |
| 6:9.703 | דעת |  |
| 6:9.704 | בעלי |  |
| 6:9.705 | השלוש |  |
| 6:9.706 | המאמינים |  |
| 6:9.707 | באב |  |
| 6:9.708 | ובן |  |
| 6:9.709 | ורוח |  |
| 6:9.710 | הקודש |  |
| 6:9.711 | כמובן |  |
| 6:9.712 | מהמשל |  |
| 6:9.713 | המבואר |  |
| 6:9.714 | בס | ' |
| 6:9.715 | עז |  |
| 6:9.716 | לאלהים |  |
| 6:9.717 | בבית |  |
| 6:9.718 | קודש |  |
| 6:9.719 | הקדשים |  |
| 6:9.720 | דף |  |
| 6:9.721 | נ"ט |  |
| 6:9.722 | ע"ב |  |
| 6:9.723 | וז"ל |  |
| 6:9.724 | ולפי |  |
| 6:9.725 | שמלכא |  |
| 6:9.726 | קדישא |  |
| 6:9.727 | דכל |  |
| 6:9.728 | קדישין |  |
| 6:9.729 | לא |  |
| 6:9.730 | מתפרש |  |
| 6:9.731 | מניה |  |
| 6:9.732 | לעלמין |  |
| 6:9.733 | והוא |  |
| 6:9.734 | מאיר |  |
| 6:9.735 | בזעיר |  |
| 6:9.736 | נאפין |  |
| 6:9.737 | דהוא |  |
| 6:9.738 | לב |  |
| 6:9.739 | האצילות |  |
| 6:9.740 | לכן |  |
| 6:9.741 | הוא |  |
| 6:9.742 | ג"כ |  |
| 6:9.743 | יקרא |  |
| 6:9.744 | בוצינא |  |
| 6:9.745 | דקרדינותא |  |
| 6:9.746 | שמאיר |  |
| 6:9.747 | בלב |  |
| 6:9.748 | שהוא |  |
| 6:9.749 | זעיר |  |
| 6:9.750 | אנפין |  |
| 6:9.751 | דאצילות | . |
| 6:9.752 | נמצא |  |
| 6:9.753 | שמה |  |
| 6:9.754 | שנקרא |  |
| 6:9.755 | זה |  |
| 6:9.756 | השם |  |
| 6:9.757 | הוא |  |
| 6:9.758 | על |  |
| 6:9.759 | שם |  |
| 6:9.760 | נשמתו |  |
| 6:9.761 | ולא |  |
| 6:9.762 | על |  |
| 6:9.763 | שם |  |
| 6:9.764 | מלבושו |  |
| 6:9.765 | כשאר |  |
| 6:9.766 | השמות | . |
| 6:9.767 | והוא |  |
| 6:9.768 | הנקרא |  |
| 6:9.769 | מלך |  |
| 6:9.770 | לבני |  |
| 6:9.771 | ישראל |  |
| 6:9.772 | כאשר |  |
| 6:9.773 | נתבאר |  |
| 6:9.774 | בפרשת |  |
| 6:9.775 | ואלה |  |
| 6:9.776 | המלכים | . |
| 6:9.777 | והוא |  |
| 6:9.778 | מלכינו |  |
| 6:9.779 | ואנחנו |  |
| 6:9.780 | עבדיו | . |
| 6:9.781 | ואליו |  |
| 6:9.782 | תפלותינו |  |
| 6:9.783 | וכוונותינו | . |
| 6:9.784 | ואפי | ' |
| 6:9.785 | אבא |  |
| 6:9.786 | ואימא |  |
| 6:9.787 | מוזהרים |  |
| 6:9.788 | לכבד |  |
| 6:9.789 | את |  |
| 6:9.790 | הבן |  |
| 6:9.791 | הזה |  |
| 6:9.792 | וליראה |  |
| 6:9.793 | ממנו | , |
| 6:9.794 | לפי |  |
| 6:9.795 | שהוא |  |
| 6:9.796 | המלך |  |
| 6:9.797 | השליט |  |
| 6:9.798 | על |  |
| 6:9.799 | כלא | . |
| 6:9.800 | ואפי | ' |
| 6:9.801 | שיצא |  |
| 6:9.802 | הבן |  |
| 6:9.803 | הקדוש |  |
| 6:9.804 | הזה |  |
| 6:9.805 | מאבא |  |
| 6:9.806 | ואימא |  |
| 6:9.807 | יתבאר |  |
| 6:9.808 | לקטן |  |
| 6:9.809 | שהוא |  |
| 6:9.810 | העיקר | , |
| 6:9.811 | אבל |  |
| 6:9.812 | בכל |  |
| 6:9.813 | זאת |  |
| 6:9.814 | אמשול |  |
| 6:9.815 | לך |  |
| 6:9.816 | משל |  |
| 6:9.817 | אם |  |
| 6:9.818 | ימלוך |  |
| 6:9.819 | א | ' |
| 6:9.820 | ויהיה |  |
| 6:9.821 | לו |  |
| 6:9.822 | אב |  |
| 6:9.823 | ואם | . |
| 6:9.824 | למי |  |
| 6:9.825 | צריכים |  |
| 6:9.826 | לעבוד | ? |
| 6:9.827 | בודאי |  |
| 6:9.828 | למלך | . |
| 6:9.829 | ואפי | ' |
| 6:9.830 | אביו |  |
| 6:9.831 | ואמו |  |
| 6:9.832 | צריכים |  |
| 6:9.833 | לו |  |
| 6:9.834 | לפי |  |
| 6:9.835 | שהוא |  |
| 6:9.836 | מלך | . |
| 6:9.837 | על |  |
| 6:9.838 | זה |  |
| 6:9.839 | אמר |  |
| 6:9.840 | נשקו |  |
| 6:9.841 | בר |  |
| 6:9.842 | פן |  |
| 6:9.843 | יאנף | . |
| 6:9.844 | ואמר |  |
| 6:9.845 | בזהר |  |
| 6:9.846 | נשקו |  |
| 6:9.847 | ידין |  |
| 6:9.848 | להאי |  |
| 6:9.849 | ברא |  |
| 6:9.850 | קדישא | . |
| 6:9.851 | כי |  |
| 6:9.852 | שולטנו |  |
| 6:9.853 | אתיהיב |  |
| 6:9.854 | ליה |  |
| 6:9.855 | על |  |
| 6:9.856 | כלא |  |
| 6:9.857 | וכו | ' |
| 6:9.858 | הרי |  |
| 6:9.859 | שאפי | ' |
| 6:9.860 | אבא |  |
| 6:9.861 | ואימא |  |
| 6:9.862 | צריכים |  |
| 6:9.863 | לו |  |
| 6:9.864 | עכ"ל | . |
| 6:9.865 | הרי |  |
| 6:9.866 | בהדיא |  |
| 6:9.867 | שהוא |  |
| 6:9.868 | זותר |  |
| 6:9.869 | לדברי |  |
| 6:9.870 | רז"ל |  |
| 6:9.871 | במדרש |  |
| 6:9.872 | שמות |  |
| 6:9.873 | רבה |  |
| 6:9.874 | הנז | ' |
| 6:9.875 | לעיל | , |
| 6:9.876 | וירושלמי |  |
| 6:9.877 | בפ | ' |
| 6:9.878 | דיני |  |
| 6:9.879 | ממונות | , |
| 6:9.880 | דגרסינן |  |
| 6:9.881 | התם |  |
| 6:9.882 | מהו |  |
| 6:9.883 | חותמו |  |
| 6:9.884 | של |  |
| 6:9.885 | הקב"ה |  |
| 6:9.886 | רב |  |
| 6:9.887 | בבי |  |
| 6:9.888 | בשם |  |
| 6:9.889 | ר |  |
| 6:9.890 | ראובן |  |
| 6:9.891 | אמת | . |
| 6:9.892 | מהו |  |
| 6:9.893 | אמת | . |
| 6:9.894 | אמר |  |
| 6:9.895 | ר | ' |
| 6:9.896 | בון |  |
| 6:9.897 | שהוא |  |
| 6:9.898 | אלהים |  |
| 6:9.899 | חיים |  |
| 6:9.900 | ומלך |  |
| 6:9.901 | עולם | . |
| 6:9.902 | אמר |  |
| 6:9.903 | ריש |  |
| 6:9.904 | לקיש |  |
| 6:9.905 | אלף |  |
| 6:9.906 | רישא |  |
| 6:9.907 | דאלף |  |
| 6:9.908 | בית |  |
| 6:9.909 | מים |  |
| 6:9.910 | באמצעיתא | , |
| 6:9.911 | תיו |  |
| 6:9.912 | בסופא | . |
| 6:9.913 | לומר |  |
| 6:9.914 | אני |  |
| 6:9.915 | ה | ' |
| 6:9.916 | ראשון |  |
| 6:9.917 | שלא |  |
| 6:9.918 | קבלתי |  |
| 6:9.919 | מאחר | . |
| 6:9.920 | ומבלעדי |  |
| 6:9.921 | אין |  |
| 6:9.922 | אלהים |  |
| 6:9.923 | שאין |  |
| 6:9.924 | לי |  |
| 6:9.925 | שותף | . |
| 6:9.926 | ואת |  |
| 6:9.927 | אחרונים |  |
| 6:9.928 | אני |  |
| 6:9.929 | הוא |  |
| 6:9.930 | שאיני |  |
| 6:9.931 | עתיד |  |
| 6:9.932 | למוסרה |  |
| 6:9.933 | לאחר |  |
| 6:9.934 | עכ"ל |  |
| 6:9.935 | הירוש | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

