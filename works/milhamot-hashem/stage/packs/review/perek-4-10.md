# Review pack: milhamot-hashem perek-4-10 (Milchamot ha-Shem 4:16–4:18)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-4-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-10`.

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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
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
- והוא שלא: "provided that" (fixed).
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 4:16

Hebrew (Wikisource): וגם על רבי עקיבה שהוא רבו של רשב"י וגם הוא הגדול מד' שנכנסו לפרדס והגיע למקום אבני שיש טהור, ונכנס בשלום ויצא בשלום ייפלא למה ריבה תלמידי חכמים שהם בשר ודם והשוה מוראם למורא שמים, ורבי' ג"כ קבעה במשנה מסכת אבות פ' ד' בשם ר"א היה לו לרבות אשתו של זעיר אנפין שהוא ה' אלהינו לפי דבריהם? או יותר היה לו לרבות אבותיו של זעיר אנפין שהם אבא ואמא?

| id | word | punctuation after |
|---|---|---|
| 4:16.1 | וגם |  |
| 4:16.2 | על |  |
| 4:16.3 | רבי |  |
| 4:16.4 | עקיבה |  |
| 4:16.5 | שהוא |  |
| 4:16.6 | רבו |  |
| 4:16.7 | של |  |
| 4:16.8 | רשב"י |  |
| 4:16.9 | וגם |  |
| 4:16.10 | הוא |  |
| 4:16.11 | הגדול |  |
| 4:16.12 | מד | ' |
| 4:16.13 | שנכנסו |  |
| 4:16.14 | לפרדס |  |
| 4:16.15 | והגיע |  |
| 4:16.16 | למקום |  |
| 4:16.17 | אבני |  |
| 4:16.18 | שיש |  |
| 4:16.19 | טהור | , |
| 4:16.20 | ונכנס |  |
| 4:16.21 | בשלום |  |
| 4:16.22 | ויצא |  |
| 4:16.23 | בשלום |  |
| 4:16.24 | ייפלא |  |
| 4:16.25 | למה |  |
| 4:16.26 | ריבה |  |
| 4:16.27 | תלמידי |  |
| 4:16.28 | חכמים |  |
| 4:16.29 | שהם |  |
| 4:16.30 | בשר |  |
| 4:16.31 | ודם |  |
| 4:16.32 | והשוה |  |
| 4:16.33 | מוראם |  |
| 4:16.34 | למורא |  |
| 4:16.35 | שמים | , |
| 4:16.36 | ורבי | ' |
| 4:16.37 | ג"כ |  |
| 4:16.38 | קבעה |  |
| 4:16.39 | במשנה |  |
| 4:16.40 | מסכת |  |
| 4:16.41 | אבות |  |
| 4:16.42 | פ | ' |
| 4:16.43 | ד | ' |
| 4:16.44 | בשם |  |
| 4:16.45 | ר"א |  |
| 4:16.46 | היה |  |
| 4:16.47 | לו |  |
| 4:16.48 | לרבות |  |
| 4:16.49 | אשתו |  |
| 4:16.50 | של |  |
| 4:16.51 | זעיר |  |
| 4:16.52 | אנפין |  |
| 4:16.53 | שהוא |  |
| 4:16.54 | ה | ' |
| 4:16.55 | אלהינו |  |
| 4:16.56 | לפי |  |
| 4:16.57 | דבריהם | ? |
| 4:16.58 | או |  |
| 4:16.59 | יותר |  |
| 4:16.60 | היה |  |
| 4:16.61 | לו |  |
| 4:16.62 | לרבות |  |
| 4:16.63 | אבותיו |  |
| 4:16.64 | של |  |
| 4:16.65 | זעיר |  |
| 4:16.66 | אנפין |  |
| 4:16.67 | שהם |  |
| 4:16.68 | אבא |  |
| 4:16.69 | ואמא | ? |


### Translation under review (4:16)

English: And also about Rabbi Akiva, who was the teacher of Rabbi Shimon bar Yochai, he too was the greatest of the four who entered the Pardes (the Orchard), and reached the place of stones of pure marble, and entered in peace and came out in peace. It is strange why he multiplied the disciples of the wise, who are flesh and blood, and equated their reverence with the reverence of Heaven. And Rabbi also established it in the Mishnah, tractate Avot, chapter 4, in the name of Rabbi Eleazar, he should have multiplied the wife of Ze'ir Anpin, who is 'the LORD our God', according to their words? Or rather he should have multiplied the fathers of Ze'ir Anpin, who are Abba and Imma?

Literal: and also about Rabbi Akiva who was the teacher of Rabbi Shimon bar Yochai and also he the greatest of the four who entered into the Pardes (the Orchard) and reached the place of stones of marble pure and entered in peace and came out in peace it is strange why he multiplied the disciples of the wise who are flesh and blood and equated their reverence with the reverence of Heaven and Rabbi also established it in the Mishnah tractate Avot chapter 4 in the name of Rabbi Eleazar he should have for him to multiply his wife of Ze'ir Anpin who is the LORD our God according to their words or rather he should have for him to multiply the fathers of Ze'ir Anpin who are Abba and Imma

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 4:16.1 | and also |  |  |  |
| 4:16.2 | about |  |  |  |
| 4:16.3 | Rabbi |  |  |  |
| 4:16.4 | Akiva |  |  |  |
| 4:16.5 | who was |  |  |  |
| 4:16.6 | the teacher |  |  |  |
| 4:16.7 | of |  |  |  |
| 4:16.8 | Rabbi Shimon bar Yochai |  |  |  |
| 4:16.9 | and also |  |  |  |
| 4:16.10 | he |  |  |  |
| 4:16.11 | the greatest |  |  |  |
| 4:16.12 | of the four |  |  |  |
| 4:16.13 | who entered |  |  |  |
| 4:16.14 | into the Pardes (the Orchard) |  |  | Pardes, literally 'orchard': the Talmudic story of four who entered it. Of the four, the author names only Akiva here. |
| 4:16.15 | and reached |  |  |  |
| 4:16.16 | the place of |  |  |  |
| 4:16.17 | stones |  |  |  |
| 4:16.18 | of marble |  |  |  |
| 4:16.19 | pure |  |  |  |
| 4:16.20 | and entered |  |  |  |
| 4:16.21 | in peace |  |  |  |
| 4:16.22 | and came out |  |  |  |
| 4:16.23 | in peace |  |  |  |
| 4:16.24 | it is strange |  |  |  |
| 4:16.25 | why |  |  |  |
| 4:16.26 | he multiplied |  |  | The Hebrew does not say who multiplied the disciples; the English keeps the question's subject open as 'he'. |
| 4:16.27 | the disciples of |  |  |  |
| 4:16.28 | the wise |  |  |  |
| 4:16.29 | who are |  |  |  |
| 4:16.30 | flesh |  |  |  |
| 4:16.31 | and blood |  |  |  |
| 4:16.32 | and equated |  |  |  |
| 4:16.33 | their reverence |  |  |  |
| 4:16.34 | with the reverence of |  |  |  |
| 4:16.35 | Heaven |  |  |  |
| 4:16.36 | and Rabbi |  |  | The bare title 'Rabbi' with no name. In a sentence about the Mishnah it is the Mishnah's redactor, Rabbi Yehudah ha-Nasi. |
| 4:16.37 | also |  |  |  |
| 4:16.38 | established it |  |  |  |
| 4:16.39 | in the Mishnah |  |  |  |
| 4:16.40 | tractate |  |  |  |
| 4:16.41 | Avot |  |  |  |
| 4:16.42 | chapter |  |  |  |
| 4:16.43 | 4 |  |  |  |
| 4:16.44 | in the name of |  |  |  |
| 4:16.45 | Rabbi Eleazar |  |  | ר"א is written without a surname. The saying the author cites (the teacher's reverence as the reverence of Heaven) is in Avot ch. 4, and he attributes it to Rabbi Eleazar. |
| 4:16.46 | he should have |  |  | היה לו לרבות: literally 'it was for him to multiply', an idiom for 'he should have'. |
| 4:16.47 | for him |  |  |  |
| 4:16.48 | to multiply |  |  |  |
| 4:16.49 | his wife |  |  |  |
| 4:16.50 | of |  |  |  |
| 4:16.51 | Ze'ir |  |  | Ze'ir Anpin: Aramaic 'small face'. Alone, 'Ze'ir' is the name. |
| 4:16.52 | Anpin |  |  |  |
| 4:16.53 | who is |  |  |  |
| 4:16.54 | the LORD |  |  | The Tetragrammaton as a name: the author gives the Kabbalists' identification of Ze'ir Anpin with 'the LORD our God' (לפי דבריהם). |
| 4:16.55 | our God |  |  |  |
| 4:16.56 | according to |  |  |  |
| 4:16.57 | their words |  |  |  |
| 4:16.58 | or |  |  |  |
| 4:16.59 | rather |  |  |  |
| 4:16.60 | he should have |  |  |  |
| 4:16.61 | for him |  |  |  |
| 4:16.62 | to multiply |  |  |  |
| 4:16.63 | the fathers |  |  |  |
| 4:16.64 | of |  |  |  |
| 4:16.65 | Ze'ir |  |  |  |
| 4:16.66 | Anpin |  |  |  |
| 4:16.67 | who are |  |  |  |
| 4:16.68 | Abba |  |  | Abba: 'Father'; Imma: 'Mother'. In the Kabbalistic scheme they are partzufim, the parents of Ze'ir Anpin. |
| 4:16.69 | and Imma |  |  |  |

