# Translation pack: milhamot-hashem perek-5-14 (Milchamot ha-Shem 5:21–5:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-14`.

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
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
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


## Milchamot ha-Shem 5:21

Hebrew (Wikisource): מט) ומעתה אשובה ארעה בשושנים דברי רז"ל בגמרא ובמדרשים המפורסמים שממקום קדוש יצאו ואין בהם שום ספק, אשמור דבריהם היקרים מפז ויהלום באחדותו יתברך אחדות גמור בלי חיבור וצירוף ושיתוף שום נברא או נאצל וידיעת דרכיו ית' על פי התורה והמצוה אשר בעבודה בחר באבות שנאמר (בראשית י"ח י"ט): "כי ידעתיו למען אשר יצוה את בניו ואת ביתו אחריו ושמרו דרך השם" וכו' וכן עשה אברהם אבינו וציוה לבניו יצחק וישמעאל אמונת היחוד שהוא יתברך אחד אמת, ואין אחד אמת זולתו, ושהוא קדמון לכל נברא, ראשון ואין ראשית לראשיתו!

| id | word | punctuation after |
|---|---|---|
| 5:21.1 | מט | ) |
| 5:21.2 | ומעתה |  |
| 5:21.3 | אשובה |  |
| 5:21.4 | ארעה |  |
| 5:21.5 | בשושנים |  |
| 5:21.6 | דברי |  |
| 5:21.7 | רז"ל |  |
| 5:21.8 | בגמרא |  |
| 5:21.9 | ובמדרשים |  |
| 5:21.10 | המפורסמים |  |
| 5:21.11 | שממקום |  |
| 5:21.12 | קדוש |  |
| 5:21.13 | יצאו |  |
| 5:21.14 | ואין |  |
| 5:21.15 | בהם |  |
| 5:21.16 | שום |  |
| 5:21.17 | ספק | , |
| 5:21.18 | אשמור |  |
| 5:21.19 | דבריהם |  |
| 5:21.20 | היקרים |  |
| 5:21.21 | מפז |  |
| 5:21.22 | ויהלום |  |
| 5:21.23 | באחדותו |  |
| 5:21.24 | יתברך |  |
| 5:21.25 | אחדות |  |
| 5:21.26 | גמור |  |
| 5:21.27 | בלי |  |
| 5:21.28 | חיבור |  |
| 5:21.29 | וצירוף |  |
| 5:21.30 | ושיתוף |  |
| 5:21.31 | שום |  |
| 5:21.32 | נברא |  |
| 5:21.33 | או |  |
| 5:21.34 | נאצל |  |
| 5:21.35 | וידיעת |  |
| 5:21.36 | דרכיו |  |
| 5:21.37 | ית | ' |
| 5:21.38 | על |  |
| 5:21.39 | פי |  |
| 5:21.40 | התורה |  |
| 5:21.41 | והמצוה |  |
| 5:21.42 | אשר |  |
| 5:21.43 | בעבודה |  |
| 5:21.44 | בחר |  |
| 5:21.45 | באבות |  |
| 5:21.46 | שנאמר |  |
| 5:21.47 | בראשית |  |
| 5:21.48 | י"ח |  |
| 5:21.49 | י"ט | ): |
| 5:21.50 | כי |  |
| 5:21.51 | ידעתיו |  |
| 5:21.52 | למען |  |
| 5:21.53 | אשר |  |
| 5:21.54 | יצוה |  |
| 5:21.55 | את |  |
| 5:21.56 | בניו |  |
| 5:21.57 | ואת |  |
| 5:21.58 | ביתו |  |
| 5:21.59 | אחריו |  |
| 5:21.60 | ושמרו |  |
| 5:21.61 | דרך |  |
| 5:21.62 | השם | " |
| 5:21.63 | וכו | ' |
| 5:21.64 | וכן |  |
| 5:21.65 | עשה |  |
| 5:21.66 | אברהם |  |
| 5:21.67 | אבינו |  |
| 5:21.68 | וציוה |  |
| 5:21.69 | לבניו |  |
| 5:21.70 | יצחק |  |
| 5:21.71 | וישמעאל |  |
| 5:21.72 | אמונת |  |
| 5:21.73 | היחוד |  |
| 5:21.74 | שהוא |  |
| 5:21.75 | יתברך |  |
| 5:21.76 | אחד |  |
| 5:21.77 | אמת | , |
| 5:21.78 | ואין |  |
| 5:21.79 | אחד |  |
| 5:21.80 | אמת |  |
| 5:21.81 | זולתו | , |
| 5:21.82 | ושהוא |  |
| 5:21.83 | קדמון |  |
| 5:21.84 | לכל |  |
| 5:21.85 | נברא | , |
| 5:21.86 | ראשון |  |
| 5:21.87 | ואין |  |
| 5:21.88 | ראשית |  |
| 5:21.89 | לראשיתו | ! |


