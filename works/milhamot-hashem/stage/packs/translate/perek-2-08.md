# Translation pack: milhamot-hashem perek-2-08 (Milchamot ha-Shem 2:16–2:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-08`.

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
- ה׳ (the Tetragrammaton as written): "the LORD", in the author's prose and
  in verses. Where it is cited as a name ("זעיר אנפין דאיקרי ה'", 4:11), write
  "who is called 'the LORD'". Not ה׳ as Tetragrammaton: "מאת ה' אליעזר בן
  יהודה" (1:2) → "from Mr. Eliezer Ben-Yehuda"; ה׳ as a numeral (5) or a page
  letter → the number. The pack splits it into the token ה plus the
  punctuation ', so read the token in context.
- הוי"ה / הוי'ה / הויה (the Name as a name: Kabbalistic usage, and the
  author's own in 3:17) → "YHVH". Never "being". Examples: "ד' מילוי הויה ע"ב
  ס"ג מ"ה ב"ן" (2:33) → "the four fillings of YHVH, 72, 63, 45, 52" (in the
  note: these are the gematria names AB, SaG, MaH, BaN). "שהוא קורא לכל אחד
  מהם בשם הוי'ה ואדנות ואלהים" (8:24) → "who calls each of them by the names
  YHVH, Adonai and Elohim". אדנות as a name → "Adonai", with a note that the
  word is literally "lordship". עשר הויות (2:33) → "ten havayot", with a note.
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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
- ר"א: "Rabbi Eleazar" only where the paragraph names him, else "R. A." with
  a note.
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
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- ס' before a quoted verse is "verse" (סימן = "section" in the author's
  numbered sections); read it from the context and note the choice.
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
- אנדרטא: "statue"; אשם תלוי: "suspended guilt offering"; ביהכנ"ס: "synagogue".
  Each with a token note if the term is technical.
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- משותף in the Maimonidean sense (a term said of God and of creatures by
  equivocation): "equivocal"; "shared" only for a thing held in common.
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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
- ב"ע: "in Aramaic" (בארמית) where it is set against בערבי; the other readings in
  `readings`.
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
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


## Milchamot ha-Shem 2:16