Phrases:
- 4:16.1 4:16.2 4:16.3 4:16.4 4:16.5 4:16.6 4:16.7 4:16.8 | And also about Rabbi Akiva, who was the teacher of Rabbi Shimon bar Yochai, | and also about Rabbi Akiva who was the teacher of Rabbi Shimon bar Yochai
- 4:16.9 4:16.10 4:16.11 4:16.12 4:16.13 4:16.14 | he too was the greatest of the four who entered the Pardes (the Orchard), | and also he the greatest of the four who entered into the Pardes (the Orchard)
- 4:16.15 4:16.16 4:16.17 4:16.18 4:16.19 | and reached the place of stones of pure marble, | and reached the place of stones of marble pure
- 4:16.20 4:16.21 4:16.22 4:16.23 | and entered in peace and came out in peace. | and entered in peace and came out in peace
- 4:16.24 4:16.25 4:16.26 | It is strange why he multiplied | it is strange why he multiplied
- 4:16.27 4:16.28 4:16.29 4:16.30 4:16.31 | the disciples of the wise, who are flesh and blood, | the disciples of the wise who are flesh and blood
- 4:16.32 4:16.33 4:16.34 4:16.35 | and equated their reverence with the reverence of Heaven. | and equated their reverence with the reverence of Heaven
- 4:16.36 4:16.37 4:16.38 4:16.39 4:16.40 4:16.41 4:16.42 4:16.43 | And Rabbi also established it in the Mishnah, tractate Avot, chapter 4, | and Rabbi also established it in the Mishnah tractate Avot chapter 4
- 4:16.44 4:16.45 | in the name of Rabbi Eleazar, | in the name of Rabbi Eleazar
- 4:16.46 4:16.47 4:16.48 4:16.49 4:16.50 4:16.51 4:16.52 | he should have multiplied the wife of Ze'ir Anpin, | he should have for him to multiply his wife of Ze'ir Anpin
- 4:16.53 4:16.54 4:16.55 4:16.56 4:16.57 | who is 'the LORD our God', according to their words? | who is the LORD our God according to their words
- 4:16.58 4:16.59 4:16.60 4:16.61 4:16.62 4:16.63 4:16.64 4:16.65 4:16.66 4:16.67 4:16.68 4:16.69 | Or rather he should have multiplied the fathers of Ze'ir Anpin, who are Abba and Imma? | or rather he should have for him to multiply the fathers of Ze'ir Anpin who are Abba and Imma

Grammar notes:
- The question at 4:16.24-35 is introduced by ייפלא ('it is strange') and למה ('why'). Its verb has no named subject.
- קבעה (4:16.38): 'established it', the object is the saying the author has just described, which he does not restate.
- The two questions are the author's own. The second is offered as the alternative to the first, with או יותר ('or rather').
- The author's Hebrew has no verb of quotation here; the saying is cited as Mishnah, and the question is his.

Translator's note: Ze'ir Anpin and the Tetragrammaton: the author uses the LORD's name for the partzuf, as the Kabbalists do.

Commentary rulings:

## Milchamot ha-Shem 4:17

