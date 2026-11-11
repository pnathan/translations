# Translation pack: milhamot-hashem perek-2-12 (Milchamot ha-Shem 2:28–2:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-12`.

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
- **מדה / מדות / מידות**: context. Do not render it "sefirot" even where that
  is the referent.
  - The hermeneutical rules: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי"
    (1:6) → "the thirteen rules of Rabbi Ishmael and the thirty-two rules of
    Rabbi Yose the Galilean"; "שלש עשרה מדות שהתורה נדרשת בהם" (10:5) → "the
    thirteen rules by which the Torah is expounded". Add a note: literally
    "measures".
  - Exod 34:6–7: "שלש עשרה מדות של רחמים" (9:8) → "the thirteen attributes of
    mercy". These are a different thirteen from the rules. The author uses
    both, so never let one rendering cover the other.
  - Kabbalistic and philosophical attributes, God's or the sefirot's: "מדת
    הדין" → "the attribute of judgment"; "המדות העליונות" (2:11) → "the
    upper attributes"; "המידות הפנימיות" (2:7) → "the inner attributes";
    "מדת המלכות" → "the attribute of Malkhut"; "מלובש במידותיו" (5:15) →
    "clothed in His attributes"; "ממדותיו" (7:17, Exod 34:6) → "of His
    attributes".
  - Measure: "עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה" (4:32) → "ten
    measures and ten sefirot, because they have measure and limit"; "לא מקום
    ולא מידה" (7:13) → "neither place nor measure"; "במדה שמדדתם" →
    "with the measure with which you measured".
  - A quality or trait: "ממדת הענוה" (4:8) → "of the quality of humility";
    "המדה המגונה הזאת" (7:19) → "this disgraceful trait"; "אין לך מדה גדולה
    מזו" (6:14, Talmud) → "there is no greater virtue than this".
  - A proper name: מסכת מדות → "tractate Middot".
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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- המשרף (the questioner's pun on Metzaref ha-Emunah): keep "the Mesharef", with a
  note naming the pun.
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
- תשבורת: "mathematics" (mensuration); ותוכנים: "astronomy" (plural of תכונה).
- ייחדו said of dedicating a thing to one: "dedicate ... solely to"; the
  Kabbalistic "unify" in the note.
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


## Milchamot ha-Shem 2:28

Hebrew (Wikisource): תשובה לאדם מערכי לב ומה' מענה לשון. ויהי בראותי את מכתבך ישבתי משתומם ומתבהל. איך מלאך לבך לדרוש ולחקור אחר חבל גאונים, אשר הם בתורה נאמנים גדולי היחס כזרע האיתנים, ואחר כמה פסקנים וגם דרשנים ראשונים עם אחרונים חכמי מערב ואשכנז וספרד ותימנים. ושאר קהלות זרע קדש בני אמונים, זאת היא חמדתם מפז ורב פנינים. וכי תבוא תקלה מיד צדיקים נאמנים יודעי מדע חכמים ונבונים יראי אלהים וסרים מרע באונים. בחכמתם הפילו כמה צרים ומונים העומדים עליהם בכל זמנים. וכי קבלתם מרוקחת בהבל ושגיונים אשר לרוב חכמתם אושרו בפי שרים ורוזנים התיצבו לפני מלכים ושלטונים וכי ילכו בדרך לא טוב אבות ובנים? זה כמה מאות שנים? וכמה הם מופלגים בכל חכמה וספרנים מקובלים עם פילוסופיא והגיונים, רפואה ותשבורת ותוכנים. הוגים יומם ולילה שקטים ושאננים, אוכלים פרי מעשיהם הטוב דשנים ורעננים. חלילה! זה לא יאמן ולא יעלה על לב כי אם לפתיים, לכל דבר מאמינים, אשר על שכלם נשענים. ומן הראוי אין עלי להשיבך. כי מה אני להכנס לפני ולפנים מה שלא הורשתי להרהר כל שכן להשיב שואלים. אבל משום עת לעשות לה' אני משיבך ממה שמצאתי כתוב בדברי חז"ל. וזכותם יעמד לי שלא יארע שום תקלה על ידי. הנה מה שכתבת שאינך מבזה חכמי ישראל כבר יש עליך כמה עדים, לא תוכל להכחיש. ומה שהביא אותך לידי כך שאתה מיחס להם טעות ומינות. אלא מפני שהם בזויים בעיניך. אבל האדם יחוש לעצמו פן יכוה בגחלתן כי גדולים צדיקים במיתתן יותר מבחייהן כמו שאמרו בתלמוד. ומה שכתבת שאם לגלות מקום הסתירה וגם התלמוד מביאים ראיה לסתור דברי גדולים מהם, אין הנדון דומה לראיה, דהיכן מצינו שמביאים ראיה לסתור הלכה למשה מסיני? וחכמי הקבלה אומרים כך קיבלו מפי משה, ומי הוא זה ואיפה הוא שנתן לך רשות לחקור ולדרוש אחר דבריהם? ומה שכתבת שהם קורים לפרצופים אלוהות ח"ו שהם אומרים כך. ומה שכתבת צא ולמד מכוונת הברכות והתפלות. כבר ביארו שעילת כל העלות נתכנה בספירה פלונית בשם כן. ובספירה פלונית בשם כך. ומה שכתבת שביארו המקובלים שכמה עולמות יש אלפי אלפים וכו' אפילו יהיו מליוני מליונים כיון שהם אומ' שבוראם אחד והוא מחיה את כולם מה יש הפסד מזה? ומה שכתבת שהם בציור אדם בעל קומה, כבר אמרו גדול כחן של נביאים שמדמין דמות גבורה של מעלה לצורת אדם. ומה שכתבת ברמ"ח איברים כבר ביארו מה הן האיברים. כי הם רי"ו אותיות של שם בן ע"ב ול"ב נתיבות שהם רומזים לכוחות אלהיים. ומה שכתבת שכבר הודיתי שדעתי שאין שום תפלה שייכא בעילת כל העילות שקר אתה דובר. והלא מתחלת שאלתך אותי על דברי המשרף הישבתיך עיין בכס"א. אבל קבלתינו שהתפלה לא"ס ולהשפיע אל עשר ספירות כמו שכ' לחם שלמה. ומה שאמרתי לך לעיין בכסא אליהו כדי שתדע שדברי המשרף דברי כפירה באמת וראוי לבערו. וח"ו שנמצא בשום ספר כדבריו, אבל דברי כס"א אינם כדבריו ח"ו. כי דברי כסא אליהו אין הזעיר אנפין כי אם כלי לפעולת הבורא. ומה שכ' בדף כ"ה שבקראינו לז"א אנחנו קורין אל כל הפרצופים עם נשמה המסתתרת בהם שהוא אין סוף ב"ה. אין רוצה לומר שאין התפילה לא"ס חלילה, אלא רוצה לומר שעליית התפלה אל הא"ס ב"ה דרך זאת המדה. והראיה על זה מה שכ' בדף י"ז ע"ב וז"ל ומה שכתוב בכוונות התפלה והברכות לכוון בכל ברכה כוונה מיוחדת לספי' מיוחדת לא ח"ו לעצמות הספירה. כי הוא קיצוץ ח"ו, וחלילה לכוון לשום כח פרטי וספירה מיוחדת. אלא הכל לעצמות יחיד א"ס ב"ה כלל כל הכחות כולם. ועוד כ' בדף כ"ה ע"א וכל תפלותינו הם לאין סוף ית'. אך שאין אנחנו יכולים לכנותו בשום כנוי ושם. לזה אנחנו מתפללים על ידי הספירות. כי שם שייכות כל כל הכינוים והשמות. כי כל הכנויים והשמות הם שמות עצמותו המתפשט בספירות. וג"כ כ' בדף כ"ו ואם לא יעשו כן ולא ייחדו תפלתם אלא לז"א לא יהיו נענים. וכל המקובלים כיילי כללא דא וז"ל ומה שתמצא באיזה ספר מהם התפלה לשום ספירה חלילה שיהיה כן. אלא הכוונה שיכוין להעלותם אל הא"ס דרך אותה ספירה.

| id | word | punctuation after |
|---|---|---|
| 2:28.1 | תשובה |  |
| 2:28.2 | לאדם |  |
| 2:28.3 | מערכי |  |
| 2:28.4 | לב |  |
| 2:28.5 | ומה | ' |
| 2:28.6 | מענה |  |
| 2:28.7 | לשון | . |
| 2:28.8 | ויהי |  |
| 2:28.9 | בראותי |  |
| 2:28.10 | את |  |
| 2:28.11 | מכתבך |  |
| 2:28.12 | ישבתי |  |
| 2:28.13 | משתומם |  |
| 2:28.14 | ומתבהל | . |
| 2:28.15 | איך |  |
| 2:28.16 | מלאך |  |
| 2:28.17 | לבך |  |
| 2:28.18 | לדרוש |  |
| 2:28.19 | ולחקור |  |
| 2:28.20 | אחר |  |
| 2:28.21 | חבל |  |
| 2:28.22 | גאונים | , |
| 2:28.23 | אשר |  |
| 2:28.24 | הם |  |
| 2:28.25 | בתורה |  |
| 2:28.26 | נאמנים |  |
| 2:28.27 | גדולי |  |
| 2:28.28 | היחס |  |
| 2:28.29 | כזרע |  |
| 2:28.30 | האיתנים | , |
| 2:28.31 | ואחר |  |
| 2:28.32 | כמה |  |
| 2:28.33 | פסקנים |  |
| 2:28.34 | וגם |  |
| 2:28.35 | דרשנים |  |
| 2:28.36 | ראשונים |  |
| 2:28.37 | עם |  |
| 2:28.38 | אחרונים |  |
| 2:28.39 | חכמי |  |
| 2:28.40 | מערב |  |
| 2:28.41 | ואשכנז |  |
| 2:28.42 | וספרד |  |
| 2:28.43 | ותימנים | . |
| 2:28.44 | ושאר |  |
| 2:28.45 | קהלות |  |
| 2:28.46 | זרע |  |
| 2:28.47 | קדש |  |
| 2:28.48 | בני |  |
| 2:28.49 | אמונים | , |
| 2:28.50 | זאת |  |
| 2:28.51 | היא |  |
| 2:28.52 | חמדתם |  |
| 2:28.53 | מפז |  |
| 2:28.54 | ורב |  |
| 2:28.55 | פנינים | . |
| 2:28.56 | וכי |  |
| 2:28.57 | תבוא |  |
| 2:28.58 | תקלה |  |
| 2:28.59 | מיד |  |
| 2:28.60 | צדיקים |  |
| 2:28.61 | נאמנים |  |
| 2:28.62 | יודעי |  |
| 2:28.63 | מדע |  |
| 2:28.64 | חכמים |  |
| 2:28.65 | ונבונים |  |
| 2:28.66 | יראי |  |
| 2:28.67 | אלהים |  |
| 2:28.68 | וסרים |  |
| 2:28.69 | מרע |  |
| 2:28.70 | באונים | . |
| 2:28.71 | בחכמתם |  |
| 2:28.72 | הפילו |  |
| 2:28.73 | כמה |  |
| 2:28.74 | צרים |  |
| 2:28.75 | ומונים |  |
| 2:28.76 | העומדים |  |
| 2:28.77 | עליהם |  |
| 2:28.78 | בכל |  |
| 2:28.79 | זמנים | . |
| 2:28.80 | וכי |  |
| 2:28.81 | קבלתם |  |
| 2:28.82 | מרוקחת |  |
| 2:28.83 | בהבל |  |
| 2:28.84 | ושגיונים |  |
| 2:28.85 | אשר |  |
| 2:28.86 | לרוב |  |
| 2:28.87 | חכמתם |  |
| 2:28.88 | אושרו |  |
| 2:28.89 | בפי |  |
| 2:28.90 | שרים |  |
| 2:28.91 | ורוזנים |  |
| 2:28.92 | התיצבו |  |
| 2:28.93 | לפני |  |
| 2:28.94 | מלכים |  |
| 2:28.95 | ושלטונים |  |
| 2:28.96 | וכי |  |
| 2:28.97 | ילכו |  |
| 2:28.98 | בדרך |  |
| 2:28.99 | לא |  |
| 2:28.100 | טוב |  |
| 2:28.101 | אבות |  |
| 2:28.102 | ובנים | ? |
| 2:28.103 | זה |  |
| 2:28.104 | כמה |  |
| 2:28.105 | מאות |  |
| 2:28.106 | שנים | ? |
| 2:28.107 | וכמה |  |
| 2:28.108 | הם |  |
| 2:28.109 | מופלגים |  |
| 2:28.110 | בכל |  |
| 2:28.111 | חכמה |  |
| 2:28.112 | וספרנים |  |
| 2:28.113 | מקובלים |  |
| 2:28.114 | עם |  |
| 2:28.115 | פילוסופיא |  |
| 2:28.116 | והגיונים | , |
| 2:28.117 | רפואה |  |
| 2:28.118 | ותשבורת |  |
| 2:28.119 | ותוכנים | . |
| 2:28.120 | הוגים |  |
| 2:28.121 | יומם |  |
| 2:28.122 | ולילה |  |
| 2:28.123 | שקטים |  |
| 2:28.124 | ושאננים | , |
| 2:28.125 | אוכלים |  |
| 2:28.126 | פרי |  |
| 2:28.127 | מעשיהם |  |
| 2:28.128 | הטוב |  |
| 2:28.129 | דשנים |  |
| 2:28.130 | ורעננים | . |
| 2:28.131 | חלילה | ! |
| 2:28.132 | זה |  |
| 2:28.133 | לא |  |
| 2:28.134 | יאמן |  |
| 2:28.135 | ולא |  |
| 2:28.136 | יעלה |  |
| 2:28.137 | על |  |
| 2:28.138 | לב |  |
| 2:28.139 | כי |  |
| 2:28.140 | אם |  |
| 2:28.141 | לפתיים | , |
| 2:28.142 | לכל |  |
| 2:28.143 | דבר |  |
| 2:28.144 | מאמינים | , |
| 2:28.145 | אשר |  |
| 2:28.146 | על |  |
| 2:28.147 | שכלם |  |
| 2:28.148 | נשענים | . |
| 2:28.149 | ומן |  |
| 2:28.150 | הראוי |  |
| 2:28.151 | אין |  |
| 2:28.152 | עלי |  |
| 2:28.153 | להשיבך | . |
| 2:28.154 | כי |  |
| 2:28.155 | מה |  |
| 2:28.156 | אני |  |
| 2:28.157 | להכנס |  |
| 2:28.158 | לפני |  |
| 2:28.159 | ולפנים |  |
| 2:28.160 | מה |  |
| 2:28.161 | שלא |  |
| 2:28.162 | הורשתי |  |
| 2:28.163 | להרהר |  |
| 2:28.164 | כל |  |
| 2:28.165 | שכן |  |
| 2:28.166 | להשיב |  |
| 2:28.167 | שואלים | . |
| 2:28.168 | אבל |  |
| 2:28.169 | משום |  |
| 2:28.170 | עת |  |
| 2:28.171 | לעשות |  |
| 2:28.172 | לה | ' |
| 2:28.173 | אני |  |
| 2:28.174 | משיבך |  |
| 2:28.175 | ממה |  |
| 2:28.176 | שמצאתי |  |
| 2:28.177 | כתוב |  |
| 2:28.178 | בדברי |  |
| 2:28.179 | חז"ל | . |
| 2:28.180 | וזכותם |  |
| 2:28.181 | יעמד |  |
| 2:28.182 | לי |  |
| 2:28.183 | שלא |  |
| 2:28.184 | יארע |  |
| 2:28.185 | שום |  |
| 2:28.186 | תקלה |  |
| 2:28.187 | על |  |
| 2:28.188 | ידי | . |
| 2:28.189 | הנה |  |
| 2:28.190 | מה |  |
| 2:28.191 | שכתבת |  |
| 2:28.192 | שאינך |  |
| 2:28.193 | מבזה |  |
| 2:28.194 | חכמי |  |
| 2:28.195 | ישראל |  |
| 2:28.196 | כבר |  |
| 2:28.197 | יש |  |
| 2:28.198 | עליך |  |
| 2:28.199 | כמה |  |
| 2:28.200 | עדים | , |
| 2:28.201 | לא |  |
| 2:28.202 | תוכל |  |
| 2:28.203 | להכחיש | . |
| 2:28.204 | ומה |  |
| 2:28.205 | שהביא |  |
| 2:28.206 | אותך |  |
| 2:28.207 | לידי |  |
| 2:28.208 | כך |  |
| 2:28.209 | שאתה |  |
| 2:28.210 | מיחס |  |
| 2:28.211 | להם |  |
| 2:28.212 | טעות |  |
| 2:28.213 | ומינות | . |
| 2:28.214 | אלא |  |
| 2:28.215 | מפני |  |
| 2:28.216 | שהם |  |
| 2:28.217 | בזויים |  |
| 2:28.218 | בעיניך | . |
| 2:28.219 | אבל |  |
| 2:28.220 | האדם |  |
| 2:28.221 | יחוש |  |
| 2:28.222 | לעצמו |  |
| 2:28.223 | פן |  |
| 2:28.224 | יכוה |  |
| 2:28.225 | בגחלתן |  |
| 2:28.226 | כי |  |
| 2:28.227 | גדולים |  |
| 2:28.228 | צדיקים |  |
| 2:28.229 | במיתתן |  |
| 2:28.230 | יותר |  |
| 2:28.231 | מבחייהן |  |
| 2:28.232 | כמו |  |
| 2:28.233 | שאמרו |  |
| 2:28.234 | בתלמוד | . |
| 2:28.235 | ומה |  |
| 2:28.236 | שכתבת |  |
| 2:28.237 | שאם |  |
| 2:28.238 | לגלות |  |
| 2:28.239 | מקום |  |
| 2:28.240 | הסתירה |  |
| 2:28.241 | וגם |  |
| 2:28.242 | התלמוד |  |
| 2:28.243 | מביאים |  |
| 2:28.244 | ראיה |  |
| 2:28.245 | לסתור |  |
| 2:28.246 | דברי |  |
| 2:28.247 | גדולים |  |
| 2:28.248 | מהם | , |
| 2:28.249 | אין |  |
| 2:28.250 | הנדון |  |
| 2:28.251 | דומה |  |
| 2:28.252 | לראיה | , |
| 2:28.253 | דהיכן |  |
| 2:28.254 | מצינו |  |
| 2:28.255 | שמביאים |  |
| 2:28.256 | ראיה |  |
| 2:28.257 | לסתור |  |
| 2:28.258 | הלכה |  |
| 2:28.259 | למשה |  |
| 2:28.260 | מסיני | ? |
| 2:28.261 | וחכמי |  |
| 2:28.262 | הקבלה |  |
| 2:28.263 | אומרים |  |
| 2:28.264 | כך |  |
| 2:28.265 | קיבלו |  |
| 2:28.266 | מפי |  |
| 2:28.267 | משה | , |
| 2:28.268 | ומי |  |
| 2:28.269 | הוא |  |
| 2:28.270 | זה |  |
| 2:28.271 | ואיפה |  |
| 2:28.272 | הוא |  |
| 2:28.273 | שנתן |  |
| 2:28.274 | לך |  |
| 2:28.275 | רשות |  |
| 2:28.276 | לחקור |  |
| 2:28.277 | ולדרוש |  |
| 2:28.278 | אחר |  |
| 2:28.279 | דבריהם | ? |
| 2:28.280 | ומה |  |
| 2:28.281 | שכתבת |  |
| 2:28.282 | שהם |  |
| 2:28.283 | קורים |  |
| 2:28.284 | לפרצופים |  |
| 2:28.285 | אלוהות |  |
| 2:28.286 | ח"ו |  |
| 2:28.287 | שהם |  |
| 2:28.288 | אומרים |  |
| 2:28.289 | כך | . |
| 2:28.290 | ומה |  |
| 2:28.291 | שכתבת |  |
| 2:28.292 | צא |  |
| 2:28.293 | ולמד |  |
| 2:28.294 | מכוונת |  |
| 2:28.295 | הברכות |  |
| 2:28.296 | והתפלות | . |
| 2:28.297 | כבר |  |
| 2:28.298 | ביארו |  |
| 2:28.299 | שעילת |  |
| 2:28.300 | כל |  |
| 2:28.301 | העלות |  |
| 2:28.302 | נתכנה |  |
| 2:28.303 | בספירה |  |
| 2:28.304 | פלונית |  |
| 2:28.305 | בשם |  |
| 2:28.306 | כן | . |
| 2:28.307 | ובספירה |  |
| 2:28.308 | פלונית |  |
| 2:28.309 | בשם |  |
| 2:28.310 | כך | . |
| 2:28.311 | ומה |  |
| 2:28.312 | שכתבת |  |
| 2:28.313 | שביארו |  |
| 2:28.314 | המקובלים |  |
| 2:28.315 | שכמה |  |
| 2:28.316 | עולמות |  |
| 2:28.317 | יש |  |
| 2:28.318 | אלפי |  |
| 2:28.319 | אלפים |  |
| 2:28.320 | וכו | ' |
| 2:28.321 | אפילו |  |
| 2:28.322 | יהיו |  |
| 2:28.323 | מליוני |  |
| 2:28.324 | מליונים |  |
| 2:28.325 | כיון |  |
| 2:28.326 | שהם |  |
| 2:28.327 | אומ | ' |
| 2:28.328 | שבוראם |  |
| 2:28.329 | אחד |  |
| 2:28.330 | והוא |  |
| 2:28.331 | מחיה |  |
| 2:28.332 | את |  |
| 2:28.333 | כולם |  |
| 2:28.334 | מה |  |
| 2:28.335 | יש |  |
| 2:28.336 | הפסד |  |
| 2:28.337 | מזה | ? |
| 2:28.338 | ומה |  |
| 2:28.339 | שכתבת |  |
| 2:28.340 | שהם |  |
| 2:28.341 | בציור |  |
| 2:28.342 | אדם |  |
| 2:28.343 | בעל |  |
| 2:28.344 | קומה | , |
| 2:28.345 | כבר |  |
| 2:28.346 | אמרו |  |
| 2:28.347 | גדול |  |
| 2:28.348 | כחן |  |
| 2:28.349 | של |  |
| 2:28.350 | נביאים |  |
| 2:28.351 | שמדמין |  |
| 2:28.352 | דמות |  |
| 2:28.353 | גבורה |  |
| 2:28.354 | של |  |
| 2:28.355 | מעלה |  |
| 2:28.356 | לצורת |  |
| 2:28.357 | אדם | . |
| 2:28.358 | ומה |  |
| 2:28.359 | שכתבת |  |
| 2:28.360 | ברמ"ח |  |
| 2:28.361 | איברים |  |
| 2:28.362 | כבר |  |
| 2:28.363 | ביארו |  |
| 2:28.364 | מה |  |
| 2:28.365 | הן |  |
| 2:28.366 | האיברים | . |
| 2:28.367 | כי |  |
| 2:28.368 | הם |  |
| 2:28.369 | רי"ו |  |
| 2:28.370 | אותיות |  |
| 2:28.371 | של |  |
| 2:28.372 | שם |  |
| 2:28.373 | בן |  |
| 2:28.374 | ע"ב |  |
| 2:28.375 | ול"ב |  |
| 2:28.376 | נתיבות |  |
| 2:28.377 | שהם |  |
| 2:28.378 | רומזים |  |
| 2:28.379 | לכוחות |  |
| 2:28.380 | אלהיים | . |
| 2:28.381 | ומה |  |
| 2:28.382 | שכתבת |  |
| 2:28.383 | שכבר |  |
| 2:28.384 | הודיתי |  |
| 2:28.385 | שדעתי |  |
| 2:28.386 | שאין |  |
| 2:28.387 | שום |  |
| 2:28.388 | תפלה |  |
| 2:28.389 | שייכא |  |
| 2:28.390 | בעילת |  |
| 2:28.391 | כל |  |
| 2:28.392 | העילות |  |
| 2:28.393 | שקר |  |
| 2:28.394 | אתה |  |
| 2:28.395 | דובר | . |
| 2:28.396 | והלא |  |
| 2:28.397 | מתחלת |  |
| 2:28.398 | שאלתך |  |
| 2:28.399 | אותי |  |
| 2:28.400 | על |  |
| 2:28.401 | דברי |  |
| 2:28.402 | המשרף |  |
| 2:28.403 | הישבתיך |  |
| 2:28.404 | עיין |  |
| 2:28.405 | בכס"א | . |
| 2:28.406 | אבל |  |
| 2:28.407 | קבלתינו |  |
| 2:28.408 | שהתפלה |  |
| 2:28.409 | לא"ס |  |
| 2:28.410 | ולהשפיע |  |
| 2:28.411 | אל |  |
| 2:28.412 | עשר |  |
| 2:28.413 | ספירות |  |
| 2:28.414 | כמו |  |
| 2:28.415 | שכ | ' |
| 2:28.416 | לחם |  |
| 2:28.417 | שלמה | . |
| 2:28.418 | ומה |  |
| 2:28.419 | שאמרתי |  |
| 2:28.420 | לך |  |
| 2:28.421 | לעיין |  |
| 2:28.422 | בכסא |  |
| 2:28.423 | אליהו |  |
| 2:28.424 | כדי |  |
| 2:28.425 | שתדע |  |
| 2:28.426 | שדברי |  |
| 2:28.427 | המשרף |  |
| 2:28.428 | דברי |  |
| 2:28.429 | כפירה |  |
| 2:28.430 | באמת |  |
| 2:28.431 | וראוי |  |
| 2:28.432 | לבערו | . |
| 2:28.433 | וח"ו |  |
| 2:28.434 | שנמצא |  |
| 2:28.435 | בשום |  |
| 2:28.436 | ספר |  |
| 2:28.437 | כדבריו | , |
| 2:28.438 | אבל |  |
| 2:28.439 | דברי |  |
| 2:28.440 | כס"א |  |
| 2:28.441 | אינם |  |
| 2:28.442 | כדבריו |  |
| 2:28.443 | ח"ו | . |
| 2:28.444 | כי |  |
| 2:28.445 | דברי |  |
| 2:28.446 | כסא |  |
| 2:28.447 | אליהו |  |
| 2:28.448 | אין |  |
| 2:28.449 | הזעיר |  |
| 2:28.450 | אנפין |  |
| 2:28.451 | כי |  |
| 2:28.452 | אם |  |
| 2:28.453 | כלי |  |
| 2:28.454 | לפעולת |  |
| 2:28.455 | הבורא | . |
| 2:28.456 | ומה |  |
| 2:28.457 | שכ | ' |
| 2:28.458 | בדף |  |
| 2:28.459 | כ"ה |  |
| 2:28.460 | שבקראינו |  |
| 2:28.461 | לז"א |  |
| 2:28.462 | אנחנו |  |
| 2:28.463 | קורין |  |
| 2:28.464 | אל |  |
| 2:28.465 | כל |  |
| 2:28.466 | הפרצופים |  |
| 2:28.467 | עם |  |
| 2:28.468 | נשמה |  |
| 2:28.469 | המסתתרת |  |
| 2:28.470 | בהם |  |
| 2:28.471 | שהוא |  |
| 2:28.472 | אין |  |
| 2:28.473 | סוף |  |
| 2:28.474 | ב"ה | . |
| 2:28.475 | אין |  |
| 2:28.476 | רוצה |  |
| 2:28.477 | לומר |  |
| 2:28.478 | שאין |  |
| 2:28.479 | התפילה |  |
| 2:28.480 | לא"ס |  |
| 2:28.481 | חלילה | , |
| 2:28.482 | אלא |  |
| 2:28.483 | רוצה |  |
| 2:28.484 | לומר |  |
| 2:28.485 | שעליית |  |
| 2:28.486 | התפלה |  |
| 2:28.487 | אל |  |
| 2:28.488 | הא"ס |  |
| 2:28.489 | ב"ה |  |
| 2:28.490 | דרך |  |
| 2:28.491 | זאת |  |
| 2:28.492 | המדה | . |
| 2:28.493 | והראיה |  |
| 2:28.494 | על |  |
| 2:28.495 | זה |  |
| 2:28.496 | מה |  |
| 2:28.497 | שכ | ' |
| 2:28.498 | בדף |  |
| 2:28.499 | י"ז |  |
| 2:28.500 | ע"ב |  |
| 2:28.501 | וז"ל |  |
| 2:28.502 | ומה |  |
| 2:28.503 | שכתוב |  |
| 2:28.504 | בכוונות |  |
| 2:28.505 | התפלה |  |
| 2:28.506 | והברכות |  |
| 2:28.507 | לכוון |  |
| 2:28.508 | בכל |  |
| 2:28.509 | ברכה |  |
| 2:28.510 | כוונה |  |
| 2:28.511 | מיוחדת |  |
| 2:28.512 | לספי | ' |
| 2:28.513 | מיוחדת |  |
| 2:28.514 | לא |  |
| 2:28.515 | ח"ו |  |
| 2:28.516 | לעצמות |  |
| 2:28.517 | הספירה | . |
| 2:28.518 | כי |  |
| 2:28.519 | הוא |  |
| 2:28.520 | קיצוץ |  |
| 2:28.521 | ח"ו | , |
| 2:28.522 | וחלילה |  |
| 2:28.523 | לכוון |  |
| 2:28.524 | לשום |  |
| 2:28.525 | כח |  |
| 2:28.526 | פרטי |  |
| 2:28.527 | וספירה |  |
| 2:28.528 | מיוחדת | . |
| 2:28.529 | אלא |  |
| 2:28.530 | הכל |  |
| 2:28.531 | לעצמות |  |
| 2:28.532 | יחיד |  |
| 2:28.533 | א"ס |  |
| 2:28.534 | ב"ה |  |
| 2:28.535 | כלל |  |
| 2:28.536 | כל |  |
| 2:28.537 | הכחות |  |
| 2:28.538 | כולם | . |
| 2:28.539 | ועוד |  |
| 2:28.540 | כ | ' |
| 2:28.541 | בדף |  |
| 2:28.542 | כ"ה |  |
| 2:28.543 | ע"א |  |
| 2:28.544 | וכל |  |
| 2:28.545 | תפלותינו |  |
| 2:28.546 | הם |  |
| 2:28.547 | לאין |  |
| 2:28.548 | סוף |  |
| 2:28.549 | ית | '. |
| 2:28.550 | אך |  |
| 2:28.551 | שאין |  |
| 2:28.552 | אנחנו |  |
| 2:28.553 | יכולים |  |
| 2:28.554 | לכנותו |  |
| 2:28.555 | בשום |  |
| 2:28.556 | כנוי |  |
| 2:28.557 | ושם | . |
| 2:28.558 | לזה |  |
| 2:28.559 | אנחנו |  |
| 2:28.560 | מתפללים |  |
| 2:28.561 | על |  |
| 2:28.562 | ידי |  |
| 2:28.563 | הספירות | . |
| 2:28.564 | כי |  |
| 2:28.565 | שם |  |
| 2:28.566 | שייכות |  |
| 2:28.567 | כל |  |
| 2:28.568 | כל |  |
| 2:28.569 | הכינוים |  |
| 2:28.570 | והשמות | . |
| 2:28.571 | כי |  |
| 2:28.572 | כל |  |
| 2:28.573 | הכנויים |  |
| 2:28.574 | והשמות |  |
| 2:28.575 | הם |  |
| 2:28.576 | שמות |  |
| 2:28.577 | עצמותו |  |
| 2:28.578 | המתפשט |  |
| 2:28.579 | בספירות | . |
| 2:28.580 | וג"כ |  |
| 2:28.581 | כ | ' |
| 2:28.582 | בדף |  |
| 2:28.583 | כ"ו |  |
| 2:28.584 | ואם |  |
| 2:28.585 | לא |  |
| 2:28.586 | יעשו |  |
| 2:28.587 | כן |  |
| 2:28.588 | ולא |  |
| 2:28.589 | ייחדו |  |
| 2:28.590 | תפלתם |  |
| 2:28.591 | אלא |  |
| 2:28.592 | לז"א |  |
| 2:28.593 | לא |  |
| 2:28.594 | יהיו |  |
| 2:28.595 | נענים | . |
| 2:28.596 | וכל |  |
| 2:28.597 | המקובלים |  |
| 2:28.598 | כיילי |  |
| 2:28.599 | כללא |  |
| 2:28.600 | דא |  |
| 2:28.601 | וז"ל |  |
| 2:28.602 | ומה |  |
| 2:28.603 | שתמצא |  |
| 2:28.604 | באיזה |  |
| 2:28.605 | ספר |  |
| 2:28.606 | מהם |  |
| 2:28.607 | התפלה |  |
| 2:28.608 | לשום |  |
| 2:28.609 | ספירה |  |
| 2:28.610 | חלילה |  |
| 2:28.611 | שיהיה |  |
| 2:28.612 | כן | . |
| 2:28.613 | אלא |  |
| 2:28.614 | הכוונה |  |
| 2:28.615 | שיכוין |  |
| 2:28.616 | להעלותם |  |
| 2:28.617 | אל |  |
| 2:28.618 | הא"ס |  |
| 2:28.619 | דרך |  |
| 2:28.620 | אותה |  |
| 2:28.621 | ספירה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

