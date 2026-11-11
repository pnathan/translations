# Translation pack: milhamot-hashem perek-8-13 (Milchamot ha-Shem 8:16–8:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-13`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
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


## Milchamot ha-Shem 8:16

Hebrew (Wikisource): קי) ולעניות דעתי יראה בפי' המאמר בהקדים היסוד שיסדו לנו רז"ל במאמרם על ולא תתורו אחרי לבבכם ואחרי עיניכם. אחרי לבבכם זו מונות, ואחרי עיניכם זו זנות. כי לבא ועינא תרי סרסורי דחטאה נינהו לחכמים בעיניהם ההולכים אחרי דעתם שלא על דרך רז"ל בפשטי הכתובים ודרשות ויסודות המקובלים בידם איש מפי איש עד משה רבינו ע"ה, וסיקלו המסילה לפנינו שלא לנטות ימין ושמאל מן הדרל הישרה אשר דרכו בה רבותינו בעלתי התלמוד והמדרשים האמתיים ואחריהם דרכו בזאת הדרך גם הרס"ג ובעל חובת הלבבות ורבי יהודה הלוי והרמב"ם ובעל העיקרים וס' מצות גדול, ורבים מחכמינו אנשים חכמים ונבונים. וידועים לנו בחכמתם ובינתם. עד אשר נתגלה ספר הזהר במרמת המלך העכו"ם שאמר שמצאו בעפר ושלחו לחכמי טוליטולה בעת אשר נתמעטה חכמת חכמינו ובינת נבונינו הסתתרה, וחשבו הנמנעות לנמצאות. וחידות רבותינו בעלי התלמוד והמדרשים ומשליהם שהם כתפוחי זהב במשכיות כסף האמינו בפירושן ככל העולה על רוחם. על פי הטעאת הזהר אשר יהסוהו לרשב"י בשקר. ובלי מדע ובינה והשכל האמינו לו לחשוב קבלתו המזויפת לקבלה אמתית. ויעשוה לגבירה ובת מלך, ואת הקבלה האמתית המשנה והתלמוד שמו לשפחה חרופה. ויהרפוה ויכניעוה ויכריעוה תחת ממשלתה לרדותה בחזק יד ובפרך. ולקבוע מנהגים ודינים חדשים שלא כרצונה. והתירו אז פרושים את הדבר לחשוב באלהות דברים אשר לא כן על ה' אלהיהם לעשות לו דמות ותבנית ותמונה! ושלשלו מהאלהים העליון כמה פרצופים. וקראו להם בשם הוי'ה ואדנות ושאר השמות המיוחדים כאשר עשו דורו של אנוש שקראו לעצבים בשם ה' ונסמכו על משענת קנה רצוץ המיוחס לרשב"י בשקר שהמציא שני מיני אלוהות הנמשכים ומשתלשלים מהסבה הראשונה. האחד מהם מערכת אלהים קדושים פועלים הטוב והשני מערכת אלהים טמאים סטרא אחרא מסאבא פועלים הרע. וישם את הכל הקדושים והטמאים בסתרו של עולם. ואת ה' אלהים אמת האל הגדול הגבור והנורא ארך אפים ורב חסד וכו' סלקו למעלה, ואמרו שלא יאות לשבחו ולפארו ולהתפלל אליו כי הוא מסולק ומרומם מכל ברכה ותהלה. ולא יענה את הקורא ומשוע אליו, כמו שטעו דור אנוש אשר היו כעורים. מגמת פניהם אל המלך ותעו מן הדרך וחשבו כי לחפצם נגעו, ובא להם זה מחסרון חכמה ומדע. כמו שאמרו רז"ל המובא לעיל, והרמב"ם ז"ל בהל' ע"ז שזהו טעותם של דור אנוש שאמרו שלערך יקרתו ומעלתו של אלהינו חלק ממשלה לזולתו ( ובס' עז לאלהים כתב בבית קדש הקדשים שנקרא מלכא קדישא. קודשא בריך הוא על שם שהוא מתברך ממה שלמעלה ממנו מה שאין כן האין סוף אין לו ממי להתברך על כן לא יצדק בו שם קודשא בריך הוא) ואינו משגיח בתחתונים, ולא יפנה אל שבחי ותהלות המשבחים ומפארים אותו, ושגם לא נקרא לו שם כלל, וכל שבח ותהלה ותפלה יחסו, אל אלוה קצר אפים (זעיר אנפין) ואמרו שהוא המלך המהולל בכל התשבוחות, כמו שכתב בביאור הרב ספר הברית והרב יושר לבב ומהר"י לופיס ונחלת יוסף ורבים מהמקובלים החדשים שנמשכו אחר דעת המסית מחבר הזהר.