Hebrew (Wikisource): כט) ובספר הברית חלק א' מאמר כ' פרק ט"ו כתב בעל משנת חסידים בספר יושר לבב, כי שם הקב"ה המורגל בפינו. ושבשם הוי'ה אנו קורים אותו, הוא פרצוף זעיר אנפין שהוא הקב"ה, שנשמתו המסתתרת בו על ידי הפרצופים שלפנים ממנו, הוא הסבה הראשונה ממש, ולה אנו עובדים בעבודתו, הכלל העולה כי הסבה הראשונה הנקרא בפי כל המקובלים אין סוף, הוא אשר ברא הכל יש מאין והוא מקיף את היש ההוא כולו מבחוץ, ובמובחר שבו והוא הבחור שבאבות הנקרא זעיר אנפין נסתר חדר לפנים מחדר כנשמה בתוך הגוף ומחייהו. ולכך המבחר הזה והוא הזעיר אנפין השולט על כל הברואים ומנהיגם וזנם ומפרנסם, והוא אלהינו ואנחנו עמו, כי נשמותינו הם חלקו, והוא הוא המהלל בכל התשבחות הנאמרים בתורתנו הקדושה אשר הנחילנו ועל ידה כל הנסתרות גלה לנו עכ"ל. ועוד האריך בראיות מהזהר והתיקונים שהעבודה לזעיר אנפין שהוא אמצע המקשר את הכל. ועוד כתב בשם האר"י בלקוטי תנ"ך על הפסוק בטחו בה' עדי עד כי ביה ה' צור עולמים וז"ל. נמצא כשיכוין האדם בזעיר אנפין לבד מספיק, כי שם נמצא אריך אנפין ואבא ואמא מתלבשין זה תוך זה תוך זעיר אנפין עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:17.1 | כט | ) |
| 4:17.2 | ובספר |  |
| 4:17.3 | הברית |  |
| 4:17.4 | חלק |  |
| 4:17.5 | א | ' |
| 4:17.6 | מאמר |  |
| 4:17.7 | כ | ' |
| 4:17.8 | פרק |  |
| 4:17.9 | ט"ו |  |
| 4:17.10 | כתב |  |
| 4:17.11 | בעל |  |
| 4:17.12 | משנת |  |
| 4:17.13 | חסידים |  |
| 4:17.14 | בספר |  |
| 4:17.15 | יושר |  |
| 4:17.16 | לבב | , |
| 4:17.17 | כי |  |
| 4:17.18 | שם |  |
| 4:17.19 | הקב"ה |  |
| 4:17.20 | המורגל |  |
| 4:17.21 | בפינו | . |
| 4:17.22 | ושבשם |  |
| 4:17.23 | הוי'ה |  |
| 4:17.24 | אנו |  |
| 4:17.25 | קורים |  |
| 4:17.26 | אותו | , |
| 4:17.27 | הוא |  |
| 4:17.28 | פרצוף |  |
| 4:17.29 | זעיר |  |
| 4:17.30 | אנפין |  |
| 4:17.31 | שהוא |  |
| 4:17.32 | הקב"ה | , |
| 4:17.33 | שנשמתו |  |
| 4:17.34 | המסתתרת |  |
| 4:17.35 | בו |  |
| 4:17.36 | על |  |
| 4:17.37 | ידי |  |
| 4:17.38 | הפרצופים |  |
| 4:17.39 | שלפנים |  |
| 4:17.40 | ממנו | , |
| 4:17.41 | הוא |  |
| 4:17.42 | הסבה |  |
| 4:17.43 | הראשונה |  |
| 4:17.44 | ממש | , |
| 4:17.45 | ולה |  |
| 4:17.46 | אנו |  |
| 4:17.47 | עובדים |  |
| 4:17.48 | בעבודתו | , |
| 4:17.49 | הכלל |  |
| 4:17.50 | העולה |  |
| 4:17.51 | כי |  |
| 4:17.52 | הסבה |  |
| 4:17.53 | הראשונה |  |
| 4:17.54 | הנקרא |  |
| 4:17.55 | בפי |  |
| 4:17.56 | כל |  |
| 4:17.57 | המקובלים |  |
| 4:17.58 | אין |  |
| 4:17.59 | סוף | , |
| 4:17.60 | הוא |  |
| 4:17.61 | אשר |  |
| 4:17.62 | ברא |  |
| 4:17.63 | הכל |  |
| 4:17.64 | יש |  |
| 4:17.65 | מאין |  |
| 4:17.66 | והוא |  |
| 4:17.67 | מקיף |  |
| 4:17.68 | את |  |
| 4:17.69 | היש |  |
| 4:17.70 | ההוא |  |
| 4:17.71 | כולו |  |
| 4:17.72 | מבחוץ | , |
| 4:17.73 | ובמובחר |  |
| 4:17.74 | שבו |  |
| 4:17.75 | והוא |  |
| 4:17.76 | הבחור |  |
| 4:17.77 | שבאבות |  |
| 4:17.78 | הנקרא |  |
| 4:17.79 | זעיר |  |
| 4:17.80 | אנפין |  |
| 4:17.81 | נסתר |  |
| 4:17.82 | חדר |  |
| 4:17.83 | לפנים |  |
| 4:17.84 | מחדר |  |
| 4:17.85 | כנשמה |  |
| 4:17.86 | בתוך |  |
| 4:17.87 | הגוף |  |
| 4:17.88 | ומחייהו | . |
| 4:17.89 | ולכך |  |
| 4:17.90 | המבחר |  |
| 4:17.91 | הזה |  |
| 4:17.92 | והוא |  |
| 4:17.93 | הזעיר |  |
| 4:17.94 | אנפין |  |
| 4:17.95 | השולט |  |
| 4:17.96 | על |  |
| 4:17.97 | כל |  |
| 4:17.98 | הברואים |  |
| 4:17.99 | ומנהיגם |  |
| 4:17.100 | וזנם |  |
| 4:17.101 | ומפרנסם | , |
| 4:17.102 | והוא |  |
| 4:17.103 | אלהינו |  |
| 4:17.104 | ואנחנו |  |
| 4:17.105 | עמו | , |
| 4:17.106 | כי |  |
| 4:17.107 | נשמותינו |  |
| 4:17.108 | הם |  |
| 4:17.109 | חלקו | , |
| 4:17.110 | והוא |  |
| 4:17.111 | הוא |  |
| 4:17.112 | המהלל |  |
| 4:17.113 | בכל |  |
| 4:17.114 | התשבחות |  |
| 4:17.115 | הנאמרים |  |
| 4:17.116 | בתורתנו |  |
| 4:17.117 | הקדושה |  |
| 4:17.118 | אשר |  |
| 4:17.119 | הנחילנו |  |
| 4:17.120 | ועל |  |
| 4:17.121 | ידה |  |
| 4:17.122 | כל |  |
| 4:17.123 | הנסתרות |  |
| 4:17.124 | גלה |  |
| 4:17.125 | לנו |  |
| 4:17.126 | עכ"ל | . |
| 4:17.127 | ועוד |  |
| 4:17.128 | האריך |  |
| 4:17.129 | בראיות |  |
| 4:17.130 | מהזהר |  |
| 4:17.131 | והתיקונים |  |
| 4:17.132 | שהעבודה |  |
| 4:17.133 | לזעיר |  |
| 4:17.134 | אנפין |  |
| 4:17.135 | שהוא |  |
| 4:17.136 | אמצע |  |
| 4:17.137 | המקשר |  |
| 4:17.138 | את |  |
| 4:17.139 | הכל | . |
| 4:17.140 | ועוד |  |
| 4:17.141 | כתב |  |
| 4:17.142 | בשם |  |
| 4:17.143 | האר"י |  |
| 4:17.144 | בלקוטי |  |
| 4:17.145 | תנ"ך |  |
| 4:17.146 | על |  |
| 4:17.147 | הפסוק |  |
| 4:17.148 | בטחו |  |
| 4:17.149 | בה | ' |
| 4:17.150 | עדי |  |
| 4:17.151 | עד |  |
| 4:17.152 | כי |  |
| 4:17.153 | ביה |  |
| 4:17.154 | ה | ' |
| 4:17.155 | צור |  |
| 4:17.156 | עולמים |  |
| 4:17.157 | וז"ל | . |
| 4:17.158 | נמצא |  |
| 4:17.159 | כשיכוין |  |
| 4:17.160 | האדם |  |
| 4:17.161 | בזעיר |  |
| 4:17.162 | אנפין |  |
| 4:17.163 | לבד |  |
| 4:17.164 | מספיק | , |
| 4:17.165 | כי |  |
| 4:17.166 | שם |  |
| 4:17.167 | נמצא |  |
| 4:17.168 | אריך |  |
| 4:17.169 | אנפין |  |
| 4:17.170 | ואבא |  |
| 4:17.171 | ואמא |  |
| 4:17.172 | מתלבשין |  |
| 4:17.173 | זה |  |
| 4:17.174 | תוך |  |
| 4:17.175 | זה |  |
| 4:17.176 | תוך |  |
| 4:17.177 | זעיר |  |
| 4:17.178 | אנפין |  |
| 4:17.179 | עכ"ל | . |


### Translation under review (4:17)

English: §29. In Sefer ha-Berit, part 1, treatise 20, chapter 15, the author of Mishnat Chasidim wrote in the book Yosher Levav: "that the name 'the Holy One, blessed be He', which is customary in our mouths. And by the name YHVH we call it, is the partzuf of Ze'ir Anpin, which is the Holy One, blessed be He, whose soul is hidden in him through the agency of the partzufim inside him, is the First Cause itself, and to it we worship in his worship. The general result is that the First Cause, which all the Kabbalists call Ein Sof, is He who created everything out of nothing, and He encompasses that whole existent from without, and in its choicest part, and He is the choicest one of the Fathers, called Ze'ir Anpin, is hidden, chamber within chamber, as the soul in the body and gives it life. Therefore this choice that is, Ze'ir Anpin, who rules over all the creatures, directs them, sustains them, and provides for them, and He is our God, and we are His people, for our souls are His portion, and He, He is the One praised in all the praises said in our holy Torah, which He bequeathed to us, and through it He revealed all the hidden things to us." Thus far his words. And further he expanded with proofs from the Zohar, and from Tikkunei Zohar, that worship is due to Ze'ir Anpin, who is the middle that connects everything. And further he wrote in the name of the Ari in Likkutei Tanakh on the verse "Trust in the LORD forever and ever, for in Yah, the LORD, is an everlasting rock." And these are his words: "It follows that when a person intends on Ze'ir Anpin alone, it suffices, for there are found Arikh Anpin, Abba and Imma, clothed one within the other, within Ze'ir Anpin." Thus far his words.

