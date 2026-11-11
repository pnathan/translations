# Translation pack: milhamot-hashem perek-6-10 (Milchamot ha-Shem 6:15–6:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-10`.

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
- **אצילות** and its root: fixed and translated, so that the verbs the
  author strings together stay linked. אצילות → "emanation"; עולם האצילות /
  דאצילות → "the World of Emanation"; נאצל → "emanated"; הנאצלים →
  "the emanated beings"; המאציל → "the Emanator". The four worlds:
  "Emanation, Creation, Formation, Making" (אבי"ע). "נאצלו ונבראו ונוצרו
  ונעשו" → "were emanated, created, formed and made". Keep נאצל and נברא
  distinct everywhere: the author argues that Ze'ir Anpin is "נברא ולא בורא"
  → "created, not creator".
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
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
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
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
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


## Milchamot ha-Shem 6:15

Hebrew (Wikisource): סז) הן עתה אערוך לך ידידי הקורא דברי הפילסוף הנז' המבזה לימוד המשנה והתלמוד, ומזלזל בשקלא וטריא דרבנן בתלמוד שהם לפי דעתו מהסטרא אחרא הנקרא מצרים (פועל הרע) שמעבידם ומטריחם בכמה קושיות ופירוקם כאשר יבין המשכיל בדבריו איך הוא משביח את הרמזים שלו בפסוק ונהר יוצא מעדן (זוהר בראשית דף כ"ו וכ"ז) בשם סבא דסבין (עתיק) שנגלה לר' אלעזר וחביריו. ור' אלעזר סיפר זה לאביו ר' שמעון שאמר להם סבא דסבין שאותן ארבעה שנכנסו לפרדס. השלשה מהם שלא למדו הסודות שלו שהם מאבני שיש טהור שהם מאבא וברתא והם מלוחות ראשונות שנשתברו, רק למדו משנה ותלמוד שהם מלוחות שניות שיש בהם טוב ורע והוא התר ואסור, לפיכך נסתכנו, אבל ר' עקיבא למד הסודות שהם ענייני רבוי האלוהות שהוא מרמזם באותיות שם הוי'ה ב"ה ושהסודות המבוארים בדבריו הם מאבני שיש טהור שכולו טוב והם ממדרגת לוחות ראשונות שנשתברו ואין בהם תערובת רע, ולפיכך נכנס בשלום ויצא בשלום, אבל המשנה והתלמוד ושאר מדרשי רז"ל הם מעורבים מטוב ורע. ראה ידידי ותן דעתך בדבריו איך הוא מבזה לימוד המשנה והתלמוד ומדרשי רז"ל ומעיד עליהם שהם עץ הדעת טוב ורע היפך דברי שלמה המלך ע"ה שאמר ברוח הקדש כי לקח טוב נתתי לכם תורתי אל תעזובו, ורז"ל במדרשיהם הרבו בשבח תורה שבכתב, ותורה שבעל פה, יותר כמו שכתבתי לעיל בש' הירושלמי בשם רשב"י.

