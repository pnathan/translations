# Translation pack: milhamot-hashem perek-8-16 (Milchamot ha-Shem 8:19–8:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-16`.

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

- **מדה / מדות / מידות**: context. Do not render it "sefirot" even where that
  is the referent.
  - The hermeneutical rules: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי"
    (1:6) → "the thirteen rules of Rabbi Ishmael and the thirty-two rules of
    Rabbi Yose the Galilean"; "שלש עשרה מדות שהתורה נדרשת בהם" (10:5) → "the
    thirteen rules by which the Torah is expounded". Add a note: literally
    "measures".
  - Exod 34:6–7: "שלש עשרה מדות של רחמים" (9:8) → "the thirteen attributes of
    mercy". These are a different thirteen from the rules. The author uses
    both, so never let one rendering cover the other.
  - Kabbalistic and philosophical attributes, God's or the sefirot's: "מדת
    הדין" → "the attribute of judgment"; "המדות העליונות" (2:11) → "the
    upper attributes"; "המידות הפנימיות" (2:7) → "the inner attributes";
    "מדת המלכות" → "the attribute of Malkhut"; "מלובש במידותיו" (5:15) →
    "clothed in His attributes"; "ממדותיו" (7:17, Exod 34:6) → "of His
    attributes".
  - Measure: "עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה" (4:32) → "ten
    measures and ten sefirot, because they have measure and limit"; "לא מקום
    ולא מידה" (7:13) → "neither place nor measure"; "במדה שמדדתם" →
    "with the measure with which you measured".
  - A quality or trait: "ממדת הענוה" (4:8) → "of the quality of humility";
    "המדה המגונה הזאת" (7:19) → "this disgraceful trait"; "אין לך מדה גדולה
    מזו" (6:14, Talmud) → "there is no greater virtue than this".
  - A proper name: מסכת מדות → "tractate Middot".
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
- **קצר אפים**: the author's Hebrew calque of זעיר אנפין, set against ארך
  אפים in Exod 34:6. Write "quick to anger", and render ארך אפים "slow to
  anger" everywhere, so that the pair shows in English. "באל זר קצר אפים
  (ז״א)" (1:4) → "a strange god, quick to anger (Ze'ir Anpin)"; "הנקרא (קצר
  אפים) זעיר אנפין" (3:3) → "called (quick to anger) Ze'ir Anpin". Add a note
  each time: Aramaic "small face"; Prov 14:17. "קצא אפים" (6:23) is translated
  as קצר, with the spelling recorded in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
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


## Milchamot ha-Shem 8:19

Hebrew (Wikisource): ונחזור לעניינינו ובהדי דתפסוח אשתמיט ביניתא מן מזייה, היינו שתפסו לאותם האנשים הטועים ומטעים אחרים לדרוש בתורה דרשות של דופי כאלה לעבוד אלוה אחר קצר אפים מבלעדי ה' אלהינו, ולירא מאשתו (זהר וירא דף קי"ב). ואחזו בזקנם שיחזרו בהם מלפרש ולדרוש דרשות זרות כאלה בתורה הקדושה, דרשות ופירושים המביאים לידי מינות ואמונת אלהים רבים. ונשמט שער מזקנם באחזם בה, כי הדרת פנים זקן ונפחת כבודם על ידי זה, ונגרע הודם וחדרם. ולכן רמו קלא ואזל קליהו ארבע מאה פרסי, שצעקו חמס על הודם והדרם וכבודם כי אזל ועבר והלך קול זלזולם ארבע מאה פרסי, באורך ורוחב ארץ ישראל שהיא ארבע מאה פרסי, ואמרי היכי נעביד ליה, דלמא מרחמי עליה מן שמיא, שחששו שמא בזכות התורה שעסקו בה אע"פ שמרמזים בה דברי מינות ירחמו עליהם מן השמים, כמו שמצינו באלישע אחר שעמדה לו זכות תורה בקצת, ואמרו בשמים לא מידן נידייניה דעסק בתורה וכו'. אמר להו נביא שדויה בדורא דאברא וכו' שהבטיחם הנביא זכריה ואמר להם אל תחושו לקול צעקתם וצוה אותם לשקע ולשכח מלב בני ישראל כל אותן הדרשות והרמזים של מינות כדודא דאברא ששוקע בים, כמ"ד צללו כעופרת במים אדירים של תורת אמת המקובלת ממשה רבינו ע"ה להקים בה בתי מדרשות בכל עיר ובכל פלך, ולהעמיד תלמידים הרבה כדאמרינן בפ' קמא דאבות, כי אנשי כנה"ג אמרו שלשה דברים היו מתונים בדין והעמידו תלמידים הרבה, ועשו סייג לתורה שלא לדרוש בה כל איש ככל העולה על רוחו, אלא על פי הכללים והיסודות האמתיים שנמסרו למשה רבינו ע"ה בסיני! וזוהי מליצת חפיוה לפומא באברא שמנעו וסתמו את פי החכמים בעיניהם מלדרוש בתורה הזיות והבלים המביאים לידי מינות, והשבעות רוחות של מתים לבוא להגיד להם נסתרות באברא, אשר אם תקים עליו לא ישמע קול, כן נעשו דורשי דרשות של דופי והזיות כחרשים! ובזה נתבטל יצרא דע"ז, ומן השמים אסתייעא מלתא, ועלתה בידם לשכח שמות האלילים אשר ידמו התרים אחרי לבבם לחשוב רבוי באלוהות, ושיש כמה עלות וסבות המשתלשלים זה מזה, ונוטלים רשות זה מזה לברוא את אשר יחפוץ לברוא אותו כאימא שנטלה רשות מאבא לברוא אדם הראשון, וכל אחד מאלו הצורות נקרא בשם ידוע, ונקרא בשם הוי'ה ואדנות ושאר השמות המיוחדים לשי"ת, ושהעבודה והתפלה להם לבדם אעפ"י שהם נבראים, אבל האלוה העליון איננו ראוי לשום שבח ותהלה לרוב מעלתו והעלמו, ואין לו דמות וצורה ואין לו אשה ולא שם מוגבל! כל מחשבות מינות הללו השכיחום רבותינו אנשי כנסת הגדולה ז"ל.

