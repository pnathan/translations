# Review pack: milhamot-hashem perek-1-03 (Milchamot ha-Shem 1:4–1:4)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-1-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-03`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- משכילים / השכלה: "the educated" / "enlightenment" in the sense of the
  Haskalah only where the paragraph speaks of that movement; otherwise
  "those of understanding".
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
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
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
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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
- הינ"ץ: keep "H.N.Ts." with a token note giving the conjecture (ה + ינ"ץ, "may
  his Rock protect him", of the Imam); no person is named in the English.
- ב"ע: "in Aramaic" (בארמית) where it is set against בערבי; the other readings in
  `readings`.
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
- גיזומיו: by context, "exaggerations" if the sentence is about claims, else
  "trimmings"; the other in a token note.
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.
- הרא"ש: "the Rosh"; ר"ת: "Rabbeinu Tam"; הגאון alone: "the Gaon", unidentified.
  (Acronym authorities keep the acronym, by the ruling above.)

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


## Milchamot ha-Shem 1:4

Hebrew (Wikisource): ולתקופת השנה השלישית פרצה המלחמה נגד הלמודים הטובים והסדרים הנעימים כמצווה עלינו מרז״ל. ויוציאו המתנגדים דבה רעה לאמר כי על ידי זה יצאו הילדים ח״ו לתרבות רעה, והרבה מהטפשים האמינו לשקריהם והוציאו את ילדיהם מבית הספר. וגם הבנים הסכלים אשר יצאו מבית הספר עוינים את מוריהם כי לעניות דעתם הם מינים ויוצאים מן הכלל אך המשכילים שבהם יודעים ומכירים כי שקר בימין מוציאי הדבה הרעה הלזו, ודאגו מאד על שהוציאום אבותם מבית הספר בסכלותם! וכאשר התעו רבים חזקה יד החשכים ובמלשינות שקר הלשינו על המלמדים אצל הממשלה שהם חפצים להמריד את יושבי הארץ. ושהם כותבי פלסתר על הממשלה אל ממשלת אנגליא לחסות בצל ממשלתה וגם אל הינ״ץ ועבור זה המלשינות הובלו המלמדים אל בית הכלא וגם גורשו מבית הספר. והביאו תחתיהם מלמד אחר אשר לא ידע לא עברית על בריה ולא ערבית כתקנה. ואין עליו ללמדם רק פרשת השבוע עם תרגום אונקלוס ותרגום רס״ג בערבי. והפטרה עם תרגום יונתן ב״ע באין הבין כלום כשאר המלמדים! ולימים מספר בטל בית הספר לגמרי! כל זאת באתנו על אשר לא מחינו בכבוד קונינו בעת שהיינו שומעים מהלומדים בזהר ושאר מקובלים דברים שאסור לשמען ומחרישים להם כמו שאמרו רז״ל דר״א נתפש למינות וכו' שמא דבר מינות וכו' ונתקיים בנו מאמר רז״ל דמודי להון נפיל בידיהון דמתרחץ עליהון דיליה דילהון ע״כ! לכן עזבנו דרכינו בלימוד השפות העברית והערבית בדקדוקן על כורחנו מפני חמת המציק על ידי מלשינות שקר ונפן אל דרך הקדש כבוד אבינו שבשמים בבטול דעות הפילסוף נביא השקר מחבר הזהר אשר אסף בחפניו דעות הפלסופיא הקדומה עובדי האלילים תמונות וכוחות נעלמים אשר בדו מלבם. ומחבר הזהר תפס דעותיהם וחשב להכניסם באמונת תורתינו הקדושה, וקרא את שמו בשם רשב״י התנא בשקר ואמר שנגלה לו הקב״ה הנקרא אצלו עתיקא קדישא וגילה לו רזין טמירין. ותלה נבואתו ורעיונותיו ברשב״י התנא ע״ה. ועשה את חבורו הזהר ואמר שהעליונים הסכימו לחבורו לפי שראה הפילסוף הזה כי המשנה ושני התלמודים ומדרשי רז״ל האמתיים עומדים נגדו באמונת אחדותו ית' לכן חשב לפסלן ובזה ולעג להם ולכל העוסקים בהם, וכינה אותם בשמות של זלזול ופחיתות, כאשר יראה המעיין בפנים! והבטיח שלא יזכו ישראל לגאולה עד אשר ישובו מדרכם מלעסוק במשנה ובתלמוד. רק יעסקו בתורתו החדשה המלאה אלילי תמונות ופרצופים בעלי מפלצות וזווגים שונים. ובתחבולה ומרמה קברו בארון בקרקע עולם להתיישן ויתחיל להרקיב, אחר ימים פרסמו והוציאו ע״י מלך ממלכי העכו״מ (כמ״ש הר' חיד״א בספרו שם הגדולים במערכת ספרים) כאלו הוא חבור קדום מרשב״י התנא. והרבה מרבותינו האחרונים נמשכו אחר הפילסוף הנז' בעצימות עין כי לקח אותם בחלק שפתיו וגיזומיו, ולא שתו לב לכזביו! חלילה לשום בר ישראל להאמין כי רשב״י התנא או זולתו מרז״ל האמינו באמונות כאלה להחליף ולהמיר את ה' אלוהינו אשר הודיע דרכיו למשה לבני ישראל עלילותיו, כי חנון ורחום הוא ארך אפים ורב חסד וכו' באל זר קצר אפים (ז״א) ולצרף ולשתף עמו חמשה פרצופים אשר לא נתברר אפי' מציאותם במופת (ר' עובדיה ספורנו דברים י״ג ג'. ושם כ״ט כ״ה וע״ש כ״ח י״ד) ולקרותם בשם ה' אלהינו, ואת ה' אלהים אמת שנתבררה מציאותו בכמה מופתים כראי מוצק חזקים כמו שכ' הרס״ג בס' האמונות והדעות והר' בעל חובת הלבבות והרמב״ם במורה נבוכים וביד, נטוש ונעזוב ונאמר שאין לו שם ולא כנוי ונעבוד את הפרצופים והצורות אשר לפי דבריו נבראו ונשתלשלו ממנו! וכבר הארכנו בפנים בראיות נכוחות מרז״ל בביטול הדעות הזרות האלו של הפילסוף נביא השקר הנז'! כי הוא מטרת תורתינו הקדושה להרחיקנו מאמונת האלילים בין גשמיים בין רוחניים. ולדעת כי ה' הוא האלהים אין עוד מלבדו. ואין לו לא אב ולא אח ולא בן (רבה פרשת יתרו) ולא קיבל מלכותו מאחר ואינו עתיד למסרה לאחר כמו שביארו רבותינו ז״ל (רבה במדבר שיר השירים פרשה א' ירושלמי סנהדרין פרק דיני ממונות) על פסוק אני ראשון ואני אחרון ומבלעדי אין אלהים ועל מה שאמר נבוכדנצר וריויה די רביעאה דמי לבר אלהין! וראיתי להר' מנחם די לונזבנו בס' דרך החיים שהרחיב פיו לדבר תועה על החסיד הקדוש בעל חובת הלבבות על אמרו בשם הפילסוף שאמר שלא יוכל לעבוד עלת העלות רק נביא הדור בטבעו, או הפילסוף המובהק במה שקנהו מן החכמה אבל זולתם עובדים זולתו מפני שאינם מבינים נמצא אלא מורכב! ועל דברי החסיד האלו התקצף וחשב כי כאן מקום תלונה עליו ותואנה מצא להתעות השומעים ולהבהילם באמרו וכי פלוני ופלוני מחכמינו האחרונים המפורסמים עובדים זולתו? והוא ספרם בלשונו אחד אחד בשמותם. ולא שם אל לבו כי הרבה מהחכמים שַמָנָה הלכו בעצימות עינים והאמינו לדברי מחבר הזהר אשר כינה את עצמו בשם רשב״י התנא כדי לתלות נבואותיו השקריות בשמו של התנא, ודבר סרה על ה' אלהינו שנגלה לו בבית מדרשו עתיקא קדישא וגילה לו סודות נעלמים וכוחות נעלמות (זהר בראשית כ״ב) כי הבורא הוא האליל הנקרא אבא שהיה מצוה לזוגתו האלילה אימא והיא עושה כל אשר יאמר אליה חוץ מבריאת אדם הראשון שהיא אימא שאמרה ועשתה את אדם הראשון שלא ברצון אבא! ובפרשת וירא דף קי״ב וכן בפרשת בלק דף קצ״א מיחס עבודתינו ותפילותינו לאל זר קצר אפים (זעיר אנפין) אשר לא עשה את השמים ואת הארץ. וקורא הוא לעצבים הללו בשם ה' אלהינו. ועל זעיר אנפין קרא ירמיה הנביא ע״ה כדנה תימרון להום אלהיא די שמיא וארקא לא עבדו ייבדו מן ארעא ומן תחות שמיא אלה. שהרי מבואר בזהר בכמה מקומות שהוא נברא ולא בורא (עיין זהר בראשית כ״ב וביושר לבב דף ג' ע״ב) ועכ״ז הוא אומר לעבדו ולברכו ולקבלו עליו באלוה ולירא מזוגתו! והרבה מרבותינו האחרונים שהזכיר בספר דרך החיים בתמיהתו על החסיד בעל חובת הלבבות הלכו בעצימת עין ונמשכו אחריו והאמינו באמונתו וקיבלו האליל קצר אפים (ז״א) עליהם באלוה! וכל זה גרם להם על שלא התבוננו בחכמת יוצר הכל שהוא לחקור ולבחון במעשה בראשית המבוארים בדברי הר' מורה נבוכים ובדברי החסיד בעל חובת הלבבות. ורבינו סעדיה גאון ור״י הלוי בכוזרי כמו שאמר הנביא ע״ה ואת פעל ה' לא יביטו. כי חשבו דבריהם לחכמה יונית ורחקו ממנה הרחק כמטחוי קשת, וגם לא שמו אל לבם דברי אבותינו ז״ל האמתיים במדרשים שלא הכיר אברהם אבינו את בוראו אלא בהתבוננו בפעל ה' וחקר ודרש בהם והביא מופתים מהם עד שהשיג האמת ואז נגלה לו הקב״ה ואמר לו אהבת צדק ותשנא רשע וכו' משל לאחד שהיה מהלך בדרך וראה בירה וכו' (מדרש הגדול ריש פרשת לך לך) ועל כרחנו נודה על האמת שדברי החסיד בעל חובת הלבבות ברור מללו שלא יוכל לעבוד את אלהים אמת אלא וכו' אבל זולתם עובדים זולתו כמאמר הנביא וימים רבים לישראל ללא אלהי אמת, וכמו שעינינו רואות ואזנינו שומעות עד שאפשר שיאמר אדם לחבירו! עַמְּךָ עַמִּי וְאלֹהֶיך לֹא אֱלהָי! ואל אלהי אמת נשפוך שיחנו לנחותינו במעגלי יושר ולהצילנו מכל אמונה זרה אכי״ר.

| id | word | punctuation after |
|---|---|---|
| 1:4.1 | ולתקופת |  |
| 1:4.2 | השנה |  |
| 1:4.3 | השלישית |  |
| 1:4.4 | פרצה |  |
| 1:4.5 | המלחמה |  |
| 1:4.6 | נגד |  |
| 1:4.7 | הלמודים |  |
| 1:4.8 | הטובים |  |
| 1:4.9 | והסדרים |  |
| 1:4.10 | הנעימים |  |
| 1:4.11 | כמצווה |  |
| 1:4.12 | עלינו |  |
| 1:4.13 | מרז״ל | . |
| 1:4.14 | ויוציאו |  |
| 1:4.15 | המתנגדים |  |
| 1:4.16 | דבה |  |
| 1:4.17 | רעה |  |
| 1:4.18 | לאמר |  |
| 1:4.19 | כי |  |
| 1:4.20 | על |  |
| 1:4.21 | ידי |  |
| 1:4.22 | זה |  |
| 1:4.23 | יצאו |  |
| 1:4.24 | הילדים |  |
| 1:4.25 | ח״ו |  |
| 1:4.26 | לתרבות |  |
| 1:4.27 | רעה | , |
| 1:4.28 | והרבה |  |
| 1:4.29 | מהטפשים |  |
| 1:4.30 | האמינו |  |
| 1:4.31 | לשקריהם |  |
| 1:4.32 | והוציאו |  |
| 1:4.33 | את |  |
| 1:4.34 | ילדיהם |  |
| 1:4.35 | מבית |  |
| 1:4.36 | הספר | . |
| 1:4.37 | וגם |  |
| 1:4.38 | הבנים |  |
| 1:4.39 | הסכלים |  |
| 1:4.40 | אשר |  |
| 1:4.41 | יצאו |  |
| 1:4.42 | מבית |  |
| 1:4.43 | הספר |  |
| 1:4.44 | עוינים |  |
| 1:4.45 | את |  |
| 1:4.46 | מוריהם |  |
| 1:4.47 | כי |  |
| 1:4.48 | לעניות |  |
| 1:4.49 | דעתם |  |
| 1:4.50 | הם |  |
| 1:4.51 | מינים |  |
| 1:4.52 | ויוצאים |  |
| 1:4.53 | מן |  |
| 1:4.54 | הכלל |  |
| 1:4.55 | אך |  |
| 1:4.56 | המשכילים |  |
| 1:4.57 | שבהם |  |
| 1:4.58 | יודעים |  |
| 1:4.59 | ומכירים |  |
| 1:4.60 | כי |  |
| 1:4.61 | שקר |  |
| 1:4.62 | בימין |  |
| 1:4.63 | מוציאי |  |
| 1:4.64 | הדבה |  |
| 1:4.65 | הרעה |  |
| 1:4.66 | הלזו | , |
| 1:4.67 | ודאגו |  |
| 1:4.68 | מאד |  |
| 1:4.69 | על |  |
| 1:4.70 | שהוציאום |  |
| 1:4.71 | אבותם |  |
| 1:4.72 | מבית |  |
| 1:4.73 | הספר |  |
| 1:4.74 | בסכלותם | ! |
| 1:4.75 | וכאשר |  |
| 1:4.76 | התעו |  |
| 1:4.77 | רבים |  |
| 1:4.78 | חזקה |  |
| 1:4.79 | יד |  |
| 1:4.80 | החשכים |  |
| 1:4.81 | ובמלשינות |  |
| 1:4.82 | שקר |  |
| 1:4.83 | הלשינו |  |
| 1:4.84 | על |  |
| 1:4.85 | המלמדים |  |
| 1:4.86 | אצל |  |
| 1:4.87 | הממשלה |  |
| 1:4.88 | שהם |  |
| 1:4.89 | חפצים |  |
| 1:4.90 | להמריד |  |
| 1:4.91 | את |  |
| 1:4.92 | יושבי |  |
| 1:4.93 | הארץ | . |
| 1:4.94 | ושהם |  |
| 1:4.95 | כותבי |  |
| 1:4.96 | פלסתר |  |
| 1:4.97 | על |  |
| 1:4.98 | הממשלה |  |
| 1:4.99 | אל |  |
| 1:4.100 | ממשלת |  |
| 1:4.101 | אנגליא |  |
| 1:4.102 | לחסות |  |
| 1:4.103 | בצל |  |
| 1:4.104 | ממשלתה |  |
| 1:4.105 | וגם |  |
| 1:4.106 | אל |  |
| 1:4.107 | הינ״ץ |  |
| 1:4.108 | ועבור |  |
| 1:4.109 | זה |  |
| 1:4.110 | המלשינות |  |
| 1:4.111 | הובלו |  |
| 1:4.112 | המלמדים |  |
| 1:4.113 | אל |  |
| 1:4.114 | בית |  |
| 1:4.115 | הכלא |  |
| 1:4.116 | וגם |  |
| 1:4.117 | גורשו |  |
| 1:4.118 | מבית |  |
| 1:4.119 | הספר | . |
| 1:4.120 | והביאו |  |
| 1:4.121 | תחתיהם |  |
| 1:4.122 | מלמד |  |
| 1:4.123 | אחר |  |
| 1:4.124 | אשר |  |
| 1:4.125 | לא |  |
| 1:4.126 | ידע |  |
| 1:4.127 | לא |  |
| 1:4.128 | עברית |  |
| 1:4.129 | על |  |
| 1:4.130 | בריה |  |
| 1:4.131 | ולא |  |
| 1:4.132 | ערבית |  |
| 1:4.133 | כתקנה | . |
| 1:4.134 | ואין |  |
| 1:4.135 | עליו |  |
| 1:4.136 | ללמדם |  |
| 1:4.137 | רק |  |
| 1:4.138 | פרשת |  |
| 1:4.139 | השבוע |  |
| 1:4.140 | עם |  |
| 1:4.141 | תרגום |  |
| 1:4.142 | אונקלוס |  |
| 1:4.143 | ותרגום |  |
| 1:4.144 | רס״ג |  |
| 1:4.145 | בערבי | . |
| 1:4.146 | והפטרה |  |
| 1:4.147 | עם |  |
| 1:4.148 | תרגום |  |
| 1:4.149 | יונתן |  |
| 1:4.150 | ב״ע |  |
| 1:4.151 | באין |  |
| 1:4.152 | הבין |  |
| 1:4.153 | כלום |  |
| 1:4.154 | כשאר |  |
| 1:4.155 | המלמדים | ! |
| 1:4.156 | ולימים |  |
| 1:4.157 | מספר |  |
| 1:4.158 | בטל |  |
| 1:4.159 | בית |  |
| 1:4.160 | הספר |  |
| 1:4.161 | לגמרי | ! |
| 1:4.162 | כל |  |
| 1:4.163 | זאת |  |
| 1:4.164 | באתנו |  |
| 1:4.165 | על |  |
| 1:4.166 | אשר |  |
| 1:4.167 | לא |  |
| 1:4.168 | מחינו |  |
| 1:4.169 | בכבוד |  |
| 1:4.170 | קונינו |  |
| 1:4.171 | בעת |  |
| 1:4.172 | שהיינו |  |
| 1:4.173 | שומעים |  |
| 1:4.174 | מהלומדים |  |
| 1:4.175 | בזהר |  |
| 1:4.176 | ושאר |  |
| 1:4.177 | מקובלים |  |
| 1:4.178 | דברים |  |
| 1:4.179 | שאסור |  |
| 1:4.180 | לשמען |  |
| 1:4.181 | ומחרישים |  |
| 1:4.182 | להם |  |
| 1:4.183 | כמו |  |
| 1:4.184 | שאמרו |  |
| 1:4.185 | רז״ל |  |
| 1:4.186 | דר״א |  |
| 1:4.187 | נתפש |  |
| 1:4.188 | למינות |  |
| 1:4.189 | וכו | ' |
| 1:4.190 | שמא |  |
| 1:4.191 | דבר |  |
| 1:4.192 | מינות |  |
| 1:4.193 | וכו | ' |
| 1:4.194 | ונתקיים |  |
| 1:4.195 | בנו |  |
| 1:4.196 | מאמר |  |
| 1:4.197 | רז״ל |  |
| 1:4.198 | דמודי |  |
| 1:4.199 | להון |  |
| 1:4.200 | נפיל |  |
| 1:4.201 | בידיהון |  |
| 1:4.202 | דמתרחץ |  |
| 1:4.203 | עליהון |  |
| 1:4.204 | דיליה |  |
| 1:4.205 | דילהון |  |
| 1:4.206 | ע״כ | ! |
| 1:4.207 | לכן |  |
| 1:4.208 | עזבנו |  |
| 1:4.209 | דרכינו |  |
| 1:4.210 | בלימוד |  |
| 1:4.211 | השפות |  |
| 1:4.212 | העברית |  |
| 1:4.213 | והערבית |  |
| 1:4.214 | בדקדוקן |  |
| 1:4.215 | על |  |
| 1:4.216 | כורחנו |  |
| 1:4.217 | מפני |  |
| 1:4.218 | חמת |  |
| 1:4.219 | המציק |  |
| 1:4.220 | על |  |
| 1:4.221 | ידי |  |
| 1:4.222 | מלשינות |  |
| 1:4.223 | שקר |  |
| 1:4.224 | ונפן |  |
| 1:4.225 | אל |  |
| 1:4.226 | דרך |  |
| 1:4.227 | הקדש |  |
| 1:4.228 | כבוד |  |
| 1:4.229 | אבינו |  |
| 1:4.230 | שבשמים |  |
| 1:4.231 | בבטול |  |
| 1:4.232 | דעות |  |
| 1:4.233 | הפילסוף |  |
| 1:4.234 | נביא |  |
| 1:4.235 | השקר |  |
| 1:4.236 | מחבר |  |
| 1:4.237 | הזהר |  |
| 1:4.238 | אשר |  |
| 1:4.239 | אסף |  |
| 1:4.240 | בחפניו |  |
| 1:4.241 | דעות |  |
| 1:4.242 | הפלסופיא |  |
| 1:4.243 | הקדומה |  |
| 1:4.244 | עובדי |  |
| 1:4.245 | האלילים |  |
| 1:4.246 | תמונות |  |
| 1:4.247 | וכוחות |  |
| 1:4.248 | נעלמים |  |
| 1:4.249 | אשר |  |
| 1:4.250 | בדו |  |
| 1:4.251 | מלבם | . |
| 1:4.252 | ומחבר |  |
| 1:4.253 | הזהר |  |
| 1:4.254 | תפס |  |
| 1:4.255 | דעותיהם |  |
| 1:4.256 | וחשב |  |
| 1:4.257 | להכניסם |  |
| 1:4.258 | באמונת |  |
| 1:4.259 | תורתינו |  |
| 1:4.260 | הקדושה | , |
| 1:4.261 | וקרא |  |
| 1:4.262 | את |  |
| 1:4.263 | שמו |  |
| 1:4.264 | בשם |  |
| 1:4.265 | רשב״י |  |
| 1:4.266 | התנא |  |
| 1:4.267 | בשקר |  |
| 1:4.268 | ואמר |  |
| 1:4.269 | שנגלה |  |
| 1:4.270 | לו |  |
| 1:4.271 | הקב״ה |  |
| 1:4.272 | הנקרא |  |
| 1:4.273 | אצלו |  |
| 1:4.274 | עתיקא |  |
| 1:4.275 | קדישא |  |
| 1:4.276 | וגילה |  |
| 1:4.277 | לו |  |
| 1:4.278 | רזין |  |
| 1:4.279 | טמירין | . |
| 1:4.280 | ותלה |  |
| 1:4.281 | נבואתו |  |
| 1:4.282 | ורעיונותיו |  |
| 1:4.283 | ברשב״י |  |
| 1:4.284 | התנא |  |
| 1:4.285 | ע״ה | . |
| 1:4.286 | ועשה |  |
| 1:4.287 | את |  |
| 1:4.288 | חבורו |  |
| 1:4.289 | הזהר |  |
| 1:4.290 | ואמר |  |
| 1:4.291 | שהעליונים |  |
| 1:4.292 | הסכימו |  |
| 1:4.293 | לחבורו |  |
| 1:4.294 | לפי |  |
| 1:4.295 | שראה |  |
| 1:4.296 | הפילסוף |  |
| 1:4.297 | הזה |  |
| 1:4.298 | כי |  |
| 1:4.299 | המשנה |  |
| 1:4.300 | ושני |  |
| 1:4.301 | התלמודים |  |
| 1:4.302 | ומדרשי |  |
| 1:4.303 | רז״ל |  |
| 1:4.304 | האמתיים |  |
| 1:4.305 | עומדים |  |
| 1:4.306 | נגדו |  |
| 1:4.307 | באמונת |  |
| 1:4.308 | אחדותו |  |
| 1:4.309 | ית | ' |
| 1:4.310 | לכן |  |
| 1:4.311 | חשב |  |
| 1:4.312 | לפסלן |  |
| 1:4.313 | ובזה |  |
| 1:4.314 | ולעג |  |
| 1:4.315 | להם |  |
| 1:4.316 | ולכל |  |
| 1:4.317 | העוסקים |  |
| 1:4.318 | בהם | , |
| 1:4.319 | וכינה |  |
| 1:4.320 | אותם |  |
| 1:4.321 | בשמות |  |
| 1:4.322 | של |  |
| 1:4.323 | זלזול |  |
| 1:4.324 | ופחיתות | , |
| 1:4.325 | כאשר |  |
| 1:4.326 | יראה |  |
| 1:4.327 | המעיין |  |
| 1:4.328 | בפנים | ! |
| 1:4.329 | והבטיח |  |
| 1:4.330 | שלא |  |
| 1:4.331 | יזכו |  |
| 1:4.332 | ישראל |  |
| 1:4.333 | לגאולה |  |
| 1:4.334 | עד |  |
| 1:4.335 | אשר |  |
| 1:4.336 | ישובו |  |
| 1:4.337 | מדרכם |  |
| 1:4.338 | מלעסוק |  |
| 1:4.339 | במשנה |  |
| 1:4.340 | ובתלמוד | . |
| 1:4.341 | רק |  |
| 1:4.342 | יעסקו |  |
| 1:4.343 | בתורתו |  |
| 1:4.344 | החדשה |  |
| 1:4.345 | המלאה |  |
| 1:4.346 | אלילי |  |
| 1:4.347 | תמונות |  |
| 1:4.348 | ופרצופים |  |
| 1:4.349 | בעלי |  |
| 1:4.350 | מפלצות |  |
| 1:4.351 | וזווגים |  |
| 1:4.352 | שונים | . |
| 1:4.353 | ובתחבולה |  |
| 1:4.354 | ומרמה |  |
| 1:4.355 | קברו |  |
| 1:4.356 | בארון |  |
| 1:4.357 | בקרקע |  |
| 1:4.358 | עולם |  |
| 1:4.359 | להתיישן |  |
| 1:4.360 | ויתחיל |  |
| 1:4.361 | להרקיב | , |
| 1:4.362 | אחר |  |
| 1:4.363 | ימים |  |
| 1:4.364 | פרסמו |  |
| 1:4.365 | והוציאו |  |
| 1:4.366 | ע״י |  |
| 1:4.367 | מלך |  |
| 1:4.368 | ממלכי |  |
| 1:4.369 | העכו״מ |  |
| 1:4.370 | כמ״ש |  |
| 1:4.371 | הר | ' |
| 1:4.372 | חיד״א |  |
| 1:4.373 | בספרו |  |
| 1:4.374 | שם |  |
| 1:4.375 | הגדולים |  |
| 1:4.376 | במערכת |  |
| 1:4.377 | ספרים | ) |
| 1:4.378 | כאלו |  |
| 1:4.379 | הוא |  |
| 1:4.380 | חבור |  |
| 1:4.381 | קדום |  |
| 1:4.382 | מרשב״י |  |
| 1:4.383 | התנא | . |
| 1:4.384 | והרבה |  |
| 1:4.385 | מרבותינו |  |
| 1:4.386 | האחרונים |  |
| 1:4.387 | נמשכו |  |
| 1:4.388 | אחר |  |
| 1:4.389 | הפילסוף |  |
| 1:4.390 | הנז | ' |
| 1:4.391 | בעצימות |  |
| 1:4.392 | עין |  |
| 1:4.393 | כי |  |
| 1:4.394 | לקח |  |
| 1:4.395 | אותם |  |
| 1:4.396 | בחלק |  |
| 1:4.397 | שפתיו |  |
| 1:4.398 | וגיזומיו | , |
| 1:4.399 | ולא |  |
| 1:4.400 | שתו |  |
| 1:4.401 | לב |  |
| 1:4.402 | לכזביו | ! |
| 1:4.403 | חלילה |  |
| 1:4.404 | לשום |  |
| 1:4.405 | בר |  |
| 1:4.406 | ישראל |  |
| 1:4.407 | להאמין |  |
| 1:4.408 | כי |  |
| 1:4.409 | רשב״י |  |
| 1:4.410 | התנא |  |
| 1:4.411 | או |  |
| 1:4.412 | זולתו |  |
| 1:4.413 | מרז״ל |  |
| 1:4.414 | האמינו |  |
| 1:4.415 | באמונות |  |
| 1:4.416 | כאלה |  |
| 1:4.417 | להחליף |  |
| 1:4.418 | ולהמיר |  |
| 1:4.419 | את |  |
| 1:4.420 | ה | ' |
| 1:4.421 | אלוהינו |  |
| 1:4.422 | אשר |  |
| 1:4.423 | הודיע |  |
| 1:4.424 | דרכיו |  |
| 1:4.425 | למשה |  |
| 1:4.426 | לבני |  |
| 1:4.427 | ישראל |  |
| 1:4.428 | עלילותיו | , |
| 1:4.429 | כי |  |
| 1:4.430 | חנון |  |
| 1:4.431 | ורחום |  |
| 1:4.432 | הוא |  |
| 1:4.433 | ארך |  |
| 1:4.434 | אפים |  |
| 1:4.435 | ורב |  |
| 1:4.436 | חסד |  |
| 1:4.437 | וכו | ' |
| 1:4.438 | באל |  |
| 1:4.439 | זר |  |
| 1:4.440 | קצר |  |
| 1:4.441 | אפים |  |
| 1:4.442 | ז״א | ) |
| 1:4.443 | ולצרף |  |
| 1:4.444 | ולשתף |  |
| 1:4.445 | עמו |  |
| 1:4.446 | חמשה |  |
| 1:4.447 | פרצופים |  |
| 1:4.448 | אשר |  |
| 1:4.449 | לא |  |
| 1:4.450 | נתברר |  |
| 1:4.451 | אפי | ' |
| 1:4.452 | מציאותם |  |
| 1:4.453 | במופת |  |
| 1:4.454 | ר | ' |
| 1:4.455 | עובדיה |  |
| 1:4.456 | ספורנו |  |
| 1:4.457 | דברים |  |
| 1:4.458 | י״ג |  |
| 1:4.459 | ג | '. |
| 1:4.460 | ושם |  |
| 1:4.461 | כ״ט |  |
| 1:4.462 | כ״ה |  |
| 1:4.463 | וע״ש |  |
| 1:4.464 | כ״ח |  |
| 1:4.465 | י״ד | ) |
| 1:4.466 | ולקרותם |  |
| 1:4.467 | בשם |  |
| 1:4.468 | ה | ' |
| 1:4.469 | אלהינו | , |
| 1:4.470 | ואת |  |
| 1:4.471 | ה | ' |
| 1:4.472 | אלהים |  |
| 1:4.473 | אמת |  |
| 1:4.474 | שנתבררה |  |
| 1:4.475 | מציאותו |  |
| 1:4.476 | בכמה |  |
| 1:4.477 | מופתים |  |
| 1:4.478 | כראי |  |
| 1:4.479 | מוצק |  |
| 1:4.480 | חזקים |  |
| 1:4.481 | כמו |  |
| 1:4.482 | שכ | ' |
| 1:4.483 | הרס״ג |  |
| 1:4.484 | בס | ' |
| 1:4.485 | האמונות |  |
| 1:4.486 | והדעות |  |
| 1:4.487 | והר | ' |
| 1:4.488 | בעל |  |
| 1:4.489 | חובת |  |
| 1:4.490 | הלבבות |  |
| 1:4.491 | והרמב״ם |  |
| 1:4.492 | במורה |  |
| 1:4.493 | נבוכים |  |
| 1:4.494 | וביד | , |
| 1:4.495 | נטוש |  |
| 1:4.496 | ונעזוב |  |
| 1:4.497 | ונאמר |  |
| 1:4.498 | שאין |  |
| 1:4.499 | לו |  |
| 1:4.500 | שם |  |
| 1:4.501 | ולא |  |
| 1:4.502 | כנוי |  |
| 1:4.503 | ונעבוד |  |
| 1:4.504 | את |  |
| 1:4.505 | הפרצופים |  |
| 1:4.506 | והצורות |  |
| 1:4.507 | אשר |  |
| 1:4.508 | לפי |  |
| 1:4.509 | דבריו |  |
| 1:4.510 | נבראו |  |
| 1:4.511 | ונשתלשלו |  |
| 1:4.512 | ממנו | ! |
| 1:4.513 | וכבר |  |
| 1:4.514 | הארכנו |  |
| 1:4.515 | בפנים |  |
| 1:4.516 | בראיות |  |
| 1:4.517 | נכוחות |  |
| 1:4.518 | מרז״ל |  |
| 1:4.519 | בביטול |  |
| 1:4.520 | הדעות |  |
| 1:4.521 | הזרות |  |
| 1:4.522 | האלו |  |
| 1:4.523 | של |  |
| 1:4.524 | הפילסוף |  |
| 1:4.525 | נביא |  |
| 1:4.526 | השקר |  |
| 1:4.527 | הנז | '! |
| 1:4.528 | כי |  |
| 1:4.529 | הוא |  |
| 1:4.530 | מטרת |  |
| 1:4.531 | תורתינו |  |
| 1:4.532 | הקדושה |  |
| 1:4.533 | להרחיקנו |  |
| 1:4.534 | מאמונת |  |
| 1:4.535 | האלילים |  |
| 1:4.536 | בין |  |
| 1:4.537 | גשמיים |  |
| 1:4.538 | בין |  |
| 1:4.539 | רוחניים | . |
| 1:4.540 | ולדעת |  |
| 1:4.541 | כי |  |
| 1:4.542 | ה | ' |
| 1:4.543 | הוא |  |
| 1:4.544 | האלהים |  |
| 1:4.545 | אין |  |
| 1:4.546 | עוד |  |
| 1:4.547 | מלבדו | . |
| 1:4.548 | ואין |  |
| 1:4.549 | לו |  |
| 1:4.550 | לא |  |
| 1:4.551 | אב |  |
| 1:4.552 | ולא |  |
| 1:4.553 | אח |  |
| 1:4.554 | ולא |  |
| 1:4.555 | בן |  |
| 1:4.556 | רבה |  |
| 1:4.557 | פרשת |  |
| 1:4.558 | יתרו | ) |
| 1:4.559 | ולא |  |
| 1:4.560 | קיבל |  |
| 1:4.561 | מלכותו |  |
| 1:4.562 | מאחר |  |
| 1:4.563 | ואינו |  |
| 1:4.564 | עתיד |  |
| 1:4.565 | למסרה |  |
| 1:4.566 | לאחר |  |
| 1:4.567 | כמו |  |
| 1:4.568 | שביארו |  |
| 1:4.569 | רבותינו |  |
| 1:4.570 | ז״ל |  |
| 1:4.571 | רבה |  |
| 1:4.572 | במדבר |  |
| 1:4.573 | שיר |  |
| 1:4.574 | השירים |  |
| 1:4.575 | פרשה |  |
| 1:4.576 | א | ' |
| 1:4.577 | ירושלמי |  |
| 1:4.578 | סנהדרין |  |
| 1:4.579 | פרק |  |
| 1:4.580 | דיני |  |
| 1:4.581 | ממונות | ) |
| 1:4.582 | על |  |
| 1:4.583 | פסוק |  |
| 1:4.584 | אני |  |
| 1:4.585 | ראשון |  |
| 1:4.586 | ואני |  |
| 1:4.587 | אחרון |  |
| 1:4.588 | ומבלעדי |  |
| 1:4.589 | אין |  |
| 1:4.590 | אלהים |  |
| 1:4.591 | ועל |  |
| 1:4.592 | מה |  |
| 1:4.593 | שאמר |  |
| 1:4.594 | נבוכדנצר |  |
| 1:4.595 | וריויה |  |
| 1:4.596 | די |  |
| 1:4.597 | רביעאה |  |
| 1:4.598 | דמי |  |
| 1:4.599 | לבר |  |
| 1:4.600 | אלהין | ! |
| 1:4.601 | וראיתי |  |
| 1:4.602 | להר | ' |
| 1:4.603 | מנחם |  |
| 1:4.604 | די |  |
| 1:4.605 | לונזבנו |  |
| 1:4.606 | בס | ' |
| 1:4.607 | דרך |  |
| 1:4.608 | החיים |  |
| 1:4.609 | שהרחיב |  |
| 1:4.610 | פיו |  |
| 1:4.611 | לדבר |  |
| 1:4.612 | תועה |  |
| 1:4.613 | על |  |
| 1:4.614 | החסיד |  |
| 1:4.615 | הקדוש |  |
| 1:4.616 | בעל |  |
| 1:4.617 | חובת |  |
| 1:4.618 | הלבבות |  |
| 1:4.619 | על |  |
| 1:4.620 | אמרו |  |
| 1:4.621 | בשם |  |
| 1:4.622 | הפילסוף |  |
| 1:4.623 | שאמר |  |
| 1:4.624 | שלא |  |
| 1:4.625 | יוכל |  |
| 1:4.626 | לעבוד |  |
| 1:4.627 | עלת |  |
| 1:4.628 | העלות |  |
| 1:4.629 | רק |  |
| 1:4.630 | נביא |  |
| 1:4.631 | הדור |  |
| 1:4.632 | בטבעו | , |
| 1:4.633 | או |  |
| 1:4.634 | הפילסוף |  |
| 1:4.635 | המובהק |  |
| 1:4.636 | במה |  |
| 1:4.637 | שקנהו |  |
| 1:4.638 | מן |  |
| 1:4.639 | החכמה |  |
| 1:4.640 | אבל |  |
| 1:4.641 | זולתם |  |
| 1:4.642 | עובדים |  |
| 1:4.643 | זולתו |  |
| 1:4.644 | מפני |  |
| 1:4.645 | שאינם |  |
| 1:4.646 | מבינים |  |
| 1:4.647 | נמצא |  |
| 1:4.648 | אלא |  |
| 1:4.649 | מורכב | ! |
| 1:4.650 | ועל |  |
| 1:4.651 | דברי |  |
| 1:4.652 | החסיד |  |
| 1:4.653 | האלו |  |
| 1:4.654 | התקצף |  |
| 1:4.655 | וחשב |  |
| 1:4.656 | כי |  |
| 1:4.657 | כאן |  |
| 1:4.658 | מקום |  |
| 1:4.659 | תלונה |  |
| 1:4.660 | עליו |  |
| 1:4.661 | ותואנה |  |
| 1:4.662 | מצא |  |
| 1:4.663 | להתעות |  |
| 1:4.664 | השומעים |  |
| 1:4.665 | ולהבהילם |  |
| 1:4.666 | באמרו |  |
| 1:4.667 | וכי |  |
| 1:4.668 | פלוני |  |
| 1:4.669 | ופלוני |  |
| 1:4.670 | מחכמינו |  |
| 1:4.671 | האחרונים |  |
| 1:4.672 | המפורסמים |  |
| 1:4.673 | עובדים |  |
| 1:4.674 | זולתו | ? |
| 1:4.675 | והוא |  |
| 1:4.676 | ספרם |  |
| 1:4.677 | בלשונו |  |
| 1:4.678 | אחד |  |
| 1:4.679 | אחד |  |
| 1:4.680 | בשמותם | . |
| 1:4.681 | ולא |  |
| 1:4.682 | שם |  |
| 1:4.683 | אל |  |
| 1:4.684 | לבו |  |
| 1:4.685 | כי |  |
| 1:4.686 | הרבה |  |
| 1:4.687 | מהחכמים |  |
| 1:4.688 | שַמָנָה |  |
| 1:4.689 | הלכו |  |
| 1:4.690 | בעצימות |  |
| 1:4.691 | עינים |  |
| 1:4.692 | והאמינו |  |
| 1:4.693 | לדברי |  |
| 1:4.694 | מחבר |  |
| 1:4.695 | הזהר |  |
| 1:4.696 | אשר |  |
| 1:4.697 | כינה |  |
| 1:4.698 | את |  |
| 1:4.699 | עצמו |  |
| 1:4.700 | בשם |  |
| 1:4.701 | רשב״י |  |
| 1:4.702 | התנא |  |
| 1:4.703 | כדי |  |
| 1:4.704 | לתלות |  |
| 1:4.705 | נבואותיו |  |
| 1:4.706 | השקריות |  |
| 1:4.707 | בשמו |  |
| 1:4.708 | של |  |
| 1:4.709 | התנא | , |
| 1:4.710 | ודבר |  |
| 1:4.711 | סרה |  |
| 1:4.712 | על |  |
| 1:4.713 | ה | ' |
| 1:4.714 | אלהינו |  |
| 1:4.715 | שנגלה |  |
| 1:4.716 | לו |  |
| 1:4.717 | בבית |  |
| 1:4.718 | מדרשו |  |
| 1:4.719 | עתיקא |  |
| 1:4.720 | קדישא |  |
| 1:4.721 | וגילה |  |
| 1:4.722 | לו |  |
| 1:4.723 | סודות |  |
| 1:4.724 | נעלמים |  |
| 1:4.725 | וכוחות |  |
| 1:4.726 | נעלמות |  |
| 1:4.727 | זהר |  |
| 1:4.728 | בראשית |  |
| 1:4.729 | כ״ב | ) |
| 1:4.730 | כי |  |
| 1:4.731 | הבורא |  |
| 1:4.732 | הוא |  |
| 1:4.733 | האליל |  |
| 1:4.734 | הנקרא |  |
| 1:4.735 | אבא |  |
| 1:4.736 | שהיה |  |
| 1:4.737 | מצוה |  |
| 1:4.738 | לזוגתו |  |
| 1:4.739 | האלילה |  |
| 1:4.740 | אימא |  |
| 1:4.741 | והיא |  |
| 1:4.742 | עושה |  |
| 1:4.743 | כל |  |
| 1:4.744 | אשר |  |
| 1:4.745 | יאמר |  |
| 1:4.746 | אליה |  |
| 1:4.747 | חוץ |  |
| 1:4.748 | מבריאת |  |
| 1:4.749 | אדם |  |
| 1:4.750 | הראשון |  |
| 1:4.751 | שהיא |  |
| 1:4.752 | אימא |  |
| 1:4.753 | שאמרה |  |
| 1:4.754 | ועשתה |  |
| 1:4.755 | את |  |
| 1:4.756 | אדם |  |
| 1:4.757 | הראשון |  |
| 1:4.758 | שלא |  |
| 1:4.759 | ברצון |  |
| 1:4.760 | אבא | ! |
| 1:4.761 | ובפרשת |  |
| 1:4.762 | וירא |  |
| 1:4.763 | דף |  |
| 1:4.764 | קי״ב |  |
| 1:4.765 | וכן |  |
| 1:4.766 | בפרשת |  |
| 1:4.767 | בלק |  |
| 1:4.768 | דף |  |
| 1:4.769 | קצ״א |  |
| 1:4.770 | מיחס |  |
| 1:4.771 | עבודתינו |  |
| 1:4.772 | ותפילותינו |  |
| 1:4.773 | לאל |  |
| 1:4.774 | זר |  |
| 1:4.775 | קצר |  |
| 1:4.776 | אפים |  |
| 1:4.777 | זעיר |  |
| 1:4.778 | אנפין | ) |
| 1:4.779 | אשר |  |
| 1:4.780 | לא |  |
| 1:4.781 | עשה |  |
| 1:4.782 | את |  |
| 1:4.783 | השמים |  |
| 1:4.784 | ואת |  |
| 1:4.785 | הארץ | . |
| 1:4.786 | וקורא |  |
| 1:4.787 | הוא |  |
| 1:4.788 | לעצבים |  |
| 1:4.789 | הללו |  |
| 1:4.790 | בשם |  |
| 1:4.791 | ה | ' |
| 1:4.792 | אלהינו | . |
| 1:4.793 | ועל |  |
| 1:4.794 | זעיר |  |
| 1:4.795 | אנפין |  |
| 1:4.796 | קרא |  |
| 1:4.797 | ירמיה |  |
| 1:4.798 | הנביא |  |
| 1:4.799 | ע״ה |  |
| 1:4.800 | כדנה |  |
| 1:4.801 | תימרון |  |
| 1:4.802 | להום |  |
| 1:4.803 | אלהיא |  |
| 1:4.804 | די |  |
| 1:4.805 | שמיא |  |
| 1:4.806 | וארקא |  |
| 1:4.807 | לא |  |
| 1:4.808 | עבדו |  |
| 1:4.809 | ייבדו |  |
| 1:4.810 | מן |  |
| 1:4.811 | ארעא |  |
| 1:4.812 | ומן |  |
| 1:4.813 | תחות |  |
| 1:4.814 | שמיא |  |
| 1:4.815 | אלה | . |
| 1:4.816 | שהרי |  |
| 1:4.817 | מבואר |  |
| 1:4.818 | בזהר |  |
| 1:4.819 | בכמה |  |
| 1:4.820 | מקומות |  |
| 1:4.821 | שהוא |  |
| 1:4.822 | נברא |  |
| 1:4.823 | ולא |  |
| 1:4.824 | בורא |  |
| 1:4.825 | עיין |  |
| 1:4.826 | זהר |  |
| 1:4.827 | בראשית |  |
| 1:4.828 | כ״ב |  |
| 1:4.829 | וביושר |  |
| 1:4.830 | לבב |  |
| 1:4.831 | דף |  |
| 1:4.832 | ג | ' |
| 1:4.833 | ע״ב | ) |
| 1:4.834 | ועכ״ז |  |
| 1:4.835 | הוא |  |
| 1:4.836 | אומר |  |
| 1:4.837 | לעבדו |  |
| 1:4.838 | ולברכו |  |
| 1:4.839 | ולקבלו |  |
| 1:4.840 | עליו |  |
| 1:4.841 | באלוה |  |
| 1:4.842 | ולירא |  |
| 1:4.843 | מזוגתו | ! |
| 1:4.844 | והרבה |  |
| 1:4.845 | מרבותינו |  |
| 1:4.846 | האחרונים |  |
| 1:4.847 | שהזכיר |  |
| 1:4.848 | בספר |  |
| 1:4.849 | דרך |  |
| 1:4.850 | החיים |  |
| 1:4.851 | בתמיהתו |  |
| 1:4.852 | על |  |
| 1:4.853 | החסיד |  |
| 1:4.854 | בעל |  |
| 1:4.855 | חובת |  |
| 1:4.856 | הלבבות |  |
| 1:4.857 | הלכו |  |
| 1:4.858 | בעצימת |  |
| 1:4.859 | עין |  |
| 1:4.860 | ונמשכו |  |
| 1:4.861 | אחריו |  |
| 1:4.862 | והאמינו |  |
| 1:4.863 | באמונתו |  |
| 1:4.864 | וקיבלו |  |
| 1:4.865 | האליל |  |
| 1:4.866 | קצר |  |
| 1:4.867 | אפים |  |
| 1:4.868 | ז״א | ) |
| 1:4.869 | עליהם |  |
| 1:4.870 | באלוה | ! |
| 1:4.871 | וכל |  |
| 1:4.872 | זה |  |
| 1:4.873 | גרם |  |
| 1:4.874 | להם |  |
| 1:4.875 | על |  |
| 1:4.876 | שלא |  |
| 1:4.877 | התבוננו |  |
| 1:4.878 | בחכמת |  |
| 1:4.879 | יוצר |  |
| 1:4.880 | הכל |  |
| 1:4.881 | שהוא |  |
| 1:4.882 | לחקור |  |
| 1:4.883 | ולבחון |  |
| 1:4.884 | במעשה |  |
| 1:4.885 | בראשית |  |
| 1:4.886 | המבוארים |  |
| 1:4.887 | בדברי |  |
| 1:4.888 | הר | ' |
| 1:4.889 | מורה |  |
| 1:4.890 | נבוכים |  |
| 1:4.891 | ובדברי |  |
| 1:4.892 | החסיד |  |
| 1:4.893 | בעל |  |
| 1:4.894 | חובת |  |
| 1:4.895 | הלבבות | . |
| 1:4.896 | ורבינו |  |
| 1:4.897 | סעדיה |  |
| 1:4.898 | גאון |  |
| 1:4.899 | ור״י |  |
| 1:4.900 | הלוי |  |
| 1:4.901 | בכוזרי |  |
| 1:4.902 | כמו |  |
| 1:4.903 | שאמר |  |
| 1:4.904 | הנביא |  |
| 1:4.905 | ע״ה |  |
| 1:4.906 | ואת |  |
| 1:4.907 | פעל |  |
| 1:4.908 | ה | ' |
| 1:4.909 | לא |  |
| 1:4.910 | יביטו | . |
| 1:4.911 | כי |  |
| 1:4.912 | חשבו |  |
| 1:4.913 | דבריהם |  |
| 1:4.914 | לחכמה |  |
| 1:4.915 | יונית |  |
| 1:4.916 | ורחקו |  |
| 1:4.917 | ממנה |  |
| 1:4.918 | הרחק |  |
| 1:4.919 | כמטחוי |  |
| 1:4.920 | קשת | , |
| 1:4.921 | וגם |  |
| 1:4.922 | לא |  |
| 1:4.923 | שמו |  |
| 1:4.924 | אל |  |
| 1:4.925 | לבם |  |
| 1:4.926 | דברי |  |
| 1:4.927 | אבותינו |  |
| 1:4.928 | ז״ל |  |
| 1:4.929 | האמתיים |  |
| 1:4.930 | במדרשים |  |
| 1:4.931 | שלא |  |
| 1:4.932 | הכיר |  |
| 1:4.933 | אברהם |  |
| 1:4.934 | אבינו |  |
| 1:4.935 | את |  |
| 1:4.936 | בוראו |  |
| 1:4.937 | אלא |  |
| 1:4.938 | בהתבוננו |  |
| 1:4.939 | בפעל |  |
| 1:4.940 | ה | ' |
| 1:4.941 | וחקר |  |
| 1:4.942 | ודרש |  |
| 1:4.943 | בהם |  |
| 1:4.944 | והביא |  |
| 1:4.945 | מופתים |  |
| 1:4.946 | מהם |  |
| 1:4.947 | עד |  |
| 1:4.948 | שהשיג |  |
| 1:4.949 | האמת |  |
| 1:4.950 | ואז |  |
| 1:4.951 | נגלה |  |
| 1:4.952 | לו |  |
| 1:4.953 | הקב״ה |  |
| 1:4.954 | ואמר |  |
| 1:4.955 | לו |  |
| 1:4.956 | אהבת |  |
| 1:4.957 | צדק |  |
| 1:4.958 | ותשנא |  |
| 1:4.959 | רשע |  |
| 1:4.960 | וכו | ' |
| 1:4.961 | משל |  |
| 1:4.962 | לאחד |  |
| 1:4.963 | שהיה |  |
| 1:4.964 | מהלך |  |
| 1:4.965 | בדרך |  |
| 1:4.966 | וראה |  |
| 1:4.967 | בירה |  |
| 1:4.968 | וכו | ' |
| 1:4.969 | מדרש |  |
| 1:4.970 | הגדול |  |
| 1:4.971 | ריש |  |
| 1:4.972 | פרשת |  |
| 1:4.973 | לך |  |
| 1:4.974 | לך | ) |
| 1:4.975 | ועל |  |
| 1:4.976 | כרחנו |  |
| 1:4.977 | נודה |  |
| 1:4.978 | על |  |
| 1:4.979 | האמת |  |
| 1:4.980 | שדברי |  |
| 1:4.981 | החסיד |  |
| 1:4.982 | בעל |  |
| 1:4.983 | חובת |  |
| 1:4.984 | הלבבות |  |
| 1:4.985 | ברור |  |
| 1:4.986 | מללו |  |
| 1:4.987 | שלא |  |
| 1:4.988 | יוכל |  |
| 1:4.989 | לעבוד |  |
| 1:4.990 | את |  |
| 1:4.991 | אלהים |  |
| 1:4.992 | אמת |  |
| 1:4.993 | אלא |  |
| 1:4.994 | וכו | ' |
| 1:4.995 | אבל |  |
| 1:4.996 | זולתם |  |
| 1:4.997 | עובדים |  |
| 1:4.998 | זולתו |  |
| 1:4.999 | כמאמר |  |
| 1:4.1000 | הנביא |  |
| 1:4.1001 | וימים |  |
| 1:4.1002 | רבים |  |
| 1:4.1003 | לישראל |  |
| 1:4.1004 | ללא |  |
| 1:4.1005 | אלהי |  |
| 1:4.1006 | אמת | , |
| 1:4.1007 | וכמו |  |
| 1:4.1008 | שעינינו |  |
| 1:4.1009 | רואות |  |
| 1:4.1010 | ואזנינו |  |
| 1:4.1011 | שומעות |  |
| 1:4.1012 | עד |  |
| 1:4.1013 | שאפשר |  |
| 1:4.1014 | שיאמר |  |
| 1:4.1015 | אדם |  |
| 1:4.1016 | לחבירו | ! |
| 1:4.1017 | עַמְּךָ |  |
| 1:4.1018 | עַמִּי |  |
| 1:4.1019 | וְאלֹהֶיך |  |
| 1:4.1020 | לֹא |  |
| 1:4.1021 | אֱלהָי | ! |
| 1:4.1022 | ואל |  |
| 1:4.1023 | אלהי |  |
| 1:4.1024 | אמת |  |
| 1:4.1025 | נשפוך |  |
| 1:4.1026 | שיחנו |  |
| 1:4.1027 | לנחותינו |  |
| 1:4.1028 | במעגלי |  |
| 1:4.1029 | יושר |  |
| 1:4.1030 | ולהצילנו |  |
| 1:4.1031 | מכל |  |
| 1:4.1032 | אמונה |  |
| 1:4.1033 | זרה |  |
| 1:4.1034 | אכי״ר | . |


