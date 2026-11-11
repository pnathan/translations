# Translation pack: milhamot-hashem perek-7-06 (Milchamot ha-Shem 7:8–7:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-06`.

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
- ר"ל with a named authority as subject: "means to say".
- פ' + numeral after a book title: "chapter" when the numeral exceeds the
  parashot of that book (Midrash Rabbah 143 is a section/chapter number as
  the author gives it); otherwise "parashah". Note the choice.
- אם ... או after a full stop: translate as the author punctuates it only if
  the English has a main clause; otherwise join it to the previous sentence
  with a colon or render it as a direct alternative question.
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
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
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
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


## Milchamot ha-Shem 7:8

Hebrew (Wikisource): שים לבך אתה הקורא לדברי הגאון מהרש"ל ורמיזותיו שאמר שחלושי הראות לא יביטו. ולא ידעו מוצאו ומובאו של ספר הזהר. כי מיד בן נכר יצא ובא לחכמי טוליטולא כמו שכ' הר' חיד"א בספרו שם הגדולים. ומה שכתב וכוונתו, ר"ל שכל כוונתו להטעות עם ישראל להאמין באלהים רבים. ומה שכ' פקח עיניך וראה אם יש בזה ממש וכו' כלומר אם אמת דבריו שכן אמר רשב"י לא אשתמיט שום רמז בתלמוד בבלי וירושלמי וספרא ספרי. ושאלתות ופסיקתא ממה ששמעו מרשב"י. כי חבר היה לחכמי המשנה ואי אפשר שלא ישמעו ממנו דבר מכל הנאמר בזהר ותיקונים. אלא ודאי הכל שקר וכזב. ומי שלא ידע להשיג תורה על נכון בתלמוד ופוסקים יבוא לקצץ בנטיעות להאמין שאלהים רבים וכמו שמסיים בסוף דבריו והוי כמו מינות. וכמאמר שלמה המלך ע"ה פתי יאמין לכל דבר. וערום יבין לאשורו. עוד כתב רביד הזהב ואע"ג דהש"ך כ' דאין דברי המרש"ל מוכרחים, וגם הדמשק אליעזר הביאם הכנה"ג המה עלוי דמה מינות שייך במה שהאדם מחמיר על עצמו וכו' ע"כ לעינינו.

