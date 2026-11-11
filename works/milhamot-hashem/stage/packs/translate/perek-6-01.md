# Translation pack: milhamot-hashem perek-6-01 (Milchamot ha-Shem 6:1–6:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-01`.

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
- **קצר אפים**: the author's Hebrew calque of זעיר אנפין, set against ארך
  אפים in Exod 34:6. Write "quick to anger", and render ארך אפים "slow to
  anger" everywhere, so that the pair shows in English. "באל זר קצר אפים
  (ז״א)" (1:4) → "a strange god, quick to anger (Ze'ir Anpin)"; "הנקרא (קצר
  אפים) זעיר אנפין" (3:3) → "called (quick to anger) Ze'ir Anpin". Add a note
  each time: Aramaic "small face"; Prov 14:17. "קצא אפים" (6:23) is translated
  as קצר, with the spelling recorded in `apparatus`.
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
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- Names of works: "Guide for the Perplexed" is rendered as the Moreh
  Nevukhim; ממלכי as "of the kings of".
- Closing formulas (וחתם דבריו, עכ"ל, ע"כ) stand outside the closing quotation
  mark; opening formulas (וז"ל, כ' וז"ל) stand outside the opening one.
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.
- הרא"ש: "the Rosh"; ר"ת: "Rabbeinu Tam"; הגאון alone: "the Gaon", unidentified.
  (Acronym authorities keep the acronym, by the ruling above.)

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


## Milchamot ha-Shem 6:1

Hebrew (Wikisource): ס) והנה מכל זה תלמוד לומר דודאי שהדברים שנאמרו בזהר בענייני האלהות לא מפי רשב"י ע"ה נאמרו ולא יצאו דברים כאלה מפי שום תנא או אמורא מרז"ל, ורק תלו אותן ברשב"י וחביריו בשקר, שהרי הוא בעצמו מכריז ואומר כל המשתף שם שמים ודבר אחר נעקר מן העולם! חלילה, חלילה להאמין שרשב"י התנא או שום אחר מרבותינו התנאים או האמוראים יחשוב או יאמר כדברים האלה לעבוד אלוה אחר קצר אפים ולשתפו במחשבה עם ה' אלהינו ארך אפים שהוא הסבה הראשונה שהוא בתור נשמה לפי דעתם לפרצופים האלה ועובדים ומתפללים לפרצופים הנז' ואומרים כי לנשמתם שהוא חלק האין סוף שבהם הם עובדים, ומיחדים הפרצופים הזכרים עם נקיבותיהם הנרמזים בשם הוי"ה, יוד חכמה שהוא אבא, ה"א ראשונה בינה אימא, וי"ו תפארת זעיר, ה"א אחרונה מלכות, ובשמע ישראל הוי"ה ראשונה אבא, אלהינו אמא, הוי"ה אחרונה תפארת זעיר, אחד מלכות, לחברם ולקשרם יחד להיות הארבעה אחד (זהר ואתחנן דף רס"ג ע"א), כי עיקר היחוד אצלם לחבר ולקשר במחשבתו עיניינים כשהם מצויירים במחשבתו גופי אור נפרדים והוא דוחק את מחשבתו לשום אותם יחד, וגם בפיו אומר לשם קב"ה ושכינתיה וכו' ליחדא שם יו"ד וה"א בו"ו וה"א וכו'. וכוונתו על אבא ואימא וזעיר ונוקביה הנרמזים לפי דעתם בשם הוי״ה, ואריך ועתיק הנרמזים בקוצו של יו"ד ועיין לעיל ס' מ"ב. ולפי דבריהם למה לא הזכירו רז"ל במשנה ובגמרא במסכת זבחים על הא דתנן לשם ששה דברים הזבח נזבח, לשם זבח, לשם זובח, לשם השם, לשם אישים, לשם ריח, לשם ניחח, והחטאת והאשם לשם חטא, ופי' בגמרא לשם ניחח, נחת רוח להקב"ה שאמר ונעשה רצונו, ולא קחשיב לשם יחוד קב"ה ושכינתיה, אלא ודאי שדעת רז"ל שאלהינו ב"ה יחיד ומיוחד מכל שאר האחדים, ואינו צריך לקשרו ולאחדו מחלקים רבים כמו שחשבו המקובלים החדשים ולא לשתפו עם פרצופים חלוקים ושונים זה משה! חלילה לנו מהעלות על לבבינו דברים זרים לתורתינו הקדושה כאלה שלא צונו הקב"ה בתורה לאחדו ולשתפו עם שום נברא, רק לדעת ולהאמין כי הוא אחד שנאמר וידעת היום והשבות אל לבביך כי ה' הוא האלהים וכו' אין עוד, ועיין לקמן סי' קל"ו.

