# Review pack: milhamot-hashem perek-2-12 (Milchamot ha-Shem 2:28–2:28)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-12`.

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


### Translation under review (2:28)

English: Reply. To man the arrangements of the heart, and from the LORD the answer of the tongue. And it was when I saw your letter, I sat astonished and dismayed. How your heart came to inquire and investigate after the band of the Geonim, who are faithful in the Torah, distinguished in lineage, like the seed of the mighty, and after several decisors and also expounders, the early and the later, the sages of the West, Ashkenaz, Sefarad and the Yemenites. and the rest of the congregations, the holy seed, the faithful children, this is their treasure, more than gold and abundant pearls. Can a stumbling come from the hand of righteous and faithful ones, who know knowledge, the wise and discerning, who fear God and turn aside from evil with their might. By their wisdom they felled many foes and multitudes who stood against them in every age. Can it be that their Kabbalah is compounded in vanity and error, which, because of their great wisdom, were approved by the mouths of princes and nobles and stood before kings and rulers, and can fathers and sons walk on a path that is not good? How many hundreds of years is this? And how many of them are distinguished in every wisdom, and scholars, Kabbalists with philosophy and logic, medicine, tishburet and computations. meditating day and night, quiet and at ease, eating the fruit of their deeds, the good, fat and lush. Heaven forbid! This is not to be believed and does not come to mind but only for the simple, who believe every matter, who lean on their understanding. And properly, it is not incumbent on me to answer you. For what am I, to enter before and within, into what I was not permitted to ponder, all the more so to answer questioners. But because it is a time to act for the LORD, I answer you from what I found written in the words of our Sages of blessed memory. May their merit stand for me, that no stumbling befall through me. Now, as for what you wrote, that you do not despise the sages of Israel, there are already several witnesses against you; you cannot deny it. And what brought you to this that you attribute to them error and heresy. Rather, because they are despised in your eyes. But a person should take heed for himself, lest he be burned by their coals, for the righteous are greater in their death than in their life, as they said in the Talmud. And as for what you wrote, that if one reveals the place of contradiction, and the Talmud also brings proof to refute the words of those greater than them, the matter is not comparable to a proof, for where do we find that they bring proof to refute a halakhah given to Moses at Sinai? And the masters of Kabbalah say that they received thus from the mouth of Moses, and who is this one, and where is he, who gave you permission to inquire and investigate after their words? And as for what you wrote, that they call the partzufim divinities, God forbid, that they say thus. And as for what you wrote, go and learn from the intentions of the blessings and the prayers. They already explained that the cause of all causes is designated in a certain sefirah by such a name. And in a certain sefirah by that name. And as for what you wrote, that the Kabbalists explained that there are several worlds, thousands of thousands, etc., even if they were millions of millions, since they say that their Creator is one, and He gives life to all of them: what loss is there from this? And as for what you wrote, that they are in the image of a man of stature, already they said, "Great is the power of the prophets, who liken the form of might above to the shape of a man." And as for what you wrote about 248 limbs, they already explained what the limbs are. For they are 216 letters of the Name of 72 and 32 paths, which allude to divine powers. And as for what you wrote, that I already conceded that my opinion is that there is no prayer applicable to the Cause of all causes, you speak falsehood. But surely, from the beginning of your question to me about the words of the Mesharef, I answered you, see in Kisse Eliyahu. But our Kabbalah holds that prayer is to Ein Sof and bestows upon the ten sefirot, as Lechem Shelomo wrote. And as for what I told you, to consult Kisse Eliyahu, so that you would know that the words of the Mesharef are words of heresy, in truth, and fitting to root it out. And God forbid that there be found in any book words like his, but the words of Kisse Eliyahu are not like his, God forbid. For the words of Kisse Eliyahu hold that Ze'ir Anpin is nothing but a vessel for the action of the Creator. And as for what he wrote on folio 25, that when we call Ze'ir Anpin God, we call all the partzufim God, with a soul hidden in them, which is Ein Sof, blessed be He. This does not mean that prayer is not directed to Ein Sof, Heaven forbid, but it means that the ascent of prayer to Ein Sof, blessed be He, is by way of this attribute. And the proof for this is what he wrote on folio 17b, and these are his words: "As for what is written about the intentions of prayer and the blessings, to direct in each blessing a particular intention, to a particular sefirah, God forbid, not to the essence of the sefirah. For it is a cutting off, God forbid, and Heaven forbid to direct it to any particular power or particular sefirah. But all is to the essence of the one Ein Sof, blessed be He, comprising all the powers, all of them." And furthermore he wrote on folio 25a: "And all our prayers are to Ein Sof, may He be blessed. But that we cannot call Him by any appellation and name. Therefore we pray through the sefirot. For the name pertaining to all, all the appellations and names. For all the appellations and names are names of His essence, extending through the sefirot." Also he wrote on folio 26: "if they do not do so, and do not unify their prayer but to Ze'ir Anpin, they will not be answered." And all the Kabbalists hold this principle, and these are their words: "As for whatever you find in some book of theirs, that prayer is addressed to any sefirah, Heaven forbid that it be so. But the intention is that one direct them to raise them to Ein Sof by way of that sefirah."

