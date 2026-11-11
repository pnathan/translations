# Translation pack: milhamot-hashem perek-9-12 (Milchamot ha-Shem 9:14–9:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-12`.

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
- ה׳ (the Tetragrammaton as written): "the LORD", in the author's prose and
  in verses. Where it is cited as a name ("זעיר אנפין דאיקרי ה'", 4:11), write
  "who is called 'the LORD'". Not ה׳ as Tetragrammaton: "מאת ה' אליעזר בן
  יהודה" (1:2) → "from Mr. Eliezer Ben-Yehuda"; ה׳ as a numeral (5) or a page
  letter → the number. The pack splits it into the token ה plus the
  punctuation ', so read the token in context.
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

- **מדה / מדות / מידות**: context. Do not render it "sefirot" even where that
  is the referent.
  - The hermeneutical rules: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי"
    (1:6) → "the thirteen rules of Rabbi Ishmael and the thirty-two rules of
    Rabbi Yose the Galilean"; "שלש עשרה מדות שהתורה נדרשת בהם" (10:5) → "the
    thirteen rules by which the Torah is expounded". Add a note: literally
    "measures".
  - Exod 34:6–7: "שלש עשרה מדות של רחמים" (9:8) → "the thirteen attributes of
    mercy". These are a different thirteen from the rules. The author uses
    both, so never let one rendering cover the other.
  - Kabbalistic and philosophical attributes, God's or the sefirot's: "מדת
    הדין" → "the attribute of judgment"; "המדות העליונות" (2:11) → "the
    upper attributes"; "המידות הפנימיות" (2:7) → "the inner attributes";
    "מדת המלכות" → "the attribute of Malkhut"; "מלובש במידותיו" (5:15) →
    "clothed in His attributes"; "ממדותיו" (7:17, Exod 34:6) → "of His
    attributes".
  - Measure: "עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה" (4:32) → "ten
    measures and ten sefirot, because they have measure and limit"; "לא מקום
    ולא מידה" (7:13) → "neither place nor measure"; "במדה שמדדתם" →
    "with the measure with which you measured".
  - A quality or trait: "ממדת הענוה" (4:8) → "of the quality of humility";
    "המדה המגונה הזאת" (7:19) → "this disgraceful trait"; "אין לך מדה גדולה
    מזו" (6:14, Talmud) → "there is no greater virtue than this".
  - A proper name: מסכת מדות → "tractate Middot".
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
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- An initial before a surname that fixes the person (ר"מ ריקנטי): expand and
  name him in a note ("Rabbi Menachem Rikanti"); an initial no surname fixes
  stays an initial with a note, as with R. A. and ר"י חיים ("Rabbi Y. Chayyim").
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
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


## Milchamot ha-Shem 9:14

Hebrew (Wikisource): קלב) וכיון שהעיקר והיסוד המוסד אצלינו שהקב"ה אינו צריך לבריותיו. ואדרבא בריותיו הם הצריכים לו ליתן להם שפע וחיות כמאמר הכתוב ואתה מחיה את כולם. וכמו שכתב הרמב"ם בעיקר הראשון שמהשי"ת קיום כל מציאות הנמצאים ואם נצייר בלבינו מניעת מציאותו, בטלה מציאות כל נמצא, ואלו נעלה על לבינו ביטול הנמצאים כולם. לא בטלה מציאותו בשבילם ולא חסרה כלל, כי הבורא יתברך אין לו צורך במציאותו לזולתו וכו'. ואחרי הדברים והאמת האלה, מי פתי יסור להאמין לדברי הר"מ נבאי אשר הרבה דברים לקלקל עלינו את הכתובים ודברי רז"ל המדברים בפה מלא ובלשון צח שאין הקב"ה צריך לבריותיו להשלים חסרונו ח"ו. והרחיב את פיו לדבר גדולות וגבוהות, בראיות צנומות דקות ויצא עתק מפיו, על הריב"ש ז"ל, ושיבש עלינו פירוש הכתובים ודברי רז"ל ברוב דברים אשר כל משכיל ידע ויכיר שאין בהם ממש, והאמת עד לעצמו שכל עבודותינו ומעשה המצות שצונו הקב"ה בתורה אינם רק לצרף וללבן אותנו, ונהיה טהורים ונקיים במדותינו! ולהרחיק ממנו הדעות והאמונות הרעות ונהיה קרובים אליו יתברך כמאמר רז"ל לא ניתנו המצות אלא לצרוף בהן את הבריות, שלא יאמר אדם אי אפשר וכו'. וכן אמרו ז"ל וכי מה איכפת לו להקב"ה בין שוחט מן הצואר לשוחט מן העורף, הוי אומר לא ניתנו המצות אלא לצרוף בהן את הבריות שנאמר אמרת ה' צרופה! לא לזווגו הוא צריך, ולא לקשט לו את כלתו! כי הוא אחד ומיוחד מכל האחדים שבעולם. וכלל הדברים פשוט הוא שהתורה ניתנה לתועלת מקבליה ומקיימיה לא לתועלת נותנה, שנאמר לטוב לנו כל הימים.

