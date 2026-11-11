# Translation pack: milhamot-hashem perek-7-22 (Milchamot ha-Shem 7:27–7:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-22`.

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
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
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
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.

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


## Milchamot ha-Shem 7:27

Hebrew (Wikisource): צט) עוד כתבו שם הגדולים והגליון הנז' משם הרב ישר בס' מצרף לחכמה ששמע מפי מגידי אמת, שבשנת ש"ף לאלך השישי כשבזזו הספרדים עיר הירילברג"ה לקחו מהאקימי"א כמה אלפים ספרים וביניהם ספרי הקדש על קלף. ובכלל ספרי הזהר על כלל הכ"ד ספרים משא סבל ושולחו הספרים לחשמנים ודוכוסים וכו'. גם שמועה זו ששמע הרב יש"ר שקר היא שהרי בשנת שי"ח כבר נדפס ספר הזהר ונפשט. ומתיא זו דאיתי בידיה לאו שמה מתיא, ועוד דמ"מ מידי עכו"ם יצא. והגאון ר' יעקב עמדין בספרו מטפחת ספרים כ' ב' רבי יצחק דמן עכו וז"ל שנת נ"ג ה"ר משה מליאון שוכן באויל"א (והוא בואר אלחגרה) היה בזמן הזה, ולמען תספרו לדור אחרון אודיעכם מה שמצאתי כתוב שר' יצחק דמן עכו הלך לחקור על ספר הזהר שכנראה היו דבריו מפליאים ישאבו ממקור עליון. ואשר הוא לשון ירושלמי נר' שדבריו דברי קבלה ואמת. ואשר הם לשון הקדש הם דברי זיוף. אמר ושאלתי את הלמידים שנמצא בידיהם בפר הזהר מאין בא להם. ולא ראיתי דבריהים מכוונים. יש מי שאמר שהרמב"ן מצא אותו בארץ ישראל ושלחו לקאטולוניא, והביאו לארגון. ונפל ביד רבי משה ליאון. וי"א שמעולם לא היה חבור של רשב"י. רק שר' משה ידע שם הכותב, ובכחו יכתוב דברים מפליאים. ולמען יקח בהם מחיר רב. היה תולה עצמו באילן גדול שהוא רשב"י וחביריו ובנו ר' אלעזר. ואבוא אל עיר יאליאדול"ר. ואמצא שם רבי משה דיליאן וישבע לי שספר רשב"י הנו בבית ישיבתו ואמצא שם חכם זקן ושמו ר' דוד רסאן קרובו ואמצא חן בעיני, ואשביעהו לאמר אם היה יודע אמתת ס' הזהר אם היה אמת אם לא. ויאמר החכם כי נתבאר לו בודאי שר' משה בשם הכותב היה בוראם. וסבת האמתות הוא כי ר' משה הנז' היה כותב סודות ונפלאות לעשירים אשר במלכות ההיא ולוקח בהם מתנות רבות זהב וכסף. ואותו היום כל מה שהיה מרויח היה מפזר עד שביום מותו לא נשאר לו אפי' פרוטה ואשתו ובניו ברעב ובצמא ובעירום ובחסד כל וכשמעי מיתתו קמתי והלכתי אל רבי יוסף דיאוילא שהיה עשיר גדול, ובחייו לקח ממנו ממון, ואומר לו הנה עתך עת דודים, תזכה לספר הגדול, ותקח הספר מיד אשתו במפה יפה שתשלח בידי אשתך. שהיא כעת ערומה, וכן עשה שלח ביד אשתו מתנות אל אשת ר' משה הנז' ותשבע לה אשת ר' משה הנז', שמעולם לא היה ספר לאישי רק מלבו ודעתו היה כותב. ואומר לאישי למה אתה תולה עיונך בחכם אחר? הלא טוב היה לך לומר שמשכלך אתה כותב. וידוך כי תיטיב לך. ואישי ענני שאם הייתי אומר שמדעתי הם לא יחשיבו אותם ולא יתנו מתנות בהם. ולכן אני תולה אותם ברשב"י וחביריו, ואחר דברתי עם בתו ואמצא דבריהם מכוונים. אלה הדברים אשר דבר אלי החכם הזקן רבי דוד, ואסע מאויל"א ואבוא אל טלאבירה מצאתי שם זקן מופלא שמו רבי יוסף הלוי בן ר' יוסף טורדוס המקובל, ואחקורה שמו ולדעת מה שידו מגעת בידיעת ספר הזהר, ויאמר אלי דע באמת כי היה ביד ר' משה ספר הזהר אמתי, וסבת אמתות דברי הוא כי היה נותן לי קונדריסים ואני לבחון הספק הזה. לקחתי קונדריס אחד מבין הקונדריסים ואמרתי לו שנאבד ממני שיתן לי ממלא מקומו, ויאמר לי הראה לי אחרית הקונדריס הקודם לו, וראשית הקונדריס המתאחר, ויכתוב לי קונדריס. וארא והנה אין ביניהם שינוי כלל אפילו תיבה אחת, רק שפה א' ודברים אחדים. היש נסיון גדול מזה? אחר זה באתי לטוליטולא ואמרו לי שבחינת ר' יוסף הלוי אינה בחינה, שאולי קודם לחזות הקונדריסים לאדם היה מעתיק אחד לעצמו, ואחר כך מתוך אותם הקונדריסים אשר בביתו היה כותב לאחרים. ע"כ הדברים שהביאם בס' יוחסין. ושם הגיה המביא ספר יוחסין דברי מה בכך נגד הדברים האלה, וכבר השיב על דבריו הגאון יעב"ץ! והוא מאריך עור לפשר לקיים מקצת ולבטל מקצת יעויין שם באורך. והנה מזה תבין שלא נודע לספר הזהר ומקור שממנו יצא ונתפשט. כשאר חיבורי רז"ל המפורסמים מיום התגלותם. המשנה והתלמוד בבלי וירושלמי ותוספתא וספרא וספרי שכולם ידוע ומפורסם מקורן ומי הוא מחברן. אבל ספר הזהר המסית זה אומר בכה וזה אומר בכה. וקורא אני עליו מאמר רז"ל בירושלמי שבת פ' ר"א דמילה ופסחים פרק אלו דברים שאמרו שם רבי זעירא בשם רבי אלעזר כל תורה שאין לה בית אב אינה תורה עכ"ל. ואל תפנה למתק שפתיו במה שהוא מלהיב את לב קוראיו ומעוררם ליראת שמים ולהשתדל בעסק תורתו החדשה שהוא קורא לה בת מלך. לפי שאין כוונתו רק להביא את הקורא לאמונה נכרית להאמין באלהים רבים משפחות משפחות. משפחת האלוהות שבעולם האצילות, שהם אדם קדמון, ועתיק יומין, ואריך אנפין, ואבא ואימא, וזעיר אנפין ונוקביה. וכנגדן בעולם הבריאה, וכמוהם בעולם היצירה, וכן בעולם העשייה, ואת כולם הם מקבלים אותם לאלוהות ומשבחים ומפארים אותם בשעת התפלה באמרם אין כאלהינו וכו' כמו שכתבתי לעיל סי' ע"ט בשם סידור חסד לאברהם. דודאי דרושים כאלה אסור לשמען וכל שכן לקרותן ולהאמין בהם.

