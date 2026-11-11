# Translation pack: milhamot-hashem perek-4-01 (Milchamot ha-Shem 4:1–4:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-01`.

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
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
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


## Milchamot ha-Shem 4:1

Hebrew (Wikisource): כ) ומעתה הט אזנך ושמע תשובות תורת משה אמת. לכבוד ה' אלהים אמת. מה שכתבת והשבת על דברי להכחיש שהם קורים לחמשה פרצופים אלוהות ואמרת ח"ו שהם אומרים כן כסות עינים היה לך שלא ראית דברי הזהר בראשית כ"ב וז"ל ויאמר אלהים נעשה אדם בצלמינו כדמותינו. סוד ה' ליריאיו. פתח ההוא סבא דסבין ואמר שמעון שמעון מאן הוא דאמר ויאמר אלהים. מאי ניהו האי אלהים! אדהכי פרח ההוא סבא דסבין ולא חמא ליה. וכמא דשמע ר' שמעון דהוה קרי ליה שמעון שמעון ולא רבי שמעון. אמר לחברוי ודאי האי הוא קב"ה דאתאמר ביה ועתיק יומין יתיב (פירוש הר' דרך אמת זקן הזקנים והוא כתר עליון. והמקדש מלך כ' סבא דסבין הוא עתיק כמו שאמר רשב"י לקמן כי אבא ואמא נקראו זקנים ועתיק ואריך נקראין סבא דסבין וקימא ליה לרשב"י שהוא עתיק ולא אריך) הא כען איהו שעתא למפתח בהאי רזא דודאי הכא אית רזא דלא אתיהיב רשו לאתגליא. פתח ואמר למלכא דהוה ליה כמה בניינין למבני. וההוא אומנא לא הוה עבד מדעם אלא מרשו דמלכא כמה דאתאמר ואהיה אצלו אמון. מלכא ודאי איהו חכמה עילאה לעילא. ועמודא דאמצעיתא מלכא לתתא, אלהים אומנא לעילא, ודא אימא עלאה אלהים אומנא לתתא ודא שכנתא ואתתא לית לה רשו למעבד מדעם בלא רשו בעלה, וכל בניינין דהוו באורח אצילותה הוה אמר אבא באמירה לגבי אימא יהא כדין וכדין ומיד הוה, כמא דאתמר ויאמר אלהים יהי אור, ויהי אור, ויאמר הוה אמר יהי אור מארי דבניינא איהו אמר, ואומנא עביד מיד, והכי כל בנינין באורח אצילותא הוה אמר יהי רקיע, יהי מארת, וכלא אתעביד מיד, כד מטא לעלמא דפירודא דאיהו עולם הנבדלים, אמר אומנא למארי בניינא, נעשה אדם בצלמינו כדמותינו, אמר מארי בניינא ודאי טב הוא למעבד ליה אבל עתיד הוא למחטי קמך בגין דאיהו בן כסיל הה"ד בן חכם ישמח אב ובן כסיל תוגת אמו, איהי אמרה, כיון דחוביה תליא באימא ולא באבא, אנא בעינא למברי ליה בדיוקנא דילי, הה"ד ויברא אלהים את האדם בצלמו, ולא בעא לאשתתפא ביה אבא, בזמנא דחב מה כתיב ובפשעיכם שולחה אמכם. אמר מלכא לאימא, ולא אמרית לך דעתיד למחטי, בההוא זמנא תריך ליה ותריך אימא עמיה, ובגין דא כתיב בן חכם ישמח אב דא אדם דאיהו באורח אצילותא, ובן כסיל דא אדם דבריאה עכ"ל. ופירש הר' שביבי נגה דמהכא משמע שהקב"ה בעצמו כביכול לא הסכים בבריאה אדם רק השכינה התחננה להיות אדם בגורלה ע"כ, ופי' הר' מקדש מלך בן חכם דא אדם דאצילות, הוא זעיר דאצילות, אדם דבריאה הוא אדם הראשון וע"כ. ולעיל מזה כ' הר' מקדש מלך בדבור המתחיל אמר אומנא פי' אימא אמרה לאבא נעשה אדם, וזהו ויאמר אלהים שהיא אימא אמרה לאבא נעשה אדם מה שאין כן כל ויאמר דמעשה בראשית שרוצה לומר ויאמר אבא לאלהים שהיא אימא (בינה) כי אבא אומר ובינה עושה ועכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:1.1 | כ | ) |
| 4:1.2 | ומעתה |  |
| 4:1.3 | הט |  |
| 4:1.4 | אזנך |  |
| 4:1.5 | ושמע |  |
| 4:1.6 | תשובות |  |
| 4:1.7 | תורת |  |
| 4:1.8 | משה |  |
| 4:1.9 | אמת | . |
| 4:1.10 | לכבוד |  |
| 4:1.11 | ה | ' |
| 4:1.12 | אלהים |  |
| 4:1.13 | אמת | . |
| 4:1.14 | מה |  |
| 4:1.15 | שכתבת |  |
| 4:1.16 | והשבת |  |
| 4:1.17 | על |  |
| 4:1.18 | דברי |  |
| 4:1.19 | להכחיש |  |
| 4:1.20 | שהם |  |
| 4:1.21 | קורים |  |
| 4:1.22 | לחמשה |  |
| 4:1.23 | פרצופים |  |
| 4:1.24 | אלוהות |  |
| 4:1.25 | ואמרת |  |
| 4:1.26 | ח"ו |  |
| 4:1.27 | שהם |  |
| 4:1.28 | אומרים |  |
| 4:1.29 | כן |  |
| 4:1.30 | כסות |  |
| 4:1.31 | עינים |  |
| 4:1.32 | היה |  |
| 4:1.33 | לך |  |
| 4:1.34 | שלא |  |
| 4:1.35 | ראית |  |
| 4:1.36 | דברי |  |
| 4:1.37 | הזהר |  |
| 4:1.38 | בראשית |  |
| 4:1.39 | כ"ב |  |
| 4:1.40 | וז"ל |  |
| 4:1.41 | ויאמר |  |
| 4:1.42 | אלהים |  |
| 4:1.43 | נעשה |  |
| 4:1.44 | אדם |  |
| 4:1.45 | בצלמינו |  |
| 4:1.46 | כדמותינו | . |
| 4:1.47 | סוד |  |
| 4:1.48 | ה | ' |
| 4:1.49 | ליריאיו | . |
| 4:1.50 | פתח |  |
| 4:1.51 | ההוא |  |
| 4:1.52 | סבא |  |
| 4:1.53 | דסבין |  |
| 4:1.54 | ואמר |  |
| 4:1.55 | שמעון |  |
| 4:1.56 | שמעון |  |
| 4:1.57 | מאן |  |
| 4:1.58 | הוא |  |
| 4:1.59 | דאמר |  |
| 4:1.60 | ויאמר |  |
| 4:1.61 | אלהים | . |
| 4:1.62 | מאי |  |
| 4:1.63 | ניהו |  |
| 4:1.64 | האי |  |
| 4:1.65 | אלהים | ! |
| 4:1.66 | אדהכי |  |
| 4:1.67 | פרח |  |
| 4:1.68 | ההוא |  |
| 4:1.69 | סבא |  |
| 4:1.70 | דסבין |  |
| 4:1.71 | ולא |  |
| 4:1.72 | חמא |  |
| 4:1.73 | ליה | . |
| 4:1.74 | וכמא |  |
| 4:1.75 | דשמע |  |
| 4:1.76 | ר | ' |
| 4:1.77 | שמעון |  |
| 4:1.78 | דהוה |  |
| 4:1.79 | קרי |  |
| 4:1.80 | ליה |  |
| 4:1.81 | שמעון |  |
| 4:1.82 | שמעון |  |
| 4:1.83 | ולא |  |
| 4:1.84 | רבי |  |
| 4:1.85 | שמעון | . |
| 4:1.86 | אמר |  |
| 4:1.87 | לחברוי |  |
| 4:1.88 | ודאי |  |
| 4:1.89 | האי |  |
| 4:1.90 | הוא |  |
| 4:1.91 | קב"ה |  |
| 4:1.92 | דאתאמר |  |
| 4:1.93 | ביה |  |
| 4:1.94 | ועתיק |  |
| 4:1.95 | יומין |  |
| 4:1.96 | יתיב |  |
| 4:1.97 | פירוש |  |
| 4:1.98 | הר | ' |
| 4:1.99 | דרך |  |
| 4:1.100 | אמת |  |
| 4:1.101 | זקן |  |
| 4:1.102 | הזקנים |  |
| 4:1.103 | והוא |  |
| 4:1.104 | כתר |  |
| 4:1.105 | עליון | . |
| 4:1.106 | והמקדש |  |
| 4:1.107 | מלך |  |
| 4:1.108 | כ | ' |
| 4:1.109 | סבא |  |
| 4:1.110 | דסבין |  |
| 4:1.111 | הוא |  |
| 4:1.112 | עתיק |  |
| 4:1.113 | כמו |  |
| 4:1.114 | שאמר |  |
| 4:1.115 | רשב"י |  |
| 4:1.116 | לקמן |  |
| 4:1.117 | כי |  |
| 4:1.118 | אבא |  |
| 4:1.119 | ואמא |  |
| 4:1.120 | נקראו |  |
| 4:1.121 | זקנים |  |
| 4:1.122 | ועתיק |  |
| 4:1.123 | ואריך |  |
| 4:1.124 | נקראין |  |
| 4:1.125 | סבא |  |
| 4:1.126 | דסבין |  |
| 4:1.127 | וקימא |  |
| 4:1.128 | ליה |  |
| 4:1.129 | לרשב"י |  |
| 4:1.130 | שהוא |  |
| 4:1.131 | עתיק |  |
| 4:1.132 | ולא |  |
| 4:1.133 | אריך | ) |
| 4:1.134 | הא |  |
| 4:1.135 | כען |  |
| 4:1.136 | איהו |  |
| 4:1.137 | שעתא |  |
| 4:1.138 | למפתח |  |
| 4:1.139 | בהאי |  |
| 4:1.140 | רזא |  |
| 4:1.141 | דודאי |  |
| 4:1.142 | הכא |  |
| 4:1.143 | אית |  |
| 4:1.144 | רזא |  |
| 4:1.145 | דלא |  |
| 4:1.146 | אתיהיב |  |
| 4:1.147 | רשו |  |
| 4:1.148 | לאתגליא | . |
| 4:1.149 | פתח |  |
| 4:1.150 | ואמר |  |
| 4:1.151 | למלכא |  |
| 4:1.152 | דהוה |  |
| 4:1.153 | ליה |  |
| 4:1.154 | כמה |  |
| 4:1.155 | בניינין |  |
| 4:1.156 | למבני | . |
| 4:1.157 | וההוא |  |
| 4:1.158 | אומנא |  |
| 4:1.159 | לא |  |
| 4:1.160 | הוה |  |
| 4:1.161 | עבד |  |
| 4:1.162 | מדעם |  |
| 4:1.163 | אלא |  |
| 4:1.164 | מרשו |  |
| 4:1.165 | דמלכא |  |
| 4:1.166 | כמה |  |
| 4:1.167 | דאתאמר |  |
| 4:1.168 | ואהיה |  |
| 4:1.169 | אצלו |  |
| 4:1.170 | אמון | . |
| 4:1.171 | מלכא |  |
| 4:1.172 | ודאי |  |
| 4:1.173 | איהו |  |
| 4:1.174 | חכמה |  |
| 4:1.175 | עילאה |  |
| 4:1.176 | לעילא | . |
| 4:1.177 | ועמודא |  |
| 4:1.178 | דאמצעיתא |  |
| 4:1.179 | מלכא |  |
| 4:1.180 | לתתא | , |
| 4:1.181 | אלהים |  |
| 4:1.182 | אומנא |  |
| 4:1.183 | לעילא | , |
| 4:1.184 | ודא |  |
| 4:1.185 | אימא |  |
| 4:1.186 | עלאה |  |
| 4:1.187 | אלהים |  |
| 4:1.188 | אומנא |  |
| 4:1.189 | לתתא |  |
| 4:1.190 | ודא |  |
| 4:1.191 | שכנתא |  |
| 4:1.192 | ואתתא |  |
| 4:1.193 | לית |  |
| 4:1.194 | לה |  |
| 4:1.195 | רשו |  |
| 4:1.196 | למעבד |  |
| 4:1.197 | מדעם |  |
| 4:1.198 | בלא |  |
| 4:1.199 | רשו |  |
| 4:1.200 | בעלה | , |
| 4:1.201 | וכל |  |
| 4:1.202 | בניינין |  |
| 4:1.203 | דהוו |  |
| 4:1.204 | באורח |  |
| 4:1.205 | אצילותה |  |
| 4:1.206 | הוה |  |
| 4:1.207 | אמר |  |
| 4:1.208 | אבא |  |
| 4:1.209 | באמירה |  |
| 4:1.210 | לגבי |  |
| 4:1.211 | אימא |  |
| 4:1.212 | יהא |  |
| 4:1.213 | כדין |  |
| 4:1.214 | וכדין |  |
| 4:1.215 | ומיד |  |
| 4:1.216 | הוה | , |
| 4:1.217 | כמא |  |
| 4:1.218 | דאתמר |  |
| 4:1.219 | ויאמר |  |
| 4:1.220 | אלהים |  |
| 4:1.221 | יהי |  |
| 4:1.222 | אור | , |
| 4:1.223 | ויהי |  |
| 4:1.224 | אור | , |
| 4:1.225 | ויאמר |  |
| 4:1.226 | הוה |  |
| 4:1.227 | אמר |  |
| 4:1.228 | יהי |  |
| 4:1.229 | אור |  |
| 4:1.230 | מארי |  |
| 4:1.231 | דבניינא |  |
| 4:1.232 | איהו |  |
| 4:1.233 | אמר | , |
| 4:1.234 | ואומנא |  |
| 4:1.235 | עביד |  |
| 4:1.236 | מיד | , |
| 4:1.237 | והכי |  |
| 4:1.238 | כל |  |
| 4:1.239 | בנינין |  |
| 4:1.240 | באורח |  |
| 4:1.241 | אצילותא |  |
| 4:1.242 | הוה |  |
| 4:1.243 | אמר |  |
| 4:1.244 | יהי |  |
| 4:1.245 | רקיע | , |
| 4:1.246 | יהי |  |
| 4:1.247 | מארת | , |
| 4:1.248 | וכלא |  |
| 4:1.249 | אתעביד |  |
| 4:1.250 | מיד | , |
| 4:1.251 | כד |  |
| 4:1.252 | מטא |  |
| 4:1.253 | לעלמא |  |
| 4:1.254 | דפירודא |  |
| 4:1.255 | דאיהו |  |
| 4:1.256 | עולם |  |
| 4:1.257 | הנבדלים | , |
| 4:1.258 | אמר |  |
| 4:1.259 | אומנא |  |
| 4:1.260 | למארי |  |
| 4:1.261 | בניינא | , |
| 4:1.262 | נעשה |  |
| 4:1.263 | אדם |  |
| 4:1.264 | בצלמינו |  |
| 4:1.265 | כדמותינו | , |
| 4:1.266 | אמר |  |
| 4:1.267 | מארי |  |
| 4:1.268 | בניינא |  |
| 4:1.269 | ודאי |  |
| 4:1.270 | טב |  |
| 4:1.271 | הוא |  |
| 4:1.272 | למעבד |  |
| 4:1.273 | ליה |  |
| 4:1.274 | אבל |  |
| 4:1.275 | עתיד |  |
| 4:1.276 | הוא |  |
| 4:1.277 | למחטי |  |
| 4:1.278 | קמך |  |
| 4:1.279 | בגין |  |
| 4:1.280 | דאיהו |  |
| 4:1.281 | בן |  |
| 4:1.282 | כסיל |  |
| 4:1.283 | הה"ד |  |
| 4:1.284 | בן |  |
| 4:1.285 | חכם |  |
| 4:1.286 | ישמח |  |
| 4:1.287 | אב |  |
| 4:1.288 | ובן |  |
| 4:1.289 | כסיל |  |
| 4:1.290 | תוגת |  |
| 4:1.291 | אמו | , |
| 4:1.292 | איהי |  |
| 4:1.293 | אמרה | , |
| 4:1.294 | כיון |  |
| 4:1.295 | דחוביה |  |
| 4:1.296 | תליא |  |
| 4:1.297 | באימא |  |
| 4:1.298 | ולא |  |
| 4:1.299 | באבא | , |
| 4:1.300 | אנא |  |
| 4:1.301 | בעינא |  |
| 4:1.302 | למברי |  |
| 4:1.303 | ליה |  |
| 4:1.304 | בדיוקנא |  |
| 4:1.305 | דילי | , |
| 4:1.306 | הה"ד |  |
| 4:1.307 | ויברא |  |
| 4:1.308 | אלהים |  |
| 4:1.309 | את |  |
| 4:1.310 | האדם |  |
| 4:1.311 | בצלמו | , |
| 4:1.312 | ולא |  |
| 4:1.313 | בעא |  |
| 4:1.314 | לאשתתפא |  |
| 4:1.315 | ביה |  |
| 4:1.316 | אבא | , |
| 4:1.317 | בזמנא |  |
| 4:1.318 | דחב |  |
| 4:1.319 | מה |  |
| 4:1.320 | כתיב |  |
| 4:1.321 | ובפשעיכם |  |
| 4:1.322 | שולחה |  |
| 4:1.323 | אמכם | . |
| 4:1.324 | אמר |  |
| 4:1.325 | מלכא |  |
| 4:1.326 | לאימא | , |
| 4:1.327 | ולא |  |
| 4:1.328 | אמרית |  |
| 4:1.329 | לך |  |
| 4:1.330 | דעתיד |  |
| 4:1.331 | למחטי | , |
| 4:1.332 | בההוא |  |
| 4:1.333 | זמנא |  |
| 4:1.334 | תריך |  |
| 4:1.335 | ליה |  |
| 4:1.336 | ותריך |  |
| 4:1.337 | אימא |  |
| 4:1.338 | עמיה | , |
| 4:1.339 | ובגין |  |
| 4:1.340 | דא |  |
| 4:1.341 | כתיב |  |
| 4:1.342 | בן |  |
| 4:1.343 | חכם |  |
| 4:1.344 | ישמח |  |
| 4:1.345 | אב |  |
| 4:1.346 | דא |  |
| 4:1.347 | אדם |  |
| 4:1.348 | דאיהו |  |
| 4:1.349 | באורח |  |
| 4:1.350 | אצילותא | , |
| 4:1.351 | ובן |  |
| 4:1.352 | כסיל |  |
| 4:1.353 | דא |  |
| 4:1.354 | אדם |  |
| 4:1.355 | דבריאה |  |
| 4:1.356 | עכ"ל | . |
| 4:1.357 | ופירש |  |
| 4:1.358 | הר | ' |
| 4:1.359 | שביבי |  |
| 4:1.360 | נגה |  |
| 4:1.361 | דמהכא |  |
| 4:1.362 | משמע |  |
| 4:1.363 | שהקב"ה |  |
| 4:1.364 | בעצמו |  |
| 4:1.365 | כביכול |  |
| 4:1.366 | לא |  |
| 4:1.367 | הסכים |  |
| 4:1.368 | בבריאה |  |
| 4:1.369 | אדם |  |
| 4:1.370 | רק |  |
| 4:1.371 | השכינה |  |
| 4:1.372 | התחננה |  |
| 4:1.373 | להיות |  |
| 4:1.374 | אדם |  |
| 4:1.375 | בגורלה |  |
| 4:1.376 | ע"כ | , |
| 4:1.377 | ופי | ' |
| 4:1.378 | הר | ' |
| 4:1.379 | מקדש |  |
| 4:1.380 | מלך |  |
| 4:1.381 | בן |  |
| 4:1.382 | חכם |  |
| 4:1.383 | דא |  |
| 4:1.384 | אדם |  |
| 4:1.385 | דאצילות | , |
| 4:1.386 | הוא |  |
| 4:1.387 | זעיר |  |
| 4:1.388 | דאצילות | , |
| 4:1.389 | אדם |  |
| 4:1.390 | דבריאה |  |
| 4:1.391 | הוא |  |
| 4:1.392 | אדם |  |
| 4:1.393 | הראשון |  |
| 4:1.394 | וע"כ | . |
| 4:1.395 | ולעיל |  |
| 4:1.396 | מזה |  |
| 4:1.397 | כ | ' |
| 4:1.398 | הר | ' |
| 4:1.399 | מקדש |  |
| 4:1.400 | מלך |  |
| 4:1.401 | בדבור |  |
| 4:1.402 | המתחיל |  |
| 4:1.403 | אמר |  |
| 4:1.404 | אומנא |  |
| 4:1.405 | פי | ' |
| 4:1.406 | אימא |  |
| 4:1.407 | אמרה |  |
| 4:1.408 | לאבא |  |
| 4:1.409 | נעשה |  |
| 4:1.410 | אדם | , |
| 4:1.411 | וזהו |  |
| 4:1.412 | ויאמר |  |
| 4:1.413 | אלהים |  |
| 4:1.414 | שהיא |  |
| 4:1.415 | אימא |  |
| 4:1.416 | אמרה |  |
| 4:1.417 | לאבא |  |
| 4:1.418 | נעשה |  |
| 4:1.419 | אדם |  |
| 4:1.420 | מה |  |
| 4:1.421 | שאין |  |
| 4:1.422 | כן |  |
| 4:1.423 | כל |  |
| 4:1.424 | ויאמר |  |
| 4:1.425 | דמעשה |  |
| 4:1.426 | בראשית |  |
| 4:1.427 | שרוצה |  |
| 4:1.428 | לומר |  |
| 4:1.429 | ויאמר |  |
| 4:1.430 | אבא |  |
| 4:1.431 | לאלהים |  |
| 4:1.432 | שהיא |  |
| 4:1.433 | אימא |  |
| 4:1.434 | בינה | ) |
| 4:1.435 | כי |  |
| 4:1.436 | אבא |  |
| 4:1.437 | אומר |  |
| 4:1.438 | ובינה |  |
| 4:1.439 | עושה |  |
| 4:1.440 | ועכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

