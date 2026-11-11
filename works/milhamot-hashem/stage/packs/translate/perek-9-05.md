# Translation pack: milhamot-hashem perek-9-05 (Milchamot ha-Shem 9:5–9:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-05`.

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
- **אצילות** and its root: fixed and translated, so that the verbs the
  author strings together stay linked. אצילות → "emanation"; עולם האצילות /
  דאצילות → "the World of Emanation"; נאצל → "emanated"; הנאצלים →
  "the emanated beings"; המאציל → "the Emanator". The four worlds:
  "Emanation, Creation, Formation, Making" (אבי"ע). "נאצלו ונבראו ונוצרו
  ונעשו" → "were emanated, created, formed and made". Keep נאצל and נברא
  distinct everywhere: the author argues that Ze'ir Anpin is "נברא ולא בורא"
  → "created, not creator".
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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
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
- דברי תימה: "astonishing words" (words that astonish).
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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


## Milchamot ha-Shem 9:5

Hebrew (Wikisource): קכד) ועל הירא את דבר ה' ומשכיל בתורתו ובקבלת רבותינו במשנה ובתלמוד אל יפת בסתר לבו לעבוד שום צורה ופרצוף אשר לפי דבריהם חלק להם האלוה העליון מעצמותו וישם את חלק החכמה שבו נשמה לחכמה ואת חלק הבינה שיש בו נשמה לבינה הנקראת אימא ואת חלק הגדולה שבו שם בגדולה ואת חלק הגבורה שם בגבורה וכו'. ויהיו החלקים האלה להם לנשמות וגם את הגופים הרקים העליונים שמו אותם לנשמות לגופים בתחתונים מהם וקיימו וקבלו עליהם לעבדם ואמרו שבעברינו אותם אנחנו עובדים לנשמתם שהיא הסבה הראשונה ממש.

| id | word | punctuation after |
|---|---|---|
| 9:5.1 | קכד | ) |
| 9:5.2 | ועל |  |
| 9:5.3 | הירא |  |
| 9:5.4 | את |  |
| 9:5.5 | דבר |  |
| 9:5.6 | ה | ' |
| 9:5.7 | ומשכיל |  |
| 9:5.8 | בתורתו |  |
| 9:5.9 | ובקבלת |  |
| 9:5.10 | רבותינו |  |
| 9:5.11 | במשנה |  |
| 9:5.12 | ובתלמוד |  |
| 9:5.13 | אל |  |
| 9:5.14 | יפת |  |
| 9:5.15 | בסתר |  |
| 9:5.16 | לבו |  |
| 9:5.17 | לעבוד |  |
| 9:5.18 | שום |  |
| 9:5.19 | צורה |  |
| 9:5.20 | ופרצוף |  |
| 9:5.21 | אשר |  |
| 9:5.22 | לפי |  |
| 9:5.23 | דבריהם |  |
| 9:5.24 | חלק |  |
| 9:5.25 | להם |  |
| 9:5.26 | האלוה |  |
| 9:5.27 | העליון |  |
| 9:5.28 | מעצמותו |  |
| 9:5.29 | וישם |  |
| 9:5.30 | את |  |
| 9:5.31 | חלק |  |
| 9:5.32 | החכמה |  |
| 9:5.33 | שבו |  |
| 9:5.34 | נשמה |  |
| 9:5.35 | לחכמה |  |
| 9:5.36 | ואת |  |
| 9:5.37 | חלק |  |
| 9:5.38 | הבינה |  |
| 9:5.39 | שיש |  |
| 9:5.40 | בו |  |
| 9:5.41 | נשמה |  |
| 9:5.42 | לבינה |  |
| 9:5.43 | הנקראת |  |
| 9:5.44 | אימא |  |
| 9:5.45 | ואת |  |
| 9:5.46 | חלק |  |
| 9:5.47 | הגדולה |  |
| 9:5.48 | שבו |  |
| 9:5.49 | שם |  |
| 9:5.50 | בגדולה |  |
| 9:5.51 | ואת |  |
| 9:5.52 | חלק |  |
| 9:5.53 | הגבורה |  |
| 9:5.54 | שם |  |
| 9:5.55 | בגבורה |  |
| 9:5.56 | וכו | '. |
| 9:5.57 | ויהיו |  |
| 9:5.58 | החלקים |  |
| 9:5.59 | האלה |  |
| 9:5.60 | להם |  |
| 9:5.61 | לנשמות |  |
| 9:5.62 | וגם |  |
| 9:5.63 | את |  |
| 9:5.64 | הגופים |  |
| 9:5.65 | הרקים |  |
| 9:5.66 | העליונים |  |
| 9:5.67 | שמו |  |
| 9:5.68 | אותם |  |
| 9:5.69 | לנשמות |  |
| 9:5.70 | לגופים |  |
| 9:5.71 | בתחתונים |  |
| 9:5.72 | מהם |  |
| 9:5.73 | וקיימו |  |
| 9:5.74 | וקבלו |  |
| 9:5.75 | עליהם |  |
| 9:5.76 | לעבדם |  |
| 9:5.77 | ואמרו |  |
| 9:5.78 | שבעברינו |  |
| 9:5.79 | אותם |  |
| 9:5.80 | אנחנו |  |
| 9:5.81 | עובדים |  |
| 9:5.82 | לנשמתם |  |
| 9:5.83 | שהיא |  |
| 9:5.84 | הסבה |  |
| 9:5.85 | הראשונה |  |
| 9:5.86 | ממש | . |