| id | word | punctuation after |
|---|---|---|
| 6:15.1 | סז | ) |
| 6:15.2 | הן |  |
| 6:15.3 | עתה |  |
| 6:15.4 | אערוך |  |
| 6:15.5 | לך |  |
| 6:15.6 | ידידי |  |
| 6:15.7 | הקורא |  |
| 6:15.8 | דברי |  |
| 6:15.9 | הפילסוף |  |
| 6:15.10 | הנז | ' |
| 6:15.11 | המבזה |  |
| 6:15.12 | לימוד |  |
| 6:15.13 | המשנה |  |
| 6:15.14 | והתלמוד | , |
| 6:15.15 | ומזלזל |  |
| 6:15.16 | בשקלא |  |
| 6:15.17 | וטריא |  |
| 6:15.18 | דרבנן |  |
| 6:15.19 | בתלמוד |  |
| 6:15.20 | שהם |  |
| 6:15.21 | לפי |  |
| 6:15.22 | דעתו |  |
| 6:15.23 | מהסטרא |  |
| 6:15.24 | אחרא |  |
| 6:15.25 | הנקרא |  |
| 6:15.26 | מצרים |  |
| 6:15.27 | פועל |  |
| 6:15.28 | הרע | ) |
| 6:15.29 | שמעבידם |  |
| 6:15.30 | ומטריחם |  |
| 6:15.31 | בכמה |  |
| 6:15.32 | קושיות |  |
| 6:15.33 | ופירוקם |  |
| 6:15.34 | כאשר |  |
| 6:15.35 | יבין |  |
| 6:15.36 | המשכיל |  |
| 6:15.37 | בדבריו |  |
| 6:15.38 | איך |  |
| 6:15.39 | הוא |  |
| 6:15.40 | משביח |  |
| 6:15.41 | את |  |
| 6:15.42 | הרמזים |  |
| 6:15.43 | שלו |  |
| 6:15.44 | בפסוק |  |
| 6:15.45 | ונהר |  |
| 6:15.46 | יוצא |  |
| 6:15.47 | מעדן |  |
| 6:15.48 | זוהר |  |
| 6:15.49 | בראשית |  |
| 6:15.50 | דף |  |
| 6:15.51 | כ"ו |  |
| 6:15.52 | וכ"ז | ) |
| 6:15.53 | בשם |  |
| 6:15.54 | סבא |  |
| 6:15.55 | דסבין |  |
| 6:15.56 | עתיק | ) |
| 6:15.57 | שנגלה |  |
| 6:15.58 | לר | ' |
| 6:15.59 | אלעזר |  |
| 6:15.60 | וחביריו | . |
| 6:15.61 | ור | ' |
| 6:15.62 | אלעזר |  |
| 6:15.63 | סיפר |  |
| 6:15.64 | זה |  |
| 6:15.65 | לאביו |  |
| 6:15.66 | ר | ' |
| 6:15.67 | שמעון |  |
| 6:15.68 | שאמר |  |
| 6:15.69 | להם |  |
| 6:15.70 | סבא |  |
| 6:15.71 | דסבין |  |
| 6:15.72 | שאותן |  |
| 6:15.73 | ארבעה |  |
| 6:15.74 | שנכנסו |  |
| 6:15.75 | לפרדס | . |
| 6:15.76 | השלשה |  |
| 6:15.77 | מהם |  |
| 6:15.78 | שלא |  |
| 6:15.79 | למדו |  |
| 6:15.80 | הסודות |  |
| 6:15.81 | שלו |  |
| 6:15.82 | שהם |  |
| 6:15.83 | מאבני |  |
| 6:15.84 | שיש |  |
| 6:15.85 | טהור |  |
| 6:15.86 | שהם |  |
| 6:15.87 | מאבא |  |
| 6:15.88 | וברתא |  |
| 6:15.89 | והם |  |
| 6:15.90 | מלוחות |  |
| 6:15.91 | ראשונות |  |
| 6:15.92 | שנשתברו | , |
| 6:15.93 | רק |  |
| 6:15.94 | למדו |  |
| 6:15.95 | משנה |  |
| 6:15.96 | ותלמוד |  |
| 6:15.97 | שהם |  |
| 6:15.98 | מלוחות |  |
| 6:15.99 | שניות |  |
| 6:15.100 | שיש |  |
| 6:15.101 | בהם |  |
| 6:15.102 | טוב |  |
| 6:15.103 | ורע |  |
| 6:15.104 | והוא |  |
| 6:15.105 | התר |  |
| 6:15.106 | ואסור | , |
| 6:15.107 | לפיכך |  |
| 6:15.108 | נסתכנו | , |
| 6:15.109 | אבל |  |
| 6:15.110 | ר | ' |
| 6:15.111 | עקיבא |  |
| 6:15.112 | למד |  |
| 6:15.113 | הסודות |  |
| 6:15.114 | שהם |  |
| 6:15.115 | ענייני |  |
| 6:15.116 | רבוי |  |
| 6:15.117 | האלוהות |  |
| 6:15.118 | שהוא |  |
| 6:15.119 | מרמזם |  |
| 6:15.120 | באותיות |  |
| 6:15.121 | שם |  |
| 6:15.122 | הוי'ה |  |
| 6:15.123 | ב"ה |  |
| 6:15.124 | ושהסודות |  |
| 6:15.125 | המבוארים |  |
| 6:15.126 | בדבריו |  |
| 6:15.127 | הם |  |
| 6:15.128 | מאבני |  |
| 6:15.129 | שיש |  |
| 6:15.130 | טהור |  |
| 6:15.131 | שכולו |  |
| 6:15.132 | טוב |  |
| 6:15.133 | והם |  |
| 6:15.134 | ממדרגת |  |
| 6:15.135 | לוחות |  |
| 6:15.136 | ראשונות |  |
| 6:15.137 | שנשתברו |  |
| 6:15.138 | ואין |  |
| 6:15.139 | בהם |  |
| 6:15.140 | תערובת |  |
| 6:15.141 | רע | , |
| 6:15.142 | ולפיכך |  |
| 6:15.143 | נכנס |  |
| 6:15.144 | בשלום |  |
| 6:15.145 | ויצא |  |
| 6:15.146 | בשלום | , |
| 6:15.147 | אבל |  |
| 6:15.148 | המשנה |  |
| 6:15.149 | והתלמוד |  |
| 6:15.150 | ושאר |  |
| 6:15.151 | מדרשי |  |
| 6:15.152 | רז"ל |  |
| 6:15.153 | הם |  |
| 6:15.154 | מעורבים |  |
| 6:15.155 | מטוב |  |
| 6:15.156 | ורע | . |
| 6:15.157 | ראה |  |
| 6:15.158 | ידידי |  |
| 6:15.159 | ותן |  |
| 6:15.160 | דעתך |  |
| 6:15.161 | בדבריו |  |
| 6:15.162 | איך |  |
| 6:15.163 | הוא |  |
| 6:15.164 | מבזה |  |
| 6:15.165 | לימוד |  |
| 6:15.166 | המשנה |  |
| 6:15.167 | והתלמוד |  |
| 6:15.168 | ומדרשי |  |
| 6:15.169 | רז"ל |  |
| 6:15.170 | ומעיד |  |
| 6:15.171 | עליהם |  |
| 6:15.172 | שהם |  |
| 6:15.173 | עץ |  |
| 6:15.174 | הדעת |  |
| 6:15.175 | טוב |  |
| 6:15.176 | ורע |  |
| 6:15.177 | היפך |  |
| 6:15.178 | דברי |  |
| 6:15.179 | שלמה |  |
| 6:15.180 | המלך |  |
| 6:15.181 | ע"ה |  |
| 6:15.182 | שאמר |  |
| 6:15.183 | ברוח |  |
| 6:15.184 | הקדש |  |
| 6:15.185 | כי |  |
| 6:15.186 | לקח |  |
| 6:15.187 | טוב |  |
| 6:15.188 | נתתי |  |
| 6:15.189 | לכם |  |
| 6:15.190 | תורתי |  |
| 6:15.191 | אל |  |
| 6:15.192 | תעזובו | , |
| 6:15.193 | ורז"ל |  |
| 6:15.194 | במדרשיהם |  |
| 6:15.195 | הרבו |  |
| 6:15.196 | בשבח |  |
| 6:15.197 | תורה |  |
| 6:15.198 | שבכתב | , |
| 6:15.199 | ותורה |  |
| 6:15.200 | שבעל |  |
| 6:15.201 | פה | , |
| 6:15.202 | יותר |  |
| 6:15.203 | כמו |  |
| 6:15.204 | שכתבתי |  |
| 6:15.205 | לעיל |  |
| 6:15.206 | בש | ' |
| 6:15.207 | הירושלמי |  |
| 6:15.208 | בשם |  |
| 6:15.209 | רשב"י | . |


