# Translation pack: milhamot-hashem perek-5-06 (Milchamot ha-Shem 5:11–5:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-06`.

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
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- באצילות: "in Emanation" (the noun the Kabbalists use), without "World of".
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
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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


## Milchamot ha-Shem 5:11

Hebrew (Wikisource): מד) ואחר שידעת את כל אלה בא וראה איך שגגו רבותינו המקובלים והטילו דופי בידיעת השי"ת בעתידות, גם סלקו ההשגחה מהמלך העליון ותלו ההשגחה בזעיר אנפין לבד, ולא באשר הפרצופין אבא ואמא, ואריך, ועתיק, ואדם קדמון, וכ"ש באין סוף העליון, ואמרו שאין מקפידין ומשגיחין במעשה התחתונים בין טוב לרע ח"ו כטוב כחוטא, והיה כצדיק כרשע, הנשבע כאשר שבועה ירא כל אפיא שוין. ודבר זה מבואר בכתבי האר"י הובא בספר ד' מאות שקל כסף דף צ"א ע"ב ז"ל שאלה ששאל הרב המובהק המקובל האלהי כמהר"ר אברהם מונצוץ מעיר טניפיל להאר"י זצוק"ל יש לשאול דבר אחד קשה והוא כי ראינו בספרי הקבלה והזהר ובספר המדע ובספר הקנה כי בשעה שהאל מתעסק בבריאה, אינו יודע. וכשיהיה באצילות יודע. אם כן הרי נתבאר שבעולם האצילות היא הידיעה והכריחה, א"כ קשה לדברי רז"ל על פסוק ויאמר ה' אל משה אמר אל הכוהנים בני אהרן, מלך ראשון שיעמוד לישראל ידקר בחרב. אמר לו ולי אתה שואל, אמור אל הכוהנים בני אהרן שהרי הרג נוב עיר הכהנים. ונמצא שהכריחה דעתו של שאול, באיזה אופן שיהיה אם בבריאה אם באצילות. וא"כ נמצא שהכריחה על כל פנים.