| id | word | punctuation after |
|---|---|---|
| 7:27.1 | צט | ) |
| 7:27.2 | עוד |  |
| 7:27.3 | כתבו |  |
| 7:27.4 | שם |  |
| 7:27.5 | הגדולים |  |
| 7:27.6 | והגליון |  |
| 7:27.7 | הנז | ' |
| 7:27.8 | משם |  |
| 7:27.9 | הרב |  |
| 7:27.10 | ישר |  |
| 7:27.11 | בס | ' |
| 7:27.12 | מצרף |  |
| 7:27.13 | לחכמה |  |
| 7:27.14 | ששמע |  |
| 7:27.15 | מפי |  |
| 7:27.16 | מגידי |  |
| 7:27.17 | אמת | , |
| 7:27.18 | שבשנת |  |
| 7:27.19 | ש"ף |  |
| 7:27.20 | לאלך |  |
| 7:27.21 | השישי |  |
| 7:27.22 | כשבזזו |  |
| 7:27.23 | הספרדים |  |
| 7:27.24 | עיר |  |
| 7:27.25 | הירילברג"ה |  |
| 7:27.26 | לקחו |  |
| 7:27.27 | מהאקימי"א |  |
| 7:27.28 | כמה |  |
| 7:27.29 | אלפים |  |
| 7:27.30 | ספרים |  |
| 7:27.31 | וביניהם |  |
| 7:27.32 | ספרי |  |
| 7:27.33 | הקדש |  |
| 7:27.34 | על |  |
| 7:27.35 | קלף | . |
| 7:27.36 | ובכלל |  |
| 7:27.37 | ספרי |  |
| 7:27.38 | הזהר |  |
| 7:27.39 | על |  |
| 7:27.40 | כלל |  |
| 7:27.41 | הכ"ד |  |
| 7:27.42 | ספרים |  |
| 7:27.43 | משא |  |
| 7:27.44 | סבל |  |
| 7:27.45 | ושולחו |  |
| 7:27.46 | הספרים |  |
| 7:27.47 | לחשמנים |  |
| 7:27.48 | ודוכוסים |  |
| 7:27.49 | וכו | '. |
| 7:27.50 | גם |  |
| 7:27.51 | שמועה |  |
| 7:27.52 | זו |  |
| 7:27.53 | ששמע |  |
| 7:27.54 | הרב |  |
| 7:27.55 | יש"ר |  |
| 7:27.56 | שקר |  |
| 7:27.57 | היא |  |
| 7:27.58 | שהרי |  |
| 7:27.59 | בשנת |  |
| 7:27.60 | שי"ח |  |
| 7:27.61 | כבר |  |
| 7:27.62 | נדפס |  |
| 7:27.63 | ספר |  |
| 7:27.64 | הזהר |  |
| 7:27.65 | ונפשט | . |
| 7:27.66 | ומתיא |  |
| 7:27.67 | זו |  |
| 7:27.68 | דאיתי |  |
| 7:27.69 | בידיה |  |
| 7:27.70 | לאו |  |
| 7:27.71 | שמה |  |
| 7:27.72 | מתיא | , |
| 7:27.73 | ועוד |  |
| 7:27.74 | דמ"מ |  |
| 7:27.75 | מידי |  |
| 7:27.76 | עכו"ם |  |
| 7:27.77 | יצא | . |
| 7:27.78 | והגאון |  |
| 7:27.79 | ר | ' |
| 7:27.80 | יעקב |  |
| 7:27.81 | עמדין |  |
| 7:27.82 | בספרו |  |
| 7:27.83 | מטפחת |  |
| 7:27.84 | ספרים |  |
| 7:27.85 | כ | ' |
| 7:27.86 | ב | ' |
| 7:27.87 | רבי |  |
| 7:27.88 | יצחק |  |
| 7:27.89 | דמן |  |
| 7:27.90 | עכו |  |
| 7:27.91 | וז"ל |  |
| 7:27.92 | שנת |  |
| 7:27.93 | נ"ג |  |
| 7:27.94 | ה"ר |  |
| 7:27.95 | משה |  |
| 7:27.96 | מליאון |  |
| 7:27.97 | שוכן |  |
| 7:27.98 | באויל"א |  |
| 7:27.99 | והוא |  |
| 7:27.100 | בואר |  |
| 7:27.101 | אלחגרה | ) |
| 7:27.102 | היה |  |
| 7:27.103 | בזמן |  |
| 7:27.104 | הזה | , |
| 7:27.105 | ולמען |  |
| 7:27.106 | תספרו |  |
| 7:27.107 | לדור |  |
| 7:27.108 | אחרון |  |
| 7:27.109 | אודיעכם |  |
| 7:27.110 | מה |  |
| 7:27.111 | שמצאתי |  |
| 7:27.112 | כתוב |  |
| 7:27.113 | שר | ' |
| 7:27.114 | יצחק |  |
| 7:27.115 | דמן |  |
| 7:27.116 | עכו |  |
| 7:27.117 | הלך |  |
| 7:27.118 | לחקור |  |
| 7:27.119 | על |  |
| 7:27.120 | ספר |  |
| 7:27.121 | הזהר |  |
| 7:27.122 | שכנראה |  |
| 7:27.123 | היו |  |
| 7:27.124 | דבריו |  |
| 7:27.125 | מפליאים |  |
| 7:27.126 | ישאבו |  |
| 7:27.127 | ממקור |  |
| 7:27.128 | עליון | . |
| 7:27.129 | ואשר |  |
| 7:27.130 | הוא |  |
| 7:27.131 | לשון |  |
| 7:27.132 | ירושלמי |  |
| 7:27.133 | נר | ' |
| 7:27.134 | שדבריו |  |
| 7:27.135 | דברי |  |
| 7:27.136 | קבלה |  |
| 7:27.137 | ואמת | . |
| 7:27.138 | ואשר |  |
| 7:27.139 | הם |  |
| 7:27.140 | לשון |  |
| 7:27.141 | הקדש |  |
| 7:27.142 | הם |  |
| 7:27.143 | דברי |  |
| 7:27.144 | זיוף | . |
| 7:27.145 | אמר |  |
| 7:27.146 | ושאלתי |  |
| 7:27.147 | את |  |
| 7:27.148 | הלמידים |  |
| 7:27.149 | שנמצא |  |
| 7:27.150 | בידיהם |  |
| 7:27.151 | בפר |  |
| 7:27.152 | הזהר |  |
| 7:27.153 | מאין |  |
| 7:27.154 | בא |  |
| 7:27.155 | להם | . |
| 7:27.156 | ולא |  |
| 7:27.157 | ראיתי |  |
| 7:27.158 | דבריהים |  |
| 7:27.159 | מכוונים | . |
| 7:27.160 | יש |  |
| 7:27.161 | מי |  |
| 7:27.162 | שאמר |  |
| 7:27.163 | שהרמב"ן |  |
| 7:27.164 | מצא |  |
| 7:27.165 | אותו |  |
| 7:27.166 | בארץ |  |
| 7:27.167 | ישראל |  |
| 7:27.168 | ושלחו |  |
| 7:27.169 | לקאטולוניא | , |
| 7:27.170 | והביאו |  |
| 7:27.171 | לארגון | . |
| 7:27.172 | ונפל |  |
| 7:27.173 | ביד |  |
| 7:27.174 | רבי |  |
| 7:27.175 | משה |  |
| 7:27.176 | ליאון | . |
| 7:27.177 | וי"א |  |
| 7:27.178 | שמעולם |  |
| 7:27.179 | לא |  |
| 7:27.180 | היה |  |
| 7:27.181 | חבור |  |
| 7:27.182 | של |  |
| 7:27.183 | רשב"י | . |
| 7:27.184 | רק |  |
| 7:27.185 | שר | ' |
| 7:27.186 | משה |  |
| 7:27.187 | ידע |  |
| 7:27.188 | שם |  |
| 7:27.189 | הכותב | , |
| 7:27.190 | ובכחו |  |
| 7:27.191 | יכתוב |  |
| 7:27.192 | דברים |  |
| 7:27.193 | מפליאים | . |
| 7:27.194 | ולמען |  |
| 7:27.195 | יקח |  |
| 7:27.196 | בהם |  |
| 7:27.197 | מחיר |  |
| 7:27.198 | רב | . |
| 7:27.199 | היה |  |
| 7:27.200 | תולה |  |
| 7:27.201 | עצמו |  |
| 7:27.202 | באילן |  |
| 7:27.203 | גדול |  |
| 7:27.204 | שהוא |  |
| 7:27.205 | רשב"י |  |
| 7:27.206 | וחביריו |  |
| 7:27.207 | ובנו |  |
| 7:27.208 | ר | ' |
| 7:27.209 | אלעזר | . |
| 7:27.210 | ואבוא |  |
| 7:27.211 | אל |  |
| 7:27.212 | עיר |  |
| 7:27.213 | יאליאדול"ר | . |
| 7:27.214 | ואמצא |  |
| 7:27.215 | שם |  |
| 7:27.216 | רבי |  |
| 7:27.217 | משה |  |
| 7:27.218 | דיליאן |  |
| 7:27.219 | וישבע |  |
| 7:27.220 | לי |  |
| 7:27.221 | שספר |  |
| 7:27.222 | רשב"י |  |
| 7:27.223 | הנו |  |
| 7:27.224 | בבית |  |
| 7:27.225 | ישיבתו |  |
| 7:27.226 | ואמצא |  |
| 7:27.227 | שם |  |
| 7:27.228 | חכם |  |
| 7:27.229 | זקן |  |
| 7:27.230 | ושמו |  |
| 7:27.231 | ר | ' |
| 7:27.232 | דוד |  |
| 7:27.233 | רסאן |  |
| 7:27.234 | קרובו |  |
| 7:27.235 | ואמצא |  |
| 7:27.236 | חן |  |
| 7:27.237 | בעיני | , |
| 7:27.238 | ואשביעהו |  |
| 7:27.239 | לאמר |  |
| 7:27.240 | אם |  |
| 7:27.241 | היה |  |
| 7:27.242 | יודע |  |
| 7:27.243 | אמתת |  |
| 7:27.244 | ס | ' |
| 7:27.245 | הזהר |  |
| 7:27.246 | אם |  |
| 7:27.247 | היה |  |
| 7:27.248 | אמת |  |
| 7:27.249 | אם |  |
| 7:27.250 | לא | . |
| 7:27.251 | ויאמר |  |
| 7:27.252 | החכם |  |
| 7:27.253 | כי |  |
| 7:27.254 | נתבאר |  |
| 7:27.255 | לו |  |
| 7:27.256 | בודאי |  |
| 7:27.257 | שר | ' |
| 7:27.258 | משה |  |
| 7:27.259 | בשם |  |
| 7:27.260 | הכותב |  |
| 7:27.261 | היה |  |
| 7:27.262 | בוראם | . |
| 7:27.263 | וסבת |  |
| 7:27.264 | האמתות |  |
| 7:27.265 | הוא |  |
| 7:27.266 | כי |  |
| 7:27.267 | ר | ' |
| 7:27.268 | משה |  |
| 7:27.269 | הנז | ' |
| 7:27.270 | היה |  |
| 7:27.271 | כותב |  |
| 7:27.272 | סודות |  |
| 7:27.273 | ונפלאות |  |
| 7:27.274 | לעשירים |  |
| 7:27.275 | אשר |  |
| 7:27.276 | במלכות |  |
| 7:27.277 | ההיא |  |
| 7:27.278 | ולוקח |  |
| 7:27.279 | בהם |  |
| 7:27.280 | מתנות |  |
| 7:27.281 | רבות |  |
| 7:27.282 | זהב |  |
| 7:27.283 | וכסף | . |
| 7:27.284 | ואותו |  |
| 7:27.285 | היום |  |
| 7:27.286 | כל |  |
| 7:27.287 | מה |  |
| 7:27.288 | שהיה |  |
| 7:27.289 | מרויח |  |
| 7:27.290 | היה |  |
| 7:27.291 | מפזר |  |
| 7:27.292 | עד |  |
| 7:27.293 | שביום |  |
| 7:27.294 | מותו |  |
| 7:27.295 | לא |  |
| 7:27.296 | נשאר |  |
| 7:27.297 | לו |  |
| 7:27.298 | אפי | ' |
| 7:27.299 | פרוטה |  |
| 7:27.300 | ואשתו |  |
| 7:27.301 | ובניו |  |
| 7:27.302 | ברעב |  |
| 7:27.303 | ובצמא |  |
| 7:27.304 | ובעירום |  |
| 7:27.305 | ובחסד |  |
| 7:27.306 | כל |  |
| 7:27.307 | וכשמעי |  |
| 7:27.308 | מיתתו |  |
| 7:27.309 | קמתי |  |
| 7:27.310 | והלכתי |  |
| 7:27.311 | אל |  |
| 7:27.312 | רבי |  |
| 7:27.313 | יוסף |  |
| 7:27.314 | דיאוילא |  |
| 7:27.315 | שהיה |  |
| 7:27.316 | עשיר |  |
| 7:27.317 | גדול | , |
| 7:27.318 | ובחייו |  |
| 7:27.319 | לקח |  |
| 7:27.320 | ממנו |  |
| 7:27.321 | ממון | , |
| 7:27.322 | ואומר |  |
| 7:27.323 | לו |  |
| 7:27.324 | הנה |  |
| 7:27.325 | עתך |  |
| 7:27.326 | עת |  |
| 7:27.327 | דודים | , |
| 7:27.328 | תזכה |  |
| 7:27.329 | לספר |  |
| 7:27.330 | הגדול | , |
| 7:27.331 | ותקח |  |
| 7:27.332 | הספר |  |
| 7:27.333 | מיד |  |
| 7:27.334 | אשתו |  |
| 7:27.335 | במפה |  |
| 7:27.336 | יפה |  |
| 7:27.337 | שתשלח |  |
| 7:27.338 | בידי |  |
| 7:27.339 | אשתך | . |
| 7:27.340 | שהיא |  |
| 7:27.341 | כעת |  |
| 7:27.342 | ערומה | , |
| 7:27.343 | וכן |  |
| 7:27.344 | עשה |  |
| 7:27.345 | שלח |  |
| 7:27.346 | ביד |  |
| 7:27.347 | אשתו |  |
| 7:27.348 | מתנות |  |
| 7:27.349 | אל |  |
| 7:27.350 | אשת |  |
| 7:27.351 | ר | ' |
| 7:27.352 | משה |  |
| 7:27.353 | הנז | ' |
| 7:27.354 | ותשבע |  |
| 7:27.355 | לה |  |
| 7:27.356 | אשת |  |
| 7:27.357 | ר | ' |
| 7:27.358 | משה |  |
| 7:27.359 | הנז | ', |
| 7:27.360 | שמעולם |  |
| 7:27.361 | לא |  |
| 7:27.362 | היה |  |
| 7:27.363 | ספר |  |
| 7:27.364 | לאישי |  |
| 7:27.365 | רק |  |
| 7:27.366 | מלבו |  |
| 7:27.367 | ודעתו |  |
| 7:27.368 | היה |  |
| 7:27.369 | כותב | . |
| 7:27.370 | ואומר |  |
| 7:27.371 | לאישי |  |
| 7:27.372 | למה |  |
| 7:27.373 | אתה |  |
| 7:27.374 | תולה |  |
| 7:27.375 | עיונך |  |
| 7:27.376 | בחכם |  |
| 7:27.377 | אחר | ? |
| 7:27.378 | הלא |  |
| 7:27.379 | טוב |  |
| 7:27.380 | היה |  |
| 7:27.381 | לך |  |
| 7:27.382 | לומר |  |
| 7:27.383 | שמשכלך |  |
| 7:27.384 | אתה |  |
| 7:27.385 | כותב | . |
| 7:27.386 | וידוך |  |
| 7:27.387 | כי |  |
| 7:27.388 | תיטיב |  |
| 7:27.389 | לך | . |
| 7:27.390 | ואישי |  |
| 7:27.391 | ענני |  |
| 7:27.392 | שאם |  |
| 7:27.393 | הייתי |  |
| 7:27.394 | אומר |  |
| 7:27.395 | שמדעתי |  |
| 7:27.396 | הם |  |
| 7:27.397 | לא |  |
| 7:27.398 | יחשיבו |  |
| 7:27.399 | אותם |  |
| 7:27.400 | ולא |  |
| 7:27.401 | יתנו |  |
| 7:27.402 | מתנות |  |
| 7:27.403 | בהם | . |
| 7:27.404 | ולכן |  |
| 7:27.405 | אני |  |
| 7:27.406 | תולה |  |
| 7:27.407 | אותם |  |
| 7:27.408 | ברשב"י |  |
| 7:27.409 | וחביריו | , |
| 7:27.410 | ואחר |  |
| 7:27.411 | דברתי |  |
| 7:27.412 | עם |  |
| 7:27.413 | בתו |  |
| 7:27.414 | ואמצא |  |
| 7:27.415 | דבריהם |  |
| 7:27.416 | מכוונים | . |
| 7:27.417 | אלה |  |
| 7:27.418 | הדברים |  |
| 7:27.419 | אשר |  |
| 7:27.420 | דבר |  |
| 7:27.421 | אלי |  |
| 7:27.422 | החכם |  |
| 7:27.423 | הזקן |  |
| 7:27.424 | רבי |  |
| 7:27.425 | דוד | , |
| 7:27.426 | ואסע |  |
| 7:27.427 | מאויל"א |  |
| 7:27.428 | ואבוא |  |
| 7:27.429 | אל |  |
| 7:27.430 | טלאבירה |  |
| 7:27.431 | מצאתי |  |
| 7:27.432 | שם |  |
| 7:27.433 | זקן |  |
| 7:27.434 | מופלא |  |
| 7:27.435 | שמו |  |
| 7:27.436 | רבי |  |
| 7:27.437 | יוסף |  |
| 7:27.438 | הלוי |  |
| 7:27.439 | בן |  |
| 7:27.440 | ר | ' |
| 7:27.441 | יוסף |  |
| 7:27.442 | טורדוס |  |
| 7:27.443 | המקובל | , |
| 7:27.444 | ואחקורה |  |
| 7:27.445 | שמו |  |
| 7:27.446 | ולדעת |  |
| 7:27.447 | מה |  |
| 7:27.448 | שידו |  |
| 7:27.449 | מגעת |  |
| 7:27.450 | בידיעת |  |
| 7:27.451 | ספר |  |
| 7:27.452 | הזהר | , |
| 7:27.453 | ויאמר |  |
| 7:27.454 | אלי |  |
| 7:27.455 | דע |  |
| 7:27.456 | באמת |  |
| 7:27.457 | כי |  |
| 7:27.458 | היה |  |
| 7:27.459 | ביד |  |
| 7:27.460 | ר | ' |
| 7:27.461 | משה |  |
| 7:27.462 | ספר |  |
| 7:27.463 | הזהר |  |
| 7:27.464 | אמתי | , |
| 7:27.465 | וסבת |  |
| 7:27.466 | אמתות |  |
| 7:27.467 | דברי |  |
| 7:27.468 | הוא |  |
| 7:27.469 | כי |  |
| 7:27.470 | היה |  |
| 7:27.471 | נותן |  |
| 7:27.472 | לי |  |
| 7:27.473 | קונדריסים |  |
| 7:27.474 | ואני |  |
| 7:27.475 | לבחון |  |
| 7:27.476 | הספק |  |
| 7:27.477 | הזה | . |
| 7:27.478 | לקחתי |  |
| 7:27.479 | קונדריס |  |
| 7:27.480 | אחד |  |
| 7:27.481 | מבין |  |
| 7:27.482 | הקונדריסים |  |
| 7:27.483 | ואמרתי |  |
| 7:27.484 | לו |  |
| 7:27.485 | שנאבד |  |
| 7:27.486 | ממני |  |
| 7:27.487 | שיתן |  |
| 7:27.488 | לי |  |
| 7:27.489 | ממלא |  |
| 7:27.490 | מקומו | , |
| 7:27.491 | ויאמר |  |
| 7:27.492 | לי |  |
| 7:27.493 | הראה |  |
| 7:27.494 | לי |  |
| 7:27.495 | אחרית |  |
| 7:27.496 | הקונדריס |  |
| 7:27.497 | הקודם |  |
| 7:27.498 | לו | , |
| 7:27.499 | וראשית |  |
| 7:27.500 | הקונדריס |  |
| 7:27.501 | המתאחר | , |
| 7:27.502 | ויכתוב |  |
| 7:27.503 | לי |  |
| 7:27.504 | קונדריס | . |
| 7:27.505 | וארא |  |
| 7:27.506 | והנה |  |
| 7:27.507 | אין |  |
| 7:27.508 | ביניהם |  |
| 7:27.509 | שינוי |  |
| 7:27.510 | כלל |  |
| 7:27.511 | אפילו |  |
| 7:27.512 | תיבה |  |
| 7:27.513 | אחת | , |
| 7:27.514 | רק |  |
| 7:27.515 | שפה |  |
| 7:27.516 | א | ' |
| 7:27.517 | ודברים |  |
| 7:27.518 | אחדים | . |
| 7:27.519 | היש |  |
| 7:27.520 | נסיון |  |
| 7:27.521 | גדול |  |
| 7:27.522 | מזה | ? |
| 7:27.523 | אחר |  |
| 7:27.524 | זה |  |
| 7:27.525 | באתי |  |
| 7:27.526 | לטוליטולא |  |
| 7:27.527 | ואמרו |  |
| 7:27.528 | לי |  |
| 7:27.529 | שבחינת |  |
| 7:27.530 | ר | ' |
| 7:27.531 | יוסף |  |
| 7:27.532 | הלוי |  |
| 7:27.533 | אינה |  |
| 7:27.534 | בחינה | , |
| 7:27.535 | שאולי |  |
| 7:27.536 | קודם |  |
| 7:27.537 | לחזות |  |
| 7:27.538 | הקונדריסים |  |
| 7:27.539 | לאדם |  |
| 7:27.540 | היה |  |
| 7:27.541 | מעתיק |  |
| 7:27.542 | אחד |  |
| 7:27.543 | לעצמו | , |
| 7:27.544 | ואחר |  |
| 7:27.545 | כך |  |
| 7:27.546 | מתוך |  |
| 7:27.547 | אותם |  |
| 7:27.548 | הקונדריסים |  |
| 7:27.549 | אשר |  |
| 7:27.550 | בביתו |  |
| 7:27.551 | היה |  |
| 7:27.552 | כותב |  |
| 7:27.553 | לאחרים | . |
| 7:27.554 | ע"כ |  |
| 7:27.555 | הדברים |  |
| 7:27.556 | שהביאם |  |
| 7:27.557 | בס | ' |
| 7:27.558 | יוחסין | . |
| 7:27.559 | ושם |  |
| 7:27.560 | הגיה |  |
| 7:27.561 | המביא |  |
| 7:27.562 | ספר |  |
| 7:27.563 | יוחסין |  |
| 7:27.564 | דברי |  |
| 7:27.565 | מה |  |
| 7:27.566 | בכך |  |
| 7:27.567 | נגד |  |
| 7:27.568 | הדברים |  |
| 7:27.569 | האלה | , |
| 7:27.570 | וכבר |  |
| 7:27.571 | השיב |  |
| 7:27.572 | על |  |
| 7:27.573 | דבריו |  |
| 7:27.574 | הגאון |  |
| 7:27.575 | יעב"ץ | ! |
| 7:27.576 | והוא |  |
| 7:27.577 | מאריך |  |
| 7:27.578 | עור |  |
| 7:27.579 | לפשר |  |
| 7:27.580 | לקיים |  |
| 7:27.581 | מקצת |  |
| 7:27.582 | ולבטל |  |
| 7:27.583 | מקצת |  |
| 7:27.584 | יעויין |  |
| 7:27.585 | שם |  |
| 7:27.586 | באורך | . |
| 7:27.587 | והנה |  |
| 7:27.588 | מזה |  |
| 7:27.589 | תבין |  |
| 7:27.590 | שלא |  |
| 7:27.591 | נודע |  |
| 7:27.592 | לספר |  |
| 7:27.593 | הזהר |  |
| 7:27.594 | ומקור |  |
| 7:27.595 | שממנו |  |
| 7:27.596 | יצא |  |
| 7:27.597 | ונתפשט | . |
| 7:27.598 | כשאר |  |
| 7:27.599 | חיבורי |  |
| 7:27.600 | רז"ל |  |
| 7:27.601 | המפורסמים |  |
| 7:27.602 | מיום |  |
| 7:27.603 | התגלותם | . |
| 7:27.604 | המשנה |  |
| 7:27.605 | והתלמוד |  |
| 7:27.606 | בבלי |  |
| 7:27.607 | וירושלמי |  |
| 7:27.608 | ותוספתא |  |
| 7:27.609 | וספרא |  |
| 7:27.610 | וספרי |  |
| 7:27.611 | שכולם |  |
| 7:27.612 | ידוע |  |
| 7:27.613 | ומפורסם |  |
| 7:27.614 | מקורן |  |
| 7:27.615 | ומי |  |
| 7:27.616 | הוא |  |
| 7:27.617 | מחברן | . |
| 7:27.618 | אבל |  |
| 7:27.619 | ספר |  |
| 7:27.620 | הזהר |  |
| 7:27.621 | המסית |  |
| 7:27.622 | זה |  |
| 7:27.623 | אומר |  |
| 7:27.624 | בכה |  |
| 7:27.625 | וזה |  |
| 7:27.626 | אומר |  |
| 7:27.627 | בכה | . |
| 7:27.628 | וקורא |  |
| 7:27.629 | אני |  |
| 7:27.630 | עליו |  |
| 7:27.631 | מאמר |  |
| 7:27.632 | רז"ל |  |
| 7:27.633 | בירושלמי |  |
| 7:27.634 | שבת |  |
| 7:27.635 | פ | ' |
| 7:27.636 | ר"א |  |
| 7:27.637 | דמילה |  |
| 7:27.638 | ופסחים |  |
| 7:27.639 | פרק |  |
| 7:27.640 | אלו |  |
| 7:27.641 | דברים |  |
| 7:27.642 | שאמרו |  |
| 7:27.643 | שם |  |
| 7:27.644 | רבי |  |
| 7:27.645 | זעירא |  |
| 7:27.646 | בשם |  |
| 7:27.647 | רבי |  |
| 7:27.648 | אלעזר |  |
| 7:27.649 | כל |  |
| 7:27.650 | תורה |  |
| 7:27.651 | שאין |  |
| 7:27.652 | לה |  |
| 7:27.653 | בית |  |
| 7:27.654 | אב |  |
| 7:27.655 | אינה |  |
| 7:27.656 | תורה |  |
| 7:27.657 | עכ"ל | . |
| 7:27.658 | ואל |  |
| 7:27.659 | תפנה |  |
| 7:27.660 | למתק |  |
| 7:27.661 | שפתיו |  |
| 7:27.662 | במה |  |
| 7:27.663 | שהוא |  |
| 7:27.664 | מלהיב |  |
| 7:27.665 | את |  |
| 7:27.666 | לב |  |
| 7:27.667 | קוראיו |  |
| 7:27.668 | ומעוררם |  |
| 7:27.669 | ליראת |  |
| 7:27.670 | שמים |  |
| 7:27.671 | ולהשתדל |  |
| 7:27.672 | בעסק |  |
| 7:27.673 | תורתו |  |
| 7:27.674 | החדשה |  |
| 7:27.675 | שהוא |  |
| 7:27.676 | קורא |  |
| 7:27.677 | לה |  |
| 7:27.678 | בת |  |
| 7:27.679 | מלך | . |
| 7:27.680 | לפי |  |
| 7:27.681 | שאין |  |
| 7:27.682 | כוונתו |  |
| 7:27.683 | רק |  |
| 7:27.684 | להביא |  |
| 7:27.685 | את |  |
| 7:27.686 | הקורא |  |
| 7:27.687 | לאמונה |  |
| 7:27.688 | נכרית |  |
| 7:27.689 | להאמין |  |
| 7:27.690 | באלהים |  |
| 7:27.691 | רבים |  |
| 7:27.692 | משפחות |  |
| 7:27.693 | משפחות | . |
| 7:27.694 | משפחת |  |
| 7:27.695 | האלוהות |  |
| 7:27.696 | שבעולם |  |
| 7:27.697 | האצילות | , |
| 7:27.698 | שהם |  |
| 7:27.699 | אדם |  |
| 7:27.700 | קדמון | , |
| 7:27.701 | ועתיק |  |
| 7:27.702 | יומין | , |
| 7:27.703 | ואריך |  |
| 7:27.704 | אנפין | , |
| 7:27.705 | ואבא |  |
| 7:27.706 | ואימא | , |
| 7:27.707 | וזעיר |  |
| 7:27.708 | אנפין |  |
| 7:27.709 | ונוקביה | . |
| 7:27.710 | וכנגדן |  |
| 7:27.711 | בעולם |  |
| 7:27.712 | הבריאה | , |
| 7:27.713 | וכמוהם |  |
| 7:27.714 | בעולם |  |
| 7:27.715 | היצירה | , |
| 7:27.716 | וכן |  |
| 7:27.717 | בעולם |  |
| 7:27.718 | העשייה | , |
| 7:27.719 | ואת |  |
| 7:27.720 | כולם |  |
| 7:27.721 | הם |  |
| 7:27.722 | מקבלים |  |
| 7:27.723 | אותם |  |
| 7:27.724 | לאלוהות |  |
| 7:27.725 | ומשבחים |  |
| 7:27.726 | ומפארים |  |
| 7:27.727 | אותם |  |
| 7:27.728 | בשעת |  |
| 7:27.729 | התפלה |  |
| 7:27.730 | באמרם |  |
| 7:27.731 | אין |  |
| 7:27.732 | כאלהינו |  |
| 7:27.733 | וכו | ' |
| 7:27.734 | כמו |  |
| 7:27.735 | שכתבתי |  |
| 7:27.736 | לעיל |  |
| 7:27.737 | סי | ' |
| 7:27.738 | ע"ט |  |
| 7:27.739 | בשם |  |
| 7:27.740 | סידור |  |
| 7:27.741 | חסד |  |
| 7:27.742 | לאברהם | . |
| 7:27.743 | דודאי |  |
| 7:27.744 | דרושים |  |
| 7:27.745 | כאלה |  |
| 7:27.746 | אסור |  |
| 7:27.747 | לשמען |  |
| 7:27.748 | וכל |  |
| 7:27.749 | שכן |  |
| 7:27.750 | לקרותן |  |
| 7:27.751 | ולהאמין |  |
| 7:27.752 | בהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

