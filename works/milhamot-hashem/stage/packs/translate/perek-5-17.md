# Translation pack: milhamot-hashem perek-5-17 (Milchamot ha-Shem 5:25–5:26)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-17`.

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
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
- ס' before a quoted verse is "verse" (סימן = "section" in the author's
  numbered sections); read it from the context and note the choice.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
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
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
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


## Milchamot ha-Shem 5:25

Hebrew (Wikisource): נא) ובכן נשכבה בבשתינו ותכסינו כלימתינו איך עזבנו אמונתינו הצרפה המקובלת בידינו מנביאינו וחכמינו התלמודיים, ונפתינו למחבר הזהר והתקונים המסית אותנו בחלקלקות להאמין באלוהות רבים ולעבוד אלהים אחרים קצר אפים וזוגתו (זעיר ונוקביה)! והישמעלים עמדו באמונתם ביחוד השי"ת כמ ושקבלו מישראל כמ"ש הרמב"ם בתשובה לגר צדק והרבה מחכמינו שהם מיחדים את השי"ת כמונו וכפי המתבאר מחכמינו התלמודיים, ולא כמו שכתבו המקובלים החדשים שאין אמונתינו וביחוד השי"ת כדעת הישמעאלים והפילסופים האלהיים, אלא בדרך אחרת בשיתוף ברואיו שהיא כאמונת העכו"ם מאמיני השלוש כמו שנתבאר לעיל סימן מ"ב ויתבאר עוד לקמן.

| id | word | punctuation after |
|---|---|---|
| 5:25.1 | נא | ) |
| 5:25.2 | ובכן |  |
| 5:25.3 | נשכבה |  |
| 5:25.4 | בבשתינו |  |
| 5:25.5 | ותכסינו |  |
| 5:25.6 | כלימתינו |  |
| 5:25.7 | איך |  |
| 5:25.8 | עזבנו |  |
| 5:25.9 | אמונתינו |  |
| 5:25.10 | הצרפה |  |
| 5:25.11 | המקובלת |  |
| 5:25.12 | בידינו |  |
| 5:25.13 | מנביאינו |  |
| 5:25.14 | וחכמינו |  |
| 5:25.15 | התלמודיים | , |
| 5:25.16 | ונפתינו |  |
| 5:25.17 | למחבר |  |
| 5:25.18 | הזהר |  |
| 5:25.19 | והתקונים |  |
| 5:25.20 | המסית |  |
| 5:25.21 | אותנו |  |
| 5:25.22 | בחלקלקות |  |
| 5:25.23 | להאמין |  |
| 5:25.24 | באלוהות |  |
| 5:25.25 | רבים |  |
| 5:25.26 | ולעבוד |  |
| 5:25.27 | אלהים |  |
| 5:25.28 | אחרים |  |
| 5:25.29 | קצר |  |
| 5:25.30 | אפים |  |
| 5:25.31 | וזוגתו |  |
| 5:25.32 | זעיר |  |
| 5:25.33 | ונוקביה | )! |
| 5:25.34 | והישמעלים |  |
| 5:25.35 | עמדו |  |
| 5:25.36 | באמונתם |  |
| 5:25.37 | ביחוד |  |
| 5:25.38 | השי"ת |  |
| 5:25.39 | כמ |  |
| 5:25.40 | ושקבלו |  |
| 5:25.41 | מישראל |  |
| 5:25.42 | כמ"ש |  |
| 5:25.43 | הרמב"ם |  |
| 5:25.44 | בתשובה |  |
| 5:25.45 | לגר |  |
| 5:25.46 | צדק |  |
| 5:25.47 | והרבה |  |
| 5:25.48 | מחכמינו |  |
| 5:25.49 | שהם |  |
| 5:25.50 | מיחדים |  |
| 5:25.51 | את |  |
| 5:25.52 | השי"ת |  |
| 5:25.53 | כמונו |  |
| 5:25.54 | וכפי |  |
| 5:25.55 | המתבאר |  |
| 5:25.56 | מחכמינו |  |
| 5:25.57 | התלמודיים | , |
| 5:25.58 | ולא |  |
| 5:25.59 | כמו |  |
| 5:25.60 | שכתבו |  |
| 5:25.61 | המקובלים |  |
| 5:25.62 | החדשים |  |
| 5:25.63 | שאין |  |
| 5:25.64 | אמונתינו |  |
| 5:25.65 | וביחוד |  |
| 5:25.66 | השי"ת |  |
| 5:25.67 | כדעת |  |
| 5:25.68 | הישמעאלים |  |
| 5:25.69 | והפילסופים |  |
| 5:25.70 | האלהיים | , |
| 5:25.71 | אלא |  |
| 5:25.72 | בדרך |  |
| 5:25.73 | אחרת |  |
| 5:25.74 | בשיתוף |  |
| 5:25.75 | ברואיו |  |
| 5:25.76 | שהיא |  |
| 5:25.77 | כאמונת |  |
| 5:25.78 | העכו"ם |  |
| 5:25.79 | מאמיני |  |
| 5:25.80 | השלוש |  |
| 5:25.81 | כמו |  |
| 5:25.82 | שנתבאר |  |
| 5:25.83 | לעיל |  |
| 5:25.84 | סימן |  |
| 5:25.85 | מ"ב |  |
| 5:25.86 | ויתבאר |  |
| 5:25.87 | עוד |  |
| 5:25.88 | לקמן | . |


## Milchamot ha-Shem 5:26

