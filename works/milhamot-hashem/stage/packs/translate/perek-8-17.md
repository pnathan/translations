# Translation pack: milhamot-hashem perek-8-17 (Milchamot ha-Shem 8:20–8:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-17`.

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

- סתו"מ: always "Torah scrolls, tefillin and mezuzot", with a note.
- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
- משכילים / השכלה: "the educated" / "enlightenment" in the sense of the
  Haskalah only where the paragraph speaks of that movement; otherwise
  "those of understanding".
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- והוא שלא: "provided that" (fixed).
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- The words of a verse quoted in the plural (Job 32:16, עמדו): plural, no
  subject added.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 8:20

Hebrew (Wikisource): קיג) ולכן אמרו רז"ל על אנשי כנה"ג שהחזירו את העטרה לישנה באמרם האל הגדול הגבור והנורא, ושכל המאורעות שאירעו לישראל מטוב ועד רע, הן הן גבורותיו ונוראותיו של אלהים עליון, והוא אחד אשר אין יחיד כיחודו. ונמשכה ונתפשטה אמונת אחדותו ית' הטהורה בכל ישראל כל משך ימי האלף הרביעי והחמישי ומעט מזער מהאלף השישי, עד שעברו איזה שנים מאלף הששי, וינסה הקב"ה את ישראל במרמת המלך העכו"ם שהוציא את ספר הזהר כאלו נתחבר ע"י רשב"י, ונפתו רבים מחכמינו לדבריו, וההמינו לשקריו לעשות לנו תורה שבעל פה חדשה, ואלהים חדשים, ולדבר תועה על המשנה והתלמוד שהם מטוב ורע מעורבים, ולא מאת השם אלא ממטטרון כדלעיל. והרבה מחכמי ישראל לא עמדו בנסיון ויאמינו בכמה דעות מדומות ונתפשטו בישראל אצל העוסקים בו, וחברו באמונה הזאת ספרים, ונספחו אל המשנה והתלמוד, כאלו הם שוים באמונת האלהות, ולא ביחנו יפה להבין דבריו אשר ידבר במליצות סתומות ומושאלות ובלעני שפה. אמנם רבותינו התלמודיים הדבקים בעסק המשנה והתלמוד ובמדרשי רז"ל וגם המון העם אשר לא ידעו דבר מהקבלה החדשה, הנם עומדים באמונתם השלימה וטהורה מכל סיג הלאה, וכל שכן האנשים המעיינים בשום שכל בדברי הרס"ג, והכוזרי, והרמב"ם, והח"ל, וכיוצא בהם, כולם כאחד תקועים באמונת האחדות האמתית. ועטרת אלהינו ברוך הוא עומדת אצלם בישנה כבימי אנשי כנה"ג. אך דא עקא. כי המה מאמינים בספרי הקבלה החדשה הצבועים למראה עינים. כי קדושים המה ושיש בהם דעות נשגבות אשר לא תכיל דעתם להשיגם. ושהמה שלובי יד עם האמונה הטהורה. ואינם מרגישים במרורות הפתנים אשר בקרבה. העוקרים יסודות תוה"ק. ועמודיה יתפלצון. ורבים מחכמינו שנו בה. ויהיו לאבן מכשול לבאים אחריהם.