| id | word | punctuation after |
|---|---|---|
| 6:1.1 | ס | ) |
| 6:1.2 | והנה |  |
| 6:1.3 | מכל |  |
| 6:1.4 | זה |  |
| 6:1.5 | תלמוד |  |
| 6:1.6 | לומר |  |
| 6:1.7 | דודאי |  |
| 6:1.8 | שהדברים |  |
| 6:1.9 | שנאמרו |  |
| 6:1.10 | בזהר |  |
| 6:1.11 | בענייני |  |
| 6:1.12 | האלהות |  |
| 6:1.13 | לא |  |
| 6:1.14 | מפי |  |
| 6:1.15 | רשב"י |  |
| 6:1.16 | ע"ה |  |
| 6:1.17 | נאמרו |  |
| 6:1.18 | ולא |  |
| 6:1.19 | יצאו |  |
| 6:1.20 | דברים |  |
| 6:1.21 | כאלה |  |
| 6:1.22 | מפי |  |
| 6:1.23 | שום |  |
| 6:1.24 | תנא |  |
| 6:1.25 | או |  |
| 6:1.26 | אמורא |  |
| 6:1.27 | מרז"ל | , |
| 6:1.28 | ורק |  |
| 6:1.29 | תלו |  |
| 6:1.30 | אותן |  |
| 6:1.31 | ברשב"י |  |
| 6:1.32 | וחביריו |  |
| 6:1.33 | בשקר | , |
| 6:1.34 | שהרי |  |
| 6:1.35 | הוא |  |
| 6:1.36 | בעצמו |  |
| 6:1.37 | מכריז |  |
| 6:1.38 | ואומר |  |
| 6:1.39 | כל |  |
| 6:1.40 | המשתף |  |
| 6:1.41 | שם |  |
| 6:1.42 | שמים |  |
| 6:1.43 | ודבר |  |
| 6:1.44 | אחר |  |
| 6:1.45 | נעקר |  |
| 6:1.46 | מן |  |
| 6:1.47 | העולם | ! |
| 6:1.48 | חלילה | , |
| 6:1.49 | חלילה |  |
| 6:1.50 | להאמין |  |
| 6:1.51 | שרשב"י |  |
| 6:1.52 | התנא |  |
| 6:1.53 | או |  |
| 6:1.54 | שום |  |
| 6:1.55 | אחר |  |
| 6:1.56 | מרבותינו |  |
| 6:1.57 | התנאים |  |
| 6:1.58 | או |  |
| 6:1.59 | האמוראים |  |
| 6:1.60 | יחשוב |  |
| 6:1.61 | או |  |
| 6:1.62 | יאמר |  |
| 6:1.63 | כדברים |  |
| 6:1.64 | האלה |  |
| 6:1.65 | לעבוד |  |
| 6:1.66 | אלוה |  |
| 6:1.67 | אחר |  |
| 6:1.68 | קצר |  |
| 6:1.69 | אפים |  |
| 6:1.70 | ולשתפו |  |
| 6:1.71 | במחשבה |  |
| 6:1.72 | עם |  |
| 6:1.73 | ה | ' |
| 6:1.74 | אלהינו |  |
| 6:1.75 | ארך |  |
| 6:1.76 | אפים |  |
| 6:1.77 | שהוא |  |
| 6:1.78 | הסבה |  |
| 6:1.79 | הראשונה |  |
| 6:1.80 | שהוא |  |
| 6:1.81 | בתור |  |
| 6:1.82 | נשמה |  |
| 6:1.83 | לפי |  |
| 6:1.84 | דעתם |  |
| 6:1.85 | לפרצופים |  |
| 6:1.86 | האלה |  |
| 6:1.87 | ועובדים |  |
| 6:1.88 | ומתפללים |  |
| 6:1.89 | לפרצופים |  |
| 6:1.90 | הנז | ' |
| 6:1.91 | ואומרים |  |
| 6:1.92 | כי |  |
| 6:1.93 | לנשמתם |  |
| 6:1.94 | שהוא |  |
| 6:1.95 | חלק |  |
| 6:1.96 | האין |  |
| 6:1.97 | סוף |  |
| 6:1.98 | שבהם |  |
| 6:1.99 | הם |  |
| 6:1.100 | עובדים | , |
| 6:1.101 | ומיחדים |  |
| 6:1.102 | הפרצופים |  |
| 6:1.103 | הזכרים |  |
| 6:1.104 | עם |  |
| 6:1.105 | נקיבותיהם |  |
| 6:1.106 | הנרמזים |  |
| 6:1.107 | בשם |  |
| 6:1.108 | הוי"ה | , |
| 6:1.109 | יוד |  |
| 6:1.110 | חכמה |  |
| 6:1.111 | שהוא |  |
| 6:1.112 | אבא | , |
| 6:1.113 | ה"א |  |
| 6:1.114 | ראשונה |  |
| 6:1.115 | בינה |  |
| 6:1.116 | אימא | , |
| 6:1.117 | וי"ו |  |
| 6:1.118 | תפארת |  |
| 6:1.119 | זעיר | , |
| 6:1.120 | ה"א |  |
| 6:1.121 | אחרונה |  |
| 6:1.122 | מלכות | , |
| 6:1.123 | ובשמע |  |
| 6:1.124 | ישראל |  |
| 6:1.125 | הוי"ה |  |
| 6:1.126 | ראשונה |  |
| 6:1.127 | אבא | , |
| 6:1.128 | אלהינו |  |
| 6:1.129 | אמא | , |
| 6:1.130 | הוי"ה |  |
| 6:1.131 | אחרונה |  |
| 6:1.132 | תפארת |  |
| 6:1.133 | זעיר | , |
| 6:1.134 | אחד |  |
| 6:1.135 | מלכות | , |
| 6:1.136 | לחברם |  |
| 6:1.137 | ולקשרם |  |
| 6:1.138 | יחד |  |
| 6:1.139 | להיות |  |
| 6:1.140 | הארבעה |  |
| 6:1.141 | אחד |  |
| 6:1.142 | זהר |  |
| 6:1.143 | ואתחנן |  |
| 6:1.144 | דף |  |
| 6:1.145 | רס"ג |  |
| 6:1.146 | ע"א | ), |
| 6:1.147 | כי |  |
| 6:1.148 | עיקר |  |
| 6:1.149 | היחוד |  |
| 6:1.150 | אצלם |  |
| 6:1.151 | לחבר |  |
| 6:1.152 | ולקשר |  |
| 6:1.153 | במחשבתו |  |
| 6:1.154 | עיניינים |  |
| 6:1.155 | כשהם |  |
| 6:1.156 | מצויירים |  |
| 6:1.157 | במחשבתו |  |
| 6:1.158 | גופי |  |
| 6:1.159 | אור |  |
| 6:1.160 | נפרדים |  |
| 6:1.161 | והוא |  |
| 6:1.162 | דוחק |  |
| 6:1.163 | את |  |
| 6:1.164 | מחשבתו |  |
| 6:1.165 | לשום |  |
| 6:1.166 | אותם |  |
| 6:1.167 | יחד | , |
| 6:1.168 | וגם |  |
| 6:1.169 | בפיו |  |
| 6:1.170 | אומר |  |
| 6:1.171 | לשם |  |
| 6:1.172 | קב"ה |  |
| 6:1.173 | ושכינתיה |  |
| 6:1.174 | וכו | ' |
| 6:1.175 | ליחדא |  |
| 6:1.176 | שם |  |
| 6:1.177 | יו"ד |  |
| 6:1.178 | וה"א |  |
| 6:1.179 | בו"ו |  |
| 6:1.180 | וה"א |  |
| 6:1.181 | וכו | '. |
| 6:1.182 | וכוונתו |  |
| 6:1.183 | על |  |
| 6:1.184 | אבא |  |
| 6:1.185 | ואימא |  |
| 6:1.186 | וזעיר |  |
| 6:1.187 | ונוקביה |  |
| 6:1.188 | הנרמזים |  |
| 6:1.189 | לפי |  |
| 6:1.190 | דעתם |  |
| 6:1.191 | בשם |  |
| 6:1.192 | הוי״ה | , |
| 6:1.193 | ואריך |  |
| 6:1.194 | ועתיק |  |
| 6:1.195 | הנרמזים |  |
| 6:1.196 | בקוצו |  |
| 6:1.197 | של |  |
| 6:1.198 | יו"ד |  |
| 6:1.199 | ועיין |  |
| 6:1.200 | לעיל |  |
| 6:1.201 | ס | ' |
| 6:1.202 | מ"ב | . |
| 6:1.203 | ולפי |  |
| 6:1.204 | דבריהם |  |
| 6:1.205 | למה |  |
| 6:1.206 | לא |  |
| 6:1.207 | הזכירו |  |
| 6:1.208 | רז"ל |  |
| 6:1.209 | במשנה |  |
| 6:1.210 | ובגמרא |  |
| 6:1.211 | במסכת |  |
| 6:1.212 | זבחים |  |
| 6:1.213 | על |  |
| 6:1.214 | הא |  |
| 6:1.215 | דתנן |  |
| 6:1.216 | לשם |  |
| 6:1.217 | ששה |  |
| 6:1.218 | דברים |  |
| 6:1.219 | הזבח |  |
| 6:1.220 | נזבח | , |
| 6:1.221 | לשם |  |
| 6:1.222 | זבח | , |
| 6:1.223 | לשם |  |
| 6:1.224 | זובח | , |
| 6:1.225 | לשם |  |
| 6:1.226 | השם | , |
| 6:1.227 | לשם |  |
| 6:1.228 | אישים | , |
| 6:1.229 | לשם |  |
| 6:1.230 | ריח | , |
| 6:1.231 | לשם |  |
| 6:1.232 | ניחח | , |
| 6:1.233 | והחטאת |  |
| 6:1.234 | והאשם |  |
| 6:1.235 | לשם |  |
| 6:1.236 | חטא | , |
| 6:1.237 | ופי | ' |
| 6:1.238 | בגמרא |  |
| 6:1.239 | לשם |  |
| 6:1.240 | ניחח | , |
| 6:1.241 | נחת |  |
| 6:1.242 | רוח |  |
| 6:1.243 | להקב"ה |  |
| 6:1.244 | שאמר |  |
| 6:1.245 | ונעשה |  |
| 6:1.246 | רצונו | , |
| 6:1.247 | ולא |  |
| 6:1.248 | קחשיב |  |
| 6:1.249 | לשם |  |
| 6:1.250 | יחוד |  |
| 6:1.251 | קב"ה |  |
| 6:1.252 | ושכינתיה | , |
| 6:1.253 | אלא |  |
| 6:1.254 | ודאי |  |
| 6:1.255 | שדעת |  |
| 6:1.256 | רז"ל |  |
| 6:1.257 | שאלהינו |  |
| 6:1.258 | ב"ה |  |
| 6:1.259 | יחיד |  |
| 6:1.260 | ומיוחד |  |
| 6:1.261 | מכל |  |
| 6:1.262 | שאר |  |
| 6:1.263 | האחדים | , |
| 6:1.264 | ואינו |  |
| 6:1.265 | צריך |  |
| 6:1.266 | לקשרו |  |
| 6:1.267 | ולאחדו |  |
| 6:1.268 | מחלקים |  |
| 6:1.269 | רבים |  |
| 6:1.270 | כמו |  |
| 6:1.271 | שחשבו |  |
| 6:1.272 | המקובלים |  |
| 6:1.273 | החדשים |  |
| 6:1.274 | ולא |  |
| 6:1.275 | לשתפו |  |
| 6:1.276 | עם |  |
| 6:1.277 | פרצופים |  |
| 6:1.278 | חלוקים |  |
| 6:1.279 | ושונים |  |
| 6:1.280 | זה |  |
| 6:1.281 | משה | ! |
| 6:1.282 | חלילה |  |
| 6:1.283 | לנו |  |
| 6:1.284 | מהעלות |  |
| 6:1.285 | על |  |
| 6:1.286 | לבבינו |  |
| 6:1.287 | דברים |  |
| 6:1.288 | זרים |  |
| 6:1.289 | לתורתינו |  |
| 6:1.290 | הקדושה |  |
| 6:1.291 | כאלה |  |
| 6:1.292 | שלא |  |
| 6:1.293 | צונו |  |
| 6:1.294 | הקב"ה |  |
| 6:1.295 | בתורה |  |
| 6:1.296 | לאחדו |  |
| 6:1.297 | ולשתפו |  |
| 6:1.298 | עם |  |
| 6:1.299 | שום |  |
| 6:1.300 | נברא | , |
| 6:1.301 | רק |  |
| 6:1.302 | לדעת |  |
| 6:1.303 | ולהאמין |  |
| 6:1.304 | כי |  |
| 6:1.305 | הוא |  |
| 6:1.306 | אחד |  |
| 6:1.307 | שנאמר |  |
| 6:1.308 | וידעת |  |
| 6:1.309 | היום |  |
| 6:1.310 | והשבות |  |
| 6:1.311 | אל |  |
| 6:1.312 | לבביך |  |
| 6:1.313 | כי |  |
| 6:1.314 | ה | ' |
| 6:1.315 | הוא |  |
| 6:1.316 | האלהים |  |
| 6:1.317 | וכו | ' |
| 6:1.318 | אין |  |
| 6:1.319 | עוד | , |
| 6:1.320 | ועיין |  |
| 6:1.321 | לקמן |  |
| 6:1.322 | סי | ' |
| 6:1.323 | קל"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

