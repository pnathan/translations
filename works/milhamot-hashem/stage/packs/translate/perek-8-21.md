# Translation pack: milhamot-hashem perek-8-21 (Milchamot ha-Shem 8:25–8:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-21`.

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
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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


## Milchamot ha-Shem 8:25

Hebrew (Wikisource): קיח) ואם נפשך לומר לפי דבריהם שהאלהים העליון הוא נשמה של הפרצופים הנבראים ולנשמתם אנו עובדים ומתפללים, כי הוא מתפשט וממלא את הספירות וכמו שכתב הפילסוף המסית מחבר הזהר דף הנז' לעיל אנא הוא לבר מכלא ולגיו מכלא, ולכל סטרא, ולעילא מכלא, ולתתא מכלא, ולית אלה אחרא עילא ותתא ומכל סטרא, ומלגיו דעסר ספירן דמנהון כלא ובהון תליא כלא ואנת בכל ספירה בארכה ורחבה עילא ותתא. ובין כל ספירה וספירה, ואנת הוא דמקריב לקודשא בריך הוא ושכיניתיה בכל ספירה וספירה ובכל ענפין וכו'. ובדברים אלו לקח את לב קוראיו, וחשבו לתקן את הסתירה ודברי מינות הגלויה ומפורסמת שהאלהים העליון שאין כמהו בכל האלהים. איננו רק מקשר ומחבר ומזויג את הספירות הנקראים קודשא בריך הוא ושכינתיה אשר נעבדם ונשתחוה להם לבדם. אבל הוא לא נעבדנו כנזכר לעיל. והנה עם מה שהקדמתי מדברי רז"ל אין קיום ועמידה לדברים אלו כלל שכיון שקיימו וקבלו שיש לנשמה ציור ובנין הגוף ברמ"ח איברים ושס"ה גידים. גם נשמת הנשמה יהיה לה איברים וגידים. וכן נשמת הנשמה עד אין קץ יהיה לה גם היא ציור ובנין הגוף שהיא נשמה לו. אלא שהיא דקה וקלושה יותר, וכמו שכתב עז לאלהים בבית קדש הקדשים פ"ג דף ט' ע"א שבהכרח גמור יש לנשמה ציור ובנין הגוף ברמ"ח איברים ושס"ה גידים בצורת אדם, ומתפשטת כפי איבריה. ואם תמצא אבר מיותר אינה מתפשטת בו, וכן ג"כ נשמה לנשמה צריך שיהיה לה בחינת הפרצוף וכו' ע"ש. ועוד שהרי ביארו שהפרצופים כולם נמשכים ומתהוים מעצמותו ית'. גם גופם וגם נשמתם ומלבושם כמו שכתב לעיל בשם שושן סודות. וכמבואר בהדיא בעץ חיים להרח"ו שער עיגולים ויושר ובשער ההקדמות. וכיון שהן נבראים אף שהם מעצמותו לפי דבריהם והוא שוכן בתוכם. איך ניחדם ונעבדם עם בוראם. ובהדיא מזהיר רשב"י ע"ה באמרו כל המשתף וכו' שנאמר בלתי לה' לבדו בלי שתוף וצירוף שום נברא בעולם, וכמו שכתב הרמב"ם ז"ל בהל' ע"ז עיקר הציווי בע"ז וכו'. הובא לעיל סי' נ"ו. וראיה לזה שלא לצרף עמו שום נברא אע"פ שהוא שוכן בו, שהרי צונו הקב"ה על ידי משה רבי' ע"ה לעשות לו משכן וארון וכפרת וכרובים ושאר כלי המשכן שנאמר ועשו לי מקדש ושכנתי בתוכם, וכאשר הוקם המשכן נר' כבודו ית' בו, שנאמר וכבוד ה' מלא וכו' ובבא משה אל אהל מועד, וישמע את הקול מדבר אליו וכו' ואף על פי כן אין אנו משתפים ומיחדים את המשכן וכליו או המקדש וכליו עם כבודו ית' השוכן בתוכו לומר שהמשכן או המקדש אלהינו, או נאמר הארון וכפרת וכרובים אלהינו, הואיל והוא אִוָם למושב לו, ולא נשתחוה ונכרע ונתפלל רק לה' אלהינו השוכן בתוכם, בלי צירוף ושיתוף המקום שהוא שוכן בו ולא ניחדם עמו ונאמר כולא חד.

