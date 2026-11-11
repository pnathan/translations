# Translation pack: milhamot-hashem perek-6-11 (Milchamot ha-Shem 6:17–6:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-11`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- ר"ל with a named authority as subject: "means to say".
- פ' + numeral after a book title: "chapter" when the numeral exceeds the
  parashot of that book (Midrash Rabbah 143 is a section/chapter number as
  the author gives it); otherwise "parashah". Note the choice.
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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
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


## Milchamot ha-Shem 6:17

Hebrew (Wikisource): סח) ואח"כ (דף הנז"ל ע"ב) הוסיף להרבות דברים בדברי נרגן כמתלהמים, וכמתלהלה היורה זיקים לשבח סודותיו ולקרות תורתו עץ חיים היא למחזיקים בה, ואיהי מטרוניתא מלכות. והעוסקים בה בני מלכים, ואת המשנה שקבל משה מפי הגבורה ושנאה לאהרן ובניו וזקני ישראל ולכל ישראל כדגרסינן בעירובין פ' כיצד מעברין. ת"ר כיצד סדר משנה משה למד מפי הגבורה (לא מפי מטטרון) נכנס אהרן שנה לו משה פרקו, נסתלק אהרון וכו', קרא לה הפילסוף הנז' שפחה ואמר שאלו זכו ישראל היה משה נותן להם משנה שאין בה אסור וטמא ופסול, אלא הכל טהור ומותר וכשר, ועל שנתן להם משה רבינו המשנה הזאת שיש בה איסור והתר, טמא וטהור, פסול וכשר, נענש להקבר בחוצה לארץ, ולא ידע איש את קבורתו, כלומר סבת עונש קבורתו, וסבת עונש קבורתו שם, משנה זו שנתן לישראל. והעוסקים בה ומורים לעם את הדרך ילכו בה קרא להם ערב רב, לפי שהמשנה מרחקת אמונתו באלהים רבים שהוא מרמזם באותיות יהוה, ולכן קרא למשנה שפחה דשלטא על מטרוניתא שהיא גורמת להפריד בין מלכא ומטרוניתא, באמרו שם ומלכא ומטרוניתא מתפרשן מבעלה בגין כן תחת שלש רגזה ארץ, תחת עבד כי ימלך, דא איהו עבדא ידיעא (מטטרון שנתן המשנה הזאת למשה) ושפחה כי תירש גברתה דא משנה, ונבל כי ישבע לחם (מלחמה של משנה) דא ערב רב. (העוסקים במשנה) עם נבל ולא חכם, ואין יחוד לפרצופים עד דיתמחון ערב רב, העוסקים בה שהם גורמים פירוד בין המלכים ופילגשיהם. ובפרשת כי תצא הוא משבח את תורתו החדשה שהוא קורא אותה קבלה, והעוסקים בה קורא להם אדם, אבל העוסקים במשנה ובתלמוד הוא קורא להם בשם דגים ועופות השמים, והעוסקים בקבלתו למעלה מהם ושעליהם נאמר וירדו בדגת הים ובעוף השמים דאינון מארי מתניתין, ומרבה עוד דברים לגנות התנאים ואמוראים ולהפחית כבודם, ולשבח העוסקים בקבלתו החדשה בדברי חלקות היורדים חדרי בטן קלי הדעת וחלושי השכל, כמאמר שלמה ע״ה דברי נרגן כמתלהמים והם ירדו חדרי בטן, כאשר יראה המעיין המשכיל על דבר אמת הבקי בכתובים ובדברי רז"ל בתלמוד ומדרשים האמתיים. ושם בדף רע"ט גילה פנים בחוצפא יתירא בשם רעיא מהימנא לומר אית סלע ואית סלע, אית אבן ואית אבן, אית אבן דשמא דיהוה (היא תורתו החדשה המלאה פרצופים וצורות הנקראים בשם יהוה ואדנות ושאר שמות הקדש נמחקים ואין נמחקים ומרמזם בשם הוי"ה) עלה אמתר ואבנא די מחת לצלמא הות לטור רב, ואית אבן דאיהי אבן משכית, דלית תמן נביעו דחכמתא ולא דיבור אלא אבן דאיהי (חסר הפסוק בכאן ולא חפץ לפרסם כוונתו כל כך בבזוי המשנה והיה רוצה לומר אבן נגף ואבן מכשול). ראה והבן בדבריו איך הוא מחשיב לתורתו החדשה וקורא לה אבנא די מחת לצלמא, ובמלת צלמא כוונתו על המשנה והתלמוד שהוא חושב אותם כצלמי האלילים, ומבזה אותה וקורא לה אבן משכית דלית בה חכמתא, ובדבריו אלה הוא רומז לפרוש מן המשנה כמאמר הכתוב ואבן משכית לא תתנו בארצכם להשתחות עליה (פי' להשתחות וליחד השי"ת על פי המשנה המרחקת כל רבוי באלהות. רק על פי תורתו החדשה שהשי"ת מחולק לכמה פרצופים זכרים ונקבות וצריך לקשרם ולאחדם ולזווגם יחד ר"ל). ושוב חזר הפילסוף הנז' לדבריו שבראש עמוד ב' שקורא לתורתו סלע הנובע חכמה, ואת המשנה ואת התלמוד שקבל משה מסיני הוא מדמה אותן לקליפה הקשה של אגוז שכל חכמי הדורות יגעים בה ולית לון יכולת לאפקא מינה מיא אלא טיפין.

