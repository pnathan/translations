# Translation pack: milhamot-hashem perek-4-15 (Milchamot ha-Shem 4:24–4:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-15`.

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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.

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


## Milchamot ha-Shem 4:24

Hebrew (Wikisource): לג) כל משכיל מבין המעיין מעט בספרי הזהר והמקובלים, ידע ויכיר כי הספירות והפרצופים הנזכרים בדבריהם קדמו לבריאת שמים וארץ וכל צבאם כנראה מדברי מהרח"ו בספר עץ חיים שער הכללים ושער עיגולים ויושר. ומבואר עוד באר היטב שער שבירת הכלים פ"ג וז"ל ונבאר סדר יציאת מלכים ונתחיל מן הראשון שהוא הדעת אשר זה יצא ראשונה. וכאשר לא היה יכול הכלי לסבול כנזכר לעיל נשבר וירד למטה בעולם הבריאה. רצוני לומר במקום שיהיה עתיד להיות עולם הבריאה אחר כך כי הרי עדיין לא היה נברא עולם הבריאה, ונפל הכלי הזה במקום שעתיד הדעת דבריאה להיות עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:24.1 | לג | ) |
| 4:24.2 | כל |  |
| 4:24.3 | משכיל |  |
| 4:24.4 | מבין |  |
| 4:24.5 | המעיין |  |
| 4:24.6 | מעט |  |
| 4:24.7 | בספרי |  |
| 4:24.8 | הזהר |  |
| 4:24.9 | והמקובלים | , |
| 4:24.10 | ידע |  |
| 4:24.11 | ויכיר |  |
| 4:24.12 | כי |  |
| 4:24.13 | הספירות |  |
| 4:24.14 | והפרצופים |  |
| 4:24.15 | הנזכרים |  |
| 4:24.16 | בדבריהם |  |
| 4:24.17 | קדמו |  |
| 4:24.18 | לבריאת |  |
| 4:24.19 | שמים |  |
| 4:24.20 | וארץ |  |
| 4:24.21 | וכל |  |
| 4:24.22 | צבאם |  |
| 4:24.23 | כנראה |  |
| 4:24.24 | מדברי |  |
| 4:24.25 | מהרח"ו |  |
| 4:24.26 | בספר |  |
| 4:24.27 | עץ |  |
| 4:24.28 | חיים |  |
| 4:24.29 | שער |  |
| 4:24.30 | הכללים |  |
| 4:24.31 | ושער |  |
| 4:24.32 | עיגולים |  |
| 4:24.33 | ויושר | . |
| 4:24.34 | ומבואר |  |
| 4:24.35 | עוד |  |
| 4:24.36 | באר |  |
| 4:24.37 | היטב |  |
| 4:24.38 | שער |  |
| 4:24.39 | שבירת |  |
| 4:24.40 | הכלים |  |
| 4:24.41 | פ"ג |  |
| 4:24.42 | וז"ל |  |
| 4:24.43 | ונבאר |  |
| 4:24.44 | סדר |  |
| 4:24.45 | יציאת |  |
| 4:24.46 | מלכים |  |
| 4:24.47 | ונתחיל |  |
| 4:24.48 | מן |  |
| 4:24.49 | הראשון |  |
| 4:24.50 | שהוא |  |
| 4:24.51 | הדעת |  |
| 4:24.52 | אשר |  |
| 4:24.53 | זה |  |
| 4:24.54 | יצא |  |
| 4:24.55 | ראשונה | . |
| 4:24.56 | וכאשר |  |
| 4:24.57 | לא |  |
| 4:24.58 | היה |  |
| 4:24.59 | יכול |  |
| 4:24.60 | הכלי |  |
| 4:24.61 | לסבול |  |
| 4:24.62 | כנזכר |  |
| 4:24.63 | לעיל |  |
| 4:24.64 | נשבר |  |
| 4:24.65 | וירד |  |
| 4:24.66 | למטה |  |
| 4:24.67 | בעולם |  |
| 4:24.68 | הבריאה | . |
| 4:24.69 | רצוני |  |
| 4:24.70 | לומר |  |
| 4:24.71 | במקום |  |
| 4:24.72 | שיהיה |  |
| 4:24.73 | עתיד |  |
| 4:24.74 | להיות |  |
| 4:24.75 | עולם |  |
| 4:24.76 | הבריאה |  |
| 4:24.77 | אחר |  |
| 4:24.78 | כך |  |
| 4:24.79 | כי |  |
| 4:24.80 | הרי |  |
| 4:24.81 | עדיין |  |
| 4:24.82 | לא |  |
| 4:24.83 | היה |  |
| 4:24.84 | נברא |  |
| 4:24.85 | עולם |  |
| 4:24.86 | הבריאה | , |
| 4:24.87 | ונפל |  |
| 4:24.88 | הכלי |  |
| 4:24.89 | הזה |  |
| 4:24.90 | במקום |  |
| 4:24.91 | שעתיד |  |
| 4:24.92 | הדעת |  |
| 4:24.93 | דבריאה |  |
| 4:24.94 | להיות |  |
| 4:24.95 | עכ"ל | . |


