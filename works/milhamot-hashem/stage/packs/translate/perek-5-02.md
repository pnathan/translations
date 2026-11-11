# Translation pack: milhamot-hashem perek-5-02 (Milchamot ha-Shem 5:2–5:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-02`.

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
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- דברי תימה: "astonishing words" (words that astonish).
- והוא שלא: "provided that" (fixed).
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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


## Milchamot ha-Shem 5:2

Hebrew (Wikisource): מא) ועל התגדלך והתפארך בעצמך, והתרוממת על מרי שמיא לדרוש בהמון כי ידעתו שבתו וקומו של ממ"ה הקב"ה וליטול את השם ולאמר אבל קבלתנו וכו', הנה לא הודעתנו אל איזה א"ס אתה מתפלל, כי כמה מיני א"ס יש, וגם לא בארת ממי קבלת קבלה זו. ומעיקרא לא על קבלתך שאלך השואל, ולמה לך להשיב שלא כענין, היפך דברי הזהר והמקובלים כמ"ש לעיל, וכמ"ש בהדיא בכסא אליהו דף ט"ו ע"ב וז"ל: "אמנם ענין הברכה לו ית"ש אין הכוונה לעצמות אדון יחיד ב"ה חלילה חלילה, כי הוא מרומם מעל כל ברכה" ע"כ, וכ"כ ספר הברית במאמר שני כי הא"ס בבחינתו הפשוטה אשר דיברו הפילוסופים (הכופרים) שלא שייך בו שום עבודה ותפילה, על כי הוא בבחינה זאת נעלה ומרומם מכל ברכה ותהילה, ובטלה כל מצוה וכל תורה ועבודה, עי"ש ועוד יתבאר לפנינו. וכיון שקבלתך זו שהזכרת היא היפך דברי הזהר והמקובלים מקדש מלך וכסא אליהו ומצרף האמונה ויושר לבב ומשנת חסידים והרש"ב ומהרח"ו בעץ חיים וספר הכוונות ומחברת הקודש וספר הברית והאר"י בספר הליקוטים, א"כ הוי שוטה וסכל יקראו לך! על מה אתה נלחם בלומדי תורה ברכילות ומלשינות ולה"ר? ועל מה אתה מבזה אותם? העל שאינם מהגים ומצפצפים בספר הזהר כמוך? רק מחזיקים במשנה ותלמוד ופוסקים? הן על זה פירשו מללמוד בספרי הקבלה החדשה כי רק מיד בני נכר הגיעה אלינו כמבואר בספר שם הגדולים להר"ר חיד"א. ובגיליון הבא מירושלים בשם הרב הנזכר ובשם הרב יש"ר וכמו שיתבאר להלן. ולדעת רבים וגדולים מחכמינו הוי מינות. ולכן קבענו כל לימודנו במשנה ותלמוד ופוסקים, כהוראת רב האי גאון המובא בתשובת הרשב"א ובמבוא התלמוד שאין טוב לישראל כעסק המשנה והתלמוד, זה לנו יותר משלושים שנה. ובזדון לבכם השמעתם בקהל רב כי אנחנו מינים וכופרים על שאין אנו לומדים בזהר ותיקונים. הן עתה נתברר לנו מתשובותיכם כי לא ידעתם מספרי ישראל מאומה, לא במדרשי רז"ל ולא בתלמוד ולא בקבלה החדשה. כמאמר ר' יהודה בן תיבון בפתיחה לשער השני מספר חובת הלבבות בשם מבחר הפנינים: "מי שמתחכם מבלי חכמה הוא כחמור הריחיים, סובב סובב ואינו זז ממקומו". עוד שם: "אל תאמר במה שלא ידעת - ידעתי, פן תחשד במה שידעת, ואמרו החכמים: מי שסכל וחשב שידע סכלותו כפולה", עכ"ל. הוי רועים רעים המטעים את העם בשקריהם, ומחשיבים עצמם לחכמים בעיניהם! חכמים אתם להרע ולחלל את השי"ת ואת תורתו! ובדת אמת אשר לא ידעתם אותה על בריה תסחרו, להיות לכם לקרדום לחפור בה בורות נשברים אשר לא יכילו המים, כמאמר הנביא (מיכה ג' ה'): "כה אמר ה' על הנביאים המתעים את עמי, הנשכים בשניהם וקראו שלום. ואשר לא יתן אל פיהם, וקדשו עליו מלחמה". ומתשובתך ניכר שלא ידעת רק אין סוף אחד המקיף את כל המציאות. ולא ידעת שיש כמה מיני אין סוף בדברי המקובלים, ושאין לכוון בהתפללו לזעיר אנפין רק להמשיך מאין סוף הגנוז בכתר בלבד. אבל האין סוף אשר הוא לעילא לעילא לית מחשבה תפיסה ביה כלל ואין לכוון תפילתו אליו, כמ"ש לעיל בשם הרח"ו בספר ארבע מאות שקל כסף וספר הברית.

