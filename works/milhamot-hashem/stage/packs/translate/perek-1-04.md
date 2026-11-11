# Translation pack: milhamot-hashem perek-1-04 (Milchamot ha-Shem 1:5–1:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-04`.

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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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


## Milchamot ha-Shem 1:5

Hebrew (Wikisource): וקראתי שם הספר מלחמות ה' על יחודו כי הפילסוף מחבר הספר הזהר קלקל עלינו את אמיתת ייחוד השי״ת, ונמשכו אחריו כמה רבנים וגדולים בתורה בעצימת עינים ועליהם אמר שלמה בחכמתו, כי רבים חללים הפילה ועצומים כל הרוגיה כמו שדרשו רז״ל (סנהדרין רי״ף) כי רבים חללים הפילה זה תלמיד שלא הגיע להוראה ומורה (כמו שעינינו רואות כמה מחברים דמעיילי פילא בקופא דמחטא ממיעוט בקיאותם) ועצומים כל הרוגיה זה תלמיד שהגיע להוראה ואינו מורה, (שאינו נותן לב להורות לאחרים כהלכה, ועצומים לשון ועוצם עיניו רש״י ז״ל) וכל שכן אם משדל אחרים לעסוק בספ' החיצון ס' הזהר ואינו נותן לב להבין הסתירה רק הולך בעצימת עין השכל ומכשיל אחרים באמונה כוזבת שעליו נאמר ועצומים כל הרוגיה ועובר על ולפני עור לא תתן מכשול, ושגגת תלמוד עולה זדון, והוא חוטא ומחטיא את הרבים וחטא הרבים תלוי בו.

