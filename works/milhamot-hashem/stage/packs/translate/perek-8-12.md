# Translation pack: milhamot-hashem perek-8-12 (Milchamot ha-Shem 8:14–8:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-12`.

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
- **קצר אפים**: the author's Hebrew calque of זעיר אנפין, set against ארך
  אפים in Exod 34:6. Write "quick to anger", and render ארך אפים "slow to
  anger" everywhere, so that the pair shows in English. "באל זר קצר אפים
  (ז״א)" (1:4) → "a strange god, quick to anger (Ze'ir Anpin)"; "הנקרא (קצר
  אפים) זעיר אנפין" (3:3) → "called (quick to anger) Ze'ir Anpin". Add a note
  each time: Aramaic "small face"; Prov 14:17. "קצא אפים" (6:23) is translated
  as קצר, with the spelling recorded in `apparatus`.
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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- האלהי before a name: "the divine", with a note; never dropped.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".

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


## Milchamot ha-Shem 8:14

Hebrew (Wikisource): קט) וכיון שאיוו לאלוהות הרבה קיימו וקיבלו שהמציא האלוה העליון הנקרא אין סיבות ועילות רבים המשלתשלים ונשפעים זה מזה, וכל אחד מהם נקרא בשם הוי'ה או אדנות או אלהים והקב"ה וכל השמות, ודרשו את הכתוב מה שמו ומה שם בנו כי תדע, מה שמו ה' צבאות שמו, (פי' אבא) ומה שם בנו, זעיר אנפין שמו וכו' כדלעיל סי' כ"ו. וכל עבודתינו ותפלותינו לפי דבריהם אל האלהות האלו הנמשכים מהאלוה העליון שקראוהו אין, כי האין לא יענה את הקורא אליו, כי אם האלהים התחתונים המה יענו, ובלבד שניחד תפלתינו אל האלוה קצר האפים (ז"א) ונשתף במחשבתינו גם פרצופים שלמעלה ממנו. יוחשב לנו עבודת קצר האפים הנברא הזה כאלו אנחנו עובדים וקורים אל האלוה העליון, כי נשמתו היא חלק אלוה ממעל.

| id | word | punctuation after |
|---|---|---|
| 8:14.1 | קט | ) |
| 8:14.2 | וכיון |  |
| 8:14.3 | שאיוו |  |
| 8:14.4 | לאלוהות |  |
| 8:14.5 | הרבה |  |
| 8:14.6 | קיימו |  |
| 8:14.7 | וקיבלו |  |
| 8:14.8 | שהמציא |  |
| 8:14.9 | האלוה |  |
| 8:14.10 | העליון |  |
| 8:14.11 | הנקרא |  |
| 8:14.12 | אין |  |
| 8:14.13 | סיבות |  |
| 8:14.14 | ועילות |  |
| 8:14.15 | רבים |  |
| 8:14.16 | המשלתשלים |  |
| 8:14.17 | ונשפעים |  |
| 8:14.18 | זה |  |
| 8:14.19 | מזה | , |
| 8:14.20 | וכל |  |
| 8:14.21 | אחד |  |
| 8:14.22 | מהם |  |
| 8:14.23 | נקרא |  |
| 8:14.24 | בשם |  |
| 8:14.25 | הוי'ה |  |
| 8:14.26 | או |  |
| 8:14.27 | אדנות |  |
| 8:14.28 | או |  |
| 8:14.29 | אלהים |  |
| 8:14.30 | והקב"ה |  |
| 8:14.31 | וכל |  |
| 8:14.32 | השמות | , |
| 8:14.33 | ודרשו |  |
| 8:14.34 | את |  |
| 8:14.35 | הכתוב |  |
| 8:14.36 | מה |  |
| 8:14.37 | שמו |  |
| 8:14.38 | ומה |  |
| 8:14.39 | שם |  |
| 8:14.40 | בנו |  |
| 8:14.41 | כי |  |
| 8:14.42 | תדע | , |
| 8:14.43 | מה |  |
| 8:14.44 | שמו |  |
| 8:14.45 | ה | ' |
| 8:14.46 | צבאות |  |
| 8:14.47 | שמו | , |
| 8:14.48 | פי | ' |
| 8:14.49 | אבא | ) |
| 8:14.50 | ומה |  |
| 8:14.51 | שם |  |
| 8:14.52 | בנו | , |
| 8:14.53 | זעיר |  |
| 8:14.54 | אנפין |  |
| 8:14.55 | שמו |  |
| 8:14.56 | וכו | ' |
| 8:14.57 | כדלעיל |  |
| 8:14.58 | סי | ' |
| 8:14.59 | כ"ו | . |
| 8:14.60 | וכל |  |
| 8:14.61 | עבודתינו |  |
| 8:14.62 | ותפלותינו |  |
| 8:14.63 | לפי |  |
| 8:14.64 | דבריהם |  |
| 8:14.65 | אל |  |
| 8:14.66 | האלהות |  |
| 8:14.67 | האלו |  |
| 8:14.68 | הנמשכים |  |
| 8:14.69 | מהאלוה |  |
| 8:14.70 | העליון |  |
| 8:14.71 | שקראוהו |  |
| 8:14.72 | אין | , |
| 8:14.73 | כי |  |
| 8:14.74 | האין |  |
| 8:14.75 | לא |  |
| 8:14.76 | יענה |  |
| 8:14.77 | את |  |
| 8:14.78 | הקורא |  |
| 8:14.79 | אליו | , |
| 8:14.80 | כי |  |
| 8:14.81 | אם |  |
| 8:14.82 | האלהים |  |
| 8:14.83 | התחתונים |  |
| 8:14.84 | המה |  |
| 8:14.85 | יענו | , |
| 8:14.86 | ובלבד |  |
| 8:14.87 | שניחד |  |
| 8:14.88 | תפלתינו |  |
| 8:14.89 | אל |  |
| 8:14.90 | האלוה |  |
| 8:14.91 | קצר |  |
| 8:14.92 | האפים |  |
| 8:14.93 | ז"א | ) |
| 8:14.94 | ונשתף |  |
| 8:14.95 | במחשבתינו |  |
| 8:14.96 | גם |  |
| 8:14.97 | פרצופים |  |
| 8:14.98 | שלמעלה |  |
| 8:14.99 | ממנו | . |
| 8:14.100 | יוחשב |  |
| 8:14.101 | לנו |  |
| 8:14.102 | עבודת |  |
| 8:14.103 | קצר |  |
| 8:14.104 | האפים |  |
| 8:14.105 | הנברא |  |
| 8:14.106 | הזה |  |
| 8:14.107 | כאלו |  |
| 8:14.108 | אנחנו |  |
| 8:14.109 | עובדים |  |
| 8:14.110 | וקורים |  |
| 8:14.111 | אל |  |
| 8:14.112 | האלוה |  |
| 8:14.113 | העליון | , |
| 8:14.114 | כי |  |
| 8:14.115 | נשמתו |  |
| 8:14.116 | היא |  |
| 8:14.117 | חלק |  |
| 8:14.118 | אלוה |  |
| 8:14.119 | ממעל | . |


