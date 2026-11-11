# Translation pack: milhamot-hashem perek-6-26 (Milchamot ha-Shem 6:39–6:40)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-26.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-26`.

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
- הוי"ה / הוי'ה / הויה (the Name as a name: Kabbalistic usage, and the
  author's own in 3:17) → "YHVH". Never "being". Examples: "ד' מילוי הויה ע"ב
  ס"ג מ"ה ב"ן" (2:33) → "the four fillings of YHVH, 72, 63, 45, 52" (in the
  note: these are the gematria names AB, SaG, MaH, BaN). "שהוא קורא לכל אחד
  מהם בשם הוי'ה ואדנות ואלהים" (8:24) → "who calls each of them by the names
  YHVH, Adonai and Elohim". אדנות as a name → "Adonai", with a note that the
  word is literally "lordship". עשר הויות (2:33) → "ten havayot", with a note.
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
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
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


## Milchamot ha-Shem 6:39

Hebrew (Wikisource): עוד כתב וז"ל ואל יעלה בדעתך כי תוכל להשיגו ית' מצד היותך דורש וחוקר בידיעת מהותו ועצם אמתתו. כי דבר זה איננו אפשרי. ואדרבה הוא מחרף ומגדף מי שחוקר בזה. שאין כח בכל מה שזולתו ית' להשיגו. וכל שכן כל מי שהוא בעל חומר. ועל זה הזהירו רבותינו ז"ל באמרם במופלא ממך בל תדרוש ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:39.1 | עוד |  |
| 6:39.2 | כתב |  |
| 6:39.3 | וז"ל |  |
| 6:39.4 | ואל |  |
| 6:39.5 | יעלה |  |
| 6:39.6 | בדעתך |  |
| 6:39.7 | כי |  |
| 6:39.8 | תוכל |  |
| 6:39.9 | להשיגו |  |
| 6:39.10 | ית | ' |
| 6:39.11 | מצד |  |
| 6:39.12 | היותך |  |
| 6:39.13 | דורש |  |
| 6:39.14 | וחוקר |  |
| 6:39.15 | בידיעת |  |
| 6:39.16 | מהותו |  |
| 6:39.17 | ועצם |  |
| 6:39.18 | אמתתו | . |
| 6:39.19 | כי |  |
| 6:39.20 | דבר |  |
| 6:39.21 | זה |  |
| 6:39.22 | איננו |  |
| 6:39.23 | אפשרי | . |
| 6:39.24 | ואדרבה |  |
| 6:39.25 | הוא |  |
| 6:39.26 | מחרף |  |
| 6:39.27 | ומגדף |  |
| 6:39.28 | מי |  |
| 6:39.29 | שחוקר |  |
| 6:39.30 | בזה | . |
| 6:39.31 | שאין |  |
| 6:39.32 | כח |  |
| 6:39.33 | בכל |  |
| 6:39.34 | מה |  |
| 6:39.35 | שזולתו |  |
| 6:39.36 | ית | ' |
| 6:39.37 | להשיגו | . |
| 6:39.38 | וכל |  |
| 6:39.39 | שכן |  |
| 6:39.40 | כל |  |
| 6:39.41 | מי |  |
| 6:39.42 | שהוא |  |
| 6:39.43 | בעל |  |
| 6:39.44 | חומר | . |
| 6:39.45 | ועל |  |
| 6:39.46 | זה |  |
| 6:39.47 | הזהירו |  |
| 6:39.48 | רבותינו |  |
| 6:39.49 | ז"ל |  |
| 6:39.50 | באמרם |  |
| 6:39.51 | במופלא |  |
| 6:39.52 | ממך |  |
| 6:39.53 | בל |  |
| 6:39.54 | תדרוש |  |
| 6:39.55 | ע"כ | . |


## Milchamot ha-Shem 6:40