Literal: §29 in Sefer ha-Berit part 1 treatise 20 chapter 15 wrote the author of Mishnat Chasidim in the book Yosher Levav that the name the Holy One, blessed be He customary in our mouths and by the name YHVH we call it is the partzuf of Ze'ir Anpin which is the Holy One, blessed be He whose soul hidden in him through the agency of the partzufim inside him is the First Cause first itself and to it we worship in his worship the general result is that the cause first which is called by the mouth of all the Kabbalists Ein Sof is who created everything being out of nothing and He encompasses - the existent that whole from without and in the choicest of it and He is the choicest one of the Fathers called Ze'ir Anpin hidden chamber inward from chamber as a soul in the body and gives it life therefore this choice this namely Ze'ir Anpin who rules over all the creatures directs them sustains them provides for them and He is our God and we are His people for our souls are His portion and He He the praised one in all the praises that are said in our Torah holy which bequeathed to us and through it all the hidden things revealed to us Thus far his words. and further expanded with proofs from the Zohar and the Tikkunei Zohar that worship to Ze'ir Anpin who is the middle that connects - everything and further wrote in the name of the Ari in Likkutei Tanakh on the verse trust in the LORD for ever and ever for in Yah the LORD is a rock everlasting and these are his words: it follows when one intends (the person) on Ze'ir Anpin alone it suffices for there are found Arikh Anpin and Abba and Imma clothed one within the other within Ze'ir Anpin Thus far his words.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 4:17.1 | §29 |  |  |  |
| 4:17.2 | in |  |  |  |
| 4:17.3 | Sefer ha-Berit |  |  | Sefer ha-Berit: 'the Book of the Covenant'. |
| 4:17.4 | part |  |  |  |
| 4:17.5 | 1 |  |  |  |
| 4:17.6 | treatise |  |  |  |
| 4:17.7 | 20 |  |  |  |
| 4:17.8 | chapter |  |  |  |
| 4:17.9 | 15 |  |  |  |
| 4:17.10 | wrote |  |  |  |
| 4:17.11 | the author of |  |  |  |
| 4:17.12 | Mishnat |  |  | Mishnat Chasidim: 'the teaching of the pious'. The author of the book is named only by the title. |
| 4:17.13 | Chasidim |  |  |  |
| 4:17.14 | in the book |  |  |  |
| 4:17.15 | Yosher |  |  | Yosher Levav: literally 'uprightness of heart'. |
| 4:17.16 | Levav |  |  |  |
| 4:17.17 | that |  |  |  |
| 4:17.18 | the name |  |  |  |
| 4:17.19 | the Holy One, blessed be He |  |  |  |
| 4:17.20 | customary |  |  |  |
| 4:17.21 | in our mouths |  |  |  |
| 4:17.22 | and by the name |  |  |  |
| 4:17.23 | YHVH |  |  |  |
| 4:17.24 | we |  |  |  |
| 4:17.25 | call |  |  |  |
| 4:17.26 | it |  |  |  |
| 4:17.27 | is |  |  |  |
| 4:17.28 | the partzuf |  |  | partzuf: a Kabbalistic entity, the word for 'face'. Transliterated as the Kabbalists' term. |
| 4:17.29 | of Ze'ir |  |  |  |
| 4:17.30 | Anpin |  |  |  |
| 4:17.31 | which is |  |  |  |
| 4:17.32 | the Holy One, blessed be He |  |  |  |
| 4:17.33 | whose soul |  |  |  |
| 4:17.34 | hidden |  |  |  |
| 4:17.35 | in him |  |  |  |
| 4:17.36 | through |  |  |  |
| 4:17.37 | the agency of |  |  |  |
| 4:17.38 | the partzufim |  |  |  |
| 4:17.39 | inside |  |  |  |
| 4:17.40 | him |  |  |  |
| 4:17.41 | is |  |  |  |
| 4:17.42 | the First Cause |  |  |  |
| 4:17.43 | first |  |  |  |
| 4:17.44 | itself |  |  |  |
| 4:17.45 | and to it |  |  |  |
| 4:17.46 | we |  |  |  |
| 4:17.47 | worship |  |  |  |
| 4:17.48 | in his worship |  |  |  |
| 4:17.49 | the general result |  |  |  |
| 4:17.50 | is |  |  |  |
| 4:17.51 | that |  |  |  |
| 4:17.52 | the cause |  |  |  |
| 4:17.53 | first |  |  |  |
| 4:17.54 | which is called |  |  |  |
| 4:17.55 | by the mouth of |  |  |  |
| 4:17.56 | all |  |  |  |
| 4:17.57 | the Kabbalists |  |  |  |
| 4:17.58 | Ein |  |  | Ein Sof: literally 'without end'. |
| 4:17.59 | Sof |  |  |  |
| 4:17.60 | is |  |  |  |
| 4:17.61 | who |  |  |  |
| 4:17.62 | created |  |  |  |
| 4:17.63 | everything |  |  |  |
| 4:17.64 | being |  |  |  |
| 4:17.65 | out of nothing |  |  |  |
| 4:17.66 | and He |  |  |  |
| 4:17.67 | encompasses |  |  |  |
| 4:17.68 | - |  |  |  |
| 4:17.69 | the existent |  |  |  |
| 4:17.70 | that |  |  |  |
| 4:17.71 | whole |  |  |  |
| 4:17.72 | from without |  |  |  |
| 4:17.73 | and in the choicest |  |  |  |
| 4:17.74 | of it |  |  |  |
| 4:17.75 | and He is |  |  |  |
| 4:17.76 | the choicest one |  |  |  |
| 4:17.77 | of the Fathers |  |  |  |
| 4:17.78 | called |  |  |  |
| 4:17.79 | Ze'ir |  |  |  |
| 4:17.80 | Anpin |  |  |  |
| 4:17.81 | hidden |  |  |  |
| 4:17.82 | chamber |  |  |  |
| 4:17.83 | inward |  |  |  |
| 4:17.84 | from chamber |  |  |  |
| 4:17.85 | as a soul |  |  |  |
| 4:17.86 | in |  |  |  |
| 4:17.87 | the body |  |  |  |
| 4:17.88 | and gives it life |  |  |  |
| 4:17.89 | therefore |  |  |  |
| 4:17.90 | this choice |  |  |  |
| 4:17.91 | this |  |  |  |
| 4:17.92 | namely |  |  |  |
| 4:17.93 | Ze'ir |  |  |  |
| 4:17.94 | Anpin |  |  |  |
| 4:17.95 | who rules |  |  |  |
| 4:17.96 | over |  |  |  |
| 4:17.97 | all |  |  |  |
| 4:17.98 | the creatures |  |  |  |
| 4:17.99 | directs them |  |  |  |
| 4:17.100 | sustains them |  |  |  |
| 4:17.101 | provides for them |  |  |  |
| 4:17.102 | and He |  |  |  |
| 4:17.103 | is our God |  |  |  |
| 4:17.104 | and we |  |  |  |
| 4:17.105 | are His people |  |  |  |
| 4:17.106 | for |  |  |  |
| 4:17.107 | our souls |  |  |  |
| 4:17.108 | are |  |  |  |
| 4:17.109 | His portion |  |  |  |
| 4:17.110 | and He |  |  |  |
| 4:17.111 | He |  |  |  |
| 4:17.112 | the praised one |  |  |  |
| 4:17.113 | in all |  |  |  |
| 4:17.114 | the praises |  |  |  |
| 4:17.115 | that are said |  |  |  |
| 4:17.116 | in our Torah |  |  |  |
| 4:17.117 | holy |  |  |  |
| 4:17.118 | which |  |  |  |
| 4:17.119 | bequeathed to us |  |  |  |
| 4:17.120 | and through |  |  |  |
| 4:17.121 | it |  |  |  |
| 4:17.122 | all |  |  |  |
| 4:17.123 | the hidden things |  |  |  |
| 4:17.124 | revealed |  |  |  |
| 4:17.125 | to us |  |  |  |
| 4:17.126 | Thus far his words. |  |  |  |
| 4:17.127 | and further |  |  |  |
| 4:17.128 | expanded |  |  |  |
| 4:17.129 | with proofs |  |  |  |
| 4:17.130 | from the Zohar |  |  |  |
| 4:17.131 | and the Tikkunei Zohar |  |  | Tikkunei Zohar: 'the emendations of the Zohar'. |
| 4:17.132 | that worship |  |  |  |
| 4:17.133 | to Ze'ir |  |  |  |
| 4:17.134 | Anpin |  |  |  |
| 4:17.135 | who is |  |  |  |
| 4:17.136 | the middle |  |  |  |
| 4:17.137 | that connects |  |  |  |
| 4:17.138 | - |  |  |  |
| 4:17.139 | everything |  |  |  |
| 4:17.140 | and further |  |  |  |
| 4:17.141 | wrote |  |  |  |
| 4:17.142 | in the name of |  |  |  |
| 4:17.143 | the Ari |  |  |  |
| 4:17.144 | in Likkutei |  |  | Likkutei: 'collections'; the title of the book the Ari's words are cited from. |
| 4:17.145 | Tanakh |  |  |  |
| 4:17.146 | on |  |  |  |
| 4:17.147 | the verse |  |  |  |
| 4:17.148 | trust |  |  |  |
| 4:17.149 | in the LORD |  |  |  |
| 4:17.150 | for ever |  |  |  |
| 4:17.151 | and ever |  |  |  |
| 4:17.152 | for |  |  |  |
| 4:17.153 | in Yah |  |  |  |
| 4:17.154 | the LORD |  |  |  |
| 4:17.155 | is a rock |  |  |  |
| 4:17.156 | everlasting |  |  |  |
| 4:17.157 | and these are his words: |  |  |  |
| 4:17.158 | it follows |  |  |  |
| 4:17.159 | when one intends |  |  |  |
| 4:17.160 | (the person) |  |  |  |
| 4:17.161 | on Ze'ir |  |  |  |
| 4:17.162 | Anpin |  |  |  |
| 4:17.163 | alone |  |  |  |
| 4:17.164 | it suffices |  |  |  |
| 4:17.165 | for |  |  |  |
| 4:17.166 | there |  |  |  |
| 4:17.167 | are found |  |  |  |
| 4:17.168 | Arikh |  |  | Arikh Anpin: 'Long Face'. |
| 4:17.169 | Anpin |  |  |  |
| 4:17.170 | and Abba |  |  |  |
| 4:17.171 | and Imma |  |  |  |
| 4:17.172 | clothed |  |  |  |
| 4:17.173 | one |  |  |  |
| 4:17.174 | within |  |  |  |
| 4:17.175 | the other |  |  |  |
| 4:17.176 | within |  |  |  |
| 4:17.177 | Ze'ir |  |  |  |
| 4:17.178 | Anpin |  |  |  |
| 4:17.179 | Thus far his words. |  |  |  |