## Milchamot ha-Shem 4:25

Hebrew (Wikisource): והיוצא לנו מזה כי רבותינו המקובלים החדשים שגגו לשאול ולמצוא פתרונים במה שלמעלה ומה למטה מה לפנים מה לאחור. דבר שקבלו חז"ל בעלי התלמוד בבלי וירושלמי ומדרש רבה ותנחומא וזולתם ממדרשי רז"ל לאסור. כדגר' בחגיגה פרק אין דורשין. כי שאל נא לימים ראשונים. יכול ישאל אדם קדם שנברא העולם ת"ל למן היום אשר ברא אלהים אדם על הארץ. יכול לא ישאל אדם מששת ימי בראשית ת"ל לימים ראשונים (פי' רש"י מיום ראשון) יכול ישאל אדם מה למעלה מה למטה ומה לפנים ומה לאחור תלמוד לומר ולמקצה השמים ועד קצה השמים. מקצה השמים ועוד קצה השמים אתה שואל. ואי אתה שואל מה למעלה ומה למטה ומה לפנים ומה לאחור עכ"ל הגמרא. ועיין במהרש"א בחדושי אגדות במעשה דר' אלעזר שאמר לר' יוחנן תא ואגמרך מעשה מרכבה ואמר ליה אכתי לא קשאי וכד קש אמר ליה ר' אסי תא ואגמרך מעשה מרכבה. אמר ליה אי זכאי הוה גמירנא מר' יוחנן רבך ונמנע ולא רצה ללמוד. לא בילדותו ולא בזקנותו. ושם הרגיש הר' מהרש"א שדברי הקבלה החדשה היא מדברת למעלה ממעשה מרכבה וראוי יותר להסתירה ולא לדרוש בה ברבים ע"ש. ואני אומר דמדינא אסור לדרוש בה כלל לא ביחיד ולא ברבים. שהרי במעשה מרכבה לא התירו חכמים לדרוש בהם ולחכם ומבין ראשי פרקים אלא עד החשמל. ואיכא דאמרי עד וארא ולא יותר דמשם ולמעלה אסור לדבר כלל, כפי קבלת חכמים ואזהרתם. והעובר על דברי חכמים חייב מיתה. ומכל זה למדנו ששגגה גדולה שגגו רבותינו בעלי הקבלה החדשה והגיסו דעתם ומלאם לבם לדרוש במה שלא הותר והורשה להם כמו שיתבאר לקמן. ועלתה בלבם מחשבה זרה שבהתקרב זמן הגאולה תבטל הלכה זו ויהיה מותר לדרוש במה שלא הותר להם מוקדם. כמו שמבואר בפסק שעשה ר' יצחק דלמאש המודפס בתחלת ס' הזהר. וטעות הוא בידו לפסוק הלכה מדברי הזהר היפך המשנה והתלמוד והפוסקים ואין לסמוך עליו כלל. ולא שמו המקובלים החדשים אל לבם מאמר רז"ל שכל הלכות של תורה שבעל פה לא יבטלו כמו שכ' הרמב"ם ז"ל בסוף הלכות מגלה. ובעברם על דברי חז"ל אירע להם בשגגת זו להאמין רבוי באלהות ולהאמין במציאות אלהים אחרים, ולשתף שם שמים ודבר אחר שעליו אמר ר' שמעון כל המשתף שם שמים ודבר אחר נעקר מן העולם. ובשלהי פסחים גרסינן מאי למכסה עתיק, זה המכסה דברים שכיסה עתיק יומין, ומאי נינהו סתרי תורה (פ' רש"י מעשה מרכבה ומע' בראשית) ואיכא דאמרי זה המגלה דברים שכיסה עתיק יומין, ומאי נינהו טעמי תורה ע"כ, הרי לך שהמגלה מעשה מרכבה לא יזכה לטוב הצפון לצדיקים, כ"ש וקל וחומר דברים שהם למעלה ממעשה מרכבה שאמור לדבר בהם כלל, ולא ניתן רשות לגלות אלא מטעמי תורה בלבד, אבל סתרי תורה לא ניתן רשות לגלותן, וכל שכן לדרשן ברבים בהדיא בבית המדרש אפי' לתלמידי חכמים, כל שכן בין ההמון אשר לא ידעו ולא יבינו, כי יביאם לידי מינות באמונת אלוהות רבים כאשר ראו עינינו, וחלילה לחשוב על ר' שמעון ב"י התנא שיעבור על איסור זה לדבר במה שלמעלה ממעשה מרכבה ובמה שהיה לפנים קודם בריאת העולם, וגם לא לשום תנא או אמורא לעשות כן ולומר שנגלה הקב"ה והתיר לו את האסור לדרוש בו, כי זו נביאות שקר היא (ומיתתו בחנק כמו שכתב הרמב"ם בהקדמת סדר זרעים) שאין נביא רשאי לחדש דבר מעתה כי לא בשמים היא.

| id | word | punctuation after |
|---|---|---|
| 4:25.1 | והיוצא |  |
| 4:25.2 | לנו |  |
| 4:25.3 | מזה |  |
| 4:25.4 | כי |  |
| 4:25.5 | רבותינו |  |
| 4:25.6 | המקובלים |  |
| 4:25.7 | החדשים |  |
| 4:25.8 | שגגו |  |
| 4:25.9 | לשאול |  |
| 4:25.10 | ולמצוא |  |
| 4:25.11 | פתרונים |  |
| 4:25.12 | במה |  |
| 4:25.13 | שלמעלה |  |
| 4:25.14 | ומה |  |
| 4:25.15 | למטה |  |
| 4:25.16 | מה |  |
| 4:25.17 | לפנים |  |
| 4:25.18 | מה |  |
| 4:25.19 | לאחור | . |
| 4:25.20 | דבר |  |
| 4:25.21 | שקבלו |  |
| 4:25.22 | חז"ל |  |
| 4:25.23 | בעלי |  |
| 4:25.24 | התלמוד |  |
| 4:25.25 | בבלי |  |
| 4:25.26 | וירושלמי |  |
| 4:25.27 | ומדרש |  |
| 4:25.28 | רבה |  |
| 4:25.29 | ותנחומא |  |
| 4:25.30 | וזולתם |  |
| 4:25.31 | ממדרשי |  |
| 4:25.32 | רז"ל |  |
| 4:25.33 | לאסור | . |
| 4:25.34 | כדגר | ' |
| 4:25.35 | בחגיגה |  |
| 4:25.36 | פרק |  |
| 4:25.37 | אין |  |
| 4:25.38 | דורשין | . |
| 4:25.39 | כי |  |
| 4:25.40 | שאל |  |
| 4:25.41 | נא |  |
| 4:25.42 | לימים |  |
| 4:25.43 | ראשונים | . |
| 4:25.44 | יכול |  |
| 4:25.45 | ישאל |  |
| 4:25.46 | אדם |  |
| 4:25.47 | קדם |  |
| 4:25.48 | שנברא |  |
| 4:25.49 | העולם |  |
| 4:25.50 | ת"ל |  |
| 4:25.51 | למן |  |
| 4:25.52 | היום |  |
| 4:25.53 | אשר |  |
| 4:25.54 | ברא |  |
| 4:25.55 | אלהים |  |
| 4:25.56 | אדם |  |
| 4:25.57 | על |  |
| 4:25.58 | הארץ | . |
| 4:25.59 | יכול |  |
| 4:25.60 | לא |  |
| 4:25.61 | ישאל |  |
| 4:25.62 | אדם |  |
| 4:25.63 | מששת |  |
| 4:25.64 | ימי |  |
| 4:25.65 | בראשית |  |
| 4:25.66 | ת"ל |  |
| 4:25.67 | לימים |  |
| 4:25.68 | ראשונים |  |
| 4:25.69 | פי | ' |
| 4:25.70 | רש"י |  |
| 4:25.71 | מיום |  |
| 4:25.72 | ראשון | ) |
| 4:25.73 | יכול |  |
| 4:25.74 | ישאל |  |
| 4:25.75 | אדם |  |
| 4:25.76 | מה |  |
| 4:25.77 | למעלה |  |
| 4:25.78 | מה |  |
| 4:25.79 | למטה |  |
| 4:25.80 | ומה |  |
| 4:25.81 | לפנים |  |
| 4:25.82 | ומה |  |
| 4:25.83 | לאחור |  |
| 4:25.84 | תלמוד |  |
| 4:25.85 | לומר |  |
| 4:25.86 | ולמקצה |  |
| 4:25.87 | השמים |  |
| 4:25.88 | ועד |  |
| 4:25.89 | קצה |  |
| 4:25.90 | השמים | . |
| 4:25.91 | מקצה |  |
| 4:25.92 | השמים |  |
| 4:25.93 | ועוד |  |
| 4:25.94 | קצה |  |
| 4:25.95 | השמים |  |
| 4:25.96 | אתה |  |
| 4:25.97 | שואל | . |
| 4:25.98 | ואי |  |
| 4:25.99 | אתה |  |
| 4:25.100 | שואל |  |
| 4:25.101 | מה |  |
| 4:25.102 | למעלה |  |
| 4:25.103 | ומה |  |
| 4:25.104 | למטה |  |
| 4:25.105 | ומה |  |
| 4:25.106 | לפנים |  |
| 4:25.107 | ומה |  |
| 4:25.108 | לאחור |  |
| 4:25.109 | עכ"ל |  |
| 4:25.110 | הגמרא | . |
| 4:25.111 | ועיין |  |
| 4:25.112 | במהרש"א |  |
| 4:25.113 | בחדושי |  |
| 4:25.114 | אגדות |  |
| 4:25.115 | במעשה |  |
| 4:25.116 | דר | ' |
| 4:25.117 | אלעזר |  |
| 4:25.118 | שאמר |  |
| 4:25.119 | לר | ' |
| 4:25.120 | יוחנן |  |
| 4:25.121 | תא |  |
| 4:25.122 | ואגמרך |  |
| 4:25.123 | מעשה |  |
| 4:25.124 | מרכבה |  |
| 4:25.125 | ואמר |  |
| 4:25.126 | ליה |  |
| 4:25.127 | אכתי |  |
| 4:25.128 | לא |  |
| 4:25.129 | קשאי |  |
| 4:25.130 | וכד |  |
| 4:25.131 | קש |  |
| 4:25.132 | אמר |  |
| 4:25.133 | ליה |  |
| 4:25.134 | ר | ' |
| 4:25.135 | אסי |  |
| 4:25.136 | תא |  |
| 4:25.137 | ואגמרך |  |
| 4:25.138 | מעשה |  |
| 4:25.139 | מרכבה | . |
| 4:25.140 | אמר |  |
| 4:25.141 | ליה |  |
| 4:25.142 | אי |  |
| 4:25.143 | זכאי |  |
| 4:25.144 | הוה |  |
| 4:25.145 | גמירנא |  |
| 4:25.146 | מר | ' |
| 4:25.147 | יוחנן |  |
| 4:25.148 | רבך |  |
| 4:25.149 | ונמנע |  |
| 4:25.150 | ולא |  |
| 4:25.151 | רצה |  |
| 4:25.152 | ללמוד | . |
| 4:25.153 | לא |  |
| 4:25.154 | בילדותו |  |
| 4:25.155 | ולא |  |
| 4:25.156 | בזקנותו | . |
| 4:25.157 | ושם |  |
| 4:25.158 | הרגיש |  |
| 4:25.159 | הר | ' |
| 4:25.160 | מהרש"א |  |
| 4:25.161 | שדברי |  |
| 4:25.162 | הקבלה |  |
| 4:25.163 | החדשה |  |
| 4:25.164 | היא |  |
| 4:25.165 | מדברת |  |
| 4:25.166 | למעלה |  |
| 4:25.167 | ממעשה |  |
| 4:25.168 | מרכבה |  |
| 4:25.169 | וראוי |  |
| 4:25.170 | יותר |  |
| 4:25.171 | להסתירה |  |
| 4:25.172 | ולא |  |
| 4:25.173 | לדרוש |  |
| 4:25.174 | בה |  |
| 4:25.175 | ברבים |  |
| 4:25.176 | ע"ש | . |
| 4:25.177 | ואני |  |
| 4:25.178 | אומר |  |
| 4:25.179 | דמדינא |  |
| 4:25.180 | אסור |  |
| 4:25.181 | לדרוש |  |
| 4:25.182 | בה |  |
| 4:25.183 | כלל |  |
| 4:25.184 | לא |  |
| 4:25.185 | ביחיד |  |
| 4:25.186 | ולא |  |
| 4:25.187 | ברבים | . |
| 4:25.188 | שהרי |  |
| 4:25.189 | במעשה |  |
| 4:25.190 | מרכבה |  |
| 4:25.191 | לא |  |
| 4:25.192 | התירו |  |
| 4:25.193 | חכמים |  |
| 4:25.194 | לדרוש |  |
| 4:25.195 | בהם |  |
| 4:25.196 | ולחכם |  |
| 4:25.197 | ומבין |  |
| 4:25.198 | ראשי |  |
| 4:25.199 | פרקים |  |
| 4:25.200 | אלא |  |
| 4:25.201 | עד |  |
| 4:25.202 | החשמל | . |
| 4:25.203 | ואיכא |  |
| 4:25.204 | דאמרי |  |
| 4:25.205 | עד |  |
| 4:25.206 | וארא |  |
| 4:25.207 | ולא |  |
| 4:25.208 | יותר |  |
| 4:25.209 | דמשם |  |
| 4:25.210 | ולמעלה |  |
| 4:25.211 | אסור |  |
| 4:25.212 | לדבר |  |
| 4:25.213 | כלל | , |
| 4:25.214 | כפי |  |
| 4:25.215 | קבלת |  |
| 4:25.216 | חכמים |  |
| 4:25.217 | ואזהרתם | . |
| 4:25.218 | והעובר |  |
| 4:25.219 | על |  |
| 4:25.220 | דברי |  |
| 4:25.221 | חכמים |  |
| 4:25.222 | חייב |  |
| 4:25.223 | מיתה | . |
| 4:25.224 | ומכל |  |
| 4:25.225 | זה |  |
| 4:25.226 | למדנו |  |
| 4:25.227 | ששגגה |  |
| 4:25.228 | גדולה |  |
| 4:25.229 | שגגו |  |
| 4:25.230 | רבותינו |  |
| 4:25.231 | בעלי |  |
| 4:25.232 | הקבלה |  |
| 4:25.233 | החדשה |  |
| 4:25.234 | והגיסו |  |
| 4:25.235 | דעתם |  |
| 4:25.236 | ומלאם |  |
| 4:25.237 | לבם |  |
| 4:25.238 | לדרוש |  |
| 4:25.239 | במה |  |
| 4:25.240 | שלא |  |
| 4:25.241 | הותר |  |
| 4:25.242 | והורשה |  |
| 4:25.243 | להם |  |
| 4:25.244 | כמו |  |
| 4:25.245 | שיתבאר |  |
| 4:25.246 | לקמן | . |
| 4:25.247 | ועלתה |  |
| 4:25.248 | בלבם |  |
| 4:25.249 | מחשבה |  |
| 4:25.250 | זרה |  |
| 4:25.251 | שבהתקרב |  |
| 4:25.252 | זמן |  |
| 4:25.253 | הגאולה |  |
| 4:25.254 | תבטל |  |
| 4:25.255 | הלכה |  |
| 4:25.256 | זו |  |
| 4:25.257 | ויהיה |  |
| 4:25.258 | מותר |  |
| 4:25.259 | לדרוש |  |
| 4:25.260 | במה |  |
| 4:25.261 | שלא |  |
| 4:25.262 | הותר |  |
| 4:25.263 | להם |  |
| 4:25.264 | מוקדם | . |
| 4:25.265 | כמו |  |
| 4:25.266 | שמבואר |  |
| 4:25.267 | בפסק |  |
| 4:25.268 | שעשה |  |
| 4:25.269 | ר | ' |
| 4:25.270 | יצחק |  |
| 4:25.271 | דלמאש |  |
| 4:25.272 | המודפס |  |
| 4:25.273 | בתחלת |  |
| 4:25.274 | ס | ' |
| 4:25.275 | הזהר | . |
| 4:25.276 | וטעות |  |
| 4:25.277 | הוא |  |
| 4:25.278 | בידו |  |
| 4:25.279 | לפסוק |  |
| 4:25.280 | הלכה |  |
| 4:25.281 | מדברי |  |
| 4:25.282 | הזהר |  |
| 4:25.283 | היפך |  |
| 4:25.284 | המשנה |  |
| 4:25.285 | והתלמוד |  |
| 4:25.286 | והפוסקים |  |
| 4:25.287 | ואין |  |
| 4:25.288 | לסמוך |  |
| 4:25.289 | עליו |  |
| 4:25.290 | כלל | . |
| 4:25.291 | ולא |  |
| 4:25.292 | שמו |  |
| 4:25.293 | המקובלים |  |
| 4:25.294 | החדשים |  |
| 4:25.295 | אל |  |
| 4:25.296 | לבם |  |
| 4:25.297 | מאמר |  |
| 4:25.298 | רז"ל |  |
| 4:25.299 | שכל |  |
| 4:25.300 | הלכות |  |
| 4:25.301 | של |  |
| 4:25.302 | תורה |  |
| 4:25.303 | שבעל |  |
| 4:25.304 | פה |  |
| 4:25.305 | לא |  |
| 4:25.306 | יבטלו |  |
| 4:25.307 | כמו |  |
| 4:25.308 | שכ | ' |
| 4:25.309 | הרמב"ם |  |
| 4:25.310 | ז"ל |  |
| 4:25.311 | בסוף |  |
| 4:25.312 | הלכות |  |
| 4:25.313 | מגלה | . |
| 4:25.314 | ובעברם |  |
| 4:25.315 | על |  |
| 4:25.316 | דברי |  |
| 4:25.317 | חז"ל |  |
| 4:25.318 | אירע |  |
| 4:25.319 | להם |  |
| 4:25.320 | בשגגת |  |
| 4:25.321 | זו |  |
| 4:25.322 | להאמין |  |
| 4:25.323 | רבוי |  |
| 4:25.324 | באלהות |  |
| 4:25.325 | ולהאמין |  |
| 4:25.326 | במציאות |  |
| 4:25.327 | אלהים |  |
| 4:25.328 | אחרים | , |
| 4:25.329 | ולשתף |  |
| 4:25.330 | שם |  |
| 4:25.331 | שמים |  |
| 4:25.332 | ודבר |  |
| 4:25.333 | אחר |  |
| 4:25.334 | שעליו |  |
| 4:25.335 | אמר |  |
| 4:25.336 | ר | ' |
| 4:25.337 | שמעון |  |
| 4:25.338 | כל |  |
| 4:25.339 | המשתף |  |
| 4:25.340 | שם |  |
| 4:25.341 | שמים |  |
| 4:25.342 | ודבר |  |
| 4:25.343 | אחר |  |
| 4:25.344 | נעקר |  |
| 4:25.345 | מן |  |
| 4:25.346 | העולם | . |
| 4:25.347 | ובשלהי |  |
| 4:25.348 | פסחים |  |
| 4:25.349 | גרסינן |  |
| 4:25.350 | מאי |  |
| 4:25.351 | למכסה |  |
| 4:25.352 | עתיק | , |
| 4:25.353 | זה |  |
| 4:25.354 | המכסה |  |
| 4:25.355 | דברים |  |
| 4:25.356 | שכיסה |  |
| 4:25.357 | עתיק |  |
| 4:25.358 | יומין | , |
| 4:25.359 | ומאי |  |
| 4:25.360 | נינהו |  |
| 4:25.361 | סתרי |  |
| 4:25.362 | תורה |  |
| 4:25.363 | פ | ' |
| 4:25.364 | רש"י |  |
| 4:25.365 | מעשה |  |
| 4:25.366 | מרכבה |  |
| 4:25.367 | ומע | ' |
| 4:25.368 | בראשית | ) |
| 4:25.369 | ואיכא |  |
| 4:25.370 | דאמרי |  |
| 4:25.371 | זה |  |
| 4:25.372 | המגלה |  |
| 4:25.373 | דברים |  |
| 4:25.374 | שכיסה |  |
| 4:25.375 | עתיק |  |
| 4:25.376 | יומין | , |
| 4:25.377 | ומאי |  |
| 4:25.378 | נינהו |  |
| 4:25.379 | טעמי |  |
| 4:25.380 | תורה |  |
| 4:25.381 | ע"כ | , |
| 4:25.382 | הרי |  |
| 4:25.383 | לך |  |
| 4:25.384 | שהמגלה |  |
| 4:25.385 | מעשה |  |
| 4:25.386 | מרכבה |  |
| 4:25.387 | לא |  |
| 4:25.388 | יזכה |  |
| 4:25.389 | לטוב |  |
| 4:25.390 | הצפון |  |
| 4:25.391 | לצדיקים | , |
| 4:25.392 | כ"ש |  |
| 4:25.393 | וקל |  |
| 4:25.394 | וחומר |  |
| 4:25.395 | דברים |  |
| 4:25.396 | שהם |  |
| 4:25.397 | למעלה |  |
| 4:25.398 | ממעשה |  |
| 4:25.399 | מרכבה |  |
| 4:25.400 | שאמור |  |
| 4:25.401 | לדבר |  |
| 4:25.402 | בהם |  |
| 4:25.403 | כלל | , |
| 4:25.404 | ולא |  |
| 4:25.405 | ניתן |  |
| 4:25.406 | רשות |  |
| 4:25.407 | לגלות |  |
| 4:25.408 | אלא |  |
| 4:25.409 | מטעמי |  |
| 4:25.410 | תורה |  |
| 4:25.411 | בלבד | , |
| 4:25.412 | אבל |  |
| 4:25.413 | סתרי |  |
| 4:25.414 | תורה |  |
| 4:25.415 | לא |  |
| 4:25.416 | ניתן |  |
| 4:25.417 | רשות |  |
| 4:25.418 | לגלותן | , |
| 4:25.419 | וכל |  |
| 4:25.420 | שכן |  |
| 4:25.421 | לדרשן |  |
| 4:25.422 | ברבים |  |
| 4:25.423 | בהדיא |  |
| 4:25.424 | בבית |  |
| 4:25.425 | המדרש |  |
| 4:25.426 | אפי | ' |
| 4:25.427 | לתלמידי |  |
| 4:25.428 | חכמים | , |
| 4:25.429 | כל |  |
| 4:25.430 | שכן |  |
| 4:25.431 | בין |  |
| 4:25.432 | ההמון |  |
| 4:25.433 | אשר |  |
| 4:25.434 | לא |  |
| 4:25.435 | ידעו |  |
| 4:25.436 | ולא |  |
| 4:25.437 | יבינו | , |
| 4:25.438 | כי |  |
| 4:25.439 | יביאם |  |
| 4:25.440 | לידי |  |
| 4:25.441 | מינות |  |
| 4:25.442 | באמונת |  |
| 4:25.443 | אלוהות |  |
| 4:25.444 | רבים |  |
| 4:25.445 | כאשר |  |
| 4:25.446 | ראו |  |
| 4:25.447 | עינינו | , |
| 4:25.448 | וחלילה |  |
| 4:25.449 | לחשוב |  |
| 4:25.450 | על |  |
| 4:25.451 | ר | ' |
| 4:25.452 | שמעון |  |
| 4:25.453 | ב"י |  |
| 4:25.454 | התנא |  |
| 4:25.455 | שיעבור |  |
| 4:25.456 | על |  |
| 4:25.457 | איסור |  |
| 4:25.458 | זה |  |
| 4:25.459 | לדבר |  |
| 4:25.460 | במה |  |
| 4:25.461 | שלמעלה |  |
| 4:25.462 | ממעשה |  |
| 4:25.463 | מרכבה |  |
| 4:25.464 | ובמה |  |
| 4:25.465 | שהיה |  |
| 4:25.466 | לפנים |  |
| 4:25.467 | קודם |  |
| 4:25.468 | בריאת |  |
| 4:25.469 | העולם | , |
| 4:25.470 | וגם |  |
| 4:25.471 | לא |  |
| 4:25.472 | לשום |  |
| 4:25.473 | תנא |  |
| 4:25.474 | או |  |
| 4:25.475 | אמורא |  |
| 4:25.476 | לעשות |  |
| 4:25.477 | כן |  |
| 4:25.478 | ולומר |  |
| 4:25.479 | שנגלה |  |
| 4:25.480 | הקב"ה |  |
| 4:25.481 | והתיר |  |
| 4:25.482 | לו |  |
| 4:25.483 | את |  |
| 4:25.484 | האסור |  |
| 4:25.485 | לדרוש |  |
| 4:25.486 | בו | , |
| 4:25.487 | כי |  |
| 4:25.488 | זו |  |
| 4:25.489 | נביאות |  |
| 4:25.490 | שקר |  |
| 4:25.491 | היא |  |
| 4:25.492 | ומיתתו |  |
| 4:25.493 | בחנק |  |
| 4:25.494 | כמו |  |
| 4:25.495 | שכתב |  |
| 4:25.496 | הרמב"ם |  |
| 4:25.497 | בהקדמת |  |
| 4:25.498 | סדר |  |
| 4:25.499 | זרעים | ) |
| 4:25.500 | שאין |  |
| 4:25.501 | נביא |  |
| 4:25.502 | רשאי |  |
| 4:25.503 | לחדש |  |
| 4:25.504 | דבר |  |
| 4:25.505 | מעתה |  |
| 4:25.506 | כי |  |
| 4:25.507 | לא |  |
| 4:25.508 | בשמים |  |
| 4:25.509 | היא | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