Literal: Reply to-man the-arrangements-of the-heart and-from-the-LORD the-answer-of the-tongue and-it-was when-I-saw (object-marker) your-letter I-sat astonished and-dismayed how has-filled your-heart to-inquire and-to-investigate after band of-the-Geonim who they-are in-the-Torah faithful distinguished-of lineage like-seed-of the-mighty-ones and-after several decisors and-also expounders the-early with the-later sages-of the-West and-Ashkenaz and-Sefarad and-the-Yemenites and-the-rest-of congregations seed holy children-of faithful this is their-treasure more-than-gold and-abundant pearls can-it-be-that comes a-stumbling from-the-hand-of righteous-ones faithful-ones who-know knowledge the-wise and-discerning who-fear God and-who-turn-aside from-evil with-their-might by-their-wisdom they-felled many foes and-multitudes who-stood against-them in-every age can-it-be-that their-Kabbalah compounded in-vanity and-error which because-of-the-great their-wisdom were-approved by-the-mouths-of princes and-nobles they-stood before kings and-rulers can-it-be-that they-walk on-the-path not good fathers and-sons this how-many hundreds of-years and-how-many they are-distinguished in-every wisdom and-scholars Kabbalists with philosophy and-logic medicine and-tishburet and-computations meditating by-day and-by-night quiet and-at-ease eating the-fruit-of their-deeds the-good fat and-lush Heaven-forbid this is-not to-be-believed and-not come-up upon the-heart but only for-the-simple to-every matter who-believe who upon their-understanding lean and-properly what-is-fitting there-is-not upon-me to-answer-you for what am-I to-enter before and-within what that-not I-was-permitted to-ponder all the-more-so to-answer questioners but because-of a-time to-act for-the-LORD I answer-you from-what I-found written in-the-words-of our-Sages-of-blessed-memory and-their-merit will-stand for-me that-not will-befall any stumbling through my-hand behold as-for-what you-wrote that-you-do-not despise the-sages-of Israel already there-are against-you several witnesses not you-can to-deny and-what brought you to this that-you attribute to-them error and-heresy but because that-they-are despised in-your-eyes but the-person should-take-heed for-himself lest he-be-burned by-their-coals for great-ones righteous-ones in-their-death more than-in-their-life as they-said in-the-Talmud and-as-for-what you-wrote that-if to-reveal the-place-of the-contradiction and-also the-Talmud brings proof to-refute the-words-of greater than-them there-is-not the-matter comparable to-a-proof for-where do-we-find that-they-bring proof to-refute a-halakhah given-to-Moses at-Sinai and-the-masters-of Kabbalah say thus they-received from-the-mouth-of Moses and-who is-he this-one and-where is-he who-gave you permission to-inquire and-to-investigate after their-words and-as-for-what you-wrote that-they call the-partzufim divinities God-forbid that-they say thus and-as-for-what you-wrote go-out and-learn from-the-intentions-of the-blessings and-the-prayers already they-explained that-the-cause-of all the-causes is-designated in-a-sefirah a-certain by-the-name such and-in-a-sefirah a-certain by-the-name that and-as-for-what you-wrote that-they-explained the-Kabbalists that-several worlds there-are thousands-of thousands etc. even-if they-were millions-of millions since that-they say that-their-Creator is-one and-He gives-life-to (object-marker) all-of-them what is-there loss from-this and-as-for-what you-wrote that-they in-the-image-of a-man of stature already they-said great is-their-power of the-prophets who-liken the-form of-might of above to-the-shape-of a-man and-as-for-what you-wrote about-248 limbs already they-explained what they-are the-limbs for they-are 216 letters of the-Name of 72 and-32 paths which-are they-allude-to the-powers-of divine and-as-for-what you-wrote that-already I-conceded that-my-opinion that-there-is-not any prayer applicable to-the-cause-of all causes falsehood you speak but-surely from-the-beginning-of your-question to-me about the-words-of the-Mesharef I-answered-you see in-Kisse-Eliyahu but our-Kabbalah that-prayer to-Ein-Sof and-to-bestow upon ten sefirot as he-wrote Lechem Shelomo and-as-for-what I-said to-you to-consult in-the-Throne of-Elijah so-that you-would-know that-the-words-of the-Mesharef are-the-words-of heresy in-truth and-fitting to-root-it-out and-God-forbid that-there-be-found in-any book like-his-words but the-words-of Kisse-Eliyahu are-not like-his-words God-forbid for the-words-of the-Throne of-Elijah is-nothing Ze'ir Anpin but only a-vessel for-the-action-of the-Creator and-as-for-what he-wrote on-folio 25 that-when-we call-Ze'ir-Anpin we call God all the-partzufim with a-soul hidden in-them which-is Ein Sof blessed-be-He this-does-not mean to-say that-not prayer to-Ein-Sof Heaven-forbid but it-means to-say the-ascent-of prayer to the-Ein-Sof blessed-be-He by-way-of this attribute and-the-proof for this is-what he-wrote on-folio 17 b and-these-are-his-words: as-for-what is-written about-the-intentions-of prayer and-the-blessings to-direct in-each blessing an-intention particular to-a-sefirah particular not God-forbid to-the-essence-of the-sefirah for it a-cutting-off God-forbid and-Heaven-forbid to-direct to-any power particular and-sefirah particular but all to-the-essence-of the-one Ein-Sof blessed-be-He comprising all the-powers all-of-them and-furthermore he-wrote on-folio 25 a and-all our-prayers are to-Ein Sof may-He-be-blessed but that-not we can call-Him by-any appellation and-name therefore we pray by-means-of the-hand-of the-sefirot for the-name pertaining all all the-appellations and-names for all the-appellations and-the-names are names of-His-essence that-extends through-the-sefirot also he-wrote on-folio 26 if not they-do so and-not unify their-prayer but to-Ze'ir-Anpin not they-will-be answered and-all the-Kabbalists all-hold the-principle this and-these-are-their-words: as-for-whatever you-find in-some book of-theirs prayer to-any sefirah Heaven-forbid that-it-be so but the-intention that-he-direct to-raise-them to Ein-Sof by-way-of that sefirah

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:28.1 | Reply |  |  | תשובה: the heading of a reply in the correspondence of chapter 2. |
| 2:28.2 | to man |  |  |  |
| 2:28.3 | the arrangements of |  |  | מערכי: construct plural, "arrangements of"; the opening alludes to Prov 16:1 (see apparatus). |
| 2:28.4 | the heart |  |  |  |
| 2:28.5 | and from the LORD |  |  | ומה': "and from the LORD"; the apostrophe is the Tetragrammaton's mark in the punctuation column. |
| 2:28.6 | the answer of |  |  |  |
| 2:28.7 | the tongue |  |  |  |
| 2:28.8 | and it was |  |  | ויהי: a wayyiqtol opening the sentence; the narrative runs on in the verbs that follow. |
| 2:28.9 | when I saw |  |  |  |
| 2:28.10 | (object marker) |  |  | את: the object marker, with no English counterpart. |
| 2:28.11 | your letter |  |  |  |
| 2:28.12 | I sat |  |  |  |
| 2:28.13 | astonished |  |  |  |
| 2:28.14 | and dismayed |  |  |  |
| 2:28.15 | how |  |  |  |
| 2:28.16 | has filled |  |  |  |
| 2:28.17 | your heart |  |  |  |
| 2:28.18 | to inquire |  |  |  |
| 2:28.19 | and to investigate |  |  |  |
| 2:28.20 | after |  |  |  |
| 2:28.21 | band |  |  | חבל: "cord, band"; here a company or fellowship. |
| 2:28.22 | of the Geonim |  |  | גאונים: the Geonim, the heads of the Babylonian academies. |
| 2:28.23 | who |  |  |  |
| 2:28.24 | they are |  |  |  |
| 2:28.25 | in the Torah |  |  |  |
| 2:28.26 | faithful |  |  |  |
| 2:28.27 | distinguished of |  |  |  |
| 2:28.28 | lineage |  |  | היחס: "lineage, pedigree"; גדולי היחס: men of distinguished lineage. |
| 2:28.29 | like seed of |  |  |  |
| 2:28.30 | the mighty ones |  |  |  |
| 2:28.31 | and after |  |  |  |
| 2:28.32 | several |  |  |  |
| 2:28.33 | decisors |  |  |  |
| 2:28.34 | and also |  |  |  |
| 2:28.35 | expounders |  |  |  |
| 2:28.36 | the early |  |  |  |
| 2:28.37 | with |  |  |  |
| 2:28.38 | the later |  |  |  |
| 2:28.39 | sages of |  |  |  |
| 2:28.40 | the West |  |  | מערב: literally "the West"; the place is not named further. |
| 2:28.41 | and Ashkenaz |  |  |  |
| 2:28.42 | and Sefarad |  |  |  |
| 2:28.43 | and the Yemenites |  |  |  |
| 2:28.44 | and the rest of |  |  |  |
| 2:28.45 | congregations |  |  |  |
| 2:28.46 | seed |  |  |  |
| 2:28.47 | holy |  |  |  |
| 2:28.48 | children of |  |  |  |
| 2:28.49 | faithful |  |  |  |
| 2:28.50 | this |  |  |  |
| 2:28.51 | is |  |  |  |
| 2:28.52 | their treasure |  |  | חמדתם: "their delight, treasure", the noun of desire. |
| 2:28.53 | more than gold |  |  |  |
| 2:28.54 | and abundant |  |  |  |
| 2:28.55 | pearls |  |  |  |
| 2:28.56 | can it be that |  |  | וכי: the interrogative of a rhetorical question: "Can it be that ...". |
| 2:28.57 | comes |  |  |  |
| 2:28.58 | a stumbling |  |  |  |
| 2:28.59 | from the hand of |  |  |  |
| 2:28.60 | righteous ones |  |  |  |
| 2:28.61 | faithful ones |  |  |  |
| 2:28.62 | who know |  |  |  |
| 2:28.63 | knowledge |  |  |  |
| 2:28.64 | the wise |  |  |  |
| 2:28.65 | and discerning |  |  |  |
| 2:28.66 | who fear |  |  |  |
| 2:28.67 | God |  |  | אלהים: "God"; the allusion is to Job 1:1 (see apparatus). |
| 2:28.68 | and who turn aside |  |  |  |
| 2:28.69 | from evil |  |  |  |
| 2:28.70 | with their might |  |  |  |
| 2:28.71 | by their wisdom |  |  |  |
| 2:28.72 | they felled |  |  |  |
| 2:28.73 | many |  |  |  |
| 2:28.74 | foes |  |  |  |
| 2:28.75 | and multitudes |  |  | ומונים: sense uncertain; rendered "and multitudes". |
| 2:28.76 | who stood |  |  |  |
| 2:28.77 | against them |  |  |  |
| 2:28.78 | in every |  |  |  |
| 2:28.79 | age |  |  |  |
| 2:28.80 | can it be that |  |  |  |
| 2:28.81 | their Kabbalah |  |  | קבלתם: "their Kabbalah", the lore of the Kabbalists, not the rabbinic chain of tradition. |
| 2:28.82 | compounded |  |  |  |
| 2:28.83 | in vanity |  |  |  |
| 2:28.84 | and error |  |  |  |
| 2:28.85 | which |  |  |  |
| 2:28.86 | because of the great |  |  |  |
| 2:28.87 | their wisdom |  |  |  |
| 2:28.88 | were approved |  |  |  |
| 2:28.89 | by the mouths of |  |  |  |
| 2:28.90 | princes |  |  |  |
| 2:28.91 | and nobles |  |  |  |
| 2:28.92 | they stood |  |  |  |
| 2:28.93 | before |  |  |  |
| 2:28.94 | kings |  |  |  |
| 2:28.95 | and rulers |  |  |  |
| 2:28.96 | can it be that |  |  |  |
| 2:28.97 | they walk |  |  |  |
| 2:28.98 | on the path |  |  |  |
| 2:28.99 | not |  |  |  |
| 2:28.100 | good |  |  |  |
| 2:28.101 | fathers |  |  |  |
| 2:28.102 | and sons |  |  |  |
| 2:28.103 | this |  |  |  |
| 2:28.104 | how many |  |  |  |
| 2:28.105 | hundreds |  |  |  |
| 2:28.106 | of years |  |  |  |
| 2:28.107 | and how many |  |  |  |
| 2:28.108 | they |  |  |  |
| 2:28.109 | are distinguished |  |  |  |
| 2:28.110 | in every |  |  |  |
| 2:28.111 | wisdom |  |  |  |
| 2:28.112 | and scholars |  |  |  |
| 2:28.113 | Kabbalists |  |  |  |
| 2:28.114 | with |  |  |  |
| 2:28.115 | philosophy |  |  |  |
| 2:28.116 | and logic |  |  |  |
| 2:28.117 | medicine |  |  |  |
| 2:28.118 | and tishburet |  |  | ותשבורת: a technical term, transliterated; its sense is not settled here. |
| 2:28.119 | and computations |  |  | ותוכנים: read as "computations"; the sense is uncertain. |
| 2:28.120 | meditating |  |  |  |
| 2:28.121 | by day |  |  |  |
| 2:28.122 | and by night |  |  |  |
| 2:28.123 | quiet |  |  |  |
| 2:28.124 | and at ease |  |  |  |
| 2:28.125 | eating |  |  |  |
| 2:28.126 | the fruit of |  |  |  |
| 2:28.127 | their deeds |  |  |  |
| 2:28.128 | the good |  |  |  |
| 2:28.129 | fat |  |  |  |
| 2:28.130 | and lush |  |  |  |
| 2:28.131 | Heaven forbid |  |  | חלילה: "Heaven forbid", the full form of ח"ו. |
| 2:28.132 | this |  |  |  |
| 2:28.133 | is not |  |  |  |
| 2:28.134 | to be believed |  |  |  |
| 2:28.135 | and not |  |  |  |
| 2:28.136 | come up |  |  |  |
| 2:28.137 | upon |  |  |  |
| 2:28.138 | the heart |  |  |  |
| 2:28.139 | but |  |  |  |
| 2:28.140 | only |  |  |  |
| 2:28.141 | for the simple |  |  |  |
| 2:28.142 | to every |  |  |  |
| 2:28.143 | matter |  |  |  |
| 2:28.144 | who believe |  |  |  |
| 2:28.145 | who |  |  |  |
| 2:28.146 | upon |  |  |  |
| 2:28.147 | their understanding |  |  |  |
| 2:28.148 | lean |  |  |  |
| 2:28.149 | and properly |  |  |  |
| 2:28.150 | what is fitting |  |  |  |
| 2:28.151 | there is not |  |  |  |
| 2:28.152 | upon me |  |  |  |
| 2:28.153 | to answer you |  |  |  |
| 2:28.154 | for |  |  |  |
| 2:28.155 | what |  |  |  |
| 2:28.156 | am I |  |  |  |
| 2:28.157 | to enter |  |  |  |
| 2:28.158 | before |  |  |  |
| 2:28.159 | and within |  |  |  |
| 2:28.160 | what |  |  |  |
| 2:28.161 | that not |  |  |  |
| 2:28.162 | I was permitted |  |  |  |
| 2:28.163 | to ponder |  |  |  |
| 2:28.164 | all |  |  |  |
| 2:28.165 | the more so |  |  |  |
| 2:28.166 | to answer |  |  |  |
| 2:28.167 | questioners |  |  |  |
| 2:28.168 | but |  |  |  |
| 2:28.169 | because of |  |  |  |
| 2:28.170 | a time |  |  |  |
| 2:28.171 | to act |  |  |  |
| 2:28.172 | for the LORD |  |  | לה: "for the LORD", the Tetragrammaton; the words are Ps 119:126. |
| 2:28.173 | I |  |  |  |
| 2:28.174 | answer you |  |  |  |
| 2:28.175 | from what |  |  |  |
| 2:28.176 | I found |  |  |  |
| 2:28.177 | written |  |  |  |
| 2:28.178 | in the words of |  |  |  |
| 2:28.179 | our Sages of blessed memory |  |  | חז"ל: "our Sages of blessed memory". |
| 2:28.180 | and their merit |  |  |  |
| 2:28.181 | will stand |  |  |  |
| 2:28.182 | for me |  |  |  |
| 2:28.183 | that not |  |  |  |
| 2:28.184 | will befall |  |  |  |
| 2:28.185 | any |  |  |  |
| 2:28.186 | stumbling |  |  |  |
| 2:28.187 | through |  |  |  |
| 2:28.188 | my hand |  |  |  |
| 2:28.189 | behold |  |  |  |
| 2:28.190 | as for what |  |  |  |
| 2:28.191 | you wrote |  |  |  |
| 2:28.192 | that you do not |  |  |  |
| 2:28.193 | despise |  |  |  |
| 2:28.194 | the sages of |  |  |  |
| 2:28.195 | Israel |  |  |  |
| 2:28.196 | already |  |  |  |
| 2:28.197 | there are |  |  |  |
| 2:28.198 | against you |  |  |  |
| 2:28.199 | several |  |  |  |
| 2:28.200 | witnesses |  |  |  |
| 2:28.201 | not |  |  |  |
| 2:28.202 | you can |  |  |  |
| 2:28.203 | to deny |  |  |  |
| 2:28.204 | and what |  |  |  |
| 2:28.205 | brought |  |  |  |
| 2:28.206 | you |  |  |  |
| 2:28.207 | to |  |  |  |
| 2:28.208 | this |  |  |  |
| 2:28.209 | that you |  |  |  |
| 2:28.210 | attribute |  |  |  |
| 2:28.211 | to them |  |  |  |
| 2:28.212 | error |  |  |  |
| 2:28.213 | and heresy |  |  |  |
| 2:28.214 | but |  |  |  |
| 2:28.215 | because |  |  |  |
| 2:28.216 | that they are |  |  |  |
| 2:28.217 | despised |  |  |  |
| 2:28.218 | in your eyes |  |  |  |
| 2:28.219 | but |  |  |  |
| 2:28.220 | the person |  |  |  |
| 2:28.221 | should take heed |  |  |  |
| 2:28.222 | for himself |  |  |  |
| 2:28.223 | lest |  |  |  |
| 2:28.224 | he be burned |  |  |  |
| 2:28.225 | by their coals |  |  | בגחלתן: "in their coals", after Avot 2:10 (see apparatus). |
| 2:28.226 | for |  |  |  |
| 2:28.227 | great ones |  |  |  |
| 2:28.228 | righteous ones |  |  |  |
| 2:28.229 | in their death |  |  |  |
| 2:28.230 | more |  |  |  |
| 2:28.231 | than in their life |  |  |  |
| 2:28.232 | as |  |  |  |
| 2:28.233 | they said |  |  |  |
| 2:28.234 | in the Talmud |  |  |  |
| 2:28.235 | and as for what |  |  |  |
| 2:28.236 | you wrote |  |  |  |
| 2:28.237 | that if |  |  |  |
| 2:28.238 | to reveal |  |  |  |
| 2:28.239 | the place of |  |  |  |
| 2:28.240 | the contradiction |  |  |  |
| 2:28.241 | and also |  |  |  |
| 2:28.242 | the Talmud |  |  |  |
| 2:28.243 | brings |  |  |  |
| 2:28.244 | proof |  |  |  |
| 2:28.245 | to refute |  |  |  |
| 2:28.246 | the words of |  |  |  |
| 2:28.247 | greater |  |  |  |
| 2:28.248 | than them |  |  |  |
| 2:28.249 | there is not |  |  |  |
| 2:28.250 | the matter |  |  |  |
| 2:28.251 | comparable |  |  |  |
| 2:28.252 | to a proof |  |  |  |
| 2:28.253 | for where |  |  |  |
| 2:28.254 | do we find |  |  |  |
| 2:28.255 | that they bring |  |  |  |
| 2:28.256 | proof |  |  |  |
| 2:28.257 | to refute |  |  |  |
| 2:28.258 | a halakhah |  |  |  |
| 2:28.259 | given to Moses |  |  |  |
| 2:28.260 | at Sinai |  |  |  |
| 2:28.261 | and the masters of |  |  |  |
| 2:28.262 | Kabbalah |  |  | הקבלה: "the Kabbalah", the Kabbalists' lore; the author claims the same word for the rabbinic chain. |
| 2:28.263 | say |  |  |  |
| 2:28.264 | thus |  |  |  |
| 2:28.265 | they received |  |  | קיבלו: "they received", from the root of Kabbalah. |
| 2:28.266 | from the mouth of |  |  |  |
| 2:28.267 | Moses |  |  |  |
| 2:28.268 | and who |  |  |  |
| 2:28.269 | is he |  |  |  |
| 2:28.270 | this one |  |  |  |
| 2:28.271 | and where |  |  |  |
| 2:28.272 | is he |  |  |  |
| 2:28.273 | who gave |  |  |  |
| 2:28.274 | you |  |  |  |
| 2:28.275 | permission |  |  |  |
| 2:28.276 | to inquire |  |  |  |
| 2:28.277 | and to investigate |  |  |  |
| 2:28.278 | after |  |  |  |
| 2:28.279 | their words |  |  |  |
| 2:28.280 | and as for what |  |  |  |
| 2:28.281 | you wrote |  |  |  |
| 2:28.282 | that they |  |  |  |
| 2:28.283 | call |  |  |  |
| 2:28.284 | the partzufim |  |  |  |
| 2:28.285 | divinities |  |  |  |
| 2:28.286 | God forbid |  |  | ח"ו: "God forbid". |
| 2:28.287 | that they |  |  |  |
| 2:28.288 | say |  |  |  |
| 2:28.289 | thus |  |  |  |
| 2:28.290 | and as for what |  |  |  |
| 2:28.291 | you wrote |  |  |  |
| 2:28.292 | go out |  |  |  |
| 2:28.293 | and learn |  |  |  |
| 2:28.294 | from the intentions of |  |  |  |
| 2:28.295 | the blessings |  |  |  |
| 2:28.296 | and the prayers |  |  |  |
| 2:28.297 | already |  |  |  |
| 2:28.298 | they explained |  |  |  |
| 2:28.299 | that the cause of |  |  |  |
| 2:28.300 | all |  |  |  |
| 2:28.301 | the causes |  |  |  |
| 2:28.302 | is designated |  |  | נתכנה: "is designated", "is called". |
| 2:28.303 | in a sefirah |  |  |  |
| 2:28.304 | a certain |  |  |  |
| 2:28.305 | by the name |  |  |  |
| 2:28.306 | such |  |  |  |
| 2:28.307 | and in a sefirah |  |  |  |
| 2:28.308 | a certain |  |  |  |
| 2:28.309 | by the name |  |  |  |
| 2:28.310 | that |  |  |  |
| 2:28.311 | and as for what |  |  |  |
| 2:28.312 | you wrote |  |  |  |
| 2:28.313 | that they explained |  |  |  |
| 2:28.314 | the Kabbalists |  |  |  |
| 2:28.315 | that several |  |  |  |
| 2:28.316 | worlds |  |  |  |
| 2:28.317 | there are |  |  |  |
| 2:28.318 | thousands of |  |  |  |
| 2:28.319 | thousands |  |  |  |
| 2:28.320 | etc. |  |  |  |
| 2:28.321 | even if |  |  |  |
| 2:28.322 | they were |  |  |  |
| 2:28.323 | millions of |  |  |  |
| 2:28.324 | millions |  |  |  |
| 2:28.325 | since |  |  |  |
| 2:28.326 | that they |  |  |  |
| 2:28.327 | say |  |  |  |
| 2:28.328 | that their Creator |  |  |  |
| 2:28.329 | is one |  |  |  |
| 2:28.330 | and He |  |  |  |
| 2:28.331 | gives life to |  |  |  |
| 2:28.332 | (object marker) |  |  |  |
| 2:28.333 | all of them |  |  |  |
| 2:28.334 | what |  |  |  |
| 2:28.335 | is there |  |  |  |
| 2:28.336 | loss |  |  |  |
| 2:28.337 | from this |  |  |  |
| 2:28.338 | and as for what |  |  |  |
| 2:28.339 | you wrote |  |  |  |
| 2:28.340 | that they |  |  |  |
| 2:28.341 | in the image of |  |  |  |
| 2:28.342 | a man |  |  |  |
| 2:28.343 | of |  |  |  |
| 2:28.344 | stature |  |  |  |
| 2:28.345 | already |  |  |  |
| 2:28.346 | they said |  |  |  |
| 2:28.347 | great |  |  |  |
| 2:28.348 | is their power |  |  |  |
| 2:28.349 | of |  |  |  |
| 2:28.350 | the prophets |  |  |  |
| 2:28.351 | who liken |  |  |  |
| 2:28.352 | the form |  |  |  |
| 2:28.353 | of might |  |  |  |
| 2:28.354 | of |  |  |  |
| 2:28.355 | above |  |  |  |
| 2:28.356 | to the shape of |  |  |  |
| 2:28.357 | a man |  |  |  |
| 2:28.358 | and as for what |  |  |  |
| 2:28.359 | you wrote |  |  |  |
| 2:28.360 | about 248 |  |  | ברמ"ח: the letters רמ"ח are 248 in letter numerals, the traditional number of the limbs. |
| 2:28.361 | limbs |  |  |  |
| 2:28.362 | already |  |  |  |
| 2:28.363 | they explained |  |  |  |
| 2:28.364 | what |  |  |  |
| 2:28.365 | they are |  |  |  |
| 2:28.366 | the limbs |  |  |  |
| 2:28.367 | for |  |  |  |
| 2:28.368 | they are |  |  |  |
| 2:28.369 | 216 |  |  | רי"ו: 216 in letter numerals. |
| 2:28.370 | letters |  |  |  |
| 2:28.371 | of |  |  |  |
| 2:28.372 | the Name |  |  |  |
| 2:28.373 | of |  |  |  |
| 2:28.374 | 72 |  |  | ע"ב: 72 in letter numerals; the Name of 72 is AB in the gematria names. |
| 2:28.375 | and 32 |  |  | ול"ב: 32 in letter numerals. |
| 2:28.376 | paths |  |  |  |
| 2:28.377 | which are |  |  |  |
| 2:28.378 | they allude to |  |  |  |
| 2:28.379 | the powers of |  |  |  |
| 2:28.380 | divine |  |  |  |
| 2:28.381 | and as for what |  |  |  |
| 2:28.382 | you wrote |  |  |  |
| 2:28.383 | that already |  |  |  |
| 2:28.384 | I conceded |  |  |  |
| 2:28.385 | that my opinion |  |  |  |
| 2:28.386 | that there is not |  |  |  |
| 2:28.387 | any |  |  |  |
| 2:28.388 | prayer |  |  |  |
| 2:28.389 | applicable |  |  |  |
| 2:28.390 | to the cause of |  |  |  |
| 2:28.391 | all |  |  |  |
| 2:28.392 | causes |  |  |  |
| 2:28.393 | falsehood |  |  |  |
| 2:28.394 | you |  |  |  |
| 2:28.395 | speak |  |  |  |
| 2:28.396 | but surely |  |  |  |
| 2:28.397 | from the beginning of |  |  |  |
| 2:28.398 | your question |  |  |  |
| 2:28.399 | to me |  |  |  |
| 2:28.400 | about |  |  |  |
| 2:28.401 | the words of |  |  |  |
| 2:28.402 | the Mesharef |  |  | המשרף: "the one who burns"; the reply does not name him. |
| 2:28.403 | I answered you |  |  |  |
| 2:28.404 | see |  |  |  |
| 2:28.405 | in Kisse Eliyahu |  |  | בכס"א: Kisse Eliyahu, "The Throne of Elijah", a book title; the meaning is given at its first occurrence in the unit. |
| 2:28.406 | but |  |  |  |
| 2:28.407 | our Kabbalah |  |  |  |
| 2:28.408 | that prayer |  |  |  |
| 2:28.409 | to Ein Sof |  |  | לא"ס: "to Ein Sof"; Ein Sof is literally "without end". |
| 2:28.410 | and to bestow |  |  |  |
| 2:28.411 | upon |  |  |  |
| 2:28.412 | ten |  |  |  |
| 2:28.413 | sefirot |  |  |  |
| 2:28.414 | as |  |  |  |
| 2:28.415 | he wrote |  |  |  |
| 2:28.416 | Lechem |  |  | לחם: Lechem Shelomo, a book title, literally "the bread of Solomon". |
| 2:28.417 | Shelomo |  |  |  |
| 2:28.418 | and as for what |  |  |  |
| 2:28.419 | I said |  |  |  |
| 2:28.420 | to you |  |  |  |
| 2:28.421 | to consult |  |  |  |
| 2:28.422 | in the Throne |  |  |  |
| 2:28.423 | of Elijah |  |  |  |
| 2:28.424 | so that |  |  |  |
| 2:28.425 | you would know |  |  |  |
| 2:28.426 | that the words of |  |  |  |
| 2:28.427 | the Mesharef |  |  |  |
| 2:28.428 | are the words of |  |  |  |
| 2:28.429 | heresy |  |  |  |
| 2:28.430 | in truth |  |  |  |
| 2:28.431 | and fitting |  |  |  |
| 2:28.432 | to root it out |  |  |  |
| 2:28.433 | and God forbid |  |  |  |
| 2:28.434 | that there be found |  |  |  |
| 2:28.435 | in any |  |  |  |
| 2:28.436 | book |  |  |  |
| 2:28.437 | like his words |  |  |  |
| 2:28.438 | but |  |  |  |
| 2:28.439 | the words of |  |  |  |
| 2:28.440 | Kisse Eliyahu |  |  |  |
| 2:28.441 | are not |  |  |  |
| 2:28.442 | like his words |  |  |  |
| 2:28.443 | God forbid |  |  |  |
| 2:28.444 | for |  |  |  |
| 2:28.445 | the words of |  |  |  |
| 2:28.446 | the Throne |  |  |  |
| 2:28.447 | of Elijah |  |  |  |
| 2:28.448 | is nothing |  |  |  |
| 2:28.449 | Ze'ir |  |  | הזעיר: "Ze'ir", from Ze'ir Anpin, literally "Small Face" (Aramaic). |
| 2:28.450 | Anpin |  |  |  |
| 2:28.451 | but |  |  |  |
| 2:28.452 | only |  |  |  |
| 2:28.453 | a vessel |  |  | כלי: "vessel", the Kabbalistic term. |
| 2:28.454 | for the action of |  |  |  |
| 2:28.455 | the Creator |  |  |  |
| 2:28.456 | and as for what |  |  |  |
| 2:28.457 | he wrote |  |  |  |
| 2:28.458 | on folio |  |  |  |
| 2:28.459 | 25 |  |  | כ"ה: 25, folio 25 (דף כ"ה). |
| 2:28.460 | that when we |  |  |  |
| 2:28.461 | call Ze'ir Anpin |  |  |  |
| 2:28.462 | we |  |  |  |
| 2:28.463 | call |  |  |  |
| 2:28.464 | God |  |  | אל: the bare name El, here of the partzufim; rendered "God" (the alternative is "El"). |
| 2:28.465 | all |  |  |  |
| 2:28.466 | the partzufim |  |  |  |
| 2:28.467 | with |  |  |  |
| 2:28.468 | a soul |  |  |  |
| 2:28.469 | hidden |  |  |  |
| 2:28.470 | in them |  |  |  |
| 2:28.471 | which is |  |  |  |
| 2:28.472 | Ein |  |  |  |
| 2:28.473 | Sof |  |  |  |
| 2:28.474 | blessed be He |  |  | ב"ה: "blessed be He". |
| 2:28.475 | this does not |  |  |  |
| 2:28.476 | mean |  |  |  |
| 2:28.477 | to say |  |  |  |
| 2:28.478 | that not |  |  |  |
| 2:28.479 | prayer |  |  |  |
| 2:28.480 | to Ein Sof |  |  |  |
| 2:28.481 | Heaven forbid |  |  |  |
| 2:28.482 | but |  |  |  |
| 2:28.483 | it means |  |  |  |
| 2:28.484 | to say |  |  |  |
| 2:28.485 | the ascent of |  |  |  |
| 2:28.486 | prayer |  |  |  |
| 2:28.487 | to |  |  |  |
| 2:28.488 | the Ein Sof |  |  |  |
| 2:28.489 | blessed be He |  |  |  |
| 2:28.490 | by way of |  |  |  |
| 2:28.491 | this |  |  |  |
| 2:28.492 | attribute |  |  |  |
| 2:28.493 | and the proof |  |  |  |
| 2:28.494 | for |  |  |  |
| 2:28.495 | this |  |  |  |
| 2:28.496 | is what |  |  |  |
| 2:28.497 | he wrote |  |  |  |
| 2:28.498 | on folio |  |  |  |
| 2:28.499 | 17 |  |  |  |
| 2:28.500 | b |  |  | ע"ב: "b", the second side of the folio (17b). |
| 2:28.501 | and these are his words: |  |  | וז"ל: "and these are his words". |
| 2:28.502 | as for what |  |  |  |
| 2:28.503 | is written |  |  |  |
| 2:28.504 | about the intentions of |  |  |  |
| 2:28.505 | prayer |  |  |  |
| 2:28.506 | and the blessings |  |  |  |
| 2:28.507 | to direct |  |  |  |
| 2:28.508 | in each |  |  |  |
| 2:28.509 | blessing |  |  |  |
| 2:28.510 | an intention |  |  |  |
| 2:28.511 | particular |  |  |  |
| 2:28.512 | to a sefirah |  |  |  |
| 2:28.513 | particular |  |  |  |
| 2:28.514 | not |  |  |  |
| 2:28.515 | God forbid |  |  |  |
| 2:28.516 | to the essence of |  |  |  |
| 2:28.517 | the sefirah |  |  |  |
| 2:28.518 | for |  |  |  |
| 2:28.519 | it |  |  |  |
| 2:28.520 | a cutting off |  |  |  |
| 2:28.521 | God forbid |  |  |  |
| 2:28.522 | and Heaven forbid |  |  |  |
| 2:28.523 | to direct |  |  |  |
| 2:28.524 | to any |  |  |  |
| 2:28.525 | power |  |  |  |
| 2:28.526 | particular |  |  |  |
| 2:28.527 | and sefirah |  |  |  |
| 2:28.528 | particular |  |  |  |
| 2:28.529 | but |  |  |  |
| 2:28.530 | all |  |  |  |
| 2:28.531 | to the essence of |  |  |  |
| 2:28.532 | the one |  |  |  |
| 2:28.533 | Ein Sof |  |  |  |
| 2:28.534 | blessed be He |  |  |  |
| 2:28.535 | comprising |  |  |  |
| 2:28.536 | all |  |  |  |
| 2:28.537 | the powers |  |  |  |
| 2:28.538 | all of them |  |  |  |
| 2:28.539 | and furthermore |  |  |  |
| 2:28.540 | he wrote |  |  |  |
| 2:28.541 | on folio |  |  |  |
| 2:28.542 | 25 |  |  |  |
| 2:28.543 | a |  |  |  |
| 2:28.544 | and all |  |  |  |
| 2:28.545 | our prayers |  |  |  |
| 2:28.546 | are |  |  |  |
| 2:28.547 | to Ein |  |  |  |
| 2:28.548 | Sof |  |  |  |
| 2:28.549 | may He be blessed |  |  |  |
| 2:28.550 | but |  |  |  |
| 2:28.551 | that not |  |  |  |
| 2:28.552 | we |  |  |  |
| 2:28.553 | can |  |  |  |
| 2:28.554 | call Him |  |  |  |
| 2:28.555 | by any |  |  |  |
| 2:28.556 | appellation |  |  |  |
| 2:28.557 | and name |  |  |  |
| 2:28.558 | therefore |  |  |  |
| 2:28.559 | we |  |  |  |
| 2:28.560 | pray |  |  |  |
| 2:28.561 | by means of |  |  |  |
| 2:28.562 | the hand of |  |  |  |
| 2:28.563 | the sefirot |  |  |  |
| 2:28.564 | for |  |  |  |
| 2:28.565 | the name |  |  |  |
| 2:28.566 | pertaining |  |  |  |
| 2:28.567 | all |  |  |  |
| 2:28.568 | all |  |  | כל: the doubled "כל כל" is as printed. |
| 2:28.569 | the appellations |  |  |  |
| 2:28.570 | and names |  |  |  |
| 2:28.571 | for |  |  |  |
| 2:28.572 | all |  |  |  |
| 2:28.573 | the appellations |  |  |  |
| 2:28.574 | and the names |  |  |  |
| 2:28.575 | are |  |  |  |
| 2:28.576 | names |  |  |  |
| 2:28.577 | of His essence |  |  |  |
| 2:28.578 | that extends |  |  |  |
| 2:28.579 | through the sefirot |  |  |  |
| 2:28.580 | also |  |  |  |
| 2:28.581 | he wrote |  |  |  |
| 2:28.582 | on folio |  |  |  |
| 2:28.583 | 26 |  |  |  |
| 2:28.584 | if |  |  |  |
| 2:28.585 | not |  |  |  |
| 2:28.586 | they do |  |  |  |
| 2:28.587 | so |  |  |  |
| 2:28.588 | and not |  |  |  |
| 2:28.589 | unify |  |  |  |
| 2:28.590 | their prayer |  |  |  |
| 2:28.591 | but |  |  |  |
| 2:28.592 | to Ze'ir Anpin |  |  |  |
| 2:28.593 | not |  |  |  |
| 2:28.594 | they will be |  |  |  |
| 2:28.595 | answered |  |  |  |
| 2:28.596 | and all |  |  |  |
| 2:28.597 | the Kabbalists |  |  |  |
| 2:28.598 | all hold |  |  | כיילי: Aramaic, "all of them hold" (see grammar). |
| 2:28.599 | the principle |  |  |  |
| 2:28.600 | this |  |  |  |
| 2:28.601 | and these are their words: |  |  |  |
| 2:28.602 | as for whatever |  |  |  |
| 2:28.603 | you find |  |  |  |
| 2:28.604 | in some |  |  |  |
| 2:28.605 | book |  |  |  |
| 2:28.606 | of theirs |  |  |  |
| 2:28.607 | prayer |  |  |  |
| 2:28.608 | to any |  |  |  |
| 2:28.609 | sefirah |  |  |  |
| 2:28.610 | Heaven forbid |  |  |  |
| 2:28.611 | that it be |  |  |  |
| 2:28.612 | so |  |  |  |
| 2:28.613 | but |  |  |  |
| 2:28.614 | the intention |  |  |  |
| 2:28.615 | that he direct |  |  |  |
| 2:28.616 | to raise them |  |  |  |
| 2:28.617 | to |  |  |  |
| 2:28.618 | Ein Sof |  |  |  |
| 2:28.619 | by way of |  |  |  |
| 2:28.620 | that |  |  |  |
| 2:28.621 | sefirah |  |  |  |

