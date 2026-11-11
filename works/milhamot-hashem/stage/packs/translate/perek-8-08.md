# Translation pack: milhamot-hashem perek-8-08 (Milchamot ha-Shem 8:10–8:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-08`.

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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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
- A Biblical quotation whose wording differs from the verse follows the
  author's wording; the difference goes in `apparatus`.
- `literal` on a long paragraph: the glosses in Hebrew order, hyphens joining
  the parts of one word. No reasoning in it; reasons go in `note`.
- `context` may carry a short note on a story the author cites (the Pardes
  story), without a reference unless the author gives one.

## Rulings from the pilot reviews, by the editor

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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


## Milchamot ha-Shem 8:10

Hebrew (Wikisource): קו) ושרידי משכילי עמינו אשר הרגישו בזרות הקבלה החדשה כרבי' תם ן' יחיא, ורבי פרץ, ורבי שמשון מקינון, והרישב"ש והרש"ל הרחיקו את עצמם מללמדה כמו שאמר הריב"ש איני תוקע את עצמי וכו', כי הרגיש בזרותה מדרך תורתינו הקדושה והטהורה מכל סיג וחלאה, אך לא דבר כנגדה רק כלאחר יד באמרו וכי יש אלהות לספירות? והמשל שאמר לו דון יוסף ן' שושן לא ערב לו, שידע באמת שהוא כיחש לו, אך לא חפץ להעציבו לומר שקר דברת. ולפי' אמר לו והנה טוב מאד אם הוא כדבריך. כי ידע כי רק דחיה בעלמא קא מדחי ליה, כי גם אמרם כי לנשמתו אתה עובד דחייה בעלמא היא. והמה באמת עובדים להם ומקבלים אלהותם עליהם. ומיחדים את הכלים עם נשמותיהם כמו שכתב שושן סודות והרש"ב בפרשת בהר דף ק"ט ומשנת חסידים סוף מסכת העשייה ועז לאלהים שהגוף והנשמה חד אינון. וכ"כ היכל הברכה פרשת ויקרא. ואלו היתה מניעת הריב"ש מטעם שלא קיבלה מחכם מקובל היה לו ללמדה ולקבלה מפי דון יוסף ן' שושן! אלא ודאי דברים בגו! שראה בחכמתו שזו היא יציאה מדת תורתינו הקדושה לעבוד אלהים אחרים, מעשה ידי מחשבות אדם ורעיונותיו ודמיונותיו אשר עליהם באה האזהרה בתורה ולא תתורו אחרי לבבכם ואחרי עיניכם ובאה הקבלה אחרי לבבכם זו מינות, ואחרי עיניכם זו זנות, כי לבא ועינא תרי סרסורי דחטאה נינהו, ומן הידוע שההולך אחר דעתו להמציא לו רעיונות בתורה. שלא על דרך רז"ל שקבלו תורה שבעל פה שהיא המשנה והתלמוד כמפורסם ודאי יבוא לידי מינות (כדגרסינן במדרש הגדול פרשת שמיני יין ושכר אל תשת. רבי חייא פתח פקודי ה' ישרים משמחי לב, וכתיב יתהלך במישרים. משמחי לב, וכתיב ויין ישמח לבב אנוש. מה היין הזה משמח את הלב. כך דברי תורה העוסק בהן ומחשב בהן יותר מדאי סוף שהוא יוצא למינות ע"כ וכבר הבאתיו בהקדמה). ועל זה רמז הריב"ש במלות קצרות וכי יש אלהות לספירות? ולכך אמר שהמפרש את הקבלה וכו' יהיו לו לבדו, ואיני תוקע את עצמי וכו' וכן עשה רבינו תם ן' יחיא ומהרש"ל המובאים דבריהם בספר רביד הזהב ור"י אלבו בספר העיקרים.

| id | word | punctuation after |
|---|---|---|
| 8:10.1 | קו | ) |
| 8:10.2 | ושרידי |  |
| 8:10.3 | משכילי |  |
| 8:10.4 | עמינו |  |
| 8:10.5 | אשר |  |
| 8:10.6 | הרגישו |  |
| 8:10.7 | בזרות |  |
| 8:10.8 | הקבלה |  |
| 8:10.9 | החדשה |  |
| 8:10.10 | כרבי | ' |
| 8:10.11 | תם |  |
| 8:10.12 | ן | ' |
| 8:10.13 | יחיא | , |
| 8:10.14 | ורבי |  |
| 8:10.15 | פרץ | , |
| 8:10.16 | ורבי |  |
| 8:10.17 | שמשון |  |
| 8:10.18 | מקינון | , |
| 8:10.19 | והרישב"ש |  |
| 8:10.20 | והרש"ל |  |
| 8:10.21 | הרחיקו |  |
| 8:10.22 | את |  |
| 8:10.23 | עצמם |  |
| 8:10.24 | מללמדה |  |
| 8:10.25 | כמו |  |
| 8:10.26 | שאמר |  |
| 8:10.27 | הריב"ש |  |
| 8:10.28 | איני |  |
| 8:10.29 | תוקע |  |
| 8:10.30 | את |  |
| 8:10.31 | עצמי |  |
| 8:10.32 | וכו | ', |
| 8:10.33 | כי |  |
| 8:10.34 | הרגיש |  |
| 8:10.35 | בזרותה |  |
| 8:10.36 | מדרך |  |
| 8:10.37 | תורתינו |  |
| 8:10.38 | הקדושה |  |
| 8:10.39 | והטהורה |  |
| 8:10.40 | מכל |  |
| 8:10.41 | סיג |  |
| 8:10.42 | וחלאה | , |
| 8:10.43 | אך |  |
| 8:10.44 | לא |  |
| 8:10.45 | דבר |  |
| 8:10.46 | כנגדה |  |
| 8:10.47 | רק |  |
| 8:10.48 | כלאחר |  |
| 8:10.49 | יד |  |
| 8:10.50 | באמרו |  |
| 8:10.51 | וכי |  |
| 8:10.52 | יש |  |
| 8:10.53 | אלהות |  |
| 8:10.54 | לספירות | ? |
| 8:10.55 | והמשל |  |
| 8:10.56 | שאמר |  |
| 8:10.57 | לו |  |
| 8:10.58 | דון |  |
| 8:10.59 | יוסף |  |
| 8:10.60 | ן | ' |
| 8:10.61 | שושן |  |
| 8:10.62 | לא |  |
| 8:10.63 | ערב |  |
| 8:10.64 | לו | , |
| 8:10.65 | שידע |  |
| 8:10.66 | באמת |  |
| 8:10.67 | שהוא |  |
| 8:10.68 | כיחש |  |
| 8:10.69 | לו | , |
| 8:10.70 | אך |  |
| 8:10.71 | לא |  |
| 8:10.72 | חפץ |  |
| 8:10.73 | להעציבו |  |
| 8:10.74 | לומר |  |
| 8:10.75 | שקר |  |
| 8:10.76 | דברת | . |
| 8:10.77 | ולפי | ' |
| 8:10.78 | אמר |  |
| 8:10.79 | לו |  |
| 8:10.80 | והנה |  |
| 8:10.81 | טוב |  |
| 8:10.82 | מאד |  |
| 8:10.83 | אם |  |
| 8:10.84 | הוא |  |
| 8:10.85 | כדבריך | . |
| 8:10.86 | כי |  |
| 8:10.87 | ידע |  |
| 8:10.88 | כי |  |
| 8:10.89 | רק |  |
| 8:10.90 | דחיה |  |
| 8:10.91 | בעלמא |  |
| 8:10.92 | קא |  |
| 8:10.93 | מדחי |  |
| 8:10.94 | ליה | , |
| 8:10.95 | כי |  |
| 8:10.96 | גם |  |
| 8:10.97 | אמרם |  |
| 8:10.98 | כי |  |
| 8:10.99 | לנשמתו |  |
| 8:10.100 | אתה |  |
| 8:10.101 | עובד |  |
| 8:10.102 | דחייה |  |
| 8:10.103 | בעלמא |  |
| 8:10.104 | היא | . |
| 8:10.105 | והמה |  |
| 8:10.106 | באמת |  |
| 8:10.107 | עובדים |  |
| 8:10.108 | להם |  |
| 8:10.109 | ומקבלים |  |
| 8:10.110 | אלהותם |  |
| 8:10.111 | עליהם | . |
| 8:10.112 | ומיחדים |  |
| 8:10.113 | את |  |
| 8:10.114 | הכלים |  |
| 8:10.115 | עם |  |
| 8:10.116 | נשמותיהם |  |
| 8:10.117 | כמו |  |
| 8:10.118 | שכתב |  |
| 8:10.119 | שושן |  |
| 8:10.120 | סודות |  |
| 8:10.121 | והרש"ב |  |
| 8:10.122 | בפרשת |  |
| 8:10.123 | בהר |  |
| 8:10.124 | דף |  |
| 8:10.125 | ק"ט |  |
| 8:10.126 | ומשנת |  |
| 8:10.127 | חסידים |  |
| 8:10.128 | סוף |  |
| 8:10.129 | מסכת |  |
| 8:10.130 | העשייה |  |
| 8:10.131 | ועז |  |
| 8:10.132 | לאלהים |  |
| 8:10.133 | שהגוף |  |
| 8:10.134 | והנשמה |  |
| 8:10.135 | חד |  |
| 8:10.136 | אינון | . |
| 8:10.137 | וכ"כ |  |
| 8:10.138 | היכל |  |
| 8:10.139 | הברכה |  |
| 8:10.140 | פרשת |  |
| 8:10.141 | ויקרא | . |
| 8:10.142 | ואלו |  |
| 8:10.143 | היתה |  |
| 8:10.144 | מניעת |  |
| 8:10.145 | הריב"ש |  |
| 8:10.146 | מטעם |  |
| 8:10.147 | שלא |  |
| 8:10.148 | קיבלה |  |
| 8:10.149 | מחכם |  |
| 8:10.150 | מקובל |  |
| 8:10.151 | היה |  |
| 8:10.152 | לו |  |
| 8:10.153 | ללמדה |  |
| 8:10.154 | ולקבלה |  |
| 8:10.155 | מפי |  |
| 8:10.156 | דון |  |
| 8:10.157 | יוסף |  |
| 8:10.158 | ן | ' |
| 8:10.159 | שושן | ! |
| 8:10.160 | אלא |  |
| 8:10.161 | ודאי |  |
| 8:10.162 | דברים |  |
| 8:10.163 | בגו | ! |
| 8:10.164 | שראה |  |
| 8:10.165 | בחכמתו |  |
| 8:10.166 | שזו |  |
| 8:10.167 | היא |  |
| 8:10.168 | יציאה |  |
| 8:10.169 | מדת |  |
| 8:10.170 | תורתינו |  |
| 8:10.171 | הקדושה |  |
| 8:10.172 | לעבוד |  |
| 8:10.173 | אלהים |  |
| 8:10.174 | אחרים | , |
| 8:10.175 | מעשה |  |
| 8:10.176 | ידי |  |
| 8:10.177 | מחשבות |  |
| 8:10.178 | אדם |  |
| 8:10.179 | ורעיונותיו |  |
| 8:10.180 | ודמיונותיו |  |
| 8:10.181 | אשר |  |
| 8:10.182 | עליהם |  |
| 8:10.183 | באה |  |
| 8:10.184 | האזהרה |  |
| 8:10.185 | בתורה |  |
| 8:10.186 | ולא |  |
| 8:10.187 | תתורו |  |
| 8:10.188 | אחרי |  |
| 8:10.189 | לבבכם |  |
| 8:10.190 | ואחרי |  |
| 8:10.191 | עיניכם |  |
| 8:10.192 | ובאה |  |
| 8:10.193 | הקבלה |  |
| 8:10.194 | אחרי |  |
| 8:10.195 | לבבכם |  |
| 8:10.196 | זו |  |
| 8:10.197 | מינות | , |
| 8:10.198 | ואחרי |  |
| 8:10.199 | עיניכם |  |
| 8:10.200 | זו |  |
| 8:10.201 | זנות | , |
| 8:10.202 | כי |  |
| 8:10.203 | לבא |  |
| 8:10.204 | ועינא |  |
| 8:10.205 | תרי |  |
| 8:10.206 | סרסורי |  |
| 8:10.207 | דחטאה |  |
| 8:10.208 | נינהו | , |
| 8:10.209 | ומן |  |
| 8:10.210 | הידוע |  |
| 8:10.211 | שההולך |  |
| 8:10.212 | אחר |  |
| 8:10.213 | דעתו |  |
| 8:10.214 | להמציא |  |
| 8:10.215 | לו |  |
| 8:10.216 | רעיונות |  |
| 8:10.217 | בתורה | . |
| 8:10.218 | שלא |  |
| 8:10.219 | על |  |
| 8:10.220 | דרך |  |
| 8:10.221 | רז"ל |  |
| 8:10.222 | שקבלו |  |
| 8:10.223 | תורה |  |
| 8:10.224 | שבעל |  |
| 8:10.225 | פה |  |
| 8:10.226 | שהיא |  |
| 8:10.227 | המשנה |  |
| 8:10.228 | והתלמוד |  |
| 8:10.229 | כמפורסם |  |
| 8:10.230 | ודאי |  |
| 8:10.231 | יבוא |  |
| 8:10.232 | לידי |  |
| 8:10.233 | מינות |  |
| 8:10.234 | כדגרסינן |  |
| 8:10.235 | במדרש |  |
| 8:10.236 | הגדול |  |
| 8:10.237 | פרשת |  |
| 8:10.238 | שמיני |  |
| 8:10.239 | יין |  |
| 8:10.240 | ושכר |  |
| 8:10.241 | אל |  |
| 8:10.242 | תשת | . |
| 8:10.243 | רבי |  |
| 8:10.244 | חייא |  |
| 8:10.245 | פתח |  |
| 8:10.246 | פקודי |  |
| 8:10.247 | ה | ' |
| 8:10.248 | ישרים |  |
| 8:10.249 | משמחי |  |
| 8:10.250 | לב | , |
| 8:10.251 | וכתיב |  |
| 8:10.252 | יתהלך |  |
| 8:10.253 | במישרים | . |
| 8:10.254 | משמחי |  |
| 8:10.255 | לב | , |
| 8:10.256 | וכתיב |  |
| 8:10.257 | ויין |  |
| 8:10.258 | ישמח |  |
| 8:10.259 | לבב |  |
| 8:10.260 | אנוש | . |
| 8:10.261 | מה |  |
| 8:10.262 | היין |  |
| 8:10.263 | הזה |  |
| 8:10.264 | משמח |  |
| 8:10.265 | את |  |
| 8:10.266 | הלב | . |
| 8:10.267 | כך |  |
| 8:10.268 | דברי |  |
| 8:10.269 | תורה |  |
| 8:10.270 | העוסק |  |
| 8:10.271 | בהן |  |
| 8:10.272 | ומחשב |  |
| 8:10.273 | בהן |  |
| 8:10.274 | יותר |  |
| 8:10.275 | מדאי |  |
| 8:10.276 | סוף |  |
| 8:10.277 | שהוא |  |
| 8:10.278 | יוצא |  |
| 8:10.279 | למינות |  |
| 8:10.280 | ע"כ |  |
| 8:10.281 | וכבר |  |
| 8:10.282 | הבאתיו |  |
| 8:10.283 | בהקדמה | ). |
| 8:10.284 | ועל |  |
| 8:10.285 | זה |  |
| 8:10.286 | רמז |  |
| 8:10.287 | הריב"ש |  |
| 8:10.288 | במלות |  |
| 8:10.289 | קצרות |  |
| 8:10.290 | וכי |  |
| 8:10.291 | יש |  |
| 8:10.292 | אלהות |  |
| 8:10.293 | לספירות | ? |
| 8:10.294 | ולכך |  |
| 8:10.295 | אמר |  |
| 8:10.296 | שהמפרש |  |
| 8:10.297 | את |  |
| 8:10.298 | הקבלה |  |
| 8:10.299 | וכו | ' |
| 8:10.300 | יהיו |  |
| 8:10.301 | לו |  |
| 8:10.302 | לבדו | , |
| 8:10.303 | ואיני |  |
| 8:10.304 | תוקע |  |
| 8:10.305 | את |  |
| 8:10.306 | עצמי |  |
| 8:10.307 | וכו | ' |
| 8:10.308 | וכן |  |
| 8:10.309 | עשה |  |
| 8:10.310 | רבינו |  |
| 8:10.311 | תם |  |
| 8:10.312 | ן | ' |
| 8:10.313 | יחיא |  |
| 8:10.314 | ומהרש"ל |  |
| 8:10.315 | המובאים |  |
| 8:10.316 | דבריהם |  |
| 8:10.317 | בספר |  |
| 8:10.318 | רביד |  |
| 8:10.319 | הזהב |  |
| 8:10.320 | ור"י |  |
| 8:10.321 | אלבו |  |
| 8:10.322 | בספר |  |
| 8:10.323 | העיקרים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