Phrases:
- 4:17.1 4:17.2 4:17.3 4:17.4 4:17.5 4:17.6 4:17.7 4:17.8 4:17.9 | §29. In Sefer ha-Berit, part 1, treatise 20, chapter 15, | §29 in Sefer ha-Berit part 1 treatise 20 chapter 15
- 4:17.10 4:17.11 4:17.12 4:17.13 | the author of Mishnat Chasidim wrote | wrote the author of Mishnat Chasidim
- 4:17.14 4:17.15 4:17.16 | in the book Yosher Levav: | in the book Yosher Levav
- 4:17.17 4:17.18 4:17.19 4:17.20 4:17.21 | "that the name 'the Holy One, blessed be He', which is customary in our mouths. | that the name the Holy One, blessed be He customary in our mouths
- 4:17.22 4:17.23 4:17.24 4:17.25 4:17.26 | And by the name YHVH we call it, | and by the name YHVH we call it
- 4:17.27 4:17.28 4:17.29 4:17.30 4:17.31 | is the partzuf of Ze'ir Anpin, which is | is the partzuf of Ze'ir Anpin which is
- 4:17.32 4:17.33 4:17.34 4:17.35 | the Holy One, blessed be He, whose soul is hidden in him | the Holy One, blessed be He whose soul hidden in him
- 4:17.36 4:17.37 4:17.38 4:17.39 4:17.40 | through the agency of the partzufim inside him, | through the agency of the partzufim inside him
- 4:17.41 4:17.42 4:17.43 4:17.44 | is the First Cause itself, | is the First Cause first itself
- 4:17.45 4:17.46 4:17.47 4:17.48 | and to it we worship in his worship. | and to it we worship in his worship
- 4:17.49 4:17.50 4:17.51 | The general result is that | the general result is that
- 4:17.52 4:17.53 4:17.54 4:17.55 4:17.56 4:17.57 4:17.58 4:17.59 | the First Cause, which all the Kabbalists call Ein Sof, | the cause first which is called by the mouth of all the Kabbalists Ein Sof
- 4:17.60 4:17.61 4:17.62 4:17.63 4:17.64 4:17.65 | is He who created everything out of nothing, | is who created everything being out of nothing
- 4:17.66 4:17.67 4:17.68 4:17.69 4:17.70 4:17.71 4:17.72 | and He encompasses that whole existent from without, | and He encompasses - the existent that whole from without
- 4:17.73 4:17.74 | and in its choicest part, | and in the choicest of it
- 4:17.75 4:17.76 4:17.77 4:17.78 | and He is the choicest one of the Fathers, called | and He is the choicest one of the Fathers called
- 4:17.79 4:17.80 4:17.81 4:17.82 4:17.83 4:17.84 | Ze'ir Anpin, is hidden, chamber within chamber, | Ze'ir Anpin hidden chamber inward from chamber
- 4:17.85 4:17.86 4:17.87 4:17.88 | as the soul in the body and gives it life. | as a soul in the body and gives it life
- 4:17.89 4:17.90 4:17.91 | Therefore this choice | therefore this choice this
- 4:17.92 4:17.93 4:17.94 | that is, Ze'ir Anpin, | namely Ze'ir Anpin
- 4:17.95 4:17.96 4:17.97 4:17.98 | who rules over all the creatures, | who rules over all the creatures
- 4:17.99 4:17.100 4:17.101 | directs them, sustains them, and provides for them, | directs them sustains them provides for them
- 4:17.102 4:17.103 4:17.104 4:17.105 | and He is our God, and we are His people, | and He is our God and we are His people
- 4:17.106 4:17.107 4:17.108 4:17.109 | for our souls are His portion, | for our souls are His portion
- 4:17.110 4:17.111 4:17.112 4:17.113 4:17.114 | and He, He is the One praised in all the praises | and He He the praised one in all the praises
- 4:17.115 4:17.116 4:17.117 4:17.118 4:17.119 | said in our holy Torah, which He bequeathed to us, | that are said in our Torah holy which bequeathed to us
- 4:17.120 4:17.121 4:17.122 4:17.123 4:17.124 4:17.125 | and through it He revealed all the hidden things to us." | and through it all the hidden things revealed to us
- 4:17.126 | Thus far his words. | Thus far his words.
- 4:17.127 4:17.128 4:17.129 4:17.130 | And further he expanded with proofs from the Zohar, | and further expanded with proofs from the Zohar
- 4:17.131 4:17.132 4:17.133 4:17.134 | and from Tikkunei Zohar, that worship is due to Ze'ir Anpin, | and the Tikkunei Zohar that worship to Ze'ir Anpin
- 4:17.135 4:17.136 4:17.137 4:17.138 4:17.139 | who is the middle that connects everything. | who is the middle that connects - everything
- 4:17.140 4:17.141 4:17.142 | And further he wrote in the name of | and further wrote in the name of
- 4:17.143 4:17.144 4:17.145 4:17.146 4:17.147 | the Ari in Likkutei Tanakh on the verse | the Ari in Likkutei Tanakh on the verse
- 4:17.148 4:17.149 4:17.150 4:17.151 | "Trust in the LORD forever and ever, | trust in the LORD for ever and ever
- 4:17.152 4:17.153 4:17.154 4:17.155 4:17.156 | for in Yah, the LORD, is an everlasting rock." | for in Yah the LORD is a rock everlasting
- 4:17.157 | And these are his words: | and these are his words:
- 4:17.158 4:17.159 4:17.160 4:17.161 4:17.162 4:17.163 4:17.164 | "It follows that when a person intends on Ze'ir Anpin alone, it suffices, | it follows when one intends (the person) on Ze'ir Anpin alone it suffices
- 4:17.165 4:17.166 4:17.167 4:17.168 4:17.169 | for there are found Arikh Anpin, | for there are found Arikh Anpin
- 4:17.170 4:17.171 | Abba and Imma, | and Abba and Imma
- 4:17.172 4:17.173 4:17.174 4:17.175 4:17.176 4:17.177 4:17.178 | clothed one within the other, within Ze'ir Anpin." | clothed one within the other within Ze'ir Anpin
- 4:17.179 | Thus far his words. | Thus far his words.

Grammar notes:
- The first quotation runs from כי (4:17.17) to עכ"ל (4:17.126). It is introduced by כתב (4:17.10) and is in Hebrew.
- ולה אנו עובדים בעבודתו (4:17.45-48): ולה ('to it') is feminine, because the cause is feminine; בעבודתו ('in his worship') is masculine, for the partzuf.
- הכלל העולה כי (4:17.49-51), 'the general result is that', opens the author's summary of the quotation.
- The Ari's quotation is introduced by וז"ל (4:17.157), 'and these are his words'. It cites Isa 26:4 (4:17.148-156); ביה (4:17.153) is 'in Yah', the short form of the Name.
- אין סוף (4:17.58-59) is a name in the Kabbalists' usage, literally 'without end'; the author's own argument is about whether it is the First Cause.

