# Translation pack: milhamot-hashem perek-4-17 (Milchamot ha-Shem 4:28–4:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-17`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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
- ר"א: "Rabbi Eleazar" only where the paragraph names him, else "R. A." with
  a note.
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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
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
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 4:28

Hebrew (Wikisource): לה) ומזה יתבאר לך ידידי אמתת אמונתינו הטהורה המזוקקה על פי רבותינו ז"ל מקבלי השמועה ממשה רבי' ע"ה שהשם ית' שהוא הסבה הראשונה לבדו ברא את הכל בלי סיוע שום נברא דלא כמו שדרש הפילוסוף מחבר הזהר שהבאתי לעיל סי' כ' וכ"א וכ"ד כי לפי דבריו העלות הרבות שהמציא מסייעות זו את זו. וכל אחת נוטלת רשות מסבה שהיא למעלה ממנה. ויש מקום לבעל הדין לומר עתיק ממתח בדרום ואריך אנפין בצפון ואבא במזרח ואימא במערב וזעיר ואשתו בזויות ואדם קדמון הגדול שבהם ממדד באמצע. ונגד מה שאמרו רז"ל בנוהג שבעולם מלך ב"ו מתקלס במדינה וגדולי המדינה מתקלסין עמו לפי שהם נושאים עמו במשאו וכו'. כן עשה הפילוסוף נביא השקר מחבר הזהר ושאר מקובלים בכבוד אבינו מלכינו שכל הפרצופים שהמציאו ברעיוניהם הם מתקלסים עמו! אדם קדמון הגדול שבהם מתקלס במאמר ראו עתה כי אני אני הוא אני אמית ואחיה וכו' שאין עילה אחרת עליו שיטול רשות ממנה! ואבא מתקלס במאמר יהי אור ויקוו המים כדלעיל! ואימא מתקלסת במאמר נעשה אדם! ועתיק מתקלס במאמר ה' פעלך בקרב שנים וחייהו. דלעתיקא קדישא אתמר כמבואר באדרת נשא! וזעיר מתקלס במאמר אנכי ה' אלהיך אשר הוצאתיך וכו'. כמו שכ' ספר הברית (יובא לקמן סי' ט"ז) והרב נחלת יוסף דף ס"א וס"ב! ומלכות מתקלסת במאמר אלה הם האלוהים המכים את מצרים שהמלכות היתה מכה בהם בסיועת הבינה כמו שכ' הר' נחלת יוסף ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:28.1 | לה | ) |
| 4:28.2 | ומזה |  |
| 4:28.3 | יתבאר |  |
| 4:28.4 | לך |  |
| 4:28.5 | ידידי |  |
| 4:28.6 | אמתת |  |
| 4:28.7 | אמונתינו |  |
| 4:28.8 | הטהורה |  |
| 4:28.9 | המזוקקה |  |
| 4:28.10 | על |  |
| 4:28.11 | פי |  |
| 4:28.12 | רבותינו |  |
| 4:28.13 | ז"ל |  |
| 4:28.14 | מקבלי |  |
| 4:28.15 | השמועה |  |
| 4:28.16 | ממשה |  |
| 4:28.17 | רבי | ' |
| 4:28.18 | ע"ה |  |
| 4:28.19 | שהשם |  |
| 4:28.20 | ית | ' |
| 4:28.21 | שהוא |  |
| 4:28.22 | הסבה |  |
| 4:28.23 | הראשונה |  |
| 4:28.24 | לבדו |  |
| 4:28.25 | ברא |  |
| 4:28.26 | את |  |
| 4:28.27 | הכל |  |
| 4:28.28 | בלי |  |
| 4:28.29 | סיוע |  |
| 4:28.30 | שום |  |
| 4:28.31 | נברא |  |
| 4:28.32 | דלא |  |
| 4:28.33 | כמו |  |
| 4:28.34 | שדרש |  |
| 4:28.35 | הפילוסוף |  |
| 4:28.36 | מחבר |  |
| 4:28.37 | הזהר |  |
| 4:28.38 | שהבאתי |  |
| 4:28.39 | לעיל |  |
| 4:28.40 | סי | ' |
| 4:28.41 | כ | ' |
| 4:28.42 | וכ"א |  |
| 4:28.43 | וכ"ד |  |
| 4:28.44 | כי |  |
| 4:28.45 | לפי |  |
| 4:28.46 | דבריו |  |
| 4:28.47 | העלות |  |
| 4:28.48 | הרבות |  |
| 4:28.49 | שהמציא |  |
| 4:28.50 | מסייעות |  |
| 4:28.51 | זו |  |
| 4:28.52 | את |  |
| 4:28.53 | זו | . |
| 4:28.54 | וכל |  |
| 4:28.55 | אחת |  |
| 4:28.56 | נוטלת |  |
| 4:28.57 | רשות |  |
| 4:28.58 | מסבה |  |
| 4:28.59 | שהיא |  |
| 4:28.60 | למעלה |  |
| 4:28.61 | ממנה | . |
| 4:28.62 | ויש |  |
| 4:28.63 | מקום |  |
| 4:28.64 | לבעל |  |
| 4:28.65 | הדין |  |
| 4:28.66 | לומר |  |
| 4:28.67 | עתיק |  |
| 4:28.68 | ממתח |  |
| 4:28.69 | בדרום |  |
| 4:28.70 | ואריך |  |
| 4:28.71 | אנפין |  |
| 4:28.72 | בצפון |  |
| 4:28.73 | ואבא |  |
| 4:28.74 | במזרח |  |
| 4:28.75 | ואימא |  |
| 4:28.76 | במערב |  |
| 4:28.77 | וזעיר |  |
| 4:28.78 | ואשתו |  |
| 4:28.79 | בזויות |  |
| 4:28.80 | ואדם |  |
| 4:28.81 | קדמון |  |
| 4:28.82 | הגדול |  |
| 4:28.83 | שבהם |  |
| 4:28.84 | ממדד |  |
| 4:28.85 | באמצע | . |
| 4:28.86 | ונגד |  |
| 4:28.87 | מה |  |
| 4:28.88 | שאמרו |  |
| 4:28.89 | רז"ל |  |
| 4:28.90 | בנוהג |  |
| 4:28.91 | שבעולם |  |
| 4:28.92 | מלך |  |
| 4:28.93 | ב"ו |  |
| 4:28.94 | מתקלס |  |
| 4:28.95 | במדינה |  |
| 4:28.96 | וגדולי |  |
| 4:28.97 | המדינה |  |
| 4:28.98 | מתקלסין |  |
| 4:28.99 | עמו |  |
| 4:28.100 | לפי |  |
| 4:28.101 | שהם |  |
| 4:28.102 | נושאים |  |
| 4:28.103 | עמו |  |
| 4:28.104 | במשאו |  |
| 4:28.105 | וכו | '. |
| 4:28.106 | כן |  |
| 4:28.107 | עשה |  |
| 4:28.108 | הפילוסוף |  |
| 4:28.109 | נביא |  |
| 4:28.110 | השקר |  |
| 4:28.111 | מחבר |  |
| 4:28.112 | הזהר |  |
| 4:28.113 | ושאר |  |
| 4:28.114 | מקובלים |  |
| 4:28.115 | בכבוד |  |
| 4:28.116 | אבינו |  |
| 4:28.117 | מלכינו |  |
| 4:28.118 | שכל |  |
| 4:28.119 | הפרצופים |  |
| 4:28.120 | שהמציאו |  |
| 4:28.121 | ברעיוניהם |  |
| 4:28.122 | הם |  |
| 4:28.123 | מתקלסים |  |
| 4:28.124 | עמו | ! |
| 4:28.125 | אדם |  |
| 4:28.126 | קדמון |  |
| 4:28.127 | הגדול |  |
| 4:28.128 | שבהם |  |
| 4:28.129 | מתקלס |  |
| 4:28.130 | במאמר |  |
| 4:28.131 | ראו |  |
| 4:28.132 | עתה |  |
| 4:28.133 | כי |  |
| 4:28.134 | אני |  |
| 4:28.135 | אני |  |
| 4:28.136 | הוא |  |
| 4:28.137 | אני |  |
| 4:28.138 | אמית |  |
| 4:28.139 | ואחיה |  |
| 4:28.140 | וכו | ' |
| 4:28.141 | שאין |  |
| 4:28.142 | עילה |  |
| 4:28.143 | אחרת |  |
| 4:28.144 | עליו |  |
| 4:28.145 | שיטול |  |
| 4:28.146 | רשות |  |
| 4:28.147 | ממנה | ! |
| 4:28.148 | ואבא |  |
| 4:28.149 | מתקלס |  |
| 4:28.150 | במאמר |  |
| 4:28.151 | יהי |  |
| 4:28.152 | אור |  |
| 4:28.153 | ויקוו |  |
| 4:28.154 | המים |  |
| 4:28.155 | כדלעיל | ! |
| 4:28.156 | ואימא |  |
| 4:28.157 | מתקלסת |  |
| 4:28.158 | במאמר |  |
| 4:28.159 | נעשה |  |
| 4:28.160 | אדם | ! |
| 4:28.161 | ועתיק |  |
| 4:28.162 | מתקלס |  |
| 4:28.163 | במאמר |  |
| 4:28.164 | ה | ' |
| 4:28.165 | פעלך |  |
| 4:28.166 | בקרב |  |
| 4:28.167 | שנים |  |
| 4:28.168 | וחייהו | . |
| 4:28.169 | דלעתיקא |  |
| 4:28.170 | קדישא |  |
| 4:28.171 | אתמר |  |
| 4:28.172 | כמבואר |  |
| 4:28.173 | באדרת |  |
| 4:28.174 | נשא | ! |
| 4:28.175 | וזעיר |  |
| 4:28.176 | מתקלס |  |
| 4:28.177 | במאמר |  |
| 4:28.178 | אנכי |  |
| 4:28.179 | ה | ' |
| 4:28.180 | אלהיך |  |
| 4:28.181 | אשר |  |
| 4:28.182 | הוצאתיך |  |
| 4:28.183 | וכו | '. |
| 4:28.184 | כמו |  |
| 4:28.185 | שכ | ' |
| 4:28.186 | ספר |  |
| 4:28.187 | הברית |  |
| 4:28.188 | יובא |  |
| 4:28.189 | לקמן |  |
| 4:28.190 | סי | ' |
| 4:28.191 | ט"ז | ) |
| 4:28.192 | והרב |  |
| 4:28.193 | נחלת |  |
| 4:28.194 | יוסף |  |
| 4:28.195 | דף |  |
| 4:28.196 | ס"א |  |
| 4:28.197 | וס"ב | ! |
| 4:28.198 | ומלכות |  |
| 4:28.199 | מתקלסת |  |
| 4:28.200 | במאמר |  |
| 4:28.201 | אלה |  |
| 4:28.202 | הם |  |
| 4:28.203 | האלוהים |  |
| 4:28.204 | המכים |  |
| 4:28.205 | את |  |
| 4:28.206 | מצרים |  |
| 4:28.207 | שהמלכות |  |
| 4:28.208 | היתה |  |
| 4:28.209 | מכה |  |
| 4:28.210 | בהם |  |
| 4:28.211 | בסיועת |  |
| 4:28.212 | הבינה |  |
| 4:28.213 | כמו |  |
| 4:28.214 | שכ | ' |
| 4:28.215 | הר | ' |
| 4:28.216 | נחלת |  |
| 4:28.217 | יוסף |  |
| 4:28.218 | ע"ש | . |