Hebrew (Wikisource): נב) ובעזרת המלמד לאדם דעת נעתיק דברי קצת מרז"ל שוני הלכות אשר מפיהם אנו חיים והולכים לאור תורתם בדיני התורה וחוקותיה, וכבר העתקתי בקונדרים הקודם קצת דברי הרמב"ם ז"ל, ולעת עתה נעתיק דברי הר"ר אליעזר בעל הרוקח ז"ל בשורש זכירת השם וז"ל כשיזכור השם באמרו ברוך אתה ה', אל יחשוב על הכבוד הנראה בלב הנביאים ומראה על הכסא, כי אם על ה' הנמצא בשמים ובארץ ובאויר ובים ובכל העולם, אלהי האבות העומד לנגדו שויתי ה' לנגדי תמיד ע"כ. ובשורש יחוד ה' כתב לך דומיה תחלה, יעירו חסידיך, אין לדמותך למעבדיך, מה לערוך נגדיך, מי ראה לידע סודך, מי בא עדיך, הכל תראה נגדיך, וכו' כל חכם לב יש לו לדעת בשביל מדעו, לכל הפחות קצת יחודו, אמת יחוד השם, וסוד הבורא אין כח לדרוש מה הוא. כי לא יראהו כל חי, לא מלאך ולא שרף ולא נביא, כי הכל נברא, והוא ברא כל, הוא היוצר, הוא הבורא, הוא אחד ואין שני לו, ומי ראהו, על כן אין לדמותו לכל דמיון ולכל מראה כי כל דבר שיש לו מראה יש לו גבול וקץ, אבל הבורא אין לו קץ לא למעלה ולא למטה, ולא לארבע רוחות העולם. לא ראש ולא סוף, לא לחכמתו ולא ליכלתו, אך כח מעשיו הגיד לעמו, להודיע גבורותיו וכבוד הדר מלכותו, הוא בעת אחד וברגע אחד יקראו אליו יתפלו כל יצוריו, כל אחד נגע לבבו וצרותיו, ועונה לכל אחד ואחד, ידוע תדע, כי הוא קרוב לכל קוראיו וישמע, על כן החכם עיניו בראשו בני שמע לדברי ואל תיגע לחשוב מה הוא, כי לא תוכל לידע דבר מהבורא, כי הוא לא נברא, היה, הוה, ויהיה, מה דמות תערכו לו, אל מי תדמיוני ואשוה, כי כל דמות הנוצרים אין ביוצר, וכל דמיון הברואים אין בבורא, וכל תמונת העשוים אין בעושה. בספר יצירה עשר ספירות (בלי מה) שאין להם סוף, עומק ראשית, עומק אחרית, עומק טוב, עומק רע, עומק רום, עומק תחת, עומק מזרח, עומק מערב, עומק צפון, עומק דרום, ואדון יחיד אל מלך נאמן מושל בכולן וכו'. הזאת ידעת מני עד, אין לך לדרוש אלא מני שים אדם עלי ארץ, כדאיתא במדרש רבה, בגדול ממך אל תדרוש, במה שהורשת התבונן. ואין לך חקר בנסתרות, כי אם למן היום אשר ברא אלהים אדם על הארץ, אמר ר' לוי למה נברא העולם בבית, מה בית סתום מכל צד ופתוח מצד אחר, כל אין לדרוש מה למעלה מה למטה מה לפנים ומה לאחור. אמר ר' תנחומא כי גדול אתה ועושה נפלאות אתה אלהים לבדך בראת את העולם, בראשית ברא אלהים לא בעמל ולא ביגיעה. אלא בדברי ה' שמים נעשו, הבן בחכמה כי הוא מלא כל, ונעלם מעין כל, ואינו נראה, ואין לו סוף, וחכם מאליו, ואין דבר נעלם ממנו בקרבך, הלא את השמים ואת הארץ אני מלא, כדאמרינן בחגיגה פרק אין דורשין. ומלך רם ונשא שוכן עליהם, ובריש מדרש משלי ומושב יקרו בשמים ממעל ושכינת עזו בגבהי מרומים, הוא אלוהינו ואין עוד אחר, אם אסק שמים שם אתה ואציעה שאול הנך, גם שם ידך תניחני ותקחני, בבראשית רבה אמר ר' אוניא בר סיסיא פעמים שאין העולם ומלואו מחזיק אלהותו, ופעמים שהוא מדבר בין שערות ראשו, מועט מחזיק את המרובה וכו' בתורת כהנים. את השמים ואת הארץ אני מלא ראה חיבתן של ישראל. כבוד המרובה בין שני הכרובים, בבראשית רבה וירא אלהים את כל אשר עשה, התחתונים העליונים בסקירה אחת, העולם הזה והעולם הבא, והוא באחד ומיוחד, ברא את העולם בולא כלום, ברצון חפצו ונפשו אותה ויעש, ולא שלא כרצונו, בדבר ה' שמים ונעשו, הבורא אינו צריך למקום ומכון, כי היה קדם כל הויות, ואין קירות והקורות חוצצין לפניו. כי לא היה בורא דבר שהוא מזיק כנגדו. יתברך הבורא והיוצר כל, דק מכל דק, נפלא ונעלם וחבוי וספון. כל דבר בריאה אין שוה לבורא. אין שייך בו שיעור וקצב, ולא אורך ולא רוחב ולא דמות ולא צורת גולם, ולא דיבוק ולא איברים, ולא צל ולא אור, ולא דמות ולא צורה, ולא סיבוב. וכפי רצונו משמיע קול לפני כבודו המבוררה, המדמה עצמו כפי צורך השעה. הבורא הוא ההוה בכל, הבורא אין לו דמות וצורה, ולא מקרה ולא גוף, ולא שיתוף וסמך ודיבוק, ואינו צריך לכלום, ולא למעון ולא למקום, ואינו דבוק למקום, וממלא כל השמים והארץ, והים והאויר וכל העולמים, ואין שייך בו לא עמידה ולא ישיבה ולא הליכה ולא סמיכה, והוא חי בכחו ויכלתו כל הימים, ואין לו נדנוד ולא יגיעה ולא פעולת כלי הוא בורא כל, בלא עמל ולא יגיעה, כל עיני הנביאים וכל עין לא תשורנו ולא תשופנו. הבורא היה טרם כל, מכון וקדם לכל, ראשון לכל חי וקיים גדול נבור ונורא, מלך מהולל ומשובח, ואין בו לא תוספת ולא חסרון, ולא מיעוט ולא יתרון, אין קץ להוייתו. החקר אלוה תמצא, אם עד תכלית שדי המצא. לתבונתו אין מספר, לגדולתו אין חקר. הבורא רואה את הכל. עין רואה ואוזן שומעת. על כן הצנע לכת בבית הכסא ועם אשתך. כשתחשוב על בורא עולם מה הוא, והיאך חנייתו בכל מקום ומעשיהו. בלום פיך מלדבר, ולבר מלהרהר. הסר המחשבה מלבך. ואם סר לבך למחשבה זאת, חושה ומהר, ואל תהרהר, ושוב ליחוד מקומו שלעולם, לעבודתו וליראתו, בחיות הקדש כשרוצין לזקוף קומתן מיד כופפין מאימתו, ושבים ומשתחוים. כך סבב לבך מלהרהר, ועל דבר זה נכרתה ברית שלא לחשוב באלהותו. שאין כל חכמים יכולים לידע. ושוב וחשוב כי הוא יוצר הכל, והוא הבורא, האמן באלהותו כי הוא אחד, יחיד הבורא. ומראה הכבוד כפי רצונו וחפצו. מראה כבוד ה' כאש אוכלת, וקראו השכינה. ופעמים נראה בלא צורה. כי אם אורה בלא ראות לכל יצירה. ושומע הדבור ורואה החזון מראה מופלאה. ויאמר ראיתי את ה', את לרכות כבודו וכו' ומראה דמות כבודו כחפץ גזרותיו. פעמים דמות אדם, ופעמים דמות אחר כי חפצו מראה כבודו, כצבאות חיילות דבוקים יחד, ה' צבאות, ופעמים גופים מחולקים, ה' אלהי הצבאות, ופעמים נקראים בני האלהים, ופעמים עירין וקדישין, והדבור בא כמפי הכבוד הגדול, הזהר פן תכזב עליו. השמר לך ושמור נפשך מאד פן תחטיא את בשרך לומר עליו חלוקת איברים, מה דמות תערכו לו, כי לא ראיתם כל תמונה, ומה שכתוב ויראו את אלהי ישראל, הכבוד ראו (כמו שתרגם אונקלוס) ובפרק קמא דברכות מה הקב"ה רואה ואינו נראה, אף הנשמה רואה ואינה נראית, מנין שהקב"ה מניח תפילין, פי' מראה מכבודו באובנתא דליבא, כמו וירא יעקב כי יש שבר במצרים, ולבי ראה הרבה, כי לא יראני האדם וחי, ר' עקיבה אומר כמשמעו, וחי אלו מלאכי השרת וחיות הקדש, ר' אליעזר אומר וחי אינו רואה אבל רואה בשעת מיתה, ובספרי וחי בחייהן אינן רואין, אבל רואין בשעת מיתה, ר' עקיבה אומר וחי, חיות הנושאות את הכסא אינן רואות את הכבוד, אמר רשב"י אף מלאכי השרת שחייהן חיי עולם אינן רואים את הכבוד. והאריך עוד הרוקח ז"ל ומסיים שם, הבורא הזהירנו ליחדו שני פעמים ביום ה' אלהינו ה' אחד, בקבלת הלב האמנתו ולא בשמיעת האזן, בכאן נקבלהו באלוה, זה אלהינו, ולהאמין שהוא אחד, ויחוד שמו, וידעת היום והשבות אל לבביך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד, הוא קדמון וחי לעולמים, ואין לו שותף, והוא אחד ואין לדמותו, אל דמות וצורה, ונשמרתם מאד לנשפותיכם כי לא ראיתם כל תמונה, ומה דמות תערכו לו. ומשמותיו ומשבחיו ומעשיו לא נוכל לדעת. כי אין כמוהו. ברוך הוא וברוך שמו ע"כ לשון הרקח.

