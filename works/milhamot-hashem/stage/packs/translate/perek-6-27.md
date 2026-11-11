# Translation pack: milhamot-hashem perek-6-27 (Milchamot ha-Shem 6:41–6:41)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-27.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-27`.

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
- דברי תימה: "astonishing words" (words that astonish).
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
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.

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


## Milchamot ha-Shem 6:41

Hebrew (Wikisource): עט) ולא די שעשו כן לחשוב כמה אלוהות בעולם האצילות, אלא שגם בעולם הבריאה, ועולם היצירה, ועולם העשייה. המציאו ברעיונותם כל הפרצופים הנז' אצלם בעולם האצילות, ואמר שהאין סוף המתלבש בהם לאו איהו חד בהון כמו בעולם האצילות כמו שכתב הר' כסא אליהו דף כ"א ע"א. ואע"פ שאמרו שהפרצופים שבבי"ע אין האין סוף בהם א' עכ"ז הם משבחים ומפארים להם, ומקבלים עליהם אלהותם כמבואר בסידור חסד חאברהם באמרם אין כאלהינו (מלכות דעשיה) אין כאדונינו (זעיר דעשייה) אין כמלכינו (אימא דעשייה) אין כמושיענו (אבא ואריך דעשייה) מי כאלהינו (מלכות דיצירה) מי כאדונינו (זעיר כיצירה) מי כמלכינו (אימא דיצירה) מי כמושיענו (אבא ואריך דיצירה) נודה לאלהינו (מלכות דבריאה) נודה לאדונינו (סעיר דבריאה) נודה למלכינו (אימא דבריאה) נודה למושיענו (אבא ואריך אנפין דבריאה) ברוך אלהינו (מלכות דאצילות) ברוך אדונינו (זעיר דאצילות) ברוך מלכינו (אימא דאצילות) ברןך מושעינו (אבא ואריך אנפין דאצילות) אצ ההוא אלהינו (כתר דאריך אנפין) אתה הוא אדונינו (ראש עתיק יומין) אתה הוא מלכינו (אדם קדמון) אתה הוא מושיענו (אור פנימי ואור מקיף המתלבש בו) כמו שיראה המעיין כל זה בסידור חסד לאברהם. הרי לך בהדיא שהם מקבלים עליהם מלכות ואדנות ואלהות כל פרצוף ופרצוף. אפי' פרצופים שבעולמות הפירוד לפי דבריהם שהם בריאה, יצירה ועשייה. ופליאה גדולה אנו מתפלאים עליכם שאתם נלחמים עמנו בכל המלחמה הזאת לבזות ולרדוף ולירד עמנו עד לחיינו. ואנו סובלים כל חרפות ובזוים לכבוד אבינו מלכינו יחיד ומיוחד מכל האחדים, בעבור מה? לא בעבור שנקבל ונודה לאמונתכם הלזו ש]ירשנו ממנה ולא חירפנו ולא בזינו אותה? ועם כל זה אתם מסתירים ומכסים את אמונתכם שלא לפרסמה. ומראים בתשובתכם כאלו אינכם מודים בה כעין גנב שמרה נשף לאמר לא תשורני עין. וסתר פנים תשימו לה! לכן הוכרחנו להביא כל זה להוכיח את הכחשות שאתם מכחישים אותנו בתשובתכם ואומרים ומי הוא שאמר שהאין סוף מתחלק? כל זה הוצרכנו להעתיק מספרי המקובלים החדשים. ולהביא מאמרי רז"ל נגדן להוכיח טעותכם או רמאותכם שאתם מסתירים איכות אמונתכם ומכסים אותה בטלית. ומכחישים את הידוע, ואינכם מבקשים האמת. ולמה זה תבושו לומר האמת שכן קבלנו לעבוד אלוהות רבים נבראים רק שאתם רודפים אחר הכבוד המדומה. לקנות שם נגד ההמון אשר לא ידעו ולא יבינו, בחשכה יתהלכו, ולשבח את עצמכם נגדם כאלו אתם נלחמים עם ה' ועם תורתו.