Hebrew (Wikisource): עח) עיניך הרואות קורא נעים, איך רבו' המקובלים החדשים בהמשכם בעצימת עינים אחר המסית, ששגו לדבר במהותו יתברך לומר שהוא אור צח ומצוחצח, ונתנו לו ית' תבנית ותמונה, ושנויים רבים, יחסו לו באמרם כי בתחלה היה ממלא כל החלל הזה, ואחר כך צמצם את עצמו לצדדין סביב להמציא מקום וחלל לברואיו, ונעשה ככדור עגול וחלול באמצעיתו, ואחר כך נשתלשל ממנו קו וצנור דק, והיה יורד מעט ומתעגל ונעשו ספירות דעיגולים, ואחר כך נמשך ממנו דרך אותו צינור ונעשו גם ספירות דיושר בצורת אדם שלם בעל קומה ברמ"ח איברים ושס"ה גידים כמו שכתבתי לעיל בקיצור, והרוצה לעמוד על ביאורן יעיין בעץ חיים להרח"ו ושער ההקדמות וכסא אליהו ועז לאלהים ושאר ספרי קבלה. ולא הספיק להם דמות ותבנית וצורה ותמונה אחת (כאשר אמרו רז"ל מי לי בשמים. ועמך לא חפצתי בארץ ויובא לקמן סי' פ"ה) אלא פרצופים רבים בראו להם בדימיונם זה למעלה מזה וכל אחד נוטל רשות מאלוה שלמעלה ממנו. כי עשו את כל אחד מהם עלה וסבה לאלוה שלמטה ממנו. ועשו הסבה הראשונה אדם קדמון. ואמרו עליו שהוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין למעלה ממנו ליטול רשות ממנו. ואחריו עתיקא קדישא כלל זכר ונקבה בחד פרצוף. ואחריו אריך אנפין כולל ג"כ זכר ונקבה בפרצוף אחד. אלא שזה הזכר והנקבה פנים ואחור, וזה ימין ושמאל, כמבואר הכל בדבריהם. ואחריו אבא ואימא ג"כ זכר ונקבה בשני פרצופים נפרדים זה מזה, ואחר אבא ואימא זעיר ונוקביה שני פרצופים ג"כ נפרדים וגם אלו הפרצופים נחלקים עוד לשנים עשר פרצופים פרטיים, וכל אחד מן הפרצופים הנזכרים נקרא בשם הקב"ה ובשם הוי'ה ואדנות. ואלו הן השנים עשר פרצופים. עתיק ונוקביה, אריך אנפין ונוקביה, אבא ואימא עילאין, ישראל סבא ותבונה, זעיר ונוקביה, יעקב ולאה (עץ חיים שער התיקון פרק א'). גם נתנו שיעור וקצב לאלו הפרצופים שזה מגיע עד החזה של פרצוף למעלה ממנו, וזה מגיע עד הטבור וכיוצא באלו העניינים שהם ממקרי הגוף ומשיגיו, והנה שכחו באמרם שאדם קדמון הוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין לו ממי ליטול רשות, ואיה כל הפרצופים שבעולמות העליונים שהם אלפים ורבבות, וכל אחד כלול מאבי"ע? ולמה הוא אין צריך ליטול רשות ממה שלמעלה ממנו? ולמה לא הניחו רשות ושליטה לאדם קדמאה שהוא למעלה ממנו, ולא לאין סוף העולה על כולם? האם לא נמצאו מחרפים ומגדפים באמרם כי מהותו אור כמ"ש הר' הנז'? ולא באים לידי מינות בקראם לכל עלה מאלה העלות בשם הקב"ה והוי'ה ואלהים? וכיון שהם מאמינים באלוהות רבים ואומרים שהכל אחד, מה הפרש יש בין אמונתם לאמונת העכו"ם שהם שלשה והשלשה אחד? ורז"ל הרחיקו לאמונה זו כדגרסינן בע"ז והבאתיו לעיל סי' ס"ד, ולפי דבריהם למה לא הניחו ר' ישמעאל לבן דמא להתרפאת מיעקב איש כפר סכניא כיון שאומרים שהשלשה אחד? ואיין עוד בסי' ס"ד מה שכתבתי בש' ירושלמי שבת פרק ח' שרצים.

| id | word | punctuation after |
|---|---|---|
| 6:40.1 | עח | ) |
| 6:40.2 | עיניך |  |
| 6:40.3 | הרואות |  |
| 6:40.4 | קורא |  |
| 6:40.5 | נעים | , |
| 6:40.6 | איך |  |
| 6:40.7 | רבו | ' |
| 6:40.8 | המקובלים |  |
| 6:40.9 | החדשים |  |
| 6:40.10 | בהמשכם |  |
| 6:40.11 | בעצימת |  |
| 6:40.12 | עינים |  |
| 6:40.13 | אחר |  |
| 6:40.14 | המסית | , |
| 6:40.15 | ששגו |  |
| 6:40.16 | לדבר |  |
| 6:40.17 | במהותו |  |
| 6:40.18 | יתברך |  |
| 6:40.19 | לומר |  |
| 6:40.20 | שהוא |  |
| 6:40.21 | אור |  |
| 6:40.22 | צח |  |
| 6:40.23 | ומצוחצח | , |
| 6:40.24 | ונתנו |  |
| 6:40.25 | לו |  |
| 6:40.26 | ית | ' |
| 6:40.27 | תבנית |  |
| 6:40.28 | ותמונה | , |
| 6:40.29 | ושנויים |  |
| 6:40.30 | רבים | , |
| 6:40.31 | יחסו |  |
| 6:40.32 | לו |  |
| 6:40.33 | באמרם |  |
| 6:40.34 | כי |  |
| 6:40.35 | בתחלה |  |
| 6:40.36 | היה |  |
| 6:40.37 | ממלא |  |
| 6:40.38 | כל |  |
| 6:40.39 | החלל |  |
| 6:40.40 | הזה | , |
| 6:40.41 | ואחר |  |
| 6:40.42 | כך |  |
| 6:40.43 | צמצם |  |
| 6:40.44 | את |  |
| 6:40.45 | עצמו |  |
| 6:40.46 | לצדדין |  |
| 6:40.47 | סביב |  |
| 6:40.48 | להמציא |  |
| 6:40.49 | מקום |  |
| 6:40.50 | וחלל |  |
| 6:40.51 | לברואיו | , |
| 6:40.52 | ונעשה |  |
| 6:40.53 | ככדור |  |
| 6:40.54 | עגול |  |
| 6:40.55 | וחלול |  |
| 6:40.56 | באמצעיתו | , |
| 6:40.57 | ואחר |  |
| 6:40.58 | כך |  |
| 6:40.59 | נשתלשל |  |
| 6:40.60 | ממנו |  |
| 6:40.61 | קו |  |
| 6:40.62 | וצנור |  |
| 6:40.63 | דק | , |
| 6:40.64 | והיה |  |
| 6:40.65 | יורד |  |
| 6:40.66 | מעט |  |
| 6:40.67 | ומתעגל |  |
| 6:40.68 | ונעשו |  |
| 6:40.69 | ספירות |  |
| 6:40.70 | דעיגולים | , |
| 6:40.71 | ואחר |  |
| 6:40.72 | כך |  |
| 6:40.73 | נמשך |  |
| 6:40.74 | ממנו |  |
| 6:40.75 | דרך |  |
| 6:40.76 | אותו |  |
| 6:40.77 | צינור |  |
| 6:40.78 | ונעשו |  |
| 6:40.79 | גם |  |
| 6:40.80 | ספירות |  |
| 6:40.81 | דיושר |  |
| 6:40.82 | בצורת |  |
| 6:40.83 | אדם |  |
| 6:40.84 | שלם |  |
| 6:40.85 | בעל |  |
| 6:40.86 | קומה |  |
| 6:40.87 | ברמ"ח |  |
| 6:40.88 | איברים |  |
| 6:40.89 | ושס"ה |  |
| 6:40.90 | גידים |  |
| 6:40.91 | כמו |  |
| 6:40.92 | שכתבתי |  |
| 6:40.93 | לעיל |  |
| 6:40.94 | בקיצור | , |
| 6:40.95 | והרוצה |  |
| 6:40.96 | לעמוד |  |
| 6:40.97 | על |  |
| 6:40.98 | ביאורן |  |
| 6:40.99 | יעיין |  |
| 6:40.100 | בעץ |  |
| 6:40.101 | חיים |  |
| 6:40.102 | להרח"ו |  |
| 6:40.103 | ושער |  |
| 6:40.104 | ההקדמות |  |
| 6:40.105 | וכסא |  |
| 6:40.106 | אליהו |  |
| 6:40.107 | ועז |  |
| 6:40.108 | לאלהים |  |
| 6:40.109 | ושאר |  |
| 6:40.110 | ספרי |  |
| 6:40.111 | קבלה | . |
| 6:40.112 | ולא |  |
| 6:40.113 | הספיק |  |
| 6:40.114 | להם |  |
| 6:40.115 | דמות |  |
| 6:40.116 | ותבנית |  |
| 6:40.117 | וצורה |  |
| 6:40.118 | ותמונה |  |
| 6:40.119 | אחת |  |
| 6:40.120 | כאשר |  |
| 6:40.121 | אמרו |  |
| 6:40.122 | רז"ל |  |
| 6:40.123 | מי |  |
| 6:40.124 | לי |  |
| 6:40.125 | בשמים | . |
| 6:40.126 | ועמך |  |
| 6:40.127 | לא |  |
| 6:40.128 | חפצתי |  |
| 6:40.129 | בארץ |  |
| 6:40.130 | ויובא |  |
| 6:40.131 | לקמן |  |
| 6:40.132 | סי | ' |
| 6:40.133 | פ"ה | ) |
| 6:40.134 | אלא |  |
| 6:40.135 | פרצופים |  |
| 6:40.136 | רבים |  |
| 6:40.137 | בראו |  |
| 6:40.138 | להם |  |
| 6:40.139 | בדימיונם |  |
| 6:40.140 | זה |  |
| 6:40.141 | למעלה |  |
| 6:40.142 | מזה |  |
| 6:40.143 | וכל |  |
| 6:40.144 | אחד |  |
| 6:40.145 | נוטל |  |
| 6:40.146 | רשות |  |
| 6:40.147 | מאלוה |  |
| 6:40.148 | שלמעלה |  |
| 6:40.149 | ממנו | . |
| 6:40.150 | כי |  |
| 6:40.151 | עשו |  |
| 6:40.152 | את |  |
| 6:40.153 | כל |  |
| 6:40.154 | אחד |  |
| 6:40.155 | מהם |  |
| 6:40.156 | עלה |  |
| 6:40.157 | וסבה |  |
| 6:40.158 | לאלוה |  |
| 6:40.159 | שלמטה |  |
| 6:40.160 | ממנו | . |
| 6:40.161 | ועשו |  |
| 6:40.162 | הסבה |  |
| 6:40.163 | הראשונה |  |
| 6:40.164 | אדם |  |
| 6:40.165 | קדמון | . |
| 6:40.166 | ואמרו |  |
| 6:40.167 | עליו |  |
| 6:40.168 | שהוא |  |
| 6:40.169 | שאמר |  |
| 6:40.170 | ראו |  |
| 6:40.171 | עתה |  |
| 6:40.172 | כי |  |
| 6:40.173 | אני |  |
| 6:40.174 | אני |  |
| 6:40.175 | הוא |  |
| 6:40.176 | ואין |  |
| 6:40.177 | אלהים |  |
| 6:40.178 | עמדי | , |
| 6:40.179 | שאין |  |
| 6:40.180 | למעלה |  |
| 6:40.181 | ממנו |  |
| 6:40.182 | ליטול |  |
| 6:40.183 | רשות |  |
| 6:40.184 | ממנו | . |
| 6:40.185 | ואחריו |  |
| 6:40.186 | עתיקא |  |
| 6:40.187 | קדישא |  |
| 6:40.188 | כלל |  |
| 6:40.189 | זכר |  |
| 6:40.190 | ונקבה |  |
| 6:40.191 | בחד |  |
| 6:40.192 | פרצוף | . |
| 6:40.193 | ואחריו |  |
| 6:40.194 | אריך |  |
| 6:40.195 | אנפין |  |
| 6:40.196 | כולל |  |
| 6:40.197 | ג"כ |  |
| 6:40.198 | זכר |  |
| 6:40.199 | ונקבה |  |
| 6:40.200 | בפרצוף |  |
| 6:40.201 | אחד | . |
| 6:40.202 | אלא |  |
| 6:40.203 | שזה |  |
| 6:40.204 | הזכר |  |
| 6:40.205 | והנקבה |  |
| 6:40.206 | פנים |  |
| 6:40.207 | ואחור | , |
| 6:40.208 | וזה |  |
| 6:40.209 | ימין |  |
| 6:40.210 | ושמאל | , |
| 6:40.211 | כמבואר |  |
| 6:40.212 | הכל |  |
| 6:40.213 | בדבריהם | . |
| 6:40.214 | ואחריו |  |
| 6:40.215 | אבא |  |
| 6:40.216 | ואימא |  |
| 6:40.217 | ג"כ |  |
| 6:40.218 | זכר |  |
| 6:40.219 | ונקבה |  |
| 6:40.220 | בשני |  |
| 6:40.221 | פרצופים |  |
| 6:40.222 | נפרדים |  |
| 6:40.223 | זה |  |
| 6:40.224 | מזה | , |
| 6:40.225 | ואחר |  |
| 6:40.226 | אבא |  |
| 6:40.227 | ואימא |  |
| 6:40.228 | זעיר |  |
| 6:40.229 | ונוקביה |  |
| 6:40.230 | שני |  |
| 6:40.231 | פרצופים |  |
| 6:40.232 | ג"כ |  |
| 6:40.233 | נפרדים |  |
| 6:40.234 | וגם |  |
| 6:40.235 | אלו |  |
| 6:40.236 | הפרצופים |  |
| 6:40.237 | נחלקים |  |
| 6:40.238 | עוד |  |
| 6:40.239 | לשנים |  |
| 6:40.240 | עשר |  |
| 6:40.241 | פרצופים |  |
| 6:40.242 | פרטיים | , |
| 6:40.243 | וכל |  |
| 6:40.244 | אחד |  |
| 6:40.245 | מן |  |
| 6:40.246 | הפרצופים |  |
| 6:40.247 | הנזכרים |  |
| 6:40.248 | נקרא |  |
| 6:40.249 | בשם |  |
| 6:40.250 | הקב"ה |  |
| 6:40.251 | ובשם |  |
| 6:40.252 | הוי'ה |  |
| 6:40.253 | ואדנות | . |
| 6:40.254 | ואלו |  |
| 6:40.255 | הן |  |
| 6:40.256 | השנים |  |
| 6:40.257 | עשר |  |
| 6:40.258 | פרצופים | . |
| 6:40.259 | עתיק |  |
| 6:40.260 | ונוקביה | , |
| 6:40.261 | אריך |  |
| 6:40.262 | אנפין |  |
| 6:40.263 | ונוקביה | , |
| 6:40.264 | אבא |  |
| 6:40.265 | ואימא |  |
| 6:40.266 | עילאין | , |
| 6:40.267 | ישראל |  |
| 6:40.268 | סבא |  |
| 6:40.269 | ותבונה | , |
| 6:40.270 | זעיר |  |
| 6:40.271 | ונוקביה | , |
| 6:40.272 | יעקב |  |
| 6:40.273 | ולאה |  |
| 6:40.274 | עץ |  |
| 6:40.275 | חיים |  |
| 6:40.276 | שער |  |
| 6:40.277 | התיקון |  |
| 6:40.278 | פרק |  |
| 6:40.279 | א | '). |
| 6:40.280 | גם |  |
| 6:40.281 | נתנו |  |
| 6:40.282 | שיעור |  |
| 6:40.283 | וקצב |  |
| 6:40.284 | לאלו |  |
| 6:40.285 | הפרצופים |  |
| 6:40.286 | שזה |  |
| 6:40.287 | מגיע |  |
| 6:40.288 | עד |  |
| 6:40.289 | החזה |  |
| 6:40.290 | של |  |
| 6:40.291 | פרצוף |  |
| 6:40.292 | למעלה |  |
| 6:40.293 | ממנו | , |
| 6:40.294 | וזה |  |
| 6:40.295 | מגיע |  |
| 6:40.296 | עד |  |
| 6:40.297 | הטבור |  |
| 6:40.298 | וכיוצא |  |
| 6:40.299 | באלו |  |
| 6:40.300 | העניינים |  |
| 6:40.301 | שהם |  |
| 6:40.302 | ממקרי |  |
| 6:40.303 | הגוף |  |
| 6:40.304 | ומשיגיו | , |
| 6:40.305 | והנה |  |
| 6:40.306 | שכחו |  |
| 6:40.307 | באמרם |  |
| 6:40.308 | שאדם |  |
| 6:40.309 | קדמון |  |
| 6:40.310 | הוא |  |
| 6:40.311 | שאמר |  |
| 6:40.312 | ראו |  |
| 6:40.313 | עתה |  |
| 6:40.314 | כי |  |
| 6:40.315 | אני |  |
| 6:40.316 | אני |  |
| 6:40.317 | הוא |  |
| 6:40.318 | ואין |  |
| 6:40.319 | אלהים |  |
| 6:40.320 | עמדי | , |
| 6:40.321 | שאין |  |
| 6:40.322 | לו |  |
| 6:40.323 | ממי |  |
| 6:40.324 | ליטול |  |
| 6:40.325 | רשות | , |
| 6:40.326 | ואיה |  |
| 6:40.327 | כל |  |
| 6:40.328 | הפרצופים |  |
| 6:40.329 | שבעולמות |  |
| 6:40.330 | העליונים |  |
| 6:40.331 | שהם |  |
| 6:40.332 | אלפים |  |
| 6:40.333 | ורבבות | , |
| 6:40.334 | וכל |  |
| 6:40.335 | אחד |  |
| 6:40.336 | כלול |  |
| 6:40.337 | מאבי"ע | ? |
| 6:40.338 | ולמה |  |
| 6:40.339 | הוא |  |
| 6:40.340 | אין |  |
| 6:40.341 | צריך |  |
| 6:40.342 | ליטול |  |
| 6:40.343 | רשות |  |
| 6:40.344 | ממה |  |
| 6:40.345 | שלמעלה |  |
| 6:40.346 | ממנו | ? |
| 6:40.347 | ולמה |  |
| 6:40.348 | לא |  |
| 6:40.349 | הניחו |  |
| 6:40.350 | רשות |  |
| 6:40.351 | ושליטה |  |
| 6:40.352 | לאדם |  |
| 6:40.353 | קדמאה |  |
| 6:40.354 | שהוא |  |
| 6:40.355 | למעלה |  |
| 6:40.356 | ממנו | , |
| 6:40.357 | ולא |  |
| 6:40.358 | לאין |  |
| 6:40.359 | סוף |  |
| 6:40.360 | העולה |  |
| 6:40.361 | על |  |
| 6:40.362 | כולם | ? |
| 6:40.363 | האם |  |
| 6:40.364 | לא |  |
| 6:40.365 | נמצאו |  |
| 6:40.366 | מחרפים |  |
| 6:40.367 | ומגדפים |  |
| 6:40.368 | באמרם |  |
| 6:40.369 | כי |  |
| 6:40.370 | מהותו |  |
| 6:40.371 | אור |  |
| 6:40.372 | כמ"ש |  |
| 6:40.373 | הר | ' |
| 6:40.374 | הנז | '? |
| 6:40.375 | ולא |  |
| 6:40.376 | באים |  |
| 6:40.377 | לידי |  |
| 6:40.378 | מינות |  |
| 6:40.379 | בקראם |  |
| 6:40.380 | לכל |  |
| 6:40.381 | עלה |  |
| 6:40.382 | מאלה |  |
| 6:40.383 | העלות |  |
| 6:40.384 | בשם |  |
| 6:40.385 | הקב"ה |  |
| 6:40.386 | והוי'ה |  |
| 6:40.387 | ואלהים | ? |
| 6:40.388 | וכיון |  |
| 6:40.389 | שהם |  |
| 6:40.390 | מאמינים |  |
| 6:40.391 | באלוהות |  |
| 6:40.392 | רבים |  |
| 6:40.393 | ואומרים |  |
| 6:40.394 | שהכל |  |
| 6:40.395 | אחד | , |
| 6:40.396 | מה |  |
| 6:40.397 | הפרש |  |
| 6:40.398 | יש |  |
| 6:40.399 | בין |  |
| 6:40.400 | אמונתם |  |
| 6:40.401 | לאמונת |  |
| 6:40.402 | העכו"ם |  |
| 6:40.403 | שהם |  |
| 6:40.404 | שלשה |  |
| 6:40.405 | והשלשה |  |
| 6:40.406 | אחד | ? |
| 6:40.407 | ורז"ל |  |
| 6:40.408 | הרחיקו |  |
| 6:40.409 | לאמונה |  |
| 6:40.410 | זו |  |
| 6:40.411 | כדגרסינן |  |
| 6:40.412 | בע"ז |  |
| 6:40.413 | והבאתיו |  |
| 6:40.414 | לעיל |  |
| 6:40.415 | סי | ' |
| 6:40.416 | ס"ד | , |
| 6:40.417 | ולפי |  |
| 6:40.418 | דבריהם |  |
| 6:40.419 | למה |  |
| 6:40.420 | לא |  |
| 6:40.421 | הניחו |  |
| 6:40.422 | ר | ' |
| 6:40.423 | ישמעאל |  |
| 6:40.424 | לבן |  |
| 6:40.425 | דמא |  |
| 6:40.426 | להתרפאת |  |
| 6:40.427 | מיעקב |  |
| 6:40.428 | איש |  |
| 6:40.429 | כפר |  |
| 6:40.430 | סכניא |  |
| 6:40.431 | כיון |  |
| 6:40.432 | שאומרים |  |
| 6:40.433 | שהשלשה |  |
| 6:40.434 | אחד | ? |
| 6:40.435 | ואיין |  |
| 6:40.436 | עוד |  |
| 6:40.437 | בסי | ' |
| 6:40.438 | ס"ד |  |
| 6:40.439 | מה |  |
| 6:40.440 | שכתבתי |  |
| 6:40.441 | בש | ' |
| 6:40.442 | ירושלמי |  |
| 6:40.443 | שבת |  |
| 6:40.444 | פרק |  |
| 6:40.445 | ח | ' |
| 6:40.446 | שרצים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

