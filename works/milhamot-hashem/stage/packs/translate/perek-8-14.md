# Translation pack: milhamot-hashem perek-8-14 (Milchamot ha-Shem 8:17–8:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-14`.

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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
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
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
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


## Milchamot ha-Shem 8:17

Hebrew (Wikisource): קיא) ואלה הדעות הן הנה הדעות הרעות שהיו לאבותינו בימי השופטים ובימי ירבעם בן נבט ושאר מלכי ישראל וקצת מלכי יהודה בזמן בית ראשון אשר עליהם נתנבאו כל הנביאים, וקראו בגרון הוי החוקקים חקקי און, ומכתבי עמל כתבו, והיה להם דבר ה' צו לצו לצו קו לקו קו זעיר שם זעיר וכמאמר רז"ל אין צו אלא ע"ז שנאמר כי הואיל הלך אחרי צו וכו' וכן קו לקו הם הפרצופים בעלי שלשה קוים המוזכרים בדברי המקובלים כמו שכתב מהרח"ו בעץ חיים שער עגולים ויושר, שספירות היושר הם בציור שלשה קוים, כמראה האדם בעל ראש, וזרועות ושוקים וגוף ורגלים וכו', וכן זעיר שם זעיר שם הוא זעיר אנפין, (קצר אפים) ונאמר ויעזבו את ה' אלהיהם ויעבדו את הבעלים ואת העשתרות, הבעלים הם האלוהות הזכרים, אבא וזעיר אנפין, ונקראו בעלים על שם הזווגים והבעילות שמזדווגין עם הנקיבות, כמו שאמרו רז"ל בירושלמי שבת פרק אמר ר' עקיבא ובמסכת אבות ע"ז פרק כל הצלמים וז"ל אמר ר' הונא רבי חמא בשם רב, הבעל ראש גוייה היה לו ונאפון הוא, וישימו להם בעל ברית לאלהים, עכ"ל הירושלמי הרי מבואר בהדיא מדברי הירוש' שנק' הצלם הזכר מעל על שם שיש לו ראש גויה והוא בועל ומנאף, וידוע הוא כי האלילים הקדמונים עשו להם צלמים מכוונים נגד הצורות והפרצופים שחשבו שיש בשמים, כדי להמשיך השפע מהצורות העליונות להם, כי הם עשוים בצלמם ודמותם, וגם כן נקראו בעלים על שם שעובדיהם שמו אותם אדונים להם כמו שכתב הרד"ק וכלי יקר. ואת העשתרות הם נקבות האלוהות שהם אימא ומלכית, כי עשתרות צאנך הם נקבות הצאן כמו שכתבו המפרשים, וכתן נאמר ואת העשתרות בלי סמיכות מלת הצאן, והם נקיבות האלוהות הנזכרים. וגם בין הכשרים היו האמונות האלו פושטים ולכן אמר נבוכדנאצר הרשע בעת שראה מלאך ה' בכבשן האש שירד להציל לחנניה מיאל ועזריה מן האור, וריויה די רביעאה דמי לבר אלהין, לפי שחשב שהזעיר אנפין הנקרא בן לאבא ואימא ירד להצילם מן האש, כדגרסינן בירושלמי שבת פרק במה אשה, וריוה די רביעאה דמי לבא אלהין אמר רבי ראובן באותה שעה ירד מלאך וסטרו לההוא רשיעא על פיו ואמר ליה תקין מילך, ובר אית ליה? חזר ואמר בריך אלההון די שדרך מישך ועביד נגו, די שלח מלאכיה ושיזיב לעבדוהי, בריה אין כתיב כאן אלא מלאכיה, ע"כ. ומזה יתברר לך שהדעות האלו של מחבר הזהר והמקובלים החדשים, זרים המה לדעת תורתינו הקדושה, ומהעמים אשר השאיר ה' לנסות בם את ישראל באו לנו, ונטו ישראל אחריהם בימי שפוט השופטים, וגם מן הכשדים שגלו ביניהם בעת חרבן בית ראשון.

