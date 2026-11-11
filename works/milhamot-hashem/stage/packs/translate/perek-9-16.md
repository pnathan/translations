# Translation pack: milhamot-hashem perek-9-16 (Milchamot ha-Shem 9:19–9:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-16`.

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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- והוא שלא: "provided that" (fixed).
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


## Milchamot ha-Shem 9:19

Hebrew (Wikisource): קלז) ובמדרש רבה פרשת קרח דרשו רז"ל ששים המה מלכות, אלו ס' מסכתות ושמנים פילגשים אלו שמנים בתי מדרשות שהיו בירושלם כנגד פתחיה, ועלמות אין מספר, משנה החיצונה, ופירש הרב עץ יוסף ומת"כ דקרי לבתי מדרשות פילגשים על שם שנושאין ונותנים בהם וחולקים בהם וחולקים זה עם זה כמו פילגשם (ומחבר הזהר כשנאתו למשנה ולתלמוד דימה שקלא וטריא דרבנן לכלביך דצוחון הב הב ושם אותם מיורשי גהינם האומרת הב הב) ופי' משנה החיצונה מה שלמדו היחידים חוץ לאלו בתי מדרשות, וביד יוסף פי' אלו הברייתות, ע"כ, וכן מצאתי בשיר השירים רבה וז"ל ששים המה מלכות, אלו ששים מסכתות של הלכות ושמנים פילגשים אלו שמנים פרשיות שבתורת כהנים. ועלמות אין מספר, אין קץ לתוספתות, אחת היא הן חולקין אלו עם אלו וכולהון דורשין מטעם אחד, ומהלכה אחת מגזרה שוה מקל וחומר ע"כ. ראה ידידי הקורא כמה מחבבים רז"ל את המשנה והתלמוד וקראו להן מלכות ולא שפחות, כמאמר שלמה כי מלכים ימלוכו, והמה בתרה של תורה כי לא נוכל לעשות דבר ממצות התורה כי אם על פי קבלת רז"ל בעלי המשנה והתלמודף ולא נסור ימין ושמאל מקבלתם בפירוש המצות היאך נעשה אותם כמשפטם, וכ"ש וק"ו באמונת אחדות הבורא ית' ושלילת מאורעות הגוף ומקריו ממנו ית'. מריה דאברהם, לא לנו ה' לא לנו לחשוב רבוי באלהותך כאשר עשו המקובלים החדשים על פי הזהר והתקונים שמיחסים אותם לרשב"י בשקר, כי נתחו (לפי עניות דעתם) את ה' אלהים אמת האל הגדול הגבור והנורא לנתחים וצורות ופרצופים רבים שונים זה מזה באיכותן וכמותן וחומרן וטבען, ויערכו את נתחיו לעין כל קורא בספריהם, וקורא להם מערכת האלהות, ויקראו לכל נתח בשם, כשם הגדולים אשר בארץ וזהו סדר עריכתם כמבואר בספרים, אין סוף העליון אשר אין למעלה ממנו, אדם קדמאה, אדם קדמון, עתיק ונוקביה, אריך אנפין ונוקביה. אבא ואימא עילאין, ישראל סבא ותבונה, יעקב ולאה זעיר ונוקביה, והוא נקרא ישראל והיא רחל, (עץ חיים שער המלכים פ"ט). ובשער אבא ואימא פ"ב כ' וכבר נודע מ"ש בענין לאה ורחל בהיותן עומדין באחורי זעיר אנפין שאז העקביים של לאה נכנסים תוך ראש רחל העומדת תחתיה, וכן הענין בכאן שעקביים של בינה נכנסין תוך ראש תבונה עכ"ל. וישסעו את אריה דבי עלאי אריה שאג מי לא יירא ה' אלהים דבר, כשסע הגדי ידים ורגלים, ועקביים וחזה ושוק וטבור, ושאר איברים נסתרים כגיד הערוה ותרין ביעין לכל אחד מן הפרצופים! ויראו לאחרים את ערות אבותיהם שבשמים (לפי דעתם שחושבים רבוי פרצופים כאל אחד) ויכניסו אותם במכחול בשפופרת ברחם הנקבות אשר בדו מלבם, כמבואר בעץ חיים בכמה מקומות ושם מבואר שהיסוד כולו הוא בחינת פנים ואין לו אחוריים לפי שכל היסוד של זכר נכנס צוך היסוד של נקבה, וכולו בחינת פנים, מה שאין כן בשאר הגוף שיש בו בחינת פנים הנחזרים נגד פני הנקבה! אוי לאותה בושה ואני לאותה כלמה לספר דברים מכוערים כאלה באלהינו שבשמים אשר מלא כל הארץ כבודו, ולקרוא לו בשם מגונה קצר אפים (ז"א). וכבר הודעתיך בסי' צ"ב דברי רבינו הגדול הרמב"ם ז"ל בספרו הנכבד מורה נבוכים שתורתינו הקדושה וכל הנביאים וחכמינו ז"ל הרחיקו מהשי"ת. שני חושים מהחמשה חושים שבאדם, והן הטעם והמשוש, וכ"ש המשגל שהוא הפחות והמגונה וכו'. וכבר ידעת מה שדרשו רז"ל על פסוק ותמת שם מרים, תקבר שם דאמר ר' אלעזר אף מרים מתה בנשיקה. איתא שם שם, ומפני מה לא נאמר בה על פי ה', שגנאי הדבר לומר ואם כל כך הקפיד הכתוב שלא לומר במרים הנביאה על פי ה' מטעם שגנאי הדבר לומר כן. כל שכן קל וחומר ליחס לשי"ת זכרות ונקבות וזווגין והטפת טפין ולומר שמקום חיבור גיד הערוה הוא למעלה מתרין פרקין נצח והוד כמו שהוא באדם כמו שכתב מהרח"ו בעץ חיים שער נקודים פ"ג. וכבר הבאתי לעיל סי' צ"ד דברי בראשית רבה ודברי הירושלמי על פסוק תאלמנה שפתי שקר הדוברות על צדיק עתק בגאוה ובוז, יתחרשן יתפרכן ישתתקן וכו' הדוברות על צדיק עתק. על צדיקו של עולם דברים שהעתיק מבריותיו בגאוה ובוז, שהוא מתגאה לומר אני דורש במעשה בראשית, סבור שהוא כמגאה ואינו אלא ככבזה. אמר ר' חנינה המתכבד בקלון חבירו אין לו חלק לעולם הבא. המתכבד בכבוד חי העולמים לא כל שכן. מה כתיב בתריה מה רב טובך אשר צפנת ליריאיך, אל כמה רב טובך ע"כ. ומה הדורש במעשה בראשית כל, הדורש במעשה מרכבה עאכ"ו, וכ"ש וק"ו הקורש ברוכב ערבות שהוא למעלה ממעשה מרכבה שלא יהיה לו במה רב טובך.

