# Translation pack: milhamot-hashem perek-7-12 (Milchamot ha-Shem 7:16–7:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-12`.

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
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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
- האלהי before a name: "the divine", with a note; never dropped.
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


## Milchamot ha-Shem 7:16

Hebrew (Wikisource): פט) העיקר החמישי כי זה אלהינו האחד המיוחד הראשון לכל הנמצאים. הוא אשר ראוי לעבדו ולרוממו, ולהרים קול ברוממותו ובעבודתו. ולא יעשה כן לזולתו מן הנמצאים. כגון המלאכים והגלגלים והיסודות וכל מה שהורכב מהם. כי כולם הם מוטבעות ומוגבלות על פעולתם, ואין להם רשות ולא חפץ, אלא רצונו ית'. ולא יושמו מליצים אמצעיים בינינו לבינו. אלא כל הרעיונים ישימו מגמתם לעומתו ית' ולא ישיתו לב לזולתו ע"כ. גם על זה העיקר שנעקר על פי הקבלה החדשה הנכריה, נשמו אחרונים. וקדמונים אחזו שער. ולא שמעו רז"ל בעלי המשנה והתלמוד באמור להם. שהעובד לאלהים עליון רם על רמים הסבה הראשונה ומתפלל אליו לא ייענה. ולא די זה גם ענוש יענש כמו שכתבו המקובלים החדשים. ורק לסבה האחרונה קצר אפים (זעיר אנפין) תאות העבודה! כי עתה קרעו בגדיהם כמו שכתבתי לעיל שזהו דעת המקובלים החדשים.

