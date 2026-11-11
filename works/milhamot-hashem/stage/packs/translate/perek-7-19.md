# Translation pack: milhamot-hashem perek-7-19 (Milchamot ha-Shem 7:24–7:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-19`.

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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
- האלהי before a name: "the divine", with a note; never dropped.
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
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


## Milchamot ha-Shem 7:24

Hebrew (Wikisource): צו) גם הרס"ג בפירושו לס' איוב מרחיק מאד סברת האומר שבני האלהים מלאכים וביזה וזלזל את האומר שהם מלאכים כמו שעשה רשב"י בבראשית רבה וז"ל בלשון ערבי. ומא אבעד אלקיאס אן יכון הדא אלשטן מלכא, לאן אהל אלתוחיד מגמעין עלי' אן אלבארי נל ועז כלק מלאיכתה מקרבין. מן עלמה אנחם לא יעצון כמא אנה אצטפא אנביאה אלמרסלין מן עלמה אנהם לא יעצונה. ומעאנדה אלאוליא אלצאלחין מן כבאיר אלמעצי! פמן זעם אם הדא אלמעאנד לאיוב כאן מלכא פקד נעל אלמלאיכה יעצון. וכאלף מא נחן מגמעין עליה, וקד תעלם ארשדך אללה אן לאלנפס תלאת קוי, אלפכר, ואלגצב ואלשהות, פאמא קווה אלפכר פחי אלתר ימייז כתא אל אנסאן אלצואב מן אלכטא ואלחק מן אלבאטל, ואמא קווה אלגצב פהי אלתי כהא ילחק אלחיואן אלאקדאם ואלגראות ואלאנתקאם ואסתיפא אלטואיל, ואמא קוה אלשהוה פבהא ישתאק אלחיואן אלי אלגדא ואלי אלתלקיח וסאיר אלדעאוי. ותעלם מע דאלך אן צפה עדל אלנאטק אן תכון אלקוותן אעני קווה אלשהוה ואלגצב מדענה לאלפכר (פאן) גלב אליהא אחדהמא כאן דאלך גורא וכטאא, פמן גאוז אן תפעל אלמלאיכה אלחסד אלדי הו אחד אגואע אלגצב לזמה תגויז וקוע אלשהואת מן אלמלאיכה והו אלאכל ואלשרב ואלנכאח וסאיר אלאמור אלרדלה אלוסלה אלתי קד צאן אללה מלאכיתה ענהא, סימא וקד סבק פי עלמה אנהם לא יעצונה. פלא יגוז אן יכון באלפעל כלאף מא סבק פי עלמה וכו' וקד בלגני ען בעץ אלכוארג אן כאלפה בה אלי' אן זעם אן הדא אלשטן מלך. ואן אלחסד ואלגירה ואלבני יגוז וקועהם באלפעל מן אלמלאיכה כמה תוהם פי קצח' איוב וחסד אלשטן לה' תם ראא איצא אן באטלאקה עליהם דאלך ילזמה אטלאק אל שיאן ואלנבאח, פנאוז המא איצא. ואדעא אן קצה ויראו בני האלהים את בנות האדם כי טבת הנה. הם אמלאיכה אלנאכחון, אלחרם באסקאטהם מן אלסמא ותציירהם כאלנאס, ונצב להם שבחה הו קולה אני אמרתי אלהים אתם ובני עליון כלכם אבן כאדם תמותון וכאחד השרים תפלו, פלם יבק פאחשה ולא קביחה חתי וצף בהא אלמלאיכה, כל דאלך גאריא עלי' אצלה הדא, פרכב הדה אלשנאעאת אלעצימה לקלה' בצרה בתפסיר אלכלאם באלגה, ואנא מביןן לטאה פי גמיע מא וצעה בעון אללה, אעלם אסעדך אללה אן אלשבהה אלדי דעת הדא אלמבדע אלהרי וגירה אל' הדא אלקול אן אלשטן מלך הוא קול אללה צעלי' לאלשטן הנה כל אשר לו בידיך רק אליו אל תשלח ידיך, ואיצא קולה רק את נפשו שמור, ואיצא קולה ויך את איוב בשחין רע, ותוהם אן הדא מנסוב אלי' אלשטן. ואנא אשרח מא בה תזול הדה אל תלאת אלשבה במא פיה מקנע. פאקול אן אלמראד ואלראי פי לגה' בני אסראיל יסמא יד. כמא קאל דוד לאלמסתגיתה אליה היד יואב אתך בכל זאת, מענאה ראי ומראד וכו' וענד זואל הדה אלשבהה אקול אן כטא הדא אלמעבר פי תלריג ויראו בני האלהים את בנות האדם. אפצע מן מא קאלה פי קצה איוב ודאלך אנא לם יקף בתה עלי' מא יסתעמלון בני אסראיל פי אלכאץ ואלעאם, ודאלך אנה מן סבילהם אן יקולון פעל אלכואץ כדא ופעל אלעואם כדא ויכון אולאיך אלכואץ איצא נאנא, ואנמא פרדהם לתכציציהם, פעלי, הדא קולה ויראו בני האלהים נאסא איצא אנמא פרדהם מן אלעואם לתכציצהם ושרפהם. והדה עבארה מתגהה לתנפי אלמנכר ען אלמקרבין. ויזידהא איצא ויקויהא קולה בעד דאלך וירא ה' כי רבה רעת האדם, ולם ידם אלמלאיכה מעהם. וקולה פי עקב דאלך אמחה את האדם אשר בראתי ולם יעאקב אלמלאיכה וכו' עכ"ל. העתק תמצא בסוף הספר הרי בהדיא שהגאון מרחיק מאד דעת האומר שבני האלהים הם מלאכים ודוחה אותה בשתי ידים וקורא לה דעת חיצונית. ודבריו מכוונים לדברי הרשב"י שבבראשית רבה, שלא כדעת הזהר שאמר שהם מלאכים, ושונה ומשלם בה, וזו הוכחה גדולה שאינו מרשב"י שהרי רשב"י היה מקלל את האומר שהם מלאכים. ובענין השטן שפי' הגאון שהוא אדם, אע"פ שהראב"ע השיגו יש מקום סתירה להשגתו ולהעמיד דברי הגאון הרס"ג, ואע"פ שבפ"ק דבבא בתרא משמע כהראב"ע. הנה ידוע מה שכ' הנגיד בעל מבוא התלמוד בפי' הגדה, ואין לך ללמוד ממנה אלא מה שיעלה על הדעת, ויש לך לדעת שכל מה שקיימו חז"ל הלכה בענין מצוה, שהוא מפי משה רבינו עליו השלום מפי הגבורה, אין לך להוסיף עליו ולא לגרוע ממנו אבל מה שפירשו בפסוקים כל אחד כפי שיזדמן לו ומה שראה בדעתו ולפי מה שיעלה על הדעת מן הפירושים האלה לומדים אותם, והשאר אין סומכים עליהם עכ"ל. ועיין רש"י שבת פ' במה מדליקים בשאלת מהו לכבויי בוצינא מקמי באישא בשבתא. והיוצא מזה דכל למסבר קראי בעלמא ואין נוגע ממנו לדינא בעשיית המצות ובעיקרי תורתינו הקדושה, אין בומכים אלא על כל מאי דמסתבר טפי. אבל כל מאי דפנקא מינא לדינא כנדון דידן כמצות יחוד השם יתברך אין לזוז מקבלת חז"ל שאין לשתף עם השם ית' שום כח נברא בין שיהיה גשמי או רוחני כמאמר רשב"י ע"ח כל המשתף שם שמים ודבר אחר נעקר מן העולם.

