# Translation pack: milhamot-hashem perek-5-15 (Milchamot ha-Shem 5:23–5:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-15`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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


## Milchamot ha-Shem 5:23

Hebrew (Wikisource): עוד שם בריש פרשת לך לך זהו שאמר הכתוב (תהלים מ"ה ח'): "אהבת צדק ותשנא רשע על כן משחך אלהים אלהיך שמן ששון מחבריך", נאמר הפסוק הזה כנגד אברהם אבינו שאהב את הקב"ה ונתקרב תחת כנפי השכינה ושנא עבודה זרה של אביו. לפי שבבית אביו היו עושין צלמים ונותנין לאברהם למכור בשוק, וכשהיה בא אדם לקנות ממנו צלם מה היה עושה היה נוטל הקורנס ומכה על ראש כל אחד מהם ואומר לזה אתה מבקש או זה, וכיון שהיו רואים שהוא מכה אותן, מיד מניחין אותן והולכין להן. והיה משוטט בדעתו ואומר עד מתי אנו משתחוין למעשה ידינו, אין ראוי לעבוד ולהשתחות אלא לארץ שהיא הוציאה פירות ומחיה אותנו. כיון שראה שהיא צריכה לגשמים ואלמלא השמים נפתחים ומורידים גשמים לא היתה מצמחת כלום, חזר ואמר אין ראוי להשתחוות אלא לשמים. חזר וראה את השמש שהיא מאירה לעולם והיא מגדלת צמחים. אמר אין ראוי להשתחות אלא לה. כיון שראה אותה ששקעה אמר אין זה אלוה. חזר וראה את הירח ואת הכוכבים שהם מאירים בלילה, אמר לאלו ראוי לעבוד ולהשתחות. כיון שעלה עמוד השחר בטלו כולם. אמר אין אלו אלוהות, ואמר אלמלא יש להם מנהיג, מפני מה זה שוקע וזה עולה. משל למה הדבר דומה לאחד שהיה מהלך בדרך וראה בירה גדולה וגבוהה ביותר. ביקש להכנס לתוכה חיזר על כל צדדיה ולא מצא פתח. קרא כמה קולות ואין עונה אותו, נשא עיניו וראה כלי צמר אדומים שטוחים על הגג. אחר כך ראה כלי פשתן לבנים. חשכה - ראה נרות דולקות. אמר אותו האיש, ודאי אדם יש בבירה, ואלו לא היה אדם בבירה למה אלו ניטלין ואלו ניתנין? כיון שראה אותו בעל הבירה שהוא מצטער על כך, הציץ עליו ואמר אני בעל הבירה. כך היה אברהם אבינו וכו' עד שהציץ הקב"ה ואמר לו אני בעל העולם וכו' בבראשית רבה וילקוט שמעוני. ובבראשית רבה פרשה ס"ד: "עקב אשר שמע אברהם בקולי ר' יוחנן ור' חנינא תרויהון אמרין בן מ"ח שנה הכיר את בוראו, ריש לקיש אמר בן שלש שנים הכיר אברהם את בוראו, מנין? עקב שמע אברהם בקול בוראו", וכתב עץ יוסף: "נראה דמודו, דתחלת ההכרה בן שלש שנים ולא נגמרה ההכרה אלא עד מ"ח שנים", עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:23.1 | עוד |  |
| 5:23.2 | שם |  |
| 5:23.3 | בריש |  |
| 5:23.4 | פרשת |  |
| 5:23.5 | לך |  |
| 5:23.6 | לך |  |
| 5:23.7 | זהו |  |
| 5:23.8 | שאמר |  |
| 5:23.9 | הכתוב |  |
| 5:23.10 | תהלים |  |
| 5:23.11 | מ"ה |  |
| 5:23.12 | ח | '): |
| 5:23.13 | אהבת |  |
| 5:23.14 | צדק |  |
| 5:23.15 | ותשנא |  |
| 5:23.16 | רשע |  |
| 5:23.17 | על |  |
| 5:23.18 | כן |  |
| 5:23.19 | משחך |  |
| 5:23.20 | אלהים |  |
| 5:23.21 | אלהיך |  |
| 5:23.22 | שמן |  |
| 5:23.23 | ששון |  |
| 5:23.24 | מחבריך | ", |
| 5:23.25 | נאמר |  |
| 5:23.26 | הפסוק |  |
| 5:23.27 | הזה |  |
| 5:23.28 | כנגד |  |
| 5:23.29 | אברהם |  |
| 5:23.30 | אבינו |  |
| 5:23.31 | שאהב |  |
| 5:23.32 | את |  |
| 5:23.33 | הקב"ה |  |
| 5:23.34 | ונתקרב |  |
| 5:23.35 | תחת |  |
| 5:23.36 | כנפי |  |
| 5:23.37 | השכינה |  |
| 5:23.38 | ושנא |  |
| 5:23.39 | עבודה |  |
| 5:23.40 | זרה |  |
| 5:23.41 | של |  |
| 5:23.42 | אביו | . |
| 5:23.43 | לפי |  |
| 5:23.44 | שבבית |  |
| 5:23.45 | אביו |  |
| 5:23.46 | היו |  |
| 5:23.47 | עושין |  |
| 5:23.48 | צלמים |  |
| 5:23.49 | ונותנין |  |
| 5:23.50 | לאברהם |  |
| 5:23.51 | למכור |  |
| 5:23.52 | בשוק | , |
| 5:23.53 | וכשהיה |  |
| 5:23.54 | בא |  |
| 5:23.55 | אדם |  |
| 5:23.56 | לקנות |  |
| 5:23.57 | ממנו |  |
| 5:23.58 | צלם |  |
| 5:23.59 | מה |  |
| 5:23.60 | היה |  |
| 5:23.61 | עושה |  |
| 5:23.62 | היה |  |
| 5:23.63 | נוטל |  |
| 5:23.64 | הקורנס |  |
| 5:23.65 | ומכה |  |
| 5:23.66 | על |  |
| 5:23.67 | ראש |  |
| 5:23.68 | כל |  |
| 5:23.69 | אחד |  |
| 5:23.70 | מהם |  |
| 5:23.71 | ואומר |  |
| 5:23.72 | לזה |  |
| 5:23.73 | אתה |  |
| 5:23.74 | מבקש |  |
| 5:23.75 | או |  |
| 5:23.76 | זה | , |
| 5:23.77 | וכיון |  |
| 5:23.78 | שהיו |  |
| 5:23.79 | רואים |  |
| 5:23.80 | שהוא |  |
| 5:23.81 | מכה |  |
| 5:23.82 | אותן | , |
| 5:23.83 | מיד |  |
| 5:23.84 | מניחין |  |
| 5:23.85 | אותן |  |
| 5:23.86 | והולכין |  |
| 5:23.87 | להן | . |
| 5:23.88 | והיה |  |
| 5:23.89 | משוטט |  |
| 5:23.90 | בדעתו |  |
| 5:23.91 | ואומר |  |
| 5:23.92 | עד |  |
| 5:23.93 | מתי |  |
| 5:23.94 | אנו |  |
| 5:23.95 | משתחוין |  |
| 5:23.96 | למעשה |  |
| 5:23.97 | ידינו | , |
| 5:23.98 | אין |  |
| 5:23.99 | ראוי |  |
| 5:23.100 | לעבוד |  |
| 5:23.101 | ולהשתחות |  |
| 5:23.102 | אלא |  |
| 5:23.103 | לארץ |  |
| 5:23.104 | שהיא |  |
| 5:23.105 | הוציאה |  |
| 5:23.106 | פירות |  |
| 5:23.107 | ומחיה |  |
| 5:23.108 | אותנו | . |
| 5:23.109 | כיון |  |
| 5:23.110 | שראה |  |
| 5:23.111 | שהיא |  |
| 5:23.112 | צריכה |  |
| 5:23.113 | לגשמים |  |
| 5:23.114 | ואלמלא |  |
| 5:23.115 | השמים |  |
| 5:23.116 | נפתחים |  |
| 5:23.117 | ומורידים |  |
| 5:23.118 | גשמים |  |
| 5:23.119 | לא |  |
| 5:23.120 | היתה |  |
| 5:23.121 | מצמחת |  |
| 5:23.122 | כלום | , |
| 5:23.123 | חזר |  |
| 5:23.124 | ואמר |  |
| 5:23.125 | אין |  |
| 5:23.126 | ראוי |  |
| 5:23.127 | להשתחוות |  |
| 5:23.128 | אלא |  |
| 5:23.129 | לשמים | . |
| 5:23.130 | חזר |  |
| 5:23.131 | וראה |  |
| 5:23.132 | את |  |
| 5:23.133 | השמש |  |
| 5:23.134 | שהיא |  |
| 5:23.135 | מאירה |  |
| 5:23.136 | לעולם |  |
| 5:23.137 | והיא |  |
| 5:23.138 | מגדלת |  |
| 5:23.139 | צמחים | . |
| 5:23.140 | אמר |  |
| 5:23.141 | אין |  |
| 5:23.142 | ראוי |  |
| 5:23.143 | להשתחות |  |
| 5:23.144 | אלא |  |
| 5:23.145 | לה | . |
| 5:23.146 | כיון |  |
| 5:23.147 | שראה |  |
| 5:23.148 | אותה |  |
| 5:23.149 | ששקעה |  |
| 5:23.150 | אמר |  |
| 5:23.151 | אין |  |
| 5:23.152 | זה |  |
| 5:23.153 | אלוה | . |
| 5:23.154 | חזר |  |
| 5:23.155 | וראה |  |
| 5:23.156 | את |  |
| 5:23.157 | הירח |  |
| 5:23.158 | ואת |  |
| 5:23.159 | הכוכבים |  |
| 5:23.160 | שהם |  |
| 5:23.161 | מאירים |  |
| 5:23.162 | בלילה | , |
| 5:23.163 | אמר |  |
| 5:23.164 | לאלו |  |
| 5:23.165 | ראוי |  |
| 5:23.166 | לעבוד |  |
| 5:23.167 | ולהשתחות | . |
| 5:23.168 | כיון |  |
| 5:23.169 | שעלה |  |
| 5:23.170 | עמוד |  |
| 5:23.171 | השחר |  |
| 5:23.172 | בטלו |  |
| 5:23.173 | כולם | . |
| 5:23.174 | אמר |  |
| 5:23.175 | אין |  |
| 5:23.176 | אלו |  |
| 5:23.177 | אלוהות | , |
| 5:23.178 | ואמר |  |
| 5:23.179 | אלמלא |  |
| 5:23.180 | יש |  |
| 5:23.181 | להם |  |
| 5:23.182 | מנהיג | , |
| 5:23.183 | מפני |  |
| 5:23.184 | מה |  |
| 5:23.185 | זה |  |
| 5:23.186 | שוקע |  |
| 5:23.187 | וזה |  |
| 5:23.188 | עולה | . |
| 5:23.189 | משל |  |
| 5:23.190 | למה |  |
| 5:23.191 | הדבר |  |
| 5:23.192 | דומה |  |
| 5:23.193 | לאחד |  |
| 5:23.194 | שהיה |  |
| 5:23.195 | מהלך |  |
| 5:23.196 | בדרך |  |
| 5:23.197 | וראה |  |
| 5:23.198 | בירה |  |
| 5:23.199 | גדולה |  |
| 5:23.200 | וגבוהה |  |
| 5:23.201 | ביותר | . |
| 5:23.202 | ביקש |  |
| 5:23.203 | להכנס |  |
| 5:23.204 | לתוכה |  |
| 5:23.205 | חיזר |  |
| 5:23.206 | על |  |
| 5:23.207 | כל |  |
| 5:23.208 | צדדיה |  |
| 5:23.209 | ולא |  |
| 5:23.210 | מצא |  |
| 5:23.211 | פתח | . |
| 5:23.212 | קרא |  |
| 5:23.213 | כמה |  |
| 5:23.214 | קולות |  |
| 5:23.215 | ואין |  |
| 5:23.216 | עונה |  |
| 5:23.217 | אותו | , |
| 5:23.218 | נשא |  |
| 5:23.219 | עיניו |  |
| 5:23.220 | וראה |  |
| 5:23.221 | כלי |  |
| 5:23.222 | צמר |  |
| 5:23.223 | אדומים |  |
| 5:23.224 | שטוחים |  |
| 5:23.225 | על |  |
| 5:23.226 | הגג | . |
| 5:23.227 | אחר |  |
| 5:23.228 | כך |  |
| 5:23.229 | ראה |  |
| 5:23.230 | כלי |  |
| 5:23.231 | פשתן |  |
| 5:23.232 | לבנים | . |
| 5:23.233 | חשכה | - |
| 5:23.234 | ראה |  |
| 5:23.235 | נרות |  |
| 5:23.236 | דולקות | . |
| 5:23.237 | אמר |  |
| 5:23.238 | אותו |  |
| 5:23.239 | האיש | , |
| 5:23.240 | ודאי |  |
| 5:23.241 | אדם |  |
| 5:23.242 | יש |  |
| 5:23.243 | בבירה | , |
| 5:23.244 | ואלו |  |
| 5:23.245 | לא |  |
| 5:23.246 | היה |  |
| 5:23.247 | אדם |  |
| 5:23.248 | בבירה |  |
| 5:23.249 | למה |  |
| 5:23.250 | אלו |  |
| 5:23.251 | ניטלין |  |
| 5:23.252 | ואלו |  |
| 5:23.253 | ניתנין | ? |
| 5:23.254 | כיון |  |
| 5:23.255 | שראה |  |
| 5:23.256 | אותו |  |
| 5:23.257 | בעל |  |
| 5:23.258 | הבירה |  |
| 5:23.259 | שהוא |  |
| 5:23.260 | מצטער |  |
| 5:23.261 | על |  |
| 5:23.262 | כך | , |
| 5:23.263 | הציץ |  |
| 5:23.264 | עליו |  |
| 5:23.265 | ואמר |  |
| 5:23.266 | אני |  |
| 5:23.267 | בעל |  |
| 5:23.268 | הבירה | . |
| 5:23.269 | כך |  |
| 5:23.270 | היה |  |
| 5:23.271 | אברהם |  |
| 5:23.272 | אבינו |  |
| 5:23.273 | וכו | ' |
| 5:23.274 | עד |  |
| 5:23.275 | שהציץ |  |
| 5:23.276 | הקב"ה |  |
| 5:23.277 | ואמר |  |
| 5:23.278 | לו |  |
| 5:23.279 | אני |  |
| 5:23.280 | בעל |  |
| 5:23.281 | העולם |  |
| 5:23.282 | וכו | ' |
| 5:23.283 | בבראשית |  |
| 5:23.284 | רבה |  |
| 5:23.285 | וילקוט |  |
| 5:23.286 | שמעוני | . |
| 5:23.287 | ובבראשית |  |
| 5:23.288 | רבה |  |
| 5:23.289 | פרשה |  |
| 5:23.290 | ס"ד | : |
| 5:23.291 | עקב |  |
| 5:23.292 | אשר |  |
| 5:23.293 | שמע |  |
| 5:23.294 | אברהם |  |
| 5:23.295 | בקולי |  |
| 5:23.296 | ר | ' |
| 5:23.297 | יוחנן |  |
| 5:23.298 | ור | ' |
| 5:23.299 | חנינא |  |
| 5:23.300 | תרויהון |  |
| 5:23.301 | אמרין |  |
| 5:23.302 | בן |  |
| 5:23.303 | מ"ח |  |
| 5:23.304 | שנה |  |
| 5:23.305 | הכיר |  |
| 5:23.306 | את |  |
| 5:23.307 | בוראו | , |
| 5:23.308 | ריש |  |
| 5:23.309 | לקיש |  |
| 5:23.310 | אמר |  |
| 5:23.311 | בן |  |
| 5:23.312 | שלש |  |
| 5:23.313 | שנים |  |
| 5:23.314 | הכיר |  |
| 5:23.315 | אברהם |  |
| 5:23.316 | את |  |
| 5:23.317 | בוראו | , |
| 5:23.318 | מנין | ? |
| 5:23.319 | עקב |  |
| 5:23.320 | שמע |  |
| 5:23.321 | אברהם |  |
| 5:23.322 | בקול |  |
| 5:23.323 | בוראו | ", |
| 5:23.324 | וכתב |  |
| 5:23.325 | עץ |  |
| 5:23.326 | יוסף | : |
| 5:23.327 | נראה |  |
| 5:23.328 | דמודו | , |
| 5:23.329 | דתחלת |  |
| 5:23.330 | ההכרה |  |
| 5:23.331 | בן |  |
| 5:23.332 | שלש |  |
| 5:23.333 | שנים |  |
| 5:23.334 | ולא |  |
| 5:23.335 | נגמרה |  |
| 5:23.336 | ההכרה |  |
| 5:23.337 | אלא |  |
| 5:23.338 | עד |  |
| 5:23.339 | מ"ח |  |
| 5:23.340 | שנים | ", |
| 5:23.341 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

