# Translation pack: milhamot-hashem perek-6-09 (Milchamot ha-Shem 6:14–6:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-09`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- והוא שלא: "provided that" (fixed).
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
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


## Milchamot ha-Shem 6:14

Hebrew (Wikisource): ובסוף פרק קמא דחגיגה י' ע"א שנינו היתר נדרים פורחים באויר ואין להם על מה שיסמכו, הלכות שבת חגיגות והמעילות הרי הם כהררים התלויים בשערה מקרא מועט והלכות מרובות, הדינין והעבודות, הטהורות והטומאות ועריות יש להן על מה שיסמכו הן הן גופי תורה. ופי' בגמרא דהכי קאמר הן והן גופי תורה בין אלו שיש להן על מה שיסמכו בין אלו שהם כהררים התלוים בשערה אלו ואלו גופי תורה, כלומר עיקרה ועצמה של תורה. גם בסוף פרק שלישי דאבות שנינו גם כן ר' אלעזר חסמא אומר קינין ופתחי נדה הן הן גופי הלכות, ופי' ר' עובדיה הן הן גופי הלכות עיקר תורה שבעל פה, שמקבלים עליה שכר ע"כ. ובמסכת מגילה דף כ"ח ע"ב אמרו ודאשתמש בתגא חלף, תני ריש לקיש זה המשתמש במי ששונה הלכות כתרה של תורה. ע"כ מכל מה שהבאתי יתבאר לך שההלכות שבמשנה ובגמרא הן הן עיקרה ועצמה של תורה שבעל פה שקבל משה בסיני שאדם מקבל שכר עליהם וזוכה בהן לחיי העולם הבא, והן כתרה של תורה שנאמר בה בי מלכים ימלכו בי שרים ישורו. ובמדרש רבה פרשת קרח דרשו רז"ל ששים המה מלכות אלו ששים מסכתות ושמנים פלגשים אלו שמנים בתי מדרשות שהיו בירושלים כנגד פתחיה, ועלמות אין מספר זו משנה החיצונה, ופירשו המתנות כהונה, והרב עץ חיים, דקרי לבתי מדרשות פילגשים על שם שנושאים ונותנים וחולקים זה עם זה כמו פילגשם, וכי משנה החיצונה מה שלמדו היחידים חוץ לאלו בתי המדרשות. וביד יוסף פי' שכוונת המדרש על הברייתות וע"כ. וכן מצאתי סיוע לפירושו בשיר השירים רבה וז"ל ששים המה מלכות אלו ששים מסכתות של הלכות, ושמונים פילגשים אלו שמונים פרשיות שבתורת כהנים. ועלמות אין מספר אין קץ לתוספתות. אחת היא הן חולקין אלו עם אלו וכולהון דורשים מטעם אחד, מהלכה אחת, מגזירה שוה, מקל וחומר ע"כ. ראה כמה מחבבים רז"ל את המשנה והתלמוד וקראו להן מלכות. ולא שפחות כאשר עשה הפילסוף הרמאי האדוק בדעות חיצוניות הוא מחבר הזהר כמו שיתבאר לקמן. ובבא מציעא פרק אלו מציאות (בבא מציעא דף ל"ג) גרסינן. ת"ר העוסקים במקרא מדה ואינה מדה במשנה מדה ונוטל עליה שכר, בתלמוד, אין לך מדה גדולה מזו, ולעולם הוי רץ למשנה יותר מן התלמוד, ופריך היא גופה קשיא, אמרת בתלמוד אין לך מדה גדולה מזו, והדר אמרת ולעולם הוי רץ למשנה יותר מן התלמוד? אמר ר' יוחנן בימי ר' נשנית משנה זו, שבקו כולי עלמא מתניתין ואזלו בתר תלמודא, הדר דרש להו ולעולם הוי רץ למשנה יותר מן התלמוד עיין שם ובפי' רש"י. ובירושלמי מבואר בהדיא דרשב"י הוא דתני לה, כדגרסינן התם בפרק כל כתבי הדא אמרא שהמשנה קודמת למקרא, ודא מסייעא לרשב"י דתני רשב"י העוסק במקרא מדה שאינה מדה, במשנה מדה שנוטלין ממנה שכר, העוסק בתלמוד אין לך מדה גדולה מזו, לעולם הוי רץ אחר המשנה יותר מן התלמוד אמר ר' יוסי בר' בון הדא דאת אמר עד שלא שקע בו רבי רוב משניות, אבל מששיקע בו רבי רוב משניות לעולם הוי רץ אחר התלמוד יותר מן המשנה ע"כ. וכן בפ"ק דשבת סוף פסקא דלא ישב אדם לפני הספר קאמר התם דהקורא את שמע בעונתה כמשנה היא, ושלא בעונתה כאדם שקורא בתורה ששכרו פחות מהעוסק במשנה, ומסיים בה רשב"י כדעתיה דרשב"י אמר העוסק במקרא מדה שאינה מדה, ורבנין עבדין מקרא כמשנה עכ"ל, ובילקוט ריש פר' צו ר' חנינה אמר אין הגליות מתכנסות אלא בזכות משניות שנ' גם כי יתנו בגוים עתה אקבצם ע"כ. ראה תראה ידידי הקורא דברי רז"ל שהבאתי מקצת מהם כמה הם מחבבים ומחשיבים את המשנה שהיא עיקר תורה שבעל פה שנאמרה למשה בסיני, ובפרט שים עיניך לדברי תנא דידן רשב"י הנאמרים באמת בבבלי ובירושלמי כמה הוא מחבב ומחשיב את המשנה יותר מן המקרא, ואת התלמוד שהוא פי' טעמי המשנה יותר מכולם שאין לך מדה גדולה יותר ממנה, אלא דבירושלמי מוקי להא דאמר לעולם הוי רץ למשנה יותר מן התלמוד קודם שסידר רבי את המשנה שאז היו הרבה משניות שהם שגורים בפי התלמידים כל אחת בשם אומרה, והרבה בריתות היו להם, אבל לאחר שסידר רבי את המשנה והשמיט בה כמה משניות ובריתות, וגם סתם במשנה דברי יחידים שראה ר' דבריהם ושנאן סתם כדי לקבוע הלכה כמותם, לפיכך חזרו לומר הוי רץ לתלמוד יותר מן המשנה. ומדברי כולם למדנו כמה היה חביב להם לימוד המשנה והתלמוד, ושרשב"י הוא דתני לה, ומזה הוכחה גמורה ברורה שס' הזהר אינו לרשב"י התנא ע"ה, כי אם לפילסוף רמאי גדול שלמד כל חכמת ישראל בשני התלמודים ומדרשי רז"ל. והוא מאמין בדעות הכשדים והכלדיים פילסופים הראשונים שהיו חולקים על אברהם אבינו באמונתו באל אחד, ומאמינים בשתי סבות קודמות, וכל אחת פועלת היפך האחרת, והסבות הפועלות הטוב קורא להם בשם סטרא דקדושה, והם אדם קדמון וכתר עליון עם שאר הפרצופים וצורות המשתלשלים מהם, והסבות הפועלות הרע הוא קורא להם סטרא דקליפה, והם אדם בליעל וכתר עליון דקליפה עם שאר פרצופים וצורות דעשר ספירות דקליפה המשתלשלים מהם הנקראים בשם אל אחר, כי הוא מאמין ומאמת מציאות אלהים אחרים בעולם כדעת הכשדים והמצרים הקדמונים בפילוסופיא הקדומה שכבר נמר ריחה וסר טעמה. ובשפתי חלקות וגיזומים הוא מושך לב רבים מהרבנים הפתאים להאמין באלהים אחרים רבים ולזלזל בלמוד המשנה והתלמוד עד שלא ידעו מה לעשות והיאך יפרשו מן האיסורי' שנאמרו למשה בסיני כגון ח"י טריפות הנזכרים במשנה וחילוקי דיניהם, לפי שלא נתפרש בתורה בהדיא אלא דרוסה כמו שכ' הרמב"ם ז"ל בהל' מאכלות אסורות. ויכוצא בזה בשאר איסורים של תורה שלא נתפרשו בתורה שבכתב ונמסרו למשה רבי' ע"ה בעל פה. והמה מפורשים במשנה ובתלמוד כגון דיני תערובת איסור ובשר בחלב. ואיזה מהן אסור מדאוריתא ואיזה אסור מדרבנן. והפילסוף הנז' כיוון לפסול ולהשכיח מישראל כל הדינים המבוארים במשנה ובתלמוד. ובהדיא אמר שבסוף גלותא יתפרנסון מחבורא דיליה.

| id | word | punctuation after |
|---|---|---|
| 6:14.1 | ובסוף |  |
| 6:14.2 | פרק |  |
| 6:14.3 | קמא |  |
| 6:14.4 | דחגיגה |  |
| 6:14.5 | י | ' |
| 6:14.6 | ע"א |  |
| 6:14.7 | שנינו |  |
| 6:14.8 | היתר |  |
| 6:14.9 | נדרים |  |
| 6:14.10 | פורחים |  |
| 6:14.11 | באויר |  |
| 6:14.12 | ואין |  |
| 6:14.13 | להם |  |
| 6:14.14 | על |  |
| 6:14.15 | מה |  |
| 6:14.16 | שיסמכו | , |
| 6:14.17 | הלכות |  |
| 6:14.18 | שבת |  |
| 6:14.19 | חגיגות |  |
| 6:14.20 | והמעילות |  |
| 6:14.21 | הרי |  |
| 6:14.22 | הם |  |
| 6:14.23 | כהררים |  |
| 6:14.24 | התלויים |  |
| 6:14.25 | בשערה |  |
| 6:14.26 | מקרא |  |
| 6:14.27 | מועט |  |
| 6:14.28 | והלכות |  |
| 6:14.29 | מרובות | , |
| 6:14.30 | הדינין |  |
| 6:14.31 | והעבודות | , |
| 6:14.32 | הטהורות |  |
| 6:14.33 | והטומאות |  |
| 6:14.34 | ועריות |  |
| 6:14.35 | יש |  |
| 6:14.36 | להן |  |
| 6:14.37 | על |  |
| 6:14.38 | מה |  |
| 6:14.39 | שיסמכו |  |
| 6:14.40 | הן |  |
| 6:14.41 | הן |  |
| 6:14.42 | גופי |  |
| 6:14.43 | תורה | . |
| 6:14.44 | ופי | ' |
| 6:14.45 | בגמרא |  |
| 6:14.46 | דהכי |  |
| 6:14.47 | קאמר |  |
| 6:14.48 | הן |  |
| 6:14.49 | והן |  |
| 6:14.50 | גופי |  |
| 6:14.51 | תורה |  |
| 6:14.52 | בין |  |
| 6:14.53 | אלו |  |
| 6:14.54 | שיש |  |
| 6:14.55 | להן |  |
| 6:14.56 | על |  |
| 6:14.57 | מה |  |
| 6:14.58 | שיסמכו |  |
| 6:14.59 | בין |  |
| 6:14.60 | אלו |  |
| 6:14.61 | שהם |  |
| 6:14.62 | כהררים |  |
| 6:14.63 | התלוים |  |
| 6:14.64 | בשערה |  |
| 6:14.65 | אלו |  |
| 6:14.66 | ואלו |  |
| 6:14.67 | גופי |  |
| 6:14.68 | תורה | , |
| 6:14.69 | כלומר |  |
| 6:14.70 | עיקרה |  |
| 6:14.71 | ועצמה |  |
| 6:14.72 | של |  |
| 6:14.73 | תורה | . |
| 6:14.74 | גם |  |
| 6:14.75 | בסוף |  |
| 6:14.76 | פרק |  |
| 6:14.77 | שלישי |  |
| 6:14.78 | דאבות |  |
| 6:14.79 | שנינו |  |
| 6:14.80 | גם |  |
| 6:14.81 | כן |  |
| 6:14.82 | ר | ' |
| 6:14.83 | אלעזר |  |
| 6:14.84 | חסמא |  |
| 6:14.85 | אומר |  |
| 6:14.86 | קינין |  |
| 6:14.87 | ופתחי |  |
| 6:14.88 | נדה |  |
| 6:14.89 | הן |  |
| 6:14.90 | הן |  |
| 6:14.91 | גופי |  |
| 6:14.92 | הלכות | , |
| 6:14.93 | ופי | ' |
| 6:14.94 | ר | ' |
| 6:14.95 | עובדיה |  |
| 6:14.96 | הן |  |
| 6:14.97 | הן |  |
| 6:14.98 | גופי |  |
| 6:14.99 | הלכות |  |
| 6:14.100 | עיקר |  |
| 6:14.101 | תורה |  |
| 6:14.102 | שבעל |  |
| 6:14.103 | פה | , |
| 6:14.104 | שמקבלים |  |
| 6:14.105 | עליה |  |
| 6:14.106 | שכר |  |
| 6:14.107 | ע"כ | . |
| 6:14.108 | ובמסכת |  |
| 6:14.109 | מגילה |  |
| 6:14.110 | דף |  |
| 6:14.111 | כ"ח |  |
| 6:14.112 | ע"ב |  |
| 6:14.113 | אמרו |  |
| 6:14.114 | ודאשתמש |  |
| 6:14.115 | בתגא |  |
| 6:14.116 | חלף | , |
| 6:14.117 | תני |  |
| 6:14.118 | ריש |  |
| 6:14.119 | לקיש |  |
| 6:14.120 | זה |  |
| 6:14.121 | המשתמש |  |
| 6:14.122 | במי |  |
| 6:14.123 | ששונה |  |
| 6:14.124 | הלכות |  |
| 6:14.125 | כתרה |  |
| 6:14.126 | של |  |
| 6:14.127 | תורה | . |
| 6:14.128 | ע"כ |  |
| 6:14.129 | מכל |  |
| 6:14.130 | מה |  |
| 6:14.131 | שהבאתי |  |
| 6:14.132 | יתבאר |  |
| 6:14.133 | לך |  |
| 6:14.134 | שההלכות |  |
| 6:14.135 | שבמשנה |  |
| 6:14.136 | ובגמרא |  |
| 6:14.137 | הן |  |
| 6:14.138 | הן |  |
| 6:14.139 | עיקרה |  |
| 6:14.140 | ועצמה |  |
| 6:14.141 | של |  |
| 6:14.142 | תורה |  |
| 6:14.143 | שבעל |  |
| 6:14.144 | פה |  |
| 6:14.145 | שקבל |  |
| 6:14.146 | משה |  |
| 6:14.147 | בסיני |  |
| 6:14.148 | שאדם |  |
| 6:14.149 | מקבל |  |
| 6:14.150 | שכר |  |
| 6:14.151 | עליהם |  |
| 6:14.152 | וזוכה |  |
| 6:14.153 | בהן |  |
| 6:14.154 | לחיי |  |
| 6:14.155 | העולם |  |
| 6:14.156 | הבא | , |
| 6:14.157 | והן |  |
| 6:14.158 | כתרה |  |
| 6:14.159 | של |  |
| 6:14.160 | תורה |  |
| 6:14.161 | שנאמר |  |
| 6:14.162 | בה |  |
| 6:14.163 | בי |  |
| 6:14.164 | מלכים |  |
| 6:14.165 | ימלכו |  |
| 6:14.166 | בי |  |
| 6:14.167 | שרים |  |
| 6:14.168 | ישורו | . |
| 6:14.169 | ובמדרש |  |
| 6:14.170 | רבה |  |
| 6:14.171 | פרשת |  |
| 6:14.172 | קרח |  |
| 6:14.173 | דרשו |  |
| 6:14.174 | רז"ל |  |
| 6:14.175 | ששים |  |
| 6:14.176 | המה |  |
| 6:14.177 | מלכות |  |
| 6:14.178 | אלו |  |
| 6:14.179 | ששים |  |
| 6:14.180 | מסכתות |  |
| 6:14.181 | ושמנים |  |
| 6:14.182 | פלגשים |  |
| 6:14.183 | אלו |  |
| 6:14.184 | שמנים |  |
| 6:14.185 | בתי |  |
| 6:14.186 | מדרשות |  |
| 6:14.187 | שהיו |  |
| 6:14.188 | בירושלים |  |
| 6:14.189 | כנגד |  |
| 6:14.190 | פתחיה | , |
| 6:14.191 | ועלמות |  |
| 6:14.192 | אין |  |
| 6:14.193 | מספר |  |
| 6:14.194 | זו |  |
| 6:14.195 | משנה |  |
| 6:14.196 | החיצונה | , |
| 6:14.197 | ופירשו |  |
| 6:14.198 | המתנות |  |
| 6:14.199 | כהונה | , |
| 6:14.200 | והרב |  |
| 6:14.201 | עץ |  |
| 6:14.202 | חיים | , |
| 6:14.203 | דקרי |  |
| 6:14.204 | לבתי |  |
| 6:14.205 | מדרשות |  |
| 6:14.206 | פילגשים |  |
| 6:14.207 | על |  |
| 6:14.208 | שם |  |
| 6:14.209 | שנושאים |  |
| 6:14.210 | ונותנים |  |
| 6:14.211 | וחולקים |  |
| 6:14.212 | זה |  |
| 6:14.213 | עם |  |
| 6:14.214 | זה |  |
| 6:14.215 | כמו |  |
| 6:14.216 | פילגשם | , |
| 6:14.217 | וכי |  |
| 6:14.218 | משנה |  |
| 6:14.219 | החיצונה |  |
| 6:14.220 | מה |  |
| 6:14.221 | שלמדו |  |
| 6:14.222 | היחידים |  |
| 6:14.223 | חוץ |  |
| 6:14.224 | לאלו |  |
| 6:14.225 | בתי |  |
| 6:14.226 | המדרשות | . |
| 6:14.227 | וביד |  |
| 6:14.228 | יוסף |  |
| 6:14.229 | פי | ' |
| 6:14.230 | שכוונת |  |
| 6:14.231 | המדרש |  |
| 6:14.232 | על |  |
| 6:14.233 | הברייתות |  |
| 6:14.234 | וע"כ | . |
| 6:14.235 | וכן |  |
| 6:14.236 | מצאתי |  |
| 6:14.237 | סיוע |  |
| 6:14.238 | לפירושו |  |
| 6:14.239 | בשיר |  |
| 6:14.240 | השירים |  |
| 6:14.241 | רבה |  |
| 6:14.242 | וז"ל |  |
| 6:14.243 | ששים |  |
| 6:14.244 | המה |  |
| 6:14.245 | מלכות |  |
| 6:14.246 | אלו |  |
| 6:14.247 | ששים |  |
| 6:14.248 | מסכתות |  |
| 6:14.249 | של |  |
| 6:14.250 | הלכות | , |
| 6:14.251 | ושמונים |  |
| 6:14.252 | פילגשים |  |
| 6:14.253 | אלו |  |
| 6:14.254 | שמונים |  |
| 6:14.255 | פרשיות |  |
| 6:14.256 | שבתורת |  |
| 6:14.257 | כהנים | . |
| 6:14.258 | ועלמות |  |
| 6:14.259 | אין |  |
| 6:14.260 | מספר |  |
| 6:14.261 | אין |  |
| 6:14.262 | קץ |  |
| 6:14.263 | לתוספתות | . |
| 6:14.264 | אחת |  |
| 6:14.265 | היא |  |
| 6:14.266 | הן |  |
| 6:14.267 | חולקין |  |
| 6:14.268 | אלו |  |
| 6:14.269 | עם |  |
| 6:14.270 | אלו |  |
| 6:14.271 | וכולהון |  |
| 6:14.272 | דורשים |  |
| 6:14.273 | מטעם |  |
| 6:14.274 | אחד | , |
| 6:14.275 | מהלכה |  |
| 6:14.276 | אחת | , |
| 6:14.277 | מגזירה |  |
| 6:14.278 | שוה | , |
| 6:14.279 | מקל |  |
| 6:14.280 | וחומר |  |
| 6:14.281 | ע"כ | . |
| 6:14.282 | ראה |  |
| 6:14.283 | כמה |  |
| 6:14.284 | מחבבים |  |
| 6:14.285 | רז"ל |  |
| 6:14.286 | את |  |
| 6:14.287 | המשנה |  |
| 6:14.288 | והתלמוד |  |
| 6:14.289 | וקראו |  |
| 6:14.290 | להן |  |
| 6:14.291 | מלכות | . |
| 6:14.292 | ולא |  |
| 6:14.293 | שפחות |  |
| 6:14.294 | כאשר |  |
| 6:14.295 | עשה |  |
| 6:14.296 | הפילסוף |  |
| 6:14.297 | הרמאי |  |
| 6:14.298 | האדוק |  |
| 6:14.299 | בדעות |  |
| 6:14.300 | חיצוניות |  |
| 6:14.301 | הוא |  |
| 6:14.302 | מחבר |  |
| 6:14.303 | הזהר |  |
| 6:14.304 | כמו |  |
| 6:14.305 | שיתבאר |  |
| 6:14.306 | לקמן | . |
| 6:14.307 | ובבא |  |
| 6:14.308 | מציעא |  |
| 6:14.309 | פרק |  |
| 6:14.310 | אלו |  |
| 6:14.311 | מציאות |  |
| 6:14.312 | בבא |  |
| 6:14.313 | מציעא |  |
| 6:14.314 | דף |  |
| 6:14.315 | ל"ג | ) |
| 6:14.316 | גרסינן | . |
| 6:14.317 | ת"ר |  |
| 6:14.318 | העוסקים |  |
| 6:14.319 | במקרא |  |
| 6:14.320 | מדה |  |
| 6:14.321 | ואינה |  |
| 6:14.322 | מדה |  |
| 6:14.323 | במשנה |  |
| 6:14.324 | מדה |  |
| 6:14.325 | ונוטל |  |
| 6:14.326 | עליה |  |
| 6:14.327 | שכר | , |
| 6:14.328 | בתלמוד | , |
| 6:14.329 | אין |  |
| 6:14.330 | לך |  |
| 6:14.331 | מדה |  |
| 6:14.332 | גדולה |  |
| 6:14.333 | מזו | , |
| 6:14.334 | ולעולם |  |
| 6:14.335 | הוי |  |
| 6:14.336 | רץ |  |
| 6:14.337 | למשנה |  |
| 6:14.338 | יותר |  |
| 6:14.339 | מן |  |
| 6:14.340 | התלמוד | , |
| 6:14.341 | ופריך |  |
| 6:14.342 | היא |  |
| 6:14.343 | גופה |  |
| 6:14.344 | קשיא | , |
| 6:14.345 | אמרת |  |
| 6:14.346 | בתלמוד |  |
| 6:14.347 | אין |  |
| 6:14.348 | לך |  |
| 6:14.349 | מדה |  |
| 6:14.350 | גדולה |  |
| 6:14.351 | מזו | , |
| 6:14.352 | והדר |  |
| 6:14.353 | אמרת |  |
| 6:14.354 | ולעולם |  |
| 6:14.355 | הוי |  |
| 6:14.356 | רץ |  |
| 6:14.357 | למשנה |  |
| 6:14.358 | יותר |  |
| 6:14.359 | מן |  |
| 6:14.360 | התלמוד | ? |
| 6:14.361 | אמר |  |
| 6:14.362 | ר | ' |
| 6:14.363 | יוחנן |  |
| 6:14.364 | בימי |  |
| 6:14.365 | ר | ' |
| 6:14.366 | נשנית |  |
| 6:14.367 | משנה |  |
| 6:14.368 | זו | , |
| 6:14.369 | שבקו |  |
| 6:14.370 | כולי |  |
| 6:14.371 | עלמא |  |
| 6:14.372 | מתניתין |  |
| 6:14.373 | ואזלו |  |
| 6:14.374 | בתר |  |
| 6:14.375 | תלמודא | , |
| 6:14.376 | הדר |  |
| 6:14.377 | דרש |  |
| 6:14.378 | להו |  |
| 6:14.379 | ולעולם |  |
| 6:14.380 | הוי |  |
| 6:14.381 | רץ |  |
| 6:14.382 | למשנה |  |
| 6:14.383 | יותר |  |
| 6:14.384 | מן |  |
| 6:14.385 | התלמוד |  |
| 6:14.386 | עיין |  |
| 6:14.387 | שם |  |
| 6:14.388 | ובפי | ' |
| 6:14.389 | רש"י | . |
| 6:14.390 | ובירושלמי |  |
| 6:14.391 | מבואר |  |
| 6:14.392 | בהדיא |  |
| 6:14.393 | דרשב"י |  |
| 6:14.394 | הוא |  |
| 6:14.395 | דתני |  |
| 6:14.396 | לה | , |
| 6:14.397 | כדגרסינן |  |
| 6:14.398 | התם |  |
| 6:14.399 | בפרק |  |
| 6:14.400 | כל |  |
| 6:14.401 | כתבי |  |
| 6:14.402 | הדא |  |
| 6:14.403 | אמרא |  |
| 6:14.404 | שהמשנה |  |
| 6:14.405 | קודמת |  |
| 6:14.406 | למקרא | , |
| 6:14.407 | ודא |  |
| 6:14.408 | מסייעא |  |
| 6:14.409 | לרשב"י |  |
| 6:14.410 | דתני |  |
| 6:14.411 | רשב"י |  |
| 6:14.412 | העוסק |  |
| 6:14.413 | במקרא |  |
| 6:14.414 | מדה |  |
| 6:14.415 | שאינה |  |
| 6:14.416 | מדה | , |
| 6:14.417 | במשנה |  |
| 6:14.418 | מדה |  |
| 6:14.419 | שנוטלין |  |
| 6:14.420 | ממנה |  |
| 6:14.421 | שכר | , |
| 6:14.422 | העוסק |  |
| 6:14.423 | בתלמוד |  |
| 6:14.424 | אין |  |
| 6:14.425 | לך |  |
| 6:14.426 | מדה |  |
| 6:14.427 | גדולה |  |
| 6:14.428 | מזו | , |
| 6:14.429 | לעולם |  |
| 6:14.430 | הוי |  |
| 6:14.431 | רץ |  |
| 6:14.432 | אחר |  |
| 6:14.433 | המשנה |  |
| 6:14.434 | יותר |  |
| 6:14.435 | מן |  |
| 6:14.436 | התלמוד |  |
| 6:14.437 | אמר |  |
| 6:14.438 | ר | ' |
| 6:14.439 | יוסי |  |
| 6:14.440 | בר | ' |
| 6:14.441 | בון |  |
| 6:14.442 | הדא |  |
| 6:14.443 | דאת |  |
| 6:14.444 | אמר |  |
| 6:14.445 | עד |  |
| 6:14.446 | שלא |  |
| 6:14.447 | שקע |  |
| 6:14.448 | בו |  |
| 6:14.449 | רבי |  |
| 6:14.450 | רוב |  |
| 6:14.451 | משניות | , |
| 6:14.452 | אבל |  |
| 6:14.453 | מששיקע |  |
| 6:14.454 | בו |  |
| 6:14.455 | רבי |  |
| 6:14.456 | רוב |  |
| 6:14.457 | משניות |  |
| 6:14.458 | לעולם |  |
| 6:14.459 | הוי |  |
| 6:14.460 | רץ |  |
| 6:14.461 | אחר |  |
| 6:14.462 | התלמוד |  |
| 6:14.463 | יותר |  |
| 6:14.464 | מן |  |
| 6:14.465 | המשנה |  |
| 6:14.466 | ע"כ | . |
| 6:14.467 | וכן |  |
| 6:14.468 | בפ"ק |  |
| 6:14.469 | דשבת |  |
| 6:14.470 | סוף |  |
| 6:14.471 | פסקא |  |
| 6:14.472 | דלא |  |
| 6:14.473 | ישב |  |
| 6:14.474 | אדם |  |
| 6:14.475 | לפני |  |
| 6:14.476 | הספר |  |
| 6:14.477 | קאמר |  |
| 6:14.478 | התם |  |
| 6:14.479 | דהקורא |  |
| 6:14.480 | את |  |
| 6:14.481 | שמע |  |
| 6:14.482 | בעונתה |  |
| 6:14.483 | כמשנה |  |
| 6:14.484 | היא | , |
| 6:14.485 | ושלא |  |
| 6:14.486 | בעונתה |  |
| 6:14.487 | כאדם |  |
| 6:14.488 | שקורא |  |
| 6:14.489 | בתורה |  |
| 6:14.490 | ששכרו |  |
| 6:14.491 | פחות |  |
| 6:14.492 | מהעוסק |  |
| 6:14.493 | במשנה | , |
| 6:14.494 | ומסיים |  |
| 6:14.495 | בה |  |
| 6:14.496 | רשב"י |  |
| 6:14.497 | כדעתיה |  |
| 6:14.498 | דרשב"י |  |
| 6:14.499 | אמר |  |
| 6:14.500 | העוסק |  |
| 6:14.501 | במקרא |  |
| 6:14.502 | מדה |  |
| 6:14.503 | שאינה |  |
| 6:14.504 | מדה | , |
| 6:14.505 | ורבנין |  |
| 6:14.506 | עבדין |  |
| 6:14.507 | מקרא |  |
| 6:14.508 | כמשנה |  |
| 6:14.509 | עכ"ל | , |
| 6:14.510 | ובילקוט |  |
| 6:14.511 | ריש |  |
| 6:14.512 | פר | ' |
| 6:14.513 | צו |  |
| 6:14.514 | ר | ' |
| 6:14.515 | חנינה |  |
| 6:14.516 | אמר |  |
| 6:14.517 | אין |  |
| 6:14.518 | הגליות |  |
| 6:14.519 | מתכנסות |  |
| 6:14.520 | אלא |  |
| 6:14.521 | בזכות |  |
| 6:14.522 | משניות |  |
| 6:14.523 | שנ | ' |
| 6:14.524 | גם |  |
| 6:14.525 | כי |  |
| 6:14.526 | יתנו |  |
| 6:14.527 | בגוים |  |
| 6:14.528 | עתה |  |
| 6:14.529 | אקבצם |  |
| 6:14.530 | ע"כ | . |
| 6:14.531 | ראה |  |
| 6:14.532 | תראה |  |
| 6:14.533 | ידידי |  |
| 6:14.534 | הקורא |  |
| 6:14.535 | דברי |  |
| 6:14.536 | רז"ל |  |
| 6:14.537 | שהבאתי |  |
| 6:14.538 | מקצת |  |
| 6:14.539 | מהם |  |
| 6:14.540 | כמה |  |
| 6:14.541 | הם |  |
| 6:14.542 | מחבבים |  |
| 6:14.543 | ומחשיבים |  |
| 6:14.544 | את |  |
| 6:14.545 | המשנה |  |
| 6:14.546 | שהיא |  |
| 6:14.547 | עיקר |  |
| 6:14.548 | תורה |  |
| 6:14.549 | שבעל |  |
| 6:14.550 | פה |  |
| 6:14.551 | שנאמרה |  |
| 6:14.552 | למשה |  |
| 6:14.553 | בסיני | , |
| 6:14.554 | ובפרט |  |
| 6:14.555 | שים |  |
| 6:14.556 | עיניך |  |
| 6:14.557 | לדברי |  |
| 6:14.558 | תנא |  |
| 6:14.559 | דידן |  |
| 6:14.560 | רשב"י |  |
| 6:14.561 | הנאמרים |  |
| 6:14.562 | באמת |  |
| 6:14.563 | בבבלי |  |
| 6:14.564 | ובירושלמי |  |
| 6:14.565 | כמה |  |
| 6:14.566 | הוא |  |
| 6:14.567 | מחבב |  |
| 6:14.568 | ומחשיב |  |
| 6:14.569 | את |  |
| 6:14.570 | המשנה |  |
| 6:14.571 | יותר |  |
| 6:14.572 | מן |  |
| 6:14.573 | המקרא | , |
| 6:14.574 | ואת |  |
| 6:14.575 | התלמוד |  |
| 6:14.576 | שהוא |  |
| 6:14.577 | פי | ' |
| 6:14.578 | טעמי |  |
| 6:14.579 | המשנה |  |
| 6:14.580 | יותר |  |
| 6:14.581 | מכולם |  |
| 6:14.582 | שאין |  |
| 6:14.583 | לך |  |
| 6:14.584 | מדה |  |
| 6:14.585 | גדולה |  |
| 6:14.586 | יותר |  |
| 6:14.587 | ממנה | , |
| 6:14.588 | אלא |  |
| 6:14.589 | דבירושלמי |  |
| 6:14.590 | מוקי |  |
| 6:14.591 | להא |  |
| 6:14.592 | דאמר |  |
| 6:14.593 | לעולם |  |
| 6:14.594 | הוי |  |
| 6:14.595 | רץ |  |
| 6:14.596 | למשנה |  |
| 6:14.597 | יותר |  |
| 6:14.598 | מן |  |
| 6:14.599 | התלמוד |  |
| 6:14.600 | קודם |  |
| 6:14.601 | שסידר |  |
| 6:14.602 | רבי |  |
| 6:14.603 | את |  |
| 6:14.604 | המשנה |  |
| 6:14.605 | שאז |  |
| 6:14.606 | היו |  |
| 6:14.607 | הרבה |  |
| 6:14.608 | משניות |  |
| 6:14.609 | שהם |  |
| 6:14.610 | שגורים |  |
| 6:14.611 | בפי |  |
| 6:14.612 | התלמידים |  |
| 6:14.613 | כל |  |
| 6:14.614 | אחת |  |
| 6:14.615 | בשם |  |
| 6:14.616 | אומרה | , |
| 6:14.617 | והרבה |  |
| 6:14.618 | בריתות |  |
| 6:14.619 | היו |  |
| 6:14.620 | להם | , |
| 6:14.621 | אבל |  |
| 6:14.622 | לאחר |  |
| 6:14.623 | שסידר |  |
| 6:14.624 | רבי |  |
| 6:14.625 | את |  |
| 6:14.626 | המשנה |  |
| 6:14.627 | והשמיט |  |
| 6:14.628 | בה |  |
| 6:14.629 | כמה |  |
| 6:14.630 | משניות |  |
| 6:14.631 | ובריתות | , |
| 6:14.632 | וגם |  |
| 6:14.633 | סתם |  |
| 6:14.634 | במשנה |  |
| 6:14.635 | דברי |  |
| 6:14.636 | יחידים |  |
| 6:14.637 | שראה |  |
| 6:14.638 | ר | ' |
| 6:14.639 | דבריהם |  |
| 6:14.640 | ושנאן |  |
| 6:14.641 | סתם |  |
| 6:14.642 | כדי |  |
| 6:14.643 | לקבוע |  |
| 6:14.644 | הלכה |  |
| 6:14.645 | כמותם | , |
| 6:14.646 | לפיכך |  |
| 6:14.647 | חזרו |  |
| 6:14.648 | לומר |  |
| 6:14.649 | הוי |  |
| 6:14.650 | רץ |  |
| 6:14.651 | לתלמוד |  |
| 6:14.652 | יותר |  |
| 6:14.653 | מן |  |
| 6:14.654 | המשנה | . |
| 6:14.655 | ומדברי |  |
| 6:14.656 | כולם |  |
| 6:14.657 | למדנו |  |
| 6:14.658 | כמה |  |
| 6:14.659 | היה |  |
| 6:14.660 | חביב |  |
| 6:14.661 | להם |  |
| 6:14.662 | לימוד |  |
| 6:14.663 | המשנה |  |
| 6:14.664 | והתלמוד | , |
| 6:14.665 | ושרשב"י |  |
| 6:14.666 | הוא |  |
| 6:14.667 | דתני |  |
| 6:14.668 | לה | , |
| 6:14.669 | ומזה |  |
| 6:14.670 | הוכחה |  |
| 6:14.671 | גמורה |  |
| 6:14.672 | ברורה |  |
| 6:14.673 | שס | ' |
| 6:14.674 | הזהר |  |
| 6:14.675 | אינו |  |
| 6:14.676 | לרשב"י |  |
| 6:14.677 | התנא |  |
| 6:14.678 | ע"ה | , |
| 6:14.679 | כי |  |
| 6:14.680 | אם |  |
| 6:14.681 | לפילסוף |  |
| 6:14.682 | רמאי |  |
| 6:14.683 | גדול |  |
| 6:14.684 | שלמד |  |
| 6:14.685 | כל |  |
| 6:14.686 | חכמת |  |
| 6:14.687 | ישראל |  |
| 6:14.688 | בשני |  |
| 6:14.689 | התלמודים |  |
| 6:14.690 | ומדרשי |  |
| 6:14.691 | רז"ל | . |
| 6:14.692 | והוא |  |
| 6:14.693 | מאמין |  |
| 6:14.694 | בדעות |  |
| 6:14.695 | הכשדים |  |
| 6:14.696 | והכלדיים |  |
| 6:14.697 | פילסופים |  |
| 6:14.698 | הראשונים |  |
| 6:14.699 | שהיו |  |
| 6:14.700 | חולקים |  |
| 6:14.701 | על |  |
| 6:14.702 | אברהם |  |
| 6:14.703 | אבינו |  |
| 6:14.704 | באמונתו |  |
| 6:14.705 | באל |  |
| 6:14.706 | אחד | , |
| 6:14.707 | ומאמינים |  |
| 6:14.708 | בשתי |  |
| 6:14.709 | סבות |  |
| 6:14.710 | קודמות | , |
| 6:14.711 | וכל |  |
| 6:14.712 | אחת |  |
| 6:14.713 | פועלת |  |
| 6:14.714 | היפך |  |
| 6:14.715 | האחרת | , |
| 6:14.716 | והסבות |  |
| 6:14.717 | הפועלות |  |
| 6:14.718 | הטוב |  |
| 6:14.719 | קורא |  |
| 6:14.720 | להם |  |
| 6:14.721 | בשם |  |
| 6:14.722 | סטרא |  |
| 6:14.723 | דקדושה | , |
| 6:14.724 | והם |  |
| 6:14.725 | אדם |  |
| 6:14.726 | קדמון |  |
| 6:14.727 | וכתר |  |
| 6:14.728 | עליון |  |
| 6:14.729 | עם |  |
| 6:14.730 | שאר |  |
| 6:14.731 | הפרצופים |  |
| 6:14.732 | וצורות |  |
| 6:14.733 | המשתלשלים |  |
| 6:14.734 | מהם | , |
| 6:14.735 | והסבות |  |
| 6:14.736 | הפועלות |  |
| 6:14.737 | הרע |  |
| 6:14.738 | הוא |  |
| 6:14.739 | קורא |  |
| 6:14.740 | להם |  |
| 6:14.741 | סטרא |  |
| 6:14.742 | דקליפה | , |
| 6:14.743 | והם |  |
| 6:14.744 | אדם |  |
| 6:14.745 | בליעל |  |
| 6:14.746 | וכתר |  |
| 6:14.747 | עליון |  |
| 6:14.748 | דקליפה |  |
| 6:14.749 | עם |  |
| 6:14.750 | שאר |  |
| 6:14.751 | פרצופים |  |
| 6:14.752 | וצורות |  |
| 6:14.753 | דעשר |  |
| 6:14.754 | ספירות |  |
| 6:14.755 | דקליפה |  |
| 6:14.756 | המשתלשלים |  |
| 6:14.757 | מהם |  |
| 6:14.758 | הנקראים |  |
| 6:14.759 | בשם |  |
| 6:14.760 | אל |  |
| 6:14.761 | אחר | , |
| 6:14.762 | כי |  |
| 6:14.763 | הוא |  |
| 6:14.764 | מאמין |  |
| 6:14.765 | ומאמת |  |
| 6:14.766 | מציאות |  |
| 6:14.767 | אלהים |  |
| 6:14.768 | אחרים |  |
| 6:14.769 | בעולם |  |
| 6:14.770 | כדעת |  |
| 6:14.771 | הכשדים |  |
| 6:14.772 | והמצרים |  |
| 6:14.773 | הקדמונים |  |
| 6:14.774 | בפילוסופיא |  |
| 6:14.775 | הקדומה |  |
| 6:14.776 | שכבר |  |
| 6:14.777 | נמר |  |
| 6:14.778 | ריחה |  |
| 6:14.779 | וסר |  |
| 6:14.780 | טעמה | . |
| 6:14.781 | ובשפתי |  |
| 6:14.782 | חלקות |  |
| 6:14.783 | וגיזומים |  |
| 6:14.784 | הוא |  |
| 6:14.785 | מושך |  |
| 6:14.786 | לב |  |
| 6:14.787 | רבים |  |
| 6:14.788 | מהרבנים |  |
| 6:14.789 | הפתאים |  |
| 6:14.790 | להאמין |  |
| 6:14.791 | באלהים |  |
| 6:14.792 | אחרים |  |
| 6:14.793 | רבים |  |
| 6:14.794 | ולזלזל |  |
| 6:14.795 | בלמוד |  |
| 6:14.796 | המשנה |  |
| 6:14.797 | והתלמוד |  |
| 6:14.798 | עד |  |
| 6:14.799 | שלא |  |
| 6:14.800 | ידעו |  |
| 6:14.801 | מה |  |
| 6:14.802 | לעשות |  |
| 6:14.803 | והיאך |  |
| 6:14.804 | יפרשו |  |
| 6:14.805 | מן |  |
| 6:14.806 | האיסורי | ' |
| 6:14.807 | שנאמרו |  |
| 6:14.808 | למשה |  |
| 6:14.809 | בסיני |  |
| 6:14.810 | כגון |  |
| 6:14.811 | ח"י |  |
| 6:14.812 | טריפות |  |
| 6:14.813 | הנזכרים |  |
| 6:14.814 | במשנה |  |
| 6:14.815 | וחילוקי |  |
| 6:14.816 | דיניהם | , |
| 6:14.817 | לפי |  |
| 6:14.818 | שלא |  |
| 6:14.819 | נתפרש |  |
| 6:14.820 | בתורה |  |
| 6:14.821 | בהדיא |  |
| 6:14.822 | אלא |  |
| 6:14.823 | דרוסה |  |
| 6:14.824 | כמו |  |
| 6:14.825 | שכ | ' |
| 6:14.826 | הרמב"ם |  |
| 6:14.827 | ז"ל |  |
| 6:14.828 | בהל | ' |
| 6:14.829 | מאכלות |  |
| 6:14.830 | אסורות | . |
| 6:14.831 | ויכוצא |  |
| 6:14.832 | בזה |  |
| 6:14.833 | בשאר |  |
| 6:14.834 | איסורים |  |
| 6:14.835 | של |  |
| 6:14.836 | תורה |  |
| 6:14.837 | שלא |  |
| 6:14.838 | נתפרשו |  |
| 6:14.839 | בתורה |  |
| 6:14.840 | שבכתב |  |
| 6:14.841 | ונמסרו |  |
| 6:14.842 | למשה |  |
| 6:14.843 | רבי | ' |
| 6:14.844 | ע"ה |  |
| 6:14.845 | בעל |  |
| 6:14.846 | פה | . |
| 6:14.847 | והמה |  |
| 6:14.848 | מפורשים |  |
| 6:14.849 | במשנה |  |
| 6:14.850 | ובתלמוד |  |
| 6:14.851 | כגון |  |
| 6:14.852 | דיני |  |
| 6:14.853 | תערובת |  |
| 6:14.854 | איסור |  |
| 6:14.855 | ובשר |  |
| 6:14.856 | בחלב | . |
| 6:14.857 | ואיזה |  |
| 6:14.858 | מהן |  |
| 6:14.859 | אסור |  |
| 6:14.860 | מדאוריתא |  |
| 6:14.861 | ואיזה |  |
| 6:14.862 | אסור |  |
| 6:14.863 | מדרבנן | . |
| 6:14.864 | והפילסוף |  |
| 6:14.865 | הנז | ' |
| 6:14.866 | כיוון |  |
| 6:14.867 | לפסול |  |
| 6:14.868 | ולהשכיח |  |
| 6:14.869 | מישראל |  |
| 6:14.870 | כל |  |
| 6:14.871 | הדינים |  |
| 6:14.872 | המבוארים |  |
| 6:14.873 | במשנה |  |
| 6:14.874 | ובתלמוד | . |
| 6:14.875 | ובהדיא |  |
| 6:14.876 | אמר |  |
| 6:14.877 | שבסוף |  |
| 6:14.878 | גלותא |  |
| 6:14.879 | יתפרנסון |  |
| 6:14.880 | מחבורא |  |
| 6:14.881 | דיליה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