| id | word | punctuation after |
|---|---|---|
| 9:14.1 | קלב | ) |
| 9:14.2 | וכיון |  |
| 9:14.3 | שהעיקר |  |
| 9:14.4 | והיסוד |  |
| 9:14.5 | המוסד |  |
| 9:14.6 | אצלינו |  |
| 9:14.7 | שהקב"ה |  |
| 9:14.8 | אינו |  |
| 9:14.9 | צריך |  |
| 9:14.10 | לבריותיו | . |
| 9:14.11 | ואדרבא |  |
| 9:14.12 | בריותיו |  |
| 9:14.13 | הם |  |
| 9:14.14 | הצריכים |  |
| 9:14.15 | לו |  |
| 9:14.16 | ליתן |  |
| 9:14.17 | להם |  |
| 9:14.18 | שפע |  |
| 9:14.19 | וחיות |  |
| 9:14.20 | כמאמר |  |
| 9:14.21 | הכתוב |  |
| 9:14.22 | ואתה |  |
| 9:14.23 | מחיה |  |
| 9:14.24 | את |  |
| 9:14.25 | כולם | . |
| 9:14.26 | וכמו |  |
| 9:14.27 | שכתב |  |
| 9:14.28 | הרמב"ם |  |
| 9:14.29 | בעיקר |  |
| 9:14.30 | הראשון |  |
| 9:14.31 | שמהשי"ת |  |
| 9:14.32 | קיום |  |
| 9:14.33 | כל |  |
| 9:14.34 | מציאות |  |
| 9:14.35 | הנמצאים |  |
| 9:14.36 | ואם |  |
| 9:14.37 | נצייר |  |
| 9:14.38 | בלבינו |  |
| 9:14.39 | מניעת |  |
| 9:14.40 | מציאותו | , |
| 9:14.41 | בטלה |  |
| 9:14.42 | מציאות |  |
| 9:14.43 | כל |  |
| 9:14.44 | נמצא | , |
| 9:14.45 | ואלו |  |
| 9:14.46 | נעלה |  |
| 9:14.47 | על |  |
| 9:14.48 | לבינו |  |
| 9:14.49 | ביטול |  |
| 9:14.50 | הנמצאים |  |
| 9:14.51 | כולם | . |
| 9:14.52 | לא |  |
| 9:14.53 | בטלה |  |
| 9:14.54 | מציאותו |  |
| 9:14.55 | בשבילם |  |
| 9:14.56 | ולא |  |
| 9:14.57 | חסרה |  |
| 9:14.58 | כלל | , |
| 9:14.59 | כי |  |
| 9:14.60 | הבורא |  |
| 9:14.61 | יתברך |  |
| 9:14.62 | אין |  |
| 9:14.63 | לו |  |
| 9:14.64 | צורך |  |
| 9:14.65 | במציאותו |  |
| 9:14.66 | לזולתו |  |
| 9:14.67 | וכו | '. |
| 9:14.68 | ואחרי |  |
| 9:14.69 | הדברים |  |
| 9:14.70 | והאמת |  |
| 9:14.71 | האלה | , |
| 9:14.72 | מי |  |
| 9:14.73 | פתי |  |
| 9:14.74 | יסור |  |
| 9:14.75 | להאמין |  |
| 9:14.76 | לדברי |  |
| 9:14.77 | הר"מ |  |
| 9:14.78 | נבאי |  |
| 9:14.79 | אשר |  |
| 9:14.80 | הרבה |  |
| 9:14.81 | דברים |  |
| 9:14.82 | לקלקל |  |
| 9:14.83 | עלינו |  |
| 9:14.84 | את |  |
| 9:14.85 | הכתובים |  |
| 9:14.86 | ודברי |  |
| 9:14.87 | רז"ל |  |
| 9:14.88 | המדברים |  |
| 9:14.89 | בפה |  |
| 9:14.90 | מלא |  |
| 9:14.91 | ובלשון |  |
| 9:14.92 | צח |  |
| 9:14.93 | שאין |  |
| 9:14.94 | הקב"ה |  |
| 9:14.95 | צריך |  |
| 9:14.96 | לבריותיו |  |
| 9:14.97 | להשלים |  |
| 9:14.98 | חסרונו |  |
| 9:14.99 | ח"ו | . |
| 9:14.100 | והרחיב |  |
| 9:14.101 | את |  |
| 9:14.102 | פיו |  |
| 9:14.103 | לדבר |  |
| 9:14.104 | גדולות |  |
| 9:14.105 | וגבוהות | , |
| 9:14.106 | בראיות |  |
| 9:14.107 | צנומות |  |
| 9:14.108 | דקות |  |
| 9:14.109 | ויצא |  |
| 9:14.110 | עתק |  |
| 9:14.111 | מפיו | , |
| 9:14.112 | על |  |
| 9:14.113 | הריב"ש |  |
| 9:14.114 | ז"ל | , |
| 9:14.115 | ושיבש |  |
| 9:14.116 | עלינו |  |
| 9:14.117 | פירוש |  |
| 9:14.118 | הכתובים |  |
| 9:14.119 | ודברי |  |
| 9:14.120 | רז"ל |  |
| 9:14.121 | ברוב |  |
| 9:14.122 | דברים |  |
| 9:14.123 | אשר |  |
| 9:14.124 | כל |  |
| 9:14.125 | משכיל |  |
| 9:14.126 | ידע |  |
| 9:14.127 | ויכיר |  |
| 9:14.128 | שאין |  |
| 9:14.129 | בהם |  |
| 9:14.130 | ממש | , |
| 9:14.131 | והאמת |  |
| 9:14.132 | עד |  |
| 9:14.133 | לעצמו |  |
| 9:14.134 | שכל |  |
| 9:14.135 | עבודותינו |  |
| 9:14.136 | ומעשה |  |
| 9:14.137 | המצות |  |
| 9:14.138 | שצונו |  |
| 9:14.139 | הקב"ה |  |
| 9:14.140 | בתורה |  |
| 9:14.141 | אינם |  |
| 9:14.142 | רק |  |
| 9:14.143 | לצרף |  |
| 9:14.144 | וללבן |  |
| 9:14.145 | אותנו | , |
| 9:14.146 | ונהיה |  |
| 9:14.147 | טהורים |  |
| 9:14.148 | ונקיים |  |
| 9:14.149 | במדותינו | ! |
| 9:14.150 | ולהרחיק |  |
| 9:14.151 | ממנו |  |
| 9:14.152 | הדעות |  |
| 9:14.153 | והאמונות |  |
| 9:14.154 | הרעות |  |
| 9:14.155 | ונהיה |  |
| 9:14.156 | קרובים |  |
| 9:14.157 | אליו |  |
| 9:14.158 | יתברך |  |
| 9:14.159 | כמאמר |  |
| 9:14.160 | רז"ל |  |
| 9:14.161 | לא |  |
| 9:14.162 | ניתנו |  |
| 9:14.163 | המצות |  |
| 9:14.164 | אלא |  |
| 9:14.165 | לצרוף |  |
| 9:14.166 | בהן |  |
| 9:14.167 | את |  |
| 9:14.168 | הבריות | , |
| 9:14.169 | שלא |  |
| 9:14.170 | יאמר |  |
| 9:14.171 | אדם |  |
| 9:14.172 | אי |  |
| 9:14.173 | אפשר |  |
| 9:14.174 | וכו | '. |
| 9:14.175 | וכן |  |
| 9:14.176 | אמרו |  |
| 9:14.177 | ז"ל |  |
| 9:14.178 | וכי |  |
| 9:14.179 | מה |  |
| 9:14.180 | איכפת |  |
| 9:14.181 | לו |  |
| 9:14.182 | להקב"ה |  |
| 9:14.183 | בין |  |
| 9:14.184 | שוחט |  |
| 9:14.185 | מן |  |
| 9:14.186 | הצואר |  |
| 9:14.187 | לשוחט |  |
| 9:14.188 | מן |  |
| 9:14.189 | העורף | , |
| 9:14.190 | הוי |  |
| 9:14.191 | אומר |  |
| 9:14.192 | לא |  |
| 9:14.193 | ניתנו |  |
| 9:14.194 | המצות |  |
| 9:14.195 | אלא |  |
| 9:14.196 | לצרוף |  |
| 9:14.197 | בהן |  |
| 9:14.198 | את |  |
| 9:14.199 | הבריות |  |
| 9:14.200 | שנאמר |  |
| 9:14.201 | אמרת |  |
| 9:14.202 | ה | ' |
| 9:14.203 | צרופה | ! |
| 9:14.204 | לא |  |
| 9:14.205 | לזווגו |  |
| 9:14.206 | הוא |  |
| 9:14.207 | צריך | , |
| 9:14.208 | ולא |  |
| 9:14.209 | לקשט |  |
| 9:14.210 | לו |  |
| 9:14.211 | את |  |
| 9:14.212 | כלתו | ! |
| 9:14.213 | כי |  |
| 9:14.214 | הוא |  |
| 9:14.215 | אחד |  |
| 9:14.216 | ומיוחד |  |
| 9:14.217 | מכל |  |
| 9:14.218 | האחדים |  |
| 9:14.219 | שבעולם | . |
| 9:14.220 | וכלל |  |
| 9:14.221 | הדברים |  |
| 9:14.222 | פשוט |  |
| 9:14.223 | הוא |  |
| 9:14.224 | שהתורה |  |
| 9:14.225 | ניתנה |  |
| 9:14.226 | לתועלת |  |
| 9:14.227 | מקבליה |  |
| 9:14.228 | ומקיימיה |  |
| 9:14.229 | לא |  |
| 9:14.230 | לתועלת |  |
| 9:14.231 | נותנה | , |
| 9:14.232 | שנאמר |  |
| 9:14.233 | לטוב |  |
| 9:14.234 | לנו |  |
| 9:14.235 | כל |  |
| 9:14.236 | הימים | . |