| id | word | punctuation after |
|---|---|---|
| 5:26.1 | נב | ) |
| 5:26.2 | ובעזרת |  |
| 5:26.3 | המלמד |  |
| 5:26.4 | לאדם |  |
| 5:26.5 | דעת |  |
| 5:26.6 | נעתיק |  |
| 5:26.7 | דברי |  |
| 5:26.8 | קצת |  |
| 5:26.9 | מרז"ל |  |
| 5:26.10 | שוני |  |
| 5:26.11 | הלכות |  |
| 5:26.12 | אשר |  |
| 5:26.13 | מפיהם |  |
| 5:26.14 | אנו |  |
| 5:26.15 | חיים |  |
| 5:26.16 | והולכים |  |
| 5:26.17 | לאור |  |
| 5:26.18 | תורתם |  |
| 5:26.19 | בדיני |  |
| 5:26.20 | התורה |  |
| 5:26.21 | וחוקותיה | , |
| 5:26.22 | וכבר |  |
| 5:26.23 | העתקתי |  |
| 5:26.24 | בקונדרים |  |
| 5:26.25 | הקודם |  |
| 5:26.26 | קצת |  |
| 5:26.27 | דברי |  |
| 5:26.28 | הרמב"ם |  |
| 5:26.29 | ז"ל | , |
| 5:26.30 | ולעת |  |
| 5:26.31 | עתה |  |
| 5:26.32 | נעתיק |  |
| 5:26.33 | דברי |  |
| 5:26.34 | הר"ר |  |
| 5:26.35 | אליעזר |  |
| 5:26.36 | בעל |  |
| 5:26.37 | הרוקח |  |
| 5:26.38 | ז"ל |  |
| 5:26.39 | בשורש |  |
| 5:26.40 | זכירת |  |
| 5:26.41 | השם |  |
| 5:26.42 | וז"ל |  |
| 5:26.43 | כשיזכור |  |
| 5:26.44 | השם |  |
| 5:26.45 | באמרו |  |
| 5:26.46 | ברוך |  |
| 5:26.47 | אתה |  |
| 5:26.48 | ה | ', |
| 5:26.49 | אל |  |
| 5:26.50 | יחשוב |  |
| 5:26.51 | על |  |
| 5:26.52 | הכבוד |  |
| 5:26.53 | הנראה |  |
| 5:26.54 | בלב |  |
| 5:26.55 | הנביאים |  |
| 5:26.56 | ומראה |  |
| 5:26.57 | על |  |
| 5:26.58 | הכסא | , |
| 5:26.59 | כי |  |
| 5:26.60 | אם |  |
| 5:26.61 | על |  |
| 5:26.62 | ה | ' |
| 5:26.63 | הנמצא |  |
| 5:26.64 | בשמים |  |
| 5:26.65 | ובארץ |  |
| 5:26.66 | ובאויר |  |
| 5:26.67 | ובים |  |
| 5:26.68 | ובכל |  |
| 5:26.69 | העולם | , |
| 5:26.70 | אלהי |  |
| 5:26.71 | האבות |  |
| 5:26.72 | העומד |  |
| 5:26.73 | לנגדו |  |
| 5:26.74 | שויתי |  |
| 5:26.75 | ה | ' |
| 5:26.76 | לנגדי |  |
| 5:26.77 | תמיד |  |
| 5:26.78 | ע"כ | . |
| 5:26.79 | ובשורש |  |
| 5:26.80 | יחוד |  |
| 5:26.81 | ה | ' |
| 5:26.82 | כתב |  |
| 5:26.83 | לך |  |
| 5:26.84 | דומיה |  |
| 5:26.85 | תחלה | , |
| 5:26.86 | יעירו |  |
| 5:26.87 | חסידיך | , |
| 5:26.88 | אין |  |
| 5:26.89 | לדמותך |  |
| 5:26.90 | למעבדיך | , |
| 5:26.91 | מה |  |
| 5:26.92 | לערוך |  |
| 5:26.93 | נגדיך | , |
| 5:26.94 | מי |  |
| 5:26.95 | ראה |  |
| 5:26.96 | לידע |  |
| 5:26.97 | סודך | , |
| 5:26.98 | מי |  |
| 5:26.99 | בא |  |
| 5:26.100 | עדיך | , |
| 5:26.101 | הכל |  |
| 5:26.102 | תראה |  |
| 5:26.103 | נגדיך | , |
| 5:26.104 | וכו | ' |
| 5:26.105 | כל |  |
| 5:26.106 | חכם |  |
| 5:26.107 | לב |  |
| 5:26.108 | יש |  |
| 5:26.109 | לו |  |
| 5:26.110 | לדעת |  |
| 5:26.111 | בשביל |  |
| 5:26.112 | מדעו | , |
| 5:26.113 | לכל |  |
| 5:26.114 | הפחות |  |
| 5:26.115 | קצת |  |
| 5:26.116 | יחודו | , |
| 5:26.117 | אמת |  |
| 5:26.118 | יחוד |  |
| 5:26.119 | השם | , |
| 5:26.120 | וסוד |  |
| 5:26.121 | הבורא |  |
| 5:26.122 | אין |  |
| 5:26.123 | כח |  |
| 5:26.124 | לדרוש |  |
| 5:26.125 | מה |  |
| 5:26.126 | הוא | . |
| 5:26.127 | כי |  |
| 5:26.128 | לא |  |
| 5:26.129 | יראהו |  |
| 5:26.130 | כל |  |
| 5:26.131 | חי | , |
| 5:26.132 | לא |  |
| 5:26.133 | מלאך |  |
| 5:26.134 | ולא |  |
| 5:26.135 | שרף |  |
| 5:26.136 | ולא |  |
| 5:26.137 | נביא | , |
| 5:26.138 | כי |  |
| 5:26.139 | הכל |  |
| 5:26.140 | נברא | , |
| 5:26.141 | והוא |  |
| 5:26.142 | ברא |  |
| 5:26.143 | כל | , |
| 5:26.144 | הוא |  |
| 5:26.145 | היוצר | , |
| 5:26.146 | הוא |  |
| 5:26.147 | הבורא | , |
| 5:26.148 | הוא |  |
| 5:26.149 | אחד |  |
| 5:26.150 | ואין |  |
| 5:26.151 | שני |  |
| 5:26.152 | לו | , |
| 5:26.153 | ומי |  |
| 5:26.154 | ראהו | , |
| 5:26.155 | על |  |
| 5:26.156 | כן |  |
| 5:26.157 | אין |  |
| 5:26.158 | לדמותו |  |
| 5:26.159 | לכל |  |
| 5:26.160 | דמיון |  |
| 5:26.161 | ולכל |  |
| 5:26.162 | מראה |  |
| 5:26.163 | כי |  |
| 5:26.164 | כל |  |
| 5:26.165 | דבר |  |
| 5:26.166 | שיש |  |
| 5:26.167 | לו |  |
| 5:26.168 | מראה |  |
| 5:26.169 | יש |  |
| 5:26.170 | לו |  |
| 5:26.171 | גבול |  |
| 5:26.172 | וקץ | , |
| 5:26.173 | אבל |  |
| 5:26.174 | הבורא |  |
| 5:26.175 | אין |  |
| 5:26.176 | לו |  |
| 5:26.177 | קץ |  |
| 5:26.178 | לא |  |
| 5:26.179 | למעלה |  |
| 5:26.180 | ולא |  |
| 5:26.181 | למטה | , |
| 5:26.182 | ולא |  |
| 5:26.183 | לארבע |  |
| 5:26.184 | רוחות |  |
| 5:26.185 | העולם | . |
| 5:26.186 | לא |  |
| 5:26.187 | ראש |  |
| 5:26.188 | ולא |  |
| 5:26.189 | סוף | , |
| 5:26.190 | לא |  |
| 5:26.191 | לחכמתו |  |
| 5:26.192 | ולא |  |
| 5:26.193 | ליכלתו | , |
| 5:26.194 | אך |  |
| 5:26.195 | כח |  |
| 5:26.196 | מעשיו |  |
| 5:26.197 | הגיד |  |
| 5:26.198 | לעמו | , |
| 5:26.199 | להודיע |  |
| 5:26.200 | גבורותיו |  |
| 5:26.201 | וכבוד |  |
| 5:26.202 | הדר |  |
| 5:26.203 | מלכותו | , |
| 5:26.204 | הוא |  |
| 5:26.205 | בעת |  |
| 5:26.206 | אחד |  |
| 5:26.207 | וברגע |  |
| 5:26.208 | אחד |  |
| 5:26.209 | יקראו |  |
| 5:26.210 | אליו |  |
| 5:26.211 | יתפלו |  |
| 5:26.212 | כל |  |
| 5:26.213 | יצוריו | , |
| 5:26.214 | כל |  |
| 5:26.215 | אחד |  |
| 5:26.216 | נגע |  |
| 5:26.217 | לבבו |  |
| 5:26.218 | וצרותיו | , |
| 5:26.219 | ועונה |  |
| 5:26.220 | לכל |  |
| 5:26.221 | אחד |  |
| 5:26.222 | ואחד | , |
| 5:26.223 | ידוע |  |
| 5:26.224 | תדע | , |
| 5:26.225 | כי |  |
| 5:26.226 | הוא |  |
| 5:26.227 | קרוב |  |
| 5:26.228 | לכל |  |
| 5:26.229 | קוראיו |  |
| 5:26.230 | וישמע | , |
| 5:26.231 | על |  |
| 5:26.232 | כן |  |
| 5:26.233 | החכם |  |
| 5:26.234 | עיניו |  |
| 5:26.235 | בראשו |  |
| 5:26.236 | בני |  |
| 5:26.237 | שמע |  |
| 5:26.238 | לדברי |  |
| 5:26.239 | ואל |  |
| 5:26.240 | תיגע |  |
| 5:26.241 | לחשוב |  |
| 5:26.242 | מה |  |
| 5:26.243 | הוא | , |
| 5:26.244 | כי |  |
| 5:26.245 | לא |  |
| 5:26.246 | תוכל |  |
| 5:26.247 | לידע |  |
| 5:26.248 | דבר |  |
| 5:26.249 | מהבורא | , |
| 5:26.250 | כי |  |
| 5:26.251 | הוא |  |
| 5:26.252 | לא |  |
| 5:26.253 | נברא | , |
| 5:26.254 | היה | , |
| 5:26.255 | הוה | , |
| 5:26.256 | ויהיה | , |
| 5:26.257 | מה |  |
| 5:26.258 | דמות |  |
| 5:26.259 | תערכו |  |
| 5:26.260 | לו | , |
| 5:26.261 | אל |  |
| 5:26.262 | מי |  |
| 5:26.263 | תדמיוני |  |
| 5:26.264 | ואשוה | , |
| 5:26.265 | כי |  |
| 5:26.266 | כל |  |
| 5:26.267 | דמות |  |
| 5:26.268 | הנוצרים |  |
| 5:26.269 | אין |  |
| 5:26.270 | ביוצר | , |
| 5:26.271 | וכל |  |
| 5:26.272 | דמיון |  |
| 5:26.273 | הברואים |  |
| 5:26.274 | אין |  |
| 5:26.275 | בבורא | , |
| 5:26.276 | וכל |  |
| 5:26.277 | תמונת |  |
| 5:26.278 | העשוים |  |
| 5:26.279 | אין |  |
| 5:26.280 | בעושה | . |
| 5:26.281 | בספר |  |
| 5:26.282 | יצירה |  |
| 5:26.283 | עשר |  |
| 5:26.284 | ספירות |  |
| 5:26.285 | בלי |  |
| 5:26.286 | מה | ) |
| 5:26.287 | שאין |  |
| 5:26.288 | להם |  |
| 5:26.289 | סוף | , |
| 5:26.290 | עומק |  |
| 5:26.291 | ראשית | , |
| 5:26.292 | עומק |  |
| 5:26.293 | אחרית | , |
| 5:26.294 | עומק |  |
| 5:26.295 | טוב | , |
| 5:26.296 | עומק |  |
| 5:26.297 | רע | , |
| 5:26.298 | עומק |  |
| 5:26.299 | רום | , |
| 5:26.300 | עומק |  |
| 5:26.301 | תחת | , |
| 5:26.302 | עומק |  |
| 5:26.303 | מזרח | , |
| 5:26.304 | עומק |  |
| 5:26.305 | מערב | , |
| 5:26.306 | עומק |  |
| 5:26.307 | צפון | , |
| 5:26.308 | עומק |  |
| 5:26.309 | דרום | , |
| 5:26.310 | ואדון |  |
| 5:26.311 | יחיד |  |
| 5:26.312 | אל |  |
| 5:26.313 | מלך |  |
| 5:26.314 | נאמן |  |
| 5:26.315 | מושל |  |
| 5:26.316 | בכולן |  |
| 5:26.317 | וכו | '. |
| 5:26.318 | הזאת |  |
| 5:26.319 | ידעת |  |
| 5:26.320 | מני |  |
| 5:26.321 | עד | , |
| 5:26.322 | אין |  |
| 5:26.323 | לך |  |
| 5:26.324 | לדרוש |  |
| 5:26.325 | אלא |  |
| 5:26.326 | מני |  |
| 5:26.327 | שים |  |
| 5:26.328 | אדם |  |
| 5:26.329 | עלי |  |
| 5:26.330 | ארץ | , |
| 5:26.331 | כדאיתא |  |
| 5:26.332 | במדרש |  |
| 5:26.333 | רבה | , |
| 5:26.334 | בגדול |  |
| 5:26.335 | ממך |  |
| 5:26.336 | אל |  |
| 5:26.337 | תדרוש | , |
| 5:26.338 | במה |  |
| 5:26.339 | שהורשת |  |
| 5:26.340 | התבונן | . |
| 5:26.341 | ואין |  |
| 5:26.342 | לך |  |
| 5:26.343 | חקר |  |
| 5:26.344 | בנסתרות | , |
| 5:26.345 | כי |  |
| 5:26.346 | אם |  |
| 5:26.347 | למן |  |
| 5:26.348 | היום |  |
| 5:26.349 | אשר |  |
| 5:26.350 | ברא |  |
| 5:26.351 | אלהים |  |
| 5:26.352 | אדם |  |
| 5:26.353 | על |  |
| 5:26.354 | הארץ | , |
| 5:26.355 | אמר |  |
| 5:26.356 | ר | ' |
| 5:26.357 | לוי |  |
| 5:26.358 | למה |  |
| 5:26.359 | נברא |  |
| 5:26.360 | העולם |  |
| 5:26.361 | בבית | , |
| 5:26.362 | מה |  |
| 5:26.363 | בית |  |
| 5:26.364 | סתום |  |
| 5:26.365 | מכל |  |
| 5:26.366 | צד |  |
| 5:26.367 | ופתוח |  |
| 5:26.368 | מצד |  |
| 5:26.369 | אחר | , |
| 5:26.370 | כל |  |
| 5:26.371 | אין |  |
| 5:26.372 | לדרוש |  |
| 5:26.373 | מה |  |
| 5:26.374 | למעלה |  |
| 5:26.375 | מה |  |
| 5:26.376 | למטה |  |
| 5:26.377 | מה |  |
| 5:26.378 | לפנים |  |
| 5:26.379 | ומה |  |
| 5:26.380 | לאחור | . |
| 5:26.381 | אמר |  |
| 5:26.382 | ר | ' |
| 5:26.383 | תנחומא |  |
| 5:26.384 | כי |  |
| 5:26.385 | גדול |  |
| 5:26.386 | אתה |  |
| 5:26.387 | ועושה |  |
| 5:26.388 | נפלאות |  |
| 5:26.389 | אתה |  |
| 5:26.390 | אלהים |  |
| 5:26.391 | לבדך |  |
| 5:26.392 | בראת |  |
| 5:26.393 | את |  |
| 5:26.394 | העולם | , |
| 5:26.395 | בראשית |  |
| 5:26.396 | ברא |  |
| 5:26.397 | אלהים |  |
| 5:26.398 | לא |  |
| 5:26.399 | בעמל |  |
| 5:26.400 | ולא |  |
| 5:26.401 | ביגיעה | . |
| 5:26.402 | אלא |  |
| 5:26.403 | בדברי |  |
| 5:26.404 | ה | ' |
| 5:26.405 | שמים |  |
| 5:26.406 | נעשו | , |
| 5:26.407 | הבן |  |
| 5:26.408 | בחכמה |  |
| 5:26.409 | כי |  |
| 5:26.410 | הוא |  |
| 5:26.411 | מלא |  |
| 5:26.412 | כל | , |
| 5:26.413 | ונעלם |  |
| 5:26.414 | מעין |  |
| 5:26.415 | כל | , |
| 5:26.416 | ואינו |  |
| 5:26.417 | נראה | , |
| 5:26.418 | ואין |  |
| 5:26.419 | לו |  |
| 5:26.420 | סוף | , |
| 5:26.421 | וחכם |  |
| 5:26.422 | מאליו | , |
| 5:26.423 | ואין |  |
| 5:26.424 | דבר |  |
| 5:26.425 | נעלם |  |
| 5:26.426 | ממנו |  |
| 5:26.427 | בקרבך | , |
| 5:26.428 | הלא |  |
| 5:26.429 | את |  |
| 5:26.430 | השמים |  |
| 5:26.431 | ואת |  |
| 5:26.432 | הארץ |  |
| 5:26.433 | אני |  |
| 5:26.434 | מלא | , |
| 5:26.435 | כדאמרינן |  |
| 5:26.436 | בחגיגה |  |
| 5:26.437 | פרק |  |
| 5:26.438 | אין |  |
| 5:26.439 | דורשין | . |
| 5:26.440 | ומלך |  |
| 5:26.441 | רם |  |
| 5:26.442 | ונשא |  |
| 5:26.443 | שוכן |  |
| 5:26.444 | עליהם | , |
| 5:26.445 | ובריש |  |
| 5:26.446 | מדרש |  |
| 5:26.447 | משלי |  |
| 5:26.448 | ומושב |  |
| 5:26.449 | יקרו |  |
| 5:26.450 | בשמים |  |
| 5:26.451 | ממעל |  |
| 5:26.452 | ושכינת |  |
| 5:26.453 | עזו |  |
| 5:26.454 | בגבהי |  |
| 5:26.455 | מרומים | , |
| 5:26.456 | הוא |  |
| 5:26.457 | אלוהינו |  |
| 5:26.458 | ואין |  |
| 5:26.459 | עוד |  |
| 5:26.460 | אחר | , |
| 5:26.461 | אם |  |
| 5:26.462 | אסק |  |
| 5:26.463 | שמים |  |
| 5:26.464 | שם |  |
| 5:26.465 | אתה |  |
| 5:26.466 | ואציעה |  |
| 5:26.467 | שאול |  |
| 5:26.468 | הנך | , |
| 5:26.469 | גם |  |
| 5:26.470 | שם |  |
| 5:26.471 | ידך |  |
| 5:26.472 | תניחני |  |
| 5:26.473 | ותקחני | , |
| 5:26.474 | בבראשית |  |
| 5:26.475 | רבה |  |
| 5:26.476 | אמר |  |
| 5:26.477 | ר | ' |
| 5:26.478 | אוניא |  |
| 5:26.479 | בר |  |
| 5:26.480 | סיסיא |  |
| 5:26.481 | פעמים |  |
| 5:26.482 | שאין |  |
| 5:26.483 | העולם |  |
| 5:26.484 | ומלואו |  |
| 5:26.485 | מחזיק |  |
| 5:26.486 | אלהותו | , |
| 5:26.487 | ופעמים |  |
| 5:26.488 | שהוא |  |
| 5:26.489 | מדבר |  |
| 5:26.490 | בין |  |
| 5:26.491 | שערות |  |
| 5:26.492 | ראשו | , |
| 5:26.493 | מועט |  |
| 5:26.494 | מחזיק |  |
| 5:26.495 | את |  |
| 5:26.496 | המרובה |  |
| 5:26.497 | וכו | ' |
| 5:26.498 | בתורת |  |
| 5:26.499 | כהנים | . |
| 5:26.500 | את |  |
| 5:26.501 | השמים |  |
| 5:26.502 | ואת |  |
| 5:26.503 | הארץ |  |
| 5:26.504 | אני |  |
| 5:26.505 | מלא |  |
| 5:26.506 | ראה |  |
| 5:26.507 | חיבתן |  |
| 5:26.508 | של |  |
| 5:26.509 | ישראל | . |
| 5:26.510 | כבוד |  |
| 5:26.511 | המרובה |  |
| 5:26.512 | בין |  |
| 5:26.513 | שני |  |
| 5:26.514 | הכרובים | , |
| 5:26.515 | בבראשית |  |
| 5:26.516 | רבה |  |
| 5:26.517 | וירא |  |
| 5:26.518 | אלהים |  |
| 5:26.519 | את |  |
| 5:26.520 | כל |  |
| 5:26.521 | אשר |  |
| 5:26.522 | עשה | , |
| 5:26.523 | התחתונים |  |
| 5:26.524 | העליונים |  |
| 5:26.525 | בסקירה |  |
| 5:26.526 | אחת | , |
| 5:26.527 | העולם |  |
| 5:26.528 | הזה |  |
| 5:26.529 | והעולם |  |
| 5:26.530 | הבא | , |
| 5:26.531 | והוא |  |
| 5:26.532 | באחד |  |
| 5:26.533 | ומיוחד | , |
| 5:26.534 | ברא |  |
| 5:26.535 | את |  |
| 5:26.536 | העולם |  |
| 5:26.537 | בולא |  |
| 5:26.538 | כלום | , |
| 5:26.539 | ברצון |  |
| 5:26.540 | חפצו |  |
| 5:26.541 | ונפשו |  |
| 5:26.542 | אותה |  |
| 5:26.543 | ויעש | , |
| 5:26.544 | ולא |  |
| 5:26.545 | שלא |  |
| 5:26.546 | כרצונו | , |
| 5:26.547 | בדבר |  |
| 5:26.548 | ה | ' |
| 5:26.549 | שמים |  |
| 5:26.550 | ונעשו | , |
| 5:26.551 | הבורא |  |
| 5:26.552 | אינו |  |
| 5:26.553 | צריך |  |
| 5:26.554 | למקום |  |
| 5:26.555 | ומכון | , |
| 5:26.556 | כי |  |
| 5:26.557 | היה |  |
| 5:26.558 | קדם |  |
| 5:26.559 | כל |  |
| 5:26.560 | הויות | , |
| 5:26.561 | ואין |  |
| 5:26.562 | קירות |  |
| 5:26.563 | והקורות |  |
| 5:26.564 | חוצצין |  |
| 5:26.565 | לפניו | . |
| 5:26.566 | כי |  |
| 5:26.567 | לא |  |
| 5:26.568 | היה |  |
| 5:26.569 | בורא |  |
| 5:26.570 | דבר |  |
| 5:26.571 | שהוא |  |
| 5:26.572 | מזיק |  |
| 5:26.573 | כנגדו | . |
| 5:26.574 | יתברך |  |
| 5:26.575 | הבורא |  |
| 5:26.576 | והיוצר |  |
| 5:26.577 | כל | , |
| 5:26.578 | דק |  |
| 5:26.579 | מכל |  |
| 5:26.580 | דק | , |
| 5:26.581 | נפלא |  |
| 5:26.582 | ונעלם |  |
| 5:26.583 | וחבוי |  |
| 5:26.584 | וספון | . |
| 5:26.585 | כל |  |
| 5:26.586 | דבר |  |
| 5:26.587 | בריאה |  |
| 5:26.588 | אין |  |
| 5:26.589 | שוה |  |
| 5:26.590 | לבורא | . |
| 5:26.591 | אין |  |
| 5:26.592 | שייך |  |
| 5:26.593 | בו |  |
| 5:26.594 | שיעור |  |
| 5:26.595 | וקצב | , |
| 5:26.596 | ולא |  |
| 5:26.597 | אורך |  |
| 5:26.598 | ולא |  |
| 5:26.599 | רוחב |  |
| 5:26.600 | ולא |  |
| 5:26.601 | דמות |  |
| 5:26.602 | ולא |  |
| 5:26.603 | צורת |  |
| 5:26.604 | גולם | , |
| 5:26.605 | ולא |  |
| 5:26.606 | דיבוק |  |
| 5:26.607 | ולא |  |
| 5:26.608 | איברים | , |
| 5:26.609 | ולא |  |
| 5:26.610 | צל |  |
| 5:26.611 | ולא |  |
| 5:26.612 | אור | , |
| 5:26.613 | ולא |  |
| 5:26.614 | דמות |  |
| 5:26.615 | ולא |  |
| 5:26.616 | צורה | , |
| 5:26.617 | ולא |  |
| 5:26.618 | סיבוב | . |
| 5:26.619 | וכפי |  |
| 5:26.620 | רצונו |  |
| 5:26.621 | משמיע |  |
| 5:26.622 | קול |  |
| 5:26.623 | לפני |  |
| 5:26.624 | כבודו |  |
| 5:26.625 | המבוררה | , |
| 5:26.626 | המדמה |  |
| 5:26.627 | עצמו |  |
| 5:26.628 | כפי |  |
| 5:26.629 | צורך |  |
| 5:26.630 | השעה | . |
| 5:26.631 | הבורא |  |
| 5:26.632 | הוא |  |
| 5:26.633 | ההוה |  |
| 5:26.634 | בכל | , |
| 5:26.635 | הבורא |  |
| 5:26.636 | אין |  |
| 5:26.637 | לו |  |
| 5:26.638 | דמות |  |
| 5:26.639 | וצורה | , |
| 5:26.640 | ולא |  |
| 5:26.641 | מקרה |  |
| 5:26.642 | ולא |  |
| 5:26.643 | גוף | , |
| 5:26.644 | ולא |  |
| 5:26.645 | שיתוף |  |
| 5:26.646 | וסמך |  |
| 5:26.647 | ודיבוק | , |
| 5:26.648 | ואינו |  |
| 5:26.649 | צריך |  |
| 5:26.650 | לכלום | , |
| 5:26.651 | ולא |  |
| 5:26.652 | למעון |  |
| 5:26.653 | ולא |  |
| 5:26.654 | למקום | , |
| 5:26.655 | ואינו |  |
| 5:26.656 | דבוק |  |
| 5:26.657 | למקום | , |
| 5:26.658 | וממלא |  |
| 5:26.659 | כל |  |
| 5:26.660 | השמים |  |
| 5:26.661 | והארץ | , |
| 5:26.662 | והים |  |
| 5:26.663 | והאויר |  |
| 5:26.664 | וכל |  |
| 5:26.665 | העולמים | , |
| 5:26.666 | ואין |  |
| 5:26.667 | שייך |  |
| 5:26.668 | בו |  |
| 5:26.669 | לא |  |
| 5:26.670 | עמידה |  |
| 5:26.671 | ולא |  |
| 5:26.672 | ישיבה |  |
| 5:26.673 | ולא |  |
| 5:26.674 | הליכה |  |
| 5:26.675 | ולא |  |
| 5:26.676 | סמיכה | , |
| 5:26.677 | והוא |  |
| 5:26.678 | חי |  |
| 5:26.679 | בכחו |  |
| 5:26.680 | ויכלתו |  |
| 5:26.681 | כל |  |
| 5:26.682 | הימים | , |
| 5:26.683 | ואין |  |
| 5:26.684 | לו |  |
| 5:26.685 | נדנוד |  |
| 5:26.686 | ולא |  |
| 5:26.687 | יגיעה |  |
| 5:26.688 | ולא |  |
| 5:26.689 | פעולת |  |
| 5:26.690 | כלי |  |
| 5:26.691 | הוא |  |
| 5:26.692 | בורא |  |
| 5:26.693 | כל | , |
| 5:26.694 | בלא |  |
| 5:26.695 | עמל |  |
| 5:26.696 | ולא |  |
| 5:26.697 | יגיעה | , |
| 5:26.698 | כל |  |
| 5:26.699 | עיני |  |
| 5:26.700 | הנביאים |  |
| 5:26.701 | וכל |  |
| 5:26.702 | עין |  |
| 5:26.703 | לא |  |
| 5:26.704 | תשורנו |  |
| 5:26.705 | ולא |  |
| 5:26.706 | תשופנו | . |
| 5:26.707 | הבורא |  |
| 5:26.708 | היה |  |
| 5:26.709 | טרם |  |
| 5:26.710 | כל | , |
| 5:26.711 | מכון |  |
| 5:26.712 | וקדם |  |
| 5:26.713 | לכל | , |
| 5:26.714 | ראשון |  |
| 5:26.715 | לכל |  |
| 5:26.716 | חי |  |
| 5:26.717 | וקיים |  |
| 5:26.718 | גדול |  |
| 5:26.719 | נבור |  |
| 5:26.720 | ונורא | , |
| 5:26.721 | מלך |  |
| 5:26.722 | מהולל |  |
| 5:26.723 | ומשובח | , |
| 5:26.724 | ואין |  |
| 5:26.725 | בו |  |
| 5:26.726 | לא |  |
| 5:26.727 | תוספת |  |
| 5:26.728 | ולא |  |
| 5:26.729 | חסרון | , |
| 5:26.730 | ולא |  |
| 5:26.731 | מיעוט |  |
| 5:26.732 | ולא |  |
| 5:26.733 | יתרון | , |
| 5:26.734 | אין |  |
| 5:26.735 | קץ |  |
| 5:26.736 | להוייתו | . |
| 5:26.737 | החקר |  |
| 5:26.738 | אלוה |  |
| 5:26.739 | תמצא | , |
| 5:26.740 | אם |  |
| 5:26.741 | עד |  |
| 5:26.742 | תכלית |  |
| 5:26.743 | שדי |  |
| 5:26.744 | המצא | . |
| 5:26.745 | לתבונתו |  |
| 5:26.746 | אין |  |
| 5:26.747 | מספר | , |
| 5:26.748 | לגדולתו |  |
| 5:26.749 | אין |  |
| 5:26.750 | חקר | . |
| 5:26.751 | הבורא |  |
| 5:26.752 | רואה |  |
| 5:26.753 | את |  |
| 5:26.754 | הכל | . |
| 5:26.755 | עין |  |
| 5:26.756 | רואה |  |
| 5:26.757 | ואוזן |  |
| 5:26.758 | שומעת | . |
| 5:26.759 | על |  |
| 5:26.760 | כן |  |
| 5:26.761 | הצנע |  |
| 5:26.762 | לכת |  |
| 5:26.763 | בבית |  |
| 5:26.764 | הכסא |  |
| 5:26.765 | ועם |  |
| 5:26.766 | אשתך | . |
| 5:26.767 | כשתחשוב |  |
| 5:26.768 | על |  |
| 5:26.769 | בורא |  |
| 5:26.770 | עולם |  |
| 5:26.771 | מה |  |
| 5:26.772 | הוא | , |
| 5:26.773 | והיאך |  |
| 5:26.774 | חנייתו |  |
| 5:26.775 | בכל |  |
| 5:26.776 | מקום |  |
| 5:26.777 | ומעשיהו | . |
| 5:26.778 | בלום |  |
| 5:26.779 | פיך |  |
| 5:26.780 | מלדבר | , |
| 5:26.781 | ולבר |  |
| 5:26.782 | מלהרהר | . |
| 5:26.783 | הסר |  |
| 5:26.784 | המחשבה |  |
| 5:26.785 | מלבך | . |
| 5:26.786 | ואם |  |
| 5:26.787 | סר |  |
| 5:26.788 | לבך |  |
| 5:26.789 | למחשבה |  |
| 5:26.790 | זאת | , |
| 5:26.791 | חושה |  |
| 5:26.792 | ומהר | , |
| 5:26.793 | ואל |  |
| 5:26.794 | תהרהר | , |
| 5:26.795 | ושוב |  |
| 5:26.796 | ליחוד |  |
| 5:26.797 | מקומו |  |
| 5:26.798 | שלעולם | , |
| 5:26.799 | לעבודתו |  |
| 5:26.800 | וליראתו | , |
| 5:26.801 | בחיות |  |
| 5:26.802 | הקדש |  |
| 5:26.803 | כשרוצין |  |
| 5:26.804 | לזקוף |  |
| 5:26.805 | קומתן |  |
| 5:26.806 | מיד |  |
| 5:26.807 | כופפין |  |
| 5:26.808 | מאימתו | , |
| 5:26.809 | ושבים |  |
| 5:26.810 | ומשתחוים | . |
| 5:26.811 | כך |  |
| 5:26.812 | סבב |  |
| 5:26.813 | לבך |  |
| 5:26.814 | מלהרהר | , |
| 5:26.815 | ועל |  |
| 5:26.816 | דבר |  |
| 5:26.817 | זה |  |
| 5:26.818 | נכרתה |  |
| 5:26.819 | ברית |  |
| 5:26.820 | שלא |  |
| 5:26.821 | לחשוב |  |
| 5:26.822 | באלהותו | . |
| 5:26.823 | שאין |  |
| 5:26.824 | כל |  |
| 5:26.825 | חכמים |  |
| 5:26.826 | יכולים |  |
| 5:26.827 | לידע | . |
| 5:26.828 | ושוב |  |
| 5:26.829 | וחשוב |  |
| 5:26.830 | כי |  |
| 5:26.831 | הוא |  |
| 5:26.832 | יוצר |  |
| 5:26.833 | הכל | , |
| 5:26.834 | והוא |  |
| 5:26.835 | הבורא | , |
| 5:26.836 | האמן |  |
| 5:26.837 | באלהותו |  |
| 5:26.838 | כי |  |
| 5:26.839 | הוא |  |
| 5:26.840 | אחד | , |
| 5:26.841 | יחיד |  |
| 5:26.842 | הבורא | . |
| 5:26.843 | ומראה |  |
| 5:26.844 | הכבוד |  |
| 5:26.845 | כפי |  |
| 5:26.846 | רצונו |  |
| 5:26.847 | וחפצו | . |
| 5:26.848 | מראה |  |
| 5:26.849 | כבוד |  |
| 5:26.850 | ה | ' |
| 5:26.851 | כאש |  |
| 5:26.852 | אוכלת | , |
| 5:26.853 | וקראו |  |
| 5:26.854 | השכינה | . |
| 5:26.855 | ופעמים |  |
| 5:26.856 | נראה |  |
| 5:26.857 | בלא |  |
| 5:26.858 | צורה | . |
| 5:26.859 | כי |  |
| 5:26.860 | אם |  |
| 5:26.861 | אורה |  |
| 5:26.862 | בלא |  |
| 5:26.863 | ראות |  |
| 5:26.864 | לכל |  |
| 5:26.865 | יצירה | . |
| 5:26.866 | ושומע |  |
| 5:26.867 | הדבור |  |
| 5:26.868 | ורואה |  |
| 5:26.869 | החזון |  |
| 5:26.870 | מראה |  |
| 5:26.871 | מופלאה | . |
| 5:26.872 | ויאמר |  |
| 5:26.873 | ראיתי |  |
| 5:26.874 | את |  |
| 5:26.875 | ה | ', |
| 5:26.876 | את |  |
| 5:26.877 | לרכות |  |
| 5:26.878 | כבודו |  |
| 5:26.879 | וכו | ' |
| 5:26.880 | ומראה |  |
| 5:26.881 | דמות |  |
| 5:26.882 | כבודו |  |
| 5:26.883 | כחפץ |  |
| 5:26.884 | גזרותיו | . |
| 5:26.885 | פעמים |  |
| 5:26.886 | דמות |  |
| 5:26.887 | אדם | , |
| 5:26.888 | ופעמים |  |
| 5:26.889 | דמות |  |
| 5:26.890 | אחר |  |
| 5:26.891 | כי |  |
| 5:26.892 | חפצו |  |
| 5:26.893 | מראה |  |
| 5:26.894 | כבודו | , |
| 5:26.895 | כצבאות |  |
| 5:26.896 | חיילות |  |
| 5:26.897 | דבוקים |  |
| 5:26.898 | יחד | , |
| 5:26.899 | ה | ' |
| 5:26.900 | צבאות | , |
| 5:26.901 | ופעמים |  |
| 5:26.902 | גופים |  |
| 5:26.903 | מחולקים | , |
| 5:26.904 | ה | ' |
| 5:26.905 | אלהי |  |
| 5:26.906 | הצבאות | , |
| 5:26.907 | ופעמים |  |
| 5:26.908 | נקראים |  |
| 5:26.909 | בני |  |
| 5:26.910 | האלהים | , |
| 5:26.911 | ופעמים |  |
| 5:26.912 | עירין |  |
| 5:26.913 | וקדישין | , |
| 5:26.914 | והדבור |  |
| 5:26.915 | בא |  |
| 5:26.916 | כמפי |  |
| 5:26.917 | הכבוד |  |
| 5:26.918 | הגדול | , |
| 5:26.919 | הזהר |  |
| 5:26.920 | פן |  |
| 5:26.921 | תכזב |  |
| 5:26.922 | עליו | . |
| 5:26.923 | השמר |  |
| 5:26.924 | לך |  |
| 5:26.925 | ושמור |  |
| 5:26.926 | נפשך |  |
| 5:26.927 | מאד |  |
| 5:26.928 | פן |  |
| 5:26.929 | תחטיא |  |
| 5:26.930 | את |  |
| 5:26.931 | בשרך |  |
| 5:26.932 | לומר |  |
| 5:26.933 | עליו |  |
| 5:26.934 | חלוקת |  |
| 5:26.935 | איברים | , |
| 5:26.936 | מה |  |
| 5:26.937 | דמות |  |
| 5:26.938 | תערכו |  |
| 5:26.939 | לו | , |
| 5:26.940 | כי |  |
| 5:26.941 | לא |  |
| 5:26.942 | ראיתם |  |
| 5:26.943 | כל |  |
| 5:26.944 | תמונה | , |
| 5:26.945 | ומה |  |
| 5:26.946 | שכתוב |  |
| 5:26.947 | ויראו |  |
| 5:26.948 | את |  |
| 5:26.949 | אלהי |  |
| 5:26.950 | ישראל | , |
| 5:26.951 | הכבוד |  |
| 5:26.952 | ראו |  |
| 5:26.953 | כמו |  |
| 5:26.954 | שתרגם |  |
| 5:26.955 | אונקלוס | ) |
| 5:26.956 | ובפרק |  |
| 5:26.957 | קמא |  |
| 5:26.958 | דברכות |  |
| 5:26.959 | מה |  |
| 5:26.960 | הקב"ה |  |
| 5:26.961 | רואה |  |
| 5:26.962 | ואינו |  |
| 5:26.963 | נראה | , |
| 5:26.964 | אף |  |
| 5:26.965 | הנשמה |  |
| 5:26.966 | רואה |  |
| 5:26.967 | ואינה |  |
| 5:26.968 | נראית | , |
| 5:26.969 | מנין |  |
| 5:26.970 | שהקב"ה |  |
| 5:26.971 | מניח |  |
| 5:26.972 | תפילין | , |
| 5:26.973 | פי | ' |
| 5:26.974 | מראה |  |
| 5:26.975 | מכבודו |  |
| 5:26.976 | באובנתא |  |
| 5:26.977 | דליבא | , |
| 5:26.978 | כמו |  |
| 5:26.979 | וירא |  |
| 5:26.980 | יעקב |  |
| 5:26.981 | כי |  |
| 5:26.982 | יש |  |
| 5:26.983 | שבר |  |
| 5:26.984 | במצרים | , |
| 5:26.985 | ולבי |  |
| 5:26.986 | ראה |  |
| 5:26.987 | הרבה | , |
| 5:26.988 | כי |  |
| 5:26.989 | לא |  |
| 5:26.990 | יראני |  |
| 5:26.991 | האדם |  |
| 5:26.992 | וחי | , |
| 5:26.993 | ר | ' |
| 5:26.994 | עקיבה |  |
| 5:26.995 | אומר |  |
| 5:26.996 | כמשמעו | , |
| 5:26.997 | וחי |  |
| 5:26.998 | אלו |  |
| 5:26.999 | מלאכי |  |
| 5:26.1000 | השרת |  |
| 5:26.1001 | וחיות |  |
| 5:26.1002 | הקדש | , |
| 5:26.1003 | ר | ' |
| 5:26.1004 | אליעזר |  |
| 5:26.1005 | אומר |  |
| 5:26.1006 | וחי |  |
| 5:26.1007 | אינו |  |
| 5:26.1008 | רואה |  |
| 5:26.1009 | אבל |  |
| 5:26.1010 | רואה |  |
| 5:26.1011 | בשעת |  |
| 5:26.1012 | מיתה | , |
| 5:26.1013 | ובספרי |  |
| 5:26.1014 | וחי |  |
| 5:26.1015 | בחייהן |  |
| 5:26.1016 | אינן |  |
| 5:26.1017 | רואין | , |
| 5:26.1018 | אבל |  |
| 5:26.1019 | רואין |  |
| 5:26.1020 | בשעת |  |
| 5:26.1021 | מיתה | , |
| 5:26.1022 | ר | ' |
| 5:26.1023 | עקיבה |  |
| 5:26.1024 | אומר |  |
| 5:26.1025 | וחי | , |
| 5:26.1026 | חיות |  |
| 5:26.1027 | הנושאות |  |
| 5:26.1028 | את |  |
| 5:26.1029 | הכסא |  |
| 5:26.1030 | אינן |  |
| 5:26.1031 | רואות |  |
| 5:26.1032 | את |  |
| 5:26.1033 | הכבוד | , |
| 5:26.1034 | אמר |  |
| 5:26.1035 | רשב"י |  |
| 5:26.1036 | אף |  |
| 5:26.1037 | מלאכי |  |
| 5:26.1038 | השרת |  |
| 5:26.1039 | שחייהן |  |
| 5:26.1040 | חיי |  |
| 5:26.1041 | עולם |  |
| 5:26.1042 | אינן |  |
| 5:26.1043 | רואים |  |
| 5:26.1044 | את |  |
| 5:26.1045 | הכבוד | . |
| 5:26.1046 | והאריך |  |
| 5:26.1047 | עוד |  |
| 5:26.1048 | הרוקח |  |
| 5:26.1049 | ז"ל |  |
| 5:26.1050 | ומסיים |  |
| 5:26.1051 | שם | , |
| 5:26.1052 | הבורא |  |
| 5:26.1053 | הזהירנו |  |
| 5:26.1054 | ליחדו |  |
| 5:26.1055 | שני |  |
| 5:26.1056 | פעמים |  |
| 5:26.1057 | ביום |  |
| 5:26.1058 | ה | ' |
| 5:26.1059 | אלהינו |  |
| 5:26.1060 | ה | ' |
| 5:26.1061 | אחד | , |
| 5:26.1062 | בקבלת |  |
| 5:26.1063 | הלב |  |
| 5:26.1064 | האמנתו |  |
| 5:26.1065 | ולא |  |
| 5:26.1066 | בשמיעת |  |
| 5:26.1067 | האזן | , |
| 5:26.1068 | בכאן |  |
| 5:26.1069 | נקבלהו |  |
| 5:26.1070 | באלוה | , |
| 5:26.1071 | זה |  |
| 5:26.1072 | אלהינו | , |
| 5:26.1073 | ולהאמין |  |
| 5:26.1074 | שהוא |  |
| 5:26.1075 | אחד | , |
| 5:26.1076 | ויחוד |  |
| 5:26.1077 | שמו | , |
| 5:26.1078 | וידעת |  |
| 5:26.1079 | היום |  |
| 5:26.1080 | והשבות |  |
| 5:26.1081 | אל |  |
| 5:26.1082 | לבביך |  |
| 5:26.1083 | כי |  |
| 5:26.1084 | ה | ' |
| 5:26.1085 | הוא |  |
| 5:26.1086 | האלהים |  |
| 5:26.1087 | בשמים |  |
| 5:26.1088 | ממעל |  |
| 5:26.1089 | ועל |  |
| 5:26.1090 | הארץ |  |
| 5:26.1091 | מתחת |  |
| 5:26.1092 | אין |  |
| 5:26.1093 | עוד | , |
| 5:26.1094 | הוא |  |
| 5:26.1095 | קדמון |  |
| 5:26.1096 | וחי |  |
| 5:26.1097 | לעולמים | , |
| 5:26.1098 | ואין |  |
| 5:26.1099 | לו |  |
| 5:26.1100 | שותף | , |
| 5:26.1101 | והוא |  |
| 5:26.1102 | אחד |  |
| 5:26.1103 | ואין |  |
| 5:26.1104 | לדמותו | , |
| 5:26.1105 | אל |  |
| 5:26.1106 | דמות |  |
| 5:26.1107 | וצורה | , |
| 5:26.1108 | ונשמרתם |  |
| 5:26.1109 | מאד |  |
| 5:26.1110 | לנשפותיכם |  |
| 5:26.1111 | כי |  |
| 5:26.1112 | לא |  |
| 5:26.1113 | ראיתם |  |
| 5:26.1114 | כל |  |
| 5:26.1115 | תמונה | , |
| 5:26.1116 | ומה |  |
| 5:26.1117 | דמות |  |
| 5:26.1118 | תערכו |  |
| 5:26.1119 | לו | . |
| 5:26.1120 | ומשמותיו |  |
| 5:26.1121 | ומשבחיו |  |
| 5:26.1122 | ומעשיו |  |
| 5:26.1123 | לא |  |
| 5:26.1124 | נוכל |  |
| 5:26.1125 | לדעת | . |
| 5:26.1126 | כי |  |
| 5:26.1127 | אין |  |
| 5:26.1128 | כמוהו | . |
| 5:26.1129 | ברוך |  |
| 5:26.1130 | הוא |  |
| 5:26.1131 | וברוך |  |
| 5:26.1132 | שמו |  |
| 5:26.1133 | ע"כ |  |
| 5:26.1134 | לשון |  |
| 5:26.1135 | הרקח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

