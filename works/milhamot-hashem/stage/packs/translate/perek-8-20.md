# Translation pack: milhamot-hashem perek-8-20 (Milchamot ha-Shem 8:24–8:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-20`.

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
- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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


## Milchamot ha-Shem 8:24

Hebrew (Wikisource): קיז) וכלל הדבר כי הפילסוף המסית מחבר הזהר שם רעיונותיו ולבו בשאלת מה למעלה ומה למטה וכו' והפקיר איסור זה הנזכר במשנה, ולא הניח אלא מה למעלה מכיפת עיגול הא"ס אשר לדעתו נצתמצם למעלה סביב, ואע"פ שתמצא אותו אומר דאית אתר דלא קימא לשאלה ביה, אין זה אלא לגנוב את העין כאלו הוא מאמין באיסור זה, והראיה שהרי דבר במהות האלוה וצמצומו שנעשה ככדור עגול חלול, וכיון שדבר באין סוף שהוא למעלה מן הכל, איך יהיה אסור לשאול ולדבר במה שלמטה ממנו תוך חללו? וכבר גילה דעתו בהדיא, שכל אסור ופסול וטמא הנזכרים במשנה הם מסטרא דיצר הרע, ולכך נהג בזה הפקר ודבר במהות יוצר הכל ואיך ציורו ואופן חנייתו, ועשה לו גופים ופרצופים קבועים שונים זה מזה בכמותם ואיכותם ומקומם, זה למעלה וזה למטה, וזה ארוך וזה קצר כמו שיראה המעיין בו ובעץ חיים ופריו לכהרח"ו, ושבש ובלבל אמונת ישראל הטהורה המקובלת בידם, והשוה אותה לאמונת העכו"ם. ורבים מעמנו נמשכו אחריו לתומם בהאמנים שמהתנא הקדוש רשב"י יצאו הדברים הכבושים אשר היו קבורים, ולא שמו אל לבם איך אפשר לקבל תורה חדשה מיד עכו"ם, ואיך נקברה בארץ בארגז, ומי קברה? ולמה נתגלית בימי רשב"י וחזרה ונקברה ונשכח זכרה מכל היהודים? אשר על כן חלילה לכל משכיל בתורת ה' שבכתב ושבעל פה להאמין שדברים כאלה יצאו מפי רשב"י ע"ה ולא משום תנא או אמורא, כי אדרבא מצינו ראינו שהיו מרחיקים אלו הדעות הרעות ומתנגדים להם, וכמו שכבר הבאתי קצת מאמרים שמהם נלמד התנגדותם לאלו הדעות החדשות הנזכרים בקבלה החדשה, דממה נפשך הואיל ויש לפרצופים דין אלהות כמו שכתב הרש"ב בגליון פרשת בהר דף ק"ט והרב משנת חסידים סוף מסכת העשייה, ובית קדש הקדשים דף נ"ו ע"ב! נמצאנו מאמינים ועובדים לאלוהות רבים נבראים שכל אחד מהם הוא סבה למה שאחריו כמו שכתב לעיל בשם זוהר בראשית דף כ"ב שהוא קורא לכל אחד מהם בשם הוי'ה ואדנות ואלהים וקב"ה ושאר שמות המיוחדים, עד נדר שכשישמע האיש הישראלי את חבירו מזכיר את השם הנכבד הויה או אלהים או הקב"ה, צריך השומע לשאול איזה הוא הקב"ה שאתה מזכיר, אם עתיק, או אריך, או אבא, או אימא, או זעיר, או נוקבא דיליה, כמו שעשה הוא עצמו מחבר הזהר בפסוק ויאמר אלהים נעשה אדם, ואמר מאן אלהים, כדלעיל סי' כ'! כי האלהים העליון לא תוכל לומר (לפי דבריהם) כי הוא אין לו שם ולא נקודה, ולא תוכל לכנותו בשם מן השמות הקדושים כמו שכתבו המקובלים החדשים, ונמצאנו מאמינים באלהים רבים נבראים ח"ו.