| id | word | punctuation after |
|---|---|---|
| 8:19.1 | ונחזור |  |
| 8:19.2 | לעניינינו |  |
| 8:19.3 | ובהדי |  |
| 8:19.4 | דתפסוח |  |
| 8:19.5 | אשתמיט |  |
| 8:19.6 | ביניתא |  |
| 8:19.7 | מן |  |
| 8:19.8 | מזייה | , |
| 8:19.9 | היינו |  |
| 8:19.10 | שתפסו |  |
| 8:19.11 | לאותם |  |
| 8:19.12 | האנשים |  |
| 8:19.13 | הטועים |  |
| 8:19.14 | ומטעים |  |
| 8:19.15 | אחרים |  |
| 8:19.16 | לדרוש |  |
| 8:19.17 | בתורה |  |
| 8:19.18 | דרשות |  |
| 8:19.19 | של |  |
| 8:19.20 | דופי |  |
| 8:19.21 | כאלה |  |
| 8:19.22 | לעבוד |  |
| 8:19.23 | אלוה |  |
| 8:19.24 | אחר |  |
| 8:19.25 | קצר |  |
| 8:19.26 | אפים |  |
| 8:19.27 | מבלעדי |  |
| 8:19.28 | ה | ' |
| 8:19.29 | אלהינו | , |
| 8:19.30 | ולירא |  |
| 8:19.31 | מאשתו |  |
| 8:19.32 | זהר |  |
| 8:19.33 | וירא |  |
| 8:19.34 | דף |  |
| 8:19.35 | קי"ב | ). |
| 8:19.36 | ואחזו |  |
| 8:19.37 | בזקנם |  |
| 8:19.38 | שיחזרו |  |
| 8:19.39 | בהם |  |
| 8:19.40 | מלפרש |  |
| 8:19.41 | ולדרוש |  |
| 8:19.42 | דרשות |  |
| 8:19.43 | זרות |  |
| 8:19.44 | כאלה |  |
| 8:19.45 | בתורה |  |
| 8:19.46 | הקדושה | , |
| 8:19.47 | דרשות |  |
| 8:19.48 | ופירושים |  |
| 8:19.49 | המביאים |  |
| 8:19.50 | לידי |  |
| 8:19.51 | מינות |  |
| 8:19.52 | ואמונת |  |
| 8:19.53 | אלהים |  |
| 8:19.54 | רבים | . |
| 8:19.55 | ונשמט |  |
| 8:19.56 | שער |  |
| 8:19.57 | מזקנם |  |
| 8:19.58 | באחזם |  |
| 8:19.59 | בה | , |
| 8:19.60 | כי |  |
| 8:19.61 | הדרת |  |
| 8:19.62 | פנים |  |
| 8:19.63 | זקן |  |
| 8:19.64 | ונפחת |  |
| 8:19.65 | כבודם |  |
| 8:19.66 | על |  |
| 8:19.67 | ידי |  |
| 8:19.68 | זה | , |
| 8:19.69 | ונגרע |  |
| 8:19.70 | הודם |  |
| 8:19.71 | וחדרם | . |
| 8:19.72 | ולכן |  |
| 8:19.73 | רמו |  |
| 8:19.74 | קלא |  |
| 8:19.75 | ואזל |  |
| 8:19.76 | קליהו |  |
| 8:19.77 | ארבע |  |
| 8:19.78 | מאה |  |
| 8:19.79 | פרסי | , |
| 8:19.80 | שצעקו |  |
| 8:19.81 | חמס |  |
| 8:19.82 | על |  |
| 8:19.83 | הודם |  |
| 8:19.84 | והדרם |  |
| 8:19.85 | וכבודם |  |
| 8:19.86 | כי |  |
| 8:19.87 | אזל |  |
| 8:19.88 | ועבר |  |
| 8:19.89 | והלך |  |
| 8:19.90 | קול |  |
| 8:19.91 | זלזולם |  |
| 8:19.92 | ארבע |  |
| 8:19.93 | מאה |  |
| 8:19.94 | פרסי | , |
| 8:19.95 | באורך |  |
| 8:19.96 | ורוחב |  |
| 8:19.97 | ארץ |  |
| 8:19.98 | ישראל |  |
| 8:19.99 | שהיא |  |
| 8:19.100 | ארבע |  |
| 8:19.101 | מאה |  |
| 8:19.102 | פרסי | , |
| 8:19.103 | ואמרי |  |
| 8:19.104 | היכי |  |
| 8:19.105 | נעביד |  |
| 8:19.106 | ליה | , |
| 8:19.107 | דלמא |  |
| 8:19.108 | מרחמי |  |
| 8:19.109 | עליה |  |
| 8:19.110 | מן |  |
| 8:19.111 | שמיא | , |
| 8:19.112 | שחששו |  |
| 8:19.113 | שמא |  |
| 8:19.114 | בזכות |  |
| 8:19.115 | התורה |  |
| 8:19.116 | שעסקו |  |
| 8:19.117 | בה |  |
| 8:19.118 | אע"פ |  |
| 8:19.119 | שמרמזים |  |
| 8:19.120 | בה |  |
| 8:19.121 | דברי |  |
| 8:19.122 | מינות |  |
| 8:19.123 | ירחמו |  |
| 8:19.124 | עליהם |  |
| 8:19.125 | מן |  |
| 8:19.126 | השמים | , |
| 8:19.127 | כמו |  |
| 8:19.128 | שמצינו |  |
| 8:19.129 | באלישע |  |
| 8:19.130 | אחר |  |
| 8:19.131 | שעמדה |  |
| 8:19.132 | לו |  |
| 8:19.133 | זכות |  |
| 8:19.134 | תורה |  |
| 8:19.135 | בקצת | , |
| 8:19.136 | ואמרו |  |
| 8:19.137 | בשמים |  |
| 8:19.138 | לא |  |
| 8:19.139 | מידן |  |
| 8:19.140 | נידייניה |  |
| 8:19.141 | דעסק |  |
| 8:19.142 | בתורה |  |
| 8:19.143 | וכו | '. |
| 8:19.144 | אמר |  |
| 8:19.145 | להו |  |
| 8:19.146 | נביא |  |
| 8:19.147 | שדויה |  |
| 8:19.148 | בדורא |  |
| 8:19.149 | דאברא |  |
| 8:19.150 | וכו | ' |
| 8:19.151 | שהבטיחם |  |
| 8:19.152 | הנביא |  |
| 8:19.153 | זכריה |  |
| 8:19.154 | ואמר |  |
| 8:19.155 | להם |  |
| 8:19.156 | אל |  |
| 8:19.157 | תחושו |  |
| 8:19.158 | לקול |  |
| 8:19.159 | צעקתם |  |
| 8:19.160 | וצוה |  |
| 8:19.161 | אותם |  |
| 8:19.162 | לשקע |  |
| 8:19.163 | ולשכח |  |
| 8:19.164 | מלב |  |
| 8:19.165 | בני |  |
| 8:19.166 | ישראל |  |
| 8:19.167 | כל |  |
| 8:19.168 | אותן |  |
| 8:19.169 | הדרשות |  |
| 8:19.170 | והרמזים |  |
| 8:19.171 | של |  |
| 8:19.172 | מינות |  |
| 8:19.173 | כדודא |  |
| 8:19.174 | דאברא |  |
| 8:19.175 | ששוקע |  |
| 8:19.176 | בים | , |
| 8:19.177 | כמ"ד |  |
| 8:19.178 | צללו |  |
| 8:19.179 | כעופרת |  |
| 8:19.180 | במים |  |
| 8:19.181 | אדירים |  |
| 8:19.182 | של |  |
| 8:19.183 | תורת |  |
| 8:19.184 | אמת |  |
| 8:19.185 | המקובלת |  |
| 8:19.186 | ממשה |  |
| 8:19.187 | רבינו |  |
| 8:19.188 | ע"ה |  |
| 8:19.189 | להקים |  |
| 8:19.190 | בה |  |
| 8:19.191 | בתי |  |
| 8:19.192 | מדרשות |  |
| 8:19.193 | בכל |  |
| 8:19.194 | עיר |  |
| 8:19.195 | ובכל |  |
| 8:19.196 | פלך | , |
| 8:19.197 | ולהעמיד |  |
| 8:19.198 | תלמידים |  |
| 8:19.199 | הרבה |  |
| 8:19.200 | כדאמרינן |  |
| 8:19.201 | בפ | ' |
| 8:19.202 | קמא |  |
| 8:19.203 | דאבות | , |
| 8:19.204 | כי |  |
| 8:19.205 | אנשי |  |
| 8:19.206 | כנה"ג |  |
| 8:19.207 | אמרו |  |
| 8:19.208 | שלשה |  |
| 8:19.209 | דברים |  |
| 8:19.210 | היו |  |
| 8:19.211 | מתונים |  |
| 8:19.212 | בדין |  |
| 8:19.213 | והעמידו |  |
| 8:19.214 | תלמידים |  |
| 8:19.215 | הרבה | , |
| 8:19.216 | ועשו |  |
| 8:19.217 | סייג |  |
| 8:19.218 | לתורה |  |
| 8:19.219 | שלא |  |
| 8:19.220 | לדרוש |  |
| 8:19.221 | בה |  |
| 8:19.222 | כל |  |
| 8:19.223 | איש |  |
| 8:19.224 | ככל |  |
| 8:19.225 | העולה |  |
| 8:19.226 | על |  |
| 8:19.227 | רוחו | , |
| 8:19.228 | אלא |  |
| 8:19.229 | על |  |
| 8:19.230 | פי |  |
| 8:19.231 | הכללים |  |
| 8:19.232 | והיסודות |  |
| 8:19.233 | האמתיים |  |
| 8:19.234 | שנמסרו |  |
| 8:19.235 | למשה |  |
| 8:19.236 | רבינו |  |
| 8:19.237 | ע"ה |  |
| 8:19.238 | בסיני | ! |
| 8:19.239 | וזוהי |  |
| 8:19.240 | מליצת |  |
| 8:19.241 | חפיוה |  |
| 8:19.242 | לפומא |  |
| 8:19.243 | באברא |  |
| 8:19.244 | שמנעו |  |
| 8:19.245 | וסתמו |  |
| 8:19.246 | את |  |
| 8:19.247 | פי |  |
| 8:19.248 | החכמים |  |
| 8:19.249 | בעיניהם |  |
| 8:19.250 | מלדרוש |  |
| 8:19.251 | בתורה |  |
| 8:19.252 | הזיות |  |
| 8:19.253 | והבלים |  |
| 8:19.254 | המביאים |  |
| 8:19.255 | לידי |  |
| 8:19.256 | מינות | , |
| 8:19.257 | והשבעות |  |
| 8:19.258 | רוחות |  |
| 8:19.259 | של |  |
| 8:19.260 | מתים |  |
| 8:19.261 | לבוא |  |
| 8:19.262 | להגיד |  |
| 8:19.263 | להם |  |
| 8:19.264 | נסתרות |  |
| 8:19.265 | באברא | , |
| 8:19.266 | אשר |  |
| 8:19.267 | אם |  |
| 8:19.268 | תקים |  |
| 8:19.269 | עליו |  |
| 8:19.270 | לא |  |
| 8:19.271 | ישמע |  |
| 8:19.272 | קול | , |
| 8:19.273 | כן |  |
| 8:19.274 | נעשו |  |
| 8:19.275 | דורשי |  |
| 8:19.276 | דרשות |  |
| 8:19.277 | של |  |
| 8:19.278 | דופי |  |
| 8:19.279 | והזיות |  |
| 8:19.280 | כחרשים | ! |
| 8:19.281 | ובזה |  |
| 8:19.282 | נתבטל |  |
| 8:19.283 | יצרא |  |
| 8:19.284 | דע"ז | , |
| 8:19.285 | ומן |  |
| 8:19.286 | השמים |  |
| 8:19.287 | אסתייעא |  |
| 8:19.288 | מלתא | , |
| 8:19.289 | ועלתה |  |
| 8:19.290 | בידם |  |
| 8:19.291 | לשכח |  |
| 8:19.292 | שמות |  |
| 8:19.293 | האלילים |  |
| 8:19.294 | אשר |  |
| 8:19.295 | ידמו |  |
| 8:19.296 | התרים |  |
| 8:19.297 | אחרי |  |
| 8:19.298 | לבבם |  |
| 8:19.299 | לחשוב |  |
| 8:19.300 | רבוי |  |
| 8:19.301 | באלוהות | , |
| 8:19.302 | ושיש |  |
| 8:19.303 | כמה |  |
| 8:19.304 | עלות |  |
| 8:19.305 | וסבות |  |
| 8:19.306 | המשתלשלים |  |
| 8:19.307 | זה |  |
| 8:19.308 | מזה | , |
| 8:19.309 | ונוטלים |  |
| 8:19.310 | רשות |  |
| 8:19.311 | זה |  |
| 8:19.312 | מזה |  |
| 8:19.313 | לברוא |  |
| 8:19.314 | את |  |
| 8:19.315 | אשר |  |
| 8:19.316 | יחפוץ |  |
| 8:19.317 | לברוא |  |
| 8:19.318 | אותו |  |
| 8:19.319 | כאימא |  |
| 8:19.320 | שנטלה |  |
| 8:19.321 | רשות |  |
| 8:19.322 | מאבא |  |
| 8:19.323 | לברוא |  |
| 8:19.324 | אדם |  |
| 8:19.325 | הראשון | , |
| 8:19.326 | וכל |  |
| 8:19.327 | אחד |  |
| 8:19.328 | מאלו |  |
| 8:19.329 | הצורות |  |
| 8:19.330 | נקרא |  |
| 8:19.331 | בשם |  |
| 8:19.332 | ידוע | , |
| 8:19.333 | ונקרא |  |
| 8:19.334 | בשם |  |
| 8:19.335 | הוי'ה |  |
| 8:19.336 | ואדנות |  |
| 8:19.337 | ושאר |  |
| 8:19.338 | השמות |  |
| 8:19.339 | המיוחדים |  |
| 8:19.340 | לשי"ת | , |
| 8:19.341 | ושהעבודה |  |
| 8:19.342 | והתפלה |  |
| 8:19.343 | להם |  |
| 8:19.344 | לבדם |  |
| 8:19.345 | אעפ"י |  |
| 8:19.346 | שהם |  |
| 8:19.347 | נבראים | , |
| 8:19.348 | אבל |  |
| 8:19.349 | האלוה |  |
| 8:19.350 | העליון |  |
| 8:19.351 | איננו |  |
| 8:19.352 | ראוי |  |
| 8:19.353 | לשום |  |
| 8:19.354 | שבח |  |
| 8:19.355 | ותהלה |  |
| 8:19.356 | לרוב |  |
| 8:19.357 | מעלתו |  |
| 8:19.358 | והעלמו | , |
| 8:19.359 | ואין |  |
| 8:19.360 | לו |  |
| 8:19.361 | דמות |  |
| 8:19.362 | וצורה |  |
| 8:19.363 | ואין |  |
| 8:19.364 | לו |  |
| 8:19.365 | אשה |  |
| 8:19.366 | ולא |  |
| 8:19.367 | שם |  |
| 8:19.368 | מוגבל | ! |
| 8:19.369 | כל |  |
| 8:19.370 | מחשבות |  |
| 8:19.371 | מינות |  |
| 8:19.372 | הללו |  |
| 8:19.373 | השכיחום |  |
| 8:19.374 | רבותינו |  |
| 8:19.375 | אנשי |  |
| 8:19.376 | כנסת |  |
| 8:19.377 | הגדולה |  |
| 8:19.378 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

