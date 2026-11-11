# Translation pack: milhamot-hashem perek-6-25 (Milchamot ha-Shem 6:38–6:38)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-25.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-25`.

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
- **שיתוף / שתוף / לשתף** (there is no שותפות in the text): context.
  - Associating another with God: fixed, "associate / association". "כל
    המשתף שם שמים ודבר אחר נעקר מן העולם" (Sanh. 63a; 1:7, 5:29, 5:32) →
    "whoever associates the name of Heaven with something else is uprooted
    from the world"; "בשיתוף שאר הפרצופים" (5:1) → "in association with the
    other partzufim"; "בלי חיבור וצירוף ושיתוף" (5:21) → "without
    composition, conjunction or association".
  - The Maimonidean term for equivocal words: "בשתוף השם בלבד" (5:10) → "by
    equivocation of the term only"; "מושאלים או משותפים" (5:18) →
    "borrowed or equivocal". Never "association" here. The note gives
    Guide I.
  - צירוף: "conjunction" when used of God; "permutation" in letter mysticism
    ("י"ב צירופי הוי'ה", 2:33 → "the twelve permutations of YHVH").
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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- ר"ל with a named authority as subject: "means to say".
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
- דרך ארץ: "worldly occupation", as in the project's Avot.
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
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


## Milchamot ha-Shem 6:38

Hebrew (Wikisource): עז) דבר הלמד מדברי הר' הנז' כי דבר הגורם לכל אלה השגיאות ששגגו רבותינו בעלי הקבלה החדשה להאמין באלהים רבים, הוא כי האמינו לדברי נביא השקר המסית מחבר ספר הזהר. ולא נתנו אל לבם אזהרת תורתינו הקדושה באמרה ונשמרתם מאד לנפשותיכם כי לא ראיתם כל תמונה. ולא לאזהרת הנביא שאמר ואל מי תדמיון אל. ואל מי תדמיוני ואשוה. ובקשו חשבונות רבים בכמה מקראות שנאמרו על דרך ההשאלה והרחבת לשון בדרך המליצה לפרשם על דרך הנחתם הראשונה. כגון בצלמינו כדמותינו ושאר לשונות כנז' לעיל שלא על דרך שפירשו רז"ל והגאונים החכמים. עד שגרם להם שבאו לכלל טעות גדולה לעבוד לנברא שהוא קצר אפים (זעיר אנפין) או מלכא קדישא דכל קדישין לדעת עז לאלהים. והמציאו כמה רמזים ורמזי רמזים שאין להם שחר או במספר האותיות, או בחילופיהן וחילופי חלופיהן ומלואיהן. ושבוש סדריהן ובלבולן. כמו מה תצעק אלי מאצ"ל העתי"ק. כדי למצוא רמז שיתפלל לז"א ויכלול עמו פרצוף עתיק יומין לשתף יחד כל החמשה פרצופין עם העתיק. כמו שכ' כסא אליהו דבר שאסרה תורה לכלול ולשתף עם אלוהינו שום דבר. כמו שאמר רשב"י ע"ה כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתי לה' לבדו. וכל זה נמשך מחמת מיעוט התבוננות בדברי רז"ל. וגם בעבור שנביא השקר מחבר הזהר שהמציא רעיונות אלו ילדי נכרים. תלה אותם ברשב"י. התנא החסיד בשקר. והם לתומם לא חשו לבדוק אחריו. והאמינו באמונת אומן שמפי רשב"י וחביריו יצאו כבושים. וכבר ביארתי לעיל שחלילה לשום תנא או אמורא לומר כן. עוד כתב הרב שבילי אמונה למעלה מזה המאמר, כי הבורא ית' לא יאויך. ר"ל שאינו בעל איכות, ומהנמנע להיות מציאות איכות, אלא במאוייך, מפני שהאיכות מקרה, והמקרה אין לו עמידה אלא בעצם. ולכן הבורא ית' אינו נושא האיכות. גם בעל האיכות הוא אשר לו התאר בעצמו ולו ית' אין לו תואר וכו' ועיין במורה נבוכים פרק נ"ב לראשון.