| id | word | punctuation after |
|---|---|---|
| 8:20.1 | קיג | ) |
| 8:20.2 | ולכן |  |
| 8:20.3 | אמרו |  |
| 8:20.4 | רז"ל |  |
| 8:20.5 | על |  |
| 8:20.6 | אנשי |  |
| 8:20.7 | כנה"ג |  |
| 8:20.8 | שהחזירו |  |
| 8:20.9 | את |  |
| 8:20.10 | העטרה |  |
| 8:20.11 | לישנה |  |
| 8:20.12 | באמרם |  |
| 8:20.13 | האל |  |
| 8:20.14 | הגדול |  |
| 8:20.15 | הגבור |  |
| 8:20.16 | והנורא | , |
| 8:20.17 | ושכל |  |
| 8:20.18 | המאורעות |  |
| 8:20.19 | שאירעו |  |
| 8:20.20 | לישראל |  |
| 8:20.21 | מטוב |  |
| 8:20.22 | ועד |  |
| 8:20.23 | רע | , |
| 8:20.24 | הן |  |
| 8:20.25 | הן |  |
| 8:20.26 | גבורותיו |  |
| 8:20.27 | ונוראותיו |  |
| 8:20.28 | של |  |
| 8:20.29 | אלהים |  |
| 8:20.30 | עליון | , |
| 8:20.31 | והוא |  |
| 8:20.32 | אחד |  |
| 8:20.33 | אשר |  |
| 8:20.34 | אין |  |
| 8:20.35 | יחיד |  |
| 8:20.36 | כיחודו | . |
| 8:20.37 | ונמשכה |  |
| 8:20.38 | ונתפשטה |  |
| 8:20.39 | אמונת |  |
| 8:20.40 | אחדותו |  |
| 8:20.41 | ית | ' |
| 8:20.42 | הטהורה |  |
| 8:20.43 | בכל |  |
| 8:20.44 | ישראל |  |
| 8:20.45 | כל |  |
| 8:20.46 | משך |  |
| 8:20.47 | ימי |  |
| 8:20.48 | האלף |  |
| 8:20.49 | הרביעי |  |
| 8:20.50 | והחמישי |  |
| 8:20.51 | ומעט |  |
| 8:20.52 | מזער |  |
| 8:20.53 | מהאלף |  |
| 8:20.54 | השישי | , |
| 8:20.55 | עד |  |
| 8:20.56 | שעברו |  |
| 8:20.57 | איזה |  |
| 8:20.58 | שנים |  |
| 8:20.59 | מאלף |  |
| 8:20.60 | הששי | , |
| 8:20.61 | וינסה |  |
| 8:20.62 | הקב"ה |  |
| 8:20.63 | את |  |
| 8:20.64 | ישראל |  |
| 8:20.65 | במרמת |  |
| 8:20.66 | המלך |  |
| 8:20.67 | העכו"ם |  |
| 8:20.68 | שהוציא |  |
| 8:20.69 | את |  |
| 8:20.70 | ספר |  |
| 8:20.71 | הזהר |  |
| 8:20.72 | כאלו |  |
| 8:20.73 | נתחבר |  |
| 8:20.74 | ע"י |  |
| 8:20.75 | רשב"י | , |
| 8:20.76 | ונפתו |  |
| 8:20.77 | רבים |  |
| 8:20.78 | מחכמינו |  |
| 8:20.79 | לדבריו | , |
| 8:20.80 | וההמינו |  |
| 8:20.81 | לשקריו |  |
| 8:20.82 | לעשות |  |
| 8:20.83 | לנו |  |
| 8:20.84 | תורה |  |
| 8:20.85 | שבעל |  |
| 8:20.86 | פה |  |
| 8:20.87 | חדשה | , |
| 8:20.88 | ואלהים |  |
| 8:20.89 | חדשים | , |
| 8:20.90 | ולדבר |  |
| 8:20.91 | תועה |  |
| 8:20.92 | על |  |
| 8:20.93 | המשנה |  |
| 8:20.94 | והתלמוד |  |
| 8:20.95 | שהם |  |
| 8:20.96 | מטוב |  |
| 8:20.97 | ורע |  |
| 8:20.98 | מעורבים | , |
| 8:20.99 | ולא |  |
| 8:20.100 | מאת |  |
| 8:20.101 | השם |  |
| 8:20.102 | אלא |  |
| 8:20.103 | ממטטרון |  |
| 8:20.104 | כדלעיל | . |
| 8:20.105 | והרבה |  |
| 8:20.106 | מחכמי |  |
| 8:20.107 | ישראל |  |
| 8:20.108 | לא |  |
| 8:20.109 | עמדו |  |
| 8:20.110 | בנסיון |  |
| 8:20.111 | ויאמינו |  |
| 8:20.112 | בכמה |  |
| 8:20.113 | דעות |  |
| 8:20.114 | מדומות |  |
| 8:20.115 | ונתפשטו |  |
| 8:20.116 | בישראל |  |
| 8:20.117 | אצל |  |
| 8:20.118 | העוסקים |  |
| 8:20.119 | בו | , |
| 8:20.120 | וחברו |  |
| 8:20.121 | באמונה |  |
| 8:20.122 | הזאת |  |
| 8:20.123 | ספרים | , |
| 8:20.124 | ונספחו |  |
| 8:20.125 | אל |  |
| 8:20.126 | המשנה |  |
| 8:20.127 | והתלמוד | , |
| 8:20.128 | כאלו |  |
| 8:20.129 | הם |  |
| 8:20.130 | שוים |  |
| 8:20.131 | באמונת |  |
| 8:20.132 | האלהות | , |
| 8:20.133 | ולא |  |
| 8:20.134 | ביחנו |  |
| 8:20.135 | יפה |  |
| 8:20.136 | להבין |  |
| 8:20.137 | דבריו |  |
| 8:20.138 | אשר |  |
| 8:20.139 | ידבר |  |
| 8:20.140 | במליצות |  |
| 8:20.141 | סתומות |  |
| 8:20.142 | ומושאלות |  |
| 8:20.143 | ובלעני |  |
| 8:20.144 | שפה | . |
| 8:20.145 | אמנם |  |
| 8:20.146 | רבותינו |  |
| 8:20.147 | התלמודיים |  |
| 8:20.148 | הדבקים |  |
| 8:20.149 | בעסק |  |
| 8:20.150 | המשנה |  |
| 8:20.151 | והתלמוד |  |
| 8:20.152 | ובמדרשי |  |
| 8:20.153 | רז"ל |  |
| 8:20.154 | וגם |  |
| 8:20.155 | המון |  |
| 8:20.156 | העם |  |
| 8:20.157 | אשר |  |
| 8:20.158 | לא |  |
| 8:20.159 | ידעו |  |
| 8:20.160 | דבר |  |
| 8:20.161 | מהקבלה |  |
| 8:20.162 | החדשה | , |
| 8:20.163 | הנם |  |
| 8:20.164 | עומדים |  |
| 8:20.165 | באמונתם |  |
| 8:20.166 | השלימה |  |
| 8:20.167 | וטהורה |  |
| 8:20.168 | מכל |  |
| 8:20.169 | סיג |  |
| 8:20.170 | הלאה | , |
| 8:20.171 | וכל |  |
| 8:20.172 | שכן |  |
| 8:20.173 | האנשים |  |
| 8:20.174 | המעיינים |  |
| 8:20.175 | בשום |  |
| 8:20.176 | שכל |  |
| 8:20.177 | בדברי |  |
| 8:20.178 | הרס"ג | , |
| 8:20.179 | והכוזרי | , |
| 8:20.180 | והרמב"ם | , |
| 8:20.181 | והח"ל | , |
| 8:20.182 | וכיוצא |  |
| 8:20.183 | בהם | , |
| 8:20.184 | כולם |  |
| 8:20.185 | כאחד |  |
| 8:20.186 | תקועים |  |
| 8:20.187 | באמונת |  |
| 8:20.188 | האחדות |  |
| 8:20.189 | האמתית | . |
| 8:20.190 | ועטרת |  |
| 8:20.191 | אלהינו |  |
| 8:20.192 | ברוך |  |
| 8:20.193 | הוא |  |
| 8:20.194 | עומדת |  |
| 8:20.195 | אצלם |  |
| 8:20.196 | בישנה |  |
| 8:20.197 | כבימי |  |
| 8:20.198 | אנשי |  |
| 8:20.199 | כנה"ג | . |
| 8:20.200 | אך |  |
| 8:20.201 | דא |  |
| 8:20.202 | עקא | . |
| 8:20.203 | כי |  |
| 8:20.204 | המה |  |
| 8:20.205 | מאמינים |  |
| 8:20.206 | בספרי |  |
| 8:20.207 | הקבלה |  |
| 8:20.208 | החדשה |  |
| 8:20.209 | הצבועים |  |
| 8:20.210 | למראה |  |
| 8:20.211 | עינים | . |
| 8:20.212 | כי |  |
| 8:20.213 | קדושים |  |
| 8:20.214 | המה |  |
| 8:20.215 | ושיש |  |
| 8:20.216 | בהם |  |
| 8:20.217 | דעות |  |
| 8:20.218 | נשגבות |  |
| 8:20.219 | אשר |  |
| 8:20.220 | לא |  |
| 8:20.221 | תכיל |  |
| 8:20.222 | דעתם |  |
| 8:20.223 | להשיגם | . |
| 8:20.224 | ושהמה |  |
| 8:20.225 | שלובי |  |
| 8:20.226 | יד |  |
| 8:20.227 | עם |  |
| 8:20.228 | האמונה |  |
| 8:20.229 | הטהורה | . |
| 8:20.230 | ואינם |  |
| 8:20.231 | מרגישים |  |
| 8:20.232 | במרורות |  |
| 8:20.233 | הפתנים |  |
| 8:20.234 | אשר |  |
| 8:20.235 | בקרבה | . |
| 8:20.236 | העוקרים |  |
| 8:20.237 | יסודות |  |
| 8:20.238 | תוה"ק | . |
| 8:20.239 | ועמודיה |  |
| 8:20.240 | יתפלצון | . |
| 8:20.241 | ורבים |  |
| 8:20.242 | מחכמינו |  |
| 8:20.243 | שנו |  |
| 8:20.244 | בה | . |
| 8:20.245 | ויהיו |  |
| 8:20.246 | לאבן |  |
| 8:20.247 | מכשול |  |
| 8:20.248 | לבאים |  |
| 8:20.249 | אחריהם | . |