| id | word | punctuation after |
|---|---|---|
| 7:16.1 | פט | ) |
| 7:16.2 | העיקר |  |
| 7:16.3 | החמישי |  |
| 7:16.4 | כי |  |
| 7:16.5 | זה |  |
| 7:16.6 | אלהינו |  |
| 7:16.7 | האחד |  |
| 7:16.8 | המיוחד |  |
| 7:16.9 | הראשון |  |
| 7:16.10 | לכל |  |
| 7:16.11 | הנמצאים | . |
| 7:16.12 | הוא |  |
| 7:16.13 | אשר |  |
| 7:16.14 | ראוי |  |
| 7:16.15 | לעבדו |  |
| 7:16.16 | ולרוממו | , |
| 7:16.17 | ולהרים |  |
| 7:16.18 | קול |  |
| 7:16.19 | ברוממותו |  |
| 7:16.20 | ובעבודתו | . |
| 7:16.21 | ולא |  |
| 7:16.22 | יעשה |  |
| 7:16.23 | כן |  |
| 7:16.24 | לזולתו |  |
| 7:16.25 | מן |  |
| 7:16.26 | הנמצאים | . |
| 7:16.27 | כגון |  |
| 7:16.28 | המלאכים |  |
| 7:16.29 | והגלגלים |  |
| 7:16.30 | והיסודות |  |
| 7:16.31 | וכל |  |
| 7:16.32 | מה |  |
| 7:16.33 | שהורכב |  |
| 7:16.34 | מהם | . |
| 7:16.35 | כי |  |
| 7:16.36 | כולם |  |
| 7:16.37 | הם |  |
| 7:16.38 | מוטבעות |  |
| 7:16.39 | ומוגבלות |  |
| 7:16.40 | על |  |
| 7:16.41 | פעולתם | , |
| 7:16.42 | ואין |  |
| 7:16.43 | להם |  |
| 7:16.44 | רשות |  |
| 7:16.45 | ולא |  |
| 7:16.46 | חפץ | , |
| 7:16.47 | אלא |  |
| 7:16.48 | רצונו |  |
| 7:16.49 | ית | '. |
| 7:16.50 | ולא |  |
| 7:16.51 | יושמו |  |
| 7:16.52 | מליצים |  |
| 7:16.53 | אמצעיים |  |
| 7:16.54 | בינינו |  |
| 7:16.55 | לבינו | . |
| 7:16.56 | אלא |  |
| 7:16.57 | כל |  |
| 7:16.58 | הרעיונים |  |
| 7:16.59 | ישימו |  |
| 7:16.60 | מגמתם |  |
| 7:16.61 | לעומתו |  |
| 7:16.62 | ית | ' |
| 7:16.63 | ולא |  |
| 7:16.64 | ישיתו |  |
| 7:16.65 | לב |  |
| 7:16.66 | לזולתו |  |
| 7:16.67 | ע"כ | . |
| 7:16.68 | גם |  |
| 7:16.69 | על |  |
| 7:16.70 | זה |  |
| 7:16.71 | העיקר |  |
| 7:16.72 | שנעקר |  |
| 7:16.73 | על |  |
| 7:16.74 | פי |  |
| 7:16.75 | הקבלה |  |
| 7:16.76 | החדשה |  |
| 7:16.77 | הנכריה | , |
| 7:16.78 | נשמו |  |
| 7:16.79 | אחרונים | . |
| 7:16.80 | וקדמונים |  |
| 7:16.81 | אחזו |  |
| 7:16.82 | שער | . |
| 7:16.83 | ולא |  |
| 7:16.84 | שמעו |  |
| 7:16.85 | רז"ל |  |
| 7:16.86 | בעלי |  |
| 7:16.87 | המשנה |  |
| 7:16.88 | והתלמוד |  |
| 7:16.89 | באמור |  |
| 7:16.90 | להם | . |
| 7:16.91 | שהעובד |  |
| 7:16.92 | לאלהים |  |
| 7:16.93 | עליון |  |
| 7:16.94 | רם |  |
| 7:16.95 | על |  |
| 7:16.96 | רמים |  |
| 7:16.97 | הסבה |  |
| 7:16.98 | הראשונה |  |
| 7:16.99 | ומתפלל |  |
| 7:16.100 | אליו |  |
| 7:16.101 | לא |  |
| 7:16.102 | ייענה | . |
| 7:16.103 | ולא |  |
| 7:16.104 | די |  |
| 7:16.105 | זה |  |
| 7:16.106 | גם |  |
| 7:16.107 | ענוש |  |
| 7:16.108 | יענש |  |
| 7:16.109 | כמו |  |
| 7:16.110 | שכתבו |  |
| 7:16.111 | המקובלים |  |
| 7:16.112 | החדשים | . |
| 7:16.113 | ורק |  |
| 7:16.114 | לסבה |  |
| 7:16.115 | האחרונה |  |
| 7:16.116 | קצר |  |
| 7:16.117 | אפים |  |
| 7:16.118 | זעיר |  |
| 7:16.119 | אנפין | ) |
| 7:16.120 | תאות |  |
| 7:16.121 | העבודה | ! |
| 7:16.122 | כי |  |
| 7:16.123 | עתה |  |
| 7:16.124 | קרעו |  |
| 7:16.125 | בגדיהם |  |
| 7:16.126 | כמו |  |
| 7:16.127 | שכתבתי |  |
| 7:16.128 | לעיל |  |
| 7:16.129 | שזהו |  |
| 7:16.130 | דעת |  |
| 7:16.131 | המקובלים |  |
| 7:16.132 | החדשים | . |


## Milchamot ha-Shem 7:17

