# Review pack: milhamot-hashem perek-2-14 (Milchamot ha-Shem 2:30–2:30)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-14`.

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
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
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
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
- A conditional whose main clause follows a Hebrew full stop (ואם ראיתי. הם
  מגלים): join with a comma and say so in a token note.
- המתפלספים: "would-be philosophers"; הנצרים = "the Christians" (spelling variant
  of הנוצרים, noted); העשיריות: "the tens" (the ten sefirot, noted).
- Nested quotations: double marks for the outer quotation, single marks for a
  quotation inside it, double again inside that.
- Printed list markers (א), ב)): keep as the author's own, "1)", "2)"; a token
  note.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
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


## Milchamot ha-Shem 2:30

Hebrew (Wikisource): ומה שכתבת שז"א בן לאבא לאמא והוא אב לעלולים שאחריו, ויש לו אחות נוקבא דיליה, הכל משל וכינוי בספירות. ומה שכתבת שאין לעבוד שום פרצוף, זה ידוע בפי כל המקובלים, ואין צריכים לאזהרתך. ומה שכתבת בשם הרמב"ם חמשה הן הנקראים מינים וכו' מי הוא האומר שאין שם אלוה, או שהם שנים, או שהוא בעל גוף, או שאינו לבדו האחד או שעובד זולתו כדי להיות מליץ, או האומר שמצא סמנים או האומר שנברא העולם יש מיש ח"ו. אלא שכבר עלתה טינא בלבך, ואתה מעיין הספרים הקדושים ולוקח לך הדברים כפשטן וסומך על הבנתך. ושכחת מה שאמר חכם החכמים, ואל בינתך אל תשען. ומפני שכבר הודעת לי שאתה מבין תתיר לי ספקותי אשר שלחתי אליך ואשר אני עתיד לשלוח. ומה שכתבת והבנת רבים וגדולים חכמי לב הפורשים ממנה ראשונים ואחרונים, ח"ו שיש שום חכם פורש ממנה. ואם יש פורש ממנה הוא עצמו מודה בפה מלא שהוא מחסרון ידיעה בה. ומה שכתבת כמו שכ' הריב"ש וחו"י, שקר אתה דובר עליהם. והם שני עדים כשרים להעיד עליך שאתה דובר שקרים. וז"ל הריב"ש בקיצור כי ר' פרץ הכהן לא היה מדבר ולא מחשיב באותן הספירות. גם שמע מפיו שר' שמשון מקינון היה אומר אני מתפלל לדעת זה התינוק כלומר להוציא מלב המקובלים שהם מתפללים פעם לספירה אחת ופעם לספירה אחרת. גם בתפלת י"ח יש להם בכל אחת ואחת כוונה לספירה ידועה. וכל זה הוא דבר זר מאוד בעיני מי שאינו מקובל כמותם. וחושבים שזה אמונת שניות. וששמע אחד מן המתפלספים מספר בגנות המקובלים והיה אומר הנצרים מאמיני השלוש והמקובלים מאמיני העשיריות. ושאל להחכם הישיש דון יוסף ן' שושן והוא היה חכם בתלמוד. וראה פילוסופיא והיה מקובל וחסיד וגדול ומדקדק במצות. ואמר לו איך אתם המקובלים מכוונים בברכה אחת לספירה ידועה ובברכה אחת לספירה אחרת. ועוד הכי יש אלוהות לספירות? שיתפלל אדם להם? וענו לו חלילה שתהיה התפלה כי אם לשי"ת עילת העילות אלא שמכוין המחשבה להמשיך השפע לאותה ספירה המתיחסת לאותו דבר שהוא מבקש עליו כמו שתאמר שבברכת על הצדיקים יכוין לספירה הנקראת חסד שהיא מדת רחמים וכו'. והביא לו משל על זה. ענה ואמר והנה טוב מאד. ואמר איני תוקע עצמי באותה חכמה אחר שלא קיבלתיה מפי חכם מקובל. ואם ראיתי באורים על סודות הרמב"ן. הם מגלים טפח ומכסים טפחים. ואמר שחכם אחד פירש הקבלה שלא כד' הרמב"ן. ולכן אמר שדברי החכם ההוא יהיו לו לבדו. וחתם דבריו ולזה אני אומר שאין לסמוך בדברים כאלו אלא מפי חכם מקובל.

