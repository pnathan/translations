# Translation pack: milhamot-hashem perek-3-06 (Milchamot ha-Shem 3:10–3:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-06`.

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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
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
- מרגל לרגל: "from festival to festival" (Num 20 ff. sense), not "foot".
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- המשרף (the questioner's pun on Metzaref ha-Emunah): keep "the Mesharef", with a
  note naming the pun.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
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


## Milchamot ha-Shem 3:10

Hebrew (Wikisource): י) ומעתה בעזר השי"ת נשוב לעיקר שאילתנו ששאלנו לך להודיענו על פי שיטת המקובלים, למי אנו קורין בתפלותנו ומברכים כל ברכותנו ועבודתנו וממי אנו מבקשים סליחה וכפרה. והיתה תשובתך בדברי נרגן כמתלהמים טענו חטים והודה לו בשעורים והכחשת את הידוע ומפורסם בספריהם לכל, כנשבע על האיש שהוא אשה, ועל האשה שהיא איש, ועל עמוד של שיש שהוא זהב. וחרה אפך על ספר מצרף האמונה, והלעגת עליו ובזית אותו, וקראת אותו ספר המשרף וכו' ובזה גלית לכל שאתה הוא המבזה את החכמים, ועל שאלותינו כיון שזעיר אנפין לפי דבריהם הוא אלוהינו, והוא הזן אותנו ומפרנסינו, ואותו נעבוד, אם כן כל העולמות שלמעלה מעולם האצילות, מי זן אותם ומפרנסם? אם זורק הוא שפעם ופרנסתם מלמטה למעלה, והם עובדים ומשתחוים לו, או שמא אלהיהם הזנם ומפרנסם, הוא עלה הראשונה, ולו הם עובדים, ואנחנו עובדים לזעיר אנפין הזנינו ומפרנסינו? ואם כן אלהי העליונים אינו אלהי התחתונים ח"ו? וזהו מקום הסתירה אשר בקשנו ממך להודיענו ולהורותנו, ועל זה באה תשובתך שאין הנדון דומה לראיה וכי מביאין ראיה לסתור הלכה למשה מסיני? ושחכמי הקבלה אומרים כך קיבלו ממשה.

