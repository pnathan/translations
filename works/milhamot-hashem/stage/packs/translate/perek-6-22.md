# Translation pack: milhamot-hashem perek-6-22 (Milchamot ha-Shem 6:33–6:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-22`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 6:33

Hebrew (Wikisource): ובגמרא מכות דף י"כ גרסינן אמר ר"ל שלש טעויות עתיד לטעות שרה של אדם ושנ' מי זה בא מאדום חמוץ בגדים מבצרה טועה שאינה קולטת אלא בצר והוא גולה לבצרה טועה שאינה קולטת אלא אדם והוא מלאך, טועה שאינה קולטת אלא שוגג והוא מזיד, ע"כ ופי' רש"י ז"ל שרה של אדום סמאל שיברה לבצרה, שבתחלה יפרע הקב"ה ממנו כשיגיע קץ אדום ליחרב כדכתיב יפקוד ה' על צבא המרום במרום, ואחר כך ועל מלכי האדמה באדמה, חמוץ בגדים מדמו של סמאל, ואע"פ שאין המלאכים בשר ודם, כתב בו הכתוב בעין הריגת אדם לשבר את האוזן מה שהיא יכולה לשמוע עכ"ל הרי בהדיא שסמא"ל מלאך ולא אלוה, ויש לשאול לפי שיטת המקובלים החדשים, כיון שיש לכל אומה שר הממונה עליה, ושר של ישראל מיכאל, ושל מצרים דהב, ושל אדום סמאל וכו'. למה נתנו המקובלים החדשים כל זו הגדולה לשמאל להקרא בשם אל אחר, ולהיות לו עשר ספירות דקליפה, ואדם בליעל ואריך ואבא ואמא וזעיר ונבה דקליפה. ואמרו שגם זה לעומת זה עשה האלהים, היכלות הטומאה נגד היכלות הקדושה, עם כל הפרצופים שלה, ולמה לא עשו כן לשאר כל השרים של כל האומות? ולמה יגרע מיכאל השר של ישראל מכל אלו המעלות? ולמה הניחוהו בלא אשה, ושמאל השיאוהו בלילית עם ת"פ מחנותיה? כלל הדברים לא מציונ בתורה ולא בנביאים ולא בכתובים ולא בדברי רבותינו ז"ל מציאות אל אחד בעולם, לא קדושים ולא טמאים זולתי ה' אלהינו ב"ה. והכל משרתיו עושי רצונו. ורק לפי מחשבת עובדיהן דברה תורה, שהם בוחרים להם כל אחד מכל הנבראים שהמציא הקב"ה בעולם ומשתחוים לו ועובדים ומקריבים לו קרבנות, וחושבין אותו לאמצעי או לאלוה. כמבואר בדברי הרמב"ם ז"ל בפ"א מלה' ע"ז. יש שבחרו בשמש ויש שבחרו בירח או באחד הכוכבים או המזלות או לאש או למים לעבדם ולהשתחוות להם. ועל כולם הזהירנו השי"ת ואמר לא יהיה לך אלהים אחרים, לא תשתחוה לאל אחר, לא שיש מציאות אל אחר בעולם מבלעדי ה' אלהינו. ולפי שבאמת אינן אלוהות כמאמר רז"ל, לכן תרגם רבי' סעדיה אלהים אחרים מעבוד אלר, ולא תרגם אלאה אלר. וכן תרגם זבח לאלהים יחרם ומן דבח לאלמעודעת פאלי'קתל לפי שבאמת אינן אלוהות.