| id | word | punctuation after |
|---|---|---|
| 9:19.1 | קלז | ) |
| 9:19.2 | ובמדרש |  |
| 9:19.3 | רבה |  |
| 9:19.4 | פרשת |  |
| 9:19.5 | קרח |  |
| 9:19.6 | דרשו |  |
| 9:19.7 | רז"ל |  |
| 9:19.8 | ששים |  |
| 9:19.9 | המה |  |
| 9:19.10 | מלכות | , |
| 9:19.11 | אלו |  |
| 9:19.12 | ס | ' |
| 9:19.13 | מסכתות |  |
| 9:19.14 | ושמנים |  |
| 9:19.15 | פילגשים |  |
| 9:19.16 | אלו |  |
| 9:19.17 | שמנים |  |
| 9:19.18 | בתי |  |
| 9:19.19 | מדרשות |  |
| 9:19.20 | שהיו |  |
| 9:19.21 | בירושלם |  |
| 9:19.22 | כנגד |  |
| 9:19.23 | פתחיה | , |
| 9:19.24 | ועלמות |  |
| 9:19.25 | אין |  |
| 9:19.26 | מספר | , |
| 9:19.27 | משנה |  |
| 9:19.28 | החיצונה | , |
| 9:19.29 | ופירש |  |
| 9:19.30 | הרב |  |
| 9:19.31 | עץ |  |
| 9:19.32 | יוסף |  |
| 9:19.33 | ומת"כ |  |
| 9:19.34 | דקרי |  |
| 9:19.35 | לבתי |  |
| 9:19.36 | מדרשות |  |
| 9:19.37 | פילגשים |  |
| 9:19.38 | על |  |
| 9:19.39 | שם |  |
| 9:19.40 | שנושאין |  |
| 9:19.41 | ונותנים |  |
| 9:19.42 | בהם |  |
| 9:19.43 | וחולקים |  |
| 9:19.44 | בהם |  |
| 9:19.45 | וחולקים |  |
| 9:19.46 | זה |  |
| 9:19.47 | עם |  |
| 9:19.48 | זה |  |
| 9:19.49 | כמו |  |
| 9:19.50 | פילגשם |  |
| 9:19.51 | ומחבר |  |
| 9:19.52 | הזהר |  |
| 9:19.53 | כשנאתו |  |
| 9:19.54 | למשנה |  |
| 9:19.55 | ולתלמוד |  |
| 9:19.56 | דימה |  |
| 9:19.57 | שקלא |  |
| 9:19.58 | וטריא |  |
| 9:19.59 | דרבנן |  |
| 9:19.60 | לכלביך |  |
| 9:19.61 | דצוחון |  |
| 9:19.62 | הב |  |
| 9:19.63 | הב |  |
| 9:19.64 | ושם |  |
| 9:19.65 | אותם |  |
| 9:19.66 | מיורשי |  |
| 9:19.67 | גהינם |  |
| 9:19.68 | האומרת |  |
| 9:19.69 | הב |  |
| 9:19.70 | הב | ) |
| 9:19.71 | ופי | ' |
| 9:19.72 | משנה |  |
| 9:19.73 | החיצונה |  |
| 9:19.74 | מה |  |
| 9:19.75 | שלמדו |  |
| 9:19.76 | היחידים |  |
| 9:19.77 | חוץ |  |
| 9:19.78 | לאלו |  |
| 9:19.79 | בתי |  |
| 9:19.80 | מדרשות | , |
| 9:19.81 | וביד |  |
| 9:19.82 | יוסף |  |
| 9:19.83 | פי | ' |
| 9:19.84 | אלו |  |
| 9:19.85 | הברייתות | , |
| 9:19.86 | ע"כ | , |
| 9:19.87 | וכן |  |
| 9:19.88 | מצאתי |  |
| 9:19.89 | בשיר |  |
| 9:19.90 | השירים |  |
| 9:19.91 | רבה |  |
| 9:19.92 | וז"ל |  |
| 9:19.93 | ששים |  |
| 9:19.94 | המה |  |
| 9:19.95 | מלכות | , |
| 9:19.96 | אלו |  |
| 9:19.97 | ששים |  |
| 9:19.98 | מסכתות |  |
| 9:19.99 | של |  |
| 9:19.100 | הלכות |  |
| 9:19.101 | ושמנים |  |
| 9:19.102 | פילגשים |  |
| 9:19.103 | אלו |  |
| 9:19.104 | שמנים |  |
| 9:19.105 | פרשיות |  |
| 9:19.106 | שבתורת |  |
| 9:19.107 | כהנים | . |
| 9:19.108 | ועלמות |  |
| 9:19.109 | אין |  |
| 9:19.110 | מספר | , |
| 9:19.111 | אין |  |
| 9:19.112 | קץ |  |
| 9:19.113 | לתוספתות | , |
| 9:19.114 | אחת |  |
| 9:19.115 | היא |  |
| 9:19.116 | הן |  |
| 9:19.117 | חולקין |  |
| 9:19.118 | אלו |  |
| 9:19.119 | עם |  |
| 9:19.120 | אלו |  |
| 9:19.121 | וכולהון |  |
| 9:19.122 | דורשין |  |
| 9:19.123 | מטעם |  |
| 9:19.124 | אחד | , |
| 9:19.125 | ומהלכה |  |
| 9:19.126 | אחת |  |
| 9:19.127 | מגזרה |  |
| 9:19.128 | שוה |  |
| 9:19.129 | מקל |  |
| 9:19.130 | וחומר |  |
| 9:19.131 | ע"כ | . |
| 9:19.132 | ראה |  |
| 9:19.133 | ידידי |  |
| 9:19.134 | הקורא |  |
| 9:19.135 | כמה |  |
| 9:19.136 | מחבבים |  |
| 9:19.137 | רז"ל |  |
| 9:19.138 | את |  |
| 9:19.139 | המשנה |  |
| 9:19.140 | והתלמוד |  |
| 9:19.141 | וקראו |  |
| 9:19.142 | להן |  |
| 9:19.143 | מלכות |  |
| 9:19.144 | ולא |  |
| 9:19.145 | שפחות | , |
| 9:19.146 | כמאמר |  |
| 9:19.147 | שלמה |  |
| 9:19.148 | כי |  |
| 9:19.149 | מלכים |  |
| 9:19.150 | ימלוכו | , |
| 9:19.151 | והמה |  |
| 9:19.152 | בתרה |  |
| 9:19.153 | של |  |
| 9:19.154 | תורה |  |
| 9:19.155 | כי |  |
| 9:19.156 | לא |  |
| 9:19.157 | נוכל |  |
| 9:19.158 | לעשות |  |
| 9:19.159 | דבר |  |
| 9:19.160 | ממצות |  |
| 9:19.161 | התורה |  |
| 9:19.162 | כי |  |
| 9:19.163 | אם |  |
| 9:19.164 | על |  |
| 9:19.165 | פי |  |
| 9:19.166 | קבלת |  |
| 9:19.167 | רז"ל |  |
| 9:19.168 | בעלי |  |
| 9:19.169 | המשנה |  |
| 9:19.170 | והתלמודף |  |
| 9:19.171 | ולא |  |
| 9:19.172 | נסור |  |
| 9:19.173 | ימין |  |
| 9:19.174 | ושמאל |  |
| 9:19.175 | מקבלתם |  |
| 9:19.176 | בפירוש |  |
| 9:19.177 | המצות |  |
| 9:19.178 | היאך |  |
| 9:19.179 | נעשה |  |
| 9:19.180 | אותם |  |
| 9:19.181 | כמשפטם | , |
| 9:19.182 | וכ"ש |  |
| 9:19.183 | וק"ו |  |
| 9:19.184 | באמונת |  |
| 9:19.185 | אחדות |  |
| 9:19.186 | הבורא |  |
| 9:19.187 | ית | ' |
| 9:19.188 | ושלילת |  |
| 9:19.189 | מאורעות |  |
| 9:19.190 | הגוף |  |
| 9:19.191 | ומקריו |  |
| 9:19.192 | ממנו |  |
| 9:19.193 | ית | '. |
| 9:19.194 | מריה |  |
| 9:19.195 | דאברהם | , |
| 9:19.196 | לא |  |
| 9:19.197 | לנו |  |
| 9:19.198 | ה | ' |
| 9:19.199 | לא |  |
| 9:19.200 | לנו |  |
| 9:19.201 | לחשוב |  |
| 9:19.202 | רבוי |  |
| 9:19.203 | באלהותך |  |
| 9:19.204 | כאשר |  |
| 9:19.205 | עשו |  |
| 9:19.206 | המקובלים |  |
| 9:19.207 | החדשים |  |
| 9:19.208 | על |  |
| 9:19.209 | פי |  |
| 9:19.210 | הזהר |  |
| 9:19.211 | והתקונים |  |
| 9:19.212 | שמיחסים |  |
| 9:19.213 | אותם |  |
| 9:19.214 | לרשב"י |  |
| 9:19.215 | בשקר | , |
| 9:19.216 | כי |  |
| 9:19.217 | נתחו |  |
| 9:19.218 | לפי |  |
| 9:19.219 | עניות |  |
| 9:19.220 | דעתם | ) |
| 9:19.221 | את |  |
| 9:19.222 | ה | ' |
| 9:19.223 | אלהים |  |
| 9:19.224 | אמת |  |
| 9:19.225 | האל |  |
| 9:19.226 | הגדול |  |
| 9:19.227 | הגבור |  |
| 9:19.228 | והנורא |  |
| 9:19.229 | לנתחים |  |
| 9:19.230 | וצורות |  |
| 9:19.231 | ופרצופים |  |
| 9:19.232 | רבים |  |
| 9:19.233 | שונים |  |
| 9:19.234 | זה |  |
| 9:19.235 | מזה |  |
| 9:19.236 | באיכותן |  |
| 9:19.237 | וכמותן |  |
| 9:19.238 | וחומרן |  |
| 9:19.239 | וטבען | , |
| 9:19.240 | ויערכו |  |
| 9:19.241 | את |  |
| 9:19.242 | נתחיו |  |
| 9:19.243 | לעין |  |
| 9:19.244 | כל |  |
| 9:19.245 | קורא |  |
| 9:19.246 | בספריהם | , |
| 9:19.247 | וקורא |  |
| 9:19.248 | להם |  |
| 9:19.249 | מערכת |  |
| 9:19.250 | האלהות | , |
| 9:19.251 | ויקראו |  |
| 9:19.252 | לכל |  |
| 9:19.253 | נתח |  |
| 9:19.254 | בשם | , |
| 9:19.255 | כשם |  |
| 9:19.256 | הגדולים |  |
| 9:19.257 | אשר |  |
| 9:19.258 | בארץ |  |
| 9:19.259 | וזהו |  |
| 9:19.260 | סדר |  |
| 9:19.261 | עריכתם |  |
| 9:19.262 | כמבואר |  |
| 9:19.263 | בספרים | , |
| 9:19.264 | אין |  |
| 9:19.265 | סוף |  |
| 9:19.266 | העליון |  |
| 9:19.267 | אשר |  |
| 9:19.268 | אין |  |
| 9:19.269 | למעלה |  |
| 9:19.270 | ממנו | , |
| 9:19.271 | אדם |  |
| 9:19.272 | קדמאה | , |
| 9:19.273 | אדם |  |
| 9:19.274 | קדמון | , |
| 9:19.275 | עתיק |  |
| 9:19.276 | ונוקביה | , |
| 9:19.277 | אריך |  |
| 9:19.278 | אנפין |  |
| 9:19.279 | ונוקביה | . |
| 9:19.280 | אבא |  |
| 9:19.281 | ואימא |  |
| 9:19.282 | עילאין | , |
| 9:19.283 | ישראל |  |
| 9:19.284 | סבא |  |
| 9:19.285 | ותבונה | , |
| 9:19.286 | יעקב |  |
| 9:19.287 | ולאה |  |
| 9:19.288 | זעיר |  |
| 9:19.289 | ונוקביה | , |
| 9:19.290 | והוא |  |
| 9:19.291 | נקרא |  |
| 9:19.292 | ישראל |  |
| 9:19.293 | והיא |  |
| 9:19.294 | רחל | , |
| 9:19.295 | עץ |  |
| 9:19.296 | חיים |  |
| 9:19.297 | שער |  |
| 9:19.298 | המלכים |  |
| 9:19.299 | פ"ט | ). |
| 9:19.300 | ובשער |  |
| 9:19.301 | אבא |  |
| 9:19.302 | ואימא |  |
| 9:19.303 | פ"ב |  |
| 9:19.304 | כ | ' |
| 9:19.305 | וכבר |  |
| 9:19.306 | נודע |  |
| 9:19.307 | מ"ש |  |
| 9:19.308 | בענין |  |
| 9:19.309 | לאה |  |
| 9:19.310 | ורחל |  |
| 9:19.311 | בהיותן |  |
| 9:19.312 | עומדין |  |
| 9:19.313 | באחורי |  |
| 9:19.314 | זעיר |  |
| 9:19.315 | אנפין |  |
| 9:19.316 | שאז |  |
| 9:19.317 | העקביים |  |
| 9:19.318 | של |  |
| 9:19.319 | לאה |  |
| 9:19.320 | נכנסים |  |
| 9:19.321 | תוך |  |
| 9:19.322 | ראש |  |
| 9:19.323 | רחל |  |
| 9:19.324 | העומדת |  |
| 9:19.325 | תחתיה | , |
| 9:19.326 | וכן |  |
| 9:19.327 | הענין |  |
| 9:19.328 | בכאן |  |
| 9:19.329 | שעקביים |  |
| 9:19.330 | של |  |
| 9:19.331 | בינה |  |
| 9:19.332 | נכנסין |  |
| 9:19.333 | תוך |  |
| 9:19.334 | ראש |  |
| 9:19.335 | תבונה |  |
| 9:19.336 | עכ"ל | . |
| 9:19.337 | וישסעו |  |
| 9:19.338 | את |  |
| 9:19.339 | אריה |  |
| 9:19.340 | דבי |  |
| 9:19.341 | עלאי |  |
| 9:19.342 | אריה |  |
| 9:19.343 | שאג |  |
| 9:19.344 | מי |  |
| 9:19.345 | לא |  |
| 9:19.346 | יירא |  |
| 9:19.347 | ה | ' |
| 9:19.348 | אלהים |  |
| 9:19.349 | דבר | , |
| 9:19.350 | כשסע |  |
| 9:19.351 | הגדי |  |
| 9:19.352 | ידים |  |
| 9:19.353 | ורגלים | , |
| 9:19.354 | ועקביים |  |
| 9:19.355 | וחזה |  |
| 9:19.356 | ושוק |  |
| 9:19.357 | וטבור | , |
| 9:19.358 | ושאר |  |
| 9:19.359 | איברים |  |
| 9:19.360 | נסתרים |  |
| 9:19.361 | כגיד |  |
| 9:19.362 | הערוה |  |
| 9:19.363 | ותרין |  |
| 9:19.364 | ביעין |  |
| 9:19.365 | לכל |  |
| 9:19.366 | אחד |  |
| 9:19.367 | מן |  |
| 9:19.368 | הפרצופים | ! |
| 9:19.369 | ויראו |  |
| 9:19.370 | לאחרים |  |
| 9:19.371 | את |  |
| 9:19.372 | ערות |  |
| 9:19.373 | אבותיהם |  |
| 9:19.374 | שבשמים |  |
| 9:19.375 | לפי |  |
| 9:19.376 | דעתם |  |
| 9:19.377 | שחושבים |  |
| 9:19.378 | רבוי |  |
| 9:19.379 | פרצופים |  |
| 9:19.380 | כאל |  |
| 9:19.381 | אחד | ) |
| 9:19.382 | ויכניסו |  |
| 9:19.383 | אותם |  |
| 9:19.384 | במכחול |  |
| 9:19.385 | בשפופרת |  |
| 9:19.386 | ברחם |  |
| 9:19.387 | הנקבות |  |
| 9:19.388 | אשר |  |
| 9:19.389 | בדו |  |
| 9:19.390 | מלבם | , |
| 9:19.391 | כמבואר |  |
| 9:19.392 | בעץ |  |
| 9:19.393 | חיים |  |
| 9:19.394 | בכמה |  |
| 9:19.395 | מקומות |  |
| 9:19.396 | ושם |  |
| 9:19.397 | מבואר |  |
| 9:19.398 | שהיסוד |  |
| 9:19.399 | כולו |  |
| 9:19.400 | הוא |  |
| 9:19.401 | בחינת |  |
| 9:19.402 | פנים |  |
| 9:19.403 | ואין |  |
| 9:19.404 | לו |  |
| 9:19.405 | אחוריים |  |
| 9:19.406 | לפי |  |
| 9:19.407 | שכל |  |
| 9:19.408 | היסוד |  |
| 9:19.409 | של |  |
| 9:19.410 | זכר |  |
| 9:19.411 | נכנס |  |
| 9:19.412 | צוך |  |
| 9:19.413 | היסוד |  |
| 9:19.414 | של |  |
| 9:19.415 | נקבה | , |
| 9:19.416 | וכולו |  |
| 9:19.417 | בחינת |  |
| 9:19.418 | פנים | , |
| 9:19.419 | מה |  |
| 9:19.420 | שאין |  |
| 9:19.421 | כן |  |
| 9:19.422 | בשאר |  |
| 9:19.423 | הגוף |  |
| 9:19.424 | שיש |  |
| 9:19.425 | בו |  |
| 9:19.426 | בחינת |  |
| 9:19.427 | פנים |  |
| 9:19.428 | הנחזרים |  |
| 9:19.429 | נגד |  |
| 9:19.430 | פני |  |
| 9:19.431 | הנקבה | ! |
| 9:19.432 | אוי |  |
| 9:19.433 | לאותה |  |
| 9:19.434 | בושה |  |
| 9:19.435 | ואני |  |
| 9:19.436 | לאותה |  |
| 9:19.437 | כלמה |  |
| 9:19.438 | לספר |  |
| 9:19.439 | דברים |  |
| 9:19.440 | מכוערים |  |
| 9:19.441 | כאלה |  |
| 9:19.442 | באלהינו |  |
| 9:19.443 | שבשמים |  |
| 9:19.444 | אשר |  |
| 9:19.445 | מלא |  |
| 9:19.446 | כל |  |
| 9:19.447 | הארץ |  |
| 9:19.448 | כבודו | , |
| 9:19.449 | ולקרוא |  |
| 9:19.450 | לו |  |
| 9:19.451 | בשם |  |
| 9:19.452 | מגונה |  |
| 9:19.453 | קצר |  |
| 9:19.454 | אפים |  |
| 9:19.455 | ז"א | ). |
| 9:19.456 | וכבר |  |
| 9:19.457 | הודעתיך |  |
| 9:19.458 | בסי | ' |
| 9:19.459 | צ"ב |  |
| 9:19.460 | דברי |  |
| 9:19.461 | רבינו |  |
| 9:19.462 | הגדול |  |
| 9:19.463 | הרמב"ם |  |
| 9:19.464 | ז"ל |  |
| 9:19.465 | בספרו |  |
| 9:19.466 | הנכבד |  |
| 9:19.467 | מורה |  |
| 9:19.468 | נבוכים |  |
| 9:19.469 | שתורתינו |  |
| 9:19.470 | הקדושה |  |
| 9:19.471 | וכל |  |
| 9:19.472 | הנביאים |  |
| 9:19.473 | וחכמינו |  |
| 9:19.474 | ז"ל |  |
| 9:19.475 | הרחיקו |  |
| 9:19.476 | מהשי"ת | . |
| 9:19.477 | שני |  |
| 9:19.478 | חושים |  |
| 9:19.479 | מהחמשה |  |
| 9:19.480 | חושים |  |
| 9:19.481 | שבאדם | , |
| 9:19.482 | והן |  |
| 9:19.483 | הטעם |  |
| 9:19.484 | והמשוש | , |
| 9:19.485 | וכ"ש |  |
| 9:19.486 | המשגל |  |
| 9:19.487 | שהוא |  |
| 9:19.488 | הפחות |  |
| 9:19.489 | והמגונה |  |
| 9:19.490 | וכו | '. |
| 9:19.491 | וכבר |  |
| 9:19.492 | ידעת |  |
| 9:19.493 | מה |  |
| 9:19.494 | שדרשו |  |
| 9:19.495 | רז"ל |  |
| 9:19.496 | על |  |
| 9:19.497 | פסוק |  |
| 9:19.498 | ותמת |  |
| 9:19.499 | שם |  |
| 9:19.500 | מרים | , |
| 9:19.501 | תקבר |  |
| 9:19.502 | שם |  |
| 9:19.503 | דאמר |  |
| 9:19.504 | ר | ' |
| 9:19.505 | אלעזר |  |
| 9:19.506 | אף |  |
| 9:19.507 | מרים |  |
| 9:19.508 | מתה |  |
| 9:19.509 | בנשיקה | . |
| 9:19.510 | איתא |  |
| 9:19.511 | שם |  |
| 9:19.512 | שם | , |
| 9:19.513 | ומפני |  |
| 9:19.514 | מה |  |
| 9:19.515 | לא |  |
| 9:19.516 | נאמר |  |
| 9:19.517 | בה |  |
| 9:19.518 | על |  |
| 9:19.519 | פי |  |
| 9:19.520 | ה | ', |
| 9:19.521 | שגנאי |  |
| 9:19.522 | הדבר |  |
| 9:19.523 | לומר |  |
| 9:19.524 | ואם |  |
| 9:19.525 | כל |  |
| 9:19.526 | כך |  |
| 9:19.527 | הקפיד |  |
| 9:19.528 | הכתוב |  |
| 9:19.529 | שלא |  |
| 9:19.530 | לומר |  |
| 9:19.531 | במרים |  |
| 9:19.532 | הנביאה |  |
| 9:19.533 | על |  |
| 9:19.534 | פי |  |
| 9:19.535 | ה | ' |
| 9:19.536 | מטעם |  |
| 9:19.537 | שגנאי |  |
| 9:19.538 | הדבר |  |
| 9:19.539 | לומר |  |
| 9:19.540 | כן | . |
| 9:19.541 | כל |  |
| 9:19.542 | שכן |  |
| 9:19.543 | קל |  |
| 9:19.544 | וחומר |  |
| 9:19.545 | ליחס |  |
| 9:19.546 | לשי"ת |  |
| 9:19.547 | זכרות |  |
| 9:19.548 | ונקבות |  |
| 9:19.549 | וזווגין |  |
| 9:19.550 | והטפת |  |
| 9:19.551 | טפין |  |
| 9:19.552 | ולומר |  |
| 9:19.553 | שמקום |  |
| 9:19.554 | חיבור |  |
| 9:19.555 | גיד |  |
| 9:19.556 | הערוה |  |
| 9:19.557 | הוא |  |
| 9:19.558 | למעלה |  |
| 9:19.559 | מתרין |  |
| 9:19.560 | פרקין |  |
| 9:19.561 | נצח |  |
| 9:19.562 | והוד |  |
| 9:19.563 | כמו |  |
| 9:19.564 | שהוא |  |
| 9:19.565 | באדם |  |
| 9:19.566 | כמו |  |
| 9:19.567 | שכתב |  |
| 9:19.568 | מהרח"ו |  |
| 9:19.569 | בעץ |  |
| 9:19.570 | חיים |  |
| 9:19.571 | שער |  |
| 9:19.572 | נקודים |  |
| 9:19.573 | פ"ג | . |
| 9:19.574 | וכבר |  |
| 9:19.575 | הבאתי |  |
| 9:19.576 | לעיל |  |
| 9:19.577 | סי | ' |
| 9:19.578 | צ"ד |  |
| 9:19.579 | דברי |  |
| 9:19.580 | בראשית |  |
| 9:19.581 | רבה |  |
| 9:19.582 | ודברי |  |
| 9:19.583 | הירושלמי |  |
| 9:19.584 | על |  |
| 9:19.585 | פסוק |  |
| 9:19.586 | תאלמנה |  |
| 9:19.587 | שפתי |  |
| 9:19.588 | שקר |  |
| 9:19.589 | הדוברות |  |
| 9:19.590 | על |  |
| 9:19.591 | צדיק |  |
| 9:19.592 | עתק |  |
| 9:19.593 | בגאוה |  |
| 9:19.594 | ובוז | , |
| 9:19.595 | יתחרשן |  |
| 9:19.596 | יתפרכן |  |
| 9:19.597 | ישתתקן |  |
| 9:19.598 | וכו | ' |
| 9:19.599 | הדוברות |  |
| 9:19.600 | על |  |
| 9:19.601 | צדיק |  |
| 9:19.602 | עתק | . |
| 9:19.603 | על |  |
| 9:19.604 | צדיקו |  |
| 9:19.605 | של |  |
| 9:19.606 | עולם |  |
| 9:19.607 | דברים |  |
| 9:19.608 | שהעתיק |  |
| 9:19.609 | מבריותיו |  |
| 9:19.610 | בגאוה |  |
| 9:19.611 | ובוז | , |
| 9:19.612 | שהוא |  |
| 9:19.613 | מתגאה |  |
| 9:19.614 | לומר |  |
| 9:19.615 | אני |  |
| 9:19.616 | דורש |  |
| 9:19.617 | במעשה |  |
| 9:19.618 | בראשית | , |
| 9:19.619 | סבור |  |
| 9:19.620 | שהוא |  |
| 9:19.621 | כמגאה |  |
| 9:19.622 | ואינו |  |
| 9:19.623 | אלא |  |
| 9:19.624 | ככבזה | . |
| 9:19.625 | אמר |  |
| 9:19.626 | ר | ' |
| 9:19.627 | חנינה |  |
| 9:19.628 | המתכבד |  |
| 9:19.629 | בקלון |  |
| 9:19.630 | חבירו |  |
| 9:19.631 | אין |  |
| 9:19.632 | לו |  |
| 9:19.633 | חלק |  |
| 9:19.634 | לעולם |  |
| 9:19.635 | הבא | . |
| 9:19.636 | המתכבד |  |
| 9:19.637 | בכבוד |  |
| 9:19.638 | חי |  |
| 9:19.639 | העולמים |  |
| 9:19.640 | לא |  |
| 9:19.641 | כל |  |
| 9:19.642 | שכן | . |
| 9:19.643 | מה |  |
| 9:19.644 | כתיב |  |
| 9:19.645 | בתריה |  |
| 9:19.646 | מה |  |
| 9:19.647 | רב |  |
| 9:19.648 | טובך |  |
| 9:19.649 | אשר |  |
| 9:19.650 | צפנת |  |
| 9:19.651 | ליריאיך | , |
| 9:19.652 | אל |  |
| 9:19.653 | כמה |  |
| 9:19.654 | רב |  |
| 9:19.655 | טובך |  |
| 9:19.656 | ע"כ | . |
| 9:19.657 | ומה |  |
| 9:19.658 | הדורש |  |
| 9:19.659 | במעשה |  |
| 9:19.660 | בראשית |  |
| 9:19.661 | כל | , |
| 9:19.662 | הדורש |  |
| 9:19.663 | במעשה |  |
| 9:19.664 | מרכבה |  |
| 9:19.665 | עאכ"ו | , |
| 9:19.666 | וכ"ש |  |
| 9:19.667 | וק"ו |  |
| 9:19.668 | הקורש |  |
| 9:19.669 | ברוכב |  |
| 9:19.670 | ערבות |  |
| 9:19.671 | שהוא |  |
| 9:19.672 | למעלה |  |
| 9:19.673 | ממעשה |  |
| 9:19.674 | מרכבה |  |
| 9:19.675 | שלא |  |
| 9:19.676 | יהיה |  |
| 9:19.677 | לו |  |
| 9:19.678 | במה |  |
| 9:19.679 | רב |  |
| 9:19.680 | טובך | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