## Milchamot ha-Shem 8:15

Hebrew (Wikisource): ולכן הוכיח ירמיה הנביא את ישראל ואמר כחשו בה' שאמר למשה כה תאמר אל בני ישראל יהוה אלהי אבותיכם אלהי אברהם, אלהי יצחק ואלהי יעקב שלחני אליכם זה שמי לעולם וזה זכרי לדר דר. ויאמרו לא הוא המטיב ומשגיח עלינו, ולא אותו נעבוד אלא קצר האפים הוא המשגיח עלינו, ואליו נקרא ואותו נעבוד, והוא המלך הקדוש המהולל בכל התשבחות, כי לו ניתן הכח והממשלה מאביו ואמו, ושהאלהים העליון הוא המאחד ומדבק ומקשר את כל האלהות התחתונים ופרצופיהם להיות כולם נחשבים אחד, אע"פ שהם באמת חלוקים. ובקראינו אל הזעיר אנפין כאלו קראנו אל האלהים העליון, ולכך כתיב לוא מלא וא"ו ואלף קרי ביה לו, וקרי ביה לא, לא אותו נעבוד, כי אם את קצר האפים, וקרי ביה לו בוא"ו כי בעבודתינו לז"א כאלו לו אנחנו עובדים, כי זהו רצונו של מלך העליון שנעבוד את המלך התחתון (זעיר אנפין) שמינהו תחתיו והשליטו על כל הברואים התחתונים וגם ליחד אשתו עמו. ובאמרינו שמע ישראל ה' אלהינו ה' אחד אנו מיחדים לפי דבריהם גם את אבא ואימא עמו כי ה' הראשון הוא אבא, אלהינו אימא, ה' דא זעיר אנפין, אחד דא מלכות, גי א"ת ט' ספירות, ודל"ת כולל הכל להיות אחד, ובכוונה זאת אנו מאחדים ואוגדים את כולם להיות אחד, מלכא עלאה אבא ואימא, ומלכא תתהה זעיר ונוקביה, לקשרא דא בדא ולא ישתכח פירודא אלא כלא חד כדלעיל סי' מ"ב. הה! אוי! ואבוי! על רוע מצב אמונת העם אשר בחר לו יה ישראל לסגולתו. ונתן לנו תורת אמת. היא חותמו של הקב"ה אשר קבל משה רבינו עליו השלום בסיני בכתב ובעל פה. ועמד עליה מ' יום וארבעים לילה ללמוד דקדוקיה וטעמיה וכללותיה, ומסרה ליהושע וכו', ואנשי כנסת הגדולה תפסו ליצרא דע"ז ושדויה בדודא דאברא וחפיוה לפומיה באברא. הן עתה הוסר כסוי העופרת מעל פי האיפה על יד המלך העכו"ם במסרו את בפר הזהר לחכמי טוליטולה, יוצא השטן וישב וירקד בנינו, ויהי לריק כל יגיעם של אנשי כנסת הגדולה אשר טרחו ויגעו. וצעקו לפני השי"ת על יצרא דע"ז. ומן השמיים סייעו אותם ונמסר בידם. ורוב חכמתם ורוח בינתם ביטלו יצרא דע"ז על ידי עסק התורה וביאור המבינים לעם בה על פי האמת והצדק כמו שהוא ערוך לפנינו בב' התלמודים ובמדרשים האמתיים. שג' ויקראו בו לפני הרחוב מאור הבקר עד מחצית היום נגד האנשים והנשים המבינים. ונאמר הלוים מבינים את העם לתורה ויקראו בס' בתורת האלהים מפורש ושום שכל ויבינו במקרא. ועל ידי תקון נוסח התפלות והברכות אשר תקנו לכל ישראל. ועל ידי השתדלותם להעמיד תלמידים הרבה כדאיתא בפרקא קמא דאבות! על ידי כל אלה עלתה בידם לעקור יצרא דע"ז, והחזירו עטרה לישנה, ונמחה שם האלילים וכל האמונות הכוזבות המביאות לידי מינות מישראל, כדגרסינן ביומא פרק בא לו, ויקם על מעלה הלוים ישוע ובני וקדמיאל שבניה בני שרביה בני כנני ויזעקו בקול גדול אל ה' אלהיהם, מאי אמור אמר רב יהודה אמר רב ואי תימא ר' יוחנן בייא בייא היינו האיי דאחרביה לבי מקדשא וקלייה להיכלא וקטלניהו לצדיקיא ואגלינהו לישראל מארעהון ועדיין מרקיד בינא, כלום יהבתיה ניהלן אלא לקבולי ביה אגרא. לא איהו בעינן ולא אגריה, בעינן, נפל להו פתקא מרקיעא דכתיב ביה אמת. אמר רבי חנינה שמע מינה חותמו של הקב"ה אמת, יתיבו בתעניתא תלתה יומין ותלתא לילי מסריה ניהלינו נפק אתא כי גוריא דנורא מבית קדש הקדשים. אמרו להו נביא לישראל היינו הוא יצרא דע"ז שנאמר ויאמר זאת הרשעה וגומר בהדי דתפסי ליה, אשתמיט ביניתא ממזיה ורמא קלא, אזל קליה ארבע מאה פרסי אמר היכי נעביד ליה, דלמא מרחמי עליה מן שמיא. אמר להו נביא שדויה בדורא דאברא והפויה לפומא באברא דשאיב קלא, שנאמר זאת הרשעה וישלך אותה אל תוך האיפה וישלך את אבן העופרת אל פיה. אמרי הואיל ועת רצון היא, נבעי רחמי על יצרא דעבירה. בעו רחמי ואמסר בידיהו, אמר להו נביא הזו ודאי קטלתון לההוא כליא עלמא, חבשוה תלתא יומי. איבעי ביעתא בת יומא לחולה בכל ארץ ישראל ולא אשתכח. אמרי היכי נעביד, ניקטליה איחרב כולי עלמא, נבעי רחמי לפלגא, פלגא מרקיעא לא יהבי. כחלינהו לעיניה ושבקוה, ואהני דלא תגרי ביה באיניש בקרובתא ע"כ. ועיין עץ יוסף, ובענף שכתב שהמאמר תמוה הוא מאד, ויצרא דע"ז בבית קדש הקדשים מאי עבידתיה? וכו'.

