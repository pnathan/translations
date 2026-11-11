# Translation pack: milhamot-hashem perek-4-11 (Milchamot ha-Shem 4:19–4:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-11`.

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
- **יחוד / ייחוד** (both spellings, one rendering):
  - The doctrine: fixed, "unity". "ייחוד השי"ת" → "the unity of the Name,
    may He be blessed"; "אמונת היחוד" → "the belief in the unity"; "שער
    היחוד" → "the Gate of Unity". Not "monotheism" and not "oneness".
    אחדות → "oneness", kept distinct.
  - The Kabbalistic act: "unification". "לשם יחוד קב"ה ושכינתיה" (6:21) →
    "for the sake of the unification of the Holy One, blessed be He, and His
    Shekhinah"; יחודים → "unifications".
  - The verb: ליחד with God as object, in the rabbinic sense → "to declare
    one": "הזהירנו ליחדו" (5:26) → "commanded us to declare Him one".
    "מיחדים את השי"ת" → "declare the Name, may He be blessed, one". With the
    Kabbalistic entities as object → "to unify": "ליחד זוגתו עמו" (7:5) →
    "to unify his consort with him".
  - The author plays on the two senses ("עיקר היחוד אצלם לחבר ולקשר", 6:1;
    "סוד היחוד הוא ליחד את השם בכל עשר ספירות", 7:25). There, keep "unity" for
    the noun and "unify" for the verb, so that the shared stem shows, and
    explain the play in a note.
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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
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
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.
- הרא"ש: "the Rosh"; ר"ת: "Rabbeinu Tam"; הגאון alone: "the Gaon", unidentified.
  (Acronym authorities keep the acronym, by the ruling above.)

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


## Milchamot ha-Shem 4:19

Hebrew (Wikisource): ל) הנה מכל זה נתבאר לך באר היטב כי הזהר וכל מפרשיו מקדש מלך וכסא אליהו ור' עמנואל בעל משנת חסידים, ויושר לבב, והרש"ב, ומצרף האמונה וספר הכוונות, ומהרח"ו בעץ חיים, ומחברת הקדש, וספר הברית, והאר"י בספר הלקוטים ונחלת יוסף, וזולתם, כולם כאחד עונים ואומרים, דבאין סוף שהוא הסבה הראשונה לא שייך שום עבודה וקריאה. רק לזעיר אנפין שהוא סבה אחרונה שבכל הנאצלים שהוא לפי דעתם המקשר את כל הפרצופים, והוא הזן ומפרנס לכל הברואים ולו לבדו נעבוד ואליו נקרא בעת צרה, והוא יעננו ביום קראינו. אך הר' כסא אליהו מצריך להזכיר עמו בהתפללו אליו גם שאר הפרצופים שלמעלה ממנו, שאם אינו משתפם עמו אינו נענה מהרה אלא א"כ מזכיר שאר פרצופים עמו. וז"ל בדף כ"ו אלא מיבעי ליה ליחדם על הדרך שאמרנו להתפלל לזעיר אנפין. ובתפלתם להזכיר שמות הפרצופים הגבוהים ממנו כי צריכים להם, ואם לא יעשו כן, ולא יחדו תפלתם אלא לזעיר אנפין לא יהיו נענים מהרה וכדאמרן וזהו צירוף מ"ה תצע"ק אל"י, שהם אותיות מאצל העתיק וכו', ומסיים שם דגם אם לא היו מתפללים אלא לזעיר אנפין לא היו פוגמין אלא שלא היו נענים כל כך בקל ע"כ. וכן כתב הרב יושר לבב בית ב' חדר ג' פ' ז' ע"ש. הנה מה מאד ישתומם הקורא הנלבב הדבק בתורת משה רבינו ע"ה שבכתב ושבע"פ ובקי בה דהיינו במשנה ובתלמוד ומדרשי רז"ל, ועיניו יחזו דברי הגאונים, החסיד בעל חובת הלבבות, ורבי יהודה הלוי, ורבי סעדיה גאון, והרמב"ם בחיבוריו, ור' אליעזר מגרמיזא ברוקח וספר מצות גדול, וספר מצות קטן ובעל העקרים וכיוצא המדברים בענין יחוד השי"ת על אמתתו כפי קבלת רז"ל מעתיקי השמועה.

