# Translation pack: milhamot-hashem perek-5-22 (Milchamot ha-Shem 5:31–5:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-22`.

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
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- דברי תימה: "astonishing words" (words that astonish).
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- אלפין: "alefs" (the letter).
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 5:31

Hebrew (Wikisource): נז) יבין המבין וישכיל היטב בדבר רבותינו ז"ל הנז' פה כי רעיונות הפילסוף מחבר הזהר והמקובלים הנמשכים אחריו מכוונים מאד למחשבות דורו של אנוש שקראו לעצבים בשמו של הקב"ה והתפללו אליהם, והמקובלים גם הם עשו כן, כי קראו לכל אחד מן הפרצופים בשמו של הקב"ה ואמרו לעבדן ולהשתחות להם ולקרוא אליהם בעת צרה וכל התפלות והברכות שתקנו אנשי כנה"ג לשבח ולפאר לשי"ת בהן. את הכל הסבו לפרצוף ז"א בצירוף שאר פרצופים הללו אע"פ שהן נבראים. וחשבו לתקן הכפירה הנגלית ונראית באמונה זו באמרם שהם עובדים לנשמתם. ושכחו את האמונה האמתית המפורסמת ופשוטה בכל כתבי הקדש תורה שבכתב ותורה שבעל פה. כי אלהינו ב"ה אינו גוף ולא כח בגוף. והם המציאו ועשו לאלהינו כמה גופים ומלבושים מתוארים בשרד ובמחוגה אורך ורוחב להם. ונתנו להם נשמות מהאלהים העליון. והשאירו עיקרו לעילא לעילא מבלי להתפלל אליו, כי לא השאירו לו שם. ותלו כל מצותיה של תורה וספוריה המלמדים לאדם דעת בגופא דמלכא (ולא בנשמתו) מנהון בידוי מנהון ברגלוי וכו' כמו שכ' מהר"י צאהרי בשם הזהר, וסותרים ג"כ דבריהם שכתבתי לעיל סי' ט"ל בן שושן סודות ומשנת חסידים והרמ"ז דבין הגוף של ספירות בין הנשמה בין המלבוש הכל אלהות גמור. כי כולם נמשכים ונהוים מעצמות האין סוף כהדין קמצא דלבושיה מניה וביה. ומשתלשלים זה מזה כבן מן האב. ולפיכך קורין לז"א בן לאבא ואימא. והוא השולט על כל הברואים וזנם ומפרנסם במצות אביו ואמו כדאיתא בזהר בלק דף קצ"א (ח"ג קצא, א). והובא לעיל סי' כ"ו. ושם מבואר שהעבודה והברכות שייכים (לפי דעתם) לז"א דהכי פקידו אבא ואימא דכלא יפלחון ליה וכמו שכ' ג"כ ס' הברית וכס"א ויושר לבב ונחלת יוסף. הנוטע אזן הלא ישמע. ויוצר עין לאדם ולב להבין ולהשכיל שכן היא מחשבת דורו של אנוש וטעותם אשר חשבו את הצורות ואת הפרצופים אשר בדו מלבם לאלוהות וקראו להם בשמותיו של הקב"ה ועבדום להיות אמצעיים להמשיך השפע מהאלהים העליון. ואם תאמר שאני מעשי דורו של אנוש שעשו פרצופים וצורות מעשה ידי אדם וקראום בשמו של הקב"ה ועבדום. אבל המקובלים לא עשו שום צורה בםועל רק חשבו הפרצופים והצורות שהן מעשה השם לאלוהות? הלא גם דורו של אנוש כן עלה בדעתם בתחלה כמו שכ' הרמב"ם ז"ל בפ"א מהל' ע"ש. והרי כבר אמרו רז"ל יכול לא יעשה אבל יעלה בדעתו. ת"ל לא יהיה לך אלהים אחרים מכל מקום. וכן הוא אומר ופן תשא עיניך השמימה וראית את השמש ואת הירח ואת הכוכבים כל צבא השמים ונדחת. הכל בכלל, שלא יעלה בדעתו אלהות וממשלה לשום נברא או נאצל. כי כולם בשם צבא השמים שהזהיר הכתוב מלעבדן. ואף אם קורא להן בשמותיו של הקב"ה כמו שעשו דורו של אנוש אין בהן ממש, כמו שמבואר בילקוט ומדרש הגדול וברש"י שהבאתי לעיל, סי' נ"ו, כי לא ברא הקב"ה דבר מבריותיו להיות נעבד, כי אם להנהיג את העולם כחפצו ורצונו, והמה כעבדים העומדים לפני המלך לעשות רצונו שנ' אלף אלפין ישמשוניה ורבוא רבוון קדמוהי יקומון, לא להיות אלוהות ח"ו ולא לעבדן. כי אין אלוה מבלעדי ה' ואין צור כאלהינו. והוא קרוב לכל קוראיו באמת, ואף אם ברגע א', יקראו אליו כל יצוריו ויגיד כל אחד מהם, ועונה אותו וממלא משאלו שלכל אחד ואחד אם הוא ראוי לפי מעשיו, כי עונותיו של אדם מעכבים מלעשות בקשתו, שנ' לתת לאיש כדרכיו, כפרי מעלליו, ונאמר כי אם עונותיכם היו מבדילים ביניכם לבין אלהיכם וכמו שכ' הרמב"ם בהל' תשובה.