Phrases:
- 2:28.1 2:28.2 2:28.3 2:28.4 2:28.5 2:28.6 2:28.7 | Reply. To man the arrangements of the heart, and from the LORD the answer of the tongue. | 'Reply to-man the-arrangements-of the-heart and-from-the-LORD the-answer-of the-tongue'
- 2:28.8 2:28.9 2:28.10 2:28.11 | And it was when I saw your letter, | 'and-it-was when-I-saw (object-marker) your-letter'
- 2:28.12 2:28.13 2:28.14 | I sat astonished and dismayed. | 'I-sat astonished and-dismayed'
- 2:28.15 2:28.16 2:28.17 2:28.18 2:28.19 | How your heart came to inquire and investigate | 'how has-filled your-heart to-inquire and-to-investigate'
- 2:28.20 2:28.21 2:28.22 | after the band of the Geonim, | 'after band of-the-Geonim'
- 2:28.23 2:28.24 2:28.25 2:28.26 2:28.27 2:28.28 2:28.29 2:28.30 | who are faithful in the Torah, distinguished in lineage, like the seed of the mighty, | 'who they-are in-the-Torah faithful distinguished-of lineage like-seed-of the-mighty-ones'
- 2:28.31 2:28.32 2:28.33 2:28.34 2:28.35 | and after several decisors and also expounders, | 'and-after several decisors and-also expounders'
- 2:28.36 2:28.37 2:28.38 | the early and the later, | 'the-early with the-later'
- 2:28.39 2:28.40 2:28.41 2:28.42 2:28.43 | the sages of the West, Ashkenaz, Sefarad and the Yemenites. | 'sages-of the-West and-Ashkenaz and-Sefarad and-the-Yemenites'
- 2:28.44 2:28.45 2:28.46 2:28.47 2:28.48 2:28.49 | and the rest of the congregations, the holy seed, the faithful children, | 'and-the-rest-of congregations seed holy children-of faithful'
- 2:28.50 2:28.51 2:28.52 2:28.53 2:28.54 2:28.55 | this is their treasure, more than gold and abundant pearls. | 'this is their-treasure more-than-gold and-abundant pearls'
- 2:28.56 2:28.57 2:28.58 2:28.59 2:28.60 2:28.61 | Can a stumbling come from the hand of righteous and faithful ones, | 'can-it-be-that comes a-stumbling from-the-hand-of righteous-ones faithful-ones'
- 2:28.62 2:28.63 2:28.64 2:28.65 2:28.66 2:28.67 2:28.68 2:28.69 2:28.70 | who know knowledge, the wise and discerning, who fear God and turn aside from evil with their might. | 'who-know knowledge the-wise and-discerning who-fear God and-who-turn-aside from-evil with-their-might'
- 2:28.71 2:28.72 2:28.73 2:28.74 2:28.75 2:28.76 2:28.77 2:28.78 2:28.79 | By their wisdom they felled many foes and multitudes who stood against them in every age. | 'by-their-wisdom they-felled many foes and-multitudes who-stood against-them in-every age'
- 2:28.80 2:28.81 2:28.82 2:28.83 2:28.84 | Can it be that their Kabbalah is compounded in vanity and error, | 'can-it-be-that their-Kabbalah compounded in-vanity and-error'
- 2:28.85 2:28.86 2:28.87 2:28.88 2:28.89 2:28.90 2:28.91 | which, because of their great wisdom, were approved by the mouths of princes and nobles | 'which because-of-the-great their-wisdom were-approved by-the-mouths-of princes and-nobles'
- 2:28.92 2:28.93 2:28.94 2:28.95 | and stood before kings and rulers, | 'they-stood before kings and-rulers'
- 2:28.96 2:28.97 2:28.98 2:28.99 2:28.100 2:28.101 2:28.102 | and can fathers and sons walk on a path that is not good? | 'can-it-be-that they-walk on-the-path not good fathers and-sons'
- 2:28.103 2:28.104 2:28.105 2:28.106 | How many hundreds of years is this? | 'this how-many hundreds of-years'
- 2:28.107 2:28.108 2:28.109 2:28.110 2:28.111 2:28.112 | And how many of them are distinguished in every wisdom, and scholars, | 'and-how-many they are-distinguished in-every wisdom and-scholars'
- 2:28.113 2:28.114 2:28.115 2:28.116 | Kabbalists with philosophy and logic, | 'Kabbalists with philosophy and-logic'
- 2:28.117 2:28.118 2:28.119 | medicine, tishburet and computations. | 'medicine and-tishburet and-computations'
- 2:28.120 2:28.121 2:28.122 2:28.123 2:28.124 | meditating day and night, quiet and at ease, | 'meditating by-day and-by-night quiet and-at-ease'
- 2:28.125 2:28.126 2:28.127 2:28.128 2:28.129 2:28.130 | eating the fruit of their deeds, the good, fat and lush. | 'eating the-fruit-of their-deeds the-good fat and-lush'
- 2:28.131 2:28.132 2:28.133 2:28.134 | Heaven forbid! This is not to be believed | 'Heaven-forbid this is-not to-be-believed'
- 2:28.135 2:28.136 2:28.137 2:28.138 | and does not come to mind | 'and-not come-up upon the-heart'
- 2:28.139 2:28.140 2:28.141 | but only for the simple, | 'but only for-the-simple'
- 2:28.142 2:28.143 2:28.144 2:28.145 2:28.146 2:28.147 2:28.148 | who believe every matter, who lean on their understanding. | 'to-every matter who-believe who upon their-understanding lean'
- 2:28.149 2:28.150 2:28.151 2:28.152 2:28.153 | And properly, it is not incumbent on me to answer you. | 'and-properly what-is-fitting there-is-not upon-me to-answer-you'
- 2:28.154 2:28.155 2:28.156 | For what am I, | 'for what am-I'
- 2:28.157 2:28.158 2:28.159 | to enter before and within, | 'to-enter before and-within'
- 2:28.160 2:28.161 2:28.162 2:28.163 | into what I was not permitted to ponder, | 'what that-not I-was-permitted to-ponder'
- 2:28.164 2:28.165 2:28.166 2:28.167 | all the more so to answer questioners. | 'all the-more-so to-answer questioners'
- 2:28.168 2:28.169 2:28.170 2:28.171 2:28.172 | But because it is a time to act for the LORD, | 'but because-of a-time to-act for-the-LORD'
- 2:28.173 2:28.174 2:28.175 2:28.176 2:28.177 2:28.178 2:28.179 | I answer you from what I found written in the words of our Sages of blessed memory. | 'I answer-you from-what I-found written in-the-words-of our-Sages-of-blessed-memory'
- 2:28.180 2:28.181 2:28.182 2:28.183 2:28.184 2:28.185 2:28.186 2:28.187 2:28.188 | May their merit stand for me, that no stumbling befall through me. | 'and-their-merit will-stand for-me that-not will-befall any stumbling through my-hand'
- 2:28.189 2:28.190 2:28.191 2:28.192 2:28.193 2:28.194 2:28.195 | Now, as for what you wrote, that you do not despise the sages of Israel, | 'behold as-for-what you-wrote that-you-do-not despise the-sages-of Israel'
- 2:28.196 2:28.197 2:28.198 2:28.199 2:28.200 | there are already several witnesses against you; | 'already there-are against-you several witnesses'
- 2:28.201 2:28.202 2:28.203 | you cannot deny it. | 'not you-can to-deny'
- 2:28.204 2:28.205 2:28.206 2:28.207 2:28.208 | And what brought you to this | 'and-what brought you to this'
- 2:28.209 2:28.210 2:28.211 2:28.212 2:28.213 | that you attribute to them error and heresy. | 'that-you attribute to-them error and-heresy'
- 2:28.214 2:28.215 2:28.216 2:28.217 2:28.218 | Rather, because they are despised in your eyes. | 'but because that-they-are despised in-your-eyes'
- 2:28.219 2:28.220 2:28.221 2:28.222 2:28.223 2:28.224 2:28.225 | But a person should take heed for himself, lest he be burned by their coals, | 'but the-person should-take-heed for-himself lest he-be-burned by-their-coals'
- 2:28.226 2:28.227 2:28.228 2:28.229 2:28.230 2:28.231 | for the righteous are greater in their death than in their life, | 'for great-ones righteous-ones in-their-death more than-in-their-life'
- 2:28.232 2:28.233 2:28.234 | as they said in the Talmud. | 'as they-said in-the-Talmud'
- 2:28.235 2:28.236 2:28.237 2:28.238 2:28.239 2:28.240 | And as for what you wrote, that if one reveals the place of contradiction, | 'and-as-for-what you-wrote that-if to-reveal the-place-of the-contradiction'
- 2:28.241 2:28.242 2:28.243 2:28.244 2:28.245 2:28.246 2:28.247 2:28.248 | and the Talmud also brings proof to refute the words of those greater than them, | 'and-also the-Talmud brings proof to-refute the-words-of greater than-them'
- 2:28.249 2:28.250 2:28.251 2:28.252 | the matter is not comparable to a proof, | 'there-is-not the-matter comparable to-a-proof'
- 2:28.253 2:28.254 2:28.255 2:28.256 2:28.257 2:28.258 2:28.259 2:28.260 | for where do we find that they bring proof to refute a halakhah given to Moses at Sinai? | 'for-where do-we-find that-they-bring proof to-refute a-halakhah given-to-Moses at-Sinai'
- 2:28.261 2:28.262 2:28.263 2:28.264 2:28.265 2:28.266 2:28.267 | And the masters of Kabbalah say that they received thus from the mouth of Moses, | 'and-the-masters-of Kabbalah say thus they-received from-the-mouth-of Moses'
- 2:28.268 2:28.269 2:28.270 2:28.271 2:28.272 2:28.273 2:28.274 2:28.275 2:28.276 2:28.277 2:28.278 2:28.279 | and who is this one, and where is he, who gave you permission to inquire and investigate after their words? | 'and-who is-he this-one and-where is-he who-gave you permission to-inquire and-to-investigate after their-words'
- 2:28.280 2:28.281 2:28.282 2:28.283 2:28.284 2:28.285 | And as for what you wrote, that they call the partzufim divinities, | 'and-as-for-what you-wrote that-they call the-partzufim divinities'
- 2:28.286 2:28.287 2:28.288 2:28.289 | God forbid, that they say thus. | 'God-forbid that-they say thus'
- 2:28.290 2:28.291 2:28.292 2:28.293 2:28.294 2:28.295 2:28.296 | And as for what you wrote, go and learn from the intentions of the blessings and the prayers. | 'and-as-for-what you-wrote go-out and-learn from-the-intentions-of the-blessings and-the-prayers'
- 2:28.297 2:28.298 2:28.299 2:28.300 2:28.301 2:28.302 | They already explained that the cause of all causes is designated | 'already they-explained that-the-cause-of all the-causes is-designated'
- 2:28.303 2:28.304 2:28.305 2:28.306 | in a certain sefirah by such a name. | 'in-a-sefirah a-certain by-the-name such'
- 2:28.307 2:28.308 2:28.309 2:28.310 | And in a certain sefirah by that name. | 'and-in-a-sefirah a-certain by-the-name that'
- 2:28.311 2:28.312 2:28.313 2:28.314 2:28.315 2:28.316 2:28.317 2:28.318 2:28.319 2:28.320 | And as for what you wrote, that the Kabbalists explained that there are several worlds, thousands of thousands, etc., | 'and-as-for-what you-wrote that-they-explained the-Kabbalists that-several worlds there-are thousands-of thousands etc.'
- 2:28.321 2:28.322 2:28.323 2:28.324 | even if they were millions of millions, | 'even-if they-were millions-of millions'
- 2:28.325 2:28.326 2:28.327 2:28.328 2:28.329 2:28.330 2:28.331 2:28.332 2:28.333 2:28.334 2:28.335 2:28.336 2:28.337 | since they say that their Creator is one, and He gives life to all of them: what loss is there from this? | 'since that-they say that-their-Creator is-one and-He gives-life-to (object-marker) all-of-them what is-there loss from-this'
- 2:28.338 2:28.339 2:28.340 2:28.341 2:28.342 2:28.343 2:28.344 | And as for what you wrote, that they are in the image of a man of stature, | 'and-as-for-what you-wrote that-they in-the-image-of a-man of stature'
- 2:28.345 2:28.346 2:28.347 2:28.348 2:28.349 2:28.350 2:28.351 2:28.352 2:28.353 2:28.354 2:28.355 2:28.356 2:28.357 | already they said, "Great is the power of the prophets, who liken the form of might above to the shape of a man." | 'already they-said great is-their-power of the-prophets who-liken the-form of-might of above to-the-shape-of a-man'
- 2:28.358 2:28.359 2:28.360 2:28.361 | And as for what you wrote about 248 limbs, | 'and-as-for-what you-wrote about-248 limbs'
- 2:28.362 2:28.363 2:28.364 2:28.365 2:28.366 | they already explained what the limbs are. | 'already they-explained what they-are the-limbs'
- 2:28.367 2:28.368 2:28.369 2:28.370 2:28.371 2:28.372 | For they are 216 letters of the Name | 'for they-are 216 letters of the-Name'
- 2:28.373 2:28.374 2:28.375 2:28.376 | of 72 and 32 paths, | 'of 72 and-32 paths'
- 2:28.377 2:28.378 2:28.379 2:28.380 | which allude to divine powers. | 'which-are they-allude-to the-powers-of divine'
- 2:28.381 2:28.382 2:28.383 2:28.384 2:28.385 | And as for what you wrote, that I already conceded that my opinion | 'and-as-for-what you-wrote that-already I-conceded that-my-opinion'
- 2:28.386 2:28.387 2:28.388 2:28.389 2:28.390 2:28.391 2:28.392 | is that there is no prayer applicable to the Cause of all causes, | 'that-there-is-not any prayer applicable to-the-cause-of all causes'
- 2:28.393 2:28.394 2:28.395 | you speak falsehood. | 'falsehood you speak'
- 2:28.396 2:28.397 2:28.398 2:28.399 2:28.400 2:28.401 2:28.402 | But surely, from the beginning of your question to me about the words of the Mesharef, | 'but-surely from-the-beginning-of your-question to-me about the-words-of the-Mesharef'
- 2:28.403 2:28.404 2:28.405 | I answered you, see in Kisse Eliyahu. | 'I-answered-you see in-Kisse-Eliyahu'
- 2:28.406 2:28.407 2:28.408 2:28.409 2:28.410 2:28.411 2:28.412 2:28.413 | But our Kabbalah holds that prayer is to Ein Sof and bestows upon the ten sefirot, | 'but our-Kabbalah that-prayer to-Ein-Sof and-to-bestow upon ten sefirot'
- 2:28.414 2:28.415 2:28.416 2:28.417 | as Lechem Shelomo wrote. | 'as he-wrote Lechem Shelomo'
- 2:28.418 2:28.419 2:28.420 2:28.421 2:28.422 2:28.423 | And as for what I told you, to consult Kisse Eliyahu, | 'and-as-for-what I-said to-you to-consult in-the-Throne of-Elijah'
- 2:28.424 2:28.425 2:28.426 2:28.427 2:28.428 2:28.429 2:28.430 2:28.431 2:28.432 | so that you would know that the words of the Mesharef are words of heresy, in truth, and fitting to root it out. | 'so-that you-would-know that-the-words-of the-Mesharef are-the-words-of heresy in-truth and-fitting to-root-it-out'
- 2:28.433 2:28.434 2:28.435 2:28.436 2:28.437 | And God forbid that there be found in any book words like his, | 'and-God-forbid that-there-be-found in-any book like-his-words'
- 2:28.438 2:28.439 2:28.440 2:28.441 2:28.442 2:28.443 | but the words of Kisse Eliyahu are not like his, God forbid. | 'but the-words-of Kisse-Eliyahu are-not like-his-words God-forbid'
- 2:28.444 2:28.445 2:28.446 2:28.447 | For the words of Kisse Eliyahu | 'for the-words-of the-Throne of-Elijah'
- 2:28.448 2:28.449 2:28.450 2:28.451 2:28.452 2:28.453 2:28.454 2:28.455 | hold that Ze'ir Anpin is nothing but a vessel for the action of the Creator. | 'is-nothing Ze'ir Anpin but only a-vessel for-the-action-of the-Creator'
- 2:28.456 2:28.457 2:28.458 2:28.459 | And as for what he wrote on folio 25, | 'and-as-for-what he-wrote on-folio 25'
- 2:28.460 2:28.461 2:28.462 2:28.463 2:28.464 2:28.465 2:28.466 2:28.467 2:28.468 2:28.469 2:28.470 | that when we call Ze'ir Anpin God, we call all the partzufim God, with a soul hidden in them, | 'that-when-we call-Ze'ir-Anpin we call God all the-partzufim with a-soul hidden in-them'
- 2:28.471 2:28.472 2:28.473 2:28.474 | which is Ein Sof, blessed be He. | 'which-is Ein Sof blessed-be-He'
- 2:28.475 2:28.476 2:28.477 | This does not mean | 'this-does-not mean to-say'
- 2:28.478 2:28.479 2:28.480 2:28.481 | that prayer is not directed to Ein Sof, Heaven forbid, | 'that-not prayer to-Ein-Sof Heaven-forbid'
- 2:28.482 2:28.483 2:28.484 | but it means that | 'but it-means to-say'
- 2:28.485 2:28.486 2:28.487 2:28.488 2:28.489 2:28.490 2:28.491 2:28.492 | the ascent of prayer to Ein Sof, blessed be He, is by way of this attribute. | 'the-ascent-of prayer to the-Ein-Sof blessed-be-He by-way-of this attribute'
- 2:28.493 2:28.494 2:28.495 2:28.496 2:28.497 | And the proof for this is what he wrote | 'and-the-proof for this is-what he-wrote'
- 2:28.498 2:28.499 2:28.500 2:28.501 | on folio 17b, and these are his words: | 'on-folio 17 b and-these-are-his-words:'
- 2:28.502 2:28.503 2:28.504 2:28.505 2:28.506 2:28.507 2:28.508 2:28.509 2:28.510 2:28.511 | "As for what is written about the intentions of prayer and the blessings, to direct in each blessing a particular intention, | 'as-for-what is-written about-the-intentions-of prayer and-the-blessings to-direct in-each blessing an-intention particular'
- 2:28.512 2:28.513 2:28.514 2:28.515 2:28.516 2:28.517 | to a particular sefirah, God forbid, not to the essence of the sefirah. | 'to-a-sefirah particular not God-forbid to-the-essence-of the-sefirah'
- 2:28.518 2:28.519 2:28.520 2:28.521 | For it is a cutting off, God forbid, | 'for it a-cutting-off God-forbid'
- 2:28.522 2:28.523 2:28.524 2:28.525 2:28.526 2:28.527 2:28.528 | and Heaven forbid to direct it to any particular power or particular sefirah. | 'and-Heaven-forbid to-direct to-any power particular and-sefirah particular'
- 2:28.529 2:28.530 2:28.531 2:28.532 2:28.533 2:28.534 2:28.535 2:28.536 2:28.537 2:28.538 | But all is to the essence of the one Ein Sof, blessed be He, comprising all the powers, all of them." | 'but all to-the-essence-of the-one Ein-Sof blessed-be-He comprising all the-powers all-of-them'
- 2:28.539 2:28.540 2:28.541 2:28.542 2:28.543 | And furthermore he wrote on folio 25a: | 'and-furthermore he-wrote on-folio 25 a'
- 2:28.544 2:28.545 2:28.546 2:28.547 2:28.548 2:28.549 | "And all our prayers are to Ein Sof, may He be blessed. | 'and-all our-prayers are to-Ein Sof may-He-be-blessed'
- 2:28.550 2:28.551 2:28.552 2:28.553 2:28.554 2:28.555 2:28.556 2:28.557 | But that we cannot call Him by any appellation and name. | 'but that-not we can call-Him by-any appellation and-name'
- 2:28.558 2:28.559 2:28.560 2:28.561 2:28.562 2:28.563 | Therefore we pray through the sefirot. | 'therefore we pray by-means-of the-hand-of the-sefirot'
- 2:28.564 2:28.565 2:28.566 2:28.567 2:28.568 2:28.569 2:28.570 | For the name pertaining to all, all the appellations and names. | 'for the-name pertaining all all the-appellations and-names'
- 2:28.571 2:28.572 2:28.573 2:28.574 2:28.575 2:28.576 2:28.577 2:28.578 2:28.579 | For all the appellations and names are names of His essence, extending through the sefirot." | 'for all the-appellations and-the-names are names of-His-essence that-extends through-the-sefirot'
- 2:28.580 2:28.581 2:28.582 2:28.583 | Also he wrote on folio 26: | 'also he-wrote on-folio 26'
- 2:28.584 2:28.585 2:28.586 2:28.587 2:28.588 2:28.589 2:28.590 2:28.591 2:28.592 2:28.593 2:28.594 2:28.595 | "if they do not do so, and do not unify their prayer but to Ze'ir Anpin, they will not be answered." | 'if not they-do so and-not unify their-prayer but to-Ze'ir-Anpin not they-will-be answered'
- 2:28.596 2:28.597 2:28.598 2:28.599 2:28.600 | And all the Kabbalists hold this principle, | 'and-all the-Kabbalists all-hold the-principle this'
- 2:28.601 2:28.602 2:28.603 2:28.604 2:28.605 2:28.606 2:28.607 2:28.608 2:28.609 2:28.610 2:28.611 2:28.612 | and these are their words: "As for whatever you find in some book of theirs, that prayer is addressed to any sefirah, Heaven forbid that it be so. | 'and-these-are-their-words: as-for-whatever you-find in-some book of-theirs prayer to-any sefirah Heaven-forbid that-it-be so'
- 2:28.613 2:28.614 2:28.615 2:28.616 2:28.617 2:28.618 2:28.619 2:28.620 2:28.621 | But the intention is that one direct them to raise them to Ein Sof by way of that sefirah." | 'but the-intention that-he-direct to-raise-them to Ein-Sof by-way-of that sefirah'