| id | word | punctuation after |
|---|---|---|
| 3:10.1 | י | ) |
| 3:10.2 | ומעתה |  |
| 3:10.3 | בעזר |  |
| 3:10.4 | השי"ת |  |
| 3:10.5 | נשוב |  |
| 3:10.6 | לעיקר |  |
| 3:10.7 | שאילתנו |  |
| 3:10.8 | ששאלנו |  |
| 3:10.9 | לך |  |
| 3:10.10 | להודיענו |  |
| 3:10.11 | על |  |
| 3:10.12 | פי |  |
| 3:10.13 | שיטת |  |
| 3:10.14 | המקובלים | , |
| 3:10.15 | למי |  |
| 3:10.16 | אנו |  |
| 3:10.17 | קורין |  |
| 3:10.18 | בתפלותנו |  |
| 3:10.19 | ומברכים |  |
| 3:10.20 | כל |  |
| 3:10.21 | ברכותנו |  |
| 3:10.22 | ועבודתנו |  |
| 3:10.23 | וממי |  |
| 3:10.24 | אנו |  |
| 3:10.25 | מבקשים |  |
| 3:10.26 | סליחה |  |
| 3:10.27 | וכפרה | . |
| 3:10.28 | והיתה |  |
| 3:10.29 | תשובתך |  |
| 3:10.30 | בדברי |  |
| 3:10.31 | נרגן |  |
| 3:10.32 | כמתלהמים |  |
| 3:10.33 | טענו |  |
| 3:10.34 | חטים |  |
| 3:10.35 | והודה |  |
| 3:10.36 | לו |  |
| 3:10.37 | בשעורים |  |
| 3:10.38 | והכחשת |  |
| 3:10.39 | את |  |
| 3:10.40 | הידוע |  |
| 3:10.41 | ומפורסם |  |
| 3:10.42 | בספריהם |  |
| 3:10.43 | לכל | , |
| 3:10.44 | כנשבע |  |
| 3:10.45 | על |  |
| 3:10.46 | האיש |  |
| 3:10.47 | שהוא |  |
| 3:10.48 | אשה | , |
| 3:10.49 | ועל |  |
| 3:10.50 | האשה |  |
| 3:10.51 | שהיא |  |
| 3:10.52 | איש | , |
| 3:10.53 | ועל |  |
| 3:10.54 | עמוד |  |
| 3:10.55 | של |  |
| 3:10.56 | שיש |  |
| 3:10.57 | שהוא |  |
| 3:10.58 | זהב | . |
| 3:10.59 | וחרה |  |
| 3:10.60 | אפך |  |
| 3:10.61 | על |  |
| 3:10.62 | ספר |  |
| 3:10.63 | מצרף |  |
| 3:10.64 | האמונה | , |
| 3:10.65 | והלעגת |  |
| 3:10.66 | עליו |  |
| 3:10.67 | ובזית |  |
| 3:10.68 | אותו | , |
| 3:10.69 | וקראת |  |
| 3:10.70 | אותו |  |
| 3:10.71 | ספר |  |
| 3:10.72 | המשרף |  |
| 3:10.73 | וכו | ' |
| 3:10.74 | ובזה |  |
| 3:10.75 | גלית |  |
| 3:10.76 | לכל |  |
| 3:10.77 | שאתה |  |
| 3:10.78 | הוא |  |
| 3:10.79 | המבזה |  |
| 3:10.80 | את |  |
| 3:10.81 | החכמים | , |
| 3:10.82 | ועל |  |
| 3:10.83 | שאלותינו |  |
| 3:10.84 | כיון |  |
| 3:10.85 | שזעיר |  |
| 3:10.86 | אנפין |  |
| 3:10.87 | לפי |  |
| 3:10.88 | דבריהם |  |
| 3:10.89 | הוא |  |
| 3:10.90 | אלוהינו | , |
| 3:10.91 | והוא |  |
| 3:10.92 | הזן |  |
| 3:10.93 | אותנו |  |
| 3:10.94 | ומפרנסינו | , |
| 3:10.95 | ואותו |  |
| 3:10.96 | נעבוד | , |
| 3:10.97 | אם |  |
| 3:10.98 | כן |  |
| 3:10.99 | כל |  |
| 3:10.100 | העולמות |  |
| 3:10.101 | שלמעלה |  |
| 3:10.102 | מעולם |  |
| 3:10.103 | האצילות | , |
| 3:10.104 | מי |  |
| 3:10.105 | זן |  |
| 3:10.106 | אותם |  |
| 3:10.107 | ומפרנסם | ? |
| 3:10.108 | אם |  |
| 3:10.109 | זורק |  |
| 3:10.110 | הוא |  |
| 3:10.111 | שפעם |  |
| 3:10.112 | ופרנסתם |  |
| 3:10.113 | מלמטה |  |
| 3:10.114 | למעלה | , |
| 3:10.115 | והם |  |
| 3:10.116 | עובדים |  |
| 3:10.117 | ומשתחוים |  |
| 3:10.118 | לו | , |
| 3:10.119 | או |  |
| 3:10.120 | שמא |  |
| 3:10.121 | אלהיהם |  |
| 3:10.122 | הזנם |  |
| 3:10.123 | ומפרנסם | , |
| 3:10.124 | הוא |  |
| 3:10.125 | עלה |  |
| 3:10.126 | הראשונה | , |
| 3:10.127 | ולו |  |
| 3:10.128 | הם |  |
| 3:10.129 | עובדים | , |
| 3:10.130 | ואנחנו |  |
| 3:10.131 | עובדים |  |
| 3:10.132 | לזעיר |  |
| 3:10.133 | אנפין |  |
| 3:10.134 | הזנינו |  |
| 3:10.135 | ומפרנסינו | ? |
| 3:10.136 | ואם |  |
| 3:10.137 | כן |  |
| 3:10.138 | אלהי |  |
| 3:10.139 | העליונים |  |
| 3:10.140 | אינו |  |
| 3:10.141 | אלהי |  |
| 3:10.142 | התחתונים |  |
| 3:10.143 | ח"ו | ? |
| 3:10.144 | וזהו |  |
| 3:10.145 | מקום |  |
| 3:10.146 | הסתירה |  |
| 3:10.147 | אשר |  |
| 3:10.148 | בקשנו |  |
| 3:10.149 | ממך |  |
| 3:10.150 | להודיענו |  |
| 3:10.151 | ולהורותנו | , |
| 3:10.152 | ועל |  |
| 3:10.153 | זה |  |
| 3:10.154 | באה |  |
| 3:10.155 | תשובתך |  |
| 3:10.156 | שאין |  |
| 3:10.157 | הנדון |  |
| 3:10.158 | דומה |  |
| 3:10.159 | לראיה |  |
| 3:10.160 | וכי |  |
| 3:10.161 | מביאין |  |
| 3:10.162 | ראיה |  |
| 3:10.163 | לסתור |  |
| 3:10.164 | הלכה |  |
| 3:10.165 | למשה |  |
| 3:10.166 | מסיני | ? |
| 3:10.167 | ושחכמי |  |
| 3:10.168 | הקבלה |  |
| 3:10.169 | אומרים |  |
| 3:10.170 | כך |  |
| 3:10.171 | קיבלו |  |
| 3:10.172 | ממשה | . |