Hebrew (Wikisource): ומה שכתבת: ויותר מזה מה שכ' על כבוד ה' שהוא אור נברא אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחוים לנברא? לא לכבוד הנראה משתחוים, רק למי שהשרה כבודו שם. והכבוד הנראה שם לאות וסימן להם, כי השי"ת איווה לאוהל מועד למושב לו, כמו שאנחנו משתחוים ומתפללים לשי"ת נכח המקדש וירושלם, כי בחר ה' בציון אוה למושב לו. וכבר ביאר הרמב"ם בפ' ס"ד מראשון, כי כבוד ה', פעמים רוצים בו האור הנברא אשר ישכינהו הק' במקום, כמו: וישכון כבוד ה' על הר סיני ויכסהו וכו', ופעמים רוצים בו עצמותו ית' ואמתתו, כמו "הראיני נא את כבודך", ובא המענה: "כי לא יראני האדם וחי", מורה כי הכבוד הנאמר פה עצמו ית', ואמר "כבודך" להגדיל וכו'. וכיוצא בזה מצינו בדברי רז"ל ששם אלהים משותף לדיינים ולמלאכים ולשי"ת, גם שם שכינה משותף למה שהזכרתי למעלה. וראיתי להעירך פה שאין איסור להשתחוות לאדם כל שאין דעתו לקבלו באלוה וגם אינו נעבד, רק נכנע לו ומשפיל עצמו לפניו, או דרך ההודאה על טובה שעשה לו. שהרי אברהם אבינו השתחוה לבני חת על החסד שעשו עמו ליתן לו אחזת קבר. וכן תרגם רבי סעדיה גאון: וישתחו לעם הארץ לבני חת, וסגד לאהל אלבלד שברא. ויהושוע השתחווה למלאך. ונתן הנביא ובת שבע השתחוו לדוד. והרבה מצינו בכתובים שהשתחוו לנברא. וגרסינן בסנהדרין פ' ארבע מיתות (סנהדרין ס"א ע"ב), אתמר: העובד ע"ז מאהבה ומיראה, אביי אמר חייב, דהא פלחה. רבא אמר פטור, אי קבליה עליה באלוה אין, אי לא לא. ואמר אביי מנא אמינא לה, דתנן העובד ע"ז אחד העובד ואחד המזבח ואחד המקטיר ואחד המנסך ואחד המשתחוה ואחד המקבלו עליו באלוה והאומר לו אלי אתה. מאי לאו אחד העובד מאהבה או מיראה? ורבא אמר לך, לא, אלא כדמתרץ ר' ירמיה (פירש רש"י: אחד העובד כל עבודה שהיא דרכה בכך, ואחד המזבח וכו' שלא כדרכה, ולעולם רישא וסיפא לשם אלהות). אמר אביי: מנא אמינא לה? דתניא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך. יכול אפי' לעבד כהמן? ת"ל ולא תעבדם. והא המן מיראה הוא נעבד! ורבא? כהמן ולא כהמן. כהמן, דאיהו גופיה ע"ז; ולא כהמן, דאלו המן מיראה, והכא לאו מיראה. ואמר אביי: מנא אמינא לה? דתניא: כהן משיח בע"ז. ר' אומר בשגגת מעשה, וחכמים אומרים בהעלם דבר, ושווין דבשעירה כיחיד, ושווין שאין מביא אשם תלוי. האי שגגת מעשה היכי דמי? אי קסבר ביהכנ"ס והשתחוה לו, הרי לבו לשמים. אלא דחזא אנדרטא והשתחוה לו (פירוש, שראה דמות המלך שהיו רגילים לעשות כדמות המלך והרואה אותו משתחוה לו לכבוד המלך, ופעמים שעובדין אותו, וזה ראה אחד שנעבד והשתחוה לו), אי קבליה עליה באלוה מזיד הוא, ואי לא קבליה עליה באלוה לאו כלום הוא (פירש רש"י: אלא לכבוד המלך השתחוה לו, לאו כלום הוא דהא לא איכוון לע"ז). אלא לאו מאהבה ומיראה? ורבא אמר לך: באומר מותר. אומר מותר היינו העלם דבר? באומר מותר לגמרי, העלם דבר קיום מקצת וביטול מקצת, עכ"ל הגמ' עם קצת פי' רש"י. ודבר נלמד, דלכל נברא אסור להשתחוות, זולת האדם, והוא שלא נעבד כהמן שעשה עצמו ע"ז, כדאמירן במגלה (יט), דאי לאו הכי לא היה מרדכי נמנע מלהשתחות לו, כדקתני בבריתא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך, והוא שלא יהא נעבד כהמן שעשה עצמו אלוה, רק משתחוה לו לאות הכבוד ולא לאלהות, ולכך השתחוה יהושע למלאך הואיל ונגלה לו בדמות אדם, כמ"ש: והנה איש עומד לנגדו, ולכן השתחוה לו לאות הכבוד ולא לעבדו. ועיין בס' העיקרים מאמר ב' פ' כ"ח שהאריך בביאור הכתובים שקורין לכבוד הנראה לנביאים בשם הוי"ה, וגם לארון, כמו שגאמר: "ויהי בנסע הארון וכו' קומה ה', ובנחה יאמר שובה ה'", ויהושע קרא לארון "אדון כל הארץ" ע"ש. ולדעתו שמה שהשתחוה יהושע למלאך, לפי שההשתחויה ההיא היתה לכבוד השי"ת, ולהיות זה שלוחו וכו' ע"ש. ורז"ל אמרו תעלא בעידניה סגיד ליה, ר"ל: אדם פחות ובזוי כשועל, והשעה עומדת לו באיזה משרה שבידו, סגיד ליה. וכן הוא נהוג אצלנו ואצל הערביים להשתחוות ולנשק אדם גדול המעלה, ואין בזה שום איסור, שאינו מקבלו עליו באלוה, רק לאות כבוד הוא משתחוה לו ומנשק את רגליו.