| id | word | punctuation after |
|---|---|---|
| 8:17.1 | קיא | ) |
| 8:17.2 | ואלה |  |
| 8:17.3 | הדעות |  |
| 8:17.4 | הן |  |
| 8:17.5 | הנה |  |
| 8:17.6 | הדעות |  |
| 8:17.7 | הרעות |  |
| 8:17.8 | שהיו |  |
| 8:17.9 | לאבותינו |  |
| 8:17.10 | בימי |  |
| 8:17.11 | השופטים |  |
| 8:17.12 | ובימי |  |
| 8:17.13 | ירבעם |  |
| 8:17.14 | בן |  |
| 8:17.15 | נבט |  |
| 8:17.16 | ושאר |  |
| 8:17.17 | מלכי |  |
| 8:17.18 | ישראל |  |
| 8:17.19 | וקצת |  |
| 8:17.20 | מלכי |  |
| 8:17.21 | יהודה |  |
| 8:17.22 | בזמן |  |
| 8:17.23 | בית |  |
| 8:17.24 | ראשון |  |
| 8:17.25 | אשר |  |
| 8:17.26 | עליהם |  |
| 8:17.27 | נתנבאו |  |
| 8:17.28 | כל |  |
| 8:17.29 | הנביאים | , |
| 8:17.30 | וקראו |  |
| 8:17.31 | בגרון |  |
| 8:17.32 | הוי |  |
| 8:17.33 | החוקקים |  |
| 8:17.34 | חקקי |  |
| 8:17.35 | און | , |
| 8:17.36 | ומכתבי |  |
| 8:17.37 | עמל |  |
| 8:17.38 | כתבו | , |
| 8:17.39 | והיה |  |
| 8:17.40 | להם |  |
| 8:17.41 | דבר |  |
| 8:17.42 | ה | ' |
| 8:17.43 | צו |  |
| 8:17.44 | לצו |  |
| 8:17.45 | לצו |  |
| 8:17.46 | קו |  |
| 8:17.47 | לקו |  |
| 8:17.48 | קו |  |
| 8:17.49 | זעיר |  |
| 8:17.50 | שם |  |
| 8:17.51 | זעיר |  |
| 8:17.52 | וכמאמר |  |
| 8:17.53 | רז"ל |  |
| 8:17.54 | אין |  |
| 8:17.55 | צו |  |
| 8:17.56 | אלא |  |
| 8:17.57 | ע"ז |  |
| 8:17.58 | שנאמר |  |
| 8:17.59 | כי |  |
| 8:17.60 | הואיל |  |
| 8:17.61 | הלך |  |
| 8:17.62 | אחרי |  |
| 8:17.63 | צו |  |
| 8:17.64 | וכו | ' |
| 8:17.65 | וכן |  |
| 8:17.66 | קו |  |
| 8:17.67 | לקו |  |
| 8:17.68 | הם |  |
| 8:17.69 | הפרצופים |  |
| 8:17.70 | בעלי |  |
| 8:17.71 | שלשה |  |
| 8:17.72 | קוים |  |
| 8:17.73 | המוזכרים |  |
| 8:17.74 | בדברי |  |
| 8:17.75 | המקובלים |  |
| 8:17.76 | כמו |  |
| 8:17.77 | שכתב |  |
| 8:17.78 | מהרח"ו |  |
| 8:17.79 | בעץ |  |
| 8:17.80 | חיים |  |
| 8:17.81 | שער |  |
| 8:17.82 | עגולים |  |
| 8:17.83 | ויושר | , |
| 8:17.84 | שספירות |  |
| 8:17.85 | היושר |  |
| 8:17.86 | הם |  |
| 8:17.87 | בציור |  |
| 8:17.88 | שלשה |  |
| 8:17.89 | קוים | , |
| 8:17.90 | כמראה |  |
| 8:17.91 | האדם |  |
| 8:17.92 | בעל |  |
| 8:17.93 | ראש | , |
| 8:17.94 | וזרועות |  |
| 8:17.95 | ושוקים |  |
| 8:17.96 | וגוף |  |
| 8:17.97 | ורגלים |  |
| 8:17.98 | וכו | ', |
| 8:17.99 | וכן |  |
| 8:17.100 | זעיר |  |
| 8:17.101 | שם |  |
| 8:17.102 | זעיר |  |
| 8:17.103 | שם |  |
| 8:17.104 | הוא |  |
| 8:17.105 | זעיר |  |
| 8:17.106 | אנפין | , |
| 8:17.107 | קצר |  |
| 8:17.108 | אפים | ) |
| 8:17.109 | ונאמר |  |
| 8:17.110 | ויעזבו |  |
| 8:17.111 | את |  |
| 8:17.112 | ה | ' |
| 8:17.113 | אלהיהם |  |
| 8:17.114 | ויעבדו |  |
| 8:17.115 | את |  |
| 8:17.116 | הבעלים |  |
| 8:17.117 | ואת |  |
| 8:17.118 | העשתרות | , |
| 8:17.119 | הבעלים |  |
| 8:17.120 | הם |  |
| 8:17.121 | האלוהות |  |
| 8:17.122 | הזכרים | , |
| 8:17.123 | אבא |  |
| 8:17.124 | וזעיר |  |
| 8:17.125 | אנפין | , |
| 8:17.126 | ונקראו |  |
| 8:17.127 | בעלים |  |
| 8:17.128 | על |  |
| 8:17.129 | שם |  |
| 8:17.130 | הזווגים |  |
| 8:17.131 | והבעילות |  |
| 8:17.132 | שמזדווגין |  |
| 8:17.133 | עם |  |
| 8:17.134 | הנקיבות | , |
| 8:17.135 | כמו |  |
| 8:17.136 | שאמרו |  |
| 8:17.137 | רז"ל |  |
| 8:17.138 | בירושלמי |  |
| 8:17.139 | שבת |  |
| 8:17.140 | פרק |  |
| 8:17.141 | אמר |  |
| 8:17.142 | ר | ' |
| 8:17.143 | עקיבא |  |
| 8:17.144 | ובמסכת |  |
| 8:17.145 | אבות |  |
| 8:17.146 | ע"ז |  |
| 8:17.147 | פרק |  |
| 8:17.148 | כל |  |
| 8:17.149 | הצלמים |  |
| 8:17.150 | וז"ל |  |
| 8:17.151 | אמר |  |
| 8:17.152 | ר | ' |
| 8:17.153 | הונא |  |
| 8:17.154 | רבי |  |
| 8:17.155 | חמא |  |
| 8:17.156 | בשם |  |
| 8:17.157 | רב | , |
| 8:17.158 | הבעל |  |
| 8:17.159 | ראש |  |
| 8:17.160 | גוייה |  |
| 8:17.161 | היה |  |
| 8:17.162 | לו |  |
| 8:17.163 | ונאפון |  |
| 8:17.164 | הוא | , |
| 8:17.165 | וישימו |  |
| 8:17.166 | להם |  |
| 8:17.167 | בעל |  |
| 8:17.168 | ברית |  |
| 8:17.169 | לאלהים | , |
| 8:17.170 | עכ"ל |  |
| 8:17.171 | הירושלמי |  |
| 8:17.172 | הרי |  |
| 8:17.173 | מבואר |  |
| 8:17.174 | בהדיא |  |
| 8:17.175 | מדברי |  |
| 8:17.176 | הירוש | ' |
| 8:17.177 | שנק | ' |
| 8:17.178 | הצלם |  |
| 8:17.179 | הזכר |  |
| 8:17.180 | מעל |  |
| 8:17.181 | על |  |
| 8:17.182 | שם |  |
| 8:17.183 | שיש |  |
| 8:17.184 | לו |  |
| 8:17.185 | ראש |  |
| 8:17.186 | גויה |  |
| 8:17.187 | והוא |  |
| 8:17.188 | בועל |  |
| 8:17.189 | ומנאף | , |
| 8:17.190 | וידוע |  |
| 8:17.191 | הוא |  |
| 8:17.192 | כי |  |
| 8:17.193 | האלילים |  |
| 8:17.194 | הקדמונים |  |
| 8:17.195 | עשו |  |
| 8:17.196 | להם |  |
| 8:17.197 | צלמים |  |
| 8:17.198 | מכוונים |  |
| 8:17.199 | נגד |  |
| 8:17.200 | הצורות |  |
| 8:17.201 | והפרצופים |  |
| 8:17.202 | שחשבו |  |
| 8:17.203 | שיש |  |
| 8:17.204 | בשמים | , |
| 8:17.205 | כדי |  |
| 8:17.206 | להמשיך |  |
| 8:17.207 | השפע |  |
| 8:17.208 | מהצורות |  |
| 8:17.209 | העליונות |  |
| 8:17.210 | להם | , |
| 8:17.211 | כי |  |
| 8:17.212 | הם |  |
| 8:17.213 | עשוים |  |
| 8:17.214 | בצלמם |  |
| 8:17.215 | ודמותם | , |
| 8:17.216 | וגם |  |
| 8:17.217 | כן |  |
| 8:17.218 | נקראו |  |
| 8:17.219 | בעלים |  |
| 8:17.220 | על |  |
| 8:17.221 | שם |  |
| 8:17.222 | שעובדיהם |  |
| 8:17.223 | שמו |  |
| 8:17.224 | אותם |  |
| 8:17.225 | אדונים |  |
| 8:17.226 | להם |  |
| 8:17.227 | כמו |  |
| 8:17.228 | שכתב |  |
| 8:17.229 | הרד"ק |  |
| 8:17.230 | וכלי |  |
| 8:17.231 | יקר | . |
| 8:17.232 | ואת |  |
| 8:17.233 | העשתרות |  |
| 8:17.234 | הם |  |
| 8:17.235 | נקבות |  |
| 8:17.236 | האלוהות |  |
| 8:17.237 | שהם |  |
| 8:17.238 | אימא |  |
| 8:17.239 | ומלכית | , |
| 8:17.240 | כי |  |
| 8:17.241 | עשתרות |  |
| 8:17.242 | צאנך |  |
| 8:17.243 | הם |  |
| 8:17.244 | נקבות |  |
| 8:17.245 | הצאן |  |
| 8:17.246 | כמו |  |
| 8:17.247 | שכתבו |  |
| 8:17.248 | המפרשים | , |
| 8:17.249 | וכתן |  |
| 8:17.250 | נאמר |  |
| 8:17.251 | ואת |  |
| 8:17.252 | העשתרות |  |
| 8:17.253 | בלי |  |
| 8:17.254 | סמיכות |  |
| 8:17.255 | מלת |  |
| 8:17.256 | הצאן | , |
| 8:17.257 | והם |  |
| 8:17.258 | נקיבות |  |
| 8:17.259 | האלוהות |  |
| 8:17.260 | הנזכרים | . |
| 8:17.261 | וגם |  |
| 8:17.262 | בין |  |
| 8:17.263 | הכשרים |  |
| 8:17.264 | היו |  |
| 8:17.265 | האמונות |  |
| 8:17.266 | האלו |  |
| 8:17.267 | פושטים |  |
| 8:17.268 | ולכן |  |
| 8:17.269 | אמר |  |
| 8:17.270 | נבוכדנאצר |  |
| 8:17.271 | הרשע |  |
| 8:17.272 | בעת |  |
| 8:17.273 | שראה |  |
| 8:17.274 | מלאך |  |
| 8:17.275 | ה | ' |
| 8:17.276 | בכבשן |  |
| 8:17.277 | האש |  |
| 8:17.278 | שירד |  |
| 8:17.279 | להציל |  |
| 8:17.280 | לחנניה |  |
| 8:17.281 | מיאל |  |
| 8:17.282 | ועזריה |  |
| 8:17.283 | מן |  |
| 8:17.284 | האור | , |
| 8:17.285 | וריויה |  |
| 8:17.286 | די |  |
| 8:17.287 | רביעאה |  |
| 8:17.288 | דמי |  |
| 8:17.289 | לבר |  |
| 8:17.290 | אלהין | , |
| 8:17.291 | לפי |  |
| 8:17.292 | שחשב |  |
| 8:17.293 | שהזעיר |  |
| 8:17.294 | אנפין |  |
| 8:17.295 | הנקרא |  |
| 8:17.296 | בן |  |
| 8:17.297 | לאבא |  |
| 8:17.298 | ואימא |  |
| 8:17.299 | ירד |  |
| 8:17.300 | להצילם |  |
| 8:17.301 | מן |  |
| 8:17.302 | האש | , |
| 8:17.303 | כדגרסינן |  |
| 8:17.304 | בירושלמי |  |
| 8:17.305 | שבת |  |
| 8:17.306 | פרק |  |
| 8:17.307 | במה |  |
| 8:17.308 | אשה | , |
| 8:17.309 | וריוה |  |
| 8:17.310 | די |  |
| 8:17.311 | רביעאה |  |
| 8:17.312 | דמי |  |
| 8:17.313 | לבא |  |
| 8:17.314 | אלהין |  |
| 8:17.315 | אמר |  |
| 8:17.316 | רבי |  |
| 8:17.317 | ראובן |  |
| 8:17.318 | באותה |  |
| 8:17.319 | שעה |  |
| 8:17.320 | ירד |  |
| 8:17.321 | מלאך |  |
| 8:17.322 | וסטרו |  |
| 8:17.323 | לההוא |  |
| 8:17.324 | רשיעא |  |
| 8:17.325 | על |  |
| 8:17.326 | פיו |  |
| 8:17.327 | ואמר |  |
| 8:17.328 | ליה |  |
| 8:17.329 | תקין |  |
| 8:17.330 | מילך | , |
| 8:17.331 | ובר |  |
| 8:17.332 | אית |  |
| 8:17.333 | ליה | ? |
| 8:17.334 | חזר |  |
| 8:17.335 | ואמר |  |
| 8:17.336 | בריך |  |
| 8:17.337 | אלההון |  |
| 8:17.338 | די |  |
| 8:17.339 | שדרך |  |
| 8:17.340 | מישך |  |
| 8:17.341 | ועביד |  |
| 8:17.342 | נגו | , |
| 8:17.343 | די |  |
| 8:17.344 | שלח |  |
| 8:17.345 | מלאכיה |  |
| 8:17.346 | ושיזיב |  |
| 8:17.347 | לעבדוהי | , |
| 8:17.348 | בריה |  |
| 8:17.349 | אין |  |
| 8:17.350 | כתיב |  |
| 8:17.351 | כאן |  |
| 8:17.352 | אלא |  |
| 8:17.353 | מלאכיה | , |
| 8:17.354 | ע"כ | . |
| 8:17.355 | ומזה |  |
| 8:17.356 | יתברר |  |
| 8:17.357 | לך |  |
| 8:17.358 | שהדעות |  |
| 8:17.359 | האלו |  |
| 8:17.360 | של |  |
| 8:17.361 | מחבר |  |
| 8:17.362 | הזהר |  |
| 8:17.363 | והמקובלים |  |
| 8:17.364 | החדשים | , |
| 8:17.365 | זרים |  |
| 8:17.366 | המה |  |
| 8:17.367 | לדעת |  |
| 8:17.368 | תורתינו |  |
| 8:17.369 | הקדושה | , |
| 8:17.370 | ומהעמים |  |
| 8:17.371 | אשר |  |
| 8:17.372 | השאיר |  |
| 8:17.373 | ה | ' |
| 8:17.374 | לנסות |  |
| 8:17.375 | בם |  |
| 8:17.376 | את |  |
| 8:17.377 | ישראל |  |
| 8:17.378 | באו |  |
| 8:17.379 | לנו | , |
| 8:17.380 | ונטו |  |
| 8:17.381 | ישראל |  |
| 8:17.382 | אחריהם |  |
| 8:17.383 | בימי |  |
| 8:17.384 | שפוט |  |
| 8:17.385 | השופטים | , |
| 8:17.386 | וגם |  |
| 8:17.387 | מן |  |
| 8:17.388 | הכשדים |  |
| 8:17.389 | שגלו |  |
| 8:17.390 | ביניהם |  |
| 8:17.391 | בעת |  |
| 8:17.392 | חרבן |  |
| 8:17.393 | בית |  |
| 8:17.394 | ראשון | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