### Translation under review (1:4)

English: And in the third year the war broke out against the good studies and the pleasant arrangements, as commanded upon us by our Rabbis of blessed memory. The opponents spread an evil slander, saying that through this the children would go out, God forbid, into bad education, and many of the fools believed their lies and took their children out of the school. Also the foolish boys who had gone out of the school are hostile toward their teachers, because in their humble opinion they are heretics and stand outside the community. But the enlightened among them know and recognise that a lie is in the right hand of those who spread this evil slander, and they were greatly distressed that their fathers had taken them out of school in their folly! And when many went astray, the hand of the obscurantists grew strong and by lying informing they informed against the teachers before the government, that they wished to incite the inhabitants of the land. and that they were writers of libels about the government, to the government of England, to shelter in its shadow, and also to H.N.Ts. And because of this informing the teachers were led to prison, and also expelled from the school. And they brought in their place another teacher, who knew no Hebrew at all and no Arabic as it should be. And he had no duty to teach them except the weekly portion, with the Targum of Onkelos and the Targum of Rabbi Saadia Gaon in Arabic. and the Haftarah with the Targum of Jonathan in the Aramaic, with no understanding at all, like the other teachers! And after a number of years the school closed entirely! All this came upon us because we did not object, for the honour of our Creator, when we were hearing from those who study the Zohar and other Kabbalists, things that are forbidden to hear, and we were silent before them as our Rabbis of blessed memory said: "R. A. was seized for heresy, etc.; perhaps a word of heresy, etc." and there was fulfilled in us the saying of our Rabbis of blessed memory: "Whoever concedes to them falls into their hands; whoever bathes at their expense, his is theirs." Thus far! Therefore we abandoned our way of studying the languages, Hebrew and Arabic, in their grammar, against our will, because of the wrath of the oppressor, through false informing, And we turned to the path of holiness, to the honour of our Father who is in heaven, in the nullification of the opinions of the philosopher, the false prophet, the author of the Zohar, who gathered into his hands the opinions of ancient philosophy, the worshippers of idols, images and hidden powers, which they had invented from their own heart. And the author of the Zohar seized their opinions, and thought to bring them into the faith of our holy Torah, and falsely called his name by the name of Rabbi Shimon bar Yochai the Tanna, and said that the Holy One, blessed be He, who is called by him Atika Kadisha, was revealed to him and revealed to him hidden mysteries. And he hung his prophecy and his ideas on Rabbi Shimon bar Yochai the Tanna, peace be upon him. And he made his composition, the Zohar, and said that the higher ones agreed to his composition, because this philosopher saw that the Mishnah, the two Talmuds and the true midrashim of our Rabbis of blessed memory stood against him in the belief in His oneness, may He be blessed, therefore he thought to disqualify them, and he despised and mocked them and all who occupy themselves with them, and he called them by names of contempt and abasement, as the reader will see within! And he promised that Israel would not merit redemption until they turned from their way of occupying themselves with the Mishnah and the Talmud. Only they should occupy themselves with his new Torah, full of idols of images, faces of monstrous form, and various couplings. And by cunning and deceit they buried it in a chest in the everlasting ground to grow old, and it began to rot, after many days they published it and brought it out through a king of kings of the idolaters, as Rabbi Chida wrote in his book Shem ha-Gedolim, in the section on books, as if it were an ancient composition of Rabbi Shimon bar Yochai the Tanna. And many of our later rabbis were drawn after the aforementioned philosopher with eyes shut, for he took them with the smoothness of his lips and his trimmings, and they did not give heed to his lies! Heaven forbid that any Jew should believe that Rabbi Shimon bar Yochai the Tanna or another of our Rabbis of blessed memory believed in such beliefs to exchange and replace the LORD our God, who made known His ways to Moses and His deeds to the children of Israel, for "He is gracious and compassionate, slow to anger and abounding in kindness," etc. for a strange god, quick to anger (Ze'ir Anpin), and to join and associate with Him five partzufim whose existence has not been established even by proof (Rabbi Ovadiah Sforno, Deuteronomy 13:3; and there 29:25; and see there 28:14) and to call them by the name of the LORD our God, and the LORD, the God of truth, whose existence has been established by many proofs, strong as a cast mirror, as Rabbi Saadia Gaon wrote in Emunot ve-De'ot, and the Rabbi, the author of Chovot ha-Levavot, and the Rambam in the Guide for the Perplexed and in the Yad, let us abandon Him and leave Him, and say that He has no name and no appellation, and let us worship the partzufim and the forms which, according to his words, were created and descended from Him! And we have already expounded at length within, with clear proofs from our Rabbis of blessed memory, on the nullification of these strange opinions of the philosopher, the false prophet, the aforementioned! for it is the aim of our holy Torah to distance us from belief in idols, whether corporeal or spiritual. and to know that the LORD is God; there is no other besides Him. and He has neither father nor brother nor son (Rabbah, parashat Yitro), and did not receive His kingship from another, since He is not destined to hand it on to another, as our Rabbis of blessed memory explained (Rabbah, Bamidbar; Shir ha-Shirim, parashah 1; Yerushalmi, Sanhedrin, chapter Dinei Mamonot). on the verse "I am the first and I am the last, and besides Me there is no God." and on what Nebuchadnezzar said: "and the appearance of the fourth is like a son of the gods!" And I have seen that Rabbi Menchem di Lonzano, in the book Derekh ha-Chayyim, who opened his mouth wide to speak error about the holy pietist, the author of Chovot ha-Levavot, concerning his saying, in the name of the philosopher, that one could not worship the Cause of causes, except a prophet of his generation by nature, or a philosopher outstanding in what he acquired from wisdom; but others worship other than Him, because they do not understand: the result is only composite! And at these words of the pietist he took offence and thought that here was ground for complaint against him, and found a pretext to mislead the listeners and alarm them by saying: "Is it that such-and-such and such-and-such of our later famous sages worship other than Him?" And he counted them, in his language, one by one by name. and he did not take to heart that many of the sages had grown fat, went along with eyes shut and believed the words of the author of the Zohar who called himself by the name of Rabbi Shimon bar Yochai the Tanna, in order to attach his false prophecies to the name of the Tanna, and spoke slander against the LORD our God, who was revealed to him in his study house as Atika Kadisha, and revealed to him hidden secrets and hidden powers (Zohar, Bereshit 22). for the Creator is the idol called Abba, who used to command his consort, the she-idol Imma, and she does all that he says to her, except the creation of the first man, which she, Imma, said and did, not with the will of Abba! And in Parashat Vayera, folio 112, and likewise in Parashat Balak, folio 191, he ascribes our worship and our prayers to a strange god, quick to anger (Ze'ir Anpin), who did not make the heavens and the earth. And he calls these idols by the name of the LORD our God. And about Ze'ir Anpin, Jeremiah the prophet, peace be upon him, proclaimed "Thus shall you say to them: the gods that did not make the heavens and the earth shall perish from the earth and from under these heavens." for it is clear in many places in the Zohar that He is created and not creator (see Zohar, Bereshit 22; and Yosher Levav, folio 3b). And yet he says to worship him, and to bless him, and to accept him as God, and to fear his consort! And many of our later rabbis, whom he mentioned in the book Derekh ha-Chayyim, in their amazement at the pietist, the author of Chovot ha-Levavot, went along with eyes shut and were drawn after him and believed his belief and accepted upon themselves the idol, quick to anger (Ze'ir Anpin), as God! And all this came upon them because they did not contemplate the wisdom of the Creator of all, which is to investigate and examine the works of Creation as explained in the words of the Rabbi, the author of Moreh Nevukhim, and in the words of the pietist, the author of Chovot ha-Levavot. And our master Rabbi Saadia Gaon and Rabbi Yehudah Halevi in the Kuzari, as the prophet, peace be upon him, said: "They do not regard the deed of the LORD." for they thought their words were Greek wisdom, and they kept far from it, as far as a bow-shot, and they also did not take to heart the true words of our forefathers of blessed memory in the midrashim that Abraham our father did not recognise his Creator except by contemplating the work of the LORD, and investigated and expounded them and brought proofs from them until he attained the truth; then the Holy One, blessed be He, was revealed to him and said to him: "You love righteousness and hate wickedness," etc. A parable: to one walking on the road who saw a castle, etc. (Midrash ha-Gadol, at the beginning of Parashat Lekh Lekha). And perforce we must acknowledge the truth that the words of the pietist, the author of Chovot ha-Levavot, clearly say that one cannot worship the God of truth except, etc. but others worship other than Him, as the prophet said: "For many days Israel was without a true God," and as our eyes see and our ears hear, to the point that a man can say to his fellow! "Your people is my people, and your God is not my God!" And to the God of truth let us pour out our prayer, to guide us on paths of uprightness, and to deliver us from every strange belief. Amen, may it be His will.

Literal: and-in-the-period-of the-year the-third broke-out the-war against the-studies the-good and-the-arrangements the-pleasant as-commanded upon-us from-our-Rabbis-of-blessed-memory and-they-spread the-opponents slander evil saying that through the-hand-of this would-go-out the-children God-forbid into-education bad and-many of-the-fools believed their-lies and-took-out [object-marker] their-children from-the-house-of the-school and-also the-boys the-foolish who had-gone-out from-the-house-of the-school hostile [object-marker] their-teachers because in-the-humbleness-of their-opinion they heretics and-stand-outside from the-community but the-enlightened among-them know and-recognise that a-lie in-the-right-hand of-the-spreaders of-the-slander the-evil this and-they-were-distressed very that they-had-taken-them-out their-fathers from-the-house-of the-school in-their-folly and-when went-astray many grew-strong the-hand of-the-obscurantists and-by-informing of-falsehood they-informed against the-teachers before the-government that-they wished to-incite [object-marker] the-inhabitants-of the-land and-that-they were-writers-of libel about the-government to the-government-of England to-shelter in-the-shadow of-its-government and-also to H.N.Ts. and-because-of this the-informing were-led the-teachers to the-house-of the-prison and-also were-expelled from-the-house-of the-school and-they-brought in-their-place a-teacher another who not knew not Hebrew on any-creature and-not Arabic as-is-proper and-there-is-not upon-him to-teach-them only the-portion-of the-week with the-Targum of-Onkelos and-the-Targum of-Rabbi-Saadia-Gaon in-Arabic and-the-Haftarah with the-Targum of-Jonathan in-the-Aramaic with-no understanding at-all like-the-rest-of the-teachers and-after a-number-of ceased the-house-of the-school entirely all this came-upon-us because that not we-objected for-the-honour-of our-Creator at-the-time when-we-were hearing from-those-who-study the-Zohar and-the-rest-of the-Kabbalists things that-are-forbidden to-hear and-we-were-silent before-them as they-said our-Rabbis-of-blessed-memory R.-A. was-seized for-heresy etc. perhaps a-word of-heresy etc. and-was-fulfilled in-us the-saying of-our-Rabbis-of-blessed-memory whoever-concedes to-them falls into-their-hands whoever-bathes at-their-expense his theirs thus-far therefore we-abandoned our-way in-the-study-of the-languages Hebrew and-Arabic in-their-grammar against our-will because-of the-wrath-of the-oppressor through the-hand-of informing false and-we-turned to the-way-of holiness the-honour-of our-Father who-is-in-heaven in-the-nullification-of the-opinions-of the-philosopher the-prophet of-falsehood the-author-of the-Zohar who gathered in-his-hands the-opinions-of philosophy ancient the-worshippers-of the-idols images and-powers hidden which invented from-their-heart and-the-author-of the-Zohar seized their-opinions and-thought to-bring-them-in into-the-faith-of our-Torah the-holy and-called [object-marker] his-name by-the-name-of Rabbi-Shimon-bar-Yochai the-Tanna falsely and-said that-was-revealed to-him the-Holy-One,-blessed-be-He who-is-called by-him Atika Kadisha and-revealed to-him mysteries hidden and-hung his-prophecy and-his-ideas-on Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him and-he-made [object-marker] his-composition the-Zohar and-said that-the-higher-ones agreed to-his-composition because saw the-philosopher this that the-Mishnah and-the-two Talmuds and-the-midrashim-of our-Rabbis-of-blessed-memory the-true stood against-him in-the-belief-in His-oneness may-He-be-blessed therefore he-thought to-disqualify-them and-he-despised and-mocked them and-all those-who-occupy themselves-with-them and-he-called them by-names of contempt and-abasement as will-see the-reader within and-he-promised that-not would-merit Israel redemption until that return from-their-way from-occupying-themselves with-the-Mishnah and-the-Talmud only they-should-occupy with-his-Torah the-new the-full idols-of images and-faces of monstrous and-couplings various and-by-cunning and-by-deceit they-buried in-a-chest in-the-ground everlasting to-grow-old and-began to-rot after many-days they-published and-brought-out through a-king of-kings the-idolaters as-wrote Rabbi Chida in-his-book Shem ha-Gedolim in-the-section on-books as-if it a-composition ancient from-Rabbi-Shimon-bar-Yochai the-Tanna and-many of-our-rabbis the-later were-drawn after the-philosopher the-aforementioned with-closing of-the-eye for he-took them with-smoothness of-his-lips and-his-trimmings and-not they-set heart to-his-lies Heaven-forbid that-anyone son-of Israel to-believe that Rabbi-Shimon-bar-Yochai the-Tanna or another of-our-Rabbis-of-blessed-memory believed in-beliefs such to-exchange and-to-replace [object-marker] the-LORD our-God who made-known His-ways to-Moses to-the-children-of Israel His-deeds for gracious and-compassionate He slow to-anger and-abounding in-kindness etc. for-a-god strange quick to-anger Ze'ir-Anpin and-to-conjoin and-to-associate with-Him five partzufim whose not has-been-established even their-existence by-proof Rabbi Ovadiah Sforno Deuteronomy 13 3 and-there 29 25 and-see-there 28 14 and-to-call-them by-the-name-of the-LORD our-God and-[object-marker] the-LORD God of-truth who-has-been-established His-existence by-many proofs like-a-mirror of-cast strong as wrote Rabbi-Saadia-Gaon in-the-book Emunot ve-De'ot and-the-Rabbi the-author-of Chovot ha-Levavot and-the-Rambam in-the-Guide for-the-Perplexed and-in-the-Yad let-us-abandon and-leave and-say that-there-is-not for-Him a-name and-not an-appellation and-let-us-worship [object-marker] the-partzufim and-the-forms which according-to his-words were-created and-descended-in-a-chain from-Him and-already we-have-expounded-at-length within with-proofs clear from-our-Rabbis-of-blessed-memory on-the-nullification-of the-opinions the-strange these of the-philosopher the-prophet of-falsehood the-aforementioned for it is-the-aim-of our-Torah holy to-distance-us from-belief-in idols whether corporeal or spiritual and-to-know that the-LORD is God there-is-no more besides-Him and-He-has-not for-Him neither father nor brother nor son Rabbah parashat Yitro and-did-not receive His-kingship from-another since-He-is-not destined to-hand-it-over to-another as explained our-Rabbis of-blessed-memory Rabbah Bamidbar Shir ha-Shirim parashah 1 Yerushalmi Sanhedrin chapter Dinei Mamonot on the-verse I-am the-first and-I-am the-last and-besides-Me there-is-no God and-on what that-said Nebuchadnezzar and-the-appearance of the-fourth is-like a-son of-gods and-I-have-seen that-Rabbi Menchem di Lonzano in-the-book Derekh ha-Chayyim who-widened his-mouth to-speak error about the-pietist the-holy the-author-of Chovot ha-Levavot concerning his-saying in-the-name-of the-philosopher that-said that-not could worship the-cause-of the-causes except a-prophet of-his-generation by-nature or a-philosopher outstanding in-what he-acquired from wisdom but others worship other-than-Him because they-are-not understanding the-result-is only composite and-at the-words-of the-pietist these took-offence and-thought that here a-ground of-complaint against-him and-a-pretext found to-mislead the-listeners and-to-alarm-them by-saying is-it-that such-and-such and-such-and-such of-our-sages the-later the-famous worship other-than-Him and-he counted-them in-his-language one one by-name and-not he-placed to his-heart that many of-the-sages had-grown-fat went-along with-closing of-eyes and-believed the-words-of the-author-of the-Zohar who called [object-marker] himself by-the-name-of Rabbi-Shimon-bar-Yochai the-Tanna in-order-to attach his-prophecies false to-the-name-of of the-Tanna and-spoke slander against the-LORD our-God who-was-revealed to-him in-the-house-of his-study Atika Kadisha and-revealed to-him secrets hidden and-powers hidden Zohar Bereshit 22 for the-Creator is the-idol called Abba who-was commanding his-consort the-she-idol Imma and-she does all that he-says to-her except for-the-creation-of man the-first which-she Imma said and-did [object-marker] man the-first not with-the-will of-Abba and-in-Parashat Vayera folio 112 and-likewise in-Parashat Balak folio 191 ascribes our-worship and-our-prayers to-a-god strange quick to-anger Ze'ir Anpin who not made [object-marker] the-heavens and-[object-marker] the-earth and-he-calls he to-the-idols these by-the-name-of the-LORD our-God and-about Ze'ir Anpin proclaimed Jeremiah the-prophet peace-be-upon-him thus shall-you-say to-them the-gods that the-heavens and-the-earth not made shall-perish from the-earth and-from under the-heavens these for-indeed it-is-clear in-the-Zohar in-many places that-He is-created and-not creator see Zohar Bereshit 22 and-Yosher Levav folio 3 b and-yet he says to-worship-him and-to-bless-him and-to-accept-him upon-him as-God and-to-fear his-consort and-many of-our-rabbis the-later who-mentioned in-the-book Derekh ha-Chayyim in-their-amazement at the-pietist the-author-of Chovot ha-Levavot went with-closing of-the-eye and-were-drawn after-him and-believed in-his-belief and-accepted the-idol quick to-anger Ze'ir-Anpin upon-them as-God and-all this caused to-them because that-not they-contemplated the-wisdom-of the-Creator of-all which-is to-investigate and-examine the-works-of Creation as-explained in-the-words-of the-Rabbi the-author-of Moreh-Nevukhim and-in-the-words-of the-pietist the-author-of Chovot ha-Levavot and-our-master Saadia Gaon and-Rabbi-Yehudah Halevi in-the-Kuzari as said the-prophet peace-be-upon-him and-[object-marker] the-deed of-the-LORD not they-regard for they-thought their-words to-be-wisdom Greek and-they-kept-far from-it far as-far-as a-bow and-also not they-set to their-heart the-words-of our-forefathers of-blessed-memory true in-the-midrashim that-not recognised Abraham our-father [object-marker] his-Creator except by-contemplating the-work of-the-LORD and-investigated and-expounded them and-brought proofs from-them until he-attained the-truth then was-revealed to-him the-Holy-One,-blessed-be-He and-said to-him You-love righteousness and-hate wickedness etc. a-parable to-one who-was walking on-the-road and-saw a-castle etc. Midrash ha-Gadol the-beginning-of parashat Lekh Lekha and perforce we-acknowledge to the-truth that-the-words-of the-pietist the-author-of Chovot ha-Levavot clearly they-said that-not one-can worship [object-marker] God of-truth except etc. but others worship other-than-Him as-the-saying-of the-prophet for-many days Israel without God of-truth and-as our-eyes see and-our-ears hear to-the-point that-it-is-possible that-one-should-say a-man to-his-fellow your-people my-people and-your-God not my-God and-to the-God-of truth let-us-pour-out our-prayer to-guide-us in-the-paths-of uprightness and-to-deliver-us from-every belief strange Amen,-may-it-be-His-will

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:4.1 | and in the period of |  |  |  |
| 1:4.2 | the year |  |  |  |
| 1:4.3 | the third |  |  |  |
| 1:4.4 | broke out |  |  |  |
| 1:4.5 | the war |  |  |  |
| 1:4.6 | against |  |  |  |
| 1:4.7 | the studies |  |  |  |
| 1:4.8 | the good |  |  |  |
| 1:4.9 | and the arrangements |  |  |  |
| 1:4.10 | the pleasant |  |  |  |
| 1:4.11 | as commanded |  |  |  |
| 1:4.12 | upon us |  |  |  |
| 1:4.13 | from our Rabbis of blessed memory |  |  |  |
| 1:4.14 | and they spread |  |  |  |
| 1:4.15 | the opponents |  |  |  |
| 1:4.16 | slander |  |  |  |
| 1:4.17 | evil |  |  |  |
| 1:4.18 | saying |  |  |  |
| 1:4.19 | that |  |  |  |
| 1:4.20 | through |  |  |  |
| 1:4.21 | the hand of |  |  |  |
| 1:4.22 | this |  |  |  |
| 1:4.23 | would go out |  |  |  |
| 1:4.24 | the children |  |  |  |
| 1:4.25 | God forbid |  |  | Abbreviation of 'chas ve-shalom': God forbid. |
| 1:4.26 | into education |  |  |  |
| 1:4.27 | bad |  |  |  |
| 1:4.28 | and many |  |  |  |
| 1:4.29 | of the fools |  |  |  |
| 1:4.30 | believed |  |  |  |
| 1:4.31 | their lies |  |  |  |
| 1:4.32 | and took out |  |  |  |
| 1:4.33 | [object marker] |  |  |  |
| 1:4.34 | their children |  |  |  |
| 1:4.35 | from the house of |  |  |  |
| 1:4.36 | the school |  |  |  |
| 1:4.37 | and also |  |  |  |
| 1:4.38 | the boys |  |  |  |
| 1:4.39 | the foolish |  |  |  |
| 1:4.40 | who |  |  |  |
| 1:4.41 | had gone out |  |  |  |
| 1:4.42 | from the house of |  |  |  |
| 1:4.43 | the school |  |  |  |
| 1:4.44 | hostile |  |  |  |
| 1:4.45 | [object marker] |  |  |  |
| 1:4.46 | their teachers |  |  |  |
| 1:4.47 | because |  |  |  |
| 1:4.48 | in the humbleness of |  |  |  |
| 1:4.49 | their opinion |  |  |  |
| 1:4.50 | they |  |  |  |
| 1:4.51 | heretics |  |  |  |
| 1:4.52 | and stand outside |  |  |  |
| 1:4.53 | from |  |  |  |
| 1:4.54 | the community |  |  |  |
| 1:4.55 | but |  |  |  |
| 1:4.56 | the enlightened |  |  |  |
| 1:4.57 | among them |  |  |  |
| 1:4.58 | know |  |  |  |
| 1:4.59 | and recognise |  |  |  |
| 1:4.60 | that |  |  |  |
| 1:4.61 | a lie |  |  |  |
| 1:4.62 | in the right hand |  |  |  |
| 1:4.63 | of the spreaders |  |  |  |
| 1:4.64 | of the slander |  |  |  |
| 1:4.65 | the evil |  |  |  |
| 1:4.66 | this |  |  |  |
| 1:4.67 | and they were distressed |  |  |  |
| 1:4.68 | very |  |  |  |
| 1:4.69 | that |  |  |  |
| 1:4.70 | they had taken them out |  |  |  |
| 1:4.71 | their fathers |  |  |  |
| 1:4.72 | from the house of |  |  |  |
| 1:4.73 | the school |  |  |  |
| 1:4.74 | in their folly |  |  |  |
| 1:4.75 | and when |  |  |  |
| 1:4.76 | went astray |  |  |  |
| 1:4.77 | many |  |  |  |
| 1:4.78 | grew strong |  |  |  |
| 1:4.79 | the hand |  |  |  |
| 1:4.80 | of the obscurantists |  |  |  |
| 1:4.81 | and by informing |  |  |  |
| 1:4.82 | of falsehood |  |  |  |
| 1:4.83 | they informed |  |  |  |
| 1:4.84 | against |  |  |  |
| 1:4.85 | the teachers |  |  |  |
| 1:4.86 | before |  |  |  |
| 1:4.87 | the government |  |  |  |
| 1:4.88 | that they |  |  |  |
| 1:4.89 | wished |  |  |  |
| 1:4.90 | to incite |  |  |  |
| 1:4.91 | [object marker] |  |  |  |
| 1:4.92 | the inhabitants of |  |  |  |
| 1:4.93 | the land |  |  |  |
| 1:4.94 | and that they |  |  |  |
| 1:4.95 | were writers of |  |  |  |
| 1:4.96 | libel |  |  |  |
| 1:4.97 | about |  |  |  |
| 1:4.98 | the government |  |  |  |
| 1:4.99 | to |  |  |  |
| 1:4.100 | the government of |  |  |  |
| 1:4.101 | England |  |  |  |
| 1:4.102 | to shelter |  |  |  |
| 1:4.103 | in the shadow |  |  |  |
| 1:4.104 | of its government |  |  |  |
| 1:4.105 | and also |  |  |  |
| 1:4.106 | to |  |  |  |
| 1:4.107 | H.N.Ts. |  |  | Left unexpanded in the English (see apparatus); the authority it names is not certain. |
| 1:4.108 | and because of |  |  |  |
| 1:4.109 | this |  |  |  |
| 1:4.110 | the informing |  |  |  |
| 1:4.111 | were led |  |  |  |
| 1:4.112 | the teachers |  |  |  |
| 1:4.113 | to |  |  |  |
| 1:4.114 | the house of |  |  |  |
| 1:4.115 | the prison |  |  |  |
| 1:4.116 | and also |  |  |  |
| 1:4.117 | were expelled |  |  |  |
| 1:4.118 | from the house of |  |  |  |
| 1:4.119 | the school |  |  |  |
| 1:4.120 | and they brought |  |  |  |
| 1:4.121 | in their place |  |  |  |
| 1:4.122 | a teacher |  |  |  |
| 1:4.123 | another |  |  |  |
| 1:4.124 | who |  |  |  |
| 1:4.125 | not |  |  |  |
| 1:4.126 | knew |  |  |  |
| 1:4.127 | not |  |  |  |
| 1:4.128 | Hebrew |  |  |  |
| 1:4.129 | on |  |  | Idiom: 'al briyah' with a negative means 'at all'. |
| 1:4.130 | any creature |  |  |  |
| 1:4.131 | and not |  |  |  |
| 1:4.132 | Arabic |  |  |  |
| 1:4.133 | as is proper |  |  |  |
| 1:4.134 | and there is not |  |  |  |
| 1:4.135 | upon him |  |  |  |
| 1:4.136 | to teach them |  |  |  |
| 1:4.137 | only |  |  |  |
| 1:4.138 | the portion of |  |  |  |
| 1:4.139 | the week |  |  |  |
| 1:4.140 | with |  |  |  |
| 1:4.141 | the Targum |  |  |  |
| 1:4.142 | of Onkelos |  |  |  |
| 1:4.143 | and the Targum |  |  |  |
| 1:4.144 | of Rabbi Saadia Gaon |  |  |  |
| 1:4.145 | in Arabic |  |  |  |
| 1:4.146 | and the Haftarah |  |  |  |
| 1:4.147 | with |  |  |  |
| 1:4.148 | the Targum |  |  |  |
| 1:4.149 | of Jonathan |  |  |  |
| 1:4.150 | in the Aramaic |  |  | Read as 'in the Aramaic'; the expansion is not certain (see apparatus). |
| 1:4.151 | with no |  |  |  |
| 1:4.152 | understanding |  |  |  |
| 1:4.153 | at all |  |  |  |
| 1:4.154 | like the rest of |  |  |  |
| 1:4.155 | the teachers |  |  |  |
| 1:4.156 | and after |  |  |  |
| 1:4.157 | a number of |  |  |  |
| 1:4.158 | ceased |  |  |  |
| 1:4.159 | the house of |  |  |  |
| 1:4.160 | the school |  |  |  |
| 1:4.161 | entirely |  |  |  |
| 1:4.162 | all |  |  |  |
| 1:4.163 | this |  |  |  |
| 1:4.164 | came upon us |  |  |  |
| 1:4.165 | because |  |  |  |
| 1:4.166 | that |  |  |  |
| 1:4.167 | not |  |  |  |
| 1:4.168 | we objected |  |  |  |
| 1:4.169 | for the honour of |  |  |  |
| 1:4.170 | our Creator |  |  |  |
| 1:4.171 | at the time |  |  |  |
| 1:4.172 | when we were |  |  |  |
| 1:4.173 | hearing |  |  |  |
| 1:4.174 | from those who study |  |  |  |
| 1:4.175 | the Zohar |  |  |  |
| 1:4.176 | and the rest of |  |  |  |
| 1:4.177 | the Kabbalists |  |  |  |
| 1:4.178 | things |  |  |  |
| 1:4.179 | that are forbidden |  |  |  |
| 1:4.180 | to hear |  |  |  |
| 1:4.181 | and we were silent |  |  |  |
| 1:4.182 | before them |  |  |  |
| 1:4.183 | as |  |  |  |
| 1:4.184 | they said |  |  |  |
| 1:4.185 | our Rabbis of blessed memory |  |  |  |
| 1:4.186 | R. A. |  |  | The author's abbreviation for a Talmudic Rabbi; left as written (see apparatus). |
| 1:4.187 | was seized |  |  |  |
| 1:4.188 | for heresy |  |  |  |
| 1:4.189 | etc. |  |  |  |
| 1:4.190 | perhaps |  |  |  |
| 1:4.191 | a word |  |  |  |
| 1:4.192 | of heresy |  |  |  |
| 1:4.193 | etc. |  |  |  |
| 1:4.194 | and was fulfilled |  |  |  |
| 1:4.195 | in us |  |  |  |
| 1:4.196 | the saying |  |  |  |
| 1:4.197 | of our Rabbis of blessed memory |  |  |  |
| 1:4.198 | whoever concedes |  |  | Aramaic: 'whoever concedes to them'. |
| 1:4.199 | to them |  |  |  |
| 1:4.200 | falls |  |  |  |
| 1:4.201 | into their hands |  |  |  |
| 1:4.202 | whoever bathes |  |  |  |
| 1:4.203 | at their expense |  |  |  |
| 1:4.204 | his |  |  |  |
| 1:4.205 | theirs |  |  |  |
| 1:4.206 | thus far |  |  |  |
| 1:4.207 | therefore |  |  |  |
| 1:4.208 | we abandoned |  |  |  |
| 1:4.209 | our way |  |  |  |
| 1:4.210 | in the study of |  |  |  |
| 1:4.211 | the languages |  |  |  |
| 1:4.212 | Hebrew |  |  |  |
| 1:4.213 | and Arabic |  |  |  |
| 1:4.214 | in their grammar |  |  |  |
| 1:4.215 | against |  |  |  |
| 1:4.216 | our will |  |  |  |
| 1:4.217 | because of |  |  |  |
| 1:4.218 | the wrath of |  |  |  |
| 1:4.219 | the oppressor |  |  |  |
| 1:4.220 | through |  |  |  |
| 1:4.221 | the hand of |  |  |  |
| 1:4.222 | informing |  |  |  |
| 1:4.223 | false |  |  |  |
| 1:4.224 | and we turned |  |  |  |
| 1:4.225 | to |  |  |  |
| 1:4.226 | the way of |  |  |  |
| 1:4.227 | holiness |  |  |  |
| 1:4.228 | the honour of |  |  |  |
| 1:4.229 | our Father |  |  |  |
| 1:4.230 | who is in heaven |  |  |  |
| 1:4.231 | in the nullification of |  |  |  |
| 1:4.232 | the opinions of |  |  |  |
| 1:4.233 | the philosopher |  |  |  |
| 1:4.234 | the prophet |  |  |  |
| 1:4.235 | of falsehood |  |  |  |
| 1:4.236 | the author of |  |  |  |
| 1:4.237 | the Zohar |  |  |  |
| 1:4.238 | who |  |  |  |
| 1:4.239 | gathered |  |  |  |
| 1:4.240 | in his hands |  |  |  |
| 1:4.241 | the opinions of |  |  |  |
| 1:4.242 | philosophy |  |  |  |
| 1:4.243 | ancient |  |  |  |
| 1:4.244 | the worshippers of |  |  |  |
| 1:4.245 | the idols |  |  |  |
| 1:4.246 | images |  |  |  |
| 1:4.247 | and powers |  |  |  |
| 1:4.248 | hidden |  |  |  |
| 1:4.249 | which |  |  |  |
| 1:4.250 | invented |  |  |  |
| 1:4.251 | from their heart |  |  |  |
| 1:4.252 | and the author of |  |  |  |
| 1:4.253 | the Zohar |  |  |  |
| 1:4.254 | seized |  |  |  |
| 1:4.255 | their opinions |  |  |  |
| 1:4.256 | and thought |  |  |  |
| 1:4.257 | to bring them in |  |  |  |
| 1:4.258 | into the faith of |  |  |  |
| 1:4.259 | our Torah |  |  |  |
| 1:4.260 | the holy |  |  |  |
| 1:4.261 | and called |  |  |  |
| 1:4.262 | [object marker] |  |  |  |
| 1:4.263 | his name |  |  |  |
| 1:4.264 | by the name of |  |  |  |
| 1:4.265 | Rabbi Shimon bar Yochai |  |  |  |
| 1:4.266 | the Tanna |  |  |  |
| 1:4.267 | falsely |  |  |  |
| 1:4.268 | and said |  |  |  |
| 1:4.269 | that was revealed |  |  |  |
| 1:4.270 | to him |  |  |  |
| 1:4.271 | the Holy One, blessed be He |  |  |  |
| 1:4.272 | who is called |  |  |  |
| 1:4.273 | by him |  |  |  |
| 1:4.274 | Atika |  |  |  |
| 1:4.275 | Kadisha |  |  |  |
| 1:4.276 | and revealed |  |  |  |
| 1:4.277 | to him |  |  |  |
| 1:4.278 | mysteries |  |  |  |
| 1:4.279 | hidden |  |  |  |
| 1:4.280 | and hung |  |  |  |
| 1:4.281 | his prophecy |  |  |  |
| 1:4.282 | and his ideas on |  |  |  |
| 1:4.283 | Rabbi Shimon bar Yochai |  |  |  |
| 1:4.284 | the Tanna |  |  |  |
| 1:4.285 | peace be upon him |  |  |  |
| 1:4.286 | and he made |  |  |  |
| 1:4.287 | [object marker] |  |  |  |
| 1:4.288 | his composition |  |  |  |
| 1:4.289 | the Zohar |  |  |  |
| 1:4.290 | and said |  |  |  |
| 1:4.291 | that the higher ones |  |  |  |
| 1:4.292 | agreed |  |  |  |
| 1:4.293 | to his composition |  |  |  |
| 1:4.294 | because |  |  |  |
| 1:4.295 | saw |  |  |  |
| 1:4.296 | the philosopher |  |  |  |
| 1:4.297 | this |  |  |  |
| 1:4.298 | that |  |  |  |
| 1:4.299 | the Mishnah |  |  |  |
| 1:4.300 | and the two |  |  |  |
| 1:4.301 | Talmuds |  |  |  |
| 1:4.302 | and the midrashim of |  |  |  |
| 1:4.303 | our Rabbis of blessed memory |  |  |  |
| 1:4.304 | the true |  |  |  |
| 1:4.305 | stood |  |  |  |
| 1:4.306 | against him |  |  |  |
| 1:4.307 | in the belief in |  |  |  |
| 1:4.308 | His oneness |  |  |  |
| 1:4.309 | may He be blessed |  |  |  |
| 1:4.310 | therefore |  |  |  |
| 1:4.311 | he thought |  |  |  |
| 1:4.312 | to disqualify them |  |  |  |
| 1:4.313 | and he despised |  |  |  |
| 1:4.314 | and mocked |  |  |  |
| 1:4.315 | them |  |  |  |
| 1:4.316 | and all |  |  |  |
| 1:4.317 | those who occupy |  |  |  |
| 1:4.318 | themselves with them |  |  |  |
| 1:4.319 | and he called |  |  |  |
| 1:4.320 | them |  |  |  |
| 1:4.321 | by names |  |  |  |
| 1:4.322 | of |  |  |  |
| 1:4.323 | contempt |  |  |  |
| 1:4.324 | and abasement |  |  |  |
| 1:4.325 | as |  |  |  |
| 1:4.326 | will see |  |  |  |
| 1:4.327 | the reader |  |  |  |
| 1:4.328 | within |  |  |  |
| 1:4.329 | and he promised |  |  |  |
| 1:4.330 | that not |  |  |  |
| 1:4.331 | would merit |  |  |  |
| 1:4.332 | Israel |  |  |  |
| 1:4.333 | redemption |  |  |  |
| 1:4.334 | until |  |  |  |
| 1:4.335 | that |  |  |  |
| 1:4.336 | return |  |  |  |
| 1:4.337 | from their way |  |  |  |
| 1:4.338 | from occupying themselves |  |  |  |
| 1:4.339 | with the Mishnah |  |  |  |
| 1:4.340 | and the Talmud |  |  |  |
| 1:4.341 | only |  |  |  |
| 1:4.342 | they should occupy |  |  |  |
| 1:4.343 | with his Torah |  |  |  |
| 1:4.344 | the new |  |  |  |
| 1:4.345 | the full |  |  |  |
| 1:4.346 | idols of |  |  |  |
| 1:4.347 | images |  |  |  |
| 1:4.348 | and faces |  |  |  |
| 1:4.349 | of |  |  |  |
| 1:4.350 | monstrous |  |  |  |
| 1:4.351 | and couplings |  |  |  |
| 1:4.352 | various |  |  |  |
| 1:4.353 | and by cunning |  |  |  |
| 1:4.354 | and by deceit |  |  |  |
| 1:4.355 | they buried |  |  |  |
| 1:4.356 | in a chest |  |  |  |
| 1:4.357 | in the ground |  |  |  |
| 1:4.358 | everlasting |  |  |  |
| 1:4.359 | to grow old |  |  |  |
| 1:4.360 | and began |  |  |  |
| 1:4.361 | to rot |  |  |  |
| 1:4.362 | after |  |  |  |
| 1:4.363 | many days |  |  |  |
| 1:4.364 | they published |  |  |  |
| 1:4.365 | and brought out |  |  |  |
| 1:4.366 | through |  |  |  |
| 1:4.367 | a king |  |  |  |
| 1:4.368 | of kings |  |  |  |
| 1:4.369 | the idolaters |  |  | Acronym for 'idolaters' (akum): worshippers of stars and constellations. |
| 1:4.370 | as wrote |  |  |  |
| 1:4.371 | Rabbi |  |  |  |
| 1:4.372 | Chida |  |  | Rabbi Chayyim Yosef David Azulai, called the Chida (acronym). |
| 1:4.373 | in his book |  |  |  |
| 1:4.374 | Shem |  |  | Shem ha-Gedolim, literally 'The Name of the Great Ones'; the title is transliterated. |
| 1:4.375 | ha-Gedolim |  |  |  |
| 1:4.376 | in the section |  |  |  |
| 1:4.377 | on books |  |  |  |
| 1:4.378 | as if |  |  |  |
| 1:4.379 | it |  |  |  |
| 1:4.380 | a composition |  |  |  |
| 1:4.381 | ancient |  |  |  |
| 1:4.382 | from Rabbi Shimon bar Yochai |  |  |  |
| 1:4.383 | the Tanna |  |  |  |
| 1:4.384 | and many |  |  |  |
| 1:4.385 | of our rabbis |  |  |  |
| 1:4.386 | the later |  |  |  |
| 1:4.387 | were drawn |  |  |  |
| 1:4.388 | after |  |  |  |
| 1:4.389 | the philosopher |  |  |  |
| 1:4.390 | the aforementioned |  |  |  |
| 1:4.391 | with closing |  |  |  |
| 1:4.392 | of the eye |  |  |  |
| 1:4.393 | for |  |  |  |
| 1:4.394 | he took |  |  |  |
| 1:4.395 | them |  |  |  |
| 1:4.396 | with smoothness |  |  |  |
| 1:4.397 | of his lips |  |  |  |
| 1:4.398 | and his trimmings |  |  |  |
| 1:4.399 | and not |  |  |  |
| 1:4.400 | they set |  |  |  |
| 1:4.401 | heart |  |  |  |
| 1:4.402 | to his lies |  |  |  |
| 1:4.403 | Heaven forbid |  |  |  |
| 1:4.404 | that anyone |  |  |  |
| 1:4.405 | son of |  |  |  |
| 1:4.406 | Israel |  |  |  |
| 1:4.407 | to believe |  |  |  |
| 1:4.408 | that |  |  |  |
| 1:4.409 | Rabbi Shimon bar Yochai |  |  |  |
| 1:4.410 | the Tanna |  |  |  |
| 1:4.411 | or |  |  |  |
| 1:4.412 | another |  |  |  |
| 1:4.413 | of our Rabbis of blessed memory |  |  |  |
| 1:4.414 | believed |  |  |  |
| 1:4.415 | in beliefs |  |  |  |
| 1:4.416 | such |  |  |  |
| 1:4.417 | to exchange |  |  |  |
| 1:4.418 | and to replace |  |  |  |
| 1:4.419 | [object marker] |  |  |  |
| 1:4.420 | the LORD |  |  |  |
| 1:4.421 | our God |  |  |  |
| 1:4.422 | who |  |  |  |
| 1:4.423 | made known |  |  |  |
| 1:4.424 | His ways |  |  |  |
| 1:4.425 | to Moses |  |  |  |
| 1:4.426 | to the children of |  |  |  |
| 1:4.427 | Israel |  |  |  |
| 1:4.428 | His deeds |  |  |  |
| 1:4.429 | for |  |  |  |
| 1:4.430 | gracious |  |  |  |
| 1:4.431 | and compassionate |  |  |  |
| 1:4.432 | He |  |  |  |
| 1:4.433 | slow |  |  | 'Slow to anger', Exod 34:6; the pair of 'quick to anger' below. |
| 1:4.434 | to anger |  |  |  |
| 1:4.435 | and abounding |  |  |  |
| 1:4.436 | in kindness |  |  |  |
| 1:4.437 | etc. |  |  |  |
| 1:4.438 | for a god |  |  |  |
| 1:4.439 | strange |  |  |  |
| 1:4.440 | quick |  |  | Calque of Ze'ir Anpin (Aramaic 'small face'), set against 'slow to anger' (Exod 34:6). Prov 14:17. |
| 1:4.441 | to anger |  |  |  |
| 1:4.442 | Ze'ir Anpin |  |  | Ze'ir Anpin: Aramaic 'small face', one of the partzufim. The calque 'quick to anger' is set against 'slow to anger' (Exod 34:6). Prov 14:17. |
| 1:4.443 | and to conjoin |  |  |  |
| 1:4.444 | and to associate |  |  |  |
| 1:4.445 | with Him |  |  |  |
| 1:4.446 | five |  |  |  |
| 1:4.447 | partzufim |  |  |  |
| 1:4.448 | whose |  |  |  |
| 1:4.449 | not |  |  |  |
| 1:4.450 | has been established |  |  |  |
| 1:4.451 | even |  |  |  |
| 1:4.452 | their existence |  |  |  |
| 1:4.453 | by proof |  |  |  |
| 1:4.454 | Rabbi |  |  |  |
| 1:4.455 | Ovadiah |  |  |  |
| 1:4.456 | Sforno |  |  |  |
| 1:4.457 | Deuteronomy |  |  |  |
| 1:4.458 | 13 |  |  |  |
| 1:4.459 | 3 |  |  |  |
| 1:4.460 | and there |  |  |  |
| 1:4.461 | 29 |  |  |  |
| 1:4.462 | 25 |  |  |  |
| 1:4.463 | and see there |  |  |  |
| 1:4.464 | 28 |  |  |  |
| 1:4.465 | 14 |  |  |  |
| 1:4.466 | and to call them |  |  |  |
| 1:4.467 | by the name of |  |  |  |
| 1:4.468 | the LORD |  |  |  |
| 1:4.469 | our God |  |  |  |
| 1:4.470 | and [object marker] |  |  |  |
| 1:4.471 | the LORD |  |  |  |
| 1:4.472 | God |  |  |  |
| 1:4.473 | of truth |  |  |  |
| 1:4.474 | who has been established |  |  |  |
| 1:4.475 | His existence |  |  |  |
| 1:4.476 | by many |  |  |  |
| 1:4.477 | proofs |  |  |  |
| 1:4.478 | like a mirror |  |  |  |
| 1:4.479 | of cast |  |  |  |
| 1:4.480 | strong |  |  |  |
| 1:4.481 | as |  |  |  |
| 1:4.482 | wrote |  |  |  |
| 1:4.483 | Rabbi Saadia Gaon |  |  |  |
| 1:4.484 | in the book |  |  |  |
| 1:4.485 | Emunot |  |  | Emunot ve-De'ot: 'Beliefs and Opinions', Saadia Gaon's book (title transliterated). |
| 1:4.486 | ve-De'ot |  |  |  |
| 1:4.487 | and the Rabbi |  |  |  |
| 1:4.488 | the author of |  |  |  |
| 1:4.489 | Chovot |  |  |  |
| 1:4.490 | ha-Levavot |  |  | Chovot ha-Levavot: 'Duties of the Heart' (title transliterated). |
| 1:4.491 | and the Rambam |  |  |  |
| 1:4.492 | in the Guide |  |  | Moreh Nevukhim: 'Guide for the Perplexed' (title transliterated). |
| 1:4.493 | for the Perplexed |  |  |  |
| 1:4.494 | and in the Yad |  |  | Yad: the Mishneh Torah (Yad ha-Chazakah), 'the Hand'. |
| 1:4.495 | let us abandon |  |  |  |
| 1:4.496 | and leave |  |  |  |
| 1:4.497 | and say |  |  |  |
| 1:4.498 | that there is not |  |  |  |
| 1:4.499 | for Him |  |  |  |
| 1:4.500 | a name |  |  |  |
| 1:4.501 | and not |  |  |  |
| 1:4.502 | an appellation |  |  |  |
| 1:4.503 | and let us worship |  |  |  |
| 1:4.504 | [object marker] |  |  |  |
| 1:4.505 | the partzufim |  |  |  |
| 1:4.506 | and the forms |  |  |  |
| 1:4.507 | which |  |  |  |
| 1:4.508 | according to |  |  |  |
| 1:4.509 | his words |  |  |  |
| 1:4.510 | were created |  |  |  |
| 1:4.511 | and descended in a chain |  |  | Hishtalshelut: 'chain of descent', the descent of the partzufim one from another. |
| 1:4.512 | from Him |  |  |  |
| 1:4.513 | and already |  |  |  |
| 1:4.514 | we have expounded at length |  |  |  |
| 1:4.515 | within |  |  |  |
| 1:4.516 | with proofs |  |  |  |
| 1:4.517 | clear |  |  |  |
| 1:4.518 | from our Rabbis of blessed memory |  |  |  |
| 1:4.519 | on the nullification of |  |  |  |
| 1:4.520 | the opinions |  |  |  |
| 1:4.521 | the strange |  |  |  |
| 1:4.522 | these |  |  |  |
| 1:4.523 | of |  |  |  |
| 1:4.524 | the philosopher |  |  |  |
| 1:4.525 | the prophet |  |  |  |
| 1:4.526 | of falsehood |  |  |  |
| 1:4.527 | the aforementioned |  |  |  |
| 1:4.528 | for |  |  |  |
| 1:4.529 | it |  |  |  |
| 1:4.530 | is the aim of |  |  |  |
| 1:4.531 | our Torah |  |  |  |
| 1:4.532 | holy |  |  |  |
| 1:4.533 | to distance us |  |  |  |
| 1:4.534 | from belief in |  |  |  |
| 1:4.535 | idols |  |  |  |
| 1:4.536 | whether |  |  |  |
| 1:4.537 | corporeal |  |  |  |
| 1:4.538 | or |  |  |  |
| 1:4.539 | spiritual |  |  |  |
| 1:4.540 | and to know |  |  |  |
| 1:4.541 | that |  |  |  |
| 1:4.542 | the LORD |  |  |  |
| 1:4.543 | is |  |  |  |
| 1:4.544 | God |  |  |  |
| 1:4.545 | there is no |  |  |  |
| 1:4.546 | more |  |  |  |
| 1:4.547 | besides Him |  |  |  |
| 1:4.548 | and He has not |  |  |  |
| 1:4.549 | for Him |  |  |  |
| 1:4.550 | neither |  |  |  |
| 1:4.551 | father |  |  |  |
| 1:4.552 | nor |  |  |  |
| 1:4.553 | brother |  |  |  |
| 1:4.554 | nor |  |  |  |
| 1:4.555 | son |  |  |  |
| 1:4.556 | Rabbah |  |  |  |
| 1:4.557 | parashat |  |  |  |
| 1:4.558 | Yitro |  |  |  |
| 1:4.559 | and did not |  |  |  |
| 1:4.560 | receive |  |  |  |
| 1:4.561 | His kingship |  |  |  |
| 1:4.562 | from another |  |  |  |
| 1:4.563 | since He is not |  |  |  |
| 1:4.564 | destined |  |  |  |
| 1:4.565 | to hand it over |  |  |  |
| 1:4.566 | to another |  |  |  |
| 1:4.567 | as |  |  |  |
| 1:4.568 | explained |  |  |  |
| 1:4.569 | our Rabbis |  |  |  |
| 1:4.570 | of blessed memory |  |  |  |
| 1:4.571 | Rabbah |  |  |  |
| 1:4.572 | Bamidbar |  |  |  |
| 1:4.573 | Shir |  |  |  |
| 1:4.574 | ha-Shirim |  |  |  |
| 1:4.575 | parashah |  |  |  |
| 1:4.576 | 1 |  |  |  |
| 1:4.577 | Yerushalmi |  |  |  |
| 1:4.578 | Sanhedrin |  |  |  |
| 1:4.579 | chapter |  |  |  |
| 1:4.580 | Dinei |  |  |  |
| 1:4.581 | Mamonot |  |  |  |
| 1:4.582 | on |  |  |  |
| 1:4.583 | the verse |  |  |  |
| 1:4.584 | I am |  |  |  |
| 1:4.585 | the first |  |  |  |
| 1:4.586 | and I am |  |  |  |
| 1:4.587 | the last |  |  |  |
| 1:4.588 | and besides Me |  |  |  |
| 1:4.589 | there is no |  |  |  |
| 1:4.590 | God |  |  |  |
| 1:4.591 | and on |  |  |  |
| 1:4.592 | what |  |  |  |
| 1:4.593 | that said |  |  |  |
| 1:4.594 | Nebuchadnezzar |  |  |  |
| 1:4.595 | and the appearance |  |  |  |
| 1:4.596 | of |  |  |  |
| 1:4.597 | the fourth |  |  |  |
| 1:4.598 | is like |  |  |  |
| 1:4.599 | a son |  |  |  |
| 1:4.600 | of gods |  |  | Biblical Aramaic (Dan 3:25). |
| 1:4.601 | and I have seen |  |  |  |
| 1:4.602 | that Rabbi |  |  |  |
| 1:4.603 | Menchem |  |  |  |
| 1:4.604 | di |  |  |  |
| 1:4.605 | Lonzano |  |  |  |
| 1:4.606 | in the book |  |  |  |
| 1:4.607 | Derekh |  |  |  |
| 1:4.608 | ha-Chayyim |  |  | Derekh ha-Chayyim: 'The Way of Life' (title transliterated). |
| 1:4.609 | who widened |  |  |  |
| 1:4.610 | his mouth |  |  |  |
| 1:4.611 | to speak |  |  |  |
| 1:4.612 | error |  |  |  |
| 1:4.613 | about |  |  |  |
| 1:4.614 | the pietist |  |  |  |
| 1:4.615 | the holy |  |  |  |
| 1:4.616 | the author of |  |  |  |
| 1:4.617 | Chovot |  |  |  |
| 1:4.618 | ha-Levavot |  |  |  |
| 1:4.619 | concerning |  |  |  |
| 1:4.620 | his saying |  |  |  |
| 1:4.621 | in the name of |  |  |  |
| 1:4.622 | the philosopher |  |  |  |
| 1:4.623 | that said |  |  |  |
| 1:4.624 | that not |  |  |  |
| 1:4.625 | could |  |  |  |
| 1:4.626 | worship |  |  |  |
| 1:4.627 | the cause of |  |  |  |
| 1:4.628 | the causes |  |  |  |
| 1:4.629 | except |  |  |  |
| 1:4.630 | a prophet |  |  |  |
| 1:4.631 | of his generation |  |  |  |
| 1:4.632 | by nature |  |  |  |
| 1:4.633 | or |  |  |  |
| 1:4.634 | a philosopher |  |  |  |
| 1:4.635 | outstanding |  |  |  |
| 1:4.636 | in what |  |  |  |
| 1:4.637 | he acquired |  |  |  |
| 1:4.638 | from |  |  |  |
| 1:4.639 | wisdom |  |  |  |
| 1:4.640 | but |  |  |  |
| 1:4.641 | others |  |  |  |
| 1:4.642 | worship |  |  |  |
| 1:4.643 | other than Him |  |  |  |
| 1:4.644 | because |  |  |  |
| 1:4.645 | they are not |  |  |  |
| 1:4.646 | understanding |  |  |  |
| 1:4.647 | the result is |  |  |  |
| 1:4.648 | only |  |  |  |
| 1:4.649 | composite |  |  |  |
| 1:4.650 | and at |  |  |  |
| 1:4.651 | the words of |  |  |  |
| 1:4.652 | the pietist |  |  |  |
| 1:4.653 | these |  |  |  |
| 1:4.654 | took offence |  |  |  |
| 1:4.655 | and thought |  |  |  |
| 1:4.656 | that |  |  |  |
| 1:4.657 | here |  |  |  |
| 1:4.658 | a ground |  |  |  |
| 1:4.659 | of complaint |  |  |  |
| 1:4.660 | against him |  |  |  |
| 1:4.661 | and a pretext |  |  |  |
| 1:4.662 | found |  |  |  |
| 1:4.663 | to mislead |  |  |  |
| 1:4.664 | the listeners |  |  |  |
| 1:4.665 | and to alarm them |  |  |  |
| 1:4.666 | by saying |  |  |  |
| 1:4.667 | is it that |  |  |  |
| 1:4.668 | such-and-such |  |  |  |
| 1:4.669 | and such-and-such |  |  |  |
| 1:4.670 | of our sages |  |  |  |
| 1:4.671 | the later |  |  |  |
| 1:4.672 | the famous |  |  |  |
| 1:4.673 | worship |  |  |  |
| 1:4.674 | other than Him |  |  |  |
| 1:4.675 | and he |  |  |  |
| 1:4.676 | counted them |  |  |  |
| 1:4.677 | in his language |  |  |  |
| 1:4.678 | one |  |  |  |
| 1:4.679 | one |  |  |  |
| 1:4.680 | by name |  |  |  |
| 1:4.681 | and not |  |  |  |
| 1:4.682 | he placed |  |  |  |
| 1:4.683 | to |  |  |  |
| 1:4.684 | his heart |  |  |  |
| 1:4.685 | that |  |  |  |
| 1:4.686 | many |  |  |  |
| 1:4.687 | of the sages |  |  |  |
| 1:4.688 | had grown fat |  |  | Pointed shamanah: 'grew fat' (cf. Deut 32:15), a figure for complacency. |
| 1:4.689 | went along |  |  |  |
| 1:4.690 | with closing |  |  |  |
| 1:4.691 | of eyes |  |  |  |
| 1:4.692 | and believed |  |  |  |
| 1:4.693 | the words of |  |  |  |
| 1:4.694 | the author of |  |  |  |
| 1:4.695 | the Zohar |  |  |  |
| 1:4.696 | who |  |  |  |
| 1:4.697 | called |  |  |  |
| 1:4.698 | [object marker] |  |  |  |
| 1:4.699 | himself |  |  |  |
| 1:4.700 | by the name of |  |  |  |
| 1:4.701 | Rabbi Shimon bar Yochai |  |  |  |
| 1:4.702 | the Tanna |  |  |  |
| 1:4.703 | in order to |  |  |  |
| 1:4.704 | attach |  |  |  |
| 1:4.705 | his prophecies |  |  |  |
| 1:4.706 | false |  |  |  |
| 1:4.707 | to the name of |  |  |  |
| 1:4.708 | of |  |  |  |
| 1:4.709 | the Tanna |  |  |  |
| 1:4.710 | and spoke |  |  |  |
| 1:4.711 | slander |  |  |  |
| 1:4.712 | against |  |  |  |
| 1:4.713 | the LORD |  |  |  |
| 1:4.714 | our God |  |  |  |
| 1:4.715 | who was revealed |  |  |  |
| 1:4.716 | to him |  |  |  |
| 1:4.717 | in the house of |  |  |  |
| 1:4.718 | his study |  |  |  |
| 1:4.719 | Atika |  |  |  |
| 1:4.720 | Kadisha |  |  |  |
| 1:4.721 | and revealed |  |  |  |
| 1:4.722 | to him |  |  |  |
| 1:4.723 | secrets |  |  |  |
| 1:4.724 | hidden |  |  |  |
| 1:4.725 | and powers |  |  |  |
| 1:4.726 | hidden |  |  |  |
| 1:4.727 | Zohar |  |  |  |
| 1:4.728 | Bereshit |  |  |  |
| 1:4.729 | 22 |  |  |  |
| 1:4.730 | for |  |  |  |
| 1:4.731 | the Creator |  |  |  |
| 1:4.732 | is |  |  |  |
| 1:4.733 | the idol |  |  |  |
| 1:4.734 | called |  |  |  |
| 1:4.735 | Abba |  |  |  |
| 1:4.736 | who was |  |  |  |
| 1:4.737 | commanding |  |  |  |
| 1:4.738 | his consort |  |  |  |
| 1:4.739 | the she-idol |  |  | The author's feminine form for the idol, kept as written as 'she-idol'. |
| 1:4.740 | Imma |  |  |  |
| 1:4.741 | and she |  |  |  |
| 1:4.742 | does |  |  |  |
| 1:4.743 | all |  |  |  |
| 1:4.744 | that |  |  |  |
| 1:4.745 | he says |  |  |  |
| 1:4.746 | to her |  |  |  |
| 1:4.747 | except |  |  |  |
| 1:4.748 | for the creation of |  |  |  |
| 1:4.749 | man |  |  |  |
| 1:4.750 | the first |  |  |  |
| 1:4.751 | which she |  |  |  |
| 1:4.752 | Imma |  |  |  |
| 1:4.753 | said |  |  |  |
| 1:4.754 | and did |  |  |  |
| 1:4.755 | [object marker] |  |  |  |
| 1:4.756 | man |  |  |  |
| 1:4.757 | the first |  |  |  |
| 1:4.758 | not |  |  |  |
| 1:4.759 | with the will |  |  |  |
| 1:4.760 | of Abba |  |  |  |
| 1:4.761 | and in Parashat |  |  |  |
| 1:4.762 | Vayera |  |  |  |
| 1:4.763 | folio |  |  |  |
| 1:4.764 | 112 |  |  |  |
| 1:4.765 | and likewise |  |  |  |
| 1:4.766 | in Parashat |  |  |  |
| 1:4.767 | Balak |  |  |  |
| 1:4.768 | folio |  |  |  |
| 1:4.769 | 191 |  |  |  |
| 1:4.770 | ascribes |  |  |  |
| 1:4.771 | our worship |  |  |  |
| 1:4.772 | and our prayers |  |  |  |
| 1:4.773 | to a god |  |  |  |
| 1:4.774 | strange |  |  |  |
| 1:4.775 | quick |  |  |  |
| 1:4.776 | to anger |  |  |  |
| 1:4.777 | Ze'ir |  |  |  |
| 1:4.778 | Anpin |  |  |  |
| 1:4.779 | who |  |  |  |
| 1:4.780 | not |  |  |  |
| 1:4.781 | made |  |  |  |
| 1:4.782 | [object marker] |  |  |  |
| 1:4.783 | the heavens |  |  |  |
| 1:4.784 | and [object marker] |  |  |  |
| 1:4.785 | the earth |  |  |  |
| 1:4.786 | and he calls |  |  |  |
| 1:4.787 | he |  |  |  |
| 1:4.788 | to the idols |  |  |  |
| 1:4.789 | these |  |  |  |
| 1:4.790 | by the name of |  |  |  |
| 1:4.791 | the LORD |  |  |  |
| 1:4.792 | our God |  |  |  |
| 1:4.793 | and about |  |  |  |
| 1:4.794 | Ze'ir |  |  |  |
| 1:4.795 | Anpin |  |  |  |
| 1:4.796 | proclaimed |  |  | 'Proclaimed concerning': the verb with 'al' takes the sense of a pronouncement. |
| 1:4.797 | Jeremiah |  |  |  |
| 1:4.798 | the prophet |  |  |  |
| 1:4.799 | peace be upon him |  |  |  |
| 1:4.800 | thus |  |  |  |
| 1:4.801 | shall you say |  |  |  |
| 1:4.802 | to them |  |  |  |
| 1:4.803 | the gods |  |  | Aramaic 'gods'. |
| 1:4.804 | that |  |  |  |
| 1:4.805 | the heavens |  |  |  |
| 1:4.806 | and the earth |  |  |  |
| 1:4.807 | not |  |  |  |
| 1:4.808 | made |  |  |  |
| 1:4.809 | shall perish |  |  |  |
| 1:4.810 | from |  |  |  |
| 1:4.811 | the earth |  |  |  |
| 1:4.812 | and from |  |  |  |
| 1:4.813 | under |  |  |  |
| 1:4.814 | the heavens |  |  |  |
| 1:4.815 | these |  |  |  |
| 1:4.816 | for indeed |  |  |  |
| 1:4.817 | it is clear |  |  |  |
| 1:4.818 | in the Zohar |  |  |  |
| 1:4.819 | in many |  |  |  |
| 1:4.820 | places |  |  |  |
| 1:4.821 | that He |  |  | 'Created, not creator': the Zohar's contrast. |
| 1:4.822 | is created |  |  |  |
| 1:4.823 | and not |  |  |  |
| 1:4.824 | creator |  |  |  |
| 1:4.825 | see |  |  |  |
| 1:4.826 | Zohar |  |  |  |
| 1:4.827 | Bereshit |  |  |  |
| 1:4.828 | 22 |  |  |  |
| 1:4.829 | and Yosher |  |  |  |
| 1:4.830 | Levav |  |  |  |
| 1:4.831 | folio |  |  |  |
| 1:4.832 | 3 |  |  |  |
| 1:4.833 | b |  |  |  |
| 1:4.834 | and yet |  |  |  |
| 1:4.835 | he |  |  |  |
| 1:4.836 | says |  |  |  |
| 1:4.837 | to worship him |  |  |  |
| 1:4.838 | and to bless him |  |  |  |
| 1:4.839 | and to accept him |  |  |  |
| 1:4.840 | upon him |  |  |  |
| 1:4.841 | as God |  |  |  |
| 1:4.842 | and to fear |  |  |  |
| 1:4.843 | his consort |  |  |  |
| 1:4.844 | and many |  |  |  |
| 1:4.845 | of our rabbis |  |  |  |
| 1:4.846 | the later |  |  |  |
| 1:4.847 | who mentioned |  |  |  |
| 1:4.848 | in the book |  |  |  |
| 1:4.849 | Derekh |  |  |  |
| 1:4.850 | ha-Chayyim |  |  |  |
| 1:4.851 | in their amazement |  |  |  |
| 1:4.852 | at |  |  |  |
| 1:4.853 | the pietist |  |  |  |
| 1:4.854 | the author of |  |  |  |
| 1:4.855 | Chovot |  |  |  |
| 1:4.856 | ha-Levavot |  |  |  |
| 1:4.857 | went |  |  |  |
| 1:4.858 | with closing |  |  |  |
| 1:4.859 | of the eye |  |  |  |
| 1:4.860 | and were drawn |  |  |  |
| 1:4.861 | after him |  |  |  |
| 1:4.862 | and believed |  |  |  |
| 1:4.863 | in his belief |  |  |  |
| 1:4.864 | and accepted |  |  |  |
| 1:4.865 | the idol |  |  |  |
| 1:4.866 | quick |  |  |  |
| 1:4.867 | to anger |  |  |  |
| 1:4.868 | Ze'ir Anpin |  |  |  |
| 1:4.869 | upon them |  |  |  |
| 1:4.870 | as God |  |  |  |
| 1:4.871 | and all |  |  |  |
| 1:4.872 | this |  |  |  |
| 1:4.873 | caused |  |  |  |
| 1:4.874 | to them |  |  |  |
| 1:4.875 | because |  |  |  |
| 1:4.876 | that not |  |  |  |
| 1:4.877 | they contemplated |  |  |  |
| 1:4.878 | the wisdom of |  |  |  |
| 1:4.879 | the Creator |  |  |  |
| 1:4.880 | of all |  |  |  |
| 1:4.881 | which is |  |  |  |
| 1:4.882 | to investigate |  |  |  |
| 1:4.883 | and examine |  |  |  |
| 1:4.884 | the works of |  |  |  |
| 1:4.885 | Creation |  |  |  |
| 1:4.886 | as explained |  |  |  |
| 1:4.887 | in the words of |  |  |  |
| 1:4.888 | the Rabbi |  |  |  |
| 1:4.889 | the author of |  |  |  |
| 1:4.890 | Moreh Nevukhim |  |  |  |
| 1:4.891 | and in the words of |  |  |  |
| 1:4.892 | the pietist |  |  |  |
| 1:4.893 | the author of |  |  |  |
| 1:4.894 | Chovot |  |  |  |
| 1:4.895 | ha-Levavot |  |  |  |
| 1:4.896 | and our master |  |  |  |
| 1:4.897 | Saadia |  |  |  |
| 1:4.898 | Gaon |  |  |  |
| 1:4.899 | and Rabbi Yehudah |  |  |  |
| 1:4.900 | Halevi |  |  |  |
| 1:4.901 | in the Kuzari |  |  |  |
| 1:4.902 | as |  |  |  |
| 1:4.903 | said |  |  |  |
| 1:4.904 | the prophet |  |  |  |
| 1:4.905 | peace be upon him |  |  |  |
| 1:4.906 | and [object marker] |  |  |  |
| 1:4.907 | the deed |  |  |  |
| 1:4.908 | of the LORD |  |  |  |
| 1:4.909 | not |  |  |  |
| 1:4.910 | they regard |  |  | Isa 5:12. |
| 1:4.911 | for |  |  |  |
| 1:4.912 | they thought |  |  |  |
| 1:4.913 | their words |  |  |  |
| 1:4.914 | to be wisdom |  |  |  |
| 1:4.915 | Greek |  |  |  |
| 1:4.916 | and they kept far |  |  |  |
| 1:4.917 | from it |  |  |  |
| 1:4.918 | far |  |  |  |
| 1:4.919 | as far as |  |  |  |
| 1:4.920 | a bow |  |  |  |
| 1:4.921 | and also |  |  |  |
| 1:4.922 | not |  |  |  |
| 1:4.923 | they set |  |  |  |
| 1:4.924 | to |  |  |  |
| 1:4.925 | their heart |  |  |  |
| 1:4.926 | the words of |  |  |  |
| 1:4.927 | our forefathers |  |  |  |
| 1:4.928 | of blessed memory |  |  |  |
| 1:4.929 | true |  |  |  |
| 1:4.930 | in the midrashim |  |  |  |
| 1:4.931 | that not |  |  |  |
| 1:4.932 | recognised |  |  |  |
| 1:4.933 | Abraham |  |  |  |
| 1:4.934 | our father |  |  |  |
| 1:4.935 | [object marker] |  |  |  |
| 1:4.936 | his Creator |  |  |  |
| 1:4.937 | except |  |  |  |
| 1:4.938 | by contemplating |  |  |  |
| 1:4.939 | the work |  |  |  |
| 1:4.940 | of the LORD |  |  |  |
| 1:4.941 | and investigated |  |  |  |
| 1:4.942 | and expounded |  |  |  |
| 1:4.943 | them |  |  |  |
| 1:4.944 | and brought |  |  |  |
| 1:4.945 | proofs |  |  |  |
| 1:4.946 | from them |  |  |  |
| 1:4.947 | until |  |  |  |
| 1:4.948 | he attained |  |  |  |
| 1:4.949 | the truth |  |  |  |
| 1:4.950 | then |  |  |  |
| 1:4.951 | was revealed |  |  |  |
| 1:4.952 | to him |  |  |  |
| 1:4.953 | the Holy One, blessed be He |  |  |  |
| 1:4.954 | and said |  |  |  |
| 1:4.955 | to him |  |  |  |
| 1:4.956 | You love |  |  |  |
| 1:4.957 | righteousness |  |  |  |
| 1:4.958 | and hate |  |  |  |
| 1:4.959 | wickedness |  |  |  |
| 1:4.960 | etc. |  |  |  |
| 1:4.961 | a parable |  |  |  |
| 1:4.962 | to one |  |  |  |
| 1:4.963 | who was |  |  |  |
| 1:4.964 | walking |  |  |  |
| 1:4.965 | on the road |  |  |  |
| 1:4.966 | and saw |  |  |  |
| 1:4.967 | a castle |  |  |  |
| 1:4.968 | etc. |  |  |  |
| 1:4.969 | Midrash |  |  |  |
| 1:4.970 | ha-Gadol |  |  |  |
| 1:4.971 | the beginning of |  |  |  |
| 1:4.972 | parashat |  |  |  |
| 1:4.973 | Lekh |  |  |  |
| 1:4.974 | Lekha |  |  |  |
| 1:4.975 | and |  |  |  |
| 1:4.976 | perforce |  |  |  |
| 1:4.977 | we acknowledge |  |  |  |
| 1:4.978 | to |  |  |  |
| 1:4.979 | the truth |  |  |  |
| 1:4.980 | that the words of |  |  |  |
| 1:4.981 | the pietist |  |  |  |
| 1:4.982 | the author of |  |  |  |
| 1:4.983 | Chovot |  |  |  |
| 1:4.984 | ha-Levavot |  |  |  |
| 1:4.985 | clearly |  |  |  |
| 1:4.986 | they said |  |  |  |
| 1:4.987 | that not |  |  |  |
| 1:4.988 | one can |  |  |  |
| 1:4.989 | worship |  |  |  |
| 1:4.990 | [object marker] |  |  |  |
| 1:4.991 | God |  |  |  |
| 1:4.992 | of truth |  |  |  |
| 1:4.993 | except |  |  |  |
| 1:4.994 | etc. |  |  |  |
| 1:4.995 | but |  |  |  |
| 1:4.996 | others |  |  |  |
| 1:4.997 | worship |  |  |  |
| 1:4.998 | other than Him |  |  |  |
| 1:4.999 | as the saying of |  |  |  |
| 1:4.1000 | the prophet |  |  |  |
| 1:4.1001 | for many |  |  |  |
| 1:4.1002 | days |  |  |  |
| 1:4.1003 | Israel |  |  |  |
| 1:4.1004 | without |  |  |  |
| 1:4.1005 | God |  |  | 2 Chr 15:3 (see apparatus). |
| 1:4.1006 | of truth |  |  |  |
| 1:4.1007 | and as |  |  |  |
| 1:4.1008 | our eyes |  |  |  |
| 1:4.1009 | see |  |  |  |
| 1:4.1010 | and our ears |  |  |  |
| 1:4.1011 | hear |  |  |  |
| 1:4.1012 | to the point |  |  |  |
| 1:4.1013 | that it is possible |  |  |  |
| 1:4.1014 | that one should say |  |  |  |
| 1:4.1015 | a man |  |  |  |
| 1:4.1016 | to his fellow |  |  |  |
| 1:4.1017 | your people |  |  |  |
| 1:4.1018 | my people |  |  |  |
| 1:4.1019 | and your God |  |  |  |
| 1:4.1020 | not |  |  |  |
| 1:4.1021 | my God |  |  |  |
| 1:4.1022 | and to |  |  |  |
| 1:4.1023 | the God of |  |  |  |
| 1:4.1024 | truth |  |  |  |
| 1:4.1025 | let us pour out |  |  |  |
| 1:4.1026 | our prayer |  |  |  |
| 1:4.1027 | to guide us |  |  |  |
| 1:4.1028 | in the paths of |  |  |  |
| 1:4.1029 | uprightness |  |  |  |
| 1:4.1030 | and to deliver us |  |  |  |
| 1:4.1031 | from every |  |  |  |
| 1:4.1032 | belief |  |  |  |
| 1:4.1033 | strange |  |  |  |
| 1:4.1034 | Amen, may it be His will |  |  |  |

Phrases:
- 1:4.1 1:4.2 1:4.3 1:4.4 1:4.5 1:4.6 1:4.7 1:4.8 1:4.9 1:4.10 1:4.11 1:4.12 1:4.13 | And in the third year the war broke out against the good studies and the pleasant arrangements, as commanded upon us by our Rabbis of blessed memory. | and-in-the-period-of the-year the-third broke-out the-war against the-studies the-good and-the-arrangements the-pleasant as-commanded upon-us from-our-Rabbis-of-blessed-memory
- 1:4.14 1:4.15 1:4.16 1:4.17 1:4.18 | The opponents spread an evil slander, saying | and-they-spread the-opponents slander evil saying
- 1:4.19 1:4.20 1:4.21 1:4.22 1:4.23 1:4.24 1:4.25 1:4.26 1:4.27 | that through this the children would go out, God forbid, into bad education, | that through the-hand-of this would-go-out the-children God-forbid into-education bad
- 1:4.28 1:4.29 1:4.30 1:4.31 1:4.32 1:4.33 1:4.34 1:4.35 1:4.36 | and many of the fools believed their lies and took their children out of the school. | and-many of-the-fools believed their-lies and-took-out [object-marker] their-children from-the-house-of the-school
- 1:4.37 1:4.38 1:4.39 1:4.40 1:4.41 1:4.42 1:4.43 | Also the foolish boys who had gone out of the school | and-also the-boys the-foolish who had-gone-out from-the-house-of the-school
- 1:4.44 1:4.45 1:4.46 | are hostile toward their teachers, | hostile [object-marker] their-teachers
- 1:4.47 1:4.48 1:4.49 1:4.50 1:4.51 1:4.52 1:4.53 1:4.54 | because in their humble opinion they are heretics and stand outside the community. | because in-the-humbleness-of their-opinion they heretics and-stand-outside from the-community
- 1:4.55 1:4.56 1:4.57 1:4.58 1:4.59 1:4.60 1:4.61 1:4.62 1:4.63 1:4.64 1:4.65 1:4.66 | But the enlightened among them know and recognise that a lie is in the right hand of those who spread this evil slander, | but the-enlightened among-them know and-recognise that a-lie in-the-right-hand of-the-spreaders of-the-slander the-evil this
- 1:4.67 1:4.68 1:4.69 1:4.70 1:4.71 1:4.72 1:4.73 1:4.74 | and they were greatly distressed that their fathers had taken them out of school in their folly! | and-they-were-distressed very that they-had-taken-them-out their-fathers from-the-house-of the-school in-their-folly
- 1:4.75 1:4.76 1:4.77 1:4.78 1:4.79 1:4.80 | And when many went astray, the hand of the obscurantists grew strong | and-when went-astray many grew-strong the-hand of-the-obscurantists
- 1:4.81 1:4.82 1:4.83 1:4.84 1:4.85 1:4.86 1:4.87 1:4.88 1:4.89 1:4.90 1:4.91 1:4.92 1:4.93 | and by lying informing they informed against the teachers before the government, that they wished to incite the inhabitants of the land. | and-by-informing of-falsehood they-informed against the-teachers before the-government that-they wished to-incite [object-marker] the-inhabitants-of the-land
- 1:4.94 1:4.95 1:4.96 1:4.97 1:4.98 1:4.99 1:4.100 1:4.101 1:4.102 1:4.103 1:4.104 1:4.105 1:4.106 1:4.107 | and that they were writers of libels about the government, to the government of England, to shelter in its shadow, and also to H.N.Ts. | and-that-they were-writers-of libel about the-government to the-government-of England to-shelter in-the-shadow of-its-government and-also to H.N.Ts.
- 1:4.108 1:4.109 1:4.110 1:4.111 1:4.112 1:4.113 1:4.114 1:4.115 1:4.116 1:4.117 1:4.118 1:4.119 | And because of this informing the teachers were led to prison, and also expelled from the school. | and-because-of this the-informing were-led the-teachers to the-house-of the-prison and-also were-expelled from-the-house-of the-school
- 1:4.120 1:4.121 1:4.122 1:4.123 1:4.124 1:4.125 1:4.126 1:4.127 1:4.128 1:4.129 1:4.130 1:4.131 1:4.132 1:4.133 | And they brought in their place another teacher, who knew no Hebrew at all and no Arabic as it should be. | and-they-brought in-their-place a-teacher another who not knew not Hebrew on any-creature and-not Arabic as-is-proper
- 1:4.134 1:4.135 1:4.136 1:4.137 1:4.138 1:4.139 1:4.140 1:4.141 1:4.142 1:4.143 1:4.144 1:4.145 | And he had no duty to teach them except the weekly portion, with the Targum of Onkelos and the Targum of Rabbi Saadia Gaon in Arabic. | and-there-is-not upon-him to-teach-them only the-portion-of the-week with the-Targum of-Onkelos and-the-Targum of-Rabbi-Saadia-Gaon in-Arabic
- 1:4.146 1:4.147 1:4.148 1:4.149 1:4.150 1:4.151 1:4.152 1:4.153 1:4.154 1:4.155 | and the Haftarah with the Targum of Jonathan in the Aramaic, with no understanding at all, like the other teachers! | and-the-Haftarah with the-Targum of-Jonathan in-the-Aramaic with-no understanding at-all like-the-rest-of the-teachers
- 1:4.156 1:4.157 1:4.158 1:4.159 1:4.160 1:4.161 | And after a number of years the school closed entirely! | and-after a-number-of ceased the-house-of the-school entirely
- 1:4.162 1:4.163 1:4.164 1:4.165 1:4.166 1:4.167 1:4.168 1:4.169 1:4.170 1:4.171 1:4.172 1:4.173 1:4.174 1:4.175 | All this came upon us because we did not object, for the honour of our Creator, when we were hearing from those who study the Zohar | all this came-upon-us because that not we-objected for-the-honour-of our-Creator at-the-time when-we-were hearing from-those-who-study the-Zohar
- 1:4.176 1:4.177 1:4.178 1:4.179 1:4.180 1:4.181 1:4.182 | and other Kabbalists, things that are forbidden to hear, and we were silent before them | and-the-rest-of the-Kabbalists things that-are-forbidden to-hear and-we-were-silent before-them
- 1:4.183 1:4.184 1:4.185 1:4.186 | as our Rabbis of blessed memory said: "R. A. | as they-said our-Rabbis-of-blessed-memory R.-A.
- 1:4.187 1:4.188 1:4.189 1:4.190 1:4.191 1:4.192 1:4.193 | was seized for heresy, etc.; perhaps a word of heresy, etc." | was-seized for-heresy etc. perhaps a-word of-heresy etc.
- 1:4.194 1:4.195 1:4.196 1:4.197 1:4.198 1:4.199 1:4.200 1:4.201 1:4.202 1:4.203 1:4.204 1:4.205 1:4.206 | and there was fulfilled in us the saying of our Rabbis of blessed memory: "Whoever concedes to them falls into their hands; whoever bathes at their expense, his is theirs." Thus far! | and-was-fulfilled in-us the-saying of-our-Rabbis-of-blessed-memory whoever-concedes to-them falls into-their-hands whoever-bathes at-their-expense his theirs thus-far
- 1:4.207 1:4.208 1:4.209 1:4.210 1:4.211 1:4.212 1:4.213 1:4.214 1:4.215 1:4.216 1:4.217 1:4.218 1:4.219 1:4.220 1:4.221 1:4.222 1:4.223 | Therefore we abandoned our way of studying the languages, Hebrew and Arabic, in their grammar, against our will, because of the wrath of the oppressor, through false informing, | therefore we-abandoned our-way in-the-study-of the-languages Hebrew and-Arabic in-their-grammar against our-will because-of the-wrath-of the-oppressor through the-hand-of informing false
- 1:4.224 1:4.225 1:4.226 1:4.227 1:4.228 1:4.229 1:4.230 | And we turned to the path of holiness, to the honour of our Father who is in heaven, | and-we-turned to the-way-of holiness the-honour-of our-Father who-is-in-heaven
- 1:4.231 1:4.232 1:4.233 1:4.234 1:4.235 1:4.236 1:4.237 | in the nullification of the opinions of the philosopher, the false prophet, the author of the Zohar, | in-the-nullification-of the-opinions-of the-philosopher the-prophet of-falsehood the-author-of the-Zohar
- 1:4.238 1:4.239 1:4.240 1:4.241 1:4.242 1:4.243 1:4.244 1:4.245 1:4.246 1:4.247 1:4.248 1:4.249 1:4.250 1:4.251 | who gathered into his hands the opinions of ancient philosophy, the worshippers of idols, images and hidden powers, which they had invented from their own heart. | who gathered in-his-hands the-opinions-of philosophy ancient the-worshippers-of the-idols images and-powers hidden which invented from-their-heart
- 1:4.252 1:4.253 1:4.254 1:4.255 | And the author of the Zohar seized their opinions, | and-the-author-of the-Zohar seized their-opinions
- 1:4.256 1:4.257 1:4.258 1:4.259 1:4.260 | and thought to bring them into the faith of our holy Torah, | and-thought to-bring-them-in into-the-faith-of our-Torah the-holy
- 1:4.261 1:4.262 1:4.263 1:4.264 1:4.265 1:4.266 1:4.267 | and falsely called his name by the name of Rabbi Shimon bar Yochai the Tanna, | and-called [object-marker] his-name by-the-name-of Rabbi-Shimon-bar-Yochai the-Tanna falsely
- 1:4.268 1:4.269 1:4.270 1:4.271 1:4.272 1:4.273 1:4.274 1:4.275 1:4.276 1:4.277 1:4.278 1:4.279 | and said that the Holy One, blessed be He, who is called by him Atika Kadisha, was revealed to him and revealed to him hidden mysteries. | and-said that-was-revealed to-him the-Holy-One,-blessed-be-He who-is-called by-him Atika Kadisha and-revealed to-him mysteries hidden
- 1:4.280 1:4.281 1:4.282 1:4.283 1:4.284 1:4.285 | And he hung his prophecy and his ideas on Rabbi Shimon bar Yochai the Tanna, peace be upon him. | and-hung his-prophecy and-his-ideas-on Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him
- 1:4.286 1:4.287 1:4.288 1:4.289 | And he made his composition, the Zohar, | and-he-made [object-marker] his-composition the-Zohar
- 1:4.290 1:4.291 1:4.292 1:4.293 1:4.294 1:4.295 1:4.296 1:4.297 1:4.298 | and said that the higher ones agreed to his composition, because this philosopher saw that | and-said that-the-higher-ones agreed to-his-composition because saw the-philosopher this that
- 1:4.299 1:4.300 1:4.301 1:4.302 1:4.303 1:4.304 1:4.305 1:4.306 | the Mishnah, the two Talmuds and the true midrashim of our Rabbis of blessed memory stood against him | the-Mishnah and-the-two Talmuds and-the-midrashim-of our-Rabbis-of-blessed-memory the-true stood against-him
- 1:4.307 1:4.308 1:4.309 | in the belief in His oneness, may He be blessed, | in-the-belief-in His-oneness may-He-be-blessed
- 1:4.310 1:4.311 1:4.312 1:4.313 1:4.314 1:4.315 1:4.316 1:4.317 1:4.318 | therefore he thought to disqualify them, and he despised and mocked them and all who occupy themselves with them, | therefore he-thought to-disqualify-them and-he-despised and-mocked them and-all those-who-occupy themselves-with-them
- 1:4.319 1:4.320 1:4.321 1:4.322 1:4.323 1:4.324 1:4.325 1:4.326 1:4.327 1:4.328 | and he called them by names of contempt and abasement, as the reader will see within! | and-he-called them by-names of contempt and-abasement as will-see the-reader within
- 1:4.329 1:4.330 1:4.331 1:4.332 1:4.333 1:4.334 1:4.335 1:4.336 1:4.337 1:4.338 1:4.339 1:4.340 | And he promised that Israel would not merit redemption until they turned from their way of occupying themselves with the Mishnah and the Talmud. | and-he-promised that-not would-merit Israel redemption until that return from-their-way from-occupying-themselves with-the-Mishnah and-the-Talmud
- 1:4.341 1:4.342 1:4.343 1:4.344 1:4.345 1:4.346 1:4.347 1:4.348 1:4.349 1:4.350 1:4.351 1:4.352 | Only they should occupy themselves with his new Torah, full of idols of images, faces of monstrous form, and various couplings. | only they-should-occupy with-his-Torah the-new the-full idols-of images and-faces of monstrous and-couplings various
- 1:4.353 1:4.354 1:4.355 1:4.356 1:4.357 1:4.358 1:4.359 1:4.360 1:4.361 | And by cunning and deceit they buried it in a chest in the everlasting ground to grow old, and it began to rot, | and-by-cunning and-by-deceit they-buried in-a-chest in-the-ground everlasting to-grow-old and-began to-rot
- 1:4.362 1:4.363 1:4.364 1:4.365 1:4.366 1:4.367 1:4.368 1:4.369 | after many days they published it and brought it out through a king of kings of the idolaters, | after many-days they-published and-brought-out through a-king of-kings the-idolaters
- 1:4.370 1:4.371 1:4.372 1:4.373 1:4.374 1:4.375 1:4.376 1:4.377 | as Rabbi Chida wrote in his book Shem ha-Gedolim, in the section on books, | as-wrote Rabbi Chida in-his-book Shem ha-Gedolim in-the-section on-books
- 1:4.378 1:4.379 1:4.380 1:4.381 1:4.382 1:4.383 | as if it were an ancient composition of Rabbi Shimon bar Yochai the Tanna. | as-if it a-composition ancient from-Rabbi-Shimon-bar-Yochai the-Tanna
- 1:4.384 1:4.385 1:4.386 1:4.387 1:4.388 1:4.389 1:4.390 1:4.391 1:4.392 | And many of our later rabbis were drawn after the aforementioned philosopher with eyes shut, | and-many of-our-rabbis the-later were-drawn after the-philosopher the-aforementioned with-closing of-the-eye
- 1:4.393 1:4.394 1:4.395 1:4.396 1:4.397 1:4.398 | for he took them with the smoothness of his lips and his trimmings, | for he-took them with-smoothness of-his-lips and-his-trimmings
- 1:4.399 1:4.400 1:4.401 1:4.402 | and they did not give heed to his lies! | and-not they-set heart to-his-lies
- 1:4.403 1:4.404 1:4.405 1:4.406 1:4.407 1:4.408 1:4.409 1:4.410 1:4.411 1:4.412 1:4.413 1:4.414 1:4.415 1:4.416 | Heaven forbid that any Jew should believe that Rabbi Shimon bar Yochai the Tanna or another of our Rabbis of blessed memory believed in such beliefs | Heaven-forbid that-anyone son-of Israel to-believe that Rabbi-Shimon-bar-Yochai the-Tanna or another of-our-Rabbis-of-blessed-memory believed in-beliefs such
- 1:4.417 1:4.418 1:4.419 1:4.420 1:4.421 1:4.422 1:4.423 1:4.424 1:4.425 1:4.426 1:4.427 1:4.428 | to exchange and replace the LORD our God, who made known His ways to Moses and His deeds to the children of Israel, | to-exchange and-to-replace [object-marker] the-LORD our-God who made-known His-ways to-Moses to-the-children-of Israel His-deeds
- 1:4.429 1:4.430 1:4.431 1:4.432 1:4.433 1:4.434 1:4.435 1:4.436 1:4.437 | for "He is gracious and compassionate, slow to anger and abounding in kindness," etc. | for gracious and-compassionate He slow to-anger and-abounding in-kindness etc.
- 1:4.438 1:4.439 1:4.440 1:4.441 1:4.442 1:4.443 1:4.444 1:4.445 1:4.446 1:4.447 | for a strange god, quick to anger (Ze'ir Anpin), and to join and associate with Him five partzufim | for-a-god strange quick to-anger Ze'ir-Anpin and-to-conjoin and-to-associate with-Him five partzufim
- 1:4.448 1:4.449 1:4.450 1:4.451 1:4.452 1:4.453 | whose existence has not been established even by proof | whose not has-been-established even their-existence by-proof
- 1:4.454 1:4.455 1:4.456 1:4.457 1:4.458 1:4.459 1:4.460 1:4.461 1:4.462 1:4.463 1:4.464 1:4.465 | (Rabbi Ovadiah Sforno, Deuteronomy 13:3; and there 29:25; and see there 28:14) | Rabbi Ovadiah Sforno Deuteronomy 13 3 and-there 29 25 and-see-there 28 14
- 1:4.466 1:4.467 1:4.468 1:4.469 | and to call them by the name of the LORD our God, | and-to-call-them by-the-name-of the-LORD our-God
- 1:4.470 1:4.471 1:4.472 1:4.473 1:4.474 1:4.475 1:4.476 1:4.477 1:4.478 1:4.479 1:4.480 | and the LORD, the God of truth, whose existence has been established by many proofs, strong as a cast mirror, | and-[object-marker] the-LORD God of-truth who-has-been-established His-existence by-many proofs like-a-mirror of-cast strong
- 1:4.481 1:4.482 1:4.483 1:4.484 1:4.485 1:4.486 1:4.487 1:4.488 1:4.489 1:4.490 | as Rabbi Saadia Gaon wrote in Emunot ve-De'ot, and the Rabbi, the author of Chovot ha-Levavot, | as wrote Rabbi-Saadia-Gaon in-the-book Emunot ve-De'ot and-the-Rabbi the-author-of Chovot ha-Levavot
- 1:4.491 1:4.492 1:4.493 1:4.494 | and the Rambam in the Guide for the Perplexed and in the Yad, | and-the-Rambam in-the-Guide for-the-Perplexed and-in-the-Yad
- 1:4.495 1:4.496 1:4.497 1:4.498 1:4.499 1:4.500 1:4.501 1:4.502 | let us abandon Him and leave Him, and say that He has no name and no appellation, | let-us-abandon and-leave and-say that-there-is-not for-Him a-name and-not an-appellation
- 1:4.503 1:4.504 1:4.505 1:4.506 1:4.507 1:4.508 1:4.509 1:4.510 1:4.511 1:4.512 | and let us worship the partzufim and the forms which, according to his words, were created and descended from Him! | and-let-us-worship [object-marker] the-partzufim and-the-forms which according-to his-words were-created and-descended-in-a-chain from-Him
- 1:4.513 1:4.514 1:4.515 1:4.516 1:4.517 1:4.518 1:4.519 1:4.520 1:4.521 1:4.522 1:4.523 | And we have already expounded at length within, with clear proofs from our Rabbis of blessed memory, on the nullification of these strange opinions of | and-already we-have-expounded-at-length within with-proofs clear from-our-Rabbis-of-blessed-memory on-the-nullification-of the-opinions the-strange these of
- 1:4.524 1:4.525 1:4.526 1:4.527 | the philosopher, the false prophet, the aforementioned! | the-philosopher the-prophet of-falsehood the-aforementioned
- 1:4.528 1:4.529 1:4.530 1:4.531 1:4.532 1:4.533 1:4.534 1:4.535 1:4.536 1:4.537 1:4.538 1:4.539 | for it is the aim of our holy Torah to distance us from belief in idols, whether corporeal or spiritual. | for it is-the-aim-of our-Torah holy to-distance-us from-belief-in idols whether corporeal or spiritual
- 1:4.540 1:4.541 1:4.542 1:4.543 1:4.544 1:4.545 1:4.546 1:4.547 | and to know that the LORD is God; there is no other besides Him. | and-to-know that the-LORD is God there-is-no more besides-Him
- 1:4.548 1:4.549 1:4.550 1:4.551 1:4.552 1:4.553 1:4.554 1:4.555 1:4.556 1:4.557 1:4.558 | and He has neither father nor brother nor son (Rabbah, parashat Yitro), | and-He-has-not for-Him neither father nor brother nor son Rabbah parashat Yitro
- 1:4.559 1:4.560 1:4.561 1:4.562 1:4.563 1:4.564 1:4.565 1:4.566 1:4.567 1:4.568 1:4.569 1:4.570 | and did not receive His kingship from another, since He is not destined to hand it on to another, as our Rabbis of blessed memory explained | and-did-not receive His-kingship from-another since-He-is-not destined to-hand-it-over to-another as explained our-Rabbis of-blessed-memory
- 1:4.571 1:4.572 1:4.573 1:4.574 1:4.575 1:4.576 1:4.577 1:4.578 1:4.579 1:4.580 1:4.581 | (Rabbah, Bamidbar; Shir ha-Shirim, parashah 1; Yerushalmi, Sanhedrin, chapter Dinei Mamonot). | Rabbah Bamidbar Shir ha-Shirim parashah 1 Yerushalmi Sanhedrin chapter Dinei Mamonot
- 1:4.582 1:4.583 1:4.584 1:4.585 1:4.586 1:4.587 1:4.588 1:4.589 1:4.590 | on the verse "I am the first and I am the last, and besides Me there is no God." | on the-verse I-am the-first and-I-am the-last and-besides-Me there-is-no God
- 1:4.591 1:4.592 1:4.593 1:4.594 1:4.595 1:4.596 1:4.597 1:4.598 1:4.599 1:4.600 | and on what Nebuchadnezzar said: "and the appearance of the fourth is like a son of the gods!" | and-on what that-said Nebuchadnezzar and-the-appearance of the-fourth is-like a-son of-gods
- 1:4.601 1:4.602 1:4.603 1:4.604 1:4.605 1:4.606 1:4.607 1:4.608 1:4.609 1:4.610 1:4.611 1:4.612 | And I have seen that Rabbi Menchem di Lonzano, in the book Derekh ha-Chayyim, who opened his mouth wide to speak error | and-I-have-seen that-Rabbi Menchem di Lonzano in-the-book Derekh ha-Chayyim who-widened his-mouth to-speak error
- 1:4.613 1:4.614 1:4.615 1:4.616 1:4.617 1:4.618 | about the holy pietist, the author of Chovot ha-Levavot, | about the-pietist the-holy the-author-of Chovot ha-Levavot
- 1:4.619 1:4.620 1:4.621 1:4.622 1:4.623 1:4.624 1:4.625 1:4.626 | concerning his saying, in the name of the philosopher, that one could not worship | concerning his-saying in-the-name-of the-philosopher that-said that-not could worship
- 1:4.627 1:4.628 1:4.629 1:4.630 1:4.631 1:4.632 | the Cause of causes, except a prophet of his generation by nature, | the-cause-of the-causes except a-prophet of-his-generation by-nature
- 1:4.633 1:4.634 1:4.635 1:4.636 1:4.637 1:4.638 1:4.639 | or a philosopher outstanding in what he acquired from wisdom; | or a-philosopher outstanding in-what he-acquired from wisdom
- 1:4.640 1:4.641 1:4.642 1:4.643 1:4.644 1:4.645 1:4.646 1:4.647 1:4.648 1:4.649 | but others worship other than Him, because they do not understand: the result is only composite! | but others worship other-than-Him because they-are-not understanding the-result-is only composite
- 1:4.650 1:4.651 1:4.652 1:4.653 1:4.654 1:4.655 1:4.656 1:4.657 1:4.658 1:4.659 1:4.660 | And at these words of the pietist he took offence and thought that here was ground for complaint against him, | and-at the-words-of the-pietist these took-offence and-thought that here a-ground of-complaint against-him
- 1:4.661 1:4.662 1:4.663 1:4.664 1:4.665 1:4.666 | and found a pretext to mislead the listeners and alarm them by saying: | and-a-pretext found to-mislead the-listeners and-to-alarm-them by-saying
- 1:4.667 1:4.668 1:4.669 1:4.670 1:4.671 1:4.672 1:4.673 1:4.674 | "Is it that such-and-such and such-and-such of our later famous sages worship other than Him?" | is-it-that such-and-such and-such-and-such of-our-sages the-later the-famous worship other-than-Him
- 1:4.675 1:4.676 1:4.677 1:4.678 1:4.679 1:4.680 | And he counted them, in his language, one by one by name. | and-he counted-them in-his-language one one by-name
- 1:4.681 1:4.682 1:4.683 1:4.684 1:4.685 1:4.686 1:4.687 1:4.688 | and he did not take to heart that many of the sages had grown fat, | and-not he-placed to his-heart that many of-the-sages had-grown-fat
- 1:4.689 1:4.690 1:4.691 1:4.692 1:4.693 1:4.694 1:4.695 | went along with eyes shut and believed the words of the author of the Zohar | went-along with-closing of-eyes and-believed the-words-of the-author-of the-Zohar
- 1:4.696 1:4.697 1:4.698 1:4.699 1:4.700 1:4.701 1:4.702 | who called himself by the name of Rabbi Shimon bar Yochai the Tanna, | who called [object-marker] himself by-the-name-of Rabbi-Shimon-bar-Yochai the-Tanna
- 1:4.703 1:4.704 1:4.705 1:4.706 1:4.707 1:4.708 1:4.709 | in order to attach his false prophecies to the name of the Tanna, | in-order-to attach his-prophecies false to-the-name-of of the-Tanna
- 1:4.710 1:4.711 1:4.712 1:4.713 1:4.714 1:4.715 1:4.716 1:4.717 1:4.718 1:4.719 1:4.720 | and spoke slander against the LORD our God, who was revealed to him in his study house as Atika Kadisha, | and-spoke slander against the-LORD our-God who-was-revealed to-him in-the-house-of his-study Atika Kadisha
- 1:4.721 1:4.722 1:4.723 1:4.724 1:4.725 1:4.726 | and revealed to him hidden secrets and hidden powers | and-revealed to-him secrets hidden and-powers hidden
- 1:4.727 1:4.728 1:4.729 | (Zohar, Bereshit 22). | Zohar Bereshit 22
- 1:4.730 1:4.731 1:4.732 1:4.733 1:4.734 1:4.735 1:4.736 1:4.737 1:4.738 1:4.739 1:4.740 | for the Creator is the idol called Abba, who used to command his consort, the she-idol Imma, | for the-Creator is the-idol called Abba who-was commanding his-consort the-she-idol Imma
- 1:4.741 1:4.742 1:4.743 1:4.744 1:4.745 1:4.746 | and she does all that he says to her, | and-she does all that he-says to-her
- 1:4.747 1:4.748 1:4.749 1:4.750 1:4.751 1:4.752 1:4.753 1:4.754 1:4.755 1:4.756 1:4.757 1:4.758 1:4.759 1:4.760 | except the creation of the first man, which she, Imma, said and did, not with the will of Abba! | except for-the-creation-of man the-first which-she Imma said and-did [object-marker] man the-first not with-the-will of-Abba
- 1:4.761 1:4.762 1:4.763 1:4.764 1:4.765 1:4.766 1:4.767 1:4.768 1:4.769 1:4.770 1:4.771 1:4.772 | And in Parashat Vayera, folio 112, and likewise in Parashat Balak, folio 191, he ascribes our worship and our prayers | and-in-Parashat Vayera folio 112 and-likewise in-Parashat Balak folio 191 ascribes our-worship and-our-prayers
- 1:4.773 1:4.774 1:4.775 1:4.776 1:4.777 1:4.778 | to a strange god, quick to anger (Ze'ir Anpin), | to-a-god strange quick to-anger Ze'ir Anpin
- 1:4.779 1:4.780 1:4.781 1:4.782 1:4.783 1:4.784 1:4.785 | who did not make the heavens and the earth. | who not made [object-marker] the-heavens and-[object-marker] the-earth
- 1:4.786 1:4.787 1:4.788 1:4.789 1:4.790 1:4.791 1:4.792 | And he calls these idols by the name of the LORD our God. | and-he-calls he to-the-idols these by-the-name-of the-LORD our-God
- 1:4.793 1:4.794 1:4.795 1:4.796 1:4.797 1:4.798 1:4.799 | And about Ze'ir Anpin, Jeremiah the prophet, peace be upon him, proclaimed | and-about Ze'ir Anpin proclaimed Jeremiah the-prophet peace-be-upon-him
- 1:4.800 1:4.801 1:4.802 1:4.803 1:4.804 1:4.805 1:4.806 1:4.807 1:4.808 1:4.809 1:4.810 1:4.811 1:4.812 1:4.813 1:4.814 1:4.815 | "Thus shall you say to them: the gods that did not make the heavens and the earth shall perish from the earth and from under these heavens." | thus shall-you-say to-them the-gods that the-heavens and-the-earth not made shall-perish from the-earth and-from under the-heavens these
- 1:4.816 1:4.817 1:4.818 1:4.819 1:4.820 1:4.821 1:4.822 1:4.823 1:4.824 | for it is clear in many places in the Zohar that He is created and not creator | for-indeed it-is-clear in-the-Zohar in-many places that-He is-created and-not creator
- 1:4.825 1:4.826 1:4.827 1:4.828 1:4.829 1:4.830 1:4.831 1:4.832 1:4.833 | (see Zohar, Bereshit 22; and Yosher Levav, folio 3b). | see Zohar Bereshit 22 and-Yosher Levav folio 3 b
- 1:4.834 1:4.835 1:4.836 1:4.837 1:4.838 1:4.839 1:4.840 1:4.841 1:4.842 1:4.843 | And yet he says to worship him, and to bless him, and to accept him as God, and to fear his consort! | and-yet he says to-worship-him and-to-bless-him and-to-accept-him upon-him as-God and-to-fear his-consort
- 1:4.844 1:4.845 1:4.846 1:4.847 1:4.848 1:4.849 1:4.850 1:4.851 1:4.852 1:4.853 1:4.854 1:4.855 1:4.856 | And many of our later rabbis, whom he mentioned in the book Derekh ha-Chayyim, in their amazement at the pietist, the author of Chovot ha-Levavot, | and-many of-our-rabbis the-later who-mentioned in-the-book Derekh ha-Chayyim in-their-amazement at the-pietist the-author-of Chovot ha-Levavot
- 1:4.857 1:4.858 1:4.859 1:4.860 1:4.861 1:4.862 1:4.863 1:4.864 1:4.865 1:4.866 1:4.867 1:4.868 1:4.869 1:4.870 | went along with eyes shut and were drawn after him and believed his belief and accepted upon themselves the idol, quick to anger (Ze'ir Anpin), as God! | went with-closing of-the-eye and-were-drawn after-him and-believed in-his-belief and-accepted the-idol quick to-anger Ze'ir-Anpin upon-them as-God
- 1:4.871 1:4.872 1:4.873 1:4.874 1:4.875 1:4.876 1:4.877 1:4.878 1:4.879 1:4.880 1:4.881 1:4.882 1:4.883 1:4.884 1:4.885 1:4.886 | And all this came upon them because they did not contemplate the wisdom of the Creator of all, which is to investigate and examine the works of Creation as explained | and-all this caused to-them because that-not they-contemplated the-wisdom-of the-Creator of-all which-is to-investigate and-examine the-works-of Creation as-explained
- 1:4.887 1:4.888 1:4.889 1:4.890 1:4.891 1:4.892 1:4.893 1:4.894 1:4.895 | in the words of the Rabbi, the author of Moreh Nevukhim, and in the words of the pietist, the author of Chovot ha-Levavot. | in-the-words-of the-Rabbi the-author-of Moreh-Nevukhim and-in-the-words-of the-pietist the-author-of Chovot ha-Levavot
- 1:4.896 1:4.897 1:4.898 1:4.899 1:4.900 1:4.901 | And our master Rabbi Saadia Gaon and Rabbi Yehudah Halevi in the Kuzari, | and-our-master Saadia Gaon and-Rabbi-Yehudah Halevi in-the-Kuzari
- 1:4.902 1:4.903 1:4.904 1:4.905 1:4.906 1:4.907 1:4.908 1:4.909 1:4.910 | as the prophet, peace be upon him, said: "They do not regard the deed of the LORD." | as said the-prophet peace-be-upon-him and-[object-marker] the-deed of-the-LORD not they-regard
- 1:4.911 1:4.912 1:4.913 1:4.914 1:4.915 1:4.916 1:4.917 1:4.918 1:4.919 1:4.920 | for they thought their words were Greek wisdom, and they kept far from it, as far as a bow-shot, | for they-thought their-words to-be-wisdom Greek and-they-kept-far from-it far as-far-as a-bow
- 1:4.921 1:4.922 1:4.923 1:4.924 1:4.925 1:4.926 1:4.927 1:4.928 1:4.929 1:4.930 | and they also did not take to heart the true words of our forefathers of blessed memory in the midrashim | and-also not they-set to their-heart the-words-of our-forefathers of-blessed-memory true in-the-midrashim
- 1:4.931 1:4.932 1:4.933 1:4.934 1:4.935 1:4.936 1:4.937 1:4.938 1:4.939 1:4.940 | that Abraham our father did not recognise his Creator except by contemplating the work of the LORD, | that-not recognised Abraham our-father [object-marker] his-Creator except by-contemplating the-work of-the-LORD
- 1:4.941 1:4.942 1:4.943 1:4.944 1:4.945 1:4.946 1:4.947 1:4.948 1:4.949 | and investigated and expounded them and brought proofs from them until he attained the truth; | and-investigated and-expounded them and-brought proofs from-them until he-attained the-truth
- 1:4.950 1:4.951 1:4.952 1:4.953 1:4.954 1:4.955 1:4.956 1:4.957 1:4.958 1:4.959 1:4.960 | then the Holy One, blessed be He, was revealed to him and said to him: "You love righteousness and hate wickedness," etc. | then was-revealed to-him the-Holy-One,-blessed-be-He and-said to-him You-love righteousness and-hate wickedness etc.
- 1:4.961 1:4.962 1:4.963 1:4.964 1:4.965 1:4.966 1:4.967 1:4.968 1:4.969 1:4.970 1:4.971 1:4.972 1:4.973 1:4.974 | A parable: to one walking on the road who saw a castle, etc. (Midrash ha-Gadol, at the beginning of Parashat Lekh Lekha). | a-parable to-one who-was walking on-the-road and-saw a-castle etc. Midrash ha-Gadol the-beginning-of parashat Lekh Lekha
- 1:4.975 1:4.976 1:4.977 1:4.978 1:4.979 1:4.980 1:4.981 1:4.982 1:4.983 1:4.984 | And perforce we must acknowledge the truth that the words of the pietist, the author of Chovot ha-Levavot, | and perforce we-acknowledge to the-truth that-the-words-of the-pietist the-author-of Chovot ha-Levavot
- 1:4.985 1:4.986 1:4.987 1:4.988 1:4.989 1:4.990 1:4.991 1:4.992 1:4.993 1:4.994 | clearly say that one cannot worship the God of truth except, etc. | clearly they-said that-not one-can worship [object-marker] God of-truth except etc.
- 1:4.995 1:4.996 1:4.997 1:4.998 1:4.999 1:4.1000 1:4.1001 1:4.1002 1:4.1003 1:4.1004 1:4.1005 1:4.1006 | but others worship other than Him, as the prophet said: "For many days Israel was without a true God," | but others worship other-than-Him as-the-saying-of the-prophet for-many days Israel without God of-truth
- 1:4.1007 1:4.1008 1:4.1009 1:4.1010 1:4.1011 1:4.1012 1:4.1013 1:4.1014 1:4.1015 1:4.1016 | and as our eyes see and our ears hear, to the point that a man can say to his fellow! | and-as our-eyes see and-our-ears hear to-the-point that-it-is-possible that-one-should-say a-man to-his-fellow
- 1:4.1017 1:4.1018 1:4.1019 1:4.1020 1:4.1021 | "Your people is my people, and your God is not my God!" | your-people my-people and-your-God not my-God
- 1:4.1022 1:4.1023 1:4.1024 1:4.1025 1:4.1026 1:4.1027 1:4.1028 1:4.1029 1:4.1030 1:4.1031 1:4.1032 1:4.1033 1:4.1034 | And to the God of truth let us pour out our prayer, to guide us on paths of uprightness, and to deliver us from every strange belief. Amen, may it be His will. | and-to the-God-of truth let-us-pour-out our-prayer to-guide-us in-the-paths-of uprightness and-to-deliver-us from-every belief strange Amen,-may-it-be-His-will

Grammar notes:
- Voice: 1:4 is in the author's own voice; 'we' (עלינו, עזבנו, נודה) is the author and the circle of the Yemenite teachers. The long periods are divided into English sentences, and no clause is dropped.
- Aramaic quotations: 1:4.198-206 (דמודי להון) is Bavli Aramaic; 1:4.274-275 (עתיקא קדישא) and 1:4.278 (רזין טמירין) are Zohar Aramaic; 1:4.591-600 is Biblical Aramaic (Dan 3:25) and 1:4.800-815 is Biblical Aramaic (Jer 10:11).
- Allusions with no quotation formula are translated from the verse and take no marks: 'שקר בימין' (1:4.61-62, Isa 44:20), 'בחלק שפתיו' (1:4.396-397, Prov 7:21), and 'ולדעת כי ה' הוא האלהים אין עוד מלבדו' (1:4.540-547, Deut 4:35).
- 1:4.667 'וכי' opens a rhetorical question, kept as a question. 1:4.688 'שמנה' is pointed as 'grew fat' (cf. Deut 32:15), and the English follows the pointing.

Translator's note: The paragraph is the author's own voice, a continuous prose period. The pack has no morphology, lexicon or commentary, so the reasons for the choices are in the token notes and the apparatus.

Commentary rulings:
- apparatus: Base text: Wikisource transcription. Q2 additions are translated as base text and are not yet collated with the 1931 scan.
- apparatus: 1:4.62 'a lie in the right hand' alludes to Isa 44:20.
- apparatus: 1:4.396 'the smoothness of his lips' alludes to Prov 7:21.
- apparatus: 1:4.433-434 'slow to anger' is Exod 34:6, the allusion behind 'ארך אפים ורב חסד'; the author's 'quick to anger' (1:4.440-441, 1:4.866-867) is set against it.
- apparatus: 1:4.478-480 'strong as a cast mirror' is Job 37:18 (Hebrew 'strong like a cast mirror').
- apparatus: 1:4.454-465 references as the author gives them: Sforno on Deut 13:3, 29:25, 28:14.
- apparatus: 1:4.540-547 Deut 4:35.
- apparatus: 1:4.584-590 Isa 44:6: 'I am the first and I am the last, and besides Me there is no God'.
- apparatus: 1:4.591-600 Dan 3:25 (Biblical Aramaic).
- apparatus: 1:4.800-815 Jer 10:11 (Biblical Aramaic); 'gods' renders Aramaic elahaya.
- apparatus: 1:4.910 Isa 5:12, quoted by the author.
- apparatus: 1:4.1001-1006 2 Chr 15:3, quoted by the author as 'the prophet'.
- apparatus: 1:4.1017-1021 Ruth 1:16 with the negation the author gives ('your God is not my God'); the English follows his wording.
- apparatus: 1:4.107 'H.N.Ts.' is not expanded; the expansion is not certain.
- apparatus: 1:4.150 'ב״ע' read as 'in the Aramaic'; the abbreviation is not certain, and no other reading is established.
- apparatus: 1:4.186 'R. A.' is the author's abbreviation and is left as written.
- apparatus: 1:4.1034 'אכי״ר' expanded as 'may it be His will' (אמן כן יהי רצון).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