Translator's note: The section label כט) (4:17.1) is the author's own number, written in Hebrew letters: 29.

Commentary rulings:
- apparatus: Isa 26:4, quoted by the Ari (4:17.148-156): 'Trust in the LORD forever and ever, for in Yah, the LORD, is an everlasting rock.'

## Milchamot ha-Shem 4:18

Hebrew (Wikisource): גם בספר עץ חיים להרח"ו שער הכללים פ' י"א כתב שמשה רבינו אמר לישראל בכניסתן לארץ ואתם הדבקים בה' אלהיכם חיים כולכם היום, כי ה' אלהיכם הוא זעיר אנפין ונוקביה, וישראל הנכנסים לארץ היו וכו' עי"ש. גם בספר הכוונות במנהגי האר"י כתב כי כל פעם שתוציא מפיך שם הוי'ה תכוין לזעיר אנפין, ותכוין בו כך יה"ו זעיר אנפין וה"א אחרונה שבו היא מלכות. תכוין שנעשה שם שלם על אדני באופן וכו' ע"ש. וכן בספר מחברת הקודש בסדר מוסף שבת כ' כי מלאכים המוני מעלה יתנו כתר לזעיר אנפין שהוא ה' אלהינו עיין שם דף מ' ע"א. גם המקדש מלך ד' י"ב ד"ה אלהא רברבא מבואר דלא נקרא אלהא רברבא אלא בזעיר אנפין, ובחכמה נקרא חכם וכו' והמכוין לאין סוף מחמת רוממותו שאין לו שם ולא נקודה שתגבילהו אין תפלתו תפלה, רק יכוין אליו בהיותו מלובש במדותיו וכו' ועיין לקמן סי' ק' בשם נחלת יוסף.

| id | word | punctuation after |
|---|---|---|
| 4:18.1 | גם |  |
| 4:18.2 | בספר |  |
| 4:18.3 | עץ |  |
| 4:18.4 | חיים |  |
| 4:18.5 | להרח"ו |  |
| 4:18.6 | שער |  |
| 4:18.7 | הכללים |  |
| 4:18.8 | פ | ' |
| 4:18.9 | י"א |  |
| 4:18.10 | כתב |  |
| 4:18.11 | שמשה |  |
| 4:18.12 | רבינו |  |
| 4:18.13 | אמר |  |
| 4:18.14 | לישראל |  |
| 4:18.15 | בכניסתן |  |
| 4:18.16 | לארץ |  |
| 4:18.17 | ואתם |  |
| 4:18.18 | הדבקים |  |
| 4:18.19 | בה | ' |
| 4:18.20 | אלהיכם |  |
| 4:18.21 | חיים |  |
| 4:18.22 | כולכם |  |
| 4:18.23 | היום | , |
| 4:18.24 | כי |  |
| 4:18.25 | ה | ' |
| 4:18.26 | אלהיכם |  |
| 4:18.27 | הוא |  |
| 4:18.28 | זעיר |  |
| 4:18.29 | אנפין |  |
| 4:18.30 | ונוקביה | , |
| 4:18.31 | וישראל |  |
| 4:18.32 | הנכנסים |  |
| 4:18.33 | לארץ |  |
| 4:18.34 | היו |  |
| 4:18.35 | וכו | ' |
| 4:18.36 | עי"ש | . |
| 4:18.37 | גם |  |
| 4:18.38 | בספר |  |
| 4:18.39 | הכוונות |  |
| 4:18.40 | במנהגי |  |
| 4:18.41 | האר"י |  |
| 4:18.42 | כתב |  |
| 4:18.43 | כי |  |
| 4:18.44 | כל |  |
| 4:18.45 | פעם |  |
| 4:18.46 | שתוציא |  |
| 4:18.47 | מפיך |  |
| 4:18.48 | שם |  |
| 4:18.49 | הוי'ה |  |
| 4:18.50 | תכוין |  |
| 4:18.51 | לזעיר |  |
| 4:18.52 | אנפין | , |
| 4:18.53 | ותכוין |  |
| 4:18.54 | בו |  |
| 4:18.55 | כך |  |
| 4:18.56 | יה"ו |  |
| 4:18.57 | זעיר |  |
| 4:18.58 | אנפין |  |
| 4:18.59 | וה"א |  |
| 4:18.60 | אחרונה |  |
| 4:18.61 | שבו |  |
| 4:18.62 | היא |  |
| 4:18.63 | מלכות | . |
| 4:18.64 | תכוין |  |
| 4:18.65 | שנעשה |  |
| 4:18.66 | שם |  |
| 4:18.67 | שלם |  |
| 4:18.68 | על |  |
| 4:18.69 | אדני |  |
| 4:18.70 | באופן |  |
| 4:18.71 | וכו | ' |
| 4:18.72 | ע"ש | . |
| 4:18.73 | וכן |  |
| 4:18.74 | בספר |  |
| 4:18.75 | מחברת |  |
| 4:18.76 | הקודש |  |
| 4:18.77 | בסדר |  |
| 4:18.78 | מוסף |  |
| 4:18.79 | שבת |  |
| 4:18.80 | כ | ' |
| 4:18.81 | כי |  |
| 4:18.82 | מלאכים |  |
| 4:18.83 | המוני |  |
| 4:18.84 | מעלה |  |
| 4:18.85 | יתנו |  |
| 4:18.86 | כתר |  |
| 4:18.87 | לזעיר |  |
| 4:18.88 | אנפין |  |
| 4:18.89 | שהוא |  |
| 4:18.90 | ה | ' |
| 4:18.91 | אלהינו |  |
| 4:18.92 | עיין |  |
| 4:18.93 | שם |  |
| 4:18.94 | דף |  |
| 4:18.95 | מ | ' |
| 4:18.96 | ע"א | . |
| 4:18.97 | גם |  |
| 4:18.98 | המקדש |  |
| 4:18.99 | מלך |  |
| 4:18.100 | ד | ' |
| 4:18.101 | י"ב |  |
| 4:18.102 | ד"ה |  |
| 4:18.103 | אלהא |  |
| 4:18.104 | רברבא |  |
| 4:18.105 | מבואר |  |
| 4:18.106 | דלא |  |
| 4:18.107 | נקרא |  |
| 4:18.108 | אלהא |  |
| 4:18.109 | רברבא |  |
| 4:18.110 | אלא |  |
| 4:18.111 | בזעיר |  |
| 4:18.112 | אנפין | , |
| 4:18.113 | ובחכמה |  |
| 4:18.114 | נקרא |  |
| 4:18.115 | חכם |  |
| 4:18.116 | וכו | ' |
| 4:18.117 | והמכוין |  |
| 4:18.118 | לאין |  |
| 4:18.119 | סוף |  |
| 4:18.120 | מחמת |  |
| 4:18.121 | רוממותו |  |
| 4:18.122 | שאין |  |
| 4:18.123 | לו |  |
| 4:18.124 | שם |  |
| 4:18.125 | ולא |  |
| 4:18.126 | נקודה |  |
| 4:18.127 | שתגבילהו |  |
| 4:18.128 | אין |  |
| 4:18.129 | תפלתו |  |
| 4:18.130 | תפלה | , |
| 4:18.131 | רק |  |
| 4:18.132 | יכוין |  |
| 4:18.133 | אליו |  |
| 4:18.134 | בהיותו |  |
| 4:18.135 | מלובש |  |
| 4:18.136 | במדותיו |  |
| 4:18.137 | וכו | ' |
| 4:18.138 | ועיין |  |
| 4:18.139 | לקמן |  |
| 4:18.140 | סי | ' |
| 4:18.141 | ק | ' |
| 4:18.142 | בשם |  |
| 4:18.143 | נחלת |  |
| 4:18.144 | יוסף | . |


### Translation under review (4:18)