| id | word | punctuation after |
|---|---|---|
| 7:24.1 | צו | ) |
| 7:24.2 | גם |  |
| 7:24.3 | הרס"ג |  |
| 7:24.4 | בפירושו |  |
| 7:24.5 | לס | ' |
| 7:24.6 | איוב |  |
| 7:24.7 | מרחיק |  |
| 7:24.8 | מאד |  |
| 7:24.9 | סברת |  |
| 7:24.10 | האומר |  |
| 7:24.11 | שבני |  |
| 7:24.12 | האלהים |  |
| 7:24.13 | מלאכים |  |
| 7:24.14 | וביזה |  |
| 7:24.15 | וזלזל |  |
| 7:24.16 | את |  |
| 7:24.17 | האומר |  |
| 7:24.18 | שהם |  |
| 7:24.19 | מלאכים |  |
| 7:24.20 | כמו |  |
| 7:24.21 | שעשה |  |
| 7:24.22 | רשב"י |  |
| 7:24.23 | בבראשית |  |
| 7:24.24 | רבה |  |
| 7:24.25 | וז"ל |  |
| 7:24.26 | בלשון |  |
| 7:24.27 | ערבי | . |
| 7:24.28 | ומא |  |
| 7:24.29 | אבעד |  |
| 7:24.30 | אלקיאס |  |
| 7:24.31 | אן |  |
| 7:24.32 | יכון |  |
| 7:24.33 | הדא |  |
| 7:24.34 | אלשטן |  |
| 7:24.35 | מלכא | , |
| 7:24.36 | לאן |  |
| 7:24.37 | אהל |  |
| 7:24.38 | אלתוחיד |  |
| 7:24.39 | מגמעין |  |
| 7:24.40 | עלי | ' |
| 7:24.41 | אן |  |
| 7:24.42 | אלבארי |  |
| 7:24.43 | נל |  |
| 7:24.44 | ועז |  |
| 7:24.45 | כלק |  |
| 7:24.46 | מלאיכתה |  |
| 7:24.47 | מקרבין | . |
| 7:24.48 | מן |  |
| 7:24.49 | עלמה |  |
| 7:24.50 | אנחם |  |
| 7:24.51 | לא |  |
| 7:24.52 | יעצון |  |
| 7:24.53 | כמא |  |
| 7:24.54 | אנה |  |
| 7:24.55 | אצטפא |  |
| 7:24.56 | אנביאה |  |
| 7:24.57 | אלמרסלין |  |
| 7:24.58 | מן |  |
| 7:24.59 | עלמה |  |
| 7:24.60 | אנהם |  |
| 7:24.61 | לא |  |
| 7:24.62 | יעצונה | . |
| 7:24.63 | ומעאנדה |  |
| 7:24.64 | אלאוליא |  |
| 7:24.65 | אלצאלחין |  |
| 7:24.66 | מן |  |
| 7:24.67 | כבאיר |  |
| 7:24.68 | אלמעצי | ! |
| 7:24.69 | פמן |  |
| 7:24.70 | זעם |  |
| 7:24.71 | אם |  |
| 7:24.72 | הדא |  |
| 7:24.73 | אלמעאנד |  |
| 7:24.74 | לאיוב |  |
| 7:24.75 | כאן |  |
| 7:24.76 | מלכא |  |
| 7:24.77 | פקד |  |
| 7:24.78 | נעל |  |
| 7:24.79 | אלמלאיכה |  |
| 7:24.80 | יעצון | . |
| 7:24.81 | וכאלף |  |
| 7:24.82 | מא |  |
| 7:24.83 | נחן |  |
| 7:24.84 | מגמעין |  |
| 7:24.85 | עליה | , |
| 7:24.86 | וקד |  |
| 7:24.87 | תעלם |  |
| 7:24.88 | ארשדך |  |
| 7:24.89 | אללה |  |
| 7:24.90 | אן |  |
| 7:24.91 | לאלנפס |  |
| 7:24.92 | תלאת |  |
| 7:24.93 | קוי | , |
| 7:24.94 | אלפכר | , |
| 7:24.95 | ואלגצב |  |
| 7:24.96 | ואלשהות | , |
| 7:24.97 | פאמא |  |
| 7:24.98 | קווה |  |
| 7:24.99 | אלפכר |  |
| 7:24.100 | פחי |  |
| 7:24.101 | אלתר |  |
| 7:24.102 | ימייז |  |
| 7:24.103 | כתא |  |
| 7:24.104 | אל |  |
| 7:24.105 | אנסאן |  |
| 7:24.106 | אלצואב |  |
| 7:24.107 | מן |  |
| 7:24.108 | אלכטא |  |
| 7:24.109 | ואלחק |  |
| 7:24.110 | מן |  |
| 7:24.111 | אלבאטל | , |
| 7:24.112 | ואמא |  |
| 7:24.113 | קווה |  |
| 7:24.114 | אלגצב |  |
| 7:24.115 | פהי |  |
| 7:24.116 | אלתי |  |
| 7:24.117 | כהא |  |
| 7:24.118 | ילחק |  |
| 7:24.119 | אלחיואן |  |
| 7:24.120 | אלאקדאם |  |
| 7:24.121 | ואלגראות |  |
| 7:24.122 | ואלאנתקאם |  |
| 7:24.123 | ואסתיפא |  |
| 7:24.124 | אלטואיל | , |
| 7:24.125 | ואמא |  |
| 7:24.126 | קוה |  |
| 7:24.127 | אלשהוה |  |
| 7:24.128 | פבהא |  |
| 7:24.129 | ישתאק |  |
| 7:24.130 | אלחיואן |  |
| 7:24.131 | אלי |  |
| 7:24.132 | אלגדא |  |
| 7:24.133 | ואלי |  |
| 7:24.134 | אלתלקיח |  |
| 7:24.135 | וסאיר |  |
| 7:24.136 | אלדעאוי | . |
| 7:24.137 | ותעלם |  |
| 7:24.138 | מע |  |
| 7:24.139 | דאלך |  |
| 7:24.140 | אן |  |
| 7:24.141 | צפה |  |
| 7:24.142 | עדל |  |
| 7:24.143 | אלנאטק |  |
| 7:24.144 | אן |  |
| 7:24.145 | תכון |  |
| 7:24.146 | אלקוותן |  |
| 7:24.147 | אעני |  |
| 7:24.148 | קווה |  |
| 7:24.149 | אלשהוה |  |
| 7:24.150 | ואלגצב |  |
| 7:24.151 | מדענה |  |
| 7:24.152 | לאלפכר |  |
| 7:24.153 | פאן | ) |
| 7:24.154 | גלב |  |
| 7:24.155 | אליהא |  |
| 7:24.156 | אחדהמא |  |
| 7:24.157 | כאן |  |
| 7:24.158 | דאלך |  |
| 7:24.159 | גורא |  |
| 7:24.160 | וכטאא | , |
| 7:24.161 | פמן |  |
| 7:24.162 | גאוז |  |
| 7:24.163 | אן |  |
| 7:24.164 | תפעל |  |
| 7:24.165 | אלמלאיכה |  |
| 7:24.166 | אלחסד |  |
| 7:24.167 | אלדי |  |
| 7:24.168 | הו |  |
| 7:24.169 | אחד |  |
| 7:24.170 | אגואע |  |
| 7:24.171 | אלגצב |  |
| 7:24.172 | לזמה |  |
| 7:24.173 | תגויז |  |
| 7:24.174 | וקוע |  |
| 7:24.175 | אלשהואת |  |
| 7:24.176 | מן |  |
| 7:24.177 | אלמלאיכה |  |
| 7:24.178 | והו |  |
| 7:24.179 | אלאכל |  |
| 7:24.180 | ואלשרב |  |
| 7:24.181 | ואלנכאח |  |
| 7:24.182 | וסאיר |  |
| 7:24.183 | אלאמור |  |
| 7:24.184 | אלרדלה |  |
| 7:24.185 | אלוסלה |  |
| 7:24.186 | אלתי |  |
| 7:24.187 | קד |  |
| 7:24.188 | צאן |  |
| 7:24.189 | אללה |  |
| 7:24.190 | מלאכיתה |  |
| 7:24.191 | ענהא | , |
| 7:24.192 | סימא |  |
| 7:24.193 | וקד |  |
| 7:24.194 | סבק |  |
| 7:24.195 | פי |  |
| 7:24.196 | עלמה |  |
| 7:24.197 | אנהם |  |
| 7:24.198 | לא |  |
| 7:24.199 | יעצונה | . |
| 7:24.200 | פלא |  |
| 7:24.201 | יגוז |  |
| 7:24.202 | אן |  |
| 7:24.203 | יכון |  |
| 7:24.204 | באלפעל |  |
| 7:24.205 | כלאף |  |
| 7:24.206 | מא |  |
| 7:24.207 | סבק |  |
| 7:24.208 | פי |  |
| 7:24.209 | עלמה |  |
| 7:24.210 | וכו | ' |
| 7:24.211 | וקד |  |
| 7:24.212 | בלגני |  |
| 7:24.213 | ען |  |
| 7:24.214 | בעץ |  |
| 7:24.215 | אלכוארג |  |
| 7:24.216 | אן |  |
| 7:24.217 | כאלפה |  |
| 7:24.218 | בה |  |
| 7:24.219 | אלי | ' |
| 7:24.220 | אן |  |
| 7:24.221 | זעם |  |
| 7:24.222 | אן |  |
| 7:24.223 | הדא |  |
| 7:24.224 | אלשטן |  |
| 7:24.225 | מלך | . |
| 7:24.226 | ואן |  |
| 7:24.227 | אלחסד |  |
| 7:24.228 | ואלגירה |  |
| 7:24.229 | ואלבני |  |
| 7:24.230 | יגוז |  |
| 7:24.231 | וקועהם |  |
| 7:24.232 | באלפעל |  |
| 7:24.233 | מן |  |
| 7:24.234 | אלמלאיכה |  |
| 7:24.235 | כמה |  |
| 7:24.236 | תוהם |  |
| 7:24.237 | פי |  |
| 7:24.238 | קצח | ' |
| 7:24.239 | איוב |  |
| 7:24.240 | וחסד |  |
| 7:24.241 | אלשטן |  |
| 7:24.242 | לה | ' |
| 7:24.243 | תם |  |
| 7:24.244 | ראא |  |
| 7:24.245 | איצא |  |
| 7:24.246 | אן |  |
| 7:24.247 | באטלאקה |  |
| 7:24.248 | עליהם |  |
| 7:24.249 | דאלך |  |
| 7:24.250 | ילזמה |  |
| 7:24.251 | אטלאק |  |
| 7:24.252 | אל |  |
| 7:24.253 | שיאן |  |
| 7:24.254 | ואלנבאח | , |
| 7:24.255 | פנאוז |  |
| 7:24.256 | המא |  |
| 7:24.257 | איצא | . |
| 7:24.258 | ואדעא |  |
| 7:24.259 | אן |  |
| 7:24.260 | קצה |  |
| 7:24.261 | ויראו |  |
| 7:24.262 | בני |  |
| 7:24.263 | האלהים |  |
| 7:24.264 | את |  |
| 7:24.265 | בנות |  |
| 7:24.266 | האדם |  |
| 7:24.267 | כי |  |
| 7:24.268 | טבת |  |
| 7:24.269 | הנה | . |
| 7:24.270 | הם |  |
| 7:24.271 | אמלאיכה |  |
| 7:24.272 | אלנאכחון | , |
| 7:24.273 | אלחרם |  |
| 7:24.274 | באסקאטהם |  |
| 7:24.275 | מן |  |
| 7:24.276 | אלסמא |  |
| 7:24.277 | ותציירהם |  |
| 7:24.278 | כאלנאס | , |
| 7:24.279 | ונצב |  |
| 7:24.280 | להם |  |
| 7:24.281 | שבחה |  |
| 7:24.282 | הו |  |
| 7:24.283 | קולה |  |
| 7:24.284 | אני |  |
| 7:24.285 | אמרתי |  |
| 7:24.286 | אלהים |  |
| 7:24.287 | אתם |  |
| 7:24.288 | ובני |  |
| 7:24.289 | עליון |  |
| 7:24.290 | כלכם |  |
| 7:24.291 | אבן |  |
| 7:24.292 | כאדם |  |
| 7:24.293 | תמותון |  |
| 7:24.294 | וכאחד |  |
| 7:24.295 | השרים |  |
| 7:24.296 | תפלו | , |
| 7:24.297 | פלם |  |
| 7:24.298 | יבק |  |
| 7:24.299 | פאחשה |  |
| 7:24.300 | ולא |  |
| 7:24.301 | קביחה |  |
| 7:24.302 | חתי |  |
| 7:24.303 | וצף |  |
| 7:24.304 | בהא |  |
| 7:24.305 | אלמלאיכה | , |
| 7:24.306 | כל |  |
| 7:24.307 | דאלך |  |
| 7:24.308 | גאריא |  |
| 7:24.309 | עלי | ' |
| 7:24.310 | אצלה |  |
| 7:24.311 | הדא | , |
| 7:24.312 | פרכב |  |
| 7:24.313 | הדה |  |
| 7:24.314 | אלשנאעאת |  |
| 7:24.315 | אלעצימה |  |
| 7:24.316 | לקלה | ' |
| 7:24.317 | בצרה |  |
| 7:24.318 | בתפסיר |  |
| 7:24.319 | אלכלאם |  |
| 7:24.320 | באלגה | , |
| 7:24.321 | ואנא |  |
| 7:24.322 | מביןן |  |
| 7:24.323 | לטאה |  |
| 7:24.324 | פי |  |
| 7:24.325 | גמיע |  |
| 7:24.326 | מא |  |
| 7:24.327 | וצעה |  |
| 7:24.328 | בעון |  |
| 7:24.329 | אללה | , |
| 7:24.330 | אעלם |  |
| 7:24.331 | אסעדך |  |
| 7:24.332 | אללה |  |
| 7:24.333 | אן |  |
| 7:24.334 | אלשבהה |  |
| 7:24.335 | אלדי |  |
| 7:24.336 | דעת |  |
| 7:24.337 | הדא |  |
| 7:24.338 | אלמבדע |  |
| 7:24.339 | אלהרי |  |
| 7:24.340 | וגירה |  |
| 7:24.341 | אל | ' |
| 7:24.342 | הדא |  |
| 7:24.343 | אלקול |  |
| 7:24.344 | אן |  |
| 7:24.345 | אלשטן |  |
| 7:24.346 | מלך |  |
| 7:24.347 | הוא |  |
| 7:24.348 | קול |  |
| 7:24.349 | אללה |  |
| 7:24.350 | צעלי | ' |
| 7:24.351 | לאלשטן |  |
| 7:24.352 | הנה |  |
| 7:24.353 | כל |  |
| 7:24.354 | אשר |  |
| 7:24.355 | לו |  |
| 7:24.356 | בידיך |  |
| 7:24.357 | רק |  |
| 7:24.358 | אליו |  |
| 7:24.359 | אל |  |
| 7:24.360 | תשלח |  |
| 7:24.361 | ידיך | , |
| 7:24.362 | ואיצא |  |
| 7:24.363 | קולה |  |
| 7:24.364 | רק |  |
| 7:24.365 | את |  |
| 7:24.366 | נפשו |  |
| 7:24.367 | שמור | , |
| 7:24.368 | ואיצא |  |
| 7:24.369 | קולה |  |
| 7:24.370 | ויך |  |
| 7:24.371 | את |  |
| 7:24.372 | איוב |  |
| 7:24.373 | בשחין |  |
| 7:24.374 | רע | , |
| 7:24.375 | ותוהם |  |
| 7:24.376 | אן |  |
| 7:24.377 | הדא |  |
| 7:24.378 | מנסוב |  |
| 7:24.379 | אלי | ' |
| 7:24.380 | אלשטן | . |
| 7:24.381 | ואנא |  |
| 7:24.382 | אשרח |  |
| 7:24.383 | מא |  |
| 7:24.384 | בה |  |
| 7:24.385 | תזול |  |
| 7:24.386 | הדה |  |
| 7:24.387 | אל |  |
| 7:24.388 | תלאת |  |
| 7:24.389 | אלשבה |  |
| 7:24.390 | במא |  |
| 7:24.391 | פיה |  |
| 7:24.392 | מקנע | . |
| 7:24.393 | פאקול |  |
| 7:24.394 | אן |  |
| 7:24.395 | אלמראד |  |
| 7:24.396 | ואלראי |  |
| 7:24.397 | פי |  |
| 7:24.398 | לגה | ' |
| 7:24.399 | בני |  |
| 7:24.400 | אסראיל |  |
| 7:24.401 | יסמא |  |
| 7:24.402 | יד | . |
| 7:24.403 | כמא |  |
| 7:24.404 | קאל |  |
| 7:24.405 | דוד |  |
| 7:24.406 | לאלמסתגיתה |  |
| 7:24.407 | אליה |  |
| 7:24.408 | היד |  |
| 7:24.409 | יואב |  |
| 7:24.410 | אתך |  |
| 7:24.411 | בכל |  |
| 7:24.412 | זאת | , |
| 7:24.413 | מענאה |  |
| 7:24.414 | ראי |  |
| 7:24.415 | ומראד |  |
| 7:24.416 | וכו | ' |
| 7:24.417 | וענד |  |
| 7:24.418 | זואל |  |
| 7:24.419 | הדה |  |
| 7:24.420 | אלשבהה |  |
| 7:24.421 | אקול |  |
| 7:24.422 | אן |  |
| 7:24.423 | כטא |  |
| 7:24.424 | הדא |  |
| 7:24.425 | אלמעבר |  |
| 7:24.426 | פי |  |
| 7:24.427 | תלריג |  |
| 7:24.428 | ויראו |  |
| 7:24.429 | בני |  |
| 7:24.430 | האלהים |  |
| 7:24.431 | את |  |
| 7:24.432 | בנות |  |
| 7:24.433 | האדם | . |
| 7:24.434 | אפצע |  |
| 7:24.435 | מן |  |
| 7:24.436 | מא |  |
| 7:24.437 | קאלה |  |
| 7:24.438 | פי |  |
| 7:24.439 | קצה |  |
| 7:24.440 | איוב |  |
| 7:24.441 | ודאלך |  |
| 7:24.442 | אנא |  |
| 7:24.443 | לם |  |
| 7:24.444 | יקף |  |
| 7:24.445 | בתה |  |
| 7:24.446 | עלי | ' |
| 7:24.447 | מא |  |
| 7:24.448 | יסתעמלון |  |
| 7:24.449 | בני |  |
| 7:24.450 | אסראיל |  |
| 7:24.451 | פי |  |
| 7:24.452 | אלכאץ |  |
| 7:24.453 | ואלעאם | , |
| 7:24.454 | ודאלך |  |
| 7:24.455 | אנה |  |
| 7:24.456 | מן |  |
| 7:24.457 | סבילהם |  |
| 7:24.458 | אן |  |
| 7:24.459 | יקולון |  |
| 7:24.460 | פעל |  |
| 7:24.461 | אלכואץ |  |
| 7:24.462 | כדא |  |
| 7:24.463 | ופעל |  |
| 7:24.464 | אלעואם |  |
| 7:24.465 | כדא |  |
| 7:24.466 | ויכון |  |
| 7:24.467 | אולאיך |  |
| 7:24.468 | אלכואץ |  |
| 7:24.469 | איצא |  |
| 7:24.470 | נאנא | , |
| 7:24.471 | ואנמא |  |
| 7:24.472 | פרדהם |  |
| 7:24.473 | לתכציציהם | , |
| 7:24.474 | פעלי | , |
| 7:24.475 | הדא |  |
| 7:24.476 | קולה |  |
| 7:24.477 | ויראו |  |
| 7:24.478 | בני |  |
| 7:24.479 | האלהים |  |
| 7:24.480 | נאסא |  |
| 7:24.481 | איצא |  |
| 7:24.482 | אנמא |  |
| 7:24.483 | פרדהם |  |
| 7:24.484 | מן |  |
| 7:24.485 | אלעואם |  |
| 7:24.486 | לתכציצהם |  |
| 7:24.487 | ושרפהם | . |
| 7:24.488 | והדה |  |
| 7:24.489 | עבארה |  |
| 7:24.490 | מתגהה |  |
| 7:24.491 | לתנפי |  |
| 7:24.492 | אלמנכר |  |
| 7:24.493 | ען |  |
| 7:24.494 | אלמקרבין | . |
| 7:24.495 | ויזידהא |  |
| 7:24.496 | איצא |  |
| 7:24.497 | ויקויהא |  |
| 7:24.498 | קולה |  |
| 7:24.499 | בעד |  |
| 7:24.500 | דאלך |  |
| 7:24.501 | וירא |  |
| 7:24.502 | ה | ' |
| 7:24.503 | כי |  |
| 7:24.504 | רבה |  |
| 7:24.505 | רעת |  |
| 7:24.506 | האדם | , |
| 7:24.507 | ולם |  |
| 7:24.508 | ידם |  |
| 7:24.509 | אלמלאיכה |  |
| 7:24.510 | מעהם | . |
| 7:24.511 | וקולה |  |
| 7:24.512 | פי |  |
| 7:24.513 | עקב |  |
| 7:24.514 | דאלך |  |
| 7:24.515 | אמחה |  |
| 7:24.516 | את |  |
| 7:24.517 | האדם |  |
| 7:24.518 | אשר |  |
| 7:24.519 | בראתי |  |
| 7:24.520 | ולם |  |
| 7:24.521 | יעאקב |  |
| 7:24.522 | אלמלאיכה |  |
| 7:24.523 | וכו | ' |
| 7:24.524 | עכ"ל | . |
| 7:24.525 | העתק |  |
| 7:24.526 | תמצא |  |
| 7:24.527 | בסוף |  |
| 7:24.528 | הספר |  |
| 7:24.529 | הרי |  |
| 7:24.530 | בהדיא |  |
| 7:24.531 | שהגאון |  |
| 7:24.532 | מרחיק |  |
| 7:24.533 | מאד |  |
| 7:24.534 | דעת |  |
| 7:24.535 | האומר |  |
| 7:24.536 | שבני |  |
| 7:24.537 | האלהים |  |
| 7:24.538 | הם |  |
| 7:24.539 | מלאכים |  |
| 7:24.540 | ודוחה |  |
| 7:24.541 | אותה |  |
| 7:24.542 | בשתי |  |
| 7:24.543 | ידים |  |
| 7:24.544 | וקורא |  |
| 7:24.545 | לה |  |
| 7:24.546 | דעת |  |
| 7:24.547 | חיצונית | . |
| 7:24.548 | ודבריו |  |
| 7:24.549 | מכוונים |  |
| 7:24.550 | לדברי |  |
| 7:24.551 | הרשב"י |  |
| 7:24.552 | שבבראשית |  |
| 7:24.553 | רבה | , |
| 7:24.554 | שלא |  |
| 7:24.555 | כדעת |  |
| 7:24.556 | הזהר |  |
| 7:24.557 | שאמר |  |
| 7:24.558 | שהם |  |
| 7:24.559 | מלאכים | , |
| 7:24.560 | ושונה |  |
| 7:24.561 | ומשלם |  |
| 7:24.562 | בה | , |
| 7:24.563 | וזו |  |
| 7:24.564 | הוכחה |  |
| 7:24.565 | גדולה |  |
| 7:24.566 | שאינו |  |
| 7:24.567 | מרשב"י |  |
| 7:24.568 | שהרי |  |
| 7:24.569 | רשב"י |  |
| 7:24.570 | היה |  |
| 7:24.571 | מקלל |  |
| 7:24.572 | את |  |
| 7:24.573 | האומר |  |
| 7:24.574 | שהם |  |
| 7:24.575 | מלאכים | . |
| 7:24.576 | ובענין |  |
| 7:24.577 | השטן |  |
| 7:24.578 | שפי | ' |
| 7:24.579 | הגאון |  |
| 7:24.580 | שהוא |  |
| 7:24.581 | אדם | , |
| 7:24.582 | אע"פ |  |
| 7:24.583 | שהראב"ע |  |
| 7:24.584 | השיגו |  |
| 7:24.585 | יש |  |
| 7:24.586 | מקום |  |
| 7:24.587 | סתירה |  |
| 7:24.588 | להשגתו |  |
| 7:24.589 | ולהעמיד |  |
| 7:24.590 | דברי |  |
| 7:24.591 | הגאון |  |
| 7:24.592 | הרס"ג | , |
| 7:24.593 | ואע"פ |  |
| 7:24.594 | שבפ"ק |  |
| 7:24.595 | דבבא |  |
| 7:24.596 | בתרא |  |
| 7:24.597 | משמע |  |
| 7:24.598 | כהראב"ע | . |
| 7:24.599 | הנה |  |
| 7:24.600 | ידוע |  |
| 7:24.601 | מה |  |
| 7:24.602 | שכ | ' |
| 7:24.603 | הנגיד |  |
| 7:24.604 | בעל |  |
| 7:24.605 | מבוא |  |
| 7:24.606 | התלמוד |  |
| 7:24.607 | בפי | ' |
| 7:24.608 | הגדה | , |
| 7:24.609 | ואין |  |
| 7:24.610 | לך |  |
| 7:24.611 | ללמוד |  |
| 7:24.612 | ממנה |  |
| 7:24.613 | אלא |  |
| 7:24.614 | מה |  |
| 7:24.615 | שיעלה |  |
| 7:24.616 | על |  |
| 7:24.617 | הדעת | , |
| 7:24.618 | ויש |  |
| 7:24.619 | לך |  |
| 7:24.620 | לדעת |  |
| 7:24.621 | שכל |  |
| 7:24.622 | מה |  |
| 7:24.623 | שקיימו |  |
| 7:24.624 | חז"ל |  |
| 7:24.625 | הלכה |  |
| 7:24.626 | בענין |  |
| 7:24.627 | מצוה | , |
| 7:24.628 | שהוא |  |
| 7:24.629 | מפי |  |
| 7:24.630 | משה |  |
| 7:24.631 | רבינו |  |
| 7:24.632 | עליו |  |
| 7:24.633 | השלום |  |
| 7:24.634 | מפי |  |
| 7:24.635 | הגבורה | , |
| 7:24.636 | אין |  |
| 7:24.637 | לך |  |
| 7:24.638 | להוסיף |  |
| 7:24.639 | עליו |  |
| 7:24.640 | ולא |  |
| 7:24.641 | לגרוע |  |
| 7:24.642 | ממנו |  |
| 7:24.643 | אבל |  |
| 7:24.644 | מה |  |
| 7:24.645 | שפירשו |  |
| 7:24.646 | בפסוקים |  |
| 7:24.647 | כל |  |
| 7:24.648 | אחד |  |
| 7:24.649 | כפי |  |
| 7:24.650 | שיזדמן |  |
| 7:24.651 | לו |  |
| 7:24.652 | ומה |  |
| 7:24.653 | שראה |  |
| 7:24.654 | בדעתו |  |
| 7:24.655 | ולפי |  |
| 7:24.656 | מה |  |
| 7:24.657 | שיעלה |  |
| 7:24.658 | על |  |
| 7:24.659 | הדעת |  |
| 7:24.660 | מן |  |
| 7:24.661 | הפירושים |  |
| 7:24.662 | האלה |  |
| 7:24.663 | לומדים |  |
| 7:24.664 | אותם | , |
| 7:24.665 | והשאר |  |
| 7:24.666 | אין |  |
| 7:24.667 | סומכים |  |
| 7:24.668 | עליהם |  |
| 7:24.669 | עכ"ל | . |
| 7:24.670 | ועיין |  |
| 7:24.671 | רש"י |  |
| 7:24.672 | שבת |  |
| 7:24.673 | פ | ' |
| 7:24.674 | במה |  |
| 7:24.675 | מדליקים |  |
| 7:24.676 | בשאלת |  |
| 7:24.677 | מהו |  |
| 7:24.678 | לכבויי |  |
| 7:24.679 | בוצינא |  |
| 7:24.680 | מקמי |  |
| 7:24.681 | באישא |  |
| 7:24.682 | בשבתא | . |
| 7:24.683 | והיוצא |  |
| 7:24.684 | מזה |  |
| 7:24.685 | דכל |  |
| 7:24.686 | למסבר |  |
| 7:24.687 | קראי |  |
| 7:24.688 | בעלמא |  |
| 7:24.689 | ואין |  |
| 7:24.690 | נוגע |  |
| 7:24.691 | ממנו |  |
| 7:24.692 | לדינא |  |
| 7:24.693 | בעשיית |  |
| 7:24.694 | המצות |  |
| 7:24.695 | ובעיקרי |  |
| 7:24.696 | תורתינו |  |
| 7:24.697 | הקדושה | , |
| 7:24.698 | אין |  |
| 7:24.699 | בומכים |  |
| 7:24.700 | אלא |  |
| 7:24.701 | על |  |
| 7:24.702 | כל |  |
| 7:24.703 | מאי |  |
| 7:24.704 | דמסתבר |  |
| 7:24.705 | טפי | . |
| 7:24.706 | אבל |  |
| 7:24.707 | כל |  |
| 7:24.708 | מאי |  |
| 7:24.709 | דפנקא |  |
| 7:24.710 | מינא |  |
| 7:24.711 | לדינא |  |
| 7:24.712 | כנדון |  |
| 7:24.713 | דידן |  |
| 7:24.714 | כמצות |  |
| 7:24.715 | יחוד |  |
| 7:24.716 | השם |  |
| 7:24.717 | יתברך |  |
| 7:24.718 | אין |  |
| 7:24.719 | לזוז |  |
| 7:24.720 | מקבלת |  |
| 7:24.721 | חז"ל |  |
| 7:24.722 | שאין |  |
| 7:24.723 | לשתף |  |
| 7:24.724 | עם |  |
| 7:24.725 | השם |  |
| 7:24.726 | ית | ' |
| 7:24.727 | שום |  |
| 7:24.728 | כח |  |
| 7:24.729 | נברא |  |
| 7:24.730 | בין |  |
| 7:24.731 | שיהיה |  |
| 7:24.732 | גשמי |  |
| 7:24.733 | או |  |
| 7:24.734 | רוחני |  |
| 7:24.735 | כמאמר |  |
| 7:24.736 | רשב"י |  |
| 7:24.737 | ע"ח |  |
| 7:24.738 | כל |  |
| 7:24.739 | המשתף |  |
| 7:24.740 | שם |  |
| 7:24.741 | שמים |  |
| 7:24.742 | ודבר |  |
| 7:24.743 | אחר |  |
| 7:24.744 | נעקר |  |
| 7:24.745 | מן |  |
| 7:24.746 | העולם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