## Milchamot ha-Shem 9:6

Hebrew (Wikisource): והדעה הזאת לקוחה מהפילסופים החיצונים (שהזכירם הר' המורה פרק כ"ט לשלישי) שאמרו שהאלוה הוא נשמת הגלגלים והכוכבים והמזלות ולכן היו עובדים להם ואומרים שהם עובדים לנשמתם שהיא הסבה הראשונה לדבריהם. וברעיונותיו המציא את הפרצופים האלה מאדם קדמאה עד זעיר ונוקביה ושלשל אותם מהסבה הראשונה אשר צמצמה ברעיונו וסלקה למעלה וחזר ושלש ממנה כל אותן הפרצופים בעולמות אין קץ וגם בעולמות שלמטה מעולם האצילות. וישם את החוט המשתלשל בהם נשמה להם ויצו את קוראיו לבדם בטענת שלא את גופם עובדים רק לנשמתם כמו שעשו הפילסופים הראשונים החיצונים בכוכבים ובמזלות. ואת כל הפרצופים קרא בשם מעשה הרכבה וישם את חיות המרכבה אלהות בעבור החלק המשתלשל מהאלהים מאלהים המצומק אל הסביבות לפי דבריו. אשר מהחלק הזה היורד דרך צנור דק נתהוו כל הגופות עם עיגוליהם וממנו גם נשמותיהן ומלבושיהן ויהיו להם לאלוהות ויקראו שמותיהן בשמותיו של הקב"ה וכינוייו. והאלהים העליון לא השאירו לו לא שם ולא כינוי ועשוהו מסולק וריקן מכל שבח ותהלה וברכה לפי שאין לו ממי להתברך וגזרו אומר שהמתפלל אליו אין תפלתו תפלה ולא יענהו.

| id | word | punctuation after |
|---|---|---|
| 9:6.1 | והדעה |  |
| 9:6.2 | הזאת |  |
| 9:6.3 | לקוחה |  |
| 9:6.4 | מהפילסופים |  |
| 9:6.5 | החיצונים |  |
| 9:6.6 | שהזכירם |  |
| 9:6.7 | הר | ' |
| 9:6.8 | המורה |  |
| 9:6.9 | פרק |  |
| 9:6.10 | כ"ט |  |
| 9:6.11 | לשלישי | ) |
| 9:6.12 | שאמרו |  |
| 9:6.13 | שהאלוה |  |
| 9:6.14 | הוא |  |
| 9:6.15 | נשמת |  |
| 9:6.16 | הגלגלים |  |
| 9:6.17 | והכוכבים |  |
| 9:6.18 | והמזלות |  |
| 9:6.19 | ולכן |  |
| 9:6.20 | היו |  |
| 9:6.21 | עובדים |  |
| 9:6.22 | להם |  |
| 9:6.23 | ואומרים |  |
| 9:6.24 | שהם |  |
| 9:6.25 | עובדים |  |
| 9:6.26 | לנשמתם |  |
| 9:6.27 | שהיא |  |
| 9:6.28 | הסבה |  |
| 9:6.29 | הראשונה |  |
| 9:6.30 | לדבריהם | . |
| 9:6.31 | וברעיונותיו |  |
| 9:6.32 | המציא |  |
| 9:6.33 | את |  |
| 9:6.34 | הפרצופים |  |
| 9:6.35 | האלה |  |
| 9:6.36 | מאדם |  |
| 9:6.37 | קדמאה |  |
| 9:6.38 | עד |  |
| 9:6.39 | זעיר |  |
| 9:6.40 | ונוקביה |  |
| 9:6.41 | ושלשל |  |
| 9:6.42 | אותם |  |
| 9:6.43 | מהסבה |  |
| 9:6.44 | הראשונה |  |
| 9:6.45 | אשר |  |
| 9:6.46 | צמצמה |  |
| 9:6.47 | ברעיונו |  |
| 9:6.48 | וסלקה |  |
| 9:6.49 | למעלה |  |
| 9:6.50 | וחזר |  |
| 9:6.51 | ושלש |  |
| 9:6.52 | ממנה |  |
| 9:6.53 | כל |  |
| 9:6.54 | אותן |  |
| 9:6.55 | הפרצופים |  |
| 9:6.56 | בעולמות |  |
| 9:6.57 | אין |  |
| 9:6.58 | קץ |  |
| 9:6.59 | וגם |  |
| 9:6.60 | בעולמות |  |
| 9:6.61 | שלמטה |  |
| 9:6.62 | מעולם |  |
| 9:6.63 | האצילות | . |
| 9:6.64 | וישם |  |
| 9:6.65 | את |  |
| 9:6.66 | החוט |  |
| 9:6.67 | המשתלשל |  |
| 9:6.68 | בהם |  |
| 9:6.69 | נשמה |  |
| 9:6.70 | להם |  |
| 9:6.71 | ויצו |  |
| 9:6.72 | את |  |
| 9:6.73 | קוראיו |  |
| 9:6.74 | לבדם |  |
| 9:6.75 | בטענת |  |
| 9:6.76 | שלא |  |
| 9:6.77 | את |  |
| 9:6.78 | גופם |  |
| 9:6.79 | עובדים |  |
| 9:6.80 | רק |  |
| 9:6.81 | לנשמתם |  |
| 9:6.82 | כמו |  |
| 9:6.83 | שעשו |  |
| 9:6.84 | הפילסופים |  |
| 9:6.85 | הראשונים |  |
| 9:6.86 | החיצונים |  |
| 9:6.87 | בכוכבים |  |
| 9:6.88 | ובמזלות | . |
| 9:6.89 | ואת |  |
| 9:6.90 | כל |  |
| 9:6.91 | הפרצופים |  |
| 9:6.92 | קרא |  |
| 9:6.93 | בשם |  |
| 9:6.94 | מעשה |  |
| 9:6.95 | הרכבה |  |
| 9:6.96 | וישם |  |
| 9:6.97 | את |  |
| 9:6.98 | חיות |  |
| 9:6.99 | המרכבה |  |
| 9:6.100 | אלהות |  |
| 9:6.101 | בעבור |  |
| 9:6.102 | החלק |  |
| 9:6.103 | המשתלשל |  |
| 9:6.104 | מהאלהים |  |
| 9:6.105 | מאלהים |  |
| 9:6.106 | המצומק |  |
| 9:6.107 | אל |  |
| 9:6.108 | הסביבות |  |
| 9:6.109 | לפי |  |
| 9:6.110 | דבריו | . |
| 9:6.111 | אשר |  |
| 9:6.112 | מהחלק |  |
| 9:6.113 | הזה |  |
| 9:6.114 | היורד |  |
| 9:6.115 | דרך |  |
| 9:6.116 | צנור |  |
| 9:6.117 | דק |  |
| 9:6.118 | נתהוו |  |
| 9:6.119 | כל |  |
| 9:6.120 | הגופות |  |
| 9:6.121 | עם |  |
| 9:6.122 | עיגוליהם |  |
| 9:6.123 | וממנו |  |
| 9:6.124 | גם |  |
| 9:6.125 | נשמותיהן |  |
| 9:6.126 | ומלבושיהן |  |
| 9:6.127 | ויהיו |  |
| 9:6.128 | להם |  |
| 9:6.129 | לאלוהות |  |
| 9:6.130 | ויקראו |  |
| 9:6.131 | שמותיהן |  |
| 9:6.132 | בשמותיו |  |
| 9:6.133 | של |  |
| 9:6.134 | הקב"ה |  |
| 9:6.135 | וכינוייו | . |
| 9:6.136 | והאלהים |  |
| 9:6.137 | העליון |  |
| 9:6.138 | לא |  |
| 9:6.139 | השאירו |  |
| 9:6.140 | לו |  |
| 9:6.141 | לא |  |
| 9:6.142 | שם |  |
| 9:6.143 | ולא |  |
| 9:6.144 | כינוי |  |
| 9:6.145 | ועשוהו |  |
| 9:6.146 | מסולק |  |
| 9:6.147 | וריקן |  |
| 9:6.148 | מכל |  |
| 9:6.149 | שבח |  |
| 9:6.150 | ותהלה |  |
| 9:6.151 | וברכה |  |
| 9:6.152 | לפי |  |
| 9:6.153 | שאין |  |
| 9:6.154 | לו |  |
| 9:6.155 | ממי |  |
| 9:6.156 | להתברך |  |
| 9:6.157 | וגזרו |  |
| 9:6.158 | אומר |  |
| 9:6.159 | שהמתפלל |  |
| 9:6.160 | אליו |  |
| 9:6.161 | אין |  |
| 9:6.162 | תפלתו |  |
| 9:6.163 | תפלה |  |
| 9:6.164 | ולא |  |
| 9:6.165 | יענהו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

