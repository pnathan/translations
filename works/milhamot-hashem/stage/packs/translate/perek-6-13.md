# Translation pack: milhamot-hashem perek-6-13 (Milchamot ha-Shem 6:20–6:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-13`.

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
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

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


## Milchamot ha-Shem 6:20

Hebrew (Wikisource): ובהקדמת תקוניה הזהר דף ג' ע"ב, גם בתחלת הקדמה שנייה ובסוף תיקונא שתיתאה כ' וכמה בני נשא יתפרנסון מהאי חבורה דילך בסוף יומיא וכו'. ומשפתי חלקות מסית ומדיח לזלזל בלימוד משנה ותלמוד, וחשב עסקם לעון אשר חטא עד שצריך תשובה לעזוב אותם ולהתעסק בתורתו החדשה המלאה תארים ותבניות צורות ופרצופים רבים שונים זה מזה אשר עלה על רוחו לעשותם גופים, ולהוריד להם נשמות מעצמות האלהים העליון שאין לו דמות הגוף בידים ורגלים ושאר אברי הגוף רק דמות עגול גדול מקיף את כל החלל שבתוכו לפי דעתו ודרך צינור דק משך כמנו הפילסוף המתעה נשמות ומלבושים לכסות מערומי הצורות והפרצופים כמבואר בעץ חיים להרח"ו, ובשאר ספרי המקבלים ממנו אשר לא שמו אל לבם דברי רז"ל בגמרא ומדרש רבה ותנחומא וכיוצא מספרי רז"ל האמתיים.

| id | word | punctuation after |
|---|---|---|
| 6:20.1 | ובהקדמת |  |
| 6:20.2 | תקוניה |  |
| 6:20.3 | הזהר |  |
| 6:20.4 | דף |  |
| 6:20.5 | ג | ' |
| 6:20.6 | ע"ב | , |
| 6:20.7 | גם |  |
| 6:20.8 | בתחלת |  |
| 6:20.9 | הקדמה |  |
| 6:20.10 | שנייה |  |
| 6:20.11 | ובסוף |  |
| 6:20.12 | תיקונא |  |
| 6:20.13 | שתיתאה |  |
| 6:20.14 | כ | ' |
| 6:20.15 | וכמה |  |
| 6:20.16 | בני |  |
| 6:20.17 | נשא |  |
| 6:20.18 | יתפרנסון |  |
| 6:20.19 | מהאי |  |
| 6:20.20 | חבורה |  |
| 6:20.21 | דילך |  |
| 6:20.22 | בסוף |  |
| 6:20.23 | יומיא |  |
| 6:20.24 | וכו | '. |
| 6:20.25 | ומשפתי |  |
| 6:20.26 | חלקות |  |
| 6:20.27 | מסית |  |
| 6:20.28 | ומדיח |  |
| 6:20.29 | לזלזל |  |
| 6:20.30 | בלימוד |  |
| 6:20.31 | משנה |  |
| 6:20.32 | ותלמוד | , |
| 6:20.33 | וחשב |  |
| 6:20.34 | עסקם |  |
| 6:20.35 | לעון |  |
| 6:20.36 | אשר |  |
| 6:20.37 | חטא |  |
| 6:20.38 | עד |  |
| 6:20.39 | שצריך |  |
| 6:20.40 | תשובה |  |
| 6:20.41 | לעזוב |  |
| 6:20.42 | אותם |  |
| 6:20.43 | ולהתעסק |  |
| 6:20.44 | בתורתו |  |
| 6:20.45 | החדשה |  |
| 6:20.46 | המלאה |  |
| 6:20.47 | תארים |  |
| 6:20.48 | ותבניות |  |
| 6:20.49 | צורות |  |
| 6:20.50 | ופרצופים |  |
| 6:20.51 | רבים |  |
| 6:20.52 | שונים |  |
| 6:20.53 | זה |  |
| 6:20.54 | מזה |  |
| 6:20.55 | אשר |  |
| 6:20.56 | עלה |  |
| 6:20.57 | על |  |
| 6:20.58 | רוחו |  |
| 6:20.59 | לעשותם |  |
| 6:20.60 | גופים | , |
| 6:20.61 | ולהוריד |  |
| 6:20.62 | להם |  |
| 6:20.63 | נשמות |  |
| 6:20.64 | מעצמות |  |
| 6:20.65 | האלהים |  |
| 6:20.66 | העליון |  |
| 6:20.67 | שאין |  |
| 6:20.68 | לו |  |
| 6:20.69 | דמות |  |
| 6:20.70 | הגוף |  |
| 6:20.71 | בידים |  |
| 6:20.72 | ורגלים |  |
| 6:20.73 | ושאר |  |
| 6:20.74 | אברי |  |
| 6:20.75 | הגוף |  |
| 6:20.76 | רק |  |
| 6:20.77 | דמות |  |
| 6:20.78 | עגול |  |
| 6:20.79 | גדול |  |
| 6:20.80 | מקיף |  |
| 6:20.81 | את |  |
| 6:20.82 | כל |  |
| 6:20.83 | החלל |  |
| 6:20.84 | שבתוכו |  |
| 6:20.85 | לפי |  |
| 6:20.86 | דעתו |  |
| 6:20.87 | ודרך |  |
| 6:20.88 | צינור |  |
| 6:20.89 | דק |  |
| 6:20.90 | משך |  |
| 6:20.91 | כמנו |  |
| 6:20.92 | הפילסוף |  |
| 6:20.93 | המתעה |  |
| 6:20.94 | נשמות |  |
| 6:20.95 | ומלבושים |  |
| 6:20.96 | לכסות |  |
| 6:20.97 | מערומי |  |
| 6:20.98 | הצורות |  |
| 6:20.99 | והפרצופים |  |
| 6:20.100 | כמבואר |  |
| 6:20.101 | בעץ |  |
| 6:20.102 | חיים |  |
| 6:20.103 | להרח"ו | , |
| 6:20.104 | ובשאר |  |
| 6:20.105 | ספרי |  |
| 6:20.106 | המקבלים |  |
| 6:20.107 | ממנו |  |
| 6:20.108 | אשר |  |
| 6:20.109 | לא |  |
| 6:20.110 | שמו |  |
| 6:20.111 | אל |  |
| 6:20.112 | לבם |  |
| 6:20.113 | דברי |  |
| 6:20.114 | רז"ל |  |
| 6:20.115 | בגמרא |  |
| 6:20.116 | ומדרש |  |
| 6:20.117 | רבה |  |
| 6:20.118 | ותנחומא |  |
| 6:20.119 | וכיוצא |  |
| 6:20.120 | מספרי |  |
| 6:20.121 | רז"ל |  |
| 6:20.122 | האמתיים | . |