| id | word | punctuation after |
|---|---|---|
| 2:30.1 | ומה |  |
| 2:30.2 | שכתבת |  |
| 2:30.3 | שז"א |  |
| 2:30.4 | בן |  |
| 2:30.5 | לאבא |  |
| 2:30.6 | לאמא |  |
| 2:30.7 | והוא |  |
| 2:30.8 | אב |  |
| 2:30.9 | לעלולים |  |
| 2:30.10 | שאחריו | , |
| 2:30.11 | ויש |  |
| 2:30.12 | לו |  |
| 2:30.13 | אחות |  |
| 2:30.14 | נוקבא |  |
| 2:30.15 | דיליה | , |
| 2:30.16 | הכל |  |
| 2:30.17 | משל |  |
| 2:30.18 | וכינוי |  |
| 2:30.19 | בספירות | . |
| 2:30.20 | ומה |  |
| 2:30.21 | שכתבת |  |
| 2:30.22 | שאין |  |
| 2:30.23 | לעבוד |  |
| 2:30.24 | שום |  |
| 2:30.25 | פרצוף | , |
| 2:30.26 | זה |  |
| 2:30.27 | ידוע |  |
| 2:30.28 | בפי |  |
| 2:30.29 | כל |  |
| 2:30.30 | המקובלים | , |
| 2:30.31 | ואין |  |
| 2:30.32 | צריכים |  |
| 2:30.33 | לאזהרתך | . |
| 2:30.34 | ומה |  |
| 2:30.35 | שכתבת |  |
| 2:30.36 | בשם |  |
| 2:30.37 | הרמב"ם |  |
| 2:30.38 | חמשה |  |
| 2:30.39 | הן |  |
| 2:30.40 | הנקראים |  |
| 2:30.41 | מינים |  |
| 2:30.42 | וכו | ' |
| 2:30.43 | מי |  |
| 2:30.44 | הוא |  |
| 2:30.45 | האומר |  |
| 2:30.46 | שאין |  |
| 2:30.47 | שם |  |
| 2:30.48 | אלוה | , |
| 2:30.49 | או |  |
| 2:30.50 | שהם |  |
| 2:30.51 | שנים | , |
| 2:30.52 | או |  |
| 2:30.53 | שהוא |  |
| 2:30.54 | בעל |  |
| 2:30.55 | גוף | , |
| 2:30.56 | או |  |
| 2:30.57 | שאינו |  |
| 2:30.58 | לבדו |  |
| 2:30.59 | האחד |  |
| 2:30.60 | או |  |
| 2:30.61 | שעובד |  |
| 2:30.62 | זולתו |  |
| 2:30.63 | כדי |  |
| 2:30.64 | להיות |  |
| 2:30.65 | מליץ | , |
| 2:30.66 | או |  |
| 2:30.67 | האומר |  |
| 2:30.68 | שמצא |  |
| 2:30.69 | סמנים |  |
| 2:30.70 | או |  |
| 2:30.71 | האומר |  |
| 2:30.72 | שנברא |  |
| 2:30.73 | העולם |  |
| 2:30.74 | יש |  |
| 2:30.75 | מיש |  |
| 2:30.76 | ח"ו | . |
| 2:30.77 | אלא |  |
| 2:30.78 | שכבר |  |
| 2:30.79 | עלתה |  |
| 2:30.80 | טינא |  |
| 2:30.81 | בלבך | , |
| 2:30.82 | ואתה |  |
| 2:30.83 | מעיין |  |
| 2:30.84 | הספרים |  |
| 2:30.85 | הקדושים |  |
| 2:30.86 | ולוקח |  |
| 2:30.87 | לך |  |
| 2:30.88 | הדברים |  |
| 2:30.89 | כפשטן |  |
| 2:30.90 | וסומך |  |
| 2:30.91 | על |  |
| 2:30.92 | הבנתך | . |
| 2:30.93 | ושכחת |  |
| 2:30.94 | מה |  |
| 2:30.95 | שאמר |  |
| 2:30.96 | חכם |  |
| 2:30.97 | החכמים | , |
| 2:30.98 | ואל |  |
| 2:30.99 | בינתך |  |
| 2:30.100 | אל |  |
| 2:30.101 | תשען | . |
| 2:30.102 | ומפני |  |
| 2:30.103 | שכבר |  |
| 2:30.104 | הודעת |  |
| 2:30.105 | לי |  |
| 2:30.106 | שאתה |  |
| 2:30.107 | מבין |  |
| 2:30.108 | תתיר |  |
| 2:30.109 | לי |  |
| 2:30.110 | ספקותי |  |
| 2:30.111 | אשר |  |
| 2:30.112 | שלחתי |  |
| 2:30.113 | אליך |  |
| 2:30.114 | ואשר |  |
| 2:30.115 | אני |  |
| 2:30.116 | עתיד |  |
| 2:30.117 | לשלוח | . |
| 2:30.118 | ומה |  |
| 2:30.119 | שכתבת |  |
| 2:30.120 | והבנת |  |
| 2:30.121 | רבים |  |
| 2:30.122 | וגדולים |  |
| 2:30.123 | חכמי |  |
| 2:30.124 | לב |  |
| 2:30.125 | הפורשים |  |
| 2:30.126 | ממנה |  |
| 2:30.127 | ראשונים |  |
| 2:30.128 | ואחרונים | , |
| 2:30.129 | ח"ו |  |
| 2:30.130 | שיש |  |
| 2:30.131 | שום |  |
| 2:30.132 | חכם |  |
| 2:30.133 | פורש |  |
| 2:30.134 | ממנה | . |
| 2:30.135 | ואם |  |
| 2:30.136 | יש |  |
| 2:30.137 | פורש |  |
| 2:30.138 | ממנה |  |
| 2:30.139 | הוא |  |
| 2:30.140 | עצמו |  |
| 2:30.141 | מודה |  |
| 2:30.142 | בפה |  |
| 2:30.143 | מלא |  |
| 2:30.144 | שהוא |  |
| 2:30.145 | מחסרון |  |
| 2:30.146 | ידיעה |  |
| 2:30.147 | בה | . |
| 2:30.148 | ומה |  |
| 2:30.149 | שכתבת |  |
| 2:30.150 | כמו |  |
| 2:30.151 | שכ | ' |
| 2:30.152 | הריב"ש |  |
| 2:30.153 | וחו"י | , |
| 2:30.154 | שקר |  |
| 2:30.155 | אתה |  |
| 2:30.156 | דובר |  |
| 2:30.157 | עליהם | . |
| 2:30.158 | והם |  |
| 2:30.159 | שני |  |
| 2:30.160 | עדים |  |
| 2:30.161 | כשרים |  |
| 2:30.162 | להעיד |  |
| 2:30.163 | עליך |  |
| 2:30.164 | שאתה |  |
| 2:30.165 | דובר |  |
| 2:30.166 | שקרים | . |
| 2:30.167 | וז"ל |  |
| 2:30.168 | הריב"ש |  |
| 2:30.169 | בקיצור |  |
| 2:30.170 | כי |  |
| 2:30.171 | ר | ' |
| 2:30.172 | פרץ |  |
| 2:30.173 | הכהן |  |
| 2:30.174 | לא |  |
| 2:30.175 | היה |  |
| 2:30.176 | מדבר |  |
| 2:30.177 | ולא |  |
| 2:30.178 | מחשיב |  |
| 2:30.179 | באותן |  |
| 2:30.180 | הספירות | . |
| 2:30.181 | גם |  |
| 2:30.182 | שמע |  |
| 2:30.183 | מפיו |  |
| 2:30.184 | שר | ' |
| 2:30.185 | שמשון |  |
| 2:30.186 | מקינון |  |
| 2:30.187 | היה |  |
| 2:30.188 | אומר |  |
| 2:30.189 | אני |  |
| 2:30.190 | מתפלל |  |
| 2:30.191 | לדעת |  |
| 2:30.192 | זה |  |
| 2:30.193 | התינוק |  |
| 2:30.194 | כלומר |  |
| 2:30.195 | להוציא |  |
| 2:30.196 | מלב |  |
| 2:30.197 | המקובלים |  |
| 2:30.198 | שהם |  |
| 2:30.199 | מתפללים |  |
| 2:30.200 | פעם |  |
| 2:30.201 | לספירה |  |
| 2:30.202 | אחת |  |
| 2:30.203 | ופעם |  |
| 2:30.204 | לספירה |  |
| 2:30.205 | אחרת | . |
| 2:30.206 | גם |  |
| 2:30.207 | בתפלת |  |
| 2:30.208 | י"ח |  |
| 2:30.209 | יש |  |
| 2:30.210 | להם |  |
| 2:30.211 | בכל |  |
| 2:30.212 | אחת |  |
| 2:30.213 | ואחת |  |
| 2:30.214 | כוונה |  |
| 2:30.215 | לספירה |  |
| 2:30.216 | ידועה | . |
| 2:30.217 | וכל |  |
| 2:30.218 | זה |  |
| 2:30.219 | הוא |  |
| 2:30.220 | דבר |  |
| 2:30.221 | זר |  |
| 2:30.222 | מאוד |  |
| 2:30.223 | בעיני |  |
| 2:30.224 | מי |  |
| 2:30.225 | שאינו |  |
| 2:30.226 | מקובל |  |
| 2:30.227 | כמותם | . |
| 2:30.228 | וחושבים |  |
| 2:30.229 | שזה |  |
| 2:30.230 | אמונת |  |
| 2:30.231 | שניות | . |
| 2:30.232 | וששמע |  |
| 2:30.233 | אחד |  |
| 2:30.234 | מן |  |
| 2:30.235 | המתפלספים |  |
| 2:30.236 | מספר |  |
| 2:30.237 | בגנות |  |
| 2:30.238 | המקובלים |  |
| 2:30.239 | והיה |  |
| 2:30.240 | אומר |  |
| 2:30.241 | הנצרים |  |
| 2:30.242 | מאמיני |  |
| 2:30.243 | השלוש |  |
| 2:30.244 | והמקובלים |  |
| 2:30.245 | מאמיני |  |
| 2:30.246 | העשיריות | . |
| 2:30.247 | ושאל |  |
| 2:30.248 | להחכם |  |
| 2:30.249 | הישיש |  |
| 2:30.250 | דון |  |
| 2:30.251 | יוסף |  |
| 2:30.252 | ן | ' |
| 2:30.253 | שושן |  |
| 2:30.254 | והוא |  |
| 2:30.255 | היה |  |
| 2:30.256 | חכם |  |
| 2:30.257 | בתלמוד | . |
| 2:30.258 | וראה |  |
| 2:30.259 | פילוסופיא |  |
| 2:30.260 | והיה |  |
| 2:30.261 | מקובל |  |
| 2:30.262 | וחסיד |  |
| 2:30.263 | וגדול |  |
| 2:30.264 | ומדקדק |  |
| 2:30.265 | במצות | . |
| 2:30.266 | ואמר |  |
| 2:30.267 | לו |  |
| 2:30.268 | איך |  |
| 2:30.269 | אתם |  |
| 2:30.270 | המקובלים |  |
| 2:30.271 | מכוונים |  |
| 2:30.272 | בברכה |  |
| 2:30.273 | אחת |  |
| 2:30.274 | לספירה |  |
| 2:30.275 | ידועה |  |
| 2:30.276 | ובברכה |  |
| 2:30.277 | אחת |  |
| 2:30.278 | לספירה |  |
| 2:30.279 | אחרת | . |
| 2:30.280 | ועוד |  |
| 2:30.281 | הכי |  |
| 2:30.282 | יש |  |
| 2:30.283 | אלוהות |  |
| 2:30.284 | לספירות | ? |
| 2:30.285 | שיתפלל |  |
| 2:30.286 | אדם |  |
| 2:30.287 | להם | ? |
| 2:30.288 | וענו |  |
| 2:30.289 | לו |  |
| 2:30.290 | חלילה |  |
| 2:30.291 | שתהיה |  |
| 2:30.292 | התפלה |  |
| 2:30.293 | כי |  |
| 2:30.294 | אם |  |
| 2:30.295 | לשי"ת |  |
| 2:30.296 | עילת |  |
| 2:30.297 | העילות |  |
| 2:30.298 | אלא |  |
| 2:30.299 | שמכוין |  |
| 2:30.300 | המחשבה |  |
| 2:30.301 | להמשיך |  |
| 2:30.302 | השפע |  |
| 2:30.303 | לאותה |  |
| 2:30.304 | ספירה |  |
| 2:30.305 | המתיחסת |  |
| 2:30.306 | לאותו |  |
| 2:30.307 | דבר |  |
| 2:30.308 | שהוא |  |
| 2:30.309 | מבקש |  |
| 2:30.310 | עליו |  |
| 2:30.311 | כמו |  |
| 2:30.312 | שתאמר |  |
| 2:30.313 | שבברכת |  |
| 2:30.314 | על |  |
| 2:30.315 | הצדיקים |  |
| 2:30.316 | יכוין |  |
| 2:30.317 | לספירה |  |
| 2:30.318 | הנקראת |  |
| 2:30.319 | חסד |  |
| 2:30.320 | שהיא |  |
| 2:30.321 | מדת |  |
| 2:30.322 | רחמים |  |
| 2:30.323 | וכו | '. |
| 2:30.324 | והביא |  |
| 2:30.325 | לו |  |
| 2:30.326 | משל |  |
| 2:30.327 | על |  |
| 2:30.328 | זה | . |
| 2:30.329 | ענה |  |
| 2:30.330 | ואמר |  |
| 2:30.331 | והנה |  |
| 2:30.332 | טוב |  |
| 2:30.333 | מאד | . |
| 2:30.334 | ואמר |  |
| 2:30.335 | איני |  |
| 2:30.336 | תוקע |  |
| 2:30.337 | עצמי |  |
| 2:30.338 | באותה |  |
| 2:30.339 | חכמה |  |
| 2:30.340 | אחר |  |
| 2:30.341 | שלא |  |
| 2:30.342 | קיבלתיה |  |
| 2:30.343 | מפי |  |
| 2:30.344 | חכם |  |
| 2:30.345 | מקובל | . |
| 2:30.346 | ואם |  |
| 2:30.347 | ראיתי |  |
| 2:30.348 | באורים |  |
| 2:30.349 | על |  |
| 2:30.350 | סודות |  |
| 2:30.351 | הרמב"ן | . |
| 2:30.352 | הם |  |
| 2:30.353 | מגלים |  |
| 2:30.354 | טפח |  |
| 2:30.355 | ומכסים |  |
| 2:30.356 | טפחים | . |
| 2:30.357 | ואמר |  |
| 2:30.358 | שחכם |  |
| 2:30.359 | אחד |  |
| 2:30.360 | פירש |  |
| 2:30.361 | הקבלה |  |
| 2:30.362 | שלא |  |
| 2:30.363 | כד | ' |
| 2:30.364 | הרמב"ן | . |
| 2:30.365 | ולכן |  |
| 2:30.366 | אמר |  |
| 2:30.367 | שדברי |  |
| 2:30.368 | החכם |  |
| 2:30.369 | ההוא |  |
| 2:30.370 | יהיו |  |
| 2:30.371 | לו |  |
| 2:30.372 | לבדו | . |
| 2:30.373 | וחתם |  |
| 2:30.374 | דבריו |  |
| 2:30.375 | ולזה |  |
| 2:30.376 | אני |  |
| 2:30.377 | אומר |  |
| 2:30.378 | שאין |  |
| 2:30.379 | לסמוך |  |
| 2:30.380 | בדברים |  |
| 2:30.381 | כאלו |  |
| 2:30.382 | אלא |  |
| 2:30.383 | מפי |  |
| 2:30.384 | חכם |  |
| 2:30.385 | מקובל | . |


