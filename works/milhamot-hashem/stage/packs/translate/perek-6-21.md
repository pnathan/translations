# Translation pack: milhamot-hashem perek-6-21 (Milchamot ha-Shem 6:32–6:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-21`.

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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.

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


## Milchamot ha-Shem 6:32

Hebrew (Wikisource): עה) ודבריו אלה שכתב שנקראו אלהים אחרים לפי שאין להם קיום מעצמן רק על ידי אחרים, נובעים ממקור נאמן מדרש רבה והילקוט פרשת ויצא ופ' מקץ מדרש הגדול בויהי מקץ, וז"ל מדרש רבה עפ פי' יפה תואר ועץ יוסף, אמר ר"י הרשעים מתקיימין על אלהיהם, ופרעה חלם והנה עומד על היאור, אבל הצדיקים אלהיהם מתקיים עליהם והנה ה' נצב עליו, (פי' ההתיצבות היא ההשגחה להצילו מכל נזק ופגע וזהו שאמר ר"י הרשעים צריכים לעמוד על אלהיהם ולשמרם. אבל הצדיקים אלהיהם נצב עליהם לשמרם ולהצילם ופי' נצב משגיח ע"כ) ובמדרש הגדול גורס במקום מתקיימין עומדים. ובילקוט פר' יתרו. אלהים אחרים. וכי אלהות הן והלא כבר נאמר ונתן אלהיהם באש כי לא אלהים המה. ומה תלמוד לומר אלהים אחרים. שהם מאחרים לעובדיהם ובן הוא אומר הן יצעק אליו ולא יענה ומצרתו לא יושיעו, ד"א אלהים אחרים שהם עושים את עובדיהן אחרים. ר' יוסי אומר אלהים אחרים למה נאמר שלא ליתן פתחון פה לאומות העולם לומר אלו נקראו בשמו כבר היה בהם צורך הרי נקראו בשמו ואין בהם צורך. אימתי נקראו בשמו, בימי אנוש בן שת שנ' אז הוחל לקרא בשם ה'. באותה שעה אמר הקב"ה אתם עשיתם מעשה חדש וקראתם אותם אלוהות, אני אעשה מעשה חדש ואקרא עצמי ה', באותה שעה קלה עלה אוקיינוס והציף שלישו של עולם וכן הוא אומר הקורא למי הים וישפכם על פני הארץ ה' שמו, ר' אלעזר המודעי אומר אלהים אחרים שהם מחדשים להם אלוהות בכל יום, הא כיצד היה לו של זהב והנצרך לו עושהו של כסף. נצרך לו עושהו של נחשת, וכן ברזל, וכן בדיל ועופרת, שני חדשים מקרוב באו וכו' ע"ש. והנה מכל זה תבין שגגת רבותינו בעלי הקבלה החדשה שהאמינו לדברי המסית מחבר הזהר במציאות אלהים אחרים בעלי עשר ספירות וחמשה פרצופים דקליפה נגד עשר ספירות ופרצופין של קדושה, ואמרו נבראת לעומת זה עשה האלהים, גם היכלות הטומאה עומת היכלות הקדושה, בכל פרצוף שלה (משנת חסידים מסכת היכלות הקליפה וס' מחברת הקדש שער ר"ח ד' ט"ז ע"ב). וסוברים שכמו שאצל וברא ויצר ועשה אלהים קדושים כן אצל וברא ויצר ועשה אלהים אחרים וכמו שיש בקדושה אדם קדמון ועתיק ואריך, ואבא ואימא, וזעיר ונוקביה דקדושה כן בשמאל אדם בליעל סטרא אחרא מסאבא, זה לעומת זה כמו שיראה המעיין בס' חיים למהרח"ו, ומשנת חסידים שער הקליפות. ואם תעמיק חקר תבין ותשכיל שסביא זו באה ונמשה להם מאמונת השונים שסוברים שיש בעולם שני פועלים, אחד פועל הטוב, ואחד פועל הרע. כמו שכתבו הראשונים לדחותה בשתי ידים, והמקובלים הרדשים תפשוה וציירו אותה בפנים אחרים, ונדחקו מאד לכסות את נבלותה, ולהלבישה באמונת תוח"ק. ובאמת זרה היא עד מאד לאמונת יהוד השי"ת, ולא תבוא בקהל ה'. ולאו קבלה היא כלל. כמו שכ' לעיל מדברי רז"ל. כי המקובלים החדשים מאמתים מציאות אל אחד בעולם היפך דברי רז"ל. והכתובים מוכיחים כדעת חכמים ז"ל. שנ' כי הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד מלבדו. אתה הוא ה' לבדך. וכמה כתובים מכריזים שאין עוד אלוה אחר זולתי ה' אלהינו. ורבותינו מתקני התפלות התקינו לומר בכל יום אנו יודעים ומעידים ומגידים, שאין אלוה זולתך וכו' והרבה לשונות בתפלה מורים שאין מציאות אל אחר בעולם כלל כלל. ורק לפי מחשבת ובדיהן. אמרה כי תורה כי לא תשתחוה לאל אחר. וכן לא יהיה לך אלהים אחרים על פני. הכל לפי מחשבת עובידהן שטועין וחושבים אותם אלוהות. לא שהשי"ת המציא שני מיני אלוהות אלה קדושים ואלה טמאים ח"ו. כי אין אלוה זולתו. וכמה פסורים בתורה ובנביאים ובכתובים מוכיחים כן.

