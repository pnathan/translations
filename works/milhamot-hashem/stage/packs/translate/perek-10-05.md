# Translation pack: milhamot-hashem perek-10-05 (Milchamot ha-Shem 10:6–10:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-05`.

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
- והוא שלא: "provided that" (fixed).
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
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


## Milchamot ha-Shem 10:6

Hebrew (Wikisource): קמד) סוף דבר תנה לבך קורא נעים, והטה אזנך לתושיה ולתבונה, ושים נא כבוד לה' ותן לו תודה על אשר העיר והאיר ברוח קדשו את חוזינו בכתבי הקדש, וגם אחרי כן החלה רוח ה' לפעם את חכמינו בעלי המשנה ושני התלמודים ומדרשים האמתיים והביטו בעיני שכלם הטהורות מראש ומקדם את העת הרעה, אשר עוד היתה רחוקה מהם כמה מאות שנים והדליקו את האבוקות הגדולות להראות לנו את האמונה הטהורה האחדות השי"ת, בטרם יבוא אפל חשך וצלמות ספר הזהר על ידי המלך העכו"ם העומם את יפעת אמונותינו הטהורה, על ידי אמונתו התפילה והבדויה אשר בשקר הוליד אותה לסלף דברי צדיקים ולדבר עתק עד צדיקו של עולם, כי המשיך הקב"ה מאורו אלהים אחרים ומינה אותם תחתיו, וצוה לעבדם ולהתפלל אליהם. על הראשונים שעשו צלמי זהב וכסף ועץ ואבן אמר הנביא ירמיה ע"ה ובאחת יבערו ויכסלו מוסר הבלים עץ הוא! ועל אנשי דורנו זה ההולכים בעקבות נביא השקר מחבר הזהר אמרה תורה כי יקום בקרבך נביא או חולם חלום וכו' נלכה ונעבדה אלהים אחרים אשר לא ידעתם לא מצד הקבלה שהעתיקו רז"ל ולא מצד המופת כמו שפירש ר' עובדיה ספורנו (דברים י"ג) וז"ל שלא בא מופת אפילו על מציאותם, כי אמנם הנודעים בחוש או במופת נודע מהם שהם פועלים צמיד בענין אחד באופן שנודע שהם פועלים טבעיים לא רצוניים, ובכן לא יועיל להתפלל אליהם ולא לעבדם לא תשמע אל דברי הנביא ההוא לא תפנה אל דבריו לראות אם יש בקצתם, כי אין ספק שכל דבריו קשר וכו' עכ"ל. אשר יאמר שכל שלא בא מופת על מציאותם כצורות ופרצופים שהמציא מחבר הזהר בנבואצו שתלה ברשב"י התנא הכל שקר! שהרי אפי' הנודעים בחוש או במופת פועלים טבעיים הם ולא רצוניים, עליהם אמר שמואל הנביא לישראל ולא תסורו כי אחרי התהו אשר לא יועילו ולא יצילו כי תהו המה! כי כל הפרצופים והצורות והעגולים שהמציא בדמיונו תהו והבל המה, ואין להם מציאות נרגש בחוש וגם אין מופת על מציאותם בה' אלהינו שיש כמה מופתים על מציאותו ית' כמפורש בתורה ובנביאים ובכתובים ובדברי רז"ל! כי אמנם כל דבריו ספורי נפלאות ונמנעות אשר ישמטו את המשנה והתלמוד, וישכינו לעפר כבודם, אף כי בת מלך היא המשנה ועליה עמד משה רבינו בהר ארבעים יום ומ' לילה עד שלמד אותה מפי הגבורה וירד לישראל! ונביא השקר מחבר הזהר בשנאתו הגדולה למשנה ולתלמוד, הרחיה עז בנפשו לבשות את המשנה ולהוריד לארץ את כבודה כדי לשכחה ולעקור אותה ואת חקותיה מישראל בתתו אותה לשמצה! אף כמעט כל קוראיו לא הבינו את מזמותיו, כי חשבו את הכל לתורה אחת ומשפט אחד, ולא הרפו את ידם מלעסוק בזה, וגם מזה לא הניחו ידם! ישראל לא ידע, עמי לא התבונן, שכל מטרתו ומגמתו של הפילסוף מחבר הזהר לתת לנו תורה שבעל פה חדשה תמורת המשנה והתלמוד! מהלל שם ה' החונן לאדם דעת ותבונה, להבין ולהשכיל קשט אמרי אמת, ולנפץ חזון פריצים אשר ינשאו להעמיד חזון ונכשלו והכשילו רבים להאמין בשוא נתעה כי ה' אלהינו איננו אחד ומיוחד, על ידי לשון רמיה ושפת חלקות, עתה יראו כל חכמי לב וישכילו מתורת ה' תמימה שבכתב ושבעל פה המשנה והתלמוד עם מדרשי רז"ל האמתיים ועם פירושו למשניות והקדמותיו עם שאר חכמינו זכי התבונה. ונפשו תשבע להם חלות טהורות משוחות בשטן האמונה הטהורה כנתינתה מסיני. ועל ידה יזכה לחזות בנועם ה' ולבקר בהיכלו, עד כה עזרנו השי"ת לגלות פני הבלות מחשבות הפילסוף נביא השקר המתעה מחבר הזהר, אשר זמם לדבר סרה על ה' ועל תורתו שבעל פה פיו ידבר תהפוכות לבחור אלהים חדשים שאר בדא אותם בדמיונו, ונתלה באילן גדול רשב"י התנא הקדוש, למען תפוש את בית ישראל בלבם! ועל האמונה הטהורה המשנה והתלמוד ישליך שקוצים לשומה שפחה חרופה ולהתעמר בה! וירומם את תורתו החדשה בלשון רמיה ושפת חלקות לכל פתי ותועה רוח! ברוך החונן לאדם דעת להשכיל להטיב, לכו ונלכה באור ה' היא תורתו הקדושה והתמימה והטהורה, אשר תבונינו ותלמדינו דעת את ה' אלהים אמת, ועל ידי זה נזכה לחזות בנועם ה' בעולם הזה ובעולם הבא אכי"ר.

| id | word | punctuation after |
|---|---|---|
| 10:6.1 | קמד | ) |
| 10:6.2 | סוף |  |
| 10:6.3 | דבר |  |
| 10:6.4 | תנה |  |
| 10:6.5 | לבך |  |
| 10:6.6 | קורא |  |
| 10:6.7 | נעים | , |
| 10:6.8 | והטה |  |
| 10:6.9 | אזנך |  |
| 10:6.10 | לתושיה |  |
| 10:6.11 | ולתבונה | , |
| 10:6.12 | ושים |  |
| 10:6.13 | נא |  |
| 10:6.14 | כבוד |  |
| 10:6.15 | לה | ' |
| 10:6.16 | ותן |  |
| 10:6.17 | לו |  |
| 10:6.18 | תודה |  |
| 10:6.19 | על |  |
| 10:6.20 | אשר |  |
| 10:6.21 | העיר |  |
| 10:6.22 | והאיר |  |
| 10:6.23 | ברוח |  |
| 10:6.24 | קדשו |  |
| 10:6.25 | את |  |
| 10:6.26 | חוזינו |  |
| 10:6.27 | בכתבי |  |
| 10:6.28 | הקדש | , |
| 10:6.29 | וגם |  |
| 10:6.30 | אחרי |  |
| 10:6.31 | כן |  |
| 10:6.32 | החלה |  |
| 10:6.33 | רוח |  |
| 10:6.34 | ה | ' |
| 10:6.35 | לפעם |  |
| 10:6.36 | את |  |
| 10:6.37 | חכמינו |  |
| 10:6.38 | בעלי |  |
| 10:6.39 | המשנה |  |
| 10:6.40 | ושני |  |
| 10:6.41 | התלמודים |  |
| 10:6.42 | ומדרשים |  |
| 10:6.43 | האמתיים |  |
| 10:6.44 | והביטו |  |
| 10:6.45 | בעיני |  |
| 10:6.46 | שכלם |  |
| 10:6.47 | הטהורות |  |
| 10:6.48 | מראש |  |
| 10:6.49 | ומקדם |  |
| 10:6.50 | את |  |
| 10:6.51 | העת |  |
| 10:6.52 | הרעה | , |
| 10:6.53 | אשר |  |
| 10:6.54 | עוד |  |
| 10:6.55 | היתה |  |
| 10:6.56 | רחוקה |  |
| 10:6.57 | מהם |  |
| 10:6.58 | כמה |  |
| 10:6.59 | מאות |  |
| 10:6.60 | שנים |  |
| 10:6.61 | והדליקו |  |
| 10:6.62 | את |  |
| 10:6.63 | האבוקות |  |
| 10:6.64 | הגדולות |  |
| 10:6.65 | להראות |  |
| 10:6.66 | לנו |  |
| 10:6.67 | את |  |
| 10:6.68 | האמונה |  |
| 10:6.69 | הטהורה |  |
| 10:6.70 | האחדות |  |
| 10:6.71 | השי"ת | , |
| 10:6.72 | בטרם |  |
| 10:6.73 | יבוא |  |
| 10:6.74 | אפל |  |
| 10:6.75 | חשך |  |
| 10:6.76 | וצלמות |  |
| 10:6.77 | ספר |  |
| 10:6.78 | הזהר |  |
| 10:6.79 | על |  |
| 10:6.80 | ידי |  |
| 10:6.81 | המלך |  |
| 10:6.82 | העכו"ם |  |
| 10:6.83 | העומם |  |
| 10:6.84 | את |  |
| 10:6.85 | יפעת |  |
| 10:6.86 | אמונותינו |  |
| 10:6.87 | הטהורה | , |
| 10:6.88 | על |  |
| 10:6.89 | ידי |  |
| 10:6.90 | אמונתו |  |
| 10:6.91 | התפילה |  |
| 10:6.92 | והבדויה |  |
| 10:6.93 | אשר |  |
| 10:6.94 | בשקר |  |
| 10:6.95 | הוליד |  |
| 10:6.96 | אותה |  |
| 10:6.97 | לסלף |  |
| 10:6.98 | דברי |  |
| 10:6.99 | צדיקים |  |
| 10:6.100 | ולדבר |  |
| 10:6.101 | עתק |  |
| 10:6.102 | עד |  |
| 10:6.103 | צדיקו |  |
| 10:6.104 | של |  |
| 10:6.105 | עולם | , |
| 10:6.106 | כי |  |
| 10:6.107 | המשיך |  |
| 10:6.108 | הקב"ה |  |
| 10:6.109 | מאורו |  |
| 10:6.110 | אלהים |  |
| 10:6.111 | אחרים |  |
| 10:6.112 | ומינה |  |
| 10:6.113 | אותם |  |
| 10:6.114 | תחתיו | , |
| 10:6.115 | וצוה |  |
| 10:6.116 | לעבדם |  |
| 10:6.117 | ולהתפלל |  |
| 10:6.118 | אליהם | . |
| 10:6.119 | על |  |
| 10:6.120 | הראשונים |  |
| 10:6.121 | שעשו |  |
| 10:6.122 | צלמי |  |
| 10:6.123 | זהב |  |
| 10:6.124 | וכסף |  |
| 10:6.125 | ועץ |  |
| 10:6.126 | ואבן |  |
| 10:6.127 | אמר |  |
| 10:6.128 | הנביא |  |
| 10:6.129 | ירמיה |  |
| 10:6.130 | ע"ה |  |
| 10:6.131 | ובאחת |  |
| 10:6.132 | יבערו |  |
| 10:6.133 | ויכסלו |  |
| 10:6.134 | מוסר |  |
| 10:6.135 | הבלים |  |
| 10:6.136 | עץ |  |
| 10:6.137 | הוא | ! |
| 10:6.138 | ועל |  |
| 10:6.139 | אנשי |  |
| 10:6.140 | דורנו |  |
| 10:6.141 | זה |  |
| 10:6.142 | ההולכים |  |
| 10:6.143 | בעקבות |  |
| 10:6.144 | נביא |  |
| 10:6.145 | השקר |  |
| 10:6.146 | מחבר |  |
| 10:6.147 | הזהר |  |
| 10:6.148 | אמרה |  |
| 10:6.149 | תורה |  |
| 10:6.150 | כי |  |
| 10:6.151 | יקום |  |
| 10:6.152 | בקרבך |  |
| 10:6.153 | נביא |  |
| 10:6.154 | או |  |
| 10:6.155 | חולם |  |
| 10:6.156 | חלום |  |
| 10:6.157 | וכו | ' |
| 10:6.158 | נלכה |  |
| 10:6.159 | ונעבדה |  |
| 10:6.160 | אלהים |  |
| 10:6.161 | אחרים |  |
| 10:6.162 | אשר |  |
| 10:6.163 | לא |  |
| 10:6.164 | ידעתם |  |
| 10:6.165 | לא |  |
| 10:6.166 | מצד |  |
| 10:6.167 | הקבלה |  |
| 10:6.168 | שהעתיקו |  |
| 10:6.169 | רז"ל |  |
| 10:6.170 | ולא |  |
| 10:6.171 | מצד |  |
| 10:6.172 | המופת |  |
| 10:6.173 | כמו |  |
| 10:6.174 | שפירש |  |
| 10:6.175 | ר | ' |
| 10:6.176 | עובדיה |  |
| 10:6.177 | ספורנו |  |
| 10:6.178 | דברים |  |
| 10:6.179 | י"ג | ) |
| 10:6.180 | וז"ל |  |
| 10:6.181 | שלא |  |
| 10:6.182 | בא |  |
| 10:6.183 | מופת |  |
| 10:6.184 | אפילו |  |
| 10:6.185 | על |  |
| 10:6.186 | מציאותם | , |
| 10:6.187 | כי |  |
| 10:6.188 | אמנם |  |
| 10:6.189 | הנודעים |  |
| 10:6.190 | בחוש |  |
| 10:6.191 | או |  |
| 10:6.192 | במופת |  |
| 10:6.193 | נודע |  |
| 10:6.194 | מהם |  |
| 10:6.195 | שהם |  |
| 10:6.196 | פועלים |  |
| 10:6.197 | צמיד |  |
| 10:6.198 | בענין |  |
| 10:6.199 | אחד |  |
| 10:6.200 | באופן |  |
| 10:6.201 | שנודע |  |
| 10:6.202 | שהם |  |
| 10:6.203 | פועלים |  |
| 10:6.204 | טבעיים |  |
| 10:6.205 | לא |  |
| 10:6.206 | רצוניים | , |
| 10:6.207 | ובכן |  |
| 10:6.208 | לא |  |
| 10:6.209 | יועיל |  |
| 10:6.210 | להתפלל |  |
| 10:6.211 | אליהם |  |
| 10:6.212 | ולא |  |
| 10:6.213 | לעבדם |  |
| 10:6.214 | לא |  |
| 10:6.215 | תשמע |  |
| 10:6.216 | אל |  |
| 10:6.217 | דברי |  |
| 10:6.218 | הנביא |  |
| 10:6.219 | ההוא |  |
| 10:6.220 | לא |  |
| 10:6.221 | תפנה |  |
| 10:6.222 | אל |  |
| 10:6.223 | דבריו |  |
| 10:6.224 | לראות |  |
| 10:6.225 | אם |  |
| 10:6.226 | יש |  |
| 10:6.227 | בקצתם | , |
| 10:6.228 | כי |  |
| 10:6.229 | אין |  |
| 10:6.230 | ספק |  |
| 10:6.231 | שכל |  |
| 10:6.232 | דבריו |  |
| 10:6.233 | קשר |  |
| 10:6.234 | וכו | ' |
| 10:6.235 | עכ"ל | . |
| 10:6.236 | אשר |  |
| 10:6.237 | יאמר |  |
| 10:6.238 | שכל |  |
| 10:6.239 | שלא |  |
| 10:6.240 | בא |  |
| 10:6.241 | מופת |  |
| 10:6.242 | על |  |
| 10:6.243 | מציאותם |  |
| 10:6.244 | כצורות |  |
| 10:6.245 | ופרצופים |  |
| 10:6.246 | שהמציא |  |
| 10:6.247 | מחבר |  |
| 10:6.248 | הזהר |  |
| 10:6.249 | בנבואצו |  |
| 10:6.250 | שתלה |  |
| 10:6.251 | ברשב"י |  |
| 10:6.252 | התנא |  |
| 10:6.253 | הכל |  |
| 10:6.254 | שקר | ! |
| 10:6.255 | שהרי |  |
| 10:6.256 | אפי | ' |
| 10:6.257 | הנודעים |  |
| 10:6.258 | בחוש |  |
| 10:6.259 | או |  |
| 10:6.260 | במופת |  |
| 10:6.261 | פועלים |  |
| 10:6.262 | טבעיים |  |
| 10:6.263 | הם |  |
| 10:6.264 | ולא |  |
| 10:6.265 | רצוניים | , |
| 10:6.266 | עליהם |  |
| 10:6.267 | אמר |  |
| 10:6.268 | שמואל |  |
| 10:6.269 | הנביא |  |
| 10:6.270 | לישראל |  |
| 10:6.271 | ולא |  |
| 10:6.272 | תסורו |  |
| 10:6.273 | כי |  |
| 10:6.274 | אחרי |  |
| 10:6.275 | התהו |  |
| 10:6.276 | אשר |  |
| 10:6.277 | לא |  |
| 10:6.278 | יועילו |  |
| 10:6.279 | ולא |  |
| 10:6.280 | יצילו |  |
| 10:6.281 | כי |  |
| 10:6.282 | תהו |  |
| 10:6.283 | המה | ! |
| 10:6.284 | כי |  |
| 10:6.285 | כל |  |
| 10:6.286 | הפרצופים |  |
| 10:6.287 | והצורות |  |
| 10:6.288 | והעגולים |  |
| 10:6.289 | שהמציא |  |
| 10:6.290 | בדמיונו |  |
| 10:6.291 | תהו |  |
| 10:6.292 | והבל |  |
| 10:6.293 | המה | , |
| 10:6.294 | ואין |  |
| 10:6.295 | להם |  |
| 10:6.296 | מציאות |  |
| 10:6.297 | נרגש |  |
| 10:6.298 | בחוש |  |
| 10:6.299 | וגם |  |
| 10:6.300 | אין |  |
| 10:6.301 | מופת |  |
| 10:6.302 | על |  |
| 10:6.303 | מציאותם |  |
| 10:6.304 | בה | ' |
| 10:6.305 | אלהינו |  |
| 10:6.306 | שיש |  |
| 10:6.307 | כמה |  |
| 10:6.308 | מופתים |  |
| 10:6.309 | על |  |
| 10:6.310 | מציאותו |  |
| 10:6.311 | ית | ' |
| 10:6.312 | כמפורש |  |
| 10:6.313 | בתורה |  |
| 10:6.314 | ובנביאים |  |
| 10:6.315 | ובכתובים |  |
| 10:6.316 | ובדברי |  |
| 10:6.317 | רז"ל | ! |
| 10:6.318 | כי |  |
| 10:6.319 | אמנם |  |
| 10:6.320 | כל |  |
| 10:6.321 | דבריו |  |
| 10:6.322 | ספורי |  |
| 10:6.323 | נפלאות |  |
| 10:6.324 | ונמנעות |  |
| 10:6.325 | אשר |  |
| 10:6.326 | ישמטו |  |
| 10:6.327 | את |  |
| 10:6.328 | המשנה |  |
| 10:6.329 | והתלמוד | , |
| 10:6.330 | וישכינו |  |
| 10:6.331 | לעפר |  |
| 10:6.332 | כבודם | , |
| 10:6.333 | אף |  |
| 10:6.334 | כי |  |
| 10:6.335 | בת |  |
| 10:6.336 | מלך |  |
| 10:6.337 | היא |  |
| 10:6.338 | המשנה |  |
| 10:6.339 | ועליה |  |
| 10:6.340 | עמד |  |
| 10:6.341 | משה |  |
| 10:6.342 | רבינו |  |
| 10:6.343 | בהר |  |
| 10:6.344 | ארבעים |  |
| 10:6.345 | יום |  |
| 10:6.346 | ומ | ' |
| 10:6.347 | לילה |  |
| 10:6.348 | עד |  |
| 10:6.349 | שלמד |  |
| 10:6.350 | אותה |  |
| 10:6.351 | מפי |  |
| 10:6.352 | הגבורה |  |
| 10:6.353 | וירד |  |
| 10:6.354 | לישראל | ! |
| 10:6.355 | ונביא |  |
| 10:6.356 | השקר |  |
| 10:6.357 | מחבר |  |
| 10:6.358 | הזהר |  |
| 10:6.359 | בשנאתו |  |
| 10:6.360 | הגדולה |  |
| 10:6.361 | למשנה |  |
| 10:6.362 | ולתלמוד | , |
| 10:6.363 | הרחיה |  |
| 10:6.364 | עז |  |
| 10:6.365 | בנפשו |  |
| 10:6.366 | לבשות |  |
| 10:6.367 | את |  |
| 10:6.368 | המשנה |  |
| 10:6.369 | ולהוריד |  |
| 10:6.370 | לארץ |  |
| 10:6.371 | את |  |
| 10:6.372 | כבודה |  |
| 10:6.373 | כדי |  |
| 10:6.374 | לשכחה |  |
| 10:6.375 | ולעקור |  |
| 10:6.376 | אותה |  |
| 10:6.377 | ואת |  |
| 10:6.378 | חקותיה |  |
| 10:6.379 | מישראל |  |
| 10:6.380 | בתתו |  |
| 10:6.381 | אותה |  |
| 10:6.382 | לשמצה | ! |
| 10:6.383 | אף |  |
| 10:6.384 | כמעט |  |
| 10:6.385 | כל |  |
| 10:6.386 | קוראיו |  |
| 10:6.387 | לא |  |
| 10:6.388 | הבינו |  |
| 10:6.389 | את |  |
| 10:6.390 | מזמותיו | , |
| 10:6.391 | כי |  |
| 10:6.392 | חשבו |  |
| 10:6.393 | את |  |
| 10:6.394 | הכל |  |
| 10:6.395 | לתורה |  |
| 10:6.396 | אחת |  |
| 10:6.397 | ומשפט |  |
| 10:6.398 | אחד | , |
| 10:6.399 | ולא |  |
| 10:6.400 | הרפו |  |
| 10:6.401 | את |  |
| 10:6.402 | ידם |  |
| 10:6.403 | מלעסוק |  |
| 10:6.404 | בזה | , |
| 10:6.405 | וגם |  |
| 10:6.406 | מזה |  |
| 10:6.407 | לא |  |
| 10:6.408 | הניחו |  |
| 10:6.409 | ידם | ! |
| 10:6.410 | ישראל |  |
| 10:6.411 | לא |  |
| 10:6.412 | ידע | , |
| 10:6.413 | עמי |  |
| 10:6.414 | לא |  |
| 10:6.415 | התבונן | , |
| 10:6.416 | שכל |  |
| 10:6.417 | מטרתו |  |
| 10:6.418 | ומגמתו |  |
| 10:6.419 | של |  |
| 10:6.420 | הפילסוף |  |
| 10:6.421 | מחבר |  |
| 10:6.422 | הזהר |  |
| 10:6.423 | לתת |  |
| 10:6.424 | לנו |  |
| 10:6.425 | תורה |  |
| 10:6.426 | שבעל |  |
| 10:6.427 | פה |  |
| 10:6.428 | חדשה |  |
| 10:6.429 | תמורת |  |
| 10:6.430 | המשנה |  |
| 10:6.431 | והתלמוד | ! |
| 10:6.432 | מהלל |  |
| 10:6.433 | שם |  |
| 10:6.434 | ה | ' |
| 10:6.435 | החונן |  |
| 10:6.436 | לאדם |  |
| 10:6.437 | דעת |  |
| 10:6.438 | ותבונה | , |
| 10:6.439 | להבין |  |
| 10:6.440 | ולהשכיל |  |
| 10:6.441 | קשט |  |
| 10:6.442 | אמרי |  |
| 10:6.443 | אמת | , |
| 10:6.444 | ולנפץ |  |
| 10:6.445 | חזון |  |
| 10:6.446 | פריצים |  |
| 10:6.447 | אשר |  |
| 10:6.448 | ינשאו |  |
| 10:6.449 | להעמיד |  |
| 10:6.450 | חזון |  |
| 10:6.451 | ונכשלו |  |
| 10:6.452 | והכשילו |  |
| 10:6.453 | רבים |  |
| 10:6.454 | להאמין |  |
| 10:6.455 | בשוא |  |
| 10:6.456 | נתעה |  |
| 10:6.457 | כי |  |
| 10:6.458 | ה | ' |
| 10:6.459 | אלהינו |  |
| 10:6.460 | איננו |  |
| 10:6.461 | אחד |  |
| 10:6.462 | ומיוחד | , |
| 10:6.463 | על |  |
| 10:6.464 | ידי |  |
| 10:6.465 | לשון |  |
| 10:6.466 | רמיה |  |
| 10:6.467 | ושפת |  |
| 10:6.468 | חלקות | , |
| 10:6.469 | עתה |  |
| 10:6.470 | יראו |  |
| 10:6.471 | כל |  |
| 10:6.472 | חכמי |  |
| 10:6.473 | לב |  |
| 10:6.474 | וישכילו |  |
| 10:6.475 | מתורת |  |
| 10:6.476 | ה | ' |
| 10:6.477 | תמימה |  |
| 10:6.478 | שבכתב |  |
| 10:6.479 | ושבעל |  |
| 10:6.480 | פה |  |
| 10:6.481 | המשנה |  |
| 10:6.482 | והתלמוד |  |
| 10:6.483 | עם |  |
| 10:6.484 | מדרשי |  |
| 10:6.485 | רז"ל |  |
| 10:6.486 | האמתיים |  |
| 10:6.487 | ועם |  |
| 10:6.488 | פירושו |  |
| 10:6.489 | למשניות |  |
| 10:6.490 | והקדמותיו |  |
| 10:6.491 | עם |  |
| 10:6.492 | שאר |  |
| 10:6.493 | חכמינו |  |
| 10:6.494 | זכי |  |
| 10:6.495 | התבונה | . |
| 10:6.496 | ונפשו |  |
| 10:6.497 | תשבע |  |
| 10:6.498 | להם |  |
| 10:6.499 | חלות |  |
| 10:6.500 | טהורות |  |
| 10:6.501 | משוחות |  |
| 10:6.502 | בשטן |  |
| 10:6.503 | האמונה |  |
| 10:6.504 | הטהורה |  |
| 10:6.505 | כנתינתה |  |
| 10:6.506 | מסיני | . |
| 10:6.507 | ועל |  |
| 10:6.508 | ידה |  |
| 10:6.509 | יזכה |  |
| 10:6.510 | לחזות |  |
| 10:6.511 | בנועם |  |
| 10:6.512 | ה | ' |
| 10:6.513 | ולבקר |  |
| 10:6.514 | בהיכלו | , |
| 10:6.515 | עד |  |
| 10:6.516 | כה |  |
| 10:6.517 | עזרנו |  |
| 10:6.518 | השי"ת |  |
| 10:6.519 | לגלות |  |
| 10:6.520 | פני |  |
| 10:6.521 | הבלות |  |
| 10:6.522 | מחשבות |  |
| 10:6.523 | הפילסוף |  |
| 10:6.524 | נביא |  |
| 10:6.525 | השקר |  |
| 10:6.526 | המתעה |  |
| 10:6.527 | מחבר |  |
| 10:6.528 | הזהר | , |
| 10:6.529 | אשר |  |
| 10:6.530 | זמם |  |
| 10:6.531 | לדבר |  |
| 10:6.532 | סרה |  |
| 10:6.533 | על |  |
| 10:6.534 | ה | ' |
| 10:6.535 | ועל |  |
| 10:6.536 | תורתו |  |
| 10:6.537 | שבעל |  |
| 10:6.538 | פה |  |
| 10:6.539 | פיו |  |
| 10:6.540 | ידבר |  |
| 10:6.541 | תהפוכות |  |
| 10:6.542 | לבחור |  |
| 10:6.543 | אלהים |  |
| 10:6.544 | חדשים |  |
| 10:6.545 | שאר |  |
| 10:6.546 | בדא |  |
| 10:6.547 | אותם |  |
| 10:6.548 | בדמיונו | , |
| 10:6.549 | ונתלה |  |
| 10:6.550 | באילן |  |
| 10:6.551 | גדול |  |
| 10:6.552 | רשב"י |  |
| 10:6.553 | התנא |  |
| 10:6.554 | הקדוש | , |
| 10:6.555 | למען |  |
| 10:6.556 | תפוש |  |
| 10:6.557 | את |  |
| 10:6.558 | בית |  |
| 10:6.559 | ישראל |  |
| 10:6.560 | בלבם | ! |
| 10:6.561 | ועל |  |
| 10:6.562 | האמונה |  |
| 10:6.563 | הטהורה |  |
| 10:6.564 | המשנה |  |
| 10:6.565 | והתלמוד |  |
| 10:6.566 | ישליך |  |
| 10:6.567 | שקוצים |  |
| 10:6.568 | לשומה |  |
| 10:6.569 | שפחה |  |
| 10:6.570 | חרופה |  |
| 10:6.571 | ולהתעמר |  |
| 10:6.572 | בה | ! |
| 10:6.573 | וירומם |  |
| 10:6.574 | את |  |
| 10:6.575 | תורתו |  |
| 10:6.576 | החדשה |  |
| 10:6.577 | בלשון |  |
| 10:6.578 | רמיה |  |
| 10:6.579 | ושפת |  |
| 10:6.580 | חלקות |  |
| 10:6.581 | לכל |  |
| 10:6.582 | פתי |  |
| 10:6.583 | ותועה |  |
| 10:6.584 | רוח | ! |
| 10:6.585 | ברוך |  |
| 10:6.586 | החונן |  |
| 10:6.587 | לאדם |  |
| 10:6.588 | דעת |  |
| 10:6.589 | להשכיל |  |
| 10:6.590 | להטיב | , |
| 10:6.591 | לכו |  |
| 10:6.592 | ונלכה |  |
| 10:6.593 | באור |  |
| 10:6.594 | ה | ' |
| 10:6.595 | היא |  |
| 10:6.596 | תורתו |  |
| 10:6.597 | הקדושה |  |
| 10:6.598 | והתמימה |  |
| 10:6.599 | והטהורה | , |
| 10:6.600 | אשר |  |
| 10:6.601 | תבונינו |  |
| 10:6.602 | ותלמדינו |  |
| 10:6.603 | דעת |  |
| 10:6.604 | את |  |
| 10:6.605 | ה | ' |
| 10:6.606 | אלהים |  |
| 10:6.607 | אמת | , |
| 10:6.608 | ועל |  |
| 10:6.609 | ידי |  |
| 10:6.610 | זה |  |
| 10:6.611 | נזכה |  |
| 10:6.612 | לחזות |  |
| 10:6.613 | בנועם |  |
| 10:6.614 | ה | ' |
| 10:6.615 | בעולם |  |
| 10:6.616 | הזה |  |
| 10:6.617 | ובעולם |  |
| 10:6.618 | הבא |  |
| 10:6.619 | אכי"ר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

