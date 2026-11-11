# Translation pack: milhamot-hashem perek-5-13 (Milchamot ha-Shem 5:20–5:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-13`.

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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

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


## Milchamot ha-Shem 5:20

Hebrew (Wikisource): אמנם היה האר"י לפעמים אומר במקומו אדון עולם אשר מלך מפני שלא נזכר בו שום שם משמות הקודש. ושייך כל אותו השבח באלוה הראשון הנקרא אין סוף אשר לא יעבדוהו משום דלית ליה שם ידיע ולא שום כינוי. וגם נזכר באדון עולם והוא יהיה בתפארה, שלפי דעתם קצתו מתלבש בתפארת, לכן לא נמנע האר"י מלאמרו. אבל ביגדל שנזכר בו שם אלהים חי שהם מיחסים אותו לפרצופי אבא ואמא וגם לזעיר אנפין אשר אותם יעבודו שיש להם פרצוף שלם בדמות אדם ובצלמו. ואינם ראשונים שכבר קדמו להם סיבות, ואינם נעלמים, לכן הקפיד מאד על ארבעה בתים הראשונים שהם שקר לפי דעתם. הלא אזן מלים תבחן כמה גדלה שגיאת רבותינו בעלי הקבלה החדשה לעקור ארבעה שרשים מיסודי תורתינו הקדושה, ובהדיא אמר הרב יושר לבב בדף ג' ע"ב: "דע את אלהי אביך כולל החמשה פרצופים וכו' ועבדהו לזעיר אנפין אע"פ שהוא נברא כי לנשמתו אתה עובד בזה כי בלעדה אין כל נמצא, משום הכי לא אמר דע אלהי אביך אלא את אלהי אביך" וכו'. ודברי יושר לבב שאמר שהוא נברא מבוארים באדרא רבה דף קל"ח ע"ב וז"ל: "ה' פעלך דא זעיר אנפין וכו' חייהו למאן, חייהו לזעיר אנפין דכל נהירו דיליה מאינון שנים קדמוניות", וכו' ועוד הרבה ראיות בלשונות הזהר שהוא נברא, ועם כל זה מצוה לעבוד אותו. ואילו שמו אל לבם דברי רז"ל בבראשית רבה ובילקוט שהבאתי לעיל סי' ל"א כל שיש לו תולדות וכו' לא עבדו לנברא. גם הטילו חסרון בידיעת השי"ת בעתיד כמו שכתבתי לעיל בשם האר"י בספר ארבע מאות שקל כסף, כי עשו ידיעת השי"ת בעתיד ידיעת ספק כידיעת האצטגנים וכן כתב הראב"ד בהשגות פ"ז מהלכות תשובה. וז"ל: "אעפ"י שאין תשובה נצחת על זה טוב לסמוך לו קצת תשובה ואומר אם היו צדקת האדם ורשעתו תלוים בגזרת הבורא ית' היתה לנו השאלה קשה עד מאד, ועכשיו שהבורא הסיר זאת הממשלה מידו ומסרה בידי האדם עצמו אין ידיעתו גזרה, אבל היא כידיעת האצטגנינים שיודעים מכח אחר מה יהיה דרכיו של זה, והדבר ידוע שכל מקרה האדם מסרם הבורא בכל המזלות אלא שנתן בו השכל להיות מחזיקו לצאת מתחת המזל והוא הכח הנתון באדם", עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:20.1 | אמנם |  |
| 5:20.2 | היה |  |
| 5:20.3 | האר"י |  |
| 5:20.4 | לפעמים |  |
| 5:20.5 | אומר |  |
| 5:20.6 | במקומו |  |
| 5:20.7 | אדון |  |
| 5:20.8 | עולם |  |
| 5:20.9 | אשר |  |
| 5:20.10 | מלך |  |
| 5:20.11 | מפני |  |
| 5:20.12 | שלא |  |
| 5:20.13 | נזכר |  |
| 5:20.14 | בו |  |
| 5:20.15 | שום |  |
| 5:20.16 | שם |  |
| 5:20.17 | משמות |  |
| 5:20.18 | הקודש | . |
| 5:20.19 | ושייך |  |
| 5:20.20 | כל |  |
| 5:20.21 | אותו |  |
| 5:20.22 | השבח |  |
| 5:20.23 | באלוה |  |
| 5:20.24 | הראשון |  |
| 5:20.25 | הנקרא |  |
| 5:20.26 | אין |  |
| 5:20.27 | סוף |  |
| 5:20.28 | אשר |  |
| 5:20.29 | לא |  |
| 5:20.30 | יעבדוהו |  |
| 5:20.31 | משום |  |
| 5:20.32 | דלית |  |
| 5:20.33 | ליה |  |
| 5:20.34 | שם |  |
| 5:20.35 | ידיע |  |
| 5:20.36 | ולא |  |
| 5:20.37 | שום |  |
| 5:20.38 | כינוי | . |
| 5:20.39 | וגם |  |
| 5:20.40 | נזכר |  |
| 5:20.41 | באדון |  |
| 5:20.42 | עולם |  |
| 5:20.43 | והוא |  |
| 5:20.44 | יהיה |  |
| 5:20.45 | בתפארה | , |
| 5:20.46 | שלפי |  |
| 5:20.47 | דעתם |  |
| 5:20.48 | קצתו |  |
| 5:20.49 | מתלבש |  |
| 5:20.50 | בתפארת | , |
| 5:20.51 | לכן |  |
| 5:20.52 | לא |  |
| 5:20.53 | נמנע |  |
| 5:20.54 | האר"י |  |
| 5:20.55 | מלאמרו | . |
| 5:20.56 | אבל |  |
| 5:20.57 | ביגדל |  |
| 5:20.58 | שנזכר |  |
| 5:20.59 | בו |  |
| 5:20.60 | שם |  |
| 5:20.61 | אלהים |  |
| 5:20.62 | חי |  |
| 5:20.63 | שהם |  |
| 5:20.64 | מיחסים |  |
| 5:20.65 | אותו |  |
| 5:20.66 | לפרצופי |  |
| 5:20.67 | אבא |  |
| 5:20.68 | ואמא |  |
| 5:20.69 | וגם |  |
| 5:20.70 | לזעיר |  |
| 5:20.71 | אנפין |  |
| 5:20.72 | אשר |  |
| 5:20.73 | אותם |  |
| 5:20.74 | יעבודו |  |
| 5:20.75 | שיש |  |
| 5:20.76 | להם |  |
| 5:20.77 | פרצוף |  |
| 5:20.78 | שלם |  |
| 5:20.79 | בדמות |  |
| 5:20.80 | אדם |  |
| 5:20.81 | ובצלמו | . |
| 5:20.82 | ואינם |  |
| 5:20.83 | ראשונים |  |
| 5:20.84 | שכבר |  |
| 5:20.85 | קדמו |  |
| 5:20.86 | להם |  |
| 5:20.87 | סיבות | , |
| 5:20.88 | ואינם |  |
| 5:20.89 | נעלמים | , |
| 5:20.90 | לכן |  |
| 5:20.91 | הקפיד |  |
| 5:20.92 | מאד |  |
| 5:20.93 | על |  |
| 5:20.94 | ארבעה |  |
| 5:20.95 | בתים |  |
| 5:20.96 | הראשונים |  |
| 5:20.97 | שהם |  |
| 5:20.98 | שקר |  |
| 5:20.99 | לפי |  |
| 5:20.100 | דעתם | . |
| 5:20.101 | הלא |  |
| 5:20.102 | אזן |  |
| 5:20.103 | מלים |  |
| 5:20.104 | תבחן |  |
| 5:20.105 | כמה |  |
| 5:20.106 | גדלה |  |
| 5:20.107 | שגיאת |  |
| 5:20.108 | רבותינו |  |
| 5:20.109 | בעלי |  |
| 5:20.110 | הקבלה |  |
| 5:20.111 | החדשה |  |
| 5:20.112 | לעקור |  |
| 5:20.113 | ארבעה |  |
| 5:20.114 | שרשים |  |
| 5:20.115 | מיסודי |  |
| 5:20.116 | תורתינו |  |
| 5:20.117 | הקדושה | , |
| 5:20.118 | ובהדיא |  |
| 5:20.119 | אמר |  |
| 5:20.120 | הרב |  |
| 5:20.121 | יושר |  |
| 5:20.122 | לבב |  |
| 5:20.123 | בדף |  |
| 5:20.124 | ג | ' |
| 5:20.125 | ע"ב | : |
| 5:20.126 | דע |  |
| 5:20.127 | את |  |
| 5:20.128 | אלהי |  |
| 5:20.129 | אביך |  |
| 5:20.130 | כולל |  |
| 5:20.131 | החמשה |  |
| 5:20.132 | פרצופים |  |
| 5:20.133 | וכו | ' |
| 5:20.134 | ועבדהו |  |
| 5:20.135 | לזעיר |  |
| 5:20.136 | אנפין |  |
| 5:20.137 | אע"פ |  |
| 5:20.138 | שהוא |  |
| 5:20.139 | נברא |  |
| 5:20.140 | כי |  |
| 5:20.141 | לנשמתו |  |
| 5:20.142 | אתה |  |
| 5:20.143 | עובד |  |
| 5:20.144 | בזה |  |
| 5:20.145 | כי |  |
| 5:20.146 | בלעדה |  |
| 5:20.147 | אין |  |
| 5:20.148 | כל |  |
| 5:20.149 | נמצא | , |
| 5:20.150 | משום |  |
| 5:20.151 | הכי |  |
| 5:20.152 | לא |  |
| 5:20.153 | אמר |  |
| 5:20.154 | דע |  |
| 5:20.155 | אלהי |  |
| 5:20.156 | אביך |  |
| 5:20.157 | אלא |  |
| 5:20.158 | את |  |
| 5:20.159 | אלהי |  |
| 5:20.160 | אביך | " |
| 5:20.161 | וכו | '. |
| 5:20.162 | ודברי |  |
| 5:20.163 | יושר |  |
| 5:20.164 | לבב |  |
| 5:20.165 | שאמר |  |
| 5:20.166 | שהוא |  |
| 5:20.167 | נברא |  |
| 5:20.168 | מבוארים |  |
| 5:20.169 | באדרא |  |
| 5:20.170 | רבה |  |
| 5:20.171 | דף |  |
| 5:20.172 | קל"ח |  |
| 5:20.173 | ע"ב |  |
| 5:20.174 | וז"ל | : |
| 5:20.175 | ה | ' |
| 5:20.176 | פעלך |  |
| 5:20.177 | דא |  |
| 5:20.178 | זעיר |  |
| 5:20.179 | אנפין |  |
| 5:20.180 | וכו | ' |
| 5:20.181 | חייהו |  |
| 5:20.182 | למאן | , |
| 5:20.183 | חייהו |  |
| 5:20.184 | לזעיר |  |
| 5:20.185 | אנפין |  |
| 5:20.186 | דכל |  |
| 5:20.187 | נהירו |  |
| 5:20.188 | דיליה |  |
| 5:20.189 | מאינון |  |
| 5:20.190 | שנים |  |
| 5:20.191 | קדמוניות | ", |
| 5:20.192 | וכו | ' |
| 5:20.193 | ועוד |  |
| 5:20.194 | הרבה |  |
| 5:20.195 | ראיות |  |
| 5:20.196 | בלשונות |  |
| 5:20.197 | הזהר |  |
| 5:20.198 | שהוא |  |
| 5:20.199 | נברא | , |
| 5:20.200 | ועם |  |
| 5:20.201 | כל |  |
| 5:20.202 | זה |  |
| 5:20.203 | מצוה |  |
| 5:20.204 | לעבוד |  |
| 5:20.205 | אותו | . |
| 5:20.206 | ואילו |  |
| 5:20.207 | שמו |  |
| 5:20.208 | אל |  |
| 5:20.209 | לבם |  |
| 5:20.210 | דברי |  |
| 5:20.211 | רז"ל |  |
| 5:20.212 | בבראשית |  |
| 5:20.213 | רבה |  |
| 5:20.214 | ובילקוט |  |
| 5:20.215 | שהבאתי |  |
| 5:20.216 | לעיל |  |
| 5:20.217 | סי | ' |
| 5:20.218 | ל"א |  |
| 5:20.219 | כל |  |
| 5:20.220 | שיש |  |
| 5:20.221 | לו |  |
| 5:20.222 | תולדות |  |
| 5:20.223 | וכו | ' |
| 5:20.224 | לא |  |
| 5:20.225 | עבדו |  |
| 5:20.226 | לנברא | . |
| 5:20.227 | גם |  |
| 5:20.228 | הטילו |  |
| 5:20.229 | חסרון |  |
| 5:20.230 | בידיעת |  |
| 5:20.231 | השי"ת |  |
| 5:20.232 | בעתיד |  |
| 5:20.233 | כמו |  |
| 5:20.234 | שכתבתי |  |
| 5:20.235 | לעיל |  |
| 5:20.236 | בשם |  |
| 5:20.237 | האר"י |  |
| 5:20.238 | בספר |  |
| 5:20.239 | ארבע |  |
| 5:20.240 | מאות |  |
| 5:20.241 | שקל |  |
| 5:20.242 | כסף | , |
| 5:20.243 | כי |  |
| 5:20.244 | עשו |  |
| 5:20.245 | ידיעת |  |
| 5:20.246 | השי"ת |  |
| 5:20.247 | בעתיד |  |
| 5:20.248 | ידיעת |  |
| 5:20.249 | ספק |  |
| 5:20.250 | כידיעת |  |
| 5:20.251 | האצטגנים |  |
| 5:20.252 | וכן |  |
| 5:20.253 | כתב |  |
| 5:20.254 | הראב"ד |  |
| 5:20.255 | בהשגות |  |
| 5:20.256 | פ"ז |  |
| 5:20.257 | מהלכות |  |
| 5:20.258 | תשובה | . |
| 5:20.259 | וז"ל | : |
| 5:20.260 | אעפ"י |  |
| 5:20.261 | שאין |  |
| 5:20.262 | תשובה |  |
| 5:20.263 | נצחת |  |
| 5:20.264 | על |  |
| 5:20.265 | זה |  |
| 5:20.266 | טוב |  |
| 5:20.267 | לסמוך |  |
| 5:20.268 | לו |  |
| 5:20.269 | קצת |  |
| 5:20.270 | תשובה |  |
| 5:20.271 | ואומר |  |
| 5:20.272 | אם |  |
| 5:20.273 | היו |  |
| 5:20.274 | צדקת |  |
| 5:20.275 | האדם |  |
| 5:20.276 | ורשעתו |  |
| 5:20.277 | תלוים |  |
| 5:20.278 | בגזרת |  |
| 5:20.279 | הבורא |  |
| 5:20.280 | ית | ' |
| 5:20.281 | היתה |  |
| 5:20.282 | לנו |  |
| 5:20.283 | השאלה |  |
| 5:20.284 | קשה |  |
| 5:20.285 | עד |  |
| 5:20.286 | מאד | , |
| 5:20.287 | ועכשיו |  |
| 5:20.288 | שהבורא |  |
| 5:20.289 | הסיר |  |
| 5:20.290 | זאת |  |
| 5:20.291 | הממשלה |  |
| 5:20.292 | מידו |  |
| 5:20.293 | ומסרה |  |
| 5:20.294 | בידי |  |
| 5:20.295 | האדם |  |
| 5:20.296 | עצמו |  |
| 5:20.297 | אין |  |
| 5:20.298 | ידיעתו |  |
| 5:20.299 | גזרה | , |
| 5:20.300 | אבל |  |
| 5:20.301 | היא |  |
| 5:20.302 | כידיעת |  |
| 5:20.303 | האצטגנינים |  |
| 5:20.304 | שיודעים |  |
| 5:20.305 | מכח |  |
| 5:20.306 | אחר |  |
| 5:20.307 | מה |  |
| 5:20.308 | יהיה |  |
| 5:20.309 | דרכיו |  |
| 5:20.310 | של |  |
| 5:20.311 | זה | , |
| 5:20.312 | והדבר |  |
| 5:20.313 | ידוע |  |
| 5:20.314 | שכל |  |
| 5:20.315 | מקרה |  |
| 5:20.316 | האדם |  |
| 5:20.317 | מסרם |  |
| 5:20.318 | הבורא |  |
| 5:20.319 | בכל |  |
| 5:20.320 | המזלות |  |
| 5:20.321 | אלא |  |
| 5:20.322 | שנתן |  |
| 5:20.323 | בו |  |
| 5:20.324 | השכל |  |
| 5:20.325 | להיות |  |
| 5:20.326 | מחזיקו |  |
| 5:20.327 | לצאת |  |
| 5:20.328 | מתחת |  |
| 5:20.329 | המזל |  |
| 5:20.330 | והוא |  |
| 5:20.331 | הכח |  |
| 5:20.332 | הנתון |  |
| 5:20.333 | באדם | ", |
| 5:20.334 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

