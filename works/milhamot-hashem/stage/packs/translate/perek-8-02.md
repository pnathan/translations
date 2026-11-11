# Translation pack: milhamot-hashem perek-8-02 (Milchamot ha-Shem 8:3–8:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-02`.

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
- דברי תימה: "astonishing words" (words that astonish).
- מעמד: "assembly" or "occasion" by context; "the Standing at Sinai" where the
  author names that event.
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 8:3

Hebrew (Wikisource): קא) ולתקן הכפירה והמינות הגלויה באמונה זו אמרו, כי עבודתינו שאנו עובדים לו אע"פ שהוא נברא כמו שכתב הרב יושר לבב כי לנשמתו אנו עובדים. ואע"פ שאנו מאמינים כי אלהינו אינו גוף ולא כח בגוף! והנה ידוע מדבריהם כי הנשמה שלכל נברא יש לה דמות וצורת הגוף בכל חלקי איבריו נשמת הזכר זכר ונשמת הנקבה נקבה, כמו שכתב עז לאלהים בבית קדש הקדשים פ"ג דף ט' ע"א שבהכרח גמור יש לנשמה פרצוף ברמ"ח איברים ושס"ה גידים בצורת אדם. ומתפשטת בצורת איבריה, ואם תמצא אבר מיותר אינה מתפשטת בו, וכן ג"כ נשמה לנשמה צריך שיהיה לה צורת הפרצוף וכו' עיין שם. כי על כן צריך לכפול לדידהו ה' אלהיכם אמת כדי להשלים רמ"ח איברי הנשמה שבזכר, ואם לא יכפול תשאר נשמתו חסרה (וזהו למרות דברי רז"ל שאמרו האומר מלה מלה בקריאת שמע וכופלה משתקין אותו משום דמחזי כשתי רשויות, ובכפילת ה' אלהיכם אמת כל שכן דמיחזי כשתי רשויות). ובמעמד הר סיני לפי דבריהם קדמה אשתו של זעיר אנפין שהיא המלכות לרדת על הר סיני. ואחר כך הלך הזעיר אנפין אחריה כמנוח אחרי אשתו. ונמצא הוא וזוגתו כולם ירדו להר סיני. שהרי הוא ירד להתחבר עמה כמו שכתב נחלת יוסף. וגם אבא ואימא גם הם נמצאו שם בהר סיני. שהרי אימא היא שאמרה אשר הוציאתיך. כי היא הוציאתם ממצרים, והמלכות היא חכמה במצרים בסיועת הבינה שהיא אימא, והתורה יצאה מחכמה שהוא אבא, ונמצא לפי זה שבמעמד הר סיני ירדו אבא ואימא וזעיר נוקביה על הר סיני. וזעיר אנפין הוא שאמר אנכי ה' אלהיך. ואם כן הוא יקשה מאד מאמר רז"ל בשמות רבה וז"ל אנכי ה' אלהיך. מלך בשר ודם שמלך יש לו אב או אח, או בן. אמר הקב"ה אני איני כן אני ראשון שאין לי אב. ואני אחרון שאין לי בן. ומבלעדי אין אלהים שאין לי אח ע"כ. הלא לפי הקבלה החדשה אבא ואימא ומלכות כולם נגלו עם קצר אפים (זעיר אנפין) בסיני והתיצבו שם עמו, והוא המדבר אנכי ה' אלהיך. איך אמר שאין לי אב ואם הרי הם נצבים עמו? ואיך אמר שאין לי אח והרי זוגתו קדמה לו לירד על הר סיני. והוא ירד אחריה להתחבר לה? השקר הוא מדבר. ולמה זה ישקר עלינו והם עומדים עמו? האם להתל בנו נגלה עלינו? חלילה לה' אלהינו מן השקר. כי הוא ה' אלהים אמת. ותורת אמת הנחילנו, ובמצות תורתינו הקדושה הרחיקנו מאמונות זרות, וכוזבות שטעו באן האומות עובדי האלילים זולתינו.

| id | word | punctuation after |
|---|---|---|
| 8:3.1 | קא | ) |
| 8:3.2 | ולתקן |  |
| 8:3.3 | הכפירה |  |
| 8:3.4 | והמינות |  |
| 8:3.5 | הגלויה |  |
| 8:3.6 | באמונה |  |
| 8:3.7 | זו |  |
| 8:3.8 | אמרו | , |
| 8:3.9 | כי |  |
| 8:3.10 | עבודתינו |  |
| 8:3.11 | שאנו |  |
| 8:3.12 | עובדים |  |
| 8:3.13 | לו |  |
| 8:3.14 | אע"פ |  |
| 8:3.15 | שהוא |  |
| 8:3.16 | נברא |  |
| 8:3.17 | כמו |  |
| 8:3.18 | שכתב |  |
| 8:3.19 | הרב |  |
| 8:3.20 | יושר |  |
| 8:3.21 | לבב |  |
| 8:3.22 | כי |  |
| 8:3.23 | לנשמתו |  |
| 8:3.24 | אנו |  |
| 8:3.25 | עובדים | . |
| 8:3.26 | ואע"פ |  |
| 8:3.27 | שאנו |  |
| 8:3.28 | מאמינים |  |
| 8:3.29 | כי |  |
| 8:3.30 | אלהינו |  |
| 8:3.31 | אינו |  |
| 8:3.32 | גוף |  |
| 8:3.33 | ולא |  |
| 8:3.34 | כח |  |
| 8:3.35 | בגוף | ! |
| 8:3.36 | והנה |  |
| 8:3.37 | ידוע |  |
| 8:3.38 | מדבריהם |  |
| 8:3.39 | כי |  |
| 8:3.40 | הנשמה |  |
| 8:3.41 | שלכל |  |
| 8:3.42 | נברא |  |
| 8:3.43 | יש |  |
| 8:3.44 | לה |  |
| 8:3.45 | דמות |  |
| 8:3.46 | וצורת |  |
| 8:3.47 | הגוף |  |
| 8:3.48 | בכל |  |
| 8:3.49 | חלקי |  |
| 8:3.50 | איבריו |  |
| 8:3.51 | נשמת |  |
| 8:3.52 | הזכר |  |
| 8:3.53 | זכר |  |
| 8:3.54 | ונשמת |  |
| 8:3.55 | הנקבה |  |
| 8:3.56 | נקבה | , |
| 8:3.57 | כמו |  |
| 8:3.58 | שכתב |  |
| 8:3.59 | עז |  |
| 8:3.60 | לאלהים |  |
| 8:3.61 | בבית |  |
| 8:3.62 | קדש |  |
| 8:3.63 | הקדשים |  |
| 8:3.64 | פ"ג |  |
| 8:3.65 | דף |  |
| 8:3.66 | ט | ' |
| 8:3.67 | ע"א |  |
| 8:3.68 | שבהכרח |  |
| 8:3.69 | גמור |  |
| 8:3.70 | יש |  |
| 8:3.71 | לנשמה |  |
| 8:3.72 | פרצוף |  |
| 8:3.73 | ברמ"ח |  |
| 8:3.74 | איברים |  |
| 8:3.75 | ושס"ה |  |
| 8:3.76 | גידים |  |
| 8:3.77 | בצורת |  |
| 8:3.78 | אדם | . |
| 8:3.79 | ומתפשטת |  |
| 8:3.80 | בצורת |  |
| 8:3.81 | איבריה | , |
| 8:3.82 | ואם |  |
| 8:3.83 | תמצא |  |
| 8:3.84 | אבר |  |
| 8:3.85 | מיותר |  |
| 8:3.86 | אינה |  |
| 8:3.87 | מתפשטת |  |
| 8:3.88 | בו | , |
| 8:3.89 | וכן |  |
| 8:3.90 | ג"כ |  |
| 8:3.91 | נשמה |  |
| 8:3.92 | לנשמה |  |
| 8:3.93 | צריך |  |
| 8:3.94 | שיהיה |  |
| 8:3.95 | לה |  |
| 8:3.96 | צורת |  |
| 8:3.97 | הפרצוף |  |
| 8:3.98 | וכו | ' |
| 8:3.99 | עיין |  |
| 8:3.100 | שם | . |
| 8:3.101 | כי |  |
| 8:3.102 | על |  |
| 8:3.103 | כן |  |
| 8:3.104 | צריך |  |
| 8:3.105 | לכפול |  |
| 8:3.106 | לדידהו |  |
| 8:3.107 | ה | ' |
| 8:3.108 | אלהיכם |  |
| 8:3.109 | אמת |  |
| 8:3.110 | כדי |  |
| 8:3.111 | להשלים |  |
| 8:3.112 | רמ"ח |  |
| 8:3.113 | איברי |  |
| 8:3.114 | הנשמה |  |
| 8:3.115 | שבזכר | , |
| 8:3.116 | ואם |  |
| 8:3.117 | לא |  |
| 8:3.118 | יכפול |  |
| 8:3.119 | תשאר |  |
| 8:3.120 | נשמתו |  |
| 8:3.121 | חסרה |  |
| 8:3.122 | וזהו |  |
| 8:3.123 | למרות |  |
| 8:3.124 | דברי |  |
| 8:3.125 | רז"ל |  |
| 8:3.126 | שאמרו |  |
| 8:3.127 | האומר |  |
| 8:3.128 | מלה |  |
| 8:3.129 | מלה |  |
| 8:3.130 | בקריאת |  |
| 8:3.131 | שמע |  |
| 8:3.132 | וכופלה |  |
| 8:3.133 | משתקין |  |
| 8:3.134 | אותו |  |
| 8:3.135 | משום |  |
| 8:3.136 | דמחזי |  |
| 8:3.137 | כשתי |  |
| 8:3.138 | רשויות | , |
| 8:3.139 | ובכפילת |  |
| 8:3.140 | ה | ' |
| 8:3.141 | אלהיכם |  |
| 8:3.142 | אמת |  |
| 8:3.143 | כל |  |
| 8:3.144 | שכן |  |
| 8:3.145 | דמיחזי |  |
| 8:3.146 | כשתי |  |
| 8:3.147 | רשויות | ). |
| 8:3.148 | ובמעמד |  |
| 8:3.149 | הר |  |
| 8:3.150 | סיני |  |
| 8:3.151 | לפי |  |
| 8:3.152 | דבריהם |  |
| 8:3.153 | קדמה |  |
| 8:3.154 | אשתו |  |
| 8:3.155 | של |  |
| 8:3.156 | זעיר |  |
| 8:3.157 | אנפין |  |
| 8:3.158 | שהיא |  |
| 8:3.159 | המלכות |  |
| 8:3.160 | לרדת |  |
| 8:3.161 | על |  |
| 8:3.162 | הר |  |
| 8:3.163 | סיני | . |
| 8:3.164 | ואחר |  |
| 8:3.165 | כך |  |
| 8:3.166 | הלך |  |
| 8:3.167 | הזעיר |  |
| 8:3.168 | אנפין |  |
| 8:3.169 | אחריה |  |
| 8:3.170 | כמנוח |  |
| 8:3.171 | אחרי |  |
| 8:3.172 | אשתו | . |
| 8:3.173 | ונמצא |  |
| 8:3.174 | הוא |  |
| 8:3.175 | וזוגתו |  |
| 8:3.176 | כולם |  |
| 8:3.177 | ירדו |  |
| 8:3.178 | להר |  |
| 8:3.179 | סיני | . |
| 8:3.180 | שהרי |  |
| 8:3.181 | הוא |  |
| 8:3.182 | ירד |  |
| 8:3.183 | להתחבר |  |
| 8:3.184 | עמה |  |
| 8:3.185 | כמו |  |
| 8:3.186 | שכתב |  |
| 8:3.187 | נחלת |  |
| 8:3.188 | יוסף | . |
| 8:3.189 | וגם |  |
| 8:3.190 | אבא |  |
| 8:3.191 | ואימא |  |
| 8:3.192 | גם |  |
| 8:3.193 | הם |  |
| 8:3.194 | נמצאו |  |
| 8:3.195 | שם |  |
| 8:3.196 | בהר |  |
| 8:3.197 | סיני | . |
| 8:3.198 | שהרי |  |
| 8:3.199 | אימא |  |
| 8:3.200 | היא |  |
| 8:3.201 | שאמרה |  |
| 8:3.202 | אשר |  |
| 8:3.203 | הוציאתיך | . |
| 8:3.204 | כי |  |
| 8:3.205 | היא |  |
| 8:3.206 | הוציאתם |  |
| 8:3.207 | ממצרים | , |
| 8:3.208 | והמלכות |  |
| 8:3.209 | היא |  |
| 8:3.210 | חכמה |  |
| 8:3.211 | במצרים |  |
| 8:3.212 | בסיועת |  |
| 8:3.213 | הבינה |  |
| 8:3.214 | שהיא |  |
| 8:3.215 | אימא | , |
| 8:3.216 | והתורה |  |
| 8:3.217 | יצאה |  |
| 8:3.218 | מחכמה |  |
| 8:3.219 | שהוא |  |
| 8:3.220 | אבא | , |
| 8:3.221 | ונמצא |  |
| 8:3.222 | לפי |  |
| 8:3.223 | זה |  |
| 8:3.224 | שבמעמד |  |
| 8:3.225 | הר |  |
| 8:3.226 | סיני |  |
| 8:3.227 | ירדו |  |
| 8:3.228 | אבא |  |
| 8:3.229 | ואימא |  |
| 8:3.230 | וזעיר |  |
| 8:3.231 | נוקביה |  |
| 8:3.232 | על |  |
| 8:3.233 | הר |  |
| 8:3.234 | סיני | . |
| 8:3.235 | וזעיר |  |
| 8:3.236 | אנפין |  |
| 8:3.237 | הוא |  |
| 8:3.238 | שאמר |  |
| 8:3.239 | אנכי |  |
| 8:3.240 | ה | ' |
| 8:3.241 | אלהיך | . |
| 8:3.242 | ואם |  |
| 8:3.243 | כן |  |
| 8:3.244 | הוא |  |
| 8:3.245 | יקשה |  |
| 8:3.246 | מאד |  |
| 8:3.247 | מאמר |  |
| 8:3.248 | רז"ל |  |
| 8:3.249 | בשמות |  |
| 8:3.250 | רבה |  |
| 8:3.251 | וז"ל |  |
| 8:3.252 | אנכי |  |
| 8:3.253 | ה | ' |
| 8:3.254 | אלהיך | . |
| 8:3.255 | מלך |  |
| 8:3.256 | בשר |  |
| 8:3.257 | ודם |  |
| 8:3.258 | שמלך |  |
| 8:3.259 | יש |  |
| 8:3.260 | לו |  |
| 8:3.261 | אב |  |
| 8:3.262 | או |  |
| 8:3.263 | אח | , |
| 8:3.264 | או |  |
| 8:3.265 | בן | . |
| 8:3.266 | אמר |  |
| 8:3.267 | הקב"ה |  |
| 8:3.268 | אני |  |
| 8:3.269 | איני |  |
| 8:3.270 | כן |  |
| 8:3.271 | אני |  |
| 8:3.272 | ראשון |  |
| 8:3.273 | שאין |  |
| 8:3.274 | לי |  |
| 8:3.275 | אב | . |
| 8:3.276 | ואני |  |
| 8:3.277 | אחרון |  |
| 8:3.278 | שאין |  |
| 8:3.279 | לי |  |
| 8:3.280 | בן | . |
| 8:3.281 | ומבלעדי |  |
| 8:3.282 | אין |  |
| 8:3.283 | אלהים |  |
| 8:3.284 | שאין |  |
| 8:3.285 | לי |  |
| 8:3.286 | אח |  |
| 8:3.287 | ע"כ | . |
| 8:3.288 | הלא |  |
| 8:3.289 | לפי |  |
| 8:3.290 | הקבלה |  |
| 8:3.291 | החדשה |  |
| 8:3.292 | אבא |  |
| 8:3.293 | ואימא |  |
| 8:3.294 | ומלכות |  |
| 8:3.295 | כולם |  |
| 8:3.296 | נגלו |  |
| 8:3.297 | עם |  |
| 8:3.298 | קצר |  |
| 8:3.299 | אפים |  |
| 8:3.300 | זעיר |  |
| 8:3.301 | אנפין | ) |
| 8:3.302 | בסיני |  |
| 8:3.303 | והתיצבו |  |
| 8:3.304 | שם |  |
| 8:3.305 | עמו | , |
| 8:3.306 | והוא |  |
| 8:3.307 | המדבר |  |
| 8:3.308 | אנכי |  |
| 8:3.309 | ה | ' |
| 8:3.310 | אלהיך | . |
| 8:3.311 | איך |  |
| 8:3.312 | אמר |  |
| 8:3.313 | שאין |  |
| 8:3.314 | לי |  |
| 8:3.315 | אב |  |
| 8:3.316 | ואם |  |
| 8:3.317 | הרי |  |
| 8:3.318 | הם |  |
| 8:3.319 | נצבים |  |
| 8:3.320 | עמו | ? |
| 8:3.321 | ואיך |  |
| 8:3.322 | אמר |  |
| 8:3.323 | שאין |  |
| 8:3.324 | לי |  |
| 8:3.325 | אח |  |
| 8:3.326 | והרי |  |
| 8:3.327 | זוגתו |  |
| 8:3.328 | קדמה |  |
| 8:3.329 | לו |  |
| 8:3.330 | לירד |  |
| 8:3.331 | על |  |
| 8:3.332 | הר |  |
| 8:3.333 | סיני | . |
| 8:3.334 | והוא |  |
| 8:3.335 | ירד |  |
| 8:3.336 | אחריה |  |
| 8:3.337 | להתחבר |  |
| 8:3.338 | לה | ? |
| 8:3.339 | השקר |  |
| 8:3.340 | הוא |  |
| 8:3.341 | מדבר | . |
| 8:3.342 | ולמה |  |
| 8:3.343 | זה |  |
| 8:3.344 | ישקר |  |
| 8:3.345 | עלינו |  |
| 8:3.346 | והם |  |
| 8:3.347 | עומדים |  |
| 8:3.348 | עמו | ? |
| 8:3.349 | האם |  |
| 8:3.350 | להתל |  |
| 8:3.351 | בנו |  |
| 8:3.352 | נגלה |  |
| 8:3.353 | עלינו | ? |
| 8:3.354 | חלילה |  |
| 8:3.355 | לה | ' |
| 8:3.356 | אלהינו |  |
| 8:3.357 | מן |  |
| 8:3.358 | השקר | . |
| 8:3.359 | כי |  |
| 8:3.360 | הוא |  |
| 8:3.361 | ה | ' |
| 8:3.362 | אלהים |  |
| 8:3.363 | אמת | . |
| 8:3.364 | ותורת |  |
| 8:3.365 | אמת |  |
| 8:3.366 | הנחילנו | , |
| 8:3.367 | ובמצות |  |
| 8:3.368 | תורתינו |  |
| 8:3.369 | הקדושה |  |
| 8:3.370 | הרחיקנו |  |
| 8:3.371 | מאמונות |  |
| 8:3.372 | זרות | , |
| 8:3.373 | וכוזבות |  |
| 8:3.374 | שטעו |  |
| 8:3.375 | באן |  |
| 8:3.376 | האומות |  |
| 8:3.377 | עובדי |  |
| 8:3.378 | האלילים |  |
| 8:3.379 | זולתינו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