| id | word | punctuation after |
|---|---|---|
| 7:8.1 | שים |  |
| 7:8.2 | לבך |  |
| 7:8.3 | אתה |  |
| 7:8.4 | הקורא |  |
| 7:8.5 | לדברי |  |
| 7:8.6 | הגאון |  |
| 7:8.7 | מהרש"ל |  |
| 7:8.8 | ורמיזותיו |  |
| 7:8.9 | שאמר |  |
| 7:8.10 | שחלושי |  |
| 7:8.11 | הראות |  |
| 7:8.12 | לא |  |
| 7:8.13 | יביטו | . |
| 7:8.14 | ולא |  |
| 7:8.15 | ידעו |  |
| 7:8.16 | מוצאו |  |
| 7:8.17 | ומובאו |  |
| 7:8.18 | של |  |
| 7:8.19 | ספר |  |
| 7:8.20 | הזהר | . |
| 7:8.21 | כי |  |
| 7:8.22 | מיד |  |
| 7:8.23 | בן |  |
| 7:8.24 | נכר |  |
| 7:8.25 | יצא |  |
| 7:8.26 | ובא |  |
| 7:8.27 | לחכמי |  |
| 7:8.28 | טוליטולא |  |
| 7:8.29 | כמו |  |
| 7:8.30 | שכ | ' |
| 7:8.31 | הר | ' |
| 7:8.32 | חיד"א |  |
| 7:8.33 | בספרו |  |
| 7:8.34 | שם |  |
| 7:8.35 | הגדולים | . |
| 7:8.36 | ומה |  |
| 7:8.37 | שכתב |  |
| 7:8.38 | וכוונתו | , |
| 7:8.39 | ר"ל |  |
| 7:8.40 | שכל |  |
| 7:8.41 | כוונתו |  |
| 7:8.42 | להטעות |  |
| 7:8.43 | עם |  |
| 7:8.44 | ישראל |  |
| 7:8.45 | להאמין |  |
| 7:8.46 | באלהים |  |
| 7:8.47 | רבים | . |
| 7:8.48 | ומה |  |
| 7:8.49 | שכ | ' |
| 7:8.50 | פקח |  |
| 7:8.51 | עיניך |  |
| 7:8.52 | וראה |  |
| 7:8.53 | אם |  |
| 7:8.54 | יש |  |
| 7:8.55 | בזה |  |
| 7:8.56 | ממש |  |
| 7:8.57 | וכו | ' |
| 7:8.58 | כלומר |  |
| 7:8.59 | אם |  |
| 7:8.60 | אמת |  |
| 7:8.61 | דבריו |  |
| 7:8.62 | שכן |  |
| 7:8.63 | אמר |  |
| 7:8.64 | רשב"י |  |
| 7:8.65 | לא |  |
| 7:8.66 | אשתמיט |  |
| 7:8.67 | שום |  |
| 7:8.68 | רמז |  |
| 7:8.69 | בתלמוד |  |
| 7:8.70 | בבלי |  |
| 7:8.71 | וירושלמי |  |
| 7:8.72 | וספרא |  |
| 7:8.73 | ספרי | . |
| 7:8.74 | ושאלתות |  |
| 7:8.75 | ופסיקתא |  |
| 7:8.76 | ממה |  |
| 7:8.77 | ששמעו |  |
| 7:8.78 | מרשב"י | . |
| 7:8.79 | כי |  |
| 7:8.80 | חבר |  |
| 7:8.81 | היה |  |
| 7:8.82 | לחכמי |  |
| 7:8.83 | המשנה |  |
| 7:8.84 | ואי |  |
| 7:8.85 | אפשר |  |
| 7:8.86 | שלא |  |
| 7:8.87 | ישמעו |  |
| 7:8.88 | ממנו |  |
| 7:8.89 | דבר |  |
| 7:8.90 | מכל |  |
| 7:8.91 | הנאמר |  |
| 7:8.92 | בזהר |  |
| 7:8.93 | ותיקונים | . |
| 7:8.94 | אלא |  |
| 7:8.95 | ודאי |  |
| 7:8.96 | הכל |  |
| 7:8.97 | שקר |  |
| 7:8.98 | וכזב | . |
| 7:8.99 | ומי |  |
| 7:8.100 | שלא |  |
| 7:8.101 | ידע |  |
| 7:8.102 | להשיג |  |
| 7:8.103 | תורה |  |
| 7:8.104 | על |  |
| 7:8.105 | נכון |  |
| 7:8.106 | בתלמוד |  |
| 7:8.107 | ופוסקים |  |
| 7:8.108 | יבוא |  |
| 7:8.109 | לקצץ |  |
| 7:8.110 | בנטיעות |  |
| 7:8.111 | להאמין |  |
| 7:8.112 | שאלהים |  |
| 7:8.113 | רבים |  |
| 7:8.114 | וכמו |  |
| 7:8.115 | שמסיים |  |
| 7:8.116 | בסוף |  |
| 7:8.117 | דבריו |  |
| 7:8.118 | והוי |  |
| 7:8.119 | כמו |  |
| 7:8.120 | מינות | . |
| 7:8.121 | וכמאמר |  |
| 7:8.122 | שלמה |  |
| 7:8.123 | המלך |  |
| 7:8.124 | ע"ה |  |
| 7:8.125 | פתי |  |
| 7:8.126 | יאמין |  |
| 7:8.127 | לכל |  |
| 7:8.128 | דבר | . |
| 7:8.129 | וערום |  |
| 7:8.130 | יבין |  |
| 7:8.131 | לאשורו | . |
| 7:8.132 | עוד |  |
| 7:8.133 | כתב |  |
| 7:8.134 | רביד |  |
| 7:8.135 | הזהב |  |
| 7:8.136 | ואע"ג |  |
| 7:8.137 | דהש"ך |  |
| 7:8.138 | כ | ' |
| 7:8.139 | דאין |  |
| 7:8.140 | דברי |  |
| 7:8.141 | המרש"ל |  |
| 7:8.142 | מוכרחים | , |
| 7:8.143 | וגם |  |
| 7:8.144 | הדמשק |  |
| 7:8.145 | אליעזר |  |
| 7:8.146 | הביאם |  |
| 7:8.147 | הכנה"ג |  |
| 7:8.148 | המה |  |
| 7:8.149 | עלוי |  |
| 7:8.150 | דמה |  |
| 7:8.151 | מינות |  |
| 7:8.152 | שייך |  |
| 7:8.153 | במה |  |
| 7:8.154 | שהאדם |  |
| 7:8.155 | מחמיר |  |
| 7:8.156 | על |  |
| 7:8.157 | עצמו |  |
| 7:8.158 | וכו | ' |
| 7:8.159 | ע"כ |  |
| 7:8.160 | לעינינו | . |


