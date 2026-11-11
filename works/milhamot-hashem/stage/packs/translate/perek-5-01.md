# Translation pack: milhamot-hashem perek-5-01 (Milchamot ha-Shem 5:1–5:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-01`.

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
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted", with
  "tzimtzum" in a note at the first occurrence in a unit. "אחר הצמצום" →
  "after the contraction".
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
- המתפלספים: "would-be philosophers"; הנצרים = "the Christians" (spelling variant
  of הנוצרים, noted); העשיריות: "the tens" (the ten sefirot, noted).
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


## Milchamot ha-Shem 5:1

Hebrew (Wikisource): מ) והמקובלים החדשים שציירו את האין סוף המקיף אחר הצמצום ככדור עגול וחלול באמצעו ובתוכו עיגולים לאלפים ורבבות הנמשכים מעצמותו יתברך, גם פרצופים דיושר אין קץ למספרם שנמשכו ונשתלשלו מעצמותו דרך הקו המתפשט ממנו מאדם קדמאה שהוא למעלה מעולם האצילות ועד סוף פרצופי כל האבי"ע, עליהם אמר מורנו ורבנו הרמב"ם ז"ל בפרק נ' מראשון: "קרוב אתה בפיהם ורחוק מכליותיהם" (ע"פ ירמיהו י"ב ב'). כי בפיהם אומרים שהוא אחד ובכל לבבם ובכל נפשם ומחשבותם - שהוא נחלק לכמה פרצופים ומדרגות זה למעלה מזה, וצריכים הם ליטול רשות זה מזה כמו שנתבאר לעיל. וכמעט שאמונה זו גרועה מאמונת העכו"ם וכמו שכתב הריב"ש בשם אחד מן המתפלספים. כי הנוצרים אומרים שהם שלושה והשלושה אחד, והמקובלים אומרים שהם חמישה בעולם האצילות והחמישה נחלקים לי"ב, והכל אחד. וזוכר אני בספר בית יהודה כתב שהדברים הנזכרים בריב"ש בשם אחד מהמתפלספים מילתא דנפשיה היא, רק שתלה אותם באחד מן המתפלספים פן יפגעו בו הרודפים אחר הכבוד ויזיקוהו, לכן תלה הדבר באחרים. ולכך סיים שאינו תוקע עצמו ללמוד הקבלה החדשה, כי די לו בקבלה הישנה שהיא המשנה והתלמוד. והמקובלים החדשים המשילו מציאותו יתברך ואמתתו לגוף האדם שהוא מורכב מרמ"ח אברים ושס"ה גידים והוא נקרא בשם אחד ראובן, או לבית המחובר מאבנים וסיד ועצים וטיט ויש בו כמה חדרים, והוא הנקרא בשם אחד בית. ולא זו בלבד אלא גם בסיבה הראשונה הנקרא אצלם אין סוף עשו בו ריבוי כמ"ש מהרח"ו בספר ארבע מאות שקל כסף דף ע"ב וז"ל: "ראוי שתדע כי הלא כל המציאות כולו האין סוף אשר אנו קורין אותו עתיקא דכל עתיקין האמור כדי שיתגלה לנו קצת הארה ממנו, והוא גנוז וטמיר בכתר כנודע דאין סוף גנוז בכתר. וראוי שתדע כי יש עדיין אין סוף אשר הוא לעילא לעילא, כי אין שום מחשבה תפיס ביה והוא עמוק מי ימצאנו, ולכן תמצא בזהר כמה מיני אין סוף בכמה דוכתי אשר כמעט יסתכן האדם בנפשו, יען מי שאין לו הקדמות אינו יכול לשוט בימא רבה. ולכן על זה האין סוף למטה למדו בזהר אדרת נשא. ואז בהיותו שם בסוד תלת רישין אז הוא נקרא עתיקא דכל עתיקין, וז"ס מה שאמרו האי עתיקא קדישא אשתכח בתלת רישין. נמצא כי עיקר עתיקא קדישא הוא האין סוף, אמנם ע"ש התלבשותו בתלת רישין נקראים התלת רישין עתיקא קדישא", עכ"ל. מתבאר בהדיא מדבריו כי הא"ס הראשון שהוא עמוק עמוק ומקיף את כל העולמות אין לכוון תפילתו אליו כלל, כי הוא עמוק עמוק ורחוק מאד ולית מחשבה תפיסא ביה כלל וכלל. וכל תפילותנו אל הזעיר אנפין להמשיך השפע מאבא ואמא ומאין סוף הגנוז בכתר בלבד יתגלה לנו הארתו. אבל אין סוף הקדום הנעלה לית מחשבה תפיסה ביה כלל. ומי שלא ידע זה יבוא לכוון תפילותיו אליו, ויסתכן בנפשו שלא תקובל תפילתו. ואדרבה ענוש ייענש כמ"ש לעיל בשם ספר כסא אליהו. ודברים אלה היפך דבריך שכתבת בשם לחם שלמה. ואין ספר לחם שלמה מצוי אצלנו לעיין בו ואפשר שאתה מטעה אותנו וכותב לנו בשמו היפך כוונתו כמו שעשית בכוונת כסא אליהו ואינו כן. כי המוסכם על בעלי הקבלה החדשה שהאין סוף שהוא הסיבה הראשונה אין שום תפילה וברכה שייכא ביה לדעתם כלל כי הוא עמוק ורחוק מאד. ורק כל תפילותנו להמשיך מאין סוף הגנוז בכתר בלבד כדי שיתגלה לנו הארתו. ולמה זה לא נכוון בתפילותנו אל האין סוף הקדום המקיף את כל המציאות? ולמה יקצוף עלינו בעת אשר נקרא אליו? וגם למה לא הודיע כל זאת ע"י נאמן ביתו משה איש האלהים, לא בתושב"כ ולא בתושבע"פ שהיא המשנה והתלמוד המקובל מפה אל פה? וכן כל הנביאים שעמדו להם לישראל מימות משה ועד האחרון שבהם - מלאכי - המזהיר אותנו (מלאכי ג' כ"ב): "זכרו תורת משה עבדי אשר צויתי אותו בחורב על כל ישראל חוקים ומשפטים" - ולא צירופים ותיקון עולמות. ולמה לא הודיעונו כ"ז? וילכו כל הדורות שעברו שולל (לפי דבריהם) להתפלל ח"ו אל אל לא יושיע? היה להם להודיע לישראל כי לא אל אל עליון נורא יתפללו כי אינו מושג. ורק אל קצר האפים (הזעיר אנפין) שכבר נתעבה והושג בחוש יתפללו. והן עתה נכפלה קושיתנו וחזקה. כי הרבה מיני אין סוף הם והרבה פרצופים יש שכ"א עושה פעולה, וזה הוא אמונת הרבה רשויות בשמים, וכבא הבאתי לעיל דברי כסא אליהו להודיע לך טעותך שחשבת לישב דבריו. וכל דבריך הבל וריק ישאם רוח. וברור הוא שדעתו שצריך שיתפלל לזעיר אנפין בשיתוף שאר הפרצופים שלמעלה ממנו.

