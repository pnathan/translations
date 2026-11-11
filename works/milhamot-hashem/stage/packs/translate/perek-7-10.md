# Translation pack: milhamot-hashem perek-7-10 (Milchamot ha-Shem 7:14–7:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-10`.

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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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
- דמיון: "analogy" for an argument, "likeness" for resemblance.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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
- אלפין: "alefs" (the letter).
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
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


## Milchamot ha-Shem 7:14

Hebrew (Wikisource): פז) ורב סעדיה גאון בפי' דניאל כ' וז"ל ודע כי דינא יתיב וספרין פתיחו אשר ביארתי למעלה יום דין ופקודה, הוא יום העתיד לדרוש כל מעשה בני אדם החיים והמתים, ואשר כתב למעלה ועתיק יומין יתיב, כל האומר כי ראה ה' צבאות עין בעין. הרי הוא פושע, שמדמה דמות לשם ית' ואין לו דמות, דכ' ואל מי תדמיון אל, וכל האומר כי ראה מלאך הוא מן השונים, והוא מן האומרים קטון וגדול, למען כי אין שופט העולם כי אם הקב"ה לבדו, ככתוב לפני ה' כי בא לשפוט את הארץ, ואין מלאך אשר יוכל לדבר, כי אלף אלפין ישמשוניה. ועת הדע כי דניאל ראה כל אלה בחלום ולא במראה עינים וכל החלום יש לו פתרון. כי הם משל, כי גם כל נבואת הנביאים אשר ראו ותחת רגליו כמעשה לבנת נספיר. וממראה מתניו ולמעלה, יושב על כסא רם ונשא, לא ראו חק חקוקו, כי אין דמות לבורא ית' כלל, ולא דמות מלאך, כי בעת שהיה מדבר עמם מראה להם חזיון חדש. וכל דמות שראו יש לו פתרון. מקל שקד, כי שוקד אני רבבות גדולות, מעשיו בלא מספר, וכאשר ראה מיכה ברוח נבואה, וכל צבא השמים עומדים מימינו ומשמאלו, לא נוכל לדמות ליוצר הכל מימינו ומשמאלו, ולא יעמוד לפניו כל יצור, ואולם דברי כל אלה במשל כמעשה בני אדם עכ"ל. והנה כבר ראו עיניך ומלאו אזניך ממה שקדם כי ה' אלהינו לא שנה ולא ישתנה. כמאמר הנשביא אני ה' לא שניתי. וכבר האריכו רבותינו הקדמונים הרס"ג ובעל חובת הלבבות ור' יהודה הליו והרמב"ם בפי' המשנה ובס' היד ובמורה, ובעל שבילי אמונה שאין ליחס לשי"ת שום שינוי והתפעלות. וגם לא מינה אחר לשפוט את הארץ. ודברים אלו מוסמכים מכל רבותינו כמבואר בדבריהם. ורבותינו בעלי הקבלה החדשה שגו בכל אלה. כי האמינו לדברי נביא השקר המתנבא בשם עתיק יומין לעבוד אל אחר קצא אפים (זעיר אנפין) הוא מחבר הזהר אשר בילדי נכרים ישפיק. ותלה אותם ברשב"י התנא בשקר. ולכן שגו ברואה שלא די שיחסו לשי"ת רבוי וחלוק צורות ופרצופים שונים וקוים. אלא שיחסו לו שינוים רבים בעצמותו ית' ואמרו שמתחלה היה מקום וחלל גדול פנוי וריק להיות שם כל ברואיו, ונעשה ככדור עגול וחלול באמצעותו, ושוב נעשה בו שינוי אחר, כי המשיך מאורו הגדול המקיף, ספירות דעיגולי אדם קדמאה סתימא, ואחר זה עיגולי אדם קדמון. ועיגולי עתיק יומין, ועיגולי אריך אנפין, ועיגולי אבא ואימא. גם עיגולי זעיר ונוקביה. ושוב נעשו עוד בו שינויים רבים להמשיך ג"כ מעצמותו ספירות דיושר בתנית אדם. והם עיקר האלוהות לפי דעתם. תחלה פרצוף אדם קדמון. ואחריו פרצוף עתיק יומין. ואחריו אריך אנפין. ושניהם כולם זכר ונקבה. רק שזה ימין ושמאל. וזה פנים ואחור, ואחריהם פרצופי אבא ואימא זכר ונקבה נפרדים זה מזה. ואחריהם פרצופי זעיר ונוקביה ג"כ נפרדים. והם עיקר האלוהות אשר לפי דבריהם אלהים נקרא בכל תפלותינו. וכל ברכותינו אליהם יכונו. כי הם אלוהינו לפי דעתם. הרי נתבאר לך כמה שינוים וחליפות נעשו באלהינו על פי הקבלה החדשה הסותרת אמונת תוה"ק.