| id | word | punctuation after |
|---|---|---|
| 1:5.1 | וקראתי |  |
| 1:5.2 | שם |  |
| 1:5.3 | הספר |  |
| 1:5.4 | מלחמות |  |
| 1:5.5 | ה | ' |
| 1:5.6 | על |  |
| 1:5.7 | יחודו |  |
| 1:5.8 | כי |  |
| 1:5.9 | הפילסוף |  |
| 1:5.10 | מחבר |  |
| 1:5.11 | הספר |  |
| 1:5.12 | הזהר |  |
| 1:5.13 | קלקל |  |
| 1:5.14 | עלינו |  |
| 1:5.15 | את |  |
| 1:5.16 | אמיתת |  |
| 1:5.17 | ייחוד |  |
| 1:5.18 | השי״ת | , |
| 1:5.19 | ונמשכו |  |
| 1:5.20 | אחריו |  |
| 1:5.21 | כמה |  |
| 1:5.22 | רבנים |  |
| 1:5.23 | וגדולים |  |
| 1:5.24 | בתורה |  |
| 1:5.25 | בעצימת |  |
| 1:5.26 | עינים |  |
| 1:5.27 | ועליהם |  |
| 1:5.28 | אמר |  |
| 1:5.29 | שלמה |  |
| 1:5.30 | בחכמתו | , |
| 1:5.31 | כי |  |
| 1:5.32 | רבים |  |
| 1:5.33 | חללים |  |
| 1:5.34 | הפילה |  |
| 1:5.35 | ועצומים |  |
| 1:5.36 | כל |  |
| 1:5.37 | הרוגיה |  |
| 1:5.38 | כמו |  |
| 1:5.39 | שדרשו |  |
| 1:5.40 | רז״ל |  |
| 1:5.41 | סנהדרין |  |
| 1:5.42 | רי״ף | ) |
| 1:5.43 | כי |  |
| 1:5.44 | רבים |  |
| 1:5.45 | חללים |  |
| 1:5.46 | הפילה |  |
| 1:5.47 | זה |  |
| 1:5.48 | תלמיד |  |
| 1:5.49 | שלא |  |
| 1:5.50 | הגיע |  |
| 1:5.51 | להוראה |  |
| 1:5.52 | ומורה |  |
| 1:5.53 | כמו |  |
| 1:5.54 | שעינינו |  |
| 1:5.55 | רואות |  |
| 1:5.56 | כמה |  |
| 1:5.57 | מחברים |  |
| 1:5.58 | דמעיילי |  |
| 1:5.59 | פילא |  |
| 1:5.60 | בקופא |  |
| 1:5.61 | דמחטא |  |
| 1:5.62 | ממיעוט |  |
| 1:5.63 | בקיאותם | ) |
| 1:5.64 | ועצומים |  |
| 1:5.65 | כל |  |
| 1:5.66 | הרוגיה |  |
| 1:5.67 | זה |  |
| 1:5.68 | תלמיד |  |
| 1:5.69 | שהגיע |  |
| 1:5.70 | להוראה |  |
| 1:5.71 | ואינו |  |
| 1:5.72 | מורה | , |
| 1:5.73 | שאינו |  |
| 1:5.74 | נותן |  |
| 1:5.75 | לב |  |
| 1:5.76 | להורות |  |
| 1:5.77 | לאחרים |  |
| 1:5.78 | כהלכה | , |
| 1:5.79 | ועצומים |  |
| 1:5.80 | לשון |  |
| 1:5.81 | ועוצם |  |
| 1:5.82 | עיניו |  |
| 1:5.83 | רש״י |  |
| 1:5.84 | ז״ל | ) |
| 1:5.85 | וכל |  |
| 1:5.86 | שכן |  |
| 1:5.87 | אם |  |
| 1:5.88 | משדל |  |
| 1:5.89 | אחרים |  |
| 1:5.90 | לעסוק |  |
| 1:5.91 | בספ | ' |
| 1:5.92 | החיצון |  |
| 1:5.93 | ס | ' |
| 1:5.94 | הזהר |  |
| 1:5.95 | ואינו |  |
| 1:5.96 | נותן |  |
| 1:5.97 | לב |  |
| 1:5.98 | להבין |  |
| 1:5.99 | הסתירה |  |
| 1:5.100 | רק |  |
| 1:5.101 | הולך |  |
| 1:5.102 | בעצימת |  |
| 1:5.103 | עין |  |
| 1:5.104 | השכל |  |
| 1:5.105 | ומכשיל |  |
| 1:5.106 | אחרים |  |
| 1:5.107 | באמונה |  |
| 1:5.108 | כוזבת |  |
| 1:5.109 | שעליו |  |
| 1:5.110 | נאמר |  |
| 1:5.111 | ועצומים |  |
| 1:5.112 | כל |  |
| 1:5.113 | הרוגיה |  |
| 1:5.114 | ועובר |  |
| 1:5.115 | על |  |
| 1:5.116 | ולפני |  |
| 1:5.117 | עור |  |
| 1:5.118 | לא |  |
| 1:5.119 | תתן |  |
| 1:5.120 | מכשול | , |
| 1:5.121 | ושגגת |  |
| 1:5.122 | תלמוד |  |
| 1:5.123 | עולה |  |
| 1:5.124 | זדון | , |
| 1:5.125 | והוא |  |
| 1:5.126 | חוטא |  |
| 1:5.127 | ומחטיא |  |
| 1:5.128 | את |  |
| 1:5.129 | הרבים |  |
| 1:5.130 | וחטא |  |
| 1:5.131 | הרבים |  |
| 1:5.132 | תלוי |  |
| 1:5.133 | בו | . |


## Milchamot ha-Shem 1:6