| id | word | punctuation after |
|---|---|---|
| 6:33.1 | ובגמרא |  |
| 6:33.2 | מכות |  |
| 6:33.3 | דף |  |
| 6:33.4 | י"כ |  |
| 6:33.5 | גרסינן |  |
| 6:33.6 | אמר |  |
| 6:33.7 | ר"ל |  |
| 6:33.8 | שלש |  |
| 6:33.9 | טעויות |  |
| 6:33.10 | עתיד |  |
| 6:33.11 | לטעות |  |
| 6:33.12 | שרה |  |
| 6:33.13 | של |  |
| 6:33.14 | אדם |  |
| 6:33.15 | ושנ | ' |
| 6:33.16 | מי |  |
| 6:33.17 | זה |  |
| 6:33.18 | בא |  |
| 6:33.19 | מאדום |  |
| 6:33.20 | חמוץ |  |
| 6:33.21 | בגדים |  |
| 6:33.22 | מבצרה |  |
| 6:33.23 | טועה |  |
| 6:33.24 | שאינה |  |
| 6:33.25 | קולטת |  |
| 6:33.26 | אלא |  |
| 6:33.27 | בצר |  |
| 6:33.28 | והוא |  |
| 6:33.29 | גולה |  |
| 6:33.30 | לבצרה |  |
| 6:33.31 | טועה |  |
| 6:33.32 | שאינה |  |
| 6:33.33 | קולטת |  |
| 6:33.34 | אלא |  |
| 6:33.35 | אדם |  |
| 6:33.36 | והוא |  |
| 6:33.37 | מלאך | , |
| 6:33.38 | טועה |  |
| 6:33.39 | שאינה |  |
| 6:33.40 | קולטת |  |
| 6:33.41 | אלא |  |
| 6:33.42 | שוגג |  |
| 6:33.43 | והוא |  |
| 6:33.44 | מזיד | , |
| 6:33.45 | ע"כ |  |
| 6:33.46 | ופי | ' |
| 6:33.47 | רש"י |  |
| 6:33.48 | ז"ל |  |
| 6:33.49 | שרה |  |
| 6:33.50 | של |  |
| 6:33.51 | אדום |  |
| 6:33.52 | סמאל |  |
| 6:33.53 | שיברה |  |
| 6:33.54 | לבצרה | , |
| 6:33.55 | שבתחלה |  |
| 6:33.56 | יפרע |  |
| 6:33.57 | הקב"ה |  |
| 6:33.58 | ממנו |  |
| 6:33.59 | כשיגיע |  |
| 6:33.60 | קץ |  |
| 6:33.61 | אדום |  |
| 6:33.62 | ליחרב |  |
| 6:33.63 | כדכתיב |  |
| 6:33.64 | יפקוד |  |
| 6:33.65 | ה | ' |
| 6:33.66 | על |  |
| 6:33.67 | צבא |  |
| 6:33.68 | המרום |  |
| 6:33.69 | במרום | , |
| 6:33.70 | ואחר |  |
| 6:33.71 | כך |  |
| 6:33.72 | ועל |  |
| 6:33.73 | מלכי |  |
| 6:33.74 | האדמה |  |
| 6:33.75 | באדמה | , |
| 6:33.76 | חמוץ |  |
| 6:33.77 | בגדים |  |
| 6:33.78 | מדמו |  |
| 6:33.79 | של |  |
| 6:33.80 | סמאל | , |
| 6:33.81 | ואע"פ |  |
| 6:33.82 | שאין |  |
| 6:33.83 | המלאכים |  |
| 6:33.84 | בשר |  |
| 6:33.85 | ודם | , |
| 6:33.86 | כתב |  |
| 6:33.87 | בו |  |
| 6:33.88 | הכתוב |  |
| 6:33.89 | בעין |  |
| 6:33.90 | הריגת |  |
| 6:33.91 | אדם |  |
| 6:33.92 | לשבר |  |
| 6:33.93 | את |  |
| 6:33.94 | האוזן |  |
| 6:33.95 | מה |  |
| 6:33.96 | שהיא |  |
| 6:33.97 | יכולה |  |
| 6:33.98 | לשמוע |  |
| 6:33.99 | עכ"ל |  |
| 6:33.100 | הרי |  |
| 6:33.101 | בהדיא |  |
| 6:33.102 | שסמא"ל |  |
| 6:33.103 | מלאך |  |
| 6:33.104 | ולא |  |
| 6:33.105 | אלוה | , |
| 6:33.106 | ויש |  |
| 6:33.107 | לשאול |  |
| 6:33.108 | לפי |  |
| 6:33.109 | שיטת |  |
| 6:33.110 | המקובלים |  |
| 6:33.111 | החדשים | , |
| 6:33.112 | כיון |  |
| 6:33.113 | שיש |  |
| 6:33.114 | לכל |  |
| 6:33.115 | אומה |  |
| 6:33.116 | שר |  |
| 6:33.117 | הממונה |  |
| 6:33.118 | עליה | , |
| 6:33.119 | ושר |  |
| 6:33.120 | של |  |
| 6:33.121 | ישראל |  |
| 6:33.122 | מיכאל | , |
| 6:33.123 | ושל |  |
| 6:33.124 | מצרים |  |
| 6:33.125 | דהב | , |
| 6:33.126 | ושל |  |
| 6:33.127 | אדום |  |
| 6:33.128 | סמאל |  |
| 6:33.129 | וכו | '. |
| 6:33.130 | למה |  |
| 6:33.131 | נתנו |  |
| 6:33.132 | המקובלים |  |
| 6:33.133 | החדשים |  |
| 6:33.134 | כל |  |
| 6:33.135 | זו |  |
| 6:33.136 | הגדולה |  |
| 6:33.137 | לשמאל |  |
| 6:33.138 | להקרא |  |
| 6:33.139 | בשם |  |
| 6:33.140 | אל |  |
| 6:33.141 | אחר | , |
| 6:33.142 | ולהיות |  |
| 6:33.143 | לו |  |
| 6:33.144 | עשר |  |
| 6:33.145 | ספירות |  |
| 6:33.146 | דקליפה | , |
| 6:33.147 | ואדם |  |
| 6:33.148 | בליעל |  |
| 6:33.149 | ואריך |  |
| 6:33.150 | ואבא |  |
| 6:33.151 | ואמא |  |
| 6:33.152 | וזעיר |  |
| 6:33.153 | ונבה |  |
| 6:33.154 | דקליפה | . |
| 6:33.155 | ואמרו |  |
| 6:33.156 | שגם |  |
| 6:33.157 | זה |  |
| 6:33.158 | לעומת |  |
| 6:33.159 | זה |  |
| 6:33.160 | עשה |  |
| 6:33.161 | האלהים | , |
| 6:33.162 | היכלות |  |
| 6:33.163 | הטומאה |  |
| 6:33.164 | נגד |  |
| 6:33.165 | היכלות |  |
| 6:33.166 | הקדושה | , |
| 6:33.167 | עם |  |
| 6:33.168 | כל |  |
| 6:33.169 | הפרצופים |  |
| 6:33.170 | שלה | , |
| 6:33.171 | ולמה |  |
| 6:33.172 | לא |  |
| 6:33.173 | עשו |  |
| 6:33.174 | כן |  |
| 6:33.175 | לשאר |  |
| 6:33.176 | כל |  |
| 6:33.177 | השרים |  |
| 6:33.178 | של |  |
| 6:33.179 | כל |  |
| 6:33.180 | האומות | ? |
| 6:33.181 | ולמה |  |
| 6:33.182 | יגרע |  |
| 6:33.183 | מיכאל |  |
| 6:33.184 | השר |  |
| 6:33.185 | של |  |
| 6:33.186 | ישראל |  |
| 6:33.187 | מכל |  |
| 6:33.188 | אלו |  |
| 6:33.189 | המעלות | ? |
| 6:33.190 | ולמה |  |
| 6:33.191 | הניחוהו |  |
| 6:33.192 | בלא |  |
| 6:33.193 | אשה | , |
| 6:33.194 | ושמאל |  |
| 6:33.195 | השיאוהו |  |
| 6:33.196 | בלילית |  |
| 6:33.197 | עם |  |
| 6:33.198 | ת"פ |  |
| 6:33.199 | מחנותיה | ? |
| 6:33.200 | כלל |  |
| 6:33.201 | הדברים |  |
| 6:33.202 | לא |  |
| 6:33.203 | מציונ |  |
| 6:33.204 | בתורה |  |
| 6:33.205 | ולא |  |
| 6:33.206 | בנביאים |  |
| 6:33.207 | ולא |  |
| 6:33.208 | בכתובים |  |
| 6:33.209 | ולא |  |
| 6:33.210 | בדברי |  |
| 6:33.211 | רבותינו |  |
| 6:33.212 | ז"ל |  |
| 6:33.213 | מציאות |  |
| 6:33.214 | אל |  |
| 6:33.215 | אחד |  |
| 6:33.216 | בעולם | , |
| 6:33.217 | לא |  |
| 6:33.218 | קדושים |  |
| 6:33.219 | ולא |  |
| 6:33.220 | טמאים |  |
| 6:33.221 | זולתי |  |
| 6:33.222 | ה | ' |
| 6:33.223 | אלהינו |  |
| 6:33.224 | ב"ה | . |
| 6:33.225 | והכל |  |
| 6:33.226 | משרתיו |  |
| 6:33.227 | עושי |  |
| 6:33.228 | רצונו | . |
| 6:33.229 | ורק |  |
| 6:33.230 | לפי |  |
| 6:33.231 | מחשבת |  |
| 6:33.232 | עובדיהן |  |
| 6:33.233 | דברה |  |
| 6:33.234 | תורה | , |
| 6:33.235 | שהם |  |
| 6:33.236 | בוחרים |  |
| 6:33.237 | להם |  |
| 6:33.238 | כל |  |
| 6:33.239 | אחד |  |
| 6:33.240 | מכל |  |
| 6:33.241 | הנבראים |  |
| 6:33.242 | שהמציא |  |
| 6:33.243 | הקב"ה |  |
| 6:33.244 | בעולם |  |
| 6:33.245 | ומשתחוים |  |
| 6:33.246 | לו |  |
| 6:33.247 | ועובדים |  |
| 6:33.248 | ומקריבים |  |
| 6:33.249 | לו |  |
| 6:33.250 | קרבנות | , |
| 6:33.251 | וחושבין |  |
| 6:33.252 | אותו |  |
| 6:33.253 | לאמצעי |  |
| 6:33.254 | או |  |
| 6:33.255 | לאלוה | . |
| 6:33.256 | כמבואר |  |
| 6:33.257 | בדברי |  |
| 6:33.258 | הרמב"ם |  |
| 6:33.259 | ז"ל |  |
| 6:33.260 | בפ"א |  |
| 6:33.261 | מלה | ' |
| 6:33.262 | ע"ז | . |
| 6:33.263 | יש |  |
| 6:33.264 | שבחרו |  |
| 6:33.265 | בשמש |  |
| 6:33.266 | ויש |  |
| 6:33.267 | שבחרו |  |
| 6:33.268 | בירח |  |
| 6:33.269 | או |  |
| 6:33.270 | באחד |  |
| 6:33.271 | הכוכבים |  |
| 6:33.272 | או |  |
| 6:33.273 | המזלות |  |
| 6:33.274 | או |  |
| 6:33.275 | לאש |  |
| 6:33.276 | או |  |
| 6:33.277 | למים |  |
| 6:33.278 | לעבדם |  |
| 6:33.279 | ולהשתחוות |  |
| 6:33.280 | להם | . |
| 6:33.281 | ועל |  |
| 6:33.282 | כולם |  |
| 6:33.283 | הזהירנו |  |
| 6:33.284 | השי"ת |  |
| 6:33.285 | ואמר |  |
| 6:33.286 | לא |  |
| 6:33.287 | יהיה |  |
| 6:33.288 | לך |  |
| 6:33.289 | אלהים |  |
| 6:33.290 | אחרים | , |
| 6:33.291 | לא |  |
| 6:33.292 | תשתחוה |  |
| 6:33.293 | לאל |  |
| 6:33.294 | אחר | , |
| 6:33.295 | לא |  |
| 6:33.296 | שיש |  |
| 6:33.297 | מציאות |  |
| 6:33.298 | אל |  |
| 6:33.299 | אחר |  |
| 6:33.300 | בעולם |  |
| 6:33.301 | מבלעדי |  |
| 6:33.302 | ה | ' |
| 6:33.303 | אלהינו | . |
| 6:33.304 | ולפי |  |
| 6:33.305 | שבאמת |  |
| 6:33.306 | אינן |  |
| 6:33.307 | אלוהות |  |
| 6:33.308 | כמאמר |  |
| 6:33.309 | רז"ל | , |
| 6:33.310 | לכן |  |
| 6:33.311 | תרגם |  |
| 6:33.312 | רבי | ' |
| 6:33.313 | סעדיה |  |
| 6:33.314 | אלהים |  |
| 6:33.315 | אחרים |  |
| 6:33.316 | מעבוד |  |
| 6:33.317 | אלר | , |
| 6:33.318 | ולא |  |
| 6:33.319 | תרגם |  |
| 6:33.320 | אלאה |  |
| 6:33.321 | אלר | . |
| 6:33.322 | וכן |  |
| 6:33.323 | תרגם |  |
| 6:33.324 | זבח |  |
| 6:33.325 | לאלהים |  |
| 6:33.326 | יחרם |  |
| 6:33.327 | ומן |  |
| 6:33.328 | דבח |  |
| 6:33.329 | לאלמעודעת |  |
| 6:33.330 | פאלי'קתל |  |
| 6:33.331 | לפי |  |
| 6:33.332 | שבאמת |  |
| 6:33.333 | אינן |  |
| 6:33.334 | אלוהות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

