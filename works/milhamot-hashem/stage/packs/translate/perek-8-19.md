# Translation pack: milhamot-hashem perek-8-19 (Milchamot ha-Shem 8:23–8:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-19`.

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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 8:23

Hebrew (Wikisource): קיו) וכבר מלאו אזניך ורעיונך, כי המקובלים עובדים לנברא כמו שנת' לעיל סי' מ"ח בשם יושר לבב דף ג' ע"ב וז"ל דע את אלהי אביך עבדהו, דע את אלהי אביך כולל ה' פרצופין, ועבדהו לז"א אע"פ שהוא נברא, כי נשמתו את העובד בזה כי מבלעדיו אין כל נמצא, משום הכי לא אמר דע אלהי אביך ולא את אלהי אביך ע"כ. וכוונתו כי במלת את לרבות נשמת זעיר אנפין. והנה בכל תורה לא מצינו שיהא דבר הבא מרבוי עוקר את העיקר שנזכר בכתוב בהדיא, רק שיהא מצורף אל העיקר ומשתתף באיסור עמו כמו שדרשו רז"ל בכתובות כבד את אביך ואת אמך, את אביך זו אשת אביך, ואת אמך זה בעל אמך, שיהא מכבד את אשת אביו עם אביו, ואת בעל אמו עם אמו, וכן מה שדרשו רז"ל וערלתם ערלתו את פריו את הטפל לפריו, דהיינו הקליפין והגרעינין שהן אסורין בערלה כפרי עצמו ובנדון זה עשו המקובלים החדשים דבר הטפל הבא מרבוי את לעיקר ואמרו שהם עובדים לנשמתו, והגוף שהוא נברא טפל. ועובדים וגוף ולנשמה, כמו שביארנו לעיל סי' ט"ל וסי' ס"ה! ונמצאו משתפים שם שמים ודבר נברא, ועוברים על דברי רשב"י שאמר כל המשתף שם שמים ודבר אחר נעקר מן העולם. והן עתה אוסיף להעירך על קצת מאמרים מהזהר דמוכח מנייהו שזעיר אנפין נברא עיין בזהר בראית דף ט"ו ומקדש מלך שם מבואר דז"א נברא הוא וכבר הבאתי קצת מזה בסי' כ"ד, גם בדף כ"ו ע"א ויטע ה' אלהים, אבא ואימא, גן דא שכינתא תתאה, עדן דא אימא עלאה, את האדם דא עמודא דאמצעיתא (שהוא ז"א) עוד שם ויצמח ה' אלהים אבא ואימא, כל עץ נחמד דא צדיק (יסוד) וטוב למאכל, דא עמודא דאמצעיתא וכו' ועץ החיים דהוא אילנא דחי יהי נטיע בנו גינתא וכו', והראיה היותר ברורא הוא האמור באדרא רבה דף קל"ח ע"ב, ה' פעלך בקרב שנים חייהו. האי לעתיק יומין אתאמר, מאן פעלך, זעיר אנפין, חייהו למאן? לזעיר אנפין דכל נהירו דיליה מאינון שנים קדמוניות! מזה יתבאר לך בבירור שזעיר אנפין נברא, וכבר ידעת דעת הפילסוף המסית מחבר הזהר ורוב המקובלים הנמשכים אחריו, שלא תכון תפלה ועבודה רק לזעיר אנפין הנברא הלזה, ולדעת בעל עז לאלהים אין העבודה אלא למלכא קדישא דכל קדישין הנקרא לבא דכל ליבין והוא מתלבש תוך זעיר אנפין כמבואר שם בבית קדש הקדשים פרק י"ט דף נ"ט ע"ד, ואע"פ שיצא ממעי אבא ואימא הוא העיקר ואף אבא ואימא צריכים לכבדו כמלך בשר ודם שמולך שכל העם מכבדים אותו. ואף אביו ואמו צריכים לכבדו, שהוא המלך השליט, ואליו תפלותינו והוא מלכינו ואלהינו, אבל המתפלל לאל עליון שהוא הסבה הראשונה אין תפלתו תפלה, כי איך נאמר ברוך אתה ה', ממי יתברך ויבוא לו השפע? (בית קדש הקדשים דף כ"ו ע"ד). ועיין בזהר רעיא מהימנא דף ק"ט ע"ב שמבאר שם, כי האלהים העליון גלית ליה גופא ולא איברים ולא נוקבא, והוא רק מחבר ומקשר את האלהים התחתונים בעלי גוף ואיברים זכרים ונקיבות עם נקבותיהם להיות אחד, כטיט המחבר אבני הבית, או כחשוקי ברזל ונחשת הקושרים ומחברים את לוחות חביות של עץ להיות אחד. יתעלה השי"ת ויתרומם מכל מה שחשבו בו הטפשים והסכלים! ולכן אמרו המקובלים החדשים שהמתפלל אליו אין תפלתו תפלה, וכל הקורא אליו לא יענהו, והוא מרומם ומסולק מכל ברכה שאין לו ממי להתברך, ולא שם לו שיגבילהו, כי הכח והממשלה להאלהים שלמטה ממנו שהם החמשה פרצופים הכוללים, והמה נקראים בשם הוי'ה ואדנות ואלהים, ועיקר שם המיוחד שם הוי'ה נקרא להאלוה קצר אפים (זעיר אנפין) עם זוגתו, והוא השולט על כל הברואים כמו שנת' לעיל בכמה מקומות.

| id | word | punctuation after |
|---|---|---|
| 8:23.1 | קיו | ) |
| 8:23.2 | וכבר |  |
| 8:23.3 | מלאו |  |
| 8:23.4 | אזניך |  |
| 8:23.5 | ורעיונך | , |
| 8:23.6 | כי |  |
| 8:23.7 | המקובלים |  |
| 8:23.8 | עובדים |  |
| 8:23.9 | לנברא |  |
| 8:23.10 | כמו |  |
| 8:23.11 | שנת | ' |
| 8:23.12 | לעיל |  |
| 8:23.13 | סי | ' |
| 8:23.14 | מ"ח |  |
| 8:23.15 | בשם |  |
| 8:23.16 | יושר |  |
| 8:23.17 | לבב |  |
| 8:23.18 | דף |  |
| 8:23.19 | ג | ' |
| 8:23.20 | ע"ב |  |
| 8:23.21 | וז"ל |  |
| 8:23.22 | דע |  |
| 8:23.23 | את |  |
| 8:23.24 | אלהי |  |
| 8:23.25 | אביך |  |
| 8:23.26 | עבדהו | , |
| 8:23.27 | דע |  |
| 8:23.28 | את |  |
| 8:23.29 | אלהי |  |
| 8:23.30 | אביך |  |
| 8:23.31 | כולל |  |
| 8:23.32 | ה | ' |
| 8:23.33 | פרצופין | , |
| 8:23.34 | ועבדהו |  |
| 8:23.35 | לז"א |  |
| 8:23.36 | אע"פ |  |
| 8:23.37 | שהוא |  |
| 8:23.38 | נברא | , |
| 8:23.39 | כי |  |
| 8:23.40 | נשמתו |  |
| 8:23.41 | את |  |
| 8:23.42 | העובד |  |
| 8:23.43 | בזה |  |
| 8:23.44 | כי |  |
| 8:23.45 | מבלעדיו |  |
| 8:23.46 | אין |  |
| 8:23.47 | כל |  |
| 8:23.48 | נמצא | , |
| 8:23.49 | משום |  |
| 8:23.50 | הכי |  |
| 8:23.51 | לא |  |
| 8:23.52 | אמר |  |
| 8:23.53 | דע |  |
| 8:23.54 | אלהי |  |
| 8:23.55 | אביך |  |
| 8:23.56 | ולא |  |
| 8:23.57 | את |  |
| 8:23.58 | אלהי |  |
| 8:23.59 | אביך |  |
| 8:23.60 | ע"כ | . |
| 8:23.61 | וכוונתו |  |
| 8:23.62 | כי |  |
| 8:23.63 | במלת |  |
| 8:23.64 | את |  |
| 8:23.65 | לרבות |  |
| 8:23.66 | נשמת |  |
| 8:23.67 | זעיר |  |
| 8:23.68 | אנפין | . |
| 8:23.69 | והנה |  |
| 8:23.70 | בכל |  |
| 8:23.71 | תורה |  |
| 8:23.72 | לא |  |
| 8:23.73 | מצינו |  |
| 8:23.74 | שיהא |  |
| 8:23.75 | דבר |  |
| 8:23.76 | הבא |  |
| 8:23.77 | מרבוי |  |
| 8:23.78 | עוקר |  |
| 8:23.79 | את |  |
| 8:23.80 | העיקר |  |
| 8:23.81 | שנזכר |  |
| 8:23.82 | בכתוב |  |
| 8:23.83 | בהדיא | , |
| 8:23.84 | רק |  |
| 8:23.85 | שיהא |  |
| 8:23.86 | מצורף |  |
| 8:23.87 | אל |  |
| 8:23.88 | העיקר |  |
| 8:23.89 | ומשתתף |  |
| 8:23.90 | באיסור |  |
| 8:23.91 | עמו |  |
| 8:23.92 | כמו |  |
| 8:23.93 | שדרשו |  |
| 8:23.94 | רז"ל |  |
| 8:23.95 | בכתובות |  |
| 8:23.96 | כבד |  |
| 8:23.97 | את |  |
| 8:23.98 | אביך |  |
| 8:23.99 | ואת |  |
| 8:23.100 | אמך | , |
| 8:23.101 | את |  |
| 8:23.102 | אביך |  |
| 8:23.103 | זו |  |
| 8:23.104 | אשת |  |
| 8:23.105 | אביך | , |
| 8:23.106 | ואת |  |
| 8:23.107 | אמך |  |
| 8:23.108 | זה |  |
| 8:23.109 | בעל |  |
| 8:23.110 | אמך | , |
| 8:23.111 | שיהא |  |
| 8:23.112 | מכבד |  |
| 8:23.113 | את |  |
| 8:23.114 | אשת |  |
| 8:23.115 | אביו |  |
| 8:23.116 | עם |  |
| 8:23.117 | אביו | , |
| 8:23.118 | ואת |  |
| 8:23.119 | בעל |  |
| 8:23.120 | אמו |  |
| 8:23.121 | עם |  |
| 8:23.122 | אמו | , |
| 8:23.123 | וכן |  |
| 8:23.124 | מה |  |
| 8:23.125 | שדרשו |  |
| 8:23.126 | רז"ל |  |
| 8:23.127 | וערלתם |  |
| 8:23.128 | ערלתו |  |
| 8:23.129 | את |  |
| 8:23.130 | פריו |  |
| 8:23.131 | את |  |
| 8:23.132 | הטפל |  |
| 8:23.133 | לפריו | , |
| 8:23.134 | דהיינו |  |
| 8:23.135 | הקליפין |  |
| 8:23.136 | והגרעינין |  |
| 8:23.137 | שהן |  |
| 8:23.138 | אסורין |  |
| 8:23.139 | בערלה |  |
| 8:23.140 | כפרי |  |
| 8:23.141 | עצמו |  |
| 8:23.142 | ובנדון |  |
| 8:23.143 | זה |  |
| 8:23.144 | עשו |  |
| 8:23.145 | המקובלים |  |
| 8:23.146 | החדשים |  |
| 8:23.147 | דבר |  |
| 8:23.148 | הטפל |  |
| 8:23.149 | הבא |  |
| 8:23.150 | מרבוי |  |
| 8:23.151 | את |  |
| 8:23.152 | לעיקר |  |
| 8:23.153 | ואמרו |  |
| 8:23.154 | שהם |  |
| 8:23.155 | עובדים |  |
| 8:23.156 | לנשמתו | , |
| 8:23.157 | והגוף |  |
| 8:23.158 | שהוא |  |
| 8:23.159 | נברא |  |
| 8:23.160 | טפל | . |
| 8:23.161 | ועובדים |  |
| 8:23.162 | וגוף |  |
| 8:23.163 | ולנשמה | , |
| 8:23.164 | כמו |  |
| 8:23.165 | שביארנו |  |
| 8:23.166 | לעיל |  |
| 8:23.167 | סי | ' |
| 8:23.168 | ט"ל |  |
| 8:23.169 | וסי | ' |
| 8:23.170 | ס"ה | ! |
| 8:23.171 | ונמצאו |  |
| 8:23.172 | משתפים |  |
| 8:23.173 | שם |  |
| 8:23.174 | שמים |  |
| 8:23.175 | ודבר |  |
| 8:23.176 | נברא | , |
| 8:23.177 | ועוברים |  |
| 8:23.178 | על |  |
| 8:23.179 | דברי |  |
| 8:23.180 | רשב"י |  |
| 8:23.181 | שאמר |  |
| 8:23.182 | כל |  |
| 8:23.183 | המשתף |  |
| 8:23.184 | שם |  |
| 8:23.185 | שמים |  |
| 8:23.186 | ודבר |  |
| 8:23.187 | אחר |  |
| 8:23.188 | נעקר |  |
| 8:23.189 | מן |  |
| 8:23.190 | העולם | . |
| 8:23.191 | והן |  |
| 8:23.192 | עתה |  |
| 8:23.193 | אוסיף |  |
| 8:23.194 | להעירך |  |
| 8:23.195 | על |  |
| 8:23.196 | קצת |  |
| 8:23.197 | מאמרים |  |
| 8:23.198 | מהזהר |  |
| 8:23.199 | דמוכח |  |
| 8:23.200 | מנייהו |  |
| 8:23.201 | שזעיר |  |
| 8:23.202 | אנפין |  |
| 8:23.203 | נברא |  |
| 8:23.204 | עיין |  |
| 8:23.205 | בזהר |  |
| 8:23.206 | בראית |  |
| 8:23.207 | דף |  |
| 8:23.208 | ט"ו |  |
| 8:23.209 | ומקדש |  |
| 8:23.210 | מלך |  |
| 8:23.211 | שם |  |
| 8:23.212 | מבואר |  |
| 8:23.213 | דז"א |  |
| 8:23.214 | נברא |  |
| 8:23.215 | הוא |  |
| 8:23.216 | וכבר |  |
| 8:23.217 | הבאתי |  |
| 8:23.218 | קצת |  |
| 8:23.219 | מזה |  |
| 8:23.220 | בסי | ' |
| 8:23.221 | כ"ד | , |
| 8:23.222 | גם |  |
| 8:23.223 | בדף |  |
| 8:23.224 | כ"ו |  |
| 8:23.225 | ע"א |  |
| 8:23.226 | ויטע |  |
| 8:23.227 | ה | ' |
| 8:23.228 | אלהים | , |
| 8:23.229 | אבא |  |
| 8:23.230 | ואימא | , |
| 8:23.231 | גן |  |
| 8:23.232 | דא |  |
| 8:23.233 | שכינתא |  |
| 8:23.234 | תתאה | , |
| 8:23.235 | עדן |  |
| 8:23.236 | דא |  |
| 8:23.237 | אימא |  |
| 8:23.238 | עלאה | , |
| 8:23.239 | את |  |
| 8:23.240 | האדם |  |
| 8:23.241 | דא |  |
| 8:23.242 | עמודא |  |
| 8:23.243 | דאמצעיתא |  |
| 8:23.244 | שהוא |  |
| 8:23.245 | ז"א | ) |
| 8:23.246 | עוד |  |
| 8:23.247 | שם |  |
| 8:23.248 | ויצמח |  |
| 8:23.249 | ה | ' |
| 8:23.250 | אלהים |  |
| 8:23.251 | אבא |  |
| 8:23.252 | ואימא | , |
| 8:23.253 | כל |  |
| 8:23.254 | עץ |  |
| 8:23.255 | נחמד |  |
| 8:23.256 | דא |  |
| 8:23.257 | צדיק |  |
| 8:23.258 | יסוד | ) |
| 8:23.259 | וטוב |  |
| 8:23.260 | למאכל | , |
| 8:23.261 | דא |  |
| 8:23.262 | עמודא |  |
| 8:23.263 | דאמצעיתא |  |
| 8:23.264 | וכו | ' |
| 8:23.265 | ועץ |  |
| 8:23.266 | החיים |  |
| 8:23.267 | דהוא |  |
| 8:23.268 | אילנא |  |
| 8:23.269 | דחי |  |
| 8:23.270 | יהי |  |
| 8:23.271 | נטיע |  |
| 8:23.272 | בנו |  |
| 8:23.273 | גינתא |  |
| 8:23.274 | וכו | ', |
| 8:23.275 | והראיה |  |
| 8:23.276 | היותר |  |
| 8:23.277 | ברורא |  |
| 8:23.278 | הוא |  |
| 8:23.279 | האמור |  |
| 8:23.280 | באדרא |  |
| 8:23.281 | רבה |  |
| 8:23.282 | דף |  |
| 8:23.283 | קל"ח |  |
| 8:23.284 | ע"ב | , |
| 8:23.285 | ה | ' |
| 8:23.286 | פעלך |  |
| 8:23.287 | בקרב |  |
| 8:23.288 | שנים |  |
| 8:23.289 | חייהו | . |
| 8:23.290 | האי |  |
| 8:23.291 | לעתיק |  |
| 8:23.292 | יומין |  |
| 8:23.293 | אתאמר | , |
| 8:23.294 | מאן |  |
| 8:23.295 | פעלך | , |
| 8:23.296 | זעיר |  |
| 8:23.297 | אנפין | , |
| 8:23.298 | חייהו |  |
| 8:23.299 | למאן | ? |
| 8:23.300 | לזעיר |  |
| 8:23.301 | אנפין |  |
| 8:23.302 | דכל |  |
| 8:23.303 | נהירו |  |
| 8:23.304 | דיליה |  |
| 8:23.305 | מאינון |  |
| 8:23.306 | שנים |  |
| 8:23.307 | קדמוניות | ! |
| 8:23.308 | מזה |  |
| 8:23.309 | יתבאר |  |
| 8:23.310 | לך |  |
| 8:23.311 | בבירור |  |
| 8:23.312 | שזעיר |  |
| 8:23.313 | אנפין |  |
| 8:23.314 | נברא | , |
| 8:23.315 | וכבר |  |
| 8:23.316 | ידעת |  |
| 8:23.317 | דעת |  |
| 8:23.318 | הפילסוף |  |
| 8:23.319 | המסית |  |
| 8:23.320 | מחבר |  |
| 8:23.321 | הזהר |  |
| 8:23.322 | ורוב |  |
| 8:23.323 | המקובלים |  |
| 8:23.324 | הנמשכים |  |
| 8:23.325 | אחריו | , |
| 8:23.326 | שלא |  |
| 8:23.327 | תכון |  |
| 8:23.328 | תפלה |  |
| 8:23.329 | ועבודה |  |
| 8:23.330 | רק |  |
| 8:23.331 | לזעיר |  |
| 8:23.332 | אנפין |  |
| 8:23.333 | הנברא |  |
| 8:23.334 | הלזה | , |
| 8:23.335 | ולדעת |  |
| 8:23.336 | בעל |  |
| 8:23.337 | עז |  |
| 8:23.338 | לאלהים |  |
| 8:23.339 | אין |  |
| 8:23.340 | העבודה |  |
| 8:23.341 | אלא |  |
| 8:23.342 | למלכא |  |
| 8:23.343 | קדישא |  |
| 8:23.344 | דכל |  |
| 8:23.345 | קדישין |  |
| 8:23.346 | הנקרא |  |
| 8:23.347 | לבא |  |
| 8:23.348 | דכל |  |
| 8:23.349 | ליבין |  |
| 8:23.350 | והוא |  |
| 8:23.351 | מתלבש |  |
| 8:23.352 | תוך |  |
| 8:23.353 | זעיר |  |
| 8:23.354 | אנפין |  |
| 8:23.355 | כמבואר |  |
| 8:23.356 | שם |  |
| 8:23.357 | בבית |  |
| 8:23.358 | קדש |  |
| 8:23.359 | הקדשים |  |
| 8:23.360 | פרק |  |
| 8:23.361 | י"ט |  |
| 8:23.362 | דף |  |
| 8:23.363 | נ"ט |  |
| 8:23.364 | ע"ד | , |
| 8:23.365 | ואע"פ |  |
| 8:23.366 | שיצא |  |
| 8:23.367 | ממעי |  |
| 8:23.368 | אבא |  |
| 8:23.369 | ואימא |  |
| 8:23.370 | הוא |  |
| 8:23.371 | העיקר |  |
| 8:23.372 | ואף |  |
| 8:23.373 | אבא |  |
| 8:23.374 | ואימא |  |
| 8:23.375 | צריכים |  |
| 8:23.376 | לכבדו |  |
| 8:23.377 | כמלך |  |
| 8:23.378 | בשר |  |
| 8:23.379 | ודם |  |
| 8:23.380 | שמולך |  |
| 8:23.381 | שכל |  |
| 8:23.382 | העם |  |
| 8:23.383 | מכבדים |  |
| 8:23.384 | אותו | . |
| 8:23.385 | ואף |  |
| 8:23.386 | אביו |  |
| 8:23.387 | ואמו |  |
| 8:23.388 | צריכים |  |
| 8:23.389 | לכבדו | , |
| 8:23.390 | שהוא |  |
| 8:23.391 | המלך |  |
| 8:23.392 | השליט | , |
| 8:23.393 | ואליו |  |
| 8:23.394 | תפלותינו |  |
| 8:23.395 | והוא |  |
| 8:23.396 | מלכינו |  |
| 8:23.397 | ואלהינו | , |
| 8:23.398 | אבל |  |
| 8:23.399 | המתפלל |  |
| 8:23.400 | לאל |  |
| 8:23.401 | עליון |  |
| 8:23.402 | שהוא |  |
| 8:23.403 | הסבה |  |
| 8:23.404 | הראשונה |  |
| 8:23.405 | אין |  |
| 8:23.406 | תפלתו |  |
| 8:23.407 | תפלה | , |
| 8:23.408 | כי |  |
| 8:23.409 | איך |  |
| 8:23.410 | נאמר |  |
| 8:23.411 | ברוך |  |
| 8:23.412 | אתה |  |
| 8:23.413 | ה | ', |
| 8:23.414 | ממי |  |
| 8:23.415 | יתברך |  |
| 8:23.416 | ויבוא |  |
| 8:23.417 | לו |  |
| 8:23.418 | השפע | ? |
| 8:23.419 | בית |  |
| 8:23.420 | קדש |  |
| 8:23.421 | הקדשים |  |
| 8:23.422 | דף |  |
| 8:23.423 | כ"ו |  |
| 8:23.424 | ע"ד | ). |
| 8:23.425 | ועיין |  |
| 8:23.426 | בזהר |  |
| 8:23.427 | רעיא |  |
| 8:23.428 | מהימנא |  |
| 8:23.429 | דף |  |
| 8:23.430 | ק"ט |  |
| 8:23.431 | ע"ב |  |
| 8:23.432 | שמבאר |  |
| 8:23.433 | שם | , |
| 8:23.434 | כי |  |
| 8:23.435 | האלהים |  |
| 8:23.436 | העליון |  |
| 8:23.437 | גלית |  |
| 8:23.438 | ליה |  |
| 8:23.439 | גופא |  |
| 8:23.440 | ולא |  |
| 8:23.441 | איברים |  |
| 8:23.442 | ולא |  |
| 8:23.443 | נוקבא | , |
| 8:23.444 | והוא |  |
| 8:23.445 | רק |  |
| 8:23.446 | מחבר |  |
| 8:23.447 | ומקשר |  |
| 8:23.448 | את |  |
| 8:23.449 | האלהים |  |
| 8:23.450 | התחתונים |  |
| 8:23.451 | בעלי |  |
| 8:23.452 | גוף |  |
| 8:23.453 | ואיברים |  |
| 8:23.454 | זכרים |  |
| 8:23.455 | ונקיבות |  |
| 8:23.456 | עם |  |
| 8:23.457 | נקבותיהם |  |
| 8:23.458 | להיות |  |
| 8:23.459 | אחד | , |
| 8:23.460 | כטיט |  |
| 8:23.461 | המחבר |  |
| 8:23.462 | אבני |  |
| 8:23.463 | הבית | , |
| 8:23.464 | או |  |
| 8:23.465 | כחשוקי |  |
| 8:23.466 | ברזל |  |
| 8:23.467 | ונחשת |  |
| 8:23.468 | הקושרים |  |
| 8:23.469 | ומחברים |  |
| 8:23.470 | את |  |
| 8:23.471 | לוחות |  |
| 8:23.472 | חביות |  |
| 8:23.473 | של |  |
| 8:23.474 | עץ |  |
| 8:23.475 | להיות |  |
| 8:23.476 | אחד | . |
| 8:23.477 | יתעלה |  |
| 8:23.478 | השי"ת |  |
| 8:23.479 | ויתרומם |  |
| 8:23.480 | מכל |  |
| 8:23.481 | מה |  |
| 8:23.482 | שחשבו |  |
| 8:23.483 | בו |  |
| 8:23.484 | הטפשים |  |
| 8:23.485 | והסכלים | ! |
| 8:23.486 | ולכן |  |
| 8:23.487 | אמרו |  |
| 8:23.488 | המקובלים |  |
| 8:23.489 | החדשים |  |
| 8:23.490 | שהמתפלל |  |
| 8:23.491 | אליו |  |
| 8:23.492 | אין |  |
| 8:23.493 | תפלתו |  |
| 8:23.494 | תפלה | , |
| 8:23.495 | וכל |  |
| 8:23.496 | הקורא |  |
| 8:23.497 | אליו |  |
| 8:23.498 | לא |  |
| 8:23.499 | יענהו | , |
| 8:23.500 | והוא |  |
| 8:23.501 | מרומם |  |
| 8:23.502 | ומסולק |  |
| 8:23.503 | מכל |  |
| 8:23.504 | ברכה |  |
| 8:23.505 | שאין |  |
| 8:23.506 | לו |  |
| 8:23.507 | ממי |  |
| 8:23.508 | להתברך | , |
| 8:23.509 | ולא |  |
| 8:23.510 | שם |  |
| 8:23.511 | לו |  |
| 8:23.512 | שיגבילהו | , |
| 8:23.513 | כי |  |
| 8:23.514 | הכח |  |
| 8:23.515 | והממשלה |  |
| 8:23.516 | להאלהים |  |
| 8:23.517 | שלמטה |  |
| 8:23.518 | ממנו |  |
| 8:23.519 | שהם |  |
| 8:23.520 | החמשה |  |
| 8:23.521 | פרצופים |  |
| 8:23.522 | הכוללים | , |
| 8:23.523 | והמה |  |
| 8:23.524 | נקראים |  |
| 8:23.525 | בשם |  |
| 8:23.526 | הוי'ה |  |
| 8:23.527 | ואדנות |  |
| 8:23.528 | ואלהים | , |
| 8:23.529 | ועיקר |  |
| 8:23.530 | שם |  |
| 8:23.531 | המיוחד |  |
| 8:23.532 | שם |  |
| 8:23.533 | הוי'ה |  |
| 8:23.534 | נקרא |  |
| 8:23.535 | להאלוה |  |
| 8:23.536 | קצר |  |
| 8:23.537 | אפים |  |
| 8:23.538 | זעיר |  |
| 8:23.539 | אנפין | ) |
| 8:23.540 | עם |  |
| 8:23.541 | זוגתו | , |
| 8:23.542 | והוא |  |
| 8:23.543 | השולט |  |
| 8:23.544 | על |  |
| 8:23.545 | כל |  |
| 8:23.546 | הברואים |  |
| 8:23.547 | כמו |  |
| 8:23.548 | שנת | ' |
| 8:23.549 | לעיל |  |
| 8:23.550 | בכמה |  |
| 8:23.551 | מקומות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