Hebrew (Wikisource): ועל אמונה החיצונה הלזו של הפילסוף מחבר הזהר דרשו רז״ל (ילקוט פרשת קדושים ומדרש הגדול שם ורמב״ם בס' המצות לאוין י') אל תפנו אל האלילים ואלהי אם אתה פונה אחרים (לכבדם ספורנו) סופך לעשותם אלוהות ע״כ, וכך אירע באמונת הפילסוף מחבר הזהר כי הלומדים הראשונים אשר למדו בו בחשבם שהוא לרשב״י התנא, חשבו לישב דבריו ולקרות את הפרצופים כלים, והחשיבום וכיבדו אותם והבאים אחריהם עשו אותם אלוהות כמחשבת הפילסוף המחבר וקיבלום עליהם באלוהות כמו שיתבאר בפנים וכן דרשו רז״ל בפסוק יין ושכר אל תשת וז״ל ר' חייא פתח פקודי ה' ישרים וכתיב ביין יתהלך במישרים וכתיב משמחי לב וכתיב ויין ישמח לבב אנוש מה היין הזה משמח את הלב כך דברי תורה משמחין את הלב מה היין הזה השותה ממנו יותר מצרכו לסוף מקיא את הכל כך דברי תורה העוסק בהן ומחשב בהן יותר מדאי סוף שהוא יוצא למינות ע״כ (מדרש הגדול פרשת שמיני פסוק יין ושכר) כדבר הזה עשה הפילסוף נביא השקר מחבר הזהר לתור אחרי לבבו לחשב בדברי תורה יותר מדאי בדרכים הרחוקים מן האמת, ורחוקים מאד מן המדות שקבלו חכמים איש מפי איש שהתורה נדרשת בהן והן הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי דאגדתא כמו שיראה המעיין הנבון והסית והדיח לחכמים ישרי לב ההולכים לתומם, להאמינו לשקריו שנגלה לו הקב״ה הנקרא אצלו סבא דסבין ומשה רבינו ‏ע״ה ואליהו וגילו לו סתרים של מינות שאסור לחשב בהן אפי' בבית הכסא והבטיח לעוסק בתורתו בלי מבין לשכר רב יותר מגופי תורה כדי להעור עיני חכמים ולסתום דרך העיון בדבריו שלא להבין את נכליו ולהגדיל מעלת העוסקים בתורתו החדשה המלאה אלילים שונים בעצמותם ופעוליהם מבלי להקפיד על תוספת ורבוי באלהינו המיוחד מכל שאר האחדים ולא יתנו לב איך הקפידה תורה על בל תוסיף בין בתפילין בין במזוזה בין בארבעה מינין שבלולב וכן בשאר מצות אע״פ שנותן חמש הפרשיות בעור אחד והוא מחברן להיות אחד, וכן אם מוסיף במזוזה פרשה אחת אע״פ שקלף מחברן לחיות אחד לא אמרינן כלא חד וכן בציצית ובאגודת ארבעה מינין אם הוסיף מין חמישי באגודה פסל ולא אמרינן כלא חד.