## Milchamot ha-Shem 4:29

Hebrew (Wikisource): ראה ידידי איך המקובלים מקלסים לכל גדולי הפרצופים בקילוסו של מלך מלכי המלכים כאלו הם נושאים במשאי עמו היפך דברי רז"ל שאמרו אתה הוא האלהים לבדך וכו' כדלעיל! דרחמנא לישיזבן מאמונות זרות כאלה! וראיתי להר' מלמד התלמידים בפרשת יתרו שכ' שהמלאכים נבראו קודם בריאת שמים וארץ והביא דבריו בעל מנורת המאור בפ' צ"ג, ומסתברא דנפשיה כתב כן, ואומדנא בעלמא היא, ומסתייע לדייק לה מברכת ברוך שאמר, וראיתו דיוקו ליתיה דאנן מעיקרא לא גרסינן ברוך עושה בראשית כמו שכתב מהרי'ץ בסי' עץ חיים, ואף לפי גרסתם אינו מוכרח דלא מצינו שהמלאכים נקראו ראשית, ולדברי רז"ל בבראשית רבה שומעין דמלתיהו מסתברא טפי, ועוד שדבריהם דברי קבלה, ואין להאריך במילי דאגדתא כ"כ במילי דלא נפקא מינה לדינא, כי כל מטרתינו פה לברר אמתת יחד הבורא ית' על פי תורתינו הקדושה ועל פי קבלת רבותינו ומעתיקי השמועה לדעת מה יעשה ישראל באומרו בכל יום ערב ובוקר שמע ישראל וכו' וה' אלהים יעזר לי וכו'.

