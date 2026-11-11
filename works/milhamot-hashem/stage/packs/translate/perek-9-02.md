# Translation pack: milhamot-hashem perek-9-02 (Milchamot ha-Shem 9:2–9:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-02`.

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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- מעמד: "assembly" or "occasion" by context; "the Standing at Sinai" where the
  author names that event.
- והוא שלא: "provided that" (fixed).
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
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A title the printed text spells wrongly (ונור אלצלאם): translate the printed
  spelling ("Nur al-Tzalam"); the likelier title (Nur al-Zulam, Me'or ha-Afelah)
  goes in the token note and `apparatus`, not in the English.
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

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


## Milchamot ha-Shem 9:2

Hebrew (Wikisource): קכא) כתב הרמב"ם ז"ל בפרק ח' מהלכות יסודי התורה, משה רבינו לא האמינו בו ישראל מפני האותות שעשה שהמאמין על פי האותות יש בלבו דופי שאפשא שיעשה האות בלאט וכשוף וכו' ובמה האמינו במעמד הר סיני. שעינינו ראו ולא זר ואזנינו שמעו ולא אחר. האש ובקולות והלפידים. והוא נגש אל הערפל והקול מדבר אליו,ואנו שומעים, משה משה לך אמור להם כך וכך, וכן הוא אומר פנים בפנים דבר ה' עמכם. ונאמר לא את אבותינו כרת ה' את הברית הזאת. ומנין שמעמד הר סיני לבדו היא הראיה לנבואתו שהיא אמת שאין בו דופי, שנאמר הנה אנכי בא אליך בעב הענן בעבור ישמע העם בדברי עמך וגם בך יאמינו לעולם, מכלל שקודם דבר זה לא האמינו בו נאמנות שהיא עומדת לעולם אלא נאמנות שיש אחריה הרהור ומחשבה, נמצאו אלו ששולח אליהם, הם עדים על נבואתו שהיא אמת ואינו צריך לעשות להן אות שהן והוא אחד בדבר, כשני עדים שראו דבר אחד ביחד, שכל אחד מהן עד לחבירו שהוא אומר אמת ואין אחד מהן צריך להביא ראיה לחבירו, כך משה רבינו ע"ה כל ישראל עדים לו מעמד הר סיני, ואינו צריך לעשות להם אות לפיכך אם עמד נביא ועשהאותות ומופתים גדולים ובקש להכחיש נבואתו של משה רבינו (ואפילו בפרט אחד או להוסיף או לגרוע בשום מצוה קלה או חמורה) אין שומעין לו ואנו יודעין ביחוד שאותן האותות בלאט וכשוף הן. לפי שנבואת משה רבינו אינה על פי האותות כדי שנערוך אותות זה לאותות זה. אלא בעינינו ראינוה ובאזנינו שמענוה כמו שממע הוא, הא למה הדבר דומה לעדים שהעידו על דבר שראה הוא בעיניו שאינו כמו שראה שאינו שומע להן. אלא יודע בודאי שהן עידי שקר. לפיכך אמרה תורה שאם בא האות והמופת לא תשמע אל דברי הנביא ההוא. שהרי זה בא להכחיש מה שראית בעיניך עכ"ל. וכן כתב הסמ"ג בהקדמת העשין וז"ל טכאשר רצה אדון הכל ליתן תורה, הטה שמים ושמי השמים על הר סיני בקולות ולפידים נפלאים ונוראים, וקרא למשה בחירו ואמר לו הנה אנכי בא אליך בעב הענן בעבור ישמע העם בדברי עמך וגם בך יאמינו לעולם, ולמה הוצרך לזה והלא כבר כתוב על הים ויאמינו בה' ובמשה עבדו? אלא כך אמר הקב"ה למשה אני רוצה שיאמינו בך ישראל במה שקדם שאתה נביא על ידי אותות ומופתים שעשית, אבל תורה שאני רוצה ליתן איני רוצה שיאמינו בך ישראל על ידי אות ומופת. אלא על ידי שישמעו באזניהם שאני מדבר עמך ולכן עשה הקב"ה את הדבר הזה, שכשיהיו ישראל בגולה. אם יאמר להם עכו"ם או ישמעאל לעזוב תורתם ולעבוד אלהים אחרים ויתן להם אות ומופת יכולין ישראל לענות לו, אפילו תתעשה הרבה אותות ומופתים כמשה בן עמרם לא נאמין לך בחילוף תורתינו, אלא אם כן נשמע באזנינו שידבר עמך השם כמו שדיבר עם משה, כי אפילו משה בחירו לא רצה שנאמינהו בתורה אלא על ידי ששמענו שדיבר עמו. ולאחר שנתן את התורה אמר אתם ראיתם כי מן השמים דברתי עמכם. כלומר אין לכם להאמין עוד לשום אדם שיאמר שהחלפתי תורה אם לא תראו שאדבר עמכם מן השמים להחליפה, ולפי שאמרנו שאין לנו תורה לשום עכו"ם או ישמעאל באות או מופת שיעשה להחליף תורתינו, פי' בתורה היאך יש יכולת לעשות אות או מופת. שנאמר כי יקום בקרבך נביא או חולם וכו' כי מנסה ה' אלהיכם אתכם, הרי פירש לך שזה האות והמות אינו מכח ע"ז שלו שלא הקב"ה עושה זה האות לנסות בו את ישראל, ואחר פטירת משה רבינו עמדו נביאים לישראל ודברו בחידות ומשלים, ועל ידי כך אמרו הגוים כי הנביאים נבאו בתורה חדשה שלהם לכך שלח הקב"ה לישראל ביד מלאכי שהוא אחרון לכל הנביאים ואמר שני דברים בסוף דבריו לחתום כל הנבואות הנה אנכי שלח לכן את אליה הנביא וכו' כלומר הנני אחרון לכל הנביאים. וממני עד שיבוא אליהו לא יעמוד נביא, והנני אומר לכם בשם הקב"ה זכרו תורת משה עבדי אשר צויתי אותו בחורב אל כל ישראל חקים ומשפטים כלומר אל יעלה על לבבכם שהנביאים שעמדו לפני נתנבאו בחילוף תורת משה, כי כולם לחזק תורת משה באו עכ"ל הסמ"ג קרוב לדברי הרמב"ם ז"ל. וכן דרשו רז"ל בפסוק שובי שובי השולמית הובא במדרש הגדול ויש פרשת וישלח ע"ש.

