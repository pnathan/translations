# Translation pack: milhamot-hashem perek-9-03 (Milchamot ha-Shem 9:3–9:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-03`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
- והוא שלא: "provided that" (fixed).
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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


## Milchamot ha-Shem 9:3

Hebrew (Wikisource): קכב) ומדברי שני המאורות הגדולים גדולי חכמי ישראל, הרמב"ם והסמ"ג מודעא רבה לאוריתא שלא לשמוע לדברי המלך העכו"ם אשר הביא לנו את הספר החתום בשם רשב"י הזהר המסית ומטעה אותנו לחשוב באלהינו כמה עלות וסבות ולהחליף ולהמיר את כבוד אלהינו אשר אין ראשית לראשיתו באל אחר קצר אפים (זעיר אנפין) אשר לפי דבריו ניתן לו הכח והממשלה מן הסבות אשר למעלה ממנו, ושהם צוו ופקדו את כל באי עולם לעבדו ולאהבה אותו באמרם כי הוא אלהינו ואותו נעבוד, ומבת זוגו נירא. חלילה לנו להאמין כי התנא רשב"י ע"ה יאמר כדברים האלה, ושהוא עשה את החבור הזה בתור תורה שבעל פה חדשה, ויצוה עלינו לשוב בתשובה שלא לעסוק במשנה ובתלמוד. אלא בתורתו החדשה הנותנת דופי במשנה ובתלמוד לקרותה שפחה וקליפה, וסלע אחרא, ואבן משכית, ושהעוסקים בה נקראים עוסקים בתורה שלא לשמה, ודומין לכלבין דצוחין הב הב, ויורשים גיהנם האומרת הב, הב, כמו שנתבאר לעיל. חלילה לכל איש ישראל תמים עם ה' אלהיו ותורתו שבכתב ושבעל פה המסורה וחקוקה במשנה ובתלמוד, להאמין בדברים כאלה למלך העכו"ם שמסר לנו את הספר המסית אותנו לעבוד אלהים אחרים ולמשוך את לבבינו לאהבה אותם בספור מעשיות נוראות ונפלאות הנזכרים בספר הזהר אשר לא בעינינו ראינו ולא אבותינו ספרו לנו, כי עתיק יומין נראה ונגלה לרשב"י ונתן לו רשות לגלות פנים בתורה שלא כהלכה המקובלת, שהקב"ה אחד מיוחד מכל האחדים. חלילה להתנא רשב"י ע"ה וחביריו מעשות זאת להדיחנו מאחרי ה' אלהינו בחלק שפתיו. וברוב לקחו אשר ליקט מדברי רז"ל ועירבם בתוך מליצותיו להטותינו מדרך הקדש ומעבודת אבינו מלכינו יוצר בראשית. כאשה זרה ונכריה, אשר עליה אמר שלמה המלך ע"ה לשמרך מאשה זרה מנכריה אמריה החליקה וכו'. הטתו ברוב לקחה בחלק שפתיה תדיחנו וכו', להאמין ולחשוב באלהינו גופות של אור. אשר מהם ובהם תלוים כל מצות התורה ואם לא הגופות האלה אין תורה ואין מצוה, וכל מעשה מצותינו וברכותינו ותפלותינו ושבחים שאנו משבחים ומפארים לשי"ת הוא רק לגופים האלו הנבראים, דהיינו החמשה פרצופים הנחלקים עוד לשנים עשר פרצופים וביחוד לקצר אפים ובת זוגו (זו"ן) ושהוא המלך המהולל בכל התשבחות, כי כן צוו אביו ואמו לעבדו ולהללו ולשבחו לפארו ולקרותו בשם יהוה. חלילה לכל רבותינו התנאים והאמוראים מראשם ועד סופם אשר ידענו בבירור את אמונתם הטהורה ואת חכמתם חסידותם וזהירותם. לחשוב עליהם שהאמינו בפילסופיא הקדומה הטמאה המתועבת המלאה הזיות והבלים להכניסה בדת תוה"ק ולמשחה בששר, ולעשות לה ציונים ורעיונות זרות הרחוקים מן השכל הישר כדי להכניסה בתורת אלהינו תורת משה רבינו הנקייה מכל דעה זרה ונפסדת, ולהעמיס בה רבוי באלהות, בלי שום אות ומופת. רק בספורים של שקר שנגלה אליו עתיקא קדישא ונשמת משה ואליהו, ונשמות תנאים ואמוראים וגאונים שעדין לא היו ולא נבראו בימי התנא הקדוש רשב"י ז"ל. ובודאי שספורים כאלה הם ספורים בדוים שלא היו ולא נבראו, ואיך נשמע ונקבל על ידי ספורים בדוים להמיר את ה' אלהינו ארך אפים ורב חסד באל אחר קצר אפים בתבנית אדם אוכל לחם העצבים? כדגרסינן בסנהדרין פרק אלו הן הנחנקין אמר רבי אבהו אמר רבי יוחנן בכל אם יאמר לך נביא (הכדוק לך רמב"ם פ"ז מיסוה"ת) עבור על דברי תורה (לפי שעה) שמע לו חוץ מע"ז, שאפי' מעמיד לך חמה באמצע רקיע אל תשמע לו, תניא רבי יוסי הגלילי אומר ראה עד היכן הגיע סוף עובדי אלילים שנתן להם ממשלה, אפי' אם מעמידים חמה באמצע רקיע אל תשמע להן מפני מה? כי מנסה ה' אלהיכם אתכם. אמר רבי עקיבה ח"ו שהקב"ה מעמיד חמה ולבנה לעוברי רצונו. הא אינו מדבר אלא במי שהיה נביא אמת, וחזר להיות נביא שקר ע"כ. והן הן דברי הרמב"ם והסמ"ג שאפי' יעמוד נביא ויעשה אותות ומופתים בשמים ובארץ כמשה בן עמרם ויאמר להחליף תורתינו, או לבטל או להוסיף שום מצוה לא נשמע לו.

