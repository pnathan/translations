# Translation pack: milhamot-hashem perek-9-11 (Milchamot ha-Shem 9:13–9:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-11`.

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
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- האלהי before a name: "the divine", with a note; never dropped.
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


## Milchamot ha-Shem 9:13

Hebrew (Wikisource): קלא) ולפי שנשתבשה אצלם אמונת יחוד השי"ת שהיא לפי האמת אחידות שאין כמוה בכל שאר האחדים כמו שאמרו רז"ל במדרש רבה (הובא לעיל סי' פ"ב) לפי' העלימו עין ממנה והיה להם לקבוע דיני יחוד השם בש"ע ונושאי כליו ובספרי האחרונים בסימנים וסעיפים מיוחדים ולהאריך בהם יותר ממה שהאריך בהם רבי' הרמב"ם בהל' יסוה"ת ובמורה וכמה דינים היו מסתעפים ומתחדשים באמונת היחוד לפי שאין אמונת היחוד איך יאמר אלא יאמין ויקבל עליו עול מלכות שמים באמונת אחדות אחת שאין כמוה בשאר כל האחדים שבעולם בלי שום שתוף עלה אחרת מכל הפרצופים הנבראים וגם להאמין אמונה שלימה במשנה ובתלמוד המקובלים בידינו שמשה רבי' קבלם מפי הגבורה ולא ממטטרון ושאין בהם כל רע כלל לא כדברי הזהר שאמר שהם מעורבים מטוב ורע ושנקראו אבן משכית וקליפה ומוץ ותבן וסלע אחרא וכיוצא בזלזולים אלו אלא הכל מפי הגבורה טוב שאין בו רע כלל והוא חיינו ואורך ימינו. הנה זו רעה חולה עליה יהיו כל הדוים איך נעזוב את ה' שהוא אלהים חיים ומלך עולם ונשמע לדברי הספר המוצא מתוך קברים על ידי מלך ממלכי העכו"ם לעבוד אלהים חדשים פרצופים נבראים בשתוף הדבר האלהי המעורב בהם לפי דעתם שהוא נשמתם וכל מצותינו שנעשה לא לשם השי"ת שאמר ונעשה רצונו רק לקשט בהם עם ברתא שהם זעיר ומלכות, וידוע הוא ומפורש בכתוב שעובדי האלילים הקדומים היו מקשטים אליליהם בכסף וזהב ואבנים יקרות, ואלה המקובלים החדשים מקשטים את אליליהם בתפלותיהם ומעשה מצותיהם והם הנותנים להם חיות ומזון בהתפללם אליהם להמשיך להם שפע וחיות מהאלהים העליון בכדי שיוכלו להשפיע לנו משלהם וחושבים שהתפלה והעבודה היא צורך גבוה כמו שכתב הר"מ גבאי דהיינו צורך הפרצופים הנעבדים, ולא צרכי גופינו היפך דברי רשב"י בירושלמי כדגר' ההם פ"ק דשבת מפסיקין לקרית שמע ואין מפסיקין לתפלה. אמר רבי אחא קרית דבר תורה, ותפלה אינה דבר תורה. אמר ר' בא קרית שמע זמנה קבוע ותפלה אין זמנה קבוע, אמר ר' יוסה קרית שמע אינה צריכה כונה, ותפלה צריכה כונה, אמר ר' מנא קשיתה קומי דר' יוסה ואפי' תימא ק"ש אינה צריכה כונה, שלשה פיסוקים הראשונים צריכים כונה מן גו דאינון ציבחן הוא מכוין ר' יוחנן בשם רשב"י כגון אנו שעסוקין בתלמוד תורה אפי לק"ש אין אנו מפסיקין. רבי יוחנן אמר על גרמיה כגון אנו שאין אנו עסוקין בתלמוד תורה אפי' לתפלה אנו מפסיקין. דין כדעתיה ודין כדעתיה, רבי יוחנן כדעתיה דר' יוחנן אמר הלואי מתפללין כל היום למה שאין תפלה מפסדת, רשב"י כדעתיה, דרשב"י אמר אלו הוינא קאים על טורא דסיני בשעתא דאתיהיבת אוריתא לישראל. הוינא מתבע קומי רחמנא דאתברי להדין לבר נשא תרין פומין. חד דיהוי לעי באוריתא. וחד דיתעביד ביה כל תורכוי, חזר ואמר, מה אין חד הוא, לית עלמא יכיל קאים ביה מן דלטוריא דיליה. אלו הויין תרין על אחת כמה וכמה. אמר ר' יוסא קומי ר' ירמיה, אתיא דר' יוחנן כר' חנניה בן עקביה דתני כותבי ספרים תפלי ומזוזות מפסיקין לק"ש ואין מפסיקין לתפלה. ר' חנניה בו עקביה אומר כשם שמפסיקין לק"ש כך מפסיקין לתפלה ולתפלין ולשאר כל מצותיה של תורה. ופריך ולא מודי רשב"י שמפסיקין לעשות סוכה ולעשות לולב, ולית ליה לרשב"י הלמד לעשות. לא הלמד שלא לעשות? שהלמד שלא לעשות נוח לו אלו לא נברא. ואמר ר' יוחנן הלמד שלא לעשות נוח לו אלו נתהפכה שליתו על פניו ולא יצא לעולם? ומשני טעמיה דרשב"י זה שינון וזה שינון. ואין מפסיקין שינון מפני שינון. ופריך והתנינן הקורא מכאן ואילך לא הפסיד כאשדם שהוא קורא בתורה, הא בעונתה חביבה היא יותר מדברי תורה? ומשני אמר ר' יודן רשב"י על ידי שהיה חדיר בדברי תורה, לפיכך אינה חביבה עליו יותר מדברי תורה ע"כ. שמע מינה מדקאמר רשב"י כדעתיה וכו' שהתנא החסיד רשב"י ס"ל שהתפלה היא תביעת צרכי הגוף ואין לפסוק מדברי תורה שהיא עולם כדי להתפלל ולתבוע צרכי הגוף ואין לפסוק מדברי תורה שהיא עולם כדי להתפלל ותבוע צרכיו שהם חיי שעה ולכן אמר אלו הוינא קאים וכו' הוינא מתבע תרי פומי. חד לעסוק בתורה, וחד לתפלה ולשאר עסקי הגוף. כי התפלה גם היא מעסק הגוף. שמבקש בה רפואה ופרנסה ושאר צרכים. וכדגר' בפ' יציאות השבת. רבא אשכחיה לרב המנונא דקא מאריך בצלותא, אמר מניחין חי עולם ועוסקין בחיי שעה, והוא סבר זמן תורה לחוד, וזמן תפלה לחוד. עוד שם רבי ירמיה הזה יתיב קמי ר' זירא וקא גריס. מטא זמן תפלה וקא מסרהיב קרי עליה רבי זירא מסיר אזנו משמוע צורה גם תפלתו תועבה ע"כ. אלמלא דעסק תורה היה חביב עליהם יותר מן התפלה. כי הוא חיינו ואורך ימינו לעולם שכולו ארוך לעולם שכולו טוב. אלא דבתר הכי מסקינן דלא שנו דאין מפסיקין לתפלה אלא כגון רשב"י וחביריו שתורתן אומנותן. אבל כגון אנו מפסיקין בין לק"ש בין לתפלה, וכן פסקו הרי"ף והרמב"ם כר' יוחנן, והוא כדאמרינן בעלמא הרבה עשו כרשב"י ולא עלתה בידם.

| id | word | punctuation after |
|---|---|---|
| 9:13.1 | קלא | ) |
| 9:13.2 | ולפי |  |
| 9:13.3 | שנשתבשה |  |
| 9:13.4 | אצלם |  |
| 9:13.5 | אמונת |  |
| 9:13.6 | יחוד |  |
| 9:13.7 | השי"ת |  |
| 9:13.8 | שהיא |  |
| 9:13.9 | לפי |  |
| 9:13.10 | האמת |  |
| 9:13.11 | אחידות |  |
| 9:13.12 | שאין |  |
| 9:13.13 | כמוה |  |
| 9:13.14 | בכל |  |
| 9:13.15 | שאר |  |
| 9:13.16 | האחדים |  |
| 9:13.17 | כמו |  |
| 9:13.18 | שאמרו |  |
| 9:13.19 | רז"ל |  |
| 9:13.20 | במדרש |  |
| 9:13.21 | רבה |  |
| 9:13.22 | הובא |  |
| 9:13.23 | לעיל |  |
| 9:13.24 | סי | ' |
| 9:13.25 | פ"ב | ) |
| 9:13.26 | לפי | ' |
| 9:13.27 | העלימו |  |
| 9:13.28 | עין |  |
| 9:13.29 | ממנה |  |
| 9:13.30 | והיה |  |
| 9:13.31 | להם |  |
| 9:13.32 | לקבוע |  |
| 9:13.33 | דיני |  |
| 9:13.34 | יחוד |  |
| 9:13.35 | השם |  |
| 9:13.36 | בש"ע |  |
| 9:13.37 | ונושאי |  |
| 9:13.38 | כליו |  |
| 9:13.39 | ובספרי |  |
| 9:13.40 | האחרונים |  |
| 9:13.41 | בסימנים |  |
| 9:13.42 | וסעיפים |  |
| 9:13.43 | מיוחדים |  |
| 9:13.44 | ולהאריך |  |
| 9:13.45 | בהם |  |
| 9:13.46 | יותר |  |
| 9:13.47 | ממה |  |
| 9:13.48 | שהאריך |  |
| 9:13.49 | בהם |  |
| 9:13.50 | רבי | ' |
| 9:13.51 | הרמב"ם |  |
| 9:13.52 | בהל | ' |
| 9:13.53 | יסוה"ת |  |
| 9:13.54 | ובמורה |  |
| 9:13.55 | וכמה |  |
| 9:13.56 | דינים |  |
| 9:13.57 | היו |  |
| 9:13.58 | מסתעפים |  |
| 9:13.59 | ומתחדשים |  |
| 9:13.60 | באמונת |  |
| 9:13.61 | היחוד |  |
| 9:13.62 | לפי |  |
| 9:13.63 | שאין |  |
| 9:13.64 | אמונת |  |
| 9:13.65 | היחוד |  |
| 9:13.66 | איך |  |
| 9:13.67 | יאמר |  |
| 9:13.68 | אלא |  |
| 9:13.69 | יאמין |  |
| 9:13.70 | ויקבל |  |
| 9:13.71 | עליו |  |
| 9:13.72 | עול |  |
| 9:13.73 | מלכות |  |
| 9:13.74 | שמים |  |
| 9:13.75 | באמונת |  |
| 9:13.76 | אחדות |  |
| 9:13.77 | אחת |  |
| 9:13.78 | שאין |  |
| 9:13.79 | כמוה |  |
| 9:13.80 | בשאר |  |
| 9:13.81 | כל |  |
| 9:13.82 | האחדים |  |
| 9:13.83 | שבעולם |  |
| 9:13.84 | בלי |  |
| 9:13.85 | שום |  |
| 9:13.86 | שתוף |  |
| 9:13.87 | עלה |  |
| 9:13.88 | אחרת |  |
| 9:13.89 | מכל |  |
| 9:13.90 | הפרצופים |  |
| 9:13.91 | הנבראים |  |
| 9:13.92 | וגם |  |
| 9:13.93 | להאמין |  |
| 9:13.94 | אמונה |  |
| 9:13.95 | שלימה |  |
| 9:13.96 | במשנה |  |
| 9:13.97 | ובתלמוד |  |
| 9:13.98 | המקובלים |  |
| 9:13.99 | בידינו |  |
| 9:13.100 | שמשה |  |
| 9:13.101 | רבי | ' |
| 9:13.102 | קבלם |  |
| 9:13.103 | מפי |  |
| 9:13.104 | הגבורה |  |
| 9:13.105 | ולא |  |
| 9:13.106 | ממטטרון |  |
| 9:13.107 | ושאין |  |
| 9:13.108 | בהם |  |
| 9:13.109 | כל |  |
| 9:13.110 | רע |  |
| 9:13.111 | כלל |  |
| 9:13.112 | לא |  |
| 9:13.113 | כדברי |  |
| 9:13.114 | הזהר |  |
| 9:13.115 | שאמר |  |
| 9:13.116 | שהם |  |
| 9:13.117 | מעורבים |  |
| 9:13.118 | מטוב |  |
| 9:13.119 | ורע |  |
| 9:13.120 | ושנקראו |  |
| 9:13.121 | אבן |  |
| 9:13.122 | משכית |  |
| 9:13.123 | וקליפה |  |
| 9:13.124 | ומוץ |  |
| 9:13.125 | ותבן |  |
| 9:13.126 | וסלע |  |
| 9:13.127 | אחרא |  |
| 9:13.128 | וכיוצא |  |
| 9:13.129 | בזלזולים |  |
| 9:13.130 | אלו |  |
| 9:13.131 | אלא |  |
| 9:13.132 | הכל |  |
| 9:13.133 | מפי |  |
| 9:13.134 | הגבורה |  |
| 9:13.135 | טוב |  |
| 9:13.136 | שאין |  |
| 9:13.137 | בו |  |
| 9:13.138 | רע |  |
| 9:13.139 | כלל |  |
| 9:13.140 | והוא |  |
| 9:13.141 | חיינו |  |
| 9:13.142 | ואורך |  |
| 9:13.143 | ימינו | . |
| 9:13.144 | הנה |  |
| 9:13.145 | זו |  |
| 9:13.146 | רעה |  |
| 9:13.147 | חולה |  |
| 9:13.148 | עליה |  |
| 9:13.149 | יהיו |  |
| 9:13.150 | כל |  |
| 9:13.151 | הדוים |  |
| 9:13.152 | איך |  |
| 9:13.153 | נעזוב |  |
| 9:13.154 | את |  |
| 9:13.155 | ה | ' |
| 9:13.156 | שהוא |  |
| 9:13.157 | אלהים |  |
| 9:13.158 | חיים |  |
| 9:13.159 | ומלך |  |
| 9:13.160 | עולם |  |
| 9:13.161 | ונשמע |  |
| 9:13.162 | לדברי |  |
| 9:13.163 | הספר |  |
| 9:13.164 | המוצא |  |
| 9:13.165 | מתוך |  |
| 9:13.166 | קברים |  |
| 9:13.167 | על |  |
| 9:13.168 | ידי |  |
| 9:13.169 | מלך |  |
| 9:13.170 | ממלכי |  |
| 9:13.171 | העכו"ם |  |
| 9:13.172 | לעבוד |  |
| 9:13.173 | אלהים |  |
| 9:13.174 | חדשים |  |
| 9:13.175 | פרצופים |  |
| 9:13.176 | נבראים |  |
| 9:13.177 | בשתוף |  |
| 9:13.178 | הדבר |  |
| 9:13.179 | האלהי |  |
| 9:13.180 | המעורב |  |
| 9:13.181 | בהם |  |
| 9:13.182 | לפי |  |
| 9:13.183 | דעתם |  |
| 9:13.184 | שהוא |  |
| 9:13.185 | נשמתם |  |
| 9:13.186 | וכל |  |
| 9:13.187 | מצותינו |  |
| 9:13.188 | שנעשה |  |
| 9:13.189 | לא |  |
| 9:13.190 | לשם |  |
| 9:13.191 | השי"ת |  |
| 9:13.192 | שאמר |  |
| 9:13.193 | ונעשה |  |
| 9:13.194 | רצונו |  |
| 9:13.195 | רק |  |
| 9:13.196 | לקשט |  |
| 9:13.197 | בהם |  |
| 9:13.198 | עם |  |
| 9:13.199 | ברתא |  |
| 9:13.200 | שהם |  |
| 9:13.201 | זעיר |  |
| 9:13.202 | ומלכות | , |
| 9:13.203 | וידוע |  |
| 9:13.204 | הוא |  |
| 9:13.205 | ומפורש |  |
| 9:13.206 | בכתוב |  |
| 9:13.207 | שעובדי |  |
| 9:13.208 | האלילים |  |
| 9:13.209 | הקדומים |  |
| 9:13.210 | היו |  |
| 9:13.211 | מקשטים |  |
| 9:13.212 | אליליהם |  |
| 9:13.213 | בכסף |  |
| 9:13.214 | וזהב |  |
| 9:13.215 | ואבנים |  |
| 9:13.216 | יקרות | , |
| 9:13.217 | ואלה |  |
| 9:13.218 | המקובלים |  |
| 9:13.219 | החדשים |  |
| 9:13.220 | מקשטים |  |
| 9:13.221 | את |  |
| 9:13.222 | אליליהם |  |
| 9:13.223 | בתפלותיהם |  |
| 9:13.224 | ומעשה |  |
| 9:13.225 | מצותיהם |  |
| 9:13.226 | והם |  |
| 9:13.227 | הנותנים |  |
| 9:13.228 | להם |  |
| 9:13.229 | חיות |  |
| 9:13.230 | ומזון |  |
| 9:13.231 | בהתפללם |  |
| 9:13.232 | אליהם |  |
| 9:13.233 | להמשיך |  |
| 9:13.234 | להם |  |
| 9:13.235 | שפע |  |
| 9:13.236 | וחיות |  |
| 9:13.237 | מהאלהים |  |
| 9:13.238 | העליון |  |
| 9:13.239 | בכדי |  |
| 9:13.240 | שיוכלו |  |
| 9:13.241 | להשפיע |  |
| 9:13.242 | לנו |  |
| 9:13.243 | משלהם |  |
| 9:13.244 | וחושבים |  |
| 9:13.245 | שהתפלה |  |
| 9:13.246 | והעבודה |  |
| 9:13.247 | היא |  |
| 9:13.248 | צורך |  |
| 9:13.249 | גבוה |  |
| 9:13.250 | כמו |  |
| 9:13.251 | שכתב |  |
| 9:13.252 | הר"מ |  |
| 9:13.253 | גבאי |  |
| 9:13.254 | דהיינו |  |
| 9:13.255 | צורך |  |
| 9:13.256 | הפרצופים |  |
| 9:13.257 | הנעבדים | , |
| 9:13.258 | ולא |  |
| 9:13.259 | צרכי |  |
| 9:13.260 | גופינו |  |
| 9:13.261 | היפך |  |
| 9:13.262 | דברי |  |
| 9:13.263 | רשב"י |  |
| 9:13.264 | בירושלמי |  |
| 9:13.265 | כדגר | ' |
| 9:13.266 | ההם |  |
| 9:13.267 | פ"ק |  |
| 9:13.268 | דשבת |  |
| 9:13.269 | מפסיקין |  |
| 9:13.270 | לקרית |  |
| 9:13.271 | שמע |  |
| 9:13.272 | ואין |  |
| 9:13.273 | מפסיקין |  |
| 9:13.274 | לתפלה | . |
| 9:13.275 | אמר |  |
| 9:13.276 | רבי |  |
| 9:13.277 | אחא |  |
| 9:13.278 | קרית |  |
| 9:13.279 | דבר |  |
| 9:13.280 | תורה | , |
| 9:13.281 | ותפלה |  |
| 9:13.282 | אינה |  |
| 9:13.283 | דבר |  |
| 9:13.284 | תורה | . |
| 9:13.285 | אמר |  |
| 9:13.286 | ר | ' |
| 9:13.287 | בא |  |
| 9:13.288 | קרית |  |
| 9:13.289 | שמע |  |
| 9:13.290 | זמנה |  |
| 9:13.291 | קבוע |  |
| 9:13.292 | ותפלה |  |
| 9:13.293 | אין |  |
| 9:13.294 | זמנה |  |
| 9:13.295 | קבוע | , |
| 9:13.296 | אמר |  |
| 9:13.297 | ר | ' |
| 9:13.298 | יוסה |  |
| 9:13.299 | קרית |  |
| 9:13.300 | שמע |  |
| 9:13.301 | אינה |  |
| 9:13.302 | צריכה |  |
| 9:13.303 | כונה | , |
| 9:13.304 | ותפלה |  |
| 9:13.305 | צריכה |  |
| 9:13.306 | כונה | , |
| 9:13.307 | אמר |  |
| 9:13.308 | ר | ' |
| 9:13.309 | מנא |  |
| 9:13.310 | קשיתה |  |
| 9:13.311 | קומי |  |
| 9:13.312 | דר | ' |
| 9:13.313 | יוסה |  |
| 9:13.314 | ואפי | ' |
| 9:13.315 | תימא |  |
| 9:13.316 | ק"ש |  |
| 9:13.317 | אינה |  |
| 9:13.318 | צריכה |  |
| 9:13.319 | כונה | , |
| 9:13.320 | שלשה |  |
| 9:13.321 | פיסוקים |  |
| 9:13.322 | הראשונים |  |
| 9:13.323 | צריכים |  |
| 9:13.324 | כונה |  |
| 9:13.325 | מן |  |
| 9:13.326 | גו |  |
| 9:13.327 | דאינון |  |
| 9:13.328 | ציבחן |  |
| 9:13.329 | הוא |  |
| 9:13.330 | מכוין |  |
| 9:13.331 | ר | ' |
| 9:13.332 | יוחנן |  |
| 9:13.333 | בשם |  |
| 9:13.334 | רשב"י |  |
| 9:13.335 | כגון |  |
| 9:13.336 | אנו |  |
| 9:13.337 | שעסוקין |  |
| 9:13.338 | בתלמוד |  |
| 9:13.339 | תורה |  |
| 9:13.340 | אפי |  |
| 9:13.341 | לק"ש |  |
| 9:13.342 | אין |  |
| 9:13.343 | אנו |  |
| 9:13.344 | מפסיקין | . |
| 9:13.345 | רבי |  |
| 9:13.346 | יוחנן |  |
| 9:13.347 | אמר |  |
| 9:13.348 | על |  |
| 9:13.349 | גרמיה |  |
| 9:13.350 | כגון |  |
| 9:13.351 | אנו |  |
| 9:13.352 | שאין |  |
| 9:13.353 | אנו |  |
| 9:13.354 | עסוקין |  |
| 9:13.355 | בתלמוד |  |
| 9:13.356 | תורה |  |
| 9:13.357 | אפי | ' |
| 9:13.358 | לתפלה |  |
| 9:13.359 | אנו |  |
| 9:13.360 | מפסיקין | . |
| 9:13.361 | דין |  |
| 9:13.362 | כדעתיה |  |
| 9:13.363 | ודין |  |
| 9:13.364 | כדעתיה | , |
| 9:13.365 | רבי |  |
| 9:13.366 | יוחנן |  |
| 9:13.367 | כדעתיה |  |
| 9:13.368 | דר | ' |
| 9:13.369 | יוחנן |  |
| 9:13.370 | אמר |  |
| 9:13.371 | הלואי |  |
| 9:13.372 | מתפללין |  |
| 9:13.373 | כל |  |
| 9:13.374 | היום |  |
| 9:13.375 | למה |  |
| 9:13.376 | שאין |  |
| 9:13.377 | תפלה |  |
| 9:13.378 | מפסדת | , |
| 9:13.379 | רשב"י |  |
| 9:13.380 | כדעתיה | , |
| 9:13.381 | דרשב"י |  |
| 9:13.382 | אמר |  |
| 9:13.383 | אלו |  |
| 9:13.384 | הוינא |  |
| 9:13.385 | קאים |  |
| 9:13.386 | על |  |
| 9:13.387 | טורא |  |
| 9:13.388 | דסיני |  |
| 9:13.389 | בשעתא |  |
| 9:13.390 | דאתיהיבת |  |
| 9:13.391 | אוריתא |  |
| 9:13.392 | לישראל | . |
| 9:13.393 | הוינא |  |
| 9:13.394 | מתבע |  |
| 9:13.395 | קומי |  |
| 9:13.396 | רחמנא |  |
| 9:13.397 | דאתברי |  |
| 9:13.398 | להדין |  |
| 9:13.399 | לבר |  |
| 9:13.400 | נשא |  |
| 9:13.401 | תרין |  |
| 9:13.402 | פומין | . |
| 9:13.403 | חד |  |
| 9:13.404 | דיהוי |  |
| 9:13.405 | לעי |  |
| 9:13.406 | באוריתא | . |
| 9:13.407 | וחד |  |
| 9:13.408 | דיתעביד |  |
| 9:13.409 | ביה |  |
| 9:13.410 | כל |  |
| 9:13.411 | תורכוי | , |
| 9:13.412 | חזר |  |
| 9:13.413 | ואמר | , |
| 9:13.414 | מה |  |
| 9:13.415 | אין |  |
| 9:13.416 | חד |  |
| 9:13.417 | הוא | , |
| 9:13.418 | לית |  |
| 9:13.419 | עלמא |  |
| 9:13.420 | יכיל |  |
| 9:13.421 | קאים |  |
| 9:13.422 | ביה |  |
| 9:13.423 | מן |  |
| 9:13.424 | דלטוריא |  |
| 9:13.425 | דיליה | . |
| 9:13.426 | אלו |  |
| 9:13.427 | הויין |  |
| 9:13.428 | תרין |  |
| 9:13.429 | על |  |
| 9:13.430 | אחת |  |
| 9:13.431 | כמה |  |
| 9:13.432 | וכמה | . |
| 9:13.433 | אמר |  |
| 9:13.434 | ר | ' |
| 9:13.435 | יוסא |  |
| 9:13.436 | קומי |  |
| 9:13.437 | ר | ' |
| 9:13.438 | ירמיה | , |
| 9:13.439 | אתיא |  |
| 9:13.440 | דר | ' |
| 9:13.441 | יוחנן |  |
| 9:13.442 | כר | ' |
| 9:13.443 | חנניה |  |
| 9:13.444 | בן |  |
| 9:13.445 | עקביה |  |
| 9:13.446 | דתני |  |
| 9:13.447 | כותבי |  |
| 9:13.448 | ספרים |  |
| 9:13.449 | תפלי |  |
| 9:13.450 | ומזוזות |  |
| 9:13.451 | מפסיקין |  |
| 9:13.452 | לק"ש |  |
| 9:13.453 | ואין |  |
| 9:13.454 | מפסיקין |  |
| 9:13.455 | לתפלה | . |
| 9:13.456 | ר | ' |
| 9:13.457 | חנניה |  |
| 9:13.458 | בו |  |
| 9:13.459 | עקביה |  |
| 9:13.460 | אומר |  |
| 9:13.461 | כשם |  |
| 9:13.462 | שמפסיקין |  |
| 9:13.463 | לק"ש |  |
| 9:13.464 | כך |  |
| 9:13.465 | מפסיקין |  |
| 9:13.466 | לתפלה |  |
| 9:13.467 | ולתפלין |  |
| 9:13.468 | ולשאר |  |
| 9:13.469 | כל |  |
| 9:13.470 | מצותיה |  |
| 9:13.471 | של |  |
| 9:13.472 | תורה | . |
| 9:13.473 | ופריך |  |
| 9:13.474 | ולא |  |
| 9:13.475 | מודי |  |
| 9:13.476 | רשב"י |  |
| 9:13.477 | שמפסיקין |  |
| 9:13.478 | לעשות |  |
| 9:13.479 | סוכה |  |
| 9:13.480 | ולעשות |  |
| 9:13.481 | לולב | , |
| 9:13.482 | ולית |  |
| 9:13.483 | ליה |  |
| 9:13.484 | לרשב"י |  |
| 9:13.485 | הלמד |  |
| 9:13.486 | לעשות | . |
| 9:13.487 | לא |  |
| 9:13.488 | הלמד |  |
| 9:13.489 | שלא |  |
| 9:13.490 | לעשות | ? |
| 9:13.491 | שהלמד |  |
| 9:13.492 | שלא |  |
| 9:13.493 | לעשות |  |
| 9:13.494 | נוח |  |
| 9:13.495 | לו |  |
| 9:13.496 | אלו |  |
| 9:13.497 | לא |  |
| 9:13.498 | נברא | . |
| 9:13.499 | ואמר |  |
| 9:13.500 | ר | ' |
| 9:13.501 | יוחנן |  |
| 9:13.502 | הלמד |  |
| 9:13.503 | שלא |  |
| 9:13.504 | לעשות |  |
| 9:13.505 | נוח |  |
| 9:13.506 | לו |  |
| 9:13.507 | אלו |  |
| 9:13.508 | נתהפכה |  |
| 9:13.509 | שליתו |  |
| 9:13.510 | על |  |
| 9:13.511 | פניו |  |
| 9:13.512 | ולא |  |
| 9:13.513 | יצא |  |
| 9:13.514 | לעולם | ? |
| 9:13.515 | ומשני |  |
| 9:13.516 | טעמיה |  |
| 9:13.517 | דרשב"י |  |
| 9:13.518 | זה |  |
| 9:13.519 | שינון |  |
| 9:13.520 | וזה |  |
| 9:13.521 | שינון | . |
| 9:13.522 | ואין |  |
| 9:13.523 | מפסיקין |  |
| 9:13.524 | שינון |  |
| 9:13.525 | מפני |  |
| 9:13.526 | שינון | . |
| 9:13.527 | ופריך |  |
| 9:13.528 | והתנינן |  |
| 9:13.529 | הקורא |  |
| 9:13.530 | מכאן |  |
| 9:13.531 | ואילך |  |
| 9:13.532 | לא |  |
| 9:13.533 | הפסיד |  |
| 9:13.534 | כאשדם |  |
| 9:13.535 | שהוא |  |
| 9:13.536 | קורא |  |
| 9:13.537 | בתורה | , |
| 9:13.538 | הא |  |
| 9:13.539 | בעונתה |  |
| 9:13.540 | חביבה |  |
| 9:13.541 | היא |  |
| 9:13.542 | יותר |  |
| 9:13.543 | מדברי |  |
| 9:13.544 | תורה | ? |
| 9:13.545 | ומשני |  |
| 9:13.546 | אמר |  |
| 9:13.547 | ר | ' |
| 9:13.548 | יודן |  |
| 9:13.549 | רשב"י |  |
| 9:13.550 | על |  |
| 9:13.551 | ידי |  |
| 9:13.552 | שהיה |  |
| 9:13.553 | חדיר |  |
| 9:13.554 | בדברי |  |
| 9:13.555 | תורה | , |
| 9:13.556 | לפיכך |  |
| 9:13.557 | אינה |  |
| 9:13.558 | חביבה |  |
| 9:13.559 | עליו |  |
| 9:13.560 | יותר |  |
| 9:13.561 | מדברי |  |
| 9:13.562 | תורה |  |
| 9:13.563 | ע"כ | . |
| 9:13.564 | שמע |  |
| 9:13.565 | מינה |  |
| 9:13.566 | מדקאמר |  |
| 9:13.567 | רשב"י |  |
| 9:13.568 | כדעתיה |  |
| 9:13.569 | וכו | ' |
| 9:13.570 | שהתנא |  |
| 9:13.571 | החסיד |  |
| 9:13.572 | רשב"י |  |
| 9:13.573 | ס"ל |  |
| 9:13.574 | שהתפלה |  |
| 9:13.575 | היא |  |
| 9:13.576 | תביעת |  |
| 9:13.577 | צרכי |  |
| 9:13.578 | הגוף |  |
| 9:13.579 | ואין |  |
| 9:13.580 | לפסוק |  |
| 9:13.581 | מדברי |  |
| 9:13.582 | תורה |  |
| 9:13.583 | שהיא |  |
| 9:13.584 | עולם |  |
| 9:13.585 | כדי |  |
| 9:13.586 | להתפלל |  |
| 9:13.587 | ולתבוע |  |
| 9:13.588 | צרכי |  |
| 9:13.589 | הגוף |  |
| 9:13.590 | ואין |  |
| 9:13.591 | לפסוק |  |
| 9:13.592 | מדברי |  |
| 9:13.593 | תורה |  |
| 9:13.594 | שהיא |  |
| 9:13.595 | עולם |  |
| 9:13.596 | כדי |  |
| 9:13.597 | להתפלל |  |
| 9:13.598 | ותבוע |  |
| 9:13.599 | צרכיו |  |
| 9:13.600 | שהם |  |
| 9:13.601 | חיי |  |
| 9:13.602 | שעה |  |
| 9:13.603 | ולכן |  |
| 9:13.604 | אמר |  |
| 9:13.605 | אלו |  |
| 9:13.606 | הוינא |  |
| 9:13.607 | קאים |  |
| 9:13.608 | וכו | ' |
| 9:13.609 | הוינא |  |
| 9:13.610 | מתבע |  |
| 9:13.611 | תרי |  |
| 9:13.612 | פומי | . |
| 9:13.613 | חד |  |
| 9:13.614 | לעסוק |  |
| 9:13.615 | בתורה | , |
| 9:13.616 | וחד |  |
| 9:13.617 | לתפלה |  |
| 9:13.618 | ולשאר |  |
| 9:13.619 | עסקי |  |
| 9:13.620 | הגוף | . |
| 9:13.621 | כי |  |
| 9:13.622 | התפלה |  |
| 9:13.623 | גם |  |
| 9:13.624 | היא |  |
| 9:13.625 | מעסק |  |
| 9:13.626 | הגוף | . |
| 9:13.627 | שמבקש |  |
| 9:13.628 | בה |  |
| 9:13.629 | רפואה |  |
| 9:13.630 | ופרנסה |  |
| 9:13.631 | ושאר |  |
| 9:13.632 | צרכים | . |
| 9:13.633 | וכדגר | ' |
| 9:13.634 | בפ | ' |
| 9:13.635 | יציאות |  |
| 9:13.636 | השבת | . |
| 9:13.637 | רבא |  |
| 9:13.638 | אשכחיה |  |
| 9:13.639 | לרב |  |
| 9:13.640 | המנונא |  |
| 9:13.641 | דקא |  |
| 9:13.642 | מאריך |  |
| 9:13.643 | בצלותא | , |
| 9:13.644 | אמר |  |
| 9:13.645 | מניחין |  |
| 9:13.646 | חי |  |
| 9:13.647 | עולם |  |
| 9:13.648 | ועוסקין |  |
| 9:13.649 | בחיי |  |
| 9:13.650 | שעה | , |
| 9:13.651 | והוא |  |
| 9:13.652 | סבר |  |
| 9:13.653 | זמן |  |
| 9:13.654 | תורה |  |
| 9:13.655 | לחוד | , |
| 9:13.656 | וזמן |  |
| 9:13.657 | תפלה |  |
| 9:13.658 | לחוד | . |
| 9:13.659 | עוד |  |
| 9:13.660 | שם |  |
| 9:13.661 | רבי |  |
| 9:13.662 | ירמיה |  |
| 9:13.663 | הזה |  |
| 9:13.664 | יתיב |  |
| 9:13.665 | קמי |  |
| 9:13.666 | ר | ' |
| 9:13.667 | זירא |  |
| 9:13.668 | וקא |  |
| 9:13.669 | גריס | . |
| 9:13.670 | מטא |  |
| 9:13.671 | זמן |  |
| 9:13.672 | תפלה |  |
| 9:13.673 | וקא |  |
| 9:13.674 | מסרהיב |  |
| 9:13.675 | קרי |  |
| 9:13.676 | עליה |  |
| 9:13.677 | רבי |  |
| 9:13.678 | זירא |  |
| 9:13.679 | מסיר |  |
| 9:13.680 | אזנו |  |
| 9:13.681 | משמוע |  |
| 9:13.682 | צורה |  |
| 9:13.683 | גם |  |
| 9:13.684 | תפלתו |  |
| 9:13.685 | תועבה |  |
| 9:13.686 | ע"כ | . |
| 9:13.687 | אלמלא |  |
| 9:13.688 | דעסק |  |
| 9:13.689 | תורה |  |
| 9:13.690 | היה |  |
| 9:13.691 | חביב |  |
| 9:13.692 | עליהם |  |
| 9:13.693 | יותר |  |
| 9:13.694 | מן |  |
| 9:13.695 | התפלה | . |
| 9:13.696 | כי |  |
| 9:13.697 | הוא |  |
| 9:13.698 | חיינו |  |
| 9:13.699 | ואורך |  |
| 9:13.700 | ימינו |  |
| 9:13.701 | לעולם |  |
| 9:13.702 | שכולו |  |
| 9:13.703 | ארוך |  |
| 9:13.704 | לעולם |  |
| 9:13.705 | שכולו |  |
| 9:13.706 | טוב | . |
| 9:13.707 | אלא |  |
| 9:13.708 | דבתר |  |
| 9:13.709 | הכי |  |
| 9:13.710 | מסקינן |  |
| 9:13.711 | דלא |  |
| 9:13.712 | שנו |  |
| 9:13.713 | דאין |  |
| 9:13.714 | מפסיקין |  |
| 9:13.715 | לתפלה |  |
| 9:13.716 | אלא |  |
| 9:13.717 | כגון |  |
| 9:13.718 | רשב"י |  |
| 9:13.719 | וחביריו |  |
| 9:13.720 | שתורתן |  |
| 9:13.721 | אומנותן | . |
| 9:13.722 | אבל |  |
| 9:13.723 | כגון |  |
| 9:13.724 | אנו |  |
| 9:13.725 | מפסיקין |  |
| 9:13.726 | בין |  |
| 9:13.727 | לק"ש |  |
| 9:13.728 | בין |  |
| 9:13.729 | לתפלה | , |
| 9:13.730 | וכן |  |
| 9:13.731 | פסקו |  |
| 9:13.732 | הרי"ף |  |
| 9:13.733 | והרמב"ם |  |
| 9:13.734 | כר | ' |
| 9:13.735 | יוחנן | , |
| 9:13.736 | והוא |  |
| 9:13.737 | כדאמרינן |  |
| 9:13.738 | בעלמא |  |
| 9:13.739 | הרבה |  |
| 9:13.740 | עשו |  |
| 9:13.741 | כרשב"י |  |
| 9:13.742 | ולא |  |
| 9:13.743 | עלתה |  |
| 9:13.744 | בידם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