Hebrew (Wikisource): צ) הן עתה נדבר על ענין שם האלוה אשר המקובלים החדשים אומרים לעבדו על פי הזהר. וקראו לו בשם מגונה מאד והוא קצר אפים (זעיר אנפין) כי ידוע ומפורסם מאוד, שתורתינו הקדושה וכל ספרי הנביאים ורז"ל בגמרא ובמדרשים ובנוסח התפלות מתארים ומספרים שבחי השי"ת בתארים נעלים. וכתב הרמב"ם ז"ל בפ' כ"ו לראשון מספר המורה. כל שהוא שלמות אצלינו ייחס לו ית' להורות עליו שהוא שלם בכל מיני שלמות ואין בו חסרון והעדר כלל. וכל מה שישיגו ההמון שהוא חסרון או העדר לא יתואר בו. ולזה לא יתואר השי"ת באכילה ושתיה ושינה. ולא במה שידמה לזה. וכל מה שיחשוב ההמון שהוא שלמות יתואר בו, ואך על פי שהדבר ההוא אינו שלמות רק בערך אלינו. אבל בערך אליו ית'. אלו אשר נחשבם כולם שלימות הם תכלית החסרון ע"כ. וכתב הרס"ג שזהו מאמר ויברכו שם כבודיך ומרומם על כל ברכה ותהלה, רוצה לומר שכל מה שיספרו עליו המספרים. וישבחוה המשבחים בו, הוא נעלה ומרומם וגדול מן הכל ע"כ. וכן כ' הר' שבילי אמונה כי ההכרח הביא אותנו להגשים הבורא ית' ולספר אותו במדות הברואים. כדי לשער ענין שיכנס מציאות הבורא ית' בנפשות. וכיון שלא מצאנו במה שאנו מדברים בה מן הלשונות מלה שתורה על אמתת ענין הבורא ית' הוצאנו הענין ביותר ממלה א'. והרבוי הנמצא במדות הבורא ית' איננו מצד עצם כבודו, רק מצד קוצר כח מליצת המספר מהשיג עניינו במלה אחת שתורה עליו. והמדות האלהיות הפועלות שיסופר בהם הבורא ית' אפשר שישתתף בספוריו עם קצת ברואיו. והתירו חכמים לספר אותו בהם מפני הדוחק המצריך אותנו להודיעו ולעמוד על מציאותו כדי שנקבל עבודתו. ולכן מצאנו שמשתמשין במין הזה ממדות השי"ת בספר התורה ובספרי הנביאים הרבה מאד ובתושבחות הנביאים. והוציאו אותו הנביאים לבני אדם במלות גשמיות שהם קרובים לשכלם והבנתם. ואלו היו מספרים בענין שראוי לו מן המלות הרוחניות והענינים הרוחניים לא היינו מבינים לא הענין ולא המלות. ולא היה אפשר שנעבוד דבר שלא נדע כלל עכ"ל. דבר הלמד מדברי הרבנים הנזכרים שלא יתואר השי"ת בשום תואר שהוא חסרון וסכלות כתואר זעיר אנפין שענינו קצר אפים. והנה התארים והמלות אשר בחר הקב"ה להודיע למשה ולישראל ממדותיו. הם ה' ה' אל רחום וחנון ארך אפים ורב חסד ואמת. נצר חסד לאלפים נשא עון ופשע וחטאה ונקה. וכן אמר דוד המלך ע"ה יודיע דרכיו למשה לבני שיראל עלילותיו. רחום וחנון ה'. ארך אפים ורב חסד. ופי' הראב"ע ז"ל יודיע דרכיו, הם הדרכים הנז' בפסוק ויעבר ה' על פניו וכו'. רחום וחנון ה' ארך אפים ורב חסד. והזכיר קצתן הצריך להן בתפלתו ע"כ.

