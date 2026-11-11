# Translation pack: milhamot-hashem perek-9-13 (Milchamot ha-Shem 9:16–9:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-13`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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


## Milchamot ha-Shem 9:16

Hebrew (Wikisource): קלד) מכאן למדנו ממלתיה דרבא שכך היא דרכו של מסית שמפליג בשבח האלילים הרחוקים שכל היא מושכת השפע אליה, וכך חיה מטיבה לעובד אותה להשפיע מטובה שמשכה אליה לזה שעובדה ומשבחה ומפארה, וכך מריעה לפונים אליה עורף ולא השחין לה. ותורתינו הקדושה העירה אותנו באמרה הקרובים אליך או הרחוקים, שמהקרובים שהם פרצופי השמש והירח ושבתי וצדק ומאדים ונגה וכוכב, שבדו חוברי שמים עובדיהם כמו שכתב הרמב"ם ורבי עובדיה בפ' כל הצלמים שיחסו לכוכבים פרצופים, עד שאמרו פרצוף שבתי זקן שחור וישיש. ופרצוף החמה פרצוף מלך מוכתר יושב בעגלה, ופרצוף נגה נערה יפה לבושה חלי זהב, וכן מיחסים לכל הכוכבים והמזלות פרצופים רבים וכו'! מהפרצופים של הכוכבים הקרובים אלינו, נקיש ונלמד על הפרצופים של ספירות הרחוקים ממנו אשר בדא מחבר הזהר המסית אותנו לעבדם ולקרותם בשמותם של הקב"ה שאין בהם ממש! ועל כן לא נאבה לו ולא נשמע אליו, כמצות תורתינו הקדושה שלא לשמוע אל המסית שמטעים את דבריו ומשבח את העושים כמעשיו לשתף ולצרף עם אלהינו ב"ה את הפרצופים הנזכרים, ולהיות העבודה ביחוד לקצר האפים בצירוף שאר פרצופים. וליחדם בפיהם שהבל א', והוסיף על השלשה בשמע ישראל גם את את המלכות שהיא הרביעית עם השלשה שמות יהוה אלהינו יהוה. כדברי השלישים, ובמלת אחד מצורף את המלכוץ עמהם. ואחריו נמשכו המון המקובלים החדשים בתם לבבם ולא נתנו אל לבם להבין ששטתו זאת מילדי נכרים המתנגדים לאמונתינו ביחוד השי"ת, והאמינו לו בחלק שפתיו המורגל בפיו תמיד לומר זכאה חולקיה. דכאה חולקיהון דישראל, כדי למשוך ולגנוב דעת השומעים להדיחם, והבטיח שלעתיד קרוב לימות המשיח כלא יתפרנסון מחאי הבורא דיליה, שיבטלו כל הלכות של תורה שבעל פה, היפך דברי רז"ל שאמרו שכל הלכות של תורה של פה אינם בטילים לעולם (רמב"ם סוף הל' מגלה). ולא די זה אלא שהמסית הנזכר שם את רשב"י תנא דידן לעובר על דברי חכמים ח"ו שאמרו ולא במרכבה ואפי' ביחיד אלא א"כ היה חכם ומבין מדעתו מוסרין לו ראשי פרקים. ולפי דבריו שאלו הם מעשה מרכבה שהזכירו רז"ל, איך היה רשב"י דורש בהם בפומפי בין חביריו מה שלא התירו חכמים לעשות כן אפי' היו כולן חכמים גדולים? ואם תאמר שעתיקא קדושא נגלה אליו כמבואר בזהר בראשית, הא ודאי לא צייתנא ליה לעקור דבר מן התורה שקבלו חכמים, שאין נביא רשאי לחדש דבר מעתה, ובודאי נביא שקר הוא, כמו שמבואר בתלמוד בכמה מקומות שאם יבוא אליהו ויאמר וכו' אין שומעין לו, וזה אפי' כמידי דגזרו ביה רבנן, כ"ש וק"ו לאיסור עצמי שקבלו חכמים בו שאסור לעשות כן שאפילו יבוא אליהו ויאמר שמותר הוא שאין שומעין לו.