| id | word | punctuation after |
|---|---|---|
| 2:16.1 | ומה |  |
| 2:16.2 | שכתבת | : |
| 2:16.3 | ויותר |  |
| 2:16.4 | מזה |  |
| 2:16.5 | מה |  |
| 2:16.6 | שכ | ' |
| 2:16.7 | על |  |
| 2:16.8 | כבוד |  |
| 2:16.9 | ה | ' |
| 2:16.10 | שהוא |  |
| 2:16.11 | אור |  |
| 2:16.12 | נברא |  |
| 2:16.13 | אם |  |
| 2:16.14 | כן |  |
| 2:16.15 | כיצד |  |
| 2:16.16 | היו |  |
| 2:16.17 | ישראל |  |
| 2:16.18 | יוצאים |  |
| 2:16.19 | פתח |  |
| 2:16.20 | אהליהם |  |
| 2:16.21 | ומשתחוים |  |
| 2:16.22 | לנברא | ? |
| 2:16.23 | לא |  |
| 2:16.24 | לכבוד |  |
| 2:16.25 | הנראה |  |
| 2:16.26 | משתחוים | , |
| 2:16.27 | רק |  |
| 2:16.28 | למי |  |
| 2:16.29 | שהשרה |  |
| 2:16.30 | כבודו |  |
| 2:16.31 | שם | . |
| 2:16.32 | והכבוד |  |
| 2:16.33 | הנראה |  |
| 2:16.34 | שם |  |
| 2:16.35 | לאות |  |
| 2:16.36 | וסימן |  |
| 2:16.37 | להם | , |
| 2:16.38 | כי |  |
| 2:16.39 | השי"ת |  |
| 2:16.40 | איווה |  |
| 2:16.41 | לאוהל |  |
| 2:16.42 | מועד |  |
| 2:16.43 | למושב |  |
| 2:16.44 | לו | , |
| 2:16.45 | כמו |  |
| 2:16.46 | שאנחנו |  |
| 2:16.47 | משתחוים |  |
| 2:16.48 | ומתפללים |  |
| 2:16.49 | לשי"ת |  |
| 2:16.50 | נכח |  |
| 2:16.51 | המקדש |  |
| 2:16.52 | וירושלם | , |
| 2:16.53 | כי |  |
| 2:16.54 | בחר |  |
| 2:16.55 | ה | ' |
| 2:16.56 | בציון |  |
| 2:16.57 | אוה |  |
| 2:16.58 | למושב |  |
| 2:16.59 | לו | . |
| 2:16.60 | וכבר |  |
| 2:16.61 | ביאר |  |
| 2:16.62 | הרמב"ם |  |
| 2:16.63 | בפ | ' |
| 2:16.64 | ס"ד |  |
| 2:16.65 | מראשון | , |
| 2:16.66 | כי |  |
| 2:16.67 | כבוד |  |
| 2:16.68 | ה | ', |
| 2:16.69 | פעמים |  |
| 2:16.70 | רוצים |  |
| 2:16.71 | בו |  |
| 2:16.72 | האור |  |
| 2:16.73 | הנברא |  |
| 2:16.74 | אשר |  |
| 2:16.75 | ישכינהו |  |
| 2:16.76 | הק | ' |
| 2:16.77 | במקום | , |
| 2:16.78 | כמו | : |
| 2:16.79 | וישכון |  |
| 2:16.80 | כבוד |  |
| 2:16.81 | ה | ' |
| 2:16.82 | על |  |
| 2:16.83 | הר |  |
| 2:16.84 | סיני |  |
| 2:16.85 | ויכסהו |  |
| 2:16.86 | וכו | ', |
| 2:16.87 | ופעמים |  |
| 2:16.88 | רוצים |  |
| 2:16.89 | בו |  |
| 2:16.90 | עצמותו |  |
| 2:16.91 | ית | ' |
| 2:16.92 | ואמתתו | , |
| 2:16.93 | כמו |  |
| 2:16.94 | הראיני |  |
| 2:16.95 | נא |  |
| 2:16.96 | את |  |
| 2:16.97 | כבודך | ", |
| 2:16.98 | ובא |  |
| 2:16.99 | המענה | : |
| 2:16.100 | כי |  |
| 2:16.101 | לא |  |
| 2:16.102 | יראני |  |
| 2:16.103 | האדם |  |
| 2:16.104 | וחי | ", |
| 2:16.105 | מורה |  |
| 2:16.106 | כי |  |
| 2:16.107 | הכבוד |  |
| 2:16.108 | הנאמר |  |
| 2:16.109 | פה |  |
| 2:16.110 | עצמו |  |
| 2:16.111 | ית | ', |
| 2:16.112 | ואמר |  |
| 2:16.113 | כבודך | " |
| 2:16.114 | להגדיל |  |
| 2:16.115 | וכו | '. |
| 2:16.116 | וכיוצא |  |
| 2:16.117 | בזה |  |
| 2:16.118 | מצינו |  |
| 2:16.119 | בדברי |  |
| 2:16.120 | רז"ל |  |
| 2:16.121 | ששם |  |
| 2:16.122 | אלהים |  |
| 2:16.123 | משותף |  |
| 2:16.124 | לדיינים |  |
| 2:16.125 | ולמלאכים |  |
| 2:16.126 | ולשי"ת | , |
| 2:16.127 | גם |  |
| 2:16.128 | שם |  |
| 2:16.129 | שכינה |  |
| 2:16.130 | משותף |  |
| 2:16.131 | למה |  |
| 2:16.132 | שהזכרתי |  |
| 2:16.133 | למעלה | . |
| 2:16.134 | וראיתי |  |
| 2:16.135 | להעירך |  |
| 2:16.136 | פה |  |
| 2:16.137 | שאין |  |
| 2:16.138 | איסור |  |
| 2:16.139 | להשתחוות |  |
| 2:16.140 | לאדם |  |
| 2:16.141 | כל |  |
| 2:16.142 | שאין |  |
| 2:16.143 | דעתו |  |
| 2:16.144 | לקבלו |  |
| 2:16.145 | באלוה |  |
| 2:16.146 | וגם |  |
| 2:16.147 | אינו |  |
| 2:16.148 | נעבד | , |
| 2:16.149 | רק |  |
| 2:16.150 | נכנע |  |
| 2:16.151 | לו |  |
| 2:16.152 | ומשפיל |  |
| 2:16.153 | עצמו |  |
| 2:16.154 | לפניו | , |
| 2:16.155 | או |  |
| 2:16.156 | דרך |  |
| 2:16.157 | ההודאה |  |
| 2:16.158 | על |  |
| 2:16.159 | טובה |  |
| 2:16.160 | שעשה |  |
| 2:16.161 | לו | . |
| 2:16.162 | שהרי |  |
| 2:16.163 | אברהם |  |
| 2:16.164 | אבינו |  |
| 2:16.165 | השתחוה |  |
| 2:16.166 | לבני |  |
| 2:16.167 | חת |  |
| 2:16.168 | על |  |
| 2:16.169 | החסד |  |
| 2:16.170 | שעשו |  |
| 2:16.171 | עמו |  |
| 2:16.172 | ליתן |  |
| 2:16.173 | לו |  |
| 2:16.174 | אחזת |  |
| 2:16.175 | קבר | . |
| 2:16.176 | וכן |  |
| 2:16.177 | תרגם |  |
| 2:16.178 | רבי |  |
| 2:16.179 | סעדיה |  |
| 2:16.180 | גאון | : |
| 2:16.181 | וישתחו |  |
| 2:16.182 | לעם |  |
| 2:16.183 | הארץ |  |
| 2:16.184 | לבני |  |
| 2:16.185 | חת | , |
| 2:16.186 | וסגד |  |
| 2:16.187 | לאהל |  |
| 2:16.188 | אלבלד |  |
| 2:16.189 | שברא | . |
| 2:16.190 | ויהושוע |  |
| 2:16.191 | השתחווה |  |
| 2:16.192 | למלאך | . |
| 2:16.193 | ונתן |  |
| 2:16.194 | הנביא |  |
| 2:16.195 | ובת |  |
| 2:16.196 | שבע |  |
| 2:16.197 | השתחוו |  |
| 2:16.198 | לדוד | . |
| 2:16.199 | והרבה |  |
| 2:16.200 | מצינו |  |
| 2:16.201 | בכתובים |  |
| 2:16.202 | שהשתחוו |  |
| 2:16.203 | לנברא | . |
| 2:16.204 | וגרסינן |  |
| 2:16.205 | בסנהדרין |  |
| 2:16.206 | פ | ' |
| 2:16.207 | ארבע |  |
| 2:16.208 | מיתות |  |
| 2:16.209 | סנהדרין |  |
| 2:16.210 | ס"א |  |
| 2:16.211 | ע"ב | ), |
| 2:16.212 | אתמר | : |
| 2:16.213 | העובד |  |
| 2:16.214 | ע"ז |  |
| 2:16.215 | מאהבה |  |
| 2:16.216 | ומיראה | , |
| 2:16.217 | אביי |  |
| 2:16.218 | אמר |  |
| 2:16.219 | חייב | , |
| 2:16.220 | דהא |  |
| 2:16.221 | פלחה | . |
| 2:16.222 | רבא |  |
| 2:16.223 | אמר |  |
| 2:16.224 | פטור | , |
| 2:16.225 | אי |  |
| 2:16.226 | קבליה |  |
| 2:16.227 | עליה |  |
| 2:16.228 | באלוה |  |
| 2:16.229 | אין | , |
| 2:16.230 | אי |  |
| 2:16.231 | לא |  |
| 2:16.232 | לא | . |
| 2:16.233 | ואמר |  |
| 2:16.234 | אביי |  |
| 2:16.235 | מנא |  |
| 2:16.236 | אמינא |  |
| 2:16.237 | לה | , |
| 2:16.238 | דתנן |  |
| 2:16.239 | העובד |  |
| 2:16.240 | ע"ז |  |
| 2:16.241 | אחד |  |
| 2:16.242 | העובד |  |
| 2:16.243 | ואחד |  |
| 2:16.244 | המזבח |  |
| 2:16.245 | ואחד |  |
| 2:16.246 | המקטיר |  |
| 2:16.247 | ואחד |  |
| 2:16.248 | המנסך |  |
| 2:16.249 | ואחד |  |
| 2:16.250 | המשתחוה |  |
| 2:16.251 | ואחד |  |
| 2:16.252 | המקבלו |  |
| 2:16.253 | עליו |  |
| 2:16.254 | באלוה |  |
| 2:16.255 | והאומר |  |
| 2:16.256 | לו |  |
| 2:16.257 | אלי |  |
| 2:16.258 | אתה | . |
| 2:16.259 | מאי |  |
| 2:16.260 | לאו |  |
| 2:16.261 | אחד |  |
| 2:16.262 | העובד |  |
| 2:16.263 | מאהבה |  |
| 2:16.264 | או |  |
| 2:16.265 | מיראה | ? |
| 2:16.266 | ורבא |  |
| 2:16.267 | אמר |  |
| 2:16.268 | לך | , |
| 2:16.269 | לא | , |
| 2:16.270 | אלא |  |
| 2:16.271 | כדמתרץ |  |
| 2:16.272 | ר | ' |
| 2:16.273 | ירמיה |  |
| 2:16.274 | פירש |  |
| 2:16.275 | רש"י | : |
| 2:16.276 | אחד |  |
| 2:16.277 | העובד |  |
| 2:16.278 | כל |  |
| 2:16.279 | עבודה |  |
| 2:16.280 | שהיא |  |
| 2:16.281 | דרכה |  |
| 2:16.282 | בכך | , |
| 2:16.283 | ואחד |  |
| 2:16.284 | המזבח |  |
| 2:16.285 | וכו | ' |
| 2:16.286 | שלא |  |
| 2:16.287 | כדרכה | , |
| 2:16.288 | ולעולם |  |
| 2:16.289 | רישא |  |
| 2:16.290 | וסיפא |  |
| 2:16.291 | לשם |  |
| 2:16.292 | אלהות | ). |
| 2:16.293 | אמר |  |
| 2:16.294 | אביי | : |
| 2:16.295 | מנא |  |
| 2:16.296 | אמינא |  |
| 2:16.297 | לה | ? |
| 2:16.298 | דתניא | : |
| 2:16.299 | לא |  |
| 2:16.300 | תשתחוה |  |
| 2:16.301 | להם | , |
| 2:16.302 | להם |  |
| 2:16.303 | אי |  |
| 2:16.304 | אתה |  |
| 2:16.305 | משתחוה | , |
| 2:16.306 | אבל |  |
| 2:16.307 | אתה |  |
| 2:16.308 | משתחוה |  |
| 2:16.309 | לאדם |  |
| 2:16.310 | כמותך | . |
| 2:16.311 | יכול |  |
| 2:16.312 | אפי | ' |
| 2:16.313 | לעבד |  |
| 2:16.314 | כהמן | ? |
| 2:16.315 | ת"ל |  |
| 2:16.316 | ולא |  |
| 2:16.317 | תעבדם | . |
| 2:16.318 | והא |  |
| 2:16.319 | המן |  |
| 2:16.320 | מיראה |  |
| 2:16.321 | הוא |  |
| 2:16.322 | נעבד | ! |
| 2:16.323 | ורבא | ? |
| 2:16.324 | כהמן |  |
| 2:16.325 | ולא |  |
| 2:16.326 | כהמן | . |
| 2:16.327 | כהמן | , |
| 2:16.328 | דאיהו |  |
| 2:16.329 | גופיה |  |
| 2:16.330 | ע"ז | ; |
| 2:16.331 | ולא |  |
| 2:16.332 | כהמן | , |
| 2:16.333 | דאלו |  |
| 2:16.334 | המן |  |
| 2:16.335 | מיראה | , |
| 2:16.336 | והכא |  |
| 2:16.337 | לאו |  |
| 2:16.338 | מיראה | . |
| 2:16.339 | ואמר |  |
| 2:16.340 | אביי | : |
| 2:16.341 | מנא |  |
| 2:16.342 | אמינא |  |
| 2:16.343 | לה | ? |
| 2:16.344 | דתניא | : |
| 2:16.345 | כהן |  |
| 2:16.346 | משיח |  |
| 2:16.347 | בע"ז | . |
| 2:16.348 | ר | ' |
| 2:16.349 | אומר |  |
| 2:16.350 | בשגגת |  |
| 2:16.351 | מעשה | , |
| 2:16.352 | וחכמים |  |
| 2:16.353 | אומרים |  |
| 2:16.354 | בהעלם |  |
| 2:16.355 | דבר | , |
| 2:16.356 | ושווין |  |
| 2:16.357 | דבשעירה |  |
| 2:16.358 | כיחיד | , |
| 2:16.359 | ושווין |  |
| 2:16.360 | שאין |  |
| 2:16.361 | מביא |  |
| 2:16.362 | אשם |  |
| 2:16.363 | תלוי | . |
| 2:16.364 | האי |  |
| 2:16.365 | שגגת |  |
| 2:16.366 | מעשה |  |
| 2:16.367 | היכי |  |
| 2:16.368 | דמי | ? |
| 2:16.369 | אי |  |
| 2:16.370 | קסבר |  |
| 2:16.371 | ביהכנ"ס |  |
| 2:16.372 | והשתחוה |  |
| 2:16.373 | לו | , |
| 2:16.374 | הרי |  |
| 2:16.375 | לבו |  |
| 2:16.376 | לשמים | . |
| 2:16.377 | אלא |  |
| 2:16.378 | דחזא |  |
| 2:16.379 | אנדרטא |  |
| 2:16.380 | והשתחוה |  |
| 2:16.381 | לו |  |
| 2:16.382 | פירוש | , |
| 2:16.383 | שראה |  |
| 2:16.384 | דמות |  |
| 2:16.385 | המלך |  |
| 2:16.386 | שהיו |  |
| 2:16.387 | רגילים |  |
| 2:16.388 | לעשות |  |
| 2:16.389 | כדמות |  |
| 2:16.390 | המלך |  |
| 2:16.391 | והרואה |  |
| 2:16.392 | אותו |  |
| 2:16.393 | משתחוה |  |
| 2:16.394 | לו |  |
| 2:16.395 | לכבוד |  |
| 2:16.396 | המלך | , |
| 2:16.397 | ופעמים |  |
| 2:16.398 | שעובדין |  |
| 2:16.399 | אותו | , |
| 2:16.400 | וזה |  |
| 2:16.401 | ראה |  |
| 2:16.402 | אחד |  |
| 2:16.403 | שנעבד |  |
| 2:16.404 | והשתחוה |  |
| 2:16.405 | לו | ), |
| 2:16.406 | אי |  |
| 2:16.407 | קבליה |  |
| 2:16.408 | עליה |  |
| 2:16.409 | באלוה |  |
| 2:16.410 | מזיד |  |
| 2:16.411 | הוא | , |
| 2:16.412 | ואי |  |
| 2:16.413 | לא |  |
| 2:16.414 | קבליה |  |
| 2:16.415 | עליה |  |
| 2:16.416 | באלוה |  |
| 2:16.417 | לאו |  |
| 2:16.418 | כלום |  |
| 2:16.419 | הוא |  |
| 2:16.420 | פירש |  |
| 2:16.421 | רש"י | : |
| 2:16.422 | אלא |  |
| 2:16.423 | לכבוד |  |
| 2:16.424 | המלך |  |
| 2:16.425 | השתחוה |  |
| 2:16.426 | לו | , |
| 2:16.427 | לאו |  |
| 2:16.428 | כלום |  |
| 2:16.429 | הוא |  |
| 2:16.430 | דהא |  |
| 2:16.431 | לא |  |
| 2:16.432 | איכוון |  |
| 2:16.433 | לע"ז | ). |
| 2:16.434 | אלא |  |
| 2:16.435 | לאו |  |
| 2:16.436 | מאהבה |  |
| 2:16.437 | ומיראה | ? |
| 2:16.438 | ורבא |  |
| 2:16.439 | אמר |  |
| 2:16.440 | לך | : |
| 2:16.441 | באומר |  |
| 2:16.442 | מותר | . |
| 2:16.443 | אומר |  |
| 2:16.444 | מותר |  |
| 2:16.445 | היינו |  |
| 2:16.446 | העלם |  |
| 2:16.447 | דבר | ? |
| 2:16.448 | באומר |  |
| 2:16.449 | מותר |  |
| 2:16.450 | לגמרי | , |
| 2:16.451 | העלם |  |
| 2:16.452 | דבר |  |
| 2:16.453 | קיום |  |
| 2:16.454 | מקצת |  |
| 2:16.455 | וביטול |  |
| 2:16.456 | מקצת | , |
| 2:16.457 | עכ"ל |  |
| 2:16.458 | הגמ | ' |
| 2:16.459 | עם |  |
| 2:16.460 | קצת |  |
| 2:16.461 | פי | ' |
| 2:16.462 | רש"י | . |
| 2:16.463 | ודבר |  |
| 2:16.464 | נלמד | , |
| 2:16.465 | דלכל |  |
| 2:16.466 | נברא |  |
| 2:16.467 | אסור |  |
| 2:16.468 | להשתחוות | , |
| 2:16.469 | זולת |  |
| 2:16.470 | האדם | , |
| 2:16.471 | והוא |  |
| 2:16.472 | שלא |  |
| 2:16.473 | נעבד |  |
| 2:16.474 | כהמן |  |
| 2:16.475 | שעשה |  |
| 2:16.476 | עצמו |  |
| 2:16.477 | ע"ז | , |
| 2:16.478 | כדאמירן |  |
| 2:16.479 | במגלה |  |
| 2:16.480 | יט | ), |
| 2:16.481 | דאי |  |
| 2:16.482 | לאו |  |
| 2:16.483 | הכי |  |
| 2:16.484 | לא |  |
| 2:16.485 | היה |  |
| 2:16.486 | מרדכי |  |
| 2:16.487 | נמנע |  |
| 2:16.488 | מלהשתחות |  |
| 2:16.489 | לו | , |
| 2:16.490 | כדקתני |  |
| 2:16.491 | בבריתא | : |
| 2:16.492 | לא |  |
| 2:16.493 | תשתחוה |  |
| 2:16.494 | להם | , |
| 2:16.495 | להם |  |
| 2:16.496 | אי |  |
| 2:16.497 | אתה |  |
| 2:16.498 | משתחוה | , |
| 2:16.499 | אבל |  |
| 2:16.500 | אתה |  |
| 2:16.501 | משתחוה |  |
| 2:16.502 | לאדם |  |
| 2:16.503 | כמותך | , |
| 2:16.504 | והוא |  |
| 2:16.505 | שלא |  |
| 2:16.506 | יהא |  |
| 2:16.507 | נעבד |  |
| 2:16.508 | כהמן |  |
| 2:16.509 | שעשה |  |
| 2:16.510 | עצמו |  |
| 2:16.511 | אלוה | , |
| 2:16.512 | רק |  |
| 2:16.513 | משתחוה |  |
| 2:16.514 | לו |  |
| 2:16.515 | לאות |  |
| 2:16.516 | הכבוד |  |
| 2:16.517 | ולא |  |
| 2:16.518 | לאלהות | , |
| 2:16.519 | ולכך |  |
| 2:16.520 | השתחוה |  |
| 2:16.521 | יהושע |  |
| 2:16.522 | למלאך |  |
| 2:16.523 | הואיל |  |
| 2:16.524 | ונגלה |  |
| 2:16.525 | לו |  |
| 2:16.526 | בדמות |  |
| 2:16.527 | אדם | , |
| 2:16.528 | כמ"ש | : |
| 2:16.529 | והנה |  |
| 2:16.530 | איש |  |
| 2:16.531 | עומד |  |
| 2:16.532 | לנגדו | , |
| 2:16.533 | ולכן |  |
| 2:16.534 | השתחוה |  |
| 2:16.535 | לו |  |
| 2:16.536 | לאות |  |
| 2:16.537 | הכבוד |  |
| 2:16.538 | ולא |  |
| 2:16.539 | לעבדו | . |
| 2:16.540 | ועיין |  |
| 2:16.541 | בס | ' |
| 2:16.542 | העיקרים |  |
| 2:16.543 | מאמר |  |
| 2:16.544 | ב | ' |
| 2:16.545 | פ | ' |
| 2:16.546 | כ"ח |  |
| 2:16.547 | שהאריך |  |
| 2:16.548 | בביאור |  |
| 2:16.549 | הכתובים |  |
| 2:16.550 | שקורין |  |
| 2:16.551 | לכבוד |  |
| 2:16.552 | הנראה |  |
| 2:16.553 | לנביאים |  |
| 2:16.554 | בשם |  |
| 2:16.555 | הוי"ה | , |
| 2:16.556 | וגם |  |
| 2:16.557 | לארון | , |
| 2:16.558 | כמו |  |
| 2:16.559 | שגאמר | : |
| 2:16.560 | ויהי |  |
| 2:16.561 | בנסע |  |
| 2:16.562 | הארון |  |
| 2:16.563 | וכו | ' |
| 2:16.564 | קומה |  |
| 2:16.565 | ה | ', |
| 2:16.566 | ובנחה |  |
| 2:16.567 | יאמר |  |
| 2:16.568 | שובה |  |
| 2:16.569 | ה | '", |
| 2:16.570 | ויהושע |  |
| 2:16.571 | קרא |  |
| 2:16.572 | לארון |  |
| 2:16.573 | אדון |  |
| 2:16.574 | כל |  |
| 2:16.575 | הארץ | " |
| 2:16.576 | ע"ש | . |
| 2:16.577 | ולדעתו |  |
| 2:16.578 | שמה |  |
| 2:16.579 | שהשתחוה |  |
| 2:16.580 | יהושע |  |
| 2:16.581 | למלאך | , |
| 2:16.582 | לפי |  |
| 2:16.583 | שההשתחויה |  |
| 2:16.584 | ההיא |  |
| 2:16.585 | היתה |  |
| 2:16.586 | לכבוד |  |
| 2:16.587 | השי"ת | , |
| 2:16.588 | ולהיות |  |
| 2:16.589 | זה |  |
| 2:16.590 | שלוחו |  |
| 2:16.591 | וכו | ' |
| 2:16.592 | ע"ש | . |
| 2:16.593 | ורז"ל |  |
| 2:16.594 | אמרו |  |
| 2:16.595 | תעלא |  |
| 2:16.596 | בעידניה |  |
| 2:16.597 | סגיד |  |
| 2:16.598 | ליה | , |
| 2:16.599 | ר"ל | : |
| 2:16.600 | אדם |  |
| 2:16.601 | פחות |  |
| 2:16.602 | ובזוי |  |
| 2:16.603 | כשועל | , |
| 2:16.604 | והשעה |  |
| 2:16.605 | עומדת |  |
| 2:16.606 | לו |  |
| 2:16.607 | באיזה |  |
| 2:16.608 | משרה |  |
| 2:16.609 | שבידו | , |
| 2:16.610 | סגיד |  |
| 2:16.611 | ליה | . |
| 2:16.612 | וכן |  |
| 2:16.613 | הוא |  |
| 2:16.614 | נהוג |  |
| 2:16.615 | אצלנו |  |
| 2:16.616 | ואצל |  |
| 2:16.617 | הערביים |  |
| 2:16.618 | להשתחוות |  |
| 2:16.619 | ולנשק |  |
| 2:16.620 | אדם |  |
| 2:16.621 | גדול |  |
| 2:16.622 | המעלה | , |
| 2:16.623 | ואין |  |
| 2:16.624 | בזה |  |
| 2:16.625 | שום |  |
| 2:16.626 | איסור | , |
| 2:16.627 | שאינו |  |
| 2:16.628 | מקבלו |  |
| 2:16.629 | עליו |  |
| 2:16.630 | באלוה | , |
| 2:16.631 | רק |  |
| 2:16.632 | לאות |  |
| 2:16.633 | כבוד |  |
| 2:16.634 | הוא |  |
| 2:16.635 | משתחוה |  |
| 2:16.636 | לו |  |
| 2:16.637 | ומנשק |  |
| 2:16.638 | את |  |
| 2:16.639 | רגליו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

