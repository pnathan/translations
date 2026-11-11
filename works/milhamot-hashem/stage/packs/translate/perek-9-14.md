# Translation pack: milhamot-hashem perek-9-14 (Milchamot ha-Shem 9:17–9:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-14`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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
- והוא שלא: "provided that" (fixed).
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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
- The words of a verse quoted in the plural (Job 32:16, עמדו): plural, no
  subject added.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 9:17

Hebrew (Wikisource): קלה) רב לכם נשיאי ישראל מכל דברי מינות האמורים בזהר בכמה מקומות, ובפרט בפרשת בהר דף ק"כ שהאלהים העליון שאם ראוי לעבדו (לפי דבריה) ברא והמציא את האלהים התחתונים הרבים, והוא ממלא אותם בחללם וחוצה להם, והוא המאחד ומקשר את נקיבותיהם עם הזכרים לקרא להם בשם קב"ה ושכינתיה, ואותם נעבוד, ונירא את נקיבותיהם, ודוקא לפרצוף אחרון שבהם הנקרא קצר אפים (ז"א) עם מלכת השמים המזדווגת עמו, משום דלא שייך עבודה ותפלה רק לפרצופים שכבר נתגשמו ונתעבו קצת כמבואר במה שקדם מדברי כסא אליהו יושר לבב וספר הברית ועץ חיים להרח"ו שער טנת"ה סוף פרק ב' ובשער הנקודים פ"א שכל מה שהאו יורד למטה הוא מתעבה יותר וניכר ונרגש ומתגלה יותר, אבל באין סוף הוא הסבה הראשונה ואפי' בפרצופים ראשונים (שהמציאו המקובלים החדשים לאלהינו) כאדם קדמאה ואדם קדמון ועתיק ואריך שעדיין לא נתעבו ונתגשמו כל כך לא שייכא עבודה וקריאה, כי אין המה, ולא יאות לקרוא להם בשם אלהים כמפורש יוצא מספר עץ חיים להרח"ו שער הנקודים פ"ז ובהג"ה שם. שזהו האלהים העליון עשה את האדם ישר (והמה בקשו חשבונות רבים) כי האדם קדמון היה בו בחינת יושר וזהו עשה את האדם ישר (הגהה א"מ ןקרא המאציל אלהים בדרך שאלה. כי הרי אפי' באריך אנפין אין שם אלהים) ע"כ. הרי לך בהדיא דפרצופים העליונים לא שייכא בהו אלהות וקריאה כלל, כי לא אלהים המה ורק להעלולים התחתונים נקרא ונעבוד, והמה אלהינו (לפי דעתם) כמו שכתב הרב יושר לבב וספר הברית וכס"א ונח"י ומצרף האמונה! ולפי דבריהם שאומרים דלית רשותא לבר נש למיעל קמי מריה אלא בידיעת אמונה זו, איך אמרו רז"ל רצה הקב"ה לזכות את ישראל לפיכך הרבה להם תורה ומצות שנ' ה' חפץ למען צדקו יגדיל תורה ויאדיר! ולמה זה לא רצה הקב"ה לזכות אותנו ולהודענו על ידי נביאנו הרבים שעמדו להם לישראל ממשה רבינו ע"ה עד שהנביא האחרון מלאכי שאמר זכרו תורת משה עבדי וכו'. את אמונה זו שהעבודה היא להעלולים התחתונים. ושהעובד לאלהים עליון לא יענהו ואדרבא ענוש יענש ומחה ה' את שמו כמו שכתב ספר הברית?

