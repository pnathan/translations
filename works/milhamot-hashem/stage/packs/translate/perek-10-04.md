# Translation pack: milhamot-hashem perek-10-04 (Milchamot ha-Shem 10:5–10:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-04`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- מרגל לרגל: "from festival to festival" (Num 20 ff. sense), not "foot".
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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


## Milchamot ha-Shem 10:5

Hebrew (Wikisource): קמג) ולו עיני רז"ל אנשי כנה"ג ראו את כל הנעשה בדת היהדות על פי הספר המזוייף מעבודת הכחות הנעלמים כי אז קרעו את בגדיהם ולבשו שק ואפר, וצעקו בחוצות ובשווקים וברחובות, צעקה גדולה ומרה כמרדכי בשושן הבירה, הזאת היא חקת התורה? להמציא בתורה רמזים לעבודת האלילים? ולימודי ידיעת האלילים האלה קראו להם סתרי תורה, מלשון סתירה שהם סותרים ועוקרים יסודותיה, והורסים פנותיה כמו שדרשו רז"ל (ילקוט פרשת שלח) אמר ר' יצחק דבר זה מסורת בידינו מאבותינו מרגלים על שם מעשיהם נקראו, ולא עלתה בידינו אלא אחד, סתור בן מיכאל סתור שסתר דבריו של הקב"ה, בן מיכאל, שעשה עצמו מך, אמר ר' יוחנן אף אנו נאמר נחבי בן ופסי, שהחביא דבריו של הקב"ה ופסע עד מדותיו של הקב"ה. כדרך הזאת עשה הפילסוף המסית מחבר הזהר והנמשכים אחריו לסתור ולקלקל את אמונותינו באל אחד מיוחד מכל שאר האחדים, ולפסוע על מדותיו של הקב"ה ולעשותן אלוהות במה שקדם, כי לא הספיק לו לחשוב באלהות עילה אחת, והוצרך להמציא ברעיונותיו עילות רבות שבה נוטל רשות מזה וזה נעזר במעשיהו על ידי אחר, כי חלש הוא ולא יוכל לעשות לבדו כמו שנתבאר לעיל, ובחר לו לעבוד את האליל קצר האפים אשר לפי דעתו לא נשתתף בשום מעשה מכל מעשה בראשית, כי הוא נברא ועדיין היה חסר תיקון ולא נשלם תיקוני עד ימי אברהם אבינו כמו שנתבאר לעיל בשם חיי שלום משם האר"י ולא די זה אלא שהוא מבזה ומהתל בשלש עשרה מדות שהתורה נדרשת בהם שקבלם משה בסיני, ועל פ' חי"ג מדות הללו נבנית המשנה והתלמוד הקדוש מימות משה ועד רבי אשי ורבינא מיסדי התלמוד כמו שכתב הרמב"ם ז"ל בהקדמת סדר ורעים ומחבד הזהר מלעיג עליהן וקורא עליהם וימררו את חייהם בעבודה קשה דא קושיא, בחמר דא קל וחומר וכו' (זהר בראשית כ"ז) והוציא לעז ודבה רעה על משה רבינו שקבל תורה שבעל פה מפי מטטרי"ן ולא מפי הקב"ה, ולכך היא מעורבת מטוב ורע, וקורא לה אבן משכית, ובה עסקו בן עזאי ובן זומא ואלישע אחר ונסתכנו! וגם קרא על המשנה תחת שלש רגזה ארץ וכו', ודימה שקלא וטריא דרבנן במשנה ובתלמוד ככלבין דצוחין הב הב, וקורא לה גם כן קליפה (זהר פרשת תצא ותיקוני זהר חדש) ועוד הרבה דברים להבזות את לימודי המשנה והתלמוד ולהפחית כבודם ולזלזל בהם, ולתת כבוד ויקר ללימוד תורתו החדשה שהיא מילדי נכרים, ולהוקיר את הבעלים והעשתרות שהמציא ברעיונותיו להטעות את ישראל, ולשבח את סחורתו אשר חשב להנחיל אותנו תחת המשנה והתלמוד ולרשת את מקומם. ועל זה אמרו רז"ל בסנהדרין פ' חלק (סנהדרין דף צ"ט) כי דבר ה' בזה, זה האומר אין תורה מן השמים וכו' ואפילו אמר כל התורה כולה מן השמים חוץ מדקדוק זה, חוץ מקך וחומר זה מגזרה שוה זו, זהו כי דבר ה' בזה וכו' רבי נתן אומר זה שאינו משגיח על המשנה וכו'. תנא דבי ר' ישמעאל כי דבר ה' בזה זה המבזה דיבור שנאמר למשה אנכי ולא יהיה לך ע"כ! וכל הני מילי דאמור רבנן איתנהו בפילסוף מחבר הזהר.

| id | word | punctuation after |
|---|---|---|
| 10:5.1 | קמג | ) |
| 10:5.2 | ולו |  |
| 10:5.3 | עיני |  |
| 10:5.4 | רז"ל |  |
| 10:5.5 | אנשי |  |
| 10:5.6 | כנה"ג |  |
| 10:5.7 | ראו |  |
| 10:5.8 | את |  |
| 10:5.9 | כל |  |
| 10:5.10 | הנעשה |  |
| 10:5.11 | בדת |  |
| 10:5.12 | היהדות |  |
| 10:5.13 | על |  |
| 10:5.14 | פי |  |
| 10:5.15 | הספר |  |
| 10:5.16 | המזוייף |  |
| 10:5.17 | מעבודת |  |
| 10:5.18 | הכחות |  |
| 10:5.19 | הנעלמים |  |
| 10:5.20 | כי |  |
| 10:5.21 | אז |  |
| 10:5.22 | קרעו |  |
| 10:5.23 | את |  |
| 10:5.24 | בגדיהם |  |
| 10:5.25 | ולבשו |  |
| 10:5.26 | שק |  |
| 10:5.27 | ואפר | , |
| 10:5.28 | וצעקו |  |
| 10:5.29 | בחוצות |  |
| 10:5.30 | ובשווקים |  |
| 10:5.31 | וברחובות | , |
| 10:5.32 | צעקה |  |
| 10:5.33 | גדולה |  |
| 10:5.34 | ומרה |  |
| 10:5.35 | כמרדכי |  |
| 10:5.36 | בשושן |  |
| 10:5.37 | הבירה | , |
| 10:5.38 | הזאת |  |
| 10:5.39 | היא |  |
| 10:5.40 | חקת |  |
| 10:5.41 | התורה | ? |
| 10:5.42 | להמציא |  |
| 10:5.43 | בתורה |  |
| 10:5.44 | רמזים |  |
| 10:5.45 | לעבודת |  |
| 10:5.46 | האלילים | ? |
| 10:5.47 | ולימודי |  |
| 10:5.48 | ידיעת |  |
| 10:5.49 | האלילים |  |
| 10:5.50 | האלה |  |
| 10:5.51 | קראו |  |
| 10:5.52 | להם |  |
| 10:5.53 | סתרי |  |
| 10:5.54 | תורה | , |
| 10:5.55 | מלשון |  |
| 10:5.56 | סתירה |  |
| 10:5.57 | שהם |  |
| 10:5.58 | סותרים |  |
| 10:5.59 | ועוקרים |  |
| 10:5.60 | יסודותיה | , |
| 10:5.61 | והורסים |  |
| 10:5.62 | פנותיה |  |
| 10:5.63 | כמו |  |
| 10:5.64 | שדרשו |  |
| 10:5.65 | רז"ל |  |
| 10:5.66 | ילקוט |  |
| 10:5.67 | פרשת |  |
| 10:5.68 | שלח | ) |
| 10:5.69 | אמר |  |
| 10:5.70 | ר | ' |
| 10:5.71 | יצחק |  |
| 10:5.72 | דבר |  |
| 10:5.73 | זה |  |
| 10:5.74 | מסורת |  |
| 10:5.75 | בידינו |  |
| 10:5.76 | מאבותינו |  |
| 10:5.77 | מרגלים |  |
| 10:5.78 | על |  |
| 10:5.79 | שם |  |
| 10:5.80 | מעשיהם |  |
| 10:5.81 | נקראו | , |
| 10:5.82 | ולא |  |
| 10:5.83 | עלתה |  |
| 10:5.84 | בידינו |  |
| 10:5.85 | אלא |  |
| 10:5.86 | אחד | , |
| 10:5.87 | סתור |  |
| 10:5.88 | בן |  |
| 10:5.89 | מיכאל |  |
| 10:5.90 | סתור |  |
| 10:5.91 | שסתר |  |
| 10:5.92 | דבריו |  |
| 10:5.93 | של |  |
| 10:5.94 | הקב"ה | , |
| 10:5.95 | בן |  |
| 10:5.96 | מיכאל | , |
| 10:5.97 | שעשה |  |
| 10:5.98 | עצמו |  |
| 10:5.99 | מך | , |
| 10:5.100 | אמר |  |
| 10:5.101 | ר | ' |
| 10:5.102 | יוחנן |  |
| 10:5.103 | אף |  |
| 10:5.104 | אנו |  |
| 10:5.105 | נאמר |  |
| 10:5.106 | נחבי |  |
| 10:5.107 | בן |  |
| 10:5.108 | ופסי | , |
| 10:5.109 | שהחביא |  |
| 10:5.110 | דבריו |  |
| 10:5.111 | של |  |
| 10:5.112 | הקב"ה |  |
| 10:5.113 | ופסע |  |
| 10:5.114 | עד |  |
| 10:5.115 | מדותיו |  |
| 10:5.116 | של |  |
| 10:5.117 | הקב"ה | . |
| 10:5.118 | כדרך |  |
| 10:5.119 | הזאת |  |
| 10:5.120 | עשה |  |
| 10:5.121 | הפילסוף |  |
| 10:5.122 | המסית |  |
| 10:5.123 | מחבר |  |
| 10:5.124 | הזהר |  |
| 10:5.125 | והנמשכים |  |
| 10:5.126 | אחריו |  |
| 10:5.127 | לסתור |  |
| 10:5.128 | ולקלקל |  |
| 10:5.129 | את |  |
| 10:5.130 | אמונותינו |  |
| 10:5.131 | באל |  |
| 10:5.132 | אחד |  |
| 10:5.133 | מיוחד |  |
| 10:5.134 | מכל |  |
| 10:5.135 | שאר |  |
| 10:5.136 | האחדים | , |
| 10:5.137 | ולפסוע |  |
| 10:5.138 | על |  |
| 10:5.139 | מדותיו |  |
| 10:5.140 | של |  |
| 10:5.141 | הקב"ה |  |
| 10:5.142 | ולעשותן |  |
| 10:5.143 | אלוהות |  |
| 10:5.144 | במה |  |
| 10:5.145 | שקדם | , |
| 10:5.146 | כי |  |
| 10:5.147 | לא |  |
| 10:5.148 | הספיק |  |
| 10:5.149 | לו |  |
| 10:5.150 | לחשוב |  |
| 10:5.151 | באלהות |  |
| 10:5.152 | עילה |  |
| 10:5.153 | אחת | , |
| 10:5.154 | והוצרך |  |
| 10:5.155 | להמציא |  |
| 10:5.156 | ברעיונותיו |  |
| 10:5.157 | עילות |  |
| 10:5.158 | רבות |  |
| 10:5.159 | שבה |  |
| 10:5.160 | נוטל |  |
| 10:5.161 | רשות |  |
| 10:5.162 | מזה |  |
| 10:5.163 | וזה |  |
| 10:5.164 | נעזר |  |
| 10:5.165 | במעשיהו |  |
| 10:5.166 | על |  |
| 10:5.167 | ידי |  |
| 10:5.168 | אחר | , |
| 10:5.169 | כי |  |
| 10:5.170 | חלש |  |
| 10:5.171 | הוא |  |
| 10:5.172 | ולא |  |
| 10:5.173 | יוכל |  |
| 10:5.174 | לעשות |  |
| 10:5.175 | לבדו |  |
| 10:5.176 | כמו |  |
| 10:5.177 | שנתבאר |  |
| 10:5.178 | לעיל | , |
| 10:5.179 | ובחר |  |
| 10:5.180 | לו |  |
| 10:5.181 | לעבוד |  |
| 10:5.182 | את |  |
| 10:5.183 | האליל |  |
| 10:5.184 | קצר |  |
| 10:5.185 | האפים |  |
| 10:5.186 | אשר |  |
| 10:5.187 | לפי |  |
| 10:5.188 | דעתו |  |
| 10:5.189 | לא |  |
| 10:5.190 | נשתתף |  |
| 10:5.191 | בשום |  |
| 10:5.192 | מעשה |  |
| 10:5.193 | מכל |  |
| 10:5.194 | מעשה |  |
| 10:5.195 | בראשית | , |
| 10:5.196 | כי |  |
| 10:5.197 | הוא |  |
| 10:5.198 | נברא |  |
| 10:5.199 | ועדיין |  |
| 10:5.200 | היה |  |
| 10:5.201 | חסר |  |
| 10:5.202 | תיקון |  |
| 10:5.203 | ולא |  |
| 10:5.204 | נשלם |  |
| 10:5.205 | תיקוני |  |
| 10:5.206 | עד |  |
| 10:5.207 | ימי |  |
| 10:5.208 | אברהם |  |
| 10:5.209 | אבינו |  |
| 10:5.210 | כמו |  |
| 10:5.211 | שנתבאר |  |
| 10:5.212 | לעיל |  |
| 10:5.213 | בשם |  |
| 10:5.214 | חיי |  |
| 10:5.215 | שלום |  |
| 10:5.216 | משם |  |
| 10:5.217 | האר"י |  |
| 10:5.218 | ולא |  |
| 10:5.219 | די |  |
| 10:5.220 | זה |  |
| 10:5.221 | אלא |  |
| 10:5.222 | שהוא |  |
| 10:5.223 | מבזה |  |
| 10:5.224 | ומהתל |  |
| 10:5.225 | בשלש |  |
| 10:5.226 | עשרה |  |
| 10:5.227 | מדות |  |
| 10:5.228 | שהתורה |  |
| 10:5.229 | נדרשת |  |
| 10:5.230 | בהם |  |
| 10:5.231 | שקבלם |  |
| 10:5.232 | משה |  |
| 10:5.233 | בסיני | , |
| 10:5.234 | ועל |  |
| 10:5.235 | פ | ' |
| 10:5.236 | חי"ג |  |
| 10:5.237 | מדות |  |
| 10:5.238 | הללו |  |
| 10:5.239 | נבנית |  |
| 10:5.240 | המשנה |  |
| 10:5.241 | והתלמוד |  |
| 10:5.242 | הקדוש |  |
| 10:5.243 | מימות |  |
| 10:5.244 | משה |  |
| 10:5.245 | ועד |  |
| 10:5.246 | רבי |  |
| 10:5.247 | אשי |  |
| 10:5.248 | ורבינא |  |
| 10:5.249 | מיסדי |  |
| 10:5.250 | התלמוד |  |
| 10:5.251 | כמו |  |
| 10:5.252 | שכתב |  |
| 10:5.253 | הרמב"ם |  |
| 10:5.254 | ז"ל |  |
| 10:5.255 | בהקדמת |  |
| 10:5.256 | סדר |  |
| 10:5.257 | ורעים |  |
| 10:5.258 | ומחבד |  |
| 10:5.259 | הזהר |  |
| 10:5.260 | מלעיג |  |
| 10:5.261 | עליהן |  |
| 10:5.262 | וקורא |  |
| 10:5.263 | עליהם |  |
| 10:5.264 | וימררו |  |
| 10:5.265 | את |  |
| 10:5.266 | חייהם |  |
| 10:5.267 | בעבודה |  |
| 10:5.268 | קשה |  |
| 10:5.269 | דא |  |
| 10:5.270 | קושיא | , |
| 10:5.271 | בחמר |  |
| 10:5.272 | דא |  |
| 10:5.273 | קל |  |
| 10:5.274 | וחומר |  |
| 10:5.275 | וכו | ' |
| 10:5.276 | זהר |  |
| 10:5.277 | בראשית |  |
| 10:5.278 | כ"ז | ) |
| 10:5.279 | והוציא |  |
| 10:5.280 | לעז |  |
| 10:5.281 | ודבה |  |
| 10:5.282 | רעה |  |
| 10:5.283 | על |  |
| 10:5.284 | משה |  |
| 10:5.285 | רבינו |  |
| 10:5.286 | שקבל |  |
| 10:5.287 | תורה |  |
| 10:5.288 | שבעל |  |
| 10:5.289 | פה |  |
| 10:5.290 | מפי |  |
| 10:5.291 | מטטרי"ן |  |
| 10:5.292 | ולא |  |
| 10:5.293 | מפי |  |
| 10:5.294 | הקב"ה | , |
| 10:5.295 | ולכך |  |
| 10:5.296 | היא |  |
| 10:5.297 | מעורבת |  |
| 10:5.298 | מטוב |  |
| 10:5.299 | ורע | , |
| 10:5.300 | וקורא |  |
| 10:5.301 | לה |  |
| 10:5.302 | אבן |  |
| 10:5.303 | משכית | , |
| 10:5.304 | ובה |  |
| 10:5.305 | עסקו |  |
| 10:5.306 | בן |  |
| 10:5.307 | עזאי |  |
| 10:5.308 | ובן |  |
| 10:5.309 | זומא |  |
| 10:5.310 | ואלישע |  |
| 10:5.311 | אחר |  |
| 10:5.312 | ונסתכנו | ! |
| 10:5.313 | וגם |  |
| 10:5.314 | קרא |  |
| 10:5.315 | על |  |
| 10:5.316 | המשנה |  |
| 10:5.317 | תחת |  |
| 10:5.318 | שלש |  |
| 10:5.319 | רגזה |  |
| 10:5.320 | ארץ |  |
| 10:5.321 | וכו | ', |
| 10:5.322 | ודימה |  |
| 10:5.323 | שקלא |  |
| 10:5.324 | וטריא |  |
| 10:5.325 | דרבנן |  |
| 10:5.326 | במשנה |  |
| 10:5.327 | ובתלמוד |  |
| 10:5.328 | ככלבין |  |
| 10:5.329 | דצוחין |  |
| 10:5.330 | הב |  |
| 10:5.331 | הב | , |
| 10:5.332 | וקורא |  |
| 10:5.333 | לה |  |
| 10:5.334 | גם |  |
| 10:5.335 | כן |  |
| 10:5.336 | קליפה |  |
| 10:5.337 | זהר |  |
| 10:5.338 | פרשת |  |
| 10:5.339 | תצא |  |
| 10:5.340 | ותיקוני |  |
| 10:5.341 | זהר |  |
| 10:5.342 | חדש | ) |
| 10:5.343 | ועוד |  |
| 10:5.344 | הרבה |  |
| 10:5.345 | דברים |  |
| 10:5.346 | להבזות |  |
| 10:5.347 | את |  |
| 10:5.348 | לימודי |  |
| 10:5.349 | המשנה |  |
| 10:5.350 | והתלמוד |  |
| 10:5.351 | ולהפחית |  |
| 10:5.352 | כבודם |  |
| 10:5.353 | ולזלזל |  |
| 10:5.354 | בהם | , |
| 10:5.355 | ולתת |  |
| 10:5.356 | כבוד |  |
| 10:5.357 | ויקר |  |
| 10:5.358 | ללימוד |  |
| 10:5.359 | תורתו |  |
| 10:5.360 | החדשה |  |
| 10:5.361 | שהיא |  |
| 10:5.362 | מילדי |  |
| 10:5.363 | נכרים | , |
| 10:5.364 | ולהוקיר |  |
| 10:5.365 | את |  |
| 10:5.366 | הבעלים |  |
| 10:5.367 | והעשתרות |  |
| 10:5.368 | שהמציא |  |
| 10:5.369 | ברעיונותיו |  |
| 10:5.370 | להטעות |  |
| 10:5.371 | את |  |
| 10:5.372 | ישראל | , |
| 10:5.373 | ולשבח |  |
| 10:5.374 | את |  |
| 10:5.375 | סחורתו |  |
| 10:5.376 | אשר |  |
| 10:5.377 | חשב |  |
| 10:5.378 | להנחיל |  |
| 10:5.379 | אותנו |  |
| 10:5.380 | תחת |  |
| 10:5.381 | המשנה |  |
| 10:5.382 | והתלמוד |  |
| 10:5.383 | ולרשת |  |
| 10:5.384 | את |  |
| 10:5.385 | מקומם | . |
| 10:5.386 | ועל |  |
| 10:5.387 | זה |  |
| 10:5.388 | אמרו |  |
| 10:5.389 | רז"ל |  |
| 10:5.390 | בסנהדרין |  |
| 10:5.391 | פ | ' |
| 10:5.392 | חלק |  |
| 10:5.393 | סנהדרין |  |
| 10:5.394 | דף |  |
| 10:5.395 | צ"ט | ) |
| 10:5.396 | כי |  |
| 10:5.397 | דבר |  |
| 10:5.398 | ה | ' |
| 10:5.399 | בזה | , |
| 10:5.400 | זה |  |
| 10:5.401 | האומר |  |
| 10:5.402 | אין |  |
| 10:5.403 | תורה |  |
| 10:5.404 | מן |  |
| 10:5.405 | השמים |  |
| 10:5.406 | וכו | ' |
| 10:5.407 | ואפילו |  |
| 10:5.408 | אמר |  |
| 10:5.409 | כל |  |
| 10:5.410 | התורה |  |
| 10:5.411 | כולה |  |
| 10:5.412 | מן |  |
| 10:5.413 | השמים |  |
| 10:5.414 | חוץ |  |
| 10:5.415 | מדקדוק |  |
| 10:5.416 | זה | , |
| 10:5.417 | חוץ |  |
| 10:5.418 | מקך |  |
| 10:5.419 | וחומר |  |
| 10:5.420 | זה |  |
| 10:5.421 | מגזרה |  |
| 10:5.422 | שוה |  |
| 10:5.423 | זו | , |
| 10:5.424 | זהו |  |
| 10:5.425 | כי |  |
| 10:5.426 | דבר |  |
| 10:5.427 | ה | ' |
| 10:5.428 | בזה |  |
| 10:5.429 | וכו | ' |
| 10:5.430 | רבי |  |
| 10:5.431 | נתן |  |
| 10:5.432 | אומר |  |
| 10:5.433 | זה |  |
| 10:5.434 | שאינו |  |
| 10:5.435 | משגיח |  |
| 10:5.436 | על |  |
| 10:5.437 | המשנה |  |
| 10:5.438 | וכו | '. |
| 10:5.439 | תנא |  |
| 10:5.440 | דבי |  |
| 10:5.441 | ר | ' |
| 10:5.442 | ישמעאל |  |
| 10:5.443 | כי |  |
| 10:5.444 | דבר |  |
| 10:5.445 | ה | ' |
| 10:5.446 | בזה |  |
| 10:5.447 | זה |  |
| 10:5.448 | המבזה |  |
| 10:5.449 | דיבור |  |
| 10:5.450 | שנאמר |  |
| 10:5.451 | למשה |  |
| 10:5.452 | אנכי |  |
| 10:5.453 | ולא |  |
| 10:5.454 | יהיה |  |
| 10:5.455 | לך |  |
| 10:5.456 | ע"כ | ! |
| 10:5.457 | וכל |  |
| 10:5.458 | הני |  |
| 10:5.459 | מילי |  |
| 10:5.460 | דאמור |  |
| 10:5.461 | רבנן |  |
| 10:5.462 | איתנהו |  |
| 10:5.463 | בפילסוף |  |
| 10:5.464 | מחבר |  |
| 10:5.465 | הזהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