| id | word | punctuation after |
|---|---|---|
| 6:41.1 | עט | ) |
| 6:41.2 | ולא |  |
| 6:41.3 | די |  |
| 6:41.4 | שעשו |  |
| 6:41.5 | כן |  |
| 6:41.6 | לחשוב |  |
| 6:41.7 | כמה |  |
| 6:41.8 | אלוהות |  |
| 6:41.9 | בעולם |  |
| 6:41.10 | האצילות | , |
| 6:41.11 | אלא |  |
| 6:41.12 | שגם |  |
| 6:41.13 | בעולם |  |
| 6:41.14 | הבריאה | , |
| 6:41.15 | ועולם |  |
| 6:41.16 | היצירה | , |
| 6:41.17 | ועולם |  |
| 6:41.18 | העשייה | . |
| 6:41.19 | המציאו |  |
| 6:41.20 | ברעיונותם |  |
| 6:41.21 | כל |  |
| 6:41.22 | הפרצופים |  |
| 6:41.23 | הנז | ' |
| 6:41.24 | אצלם |  |
| 6:41.25 | בעולם |  |
| 6:41.26 | האצילות | , |
| 6:41.27 | ואמר |  |
| 6:41.28 | שהאין |  |
| 6:41.29 | סוף |  |
| 6:41.30 | המתלבש |  |
| 6:41.31 | בהם |  |
| 6:41.32 | לאו |  |
| 6:41.33 | איהו |  |
| 6:41.34 | חד |  |
| 6:41.35 | בהון |  |
| 6:41.36 | כמו |  |
| 6:41.37 | בעולם |  |
| 6:41.38 | האצילות |  |
| 6:41.39 | כמו |  |
| 6:41.40 | שכתב |  |
| 6:41.41 | הר | ' |
| 6:41.42 | כסא |  |
| 6:41.43 | אליהו |  |
| 6:41.44 | דף |  |
| 6:41.45 | כ"א |  |
| 6:41.46 | ע"א | . |
| 6:41.47 | ואע"פ |  |
| 6:41.48 | שאמרו |  |
| 6:41.49 | שהפרצופים |  |
| 6:41.50 | שבבי"ע |  |
| 6:41.51 | אין |  |
| 6:41.52 | האין |  |
| 6:41.53 | סוף |  |
| 6:41.54 | בהם |  |
| 6:41.55 | א | ' |
| 6:41.56 | עכ"ז |  |
| 6:41.57 | הם |  |
| 6:41.58 | משבחים |  |
| 6:41.59 | ומפארים |  |
| 6:41.60 | להם | , |
| 6:41.61 | ומקבלים |  |
| 6:41.62 | עליהם |  |
| 6:41.63 | אלהותם |  |
| 6:41.64 | כמבואר |  |
| 6:41.65 | בסידור |  |
| 6:41.66 | חסד |  |
| 6:41.67 | חאברהם |  |
| 6:41.68 | באמרם |  |
| 6:41.69 | אין |  |
| 6:41.70 | כאלהינו |  |
| 6:41.71 | מלכות |  |
| 6:41.72 | דעשיה | ) |
| 6:41.73 | אין |  |
| 6:41.74 | כאדונינו |  |
| 6:41.75 | זעיר |  |
| 6:41.76 | דעשייה | ) |
| 6:41.77 | אין |  |
| 6:41.78 | כמלכינו |  |
| 6:41.79 | אימא |  |
| 6:41.80 | דעשייה | ) |
| 6:41.81 | אין |  |
| 6:41.82 | כמושיענו |  |
| 6:41.83 | אבא |  |
| 6:41.84 | ואריך |  |
| 6:41.85 | דעשייה | ) |
| 6:41.86 | מי |  |
| 6:41.87 | כאלהינו |  |
| 6:41.88 | מלכות |  |
| 6:41.89 | דיצירה | ) |
| 6:41.90 | מי |  |
| 6:41.91 | כאדונינו |  |
| 6:41.92 | זעיר |  |
| 6:41.93 | כיצירה | ) |
| 6:41.94 | מי |  |
| 6:41.95 | כמלכינו |  |
| 6:41.96 | אימא |  |
| 6:41.97 | דיצירה | ) |
| 6:41.98 | מי |  |
| 6:41.99 | כמושיענו |  |
| 6:41.100 | אבא |  |
| 6:41.101 | ואריך |  |
| 6:41.102 | דיצירה | ) |
| 6:41.103 | נודה |  |
| 6:41.104 | לאלהינו |  |
| 6:41.105 | מלכות |  |
| 6:41.106 | דבריאה | ) |
| 6:41.107 | נודה |  |
| 6:41.108 | לאדונינו |  |
| 6:41.109 | סעיר |  |
| 6:41.110 | דבריאה | ) |
| 6:41.111 | נודה |  |
| 6:41.112 | למלכינו |  |
| 6:41.113 | אימא |  |
| 6:41.114 | דבריאה | ) |
| 6:41.115 | נודה |  |
| 6:41.116 | למושיענו |  |
| 6:41.117 | אבא |  |
| 6:41.118 | ואריך |  |
| 6:41.119 | אנפין |  |
| 6:41.120 | דבריאה | ) |
| 6:41.121 | ברוך |  |
| 6:41.122 | אלהינו |  |
| 6:41.123 | מלכות |  |
| 6:41.124 | דאצילות | ) |
| 6:41.125 | ברוך |  |
| 6:41.126 | אדונינו |  |
| 6:41.127 | זעיר |  |
| 6:41.128 | דאצילות | ) |
| 6:41.129 | ברוך |  |
| 6:41.130 | מלכינו |  |
| 6:41.131 | אימא |  |
| 6:41.132 | דאצילות | ) |
| 6:41.133 | ברןך |  |
| 6:41.134 | מושעינו |  |
| 6:41.135 | אבא |  |
| 6:41.136 | ואריך |  |
| 6:41.137 | אנפין |  |
| 6:41.138 | דאצילות | ) |
| 6:41.139 | אצ |  |
| 6:41.140 | ההוא |  |
| 6:41.141 | אלהינו |  |
| 6:41.142 | כתר |  |
| 6:41.143 | דאריך |  |
| 6:41.144 | אנפין | ) |
| 6:41.145 | אתה |  |
| 6:41.146 | הוא |  |
| 6:41.147 | אדונינו |  |
| 6:41.148 | ראש |  |
| 6:41.149 | עתיק |  |
| 6:41.150 | יומין | ) |
| 6:41.151 | אתה |  |
| 6:41.152 | הוא |  |
| 6:41.153 | מלכינו |  |
| 6:41.154 | אדם |  |
| 6:41.155 | קדמון | ) |
| 6:41.156 | אתה |  |
| 6:41.157 | הוא |  |
| 6:41.158 | מושיענו |  |
| 6:41.159 | אור |  |
| 6:41.160 | פנימי |  |
| 6:41.161 | ואור |  |
| 6:41.162 | מקיף |  |
| 6:41.163 | המתלבש |  |
| 6:41.164 | בו | ) |
| 6:41.165 | כמו |  |
| 6:41.166 | שיראה |  |
| 6:41.167 | המעיין |  |
| 6:41.168 | כל |  |
| 6:41.169 | זה |  |
| 6:41.170 | בסידור |  |
| 6:41.171 | חסד |  |
| 6:41.172 | לאברהם | . |
| 6:41.173 | הרי |  |
| 6:41.174 | לך |  |
| 6:41.175 | בהדיא |  |
| 6:41.176 | שהם |  |
| 6:41.177 | מקבלים |  |
| 6:41.178 | עליהם |  |
| 6:41.179 | מלכות |  |
| 6:41.180 | ואדנות |  |
| 6:41.181 | ואלהות |  |
| 6:41.182 | כל |  |
| 6:41.183 | פרצוף |  |
| 6:41.184 | ופרצוף | . |
| 6:41.185 | אפי | ' |
| 6:41.186 | פרצופים |  |
| 6:41.187 | שבעולמות |  |
| 6:41.188 | הפירוד |  |
| 6:41.189 | לפי |  |
| 6:41.190 | דבריהם |  |
| 6:41.191 | שהם |  |
| 6:41.192 | בריאה | , |
| 6:41.193 | יצירה |  |
| 6:41.194 | ועשייה | . |
| 6:41.195 | ופליאה |  |
| 6:41.196 | גדולה |  |
| 6:41.197 | אנו |  |
| 6:41.198 | מתפלאים |  |
| 6:41.199 | עליכם |  |
| 6:41.200 | שאתם |  |
| 6:41.201 | נלחמים |  |
| 6:41.202 | עמנו |  |
| 6:41.203 | בכל |  |
| 6:41.204 | המלחמה |  |
| 6:41.205 | הזאת |  |
| 6:41.206 | לבזות |  |
| 6:41.207 | ולרדוף |  |
| 6:41.208 | ולירד |  |
| 6:41.209 | עמנו |  |
| 6:41.210 | עד |  |
| 6:41.211 | לחיינו | . |
| 6:41.212 | ואנו |  |
| 6:41.213 | סובלים |  |
| 6:41.214 | כל |  |
| 6:41.215 | חרפות |  |
| 6:41.216 | ובזוים |  |
| 6:41.217 | לכבוד |  |
| 6:41.218 | אבינו |  |
| 6:41.219 | מלכינו |  |
| 6:41.220 | יחיד |  |
| 6:41.221 | ומיוחד |  |
| 6:41.222 | מכל |  |
| 6:41.223 | האחדים | , |
| 6:41.224 | בעבור |  |
| 6:41.225 | מה | ? |
| 6:41.226 | לא |  |
| 6:41.227 | בעבור |  |
| 6:41.228 | שנקבל |  |
| 6:41.229 | ונודה |  |
| 6:41.230 | לאמונתכם |  |
| 6:41.231 | הלזו |  |
| 6:41.232 | ש]ירשנו |  |
| 6:41.233 | ממנה |  |
| 6:41.234 | ולא |  |
| 6:41.235 | חירפנו |  |
| 6:41.236 | ולא |  |
| 6:41.237 | בזינו |  |
| 6:41.238 | אותה | ? |
| 6:41.239 | ועם |  |
| 6:41.240 | כל |  |
| 6:41.241 | זה |  |
| 6:41.242 | אתם |  |
| 6:41.243 | מסתירים |  |
| 6:41.244 | ומכסים |  |
| 6:41.245 | את |  |
| 6:41.246 | אמונתכם |  |
| 6:41.247 | שלא |  |
| 6:41.248 | לפרסמה | . |
| 6:41.249 | ומראים |  |
| 6:41.250 | בתשובתכם |  |
| 6:41.251 | כאלו |  |
| 6:41.252 | אינכם |  |
| 6:41.253 | מודים |  |
| 6:41.254 | בה |  |
| 6:41.255 | כעין |  |
| 6:41.256 | גנב |  |
| 6:41.257 | שמרה |  |
| 6:41.258 | נשף |  |
| 6:41.259 | לאמר |  |
| 6:41.260 | לא |  |
| 6:41.261 | תשורני |  |
| 6:41.262 | עין | . |
| 6:41.263 | וסתר |  |
| 6:41.264 | פנים |  |
| 6:41.265 | תשימו |  |
| 6:41.266 | לה | ! |
| 6:41.267 | לכן |  |
| 6:41.268 | הוכרחנו |  |
| 6:41.269 | להביא |  |
| 6:41.270 | כל |  |
| 6:41.271 | זה |  |
| 6:41.272 | להוכיח |  |
| 6:41.273 | את |  |
| 6:41.274 | הכחשות |  |
| 6:41.275 | שאתם |  |
| 6:41.276 | מכחישים |  |
| 6:41.277 | אותנו |  |
| 6:41.278 | בתשובתכם |  |
| 6:41.279 | ואומרים |  |
| 6:41.280 | ומי |  |
| 6:41.281 | הוא |  |
| 6:41.282 | שאמר |  |
| 6:41.283 | שהאין |  |
| 6:41.284 | סוף |  |
| 6:41.285 | מתחלק | ? |
| 6:41.286 | כל |  |
| 6:41.287 | זה |  |
| 6:41.288 | הוצרכנו |  |
| 6:41.289 | להעתיק |  |
| 6:41.290 | מספרי |  |
| 6:41.291 | המקובלים |  |
| 6:41.292 | החדשים | . |
| 6:41.293 | ולהביא |  |
| 6:41.294 | מאמרי |  |
| 6:41.295 | רז"ל |  |
| 6:41.296 | נגדן |  |
| 6:41.297 | להוכיח |  |
| 6:41.298 | טעותכם |  |
| 6:41.299 | או |  |
| 6:41.300 | רמאותכם |  |
| 6:41.301 | שאתם |  |
| 6:41.302 | מסתירים |  |
| 6:41.303 | איכות |  |
| 6:41.304 | אמונתכם |  |
| 6:41.305 | ומכסים |  |
| 6:41.306 | אותה |  |
| 6:41.307 | בטלית | . |
| 6:41.308 | ומכחישים |  |
| 6:41.309 | את |  |
| 6:41.310 | הידוע | , |
| 6:41.311 | ואינכם |  |
| 6:41.312 | מבקשים |  |
| 6:41.313 | האמת | . |
| 6:41.314 | ולמה |  |
| 6:41.315 | זה |  |
| 6:41.316 | תבושו |  |
| 6:41.317 | לומר |  |
| 6:41.318 | האמת |  |
| 6:41.319 | שכן |  |
| 6:41.320 | קבלנו |  |
| 6:41.321 | לעבוד |  |
| 6:41.322 | אלוהות |  |
| 6:41.323 | רבים |  |
| 6:41.324 | נבראים |  |
| 6:41.325 | רק |  |
| 6:41.326 | שאתם |  |
| 6:41.327 | רודפים |  |
| 6:41.328 | אחר |  |
| 6:41.329 | הכבוד |  |
| 6:41.330 | המדומה | . |
| 6:41.331 | לקנות |  |
| 6:41.332 | שם |  |
| 6:41.333 | נגד |  |
| 6:41.334 | ההמון |  |
| 6:41.335 | אשר |  |
| 6:41.336 | לא |  |
| 6:41.337 | ידעו |  |
| 6:41.338 | ולא |  |
| 6:41.339 | יבינו | , |
| 6:41.340 | בחשכה |  |
| 6:41.341 | יתהלכו | , |
| 6:41.342 | ולשבח |  |
| 6:41.343 | את |  |
| 6:41.344 | עצמכם |  |
| 6:41.345 | נגדם |  |
| 6:41.346 | כאלו |  |
| 6:41.347 | אתם |  |
| 6:41.348 | נלחמים |  |
| 6:41.349 | עם |  |
| 6:41.350 | ה | ' |
| 6:41.351 | ועם |  |
| 6:41.352 | תורתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