| id | word | punctuation after |
|---|---|---|
| 5:31.1 | נז | ) |
| 5:31.2 | יבין |  |
| 5:31.3 | המבין |  |
| 5:31.4 | וישכיל |  |
| 5:31.5 | היטב |  |
| 5:31.6 | בדבר |  |
| 5:31.7 | רבותינו |  |
| 5:31.8 | ז"ל |  |
| 5:31.9 | הנז | ' |
| 5:31.10 | פה |  |
| 5:31.11 | כי |  |
| 5:31.12 | רעיונות |  |
| 5:31.13 | הפילסוף |  |
| 5:31.14 | מחבר |  |
| 5:31.15 | הזהר |  |
| 5:31.16 | והמקובלים |  |
| 5:31.17 | הנמשכים |  |
| 5:31.18 | אחריו |  |
| 5:31.19 | מכוונים |  |
| 5:31.20 | מאד |  |
| 5:31.21 | למחשבות |  |
| 5:31.22 | דורו |  |
| 5:31.23 | של |  |
| 5:31.24 | אנוש |  |
| 5:31.25 | שקראו |  |
| 5:31.26 | לעצבים |  |
| 5:31.27 | בשמו |  |
| 5:31.28 | של |  |
| 5:31.29 | הקב"ה |  |
| 5:31.30 | והתפללו |  |
| 5:31.31 | אליהם | , |
| 5:31.32 | והמקובלים |  |
| 5:31.33 | גם |  |
| 5:31.34 | הם |  |
| 5:31.35 | עשו |  |
| 5:31.36 | כן | , |
| 5:31.37 | כי |  |
| 5:31.38 | קראו |  |
| 5:31.39 | לכל |  |
| 5:31.40 | אחד |  |
| 5:31.41 | מן |  |
| 5:31.42 | הפרצופים |  |
| 5:31.43 | בשמו |  |
| 5:31.44 | של |  |
| 5:31.45 | הקב"ה |  |
| 5:31.46 | ואמרו |  |
| 5:31.47 | לעבדן |  |
| 5:31.48 | ולהשתחות |  |
| 5:31.49 | להם |  |
| 5:31.50 | ולקרוא |  |
| 5:31.51 | אליהם |  |
| 5:31.52 | בעת |  |
| 5:31.53 | צרה |  |
| 5:31.54 | וכל |  |
| 5:31.55 | התפלות |  |
| 5:31.56 | והברכות |  |
| 5:31.57 | שתקנו |  |
| 5:31.58 | אנשי |  |
| 5:31.59 | כנה"ג |  |
| 5:31.60 | לשבח |  |
| 5:31.61 | ולפאר |  |
| 5:31.62 | לשי"ת |  |
| 5:31.63 | בהן | . |
| 5:31.64 | את |  |
| 5:31.65 | הכל |  |
| 5:31.66 | הסבו |  |
| 5:31.67 | לפרצוף |  |
| 5:31.68 | ז"א |  |
| 5:31.69 | בצירוף |  |
| 5:31.70 | שאר |  |
| 5:31.71 | פרצופים |  |
| 5:31.72 | הללו |  |
| 5:31.73 | אע"פ |  |
| 5:31.74 | שהן |  |
| 5:31.75 | נבראים | . |
| 5:31.76 | וחשבו |  |
| 5:31.77 | לתקן |  |
| 5:31.78 | הכפירה |  |
| 5:31.79 | הנגלית |  |
| 5:31.80 | ונראית |  |
| 5:31.81 | באמונה |  |
| 5:31.82 | זו |  |
| 5:31.83 | באמרם |  |
| 5:31.84 | שהם |  |
| 5:31.85 | עובדים |  |
| 5:31.86 | לנשמתם | . |
| 5:31.87 | ושכחו |  |
| 5:31.88 | את |  |
| 5:31.89 | האמונה |  |
| 5:31.90 | האמתית |  |
| 5:31.91 | המפורסמת |  |
| 5:31.92 | ופשוטה |  |
| 5:31.93 | בכל |  |
| 5:31.94 | כתבי |  |
| 5:31.95 | הקדש |  |
| 5:31.96 | תורה |  |
| 5:31.97 | שבכתב |  |
| 5:31.98 | ותורה |  |
| 5:31.99 | שבעל |  |
| 5:31.100 | פה | . |
| 5:31.101 | כי |  |
| 5:31.102 | אלהינו |  |
| 5:31.103 | ב"ה |  |
| 5:31.104 | אינו |  |
| 5:31.105 | גוף |  |
| 5:31.106 | ולא |  |
| 5:31.107 | כח |  |
| 5:31.108 | בגוף | . |
| 5:31.109 | והם |  |
| 5:31.110 | המציאו |  |
| 5:31.111 | ועשו |  |
| 5:31.112 | לאלהינו |  |
| 5:31.113 | כמה |  |
| 5:31.114 | גופים |  |
| 5:31.115 | ומלבושים |  |
| 5:31.116 | מתוארים |  |
| 5:31.117 | בשרד |  |
| 5:31.118 | ובמחוגה |  |
| 5:31.119 | אורך |  |
| 5:31.120 | ורוחב |  |
| 5:31.121 | להם | . |
| 5:31.122 | ונתנו |  |
| 5:31.123 | להם |  |
| 5:31.124 | נשמות |  |
| 5:31.125 | מהאלהים |  |
| 5:31.126 | העליון | . |
| 5:31.127 | והשאירו |  |
| 5:31.128 | עיקרו |  |
| 5:31.129 | לעילא |  |
| 5:31.130 | לעילא |  |
| 5:31.131 | מבלי |  |
| 5:31.132 | להתפלל |  |
| 5:31.133 | אליו | , |
| 5:31.134 | כי |  |
| 5:31.135 | לא |  |
| 5:31.136 | השאירו |  |
| 5:31.137 | לו |  |
| 5:31.138 | שם | . |
| 5:31.139 | ותלו |  |
| 5:31.140 | כל |  |
| 5:31.141 | מצותיה |  |
| 5:31.142 | של |  |
| 5:31.143 | תורה |  |
| 5:31.144 | וספוריה |  |
| 5:31.145 | המלמדים |  |
| 5:31.146 | לאדם |  |
| 5:31.147 | דעת |  |
| 5:31.148 | בגופא |  |
| 5:31.149 | דמלכא |  |
| 5:31.150 | ולא |  |
| 5:31.151 | בנשמתו | ) |
| 5:31.152 | מנהון |  |
| 5:31.153 | בידוי |  |
| 5:31.154 | מנהון |  |
| 5:31.155 | ברגלוי |  |
| 5:31.156 | וכו | ' |
| 5:31.157 | כמו |  |
| 5:31.158 | שכ | ' |
| 5:31.159 | מהר"י |  |
| 5:31.160 | צאהרי |  |
| 5:31.161 | בשם |  |
| 5:31.162 | הזהר | , |
| 5:31.163 | וסותרים |  |
| 5:31.164 | ג"כ |  |
| 5:31.165 | דבריהם |  |
| 5:31.166 | שכתבתי |  |
| 5:31.167 | לעיל |  |
| 5:31.168 | סי | ' |
| 5:31.169 | ט"ל |  |
| 5:31.170 | בן |  |
| 5:31.171 | שושן |  |
| 5:31.172 | סודות |  |
| 5:31.173 | ומשנת |  |
| 5:31.174 | חסידים |  |
| 5:31.175 | והרמ"ז |  |
| 5:31.176 | דבין |  |
| 5:31.177 | הגוף |  |
| 5:31.178 | של |  |
| 5:31.179 | ספירות |  |
| 5:31.180 | בין |  |
| 5:31.181 | הנשמה |  |
| 5:31.182 | בין |  |
| 5:31.183 | המלבוש |  |
| 5:31.184 | הכל |  |
| 5:31.185 | אלהות |  |
| 5:31.186 | גמור | . |
| 5:31.187 | כי |  |
| 5:31.188 | כולם |  |
| 5:31.189 | נמשכים |  |
| 5:31.190 | ונהוים |  |
| 5:31.191 | מעצמות |  |
| 5:31.192 | האין |  |
| 5:31.193 | סוף |  |
| 5:31.194 | כהדין |  |
| 5:31.195 | קמצא |  |
| 5:31.196 | דלבושיה |  |
| 5:31.197 | מניה |  |
| 5:31.198 | וביה | . |
| 5:31.199 | ומשתלשלים |  |
| 5:31.200 | זה |  |
| 5:31.201 | מזה |  |
| 5:31.202 | כבן |  |
| 5:31.203 | מן |  |
| 5:31.204 | האב | . |
| 5:31.205 | ולפיכך |  |
| 5:31.206 | קורין |  |
| 5:31.207 | לז"א |  |
| 5:31.208 | בן |  |
| 5:31.209 | לאבא |  |
| 5:31.210 | ואימא | . |
| 5:31.211 | והוא |  |
| 5:31.212 | השולט |  |
| 5:31.213 | על |  |
| 5:31.214 | כל |  |
| 5:31.215 | הברואים |  |
| 5:31.216 | וזנם |  |
| 5:31.217 | ומפרנסם |  |
| 5:31.218 | במצות |  |
| 5:31.219 | אביו |  |
| 5:31.220 | ואמו |  |
| 5:31.221 | כדאיתא |  |
| 5:31.222 | בזהר |  |
| 5:31.223 | בלק |  |
| 5:31.224 | דף |  |
| 5:31.225 | קצ"א |  |
| 5:31.226 | ח"ג |  |
| 5:31.227 | קצא | , |
| 5:31.228 | א | ). |
| 5:31.229 | והובא |  |
| 5:31.230 | לעיל |  |
| 5:31.231 | סי | ' |
| 5:31.232 | כ"ו | . |
| 5:31.233 | ושם |  |
| 5:31.234 | מבואר |  |
| 5:31.235 | שהעבודה |  |
| 5:31.236 | והברכות |  |
| 5:31.237 | שייכים |  |
| 5:31.238 | לפי |  |
| 5:31.239 | דעתם | ) |
| 5:31.240 | לז"א |  |
| 5:31.241 | דהכי |  |
| 5:31.242 | פקידו |  |
| 5:31.243 | אבא |  |
| 5:31.244 | ואימא |  |
| 5:31.245 | דכלא |  |
| 5:31.246 | יפלחון |  |
| 5:31.247 | ליה |  |
| 5:31.248 | וכמו |  |
| 5:31.249 | שכ | ' |
| 5:31.250 | ג"כ |  |
| 5:31.251 | ס | ' |
| 5:31.252 | הברית |  |
| 5:31.253 | וכס"א |  |
| 5:31.254 | ויושר |  |
| 5:31.255 | לבב |  |
| 5:31.256 | ונחלת |  |
| 5:31.257 | יוסף | . |
| 5:31.258 | הנוטע |  |
| 5:31.259 | אזן |  |
| 5:31.260 | הלא |  |
| 5:31.261 | ישמע | . |
| 5:31.262 | ויוצר |  |
| 5:31.263 | עין |  |
| 5:31.264 | לאדם |  |
| 5:31.265 | ולב |  |
| 5:31.266 | להבין |  |
| 5:31.267 | ולהשכיל |  |
| 5:31.268 | שכן |  |
| 5:31.269 | היא |  |
| 5:31.270 | מחשבת |  |
| 5:31.271 | דורו |  |
| 5:31.272 | של |  |
| 5:31.273 | אנוש |  |
| 5:31.274 | וטעותם |  |
| 5:31.275 | אשר |  |
| 5:31.276 | חשבו |  |
| 5:31.277 | את |  |
| 5:31.278 | הצורות |  |
| 5:31.279 | ואת |  |
| 5:31.280 | הפרצופים |  |
| 5:31.281 | אשר |  |
| 5:31.282 | בדו |  |
| 5:31.283 | מלבם |  |
| 5:31.284 | לאלוהות |  |
| 5:31.285 | וקראו |  |
| 5:31.286 | להם |  |
| 5:31.287 | בשמותיו |  |
| 5:31.288 | של |  |
| 5:31.289 | הקב"ה |  |
| 5:31.290 | ועבדום |  |
| 5:31.291 | להיות |  |
| 5:31.292 | אמצעיים |  |
| 5:31.293 | להמשיך |  |
| 5:31.294 | השפע |  |
| 5:31.295 | מהאלהים |  |
| 5:31.296 | העליון | . |
| 5:31.297 | ואם |  |
| 5:31.298 | תאמר |  |
| 5:31.299 | שאני |  |
| 5:31.300 | מעשי |  |
| 5:31.301 | דורו |  |
| 5:31.302 | של |  |
| 5:31.303 | אנוש |  |
| 5:31.304 | שעשו |  |
| 5:31.305 | פרצופים |  |
| 5:31.306 | וצורות |  |
| 5:31.307 | מעשה |  |
| 5:31.308 | ידי |  |
| 5:31.309 | אדם |  |
| 5:31.310 | וקראום |  |
| 5:31.311 | בשמו |  |
| 5:31.312 | של |  |
| 5:31.313 | הקב"ה |  |
| 5:31.314 | ועבדום | . |
| 5:31.315 | אבל |  |
| 5:31.316 | המקובלים |  |
| 5:31.317 | לא |  |
| 5:31.318 | עשו |  |
| 5:31.319 | שום |  |
| 5:31.320 | צורה |  |
| 5:31.321 | בםועל |  |
| 5:31.322 | רק |  |
| 5:31.323 | חשבו |  |
| 5:31.324 | הפרצופים |  |
| 5:31.325 | והצורות |  |
| 5:31.326 | שהן |  |
| 5:31.327 | מעשה |  |
| 5:31.328 | השם |  |
| 5:31.329 | לאלוהות | ? |
| 5:31.330 | הלא |  |
| 5:31.331 | גם |  |
| 5:31.332 | דורו |  |
| 5:31.333 | של |  |
| 5:31.334 | אנוש |  |
| 5:31.335 | כן |  |
| 5:31.336 | עלה |  |
| 5:31.337 | בדעתם |  |
| 5:31.338 | בתחלה |  |
| 5:31.339 | כמו |  |
| 5:31.340 | שכ | ' |
| 5:31.341 | הרמב"ם |  |
| 5:31.342 | ז"ל |  |
| 5:31.343 | בפ"א |  |
| 5:31.344 | מהל | ' |
| 5:31.345 | ע"ש | . |
| 5:31.346 | והרי |  |
| 5:31.347 | כבר |  |
| 5:31.348 | אמרו |  |
| 5:31.349 | רז"ל |  |
| 5:31.350 | יכול |  |
| 5:31.351 | לא |  |
| 5:31.352 | יעשה |  |
| 5:31.353 | אבל |  |
| 5:31.354 | יעלה |  |
| 5:31.355 | בדעתו | . |
| 5:31.356 | ת"ל |  |
| 5:31.357 | לא |  |
| 5:31.358 | יהיה |  |
| 5:31.359 | לך |  |
| 5:31.360 | אלהים |  |
| 5:31.361 | אחרים |  |
| 5:31.362 | מכל |  |
| 5:31.363 | מקום | . |
| 5:31.364 | וכן |  |
| 5:31.365 | הוא |  |
| 5:31.366 | אומר |  |
| 5:31.367 | ופן |  |
| 5:31.368 | תשא |  |
| 5:31.369 | עיניך |  |
| 5:31.370 | השמימה |  |
| 5:31.371 | וראית |  |
| 5:31.372 | את |  |
| 5:31.373 | השמש |  |
| 5:31.374 | ואת |  |
| 5:31.375 | הירח |  |
| 5:31.376 | ואת |  |
| 5:31.377 | הכוכבים |  |
| 5:31.378 | כל |  |
| 5:31.379 | צבא |  |
| 5:31.380 | השמים |  |
| 5:31.381 | ונדחת | . |
| 5:31.382 | הכל |  |
| 5:31.383 | בכלל | , |
| 5:31.384 | שלא |  |
| 5:31.385 | יעלה |  |
| 5:31.386 | בדעתו |  |
| 5:31.387 | אלהות |  |
| 5:31.388 | וממשלה |  |
| 5:31.389 | לשום |  |
| 5:31.390 | נברא |  |
| 5:31.391 | או |  |
| 5:31.392 | נאצל | . |
| 5:31.393 | כי |  |
| 5:31.394 | כולם |  |
| 5:31.395 | בשם |  |
| 5:31.396 | צבא |  |
| 5:31.397 | השמים |  |
| 5:31.398 | שהזהיר |  |
| 5:31.399 | הכתוב |  |
| 5:31.400 | מלעבדן | . |
| 5:31.401 | ואף |  |
| 5:31.402 | אם |  |
| 5:31.403 | קורא |  |
| 5:31.404 | להן |  |
| 5:31.405 | בשמותיו |  |
| 5:31.406 | של |  |
| 5:31.407 | הקב"ה |  |
| 5:31.408 | כמו |  |
| 5:31.409 | שעשו |  |
| 5:31.410 | דורו |  |
| 5:31.411 | של |  |
| 5:31.412 | אנוש |  |
| 5:31.413 | אין |  |
| 5:31.414 | בהן |  |
| 5:31.415 | ממש | , |
| 5:31.416 | כמו |  |
| 5:31.417 | שמבואר |  |
| 5:31.418 | בילקוט |  |
| 5:31.419 | ומדרש |  |
| 5:31.420 | הגדול |  |
| 5:31.421 | וברש"י |  |
| 5:31.422 | שהבאתי |  |
| 5:31.423 | לעיל | , |
| 5:31.424 | סי | ' |
| 5:31.425 | נ"ו | , |
| 5:31.426 | כי |  |
| 5:31.427 | לא |  |
| 5:31.428 | ברא |  |
| 5:31.429 | הקב"ה |  |
| 5:31.430 | דבר |  |
| 5:31.431 | מבריותיו |  |
| 5:31.432 | להיות |  |
| 5:31.433 | נעבד | , |
| 5:31.434 | כי |  |
| 5:31.435 | אם |  |
| 5:31.436 | להנהיג |  |
| 5:31.437 | את |  |
| 5:31.438 | העולם |  |
| 5:31.439 | כחפצו |  |
| 5:31.440 | ורצונו | , |
| 5:31.441 | והמה |  |
| 5:31.442 | כעבדים |  |
| 5:31.443 | העומדים |  |
| 5:31.444 | לפני |  |
| 5:31.445 | המלך |  |
| 5:31.446 | לעשות |  |
| 5:31.447 | רצונו |  |
| 5:31.448 | שנ | ' |
| 5:31.449 | אלף |  |
| 5:31.450 | אלפין |  |
| 5:31.451 | ישמשוניה |  |
| 5:31.452 | ורבוא |  |
| 5:31.453 | רבוון |  |
| 5:31.454 | קדמוהי |  |
| 5:31.455 | יקומון | , |
| 5:31.456 | לא |  |
| 5:31.457 | להיות |  |
| 5:31.458 | אלוהות |  |
| 5:31.459 | ח"ו |  |
| 5:31.460 | ולא |  |
| 5:31.461 | לעבדן | . |
| 5:31.462 | כי |  |
| 5:31.463 | אין |  |
| 5:31.464 | אלוה |  |
| 5:31.465 | מבלעדי |  |
| 5:31.466 | ה | ' |
| 5:31.467 | ואין |  |
| 5:31.468 | צור |  |
| 5:31.469 | כאלהינו | . |
| 5:31.470 | והוא |  |
| 5:31.471 | קרוב |  |
| 5:31.472 | לכל |  |
| 5:31.473 | קוראיו |  |
| 5:31.474 | באמת | , |
| 5:31.475 | ואף |  |
| 5:31.476 | אם |  |
| 5:31.477 | ברגע |  |
| 5:31.478 | א | ', |
| 5:31.479 | יקראו |  |
| 5:31.480 | אליו |  |
| 5:31.481 | כל |  |
| 5:31.482 | יצוריו |  |
| 5:31.483 | ויגיד |  |
| 5:31.484 | כל |  |
| 5:31.485 | אחד |  |
| 5:31.486 | מהם | , |
| 5:31.487 | ועונה |  |
| 5:31.488 | אותו |  |
| 5:31.489 | וממלא |  |
| 5:31.490 | משאלו |  |
| 5:31.491 | שלכל |  |
| 5:31.492 | אחד |  |
| 5:31.493 | ואחד |  |
| 5:31.494 | אם |  |
| 5:31.495 | הוא |  |
| 5:31.496 | ראוי |  |
| 5:31.497 | לפי |  |
| 5:31.498 | מעשיו | , |
| 5:31.499 | כי |  |
| 5:31.500 | עונותיו |  |
| 5:31.501 | של |  |
| 5:31.502 | אדם |  |
| 5:31.503 | מעכבים |  |
| 5:31.504 | מלעשות |  |
| 5:31.505 | בקשתו | , |
| 5:31.506 | שנ | ' |
| 5:31.507 | לתת |  |
| 5:31.508 | לאיש |  |
| 5:31.509 | כדרכיו | , |
| 5:31.510 | כפרי |  |
| 5:31.511 | מעלליו | , |
| 5:31.512 | ונאמר |  |
| 5:31.513 | כי |  |
| 5:31.514 | אם |  |
| 5:31.515 | עונותיכם |  |
| 5:31.516 | היו |  |
| 5:31.517 | מבדילים |  |
| 5:31.518 | ביניכם |  |
| 5:31.519 | לבין |  |
| 5:31.520 | אלהיכם |  |
| 5:31.521 | וכמו |  |
| 5:31.522 | שכ | ' |
| 5:31.523 | הרמב"ם |  |
| 5:31.524 | בהל | ' |
| 5:31.525 | תשובה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

