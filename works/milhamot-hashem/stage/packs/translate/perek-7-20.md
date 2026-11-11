# Translation pack: milhamot-hashem perek-7-20 (Milchamot ha-Shem 7:25–7:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-20`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
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


## Milchamot ha-Shem 7:25

Hebrew (Wikisource): צז) וראיתי לאחד קדוש מחכמי ימי הבינים הוא מהר"י צאהרי אשר חי בתימן בעת התפשטות ספרי הזהר והקבלה הנכריה החדשה בארץ התימן כמו שכ' בפרשת לך לך שהן חדשים מקרוב באו. ובפרשת בראשית ביאר שיש שתי דיעות ביחוד השי"ת וז"ל ואחר הקדמתי זאת צריך אני לגלות לך מעט כי בזה הענין (ר"ל יחוד השם) יש שתי דעות, האחת דעת הרמב"ם והנמשכים אחריו שלא הגיע אליהם דבר מדברי הקבלה (הנכריה החדשה). וסברתם ביחוד השם על דרך שלילה רצוני לומר לסלק הספק מלבו ולסלק מהשם ית' כל החסרונות ואז הוא אחד פשוט ועוד אמרו (גמרא ברכות דף י"ג ע"ב) צריך להאריך בדלת של אחד כדי שימליכהו על ארבע פנות העולם כולל הדברים לא תמצא מגמתם מגעת בכל מה שהביאו בדיבריהם אלא עד כאן. והדעת האחרת הם בעלי הקבלה העושים את כל מלאכת הקדש אמרו כי סוד היחוד הוא ליחד את השם בכל עשר ספירות. והקדושה גם כן צריך שיכוין בכל עשר ספירות שהם מערכת האלהות, וכל המצות מתפשטות מן העשר רצוני לומר שכל מצוה ומצוה יש לה סוד באחת מן העשר כמו שאמר בוצנא קדישא המאיר לארץ ולדרים רשב"י בפ' וישמע יתרו, כל פקודי אוריתא מתאחדן בגופא דמלכא מנהון ברישא ומנהון בגופא, ומנהון בידי מלכא ומנהון ברגלוי. ולית מאן דנפיק מגופא דמלכא לבר וכו'. ואתה המעיין אל תתמה במה שאמרו רישא ידא ורגלא שהרי כל אלו רומזים אל הספירות כמו שאצייר לך עתה שיח השדה הקדושה והטהורה העליונה כפי מה שסידר הרב שערי אורה אחת לאחת למצוא חשבון. אבל כי חג' ראשונות רומזים לראש. גדולה גבורה זרועות עולם. תפארת לשדרה וללב. נצח והוד תרין ירכין. ויסוד דא ברית. ומלכות היא מקבלת את הכל ומשפעת לכל התחתונים. וזהו סוד היחוד והקדושה, וצריך לחבר את האהל להיות אחד השני שמות ואהדונה"א. נמצאת למד כשתאמר מלת אחד, תכוין לתשע ספירות, אל"ף וחי"ת, ומלכות רומזת לעשירית, שהיא דל"ת, וחסוד למען יאריך ימים על ממלכתו, על שם הבורא מלאך וגלגל ואופן עכ"ל. ראה אבי גם ראה איך החכם הנז' מהר"י צאהרי הרגיש וידע שהם שתי דעות מתנגדות ביחוד השי"ת, אך ורק כי חשב את הראשונים חכמי התלמוד שאמרו להאריך בדל"ת דאחד כדי להמליכו בארבע פינות העולם לחסידי מדע. שלא ידעו בקבלה הנכריה החדשה ולא הגיע אליהם דבר ממנה, ולא ידעו ליחד את השם עם ברואיו הספירות המדומות שעשו אותם גופים ופרצופים לאלהינו ב"ה והוא נשמה להם. ואדרבא נהפוך הוא שאיסור גמור הוא לכלול ולשתף עם השי"ת שום נברא כמו שנתבאר לעיל סי' ע"ד. וסי' פ"ב בשם מדרש רבה, והטעם כי הוא ית' נעלה ומרומם ושגיא כה, ודי לו בעצמותו ואין צריך לבריותיו להשתתף עם שום אחד מהם. שכל בריותיו צריכים לו, והוא אינו צריך להם להעזר בהם ח"ו כי הוא אמר ויהי הוא לבדו צוה ויעמוד. ולא ירד הרב הנז' לעומק דברי ה רס"ג והרמב"ם שהוא עסוק בהם להעתיק דבריהם והוא מחבבן, ועשה אותן כעין הצעה והקדמה לאמתת יחוד השי"ת, והרגיש שהם שתי סברות. ולא הרגיש הסתירה בפרק ג' לראשון שהעתיק מקצת דבריו, שהם מתנגדים לגמרי לשטת הקבלה הנכריה החדשה, רק חשב אותם לחסידי מדע. שלא ידעו לצרף ולשתף ולכלול את הצורות והפרצופים עם אמתתו ית' וליחדו אמהם דוקא. ולא ידעו ולא יבינו כי הרס"ג והרמב"ם ורבי יהודה הלוי ובעל חובת הלבבות נלחמים המה בחזקת היד עם דעות הנכריות הזרות שהתחילו להנץ בין היהודים מזמן הגאונים, ועדיין היו בעל פה ולא הוחקו על ספר. עד שיצא הספר הנכרי הזהר המבזה את תורה שבעל פה המשנה והתלמוד, וקורא להם אבן משכית, ועץ הדעת טוב ורע. וקליפה, ושפחה כי תירש גברתה, נגד דברי דוד המלך ע"ה שאמר עליהם הנחמדים מזהב ומפז רב ומתוקים מדבש ומנופת צופים, והסית והדיח את הרבים לעבוד לאל אחר קצר אפים נברא, כמו שכתב הר' יושר לבב לעבדו אע"פ שהוא נברא, לפי שהאמינו באמונה שלימה כי מרשב"י יצאו הדברים ההם. והסבה שגרמה להם זה שלא שמו אל לבם ונגד עיניהם דברי הראשונים בעלי הכללים המעמיקים בתלמוד, המה דברי הרא"ש ז"ל שהבאתי לעיל סי' מ"ב שכל תוספתא או בריתא שלא נתפשטה עד אחר סידור הגמרא לאו דסמכא היא. דמסתמא מחברי התלמוד חקרו ודרשו לידע כל הספרים שנכתבו על דעת חכמים וביררו אותם שהם בר סמכא. וכל תוספתא או ברייתא שהיא היפך הגמרא לאו דסמכא היא דודאי משבשתא היא או מזויפת ע"ש.

