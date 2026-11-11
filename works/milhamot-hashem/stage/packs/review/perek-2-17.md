# Review pack: milhamot-hashem perek-2-17 (Milchamot ha-Shem 2:33–2:33)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-17`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
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
- **אצילות** and its root: fixed and translated, so that the verbs the
  author strings together stay linked. אצילות → "emanation"; עולם האצילות /
  דאצילות → "the World of Emanation"; נאצל → "emanated"; הנאצלים →
  "the emanated beings"; המאציל → "the Emanator". The four worlds:
  "Emanation, Creation, Formation, Making" (אבי"ע). "נאצלו ונבראו ונוצרו
  ונעשו" → "were emanated, created, formed and made". Keep נאצל and נברא
  distinct everywhere: the author argues that Ze'ir Anpin is "נברא ולא בורא"
  → "created, not creator".
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
- ב"ע: "in Aramaic" (בארמית) where it is set against בערבי; the other readings in
  `readings`.
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


## Milchamot ha-Shem 2:33

Hebrew (Wikisource): הראית לדעת דברי הגאון חוות יאיר כמה הוא ירא וחרד מלהורות על הפרישה ומסתייע בדברי הגאון בעל שמן למאור הובאו דבריו למעלה, ועדין הוא ירא לאמר בהדיא פן יגרום רפיון ידים מעסק החכמה הקדושה הזאת, וכמה הוא משבח החכמה הזאת לדזכי לה ללמדה מפי מקובל, ומודה שכל החושך עצמו מללמוד בה אחר שקיבלה הוא נדחה ממחיצתו, וגם כן הוא מתיר לעסוק בדברי קבלה ואם יראה לפניו דברים תמוהים שיקבלם באמונת אומן, ויתלה בקצר דעתו. א"כ שני עדיך הריב"ש וחוות יאיר יעידון ויגידון בך שאתה בודה דברים מלבך, והיכן היא פרישתו שפירש כמו שאמרת? והלא הוא עצמו אמר מי יתן ויערה עלינו רוח ממרום וכו'. ואני מביא לך ראיה שזכה לחכמת הקבלה, שהרי בסי' קע"ב כתב וז"ל שאפשר שיש עוד דברי סוד וסתר בשני פסוקים דאלפים בת יכיל וג' אלפים בת' כי בחטא אדם הראשון נתקלקלו שלשה אלפין מאמת ומאדנות ומאדם, ונשאר מת-דין-דם נגד ע"ז, גילוי עריות ושפיכות דמים. ודוד המלך תיקן אמת, לכן לא טעם שינה בלילה, ולא הספיק לתקן דין גם דם, כמו שאמר הקב"ה לא תבנה הבית כי דמים רבים שפכת, רק בנך שלמה כי שלום וכו'. ותקן שלמה שני אלפין ועם האלף דאביו נתקנו ג' אלפין, וענין בת ידוע לי"ח כי עיקר הפגם שם, ועוד האריך בסוד הענין שם. וסי' קע"ו לענין הזן סומא באחת מעיניו, ועם כל זה קראתי תיגר כי נר לי דבתרווייהו איכא למיחד מיהו, היכא דאיכא אחר הגון וראוי כיוצא בו כי ידוע שרמ"ח איברים הם כסא ודמות לרמ"ח אורות עליונים ורמ"ח איברים של נשמה. ואם כן כל כי האי גונא הכסא פגום, כי על פי הנסתר יש תלי תלים סודות נסתרים באיברי הגוף, ויש חסרון שפע בכוס של ברכה. העליון ג"כ בחתימת הספר כ' שעשה תרל"ה תשובות וליקט מהן רל"ב מספר יהי אור והוא בגימטריא ד' מילוי הויה ע"ב ס"ג מ"ה בן שבאבא ואימא וזעיר ונוקביה. ועוד יהי אור יה יאיר. שהם מקורי תפארת ומלכות שבבינה שהוא אור הנאצל ביום ראשון ועוד האריך בדברי סוד וסתר הרבה. ג"כ אביו בעל שמן למאור היה מקובל אלהי שהרי נדפסו חרוזיו בסוף הספר. והן על פי שם בן מ"ב וי"ב צירופי הוי'ה וכ"ד צירופי אדנות ושם בן כ"ב, גם החכם האמתי בעל מנורת המאור היה מקובל אלהי שהרי כתב בפ' צ"ג על ברוך שאמר דרכו בה מספר העשרה העשרה וכו'. ובפ' צ"ה כ' אין אומרים קדיש בפחות מעשרה, והטעם לפי שבעשרה נשלמה הקדושה שהיא למעלה מעשר ספירות, ובפ' קל"ג כתב וכשהיה יעקב יורד למצרים ראה כי הגלות התחיל בו ובזרעו ופחד ממנו וזבח זבחים לפחד אביו יצחק שלא תהא מדת הדין מתוחה כנגדו, ועשה זה בבאר שבע שהיא בית תפלה לאבותיו, והקריב שלמים להשלים אליו כל המדות כמו שדרשו שלמים שמטילים שלום להשלים אליו מדת הגבורה, ובפ' קמ"ז כתב נמצא כי שלמות הקדושה היא בעשר הרומזים לעשר הויות והם עשרה מאמרות שבהם נברא העולם. ובפ' קנ"ג כתב ועל דרך האמת כי ששת ימים עשה ה' וכו' וביום השביעי הוא שבת ואין לו בן זוג, וכנסת ישראל היא בת זוגו שנ' ואת הארץ והנה הוא שמיני עצרת, כי שם נעצר הכל. וזהו מאמרם שמיני רגל בפני עצמו ותשלומי ראשון הוא כי הוא אצילות הראשונים ואינו באחדות שלהם, ומי שיבין בהויות ישכיל וימצא האמת, ובחתימת הספר כתב והחמישי עזבתי הדברים החמורים שהם נסתרים וגבוהים ואין מי שישיג אותם אלא יחיד בדור, הרי שהיה רוצה לכתוב דברי סתר בחבורו. אלא מפני שאינן ראוים לכל אדם. הראית לדעת שאין מי שיחלוק על הקבלה. לא הראשונים ולא מאחרונים לא רבין ולא יחידים, לא גדולים ולא קטנים. אדרבה הכל מסכימים שמי שיזכה לקבלה מפה לאזן אין לך מדה טובה ממנה. ושאסור לדרוש ולחקור עליה, אלא אם ירצה יקבלה באמונה שלימה ולא יפרשה בשכלו האנושי. כי יתלה מה שהקשה עליו בקוצר דעתו ולא יהרוס לחשוב שום הנשמה. ובזה נשלמו כל דברי שפל השפלים. סדר ואתה תשוב ושמעת בקול ה'.

| id | word | punctuation after |
|---|---|---|
| 2:33.1 | הראית |  |
| 2:33.2 | לדעת |  |
| 2:33.3 | דברי |  |
| 2:33.4 | הגאון |  |
| 2:33.5 | חוות |  |
| 2:33.6 | יאיר |  |
| 2:33.7 | כמה |  |
| 2:33.8 | הוא |  |
| 2:33.9 | ירא |  |
| 2:33.10 | וחרד |  |
| 2:33.11 | מלהורות |  |
| 2:33.12 | על |  |
| 2:33.13 | הפרישה |  |
| 2:33.14 | ומסתייע |  |
| 2:33.15 | בדברי |  |
| 2:33.16 | הגאון |  |
| 2:33.17 | בעל |  |
| 2:33.18 | שמן |  |
| 2:33.19 | למאור |  |
| 2:33.20 | הובאו |  |
| 2:33.21 | דבריו |  |
| 2:33.22 | למעלה | , |
| 2:33.23 | ועדין |  |
| 2:33.24 | הוא |  |
| 2:33.25 | ירא |  |
| 2:33.26 | לאמר |  |
| 2:33.27 | בהדיא |  |
| 2:33.28 | פן |  |
| 2:33.29 | יגרום |  |
| 2:33.30 | רפיון |  |
| 2:33.31 | ידים |  |
| 2:33.32 | מעסק |  |
| 2:33.33 | החכמה |  |
| 2:33.34 | הקדושה |  |
| 2:33.35 | הזאת | , |
| 2:33.36 | וכמה |  |
| 2:33.37 | הוא |  |
| 2:33.38 | משבח |  |
| 2:33.39 | החכמה |  |
| 2:33.40 | הזאת |  |
| 2:33.41 | לדזכי |  |
| 2:33.42 | לה |  |
| 2:33.43 | ללמדה |  |
| 2:33.44 | מפי |  |
| 2:33.45 | מקובל | , |
| 2:33.46 | ומודה |  |
| 2:33.47 | שכל |  |
| 2:33.48 | החושך |  |
| 2:33.49 | עצמו |  |
| 2:33.50 | מללמוד |  |
| 2:33.51 | בה |  |
| 2:33.52 | אחר |  |
| 2:33.53 | שקיבלה |  |
| 2:33.54 | הוא |  |
| 2:33.55 | נדחה |  |
| 2:33.56 | ממחיצתו | , |
| 2:33.57 | וגם |  |
| 2:33.58 | כן |  |
| 2:33.59 | הוא |  |
| 2:33.60 | מתיר |  |
| 2:33.61 | לעסוק |  |
| 2:33.62 | בדברי |  |
| 2:33.63 | קבלה |  |
| 2:33.64 | ואם |  |
| 2:33.65 | יראה |  |
| 2:33.66 | לפניו |  |
| 2:33.67 | דברים |  |
| 2:33.68 | תמוהים |  |
| 2:33.69 | שיקבלם |  |
| 2:33.70 | באמונת |  |
| 2:33.71 | אומן | , |
| 2:33.72 | ויתלה |  |
| 2:33.73 | בקצר |  |
| 2:33.74 | דעתו | . |
| 2:33.75 | א"כ |  |
| 2:33.76 | שני |  |
| 2:33.77 | עדיך |  |
| 2:33.78 | הריב"ש |  |
| 2:33.79 | וחוות |  |
| 2:33.80 | יאיר |  |
| 2:33.81 | יעידון |  |
| 2:33.82 | ויגידון |  |
| 2:33.83 | בך |  |
| 2:33.84 | שאתה |  |
| 2:33.85 | בודה |  |
| 2:33.86 | דברים |  |
| 2:33.87 | מלבך | , |
| 2:33.88 | והיכן |  |
| 2:33.89 | היא |  |
| 2:33.90 | פרישתו |  |
| 2:33.91 | שפירש |  |
| 2:33.92 | כמו |  |
| 2:33.93 | שאמרת | ? |
| 2:33.94 | והלא |  |
| 2:33.95 | הוא |  |
| 2:33.96 | עצמו |  |
| 2:33.97 | אמר |  |
| 2:33.98 | מי |  |
| 2:33.99 | יתן |  |
| 2:33.100 | ויערה |  |
| 2:33.101 | עלינו |  |
| 2:33.102 | רוח |  |
| 2:33.103 | ממרום |  |
| 2:33.104 | וכו | '. |
| 2:33.105 | ואני |  |
| 2:33.106 | מביא |  |
| 2:33.107 | לך |  |
| 2:33.108 | ראיה |  |
| 2:33.109 | שזכה |  |
| 2:33.110 | לחכמת |  |
| 2:33.111 | הקבלה | , |
| 2:33.112 | שהרי |  |
| 2:33.113 | בסי | ' |
| 2:33.114 | קע"ב |  |
| 2:33.115 | כתב |  |
| 2:33.116 | וז"ל |  |
| 2:33.117 | שאפשר |  |
| 2:33.118 | שיש |  |
| 2:33.119 | עוד |  |
| 2:33.120 | דברי |  |
| 2:33.121 | סוד |  |
| 2:33.122 | וסתר |  |
| 2:33.123 | בשני |  |
| 2:33.124 | פסוקים |  |
| 2:33.125 | דאלפים |  |
| 2:33.126 | בת |  |
| 2:33.127 | יכיל |  |
| 2:33.128 | וג | ' |
| 2:33.129 | אלפים |  |
| 2:33.130 | בת | ' |
| 2:33.131 | כי |  |
| 2:33.132 | בחטא |  |
| 2:33.133 | אדם |  |
| 2:33.134 | הראשון |  |
| 2:33.135 | נתקלקלו |  |
| 2:33.136 | שלשה |  |
| 2:33.137 | אלפין |  |
| 2:33.138 | מאמת |  |
| 2:33.139 | ומאדנות |  |
| 2:33.140 | ומאדם | , |
| 2:33.141 | ונשאר |  |
| 2:33.142 | מת-דין-דם |  |
| 2:33.143 | נגד |  |
| 2:33.144 | ע"ז | , |
| 2:33.145 | גילוי |  |
| 2:33.146 | עריות |  |
| 2:33.147 | ושפיכות |  |
| 2:33.148 | דמים | . |
| 2:33.149 | ודוד |  |
| 2:33.150 | המלך |  |
| 2:33.151 | תיקן |  |
| 2:33.152 | אמת | , |
| 2:33.153 | לכן |  |
| 2:33.154 | לא |  |
| 2:33.155 | טעם |  |
| 2:33.156 | שינה |  |
| 2:33.157 | בלילה | , |
| 2:33.158 | ולא |  |
| 2:33.159 | הספיק |  |
| 2:33.160 | לתקן |  |
| 2:33.161 | דין |  |
| 2:33.162 | גם |  |
| 2:33.163 | דם | , |
| 2:33.164 | כמו |  |
| 2:33.165 | שאמר |  |
| 2:33.166 | הקב"ה |  |
| 2:33.167 | לא |  |
| 2:33.168 | תבנה |  |
| 2:33.169 | הבית |  |
| 2:33.170 | כי |  |
| 2:33.171 | דמים |  |
| 2:33.172 | רבים |  |
| 2:33.173 | שפכת | , |
| 2:33.174 | רק |  |
| 2:33.175 | בנך |  |
| 2:33.176 | שלמה |  |
| 2:33.177 | כי |  |
| 2:33.178 | שלום |  |
| 2:33.179 | וכו | '. |
| 2:33.180 | ותקן |  |
| 2:33.181 | שלמה |  |
| 2:33.182 | שני |  |
| 2:33.183 | אלפין |  |
| 2:33.184 | ועם |  |
| 2:33.185 | האלף |  |
| 2:33.186 | דאביו |  |
| 2:33.187 | נתקנו |  |
| 2:33.188 | ג | ' |
| 2:33.189 | אלפין | , |
| 2:33.190 | וענין |  |
| 2:33.191 | בת |  |
| 2:33.192 | ידוע |  |
| 2:33.193 | לי"ח |  |
| 2:33.194 | כי |  |
| 2:33.195 | עיקר |  |
| 2:33.196 | הפגם |  |
| 2:33.197 | שם | , |
| 2:33.198 | ועוד |  |
| 2:33.199 | האריך |  |
| 2:33.200 | בסוד |  |
| 2:33.201 | הענין |  |
| 2:33.202 | שם | . |
| 2:33.203 | וסי | ' |
| 2:33.204 | קע"ו |  |
| 2:33.205 | לענין |  |
| 2:33.206 | הזן |  |
| 2:33.207 | סומא |  |
| 2:33.208 | באחת |  |
| 2:33.209 | מעיניו | , |
| 2:33.210 | ועם |  |
| 2:33.211 | כל |  |
| 2:33.212 | זה |  |
| 2:33.213 | קראתי |  |
| 2:33.214 | תיגר |  |
| 2:33.215 | כי |  |
| 2:33.216 | נר |  |
| 2:33.217 | לי |  |
| 2:33.218 | דבתרווייהו |  |
| 2:33.219 | איכא |  |
| 2:33.220 | למיחד |  |
| 2:33.221 | מיהו | , |
| 2:33.222 | היכא |  |
| 2:33.223 | דאיכא |  |
| 2:33.224 | אחר |  |
| 2:33.225 | הגון |  |
| 2:33.226 | וראוי |  |
| 2:33.227 | כיוצא |  |
| 2:33.228 | בו |  |
| 2:33.229 | כי |  |
| 2:33.230 | ידוע |  |
| 2:33.231 | שרמ"ח |  |
| 2:33.232 | איברים |  |
| 2:33.233 | הם |  |
| 2:33.234 | כסא |  |
| 2:33.235 | ודמות |  |
| 2:33.236 | לרמ"ח |  |
| 2:33.237 | אורות |  |
| 2:33.238 | עליונים |  |
| 2:33.239 | ורמ"ח |  |
| 2:33.240 | איברים |  |
| 2:33.241 | של |  |
| 2:33.242 | נשמה | . |
| 2:33.243 | ואם |  |
| 2:33.244 | כן |  |
| 2:33.245 | כל |  |
| 2:33.246 | כי |  |
| 2:33.247 | האי |  |
| 2:33.248 | גונא |  |
| 2:33.249 | הכסא |  |
| 2:33.250 | פגום | , |
| 2:33.251 | כי |  |
| 2:33.252 | על |  |
| 2:33.253 | פי |  |
| 2:33.254 | הנסתר |  |
| 2:33.255 | יש |  |
| 2:33.256 | תלי |  |
| 2:33.257 | תלים |  |
| 2:33.258 | סודות |  |
| 2:33.259 | נסתרים |  |
| 2:33.260 | באיברי |  |
| 2:33.261 | הגוף | , |
| 2:33.262 | ויש |  |
| 2:33.263 | חסרון |  |
| 2:33.264 | שפע |  |
| 2:33.265 | בכוס |  |
| 2:33.266 | של |  |
| 2:33.267 | ברכה | . |
| 2:33.268 | העליון |  |
| 2:33.269 | ג"כ |  |
| 2:33.270 | בחתימת |  |
| 2:33.271 | הספר |  |
| 2:33.272 | כ | ' |
| 2:33.273 | שעשה |  |
| 2:33.274 | תרל"ה |  |
| 2:33.275 | תשובות |  |
| 2:33.276 | וליקט |  |
| 2:33.277 | מהן |  |
| 2:33.278 | רל"ב |  |
| 2:33.279 | מספר |  |
| 2:33.280 | יהי |  |
| 2:33.281 | אור |  |
| 2:33.282 | והוא |  |
| 2:33.283 | בגימטריא |  |
| 2:33.284 | ד | ' |
| 2:33.285 | מילוי |  |
| 2:33.286 | הויה |  |
| 2:33.287 | ע"ב |  |
| 2:33.288 | ס"ג |  |
| 2:33.289 | מ"ה |  |
| 2:33.290 | בן |  |
| 2:33.291 | שבאבא |  |
| 2:33.292 | ואימא |  |
| 2:33.293 | וזעיר |  |
| 2:33.294 | ונוקביה | . |
| 2:33.295 | ועוד |  |
| 2:33.296 | יהי |  |
| 2:33.297 | אור |  |
| 2:33.298 | יה |  |
| 2:33.299 | יאיר | . |
| 2:33.300 | שהם |  |
| 2:33.301 | מקורי |  |
| 2:33.302 | תפארת |  |
| 2:33.303 | ומלכות |  |
| 2:33.304 | שבבינה |  |
| 2:33.305 | שהוא |  |
| 2:33.306 | אור |  |
| 2:33.307 | הנאצל |  |
| 2:33.308 | ביום |  |
| 2:33.309 | ראשון |  |
| 2:33.310 | ועוד |  |
| 2:33.311 | האריך |  |
| 2:33.312 | בדברי |  |
| 2:33.313 | סוד |  |
| 2:33.314 | וסתר |  |
| 2:33.315 | הרבה | . |
| 2:33.316 | ג"כ |  |
| 2:33.317 | אביו |  |
| 2:33.318 | בעל |  |
| 2:33.319 | שמן |  |
| 2:33.320 | למאור |  |
| 2:33.321 | היה |  |
| 2:33.322 | מקובל |  |
| 2:33.323 | אלהי |  |
| 2:33.324 | שהרי |  |
| 2:33.325 | נדפסו |  |
| 2:33.326 | חרוזיו |  |
| 2:33.327 | בסוף |  |
| 2:33.328 | הספר | . |
| 2:33.329 | והן |  |
| 2:33.330 | על |  |
| 2:33.331 | פי |  |
| 2:33.332 | שם |  |
| 2:33.333 | בן |  |
| 2:33.334 | מ"ב |  |
| 2:33.335 | וי"ב |  |
| 2:33.336 | צירופי |  |
| 2:33.337 | הוי'ה |  |
| 2:33.338 | וכ"ד |  |
| 2:33.339 | צירופי |  |
| 2:33.340 | אדנות |  |
| 2:33.341 | ושם |  |
| 2:33.342 | בן |  |
| 2:33.343 | כ"ב | , |
| 2:33.344 | גם |  |
| 2:33.345 | החכם |  |
| 2:33.346 | האמתי |  |
| 2:33.347 | בעל |  |
| 2:33.348 | מנורת |  |
| 2:33.349 | המאור |  |
| 2:33.350 | היה |  |
| 2:33.351 | מקובל |  |
| 2:33.352 | אלהי |  |
| 2:33.353 | שהרי |  |
| 2:33.354 | כתב |  |
| 2:33.355 | בפ | ' |
| 2:33.356 | צ"ג |  |
| 2:33.357 | על |  |
| 2:33.358 | ברוך |  |
| 2:33.359 | שאמר |  |
| 2:33.360 | דרכו |  |
| 2:33.361 | בה |  |
| 2:33.362 | מספר |  |
| 2:33.363 | העשרה |  |
| 2:33.364 | העשרה |  |
| 2:33.365 | וכו | '. |
| 2:33.366 | ובפ | ' |
| 2:33.367 | צ"ה |  |
| 2:33.368 | כ | ' |
| 2:33.369 | אין |  |
| 2:33.370 | אומרים |  |
| 2:33.371 | קדיש |  |
| 2:33.372 | בפחות |  |
| 2:33.373 | מעשרה | , |
| 2:33.374 | והטעם |  |
| 2:33.375 | לפי |  |
| 2:33.376 | שבעשרה |  |
| 2:33.377 | נשלמה |  |
| 2:33.378 | הקדושה |  |
| 2:33.379 | שהיא |  |
| 2:33.380 | למעלה |  |
| 2:33.381 | מעשר |  |
| 2:33.382 | ספירות | , |
| 2:33.383 | ובפ | ' |
| 2:33.384 | קל"ג |  |
| 2:33.385 | כתב |  |
| 2:33.386 | וכשהיה |  |
| 2:33.387 | יעקב |  |
| 2:33.388 | יורד |  |
| 2:33.389 | למצרים |  |
| 2:33.390 | ראה |  |
| 2:33.391 | כי |  |
| 2:33.392 | הגלות |  |
| 2:33.393 | התחיל |  |
| 2:33.394 | בו |  |
| 2:33.395 | ובזרעו |  |
| 2:33.396 | ופחד |  |
| 2:33.397 | ממנו |  |
| 2:33.398 | וזבח |  |
| 2:33.399 | זבחים |  |
| 2:33.400 | לפחד |  |
| 2:33.401 | אביו |  |
| 2:33.402 | יצחק |  |
| 2:33.403 | שלא |  |
| 2:33.404 | תהא |  |
| 2:33.405 | מדת |  |
| 2:33.406 | הדין |  |
| 2:33.407 | מתוחה |  |
| 2:33.408 | כנגדו | , |
| 2:33.409 | ועשה |  |
| 2:33.410 | זה |  |
| 2:33.411 | בבאר |  |
| 2:33.412 | שבע |  |
| 2:33.413 | שהיא |  |
| 2:33.414 | בית |  |
| 2:33.415 | תפלה |  |
| 2:33.416 | לאבותיו | , |
| 2:33.417 | והקריב |  |
| 2:33.418 | שלמים |  |
| 2:33.419 | להשלים |  |
| 2:33.420 | אליו |  |
| 2:33.421 | כל |  |
| 2:33.422 | המדות |  |
| 2:33.423 | כמו |  |
| 2:33.424 | שדרשו |  |
| 2:33.425 | שלמים |  |
| 2:33.426 | שמטילים |  |
| 2:33.427 | שלום |  |
| 2:33.428 | להשלים |  |
| 2:33.429 | אליו |  |
| 2:33.430 | מדת |  |
| 2:33.431 | הגבורה | , |
| 2:33.432 | ובפ | ' |
| 2:33.433 | קמ"ז |  |
| 2:33.434 | כתב |  |
| 2:33.435 | נמצא |  |
| 2:33.436 | כי |  |
| 2:33.437 | שלמות |  |
| 2:33.438 | הקדושה |  |
| 2:33.439 | היא |  |
| 2:33.440 | בעשר |  |
| 2:33.441 | הרומזים |  |
| 2:33.442 | לעשר |  |
| 2:33.443 | הויות |  |
| 2:33.444 | והם |  |
| 2:33.445 | עשרה |  |
| 2:33.446 | מאמרות |  |
| 2:33.447 | שבהם |  |
| 2:33.448 | נברא |  |
| 2:33.449 | העולם | . |
| 2:33.450 | ובפ | ' |
| 2:33.451 | קנ"ג |  |
| 2:33.452 | כתב |  |
| 2:33.453 | ועל |  |
| 2:33.454 | דרך |  |
| 2:33.455 | האמת |  |
| 2:33.456 | כי |  |
| 2:33.457 | ששת |  |
| 2:33.458 | ימים |  |
| 2:33.459 | עשה |  |
| 2:33.460 | ה | ' |
| 2:33.461 | וכו | ' |
| 2:33.462 | וביום |  |
| 2:33.463 | השביעי |  |
| 2:33.464 | הוא |  |
| 2:33.465 | שבת |  |
| 2:33.466 | ואין |  |
| 2:33.467 | לו |  |
| 2:33.468 | בן |  |
| 2:33.469 | זוג | , |
| 2:33.470 | וכנסת |  |
| 2:33.471 | ישראל |  |
| 2:33.472 | היא |  |
| 2:33.473 | בת |  |
| 2:33.474 | זוגו |  |
| 2:33.475 | שנ | ' |
| 2:33.476 | ואת |  |
| 2:33.477 | הארץ |  |
| 2:33.478 | והנה |  |
| 2:33.479 | הוא |  |
| 2:33.480 | שמיני |  |
| 2:33.481 | עצרת | , |
| 2:33.482 | כי |  |
| 2:33.483 | שם |  |
| 2:33.484 | נעצר |  |
| 2:33.485 | הכל | . |
| 2:33.486 | וזהו |  |
| 2:33.487 | מאמרם |  |
| 2:33.488 | שמיני |  |
| 2:33.489 | רגל |  |
| 2:33.490 | בפני |  |
| 2:33.491 | עצמו |  |
| 2:33.492 | ותשלומי |  |
| 2:33.493 | ראשון |  |
| 2:33.494 | הוא |  |
| 2:33.495 | כי |  |
| 2:33.496 | הוא |  |
| 2:33.497 | אצילות |  |
| 2:33.498 | הראשונים |  |
| 2:33.499 | ואינו |  |
| 2:33.500 | באחדות |  |
| 2:33.501 | שלהם | , |
| 2:33.502 | ומי |  |
| 2:33.503 | שיבין |  |
| 2:33.504 | בהויות |  |
| 2:33.505 | ישכיל |  |
| 2:33.506 | וימצא |  |
| 2:33.507 | האמת | , |
| 2:33.508 | ובחתימת |  |
| 2:33.509 | הספר |  |
| 2:33.510 | כתב |  |
| 2:33.511 | והחמישי |  |
| 2:33.512 | עזבתי |  |
| 2:33.513 | הדברים |  |
| 2:33.514 | החמורים |  |
| 2:33.515 | שהם |  |
| 2:33.516 | נסתרים |  |
| 2:33.517 | וגבוהים |  |
| 2:33.518 | ואין |  |
| 2:33.519 | מי |  |
| 2:33.520 | שישיג |  |
| 2:33.521 | אותם |  |
| 2:33.522 | אלא |  |
| 2:33.523 | יחיד |  |
| 2:33.524 | בדור | , |
| 2:33.525 | הרי |  |
| 2:33.526 | שהיה |  |
| 2:33.527 | רוצה |  |
| 2:33.528 | לכתוב |  |
| 2:33.529 | דברי |  |
| 2:33.530 | סתר |  |
| 2:33.531 | בחבורו | . |
| 2:33.532 | אלא |  |
| 2:33.533 | מפני |  |
| 2:33.534 | שאינן |  |
| 2:33.535 | ראוים |  |
| 2:33.536 | לכל |  |
| 2:33.537 | אדם | . |
| 2:33.538 | הראית |  |
| 2:33.539 | לדעת |  |
| 2:33.540 | שאין |  |
| 2:33.541 | מי |  |
| 2:33.542 | שיחלוק |  |
| 2:33.543 | על |  |
| 2:33.544 | הקבלה | . |
| 2:33.545 | לא |  |
| 2:33.546 | הראשונים |  |
| 2:33.547 | ולא |  |
| 2:33.548 | מאחרונים |  |
| 2:33.549 | לא |  |
| 2:33.550 | רבין |  |
| 2:33.551 | ולא |  |
| 2:33.552 | יחידים | , |
| 2:33.553 | לא |  |
| 2:33.554 | גדולים |  |
| 2:33.555 | ולא |  |
| 2:33.556 | קטנים | . |
| 2:33.557 | אדרבה |  |
| 2:33.558 | הכל |  |
| 2:33.559 | מסכימים |  |
| 2:33.560 | שמי |  |
| 2:33.561 | שיזכה |  |
| 2:33.562 | לקבלה |  |
| 2:33.563 | מפה |  |
| 2:33.564 | לאזן |  |
| 2:33.565 | אין |  |
| 2:33.566 | לך |  |
| 2:33.567 | מדה |  |
| 2:33.568 | טובה |  |
| 2:33.569 | ממנה | . |
| 2:33.570 | ושאסור |  |
| 2:33.571 | לדרוש |  |
| 2:33.572 | ולחקור |  |
| 2:33.573 | עליה | , |
| 2:33.574 | אלא |  |
| 2:33.575 | אם |  |
| 2:33.576 | ירצה |  |
| 2:33.577 | יקבלה |  |
| 2:33.578 | באמונה |  |
| 2:33.579 | שלימה |  |
| 2:33.580 | ולא |  |
| 2:33.581 | יפרשה |  |
| 2:33.582 | בשכלו |  |
| 2:33.583 | האנושי | . |
| 2:33.584 | כי |  |
| 2:33.585 | יתלה |  |
| 2:33.586 | מה |  |
| 2:33.587 | שהקשה |  |
| 2:33.588 | עליו |  |
| 2:33.589 | בקוצר |  |
| 2:33.590 | דעתו |  |
| 2:33.591 | ולא |  |
| 2:33.592 | יהרוס |  |
| 2:33.593 | לחשוב |  |
| 2:33.594 | שום |  |
| 2:33.595 | הנשמה | . |
| 2:33.596 | ובזה |  |
| 2:33.597 | נשלמו |  |
| 2:33.598 | כל |  |
| 2:33.599 | דברי |  |
| 2:33.600 | שפל |  |
| 2:33.601 | השפלים | . |
| 2:33.602 | סדר |  |
| 2:33.603 | ואתה |  |
| 2:33.604 | תשוב |  |
| 2:33.605 | ושמעת |  |
| 2:33.606 | בקול |  |
| 2:33.607 | ה | '. |


### Translation under review (2:33)

English: Have you seen, to know the words of the Gaon Chavot Yair, how afraid and trembling he is to teach about separation, and relies on the words of the Gaon, the author of Shemen la-Ma'or, whose words were cited above, yet he is still afraid to say it outright, lest it slacken the hands from this holy wisdom, and how much he praises this wisdom, that one may merit it by learning it from a Kabbalist, and admits that anyone who holds himself back from studying it, after receiving it, is pushed away from its precinct; and also he permits engaging in the words of Kabbalah, and if he sees puzzling things before him, he may accept them in the faith of an artisan, and attribute it to his own shortness of understanding. If so, your two witnesses, the Ribash and Chavot Yair, will testify and bear witness against you that you invent things from your own heart; and where is his separation, which he explained, as you said? Did he not himself say, "Would that a spirit from on high were poured upon us," etc.? And I bring you proof that he merited the wisdom of Kabbalah, for in §272 he wrote, and these are his words: "that it is possible there are further matters of secret and hidden lore in two verses, alafim bat yakhil and gimel alafim bat, for at the sin of the first man three alafin were corrupted, of truth, of Adonai and of man, and there remained mat-din-dam against idolatry, illicit unions and bloodshed. And King David repaired truth; therefore he did not taste sleep at night, and did not manage to repair judgment, and also blood, as the Holy One, blessed be He, said: 'You shall not build the house, for you have shed much blood; only your son Solomon, for peace, etc.' And Solomon repaired two alafin, and with the alef of his father, three alafin were repaired, and the matter of Bat is known to li-yach, for the main flaw is there, and he further expanded on the secret of the matter there. And §276, on the matter of one who feeds a man blind in one of his eyes, and with all this I protested, for it seems to me that in both cases there is room to distinguish. However, where there is another fitting and suitable one like him, for it is known that the 248 limbs are a throne and a likeness of the 248 supernal lights, and the 248 limbs of the soul. And so, in every such case the throne is flawed, for according to the hidden there are heaps upon heaps of secret mysteries in the parts of the body, and there is a lack of flow in the cup of blessing." The one above, too, in the colophon of the book, wrote that he made 635 responses and gathered from them 232, the book Yehi Or, and it has the gematria of the four fillings of YHVH, 72, 63, 45, 52: the son who is in Abba and Imma, and Ze'ir and his Nukva. And further: let there be light. Yah, Yair. These are the sources of Tif'eret (Beauty) and Malkhut (Kingship) in Binah (Understanding), which is the light emanated on the first day. And he further expanded on many matters of secret and hidden lore. Also his father, the author of Shemen la-Ma'or, was a divine Kabbalist, for his verses were printed at the end of the book. And these are according to the Name of forty-two, and the twelve combinations of YHVH, and the twenty-four combinations of Adonai, and the Name of twenty-two. Also the true sage, the author of Menorat ha-Ma'or, was a divine Kabbalist, for in parashah 93 he wrote on "Baruch She'amar," that its way is in it, the number of the ten, the ten, etc. And in parashah 95 he wrote: "Kaddish is not said with fewer than ten," and the reason is that with ten the holiness is completed, which is above the ten sefirot. And in parashah 133 he wrote: "When Jacob was going down to Egypt, he saw that the exile had begun in him and in his seed, and he feared it and sacrificed sacrifices to the Fear of his father Isaac, so that the attribute of judgment would not be stretched out against him; and did this at Beersheba, which is a house of prayer for his fathers, and offered peace-offerings to complete to him all the attributes, as they expounded: 'shelamim, which cast peace', to complete to him the attribute of Gevurah (Might)." And in parashah 147 he wrote: "It is found that the completeness of holiness is in the ten, which allude to the ten havayot, and they are the ten sayings by which the world was created." And in parashah 153 he wrote: "And in the way of truth, that in six days the LORD made, etc. and on the seventh day it is Sabbath, and it has no mate, and the Assembly of Israel is its partner, as it is said: 'and the earth, and behold,' it is the Eighth Day of Assembly, for there everything is held back." And this is their saying: "The Eighth Day is a festival in itself," and the completion of the first is that it is the emanation of the earlier ones, and it is not in their unity. and whoever understands the havayot will be wise and find the truth. And in the colophon of the book he wrote: "And the fifth, I left out the grave matters, which are hidden and lofty, and there is no one who can attain them but a single individual in a generation." Thus he wished to write words of secret in his work, but because they are not fit for every person. Have you seen, to know that no one disputes the Kabbalah. neither the early ones nor the later ones, neither the many nor the individuals, neither the great nor the small. On the contrary, all agree that whoever merits the tradition from mouth to ear, there is no finer virtue for you than it. and that it is forbidden to expound and investigate it; but if he wishes, he should accept it in complete faith and not explain it by his human intellect. For he should attribute what troubles him to the shortness of his understanding, and not ruin any soul by thinking. And with this, all the words of the lowest of the low are complete. Order. And you shall return and hearken to the voice of the LORD.

Literal: have you seen to know the words of the Gaon Chavot Yair how he afraid and trembling to teach about separation and relies on the words of the Gaon the author of Shemen la-Ma'or were brought his words above yet he afraid to say it outright lest it cause slackening of the hands from engaging the wisdom the holy this and how much he praises the wisdom this to merit it to learn it from the mouth of a Kabbalist and admits that anyone who holds back himself from studying it after that he received it he is pushed away from its precinct and also also he permits to engage in the words of Kabbalah and if he sees before him things puzzling that he accept them in the faith of an artisan and attribute it to the shortness of his understanding if so two your witnesses the Ribash and Chavot Yair will testify and will bear witness against you that you invent things from your heart and where is it his separation that he explained as you said did not he himself say would that be given and be poured upon us a spirit from on high etc. and I bring to you proof that he merited the wisdom of Kabbalah for behold in section 272 he wrote and these are his words that it is possible that there are further matters of secret and hidden in two verses alafim bat yakhil bat yakhil and gimel alafim bat for at the sin of Adam the first were corrupted three alafin of truth and of Adonai and of man and there remained mat-din-dam against idolatry uncovering illicit unions and shedding of blood and David the king repaired truth therefore not tasted sleep at night and not manage to repair judgment also blood as that said the Holy One, blessed be He not you shall build the house for blood much you have shed only your son Solomon for peace etc. and repaired Solomon two alafin and with the alef of his father were repaired three alafin and the matter of Bat is known to li-yach for the main flaw there and further expanded on the secret of the matter there and section 276 on the matter of the one who feeds a blind man in one (of) his eyes and with all this I protested a challenge for it seems to me that in both of them there is room to distinguish however where there is another fitting and suitable like him for it is known that the 248 limbs they are a throne and a likeness of the 248 lights supernal and the 248 limbs of the soul and if so every that such case the throne flawed for according to the mouth of the hidden there are heaps upon heaps secrets hidden in the limbs of the body and there is a lack of flow in the cup of blessing the one above also in the colophon of the book wrote that he made 635 responses and gathered from them 232 the book Yehi Or and it is of the gematria four fillings of YHVH 72 63 45 52, the son who is in Abba and Imma and Ze'ir and his Nukva and further let there be light Yah Yair these are the sources of Tif'eret (Beauty) and Malkhut (Kingship) in Binah (Understanding) which is the light emanated on the day first and further expanded on the words of secret and hidden much also his father the author of Shemen la-Ma'or was a Kabbalist divine for were printed his verses at the end of the book and these according to the mouth of the Name of forty-two and twelve combinations of YHVH and twenty-four combinations of Adonai and the Name of twenty-two also the sage the true the author of Menorat ha-Ma'or was a Kabbalist divine for he wrote in parashah 93 on Baruch She-amar its way in it the number of the ten the ten etc. and in parashah 95 wrote not they say Kaddish with fewer than ten and the reason is that in ten is completed the holiness which is above the ten sefirot and in parashah 133 he wrote when Jacob was going down to Egypt he saw that the exile had begun in him and in his seed and feared it and sacrificed sacrifices to the Fear of his father Isaac so that not would be the attribute of judgment stretched against him and did this at Beersheba Beersheba which is a house of prayer for his fathers and offered peace-offerings to complete to him all the attributes as they expounded shelamim which cast peace to complete to him the attribute of Gevurah (Might) and in parashah 147 he wrote it is found that the completeness of holiness is in the ten which allude to the ten havayot and they are ten sayings by which was created the world and in parashah 153 he wrote and in the way of truth that six days made the LORD etc. and on the day the seventh it is Sabbath and there is no for it a mate and the Assembly of Israel is the partner of its as it is said and the earth and behold it is the eighth of Assembly for there is held back everything and this is their saying the eighth festival in itself and the completion of the first is that it is emanation of the earlier ones and it is not in the unity of theirs and whoever understands the havayot will be wise and will find the truth and in the colophon of the book he wrote and the fifth I left out the matters the grave which are hidden and lofty and there is no one who can attain them but a single individual in a generation thus that he was wishing to write words of secret in his work but because of that they are not fit for every person have you seen to know that there is no one who disputes about the Kabbalah neither the early ones nor the later ones neither the many nor the individuals neither the great nor the small on the contrary all agree that whoever merits the tradition from mouth to ear there is no for you virtue finer than it and that it is forbidden to expound and investigate it but if he wishes he should accept it in faith complete and not explain it by his intellect human for he should attribute what he found difficult for him to the shortness of his understanding and not he should ruin by thinking any the soul and with this are complete all the words of the lowest of the lowly order and you shall return and hearken to the voice of the LORD

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:33.1 | have you seen |  |  | interrogative ha- on a perfect verb; the opening question |
| 2:33.2 | to know |  |  |  |
| 2:33.3 | the words of |  |  |  |
| 2:33.4 | the Gaon |  |  |  |
| 2:33.5 | Chavot |  |  | Chavot Yair ('cords of Yair') is the title of the responsa of Rabbi Yair Chayyim Bacharach, called here 'the Gaon' |
| 2:33.6 | Yair |  |  |  |
| 2:33.7 | how |  |  |  |
| 2:33.8 | he |  |  |  |
| 2:33.9 | afraid |  |  |  |
| 2:33.10 | and trembling |  |  |  |
| 2:33.11 | to teach |  |  |  |
| 2:33.12 | about |  |  |  |
| 2:33.13 | separation |  |  | the separation (abstention) from this wisdom |
| 2:33.14 | and relies |  |  |  |
| 2:33.15 | on the words of |  |  |  |
| 2:33.16 | the Gaon |  |  |  |
| 2:33.17 | the author of |  |  |  |
| 2:33.18 | Shemen |  |  | Shemen la-Ma'or ('Oil for the Light') is a book; the text names its author again below as 'his father' |
| 2:33.19 | la-Ma'or |  |  |  |
| 2:33.20 | were brought |  |  |  |
| 2:33.21 | his words |  |  |  |
| 2:33.22 | above |  |  | refers back to earlier in the book |
| 2:33.23 | yet |  |  |  |
| 2:33.24 | he |  |  |  |
| 2:33.25 | afraid |  |  |  |
| 2:33.26 | to say it |  |  |  |
| 2:33.27 | outright |  |  | openly |
| 2:33.28 | lest |  |  |  |
| 2:33.29 | it cause |  |  |  |
| 2:33.30 | slackening |  |  |  |
| 2:33.31 | of the hands |  |  |  |
| 2:33.32 | from engaging |  |  |  |
| 2:33.33 | the wisdom |  |  |  |
| 2:33.34 | the holy |  |  |  |
| 2:33.35 | this |  |  |  |
| 2:33.36 | and how much |  |  |  |
| 2:33.37 | he |  |  |  |
| 2:33.38 | praises |  |  |  |
| 2:33.39 | the wisdom |  |  |  |
| 2:33.40 | this |  |  |  |
| 2:33.41 | to merit |  |  | the spelling is as printed; the sense is 'to merit' (lizkot) |
| 2:33.42 | it |  |  |  |
| 2:33.43 | to learn it |  |  |  |
| 2:33.44 | from the mouth of |  |  |  |
| 2:33.45 | a Kabbalist |  |  | mekubbal: a master of the Kabbalah |
| 2:33.46 | and admits |  |  |  |
| 2:33.47 | that anyone |  |  |  |
| 2:33.48 | who holds back |  |  |  |
| 2:33.49 | himself |  |  |  |
| 2:33.50 | from studying |  |  |  |
| 2:33.51 | it |  |  |  |
| 2:33.52 | after |  |  |  |
| 2:33.53 | that he received it |  |  |  |
| 2:33.54 | he |  |  |  |
| 2:33.55 | is pushed away |  |  |  |
| 2:33.56 | from its precinct |  |  |  |
| 2:33.57 | and also |  |  |  |
| 2:33.58 | also |  |  |  |
| 2:33.59 | he |  |  |  |
| 2:33.60 | permits |  |  |  |
| 2:33.61 | to engage |  |  |  |
| 2:33.62 | in the words of |  |  |  |
| 2:33.63 | Kabbalah |  |  |  |
| 2:33.64 | and if |  |  |  |
| 2:33.65 | he sees |  |  |  |
| 2:33.66 | before him |  |  |  |
| 2:33.67 | things |  |  |  |
| 2:33.68 | puzzling |  |  | astonishing |
| 2:33.69 | that he accept them |  |  |  |
| 2:33.70 | in the faith of |  |  |  |
| 2:33.71 | an artisan |  |  | emunat omen: the trust a layman places in an expert craftsman |
| 2:33.72 | and attribute it |  |  |  |
| 2:33.73 | to the shortness |  |  |  |
| 2:33.74 | of his understanding |  |  | his: the reader's |
| 2:33.75 | if so |  |  |  |
| 2:33.76 | two |  |  |  |
| 2:33.77 | your witnesses |  |  |  |
| 2:33.78 | the Ribash |  |  | Rabbi Isaac ben Sheshet, the Ribash |
| 2:33.79 | and Chavot |  |  |  |
| 2:33.80 | Yair |  |  |  |
| 2:33.81 | will testify |  |  |  |
| 2:33.82 | and will bear witness |  |  |  |
| 2:33.83 | against you |  |  |  |
| 2:33.84 | that you |  |  |  |
| 2:33.85 | invent |  |  |  |
| 2:33.86 | things |  |  |  |
| 2:33.87 | from your heart |  |  |  |
| 2:33.88 | and where |  |  |  |
| 2:33.89 | is it |  |  |  |
| 2:33.90 | his separation |  |  |  |
| 2:33.91 | that he explained |  |  |  |
| 2:33.92 | as |  |  |  |
| 2:33.93 | you said |  |  |  |
| 2:33.94 | did not |  |  |  |
| 2:33.95 | he |  |  |  |
| 2:33.96 | himself |  |  |  |
| 2:33.97 | say |  |  |  |
| 2:33.98 | would that |  |  |  |
| 2:33.99 | be given |  |  |  |
| 2:33.100 | and be poured |  |  |  |
| 2:33.101 | upon us |  |  |  |
| 2:33.102 | a spirit |  |  |  |
| 2:33.103 | from on high |  |  |  |
| 2:33.104 | etc. |  |  | the author's truncated quotation; a possible echo of Isa 32:15 (not verified) |
| 2:33.105 | and I |  |  |  |
| 2:33.106 | bring |  |  |  |
| 2:33.107 | to you |  |  |  |
| 2:33.108 | proof |  |  |  |
| 2:33.109 | that he merited |  |  |  |
| 2:33.110 | the wisdom of |  |  |  |
| 2:33.111 | Kabbalah |  |  |  |
| 2:33.112 | for behold |  |  |  |
| 2:33.113 | in section |  |  |  |
| 2:33.114 | 272 |  |  |  |
| 2:33.115 | he wrote |  |  |  |
| 2:33.116 | and these are his words |  |  | quotation formula; the quotation runs to 267 |
| 2:33.117 | that it is possible |  |  |  |
| 2:33.118 | that there are |  |  |  |
| 2:33.119 | further |  |  |  |
| 2:33.120 | matters of |  |  |  |
| 2:33.121 | secret |  |  |  |
| 2:33.122 | and hidden |  |  |  |
| 2:33.123 | in two |  |  |  |
| 2:33.124 | verses |  |  |  |
| 2:33.125 | alafim bat yakhil |  |  | Kabbalistic designation, not identified; kept transliterated |
| 2:33.126 | bat |  |  | Kabbalistic term; kept transliterated |
| 2:33.127 | yakhil |  |  |  |
| 2:33.128 | and gimel |  |  | the letter gimel, the number three |
| 2:33.129 | alafim |  |  | Kabbalistic 'thousands' |
| 2:33.130 | bat |  |  |  |
| 2:33.131 | for |  |  |  |
| 2:33.132 | at the sin of |  |  |  |
| 2:33.133 | Adam |  |  |  |
| 2:33.134 | the first |  |  |  |
| 2:33.135 | were corrupted |  |  |  |
| 2:33.136 | three |  |  |  |
| 2:33.137 | alafin |  |  | Kabbalistic term |
| 2:33.138 | of truth |  |  |  |
| 2:33.139 | and of Adonai |  |  | the name; the Hebrew word is adnut, literally 'lordship' |
| 2:33.140 | and of man |  |  |  |
| 2:33.141 | and there remained |  |  |  |
| 2:33.142 | mat-din-dam |  |  | the hyphenated Hebrew kept as printed |
| 2:33.143 | against |  |  |  |
| 2:33.144 | idolatry |  |  | literally 'strange worship' (avodah zarah) |
| 2:33.145 | uncovering |  |  |  |
| 2:33.146 | illicit unions |  |  | literally 'nakedness' |
| 2:33.147 | and shedding |  |  |  |
| 2:33.148 | of blood |  |  |  |
| 2:33.149 | and David |  |  |  |
| 2:33.150 | the king |  |  |  |
| 2:33.151 | repaired |  |  |  |
| 2:33.152 | truth |  |  |  |
| 2:33.153 | therefore |  |  |  |
| 2:33.154 | not |  |  |  |
| 2:33.155 | tasted |  |  |  |
| 2:33.156 | sleep |  |  |  |
| 2:33.157 | at night |  |  |  |
| 2:33.158 | and not |  |  |  |
| 2:33.159 | manage |  |  |  |
| 2:33.160 | to repair |  |  |  |
| 2:33.161 | judgment |  |  |  |
| 2:33.162 | also |  |  |  |
| 2:33.163 | blood |  |  |  |
| 2:33.164 | as |  |  |  |
| 2:33.165 | that said |  |  |  |
| 2:33.166 | the Holy One, blessed be He |  |  |  |
| 2:33.167 | not |  |  |  |
| 2:33.168 | you shall build |  |  |  |
| 2:33.169 | the house |  |  |  |
| 2:33.170 | for |  |  |  |
| 2:33.171 | blood |  |  |  |
| 2:33.172 | much |  |  |  |
| 2:33.173 | you have shed |  |  |  |
| 2:33.174 | only |  |  |  |
| 2:33.175 | your son |  |  |  |
| 2:33.176 | Solomon |  |  |  |
| 2:33.177 | for |  |  |  |
| 2:33.178 | peace |  |  |  |
| 2:33.179 | etc. |  |  | the author's wording; compare 1 Chr 22:8–9 |
| 2:33.180 | and repaired |  |  |  |
| 2:33.181 | Solomon |  |  |  |
| 2:33.182 | two |  |  |  |
| 2:33.183 | alafin |  |  |  |
| 2:33.184 | and with |  |  |  |
| 2:33.185 | the alef |  |  |  |
| 2:33.186 | of his father |  |  |  |
| 2:33.187 | were repaired |  |  |  |
| 2:33.188 | three |  |  |  |
| 2:33.189 | alafin |  |  |  |
| 2:33.190 | and the matter of |  |  |  |
| 2:33.191 | Bat |  |  |  |
| 2:33.192 | is known |  |  |  |
| 2:33.193 | to li-yach |  |  | abbreviation not resolved (perhaps yod-chet, eighteen); rendered as printed |
| 2:33.194 | for |  |  |  |
| 2:33.195 | the main |  |  |  |
| 2:33.196 | flaw |  |  |  |
| 2:33.197 | there |  |  |  |
| 2:33.198 | and further |  |  |  |
| 2:33.199 | expanded |  |  |  |
| 2:33.200 | on the secret |  |  |  |
| 2:33.201 | of the matter |  |  |  |
| 2:33.202 | there |  |  |  |
| 2:33.203 | and section |  |  |  |
| 2:33.204 | 276 |  |  | this sentence has no verb in the Hebrew; the English supplies none |
| 2:33.205 | on the matter of |  |  |  |
| 2:33.206 | the one who feeds |  |  | a halakhic case; the Hebrew is compressed |
| 2:33.207 | a blind man |  |  |  |
| 2:33.208 | in one (of) |  |  |  |
| 2:33.209 | his eyes |  |  |  |
| 2:33.210 | and with |  |  |  |
| 2:33.211 | all |  |  |  |
| 2:33.212 | this |  |  |  |
| 2:33.213 | I protested |  |  | literally 'I called out a challenge'; the 'I' is the author of the quotation, as the text runs |
| 2:33.214 | a challenge |  |  |  |
| 2:33.215 | for |  |  |  |
| 2:33.216 | it seems |  |  |  |
| 2:33.217 | to me |  |  |  |
| 2:33.218 | that in both of them |  |  |  |
| 2:33.219 | there is |  |  |  |
| 2:33.220 | room to distinguish |  |  |  |
| 2:33.221 | however |  |  |  |
| 2:33.222 | where |  |  |  |
| 2:33.223 | there is |  |  |  |
| 2:33.224 | another |  |  |  |
| 2:33.225 | fitting |  |  |  |
| 2:33.226 | and suitable |  |  |  |
| 2:33.227 | like |  |  |  |
| 2:33.228 | him |  |  |  |
| 2:33.229 | for |  |  |  |
| 2:33.230 | it is known |  |  |  |
| 2:33.231 | that the 248 |  |  | ramah: the number of the limbs (248) |
| 2:33.232 | limbs |  |  |  |
| 2:33.233 | they are |  |  |  |
| 2:33.234 | a throne |  |  | kisse: a seat |
| 2:33.235 | and a likeness |  |  |  |
| 2:33.236 | of the 248 |  |  |  |
| 2:33.237 | lights |  |  |  |
| 2:33.238 | supernal |  |  |  |
| 2:33.239 | and the 248 |  |  |  |
| 2:33.240 | limbs |  |  |  |
| 2:33.241 | of |  |  |  |
| 2:33.242 | the soul |  |  |  |
| 2:33.243 | and if |  |  | a conditional that ends in a main clause |
| 2:33.244 | so |  |  |  |
| 2:33.245 | every |  |  |  |
| 2:33.246 | that |  |  |  |
| 2:33.247 | such |  |  |  |
| 2:33.248 | case |  |  |  |
| 2:33.249 | the throne |  |  |  |
| 2:33.250 | flawed |  |  |  |
| 2:33.251 | for |  |  |  |
| 2:33.252 | according to |  |  |  |
| 2:33.253 | the mouth of |  |  |  |
| 2:33.254 | the hidden |  |  |  |
| 2:33.255 | there are |  |  |  |
| 2:33.256 | heaps |  |  |  |
| 2:33.257 | upon heaps |  |  |  |
| 2:33.258 | secrets |  |  |  |
| 2:33.259 | hidden |  |  |  |
| 2:33.260 | in the limbs |  |  |  |
| 2:33.261 | of the body |  |  |  |
| 2:33.262 | and there is |  |  |  |
| 2:33.263 | a lack |  |  |  |
| 2:33.264 | of flow |  |  |  |
| 2:33.265 | in the cup |  |  |  |
| 2:33.266 | of |  |  |  |
| 2:33.267 | blessing |  |  |  |
| 2:33.268 | the one above |  |  | the one named before; not further identified |
| 2:33.269 | also |  |  |  |
| 2:33.270 | in the colophon of |  |  |  |
| 2:33.271 | the book |  |  |  |
| 2:33.272 | wrote |  |  | abbreviation of כתב |
| 2:33.273 | that he made |  |  |  |
| 2:33.274 | 635 |  |  | the letters tav-resh-lamed-heh |
| 2:33.275 | responses |  |  |  |
| 2:33.276 | and gathered |  |  |  |
| 2:33.277 | from them |  |  |  |
| 2:33.278 | 232 |  |  | the letters resh-lamed-bet |
| 2:33.279 | the book |  |  |  |
| 2:33.280 | Yehi |  |  | Yehi Or ('Let there be light'), a book title; its letters total 232 |
| 2:33.281 | Or |  |  |  |
| 2:33.282 | and it is |  |  |  |
| 2:33.283 | of the gematria |  |  |  |
| 2:33.284 | four |  |  |  |
| 2:33.285 | fillings |  |  |  |
| 2:33.286 | of YHVH |  |  | the four fillings of the Name; the gematria names are AB, SaG, MaH, BaN |
| 2:33.287 | 72 |  |  | AB |
| 2:33.288 | 63 |  |  | SaG |
| 2:33.289 | 45 |  |  | MaH |
| 2:33.290 | 52, the son |  |  | BaN |
| 2:33.291 | who is in Abba |  |  | Abba (Father), a partzuf |
| 2:33.292 | and Imma |  |  | Imma (Mother), a partzuf |
| 2:33.293 | and Ze'ir |  |  | Ze'ir Anpin (Small Face) |
| 2:33.294 | and his Nukva |  |  | his Nukva (Female) |
| 2:33.295 | and further |  |  |  |
| 2:33.296 | let there be |  |  |  |
| 2:33.297 | light |  |  | Gen 1:3, an allusion without quotation marks |
| 2:33.298 | Yah |  |  | a divine name; with Yair, a play on the author's name |
| 2:33.299 | Yair |  |  |  |
| 2:33.300 | these are |  |  | a sentence opening after a full stop; 'these' supplies the main clause the Hebrew's syntax requires |
| 2:33.301 | the sources of |  |  |  |
| 2:33.302 | Tif'eret (Beauty) |  |  | a sefirah |
| 2:33.303 | and Malkhut (Kingship) |  |  | a sefirah |
| 2:33.304 | in Binah (Understanding) |  |  | a sefirah |
| 2:33.305 | which is |  |  |  |
| 2:33.306 | the light |  |  |  |
| 2:33.307 | emanated |  |  |  |
| 2:33.308 | on the day |  |  |  |
| 2:33.309 | first |  |  |  |
| 2:33.310 | and further |  |  |  |
| 2:33.311 | expanded |  |  |  |
| 2:33.312 | on the words of |  |  |  |
| 2:33.313 | secret |  |  |  |
| 2:33.314 | and hidden |  |  |  |
| 2:33.315 | much |  |  |  |
| 2:33.316 | also |  |  |  |
| 2:33.317 | his father |  |  | the father of the Gaon named above |
| 2:33.318 | the author of |  |  |  |
| 2:33.319 | Shemen |  |  | Shemen la-Ma'or ('Oil for the Light') |
| 2:33.320 | la-Ma'or |  |  |  |
| 2:33.321 | was |  |  |  |
| 2:33.322 | a Kabbalist |  |  |  |
| 2:33.323 | divine |  |  | literally 'a Kabbalist of God' |
| 2:33.324 | for |  |  |  |
| 2:33.325 | were printed |  |  |  |
| 2:33.326 | his verses |  |  | his rhymed verses |
| 2:33.327 | at the end of |  |  |  |
| 2:33.328 | the book |  |  |  |
| 2:33.329 | and these |  |  |  |
| 2:33.330 | according to |  |  |  |
| 2:33.331 | the mouth of |  |  |  |
| 2:33.332 | the Name |  |  |  |
| 2:33.333 | of |  |  |  |
| 2:33.334 | forty-two |  |  |  |
| 2:33.335 | and twelve |  |  |  |
| 2:33.336 | combinations |  |  |  |
| 2:33.337 | of YHVH |  |  | the Name written with an apostrophe |
| 2:33.338 | and twenty-four |  |  |  |
| 2:33.339 | combinations |  |  |  |
| 2:33.340 | of Adonai |  |  | the name; literally 'lordship' |
| 2:33.341 | and the Name |  |  |  |
| 2:33.342 | of |  |  |  |
| 2:33.343 | twenty-two |  |  |  |
| 2:33.344 | also |  |  |  |
| 2:33.345 | the sage |  |  |  |
| 2:33.346 | the true |  |  |  |
| 2:33.347 | the author of |  |  |  |
| 2:33.348 | Menorat |  |  | Menorat ha-Ma'or ('The Candelabrum of the Light'), a book title |
| 2:33.349 | ha-Ma'or |  |  | see the previous note |
| 2:33.350 | was |  |  |  |
| 2:33.351 | a Kabbalist |  |  |  |
| 2:33.352 | divine |  |  | literally 'a Kabbalist of God' |
| 2:33.353 | for |  |  |  |
| 2:33.354 | he wrote |  |  |  |
| 2:33.355 | in parashah |  |  | the numeral is kept as written; 'parashah' per the project's chapter-numbering ruling |
| 2:33.356 | 93 |  |  |  |
| 2:33.357 | on |  |  |  |
| 2:33.358 | Baruch |  |  | the blessing 'Baruch She-amar' |
| 2:33.359 | She-amar |  |  |  |
| 2:33.360 | its way |  |  |  |
| 2:33.361 | in it |  |  |  |
| 2:33.362 | the number of |  |  |  |
| 2:33.363 | the ten |  |  |  |
| 2:33.364 | the ten |  |  |  |
| 2:33.365 | etc. |  |  | the author's truncated quotation |
| 2:33.366 | and in parashah |  |  | parashah 95, as the author gives it |
| 2:33.367 | 95 |  |  |  |
| 2:33.368 | wrote |  |  | abbreviation of כתב |
| 2:33.369 | not |  |  |  |
| 2:33.370 | they say |  |  |  |
| 2:33.371 | Kaddish |  |  | the prayer that requires a quorum of ten |
| 2:33.372 | with fewer |  |  |  |
| 2:33.373 | than ten |  |  |  |
| 2:33.374 | and the reason |  |  |  |
| 2:33.375 | is that |  |  |  |
| 2:33.376 | in ten |  |  |  |
| 2:33.377 | is completed |  |  |  |
| 2:33.378 | the holiness |  |  |  |
| 2:33.379 | which is |  |  |  |
| 2:33.380 | above |  |  |  |
| 2:33.381 | the ten |  |  |  |
| 2:33.382 | sefirot |  |  |  |
| 2:33.383 | and in parashah |  |  | parashah 133, as the author gives it |
| 2:33.384 | 133 |  |  |  |
| 2:33.385 | he wrote |  |  |  |
| 2:33.386 | when |  |  |  |
| 2:33.387 | Jacob |  |  |  |
| 2:33.388 | was going down |  |  |  |
| 2:33.389 | to Egypt |  |  |  |
| 2:33.390 | he saw |  |  |  |
| 2:33.391 | that |  |  |  |
| 2:33.392 | the exile |  |  |  |
| 2:33.393 | had begun |  |  |  |
| 2:33.394 | in him |  |  |  |
| 2:33.395 | and in his seed |  |  |  |
| 2:33.396 | and feared |  |  |  |
| 2:33.397 | it |  |  | he feared it (literally 'feared from him') |
| 2:33.398 | and sacrificed |  |  |  |
| 2:33.399 | sacrifices |  |  |  |
| 2:33.400 | to the Fear of |  |  | Gen 31:42, 'Fear of Isaac' as a divine epithet |
| 2:33.401 | his father |  |  |  |
| 2:33.402 | Isaac |  |  |  |
| 2:33.403 | so that not |  |  |  |
| 2:33.404 | would be |  |  |  |
| 2:33.405 | the attribute of |  |  |  |
| 2:33.406 | judgment |  |  |  |
| 2:33.407 | stretched |  |  |  |
| 2:33.408 | against him |  |  |  |
| 2:33.409 | and did |  |  |  |
| 2:33.410 | this |  |  |  |
| 2:33.411 | at Beersheba |  |  | literally 'at the well' |
| 2:33.412 | Beersheba |  |  |  |
| 2:33.413 | which is |  |  |  |
| 2:33.414 | a house |  |  |  |
| 2:33.415 | of prayer |  |  |  |
| 2:33.416 | for his fathers |  |  |  |
| 2:33.417 | and offered |  |  |  |
| 2:33.418 | peace-offerings |  |  | shelamim, read with shalom (peace) |
| 2:33.419 | to complete |  |  |  |
| 2:33.420 | to him |  |  |  |
| 2:33.421 | all |  |  |  |
| 2:33.422 | the attributes |  |  |  |
| 2:33.423 | as |  |  |  |
| 2:33.424 | they expounded |  |  | the expounders, not named |
| 2:33.425 | shelamim |  |  |  |
| 2:33.426 | which cast |  |  |  |
| 2:33.427 | peace |  |  |  |
| 2:33.428 | to complete |  |  |  |
| 2:33.429 | to him |  |  |  |
| 2:33.430 | the attribute of |  |  |  |
| 2:33.431 | Gevurah (Might) |  |  | a sefirah |
| 2:33.432 | and in parashah |  |  | parashah 147, as the author gives it |
| 2:33.433 | 147 |  |  |  |
| 2:33.434 | he wrote |  |  |  |
| 2:33.435 | it is found |  |  |  |
| 2:33.436 | that |  |  |  |
| 2:33.437 | the completeness |  |  |  |
| 2:33.438 | of holiness |  |  |  |
| 2:33.439 | is |  |  |  |
| 2:33.440 | in the ten |  |  |  |
| 2:33.441 | which allude |  |  |  |
| 2:33.442 | to the ten |  |  |  |
| 2:33.443 | havayot |  |  | the plural of havaya, 'being'; the ten havayot |
| 2:33.444 | and they are |  |  |  |
| 2:33.445 | ten |  |  |  |
| 2:33.446 | sayings |  |  | Avot 5:1 |
| 2:33.447 | by which |  |  |  |
| 2:33.448 | was created |  |  | distinct from 'emanated' |
| 2:33.449 | the world |  |  |  |
| 2:33.450 | and in parashah |  |  | parashah 153, as the author gives it |
| 2:33.451 | 153 |  |  |  |
| 2:33.452 | he wrote |  |  |  |
| 2:33.453 | and in |  |  |  |
| 2:33.454 | the way of |  |  |  |
| 2:33.455 | truth |  |  |  |
| 2:33.456 | that |  |  |  |
| 2:33.457 | six |  |  |  |
| 2:33.458 | days |  |  |  |
| 2:33.459 | made |  |  |  |
| 2:33.460 | the LORD |  |  | the Tetragrammaton as written |
| 2:33.461 | etc. |  |  | Exod 20:11, cut short |
| 2:33.462 | and on the day |  |  |  |
| 2:33.463 | the seventh |  |  |  |
| 2:33.464 | it is |  |  |  |
| 2:33.465 | Sabbath |  |  |  |
| 2:33.466 | and there is no |  |  |  |
| 2:33.467 | for it |  |  |  |
| 2:33.468 | a |  |  |  |
| 2:33.469 | mate |  |  |  |
| 2:33.470 | and the Assembly of |  |  |  |
| 2:33.471 | Israel |  |  |  |
| 2:33.472 | is |  |  |  |
| 2:33.473 | the partner |  |  |  |
| 2:33.474 | of its |  |  |  |
| 2:33.475 | as it is said |  |  | שנאמר |
| 2:33.476 | and the |  |  | the verse is not identified; the words are cut short and are given as printed |
| 2:33.477 | earth |  |  |  |
| 2:33.478 | and behold |  |  |  |
| 2:33.479 | it is |  |  |  |
| 2:33.480 | the eighth |  |  |  |
| 2:33.481 | of Assembly |  |  | Shemini Atzeret, the Eighth Day of Assembly |
| 2:33.482 | for |  |  |  |
| 2:33.483 | there |  |  |  |
| 2:33.484 | is held back |  |  |  |
| 2:33.485 | everything |  |  |  |
| 2:33.486 | and this is |  |  |  |
| 2:33.487 | their saying |  |  | the Talmudic formula |
| 2:33.488 | the eighth |  |  |  |
| 2:33.489 | festival |  |  |  |
| 2:33.490 | in |  |  |  |
| 2:33.491 | itself |  |  |  |
| 2:33.492 | and the completion of |  |  |  |
| 2:33.493 | the first |  |  |  |
| 2:33.494 | is |  |  |  |
| 2:33.495 | that |  |  |  |
| 2:33.496 | it is |  |  |  |
| 2:33.497 | emanation |  |  |  |
| 2:33.498 | of the earlier ones |  |  |  |
| 2:33.499 | and it is not |  |  |  |
| 2:33.500 | in the unity |  |  |  |
| 2:33.501 | of theirs |  |  |  |
| 2:33.502 | and whoever |  |  |  |
| 2:33.503 | understands |  |  |  |
| 2:33.504 | the havayot |  |  |  |
| 2:33.505 | will be wise |  |  |  |
| 2:33.506 | and will find |  |  |  |
| 2:33.507 | the truth |  |  |  |
| 2:33.508 | and in the colophon of |  |  |  |
| 2:33.509 | the book |  |  |  |
| 2:33.510 | he wrote |  |  |  |
| 2:33.511 | and the fifth |  |  | the fifth part; not identified in the text |
| 2:33.512 | I left out |  |  |  |
| 2:33.513 | the matters |  |  |  |
| 2:33.514 | the grave |  |  |  |
| 2:33.515 | which are |  |  |  |
| 2:33.516 | hidden |  |  |  |
| 2:33.517 | and lofty |  |  |  |
| 2:33.518 | and there is no |  |  |  |
| 2:33.519 | one |  |  |  |
| 2:33.520 | who can attain |  |  |  |
| 2:33.521 | them |  |  |  |
| 2:33.522 | but |  |  |  |
| 2:33.523 | a single individual |  |  |  |
| 2:33.524 | in a generation |  |  |  |
| 2:33.525 | thus |  |  |  |
| 2:33.526 | that he was |  |  |  |
| 2:33.527 | wishing |  |  |  |
| 2:33.528 | to write |  |  |  |
| 2:33.529 | words of |  |  |  |
| 2:33.530 | secret |  |  |  |
| 2:33.531 | in his work |  |  |  |
| 2:33.532 | but |  |  |  |
| 2:33.533 | because of |  |  |  |
| 2:33.534 | that they are not |  |  |  |
| 2:33.535 | fit |  |  |  |
| 2:33.536 | for every |  |  |  |
| 2:33.537 | person |  |  |  |
| 2:33.538 | have you seen |  |  | as in the opening sentence |
| 2:33.539 | to know |  |  |  |
| 2:33.540 | that there is no |  |  |  |
| 2:33.541 | one |  |  |  |
| 2:33.542 | who disputes |  |  |  |
| 2:33.543 | about |  |  |  |
| 2:33.544 | the Kabbalah |  |  |  |
| 2:33.545 | neither |  |  |  |
| 2:33.546 | the early ones |  |  |  |
| 2:33.547 | nor |  |  |  |
| 2:33.548 | the later ones |  |  |  |
| 2:33.549 | neither |  |  |  |
| 2:33.550 | the many |  |  |  |
| 2:33.551 | nor |  |  |  |
| 2:33.552 | the individuals |  |  |  |
| 2:33.553 | neither |  |  |  |
| 2:33.554 | the great |  |  |  |
| 2:33.555 | nor |  |  |  |
| 2:33.556 | the small |  |  |  |
| 2:33.557 | on the contrary |  |  |  |
| 2:33.558 | all |  |  |  |
| 2:33.559 | agree |  |  |  |
| 2:33.560 | that whoever |  |  |  |
| 2:33.561 | merits |  |  |  |
| 2:33.562 | the tradition |  |  | the word is Kabbalah; here the transmitted tradition |
| 2:33.563 | from mouth |  |  |  |
| 2:33.564 | to ear |  |  |  |
| 2:33.565 | there is no |  |  |  |
| 2:33.566 | for you |  |  |  |
| 2:33.567 | virtue |  |  | literally 'measure' |
| 2:33.568 | finer |  |  | good |
| 2:33.569 | than it |  |  |  |
| 2:33.570 | and that it is forbidden |  |  |  |
| 2:33.571 | to expound |  |  |  |
| 2:33.572 | and investigate |  |  |  |
| 2:33.573 | it |  |  |  |
| 2:33.574 | but |  |  |  |
| 2:33.575 | if |  |  |  |
| 2:33.576 | he wishes |  |  |  |
| 2:33.577 | he should accept it |  |  |  |
| 2:33.578 | in faith |  |  |  |
| 2:33.579 | complete |  |  |  |
| 2:33.580 | and not |  |  |  |
| 2:33.581 | explain it |  |  |  |
| 2:33.582 | by his intellect |  |  |  |
| 2:33.583 | human |  |  |  |
| 2:33.584 | for |  |  |  |
| 2:33.585 | he should attribute |  |  |  |
| 2:33.586 | what |  |  |  |
| 2:33.587 | he found difficult |  |  |  |
| 2:33.588 | for him |  |  |  |
| 2:33.589 | to the shortness |  |  |  |
| 2:33.590 | of his understanding |  |  |  |
| 2:33.591 | and not |  |  |  |
| 2:33.592 | he should ruin |  |  |  |
| 2:33.593 | by thinking |  |  | the sense of this clause is uncertain; rendered by its likeliest sense |
| 2:33.594 | any |  |  |  |
| 2:33.595 | the soul |  |  |  |
| 2:33.596 | and with this |  |  |  |
| 2:33.597 | are complete |  |  |  |
| 2:33.598 | all |  |  |  |
| 2:33.599 | the words of |  |  |  |
| 2:33.600 | the lowest |  |  |  |
| 2:33.601 | of the lowly |  |  | a self-deprecating formula |
| 2:33.602 | order |  |  | a marker whose function the text does not state |
| 2:33.603 | and you |  |  |  |
| 2:33.604 | shall return |  |  |  |
| 2:33.605 | and hearken |  |  |  |
| 2:33.606 | to the voice of |  |  |  |
| 2:33.607 | the LORD |  |  | the Tetragrammaton as written |

Phrases:
- 2:33.1 2:33.2 2:33.3 2:33.4 2:33.5 2:33.6 2:33.7 2:33.8 2:33.9 2:33.10 2:33.11 2:33.12 2:33.13 | Have you seen, to know the words of the Gaon Chavot Yair, how afraid and trembling he is to teach about separation, | have you seen to know the words of the Gaon Chavot Yair how he afraid and trembling to teach about separation
- 2:33.14 2:33.15 2:33.16 2:33.17 2:33.18 2:33.19 2:33.20 2:33.21 2:33.22 | and relies on the words of the Gaon, the author of Shemen la-Ma'or, whose words were cited above, | and relies on the words of the Gaon the author of Shemen la-Ma'or were brought his words above
- 2:33.23 2:33.24 2:33.25 2:33.26 2:33.27 2:33.28 2:33.29 2:33.30 2:33.31 2:33.32 2:33.33 2:33.34 2:33.35 | yet he is still afraid to say it outright, lest it slacken the hands from this holy wisdom, | yet he afraid to say it outright lest it cause slackening of the hands from engaging the wisdom the holy this
- 2:33.36 2:33.37 2:33.38 2:33.39 2:33.40 2:33.41 2:33.42 2:33.43 2:33.44 2:33.45 2:33.46 2:33.47 2:33.48 2:33.49 2:33.50 2:33.51 2:33.52 2:33.53 2:33.54 2:33.55 2:33.56 | and how much he praises this wisdom, that one may merit it by learning it from a Kabbalist, and admits that anyone who holds himself back from studying it, after receiving it, is pushed away from its precinct; | and how much he praises the wisdom this to merit it to learn it from the mouth of a Kabbalist and admits that anyone who holds back himself from studying it after that he received it he is pushed away from its precinct
- 2:33.57 2:33.58 2:33.59 2:33.60 2:33.61 2:33.62 2:33.63 2:33.64 2:33.65 2:33.66 2:33.67 2:33.68 2:33.69 2:33.70 2:33.71 2:33.72 2:33.73 2:33.74 | and also he permits engaging in the words of Kabbalah, and if he sees puzzling things before him, he may accept them in the faith of an artisan, and attribute it to his own shortness of understanding. | and also also he permits to engage in the words of Kabbalah and if he sees before him things puzzling that he accept them in the faith of an artisan and attribute it to the shortness of his understanding
- 2:33.75 2:33.76 2:33.77 2:33.78 2:33.79 2:33.80 2:33.81 2:33.82 2:33.83 2:33.84 2:33.85 2:33.86 2:33.87 | If so, your two witnesses, the Ribash and Chavot Yair, will testify and bear witness against you that you invent things from your own heart; | if so two your witnesses the Ribash and Chavot Yair will testify and will bear witness against you that you invent things from your heart
- 2:33.88 2:33.89 2:33.90 2:33.91 2:33.92 2:33.93 | and where is his separation, which he explained, as you said? | and where is it his separation that he explained as you said
- 2:33.94 2:33.95 2:33.96 2:33.97 2:33.98 2:33.99 2:33.100 2:33.101 2:33.102 2:33.103 2:33.104 | Did he not himself say, "Would that a spirit from on high were poured upon us," etc.? | did not he himself say would that be given and be poured upon us a spirit from on high etc.
- 2:33.105 2:33.106 2:33.107 2:33.108 2:33.109 2:33.110 2:33.111 | And I bring you proof that he merited the wisdom of Kabbalah, | and I bring to you proof that he merited the wisdom of Kabbalah
- 2:33.112 2:33.113 2:33.114 2:33.115 2:33.116 | for in §272 he wrote, and these are his words: | for behold in section 272 he wrote and these are his words
- 2:33.117 2:33.118 2:33.119 2:33.120 2:33.121 2:33.122 2:33.123 2:33.124 2:33.125 2:33.126 2:33.127 2:33.128 2:33.129 2:33.130 2:33.131 2:33.132 2:33.133 2:33.134 | "that it is possible there are further matters of secret and hidden lore in two verses, alafim bat yakhil and gimel alafim bat, for at the sin of the first man | that it is possible that there are further matters of secret and hidden in two verses alafim bat yakhil bat yakhil and gimel alafim bat for at the sin of Adam the first
- 2:33.135 2:33.136 2:33.137 2:33.138 2:33.139 2:33.140 2:33.141 2:33.142 2:33.143 2:33.144 2:33.145 2:33.146 2:33.147 2:33.148 | three alafin were corrupted, of truth, of Adonai and of man, and there remained mat-din-dam against idolatry, illicit unions and bloodshed. | were corrupted three alafin of truth and of Adonai and of man and there remained mat-din-dam against idolatry uncovering illicit unions and shedding of blood
- 2:33.149 2:33.150 2:33.151 2:33.152 2:33.153 2:33.154 2:33.155 2:33.156 2:33.157 2:33.158 2:33.159 2:33.160 2:33.161 2:33.162 2:33.163 | And King David repaired truth; therefore he did not taste sleep at night, and did not manage to repair judgment, and also blood, | and David the king repaired truth therefore not tasted sleep at night and not manage to repair judgment also blood
- 2:33.164 2:33.165 2:33.166 2:33.167 2:33.168 2:33.169 2:33.170 2:33.171 2:33.172 2:33.173 2:33.174 2:33.175 2:33.176 2:33.177 2:33.178 2:33.179 | as the Holy One, blessed be He, said: 'You shall not build the house, for you have shed much blood; only your son Solomon, for peace, etc.' | as that said the Holy One, blessed be He not you shall build the house for blood much you have shed only your son Solomon for peace etc.
- 2:33.180 2:33.181 2:33.182 2:33.183 2:33.184 2:33.185 2:33.186 2:33.187 2:33.188 2:33.189 | And Solomon repaired two alafin, and with the alef of his father, three alafin were repaired, | and repaired Solomon two alafin and with the alef of his father were repaired three alafin
- 2:33.190 2:33.191 2:33.192 2:33.193 2:33.194 2:33.195 2:33.196 2:33.197 2:33.198 2:33.199 2:33.200 2:33.201 2:33.202 | and the matter of Bat is known to li-yach, for the main flaw is there, and he further expanded on the secret of the matter there. | and the matter of Bat is known to li-yach for the main flaw there and further expanded on the secret of the matter there
- 2:33.203 2:33.204 2:33.205 2:33.206 2:33.207 2:33.208 2:33.209 | And §276, on the matter of one who feeds a man blind in one of his eyes, | and section 276 on the matter of the one who feeds a blind man in one (of) his eyes
- 2:33.210 2:33.211 2:33.212 2:33.213 2:33.214 2:33.215 2:33.216 2:33.217 2:33.218 2:33.219 2:33.220 2:33.221 | and with all this I protested, for it seems to me that in both cases there is room to distinguish. However, | and with all this I protested a challenge for it seems to me that in both of them there is room to distinguish however
- 2:33.222 2:33.223 2:33.224 2:33.225 2:33.226 2:33.227 2:33.228 | where there is another fitting and suitable one like him, | where there is another fitting and suitable like him
- 2:33.229 2:33.230 2:33.231 2:33.232 2:33.233 2:33.234 2:33.235 2:33.236 2:33.237 2:33.238 2:33.239 2:33.240 2:33.241 2:33.242 | for it is known that the 248 limbs are a throne and a likeness of the 248 supernal lights, and the 248 limbs of the soul. | for it is known that the 248 limbs they are a throne and a likeness of the 248 lights supernal and the 248 limbs of the soul
- 2:33.243 2:33.244 2:33.245 2:33.246 2:33.247 2:33.248 2:33.249 2:33.250 2:33.251 2:33.252 2:33.253 2:33.254 2:33.255 2:33.256 2:33.257 2:33.258 2:33.259 2:33.260 2:33.261 2:33.262 2:33.263 2:33.264 2:33.265 2:33.266 2:33.267 | And so, in every such case the throne is flawed, for according to the hidden there are heaps upon heaps of secret mysteries in the parts of the body, and there is a lack of flow in the cup of blessing." | and if so every that such case the throne flawed for according to the mouth of the hidden there are heaps upon heaps secrets hidden in the limbs of the body and there is a lack of flow in the cup of blessing
- 2:33.268 2:33.269 2:33.270 2:33.271 2:33.272 2:33.273 2:33.274 2:33.275 | The one above, too, in the colophon of the book, wrote that he made 635 responses | the one above also in the colophon of the book wrote that he made 635 responses
- 2:33.276 2:33.277 2:33.278 2:33.279 2:33.280 2:33.281 | and gathered from them 232, the book Yehi Or, | and gathered from them 232 the book Yehi Or
- 2:33.282 2:33.283 2:33.284 2:33.285 2:33.286 2:33.287 2:33.288 2:33.289 2:33.290 2:33.291 2:33.292 2:33.293 2:33.294 | and it has the gematria of the four fillings of YHVH, 72, 63, 45, 52: the son who is in Abba and Imma, and Ze'ir and his Nukva. | and it is of the gematria four fillings of YHVH 72 63 45 52, the son who is in Abba and Imma and Ze'ir and his Nukva
- 2:33.295 2:33.296 2:33.297 2:33.298 2:33.299 2:33.300 2:33.301 2:33.302 2:33.303 2:33.304 2:33.305 2:33.306 2:33.307 2:33.308 2:33.309 | And further: let there be light. Yah, Yair. These are the sources of Tif'eret (Beauty) and Malkhut (Kingship) in Binah (Understanding), which is the light emanated on the first day. | and further let there be light Yah Yair these are the sources of Tif'eret (Beauty) and Malkhut (Kingship) in Binah (Understanding) which is the light emanated on the day first
- 2:33.310 2:33.311 2:33.312 2:33.313 2:33.314 2:33.315 | And he further expanded on many matters of secret and hidden lore. | and further expanded on the words of secret and hidden much
- 2:33.316 2:33.317 2:33.318 2:33.319 2:33.320 2:33.321 2:33.322 2:33.323 2:33.324 2:33.325 2:33.326 2:33.327 2:33.328 | Also his father, the author of Shemen la-Ma'or, was a divine Kabbalist, for his verses were printed at the end of the book. | also his father the author of Shemen la-Ma'or was a Kabbalist divine for were printed his verses at the end of the book
- 2:33.329 2:33.330 2:33.331 2:33.332 2:33.333 2:33.334 2:33.335 2:33.336 2:33.337 2:33.338 2:33.339 2:33.340 2:33.341 2:33.342 2:33.343 | And these are according to the Name of forty-two, and the twelve combinations of YHVH, and the twenty-four combinations of Adonai, and the Name of twenty-two. | and these according to the mouth of the Name of forty-two and twelve combinations of YHVH and twenty-four combinations of Adonai and the Name of twenty-two
- 2:33.344 2:33.345 2:33.346 2:33.347 2:33.348 2:33.349 2:33.350 2:33.351 2:33.352 2:33.353 2:33.354 2:33.355 2:33.356 | Also the true sage, the author of Menorat ha-Ma'or, was a divine Kabbalist, for in parashah 93 he wrote | also the sage the true the author of Menorat ha-Ma'or was a Kabbalist divine for he wrote in parashah 93
- 2:33.357 2:33.358 2:33.359 2:33.360 2:33.361 2:33.362 2:33.363 2:33.364 2:33.365 | on "Baruch She'amar," that its way is in it, the number of the ten, the ten, etc. | on Baruch She-amar its way in it the number of the ten the ten etc.
- 2:33.366 2:33.367 2:33.368 2:33.369 2:33.370 2:33.371 2:33.372 2:33.373 2:33.374 2:33.375 2:33.376 2:33.377 2:33.378 2:33.379 2:33.380 2:33.381 2:33.382 | And in parashah 95 he wrote: "Kaddish is not said with fewer than ten," and the reason is that with ten the holiness is completed, which is above the ten sefirot. | and in parashah 95 wrote not they say Kaddish with fewer than ten and the reason is that in ten is completed the holiness which is above the ten sefirot
- 2:33.383 2:33.384 2:33.385 2:33.386 2:33.387 2:33.388 2:33.389 2:33.390 2:33.391 2:33.392 2:33.393 2:33.394 2:33.395 2:33.396 2:33.397 | And in parashah 133 he wrote: "When Jacob was going down to Egypt, he saw that the exile had begun in him and in his seed, and he feared it | and in parashah 133 he wrote when Jacob was going down to Egypt he saw that the exile had begun in him and in his seed and feared it
- 2:33.398 2:33.399 2:33.400 2:33.401 2:33.402 2:33.403 2:33.404 2:33.405 2:33.406 2:33.407 2:33.408 | and sacrificed sacrifices to the Fear of his father Isaac, so that the attribute of judgment would not be stretched out against him; | and sacrificed sacrifices to the Fear of his father Isaac so that not would be the attribute of judgment stretched against him
- 2:33.409 2:33.410 2:33.411 2:33.412 2:33.413 2:33.414 2:33.415 2:33.416 2:33.417 2:33.418 2:33.419 2:33.420 | and did this at Beersheba, which is a house of prayer for his fathers, and offered peace-offerings to complete to him | and did this at Beersheba Beersheba which is a house of prayer for his fathers and offered peace-offerings to complete to him
- 2:33.421 2:33.422 2:33.423 2:33.424 2:33.425 2:33.426 2:33.427 2:33.428 2:33.429 2:33.430 2:33.431 | all the attributes, as they expounded: 'shelamim, which cast peace', to complete to him the attribute of Gevurah (Might)." | all the attributes as they expounded shelamim which cast peace to complete to him the attribute of Gevurah (Might)
- 2:33.432 2:33.433 2:33.434 2:33.435 2:33.436 2:33.437 2:33.438 2:33.439 2:33.440 2:33.441 2:33.442 2:33.443 2:33.444 2:33.445 2:33.446 2:33.447 2:33.448 2:33.449 | And in parashah 147 he wrote: "It is found that the completeness of holiness is in the ten, which allude to the ten havayot, and they are the ten sayings by which the world was created." | and in parashah 147 he wrote it is found that the completeness of holiness is in the ten which allude to the ten havayot and they are ten sayings by which was created the world
- 2:33.450 2:33.451 2:33.452 2:33.453 2:33.454 2:33.455 2:33.456 2:33.457 2:33.458 2:33.459 2:33.460 2:33.461 | And in parashah 153 he wrote: "And in the way of truth, that in six days the LORD made, etc. | and in parashah 153 he wrote and in the way of truth that six days made the LORD etc.
- 2:33.462 2:33.463 2:33.464 2:33.465 2:33.466 2:33.467 2:33.468 2:33.469 | and on the seventh day it is Sabbath, and it has no mate, | and on the day the seventh it is Sabbath and there is no for it a mate
- 2:33.470 2:33.471 2:33.472 2:33.473 2:33.474 2:33.475 | and the Assembly of Israel is its partner, as it is said: | and the Assembly of Israel is the partner of its as it is said
- 2:33.476 2:33.477 2:33.478 2:33.479 2:33.480 2:33.481 | 'and the earth, and behold,' it is the Eighth Day of Assembly, | and the earth and behold it is the eighth of Assembly
- 2:33.482 2:33.483 2:33.484 2:33.485 | for there everything is held back." | for there is held back everything
- 2:33.486 2:33.487 2:33.488 2:33.489 2:33.490 2:33.491 2:33.492 2:33.493 2:33.494 2:33.495 | And this is their saying: "The Eighth Day is a festival in itself," and the completion of the first is that | and this is their saying the eighth festival in itself and the completion of the first is that
- 2:33.496 2:33.497 2:33.498 2:33.499 2:33.500 2:33.501 | it is the emanation of the earlier ones, and it is not in their unity. | it is emanation of the earlier ones and it is not in the unity of theirs
- 2:33.502 2:33.503 2:33.504 2:33.505 2:33.506 2:33.507 | and whoever understands the havayot will be wise and find the truth. | and whoever understands the havayot will be wise and will find the truth
- 2:33.508 2:33.509 2:33.510 2:33.511 2:33.512 2:33.513 2:33.514 2:33.515 2:33.516 2:33.517 2:33.518 2:33.519 2:33.520 2:33.521 2:33.522 2:33.523 2:33.524 | And in the colophon of the book he wrote: "And the fifth, I left out the grave matters, which are hidden and lofty, and there is no one who can attain them but a single individual in a generation." | and in the colophon of the book he wrote and the fifth I left out the matters the grave which are hidden and lofty and there is no one who can attain them but a single individual in a generation
- 2:33.525 2:33.526 2:33.527 2:33.528 2:33.529 2:33.530 2:33.531 2:33.532 2:33.533 2:33.534 2:33.535 2:33.536 2:33.537 | Thus he wished to write words of secret in his work, but because they are not fit for every person. | thus that he was wishing to write words of secret in his work but because of that they are not fit for every person
- 2:33.538 2:33.539 2:33.540 2:33.541 2:33.542 2:33.543 2:33.544 | Have you seen, to know that no one disputes the Kabbalah. | have you seen to know that there is no one who disputes about the Kabbalah
- 2:33.545 2:33.546 2:33.547 2:33.548 2:33.549 2:33.550 2:33.551 2:33.552 2:33.553 2:33.554 2:33.555 2:33.556 | neither the early ones nor the later ones, neither the many nor the individuals, neither the great nor the small. | neither the early ones nor the later ones neither the many nor the individuals neither the great nor the small
- 2:33.557 2:33.558 2:33.559 2:33.560 2:33.561 2:33.562 2:33.563 2:33.564 2:33.565 2:33.566 2:33.567 2:33.568 2:33.569 | On the contrary, all agree that whoever merits the tradition from mouth to ear, there is no finer virtue for you than it. | on the contrary all agree that whoever merits the tradition from mouth to ear there is no for you virtue finer than it
- 2:33.570 2:33.571 2:33.572 2:33.573 2:33.574 2:33.575 2:33.576 2:33.577 2:33.578 2:33.579 2:33.580 2:33.581 2:33.582 2:33.583 | and that it is forbidden to expound and investigate it; but if he wishes, he should accept it in complete faith and not explain it by his human intellect. | and that it is forbidden to expound and investigate it but if he wishes he should accept it in faith complete and not explain it by his intellect human
- 2:33.584 2:33.585 2:33.586 2:33.587 2:33.588 2:33.589 2:33.590 2:33.591 2:33.592 2:33.593 2:33.594 2:33.595 | For he should attribute what troubles him to the shortness of his understanding, and not ruin any soul by thinking. | for he should attribute what he found difficult for him to the shortness of his understanding and not he should ruin by thinking any the soul
- 2:33.596 2:33.597 2:33.598 2:33.599 2:33.600 2:33.601 | And with this, all the words of the lowest of the low are complete. | and with this are complete all the words of the lowest of the lowly
- 2:33.602 2:33.603 2:33.604 2:33.605 2:33.606 2:33.607 | Order. And you shall return and hearken to the voice of the LORD. | order and you shall return and hearken to the voice of the LORD

Grammar notes:
- The opening and the closing both begin 'Have you seen, to know' (הראית לדעת, 1 and 538): the speaker addresses the questioner with a rhetorical question.
- The quotation from §272 is opened by כתב וז"ל (115-116) and runs to 267. The Holy One's words inside it (164-179) take single marks; the quoted words of Kabbalistic figures are in double marks where they stand.
- The conditionals ואם כן (243) and א"כ (75) carry the argument: the respondent draws the consequence from the witnesses he has just named.
- The fragment at 203-209 (וסי' קע"ו) has no verb in the Hebrew; the English supplies none and keeps it a reference.
- שמיני רגל (488) is the Talmudic formula for the Eighth Day, which is Shemini Atzeret (481); the author argues from its separateness to the emanated day.

Translator's note: The paragraph is the respondent's rebuke of the questioner, in the respondent's voice. The quotations from Chavot Yair are marked as far as the Hebrew shows them; the end of the §272 quotation is inferred (see apparatus). 'קראתי תיגר' (213) is read as the quoted author's own protest.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated, so the base text is translated as it stands.
- apparatus: Quotation boundaries: the quotation from §272 is taken to run from שאפשר (117) to the cup of blessing (267); the Hebrew marks no closing formula, so the end is inferred from the sense.
- apparatus: 1 Chr 22:8-9: the words 'לא תבנה הבית כי דמים רבים שפכת, רק בנך שלמה כי שלום' follow the author's wording, which differs from the verse.
- apparatus: Gen 1:3: יהי אור (296-297), an allusion without quotation marks.
- apparatus: Gen 31:42: פחד יצחק, 'Fear of Isaac' as a divine epithet (400-402).
- apparatus: Exod 20:11: ששת ימים עשה ה' (457-460), cut short with וכו' (461).
- apparatus: Deut 30:2: ואתה תשוב ושמעת בקול ה' (603-607), in the author's wording; the verse as usually printed reads differently.
- apparatus: Avot 5:1: עשרה מאמרות (445-446).
- apparatus: Possible echo of Isa 32:15 in מי יתן ויערה עלינו רוח ממרום (98-103); not verified.
- apparatus: Numbers: 'Shemen la-Ma'or' and 'Menorat ha-Ma'or' are book titles, meanings in the token notes; Yehi Or is a title whose letters total 232.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