## Milchamot ha-Shem 7:9

Hebrew (Wikisource): הנה מה מאד יתפלא הקורא הנכון על דברי הש"ך שכתב על דברי הגאון מהרש"ל שאין דבריו מוכרחים. והוא בעצמו לא ירד לסוף דעתו, וכל שכן על תמיהת הדמשק אליעזר איך מלאו לבו לתמוה על אדם גדול כמהרש"ל אשר שט בים התלמוד וספרא וספרי ושאלתות ועיניו ראו ועפעפיו בחנו, כי דרך עצב בדבר הזהר והמקובלים החדשים, והוא לבדו ידע מהיכן מוצאו ומובאו, כי מיד עכו"ם יצא ובא אלינו, להסית ולהדיח אותנו להאמין באלהים רבים ולעבוד אלהים אחרים. סבות רבות מחודשות נבראות המה הצורות והפרצופים שהמציא מחבר הזהר ברעיונותיו, וזהו שכ' רש"ל שלא ידעו מוצאו ומובאו וכוונתו. והוסיף עוד לומר פקח עיניך וראה אם יש ממש בדבריו וממקור נאמן יצאו דברים כאלה רבוי באלהות בפרצופים שונים ומלכות וממשלה להם, לא לשתמיט שום רמז מזה בתלמוד בבלי וירושלמי וספרא וספרי ופסיקתא ושאלתות, אלא ודאי לא מישראל יצאו הדברים האלה וחיצונים הם, ולכך אמר לכן אהובי אל תלך בדרך אתם וכו', ואף מי שדעתו להחמיר על עצמו, הואיל ואין לזאת החומרא שום רמז בתלמוד והוא מקפיד בדבר שלא הקפידו בו חכמים הוי כמו מינות כדתנן בצבועים איני עובר אף בלבנים לא יעבור וכמ"ש התוספת י"ט ויובא לקמן סי כ"ה. אכן עוד תגדל הפליאה על הדמשק אליעזר באמרו מה מינות שייך במה שאדם מחמיר על עצמו, כי ידוע הוא ומפורסם בכל האיסורים והחומרות והמנהגים שחידשו הזהר והמקובלים החדשים הבאים אחריו, כולהו אית בהו טעמא לפי שטתם באמונת אלוהות רבים, והרש"ל ז"ל הרגיש בזרותם כי עיניו יחזו ועפעפיו יבחנו, כי בילדי נכרים עובדי ע"ז ישפיקו, ודרך עוצב יש בטעמי האיסורים והחומרות והמנהגים שחידשו שלא כדרך הגמרא ואית בהו לתא דע"ז. ולפי שהראשונים כר' תם ן' יחיא ומהר"י אלבו והר' מהרש"ל וכיוצא בהם יראו לנפשם מאד פן יפגעו להם הרודפים הקנאים כאיזבל לאליהו, לכן תלו הטעות בלומדים והזהירום ממנה, ולפי האמת הברור כי אותן החומרות מינות גמורה. והא לך דוגמא מה שכתב הרב משנת חסידים מסכת ליל פסח פרק י"א שכשאומר בהגדה דם ואש ותימרות עשן וכן בהזכירו עשר מכות בפרט ובכלל ישפוך מהכוס לתוך כלי שבור הרומז אל הקליפה הנקרא ארור, ועל ידי זה תפסוק המלכות (שהיא דשתנא) מדם נדתה ותהיה ראויה לזווג ונקתה ונזרעה זרע, כי היין הנשאר בה הוא יין המשמח דם טהור המתהפך לחלב אחר כך ע"כ. והחכם עיניו בראשו להשכיל, דכל כי האי גונא מנסך הוא לע"ז שמנסך מכוסו לקליפה הנקראת ארור כדי לתקן המלכות ויינו אסור בהנאה מן התורה שהרי ניסך ממנו לקליפה הנקראת ארור כדי לתקן, והמנסך היא אחת מארבע עבודות שחייב עליהן בין שדרך עבודתה בכך, בין שאין דרך עבודתה בכך כדתנן בסנהדרין פ' ארבע מיתות, (סנהדרין דף ס' ע"ב). וכן מה שמקפידין לחלוץ התפלין קודם מוסף ראש חדש טעמא הוא לפי שהתפילין הם בז"א שלפי דבריהם עליו אמרו רז"ל במס' ברכות שהקב"ה מניח תפילין, ובמוסף ר"ח הוא בא אור גדול מעתיק שהוא האלוה העליון שאינו מניח תפלין, כי התפלין הם שייכים לקצר אפים שהוא זעיר אנפין, וכי האיי גונא מינות הוא להאמין ברשויות הרבה, זה מניח תפילין וזה אין מניח תפילין, והרי אין הקב"ה אחד.

| id | word | punctuation after |
|---|---|---|
| 7:9.1 | הנה |  |
| 7:9.2 | מה |  |
| 7:9.3 | מאד |  |
| 7:9.4 | יתפלא |  |
| 7:9.5 | הקורא |  |
| 7:9.6 | הנכון |  |
| 7:9.7 | על |  |
| 7:9.8 | דברי |  |
| 7:9.9 | הש"ך |  |
| 7:9.10 | שכתב |  |
| 7:9.11 | על |  |
| 7:9.12 | דברי |  |
| 7:9.13 | הגאון |  |
| 7:9.14 | מהרש"ל |  |
| 7:9.15 | שאין |  |
| 7:9.16 | דבריו |  |
| 7:9.17 | מוכרחים | . |
| 7:9.18 | והוא |  |
| 7:9.19 | בעצמו |  |
| 7:9.20 | לא |  |
| 7:9.21 | ירד |  |
| 7:9.22 | לסוף |  |
| 7:9.23 | דעתו | , |
| 7:9.24 | וכל |  |
| 7:9.25 | שכן |  |
| 7:9.26 | על |  |
| 7:9.27 | תמיהת |  |
| 7:9.28 | הדמשק |  |
| 7:9.29 | אליעזר |  |
| 7:9.30 | איך |  |
| 7:9.31 | מלאו |  |
| 7:9.32 | לבו |  |
| 7:9.33 | לתמוה |  |
| 7:9.34 | על |  |
| 7:9.35 | אדם |  |
| 7:9.36 | גדול |  |
| 7:9.37 | כמהרש"ל |  |
| 7:9.38 | אשר |  |
| 7:9.39 | שט |  |
| 7:9.40 | בים |  |
| 7:9.41 | התלמוד |  |
| 7:9.42 | וספרא |  |
| 7:9.43 | וספרי |  |
| 7:9.44 | ושאלתות |  |
| 7:9.45 | ועיניו |  |
| 7:9.46 | ראו |  |
| 7:9.47 | ועפעפיו |  |
| 7:9.48 | בחנו | , |
| 7:9.49 | כי |  |
| 7:9.50 | דרך |  |
| 7:9.51 | עצב |  |
| 7:9.52 | בדבר |  |
| 7:9.53 | הזהר |  |
| 7:9.54 | והמקובלים |  |
| 7:9.55 | החדשים | , |
| 7:9.56 | והוא |  |
| 7:9.57 | לבדו |  |
| 7:9.58 | ידע |  |
| 7:9.59 | מהיכן |  |
| 7:9.60 | מוצאו |  |
| 7:9.61 | ומובאו | , |
| 7:9.62 | כי |  |
| 7:9.63 | מיד |  |
| 7:9.64 | עכו"ם |  |
| 7:9.65 | יצא |  |
| 7:9.66 | ובא |  |
| 7:9.67 | אלינו | , |
| 7:9.68 | להסית |  |
| 7:9.69 | ולהדיח |  |
| 7:9.70 | אותנו |  |
| 7:9.71 | להאמין |  |
| 7:9.72 | באלהים |  |
| 7:9.73 | רבים |  |
| 7:9.74 | ולעבוד |  |
| 7:9.75 | אלהים |  |
| 7:9.76 | אחרים | . |
| 7:9.77 | סבות |  |
| 7:9.78 | רבות |  |
| 7:9.79 | מחודשות |  |
| 7:9.80 | נבראות |  |
| 7:9.81 | המה |  |
| 7:9.82 | הצורות |  |
| 7:9.83 | והפרצופים |  |
| 7:9.84 | שהמציא |  |
| 7:9.85 | מחבר |  |
| 7:9.86 | הזהר |  |
| 7:9.87 | ברעיונותיו | , |
| 7:9.88 | וזהו |  |
| 7:9.89 | שכ | ' |
| 7:9.90 | רש"ל |  |
| 7:9.91 | שלא |  |
| 7:9.92 | ידעו |  |
| 7:9.93 | מוצאו |  |
| 7:9.94 | ומובאו |  |
| 7:9.95 | וכוונתו | . |
| 7:9.96 | והוסיף |  |
| 7:9.97 | עוד |  |
| 7:9.98 | לומר |  |
| 7:9.99 | פקח |  |
| 7:9.100 | עיניך |  |
| 7:9.101 | וראה |  |
| 7:9.102 | אם |  |
| 7:9.103 | יש |  |
| 7:9.104 | ממש |  |
| 7:9.105 | בדבריו |  |
| 7:9.106 | וממקור |  |
| 7:9.107 | נאמן |  |
| 7:9.108 | יצאו |  |
| 7:9.109 | דברים |  |
| 7:9.110 | כאלה |  |
| 7:9.111 | רבוי |  |
| 7:9.112 | באלהות |  |
| 7:9.113 | בפרצופים |  |
| 7:9.114 | שונים |  |
| 7:9.115 | ומלכות |  |
| 7:9.116 | וממשלה |  |
| 7:9.117 | להם | , |
| 7:9.118 | לא |  |
| 7:9.119 | לשתמיט |  |
| 7:9.120 | שום |  |
| 7:9.121 | רמז |  |
| 7:9.122 | מזה |  |
| 7:9.123 | בתלמוד |  |
| 7:9.124 | בבלי |  |
| 7:9.125 | וירושלמי |  |
| 7:9.126 | וספרא |  |
| 7:9.127 | וספרי |  |
| 7:9.128 | ופסיקתא |  |
| 7:9.129 | ושאלתות | , |
| 7:9.130 | אלא |  |
| 7:9.131 | ודאי |  |
| 7:9.132 | לא |  |
| 7:9.133 | מישראל |  |
| 7:9.134 | יצאו |  |
| 7:9.135 | הדברים |  |
| 7:9.136 | האלה |  |
| 7:9.137 | וחיצונים |  |
| 7:9.138 | הם | , |
| 7:9.139 | ולכך |  |
| 7:9.140 | אמר |  |
| 7:9.141 | לכן |  |
| 7:9.142 | אהובי |  |
| 7:9.143 | אל |  |
| 7:9.144 | תלך |  |
| 7:9.145 | בדרך |  |
| 7:9.146 | אתם |  |
| 7:9.147 | וכו | ', |
| 7:9.148 | ואף |  |
| 7:9.149 | מי |  |
| 7:9.150 | שדעתו |  |
| 7:9.151 | להחמיר |  |
| 7:9.152 | על |  |
| 7:9.153 | עצמו | , |
| 7:9.154 | הואיל |  |
| 7:9.155 | ואין |  |
| 7:9.156 | לזאת |  |
| 7:9.157 | החומרא |  |
| 7:9.158 | שום |  |
| 7:9.159 | רמז |  |
| 7:9.160 | בתלמוד |  |
| 7:9.161 | והוא |  |
| 7:9.162 | מקפיד |  |
| 7:9.163 | בדבר |  |
| 7:9.164 | שלא |  |
| 7:9.165 | הקפידו |  |
| 7:9.166 | בו |  |
| 7:9.167 | חכמים |  |
| 7:9.168 | הוי |  |
| 7:9.169 | כמו |  |
| 7:9.170 | מינות |  |
| 7:9.171 | כדתנן |  |
| 7:9.172 | בצבועים |  |
| 7:9.173 | איני |  |
| 7:9.174 | עובר |  |
| 7:9.175 | אף |  |
| 7:9.176 | בלבנים |  |
| 7:9.177 | לא |  |
| 7:9.178 | יעבור |  |
| 7:9.179 | וכמ"ש |  |
| 7:9.180 | התוספת |  |
| 7:9.181 | י"ט |  |
| 7:9.182 | ויובא |  |
| 7:9.183 | לקמן |  |
| 7:9.184 | סי |  |
| 7:9.185 | כ"ה | . |
| 7:9.186 | אכן |  |
| 7:9.187 | עוד |  |
| 7:9.188 | תגדל |  |
| 7:9.189 | הפליאה |  |
| 7:9.190 | על |  |
| 7:9.191 | הדמשק |  |
| 7:9.192 | אליעזר |  |
| 7:9.193 | באמרו |  |
| 7:9.194 | מה |  |
| 7:9.195 | מינות |  |
| 7:9.196 | שייך |  |
| 7:9.197 | במה |  |
| 7:9.198 | שאדם |  |
| 7:9.199 | מחמיר |  |
| 7:9.200 | על |  |
| 7:9.201 | עצמו | , |
| 7:9.202 | כי |  |
| 7:9.203 | ידוע |  |
| 7:9.204 | הוא |  |
| 7:9.205 | ומפורסם |  |
| 7:9.206 | בכל |  |
| 7:9.207 | האיסורים |  |
| 7:9.208 | והחומרות |  |
| 7:9.209 | והמנהגים |  |
| 7:9.210 | שחידשו |  |
| 7:9.211 | הזהר |  |
| 7:9.212 | והמקובלים |  |
| 7:9.213 | החדשים |  |
| 7:9.214 | הבאים |  |
| 7:9.215 | אחריו | , |
| 7:9.216 | כולהו |  |
| 7:9.217 | אית |  |
| 7:9.218 | בהו |  |
| 7:9.219 | טעמא |  |
| 7:9.220 | לפי |  |
| 7:9.221 | שטתם |  |
| 7:9.222 | באמונת |  |
| 7:9.223 | אלוהות |  |
| 7:9.224 | רבים | , |
| 7:9.225 | והרש"ל |  |
| 7:9.226 | ז"ל |  |
| 7:9.227 | הרגיש |  |
| 7:9.228 | בזרותם |  |
| 7:9.229 | כי |  |
| 7:9.230 | עיניו |  |
| 7:9.231 | יחזו |  |
| 7:9.232 | ועפעפיו |  |
| 7:9.233 | יבחנו | , |
| 7:9.234 | כי |  |
| 7:9.235 | בילדי |  |
| 7:9.236 | נכרים |  |
| 7:9.237 | עובדי |  |
| 7:9.238 | ע"ז |  |
| 7:9.239 | ישפיקו | , |
| 7:9.240 | ודרך |  |
| 7:9.241 | עוצב |  |
| 7:9.242 | יש |  |
| 7:9.243 | בטעמי |  |
| 7:9.244 | האיסורים |  |
| 7:9.245 | והחומרות |  |
| 7:9.246 | והמנהגים |  |
| 7:9.247 | שחידשו |  |
| 7:9.248 | שלא |  |
| 7:9.249 | כדרך |  |
| 7:9.250 | הגמרא |  |
| 7:9.251 | ואית |  |
| 7:9.252 | בהו |  |
| 7:9.253 | לתא |  |
| 7:9.254 | דע"ז | . |
| 7:9.255 | ולפי |  |
| 7:9.256 | שהראשונים |  |
| 7:9.257 | כר | ' |
| 7:9.258 | תם |  |
| 7:9.259 | ן | ' |
| 7:9.260 | יחיא |  |
| 7:9.261 | ומהר"י |  |
| 7:9.262 | אלבו |  |
| 7:9.263 | והר | ' |
| 7:9.264 | מהרש"ל |  |
| 7:9.265 | וכיוצא |  |
| 7:9.266 | בהם |  |
| 7:9.267 | יראו |  |
| 7:9.268 | לנפשם |  |
| 7:9.269 | מאד |  |
| 7:9.270 | פן |  |
| 7:9.271 | יפגעו |  |
| 7:9.272 | להם |  |
| 7:9.273 | הרודפים |  |
| 7:9.274 | הקנאים |  |
| 7:9.275 | כאיזבל |  |
| 7:9.276 | לאליהו | , |
| 7:9.277 | לכן |  |
| 7:9.278 | תלו |  |
| 7:9.279 | הטעות |  |
| 7:9.280 | בלומדים |  |
| 7:9.281 | והזהירום |  |
| 7:9.282 | ממנה | , |
| 7:9.283 | ולפי |  |
| 7:9.284 | האמת |  |
| 7:9.285 | הברור |  |
| 7:9.286 | כי |  |
| 7:9.287 | אותן |  |
| 7:9.288 | החומרות |  |
| 7:9.289 | מינות |  |
| 7:9.290 | גמורה | . |
| 7:9.291 | והא |  |
| 7:9.292 | לך |  |
| 7:9.293 | דוגמא |  |
| 7:9.294 | מה |  |
| 7:9.295 | שכתב |  |
| 7:9.296 | הרב |  |
| 7:9.297 | משנת |  |
| 7:9.298 | חסידים |  |
| 7:9.299 | מסכת |  |
| 7:9.300 | ליל |  |
| 7:9.301 | פסח |  |
| 7:9.302 | פרק |  |
| 7:9.303 | י"א |  |
| 7:9.304 | שכשאומר |  |
| 7:9.305 | בהגדה |  |
| 7:9.306 | דם |  |
| 7:9.307 | ואש |  |
| 7:9.308 | ותימרות |  |
| 7:9.309 | עשן |  |
| 7:9.310 | וכן |  |
| 7:9.311 | בהזכירו |  |
| 7:9.312 | עשר |  |
| 7:9.313 | מכות |  |
| 7:9.314 | בפרט |  |
| 7:9.315 | ובכלל |  |
| 7:9.316 | ישפוך |  |
| 7:9.317 | מהכוס |  |
| 7:9.318 | לתוך |  |
| 7:9.319 | כלי |  |
| 7:9.320 | שבור |  |
| 7:9.321 | הרומז |  |
| 7:9.322 | אל |  |
| 7:9.323 | הקליפה |  |
| 7:9.324 | הנקרא |  |
| 7:9.325 | ארור | , |
| 7:9.326 | ועל |  |
| 7:9.327 | ידי |  |
| 7:9.328 | זה |  |
| 7:9.329 | תפסוק |  |
| 7:9.330 | המלכות |  |
| 7:9.331 | שהיא |  |
| 7:9.332 | דשתנא | ) |
| 7:9.333 | מדם |  |
| 7:9.334 | נדתה |  |
| 7:9.335 | ותהיה |  |
| 7:9.336 | ראויה |  |
| 7:9.337 | לזווג |  |
| 7:9.338 | ונקתה |  |
| 7:9.339 | ונזרעה |  |
| 7:9.340 | זרע | , |
| 7:9.341 | כי |  |
| 7:9.342 | היין |  |
| 7:9.343 | הנשאר |  |
| 7:9.344 | בה |  |
| 7:9.345 | הוא |  |
| 7:9.346 | יין |  |
| 7:9.347 | המשמח |  |
| 7:9.348 | דם |  |
| 7:9.349 | טהור |  |
| 7:9.350 | המתהפך |  |
| 7:9.351 | לחלב |  |
| 7:9.352 | אחר |  |
| 7:9.353 | כך |  |
| 7:9.354 | ע"כ | . |
| 7:9.355 | והחכם |  |
| 7:9.356 | עיניו |  |
| 7:9.357 | בראשו |  |
| 7:9.358 | להשכיל | , |
| 7:9.359 | דכל |  |
| 7:9.360 | כי |  |
| 7:9.361 | האי |  |
| 7:9.362 | גונא |  |
| 7:9.363 | מנסך |  |
| 7:9.364 | הוא |  |
| 7:9.365 | לע"ז |  |
| 7:9.366 | שמנסך |  |
| 7:9.367 | מכוסו |  |
| 7:9.368 | לקליפה |  |
| 7:9.369 | הנקראת |  |
| 7:9.370 | ארור |  |
| 7:9.371 | כדי |  |
| 7:9.372 | לתקן |  |
| 7:9.373 | המלכות |  |
| 7:9.374 | ויינו |  |
| 7:9.375 | אסור |  |
| 7:9.376 | בהנאה |  |
| 7:9.377 | מן |  |
| 7:9.378 | התורה |  |
| 7:9.379 | שהרי |  |
| 7:9.380 | ניסך |  |
| 7:9.381 | ממנו |  |
| 7:9.382 | לקליפה |  |
| 7:9.383 | הנקראת |  |
| 7:9.384 | ארור |  |
| 7:9.385 | כדי |  |
| 7:9.386 | לתקן | , |
| 7:9.387 | והמנסך |  |
| 7:9.388 | היא |  |
| 7:9.389 | אחת |  |
| 7:9.390 | מארבע |  |
| 7:9.391 | עבודות |  |
| 7:9.392 | שחייב |  |
| 7:9.393 | עליהן |  |
| 7:9.394 | בין |  |
| 7:9.395 | שדרך |  |
| 7:9.396 | עבודתה |  |
| 7:9.397 | בכך | , |
| 7:9.398 | בין |  |
| 7:9.399 | שאין |  |
| 7:9.400 | דרך |  |
| 7:9.401 | עבודתה |  |
| 7:9.402 | בכך |  |
| 7:9.403 | כדתנן |  |
| 7:9.404 | בסנהדרין |  |
| 7:9.405 | פ | ' |
| 7:9.406 | ארבע |  |
| 7:9.407 | מיתות | , |
| 7:9.408 | סנהדרין |  |
| 7:9.409 | דף |  |
| 7:9.410 | ס | ' |
| 7:9.411 | ע"ב | ). |
| 7:9.412 | וכן |  |
| 7:9.413 | מה |  |
| 7:9.414 | שמקפידין |  |
| 7:9.415 | לחלוץ |  |
| 7:9.416 | התפלין |  |
| 7:9.417 | קודם |  |
| 7:9.418 | מוסף |  |
| 7:9.419 | ראש |  |
| 7:9.420 | חדש |  |
| 7:9.421 | טעמא |  |
| 7:9.422 | הוא |  |
| 7:9.423 | לפי |  |
| 7:9.424 | שהתפילין |  |
| 7:9.425 | הם |  |
| 7:9.426 | בז"א |  |
| 7:9.427 | שלפי |  |
| 7:9.428 | דבריהם |  |
| 7:9.429 | עליו |  |
| 7:9.430 | אמרו |  |
| 7:9.431 | רז"ל |  |
| 7:9.432 | במס | ' |
| 7:9.433 | ברכות |  |
| 7:9.434 | שהקב"ה |  |
| 7:9.435 | מניח |  |
| 7:9.436 | תפילין | , |
| 7:9.437 | ובמוסף |  |
| 7:9.438 | ר"ח |  |
| 7:9.439 | הוא |  |
| 7:9.440 | בא |  |
| 7:9.441 | אור |  |
| 7:9.442 | גדול |  |
| 7:9.443 | מעתיק |  |
| 7:9.444 | שהוא |  |
| 7:9.445 | האלוה |  |
| 7:9.446 | העליון |  |
| 7:9.447 | שאינו |  |
| 7:9.448 | מניח |  |
| 7:9.449 | תפלין | , |
| 7:9.450 | כי |  |
| 7:9.451 | התפלין |  |
| 7:9.452 | הם |  |
| 7:9.453 | שייכים |  |
| 7:9.454 | לקצר |  |
| 7:9.455 | אפים |  |
| 7:9.456 | שהוא |  |
| 7:9.457 | זעיר |  |
| 7:9.458 | אנפין | , |
| 7:9.459 | וכי |  |
| 7:9.460 | האיי |  |
| 7:9.461 | גונא |  |
| 7:9.462 | מינות |  |
| 7:9.463 | הוא |  |
| 7:9.464 | להאמין |  |
| 7:9.465 | ברשויות |  |
| 7:9.466 | הרבה | , |
| 7:9.467 | זה |  |
| 7:9.468 | מניח |  |
| 7:9.469 | תפילין |  |
| 7:9.470 | וזה |  |
| 7:9.471 | אין |  |
| 7:9.472 | מניח |  |
| 7:9.473 | תפילין | , |
| 7:9.474 | והרי |  |
| 7:9.475 | אין |  |
| 7:9.476 | הקב"ה |  |
| 7:9.477 | אחד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

