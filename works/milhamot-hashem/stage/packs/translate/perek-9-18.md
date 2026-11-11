# Translation pack: milhamot-hashem perek-9-18 (Milchamot ha-Shem 9:21–9:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-18`.

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
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- מעמד: "assembly" or "occasion" by context; "the Standing at Sinai" where the
  author names that event.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
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
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
- A title the printed text spells wrongly (ונור אלצלאם): translate the printed
  spelling ("Nur al-Tzalam"); the likelier title (Nur al-Zulam, Me'or ha-Afelah)
  goes in the token note and `apparatus`, not in the English.
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
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


## Milchamot ha-Shem 9:21

Hebrew (Wikisource): קלט) לפיכך כל מי שלבו שלם עם ה' אלהינו שהוא הסבה הראשונה לכל הנמצאים בלי שום שתוף כח אחר, אל ישמע לדברי המסית מחבר הזהר והנמשכים אחריו! כי ה' אלהינו הוא לבדו הנמציא כל הנמצאים ברוב טובו וחסדו, והוא היחיד ומיוחד מכל שאר האחדים, והוא לבדו ראשון שאין ראשית לראשיתו, והוא שהוציאנו ממצרים, והוא שנגלה בסיני ונתן לנו את התורה בקולות וברקים ולפידים ורול שופר, ורעשו כל ההרים ונזדעזע כל העולם לקול נבורתו, ונתן בקולו קול עוז אנכי ה' אלהיך אשר הוציאתיך, לא היה לך אלהים אחרים על פני, וכל עשרה דברים הנאמרים במעמד הנבחר ההוא בפרסום גדול, עד שכל מלכי מזרח ומערב חרדו ותמהו, והלכו אצל בלעם נביאם להודיעם מה זה, כמו שאמרו רז"ל במדרש, ואבותינו כולם ששים רבוא איש רגלי מלבד הזקינים והנשים והטף כולם נצבים בתחתית הההר הר סיני לקבל תורה ומצות, וענו כולם ואמרו כל אשר דבר ה' נעשה ונשמע, ומיראה ופחד גדול מקול ה' בכח, אמרו למשה רבינו דבר אתה עמנו ונשמעה וכו'. וכל זה עשה הקב"ה עמנו כדי להאמין במשה בחירו שתורת אמת נתן לנו, ולא יחליפה ולא יחליפה ולא ימירה וגם אמר את ה' האמרת וכו' וה' האמירך היום להיות לו לעם סגלה, ולא נאמין לשום נביא או קוסם שיפתה אותנו ל עבוד אלוה אחר מבלעדיו, רק אותו לבדו נעבוד שלי שום שתוף דבר אחר עמו, והוא הודיע דרכיו למשה. כי הוא חנון ורחום ארף אפים וכו' ולא קצר אפים ואם כן איך נעזוב את אלהינו ארך אפים ונעבוד לאלוה קצא אפים? וגם אם יעשה נביא השקר אותות ומופתים לעינינו בשמים ובארץ. כמו שנאמר ונתן אליך אות או מופת וכי' ובא האות והמופת אשר דבר אליך לאמר נלכה אחרי אלהים אחרים אשר לא ידעתם (לא בנביאנו ולא מרז"ל) ונעבדם. לא תשמע אל דברי הנביא ההוא וכו' כל שכן וק"ו שלא נשמע לדברי הפילסוף מחבר הזהר, אשר זה מקרוב צמח נן הארץ ולא נודע לו בית אב מי הוא זה מחברו באמת, כשאר ספרי רז"ל הידועים ומפורסמים מחבריהם מיום שחוברו על ידי גדולי ישראל מעתיקי השמועה ועד עתה. וגדול ועצום הוא הק"ו הזה. ומה אם נביא העושה אותות ומופתים לענינינו בשמים ובארץ כמו שפירשו רז"ל ומביאו רש"י ז"ל, אמרה תורה לא תשמע אל דברי הנביא ההוא כי מנסה ה' אלהיכם אתכם? ספר הזהר הבא לנו מיד עכו"ם, ומשמיע גדולות ונוראות משם תנאים ואמוראים ונשמות צדיקים שקדמו לרשב"י כמה דורות. ואמוראים המאוחרים מן רב מאד אחר רשב"י, ואליהו ועתיקא קדישא שנגלה לנו בבית מדרשו (בין חביריו ותלמידיו) וכיוצא בספורים אלו הרשומים בדיו שחור על נייר לבן, אשר לא ראינו בעינינו ולא אבותינו ספרו לנו כל אותן הדברים והמעשיות הרשומים שם! כ"ש שלא נשמע בליו לעבוד שום פרצוף מכל הפרצופים, ובפרט קצר אפים (ז"א) כי אם את אלהינו בהו"ש אשר קבלנו מאבותינו, ותוה"ק ספרה לנו שאין לו דמות הגוף, ואין ראשית לראשיתו, והוא אחד ואין יחוד כיחודו לפי שכל הפרצופים והעגולים כולם יש ראשית לראשיתם, ואין לעבוד רק לצור הכל בלי שום שתוף נברא או נאצל מכל הנבראים והנאצלים, וגם אין ליחד ולקשר עם אחדותו שום אחד מהם, כמאמר רשב"י כל המשתף וכו'.

| id | word | punctuation after |
|---|---|---|
| 9:21.1 | קלט | ) |
| 9:21.2 | לפיכך |  |
| 9:21.3 | כל |  |
| 9:21.4 | מי |  |
| 9:21.5 | שלבו |  |
| 9:21.6 | שלם |  |
| 9:21.7 | עם |  |
| 9:21.8 | ה | ' |
| 9:21.9 | אלהינו |  |
| 9:21.10 | שהוא |  |
| 9:21.11 | הסבה |  |
| 9:21.12 | הראשונה |  |
| 9:21.13 | לכל |  |
| 9:21.14 | הנמצאים |  |
| 9:21.15 | בלי |  |
| 9:21.16 | שום |  |
| 9:21.17 | שתוף |  |
| 9:21.18 | כח |  |
| 9:21.19 | אחר | , |
| 9:21.20 | אל |  |
| 9:21.21 | ישמע |  |
| 9:21.22 | לדברי |  |
| 9:21.23 | המסית |  |
| 9:21.24 | מחבר |  |
| 9:21.25 | הזהר |  |
| 9:21.26 | והנמשכים |  |
| 9:21.27 | אחריו | ! |
| 9:21.28 | כי |  |
| 9:21.29 | ה | ' |
| 9:21.30 | אלהינו |  |
| 9:21.31 | הוא |  |
| 9:21.32 | לבדו |  |
| 9:21.33 | הנמציא |  |
| 9:21.34 | כל |  |
| 9:21.35 | הנמצאים |  |
| 9:21.36 | ברוב |  |
| 9:21.37 | טובו |  |
| 9:21.38 | וחסדו | , |
| 9:21.39 | והוא |  |
| 9:21.40 | היחיד |  |
| 9:21.41 | ומיוחד |  |
| 9:21.42 | מכל |  |
| 9:21.43 | שאר |  |
| 9:21.44 | האחדים | , |
| 9:21.45 | והוא |  |
| 9:21.46 | לבדו |  |
| 9:21.47 | ראשון |  |
| 9:21.48 | שאין |  |
| 9:21.49 | ראשית |  |
| 9:21.50 | לראשיתו | , |
| 9:21.51 | והוא |  |
| 9:21.52 | שהוציאנו |  |
| 9:21.53 | ממצרים | , |
| 9:21.54 | והוא |  |
| 9:21.55 | שנגלה |  |
| 9:21.56 | בסיני |  |
| 9:21.57 | ונתן |  |
| 9:21.58 | לנו |  |
| 9:21.59 | את |  |
| 9:21.60 | התורה |  |
| 9:21.61 | בקולות |  |
| 9:21.62 | וברקים |  |
| 9:21.63 | ולפידים |  |
| 9:21.64 | ורול |  |
| 9:21.65 | שופר | , |
| 9:21.66 | ורעשו |  |
| 9:21.67 | כל |  |
| 9:21.68 | ההרים |  |
| 9:21.69 | ונזדעזע |  |
| 9:21.70 | כל |  |
| 9:21.71 | העולם |  |
| 9:21.72 | לקול |  |
| 9:21.73 | נבורתו | , |
| 9:21.74 | ונתן |  |
| 9:21.75 | בקולו |  |
| 9:21.76 | קול |  |
| 9:21.77 | עוז |  |
| 9:21.78 | אנכי |  |
| 9:21.79 | ה | ' |
| 9:21.80 | אלהיך |  |
| 9:21.81 | אשר |  |
| 9:21.82 | הוציאתיך | , |
| 9:21.83 | לא |  |
| 9:21.84 | היה |  |
| 9:21.85 | לך |  |
| 9:21.86 | אלהים |  |
| 9:21.87 | אחרים |  |
| 9:21.88 | על |  |
| 9:21.89 | פני | , |
| 9:21.90 | וכל |  |
| 9:21.91 | עשרה |  |
| 9:21.92 | דברים |  |
| 9:21.93 | הנאמרים |  |
| 9:21.94 | במעמד |  |
| 9:21.95 | הנבחר |  |
| 9:21.96 | ההוא |  |
| 9:21.97 | בפרסום |  |
| 9:21.98 | גדול | , |
| 9:21.99 | עד |  |
| 9:21.100 | שכל |  |
| 9:21.101 | מלכי |  |
| 9:21.102 | מזרח |  |
| 9:21.103 | ומערב |  |
| 9:21.104 | חרדו |  |
| 9:21.105 | ותמהו | , |
| 9:21.106 | והלכו |  |
| 9:21.107 | אצל |  |
| 9:21.108 | בלעם |  |
| 9:21.109 | נביאם |  |
| 9:21.110 | להודיעם |  |
| 9:21.111 | מה |  |
| 9:21.112 | זה | , |
| 9:21.113 | כמו |  |
| 9:21.114 | שאמרו |  |
| 9:21.115 | רז"ל |  |
| 9:21.116 | במדרש | , |
| 9:21.117 | ואבותינו |  |
| 9:21.118 | כולם |  |
| 9:21.119 | ששים |  |
| 9:21.120 | רבוא |  |
| 9:21.121 | איש |  |
| 9:21.122 | רגלי |  |
| 9:21.123 | מלבד |  |
| 9:21.124 | הזקינים |  |
| 9:21.125 | והנשים |  |
| 9:21.126 | והטף |  |
| 9:21.127 | כולם |  |
| 9:21.128 | נצבים |  |
| 9:21.129 | בתחתית |  |
| 9:21.130 | הההר |  |
| 9:21.131 | הר |  |
| 9:21.132 | סיני |  |
| 9:21.133 | לקבל |  |
| 9:21.134 | תורה |  |
| 9:21.135 | ומצות | , |
| 9:21.136 | וענו |  |
| 9:21.137 | כולם |  |
| 9:21.138 | ואמרו |  |
| 9:21.139 | כל |  |
| 9:21.140 | אשר |  |
| 9:21.141 | דבר |  |
| 9:21.142 | ה | ' |
| 9:21.143 | נעשה |  |
| 9:21.144 | ונשמע | , |
| 9:21.145 | ומיראה |  |
| 9:21.146 | ופחד |  |
| 9:21.147 | גדול |  |
| 9:21.148 | מקול |  |
| 9:21.149 | ה | ' |
| 9:21.150 | בכח | , |
| 9:21.151 | אמרו |  |
| 9:21.152 | למשה |  |
| 9:21.153 | רבינו |  |
| 9:21.154 | דבר |  |
| 9:21.155 | אתה |  |
| 9:21.156 | עמנו |  |
| 9:21.157 | ונשמעה |  |
| 9:21.158 | וכו | '. |
| 9:21.159 | וכל |  |
| 9:21.160 | זה |  |
| 9:21.161 | עשה |  |
| 9:21.162 | הקב"ה |  |
| 9:21.163 | עמנו |  |
| 9:21.164 | כדי |  |
| 9:21.165 | להאמין |  |
| 9:21.166 | במשה |  |
| 9:21.167 | בחירו |  |
| 9:21.168 | שתורת |  |
| 9:21.169 | אמת |  |
| 9:21.170 | נתן |  |
| 9:21.171 | לנו | , |
| 9:21.172 | ולא |  |
| 9:21.173 | יחליפה |  |
| 9:21.174 | ולא |  |
| 9:21.175 | יחליפה |  |
| 9:21.176 | ולא |  |
| 9:21.177 | ימירה |  |
| 9:21.178 | וגם |  |
| 9:21.179 | אמר |  |
| 9:21.180 | את |  |
| 9:21.181 | ה | ' |
| 9:21.182 | האמרת |  |
| 9:21.183 | וכו | ' |
| 9:21.184 | וה | ' |
| 9:21.185 | האמירך |  |
| 9:21.186 | היום |  |
| 9:21.187 | להיות |  |
| 9:21.188 | לו |  |
| 9:21.189 | לעם |  |
| 9:21.190 | סגלה | , |
| 9:21.191 | ולא |  |
| 9:21.192 | נאמין |  |
| 9:21.193 | לשום |  |
| 9:21.194 | נביא |  |
| 9:21.195 | או |  |
| 9:21.196 | קוסם |  |
| 9:21.197 | שיפתה |  |
| 9:21.198 | אותנו |  |
| 9:21.199 | ל |  |
| 9:21.200 | עבוד |  |
| 9:21.201 | אלוה |  |
| 9:21.202 | אחר |  |
| 9:21.203 | מבלעדיו | , |
| 9:21.204 | רק |  |
| 9:21.205 | אותו |  |
| 9:21.206 | לבדו |  |
| 9:21.207 | נעבוד |  |
| 9:21.208 | שלי |  |
| 9:21.209 | שום |  |
| 9:21.210 | שתוף |  |
| 9:21.211 | דבר |  |
| 9:21.212 | אחר |  |
| 9:21.213 | עמו | , |
| 9:21.214 | והוא |  |
| 9:21.215 | הודיע |  |
| 9:21.216 | דרכיו |  |
| 9:21.217 | למשה | . |
| 9:21.218 | כי |  |
| 9:21.219 | הוא |  |
| 9:21.220 | חנון |  |
| 9:21.221 | ורחום |  |
| 9:21.222 | ארף |  |
| 9:21.223 | אפים |  |
| 9:21.224 | וכו | ' |
| 9:21.225 | ולא |  |
| 9:21.226 | קצר |  |
| 9:21.227 | אפים |  |
| 9:21.228 | ואם |  |
| 9:21.229 | כן |  |
| 9:21.230 | איך |  |
| 9:21.231 | נעזוב |  |
| 9:21.232 | את |  |
| 9:21.233 | אלהינו |  |
| 9:21.234 | ארך |  |
| 9:21.235 | אפים |  |
| 9:21.236 | ונעבוד |  |
| 9:21.237 | לאלוה |  |
| 9:21.238 | קצא |  |
| 9:21.239 | אפים | ? |
| 9:21.240 | וגם |  |
| 9:21.241 | אם |  |
| 9:21.242 | יעשה |  |
| 9:21.243 | נביא |  |
| 9:21.244 | השקר |  |
| 9:21.245 | אותות |  |
| 9:21.246 | ומופתים |  |
| 9:21.247 | לעינינו |  |
| 9:21.248 | בשמים |  |
| 9:21.249 | ובארץ | . |
| 9:21.250 | כמו |  |
| 9:21.251 | שנאמר |  |
| 9:21.252 | ונתן |  |
| 9:21.253 | אליך |  |
| 9:21.254 | אות |  |
| 9:21.255 | או |  |
| 9:21.256 | מופת |  |
| 9:21.257 | וכי | ' |
| 9:21.258 | ובא |  |
| 9:21.259 | האות |  |
| 9:21.260 | והמופת |  |
| 9:21.261 | אשר |  |
| 9:21.262 | דבר |  |
| 9:21.263 | אליך |  |
| 9:21.264 | לאמר |  |
| 9:21.265 | נלכה |  |
| 9:21.266 | אחרי |  |
| 9:21.267 | אלהים |  |
| 9:21.268 | אחרים |  |
| 9:21.269 | אשר |  |
| 9:21.270 | לא |  |
| 9:21.271 | ידעתם |  |
| 9:21.272 | לא |  |
| 9:21.273 | בנביאנו |  |
| 9:21.274 | ולא |  |
| 9:21.275 | מרז"ל | ) |
| 9:21.276 | ונעבדם | . |
| 9:21.277 | לא |  |
| 9:21.278 | תשמע |  |
| 9:21.279 | אל |  |
| 9:21.280 | דברי |  |
| 9:21.281 | הנביא |  |
| 9:21.282 | ההוא |  |
| 9:21.283 | וכו | ' |
| 9:21.284 | כל |  |
| 9:21.285 | שכן |  |
| 9:21.286 | וק"ו |  |
| 9:21.287 | שלא |  |
| 9:21.288 | נשמע |  |
| 9:21.289 | לדברי |  |
| 9:21.290 | הפילסוף |  |
| 9:21.291 | מחבר |  |
| 9:21.292 | הזהר | , |
| 9:21.293 | אשר |  |
| 9:21.294 | זה |  |
| 9:21.295 | מקרוב |  |
| 9:21.296 | צמח |  |
| 9:21.297 | נן |  |
| 9:21.298 | הארץ |  |
| 9:21.299 | ולא |  |
| 9:21.300 | נודע |  |
| 9:21.301 | לו |  |
| 9:21.302 | בית |  |
| 9:21.303 | אב |  |
| 9:21.304 | מי |  |
| 9:21.305 | הוא |  |
| 9:21.306 | זה |  |
| 9:21.307 | מחברו |  |
| 9:21.308 | באמת | , |
| 9:21.309 | כשאר |  |
| 9:21.310 | ספרי |  |
| 9:21.311 | רז"ל |  |
| 9:21.312 | הידועים |  |
| 9:21.313 | ומפורסמים |  |
| 9:21.314 | מחבריהם |  |
| 9:21.315 | מיום |  |
| 9:21.316 | שחוברו |  |
| 9:21.317 | על |  |
| 9:21.318 | ידי |  |
| 9:21.319 | גדולי |  |
| 9:21.320 | ישראל |  |
| 9:21.321 | מעתיקי |  |
| 9:21.322 | השמועה |  |
| 9:21.323 | ועד |  |
| 9:21.324 | עתה | . |
| 9:21.325 | וגדול |  |
| 9:21.326 | ועצום |  |
| 9:21.327 | הוא |  |
| 9:21.328 | הק"ו |  |
| 9:21.329 | הזה | . |
| 9:21.330 | ומה |  |
| 9:21.331 | אם |  |
| 9:21.332 | נביא |  |
| 9:21.333 | העושה |  |
| 9:21.334 | אותות |  |
| 9:21.335 | ומופתים |  |
| 9:21.336 | לענינינו |  |
| 9:21.337 | בשמים |  |
| 9:21.338 | ובארץ |  |
| 9:21.339 | כמו |  |
| 9:21.340 | שפירשו |  |
| 9:21.341 | רז"ל |  |
| 9:21.342 | ומביאו |  |
| 9:21.343 | רש"י |  |
| 9:21.344 | ז"ל | , |
| 9:21.345 | אמרה |  |
| 9:21.346 | תורה |  |
| 9:21.347 | לא |  |
| 9:21.348 | תשמע |  |
| 9:21.349 | אל |  |
| 9:21.350 | דברי |  |
| 9:21.351 | הנביא |  |
| 9:21.352 | ההוא |  |
| 9:21.353 | כי |  |
| 9:21.354 | מנסה |  |
| 9:21.355 | ה | ' |
| 9:21.356 | אלהיכם |  |
| 9:21.357 | אתכם | ? |
| 9:21.358 | ספר |  |
| 9:21.359 | הזהר |  |
| 9:21.360 | הבא |  |
| 9:21.361 | לנו |  |
| 9:21.362 | מיד |  |
| 9:21.363 | עכו"ם | , |
| 9:21.364 | ומשמיע |  |
| 9:21.365 | גדולות |  |
| 9:21.366 | ונוראות |  |
| 9:21.367 | משם |  |
| 9:21.368 | תנאים |  |
| 9:21.369 | ואמוראים |  |
| 9:21.370 | ונשמות |  |
| 9:21.371 | צדיקים |  |
| 9:21.372 | שקדמו |  |
| 9:21.373 | לרשב"י |  |
| 9:21.374 | כמה |  |
| 9:21.375 | דורות | . |
| 9:21.376 | ואמוראים |  |
| 9:21.377 | המאוחרים |  |
| 9:21.378 | מן |  |
| 9:21.379 | רב |  |
| 9:21.380 | מאד |  |
| 9:21.381 | אחר |  |
| 9:21.382 | רשב"י | , |
| 9:21.383 | ואליהו |  |
| 9:21.384 | ועתיקא |  |
| 9:21.385 | קדישא |  |
| 9:21.386 | שנגלה |  |
| 9:21.387 | לנו |  |
| 9:21.388 | בבית |  |
| 9:21.389 | מדרשו |  |
| 9:21.390 | בין |  |
| 9:21.391 | חביריו |  |
| 9:21.392 | ותלמידיו | ) |
| 9:21.393 | וכיוצא |  |
| 9:21.394 | בספורים |  |
| 9:21.395 | אלו |  |
| 9:21.396 | הרשומים |  |
| 9:21.397 | בדיו |  |
| 9:21.398 | שחור |  |
| 9:21.399 | על |  |
| 9:21.400 | נייר |  |
| 9:21.401 | לבן | , |
| 9:21.402 | אשר |  |
| 9:21.403 | לא |  |
| 9:21.404 | ראינו |  |
| 9:21.405 | בעינינו |  |
| 9:21.406 | ולא |  |
| 9:21.407 | אבותינו |  |
| 9:21.408 | ספרו |  |
| 9:21.409 | לנו |  |
| 9:21.410 | כל |  |
| 9:21.411 | אותן |  |
| 9:21.412 | הדברים |  |
| 9:21.413 | והמעשיות |  |
| 9:21.414 | הרשומים |  |
| 9:21.415 | שם | ! |
| 9:21.416 | כ"ש |  |
| 9:21.417 | שלא |  |
| 9:21.418 | נשמע |  |
| 9:21.419 | בליו |  |
| 9:21.420 | לעבוד |  |
| 9:21.421 | שום |  |
| 9:21.422 | פרצוף |  |
| 9:21.423 | מכל |  |
| 9:21.424 | הפרצופים | , |
| 9:21.425 | ובפרט |  |
| 9:21.426 | קצר |  |
| 9:21.427 | אפים |  |
| 9:21.428 | ז"א | ) |
| 9:21.429 | כי |  |
| 9:21.430 | אם |  |
| 9:21.431 | את |  |
| 9:21.432 | אלהינו |  |
| 9:21.433 | בהו"ש |  |
| 9:21.434 | אשר |  |
| 9:21.435 | קבלנו |  |
| 9:21.436 | מאבותינו | , |
| 9:21.437 | ותוה"ק |  |
| 9:21.438 | ספרה |  |
| 9:21.439 | לנו |  |
| 9:21.440 | שאין |  |
| 9:21.441 | לו |  |
| 9:21.442 | דמות |  |
| 9:21.443 | הגוף | , |
| 9:21.444 | ואין |  |
| 9:21.445 | ראשית |  |
| 9:21.446 | לראשיתו | , |
| 9:21.447 | והוא |  |
| 9:21.448 | אחד |  |
| 9:21.449 | ואין |  |
| 9:21.450 | יחוד |  |
| 9:21.451 | כיחודו |  |
| 9:21.452 | לפי |  |
| 9:21.453 | שכל |  |
| 9:21.454 | הפרצופים |  |
| 9:21.455 | והעגולים |  |
| 9:21.456 | כולם |  |
| 9:21.457 | יש |  |
| 9:21.458 | ראשית |  |
| 9:21.459 | לראשיתם | , |
| 9:21.460 | ואין |  |
| 9:21.461 | לעבוד |  |
| 9:21.462 | רק |  |
| 9:21.463 | לצור |  |
| 9:21.464 | הכל |  |
| 9:21.465 | בלי |  |
| 9:21.466 | שום |  |
| 9:21.467 | שתוף |  |
| 9:21.468 | נברא |  |
| 9:21.469 | או |  |
| 9:21.470 | נאצל |  |
| 9:21.471 | מכל |  |
| 9:21.472 | הנבראים |  |
| 9:21.473 | והנאצלים | , |
| 9:21.474 | וגם |  |
| 9:21.475 | אין |  |
| 9:21.476 | ליחד |  |
| 9:21.477 | ולקשר |  |
| 9:21.478 | עם |  |
| 9:21.479 | אחדותו |  |
| 9:21.480 | שום |  |
| 9:21.481 | אחד |  |
| 9:21.482 | מהם | , |
| 9:21.483 | כמאמר |  |
| 9:21.484 | רשב"י |  |
| 9:21.485 | כל |  |
| 9:21.486 | המשתף |  |
| 9:21.487 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

