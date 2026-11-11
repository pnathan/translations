# Translation pack: milhamot-hashem perek-8-04 (Milchamot ha-Shem 8:5–8:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-04`.

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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- ר"א: "Rabbi Eleazar" only where the paragraph names him, else "R. A." with
  a note.
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
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
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


## Milchamot ha-Shem 8:5

Hebrew (Wikisource): קב) ולפי דברי הרב סוד ישרים לא הרוחנו מאומה בידיעת ענין הצמצום עם השתלשלות העגולים והפרצופים של יושר זה אחר זה, רק שעברנו על דברי משיח אלהי יעקב שאמר מי כה' אלהינו המגביהי לשבת המשפילי לראות בשמים ובארץ, ואומר אמרו בלבבכם על משכבכם ודומו סלה, והיה די לנו במאמר הנביא לכו ונלכה באור ה' (ולפי האמת האור הנזכר בפסוק הזה הוא אור התורה הקדושה שנאמר כי נר מצוה ותורה אור שהיה מאירה לאדם ומישרת דרכיו לשמור מצותיו וחוקותיו ית', אבל האורות הרוחניים אינם מאירים ואין בכחם ליתן תורה אחרת, כי לא יחליף האל לא ימיר תורתינו באחרת) ולמה זה נבקש חשבונות רבים המביאים לידי מינות לעבוד אל אחר קצר אפים, והאלהים עשה את האדם בשכל ישר, ונתן לו תורת אמת, והראנו את כבודו ואת גדלו ואמר כי לא ראיתם כל תמונה! ומ"מ גם לפי הנחת סוד ישרים שהאור הנברא עם כמה אורות קדושים רוחנים למעלה מהלמאכים ונקראים בשם ספירות ואין להם שום צורה ודמות, מ"מ בכלל צבא השמים שברא אלהינו יחשבו, ואסור לעבדן כיון שהם נבראים, ואין ראוי לעבוד אלא לאלהי קדם שבראם, בין שהם גשמים או רוחנים, וגם לצרפם ולאחדם ולשתפם לבוראם אסור לעבוד בלתי לה' לבדו שהוא הסבה הראשונה כמו שקבלו רז"ל וכמו שנבאר לעיל שכפי האמת הגמור היא שוה לאמונת הישמעאלים שקבלוה ממחוקקם שקבל מישראל כידוע, וזכו בני ישמעאל לזאת האמונה הטהורה ביחוד השי"ת בשכר תשובה שעשה ישמעאל אביהם, כד קימא לך שישמעאל עשה תשובה והאמין בה' שהוא הסבה הראשונה, ונקרא בלשונם אללה, וזה זכרי לדר דר וכו', ובלשון רז"ל הקב"ה או רבש"ע, והוא אלהינו ואנחנו עמו, ולו אנחנו עובדים ומשתחוים ומברכים, כמקובל בידינו מאבות אבותינו ומרז"ל (ואיך יאמר בעל ספר הברית שאין אמונתינו באמונת הישמעאלים) ולחנם נכתבו בדברי המקובלים החדשים כל אותן הציורים וההגשמות והפרצופים ומעלותיהם השונים זה מזה, והוא ב"ה רק בחפץ רצונו ית' מראה מכבודו לנביאים ובמשכן ובבית עולמים כדי לקבוע אמונתו בלבינו ככתוב כי לבעבור גסות אתכם בא האלהים ובעבור תהיה יראתו על פניכם לבלתי תחטאו, ואין כאן דבר ידוע וקבוע באור הנזכר לומר בו זה פרצוף פלוני, וזה פרצוף לפוני אלמוני, וזה קצר, וזה ארוך, וזה מעלתו למעלה, וזה למטה, ועד היכן מגיעים רגליו כמו שחשבו המקובלים החדשים, אלא הכל כחפצו ית' הוא מראה לנו.

| id | word | punctuation after |
|---|---|---|
| 8:5.1 | קב | ) |
| 8:5.2 | ולפי |  |
| 8:5.3 | דברי |  |
| 8:5.4 | הרב |  |
| 8:5.5 | סוד |  |
| 8:5.6 | ישרים |  |
| 8:5.7 | לא |  |
| 8:5.8 | הרוחנו |  |
| 8:5.9 | מאומה |  |
| 8:5.10 | בידיעת |  |
| 8:5.11 | ענין |  |
| 8:5.12 | הצמצום |  |
| 8:5.13 | עם |  |
| 8:5.14 | השתלשלות |  |
| 8:5.15 | העגולים |  |
| 8:5.16 | והפרצופים |  |
| 8:5.17 | של |  |
| 8:5.18 | יושר |  |
| 8:5.19 | זה |  |
| 8:5.20 | אחר |  |
| 8:5.21 | זה | , |
| 8:5.22 | רק |  |
| 8:5.23 | שעברנו |  |
| 8:5.24 | על |  |
| 8:5.25 | דברי |  |
| 8:5.26 | משיח |  |
| 8:5.27 | אלהי |  |
| 8:5.28 | יעקב |  |
| 8:5.29 | שאמר |  |
| 8:5.30 | מי |  |
| 8:5.31 | כה | ' |
| 8:5.32 | אלהינו |  |
| 8:5.33 | המגביהי |  |
| 8:5.34 | לשבת |  |
| 8:5.35 | המשפילי |  |
| 8:5.36 | לראות |  |
| 8:5.37 | בשמים |  |
| 8:5.38 | ובארץ | , |
| 8:5.39 | ואומר |  |
| 8:5.40 | אמרו |  |
| 8:5.41 | בלבבכם |  |
| 8:5.42 | על |  |
| 8:5.43 | משכבכם |  |
| 8:5.44 | ודומו |  |
| 8:5.45 | סלה | , |
| 8:5.46 | והיה |  |
| 8:5.47 | די |  |
| 8:5.48 | לנו |  |
| 8:5.49 | במאמר |  |
| 8:5.50 | הנביא |  |
| 8:5.51 | לכו |  |
| 8:5.52 | ונלכה |  |
| 8:5.53 | באור |  |
| 8:5.54 | ה | ' |
| 8:5.55 | ולפי |  |
| 8:5.56 | האמת |  |
| 8:5.57 | האור |  |
| 8:5.58 | הנזכר |  |
| 8:5.59 | בפסוק |  |
| 8:5.60 | הזה |  |
| 8:5.61 | הוא |  |
| 8:5.62 | אור |  |
| 8:5.63 | התורה |  |
| 8:5.64 | הקדושה |  |
| 8:5.65 | שנאמר |  |
| 8:5.66 | כי |  |
| 8:5.67 | נר |  |
| 8:5.68 | מצוה |  |
| 8:5.69 | ותורה |  |
| 8:5.70 | אור |  |
| 8:5.71 | שהיה |  |
| 8:5.72 | מאירה |  |
| 8:5.73 | לאדם |  |
| 8:5.74 | ומישרת |  |
| 8:5.75 | דרכיו |  |
| 8:5.76 | לשמור |  |
| 8:5.77 | מצותיו |  |
| 8:5.78 | וחוקותיו |  |
| 8:5.79 | ית | ', |
| 8:5.80 | אבל |  |
| 8:5.81 | האורות |  |
| 8:5.82 | הרוחניים |  |
| 8:5.83 | אינם |  |
| 8:5.84 | מאירים |  |
| 8:5.85 | ואין |  |
| 8:5.86 | בכחם |  |
| 8:5.87 | ליתן |  |
| 8:5.88 | תורה |  |
| 8:5.89 | אחרת | , |
| 8:5.90 | כי |  |
| 8:5.91 | לא |  |
| 8:5.92 | יחליף |  |
| 8:5.93 | האל |  |
| 8:5.94 | לא |  |
| 8:5.95 | ימיר |  |
| 8:5.96 | תורתינו |  |
| 8:5.97 | באחרת | ) |
| 8:5.98 | ולמה |  |
| 8:5.99 | זה |  |
| 8:5.100 | נבקש |  |
| 8:5.101 | חשבונות |  |
| 8:5.102 | רבים |  |
| 8:5.103 | המביאים |  |
| 8:5.104 | לידי |  |
| 8:5.105 | מינות |  |
| 8:5.106 | לעבוד |  |
| 8:5.107 | אל |  |
| 8:5.108 | אחר |  |
| 8:5.109 | קצר |  |
| 8:5.110 | אפים | , |
| 8:5.111 | והאלהים |  |
| 8:5.112 | עשה |  |
| 8:5.113 | את |  |
| 8:5.114 | האדם |  |
| 8:5.115 | בשכל |  |
| 8:5.116 | ישר | , |
| 8:5.117 | ונתן |  |
| 8:5.118 | לו |  |
| 8:5.119 | תורת |  |
| 8:5.120 | אמת | , |
| 8:5.121 | והראנו |  |
| 8:5.122 | את |  |
| 8:5.123 | כבודו |  |
| 8:5.124 | ואת |  |
| 8:5.125 | גדלו |  |
| 8:5.126 | ואמר |  |
| 8:5.127 | כי |  |
| 8:5.128 | לא |  |
| 8:5.129 | ראיתם |  |
| 8:5.130 | כל |  |
| 8:5.131 | תמונה | ! |
| 8:5.132 | ומ"מ |  |
| 8:5.133 | גם |  |
| 8:5.134 | לפי |  |
| 8:5.135 | הנחת |  |
| 8:5.136 | סוד |  |
| 8:5.137 | ישרים |  |
| 8:5.138 | שהאור |  |
| 8:5.139 | הנברא |  |
| 8:5.140 | עם |  |
| 8:5.141 | כמה |  |
| 8:5.142 | אורות |  |
| 8:5.143 | קדושים |  |
| 8:5.144 | רוחנים |  |
| 8:5.145 | למעלה |  |
| 8:5.146 | מהלמאכים |  |
| 8:5.147 | ונקראים |  |
| 8:5.148 | בשם |  |
| 8:5.149 | ספירות |  |
| 8:5.150 | ואין |  |
| 8:5.151 | להם |  |
| 8:5.152 | שום |  |
| 8:5.153 | צורה |  |
| 8:5.154 | ודמות | , |
| 8:5.155 | מ"מ |  |
| 8:5.156 | בכלל |  |
| 8:5.157 | צבא |  |
| 8:5.158 | השמים |  |
| 8:5.159 | שברא |  |
| 8:5.160 | אלהינו |  |
| 8:5.161 | יחשבו | , |
| 8:5.162 | ואסור |  |
| 8:5.163 | לעבדן |  |
| 8:5.164 | כיון |  |
| 8:5.165 | שהם |  |
| 8:5.166 | נבראים | , |
| 8:5.167 | ואין |  |
| 8:5.168 | ראוי |  |
| 8:5.169 | לעבוד |  |
| 8:5.170 | אלא |  |
| 8:5.171 | לאלהי |  |
| 8:5.172 | קדם |  |
| 8:5.173 | שבראם | , |
| 8:5.174 | בין |  |
| 8:5.175 | שהם |  |
| 8:5.176 | גשמים |  |
| 8:5.177 | או |  |
| 8:5.178 | רוחנים | , |
| 8:5.179 | וגם |  |
| 8:5.180 | לצרפם |  |
| 8:5.181 | ולאחדם |  |
| 8:5.182 | ולשתפם |  |
| 8:5.183 | לבוראם |  |
| 8:5.184 | אסור |  |
| 8:5.185 | לעבוד |  |
| 8:5.186 | בלתי |  |
| 8:5.187 | לה | ' |
| 8:5.188 | לבדו |  |
| 8:5.189 | שהוא |  |
| 8:5.190 | הסבה |  |
| 8:5.191 | הראשונה |  |
| 8:5.192 | כמו |  |
| 8:5.193 | שקבלו |  |
| 8:5.194 | רז"ל |  |
| 8:5.195 | וכמו |  |
| 8:5.196 | שנבאר |  |
| 8:5.197 | לעיל |  |
| 8:5.198 | שכפי |  |
| 8:5.199 | האמת |  |
| 8:5.200 | הגמור |  |
| 8:5.201 | היא |  |
| 8:5.202 | שוה |  |
| 8:5.203 | לאמונת |  |
| 8:5.204 | הישמעאלים |  |
| 8:5.205 | שקבלוה |  |
| 8:5.206 | ממחוקקם |  |
| 8:5.207 | שקבל |  |
| 8:5.208 | מישראל |  |
| 8:5.209 | כידוע | , |
| 8:5.210 | וזכו |  |
| 8:5.211 | בני |  |
| 8:5.212 | ישמעאל |  |
| 8:5.213 | לזאת |  |
| 8:5.214 | האמונה |  |
| 8:5.215 | הטהורה |  |
| 8:5.216 | ביחוד |  |
| 8:5.217 | השי"ת |  |
| 8:5.218 | בשכר |  |
| 8:5.219 | תשובה |  |
| 8:5.220 | שעשה |  |
| 8:5.221 | ישמעאל |  |
| 8:5.222 | אביהם | , |
| 8:5.223 | כד |  |
| 8:5.224 | קימא |  |
| 8:5.225 | לך |  |
| 8:5.226 | שישמעאל |  |
| 8:5.227 | עשה |  |
| 8:5.228 | תשובה |  |
| 8:5.229 | והאמין |  |
| 8:5.230 | בה | ' |
| 8:5.231 | שהוא |  |
| 8:5.232 | הסבה |  |
| 8:5.233 | הראשונה | , |
| 8:5.234 | ונקרא |  |
| 8:5.235 | בלשונם |  |
| 8:5.236 | אללה | , |
| 8:5.237 | וזה |  |
| 8:5.238 | זכרי |  |
| 8:5.239 | לדר |  |
| 8:5.240 | דר |  |
| 8:5.241 | וכו | ', |
| 8:5.242 | ובלשון |  |
| 8:5.243 | רז"ל |  |
| 8:5.244 | הקב"ה |  |
| 8:5.245 | או |  |
| 8:5.246 | רבש"ע | , |
| 8:5.247 | והוא |  |
| 8:5.248 | אלהינו |  |
| 8:5.249 | ואנחנו |  |
| 8:5.250 | עמו | , |
| 8:5.251 | ולו |  |
| 8:5.252 | אנחנו |  |
| 8:5.253 | עובדים |  |
| 8:5.254 | ומשתחוים |  |
| 8:5.255 | ומברכים | , |
| 8:5.256 | כמקובל |  |
| 8:5.257 | בידינו |  |
| 8:5.258 | מאבות |  |
| 8:5.259 | אבותינו |  |
| 8:5.260 | ומרז"ל |  |
| 8:5.261 | ואיך |  |
| 8:5.262 | יאמר |  |
| 8:5.263 | בעל |  |
| 8:5.264 | ספר |  |
| 8:5.265 | הברית |  |
| 8:5.266 | שאין |  |
| 8:5.267 | אמונתינו |  |
| 8:5.268 | באמונת |  |
| 8:5.269 | הישמעאלים | ) |
| 8:5.270 | ולחנם |  |
| 8:5.271 | נכתבו |  |
| 8:5.272 | בדברי |  |
| 8:5.273 | המקובלים |  |
| 8:5.274 | החדשים |  |
| 8:5.275 | כל |  |
| 8:5.276 | אותן |  |
| 8:5.277 | הציורים |  |
| 8:5.278 | וההגשמות |  |
| 8:5.279 | והפרצופים |  |
| 8:5.280 | ומעלותיהם |  |
| 8:5.281 | השונים |  |
| 8:5.282 | זה |  |
| 8:5.283 | מזה | , |
| 8:5.284 | והוא |  |
| 8:5.285 | ב"ה |  |
| 8:5.286 | רק |  |
| 8:5.287 | בחפץ |  |
| 8:5.288 | רצונו |  |
| 8:5.289 | ית | ' |
| 8:5.290 | מראה |  |
| 8:5.291 | מכבודו |  |
| 8:5.292 | לנביאים |  |
| 8:5.293 | ובמשכן |  |
| 8:5.294 | ובבית |  |
| 8:5.295 | עולמים |  |
| 8:5.296 | כדי |  |
| 8:5.297 | לקבוע |  |
| 8:5.298 | אמונתו |  |
| 8:5.299 | בלבינו |  |
| 8:5.300 | ככתוב |  |
| 8:5.301 | כי |  |
| 8:5.302 | לבעבור |  |
| 8:5.303 | גסות |  |
| 8:5.304 | אתכם |  |
| 8:5.305 | בא |  |
| 8:5.306 | האלהים |  |
| 8:5.307 | ובעבור |  |
| 8:5.308 | תהיה |  |
| 8:5.309 | יראתו |  |
| 8:5.310 | על |  |
| 8:5.311 | פניכם |  |
| 8:5.312 | לבלתי |  |
| 8:5.313 | תחטאו | , |
| 8:5.314 | ואין |  |
| 8:5.315 | כאן |  |
| 8:5.316 | דבר |  |
| 8:5.317 | ידוע |  |
| 8:5.318 | וקבוע |  |
| 8:5.319 | באור |  |
| 8:5.320 | הנזכר |  |
| 8:5.321 | לומר |  |
| 8:5.322 | בו |  |
| 8:5.323 | זה |  |
| 8:5.324 | פרצוף |  |
| 8:5.325 | פלוני | , |
| 8:5.326 | וזה |  |
| 8:5.327 | פרצוף |  |
| 8:5.328 | לפוני |  |
| 8:5.329 | אלמוני | , |
| 8:5.330 | וזה |  |
| 8:5.331 | קצר | , |
| 8:5.332 | וזה |  |
| 8:5.333 | ארוך | , |
| 8:5.334 | וזה |  |
| 8:5.335 | מעלתו |  |
| 8:5.336 | למעלה | , |
| 8:5.337 | וזה |  |
| 8:5.338 | למטה | , |
| 8:5.339 | ועד |  |
| 8:5.340 | היכן |  |
| 8:5.341 | מגיעים |  |
| 8:5.342 | רגליו |  |
| 8:5.343 | כמו |  |
| 8:5.344 | שחשבו |  |
| 8:5.345 | המקובלים |  |
| 8:5.346 | החדשים | , |
| 8:5.347 | אלא |  |
| 8:5.348 | הכל |  |
| 8:5.349 | כחפצו |  |
| 8:5.350 | ית | ' |
| 8:5.351 | הוא |  |
| 8:5.352 | מראה |  |
| 8:5.353 | לנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