## Milchamot ha-Shem 9:15

Hebrew (Wikisource): קלג) ועל כל אלה הדעות הרעות החיצונות והזרות לנו לתורתינו הקדושה הזהיר הכתוב אותנו באמרו כי יסיתך איחך בן אמך או בנך או בתך או אשת חיקך או רעך אשר כנפשך בסתר לאמר נלכה ונעבדה אלהים אחרים אשר לא ידעת אתה ואבתיך, מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים ממך בסתר! אם יאמר בא ואגלה לך דברים נסתרים וסודות עמוקים לאמר, אלילים יש בשמים ממעל בכמה פרצופים זה למעלה מזה, וכך היא אוכלת וכך היא שותה, שמושכת לה מוחין למזונה וחיותה תמיד מאצל העתיק, ועל ידי מוחין שמושך ויונק פרצוף התחתון מהפרצוף העליון על ידי תפלותינו ובקשותינו ממנו הוא משפיע לנו שפע והארה, ובכך היא מטיבה לנו! וכך היא מריעה למי שלא יעבוד אותה, שכן צוו עליה אבא ואימא דכלא יפלחון ליה לזעיר אנפין הוא ואשתו, ומי שלא יעבוד רק לאלהים עליון ענוש יענש, שכך רצון המלך העליון שנתפלל ונקרא לאלוה קצר אפים (זעור) בצירוף הפרצופים שלמעלה ממנו כנזכר לעיל כמה פעמים בשם ספר הברית וכס"א ונח"י ויושר לבב! לא תאבה לו ולא תשמע אליו וכו'! ובסנהדרין דף ס"א אמרינן לא תאבה לו ולא תשמע אליו, הא אבה ושמע חייב, ואביי מפליג בה בין ניסת מפי עצמו לניסת מפי אחרים, דניסת מפי עצמו מימליך, אבל ניסת מפי אחרים גריד בתריהו, ורבא אמר אידי ואידי בניסת מפי אחרים, הא אמר ליה כך אוכלת, כך שותה, כך מטיבה כך מריעה, הא דלא אמר ליה כך אוכלת וכך שותה וכו'. אמר רבא מנא אמינא לה דכתיב מאלהי העמים אשר סביבותיכם הקרובים אליך או הרחוקים ממך, מה לי קרובים מה לי רחוקים אלא הכי קאמר, מטיבותן של קרובים אצה למד טיבותן של רחוקים ע"כ, ופירש רש"י ז"ל דכתיב בפרשת מסית מאלהי העמים הקרובים אליך או הרחוקים ממך, מה לי קרובים מה לי רחוקים. ולמה הזכיר קרוב ורחוק אלא מפני שדרכו של מסית לדבר על ע"ז רחוקה שאינו מכיר בה, ומוציא מפיו השקר ואומר כך אוכלת כך שותה, לפיכך אני אומר לך הסתכל בע"ז הקרובות אליך ותראה שאין בהם ממש, ומהן אתה למד טיבן של רחוקים ע"כ. ועיין לעיל סי' ט"ל.