| id | word | punctuation after |
|---|---|---|
| 6:17.1 | סח | ) |
| 6:17.2 | ואח"כ |  |
| 6:17.3 | דף |  |
| 6:17.4 | הנז"ל |  |
| 6:17.5 | ע"ב | ) |
| 6:17.6 | הוסיף |  |
| 6:17.7 | להרבות |  |
| 6:17.8 | דברים |  |
| 6:17.9 | בדברי |  |
| 6:17.10 | נרגן |  |
| 6:17.11 | כמתלהמים | , |
| 6:17.12 | וכמתלהלה |  |
| 6:17.13 | היורה |  |
| 6:17.14 | זיקים |  |
| 6:17.15 | לשבח |  |
| 6:17.16 | סודותיו |  |
| 6:17.17 | ולקרות |  |
| 6:17.18 | תורתו |  |
| 6:17.19 | עץ |  |
| 6:17.20 | חיים |  |
| 6:17.21 | היא |  |
| 6:17.22 | למחזיקים |  |
| 6:17.23 | בה | , |
| 6:17.24 | ואיהי |  |
| 6:17.25 | מטרוניתא |  |
| 6:17.26 | מלכות | . |
| 6:17.27 | והעוסקים |  |
| 6:17.28 | בה |  |
| 6:17.29 | בני |  |
| 6:17.30 | מלכים | , |
| 6:17.31 | ואת |  |
| 6:17.32 | המשנה |  |
| 6:17.33 | שקבל |  |
| 6:17.34 | משה |  |
| 6:17.35 | מפי |  |
| 6:17.36 | הגבורה |  |
| 6:17.37 | ושנאה |  |
| 6:17.38 | לאהרן |  |
| 6:17.39 | ובניו |  |
| 6:17.40 | וזקני |  |
| 6:17.41 | ישראל |  |
| 6:17.42 | ולכל |  |
| 6:17.43 | ישראל |  |
| 6:17.44 | כדגרסינן |  |
| 6:17.45 | בעירובין |  |
| 6:17.46 | פ | ' |
| 6:17.47 | כיצד |  |
| 6:17.48 | מעברין | . |
| 6:17.49 | ת"ר |  |
| 6:17.50 | כיצד |  |
| 6:17.51 | סדר |  |
| 6:17.52 | משנה |  |
| 6:17.53 | משה |  |
| 6:17.54 | למד |  |
| 6:17.55 | מפי |  |
| 6:17.56 | הגבורה |  |
| 6:17.57 | לא |  |
| 6:17.58 | מפי |  |
| 6:17.59 | מטטרון | ) |
| 6:17.60 | נכנס |  |
| 6:17.61 | אהרן |  |
| 6:17.62 | שנה |  |
| 6:17.63 | לו |  |
| 6:17.64 | משה |  |
| 6:17.65 | פרקו | , |
| 6:17.66 | נסתלק |  |
| 6:17.67 | אהרון |  |
| 6:17.68 | וכו | ', |
| 6:17.69 | קרא |  |
| 6:17.70 | לה |  |
| 6:17.71 | הפילסוף |  |
| 6:17.72 | הנז | ' |
| 6:17.73 | שפחה |  |
| 6:17.74 | ואמר |  |
| 6:17.75 | שאלו |  |
| 6:17.76 | זכו |  |
| 6:17.77 | ישראל |  |
| 6:17.78 | היה |  |
| 6:17.79 | משה |  |
| 6:17.80 | נותן |  |
| 6:17.81 | להם |  |
| 6:17.82 | משנה |  |
| 6:17.83 | שאין |  |
| 6:17.84 | בה |  |
| 6:17.85 | אסור |  |
| 6:17.86 | וטמא |  |
| 6:17.87 | ופסול | , |
| 6:17.88 | אלא |  |
| 6:17.89 | הכל |  |
| 6:17.90 | טהור |  |
| 6:17.91 | ומותר |  |
| 6:17.92 | וכשר | , |
| 6:17.93 | ועל |  |
| 6:17.94 | שנתן |  |
| 6:17.95 | להם |  |
| 6:17.96 | משה |  |
| 6:17.97 | רבינו |  |
| 6:17.98 | המשנה |  |
| 6:17.99 | הזאת |  |
| 6:17.100 | שיש |  |
| 6:17.101 | בה |  |
| 6:17.102 | איסור |  |
| 6:17.103 | והתר | , |
| 6:17.104 | טמא |  |
| 6:17.105 | וטהור | , |
| 6:17.106 | פסול |  |
| 6:17.107 | וכשר | , |
| 6:17.108 | נענש |  |
| 6:17.109 | להקבר |  |
| 6:17.110 | בחוצה |  |
| 6:17.111 | לארץ | , |
| 6:17.112 | ולא |  |
| 6:17.113 | ידע |  |
| 6:17.114 | איש |  |
| 6:17.115 | את |  |
| 6:17.116 | קבורתו | , |
| 6:17.117 | כלומר |  |
| 6:17.118 | סבת |  |
| 6:17.119 | עונש |  |
| 6:17.120 | קבורתו | , |
| 6:17.121 | וסבת |  |
| 6:17.122 | עונש |  |
| 6:17.123 | קבורתו |  |
| 6:17.124 | שם | , |
| 6:17.125 | משנה |  |
| 6:17.126 | זו |  |
| 6:17.127 | שנתן |  |
| 6:17.128 | לישראל | . |
| 6:17.129 | והעוסקים |  |
| 6:17.130 | בה |  |
| 6:17.131 | ומורים |  |
| 6:17.132 | לעם |  |
| 6:17.133 | את |  |
| 6:17.134 | הדרך |  |
| 6:17.135 | ילכו |  |
| 6:17.136 | בה |  |
| 6:17.137 | קרא |  |
| 6:17.138 | להם |  |
| 6:17.139 | ערב |  |
| 6:17.140 | רב | , |
| 6:17.141 | לפי |  |
| 6:17.142 | שהמשנה |  |
| 6:17.143 | מרחקת |  |
| 6:17.144 | אמונתו |  |
| 6:17.145 | באלהים |  |
| 6:17.146 | רבים |  |
| 6:17.147 | שהוא |  |
| 6:17.148 | מרמזם |  |
| 6:17.149 | באותיות |  |
| 6:17.150 | יהוה | , |
| 6:17.151 | ולכן |  |
| 6:17.152 | קרא |  |
| 6:17.153 | למשנה |  |
| 6:17.154 | שפחה |  |
| 6:17.155 | דשלטא |  |
| 6:17.156 | על |  |
| 6:17.157 | מטרוניתא |  |
| 6:17.158 | שהיא |  |
| 6:17.159 | גורמת |  |
| 6:17.160 | להפריד |  |
| 6:17.161 | בין |  |
| 6:17.162 | מלכא |  |
| 6:17.163 | ומטרוניתא | , |
| 6:17.164 | באמרו |  |
| 6:17.165 | שם |  |
| 6:17.166 | ומלכא |  |
| 6:17.167 | ומטרוניתא |  |
| 6:17.168 | מתפרשן |  |
| 6:17.169 | מבעלה |  |
| 6:17.170 | בגין |  |
| 6:17.171 | כן |  |
| 6:17.172 | תחת |  |
| 6:17.173 | שלש |  |
| 6:17.174 | רגזה |  |
| 6:17.175 | ארץ | , |
| 6:17.176 | תחת |  |
| 6:17.177 | עבד |  |
| 6:17.178 | כי |  |
| 6:17.179 | ימלך | , |
| 6:17.180 | דא |  |
| 6:17.181 | איהו |  |
| 6:17.182 | עבדא |  |
| 6:17.183 | ידיעא |  |
| 6:17.184 | מטטרון |  |
| 6:17.185 | שנתן |  |
| 6:17.186 | המשנה |  |
| 6:17.187 | הזאת |  |
| 6:17.188 | למשה | ) |
| 6:17.189 | ושפחה |  |
| 6:17.190 | כי |  |
| 6:17.191 | תירש |  |
| 6:17.192 | גברתה |  |
| 6:17.193 | דא |  |
| 6:17.194 | משנה | , |
| 6:17.195 | ונבל |  |
| 6:17.196 | כי |  |
| 6:17.197 | ישבע |  |
| 6:17.198 | לחם |  |
| 6:17.199 | מלחמה |  |
| 6:17.200 | של |  |
| 6:17.201 | משנה | ) |
| 6:17.202 | דא |  |
| 6:17.203 | ערב |  |
| 6:17.204 | רב | . |
| 6:17.205 | העוסקים |  |
| 6:17.206 | במשנה | ) |
| 6:17.207 | עם |  |
| 6:17.208 | נבל |  |
| 6:17.209 | ולא |  |
| 6:17.210 | חכם | , |
| 6:17.211 | ואין |  |
| 6:17.212 | יחוד |  |
| 6:17.213 | לפרצופים |  |
| 6:17.214 | עד |  |
| 6:17.215 | דיתמחון |  |
| 6:17.216 | ערב |  |
| 6:17.217 | רב | , |
| 6:17.218 | העוסקים |  |
| 6:17.219 | בה |  |
| 6:17.220 | שהם |  |
| 6:17.221 | גורמים |  |
| 6:17.222 | פירוד |  |
| 6:17.223 | בין |  |
| 6:17.224 | המלכים |  |
| 6:17.225 | ופילגשיהם | . |
| 6:17.226 | ובפרשת |  |
| 6:17.227 | כי |  |
| 6:17.228 | תצא |  |
| 6:17.229 | הוא |  |
| 6:17.230 | משבח |  |
| 6:17.231 | את |  |
| 6:17.232 | תורתו |  |
| 6:17.233 | החדשה |  |
| 6:17.234 | שהוא |  |
| 6:17.235 | קורא |  |
| 6:17.236 | אותה |  |
| 6:17.237 | קבלה | , |
| 6:17.238 | והעוסקים |  |
| 6:17.239 | בה |  |
| 6:17.240 | קורא |  |
| 6:17.241 | להם |  |
| 6:17.242 | אדם | , |
| 6:17.243 | אבל |  |
| 6:17.244 | העוסקים |  |
| 6:17.245 | במשנה |  |
| 6:17.246 | ובתלמוד |  |
| 6:17.247 | הוא |  |
| 6:17.248 | קורא |  |
| 6:17.249 | להם |  |
| 6:17.250 | בשם |  |
| 6:17.251 | דגים |  |
| 6:17.252 | ועופות |  |
| 6:17.253 | השמים | , |
| 6:17.254 | והעוסקים |  |
| 6:17.255 | בקבלתו |  |
| 6:17.256 | למעלה |  |
| 6:17.257 | מהם |  |
| 6:17.258 | ושעליהם |  |
| 6:17.259 | נאמר |  |
| 6:17.260 | וירדו |  |
| 6:17.261 | בדגת |  |
| 6:17.262 | הים |  |
| 6:17.263 | ובעוף |  |
| 6:17.264 | השמים |  |
| 6:17.265 | דאינון |  |
| 6:17.266 | מארי |  |
| 6:17.267 | מתניתין | , |
| 6:17.268 | ומרבה |  |
| 6:17.269 | עוד |  |
| 6:17.270 | דברים |  |
| 6:17.271 | לגנות |  |
| 6:17.272 | התנאים |  |
| 6:17.273 | ואמוראים |  |
| 6:17.274 | ולהפחית |  |
| 6:17.275 | כבודם | , |
| 6:17.276 | ולשבח |  |
| 6:17.277 | העוסקים |  |
| 6:17.278 | בקבלתו |  |
| 6:17.279 | החדשה |  |
| 6:17.280 | בדברי |  |
| 6:17.281 | חלקות |  |
| 6:17.282 | היורדים |  |
| 6:17.283 | חדרי |  |
| 6:17.284 | בטן |  |
| 6:17.285 | קלי |  |
| 6:17.286 | הדעת |  |
| 6:17.287 | וחלושי |  |
| 6:17.288 | השכל | , |
| 6:17.289 | כמאמר |  |
| 6:17.290 | שלמה |  |
| 6:17.291 | ע״ה |  |
| 6:17.292 | דברי |  |
| 6:17.293 | נרגן |  |
| 6:17.294 | כמתלהמים |  |
| 6:17.295 | והם |  |
| 6:17.296 | ירדו |  |
| 6:17.297 | חדרי |  |
| 6:17.298 | בטן | , |
| 6:17.299 | כאשר |  |
| 6:17.300 | יראה |  |
| 6:17.301 | המעיין |  |
| 6:17.302 | המשכיל |  |
| 6:17.303 | על |  |
| 6:17.304 | דבר |  |
| 6:17.305 | אמת |  |
| 6:17.306 | הבקי |  |
| 6:17.307 | בכתובים |  |
| 6:17.308 | ובדברי |  |
| 6:17.309 | רז"ל |  |
| 6:17.310 | בתלמוד |  |
| 6:17.311 | ומדרשים |  |
| 6:17.312 | האמתיים | . |
| 6:17.313 | ושם |  |
| 6:17.314 | בדף |  |
| 6:17.315 | רע"ט |  |
| 6:17.316 | גילה |  |
| 6:17.317 | פנים |  |
| 6:17.318 | בחוצפא |  |
| 6:17.319 | יתירא |  |
| 6:17.320 | בשם |  |
| 6:17.321 | רעיא |  |
| 6:17.322 | מהימנא |  |
| 6:17.323 | לומר |  |
| 6:17.324 | אית |  |
| 6:17.325 | סלע |  |
| 6:17.326 | ואית |  |
| 6:17.327 | סלע | , |
| 6:17.328 | אית |  |
| 6:17.329 | אבן |  |
| 6:17.330 | ואית |  |
| 6:17.331 | אבן | , |
| 6:17.332 | אית |  |
| 6:17.333 | אבן |  |
| 6:17.334 | דשמא |  |
| 6:17.335 | דיהוה |  |
| 6:17.336 | היא |  |
| 6:17.337 | תורתו |  |
| 6:17.338 | החדשה |  |
| 6:17.339 | המלאה |  |
| 6:17.340 | פרצופים |  |
| 6:17.341 | וצורות |  |
| 6:17.342 | הנקראים |  |
| 6:17.343 | בשם |  |
| 6:17.344 | יהוה |  |
| 6:17.345 | ואדנות |  |
| 6:17.346 | ושאר |  |
| 6:17.347 | שמות |  |
| 6:17.348 | הקדש |  |
| 6:17.349 | נמחקים |  |
| 6:17.350 | ואין |  |
| 6:17.351 | נמחקים |  |
| 6:17.352 | ומרמזם |  |
| 6:17.353 | בשם |  |
| 6:17.354 | הוי"ה | ) |
| 6:17.355 | עלה |  |
| 6:17.356 | אמתר |  |
| 6:17.357 | ואבנא |  |
| 6:17.358 | די |  |
| 6:17.359 | מחת |  |
| 6:17.360 | לצלמא |  |
| 6:17.361 | הות |  |
| 6:17.362 | לטור |  |
| 6:17.363 | רב | , |
| 6:17.364 | ואית |  |
| 6:17.365 | אבן |  |
| 6:17.366 | דאיהי |  |
| 6:17.367 | אבן |  |
| 6:17.368 | משכית | , |
| 6:17.369 | דלית |  |
| 6:17.370 | תמן |  |
| 6:17.371 | נביעו |  |
| 6:17.372 | דחכמתא |  |
| 6:17.373 | ולא |  |
| 6:17.374 | דיבור |  |
| 6:17.375 | אלא |  |
| 6:17.376 | אבן |  |
| 6:17.377 | דאיהי |  |
| 6:17.378 | חסר |  |
| 6:17.379 | הפסוק |  |
| 6:17.380 | בכאן |  |
| 6:17.381 | ולא |  |
| 6:17.382 | חפץ |  |
| 6:17.383 | לפרסם |  |
| 6:17.384 | כוונתו |  |
| 6:17.385 | כל |  |
| 6:17.386 | כך |  |
| 6:17.387 | בבזוי |  |
| 6:17.388 | המשנה |  |
| 6:17.389 | והיה |  |
| 6:17.390 | רוצה |  |
| 6:17.391 | לומר |  |
| 6:17.392 | אבן |  |
| 6:17.393 | נגף |  |
| 6:17.394 | ואבן |  |
| 6:17.395 | מכשול | ). |
| 6:17.396 | ראה |  |
| 6:17.397 | והבן |  |
| 6:17.398 | בדבריו |  |
| 6:17.399 | איך |  |
| 6:17.400 | הוא |  |
| 6:17.401 | מחשיב |  |
| 6:17.402 | לתורתו |  |
| 6:17.403 | החדשה |  |
| 6:17.404 | וקורא |  |
| 6:17.405 | לה |  |
| 6:17.406 | אבנא |  |
| 6:17.407 | די |  |
| 6:17.408 | מחת |  |
| 6:17.409 | לצלמא | , |
| 6:17.410 | ובמלת |  |
| 6:17.411 | צלמא |  |
| 6:17.412 | כוונתו |  |
| 6:17.413 | על |  |
| 6:17.414 | המשנה |  |
| 6:17.415 | והתלמוד |  |
| 6:17.416 | שהוא |  |
| 6:17.417 | חושב |  |
| 6:17.418 | אותם |  |
| 6:17.419 | כצלמי |  |
| 6:17.420 | האלילים | , |
| 6:17.421 | ומבזה |  |
| 6:17.422 | אותה |  |
| 6:17.423 | וקורא |  |
| 6:17.424 | לה |  |
| 6:17.425 | אבן |  |
| 6:17.426 | משכית |  |
| 6:17.427 | דלית |  |
| 6:17.428 | בה |  |
| 6:17.429 | חכמתא | , |
| 6:17.430 | ובדבריו |  |
| 6:17.431 | אלה |  |
| 6:17.432 | הוא |  |
| 6:17.433 | רומז |  |
| 6:17.434 | לפרוש |  |
| 6:17.435 | מן |  |
| 6:17.436 | המשנה |  |
| 6:17.437 | כמאמר |  |
| 6:17.438 | הכתוב |  |
| 6:17.439 | ואבן |  |
| 6:17.440 | משכית |  |
| 6:17.441 | לא |  |
| 6:17.442 | תתנו |  |
| 6:17.443 | בארצכם |  |
| 6:17.444 | להשתחות |  |
| 6:17.445 | עליה |  |
| 6:17.446 | פי | ' |
| 6:17.447 | להשתחות |  |
| 6:17.448 | וליחד |  |
| 6:17.449 | השי"ת |  |
| 6:17.450 | על |  |
| 6:17.451 | פי |  |
| 6:17.452 | המשנה |  |
| 6:17.453 | המרחקת |  |
| 6:17.454 | כל |  |
| 6:17.455 | רבוי |  |
| 6:17.456 | באלהות | . |
| 6:17.457 | רק |  |
| 6:17.458 | על |  |
| 6:17.459 | פי |  |
| 6:17.460 | תורתו |  |
| 6:17.461 | החדשה |  |
| 6:17.462 | שהשי"ת |  |
| 6:17.463 | מחולק |  |
| 6:17.464 | לכמה |  |
| 6:17.465 | פרצופים |  |
| 6:17.466 | זכרים |  |
| 6:17.467 | ונקבות |  |
| 6:17.468 | וצריך |  |
| 6:17.469 | לקשרם |  |
| 6:17.470 | ולאחדם |  |
| 6:17.471 | ולזווגם |  |
| 6:17.472 | יחד |  |
| 6:17.473 | ר"ל | ). |
| 6:17.474 | ושוב |  |
| 6:17.475 | חזר |  |
| 6:17.476 | הפילסוף |  |
| 6:17.477 | הנז | ' |
| 6:17.478 | לדבריו |  |
| 6:17.479 | שבראש |  |
| 6:17.480 | עמוד |  |
| 6:17.481 | ב | ' |
| 6:17.482 | שקורא |  |
| 6:17.483 | לתורתו |  |
| 6:17.484 | סלע |  |
| 6:17.485 | הנובע |  |
| 6:17.486 | חכמה | , |
| 6:17.487 | ואת |  |
| 6:17.488 | המשנה |  |
| 6:17.489 | ואת |  |
| 6:17.490 | התלמוד |  |
| 6:17.491 | שקבל |  |
| 6:17.492 | משה |  |
| 6:17.493 | מסיני |  |
| 6:17.494 | הוא |  |
| 6:17.495 | מדמה |  |
| 6:17.496 | אותן |  |
| 6:17.497 | לקליפה |  |
| 6:17.498 | הקשה |  |
| 6:17.499 | של |  |
| 6:17.500 | אגוז |  |
| 6:17.501 | שכל |  |
| 6:17.502 | חכמי |  |
| 6:17.503 | הדורות |  |
| 6:17.504 | יגעים |  |
| 6:17.505 | בה |  |
| 6:17.506 | ולית |  |
| 6:17.507 | לון |  |
| 6:17.508 | יכולת |  |
| 6:17.509 | לאפקא |  |
| 6:17.510 | מינה |  |
| 6:17.511 | מיא |  |
| 6:17.512 | אלא |  |
| 6:17.513 | טיפין | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