| id | word | punctuation after |
|---|---|---|
| 5:11.1 | מד | ) |
| 5:11.2 | ואחר |  |
| 5:11.3 | שידעת |  |
| 5:11.4 | את |  |
| 5:11.5 | כל |  |
| 5:11.6 | אלה |  |
| 5:11.7 | בא |  |
| 5:11.8 | וראה |  |
| 5:11.9 | איך |  |
| 5:11.10 | שגגו |  |
| 5:11.11 | רבותינו |  |
| 5:11.12 | המקובלים |  |
| 5:11.13 | והטילו |  |
| 5:11.14 | דופי |  |
| 5:11.15 | בידיעת |  |
| 5:11.16 | השי"ת |  |
| 5:11.17 | בעתידות | , |
| 5:11.18 | גם |  |
| 5:11.19 | סלקו |  |
| 5:11.20 | ההשגחה |  |
| 5:11.21 | מהמלך |  |
| 5:11.22 | העליון |  |
| 5:11.23 | ותלו |  |
| 5:11.24 | ההשגחה |  |
| 5:11.25 | בזעיר |  |
| 5:11.26 | אנפין |  |
| 5:11.27 | לבד | , |
| 5:11.28 | ולא |  |
| 5:11.29 | באשר |  |
| 5:11.30 | הפרצופין |  |
| 5:11.31 | אבא |  |
| 5:11.32 | ואמא | , |
| 5:11.33 | ואריך | , |
| 5:11.34 | ועתיק | , |
| 5:11.35 | ואדם |  |
| 5:11.36 | קדמון | , |
| 5:11.37 | וכ"ש |  |
| 5:11.38 | באין |  |
| 5:11.39 | סוף |  |
| 5:11.40 | העליון | , |
| 5:11.41 | ואמרו |  |
| 5:11.42 | שאין |  |
| 5:11.43 | מקפידין |  |
| 5:11.44 | ומשגיחין |  |
| 5:11.45 | במעשה |  |
| 5:11.46 | התחתונים |  |
| 5:11.47 | בין |  |
| 5:11.48 | טוב |  |
| 5:11.49 | לרע |  |
| 5:11.50 | ח"ו |  |
| 5:11.51 | כטוב |  |
| 5:11.52 | כחוטא | , |
| 5:11.53 | והיה |  |
| 5:11.54 | כצדיק |  |
| 5:11.55 | כרשע | , |
| 5:11.56 | הנשבע |  |
| 5:11.57 | כאשר |  |
| 5:11.58 | שבועה |  |
| 5:11.59 | ירא |  |
| 5:11.60 | כל |  |
| 5:11.61 | אפיא |  |
| 5:11.62 | שוין | . |
| 5:11.63 | ודבר |  |
| 5:11.64 | זה |  |
| 5:11.65 | מבואר |  |
| 5:11.66 | בכתבי |  |
| 5:11.67 | האר"י |  |
| 5:11.68 | הובא |  |
| 5:11.69 | בספר |  |
| 5:11.70 | ד | ' |
| 5:11.71 | מאות |  |
| 5:11.72 | שקל |  |
| 5:11.73 | כסף |  |
| 5:11.74 | דף |  |
| 5:11.75 | צ"א |  |
| 5:11.76 | ע"ב |  |
| 5:11.77 | ז"ל |  |
| 5:11.78 | שאלה |  |
| 5:11.79 | ששאל |  |
| 5:11.80 | הרב |  |
| 5:11.81 | המובהק |  |
| 5:11.82 | המקובל |  |
| 5:11.83 | האלהי |  |
| 5:11.84 | כמהר"ר |  |
| 5:11.85 | אברהם |  |
| 5:11.86 | מונצוץ |  |
| 5:11.87 | מעיר |  |
| 5:11.88 | טניפיל |  |
| 5:11.89 | להאר"י |  |
| 5:11.90 | זצוק"ל |  |
| 5:11.91 | יש |  |
| 5:11.92 | לשאול |  |
| 5:11.93 | דבר |  |
| 5:11.94 | אחד |  |
| 5:11.95 | קשה |  |
| 5:11.96 | והוא |  |
| 5:11.97 | כי |  |
| 5:11.98 | ראינו |  |
| 5:11.99 | בספרי |  |
| 5:11.100 | הקבלה |  |
| 5:11.101 | והזהר |  |
| 5:11.102 | ובספר |  |
| 5:11.103 | המדע |  |
| 5:11.104 | ובספר |  |
| 5:11.105 | הקנה |  |
| 5:11.106 | כי |  |
| 5:11.107 | בשעה |  |
| 5:11.108 | שהאל |  |
| 5:11.109 | מתעסק |  |
| 5:11.110 | בבריאה | , |
| 5:11.111 | אינו |  |
| 5:11.112 | יודע | . |
| 5:11.113 | וכשיהיה |  |
| 5:11.114 | באצילות |  |
| 5:11.115 | יודע | . |
| 5:11.116 | אם |  |
| 5:11.117 | כן |  |
| 5:11.118 | הרי |  |
| 5:11.119 | נתבאר |  |
| 5:11.120 | שבעולם |  |
| 5:11.121 | האצילות |  |
| 5:11.122 | היא |  |
| 5:11.123 | הידיעה |  |
| 5:11.124 | והכריחה | , |
| 5:11.125 | א"כ |  |
| 5:11.126 | קשה |  |
| 5:11.127 | לדברי |  |
| 5:11.128 | רז"ל |  |
| 5:11.129 | על |  |
| 5:11.130 | פסוק |  |
| 5:11.131 | ויאמר |  |
| 5:11.132 | ה | ' |
| 5:11.133 | אל |  |
| 5:11.134 | משה |  |
| 5:11.135 | אמר |  |
| 5:11.136 | אל |  |
| 5:11.137 | הכוהנים |  |
| 5:11.138 | בני |  |
| 5:11.139 | אהרן | , |
| 5:11.140 | מלך |  |
| 5:11.141 | ראשון |  |
| 5:11.142 | שיעמוד |  |
| 5:11.143 | לישראל |  |
| 5:11.144 | ידקר |  |
| 5:11.145 | בחרב | . |
| 5:11.146 | אמר |  |
| 5:11.147 | לו |  |
| 5:11.148 | ולי |  |
| 5:11.149 | אתה |  |
| 5:11.150 | שואל | , |
| 5:11.151 | אמור |  |
| 5:11.152 | אל |  |
| 5:11.153 | הכוהנים |  |
| 5:11.154 | בני |  |
| 5:11.155 | אהרן |  |
| 5:11.156 | שהרי |  |
| 5:11.157 | הרג |  |
| 5:11.158 | נוב |  |
| 5:11.159 | עיר |  |
| 5:11.160 | הכהנים | . |
| 5:11.161 | ונמצא |  |
| 5:11.162 | שהכריחה |  |
| 5:11.163 | דעתו |  |
| 5:11.164 | של |  |
| 5:11.165 | שאול | , |
| 5:11.166 | באיזה |  |
| 5:11.167 | אופן |  |
| 5:11.168 | שיהיה |  |
| 5:11.169 | אם |  |
| 5:11.170 | בבריאה |  |
| 5:11.171 | אם |  |
| 5:11.172 | באצילות | . |
| 5:11.173 | וא"כ |  |
| 5:11.174 | נמצא |  |
| 5:11.175 | שהכריחה |  |
| 5:11.176 | על |  |
| 5:11.177 | כל |  |
| 5:11.178 | פנים | . |