### Translation under review (2:30)

English: As for what you wrote, that Ze'ir Anpin is the son of Abba and Imma, and that he is the father of the effects that come after him, and he has a sister, his Nukva; all of it is a parable and an allusion concerning the sefirot. As for what you wrote, that no partzuf is to be worshipped, this is known on the lips of all the Kabbalists, and they have no need of your warning. As for what you wrote in the name of the Rambam, that they are five, those called heretics, etc. Who is the one who says there is no God, or that there are two, or that He has a body, or that He is not alone, the One, or that he worships another in order to be an intercessor, or the one who says he found signs, or the one who says the world was created something from something, God forbid. But a grudge has already risen in your heart, and you read the holy books, take their words at their plain sense, and rely on your understanding. And you forgot what the wisest of the wise said: "and do not lean on your own understanding." Since you have already informed me that you understand, resolve my doubts, which I sent you and which I am about to send. As for what you wrote, that you understood that many and great men of wise heart, early and late, withdraw from it: God forbid that there be any sage who withdraws from it. And if there is one who withdraws from it, he himself admits openly that it is from his lack of knowledge of it. As for what you wrote, as the Ribash and the Chavot Yair wrote, you speak falsehood about them. And they are two valid witnesses to testify against you that you speak lies. And these are the words of the Ribash, in brief: "Rabbi Peretz ha-Kohen did not speak of those sefirot and did not attach importance to them. Also he heard from his mouth that Rabbi Shimshon of Kinon used to say: 'I pray to know this child,' that is, to remove from the hearts of the Kabbalists the idea that they pray at one time to one sefirah and at another time to another sefirah. Also in the Eighteen Benedictions they have, in each and every one, a specific intention toward a known sefirah. And all this is a very strange thing in the eyes of anyone who is not a Kabbalist like them, and they think it is the belief in duality. And he heard one of the would-be philosophers speaking in disparagement of the Kabbalists, and he would say: 'The Christians are believers in the Trinity, and the Kabbalists are believers in the tens.' And he asked the elderly sage Don Yosef ibn Shushan, who was a sage in the Talmud, who had studied philosophy and was a Kabbalist, pious, great, and scrupulous in the commandments. And he said to him: 'How is it that you Kabbalists direct one blessing to one known sefirah and another blessing to another sefirah. Furthermore, are there divinities for the sefirot? That a man should pray to them?' And they answered him: 'God forbid that the prayer be anything but to the Name, may He be blessed, the Cause of causes; but one directs the thought to draw the flow to that sefirah which relates to the thing he asks for, just as you would say that in the blessing for the righteous one should direct to the sefirah called Chesed, which is the attribute of mercy, etc.' And he brought him a parable on this. He answered and said: 'And behold, very good.' And he said: 'I do not commit myself to that wisdom, since I did not receive it from the mouth of a Kabbalist sage.' 'And if I have seen commentaries on the secrets of the Ramban, they reveal a handbreadth and cover handbreadths.' And he said that a certain sage explained the Kabbalah not as the Ramban holds, and therefore he said that the words of that sage shall be his alone. And he ended his words." Therefore I say that one must not rely on such words, except from the mouth of a Kabbalist sage.