| id | word | punctuation after |
|---|---|---|
| 7:14.1 | פז | ) |
| 7:14.2 | ורב |  |
| 7:14.3 | סעדיה |  |
| 7:14.4 | גאון |  |
| 7:14.5 | בפי | ' |
| 7:14.6 | דניאל |  |
| 7:14.7 | כ | ' |
| 7:14.8 | וז"ל |  |
| 7:14.9 | ודע |  |
| 7:14.10 | כי |  |
| 7:14.11 | דינא |  |
| 7:14.12 | יתיב |  |
| 7:14.13 | וספרין |  |
| 7:14.14 | פתיחו |  |
| 7:14.15 | אשר |  |
| 7:14.16 | ביארתי |  |
| 7:14.17 | למעלה |  |
| 7:14.18 | יום |  |
| 7:14.19 | דין |  |
| 7:14.20 | ופקודה | , |
| 7:14.21 | הוא |  |
| 7:14.22 | יום |  |
| 7:14.23 | העתיד |  |
| 7:14.24 | לדרוש |  |
| 7:14.25 | כל |  |
| 7:14.26 | מעשה |  |
| 7:14.27 | בני |  |
| 7:14.28 | אדם |  |
| 7:14.29 | החיים |  |
| 7:14.30 | והמתים | , |
| 7:14.31 | ואשר |  |
| 7:14.32 | כתב |  |
| 7:14.33 | למעלה |  |
| 7:14.34 | ועתיק |  |
| 7:14.35 | יומין |  |
| 7:14.36 | יתיב | , |
| 7:14.37 | כל |  |
| 7:14.38 | האומר |  |
| 7:14.39 | כי |  |
| 7:14.40 | ראה |  |
| 7:14.41 | ה | ' |
| 7:14.42 | צבאות |  |
| 7:14.43 | עין |  |
| 7:14.44 | בעין | . |
| 7:14.45 | הרי |  |
| 7:14.46 | הוא |  |
| 7:14.47 | פושע | , |
| 7:14.48 | שמדמה |  |
| 7:14.49 | דמות |  |
| 7:14.50 | לשם |  |
| 7:14.51 | ית | ' |
| 7:14.52 | ואין |  |
| 7:14.53 | לו |  |
| 7:14.54 | דמות | , |
| 7:14.55 | דכ | ' |
| 7:14.56 | ואל |  |
| 7:14.57 | מי |  |
| 7:14.58 | תדמיון |  |
| 7:14.59 | אל | , |
| 7:14.60 | וכל |  |
| 7:14.61 | האומר |  |
| 7:14.62 | כי |  |
| 7:14.63 | ראה |  |
| 7:14.64 | מלאך |  |
| 7:14.65 | הוא |  |
| 7:14.66 | מן |  |
| 7:14.67 | השונים | , |
| 7:14.68 | והוא |  |
| 7:14.69 | מן |  |
| 7:14.70 | האומרים |  |
| 7:14.71 | קטון |  |
| 7:14.72 | וגדול | , |
| 7:14.73 | למען |  |
| 7:14.74 | כי |  |
| 7:14.75 | אין |  |
| 7:14.76 | שופט |  |
| 7:14.77 | העולם |  |
| 7:14.78 | כי |  |
| 7:14.79 | אם |  |
| 7:14.80 | הקב"ה |  |
| 7:14.81 | לבדו | , |
| 7:14.82 | ככתוב |  |
| 7:14.83 | לפני |  |
| 7:14.84 | ה | ' |
| 7:14.85 | כי |  |
| 7:14.86 | בא |  |
| 7:14.87 | לשפוט |  |
| 7:14.88 | את |  |
| 7:14.89 | הארץ | , |
| 7:14.90 | ואין |  |
| 7:14.91 | מלאך |  |
| 7:14.92 | אשר |  |
| 7:14.93 | יוכל |  |
| 7:14.94 | לדבר | , |
| 7:14.95 | כי |  |
| 7:14.96 | אלף |  |
| 7:14.97 | אלפין |  |
| 7:14.98 | ישמשוניה | . |
| 7:14.99 | ועת |  |
| 7:14.100 | הדע |  |
| 7:14.101 | כי |  |
| 7:14.102 | דניאל |  |
| 7:14.103 | ראה |  |
| 7:14.104 | כל |  |
| 7:14.105 | אלה |  |
| 7:14.106 | בחלום |  |
| 7:14.107 | ולא |  |
| 7:14.108 | במראה |  |
| 7:14.109 | עינים |  |
| 7:14.110 | וכל |  |
| 7:14.111 | החלום |  |
| 7:14.112 | יש |  |
| 7:14.113 | לו |  |
| 7:14.114 | פתרון | . |
| 7:14.115 | כי |  |
| 7:14.116 | הם |  |
| 7:14.117 | משל | , |
| 7:14.118 | כי |  |
| 7:14.119 | גם |  |
| 7:14.120 | כל |  |
| 7:14.121 | נבואת |  |
| 7:14.122 | הנביאים |  |
| 7:14.123 | אשר |  |
| 7:14.124 | ראו |  |
| 7:14.125 | ותחת |  |
| 7:14.126 | רגליו |  |
| 7:14.127 | כמעשה |  |
| 7:14.128 | לבנת |  |
| 7:14.129 | נספיר | . |
| 7:14.130 | וממראה |  |
| 7:14.131 | מתניו |  |
| 7:14.132 | ולמעלה | , |
| 7:14.133 | יושב |  |
| 7:14.134 | על |  |
| 7:14.135 | כסא |  |
| 7:14.136 | רם |  |
| 7:14.137 | ונשא | , |
| 7:14.138 | לא |  |
| 7:14.139 | ראו |  |
| 7:14.140 | חק |  |
| 7:14.141 | חקוקו | , |
| 7:14.142 | כי |  |
| 7:14.143 | אין |  |
| 7:14.144 | דמות |  |
| 7:14.145 | לבורא |  |
| 7:14.146 | ית | ' |
| 7:14.147 | כלל | , |
| 7:14.148 | ולא |  |
| 7:14.149 | דמות |  |
| 7:14.150 | מלאך | , |
| 7:14.151 | כי |  |
| 7:14.152 | בעת |  |
| 7:14.153 | שהיה |  |
| 7:14.154 | מדבר |  |
| 7:14.155 | עמם |  |
| 7:14.156 | מראה |  |
| 7:14.157 | להם |  |
| 7:14.158 | חזיון |  |
| 7:14.159 | חדש | . |
| 7:14.160 | וכל |  |
| 7:14.161 | דמות |  |
| 7:14.162 | שראו |  |
| 7:14.163 | יש |  |
| 7:14.164 | לו |  |
| 7:14.165 | פתרון | . |
| 7:14.166 | מקל |  |
| 7:14.167 | שקד | , |
| 7:14.168 | כי |  |
| 7:14.169 | שוקד |  |
| 7:14.170 | אני |  |
| 7:14.171 | רבבות |  |
| 7:14.172 | גדולות | , |
| 7:14.173 | מעשיו |  |
| 7:14.174 | בלא |  |
| 7:14.175 | מספר | , |
| 7:14.176 | וכאשר |  |
| 7:14.177 | ראה |  |
| 7:14.178 | מיכה |  |
| 7:14.179 | ברוח |  |
| 7:14.180 | נבואה | , |
| 7:14.181 | וכל |  |
| 7:14.182 | צבא |  |
| 7:14.183 | השמים |  |
| 7:14.184 | עומדים |  |
| 7:14.185 | מימינו |  |
| 7:14.186 | ומשמאלו | , |
| 7:14.187 | לא |  |
| 7:14.188 | נוכל |  |
| 7:14.189 | לדמות |  |
| 7:14.190 | ליוצר |  |
| 7:14.191 | הכל |  |
| 7:14.192 | מימינו |  |
| 7:14.193 | ומשמאלו | , |
| 7:14.194 | ולא |  |
| 7:14.195 | יעמוד |  |
| 7:14.196 | לפניו |  |
| 7:14.197 | כל |  |
| 7:14.198 | יצור | , |
| 7:14.199 | ואולם |  |
| 7:14.200 | דברי |  |
| 7:14.201 | כל |  |
| 7:14.202 | אלה |  |
| 7:14.203 | במשל |  |
| 7:14.204 | כמעשה |  |
| 7:14.205 | בני |  |
| 7:14.206 | אדם |  |
| 7:14.207 | עכ"ל | . |
| 7:14.208 | והנה |  |
| 7:14.209 | כבר |  |
| 7:14.210 | ראו |  |
| 7:14.211 | עיניך |  |
| 7:14.212 | ומלאו |  |
| 7:14.213 | אזניך |  |
| 7:14.214 | ממה |  |
| 7:14.215 | שקדם |  |
| 7:14.216 | כי |  |
| 7:14.217 | ה | ' |
| 7:14.218 | אלהינו |  |
| 7:14.219 | לא |  |
| 7:14.220 | שנה |  |
| 7:14.221 | ולא |  |
| 7:14.222 | ישתנה | . |
| 7:14.223 | כמאמר |  |
| 7:14.224 | הנשביא |  |
| 7:14.225 | אני |  |
| 7:14.226 | ה | ' |
| 7:14.227 | לא |  |
| 7:14.228 | שניתי | . |
| 7:14.229 | וכבר |  |
| 7:14.230 | האריכו |  |
| 7:14.231 | רבותינו |  |
| 7:14.232 | הקדמונים |  |
| 7:14.233 | הרס"ג |  |
| 7:14.234 | ובעל |  |
| 7:14.235 | חובת |  |
| 7:14.236 | הלבבות |  |
| 7:14.237 | ור | ' |
| 7:14.238 | יהודה |  |
| 7:14.239 | הליו |  |
| 7:14.240 | והרמב"ם |  |
| 7:14.241 | בפי | ' |
| 7:14.242 | המשנה |  |
| 7:14.243 | ובס | ' |
| 7:14.244 | היד |  |
| 7:14.245 | ובמורה | , |
| 7:14.246 | ובעל |  |
| 7:14.247 | שבילי |  |
| 7:14.248 | אמונה |  |
| 7:14.249 | שאין |  |
| 7:14.250 | ליחס |  |
| 7:14.251 | לשי"ת |  |
| 7:14.252 | שום |  |
| 7:14.253 | שינוי |  |
| 7:14.254 | והתפעלות | . |
| 7:14.255 | וגם |  |
| 7:14.256 | לא |  |
| 7:14.257 | מינה |  |
| 7:14.258 | אחר |  |
| 7:14.259 | לשפוט |  |
| 7:14.260 | את |  |
| 7:14.261 | הארץ | . |
| 7:14.262 | ודברים |  |
| 7:14.263 | אלו |  |
| 7:14.264 | מוסמכים |  |
| 7:14.265 | מכל |  |
| 7:14.266 | רבותינו |  |
| 7:14.267 | כמבואר |  |
| 7:14.268 | בדבריהם | . |
| 7:14.269 | ורבותינו |  |
| 7:14.270 | בעלי |  |
| 7:14.271 | הקבלה |  |
| 7:14.272 | החדשה |  |
| 7:14.273 | שגו |  |
| 7:14.274 | בכל |  |
| 7:14.275 | אלה | . |
| 7:14.276 | כי |  |
| 7:14.277 | האמינו |  |
| 7:14.278 | לדברי |  |
| 7:14.279 | נביא |  |
| 7:14.280 | השקר |  |
| 7:14.281 | המתנבא |  |
| 7:14.282 | בשם |  |
| 7:14.283 | עתיק |  |
| 7:14.284 | יומין |  |
| 7:14.285 | לעבוד |  |
| 7:14.286 | אל |  |
| 7:14.287 | אחר |  |
| 7:14.288 | קצא |  |
| 7:14.289 | אפים |  |
| 7:14.290 | זעיר |  |
| 7:14.291 | אנפין | ) |
| 7:14.292 | הוא |  |
| 7:14.293 | מחבר |  |
| 7:14.294 | הזהר |  |
| 7:14.295 | אשר |  |
| 7:14.296 | בילדי |  |
| 7:14.297 | נכרים |  |
| 7:14.298 | ישפיק | . |
| 7:14.299 | ותלה |  |
| 7:14.300 | אותם |  |
| 7:14.301 | ברשב"י |  |
| 7:14.302 | התנא |  |
| 7:14.303 | בשקר | . |
| 7:14.304 | ולכן |  |
| 7:14.305 | שגו |  |
| 7:14.306 | ברואה |  |
| 7:14.307 | שלא |  |
| 7:14.308 | די |  |
| 7:14.309 | שיחסו |  |
| 7:14.310 | לשי"ת |  |
| 7:14.311 | רבוי |  |
| 7:14.312 | וחלוק |  |
| 7:14.313 | צורות |  |
| 7:14.314 | ופרצופים |  |
| 7:14.315 | שונים |  |
| 7:14.316 | וקוים | . |
| 7:14.317 | אלא |  |
| 7:14.318 | שיחסו |  |
| 7:14.319 | לו |  |
| 7:14.320 | שינוים |  |
| 7:14.321 | רבים |  |
| 7:14.322 | בעצמותו |  |
| 7:14.323 | ית | ' |
| 7:14.324 | ואמרו |  |
| 7:14.325 | שמתחלה |  |
| 7:14.326 | היה |  |
| 7:14.327 | מקום |  |
| 7:14.328 | וחלל |  |
| 7:14.329 | גדול |  |
| 7:14.330 | פנוי |  |
| 7:14.331 | וריק |  |
| 7:14.332 | להיות |  |
| 7:14.333 | שם |  |
| 7:14.334 | כל |  |
| 7:14.335 | ברואיו | , |
| 7:14.336 | ונעשה |  |
| 7:14.337 | ככדור |  |
| 7:14.338 | עגול |  |
| 7:14.339 | וחלול |  |
| 7:14.340 | באמצעותו | , |
| 7:14.341 | ושוב |  |
| 7:14.342 | נעשה |  |
| 7:14.343 | בו |  |
| 7:14.344 | שינוי |  |
| 7:14.345 | אחר | , |
| 7:14.346 | כי |  |
| 7:14.347 | המשיך |  |
| 7:14.348 | מאורו |  |
| 7:14.349 | הגדול |  |
| 7:14.350 | המקיף | , |
| 7:14.351 | ספירות |  |
| 7:14.352 | דעיגולי |  |
| 7:14.353 | אדם |  |
| 7:14.354 | קדמאה |  |
| 7:14.355 | סתימא | , |
| 7:14.356 | ואחר |  |
| 7:14.357 | זה |  |
| 7:14.358 | עיגולי |  |
| 7:14.359 | אדם |  |
| 7:14.360 | קדמון | . |
| 7:14.361 | ועיגולי |  |
| 7:14.362 | עתיק |  |
| 7:14.363 | יומין | , |
| 7:14.364 | ועיגולי |  |
| 7:14.365 | אריך |  |
| 7:14.366 | אנפין | , |
| 7:14.367 | ועיגולי |  |
| 7:14.368 | אבא |  |
| 7:14.369 | ואימא | . |
| 7:14.370 | גם |  |
| 7:14.371 | עיגולי |  |
| 7:14.372 | זעיר |  |
| 7:14.373 | ונוקביה | . |
| 7:14.374 | ושוב |  |
| 7:14.375 | נעשו |  |
| 7:14.376 | עוד |  |
| 7:14.377 | בו |  |
| 7:14.378 | שינויים |  |
| 7:14.379 | רבים |  |
| 7:14.380 | להמשיך |  |
| 7:14.381 | ג"כ |  |
| 7:14.382 | מעצמותו |  |
| 7:14.383 | ספירות |  |
| 7:14.384 | דיושר |  |
| 7:14.385 | בתנית |  |
| 7:14.386 | אדם | . |
| 7:14.387 | והם |  |
| 7:14.388 | עיקר |  |
| 7:14.389 | האלוהות |  |
| 7:14.390 | לפי |  |
| 7:14.391 | דעתם | . |
| 7:14.392 | תחלה |  |
| 7:14.393 | פרצוף |  |
| 7:14.394 | אדם |  |
| 7:14.395 | קדמון | . |
| 7:14.396 | ואחריו |  |
| 7:14.397 | פרצוף |  |
| 7:14.398 | עתיק |  |
| 7:14.399 | יומין | . |
| 7:14.400 | ואחריו |  |
| 7:14.401 | אריך |  |
| 7:14.402 | אנפין | . |
| 7:14.403 | ושניהם |  |
| 7:14.404 | כולם |  |
| 7:14.405 | זכר |  |
| 7:14.406 | ונקבה | . |
| 7:14.407 | רק |  |
| 7:14.408 | שזה |  |
| 7:14.409 | ימין |  |
| 7:14.410 | ושמאל | . |
| 7:14.411 | וזה |  |
| 7:14.412 | פנים |  |
| 7:14.413 | ואחור | , |
| 7:14.414 | ואחריהם |  |
| 7:14.415 | פרצופי |  |
| 7:14.416 | אבא |  |
| 7:14.417 | ואימא |  |
| 7:14.418 | זכר |  |
| 7:14.419 | ונקבה |  |
| 7:14.420 | נפרדים |  |
| 7:14.421 | זה |  |
| 7:14.422 | מזה | . |
| 7:14.423 | ואחריהם |  |
| 7:14.424 | פרצופי |  |
| 7:14.425 | זעיר |  |
| 7:14.426 | ונוקביה |  |
| 7:14.427 | ג"כ |  |
| 7:14.428 | נפרדים | . |
| 7:14.429 | והם |  |
| 7:14.430 | עיקר |  |
| 7:14.431 | האלוהות |  |
| 7:14.432 | אשר |  |
| 7:14.433 | לפי |  |
| 7:14.434 | דבריהם |  |
| 7:14.435 | אלהים |  |
| 7:14.436 | נקרא |  |
| 7:14.437 | בכל |  |
| 7:14.438 | תפלותינו | . |
| 7:14.439 | וכל |  |
| 7:14.440 | ברכותינו |  |
| 7:14.441 | אליהם |  |
| 7:14.442 | יכונו | . |
| 7:14.443 | כי |  |
| 7:14.444 | הם |  |
| 7:14.445 | אלוהינו |  |
| 7:14.446 | לפי |  |
| 7:14.447 | דעתם | . |
| 7:14.448 | הרי |  |
| 7:14.449 | נתבאר |  |
| 7:14.450 | לך |  |
| 7:14.451 | כמה |  |
| 7:14.452 | שינוים |  |
| 7:14.453 | וחליפות |  |
| 7:14.454 | נעשו |  |
| 7:14.455 | באלהינו |  |
| 7:14.456 | על |  |
| 7:14.457 | פי |  |
| 7:14.458 | הקבלה |  |
| 7:14.459 | החדשה |  |
| 7:14.460 | הסותרת |  |
| 7:14.461 | אמונת |  |
| 7:14.462 | תוה"ק | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