| id | word | punctuation after |
|---|---|---|
| 5:2.1 | מא | ) |
| 5:2.2 | ועל |  |
| 5:2.3 | התגדלך |  |
| 5:2.4 | והתפארך |  |
| 5:2.5 | בעצמך | , |
| 5:2.6 | והתרוממת |  |
| 5:2.7 | על |  |
| 5:2.8 | מרי |  |
| 5:2.9 | שמיא |  |
| 5:2.10 | לדרוש |  |
| 5:2.11 | בהמון |  |
| 5:2.12 | כי |  |
| 5:2.13 | ידעתו |  |
| 5:2.14 | שבתו |  |
| 5:2.15 | וקומו |  |
| 5:2.16 | של |  |
| 5:2.17 | ממ"ה |  |
| 5:2.18 | הקב"ה |  |
| 5:2.19 | וליטול |  |
| 5:2.20 | את |  |
| 5:2.21 | השם |  |
| 5:2.22 | ולאמר |  |
| 5:2.23 | אבל |  |
| 5:2.24 | קבלתנו |  |
| 5:2.25 | וכו | ', |
| 5:2.26 | הנה |  |
| 5:2.27 | לא |  |
| 5:2.28 | הודעתנו |  |
| 5:2.29 | אל |  |
| 5:2.30 | איזה |  |
| 5:2.31 | א"ס |  |
| 5:2.32 | אתה |  |
| 5:2.33 | מתפלל | , |
| 5:2.34 | כי |  |
| 5:2.35 | כמה |  |
| 5:2.36 | מיני |  |
| 5:2.37 | א"ס |  |
| 5:2.38 | יש | , |
| 5:2.39 | וגם |  |
| 5:2.40 | לא |  |
| 5:2.41 | בארת |  |
| 5:2.42 | ממי |  |
| 5:2.43 | קבלת |  |
| 5:2.44 | קבלה |  |
| 5:2.45 | זו | . |
| 5:2.46 | ומעיקרא |  |
| 5:2.47 | לא |  |
| 5:2.48 | על |  |
| 5:2.49 | קבלתך |  |
| 5:2.50 | שאלך |  |
| 5:2.51 | השואל | , |
| 5:2.52 | ולמה |  |
| 5:2.53 | לך |  |
| 5:2.54 | להשיב |  |
| 5:2.55 | שלא |  |
| 5:2.56 | כענין | , |
| 5:2.57 | היפך |  |
| 5:2.58 | דברי |  |
| 5:2.59 | הזהר |  |
| 5:2.60 | והמקובלים |  |
| 5:2.61 | כמ"ש |  |
| 5:2.62 | לעיל | , |
| 5:2.63 | וכמ"ש |  |
| 5:2.64 | בהדיא |  |
| 5:2.65 | בכסא |  |
| 5:2.66 | אליהו |  |
| 5:2.67 | דף |  |
| 5:2.68 | ט"ו |  |
| 5:2.69 | ע"ב |  |
| 5:2.70 | וז"ל | : |
| 5:2.71 | אמנם |  |
| 5:2.72 | ענין |  |
| 5:2.73 | הברכה |  |
| 5:2.74 | לו |  |
| 5:2.75 | ית"ש |  |
| 5:2.76 | אין |  |
| 5:2.77 | הכוונה |  |
| 5:2.78 | לעצמות |  |
| 5:2.79 | אדון |  |
| 5:2.80 | יחיד |  |
| 5:2.81 | ב"ה |  |
| 5:2.82 | חלילה |  |
| 5:2.83 | חלילה | , |
| 5:2.84 | כי |  |
| 5:2.85 | הוא |  |
| 5:2.86 | מרומם |  |
| 5:2.87 | מעל |  |
| 5:2.88 | כל |  |
| 5:2.89 | ברכה | " |
| 5:2.90 | ע"כ | , |
| 5:2.91 | וכ"כ |  |
| 5:2.92 | ספר |  |
| 5:2.93 | הברית |  |
| 5:2.94 | במאמר |  |
| 5:2.95 | שני |  |
| 5:2.96 | כי |  |
| 5:2.97 | הא"ס |  |
| 5:2.98 | בבחינתו |  |
| 5:2.99 | הפשוטה |  |
| 5:2.100 | אשר |  |
| 5:2.101 | דיברו |  |
| 5:2.102 | הפילוסופים |  |
| 5:2.103 | הכופרים | ) |
| 5:2.104 | שלא |  |
| 5:2.105 | שייך |  |
| 5:2.106 | בו |  |
| 5:2.107 | שום |  |
| 5:2.108 | עבודה |  |
| 5:2.109 | ותפילה | , |
| 5:2.110 | על |  |
| 5:2.111 | כי |  |
| 5:2.112 | הוא |  |
| 5:2.113 | בבחינה |  |
| 5:2.114 | זאת |  |
| 5:2.115 | נעלה |  |
| 5:2.116 | ומרומם |  |
| 5:2.117 | מכל |  |
| 5:2.118 | ברכה |  |
| 5:2.119 | ותהילה | , |
| 5:2.120 | ובטלה |  |
| 5:2.121 | כל |  |
| 5:2.122 | מצוה |  |
| 5:2.123 | וכל |  |
| 5:2.124 | תורה |  |
| 5:2.125 | ועבודה | , |
| 5:2.126 | עי"ש |  |
| 5:2.127 | ועוד |  |
| 5:2.128 | יתבאר |  |
| 5:2.129 | לפנינו | . |
| 5:2.130 | וכיון |  |
| 5:2.131 | שקבלתך |  |
| 5:2.132 | זו |  |
| 5:2.133 | שהזכרת |  |
| 5:2.134 | היא |  |
| 5:2.135 | היפך |  |
| 5:2.136 | דברי |  |
| 5:2.137 | הזהר |  |
| 5:2.138 | והמקובלים |  |
| 5:2.139 | מקדש |  |
| 5:2.140 | מלך |  |
| 5:2.141 | וכסא |  |
| 5:2.142 | אליהו |  |
| 5:2.143 | ומצרף |  |
| 5:2.144 | האמונה |  |
| 5:2.145 | ויושר |  |
| 5:2.146 | לבב |  |
| 5:2.147 | ומשנת |  |
| 5:2.148 | חסידים |  |
| 5:2.149 | והרש"ב |  |
| 5:2.150 | ומהרח"ו |  |
| 5:2.151 | בעץ |  |
| 5:2.152 | חיים |  |
| 5:2.153 | וספר |  |
| 5:2.154 | הכוונות |  |
| 5:2.155 | ומחברת |  |
| 5:2.156 | הקודש |  |
| 5:2.157 | וספר |  |
| 5:2.158 | הברית |  |
| 5:2.159 | והאר"י |  |
| 5:2.160 | בספר |  |
| 5:2.161 | הליקוטים | , |
| 5:2.162 | א"כ |  |
| 5:2.163 | הוי |  |
| 5:2.164 | שוטה |  |
| 5:2.165 | וסכל |  |
| 5:2.166 | יקראו |  |
| 5:2.167 | לך | ! |
| 5:2.168 | על |  |
| 5:2.169 | מה |  |
| 5:2.170 | אתה |  |
| 5:2.171 | נלחם |  |
| 5:2.172 | בלומדי |  |
| 5:2.173 | תורה |  |
| 5:2.174 | ברכילות |  |
| 5:2.175 | ומלשינות |  |
| 5:2.176 | ולה"ר | ? |
| 5:2.177 | ועל |  |
| 5:2.178 | מה |  |
| 5:2.179 | אתה |  |
| 5:2.180 | מבזה |  |
| 5:2.181 | אותם | ? |
| 5:2.182 | העל |  |
| 5:2.183 | שאינם |  |
| 5:2.184 | מהגים |  |
| 5:2.185 | ומצפצפים |  |
| 5:2.186 | בספר |  |
| 5:2.187 | הזהר |  |
| 5:2.188 | כמוך | ? |
| 5:2.189 | רק |  |
| 5:2.190 | מחזיקים |  |
| 5:2.191 | במשנה |  |
| 5:2.192 | ותלמוד |  |
| 5:2.193 | ופוסקים | ? |
| 5:2.194 | הן |  |
| 5:2.195 | על |  |
| 5:2.196 | זה |  |
| 5:2.197 | פירשו |  |
| 5:2.198 | מללמוד |  |
| 5:2.199 | בספרי |  |
| 5:2.200 | הקבלה |  |
| 5:2.201 | החדשה |  |
| 5:2.202 | כי |  |
| 5:2.203 | רק |  |
| 5:2.204 | מיד |  |
| 5:2.205 | בני |  |
| 5:2.206 | נכר |  |
| 5:2.207 | הגיעה |  |
| 5:2.208 | אלינו |  |
| 5:2.209 | כמבואר |  |
| 5:2.210 | בספר |  |
| 5:2.211 | שם |  |
| 5:2.212 | הגדולים |  |
| 5:2.213 | להר"ר |  |
| 5:2.214 | חיד"א | . |
| 5:2.215 | ובגיליון |  |
| 5:2.216 | הבא |  |
| 5:2.217 | מירושלים |  |
| 5:2.218 | בשם |  |
| 5:2.219 | הרב |  |
| 5:2.220 | הנזכר |  |
| 5:2.221 | ובשם |  |
| 5:2.222 | הרב |  |
| 5:2.223 | יש"ר |  |
| 5:2.224 | וכמו |  |
| 5:2.225 | שיתבאר |  |
| 5:2.226 | להלן | . |
| 5:2.227 | ולדעת |  |
| 5:2.228 | רבים |  |
| 5:2.229 | וגדולים |  |
| 5:2.230 | מחכמינו |  |
| 5:2.231 | הוי |  |
| 5:2.232 | מינות | . |
| 5:2.233 | ולכן |  |
| 5:2.234 | קבענו |  |
| 5:2.235 | כל |  |
| 5:2.236 | לימודנו |  |
| 5:2.237 | במשנה |  |
| 5:2.238 | ותלמוד |  |
| 5:2.239 | ופוסקים | , |
| 5:2.240 | כהוראת |  |
| 5:2.241 | רב |  |
| 5:2.242 | האי |  |
| 5:2.243 | גאון |  |
| 5:2.244 | המובא |  |
| 5:2.245 | בתשובת |  |
| 5:2.246 | הרשב"א |  |
| 5:2.247 | ובמבוא |  |
| 5:2.248 | התלמוד |  |
| 5:2.249 | שאין |  |
| 5:2.250 | טוב |  |
| 5:2.251 | לישראל |  |
| 5:2.252 | כעסק |  |
| 5:2.253 | המשנה |  |
| 5:2.254 | והתלמוד | , |
| 5:2.255 | זה |  |
| 5:2.256 | לנו |  |
| 5:2.257 | יותר |  |
| 5:2.258 | משלושים |  |
| 5:2.259 | שנה | . |
| 5:2.260 | ובזדון |  |
| 5:2.261 | לבכם |  |
| 5:2.262 | השמעתם |  |
| 5:2.263 | בקהל |  |
| 5:2.264 | רב |  |
| 5:2.265 | כי |  |
| 5:2.266 | אנחנו |  |
| 5:2.267 | מינים |  |
| 5:2.268 | וכופרים |  |
| 5:2.269 | על |  |
| 5:2.270 | שאין |  |
| 5:2.271 | אנו |  |
| 5:2.272 | לומדים |  |
| 5:2.273 | בזהר |  |
| 5:2.274 | ותיקונים | . |
| 5:2.275 | הן |  |
| 5:2.276 | עתה |  |
| 5:2.277 | נתברר |  |
| 5:2.278 | לנו |  |
| 5:2.279 | מתשובותיכם |  |
| 5:2.280 | כי |  |
| 5:2.281 | לא |  |
| 5:2.282 | ידעתם |  |
| 5:2.283 | מספרי |  |
| 5:2.284 | ישראל |  |
| 5:2.285 | מאומה | , |
| 5:2.286 | לא |  |
| 5:2.287 | במדרשי |  |
| 5:2.288 | רז"ל |  |
| 5:2.289 | ולא |  |
| 5:2.290 | בתלמוד |  |
| 5:2.291 | ולא |  |
| 5:2.292 | בקבלה |  |
| 5:2.293 | החדשה | . |
| 5:2.294 | כמאמר |  |
| 5:2.295 | ר | ' |
| 5:2.296 | יהודה |  |
| 5:2.297 | בן |  |
| 5:2.298 | תיבון |  |
| 5:2.299 | בפתיחה |  |
| 5:2.300 | לשער |  |
| 5:2.301 | השני |  |
| 5:2.302 | מספר |  |
| 5:2.303 | חובת |  |
| 5:2.304 | הלבבות |  |
| 5:2.305 | בשם |  |
| 5:2.306 | מבחר |  |
| 5:2.307 | הפנינים | : |
| 5:2.308 | מי |  |
| 5:2.309 | שמתחכם |  |
| 5:2.310 | מבלי |  |
| 5:2.311 | חכמה |  |
| 5:2.312 | הוא |  |
| 5:2.313 | כחמור |  |
| 5:2.314 | הריחיים | , |
| 5:2.315 | סובב |  |
| 5:2.316 | סובב |  |
| 5:2.317 | ואינו |  |
| 5:2.318 | זז |  |
| 5:2.319 | ממקומו | ". |
| 5:2.320 | עוד |  |
| 5:2.321 | שם | : |
| 5:2.322 | אל |  |
| 5:2.323 | תאמר |  |
| 5:2.324 | במה |  |
| 5:2.325 | שלא |  |
| 5:2.326 | ידעת | - |
| 5:2.327 | ידעתי | , |
| 5:2.328 | פן |  |
| 5:2.329 | תחשד |  |
| 5:2.330 | במה |  |
| 5:2.331 | שידעת | , |
| 5:2.332 | ואמרו |  |
| 5:2.333 | החכמים | : |
| 5:2.334 | מי |  |
| 5:2.335 | שסכל |  |
| 5:2.336 | וחשב |  |
| 5:2.337 | שידע |  |
| 5:2.338 | סכלותו |  |
| 5:2.339 | כפולה | ", |
| 5:2.340 | עכ"ל | . |
| 5:2.341 | הוי |  |
| 5:2.342 | רועים |  |
| 5:2.343 | רעים |  |
| 5:2.344 | המטעים |  |
| 5:2.345 | את |  |
| 5:2.346 | העם |  |
| 5:2.347 | בשקריהם | , |
| 5:2.348 | ומחשיבים |  |
| 5:2.349 | עצמם |  |
| 5:2.350 | לחכמים |  |
| 5:2.351 | בעיניהם | ! |
| 5:2.352 | חכמים |  |
| 5:2.353 | אתם |  |
| 5:2.354 | להרע |  |
| 5:2.355 | ולחלל |  |
| 5:2.356 | את |  |
| 5:2.357 | השי"ת |  |
| 5:2.358 | ואת |  |
| 5:2.359 | תורתו | ! |
| 5:2.360 | ובדת |  |
| 5:2.361 | אמת |  |
| 5:2.362 | אשר |  |
| 5:2.363 | לא |  |
| 5:2.364 | ידעתם |  |
| 5:2.365 | אותה |  |
| 5:2.366 | על |  |
| 5:2.367 | בריה |  |
| 5:2.368 | תסחרו | , |
| 5:2.369 | להיות |  |
| 5:2.370 | לכם |  |
| 5:2.371 | לקרדום |  |
| 5:2.372 | לחפור |  |
| 5:2.373 | בה |  |
| 5:2.374 | בורות |  |
| 5:2.375 | נשברים |  |
| 5:2.376 | אשר |  |
| 5:2.377 | לא |  |
| 5:2.378 | יכילו |  |
| 5:2.379 | המים | , |
| 5:2.380 | כמאמר |  |
| 5:2.381 | הנביא |  |
| 5:2.382 | מיכה |  |
| 5:2.383 | ג | ' |
| 5:2.384 | ה | '): |
| 5:2.385 | כה |  |
| 5:2.386 | אמר |  |
| 5:2.387 | ה | ' |
| 5:2.388 | על |  |
| 5:2.389 | הנביאים |  |
| 5:2.390 | המתעים |  |
| 5:2.391 | את |  |
| 5:2.392 | עמי | , |
| 5:2.393 | הנשכים |  |
| 5:2.394 | בשניהם |  |
| 5:2.395 | וקראו |  |
| 5:2.396 | שלום | . |
| 5:2.397 | ואשר |  |
| 5:2.398 | לא |  |
| 5:2.399 | יתן |  |
| 5:2.400 | אל |  |
| 5:2.401 | פיהם | , |
| 5:2.402 | וקדשו |  |
| 5:2.403 | עליו |  |
| 5:2.404 | מלחמה | ". |
| 5:2.405 | ומתשובתך |  |
| 5:2.406 | ניכר |  |
| 5:2.407 | שלא |  |
| 5:2.408 | ידעת |  |
| 5:2.409 | רק |  |
| 5:2.410 | אין |  |
| 5:2.411 | סוף |  |
| 5:2.412 | אחד |  |
| 5:2.413 | המקיף |  |
| 5:2.414 | את |  |
| 5:2.415 | כל |  |
| 5:2.416 | המציאות | . |
| 5:2.417 | ולא |  |
| 5:2.418 | ידעת |  |
| 5:2.419 | שיש |  |
| 5:2.420 | כמה |  |
| 5:2.421 | מיני |  |
| 5:2.422 | אין |  |
| 5:2.423 | סוף |  |
| 5:2.424 | בדברי |  |
| 5:2.425 | המקובלים | , |
| 5:2.426 | ושאין |  |
| 5:2.427 | לכוון |  |
| 5:2.428 | בהתפללו |  |
| 5:2.429 | לזעיר |  |
| 5:2.430 | אנפין |  |
| 5:2.431 | רק |  |
| 5:2.432 | להמשיך |  |
| 5:2.433 | מאין |  |
| 5:2.434 | סוף |  |
| 5:2.435 | הגנוז |  |
| 5:2.436 | בכתר |  |
| 5:2.437 | בלבד | . |
| 5:2.438 | אבל |  |
| 5:2.439 | האין |  |
| 5:2.440 | סוף |  |
| 5:2.441 | אשר |  |
| 5:2.442 | הוא |  |
| 5:2.443 | לעילא |  |
| 5:2.444 | לעילא |  |
| 5:2.445 | לית |  |
| 5:2.446 | מחשבה |  |
| 5:2.447 | תפיסה |  |
| 5:2.448 | ביה |  |
| 5:2.449 | כלל |  |
| 5:2.450 | ואין |  |
| 5:2.451 | לכוון |  |
| 5:2.452 | תפילתו |  |
| 5:2.453 | אליו | , |
| 5:2.454 | כמ"ש |  |
| 5:2.455 | לעיל |  |
| 5:2.456 | בשם |  |
| 5:2.457 | הרח"ו |  |
| 5:2.458 | בספר |  |
| 5:2.459 | ארבע |  |
| 5:2.460 | מאות |  |
| 5:2.461 | שקל |  |
| 5:2.462 | כסף |  |
| 5:2.463 | וספר |  |
| 5:2.464 | הברית | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