## Milchamot ha-Shem 5:22

Hebrew (Wikisource): וזה לשון מדרש הגדול: תרח אבי אברהם היה עובד ע"ז והיה עושה צלמים וכו' וכשמלך נמרוד הרשע באותו זמן ועשה עצמו אלוה וכו' והיו כל האומות באים ומשתחוים לו וכורעין לפניו עד שזרח אורו של עולם אברהם אבינו, וכשנולד אברהם אבינו אמרו לו אצטגניניו קוסמיו וחרטומיו: היום נולד נער שהוא עתיד לירש את העולם ועתיד לאבד את המלכיות ולהוציאך ממלכותך, אם רצונך שלח אצל אביו ופדהו ממנו בבית מלא כסף וזהב ועקרהו מן העולם, והיה תרח יושב ביניהם, אמר להם דומין דבריכם לפרד שאומרים לו טול כור שעורים ונחתוך את ראשך. מה יעשה בשעורים אחר שיחתכו את ראשו? כך זה כשייהרג בנו מה יעשה בדמיו? אמרו לו מתוך דבריך ניכר שנולד לך בן זה, אמר להם לא אכחד מכם בן נולד לי וכבר מת, אמרו לו לא על מת אנו מספרים אלא על חי ואנו רואים מזלו ברקיע, עכשיו לך והבא אותו, מיד הלך תרח ולקח את אברהם ואת מנקתו והטמינם במערה והיו שם שלוש שנים. אחר כך יצא מן המערה ובא אצל אביו והיה לבו משוטט ביצירת העולם ובמאורות להשתחוות להן ולעבדן ולידע איזה מהן הוא אלוה, ראה הלבנה שאורה מבהיק בלילה מסוף העולם ועד סופו ופמליא שלה מרובה, אמר זה אלוה. עברה כל הלילה. בשחרית כיון שזרחה חמה וחשך מאור הלבנה ותשש כחה, אמר זה אלוה. לערב חשכה ויצאה הלבנה והכוכבים והמזלות, אמר בודאי יש אדון לאלו וכו' כיון שנעשה לו נס ויצא בשלום, באותה שעה כיון אברהם אבינו דעתו לשמים ואמר, יתברך שמו של הקב"ה שחמה ולבנה כוכבים ומזלות אינם משתחוים אלא לו. ושמא תאמר רב היה לאברהם אלא זימן לו הקב"ה שתי כליותיו כשני מעינות והיו מוציאין לו חכמה שנאמר הן אמת חפצת בטוחות ובסתום חכמה תודיעני" וכו'.