| id | word | punctuation after |
|---|---|---|
| 7:25.1 | צז | ) |
| 7:25.2 | וראיתי |  |
| 7:25.3 | לאחד |  |
| 7:25.4 | קדוש |  |
| 7:25.5 | מחכמי |  |
| 7:25.6 | ימי |  |
| 7:25.7 | הבינים |  |
| 7:25.8 | הוא |  |
| 7:25.9 | מהר"י |  |
| 7:25.10 | צאהרי |  |
| 7:25.11 | אשר |  |
| 7:25.12 | חי |  |
| 7:25.13 | בתימן |  |
| 7:25.14 | בעת |  |
| 7:25.15 | התפשטות |  |
| 7:25.16 | ספרי |  |
| 7:25.17 | הזהר |  |
| 7:25.18 | והקבלה |  |
| 7:25.19 | הנכריה |  |
| 7:25.20 | החדשה |  |
| 7:25.21 | בארץ |  |
| 7:25.22 | התימן |  |
| 7:25.23 | כמו |  |
| 7:25.24 | שכ | ' |
| 7:25.25 | בפרשת |  |
| 7:25.26 | לך |  |
| 7:25.27 | לך |  |
| 7:25.28 | שהן |  |
| 7:25.29 | חדשים |  |
| 7:25.30 | מקרוב |  |
| 7:25.31 | באו | . |
| 7:25.32 | ובפרשת |  |
| 7:25.33 | בראשית |  |
| 7:25.34 | ביאר |  |
| 7:25.35 | שיש |  |
| 7:25.36 | שתי |  |
| 7:25.37 | דיעות |  |
| 7:25.38 | ביחוד |  |
| 7:25.39 | השי"ת |  |
| 7:25.40 | וז"ל |  |
| 7:25.41 | ואחר |  |
| 7:25.42 | הקדמתי |  |
| 7:25.43 | זאת |  |
| 7:25.44 | צריך |  |
| 7:25.45 | אני |  |
| 7:25.46 | לגלות |  |
| 7:25.47 | לך |  |
| 7:25.48 | מעט |  |
| 7:25.49 | כי |  |
| 7:25.50 | בזה |  |
| 7:25.51 | הענין |  |
| 7:25.52 | ר"ל |  |
| 7:25.53 | יחוד |  |
| 7:25.54 | השם | ) |
| 7:25.55 | יש |  |
| 7:25.56 | שתי |  |
| 7:25.57 | דעות | , |
| 7:25.58 | האחת |  |
| 7:25.59 | דעת |  |
| 7:25.60 | הרמב"ם |  |
| 7:25.61 | והנמשכים |  |
| 7:25.62 | אחריו |  |
| 7:25.63 | שלא |  |
| 7:25.64 | הגיע |  |
| 7:25.65 | אליהם |  |
| 7:25.66 | דבר |  |
| 7:25.67 | מדברי |  |
| 7:25.68 | הקבלה |  |
| 7:25.69 | הנכריה |  |
| 7:25.70 | החדשה | ). |
| 7:25.71 | וסברתם |  |
| 7:25.72 | ביחוד |  |
| 7:25.73 | השם |  |
| 7:25.74 | על |  |
| 7:25.75 | דרך |  |
| 7:25.76 | שלילה |  |
| 7:25.77 | רצוני |  |
| 7:25.78 | לומר |  |
| 7:25.79 | לסלק |  |
| 7:25.80 | הספק |  |
| 7:25.81 | מלבו |  |
| 7:25.82 | ולסלק |  |
| 7:25.83 | מהשם |  |
| 7:25.84 | ית | ' |
| 7:25.85 | כל |  |
| 7:25.86 | החסרונות |  |
| 7:25.87 | ואז |  |
| 7:25.88 | הוא |  |
| 7:25.89 | אחד |  |
| 7:25.90 | פשוט |  |
| 7:25.91 | ועוד |  |
| 7:25.92 | אמרו |  |
| 7:25.93 | גמרא |  |
| 7:25.94 | ברכות |  |
| 7:25.95 | דף |  |
| 7:25.96 | י"ג |  |
| 7:25.97 | ע"ב | ) |
| 7:25.98 | צריך |  |
| 7:25.99 | להאריך |  |
| 7:25.100 | בדלת |  |
| 7:25.101 | של |  |
| 7:25.102 | אחד |  |
| 7:25.103 | כדי |  |
| 7:25.104 | שימליכהו |  |
| 7:25.105 | על |  |
| 7:25.106 | ארבע |  |
| 7:25.107 | פנות |  |
| 7:25.108 | העולם |  |
| 7:25.109 | כולל |  |
| 7:25.110 | הדברים |  |
| 7:25.111 | לא |  |
| 7:25.112 | תמצא |  |
| 7:25.113 | מגמתם |  |
| 7:25.114 | מגעת |  |
| 7:25.115 | בכל |  |
| 7:25.116 | מה |  |
| 7:25.117 | שהביאו |  |
| 7:25.118 | בדיבריהם |  |
| 7:25.119 | אלא |  |
| 7:25.120 | עד |  |
| 7:25.121 | כאן | . |
| 7:25.122 | והדעת |  |
| 7:25.123 | האחרת |  |
| 7:25.124 | הם |  |
| 7:25.125 | בעלי |  |
| 7:25.126 | הקבלה |  |
| 7:25.127 | העושים |  |
| 7:25.128 | את |  |
| 7:25.129 | כל |  |
| 7:25.130 | מלאכת |  |
| 7:25.131 | הקדש |  |
| 7:25.132 | אמרו |  |
| 7:25.133 | כי |  |
| 7:25.134 | סוד |  |
| 7:25.135 | היחוד |  |
| 7:25.136 | הוא |  |
| 7:25.137 | ליחד |  |
| 7:25.138 | את |  |
| 7:25.139 | השם |  |
| 7:25.140 | בכל |  |
| 7:25.141 | עשר |  |
| 7:25.142 | ספירות | . |
| 7:25.143 | והקדושה |  |
| 7:25.144 | גם |  |
| 7:25.145 | כן |  |
| 7:25.146 | צריך |  |
| 7:25.147 | שיכוין |  |
| 7:25.148 | בכל |  |
| 7:25.149 | עשר |  |
| 7:25.150 | ספירות |  |
| 7:25.151 | שהם |  |
| 7:25.152 | מערכת |  |
| 7:25.153 | האלהות | , |
| 7:25.154 | וכל |  |
| 7:25.155 | המצות |  |
| 7:25.156 | מתפשטות |  |
| 7:25.157 | מן |  |
| 7:25.158 | העשר |  |
| 7:25.159 | רצוני |  |
| 7:25.160 | לומר |  |
| 7:25.161 | שכל |  |
| 7:25.162 | מצוה |  |
| 7:25.163 | ומצוה |  |
| 7:25.164 | יש |  |
| 7:25.165 | לה |  |
| 7:25.166 | סוד |  |
| 7:25.167 | באחת |  |
| 7:25.168 | מן |  |
| 7:25.169 | העשר |  |
| 7:25.170 | כמו |  |
| 7:25.171 | שאמר |  |
| 7:25.172 | בוצנא |  |
| 7:25.173 | קדישא |  |
| 7:25.174 | המאיר |  |
| 7:25.175 | לארץ |  |
| 7:25.176 | ולדרים |  |
| 7:25.177 | רשב"י |  |
| 7:25.178 | בפ | ' |
| 7:25.179 | וישמע |  |
| 7:25.180 | יתרו | , |
| 7:25.181 | כל |  |
| 7:25.182 | פקודי |  |
| 7:25.183 | אוריתא |  |
| 7:25.184 | מתאחדן |  |
| 7:25.185 | בגופא |  |
| 7:25.186 | דמלכא |  |
| 7:25.187 | מנהון |  |
| 7:25.188 | ברישא |  |
| 7:25.189 | ומנהון |  |
| 7:25.190 | בגופא | , |
| 7:25.191 | ומנהון |  |
| 7:25.192 | בידי |  |
| 7:25.193 | מלכא |  |
| 7:25.194 | ומנהון |  |
| 7:25.195 | ברגלוי | . |
| 7:25.196 | ולית |  |
| 7:25.197 | מאן |  |
| 7:25.198 | דנפיק |  |
| 7:25.199 | מגופא |  |
| 7:25.200 | דמלכא |  |
| 7:25.201 | לבר |  |
| 7:25.202 | וכו | '. |
| 7:25.203 | ואתה |  |
| 7:25.204 | המעיין |  |
| 7:25.205 | אל |  |
| 7:25.206 | תתמה |  |
| 7:25.207 | במה |  |
| 7:25.208 | שאמרו |  |
| 7:25.209 | רישא |  |
| 7:25.210 | ידא |  |
| 7:25.211 | ורגלא |  |
| 7:25.212 | שהרי |  |
| 7:25.213 | כל |  |
| 7:25.214 | אלו |  |
| 7:25.215 | רומזים |  |
| 7:25.216 | אל |  |
| 7:25.217 | הספירות |  |
| 7:25.218 | כמו |  |
| 7:25.219 | שאצייר |  |
| 7:25.220 | לך |  |
| 7:25.221 | עתה |  |
| 7:25.222 | שיח |  |
| 7:25.223 | השדה |  |
| 7:25.224 | הקדושה |  |
| 7:25.225 | והטהורה |  |
| 7:25.226 | העליונה |  |
| 7:25.227 | כפי |  |
| 7:25.228 | מה |  |
| 7:25.229 | שסידר |  |
| 7:25.230 | הרב |  |
| 7:25.231 | שערי |  |
| 7:25.232 | אורה |  |
| 7:25.233 | אחת |  |
| 7:25.234 | לאחת |  |
| 7:25.235 | למצוא |  |
| 7:25.236 | חשבון | . |
| 7:25.237 | אבל |  |
| 7:25.238 | כי |  |
| 7:25.239 | חג | ' |
| 7:25.240 | ראשונות |  |
| 7:25.241 | רומזים |  |
| 7:25.242 | לראש | . |
| 7:25.243 | גדולה |  |
| 7:25.244 | גבורה |  |
| 7:25.245 | זרועות |  |
| 7:25.246 | עולם | . |
| 7:25.247 | תפארת |  |
| 7:25.248 | לשדרה |  |
| 7:25.249 | וללב | . |
| 7:25.250 | נצח |  |
| 7:25.251 | והוד |  |
| 7:25.252 | תרין |  |
| 7:25.253 | ירכין | . |
| 7:25.254 | ויסוד |  |
| 7:25.255 | דא |  |
| 7:25.256 | ברית | . |
| 7:25.257 | ומלכות |  |
| 7:25.258 | היא |  |
| 7:25.259 | מקבלת |  |
| 7:25.260 | את |  |
| 7:25.261 | הכל |  |
| 7:25.262 | ומשפעת |  |
| 7:25.263 | לכל |  |
| 7:25.264 | התחתונים | . |
| 7:25.265 | וזהו |  |
| 7:25.266 | סוד |  |
| 7:25.267 | היחוד |  |
| 7:25.268 | והקדושה | , |
| 7:25.269 | וצריך |  |
| 7:25.270 | לחבר |  |
| 7:25.271 | את |  |
| 7:25.272 | האהל |  |
| 7:25.273 | להיות |  |
| 7:25.274 | אחד |  |
| 7:25.275 | השני |  |
| 7:25.276 | שמות |  |
| 7:25.277 | ואהדונה"א | . |
| 7:25.278 | נמצאת |  |
| 7:25.279 | למד |  |
| 7:25.280 | כשתאמר |  |
| 7:25.281 | מלת |  |
| 7:25.282 | אחד | , |
| 7:25.283 | תכוין |  |
| 7:25.284 | לתשע |  |
| 7:25.285 | ספירות | , |
| 7:25.286 | אל"ף |  |
| 7:25.287 | וחי"ת | , |
| 7:25.288 | ומלכות |  |
| 7:25.289 | רומזת |  |
| 7:25.290 | לעשירית | , |
| 7:25.291 | שהיא |  |
| 7:25.292 | דל"ת | , |
| 7:25.293 | וחסוד |  |
| 7:25.294 | למען |  |
| 7:25.295 | יאריך |  |
| 7:25.296 | ימים |  |
| 7:25.297 | על |  |
| 7:25.298 | ממלכתו | , |
| 7:25.299 | על |  |
| 7:25.300 | שם |  |
| 7:25.301 | הבורא |  |
| 7:25.302 | מלאך |  |
| 7:25.303 | וגלגל |  |
| 7:25.304 | ואופן |  |
| 7:25.305 | עכ"ל | . |
| 7:25.306 | ראה |  |
| 7:25.307 | אבי |  |
| 7:25.308 | גם |  |
| 7:25.309 | ראה |  |
| 7:25.310 | איך |  |
| 7:25.311 | החכם |  |
| 7:25.312 | הנז | ' |
| 7:25.313 | מהר"י |  |
| 7:25.314 | צאהרי |  |
| 7:25.315 | הרגיש |  |
| 7:25.316 | וידע |  |
| 7:25.317 | שהם |  |
| 7:25.318 | שתי |  |
| 7:25.319 | דעות |  |
| 7:25.320 | מתנגדות |  |
| 7:25.321 | ביחוד |  |
| 7:25.322 | השי"ת | , |
| 7:25.323 | אך |  |
| 7:25.324 | ורק |  |
| 7:25.325 | כי |  |
| 7:25.326 | חשב |  |
| 7:25.327 | את |  |
| 7:25.328 | הראשונים |  |
| 7:25.329 | חכמי |  |
| 7:25.330 | התלמוד |  |
| 7:25.331 | שאמרו |  |
| 7:25.332 | להאריך |  |
| 7:25.333 | בדל"ת |  |
| 7:25.334 | דאחד |  |
| 7:25.335 | כדי |  |
| 7:25.336 | להמליכו |  |
| 7:25.337 | בארבע |  |
| 7:25.338 | פינות |  |
| 7:25.339 | העולם |  |
| 7:25.340 | לחסידי |  |
| 7:25.341 | מדע | . |
| 7:25.342 | שלא |  |
| 7:25.343 | ידעו |  |
| 7:25.344 | בקבלה |  |
| 7:25.345 | הנכריה |  |
| 7:25.346 | החדשה |  |
| 7:25.347 | ולא |  |
| 7:25.348 | הגיע |  |
| 7:25.349 | אליהם |  |
| 7:25.350 | דבר |  |
| 7:25.351 | ממנה | , |
| 7:25.352 | ולא |  |
| 7:25.353 | ידעו |  |
| 7:25.354 | ליחד |  |
| 7:25.355 | את |  |
| 7:25.356 | השם |  |
| 7:25.357 | עם |  |
| 7:25.358 | ברואיו |  |
| 7:25.359 | הספירות |  |
| 7:25.360 | המדומות |  |
| 7:25.361 | שעשו |  |
| 7:25.362 | אותם |  |
| 7:25.363 | גופים |  |
| 7:25.364 | ופרצופים |  |
| 7:25.365 | לאלהינו |  |
| 7:25.366 | ב"ה |  |
| 7:25.367 | והוא |  |
| 7:25.368 | נשמה |  |
| 7:25.369 | להם | . |
| 7:25.370 | ואדרבא |  |
| 7:25.371 | נהפוך |  |
| 7:25.372 | הוא |  |
| 7:25.373 | שאיסור |  |
| 7:25.374 | גמור |  |
| 7:25.375 | הוא |  |
| 7:25.376 | לכלול |  |
| 7:25.377 | ולשתף |  |
| 7:25.378 | עם |  |
| 7:25.379 | השי"ת |  |
| 7:25.380 | שום |  |
| 7:25.381 | נברא |  |
| 7:25.382 | כמו |  |
| 7:25.383 | שנתבאר |  |
| 7:25.384 | לעיל |  |
| 7:25.385 | סי | ' |
| 7:25.386 | ע"ד | . |
| 7:25.387 | וסי | ' |
| 7:25.388 | פ"ב |  |
| 7:25.389 | בשם |  |
| 7:25.390 | מדרש |  |
| 7:25.391 | רבה | , |
| 7:25.392 | והטעם |  |
| 7:25.393 | כי |  |
| 7:25.394 | הוא |  |
| 7:25.395 | ית | ' |
| 7:25.396 | נעלה |  |
| 7:25.397 | ומרומם |  |
| 7:25.398 | ושגיא |  |
| 7:25.399 | כה | , |
| 7:25.400 | ודי |  |
| 7:25.401 | לו |  |
| 7:25.402 | בעצמותו |  |
| 7:25.403 | ואין |  |
| 7:25.404 | צריך |  |
| 7:25.405 | לבריותיו |  |
| 7:25.406 | להשתתף |  |
| 7:25.407 | עם |  |
| 7:25.408 | שום |  |
| 7:25.409 | אחד |  |
| 7:25.410 | מהם | . |
| 7:25.411 | שכל |  |
| 7:25.412 | בריותיו |  |
| 7:25.413 | צריכים |  |
| 7:25.414 | לו | , |
| 7:25.415 | והוא |  |
| 7:25.416 | אינו |  |
| 7:25.417 | צריך |  |
| 7:25.418 | להם |  |
| 7:25.419 | להעזר |  |
| 7:25.420 | בהם |  |
| 7:25.421 | ח"ו |  |
| 7:25.422 | כי |  |
| 7:25.423 | הוא |  |
| 7:25.424 | אמר |  |
| 7:25.425 | ויהי |  |
| 7:25.426 | הוא |  |
| 7:25.427 | לבדו |  |
| 7:25.428 | צוה |  |
| 7:25.429 | ויעמוד | . |
| 7:25.430 | ולא |  |
| 7:25.431 | ירד |  |
| 7:25.432 | הרב |  |
| 7:25.433 | הנז | ' |
| 7:25.434 | לעומק |  |
| 7:25.435 | דברי |  |
| 7:25.436 | ה |  |
| 7:25.437 | רס"ג |  |
| 7:25.438 | והרמב"ם |  |
| 7:25.439 | שהוא |  |
| 7:25.440 | עסוק |  |
| 7:25.441 | בהם |  |
| 7:25.442 | להעתיק |  |
| 7:25.443 | דבריהם |  |
| 7:25.444 | והוא |  |
| 7:25.445 | מחבבן | , |
| 7:25.446 | ועשה |  |
| 7:25.447 | אותן |  |
| 7:25.448 | כעין |  |
| 7:25.449 | הצעה |  |
| 7:25.450 | והקדמה |  |
| 7:25.451 | לאמתת |  |
| 7:25.452 | יחוד |  |
| 7:25.453 | השי"ת | , |
| 7:25.454 | והרגיש |  |
| 7:25.455 | שהם |  |
| 7:25.456 | שתי |  |
| 7:25.457 | סברות | . |
| 7:25.458 | ולא |  |
| 7:25.459 | הרגיש |  |
| 7:25.460 | הסתירה |  |
| 7:25.461 | בפרק |  |
| 7:25.462 | ג | ' |
| 7:25.463 | לראשון |  |
| 7:25.464 | שהעתיק |  |
| 7:25.465 | מקצת |  |
| 7:25.466 | דבריו | , |
| 7:25.467 | שהם |  |
| 7:25.468 | מתנגדים |  |
| 7:25.469 | לגמרי |  |
| 7:25.470 | לשטת |  |
| 7:25.471 | הקבלה |  |
| 7:25.472 | הנכריה |  |
| 7:25.473 | החדשה | , |
| 7:25.474 | רק |  |
| 7:25.475 | חשב |  |
| 7:25.476 | אותם |  |
| 7:25.477 | לחסידי |  |
| 7:25.478 | מדע | . |
| 7:25.479 | שלא |  |
| 7:25.480 | ידעו |  |
| 7:25.481 | לצרף |  |
| 7:25.482 | ולשתף |  |
| 7:25.483 | ולכלול |  |
| 7:25.484 | את |  |
| 7:25.485 | הצורות |  |
| 7:25.486 | והפרצופים |  |
| 7:25.487 | עם |  |
| 7:25.488 | אמתתו |  |
| 7:25.489 | ית | ' |
| 7:25.490 | וליחדו |  |
| 7:25.491 | אמהם |  |
| 7:25.492 | דוקא | . |
| 7:25.493 | ולא |  |
| 7:25.494 | ידעו |  |
| 7:25.495 | ולא |  |
| 7:25.496 | יבינו |  |
| 7:25.497 | כי |  |
| 7:25.498 | הרס"ג |  |
| 7:25.499 | והרמב"ם |  |
| 7:25.500 | ורבי |  |
| 7:25.501 | יהודה |  |
| 7:25.502 | הלוי |  |
| 7:25.503 | ובעל |  |
| 7:25.504 | חובת |  |
| 7:25.505 | הלבבות |  |
| 7:25.506 | נלחמים |  |
| 7:25.507 | המה |  |
| 7:25.508 | בחזקת |  |
| 7:25.509 | היד |  |
| 7:25.510 | עם |  |
| 7:25.511 | דעות |  |
| 7:25.512 | הנכריות |  |
| 7:25.513 | הזרות |  |
| 7:25.514 | שהתחילו |  |
| 7:25.515 | להנץ |  |
| 7:25.516 | בין |  |
| 7:25.517 | היהודים |  |
| 7:25.518 | מזמן |  |
| 7:25.519 | הגאונים | , |
| 7:25.520 | ועדיין |  |
| 7:25.521 | היו |  |
| 7:25.522 | בעל |  |
| 7:25.523 | פה |  |
| 7:25.524 | ולא |  |
| 7:25.525 | הוחקו |  |
| 7:25.526 | על |  |
| 7:25.527 | ספר | . |
| 7:25.528 | עד |  |
| 7:25.529 | שיצא |  |
| 7:25.530 | הספר |  |
| 7:25.531 | הנכרי |  |
| 7:25.532 | הזהר |  |
| 7:25.533 | המבזה |  |
| 7:25.534 | את |  |
| 7:25.535 | תורה |  |
| 7:25.536 | שבעל |  |
| 7:25.537 | פה |  |
| 7:25.538 | המשנה |  |
| 7:25.539 | והתלמוד | , |
| 7:25.540 | וקורא |  |
| 7:25.541 | להם |  |
| 7:25.542 | אבן |  |
| 7:25.543 | משכית | , |
| 7:25.544 | ועץ |  |
| 7:25.545 | הדעת |  |
| 7:25.546 | טוב |  |
| 7:25.547 | ורע | . |
| 7:25.548 | וקליפה | , |
| 7:25.549 | ושפחה |  |
| 7:25.550 | כי |  |
| 7:25.551 | תירש |  |
| 7:25.552 | גברתה | , |
| 7:25.553 | נגד |  |
| 7:25.554 | דברי |  |
| 7:25.555 | דוד |  |
| 7:25.556 | המלך |  |
| 7:25.557 | ע"ה |  |
| 7:25.558 | שאמר |  |
| 7:25.559 | עליהם |  |
| 7:25.560 | הנחמדים |  |
| 7:25.561 | מזהב |  |
| 7:25.562 | ומפז |  |
| 7:25.563 | רב |  |
| 7:25.564 | ומתוקים |  |
| 7:25.565 | מדבש |  |
| 7:25.566 | ומנופת |  |
| 7:25.567 | צופים | , |
| 7:25.568 | והסית |  |
| 7:25.569 | והדיח |  |
| 7:25.570 | את |  |
| 7:25.571 | הרבים |  |
| 7:25.572 | לעבוד |  |
| 7:25.573 | לאל |  |
| 7:25.574 | אחר |  |
| 7:25.575 | קצר |  |
| 7:25.576 | אפים |  |
| 7:25.577 | נברא | , |
| 7:25.578 | כמו |  |
| 7:25.579 | שכתב |  |
| 7:25.580 | הר | ' |
| 7:25.581 | יושר |  |
| 7:25.582 | לבב |  |
| 7:25.583 | לעבדו |  |
| 7:25.584 | אע"פ |  |
| 7:25.585 | שהוא |  |
| 7:25.586 | נברא | , |
| 7:25.587 | לפי |  |
| 7:25.588 | שהאמינו |  |
| 7:25.589 | באמונה |  |
| 7:25.590 | שלימה |  |
| 7:25.591 | כי |  |
| 7:25.592 | מרשב"י |  |
| 7:25.593 | יצאו |  |
| 7:25.594 | הדברים |  |
| 7:25.595 | ההם | . |
| 7:25.596 | והסבה |  |
| 7:25.597 | שגרמה |  |
| 7:25.598 | להם |  |
| 7:25.599 | זה |  |
| 7:25.600 | שלא |  |
| 7:25.601 | שמו |  |
| 7:25.602 | אל |  |
| 7:25.603 | לבם |  |
| 7:25.604 | ונגד |  |
| 7:25.605 | עיניהם |  |
| 7:25.606 | דברי |  |
| 7:25.607 | הראשונים |  |
| 7:25.608 | בעלי |  |
| 7:25.609 | הכללים |  |
| 7:25.610 | המעמיקים |  |
| 7:25.611 | בתלמוד | , |
| 7:25.612 | המה |  |
| 7:25.613 | דברי |  |
| 7:25.614 | הרא"ש |  |
| 7:25.615 | ז"ל |  |
| 7:25.616 | שהבאתי |  |
| 7:25.617 | לעיל |  |
| 7:25.618 | סי | ' |
| 7:25.619 | מ"ב |  |
| 7:25.620 | שכל |  |
| 7:25.621 | תוספתא |  |
| 7:25.622 | או |  |
| 7:25.623 | בריתא |  |
| 7:25.624 | שלא |  |
| 7:25.625 | נתפשטה |  |
| 7:25.626 | עד |  |
| 7:25.627 | אחר |  |
| 7:25.628 | סידור |  |
| 7:25.629 | הגמרא |  |
| 7:25.630 | לאו |  |
| 7:25.631 | דסמכא |  |
| 7:25.632 | היא | . |
| 7:25.633 | דמסתמא |  |
| 7:25.634 | מחברי |  |
| 7:25.635 | התלמוד |  |
| 7:25.636 | חקרו |  |
| 7:25.637 | ודרשו |  |
| 7:25.638 | לידע |  |
| 7:25.639 | כל |  |
| 7:25.640 | הספרים |  |
| 7:25.641 | שנכתבו |  |
| 7:25.642 | על |  |
| 7:25.643 | דעת |  |
| 7:25.644 | חכמים |  |
| 7:25.645 | וביררו |  |
| 7:25.646 | אותם |  |
| 7:25.647 | שהם |  |
| 7:25.648 | בר |  |
| 7:25.649 | סמכא | . |
| 7:25.650 | וכל |  |
| 7:25.651 | תוספתא |  |
| 7:25.652 | או |  |
| 7:25.653 | ברייתא |  |
| 7:25.654 | שהיא |  |
| 7:25.655 | היפך |  |
| 7:25.656 | הגמרא |  |
| 7:25.657 | לאו |  |
| 7:25.658 | דסמכא |  |
| 7:25.659 | היא |  |
| 7:25.660 | דודאי |  |
| 7:25.661 | משבשתא |  |
| 7:25.662 | היא |  |
| 7:25.663 | או |  |
| 7:25.664 | מזויפת |  |
| 7:25.665 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

