# Translation pack: milhamot-hashem perek-7-07 (Milchamot ha-Shem 7:10–7:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-07`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
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
- מנחה in הגשת מנחה where the author argues from Gen 43 / 1 Sam 10: "gift";
  "meal offering" only in a sacrificial context.
- והוא שלא: "provided that" (fixed).
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


## Milchamot ha-Shem 7:10

Hebrew (Wikisource): פד) וכן ידוע הוא ומפורסם שבכל ארץ התימן נוהגין לומר סדר קדושה בחזרת התפלה בנוסח שכתב הרמב"ם ז"ל הסדר התפלה שלעולם שליח צבור אומר ברכה שלישית בנוסח שיש בו ומשרתיו שואלים זה לזה איה מקום כבודו ובין בשחרית בין במנחה בין במוסף שבת ור"ח ויום טוב וחול המועד לעולם אומרין איה מקום כבודו והוא מיוסד מרבותינו אנשי כנה"ג מתקני התפלות והברכות, וכמ"ש רז"ל בספרי ומדרש רבה שאפילו מלאכים וחיות הקדש שואלים איה מקום כבודו כמו שכ' לעיל סי' נ"ב בשם הרקח. וחדשים מקרוב באו לערך מאה שנה או יותר מעט שינו הנוסח שלא לומר ומשרתיו שואלים זה לזה איה מקום כבודו רק במוסף שבת בלבד מקפידים על האומרו בשאר ימים ומשתקים אותו. וטעמם כמו שכ' הר' משנת חסידים במסכת ר"ח פ"ג משנה ט' וז"ל ומפני שאין לזעיר ונוקביה. עלייה (במוסף ר"ח) עתה למעלה ממקום אצילותם כמו בשבת. לפיכך אין אומרים ומשרתיו שואלים זה לזה איה מקום כבודו. כי משרתיו יודעים מקומו עכ"ל. הרי בהדיא מבואר בדבריו דסבירא להו לבעלי הקבלה החדשה כי מאמר רבותי' מסדרי התפלות ובעלי התלמוד שהקב"ה נעלם ורואה ואינו נראה עד שאפילו מלאכים וחיות הקדש שואלים איה מקום כבודו, הוא מצד העתקו בזמן מה לעלות למעלה ממקום משכן כבודו. דהיינו במוסף שבת דוקא שאז נעתק ועלה למעלה. ואז כשרואים מקומו המורגל להיות שוכן שם פנוי. אז יפקד מקום מושבו ואז שואלים איה מקום כבודו. אמנם בשאר זמנים שהוא שוכן במקומו כל משרתיו יודעים מקום משכן כבודו ואין צריכים לשאול עליו כמאמר יהונתן בן שאול לדוד מחר חדש ונפקדת כי יפקד מושבך. ודבריהם אלו סותרים העיקר שהזכירו הראשונים האמור ביגדל אלהים חי ומבואר ג"כ בדברי הרמב"ם בפ' א' מיסודי התורה ובמורה. ובס' אהל יעקב פרשת וירא על מאמר אברהם אבינו לשי"ת אל נא תעבור מעל עבדך כתב דבהשקפה ראשונה יפלא מאד על אברהם אבינו איך ערב אל לבו לומר לפני השי"ת אל נא תעבור מעל עבדך. ובגמרא שבת קכ"ז אר"א מדת ב"ו אין קטן יכול לומר לגדול המתן עד שאבוא אצלך, ואלו בהקב"ה כתיב ויאמר אל נא תעבר מעל עבדך. אבל הבנת הענין הוא עם המאומת והמקובל אצלינו מפי קדמונינו ע"ה ובראשם הר' המורה, כי הקב"ה נעלה ומרומם מכל שינוי ותנועה והעתק מקום ובשגם מצינו בפרשת וירא וילך ה' וכו' אין הרצון בו חלילה שהקב"ה נעתק וסר ממקומו ובא להתראות אל מקום מצב מי שיראה אליו. אבל כל התקרבותו ית' והתגלותו לאדם תלוי אך בהתקרבות האדם אליו כמאמר אני לדודי ודודי לי, וזהו שאמר אני לדודי ועלי תשוקתו, כלומר תשוקתו ית' אלינו הוא דבר תלוי בידינו, ואברהם אבינו בעוצם דבקותו בו ית' נגלה אליו, ובהזדמן האורחים אצלו וצידד להכניסם ולטבוח טבח והכן ירא לנפשו שלא יפול מהדבק בשי"ת ושאל ממנו אל נא תעבור מעל עבדך, לבל יסיר הקב"ה עזרתו מאתו ולא ינתק מדבקותו בטרדת הכנסת אורחים וכו' עיין שם.

| id | word | punctuation after |
|---|---|---|
| 7:10.1 | פד | ) |
| 7:10.2 | וכן |  |
| 7:10.3 | ידוע |  |
| 7:10.4 | הוא |  |
| 7:10.5 | ומפורסם |  |
| 7:10.6 | שבכל |  |
| 7:10.7 | ארץ |  |
| 7:10.8 | התימן |  |
| 7:10.9 | נוהגין |  |
| 7:10.10 | לומר |  |
| 7:10.11 | סדר |  |
| 7:10.12 | קדושה |  |
| 7:10.13 | בחזרת |  |
| 7:10.14 | התפלה |  |
| 7:10.15 | בנוסח |  |
| 7:10.16 | שכתב |  |
| 7:10.17 | הרמב"ם |  |
| 7:10.18 | ז"ל |  |
| 7:10.19 | הסדר |  |
| 7:10.20 | התפלה |  |
| 7:10.21 | שלעולם |  |
| 7:10.22 | שליח |  |
| 7:10.23 | צבור |  |
| 7:10.24 | אומר |  |
| 7:10.25 | ברכה |  |
| 7:10.26 | שלישית |  |
| 7:10.27 | בנוסח |  |
| 7:10.28 | שיש |  |
| 7:10.29 | בו |  |
| 7:10.30 | ומשרתיו |  |
| 7:10.31 | שואלים |  |
| 7:10.32 | זה |  |
| 7:10.33 | לזה |  |
| 7:10.34 | איה |  |
| 7:10.35 | מקום |  |
| 7:10.36 | כבודו |  |
| 7:10.37 | ובין |  |
| 7:10.38 | בשחרית |  |
| 7:10.39 | בין |  |
| 7:10.40 | במנחה |  |
| 7:10.41 | בין |  |
| 7:10.42 | במוסף |  |
| 7:10.43 | שבת |  |
| 7:10.44 | ור"ח |  |
| 7:10.45 | ויום |  |
| 7:10.46 | טוב |  |
| 7:10.47 | וחול |  |
| 7:10.48 | המועד |  |
| 7:10.49 | לעולם |  |
| 7:10.50 | אומרין |  |
| 7:10.51 | איה |  |
| 7:10.52 | מקום |  |
| 7:10.53 | כבודו |  |
| 7:10.54 | והוא |  |
| 7:10.55 | מיוסד |  |
| 7:10.56 | מרבותינו |  |
| 7:10.57 | אנשי |  |
| 7:10.58 | כנה"ג |  |
| 7:10.59 | מתקני |  |
| 7:10.60 | התפלות |  |
| 7:10.61 | והברכות | , |
| 7:10.62 | וכמ"ש |  |
| 7:10.63 | רז"ל |  |
| 7:10.64 | בספרי |  |
| 7:10.65 | ומדרש |  |
| 7:10.66 | רבה |  |
| 7:10.67 | שאפילו |  |
| 7:10.68 | מלאכים |  |
| 7:10.69 | וחיות |  |
| 7:10.70 | הקדש |  |
| 7:10.71 | שואלים |  |
| 7:10.72 | איה |  |
| 7:10.73 | מקום |  |
| 7:10.74 | כבודו |  |
| 7:10.75 | כמו |  |
| 7:10.76 | שכ | ' |
| 7:10.77 | לעיל |  |
| 7:10.78 | סי | ' |
| 7:10.79 | נ"ב |  |
| 7:10.80 | בשם |  |
| 7:10.81 | הרקח | . |
| 7:10.82 | וחדשים |  |
| 7:10.83 | מקרוב |  |
| 7:10.84 | באו |  |
| 7:10.85 | לערך |  |
| 7:10.86 | מאה |  |
| 7:10.87 | שנה |  |
| 7:10.88 | או |  |
| 7:10.89 | יותר |  |
| 7:10.90 | מעט |  |
| 7:10.91 | שינו |  |
| 7:10.92 | הנוסח |  |
| 7:10.93 | שלא |  |
| 7:10.94 | לומר |  |
| 7:10.95 | ומשרתיו |  |
| 7:10.96 | שואלים |  |
| 7:10.97 | זה |  |
| 7:10.98 | לזה |  |
| 7:10.99 | איה |  |
| 7:10.100 | מקום |  |
| 7:10.101 | כבודו |  |
| 7:10.102 | רק |  |
| 7:10.103 | במוסף |  |
| 7:10.104 | שבת |  |
| 7:10.105 | בלבד |  |
| 7:10.106 | מקפידים |  |
| 7:10.107 | על |  |
| 7:10.108 | האומרו |  |
| 7:10.109 | בשאר |  |
| 7:10.110 | ימים |  |
| 7:10.111 | ומשתקים |  |
| 7:10.112 | אותו | . |
| 7:10.113 | וטעמם |  |
| 7:10.114 | כמו |  |
| 7:10.115 | שכ | ' |
| 7:10.116 | הר | ' |
| 7:10.117 | משנת |  |
| 7:10.118 | חסידים |  |
| 7:10.119 | במסכת |  |
| 7:10.120 | ר"ח |  |
| 7:10.121 | פ"ג |  |
| 7:10.122 | משנה |  |
| 7:10.123 | ט | ' |
| 7:10.124 | וז"ל |  |
| 7:10.125 | ומפני |  |
| 7:10.126 | שאין |  |
| 7:10.127 | לזעיר |  |
| 7:10.128 | ונוקביה | . |
| 7:10.129 | עלייה |  |
| 7:10.130 | במוסף |  |
| 7:10.131 | ר"ח | ) |
| 7:10.132 | עתה |  |
| 7:10.133 | למעלה |  |
| 7:10.134 | ממקום |  |
| 7:10.135 | אצילותם |  |
| 7:10.136 | כמו |  |
| 7:10.137 | בשבת | . |
| 7:10.138 | לפיכך |  |
| 7:10.139 | אין |  |
| 7:10.140 | אומרים |  |
| 7:10.141 | ומשרתיו |  |
| 7:10.142 | שואלים |  |
| 7:10.143 | זה |  |
| 7:10.144 | לזה |  |
| 7:10.145 | איה |  |
| 7:10.146 | מקום |  |
| 7:10.147 | כבודו | . |
| 7:10.148 | כי |  |
| 7:10.149 | משרתיו |  |
| 7:10.150 | יודעים |  |
| 7:10.151 | מקומו |  |
| 7:10.152 | עכ"ל | . |
| 7:10.153 | הרי |  |
| 7:10.154 | בהדיא |  |
| 7:10.155 | מבואר |  |
| 7:10.156 | בדבריו |  |
| 7:10.157 | דסבירא |  |
| 7:10.158 | להו |  |
| 7:10.159 | לבעלי |  |
| 7:10.160 | הקבלה |  |
| 7:10.161 | החדשה |  |
| 7:10.162 | כי |  |
| 7:10.163 | מאמר |  |
| 7:10.164 | רבותי | ' |
| 7:10.165 | מסדרי |  |
| 7:10.166 | התפלות |  |
| 7:10.167 | ובעלי |  |
| 7:10.168 | התלמוד |  |
| 7:10.169 | שהקב"ה |  |
| 7:10.170 | נעלם |  |
| 7:10.171 | ורואה |  |
| 7:10.172 | ואינו |  |
| 7:10.173 | נראה |  |
| 7:10.174 | עד |  |
| 7:10.175 | שאפילו |  |
| 7:10.176 | מלאכים |  |
| 7:10.177 | וחיות |  |
| 7:10.178 | הקדש |  |
| 7:10.179 | שואלים |  |
| 7:10.180 | איה |  |
| 7:10.181 | מקום |  |
| 7:10.182 | כבודו | , |
| 7:10.183 | הוא |  |
| 7:10.184 | מצד |  |
| 7:10.185 | העתקו |  |
| 7:10.186 | בזמן |  |
| 7:10.187 | מה |  |
| 7:10.188 | לעלות |  |
| 7:10.189 | למעלה |  |
| 7:10.190 | ממקום |  |
| 7:10.191 | משכן |  |
| 7:10.192 | כבודו | . |
| 7:10.193 | דהיינו |  |
| 7:10.194 | במוסף |  |
| 7:10.195 | שבת |  |
| 7:10.196 | דוקא |  |
| 7:10.197 | שאז |  |
| 7:10.198 | נעתק |  |
| 7:10.199 | ועלה |  |
| 7:10.200 | למעלה | . |
| 7:10.201 | ואז |  |
| 7:10.202 | כשרואים |  |
| 7:10.203 | מקומו |  |
| 7:10.204 | המורגל |  |
| 7:10.205 | להיות |  |
| 7:10.206 | שוכן |  |
| 7:10.207 | שם |  |
| 7:10.208 | פנוי | . |
| 7:10.209 | אז |  |
| 7:10.210 | יפקד |  |
| 7:10.211 | מקום |  |
| 7:10.212 | מושבו |  |
| 7:10.213 | ואז |  |
| 7:10.214 | שואלים |  |
| 7:10.215 | איה |  |
| 7:10.216 | מקום |  |
| 7:10.217 | כבודו | . |
| 7:10.218 | אמנם |  |
| 7:10.219 | בשאר |  |
| 7:10.220 | זמנים |  |
| 7:10.221 | שהוא |  |
| 7:10.222 | שוכן |  |
| 7:10.223 | במקומו |  |
| 7:10.224 | כל |  |
| 7:10.225 | משרתיו |  |
| 7:10.226 | יודעים |  |
| 7:10.227 | מקום |  |
| 7:10.228 | משכן |  |
| 7:10.229 | כבודו |  |
| 7:10.230 | ואין |  |
| 7:10.231 | צריכים |  |
| 7:10.232 | לשאול |  |
| 7:10.233 | עליו |  |
| 7:10.234 | כמאמר |  |
| 7:10.235 | יהונתן |  |
| 7:10.236 | בן |  |
| 7:10.237 | שאול |  |
| 7:10.238 | לדוד |  |
| 7:10.239 | מחר |  |
| 7:10.240 | חדש |  |
| 7:10.241 | ונפקדת |  |
| 7:10.242 | כי |  |
| 7:10.243 | יפקד |  |
| 7:10.244 | מושבך | . |
| 7:10.245 | ודבריהם |  |
| 7:10.246 | אלו |  |
| 7:10.247 | סותרים |  |
| 7:10.248 | העיקר |  |
| 7:10.249 | שהזכירו |  |
| 7:10.250 | הראשונים |  |
| 7:10.251 | האמור |  |
| 7:10.252 | ביגדל |  |
| 7:10.253 | אלהים |  |
| 7:10.254 | חי |  |
| 7:10.255 | ומבואר |  |
| 7:10.256 | ג"כ |  |
| 7:10.257 | בדברי |  |
| 7:10.258 | הרמב"ם |  |
| 7:10.259 | בפ | ' |
| 7:10.260 | א | ' |
| 7:10.261 | מיסודי |  |
| 7:10.262 | התורה |  |
| 7:10.263 | ובמורה | . |
| 7:10.264 | ובס | ' |
| 7:10.265 | אהל |  |
| 7:10.266 | יעקב |  |
| 7:10.267 | פרשת |  |
| 7:10.268 | וירא |  |
| 7:10.269 | על |  |
| 7:10.270 | מאמר |  |
| 7:10.271 | אברהם |  |
| 7:10.272 | אבינו |  |
| 7:10.273 | לשי"ת |  |
| 7:10.274 | אל |  |
| 7:10.275 | נא |  |
| 7:10.276 | תעבור |  |
| 7:10.277 | מעל |  |
| 7:10.278 | עבדך |  |
| 7:10.279 | כתב |  |
| 7:10.280 | דבהשקפה |  |
| 7:10.281 | ראשונה |  |
| 7:10.282 | יפלא |  |
| 7:10.283 | מאד |  |
| 7:10.284 | על |  |
| 7:10.285 | אברהם |  |
| 7:10.286 | אבינו |  |
| 7:10.287 | איך |  |
| 7:10.288 | ערב |  |
| 7:10.289 | אל |  |
| 7:10.290 | לבו |  |
| 7:10.291 | לומר |  |
| 7:10.292 | לפני |  |
| 7:10.293 | השי"ת |  |
| 7:10.294 | אל |  |
| 7:10.295 | נא |  |
| 7:10.296 | תעבור |  |
| 7:10.297 | מעל |  |
| 7:10.298 | עבדך | . |
| 7:10.299 | ובגמרא |  |
| 7:10.300 | שבת |  |
| 7:10.301 | קכ"ז |  |
| 7:10.302 | אר"א |  |
| 7:10.303 | מדת |  |
| 7:10.304 | ב"ו |  |
| 7:10.305 | אין |  |
| 7:10.306 | קטן |  |
| 7:10.307 | יכול |  |
| 7:10.308 | לומר |  |
| 7:10.309 | לגדול |  |
| 7:10.310 | המתן |  |
| 7:10.311 | עד |  |
| 7:10.312 | שאבוא |  |
| 7:10.313 | אצלך | , |
| 7:10.314 | ואלו |  |
| 7:10.315 | בהקב"ה |  |
| 7:10.316 | כתיב |  |
| 7:10.317 | ויאמר |  |
| 7:10.318 | אל |  |
| 7:10.319 | נא |  |
| 7:10.320 | תעבר |  |
| 7:10.321 | מעל |  |
| 7:10.322 | עבדך | . |
| 7:10.323 | אבל |  |
| 7:10.324 | הבנת |  |
| 7:10.325 | הענין |  |
| 7:10.326 | הוא |  |
| 7:10.327 | עם |  |
| 7:10.328 | המאומת |  |
| 7:10.329 | והמקובל |  |
| 7:10.330 | אצלינו |  |
| 7:10.331 | מפי |  |
| 7:10.332 | קדמונינו |  |
| 7:10.333 | ע"ה |  |
| 7:10.334 | ובראשם |  |
| 7:10.335 | הר | ' |
| 7:10.336 | המורה | , |
| 7:10.337 | כי |  |
| 7:10.338 | הקב"ה |  |
| 7:10.339 | נעלה |  |
| 7:10.340 | ומרומם |  |
| 7:10.341 | מכל |  |
| 7:10.342 | שינוי |  |
| 7:10.343 | ותנועה |  |
| 7:10.344 | והעתק |  |
| 7:10.345 | מקום |  |
| 7:10.346 | ובשגם |  |
| 7:10.347 | מצינו |  |
| 7:10.348 | בפרשת |  |
| 7:10.349 | וירא |  |
| 7:10.350 | וילך |  |
| 7:10.351 | ה | ' |
| 7:10.352 | וכו | ' |
| 7:10.353 | אין |  |
| 7:10.354 | הרצון |  |
| 7:10.355 | בו |  |
| 7:10.356 | חלילה |  |
| 7:10.357 | שהקב"ה |  |
| 7:10.358 | נעתק |  |
| 7:10.359 | וסר |  |
| 7:10.360 | ממקומו |  |
| 7:10.361 | ובא |  |
| 7:10.362 | להתראות |  |
| 7:10.363 | אל |  |
| 7:10.364 | מקום |  |
| 7:10.365 | מצב |  |
| 7:10.366 | מי |  |
| 7:10.367 | שיראה |  |
| 7:10.368 | אליו | . |
| 7:10.369 | אבל |  |
| 7:10.370 | כל |  |
| 7:10.371 | התקרבותו |  |
| 7:10.372 | ית | ' |
| 7:10.373 | והתגלותו |  |
| 7:10.374 | לאדם |  |
| 7:10.375 | תלוי |  |
| 7:10.376 | אך |  |
| 7:10.377 | בהתקרבות |  |
| 7:10.378 | האדם |  |
| 7:10.379 | אליו |  |
| 7:10.380 | כמאמר |  |
| 7:10.381 | אני |  |
| 7:10.382 | לדודי |  |
| 7:10.383 | ודודי |  |
| 7:10.384 | לי | , |
| 7:10.385 | וזהו |  |
| 7:10.386 | שאמר |  |
| 7:10.387 | אני |  |
| 7:10.388 | לדודי |  |
| 7:10.389 | ועלי |  |
| 7:10.390 | תשוקתו | , |
| 7:10.391 | כלומר |  |
| 7:10.392 | תשוקתו |  |
| 7:10.393 | ית | ' |
| 7:10.394 | אלינו |  |
| 7:10.395 | הוא |  |
| 7:10.396 | דבר |  |
| 7:10.397 | תלוי |  |
| 7:10.398 | בידינו | , |
| 7:10.399 | ואברהם |  |
| 7:10.400 | אבינו |  |
| 7:10.401 | בעוצם |  |
| 7:10.402 | דבקותו |  |
| 7:10.403 | בו |  |
| 7:10.404 | ית | ' |
| 7:10.405 | נגלה |  |
| 7:10.406 | אליו | , |
| 7:10.407 | ובהזדמן |  |
| 7:10.408 | האורחים |  |
| 7:10.409 | אצלו |  |
| 7:10.410 | וצידד |  |
| 7:10.411 | להכניסם |  |
| 7:10.412 | ולטבוח |  |
| 7:10.413 | טבח |  |
| 7:10.414 | והכן |  |
| 7:10.415 | ירא |  |
| 7:10.416 | לנפשו |  |
| 7:10.417 | שלא |  |
| 7:10.418 | יפול |  |
| 7:10.419 | מהדבק |  |
| 7:10.420 | בשי"ת |  |
| 7:10.421 | ושאל |  |
| 7:10.422 | ממנו |  |
| 7:10.423 | אל |  |
| 7:10.424 | נא |  |
| 7:10.425 | תעבור |  |
| 7:10.426 | מעל |  |
| 7:10.427 | עבדך | , |
| 7:10.428 | לבל |  |
| 7:10.429 | יסיר |  |
| 7:10.430 | הקב"ה |  |
| 7:10.431 | עזרתו |  |
| 7:10.432 | מאתו |  |
| 7:10.433 | ולא |  |
| 7:10.434 | ינתק |  |
| 7:10.435 | מדבקותו |  |
| 7:10.436 | בטרדת |  |
| 7:10.437 | הכנסת |  |
| 7:10.438 | אורחים |  |
| 7:10.439 | וכו | ' |
| 7:10.440 | עיין |  |
| 7:10.441 | שם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

