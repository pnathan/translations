# Translation pack: milhamot-hashem perek-6-05 (Milchamot ha-Shem 6:10–6:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-05`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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


## Milchamot ha-Shem 6:10

Hebrew (Wikisource): ובשיר השירים רבה פר' א' איתה, אמר ריש לקיש ולמה הוא אמת. אלף בראש האותיות, מ"ם באצמע, תיו בסופן. לומר אני ראשון ואני אחרון ומבלעדי אין אלהים. אני ראשון שלא קיבלתי מלכותי מאחר, ואני אחרון שאיני מוסרה לאחר שאינו בעולם. ומבלעדי אין אלהים שאין לי שני ע"כ. ופי' הר' עץ יוסף ששלשה שרשים מיוחדים באלהותו ית' הא' בקדמותו, והשני אחדותו, והג' נצחיותו ע"כ. הרי נתבאר לך מזה שהירושלמי ושמות רבה ושיר השירים רבה כולם מסכימים לעיקרים שכתב הרמב"ם באחדותו ית' וקדמותו. ודברי הזהר הומקובלים החדשים סותרים העיקרים ההם. לפי שנמשכו אחר מאמיני השלוש לעבוד את הבן כמו שנתבאר. ורז"ל הפליגו בהרחקת בעלי אמונה זו עד שאסרו לישא וליתן עמהם כדגרסינן בע"ז פרק אין מעמידין לא ישא ויתן עם המינים ואין מתרפאין מהן אפי' לחיי שעה. ומעשה בבן דמא בן אחותו של ר' ישמעאל שהכישו נחש ובא יעקב איש כפר סכניא לרפאתו ולא הניחו ר' ישמעאל. אמר לו ישמעאל אחי, הנח לי ואתרפא ממנו. ואני אביא לך מקרא מן התורה שהוא מותר. ולא הספיק לגמור את הדבר עד שיצאה נשמתו ומת. קרי עליה ר' ישמעאל אשריך בן דמא שגופך טהור ונשמתך יצאה בטהרה, ולא עברת על דברי חביריך שהיו אומרים ופורץ גדר ישכנו נחש, ומסיים שם שאני מינות משאר איסורין של תורה. משום דמשכא ואתי למשך בתריהו ע"כ. טהכי נמי אמרו בירוש' שבת פרק ח' שרצים. בר ברוה דר' יהושע בן לוי הוה ליה בלי את בר נש ולחש ליה מן שמיה דישו פנדירא ואינשם כד נפיק אמר ליה מאן לחשת ליה, אמר ליה מילת פלן. אמר ליה נוח הוה ליה אלו הוה מאית. והות לי' כן. בשגגה שיוצא מלפני השליט. עוד שם מעשה בר' אלעזר בן דמא שנשכו נחש ובא יעקב איש כפר סמא (הוא יעקב איש כפר סכניא הנז' בבבלי שהוא מתלמידי יש"ו) משם של יש"ו פנדריא לרפותו ולא הניח לו ר' ישמעאל. אמר לו אני מביא לך ראיה שירפאני. לא הספיק להביא ראיה עד שמת בן דמא. אמר לו ר' ישמעאל אשריך בן דמא שיצאת בשלום מן העולם ולא פרצת גדירן של חכמים דכ' ופורץ ישכנו נחש וכו' ומאי הוה ליה מימר אשר יעשה אותם האדם וחי בהם ע"כ. מכל זה אתה למד כמה הפליגו והרחיקו שלא לישא וליתן עם מאמיני השלשה ואין מתרפאין מהן אפי' במקום סכנה. ואע"פ שהם אומרים שהשלשה אחד, ולפי שכל ישראל בזמן רז"ל אדוקים ודבוקים ביחוד השי"ת באמת לא הוצרכו במשנה ובתלמוד לבאר דבר זה ולומר שמי שחושב עלות נאצלים מהבורא ואומר שהם אחד עם מאצילם שהוא עובד ע"ז. כי זו היא אמונת המשיחים. ועליה אמר ר' שמעון ב"י כל המשתף שם שמים ודבר אחר נעקר מן העולם. וכבר אמרו הנצרים עובדי ע"ז הם. ויום ראשון יום אידם הוא. לפיכך אסור לשאת ולתת עמהן בארץ ישראל יום ה' ויום ו' שבכל שבת ושבת ואין צריך לומר יום ראשון עצמו שהוא אסור בכל מקום. וכן בכל אידיהן (הרמב"ם פ"ט מע"ז) ולכן כ' הסמ"ג שאין צריך להאריך בזה מפני שכל ישראל אדוקים באמונת היחוד.

| id | word | punctuation after |
|---|---|---|
| 6:10.1 | ובשיר |  |
| 6:10.2 | השירים |  |
| 6:10.3 | רבה |  |
| 6:10.4 | פר | ' |
| 6:10.5 | א | ' |
| 6:10.6 | איתה | , |
| 6:10.7 | אמר |  |
| 6:10.8 | ריש |  |
| 6:10.9 | לקיש |  |
| 6:10.10 | ולמה |  |
| 6:10.11 | הוא |  |
| 6:10.12 | אמת | . |
| 6:10.13 | אלף |  |
| 6:10.14 | בראש |  |
| 6:10.15 | האותיות | , |
| 6:10.16 | מ"ם |  |
| 6:10.17 | באצמע | , |
| 6:10.18 | תיו |  |
| 6:10.19 | בסופן | . |
| 6:10.20 | לומר |  |
| 6:10.21 | אני |  |
| 6:10.22 | ראשון |  |
| 6:10.23 | ואני |  |
| 6:10.24 | אחרון |  |
| 6:10.25 | ומבלעדי |  |
| 6:10.26 | אין |  |
| 6:10.27 | אלהים | . |
| 6:10.28 | אני |  |
| 6:10.29 | ראשון |  |
| 6:10.30 | שלא |  |
| 6:10.31 | קיבלתי |  |
| 6:10.32 | מלכותי |  |
| 6:10.33 | מאחר | , |
| 6:10.34 | ואני |  |
| 6:10.35 | אחרון |  |
| 6:10.36 | שאיני |  |
| 6:10.37 | מוסרה |  |
| 6:10.38 | לאחר |  |
| 6:10.39 | שאינו |  |
| 6:10.40 | בעולם | . |
| 6:10.41 | ומבלעדי |  |
| 6:10.42 | אין |  |
| 6:10.43 | אלהים |  |
| 6:10.44 | שאין |  |
| 6:10.45 | לי |  |
| 6:10.46 | שני |  |
| 6:10.47 | ע"כ | . |
| 6:10.48 | ופי | ' |
| 6:10.49 | הר | ' |
| 6:10.50 | עץ |  |
| 6:10.51 | יוסף |  |
| 6:10.52 | ששלשה |  |
| 6:10.53 | שרשים |  |
| 6:10.54 | מיוחדים |  |
| 6:10.55 | באלהותו |  |
| 6:10.56 | ית | ' |
| 6:10.57 | הא | ' |
| 6:10.58 | בקדמותו | , |
| 6:10.59 | והשני |  |
| 6:10.60 | אחדותו | , |
| 6:10.61 | והג | ' |
| 6:10.62 | נצחיותו |  |
| 6:10.63 | ע"כ | . |
| 6:10.64 | הרי |  |
| 6:10.65 | נתבאר |  |
| 6:10.66 | לך |  |
| 6:10.67 | מזה |  |
| 6:10.68 | שהירושלמי |  |
| 6:10.69 | ושמות |  |
| 6:10.70 | רבה |  |
| 6:10.71 | ושיר |  |
| 6:10.72 | השירים |  |
| 6:10.73 | רבה |  |
| 6:10.74 | כולם |  |
| 6:10.75 | מסכימים |  |
| 6:10.76 | לעיקרים |  |
| 6:10.77 | שכתב |  |
| 6:10.78 | הרמב"ם |  |
| 6:10.79 | באחדותו |  |
| 6:10.80 | ית | ' |
| 6:10.81 | וקדמותו | . |
| 6:10.82 | ודברי |  |
| 6:10.83 | הזהר |  |
| 6:10.84 | הומקובלים |  |
| 6:10.85 | החדשים |  |
| 6:10.86 | סותרים |  |
| 6:10.87 | העיקרים |  |
| 6:10.88 | ההם | . |
| 6:10.89 | לפי |  |
| 6:10.90 | שנמשכו |  |
| 6:10.91 | אחר |  |
| 6:10.92 | מאמיני |  |
| 6:10.93 | השלוש |  |
| 6:10.94 | לעבוד |  |
| 6:10.95 | את |  |
| 6:10.96 | הבן |  |
| 6:10.97 | כמו |  |
| 6:10.98 | שנתבאר | . |
| 6:10.99 | ורז"ל |  |
| 6:10.100 | הפליגו |  |
| 6:10.101 | בהרחקת |  |
| 6:10.102 | בעלי |  |
| 6:10.103 | אמונה |  |
| 6:10.104 | זו |  |
| 6:10.105 | עד |  |
| 6:10.106 | שאסרו |  |
| 6:10.107 | לישא |  |
| 6:10.108 | וליתן |  |
| 6:10.109 | עמהם |  |
| 6:10.110 | כדגרסינן |  |
| 6:10.111 | בע"ז |  |
| 6:10.112 | פרק |  |
| 6:10.113 | אין |  |
| 6:10.114 | מעמידין |  |
| 6:10.115 | לא |  |
| 6:10.116 | ישא |  |
| 6:10.117 | ויתן |  |
| 6:10.118 | עם |  |
| 6:10.119 | המינים |  |
| 6:10.120 | ואין |  |
| 6:10.121 | מתרפאין |  |
| 6:10.122 | מהן |  |
| 6:10.123 | אפי | ' |
| 6:10.124 | לחיי |  |
| 6:10.125 | שעה | . |
| 6:10.126 | ומעשה |  |
| 6:10.127 | בבן |  |
| 6:10.128 | דמא |  |
| 6:10.129 | בן |  |
| 6:10.130 | אחותו |  |
| 6:10.131 | של |  |
| 6:10.132 | ר | ' |
| 6:10.133 | ישמעאל |  |
| 6:10.134 | שהכישו |  |
| 6:10.135 | נחש |  |
| 6:10.136 | ובא |  |
| 6:10.137 | יעקב |  |
| 6:10.138 | איש |  |
| 6:10.139 | כפר |  |
| 6:10.140 | סכניא |  |
| 6:10.141 | לרפאתו |  |
| 6:10.142 | ולא |  |
| 6:10.143 | הניחו |  |
| 6:10.144 | ר | ' |
| 6:10.145 | ישמעאל | . |
| 6:10.146 | אמר |  |
| 6:10.147 | לו |  |
| 6:10.148 | ישמעאל |  |
| 6:10.149 | אחי | , |
| 6:10.150 | הנח |  |
| 6:10.151 | לי |  |
| 6:10.152 | ואתרפא |  |
| 6:10.153 | ממנו | . |
| 6:10.154 | ואני |  |
| 6:10.155 | אביא |  |
| 6:10.156 | לך |  |
| 6:10.157 | מקרא |  |
| 6:10.158 | מן |  |
| 6:10.159 | התורה |  |
| 6:10.160 | שהוא |  |
| 6:10.161 | מותר | . |
| 6:10.162 | ולא |  |
| 6:10.163 | הספיק |  |
| 6:10.164 | לגמור |  |
| 6:10.165 | את |  |
| 6:10.166 | הדבר |  |
| 6:10.167 | עד |  |
| 6:10.168 | שיצאה |  |
| 6:10.169 | נשמתו |  |
| 6:10.170 | ומת | . |
| 6:10.171 | קרי |  |
| 6:10.172 | עליה |  |
| 6:10.173 | ר | ' |
| 6:10.174 | ישמעאל |  |
| 6:10.175 | אשריך |  |
| 6:10.176 | בן |  |
| 6:10.177 | דמא |  |
| 6:10.178 | שגופך |  |
| 6:10.179 | טהור |  |
| 6:10.180 | ונשמתך |  |
| 6:10.181 | יצאה |  |
| 6:10.182 | בטהרה | , |
| 6:10.183 | ולא |  |
| 6:10.184 | עברת |  |
| 6:10.185 | על |  |
| 6:10.186 | דברי |  |
| 6:10.187 | חביריך |  |
| 6:10.188 | שהיו |  |
| 6:10.189 | אומרים |  |
| 6:10.190 | ופורץ |  |
| 6:10.191 | גדר |  |
| 6:10.192 | ישכנו |  |
| 6:10.193 | נחש | , |
| 6:10.194 | ומסיים |  |
| 6:10.195 | שם |  |
| 6:10.196 | שאני |  |
| 6:10.197 | מינות |  |
| 6:10.198 | משאר |  |
| 6:10.199 | איסורין |  |
| 6:10.200 | של |  |
| 6:10.201 | תורה | . |
| 6:10.202 | משום |  |
| 6:10.203 | דמשכא |  |
| 6:10.204 | ואתי |  |
| 6:10.205 | למשך |  |
| 6:10.206 | בתריהו |  |
| 6:10.207 | ע"כ | . |
| 6:10.208 | טהכי |  |
| 6:10.209 | נמי |  |
| 6:10.210 | אמרו |  |
| 6:10.211 | בירוש | ' |
| 6:10.212 | שבת |  |
| 6:10.213 | פרק |  |
| 6:10.214 | ח | ' |
| 6:10.215 | שרצים | . |
| 6:10.216 | בר |  |
| 6:10.217 | ברוה |  |
| 6:10.218 | דר | ' |
| 6:10.219 | יהושע |  |
| 6:10.220 | בן |  |
| 6:10.221 | לוי |  |
| 6:10.222 | הוה |  |
| 6:10.223 | ליה |  |
| 6:10.224 | בלי |  |
| 6:10.225 | את |  |
| 6:10.226 | בר |  |
| 6:10.227 | נש |  |
| 6:10.228 | ולחש |  |
| 6:10.229 | ליה |  |
| 6:10.230 | מן |  |
| 6:10.231 | שמיה |  |
| 6:10.232 | דישו |  |
| 6:10.233 | פנדירא |  |
| 6:10.234 | ואינשם |  |
| 6:10.235 | כד |  |
| 6:10.236 | נפיק |  |
| 6:10.237 | אמר |  |
| 6:10.238 | ליה |  |
| 6:10.239 | מאן |  |
| 6:10.240 | לחשת |  |
| 6:10.241 | ליה | , |
| 6:10.242 | אמר |  |
| 6:10.243 | ליה |  |
| 6:10.244 | מילת |  |
| 6:10.245 | פלן | . |
| 6:10.246 | אמר |  |
| 6:10.247 | ליה |  |
| 6:10.248 | נוח |  |
| 6:10.249 | הוה |  |
| 6:10.250 | ליה |  |
| 6:10.251 | אלו |  |
| 6:10.252 | הוה |  |
| 6:10.253 | מאית | . |
| 6:10.254 | והות |  |
| 6:10.255 | לי | ' |
| 6:10.256 | כן | . |
| 6:10.257 | בשגגה |  |
| 6:10.258 | שיוצא |  |
| 6:10.259 | מלפני |  |
| 6:10.260 | השליט | . |
| 6:10.261 | עוד |  |
| 6:10.262 | שם |  |
| 6:10.263 | מעשה |  |
| 6:10.264 | בר | ' |
| 6:10.265 | אלעזר |  |
| 6:10.266 | בן |  |
| 6:10.267 | דמא |  |
| 6:10.268 | שנשכו |  |
| 6:10.269 | נחש |  |
| 6:10.270 | ובא |  |
| 6:10.271 | יעקב |  |
| 6:10.272 | איש |  |
| 6:10.273 | כפר |  |
| 6:10.274 | סמא |  |
| 6:10.275 | הוא |  |
| 6:10.276 | יעקב |  |
| 6:10.277 | איש |  |
| 6:10.278 | כפר |  |
| 6:10.279 | סכניא |  |
| 6:10.280 | הנז | ' |
| 6:10.281 | בבבלי |  |
| 6:10.282 | שהוא |  |
| 6:10.283 | מתלמידי |  |
| 6:10.284 | יש"ו | ) |
| 6:10.285 | משם |  |
| 6:10.286 | של |  |
| 6:10.287 | יש"ו |  |
| 6:10.288 | פנדריא |  |
| 6:10.289 | לרפותו |  |
| 6:10.290 | ולא |  |
| 6:10.291 | הניח |  |
| 6:10.292 | לו |  |
| 6:10.293 | ר | ' |
| 6:10.294 | ישמעאל | . |
| 6:10.295 | אמר |  |
| 6:10.296 | לו |  |
| 6:10.297 | אני |  |
| 6:10.298 | מביא |  |
| 6:10.299 | לך |  |
| 6:10.300 | ראיה |  |
| 6:10.301 | שירפאני | . |
| 6:10.302 | לא |  |
| 6:10.303 | הספיק |  |
| 6:10.304 | להביא |  |
| 6:10.305 | ראיה |  |
| 6:10.306 | עד |  |
| 6:10.307 | שמת |  |
| 6:10.308 | בן |  |
| 6:10.309 | דמא | . |
| 6:10.310 | אמר |  |
| 6:10.311 | לו |  |
| 6:10.312 | ר | ' |
| 6:10.313 | ישמעאל |  |
| 6:10.314 | אשריך |  |
| 6:10.315 | בן |  |
| 6:10.316 | דמא |  |
| 6:10.317 | שיצאת |  |
| 6:10.318 | בשלום |  |
| 6:10.319 | מן |  |
| 6:10.320 | העולם |  |
| 6:10.321 | ולא |  |
| 6:10.322 | פרצת |  |
| 6:10.323 | גדירן |  |
| 6:10.324 | של |  |
| 6:10.325 | חכמים |  |
| 6:10.326 | דכ | ' |
| 6:10.327 | ופורץ |  |
| 6:10.328 | ישכנו |  |
| 6:10.329 | נחש |  |
| 6:10.330 | וכו | ' |
| 6:10.331 | ומאי |  |
| 6:10.332 | הוה |  |
| 6:10.333 | ליה |  |
| 6:10.334 | מימר |  |
| 6:10.335 | אשר |  |
| 6:10.336 | יעשה |  |
| 6:10.337 | אותם |  |
| 6:10.338 | האדם |  |
| 6:10.339 | וחי |  |
| 6:10.340 | בהם |  |
| 6:10.341 | ע"כ | . |
| 6:10.342 | מכל |  |
| 6:10.343 | זה |  |
| 6:10.344 | אתה |  |
| 6:10.345 | למד |  |
| 6:10.346 | כמה |  |
| 6:10.347 | הפליגו |  |
| 6:10.348 | והרחיקו |  |
| 6:10.349 | שלא |  |
| 6:10.350 | לישא |  |
| 6:10.351 | וליתן |  |
| 6:10.352 | עם |  |
| 6:10.353 | מאמיני |  |
| 6:10.354 | השלשה |  |
| 6:10.355 | ואין |  |
| 6:10.356 | מתרפאין |  |
| 6:10.357 | מהן |  |
| 6:10.358 | אפי | ' |
| 6:10.359 | במקום |  |
| 6:10.360 | סכנה | . |
| 6:10.361 | ואע"פ |  |
| 6:10.362 | שהם |  |
| 6:10.363 | אומרים |  |
| 6:10.364 | שהשלשה |  |
| 6:10.365 | אחד | , |
| 6:10.366 | ולפי |  |
| 6:10.367 | שכל |  |
| 6:10.368 | ישראל |  |
| 6:10.369 | בזמן |  |
| 6:10.370 | רז"ל |  |
| 6:10.371 | אדוקים |  |
| 6:10.372 | ודבוקים |  |
| 6:10.373 | ביחוד |  |
| 6:10.374 | השי"ת |  |
| 6:10.375 | באמת |  |
| 6:10.376 | לא |  |
| 6:10.377 | הוצרכו |  |
| 6:10.378 | במשנה |  |
| 6:10.379 | ובתלמוד |  |
| 6:10.380 | לבאר |  |
| 6:10.381 | דבר |  |
| 6:10.382 | זה |  |
| 6:10.383 | ולומר |  |
| 6:10.384 | שמי |  |
| 6:10.385 | שחושב |  |
| 6:10.386 | עלות |  |
| 6:10.387 | נאצלים |  |
| 6:10.388 | מהבורא |  |
| 6:10.389 | ואומר |  |
| 6:10.390 | שהם |  |
| 6:10.391 | אחד |  |
| 6:10.392 | עם |  |
| 6:10.393 | מאצילם |  |
| 6:10.394 | שהוא |  |
| 6:10.395 | עובד |  |
| 6:10.396 | ע"ז | . |
| 6:10.397 | כי |  |
| 6:10.398 | זו |  |
| 6:10.399 | היא |  |
| 6:10.400 | אמונת |  |
| 6:10.401 | המשיחים | . |
| 6:10.402 | ועליה |  |
| 6:10.403 | אמר |  |
| 6:10.404 | ר | ' |
| 6:10.405 | שמעון |  |
| 6:10.406 | ב"י |  |
| 6:10.407 | כל |  |
| 6:10.408 | המשתף |  |
| 6:10.409 | שם |  |
| 6:10.410 | שמים |  |
| 6:10.411 | ודבר |  |
| 6:10.412 | אחר |  |
| 6:10.413 | נעקר |  |
| 6:10.414 | מן |  |
| 6:10.415 | העולם | . |
| 6:10.416 | וכבר |  |
| 6:10.417 | אמרו |  |
| 6:10.418 | הנצרים |  |
| 6:10.419 | עובדי |  |
| 6:10.420 | ע"ז |  |
| 6:10.421 | הם | . |
| 6:10.422 | ויום |  |
| 6:10.423 | ראשון |  |
| 6:10.424 | יום |  |
| 6:10.425 | אידם |  |
| 6:10.426 | הוא | . |
| 6:10.427 | לפיכך |  |
| 6:10.428 | אסור |  |
| 6:10.429 | לשאת |  |
| 6:10.430 | ולתת |  |
| 6:10.431 | עמהן |  |
| 6:10.432 | בארץ |  |
| 6:10.433 | ישראל |  |
| 6:10.434 | יום |  |
| 6:10.435 | ה | ' |
| 6:10.436 | ויום |  |
| 6:10.437 | ו | ' |
| 6:10.438 | שבכל |  |
| 6:10.439 | שבת |  |
| 6:10.440 | ושבת |  |
| 6:10.441 | ואין |  |
| 6:10.442 | צריך |  |
| 6:10.443 | לומר |  |
| 6:10.444 | יום |  |
| 6:10.445 | ראשון |  |
| 6:10.446 | עצמו |  |
| 6:10.447 | שהוא |  |
| 6:10.448 | אסור |  |
| 6:10.449 | בכל |  |
| 6:10.450 | מקום | . |
| 6:10.451 | וכן |  |
| 6:10.452 | בכל |  |
| 6:10.453 | אידיהן |  |
| 6:10.454 | הרמב"ם |  |
| 6:10.455 | פ"ט |  |
| 6:10.456 | מע"ז | ) |
| 6:10.457 | ולכן |  |
| 6:10.458 | כ | ' |
| 6:10.459 | הסמ"ג |  |
| 6:10.460 | שאין |  |
| 6:10.461 | צריך |  |
| 6:10.462 | להאריך |  |
| 6:10.463 | בזה |  |
| 6:10.464 | מפני |  |
| 6:10.465 | שכל |  |
| 6:10.466 | ישראל |  |
| 6:10.467 | אדוקים |  |
| 6:10.468 | באמונת |  |
| 6:10.469 | היחוד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

