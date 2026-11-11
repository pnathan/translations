# Translation pack: milhamot-hashem perek-5-20 (Milchamot ha-Shem 5:29–5:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-20`.

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
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
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
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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


## Milchamot ha-Shem 5:29

Hebrew (Wikisource): נה) הן אלה קצות דרכי המקובלים החדשים בענין האלהות, אשר יראה הרואה שהם מתנגדים לגמרי לדברי רבותינו בבראשית רבה וילקוט המובא לעיל סי' ל"א ומדרש רבא פרשת יתרו וירושלמי שבת פרק במה אשה המובא בקונדריס שקדם דמלך בשר ודם שמלך יש לו אב או אח או בן וכו'. והאמת כדברי רז"ל כי אלהינו בהו"ש אחד ואין יחיד כיחודו, עוד כ' הרקח בשורש תשובה וז"ל אלו הן המינים האומר עולם כמנהגו נוהג ואין אלוה בעולם והאומר אין בורא קדמון לכל ( זהו דעת הפילסוף מחבר הזהר דמ"ל שעתיק נגלה לרשב"י והתיר לו לגלות הסוד שהאלוהות עלות וסבות רבות זו למעלה מזו ושאבא הוא הבורא והוא שאמר יהי אור, ויהי רקיע, ויקוו המים, ויהי מארת, אבל נעשה אדם אמרה אימא. ואדם קדמון שהוא למעלה מהם אמר ראו עתה כי אני אני הוא ואין אלהים עמדי) או האומר יש אלוה אחד עמו. (גם בזה שגגו המקובלים ונפתו למחבר הזהר שחושבים לכל אחד מן הפרצופים אלוה הוא ואשתו כדלעיל סי' כ' וכ"א וכ"ד וגם מאמינים במציאות עשר ספירות בקליפה הנקרא בשם אל אחד) או האומר יש לו דמות וצורה כאדם וגוף ותמונה (זה ג"כ דעת המקובלים שאמרו שזהו סוד ויברא אלהים את האדם בצלמו) או משתף דבר אחר עמו (ויתבאר לקטן סי' נ"ח איך המקובלים משתפים, וכבר ביארתי לעיל סי' מ"ח שהמקובלים ממאנים למור ד' בתים הראשונים שביגדל אלהים חי) או האומר שאין שם נבואה בעולם מן השי"ת או מכחיש תורת משה רבינו אפי' חבור אחד או דקדוק אחד או אומר אין הבורא יודע עשת מחשבות האדם, או אין האדם נדון לפי מעשיו (זה אומרים המקובלים על עתיק ואריך ופרצופים שלמעלה מהם. שאינם דנים את האדם לפי מעשיו. ואין אדם יוצא חייב בדיניהם כצדיק כרשע לפי שלא נתכנו להם עלילות, רק לז"א נתכנו עלילות סי' מ"ד ומ"ה). או שאומר אין העולם הבא ואין תשלום שכר, אלו המינים והאפיקורסין שמורידין אותן ולא מעלין אם לא עשו תשובה עכ"ל הרקח. ובספר מצות גדול לאוין א' המצוה הראשונה שלא לעולות במחשבה שיש שם אלוה זולתי השי"ת שנ' לא יהיה לך אלהים אחרים על פני, ותניא במכלתא לפי שנ' לא תעשה לך פסל. העשוי כבר מנין שלא יקיים ת"ל לא יהיה לך. ואין מקרא יוצא מדי פשוטו. ובפרק ד' מיתות אמרו שם כל המשתף שם שמים ודבר אחד נעקר מן העולם שנ' בלתי לה' לבדו. והמקרא הזה בא לעונש. וזה לאזהרה, ואומר ר' יעקב כי מה שנ' חרב לה' ולגדעון, אין זה שתוף לאלהות ולאדנות עכ"ל הסמ"ג.

| id | word | punctuation after |
|---|---|---|
| 5:29.1 | נה | ) |
| 5:29.2 | הן |  |
| 5:29.3 | אלה |  |
| 5:29.4 | קצות |  |
| 5:29.5 | דרכי |  |
| 5:29.6 | המקובלים |  |
| 5:29.7 | החדשים |  |
| 5:29.8 | בענין |  |
| 5:29.9 | האלהות | , |
| 5:29.10 | אשר |  |
| 5:29.11 | יראה |  |
| 5:29.12 | הרואה |  |
| 5:29.13 | שהם |  |
| 5:29.14 | מתנגדים |  |
| 5:29.15 | לגמרי |  |
| 5:29.16 | לדברי |  |
| 5:29.17 | רבותינו |  |
| 5:29.18 | בבראשית |  |
| 5:29.19 | רבה |  |
| 5:29.20 | וילקוט |  |
| 5:29.21 | המובא |  |
| 5:29.22 | לעיל |  |
| 5:29.23 | סי | ' |
| 5:29.24 | ל"א |  |
| 5:29.25 | ומדרש |  |
| 5:29.26 | רבא |  |
| 5:29.27 | פרשת |  |
| 5:29.28 | יתרו |  |
| 5:29.29 | וירושלמי |  |
| 5:29.30 | שבת |  |
| 5:29.31 | פרק |  |
| 5:29.32 | במה |  |
| 5:29.33 | אשה |  |
| 5:29.34 | המובא |  |
| 5:29.35 | בקונדריס |  |
| 5:29.36 | שקדם |  |
| 5:29.37 | דמלך |  |
| 5:29.38 | בשר |  |
| 5:29.39 | ודם |  |
| 5:29.40 | שמלך |  |
| 5:29.41 | יש |  |
| 5:29.42 | לו |  |
| 5:29.43 | אב |  |
| 5:29.44 | או |  |
| 5:29.45 | אח |  |
| 5:29.46 | או |  |
| 5:29.47 | בן |  |
| 5:29.48 | וכו | '. |
| 5:29.49 | והאמת |  |
| 5:29.50 | כדברי |  |
| 5:29.51 | רז"ל |  |
| 5:29.52 | כי |  |
| 5:29.53 | אלהינו |  |
| 5:29.54 | בהו"ש |  |
| 5:29.55 | אחד |  |
| 5:29.56 | ואין |  |
| 5:29.57 | יחיד |  |
| 5:29.58 | כיחודו | , |
| 5:29.59 | עוד |  |
| 5:29.60 | כ | ' |
| 5:29.61 | הרקח |  |
| 5:29.62 | בשורש |  |
| 5:29.63 | תשובה |  |
| 5:29.64 | וז"ל |  |
| 5:29.65 | אלו |  |
| 5:29.66 | הן |  |
| 5:29.67 | המינים |  |
| 5:29.68 | האומר |  |
| 5:29.69 | עולם |  |
| 5:29.70 | כמנהגו |  |
| 5:29.71 | נוהג |  |
| 5:29.72 | ואין |  |
| 5:29.73 | אלוה |  |
| 5:29.74 | בעולם |  |
| 5:29.75 | והאומר |  |
| 5:29.76 | אין |  |
| 5:29.77 | בורא |  |
| 5:29.78 | קדמון |  |
| 5:29.79 | לכל | ( |
| 5:29.80 | זהו |  |
| 5:29.81 | דעת |  |
| 5:29.82 | הפילסוף |  |
| 5:29.83 | מחבר |  |
| 5:29.84 | הזהר |  |
| 5:29.85 | דמ"ל |  |
| 5:29.86 | שעתיק |  |
| 5:29.87 | נגלה |  |
| 5:29.88 | לרשב"י |  |
| 5:29.89 | והתיר |  |
| 5:29.90 | לו |  |
| 5:29.91 | לגלות |  |
| 5:29.92 | הסוד |  |
| 5:29.93 | שהאלוהות |  |
| 5:29.94 | עלות |  |
| 5:29.95 | וסבות |  |
| 5:29.96 | רבות |  |
| 5:29.97 | זו |  |
| 5:29.98 | למעלה |  |
| 5:29.99 | מזו |  |
| 5:29.100 | ושאבא |  |
| 5:29.101 | הוא |  |
| 5:29.102 | הבורא |  |
| 5:29.103 | והוא |  |
| 5:29.104 | שאמר |  |
| 5:29.105 | יהי |  |
| 5:29.106 | אור | , |
| 5:29.107 | ויהי |  |
| 5:29.108 | רקיע | , |
| 5:29.109 | ויקוו |  |
| 5:29.110 | המים | , |
| 5:29.111 | ויהי |  |
| 5:29.112 | מארת | , |
| 5:29.113 | אבל |  |
| 5:29.114 | נעשה |  |
| 5:29.115 | אדם |  |
| 5:29.116 | אמרה |  |
| 5:29.117 | אימא | . |
| 5:29.118 | ואדם |  |
| 5:29.119 | קדמון |  |
| 5:29.120 | שהוא |  |
| 5:29.121 | למעלה |  |
| 5:29.122 | מהם |  |
| 5:29.123 | אמר |  |
| 5:29.124 | ראו |  |
| 5:29.125 | עתה |  |
| 5:29.126 | כי |  |
| 5:29.127 | אני |  |
| 5:29.128 | אני |  |
| 5:29.129 | הוא |  |
| 5:29.130 | ואין |  |
| 5:29.131 | אלהים |  |
| 5:29.132 | עמדי | ) |
| 5:29.133 | או |  |
| 5:29.134 | האומר |  |
| 5:29.135 | יש |  |
| 5:29.136 | אלוה |  |
| 5:29.137 | אחד |  |
| 5:29.138 | עמו | . |
| 5:29.139 | גם |  |
| 5:29.140 | בזה |  |
| 5:29.141 | שגגו |  |
| 5:29.142 | המקובלים |  |
| 5:29.143 | ונפתו |  |
| 5:29.144 | למחבר |  |
| 5:29.145 | הזהר |  |
| 5:29.146 | שחושבים |  |
| 5:29.147 | לכל |  |
| 5:29.148 | אחד |  |
| 5:29.149 | מן |  |
| 5:29.150 | הפרצופים |  |
| 5:29.151 | אלוה |  |
| 5:29.152 | הוא |  |
| 5:29.153 | ואשתו |  |
| 5:29.154 | כדלעיל |  |
| 5:29.155 | סי | ' |
| 5:29.156 | כ | ' |
| 5:29.157 | וכ"א |  |
| 5:29.158 | וכ"ד |  |
| 5:29.159 | וגם |  |
| 5:29.160 | מאמינים |  |
| 5:29.161 | במציאות |  |
| 5:29.162 | עשר |  |
| 5:29.163 | ספירות |  |
| 5:29.164 | בקליפה |  |
| 5:29.165 | הנקרא |  |
| 5:29.166 | בשם |  |
| 5:29.167 | אל |  |
| 5:29.168 | אחד | ) |
| 5:29.169 | או |  |
| 5:29.170 | האומר |  |
| 5:29.171 | יש |  |
| 5:29.172 | לו |  |
| 5:29.173 | דמות |  |
| 5:29.174 | וצורה |  |
| 5:29.175 | כאדם |  |
| 5:29.176 | וגוף |  |
| 5:29.177 | ותמונה |  |
| 5:29.178 | זה |  |
| 5:29.179 | ג"כ |  |
| 5:29.180 | דעת |  |
| 5:29.181 | המקובלים |  |
| 5:29.182 | שאמרו |  |
| 5:29.183 | שזהו |  |
| 5:29.184 | סוד |  |
| 5:29.185 | ויברא |  |
| 5:29.186 | אלהים |  |
| 5:29.187 | את |  |
| 5:29.188 | האדם |  |
| 5:29.189 | בצלמו | ) |
| 5:29.190 | או |  |
| 5:29.191 | משתף |  |
| 5:29.192 | דבר |  |
| 5:29.193 | אחר |  |
| 5:29.194 | עמו |  |
| 5:29.195 | ויתבאר |  |
| 5:29.196 | לקטן |  |
| 5:29.197 | סי | ' |
| 5:29.198 | נ"ח |  |
| 5:29.199 | איך |  |
| 5:29.200 | המקובלים |  |
| 5:29.201 | משתפים | , |
| 5:29.202 | וכבר |  |
| 5:29.203 | ביארתי |  |
| 5:29.204 | לעיל |  |
| 5:29.205 | סי | ' |
| 5:29.206 | מ"ח |  |
| 5:29.207 | שהמקובלים |  |
| 5:29.208 | ממאנים |  |
| 5:29.209 | למור |  |
| 5:29.210 | ד | ' |
| 5:29.211 | בתים |  |
| 5:29.212 | הראשונים |  |
| 5:29.213 | שביגדל |  |
| 5:29.214 | אלהים |  |
| 5:29.215 | חי | ) |
| 5:29.216 | או |  |
| 5:29.217 | האומר |  |
| 5:29.218 | שאין |  |
| 5:29.219 | שם |  |
| 5:29.220 | נבואה |  |
| 5:29.221 | בעולם |  |
| 5:29.222 | מן |  |
| 5:29.223 | השי"ת |  |
| 5:29.224 | או |  |
| 5:29.225 | מכחיש |  |
| 5:29.226 | תורת |  |
| 5:29.227 | משה |  |
| 5:29.228 | רבינו |  |
| 5:29.229 | אפי | ' |
| 5:29.230 | חבור |  |
| 5:29.231 | אחד |  |
| 5:29.232 | או |  |
| 5:29.233 | דקדוק |  |
| 5:29.234 | אחד |  |
| 5:29.235 | או |  |
| 5:29.236 | אומר |  |
| 5:29.237 | אין |  |
| 5:29.238 | הבורא |  |
| 5:29.239 | יודע |  |
| 5:29.240 | עשת |  |
| 5:29.241 | מחשבות |  |
| 5:29.242 | האדם | , |
| 5:29.243 | או |  |
| 5:29.244 | אין |  |
| 5:29.245 | האדם |  |
| 5:29.246 | נדון |  |
| 5:29.247 | לפי |  |
| 5:29.248 | מעשיו |  |
| 5:29.249 | זה |  |
| 5:29.250 | אומרים |  |
| 5:29.251 | המקובלים |  |
| 5:29.252 | על |  |
| 5:29.253 | עתיק |  |
| 5:29.254 | ואריך |  |
| 5:29.255 | ופרצופים |  |
| 5:29.256 | שלמעלה |  |
| 5:29.257 | מהם | . |
| 5:29.258 | שאינם |  |
| 5:29.259 | דנים |  |
| 5:29.260 | את |  |
| 5:29.261 | האדם |  |
| 5:29.262 | לפי |  |
| 5:29.263 | מעשיו | . |
| 5:29.264 | ואין |  |
| 5:29.265 | אדם |  |
| 5:29.266 | יוצא |  |
| 5:29.267 | חייב |  |
| 5:29.268 | בדיניהם |  |
| 5:29.269 | כצדיק |  |
| 5:29.270 | כרשע |  |
| 5:29.271 | לפי |  |
| 5:29.272 | שלא |  |
| 5:29.273 | נתכנו |  |
| 5:29.274 | להם |  |
| 5:29.275 | עלילות | , |
| 5:29.276 | רק |  |
| 5:29.277 | לז"א |  |
| 5:29.278 | נתכנו |  |
| 5:29.279 | עלילות |  |
| 5:29.280 | סי | ' |
| 5:29.281 | מ"ד |  |
| 5:29.282 | ומ"ה | ). |
| 5:29.283 | או |  |
| 5:29.284 | שאומר |  |
| 5:29.285 | אין |  |
| 5:29.286 | העולם |  |
| 5:29.287 | הבא |  |
| 5:29.288 | ואין |  |
| 5:29.289 | תשלום |  |
| 5:29.290 | שכר | , |
| 5:29.291 | אלו |  |
| 5:29.292 | המינים |  |
| 5:29.293 | והאפיקורסין |  |
| 5:29.294 | שמורידין |  |
| 5:29.295 | אותן |  |
| 5:29.296 | ולא |  |
| 5:29.297 | מעלין |  |
| 5:29.298 | אם |  |
| 5:29.299 | לא |  |
| 5:29.300 | עשו |  |
| 5:29.301 | תשובה |  |
| 5:29.302 | עכ"ל |  |
| 5:29.303 | הרקח | . |
| 5:29.304 | ובספר |  |
| 5:29.305 | מצות |  |
| 5:29.306 | גדול |  |
| 5:29.307 | לאוין |  |
| 5:29.308 | א | ' |
| 5:29.309 | המצוה |  |
| 5:29.310 | הראשונה |  |
| 5:29.311 | שלא |  |
| 5:29.312 | לעולות |  |
| 5:29.313 | במחשבה |  |
| 5:29.314 | שיש |  |
| 5:29.315 | שם |  |
| 5:29.316 | אלוה |  |
| 5:29.317 | זולתי |  |
| 5:29.318 | השי"ת |  |
| 5:29.319 | שנ | ' |
| 5:29.320 | לא |  |
| 5:29.321 | יהיה |  |
| 5:29.322 | לך |  |
| 5:29.323 | אלהים |  |
| 5:29.324 | אחרים |  |
| 5:29.325 | על |  |
| 5:29.326 | פני | , |
| 5:29.327 | ותניא |  |
| 5:29.328 | במכלתא |  |
| 5:29.329 | לפי |  |
| 5:29.330 | שנ | ' |
| 5:29.331 | לא |  |
| 5:29.332 | תעשה |  |
| 5:29.333 | לך |  |
| 5:29.334 | פסל | . |
| 5:29.335 | העשוי |  |
| 5:29.336 | כבר |  |
| 5:29.337 | מנין |  |
| 5:29.338 | שלא |  |
| 5:29.339 | יקיים |  |
| 5:29.340 | ת"ל |  |
| 5:29.341 | לא |  |
| 5:29.342 | יהיה |  |
| 5:29.343 | לך | . |
| 5:29.344 | ואין |  |
| 5:29.345 | מקרא |  |
| 5:29.346 | יוצא |  |
| 5:29.347 | מדי |  |
| 5:29.348 | פשוטו | . |
| 5:29.349 | ובפרק |  |
| 5:29.350 | ד | ' |
| 5:29.351 | מיתות |  |
| 5:29.352 | אמרו |  |
| 5:29.353 | שם |  |
| 5:29.354 | כל |  |
| 5:29.355 | המשתף |  |
| 5:29.356 | שם |  |
| 5:29.357 | שמים |  |
| 5:29.358 | ודבר |  |
| 5:29.359 | אחד |  |
| 5:29.360 | נעקר |  |
| 5:29.361 | מן |  |
| 5:29.362 | העולם |  |
| 5:29.363 | שנ | ' |
| 5:29.364 | בלתי |  |
| 5:29.365 | לה | ' |
| 5:29.366 | לבדו | . |
| 5:29.367 | והמקרא |  |
| 5:29.368 | הזה |  |
| 5:29.369 | בא |  |
| 5:29.370 | לעונש | . |
| 5:29.371 | וזה |  |
| 5:29.372 | לאזהרה | , |
| 5:29.373 | ואומר |  |
| 5:29.374 | ר | ' |
| 5:29.375 | יעקב |  |
| 5:29.376 | כי |  |
| 5:29.377 | מה |  |
| 5:29.378 | שנ | ' |
| 5:29.379 | חרב |  |
| 5:29.380 | לה | ' |
| 5:29.381 | ולגדעון | , |
| 5:29.382 | אין |  |
| 5:29.383 | זה |  |
| 5:29.384 | שתוף |  |
| 5:29.385 | לאלהות |  |
| 5:29.386 | ולאדנות |  |
| 5:29.387 | עכ"ל |  |
| 5:29.388 | הסמ"ג | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