English: Also, in the book Etz Chayyim by Rabbi Chayyim Vital, Gate of the Generalities, chapter 11, he wrote that our teacher Moses said to Israel at their entering the land: "But you who cleave to the LORD your God are alive, every one of you, today," for the LORD your God is Ze'ir Anpin and his Nukva, and Israel who entered the land were, etc. See there. Also in the book Sefer ha-Khavanot, in the customs of the Ari, he wrote that every time you utter the name YHVH, intend toward Ze'ir Anpin, and intend in it thus: Yod-Heh-Vav Ze'ir Anpin, and the last Heh of it is Malkhut (Kingship). intend that the Name be made complete on the bases, in such a way, etc. See there. Likewise, in the book Machberet ha-Kodesh, in the order of the Musaf of Shabbat, 20, that the hosts of angels above will give a crown to Ze'ir Anpin, who is the LORD our God. See there, folio 40a. Also in Mikdash Melekh, 4, 12, beginning with the words it is explained that 'the Great God' is called only in Ze'ir Anpin. and in Chokhmah (Wisdom) he is called wise, etc. And whoever intends toward Ein Sof, because of His exaltedness that He has no name, and no point that would limit Him, his prayer is not a prayer, rather he should intend toward Him while He is clothed in His attributes, etc. And see below, §100, in the name of Nachalat Yosef.

Literal: also in the book Etz Chayyim by Rabbi Chayyim Vital Gate of the Generalities chapter 11 wrote that Moses our teacher said to Israel at their entering the land but you who cleave to the LORD your God are alive every one of you today for the LORD your God is Ze'ir Anpin and his Nukva (the female) and Israel who entered the land were etc. see there also in the book Sefer ha-Khavanot in the customs of the Ari wrote that every time you utter from your mouth the name YHVH intend toward Ze'ir Anpin and intend in it thus Yod-Heh-Vav Ze'ir Anpin and the Heh last of it is Malkhut (Kingship) intend that the Name complete on the bases in such a way etc. see there likewise in the book Machberet ha-Kodesh in the order of the Musaf of Shabbat 20 that angels the hosts of above will give a crown to Ze'ir Anpin who is the LORD our God see there folio 40 a also Mikdash Melekh 4 12 beginning with the words the Great God (great) is explained that is called the Great God (great) only in Ze'ir Anpin and in Chokhmah (Wisdom) he is called wise etc. and whoever intends toward Ein Sof because of His exaltedness that there is not to Him a name and not a point that would limit Him is not his prayer a prayer rather should intend toward Him while He is clothed in His attributes etc. and see below section 100 in the name of Nachalat Yosef

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 4:18.1 | also |  |  |  |
| 4:18.2 | in the book |  |  |  |
| 4:18.3 | Etz |  |  | Etz Chayyim: 'Tree of Life'. |
| 4:18.4 | Chayyim |  |  |  |
| 4:18.5 | by Rabbi Chayyim Vital |  |  |  |
| 4:18.6 | Gate |  |  |  |
| 4:18.7 | of the Generalities |  |  |  |
| 4:18.8 | chapter |  |  |  |
| 4:18.9 | 11 |  |  |  |
| 4:18.10 | wrote |  |  |  |
| 4:18.11 | that Moses |  |  |  |
| 4:18.12 | our teacher |  |  |  |
| 4:18.13 | said |  |  |  |
| 4:18.14 | to Israel |  |  |  |
| 4:18.15 | at their entering |  |  |  |
| 4:18.16 | the land |  |  |  |
| 4:18.17 | but you |  |  |  |
| 4:18.18 | who cleave |  |  |  |
| 4:18.19 | to the LORD |  |  |  |
| 4:18.20 | your God |  |  |  |
| 4:18.21 | are alive |  |  |  |
| 4:18.22 | every one of you |  |  |  |
| 4:18.23 | today |  |  |  |
| 4:18.24 | for |  |  |  |
| 4:18.25 | the LORD |  |  |  |
| 4:18.26 | your God |  |  |  |
| 4:18.27 | is |  |  |  |
| 4:18.28 | Ze'ir |  |  |  |
| 4:18.29 | Anpin |  |  |  |
| 4:18.30 | and his Nukva (the female) |  |  | Nukva: 'female', the feminine partner of Ze'ir Anpin (Aramaic ונוקביה). |
| 4:18.31 | and Israel |  |  |  |
| 4:18.32 | who entered |  |  |  |
| 4:18.33 | the land |  |  |  |
| 4:18.34 | were |  |  |  |
| 4:18.35 | etc. |  |  |  |
| 4:18.36 | see there |  |  |  |
| 4:18.37 | also |  |  |  |
| 4:18.38 | in the book |  |  |  |
| 4:18.39 | Sefer ha-Khavanot |  |  | Sefer ha-Khavanot: 'the book of intentions'. Kavanot are the intentions of the mind in prayer and in the Name. |
| 4:18.40 | in the customs of |  |  |  |
| 4:18.41 | the Ari |  |  |  |
| 4:18.42 | wrote |  |  |  |
| 4:18.43 | that |  |  |  |
| 4:18.44 | every |  |  |  |
| 4:18.45 | time |  |  |  |
| 4:18.46 | you utter |  |  |  |
| 4:18.47 | from your mouth |  |  |  |
| 4:18.48 | the name |  |  |  |
| 4:18.49 | YHVH |  |  |  |
| 4:18.50 | intend |  |  |  |
| 4:18.51 | toward Ze'ir |  |  |  |
| 4:18.52 | Anpin |  |  |  |
| 4:18.53 | and intend |  |  |  |
| 4:18.54 | in it |  |  |  |
| 4:18.55 | thus |  |  |  |
| 4:18.56 | Yod-Heh-Vav |  |  | Yod, Heh, Vav: the letters of the Name spelled out. |
| 4:18.57 | Ze'ir |  |  |  |
| 4:18.58 | Anpin |  |  |  |
| 4:18.59 | and the Heh |  |  | The last Heh: the letter ה of the Name, as the author takes it. |
| 4:18.60 | last |  |  |  |
| 4:18.61 | of it |  |  |  |
| 4:18.62 | is |  |  |  |
| 4:18.63 | Malkhut (Kingship) |  |  | Malkhut (Kingship): the tenth sefirah. |
| 4:18.64 | intend |  |  |  |
| 4:18.65 | that |  |  |  |
| 4:18.66 | the Name |  |  |  |
| 4:18.67 | complete |  |  |  |
| 4:18.68 | on |  |  |  |
| 4:18.69 | the bases |  |  | אדנים (adanim): 'bases' or sockets. |
| 4:18.70 | in such a way |  |  |  |
| 4:18.71 | etc. |  |  |  |
| 4:18.72 | see there |  |  |  |
| 4:18.73 | likewise |  |  |  |
| 4:18.74 | in the book |  |  |  |
| 4:18.75 | Machberet |  |  | Machberet: 'composition'; the title means 'the holy composition'. |
| 4:18.76 | ha-Kodesh |  |  |  |
| 4:18.77 | in the order of |  |  |  |
| 4:18.78 | the Musaf |  |  |  |
| 4:18.79 | of Shabbat |  |  |  |
| 4:18.80 | 20 |  |  |  |
| 4:18.81 | that |  |  |  |
| 4:18.82 | angels |  |  |  |
| 4:18.83 | the hosts of |  |  |  |
| 4:18.84 | above |  |  |  |
| 4:18.85 | will give |  |  |  |
| 4:18.86 | a crown |  |  |  |
| 4:18.87 | to Ze'ir |  |  |  |
| 4:18.88 | Anpin |  |  |  |
| 4:18.89 | who is |  |  |  |
| 4:18.90 | the LORD |  |  |  |
| 4:18.91 | our God |  |  |  |
| 4:18.92 | see |  |  |  |
| 4:18.93 | there |  |  |  |
| 4:18.94 | folio |  |  |  |
| 4:18.95 | 40 |  |  |  |
| 4:18.96 | a |  |  |  |
| 4:18.97 | also |  |  |  |
| 4:18.98 | Mikdash |  |  |  |
| 4:18.99 | Melekh |  |  |  |
| 4:18.100 | 4 |  |  |  |
| 4:18.101 | 12 |  |  |  |
| 4:18.102 | beginning with the words |  |  |  |
| 4:18.103 | the Great God |  |  | Elaha Rabba: Aramaic for 'the Great God'. The next word, rabba, is 'great'. |
| 4:18.104 | (great) |  |  |  |
| 4:18.105 | is explained |  |  |  |
| 4:18.106 | that |  |  |  |
| 4:18.107 | is called |  |  |  |
| 4:18.108 | the Great God |  |  |  |
| 4:18.109 | (great) |  |  |  |
| 4:18.110 | only |  |  |  |
| 4:18.111 | in Ze'ir |  |  |  |
| 4:18.112 | Anpin |  |  |  |
| 4:18.113 | and in Chokhmah (Wisdom) |  |  | Chokhmah: the second sefirah, Wisdom. |
| 4:18.114 | he is called |  |  |  |
| 4:18.115 | wise |  |  |  |
| 4:18.116 | etc. |  |  |  |
| 4:18.117 | and whoever intends |  |  |  |
| 4:18.118 | toward Ein |  |  |  |
| 4:18.119 | Sof |  |  |  |
| 4:18.120 | because of |  |  |  |
| 4:18.121 | His exaltedness |  |  |  |
| 4:18.122 | that there is not |  |  |  |
| 4:18.123 | to Him |  |  |  |
| 4:18.124 | a name |  |  |  |
| 4:18.125 | and not |  |  |  |
| 4:18.126 | a point |  |  |  |
| 4:18.127 | that would limit Him |  |  |  |
| 4:18.128 | is not |  |  |  |
| 4:18.129 | his prayer |  |  |  |
| 4:18.130 | a prayer |  |  |  |
| 4:18.131 | rather |  |  |  |
| 4:18.132 | should intend |  |  |  |
| 4:18.133 | toward Him |  |  |  |
| 4:18.134 | while He is |  |  |  |
| 4:18.135 | clothed |  |  |  |
| 4:18.136 | in His attributes |  |  |  |
| 4:18.137 | etc. |  |  |  |
| 4:18.138 | and see |  |  |  |
| 4:18.139 | below |  |  |  |
| 4:18.140 | section |  |  |  |
| 4:18.141 | 100 |  |  |  |
| 4:18.142 | in the name of |  |  |  |
| 4:18.143 | Nachalat |  |  | Nachalat Yosef: 'the inheritance of Yosef'. |
| 4:18.144 | Yosef |  |  |  |