Grammar notes:
- The paragraph is the respondent's reply in chapter 2. "You" is the questioner; the reply takes up each point of his letter in turn, introduced by ומה שכתבת ("as for what you wrote") and answered after it.
- The rhetorical questions opened by וכי ("Can it be that ...", 2:28.56, 2:28.80, 2:28.96) and the sentences ending in a question mark (2:28.102, 2:28.106, 2:28.260, 2:28.279, 2:28.337) keep their question form; the text gives no answer to them within the paragraph.
- Quotations of the Kabbalists are in double quotation marks. They open after the formula וז"ל (2:28.501, 2:28.601) or after כ' with a folio (2:28.539, 2:28.580), and close at the end of each passage, since the text gives no closing formula. The quoted words are Hebrew, except כיילי כללא דא (2:28.598-600), which is Aramaic and is rendered by its sense.
- The verb ייחדו (2:28.589) takes the prayer as its object and is rendered "unify", the Kabbalistic act; לז"א (2:28.592) is the entity the prayer is directed to.
- Two Biblical allusions carry the opening and the close of the argument: Prov 16:1 at the start (2:28.1-7), and the word לה' from Ps 119:126 (2:28.172), which the author reads as "for the LORD".

Translator's note: Prose paragraph: the respondent's reply in chapter 2. Each word is glossed. Letter numerals are given as Arabic numerals (רמ"ח 248, רי"ו 216, ע"ב 72, ל"ב 32, כ"ה 25, י"ז ע"ב 17b). Kisse Eliyahu and Lechem Shelomo are book titles; Ein Sof and Ze'ir Anpin are kept as names.

Commentary rulings:
- apparatus: Base text: Wikisource, as printed; the 1931 scan is not yet collated (Q2), so no reading is added or corrected here.
- apparatus: Prov 16:1 (allusion at the opening, 2:28.1-7): the verse reads "The arrangements of the heart belong to man, and the answer of the tongue is from the LORD." The author's Hebrew is the verse's words with תשובה prefixed.
- apparatus: Job 1:1 (allusion, 2:28.66-70): "וִירֵא אֱלֹהִים וְסָר מֵרָע", rendered "fear God and turn from evil", with the author's plural form.
- apparatus: Isa 6:13 (allusion, 2:28.46-47): "זֶרַע קֹדֶשׁ", "the holy seed".
- apparatus: Isa 65:2 (allusion, 2:28.97-100): "בְּדֶרֶךְ לֹא־טוֹב", "in a way that is not good".
- apparatus: Ps 92:15 (allusion, 2:28.129-130): "דְּשֵׁנִים וְרַעֲנַנִּים", "fat and luxuriant".
- apparatus: Avot 2:10 (allusion, 2:28.224-225): the author's "פן יכוה בגחלתן" recalls the saying to beware of the coals of the sages; the words are his, not a quotation.
- apparatus: Ps 119:126 (allusion, 2:28.171-172): "עֵת לַעֲשׂוֹת לַה'", "It is a time to act for the LORD".
- apparatus: Folio references as the author gives them: folio 17b (2:28.498-500), folio 25 (2:28.458-459, 2:28.541-543), folio 26 (2:28.582-583). No reference is added.
- apparatus: The doubled "כל כל" (2:28.567-568) is as printed; both words are translated.
- apparatus: ומונים (2:28.75): the word is uncertain; rendered "and multitudes".
- apparatus: המשרף (2:28.402, 2:28.427): the reply does not name him; the transliteration "Mesharef" is kept.
- apparatus: ותשבורת and ותוכנים (2:28.118-119): technical terms; "tishburet" is kept, and "computations" is a guess at the second.
- apparatus: The Talmudic sayings (2:28.233-234, 2:28.241-260, 2:28.258-260) are given without a reference, since the author gives none.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