Literal: and what you wrote that Ze'ir Anpin son of of Abba of Imma and he father of of effects that come after him and there is to him a sister his Nukva his all parable and allusion concerning the sefirot and what you wrote that no to worship any partzuf this is known on the lips of all the Kabbalists and not need for your warning and what you wrote in the name of the Rambam five they are the ones called heretics etc. who is he the one who says that there is not there God or that they are two or that He is possessor of a body or that not alone the One or that worships another in order to be intercessor or the one who says that he found signs or the one who says that was created the world something from something God forbid but already rose up a grudge in your heart and you read the books the holy and take for yourself the matters as their plain sense and rely on your understanding and you forgot what that said wise one of the wise and to your understanding not lean and since already you informed to me that you understand resolve for me my doubts which I sent to you and which I about to send and what you wrote and you understood many and great sages of heart the withdrawers from it the early and the late God forbid that there is any sage withdraws from it and if there is one who withdraws from it he himself admits with mouth full that it is from lack of knowledge of it and what you wrote as wrote the Ribash and the Chavot Yair falsehood you speak about them and they two witnesses valid to testify against you that you speak lies and these are the words of the Ribash in brief that Rabbi Peretz ha-Kohen not was speak and not attach importance to those the sefirot also heard from his mouth that Rabbi Shimshon of Kinon used to say I pray to know this the child that is to remove from the heart of the Kabbalists that they pray once to a sefirah one and once to a sefirah another also in the prayer of the eighteen there is for them in every one and one an intention toward a sefirah known and all this is a thing strange very in the eyes of anyone who is not a Kabbalist like them and they think that this the belief in duality and he heard one of the would-be philosophers speaking in disparagement of the Kabbalists and he would say the Christians believers in the Trinity and the Kabbalists believers in the tens and he asked the sage the elderly Don Yosef ibn Shushan who was a sage in the Talmud and had seen philosophy and was a Kabbalist pious great and scrupulous in the commandments and said to him how you the Kabbalists direct in a blessing one to a sefirah known and in a blessing one to a sefirah another and furthermore so is there divinities for sefirot that one should pray a man to them and they answered him far be it that there be the prayer but only to the Name, may He be blessed cause of the causes but that one directs the thought to draw the flow to that sefirah that relates to that matter which he asks for it as you would say that in the blessing of for the righteous one should direct to a sefirah called Chesed which is attribute of mercy etc. and brought him a parable on this he answered and said and behold good very and said I am not committing myself to that wisdom since not I received it from the mouth of a sage Kabbalist and if I saw in commentaries (lights) on the secrets of the Ramban they reveal a handbreadth and cover handbreadths and said that a sage one explained the Kabbalah not as the view of the Ramban therefore said that the words of the sage that shall be his alone and ended his words and therefore I say that not to rely on words such as these except from the mouth of a sage a Kabbalist

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:30.1 | and what |  |  |  |
| 2:30.2 | you wrote |  |  |  |
| 2:30.3 | that Ze'ir Anpin |  |  | Ze'ir Anpin: Aramaic for 'Small Face'; the Hebrew writes the initials ז"א. Ze'ir alone is 'Ze'ir'. |
| 2:30.4 | son of |  |  |  |
| 2:30.5 | of Abba |  |  | Abba: 'Father', the name of a partzuf. |
| 2:30.6 | of Imma |  |  | Imma: 'Mother', the name of a partzuf. |
| 2:30.7 | and he |  |  |  |
| 2:30.8 | father of |  |  |  |
| 2:30.9 | of effects |  |  | Cause and effect: the author's עלה and עלול. 'Effects' are what a cause produces. |
| 2:30.10 | that come after him |  |  |  |
| 2:30.11 | and there is |  |  |  |
| 2:30.12 | to him |  |  |  |
| 2:30.13 | a sister |  |  |  |
| 2:30.14 | his Nukva |  |  | Nukva: 'Female', the feminine partner of the partzuf (his Nukva). |
| 2:30.15 | his |  |  |  |
| 2:30.16 | all |  |  |  |
| 2:30.17 | parable |  |  |  |
| 2:30.18 | and allusion |  |  |  |
| 2:30.19 | concerning the sefirot |  |  |  |
| 2:30.20 | and what |  |  |  |
| 2:30.21 | you wrote |  |  |  |
| 2:30.22 | that no |  |  |  |
| 2:30.23 | to worship |  |  |  |
| 2:30.24 | any |  |  |  |
| 2:30.25 | partzuf |  |  | Partzuf: a Kabbalistic entity, kept transliterated (the ordinary sense is 'face'). |
| 2:30.26 | this |  |  |  |
| 2:30.27 | is known |  |  |  |
| 2:30.28 | on the lips of |  |  |  |
| 2:30.29 | all |  |  |  |
| 2:30.30 | the Kabbalists |  |  |  |
| 2:30.31 | and not |  |  |  |
| 2:30.32 | need |  |  |  |
| 2:30.33 | for your warning |  |  |  |
| 2:30.34 | and what |  |  |  |
| 2:30.35 | you wrote |  |  |  |
| 2:30.36 | in the name of |  |  |  |
| 2:30.37 | the Rambam |  |  | The Rambam (Maimonides). The list is the one the questioner gives in the Rambam's name. |
| 2:30.38 | five |  |  |  |
| 2:30.39 | they are |  |  |  |
| 2:30.40 | the ones called |  |  |  |
| 2:30.41 | heretics |  |  | Heretics (minim). |
| 2:30.42 | etc. |  |  | 'Etc.' marks a truncated quotation: the list goes on in the source, and the author does not give it. |
| 2:30.43 | who |  |  |  |
| 2:30.44 | is he |  |  |  |
| 2:30.45 | the one who says |  |  |  |
| 2:30.46 | that there is not |  |  |  |
| 2:30.47 | there |  |  |  |
| 2:30.48 | God |  |  |  |
| 2:30.49 | or |  |  |  |
| 2:30.50 | that they are |  |  |  |
| 2:30.51 | two |  |  |  |
| 2:30.52 | or |  |  |  |
| 2:30.53 | that He is |  |  |  |
| 2:30.54 | possessor of |  |  |  |
| 2:30.55 | a body |  |  |  |
| 2:30.56 | or |  |  |  |
| 2:30.57 | that not |  |  |  |
| 2:30.58 | alone |  |  |  |
| 2:30.59 | the One |  |  |  |
| 2:30.60 | or |  |  |  |
| 2:30.61 | that worships |  |  |  |
| 2:30.62 | another |  |  |  |
| 2:30.63 | in order |  |  |  |
| 2:30.64 | to be |  |  |  |
| 2:30.65 | intercessor |  |  |  |
| 2:30.66 | or |  |  |  |
| 2:30.67 | the one who says |  |  |  |
| 2:30.68 | that he found |  |  |  |
| 2:30.69 | signs |  |  | Literally 'signs'. |
| 2:30.70 | or |  |  |  |
| 2:30.71 | the one who says |  |  |  |
| 2:30.72 | that was created |  |  |  |
| 2:30.73 | the world |  |  |  |
| 2:30.74 | something |  |  |  |
| 2:30.75 | from something |  |  | Literally 'something from something' (יש מיש). |
| 2:30.76 | God forbid |  |  | God forbid (ח"ו). |
| 2:30.77 | but |  |  |  |
| 2:30.78 | already |  |  |  |
| 2:30.79 | rose up |  |  |  |
| 2:30.80 | a grudge |  |  |  |
| 2:30.81 | in your heart |  |  |  |
| 2:30.82 | and you |  |  |  |
| 2:30.83 | read |  |  |  |
| 2:30.84 | the books |  |  |  |
| 2:30.85 | the holy |  |  |  |
| 2:30.86 | and take |  |  |  |
| 2:30.87 | for yourself |  |  |  |
| 2:30.88 | the matters |  |  |  |
| 2:30.89 | as their plain sense |  |  | Literally 'according to its plain sense'; the Hebrew is in the plural, 'the books' plain sense'. |
| 2:30.90 | and rely |  |  |  |
| 2:30.91 | on |  |  |  |
| 2:30.92 | your understanding |  |  |  |
| 2:30.93 | and you forgot |  |  |  |
| 2:30.94 | what |  |  |  |
| 2:30.95 | that said |  |  | The verse is quoted in the following words. |
| 2:30.96 | wise one of |  |  | Literally 'the wise one of the wise' (חכם החכמים). |
| 2:30.97 | the wise |  |  |  |
| 2:30.98 | and to |  |  |  |
| 2:30.99 | your understanding |  |  |  |
| 2:30.100 | not |  |  |  |
| 2:30.101 | lean |  |  | Prov 3:5, 'and do not lean on your own understanding'. |
| 2:30.102 | and since |  |  |  |
| 2:30.103 | already |  |  |  |
| 2:30.104 | you informed |  |  |  |
| 2:30.105 | to me |  |  |  |
| 2:30.106 | that you |  |  |  |
| 2:30.107 | understand |  |  |  |
| 2:30.108 | resolve |  |  |  |
| 2:30.109 | for me |  |  |  |
| 2:30.110 | my doubts |  |  |  |
| 2:30.111 | which |  |  |  |
| 2:30.112 | I sent |  |  |  |
| 2:30.113 | to you |  |  |  |
| 2:30.114 | and which |  |  |  |
| 2:30.115 | I |  |  |  |
| 2:30.116 | about |  |  |  |
| 2:30.117 | to send |  |  |  |
| 2:30.118 | and what |  |  |  |
| 2:30.119 | you wrote |  |  |  |
| 2:30.120 | and you understood |  |  |  |
| 2:30.121 | many |  |  |  |
| 2:30.122 | and great |  |  |  |
| 2:30.123 | sages of |  |  |  |
| 2:30.124 | heart |  |  |  |
| 2:30.125 | the withdrawers |  |  |  |
| 2:30.126 | from it |  |  |  |
| 2:30.127 | the early |  |  |  |
| 2:30.128 | and the late |  |  |  |
| 2:30.129 | God forbid |  |  |  |
| 2:30.130 | that there is |  |  |  |
| 2:30.131 | any |  |  |  |
| 2:30.132 | sage |  |  |  |
| 2:30.133 | withdraws |  |  |  |
| 2:30.134 | from it |  |  |  |
| 2:30.135 | and if |  |  |  |
| 2:30.136 | there is |  |  |  |
| 2:30.137 | one who withdraws |  |  |  |
| 2:30.138 | from it |  |  |  |
| 2:30.139 | he |  |  |  |
| 2:30.140 | himself |  |  |  |
| 2:30.141 | admits |  |  |  |
| 2:30.142 | with mouth |  |  | Literally 'with a full mouth' (בפה מלא): openly, without reservation. |
| 2:30.143 | full |  |  |  |
| 2:30.144 | that it is |  |  |  |
| 2:30.145 | from lack of |  |  |  |
| 2:30.146 | knowledge |  |  |  |
| 2:30.147 | of it |  |  |  |
| 2:30.148 | and what |  |  |  |
| 2:30.149 | you wrote |  |  |  |
| 2:30.150 | as |  |  |  |
| 2:30.151 | wrote |  |  |  |
| 2:30.152 | the Ribash |  |  | The Ribash: Rabbi Yitzchak ben Sheshet. |
| 2:30.153 | and the Chavot Yair |  |  | Chavot Yair, the book the abbreviation names; cited by its title. |
| 2:30.154 | falsehood |  |  |  |
| 2:30.155 | you |  |  |  |
| 2:30.156 | speak |  |  |  |
| 2:30.157 | about them |  |  |  |
| 2:30.158 | and they |  |  |  |
| 2:30.159 | two |  |  |  |
| 2:30.160 | witnesses |  |  |  |
| 2:30.161 | valid |  |  |  |
| 2:30.162 | to testify |  |  |  |
| 2:30.163 | against you |  |  |  |
| 2:30.164 | that you |  |  |  |
| 2:30.165 | speak |  |  |  |
| 2:30.166 | lies |  |  |  |
| 2:30.167 | and these are the words of |  |  |  |
| 2:30.168 | the Ribash |  |  |  |
| 2:30.169 | in brief |  |  |  |
| 2:30.170 | that |  |  |  |
| 2:30.171 | Rabbi |  |  | ר' written in full: Rabbi. |
| 2:30.172 | Peretz |  |  |  |
| 2:30.173 | ha-Kohen |  |  |  |
| 2:30.174 | not |  |  |  |
| 2:30.175 | was |  |  |  |
| 2:30.176 | speak |  |  |  |
| 2:30.177 | and not |  |  |  |
| 2:30.178 | attach importance |  |  |  |
| 2:30.179 | to those |  |  |  |
| 2:30.180 | the sefirot |  |  |  |
| 2:30.181 | also |  |  |  |
| 2:30.182 | heard |  |  |  |
| 2:30.183 | from his mouth |  |  |  |
| 2:30.184 | that Rabbi |  |  | שר' is ש' ('that') with ר' ('Rabbi'). |
| 2:30.185 | Shimshon |  |  |  |
| 2:30.186 | of Kinon |  |  |  |
| 2:30.187 | used |  |  |  |
| 2:30.188 | to say |  |  |  |
| 2:30.189 | I |  |  |  |
| 2:30.190 | pray |  |  |  |
| 2:30.191 | to know |  |  |  |
| 2:30.192 | this |  |  |  |
| 2:30.193 | the child |  |  | 'The child': as written. The text does not explain the phrase. |
| 2:30.194 | that is |  |  |  |
| 2:30.195 | to remove |  |  |  |
| 2:30.196 | from the heart of |  |  |  |
| 2:30.197 | the Kabbalists |  |  |  |
| 2:30.198 | that they |  |  |  |
| 2:30.199 | pray |  |  |  |
| 2:30.200 | once |  |  |  |
| 2:30.201 | to a sefirah |  |  |  |
| 2:30.202 | one |  |  |  |
| 2:30.203 | and once |  |  |  |
| 2:30.204 | to a sefirah |  |  |  |
| 2:30.205 | another |  |  |  |
| 2:30.206 | also |  |  |  |
| 2:30.207 | in the prayer of |  |  |  |
| 2:30.208 | the eighteen |  |  | The Eighteen Benedictions (the Amidah). |
| 2:30.209 | there is |  |  |  |
| 2:30.210 | for them |  |  |  |
| 2:30.211 | in every |  |  |  |
| 2:30.212 | one |  |  |  |
| 2:30.213 | and one |  |  |  |
| 2:30.214 | an intention |  |  |  |
| 2:30.215 | toward a sefirah |  |  |  |
| 2:30.216 | known |  |  |  |
| 2:30.217 | and all |  |  |  |
| 2:30.218 | this |  |  |  |
| 2:30.219 | is |  |  |  |
| 2:30.220 | a thing |  |  |  |
| 2:30.221 | strange |  |  |  |
| 2:30.222 | very |  |  |  |
| 2:30.223 | in the eyes of |  |  |  |
| 2:30.224 | anyone |  |  |  |
| 2:30.225 | who is not |  |  |  |
| 2:30.226 | a Kabbalist |  |  |  |
| 2:30.227 | like them |  |  |  |
| 2:30.228 | and they think |  |  |  |
| 2:30.229 | that this |  |  |  |
| 2:30.230 | the belief in |  |  |  |
| 2:30.231 | duality |  |  |  |
| 2:30.232 | and he heard |  |  |  |
| 2:30.233 | one |  |  |  |
| 2:30.234 | of |  |  |  |
| 2:30.235 | the would-be philosophers |  |  | 'Would-be philosophers' (המתפלספים): a term of disparagement. |
| 2:30.236 | speaking |  |  |  |
| 2:30.237 | in disparagement of |  |  |  |
| 2:30.238 | the Kabbalists |  |  |  |
| 2:30.239 | and he |  |  |  |
| 2:30.240 | would say |  |  |  |
| 2:30.241 | the Christians |  |  | The Hebrew writes הנצרים (Notzrim), the term for the Christians. |
| 2:30.242 | believers in |  |  |  |
| 2:30.243 | the Trinity |  |  |  |
| 2:30.244 | and the Kabbalists |  |  |  |
| 2:30.245 | believers in |  |  |  |
| 2:30.246 | the tens |  |  | 'The tens': the ten sefirot. |
| 2:30.247 | and he asked |  |  |  |
| 2:30.248 | the sage |  |  |  |
| 2:30.249 | the elderly |  |  |  |
| 2:30.250 | Don |  |  |  |
| 2:30.251 | Yosef |  |  |  |
| 2:30.252 | ibn |  |  | ibn: 'son of'. |
| 2:30.253 | Shushan |  |  |  |
| 2:30.254 | who |  |  |  |
| 2:30.255 | was |  |  |  |
| 2:30.256 | a sage |  |  |  |
| 2:30.257 | in the Talmud |  |  |  |
| 2:30.258 | and had seen |  |  |  |
| 2:30.259 | philosophy |  |  |  |
| 2:30.260 | and was |  |  |  |
| 2:30.261 | a Kabbalist |  |  |  |
| 2:30.262 | pious |  |  |  |
| 2:30.263 | great |  |  |  |
| 2:30.264 | and scrupulous |  |  |  |
| 2:30.265 | in the commandments |  |  |  |
| 2:30.266 | and said |  |  |  |
| 2:30.267 | to him |  |  |  |
| 2:30.268 | how |  |  |  |
| 2:30.269 | you |  |  |  |
| 2:30.270 | the Kabbalists |  |  |  |
| 2:30.271 | direct |  |  |  |
| 2:30.272 | in a blessing |  |  |  |
| 2:30.273 | one |  |  |  |
| 2:30.274 | to a sefirah |  |  |  |
| 2:30.275 | known |  |  |  |
| 2:30.276 | and in a blessing |  |  |  |
| 2:30.277 | one |  |  |  |
| 2:30.278 | to a sefirah |  |  |  |
| 2:30.279 | another |  |  |  |
| 2:30.280 | and furthermore |  |  |  |
| 2:30.281 | so |  |  |  |
| 2:30.282 | is there |  |  |  |
| 2:30.283 | divinities |  |  |  |
| 2:30.284 | for sefirot |  |  |  |
| 2:30.285 | that one should pray |  |  |  |
| 2:30.286 | a man |  |  |  |
| 2:30.287 | to them |  |  |  |
| 2:30.288 | and they answered |  |  |  |
| 2:30.289 | him |  |  |  |
| 2:30.290 | far be it |  |  |  |
| 2:30.291 | that there be |  |  |  |
| 2:30.292 | the prayer |  |  |  |
| 2:30.293 | but |  |  |  |
| 2:30.294 | only |  |  |  |
| 2:30.295 | to the Name, may He be blessed |  |  | השי"ת: the Name, may He be blessed. |
| 2:30.296 | cause of |  |  | 'Cause of causes' (עלת העלות). |
| 2:30.297 | the causes |  |  |  |
| 2:30.298 | but |  |  |  |
| 2:30.299 | that one directs |  |  |  |
| 2:30.300 | the thought |  |  |  |
| 2:30.301 | to draw |  |  |  |
| 2:30.302 | the flow |  |  |  |
| 2:30.303 | to that |  |  |  |
| 2:30.304 | sefirah |  |  |  |
| 2:30.305 | that relates |  |  |  |
| 2:30.306 | to that |  |  |  |
| 2:30.307 | matter |  |  |  |
| 2:30.308 | which he |  |  |  |
| 2:30.309 | asks |  |  |  |
| 2:30.310 | for it |  |  |  |
| 2:30.311 | as |  |  |  |
| 2:30.312 | you would say |  |  |  |
| 2:30.313 | that in the blessing of |  |  |  |
| 2:30.314 | for |  |  |  |
| 2:30.315 | the righteous |  |  |  |
| 2:30.316 | one should direct |  |  |  |
| 2:30.317 | to a sefirah |  |  |  |
| 2:30.318 | called |  |  |  |
| 2:30.319 | Chesed |  |  |  |
| 2:30.320 | which is |  |  |  |
| 2:30.321 | attribute of |  |  |  |
| 2:30.322 | mercy |  |  |  |
| 2:30.323 | etc. |  |  |  |
| 2:30.324 | and brought |  |  |  |
| 2:30.325 | him |  |  |  |
| 2:30.326 | a parable |  |  |  |
| 2:30.327 | on |  |  |  |
| 2:30.328 | this |  |  |  |
| 2:30.329 | he answered |  |  |  |
| 2:30.330 | and said |  |  |  |
| 2:30.331 | and behold |  |  |  |
| 2:30.332 | good |  |  |  |
| 2:30.333 | very |  |  |  |
| 2:30.334 | and said |  |  |  |
| 2:30.335 | I am not |  |  |  |
| 2:30.336 | committing |  |  |  |
| 2:30.337 | myself |  |  |  |
| 2:30.338 | to that |  |  |  |
| 2:30.339 | wisdom |  |  |  |
| 2:30.340 | since |  |  |  |
| 2:30.341 | not |  |  |  |
| 2:30.342 | I received it |  |  |  |
| 2:30.343 | from the mouth of |  |  |  |
| 2:30.344 | a sage |  |  |  |
| 2:30.345 | Kabbalist |  |  |  |
| 2:30.346 | and if |  |  |  |
| 2:30.347 | I saw |  |  |  |
| 2:30.348 | in commentaries (lights) |  |  | Literally 'lights': commentaries called 'lights' on the Ramban's secrets. |
| 2:30.349 | on |  |  |  |
| 2:30.350 | the secrets of |  |  |  |
| 2:30.351 | the Ramban |  |  | The Ramban (Nachmanides). |
| 2:30.352 | they |  |  |  |
| 2:30.353 | reveal |  |  |  |
| 2:30.354 | a handbreadth |  |  |  |
| 2:30.355 | and cover |  |  |  |
| 2:30.356 | handbreadths |  |  | An idiom: they reveal a handbreadth and conceal handbreadths, that is, reveal little and hide much. |
| 2:30.357 | and said |  |  |  |
| 2:30.358 | that a sage |  |  |  |
| 2:30.359 | one |  |  |  |
| 2:30.360 | explained |  |  |  |
| 2:30.361 | the Kabbalah |  |  |  |
| 2:30.362 | not |  |  |  |
| 2:30.363 | as the view of |  |  | כד' is כדעת: 'as the view of'. |
| 2:30.364 | the Ramban |  |  |  |
| 2:30.365 | therefore |  |  |  |
| 2:30.366 | said |  |  |  |
| 2:30.367 | that the words of |  |  |  |
| 2:30.368 | the sage |  |  |  |
| 2:30.369 | that |  |  |  |
| 2:30.370 | shall be |  |  |  |
| 2:30.371 | his |  |  |  |
| 2:30.372 | alone |  |  |  |
| 2:30.373 | and ended |  |  |  |
| 2:30.374 | his words |  |  |  |
| 2:30.375 | and therefore |  |  |  |
| 2:30.376 | I |  |  |  |
| 2:30.377 | say |  |  |  |
| 2:30.378 | that not |  |  |  |
| 2:30.379 | to rely |  |  |  |
| 2:30.380 | on words |  |  |  |
| 2:30.381 | such as these |  |  |  |
| 2:30.382 | except |  |  |  |
| 2:30.383 | from the mouth of |  |  |  |
| 2:30.384 | a sage |  |  |  |
| 2:30.385 | a Kabbalist |  |  |  |

Phrases:
- 2:30.1 2:30.2 2:30.3 2:30.4 2:30.5 2:30.6 2:30.7 2:30.8 2:30.9 2:30.10 | As for what you wrote, that Ze'ir Anpin is the son of Abba and Imma, and that he is the father of the effects that come after him, | and what you wrote that Ze'ir Anpin son of of Abba of Imma and he father of of effects that come after him
- 2:30.11 2:30.12 2:30.13 2:30.14 2:30.15 2:30.16 2:30.17 2:30.18 2:30.19 | and he has a sister, his Nukva; all of it is a parable and an allusion concerning the sefirot. | and there is to him a sister his Nukva his all parable and allusion concerning the sefirot
- 2:30.20 2:30.21 2:30.22 2:30.23 2:30.24 2:30.25 2:30.26 2:30.27 2:30.28 2:30.29 2:30.30 2:30.31 2:30.32 2:30.33 | As for what you wrote, that no partzuf is to be worshipped, this is known on the lips of all the Kabbalists, and they have no need of your warning. | and what you wrote that no to worship any partzuf this is known on the lips of all the Kabbalists and not need for your warning
- 2:30.34 2:30.35 2:30.36 2:30.37 2:30.38 2:30.39 2:30.40 2:30.41 2:30.42 | As for what you wrote in the name of the Rambam, that they are five, those called heretics, etc. | and what you wrote in the name of the Rambam five they are the ones called heretics etc.
- 2:30.43 2:30.44 2:30.45 2:30.46 2:30.47 2:30.48 2:30.49 2:30.50 2:30.51 2:30.52 2:30.53 2:30.54 2:30.55 2:30.56 2:30.57 2:30.58 2:30.59 2:30.60 2:30.61 2:30.62 2:30.63 2:30.64 2:30.65 2:30.66 2:30.67 2:30.68 2:30.69 2:30.70 2:30.71 2:30.72 2:30.73 2:30.74 2:30.75 2:30.76 | Who is the one who says there is no God, or that there are two, or that He has a body, or that He is not alone, the One, or that he worships another in order to be an intercessor, or the one who says he found signs, or the one who says the world was created something from something, God forbid. | who is he the one who says that there is not there God or that they are two or that He is possessor of a body or that not alone the One or that worships another in order to be intercessor or the one who says that he found signs or the one who says that was created the world something from something God forbid
- 2:30.77 2:30.78 2:30.79 2:30.80 2:30.81 2:30.82 2:30.83 2:30.84 2:30.85 2:30.86 2:30.87 2:30.88 2:30.89 2:30.90 2:30.91 2:30.92 | But a grudge has already risen in your heart, and you read the holy books, take their words at their plain sense, and rely on your understanding. | but already rose up a grudge in your heart and you read the books the holy and take for yourself the matters as their plain sense and rely on your understanding
- 2:30.93 2:30.94 2:30.95 2:30.96 2:30.97 2:30.98 2:30.99 2:30.100 2:30.101 | And you forgot what the wisest of the wise said: "and do not lean on your own understanding." | and you forgot what that said wise one of the wise and to your understanding not lean
- 2:30.102 2:30.103 2:30.104 2:30.105 2:30.106 2:30.107 2:30.108 2:30.109 2:30.110 2:30.111 2:30.112 2:30.113 2:30.114 2:30.115 2:30.116 2:30.117 | Since you have already informed me that you understand, resolve my doubts, which I sent you and which I am about to send. | and since already you informed to me that you understand resolve for me my doubts which I sent to you and which I about to send
- 2:30.118 2:30.119 2:30.120 2:30.121 2:30.122 2:30.123 2:30.124 2:30.125 2:30.126 2:30.127 2:30.128 2:30.129 2:30.130 2:30.131 2:30.132 2:30.133 2:30.134 | As for what you wrote, that you understood that many and great men of wise heart, early and late, withdraw from it: God forbid that there be any sage who withdraws from it. | and what you wrote and you understood many and great sages of heart the withdrawers from it the early and the late God forbid that there is any sage withdraws from it
- 2:30.135 2:30.136 2:30.137 2:30.138 2:30.139 2:30.140 2:30.141 2:30.142 2:30.143 2:30.144 2:30.145 2:30.146 2:30.147 | And if there is one who withdraws from it, he himself admits openly that it is from his lack of knowledge of it. | and if there is one who withdraws from it he himself admits with mouth full that it is from lack of knowledge of it
- 2:30.148 2:30.149 2:30.150 2:30.151 2:30.152 2:30.153 2:30.154 2:30.155 2:30.156 2:30.157 | As for what you wrote, as the Ribash and the Chavot Yair wrote, you speak falsehood about them. | and what you wrote as wrote the Ribash and the Chavot Yair falsehood you speak about them
- 2:30.158 2:30.159 2:30.160 2:30.161 2:30.162 2:30.163 2:30.164 2:30.165 2:30.166 | And they are two valid witnesses to testify against you that you speak lies. | and they two witnesses valid to testify against you that you speak lies
- 2:30.167 2:30.168 2:30.169 2:30.170 2:30.171 2:30.172 2:30.173 2:30.174 2:30.175 2:30.176 2:30.177 2:30.178 2:30.179 2:30.180 | And these are the words of the Ribash, in brief: "Rabbi Peretz ha-Kohen did not speak of those sefirot and did not attach importance to them. | and these are the words of the Ribash in brief that Rabbi Peretz ha-Kohen not was speak and not attach importance to those the sefirot
- 2:30.181 2:30.182 2:30.183 2:30.184 2:30.185 2:30.186 2:30.187 2:30.188 2:30.189 2:30.190 2:30.191 2:30.192 2:30.193 2:30.194 2:30.195 2:30.196 2:30.197 2:30.198 2:30.199 2:30.200 2:30.201 2:30.202 2:30.203 2:30.204 2:30.205 | Also he heard from his mouth that Rabbi Shimshon of Kinon used to say: 'I pray to know this child,' that is, to remove from the hearts of the Kabbalists the idea that they pray at one time to one sefirah and at another time to another sefirah. | also heard from his mouth that Rabbi Shimshon of Kinon used to say I pray to know this the child that is to remove from the heart of the Kabbalists that they pray once to a sefirah one and once to a sefirah another
- 2:30.206 2:30.207 2:30.208 2:30.209 2:30.210 2:30.211 2:30.212 2:30.213 2:30.214 2:30.215 2:30.216 | Also in the Eighteen Benedictions they have, in each and every one, a specific intention toward a known sefirah. | also in the prayer of the eighteen there is for them in every one and one an intention toward a sefirah known
- 2:30.217 2:30.218 2:30.219 2:30.220 2:30.221 2:30.222 2:30.223 2:30.224 2:30.225 2:30.226 2:30.227 2:30.228 2:30.229 2:30.230 2:30.231 | And all this is a very strange thing in the eyes of anyone who is not a Kabbalist like them, and they think it is the belief in duality. | and all this is a thing strange very in the eyes of anyone who is not a Kabbalist like them and they think that this the belief in duality
- 2:30.232 2:30.233 2:30.234 2:30.235 2:30.236 2:30.237 2:30.238 2:30.239 2:30.240 | And he heard one of the would-be philosophers speaking in disparagement of the Kabbalists, and he would say: | and he heard one of the would-be philosophers speaking in disparagement of the Kabbalists and he would say
- 2:30.241 2:30.242 2:30.243 2:30.244 2:30.245 2:30.246 | 'The Christians are believers in the Trinity, and the Kabbalists are believers in the tens.' | the Christians believers in the Trinity and the Kabbalists believers in the tens
- 2:30.247 2:30.248 2:30.249 2:30.250 2:30.251 2:30.252 2:30.253 2:30.254 2:30.255 2:30.256 2:30.257 | And he asked the elderly sage Don Yosef ibn Shushan, who was a sage in the Talmud, | and he asked the sage the elderly Don Yosef ibn Shushan who was a sage in the Talmud
- 2:30.258 2:30.259 2:30.260 2:30.261 2:30.262 2:30.263 2:30.264 2:30.265 | who had studied philosophy and was a Kabbalist, pious, great, and scrupulous in the commandments. | and had seen philosophy and was a Kabbalist pious great and scrupulous in the commandments
- 2:30.266 2:30.267 2:30.268 2:30.269 2:30.270 2:30.271 2:30.272 2:30.273 2:30.274 2:30.275 2:30.276 2:30.277 2:30.278 2:30.279 | And he said to him: 'How is it that you Kabbalists direct one blessing to one known sefirah and another blessing to another sefirah. | and said to him how you the Kabbalists direct in a blessing one to a sefirah known and in a blessing one to a sefirah another
- 2:30.280 2:30.281 2:30.282 2:30.283 2:30.284 2:30.285 2:30.286 2:30.287 | Furthermore, are there divinities for the sefirot? That a man should pray to them?' | and furthermore so is there divinities for sefirot that one should pray a man to them
- 2:30.288 2:30.289 2:30.290 2:30.291 2:30.292 2:30.293 2:30.294 2:30.295 2:30.296 2:30.297 | And they answered him: 'God forbid that the prayer be anything but to the Name, may He be blessed, the Cause of causes; | and they answered him far be it that there be the prayer but only to the Name, may He be blessed cause of the causes
- 2:30.298 2:30.299 2:30.300 2:30.301 2:30.302 2:30.303 2:30.304 2:30.305 2:30.306 2:30.307 2:30.308 2:30.309 2:30.310 | but one directs the thought to draw the flow to that sefirah which relates to the thing he asks for, | but that one directs the thought to draw the flow to that sefirah that relates to that matter which he asks for it
- 2:30.311 2:30.312 2:30.313 2:30.314 2:30.315 2:30.316 2:30.317 2:30.318 2:30.319 2:30.320 2:30.321 2:30.322 2:30.323 | just as you would say that in the blessing for the righteous one should direct to the sefirah called Chesed, which is the attribute of mercy, etc.' | as you would say that in the blessing of for the righteous one should direct to a sefirah called Chesed which is attribute of mercy etc.
- 2:30.324 2:30.325 2:30.326 2:30.327 2:30.328 | And he brought him a parable on this. | and brought him a parable on this
- 2:30.329 2:30.330 2:30.331 2:30.332 2:30.333 | He answered and said: 'And behold, very good.' | he answered and said and behold good very
- 2:30.334 2:30.335 2:30.336 2:30.337 2:30.338 2:30.339 2:30.340 2:30.341 2:30.342 2:30.343 2:30.344 2:30.345 | And he said: 'I do not commit myself to that wisdom, since I did not receive it from the mouth of a Kabbalist sage.' | and said I am not committing myself to that wisdom since not I received it from the mouth of a sage Kabbalist
- 2:30.346 2:30.347 2:30.348 2:30.349 2:30.350 2:30.351 2:30.352 2:30.353 2:30.354 2:30.355 2:30.356 | 'And if I have seen commentaries on the secrets of the Ramban, they reveal a handbreadth and cover handbreadths.' | and if I saw in commentaries (lights) on the secrets of the Ramban they reveal a handbreadth and cover handbreadths
- 2:30.357 2:30.358 2:30.359 2:30.360 2:30.361 2:30.362 2:30.363 2:30.364 2:30.365 2:30.366 | And he said that a certain sage explained the Kabbalah not as the Ramban holds, and therefore he said | and said that a sage one explained the Kabbalah not as the view of the Ramban therefore said
- 2:30.367 2:30.368 2:30.369 2:30.370 2:30.371 2:30.372 2:30.373 2:30.374 | that the words of that sage shall be his alone. And he ended his words." | that the words of the sage that shall be his alone and ended his words
- 2:30.375 2:30.376 2:30.377 2:30.378 2:30.379 2:30.380 2:30.381 2:30.382 2:30.383 2:30.384 2:30.385 | Therefore I say that one must not rely on such words, except from the mouth of a Kabbalist sage. | and therefore I say that not to rely on words such as these except from the mouth of a sage a Kabbalist

Grammar notes:
- This paragraph is the respondent's letter, answering the questioner in the second person. Each point opens with ומה שכתבת ('as for what you wrote') and is answered in turn.
- The Ribash passage opens with the formula וז"ל and is given in double quotation marks. The speakers inside it are in single quotation marks.
- ואל בינתך אל תשען (Prov 3:5) is a verse quoted as a saying. תשען is the imperfect with אל, a negative command: 'do not lean'.
- שאין לעבוד is impersonal: 'one may not worship', rendered 'that no partzuf is to be worshipped'.
- היה with the participle (היה אומר) gives a habitual past: 'used to say'.

Translator's note: The paragraph is in the respondent's voice. The quotation from the Ribash is given as the writer gives it; the sentence in the Ribash's own voice that closes with וחתם דבריו ends the quotation. 'The Christians' renders הנצרים and 'something from something' renders יש מיש, both by the ruling. The Kabbalists' answer in the story is kept as Don Yosef's account gives it.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated (Q2).
- apparatus: Prov 3:5: 'ואל בינתך אל תשען' ('and do not lean on your own understanding'), quoted as the sage's words.
- apparatus: The text writes הנצרים (Notzrim) for the Christians.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