| id | word | punctuation after |
|---|---|---|
| 4:29.1 | ראה |  |
| 4:29.2 | ידידי |  |
| 4:29.3 | איך |  |
| 4:29.4 | המקובלים |  |
| 4:29.5 | מקלסים |  |
| 4:29.6 | לכל |  |
| 4:29.7 | גדולי |  |
| 4:29.8 | הפרצופים |  |
| 4:29.9 | בקילוסו |  |
| 4:29.10 | של |  |
| 4:29.11 | מלך |  |
| 4:29.12 | מלכי |  |
| 4:29.13 | המלכים |  |
| 4:29.14 | כאלו |  |
| 4:29.15 | הם |  |
| 4:29.16 | נושאים |  |
| 4:29.17 | במשאי |  |
| 4:29.18 | עמו |  |
| 4:29.19 | היפך |  |
| 4:29.20 | דברי |  |
| 4:29.21 | רז"ל |  |
| 4:29.22 | שאמרו |  |
| 4:29.23 | אתה |  |
| 4:29.24 | הוא |  |
| 4:29.25 | האלהים |  |
| 4:29.26 | לבדך |  |
| 4:29.27 | וכו | ' |
| 4:29.28 | כדלעיל | ! |
| 4:29.29 | דרחמנא |  |
| 4:29.30 | לישיזבן |  |
| 4:29.31 | מאמונות |  |
| 4:29.32 | זרות |  |
| 4:29.33 | כאלה | ! |
| 4:29.34 | וראיתי |  |
| 4:29.35 | להר | ' |
| 4:29.36 | מלמד |  |
| 4:29.37 | התלמידים |  |
| 4:29.38 | בפרשת |  |
| 4:29.39 | יתרו |  |
| 4:29.40 | שכ | ' |
| 4:29.41 | שהמלאכים |  |
| 4:29.42 | נבראו |  |
| 4:29.43 | קודם |  |
| 4:29.44 | בריאת |  |
| 4:29.45 | שמים |  |
| 4:29.46 | וארץ |  |
| 4:29.47 | והביא |  |
| 4:29.48 | דבריו |  |
| 4:29.49 | בעל |  |
| 4:29.50 | מנורת |  |
| 4:29.51 | המאור |  |
| 4:29.52 | בפ | ' |
| 4:29.53 | צ"ג | , |
| 4:29.54 | ומסתברא |  |
| 4:29.55 | דנפשיה |  |
| 4:29.56 | כתב |  |
| 4:29.57 | כן | , |
| 4:29.58 | ואומדנא |  |
| 4:29.59 | בעלמא |  |
| 4:29.60 | היא | , |
| 4:29.61 | ומסתייע |  |
| 4:29.62 | לדייק |  |
| 4:29.63 | לה |  |
| 4:29.64 | מברכת |  |
| 4:29.65 | ברוך |  |
| 4:29.66 | שאמר | , |
| 4:29.67 | וראיתו |  |
| 4:29.68 | דיוקו |  |
| 4:29.69 | ליתיה |  |
| 4:29.70 | דאנן |  |
| 4:29.71 | מעיקרא |  |
| 4:29.72 | לא |  |
| 4:29.73 | גרסינן |  |
| 4:29.74 | ברוך |  |
| 4:29.75 | עושה |  |
| 4:29.76 | בראשית |  |
| 4:29.77 | כמו |  |
| 4:29.78 | שכתב |  |
| 4:29.79 | מהרי'ץ |  |
| 4:29.80 | בסי | ' |
| 4:29.81 | עץ |  |
| 4:29.82 | חיים | , |
| 4:29.83 | ואף |  |
| 4:29.84 | לפי |  |
| 4:29.85 | גרסתם |  |
| 4:29.86 | אינו |  |
| 4:29.87 | מוכרח |  |
| 4:29.88 | דלא |  |
| 4:29.89 | מצינו |  |
| 4:29.90 | שהמלאכים |  |
| 4:29.91 | נקראו |  |
| 4:29.92 | ראשית | , |
| 4:29.93 | ולדברי |  |
| 4:29.94 | רז"ל |  |
| 4:29.95 | בבראשית |  |
| 4:29.96 | רבה |  |
| 4:29.97 | שומעין |  |
| 4:29.98 | דמלתיהו |  |
| 4:29.99 | מסתברא |  |
| 4:29.100 | טפי | , |
| 4:29.101 | ועוד |  |
| 4:29.102 | שדבריהם |  |
| 4:29.103 | דברי |  |
| 4:29.104 | קבלה | , |
| 4:29.105 | ואין |  |
| 4:29.106 | להאריך |  |
| 4:29.107 | במילי |  |
| 4:29.108 | דאגדתא |  |
| 4:29.109 | כ"כ |  |
| 4:29.110 | במילי |  |
| 4:29.111 | דלא |  |
| 4:29.112 | נפקא |  |
| 4:29.113 | מינה |  |
| 4:29.114 | לדינא | , |
| 4:29.115 | כי |  |
| 4:29.116 | כל |  |
| 4:29.117 | מטרתינו |  |
| 4:29.118 | פה |  |
| 4:29.119 | לברר |  |
| 4:29.120 | אמתת |  |
| 4:29.121 | יחד |  |
| 4:29.122 | הבורא |  |
| 4:29.123 | ית | ' |
| 4:29.124 | על |  |
| 4:29.125 | פי |  |
| 4:29.126 | תורתינו |  |
| 4:29.127 | הקדושה |  |
| 4:29.128 | ועל |  |
| 4:29.129 | פי |  |
| 4:29.130 | קבלת |  |
| 4:29.131 | רבותינו |  |
| 4:29.132 | ומעתיקי |  |
| 4:29.133 | השמועה |  |
| 4:29.134 | לדעת |  |
| 4:29.135 | מה |  |
| 4:29.136 | יעשה |  |
| 4:29.137 | ישראל |  |
| 4:29.138 | באומרו |  |
| 4:29.139 | בכל |  |
| 4:29.140 | יום |  |
| 4:29.141 | ערב |  |
| 4:29.142 | ובוקר |  |
| 4:29.143 | שמע |  |
| 4:29.144 | ישראל |  |
| 4:29.145 | וכו | ' |
| 4:29.146 | וה | ' |
| 4:29.147 | אלהים |  |
| 4:29.148 | יעזר |  |
| 4:29.149 | לי |  |
| 4:29.150 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