## Milchamot ha-Shem 8:21

Hebrew (Wikisource): קיד) ורבים מהמעיינים יודעים ומעידים ומגידים בצנעה ובהסתר גדול כי זרה היא עד מאוד לאמונת תורתינו באמונת יחוד השי"ת. ולא יחפצו לגלות דעתם, מיראת החסידים והמתחסדים אוהבי ורודפי כבוד וגדולה. העושים כוונות למלכת השמים המדומה אשר שמו ברעיונותם בסתרו של עולם. שלא הועילה אצלם תפלת אנשי כנה"ג ותקנותם וכל יסודותם רק שלא לעשות צלמים ומסכות כסף וזהב ועץ ואבן מעשה ידי חרש. כמו שפתח מחבר הזהר לומר ארור האיש אשר יעשה פסל ומסכה וכו' מעשה ידי חרש. אבל במחשבתם יש הרבה צלמים ובעלים ועשתרות העובדים ומשתחוים להם. ומתפללים אליהם בטהרה. וקורים להם בכל השמות המיוחדים לשי"ת. וכוונות עושים להם לחברם כמכנס חבלים רבים ומושך גמלים. ואומרים ובולא חד. והמשכילים המאמינים באחדותו ית' על פי האמת והצדק. ומחזיקים בתורת האמת בכתב ובע"פ. יראים לנפשם להגיד ולהעלות על דל שפתם את האמת הגמור. מפני תרמית האמת אשר בפי המתחסדים פן יפגעו בהם ויתנפלו עליהם בנידוים וחרמות כתוקף ובחזק יד שלא כרצונו וכרצון יראיו, כנפילים אשר מעולם אנשי השם. ומפחד המשכילים מהם. המה יראים למסור כבודם על קדושת השם הנכבד והנורא. הוא ה' אלהינו אשר בכל הארץ משפטיו. והמה מעלימים מעיניהם מעשה אבותינו שמסרו נפשם וגופם על קדושת השם. אברהם אבינו הולשך לאור כשדים. חנניה מיקאל לועזריה מסרו עצמן והושלכו לכבשן האש. דניאל הושלך לגוב אריות, וכמה צדיקים וחסידים המסופר עליהם בגמרא ומדרשי רז"ל שמסרו עצמן על קדושת השם. יש מהן שניצולו. ויש שנתפסו ונהרגו בעון דוריהם.