| id | word | punctuation after |
|---|---|---|
| 8:25.1 | קיח | ) |
| 8:25.2 | ואם |  |
| 8:25.3 | נפשך |  |
| 8:25.4 | לומר |  |
| 8:25.5 | לפי |  |
| 8:25.6 | דבריהם |  |
| 8:25.7 | שהאלהים |  |
| 8:25.8 | העליון |  |
| 8:25.9 | הוא |  |
| 8:25.10 | נשמה |  |
| 8:25.11 | של |  |
| 8:25.12 | הפרצופים |  |
| 8:25.13 | הנבראים |  |
| 8:25.14 | ולנשמתם |  |
| 8:25.15 | אנו |  |
| 8:25.16 | עובדים |  |
| 8:25.17 | ומתפללים | , |
| 8:25.18 | כי |  |
| 8:25.19 | הוא |  |
| 8:25.20 | מתפשט |  |
| 8:25.21 | וממלא |  |
| 8:25.22 | את |  |
| 8:25.23 | הספירות |  |
| 8:25.24 | וכמו |  |
| 8:25.25 | שכתב |  |
| 8:25.26 | הפילסוף |  |
| 8:25.27 | המסית |  |
| 8:25.28 | מחבר |  |
| 8:25.29 | הזהר |  |
| 8:25.30 | דף |  |
| 8:25.31 | הנז | ' |
| 8:25.32 | לעיל |  |
| 8:25.33 | אנא |  |
| 8:25.34 | הוא |  |
| 8:25.35 | לבר |  |
| 8:25.36 | מכלא |  |
| 8:25.37 | ולגיו |  |
| 8:25.38 | מכלא | , |
| 8:25.39 | ולכל |  |
| 8:25.40 | סטרא | , |
| 8:25.41 | ולעילא |  |
| 8:25.42 | מכלא | , |
| 8:25.43 | ולתתא |  |
| 8:25.44 | מכלא | , |
| 8:25.45 | ולית |  |
| 8:25.46 | אלה |  |
| 8:25.47 | אחרא |  |
| 8:25.48 | עילא |  |
| 8:25.49 | ותתא |  |
| 8:25.50 | ומכל |  |
| 8:25.51 | סטרא | , |
| 8:25.52 | ומלגיו |  |
| 8:25.53 | דעסר |  |
| 8:25.54 | ספירן |  |
| 8:25.55 | דמנהון |  |
| 8:25.56 | כלא |  |
| 8:25.57 | ובהון |  |
| 8:25.58 | תליא |  |
| 8:25.59 | כלא |  |
| 8:25.60 | ואנת |  |
| 8:25.61 | בכל |  |
| 8:25.62 | ספירה |  |
| 8:25.63 | בארכה |  |
| 8:25.64 | ורחבה |  |
| 8:25.65 | עילא |  |
| 8:25.66 | ותתא | . |
| 8:25.67 | ובין |  |
| 8:25.68 | כל |  |
| 8:25.69 | ספירה |  |
| 8:25.70 | וספירה | , |
| 8:25.71 | ואנת |  |
| 8:25.72 | הוא |  |
| 8:25.73 | דמקריב |  |
| 8:25.74 | לקודשא |  |
| 8:25.75 | בריך |  |
| 8:25.76 | הוא |  |
| 8:25.77 | ושכיניתיה |  |
| 8:25.78 | בכל |  |
| 8:25.79 | ספירה |  |
| 8:25.80 | וספירה |  |
| 8:25.81 | ובכל |  |
| 8:25.82 | ענפין |  |
| 8:25.83 | וכו | '. |
| 8:25.84 | ובדברים |  |
| 8:25.85 | אלו |  |
| 8:25.86 | לקח |  |
| 8:25.87 | את |  |
| 8:25.88 | לב |  |
| 8:25.89 | קוראיו | , |
| 8:25.90 | וחשבו |  |
| 8:25.91 | לתקן |  |
| 8:25.92 | את |  |
| 8:25.93 | הסתירה |  |
| 8:25.94 | ודברי |  |
| 8:25.95 | מינות |  |
| 8:25.96 | הגלויה |  |
| 8:25.97 | ומפורסמת |  |
| 8:25.98 | שהאלהים |  |
| 8:25.99 | העליון |  |
| 8:25.100 | שאין |  |
| 8:25.101 | כמהו |  |
| 8:25.102 | בכל |  |
| 8:25.103 | האלהים | . |
| 8:25.104 | איננו |  |
| 8:25.105 | רק |  |
| 8:25.106 | מקשר |  |
| 8:25.107 | ומחבר |  |
| 8:25.108 | ומזויג |  |
| 8:25.109 | את |  |
| 8:25.110 | הספירות |  |
| 8:25.111 | הנקראים |  |
| 8:25.112 | קודשא |  |
| 8:25.113 | בריך |  |
| 8:25.114 | הוא |  |
| 8:25.115 | ושכינתיה |  |
| 8:25.116 | אשר |  |
| 8:25.117 | נעבדם |  |
| 8:25.118 | ונשתחוה |  |
| 8:25.119 | להם |  |
| 8:25.120 | לבדם | . |
| 8:25.121 | אבל |  |
| 8:25.122 | הוא |  |
| 8:25.123 | לא |  |
| 8:25.124 | נעבדנו |  |
| 8:25.125 | כנזכר |  |
| 8:25.126 | לעיל | . |
| 8:25.127 | והנה |  |
| 8:25.128 | עם |  |
| 8:25.129 | מה |  |
| 8:25.130 | שהקדמתי |  |
| 8:25.131 | מדברי |  |
| 8:25.132 | רז"ל |  |
| 8:25.133 | אין |  |
| 8:25.134 | קיום |  |
| 8:25.135 | ועמידה |  |
| 8:25.136 | לדברים |  |
| 8:25.137 | אלו |  |
| 8:25.138 | כלל |  |
| 8:25.139 | שכיון |  |
| 8:25.140 | שקיימו |  |
| 8:25.141 | וקבלו |  |
| 8:25.142 | שיש |  |
| 8:25.143 | לנשמה |  |
| 8:25.144 | ציור |  |
| 8:25.145 | ובנין |  |
| 8:25.146 | הגוף |  |
| 8:25.147 | ברמ"ח |  |
| 8:25.148 | איברים |  |
| 8:25.149 | ושס"ה |  |
| 8:25.150 | גידים | . |
| 8:25.151 | גם |  |
| 8:25.152 | נשמת |  |
| 8:25.153 | הנשמה |  |
| 8:25.154 | יהיה |  |
| 8:25.155 | לה |  |
| 8:25.156 | איברים |  |
| 8:25.157 | וגידים | . |
| 8:25.158 | וכן |  |
| 8:25.159 | נשמת |  |
| 8:25.160 | הנשמה |  |
| 8:25.161 | עד |  |
| 8:25.162 | אין |  |
| 8:25.163 | קץ |  |
| 8:25.164 | יהיה |  |
| 8:25.165 | לה |  |
| 8:25.166 | גם |  |
| 8:25.167 | היא |  |
| 8:25.168 | ציור |  |
| 8:25.169 | ובנין |  |
| 8:25.170 | הגוף |  |
| 8:25.171 | שהיא |  |
| 8:25.172 | נשמה |  |
| 8:25.173 | לו | . |
| 8:25.174 | אלא |  |
| 8:25.175 | שהיא |  |
| 8:25.176 | דקה |  |
| 8:25.177 | וקלושה |  |
| 8:25.178 | יותר | , |
| 8:25.179 | וכמו |  |
| 8:25.180 | שכתב |  |
| 8:25.181 | עז |  |
| 8:25.182 | לאלהים |  |
| 8:25.183 | בבית |  |
| 8:25.184 | קדש |  |
| 8:25.185 | הקדשים |  |
| 8:25.186 | פ"ג |  |
| 8:25.187 | דף |  |
| 8:25.188 | ט | ' |
| 8:25.189 | ע"א |  |
| 8:25.190 | שבהכרח |  |
| 8:25.191 | גמור |  |
| 8:25.192 | יש |  |
| 8:25.193 | לנשמה |  |
| 8:25.194 | ציור |  |
| 8:25.195 | ובנין |  |
| 8:25.196 | הגוף |  |
| 8:25.197 | ברמ"ח |  |
| 8:25.198 | איברים |  |
| 8:25.199 | ושס"ה |  |
| 8:25.200 | גידים |  |
| 8:25.201 | בצורת |  |
| 8:25.202 | אדם | , |
| 8:25.203 | ומתפשטת |  |
| 8:25.204 | כפי |  |
| 8:25.205 | איבריה | . |
| 8:25.206 | ואם |  |
| 8:25.207 | תמצא |  |
| 8:25.208 | אבר |  |
| 8:25.209 | מיותר |  |
| 8:25.210 | אינה |  |
| 8:25.211 | מתפשטת |  |
| 8:25.212 | בו | , |
| 8:25.213 | וכן |  |
| 8:25.214 | ג"כ |  |
| 8:25.215 | נשמה |  |
| 8:25.216 | לנשמה |  |
| 8:25.217 | צריך |  |
| 8:25.218 | שיהיה |  |
| 8:25.219 | לה |  |
| 8:25.220 | בחינת |  |
| 8:25.221 | הפרצוף |  |
| 8:25.222 | וכו | ' |
| 8:25.223 | ע"ש | . |
| 8:25.224 | ועוד |  |
| 8:25.225 | שהרי |  |
| 8:25.226 | ביארו |  |
| 8:25.227 | שהפרצופים |  |
| 8:25.228 | כולם |  |
| 8:25.229 | נמשכים |  |
| 8:25.230 | ומתהוים |  |
| 8:25.231 | מעצמותו |  |
| 8:25.232 | ית | '. |
| 8:25.233 | גם |  |
| 8:25.234 | גופם |  |
| 8:25.235 | וגם |  |
| 8:25.236 | נשמתם |  |
| 8:25.237 | ומלבושם |  |
| 8:25.238 | כמו |  |
| 8:25.239 | שכתב |  |
| 8:25.240 | לעיל |  |
| 8:25.241 | בשם |  |
| 8:25.242 | שושן |  |
| 8:25.243 | סודות | . |
| 8:25.244 | וכמבואר |  |
| 8:25.245 | בהדיא |  |
| 8:25.246 | בעץ |  |
| 8:25.247 | חיים |  |
| 8:25.248 | להרח"ו |  |
| 8:25.249 | שער |  |
| 8:25.250 | עיגולים |  |
| 8:25.251 | ויושר |  |
| 8:25.252 | ובשער |  |
| 8:25.253 | ההקדמות | . |
| 8:25.254 | וכיון |  |
| 8:25.255 | שהן |  |
| 8:25.256 | נבראים |  |
| 8:25.257 | אף |  |
| 8:25.258 | שהם |  |
| 8:25.259 | מעצמותו |  |
| 8:25.260 | לפי |  |
| 8:25.261 | דבריהם |  |
| 8:25.262 | והוא |  |
| 8:25.263 | שוכן |  |
| 8:25.264 | בתוכם | . |
| 8:25.265 | איך |  |
| 8:25.266 | ניחדם |  |
| 8:25.267 | ונעבדם |  |
| 8:25.268 | עם |  |
| 8:25.269 | בוראם | . |
| 8:25.270 | ובהדיא |  |
| 8:25.271 | מזהיר |  |
| 8:25.272 | רשב"י |  |
| 8:25.273 | ע"ה |  |
| 8:25.274 | באמרו |  |
| 8:25.275 | כל |  |
| 8:25.276 | המשתף |  |
| 8:25.277 | וכו | ' |
| 8:25.278 | שנאמר |  |
| 8:25.279 | בלתי |  |
| 8:25.280 | לה | ' |
| 8:25.281 | לבדו |  |
| 8:25.282 | בלי |  |
| 8:25.283 | שתוף |  |
| 8:25.284 | וצירוף |  |
| 8:25.285 | שום |  |
| 8:25.286 | נברא |  |
| 8:25.287 | בעולם | , |
| 8:25.288 | וכמו |  |
| 8:25.289 | שכתב |  |
| 8:25.290 | הרמב"ם |  |
| 8:25.291 | ז"ל |  |
| 8:25.292 | בהל | ' |
| 8:25.293 | ע"ז |  |
| 8:25.294 | עיקר |  |
| 8:25.295 | הציווי |  |
| 8:25.296 | בע"ז |  |
| 8:25.297 | וכו | '. |
| 8:25.298 | הובא |  |
| 8:25.299 | לעיל |  |
| 8:25.300 | סי | ' |
| 8:25.301 | נ"ו | . |
| 8:25.302 | וראיה |  |
| 8:25.303 | לזה |  |
| 8:25.304 | שלא |  |
| 8:25.305 | לצרף |  |
| 8:25.306 | עמו |  |
| 8:25.307 | שום |  |
| 8:25.308 | נברא |  |
| 8:25.309 | אע"פ |  |
| 8:25.310 | שהוא |  |
| 8:25.311 | שוכן |  |
| 8:25.312 | בו | , |
| 8:25.313 | שהרי |  |
| 8:25.314 | צונו |  |
| 8:25.315 | הקב"ה |  |
| 8:25.316 | על |  |
| 8:25.317 | ידי |  |
| 8:25.318 | משה |  |
| 8:25.319 | רבי | ' |
| 8:25.320 | ע"ה |  |
| 8:25.321 | לעשות |  |
| 8:25.322 | לו |  |
| 8:25.323 | משכן |  |
| 8:25.324 | וארון |  |
| 8:25.325 | וכפרת |  |
| 8:25.326 | וכרובים |  |
| 8:25.327 | ושאר |  |
| 8:25.328 | כלי |  |
| 8:25.329 | המשכן |  |
| 8:25.330 | שנאמר |  |
| 8:25.331 | ועשו |  |
| 8:25.332 | לי |  |
| 8:25.333 | מקדש |  |
| 8:25.334 | ושכנתי |  |
| 8:25.335 | בתוכם | , |
| 8:25.336 | וכאשר |  |
| 8:25.337 | הוקם |  |
| 8:25.338 | המשכן |  |
| 8:25.339 | נר | ' |
| 8:25.340 | כבודו |  |
| 8:25.341 | ית | ' |
| 8:25.342 | בו | , |
| 8:25.343 | שנאמר |  |
| 8:25.344 | וכבוד |  |
| 8:25.345 | ה | ' |
| 8:25.346 | מלא |  |
| 8:25.347 | וכו | ' |
| 8:25.348 | ובבא |  |
| 8:25.349 | משה |  |
| 8:25.350 | אל |  |
| 8:25.351 | אהל |  |
| 8:25.352 | מועד | , |
| 8:25.353 | וישמע |  |
| 8:25.354 | את |  |
| 8:25.355 | הקול |  |
| 8:25.356 | מדבר |  |
| 8:25.357 | אליו |  |
| 8:25.358 | וכו | ' |
| 8:25.359 | ואף |  |
| 8:25.360 | על |  |
| 8:25.361 | פי |  |
| 8:25.362 | כן |  |
| 8:25.363 | אין |  |
| 8:25.364 | אנו |  |
| 8:25.365 | משתפים |  |
| 8:25.366 | ומיחדים |  |
| 8:25.367 | את |  |
| 8:25.368 | המשכן |  |
| 8:25.369 | וכליו |  |
| 8:25.370 | או |  |
| 8:25.371 | המקדש |  |
| 8:25.372 | וכליו |  |
| 8:25.373 | עם |  |
| 8:25.374 | כבודו |  |
| 8:25.375 | ית | ' |
| 8:25.376 | השוכן |  |
| 8:25.377 | בתוכו |  |
| 8:25.378 | לומר |  |
| 8:25.379 | שהמשכן |  |
| 8:25.380 | או |  |
| 8:25.381 | המקדש |  |
| 8:25.382 | אלהינו | , |
| 8:25.383 | או |  |
| 8:25.384 | נאמר |  |
| 8:25.385 | הארון |  |
| 8:25.386 | וכפרת |  |
| 8:25.387 | וכרובים |  |
| 8:25.388 | אלהינו | , |
| 8:25.389 | הואיל |  |
| 8:25.390 | והוא |  |
| 8:25.391 | אִוָם |  |
| 8:25.392 | למושב |  |
| 8:25.393 | לו | , |
| 8:25.394 | ולא |  |
| 8:25.395 | נשתחוה |  |
| 8:25.396 | ונכרע |  |
| 8:25.397 | ונתפלל |  |
| 8:25.398 | רק |  |
| 8:25.399 | לה | ' |
| 8:25.400 | אלהינו |  |
| 8:25.401 | השוכן |  |
| 8:25.402 | בתוכם | , |
| 8:25.403 | בלי |  |
| 8:25.404 | צירוף |  |
| 8:25.405 | ושיתוף |  |
| 8:25.406 | המקום |  |
| 8:25.407 | שהוא |  |
| 8:25.408 | שוכן |  |
| 8:25.409 | בו |  |
| 8:25.410 | ולא |  |
| 8:25.411 | ניחדם |  |
| 8:25.412 | עמו |  |
| 8:25.413 | ונאמר |  |
| 8:25.414 | כולא |  |
| 8:25.415 | חד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