## Milchamot ha-Shem 6:16

Hebrew (Wikisource): ובכמה מקומות קורים רז"ל לתורה טוב כמו שאמרו על זנח ישראל טוב אויב ירדפו, ואין טוב אלא תורה שנ' כי לקח טוב נתתי לכם תורתי אל תעזובו. וכתיב עץ החיים היא למחזיקים בה, ולא עץ הדעת טוב ורע, ושוב חזר הפילסוף הנזכר לגנות את המשנה והתלמוד באמרו דבר אחר (שם דף כ"ז ע"א) ומשבח את סחורתו הסודות שהוא בודאם מלבו שהם מעץ החיים מזעיר אנפין דאצילות ששם אין שליטא לקליפות, ולא ימצא בהם קושיות ומשא ומתן של הלכה (אבל הם המצאת אלהים חדשים מקרוב באו לא שערום אבותינו והם עוקרים יסוד יחוד השם שהוא העיקר הגדול). אבל המשנה והתלמוד הם יוצאים ממטטרון דיצירה והוא זעיר אנפין דיצירה ששם נכנסו בן עזאי ובן זומא ואלישע, אחר ששם יש אחיזה לקליפות והם מעץ הדעת טוב ורע, מצד אחד טוב, והוא היתר כשר טהור, ומצד אחד רע והוא אסור ופסול וטמא (כאן גילה מצפוני לבו דניחא ליה בהפקירא ולא יזכר שם טמא ופסול ואסור החכם יבין). ועוד חזר שם לדרוש בגנות המשנה והתלמוד בשם מארי מתניתין וז"ל וימררו את חייהם (המצרים שהם הקליפות) בעבודה קשה, דא קושיא, בחומר, בקל וחומר, ובלבנים, בליבון הלכתא, ובכל עבודה בשדה, דא ברייתא, את כל עבודתם אשר עבדו בהם בפרך, דא משנה, אם תיבין בתיובתא. אתאמר בהון ויורהו ה' עץ, ודא עץ חיים, וביה ויתמתקו המים (פי' המקדש מלך בתחלה היו עומקים בפשטי התורה (דהיינו משנה וגמרא וברייתא שהזכיר לעיל) אזי וימררו את חייהם שאינם מקבלים שפע מזעיר דאצילות. ואחר ששבו בתשוּבה וחזרו ללמוד סודות התורה אז מושכים שפע משם מ"ה שהוא שקיו דאילנא שהוא פנימיות זעיר וכו') ודא משה משיח דאתמר ביה ומטה אלהים בידו, מטה דא מטטרון, מסטריה חיים, ומסטריה מיתה. כד אתהפיך למטה איהו עזר מסטרה דטוב. כך אתהפיך לחויא איהו כנגדו. מיד וינס משה מפניו. וקב"ה מסר ליה בידיה דמשה ואיהו אוריתא דבעל פה דביה איסור והתר (פי' המק"מ כי תורה שבע"פ היא במטטרון ששם טוב ורע שהם איסור והתר) הנך רואה בעיניך ידידי את דברי הפילסוף החובר מחבר הזהר איך הוא מכוין ליתן דופי במשנת חכמים ותלמודם המקובל בידם איש מפי איש ממשה רבינו ע"ה. וליחס את המשנה למטטרון שיש בו טוב ורע. ולא מפי הגבורה ניתנה אלא ממטטרון. והטוב שבה הוא התר טהור כשר, ורע שבה אסור וטמא ופסול. וחשב את העסק בה לעון אשר חטא. אשר אם ישובו ישראל בתשובה ויעזבו המשנה, ויתעסקו בסודות שערך לפניהם ברבוי האלוהות הקדושים לו פועלי הטוב, וגם בידיעת הקליפות אלהים אחרים טמאים פועלי הרע ימתקו להם המים המרים המאררים. כי אז לא ידעו עוד להבדיל בין טמא לטהור ובין קדש לחול, כאשר מצאנו ראינו האנשים העוסקים כל ימיהם בסודות הנזכרים ואינם משגיחים במשנה ותלמוד ופוסקים, שאינם יודעים כלל להבדיל בין הטמא ובין הטהור, כשר ופסול, אסור ומותר, ולא ידעו במה יכשלו, כי מוחם מלא מלי דחסידותא דאית בהו אפיקרותא, עבודת הפרצופים וכפירות שונים אשר לא צוה ה', ומוחין דקטנות ומוחין דגדלות, וזווגים וטפין הנוזלים מן הפרצופים הזכרים לרחם נשיהם ופילגשיהם, עד אפס מקום במוחם לנוח שם אסור ומותר. להבדיל בין הטמא ובין הטהור, ובין דבר הנאכל לדבר שאסור לאכלו אשר על זה נצטוינו בתורה, וקיימו בעצמם מאמר ר' נתן שאמר כי דבר ה' בזה זה שלא השגיח על המשנה (פי' רש"י שעשאה כפי שאינה עיקר. או שאינו מחזר לשנותה, סנהדרין דף צ"ט ע"ש ברש"י ומהרש"א).

| id | word | punctuation after |
|---|---|---|
| 6:16.1 | ובכמה |  |
| 6:16.2 | מקומות |  |
| 6:16.3 | קורים |  |
| 6:16.4 | רז"ל |  |
| 6:16.5 | לתורה |  |
| 6:16.6 | טוב |  |
| 6:16.7 | כמו |  |
| 6:16.8 | שאמרו |  |
| 6:16.9 | על |  |
| 6:16.10 | זנח |  |
| 6:16.11 | ישראל |  |
| 6:16.12 | טוב |  |
| 6:16.13 | אויב |  |
| 6:16.14 | ירדפו | , |
| 6:16.15 | ואין |  |
| 6:16.16 | טוב |  |
| 6:16.17 | אלא |  |
| 6:16.18 | תורה |  |
| 6:16.19 | שנ | ' |
| 6:16.20 | כי |  |
| 6:16.21 | לקח |  |
| 6:16.22 | טוב |  |
| 6:16.23 | נתתי |  |
| 6:16.24 | לכם |  |
| 6:16.25 | תורתי |  |
| 6:16.26 | אל |  |
| 6:16.27 | תעזובו | . |
| 6:16.28 | וכתיב |  |
| 6:16.29 | עץ |  |
| 6:16.30 | החיים |  |
| 6:16.31 | היא |  |
| 6:16.32 | למחזיקים |  |
| 6:16.33 | בה | , |
| 6:16.34 | ולא |  |
| 6:16.35 | עץ |  |
| 6:16.36 | הדעת |  |
| 6:16.37 | טוב |  |
| 6:16.38 | ורע | , |
| 6:16.39 | ושוב |  |
| 6:16.40 | חזר |  |
| 6:16.41 | הפילסוף |  |
| 6:16.42 | הנזכר |  |
| 6:16.43 | לגנות |  |
| 6:16.44 | את |  |
| 6:16.45 | המשנה |  |
| 6:16.46 | והתלמוד |  |
| 6:16.47 | באמרו |  |
| 6:16.48 | דבר |  |
| 6:16.49 | אחר |  |
| 6:16.50 | שם |  |
| 6:16.51 | דף |  |
| 6:16.52 | כ"ז |  |
| 6:16.53 | ע"א | ) |
| 6:16.54 | ומשבח |  |
| 6:16.55 | את |  |
| 6:16.56 | סחורתו |  |
| 6:16.57 | הסודות |  |
| 6:16.58 | שהוא |  |
| 6:16.59 | בודאם |  |
| 6:16.60 | מלבו |  |
| 6:16.61 | שהם |  |
| 6:16.62 | מעץ |  |
| 6:16.63 | החיים |  |
| 6:16.64 | מזעיר |  |
| 6:16.65 | אנפין |  |
| 6:16.66 | דאצילות |  |
| 6:16.67 | ששם |  |
| 6:16.68 | אין |  |
| 6:16.69 | שליטא |  |
| 6:16.70 | לקליפות | , |
| 6:16.71 | ולא |  |
| 6:16.72 | ימצא |  |
| 6:16.73 | בהם |  |
| 6:16.74 | קושיות |  |
| 6:16.75 | ומשא |  |
| 6:16.76 | ומתן |  |
| 6:16.77 | של |  |
| 6:16.78 | הלכה |  |
| 6:16.79 | אבל |  |
| 6:16.80 | הם |  |
| 6:16.81 | המצאת |  |
| 6:16.82 | אלהים |  |
| 6:16.83 | חדשים |  |
| 6:16.84 | מקרוב |  |
| 6:16.85 | באו |  |
| 6:16.86 | לא |  |
| 6:16.87 | שערום |  |
| 6:16.88 | אבותינו |  |
| 6:16.89 | והם |  |
| 6:16.90 | עוקרים |  |
| 6:16.91 | יסוד |  |
| 6:16.92 | יחוד |  |
| 6:16.93 | השם |  |
| 6:16.94 | שהוא |  |
| 6:16.95 | העיקר |  |
| 6:16.96 | הגדול | ). |
| 6:16.97 | אבל |  |
| 6:16.98 | המשנה |  |
| 6:16.99 | והתלמוד |  |
| 6:16.100 | הם |  |
| 6:16.101 | יוצאים |  |
| 6:16.102 | ממטטרון |  |
| 6:16.103 | דיצירה |  |
| 6:16.104 | והוא |  |
| 6:16.105 | זעיר |  |
| 6:16.106 | אנפין |  |
| 6:16.107 | דיצירה |  |
| 6:16.108 | ששם |  |
| 6:16.109 | נכנסו |  |
| 6:16.110 | בן |  |
| 6:16.111 | עזאי |  |
| 6:16.112 | ובן |  |
| 6:16.113 | זומא |  |
| 6:16.114 | ואלישע | , |
| 6:16.115 | אחר |  |
| 6:16.116 | ששם |  |
| 6:16.117 | יש |  |
| 6:16.118 | אחיזה |  |
| 6:16.119 | לקליפות |  |
| 6:16.120 | והם |  |
| 6:16.121 | מעץ |  |
| 6:16.122 | הדעת |  |
| 6:16.123 | טוב |  |
| 6:16.124 | ורע | , |
| 6:16.125 | מצד |  |
| 6:16.126 | אחד |  |
| 6:16.127 | טוב | , |
| 6:16.128 | והוא |  |
| 6:16.129 | היתר |  |
| 6:16.130 | כשר |  |
| 6:16.131 | טהור | , |
| 6:16.132 | ומצד |  |
| 6:16.133 | אחד |  |
| 6:16.134 | רע |  |
| 6:16.135 | והוא |  |
| 6:16.136 | אסור |  |
| 6:16.137 | ופסול |  |
| 6:16.138 | וטמא |  |
| 6:16.139 | כאן |  |
| 6:16.140 | גילה |  |
| 6:16.141 | מצפוני |  |
| 6:16.142 | לבו |  |
| 6:16.143 | דניחא |  |
| 6:16.144 | ליה |  |
| 6:16.145 | בהפקירא |  |
| 6:16.146 | ולא |  |
| 6:16.147 | יזכר |  |
| 6:16.148 | שם |  |
| 6:16.149 | טמא |  |
| 6:16.150 | ופסול |  |
| 6:16.151 | ואסור |  |
| 6:16.152 | החכם |  |
| 6:16.153 | יבין | ). |
| 6:16.154 | ועוד |  |
| 6:16.155 | חזר |  |
| 6:16.156 | שם |  |
| 6:16.157 | לדרוש |  |
| 6:16.158 | בגנות |  |
| 6:16.159 | המשנה |  |
| 6:16.160 | והתלמוד |  |
| 6:16.161 | בשם |  |
| 6:16.162 | מארי |  |
| 6:16.163 | מתניתין |  |
| 6:16.164 | וז"ל |  |
| 6:16.165 | וימררו |  |
| 6:16.166 | את |  |
| 6:16.167 | חייהם |  |
| 6:16.168 | המצרים |  |
| 6:16.169 | שהם |  |
| 6:16.170 | הקליפות | ) |
| 6:16.171 | בעבודה |  |
| 6:16.172 | קשה | , |
| 6:16.173 | דא |  |
| 6:16.174 | קושיא | , |
| 6:16.175 | בחומר | , |
| 6:16.176 | בקל |  |
| 6:16.177 | וחומר | , |
| 6:16.178 | ובלבנים | , |
| 6:16.179 | בליבון |  |
| 6:16.180 | הלכתא | , |
| 6:16.181 | ובכל |  |
| 6:16.182 | עבודה |  |
| 6:16.183 | בשדה | , |
| 6:16.184 | דא |  |
| 6:16.185 | ברייתא | , |
| 6:16.186 | את |  |
| 6:16.187 | כל |  |
| 6:16.188 | עבודתם |  |
| 6:16.189 | אשר |  |
| 6:16.190 | עבדו |  |
| 6:16.191 | בהם |  |
| 6:16.192 | בפרך | , |
| 6:16.193 | דא |  |
| 6:16.194 | משנה | , |
| 6:16.195 | אם |  |
| 6:16.196 | תיבין |  |
| 6:16.197 | בתיובתא | . |
| 6:16.198 | אתאמר |  |
| 6:16.199 | בהון |  |
| 6:16.200 | ויורהו |  |
| 6:16.201 | ה | ' |
| 6:16.202 | עץ | , |
| 6:16.203 | ודא |  |
| 6:16.204 | עץ |  |
| 6:16.205 | חיים | , |
| 6:16.206 | וביה |  |
| 6:16.207 | ויתמתקו |  |
| 6:16.208 | המים |  |
| 6:16.209 | פי | ' |
| 6:16.210 | המקדש |  |
| 6:16.211 | מלך |  |
| 6:16.212 | בתחלה |  |
| 6:16.213 | היו |  |
| 6:16.214 | עומקים |  |
| 6:16.215 | בפשטי |  |
| 6:16.216 | התורה |  |
| 6:16.217 | דהיינו |  |
| 6:16.218 | משנה |  |
| 6:16.219 | וגמרא |  |
| 6:16.220 | וברייתא |  |
| 6:16.221 | שהזכיר |  |
| 6:16.222 | לעיל | ) |
| 6:16.223 | אזי |  |
| 6:16.224 | וימררו |  |
| 6:16.225 | את |  |
| 6:16.226 | חייהם |  |
| 6:16.227 | שאינם |  |
| 6:16.228 | מקבלים |  |
| 6:16.229 | שפע |  |
| 6:16.230 | מזעיר |  |
| 6:16.231 | דאצילות | . |
| 6:16.232 | ואחר |  |
| 6:16.233 | ששבו |  |
| 6:16.234 | בתשוּבה |  |
| 6:16.235 | וחזרו |  |
| 6:16.236 | ללמוד |  |
| 6:16.237 | סודות |  |
| 6:16.238 | התורה |  |
| 6:16.239 | אז |  |
| 6:16.240 | מושכים |  |
| 6:16.241 | שפע |  |
| 6:16.242 | משם |  |
| 6:16.243 | מ"ה |  |
| 6:16.244 | שהוא |  |
| 6:16.245 | שקיו |  |
| 6:16.246 | דאילנא |  |
| 6:16.247 | שהוא |  |
| 6:16.248 | פנימיות |  |
| 6:16.249 | זעיר |  |
| 6:16.250 | וכו | ') |
| 6:16.251 | ודא |  |
| 6:16.252 | משה |  |
| 6:16.253 | משיח |  |
| 6:16.254 | דאתמר |  |
| 6:16.255 | ביה |  |
| 6:16.256 | ומטה |  |
| 6:16.257 | אלהים |  |
| 6:16.258 | בידו | , |
| 6:16.259 | מטה |  |
| 6:16.260 | דא |  |
| 6:16.261 | מטטרון | , |
| 6:16.262 | מסטריה |  |
| 6:16.263 | חיים | , |
| 6:16.264 | ומסטריה |  |
| 6:16.265 | מיתה | . |
| 6:16.266 | כד |  |
| 6:16.267 | אתהפיך |  |
| 6:16.268 | למטה |  |
| 6:16.269 | איהו |  |
| 6:16.270 | עזר |  |
| 6:16.271 | מסטרה |  |
| 6:16.272 | דטוב | . |
| 6:16.273 | כך |  |
| 6:16.274 | אתהפיך |  |
| 6:16.275 | לחויא |  |
| 6:16.276 | איהו |  |
| 6:16.277 | כנגדו | . |
| 6:16.278 | מיד |  |
| 6:16.279 | וינס |  |
| 6:16.280 | משה |  |
| 6:16.281 | מפניו | . |
| 6:16.282 | וקב"ה |  |
| 6:16.283 | מסר |  |
| 6:16.284 | ליה |  |
| 6:16.285 | בידיה |  |
| 6:16.286 | דמשה |  |
| 6:16.287 | ואיהו |  |
| 6:16.288 | אוריתא |  |
| 6:16.289 | דבעל |  |
| 6:16.290 | פה |  |
| 6:16.291 | דביה |  |
| 6:16.292 | איסור |  |
| 6:16.293 | והתר |  |
| 6:16.294 | פי | ' |
| 6:16.295 | המק"מ |  |
| 6:16.296 | כי |  |
| 6:16.297 | תורה |  |
| 6:16.298 | שבע"פ |  |
| 6:16.299 | היא |  |
| 6:16.300 | במטטרון |  |
| 6:16.301 | ששם |  |
| 6:16.302 | טוב |  |
| 6:16.303 | ורע |  |
| 6:16.304 | שהם |  |
| 6:16.305 | איסור |  |
| 6:16.306 | והתר | ) |
| 6:16.307 | הנך |  |
| 6:16.308 | רואה |  |
| 6:16.309 | בעיניך |  |
| 6:16.310 | ידידי |  |
| 6:16.311 | את |  |
| 6:16.312 | דברי |  |
| 6:16.313 | הפילסוף |  |
| 6:16.314 | החובר |  |
| 6:16.315 | מחבר |  |
| 6:16.316 | הזהר |  |
| 6:16.317 | איך |  |
| 6:16.318 | הוא |  |
| 6:16.319 | מכוין |  |
| 6:16.320 | ליתן |  |
| 6:16.321 | דופי |  |
| 6:16.322 | במשנת |  |
| 6:16.323 | חכמים |  |
| 6:16.324 | ותלמודם |  |
| 6:16.325 | המקובל |  |
| 6:16.326 | בידם |  |
| 6:16.327 | איש |  |
| 6:16.328 | מפי |  |
| 6:16.329 | איש |  |
| 6:16.330 | ממשה |  |
| 6:16.331 | רבינו |  |
| 6:16.332 | ע"ה | . |
| 6:16.333 | וליחס |  |
| 6:16.334 | את |  |
| 6:16.335 | המשנה |  |
| 6:16.336 | למטטרון |  |
| 6:16.337 | שיש |  |
| 6:16.338 | בו |  |
| 6:16.339 | טוב |  |
| 6:16.340 | ורע | . |
| 6:16.341 | ולא |  |
| 6:16.342 | מפי |  |
| 6:16.343 | הגבורה |  |
| 6:16.344 | ניתנה |  |
| 6:16.345 | אלא |  |
| 6:16.346 | ממטטרון | . |
| 6:16.347 | והטוב |  |
| 6:16.348 | שבה |  |
| 6:16.349 | הוא |  |
| 6:16.350 | התר |  |
| 6:16.351 | טהור |  |
| 6:16.352 | כשר | , |
| 6:16.353 | ורע |  |
| 6:16.354 | שבה |  |
| 6:16.355 | אסור |  |
| 6:16.356 | וטמא |  |
| 6:16.357 | ופסול | . |
| 6:16.358 | וחשב |  |
| 6:16.359 | את |  |
| 6:16.360 | העסק |  |
| 6:16.361 | בה |  |
| 6:16.362 | לעון |  |
| 6:16.363 | אשר |  |
| 6:16.364 | חטא | . |
| 6:16.365 | אשר |  |
| 6:16.366 | אם |  |
| 6:16.367 | ישובו |  |
| 6:16.368 | ישראל |  |
| 6:16.369 | בתשובה |  |
| 6:16.370 | ויעזבו |  |
| 6:16.371 | המשנה | , |
| 6:16.372 | ויתעסקו |  |
| 6:16.373 | בסודות |  |
| 6:16.374 | שערך |  |
| 6:16.375 | לפניהם |  |
| 6:16.376 | ברבוי |  |
| 6:16.377 | האלוהות |  |
| 6:16.378 | הקדושים |  |
| 6:16.379 | לו |  |
| 6:16.380 | פועלי |  |
| 6:16.381 | הטוב | , |
| 6:16.382 | וגם |  |
| 6:16.383 | בידיעת |  |
| 6:16.384 | הקליפות |  |
| 6:16.385 | אלהים |  |
| 6:16.386 | אחרים |  |
| 6:16.387 | טמאים |  |
| 6:16.388 | פועלי |  |
| 6:16.389 | הרע |  |
| 6:16.390 | ימתקו |  |
| 6:16.391 | להם |  |
| 6:16.392 | המים |  |
| 6:16.393 | המרים |  |
| 6:16.394 | המאררים | . |
| 6:16.395 | כי |  |
| 6:16.396 | אז |  |
| 6:16.397 | לא |  |
| 6:16.398 | ידעו |  |
| 6:16.399 | עוד |  |
| 6:16.400 | להבדיל |  |
| 6:16.401 | בין |  |
| 6:16.402 | טמא |  |
| 6:16.403 | לטהור |  |
| 6:16.404 | ובין |  |
| 6:16.405 | קדש |  |
| 6:16.406 | לחול | , |
| 6:16.407 | כאשר |  |
| 6:16.408 | מצאנו |  |
| 6:16.409 | ראינו |  |
| 6:16.410 | האנשים |  |
| 6:16.411 | העוסקים |  |
| 6:16.412 | כל |  |
| 6:16.413 | ימיהם |  |
| 6:16.414 | בסודות |  |
| 6:16.415 | הנזכרים |  |
| 6:16.416 | ואינם |  |
| 6:16.417 | משגיחים |  |
| 6:16.418 | במשנה |  |
| 6:16.419 | ותלמוד |  |
| 6:16.420 | ופוסקים | , |
| 6:16.421 | שאינם |  |
| 6:16.422 | יודעים |  |
| 6:16.423 | כלל |  |
| 6:16.424 | להבדיל |  |
| 6:16.425 | בין |  |
| 6:16.426 | הטמא |  |
| 6:16.427 | ובין |  |
| 6:16.428 | הטהור | , |
| 6:16.429 | כשר |  |
| 6:16.430 | ופסול | , |
| 6:16.431 | אסור |  |
| 6:16.432 | ומותר | , |
| 6:16.433 | ולא |  |
| 6:16.434 | ידעו |  |
| 6:16.435 | במה |  |
| 6:16.436 | יכשלו | , |
| 6:16.437 | כי |  |
| 6:16.438 | מוחם |  |
| 6:16.439 | מלא |  |
| 6:16.440 | מלי |  |
| 6:16.441 | דחסידותא |  |
| 6:16.442 | דאית |  |
| 6:16.443 | בהו |  |
| 6:16.444 | אפיקרותא | , |
| 6:16.445 | עבודת |  |
| 6:16.446 | הפרצופים |  |
| 6:16.447 | וכפירות |  |
| 6:16.448 | שונים |  |
| 6:16.449 | אשר |  |
| 6:16.450 | לא |  |
| 6:16.451 | צוה |  |
| 6:16.452 | ה | ', |
| 6:16.453 | ומוחין |  |
| 6:16.454 | דקטנות |  |
| 6:16.455 | ומוחין |  |
| 6:16.456 | דגדלות | , |
| 6:16.457 | וזווגים |  |
| 6:16.458 | וטפין |  |
| 6:16.459 | הנוזלים |  |
| 6:16.460 | מן |  |
| 6:16.461 | הפרצופים |  |
| 6:16.462 | הזכרים |  |
| 6:16.463 | לרחם |  |
| 6:16.464 | נשיהם |  |
| 6:16.465 | ופילגשיהם | , |
| 6:16.466 | עד |  |
| 6:16.467 | אפס |  |
| 6:16.468 | מקום |  |
| 6:16.469 | במוחם |  |
| 6:16.470 | לנוח |  |
| 6:16.471 | שם |  |
| 6:16.472 | אסור |  |
| 6:16.473 | ומותר | . |
| 6:16.474 | להבדיל |  |
| 6:16.475 | בין |  |
| 6:16.476 | הטמא |  |
| 6:16.477 | ובין |  |
| 6:16.478 | הטהור | , |
| 6:16.479 | ובין |  |
| 6:16.480 | דבר |  |
| 6:16.481 | הנאכל |  |
| 6:16.482 | לדבר |  |
| 6:16.483 | שאסור |  |
| 6:16.484 | לאכלו |  |
| 6:16.485 | אשר |  |
| 6:16.486 | על |  |
| 6:16.487 | זה |  |
| 6:16.488 | נצטוינו |  |
| 6:16.489 | בתורה | , |
| 6:16.490 | וקיימו |  |
| 6:16.491 | בעצמם |  |
| 6:16.492 | מאמר |  |
| 6:16.493 | ר | ' |
| 6:16.494 | נתן |  |
| 6:16.495 | שאמר |  |
| 6:16.496 | כי |  |
| 6:16.497 | דבר |  |
| 6:16.498 | ה | ' |
| 6:16.499 | בזה |  |
| 6:16.500 | זה |  |
| 6:16.501 | שלא |  |
| 6:16.502 | השגיח |  |
| 6:16.503 | על |  |
| 6:16.504 | המשנה |  |
| 6:16.505 | פי | ' |
| 6:16.506 | רש"י |  |
| 6:16.507 | שעשאה |  |
| 6:16.508 | כפי |  |
| 6:16.509 | שאינה |  |
| 6:16.510 | עיקר | . |
| 6:16.511 | או |  |
| 6:16.512 | שאינו |  |
| 6:16.513 | מחזר |  |
| 6:16.514 | לשנותה | , |
| 6:16.515 | סנהדרין |  |
| 6:16.516 | דף |  |
| 6:16.517 | צ"ט |  |
| 6:16.518 | ע"ש |  |
| 6:16.519 | ברש"י |  |
| 6:16.520 | ומהרש"א | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