| id | word | punctuation after |
|---|---|---|
| 4:19.1 | ל | ) |
| 4:19.2 | הנה |  |
| 4:19.3 | מכל |  |
| 4:19.4 | זה |  |
| 4:19.5 | נתבאר |  |
| 4:19.6 | לך |  |
| 4:19.7 | באר |  |
| 4:19.8 | היטב |  |
| 4:19.9 | כי |  |
| 4:19.10 | הזהר |  |
| 4:19.11 | וכל |  |
| 4:19.12 | מפרשיו |  |
| 4:19.13 | מקדש |  |
| 4:19.14 | מלך |  |
| 4:19.15 | וכסא |  |
| 4:19.16 | אליהו |  |
| 4:19.17 | ור | ' |
| 4:19.18 | עמנואל |  |
| 4:19.19 | בעל |  |
| 4:19.20 | משנת |  |
| 4:19.21 | חסידים | , |
| 4:19.22 | ויושר |  |
| 4:19.23 | לבב | , |
| 4:19.24 | והרש"ב | , |
| 4:19.25 | ומצרף |  |
| 4:19.26 | האמונה |  |
| 4:19.27 | וספר |  |
| 4:19.28 | הכוונות | , |
| 4:19.29 | ומהרח"ו |  |
| 4:19.30 | בעץ |  |
| 4:19.31 | חיים | , |
| 4:19.32 | ומחברת |  |
| 4:19.33 | הקדש | , |
| 4:19.34 | וספר |  |
| 4:19.35 | הברית | , |
| 4:19.36 | והאר"י |  |
| 4:19.37 | בספר |  |
| 4:19.38 | הלקוטים |  |
| 4:19.39 | ונחלת |  |
| 4:19.40 | יוסף | , |
| 4:19.41 | וזולתם | , |
| 4:19.42 | כולם |  |
| 4:19.43 | כאחד |  |
| 4:19.44 | עונים |  |
| 4:19.45 | ואומרים | , |
| 4:19.46 | דבאין |  |
| 4:19.47 | סוף |  |
| 4:19.48 | שהוא |  |
| 4:19.49 | הסבה |  |
| 4:19.50 | הראשונה |  |
| 4:19.51 | לא |  |
| 4:19.52 | שייך |  |
| 4:19.53 | שום |  |
| 4:19.54 | עבודה |  |
| 4:19.55 | וקריאה | . |
| 4:19.56 | רק |  |
| 4:19.57 | לזעיר |  |
| 4:19.58 | אנפין |  |
| 4:19.59 | שהוא |  |
| 4:19.60 | סבה |  |
| 4:19.61 | אחרונה |  |
| 4:19.62 | שבכל |  |
| 4:19.63 | הנאצלים |  |
| 4:19.64 | שהוא |  |
| 4:19.65 | לפי |  |
| 4:19.66 | דעתם |  |
| 4:19.67 | המקשר |  |
| 4:19.68 | את |  |
| 4:19.69 | כל |  |
| 4:19.70 | הפרצופים | , |
| 4:19.71 | והוא |  |
| 4:19.72 | הזן |  |
| 4:19.73 | ומפרנס |  |
| 4:19.74 | לכל |  |
| 4:19.75 | הברואים |  |
| 4:19.76 | ולו |  |
| 4:19.77 | לבדו |  |
| 4:19.78 | נעבוד |  |
| 4:19.79 | ואליו |  |
| 4:19.80 | נקרא |  |
| 4:19.81 | בעת |  |
| 4:19.82 | צרה | , |
| 4:19.83 | והוא |  |
| 4:19.84 | יעננו |  |
| 4:19.85 | ביום |  |
| 4:19.86 | קראינו | . |
| 4:19.87 | אך |  |
| 4:19.88 | הר | ' |
| 4:19.89 | כסא |  |
| 4:19.90 | אליהו |  |
| 4:19.91 | מצריך |  |
| 4:19.92 | להזכיר |  |
| 4:19.93 | עמו |  |
| 4:19.94 | בהתפללו |  |
| 4:19.95 | אליו |  |
| 4:19.96 | גם |  |
| 4:19.97 | שאר |  |
| 4:19.98 | הפרצופים |  |
| 4:19.99 | שלמעלה |  |
| 4:19.100 | ממנו | , |
| 4:19.101 | שאם |  |
| 4:19.102 | אינו |  |
| 4:19.103 | משתפם |  |
| 4:19.104 | עמו |  |
| 4:19.105 | אינו |  |
| 4:19.106 | נענה |  |
| 4:19.107 | מהרה |  |
| 4:19.108 | אלא |  |
| 4:19.109 | א"כ |  |
| 4:19.110 | מזכיר |  |
| 4:19.111 | שאר |  |
| 4:19.112 | פרצופים |  |
| 4:19.113 | עמו | . |
| 4:19.114 | וז"ל |  |
| 4:19.115 | בדף |  |
| 4:19.116 | כ"ו |  |
| 4:19.117 | אלא |  |
| 4:19.118 | מיבעי |  |
| 4:19.119 | ליה |  |
| 4:19.120 | ליחדם |  |
| 4:19.121 | על |  |
| 4:19.122 | הדרך |  |
| 4:19.123 | שאמרנו |  |
| 4:19.124 | להתפלל |  |
| 4:19.125 | לזעיר |  |
| 4:19.126 | אנפין | . |
| 4:19.127 | ובתפלתם |  |
| 4:19.128 | להזכיר |  |
| 4:19.129 | שמות |  |
| 4:19.130 | הפרצופים |  |
| 4:19.131 | הגבוהים |  |
| 4:19.132 | ממנו |  |
| 4:19.133 | כי |  |
| 4:19.134 | צריכים |  |
| 4:19.135 | להם | , |
| 4:19.136 | ואם |  |
| 4:19.137 | לא |  |
| 4:19.138 | יעשו |  |
| 4:19.139 | כן | , |
| 4:19.140 | ולא |  |
| 4:19.141 | יחדו |  |
| 4:19.142 | תפלתם |  |
| 4:19.143 | אלא |  |
| 4:19.144 | לזעיר |  |
| 4:19.145 | אנפין |  |
| 4:19.146 | לא |  |
| 4:19.147 | יהיו |  |
| 4:19.148 | נענים |  |
| 4:19.149 | מהרה |  |
| 4:19.150 | וכדאמרן |  |
| 4:19.151 | וזהו |  |
| 4:19.152 | צירוף |  |
| 4:19.153 | מ"ה |  |
| 4:19.154 | תצע"ק |  |
| 4:19.155 | אל"י | , |
| 4:19.156 | שהם |  |
| 4:19.157 | אותיות |  |
| 4:19.158 | מאצל |  |
| 4:19.159 | העתיק |  |
| 4:19.160 | וכו | ', |
| 4:19.161 | ומסיים |  |
| 4:19.162 | שם |  |
| 4:19.163 | דגם |  |
| 4:19.164 | אם |  |
| 4:19.165 | לא |  |
| 4:19.166 | היו |  |
| 4:19.167 | מתפללים |  |
| 4:19.168 | אלא |  |
| 4:19.169 | לזעיר |  |
| 4:19.170 | אנפין |  |
| 4:19.171 | לא |  |
| 4:19.172 | היו |  |
| 4:19.173 | פוגמין |  |
| 4:19.174 | אלא |  |
| 4:19.175 | שלא |  |
| 4:19.176 | היו |  |
| 4:19.177 | נענים |  |
| 4:19.178 | כל |  |
| 4:19.179 | כך |  |
| 4:19.180 | בקל |  |
| 4:19.181 | ע"כ | . |
| 4:19.182 | וכן |  |
| 4:19.183 | כתב |  |
| 4:19.184 | הרב |  |
| 4:19.185 | יושר |  |
| 4:19.186 | לבב |  |
| 4:19.187 | בית |  |
| 4:19.188 | ב | ' |
| 4:19.189 | חדר |  |
| 4:19.190 | ג | ' |
| 4:19.191 | פ | ' |
| 4:19.192 | ז | ' |
| 4:19.193 | ע"ש | . |
| 4:19.194 | הנה |  |
| 4:19.195 | מה |  |
| 4:19.196 | מאד |  |
| 4:19.197 | ישתומם |  |
| 4:19.198 | הקורא |  |
| 4:19.199 | הנלבב |  |
| 4:19.200 | הדבק |  |
| 4:19.201 | בתורת |  |
| 4:19.202 | משה |  |
| 4:19.203 | רבינו |  |
| 4:19.204 | ע"ה |  |
| 4:19.205 | שבכתב |  |
| 4:19.206 | ושבע"פ |  |
| 4:19.207 | ובקי |  |
| 4:19.208 | בה |  |
| 4:19.209 | דהיינו |  |
| 4:19.210 | במשנה |  |
| 4:19.211 | ובתלמוד |  |
| 4:19.212 | ומדרשי |  |
| 4:19.213 | רז"ל | , |
| 4:19.214 | ועיניו |  |
| 4:19.215 | יחזו |  |
| 4:19.216 | דברי |  |
| 4:19.217 | הגאונים | , |
| 4:19.218 | החסיד |  |
| 4:19.219 | בעל |  |
| 4:19.220 | חובת |  |
| 4:19.221 | הלבבות | , |
| 4:19.222 | ורבי |  |
| 4:19.223 | יהודה |  |
| 4:19.224 | הלוי | , |
| 4:19.225 | ורבי |  |
| 4:19.226 | סעדיה |  |
| 4:19.227 | גאון | , |
| 4:19.228 | והרמב"ם |  |
| 4:19.229 | בחיבוריו | , |
| 4:19.230 | ור | ' |
| 4:19.231 | אליעזר |  |
| 4:19.232 | מגרמיזא |  |
| 4:19.233 | ברוקח |  |
| 4:19.234 | וספר |  |
| 4:19.235 | מצות |  |
| 4:19.236 | גדול | , |
| 4:19.237 | וספר |  |
| 4:19.238 | מצות |  |
| 4:19.239 | קטן |  |
| 4:19.240 | ובעל |  |
| 4:19.241 | העקרים |  |
| 4:19.242 | וכיוצא |  |
| 4:19.243 | המדברים |  |
| 4:19.244 | בענין |  |
| 4:19.245 | יחוד |  |
| 4:19.246 | השי"ת |  |
| 4:19.247 | על |  |
| 4:19.248 | אמתתו |  |
| 4:19.249 | כפי |  |
| 4:19.250 | קבלת |  |
| 4:19.251 | רז"ל |  |
| 4:19.252 | מעתיקי |  |
| 4:19.253 | השמועה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

