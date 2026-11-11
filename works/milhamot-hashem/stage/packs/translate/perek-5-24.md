# Translation pack: milhamot-hashem perek-5-24 (Milchamot ha-Shem 5:33–5:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-24.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-24`.

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
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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


## Milchamot ha-Shem 5:33

Hebrew (Wikisource): נט) ותו גרסינן במנחות פרק הרי עלי עשרון (מנחות דף ק"י) (דף קי.) אמר ר"א בר יצחק ואמרי לה אמר רב יהודה אמר רב מצור ועד קרתגני מכירין את ישראל ואת אביהם שבשמים, מצור כלפי מערב. ומקרתגני כלפי מזרח אין מכירין לא את ישראל ולא את אביהם שבשמים. אתיביה רב שימי רב חייא לרב. כי ממזרח שמש ועד מבואו גדול שמי בגוים ובכל מקום וכו' אמר ליה שימי את? דקרו ליה (לאלהי ישראל) אלהא דאלהיא ע"כ ופי' רש"י מכירים את ישראל מכירים בכבודן הראשון ואין מכירים את עולם ואת אביהם שבשמים שמאמינים בהקב"ה ואינם עובדים ע"ז. וכתב הר' מהרש"א נקט מתחלה ישראל דעל ידי ישראל למטה נתפרסם אלהותו ית' ויכלתו בעולם התחתון ושהוא אלהי ישראל ומשגיח על הארץ דהיינו אביהם שבשמים אבל מצור כלפי מערב כיון שאין מכירים את ישראל. אין מכירים את אביהם שבשמים שהוא משגיח בעולם השפל. אבל הוא אלהא דאלהא כדמסיק שהוא אלהי האלהים. ולאותן אלוהות שתחתיו נתן כח וממשלה למטה ואינו משגיח בעולם השפל התחתון עכ"ל. הט לבך קורא נעים לתושיה ולתבונה בדברי רז"ל בעלי התלמוד ותפקחנה עיניך ונהרת ונוכחת לדעת כי דעת המקובלים שאמרו שהאין סוף שהוא לעילא לעילא לית מחשבה תפיסה ביה ואין להתפלל ולקרא אליו רק לזעיר אנפין דשולטנו יהיב ליה על כלא. ושהקורא ומתפלל לאין סוף שהוא אלהים עליון לא יענה! היא בעצמה אמונת אוה"ע מצור כלפי מערב ומקרתגני כלפי מזרח האומרים שאבינו שבשמים המגביהי לשבת איננו משגיח בתחתונים. אבל הוא אלהי האלהים שהם חמשה פרצופים ושגם פרצופים העליונים עצמן נתנו הממשלה והכח לאלוה קצר אפים (זעיר אנפין) דאעטרו ליה ויהבו ליה שולטנו על כלא (זהר בלק דף קצ"א) כי רק לזעיר אנפין נתכנו עלילות. אבל האלהים העליונים לא נתכנו להם עלילות (זהר האזינו דף רצ"ב) כמו שנת' לעיל מסי' כ"ט והלאה. הוא שהנביא ישעיה מוכיח את ישראל במאמרו ית' לאמר איזה ספר וכו' הקצור קצרה ידי מפדות ואם אין בי כח להציל וכו' עד שאצטרך להעזר בפרצופים שהם למטה ממני. או שאמנה אותם במקומי? כי הוא ית' הסבה הראשונה לא צוה לעבוד לשום נברא כלל, ואינו כח בגוף כדי לעבוד לנשמתו כמו שחשבו המקובלים החדשים.

| id | word | punctuation after |
|---|---|---|
| 5:33.1 | נט | ) |
| 5:33.2 | ותו |  |
| 5:33.3 | גרסינן |  |
| 5:33.4 | במנחות |  |
| 5:33.5 | פרק |  |
| 5:33.6 | הרי |  |
| 5:33.7 | עלי |  |
| 5:33.8 | עשרון |  |
| 5:33.9 | מנחות |  |
| 5:33.10 | דף |  |
| 5:33.11 | ק"י | ) |
| 5:33.12 | דף |  |
| 5:33.13 | קי | .) |
| 5:33.14 | אמר |  |
| 5:33.15 | ר"א |  |
| 5:33.16 | בר |  |
| 5:33.17 | יצחק |  |
| 5:33.18 | ואמרי |  |
| 5:33.19 | לה |  |
| 5:33.20 | אמר |  |
| 5:33.21 | רב |  |
| 5:33.22 | יהודה |  |
| 5:33.23 | אמר |  |
| 5:33.24 | רב |  |
| 5:33.25 | מצור |  |
| 5:33.26 | ועד |  |
| 5:33.27 | קרתגני |  |
| 5:33.28 | מכירין |  |
| 5:33.29 | את |  |
| 5:33.30 | ישראל |  |
| 5:33.31 | ואת |  |
| 5:33.32 | אביהם |  |
| 5:33.33 | שבשמים | , |
| 5:33.34 | מצור |  |
| 5:33.35 | כלפי |  |
| 5:33.36 | מערב | . |
| 5:33.37 | ומקרתגני |  |
| 5:33.38 | כלפי |  |
| 5:33.39 | מזרח |  |
| 5:33.40 | אין |  |
| 5:33.41 | מכירין |  |
| 5:33.42 | לא |  |
| 5:33.43 | את |  |
| 5:33.44 | ישראל |  |
| 5:33.45 | ולא |  |
| 5:33.46 | את |  |
| 5:33.47 | אביהם |  |
| 5:33.48 | שבשמים | . |
| 5:33.49 | אתיביה |  |
| 5:33.50 | רב |  |
| 5:33.51 | שימי |  |
| 5:33.52 | רב |  |
| 5:33.53 | חייא |  |
| 5:33.54 | לרב | . |
| 5:33.55 | כי |  |
| 5:33.56 | ממזרח |  |
| 5:33.57 | שמש |  |
| 5:33.58 | ועד |  |
| 5:33.59 | מבואו |  |
| 5:33.60 | גדול |  |
| 5:33.61 | שמי |  |
| 5:33.62 | בגוים |  |
| 5:33.63 | ובכל |  |
| 5:33.64 | מקום |  |
| 5:33.65 | וכו | ' |
| 5:33.66 | אמר |  |
| 5:33.67 | ליה |  |
| 5:33.68 | שימי |  |
| 5:33.69 | את | ? |
| 5:33.70 | דקרו |  |
| 5:33.71 | ליה |  |
| 5:33.72 | לאלהי |  |
| 5:33.73 | ישראל | ) |
| 5:33.74 | אלהא |  |
| 5:33.75 | דאלהיא |  |
| 5:33.76 | ע"כ |  |
| 5:33.77 | ופי | ' |
| 5:33.78 | רש"י |  |
| 5:33.79 | מכירים |  |
| 5:33.80 | את |  |
| 5:33.81 | ישראל |  |
| 5:33.82 | מכירים |  |
| 5:33.83 | בכבודן |  |
| 5:33.84 | הראשון |  |
| 5:33.85 | ואין |  |
| 5:33.86 | מכירים |  |
| 5:33.87 | את |  |
| 5:33.88 | עולם |  |
| 5:33.89 | ואת |  |
| 5:33.90 | אביהם |  |
| 5:33.91 | שבשמים |  |
| 5:33.92 | שמאמינים |  |
| 5:33.93 | בהקב"ה |  |
| 5:33.94 | ואינם |  |
| 5:33.95 | עובדים |  |
| 5:33.96 | ע"ז | . |
| 5:33.97 | וכתב |  |
| 5:33.98 | הר | ' |
| 5:33.99 | מהרש"א |  |
| 5:33.100 | נקט |  |
| 5:33.101 | מתחלה |  |
| 5:33.102 | ישראל |  |
| 5:33.103 | דעל |  |
| 5:33.104 | ידי |  |
| 5:33.105 | ישראל |  |
| 5:33.106 | למטה |  |
| 5:33.107 | נתפרסם |  |
| 5:33.108 | אלהותו |  |
| 5:33.109 | ית | ' |
| 5:33.110 | ויכלתו |  |
| 5:33.111 | בעולם |  |
| 5:33.112 | התחתון |  |
| 5:33.113 | ושהוא |  |
| 5:33.114 | אלהי |  |
| 5:33.115 | ישראל |  |
| 5:33.116 | ומשגיח |  |
| 5:33.117 | על |  |
| 5:33.118 | הארץ |  |
| 5:33.119 | דהיינו |  |
| 5:33.120 | אביהם |  |
| 5:33.121 | שבשמים |  |
| 5:33.122 | אבל |  |
| 5:33.123 | מצור |  |
| 5:33.124 | כלפי |  |
| 5:33.125 | מערב |  |
| 5:33.126 | כיון |  |
| 5:33.127 | שאין |  |
| 5:33.128 | מכירים |  |
| 5:33.129 | את |  |
| 5:33.130 | ישראל | . |
| 5:33.131 | אין |  |
| 5:33.132 | מכירים |  |
| 5:33.133 | את |  |
| 5:33.134 | אביהם |  |
| 5:33.135 | שבשמים |  |
| 5:33.136 | שהוא |  |
| 5:33.137 | משגיח |  |
| 5:33.138 | בעולם |  |
| 5:33.139 | השפל | . |
| 5:33.140 | אבל |  |
| 5:33.141 | הוא |  |
| 5:33.142 | אלהא |  |
| 5:33.143 | דאלהא |  |
| 5:33.144 | כדמסיק |  |
| 5:33.145 | שהוא |  |
| 5:33.146 | אלהי |  |
| 5:33.147 | האלהים | . |
| 5:33.148 | ולאותן |  |
| 5:33.149 | אלוהות |  |
| 5:33.150 | שתחתיו |  |
| 5:33.151 | נתן |  |
| 5:33.152 | כח |  |
| 5:33.153 | וממשלה |  |
| 5:33.154 | למטה |  |
| 5:33.155 | ואינו |  |
| 5:33.156 | משגיח |  |
| 5:33.157 | בעולם |  |
| 5:33.158 | השפל |  |
| 5:33.159 | התחתון |  |
| 5:33.160 | עכ"ל | . |
| 5:33.161 | הט |  |
| 5:33.162 | לבך |  |
| 5:33.163 | קורא |  |
| 5:33.164 | נעים |  |
| 5:33.165 | לתושיה |  |
| 5:33.166 | ולתבונה |  |
| 5:33.167 | בדברי |  |
| 5:33.168 | רז"ל |  |
| 5:33.169 | בעלי |  |
| 5:33.170 | התלמוד |  |
| 5:33.171 | ותפקחנה |  |
| 5:33.172 | עיניך |  |
| 5:33.173 | ונהרת |  |
| 5:33.174 | ונוכחת |  |
| 5:33.175 | לדעת |  |
| 5:33.176 | כי |  |
| 5:33.177 | דעת |  |
| 5:33.178 | המקובלים |  |
| 5:33.179 | שאמרו |  |
| 5:33.180 | שהאין |  |
| 5:33.181 | סוף |  |
| 5:33.182 | שהוא |  |
| 5:33.183 | לעילא |  |
| 5:33.184 | לעילא |  |
| 5:33.185 | לית |  |
| 5:33.186 | מחשבה |  |
| 5:33.187 | תפיסה |  |
| 5:33.188 | ביה |  |
| 5:33.189 | ואין |  |
| 5:33.190 | להתפלל |  |
| 5:33.191 | ולקרא |  |
| 5:33.192 | אליו |  |
| 5:33.193 | רק |  |
| 5:33.194 | לזעיר |  |
| 5:33.195 | אנפין |  |
| 5:33.196 | דשולטנו |  |
| 5:33.197 | יהיב |  |
| 5:33.198 | ליה |  |
| 5:33.199 | על |  |
| 5:33.200 | כלא | . |
| 5:33.201 | ושהקורא |  |
| 5:33.202 | ומתפלל |  |
| 5:33.203 | לאין |  |
| 5:33.204 | סוף |  |
| 5:33.205 | שהוא |  |
| 5:33.206 | אלהים |  |
| 5:33.207 | עליון |  |
| 5:33.208 | לא |  |
| 5:33.209 | יענה | ! |
| 5:33.210 | היא |  |
| 5:33.211 | בעצמה |  |
| 5:33.212 | אמונת |  |
| 5:33.213 | אוה"ע |  |
| 5:33.214 | מצור |  |
| 5:33.215 | כלפי |  |
| 5:33.216 | מערב |  |
| 5:33.217 | ומקרתגני |  |
| 5:33.218 | כלפי |  |
| 5:33.219 | מזרח |  |
| 5:33.220 | האומרים |  |
| 5:33.221 | שאבינו |  |
| 5:33.222 | שבשמים |  |
| 5:33.223 | המגביהי |  |
| 5:33.224 | לשבת |  |
| 5:33.225 | איננו |  |
| 5:33.226 | משגיח |  |
| 5:33.227 | בתחתונים | . |
| 5:33.228 | אבל |  |
| 5:33.229 | הוא |  |
| 5:33.230 | אלהי |  |
| 5:33.231 | האלהים |  |
| 5:33.232 | שהם |  |
| 5:33.233 | חמשה |  |
| 5:33.234 | פרצופים |  |
| 5:33.235 | ושגם |  |
| 5:33.236 | פרצופים |  |
| 5:33.237 | העליונים |  |
| 5:33.238 | עצמן |  |
| 5:33.239 | נתנו |  |
| 5:33.240 | הממשלה |  |
| 5:33.241 | והכח |  |
| 5:33.242 | לאלוה |  |
| 5:33.243 | קצר |  |
| 5:33.244 | אפים |  |
| 5:33.245 | זעיר |  |
| 5:33.246 | אנפין | ) |
| 5:33.247 | דאעטרו |  |
| 5:33.248 | ליה |  |
| 5:33.249 | ויהבו |  |
| 5:33.250 | ליה |  |
| 5:33.251 | שולטנו |  |
| 5:33.252 | על |  |
| 5:33.253 | כלא |  |
| 5:33.254 | זהר |  |
| 5:33.255 | בלק |  |
| 5:33.256 | דף |  |
| 5:33.257 | קצ"א | ) |
| 5:33.258 | כי |  |
| 5:33.259 | רק |  |
| 5:33.260 | לזעיר |  |
| 5:33.261 | אנפין |  |
| 5:33.262 | נתכנו |  |
| 5:33.263 | עלילות | . |
| 5:33.264 | אבל |  |
| 5:33.265 | האלהים |  |
| 5:33.266 | העליונים |  |
| 5:33.267 | לא |  |
| 5:33.268 | נתכנו |  |
| 5:33.269 | להם |  |
| 5:33.270 | עלילות |  |
| 5:33.271 | זהר |  |
| 5:33.272 | האזינו |  |
| 5:33.273 | דף |  |
| 5:33.274 | רצ"ב | ) |
| 5:33.275 | כמו |  |
| 5:33.276 | שנת | ' |
| 5:33.277 | לעיל |  |
| 5:33.278 | מסי | ' |
| 5:33.279 | כ"ט |  |
| 5:33.280 | והלאה | . |
| 5:33.281 | הוא |  |
| 5:33.282 | שהנביא |  |
| 5:33.283 | ישעיה |  |
| 5:33.284 | מוכיח |  |
| 5:33.285 | את |  |
| 5:33.286 | ישראל |  |
| 5:33.287 | במאמרו |  |
| 5:33.288 | ית | ' |
| 5:33.289 | לאמר |  |
| 5:33.290 | איזה |  |
| 5:33.291 | ספר |  |
| 5:33.292 | וכו | ' |
| 5:33.293 | הקצור |  |
| 5:33.294 | קצרה |  |
| 5:33.295 | ידי |  |
| 5:33.296 | מפדות |  |
| 5:33.297 | ואם |  |
| 5:33.298 | אין |  |
| 5:33.299 | בי |  |
| 5:33.300 | כח |  |
| 5:33.301 | להציל |  |
| 5:33.302 | וכו | ' |
| 5:33.303 | עד |  |
| 5:33.304 | שאצטרך |  |
| 5:33.305 | להעזר |  |
| 5:33.306 | בפרצופים |  |
| 5:33.307 | שהם |  |
| 5:33.308 | למטה |  |
| 5:33.309 | ממני | . |
| 5:33.310 | או |  |
| 5:33.311 | שאמנה |  |
| 5:33.312 | אותם |  |
| 5:33.313 | במקומי | ? |
| 5:33.314 | כי |  |
| 5:33.315 | הוא |  |
| 5:33.316 | ית | ' |
| 5:33.317 | הסבה |  |
| 5:33.318 | הראשונה |  |
| 5:33.319 | לא |  |
| 5:33.320 | צוה |  |
| 5:33.321 | לעבוד |  |
| 5:33.322 | לשום |  |
| 5:33.323 | נברא |  |
| 5:33.324 | כלל | , |
| 5:33.325 | ואינו |  |
| 5:33.326 | כח |  |
| 5:33.327 | בגוף |  |
| 5:33.328 | כדי |  |
| 5:33.329 | לעבוד |  |
| 5:33.330 | לנשמתו |  |
| 5:33.331 | כמו |  |
| 5:33.332 | שחשבו |  |
| 5:33.333 | המקובלים |  |
| 5:33.334 | החדשים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

