# Translation pack: milhamot-hashem perek-7-17 (Milchamot ha-Shem 7:22–7:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-17`.

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
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
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


## Milchamot ha-Shem 7:22

Hebrew (Wikisource): צד) ואם בספור עסקי ב"ו הזהירו חכמים ז"ל שלא יוציא אדם דבר מגונה מפיו! למלך מלכי המלכים הקב"ה עאכ"ו. כמאמר שלמה המלך ע"ה אל תתן את פיך לחטיא את בשרך ואל תאמר לפני המלאך כי שגגה היא למה יקצוף וכו'! ובירושלמי פרק אין דורשין גרסינן רב אמר תאלמנה שפתי שקר, יתחרשן, יתפרכן, ישתתקן. יתחרשן כמה דאתאמר ויאמר ה' אליו מי שם פה לאדם אי מי ישום אלם. יתפרכן כמה דתימר והנה אנחנו מאלמים אלומים (פי' הערוך בערך פרך ענין קשירה) ישתתקן כמשמעו. הדוברות על צדיק עתק. בדוברות על צדיקו של עולם דברים שהעתיק מבריותיו. בגאוה ובוז. זה שהוא מתגאה לומר אני דורש במעשה בראשית. סבור שהוא מתגאה ואינו אלא כמבזה. אמר ר' יוסי בר' חנינה המתכבד בקלון חבריו אין לו חלק לעה"ב. המתכבד בכבוד חי העולמים לא כל שכן. מה כתיב בתריה מה רב טובך אשר צפנת ליראיך. אל יהי לו במה רב טובך ע"כ. וכן אמרו בבראשית רבה פ' א' רב הונא ב' בר קפרא פתח תאלמנה שפתי שקר הדוברות על צדיק עתק בגאוה ובוז. יתחרשן. יתפרכן. ישתתקון. יתחרשון היך מה דאתאמר מי שם פה לאדם או מי ישום אלם. או חרש. יתפרכון. היך מה דאתאמר והנה אנחנו מלאמים אלומים והנה קמה אלומתי. ישתתקון כמשמעו. הדוברות על צדיק חי העולמים. עתק דברים שהעתיק מבריותיו. בגאוה אתמהא? בשביל להתגאות ולומר אני דורש במעשה בראשית. ובוז אתמהא? מבזה על כבודו דאמר ר' חנינה כל המתכבד בקלון חבירו אין לו חלק לעה"ב. בכבודו של מקום עאכ"ו. ומה כתיב אחריו. מה רב טובך אשר צפנת ליראיך. ליראיך ולא לבוזים את מוראך הרב אל יהיו במה רב טובך ע"כ. ובכתובות פ' המדיר אהא דתנן על מנת שתאמרי לפלוני מה שאמרתי לך או מה שאמרת לי, פריך בגמרא ותאמר? ופריק דברים של קלון. ופי' הרמב״ם שם שמכריחה לומר לאחר מענייני תשמיש וזהו אמרם דברים של קלון. הרי דרבותי' ז"ל בעלי הגמרא קראו לעסקי תשמיש דברים של קלון. ודברים כאלה הם מהמפורסמות שאין צריך ראיה רק למתעקש! ואיה בוז וחרפה וקלון יותר מזה ליחס לאלהים אשר נעבוד ונשתחוה לו אברי תשמיש. גיד וביצים, ואשה, ורחם (עיין זהר פרשת בהר דף ק"ט וסוף אדרא דף רצ"ו ע"א) והשתוקקות לחבקה ולנשקה בהיותה מתקשטת לפניו. ובהזדווגו עמה עושה לה נחת רוח ברחמה עד שהפיטן מהם בעזות מצחו אמר יחבק לה בעלה וביסודא דילה דעביד ניחא לה יהא כתיש כתישין. הלה' אלהיכם שאתם אומרים שהוא קצר אפים (זעיר אנפין) תגמלו זאת? והלא כבר נאמר בן יכבד אב ועבד אדוניו. ואם אב אני איה כבודי ואם אדונים אני איה מוראי? האם כבודו של אב שיספר בנו עליו שהוא מחבק את אשתו. ועושה לה נחת רוח ברחמה בדישו שם דיש? ואם זהו מוראו של אדון שיקראהו עבדו קצר אפים (זעיר נאפין) ויספר עליו דברי קלון כאלה? למה יגרע כבוד אבא ואימא וזעיר אנפין ונוקביה שאתם אומרים שהוא ה' אלהיכם והוא שנגלה בסיני ונתן לנו את התורה. מכבוד האב והאדון? לא על הדברים האלה אמר הנביא ויחפאו בני ישראל דברים אשר לא כן על ה' אלהיהם? ואיך נאמר קדוש קדוש קדוש ה' צבאות. וגם נאמר שהוא קדוש ומשרתיו קדושים! ונספר עליו דברי בוז וקלון וחרפה כאלה? דברים שאין להם שום שורש ולא שמץ דבר, לא בתורה ולא בנביאים ולא בכתובים לא בדברי רז"ל שקבלו איש מפי איש (כמבואר במסכת אבות) לא במשנה ולא בתלמוד בבלי וירושלמי ולא במדרשים האמתיים. שהרי תוה"ק לא יחסה לשי"ת חוש המשוש גם בחלק המעולה שבו כמו שכתב הרב המורה ז"ל. כ"ש וק"ו החלק הפוחת והמגונה שבו דהיינו המשגל. וחלילה להאמין כי רבותינו ז"ל הקדושים יוציאו מפיהם דברים כאלה על ה' אלהינו ועל כל משרתיו. ולקראו בשם קצר אפים ולספר עליו או על משרתיו ענייני משגל ופריה ורביה להוליד נשמות למלאכים ולבני תמותה, והנביא אמר נותן נשמה לעם עליה ורוח להולכים בה, ולא מוליד. האם על הקבלה החדשה הנכריה יאמרו הגוים עלינו רק עם חכם ונבון הגוי הגדול הזה? האם יחוד ה' שנצטוינו בתורה הוא לחשוב את כל הצורות והפרצופים הנזכרים באלו הם אחד? האם עליהם אמר דוד המלך ע"ה ואדברה בעדותיך נגד מלכים ולא אבוש? הלא נשכבה בבשתינו ותכסינו כלמתינו אם נספר באזני ערבי אחד כדברים האלה! ובילקוט פר' קרח רמז תשנ"ו וז"ל חביבה תורה. שכששאל מלך ישראל לא שאל אלא תורה שנ' טוב אתה ומטיב למדני חקיך. ואמר סעדני ואושעה ואשעה בחקיך וכו' שלא אהא למד תורה ואין יצר הרע מניח לי לשנות. או אהא למד ושוכח או אטמא את הטהור ומטהר את הטמא ונמצאתי אבד לעה"ב. או שמא ישאלוני מגויי הארץ וממשפחות האדמה ואיני יודע להשיבן, ונמצאתי בוש לעיניהם. וכן הוא אומר ואדברה בעדותיך נגד מלכים ולא אבוש וכו' עכ"ל. הרי בהדיא שחוקי תוה"ק ידיעותיה ומשפטיה לא נבוש לפסרם נגד מלכי אוה"ע! מה שאין כן יחודים ופרצופים וזווגים הנז' בזהר וספרי המקובלים החדשים. ולא לחנם הזהירו גדולי חכמי ישראל שלא ללמדה ולילך במנהגים הנז' בה היפך התלמוד והפוסקים, הלא הרב מהרש"ל ור"ת ן' יחיא המובאים בשו"ת רביד הזהב, גם ר' פרץ הובא בשו"ת הריב"ש סי' קנ"ז לא היה מדבר ולא מחשיב לאותן ספירות, (פי' לא היה מחשיב לאותן ספירות שום אלהות), גם ר' שמשון מקינון מאן לכוון לשום ספירה והיה אומר שהוא מתפלל לדעת זה התינוק (פי' כתינוק זה שמתפלל לשוכן שמים) להוציא מלב המקובלים וכו', והריב"ש ז"ל הרגיש בזרותה וכתב איני תוקע עצמי ללמדה וכו', וכתב עוד וכי הספירות יש להם אלהות? והמשיב רבי יוסף ן' שושן כיחש לו והביא לו משל שאינו לפי האמת! כי האמת הגמור שהמקובלים סוברים דספירות דיושר הם עיקר האלהות. והוא ז"ל הרגיש בסתירה, ולכן אמר איני תוקע עצמי ללמדה וכו' גם הרב בעל העיקרים פרק כ"ח משני הזהיר מאד שלא ללמוד הקבלה וז"ל וכלל אומר לך השמר לך ושמור נפשך מאד פן תנקש אחריהם ותלכד במצודתם, כי הם עוזבים ארחות יושר ללכת בדרכי חשך (לעבוד אלהים אחרים) כי לא ידעו ולא יבינו בחשכה יתהלכו העוסקים בקבלה מדעת עצמם מזולת קבלה מפי חכם מקובל עכ״ל. וכבר נתבאר לעיל סי' י"ח משם מהר"י צאהרי שאמונת הקבלה הנכריה החדשה הזו, חדשה היא בארץ התימן על ידי הספרים שבאו בה מאותו זמן כעדותו של מהר"י הנז'. ומעולם לא נמצא בתימן איש מקובל מפה אל פה, וספרי קדמונינו הנשארים בידינו יעידון על זה. וכפי דברי הרב שם הגדולים במערכת ספרים ערך זהר גם בכל ישראל חדשה היא וז"ל ראיתי כתוב מהר' מהר"א רויגו וז"ל מצאתי בזהר כ"י ישן נושן אצל מורי מהר"ם זכות וז"ל מצאתי כתוב באמת כי ראש המקובלים רבי נחוניא בן הקנה והוא חיבר ס' הבהיר, ואחריו רשב"י עשה ס' הזהר וחיבר בו כמה חיבורים כמו התקונים, וכשמת רשב"י ור"א וכל הדור ההוא אבדה חכמת הקבלה, עד שהקרה ה' לפני מלך ממלכי מזרח שצוה לחפור במקום אחד על עסקי ממון, ונמצא שם ארון אחד, ובו ספר הזהר ושלח לחכמי אדום ולא ידעו ולא יבינו, שלח אחר היהודים באו אצלו, וראו הספר ואמרו לו אדונינו המלך, זה הס' עשאו חכם אחד והוא עמוק מאד ואין אנו מבינים אותו, אמר להם וכי אין יהודי בעולם שמבין אותו? אמרו לו יש במדינת טוליטולא, והמלך שלח הספרים עם גבוריו לטוליטולה, וכשראוהו חכמי טוליטולה שמחו בו שמחה גדולה, ושלחו למלך מתנות רבות, ומשם נתפרסמה הקבלה בישראל ע"כ מצאתי מהר' הנז' עכ"ל. וכדברים האלה כתבו חכמי ירושלם אלינו, והרואה יבין דבר לאשורו ממה שכתב שאחר מות רשב"י אבדה חכמת הקבלה ונשכלה מישראל, ולא למדו דורות האחרונים רק מתוך הספר המזויף שקבלו אותו חכמי טוליטולה מיד המלך הנכרי ושמחו בו שמחה גדולה, ועל ידו המירו את כבודם ולא קיבלו מפי חכם מקובל, גם כל האנשים שראינו אותם הולכים בדרכי הקבלה הלזו כל ידיעתם ואמונתם בה מפי הספרים ולא מפי סופרים, וכפי דברי הרמב"ן ודאי שהלומד מתוך הספרים ישגה בה ויבא לידי טעות להאמין ברשויות הרבה ולעבוד קצר אפים נברא, אשר לפי דבריהם נשמתו חלק אלוה ממעל והמה עובדים לחלק ההוא שבזעיר אנפין בחשבם כי האלהים האמצעים אבא ואימא נתנו הכח והממשלה לשלוט בתחתונים. ואמר לו בני אתה (זהר בלק דף קצ"א) וצוו לעבדו ולברך בשמו, וקראו לו בשם ידוד כשם ה' אלהינו כמו שעשו דורו של אנוש כדלעיל סי' נ"ו ונ"ז.

| id | word | punctuation after |
|---|---|---|
| 7:22.1 | צד | ) |
| 7:22.2 | ואם |  |
| 7:22.3 | בספור |  |
| 7:22.4 | עסקי |  |
| 7:22.5 | ב"ו |  |
| 7:22.6 | הזהירו |  |
| 7:22.7 | חכמים |  |
| 7:22.8 | ז"ל |  |
| 7:22.9 | שלא |  |
| 7:22.10 | יוציא |  |
| 7:22.11 | אדם |  |
| 7:22.12 | דבר |  |
| 7:22.13 | מגונה |  |
| 7:22.14 | מפיו | ! |
| 7:22.15 | למלך |  |
| 7:22.16 | מלכי |  |
| 7:22.17 | המלכים |  |
| 7:22.18 | הקב"ה |  |
| 7:22.19 | עאכ"ו | . |
| 7:22.20 | כמאמר |  |
| 7:22.21 | שלמה |  |
| 7:22.22 | המלך |  |
| 7:22.23 | ע"ה |  |
| 7:22.24 | אל |  |
| 7:22.25 | תתן |  |
| 7:22.26 | את |  |
| 7:22.27 | פיך |  |
| 7:22.28 | לחטיא |  |
| 7:22.29 | את |  |
| 7:22.30 | בשרך |  |
| 7:22.31 | ואל |  |
| 7:22.32 | תאמר |  |
| 7:22.33 | לפני |  |
| 7:22.34 | המלאך |  |
| 7:22.35 | כי |  |
| 7:22.36 | שגגה |  |
| 7:22.37 | היא |  |
| 7:22.38 | למה |  |
| 7:22.39 | יקצוף |  |
| 7:22.40 | וכו | '! |
| 7:22.41 | ובירושלמי |  |
| 7:22.42 | פרק |  |
| 7:22.43 | אין |  |
| 7:22.44 | דורשין |  |
| 7:22.45 | גרסינן |  |
| 7:22.46 | רב |  |
| 7:22.47 | אמר |  |
| 7:22.48 | תאלמנה |  |
| 7:22.49 | שפתי |  |
| 7:22.50 | שקר | , |
| 7:22.51 | יתחרשן | , |
| 7:22.52 | יתפרכן | , |
| 7:22.53 | ישתתקן | . |
| 7:22.54 | יתחרשן |  |
| 7:22.55 | כמה |  |
| 7:22.56 | דאתאמר |  |
| 7:22.57 | ויאמר |  |
| 7:22.58 | ה | ' |
| 7:22.59 | אליו |  |
| 7:22.60 | מי |  |
| 7:22.61 | שם |  |
| 7:22.62 | פה |  |
| 7:22.63 | לאדם |  |
| 7:22.64 | אי |  |
| 7:22.65 | מי |  |
| 7:22.66 | ישום |  |
| 7:22.67 | אלם | . |
| 7:22.68 | יתפרכן |  |
| 7:22.69 | כמה |  |
| 7:22.70 | דתימר |  |
| 7:22.71 | והנה |  |
| 7:22.72 | אנחנו |  |
| 7:22.73 | מאלמים |  |
| 7:22.74 | אלומים |  |
| 7:22.75 | פי | ' |
| 7:22.76 | הערוך |  |
| 7:22.77 | בערך |  |
| 7:22.78 | פרך |  |
| 7:22.79 | ענין |  |
| 7:22.80 | קשירה | ) |
| 7:22.81 | ישתתקן |  |
| 7:22.82 | כמשמעו | . |
| 7:22.83 | הדוברות |  |
| 7:22.84 | על |  |
| 7:22.85 | צדיק |  |
| 7:22.86 | עתק | . |
| 7:22.87 | בדוברות |  |
| 7:22.88 | על |  |
| 7:22.89 | צדיקו |  |
| 7:22.90 | של |  |
| 7:22.91 | עולם |  |
| 7:22.92 | דברים |  |
| 7:22.93 | שהעתיק |  |
| 7:22.94 | מבריותיו | . |
| 7:22.95 | בגאוה |  |
| 7:22.96 | ובוז | . |
| 7:22.97 | זה |  |
| 7:22.98 | שהוא |  |
| 7:22.99 | מתגאה |  |
| 7:22.100 | לומר |  |
| 7:22.101 | אני |  |
| 7:22.102 | דורש |  |
| 7:22.103 | במעשה |  |
| 7:22.104 | בראשית | . |
| 7:22.105 | סבור |  |
| 7:22.106 | שהוא |  |
| 7:22.107 | מתגאה |  |
| 7:22.108 | ואינו |  |
| 7:22.109 | אלא |  |
| 7:22.110 | כמבזה | . |
| 7:22.111 | אמר |  |
| 7:22.112 | ר | ' |
| 7:22.113 | יוסי |  |
| 7:22.114 | בר | ' |
| 7:22.115 | חנינה |  |
| 7:22.116 | המתכבד |  |
| 7:22.117 | בקלון |  |
| 7:22.118 | חבריו |  |
| 7:22.119 | אין |  |
| 7:22.120 | לו |  |
| 7:22.121 | חלק |  |
| 7:22.122 | לעה"ב | . |
| 7:22.123 | המתכבד |  |
| 7:22.124 | בכבוד |  |
| 7:22.125 | חי |  |
| 7:22.126 | העולמים |  |
| 7:22.127 | לא |  |
| 7:22.128 | כל |  |
| 7:22.129 | שכן | . |
| 7:22.130 | מה |  |
| 7:22.131 | כתיב |  |
| 7:22.132 | בתריה |  |
| 7:22.133 | מה |  |
| 7:22.134 | רב |  |
| 7:22.135 | טובך |  |
| 7:22.136 | אשר |  |
| 7:22.137 | צפנת |  |
| 7:22.138 | ליראיך | . |
| 7:22.139 | אל |  |
| 7:22.140 | יהי |  |
| 7:22.141 | לו |  |
| 7:22.142 | במה |  |
| 7:22.143 | רב |  |
| 7:22.144 | טובך |  |
| 7:22.145 | ע"כ | . |
| 7:22.146 | וכן |  |
| 7:22.147 | אמרו |  |
| 7:22.148 | בבראשית |  |
| 7:22.149 | רבה |  |
| 7:22.150 | פ | ' |
| 7:22.151 | א | ' |
| 7:22.152 | רב |  |
| 7:22.153 | הונא |  |
| 7:22.154 | ב | ' |
| 7:22.155 | בר |  |
| 7:22.156 | קפרא |  |
| 7:22.157 | פתח |  |
| 7:22.158 | תאלמנה |  |
| 7:22.159 | שפתי |  |
| 7:22.160 | שקר |  |
| 7:22.161 | הדוברות |  |
| 7:22.162 | על |  |
| 7:22.163 | צדיק |  |
| 7:22.164 | עתק |  |
| 7:22.165 | בגאוה |  |
| 7:22.166 | ובוז | . |
| 7:22.167 | יתחרשן | . |
| 7:22.168 | יתפרכן | . |
| 7:22.169 | ישתתקון | . |
| 7:22.170 | יתחרשון |  |
| 7:22.171 | היך |  |
| 7:22.172 | מה |  |
| 7:22.173 | דאתאמר |  |
| 7:22.174 | מי |  |
| 7:22.175 | שם |  |
| 7:22.176 | פה |  |
| 7:22.177 | לאדם |  |
| 7:22.178 | או |  |
| 7:22.179 | מי |  |
| 7:22.180 | ישום |  |
| 7:22.181 | אלם | . |
| 7:22.182 | או |  |
| 7:22.183 | חרש | . |
| 7:22.184 | יתפרכון | . |
| 7:22.185 | היך |  |
| 7:22.186 | מה |  |
| 7:22.187 | דאתאמר |  |
| 7:22.188 | והנה |  |
| 7:22.189 | אנחנו |  |
| 7:22.190 | מלאמים |  |
| 7:22.191 | אלומים |  |
| 7:22.192 | והנה |  |
| 7:22.193 | קמה |  |
| 7:22.194 | אלומתי | . |
| 7:22.195 | ישתתקון |  |
| 7:22.196 | כמשמעו | . |
| 7:22.197 | הדוברות |  |
| 7:22.198 | על |  |
| 7:22.199 | צדיק |  |
| 7:22.200 | חי |  |
| 7:22.201 | העולמים | . |
| 7:22.202 | עתק |  |
| 7:22.203 | דברים |  |
| 7:22.204 | שהעתיק |  |
| 7:22.205 | מבריותיו | . |
| 7:22.206 | בגאוה |  |
| 7:22.207 | אתמהא | ? |
| 7:22.208 | בשביל |  |
| 7:22.209 | להתגאות |  |
| 7:22.210 | ולומר |  |
| 7:22.211 | אני |  |
| 7:22.212 | דורש |  |
| 7:22.213 | במעשה |  |
| 7:22.214 | בראשית | . |
| 7:22.215 | ובוז |  |
| 7:22.216 | אתמהא | ? |
| 7:22.217 | מבזה |  |
| 7:22.218 | על |  |
| 7:22.219 | כבודו |  |
| 7:22.220 | דאמר |  |
| 7:22.221 | ר | ' |
| 7:22.222 | חנינה |  |
| 7:22.223 | כל |  |
| 7:22.224 | המתכבד |  |
| 7:22.225 | בקלון |  |
| 7:22.226 | חבירו |  |
| 7:22.227 | אין |  |
| 7:22.228 | לו |  |
| 7:22.229 | חלק |  |
| 7:22.230 | לעה"ב | . |
| 7:22.231 | בכבודו |  |
| 7:22.232 | של |  |
| 7:22.233 | מקום |  |
| 7:22.234 | עאכ"ו | . |
| 7:22.235 | ומה |  |
| 7:22.236 | כתיב |  |
| 7:22.237 | אחריו | . |
| 7:22.238 | מה |  |
| 7:22.239 | רב |  |
| 7:22.240 | טובך |  |
| 7:22.241 | אשר |  |
| 7:22.242 | צפנת |  |
| 7:22.243 | ליראיך | . |
| 7:22.244 | ליראיך |  |
| 7:22.245 | ולא |  |
| 7:22.246 | לבוזים |  |
| 7:22.247 | את |  |
| 7:22.248 | מוראך |  |
| 7:22.249 | הרב |  |
| 7:22.250 | אל |  |
| 7:22.251 | יהיו |  |
| 7:22.252 | במה |  |
| 7:22.253 | רב |  |
| 7:22.254 | טובך |  |
| 7:22.255 | ע"כ | . |
| 7:22.256 | ובכתובות |  |
| 7:22.257 | פ | ' |
| 7:22.258 | המדיר |  |
| 7:22.259 | אהא |  |
| 7:22.260 | דתנן |  |
| 7:22.261 | על |  |
| 7:22.262 | מנת |  |
| 7:22.263 | שתאמרי |  |
| 7:22.264 | לפלוני |  |
| 7:22.265 | מה |  |
| 7:22.266 | שאמרתי |  |
| 7:22.267 | לך |  |
| 7:22.268 | או |  |
| 7:22.269 | מה |  |
| 7:22.270 | שאמרת |  |
| 7:22.271 | לי | , |
| 7:22.272 | פריך |  |
| 7:22.273 | בגמרא |  |
| 7:22.274 | ותאמר | ? |
| 7:22.275 | ופריק |  |
| 7:22.276 | דברים |  |
| 7:22.277 | של |  |
| 7:22.278 | קלון | . |
| 7:22.279 | ופי | ' |
| 7:22.280 | הרמב״ם |  |
| 7:22.281 | שם |  |
| 7:22.282 | שמכריחה |  |
| 7:22.283 | לומר |  |
| 7:22.284 | לאחר |  |
| 7:22.285 | מענייני |  |
| 7:22.286 | תשמיש |  |
| 7:22.287 | וזהו |  |
| 7:22.288 | אמרם |  |
| 7:22.289 | דברים |  |
| 7:22.290 | של |  |
| 7:22.291 | קלון | . |
| 7:22.292 | הרי |  |
| 7:22.293 | דרבותי | ' |
| 7:22.294 | ז"ל |  |
| 7:22.295 | בעלי |  |
| 7:22.296 | הגמרא |  |
| 7:22.297 | קראו |  |
| 7:22.298 | לעסקי |  |
| 7:22.299 | תשמיש |  |
| 7:22.300 | דברים |  |
| 7:22.301 | של |  |
| 7:22.302 | קלון | . |
| 7:22.303 | ודברים |  |
| 7:22.304 | כאלה |  |
| 7:22.305 | הם |  |
| 7:22.306 | מהמפורסמות |  |
| 7:22.307 | שאין |  |
| 7:22.308 | צריך |  |
| 7:22.309 | ראיה |  |
| 7:22.310 | רק |  |
| 7:22.311 | למתעקש | ! |
| 7:22.312 | ואיה |  |
| 7:22.313 | בוז |  |
| 7:22.314 | וחרפה |  |
| 7:22.315 | וקלון |  |
| 7:22.316 | יותר |  |
| 7:22.317 | מזה |  |
| 7:22.318 | ליחס |  |
| 7:22.319 | לאלהים |  |
| 7:22.320 | אשר |  |
| 7:22.321 | נעבוד |  |
| 7:22.322 | ונשתחוה |  |
| 7:22.323 | לו |  |
| 7:22.324 | אברי |  |
| 7:22.325 | תשמיש | . |
| 7:22.326 | גיד |  |
| 7:22.327 | וביצים | , |
| 7:22.328 | ואשה | , |
| 7:22.329 | ורחם |  |
| 7:22.330 | עיין |  |
| 7:22.331 | זהר |  |
| 7:22.332 | פרשת |  |
| 7:22.333 | בהר |  |
| 7:22.334 | דף |  |
| 7:22.335 | ק"ט |  |
| 7:22.336 | וסוף |  |
| 7:22.337 | אדרא |  |
| 7:22.338 | דף |  |
| 7:22.339 | רצ"ו |  |
| 7:22.340 | ע"א | ) |
| 7:22.341 | והשתוקקות |  |
| 7:22.342 | לחבקה |  |
| 7:22.343 | ולנשקה |  |
| 7:22.344 | בהיותה |  |
| 7:22.345 | מתקשטת |  |
| 7:22.346 | לפניו | . |
| 7:22.347 | ובהזדווגו |  |
| 7:22.348 | עמה |  |
| 7:22.349 | עושה |  |
| 7:22.350 | לה |  |
| 7:22.351 | נחת |  |
| 7:22.352 | רוח |  |
| 7:22.353 | ברחמה |  |
| 7:22.354 | עד |  |
| 7:22.355 | שהפיטן |  |
| 7:22.356 | מהם |  |
| 7:22.357 | בעזות |  |
| 7:22.358 | מצחו |  |
| 7:22.359 | אמר |  |
| 7:22.360 | יחבק |  |
| 7:22.361 | לה |  |
| 7:22.362 | בעלה |  |
| 7:22.363 | וביסודא |  |
| 7:22.364 | דילה |  |
| 7:22.365 | דעביד |  |
| 7:22.366 | ניחא |  |
| 7:22.367 | לה |  |
| 7:22.368 | יהא |  |
| 7:22.369 | כתיש |  |
| 7:22.370 | כתישין | . |
| 7:22.371 | הלה | ' |
| 7:22.372 | אלהיכם |  |
| 7:22.373 | שאתם |  |
| 7:22.374 | אומרים |  |
| 7:22.375 | שהוא |  |
| 7:22.376 | קצר |  |
| 7:22.377 | אפים |  |
| 7:22.378 | זעיר |  |
| 7:22.379 | אנפין | ) |
| 7:22.380 | תגמלו |  |
| 7:22.381 | זאת | ? |
| 7:22.382 | והלא |  |
| 7:22.383 | כבר |  |
| 7:22.384 | נאמר |  |
| 7:22.385 | בן |  |
| 7:22.386 | יכבד |  |
| 7:22.387 | אב |  |
| 7:22.388 | ועבד |  |
| 7:22.389 | אדוניו | . |
| 7:22.390 | ואם |  |
| 7:22.391 | אב |  |
| 7:22.392 | אני |  |
| 7:22.393 | איה |  |
| 7:22.394 | כבודי |  |
| 7:22.395 | ואם |  |
| 7:22.396 | אדונים |  |
| 7:22.397 | אני |  |
| 7:22.398 | איה |  |
| 7:22.399 | מוראי | ? |
| 7:22.400 | האם |  |
| 7:22.401 | כבודו |  |
| 7:22.402 | של |  |
| 7:22.403 | אב |  |
| 7:22.404 | שיספר |  |
| 7:22.405 | בנו |  |
| 7:22.406 | עליו |  |
| 7:22.407 | שהוא |  |
| 7:22.408 | מחבק |  |
| 7:22.409 | את |  |
| 7:22.410 | אשתו | . |
| 7:22.411 | ועושה |  |
| 7:22.412 | לה |  |
| 7:22.413 | נחת |  |
| 7:22.414 | רוח |  |
| 7:22.415 | ברחמה |  |
| 7:22.416 | בדישו |  |
| 7:22.417 | שם |  |
| 7:22.418 | דיש | ? |
| 7:22.419 | ואם |  |
| 7:22.420 | זהו |  |
| 7:22.421 | מוראו |  |
| 7:22.422 | של |  |
| 7:22.423 | אדון |  |
| 7:22.424 | שיקראהו |  |
| 7:22.425 | עבדו |  |
| 7:22.426 | קצר |  |
| 7:22.427 | אפים |  |
| 7:22.428 | זעיר |  |
| 7:22.429 | נאפין | ) |
| 7:22.430 | ויספר |  |
| 7:22.431 | עליו |  |
| 7:22.432 | דברי |  |
| 7:22.433 | קלון |  |
| 7:22.434 | כאלה | ? |
| 7:22.435 | למה |  |
| 7:22.436 | יגרע |  |
| 7:22.437 | כבוד |  |
| 7:22.438 | אבא |  |
| 7:22.439 | ואימא |  |
| 7:22.440 | וזעיר |  |
| 7:22.441 | אנפין |  |
| 7:22.442 | ונוקביה |  |
| 7:22.443 | שאתם |  |
| 7:22.444 | אומרים |  |
| 7:22.445 | שהוא |  |
| 7:22.446 | ה | ' |
| 7:22.447 | אלהיכם |  |
| 7:22.448 | והוא |  |
| 7:22.449 | שנגלה |  |
| 7:22.450 | בסיני |  |
| 7:22.451 | ונתן |  |
| 7:22.452 | לנו |  |
| 7:22.453 | את |  |
| 7:22.454 | התורה | . |
| 7:22.455 | מכבוד |  |
| 7:22.456 | האב |  |
| 7:22.457 | והאדון | ? |
| 7:22.458 | לא |  |
| 7:22.459 | על |  |
| 7:22.460 | הדברים |  |
| 7:22.461 | האלה |  |
| 7:22.462 | אמר |  |
| 7:22.463 | הנביא |  |
| 7:22.464 | ויחפאו |  |
| 7:22.465 | בני |  |
| 7:22.466 | ישראל |  |
| 7:22.467 | דברים |  |
| 7:22.468 | אשר |  |
| 7:22.469 | לא |  |
| 7:22.470 | כן |  |
| 7:22.471 | על |  |
| 7:22.472 | ה | ' |
| 7:22.473 | אלהיהם | ? |
| 7:22.474 | ואיך |  |
| 7:22.475 | נאמר |  |
| 7:22.476 | קדוש |  |
| 7:22.477 | קדוש |  |
| 7:22.478 | קדוש |  |
| 7:22.479 | ה | ' |
| 7:22.480 | צבאות | . |
| 7:22.481 | וגם |  |
| 7:22.482 | נאמר |  |
| 7:22.483 | שהוא |  |
| 7:22.484 | קדוש |  |
| 7:22.485 | ומשרתיו |  |
| 7:22.486 | קדושים | ! |
| 7:22.487 | ונספר |  |
| 7:22.488 | עליו |  |
| 7:22.489 | דברי |  |
| 7:22.490 | בוז |  |
| 7:22.491 | וקלון |  |
| 7:22.492 | וחרפה |  |
| 7:22.493 | כאלה | ? |
| 7:22.494 | דברים |  |
| 7:22.495 | שאין |  |
| 7:22.496 | להם |  |
| 7:22.497 | שום |  |
| 7:22.498 | שורש |  |
| 7:22.499 | ולא |  |
| 7:22.500 | שמץ |  |
| 7:22.501 | דבר | , |
| 7:22.502 | לא |  |
| 7:22.503 | בתורה |  |
| 7:22.504 | ולא |  |
| 7:22.505 | בנביאים |  |
| 7:22.506 | ולא |  |
| 7:22.507 | בכתובים |  |
| 7:22.508 | לא |  |
| 7:22.509 | בדברי |  |
| 7:22.510 | רז"ל |  |
| 7:22.511 | שקבלו |  |
| 7:22.512 | איש |  |
| 7:22.513 | מפי |  |
| 7:22.514 | איש |  |
| 7:22.515 | כמבואר |  |
| 7:22.516 | במסכת |  |
| 7:22.517 | אבות | ) |
| 7:22.518 | לא |  |
| 7:22.519 | במשנה |  |
| 7:22.520 | ולא |  |
| 7:22.521 | בתלמוד |  |
| 7:22.522 | בבלי |  |
| 7:22.523 | וירושלמי |  |
| 7:22.524 | ולא |  |
| 7:22.525 | במדרשים |  |
| 7:22.526 | האמתיים | . |
| 7:22.527 | שהרי |  |
| 7:22.528 | תוה"ק |  |
| 7:22.529 | לא |  |
| 7:22.530 | יחסה |  |
| 7:22.531 | לשי"ת |  |
| 7:22.532 | חוש |  |
| 7:22.533 | המשוש |  |
| 7:22.534 | גם |  |
| 7:22.535 | בחלק |  |
| 7:22.536 | המעולה |  |
| 7:22.537 | שבו |  |
| 7:22.538 | כמו |  |
| 7:22.539 | שכתב |  |
| 7:22.540 | הרב |  |
| 7:22.541 | המורה |  |
| 7:22.542 | ז"ל | . |
| 7:22.543 | כ"ש |  |
| 7:22.544 | וק"ו |  |
| 7:22.545 | החלק |  |
| 7:22.546 | הפוחת |  |
| 7:22.547 | והמגונה |  |
| 7:22.548 | שבו |  |
| 7:22.549 | דהיינו |  |
| 7:22.550 | המשגל | . |
| 7:22.551 | וחלילה |  |
| 7:22.552 | להאמין |  |
| 7:22.553 | כי |  |
| 7:22.554 | רבותינו |  |
| 7:22.555 | ז"ל |  |
| 7:22.556 | הקדושים |  |
| 7:22.557 | יוציאו |  |
| 7:22.558 | מפיהם |  |
| 7:22.559 | דברים |  |
| 7:22.560 | כאלה |  |
| 7:22.561 | על |  |
| 7:22.562 | ה | ' |
| 7:22.563 | אלהינו |  |
| 7:22.564 | ועל |  |
| 7:22.565 | כל |  |
| 7:22.566 | משרתיו | . |
| 7:22.567 | ולקראו |  |
| 7:22.568 | בשם |  |
| 7:22.569 | קצר |  |
| 7:22.570 | אפים |  |
| 7:22.571 | ולספר |  |
| 7:22.572 | עליו |  |
| 7:22.573 | או |  |
| 7:22.574 | על |  |
| 7:22.575 | משרתיו |  |
| 7:22.576 | ענייני |  |
| 7:22.577 | משגל |  |
| 7:22.578 | ופריה |  |
| 7:22.579 | ורביה |  |
| 7:22.580 | להוליד |  |
| 7:22.581 | נשמות |  |
| 7:22.582 | למלאכים |  |
| 7:22.583 | ולבני |  |
| 7:22.584 | תמותה | , |
| 7:22.585 | והנביא |  |
| 7:22.586 | אמר |  |
| 7:22.587 | נותן |  |
| 7:22.588 | נשמה |  |
| 7:22.589 | לעם |  |
| 7:22.590 | עליה |  |
| 7:22.591 | ורוח |  |
| 7:22.592 | להולכים |  |
| 7:22.593 | בה | , |
| 7:22.594 | ולא |  |
| 7:22.595 | מוליד | . |
| 7:22.596 | האם |  |
| 7:22.597 | על |  |
| 7:22.598 | הקבלה |  |
| 7:22.599 | החדשה |  |
| 7:22.600 | הנכריה |  |
| 7:22.601 | יאמרו |  |
| 7:22.602 | הגוים |  |
| 7:22.603 | עלינו |  |
| 7:22.604 | רק |  |
| 7:22.605 | עם |  |
| 7:22.606 | חכם |  |
| 7:22.607 | ונבון |  |
| 7:22.608 | הגוי |  |
| 7:22.609 | הגדול |  |
| 7:22.610 | הזה | ? |
| 7:22.611 | האם |  |
| 7:22.612 | יחוד |  |
| 7:22.613 | ה | ' |
| 7:22.614 | שנצטוינו |  |
| 7:22.615 | בתורה |  |
| 7:22.616 | הוא |  |
| 7:22.617 | לחשוב |  |
| 7:22.618 | את |  |
| 7:22.619 | כל |  |
| 7:22.620 | הצורות |  |
| 7:22.621 | והפרצופים |  |
| 7:22.622 | הנזכרים |  |
| 7:22.623 | באלו |  |
| 7:22.624 | הם |  |
| 7:22.625 | אחד | ? |
| 7:22.626 | האם |  |
| 7:22.627 | עליהם |  |
| 7:22.628 | אמר |  |
| 7:22.629 | דוד |  |
| 7:22.630 | המלך |  |
| 7:22.631 | ע"ה |  |
| 7:22.632 | ואדברה |  |
| 7:22.633 | בעדותיך |  |
| 7:22.634 | נגד |  |
| 7:22.635 | מלכים |  |
| 7:22.636 | ולא |  |
| 7:22.637 | אבוש | ? |
| 7:22.638 | הלא |  |
| 7:22.639 | נשכבה |  |
| 7:22.640 | בבשתינו |  |
| 7:22.641 | ותכסינו |  |
| 7:22.642 | כלמתינו |  |
| 7:22.643 | אם |  |
| 7:22.644 | נספר |  |
| 7:22.645 | באזני |  |
| 7:22.646 | ערבי |  |
| 7:22.647 | אחד |  |
| 7:22.648 | כדברים |  |
| 7:22.649 | האלה | ! |
| 7:22.650 | ובילקוט |  |
| 7:22.651 | פר | ' |
| 7:22.652 | קרח |  |
| 7:22.653 | רמז |  |
| 7:22.654 | תשנ"ו |  |
| 7:22.655 | וז"ל |  |
| 7:22.656 | חביבה |  |
| 7:22.657 | תורה | . |
| 7:22.658 | שכששאל |  |
| 7:22.659 | מלך |  |
| 7:22.660 | ישראל |  |
| 7:22.661 | לא |  |
| 7:22.662 | שאל |  |
| 7:22.663 | אלא |  |
| 7:22.664 | תורה |  |
| 7:22.665 | שנ | ' |
| 7:22.666 | טוב |  |
| 7:22.667 | אתה |  |
| 7:22.668 | ומטיב |  |
| 7:22.669 | למדני |  |
| 7:22.670 | חקיך | . |
| 7:22.671 | ואמר |  |
| 7:22.672 | סעדני |  |
| 7:22.673 | ואושעה |  |
| 7:22.674 | ואשעה |  |
| 7:22.675 | בחקיך |  |
| 7:22.676 | וכו | ' |
| 7:22.677 | שלא |  |
| 7:22.678 | אהא |  |
| 7:22.679 | למד |  |
| 7:22.680 | תורה |  |
| 7:22.681 | ואין |  |
| 7:22.682 | יצר |  |
| 7:22.683 | הרע |  |
| 7:22.684 | מניח |  |
| 7:22.685 | לי |  |
| 7:22.686 | לשנות | . |
| 7:22.687 | או |  |
| 7:22.688 | אהא |  |
| 7:22.689 | למד |  |
| 7:22.690 | ושוכח |  |
| 7:22.691 | או |  |
| 7:22.692 | אטמא |  |
| 7:22.693 | את |  |
| 7:22.694 | הטהור |  |
| 7:22.695 | ומטהר |  |
| 7:22.696 | את |  |
| 7:22.697 | הטמא |  |
| 7:22.698 | ונמצאתי |  |
| 7:22.699 | אבד |  |
| 7:22.700 | לעה"ב | . |
| 7:22.701 | או |  |
| 7:22.702 | שמא |  |
| 7:22.703 | ישאלוני |  |
| 7:22.704 | מגויי |  |
| 7:22.705 | הארץ |  |
| 7:22.706 | וממשפחות |  |
| 7:22.707 | האדמה |  |
| 7:22.708 | ואיני |  |
| 7:22.709 | יודע |  |
| 7:22.710 | להשיבן | , |
| 7:22.711 | ונמצאתי |  |
| 7:22.712 | בוש |  |
| 7:22.713 | לעיניהם | . |
| 7:22.714 | וכן |  |
| 7:22.715 | הוא |  |
| 7:22.716 | אומר |  |
| 7:22.717 | ואדברה |  |
| 7:22.718 | בעדותיך |  |
| 7:22.719 | נגד |  |
| 7:22.720 | מלכים |  |
| 7:22.721 | ולא |  |
| 7:22.722 | אבוש |  |
| 7:22.723 | וכו | ' |
| 7:22.724 | עכ"ל | . |
| 7:22.725 | הרי |  |
| 7:22.726 | בהדיא |  |
| 7:22.727 | שחוקי |  |
| 7:22.728 | תוה"ק |  |
| 7:22.729 | ידיעותיה |  |
| 7:22.730 | ומשפטיה |  |
| 7:22.731 | לא |  |
| 7:22.732 | נבוש |  |
| 7:22.733 | לפסרם |  |
| 7:22.734 | נגד |  |
| 7:22.735 | מלכי |  |
| 7:22.736 | אוה"ע | ! |
| 7:22.737 | מה |  |
| 7:22.738 | שאין |  |
| 7:22.739 | כן |  |
| 7:22.740 | יחודים |  |
| 7:22.741 | ופרצופים |  |
| 7:22.742 | וזווגים |  |
| 7:22.743 | הנז | ' |
| 7:22.744 | בזהר |  |
| 7:22.745 | וספרי |  |
| 7:22.746 | המקובלים |  |
| 7:22.747 | החדשים | . |
| 7:22.748 | ולא |  |
| 7:22.749 | לחנם |  |
| 7:22.750 | הזהירו |  |
| 7:22.751 | גדולי |  |
| 7:22.752 | חכמי |  |
| 7:22.753 | ישראל |  |
| 7:22.754 | שלא |  |
| 7:22.755 | ללמדה |  |
| 7:22.756 | ולילך |  |
| 7:22.757 | במנהגים |  |
| 7:22.758 | הנז | ' |
| 7:22.759 | בה |  |
| 7:22.760 | היפך |  |
| 7:22.761 | התלמוד |  |
| 7:22.762 | והפוסקים | , |
| 7:22.763 | הלא |  |
| 7:22.764 | הרב |  |
| 7:22.765 | מהרש"ל |  |
| 7:22.766 | ור"ת |  |
| 7:22.767 | ן | ' |
| 7:22.768 | יחיא |  |
| 7:22.769 | המובאים |  |
| 7:22.770 | בשו"ת |  |
| 7:22.771 | רביד |  |
| 7:22.772 | הזהב | , |
| 7:22.773 | גם |  |
| 7:22.774 | ר | ' |
| 7:22.775 | פרץ |  |
| 7:22.776 | הובא |  |
| 7:22.777 | בשו"ת |  |
| 7:22.778 | הריב"ש |  |
| 7:22.779 | סי | ' |
| 7:22.780 | קנ"ז |  |
| 7:22.781 | לא |  |
| 7:22.782 | היה |  |
| 7:22.783 | מדבר |  |
| 7:22.784 | ולא |  |
| 7:22.785 | מחשיב |  |
| 7:22.786 | לאותן |  |
| 7:22.787 | ספירות | , |
| 7:22.788 | פי | ' |
| 7:22.789 | לא |  |
| 7:22.790 | היה |  |
| 7:22.791 | מחשיב |  |
| 7:22.792 | לאותן |  |
| 7:22.793 | ספירות |  |
| 7:22.794 | שום |  |
| 7:22.795 | אלהות | ), |
| 7:22.796 | גם |  |
| 7:22.797 | ר | ' |
| 7:22.798 | שמשון |  |
| 7:22.799 | מקינון |  |
| 7:22.800 | מאן |  |
| 7:22.801 | לכוון |  |
| 7:22.802 | לשום |  |
| 7:22.803 | ספירה |  |
| 7:22.804 | והיה |  |
| 7:22.805 | אומר |  |
| 7:22.806 | שהוא |  |
| 7:22.807 | מתפלל |  |
| 7:22.808 | לדעת |  |
| 7:22.809 | זה |  |
| 7:22.810 | התינוק |  |
| 7:22.811 | פי | ' |
| 7:22.812 | כתינוק |  |
| 7:22.813 | זה |  |
| 7:22.814 | שמתפלל |  |
| 7:22.815 | לשוכן |  |
| 7:22.816 | שמים | ) |
| 7:22.817 | להוציא |  |
| 7:22.818 | מלב |  |
| 7:22.819 | המקובלים |  |
| 7:22.820 | וכו | ', |
| 7:22.821 | והריב"ש |  |
| 7:22.822 | ז"ל |  |
| 7:22.823 | הרגיש |  |
| 7:22.824 | בזרותה |  |
| 7:22.825 | וכתב |  |
| 7:22.826 | איני |  |
| 7:22.827 | תוקע |  |
| 7:22.828 | עצמי |  |
| 7:22.829 | ללמדה |  |
| 7:22.830 | וכו | ', |
| 7:22.831 | וכתב |  |
| 7:22.832 | עוד |  |
| 7:22.833 | וכי |  |
| 7:22.834 | הספירות |  |
| 7:22.835 | יש |  |
| 7:22.836 | להם |  |
| 7:22.837 | אלהות | ? |
| 7:22.838 | והמשיב |  |
| 7:22.839 | רבי |  |
| 7:22.840 | יוסף |  |
| 7:22.841 | ן | ' |
| 7:22.842 | שושן |  |
| 7:22.843 | כיחש |  |
| 7:22.844 | לו |  |
| 7:22.845 | והביא |  |
| 7:22.846 | לו |  |
| 7:22.847 | משל |  |
| 7:22.848 | שאינו |  |
| 7:22.849 | לפי |  |
| 7:22.850 | האמת | ! |
| 7:22.851 | כי |  |
| 7:22.852 | האמת |  |
| 7:22.853 | הגמור |  |
| 7:22.854 | שהמקובלים |  |
| 7:22.855 | סוברים |  |
| 7:22.856 | דספירות |  |
| 7:22.857 | דיושר |  |
| 7:22.858 | הם |  |
| 7:22.859 | עיקר |  |
| 7:22.860 | האלהות | . |
| 7:22.861 | והוא |  |
| 7:22.862 | ז"ל |  |
| 7:22.863 | הרגיש |  |
| 7:22.864 | בסתירה | , |
| 7:22.865 | ולכן |  |
| 7:22.866 | אמר |  |
| 7:22.867 | איני |  |
| 7:22.868 | תוקע |  |
| 7:22.869 | עצמי |  |
| 7:22.870 | ללמדה |  |
| 7:22.871 | וכו | ' |
| 7:22.872 | גם |  |
| 7:22.873 | הרב |  |
| 7:22.874 | בעל |  |
| 7:22.875 | העיקרים |  |
| 7:22.876 | פרק |  |
| 7:22.877 | כ"ח |  |
| 7:22.878 | משני |  |
| 7:22.879 | הזהיר |  |
| 7:22.880 | מאד |  |
| 7:22.881 | שלא |  |
| 7:22.882 | ללמוד |  |
| 7:22.883 | הקבלה |  |
| 7:22.884 | וז"ל |  |
| 7:22.885 | וכלל |  |
| 7:22.886 | אומר |  |
| 7:22.887 | לך |  |
| 7:22.888 | השמר |  |
| 7:22.889 | לך |  |
| 7:22.890 | ושמור |  |
| 7:22.891 | נפשך |  |
| 7:22.892 | מאד |  |
| 7:22.893 | פן |  |
| 7:22.894 | תנקש |  |
| 7:22.895 | אחריהם |  |
| 7:22.896 | ותלכד |  |
| 7:22.897 | במצודתם | , |
| 7:22.898 | כי |  |
| 7:22.899 | הם |  |
| 7:22.900 | עוזבים |  |
| 7:22.901 | ארחות |  |
| 7:22.902 | יושר |  |
| 7:22.903 | ללכת |  |
| 7:22.904 | בדרכי |  |
| 7:22.905 | חשך |  |
| 7:22.906 | לעבוד |  |
| 7:22.907 | אלהים |  |
| 7:22.908 | אחרים | ) |
| 7:22.909 | כי |  |
| 7:22.910 | לא |  |
| 7:22.911 | ידעו |  |
| 7:22.912 | ולא |  |
| 7:22.913 | יבינו |  |
| 7:22.914 | בחשכה |  |
| 7:22.915 | יתהלכו |  |
| 7:22.916 | העוסקים |  |
| 7:22.917 | בקבלה |  |
| 7:22.918 | מדעת |  |
| 7:22.919 | עצמם |  |
| 7:22.920 | מזולת |  |
| 7:22.921 | קבלה |  |
| 7:22.922 | מפי |  |
| 7:22.923 | חכם |  |
| 7:22.924 | מקובל |  |
| 7:22.925 | עכ״ל | . |
| 7:22.926 | וכבר |  |
| 7:22.927 | נתבאר |  |
| 7:22.928 | לעיל |  |
| 7:22.929 | סי | ' |
| 7:22.930 | י"ח |  |
| 7:22.931 | משם |  |
| 7:22.932 | מהר"י |  |
| 7:22.933 | צאהרי |  |
| 7:22.934 | שאמונת |  |
| 7:22.935 | הקבלה |  |
| 7:22.936 | הנכריה |  |
| 7:22.937 | החדשה |  |
| 7:22.938 | הזו | , |
| 7:22.939 | חדשה |  |
| 7:22.940 | היא |  |
| 7:22.941 | בארץ |  |
| 7:22.942 | התימן |  |
| 7:22.943 | על |  |
| 7:22.944 | ידי |  |
| 7:22.945 | הספרים |  |
| 7:22.946 | שבאו |  |
| 7:22.947 | בה |  |
| 7:22.948 | מאותו |  |
| 7:22.949 | זמן |  |
| 7:22.950 | כעדותו |  |
| 7:22.951 | של |  |
| 7:22.952 | מהר"י |  |
| 7:22.953 | הנז | '. |
| 7:22.954 | ומעולם |  |
| 7:22.955 | לא |  |
| 7:22.956 | נמצא |  |
| 7:22.957 | בתימן |  |
| 7:22.958 | איש |  |
| 7:22.959 | מקובל |  |
| 7:22.960 | מפה |  |
| 7:22.961 | אל |  |
| 7:22.962 | פה | , |
| 7:22.963 | וספרי |  |
| 7:22.964 | קדמונינו |  |
| 7:22.965 | הנשארים |  |
| 7:22.966 | בידינו |  |
| 7:22.967 | יעידון |  |
| 7:22.968 | על |  |
| 7:22.969 | זה | . |
| 7:22.970 | וכפי |  |
| 7:22.971 | דברי |  |
| 7:22.972 | הרב |  |
| 7:22.973 | שם |  |
| 7:22.974 | הגדולים |  |
| 7:22.975 | במערכת |  |
| 7:22.976 | ספרים |  |
| 7:22.977 | ערך |  |
| 7:22.978 | זהר |  |
| 7:22.979 | גם |  |
| 7:22.980 | בכל |  |
| 7:22.981 | ישראל |  |
| 7:22.982 | חדשה |  |
| 7:22.983 | היא |  |
| 7:22.984 | וז"ל |  |
| 7:22.985 | ראיתי |  |
| 7:22.986 | כתוב |  |
| 7:22.987 | מהר | ' |
| 7:22.988 | מהר"א |  |
| 7:22.989 | רויגו |  |
| 7:22.990 | וז"ל |  |
| 7:22.991 | מצאתי |  |
| 7:22.992 | בזהר |  |
| 7:22.993 | כ"י |  |
| 7:22.994 | ישן |  |
| 7:22.995 | נושן |  |
| 7:22.996 | אצל |  |
| 7:22.997 | מורי |  |
| 7:22.998 | מהר"ם |  |
| 7:22.999 | זכות |  |
| 7:22.1000 | וז"ל |  |
| 7:22.1001 | מצאתי |  |
| 7:22.1002 | כתוב |  |
| 7:22.1003 | באמת |  |
| 7:22.1004 | כי |  |
| 7:22.1005 | ראש |  |
| 7:22.1006 | המקובלים |  |
| 7:22.1007 | רבי |  |
| 7:22.1008 | נחוניא |  |
| 7:22.1009 | בן |  |
| 7:22.1010 | הקנה |  |
| 7:22.1011 | והוא |  |
| 7:22.1012 | חיבר |  |
| 7:22.1013 | ס | ' |
| 7:22.1014 | הבהיר | , |
| 7:22.1015 | ואחריו |  |
| 7:22.1016 | רשב"י |  |
| 7:22.1017 | עשה |  |
| 7:22.1018 | ס | ' |
| 7:22.1019 | הזהר |  |
| 7:22.1020 | וחיבר |  |
| 7:22.1021 | בו |  |
| 7:22.1022 | כמה |  |
| 7:22.1023 | חיבורים |  |
| 7:22.1024 | כמו |  |
| 7:22.1025 | התקונים | , |
| 7:22.1026 | וכשמת |  |
| 7:22.1027 | רשב"י |  |
| 7:22.1028 | ור"א |  |
| 7:22.1029 | וכל |  |
| 7:22.1030 | הדור |  |
| 7:22.1031 | ההוא |  |
| 7:22.1032 | אבדה |  |
| 7:22.1033 | חכמת |  |
| 7:22.1034 | הקבלה | , |
| 7:22.1035 | עד |  |
| 7:22.1036 | שהקרה |  |
| 7:22.1037 | ה | ' |
| 7:22.1038 | לפני |  |
| 7:22.1039 | מלך |  |
| 7:22.1040 | ממלכי |  |
| 7:22.1041 | מזרח |  |
| 7:22.1042 | שצוה |  |
| 7:22.1043 | לחפור |  |
| 7:22.1044 | במקום |  |
| 7:22.1045 | אחד |  |
| 7:22.1046 | על |  |
| 7:22.1047 | עסקי |  |
| 7:22.1048 | ממון | , |
| 7:22.1049 | ונמצא |  |
| 7:22.1050 | שם |  |
| 7:22.1051 | ארון |  |
| 7:22.1052 | אחד | , |
| 7:22.1053 | ובו |  |
| 7:22.1054 | ספר |  |
| 7:22.1055 | הזהר |  |
| 7:22.1056 | ושלח |  |
| 7:22.1057 | לחכמי |  |
| 7:22.1058 | אדום |  |
| 7:22.1059 | ולא |  |
| 7:22.1060 | ידעו |  |
| 7:22.1061 | ולא |  |
| 7:22.1062 | יבינו | , |
| 7:22.1063 | שלח |  |
| 7:22.1064 | אחר |  |
| 7:22.1065 | היהודים |  |
| 7:22.1066 | באו |  |
| 7:22.1067 | אצלו | , |
| 7:22.1068 | וראו |  |
| 7:22.1069 | הספר |  |
| 7:22.1070 | ואמרו |  |
| 7:22.1071 | לו |  |
| 7:22.1072 | אדונינו |  |
| 7:22.1073 | המלך | , |
| 7:22.1074 | זה |  |
| 7:22.1075 | הס | ' |
| 7:22.1076 | עשאו |  |
| 7:22.1077 | חכם |  |
| 7:22.1078 | אחד |  |
| 7:22.1079 | והוא |  |
| 7:22.1080 | עמוק |  |
| 7:22.1081 | מאד |  |
| 7:22.1082 | ואין |  |
| 7:22.1083 | אנו |  |
| 7:22.1084 | מבינים |  |
| 7:22.1085 | אותו | , |
| 7:22.1086 | אמר |  |
| 7:22.1087 | להם |  |
| 7:22.1088 | וכי |  |
| 7:22.1089 | אין |  |
| 7:22.1090 | יהודי |  |
| 7:22.1091 | בעולם |  |
| 7:22.1092 | שמבין |  |
| 7:22.1093 | אותו | ? |
| 7:22.1094 | אמרו |  |
| 7:22.1095 | לו |  |
| 7:22.1096 | יש |  |
| 7:22.1097 | במדינת |  |
| 7:22.1098 | טוליטולא | , |
| 7:22.1099 | והמלך |  |
| 7:22.1100 | שלח |  |
| 7:22.1101 | הספרים |  |
| 7:22.1102 | עם |  |
| 7:22.1103 | גבוריו |  |
| 7:22.1104 | לטוליטולה | , |
| 7:22.1105 | וכשראוהו |  |
| 7:22.1106 | חכמי |  |
| 7:22.1107 | טוליטולה |  |
| 7:22.1108 | שמחו |  |
| 7:22.1109 | בו |  |
| 7:22.1110 | שמחה |  |
| 7:22.1111 | גדולה | , |
| 7:22.1112 | ושלחו |  |
| 7:22.1113 | למלך |  |
| 7:22.1114 | מתנות |  |
| 7:22.1115 | רבות | , |
| 7:22.1116 | ומשם |  |
| 7:22.1117 | נתפרסמה |  |
| 7:22.1118 | הקבלה |  |
| 7:22.1119 | בישראל |  |
| 7:22.1120 | ע"כ |  |
| 7:22.1121 | מצאתי |  |
| 7:22.1122 | מהר | ' |
| 7:22.1123 | הנז | ' |
| 7:22.1124 | עכ"ל | . |
| 7:22.1125 | וכדברים |  |
| 7:22.1126 | האלה |  |
| 7:22.1127 | כתבו |  |
| 7:22.1128 | חכמי |  |
| 7:22.1129 | ירושלם |  |
| 7:22.1130 | אלינו | , |
| 7:22.1131 | והרואה |  |
| 7:22.1132 | יבין |  |
| 7:22.1133 | דבר |  |
| 7:22.1134 | לאשורו |  |
| 7:22.1135 | ממה |  |
| 7:22.1136 | שכתב |  |
| 7:22.1137 | שאחר |  |
| 7:22.1138 | מות |  |
| 7:22.1139 | רשב"י |  |
| 7:22.1140 | אבדה |  |
| 7:22.1141 | חכמת |  |
| 7:22.1142 | הקבלה |  |
| 7:22.1143 | ונשכלה |  |
| 7:22.1144 | מישראל | , |
| 7:22.1145 | ולא |  |
| 7:22.1146 | למדו |  |
| 7:22.1147 | דורות |  |
| 7:22.1148 | האחרונים |  |
| 7:22.1149 | רק |  |
| 7:22.1150 | מתוך |  |
| 7:22.1151 | הספר |  |
| 7:22.1152 | המזויף |  |
| 7:22.1153 | שקבלו |  |
| 7:22.1154 | אותו |  |
| 7:22.1155 | חכמי |  |
| 7:22.1156 | טוליטולה |  |
| 7:22.1157 | מיד |  |
| 7:22.1158 | המלך |  |
| 7:22.1159 | הנכרי |  |
| 7:22.1160 | ושמחו |  |
| 7:22.1161 | בו |  |
| 7:22.1162 | שמחה |  |
| 7:22.1163 | גדולה | , |
| 7:22.1164 | ועל |  |
| 7:22.1165 | ידו |  |
| 7:22.1166 | המירו |  |
| 7:22.1167 | את |  |
| 7:22.1168 | כבודם |  |
| 7:22.1169 | ולא |  |
| 7:22.1170 | קיבלו |  |
| 7:22.1171 | מפי |  |
| 7:22.1172 | חכם |  |
| 7:22.1173 | מקובל | , |
| 7:22.1174 | גם |  |
| 7:22.1175 | כל |  |
| 7:22.1176 | האנשים |  |
| 7:22.1177 | שראינו |  |
| 7:22.1178 | אותם |  |
| 7:22.1179 | הולכים |  |
| 7:22.1180 | בדרכי |  |
| 7:22.1181 | הקבלה |  |
| 7:22.1182 | הלזו |  |
| 7:22.1183 | כל |  |
| 7:22.1184 | ידיעתם |  |
| 7:22.1185 | ואמונתם |  |
| 7:22.1186 | בה |  |
| 7:22.1187 | מפי |  |
| 7:22.1188 | הספרים |  |
| 7:22.1189 | ולא |  |
| 7:22.1190 | מפי |  |
| 7:22.1191 | סופרים | , |
| 7:22.1192 | וכפי |  |
| 7:22.1193 | דברי |  |
| 7:22.1194 | הרמב"ן |  |
| 7:22.1195 | ודאי |  |
| 7:22.1196 | שהלומד |  |
| 7:22.1197 | מתוך |  |
| 7:22.1198 | הספרים |  |
| 7:22.1199 | ישגה |  |
| 7:22.1200 | בה |  |
| 7:22.1201 | ויבא |  |
| 7:22.1202 | לידי |  |
| 7:22.1203 | טעות |  |
| 7:22.1204 | להאמין |  |
| 7:22.1205 | ברשויות |  |
| 7:22.1206 | הרבה |  |
| 7:22.1207 | ולעבוד |  |
| 7:22.1208 | קצר |  |
| 7:22.1209 | אפים |  |
| 7:22.1210 | נברא | , |
| 7:22.1211 | אשר |  |
| 7:22.1212 | לפי |  |
| 7:22.1213 | דבריהם |  |
| 7:22.1214 | נשמתו |  |
| 7:22.1215 | חלק |  |
| 7:22.1216 | אלוה |  |
| 7:22.1217 | ממעל |  |
| 7:22.1218 | והמה |  |
| 7:22.1219 | עובדים |  |
| 7:22.1220 | לחלק |  |
| 7:22.1221 | ההוא |  |
| 7:22.1222 | שבזעיר |  |
| 7:22.1223 | אנפין |  |
| 7:22.1224 | בחשבם |  |
| 7:22.1225 | כי |  |
| 7:22.1226 | האלהים |  |
| 7:22.1227 | האמצעים |  |
| 7:22.1228 | אבא |  |
| 7:22.1229 | ואימא |  |
| 7:22.1230 | נתנו |  |
| 7:22.1231 | הכח |  |
| 7:22.1232 | והממשלה |  |
| 7:22.1233 | לשלוט |  |
| 7:22.1234 | בתחתונים | . |
| 7:22.1235 | ואמר |  |
| 7:22.1236 | לו |  |
| 7:22.1237 | בני |  |
| 7:22.1238 | אתה |  |
| 7:22.1239 | זהר |  |
| 7:22.1240 | בלק |  |
| 7:22.1241 | דף |  |
| 7:22.1242 | קצ"א | ) |
| 7:22.1243 | וצוו |  |
| 7:22.1244 | לעבדו |  |
| 7:22.1245 | ולברך |  |
| 7:22.1246 | בשמו | , |
| 7:22.1247 | וקראו |  |
| 7:22.1248 | לו |  |
| 7:22.1249 | בשם |  |
| 7:22.1250 | ידוד |  |
| 7:22.1251 | כשם |  |
| 7:22.1252 | ה | ' |
| 7:22.1253 | אלהינו |  |
| 7:22.1254 | כמו |  |
| 7:22.1255 | שעשו |  |
| 7:22.1256 | דורו |  |
| 7:22.1257 | של |  |
| 7:22.1258 | אנוש |  |
| 7:22.1259 | כדלעיל |  |
| 7:22.1260 | סי | ' |
| 7:22.1261 | נ"ו |  |
| 7:22.1262 | ונ"ז | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