| id | word | punctuation after |
|---|---|---|
| 9:2.1 | קכא | ) |
| 9:2.2 | כתב |  |
| 9:2.3 | הרמב"ם |  |
| 9:2.4 | ז"ל |  |
| 9:2.5 | בפרק |  |
| 9:2.6 | ח | ' |
| 9:2.7 | מהלכות |  |
| 9:2.8 | יסודי |  |
| 9:2.9 | התורה | , |
| 9:2.10 | משה |  |
| 9:2.11 | רבינו |  |
| 9:2.12 | לא |  |
| 9:2.13 | האמינו |  |
| 9:2.14 | בו |  |
| 9:2.15 | ישראל |  |
| 9:2.16 | מפני |  |
| 9:2.17 | האותות |  |
| 9:2.18 | שעשה |  |
| 9:2.19 | שהמאמין |  |
| 9:2.20 | על |  |
| 9:2.21 | פי |  |
| 9:2.22 | האותות |  |
| 9:2.23 | יש |  |
| 9:2.24 | בלבו |  |
| 9:2.25 | דופי |  |
| 9:2.26 | שאפשא |  |
| 9:2.27 | שיעשה |  |
| 9:2.28 | האות |  |
| 9:2.29 | בלאט |  |
| 9:2.30 | וכשוף |  |
| 9:2.31 | וכו | ' |
| 9:2.32 | ובמה |  |
| 9:2.33 | האמינו |  |
| 9:2.34 | במעמד |  |
| 9:2.35 | הר |  |
| 9:2.36 | סיני | . |
| 9:2.37 | שעינינו |  |
| 9:2.38 | ראו |  |
| 9:2.39 | ולא |  |
| 9:2.40 | זר |  |
| 9:2.41 | ואזנינו |  |
| 9:2.42 | שמעו |  |
| 9:2.43 | ולא |  |
| 9:2.44 | אחר | . |
| 9:2.45 | האש |  |
| 9:2.46 | ובקולות |  |
| 9:2.47 | והלפידים | . |
| 9:2.48 | והוא |  |
| 9:2.49 | נגש |  |
| 9:2.50 | אל |  |
| 9:2.51 | הערפל |  |
| 9:2.52 | והקול |  |
| 9:2.53 | מדבר |  |
| 9:2.54 | אליו,ואנו |  |
| 9:2.55 | שומעים | , |
| 9:2.56 | משה |  |
| 9:2.57 | משה |  |
| 9:2.58 | לך |  |
| 9:2.59 | אמור |  |
| 9:2.60 | להם |  |
| 9:2.61 | כך |  |
| 9:2.62 | וכך | , |
| 9:2.63 | וכן |  |
| 9:2.64 | הוא |  |
| 9:2.65 | אומר |  |
| 9:2.66 | פנים |  |
| 9:2.67 | בפנים |  |
| 9:2.68 | דבר |  |
| 9:2.69 | ה | ' |
| 9:2.70 | עמכם | . |
| 9:2.71 | ונאמר |  |
| 9:2.72 | לא |  |
| 9:2.73 | את |  |
| 9:2.74 | אבותינו |  |
| 9:2.75 | כרת |  |
| 9:2.76 | ה | ' |
| 9:2.77 | את |  |
| 9:2.78 | הברית |  |
| 9:2.79 | הזאת | . |
| 9:2.80 | ומנין |  |
| 9:2.81 | שמעמד |  |
| 9:2.82 | הר |  |
| 9:2.83 | סיני |  |
| 9:2.84 | לבדו |  |
| 9:2.85 | היא |  |
| 9:2.86 | הראיה |  |
| 9:2.87 | לנבואתו |  |
| 9:2.88 | שהיא |  |
| 9:2.89 | אמת |  |
| 9:2.90 | שאין |  |
| 9:2.91 | בו |  |
| 9:2.92 | דופי | , |
| 9:2.93 | שנאמר |  |
| 9:2.94 | הנה |  |
| 9:2.95 | אנכי |  |
| 9:2.96 | בא |  |
| 9:2.97 | אליך |  |
| 9:2.98 | בעב |  |
| 9:2.99 | הענן |  |
| 9:2.100 | בעבור |  |
| 9:2.101 | ישמע |  |
| 9:2.102 | העם |  |
| 9:2.103 | בדברי |  |
| 9:2.104 | עמך |  |
| 9:2.105 | וגם |  |
| 9:2.106 | בך |  |
| 9:2.107 | יאמינו |  |
| 9:2.108 | לעולם | , |
| 9:2.109 | מכלל |  |
| 9:2.110 | שקודם |  |
| 9:2.111 | דבר |  |
| 9:2.112 | זה |  |
| 9:2.113 | לא |  |
| 9:2.114 | האמינו |  |
| 9:2.115 | בו |  |
| 9:2.116 | נאמנות |  |
| 9:2.117 | שהיא |  |
| 9:2.118 | עומדת |  |
| 9:2.119 | לעולם |  |
| 9:2.120 | אלא |  |
| 9:2.121 | נאמנות |  |
| 9:2.122 | שיש |  |
| 9:2.123 | אחריה |  |
| 9:2.124 | הרהור |  |
| 9:2.125 | ומחשבה | , |
| 9:2.126 | נמצאו |  |
| 9:2.127 | אלו |  |
| 9:2.128 | ששולח |  |
| 9:2.129 | אליהם | , |
| 9:2.130 | הם |  |
| 9:2.131 | עדים |  |
| 9:2.132 | על |  |
| 9:2.133 | נבואתו |  |
| 9:2.134 | שהיא |  |
| 9:2.135 | אמת |  |
| 9:2.136 | ואינו |  |
| 9:2.137 | צריך |  |
| 9:2.138 | לעשות |  |
| 9:2.139 | להן |  |
| 9:2.140 | אות |  |
| 9:2.141 | שהן |  |
| 9:2.142 | והוא |  |
| 9:2.143 | אחד |  |
| 9:2.144 | בדבר | , |
| 9:2.145 | כשני |  |
| 9:2.146 | עדים |  |
| 9:2.147 | שראו |  |
| 9:2.148 | דבר |  |
| 9:2.149 | אחד |  |
| 9:2.150 | ביחד | , |
| 9:2.151 | שכל |  |
| 9:2.152 | אחד |  |
| 9:2.153 | מהן |  |
| 9:2.154 | עד |  |
| 9:2.155 | לחבירו |  |
| 9:2.156 | שהוא |  |
| 9:2.157 | אומר |  |
| 9:2.158 | אמת |  |
| 9:2.159 | ואין |  |
| 9:2.160 | אחד |  |
| 9:2.161 | מהן |  |
| 9:2.162 | צריך |  |
| 9:2.163 | להביא |  |
| 9:2.164 | ראיה |  |
| 9:2.165 | לחבירו | , |
| 9:2.166 | כך |  |
| 9:2.167 | משה |  |
| 9:2.168 | רבינו |  |
| 9:2.169 | ע"ה |  |
| 9:2.170 | כל |  |
| 9:2.171 | ישראל |  |
| 9:2.172 | עדים |  |
| 9:2.173 | לו |  |
| 9:2.174 | מעמד |  |
| 9:2.175 | הר |  |
| 9:2.176 | סיני | , |
| 9:2.177 | ואינו |  |
| 9:2.178 | צריך |  |
| 9:2.179 | לעשות |  |
| 9:2.180 | להם |  |
| 9:2.181 | אות |  |
| 9:2.182 | לפיכך |  |
| 9:2.183 | אם |  |
| 9:2.184 | עמד |  |
| 9:2.185 | נביא |  |
| 9:2.186 | ועשהאותות |  |
| 9:2.187 | ומופתים |  |
| 9:2.188 | גדולים |  |
| 9:2.189 | ובקש |  |
| 9:2.190 | להכחיש |  |
| 9:2.191 | נבואתו |  |
| 9:2.192 | של |  |
| 9:2.193 | משה |  |
| 9:2.194 | רבינו |  |
| 9:2.195 | ואפילו |  |
| 9:2.196 | בפרט |  |
| 9:2.197 | אחד |  |
| 9:2.198 | או |  |
| 9:2.199 | להוסיף |  |
| 9:2.200 | או |  |
| 9:2.201 | לגרוע |  |
| 9:2.202 | בשום |  |
| 9:2.203 | מצוה |  |
| 9:2.204 | קלה |  |
| 9:2.205 | או |  |
| 9:2.206 | חמורה | ) |
| 9:2.207 | אין |  |
| 9:2.208 | שומעין |  |
| 9:2.209 | לו |  |
| 9:2.210 | ואנו |  |
| 9:2.211 | יודעין |  |
| 9:2.212 | ביחוד |  |
| 9:2.213 | שאותן |  |
| 9:2.214 | האותות |  |
| 9:2.215 | בלאט |  |
| 9:2.216 | וכשוף |  |
| 9:2.217 | הן | . |
| 9:2.218 | לפי |  |
| 9:2.219 | שנבואת |  |
| 9:2.220 | משה |  |
| 9:2.221 | רבינו |  |
| 9:2.222 | אינה |  |
| 9:2.223 | על |  |
| 9:2.224 | פי |  |
| 9:2.225 | האותות |  |
| 9:2.226 | כדי |  |
| 9:2.227 | שנערוך |  |
| 9:2.228 | אותות |  |
| 9:2.229 | זה |  |
| 9:2.230 | לאותות |  |
| 9:2.231 | זה | . |
| 9:2.232 | אלא |  |
| 9:2.233 | בעינינו |  |
| 9:2.234 | ראינוה |  |
| 9:2.235 | ובאזנינו |  |
| 9:2.236 | שמענוה |  |
| 9:2.237 | כמו |  |
| 9:2.238 | שממע |  |
| 9:2.239 | הוא | , |
| 9:2.240 | הא |  |
| 9:2.241 | למה |  |
| 9:2.242 | הדבר |  |
| 9:2.243 | דומה |  |
| 9:2.244 | לעדים |  |
| 9:2.245 | שהעידו |  |
| 9:2.246 | על |  |
| 9:2.247 | דבר |  |
| 9:2.248 | שראה |  |
| 9:2.249 | הוא |  |
| 9:2.250 | בעיניו |  |
| 9:2.251 | שאינו |  |
| 9:2.252 | כמו |  |
| 9:2.253 | שראה |  |
| 9:2.254 | שאינו |  |
| 9:2.255 | שומע |  |
| 9:2.256 | להן | . |
| 9:2.257 | אלא |  |
| 9:2.258 | יודע |  |
| 9:2.259 | בודאי |  |
| 9:2.260 | שהן |  |
| 9:2.261 | עידי |  |
| 9:2.262 | שקר | . |
| 9:2.263 | לפיכך |  |
| 9:2.264 | אמרה |  |
| 9:2.265 | תורה |  |
| 9:2.266 | שאם |  |
| 9:2.267 | בא |  |
| 9:2.268 | האות |  |
| 9:2.269 | והמופת |  |
| 9:2.270 | לא |  |
| 9:2.271 | תשמע |  |
| 9:2.272 | אל |  |
| 9:2.273 | דברי |  |
| 9:2.274 | הנביא |  |
| 9:2.275 | ההוא | . |
| 9:2.276 | שהרי |  |
| 9:2.277 | זה |  |
| 9:2.278 | בא |  |
| 9:2.279 | להכחיש |  |
| 9:2.280 | מה |  |
| 9:2.281 | שראית |  |
| 9:2.282 | בעיניך |  |
| 9:2.283 | עכ"ל | . |
| 9:2.284 | וכן |  |
| 9:2.285 | כתב |  |
| 9:2.286 | הסמ"ג |  |
| 9:2.287 | בהקדמת |  |
| 9:2.288 | העשין |  |
| 9:2.289 | וז"ל |  |
| 9:2.290 | טכאשר |  |
| 9:2.291 | רצה |  |
| 9:2.292 | אדון |  |
| 9:2.293 | הכל |  |
| 9:2.294 | ליתן |  |
| 9:2.295 | תורה | , |
| 9:2.296 | הטה |  |
| 9:2.297 | שמים |  |
| 9:2.298 | ושמי |  |
| 9:2.299 | השמים |  |
| 9:2.300 | על |  |
| 9:2.301 | הר |  |
| 9:2.302 | סיני |  |
| 9:2.303 | בקולות |  |
| 9:2.304 | ולפידים |  |
| 9:2.305 | נפלאים |  |
| 9:2.306 | ונוראים | , |
| 9:2.307 | וקרא |  |
| 9:2.308 | למשה |  |
| 9:2.309 | בחירו |  |
| 9:2.310 | ואמר |  |
| 9:2.311 | לו |  |
| 9:2.312 | הנה |  |
| 9:2.313 | אנכי |  |
| 9:2.314 | בא |  |
| 9:2.315 | אליך |  |
| 9:2.316 | בעב |  |
| 9:2.317 | הענן |  |
| 9:2.318 | בעבור |  |
| 9:2.319 | ישמע |  |
| 9:2.320 | העם |  |
| 9:2.321 | בדברי |  |
| 9:2.322 | עמך |  |
| 9:2.323 | וגם |  |
| 9:2.324 | בך |  |
| 9:2.325 | יאמינו |  |
| 9:2.326 | לעולם | , |
| 9:2.327 | ולמה |  |
| 9:2.328 | הוצרך |  |
| 9:2.329 | לזה |  |
| 9:2.330 | והלא |  |
| 9:2.331 | כבר |  |
| 9:2.332 | כתוב |  |
| 9:2.333 | על |  |
| 9:2.334 | הים |  |
| 9:2.335 | ויאמינו |  |
| 9:2.336 | בה | ' |
| 9:2.337 | ובמשה |  |
| 9:2.338 | עבדו | ? |
| 9:2.339 | אלא |  |
| 9:2.340 | כך |  |
| 9:2.341 | אמר |  |
| 9:2.342 | הקב"ה |  |
| 9:2.343 | למשה |  |
| 9:2.344 | אני |  |
| 9:2.345 | רוצה |  |
| 9:2.346 | שיאמינו |  |
| 9:2.347 | בך |  |
| 9:2.348 | ישראל |  |
| 9:2.349 | במה |  |
| 9:2.350 | שקדם |  |
| 9:2.351 | שאתה |  |
| 9:2.352 | נביא |  |
| 9:2.353 | על |  |
| 9:2.354 | ידי |  |
| 9:2.355 | אותות |  |
| 9:2.356 | ומופתים |  |
| 9:2.357 | שעשית | , |
| 9:2.358 | אבל |  |
| 9:2.359 | תורה |  |
| 9:2.360 | שאני |  |
| 9:2.361 | רוצה |  |
| 9:2.362 | ליתן |  |
| 9:2.363 | איני |  |
| 9:2.364 | רוצה |  |
| 9:2.365 | שיאמינו |  |
| 9:2.366 | בך |  |
| 9:2.367 | ישראל |  |
| 9:2.368 | על |  |
| 9:2.369 | ידי |  |
| 9:2.370 | אות |  |
| 9:2.371 | ומופת | . |
| 9:2.372 | אלא |  |
| 9:2.373 | על |  |
| 9:2.374 | ידי |  |
| 9:2.375 | שישמעו |  |
| 9:2.376 | באזניהם |  |
| 9:2.377 | שאני |  |
| 9:2.378 | מדבר |  |
| 9:2.379 | עמך |  |
| 9:2.380 | ולכן |  |
| 9:2.381 | עשה |  |
| 9:2.382 | הקב"ה |  |
| 9:2.383 | את |  |
| 9:2.384 | הדבר |  |
| 9:2.385 | הזה | , |
| 9:2.386 | שכשיהיו |  |
| 9:2.387 | ישראל |  |
| 9:2.388 | בגולה | . |
| 9:2.389 | אם |  |
| 9:2.390 | יאמר |  |
| 9:2.391 | להם |  |
| 9:2.392 | עכו"ם |  |
| 9:2.393 | או |  |
| 9:2.394 | ישמעאל |  |
| 9:2.395 | לעזוב |  |
| 9:2.396 | תורתם |  |
| 9:2.397 | ולעבוד |  |
| 9:2.398 | אלהים |  |
| 9:2.399 | אחרים |  |
| 9:2.400 | ויתן |  |
| 9:2.401 | להם |  |
| 9:2.402 | אות |  |
| 9:2.403 | ומופת |  |
| 9:2.404 | יכולין |  |
| 9:2.405 | ישראל |  |
| 9:2.406 | לענות |  |
| 9:2.407 | לו | , |
| 9:2.408 | אפילו |  |
| 9:2.409 | תתעשה |  |
| 9:2.410 | הרבה |  |
| 9:2.411 | אותות |  |
| 9:2.412 | ומופתים |  |
| 9:2.413 | כמשה |  |
| 9:2.414 | בן |  |
| 9:2.415 | עמרם |  |
| 9:2.416 | לא |  |
| 9:2.417 | נאמין |  |
| 9:2.418 | לך |  |
| 9:2.419 | בחילוף |  |
| 9:2.420 | תורתינו | , |
| 9:2.421 | אלא |  |
| 9:2.422 | אם |  |
| 9:2.423 | כן |  |
| 9:2.424 | נשמע |  |
| 9:2.425 | באזנינו |  |
| 9:2.426 | שידבר |  |
| 9:2.427 | עמך |  |
| 9:2.428 | השם |  |
| 9:2.429 | כמו |  |
| 9:2.430 | שדיבר |  |
| 9:2.431 | עם |  |
| 9:2.432 | משה | , |
| 9:2.433 | כי |  |
| 9:2.434 | אפילו |  |
| 9:2.435 | משה |  |
| 9:2.436 | בחירו |  |
| 9:2.437 | לא |  |
| 9:2.438 | רצה |  |
| 9:2.439 | שנאמינהו |  |
| 9:2.440 | בתורה |  |
| 9:2.441 | אלא |  |
| 9:2.442 | על |  |
| 9:2.443 | ידי |  |
| 9:2.444 | ששמענו |  |
| 9:2.445 | שדיבר |  |
| 9:2.446 | עמו | . |
| 9:2.447 | ולאחר |  |
| 9:2.448 | שנתן |  |
| 9:2.449 | את |  |
| 9:2.450 | התורה |  |
| 9:2.451 | אמר |  |
| 9:2.452 | אתם |  |
| 9:2.453 | ראיתם |  |
| 9:2.454 | כי |  |
| 9:2.455 | מן |  |
| 9:2.456 | השמים |  |
| 9:2.457 | דברתי |  |
| 9:2.458 | עמכם | . |
| 9:2.459 | כלומר |  |
| 9:2.460 | אין |  |
| 9:2.461 | לכם |  |
| 9:2.462 | להאמין |  |
| 9:2.463 | עוד |  |
| 9:2.464 | לשום |  |
| 9:2.465 | אדם |  |
| 9:2.466 | שיאמר |  |
| 9:2.467 | שהחלפתי |  |
| 9:2.468 | תורה |  |
| 9:2.469 | אם |  |
| 9:2.470 | לא |  |
| 9:2.471 | תראו |  |
| 9:2.472 | שאדבר |  |
| 9:2.473 | עמכם |  |
| 9:2.474 | מן |  |
| 9:2.475 | השמים |  |
| 9:2.476 | להחליפה | , |
| 9:2.477 | ולפי |  |
| 9:2.478 | שאמרנו |  |
| 9:2.479 | שאין |  |
| 9:2.480 | לנו |  |
| 9:2.481 | תורה |  |
| 9:2.482 | לשום |  |
| 9:2.483 | עכו"ם |  |
| 9:2.484 | או |  |
| 9:2.485 | ישמעאל |  |
| 9:2.486 | באות |  |
| 9:2.487 | או |  |
| 9:2.488 | מופת |  |
| 9:2.489 | שיעשה |  |
| 9:2.490 | להחליף |  |
| 9:2.491 | תורתינו | , |
| 9:2.492 | פי | ' |
| 9:2.493 | בתורה |  |
| 9:2.494 | היאך |  |
| 9:2.495 | יש |  |
| 9:2.496 | יכולת |  |
| 9:2.497 | לעשות |  |
| 9:2.498 | אות |  |
| 9:2.499 | או |  |
| 9:2.500 | מופת | . |
| 9:2.501 | שנאמר |  |
| 9:2.502 | כי |  |
| 9:2.503 | יקום |  |
| 9:2.504 | בקרבך |  |
| 9:2.505 | נביא |  |
| 9:2.506 | או |  |
| 9:2.507 | חולם |  |
| 9:2.508 | וכו | ' |
| 9:2.509 | כי |  |
| 9:2.510 | מנסה |  |
| 9:2.511 | ה | ' |
| 9:2.512 | אלהיכם |  |
| 9:2.513 | אתכם | , |
| 9:2.514 | הרי |  |
| 9:2.515 | פירש |  |
| 9:2.516 | לך |  |
| 9:2.517 | שזה |  |
| 9:2.518 | האות |  |
| 9:2.519 | והמות |  |
| 9:2.520 | אינו |  |
| 9:2.521 | מכח |  |
| 9:2.522 | ע"ז |  |
| 9:2.523 | שלו |  |
| 9:2.524 | שלא |  |
| 9:2.525 | הקב"ה |  |
| 9:2.526 | עושה |  |
| 9:2.527 | זה |  |
| 9:2.528 | האות |  |
| 9:2.529 | לנסות |  |
| 9:2.530 | בו |  |
| 9:2.531 | את |  |
| 9:2.532 | ישראל | , |
| 9:2.533 | ואחר |  |
| 9:2.534 | פטירת |  |
| 9:2.535 | משה |  |
| 9:2.536 | רבינו |  |
| 9:2.537 | עמדו |  |
| 9:2.538 | נביאים |  |
| 9:2.539 | לישראל |  |
| 9:2.540 | ודברו |  |
| 9:2.541 | בחידות |  |
| 9:2.542 | ומשלים | , |
| 9:2.543 | ועל |  |
| 9:2.544 | ידי |  |
| 9:2.545 | כך |  |
| 9:2.546 | אמרו |  |
| 9:2.547 | הגוים |  |
| 9:2.548 | כי |  |
| 9:2.549 | הנביאים |  |
| 9:2.550 | נבאו |  |
| 9:2.551 | בתורה |  |
| 9:2.552 | חדשה |  |
| 9:2.553 | שלהם |  |
| 9:2.554 | לכך |  |
| 9:2.555 | שלח |  |
| 9:2.556 | הקב"ה |  |
| 9:2.557 | לישראל |  |
| 9:2.558 | ביד |  |
| 9:2.559 | מלאכי |  |
| 9:2.560 | שהוא |  |
| 9:2.561 | אחרון |  |
| 9:2.562 | לכל |  |
| 9:2.563 | הנביאים |  |
| 9:2.564 | ואמר |  |
| 9:2.565 | שני |  |
| 9:2.566 | דברים |  |
| 9:2.567 | בסוף |  |
| 9:2.568 | דבריו |  |
| 9:2.569 | לחתום |  |
| 9:2.570 | כל |  |
| 9:2.571 | הנבואות |  |
| 9:2.572 | הנה |  |
| 9:2.573 | אנכי |  |
| 9:2.574 | שלח |  |
| 9:2.575 | לכן |  |
| 9:2.576 | את |  |
| 9:2.577 | אליה |  |
| 9:2.578 | הנביא |  |
| 9:2.579 | וכו | ' |
| 9:2.580 | כלומר |  |
| 9:2.581 | הנני |  |
| 9:2.582 | אחרון |  |
| 9:2.583 | לכל |  |
| 9:2.584 | הנביאים | . |
| 9:2.585 | וממני |  |
| 9:2.586 | עד |  |
| 9:2.587 | שיבוא |  |
| 9:2.588 | אליהו |  |
| 9:2.589 | לא |  |
| 9:2.590 | יעמוד |  |
| 9:2.591 | נביא | , |
| 9:2.592 | והנני |  |
| 9:2.593 | אומר |  |
| 9:2.594 | לכם |  |
| 9:2.595 | בשם |  |
| 9:2.596 | הקב"ה |  |
| 9:2.597 | זכרו |  |
| 9:2.598 | תורת |  |
| 9:2.599 | משה |  |
| 9:2.600 | עבדי |  |
| 9:2.601 | אשר |  |
| 9:2.602 | צויתי |  |
| 9:2.603 | אותו |  |
| 9:2.604 | בחורב |  |
| 9:2.605 | אל |  |
| 9:2.606 | כל |  |
| 9:2.607 | ישראל |  |
| 9:2.608 | חקים |  |
| 9:2.609 | ומשפטים |  |
| 9:2.610 | כלומר |  |
| 9:2.611 | אל |  |
| 9:2.612 | יעלה |  |
| 9:2.613 | על |  |
| 9:2.614 | לבבכם |  |
| 9:2.615 | שהנביאים |  |
| 9:2.616 | שעמדו |  |
| 9:2.617 | לפני |  |
| 9:2.618 | נתנבאו |  |
| 9:2.619 | בחילוף |  |
| 9:2.620 | תורת |  |
| 9:2.621 | משה | , |
| 9:2.622 | כי |  |
| 9:2.623 | כולם |  |
| 9:2.624 | לחזק |  |
| 9:2.625 | תורת |  |
| 9:2.626 | משה |  |
| 9:2.627 | באו |  |
| 9:2.628 | עכ"ל |  |
| 9:2.629 | הסמ"ג |  |
| 9:2.630 | קרוב |  |
| 9:2.631 | לדברי |  |
| 9:2.632 | הרמב"ם |  |
| 9:2.633 | ז"ל | . |
| 9:2.634 | וכן |  |
| 9:2.635 | דרשו |  |
| 9:2.636 | רז"ל |  |
| 9:2.637 | בפסוק |  |
| 9:2.638 | שובי |  |
| 9:2.639 | שובי |  |
| 9:2.640 | השולמית |  |
| 9:2.641 | הובא |  |
| 9:2.642 | במדרש |  |
| 9:2.643 | הגדול |  |
| 9:2.644 | ויש |  |
| 9:2.645 | פרשת |  |
| 9:2.646 | וישלח |  |
| 9:2.647 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