| id | word | punctuation after |
|---|---|---|
| 9:16.1 | קלד | ) |
| 9:16.2 | מכאן |  |
| 9:16.3 | למדנו |  |
| 9:16.4 | ממלתיה |  |
| 9:16.5 | דרבא |  |
| 9:16.6 | שכך |  |
| 9:16.7 | היא |  |
| 9:16.8 | דרכו |  |
| 9:16.9 | של |  |
| 9:16.10 | מסית |  |
| 9:16.11 | שמפליג |  |
| 9:16.12 | בשבח |  |
| 9:16.13 | האלילים |  |
| 9:16.14 | הרחוקים |  |
| 9:16.15 | שכל |  |
| 9:16.16 | היא |  |
| 9:16.17 | מושכת |  |
| 9:16.18 | השפע |  |
| 9:16.19 | אליה | , |
| 9:16.20 | וכך |  |
| 9:16.21 | חיה |  |
| 9:16.22 | מטיבה |  |
| 9:16.23 | לעובד |  |
| 9:16.24 | אותה |  |
| 9:16.25 | להשפיע |  |
| 9:16.26 | מטובה |  |
| 9:16.27 | שמשכה |  |
| 9:16.28 | אליה |  |
| 9:16.29 | לזה |  |
| 9:16.30 | שעובדה |  |
| 9:16.31 | ומשבחה |  |
| 9:16.32 | ומפארה | , |
| 9:16.33 | וכך |  |
| 9:16.34 | מריעה |  |
| 9:16.35 | לפונים |  |
| 9:16.36 | אליה |  |
| 9:16.37 | עורף |  |
| 9:16.38 | ולא |  |
| 9:16.39 | השחין |  |
| 9:16.40 | לה | . |
| 9:16.41 | ותורתינו |  |
| 9:16.42 | הקדושה |  |
| 9:16.43 | העירה |  |
| 9:16.44 | אותנו |  |
| 9:16.45 | באמרה |  |
| 9:16.46 | הקרובים |  |
| 9:16.47 | אליך |  |
| 9:16.48 | או |  |
| 9:16.49 | הרחוקים | , |
| 9:16.50 | שמהקרובים |  |
| 9:16.51 | שהם |  |
| 9:16.52 | פרצופי |  |
| 9:16.53 | השמש |  |
| 9:16.54 | והירח |  |
| 9:16.55 | ושבתי |  |
| 9:16.56 | וצדק |  |
| 9:16.57 | ומאדים |  |
| 9:16.58 | ונגה |  |
| 9:16.59 | וכוכב | , |
| 9:16.60 | שבדו |  |
| 9:16.61 | חוברי |  |
| 9:16.62 | שמים |  |
| 9:16.63 | עובדיהם |  |
| 9:16.64 | כמו |  |
| 9:16.65 | שכתב |  |
| 9:16.66 | הרמב"ם |  |
| 9:16.67 | ורבי |  |
| 9:16.68 | עובדיה |  |
| 9:16.69 | בפ | ' |
| 9:16.70 | כל |  |
| 9:16.71 | הצלמים |  |
| 9:16.72 | שיחסו |  |
| 9:16.73 | לכוכבים |  |
| 9:16.74 | פרצופים | , |
| 9:16.75 | עד |  |
| 9:16.76 | שאמרו |  |
| 9:16.77 | פרצוף |  |
| 9:16.78 | שבתי |  |
| 9:16.79 | זקן |  |
| 9:16.80 | שחור |  |
| 9:16.81 | וישיש | . |
| 9:16.82 | ופרצוף |  |
| 9:16.83 | החמה |  |
| 9:16.84 | פרצוף |  |
| 9:16.85 | מלך |  |
| 9:16.86 | מוכתר |  |
| 9:16.87 | יושב |  |
| 9:16.88 | בעגלה | , |
| 9:16.89 | ופרצוף |  |
| 9:16.90 | נגה |  |
| 9:16.91 | נערה |  |
| 9:16.92 | יפה |  |
| 9:16.93 | לבושה |  |
| 9:16.94 | חלי |  |
| 9:16.95 | זהב | , |
| 9:16.96 | וכן |  |
| 9:16.97 | מיחסים |  |
| 9:16.98 | לכל |  |
| 9:16.99 | הכוכבים |  |
| 9:16.100 | והמזלות |  |
| 9:16.101 | פרצופים |  |
| 9:16.102 | רבים |  |
| 9:16.103 | וכו | '! |
| 9:16.104 | מהפרצופים |  |
| 9:16.105 | של |  |
| 9:16.106 | הכוכבים |  |
| 9:16.107 | הקרובים |  |
| 9:16.108 | אלינו | , |
| 9:16.109 | נקיש |  |
| 9:16.110 | ונלמד |  |
| 9:16.111 | על |  |
| 9:16.112 | הפרצופים |  |
| 9:16.113 | של |  |
| 9:16.114 | ספירות |  |
| 9:16.115 | הרחוקים |  |
| 9:16.116 | ממנו |  |
| 9:16.117 | אשר |  |
| 9:16.118 | בדא |  |
| 9:16.119 | מחבר |  |
| 9:16.120 | הזהר |  |
| 9:16.121 | המסית |  |
| 9:16.122 | אותנו |  |
| 9:16.123 | לעבדם |  |
| 9:16.124 | ולקרותם |  |
| 9:16.125 | בשמותם |  |
| 9:16.126 | של |  |
| 9:16.127 | הקב"ה |  |
| 9:16.128 | שאין |  |
| 9:16.129 | בהם |  |
| 9:16.130 | ממש | ! |
| 9:16.131 | ועל |  |
| 9:16.132 | כן |  |
| 9:16.133 | לא |  |
| 9:16.134 | נאבה |  |
| 9:16.135 | לו |  |
| 9:16.136 | ולא |  |
| 9:16.137 | נשמע |  |
| 9:16.138 | אליו | , |
| 9:16.139 | כמצות |  |
| 9:16.140 | תורתינו |  |
| 9:16.141 | הקדושה |  |
| 9:16.142 | שלא |  |
| 9:16.143 | לשמוע |  |
| 9:16.144 | אל |  |
| 9:16.145 | המסית |  |
| 9:16.146 | שמטעים |  |
| 9:16.147 | את |  |
| 9:16.148 | דבריו |  |
| 9:16.149 | ומשבח |  |
| 9:16.150 | את |  |
| 9:16.151 | העושים |  |
| 9:16.152 | כמעשיו |  |
| 9:16.153 | לשתף |  |
| 9:16.154 | ולצרף |  |
| 9:16.155 | עם |  |
| 9:16.156 | אלהינו |  |
| 9:16.157 | ב"ה |  |
| 9:16.158 | את |  |
| 9:16.159 | הפרצופים |  |
| 9:16.160 | הנזכרים | , |
| 9:16.161 | ולהיות |  |
| 9:16.162 | העבודה |  |
| 9:16.163 | ביחוד |  |
| 9:16.164 | לקצר |  |
| 9:16.165 | האפים |  |
| 9:16.166 | בצירוף |  |
| 9:16.167 | שאר |  |
| 9:16.168 | פרצופים | . |
| 9:16.169 | וליחדם |  |
| 9:16.170 | בפיהם |  |
| 9:16.171 | שהבל |  |
| 9:16.172 | א | ', |
| 9:16.173 | והוסיף |  |
| 9:16.174 | על |  |
| 9:16.175 | השלשה |  |
| 9:16.176 | בשמע |  |
| 9:16.177 | ישראל |  |
| 9:16.178 | גם |  |
| 9:16.179 | את |  |
| 9:16.180 | את |  |
| 9:16.181 | המלכות |  |
| 9:16.182 | שהיא |  |
| 9:16.183 | הרביעית |  |
| 9:16.184 | עם |  |
| 9:16.185 | השלשה |  |
| 9:16.186 | שמות |  |
| 9:16.187 | יהוה |  |
| 9:16.188 | אלהינו |  |
| 9:16.189 | יהוה | . |
| 9:16.190 | כדברי |  |
| 9:16.191 | השלישים | , |
| 9:16.192 | ובמלת |  |
| 9:16.193 | אחד |  |
| 9:16.194 | מצורף |  |
| 9:16.195 | את |  |
| 9:16.196 | המלכוץ |  |
| 9:16.197 | עמהם | . |
| 9:16.198 | ואחריו |  |
| 9:16.199 | נמשכו |  |
| 9:16.200 | המון |  |
| 9:16.201 | המקובלים |  |
| 9:16.202 | החדשים |  |
| 9:16.203 | בתם |  |
| 9:16.204 | לבבם |  |
| 9:16.205 | ולא |  |
| 9:16.206 | נתנו |  |
| 9:16.207 | אל |  |
| 9:16.208 | לבם |  |
| 9:16.209 | להבין |  |
| 9:16.210 | ששטתו |  |
| 9:16.211 | זאת |  |
| 9:16.212 | מילדי |  |
| 9:16.213 | נכרים |  |
| 9:16.214 | המתנגדים |  |
| 9:16.215 | לאמונתינו |  |
| 9:16.216 | ביחוד |  |
| 9:16.217 | השי"ת | , |
| 9:16.218 | והאמינו |  |
| 9:16.219 | לו |  |
| 9:16.220 | בחלק |  |
| 9:16.221 | שפתיו |  |
| 9:16.222 | המורגל |  |
| 9:16.223 | בפיו |  |
| 9:16.224 | תמיד |  |
| 9:16.225 | לומר |  |
| 9:16.226 | זכאה |  |
| 9:16.227 | חולקיה | . |
| 9:16.228 | דכאה |  |
| 9:16.229 | חולקיהון |  |
| 9:16.230 | דישראל | , |
| 9:16.231 | כדי |  |
| 9:16.232 | למשוך |  |
| 9:16.233 | ולגנוב |  |
| 9:16.234 | דעת |  |
| 9:16.235 | השומעים |  |
| 9:16.236 | להדיחם | , |
| 9:16.237 | והבטיח |  |
| 9:16.238 | שלעתיד |  |
| 9:16.239 | קרוב |  |
| 9:16.240 | לימות |  |
| 9:16.241 | המשיח |  |
| 9:16.242 | כלא |  |
| 9:16.243 | יתפרנסון |  |
| 9:16.244 | מחאי |  |
| 9:16.245 | הבורא |  |
| 9:16.246 | דיליה | , |
| 9:16.247 | שיבטלו |  |
| 9:16.248 | כל |  |
| 9:16.249 | הלכות |  |
| 9:16.250 | של |  |
| 9:16.251 | תורה |  |
| 9:16.252 | שבעל |  |
| 9:16.253 | פה | , |
| 9:16.254 | היפך |  |
| 9:16.255 | דברי |  |
| 9:16.256 | רז"ל |  |
| 9:16.257 | שאמרו |  |
| 9:16.258 | שכל |  |
| 9:16.259 | הלכות |  |
| 9:16.260 | של |  |
| 9:16.261 | תורה |  |
| 9:16.262 | של |  |
| 9:16.263 | פה |  |
| 9:16.264 | אינם |  |
| 9:16.265 | בטילים |  |
| 9:16.266 | לעולם |  |
| 9:16.267 | רמב"ם |  |
| 9:16.268 | סוף |  |
| 9:16.269 | הל | ' |
| 9:16.270 | מגלה | ). |
| 9:16.271 | ולא |  |
| 9:16.272 | די |  |
| 9:16.273 | זה |  |
| 9:16.274 | אלא |  |
| 9:16.275 | שהמסית |  |
| 9:16.276 | הנזכר |  |
| 9:16.277 | שם |  |
| 9:16.278 | את |  |
| 9:16.279 | רשב"י |  |
| 9:16.280 | תנא |  |
| 9:16.281 | דידן |  |
| 9:16.282 | לעובר |  |
| 9:16.283 | על |  |
| 9:16.284 | דברי |  |
| 9:16.285 | חכמים |  |
| 9:16.286 | ח"ו |  |
| 9:16.287 | שאמרו |  |
| 9:16.288 | ולא |  |
| 9:16.289 | במרכבה |  |
| 9:16.290 | ואפי | ' |
| 9:16.291 | ביחיד |  |
| 9:16.292 | אלא |  |
| 9:16.293 | א"כ |  |
| 9:16.294 | היה |  |
| 9:16.295 | חכם |  |
| 9:16.296 | ומבין |  |
| 9:16.297 | מדעתו |  |
| 9:16.298 | מוסרין |  |
| 9:16.299 | לו |  |
| 9:16.300 | ראשי |  |
| 9:16.301 | פרקים | . |
| 9:16.302 | ולפי |  |
| 9:16.303 | דבריו |  |
| 9:16.304 | שאלו |  |
| 9:16.305 | הם |  |
| 9:16.306 | מעשה |  |
| 9:16.307 | מרכבה |  |
| 9:16.308 | שהזכירו |  |
| 9:16.309 | רז"ל | , |
| 9:16.310 | איך |  |
| 9:16.311 | היה |  |
| 9:16.312 | רשב"י |  |
| 9:16.313 | דורש |  |
| 9:16.314 | בהם |  |
| 9:16.315 | בפומפי |  |
| 9:16.316 | בין |  |
| 9:16.317 | חביריו |  |
| 9:16.318 | מה |  |
| 9:16.319 | שלא |  |
| 9:16.320 | התירו |  |
| 9:16.321 | חכמים |  |
| 9:16.322 | לעשות |  |
| 9:16.323 | כן |  |
| 9:16.324 | אפי | ' |
| 9:16.325 | היו |  |
| 9:16.326 | כולן |  |
| 9:16.327 | חכמים |  |
| 9:16.328 | גדולים | ? |
| 9:16.329 | ואם |  |
| 9:16.330 | תאמר |  |
| 9:16.331 | שעתיקא |  |
| 9:16.332 | קדושא |  |
| 9:16.333 | נגלה |  |
| 9:16.334 | אליו |  |
| 9:16.335 | כמבואר |  |
| 9:16.336 | בזהר |  |
| 9:16.337 | בראשית | , |
| 9:16.338 | הא |  |
| 9:16.339 | ודאי |  |
| 9:16.340 | לא |  |
| 9:16.341 | צייתנא |  |
| 9:16.342 | ליה |  |
| 9:16.343 | לעקור |  |
| 9:16.344 | דבר |  |
| 9:16.345 | מן |  |
| 9:16.346 | התורה |  |
| 9:16.347 | שקבלו |  |
| 9:16.348 | חכמים | , |
| 9:16.349 | שאין |  |
| 9:16.350 | נביא |  |
| 9:16.351 | רשאי |  |
| 9:16.352 | לחדש |  |
| 9:16.353 | דבר |  |
| 9:16.354 | מעתה | , |
| 9:16.355 | ובודאי |  |
| 9:16.356 | נביא |  |
| 9:16.357 | שקר |  |
| 9:16.358 | הוא | , |
| 9:16.359 | כמו |  |
| 9:16.360 | שמבואר |  |
| 9:16.361 | בתלמוד |  |
| 9:16.362 | בכמה |  |
| 9:16.363 | מקומות |  |
| 9:16.364 | שאם |  |
| 9:16.365 | יבוא |  |
| 9:16.366 | אליהו |  |
| 9:16.367 | ויאמר |  |
| 9:16.368 | וכו | ' |
| 9:16.369 | אין |  |
| 9:16.370 | שומעין |  |
| 9:16.371 | לו | , |
| 9:16.372 | וזה |  |
| 9:16.373 | אפי | ' |
| 9:16.374 | כמידי |  |
| 9:16.375 | דגזרו |  |
| 9:16.376 | ביה |  |
| 9:16.377 | רבנן | , |
| 9:16.378 | כ"ש |  |
| 9:16.379 | וק"ו |  |
| 9:16.380 | לאיסור |  |
| 9:16.381 | עצמי |  |
| 9:16.382 | שקבלו |  |
| 9:16.383 | חכמים |  |
| 9:16.384 | בו |  |
| 9:16.385 | שאסור |  |
| 9:16.386 | לעשות |  |
| 9:16.387 | כן |  |
| 9:16.388 | שאפילו |  |
| 9:16.389 | יבוא |  |
| 9:16.390 | אליהו |  |
| 9:16.391 | ויאמר |  |
| 9:16.392 | שמותר |  |
| 9:16.393 | הוא |  |
| 9:16.394 | שאין |  |
| 9:16.395 | שומעין |  |
| 9:16.396 | לו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

