# Translation pack: milhamot-hashem perek-2-15 (Milchamot ha-Shem 2:31–2:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-15`.

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
- **אצילות** and its root: fixed and translated, so that the verbs the
  author strings together stay linked. אצילות → "emanation"; עולם האצילות /
  דאצילות → "the World of Emanation"; נאצל → "emanated"; הנאצלים →
  "the emanated beings"; המאציל → "the Emanator". The four worlds:
  "Emanation, Creation, Formation, Making" (אבי"ע). "נאצלו ונבראו ונוצרו
  ונעשו" → "were emanated, created, formed and made". Keep נאצל and נברא
  distinct everywhere: the author argues that Ze'ir Anpin is "נברא ולא בורא"
  → "created, not creator".
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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- An initial before a surname that fixes the person (ר"מ ריקנטי): expand and
  name him in a note ("Rabbi Menachem Rikanti"); an initial no surname fixes
  stays an initial with a note, as with R. A. and ר"י חיים ("Rabbi Y. Chayyim").
- דרך ארץ: "worldly occupation", as in the project's Avot.
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
- Names of works: "Guide for the Perplexed" is rendered as the Moreh
  Nevukhim; ממלכי as "of the kings of".
- Closing formulas (וחתם דבריו, עכ"ל, ע"כ) stand outside the closing quotation
  mark; opening formulas (וז"ל, כ' וז"ל) stand outside the opening one.
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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


## Milchamot ha-Shem 2:31

Hebrew (Wikisource): הראית לדעת מה שלא נתקע הריב"ש בה הוא מחמת שלא קבלה מפה אל פה, לא מחמת שום נדנוד כמו שאתה האומר. והא לך לשון העד השני הגאון בעל חוות יאיר ז"ל בקיצור גדול מאד ששאלו חכם אחד שבילה ימיו בגמרא ופוסקים, והנה רוח עלתה בקרבו ועוררה אותו ללמוד הקבלה להפלגת מעלותיה שהיא נשמת התורה, וראה כמה מאמרים מהזהר דאי אפשר ליעול קמי מלכא קדישא זולתה, וחכמי התלמוד לא צוו עליה רק הזהירו שלא לדרוש בה, במופלא ממך אל תדרוש וכו' ויבקש מהר' להוציאו ממבוכה, ויען ויאמר לו, אוי לי אם איעצך על הפרישה ולגרום רפיון ידים ומניעה מחכמה זו שהיא בלי ספק נשמת התורה ושורש האמונה כחכמה ודעת ותבונה, ובלי ספק דזכי לה הוא מבני עלייה, אשרי לו, מה טוב חלקו, ומה נעים גודלו, אהוב למעלה וכו'. העלה דברי יש"ר מקנדי"א שהאריך מאד וקבץ כעמיר גורנה רוב דעות שקדמוהו לקרב ולא לרחק, נגד ספר בחינת הדת שבקש להדיח רבים! ואמר שמצא יתד תקוע בעמוד ברזל במקום נאמן ליתלות בו להסכים לקבל שכר על הפרישה יותר מעל הדרישה בס' שמן למאור לאביו של הרב (וז"ל בקיצור והאי מלתא שכתב, בשם ר"מ גבאי דכל מי שלא קבל סתרי תורה מרבו וחושב לשקלן בדעתו עובר על לא תעשה לך פסל, אמרתי עם לבי לעורר על דרכי בני אדם בדורינו, ראיתי עוללים לא ראו אור יאמרו חמותי ראיתי אור, בקרבם טמון כשאור אפרוחי שלא נתפתהו עיניהם וכו'. ואפי' הגדולים אשר בארץ מי שהוא חריף וחרוץ בקי בחדרי תורה. אין כח בשכלו להשיג בחדרי אורה, כי אם מפי זקן ומורה. וז"ל הרמב"ן בהקדמתו לפי' התורה ואני הנני מביא בברית נאמנת והיא הנותנת עצה הוגנת לכל מסתכל בספר הזה, לבל יסבור סברא, ולא יחשוב מחשבות בדבר מכל הרמזים אשר אני כותב בסתרי התורה כי אני מודיע נאמנה שלא יושגו דברי בשום שכל ובינה, זולתי מפי חכם מקובל לאזן מקבל מבין. והסברא בהם איוולת רבת הנזקים, מונעת התועלת וכו' ואם תאמר שאני פי' הרמב"ן, שכל דבריו בקיצור ברמיזה, משא"כ שאר ספרי קבלה שביארו דבריהם. הלא תראה בשערי אורה כתב בכל דף לשון הנני רומז. וכן בס' בחיי קיצר מאד ורגיל על לשונו לומר והמשכיל יבין כ"ש ספר הזהר שהוא סגור ומסוגר, וז"ל עץ חיים בהקדמתו: עם היות שבודאי בדורות האחרונים מתפרנסין מהאי חיבורא. אין ספק כי מי החכמה לא יתגלו לכל אדם בכח עיון כפי שכלם האנושי אלא ע"י שפע אלהי המושפע עליו. ואם האחרון שבמקובלים הפליג לומר שלא יושגו דבריו איך יעלה בדעת שכל אנושי בפרט באנשי דורנו זה להבין ולהשכיל דברי אלהים חיים הרשב"י אשר קול דברותיו מלהבי אש, והם חתומים באלף עזקן וכו'. לכן אנשי לבב שמעו לי ואל תהרסון לראות ספרים אחרונים הבנויים על פי שכלם האנושי ע"ש ולדעתי כי כשאדם חוטא בהרהור המחשבה בעולם האצילות גדול עונו מנשוא. אע"פ שרבים וכן שלמים ממקובלים אחרונים עוררו לבות בני אדם מאד על לימוד הקבלה וכתבו שהחושך עצמו מללמוד קבלה הוא נדחה ממחיצתו ומפסיד עולמו. נראה לי שלא פליגי ובודאי מי שיכול ללמוד מפי מקובל אלהי אשר יסביר לו העניינים כאשר עשה האר"י לתלמידיו אשרי חלקו. ואע"ג דהאחרונים התירו ללמוד מן הספרים ולהבין מדעתו. טוב ציפורנם של ראשונים וכו' ושב ואל תעשה עדיף דחמירא סכנתא דגופא וכ"ש דנפש וכו' ואין לנו לחקור בנסתרות כי אם לשמוע בלימודי פשטי המקראות ותוכחת מוסר להתנהג בדרכי ישרים כי ידיעתינו קצרה ע"כ לשון שמן המאור אשר הראיה ממנו להסכים על הפרישה וכו'. ועם כל זה גליתי לך דעתי ולא תסמוך עלי לא אתה ולא אחר כי איני כדאי להכריע והנה אני מדמה לימוד קדש קדשים הזה לנסיעה לארץ הקדש כי ידוע מעלתה כמו שאמרו בתלמוד שהרי בזולתה דומה כמי שאין לו אלוה ובעובד ע"ז והמהלך ד' אמות בא"י הוא בן העה"ב וכמה שבחים. ואפילו הכי כתבו התוס' בשם רבי חיים דעכשיו אינו מצוה לדור בארץ. כי יש כמה מצות התלוים בארץ שאין אנו יכולים לעמוד עליהם. ועוד מפני סכנת דרכים, ומפני העוני שם יעבירנו על דעתו ועל דעת קונו שם. ובאם אין לפניו בקי ושמירה מעולה מכל נזקים טוב למנוע מזה! הלא נמי מי יכחיש מעלת חכמה זו הגבה למעלה בלי ערך. בתנאי לדזכי לה ללמדה מפי מקובל שקבלה ג"כ מרבו. והן אמת שאני תמיה על ספרים שחוברו מאחרונים שהביאו מאמרים מהזהר. והמה מקשים בהם ומבארים כפי שיקול דעתם. מי נתן רשות לשכל אנושי להמציא משיקול דעתו במציאות שלש ראשונות ושלש צחצחות, כי מבואר בזהר דאית מקום דלא קימא לשאלה וכו' ואין ראיה ממאמר הזהר וס' יצירה ור"מ שקבילו מפי סבא דסבין ואליהו ומקדישין עילאין וכל הנך ינוקי דבזהר רוח ה' דבר בם ומלתו על לשונם, וגם האר"י וגוריו שקיבלו ממנו פה אל פה מפי עליונים יצאו הדברים ואנן בלמודינו באותן לא יעלה בדעתינו זולת קריאת המלות, מה שאין כן אלו האחרונים והמבארים משיקול הדעת.

| id | word | punctuation after |
|---|---|---|
| 2:31.1 | הראית |  |
| 2:31.2 | לדעת |  |
| 2:31.3 | מה |  |
| 2:31.4 | שלא |  |
| 2:31.5 | נתקע |  |
| 2:31.6 | הריב"ש |  |
| 2:31.7 | בה |  |
| 2:31.8 | הוא |  |
| 2:31.9 | מחמת |  |
| 2:31.10 | שלא |  |
| 2:31.11 | קבלה |  |
| 2:31.12 | מפה |  |
| 2:31.13 | אל |  |
| 2:31.14 | פה | , |
| 2:31.15 | לא |  |
| 2:31.16 | מחמת |  |
| 2:31.17 | שום |  |
| 2:31.18 | נדנוד |  |
| 2:31.19 | כמו |  |
| 2:31.20 | שאתה |  |
| 2:31.21 | האומר | . |
| 2:31.22 | והא |  |
| 2:31.23 | לך |  |
| 2:31.24 | לשון |  |
| 2:31.25 | העד |  |
| 2:31.26 | השני |  |
| 2:31.27 | הגאון |  |
| 2:31.28 | בעל |  |
| 2:31.29 | חוות |  |
| 2:31.30 | יאיר |  |
| 2:31.31 | ז"ל |  |
| 2:31.32 | בקיצור |  |
| 2:31.33 | גדול |  |
| 2:31.34 | מאד |  |
| 2:31.35 | ששאלו |  |
| 2:31.36 | חכם |  |
| 2:31.37 | אחד |  |
| 2:31.38 | שבילה |  |
| 2:31.39 | ימיו |  |
| 2:31.40 | בגמרא |  |
| 2:31.41 | ופוסקים | , |
| 2:31.42 | והנה |  |
| 2:31.43 | רוח |  |
| 2:31.44 | עלתה |  |
| 2:31.45 | בקרבו |  |
| 2:31.46 | ועוררה |  |
| 2:31.47 | אותו |  |
| 2:31.48 | ללמוד |  |
| 2:31.49 | הקבלה |  |
| 2:31.50 | להפלגת |  |
| 2:31.51 | מעלותיה |  |
| 2:31.52 | שהיא |  |
| 2:31.53 | נשמת |  |
| 2:31.54 | התורה | , |
| 2:31.55 | וראה |  |
| 2:31.56 | כמה |  |
| 2:31.57 | מאמרים |  |
| 2:31.58 | מהזהר |  |
| 2:31.59 | דאי |  |
| 2:31.60 | אפשר |  |
| 2:31.61 | ליעול |  |
| 2:31.62 | קמי |  |
| 2:31.63 | מלכא |  |
| 2:31.64 | קדישא |  |
| 2:31.65 | זולתה | , |
| 2:31.66 | וחכמי |  |
| 2:31.67 | התלמוד |  |
| 2:31.68 | לא |  |
| 2:31.69 | צוו |  |
| 2:31.70 | עליה |  |
| 2:31.71 | רק |  |
| 2:31.72 | הזהירו |  |
| 2:31.73 | שלא |  |
| 2:31.74 | לדרוש |  |
| 2:31.75 | בה | , |
| 2:31.76 | במופלא |  |
| 2:31.77 | ממך |  |
| 2:31.78 | אל |  |
| 2:31.79 | תדרוש |  |
| 2:31.80 | וכו | ' |
| 2:31.81 | ויבקש |  |
| 2:31.82 | מהר | ' |
| 2:31.83 | להוציאו |  |
| 2:31.84 | ממבוכה | , |
| 2:31.85 | ויען |  |
| 2:31.86 | ויאמר |  |
| 2:31.87 | לו | , |
| 2:31.88 | אוי |  |
| 2:31.89 | לי |  |
| 2:31.90 | אם |  |
| 2:31.91 | איעצך |  |
| 2:31.92 | על |  |
| 2:31.93 | הפרישה |  |
| 2:31.94 | ולגרום |  |
| 2:31.95 | רפיון |  |
| 2:31.96 | ידים |  |
| 2:31.97 | ומניעה |  |
| 2:31.98 | מחכמה |  |
| 2:31.99 | זו |  |
| 2:31.100 | שהיא |  |
| 2:31.101 | בלי |  |
| 2:31.102 | ספק |  |
| 2:31.103 | נשמת |  |
| 2:31.104 | התורה |  |
| 2:31.105 | ושורש |  |
| 2:31.106 | האמונה |  |
| 2:31.107 | כחכמה |  |
| 2:31.108 | ודעת |  |
| 2:31.109 | ותבונה | , |
| 2:31.110 | ובלי |  |
| 2:31.111 | ספק |  |
| 2:31.112 | דזכי |  |
| 2:31.113 | לה |  |
| 2:31.114 | הוא |  |
| 2:31.115 | מבני |  |
| 2:31.116 | עלייה | , |
| 2:31.117 | אשרי |  |
| 2:31.118 | לו | , |
| 2:31.119 | מה |  |
| 2:31.120 | טוב |  |
| 2:31.121 | חלקו | , |
| 2:31.122 | ומה |  |
| 2:31.123 | נעים |  |
| 2:31.124 | גודלו | , |
| 2:31.125 | אהוב |  |
| 2:31.126 | למעלה |  |
| 2:31.127 | וכו | '. |
| 2:31.128 | העלה |  |
| 2:31.129 | דברי |  |
| 2:31.130 | יש"ר |  |
| 2:31.131 | מקנדי"א |  |
| 2:31.132 | שהאריך |  |
| 2:31.133 | מאד |  |
| 2:31.134 | וקבץ |  |
| 2:31.135 | כעמיר |  |
| 2:31.136 | גורנה |  |
| 2:31.137 | רוב |  |
| 2:31.138 | דעות |  |
| 2:31.139 | שקדמוהו |  |
| 2:31.140 | לקרב |  |
| 2:31.141 | ולא |  |
| 2:31.142 | לרחק | , |
| 2:31.143 | נגד |  |
| 2:31.144 | ספר |  |
| 2:31.145 | בחינת |  |
| 2:31.146 | הדת |  |
| 2:31.147 | שבקש |  |
| 2:31.148 | להדיח |  |
| 2:31.149 | רבים | ! |
| 2:31.150 | ואמר |  |
| 2:31.151 | שמצא |  |
| 2:31.152 | יתד |  |
| 2:31.153 | תקוע |  |
| 2:31.154 | בעמוד |  |
| 2:31.155 | ברזל |  |
| 2:31.156 | במקום |  |
| 2:31.157 | נאמן |  |
| 2:31.158 | ליתלות |  |
| 2:31.159 | בו |  |
| 2:31.160 | להסכים |  |
| 2:31.161 | לקבל |  |
| 2:31.162 | שכר |  |
| 2:31.163 | על |  |
| 2:31.164 | הפרישה |  |
| 2:31.165 | יותר |  |
| 2:31.166 | מעל |  |
| 2:31.167 | הדרישה |  |
| 2:31.168 | בס | ' |
| 2:31.169 | שמן |  |
| 2:31.170 | למאור |  |
| 2:31.171 | לאביו |  |
| 2:31.172 | של |  |
| 2:31.173 | הרב |  |
| 2:31.174 | וז"ל |  |
| 2:31.175 | בקיצור |  |
| 2:31.176 | והאי |  |
| 2:31.177 | מלתא |  |
| 2:31.178 | שכתב | , |
| 2:31.179 | בשם |  |
| 2:31.180 | ר"מ |  |
| 2:31.181 | גבאי |  |
| 2:31.182 | דכל |  |
| 2:31.183 | מי |  |
| 2:31.184 | שלא |  |
| 2:31.185 | קבל |  |
| 2:31.186 | סתרי |  |
| 2:31.187 | תורה |  |
| 2:31.188 | מרבו |  |
| 2:31.189 | וחושב |  |
| 2:31.190 | לשקלן |  |
| 2:31.191 | בדעתו |  |
| 2:31.192 | עובר |  |
| 2:31.193 | על |  |
| 2:31.194 | לא |  |
| 2:31.195 | תעשה |  |
| 2:31.196 | לך |  |
| 2:31.197 | פסל | , |
| 2:31.198 | אמרתי |  |
| 2:31.199 | עם |  |
| 2:31.200 | לבי |  |
| 2:31.201 | לעורר |  |
| 2:31.202 | על |  |
| 2:31.203 | דרכי |  |
| 2:31.204 | בני |  |
| 2:31.205 | אדם |  |
| 2:31.206 | בדורינו | , |
| 2:31.207 | ראיתי |  |
| 2:31.208 | עוללים |  |
| 2:31.209 | לא |  |
| 2:31.210 | ראו |  |
| 2:31.211 | אור |  |
| 2:31.212 | יאמרו |  |
| 2:31.213 | חמותי |  |
| 2:31.214 | ראיתי |  |
| 2:31.215 | אור | , |
| 2:31.216 | בקרבם |  |
| 2:31.217 | טמון |  |
| 2:31.218 | כשאור |  |
| 2:31.219 | אפרוחי |  |
| 2:31.220 | שלא |  |
| 2:31.221 | נתפתהו |  |
| 2:31.222 | עיניהם |  |
| 2:31.223 | וכו | '. |
| 2:31.224 | ואפי | ' |
| 2:31.225 | הגדולים |  |
| 2:31.226 | אשר |  |
| 2:31.227 | בארץ |  |
| 2:31.228 | מי |  |
| 2:31.229 | שהוא |  |
| 2:31.230 | חריף |  |
| 2:31.231 | וחרוץ |  |
| 2:31.232 | בקי |  |
| 2:31.233 | בחדרי |  |
| 2:31.234 | תורה | . |
| 2:31.235 | אין |  |
| 2:31.236 | כח |  |
| 2:31.237 | בשכלו |  |
| 2:31.238 | להשיג |  |
| 2:31.239 | בחדרי |  |
| 2:31.240 | אורה | , |
| 2:31.241 | כי |  |
| 2:31.242 | אם |  |
| 2:31.243 | מפי |  |
| 2:31.244 | זקן |  |
| 2:31.245 | ומורה | . |
| 2:31.246 | וז"ל |  |
| 2:31.247 | הרמב"ן |  |
| 2:31.248 | בהקדמתו |  |
| 2:31.249 | לפי | ' |
| 2:31.250 | התורה |  |
| 2:31.251 | ואני |  |
| 2:31.252 | הנני |  |
| 2:31.253 | מביא |  |
| 2:31.254 | בברית |  |
| 2:31.255 | נאמנת |  |
| 2:31.256 | והיא |  |
| 2:31.257 | הנותנת |  |
| 2:31.258 | עצה |  |
| 2:31.259 | הוגנת |  |
| 2:31.260 | לכל |  |
| 2:31.261 | מסתכל |  |
| 2:31.262 | בספר |  |
| 2:31.263 | הזה | , |
| 2:31.264 | לבל |  |
| 2:31.265 | יסבור |  |
| 2:31.266 | סברא | , |
| 2:31.267 | ולא |  |
| 2:31.268 | יחשוב |  |
| 2:31.269 | מחשבות |  |
| 2:31.270 | בדבר |  |
| 2:31.271 | מכל |  |
| 2:31.272 | הרמזים |  |
| 2:31.273 | אשר |  |
| 2:31.274 | אני |  |
| 2:31.275 | כותב |  |
| 2:31.276 | בסתרי |  |
| 2:31.277 | התורה |  |
| 2:31.278 | כי |  |
| 2:31.279 | אני |  |
| 2:31.280 | מודיע |  |
| 2:31.281 | נאמנה |  |
| 2:31.282 | שלא |  |
| 2:31.283 | יושגו |  |
| 2:31.284 | דברי |  |
| 2:31.285 | בשום |  |
| 2:31.286 | שכל |  |
| 2:31.287 | ובינה | , |
| 2:31.288 | זולתי |  |
| 2:31.289 | מפי |  |
| 2:31.290 | חכם |  |
| 2:31.291 | מקובל |  |
| 2:31.292 | לאזן |  |
| 2:31.293 | מקבל |  |
| 2:31.294 | מבין | . |
| 2:31.295 | והסברא |  |
| 2:31.296 | בהם |  |
| 2:31.297 | איוולת |  |
| 2:31.298 | רבת |  |
| 2:31.299 | הנזקים | , |
| 2:31.300 | מונעת |  |
| 2:31.301 | התועלת |  |
| 2:31.302 | וכו | ' |
| 2:31.303 | ואם |  |
| 2:31.304 | תאמר |  |
| 2:31.305 | שאני |  |
| 2:31.306 | פי | ' |
| 2:31.307 | הרמב"ן | , |
| 2:31.308 | שכל |  |
| 2:31.309 | דבריו |  |
| 2:31.310 | בקיצור |  |
| 2:31.311 | ברמיזה | , |
| 2:31.312 | משא"כ |  |
| 2:31.313 | שאר |  |
| 2:31.314 | ספרי |  |
| 2:31.315 | קבלה |  |
| 2:31.316 | שביארו |  |
| 2:31.317 | דבריהם | . |
| 2:31.318 | הלא |  |
| 2:31.319 | תראה |  |
| 2:31.320 | בשערי |  |
| 2:31.321 | אורה |  |
| 2:31.322 | כתב |  |
| 2:31.323 | בכל |  |
| 2:31.324 | דף |  |
| 2:31.325 | לשון |  |
| 2:31.326 | הנני |  |
| 2:31.327 | רומז | . |
| 2:31.328 | וכן |  |
| 2:31.329 | בס | ' |
| 2:31.330 | בחיי |  |
| 2:31.331 | קיצר |  |
| 2:31.332 | מאד |  |
| 2:31.333 | ורגיל |  |
| 2:31.334 | על |  |
| 2:31.335 | לשונו |  |
| 2:31.336 | לומר |  |
| 2:31.337 | והמשכיל |  |
| 2:31.338 | יבין |  |
| 2:31.339 | כ"ש |  |
| 2:31.340 | ספר |  |
| 2:31.341 | הזהר |  |
| 2:31.342 | שהוא |  |
| 2:31.343 | סגור |  |
| 2:31.344 | ומסוגר | , |
| 2:31.345 | וז"ל |  |
| 2:31.346 | עץ |  |
| 2:31.347 | חיים |  |
| 2:31.348 | בהקדמתו | : |
| 2:31.349 | עם |  |
| 2:31.350 | היות |  |
| 2:31.351 | שבודאי |  |
| 2:31.352 | בדורות |  |
| 2:31.353 | האחרונים |  |
| 2:31.354 | מתפרנסין |  |
| 2:31.355 | מהאי |  |
| 2:31.356 | חיבורא | . |
| 2:31.357 | אין |  |
| 2:31.358 | ספק |  |
| 2:31.359 | כי |  |
| 2:31.360 | מי |  |
| 2:31.361 | החכמה |  |
| 2:31.362 | לא |  |
| 2:31.363 | יתגלו |  |
| 2:31.364 | לכל |  |
| 2:31.365 | אדם |  |
| 2:31.366 | בכח |  |
| 2:31.367 | עיון |  |
| 2:31.368 | כפי |  |
| 2:31.369 | שכלם |  |
| 2:31.370 | האנושי |  |
| 2:31.371 | אלא |  |
| 2:31.372 | ע"י |  |
| 2:31.373 | שפע |  |
| 2:31.374 | אלהי |  |
| 2:31.375 | המושפע |  |
| 2:31.376 | עליו | . |
| 2:31.377 | ואם |  |
| 2:31.378 | האחרון |  |
| 2:31.379 | שבמקובלים |  |
| 2:31.380 | הפליג |  |
| 2:31.381 | לומר |  |
| 2:31.382 | שלא |  |
| 2:31.383 | יושגו |  |
| 2:31.384 | דבריו |  |
| 2:31.385 | איך |  |
| 2:31.386 | יעלה |  |
| 2:31.387 | בדעת |  |
| 2:31.388 | שכל |  |
| 2:31.389 | אנושי |  |
| 2:31.390 | בפרט |  |
| 2:31.391 | באנשי |  |
| 2:31.392 | דורנו |  |
| 2:31.393 | זה |  |
| 2:31.394 | להבין |  |
| 2:31.395 | ולהשכיל |  |
| 2:31.396 | דברי |  |
| 2:31.397 | אלהים |  |
| 2:31.398 | חיים |  |
| 2:31.399 | הרשב"י |  |
| 2:31.400 | אשר |  |
| 2:31.401 | קול |  |
| 2:31.402 | דברותיו |  |
| 2:31.403 | מלהבי |  |
| 2:31.404 | אש | , |
| 2:31.405 | והם |  |
| 2:31.406 | חתומים |  |
| 2:31.407 | באלף |  |
| 2:31.408 | עזקן |  |
| 2:31.409 | וכו | '. |
| 2:31.410 | לכן |  |
| 2:31.411 | אנשי |  |
| 2:31.412 | לבב |  |
| 2:31.413 | שמעו |  |
| 2:31.414 | לי |  |
| 2:31.415 | ואל |  |
| 2:31.416 | תהרסון |  |
| 2:31.417 | לראות |  |
| 2:31.418 | ספרים |  |
| 2:31.419 | אחרונים |  |
| 2:31.420 | הבנויים |  |
| 2:31.421 | על |  |
| 2:31.422 | פי |  |
| 2:31.423 | שכלם |  |
| 2:31.424 | האנושי |  |
| 2:31.425 | ע"ש |  |
| 2:31.426 | ולדעתי |  |
| 2:31.427 | כי |  |
| 2:31.428 | כשאדם |  |
| 2:31.429 | חוטא |  |
| 2:31.430 | בהרהור |  |
| 2:31.431 | המחשבה |  |
| 2:31.432 | בעולם |  |
| 2:31.433 | האצילות |  |
| 2:31.434 | גדול |  |
| 2:31.435 | עונו |  |
| 2:31.436 | מנשוא | . |
| 2:31.437 | אע"פ |  |
| 2:31.438 | שרבים |  |
| 2:31.439 | וכן |  |
| 2:31.440 | שלמים |  |
| 2:31.441 | ממקובלים |  |
| 2:31.442 | אחרונים |  |
| 2:31.443 | עוררו |  |
| 2:31.444 | לבות |  |
| 2:31.445 | בני |  |
| 2:31.446 | אדם |  |
| 2:31.447 | מאד |  |
| 2:31.448 | על |  |
| 2:31.449 | לימוד |  |
| 2:31.450 | הקבלה |  |
| 2:31.451 | וכתבו |  |
| 2:31.452 | שהחושך |  |
| 2:31.453 | עצמו |  |
| 2:31.454 | מללמוד |  |
| 2:31.455 | קבלה |  |
| 2:31.456 | הוא |  |
| 2:31.457 | נדחה |  |
| 2:31.458 | ממחיצתו |  |
| 2:31.459 | ומפסיד |  |
| 2:31.460 | עולמו | . |
| 2:31.461 | נראה |  |
| 2:31.462 | לי |  |
| 2:31.463 | שלא |  |
| 2:31.464 | פליגי |  |
| 2:31.465 | ובודאי |  |
| 2:31.466 | מי |  |
| 2:31.467 | שיכול |  |
| 2:31.468 | ללמוד |  |
| 2:31.469 | מפי |  |
| 2:31.470 | מקובל |  |
| 2:31.471 | אלהי |  |
| 2:31.472 | אשר |  |
| 2:31.473 | יסביר |  |
| 2:31.474 | לו |  |
| 2:31.475 | העניינים |  |
| 2:31.476 | כאשר |  |
| 2:31.477 | עשה |  |
| 2:31.478 | האר"י |  |
| 2:31.479 | לתלמידיו |  |
| 2:31.480 | אשרי |  |
| 2:31.481 | חלקו | . |
| 2:31.482 | ואע"ג |  |
| 2:31.483 | דהאחרונים |  |
| 2:31.484 | התירו |  |
| 2:31.485 | ללמוד |  |
| 2:31.486 | מן |  |
| 2:31.487 | הספרים |  |
| 2:31.488 | ולהבין |  |
| 2:31.489 | מדעתו | . |
| 2:31.490 | טוב |  |
| 2:31.491 | ציפורנם |  |
| 2:31.492 | של |  |
| 2:31.493 | ראשונים |  |
| 2:31.494 | וכו | ' |
| 2:31.495 | ושב |  |
| 2:31.496 | ואל |  |
| 2:31.497 | תעשה |  |
| 2:31.498 | עדיף |  |
| 2:31.499 | דחמירא |  |
| 2:31.500 | סכנתא |  |
| 2:31.501 | דגופא |  |
| 2:31.502 | וכ"ש |  |
| 2:31.503 | דנפש |  |
| 2:31.504 | וכו | ' |
| 2:31.505 | ואין |  |
| 2:31.506 | לנו |  |
| 2:31.507 | לחקור |  |
| 2:31.508 | בנסתרות |  |
| 2:31.509 | כי |  |
| 2:31.510 | אם |  |
| 2:31.511 | לשמוע |  |
| 2:31.512 | בלימודי |  |
| 2:31.513 | פשטי |  |
| 2:31.514 | המקראות |  |
| 2:31.515 | ותוכחת |  |
| 2:31.516 | מוסר |  |
| 2:31.517 | להתנהג |  |
| 2:31.518 | בדרכי |  |
| 2:31.519 | ישרים |  |
| 2:31.520 | כי |  |
| 2:31.521 | ידיעתינו |  |
| 2:31.522 | קצרה |  |
| 2:31.523 | ע"כ |  |
| 2:31.524 | לשון |  |
| 2:31.525 | שמן |  |
| 2:31.526 | המאור |  |
| 2:31.527 | אשר |  |
| 2:31.528 | הראיה |  |
| 2:31.529 | ממנו |  |
| 2:31.530 | להסכים |  |
| 2:31.531 | על |  |
| 2:31.532 | הפרישה |  |
| 2:31.533 | וכו | '. |
| 2:31.534 | ועם |  |
| 2:31.535 | כל |  |
| 2:31.536 | זה |  |
| 2:31.537 | גליתי |  |
| 2:31.538 | לך |  |
| 2:31.539 | דעתי |  |
| 2:31.540 | ולא |  |
| 2:31.541 | תסמוך |  |
| 2:31.542 | עלי |  |
| 2:31.543 | לא |  |
| 2:31.544 | אתה |  |
| 2:31.545 | ולא |  |
| 2:31.546 | אחר |  |
| 2:31.547 | כי |  |
| 2:31.548 | איני |  |
| 2:31.549 | כדאי |  |
| 2:31.550 | להכריע |  |
| 2:31.551 | והנה |  |
| 2:31.552 | אני |  |
| 2:31.553 | מדמה |  |
| 2:31.554 | לימוד |  |
| 2:31.555 | קדש |  |
| 2:31.556 | קדשים |  |
| 2:31.557 | הזה |  |
| 2:31.558 | לנסיעה |  |
| 2:31.559 | לארץ |  |
| 2:31.560 | הקדש |  |
| 2:31.561 | כי |  |
| 2:31.562 | ידוע |  |
| 2:31.563 | מעלתה |  |
| 2:31.564 | כמו |  |
| 2:31.565 | שאמרו |  |
| 2:31.566 | בתלמוד |  |
| 2:31.567 | שהרי |  |
| 2:31.568 | בזולתה |  |
| 2:31.569 | דומה |  |
| 2:31.570 | כמי |  |
| 2:31.571 | שאין |  |
| 2:31.572 | לו |  |
| 2:31.573 | אלוה |  |
| 2:31.574 | ובעובד |  |
| 2:31.575 | ע"ז |  |
| 2:31.576 | והמהלך |  |
| 2:31.577 | ד | ' |
| 2:31.578 | אמות |  |
| 2:31.579 | בא"י |  |
| 2:31.580 | הוא |  |
| 2:31.581 | בן |  |
| 2:31.582 | העה"ב |  |
| 2:31.583 | וכמה |  |
| 2:31.584 | שבחים | . |
| 2:31.585 | ואפילו |  |
| 2:31.586 | הכי |  |
| 2:31.587 | כתבו |  |
| 2:31.588 | התוס | ' |
| 2:31.589 | בשם |  |
| 2:31.590 | רבי |  |
| 2:31.591 | חיים |  |
| 2:31.592 | דעכשיו |  |
| 2:31.593 | אינו |  |
| 2:31.594 | מצוה |  |
| 2:31.595 | לדור |  |
| 2:31.596 | בארץ | . |
| 2:31.597 | כי |  |
| 2:31.598 | יש |  |
| 2:31.599 | כמה |  |
| 2:31.600 | מצות |  |
| 2:31.601 | התלוים |  |
| 2:31.602 | בארץ |  |
| 2:31.603 | שאין |  |
| 2:31.604 | אנו |  |
| 2:31.605 | יכולים |  |
| 2:31.606 | לעמוד |  |
| 2:31.607 | עליהם | . |
| 2:31.608 | ועוד |  |
| 2:31.609 | מפני |  |
| 2:31.610 | סכנת |  |
| 2:31.611 | דרכים | , |
| 2:31.612 | ומפני |  |
| 2:31.613 | העוני |  |
| 2:31.614 | שם |  |
| 2:31.615 | יעבירנו |  |
| 2:31.616 | על |  |
| 2:31.617 | דעתו |  |
| 2:31.618 | ועל |  |
| 2:31.619 | דעת |  |
| 2:31.620 | קונו |  |
| 2:31.621 | שם | . |
| 2:31.622 | ובאם |  |
| 2:31.623 | אין |  |
| 2:31.624 | לפניו |  |
| 2:31.625 | בקי |  |
| 2:31.626 | ושמירה |  |
| 2:31.627 | מעולה |  |
| 2:31.628 | מכל |  |
| 2:31.629 | נזקים |  |
| 2:31.630 | טוב |  |
| 2:31.631 | למנוע |  |
| 2:31.632 | מזה | ! |
| 2:31.633 | הלא |  |
| 2:31.634 | נמי |  |
| 2:31.635 | מי |  |
| 2:31.636 | יכחיש |  |
| 2:31.637 | מעלת |  |
| 2:31.638 | חכמה |  |
| 2:31.639 | זו |  |
| 2:31.640 | הגבה |  |
| 2:31.641 | למעלה |  |
| 2:31.642 | בלי |  |
| 2:31.643 | ערך | . |
| 2:31.644 | בתנאי |  |
| 2:31.645 | לדזכי |  |
| 2:31.646 | לה |  |
| 2:31.647 | ללמדה |  |
| 2:31.648 | מפי |  |
| 2:31.649 | מקובל |  |
| 2:31.650 | שקבלה |  |
| 2:31.651 | ג"כ |  |
| 2:31.652 | מרבו | . |
| 2:31.653 | והן |  |
| 2:31.654 | אמת |  |
| 2:31.655 | שאני |  |
| 2:31.656 | תמיה |  |
| 2:31.657 | על |  |
| 2:31.658 | ספרים |  |
| 2:31.659 | שחוברו |  |
| 2:31.660 | מאחרונים |  |
| 2:31.661 | שהביאו |  |
| 2:31.662 | מאמרים |  |
| 2:31.663 | מהזהר | . |
| 2:31.664 | והמה |  |
| 2:31.665 | מקשים |  |
| 2:31.666 | בהם |  |
| 2:31.667 | ומבארים |  |
| 2:31.668 | כפי |  |
| 2:31.669 | שיקול |  |
| 2:31.670 | דעתם | . |
| 2:31.671 | מי |  |
| 2:31.672 | נתן |  |
| 2:31.673 | רשות |  |
| 2:31.674 | לשכל |  |
| 2:31.675 | אנושי |  |
| 2:31.676 | להמציא |  |
| 2:31.677 | משיקול |  |
| 2:31.678 | דעתו |  |
| 2:31.679 | במציאות |  |
| 2:31.680 | שלש |  |
| 2:31.681 | ראשונות |  |
| 2:31.682 | ושלש |  |
| 2:31.683 | צחצחות | , |
| 2:31.684 | כי |  |
| 2:31.685 | מבואר |  |
| 2:31.686 | בזהר |  |
| 2:31.687 | דאית |  |
| 2:31.688 | מקום |  |
| 2:31.689 | דלא |  |
| 2:31.690 | קימא |  |
| 2:31.691 | לשאלה |  |
| 2:31.692 | וכו | ' |
| 2:31.693 | ואין |  |
| 2:31.694 | ראיה |  |
| 2:31.695 | ממאמר |  |
| 2:31.696 | הזהר |  |
| 2:31.697 | וס | ' |
| 2:31.698 | יצירה |  |
| 2:31.699 | ור"מ |  |
| 2:31.700 | שקבילו |  |
| 2:31.701 | מפי |  |
| 2:31.702 | סבא |  |
| 2:31.703 | דסבין |  |
| 2:31.704 | ואליהו |  |
| 2:31.705 | ומקדישין |  |
| 2:31.706 | עילאין |  |
| 2:31.707 | וכל |  |
| 2:31.708 | הנך |  |
| 2:31.709 | ינוקי |  |
| 2:31.710 | דבזהר |  |
| 2:31.711 | רוח |  |
| 2:31.712 | ה | ' |
| 2:31.713 | דבר |  |
| 2:31.714 | בם |  |
| 2:31.715 | ומלתו |  |
| 2:31.716 | על |  |
| 2:31.717 | לשונם | , |
| 2:31.718 | וגם |  |
| 2:31.719 | האר"י |  |
| 2:31.720 | וגוריו |  |
| 2:31.721 | שקיבלו |  |
| 2:31.722 | ממנו |  |
| 2:31.723 | פה |  |
| 2:31.724 | אל |  |
| 2:31.725 | פה |  |
| 2:31.726 | מפי |  |
| 2:31.727 | עליונים |  |
| 2:31.728 | יצאו |  |
| 2:31.729 | הדברים |  |
| 2:31.730 | ואנן |  |
| 2:31.731 | בלמודינו |  |
| 2:31.732 | באותן |  |
| 2:31.733 | לא |  |
| 2:31.734 | יעלה |  |
| 2:31.735 | בדעתינו |  |
| 2:31.736 | זולת |  |
| 2:31.737 | קריאת |  |
| 2:31.738 | המלות | , |
| 2:31.739 | מה |  |
| 2:31.740 | שאין |  |
| 2:31.741 | כן |  |
| 2:31.742 | אלו |  |
| 2:31.743 | האחרונים |  |
| 2:31.744 | והמבארים |  |
| 2:31.745 | משיקול |  |
| 2:31.746 | הדעת | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