Phrases:
- 4:18.1 4:18.2 4:18.3 4:18.4 4:18.5 4:18.6 4:18.7 4:18.8 4:18.9 | Also, in the book Etz Chayyim by Rabbi Chayyim Vital, Gate of the Generalities, chapter 11, | also in the book Etz Chayyim by Rabbi Chayyim Vital Gate of the Generalities chapter 11
- 4:18.10 4:18.11 4:18.12 4:18.13 4:18.14 4:18.15 4:18.16 | he wrote that our teacher Moses said to Israel at their entering the land: | wrote that Moses our teacher said to Israel at their entering the land
- 4:18.17 4:18.18 4:18.19 4:18.20 4:18.21 4:18.22 4:18.23 | "But you who cleave to the LORD your God are alive, every one of you, today," | but you who cleave to the LORD your God are alive every one of you today
- 4:18.24 4:18.25 4:18.26 4:18.27 4:18.28 4:18.29 4:18.30 | for the LORD your God is Ze'ir Anpin and his Nukva, | for the LORD your God is Ze'ir Anpin and his Nukva (the female)
- 4:18.31 4:18.32 4:18.33 4:18.34 4:18.35 4:18.36 | and Israel who entered the land were, etc. See there. | and Israel who entered the land were etc. see there
- 4:18.37 4:18.38 4:18.39 4:18.40 4:18.41 4:18.42 4:18.43 | Also in the book Sefer ha-Khavanot, in the customs of the Ari, he wrote that | also in the book Sefer ha-Khavanot in the customs of the Ari wrote that
- 4:18.44 4:18.45 4:18.46 4:18.47 4:18.48 4:18.49 4:18.50 4:18.51 4:18.52 | every time you utter the name YHVH, intend toward Ze'ir Anpin, | every time you utter from your mouth the name YHVH intend toward Ze'ir Anpin
- 4:18.53 4:18.54 4:18.55 4:18.56 4:18.57 4:18.58 4:18.59 4:18.60 4:18.61 4:18.62 4:18.63 | and intend in it thus: Yod-Heh-Vav Ze'ir Anpin, and the last Heh of it is Malkhut (Kingship). | and intend in it thus Yod-Heh-Vav Ze'ir Anpin and the Heh last of it is Malkhut (Kingship)
- 4:18.64 4:18.65 4:18.66 4:18.67 4:18.68 4:18.69 4:18.70 4:18.71 4:18.72 | intend that the Name be made complete on the bases, in such a way, etc. See there. | intend that the Name complete on the bases in such a way etc. see there
- 4:18.73 4:18.74 4:18.75 4:18.76 4:18.77 4:18.78 4:18.79 4:18.80 | Likewise, in the book Machberet ha-Kodesh, in the order of the Musaf of Shabbat, 20, | likewise in the book Machberet ha-Kodesh in the order of the Musaf of Shabbat 20
- 4:18.81 4:18.82 4:18.83 4:18.84 4:18.85 4:18.86 4:18.87 4:18.88 | that the hosts of angels above will give a crown to Ze'ir Anpin, | that angels the hosts of above will give a crown to Ze'ir Anpin
- 4:18.89 4:18.90 4:18.91 4:18.92 4:18.93 4:18.94 4:18.95 4:18.96 | who is the LORD our God. See there, folio 40a. | who is the LORD our God see there folio 40 a
- 4:18.97 4:18.98 4:18.99 4:18.100 4:18.101 4:18.102 | Also in Mikdash Melekh, 4, 12, beginning with the words | also Mikdash Melekh 4 12 beginning with the words
- 4:18.103 4:18.104 4:18.105 4:18.106 4:18.107 4:18.108 4:18.109 4:18.110 4:18.111 4:18.112 | it is explained that 'the Great God' is called only in Ze'ir Anpin. | the Great God (great) is explained that is called the Great God (great) only in Ze'ir Anpin
- 4:18.113 4:18.114 4:18.115 4:18.116 | and in Chokhmah (Wisdom) he is called wise, etc. | and in Chokhmah (Wisdom) he is called wise etc.
- 4:18.117 4:18.118 4:18.119 4:18.120 4:18.121 4:18.122 4:18.123 4:18.124 | And whoever intends toward Ein Sof, because of His exaltedness that He has no name, | and whoever intends toward Ein Sof because of His exaltedness that there is not to Him a name
- 4:18.125 4:18.126 4:18.127 | and no point that would limit Him, | and not a point that would limit Him
- 4:18.128 4:18.129 4:18.130 | his prayer is not a prayer, | is not his prayer a prayer
- 4:18.131 4:18.132 4:18.133 4:18.134 4:18.135 4:18.136 4:18.137 | rather he should intend toward Him while He is clothed in His attributes, etc. | rather should intend toward Him while He is clothed in His attributes etc.
- 4:18.138 4:18.139 4:18.140 4:18.141 | And see below, §100, | and see below section 100
- 4:18.142 4:18.143 4:18.144 | in the name of Nachalat Yosef. | in the name of Nachalat Yosef

Grammar notes:
- The quotation of Deut 4:4 (4:18.17-23) has no formula; it is set in quotation marks as the author's direct words. The author's own gloss begins with כי (4:18.24) and is not part of the verse.
- וכו' (4:18.35, 4:18.71, 4:18.116, 4:18.137) marks an abbreviated passage. It stays as 'etc.' because it marks the truncation.
- יה"ו (4:18.56) spells the letters yod, heh, vav of the Name; וה"א אחרונה (4:18.59) is the last heh.
- על אדני (4:18.68-69): 'bases', with the verb שנעשה ('that it be made') in a passive-like construction, 'that the Name be made complete'.
- אלהא רברבא (4:18.103-104) is Aramaic. מבואר ('it is explained') introduces indirect speech, so no quotation marks are used for the sentence.

Translator's note: Ein Sof (4:18.118-119) is the same name as in 4:17; see the note there.

Commentary rulings:
- apparatus: Deut 4:4 (4:18.17-23): 'But you who cleave to the LORD your God are alive, every one of you, today.'

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