| id | word | punctuation after |
|---|---|---|
| 9:3.1 | קכב | ) |
| 9:3.2 | ומדברי |  |
| 9:3.3 | שני |  |
| 9:3.4 | המאורות |  |
| 9:3.5 | הגדולים |  |
| 9:3.6 | גדולי |  |
| 9:3.7 | חכמי |  |
| 9:3.8 | ישראל | , |
| 9:3.9 | הרמב"ם |  |
| 9:3.10 | והסמ"ג |  |
| 9:3.11 | מודעא |  |
| 9:3.12 | רבה |  |
| 9:3.13 | לאוריתא |  |
| 9:3.14 | שלא |  |
| 9:3.15 | לשמוע |  |
| 9:3.16 | לדברי |  |
| 9:3.17 | המלך |  |
| 9:3.18 | העכו"ם |  |
| 9:3.19 | אשר |  |
| 9:3.20 | הביא |  |
| 9:3.21 | לנו |  |
| 9:3.22 | את |  |
| 9:3.23 | הספר |  |
| 9:3.24 | החתום |  |
| 9:3.25 | בשם |  |
| 9:3.26 | רשב"י |  |
| 9:3.27 | הזהר |  |
| 9:3.28 | המסית |  |
| 9:3.29 | ומטעה |  |
| 9:3.30 | אותנו |  |
| 9:3.31 | לחשוב |  |
| 9:3.32 | באלהינו |  |
| 9:3.33 | כמה |  |
| 9:3.34 | עלות |  |
| 9:3.35 | וסבות |  |
| 9:3.36 | ולהחליף |  |
| 9:3.37 | ולהמיר |  |
| 9:3.38 | את |  |
| 9:3.39 | כבוד |  |
| 9:3.40 | אלהינו |  |
| 9:3.41 | אשר |  |
| 9:3.42 | אין |  |
| 9:3.43 | ראשית |  |
| 9:3.44 | לראשיתו |  |
| 9:3.45 | באל |  |
| 9:3.46 | אחר |  |
| 9:3.47 | קצר |  |
| 9:3.48 | אפים |  |
| 9:3.49 | זעיר |  |
| 9:3.50 | אנפין | ) |
| 9:3.51 | אשר |  |
| 9:3.52 | לפי |  |
| 9:3.53 | דבריו |  |
| 9:3.54 | ניתן |  |
| 9:3.55 | לו |  |
| 9:3.56 | הכח |  |
| 9:3.57 | והממשלה |  |
| 9:3.58 | מן |  |
| 9:3.59 | הסבות |  |
| 9:3.60 | אשר |  |
| 9:3.61 | למעלה |  |
| 9:3.62 | ממנו | , |
| 9:3.63 | ושהם |  |
| 9:3.64 | צוו |  |
| 9:3.65 | ופקדו |  |
| 9:3.66 | את |  |
| 9:3.67 | כל |  |
| 9:3.68 | באי |  |
| 9:3.69 | עולם |  |
| 9:3.70 | לעבדו |  |
| 9:3.71 | ולאהבה |  |
| 9:3.72 | אותו |  |
| 9:3.73 | באמרם |  |
| 9:3.74 | כי |  |
| 9:3.75 | הוא |  |
| 9:3.76 | אלהינו |  |
| 9:3.77 | ואותו |  |
| 9:3.78 | נעבוד | , |
| 9:3.79 | ומבת |  |
| 9:3.80 | זוגו |  |
| 9:3.81 | נירא | . |
| 9:3.82 | חלילה |  |
| 9:3.83 | לנו |  |
| 9:3.84 | להאמין |  |
| 9:3.85 | כי |  |
| 9:3.86 | התנא |  |
| 9:3.87 | רשב"י |  |
| 9:3.88 | ע"ה |  |
| 9:3.89 | יאמר |  |
| 9:3.90 | כדברים |  |
| 9:3.91 | האלה | , |
| 9:3.92 | ושהוא |  |
| 9:3.93 | עשה |  |
| 9:3.94 | את |  |
| 9:3.95 | החבור |  |
| 9:3.96 | הזה |  |
| 9:3.97 | בתור |  |
| 9:3.98 | תורה |  |
| 9:3.99 | שבעל |  |
| 9:3.100 | פה |  |
| 9:3.101 | חדשה | , |
| 9:3.102 | ויצוה |  |
| 9:3.103 | עלינו |  |
| 9:3.104 | לשוב |  |
| 9:3.105 | בתשובה |  |
| 9:3.106 | שלא |  |
| 9:3.107 | לעסוק |  |
| 9:3.108 | במשנה |  |
| 9:3.109 | ובתלמוד | . |
| 9:3.110 | אלא |  |
| 9:3.111 | בתורתו |  |
| 9:3.112 | החדשה |  |
| 9:3.113 | הנותנת |  |
| 9:3.114 | דופי |  |
| 9:3.115 | במשנה |  |
| 9:3.116 | ובתלמוד |  |
| 9:3.117 | לקרותה |  |
| 9:3.118 | שפחה |  |
| 9:3.119 | וקליפה | , |
| 9:3.120 | וסלע |  |
| 9:3.121 | אחרא | , |
| 9:3.122 | ואבן |  |
| 9:3.123 | משכית | , |
| 9:3.124 | ושהעוסקים |  |
| 9:3.125 | בה |  |
| 9:3.126 | נקראים |  |
| 9:3.127 | עוסקים |  |
| 9:3.128 | בתורה |  |
| 9:3.129 | שלא |  |
| 9:3.130 | לשמה | , |
| 9:3.131 | ודומין |  |
| 9:3.132 | לכלבין |  |
| 9:3.133 | דצוחין |  |
| 9:3.134 | הב |  |
| 9:3.135 | הב | , |
| 9:3.136 | ויורשים |  |
| 9:3.137 | גיהנם |  |
| 9:3.138 | האומרת |  |
| 9:3.139 | הב | , |
| 9:3.140 | הב | , |
| 9:3.141 | כמו |  |
| 9:3.142 | שנתבאר |  |
| 9:3.143 | לעיל | . |
| 9:3.144 | חלילה |  |
| 9:3.145 | לכל |  |
| 9:3.146 | איש |  |
| 9:3.147 | ישראל |  |
| 9:3.148 | תמים |  |
| 9:3.149 | עם |  |
| 9:3.150 | ה | ' |
| 9:3.151 | אלהיו |  |
| 9:3.152 | ותורתו |  |
| 9:3.153 | שבכתב |  |
| 9:3.154 | ושבעל |  |
| 9:3.155 | פה |  |
| 9:3.156 | המסורה |  |
| 9:3.157 | וחקוקה |  |
| 9:3.158 | במשנה |  |
| 9:3.159 | ובתלמוד | , |
| 9:3.160 | להאמין |  |
| 9:3.161 | בדברים |  |
| 9:3.162 | כאלה |  |
| 9:3.163 | למלך |  |
| 9:3.164 | העכו"ם |  |
| 9:3.165 | שמסר |  |
| 9:3.166 | לנו |  |
| 9:3.167 | את |  |
| 9:3.168 | הספר |  |
| 9:3.169 | המסית |  |
| 9:3.170 | אותנו |  |
| 9:3.171 | לעבוד |  |
| 9:3.172 | אלהים |  |
| 9:3.173 | אחרים |  |
| 9:3.174 | ולמשוך |  |
| 9:3.175 | את |  |
| 9:3.176 | לבבינו |  |
| 9:3.177 | לאהבה |  |
| 9:3.178 | אותם |  |
| 9:3.179 | בספור |  |
| 9:3.180 | מעשיות |  |
| 9:3.181 | נוראות |  |
| 9:3.182 | ונפלאות |  |
| 9:3.183 | הנזכרים |  |
| 9:3.184 | בספר |  |
| 9:3.185 | הזהר |  |
| 9:3.186 | אשר |  |
| 9:3.187 | לא |  |
| 9:3.188 | בעינינו |  |
| 9:3.189 | ראינו |  |
| 9:3.190 | ולא |  |
| 9:3.191 | אבותינו |  |
| 9:3.192 | ספרו |  |
| 9:3.193 | לנו | , |
| 9:3.194 | כי |  |
| 9:3.195 | עתיק |  |
| 9:3.196 | יומין |  |
| 9:3.197 | נראה |  |
| 9:3.198 | ונגלה |  |
| 9:3.199 | לרשב"י |  |
| 9:3.200 | ונתן |  |
| 9:3.201 | לו |  |
| 9:3.202 | רשות |  |
| 9:3.203 | לגלות |  |
| 9:3.204 | פנים |  |
| 9:3.205 | בתורה |  |
| 9:3.206 | שלא |  |
| 9:3.207 | כהלכה |  |
| 9:3.208 | המקובלת | , |
| 9:3.209 | שהקב"ה |  |
| 9:3.210 | אחד |  |
| 9:3.211 | מיוחד |  |
| 9:3.212 | מכל |  |
| 9:3.213 | האחדים | . |
| 9:3.214 | חלילה |  |
| 9:3.215 | להתנא |  |
| 9:3.216 | רשב"י |  |
| 9:3.217 | ע"ה |  |
| 9:3.218 | וחביריו |  |
| 9:3.219 | מעשות |  |
| 9:3.220 | זאת |  |
| 9:3.221 | להדיחנו |  |
| 9:3.222 | מאחרי |  |
| 9:3.223 | ה | ' |
| 9:3.224 | אלהינו |  |
| 9:3.225 | בחלק |  |
| 9:3.226 | שפתיו | . |
| 9:3.227 | וברוב |  |
| 9:3.228 | לקחו |  |
| 9:3.229 | אשר |  |
| 9:3.230 | ליקט |  |
| 9:3.231 | מדברי |  |
| 9:3.232 | רז"ל |  |
| 9:3.233 | ועירבם |  |
| 9:3.234 | בתוך |  |
| 9:3.235 | מליצותיו |  |
| 9:3.236 | להטותינו |  |
| 9:3.237 | מדרך |  |
| 9:3.238 | הקדש |  |
| 9:3.239 | ומעבודת |  |
| 9:3.240 | אבינו |  |
| 9:3.241 | מלכינו |  |
| 9:3.242 | יוצר |  |
| 9:3.243 | בראשית | . |
| 9:3.244 | כאשה |  |
| 9:3.245 | זרה |  |
| 9:3.246 | ונכריה | , |
| 9:3.247 | אשר |  |
| 9:3.248 | עליה |  |
| 9:3.249 | אמר |  |
| 9:3.250 | שלמה |  |
| 9:3.251 | המלך |  |
| 9:3.252 | ע"ה |  |
| 9:3.253 | לשמרך |  |
| 9:3.254 | מאשה |  |
| 9:3.255 | זרה |  |
| 9:3.256 | מנכריה |  |
| 9:3.257 | אמריה |  |
| 9:3.258 | החליקה |  |
| 9:3.259 | וכו | '. |
| 9:3.260 | הטתו |  |
| 9:3.261 | ברוב |  |
| 9:3.262 | לקחה |  |
| 9:3.263 | בחלק |  |
| 9:3.264 | שפתיה |  |
| 9:3.265 | תדיחנו |  |
| 9:3.266 | וכו | ', |
| 9:3.267 | להאמין |  |
| 9:3.268 | ולחשוב |  |
| 9:3.269 | באלהינו |  |
| 9:3.270 | גופות |  |
| 9:3.271 | של |  |
| 9:3.272 | אור | . |
| 9:3.273 | אשר |  |
| 9:3.274 | מהם |  |
| 9:3.275 | ובהם |  |
| 9:3.276 | תלוים |  |
| 9:3.277 | כל |  |
| 9:3.278 | מצות |  |
| 9:3.279 | התורה |  |
| 9:3.280 | ואם |  |
| 9:3.281 | לא |  |
| 9:3.282 | הגופות |  |
| 9:3.283 | האלה |  |
| 9:3.284 | אין |  |
| 9:3.285 | תורה |  |
| 9:3.286 | ואין |  |
| 9:3.287 | מצוה | , |
| 9:3.288 | וכל |  |
| 9:3.289 | מעשה |  |
| 9:3.290 | מצותינו |  |
| 9:3.291 | וברכותינו |  |
| 9:3.292 | ותפלותינו |  |
| 9:3.293 | ושבחים |  |
| 9:3.294 | שאנו |  |
| 9:3.295 | משבחים |  |
| 9:3.296 | ומפארים |  |
| 9:3.297 | לשי"ת |  |
| 9:3.298 | הוא |  |
| 9:3.299 | רק |  |
| 9:3.300 | לגופים |  |
| 9:3.301 | האלו |  |
| 9:3.302 | הנבראים | , |
| 9:3.303 | דהיינו |  |
| 9:3.304 | החמשה |  |
| 9:3.305 | פרצופים |  |
| 9:3.306 | הנחלקים |  |
| 9:3.307 | עוד |  |
| 9:3.308 | לשנים |  |
| 9:3.309 | עשר |  |
| 9:3.310 | פרצופים |  |
| 9:3.311 | וביחוד |  |
| 9:3.312 | לקצר |  |
| 9:3.313 | אפים |  |
| 9:3.314 | ובת |  |
| 9:3.315 | זוגו |  |
| 9:3.316 | זו"ן | ) |
| 9:3.317 | ושהוא |  |
| 9:3.318 | המלך |  |
| 9:3.319 | המהולל |  |
| 9:3.320 | בכל |  |
| 9:3.321 | התשבחות | , |
| 9:3.322 | כי |  |
| 9:3.323 | כן |  |
| 9:3.324 | צוו |  |
| 9:3.325 | אביו |  |
| 9:3.326 | ואמו |  |
| 9:3.327 | לעבדו |  |
| 9:3.328 | ולהללו |  |
| 9:3.329 | ולשבחו |  |
| 9:3.330 | לפארו |  |
| 9:3.331 | ולקרותו |  |
| 9:3.332 | בשם |  |
| 9:3.333 | יהוה | . |
| 9:3.334 | חלילה |  |
| 9:3.335 | לכל |  |
| 9:3.336 | רבותינו |  |
| 9:3.337 | התנאים |  |
| 9:3.338 | והאמוראים |  |
| 9:3.339 | מראשם |  |
| 9:3.340 | ועד |  |
| 9:3.341 | סופם |  |
| 9:3.342 | אשר |  |
| 9:3.343 | ידענו |  |
| 9:3.344 | בבירור |  |
| 9:3.345 | את |  |
| 9:3.346 | אמונתם |  |
| 9:3.347 | הטהורה |  |
| 9:3.348 | ואת |  |
| 9:3.349 | חכמתם |  |
| 9:3.350 | חסידותם |  |
| 9:3.351 | וזהירותם | . |
| 9:3.352 | לחשוב |  |
| 9:3.353 | עליהם |  |
| 9:3.354 | שהאמינו |  |
| 9:3.355 | בפילסופיא |  |
| 9:3.356 | הקדומה |  |
| 9:3.357 | הטמאה |  |
| 9:3.358 | המתועבת |  |
| 9:3.359 | המלאה |  |
| 9:3.360 | הזיות |  |
| 9:3.361 | והבלים |  |
| 9:3.362 | להכניסה |  |
| 9:3.363 | בדת |  |
| 9:3.364 | תוה"ק |  |
| 9:3.365 | ולמשחה |  |
| 9:3.366 | בששר | , |
| 9:3.367 | ולעשות |  |
| 9:3.368 | לה |  |
| 9:3.369 | ציונים |  |
| 9:3.370 | ורעיונות |  |
| 9:3.371 | זרות |  |
| 9:3.372 | הרחוקים |  |
| 9:3.373 | מן |  |
| 9:3.374 | השכל |  |
| 9:3.375 | הישר |  |
| 9:3.376 | כדי |  |
| 9:3.377 | להכניסה |  |
| 9:3.378 | בתורת |  |
| 9:3.379 | אלהינו |  |
| 9:3.380 | תורת |  |
| 9:3.381 | משה |  |
| 9:3.382 | רבינו |  |
| 9:3.383 | הנקייה |  |
| 9:3.384 | מכל |  |
| 9:3.385 | דעה |  |
| 9:3.386 | זרה |  |
| 9:3.387 | ונפסדת | , |
| 9:3.388 | ולהעמיס |  |
| 9:3.389 | בה |  |
| 9:3.390 | רבוי |  |
| 9:3.391 | באלהות | , |
| 9:3.392 | בלי |  |
| 9:3.393 | שום |  |
| 9:3.394 | אות |  |
| 9:3.395 | ומופת | . |
| 9:3.396 | רק |  |
| 9:3.397 | בספורים |  |
| 9:3.398 | של |  |
| 9:3.399 | שקר |  |
| 9:3.400 | שנגלה |  |
| 9:3.401 | אליו |  |
| 9:3.402 | עתיקא |  |
| 9:3.403 | קדישא |  |
| 9:3.404 | ונשמת |  |
| 9:3.405 | משה |  |
| 9:3.406 | ואליהו | , |
| 9:3.407 | ונשמות |  |
| 9:3.408 | תנאים |  |
| 9:3.409 | ואמוראים |  |
| 9:3.410 | וגאונים |  |
| 9:3.411 | שעדין |  |
| 9:3.412 | לא |  |
| 9:3.413 | היו |  |
| 9:3.414 | ולא |  |
| 9:3.415 | נבראו |  |
| 9:3.416 | בימי |  |
| 9:3.417 | התנא |  |
| 9:3.418 | הקדוש |  |
| 9:3.419 | רשב"י |  |
| 9:3.420 | ז"ל | . |
| 9:3.421 | ובודאי |  |
| 9:3.422 | שספורים |  |
| 9:3.423 | כאלה |  |
| 9:3.424 | הם |  |
| 9:3.425 | ספורים |  |
| 9:3.426 | בדוים |  |
| 9:3.427 | שלא |  |
| 9:3.428 | היו |  |
| 9:3.429 | ולא |  |
| 9:3.430 | נבראו | , |
| 9:3.431 | ואיך |  |
| 9:3.432 | נשמע |  |
| 9:3.433 | ונקבל |  |
| 9:3.434 | על |  |
| 9:3.435 | ידי |  |
| 9:3.436 | ספורים |  |
| 9:3.437 | בדוים |  |
| 9:3.438 | להמיר |  |
| 9:3.439 | את |  |
| 9:3.440 | ה | ' |
| 9:3.441 | אלהינו |  |
| 9:3.442 | ארך |  |
| 9:3.443 | אפים |  |
| 9:3.444 | ורב |  |
| 9:3.445 | חסד |  |
| 9:3.446 | באל |  |
| 9:3.447 | אחר |  |
| 9:3.448 | קצר |  |
| 9:3.449 | אפים |  |
| 9:3.450 | בתבנית |  |
| 9:3.451 | אדם |  |
| 9:3.452 | אוכל |  |
| 9:3.453 | לחם |  |
| 9:3.454 | העצבים | ? |
| 9:3.455 | כדגרסינן |  |
| 9:3.456 | בסנהדרין |  |
| 9:3.457 | פרק |  |
| 9:3.458 | אלו |  |
| 9:3.459 | הן |  |
| 9:3.460 | הנחנקין |  |
| 9:3.461 | אמר |  |
| 9:3.462 | רבי |  |
| 9:3.463 | אבהו |  |
| 9:3.464 | אמר |  |
| 9:3.465 | רבי |  |
| 9:3.466 | יוחנן |  |
| 9:3.467 | בכל |  |
| 9:3.468 | אם |  |
| 9:3.469 | יאמר |  |
| 9:3.470 | לך |  |
| 9:3.471 | נביא |  |
| 9:3.472 | הכדוק |  |
| 9:3.473 | לך |  |
| 9:3.474 | רמב"ם |  |
| 9:3.475 | פ"ז |  |
| 9:3.476 | מיסוה"ת | ) |
| 9:3.477 | עבור |  |
| 9:3.478 | על |  |
| 9:3.479 | דברי |  |
| 9:3.480 | תורה |  |
| 9:3.481 | לפי |  |
| 9:3.482 | שעה | ) |
| 9:3.483 | שמע |  |
| 9:3.484 | לו |  |
| 9:3.485 | חוץ |  |
| 9:3.486 | מע"ז | , |
| 9:3.487 | שאפי | ' |
| 9:3.488 | מעמיד |  |
| 9:3.489 | לך |  |
| 9:3.490 | חמה |  |
| 9:3.491 | באמצע |  |
| 9:3.492 | רקיע |  |
| 9:3.493 | אל |  |
| 9:3.494 | תשמע |  |
| 9:3.495 | לו | , |
| 9:3.496 | תניא |  |
| 9:3.497 | רבי |  |
| 9:3.498 | יוסי |  |
| 9:3.499 | הגלילי |  |
| 9:3.500 | אומר |  |
| 9:3.501 | ראה |  |
| 9:3.502 | עד |  |
| 9:3.503 | היכן |  |
| 9:3.504 | הגיע |  |
| 9:3.505 | סוף |  |
| 9:3.506 | עובדי |  |
| 9:3.507 | אלילים |  |
| 9:3.508 | שנתן |  |
| 9:3.509 | להם |  |
| 9:3.510 | ממשלה | , |
| 9:3.511 | אפי | ' |
| 9:3.512 | אם |  |
| 9:3.513 | מעמידים |  |
| 9:3.514 | חמה |  |
| 9:3.515 | באמצע |  |
| 9:3.516 | רקיע |  |
| 9:3.517 | אל |  |
| 9:3.518 | תשמע |  |
| 9:3.519 | להן |  |
| 9:3.520 | מפני |  |
| 9:3.521 | מה | ? |
| 9:3.522 | כי |  |
| 9:3.523 | מנסה |  |
| 9:3.524 | ה | ' |
| 9:3.525 | אלהיכם |  |
| 9:3.526 | אתכם | . |
| 9:3.527 | אמר |  |
| 9:3.528 | רבי |  |
| 9:3.529 | עקיבה |  |
| 9:3.530 | ח"ו |  |
| 9:3.531 | שהקב"ה |  |
| 9:3.532 | מעמיד |  |
| 9:3.533 | חמה |  |
| 9:3.534 | ולבנה |  |
| 9:3.535 | לעוברי |  |
| 9:3.536 | רצונו | . |
| 9:3.537 | הא |  |
| 9:3.538 | אינו |  |
| 9:3.539 | מדבר |  |
| 9:3.540 | אלא |  |
| 9:3.541 | במי |  |
| 9:3.542 | שהיה |  |
| 9:3.543 | נביא |  |
| 9:3.544 | אמת | , |
| 9:3.545 | וחזר |  |
| 9:3.546 | להיות |  |
| 9:3.547 | נביא |  |
| 9:3.548 | שקר |  |
| 9:3.549 | ע"כ | . |
| 9:3.550 | והן |  |
| 9:3.551 | הן |  |
| 9:3.552 | דברי |  |
| 9:3.553 | הרמב"ם |  |
| 9:3.554 | והסמ"ג |  |
| 9:3.555 | שאפי | ' |
| 9:3.556 | יעמוד |  |
| 9:3.557 | נביא |  |
| 9:3.558 | ויעשה |  |
| 9:3.559 | אותות |  |
| 9:3.560 | ומופתים |  |
| 9:3.561 | בשמים |  |
| 9:3.562 | ובארץ |  |
| 9:3.563 | כמשה |  |
| 9:3.564 | בן |  |
| 9:3.565 | עמרם |  |
| 9:3.566 | ויאמר |  |
| 9:3.567 | להחליף |  |
| 9:3.568 | תורתינו | , |
| 9:3.569 | או |  |
| 9:3.570 | לבטל |  |
| 9:3.571 | או |  |
| 9:3.572 | להוסיף |  |
| 9:3.573 | שום |  |
| 9:3.574 | מצוה |  |
| 9:3.575 | לא |  |
| 9:3.576 | נשמע |  |
| 9:3.577 | לו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