## Milchamot ha-Shem 6:21

Hebrew (Wikisource): ובתיקוני זהר חדש (דף ו' וז' ודף י') הרחיב עוד את פיו לספר בגנות העוסקים במשנה ובתלמוד כדי שיקרא רבי או כדי להשיג עושר בעה"ז או שיזכה לחיי העה"ב, שאמרו חכמים עליהם לעולם יעסוק אדם בתורה ובמצות אפי' שלא לשמה, שמתוך וכו' וגם אמרו האומר סלע זו לצדקה בשביל שיחיו בני בשביל שאזכה לחיי העולם הבא הרי זה צדיק גמור. והפילסוף מחבר הזהר שם אותם ליורשי גהינם, ודימה שקלא וטריא דרבנן העוסקים בתורה כדי שיזכו לחיי העולם הבא לכלבים המנבחים ואומרים חף חף הב הב. ואמר שהם יורשים גיהנם האומרת הב הב. כמו שמסיים שם שעליהם נאמר ואלה לחרפות ולדראון עולם. ולעג על מאמר התנא למוד תורה הרבה ויתנו לך שכר הרבה. ואף אם עוסק בתורה להבדיל בין הטהור ובין הטמא כאשר ציונו ה' אלהינו להבדיל בין החיה הנאכלת ובין החיה אשר לא תאכל, ושאר איסורי מאכלות הכל נקרא אצלו עוסק שלא לשמה כמו שאמר שם - לא משתדלין אלא לגרמייהו לדבתא בשרא וכו'. ולא נקרא עוסק לשמה אלא העוסק בתורה כדי לחבר האלהות הזכרים בנקבות, אבא עם אימא, וזעיר עם נוקביה. לזוגם יחד והיו לבשר אחד. וכן בקריאת שמע וציצית ותפילין ושאר כל המצות. ולכן תקנו המקובלים לומר קודם כל מצוה שעושים לשם יחוד קב"ה ושכינתיה וכו', כי תכלית כוונת הפילסוף להכניס בישראל אמונת רבוי באלוהות זכרים ונקבות ולעשות חבורים שלו שהם הזהר והתיקונים עיקר לעסוק בהם חלף המשנה והתלמוד. ולשכח מישראל עסק המשנה והתלמוד שכל כללותיה ופרטותיה של תורה כלולים בהם. ובכמה מקומות הוא משבח את חבורו זה ומהם בתיקוני זוהר חדש דף י"ו ע"ב. ובנפת אשר יטיף הפילסוף הנז' בשפתי אמונה שהיא זרה לאמונת תורתינו הקדושה משך לב רבים בחלקלקותיו, וישיאם לעבוד אלהים אחרים המה הצורות והפרצופים אשר בדה מלבו להטעותנו בפירוש עשר ספירות הנז' בספר יצירה לחשבם אלהות ולעבוד אותם בעבור שנשמותיהם חצובים מהאלהים העליון (לפי דעתו) אשר חלק להם מעצמותו ומהותו הם ועיגוליהם וגופם ונשמתם ולבושם, כולם חוצבו ממקור אחד (הוא הא"ם) אשר לפי דעתם אין לו שם ונקודה ולית מחשבה תפיסה ביה, ולא שייך שום עבודה וקריאה אליו, ולא יענה את הקורא אליו שכבר מינה והפקיד את הפרצופים על כל מלאכת הבריאה. ועשה את אבא בורא, והוא אומר ומצוה לאימא יהי כך וכך, והיא עושה כדבריו, כאומן שעושה מלאכת בעל הבית באמונה. ואבא ואמא גם הם צוו והפקידו את זעיר אנפין ונוקבתיה להשגיח ולרון על מעשה התחתונים. מאן דזכי לדינא לדינא, ומאן דזכי לרחמי לרחמי, הכל נעשה כחפצו ורצונו של זעיר אנפין. וכל דברי המחבר הזה בנוים על הפילוסופיא הקדומה אשר כבר נמר ריחה וטעמה מעת אשר נגלה כבודו ית' על הר סיני ואמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים. רק שאחרי מות יהושע והזקנים חזרו להדיעות האלו הנפסדות ויעבדו את הבעלים והעשתרות האלו ויעשו צלמיהם כדי להוריד השפע מהן על אותן הצורות והצלמים לפי דעתם.

| id | word | punctuation after |
|---|---|---|
| 6:21.1 | ובתיקוני |  |
| 6:21.2 | זהר |  |
| 6:21.3 | חדש |  |
| 6:21.4 | דף |  |
| 6:21.5 | ו | ' |
| 6:21.6 | וז | ' |
| 6:21.7 | ודף |  |
| 6:21.8 | י | ') |
| 6:21.9 | הרחיב |  |
| 6:21.10 | עוד |  |
| 6:21.11 | את |  |
| 6:21.12 | פיו |  |
| 6:21.13 | לספר |  |
| 6:21.14 | בגנות |  |
| 6:21.15 | העוסקים |  |
| 6:21.16 | במשנה |  |
| 6:21.17 | ובתלמוד |  |
| 6:21.18 | כדי |  |
| 6:21.19 | שיקרא |  |
| 6:21.20 | רבי |  |
| 6:21.21 | או |  |
| 6:21.22 | כדי |  |
| 6:21.23 | להשיג |  |
| 6:21.24 | עושר |  |
| 6:21.25 | בעה"ז |  |
| 6:21.26 | או |  |
| 6:21.27 | שיזכה |  |
| 6:21.28 | לחיי |  |
| 6:21.29 | העה"ב | , |
| 6:21.30 | שאמרו |  |
| 6:21.31 | חכמים |  |
| 6:21.32 | עליהם |  |
| 6:21.33 | לעולם |  |
| 6:21.34 | יעסוק |  |
| 6:21.35 | אדם |  |
| 6:21.36 | בתורה |  |
| 6:21.37 | ובמצות |  |
| 6:21.38 | אפי | ' |
| 6:21.39 | שלא |  |
| 6:21.40 | לשמה | , |
| 6:21.41 | שמתוך |  |
| 6:21.42 | וכו | ' |
| 6:21.43 | וגם |  |
| 6:21.44 | אמרו |  |
| 6:21.45 | האומר |  |
| 6:21.46 | סלע |  |
| 6:21.47 | זו |  |
| 6:21.48 | לצדקה |  |
| 6:21.49 | בשביל |  |
| 6:21.50 | שיחיו |  |
| 6:21.51 | בני |  |
| 6:21.52 | בשביל |  |
| 6:21.53 | שאזכה |  |
| 6:21.54 | לחיי |  |
| 6:21.55 | העולם |  |
| 6:21.56 | הבא |  |
| 6:21.57 | הרי |  |
| 6:21.58 | זה |  |
| 6:21.59 | צדיק |  |
| 6:21.60 | גמור | . |
| 6:21.61 | והפילסוף |  |
| 6:21.62 | מחבר |  |
| 6:21.63 | הזהר |  |
| 6:21.64 | שם |  |
| 6:21.65 | אותם |  |
| 6:21.66 | ליורשי |  |
| 6:21.67 | גהינם | , |
| 6:21.68 | ודימה |  |
| 6:21.69 | שקלא |  |
| 6:21.70 | וטריא |  |
| 6:21.71 | דרבנן |  |
| 6:21.72 | העוסקים |  |
| 6:21.73 | בתורה |  |
| 6:21.74 | כדי |  |
| 6:21.75 | שיזכו |  |
| 6:21.76 | לחיי |  |
| 6:21.77 | העולם |  |
| 6:21.78 | הבא |  |
| 6:21.79 | לכלבים |  |
| 6:21.80 | המנבחים |  |
| 6:21.81 | ואומרים |  |
| 6:21.82 | חף |  |
| 6:21.83 | חף |  |
| 6:21.84 | הב |  |
| 6:21.85 | הב | . |
| 6:21.86 | ואמר |  |
| 6:21.87 | שהם |  |
| 6:21.88 | יורשים |  |
| 6:21.89 | גיהנם |  |
| 6:21.90 | האומרת |  |
| 6:21.91 | הב |  |
| 6:21.92 | הב | . |
| 6:21.93 | כמו |  |
| 6:21.94 | שמסיים |  |
| 6:21.95 | שם |  |
| 6:21.96 | שעליהם |  |
| 6:21.97 | נאמר |  |
| 6:21.98 | ואלה |  |
| 6:21.99 | לחרפות |  |
| 6:21.100 | ולדראון |  |
| 6:21.101 | עולם | . |
| 6:21.102 | ולעג |  |
| 6:21.103 | על |  |
| 6:21.104 | מאמר |  |
| 6:21.105 | התנא |  |
| 6:21.106 | למוד |  |
| 6:21.107 | תורה |  |
| 6:21.108 | הרבה |  |
| 6:21.109 | ויתנו |  |
| 6:21.110 | לך |  |
| 6:21.111 | שכר |  |
| 6:21.112 | הרבה | . |
| 6:21.113 | ואף |  |
| 6:21.114 | אם |  |
| 6:21.115 | עוסק |  |
| 6:21.116 | בתורה |  |
| 6:21.117 | להבדיל |  |
| 6:21.118 | בין |  |
| 6:21.119 | הטהור |  |
| 6:21.120 | ובין |  |
| 6:21.121 | הטמא |  |
| 6:21.122 | כאשר |  |
| 6:21.123 | ציונו |  |
| 6:21.124 | ה | ' |
| 6:21.125 | אלהינו |  |
| 6:21.126 | להבדיל |  |
| 6:21.127 | בין |  |
| 6:21.128 | החיה |  |
| 6:21.129 | הנאכלת |  |
| 6:21.130 | ובין |  |
| 6:21.131 | החיה |  |
| 6:21.132 | אשר |  |
| 6:21.133 | לא |  |
| 6:21.134 | תאכל | , |
| 6:21.135 | ושאר |  |
| 6:21.136 | איסורי |  |
| 6:21.137 | מאכלות |  |
| 6:21.138 | הכל |  |
| 6:21.139 | נקרא |  |
| 6:21.140 | אצלו |  |
| 6:21.141 | עוסק |  |
| 6:21.142 | שלא |  |
| 6:21.143 | לשמה |  |
| 6:21.144 | כמו |  |
| 6:21.145 | שאמר |  |
| 6:21.146 | שם | - |
| 6:21.147 | לא |  |
| 6:21.148 | משתדלין |  |
| 6:21.149 | אלא |  |
| 6:21.150 | לגרמייהו |  |
| 6:21.151 | לדבתא |  |
| 6:21.152 | בשרא |  |
| 6:21.153 | וכו | '. |
| 6:21.154 | ולא |  |
| 6:21.155 | נקרא |  |
| 6:21.156 | עוסק |  |
| 6:21.157 | לשמה |  |
| 6:21.158 | אלא |  |
| 6:21.159 | העוסק |  |
| 6:21.160 | בתורה |  |
| 6:21.161 | כדי |  |
| 6:21.162 | לחבר |  |
| 6:21.163 | האלהות |  |
| 6:21.164 | הזכרים |  |
| 6:21.165 | בנקבות | , |
| 6:21.166 | אבא |  |
| 6:21.167 | עם |  |
| 6:21.168 | אימא | , |
| 6:21.169 | וזעיר |  |
| 6:21.170 | עם |  |
| 6:21.171 | נוקביה | . |
| 6:21.172 | לזוגם |  |
| 6:21.173 | יחד |  |
| 6:21.174 | והיו |  |
| 6:21.175 | לבשר |  |
| 6:21.176 | אחד | . |
| 6:21.177 | וכן |  |
| 6:21.178 | בקריאת |  |
| 6:21.179 | שמע |  |
| 6:21.180 | וציצית |  |
| 6:21.181 | ותפילין |  |
| 6:21.182 | ושאר |  |
| 6:21.183 | כל |  |
| 6:21.184 | המצות | . |
| 6:21.185 | ולכן |  |
| 6:21.186 | תקנו |  |
| 6:21.187 | המקובלים |  |
| 6:21.188 | לומר |  |
| 6:21.189 | קודם |  |
| 6:21.190 | כל |  |
| 6:21.191 | מצוה |  |
| 6:21.192 | שעושים |  |
| 6:21.193 | לשם |  |
| 6:21.194 | יחוד |  |
| 6:21.195 | קב"ה |  |
| 6:21.196 | ושכינתיה |  |
| 6:21.197 | וכו | ', |
| 6:21.198 | כי |  |
| 6:21.199 | תכלית |  |
| 6:21.200 | כוונת |  |
| 6:21.201 | הפילסוף |  |
| 6:21.202 | להכניס |  |
| 6:21.203 | בישראל |  |
| 6:21.204 | אמונת |  |
| 6:21.205 | רבוי |  |
| 6:21.206 | באלוהות |  |
| 6:21.207 | זכרים |  |
| 6:21.208 | ונקבות |  |
| 6:21.209 | ולעשות |  |
| 6:21.210 | חבורים |  |
| 6:21.211 | שלו |  |
| 6:21.212 | שהם |  |
| 6:21.213 | הזהר |  |
| 6:21.214 | והתיקונים |  |
| 6:21.215 | עיקר |  |
| 6:21.216 | לעסוק |  |
| 6:21.217 | בהם |  |
| 6:21.218 | חלף |  |
| 6:21.219 | המשנה |  |
| 6:21.220 | והתלמוד | . |
| 6:21.221 | ולשכח |  |
| 6:21.222 | מישראל |  |
| 6:21.223 | עסק |  |
| 6:21.224 | המשנה |  |
| 6:21.225 | והתלמוד |  |
| 6:21.226 | שכל |  |
| 6:21.227 | כללותיה |  |
| 6:21.228 | ופרטותיה |  |
| 6:21.229 | של |  |
| 6:21.230 | תורה |  |
| 6:21.231 | כלולים |  |
| 6:21.232 | בהם | . |
| 6:21.233 | ובכמה |  |
| 6:21.234 | מקומות |  |
| 6:21.235 | הוא |  |
| 6:21.236 | משבח |  |
| 6:21.237 | את |  |
| 6:21.238 | חבורו |  |
| 6:21.239 | זה |  |
| 6:21.240 | ומהם |  |
| 6:21.241 | בתיקוני |  |
| 6:21.242 | זוהר |  |
| 6:21.243 | חדש |  |
| 6:21.244 | דף |  |
| 6:21.245 | י"ו |  |
| 6:21.246 | ע"ב | . |
| 6:21.247 | ובנפת |  |
| 6:21.248 | אשר |  |
| 6:21.249 | יטיף |  |
| 6:21.250 | הפילסוף |  |
| 6:21.251 | הנז | ' |
| 6:21.252 | בשפתי |  |
| 6:21.253 | אמונה |  |
| 6:21.254 | שהיא |  |
| 6:21.255 | זרה |  |
| 6:21.256 | לאמונת |  |
| 6:21.257 | תורתינו |  |
| 6:21.258 | הקדושה |  |
| 6:21.259 | משך |  |
| 6:21.260 | לב |  |
| 6:21.261 | רבים |  |
| 6:21.262 | בחלקלקותיו | , |
| 6:21.263 | וישיאם |  |
| 6:21.264 | לעבוד |  |
| 6:21.265 | אלהים |  |
| 6:21.266 | אחרים |  |
| 6:21.267 | המה |  |
| 6:21.268 | הצורות |  |
| 6:21.269 | והפרצופים |  |
| 6:21.270 | אשר |  |
| 6:21.271 | בדה |  |
| 6:21.272 | מלבו |  |
| 6:21.273 | להטעותנו |  |
| 6:21.274 | בפירוש |  |
| 6:21.275 | עשר |  |
| 6:21.276 | ספירות |  |
| 6:21.277 | הנז | ' |
| 6:21.278 | בספר |  |
| 6:21.279 | יצירה |  |
| 6:21.280 | לחשבם |  |
| 6:21.281 | אלהות |  |
| 6:21.282 | ולעבוד |  |
| 6:21.283 | אותם |  |
| 6:21.284 | בעבור |  |
| 6:21.285 | שנשמותיהם |  |
| 6:21.286 | חצובים |  |
| 6:21.287 | מהאלהים |  |
| 6:21.288 | העליון |  |
| 6:21.289 | לפי |  |
| 6:21.290 | דעתו | ) |
| 6:21.291 | אשר |  |
| 6:21.292 | חלק |  |
| 6:21.293 | להם |  |
| 6:21.294 | מעצמותו |  |
| 6:21.295 | ומהותו |  |
| 6:21.296 | הם |  |
| 6:21.297 | ועיגוליהם |  |
| 6:21.298 | וגופם |  |
| 6:21.299 | ונשמתם |  |
| 6:21.300 | ולבושם | , |
| 6:21.301 | כולם |  |
| 6:21.302 | חוצבו |  |
| 6:21.303 | ממקור |  |
| 6:21.304 | אחד |  |
| 6:21.305 | הוא |  |
| 6:21.306 | הא"ם | ) |
| 6:21.307 | אשר |  |
| 6:21.308 | לפי |  |
| 6:21.309 | דעתם |  |
| 6:21.310 | אין |  |
| 6:21.311 | לו |  |
| 6:21.312 | שם |  |
| 6:21.313 | ונקודה |  |
| 6:21.314 | ולית |  |
| 6:21.315 | מחשבה |  |
| 6:21.316 | תפיסה |  |
| 6:21.317 | ביה | , |
| 6:21.318 | ולא |  |
| 6:21.319 | שייך |  |
| 6:21.320 | שום |  |
| 6:21.321 | עבודה |  |
| 6:21.322 | וקריאה |  |
| 6:21.323 | אליו | , |
| 6:21.324 | ולא |  |
| 6:21.325 | יענה |  |
| 6:21.326 | את |  |
| 6:21.327 | הקורא |  |
| 6:21.328 | אליו |  |
| 6:21.329 | שכבר |  |
| 6:21.330 | מינה |  |
| 6:21.331 | והפקיד |  |
| 6:21.332 | את |  |
| 6:21.333 | הפרצופים |  |
| 6:21.334 | על |  |
| 6:21.335 | כל |  |
| 6:21.336 | מלאכת |  |
| 6:21.337 | הבריאה | . |
| 6:21.338 | ועשה |  |
| 6:21.339 | את |  |
| 6:21.340 | אבא |  |
| 6:21.341 | בורא | , |
| 6:21.342 | והוא |  |
| 6:21.343 | אומר |  |
| 6:21.344 | ומצוה |  |
| 6:21.345 | לאימא |  |
| 6:21.346 | יהי |  |
| 6:21.347 | כך |  |
| 6:21.348 | וכך | , |
| 6:21.349 | והיא |  |
| 6:21.350 | עושה |  |
| 6:21.351 | כדבריו | , |
| 6:21.352 | כאומן |  |
| 6:21.353 | שעושה |  |
| 6:21.354 | מלאכת |  |
| 6:21.355 | בעל |  |
| 6:21.356 | הבית |  |
| 6:21.357 | באמונה | . |
| 6:21.358 | ואבא |  |
| 6:21.359 | ואמא |  |
| 6:21.360 | גם |  |
| 6:21.361 | הם |  |
| 6:21.362 | צוו |  |
| 6:21.363 | והפקידו |  |
| 6:21.364 | את |  |
| 6:21.365 | זעיר |  |
| 6:21.366 | אנפין |  |
| 6:21.367 | ונוקבתיה |  |
| 6:21.368 | להשגיח |  |
| 6:21.369 | ולרון |  |
| 6:21.370 | על |  |
| 6:21.371 | מעשה |  |
| 6:21.372 | התחתונים | . |
| 6:21.373 | מאן |  |
| 6:21.374 | דזכי |  |
| 6:21.375 | לדינא |  |
| 6:21.376 | לדינא | , |
| 6:21.377 | ומאן |  |
| 6:21.378 | דזכי |  |
| 6:21.379 | לרחמי |  |
| 6:21.380 | לרחמי | , |
| 6:21.381 | הכל |  |
| 6:21.382 | נעשה |  |
| 6:21.383 | כחפצו |  |
| 6:21.384 | ורצונו |  |
| 6:21.385 | של |  |
| 6:21.386 | זעיר |  |
| 6:21.387 | אנפין | . |
| 6:21.388 | וכל |  |
| 6:21.389 | דברי |  |
| 6:21.390 | המחבר |  |
| 6:21.391 | הזה |  |
| 6:21.392 | בנוים |  |
| 6:21.393 | על |  |
| 6:21.394 | הפילוסופיא |  |
| 6:21.395 | הקדומה |  |
| 6:21.396 | אשר |  |
| 6:21.397 | כבר |  |
| 6:21.398 | נמר |  |
| 6:21.399 | ריחה |  |
| 6:21.400 | וטעמה |  |
| 6:21.401 | מעת |  |
| 6:21.402 | אשר |  |
| 6:21.403 | נגלה |  |
| 6:21.404 | כבודו |  |
| 6:21.405 | ית | ' |
| 6:21.406 | על |  |
| 6:21.407 | הר |  |
| 6:21.408 | סיני |  |
| 6:21.409 | ואמר |  |
| 6:21.410 | אנכי |  |
| 6:21.411 | ה | ' |
| 6:21.412 | אלהיך | , |
| 6:21.413 | לא |  |
| 6:21.414 | יהיה |  |
| 6:21.415 | לך |  |
| 6:21.416 | אלהים |  |
| 6:21.417 | אחרים | . |
| 6:21.418 | רק |  |
| 6:21.419 | שאחרי |  |
| 6:21.420 | מות |  |
| 6:21.421 | יהושע |  |
| 6:21.422 | והזקנים |  |
| 6:21.423 | חזרו |  |
| 6:21.424 | להדיעות |  |
| 6:21.425 | האלו |  |
| 6:21.426 | הנפסדות |  |
| 6:21.427 | ויעבדו |  |
| 6:21.428 | את |  |
| 6:21.429 | הבעלים |  |
| 6:21.430 | והעשתרות |  |
| 6:21.431 | האלו |  |
| 6:21.432 | ויעשו |  |
| 6:21.433 | צלמיהם |  |
| 6:21.434 | כדי |  |
| 6:21.435 | להוריד |  |
| 6:21.436 | השפע |  |
| 6:21.437 | מהן |  |
| 6:21.438 | על |  |
| 6:21.439 | אותן |  |
| 6:21.440 | הצורות |  |
| 6:21.441 | והצלמים |  |
| 6:21.442 | לפי |  |
| 6:21.443 | דעתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

