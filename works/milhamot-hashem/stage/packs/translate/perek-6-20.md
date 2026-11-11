# Translation pack: milhamot-hashem perek-6-20 (Milchamot ha-Shem 6:30–6:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-20`.

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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- An initial before a surname that fixes the person (ר"מ ריקנטי): expand and
  name him in a note ("Rabbi Menachem Rikanti"); an initial no surname fixes
  stays an initial with a note, as with R. A. and ר"י חיים ("Rabbi Y. Chayyim").
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 6:30

Hebrew (Wikisource): בהא סלקינן ובהא נחתינן. כיון שהמסית מחבר הזהר בא להסתינו ולהדיחנו מעל ה' אלוהינו. ופגע בכבוד אבינו מלכינו וביסודי תורתינו הקדושה ושרשיה הטהורים ביחוד השי"ת שהוא הסבה הראשונה. ובא ליסד לנו דת ותורה חדשה להאמין בסבות רבות שהם משתלשלות זו מזו. וכל אחת נקראת בשם הוי'ה ואדנות ואלהים והקב"ה ויחס התפלה והעבודה והקרבנות שמקריבין ישראל בזמן שבית המקדש קיים לנברא קצר אפים (זעיר אנפין) והוא הסבה האחרונה שבפרצופי האצילות. אע"פ שאמר שנגלה לו עתיקא קדישא או אליהו ור"מ וכיוצא לא נאמין לו, כי שקר בפיו. שכן קבלנו מרז"ל שאין נביא רשאי לחדש דבר מעתה. ואף אם בת קול תכריז מן השמים לאמר שמעו אליו, לא נשמע לו לסור מן הדרך אשר צוה ה' אלהינו לזלזל בשום עיקר מעיקרי תורתינו. כל שכן זה שהוא עוקר ארבעה עיקרים מיסודי תוה"ק שמנאם הרמב"ם ז"ל בפירוש המשנה, כמו שנתבאר לעיל.

| id | word | punctuation after |
|---|---|---|
| 6:30.1 | בהא |  |
| 6:30.2 | סלקינן |  |
| 6:30.3 | ובהא |  |
| 6:30.4 | נחתינן | . |
| 6:30.5 | כיון |  |
| 6:30.6 | שהמסית |  |
| 6:30.7 | מחבר |  |
| 6:30.8 | הזהר |  |
| 6:30.9 | בא |  |
| 6:30.10 | להסתינו |  |
| 6:30.11 | ולהדיחנו |  |
| 6:30.12 | מעל |  |
| 6:30.13 | ה | ' |
| 6:30.14 | אלוהינו | . |
| 6:30.15 | ופגע |  |
| 6:30.16 | בכבוד |  |
| 6:30.17 | אבינו |  |
| 6:30.18 | מלכינו |  |
| 6:30.19 | וביסודי |  |
| 6:30.20 | תורתינו |  |
| 6:30.21 | הקדושה |  |
| 6:30.22 | ושרשיה |  |
| 6:30.23 | הטהורים |  |
| 6:30.24 | ביחוד |  |
| 6:30.25 | השי"ת |  |
| 6:30.26 | שהוא |  |
| 6:30.27 | הסבה |  |
| 6:30.28 | הראשונה | . |
| 6:30.29 | ובא |  |
| 6:30.30 | ליסד |  |
| 6:30.31 | לנו |  |
| 6:30.32 | דת |  |
| 6:30.33 | ותורה |  |
| 6:30.34 | חדשה |  |
| 6:30.35 | להאמין |  |
| 6:30.36 | בסבות |  |
| 6:30.37 | רבות |  |
| 6:30.38 | שהם |  |
| 6:30.39 | משתלשלות |  |
| 6:30.40 | זו |  |
| 6:30.41 | מזו | . |
| 6:30.42 | וכל |  |
| 6:30.43 | אחת |  |
| 6:30.44 | נקראת |  |
| 6:30.45 | בשם |  |
| 6:30.46 | הוי'ה |  |
| 6:30.47 | ואדנות |  |
| 6:30.48 | ואלהים |  |
| 6:30.49 | והקב"ה |  |
| 6:30.50 | ויחס |  |
| 6:30.51 | התפלה |  |
| 6:30.52 | והעבודה |  |
| 6:30.53 | והקרבנות |  |
| 6:30.54 | שמקריבין |  |
| 6:30.55 | ישראל |  |
| 6:30.56 | בזמן |  |
| 6:30.57 | שבית |  |
| 6:30.58 | המקדש |  |
| 6:30.59 | קיים |  |
| 6:30.60 | לנברא |  |
| 6:30.61 | קצר |  |
| 6:30.62 | אפים |  |
| 6:30.63 | זעיר |  |
| 6:30.64 | אנפין | ) |
| 6:30.65 | והוא |  |
| 6:30.66 | הסבה |  |
| 6:30.67 | האחרונה |  |
| 6:30.68 | שבפרצופי |  |
| 6:30.69 | האצילות | . |
| 6:30.70 | אע"פ |  |
| 6:30.71 | שאמר |  |
| 6:30.72 | שנגלה |  |
| 6:30.73 | לו |  |
| 6:30.74 | עתיקא |  |
| 6:30.75 | קדישא |  |
| 6:30.76 | או |  |
| 6:30.77 | אליהו |  |
| 6:30.78 | ור"מ |  |
| 6:30.79 | וכיוצא |  |
| 6:30.80 | לא |  |
| 6:30.81 | נאמין |  |
| 6:30.82 | לו | , |
| 6:30.83 | כי |  |
| 6:30.84 | שקר |  |
| 6:30.85 | בפיו | . |
| 6:30.86 | שכן |  |
| 6:30.87 | קבלנו |  |
| 6:30.88 | מרז"ל |  |
| 6:30.89 | שאין |  |
| 6:30.90 | נביא |  |
| 6:30.91 | רשאי |  |
| 6:30.92 | לחדש |  |
| 6:30.93 | דבר |  |
| 6:30.94 | מעתה | . |
| 6:30.95 | ואף |  |
| 6:30.96 | אם |  |
| 6:30.97 | בת |  |
| 6:30.98 | קול |  |
| 6:30.99 | תכריז |  |
| 6:30.100 | מן |  |
| 6:30.101 | השמים |  |
| 6:30.102 | לאמר |  |
| 6:30.103 | שמעו |  |
| 6:30.104 | אליו | , |
| 6:30.105 | לא |  |
| 6:30.106 | נשמע |  |
| 6:30.107 | לו |  |
| 6:30.108 | לסור |  |
| 6:30.109 | מן |  |
| 6:30.110 | הדרך |  |
| 6:30.111 | אשר |  |
| 6:30.112 | צוה |  |
| 6:30.113 | ה | ' |
| 6:30.114 | אלהינו |  |
| 6:30.115 | לזלזל |  |
| 6:30.116 | בשום |  |
| 6:30.117 | עיקר |  |
| 6:30.118 | מעיקרי |  |
| 6:30.119 | תורתינו | . |
| 6:30.120 | כל |  |
| 6:30.121 | שכן |  |
| 6:30.122 | זה |  |
| 6:30.123 | שהוא |  |
| 6:30.124 | עוקר |  |
| 6:30.125 | ארבעה |  |
| 6:30.126 | עיקרים |  |
| 6:30.127 | מיסודי |  |
| 6:30.128 | תוה"ק |  |
| 6:30.129 | שמנאם |  |
| 6:30.130 | הרמב"ם |  |
| 6:30.131 | ז"ל |  |
| 6:30.132 | בפירוש |  |
| 6:30.133 | המשנה | , |
| 6:30.134 | כמו |  |
| 6:30.135 | שנתבאר |  |
| 6:30.136 | לעיל | . |