| id | word | punctuation after |
|---|---|---|
| 8:15.1 | ולכן |  |
| 8:15.2 | הוכיח |  |
| 8:15.3 | ירמיה |  |
| 8:15.4 | הנביא |  |
| 8:15.5 | את |  |
| 8:15.6 | ישראל |  |
| 8:15.7 | ואמר |  |
| 8:15.8 | כחשו |  |
| 8:15.9 | בה | ' |
| 8:15.10 | שאמר |  |
| 8:15.11 | למשה |  |
| 8:15.12 | כה |  |
| 8:15.13 | תאמר |  |
| 8:15.14 | אל |  |
| 8:15.15 | בני |  |
| 8:15.16 | ישראל |  |
| 8:15.17 | יהוה |  |
| 8:15.18 | אלהי |  |
| 8:15.19 | אבותיכם |  |
| 8:15.20 | אלהי |  |
| 8:15.21 | אברהם | , |
| 8:15.22 | אלהי |  |
| 8:15.23 | יצחק |  |
| 8:15.24 | ואלהי |  |
| 8:15.25 | יעקב |  |
| 8:15.26 | שלחני |  |
| 8:15.27 | אליכם |  |
| 8:15.28 | זה |  |
| 8:15.29 | שמי |  |
| 8:15.30 | לעולם |  |
| 8:15.31 | וזה |  |
| 8:15.32 | זכרי |  |
| 8:15.33 | לדר |  |
| 8:15.34 | דר | . |
| 8:15.35 | ויאמרו |  |
| 8:15.36 | לא |  |
| 8:15.37 | הוא |  |
| 8:15.38 | המטיב |  |
| 8:15.39 | ומשגיח |  |
| 8:15.40 | עלינו | , |
| 8:15.41 | ולא |  |
| 8:15.42 | אותו |  |
| 8:15.43 | נעבוד |  |
| 8:15.44 | אלא |  |
| 8:15.45 | קצר |  |
| 8:15.46 | האפים |  |
| 8:15.47 | הוא |  |
| 8:15.48 | המשגיח |  |
| 8:15.49 | עלינו | , |
| 8:15.50 | ואליו |  |
| 8:15.51 | נקרא |  |
| 8:15.52 | ואותו |  |
| 8:15.53 | נעבוד | , |
| 8:15.54 | והוא |  |
| 8:15.55 | המלך |  |
| 8:15.56 | הקדוש |  |
| 8:15.57 | המהולל |  |
| 8:15.58 | בכל |  |
| 8:15.59 | התשבחות | , |
| 8:15.60 | כי |  |
| 8:15.61 | לו |  |
| 8:15.62 | ניתן |  |
| 8:15.63 | הכח |  |
| 8:15.64 | והממשלה |  |
| 8:15.65 | מאביו |  |
| 8:15.66 | ואמו | , |
| 8:15.67 | ושהאלהים |  |
| 8:15.68 | העליון |  |
| 8:15.69 | הוא |  |
| 8:15.70 | המאחד |  |
| 8:15.71 | ומדבק |  |
| 8:15.72 | ומקשר |  |
| 8:15.73 | את |  |
| 8:15.74 | כל |  |
| 8:15.75 | האלהות |  |
| 8:15.76 | התחתונים |  |
| 8:15.77 | ופרצופיהם |  |
| 8:15.78 | להיות |  |
| 8:15.79 | כולם |  |
| 8:15.80 | נחשבים |  |
| 8:15.81 | אחד | , |
| 8:15.82 | אע"פ |  |
| 8:15.83 | שהם |  |
| 8:15.84 | באמת |  |
| 8:15.85 | חלוקים | . |
| 8:15.86 | ובקראינו |  |
| 8:15.87 | אל |  |
| 8:15.88 | הזעיר |  |
| 8:15.89 | אנפין |  |
| 8:15.90 | כאלו |  |
| 8:15.91 | קראנו |  |
| 8:15.92 | אל |  |
| 8:15.93 | האלהים |  |
| 8:15.94 | העליון | , |
| 8:15.95 | ולכך |  |
| 8:15.96 | כתיב |  |
| 8:15.97 | לוא |  |
| 8:15.98 | מלא |  |
| 8:15.99 | וא"ו |  |
| 8:15.100 | ואלף |  |
| 8:15.101 | קרי |  |
| 8:15.102 | ביה |  |
| 8:15.103 | לו | , |
| 8:15.104 | וקרי |  |
| 8:15.105 | ביה |  |
| 8:15.106 | לא | , |
| 8:15.107 | לא |  |
| 8:15.108 | אותו |  |
| 8:15.109 | נעבוד | , |
| 8:15.110 | כי |  |
| 8:15.111 | אם |  |
| 8:15.112 | את |  |
| 8:15.113 | קצר |  |
| 8:15.114 | האפים | , |
| 8:15.115 | וקרי |  |
| 8:15.116 | ביה |  |
| 8:15.117 | לו |  |
| 8:15.118 | בוא"ו |  |
| 8:15.119 | כי |  |
| 8:15.120 | בעבודתינו |  |
| 8:15.121 | לז"א |  |
| 8:15.122 | כאלו |  |
| 8:15.123 | לו |  |
| 8:15.124 | אנחנו |  |
| 8:15.125 | עובדים | , |
| 8:15.126 | כי |  |
| 8:15.127 | זהו |  |
| 8:15.128 | רצונו |  |
| 8:15.129 | של |  |
| 8:15.130 | מלך |  |
| 8:15.131 | העליון |  |
| 8:15.132 | שנעבוד |  |
| 8:15.133 | את |  |
| 8:15.134 | המלך |  |
| 8:15.135 | התחתון |  |
| 8:15.136 | זעיר |  |
| 8:15.137 | אנפין | ) |
| 8:15.138 | שמינהו |  |
| 8:15.139 | תחתיו |  |
| 8:15.140 | והשליטו |  |
| 8:15.141 | על |  |
| 8:15.142 | כל |  |
| 8:15.143 | הברואים |  |
| 8:15.144 | התחתונים |  |
| 8:15.145 | וגם |  |
| 8:15.146 | ליחד |  |
| 8:15.147 | אשתו |  |
| 8:15.148 | עמו | . |
| 8:15.149 | ובאמרינו |  |
| 8:15.150 | שמע |  |
| 8:15.151 | ישראל |  |
| 8:15.152 | ה | ' |
| 8:15.153 | אלהינו |  |
| 8:15.154 | ה | ' |
| 8:15.155 | אחד |  |
| 8:15.156 | אנו |  |
| 8:15.157 | מיחדים |  |
| 8:15.158 | לפי |  |
| 8:15.159 | דבריהם |  |
| 8:15.160 | גם |  |
| 8:15.161 | את |  |
| 8:15.162 | אבא |  |
| 8:15.163 | ואימא |  |
| 8:15.164 | עמו |  |
| 8:15.165 | כי |  |
| 8:15.166 | ה | ' |
| 8:15.167 | הראשון |  |
| 8:15.168 | הוא |  |
| 8:15.169 | אבא | , |
| 8:15.170 | אלהינו |  |
| 8:15.171 | אימא | , |
| 8:15.172 | ה | ' |
| 8:15.173 | דא |  |
| 8:15.174 | זעיר |  |
| 8:15.175 | אנפין | , |
| 8:15.176 | אחד |  |
| 8:15.177 | דא |  |
| 8:15.178 | מלכות | , |
| 8:15.179 | גי |  |
| 8:15.180 | א"ת |  |
| 8:15.181 | ט | ' |
| 8:15.182 | ספירות | , |
| 8:15.183 | ודל"ת |  |
| 8:15.184 | כולל |  |
| 8:15.185 | הכל |  |
| 8:15.186 | להיות |  |
| 8:15.187 | אחד | , |
| 8:15.188 | ובכוונה |  |
| 8:15.189 | זאת |  |
| 8:15.190 | אנו |  |
| 8:15.191 | מאחדים |  |
| 8:15.192 | ואוגדים |  |
| 8:15.193 | את |  |
| 8:15.194 | כולם |  |
| 8:15.195 | להיות |  |
| 8:15.196 | אחד | , |
| 8:15.197 | מלכא |  |
| 8:15.198 | עלאה |  |
| 8:15.199 | אבא |  |
| 8:15.200 | ואימא | , |
| 8:15.201 | ומלכא |  |
| 8:15.202 | תתהה |  |
| 8:15.203 | זעיר |  |
| 8:15.204 | ונוקביה | , |
| 8:15.205 | לקשרא |  |
| 8:15.206 | דא |  |
| 8:15.207 | בדא |  |
| 8:15.208 | ולא |  |
| 8:15.209 | ישתכח |  |
| 8:15.210 | פירודא |  |
| 8:15.211 | אלא |  |
| 8:15.212 | כלא |  |
| 8:15.213 | חד |  |
| 8:15.214 | כדלעיל |  |
| 8:15.215 | סי | ' |
| 8:15.216 | מ"ב | . |
| 8:15.217 | הה | ! |
| 8:15.218 | אוי | ! |
| 8:15.219 | ואבוי | ! |
| 8:15.220 | על |  |
| 8:15.221 | רוע |  |
| 8:15.222 | מצב |  |
| 8:15.223 | אמונת |  |
| 8:15.224 | העם |  |
| 8:15.225 | אשר |  |
| 8:15.226 | בחר |  |
| 8:15.227 | לו |  |
| 8:15.228 | יה |  |
| 8:15.229 | ישראל |  |
| 8:15.230 | לסגולתו | . |
| 8:15.231 | ונתן |  |
| 8:15.232 | לנו |  |
| 8:15.233 | תורת |  |
| 8:15.234 | אמת | . |
| 8:15.235 | היא |  |
| 8:15.236 | חותמו |  |
| 8:15.237 | של |  |
| 8:15.238 | הקב"ה |  |
| 8:15.239 | אשר |  |
| 8:15.240 | קבל |  |
| 8:15.241 | משה |  |
| 8:15.242 | רבינו |  |
| 8:15.243 | עליו |  |
| 8:15.244 | השלום |  |
| 8:15.245 | בסיני |  |
| 8:15.246 | בכתב |  |
| 8:15.247 | ובעל |  |
| 8:15.248 | פה | . |
| 8:15.249 | ועמד |  |
| 8:15.250 | עליה |  |
| 8:15.251 | מ | ' |
| 8:15.252 | יום |  |
| 8:15.253 | וארבעים |  |
| 8:15.254 | לילה |  |
| 8:15.255 | ללמוד |  |
| 8:15.256 | דקדוקיה |  |
| 8:15.257 | וטעמיה |  |
| 8:15.258 | וכללותיה | , |
| 8:15.259 | ומסרה |  |
| 8:15.260 | ליהושע |  |
| 8:15.261 | וכו | ', |
| 8:15.262 | ואנשי |  |
| 8:15.263 | כנסת |  |
| 8:15.264 | הגדולה |  |
| 8:15.265 | תפסו |  |
| 8:15.266 | ליצרא |  |
| 8:15.267 | דע"ז |  |
| 8:15.268 | ושדויה |  |
| 8:15.269 | בדודא |  |
| 8:15.270 | דאברא |  |
| 8:15.271 | וחפיוה |  |
| 8:15.272 | לפומיה |  |
| 8:15.273 | באברא | . |
| 8:15.274 | הן |  |
| 8:15.275 | עתה |  |
| 8:15.276 | הוסר |  |
| 8:15.277 | כסוי |  |
| 8:15.278 | העופרת |  |
| 8:15.279 | מעל |  |
| 8:15.280 | פי |  |
| 8:15.281 | האיפה |  |
| 8:15.282 | על |  |
| 8:15.283 | יד |  |
| 8:15.284 | המלך |  |
| 8:15.285 | העכו"ם |  |
| 8:15.286 | במסרו |  |
| 8:15.287 | את |  |
| 8:15.288 | בפר |  |
| 8:15.289 | הזהר |  |
| 8:15.290 | לחכמי |  |
| 8:15.291 | טוליטולה | , |
| 8:15.292 | יוצא |  |
| 8:15.293 | השטן |  |
| 8:15.294 | וישב |  |
| 8:15.295 | וירקד |  |
| 8:15.296 | בנינו | , |
| 8:15.297 | ויהי |  |
| 8:15.298 | לריק |  |
| 8:15.299 | כל |  |
| 8:15.300 | יגיעם |  |
| 8:15.301 | של |  |
| 8:15.302 | אנשי |  |
| 8:15.303 | כנסת |  |
| 8:15.304 | הגדולה |  |
| 8:15.305 | אשר |  |
| 8:15.306 | טרחו |  |
| 8:15.307 | ויגעו | . |
| 8:15.308 | וצעקו |  |
| 8:15.309 | לפני |  |
| 8:15.310 | השי"ת |  |
| 8:15.311 | על |  |
| 8:15.312 | יצרא |  |
| 8:15.313 | דע"ז | . |
| 8:15.314 | ומן |  |
| 8:15.315 | השמיים |  |
| 8:15.316 | סייעו |  |
| 8:15.317 | אותם |  |
| 8:15.318 | ונמסר |  |
| 8:15.319 | בידם | . |
| 8:15.320 | ורוב |  |
| 8:15.321 | חכמתם |  |
| 8:15.322 | ורוח |  |
| 8:15.323 | בינתם |  |
| 8:15.324 | ביטלו |  |
| 8:15.325 | יצרא |  |
| 8:15.326 | דע"ז |  |
| 8:15.327 | על |  |
| 8:15.328 | ידי |  |
| 8:15.329 | עסק |  |
| 8:15.330 | התורה |  |
| 8:15.331 | וביאור |  |
| 8:15.332 | המבינים |  |
| 8:15.333 | לעם |  |
| 8:15.334 | בה |  |
| 8:15.335 | על |  |
| 8:15.336 | פי |  |
| 8:15.337 | האמת |  |
| 8:15.338 | והצדק |  |
| 8:15.339 | כמו |  |
| 8:15.340 | שהוא |  |
| 8:15.341 | ערוך |  |
| 8:15.342 | לפנינו |  |
| 8:15.343 | בב | ' |
| 8:15.344 | התלמודים |  |
| 8:15.345 | ובמדרשים |  |
| 8:15.346 | האמתיים | . |
| 8:15.347 | שג | ' |
| 8:15.348 | ויקראו |  |
| 8:15.349 | בו |  |
| 8:15.350 | לפני |  |
| 8:15.351 | הרחוב |  |
| 8:15.352 | מאור |  |
| 8:15.353 | הבקר |  |
| 8:15.354 | עד |  |
| 8:15.355 | מחצית |  |
| 8:15.356 | היום |  |
| 8:15.357 | נגד |  |
| 8:15.358 | האנשים |  |
| 8:15.359 | והנשים |  |
| 8:15.360 | המבינים | . |
| 8:15.361 | ונאמר |  |
| 8:15.362 | הלוים |  |
| 8:15.363 | מבינים |  |
| 8:15.364 | את |  |
| 8:15.365 | העם |  |
| 8:15.366 | לתורה |  |
| 8:15.367 | ויקראו |  |
| 8:15.368 | בס | ' |
| 8:15.369 | בתורת |  |
| 8:15.370 | האלהים |  |
| 8:15.371 | מפורש |  |
| 8:15.372 | ושום |  |
| 8:15.373 | שכל |  |
| 8:15.374 | ויבינו |  |
| 8:15.375 | במקרא | . |
| 8:15.376 | ועל |  |
| 8:15.377 | ידי |  |
| 8:15.378 | תקון |  |
| 8:15.379 | נוסח |  |
| 8:15.380 | התפלות |  |
| 8:15.381 | והברכות |  |
| 8:15.382 | אשר |  |
| 8:15.383 | תקנו |  |
| 8:15.384 | לכל |  |
| 8:15.385 | ישראל | . |
| 8:15.386 | ועל |  |
| 8:15.387 | ידי |  |
| 8:15.388 | השתדלותם |  |
| 8:15.389 | להעמיד |  |
| 8:15.390 | תלמידים |  |
| 8:15.391 | הרבה |  |
| 8:15.392 | כדאיתא |  |
| 8:15.393 | בפרקא |  |
| 8:15.394 | קמא |  |
| 8:15.395 | דאבות | ! |
| 8:15.396 | על |  |
| 8:15.397 | ידי |  |
| 8:15.398 | כל |  |
| 8:15.399 | אלה |  |
| 8:15.400 | עלתה |  |
| 8:15.401 | בידם |  |
| 8:15.402 | לעקור |  |
| 8:15.403 | יצרא |  |
| 8:15.404 | דע"ז | , |
| 8:15.405 | והחזירו |  |
| 8:15.406 | עטרה |  |
| 8:15.407 | לישנה | , |
| 8:15.408 | ונמחה |  |
| 8:15.409 | שם |  |
| 8:15.410 | האלילים |  |
| 8:15.411 | וכל |  |
| 8:15.412 | האמונות |  |
| 8:15.413 | הכוזבות |  |
| 8:15.414 | המביאות |  |
| 8:15.415 | לידי |  |
| 8:15.416 | מינות |  |
| 8:15.417 | מישראל | , |
| 8:15.418 | כדגרסינן |  |
| 8:15.419 | ביומא |  |
| 8:15.420 | פרק |  |
| 8:15.421 | בא |  |
| 8:15.422 | לו | , |
| 8:15.423 | ויקם |  |
| 8:15.424 | על |  |
| 8:15.425 | מעלה |  |
| 8:15.426 | הלוים |  |
| 8:15.427 | ישוע |  |
| 8:15.428 | ובני |  |
| 8:15.429 | וקדמיאל |  |
| 8:15.430 | שבניה |  |
| 8:15.431 | בני |  |
| 8:15.432 | שרביה |  |
| 8:15.433 | בני |  |
| 8:15.434 | כנני |  |
| 8:15.435 | ויזעקו |  |
| 8:15.436 | בקול |  |
| 8:15.437 | גדול |  |
| 8:15.438 | אל |  |
| 8:15.439 | ה | ' |
| 8:15.440 | אלהיהם | , |
| 8:15.441 | מאי |  |
| 8:15.442 | אמור |  |
| 8:15.443 | אמר |  |
| 8:15.444 | רב |  |
| 8:15.445 | יהודה |  |
| 8:15.446 | אמר |  |
| 8:15.447 | רב |  |
| 8:15.448 | ואי |  |
| 8:15.449 | תימא |  |
| 8:15.450 | ר | ' |
| 8:15.451 | יוחנן |  |
| 8:15.452 | בייא |  |
| 8:15.453 | בייא |  |
| 8:15.454 | היינו |  |
| 8:15.455 | האיי |  |
| 8:15.456 | דאחרביה |  |
| 8:15.457 | לבי |  |
| 8:15.458 | מקדשא |  |
| 8:15.459 | וקלייה |  |
| 8:15.460 | להיכלא |  |
| 8:15.461 | וקטלניהו |  |
| 8:15.462 | לצדיקיא |  |
| 8:15.463 | ואגלינהו |  |
| 8:15.464 | לישראל |  |
| 8:15.465 | מארעהון |  |
| 8:15.466 | ועדיין |  |
| 8:15.467 | מרקיד |  |
| 8:15.468 | בינא | , |
| 8:15.469 | כלום |  |
| 8:15.470 | יהבתיה |  |
| 8:15.471 | ניהלן |  |
| 8:15.472 | אלא |  |
| 8:15.473 | לקבולי |  |
| 8:15.474 | ביה |  |
| 8:15.475 | אגרא | . |
| 8:15.476 | לא |  |
| 8:15.477 | איהו |  |
| 8:15.478 | בעינן |  |
| 8:15.479 | ולא |  |
| 8:15.480 | אגריה | , |
| 8:15.481 | בעינן | , |
| 8:15.482 | נפל |  |
| 8:15.483 | להו |  |
| 8:15.484 | פתקא |  |
| 8:15.485 | מרקיעא |  |
| 8:15.486 | דכתיב |  |
| 8:15.487 | ביה |  |
| 8:15.488 | אמת | . |
| 8:15.489 | אמר |  |
| 8:15.490 | רבי |  |
| 8:15.491 | חנינה |  |
| 8:15.492 | שמע |  |
| 8:15.493 | מינה |  |
| 8:15.494 | חותמו |  |
| 8:15.495 | של |  |
| 8:15.496 | הקב"ה |  |
| 8:15.497 | אמת | , |
| 8:15.498 | יתיבו |  |
| 8:15.499 | בתעניתא |  |
| 8:15.500 | תלתה |  |
| 8:15.501 | יומין |  |
| 8:15.502 | ותלתא |  |
| 8:15.503 | לילי |  |
| 8:15.504 | מסריה |  |
| 8:15.505 | ניהלינו |  |
| 8:15.506 | נפק |  |
| 8:15.507 | אתא |  |
| 8:15.508 | כי |  |
| 8:15.509 | גוריא |  |
| 8:15.510 | דנורא |  |
| 8:15.511 | מבית |  |
| 8:15.512 | קדש |  |
| 8:15.513 | הקדשים | . |
| 8:15.514 | אמרו |  |
| 8:15.515 | להו |  |
| 8:15.516 | נביא |  |
| 8:15.517 | לישראל |  |
| 8:15.518 | היינו |  |
| 8:15.519 | הוא |  |
| 8:15.520 | יצרא |  |
| 8:15.521 | דע"ז |  |
| 8:15.522 | שנאמר |  |
| 8:15.523 | ויאמר |  |
| 8:15.524 | זאת |  |
| 8:15.525 | הרשעה |  |
| 8:15.526 | וגומר |  |
| 8:15.527 | בהדי |  |
| 8:15.528 | דתפסי |  |
| 8:15.529 | ליה | , |
| 8:15.530 | אשתמיט |  |
| 8:15.531 | ביניתא |  |
| 8:15.532 | ממזיה |  |
| 8:15.533 | ורמא |  |
| 8:15.534 | קלא | , |
| 8:15.535 | אזל |  |
| 8:15.536 | קליה |  |
| 8:15.537 | ארבע |  |
| 8:15.538 | מאה |  |
| 8:15.539 | פרסי |  |
| 8:15.540 | אמר |  |
| 8:15.541 | היכי |  |
| 8:15.542 | נעביד |  |
| 8:15.543 | ליה | , |
| 8:15.544 | דלמא |  |
| 8:15.545 | מרחמי |  |
| 8:15.546 | עליה |  |
| 8:15.547 | מן |  |
| 8:15.548 | שמיא | . |
| 8:15.549 | אמר |  |
| 8:15.550 | להו |  |
| 8:15.551 | נביא |  |
| 8:15.552 | שדויה |  |
| 8:15.553 | בדורא |  |
| 8:15.554 | דאברא |  |
| 8:15.555 | והפויה |  |
| 8:15.556 | לפומא |  |
| 8:15.557 | באברא |  |
| 8:15.558 | דשאיב |  |
| 8:15.559 | קלא | , |
| 8:15.560 | שנאמר |  |
| 8:15.561 | זאת |  |
| 8:15.562 | הרשעה |  |
| 8:15.563 | וישלך |  |
| 8:15.564 | אותה |  |
| 8:15.565 | אל |  |
| 8:15.566 | תוך |  |
| 8:15.567 | האיפה |  |
| 8:15.568 | וישלך |  |
| 8:15.569 | את |  |
| 8:15.570 | אבן |  |
| 8:15.571 | העופרת |  |
| 8:15.572 | אל |  |
| 8:15.573 | פיה | . |
| 8:15.574 | אמרי |  |
| 8:15.575 | הואיל |  |
| 8:15.576 | ועת |  |
| 8:15.577 | רצון |  |
| 8:15.578 | היא | , |
| 8:15.579 | נבעי |  |
| 8:15.580 | רחמי |  |
| 8:15.581 | על |  |
| 8:15.582 | יצרא |  |
| 8:15.583 | דעבירה | . |
| 8:15.584 | בעו |  |
| 8:15.585 | רחמי |  |
| 8:15.586 | ואמסר |  |
| 8:15.587 | בידיהו | , |
| 8:15.588 | אמר |  |
| 8:15.589 | להו |  |
| 8:15.590 | נביא |  |
| 8:15.591 | הזו |  |
| 8:15.592 | ודאי |  |
| 8:15.593 | קטלתון |  |
| 8:15.594 | לההוא |  |
| 8:15.595 | כליא |  |
| 8:15.596 | עלמא | , |
| 8:15.597 | חבשוה |  |
| 8:15.598 | תלתא |  |
| 8:15.599 | יומי | . |
| 8:15.600 | איבעי |  |
| 8:15.601 | ביעתא |  |
| 8:15.602 | בת |  |
| 8:15.603 | יומא |  |
| 8:15.604 | לחולה |  |
| 8:15.605 | בכל |  |
| 8:15.606 | ארץ |  |
| 8:15.607 | ישראל |  |
| 8:15.608 | ולא |  |
| 8:15.609 | אשתכח | . |
| 8:15.610 | אמרי |  |
| 8:15.611 | היכי |  |
| 8:15.612 | נעביד | , |
| 8:15.613 | ניקטליה |  |
| 8:15.614 | איחרב |  |
| 8:15.615 | כולי |  |
| 8:15.616 | עלמא | , |
| 8:15.617 | נבעי |  |
| 8:15.618 | רחמי |  |
| 8:15.619 | לפלגא | , |
| 8:15.620 | פלגא |  |
| 8:15.621 | מרקיעא |  |
| 8:15.622 | לא |  |
| 8:15.623 | יהבי | . |
| 8:15.624 | כחלינהו |  |
| 8:15.625 | לעיניה |  |
| 8:15.626 | ושבקוה | , |
| 8:15.627 | ואהני |  |
| 8:15.628 | דלא |  |
| 8:15.629 | תגרי |  |
| 8:15.630 | ביה |  |
| 8:15.631 | באיניש |  |
| 8:15.632 | בקרובתא |  |
| 8:15.633 | ע"כ | . |
| 8:15.634 | ועיין |  |
| 8:15.635 | עץ |  |
| 8:15.636 | יוסף | , |
| 8:15.637 | ובענף |  |
| 8:15.638 | שכתב |  |
| 8:15.639 | שהמאמר |  |
| 8:15.640 | תמוה |  |
| 8:15.641 | הוא |  |
| 8:15.642 | מאד | , |
| 8:15.643 | ויצרא |  |
| 8:15.644 | דע"ז |  |
| 8:15.645 | בבית |  |
| 8:15.646 | קדש |  |
| 8:15.647 | הקדשים |  |
| 8:15.648 | מאי |  |
| 8:15.649 | עבידתיה | ? |
| 8:15.650 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

