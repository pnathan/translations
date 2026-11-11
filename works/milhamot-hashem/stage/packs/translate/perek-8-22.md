# Translation pack: milhamot-hashem perek-8-22 (Milchamot ha-Shem 8:26–8:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-22`.

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
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
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


## Milchamot ha-Shem 8:26

Hebrew (Wikisource): ואמר השי"ת וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד, ואמר על ידי הנביא ירמיה, אם יסתר איש במסתרים ואני לא אראנו נאם ה' הלא את השמים ואת הארץ אני מלא נאם ה' ונאמר מלא כל הארץ כבודו, ורבותינו אנשי כנה"ג בתיקוני התפלה אמרו כבודו והודו מלא העולם וכו'! האים על פי הדברים האלה נוכל לומר כיון שהוא ית' מלא הכוכבים ומסלותיהם, והשמש והירח בכל חלקיהם בפנים ובחוץ ובארץ ובימים ובנהרות ובגיאות ובמערות ובחוחים בפנים בפנים ובחוץ! הנוכל לומר על שום כוכב מהם הוא אלהינו, או על השמש והירח והארץ אשר הוא מלא אותם המה אלהינו ונצרפם יחד עם כבודו ית' השוכן בהם? לא! ולא! ואיך נומר על אבא ואימא קודשא בריך הוא ושכינתה דלעילא, ועל זעיר אנפין ונוקביה קודשא בריך הוא ושכינתיה דלתתא, ונאמר הוא ה' אלהינו והוא המלך המהולל בכל התושבחות, ובעבור מה נקרא להצורות והפרצופים שהמציא המסית מחבר הזהר בשם הוי'ה ואדנות ושאר השמות. יענה המקובל החדש בעבור שהאלהים העליון ממלא את הפרצופים בפנים ובחוץ ובין כל פרצוף ופרצוף! אף אתה אמור לו הלא הקב"ה מכריז על ידי נביאו הלא את השמים ואת הארץ אני מלא נאם ה' ועם כל זה הוא מזהיר אותנו בפירוש בתורה שלא נעבוד שום נברא זולתו ולא נשתף עמו שום דבר אחר לא מן העליונים ולא מן התחתונים וכמאמר רשב"י הנז' לעיל וכמו שנת' לעיל סי' נ"ח וסי' ס"ד.

