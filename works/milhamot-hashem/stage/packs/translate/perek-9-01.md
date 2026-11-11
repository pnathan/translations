# Translation pack: milhamot-hashem perek-9-01 (Milchamot ha-Shem 9:1–9:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-01`.

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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- ר"ל with a named authority as subject: "means to say".
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
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


## Milchamot ha-Shem 9:1

Hebrew (Wikisource): קך) וכתב רבינו סעדיה גאון בהקדמת ספרו האמונות והדעות. וז"ל והאמונה על ב' דרכים. אמת ושקר, והאמונה האמתית היא שידע הדבר כאשר הוא וכו' והאמונה השקרית הוא שידע הדבר בהפך מה הוא וכו'! והחכם המשובח מי ששם אמיתת הדברים שרש. ומנהיג דעתו עליהם ועם חכמתו בטוח אל המובטח, ונזהר ממה שצריך להזהר ממנו (פירש הרב שביל האמונה כלומר עם שהוא חכם. בחכמתו הוא ירא בכל מה שהתורה האלהית הגוזר עליו כהפך עיונו, והוא בוטח על הדבר הראוי לבטוח וירא ונזהר מהדבר אשר ראוי להזהר ממנו יען שנפשו יודעת מאד כי יש מביאות גם לדברים שלא נפל דעתו עליהם עד אשר ייטיב עיונו עליהם עד שיסורו כל הספיקות מלבו. ויתברר לו אמתתו) והכסיל המגונה מי ששם דעתו לשורש, וחושב כי אמתות הדברים הולכות אחר דעתו, ועם סכלותו בוטח במה שצריך להזהר ממנו, ונזהר מהבטוח בו, וכאשר אמר החכם (משלי ד') חכם ירא וסר מרע, וכסיל מתעבר ובוטח ע"כ. ואחרי הדברים והאמת האלה תבין ותשכיל כי קבלת המסית מחבר הזהר והנמשכים אחריו לתומם היא בעצמה דעת הפילסופים הקדמונים הרחוקים מן התורה והמצוה אשר צונו ה' אלהינו על ידי משה בחירו, וזקינים ונביאים שקיבלו איש מפי איש עד חיבור המשנה והתלמוד ועד סוף הגאונים הרס"ג ז"ל אשר נתן אל לבו להשיב על כתות הפילסופים התועים ההולכים אחרי ההבל, לבל יתערבו דעותיהם באמונת תורתינו הטהורה וז"ל הכת האחת היא האומרת שהבורא יש לו גשמים רוחניים דקים ונקודות דקות מאוד כאבק דק וכו' שם במאמר א' חידוש. והכת האחרת היא האומרת שבורא הגשמים בראם מעצמו (רוצה לומר שהפריש חלק מעצמותו ומאותו החלק ברא את העולם) מצאתי אלה האנשים לא נתכן להם לכחש בעושה, ועם כל זה לא קבל שכלם היות דבר מלא דבר. וכיון שאין דבר כי אם הבורא, האמינו כי ברא הדברים מן עצמו. ואלה יותר כסילים מן הראשונים. וראיתי לגלות אולתם בשלש עשרה פנים וכו' ע"ש בדברי הגאון ועומק תשובותיו בזה, ותבין להשיב על הטעאות מחבר הזהר והמקובלים החדשים שנטו אחריו. כי שלי שום ספק שהיא לקוחה מאלה השתי דעות הזרות. ועיין בפי' הרמב"ן לשיר השירים בפסוק אפררון עשה לו המלך שלמה מעצי הלבנון שגם הוא טעה בזה ונטה לרעת הזרה הזאת. וחשב ליחסה לרבי אלעזר הגדול שדעתו כדעת אפלטון האומר כי שוא הוא שהבורא ימציא דבר מלא דבר. מעצמותו ית' והוא דרך משל כחומר ליוצר וכברזל לנפח, יצייר ממנו מה שירצה מהכלים. כן הבורא יצייר מזה הנאצל שמים וארץ וזולתם אין זה חסרון בחוקו ית' כשלא יברא דבר אלא מדבר ודבריו אלה היפך המקובל בכל האומה הישראלית, שהקב"ה המציא כל עם ישראל שהיא מקובלת בידם. ויאמינו סברא פחותה כזו שהקב"ה המציא כל המציאות מעצמותו! וכבר העיר הרמב"ם ז"ל במורה נבוכים פרק כ"ו משני על דברי רבי עליאזר והניח בקושיא. ולא דן אותו לחולק על הקבלה האמתית שהיא פשוטה בכל ישראל. ולעניות דעתי נראה כוונת ר' אליעזר באמרו מאוד לבושו ר"ל מרוב כוחו ית' ועצמותו וגאותו המכונה בכתוב בשם לבוש שנ' לבשי עזך ציון, עורי לבשי עז. לבש ה' עז! ונגד זה נקרא חסרון הכח והיכולת בשם לבוש שנאמר ה' מלך גאות לבש (תהילים צ״ג:א׳). וכן העוז שהוא החוזק והיכולת נקרא חסרון הכח והיכולת בשם לבוש בשת, ילבשו, שוטוני כלמה, שונאיך ילבשו בשת ילבשו בשת וכלמה! ורצה ר' אליעזר ע"ה לומר שברוב כחו ית' ועצמתו ברא השמים יש מאין המוחלט, והשמים נקראו כסא, שנאמר השמים כסאי, ומהגרמים השמימים הנקראים שלג, ברא את הארץ, וזהו כי לשלג יאמר הוי ארץ! והרצון יקרא אור, כמו שנאמר ואור פני לא יפילון, ר"ל שלא יבטלו דברי ויעברו על מה שרציתי כן נ"ל. וזו ראיה גמורה (ר"ל סברת הרמב"ם שמיחס לר"א) שהקבלה החדשה לקוחה מדברי הפילסופים הקדמונים, וגם מאמונת השונים המאמינים בשני פועלים פועל הטוב ופועל הרע, סטרא דקדושה, וסטרא אחרא מסאבא, עיין עליהם ועל תשובת הגאון הרס"ג ז"ל במאמר חדוש, ובפרק ז' מאמר האחדות מצאתי בדברי הגאון ראיה וחזוק למה שכתבתי לעיל וז"ל והכת האומרת כי קצת האל זב גוף ורוח, יתחייבו עם שני התשובות האלה כל מה שהתחייבו בו האומרים כי הדברים הנבראים הם מהבורא והכת האומרת שגופו ורוחו מחודשים ויש בו רוח אחרת מהבורא תחייב כי גוף מחודש שב להיות אלוה בשיתוף דבר אלהי בו ומקישים אותו להשראת הכבוד על הר סיני ובעץ הסנה ובאהל מועד, וזה מחייב אותם להאמין כי האלה והסנה וההר אלהים ג"כ ויוספו רעה וכו' ע"כ. ומדברי הגאון תבין אמתת דברינו שכתבתי לעיל כי נכונים הדברים על פי תורתינו הקדושה וקבלת רז"ל וכמו שכתבו הפוסקים דכל דאתאמר בדברי גאון אליו שומעין שדבריהם דברי קבלה.

| id | word | punctuation after |
|---|---|---|
| 9:1.1 | קך | ) |
| 9:1.2 | וכתב |  |
| 9:1.3 | רבינו |  |
| 9:1.4 | סעדיה |  |
| 9:1.5 | גאון |  |
| 9:1.6 | בהקדמת |  |
| 9:1.7 | ספרו |  |
| 9:1.8 | האמונות |  |
| 9:1.9 | והדעות | . |
| 9:1.10 | וז"ל |  |
| 9:1.11 | והאמונה |  |
| 9:1.12 | על |  |
| 9:1.13 | ב | ' |
| 9:1.14 | דרכים | . |
| 9:1.15 | אמת |  |
| 9:1.16 | ושקר | , |
| 9:1.17 | והאמונה |  |
| 9:1.18 | האמתית |  |
| 9:1.19 | היא |  |
| 9:1.20 | שידע |  |
| 9:1.21 | הדבר |  |
| 9:1.22 | כאשר |  |
| 9:1.23 | הוא |  |
| 9:1.24 | וכו | ' |
| 9:1.25 | והאמונה |  |
| 9:1.26 | השקרית |  |
| 9:1.27 | הוא |  |
| 9:1.28 | שידע |  |
| 9:1.29 | הדבר |  |
| 9:1.30 | בהפך |  |
| 9:1.31 | מה |  |
| 9:1.32 | הוא |  |
| 9:1.33 | וכו | '! |
| 9:1.34 | והחכם |  |
| 9:1.35 | המשובח |  |
| 9:1.36 | מי |  |
| 9:1.37 | ששם |  |
| 9:1.38 | אמיתת |  |
| 9:1.39 | הדברים |  |
| 9:1.40 | שרש | . |
| 9:1.41 | ומנהיג |  |
| 9:1.42 | דעתו |  |
| 9:1.43 | עליהם |  |
| 9:1.44 | ועם |  |
| 9:1.45 | חכמתו |  |
| 9:1.46 | בטוח |  |
| 9:1.47 | אל |  |
| 9:1.48 | המובטח | , |
| 9:1.49 | ונזהר |  |
| 9:1.50 | ממה |  |
| 9:1.51 | שצריך |  |
| 9:1.52 | להזהר |  |
| 9:1.53 | ממנו |  |
| 9:1.54 | פירש |  |
| 9:1.55 | הרב |  |
| 9:1.56 | שביל |  |
| 9:1.57 | האמונה |  |
| 9:1.58 | כלומר |  |
| 9:1.59 | עם |  |
| 9:1.60 | שהוא |  |
| 9:1.61 | חכם | . |
| 9:1.62 | בחכמתו |  |
| 9:1.63 | הוא |  |
| 9:1.64 | ירא |  |
| 9:1.65 | בכל |  |
| 9:1.66 | מה |  |
| 9:1.67 | שהתורה |  |
| 9:1.68 | האלהית |  |
| 9:1.69 | הגוזר |  |
| 9:1.70 | עליו |  |
| 9:1.71 | כהפך |  |
| 9:1.72 | עיונו | , |
| 9:1.73 | והוא |  |
| 9:1.74 | בוטח |  |
| 9:1.75 | על |  |
| 9:1.76 | הדבר |  |
| 9:1.77 | הראוי |  |
| 9:1.78 | לבטוח |  |
| 9:1.79 | וירא |  |
| 9:1.80 | ונזהר |  |
| 9:1.81 | מהדבר |  |
| 9:1.82 | אשר |  |
| 9:1.83 | ראוי |  |
| 9:1.84 | להזהר |  |
| 9:1.85 | ממנו |  |
| 9:1.86 | יען |  |
| 9:1.87 | שנפשו |  |
| 9:1.88 | יודעת |  |
| 9:1.89 | מאד |  |
| 9:1.90 | כי |  |
| 9:1.91 | יש |  |
| 9:1.92 | מביאות |  |
| 9:1.93 | גם |  |
| 9:1.94 | לדברים |  |
| 9:1.95 | שלא |  |
| 9:1.96 | נפל |  |
| 9:1.97 | דעתו |  |
| 9:1.98 | עליהם |  |
| 9:1.99 | עד |  |
| 9:1.100 | אשר |  |
| 9:1.101 | ייטיב |  |
| 9:1.102 | עיונו |  |
| 9:1.103 | עליהם |  |
| 9:1.104 | עד |  |
| 9:1.105 | שיסורו |  |
| 9:1.106 | כל |  |
| 9:1.107 | הספיקות |  |
| 9:1.108 | מלבו | . |
| 9:1.109 | ויתברר |  |
| 9:1.110 | לו |  |
| 9:1.111 | אמתתו | ) |
| 9:1.112 | והכסיל |  |
| 9:1.113 | המגונה |  |
| 9:1.114 | מי |  |
| 9:1.115 | ששם |  |
| 9:1.116 | דעתו |  |
| 9:1.117 | לשורש | , |
| 9:1.118 | וחושב |  |
| 9:1.119 | כי |  |
| 9:1.120 | אמתות |  |
| 9:1.121 | הדברים |  |
| 9:1.122 | הולכות |  |
| 9:1.123 | אחר |  |
| 9:1.124 | דעתו | , |
| 9:1.125 | ועם |  |
| 9:1.126 | סכלותו |  |
| 9:1.127 | בוטח |  |
| 9:1.128 | במה |  |
| 9:1.129 | שצריך |  |
| 9:1.130 | להזהר |  |
| 9:1.131 | ממנו | , |
| 9:1.132 | ונזהר |  |
| 9:1.133 | מהבטוח |  |
| 9:1.134 | בו | , |
| 9:1.135 | וכאשר |  |
| 9:1.136 | אמר |  |
| 9:1.137 | החכם |  |
| 9:1.138 | משלי |  |
| 9:1.139 | ד | ') |
| 9:1.140 | חכם |  |
| 9:1.141 | ירא |  |
| 9:1.142 | וסר |  |
| 9:1.143 | מרע | , |
| 9:1.144 | וכסיל |  |
| 9:1.145 | מתעבר |  |
| 9:1.146 | ובוטח |  |
| 9:1.147 | ע"כ | . |
| 9:1.148 | ואחרי |  |
| 9:1.149 | הדברים |  |
| 9:1.150 | והאמת |  |
| 9:1.151 | האלה |  |
| 9:1.152 | תבין |  |
| 9:1.153 | ותשכיל |  |
| 9:1.154 | כי |  |
| 9:1.155 | קבלת |  |
| 9:1.156 | המסית |  |
| 9:1.157 | מחבר |  |
| 9:1.158 | הזהר |  |
| 9:1.159 | והנמשכים |  |
| 9:1.160 | אחריו |  |
| 9:1.161 | לתומם |  |
| 9:1.162 | היא |  |
| 9:1.163 | בעצמה |  |
| 9:1.164 | דעת |  |
| 9:1.165 | הפילסופים |  |
| 9:1.166 | הקדמונים |  |
| 9:1.167 | הרחוקים |  |
| 9:1.168 | מן |  |
| 9:1.169 | התורה |  |
| 9:1.170 | והמצוה |  |
| 9:1.171 | אשר |  |
| 9:1.172 | צונו |  |
| 9:1.173 | ה | ' |
| 9:1.174 | אלהינו |  |
| 9:1.175 | על |  |
| 9:1.176 | ידי |  |
| 9:1.177 | משה |  |
| 9:1.178 | בחירו | , |
| 9:1.179 | וזקינים |  |
| 9:1.180 | ונביאים |  |
| 9:1.181 | שקיבלו |  |
| 9:1.182 | איש |  |
| 9:1.183 | מפי |  |
| 9:1.184 | איש |  |
| 9:1.185 | עד |  |
| 9:1.186 | חיבור |  |
| 9:1.187 | המשנה |  |
| 9:1.188 | והתלמוד |  |
| 9:1.189 | ועד |  |
| 9:1.190 | סוף |  |
| 9:1.191 | הגאונים |  |
| 9:1.192 | הרס"ג |  |
| 9:1.193 | ז"ל |  |
| 9:1.194 | אשר |  |
| 9:1.195 | נתן |  |
| 9:1.196 | אל |  |
| 9:1.197 | לבו |  |
| 9:1.198 | להשיב |  |
| 9:1.199 | על |  |
| 9:1.200 | כתות |  |
| 9:1.201 | הפילסופים |  |
| 9:1.202 | התועים |  |
| 9:1.203 | ההולכים |  |
| 9:1.204 | אחרי |  |
| 9:1.205 | ההבל | , |
| 9:1.206 | לבל |  |
| 9:1.207 | יתערבו |  |
| 9:1.208 | דעותיהם |  |
| 9:1.209 | באמונת |  |
| 9:1.210 | תורתינו |  |
| 9:1.211 | הטהורה |  |
| 9:1.212 | וז"ל |  |
| 9:1.213 | הכת |  |
| 9:1.214 | האחת |  |
| 9:1.215 | היא |  |
| 9:1.216 | האומרת |  |
| 9:1.217 | שהבורא |  |
| 9:1.218 | יש |  |
| 9:1.219 | לו |  |
| 9:1.220 | גשמים |  |
| 9:1.221 | רוחניים |  |
| 9:1.222 | דקים |  |
| 9:1.223 | ונקודות |  |
| 9:1.224 | דקות |  |
| 9:1.225 | מאוד |  |
| 9:1.226 | כאבק |  |
| 9:1.227 | דק |  |
| 9:1.228 | וכו | ' |
| 9:1.229 | שם |  |
| 9:1.230 | במאמר |  |
| 9:1.231 | א | ' |
| 9:1.232 | חידוש | . |
| 9:1.233 | והכת |  |
| 9:1.234 | האחרת |  |
| 9:1.235 | היא |  |
| 9:1.236 | האומרת |  |
| 9:1.237 | שבורא |  |
| 9:1.238 | הגשמים |  |
| 9:1.239 | בראם |  |
| 9:1.240 | מעצמו |  |
| 9:1.241 | רוצה |  |
| 9:1.242 | לומר |  |
| 9:1.243 | שהפריש |  |
| 9:1.244 | חלק |  |
| 9:1.245 | מעצמותו |  |
| 9:1.246 | ומאותו |  |
| 9:1.247 | החלק |  |
| 9:1.248 | ברא |  |
| 9:1.249 | את |  |
| 9:1.250 | העולם | ) |
| 9:1.251 | מצאתי |  |
| 9:1.252 | אלה |  |
| 9:1.253 | האנשים |  |
| 9:1.254 | לא |  |
| 9:1.255 | נתכן |  |
| 9:1.256 | להם |  |
| 9:1.257 | לכחש |  |
| 9:1.258 | בעושה | , |
| 9:1.259 | ועם |  |
| 9:1.260 | כל |  |
| 9:1.261 | זה |  |
| 9:1.262 | לא |  |
| 9:1.263 | קבל |  |
| 9:1.264 | שכלם |  |
| 9:1.265 | היות |  |
| 9:1.266 | דבר |  |
| 9:1.267 | מלא |  |
| 9:1.268 | דבר | . |
| 9:1.269 | וכיון |  |
| 9:1.270 | שאין |  |
| 9:1.271 | דבר |  |
| 9:1.272 | כי |  |
| 9:1.273 | אם |  |
| 9:1.274 | הבורא | , |
| 9:1.275 | האמינו |  |
| 9:1.276 | כי |  |
| 9:1.277 | ברא |  |
| 9:1.278 | הדברים |  |
| 9:1.279 | מן |  |
| 9:1.280 | עצמו | . |
| 9:1.281 | ואלה |  |
| 9:1.282 | יותר |  |
| 9:1.283 | כסילים |  |
| 9:1.284 | מן |  |
| 9:1.285 | הראשונים | . |
| 9:1.286 | וראיתי |  |
| 9:1.287 | לגלות |  |
| 9:1.288 | אולתם |  |
| 9:1.289 | בשלש |  |
| 9:1.290 | עשרה |  |
| 9:1.291 | פנים |  |
| 9:1.292 | וכו | ' |
| 9:1.293 | ע"ש |  |
| 9:1.294 | בדברי |  |
| 9:1.295 | הגאון |  |
| 9:1.296 | ועומק |  |
| 9:1.297 | תשובותיו |  |
| 9:1.298 | בזה | , |
| 9:1.299 | ותבין |  |
| 9:1.300 | להשיב |  |
| 9:1.301 | על |  |
| 9:1.302 | הטעאות |  |
| 9:1.303 | מחבר |  |
| 9:1.304 | הזהר |  |
| 9:1.305 | והמקובלים |  |
| 9:1.306 | החדשים |  |
| 9:1.307 | שנטו |  |
| 9:1.308 | אחריו | . |
| 9:1.309 | כי |  |
| 9:1.310 | שלי |  |
| 9:1.311 | שום |  |
| 9:1.312 | ספק |  |
| 9:1.313 | שהיא |  |
| 9:1.314 | לקוחה |  |
| 9:1.315 | מאלה |  |
| 9:1.316 | השתי |  |
| 9:1.317 | דעות |  |
| 9:1.318 | הזרות | . |
| 9:1.319 | ועיין |  |
| 9:1.320 | בפי | ' |
| 9:1.321 | הרמב"ן |  |
| 9:1.322 | לשיר |  |
| 9:1.323 | השירים |  |
| 9:1.324 | בפסוק |  |
| 9:1.325 | אפררון |  |
| 9:1.326 | עשה |  |
| 9:1.327 | לו |  |
| 9:1.328 | המלך |  |
| 9:1.329 | שלמה |  |
| 9:1.330 | מעצי |  |
| 9:1.331 | הלבנון |  |
| 9:1.332 | שגם |  |
| 9:1.333 | הוא |  |
| 9:1.334 | טעה |  |
| 9:1.335 | בזה |  |
| 9:1.336 | ונטה |  |
| 9:1.337 | לרעת |  |
| 9:1.338 | הזרה |  |
| 9:1.339 | הזאת | . |
| 9:1.340 | וחשב |  |
| 9:1.341 | ליחסה |  |
| 9:1.342 | לרבי |  |
| 9:1.343 | אלעזר |  |
| 9:1.344 | הגדול |  |
| 9:1.345 | שדעתו |  |
| 9:1.346 | כדעת |  |
| 9:1.347 | אפלטון |  |
| 9:1.348 | האומר |  |
| 9:1.349 | כי |  |
| 9:1.350 | שוא |  |
| 9:1.351 | הוא |  |
| 9:1.352 | שהבורא |  |
| 9:1.353 | ימציא |  |
| 9:1.354 | דבר |  |
| 9:1.355 | מלא |  |
| 9:1.356 | דבר | . |
| 9:1.357 | מעצמותו |  |
| 9:1.358 | ית | ' |
| 9:1.359 | והוא |  |
| 9:1.360 | דרך |  |
| 9:1.361 | משל |  |
| 9:1.362 | כחומר |  |
| 9:1.363 | ליוצר |  |
| 9:1.364 | וכברזל |  |
| 9:1.365 | לנפח | , |
| 9:1.366 | יצייר |  |
| 9:1.367 | ממנו |  |
| 9:1.368 | מה |  |
| 9:1.369 | שירצה |  |
| 9:1.370 | מהכלים | . |
| 9:1.371 | כן |  |
| 9:1.372 | הבורא |  |
| 9:1.373 | יצייר |  |
| 9:1.374 | מזה |  |
| 9:1.375 | הנאצל |  |
| 9:1.376 | שמים |  |
| 9:1.377 | וארץ |  |
| 9:1.378 | וזולתם |  |
| 9:1.379 | אין |  |
| 9:1.380 | זה |  |
| 9:1.381 | חסרון |  |
| 9:1.382 | בחוקו |  |
| 9:1.383 | ית | ' |
| 9:1.384 | כשלא |  |
| 9:1.385 | יברא |  |
| 9:1.386 | דבר |  |
| 9:1.387 | אלא |  |
| 9:1.388 | מדבר |  |
| 9:1.389 | ודבריו |  |
| 9:1.390 | אלה |  |
| 9:1.391 | היפך |  |
| 9:1.392 | המקובל |  |
| 9:1.393 | בכל |  |
| 9:1.394 | האומה |  |
| 9:1.395 | הישראלית | , |
| 9:1.396 | שהקב"ה |  |
| 9:1.397 | המציא |  |
| 9:1.398 | כל |  |
| 9:1.399 | עם |  |
| 9:1.400 | ישראל |  |
| 9:1.401 | שהיא |  |
| 9:1.402 | מקובלת |  |
| 9:1.403 | בידם | . |
| 9:1.404 | ויאמינו |  |
| 9:1.405 | סברא |  |
| 9:1.406 | פחותה |  |
| 9:1.407 | כזו |  |
| 9:1.408 | שהקב"ה |  |
| 9:1.409 | המציא |  |
| 9:1.410 | כל |  |
| 9:1.411 | המציאות |  |
| 9:1.412 | מעצמותו | ! |
| 9:1.413 | וכבר |  |
| 9:1.414 | העיר |  |
| 9:1.415 | הרמב"ם |  |
| 9:1.416 | ז"ל |  |
| 9:1.417 | במורה |  |
| 9:1.418 | נבוכים |  |
| 9:1.419 | פרק |  |
| 9:1.420 | כ"ו |  |
| 9:1.421 | משני |  |
| 9:1.422 | על |  |
| 9:1.423 | דברי |  |
| 9:1.424 | רבי |  |
| 9:1.425 | עליאזר |  |
| 9:1.426 | והניח |  |
| 9:1.427 | בקושיא | . |
| 9:1.428 | ולא |  |
| 9:1.429 | דן |  |
| 9:1.430 | אותו |  |
| 9:1.431 | לחולק |  |
| 9:1.432 | על |  |
| 9:1.433 | הקבלה |  |
| 9:1.434 | האמתית |  |
| 9:1.435 | שהיא |  |
| 9:1.436 | פשוטה |  |
| 9:1.437 | בכל |  |
| 9:1.438 | ישראל | . |
| 9:1.439 | ולעניות |  |
| 9:1.440 | דעתי |  |
| 9:1.441 | נראה |  |
| 9:1.442 | כוונת |  |
| 9:1.443 | ר | ' |
| 9:1.444 | אליעזר |  |
| 9:1.445 | באמרו |  |
| 9:1.446 | מאוד |  |
| 9:1.447 | לבושו |  |
| 9:1.448 | ר"ל |  |
| 9:1.449 | מרוב |  |
| 9:1.450 | כוחו |  |
| 9:1.451 | ית | ' |
| 9:1.452 | ועצמותו |  |
| 9:1.453 | וגאותו |  |
| 9:1.454 | המכונה |  |
| 9:1.455 | בכתוב |  |
| 9:1.456 | בשם |  |
| 9:1.457 | לבוש |  |
| 9:1.458 | שנ | ' |
| 9:1.459 | לבשי |  |
| 9:1.460 | עזך |  |
| 9:1.461 | ציון | , |
| 9:1.462 | עורי |  |
| 9:1.463 | לבשי |  |
| 9:1.464 | עז | . |
| 9:1.465 | לבש |  |
| 9:1.466 | ה | ' |
| 9:1.467 | עז | ! |
| 9:1.468 | ונגד |  |
| 9:1.469 | זה |  |
| 9:1.470 | נקרא |  |
| 9:1.471 | חסרון |  |
| 9:1.472 | הכח |  |
| 9:1.473 | והיכולת |  |
| 9:1.474 | בשם |  |
| 9:1.475 | לבוש |  |
| 9:1.476 | שנאמר |  |
| 9:1.477 | ה | ' |
| 9:1.478 | מלך |  |
| 9:1.479 | גאות |  |
| 9:1.480 | לבש |  |
| 9:1.481 | תהילים |  |
| 9:1.482 | צ״ג:א׳ | ). |
| 9:1.483 | וכן |  |
| 9:1.484 | העוז |  |
| 9:1.485 | שהוא |  |
| 9:1.486 | החוזק |  |
| 9:1.487 | והיכולת |  |
| 9:1.488 | נקרא |  |
| 9:1.489 | חסרון |  |
| 9:1.490 | הכח |  |
| 9:1.491 | והיכולת |  |
| 9:1.492 | בשם |  |
| 9:1.493 | לבוש |  |
| 9:1.494 | בשת | , |
| 9:1.495 | ילבשו | , |
| 9:1.496 | שוטוני |  |
| 9:1.497 | כלמה | , |
| 9:1.498 | שונאיך |  |
| 9:1.499 | ילבשו |  |
| 9:1.500 | בשת |  |
| 9:1.501 | ילבשו |  |
| 9:1.502 | בשת |  |
| 9:1.503 | וכלמה | ! |
| 9:1.504 | ורצה |  |
| 9:1.505 | ר | ' |
| 9:1.506 | אליעזר |  |
| 9:1.507 | ע"ה |  |
| 9:1.508 | לומר |  |
| 9:1.509 | שברוב |  |
| 9:1.510 | כחו |  |
| 9:1.511 | ית | ' |
| 9:1.512 | ועצמתו |  |
| 9:1.513 | ברא |  |
| 9:1.514 | השמים |  |
| 9:1.515 | יש |  |
| 9:1.516 | מאין |  |
| 9:1.517 | המוחלט | , |
| 9:1.518 | והשמים |  |
| 9:1.519 | נקראו |  |
| 9:1.520 | כסא | , |
| 9:1.521 | שנאמר |  |
| 9:1.522 | השמים |  |
| 9:1.523 | כסאי | , |
| 9:1.524 | ומהגרמים |  |
| 9:1.525 | השמימים |  |
| 9:1.526 | הנקראים |  |
| 9:1.527 | שלג | , |
| 9:1.528 | ברא |  |
| 9:1.529 | את |  |
| 9:1.530 | הארץ | , |
| 9:1.531 | וזהו |  |
| 9:1.532 | כי |  |
| 9:1.533 | לשלג |  |
| 9:1.534 | יאמר |  |
| 9:1.535 | הוי |  |
| 9:1.536 | ארץ | ! |
| 9:1.537 | והרצון |  |
| 9:1.538 | יקרא |  |
| 9:1.539 | אור | , |
| 9:1.540 | כמו |  |
| 9:1.541 | שנאמר |  |
| 9:1.542 | ואור |  |
| 9:1.543 | פני |  |
| 9:1.544 | לא |  |
| 9:1.545 | יפילון | , |
| 9:1.546 | ר"ל |  |
| 9:1.547 | שלא |  |
| 9:1.548 | יבטלו |  |
| 9:1.549 | דברי |  |
| 9:1.550 | ויעברו |  |
| 9:1.551 | על |  |
| 9:1.552 | מה |  |
| 9:1.553 | שרציתי |  |
| 9:1.554 | כן |  |
| 9:1.555 | נ"ל | . |
| 9:1.556 | וזו |  |
| 9:1.557 | ראיה |  |
| 9:1.558 | גמורה |  |
| 9:1.559 | ר"ל |  |
| 9:1.560 | סברת |  |
| 9:1.561 | הרמב"ם |  |
| 9:1.562 | שמיחס |  |
| 9:1.563 | לר"א | ) |
| 9:1.564 | שהקבלה |  |
| 9:1.565 | החדשה |  |
| 9:1.566 | לקוחה |  |
| 9:1.567 | מדברי |  |
| 9:1.568 | הפילסופים |  |
| 9:1.569 | הקדמונים | , |
| 9:1.570 | וגם |  |
| 9:1.571 | מאמונת |  |
| 9:1.572 | השונים |  |
| 9:1.573 | המאמינים |  |
| 9:1.574 | בשני |  |
| 9:1.575 | פועלים |  |
| 9:1.576 | פועל |  |
| 9:1.577 | הטוב |  |
| 9:1.578 | ופועל |  |
| 9:1.579 | הרע | , |
| 9:1.580 | סטרא |  |
| 9:1.581 | דקדושה | , |
| 9:1.582 | וסטרא |  |
| 9:1.583 | אחרא |  |
| 9:1.584 | מסאבא | , |
| 9:1.585 | עיין |  |
| 9:1.586 | עליהם |  |
| 9:1.587 | ועל |  |
| 9:1.588 | תשובת |  |
| 9:1.589 | הגאון |  |
| 9:1.590 | הרס"ג |  |
| 9:1.591 | ז"ל |  |
| 9:1.592 | במאמר |  |
| 9:1.593 | חדוש | , |
| 9:1.594 | ובפרק |  |
| 9:1.595 | ז | ' |
| 9:1.596 | מאמר |  |
| 9:1.597 | האחדות |  |
| 9:1.598 | מצאתי |  |
| 9:1.599 | בדברי |  |
| 9:1.600 | הגאון |  |
| 9:1.601 | ראיה |  |
| 9:1.602 | וחזוק |  |
| 9:1.603 | למה |  |
| 9:1.604 | שכתבתי |  |
| 9:1.605 | לעיל |  |
| 9:1.606 | וז"ל |  |
| 9:1.607 | והכת |  |
| 9:1.608 | האומרת |  |
| 9:1.609 | כי |  |
| 9:1.610 | קצת |  |
| 9:1.611 | האל |  |
| 9:1.612 | זב |  |
| 9:1.613 | גוף |  |
| 9:1.614 | ורוח | , |
| 9:1.615 | יתחייבו |  |
| 9:1.616 | עם |  |
| 9:1.617 | שני |  |
| 9:1.618 | התשובות |  |
| 9:1.619 | האלה |  |
| 9:1.620 | כל |  |
| 9:1.621 | מה |  |
| 9:1.622 | שהתחייבו |  |
| 9:1.623 | בו |  |
| 9:1.624 | האומרים |  |
| 9:1.625 | כי |  |
| 9:1.626 | הדברים |  |
| 9:1.627 | הנבראים |  |
| 9:1.628 | הם |  |
| 9:1.629 | מהבורא |  |
| 9:1.630 | והכת |  |
| 9:1.631 | האומרת |  |
| 9:1.632 | שגופו |  |
| 9:1.633 | ורוחו |  |
| 9:1.634 | מחודשים |  |
| 9:1.635 | ויש |  |
| 9:1.636 | בו |  |
| 9:1.637 | רוח |  |
| 9:1.638 | אחרת |  |
| 9:1.639 | מהבורא |  |
| 9:1.640 | תחייב |  |
| 9:1.641 | כי |  |
| 9:1.642 | גוף |  |
| 9:1.643 | מחודש |  |
| 9:1.644 | שב |  |
| 9:1.645 | להיות |  |
| 9:1.646 | אלוה |  |
| 9:1.647 | בשיתוף |  |
| 9:1.648 | דבר |  |
| 9:1.649 | אלהי |  |
| 9:1.650 | בו |  |
| 9:1.651 | ומקישים |  |
| 9:1.652 | אותו |  |
| 9:1.653 | להשראת |  |
| 9:1.654 | הכבוד |  |
| 9:1.655 | על |  |
| 9:1.656 | הר |  |
| 9:1.657 | סיני |  |
| 9:1.658 | ובעץ |  |
| 9:1.659 | הסנה |  |
| 9:1.660 | ובאהל |  |
| 9:1.661 | מועד | , |
| 9:1.662 | וזה |  |
| 9:1.663 | מחייב |  |
| 9:1.664 | אותם |  |
| 9:1.665 | להאמין |  |
| 9:1.666 | כי |  |
| 9:1.667 | האלה |  |
| 9:1.668 | והסנה |  |
| 9:1.669 | וההר |  |
| 9:1.670 | אלהים |  |
| 9:1.671 | ג"כ |  |
| 9:1.672 | ויוספו |  |
| 9:1.673 | רעה |  |
| 9:1.674 | וכו | ' |
| 9:1.675 | ע"כ | . |
| 9:1.676 | ומדברי |  |
| 9:1.677 | הגאון |  |
| 9:1.678 | תבין |  |
| 9:1.679 | אמתת |  |
| 9:1.680 | דברינו |  |
| 9:1.681 | שכתבתי |  |
| 9:1.682 | לעיל |  |
| 9:1.683 | כי |  |
| 9:1.684 | נכונים |  |
| 9:1.685 | הדברים |  |
| 9:1.686 | על |  |
| 9:1.687 | פי |  |
| 9:1.688 | תורתינו |  |
| 9:1.689 | הקדושה |  |
| 9:1.690 | וקבלת |  |
| 9:1.691 | רז"ל |  |
| 9:1.692 | וכמו |  |
| 9:1.693 | שכתבו |  |
| 9:1.694 | הפוסקים |  |
| 9:1.695 | דכל |  |
| 9:1.696 | דאתאמר |  |
| 9:1.697 | בדברי |  |
| 9:1.698 | גאון |  |
| 9:1.699 | אליו |  |
| 9:1.700 | שומעין |  |
| 9:1.701 | שדבריהם |  |
| 9:1.702 | דברי |  |
| 9:1.703 | קבלה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