| id | word | punctuation after |
|---|---|---|
| 9:17.1 | קלה | ) |
| 9:17.2 | רב |  |
| 9:17.3 | לכם |  |
| 9:17.4 | נשיאי |  |
| 9:17.5 | ישראל |  |
| 9:17.6 | מכל |  |
| 9:17.7 | דברי |  |
| 9:17.8 | מינות |  |
| 9:17.9 | האמורים |  |
| 9:17.10 | בזהר |  |
| 9:17.11 | בכמה |  |
| 9:17.12 | מקומות | , |
| 9:17.13 | ובפרט |  |
| 9:17.14 | בפרשת |  |
| 9:17.15 | בהר |  |
| 9:17.16 | דף |  |
| 9:17.17 | ק"כ |  |
| 9:17.18 | שהאלהים |  |
| 9:17.19 | העליון |  |
| 9:17.20 | שאם |  |
| 9:17.21 | ראוי |  |
| 9:17.22 | לעבדו |  |
| 9:17.23 | לפי |  |
| 9:17.24 | דבריה | ) |
| 9:17.25 | ברא |  |
| 9:17.26 | והמציא |  |
| 9:17.27 | את |  |
| 9:17.28 | האלהים |  |
| 9:17.29 | התחתונים |  |
| 9:17.30 | הרבים | , |
| 9:17.31 | והוא |  |
| 9:17.32 | ממלא |  |
| 9:17.33 | אותם |  |
| 9:17.34 | בחללם |  |
| 9:17.35 | וחוצה |  |
| 9:17.36 | להם | , |
| 9:17.37 | והוא |  |
| 9:17.38 | המאחד |  |
| 9:17.39 | ומקשר |  |
| 9:17.40 | את |  |
| 9:17.41 | נקיבותיהם |  |
| 9:17.42 | עם |  |
| 9:17.43 | הזכרים |  |
| 9:17.44 | לקרא |  |
| 9:17.45 | להם |  |
| 9:17.46 | בשם |  |
| 9:17.47 | קב"ה |  |
| 9:17.48 | ושכינתיה | , |
| 9:17.49 | ואותם |  |
| 9:17.50 | נעבוד | , |
| 9:17.51 | ונירא |  |
| 9:17.52 | את |  |
| 9:17.53 | נקיבותיהם | , |
| 9:17.54 | ודוקא |  |
| 9:17.55 | לפרצוף |  |
| 9:17.56 | אחרון |  |
| 9:17.57 | שבהם |  |
| 9:17.58 | הנקרא |  |
| 9:17.59 | קצר |  |
| 9:17.60 | אפים |  |
| 9:17.61 | ז"א | ) |
| 9:17.62 | עם |  |
| 9:17.63 | מלכת |  |
| 9:17.64 | השמים |  |
| 9:17.65 | המזדווגת |  |
| 9:17.66 | עמו | , |
| 9:17.67 | משום |  |
| 9:17.68 | דלא |  |
| 9:17.69 | שייך |  |
| 9:17.70 | עבודה |  |
| 9:17.71 | ותפלה |  |
| 9:17.72 | רק |  |
| 9:17.73 | לפרצופים |  |
| 9:17.74 | שכבר |  |
| 9:17.75 | נתגשמו |  |
| 9:17.76 | ונתעבו |  |
| 9:17.77 | קצת |  |
| 9:17.78 | כמבואר |  |
| 9:17.79 | במה |  |
| 9:17.80 | שקדם |  |
| 9:17.81 | מדברי |  |
| 9:17.82 | כסא |  |
| 9:17.83 | אליהו |  |
| 9:17.84 | יושר |  |
| 9:17.85 | לבב |  |
| 9:17.86 | וספר |  |
| 9:17.87 | הברית |  |
| 9:17.88 | ועץ |  |
| 9:17.89 | חיים |  |
| 9:17.90 | להרח"ו |  |
| 9:17.91 | שער |  |
| 9:17.92 | טנת"ה |  |
| 9:17.93 | סוף |  |
| 9:17.94 | פרק |  |
| 9:17.95 | ב | ' |
| 9:17.96 | ובשער |  |
| 9:17.97 | הנקודים |  |
| 9:17.98 | פ"א |  |
| 9:17.99 | שכל |  |
| 9:17.100 | מה |  |
| 9:17.101 | שהאו |  |
| 9:17.102 | יורד |  |
| 9:17.103 | למטה |  |
| 9:17.104 | הוא |  |
| 9:17.105 | מתעבה |  |
| 9:17.106 | יותר |  |
| 9:17.107 | וניכר |  |
| 9:17.108 | ונרגש |  |
| 9:17.109 | ומתגלה |  |
| 9:17.110 | יותר | , |
| 9:17.111 | אבל |  |
| 9:17.112 | באין |  |
| 9:17.113 | סוף |  |
| 9:17.114 | הוא |  |
| 9:17.115 | הסבה |  |
| 9:17.116 | הראשונה |  |
| 9:17.117 | ואפי | ' |
| 9:17.118 | בפרצופים |  |
| 9:17.119 | ראשונים |  |
| 9:17.120 | שהמציאו |  |
| 9:17.121 | המקובלים |  |
| 9:17.122 | החדשים |  |
| 9:17.123 | לאלהינו | ) |
| 9:17.124 | כאדם |  |
| 9:17.125 | קדמאה |  |
| 9:17.126 | ואדם |  |
| 9:17.127 | קדמון |  |
| 9:17.128 | ועתיק |  |
| 9:17.129 | ואריך |  |
| 9:17.130 | שעדיין |  |
| 9:17.131 | לא |  |
| 9:17.132 | נתעבו |  |
| 9:17.133 | ונתגשמו |  |
| 9:17.134 | כל |  |
| 9:17.135 | כך |  |
| 9:17.136 | לא |  |
| 9:17.137 | שייכא |  |
| 9:17.138 | עבודה |  |
| 9:17.139 | וקריאה | , |
| 9:17.140 | כי |  |
| 9:17.141 | אין |  |
| 9:17.142 | המה | , |
| 9:17.143 | ולא |  |
| 9:17.144 | יאות |  |
| 9:17.145 | לקרוא |  |
| 9:17.146 | להם |  |
| 9:17.147 | בשם |  |
| 9:17.148 | אלהים |  |
| 9:17.149 | כמפורש |  |
| 9:17.150 | יוצא |  |
| 9:17.151 | מספר |  |
| 9:17.152 | עץ |  |
| 9:17.153 | חיים |  |
| 9:17.154 | להרח"ו |  |
| 9:17.155 | שער |  |
| 9:17.156 | הנקודים |  |
| 9:17.157 | פ"ז |  |
| 9:17.158 | ובהג"ה |  |
| 9:17.159 | שם | . |
| 9:17.160 | שזהו |  |
| 9:17.161 | האלהים |  |
| 9:17.162 | העליון |  |
| 9:17.163 | עשה |  |
| 9:17.164 | את |  |
| 9:17.165 | האדם |  |
| 9:17.166 | ישר |  |
| 9:17.167 | והמה |  |
| 9:17.168 | בקשו |  |
| 9:17.169 | חשבונות |  |
| 9:17.170 | רבים | ) |
| 9:17.171 | כי |  |
| 9:17.172 | האדם |  |
| 9:17.173 | קדמון |  |
| 9:17.174 | היה |  |
| 9:17.175 | בו |  |
| 9:17.176 | בחינת |  |
| 9:17.177 | יושר |  |
| 9:17.178 | וזהו |  |
| 9:17.179 | עשה |  |
| 9:17.180 | את |  |
| 9:17.181 | האדם |  |
| 9:17.182 | ישר |  |
| 9:17.183 | הגהה |  |
| 9:17.184 | א"מ |  |
| 9:17.185 | ןקרא |  |
| 9:17.186 | המאציל |  |
| 9:17.187 | אלהים |  |
| 9:17.188 | בדרך |  |
| 9:17.189 | שאלה | . |
| 9:17.190 | כי |  |
| 9:17.191 | הרי |  |
| 9:17.192 | אפי | ' |
| 9:17.193 | באריך |  |
| 9:17.194 | אנפין |  |
| 9:17.195 | אין |  |
| 9:17.196 | שם |  |
| 9:17.197 | אלהים | ) |
| 9:17.198 | ע"כ | . |
| 9:17.199 | הרי |  |
| 9:17.200 | לך |  |
| 9:17.201 | בהדיא |  |
| 9:17.202 | דפרצופים |  |
| 9:17.203 | העליונים |  |
| 9:17.204 | לא |  |
| 9:17.205 | שייכא |  |
| 9:17.206 | בהו |  |
| 9:17.207 | אלהות |  |
| 9:17.208 | וקריאה |  |
| 9:17.209 | כלל | , |
| 9:17.210 | כי |  |
| 9:17.211 | לא |  |
| 9:17.212 | אלהים |  |
| 9:17.213 | המה |  |
| 9:17.214 | ורק |  |
| 9:17.215 | להעלולים |  |
| 9:17.216 | התחתונים |  |
| 9:17.217 | נקרא |  |
| 9:17.218 | ונעבוד | , |
| 9:17.219 | והמה |  |
| 9:17.220 | אלהינו |  |
| 9:17.221 | לפי |  |
| 9:17.222 | דעתם | ) |
| 9:17.223 | כמו |  |
| 9:17.224 | שכתב |  |
| 9:17.225 | הרב |  |
| 9:17.226 | יושר |  |
| 9:17.227 | לבב |  |
| 9:17.228 | וספר |  |
| 9:17.229 | הברית |  |
| 9:17.230 | וכס"א |  |
| 9:17.231 | ונח"י |  |
| 9:17.232 | ומצרף |  |
| 9:17.233 | האמונה | ! |
| 9:17.234 | ולפי |  |
| 9:17.235 | דבריהם |  |
| 9:17.236 | שאומרים |  |
| 9:17.237 | דלית |  |
| 9:17.238 | רשותא |  |
| 9:17.239 | לבר |  |
| 9:17.240 | נש |  |
| 9:17.241 | למיעל |  |
| 9:17.242 | קמי |  |
| 9:17.243 | מריה |  |
| 9:17.244 | אלא |  |
| 9:17.245 | בידיעת |  |
| 9:17.246 | אמונה |  |
| 9:17.247 | זו | , |
| 9:17.248 | איך |  |
| 9:17.249 | אמרו |  |
| 9:17.250 | רז"ל |  |
| 9:17.251 | רצה |  |
| 9:17.252 | הקב"ה |  |
| 9:17.253 | לזכות |  |
| 9:17.254 | את |  |
| 9:17.255 | ישראל |  |
| 9:17.256 | לפיכך |  |
| 9:17.257 | הרבה |  |
| 9:17.258 | להם |  |
| 9:17.259 | תורה |  |
| 9:17.260 | ומצות |  |
| 9:17.261 | שנ | ' |
| 9:17.262 | ה | ' |
| 9:17.263 | חפץ |  |
| 9:17.264 | למען |  |
| 9:17.265 | צדקו |  |
| 9:17.266 | יגדיל |  |
| 9:17.267 | תורה |  |
| 9:17.268 | ויאדיר | ! |
| 9:17.269 | ולמה |  |
| 9:17.270 | זה |  |
| 9:17.271 | לא |  |
| 9:17.272 | רצה |  |
| 9:17.273 | הקב"ה |  |
| 9:17.274 | לזכות |  |
| 9:17.275 | אותנו |  |
| 9:17.276 | ולהודענו |  |
| 9:17.277 | על |  |
| 9:17.278 | ידי |  |
| 9:17.279 | נביאנו |  |
| 9:17.280 | הרבים |  |
| 9:17.281 | שעמדו |  |
| 9:17.282 | להם |  |
| 9:17.283 | לישראל |  |
| 9:17.284 | ממשה |  |
| 9:17.285 | רבינו |  |
| 9:17.286 | ע"ה |  |
| 9:17.287 | עד |  |
| 9:17.288 | שהנביא |  |
| 9:17.289 | האחרון |  |
| 9:17.290 | מלאכי |  |
| 9:17.291 | שאמר |  |
| 9:17.292 | זכרו |  |
| 9:17.293 | תורת |  |
| 9:17.294 | משה |  |
| 9:17.295 | עבדי |  |
| 9:17.296 | וכו | '. |
| 9:17.297 | את |  |
| 9:17.298 | אמונה |  |
| 9:17.299 | זו |  |
| 9:17.300 | שהעבודה |  |
| 9:17.301 | היא |  |
| 9:17.302 | להעלולים |  |
| 9:17.303 | התחתונים | . |
| 9:17.304 | ושהעובד |  |
| 9:17.305 | לאלהים |  |
| 9:17.306 | עליון |  |
| 9:17.307 | לא |  |
| 9:17.308 | יענהו |  |
| 9:17.309 | ואדרבא |  |
| 9:17.310 | ענוש |  |
| 9:17.311 | יענש |  |
| 9:17.312 | ומחה |  |
| 9:17.313 | ה | ' |
| 9:17.314 | את |  |
| 9:17.315 | שמו |  |
| 9:17.316 | כמו |  |
| 9:17.317 | שכתב |  |
| 9:17.318 | ספר |  |
| 9:17.319 | הברית | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