| id | word | punctuation after |
|---|---|---|
| 8:26.1 | ואמר |  |
| 8:26.2 | השי"ת |  |
| 8:26.3 | וידעת |  |
| 8:26.4 | היום |  |
| 8:26.5 | והשבות |  |
| 8:26.6 | אל |  |
| 8:26.7 | לבבך |  |
| 8:26.8 | כי |  |
| 8:26.9 | ה | ' |
| 8:26.10 | הוא |  |
| 8:26.11 | האלהים |  |
| 8:26.12 | בשמים |  |
| 8:26.13 | ממעל |  |
| 8:26.14 | ועל |  |
| 8:26.15 | הארץ |  |
| 8:26.16 | מתחת |  |
| 8:26.17 | אין |  |
| 8:26.18 | עוד | , |
| 8:26.19 | ואמר |  |
| 8:26.20 | על |  |
| 8:26.21 | ידי |  |
| 8:26.22 | הנביא |  |
| 8:26.23 | ירמיה | , |
| 8:26.24 | אם |  |
| 8:26.25 | יסתר |  |
| 8:26.26 | איש |  |
| 8:26.27 | במסתרים |  |
| 8:26.28 | ואני |  |
| 8:26.29 | לא |  |
| 8:26.30 | אראנו |  |
| 8:26.31 | נאם |  |
| 8:26.32 | ה | ' |
| 8:26.33 | הלא |  |
| 8:26.34 | את |  |
| 8:26.35 | השמים |  |
| 8:26.36 | ואת |  |
| 8:26.37 | הארץ |  |
| 8:26.38 | אני |  |
| 8:26.39 | מלא |  |
| 8:26.40 | נאם |  |
| 8:26.41 | ה | ' |
| 8:26.42 | ונאמר |  |
| 8:26.43 | מלא |  |
| 8:26.44 | כל |  |
| 8:26.45 | הארץ |  |
| 8:26.46 | כבודו | , |
| 8:26.47 | ורבותינו |  |
| 8:26.48 | אנשי |  |
| 8:26.49 | כנה"ג |  |
| 8:26.50 | בתיקוני |  |
| 8:26.51 | התפלה |  |
| 8:26.52 | אמרו |  |
| 8:26.53 | כבודו |  |
| 8:26.54 | והודו |  |
| 8:26.55 | מלא |  |
| 8:26.56 | העולם |  |
| 8:26.57 | וכו | '! |
| 8:26.58 | האים |  |
| 8:26.59 | על |  |
| 8:26.60 | פי |  |
| 8:26.61 | הדברים |  |
| 8:26.62 | האלה |  |
| 8:26.63 | נוכל |  |
| 8:26.64 | לומר |  |
| 8:26.65 | כיון |  |
| 8:26.66 | שהוא |  |
| 8:26.67 | ית | ' |
| 8:26.68 | מלא |  |
| 8:26.69 | הכוכבים |  |
| 8:26.70 | ומסלותיהם | , |
| 8:26.71 | והשמש |  |
| 8:26.72 | והירח |  |
| 8:26.73 | בכל |  |
| 8:26.74 | חלקיהם |  |
| 8:26.75 | בפנים |  |
| 8:26.76 | ובחוץ |  |
| 8:26.77 | ובארץ |  |
| 8:26.78 | ובימים |  |
| 8:26.79 | ובנהרות |  |
| 8:26.80 | ובגיאות |  |
| 8:26.81 | ובמערות |  |
| 8:26.82 | ובחוחים |  |
| 8:26.83 | בפנים |  |
| 8:26.84 | בפנים |  |
| 8:26.85 | ובחוץ | ! |
| 8:26.86 | הנוכל |  |
| 8:26.87 | לומר |  |
| 8:26.88 | על |  |
| 8:26.89 | שום |  |
| 8:26.90 | כוכב |  |
| 8:26.91 | מהם |  |
| 8:26.92 | הוא |  |
| 8:26.93 | אלהינו | , |
| 8:26.94 | או |  |
| 8:26.95 | על |  |
| 8:26.96 | השמש |  |
| 8:26.97 | והירח |  |
| 8:26.98 | והארץ |  |
| 8:26.99 | אשר |  |
| 8:26.100 | הוא |  |
| 8:26.101 | מלא |  |
| 8:26.102 | אותם |  |
| 8:26.103 | המה |  |
| 8:26.104 | אלהינו |  |
| 8:26.105 | ונצרפם |  |
| 8:26.106 | יחד |  |
| 8:26.107 | עם |  |
| 8:26.108 | כבודו |  |
| 8:26.109 | ית | ' |
| 8:26.110 | השוכן |  |
| 8:26.111 | בהם | ? |
| 8:26.112 | לא | ! |
| 8:26.113 | ולא | ! |
| 8:26.114 | ואיך |  |
| 8:26.115 | נומר |  |
| 8:26.116 | על |  |
| 8:26.117 | אבא |  |
| 8:26.118 | ואימא |  |
| 8:26.119 | קודשא |  |
| 8:26.120 | בריך |  |
| 8:26.121 | הוא |  |
| 8:26.122 | ושכינתה |  |
| 8:26.123 | דלעילא | , |
| 8:26.124 | ועל |  |
| 8:26.125 | זעיר |  |
| 8:26.126 | אנפין |  |
| 8:26.127 | ונוקביה |  |
| 8:26.128 | קודשא |  |
| 8:26.129 | בריך |  |
| 8:26.130 | הוא |  |
| 8:26.131 | ושכינתיה |  |
| 8:26.132 | דלתתא | , |
| 8:26.133 | ונאמר |  |
| 8:26.134 | הוא |  |
| 8:26.135 | ה | ' |
| 8:26.136 | אלהינו |  |
| 8:26.137 | והוא |  |
| 8:26.138 | המלך |  |
| 8:26.139 | המהולל |  |
| 8:26.140 | בכל |  |
| 8:26.141 | התושבחות | , |
| 8:26.142 | ובעבור |  |
| 8:26.143 | מה |  |
| 8:26.144 | נקרא |  |
| 8:26.145 | להצורות |  |
| 8:26.146 | והפרצופים |  |
| 8:26.147 | שהמציא |  |
| 8:26.148 | המסית |  |
| 8:26.149 | מחבר |  |
| 8:26.150 | הזהר |  |
| 8:26.151 | בשם |  |
| 8:26.152 | הוי'ה |  |
| 8:26.153 | ואדנות |  |
| 8:26.154 | ושאר |  |
| 8:26.155 | השמות | . |
| 8:26.156 | יענה |  |
| 8:26.157 | המקובל |  |
| 8:26.158 | החדש |  |
| 8:26.159 | בעבור |  |
| 8:26.160 | שהאלהים |  |
| 8:26.161 | העליון |  |
| 8:26.162 | ממלא |  |
| 8:26.163 | את |  |
| 8:26.164 | הפרצופים |  |
| 8:26.165 | בפנים |  |
| 8:26.166 | ובחוץ |  |
| 8:26.167 | ובין |  |
| 8:26.168 | כל |  |
| 8:26.169 | פרצוף |  |
| 8:26.170 | ופרצוף | ! |
| 8:26.171 | אף |  |
| 8:26.172 | אתה |  |
| 8:26.173 | אמור |  |
| 8:26.174 | לו |  |
| 8:26.175 | הלא |  |
| 8:26.176 | הקב"ה |  |
| 8:26.177 | מכריז |  |
| 8:26.178 | על |  |
| 8:26.179 | ידי |  |
| 8:26.180 | נביאו |  |
| 8:26.181 | הלא |  |
| 8:26.182 | את |  |
| 8:26.183 | השמים |  |
| 8:26.184 | ואת |  |
| 8:26.185 | הארץ |  |
| 8:26.186 | אני |  |
| 8:26.187 | מלא |  |
| 8:26.188 | נאם |  |
| 8:26.189 | ה | ' |
| 8:26.190 | ועם |  |
| 8:26.191 | כל |  |
| 8:26.192 | זה |  |
| 8:26.193 | הוא |  |
| 8:26.194 | מזהיר |  |
| 8:26.195 | אותנו |  |
| 8:26.196 | בפירוש |  |
| 8:26.197 | בתורה |  |
| 8:26.198 | שלא |  |
| 8:26.199 | נעבוד |  |
| 8:26.200 | שום |  |
| 8:26.201 | נברא |  |
| 8:26.202 | זולתו |  |
| 8:26.203 | ולא |  |
| 8:26.204 | נשתף |  |
| 8:26.205 | עמו |  |
| 8:26.206 | שום |  |
| 8:26.207 | דבר |  |
| 8:26.208 | אחר |  |
| 8:26.209 | לא |  |
| 8:26.210 | מן |  |
| 8:26.211 | העליונים |  |
| 8:26.212 | ולא |  |
| 8:26.213 | מן |  |
| 8:26.214 | התחתונים |  |
| 8:26.215 | וכמאמר |  |
| 8:26.216 | רשב"י |  |
| 8:26.217 | הנז | ' |
| 8:26.218 | לעיל |  |
| 8:26.219 | וכמו |  |
| 8:26.220 | שנת | ' |
| 8:26.221 | לעיל |  |
| 8:26.222 | סי | ' |
| 8:26.223 | נ"ח |  |
| 8:26.224 | וסי | ' |
| 8:26.225 | ס"ד | . |