## Milchamot ha-Shem 3:11

Hebrew (Wikisource): יא) מאד אנו תמהים ומתפלאים על דבריך ועל תשובתך זאת, וכי הלכה למשה מסיני להאמין באלוהות רבים? ולעבוד אלהים אחרים מבלעדי ה' אלהינו הראשון וצור לכל? והוא סבת כל הסבות כולם, ובמרמה את בא בתשובותיך אלה לאנשים עזי לב אשר התפארת עליהם והראית להם דבריך אלה לאמר. ראו כי סתרתי את דבריהם, ודחיתי את ראיותיהם בקש, וחוזר קונדריסך על כל המסובין, ואחר אשר קילסוך ושבחוך שאתה חריף ושנון הבאת אותה לנו. והנה כולה הטעאות כאלו לא ידעת על מה תסוב שאלתינו, ועל טעותך הכוזבת הלזו נשוב עוד, כי ידוע ומפורסם בכללי הפוסקים (כנה"ג ורדב"ז ושה"ז) שכל מקום שהקבלה חולקת על הפסק הלך אחרי הפסק, והכלל הזה מפורסם בישראל. ואם דבריך שהקבלה הלזו היא הלכה למשה מסיני, וכי תדחה הלכה למשה מפני סברת איזה תנא או אמורא, וכ"ש פוסק אחרון (עיין לקמן סי' ט"ו) ועליך להביא ראיה ולהודיענו היכן מצינו זה ששום תנא או אמורא יעיז פניו לחלוק ולהורות היפך הלכה המקובלת ממשה רבי' ע"ה. הלא שוטה או רשע וגס רוח יקרא? ובכמה מקומות בתלמוד מרגלא בפומיהו דרבנן לומר אם הלכה נקבל, ואם לדין יש תשובה.