| id | word | punctuation after |
|---|---|---|
| 8:16.1 | קי | ) |
| 8:16.2 | ולעניות |  |
| 8:16.3 | דעתי |  |
| 8:16.4 | יראה |  |
| 8:16.5 | בפי | ' |
| 8:16.6 | המאמר |  |
| 8:16.7 | בהקדים |  |
| 8:16.8 | היסוד |  |
| 8:16.9 | שיסדו |  |
| 8:16.10 | לנו |  |
| 8:16.11 | רז"ל |  |
| 8:16.12 | במאמרם |  |
| 8:16.13 | על |  |
| 8:16.14 | ולא |  |
| 8:16.15 | תתורו |  |
| 8:16.16 | אחרי |  |
| 8:16.17 | לבבכם |  |
| 8:16.18 | ואחרי |  |
| 8:16.19 | עיניכם | . |
| 8:16.20 | אחרי |  |
| 8:16.21 | לבבכם |  |
| 8:16.22 | זו |  |
| 8:16.23 | מונות | , |
| 8:16.24 | ואחרי |  |
| 8:16.25 | עיניכם |  |
| 8:16.26 | זו |  |
| 8:16.27 | זנות | . |
| 8:16.28 | כי |  |
| 8:16.29 | לבא |  |
| 8:16.30 | ועינא |  |
| 8:16.31 | תרי |  |
| 8:16.32 | סרסורי |  |
| 8:16.33 | דחטאה |  |
| 8:16.34 | נינהו |  |
| 8:16.35 | לחכמים |  |
| 8:16.36 | בעיניהם |  |
| 8:16.37 | ההולכים |  |
| 8:16.38 | אחרי |  |
| 8:16.39 | דעתם |  |
| 8:16.40 | שלא |  |
| 8:16.41 | על |  |
| 8:16.42 | דרך |  |
| 8:16.43 | רז"ל |  |
| 8:16.44 | בפשטי |  |
| 8:16.45 | הכתובים |  |
| 8:16.46 | ודרשות |  |
| 8:16.47 | ויסודות |  |
| 8:16.48 | המקובלים |  |
| 8:16.49 | בידם |  |
| 8:16.50 | איש |  |
| 8:16.51 | מפי |  |
| 8:16.52 | איש |  |
| 8:16.53 | עד |  |
| 8:16.54 | משה |  |
| 8:16.55 | רבינו |  |
| 8:16.56 | ע"ה | , |
| 8:16.57 | וסיקלו |  |
| 8:16.58 | המסילה |  |
| 8:16.59 | לפנינו |  |
| 8:16.60 | שלא |  |
| 8:16.61 | לנטות |  |
| 8:16.62 | ימין |  |
| 8:16.63 | ושמאל |  |
| 8:16.64 | מן |  |
| 8:16.65 | הדרל |  |
| 8:16.66 | הישרה |  |
| 8:16.67 | אשר |  |
| 8:16.68 | דרכו |  |
| 8:16.69 | בה |  |
| 8:16.70 | רבותינו |  |
| 8:16.71 | בעלתי |  |
| 8:16.72 | התלמוד |  |
| 8:16.73 | והמדרשים |  |
| 8:16.74 | האמתיים |  |
| 8:16.75 | ואחריהם |  |
| 8:16.76 | דרכו |  |
| 8:16.77 | בזאת |  |
| 8:16.78 | הדרך |  |
| 8:16.79 | גם |  |
| 8:16.80 | הרס"ג |  |
| 8:16.81 | ובעל |  |
| 8:16.82 | חובת |  |
| 8:16.83 | הלבבות |  |
| 8:16.84 | ורבי |  |
| 8:16.85 | יהודה |  |
| 8:16.86 | הלוי |  |
| 8:16.87 | והרמב"ם |  |
| 8:16.88 | ובעל |  |
| 8:16.89 | העיקרים |  |
| 8:16.90 | וס | ' |
| 8:16.91 | מצות |  |
| 8:16.92 | גדול | , |
| 8:16.93 | ורבים |  |
| 8:16.94 | מחכמינו |  |
| 8:16.95 | אנשים |  |
| 8:16.96 | חכמים |  |
| 8:16.97 | ונבונים | . |
| 8:16.98 | וידועים |  |
| 8:16.99 | לנו |  |
| 8:16.100 | בחכמתם |  |
| 8:16.101 | ובינתם | . |
| 8:16.102 | עד |  |
| 8:16.103 | אשר |  |
| 8:16.104 | נתגלה |  |
| 8:16.105 | ספר |  |
| 8:16.106 | הזהר |  |
| 8:16.107 | במרמת |  |
| 8:16.108 | המלך |  |
| 8:16.109 | העכו"ם |  |
| 8:16.110 | שאמר |  |
| 8:16.111 | שמצאו |  |
| 8:16.112 | בעפר |  |
| 8:16.113 | ושלחו |  |
| 8:16.114 | לחכמי |  |
| 8:16.115 | טוליטולה |  |
| 8:16.116 | בעת |  |
| 8:16.117 | אשר |  |
| 8:16.118 | נתמעטה |  |
| 8:16.119 | חכמת |  |
| 8:16.120 | חכמינו |  |
| 8:16.121 | ובינת |  |
| 8:16.122 | נבונינו |  |
| 8:16.123 | הסתתרה | , |
| 8:16.124 | וחשבו |  |
| 8:16.125 | הנמנעות |  |
| 8:16.126 | לנמצאות | . |
| 8:16.127 | וחידות |  |
| 8:16.128 | רבותינו |  |
| 8:16.129 | בעלי |  |
| 8:16.130 | התלמוד |  |
| 8:16.131 | והמדרשים |  |
| 8:16.132 | ומשליהם |  |
| 8:16.133 | שהם |  |
| 8:16.134 | כתפוחי |  |
| 8:16.135 | זהב |  |
| 8:16.136 | במשכיות |  |
| 8:16.137 | כסף |  |
| 8:16.138 | האמינו |  |
| 8:16.139 | בפירושן |  |
| 8:16.140 | ככל |  |
| 8:16.141 | העולה |  |
| 8:16.142 | על |  |
| 8:16.143 | רוחם | . |
| 8:16.144 | על |  |
| 8:16.145 | פי |  |
| 8:16.146 | הטעאת |  |
| 8:16.147 | הזהר |  |
| 8:16.148 | אשר |  |
| 8:16.149 | יהסוהו |  |
| 8:16.150 | לרשב"י |  |
| 8:16.151 | בשקר | . |
| 8:16.152 | ובלי |  |
| 8:16.153 | מדע |  |
| 8:16.154 | ובינה |  |
| 8:16.155 | והשכל |  |
| 8:16.156 | האמינו |  |
| 8:16.157 | לו |  |
| 8:16.158 | לחשוב |  |
| 8:16.159 | קבלתו |  |
| 8:16.160 | המזויפת |  |
| 8:16.161 | לקבלה |  |
| 8:16.162 | אמתית | . |
| 8:16.163 | ויעשוה |  |
| 8:16.164 | לגבירה |  |
| 8:16.165 | ובת |  |
| 8:16.166 | מלך | , |
| 8:16.167 | ואת |  |
| 8:16.168 | הקבלה |  |
| 8:16.169 | האמתית |  |
| 8:16.170 | המשנה |  |
| 8:16.171 | והתלמוד |  |
| 8:16.172 | שמו |  |
| 8:16.173 | לשפחה |  |
| 8:16.174 | חרופה | . |
| 8:16.175 | ויהרפוה |  |
| 8:16.176 | ויכניעוה |  |
| 8:16.177 | ויכריעוה |  |
| 8:16.178 | תחת |  |
| 8:16.179 | ממשלתה |  |
| 8:16.180 | לרדותה |  |
| 8:16.181 | בחזק |  |
| 8:16.182 | יד |  |
| 8:16.183 | ובפרך | . |
| 8:16.184 | ולקבוע |  |
| 8:16.185 | מנהגים |  |
| 8:16.186 | ודינים |  |
| 8:16.187 | חדשים |  |
| 8:16.188 | שלא |  |
| 8:16.189 | כרצונה | . |
| 8:16.190 | והתירו |  |
| 8:16.191 | אז |  |
| 8:16.192 | פרושים |  |
| 8:16.193 | את |  |
| 8:16.194 | הדבר |  |
| 8:16.195 | לחשוב |  |
| 8:16.196 | באלהות |  |
| 8:16.197 | דברים |  |
| 8:16.198 | אשר |  |
| 8:16.199 | לא |  |
| 8:16.200 | כן |  |
| 8:16.201 | על |  |
| 8:16.202 | ה | ' |
| 8:16.203 | אלהיהם |  |
| 8:16.204 | לעשות |  |
| 8:16.205 | לו |  |
| 8:16.206 | דמות |  |
| 8:16.207 | ותבנית |  |
| 8:16.208 | ותמונה | ! |
| 8:16.209 | ושלשלו |  |
| 8:16.210 | מהאלהים |  |
| 8:16.211 | העליון |  |
| 8:16.212 | כמה |  |
| 8:16.213 | פרצופים | . |
| 8:16.214 | וקראו |  |
| 8:16.215 | להם |  |
| 8:16.216 | בשם |  |
| 8:16.217 | הוי'ה |  |
| 8:16.218 | ואדנות |  |
| 8:16.219 | ושאר |  |
| 8:16.220 | השמות |  |
| 8:16.221 | המיוחדים |  |
| 8:16.222 | כאשר |  |
| 8:16.223 | עשו |  |
| 8:16.224 | דורו |  |
| 8:16.225 | של |  |
| 8:16.226 | אנוש |  |
| 8:16.227 | שקראו |  |
| 8:16.228 | לעצבים |  |
| 8:16.229 | בשם |  |
| 8:16.230 | ה | ' |
| 8:16.231 | ונסמכו |  |
| 8:16.232 | על |  |
| 8:16.233 | משענת |  |
| 8:16.234 | קנה |  |
| 8:16.235 | רצוץ |  |
| 8:16.236 | המיוחס |  |
| 8:16.237 | לרשב"י |  |
| 8:16.238 | בשקר |  |
| 8:16.239 | שהמציא |  |
| 8:16.240 | שני |  |
| 8:16.241 | מיני |  |
| 8:16.242 | אלוהות |  |
| 8:16.243 | הנמשכים |  |
| 8:16.244 | ומשתלשלים |  |
| 8:16.245 | מהסבה |  |
| 8:16.246 | הראשונה | . |
| 8:16.247 | האחד |  |
| 8:16.248 | מהם |  |
| 8:16.249 | מערכת |  |
| 8:16.250 | אלהים |  |
| 8:16.251 | קדושים |  |
| 8:16.252 | פועלים |  |
| 8:16.253 | הטוב |  |
| 8:16.254 | והשני |  |
| 8:16.255 | מערכת |  |
| 8:16.256 | אלהים |  |
| 8:16.257 | טמאים |  |
| 8:16.258 | סטרא |  |
| 8:16.259 | אחרא |  |
| 8:16.260 | מסאבא |  |
| 8:16.261 | פועלים |  |
| 8:16.262 | הרע | . |
| 8:16.263 | וישם |  |
| 8:16.264 | את |  |
| 8:16.265 | הכל |  |
| 8:16.266 | הקדושים |  |
| 8:16.267 | והטמאים |  |
| 8:16.268 | בסתרו |  |
| 8:16.269 | של |  |
| 8:16.270 | עולם | . |
| 8:16.271 | ואת |  |
| 8:16.272 | ה | ' |
| 8:16.273 | אלהים |  |
| 8:16.274 | אמת |  |
| 8:16.275 | האל |  |
| 8:16.276 | הגדול |  |
| 8:16.277 | הגבור |  |
| 8:16.278 | והנורא |  |
| 8:16.279 | ארך |  |
| 8:16.280 | אפים |  |
| 8:16.281 | ורב |  |
| 8:16.282 | חסד |  |
| 8:16.283 | וכו | ' |
| 8:16.284 | סלקו |  |
| 8:16.285 | למעלה | , |
| 8:16.286 | ואמרו |  |
| 8:16.287 | שלא |  |
| 8:16.288 | יאות |  |
| 8:16.289 | לשבחו |  |
| 8:16.290 | ולפארו |  |
| 8:16.291 | ולהתפלל |  |
| 8:16.292 | אליו |  |
| 8:16.293 | כי |  |
| 8:16.294 | הוא |  |
| 8:16.295 | מסולק |  |
| 8:16.296 | ומרומם |  |
| 8:16.297 | מכל |  |
| 8:16.298 | ברכה |  |
| 8:16.299 | ותהלה | . |
| 8:16.300 | ולא |  |
| 8:16.301 | יענה |  |
| 8:16.302 | את |  |
| 8:16.303 | הקורא |  |
| 8:16.304 | ומשוע |  |
| 8:16.305 | אליו | , |
| 8:16.306 | כמו |  |
| 8:16.307 | שטעו |  |
| 8:16.308 | דור |  |
| 8:16.309 | אנוש |  |
| 8:16.310 | אשר |  |
| 8:16.311 | היו |  |
| 8:16.312 | כעורים | . |
| 8:16.313 | מגמת |  |
| 8:16.314 | פניהם |  |
| 8:16.315 | אל |  |
| 8:16.316 | המלך |  |
| 8:16.317 | ותעו |  |
| 8:16.318 | מן |  |
| 8:16.319 | הדרך |  |
| 8:16.320 | וחשבו |  |
| 8:16.321 | כי |  |
| 8:16.322 | לחפצם |  |
| 8:16.323 | נגעו | , |
| 8:16.324 | ובא |  |
| 8:16.325 | להם |  |
| 8:16.326 | זה |  |
| 8:16.327 | מחסרון |  |
| 8:16.328 | חכמה |  |
| 8:16.329 | ומדע | . |
| 8:16.330 | כמו |  |
| 8:16.331 | שאמרו |  |
| 8:16.332 | רז"ל |  |
| 8:16.333 | המובא |  |
| 8:16.334 | לעיל | , |
| 8:16.335 | והרמב"ם |  |
| 8:16.336 | ז"ל |  |
| 8:16.337 | בהל | ' |
| 8:16.338 | ע"ז |  |
| 8:16.339 | שזהו |  |
| 8:16.340 | טעותם |  |
| 8:16.341 | של |  |
| 8:16.342 | דור |  |
| 8:16.343 | אנוש |  |
| 8:16.344 | שאמרו |  |
| 8:16.345 | שלערך |  |
| 8:16.346 | יקרתו |  |
| 8:16.347 | ומעלתו |  |
| 8:16.348 | של |  |
| 8:16.349 | אלהינו |  |
| 8:16.350 | חלק |  |
| 8:16.351 | ממשלה |  |
| 8:16.352 | לזולתו | ( |
| 8:16.353 | ובס | ' |
| 8:16.354 | עז |  |
| 8:16.355 | לאלהים |  |
| 8:16.356 | כתב |  |
| 8:16.357 | בבית |  |
| 8:16.358 | קדש |  |
| 8:16.359 | הקדשים |  |
| 8:16.360 | שנקרא |  |
| 8:16.361 | מלכא |  |
| 8:16.362 | קדישא | . |
| 8:16.363 | קודשא |  |
| 8:16.364 | בריך |  |
| 8:16.365 | הוא |  |
| 8:16.366 | על |  |
| 8:16.367 | שם |  |
| 8:16.368 | שהוא |  |
| 8:16.369 | מתברך |  |
| 8:16.370 | ממה |  |
| 8:16.371 | שלמעלה |  |
| 8:16.372 | ממנו |  |
| 8:16.373 | מה |  |
| 8:16.374 | שאין |  |
| 8:16.375 | כן |  |
| 8:16.376 | האין |  |
| 8:16.377 | סוף |  |
| 8:16.378 | אין |  |
| 8:16.379 | לו |  |
| 8:16.380 | ממי |  |
| 8:16.381 | להתברך |  |
| 8:16.382 | על |  |
| 8:16.383 | כן |  |
| 8:16.384 | לא |  |
| 8:16.385 | יצדק |  |
| 8:16.386 | בו |  |
| 8:16.387 | שם |  |
| 8:16.388 | קודשא |  |
| 8:16.389 | בריך |  |
| 8:16.390 | הוא | ) |
| 8:16.391 | ואינו |  |
| 8:16.392 | משגיח |  |
| 8:16.393 | בתחתונים | , |
| 8:16.394 | ולא |  |
| 8:16.395 | יפנה |  |
| 8:16.396 | אל |  |
| 8:16.397 | שבחי |  |
| 8:16.398 | ותהלות |  |
| 8:16.399 | המשבחים |  |
| 8:16.400 | ומפארים |  |
| 8:16.401 | אותו | , |
| 8:16.402 | ושגם |  |
| 8:16.403 | לא |  |
| 8:16.404 | נקרא |  |
| 8:16.405 | לו |  |
| 8:16.406 | שם |  |
| 8:16.407 | כלל | , |
| 8:16.408 | וכל |  |
| 8:16.409 | שבח |  |
| 8:16.410 | ותהלה |  |
| 8:16.411 | ותפלה |  |
| 8:16.412 | יחסו | , |
| 8:16.413 | אל |  |
| 8:16.414 | אלוה |  |
| 8:16.415 | קצר |  |
| 8:16.416 | אפים |  |
| 8:16.417 | זעיר |  |
| 8:16.418 | אנפין | ) |
| 8:16.419 | ואמרו |  |
| 8:16.420 | שהוא |  |
| 8:16.421 | המלך |  |
| 8:16.422 | המהולל |  |
| 8:16.423 | בכל |  |
| 8:16.424 | התשבוחות | , |
| 8:16.425 | כמו |  |
| 8:16.426 | שכתב |  |
| 8:16.427 | בביאור |  |
| 8:16.428 | הרב |  |
| 8:16.429 | ספר |  |
| 8:16.430 | הברית |  |
| 8:16.431 | והרב |  |
| 8:16.432 | יושר |  |
| 8:16.433 | לבב |  |
| 8:16.434 | ומהר"י |  |
| 8:16.435 | לופיס |  |
| 8:16.436 | ונחלת |  |
| 8:16.437 | יוסף |  |
| 8:16.438 | ורבים |  |
| 8:16.439 | מהמקובלים |  |
| 8:16.440 | החדשים |  |
| 8:16.441 | שנמשכו |  |
| 8:16.442 | אחר |  |
| 8:16.443 | דעת |  |
| 8:16.444 | המסית |  |
| 8:16.445 | מחבר |  |
| 8:16.446 | הזהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