## Milchamot ha-Shem 8:27

Hebrew (Wikisource): קיט) ואם תאמר שאני שאר הנבראים התחתונים והשפלים שהם רחוקים מהבורא שאין להם דין אלהות! אבל הקרובים לאלוה העליון שהוא הסבה העליונה יש להם דין אלהות! והלא כבר נאמר האלהי מקרוב אני נאם ה' ולא אלהי מרחוק, יאמר שהוא ית' מושל בעליונים ובתחתונים, וכולם נסקרים בסקירה אחת באין הפרש והבדל ביניהם כלל, ונאמר וכבודי לאחר לא אתן, שלא מסר עולמו לשום נברא לעשות בעולם כחפצו, רק הוא אמר ויהי אור צוה ויעמוד, ולא ישנו את תפקידתם אשר תוה עליהם, וכמו שאמרו רז"ל בפ' קמא דפסחים קרייה רחמנא לנהורא ופקדיה אמצותא דיממא, וקרייה לחשוא ופקדיה אמצותא דליליא, אבל אי לשום אחד מכל הנבראים דין אלהות כלל! ואע"פ שהם מלאים מכבודו, ובין שהם קרובים ובין שהם רחוקים, שהרי און שפל ורחוק כמונו יושבי חלד, ועם כל זה אמר הכתוב מי גוי גדול אשר לו אלהים קרובים אליו כה' אלהינו! שהוא קרוב בכל מיני קריבות כמו שאמרו רז"ל שאדם עומד אחורי העמוד בבית הכנסת ומתפלל בלחש, והקב"ה מאזין. כפה לאזן כדלעיל, וכן אמר נעים זמירות ישראל מי כה' אלהינו המגביהי לשבת, המשפילי לראות בשמים ובארץ, כקרוב כרחוק, הכל גלוי לפניו, והכל עבדיו לעשות רצונו וחפצו ובידו לשנוהן כחפצו ורצונו, ואין להם דין אלהות כלל, כי לעולם לא נתן ולא יתן כבודו לאחרים, שנאמר וכבודי לאחר לא אתן. והמופת הגדול לנו יציאת מצרים וקריעת ים סוף וכל המופתים שנעשו שמה. ולא מחשבותינו מחשבותיו ולא דרכינו דרכיו. כי לא נוכל להשיג מהותו ואיכותו ואופן חנייתו כמו שנתבאר לעיל סי' נ"ב ונ"ג בשם הרוקח.