| id | word | punctuation after |
|---|---|---|
| 7:17.1 | צ | ) |
| 7:17.2 | הן |  |
| 7:17.3 | עתה |  |
| 7:17.4 | נדבר |  |
| 7:17.5 | על |  |
| 7:17.6 | ענין |  |
| 7:17.7 | שם |  |
| 7:17.8 | האלוה |  |
| 7:17.9 | אשר |  |
| 7:17.10 | המקובלים |  |
| 7:17.11 | החדשים |  |
| 7:17.12 | אומרים |  |
| 7:17.13 | לעבדו |  |
| 7:17.14 | על |  |
| 7:17.15 | פי |  |
| 7:17.16 | הזהר | . |
| 7:17.17 | וקראו |  |
| 7:17.18 | לו |  |
| 7:17.19 | בשם |  |
| 7:17.20 | מגונה |  |
| 7:17.21 | מאד |  |
| 7:17.22 | והוא |  |
| 7:17.23 | קצר |  |
| 7:17.24 | אפים |  |
| 7:17.25 | זעיר |  |
| 7:17.26 | אנפין | ) |
| 7:17.27 | כי |  |
| 7:17.28 | ידוע |  |
| 7:17.29 | ומפורסם |  |
| 7:17.30 | מאוד | , |
| 7:17.31 | שתורתינו |  |
| 7:17.32 | הקדושה |  |
| 7:17.33 | וכל |  |
| 7:17.34 | ספרי |  |
| 7:17.35 | הנביאים |  |
| 7:17.36 | ורז"ל |  |
| 7:17.37 | בגמרא |  |
| 7:17.38 | ובמדרשים |  |
| 7:17.39 | ובנוסח |  |
| 7:17.40 | התפלות |  |
| 7:17.41 | מתארים |  |
| 7:17.42 | ומספרים |  |
| 7:17.43 | שבחי |  |
| 7:17.44 | השי"ת |  |
| 7:17.45 | בתארים |  |
| 7:17.46 | נעלים | . |
| 7:17.47 | וכתב |  |
| 7:17.48 | הרמב"ם |  |
| 7:17.49 | ז"ל |  |
| 7:17.50 | בפ | ' |
| 7:17.51 | כ"ו |  |
| 7:17.52 | לראשון |  |
| 7:17.53 | מספר |  |
| 7:17.54 | המורה | . |
| 7:17.55 | כל |  |
| 7:17.56 | שהוא |  |
| 7:17.57 | שלמות |  |
| 7:17.58 | אצלינו |  |
| 7:17.59 | ייחס |  |
| 7:17.60 | לו |  |
| 7:17.61 | ית | ' |
| 7:17.62 | להורות |  |
| 7:17.63 | עליו |  |
| 7:17.64 | שהוא |  |
| 7:17.65 | שלם |  |
| 7:17.66 | בכל |  |
| 7:17.67 | מיני |  |
| 7:17.68 | שלמות |  |
| 7:17.69 | ואין |  |
| 7:17.70 | בו |  |
| 7:17.71 | חסרון |  |
| 7:17.72 | והעדר |  |
| 7:17.73 | כלל | . |
| 7:17.74 | וכל |  |
| 7:17.75 | מה |  |
| 7:17.76 | שישיגו |  |
| 7:17.77 | ההמון |  |
| 7:17.78 | שהוא |  |
| 7:17.79 | חסרון |  |
| 7:17.80 | או |  |
| 7:17.81 | העדר |  |
| 7:17.82 | לא |  |
| 7:17.83 | יתואר |  |
| 7:17.84 | בו | . |
| 7:17.85 | ולזה |  |
| 7:17.86 | לא |  |
| 7:17.87 | יתואר |  |
| 7:17.88 | השי"ת |  |
| 7:17.89 | באכילה |  |
| 7:17.90 | ושתיה |  |
| 7:17.91 | ושינה | . |
| 7:17.92 | ולא |  |
| 7:17.93 | במה |  |
| 7:17.94 | שידמה |  |
| 7:17.95 | לזה | . |
| 7:17.96 | וכל |  |
| 7:17.97 | מה |  |
| 7:17.98 | שיחשוב |  |
| 7:17.99 | ההמון |  |
| 7:17.100 | שהוא |  |
| 7:17.101 | שלמות |  |
| 7:17.102 | יתואר |  |
| 7:17.103 | בו | , |
| 7:17.104 | ואך |  |
| 7:17.105 | על |  |
| 7:17.106 | פי |  |
| 7:17.107 | שהדבר |  |
| 7:17.108 | ההוא |  |
| 7:17.109 | אינו |  |
| 7:17.110 | שלמות |  |
| 7:17.111 | רק |  |
| 7:17.112 | בערך |  |
| 7:17.113 | אלינו | . |
| 7:17.114 | אבל |  |
| 7:17.115 | בערך |  |
| 7:17.116 | אליו |  |
| 7:17.117 | ית | '. |
| 7:17.118 | אלו |  |
| 7:17.119 | אשר |  |
| 7:17.120 | נחשבם |  |
| 7:17.121 | כולם |  |
| 7:17.122 | שלימות |  |
| 7:17.123 | הם |  |
| 7:17.124 | תכלית |  |
| 7:17.125 | החסרון |  |
| 7:17.126 | ע"כ | . |
| 7:17.127 | וכתב |  |
| 7:17.128 | הרס"ג |  |
| 7:17.129 | שזהו |  |
| 7:17.130 | מאמר |  |
| 7:17.131 | ויברכו |  |
| 7:17.132 | שם |  |
| 7:17.133 | כבודיך |  |
| 7:17.134 | ומרומם |  |
| 7:17.135 | על |  |
| 7:17.136 | כל |  |
| 7:17.137 | ברכה |  |
| 7:17.138 | ותהלה | , |
| 7:17.139 | רוצה |  |
| 7:17.140 | לומר |  |
| 7:17.141 | שכל |  |
| 7:17.142 | מה |  |
| 7:17.143 | שיספרו |  |
| 7:17.144 | עליו |  |
| 7:17.145 | המספרים | . |
| 7:17.146 | וישבחוה |  |
| 7:17.147 | המשבחים |  |
| 7:17.148 | בו | , |
| 7:17.149 | הוא |  |
| 7:17.150 | נעלה |  |
| 7:17.151 | ומרומם |  |
| 7:17.152 | וגדול |  |
| 7:17.153 | מן |  |
| 7:17.154 | הכל |  |
| 7:17.155 | ע"כ | . |
| 7:17.156 | וכן |  |
| 7:17.157 | כ | ' |
| 7:17.158 | הר | ' |
| 7:17.159 | שבילי |  |
| 7:17.160 | אמונה |  |
| 7:17.161 | כי |  |
| 7:17.162 | ההכרח |  |
| 7:17.163 | הביא |  |
| 7:17.164 | אותנו |  |
| 7:17.165 | להגשים |  |
| 7:17.166 | הבורא |  |
| 7:17.167 | ית | ' |
| 7:17.168 | ולספר |  |
| 7:17.169 | אותו |  |
| 7:17.170 | במדות |  |
| 7:17.171 | הברואים | . |
| 7:17.172 | כדי |  |
| 7:17.173 | לשער |  |
| 7:17.174 | ענין |  |
| 7:17.175 | שיכנס |  |
| 7:17.176 | מציאות |  |
| 7:17.177 | הבורא |  |
| 7:17.178 | ית | ' |
| 7:17.179 | בנפשות | . |
| 7:17.180 | וכיון |  |
| 7:17.181 | שלא |  |
| 7:17.182 | מצאנו |  |
| 7:17.183 | במה |  |
| 7:17.184 | שאנו |  |
| 7:17.185 | מדברים |  |
| 7:17.186 | בה |  |
| 7:17.187 | מן |  |
| 7:17.188 | הלשונות |  |
| 7:17.189 | מלה |  |
| 7:17.190 | שתורה |  |
| 7:17.191 | על |  |
| 7:17.192 | אמתת |  |
| 7:17.193 | ענין |  |
| 7:17.194 | הבורא |  |
| 7:17.195 | ית | ' |
| 7:17.196 | הוצאנו |  |
| 7:17.197 | הענין |  |
| 7:17.198 | ביותר |  |
| 7:17.199 | ממלה |  |
| 7:17.200 | א | '. |
| 7:17.201 | והרבוי |  |
| 7:17.202 | הנמצא |  |
| 7:17.203 | במדות |  |
| 7:17.204 | הבורא |  |
| 7:17.205 | ית | ' |
| 7:17.206 | איננו |  |
| 7:17.207 | מצד |  |
| 7:17.208 | עצם |  |
| 7:17.209 | כבודו | , |
| 7:17.210 | רק |  |
| 7:17.211 | מצד |  |
| 7:17.212 | קוצר |  |
| 7:17.213 | כח |  |
| 7:17.214 | מליצת |  |
| 7:17.215 | המספר |  |
| 7:17.216 | מהשיג |  |
| 7:17.217 | עניינו |  |
| 7:17.218 | במלה |  |
| 7:17.219 | אחת |  |
| 7:17.220 | שתורה |  |
| 7:17.221 | עליו | . |
| 7:17.222 | והמדות |  |
| 7:17.223 | האלהיות |  |
| 7:17.224 | הפועלות |  |
| 7:17.225 | שיסופר |  |
| 7:17.226 | בהם |  |
| 7:17.227 | הבורא |  |
| 7:17.228 | ית | ' |
| 7:17.229 | אפשר |  |
| 7:17.230 | שישתתף |  |
| 7:17.231 | בספוריו |  |
| 7:17.232 | עם |  |
| 7:17.233 | קצת |  |
| 7:17.234 | ברואיו | . |
| 7:17.235 | והתירו |  |
| 7:17.236 | חכמים |  |
| 7:17.237 | לספר |  |
| 7:17.238 | אותו |  |
| 7:17.239 | בהם |  |
| 7:17.240 | מפני |  |
| 7:17.241 | הדוחק |  |
| 7:17.242 | המצריך |  |
| 7:17.243 | אותנו |  |
| 7:17.244 | להודיעו |  |
| 7:17.245 | ולעמוד |  |
| 7:17.246 | על |  |
| 7:17.247 | מציאותו |  |
| 7:17.248 | כדי |  |
| 7:17.249 | שנקבל |  |
| 7:17.250 | עבודתו | . |
| 7:17.251 | ולכן |  |
| 7:17.252 | מצאנו |  |
| 7:17.253 | שמשתמשין |  |
| 7:17.254 | במין |  |
| 7:17.255 | הזה |  |
| 7:17.256 | ממדות |  |
| 7:17.257 | השי"ת |  |
| 7:17.258 | בספר |  |
| 7:17.259 | התורה |  |
| 7:17.260 | ובספרי |  |
| 7:17.261 | הנביאים |  |
| 7:17.262 | הרבה |  |
| 7:17.263 | מאד |  |
| 7:17.264 | ובתושבחות |  |
| 7:17.265 | הנביאים | . |
| 7:17.266 | והוציאו |  |
| 7:17.267 | אותו |  |
| 7:17.268 | הנביאים |  |
| 7:17.269 | לבני |  |
| 7:17.270 | אדם |  |
| 7:17.271 | במלות |  |
| 7:17.272 | גשמיות |  |
| 7:17.273 | שהם |  |
| 7:17.274 | קרובים |  |
| 7:17.275 | לשכלם |  |
| 7:17.276 | והבנתם | . |
| 7:17.277 | ואלו |  |
| 7:17.278 | היו |  |
| 7:17.279 | מספרים |  |
| 7:17.280 | בענין |  |
| 7:17.281 | שראוי |  |
| 7:17.282 | לו |  |
| 7:17.283 | מן |  |
| 7:17.284 | המלות |  |
| 7:17.285 | הרוחניות |  |
| 7:17.286 | והענינים |  |
| 7:17.287 | הרוחניים |  |
| 7:17.288 | לא |  |
| 7:17.289 | היינו |  |
| 7:17.290 | מבינים |  |
| 7:17.291 | לא |  |
| 7:17.292 | הענין |  |
| 7:17.293 | ולא |  |
| 7:17.294 | המלות | . |
| 7:17.295 | ולא |  |
| 7:17.296 | היה |  |
| 7:17.297 | אפשר |  |
| 7:17.298 | שנעבוד |  |
| 7:17.299 | דבר |  |
| 7:17.300 | שלא |  |
| 7:17.301 | נדע |  |
| 7:17.302 | כלל |  |
| 7:17.303 | עכ"ל | . |
| 7:17.304 | דבר |  |
| 7:17.305 | הלמד |  |
| 7:17.306 | מדברי |  |
| 7:17.307 | הרבנים |  |
| 7:17.308 | הנזכרים |  |
| 7:17.309 | שלא |  |
| 7:17.310 | יתואר |  |
| 7:17.311 | השי"ת |  |
| 7:17.312 | בשום |  |
| 7:17.313 | תואר |  |
| 7:17.314 | שהוא |  |
| 7:17.315 | חסרון |  |
| 7:17.316 | וסכלות |  |
| 7:17.317 | כתואר |  |
| 7:17.318 | זעיר |  |
| 7:17.319 | אנפין |  |
| 7:17.320 | שענינו |  |
| 7:17.321 | קצר |  |
| 7:17.322 | אפים | . |
| 7:17.323 | והנה |  |
| 7:17.324 | התארים |  |
| 7:17.325 | והמלות |  |
| 7:17.326 | אשר |  |
| 7:17.327 | בחר |  |
| 7:17.328 | הקב"ה |  |
| 7:17.329 | להודיע |  |
| 7:17.330 | למשה |  |
| 7:17.331 | ולישראל |  |
| 7:17.332 | ממדותיו | . |
| 7:17.333 | הם |  |
| 7:17.334 | ה | ' |
| 7:17.335 | ה | ' |
| 7:17.336 | אל |  |
| 7:17.337 | רחום |  |
| 7:17.338 | וחנון |  |
| 7:17.339 | ארך |  |
| 7:17.340 | אפים |  |
| 7:17.341 | ורב |  |
| 7:17.342 | חסד |  |
| 7:17.343 | ואמת | . |
| 7:17.344 | נצר |  |
| 7:17.345 | חסד |  |
| 7:17.346 | לאלפים |  |
| 7:17.347 | נשא |  |
| 7:17.348 | עון |  |
| 7:17.349 | ופשע |  |
| 7:17.350 | וחטאה |  |
| 7:17.351 | ונקה | . |
| 7:17.352 | וכן |  |
| 7:17.353 | אמר |  |
| 7:17.354 | דוד |  |
| 7:17.355 | המלך |  |
| 7:17.356 | ע"ה |  |
| 7:17.357 | יודיע |  |
| 7:17.358 | דרכיו |  |
| 7:17.359 | למשה |  |
| 7:17.360 | לבני |  |
| 7:17.361 | שיראל |  |
| 7:17.362 | עלילותיו | . |
| 7:17.363 | רחום |  |
| 7:17.364 | וחנון |  |
| 7:17.365 | ה | '. |
| 7:17.366 | ארך |  |
| 7:17.367 | אפים |  |
| 7:17.368 | ורב |  |
| 7:17.369 | חסד | . |
| 7:17.370 | ופי | ' |
| 7:17.371 | הראב"ע |  |
| 7:17.372 | ז"ל |  |
| 7:17.373 | יודיע |  |
| 7:17.374 | דרכיו | , |
| 7:17.375 | הם |  |
| 7:17.376 | הדרכים |  |
| 7:17.377 | הנז | ' |
| 7:17.378 | בפסוק |  |
| 7:17.379 | ויעבר |  |
| 7:17.380 | ה | ' |
| 7:17.381 | על |  |
| 7:17.382 | פניו |  |
| 7:17.383 | וכו | '. |
| 7:17.384 | רחום |  |
| 7:17.385 | וחנון |  |
| 7:17.386 | ה | ' |
| 7:17.387 | ארך |  |
| 7:17.388 | אפים |  |
| 7:17.389 | ורב |  |
| 7:17.390 | חסד | . |
| 7:17.391 | והזכיר |  |
| 7:17.392 | קצתן |  |
| 7:17.393 | הצריך |  |
| 7:17.394 | להן |  |
| 7:17.395 | בתפלתו |  |
| 7:17.396 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