| id | word | punctuation after |
|---|---|---|
| 1:6.1 | ועל |  |
| 1:6.2 | אמונה |  |
| 1:6.3 | החיצונה |  |
| 1:6.4 | הלזו |  |
| 1:6.5 | של |  |
| 1:6.6 | הפילסוף |  |
| 1:6.7 | מחבר |  |
| 1:6.8 | הזהר |  |
| 1:6.9 | דרשו |  |
| 1:6.10 | רז״ל |  |
| 1:6.11 | ילקוט |  |
| 1:6.12 | פרשת |  |
| 1:6.13 | קדושים |  |
| 1:6.14 | ומדרש |  |
| 1:6.15 | הגדול |  |
| 1:6.16 | שם |  |
| 1:6.17 | ורמב״ם |  |
| 1:6.18 | בס | ' |
| 1:6.19 | המצות |  |
| 1:6.20 | לאוין |  |
| 1:6.21 | י | ') |
| 1:6.22 | אל |  |
| 1:6.23 | תפנו |  |
| 1:6.24 | אל |  |
| 1:6.25 | האלילים |  |
| 1:6.26 | ואלהי |  |
| 1:6.27 | אם |  |
| 1:6.28 | אתה |  |
| 1:6.29 | פונה |  |
| 1:6.30 | אחרים |  |
| 1:6.31 | לכבדם |  |
| 1:6.32 | ספורנו | ) |
| 1:6.33 | סופך |  |
| 1:6.34 | לעשותם |  |
| 1:6.35 | אלוהות |  |
| 1:6.36 | ע״כ | , |
| 1:6.37 | וכך |  |
| 1:6.38 | אירע |  |
| 1:6.39 | באמונת |  |
| 1:6.40 | הפילסוף |  |
| 1:6.41 | מחבר |  |
| 1:6.42 | הזהר |  |
| 1:6.43 | כי |  |
| 1:6.44 | הלומדים |  |
| 1:6.45 | הראשונים |  |
| 1:6.46 | אשר |  |
| 1:6.47 | למדו |  |
| 1:6.48 | בו |  |
| 1:6.49 | בחשבם |  |
| 1:6.50 | שהוא |  |
| 1:6.51 | לרשב״י |  |
| 1:6.52 | התנא | , |
| 1:6.53 | חשבו |  |
| 1:6.54 | לישב |  |
| 1:6.55 | דבריו |  |
| 1:6.56 | ולקרות |  |
| 1:6.57 | את |  |
| 1:6.58 | הפרצופים |  |
| 1:6.59 | כלים | , |
| 1:6.60 | והחשיבום |  |
| 1:6.61 | וכיבדו |  |
| 1:6.62 | אותם |  |
| 1:6.63 | והבאים |  |
| 1:6.64 | אחריהם |  |
| 1:6.65 | עשו |  |
| 1:6.66 | אותם |  |
| 1:6.67 | אלוהות |  |
| 1:6.68 | כמחשבת |  |
| 1:6.69 | הפילסוף |  |
| 1:6.70 | המחבר |  |
| 1:6.71 | וקיבלום |  |
| 1:6.72 | עליהם |  |
| 1:6.73 | באלוהות |  |
| 1:6.74 | כמו |  |
| 1:6.75 | שיתבאר |  |
| 1:6.76 | בפנים |  |
| 1:6.77 | וכן |  |
| 1:6.78 | דרשו |  |
| 1:6.79 | רז״ל |  |
| 1:6.80 | בפסוק |  |
| 1:6.81 | יין |  |
| 1:6.82 | ושכר |  |
| 1:6.83 | אל |  |
| 1:6.84 | תשת |  |
| 1:6.85 | וז״ל |  |
| 1:6.86 | ר | ' |
| 1:6.87 | חייא |  |
| 1:6.88 | פתח |  |
| 1:6.89 | פקודי |  |
| 1:6.90 | ה | ' |
| 1:6.91 | ישרים |  |
| 1:6.92 | וכתיב |  |
| 1:6.93 | ביין |  |
| 1:6.94 | יתהלך |  |
| 1:6.95 | במישרים |  |
| 1:6.96 | וכתיב |  |
| 1:6.97 | משמחי |  |
| 1:6.98 | לב |  |
| 1:6.99 | וכתיב |  |
| 1:6.100 | ויין |  |
| 1:6.101 | ישמח |  |
| 1:6.102 | לבב |  |
| 1:6.103 | אנוש |  |
| 1:6.104 | מה |  |
| 1:6.105 | היין |  |
| 1:6.106 | הזה |  |
| 1:6.107 | משמח |  |
| 1:6.108 | את |  |
| 1:6.109 | הלב |  |
| 1:6.110 | כך |  |
| 1:6.111 | דברי |  |
| 1:6.112 | תורה |  |
| 1:6.113 | משמחין |  |
| 1:6.114 | את |  |
| 1:6.115 | הלב |  |
| 1:6.116 | מה |  |
| 1:6.117 | היין |  |
| 1:6.118 | הזה |  |
| 1:6.119 | השותה |  |
| 1:6.120 | ממנו |  |
| 1:6.121 | יותר |  |
| 1:6.122 | מצרכו |  |
| 1:6.123 | לסוף |  |
| 1:6.124 | מקיא |  |
| 1:6.125 | את |  |
| 1:6.126 | הכל |  |
| 1:6.127 | כך |  |
| 1:6.128 | דברי |  |
| 1:6.129 | תורה |  |
| 1:6.130 | העוסק |  |
| 1:6.131 | בהן |  |
| 1:6.132 | ומחשב |  |
| 1:6.133 | בהן |  |
| 1:6.134 | יותר |  |
| 1:6.135 | מדאי |  |
| 1:6.136 | סוף |  |
| 1:6.137 | שהוא |  |
| 1:6.138 | יוצא |  |
| 1:6.139 | למינות |  |
| 1:6.140 | ע״כ |  |
| 1:6.141 | מדרש |  |
| 1:6.142 | הגדול |  |
| 1:6.143 | פרשת |  |
| 1:6.144 | שמיני |  |
| 1:6.145 | פסוק |  |
| 1:6.146 | יין |  |
| 1:6.147 | ושכר | ) |
| 1:6.148 | כדבר |  |
| 1:6.149 | הזה |  |
| 1:6.150 | עשה |  |
| 1:6.151 | הפילסוף |  |
| 1:6.152 | נביא |  |
| 1:6.153 | השקר |  |
| 1:6.154 | מחבר |  |
| 1:6.155 | הזהר |  |
| 1:6.156 | לתור |  |
| 1:6.157 | אחרי |  |
| 1:6.158 | לבבו |  |
| 1:6.159 | לחשב |  |
| 1:6.160 | בדברי |  |
| 1:6.161 | תורה |  |
| 1:6.162 | יותר |  |
| 1:6.163 | מדאי |  |
| 1:6.164 | בדרכים |  |
| 1:6.165 | הרחוקים |  |
| 1:6.166 | מן |  |
| 1:6.167 | האמת | , |
| 1:6.168 | ורחוקים |  |
| 1:6.169 | מאד |  |
| 1:6.170 | מן |  |
| 1:6.171 | המדות |  |
| 1:6.172 | שקבלו |  |
| 1:6.173 | חכמים |  |
| 1:6.174 | איש |  |
| 1:6.175 | מפי |  |
| 1:6.176 | איש |  |
| 1:6.177 | שהתורה |  |
| 1:6.178 | נדרשת |  |
| 1:6.179 | בהן |  |
| 1:6.180 | והן |  |
| 1:6.181 | הי״ג |  |
| 1:6.182 | מדות |  |
| 1:6.183 | דר | ' |
| 1:6.184 | ישמעאל |  |
| 1:6.185 | ול״ג |  |
| 1:6.186 | מדות |  |
| 1:6.187 | דר | ' |
| 1:6.188 | יוסי |  |
| 1:6.189 | הגלילי |  |
| 1:6.190 | דאגדתא |  |
| 1:6.191 | כמו |  |
| 1:6.192 | שיראה |  |
| 1:6.193 | המעיין |  |
| 1:6.194 | הנבון |  |
| 1:6.195 | והסית |  |
| 1:6.196 | והדיח |  |
| 1:6.197 | לחכמים |  |
| 1:6.198 | ישרי |  |
| 1:6.199 | לב |  |
| 1:6.200 | ההולכים |  |
| 1:6.201 | לתומם | , |
| 1:6.202 | להאמינו |  |
| 1:6.203 | לשקריו |  |
| 1:6.204 | שנגלה |  |
| 1:6.205 | לו |  |
| 1:6.206 | הקב״ה |  |
| 1:6.207 | הנקרא |  |
| 1:6.208 | אצלו |  |
| 1:6.209 | סבא |  |
| 1:6.210 | דסבין |  |
| 1:6.211 | ומשה |  |
| 1:6.212 | רבינו |  |
| 1:6.213 | ע״ה |  |
| 1:6.214 | ואליהו |  |
| 1:6.215 | וגילו |  |
| 1:6.216 | לו |  |
| 1:6.217 | סתרים |  |
| 1:6.218 | של |  |
| 1:6.219 | מינות |  |
| 1:6.220 | שאסור |  |
| 1:6.221 | לחשב |  |
| 1:6.222 | בהן |  |
| 1:6.223 | אפי | ' |
| 1:6.224 | בבית |  |
| 1:6.225 | הכסא |  |
| 1:6.226 | והבטיח |  |
| 1:6.227 | לעוסק |  |
| 1:6.228 | בתורתו |  |
| 1:6.229 | בלי |  |
| 1:6.230 | מבין |  |
| 1:6.231 | לשכר |  |
| 1:6.232 | רב |  |
| 1:6.233 | יותר |  |
| 1:6.234 | מגופי |  |
| 1:6.235 | תורה |  |
| 1:6.236 | כדי |  |
| 1:6.237 | להעור |  |
| 1:6.238 | עיני |  |
| 1:6.239 | חכמים |  |
| 1:6.240 | ולסתום |  |
| 1:6.241 | דרך |  |
| 1:6.242 | העיון |  |
| 1:6.243 | בדבריו |  |
| 1:6.244 | שלא |  |
| 1:6.245 | להבין |  |
| 1:6.246 | את |  |
| 1:6.247 | נכליו |  |
| 1:6.248 | ולהגדיל |  |
| 1:6.249 | מעלת |  |
| 1:6.250 | העוסקים |  |
| 1:6.251 | בתורתו |  |
| 1:6.252 | החדשה |  |
| 1:6.253 | המלאה |  |
| 1:6.254 | אלילים |  |
| 1:6.255 | שונים |  |
| 1:6.256 | בעצמותם |  |
| 1:6.257 | ופעוליהם |  |
| 1:6.258 | מבלי |  |
| 1:6.259 | להקפיד |  |
| 1:6.260 | על |  |
| 1:6.261 | תוספת |  |
| 1:6.262 | ורבוי |  |
| 1:6.263 | באלהינו |  |
| 1:6.264 | המיוחד |  |
| 1:6.265 | מכל |  |
| 1:6.266 | שאר |  |
| 1:6.267 | האחדים |  |
| 1:6.268 | ולא |  |
| 1:6.269 | יתנו |  |
| 1:6.270 | לב |  |
| 1:6.271 | איך |  |
| 1:6.272 | הקפידה |  |
| 1:6.273 | תורה |  |
| 1:6.274 | על |  |
| 1:6.275 | בל |  |
| 1:6.276 | תוסיף |  |
| 1:6.277 | בין |  |
| 1:6.278 | בתפילין |  |
| 1:6.279 | בין |  |
| 1:6.280 | במזוזה |  |
| 1:6.281 | בין |  |
| 1:6.282 | בארבעה |  |
| 1:6.283 | מינין |  |
| 1:6.284 | שבלולב |  |
| 1:6.285 | וכן |  |
| 1:6.286 | בשאר |  |
| 1:6.287 | מצות |  |
| 1:6.288 | אע״פ |  |
| 1:6.289 | שנותן |  |
| 1:6.290 | חמש |  |
| 1:6.291 | הפרשיות |  |
| 1:6.292 | בעור |  |
| 1:6.293 | אחד |  |
| 1:6.294 | והוא |  |
| 1:6.295 | מחברן |  |
| 1:6.296 | להיות |  |
| 1:6.297 | אחד | , |
| 1:6.298 | וכן |  |
| 1:6.299 | אם |  |
| 1:6.300 | מוסיף |  |
| 1:6.301 | במזוזה |  |
| 1:6.302 | פרשה |  |
| 1:6.303 | אחת |  |
| 1:6.304 | אע״פ |  |
| 1:6.305 | שקלף |  |
| 1:6.306 | מחברן |  |
| 1:6.307 | לחיות |  |
| 1:6.308 | אחד |  |
| 1:6.309 | לא |  |
| 1:6.310 | אמרינן |  |
| 1:6.311 | כלא |  |
| 1:6.312 | חד |  |
| 1:6.313 | וכן |  |
| 1:6.314 | בציצית |  |
| 1:6.315 | ובאגודת |  |
| 1:6.316 | ארבעה |  |
| 1:6.317 | מינין |  |
| 1:6.318 | אם |  |
| 1:6.319 | הוסיף |  |
| 1:6.320 | מין |  |
| 1:6.321 | חמישי |  |
| 1:6.322 | באגודה |  |
| 1:6.323 | פסל |  |
| 1:6.324 | ולא |  |
| 1:6.325 | אמרינן |  |
| 1:6.326 | כלא |  |
| 1:6.327 | חד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

