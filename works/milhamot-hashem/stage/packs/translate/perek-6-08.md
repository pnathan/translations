# Translation pack: milhamot-hashem perek-6-08 (Milchamot ha-Shem 6:13–6:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-08`.

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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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


## Milchamot ha-Shem 6:13

Hebrew (Wikisource): סו) הן עתה אערוך לפניך ידידי הקורא, השם נגד פניו כבוד השי"ת וכבוד תורה שבעל פה שהיא המשנה והתלמוד אשר כל דורש את דבר ה' בתורת אלהים יפנה ואעתיק פה קצת מאמרים מרז"ל להודיעך איך הם מחבבים ומוקירים את המשנה ואת התלמוד. וכי הם עיקר תורה שבעל פה שקבל משה מסיני, ולא שינה מכל אשר צוהו ה' אלהינו כאשר העיד עליו השי"ת באמרו בכל ביתי נאמן הוא. ואחרי כן אערוך לעיניך הקורא את דברי נביא השקר והפילסוף הרמאי מחבר ס' הזהר איך הוא מבזה את המשנה והתלמוד והוציא לעז וחשד למשה רבי' ע"ה שמפי עצמו אמרה, ושעל זה נקנסה עליו מיתה ולהקבר בארץ נכריה חוץ לארץ, ומן המפורסם וידוע לכל שלפי רוב חבתה של משנה בעיני רבותינו ז"ל בעלי התלמוד הם מדקדקים בלשונה כדרך שעשו במקראי קדש כדגרסינן בעירובין וסוכה מאי שנא מבוי דתני תקנתא ומאי שנא סוכה דתני פסולה וכו'. לפי שרז"ל דבר ברור הוא אצלם שהמשנה היא בעצמה ההלכות שקבל משה בסיני ועליה עמד בהר ארבעים וארבעים לילה ללמוד כללותיה ופרטותיה אשר בהם נזכה לחיי העוה"ב כמקובל בכל האומה הישראלית וכמו שאמרו רז"ל כל השונה הלכות בכל יום מובטח לו שהוא בן העולם הבא. ואמרו עוד שאין לו להקב"ה בעולמו אלא ארבע אמות של הלכה. ור' אמי ור' אסי לא הוו מצלו בבי כנשתא אלא ביני עמודי היכי דגרסי. ובהוריות דף י"ג במעשה דר' מאיר ור' נתן שנתקנאו ברבן שמעון בן גמליאל שתיקן שכשיכנס הנשיא כל העם עומדים מפניו, וכשאב ב"ד נכנס עושין לו שתי שורות אחת מכאן ואחת מכאן, וכשחכם נכנס אחד ועומד ואחד יושב עד שישב במקומו, ובקשו להעביר את רשב"ג מנשיאותו על ידי שישאלו ממנו לשנות להם מסכת עוקצין דלא הוה גמיר לה, וכי לא ידע נימרו ליה מי ימלל גבורות ה' ישמיע כל תהלתו, למי נאה למלל גבורות ה', למי שיכול להשמיע כל תהלתו וכו', אלמא דקרו רבנן למשנה תהלות ה', ובמדרש תנחומא פרשת וירא גרסינן אמר רב יהודה בן שלום בקש משה שתהא המשנה אף היא בכתב וצפה הקב"ה שאומות העולם עתידן לתרגם את התורה ולהיות קורין אותה יונית והן אומרין אנו הן של ישראל. אמר לו הקב"ה למשה אכתב לו רובי תורתי ואם כן כמו זר נחשבו, וכל כך למה מפני שהמשנה היא מסתורין של הקב"ה. ואין הקב"ה מוסר מסתורין שלו אלא לצדיקים שנ' סוד ה' ליריאיו. ועיין מה שכ' בזה מהר"י מדובנא בספרו הנחמד אהל יעקב בחתימת ספר במדבר וינעם לך. ראה כמה היה חביבה עליהם המשנה וקראו לה מסתורין של הקב"ה שאינו מוסרה אלא לצדיקים שנ' סוד ה' ליריאיו. וכן במדרש שמות רבה פרשה מ"ז אמרו כתב לך את הדברים האלה וגו' הדא הוא דכתיב אכתב לו רובי תורתי כמו זר נחשבו בשעה שנגלה הקב"ה בסיני ליתן תורה לישראל אמרה למשה על הסדר מקרא, משנה, תלמוד, ואגדה, שנאמר וידבר אלהים את כל הדברים האלה, אפי' מה שתלמיד שואל לרב אמר הקב"ה למשה באותה שעה ומאחר שלמד אותה מפי הקב"ה אמר לו למדה לישראל, אמר לפניו רבונו של עולם אכתוב אותה להם, אמר איני מבקש ליתנה להם בכתב, מפני שגלוי לפני שאומות העולם עתידים לשלוט בהם וליטול אותה מהם ויהיו בזוים באומות (פי' שיאמרו אומות העולם שניתנה להם) אלא המקרא אני נותן להם בכתב והמשנה והתלמוד והאגדה בעל פה וכו' ומסיים כי על פי הדברים האלה זו המשנה והתלמוד שהם מבדילים בין ישראל לבין האומות ע"כ. וכן כתבו קדמונינו הר' חטר ור' זכריה הרופא בפי' הי"ג עיקרים בעיקר השמיני וז"ל ואנמי' כאן על פה לפי שידע הקב"ה והודיע למשה שאומות העולם עתידין ליקח התורה על ישראל. וכיון שלקחו תורה שבכתב לא יוכלו ליקח תורה שבעל פה מפני שהיא בסופרים ולא בספרים. ואלו לקחו תורה שבכתב עם תורה שבעל פה היו אומרים שהם ישראל בעדות נאמנה ע"כ. ראה ידידי הקורא והבן דברי רז"ל שאמרו כי על המשנה והתלמוד שהם יסוד כל תורה שבעל פה כרת השי"ת ברית עם ישראל. כי המה חיי רוחנו לעולם הבא והם רובה של תורה כדאיתא בגיטין פרק הנזקין. ובפרק קמא דברכות גרסינן ואמר ר' לוי בר חמא אמר ריש לקיש מאי דכתיב ואתנה לך את לוחות האבן והתורה והמצוה אשר כתבתי להורותם, לוחות אלו עשרת הדברות, והתורה זו מקרא. והמצוה זו משנה, אשר כתבתי אלו נביאים וכתובים, להורותם זו גמרא. מלמד שכולם ניתנו למשה מסיני ע"כ. ומביאו רבינו הרמב"ם בהקדמת ספר היד עיין שם.

| id | word | punctuation after |
|---|---|---|
| 6:13.1 | סו | ) |
| 6:13.2 | הן |  |
| 6:13.3 | עתה |  |
| 6:13.4 | אערוך |  |
| 6:13.5 | לפניך |  |
| 6:13.6 | ידידי |  |
| 6:13.7 | הקורא | , |
| 6:13.8 | השם |  |
| 6:13.9 | נגד |  |
| 6:13.10 | פניו |  |
| 6:13.11 | כבוד |  |
| 6:13.12 | השי"ת |  |
| 6:13.13 | וכבוד |  |
| 6:13.14 | תורה |  |
| 6:13.15 | שבעל |  |
| 6:13.16 | פה |  |
| 6:13.17 | שהיא |  |
| 6:13.18 | המשנה |  |
| 6:13.19 | והתלמוד |  |
| 6:13.20 | אשר |  |
| 6:13.21 | כל |  |
| 6:13.22 | דורש |  |
| 6:13.23 | את |  |
| 6:13.24 | דבר |  |
| 6:13.25 | ה | ' |
| 6:13.26 | בתורת |  |
| 6:13.27 | אלהים |  |
| 6:13.28 | יפנה |  |
| 6:13.29 | ואעתיק |  |
| 6:13.30 | פה |  |
| 6:13.31 | קצת |  |
| 6:13.32 | מאמרים |  |
| 6:13.33 | מרז"ל |  |
| 6:13.34 | להודיעך |  |
| 6:13.35 | איך |  |
| 6:13.36 | הם |  |
| 6:13.37 | מחבבים |  |
| 6:13.38 | ומוקירים |  |
| 6:13.39 | את |  |
| 6:13.40 | המשנה |  |
| 6:13.41 | ואת |  |
| 6:13.42 | התלמוד | . |
| 6:13.43 | וכי |  |
| 6:13.44 | הם |  |
| 6:13.45 | עיקר |  |
| 6:13.46 | תורה |  |
| 6:13.47 | שבעל |  |
| 6:13.48 | פה |  |
| 6:13.49 | שקבל |  |
| 6:13.50 | משה |  |
| 6:13.51 | מסיני | , |
| 6:13.52 | ולא |  |
| 6:13.53 | שינה |  |
| 6:13.54 | מכל |  |
| 6:13.55 | אשר |  |
| 6:13.56 | צוהו |  |
| 6:13.57 | ה | ' |
| 6:13.58 | אלהינו |  |
| 6:13.59 | כאשר |  |
| 6:13.60 | העיד |  |
| 6:13.61 | עליו |  |
| 6:13.62 | השי"ת |  |
| 6:13.63 | באמרו |  |
| 6:13.64 | בכל |  |
| 6:13.65 | ביתי |  |
| 6:13.66 | נאמן |  |
| 6:13.67 | הוא | . |
| 6:13.68 | ואחרי |  |
| 6:13.69 | כן |  |
| 6:13.70 | אערוך |  |
| 6:13.71 | לעיניך |  |
| 6:13.72 | הקורא |  |
| 6:13.73 | את |  |
| 6:13.74 | דברי |  |
| 6:13.75 | נביא |  |
| 6:13.76 | השקר |  |
| 6:13.77 | והפילסוף |  |
| 6:13.78 | הרמאי |  |
| 6:13.79 | מחבר |  |
| 6:13.80 | ס | ' |
| 6:13.81 | הזהר |  |
| 6:13.82 | איך |  |
| 6:13.83 | הוא |  |
| 6:13.84 | מבזה |  |
| 6:13.85 | את |  |
| 6:13.86 | המשנה |  |
| 6:13.87 | והתלמוד |  |
| 6:13.88 | והוציא |  |
| 6:13.89 | לעז |  |
| 6:13.90 | וחשד |  |
| 6:13.91 | למשה |  |
| 6:13.92 | רבי | ' |
| 6:13.93 | ע"ה |  |
| 6:13.94 | שמפי |  |
| 6:13.95 | עצמו |  |
| 6:13.96 | אמרה | , |
| 6:13.97 | ושעל |  |
| 6:13.98 | זה |  |
| 6:13.99 | נקנסה |  |
| 6:13.100 | עליו |  |
| 6:13.101 | מיתה |  |
| 6:13.102 | ולהקבר |  |
| 6:13.103 | בארץ |  |
| 6:13.104 | נכריה |  |
| 6:13.105 | חוץ |  |
| 6:13.106 | לארץ | , |
| 6:13.107 | ומן |  |
| 6:13.108 | המפורסם |  |
| 6:13.109 | וידוע |  |
| 6:13.110 | לכל |  |
| 6:13.111 | שלפי |  |
| 6:13.112 | רוב |  |
| 6:13.113 | חבתה |  |
| 6:13.114 | של |  |
| 6:13.115 | משנה |  |
| 6:13.116 | בעיני |  |
| 6:13.117 | רבותינו |  |
| 6:13.118 | ז"ל |  |
| 6:13.119 | בעלי |  |
| 6:13.120 | התלמוד |  |
| 6:13.121 | הם |  |
| 6:13.122 | מדקדקים |  |
| 6:13.123 | בלשונה |  |
| 6:13.124 | כדרך |  |
| 6:13.125 | שעשו |  |
| 6:13.126 | במקראי |  |
| 6:13.127 | קדש |  |
| 6:13.128 | כדגרסינן |  |
| 6:13.129 | בעירובין |  |
| 6:13.130 | וסוכה |  |
| 6:13.131 | מאי |  |
| 6:13.132 | שנא |  |
| 6:13.133 | מבוי |  |
| 6:13.134 | דתני |  |
| 6:13.135 | תקנתא |  |
| 6:13.136 | ומאי |  |
| 6:13.137 | שנא |  |
| 6:13.138 | סוכה |  |
| 6:13.139 | דתני |  |
| 6:13.140 | פסולה |  |
| 6:13.141 | וכו | '. |
| 6:13.142 | לפי |  |
| 6:13.143 | שרז"ל |  |
| 6:13.144 | דבר |  |
| 6:13.145 | ברור |  |
| 6:13.146 | הוא |  |
| 6:13.147 | אצלם |  |
| 6:13.148 | שהמשנה |  |
| 6:13.149 | היא |  |
| 6:13.150 | בעצמה |  |
| 6:13.151 | ההלכות |  |
| 6:13.152 | שקבל |  |
| 6:13.153 | משה |  |
| 6:13.154 | בסיני |  |
| 6:13.155 | ועליה |  |
| 6:13.156 | עמד |  |
| 6:13.157 | בהר |  |
| 6:13.158 | ארבעים |  |
| 6:13.159 | וארבעים |  |
| 6:13.160 | לילה |  |
| 6:13.161 | ללמוד |  |
| 6:13.162 | כללותיה |  |
| 6:13.163 | ופרטותיה |  |
| 6:13.164 | אשר |  |
| 6:13.165 | בהם |  |
| 6:13.166 | נזכה |  |
| 6:13.167 | לחיי |  |
| 6:13.168 | העוה"ב |  |
| 6:13.169 | כמקובל |  |
| 6:13.170 | בכל |  |
| 6:13.171 | האומה |  |
| 6:13.172 | הישראלית |  |
| 6:13.173 | וכמו |  |
| 6:13.174 | שאמרו |  |
| 6:13.175 | רז"ל |  |
| 6:13.176 | כל |  |
| 6:13.177 | השונה |  |
| 6:13.178 | הלכות |  |
| 6:13.179 | בכל |  |
| 6:13.180 | יום |  |
| 6:13.181 | מובטח |  |
| 6:13.182 | לו |  |
| 6:13.183 | שהוא |  |
| 6:13.184 | בן |  |
| 6:13.185 | העולם |  |
| 6:13.186 | הבא | . |
| 6:13.187 | ואמרו |  |
| 6:13.188 | עוד |  |
| 6:13.189 | שאין |  |
| 6:13.190 | לו |  |
| 6:13.191 | להקב"ה |  |
| 6:13.192 | בעולמו |  |
| 6:13.193 | אלא |  |
| 6:13.194 | ארבע |  |
| 6:13.195 | אמות |  |
| 6:13.196 | של |  |
| 6:13.197 | הלכה | . |
| 6:13.198 | ור | ' |
| 6:13.199 | אמי |  |
| 6:13.200 | ור | ' |
| 6:13.201 | אסי |  |
| 6:13.202 | לא |  |
| 6:13.203 | הוו |  |
| 6:13.204 | מצלו |  |
| 6:13.205 | בבי |  |
| 6:13.206 | כנשתא |  |
| 6:13.207 | אלא |  |
| 6:13.208 | ביני |  |
| 6:13.209 | עמודי |  |
| 6:13.210 | היכי |  |
| 6:13.211 | דגרסי | . |
| 6:13.212 | ובהוריות |  |
| 6:13.213 | דף |  |
| 6:13.214 | י"ג |  |
| 6:13.215 | במעשה |  |
| 6:13.216 | דר | ' |
| 6:13.217 | מאיר |  |
| 6:13.218 | ור | ' |
| 6:13.219 | נתן |  |
| 6:13.220 | שנתקנאו |  |
| 6:13.221 | ברבן |  |
| 6:13.222 | שמעון |  |
| 6:13.223 | בן |  |
| 6:13.224 | גמליאל |  |
| 6:13.225 | שתיקן |  |
| 6:13.226 | שכשיכנס |  |
| 6:13.227 | הנשיא |  |
| 6:13.228 | כל |  |
| 6:13.229 | העם |  |
| 6:13.230 | עומדים |  |
| 6:13.231 | מפניו | , |
| 6:13.232 | וכשאב |  |
| 6:13.233 | ב"ד |  |
| 6:13.234 | נכנס |  |
| 6:13.235 | עושין |  |
| 6:13.236 | לו |  |
| 6:13.237 | שתי |  |
| 6:13.238 | שורות |  |
| 6:13.239 | אחת |  |
| 6:13.240 | מכאן |  |
| 6:13.241 | ואחת |  |
| 6:13.242 | מכאן | , |
| 6:13.243 | וכשחכם |  |
| 6:13.244 | נכנס |  |
| 6:13.245 | אחד |  |
| 6:13.246 | ועומד |  |
| 6:13.247 | ואחד |  |
| 6:13.248 | יושב |  |
| 6:13.249 | עד |  |
| 6:13.250 | שישב |  |
| 6:13.251 | במקומו | , |
| 6:13.252 | ובקשו |  |
| 6:13.253 | להעביר |  |
| 6:13.254 | את |  |
| 6:13.255 | רשב"ג |  |
| 6:13.256 | מנשיאותו |  |
| 6:13.257 | על |  |
| 6:13.258 | ידי |  |
| 6:13.259 | שישאלו |  |
| 6:13.260 | ממנו |  |
| 6:13.261 | לשנות |  |
| 6:13.262 | להם |  |
| 6:13.263 | מסכת |  |
| 6:13.264 | עוקצין |  |
| 6:13.265 | דלא |  |
| 6:13.266 | הוה |  |
| 6:13.267 | גמיר |  |
| 6:13.268 | לה | , |
| 6:13.269 | וכי |  |
| 6:13.270 | לא |  |
| 6:13.271 | ידע |  |
| 6:13.272 | נימרו |  |
| 6:13.273 | ליה |  |
| 6:13.274 | מי |  |
| 6:13.275 | ימלל |  |
| 6:13.276 | גבורות |  |
| 6:13.277 | ה | ' |
| 6:13.278 | ישמיע |  |
| 6:13.279 | כל |  |
| 6:13.280 | תהלתו | , |
| 6:13.281 | למי |  |
| 6:13.282 | נאה |  |
| 6:13.283 | למלל |  |
| 6:13.284 | גבורות |  |
| 6:13.285 | ה | ', |
| 6:13.286 | למי |  |
| 6:13.287 | שיכול |  |
| 6:13.288 | להשמיע |  |
| 6:13.289 | כל |  |
| 6:13.290 | תהלתו |  |
| 6:13.291 | וכו | ', |
| 6:13.292 | אלמא |  |
| 6:13.293 | דקרו |  |
| 6:13.294 | רבנן |  |
| 6:13.295 | למשנה |  |
| 6:13.296 | תהלות |  |
| 6:13.297 | ה | ', |
| 6:13.298 | ובמדרש |  |
| 6:13.299 | תנחומא |  |
| 6:13.300 | פרשת |  |
| 6:13.301 | וירא |  |
| 6:13.302 | גרסינן |  |
| 6:13.303 | אמר |  |
| 6:13.304 | רב |  |
| 6:13.305 | יהודה |  |
| 6:13.306 | בן |  |
| 6:13.307 | שלום |  |
| 6:13.308 | בקש |  |
| 6:13.309 | משה |  |
| 6:13.310 | שתהא |  |
| 6:13.311 | המשנה |  |
| 6:13.312 | אף |  |
| 6:13.313 | היא |  |
| 6:13.314 | בכתב |  |
| 6:13.315 | וצפה |  |
| 6:13.316 | הקב"ה |  |
| 6:13.317 | שאומות |  |
| 6:13.318 | העולם |  |
| 6:13.319 | עתידן |  |
| 6:13.320 | לתרגם |  |
| 6:13.321 | את |  |
| 6:13.322 | התורה |  |
| 6:13.323 | ולהיות |  |
| 6:13.324 | קורין |  |
| 6:13.325 | אותה |  |
| 6:13.326 | יונית |  |
| 6:13.327 | והן |  |
| 6:13.328 | אומרין |  |
| 6:13.329 | אנו |  |
| 6:13.330 | הן |  |
| 6:13.331 | של |  |
| 6:13.332 | ישראל | . |
| 6:13.333 | אמר |  |
| 6:13.334 | לו |  |
| 6:13.335 | הקב"ה |  |
| 6:13.336 | למשה |  |
| 6:13.337 | אכתב |  |
| 6:13.338 | לו |  |
| 6:13.339 | רובי |  |
| 6:13.340 | תורתי |  |
| 6:13.341 | ואם |  |
| 6:13.342 | כן |  |
| 6:13.343 | כמו |  |
| 6:13.344 | זר |  |
| 6:13.345 | נחשבו | , |
| 6:13.346 | וכל |  |
| 6:13.347 | כך |  |
| 6:13.348 | למה |  |
| 6:13.349 | מפני |  |
| 6:13.350 | שהמשנה |  |
| 6:13.351 | היא |  |
| 6:13.352 | מסתורין |  |
| 6:13.353 | של |  |
| 6:13.354 | הקב"ה | . |
| 6:13.355 | ואין |  |
| 6:13.356 | הקב"ה |  |
| 6:13.357 | מוסר |  |
| 6:13.358 | מסתורין |  |
| 6:13.359 | שלו |  |
| 6:13.360 | אלא |  |
| 6:13.361 | לצדיקים |  |
| 6:13.362 | שנ | ' |
| 6:13.363 | סוד |  |
| 6:13.364 | ה | ' |
| 6:13.365 | ליריאיו | . |
| 6:13.366 | ועיין |  |
| 6:13.367 | מה |  |
| 6:13.368 | שכ | ' |
| 6:13.369 | בזה |  |
| 6:13.370 | מהר"י |  |
| 6:13.371 | מדובנא |  |
| 6:13.372 | בספרו |  |
| 6:13.373 | הנחמד |  |
| 6:13.374 | אהל |  |
| 6:13.375 | יעקב |  |
| 6:13.376 | בחתימת |  |
| 6:13.377 | ספר |  |
| 6:13.378 | במדבר |  |
| 6:13.379 | וינעם |  |
| 6:13.380 | לך | . |
| 6:13.381 | ראה |  |
| 6:13.382 | כמה |  |
| 6:13.383 | היה |  |
| 6:13.384 | חביבה |  |
| 6:13.385 | עליהם |  |
| 6:13.386 | המשנה |  |
| 6:13.387 | וקראו |  |
| 6:13.388 | לה |  |
| 6:13.389 | מסתורין |  |
| 6:13.390 | של |  |
| 6:13.391 | הקב"ה |  |
| 6:13.392 | שאינו |  |
| 6:13.393 | מוסרה |  |
| 6:13.394 | אלא |  |
| 6:13.395 | לצדיקים |  |
| 6:13.396 | שנ | ' |
| 6:13.397 | סוד |  |
| 6:13.398 | ה | ' |
| 6:13.399 | ליריאיו | . |
| 6:13.400 | וכן |  |
| 6:13.401 | במדרש |  |
| 6:13.402 | שמות |  |
| 6:13.403 | רבה |  |
| 6:13.404 | פרשה |  |
| 6:13.405 | מ"ז |  |
| 6:13.406 | אמרו |  |
| 6:13.407 | כתב |  |
| 6:13.408 | לך |  |
| 6:13.409 | את |  |
| 6:13.410 | הדברים |  |
| 6:13.411 | האלה |  |
| 6:13.412 | וגו | ' |
| 6:13.413 | הדא |  |
| 6:13.414 | הוא |  |
| 6:13.415 | דכתיב |  |
| 6:13.416 | אכתב |  |
| 6:13.417 | לו |  |
| 6:13.418 | רובי |  |
| 6:13.419 | תורתי |  |
| 6:13.420 | כמו |  |
| 6:13.421 | זר |  |
| 6:13.422 | נחשבו |  |
| 6:13.423 | בשעה |  |
| 6:13.424 | שנגלה |  |
| 6:13.425 | הקב"ה |  |
| 6:13.426 | בסיני |  |
| 6:13.427 | ליתן |  |
| 6:13.428 | תורה |  |
| 6:13.429 | לישראל |  |
| 6:13.430 | אמרה |  |
| 6:13.431 | למשה |  |
| 6:13.432 | על |  |
| 6:13.433 | הסדר |  |
| 6:13.434 | מקרא | , |
| 6:13.435 | משנה | , |
| 6:13.436 | תלמוד | , |
| 6:13.437 | ואגדה | , |
| 6:13.438 | שנאמר |  |
| 6:13.439 | וידבר |  |
| 6:13.440 | אלהים |  |
| 6:13.441 | את |  |
| 6:13.442 | כל |  |
| 6:13.443 | הדברים |  |
| 6:13.444 | האלה | , |
| 6:13.445 | אפי | ' |
| 6:13.446 | מה |  |
| 6:13.447 | שתלמיד |  |
| 6:13.448 | שואל |  |
| 6:13.449 | לרב |  |
| 6:13.450 | אמר |  |
| 6:13.451 | הקב"ה |  |
| 6:13.452 | למשה |  |
| 6:13.453 | באותה |  |
| 6:13.454 | שעה |  |
| 6:13.455 | ומאחר |  |
| 6:13.456 | שלמד |  |
| 6:13.457 | אותה |  |
| 6:13.458 | מפי |  |
| 6:13.459 | הקב"ה |  |
| 6:13.460 | אמר |  |
| 6:13.461 | לו |  |
| 6:13.462 | למדה |  |
| 6:13.463 | לישראל | , |
| 6:13.464 | אמר |  |
| 6:13.465 | לפניו |  |
| 6:13.466 | רבונו |  |
| 6:13.467 | של |  |
| 6:13.468 | עולם |  |
| 6:13.469 | אכתוב |  |
| 6:13.470 | אותה |  |
| 6:13.471 | להם | , |
| 6:13.472 | אמר |  |
| 6:13.473 | איני |  |
| 6:13.474 | מבקש |  |
| 6:13.475 | ליתנה |  |
| 6:13.476 | להם |  |
| 6:13.477 | בכתב | , |
| 6:13.478 | מפני |  |
| 6:13.479 | שגלוי |  |
| 6:13.480 | לפני |  |
| 6:13.481 | שאומות |  |
| 6:13.482 | העולם |  |
| 6:13.483 | עתידים |  |
| 6:13.484 | לשלוט |  |
| 6:13.485 | בהם |  |
| 6:13.486 | וליטול |  |
| 6:13.487 | אותה |  |
| 6:13.488 | מהם |  |
| 6:13.489 | ויהיו |  |
| 6:13.490 | בזוים |  |
| 6:13.491 | באומות |  |
| 6:13.492 | פי | ' |
| 6:13.493 | שיאמרו |  |
| 6:13.494 | אומות |  |
| 6:13.495 | העולם |  |
| 6:13.496 | שניתנה |  |
| 6:13.497 | להם | ) |
| 6:13.498 | אלא |  |
| 6:13.499 | המקרא |  |
| 6:13.500 | אני |  |
| 6:13.501 | נותן |  |
| 6:13.502 | להם |  |
| 6:13.503 | בכתב |  |
| 6:13.504 | והמשנה |  |
| 6:13.505 | והתלמוד |  |
| 6:13.506 | והאגדה |  |
| 6:13.507 | בעל |  |
| 6:13.508 | פה |  |
| 6:13.509 | וכו | ' |
| 6:13.510 | ומסיים |  |
| 6:13.511 | כי |  |
| 6:13.512 | על |  |
| 6:13.513 | פי |  |
| 6:13.514 | הדברים |  |
| 6:13.515 | האלה |  |
| 6:13.516 | זו |  |
| 6:13.517 | המשנה |  |
| 6:13.518 | והתלמוד |  |
| 6:13.519 | שהם |  |
| 6:13.520 | מבדילים |  |
| 6:13.521 | בין |  |
| 6:13.522 | ישראל |  |
| 6:13.523 | לבין |  |
| 6:13.524 | האומות |  |
| 6:13.525 | ע"כ | . |
| 6:13.526 | וכן |  |
| 6:13.527 | כתבו |  |
| 6:13.528 | קדמונינו |  |
| 6:13.529 | הר | ' |
| 6:13.530 | חטר |  |
| 6:13.531 | ור | ' |
| 6:13.532 | זכריה |  |
| 6:13.533 | הרופא |  |
| 6:13.534 | בפי | ' |
| 6:13.535 | הי"ג |  |
| 6:13.536 | עיקרים |  |
| 6:13.537 | בעיקר |  |
| 6:13.538 | השמיני |  |
| 6:13.539 | וז"ל |  |
| 6:13.540 | ואנמי | ' |
| 6:13.541 | כאן |  |
| 6:13.542 | על |  |
| 6:13.543 | פה |  |
| 6:13.544 | לפי |  |
| 6:13.545 | שידע |  |
| 6:13.546 | הקב"ה |  |
| 6:13.547 | והודיע |  |
| 6:13.548 | למשה |  |
| 6:13.549 | שאומות |  |
| 6:13.550 | העולם |  |
| 6:13.551 | עתידין |  |
| 6:13.552 | ליקח |  |
| 6:13.553 | התורה |  |
| 6:13.554 | על |  |
| 6:13.555 | ישראל | . |
| 6:13.556 | וכיון |  |
| 6:13.557 | שלקחו |  |
| 6:13.558 | תורה |  |
| 6:13.559 | שבכתב |  |
| 6:13.560 | לא |  |
| 6:13.561 | יוכלו |  |
| 6:13.562 | ליקח |  |
| 6:13.563 | תורה |  |
| 6:13.564 | שבעל |  |
| 6:13.565 | פה |  |
| 6:13.566 | מפני |  |
| 6:13.567 | שהיא |  |
| 6:13.568 | בסופרים |  |
| 6:13.569 | ולא |  |
| 6:13.570 | בספרים | . |
| 6:13.571 | ואלו |  |
| 6:13.572 | לקחו |  |
| 6:13.573 | תורה |  |
| 6:13.574 | שבכתב |  |
| 6:13.575 | עם |  |
| 6:13.576 | תורה |  |
| 6:13.577 | שבעל |  |
| 6:13.578 | פה |  |
| 6:13.579 | היו |  |
| 6:13.580 | אומרים |  |
| 6:13.581 | שהם |  |
| 6:13.582 | ישראל |  |
| 6:13.583 | בעדות |  |
| 6:13.584 | נאמנה |  |
| 6:13.585 | ע"כ | . |
| 6:13.586 | ראה |  |
| 6:13.587 | ידידי |  |
| 6:13.588 | הקורא |  |
| 6:13.589 | והבן |  |
| 6:13.590 | דברי |  |
| 6:13.591 | רז"ל |  |
| 6:13.592 | שאמרו |  |
| 6:13.593 | כי |  |
| 6:13.594 | על |  |
| 6:13.595 | המשנה |  |
| 6:13.596 | והתלמוד |  |
| 6:13.597 | שהם |  |
| 6:13.598 | יסוד |  |
| 6:13.599 | כל |  |
| 6:13.600 | תורה |  |
| 6:13.601 | שבעל |  |
| 6:13.602 | פה |  |
| 6:13.603 | כרת |  |
| 6:13.604 | השי"ת |  |
| 6:13.605 | ברית |  |
| 6:13.606 | עם |  |
| 6:13.607 | ישראל | . |
| 6:13.608 | כי |  |
| 6:13.609 | המה |  |
| 6:13.610 | חיי |  |
| 6:13.611 | רוחנו |  |
| 6:13.612 | לעולם |  |
| 6:13.613 | הבא |  |
| 6:13.614 | והם |  |
| 6:13.615 | רובה |  |
| 6:13.616 | של |  |
| 6:13.617 | תורה |  |
| 6:13.618 | כדאיתא |  |
| 6:13.619 | בגיטין |  |
| 6:13.620 | פרק |  |
| 6:13.621 | הנזקין | . |
| 6:13.622 | ובפרק |  |
| 6:13.623 | קמא |  |
| 6:13.624 | דברכות |  |
| 6:13.625 | גרסינן |  |
| 6:13.626 | ואמר |  |
| 6:13.627 | ר | ' |
| 6:13.628 | לוי |  |
| 6:13.629 | בר |  |
| 6:13.630 | חמא |  |
| 6:13.631 | אמר |  |
| 6:13.632 | ריש |  |
| 6:13.633 | לקיש |  |
| 6:13.634 | מאי |  |
| 6:13.635 | דכתיב |  |
| 6:13.636 | ואתנה |  |
| 6:13.637 | לך |  |
| 6:13.638 | את |  |
| 6:13.639 | לוחות |  |
| 6:13.640 | האבן |  |
| 6:13.641 | והתורה |  |
| 6:13.642 | והמצוה |  |
| 6:13.643 | אשר |  |
| 6:13.644 | כתבתי |  |
| 6:13.645 | להורותם | , |
| 6:13.646 | לוחות |  |
| 6:13.647 | אלו |  |
| 6:13.648 | עשרת |  |
| 6:13.649 | הדברות | , |
| 6:13.650 | והתורה |  |
| 6:13.651 | זו |  |
| 6:13.652 | מקרא | . |
| 6:13.653 | והמצוה |  |
| 6:13.654 | זו |  |
| 6:13.655 | משנה | , |
| 6:13.656 | אשר |  |
| 6:13.657 | כתבתי |  |
| 6:13.658 | אלו |  |
| 6:13.659 | נביאים |  |
| 6:13.660 | וכתובים | , |
| 6:13.661 | להורותם |  |
| 6:13.662 | זו |  |
| 6:13.663 | גמרא | . |
| 6:13.664 | מלמד |  |
| 6:13.665 | שכולם |  |
| 6:13.666 | ניתנו |  |
| 6:13.667 | למשה |  |
| 6:13.668 | מסיני |  |
| 6:13.669 | ע"כ | . |
| 6:13.670 | ומביאו |  |
| 6:13.671 | רבינו |  |
| 6:13.672 | הרמב"ם |  |
| 6:13.673 | בהקדמת |  |
| 6:13.674 | ספר |  |
| 6:13.675 | היד |  |
| 6:13.676 | עיין |  |
| 6:13.677 | שם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

