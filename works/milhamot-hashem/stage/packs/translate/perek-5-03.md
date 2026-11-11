# Translation pack: milhamot-hashem perek-5-03 (Milchamot ha-Shem 5:3–5:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-03`.

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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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


## Milchamot ha-Shem 5:3

Hebrew (Wikisource): מב) נטה לך קורא נעים מהאמונות המדומות האלה. וראה מה שכ' אחד מעמודי הוראה אשר בית ישראל נשענים עליהם הוא רבי' הראש ז"ל במס' חולין פ' השוחט והובאו דבריו במבוא התלמוד וז"ל ומסתברא שכל תוספתא שלא נתפשטה עד אחר סידור הגמרא לאו דסמכא היא דמסתמא כיון שרצו חכמי ישראל לעשות חבור קיימא אמת. חקרו ודרשו וחיזרו לידע כל הספרים שנכתבו על דברי חכמים וביררו אותם שהם בר סמכא עכ"ל. ועיין לעיל סי' י"ב וי"ג כ' הר' סעדיה גאון שדבריו מסכימים לדברי הרא"ש. ומה שכ' המגיה במבוא התלמוד בגיליון לעיין בהרא"ש דהרמב"ם לא ס"ל הכי כבר ישב הר' ב"י דברי הרמב"ם שפסק בהתוספתא ושגם מסדר הגמרא סבירא ליה כהתוספתא ושגירסת הרמב"ם בגמרא שהה כמיעוט סימנים וכבר הביאה אבי' לעיל שהה כדי שיגביה הבהמה וירביצנה וישחוט מיעוט הסימנים. וכן היא הגירסא בהרמב"ם כ"י ישנים. ונמצא שגם הרמב"ם הכי ס"ל, וכבר העיר על זה ג"כ המ"מ הנדפס מחדש. ולפי זה כולי עלמא מודו שכל חיבור המיוחס לשום תנא או אמורא שלא נתפשט רק אחר סידור הגמ' לאו דסמכא הוא. וכמו שאמרו בירושלמי שכל משנה שלא נכנסה לחבורה אינה משנה. כמו שכתבתי לעיל סי' י"ב. ואם כך היא אמונתך שכל תפלותינו ועבודותינו לסבה הראשונה שהוא ה' אלהים אמת אשר אין ראשית לראשיתו והוא שהוציא את אבותינו ממצרים וכבודו נגלה עליהם בהר סיני בתתו לנו את התורה ואמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים על פני וכו' א"כ למה תבזה את בעלי אמונתך העוסקים יומם ולילה במשנה ובתלמוד ותקרא להם מינים ואפיקורסים ותחזיק ידי זרים לאמונתך אשר בילדי נכרים ישפיקו הוא ספר הזהר הנמסר לישראל מיד גוי ויחסוהו לרשב"י התנא בשקר. והטוב והנאות לנו להחזיק באמונת קדמונינו ביחוד השי"ת בלי צירוף ושתוף. לקיים ואל תטוש תורת אמך, ואל תבוז כי זקנה אמך. רצוני לומר שלא תבוז ותלעיג על אמונת קדמונינו ביחוד השי"ת ולא תזלזל בספריהם אף על פי שבלו והזקינו ונתישנו. לא כאשר עשו בימי הבינים שזלזלו בהם ואיבדו אותן ועשו אותן כריכות לספרים חדשים המודפסים. ויש שמכרו אותן בזול מאד. ובעיני ראיתי אחד מבני עמנו שהחליף את ספר מורה הנבוכים כ"י בלשון ערבי השוה לנו עתה שנים ושלשה זהובים בספר קטן הכמות והאיכות מודפס שוה חצי ריאל. והרבה מהמונינו מכרו ספרים ישנים שירשו מאבותיהם לערלים שבאו לארץ התימן כמשה בן נתנאל שישב בבית מ"ו אברהם צאלח. והוא בעצמו מכר לו אז נביאים וכתובים עם תרגום יונתן בן עוזיאל על נביאים. ותר' ר' סעדיה גאון ערבי על ישעיה ועל כתובים. גם משנה עם פי' הרמב"ם בערבי ס' קדשים בכתב מפואר עם ציורי בית קדשינו ותפארתינו וכלי הקדש במסכת מדות. ובעיני ראיתיה בידו אני והמנוח מו"ר יחיא קרח זלה"ה. ואין לאל ידינו להוציאה מתחת ידו. ועד עתה חסר לנו סדר קדשים ונשים מהמשנה כ"י עם פי' הרמב"ם ז"ל. ולא אלו בלבד מכרו לו. כי הרבה ספרים כתבי יד עתיקים נמכרו לו גם לזולתו בזול. כי היו בידם ככלי אין חפץ בו ולא היו יודעים לקרות בהם ואנן הוא דעבדינן לתורת אבותינו שלא תשתכח מישראל בעיר תהלה צנעא, כי קבענו למודינו במשנה עם פי' הרמב"ם בערבי. ופירושי רס"ג במקרא עם תלמוד ופוסקים. ומקודם לא היו לומדים בכל בתי כנסיות בלילות כי אם בספר הזהר בלי הבין כמפורסם. והיו משתבחים בעצמם שלמדו באותה לילה כך וכך דפים. ולימוד המשנה והתלמוד היה יקר בימים ההם אין חזון נפרץ. ורק בבקר לומדים מעט משנה לפני מו"ר יוסף קרח ז"ל עם פי' ר' עובדיה בבית הכנסת שרעבי. ותחלת לימודנו היה שם, וגם אחרי כן לפני בנו הרה"ג כמו"ר חיים קרח ז"ל בלילה. ורבים ממשכילי עמנו היו באים ללמוד. ואחרי כן התחלנו ללמוד פירוש הרמב"ם בערבי בלילי שבת מחצות. ואז נשתדל ידינו המנוח מו"ר חיים אלכסאר לכתוב לו משנה בפי' הרמב"ם ז"ל. ואחרים מהמשכילים גם הם נשתדלו לכתבם וללמוד בהם. ונמשך הדבר שגם בקצת בתי כנסיות נתקנאו לעשות כמעשנו כמאמר רז"ל קנאת סופרים תרבה חכמה. ולמדו גם כן הקדמת הרמב"ם לסדר זרעים ושלש עשרה עיקרים שכ' רבינו בפרק עשירי דסנהדרין בשום שכל ויבינו בהם ובדברי ר' סעדיה גאון וחובת הלבבות שהם עומדים נגד דברי הזהר והמקובלים החדשים. וראיתי להביא פה העיקר השני מהי"ג עיקרים עם קצת מדברי אחד מקדמונינו הר"ר חטר בן שלמה בן אלמעלם בפי' העיקר הזה וז"ל. ואלקאעדה אלתאניה וחדתה תעלי ודאלך אן הדא עלה אלכל ואחד, ליס כואחד אלגנס, ולא כואחד אלנוע. ולא כאלשכץ אלואחד אלמרכב אלדי ינקסם לאחאד כתירה ולא ואחד כאלגסם אלביסט אלואחד באלעדר אלדי יקבל אלאנקסאם ואלתגזי אלי לא נהאה. בל הו תע' ואחד בוחדה ליס כמתלהא וחדה בוגה, והדה אלקאעדה אלתאניה הי אלמדלול עליהא בקולה שמע ישראל ה' אלהינו ה' אחד. שרחהא אנה למא ביין לנא ונודה תע' אפתקר לביאן וחדאנייתה, אד האתין אלקאעדתין המא אצל שריעה משה רבי ע"ה, והמא אוול כל שריעה אלאהייה. ואכר כל שריעה נאמוסייה, לאן אלשריעה אלאלאהייה תבתדי מן אלאלאהיאת ותנתהי אלי אלנסטאנייאת, ואלשריעה אלנמוסייה עכס דאלך. ביאן דאך אן אלאנסאן אלא עלם אלונוד מן אבתדאיה אלי אנתהאיה אעני מם אעלא אעאליין אלי אספל אלסאפלין, פכיך מא נזל ואכר פי אלהבוט ספלא קלת אלוחדה צרורה ואזאדת ואלכהרה וכו. ואמא אדא אכד אלאנסאן אן יתרקא פי אלוגוד מן אספל אלי אלעא, פכיך מא אזאד פי אלתרקא עלווא צעפת אלכתרה ותקוות אלוהדה אלי אלוהחד אלאחד אלדי נשת מנה כל וחדה אחדיהה פי אלוגוד, והו ואחד אלוהדאנייה, והא אנא אדבר תלאתה מדאהב, אתנוייה, ואמגוס, ואנצארא והא אנא אביין כל מדהב מנהם ומן אטלע עלי כתבהם פצל דאלך, אמא אלתנוייה פאנהם יקולון אן לאלעאלם צאנעאן דקימאן, אחדהמא אלנור, והו יפעל אלכיר פי אלוגוד ולא יקדר עלי פעל אלשר! ואלתאני אלצלמה, והו יפעל אלשר פי אלוגוד ולא יקדר עלי פעל אלכיר, והם מכתלפין, מנהם מן יקול אן אלנור ואלצלמה חיאן, ומנחם מן יקול מייתאן. ואמא אלמגוס פאנהם יקולין אן לאלעאלם צאנעא חכימא קדירא מדברא גיר גסם, תזהוה ען פעל אלשדור וסמוה יזדאן, תם אנהם אצאפו היה אלאלאם אלכיאנה פי אלעאלם לאלשיטאן אהרמין בלגתהם, והם איצא מכתלפין, מנאהם מן יקול אנה מהדת וגסם, ומנאהם מאן יקול בקרם אהרמין ואנא גיר גסם, תם אכתלפו בסבב חדותה איצא, פקאלו אנה חדת ליזדאן פכרה, פקאל פי נפסה, פכיף יכון חאלי וחדי או מן ינאזעני, וקאלו אן אלשיטאן תולד מן תלך אלפכרה, תם אנה חדאת יזדאן וצאלחה אלי מרה, ודכרו אשיא ליס הנא מוצע דכרהא, ואלכיר ענדהם גניע אעני אלתנוייה ואלמגוס הו מא תשתהיה אלנפוס, ואלשר הו מא תכרהה אלנפוס! ואמא אלנצרא פאנהם יקולו אם אללה תע ואחד (עלי אלחקיקה) תלאתה אקאנים עלי אלחקיקה, אקנים אלאב, והו דאת אלבארי ענדהם, ואקנוס אלאכן והו אלעלם (חכמה) ואוקנום רוח אלקודש והו אלחיה (בינה הנקרא אלהים חיים) ואלתלאתה ורגעו ואחד. והם פריק מן ישראל. ואסתדלאלהם עלי אנה אלתלאתה ירגעו ואחד, מן קולח תע שמע ישראל ידוד אלהינו ידוד אחד, קאלו הדה תלתלאה אסמא לתלאתה מסמיין, וכלן אלתלאתה ירגעו ואחד, כמא קאל תע ידוד אלהינו ידוד אחד, ואלתאלב אנהם יעתקדו פי אלתוחיד לתהדר מן דאלך ארגע אלי גרץ אלקאעה וכו' עכ"ל הר' חטר ז"ל ע"ש. [תרגום הרב המחבר: היסוד השני יחודו יתברך וזהו לידע שזה האחד שהוא סיבת הכול, אחד לא כאחד הסוג, ולא כאחד המין, ולא כאחד המורכב אשר יחולק לחלקים רבים, ולא כגוף הפשוט שהוא אחד במנין (כמו המים שהוא יסוד פשוט) (חסר) ומתחלק לאין קץ. אבל הוא יתברך אחד, אחדות שאין כמוה אחדות בשום פנים, וזה היסוד השני הוא הנלמד מאומרו: "שמע ישראל ה"א ה' אחד!". ביאור היסוד הזה, הואיל וכבר באר רבנו מציאותו ית', הוצרך לבאר אחדותו. כי אלו השני יסודות הם עיקר תורת משרע"ה. והמה ראשית לכל תורה אלוהית וסוף לכל תורה נימוסית. כי התורה האלוהית (מתחלתאמן האלהיות) [מתחלת באלוהיות] וגומרת בגשמיות והתורה הנימוסית היפך זה. ביאור זה שהאדם אם ידע מציאות כל הנמצאים מראש ועד סוף, ר"ך ממעלה מעלה ועד מטה מטה, כל עוד שירד בידיעתו מטה נתמעטה האחדות ונתוסף הריבוי עד שמגיע לעולם השפל שבלי ספק לא ימד ולא יספר סוגי מיניו כ"ש אישי כל מין ומין וכו'. אבל אם בא לעלות בנמצאים ממטה למעלה, הנה כל מה שיוסיף לעלות יתמעט הריבוי ותתחזק האחדות עד האחד המיוחד שממנו נמצאו כל האחדים שלכל הנמצאים. והוא אחד מיוחד. שלש אמונות יש: השונים, והאמגושים והנוצרים, והנני מבאר כ"א מהם ומי שיעיין בספריהם יתבאר לו זה. השונים אומרים שיש לעולם שני פועלים: אחד מהם האור, והוא פועל הטוב בנמצאים ולא יוכל לעשות רע, והפועל השני החושך, והוא פועל הרע בעולם ולא יוכל לעשות טוב, והמה חלוקים: יש מהם אומרים שהאור והחושך חיים וי"א שהם מתים. אמנם האמגוש הנם אומרים שיש לעולם פועל חכם גיבור מנהיג ואינו גוף, והרחיקו ממנו מעשות כל רע, וקראוהו "יזדאן". וייחסו כל הרע והצער ההווים בעולם לשטן, וקראוהו "אהרמן" בלשונם. והמה חלוקים:

| id | word | punctuation after |
|---|---|---|
| 5:3.1 | מב | ) |
| 5:3.2 | נטה |  |
| 5:3.3 | לך |  |
| 5:3.4 | קורא |  |
| 5:3.5 | נעים |  |
| 5:3.6 | מהאמונות |  |
| 5:3.7 | המדומות |  |
| 5:3.8 | האלה | . |
| 5:3.9 | וראה |  |
| 5:3.10 | מה |  |
| 5:3.11 | שכ | ' |
| 5:3.12 | אחד |  |
| 5:3.13 | מעמודי |  |
| 5:3.14 | הוראה |  |
| 5:3.15 | אשר |  |
| 5:3.16 | בית |  |
| 5:3.17 | ישראל |  |
| 5:3.18 | נשענים |  |
| 5:3.19 | עליהם |  |
| 5:3.20 | הוא |  |
| 5:3.21 | רבי | ' |
| 5:3.22 | הראש |  |
| 5:3.23 | ז"ל |  |
| 5:3.24 | במס | ' |
| 5:3.25 | חולין |  |
| 5:3.26 | פ | ' |
| 5:3.27 | השוחט |  |
| 5:3.28 | והובאו |  |
| 5:3.29 | דבריו |  |
| 5:3.30 | במבוא |  |
| 5:3.31 | התלמוד |  |
| 5:3.32 | וז"ל |  |
| 5:3.33 | ומסתברא |  |
| 5:3.34 | שכל |  |
| 5:3.35 | תוספתא |  |
| 5:3.36 | שלא |  |
| 5:3.37 | נתפשטה |  |
| 5:3.38 | עד |  |
| 5:3.39 | אחר |  |
| 5:3.40 | סידור |  |
| 5:3.41 | הגמרא |  |
| 5:3.42 | לאו |  |
| 5:3.43 | דסמכא |  |
| 5:3.44 | היא |  |
| 5:3.45 | דמסתמא |  |
| 5:3.46 | כיון |  |
| 5:3.47 | שרצו |  |
| 5:3.48 | חכמי |  |
| 5:3.49 | ישראל |  |
| 5:3.50 | לעשות |  |
| 5:3.51 | חבור |  |
| 5:3.52 | קיימא |  |
| 5:3.53 | אמת | . |
| 5:3.54 | חקרו |  |
| 5:3.55 | ודרשו |  |
| 5:3.56 | וחיזרו |  |
| 5:3.57 | לידע |  |
| 5:3.58 | כל |  |
| 5:3.59 | הספרים |  |
| 5:3.60 | שנכתבו |  |
| 5:3.61 | על |  |
| 5:3.62 | דברי |  |
| 5:3.63 | חכמים |  |
| 5:3.64 | וביררו |  |
| 5:3.65 | אותם |  |
| 5:3.66 | שהם |  |
| 5:3.67 | בר |  |
| 5:3.68 | סמכא |  |
| 5:3.69 | עכ"ל | . |
| 5:3.70 | ועיין |  |
| 5:3.71 | לעיל |  |
| 5:3.72 | סי | ' |
| 5:3.73 | י"ב |  |
| 5:3.74 | וי"ג |  |
| 5:3.75 | כ | ' |
| 5:3.76 | הר | ' |
| 5:3.77 | סעדיה |  |
| 5:3.78 | גאון |  |
| 5:3.79 | שדבריו |  |
| 5:3.80 | מסכימים |  |
| 5:3.81 | לדברי |  |
| 5:3.82 | הרא"ש | . |
| 5:3.83 | ומה |  |
| 5:3.84 | שכ | ' |
| 5:3.85 | המגיה |  |
| 5:3.86 | במבוא |  |
| 5:3.87 | התלמוד |  |
| 5:3.88 | בגיליון |  |
| 5:3.89 | לעיין |  |
| 5:3.90 | בהרא"ש |  |
| 5:3.91 | דהרמב"ם |  |
| 5:3.92 | לא |  |
| 5:3.93 | ס"ל |  |
| 5:3.94 | הכי |  |
| 5:3.95 | כבר |  |
| 5:3.96 | ישב |  |
| 5:3.97 | הר | ' |
| 5:3.98 | ב"י |  |
| 5:3.99 | דברי |  |
| 5:3.100 | הרמב"ם |  |
| 5:3.101 | שפסק |  |
| 5:3.102 | בהתוספתא |  |
| 5:3.103 | ושגם |  |
| 5:3.104 | מסדר |  |
| 5:3.105 | הגמרא |  |
| 5:3.106 | סבירא |  |
| 5:3.107 | ליה |  |
| 5:3.108 | כהתוספתא |  |
| 5:3.109 | ושגירסת |  |
| 5:3.110 | הרמב"ם |  |
| 5:3.111 | בגמרא |  |
| 5:3.112 | שהה |  |
| 5:3.113 | כמיעוט |  |
| 5:3.114 | סימנים |  |
| 5:3.115 | וכבר |  |
| 5:3.116 | הביאה |  |
| 5:3.117 | אבי | ' |
| 5:3.118 | לעיל |  |
| 5:3.119 | שהה |  |
| 5:3.120 | כדי |  |
| 5:3.121 | שיגביה |  |
| 5:3.122 | הבהמה |  |
| 5:3.123 | וירביצנה |  |
| 5:3.124 | וישחוט |  |
| 5:3.125 | מיעוט |  |
| 5:3.126 | הסימנים | . |
| 5:3.127 | וכן |  |
| 5:3.128 | היא |  |
| 5:3.129 | הגירסא |  |
| 5:3.130 | בהרמב"ם |  |
| 5:3.131 | כ"י |  |
| 5:3.132 | ישנים | . |
| 5:3.133 | ונמצא |  |
| 5:3.134 | שגם |  |
| 5:3.135 | הרמב"ם |  |
| 5:3.136 | הכי |  |
| 5:3.137 | ס"ל | , |
| 5:3.138 | וכבר |  |
| 5:3.139 | העיר |  |
| 5:3.140 | על |  |
| 5:3.141 | זה |  |
| 5:3.142 | ג"כ |  |
| 5:3.143 | המ"מ |  |
| 5:3.144 | הנדפס |  |
| 5:3.145 | מחדש | . |
| 5:3.146 | ולפי |  |
| 5:3.147 | זה |  |
| 5:3.148 | כולי |  |
| 5:3.149 | עלמא |  |
| 5:3.150 | מודו |  |
| 5:3.151 | שכל |  |
| 5:3.152 | חיבור |  |
| 5:3.153 | המיוחס |  |
| 5:3.154 | לשום |  |
| 5:3.155 | תנא |  |
| 5:3.156 | או |  |
| 5:3.157 | אמורא |  |
| 5:3.158 | שלא |  |
| 5:3.159 | נתפשט |  |
| 5:3.160 | רק |  |
| 5:3.161 | אחר |  |
| 5:3.162 | סידור |  |
| 5:3.163 | הגמ | ' |
| 5:3.164 | לאו |  |
| 5:3.165 | דסמכא |  |
| 5:3.166 | הוא | . |
| 5:3.167 | וכמו |  |
| 5:3.168 | שאמרו |  |
| 5:3.169 | בירושלמי |  |
| 5:3.170 | שכל |  |
| 5:3.171 | משנה |  |
| 5:3.172 | שלא |  |
| 5:3.173 | נכנסה |  |
| 5:3.174 | לחבורה |  |
| 5:3.175 | אינה |  |
| 5:3.176 | משנה | . |
| 5:3.177 | כמו |  |
| 5:3.178 | שכתבתי |  |
| 5:3.179 | לעיל |  |
| 5:3.180 | סי | ' |
| 5:3.181 | י"ב | . |
| 5:3.182 | ואם |  |
| 5:3.183 | כך |  |
| 5:3.184 | היא |  |
| 5:3.185 | אמונתך |  |
| 5:3.186 | שכל |  |
| 5:3.187 | תפלותינו |  |
| 5:3.188 | ועבודותינו |  |
| 5:3.189 | לסבה |  |
| 5:3.190 | הראשונה |  |
| 5:3.191 | שהוא |  |
| 5:3.192 | ה | ' |
| 5:3.193 | אלהים |  |
| 5:3.194 | אמת |  |
| 5:3.195 | אשר |  |
| 5:3.196 | אין |  |
| 5:3.197 | ראשית |  |
| 5:3.198 | לראשיתו |  |
| 5:3.199 | והוא |  |
| 5:3.200 | שהוציא |  |
| 5:3.201 | את |  |
| 5:3.202 | אבותינו |  |
| 5:3.203 | ממצרים |  |
| 5:3.204 | וכבודו |  |
| 5:3.205 | נגלה |  |
| 5:3.206 | עליהם |  |
| 5:3.207 | בהר |  |
| 5:3.208 | סיני |  |
| 5:3.209 | בתתו |  |
| 5:3.210 | לנו |  |
| 5:3.211 | את |  |
| 5:3.212 | התורה |  |
| 5:3.213 | ואמר |  |
| 5:3.214 | אנכי |  |
| 5:3.215 | ה | ' |
| 5:3.216 | אלהיך | , |
| 5:3.217 | לא |  |
| 5:3.218 | יהיה |  |
| 5:3.219 | לך |  |
| 5:3.220 | אלהים |  |
| 5:3.221 | אחרים |  |
| 5:3.222 | על |  |
| 5:3.223 | פני |  |
| 5:3.224 | וכו | ' |
| 5:3.225 | א"כ |  |
| 5:3.226 | למה |  |
| 5:3.227 | תבזה |  |
| 5:3.228 | את |  |
| 5:3.229 | בעלי |  |
| 5:3.230 | אמונתך |  |
| 5:3.231 | העוסקים |  |
| 5:3.232 | יומם |  |
| 5:3.233 | ולילה |  |
| 5:3.234 | במשנה |  |
| 5:3.235 | ובתלמוד |  |
| 5:3.236 | ותקרא |  |
| 5:3.237 | להם |  |
| 5:3.238 | מינים |  |
| 5:3.239 | ואפיקורסים |  |
| 5:3.240 | ותחזיק |  |
| 5:3.241 | ידי |  |
| 5:3.242 | זרים |  |
| 5:3.243 | לאמונתך |  |
| 5:3.244 | אשר |  |
| 5:3.245 | בילדי |  |
| 5:3.246 | נכרים |  |
| 5:3.247 | ישפיקו |  |
| 5:3.248 | הוא |  |
| 5:3.249 | ספר |  |
| 5:3.250 | הזהר |  |
| 5:3.251 | הנמסר |  |
| 5:3.252 | לישראל |  |
| 5:3.253 | מיד |  |
| 5:3.254 | גוי |  |
| 5:3.255 | ויחסוהו |  |
| 5:3.256 | לרשב"י |  |
| 5:3.257 | התנא |  |
| 5:3.258 | בשקר | . |
| 5:3.259 | והטוב |  |
| 5:3.260 | והנאות |  |
| 5:3.261 | לנו |  |
| 5:3.262 | להחזיק |  |
| 5:3.263 | באמונת |  |
| 5:3.264 | קדמונינו |  |
| 5:3.265 | ביחוד |  |
| 5:3.266 | השי"ת |  |
| 5:3.267 | בלי |  |
| 5:3.268 | צירוף |  |
| 5:3.269 | ושתוף | . |
| 5:3.270 | לקיים |  |
| 5:3.271 | ואל |  |
| 5:3.272 | תטוש |  |
| 5:3.273 | תורת |  |
| 5:3.274 | אמך | , |
| 5:3.275 | ואל |  |
| 5:3.276 | תבוז |  |
| 5:3.277 | כי |  |
| 5:3.278 | זקנה |  |
| 5:3.279 | אמך | . |
| 5:3.280 | רצוני |  |
| 5:3.281 | לומר |  |
| 5:3.282 | שלא |  |
| 5:3.283 | תבוז |  |
| 5:3.284 | ותלעיג |  |
| 5:3.285 | על |  |
| 5:3.286 | אמונת |  |
| 5:3.287 | קדמונינו |  |
| 5:3.288 | ביחוד |  |
| 5:3.289 | השי"ת |  |
| 5:3.290 | ולא |  |
| 5:3.291 | תזלזל |  |
| 5:3.292 | בספריהם |  |
| 5:3.293 | אף |  |
| 5:3.294 | על |  |
| 5:3.295 | פי |  |
| 5:3.296 | שבלו |  |
| 5:3.297 | והזקינו |  |
| 5:3.298 | ונתישנו | . |
| 5:3.299 | לא |  |
| 5:3.300 | כאשר |  |
| 5:3.301 | עשו |  |
| 5:3.302 | בימי |  |
| 5:3.303 | הבינים |  |
| 5:3.304 | שזלזלו |  |
| 5:3.305 | בהם |  |
| 5:3.306 | ואיבדו |  |
| 5:3.307 | אותן |  |
| 5:3.308 | ועשו |  |
| 5:3.309 | אותן |  |
| 5:3.310 | כריכות |  |
| 5:3.311 | לספרים |  |
| 5:3.312 | חדשים |  |
| 5:3.313 | המודפסים | . |
| 5:3.314 | ויש |  |
| 5:3.315 | שמכרו |  |
| 5:3.316 | אותן |  |
| 5:3.317 | בזול |  |
| 5:3.318 | מאד | . |
| 5:3.319 | ובעיני |  |
| 5:3.320 | ראיתי |  |
| 5:3.321 | אחד |  |
| 5:3.322 | מבני |  |
| 5:3.323 | עמנו |  |
| 5:3.324 | שהחליף |  |
| 5:3.325 | את |  |
| 5:3.326 | ספר |  |
| 5:3.327 | מורה |  |
| 5:3.328 | הנבוכים |  |
| 5:3.329 | כ"י |  |
| 5:3.330 | בלשון |  |
| 5:3.331 | ערבי |  |
| 5:3.332 | השוה |  |
| 5:3.333 | לנו |  |
| 5:3.334 | עתה |  |
| 5:3.335 | שנים |  |
| 5:3.336 | ושלשה |  |
| 5:3.337 | זהובים |  |
| 5:3.338 | בספר |  |
| 5:3.339 | קטן |  |
| 5:3.340 | הכמות |  |
| 5:3.341 | והאיכות |  |
| 5:3.342 | מודפס |  |
| 5:3.343 | שוה |  |
| 5:3.344 | חצי |  |
| 5:3.345 | ריאל | . |
| 5:3.346 | והרבה |  |
| 5:3.347 | מהמונינו |  |
| 5:3.348 | מכרו |  |
| 5:3.349 | ספרים |  |
| 5:3.350 | ישנים |  |
| 5:3.351 | שירשו |  |
| 5:3.352 | מאבותיהם |  |
| 5:3.353 | לערלים |  |
| 5:3.354 | שבאו |  |
| 5:3.355 | לארץ |  |
| 5:3.356 | התימן |  |
| 5:3.357 | כמשה |  |
| 5:3.358 | בן |  |
| 5:3.359 | נתנאל |  |
| 5:3.360 | שישב |  |
| 5:3.361 | בבית |  |
| 5:3.362 | מ"ו |  |
| 5:3.363 | אברהם |  |
| 5:3.364 | צאלח | . |
| 5:3.365 | והוא |  |
| 5:3.366 | בעצמו |  |
| 5:3.367 | מכר |  |
| 5:3.368 | לו |  |
| 5:3.369 | אז |  |
| 5:3.370 | נביאים |  |
| 5:3.371 | וכתובים |  |
| 5:3.372 | עם |  |
| 5:3.373 | תרגום |  |
| 5:3.374 | יונתן |  |
| 5:3.375 | בן |  |
| 5:3.376 | עוזיאל |  |
| 5:3.377 | על |  |
| 5:3.378 | נביאים | . |
| 5:3.379 | ותר | ' |
| 5:3.380 | ר | ' |
| 5:3.381 | סעדיה |  |
| 5:3.382 | גאון |  |
| 5:3.383 | ערבי |  |
| 5:3.384 | על |  |
| 5:3.385 | ישעיה |  |
| 5:3.386 | ועל |  |
| 5:3.387 | כתובים | . |
| 5:3.388 | גם |  |
| 5:3.389 | משנה |  |
| 5:3.390 | עם |  |
| 5:3.391 | פי | ' |
| 5:3.392 | הרמב"ם |  |
| 5:3.393 | בערבי |  |
| 5:3.394 | ס | ' |
| 5:3.395 | קדשים |  |
| 5:3.396 | בכתב |  |
| 5:3.397 | מפואר |  |
| 5:3.398 | עם |  |
| 5:3.399 | ציורי |  |
| 5:3.400 | בית |  |
| 5:3.401 | קדשינו |  |
| 5:3.402 | ותפארתינו |  |
| 5:3.403 | וכלי |  |
| 5:3.404 | הקדש |  |
| 5:3.405 | במסכת |  |
| 5:3.406 | מדות | . |
| 5:3.407 | ובעיני |  |
| 5:3.408 | ראיתיה |  |
| 5:3.409 | בידו |  |
| 5:3.410 | אני |  |
| 5:3.411 | והמנוח |  |
| 5:3.412 | מו"ר |  |
| 5:3.413 | יחיא |  |
| 5:3.414 | קרח |  |
| 5:3.415 | זלה"ה | . |
| 5:3.416 | ואין |  |
| 5:3.417 | לאל |  |
| 5:3.418 | ידינו |  |
| 5:3.419 | להוציאה |  |
| 5:3.420 | מתחת |  |
| 5:3.421 | ידו | . |
| 5:3.422 | ועד |  |
| 5:3.423 | עתה |  |
| 5:3.424 | חסר |  |
| 5:3.425 | לנו |  |
| 5:3.426 | סדר |  |
| 5:3.427 | קדשים |  |
| 5:3.428 | ונשים |  |
| 5:3.429 | מהמשנה |  |
| 5:3.430 | כ"י |  |
| 5:3.431 | עם |  |
| 5:3.432 | פי | ' |
| 5:3.433 | הרמב"ם |  |
| 5:3.434 | ז"ל | . |
| 5:3.435 | ולא |  |
| 5:3.436 | אלו |  |
| 5:3.437 | בלבד |  |
| 5:3.438 | מכרו |  |
| 5:3.439 | לו | . |
| 5:3.440 | כי |  |
| 5:3.441 | הרבה |  |
| 5:3.442 | ספרים |  |
| 5:3.443 | כתבי |  |
| 5:3.444 | יד |  |
| 5:3.445 | עתיקים |  |
| 5:3.446 | נמכרו |  |
| 5:3.447 | לו |  |
| 5:3.448 | גם |  |
| 5:3.449 | לזולתו |  |
| 5:3.450 | בזול | . |
| 5:3.451 | כי |  |
| 5:3.452 | היו |  |
| 5:3.453 | בידם |  |
| 5:3.454 | ככלי |  |
| 5:3.455 | אין |  |
| 5:3.456 | חפץ |  |
| 5:3.457 | בו |  |
| 5:3.458 | ולא |  |
| 5:3.459 | היו |  |
| 5:3.460 | יודעים |  |
| 5:3.461 | לקרות |  |
| 5:3.462 | בהם |  |
| 5:3.463 | ואנן |  |
| 5:3.464 | הוא |  |
| 5:3.465 | דעבדינן |  |
| 5:3.466 | לתורת |  |
| 5:3.467 | אבותינו |  |
| 5:3.468 | שלא |  |
| 5:3.469 | תשתכח |  |
| 5:3.470 | מישראל |  |
| 5:3.471 | בעיר |  |
| 5:3.472 | תהלה |  |
| 5:3.473 | צנעא | , |
| 5:3.474 | כי |  |
| 5:3.475 | קבענו |  |
| 5:3.476 | למודינו |  |
| 5:3.477 | במשנה |  |
| 5:3.478 | עם |  |
| 5:3.479 | פי | ' |
| 5:3.480 | הרמב"ם |  |
| 5:3.481 | בערבי | . |
| 5:3.482 | ופירושי |  |
| 5:3.483 | רס"ג |  |
| 5:3.484 | במקרא |  |
| 5:3.485 | עם |  |
| 5:3.486 | תלמוד |  |
| 5:3.487 | ופוסקים | . |
| 5:3.488 | ומקודם |  |
| 5:3.489 | לא |  |
| 5:3.490 | היו |  |
| 5:3.491 | לומדים |  |
| 5:3.492 | בכל |  |
| 5:3.493 | בתי |  |
| 5:3.494 | כנסיות |  |
| 5:3.495 | בלילות |  |
| 5:3.496 | כי |  |
| 5:3.497 | אם |  |
| 5:3.498 | בספר |  |
| 5:3.499 | הזהר |  |
| 5:3.500 | בלי |  |
| 5:3.501 | הבין |  |
| 5:3.502 | כמפורסם | . |
| 5:3.503 | והיו |  |
| 5:3.504 | משתבחים |  |
| 5:3.505 | בעצמם |  |
| 5:3.506 | שלמדו |  |
| 5:3.507 | באותה |  |
| 5:3.508 | לילה |  |
| 5:3.509 | כך |  |
| 5:3.510 | וכך |  |
| 5:3.511 | דפים | . |
| 5:3.512 | ולימוד |  |
| 5:3.513 | המשנה |  |
| 5:3.514 | והתלמוד |  |
| 5:3.515 | היה |  |
| 5:3.516 | יקר |  |
| 5:3.517 | בימים |  |
| 5:3.518 | ההם |  |
| 5:3.519 | אין |  |
| 5:3.520 | חזון |  |
| 5:3.521 | נפרץ | . |
| 5:3.522 | ורק |  |
| 5:3.523 | בבקר |  |
| 5:3.524 | לומדים |  |
| 5:3.525 | מעט |  |
| 5:3.526 | משנה |  |
| 5:3.527 | לפני |  |
| 5:3.528 | מו"ר |  |
| 5:3.529 | יוסף |  |
| 5:3.530 | קרח |  |
| 5:3.531 | ז"ל |  |
| 5:3.532 | עם |  |
| 5:3.533 | פי | ' |
| 5:3.534 | ר | ' |
| 5:3.535 | עובדיה |  |
| 5:3.536 | בבית |  |
| 5:3.537 | הכנסת |  |
| 5:3.538 | שרעבי | . |
| 5:3.539 | ותחלת |  |
| 5:3.540 | לימודנו |  |
| 5:3.541 | היה |  |
| 5:3.542 | שם | , |
| 5:3.543 | וגם |  |
| 5:3.544 | אחרי |  |
| 5:3.545 | כן |  |
| 5:3.546 | לפני |  |
| 5:3.547 | בנו |  |
| 5:3.548 | הרה"ג |  |
| 5:3.549 | כמו"ר |  |
| 5:3.550 | חיים |  |
| 5:3.551 | קרח |  |
| 5:3.552 | ז"ל |  |
| 5:3.553 | בלילה | . |
| 5:3.554 | ורבים |  |
| 5:3.555 | ממשכילי |  |
| 5:3.556 | עמנו |  |
| 5:3.557 | היו |  |
| 5:3.558 | באים |  |
| 5:3.559 | ללמוד | . |
| 5:3.560 | ואחרי |  |
| 5:3.561 | כן |  |
| 5:3.562 | התחלנו |  |
| 5:3.563 | ללמוד |  |
| 5:3.564 | פירוש |  |
| 5:3.565 | הרמב"ם |  |
| 5:3.566 | בערבי |  |
| 5:3.567 | בלילי |  |
| 5:3.568 | שבת |  |
| 5:3.569 | מחצות | . |
| 5:3.570 | ואז |  |
| 5:3.571 | נשתדל |  |
| 5:3.572 | ידינו |  |
| 5:3.573 | המנוח |  |
| 5:3.574 | מו"ר |  |
| 5:3.575 | חיים |  |
| 5:3.576 | אלכסאר |  |
| 5:3.577 | לכתוב |  |
| 5:3.578 | לו |  |
| 5:3.579 | משנה |  |
| 5:3.580 | בפי | ' |
| 5:3.581 | הרמב"ם |  |
| 5:3.582 | ז"ל | . |
| 5:3.583 | ואחרים |  |
| 5:3.584 | מהמשכילים |  |
| 5:3.585 | גם |  |
| 5:3.586 | הם |  |
| 5:3.587 | נשתדלו |  |
| 5:3.588 | לכתבם |  |
| 5:3.589 | וללמוד |  |
| 5:3.590 | בהם | . |
| 5:3.591 | ונמשך |  |
| 5:3.592 | הדבר |  |
| 5:3.593 | שגם |  |
| 5:3.594 | בקצת |  |
| 5:3.595 | בתי |  |
| 5:3.596 | כנסיות |  |
| 5:3.597 | נתקנאו |  |
| 5:3.598 | לעשות |  |
| 5:3.599 | כמעשנו |  |
| 5:3.600 | כמאמר |  |
| 5:3.601 | רז"ל |  |
| 5:3.602 | קנאת |  |
| 5:3.603 | סופרים |  |
| 5:3.604 | תרבה |  |
| 5:3.605 | חכמה | . |
| 5:3.606 | ולמדו |  |
| 5:3.607 | גם |  |
| 5:3.608 | כן |  |
| 5:3.609 | הקדמת |  |
| 5:3.610 | הרמב"ם |  |
| 5:3.611 | לסדר |  |
| 5:3.612 | זרעים |  |
| 5:3.613 | ושלש |  |
| 5:3.614 | עשרה |  |
| 5:3.615 | עיקרים |  |
| 5:3.616 | שכ | ' |
| 5:3.617 | רבינו |  |
| 5:3.618 | בפרק |  |
| 5:3.619 | עשירי |  |
| 5:3.620 | דסנהדרין |  |
| 5:3.621 | בשום |  |
| 5:3.622 | שכל |  |
| 5:3.623 | ויבינו |  |
| 5:3.624 | בהם |  |
| 5:3.625 | ובדברי |  |
| 5:3.626 | ר | ' |
| 5:3.627 | סעדיה |  |
| 5:3.628 | גאון |  |
| 5:3.629 | וחובת |  |
| 5:3.630 | הלבבות |  |
| 5:3.631 | שהם |  |
| 5:3.632 | עומדים |  |
| 5:3.633 | נגד |  |
| 5:3.634 | דברי |  |
| 5:3.635 | הזהר |  |
| 5:3.636 | והמקובלים |  |
| 5:3.637 | החדשים | . |
| 5:3.638 | וראיתי |  |
| 5:3.639 | להביא |  |
| 5:3.640 | פה |  |
| 5:3.641 | העיקר |  |
| 5:3.642 | השני |  |
| 5:3.643 | מהי"ג |  |
| 5:3.644 | עיקרים |  |
| 5:3.645 | עם |  |
| 5:3.646 | קצת |  |
| 5:3.647 | מדברי |  |
| 5:3.648 | אחד |  |
| 5:3.649 | מקדמונינו |  |
| 5:3.650 | הר"ר |  |
| 5:3.651 | חטר |  |
| 5:3.652 | בן |  |
| 5:3.653 | שלמה |  |
| 5:3.654 | בן |  |
| 5:3.655 | אלמעלם |  |
| 5:3.656 | בפי | ' |
| 5:3.657 | העיקר |  |
| 5:3.658 | הזה |  |
| 5:3.659 | וז"ל | . |
| 5:3.660 | ואלקאעדה |  |
| 5:3.661 | אלתאניה |  |
| 5:3.662 | וחדתה |  |
| 5:3.663 | תעלי |  |
| 5:3.664 | ודאלך |  |
| 5:3.665 | אן |  |
| 5:3.666 | הדא |  |
| 5:3.667 | עלה |  |
| 5:3.668 | אלכל |  |
| 5:3.669 | ואחד | , |
| 5:3.670 | ליס |  |
| 5:3.671 | כואחד |  |
| 5:3.672 | אלגנס | , |
| 5:3.673 | ולא |  |
| 5:3.674 | כואחד |  |
| 5:3.675 | אלנוע | . |
| 5:3.676 | ולא |  |
| 5:3.677 | כאלשכץ |  |
| 5:3.678 | אלואחד |  |
| 5:3.679 | אלמרכב |  |
| 5:3.680 | אלדי |  |
| 5:3.681 | ינקסם |  |
| 5:3.682 | לאחאד |  |
| 5:3.683 | כתירה |  |
| 5:3.684 | ולא |  |
| 5:3.685 | ואחד |  |
| 5:3.686 | כאלגסם |  |
| 5:3.687 | אלביסט |  |
| 5:3.688 | אלואחד |  |
| 5:3.689 | באלעדר |  |
| 5:3.690 | אלדי |  |
| 5:3.691 | יקבל |  |
| 5:3.692 | אלאנקסאם |  |
| 5:3.693 | ואלתגזי |  |
| 5:3.694 | אלי |  |
| 5:3.695 | לא |  |
| 5:3.696 | נהאה | . |
| 5:3.697 | בל |  |
| 5:3.698 | הו |  |
| 5:3.699 | תע | ' |
| 5:3.700 | ואחד |  |
| 5:3.701 | בוחדה |  |
| 5:3.702 | ליס |  |
| 5:3.703 | כמתלהא |  |
| 5:3.704 | וחדה |  |
| 5:3.705 | בוגה | , |
| 5:3.706 | והדה |  |
| 5:3.707 | אלקאעדה |  |
| 5:3.708 | אלתאניה |  |
| 5:3.709 | הי |  |
| 5:3.710 | אלמדלול |  |
| 5:3.711 | עליהא |  |
| 5:3.712 | בקולה |  |
| 5:3.713 | שמע |  |
| 5:3.714 | ישראל |  |
| 5:3.715 | ה | ' |
| 5:3.716 | אלהינו |  |
| 5:3.717 | ה | ' |
| 5:3.718 | אחד | . |
| 5:3.719 | שרחהא |  |
| 5:3.720 | אנה |  |
| 5:3.721 | למא |  |
| 5:3.722 | ביין |  |
| 5:3.723 | לנא |  |
| 5:3.724 | ונודה |  |
| 5:3.725 | תע | ' |
| 5:3.726 | אפתקר |  |
| 5:3.727 | לביאן |  |
| 5:3.728 | וחדאנייתה | , |
| 5:3.729 | אד |  |
| 5:3.730 | האתין |  |
| 5:3.731 | אלקאעדתין |  |
| 5:3.732 | המא |  |
| 5:3.733 | אצל |  |
| 5:3.734 | שריעה |  |
| 5:3.735 | משה |  |
| 5:3.736 | רבי |  |
| 5:3.737 | ע"ה | , |
| 5:3.738 | והמא |  |
| 5:3.739 | אוול |  |
| 5:3.740 | כל |  |
| 5:3.741 | שריעה |  |
| 5:3.742 | אלאהייה | . |
| 5:3.743 | ואכר |  |
| 5:3.744 | כל |  |
| 5:3.745 | שריעה |  |
| 5:3.746 | נאמוסייה | , |
| 5:3.747 | לאן |  |
| 5:3.748 | אלשריעה |  |
| 5:3.749 | אלאלאהייה |  |
| 5:3.750 | תבתדי |  |
| 5:3.751 | מן |  |
| 5:3.752 | אלאלאהיאת |  |
| 5:3.753 | ותנתהי |  |
| 5:3.754 | אלי |  |
| 5:3.755 | אלנסטאנייאת | , |
| 5:3.756 | ואלשריעה |  |
| 5:3.757 | אלנמוסייה |  |
| 5:3.758 | עכס |  |
| 5:3.759 | דאלך | . |
| 5:3.760 | ביאן |  |
| 5:3.761 | דאך |  |
| 5:3.762 | אן |  |
| 5:3.763 | אלאנסאן |  |
| 5:3.764 | אלא |  |
| 5:3.765 | עלם |  |
| 5:3.766 | אלונוד |  |
| 5:3.767 | מן |  |
| 5:3.768 | אבתדאיה |  |
| 5:3.769 | אלי |  |
| 5:3.770 | אנתהאיה |  |
| 5:3.771 | אעני |  |
| 5:3.772 | מם |  |
| 5:3.773 | אעלא |  |
| 5:3.774 | אעאליין |  |
| 5:3.775 | אלי |  |
| 5:3.776 | אספל |  |
| 5:3.777 | אלסאפלין | , |
| 5:3.778 | פכיך |  |
| 5:3.779 | מא |  |
| 5:3.780 | נזל |  |
| 5:3.781 | ואכר |  |
| 5:3.782 | פי |  |
| 5:3.783 | אלהבוט |  |
| 5:3.784 | ספלא |  |
| 5:3.785 | קלת |  |
| 5:3.786 | אלוחדה |  |
| 5:3.787 | צרורה |  |
| 5:3.788 | ואזאדת |  |
| 5:3.789 | ואלכהרה |  |
| 5:3.790 | וכו | . |
| 5:3.791 | ואמא |  |
| 5:3.792 | אדא |  |
| 5:3.793 | אכד |  |
| 5:3.794 | אלאנסאן |  |
| 5:3.795 | אן |  |
| 5:3.796 | יתרקא |  |
| 5:3.797 | פי |  |
| 5:3.798 | אלוגוד |  |
| 5:3.799 | מן |  |
| 5:3.800 | אספל |  |
| 5:3.801 | אלי |  |
| 5:3.802 | אלעא | , |
| 5:3.803 | פכיך |  |
| 5:3.804 | מא |  |
| 5:3.805 | אזאד |  |
| 5:3.806 | פי |  |
| 5:3.807 | אלתרקא |  |
| 5:3.808 | עלווא |  |
| 5:3.809 | צעפת |  |
| 5:3.810 | אלכתרה |  |
| 5:3.811 | ותקוות |  |
| 5:3.812 | אלוהדה |  |
| 5:3.813 | אלי |  |
| 5:3.814 | אלוהחד |  |
| 5:3.815 | אלאחד |  |
| 5:3.816 | אלדי |  |
| 5:3.817 | נשת |  |
| 5:3.818 | מנה |  |
| 5:3.819 | כל |  |
| 5:3.820 | וחדה |  |
| 5:3.821 | אחדיהה |  |
| 5:3.822 | פי |  |
| 5:3.823 | אלוגוד | , |
| 5:3.824 | והו |  |
| 5:3.825 | ואחד |  |
| 5:3.826 | אלוהדאנייה | , |
| 5:3.827 | והא |  |
| 5:3.828 | אנא |  |
| 5:3.829 | אדבר |  |
| 5:3.830 | תלאתה |  |
| 5:3.831 | מדאהב | , |
| 5:3.832 | אתנוייה | , |
| 5:3.833 | ואמגוס | , |
| 5:3.834 | ואנצארא |  |
| 5:3.835 | והא |  |
| 5:3.836 | אנא |  |
| 5:3.837 | אביין |  |
| 5:3.838 | כל |  |
| 5:3.839 | מדהב |  |
| 5:3.840 | מנהם |  |
| 5:3.841 | ומן |  |
| 5:3.842 | אטלע |  |
| 5:3.843 | עלי |  |
| 5:3.844 | כתבהם |  |
| 5:3.845 | פצל |  |
| 5:3.846 | דאלך | , |
| 5:3.847 | אמא |  |
| 5:3.848 | אלתנוייה |  |
| 5:3.849 | פאנהם |  |
| 5:3.850 | יקולון |  |
| 5:3.851 | אן |  |
| 5:3.852 | לאלעאלם |  |
| 5:3.853 | צאנעאן |  |
| 5:3.854 | דקימאן | , |
| 5:3.855 | אחדהמא |  |
| 5:3.856 | אלנור | , |
| 5:3.857 | והו |  |
| 5:3.858 | יפעל |  |
| 5:3.859 | אלכיר |  |
| 5:3.860 | פי |  |
| 5:3.861 | אלוגוד |  |
| 5:3.862 | ולא |  |
| 5:3.863 | יקדר |  |
| 5:3.864 | עלי |  |
| 5:3.865 | פעל |  |
| 5:3.866 | אלשר | ! |
| 5:3.867 | ואלתאני |  |
| 5:3.868 | אלצלמה | , |
| 5:3.869 | והו |  |
| 5:3.870 | יפעל |  |
| 5:3.871 | אלשר |  |
| 5:3.872 | פי |  |
| 5:3.873 | אלוגוד |  |
| 5:3.874 | ולא |  |
| 5:3.875 | יקדר |  |
| 5:3.876 | עלי |  |
| 5:3.877 | פעל |  |
| 5:3.878 | אלכיר | , |
| 5:3.879 | והם |  |
| 5:3.880 | מכתלפין | , |
| 5:3.881 | מנהם |  |
| 5:3.882 | מן |  |
| 5:3.883 | יקול |  |
| 5:3.884 | אן |  |
| 5:3.885 | אלנור |  |
| 5:3.886 | ואלצלמה |  |
| 5:3.887 | חיאן | , |
| 5:3.888 | ומנחם |  |
| 5:3.889 | מן |  |
| 5:3.890 | יקול |  |
| 5:3.891 | מייתאן | . |
| 5:3.892 | ואמא |  |
| 5:3.893 | אלמגוס |  |
| 5:3.894 | פאנהם |  |
| 5:3.895 | יקולין |  |
| 5:3.896 | אן |  |
| 5:3.897 | לאלעאלם |  |
| 5:3.898 | צאנעא |  |
| 5:3.899 | חכימא |  |
| 5:3.900 | קדירא |  |
| 5:3.901 | מדברא |  |
| 5:3.902 | גיר |  |
| 5:3.903 | גסם | , |
| 5:3.904 | תזהוה |  |
| 5:3.905 | ען |  |
| 5:3.906 | פעל |  |
| 5:3.907 | אלשדור |  |
| 5:3.908 | וסמוה |  |
| 5:3.909 | יזדאן | , |
| 5:3.910 | תם |  |
| 5:3.911 | אנהם |  |
| 5:3.912 | אצאפו |  |
| 5:3.913 | היה |  |
| 5:3.914 | אלאלאם |  |
| 5:3.915 | אלכיאנה |  |
| 5:3.916 | פי |  |
| 5:3.917 | אלעאלם |  |
| 5:3.918 | לאלשיטאן |  |
| 5:3.919 | אהרמין |  |
| 5:3.920 | בלגתהם | , |
| 5:3.921 | והם |  |
| 5:3.922 | איצא |  |
| 5:3.923 | מכתלפין | , |
| 5:3.924 | מנאהם |  |
| 5:3.925 | מן |  |
| 5:3.926 | יקול |  |
| 5:3.927 | אנה |  |
| 5:3.928 | מהדת |  |
| 5:3.929 | וגסם | , |
| 5:3.930 | ומנאהם |  |
| 5:3.931 | מאן |  |
| 5:3.932 | יקול |  |
| 5:3.933 | בקרם |  |
| 5:3.934 | אהרמין |  |
| 5:3.935 | ואנא |  |
| 5:3.936 | גיר |  |
| 5:3.937 | גסם | , |
| 5:3.938 | תם |  |
| 5:3.939 | אכתלפו |  |
| 5:3.940 | בסבב |  |
| 5:3.941 | חדותה |  |
| 5:3.942 | איצא | , |
| 5:3.943 | פקאלו |  |
| 5:3.944 | אנה |  |
| 5:3.945 | חדת |  |
| 5:3.946 | ליזדאן |  |
| 5:3.947 | פכרה | , |
| 5:3.948 | פקאל |  |
| 5:3.949 | פי |  |
| 5:3.950 | נפסה | , |
| 5:3.951 | פכיף |  |
| 5:3.952 | יכון |  |
| 5:3.953 | חאלי |  |
| 5:3.954 | וחדי |  |
| 5:3.955 | או |  |
| 5:3.956 | מן |  |
| 5:3.957 | ינאזעני | , |
| 5:3.958 | וקאלו |  |
| 5:3.959 | אן |  |
| 5:3.960 | אלשיטאן |  |
| 5:3.961 | תולד |  |
| 5:3.962 | מן |  |
| 5:3.963 | תלך |  |
| 5:3.964 | אלפכרה | , |
| 5:3.965 | תם |  |
| 5:3.966 | אנה |  |
| 5:3.967 | חדאת |  |
| 5:3.968 | יזדאן |  |
| 5:3.969 | וצאלחה |  |
| 5:3.970 | אלי |  |
| 5:3.971 | מרה | , |
| 5:3.972 | ודכרו |  |
| 5:3.973 | אשיא |  |
| 5:3.974 | ליס |  |
| 5:3.975 | הנא |  |
| 5:3.976 | מוצע |  |
| 5:3.977 | דכרהא | , |
| 5:3.978 | ואלכיר |  |
| 5:3.979 | ענדהם |  |
| 5:3.980 | גניע |  |
| 5:3.981 | אעני |  |
| 5:3.982 | אלתנוייה |  |
| 5:3.983 | ואלמגוס |  |
| 5:3.984 | הו |  |
| 5:3.985 | מא |  |
| 5:3.986 | תשתהיה |  |
| 5:3.987 | אלנפוס | , |
| 5:3.988 | ואלשר |  |
| 5:3.989 | הו |  |
| 5:3.990 | מא |  |
| 5:3.991 | תכרהה |  |
| 5:3.992 | אלנפוס | ! |
| 5:3.993 | ואמא |  |
| 5:3.994 | אלנצרא |  |
| 5:3.995 | פאנהם |  |
| 5:3.996 | יקולו |  |
| 5:3.997 | אם |  |
| 5:3.998 | אללה |  |
| 5:3.999 | תע |  |
| 5:3.1000 | ואחד |  |
| 5:3.1001 | עלי |  |
| 5:3.1002 | אלחקיקה | ) |
| 5:3.1003 | תלאתה |  |
| 5:3.1004 | אקאנים |  |
| 5:3.1005 | עלי |  |
| 5:3.1006 | אלחקיקה | , |
| 5:3.1007 | אקנים |  |
| 5:3.1008 | אלאב | , |
| 5:3.1009 | והו |  |
| 5:3.1010 | דאת |  |
| 5:3.1011 | אלבארי |  |
| 5:3.1012 | ענדהם | , |
| 5:3.1013 | ואקנוס |  |
| 5:3.1014 | אלאכן |  |
| 5:3.1015 | והו |  |
| 5:3.1016 | אלעלם |  |
| 5:3.1017 | חכמה | ) |
| 5:3.1018 | ואוקנום |  |
| 5:3.1019 | רוח |  |
| 5:3.1020 | אלקודש |  |
| 5:3.1021 | והו |  |
| 5:3.1022 | אלחיה |  |
| 5:3.1023 | בינה |  |
| 5:3.1024 | הנקרא |  |
| 5:3.1025 | אלהים |  |
| 5:3.1026 | חיים | ) |
| 5:3.1027 | ואלתלאתה |  |
| 5:3.1028 | ורגעו |  |
| 5:3.1029 | ואחד | . |
| 5:3.1030 | והם |  |
| 5:3.1031 | פריק |  |
| 5:3.1032 | מן |  |
| 5:3.1033 | ישראל | . |
| 5:3.1034 | ואסתדלאלהם |  |
| 5:3.1035 | עלי |  |
| 5:3.1036 | אנה |  |
| 5:3.1037 | אלתלאתה |  |
| 5:3.1038 | ירגעו |  |
| 5:3.1039 | ואחד | , |
| 5:3.1040 | מן |  |
| 5:3.1041 | קולח |  |
| 5:3.1042 | תע |  |
| 5:3.1043 | שמע |  |
| 5:3.1044 | ישראל |  |
| 5:3.1045 | ידוד |  |
| 5:3.1046 | אלהינו |  |
| 5:3.1047 | ידוד |  |
| 5:3.1048 | אחד | , |
| 5:3.1049 | קאלו |  |
| 5:3.1050 | הדה |  |
| 5:3.1051 | תלתלאה |  |
| 5:3.1052 | אסמא |  |
| 5:3.1053 | לתלאתה |  |
| 5:3.1054 | מסמיין | , |
| 5:3.1055 | וכלן |  |
| 5:3.1056 | אלתלאתה |  |
| 5:3.1057 | ירגעו |  |
| 5:3.1058 | ואחד | , |
| 5:3.1059 | כמא |  |
| 5:3.1060 | קאל |  |
| 5:3.1061 | תע |  |
| 5:3.1062 | ידוד |  |
| 5:3.1063 | אלהינו |  |
| 5:3.1064 | ידוד |  |
| 5:3.1065 | אחד | , |
| 5:3.1066 | ואלתאלב |  |
| 5:3.1067 | אנהם |  |
| 5:3.1068 | יעתקדו |  |
| 5:3.1069 | פי |  |
| 5:3.1070 | אלתוחיד |  |
| 5:3.1071 | לתהדר |  |
| 5:3.1072 | מן |  |
| 5:3.1073 | דאלך |  |
| 5:3.1074 | ארגע |  |
| 5:3.1075 | אלי |  |
| 5:3.1076 | גרץ |  |
| 5:3.1077 | אלקאעה |  |
| 5:3.1078 | וכו | ' |
| 5:3.1079 | עכ"ל |  |
| 5:3.1080 | הר | ' |
| 5:3.1081 | חטר |  |
| 5:3.1082 | ז"ל |  |
| 5:3.1083 | ע"ש | . |
| 5:3.1084 | תרגום |  |
| 5:3.1085 | הרב |  |
| 5:3.1086 | המחבר | : |
| 5:3.1087 | היסוד |  |
| 5:3.1088 | השני |  |
| 5:3.1089 | יחודו |  |
| 5:3.1090 | יתברך |  |
| 5:3.1091 | וזהו |  |
| 5:3.1092 | לידע |  |
| 5:3.1093 | שזה |  |
| 5:3.1094 | האחד |  |
| 5:3.1095 | שהוא |  |
| 5:3.1096 | סיבת |  |
| 5:3.1097 | הכול | , |
| 5:3.1098 | אחד |  |
| 5:3.1099 | לא |  |
| 5:3.1100 | כאחד |  |
| 5:3.1101 | הסוג | , |
| 5:3.1102 | ולא |  |
| 5:3.1103 | כאחד |  |
| 5:3.1104 | המין | , |
| 5:3.1105 | ולא |  |
| 5:3.1106 | כאחד |  |
| 5:3.1107 | המורכב |  |
| 5:3.1108 | אשר |  |
| 5:3.1109 | יחולק |  |
| 5:3.1110 | לחלקים |  |
| 5:3.1111 | רבים | , |
| 5:3.1112 | ולא |  |
| 5:3.1113 | כגוף |  |
| 5:3.1114 | הפשוט |  |
| 5:3.1115 | שהוא |  |
| 5:3.1116 | אחד |  |
| 5:3.1117 | במנין |  |
| 5:3.1118 | כמו |  |
| 5:3.1119 | המים |  |
| 5:3.1120 | שהוא |  |
| 5:3.1121 | יסוד |  |
| 5:3.1122 | פשוט | ) |
| 5:3.1123 | חסר | ) |
| 5:3.1124 | ומתחלק |  |
| 5:3.1125 | לאין |  |
| 5:3.1126 | קץ | . |
| 5:3.1127 | אבל |  |
| 5:3.1128 | הוא |  |
| 5:3.1129 | יתברך |  |
| 5:3.1130 | אחד | , |
| 5:3.1131 | אחדות |  |
| 5:3.1132 | שאין |  |
| 5:3.1133 | כמוה |  |
| 5:3.1134 | אחדות |  |
| 5:3.1135 | בשום |  |
| 5:3.1136 | פנים | , |
| 5:3.1137 | וזה |  |
| 5:3.1138 | היסוד |  |
| 5:3.1139 | השני |  |
| 5:3.1140 | הוא |  |
| 5:3.1141 | הנלמד |  |
| 5:3.1142 | מאומרו | : |
| 5:3.1143 | שמע |  |
| 5:3.1144 | ישראל |  |
| 5:3.1145 | ה"א |  |
| 5:3.1146 | ה | ' |
| 5:3.1147 | אחד | !". |
| 5:3.1148 | ביאור |  |
| 5:3.1149 | היסוד |  |
| 5:3.1150 | הזה | , |
| 5:3.1151 | הואיל |  |
| 5:3.1152 | וכבר |  |
| 5:3.1153 | באר |  |
| 5:3.1154 | רבנו |  |
| 5:3.1155 | מציאותו |  |
| 5:3.1156 | ית | ', |
| 5:3.1157 | הוצרך |  |
| 5:3.1158 | לבאר |  |
| 5:3.1159 | אחדותו | . |
| 5:3.1160 | כי |  |
| 5:3.1161 | אלו |  |
| 5:3.1162 | השני |  |
| 5:3.1163 | יסודות |  |
| 5:3.1164 | הם |  |
| 5:3.1165 | עיקר |  |
| 5:3.1166 | תורת |  |
| 5:3.1167 | משרע"ה | . |
| 5:3.1168 | והמה |  |
| 5:3.1169 | ראשית |  |
| 5:3.1170 | לכל |  |
| 5:3.1171 | תורה |  |
| 5:3.1172 | אלוהית |  |
| 5:3.1173 | וסוף |  |
| 5:3.1174 | לכל |  |
| 5:3.1175 | תורה |  |
| 5:3.1176 | נימוסית | . |
| 5:3.1177 | כי |  |
| 5:3.1178 | התורה |  |
| 5:3.1179 | האלוהית |  |
| 5:3.1180 | מתחלתאמן |  |
| 5:3.1181 | האלהיות | ) |
| 5:3.1182 | מתחלת |  |
| 5:3.1183 | באלוהיות | ] |
| 5:3.1184 | וגומרת |  |
| 5:3.1185 | בגשמיות |  |
| 5:3.1186 | והתורה |  |
| 5:3.1187 | הנימוסית |  |
| 5:3.1188 | היפך |  |
| 5:3.1189 | זה | . |
| 5:3.1190 | ביאור |  |
| 5:3.1191 | זה |  |
| 5:3.1192 | שהאדם |  |
| 5:3.1193 | אם |  |
| 5:3.1194 | ידע |  |
| 5:3.1195 | מציאות |  |
| 5:3.1196 | כל |  |
| 5:3.1197 | הנמצאים |  |
| 5:3.1198 | מראש |  |
| 5:3.1199 | ועד |  |
| 5:3.1200 | סוף | , |
| 5:3.1201 | ר"ך |  |
| 5:3.1202 | ממעלה |  |
| 5:3.1203 | מעלה |  |
| 5:3.1204 | ועד |  |
| 5:3.1205 | מטה |  |
| 5:3.1206 | מטה | , |
| 5:3.1207 | כל |  |
| 5:3.1208 | עוד |  |
| 5:3.1209 | שירד |  |
| 5:3.1210 | בידיעתו |  |
| 5:3.1211 | מטה |  |
| 5:3.1212 | נתמעטה |  |
| 5:3.1213 | האחדות |  |
| 5:3.1214 | ונתוסף |  |
| 5:3.1215 | הריבוי |  |
| 5:3.1216 | עד |  |
| 5:3.1217 | שמגיע |  |
| 5:3.1218 | לעולם |  |
| 5:3.1219 | השפל |  |
| 5:3.1220 | שבלי |  |
| 5:3.1221 | ספק |  |
| 5:3.1222 | לא |  |
| 5:3.1223 | ימד |  |
| 5:3.1224 | ולא |  |
| 5:3.1225 | יספר |  |
| 5:3.1226 | סוגי |  |
| 5:3.1227 | מיניו |  |
| 5:3.1228 | כ"ש |  |
| 5:3.1229 | אישי |  |
| 5:3.1230 | כל |  |
| 5:3.1231 | מין |  |
| 5:3.1232 | ומין |  |
| 5:3.1233 | וכו | '. |
| 5:3.1234 | אבל |  |
| 5:3.1235 | אם |  |
| 5:3.1236 | בא |  |
| 5:3.1237 | לעלות |  |
| 5:3.1238 | בנמצאים |  |
| 5:3.1239 | ממטה |  |
| 5:3.1240 | למעלה | , |
| 5:3.1241 | הנה |  |
| 5:3.1242 | כל |  |
| 5:3.1243 | מה |  |
| 5:3.1244 | שיוסיף |  |
| 5:3.1245 | לעלות |  |
| 5:3.1246 | יתמעט |  |
| 5:3.1247 | הריבוי |  |
| 5:3.1248 | ותתחזק |  |
| 5:3.1249 | האחדות |  |
| 5:3.1250 | עד |  |
| 5:3.1251 | האחד |  |
| 5:3.1252 | המיוחד |  |
| 5:3.1253 | שממנו |  |
| 5:3.1254 | נמצאו |  |
| 5:3.1255 | כל |  |
| 5:3.1256 | האחדים |  |
| 5:3.1257 | שלכל |  |
| 5:3.1258 | הנמצאים | . |
| 5:3.1259 | והוא |  |
| 5:3.1260 | אחד |  |
| 5:3.1261 | מיוחד | . |
| 5:3.1262 | שלש |  |
| 5:3.1263 | אמונות |  |
| 5:3.1264 | יש | : |
| 5:3.1265 | השונים | , |
| 5:3.1266 | והאמגושים |  |
| 5:3.1267 | והנוצרים | , |
| 5:3.1268 | והנני |  |
| 5:3.1269 | מבאר |  |
| 5:3.1270 | כ"א |  |
| 5:3.1271 | מהם |  |
| 5:3.1272 | ומי |  |
| 5:3.1273 | שיעיין |  |
| 5:3.1274 | בספריהם |  |
| 5:3.1275 | יתבאר |  |
| 5:3.1276 | לו |  |
| 5:3.1277 | זה | . |
| 5:3.1278 | השונים |  |
| 5:3.1279 | אומרים |  |
| 5:3.1280 | שיש |  |
| 5:3.1281 | לעולם |  |
| 5:3.1282 | שני |  |
| 5:3.1283 | פועלים | : |
| 5:3.1284 | אחד |  |
| 5:3.1285 | מהם |  |
| 5:3.1286 | האור | , |
| 5:3.1287 | והוא |  |
| 5:3.1288 | פועל |  |
| 5:3.1289 | הטוב |  |
| 5:3.1290 | בנמצאים |  |
| 5:3.1291 | ולא |  |
| 5:3.1292 | יוכל |  |
| 5:3.1293 | לעשות |  |
| 5:3.1294 | רע | , |
| 5:3.1295 | והפועל |  |
| 5:3.1296 | השני |  |
| 5:3.1297 | החושך | , |
| 5:3.1298 | והוא |  |
| 5:3.1299 | פועל |  |
| 5:3.1300 | הרע |  |
| 5:3.1301 | בעולם |  |
| 5:3.1302 | ולא |  |
| 5:3.1303 | יוכל |  |
| 5:3.1304 | לעשות |  |
| 5:3.1305 | טוב | , |
| 5:3.1306 | והמה |  |
| 5:3.1307 | חלוקים | : |
| 5:3.1308 | יש |  |
| 5:3.1309 | מהם |  |
| 5:3.1310 | אומרים |  |
| 5:3.1311 | שהאור |  |
| 5:3.1312 | והחושך |  |
| 5:3.1313 | חיים |  |
| 5:3.1314 | וי"א |  |
| 5:3.1315 | שהם |  |
| 5:3.1316 | מתים | . |
| 5:3.1317 | אמנם |  |
| 5:3.1318 | האמגוש |  |
| 5:3.1319 | הנם |  |
| 5:3.1320 | אומרים |  |
| 5:3.1321 | שיש |  |
| 5:3.1322 | לעולם |  |
| 5:3.1323 | פועל |  |
| 5:3.1324 | חכם |  |
| 5:3.1325 | גיבור |  |
| 5:3.1326 | מנהיג |  |
| 5:3.1327 | ואינו |  |
| 5:3.1328 | גוף | , |
| 5:3.1329 | והרחיקו |  |
| 5:3.1330 | ממנו |  |
| 5:3.1331 | מעשות |  |
| 5:3.1332 | כל |  |
| 5:3.1333 | רע | , |
| 5:3.1334 | וקראוהו |  |
| 5:3.1335 | יזדאן | ". |
| 5:3.1336 | וייחסו |  |
| 5:3.1337 | כל |  |
| 5:3.1338 | הרע |  |
| 5:3.1339 | והצער |  |
| 5:3.1340 | ההווים |  |
| 5:3.1341 | בעולם |  |
| 5:3.1342 | לשטן | , |
| 5:3.1343 | וקראוהו |  |
| 5:3.1344 | אהרמן | " |
| 5:3.1345 | בלשונם | . |
| 5:3.1346 | והמה |  |
| 5:3.1347 | חלוקים | : |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