| id | word | punctuation after |
|---|---|---|
| 6:32.1 | עה | ) |
| 6:32.2 | ודבריו |  |
| 6:32.3 | אלה |  |
| 6:32.4 | שכתב |  |
| 6:32.5 | שנקראו |  |
| 6:32.6 | אלהים |  |
| 6:32.7 | אחרים |  |
| 6:32.8 | לפי |  |
| 6:32.9 | שאין |  |
| 6:32.10 | להם |  |
| 6:32.11 | קיום |  |
| 6:32.12 | מעצמן |  |
| 6:32.13 | רק |  |
| 6:32.14 | על |  |
| 6:32.15 | ידי |  |
| 6:32.16 | אחרים | , |
| 6:32.17 | נובעים |  |
| 6:32.18 | ממקור |  |
| 6:32.19 | נאמן |  |
| 6:32.20 | מדרש |  |
| 6:32.21 | רבה |  |
| 6:32.22 | והילקוט |  |
| 6:32.23 | פרשת |  |
| 6:32.24 | ויצא |  |
| 6:32.25 | ופ | ' |
| 6:32.26 | מקץ |  |
| 6:32.27 | מדרש |  |
| 6:32.28 | הגדול |  |
| 6:32.29 | בויהי |  |
| 6:32.30 | מקץ | , |
| 6:32.31 | וז"ל |  |
| 6:32.32 | מדרש |  |
| 6:32.33 | רבה |  |
| 6:32.34 | עפ |  |
| 6:32.35 | פי | ' |
| 6:32.36 | יפה |  |
| 6:32.37 | תואר |  |
| 6:32.38 | ועץ |  |
| 6:32.39 | יוסף | , |
| 6:32.40 | אמר |  |
| 6:32.41 | ר"י |  |
| 6:32.42 | הרשעים |  |
| 6:32.43 | מתקיימין |  |
| 6:32.44 | על |  |
| 6:32.45 | אלהיהם | , |
| 6:32.46 | ופרעה |  |
| 6:32.47 | חלם |  |
| 6:32.48 | והנה |  |
| 6:32.49 | עומד |  |
| 6:32.50 | על |  |
| 6:32.51 | היאור | , |
| 6:32.52 | אבל |  |
| 6:32.53 | הצדיקים |  |
| 6:32.54 | אלהיהם |  |
| 6:32.55 | מתקיים |  |
| 6:32.56 | עליהם |  |
| 6:32.57 | והנה |  |
| 6:32.58 | ה | ' |
| 6:32.59 | נצב |  |
| 6:32.60 | עליו | , |
| 6:32.61 | פי | ' |
| 6:32.62 | ההתיצבות |  |
| 6:32.63 | היא |  |
| 6:32.64 | ההשגחה |  |
| 6:32.65 | להצילו |  |
| 6:32.66 | מכל |  |
| 6:32.67 | נזק |  |
| 6:32.68 | ופגע |  |
| 6:32.69 | וזהו |  |
| 6:32.70 | שאמר |  |
| 6:32.71 | ר"י |  |
| 6:32.72 | הרשעים |  |
| 6:32.73 | צריכים |  |
| 6:32.74 | לעמוד |  |
| 6:32.75 | על |  |
| 6:32.76 | אלהיהם |  |
| 6:32.77 | ולשמרם | . |
| 6:32.78 | אבל |  |
| 6:32.79 | הצדיקים |  |
| 6:32.80 | אלהיהם |  |
| 6:32.81 | נצב |  |
| 6:32.82 | עליהם |  |
| 6:32.83 | לשמרם |  |
| 6:32.84 | ולהצילם |  |
| 6:32.85 | ופי | ' |
| 6:32.86 | נצב |  |
| 6:32.87 | משגיח |  |
| 6:32.88 | ע"כ | ) |
| 6:32.89 | ובמדרש |  |
| 6:32.90 | הגדול |  |
| 6:32.91 | גורס |  |
| 6:32.92 | במקום |  |
| 6:32.93 | מתקיימין |  |
| 6:32.94 | עומדים | . |
| 6:32.95 | ובילקוט |  |
| 6:32.96 | פר | ' |
| 6:32.97 | יתרו | . |
| 6:32.98 | אלהים |  |
| 6:32.99 | אחרים | . |
| 6:32.100 | וכי |  |
| 6:32.101 | אלהות |  |
| 6:32.102 | הן |  |
| 6:32.103 | והלא |  |
| 6:32.104 | כבר |  |
| 6:32.105 | נאמר |  |
| 6:32.106 | ונתן |  |
| 6:32.107 | אלהיהם |  |
| 6:32.108 | באש |  |
| 6:32.109 | כי |  |
| 6:32.110 | לא |  |
| 6:32.111 | אלהים |  |
| 6:32.112 | המה | . |
| 6:32.113 | ומה |  |
| 6:32.114 | תלמוד |  |
| 6:32.115 | לומר |  |
| 6:32.116 | אלהים |  |
| 6:32.117 | אחרים | . |
| 6:32.118 | שהם |  |
| 6:32.119 | מאחרים |  |
| 6:32.120 | לעובדיהם |  |
| 6:32.121 | ובן |  |
| 6:32.122 | הוא |  |
| 6:32.123 | אומר |  |
| 6:32.124 | הן |  |
| 6:32.125 | יצעק |  |
| 6:32.126 | אליו |  |
| 6:32.127 | ולא |  |
| 6:32.128 | יענה |  |
| 6:32.129 | ומצרתו |  |
| 6:32.130 | לא |  |
| 6:32.131 | יושיעו | , |
| 6:32.132 | ד"א |  |
| 6:32.133 | אלהים |  |
| 6:32.134 | אחרים |  |
| 6:32.135 | שהם |  |
| 6:32.136 | עושים |  |
| 6:32.137 | את |  |
| 6:32.138 | עובדיהן |  |
| 6:32.139 | אחרים | . |
| 6:32.140 | ר | ' |
| 6:32.141 | יוסי |  |
| 6:32.142 | אומר |  |
| 6:32.143 | אלהים |  |
| 6:32.144 | אחרים |  |
| 6:32.145 | למה |  |
| 6:32.146 | נאמר |  |
| 6:32.147 | שלא |  |
| 6:32.148 | ליתן |  |
| 6:32.149 | פתחון |  |
| 6:32.150 | פה |  |
| 6:32.151 | לאומות |  |
| 6:32.152 | העולם |  |
| 6:32.153 | לומר |  |
| 6:32.154 | אלו |  |
| 6:32.155 | נקראו |  |
| 6:32.156 | בשמו |  |
| 6:32.157 | כבר |  |
| 6:32.158 | היה |  |
| 6:32.159 | בהם |  |
| 6:32.160 | צורך |  |
| 6:32.161 | הרי |  |
| 6:32.162 | נקראו |  |
| 6:32.163 | בשמו |  |
| 6:32.164 | ואין |  |
| 6:32.165 | בהם |  |
| 6:32.166 | צורך | . |
| 6:32.167 | אימתי |  |
| 6:32.168 | נקראו |  |
| 6:32.169 | בשמו | , |
| 6:32.170 | בימי |  |
| 6:32.171 | אנוש |  |
| 6:32.172 | בן |  |
| 6:32.173 | שת |  |
| 6:32.174 | שנ | ' |
| 6:32.175 | אז |  |
| 6:32.176 | הוחל |  |
| 6:32.177 | לקרא |  |
| 6:32.178 | בשם |  |
| 6:32.179 | ה | '. |
| 6:32.180 | באותה |  |
| 6:32.181 | שעה |  |
| 6:32.182 | אמר |  |
| 6:32.183 | הקב"ה |  |
| 6:32.184 | אתם |  |
| 6:32.185 | עשיתם |  |
| 6:32.186 | מעשה |  |
| 6:32.187 | חדש |  |
| 6:32.188 | וקראתם |  |
| 6:32.189 | אותם |  |
| 6:32.190 | אלוהות | , |
| 6:32.191 | אני |  |
| 6:32.192 | אעשה |  |
| 6:32.193 | מעשה |  |
| 6:32.194 | חדש |  |
| 6:32.195 | ואקרא |  |
| 6:32.196 | עצמי |  |
| 6:32.197 | ה | ', |
| 6:32.198 | באותה |  |
| 6:32.199 | שעה |  |
| 6:32.200 | קלה |  |
| 6:32.201 | עלה |  |
| 6:32.202 | אוקיינוס |  |
| 6:32.203 | והציף |  |
| 6:32.204 | שלישו |  |
| 6:32.205 | של |  |
| 6:32.206 | עולם |  |
| 6:32.207 | וכן |  |
| 6:32.208 | הוא |  |
| 6:32.209 | אומר |  |
| 6:32.210 | הקורא |  |
| 6:32.211 | למי |  |
| 6:32.212 | הים |  |
| 6:32.213 | וישפכם |  |
| 6:32.214 | על |  |
| 6:32.215 | פני |  |
| 6:32.216 | הארץ |  |
| 6:32.217 | ה | ' |
| 6:32.218 | שמו | , |
| 6:32.219 | ר | ' |
| 6:32.220 | אלעזר |  |
| 6:32.221 | המודעי |  |
| 6:32.222 | אומר |  |
| 6:32.223 | אלהים |  |
| 6:32.224 | אחרים |  |
| 6:32.225 | שהם |  |
| 6:32.226 | מחדשים |  |
| 6:32.227 | להם |  |
| 6:32.228 | אלוהות |  |
| 6:32.229 | בכל |  |
| 6:32.230 | יום | , |
| 6:32.231 | הא |  |
| 6:32.232 | כיצד |  |
| 6:32.233 | היה |  |
| 6:32.234 | לו |  |
| 6:32.235 | של |  |
| 6:32.236 | זהב |  |
| 6:32.237 | והנצרך |  |
| 6:32.238 | לו |  |
| 6:32.239 | עושהו |  |
| 6:32.240 | של |  |
| 6:32.241 | כסף | . |
| 6:32.242 | נצרך |  |
| 6:32.243 | לו |  |
| 6:32.244 | עושהו |  |
| 6:32.245 | של |  |
| 6:32.246 | נחשת | , |
| 6:32.247 | וכן |  |
| 6:32.248 | ברזל | , |
| 6:32.249 | וכן |  |
| 6:32.250 | בדיל |  |
| 6:32.251 | ועופרת | , |
| 6:32.252 | שני |  |
| 6:32.253 | חדשים |  |
| 6:32.254 | מקרוב |  |
| 6:32.255 | באו |  |
| 6:32.256 | וכו | ' |
| 6:32.257 | ע"ש | . |
| 6:32.258 | והנה |  |
| 6:32.259 | מכל |  |
| 6:32.260 | זה |  |
| 6:32.261 | תבין |  |
| 6:32.262 | שגגת |  |
| 6:32.263 | רבותינו |  |
| 6:32.264 | בעלי |  |
| 6:32.265 | הקבלה |  |
| 6:32.266 | החדשה |  |
| 6:32.267 | שהאמינו |  |
| 6:32.268 | לדברי |  |
| 6:32.269 | המסית |  |
| 6:32.270 | מחבר |  |
| 6:32.271 | הזהר |  |
| 6:32.272 | במציאות |  |
| 6:32.273 | אלהים |  |
| 6:32.274 | אחרים |  |
| 6:32.275 | בעלי |  |
| 6:32.276 | עשר |  |
| 6:32.277 | ספירות |  |
| 6:32.278 | וחמשה |  |
| 6:32.279 | פרצופים |  |
| 6:32.280 | דקליפה |  |
| 6:32.281 | נגד |  |
| 6:32.282 | עשר |  |
| 6:32.283 | ספירות |  |
| 6:32.284 | ופרצופין |  |
| 6:32.285 | של |  |
| 6:32.286 | קדושה | , |
| 6:32.287 | ואמרו |  |
| 6:32.288 | נבראת |  |
| 6:32.289 | לעומת |  |
| 6:32.290 | זה |  |
| 6:32.291 | עשה |  |
| 6:32.292 | האלהים | , |
| 6:32.293 | גם |  |
| 6:32.294 | היכלות |  |
| 6:32.295 | הטומאה |  |
| 6:32.296 | עומת |  |
| 6:32.297 | היכלות |  |
| 6:32.298 | הקדושה | , |
| 6:32.299 | בכל |  |
| 6:32.300 | פרצוף |  |
| 6:32.301 | שלה |  |
| 6:32.302 | משנת |  |
| 6:32.303 | חסידים |  |
| 6:32.304 | מסכת |  |
| 6:32.305 | היכלות |  |
| 6:32.306 | הקליפה |  |
| 6:32.307 | וס | ' |
| 6:32.308 | מחברת |  |
| 6:32.309 | הקדש |  |
| 6:32.310 | שער |  |
| 6:32.311 | ר"ח |  |
| 6:32.312 | ד | ' |
| 6:32.313 | ט"ז |  |
| 6:32.314 | ע"ב | ). |
| 6:32.315 | וסוברים |  |
| 6:32.316 | שכמו |  |
| 6:32.317 | שאצל |  |
| 6:32.318 | וברא |  |
| 6:32.319 | ויצר |  |
| 6:32.320 | ועשה |  |
| 6:32.321 | אלהים |  |
| 6:32.322 | קדושים |  |
| 6:32.323 | כן |  |
| 6:32.324 | אצל |  |
| 6:32.325 | וברא |  |
| 6:32.326 | ויצר |  |
| 6:32.327 | ועשה |  |
| 6:32.328 | אלהים |  |
| 6:32.329 | אחרים |  |
| 6:32.330 | וכמו |  |
| 6:32.331 | שיש |  |
| 6:32.332 | בקדושה |  |
| 6:32.333 | אדם |  |
| 6:32.334 | קדמון |  |
| 6:32.335 | ועתיק |  |
| 6:32.336 | ואריך | , |
| 6:32.337 | ואבא |  |
| 6:32.338 | ואימא | , |
| 6:32.339 | וזעיר |  |
| 6:32.340 | ונוקביה |  |
| 6:32.341 | דקדושה |  |
| 6:32.342 | כן |  |
| 6:32.343 | בשמאל |  |
| 6:32.344 | אדם |  |
| 6:32.345 | בליעל |  |
| 6:32.346 | סטרא |  |
| 6:32.347 | אחרא |  |
| 6:32.348 | מסאבא | , |
| 6:32.349 | זה |  |
| 6:32.350 | לעומת |  |
| 6:32.351 | זה |  |
| 6:32.352 | כמו |  |
| 6:32.353 | שיראה |  |
| 6:32.354 | המעיין |  |
| 6:32.355 | בס | ' |
| 6:32.356 | חיים |  |
| 6:32.357 | למהרח"ו | , |
| 6:32.358 | ומשנת |  |
| 6:32.359 | חסידים |  |
| 6:32.360 | שער |  |
| 6:32.361 | הקליפות | . |
| 6:32.362 | ואם |  |
| 6:32.363 | תעמיק |  |
| 6:32.364 | חקר |  |
| 6:32.365 | תבין |  |
| 6:32.366 | ותשכיל |  |
| 6:32.367 | שסביא |  |
| 6:32.368 | זו |  |
| 6:32.369 | באה |  |
| 6:32.370 | ונמשה |  |
| 6:32.371 | להם |  |
| 6:32.372 | מאמונת |  |
| 6:32.373 | השונים |  |
| 6:32.374 | שסוברים |  |
| 6:32.375 | שיש |  |
| 6:32.376 | בעולם |  |
| 6:32.377 | שני |  |
| 6:32.378 | פועלים | , |
| 6:32.379 | אחד |  |
| 6:32.380 | פועל |  |
| 6:32.381 | הטוב | , |
| 6:32.382 | ואחד |  |
| 6:32.383 | פועל |  |
| 6:32.384 | הרע | . |
| 6:32.385 | כמו |  |
| 6:32.386 | שכתבו |  |
| 6:32.387 | הראשונים |  |
| 6:32.388 | לדחותה |  |
| 6:32.389 | בשתי |  |
| 6:32.390 | ידים | , |
| 6:32.391 | והמקובלים |  |
| 6:32.392 | הרדשים |  |
| 6:32.393 | תפשוה |  |
| 6:32.394 | וציירו |  |
| 6:32.395 | אותה |  |
| 6:32.396 | בפנים |  |
| 6:32.397 | אחרים | , |
| 6:32.398 | ונדחקו |  |
| 6:32.399 | מאד |  |
| 6:32.400 | לכסות |  |
| 6:32.401 | את |  |
| 6:32.402 | נבלותה | , |
| 6:32.403 | ולהלבישה |  |
| 6:32.404 | באמונת |  |
| 6:32.405 | תוח"ק | . |
| 6:32.406 | ובאמת |  |
| 6:32.407 | זרה |  |
| 6:32.408 | היא |  |
| 6:32.409 | עד |  |
| 6:32.410 | מאד |  |
| 6:32.411 | לאמונת |  |
| 6:32.412 | יהוד |  |
| 6:32.413 | השי"ת | , |
| 6:32.414 | ולא |  |
| 6:32.415 | תבוא |  |
| 6:32.416 | בקהל |  |
| 6:32.417 | ה | '. |
| 6:32.418 | ולאו |  |
| 6:32.419 | קבלה |  |
| 6:32.420 | היא |  |
| 6:32.421 | כלל | . |
| 6:32.422 | כמו |  |
| 6:32.423 | שכ | ' |
| 6:32.424 | לעיל |  |
| 6:32.425 | מדברי |  |
| 6:32.426 | רז"ל | . |
| 6:32.427 | כי |  |
| 6:32.428 | המקובלים |  |
| 6:32.429 | החדשים |  |
| 6:32.430 | מאמתים |  |
| 6:32.431 | מציאות |  |
| 6:32.432 | אל |  |
| 6:32.433 | אחד |  |
| 6:32.434 | בעולם |  |
| 6:32.435 | היפך |  |
| 6:32.436 | דברי |  |
| 6:32.437 | רז"ל | . |
| 6:32.438 | והכתובים |  |
| 6:32.439 | מוכיחים |  |
| 6:32.440 | כדעת |  |
| 6:32.441 | חכמים |  |
| 6:32.442 | ז"ל | . |
| 6:32.443 | שנ | ' |
| 6:32.444 | כי |  |
| 6:32.445 | הוא |  |
| 6:32.446 | האלהים |  |
| 6:32.447 | בשמים |  |
| 6:32.448 | ממעל |  |
| 6:32.449 | ועל |  |
| 6:32.450 | הארץ |  |
| 6:32.451 | מתחת |  |
| 6:32.452 | אין |  |
| 6:32.453 | עוד |  |
| 6:32.454 | מלבדו | . |
| 6:32.455 | אתה |  |
| 6:32.456 | הוא |  |
| 6:32.457 | ה | ' |
| 6:32.458 | לבדך | . |
| 6:32.459 | וכמה |  |
| 6:32.460 | כתובים |  |
| 6:32.461 | מכריזים |  |
| 6:32.462 | שאין |  |
| 6:32.463 | עוד |  |
| 6:32.464 | אלוה |  |
| 6:32.465 | אחר |  |
| 6:32.466 | זולתי |  |
| 6:32.467 | ה | ' |
| 6:32.468 | אלהינו | . |
| 6:32.469 | ורבותינו |  |
| 6:32.470 | מתקני |  |
| 6:32.471 | התפלות |  |
| 6:32.472 | התקינו |  |
| 6:32.473 | לומר |  |
| 6:32.474 | בכל |  |
| 6:32.475 | יום |  |
| 6:32.476 | אנו |  |
| 6:32.477 | יודעים |  |
| 6:32.478 | ומעידים |  |
| 6:32.479 | ומגידים | , |
| 6:32.480 | שאין |  |
| 6:32.481 | אלוה |  |
| 6:32.482 | זולתך |  |
| 6:32.483 | וכו | ' |
| 6:32.484 | והרבה |  |
| 6:32.485 | לשונות |  |
| 6:32.486 | בתפלה |  |
| 6:32.487 | מורים |  |
| 6:32.488 | שאין |  |
| 6:32.489 | מציאות |  |
| 6:32.490 | אל |  |
| 6:32.491 | אחר |  |
| 6:32.492 | בעולם |  |
| 6:32.493 | כלל |  |
| 6:32.494 | כלל | . |
| 6:32.495 | ורק |  |
| 6:32.496 | לפי |  |
| 6:32.497 | מחשבת |  |
| 6:32.498 | ובדיהן | . |
| 6:32.499 | אמרה |  |
| 6:32.500 | כי |  |
| 6:32.501 | תורה |  |
| 6:32.502 | כי |  |
| 6:32.503 | לא |  |
| 6:32.504 | תשתחוה |  |
| 6:32.505 | לאל |  |
| 6:32.506 | אחר | . |
| 6:32.507 | וכן |  |
| 6:32.508 | לא |  |
| 6:32.509 | יהיה |  |
| 6:32.510 | לך |  |
| 6:32.511 | אלהים |  |
| 6:32.512 | אחרים |  |
| 6:32.513 | על |  |
| 6:32.514 | פני | . |
| 6:32.515 | הכל |  |
| 6:32.516 | לפי |  |
| 6:32.517 | מחשבת |  |
| 6:32.518 | עובידהן |  |
| 6:32.519 | שטועין |  |
| 6:32.520 | וחושבים |  |
| 6:32.521 | אותם |  |
| 6:32.522 | אלוהות | . |
| 6:32.523 | לא |  |
| 6:32.524 | שהשי"ת |  |
| 6:32.525 | המציא |  |
| 6:32.526 | שני |  |
| 6:32.527 | מיני |  |
| 6:32.528 | אלוהות |  |
| 6:32.529 | אלה |  |
| 6:32.530 | קדושים |  |
| 6:32.531 | ואלה |  |
| 6:32.532 | טמאים |  |
| 6:32.533 | ח"ו | . |
| 6:32.534 | כי |  |
| 6:32.535 | אין |  |
| 6:32.536 | אלוה |  |
| 6:32.537 | זולתו | . |
| 6:32.538 | וכמה |  |
| 6:32.539 | פסורים |  |
| 6:32.540 | בתורה |  |
| 6:32.541 | ובנביאים |  |
| 6:32.542 | ובכתובים |  |
| 6:32.543 | מוכיחים |  |
| 6:32.544 | כן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