| id | word | punctuation after |
|---|---|---|
| 3:11.1 | יא | ) |
| 3:11.2 | מאד |  |
| 3:11.3 | אנו |  |
| 3:11.4 | תמהים |  |
| 3:11.5 | ומתפלאים |  |
| 3:11.6 | על |  |
| 3:11.7 | דבריך |  |
| 3:11.8 | ועל |  |
| 3:11.9 | תשובתך |  |
| 3:11.10 | זאת | , |
| 3:11.11 | וכי |  |
| 3:11.12 | הלכה |  |
| 3:11.13 | למשה |  |
| 3:11.14 | מסיני |  |
| 3:11.15 | להאמין |  |
| 3:11.16 | באלוהות |  |
| 3:11.17 | רבים | ? |
| 3:11.18 | ולעבוד |  |
| 3:11.19 | אלהים |  |
| 3:11.20 | אחרים |  |
| 3:11.21 | מבלעדי |  |
| 3:11.22 | ה | ' |
| 3:11.23 | אלהינו |  |
| 3:11.24 | הראשון |  |
| 3:11.25 | וצור |  |
| 3:11.26 | לכל | ? |
| 3:11.27 | והוא |  |
| 3:11.28 | סבת |  |
| 3:11.29 | כל |  |
| 3:11.30 | הסבות |  |
| 3:11.31 | כולם | , |
| 3:11.32 | ובמרמה |  |
| 3:11.33 | את |  |
| 3:11.34 | בא |  |
| 3:11.35 | בתשובותיך |  |
| 3:11.36 | אלה |  |
| 3:11.37 | לאנשים |  |
| 3:11.38 | עזי |  |
| 3:11.39 | לב |  |
| 3:11.40 | אשר |  |
| 3:11.41 | התפארת |  |
| 3:11.42 | עליהם |  |
| 3:11.43 | והראית |  |
| 3:11.44 | להם |  |
| 3:11.45 | דבריך |  |
| 3:11.46 | אלה |  |
| 3:11.47 | לאמר | . |
| 3:11.48 | ראו |  |
| 3:11.49 | כי |  |
| 3:11.50 | סתרתי |  |
| 3:11.51 | את |  |
| 3:11.52 | דבריהם | , |
| 3:11.53 | ודחיתי |  |
| 3:11.54 | את |  |
| 3:11.55 | ראיותיהם |  |
| 3:11.56 | בקש | , |
| 3:11.57 | וחוזר |  |
| 3:11.58 | קונדריסך |  |
| 3:11.59 | על |  |
| 3:11.60 | כל |  |
| 3:11.61 | המסובין | , |
| 3:11.62 | ואחר |  |
| 3:11.63 | אשר |  |
| 3:11.64 | קילסוך |  |
| 3:11.65 | ושבחוך |  |
| 3:11.66 | שאתה |  |
| 3:11.67 | חריף |  |
| 3:11.68 | ושנון |  |
| 3:11.69 | הבאת |  |
| 3:11.70 | אותה |  |
| 3:11.71 | לנו | . |
| 3:11.72 | והנה |  |
| 3:11.73 | כולה |  |
| 3:11.74 | הטעאות |  |
| 3:11.75 | כאלו |  |
| 3:11.76 | לא |  |
| 3:11.77 | ידעת |  |
| 3:11.78 | על |  |
| 3:11.79 | מה |  |
| 3:11.80 | תסוב |  |
| 3:11.81 | שאלתינו | , |
| 3:11.82 | ועל |  |
| 3:11.83 | טעותך |  |
| 3:11.84 | הכוזבת |  |
| 3:11.85 | הלזו |  |
| 3:11.86 | נשוב |  |
| 3:11.87 | עוד | , |
| 3:11.88 | כי |  |
| 3:11.89 | ידוע |  |
| 3:11.90 | ומפורסם |  |
| 3:11.91 | בכללי |  |
| 3:11.92 | הפוסקים |  |
| 3:11.93 | כנה"ג |  |
| 3:11.94 | ורדב"ז |  |
| 3:11.95 | ושה"ז | ) |
| 3:11.96 | שכל |  |
| 3:11.97 | מקום |  |
| 3:11.98 | שהקבלה |  |
| 3:11.99 | חולקת |  |
| 3:11.100 | על |  |
| 3:11.101 | הפסק |  |
| 3:11.102 | הלך |  |
| 3:11.103 | אחרי |  |
| 3:11.104 | הפסק | , |
| 3:11.105 | והכלל |  |
| 3:11.106 | הזה |  |
| 3:11.107 | מפורסם |  |
| 3:11.108 | בישראל | . |
| 3:11.109 | ואם |  |
| 3:11.110 | דבריך |  |
| 3:11.111 | שהקבלה |  |
| 3:11.112 | הלזו |  |
| 3:11.113 | היא |  |
| 3:11.114 | הלכה |  |
| 3:11.115 | למשה |  |
| 3:11.116 | מסיני | , |
| 3:11.117 | וכי |  |
| 3:11.118 | תדחה |  |
| 3:11.119 | הלכה |  |
| 3:11.120 | למשה |  |
| 3:11.121 | מפני |  |
| 3:11.122 | סברת |  |
| 3:11.123 | איזה |  |
| 3:11.124 | תנא |  |
| 3:11.125 | או |  |
| 3:11.126 | אמורא | , |
| 3:11.127 | וכ"ש |  |
| 3:11.128 | פוסק |  |
| 3:11.129 | אחרון |  |
| 3:11.130 | עיין |  |
| 3:11.131 | לקמן |  |
| 3:11.132 | סי | ' |
| 3:11.133 | ט"ו | ) |
| 3:11.134 | ועליך |  |
| 3:11.135 | להביא |  |
| 3:11.136 | ראיה |  |
| 3:11.137 | ולהודיענו |  |
| 3:11.138 | היכן |  |
| 3:11.139 | מצינו |  |
| 3:11.140 | זה |  |
| 3:11.141 | ששום |  |
| 3:11.142 | תנא |  |
| 3:11.143 | או |  |
| 3:11.144 | אמורא |  |
| 3:11.145 | יעיז |  |
| 3:11.146 | פניו |  |
| 3:11.147 | לחלוק |  |
| 3:11.148 | ולהורות |  |
| 3:11.149 | היפך |  |
| 3:11.150 | הלכה |  |
| 3:11.151 | המקובלת |  |
| 3:11.152 | ממשה |  |
| 3:11.153 | רבי | ' |
| 3:11.154 | ע"ה | . |
| 3:11.155 | הלא |  |
| 3:11.156 | שוטה |  |
| 3:11.157 | או |  |
| 3:11.158 | רשע |  |
| 3:11.159 | וגס |  |
| 3:11.160 | רוח |  |
| 3:11.161 | יקרא | ? |
| 3:11.162 | ובכמה |  |
| 3:11.163 | מקומות |  |
| 3:11.164 | בתלמוד |  |
| 3:11.165 | מרגלא |  |
| 3:11.166 | בפומיהו |  |
| 3:11.167 | דרבנן |  |
| 3:11.168 | לומר |  |
| 3:11.169 | אם |  |
| 3:11.170 | הלכה |  |
| 3:11.171 | נקבל | , |
| 3:11.172 | ואם |  |
| 3:11.173 | לדין |  |
| 3:11.174 | יש |  |
| 3:11.175 | תשובה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