| id | word | punctuation after |
|---|---|---|
| 6:38.1 | עז | ) |
| 6:38.2 | דבר |  |
| 6:38.3 | הלמד |  |
| 6:38.4 | מדברי |  |
| 6:38.5 | הר | ' |
| 6:38.6 | הנז | ' |
| 6:38.7 | כי |  |
| 6:38.8 | דבר |  |
| 6:38.9 | הגורם |  |
| 6:38.10 | לכל |  |
| 6:38.11 | אלה |  |
| 6:38.12 | השגיאות |  |
| 6:38.13 | ששגגו |  |
| 6:38.14 | רבותינו |  |
| 6:38.15 | בעלי |  |
| 6:38.16 | הקבלה |  |
| 6:38.17 | החדשה |  |
| 6:38.18 | להאמין |  |
| 6:38.19 | באלהים |  |
| 6:38.20 | רבים | , |
| 6:38.21 | הוא |  |
| 6:38.22 | כי |  |
| 6:38.23 | האמינו |  |
| 6:38.24 | לדברי |  |
| 6:38.25 | נביא |  |
| 6:38.26 | השקר |  |
| 6:38.27 | המסית |  |
| 6:38.28 | מחבר |  |
| 6:38.29 | ספר |  |
| 6:38.30 | הזהר | . |
| 6:38.31 | ולא |  |
| 6:38.32 | נתנו |  |
| 6:38.33 | אל |  |
| 6:38.34 | לבם |  |
| 6:38.35 | אזהרת |  |
| 6:38.36 | תורתינו |  |
| 6:38.37 | הקדושה |  |
| 6:38.38 | באמרה |  |
| 6:38.39 | ונשמרתם |  |
| 6:38.40 | מאד |  |
| 6:38.41 | לנפשותיכם |  |
| 6:38.42 | כי |  |
| 6:38.43 | לא |  |
| 6:38.44 | ראיתם |  |
| 6:38.45 | כל |  |
| 6:38.46 | תמונה | . |
| 6:38.47 | ולא |  |
| 6:38.48 | לאזהרת |  |
| 6:38.49 | הנביא |  |
| 6:38.50 | שאמר |  |
| 6:38.51 | ואל |  |
| 6:38.52 | מי |  |
| 6:38.53 | תדמיון |  |
| 6:38.54 | אל | . |
| 6:38.55 | ואל |  |
| 6:38.56 | מי |  |
| 6:38.57 | תדמיוני |  |
| 6:38.58 | ואשוה | . |
| 6:38.59 | ובקשו |  |
| 6:38.60 | חשבונות |  |
| 6:38.61 | רבים |  |
| 6:38.62 | בכמה |  |
| 6:38.63 | מקראות |  |
| 6:38.64 | שנאמרו |  |
| 6:38.65 | על |  |
| 6:38.66 | דרך |  |
| 6:38.67 | ההשאלה |  |
| 6:38.68 | והרחבת |  |
| 6:38.69 | לשון |  |
| 6:38.70 | בדרך |  |
| 6:38.71 | המליצה |  |
| 6:38.72 | לפרשם |  |
| 6:38.73 | על |  |
| 6:38.74 | דרך |  |
| 6:38.75 | הנחתם |  |
| 6:38.76 | הראשונה | . |
| 6:38.77 | כגון |  |
| 6:38.78 | בצלמינו |  |
| 6:38.79 | כדמותינו |  |
| 6:38.80 | ושאר |  |
| 6:38.81 | לשונות |  |
| 6:38.82 | כנז | ' |
| 6:38.83 | לעיל |  |
| 6:38.84 | שלא |  |
| 6:38.85 | על |  |
| 6:38.86 | דרך |  |
| 6:38.87 | שפירשו |  |
| 6:38.88 | רז"ל |  |
| 6:38.89 | והגאונים |  |
| 6:38.90 | החכמים | . |
| 6:38.91 | עד |  |
| 6:38.92 | שגרם |  |
| 6:38.93 | להם |  |
| 6:38.94 | שבאו |  |
| 6:38.95 | לכלל |  |
| 6:38.96 | טעות |  |
| 6:38.97 | גדולה |  |
| 6:38.98 | לעבוד |  |
| 6:38.99 | לנברא |  |
| 6:38.100 | שהוא |  |
| 6:38.101 | קצר |  |
| 6:38.102 | אפים |  |
| 6:38.103 | זעיר |  |
| 6:38.104 | אנפין | ) |
| 6:38.105 | או |  |
| 6:38.106 | מלכא |  |
| 6:38.107 | קדישא |  |
| 6:38.108 | דכל |  |
| 6:38.109 | קדישין |  |
| 6:38.110 | לדעת |  |
| 6:38.111 | עז |  |
| 6:38.112 | לאלהים | . |
| 6:38.113 | והמציאו |  |
| 6:38.114 | כמה |  |
| 6:38.115 | רמזים |  |
| 6:38.116 | ורמזי |  |
| 6:38.117 | רמזים |  |
| 6:38.118 | שאין |  |
| 6:38.119 | להם |  |
| 6:38.120 | שחר |  |
| 6:38.121 | או |  |
| 6:38.122 | במספר |  |
| 6:38.123 | האותיות | , |
| 6:38.124 | או |  |
| 6:38.125 | בחילופיהן |  |
| 6:38.126 | וחילופי |  |
| 6:38.127 | חלופיהן |  |
| 6:38.128 | ומלואיהן | . |
| 6:38.129 | ושבוש |  |
| 6:38.130 | סדריהן |  |
| 6:38.131 | ובלבולן | . |
| 6:38.132 | כמו |  |
| 6:38.133 | מה |  |
| 6:38.134 | תצעק |  |
| 6:38.135 | אלי |  |
| 6:38.136 | מאצ"ל |  |
| 6:38.137 | העתי"ק | . |
| 6:38.138 | כדי |  |
| 6:38.139 | למצוא |  |
| 6:38.140 | רמז |  |
| 6:38.141 | שיתפלל |  |
| 6:38.142 | לז"א |  |
| 6:38.143 | ויכלול |  |
| 6:38.144 | עמו |  |
| 6:38.145 | פרצוף |  |
| 6:38.146 | עתיק |  |
| 6:38.147 | יומין |  |
| 6:38.148 | לשתף |  |
| 6:38.149 | יחד |  |
| 6:38.150 | כל |  |
| 6:38.151 | החמשה |  |
| 6:38.152 | פרצופין |  |
| 6:38.153 | עם |  |
| 6:38.154 | העתיק | . |
| 6:38.155 | כמו |  |
| 6:38.156 | שכ | ' |
| 6:38.157 | כסא |  |
| 6:38.158 | אליהו |  |
| 6:38.159 | דבר |  |
| 6:38.160 | שאסרה |  |
| 6:38.161 | תורה |  |
| 6:38.162 | לכלול |  |
| 6:38.163 | ולשתף |  |
| 6:38.164 | עם |  |
| 6:38.165 | אלוהינו |  |
| 6:38.166 | שום |  |
| 6:38.167 | דבר | . |
| 6:38.168 | כמו |  |
| 6:38.169 | שאמר |  |
| 6:38.170 | רשב"י |  |
| 6:38.171 | ע"ה |  |
| 6:38.172 | כל |  |
| 6:38.173 | המשתף |  |
| 6:38.174 | שם |  |
| 6:38.175 | שמים |  |
| 6:38.176 | ודבר |  |
| 6:38.177 | אחר |  |
| 6:38.178 | נעקר |  |
| 6:38.179 | מן |  |
| 6:38.180 | העולם |  |
| 6:38.181 | שנ | ' |
| 6:38.182 | בלתי |  |
| 6:38.183 | לה | ' |
| 6:38.184 | לבדו | . |
| 6:38.185 | וכל |  |
| 6:38.186 | זה |  |
| 6:38.187 | נמשך |  |
| 6:38.188 | מחמת |  |
| 6:38.189 | מיעוט |  |
| 6:38.190 | התבוננות |  |
| 6:38.191 | בדברי |  |
| 6:38.192 | רז"ל | . |
| 6:38.193 | וגם |  |
| 6:38.194 | בעבור |  |
| 6:38.195 | שנביא |  |
| 6:38.196 | השקר |  |
| 6:38.197 | מחבר |  |
| 6:38.198 | הזהר |  |
| 6:38.199 | שהמציא |  |
| 6:38.200 | רעיונות |  |
| 6:38.201 | אלו |  |
| 6:38.202 | ילדי |  |
| 6:38.203 | נכרים | . |
| 6:38.204 | תלה |  |
| 6:38.205 | אותם |  |
| 6:38.206 | ברשב"י | . |
| 6:38.207 | התנא |  |
| 6:38.208 | החסיד |  |
| 6:38.209 | בשקר | . |
| 6:38.210 | והם |  |
| 6:38.211 | לתומם |  |
| 6:38.212 | לא |  |
| 6:38.213 | חשו |  |
| 6:38.214 | לבדוק |  |
| 6:38.215 | אחריו | . |
| 6:38.216 | והאמינו |  |
| 6:38.217 | באמונת |  |
| 6:38.218 | אומן |  |
| 6:38.219 | שמפי |  |
| 6:38.220 | רשב"י |  |
| 6:38.221 | וחביריו |  |
| 6:38.222 | יצאו |  |
| 6:38.223 | כבושים | . |
| 6:38.224 | וכבר |  |
| 6:38.225 | ביארתי |  |
| 6:38.226 | לעיל |  |
| 6:38.227 | שחלילה |  |
| 6:38.228 | לשום |  |
| 6:38.229 | תנא |  |
| 6:38.230 | או |  |
| 6:38.231 | אמורא |  |
| 6:38.232 | לומר |  |
| 6:38.233 | כן | . |
| 6:38.234 | עוד |  |
| 6:38.235 | כתב |  |
| 6:38.236 | הרב |  |
| 6:38.237 | שבילי |  |
| 6:38.238 | אמונה |  |
| 6:38.239 | למעלה |  |
| 6:38.240 | מזה |  |
| 6:38.241 | המאמר | , |
| 6:38.242 | כי |  |
| 6:38.243 | הבורא |  |
| 6:38.244 | ית | ' |
| 6:38.245 | לא |  |
| 6:38.246 | יאויך | . |
| 6:38.247 | ר"ל |  |
| 6:38.248 | שאינו |  |
| 6:38.249 | בעל |  |
| 6:38.250 | איכות | , |
| 6:38.251 | ומהנמנע |  |
| 6:38.252 | להיות |  |
| 6:38.253 | מציאות |  |
| 6:38.254 | איכות | , |
| 6:38.255 | אלא |  |
| 6:38.256 | במאוייך | , |
| 6:38.257 | מפני |  |
| 6:38.258 | שהאיכות |  |
| 6:38.259 | מקרה | , |
| 6:38.260 | והמקרה |  |
| 6:38.261 | אין |  |
| 6:38.262 | לו |  |
| 6:38.263 | עמידה |  |
| 6:38.264 | אלא |  |
| 6:38.265 | בעצם | . |
| 6:38.266 | ולכן |  |
| 6:38.267 | הבורא |  |
| 6:38.268 | ית | ' |
| 6:38.269 | אינו |  |
| 6:38.270 | נושא |  |
| 6:38.271 | האיכות | . |
| 6:38.272 | גם |  |
| 6:38.273 | בעל |  |
| 6:38.274 | האיכות |  |
| 6:38.275 | הוא |  |
| 6:38.276 | אשר |  |
| 6:38.277 | לו |  |
| 6:38.278 | התאר |  |
| 6:38.279 | בעצמו |  |
| 6:38.280 | ולו |  |
| 6:38.281 | ית | ' |
| 6:38.282 | אין |  |
| 6:38.283 | לו |  |
| 6:38.284 | תואר |  |
| 6:38.285 | וכו | ' |
| 6:38.286 | ועיין |  |
| 6:38.287 | במורה |  |
| 6:38.288 | נבוכים |  |
| 6:38.289 | פרק |  |
| 6:38.290 | נ"ב |  |
| 6:38.291 | לראשון | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