## Milchamot ha-Shem 6:31

Hebrew (Wikisource): וכתב בספר שבילי אמונה תחלת הנתיב הראשון וז"ל דרך אמונה בחרתי משפטיך שויתי וכו' וכתבו חכמי המחקר, לא יוכל לעבוד עלת כל העילות אלא נביא הדור בטבעו או החכם המובהק במה שקנה מן החכמה. וכאשר ישיג האדם לזה, אזי יגיע לקבל שכר הטוב ויתקיים בו כוונת הבריאה כאמרו כל הנקרא בשמי ולכבודי בראתיו. ולפי עיקר זה, נאמר בראשית עשרת הדברות אנכי ה' אלהיך כלומר האמן בלבך אמתת מציאותי. ואח"כ תוכל לקבל עול מלכותי. לפי שאין קבלת עול מלכות ראויה. רק אחר אמונת מציאותו. ואמר ה' אלהיך כלומר הוה קדמון. ומאתו היה הכל בחפץ ויכולת. והוא אלהים להם, שחייבים לעבוד אותו. ואמר אשר הוצאתיך מארץ מצרים, כי הוצאתו אותנו משם. מורה על המציאות והחפץ. כי בידיעה ובהשגחה ממנו ית' יצאו. וגם תורה על החדוש. כי אם העולם קדמון לא ישתנה דבר מטבעו. ותורה על היכולת, ותורה על היחוד. כמו שאמר בעבור תדע כי אין כמוני בכל הארץ ומצוה זו לא נאמרה בדברות על דרך הצווי, לפי שאין ראוי לצוות אלא אחר שידע השמוע מי הוא המצווה ויאמין בלבו שהוא נמצא על ענין כך. ולפיכך אמר בדבור זה אנכי ה' אלהיך, ולא אמר על צד הצווי אלא על צד ההאמנה. ואחר שהודיענו אמתת מציאותו ודרכי הנהגתו לעולם, אזי התחיל לצוות אותנו ולהודיע לנו שאין דבר זולתו ראוי לעבודה. לפי שכל הנמצאים זולתו נמצאו במאמרו, ואין ראוי לכלול עמו ובקבלת עול מלכותו אחד מכל העבדים. שלא להשוות כבוד הבורא יתברך לכבוד הנבראים. לפיכך סמך מצוה זו לאמתת מציאותו. לפי שאמונת עבודת זולתו מכחשת מציאותו ועל דבר זה התחיל לצוות ולהודיע שאין ראוי לעבוד לזולתו. ולכן הזכיר ואמר לא יהיה לך אלהים אחרים על פני. ועל יעלה בדעתך ומחשבתך ממה שאמר לא יהיה לך אלהים אחרים שיש מציאות אלוה אחר זולתו. כי מה שאמר אלהים הוא לפי מחשבת עובדיהם. כמו והאנשים רדפו אחריהם. ואמר אחרים, כלומר אלהים נמשכים אחר אחרים, שאין להם כח לעצמן, שקיומן תלוי בדבר אחר זולתן. ומי שאין לו כח קיים מעצמו אלא מזולתו אין ראוי להיעבד. ולא להשבע בשמו. אבל העובדה הראויה היא ליחיד אשר הוא נושא ומקיים כל הבריות הנמשכים זה אחר זה, ועדיו ית' יגיע תכליתן והוא אין לו נושא ואין מקיים ית' שמו ויתעלה זכרו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:31.1 | וכתב |  |
| 6:31.2 | בספר |  |
| 6:31.3 | שבילי |  |
| 6:31.4 | אמונה |  |
| 6:31.5 | תחלת |  |
| 6:31.6 | הנתיב |  |
| 6:31.7 | הראשון |  |
| 6:31.8 | וז"ל |  |
| 6:31.9 | דרך |  |
| 6:31.10 | אמונה |  |
| 6:31.11 | בחרתי |  |
| 6:31.12 | משפטיך |  |
| 6:31.13 | שויתי |  |
| 6:31.14 | וכו | ' |
| 6:31.15 | וכתבו |  |
| 6:31.16 | חכמי |  |
| 6:31.17 | המחקר | , |
| 6:31.18 | לא |  |
| 6:31.19 | יוכל |  |
| 6:31.20 | לעבוד |  |
| 6:31.21 | עלת |  |
| 6:31.22 | כל |  |
| 6:31.23 | העילות |  |
| 6:31.24 | אלא |  |
| 6:31.25 | נביא |  |
| 6:31.26 | הדור |  |
| 6:31.27 | בטבעו |  |
| 6:31.28 | או |  |
| 6:31.29 | החכם |  |
| 6:31.30 | המובהק |  |
| 6:31.31 | במה |  |
| 6:31.32 | שקנה |  |
| 6:31.33 | מן |  |
| 6:31.34 | החכמה | . |
| 6:31.35 | וכאשר |  |
| 6:31.36 | ישיג |  |
| 6:31.37 | האדם |  |
| 6:31.38 | לזה | , |
| 6:31.39 | אזי |  |
| 6:31.40 | יגיע |  |
| 6:31.41 | לקבל |  |
| 6:31.42 | שכר |  |
| 6:31.43 | הטוב |  |
| 6:31.44 | ויתקיים |  |
| 6:31.45 | בו |  |
| 6:31.46 | כוונת |  |
| 6:31.47 | הבריאה |  |
| 6:31.48 | כאמרו |  |
| 6:31.49 | כל |  |
| 6:31.50 | הנקרא |  |
| 6:31.51 | בשמי |  |
| 6:31.52 | ולכבודי |  |
| 6:31.53 | בראתיו | . |
| 6:31.54 | ולפי |  |
| 6:31.55 | עיקר |  |
| 6:31.56 | זה | , |
| 6:31.57 | נאמר |  |
| 6:31.58 | בראשית |  |
| 6:31.59 | עשרת |  |
| 6:31.60 | הדברות |  |
| 6:31.61 | אנכי |  |
| 6:31.62 | ה | ' |
| 6:31.63 | אלהיך |  |
| 6:31.64 | כלומר |  |
| 6:31.65 | האמן |  |
| 6:31.66 | בלבך |  |
| 6:31.67 | אמתת |  |
| 6:31.68 | מציאותי | . |
| 6:31.69 | ואח"כ |  |
| 6:31.70 | תוכל |  |
| 6:31.71 | לקבל |  |
| 6:31.72 | עול |  |
| 6:31.73 | מלכותי | . |
| 6:31.74 | לפי |  |
| 6:31.75 | שאין |  |
| 6:31.76 | קבלת |  |
| 6:31.77 | עול |  |
| 6:31.78 | מלכות |  |
| 6:31.79 | ראויה | . |
| 6:31.80 | רק |  |
| 6:31.81 | אחר |  |
| 6:31.82 | אמונת |  |
| 6:31.83 | מציאותו | . |
| 6:31.84 | ואמר |  |
| 6:31.85 | ה | ' |
| 6:31.86 | אלהיך |  |
| 6:31.87 | כלומר |  |
| 6:31.88 | הוה |  |
| 6:31.89 | קדמון | . |
| 6:31.90 | ומאתו |  |
| 6:31.91 | היה |  |
| 6:31.92 | הכל |  |
| 6:31.93 | בחפץ |  |
| 6:31.94 | ויכולת | . |
| 6:31.95 | והוא |  |
| 6:31.96 | אלהים |  |
| 6:31.97 | להם | , |
| 6:31.98 | שחייבים |  |
| 6:31.99 | לעבוד |  |
| 6:31.100 | אותו | . |
| 6:31.101 | ואמר |  |
| 6:31.102 | אשר |  |
| 6:31.103 | הוצאתיך |  |
| 6:31.104 | מארץ |  |
| 6:31.105 | מצרים | , |
| 6:31.106 | כי |  |
| 6:31.107 | הוצאתו |  |
| 6:31.108 | אותנו |  |
| 6:31.109 | משם | . |
| 6:31.110 | מורה |  |
| 6:31.111 | על |  |
| 6:31.112 | המציאות |  |
| 6:31.113 | והחפץ | . |
| 6:31.114 | כי |  |
| 6:31.115 | בידיעה |  |
| 6:31.116 | ובהשגחה |  |
| 6:31.117 | ממנו |  |
| 6:31.118 | ית | ' |
| 6:31.119 | יצאו | . |
| 6:31.120 | וגם |  |
| 6:31.121 | תורה |  |
| 6:31.122 | על |  |
| 6:31.123 | החדוש | . |
| 6:31.124 | כי |  |
| 6:31.125 | אם |  |
| 6:31.126 | העולם |  |
| 6:31.127 | קדמון |  |
| 6:31.128 | לא |  |
| 6:31.129 | ישתנה |  |
| 6:31.130 | דבר |  |
| 6:31.131 | מטבעו | . |
| 6:31.132 | ותורה |  |
| 6:31.133 | על |  |
| 6:31.134 | היכולת | , |
| 6:31.135 | ותורה |  |
| 6:31.136 | על |  |
| 6:31.137 | היחוד | . |
| 6:31.138 | כמו |  |
| 6:31.139 | שאמר |  |
| 6:31.140 | בעבור |  |
| 6:31.141 | תדע |  |
| 6:31.142 | כי |  |
| 6:31.143 | אין |  |
| 6:31.144 | כמוני |  |
| 6:31.145 | בכל |  |
| 6:31.146 | הארץ |  |
| 6:31.147 | ומצוה |  |
| 6:31.148 | זו |  |
| 6:31.149 | לא |  |
| 6:31.150 | נאמרה |  |
| 6:31.151 | בדברות |  |
| 6:31.152 | על |  |
| 6:31.153 | דרך |  |
| 6:31.154 | הצווי | , |
| 6:31.155 | לפי |  |
| 6:31.156 | שאין |  |
| 6:31.157 | ראוי |  |
| 6:31.158 | לצוות |  |
| 6:31.159 | אלא |  |
| 6:31.160 | אחר |  |
| 6:31.161 | שידע |  |
| 6:31.162 | השמוע |  |
| 6:31.163 | מי |  |
| 6:31.164 | הוא |  |
| 6:31.165 | המצווה |  |
| 6:31.166 | ויאמין |  |
| 6:31.167 | בלבו |  |
| 6:31.168 | שהוא |  |
| 6:31.169 | נמצא |  |
| 6:31.170 | על |  |
| 6:31.171 | ענין |  |
| 6:31.172 | כך | . |
| 6:31.173 | ולפיכך |  |
| 6:31.174 | אמר |  |
| 6:31.175 | בדבור |  |
| 6:31.176 | זה |  |
| 6:31.177 | אנכי |  |
| 6:31.178 | ה | ' |
| 6:31.179 | אלהיך | , |
| 6:31.180 | ולא |  |
| 6:31.181 | אמר |  |
| 6:31.182 | על |  |
| 6:31.183 | צד |  |
| 6:31.184 | הצווי |  |
| 6:31.185 | אלא |  |
| 6:31.186 | על |  |
| 6:31.187 | צד |  |
| 6:31.188 | ההאמנה | . |
| 6:31.189 | ואחר |  |
| 6:31.190 | שהודיענו |  |
| 6:31.191 | אמתת |  |
| 6:31.192 | מציאותו |  |
| 6:31.193 | ודרכי |  |
| 6:31.194 | הנהגתו |  |
| 6:31.195 | לעולם | , |
| 6:31.196 | אזי |  |
| 6:31.197 | התחיל |  |
| 6:31.198 | לצוות |  |
| 6:31.199 | אותנו |  |
| 6:31.200 | ולהודיע |  |
| 6:31.201 | לנו |  |
| 6:31.202 | שאין |  |
| 6:31.203 | דבר |  |
| 6:31.204 | זולתו |  |
| 6:31.205 | ראוי |  |
| 6:31.206 | לעבודה | . |
| 6:31.207 | לפי |  |
| 6:31.208 | שכל |  |
| 6:31.209 | הנמצאים |  |
| 6:31.210 | זולתו |  |
| 6:31.211 | נמצאו |  |
| 6:31.212 | במאמרו | , |
| 6:31.213 | ואין |  |
| 6:31.214 | ראוי |  |
| 6:31.215 | לכלול |  |
| 6:31.216 | עמו |  |
| 6:31.217 | ובקבלת |  |
| 6:31.218 | עול |  |
| 6:31.219 | מלכותו |  |
| 6:31.220 | אחד |  |
| 6:31.221 | מכל |  |
| 6:31.222 | העבדים | . |
| 6:31.223 | שלא |  |
| 6:31.224 | להשוות |  |
| 6:31.225 | כבוד |  |
| 6:31.226 | הבורא |  |
| 6:31.227 | יתברך |  |
| 6:31.228 | לכבוד |  |
| 6:31.229 | הנבראים | . |
| 6:31.230 | לפיכך |  |
| 6:31.231 | סמך |  |
| 6:31.232 | מצוה |  |
| 6:31.233 | זו |  |
| 6:31.234 | לאמתת |  |
| 6:31.235 | מציאותו | . |
| 6:31.236 | לפי |  |
| 6:31.237 | שאמונת |  |
| 6:31.238 | עבודת |  |
| 6:31.239 | זולתו |  |
| 6:31.240 | מכחשת |  |
| 6:31.241 | מציאותו |  |
| 6:31.242 | ועל |  |
| 6:31.243 | דבר |  |
| 6:31.244 | זה |  |
| 6:31.245 | התחיל |  |
| 6:31.246 | לצוות |  |
| 6:31.247 | ולהודיע |  |
| 6:31.248 | שאין |  |
| 6:31.249 | ראוי |  |
| 6:31.250 | לעבוד |  |
| 6:31.251 | לזולתו | . |
| 6:31.252 | ולכן |  |
| 6:31.253 | הזכיר |  |
| 6:31.254 | ואמר |  |
| 6:31.255 | לא |  |
| 6:31.256 | יהיה |  |
| 6:31.257 | לך |  |
| 6:31.258 | אלהים |  |
| 6:31.259 | אחרים |  |
| 6:31.260 | על |  |
| 6:31.261 | פני | . |
| 6:31.262 | ועל |  |
| 6:31.263 | יעלה |  |
| 6:31.264 | בדעתך |  |
| 6:31.265 | ומחשבתך |  |
| 6:31.266 | ממה |  |
| 6:31.267 | שאמר |  |
| 6:31.268 | לא |  |
| 6:31.269 | יהיה |  |
| 6:31.270 | לך |  |
| 6:31.271 | אלהים |  |
| 6:31.272 | אחרים |  |
| 6:31.273 | שיש |  |
| 6:31.274 | מציאות |  |
| 6:31.275 | אלוה |  |
| 6:31.276 | אחר |  |
| 6:31.277 | זולתו | . |
| 6:31.278 | כי |  |
| 6:31.279 | מה |  |
| 6:31.280 | שאמר |  |
| 6:31.281 | אלהים |  |
| 6:31.282 | הוא |  |
| 6:31.283 | לפי |  |
| 6:31.284 | מחשבת |  |
| 6:31.285 | עובדיהם | . |
| 6:31.286 | כמו |  |
| 6:31.287 | והאנשים |  |
| 6:31.288 | רדפו |  |
| 6:31.289 | אחריהם | . |
| 6:31.290 | ואמר |  |
| 6:31.291 | אחרים | , |
| 6:31.292 | כלומר |  |
| 6:31.293 | אלהים |  |
| 6:31.294 | נמשכים |  |
| 6:31.295 | אחר |  |
| 6:31.296 | אחרים | , |
| 6:31.297 | שאין |  |
| 6:31.298 | להם |  |
| 6:31.299 | כח |  |
| 6:31.300 | לעצמן | , |
| 6:31.301 | שקיומן |  |
| 6:31.302 | תלוי |  |
| 6:31.303 | בדבר |  |
| 6:31.304 | אחר |  |
| 6:31.305 | זולתן | . |
| 6:31.306 | ומי |  |
| 6:31.307 | שאין |  |
| 6:31.308 | לו |  |
| 6:31.309 | כח |  |
| 6:31.310 | קיים |  |
| 6:31.311 | מעצמו |  |
| 6:31.312 | אלא |  |
| 6:31.313 | מזולתו |  |
| 6:31.314 | אין |  |
| 6:31.315 | ראוי |  |
| 6:31.316 | להיעבד | . |
| 6:31.317 | ולא |  |
| 6:31.318 | להשבע |  |
| 6:31.319 | בשמו | . |
| 6:31.320 | אבל |  |
| 6:31.321 | העובדה |  |
| 6:31.322 | הראויה |  |
| 6:31.323 | היא |  |
| 6:31.324 | ליחיד |  |
| 6:31.325 | אשר |  |
| 6:31.326 | הוא |  |
| 6:31.327 | נושא |  |
| 6:31.328 | ומקיים |  |
| 6:31.329 | כל |  |
| 6:31.330 | הבריות |  |
| 6:31.331 | הנמשכים |  |
| 6:31.332 | זה |  |
| 6:31.333 | אחר |  |
| 6:31.334 | זה | , |
| 6:31.335 | ועדיו |  |
| 6:31.336 | ית | ' |
| 6:31.337 | יגיע |  |
| 6:31.338 | תכליתן |  |
| 6:31.339 | והוא |  |
| 6:31.340 | אין |  |
| 6:31.341 | לו |  |
| 6:31.342 | נושא |  |
| 6:31.343 | ואין |  |
| 6:31.344 | מקיים |  |
| 6:31.345 | ית | ' |
| 6:31.346 | שמו |  |
| 6:31.347 | ויתעלה |  |
| 6:31.348 | זכרו |  |
| 6:31.349 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