| id | word | punctuation after |
|---|---|---|
| 5:22.1 | וזה |  |
| 5:22.2 | לשון |  |
| 5:22.3 | מדרש |  |
| 5:22.4 | הגדול | : |
| 5:22.5 | תרח |  |
| 5:22.6 | אבי |  |
| 5:22.7 | אברהם |  |
| 5:22.8 | היה |  |
| 5:22.9 | עובד |  |
| 5:22.10 | ע"ז |  |
| 5:22.11 | והיה |  |
| 5:22.12 | עושה |  |
| 5:22.13 | צלמים |  |
| 5:22.14 | וכו | ' |
| 5:22.15 | וכשמלך |  |
| 5:22.16 | נמרוד |  |
| 5:22.17 | הרשע |  |
| 5:22.18 | באותו |  |
| 5:22.19 | זמן |  |
| 5:22.20 | ועשה |  |
| 5:22.21 | עצמו |  |
| 5:22.22 | אלוה |  |
| 5:22.23 | וכו | ' |
| 5:22.24 | והיו |  |
| 5:22.25 | כל |  |
| 5:22.26 | האומות |  |
| 5:22.27 | באים |  |
| 5:22.28 | ומשתחוים |  |
| 5:22.29 | לו |  |
| 5:22.30 | וכורעין |  |
| 5:22.31 | לפניו |  |
| 5:22.32 | עד |  |
| 5:22.33 | שזרח |  |
| 5:22.34 | אורו |  |
| 5:22.35 | של |  |
| 5:22.36 | עולם |  |
| 5:22.37 | אברהם |  |
| 5:22.38 | אבינו | , |
| 5:22.39 | וכשנולד |  |
| 5:22.40 | אברהם |  |
| 5:22.41 | אבינו |  |
| 5:22.42 | אמרו |  |
| 5:22.43 | לו |  |
| 5:22.44 | אצטגניניו |  |
| 5:22.45 | קוסמיו |  |
| 5:22.46 | וחרטומיו | : |
| 5:22.47 | היום |  |
| 5:22.48 | נולד |  |
| 5:22.49 | נער |  |
| 5:22.50 | שהוא |  |
| 5:22.51 | עתיד |  |
| 5:22.52 | לירש |  |
| 5:22.53 | את |  |
| 5:22.54 | העולם |  |
| 5:22.55 | ועתיד |  |
| 5:22.56 | לאבד |  |
| 5:22.57 | את |  |
| 5:22.58 | המלכיות |  |
| 5:22.59 | ולהוציאך |  |
| 5:22.60 | ממלכותך | , |
| 5:22.61 | אם |  |
| 5:22.62 | רצונך |  |
| 5:22.63 | שלח |  |
| 5:22.64 | אצל |  |
| 5:22.65 | אביו |  |
| 5:22.66 | ופדהו |  |
| 5:22.67 | ממנו |  |
| 5:22.68 | בבית |  |
| 5:22.69 | מלא |  |
| 5:22.70 | כסף |  |
| 5:22.71 | וזהב |  |
| 5:22.72 | ועקרהו |  |
| 5:22.73 | מן |  |
| 5:22.74 | העולם | , |
| 5:22.75 | והיה |  |
| 5:22.76 | תרח |  |
| 5:22.77 | יושב |  |
| 5:22.78 | ביניהם | , |
| 5:22.79 | אמר |  |
| 5:22.80 | להם |  |
| 5:22.81 | דומין |  |
| 5:22.82 | דבריכם |  |
| 5:22.83 | לפרד |  |
| 5:22.84 | שאומרים |  |
| 5:22.85 | לו |  |
| 5:22.86 | טול |  |
| 5:22.87 | כור |  |
| 5:22.88 | שעורים |  |
| 5:22.89 | ונחתוך |  |
| 5:22.90 | את |  |
| 5:22.91 | ראשך | . |
| 5:22.92 | מה |  |
| 5:22.93 | יעשה |  |
| 5:22.94 | בשעורים |  |
| 5:22.95 | אחר |  |
| 5:22.96 | שיחתכו |  |
| 5:22.97 | את |  |
| 5:22.98 | ראשו | ? |
| 5:22.99 | כך |  |
| 5:22.100 | זה |  |
| 5:22.101 | כשייהרג |  |
| 5:22.102 | בנו |  |
| 5:22.103 | מה |  |
| 5:22.104 | יעשה |  |
| 5:22.105 | בדמיו | ? |
| 5:22.106 | אמרו |  |
| 5:22.107 | לו |  |
| 5:22.108 | מתוך |  |
| 5:22.109 | דבריך |  |
| 5:22.110 | ניכר |  |
| 5:22.111 | שנולד |  |
| 5:22.112 | לך |  |
| 5:22.113 | בן |  |
| 5:22.114 | זה | , |
| 5:22.115 | אמר |  |
| 5:22.116 | להם |  |
| 5:22.117 | לא |  |
| 5:22.118 | אכחד |  |
| 5:22.119 | מכם |  |
| 5:22.120 | בן |  |
| 5:22.121 | נולד |  |
| 5:22.122 | לי |  |
| 5:22.123 | וכבר |  |
| 5:22.124 | מת | , |
| 5:22.125 | אמרו |  |
| 5:22.126 | לו |  |
| 5:22.127 | לא |  |
| 5:22.128 | על |  |
| 5:22.129 | מת |  |
| 5:22.130 | אנו |  |
| 5:22.131 | מספרים |  |
| 5:22.132 | אלא |  |
| 5:22.133 | על |  |
| 5:22.134 | חי |  |
| 5:22.135 | ואנו |  |
| 5:22.136 | רואים |  |
| 5:22.137 | מזלו |  |
| 5:22.138 | ברקיע | , |
| 5:22.139 | עכשיו |  |
| 5:22.140 | לך |  |
| 5:22.141 | והבא |  |
| 5:22.142 | אותו | , |
| 5:22.143 | מיד |  |
| 5:22.144 | הלך |  |
| 5:22.145 | תרח |  |
| 5:22.146 | ולקח |  |
| 5:22.147 | את |  |
| 5:22.148 | אברהם |  |
| 5:22.149 | ואת |  |
| 5:22.150 | מנקתו |  |
| 5:22.151 | והטמינם |  |
| 5:22.152 | במערה |  |
| 5:22.153 | והיו |  |
| 5:22.154 | שם |  |
| 5:22.155 | שלוש |  |
| 5:22.156 | שנים | . |
| 5:22.157 | אחר |  |
| 5:22.158 | כך |  |
| 5:22.159 | יצא |  |
| 5:22.160 | מן |  |
| 5:22.161 | המערה |  |
| 5:22.162 | ובא |  |
| 5:22.163 | אצל |  |
| 5:22.164 | אביו |  |
| 5:22.165 | והיה |  |
| 5:22.166 | לבו |  |
| 5:22.167 | משוטט |  |
| 5:22.168 | ביצירת |  |
| 5:22.169 | העולם |  |
| 5:22.170 | ובמאורות |  |
| 5:22.171 | להשתחוות |  |
| 5:22.172 | להן |  |
| 5:22.173 | ולעבדן |  |
| 5:22.174 | ולידע |  |
| 5:22.175 | איזה |  |
| 5:22.176 | מהן |  |
| 5:22.177 | הוא |  |
| 5:22.178 | אלוה | , |
| 5:22.179 | ראה |  |
| 5:22.180 | הלבנה |  |
| 5:22.181 | שאורה |  |
| 5:22.182 | מבהיק |  |
| 5:22.183 | בלילה |  |
| 5:22.184 | מסוף |  |
| 5:22.185 | העולם |  |
| 5:22.186 | ועד |  |
| 5:22.187 | סופו |  |
| 5:22.188 | ופמליא |  |
| 5:22.189 | שלה |  |
| 5:22.190 | מרובה | , |
| 5:22.191 | אמר |  |
| 5:22.192 | זה |  |
| 5:22.193 | אלוה | . |
| 5:22.194 | עברה |  |
| 5:22.195 | כל |  |
| 5:22.196 | הלילה | . |
| 5:22.197 | בשחרית |  |
| 5:22.198 | כיון |  |
| 5:22.199 | שזרחה |  |
| 5:22.200 | חמה |  |
| 5:22.201 | וחשך |  |
| 5:22.202 | מאור |  |
| 5:22.203 | הלבנה |  |
| 5:22.204 | ותשש |  |
| 5:22.205 | כחה | , |
| 5:22.206 | אמר |  |
| 5:22.207 | זה |  |
| 5:22.208 | אלוה | . |
| 5:22.209 | לערב |  |
| 5:22.210 | חשכה |  |
| 5:22.211 | ויצאה |  |
| 5:22.212 | הלבנה |  |
| 5:22.213 | והכוכבים |  |
| 5:22.214 | והמזלות | , |
| 5:22.215 | אמר |  |
| 5:22.216 | בודאי |  |
| 5:22.217 | יש |  |
| 5:22.218 | אדון |  |
| 5:22.219 | לאלו |  |
| 5:22.220 | וכו | ' |
| 5:22.221 | כיון |  |
| 5:22.222 | שנעשה |  |
| 5:22.223 | לו |  |
| 5:22.224 | נס |  |
| 5:22.225 | ויצא |  |
| 5:22.226 | בשלום | , |
| 5:22.227 | באותה |  |
| 5:22.228 | שעה |  |
| 5:22.229 | כיון |  |
| 5:22.230 | אברהם |  |
| 5:22.231 | אבינו |  |
| 5:22.232 | דעתו |  |
| 5:22.233 | לשמים |  |
| 5:22.234 | ואמר | , |
| 5:22.235 | יתברך |  |
| 5:22.236 | שמו |  |
| 5:22.237 | של |  |
| 5:22.238 | הקב"ה |  |
| 5:22.239 | שחמה |  |
| 5:22.240 | ולבנה |  |
| 5:22.241 | כוכבים |  |
| 5:22.242 | ומזלות |  |
| 5:22.243 | אינם |  |
| 5:22.244 | משתחוים |  |
| 5:22.245 | אלא |  |
| 5:22.246 | לו | . |
| 5:22.247 | ושמא |  |
| 5:22.248 | תאמר |  |
| 5:22.249 | רב |  |
| 5:22.250 | היה |  |
| 5:22.251 | לאברהם |  |
| 5:22.252 | אלא |  |
| 5:22.253 | זימן |  |
| 5:22.254 | לו |  |
| 5:22.255 | הקב"ה |  |
| 5:22.256 | שתי |  |
| 5:22.257 | כליותיו |  |
| 5:22.258 | כשני |  |
| 5:22.259 | מעינות |  |
| 5:22.260 | והיו |  |
| 5:22.261 | מוציאין |  |
| 5:22.262 | לו |  |
| 5:22.263 | חכמה |  |
| 5:22.264 | שנאמר |  |
| 5:22.265 | הן |  |
| 5:22.266 | אמת |  |
| 5:22.267 | חפצת |  |
| 5:22.268 | בטוחות |  |
| 5:22.269 | ובסתום |  |
| 5:22.270 | חכמה |  |
| 5:22.271 | תודיעני | " |
| 5:22.272 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