| id | word | punctuation after |
|---|---|---|
| 8:24.1 | קיז | ) |
| 8:24.2 | וכלל |  |
| 8:24.3 | הדבר |  |
| 8:24.4 | כי |  |
| 8:24.5 | הפילסוף |  |
| 8:24.6 | המסית |  |
| 8:24.7 | מחבר |  |
| 8:24.8 | הזהר |  |
| 8:24.9 | שם |  |
| 8:24.10 | רעיונותיו |  |
| 8:24.11 | ולבו |  |
| 8:24.12 | בשאלת |  |
| 8:24.13 | מה |  |
| 8:24.14 | למעלה |  |
| 8:24.15 | ומה |  |
| 8:24.16 | למטה |  |
| 8:24.17 | וכו | ' |
| 8:24.18 | והפקיר |  |
| 8:24.19 | איסור |  |
| 8:24.20 | זה |  |
| 8:24.21 | הנזכר |  |
| 8:24.22 | במשנה | , |
| 8:24.23 | ולא |  |
| 8:24.24 | הניח |  |
| 8:24.25 | אלא |  |
| 8:24.26 | מה |  |
| 8:24.27 | למעלה |  |
| 8:24.28 | מכיפת |  |
| 8:24.29 | עיגול |  |
| 8:24.30 | הא"ס |  |
| 8:24.31 | אשר |  |
| 8:24.32 | לדעתו |  |
| 8:24.33 | נצתמצם |  |
| 8:24.34 | למעלה |  |
| 8:24.35 | סביב | , |
| 8:24.36 | ואע"פ |  |
| 8:24.37 | שתמצא |  |
| 8:24.38 | אותו |  |
| 8:24.39 | אומר |  |
| 8:24.40 | דאית |  |
| 8:24.41 | אתר |  |
| 8:24.42 | דלא |  |
| 8:24.43 | קימא |  |
| 8:24.44 | לשאלה |  |
| 8:24.45 | ביה | , |
| 8:24.46 | אין |  |
| 8:24.47 | זה |  |
| 8:24.48 | אלא |  |
| 8:24.49 | לגנוב |  |
| 8:24.50 | את |  |
| 8:24.51 | העין |  |
| 8:24.52 | כאלו |  |
| 8:24.53 | הוא |  |
| 8:24.54 | מאמין |  |
| 8:24.55 | באיסור |  |
| 8:24.56 | זה | , |
| 8:24.57 | והראיה |  |
| 8:24.58 | שהרי |  |
| 8:24.59 | דבר |  |
| 8:24.60 | במהות |  |
| 8:24.61 | האלוה |  |
| 8:24.62 | וצמצומו |  |
| 8:24.63 | שנעשה |  |
| 8:24.64 | ככדור |  |
| 8:24.65 | עגול |  |
| 8:24.66 | חלול | , |
| 8:24.67 | וכיון |  |
| 8:24.68 | שדבר |  |
| 8:24.69 | באין |  |
| 8:24.70 | סוף |  |
| 8:24.71 | שהוא |  |
| 8:24.72 | למעלה |  |
| 8:24.73 | מן |  |
| 8:24.74 | הכל | , |
| 8:24.75 | איך |  |
| 8:24.76 | יהיה |  |
| 8:24.77 | אסור |  |
| 8:24.78 | לשאול |  |
| 8:24.79 | ולדבר |  |
| 8:24.80 | במה |  |
| 8:24.81 | שלמטה |  |
| 8:24.82 | ממנו |  |
| 8:24.83 | תוך |  |
| 8:24.84 | חללו | ? |
| 8:24.85 | וכבר |  |
| 8:24.86 | גילה |  |
| 8:24.87 | דעתו |  |
| 8:24.88 | בהדיא | , |
| 8:24.89 | שכל |  |
| 8:24.90 | אסור |  |
| 8:24.91 | ופסול |  |
| 8:24.92 | וטמא |  |
| 8:24.93 | הנזכרים |  |
| 8:24.94 | במשנה |  |
| 8:24.95 | הם |  |
| 8:24.96 | מסטרא |  |
| 8:24.97 | דיצר |  |
| 8:24.98 | הרע | , |
| 8:24.99 | ולכך |  |
| 8:24.100 | נהג |  |
| 8:24.101 | בזה |  |
| 8:24.102 | הפקר |  |
| 8:24.103 | ודבר |  |
| 8:24.104 | במהות |  |
| 8:24.105 | יוצר |  |
| 8:24.106 | הכל |  |
| 8:24.107 | ואיך |  |
| 8:24.108 | ציורו |  |
| 8:24.109 | ואופן |  |
| 8:24.110 | חנייתו | , |
| 8:24.111 | ועשה |  |
| 8:24.112 | לו |  |
| 8:24.113 | גופים |  |
| 8:24.114 | ופרצופים |  |
| 8:24.115 | קבועים |  |
| 8:24.116 | שונים |  |
| 8:24.117 | זה |  |
| 8:24.118 | מזה |  |
| 8:24.119 | בכמותם |  |
| 8:24.120 | ואיכותם |  |
| 8:24.121 | ומקומם | , |
| 8:24.122 | זה |  |
| 8:24.123 | למעלה |  |
| 8:24.124 | וזה |  |
| 8:24.125 | למטה | , |
| 8:24.126 | וזה |  |
| 8:24.127 | ארוך |  |
| 8:24.128 | וזה |  |
| 8:24.129 | קצר |  |
| 8:24.130 | כמו |  |
| 8:24.131 | שיראה |  |
| 8:24.132 | המעיין |  |
| 8:24.133 | בו |  |
| 8:24.134 | ובעץ |  |
| 8:24.135 | חיים |  |
| 8:24.136 | ופריו |  |
| 8:24.137 | לכהרח"ו | , |
| 8:24.138 | ושבש |  |
| 8:24.139 | ובלבל |  |
| 8:24.140 | אמונת |  |
| 8:24.141 | ישראל |  |
| 8:24.142 | הטהורה |  |
| 8:24.143 | המקובלת |  |
| 8:24.144 | בידם | , |
| 8:24.145 | והשוה |  |
| 8:24.146 | אותה |  |
| 8:24.147 | לאמונת |  |
| 8:24.148 | העכו"ם | . |
| 8:24.149 | ורבים |  |
| 8:24.150 | מעמנו |  |
| 8:24.151 | נמשכו |  |
| 8:24.152 | אחריו |  |
| 8:24.153 | לתומם |  |
| 8:24.154 | בהאמנים |  |
| 8:24.155 | שמהתנא |  |
| 8:24.156 | הקדוש |  |
| 8:24.157 | רשב"י |  |
| 8:24.158 | יצאו |  |
| 8:24.159 | הדברים |  |
| 8:24.160 | הכבושים |  |
| 8:24.161 | אשר |  |
| 8:24.162 | היו |  |
| 8:24.163 | קבורים | , |
| 8:24.164 | ולא |  |
| 8:24.165 | שמו |  |
| 8:24.166 | אל |  |
| 8:24.167 | לבם |  |
| 8:24.168 | איך |  |
| 8:24.169 | אפשר |  |
| 8:24.170 | לקבל |  |
| 8:24.171 | תורה |  |
| 8:24.172 | חדשה |  |
| 8:24.173 | מיד |  |
| 8:24.174 | עכו"ם | , |
| 8:24.175 | ואיך |  |
| 8:24.176 | נקברה |  |
| 8:24.177 | בארץ |  |
| 8:24.178 | בארגז | , |
| 8:24.179 | ומי |  |
| 8:24.180 | קברה | ? |
| 8:24.181 | ולמה |  |
| 8:24.182 | נתגלית |  |
| 8:24.183 | בימי |  |
| 8:24.184 | רשב"י |  |
| 8:24.185 | וחזרה |  |
| 8:24.186 | ונקברה |  |
| 8:24.187 | ונשכח |  |
| 8:24.188 | זכרה |  |
| 8:24.189 | מכל |  |
| 8:24.190 | היהודים | ? |
| 8:24.191 | אשר |  |
| 8:24.192 | על |  |
| 8:24.193 | כן |  |
| 8:24.194 | חלילה |  |
| 8:24.195 | לכל |  |
| 8:24.196 | משכיל |  |
| 8:24.197 | בתורת |  |
| 8:24.198 | ה | ' |
| 8:24.199 | שבכתב |  |
| 8:24.200 | ושבעל |  |
| 8:24.201 | פה |  |
| 8:24.202 | להאמין |  |
| 8:24.203 | שדברים |  |
| 8:24.204 | כאלה |  |
| 8:24.205 | יצאו |  |
| 8:24.206 | מפי |  |
| 8:24.207 | רשב"י |  |
| 8:24.208 | ע"ה |  |
| 8:24.209 | ולא |  |
| 8:24.210 | משום |  |
| 8:24.211 | תנא |  |
| 8:24.212 | או |  |
| 8:24.213 | אמורא | , |
| 8:24.214 | כי |  |
| 8:24.215 | אדרבא |  |
| 8:24.216 | מצינו |  |
| 8:24.217 | ראינו |  |
| 8:24.218 | שהיו |  |
| 8:24.219 | מרחיקים |  |
| 8:24.220 | אלו |  |
| 8:24.221 | הדעות |  |
| 8:24.222 | הרעות |  |
| 8:24.223 | ומתנגדים |  |
| 8:24.224 | להם | , |
| 8:24.225 | וכמו |  |
| 8:24.226 | שכבר |  |
| 8:24.227 | הבאתי |  |
| 8:24.228 | קצת |  |
| 8:24.229 | מאמרים |  |
| 8:24.230 | שמהם |  |
| 8:24.231 | נלמד |  |
| 8:24.232 | התנגדותם |  |
| 8:24.233 | לאלו |  |
| 8:24.234 | הדעות |  |
| 8:24.235 | החדשות |  |
| 8:24.236 | הנזכרים |  |
| 8:24.237 | בקבלה |  |
| 8:24.238 | החדשה | , |
| 8:24.239 | דממה |  |
| 8:24.240 | נפשך |  |
| 8:24.241 | הואיל |  |
| 8:24.242 | ויש |  |
| 8:24.243 | לפרצופים |  |
| 8:24.244 | דין |  |
| 8:24.245 | אלהות |  |
| 8:24.246 | כמו |  |
| 8:24.247 | שכתב |  |
| 8:24.248 | הרש"ב |  |
| 8:24.249 | בגליון |  |
| 8:24.250 | פרשת |  |
| 8:24.251 | בהר |  |
| 8:24.252 | דף |  |
| 8:24.253 | ק"ט |  |
| 8:24.254 | והרב |  |
| 8:24.255 | משנת |  |
| 8:24.256 | חסידים |  |
| 8:24.257 | סוף |  |
| 8:24.258 | מסכת |  |
| 8:24.259 | העשייה | , |
| 8:24.260 | ובית |  |
| 8:24.261 | קדש |  |
| 8:24.262 | הקדשים |  |
| 8:24.263 | דף |  |
| 8:24.264 | נ"ו |  |
| 8:24.265 | ע"ב | ! |
| 8:24.266 | נמצאנו |  |
| 8:24.267 | מאמינים |  |
| 8:24.268 | ועובדים |  |
| 8:24.269 | לאלוהות |  |
| 8:24.270 | רבים |  |
| 8:24.271 | נבראים |  |
| 8:24.272 | שכל |  |
| 8:24.273 | אחד |  |
| 8:24.274 | מהם |  |
| 8:24.275 | הוא |  |
| 8:24.276 | סבה |  |
| 8:24.277 | למה |  |
| 8:24.278 | שאחריו |  |
| 8:24.279 | כמו |  |
| 8:24.280 | שכתב |  |
| 8:24.281 | לעיל |  |
| 8:24.282 | בשם |  |
| 8:24.283 | זוהר |  |
| 8:24.284 | בראשית |  |
| 8:24.285 | דף |  |
| 8:24.286 | כ"ב |  |
| 8:24.287 | שהוא |  |
| 8:24.288 | קורא |  |
| 8:24.289 | לכל |  |
| 8:24.290 | אחד |  |
| 8:24.291 | מהם |  |
| 8:24.292 | בשם |  |
| 8:24.293 | הוי'ה |  |
| 8:24.294 | ואדנות |  |
| 8:24.295 | ואלהים |  |
| 8:24.296 | וקב"ה |  |
| 8:24.297 | ושאר |  |
| 8:24.298 | שמות |  |
| 8:24.299 | המיוחדים | , |
| 8:24.300 | עד |  |
| 8:24.301 | נדר |  |
| 8:24.302 | שכשישמע |  |
| 8:24.303 | האיש |  |
| 8:24.304 | הישראלי |  |
| 8:24.305 | את |  |
| 8:24.306 | חבירו |  |
| 8:24.307 | מזכיר |  |
| 8:24.308 | את |  |
| 8:24.309 | השם |  |
| 8:24.310 | הנכבד |  |
| 8:24.311 | הויה |  |
| 8:24.312 | או |  |
| 8:24.313 | אלהים |  |
| 8:24.314 | או |  |
| 8:24.315 | הקב"ה | , |
| 8:24.316 | צריך |  |
| 8:24.317 | השומע |  |
| 8:24.318 | לשאול |  |
| 8:24.319 | איזה |  |
| 8:24.320 | הוא |  |
| 8:24.321 | הקב"ה |  |
| 8:24.322 | שאתה |  |
| 8:24.323 | מזכיר | , |
| 8:24.324 | אם |  |
| 8:24.325 | עתיק | , |
| 8:24.326 | או |  |
| 8:24.327 | אריך | , |
| 8:24.328 | או |  |
| 8:24.329 | אבא | , |
| 8:24.330 | או |  |
| 8:24.331 | אימא | , |
| 8:24.332 | או |  |
| 8:24.333 | זעיר | , |
| 8:24.334 | או |  |
| 8:24.335 | נוקבא |  |
| 8:24.336 | דיליה | , |
| 8:24.337 | כמו |  |
| 8:24.338 | שעשה |  |
| 8:24.339 | הוא |  |
| 8:24.340 | עצמו |  |
| 8:24.341 | מחבר |  |
| 8:24.342 | הזהר |  |
| 8:24.343 | בפסוק |  |
| 8:24.344 | ויאמר |  |
| 8:24.345 | אלהים |  |
| 8:24.346 | נעשה |  |
| 8:24.347 | אדם | , |
| 8:24.348 | ואמר |  |
| 8:24.349 | מאן |  |
| 8:24.350 | אלהים | , |
| 8:24.351 | כדלעיל |  |
| 8:24.352 | סי | ' |
| 8:24.353 | כ | '! |
| 8:24.354 | כי |  |
| 8:24.355 | האלהים |  |
| 8:24.356 | העליון |  |
| 8:24.357 | לא |  |
| 8:24.358 | תוכל |  |
| 8:24.359 | לומר |  |
| 8:24.360 | לפי |  |
| 8:24.361 | דבריהם | ) |
| 8:24.362 | כי |  |
| 8:24.363 | הוא |  |
| 8:24.364 | אין |  |
| 8:24.365 | לו |  |
| 8:24.366 | שם |  |
| 8:24.367 | ולא |  |
| 8:24.368 | נקודה | , |
| 8:24.369 | ולא |  |
| 8:24.370 | תוכל |  |
| 8:24.371 | לכנותו |  |
| 8:24.372 | בשם |  |
| 8:24.373 | מן |  |
| 8:24.374 | השמות |  |
| 8:24.375 | הקדושים |  |
| 8:24.376 | כמו |  |
| 8:24.377 | שכתבו |  |
| 8:24.378 | המקובלים |  |
| 8:24.379 | החדשים | , |
| 8:24.380 | ונמצאנו |  |
| 8:24.381 | מאמינים |  |
| 8:24.382 | באלהים |  |
| 8:24.383 | רבים |  |
| 8:24.384 | נבראים |  |
| 8:24.385 | ח"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