| id | word | punctuation after |
|---|---|---|
| 8:27.1 | קיט | ) |
| 8:27.2 | ואם |  |
| 8:27.3 | תאמר |  |
| 8:27.4 | שאני |  |
| 8:27.5 | שאר |  |
| 8:27.6 | הנבראים |  |
| 8:27.7 | התחתונים |  |
| 8:27.8 | והשפלים |  |
| 8:27.9 | שהם |  |
| 8:27.10 | רחוקים |  |
| 8:27.11 | מהבורא |  |
| 8:27.12 | שאין |  |
| 8:27.13 | להם |  |
| 8:27.14 | דין |  |
| 8:27.15 | אלהות | ! |
| 8:27.16 | אבל |  |
| 8:27.17 | הקרובים |  |
| 8:27.18 | לאלוה |  |
| 8:27.19 | העליון |  |
| 8:27.20 | שהוא |  |
| 8:27.21 | הסבה |  |
| 8:27.22 | העליונה |  |
| 8:27.23 | יש |  |
| 8:27.24 | להם |  |
| 8:27.25 | דין |  |
| 8:27.26 | אלהות | ! |
| 8:27.27 | והלא |  |
| 8:27.28 | כבר |  |
| 8:27.29 | נאמר |  |
| 8:27.30 | האלהי |  |
| 8:27.31 | מקרוב |  |
| 8:27.32 | אני |  |
| 8:27.33 | נאם |  |
| 8:27.34 | ה | ' |
| 8:27.35 | ולא |  |
| 8:27.36 | אלהי |  |
| 8:27.37 | מרחוק | , |
| 8:27.38 | יאמר |  |
| 8:27.39 | שהוא |  |
| 8:27.40 | ית | ' |
| 8:27.41 | מושל |  |
| 8:27.42 | בעליונים |  |
| 8:27.43 | ובתחתונים | , |
| 8:27.44 | וכולם |  |
| 8:27.45 | נסקרים |  |
| 8:27.46 | בסקירה |  |
| 8:27.47 | אחת |  |
| 8:27.48 | באין |  |
| 8:27.49 | הפרש |  |
| 8:27.50 | והבדל |  |
| 8:27.51 | ביניהם |  |
| 8:27.52 | כלל | , |
| 8:27.53 | ונאמר |  |
| 8:27.54 | וכבודי |  |
| 8:27.55 | לאחר |  |
| 8:27.56 | לא |  |
| 8:27.57 | אתן | , |
| 8:27.58 | שלא |  |
| 8:27.59 | מסר |  |
| 8:27.60 | עולמו |  |
| 8:27.61 | לשום |  |
| 8:27.62 | נברא |  |
| 8:27.63 | לעשות |  |
| 8:27.64 | בעולם |  |
| 8:27.65 | כחפצו | , |
| 8:27.66 | רק |  |
| 8:27.67 | הוא |  |
| 8:27.68 | אמר |  |
| 8:27.69 | ויהי |  |
| 8:27.70 | אור |  |
| 8:27.71 | צוה |  |
| 8:27.72 | ויעמוד | , |
| 8:27.73 | ולא |  |
| 8:27.74 | ישנו |  |
| 8:27.75 | את |  |
| 8:27.76 | תפקידתם |  |
| 8:27.77 | אשר |  |
| 8:27.78 | תוה |  |
| 8:27.79 | עליהם | , |
| 8:27.80 | וכמו |  |
| 8:27.81 | שאמרו |  |
| 8:27.82 | רז"ל |  |
| 8:27.83 | בפ | ' |
| 8:27.84 | קמא |  |
| 8:27.85 | דפסחים |  |
| 8:27.86 | קרייה |  |
| 8:27.87 | רחמנא |  |
| 8:27.88 | לנהורא |  |
| 8:27.89 | ופקדיה |  |
| 8:27.90 | אמצותא |  |
| 8:27.91 | דיממא | , |
| 8:27.92 | וקרייה |  |
| 8:27.93 | לחשוא |  |
| 8:27.94 | ופקדיה |  |
| 8:27.95 | אמצותא |  |
| 8:27.96 | דליליא | , |
| 8:27.97 | אבל |  |
| 8:27.98 | אי |  |
| 8:27.99 | לשום |  |
| 8:27.100 | אחד |  |
| 8:27.101 | מכל |  |
| 8:27.102 | הנבראים |  |
| 8:27.103 | דין |  |
| 8:27.104 | אלהות |  |
| 8:27.105 | כלל | ! |
| 8:27.106 | ואע"פ |  |
| 8:27.107 | שהם |  |
| 8:27.108 | מלאים |  |
| 8:27.109 | מכבודו | , |
| 8:27.110 | ובין |  |
| 8:27.111 | שהם |  |
| 8:27.112 | קרובים |  |
| 8:27.113 | ובין |  |
| 8:27.114 | שהם |  |
| 8:27.115 | רחוקים | , |
| 8:27.116 | שהרי |  |
| 8:27.117 | און |  |
| 8:27.118 | שפל |  |
| 8:27.119 | ורחוק |  |
| 8:27.120 | כמונו |  |
| 8:27.121 | יושבי |  |
| 8:27.122 | חלד | , |
| 8:27.123 | ועם |  |
| 8:27.124 | כל |  |
| 8:27.125 | זה |  |
| 8:27.126 | אמר |  |
| 8:27.127 | הכתוב |  |
| 8:27.128 | מי |  |
| 8:27.129 | גוי |  |
| 8:27.130 | גדול |  |
| 8:27.131 | אשר |  |
| 8:27.132 | לו |  |
| 8:27.133 | אלהים |  |
| 8:27.134 | קרובים |  |
| 8:27.135 | אליו |  |
| 8:27.136 | כה | ' |
| 8:27.137 | אלהינו | ! |
| 8:27.138 | שהוא |  |
| 8:27.139 | קרוב |  |
| 8:27.140 | בכל |  |
| 8:27.141 | מיני |  |
| 8:27.142 | קריבות |  |
| 8:27.143 | כמו |  |
| 8:27.144 | שאמרו |  |
| 8:27.145 | רז"ל |  |
| 8:27.146 | שאדם |  |
| 8:27.147 | עומד |  |
| 8:27.148 | אחורי |  |
| 8:27.149 | העמוד |  |
| 8:27.150 | בבית |  |
| 8:27.151 | הכנסת |  |
| 8:27.152 | ומתפלל |  |
| 8:27.153 | בלחש | , |
| 8:27.154 | והקב"ה |  |
| 8:27.155 | מאזין | . |
| 8:27.156 | כפה |  |
| 8:27.157 | לאזן |  |
| 8:27.158 | כדלעיל | , |
| 8:27.159 | וכן |  |
| 8:27.160 | אמר |  |
| 8:27.161 | נעים |  |
| 8:27.162 | זמירות |  |
| 8:27.163 | ישראל |  |
| 8:27.164 | מי |  |
| 8:27.165 | כה | ' |
| 8:27.166 | אלהינו |  |
| 8:27.167 | המגביהי |  |
| 8:27.168 | לשבת | , |
| 8:27.169 | המשפילי |  |
| 8:27.170 | לראות |  |
| 8:27.171 | בשמים |  |
| 8:27.172 | ובארץ | , |
| 8:27.173 | כקרוב |  |
| 8:27.174 | כרחוק | , |
| 8:27.175 | הכל |  |
| 8:27.176 | גלוי |  |
| 8:27.177 | לפניו | , |
| 8:27.178 | והכל |  |
| 8:27.179 | עבדיו |  |
| 8:27.180 | לעשות |  |
| 8:27.181 | רצונו |  |
| 8:27.182 | וחפצו |  |
| 8:27.183 | ובידו |  |
| 8:27.184 | לשנוהן |  |
| 8:27.185 | כחפצו |  |
| 8:27.186 | ורצונו | , |
| 8:27.187 | ואין |  |
| 8:27.188 | להם |  |
| 8:27.189 | דין |  |
| 8:27.190 | אלהות |  |
| 8:27.191 | כלל | , |
| 8:27.192 | כי |  |
| 8:27.193 | לעולם |  |
| 8:27.194 | לא |  |
| 8:27.195 | נתן |  |
| 8:27.196 | ולא |  |
| 8:27.197 | יתן |  |
| 8:27.198 | כבודו |  |
| 8:27.199 | לאחרים | , |
| 8:27.200 | שנאמר |  |
| 8:27.201 | וכבודי |  |
| 8:27.202 | לאחר |  |
| 8:27.203 | לא |  |
| 8:27.204 | אתן | . |
| 8:27.205 | והמופת |  |
| 8:27.206 | הגדול |  |
| 8:27.207 | לנו |  |
| 8:27.208 | יציאת |  |
| 8:27.209 | מצרים |  |
| 8:27.210 | וקריעת |  |
| 8:27.211 | ים |  |
| 8:27.212 | סוף |  |
| 8:27.213 | וכל |  |
| 8:27.214 | המופתים |  |
| 8:27.215 | שנעשו |  |
| 8:27.216 | שמה | . |
| 8:27.217 | ולא |  |
| 8:27.218 | מחשבותינו |  |
| 8:27.219 | מחשבותיו |  |
| 8:27.220 | ולא |  |
| 8:27.221 | דרכינו |  |
| 8:27.222 | דרכיו | . |
| 8:27.223 | כי |  |
| 8:27.224 | לא |  |
| 8:27.225 | נוכל |  |
| 8:27.226 | להשיג |  |
| 8:27.227 | מהותו |  |
| 8:27.228 | ואיכותו |  |
| 8:27.229 | ואופן |  |
| 8:27.230 | חנייתו |  |
| 8:27.231 | כמו |  |
| 8:27.232 | שנתבאר |  |
| 8:27.233 | לעיל |  |
| 8:27.234 | סי | ' |
| 8:27.235 | נ"ב |  |
| 8:27.236 | ונ"ג |  |
| 8:27.237 | בשם |  |
| 8:27.238 | הרוקח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