## Milchamot ha-Shem 5:12

Hebrew (Wikisource): והתשובה שהשיב מורי זלה"ה, הוא אמת שיש באצילות ידיעה, אבל יש כח באדם לבחור זולתה כמו שאמר הכתוב ראה נתתי לפניך היום את החיים ואת הטוב ואת המות ואת הרע ובחרת בחיים למען תחיה אתה וזרעיך. הרי מכאן שהכריחה ממעשה שאול, ומן הכתוב מוכיח שהיא בחירה שנ' ובחרת בחיים וגם לרז"ל שאמרו אין רע יורד מן השמים, כי למעלה באצילות הכל פשוט, ואין אותה ידיעה יורד למטה להכריח על האדם כי אין שכר ועונש ואין בחירה ורצון למעלה, וזה סוד הלא אח עשו ליעקב נאם ה' ואהב את יעקב, וכי השי"ת אומר שכמו יעקב כן עשו. כי כבר העידה תורה כי יעקב איש תם יושב אהלים, ועשו עובד ע"ז, והשי"ת בוחר בצדיק, וא"כ מהו זה שאמר השי"ת ואהב את יעקב, זה צדיק וזה רשע, אלא הכוונה היא, כי למעלה באצילות כמו יעקב כן עשו, כי שם לא יש שכר ועונש, וזהו סוד הלא אח עשו ליעקב נאם ה', שהוא עתיקא קדישא, כמו שפירש באדרא, בי נשבעתי נאם ה' עתיקא קדישא. על כן אמר יש לכם לשמוע בקלי ולקבל מלכותי כי בחרתי ביעקב ואף שהוא כמו עשו וגם אם נאמר שיש ידיעה מכרחת, אין צרך לתורה ולא למצות, כי כבר הוכרח האדם במעשיו, וגם מאמר, ר' חנניה בן עקשיא אומר רצה הקב"ה לזכות את ישראל לפיכך הרבה להם תורה ומצות. אם נאמר שיש ידיעה מכרחת, אין צורך בהרבות תורה ומצות. אלא הכוונה היא למעלה באצילות יש ידיעה ואין ידיעת הרע של אדם יורד מלמעלה כי נשאר שם וביד האדם שהוא בעל בחירה לבחור בזולתה. על כן צותה תורה לשמור מצותיה וגם אם האדם דבק בתורה מושך עליו הטוב. כי אין רע יורד מעצמו מלמעלה אלא ימשוך האדם עליו הרע אז יורד. ועם זה יובן מאמר אחר פר' ויקרא או ראה או ידע כי שם אומר כי הנפש קודם שתרד משביע אותה הקב"ה תהי צדיק ואל תהי רשע. כי אם נאמר שיש ידיעה נמצא משביעה לשקר. שהרי הוא יודע בו שיחטא ומשביע. וגם מאמר אחר לרז"ל משביעין אותו תהי צדיק ואל תהי רשע. אם נאמר שידע נמצא משביעו לשקר. ואין הקב"ה רצה שישבעו בשמו לשקר. אלא אלו המאמרים וגם הרבה כתובים מוכיחים שבשעת הבריאה אין לו ידיעה. וממה שביארנו יש ראיה שאפי' ידיעת האצילות היא ידיעה פשוטה, הויא תלויה ברצון האדם שכבר נאמר לו כי הוא כמו קונו שהוא בעל בחירה ורצון כמו שאמר הכתוב כי יעקב בחר לו יה ישראל לסגולתו. הרי הוא מובן שיש לו בחירה ורצון. גם מה שנ' ורציתי אתכם וכו' הרי מן הכתובים שיש לו בחירה ורצון. וגם מדמה ישראל אליו שאמר והייתם קדושים כי קדוש אני כמו שאני קדוש כן אתם תהיו קדושים. אחי ורעי בני כדמותו כצלמו. כי אם יאמר שיש לו ידיעה ואינה מכרחת כמו שכתב הרמב"ם נמצא שאול שהכריחה אותו ידיעתו עליו יש לו תלונה על זה, אמנם אם נאמר כמו שאמרנו יבוא הכל באמת ובאמונה עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:12.1 | והתשובה |  |
| 5:12.2 | שהשיב |  |
| 5:12.3 | מורי |  |
| 5:12.4 | זלה"ה | , |
| 5:12.5 | הוא |  |
| 5:12.6 | אמת |  |
| 5:12.7 | שיש |  |
| 5:12.8 | באצילות |  |
| 5:12.9 | ידיעה | , |
| 5:12.10 | אבל |  |
| 5:12.11 | יש |  |
| 5:12.12 | כח |  |
| 5:12.13 | באדם |  |
| 5:12.14 | לבחור |  |
| 5:12.15 | זולתה |  |
| 5:12.16 | כמו |  |
| 5:12.17 | שאמר |  |
| 5:12.18 | הכתוב |  |
| 5:12.19 | ראה |  |
| 5:12.20 | נתתי |  |
| 5:12.21 | לפניך |  |
| 5:12.22 | היום |  |
| 5:12.23 | את |  |
| 5:12.24 | החיים |  |
| 5:12.25 | ואת |  |
| 5:12.26 | הטוב |  |
| 5:12.27 | ואת |  |
| 5:12.28 | המות |  |
| 5:12.29 | ואת |  |
| 5:12.30 | הרע |  |
| 5:12.31 | ובחרת |  |
| 5:12.32 | בחיים |  |
| 5:12.33 | למען |  |
| 5:12.34 | תחיה |  |
| 5:12.35 | אתה |  |
| 5:12.36 | וזרעיך | . |
| 5:12.37 | הרי |  |
| 5:12.38 | מכאן |  |
| 5:12.39 | שהכריחה |  |
| 5:12.40 | ממעשה |  |
| 5:12.41 | שאול | , |
| 5:12.42 | ומן |  |
| 5:12.43 | הכתוב |  |
| 5:12.44 | מוכיח |  |
| 5:12.45 | שהיא |  |
| 5:12.46 | בחירה |  |
| 5:12.47 | שנ | ' |
| 5:12.48 | ובחרת |  |
| 5:12.49 | בחיים |  |
| 5:12.50 | וגם |  |
| 5:12.51 | לרז"ל |  |
| 5:12.52 | שאמרו |  |
| 5:12.53 | אין |  |
| 5:12.54 | רע |  |
| 5:12.55 | יורד |  |
| 5:12.56 | מן |  |
| 5:12.57 | השמים | , |
| 5:12.58 | כי |  |
| 5:12.59 | למעלה |  |
| 5:12.60 | באצילות |  |
| 5:12.61 | הכל |  |
| 5:12.62 | פשוט | , |
| 5:12.63 | ואין |  |
| 5:12.64 | אותה |  |
| 5:12.65 | ידיעה |  |
| 5:12.66 | יורד |  |
| 5:12.67 | למטה |  |
| 5:12.68 | להכריח |  |
| 5:12.69 | על |  |
| 5:12.70 | האדם |  |
| 5:12.71 | כי |  |
| 5:12.72 | אין |  |
| 5:12.73 | שכר |  |
| 5:12.74 | ועונש |  |
| 5:12.75 | ואין |  |
| 5:12.76 | בחירה |  |
| 5:12.77 | ורצון |  |
| 5:12.78 | למעלה | , |
| 5:12.79 | וזה |  |
| 5:12.80 | סוד |  |
| 5:12.81 | הלא |  |
| 5:12.82 | אח |  |
| 5:12.83 | עשו |  |
| 5:12.84 | ליעקב |  |
| 5:12.85 | נאם |  |
| 5:12.86 | ה | ' |
| 5:12.87 | ואהב |  |
| 5:12.88 | את |  |
| 5:12.89 | יעקב | , |
| 5:12.90 | וכי |  |
| 5:12.91 | השי"ת |  |
| 5:12.92 | אומר |  |
| 5:12.93 | שכמו |  |
| 5:12.94 | יעקב |  |
| 5:12.95 | כן |  |
| 5:12.96 | עשו | . |
| 5:12.97 | כי |  |
| 5:12.98 | כבר |  |
| 5:12.99 | העידה |  |
| 5:12.100 | תורה |  |
| 5:12.101 | כי |  |
| 5:12.102 | יעקב |  |
| 5:12.103 | איש |  |
| 5:12.104 | תם |  |
| 5:12.105 | יושב |  |
| 5:12.106 | אהלים | , |
| 5:12.107 | ועשו |  |
| 5:12.108 | עובד |  |
| 5:12.109 | ע"ז | , |
| 5:12.110 | והשי"ת |  |
| 5:12.111 | בוחר |  |
| 5:12.112 | בצדיק | , |
| 5:12.113 | וא"כ |  |
| 5:12.114 | מהו |  |
| 5:12.115 | זה |  |
| 5:12.116 | שאמר |  |
| 5:12.117 | השי"ת |  |
| 5:12.118 | ואהב |  |
| 5:12.119 | את |  |
| 5:12.120 | יעקב | , |
| 5:12.121 | זה |  |
| 5:12.122 | צדיק |  |
| 5:12.123 | וזה |  |
| 5:12.124 | רשע | , |
| 5:12.125 | אלא |  |
| 5:12.126 | הכוונה |  |
| 5:12.127 | היא | , |
| 5:12.128 | כי |  |
| 5:12.129 | למעלה |  |
| 5:12.130 | באצילות |  |
| 5:12.131 | כמו |  |
| 5:12.132 | יעקב |  |
| 5:12.133 | כן |  |
| 5:12.134 | עשו | , |
| 5:12.135 | כי |  |
| 5:12.136 | שם |  |
| 5:12.137 | לא |  |
| 5:12.138 | יש |  |
| 5:12.139 | שכר |  |
| 5:12.140 | ועונש | , |
| 5:12.141 | וזהו |  |
| 5:12.142 | סוד |  |
| 5:12.143 | הלא |  |
| 5:12.144 | אח |  |
| 5:12.145 | עשו |  |
| 5:12.146 | ליעקב |  |
| 5:12.147 | נאם |  |
| 5:12.148 | ה | ', |
| 5:12.149 | שהוא |  |
| 5:12.150 | עתיקא |  |
| 5:12.151 | קדישא | , |
| 5:12.152 | כמו |  |
| 5:12.153 | שפירש |  |
| 5:12.154 | באדרא | , |
| 5:12.155 | בי |  |
| 5:12.156 | נשבעתי |  |
| 5:12.157 | נאם |  |
| 5:12.158 | ה | ' |
| 5:12.159 | עתיקא |  |
| 5:12.160 | קדישא | . |
| 5:12.161 | על |  |
| 5:12.162 | כן |  |
| 5:12.163 | אמר |  |
| 5:12.164 | יש |  |
| 5:12.165 | לכם |  |
| 5:12.166 | לשמוע |  |
| 5:12.167 | בקלי |  |
| 5:12.168 | ולקבל |  |
| 5:12.169 | מלכותי |  |
| 5:12.170 | כי |  |
| 5:12.171 | בחרתי |  |
| 5:12.172 | ביעקב |  |
| 5:12.173 | ואף |  |
| 5:12.174 | שהוא |  |
| 5:12.175 | כמו |  |
| 5:12.176 | עשו |  |
| 5:12.177 | וגם |  |
| 5:12.178 | אם |  |
| 5:12.179 | נאמר |  |
| 5:12.180 | שיש |  |
| 5:12.181 | ידיעה |  |
| 5:12.182 | מכרחת | , |
| 5:12.183 | אין |  |
| 5:12.184 | צרך |  |
| 5:12.185 | לתורה |  |
| 5:12.186 | ולא |  |
| 5:12.187 | למצות | , |
| 5:12.188 | כי |  |
| 5:12.189 | כבר |  |
| 5:12.190 | הוכרח |  |
| 5:12.191 | האדם |  |
| 5:12.192 | במעשיו | , |
| 5:12.193 | וגם |  |
| 5:12.194 | מאמר | , |
| 5:12.195 | ר | ' |
| 5:12.196 | חנניה |  |
| 5:12.197 | בן |  |
| 5:12.198 | עקשיא |  |
| 5:12.199 | אומר |  |
| 5:12.200 | רצה |  |
| 5:12.201 | הקב"ה |  |
| 5:12.202 | לזכות |  |
| 5:12.203 | את |  |
| 5:12.204 | ישראל |  |
| 5:12.205 | לפיכך |  |
| 5:12.206 | הרבה |  |
| 5:12.207 | להם |  |
| 5:12.208 | תורה |  |
| 5:12.209 | ומצות | . |
| 5:12.210 | אם |  |
| 5:12.211 | נאמר |  |
| 5:12.212 | שיש |  |
| 5:12.213 | ידיעה |  |
| 5:12.214 | מכרחת | , |
| 5:12.215 | אין |  |
| 5:12.216 | צורך |  |
| 5:12.217 | בהרבות |  |
| 5:12.218 | תורה |  |
| 5:12.219 | ומצות | . |
| 5:12.220 | אלא |  |
| 5:12.221 | הכוונה |  |
| 5:12.222 | היא |  |
| 5:12.223 | למעלה |  |
| 5:12.224 | באצילות |  |
| 5:12.225 | יש |  |
| 5:12.226 | ידיעה |  |
| 5:12.227 | ואין |  |
| 5:12.228 | ידיעת |  |
| 5:12.229 | הרע |  |
| 5:12.230 | של |  |
| 5:12.231 | אדם |  |
| 5:12.232 | יורד |  |
| 5:12.233 | מלמעלה |  |
| 5:12.234 | כי |  |
| 5:12.235 | נשאר |  |
| 5:12.236 | שם |  |
| 5:12.237 | וביד |  |
| 5:12.238 | האדם |  |
| 5:12.239 | שהוא |  |
| 5:12.240 | בעל |  |
| 5:12.241 | בחירה |  |
| 5:12.242 | לבחור |  |
| 5:12.243 | בזולתה | . |
| 5:12.244 | על |  |
| 5:12.245 | כן |  |
| 5:12.246 | צותה |  |
| 5:12.247 | תורה |  |
| 5:12.248 | לשמור |  |
| 5:12.249 | מצותיה |  |
| 5:12.250 | וגם |  |
| 5:12.251 | אם |  |
| 5:12.252 | האדם |  |
| 5:12.253 | דבק |  |
| 5:12.254 | בתורה |  |
| 5:12.255 | מושך |  |
| 5:12.256 | עליו |  |
| 5:12.257 | הטוב | . |
| 5:12.258 | כי |  |
| 5:12.259 | אין |  |
| 5:12.260 | רע |  |
| 5:12.261 | יורד |  |
| 5:12.262 | מעצמו |  |
| 5:12.263 | מלמעלה |  |
| 5:12.264 | אלא |  |
| 5:12.265 | ימשוך |  |
| 5:12.266 | האדם |  |
| 5:12.267 | עליו |  |
| 5:12.268 | הרע |  |
| 5:12.269 | אז |  |
| 5:12.270 | יורד | . |
| 5:12.271 | ועם |  |
| 5:12.272 | זה |  |
| 5:12.273 | יובן |  |
| 5:12.274 | מאמר |  |
| 5:12.275 | אחר |  |
| 5:12.276 | פר | ' |
| 5:12.277 | ויקרא |  |
| 5:12.278 | או |  |
| 5:12.279 | ראה |  |
| 5:12.280 | או |  |
| 5:12.281 | ידע |  |
| 5:12.282 | כי |  |
| 5:12.283 | שם |  |
| 5:12.284 | אומר |  |
| 5:12.285 | כי |  |
| 5:12.286 | הנפש |  |
| 5:12.287 | קודם |  |
| 5:12.288 | שתרד |  |
| 5:12.289 | משביע |  |
| 5:12.290 | אותה |  |
| 5:12.291 | הקב"ה |  |
| 5:12.292 | תהי |  |
| 5:12.293 | צדיק |  |
| 5:12.294 | ואל |  |
| 5:12.295 | תהי |  |
| 5:12.296 | רשע | . |
| 5:12.297 | כי |  |
| 5:12.298 | אם |  |
| 5:12.299 | נאמר |  |
| 5:12.300 | שיש |  |
| 5:12.301 | ידיעה |  |
| 5:12.302 | נמצא |  |
| 5:12.303 | משביעה |  |
| 5:12.304 | לשקר | . |
| 5:12.305 | שהרי |  |
| 5:12.306 | הוא |  |
| 5:12.307 | יודע |  |
| 5:12.308 | בו |  |
| 5:12.309 | שיחטא |  |
| 5:12.310 | ומשביע | . |
| 5:12.311 | וגם |  |
| 5:12.312 | מאמר |  |
| 5:12.313 | אחר |  |
| 5:12.314 | לרז"ל |  |
| 5:12.315 | משביעין |  |
| 5:12.316 | אותו |  |
| 5:12.317 | תהי |  |
| 5:12.318 | צדיק |  |
| 5:12.319 | ואל |  |
| 5:12.320 | תהי |  |
| 5:12.321 | רשע | . |
| 5:12.322 | אם |  |
| 5:12.323 | נאמר |  |
| 5:12.324 | שידע |  |
| 5:12.325 | נמצא |  |
| 5:12.326 | משביעו |  |
| 5:12.327 | לשקר | . |
| 5:12.328 | ואין |  |
| 5:12.329 | הקב"ה |  |
| 5:12.330 | רצה |  |
| 5:12.331 | שישבעו |  |
| 5:12.332 | בשמו |  |
| 5:12.333 | לשקר | . |
| 5:12.334 | אלא |  |
| 5:12.335 | אלו |  |
| 5:12.336 | המאמרים |  |
| 5:12.337 | וגם |  |
| 5:12.338 | הרבה |  |
| 5:12.339 | כתובים |  |
| 5:12.340 | מוכיחים |  |
| 5:12.341 | שבשעת |  |
| 5:12.342 | הבריאה |  |
| 5:12.343 | אין |  |
| 5:12.344 | לו |  |
| 5:12.345 | ידיעה | . |
| 5:12.346 | וממה |  |
| 5:12.347 | שביארנו |  |
| 5:12.348 | יש |  |
| 5:12.349 | ראיה |  |
| 5:12.350 | שאפי | ' |
| 5:12.351 | ידיעת |  |
| 5:12.352 | האצילות |  |
| 5:12.353 | היא |  |
| 5:12.354 | ידיעה |  |
| 5:12.355 | פשוטה | , |
| 5:12.356 | הויא |  |
| 5:12.357 | תלויה |  |
| 5:12.358 | ברצון |  |
| 5:12.359 | האדם |  |
| 5:12.360 | שכבר |  |
| 5:12.361 | נאמר |  |
| 5:12.362 | לו |  |
| 5:12.363 | כי |  |
| 5:12.364 | הוא |  |
| 5:12.365 | כמו |  |
| 5:12.366 | קונו |  |
| 5:12.367 | שהוא |  |
| 5:12.368 | בעל |  |
| 5:12.369 | בחירה |  |
| 5:12.370 | ורצון |  |
| 5:12.371 | כמו |  |
| 5:12.372 | שאמר |  |
| 5:12.373 | הכתוב |  |
| 5:12.374 | כי |  |
| 5:12.375 | יעקב |  |
| 5:12.376 | בחר |  |
| 5:12.377 | לו |  |
| 5:12.378 | יה |  |
| 5:12.379 | ישראל |  |
| 5:12.380 | לסגולתו | . |
| 5:12.381 | הרי |  |
| 5:12.382 | הוא |  |
| 5:12.383 | מובן |  |
| 5:12.384 | שיש |  |
| 5:12.385 | לו |  |
| 5:12.386 | בחירה |  |
| 5:12.387 | ורצון | . |
| 5:12.388 | גם |  |
| 5:12.389 | מה |  |
| 5:12.390 | שנ | ' |
| 5:12.391 | ורציתי |  |
| 5:12.392 | אתכם |  |
| 5:12.393 | וכו | ' |
| 5:12.394 | הרי |  |
| 5:12.395 | מן |  |
| 5:12.396 | הכתובים |  |
| 5:12.397 | שיש |  |
| 5:12.398 | לו |  |
| 5:12.399 | בחירה |  |
| 5:12.400 | ורצון | . |
| 5:12.401 | וגם |  |
| 5:12.402 | מדמה |  |
| 5:12.403 | ישראל |  |
| 5:12.404 | אליו |  |
| 5:12.405 | שאמר |  |
| 5:12.406 | והייתם |  |
| 5:12.407 | קדושים |  |
| 5:12.408 | כי |  |
| 5:12.409 | קדוש |  |
| 5:12.410 | אני |  |
| 5:12.411 | כמו |  |
| 5:12.412 | שאני |  |
| 5:12.413 | קדוש |  |
| 5:12.414 | כן |  |
| 5:12.415 | אתם |  |
| 5:12.416 | תהיו |  |
| 5:12.417 | קדושים | . |
| 5:12.418 | אחי |  |
| 5:12.419 | ורעי |  |
| 5:12.420 | בני |  |
| 5:12.421 | כדמותו |  |
| 5:12.422 | כצלמו | . |
| 5:12.423 | כי |  |
| 5:12.424 | אם |  |
| 5:12.425 | יאמר |  |
| 5:12.426 | שיש |  |
| 5:12.427 | לו |  |
| 5:12.428 | ידיעה |  |
| 5:12.429 | ואינה |  |
| 5:12.430 | מכרחת |  |
| 5:12.431 | כמו |  |
| 5:12.432 | שכתב |  |
| 5:12.433 | הרמב"ם |  |
| 5:12.434 | נמצא |  |
| 5:12.435 | שאול |  |
| 5:12.436 | שהכריחה |  |
| 5:12.437 | אותו |  |
| 5:12.438 | ידיעתו |  |
| 5:12.439 | עליו |  |
| 5:12.440 | יש |  |
| 5:12.441 | לו |  |
| 5:12.442 | תלונה |  |
| 5:12.443 | על |  |
| 5:12.444 | זה | , |
| 5:12.445 | אמנם |  |
| 5:12.446 | אם |  |
| 5:12.447 | נאמר |  |
| 5:12.448 | כמו |  |
| 5:12.449 | שאמרנו |  |
| 5:12.450 | יבוא |  |
| 5:12.451 | הכל |  |
| 5:12.452 | באמת |  |
| 5:12.453 | ובאמונה |  |
| 5:12.454 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

