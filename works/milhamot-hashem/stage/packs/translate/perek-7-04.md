# Translation pack: milhamot-hashem perek-7-04 (Milchamot ha-Shem 7:5–7:5)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-04`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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


## Milchamot ha-Shem 7:5

Hebrew (Wikisource): פב) ובמדרש רבה פרשת ואתחנן. שמע ישאל ה' אלהינו ה' אחד, זהו שאמר הכתוב מי לי בשמים ועמך לא חפצתי בארץ, רב אמר ב' רקיעים הם שמים ושמי השמים. ר' אלעזר אמר שבעה רקיעים הם שמים ושמי השמים. רקיע, שחקים, מעון, זבול, ערפל, וכולו פתח הקב"ה לישראל להודיען שאין אלוה אחר אלא הוא. אמרה כנסת ישראל לפני הקב"ה. מי לי בשמים חוץ מכבודך. כשם שאין לי בשמים אלא אתה, כך לא חפצתי אחר בארץ. אלא אני נכנסת בכל יום לבתי כנסיות ומעידה עליך שאין אלוה אחר אלא אתה ואומרת שמע ישראל ה' אלהינו ה' אחד. הנך רואה ידידי דברי רז"ל הנז' במדרש עומדים נגד הקבלה החדשה. שהרי בשמע ישראל שהוא קבלת עול מלכות שמים, המקובלים החדשים משתפים בו אבא ואימא וזעיר ונוקביה כמבואר בזהר ואתחנן. הובא לעיל סי' מ"ב, מה שלא עלה על דעת רז"ל שייעשה כן בישראל וגדולה מזו. אמרו בעלי הקבלה החדשה, שהמתפלל לאין סוף הסבה הראשונה אין תפלתו תפלה, אלא צריך להתפלל דוקא לז"א ליחד זוגתו עמו וגם לשתף במחשבה אבא ואימא ואריך אנפין ועתיק יומין כמו שכתבתי לעיל ב' כסא אליהו ויושר לבב. שהמיחד תפלתו לז"א בלי צירוף ושתוף פרצופים הנז' אינו נענה מהרה, וזה סוד מה תצעק אלי כדלעיל סי' ל'. עוד אמרו במדרש רבה פרשה הנז"ל. וז"ל שמע ישראל ה' אלהינו ה' אחד, רבנן אמרי אמר הקב"ה בני כל מה שבראתי בראתי זוגות, שמים וארץ זוגות, חמה ולבנה זוגות, אדם וחוה זוגות, העולם הזה והעה"ב זוגות, אבל כבודי אחד ומיוחד בעולם, מנין ממה שקרינו בענין שמע ישראל ה' אלהינו ה' אחד עכ"ל. גם דברי המדרש אלו עומדים נגד הקבלה החדשה שאמרה שזעיר ונוקביה הוא אלוהינו ואנחנו עמו וצאן מריעתו. ואותו נעבוד במצות אבא ואימא, והוא משתבח בעצמו ואומר ה' אמר אלי בני אתה וכו' כדלעיל סי' כ"ו. שהרי במדרש אמרו כי ה' אלהינו אחד ומיוחד מכל האחדים שבעולם כנז' לעיל, ואין לו עזר כנגדו לפי שהוא יכול ואינו צריך לעזר זולתו.