| id | word | punctuation after |
|---|---|---|
| 5:1.1 | מ | ) |
| 5:1.2 | והמקובלים |  |
| 5:1.3 | החדשים |  |
| 5:1.4 | שציירו |  |
| 5:1.5 | את |  |
| 5:1.6 | האין |  |
| 5:1.7 | סוף |  |
| 5:1.8 | המקיף |  |
| 5:1.9 | אחר |  |
| 5:1.10 | הצמצום |  |
| 5:1.11 | ככדור |  |
| 5:1.12 | עגול |  |
| 5:1.13 | וחלול |  |
| 5:1.14 | באמצעו |  |
| 5:1.15 | ובתוכו |  |
| 5:1.16 | עיגולים |  |
| 5:1.17 | לאלפים |  |
| 5:1.18 | ורבבות |  |
| 5:1.19 | הנמשכים |  |
| 5:1.20 | מעצמותו |  |
| 5:1.21 | יתברך | , |
| 5:1.22 | גם |  |
| 5:1.23 | פרצופים |  |
| 5:1.24 | דיושר |  |
| 5:1.25 | אין |  |
| 5:1.26 | קץ |  |
| 5:1.27 | למספרם |  |
| 5:1.28 | שנמשכו |  |
| 5:1.29 | ונשתלשלו |  |
| 5:1.30 | מעצמותו |  |
| 5:1.31 | דרך |  |
| 5:1.32 | הקו |  |
| 5:1.33 | המתפשט |  |
| 5:1.34 | ממנו |  |
| 5:1.35 | מאדם |  |
| 5:1.36 | קדמאה |  |
| 5:1.37 | שהוא |  |
| 5:1.38 | למעלה |  |
| 5:1.39 | מעולם |  |
| 5:1.40 | האצילות |  |
| 5:1.41 | ועד |  |
| 5:1.42 | סוף |  |
| 5:1.43 | פרצופי |  |
| 5:1.44 | כל |  |
| 5:1.45 | האבי"ע | , |
| 5:1.46 | עליהם |  |
| 5:1.47 | אמר |  |
| 5:1.48 | מורנו |  |
| 5:1.49 | ורבנו |  |
| 5:1.50 | הרמב"ם |  |
| 5:1.51 | ז"ל |  |
| 5:1.52 | בפרק |  |
| 5:1.53 | נ | ' |
| 5:1.54 | מראשון | : |
| 5:1.55 | קרוב |  |
| 5:1.56 | אתה |  |
| 5:1.57 | בפיהם |  |
| 5:1.58 | ורחוק |  |
| 5:1.59 | מכליותיהם | " |
| 5:1.60 | ע"פ |  |
| 5:1.61 | ירמיהו |  |
| 5:1.62 | י"ב |  |
| 5:1.63 | ב | '). |
| 5:1.64 | כי |  |
| 5:1.65 | בפיהם |  |
| 5:1.66 | אומרים |  |
| 5:1.67 | שהוא |  |
| 5:1.68 | אחד |  |
| 5:1.69 | ובכל |  |
| 5:1.70 | לבבם |  |
| 5:1.71 | ובכל |  |
| 5:1.72 | נפשם |  |
| 5:1.73 | ומחשבותם | - |
| 5:1.74 | שהוא |  |
| 5:1.75 | נחלק |  |
| 5:1.76 | לכמה |  |
| 5:1.77 | פרצופים |  |
| 5:1.78 | ומדרגות |  |
| 5:1.79 | זה |  |
| 5:1.80 | למעלה |  |
| 5:1.81 | מזה | , |
| 5:1.82 | וצריכים |  |
| 5:1.83 | הם |  |
| 5:1.84 | ליטול |  |
| 5:1.85 | רשות |  |
| 5:1.86 | זה |  |
| 5:1.87 | מזה |  |
| 5:1.88 | כמו |  |
| 5:1.89 | שנתבאר |  |
| 5:1.90 | לעיל | . |
| 5:1.91 | וכמעט |  |
| 5:1.92 | שאמונה |  |
| 5:1.93 | זו |  |
| 5:1.94 | גרועה |  |
| 5:1.95 | מאמונת |  |
| 5:1.96 | העכו"ם |  |
| 5:1.97 | וכמו |  |
| 5:1.98 | שכתב |  |
| 5:1.99 | הריב"ש |  |
| 5:1.100 | בשם |  |
| 5:1.101 | אחד |  |
| 5:1.102 | מן |  |
| 5:1.103 | המתפלספים | . |
| 5:1.104 | כי |  |
| 5:1.105 | הנוצרים |  |
| 5:1.106 | אומרים |  |
| 5:1.107 | שהם |  |
| 5:1.108 | שלושה |  |
| 5:1.109 | והשלושה |  |
| 5:1.110 | אחד | , |
| 5:1.111 | והמקובלים |  |
| 5:1.112 | אומרים |  |
| 5:1.113 | שהם |  |
| 5:1.114 | חמישה |  |
| 5:1.115 | בעולם |  |
| 5:1.116 | האצילות |  |
| 5:1.117 | והחמישה |  |
| 5:1.118 | נחלקים |  |
| 5:1.119 | לי"ב | , |
| 5:1.120 | והכל |  |
| 5:1.121 | אחד | . |
| 5:1.122 | וזוכר |  |
| 5:1.123 | אני |  |
| 5:1.124 | בספר |  |
| 5:1.125 | בית |  |
| 5:1.126 | יהודה |  |
| 5:1.127 | כתב |  |
| 5:1.128 | שהדברים |  |
| 5:1.129 | הנזכרים |  |
| 5:1.130 | בריב"ש |  |
| 5:1.131 | בשם |  |
| 5:1.132 | אחד |  |
| 5:1.133 | מהמתפלספים |  |
| 5:1.134 | מילתא |  |
| 5:1.135 | דנפשיה |  |
| 5:1.136 | היא | , |
| 5:1.137 | רק |  |
| 5:1.138 | שתלה |  |
| 5:1.139 | אותם |  |
| 5:1.140 | באחד |  |
| 5:1.141 | מן |  |
| 5:1.142 | המתפלספים |  |
| 5:1.143 | פן |  |
| 5:1.144 | יפגעו |  |
| 5:1.145 | בו |  |
| 5:1.146 | הרודפים |  |
| 5:1.147 | אחר |  |
| 5:1.148 | הכבוד |  |
| 5:1.149 | ויזיקוהו | , |
| 5:1.150 | לכן |  |
| 5:1.151 | תלה |  |
| 5:1.152 | הדבר |  |
| 5:1.153 | באחרים | . |
| 5:1.154 | ולכך |  |
| 5:1.155 | סיים |  |
| 5:1.156 | שאינו |  |
| 5:1.157 | תוקע |  |
| 5:1.158 | עצמו |  |
| 5:1.159 | ללמוד |  |
| 5:1.160 | הקבלה |  |
| 5:1.161 | החדשה | , |
| 5:1.162 | כי |  |
| 5:1.163 | די |  |
| 5:1.164 | לו |  |
| 5:1.165 | בקבלה |  |
| 5:1.166 | הישנה |  |
| 5:1.167 | שהיא |  |
| 5:1.168 | המשנה |  |
| 5:1.169 | והתלמוד | . |
| 5:1.170 | והמקובלים |  |
| 5:1.171 | החדשים |  |
| 5:1.172 | המשילו |  |
| 5:1.173 | מציאותו |  |
| 5:1.174 | יתברך |  |
| 5:1.175 | ואמתתו |  |
| 5:1.176 | לגוף |  |
| 5:1.177 | האדם |  |
| 5:1.178 | שהוא |  |
| 5:1.179 | מורכב |  |
| 5:1.180 | מרמ"ח |  |
| 5:1.181 | אברים |  |
| 5:1.182 | ושס"ה |  |
| 5:1.183 | גידים |  |
| 5:1.184 | והוא |  |
| 5:1.185 | נקרא |  |
| 5:1.186 | בשם |  |
| 5:1.187 | אחד |  |
| 5:1.188 | ראובן | , |
| 5:1.189 | או |  |
| 5:1.190 | לבית |  |
| 5:1.191 | המחובר |  |
| 5:1.192 | מאבנים |  |
| 5:1.193 | וסיד |  |
| 5:1.194 | ועצים |  |
| 5:1.195 | וטיט |  |
| 5:1.196 | ויש |  |
| 5:1.197 | בו |  |
| 5:1.198 | כמה |  |
| 5:1.199 | חדרים | , |
| 5:1.200 | והוא |  |
| 5:1.201 | הנקרא |  |
| 5:1.202 | בשם |  |
| 5:1.203 | אחד |  |
| 5:1.204 | בית | . |
| 5:1.205 | ולא |  |
| 5:1.206 | זו |  |
| 5:1.207 | בלבד |  |
| 5:1.208 | אלא |  |
| 5:1.209 | גם |  |
| 5:1.210 | בסיבה |  |
| 5:1.211 | הראשונה |  |
| 5:1.212 | הנקרא |  |
| 5:1.213 | אצלם |  |
| 5:1.214 | אין |  |
| 5:1.215 | סוף |  |
| 5:1.216 | עשו |  |
| 5:1.217 | בו |  |
| 5:1.218 | ריבוי |  |
| 5:1.219 | כמ"ש |  |
| 5:1.220 | מהרח"ו |  |
| 5:1.221 | בספר |  |
| 5:1.222 | ארבע |  |
| 5:1.223 | מאות |  |
| 5:1.224 | שקל |  |
| 5:1.225 | כסף |  |
| 5:1.226 | דף |  |
| 5:1.227 | ע"ב |  |
| 5:1.228 | וז"ל | : |
| 5:1.229 | ראוי |  |
| 5:1.230 | שתדע |  |
| 5:1.231 | כי |  |
| 5:1.232 | הלא |  |
| 5:1.233 | כל |  |
| 5:1.234 | המציאות |  |
| 5:1.235 | כולו |  |
| 5:1.236 | האין |  |
| 5:1.237 | סוף |  |
| 5:1.238 | אשר |  |
| 5:1.239 | אנו |  |
| 5:1.240 | קורין |  |
| 5:1.241 | אותו |  |
| 5:1.242 | עתיקא |  |
| 5:1.243 | דכל |  |
| 5:1.244 | עתיקין |  |
| 5:1.245 | האמור |  |
| 5:1.246 | כדי |  |
| 5:1.247 | שיתגלה |  |
| 5:1.248 | לנו |  |
| 5:1.249 | קצת |  |
| 5:1.250 | הארה |  |
| 5:1.251 | ממנו | , |
| 5:1.252 | והוא |  |
| 5:1.253 | גנוז |  |
| 5:1.254 | וטמיר |  |
| 5:1.255 | בכתר |  |
| 5:1.256 | כנודע |  |
| 5:1.257 | דאין |  |
| 5:1.258 | סוף |  |
| 5:1.259 | גנוז |  |
| 5:1.260 | בכתר | . |
| 5:1.261 | וראוי |  |
| 5:1.262 | שתדע |  |
| 5:1.263 | כי |  |
| 5:1.264 | יש |  |
| 5:1.265 | עדיין |  |
| 5:1.266 | אין |  |
| 5:1.267 | סוף |  |
| 5:1.268 | אשר |  |
| 5:1.269 | הוא |  |
| 5:1.270 | לעילא |  |
| 5:1.271 | לעילא | , |
| 5:1.272 | כי |  |
| 5:1.273 | אין |  |
| 5:1.274 | שום |  |
| 5:1.275 | מחשבה |  |
| 5:1.276 | תפיס |  |
| 5:1.277 | ביה |  |
| 5:1.278 | והוא |  |
| 5:1.279 | עמוק |  |
| 5:1.280 | מי |  |
| 5:1.281 | ימצאנו | , |
| 5:1.282 | ולכן |  |
| 5:1.283 | תמצא |  |
| 5:1.284 | בזהר |  |
| 5:1.285 | כמה |  |
| 5:1.286 | מיני |  |
| 5:1.287 | אין |  |
| 5:1.288 | סוף |  |
| 5:1.289 | בכמה |  |
| 5:1.290 | דוכתי |  |
| 5:1.291 | אשר |  |
| 5:1.292 | כמעט |  |
| 5:1.293 | יסתכן |  |
| 5:1.294 | האדם |  |
| 5:1.295 | בנפשו | , |
| 5:1.296 | יען |  |
| 5:1.297 | מי |  |
| 5:1.298 | שאין |  |
| 5:1.299 | לו |  |
| 5:1.300 | הקדמות |  |
| 5:1.301 | אינו |  |
| 5:1.302 | יכול |  |
| 5:1.303 | לשוט |  |
| 5:1.304 | בימא |  |
| 5:1.305 | רבה | . |
| 5:1.306 | ולכן |  |
| 5:1.307 | על |  |
| 5:1.308 | זה |  |
| 5:1.309 | האין |  |
| 5:1.310 | סוף |  |
| 5:1.311 | למטה |  |
| 5:1.312 | למדו |  |
| 5:1.313 | בזהר |  |
| 5:1.314 | אדרת |  |
| 5:1.315 | נשא | . |
| 5:1.316 | ואז |  |
| 5:1.317 | בהיותו |  |
| 5:1.318 | שם |  |
| 5:1.319 | בסוד |  |
| 5:1.320 | תלת |  |
| 5:1.321 | רישין |  |
| 5:1.322 | אז |  |
| 5:1.323 | הוא |  |
| 5:1.324 | נקרא |  |
| 5:1.325 | עתיקא |  |
| 5:1.326 | דכל |  |
| 5:1.327 | עתיקין | , |
| 5:1.328 | וז"ס |  |
| 5:1.329 | מה |  |
| 5:1.330 | שאמרו |  |
| 5:1.331 | האי |  |
| 5:1.332 | עתיקא |  |
| 5:1.333 | קדישא |  |
| 5:1.334 | אשתכח |  |
| 5:1.335 | בתלת |  |
| 5:1.336 | רישין | . |
| 5:1.337 | נמצא |  |
| 5:1.338 | כי |  |
| 5:1.339 | עיקר |  |
| 5:1.340 | עתיקא |  |
| 5:1.341 | קדישא |  |
| 5:1.342 | הוא |  |
| 5:1.343 | האין |  |
| 5:1.344 | סוף | , |
| 5:1.345 | אמנם |  |
| 5:1.346 | ע"ש |  |
| 5:1.347 | התלבשותו |  |
| 5:1.348 | בתלת |  |
| 5:1.349 | רישין |  |
| 5:1.350 | נקראים |  |
| 5:1.351 | התלת |  |
| 5:1.352 | רישין |  |
| 5:1.353 | עתיקא |  |
| 5:1.354 | קדישא | ", |
| 5:1.355 | עכ"ל | . |
| 5:1.356 | מתבאר |  |
| 5:1.357 | בהדיא |  |
| 5:1.358 | מדבריו |  |
| 5:1.359 | כי |  |
| 5:1.360 | הא"ס |  |
| 5:1.361 | הראשון |  |
| 5:1.362 | שהוא |  |
| 5:1.363 | עמוק |  |
| 5:1.364 | עמוק |  |
| 5:1.365 | ומקיף |  |
| 5:1.366 | את |  |
| 5:1.367 | כל |  |
| 5:1.368 | העולמות |  |
| 5:1.369 | אין |  |
| 5:1.370 | לכוון |  |
| 5:1.371 | תפילתו |  |
| 5:1.372 | אליו |  |
| 5:1.373 | כלל | , |
| 5:1.374 | כי |  |
| 5:1.375 | הוא |  |
| 5:1.376 | עמוק |  |
| 5:1.377 | עמוק |  |
| 5:1.378 | ורחוק |  |
| 5:1.379 | מאד |  |
| 5:1.380 | ולית |  |
| 5:1.381 | מחשבה |  |
| 5:1.382 | תפיסא |  |
| 5:1.383 | ביה |  |
| 5:1.384 | כלל |  |
| 5:1.385 | וכלל | . |
| 5:1.386 | וכל |  |
| 5:1.387 | תפילותנו |  |
| 5:1.388 | אל |  |
| 5:1.389 | הזעיר |  |
| 5:1.390 | אנפין |  |
| 5:1.391 | להמשיך |  |
| 5:1.392 | השפע |  |
| 5:1.393 | מאבא |  |
| 5:1.394 | ואמא |  |
| 5:1.395 | ומאין |  |
| 5:1.396 | סוף |  |
| 5:1.397 | הגנוז |  |
| 5:1.398 | בכתר |  |
| 5:1.399 | בלבד |  |
| 5:1.400 | יתגלה |  |
| 5:1.401 | לנו |  |
| 5:1.402 | הארתו | . |
| 5:1.403 | אבל |  |
| 5:1.404 | אין |  |
| 5:1.405 | סוף |  |
| 5:1.406 | הקדום |  |
| 5:1.407 | הנעלה |  |
| 5:1.408 | לית |  |
| 5:1.409 | מחשבה |  |
| 5:1.410 | תפיסה |  |
| 5:1.411 | ביה |  |
| 5:1.412 | כלל | . |
| 5:1.413 | ומי |  |
| 5:1.414 | שלא |  |
| 5:1.415 | ידע |  |
| 5:1.416 | זה |  |
| 5:1.417 | יבוא |  |
| 5:1.418 | לכוון |  |
| 5:1.419 | תפילותיו |  |
| 5:1.420 | אליו | , |
| 5:1.421 | ויסתכן |  |
| 5:1.422 | בנפשו |  |
| 5:1.423 | שלא |  |
| 5:1.424 | תקובל |  |
| 5:1.425 | תפילתו | . |
| 5:1.426 | ואדרבה |  |
| 5:1.427 | ענוש |  |
| 5:1.428 | ייענש |  |
| 5:1.429 | כמ"ש |  |
| 5:1.430 | לעיל |  |
| 5:1.431 | בשם |  |
| 5:1.432 | ספר |  |
| 5:1.433 | כסא |  |
| 5:1.434 | אליהו | . |
| 5:1.435 | ודברים |  |
| 5:1.436 | אלה |  |
| 5:1.437 | היפך |  |
| 5:1.438 | דבריך |  |
| 5:1.439 | שכתבת |  |
| 5:1.440 | בשם |  |
| 5:1.441 | לחם |  |
| 5:1.442 | שלמה | . |
| 5:1.443 | ואין |  |
| 5:1.444 | ספר |  |
| 5:1.445 | לחם |  |
| 5:1.446 | שלמה |  |
| 5:1.447 | מצוי |  |
| 5:1.448 | אצלנו |  |
| 5:1.449 | לעיין |  |
| 5:1.450 | בו |  |
| 5:1.451 | ואפשר |  |
| 5:1.452 | שאתה |  |
| 5:1.453 | מטעה |  |
| 5:1.454 | אותנו |  |
| 5:1.455 | וכותב |  |
| 5:1.456 | לנו |  |
| 5:1.457 | בשמו |  |
| 5:1.458 | היפך |  |
| 5:1.459 | כוונתו |  |
| 5:1.460 | כמו |  |
| 5:1.461 | שעשית |  |
| 5:1.462 | בכוונת |  |
| 5:1.463 | כסא |  |
| 5:1.464 | אליהו |  |
| 5:1.465 | ואינו |  |
| 5:1.466 | כן | . |
| 5:1.467 | כי |  |
| 5:1.468 | המוסכם |  |
| 5:1.469 | על |  |
| 5:1.470 | בעלי |  |
| 5:1.471 | הקבלה |  |
| 5:1.472 | החדשה |  |
| 5:1.473 | שהאין |  |
| 5:1.474 | סוף |  |
| 5:1.475 | שהוא |  |
| 5:1.476 | הסיבה |  |
| 5:1.477 | הראשונה |  |
| 5:1.478 | אין |  |
| 5:1.479 | שום |  |
| 5:1.480 | תפילה |  |
| 5:1.481 | וברכה |  |
| 5:1.482 | שייכא |  |
| 5:1.483 | ביה |  |
| 5:1.484 | לדעתם |  |
| 5:1.485 | כלל |  |
| 5:1.486 | כי |  |
| 5:1.487 | הוא |  |
| 5:1.488 | עמוק |  |
| 5:1.489 | ורחוק |  |
| 5:1.490 | מאד | . |
| 5:1.491 | ורק |  |
| 5:1.492 | כל |  |
| 5:1.493 | תפילותנו |  |
| 5:1.494 | להמשיך |  |
| 5:1.495 | מאין |  |
| 5:1.496 | סוף |  |
| 5:1.497 | הגנוז |  |
| 5:1.498 | בכתר |  |
| 5:1.499 | בלבד |  |
| 5:1.500 | כדי |  |
| 5:1.501 | שיתגלה |  |
| 5:1.502 | לנו |  |
| 5:1.503 | הארתו | . |
| 5:1.504 | ולמה |  |
| 5:1.505 | זה |  |
| 5:1.506 | לא |  |
| 5:1.507 | נכוון |  |
| 5:1.508 | בתפילותנו |  |
| 5:1.509 | אל |  |
| 5:1.510 | האין |  |
| 5:1.511 | סוף |  |
| 5:1.512 | הקדום |  |
| 5:1.513 | המקיף |  |
| 5:1.514 | את |  |
| 5:1.515 | כל |  |
| 5:1.516 | המציאות | ? |
| 5:1.517 | ולמה |  |
| 5:1.518 | יקצוף |  |
| 5:1.519 | עלינו |  |
| 5:1.520 | בעת |  |
| 5:1.521 | אשר |  |
| 5:1.522 | נקרא |  |
| 5:1.523 | אליו | ? |
| 5:1.524 | וגם |  |
| 5:1.525 | למה |  |
| 5:1.526 | לא |  |
| 5:1.527 | הודיע |  |
| 5:1.528 | כל |  |
| 5:1.529 | זאת |  |
| 5:1.530 | ע"י |  |
| 5:1.531 | נאמן |  |
| 5:1.532 | ביתו |  |
| 5:1.533 | משה |  |
| 5:1.534 | איש |  |
| 5:1.535 | האלהים | , |
| 5:1.536 | לא |  |
| 5:1.537 | בתושב"כ |  |
| 5:1.538 | ולא |  |
| 5:1.539 | בתושבע"פ |  |
| 5:1.540 | שהיא |  |
| 5:1.541 | המשנה |  |
| 5:1.542 | והתלמוד |  |
| 5:1.543 | המקובל |  |
| 5:1.544 | מפה |  |
| 5:1.545 | אל |  |
| 5:1.546 | פה | ? |
| 5:1.547 | וכן |  |
| 5:1.548 | כל |  |
| 5:1.549 | הנביאים |  |
| 5:1.550 | שעמדו |  |
| 5:1.551 | להם |  |
| 5:1.552 | לישראל |  |
| 5:1.553 | מימות |  |
| 5:1.554 | משה |  |
| 5:1.555 | ועד |  |
| 5:1.556 | האחרון |  |
| 5:1.557 | שבהם | - |
| 5:1.558 | מלאכי | - |
| 5:1.559 | המזהיר |  |
| 5:1.560 | אותנו |  |
| 5:1.561 | מלאכי |  |
| 5:1.562 | ג | ' |
| 5:1.563 | כ"ב | ): |
| 5:1.564 | זכרו |  |
| 5:1.565 | תורת |  |
| 5:1.566 | משה |  |
| 5:1.567 | עבדי |  |
| 5:1.568 | אשר |  |
| 5:1.569 | צויתי |  |
| 5:1.570 | אותו |  |
| 5:1.571 | בחורב |  |
| 5:1.572 | על |  |
| 5:1.573 | כל |  |
| 5:1.574 | ישראל |  |
| 5:1.575 | חוקים |  |
| 5:1.576 | ומשפטים | "- |
| 5:1.577 | ולא |  |
| 5:1.578 | צירופים |  |
| 5:1.579 | ותיקון |  |
| 5:1.580 | עולמות | . |
| 5:1.581 | ולמה |  |
| 5:1.582 | לא |  |
| 5:1.583 | הודיעונו |  |
| 5:1.584 | כ"ז | ? |
| 5:1.585 | וילכו |  |
| 5:1.586 | כל |  |
| 5:1.587 | הדורות |  |
| 5:1.588 | שעברו |  |
| 5:1.589 | שולל |  |
| 5:1.590 | לפי |  |
| 5:1.591 | דבריהם | ) |
| 5:1.592 | להתפלל |  |
| 5:1.593 | ח"ו |  |
| 5:1.594 | אל |  |
| 5:1.595 | אל |  |
| 5:1.596 | לא |  |
| 5:1.597 | יושיע | ? |
| 5:1.598 | היה |  |
| 5:1.599 | להם |  |
| 5:1.600 | להודיע |  |
| 5:1.601 | לישראל |  |
| 5:1.602 | כי |  |
| 5:1.603 | לא |  |
| 5:1.604 | אל |  |
| 5:1.605 | אל |  |
| 5:1.606 | עליון |  |
| 5:1.607 | נורא |  |
| 5:1.608 | יתפללו |  |
| 5:1.609 | כי |  |
| 5:1.610 | אינו |  |
| 5:1.611 | מושג | . |
| 5:1.612 | ורק |  |
| 5:1.613 | אל |  |
| 5:1.614 | קצר |  |
| 5:1.615 | האפים |  |
| 5:1.616 | הזעיר |  |
| 5:1.617 | אנפין | ) |
| 5:1.618 | שכבר |  |
| 5:1.619 | נתעבה |  |
| 5:1.620 | והושג |  |
| 5:1.621 | בחוש |  |
| 5:1.622 | יתפללו | . |
| 5:1.623 | והן |  |
| 5:1.624 | עתה |  |
| 5:1.625 | נכפלה |  |
| 5:1.626 | קושיתנו |  |
| 5:1.627 | וחזקה | . |
| 5:1.628 | כי |  |
| 5:1.629 | הרבה |  |
| 5:1.630 | מיני |  |
| 5:1.631 | אין |  |
| 5:1.632 | סוף |  |
| 5:1.633 | הם |  |
| 5:1.634 | והרבה |  |
| 5:1.635 | פרצופים |  |
| 5:1.636 | יש |  |
| 5:1.637 | שכ"א |  |
| 5:1.638 | עושה |  |
| 5:1.639 | פעולה | , |
| 5:1.640 | וזה |  |
| 5:1.641 | הוא |  |
| 5:1.642 | אמונת |  |
| 5:1.643 | הרבה |  |
| 5:1.644 | רשויות |  |
| 5:1.645 | בשמים | , |
| 5:1.646 | וכבא |  |
| 5:1.647 | הבאתי |  |
| 5:1.648 | לעיל |  |
| 5:1.649 | דברי |  |
| 5:1.650 | כסא |  |
| 5:1.651 | אליהו |  |
| 5:1.652 | להודיע |  |
| 5:1.653 | לך |  |
| 5:1.654 | טעותך |  |
| 5:1.655 | שחשבת |  |
| 5:1.656 | לישב |  |
| 5:1.657 | דבריו | . |
| 5:1.658 | וכל |  |
| 5:1.659 | דבריך |  |
| 5:1.660 | הבל |  |
| 5:1.661 | וריק |  |
| 5:1.662 | ישאם |  |
| 5:1.663 | רוח | . |
| 5:1.664 | וברור |  |
| 5:1.665 | הוא |  |
| 5:1.666 | שדעתו |  |
| 5:1.667 | שצריך |  |
| 5:1.668 | שיתפלל |  |
| 5:1.669 | לזעיר |  |
| 5:1.670 | אנפין |  |
| 5:1.671 | בשיתוף |  |
| 5:1.672 | שאר |  |
| 5:1.673 | הפרצופים |  |
| 5:1.674 | שלמעלה |  |
| 5:1.675 | ממנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