| id | word | punctuation after |
|---|---|---|
| 9:15.1 | קלג | ) |
| 9:15.2 | ועל |  |
| 9:15.3 | כל |  |
| 9:15.4 | אלה |  |
| 9:15.5 | הדעות |  |
| 9:15.6 | הרעות |  |
| 9:15.7 | החיצונות |  |
| 9:15.8 | והזרות |  |
| 9:15.9 | לנו |  |
| 9:15.10 | לתורתינו |  |
| 9:15.11 | הקדושה |  |
| 9:15.12 | הזהיר |  |
| 9:15.13 | הכתוב |  |
| 9:15.14 | אותנו |  |
| 9:15.15 | באמרו |  |
| 9:15.16 | כי |  |
| 9:15.17 | יסיתך |  |
| 9:15.18 | איחך |  |
| 9:15.19 | בן |  |
| 9:15.20 | אמך |  |
| 9:15.21 | או |  |
| 9:15.22 | בנך |  |
| 9:15.23 | או |  |
| 9:15.24 | בתך |  |
| 9:15.25 | או |  |
| 9:15.26 | אשת |  |
| 9:15.27 | חיקך |  |
| 9:15.28 | או |  |
| 9:15.29 | רעך |  |
| 9:15.30 | אשר |  |
| 9:15.31 | כנפשך |  |
| 9:15.32 | בסתר |  |
| 9:15.33 | לאמר |  |
| 9:15.34 | נלכה |  |
| 9:15.35 | ונעבדה |  |
| 9:15.36 | אלהים |  |
| 9:15.37 | אחרים |  |
| 9:15.38 | אשר |  |
| 9:15.39 | לא |  |
| 9:15.40 | ידעת |  |
| 9:15.41 | אתה |  |
| 9:15.42 | ואבתיך | , |
| 9:15.43 | מאלהי |  |
| 9:15.44 | העמים |  |
| 9:15.45 | אשר |  |
| 9:15.46 | סביבותיכם |  |
| 9:15.47 | הקרובים |  |
| 9:15.48 | אליך |  |
| 9:15.49 | או |  |
| 9:15.50 | הרחוקים |  |
| 9:15.51 | ממך |  |
| 9:15.52 | בסתר | ! |
| 9:15.53 | אם |  |
| 9:15.54 | יאמר |  |
| 9:15.55 | בא |  |
| 9:15.56 | ואגלה |  |
| 9:15.57 | לך |  |
| 9:15.58 | דברים |  |
| 9:15.59 | נסתרים |  |
| 9:15.60 | וסודות |  |
| 9:15.61 | עמוקים |  |
| 9:15.62 | לאמר | , |
| 9:15.63 | אלילים |  |
| 9:15.64 | יש |  |
| 9:15.65 | בשמים |  |
| 9:15.66 | ממעל |  |
| 9:15.67 | בכמה |  |
| 9:15.68 | פרצופים |  |
| 9:15.69 | זה |  |
| 9:15.70 | למעלה |  |
| 9:15.71 | מזה | , |
| 9:15.72 | וכך |  |
| 9:15.73 | היא |  |
| 9:15.74 | אוכלת |  |
| 9:15.75 | וכך |  |
| 9:15.76 | היא |  |
| 9:15.77 | שותה | , |
| 9:15.78 | שמושכת |  |
| 9:15.79 | לה |  |
| 9:15.80 | מוחין |  |
| 9:15.81 | למזונה |  |
| 9:15.82 | וחיותה |  |
| 9:15.83 | תמיד |  |
| 9:15.84 | מאצל |  |
| 9:15.85 | העתיק | , |
| 9:15.86 | ועל |  |
| 9:15.87 | ידי |  |
| 9:15.88 | מוחין |  |
| 9:15.89 | שמושך |  |
| 9:15.90 | ויונק |  |
| 9:15.91 | פרצוף |  |
| 9:15.92 | התחתון |  |
| 9:15.93 | מהפרצוף |  |
| 9:15.94 | העליון |  |
| 9:15.95 | על |  |
| 9:15.96 | ידי |  |
| 9:15.97 | תפלותינו |  |
| 9:15.98 | ובקשותינו |  |
| 9:15.99 | ממנו |  |
| 9:15.100 | הוא |  |
| 9:15.101 | משפיע |  |
| 9:15.102 | לנו |  |
| 9:15.103 | שפע |  |
| 9:15.104 | והארה | , |
| 9:15.105 | ובכך |  |
| 9:15.106 | היא |  |
| 9:15.107 | מטיבה |  |
| 9:15.108 | לנו | ! |
| 9:15.109 | וכך |  |
| 9:15.110 | היא |  |
| 9:15.111 | מריעה |  |
| 9:15.112 | למי |  |
| 9:15.113 | שלא |  |
| 9:15.114 | יעבוד |  |
| 9:15.115 | אותה | , |
| 9:15.116 | שכן |  |
| 9:15.117 | צוו |  |
| 9:15.118 | עליה |  |
| 9:15.119 | אבא |  |
| 9:15.120 | ואימא |  |
| 9:15.121 | דכלא |  |
| 9:15.122 | יפלחון |  |
| 9:15.123 | ליה |  |
| 9:15.124 | לזעיר |  |
| 9:15.125 | אנפין |  |
| 9:15.126 | הוא |  |
| 9:15.127 | ואשתו | , |
| 9:15.128 | ומי |  |
| 9:15.129 | שלא |  |
| 9:15.130 | יעבוד |  |
| 9:15.131 | רק |  |
| 9:15.132 | לאלהים |  |
| 9:15.133 | עליון |  |
| 9:15.134 | ענוש |  |
| 9:15.135 | יענש | , |
| 9:15.136 | שכך |  |
| 9:15.137 | רצון |  |
| 9:15.138 | המלך |  |
| 9:15.139 | העליון |  |
| 9:15.140 | שנתפלל |  |
| 9:15.141 | ונקרא |  |
| 9:15.142 | לאלוה |  |
| 9:15.143 | קצר |  |
| 9:15.144 | אפים |  |
| 9:15.145 | זעור | ) |
| 9:15.146 | בצירוף |  |
| 9:15.147 | הפרצופים |  |
| 9:15.148 | שלמעלה |  |
| 9:15.149 | ממנו |  |
| 9:15.150 | כנזכר |  |
| 9:15.151 | לעיל |  |
| 9:15.152 | כמה |  |
| 9:15.153 | פעמים |  |
| 9:15.154 | בשם |  |
| 9:15.155 | ספר |  |
| 9:15.156 | הברית |  |
| 9:15.157 | וכס"א |  |
| 9:15.158 | ונח"י |  |
| 9:15.159 | ויושר |  |
| 9:15.160 | לבב | ! |
| 9:15.161 | לא |  |
| 9:15.162 | תאבה |  |
| 9:15.163 | לו |  |
| 9:15.164 | ולא |  |
| 9:15.165 | תשמע |  |
| 9:15.166 | אליו |  |
| 9:15.167 | וכו | '! |
| 9:15.168 | ובסנהדרין |  |
| 9:15.169 | דף |  |
| 9:15.170 | ס"א |  |
| 9:15.171 | אמרינן |  |
| 9:15.172 | לא |  |
| 9:15.173 | תאבה |  |
| 9:15.174 | לו |  |
| 9:15.175 | ולא |  |
| 9:15.176 | תשמע |  |
| 9:15.177 | אליו | , |
| 9:15.178 | הא |  |
| 9:15.179 | אבה |  |
| 9:15.180 | ושמע |  |
| 9:15.181 | חייב | , |
| 9:15.182 | ואביי |  |
| 9:15.183 | מפליג |  |
| 9:15.184 | בה |  |
| 9:15.185 | בין |  |
| 9:15.186 | ניסת |  |
| 9:15.187 | מפי |  |
| 9:15.188 | עצמו |  |
| 9:15.189 | לניסת |  |
| 9:15.190 | מפי |  |
| 9:15.191 | אחרים | , |
| 9:15.192 | דניסת |  |
| 9:15.193 | מפי |  |
| 9:15.194 | עצמו |  |
| 9:15.195 | מימליך | , |
| 9:15.196 | אבל |  |
| 9:15.197 | ניסת |  |
| 9:15.198 | מפי |  |
| 9:15.199 | אחרים |  |
| 9:15.200 | גריד |  |
| 9:15.201 | בתריהו | , |
| 9:15.202 | ורבא |  |
| 9:15.203 | אמר |  |
| 9:15.204 | אידי |  |
| 9:15.205 | ואידי |  |
| 9:15.206 | בניסת |  |
| 9:15.207 | מפי |  |
| 9:15.208 | אחרים | , |
| 9:15.209 | הא |  |
| 9:15.210 | אמר |  |
| 9:15.211 | ליה |  |
| 9:15.212 | כך |  |
| 9:15.213 | אוכלת | , |
| 9:15.214 | כך |  |
| 9:15.215 | שותה | , |
| 9:15.216 | כך |  |
| 9:15.217 | מטיבה |  |
| 9:15.218 | כך |  |
| 9:15.219 | מריעה | , |
| 9:15.220 | הא |  |
| 9:15.221 | דלא |  |
| 9:15.222 | אמר |  |
| 9:15.223 | ליה |  |
| 9:15.224 | כך |  |
| 9:15.225 | אוכלת |  |
| 9:15.226 | וכך |  |
| 9:15.227 | שותה |  |
| 9:15.228 | וכו | '. |
| 9:15.229 | אמר |  |
| 9:15.230 | רבא |  |
| 9:15.231 | מנא |  |
| 9:15.232 | אמינא |  |
| 9:15.233 | לה |  |
| 9:15.234 | דכתיב |  |
| 9:15.235 | מאלהי |  |
| 9:15.236 | העמים |  |
| 9:15.237 | אשר |  |
| 9:15.238 | סביבותיכם |  |
| 9:15.239 | הקרובים |  |
| 9:15.240 | אליך |  |
| 9:15.241 | או |  |
| 9:15.242 | הרחוקים |  |
| 9:15.243 | ממך | , |
| 9:15.244 | מה |  |
| 9:15.245 | לי |  |
| 9:15.246 | קרובים |  |
| 9:15.247 | מה |  |
| 9:15.248 | לי |  |
| 9:15.249 | רחוקים |  |
| 9:15.250 | אלא |  |
| 9:15.251 | הכי |  |
| 9:15.252 | קאמר | , |
| 9:15.253 | מטיבותן |  |
| 9:15.254 | של |  |
| 9:15.255 | קרובים |  |
| 9:15.256 | אצה |  |
| 9:15.257 | למד |  |
| 9:15.258 | טיבותן |  |
| 9:15.259 | של |  |
| 9:15.260 | רחוקים |  |
| 9:15.261 | ע"כ | , |
| 9:15.262 | ופירש |  |
| 9:15.263 | רש"י |  |
| 9:15.264 | ז"ל |  |
| 9:15.265 | דכתיב |  |
| 9:15.266 | בפרשת |  |
| 9:15.267 | מסית |  |
| 9:15.268 | מאלהי |  |
| 9:15.269 | העמים |  |
| 9:15.270 | הקרובים |  |
| 9:15.271 | אליך |  |
| 9:15.272 | או |  |
| 9:15.273 | הרחוקים |  |
| 9:15.274 | ממך | , |
| 9:15.275 | מה |  |
| 9:15.276 | לי |  |
| 9:15.277 | קרובים |  |
| 9:15.278 | מה |  |
| 9:15.279 | לי |  |
| 9:15.280 | רחוקים | . |
| 9:15.281 | ולמה |  |
| 9:15.282 | הזכיר |  |
| 9:15.283 | קרוב |  |
| 9:15.284 | ורחוק |  |
| 9:15.285 | אלא |  |
| 9:15.286 | מפני |  |
| 9:15.287 | שדרכו |  |
| 9:15.288 | של |  |
| 9:15.289 | מסית |  |
| 9:15.290 | לדבר |  |
| 9:15.291 | על |  |
| 9:15.292 | ע"ז |  |
| 9:15.293 | רחוקה |  |
| 9:15.294 | שאינו |  |
| 9:15.295 | מכיר |  |
| 9:15.296 | בה | , |
| 9:15.297 | ומוציא |  |
| 9:15.298 | מפיו |  |
| 9:15.299 | השקר |  |
| 9:15.300 | ואומר |  |
| 9:15.301 | כך |  |
| 9:15.302 | אוכלת |  |
| 9:15.303 | כך |  |
| 9:15.304 | שותה | , |
| 9:15.305 | לפיכך |  |
| 9:15.306 | אני |  |
| 9:15.307 | אומר |  |
| 9:15.308 | לך |  |
| 9:15.309 | הסתכל |  |
| 9:15.310 | בע"ז |  |
| 9:15.311 | הקרובות |  |
| 9:15.312 | אליך |  |
| 9:15.313 | ותראה |  |
| 9:15.314 | שאין |  |
| 9:15.315 | בהם |  |
| 9:15.316 | ממש | , |
| 9:15.317 | ומהן |  |
| 9:15.318 | אתה |  |
| 9:15.319 | למד |  |
| 9:15.320 | טיבן |  |
| 9:15.321 | של |  |
| 9:15.322 | רחוקים |  |
| 9:15.323 | ע"כ | . |
| 9:15.324 | ועיין |  |
| 9:15.325 | לעיל |  |
| 9:15.326 | סי | ' |
| 9:15.327 | ט"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