| id | word | punctuation after |
|---|---|---|
| 8:21.1 | קיד | ) |
| 8:21.2 | ורבים |  |
| 8:21.3 | מהמעיינים |  |
| 8:21.4 | יודעים |  |
| 8:21.5 | ומעידים |  |
| 8:21.6 | ומגידים |  |
| 8:21.7 | בצנעה |  |
| 8:21.8 | ובהסתר |  |
| 8:21.9 | גדול |  |
| 8:21.10 | כי |  |
| 8:21.11 | זרה |  |
| 8:21.12 | היא |  |
| 8:21.13 | עד |  |
| 8:21.14 | מאוד |  |
| 8:21.15 | לאמונת |  |
| 8:21.16 | תורתינו |  |
| 8:21.17 | באמונת |  |
| 8:21.18 | יחוד |  |
| 8:21.19 | השי"ת | . |
| 8:21.20 | ולא |  |
| 8:21.21 | יחפצו |  |
| 8:21.22 | לגלות |  |
| 8:21.23 | דעתם | , |
| 8:21.24 | מיראת |  |
| 8:21.25 | החסידים |  |
| 8:21.26 | והמתחסדים |  |
| 8:21.27 | אוהבי |  |
| 8:21.28 | ורודפי |  |
| 8:21.29 | כבוד |  |
| 8:21.30 | וגדולה | . |
| 8:21.31 | העושים |  |
| 8:21.32 | כוונות |  |
| 8:21.33 | למלכת |  |
| 8:21.34 | השמים |  |
| 8:21.35 | המדומה |  |
| 8:21.36 | אשר |  |
| 8:21.37 | שמו |  |
| 8:21.38 | ברעיונותם |  |
| 8:21.39 | בסתרו |  |
| 8:21.40 | של |  |
| 8:21.41 | עולם | . |
| 8:21.42 | שלא |  |
| 8:21.43 | הועילה |  |
| 8:21.44 | אצלם |  |
| 8:21.45 | תפלת |  |
| 8:21.46 | אנשי |  |
| 8:21.47 | כנה"ג |  |
| 8:21.48 | ותקנותם |  |
| 8:21.49 | וכל |  |
| 8:21.50 | יסודותם |  |
| 8:21.51 | רק |  |
| 8:21.52 | שלא |  |
| 8:21.53 | לעשות |  |
| 8:21.54 | צלמים |  |
| 8:21.55 | ומסכות |  |
| 8:21.56 | כסף |  |
| 8:21.57 | וזהב |  |
| 8:21.58 | ועץ |  |
| 8:21.59 | ואבן |  |
| 8:21.60 | מעשה |  |
| 8:21.61 | ידי |  |
| 8:21.62 | חרש | . |
| 8:21.63 | כמו |  |
| 8:21.64 | שפתח |  |
| 8:21.65 | מחבר |  |
| 8:21.66 | הזהר |  |
| 8:21.67 | לומר |  |
| 8:21.68 | ארור |  |
| 8:21.69 | האיש |  |
| 8:21.70 | אשר |  |
| 8:21.71 | יעשה |  |
| 8:21.72 | פסל |  |
| 8:21.73 | ומסכה |  |
| 8:21.74 | וכו | ' |
| 8:21.75 | מעשה |  |
| 8:21.76 | ידי |  |
| 8:21.77 | חרש | . |
| 8:21.78 | אבל |  |
| 8:21.79 | במחשבתם |  |
| 8:21.80 | יש |  |
| 8:21.81 | הרבה |  |
| 8:21.82 | צלמים |  |
| 8:21.83 | ובעלים |  |
| 8:21.84 | ועשתרות |  |
| 8:21.85 | העובדים |  |
| 8:21.86 | ומשתחוים |  |
| 8:21.87 | להם | . |
| 8:21.88 | ומתפללים |  |
| 8:21.89 | אליהם |  |
| 8:21.90 | בטהרה | . |
| 8:21.91 | וקורים |  |
| 8:21.92 | להם |  |
| 8:21.93 | בכל |  |
| 8:21.94 | השמות |  |
| 8:21.95 | המיוחדים |  |
| 8:21.96 | לשי"ת | . |
| 8:21.97 | וכוונות |  |
| 8:21.98 | עושים |  |
| 8:21.99 | להם |  |
| 8:21.100 | לחברם |  |
| 8:21.101 | כמכנס |  |
| 8:21.102 | חבלים |  |
| 8:21.103 | רבים |  |
| 8:21.104 | ומושך |  |
| 8:21.105 | גמלים | . |
| 8:21.106 | ואומרים |  |
| 8:21.107 | ובולא |  |
| 8:21.108 | חד | . |
| 8:21.109 | והמשכילים |  |
| 8:21.110 | המאמינים |  |
| 8:21.111 | באחדותו |  |
| 8:21.112 | ית | ' |
| 8:21.113 | על |  |
| 8:21.114 | פי |  |
| 8:21.115 | האמת |  |
| 8:21.116 | והצדק | . |
| 8:21.117 | ומחזיקים |  |
| 8:21.118 | בתורת |  |
| 8:21.119 | האמת |  |
| 8:21.120 | בכתב |  |
| 8:21.121 | ובע"פ | . |
| 8:21.122 | יראים |  |
| 8:21.123 | לנפשם |  |
| 8:21.124 | להגיד |  |
| 8:21.125 | ולהעלות |  |
| 8:21.126 | על |  |
| 8:21.127 | דל |  |
| 8:21.128 | שפתם |  |
| 8:21.129 | את |  |
| 8:21.130 | האמת |  |
| 8:21.131 | הגמור | . |
| 8:21.132 | מפני |  |
| 8:21.133 | תרמית |  |
| 8:21.134 | האמת |  |
| 8:21.135 | אשר |  |
| 8:21.136 | בפי |  |
| 8:21.137 | המתחסדים |  |
| 8:21.138 | פן |  |
| 8:21.139 | יפגעו |  |
| 8:21.140 | בהם |  |
| 8:21.141 | ויתנפלו |  |
| 8:21.142 | עליהם |  |
| 8:21.143 | בנידוים |  |
| 8:21.144 | וחרמות |  |
| 8:21.145 | כתוקף |  |
| 8:21.146 | ובחזק |  |
| 8:21.147 | יד |  |
| 8:21.148 | שלא |  |
| 8:21.149 | כרצונו |  |
| 8:21.150 | וכרצון |  |
| 8:21.151 | יראיו | , |
| 8:21.152 | כנפילים |  |
| 8:21.153 | אשר |  |
| 8:21.154 | מעולם |  |
| 8:21.155 | אנשי |  |
| 8:21.156 | השם | . |
| 8:21.157 | ומפחד |  |
| 8:21.158 | המשכילים |  |
| 8:21.159 | מהם | . |
| 8:21.160 | המה |  |
| 8:21.161 | יראים |  |
| 8:21.162 | למסור |  |
| 8:21.163 | כבודם |  |
| 8:21.164 | על |  |
| 8:21.165 | קדושת |  |
| 8:21.166 | השם |  |
| 8:21.167 | הנכבד |  |
| 8:21.168 | והנורא | . |
| 8:21.169 | הוא |  |
| 8:21.170 | ה | ' |
| 8:21.171 | אלהינו |  |
| 8:21.172 | אשר |  |
| 8:21.173 | בכל |  |
| 8:21.174 | הארץ |  |
| 8:21.175 | משפטיו | . |
| 8:21.176 | והמה |  |
| 8:21.177 | מעלימים |  |
| 8:21.178 | מעיניהם |  |
| 8:21.179 | מעשה |  |
| 8:21.180 | אבותינו |  |
| 8:21.181 | שמסרו |  |
| 8:21.182 | נפשם |  |
| 8:21.183 | וגופם |  |
| 8:21.184 | על |  |
| 8:21.185 | קדושת |  |
| 8:21.186 | השם | . |
| 8:21.187 | אברהם |  |
| 8:21.188 | אבינו |  |
| 8:21.189 | הולשך |  |
| 8:21.190 | לאור |  |
| 8:21.191 | כשדים | . |
| 8:21.192 | חנניה |  |
| 8:21.193 | מיקאל |  |
| 8:21.194 | לועזריה |  |
| 8:21.195 | מסרו |  |
| 8:21.196 | עצמן |  |
| 8:21.197 | והושלכו |  |
| 8:21.198 | לכבשן |  |
| 8:21.199 | האש | . |
| 8:21.200 | דניאל |  |
| 8:21.201 | הושלך |  |
| 8:21.202 | לגוב |  |
| 8:21.203 | אריות | , |
| 8:21.204 | וכמה |  |
| 8:21.205 | צדיקים |  |
| 8:21.206 | וחסידים |  |
| 8:21.207 | המסופר |  |
| 8:21.208 | עליהם |  |
| 8:21.209 | בגמרא |  |
| 8:21.210 | ומדרשי |  |
| 8:21.211 | רז"ל |  |
| 8:21.212 | שמסרו |  |
| 8:21.213 | עצמן |  |
| 8:21.214 | על |  |
| 8:21.215 | קדושת |  |
| 8:21.216 | השם | . |
| 8:21.217 | יש |  |
| 8:21.218 | מהן |  |
| 8:21.219 | שניצולו | . |
| 8:21.220 | ויש |  |
| 8:21.221 | שנתפסו |  |
| 8:21.222 | ונהרגו |  |
| 8:21.223 | בעון |  |
| 8:21.224 | דוריהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