| id | word | punctuation after |
|---|---|---|
| 7:5.1 | פב | ) |
| 7:5.2 | ובמדרש |  |
| 7:5.3 | רבה |  |
| 7:5.4 | פרשת |  |
| 7:5.5 | ואתחנן | . |
| 7:5.6 | שמע |  |
| 7:5.7 | ישאל |  |
| 7:5.8 | ה | ' |
| 7:5.9 | אלהינו |  |
| 7:5.10 | ה | ' |
| 7:5.11 | אחד | , |
| 7:5.12 | זהו |  |
| 7:5.13 | שאמר |  |
| 7:5.14 | הכתוב |  |
| 7:5.15 | מי |  |
| 7:5.16 | לי |  |
| 7:5.17 | בשמים |  |
| 7:5.18 | ועמך |  |
| 7:5.19 | לא |  |
| 7:5.20 | חפצתי |  |
| 7:5.21 | בארץ | , |
| 7:5.22 | רב |  |
| 7:5.23 | אמר |  |
| 7:5.24 | ב | ' |
| 7:5.25 | רקיעים |  |
| 7:5.26 | הם |  |
| 7:5.27 | שמים |  |
| 7:5.28 | ושמי |  |
| 7:5.29 | השמים | . |
| 7:5.30 | ר | ' |
| 7:5.31 | אלעזר |  |
| 7:5.32 | אמר |  |
| 7:5.33 | שבעה |  |
| 7:5.34 | רקיעים |  |
| 7:5.35 | הם |  |
| 7:5.36 | שמים |  |
| 7:5.37 | ושמי |  |
| 7:5.38 | השמים | . |
| 7:5.39 | רקיע | , |
| 7:5.40 | שחקים | , |
| 7:5.41 | מעון | , |
| 7:5.42 | זבול | , |
| 7:5.43 | ערפל | , |
| 7:5.44 | וכולו |  |
| 7:5.45 | פתח |  |
| 7:5.46 | הקב"ה |  |
| 7:5.47 | לישראל |  |
| 7:5.48 | להודיען |  |
| 7:5.49 | שאין |  |
| 7:5.50 | אלוה |  |
| 7:5.51 | אחר |  |
| 7:5.52 | אלא |  |
| 7:5.53 | הוא | . |
| 7:5.54 | אמרה |  |
| 7:5.55 | כנסת |  |
| 7:5.56 | ישראל |  |
| 7:5.57 | לפני |  |
| 7:5.58 | הקב"ה | . |
| 7:5.59 | מי |  |
| 7:5.60 | לי |  |
| 7:5.61 | בשמים |  |
| 7:5.62 | חוץ |  |
| 7:5.63 | מכבודך | . |
| 7:5.64 | כשם |  |
| 7:5.65 | שאין |  |
| 7:5.66 | לי |  |
| 7:5.67 | בשמים |  |
| 7:5.68 | אלא |  |
| 7:5.69 | אתה | , |
| 7:5.70 | כך |  |
| 7:5.71 | לא |  |
| 7:5.72 | חפצתי |  |
| 7:5.73 | אחר |  |
| 7:5.74 | בארץ | . |
| 7:5.75 | אלא |  |
| 7:5.76 | אני |  |
| 7:5.77 | נכנסת |  |
| 7:5.78 | בכל |  |
| 7:5.79 | יום |  |
| 7:5.80 | לבתי |  |
| 7:5.81 | כנסיות |  |
| 7:5.82 | ומעידה |  |
| 7:5.83 | עליך |  |
| 7:5.84 | שאין |  |
| 7:5.85 | אלוה |  |
| 7:5.86 | אחר |  |
| 7:5.87 | אלא |  |
| 7:5.88 | אתה |  |
| 7:5.89 | ואומרת |  |
| 7:5.90 | שמע |  |
| 7:5.91 | ישראל |  |
| 7:5.92 | ה | ' |
| 7:5.93 | אלהינו |  |
| 7:5.94 | ה | ' |
| 7:5.95 | אחד | . |
| 7:5.96 | הנך |  |
| 7:5.97 | רואה |  |
| 7:5.98 | ידידי |  |
| 7:5.99 | דברי |  |
| 7:5.100 | רז"ל |  |
| 7:5.101 | הנז | ' |
| 7:5.102 | במדרש |  |
| 7:5.103 | עומדים |  |
| 7:5.104 | נגד |  |
| 7:5.105 | הקבלה |  |
| 7:5.106 | החדשה | . |
| 7:5.107 | שהרי |  |
| 7:5.108 | בשמע |  |
| 7:5.109 | ישראל |  |
| 7:5.110 | שהוא |  |
| 7:5.111 | קבלת |  |
| 7:5.112 | עול |  |
| 7:5.113 | מלכות |  |
| 7:5.114 | שמים | , |
| 7:5.115 | המקובלים |  |
| 7:5.116 | החדשים |  |
| 7:5.117 | משתפים |  |
| 7:5.118 | בו |  |
| 7:5.119 | אבא |  |
| 7:5.120 | ואימא |  |
| 7:5.121 | וזעיר |  |
| 7:5.122 | ונוקביה |  |
| 7:5.123 | כמבואר |  |
| 7:5.124 | בזהר |  |
| 7:5.125 | ואתחנן | . |
| 7:5.126 | הובא |  |
| 7:5.127 | לעיל |  |
| 7:5.128 | סי | ' |
| 7:5.129 | מ"ב | , |
| 7:5.130 | מה |  |
| 7:5.131 | שלא |  |
| 7:5.132 | עלה |  |
| 7:5.133 | על |  |
| 7:5.134 | דעת |  |
| 7:5.135 | רז"ל |  |
| 7:5.136 | שייעשה |  |
| 7:5.137 | כן |  |
| 7:5.138 | בישראל |  |
| 7:5.139 | וגדולה |  |
| 7:5.140 | מזו | . |
| 7:5.141 | אמרו |  |
| 7:5.142 | בעלי |  |
| 7:5.143 | הקבלה |  |
| 7:5.144 | החדשה | , |
| 7:5.145 | שהמתפלל |  |
| 7:5.146 | לאין |  |
| 7:5.147 | סוף |  |
| 7:5.148 | הסבה |  |
| 7:5.149 | הראשונה |  |
| 7:5.150 | אין |  |
| 7:5.151 | תפלתו |  |
| 7:5.152 | תפלה | , |
| 7:5.153 | אלא |  |
| 7:5.154 | צריך |  |
| 7:5.155 | להתפלל |  |
| 7:5.156 | דוקא |  |
| 7:5.157 | לז"א |  |
| 7:5.158 | ליחד |  |
| 7:5.159 | זוגתו |  |
| 7:5.160 | עמו |  |
| 7:5.161 | וגם |  |
| 7:5.162 | לשתף |  |
| 7:5.163 | במחשבה |  |
| 7:5.164 | אבא |  |
| 7:5.165 | ואימא |  |
| 7:5.166 | ואריך |  |
| 7:5.167 | אנפין |  |
| 7:5.168 | ועתיק |  |
| 7:5.169 | יומין |  |
| 7:5.170 | כמו |  |
| 7:5.171 | שכתבתי |  |
| 7:5.172 | לעיל |  |
| 7:5.173 | ב | ' |
| 7:5.174 | כסא |  |
| 7:5.175 | אליהו |  |
| 7:5.176 | ויושר |  |
| 7:5.177 | לבב | . |
| 7:5.178 | שהמיחד |  |
| 7:5.179 | תפלתו |  |
| 7:5.180 | לז"א |  |
| 7:5.181 | בלי |  |
| 7:5.182 | צירוף |  |
| 7:5.183 | ושתוף |  |
| 7:5.184 | פרצופים |  |
| 7:5.185 | הנז | ' |
| 7:5.186 | אינו |  |
| 7:5.187 | נענה |  |
| 7:5.188 | מהרה | , |
| 7:5.189 | וזה |  |
| 7:5.190 | סוד |  |
| 7:5.191 | מה |  |
| 7:5.192 | תצעק |  |
| 7:5.193 | אלי |  |
| 7:5.194 | כדלעיל |  |
| 7:5.195 | סי | ' |
| 7:5.196 | ל | '. |
| 7:5.197 | עוד |  |
| 7:5.198 | אמרו |  |
| 7:5.199 | במדרש |  |
| 7:5.200 | רבה |  |
| 7:5.201 | פרשה |  |
| 7:5.202 | הנז"ל | . |
| 7:5.203 | וז"ל |  |
| 7:5.204 | שמע |  |
| 7:5.205 | ישראל |  |
| 7:5.206 | ה | ' |
| 7:5.207 | אלהינו |  |
| 7:5.208 | ה | ' |
| 7:5.209 | אחד | , |
| 7:5.210 | רבנן |  |
| 7:5.211 | אמרי |  |
| 7:5.212 | אמר |  |
| 7:5.213 | הקב"ה |  |
| 7:5.214 | בני |  |
| 7:5.215 | כל |  |
| 7:5.216 | מה |  |
| 7:5.217 | שבראתי |  |
| 7:5.218 | בראתי |  |
| 7:5.219 | זוגות | , |
| 7:5.220 | שמים |  |
| 7:5.221 | וארץ |  |
| 7:5.222 | זוגות | , |
| 7:5.223 | חמה |  |
| 7:5.224 | ולבנה |  |
| 7:5.225 | זוגות | , |
| 7:5.226 | אדם |  |
| 7:5.227 | וחוה |  |
| 7:5.228 | זוגות | , |
| 7:5.229 | העולם |  |
| 7:5.230 | הזה |  |
| 7:5.231 | והעה"ב |  |
| 7:5.232 | זוגות | , |
| 7:5.233 | אבל |  |
| 7:5.234 | כבודי |  |
| 7:5.235 | אחד |  |
| 7:5.236 | ומיוחד |  |
| 7:5.237 | בעולם | , |
| 7:5.238 | מנין |  |
| 7:5.239 | ממה |  |
| 7:5.240 | שקרינו |  |
| 7:5.241 | בענין |  |
| 7:5.242 | שמע |  |
| 7:5.243 | ישראל |  |
| 7:5.244 | ה | ' |
| 7:5.245 | אלהינו |  |
| 7:5.246 | ה | ' |
| 7:5.247 | אחד |  |
| 7:5.248 | עכ"ל | . |
| 7:5.249 | גם |  |
| 7:5.250 | דברי |  |
| 7:5.251 | המדרש |  |
| 7:5.252 | אלו |  |
| 7:5.253 | עומדים |  |
| 7:5.254 | נגד |  |
| 7:5.255 | הקבלה |  |
| 7:5.256 | החדשה |  |
| 7:5.257 | שאמרה |  |
| 7:5.258 | שזעיר |  |
| 7:5.259 | ונוקביה |  |
| 7:5.260 | הוא |  |
| 7:5.261 | אלוהינו |  |
| 7:5.262 | ואנחנו |  |
| 7:5.263 | עמו |  |
| 7:5.264 | וצאן |  |
| 7:5.265 | מריעתו | . |
| 7:5.266 | ואותו |  |
| 7:5.267 | נעבוד |  |
| 7:5.268 | במצות |  |
| 7:5.269 | אבא |  |
| 7:5.270 | ואימא | , |
| 7:5.271 | והוא |  |
| 7:5.272 | משתבח |  |
| 7:5.273 | בעצמו |  |
| 7:5.274 | ואומר |  |
| 7:5.275 | ה | ' |
| 7:5.276 | אמר |  |
| 7:5.277 | אלי |  |
| 7:5.278 | בני |  |
| 7:5.279 | אתה |  |
| 7:5.280 | וכו | ' |
| 7:5.281 | כדלעיל |  |
| 7:5.282 | סי | ' |
| 7:5.283 | כ"ו | . |
| 7:5.284 | שהרי |  |
| 7:5.285 | במדרש |  |
| 7:5.286 | אמרו |  |
| 7:5.287 | כי |  |
| 7:5.288 | ה | ' |
| 7:5.289 | אלהינו |  |
| 7:5.290 | אחד |  |
| 7:5.291 | ומיוחד |  |
| 7:5.292 | מכל |  |
| 7:5.293 | האחדים |  |
| 7:5.294 | שבעולם |  |
| 7:5.295 | כנז | ' |
| 7:5.296 | לעיל | , |
| 7:5.297 | ואין |  |
| 7:5.298 | לו |  |
| 7:5.299 | עזר |  |
| 7:5.300 | כנגדו |  |
| 7:5.301 | לפי |  |
| 7:5.302 | שהוא |  |
| 7:5.303 | יכול |  |
| 7:5.304 | ואינו |  |
| 7:5.305 | צריך |  |
| 7:5.306 | לעזר |  |
| 7:5.307 | זולתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

