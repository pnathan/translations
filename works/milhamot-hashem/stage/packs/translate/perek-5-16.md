# Translation pack: milhamot-hashem perek-5-16 (Milchamot ha-Shem 5:24–5:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-16`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
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


## Milchamot ha-Shem 5:24

Hebrew (Wikisource): נ) וכ"כ הרמב"ם פרק א' מהל' ע"ז וז"ל כיון שנגמל איתן זה התחיל לשוטט בדעתו והוא קטן ולחשב ביום ובלילה, והיה תמיה היאך אפשר שיהיה הגלגל הזה נוהג תמיד ואין לו מנהיג, ומי יסבב אותו לפי שאי אפשר שיסבב את עצמו. ולא היה לו מלמד ולא מודיע לו דבר אלא מושקע באור כשדים בין עובדי עבודה זרה הטפשים. ואביו ואמו וכל העם עובדים ע"ז והוא היה עובד עמהם ולבו משוטט ומבין עד שהשיג דרך האמת והבין קו הצדק מדעתו הנכונה והבין וידע שיש שם אלוה אחד והוא מנהיג את הגלגל. והוא האל בורא הכל ואין בכל הנמצאים אלוה חוץ ממנו, וידע שכל העם טועים, ודבר שגרם להם טעות זו שעובדים את הכוכבים ואת הצורות עד שאבד האמת מדעתן. ובן ארבעים שנה הכיר אברהם את בוראו כיון שהכיר וידע התחיל להשיב תשובות על בני אור כשדים ולערוך דין עמהם ולומר שאין זו דרך האמת שאתם הולכים בה וכו' וכיון שהיו העם מתקבצין אליו היה מודיע לכל אחד ואחד כפי דעתו עד שיחזירוהו למוטב לדרך האמת עד שנתקבצו אליו אלפים ורבבות והם אנשי בית אברהם ושתל בלבן העיקר הגדול הזה וחיבר בו ספרים והודיע ליצחק בנו ויצחק ישב ללמד ומחזיר ומזהיר, ויצחק ליעקב ומינהו ללמד ומחזיר כל הנלוים אליו. ויעקב אבינו לימד בניו כולן, והבדיל לוי ומינהו ראש והושיבו בישיבה ללמד דרך ה' ולשמור מצות אברהם, וצוה את בניו שלא יפסיקו מבני לוי ממונה כדי שלא ישתכח הלמוד, והיה הדבר הולך ומתגבר בבני יעקב והנלוים אליהם ונעשית בעולם אומה שהיא יודעת את השם עד שארכו הימים לישראל במצרים, וחזרו ללמוד מעשיהם ולעבוד ע"ז כמותן חוץ משבט לוי שעמד במצות אבותיו ומעולם לא עבד שבט לוי ע"ז עכ"ל. ובהל' קריאת שמע כתוב כשקרבו ימי ישראל למות קבץ בניו וצוום וזריזם על יחוד ה' ועל דרך ה' שהלך בה אברהם ויצחק אביו, ואמר להם בני שמא יש בכם פסלות מי שאינו עומד עמו ביחוד אדון העולם ענו כולם ואמרו שמע ישראל ה' אלהינו ה' אחד (כלומר שמע ממנו אבינו ישראל ה' אלהינו ה' אחד) פתח הזקן ואמר ברוך שם כבוד מלכותו לעולם ועד, לפיכך וכו' ע"כ. דבר הלמד מכל האמור שאברהם אבינו ע"ה לימד לכל העולם אמונת יחוד השם וחיבר בה ספרים והודיעם לבניו ולכל הנמשכים אחריו וישמעאל בנו בכללם דהא קי"ל שישמעאל עשה תשובה. ואפשר שבשכר שעשה ישמעאל תשובה זכו בניו לחזור לאמונת יחוד ה' אחרי ימים רבים על ידי מחמד שלמד אצל ישראל כידוע, אך בזרע יעקב לא פסקה, ונשארה בזרע לוי ובשרידים משאר שבטים עד שבא מחוקקינו אבי כל הנביאים וראש לחכמים משה רבינו ע"ה ונגלו לו מלאך ה' בסנה בשמרו את השבועה שנשבע לאברהם אבינו וגאל את אבותינו ממצרים על ידו, וקירבנו לפני הר סיני, והכתיר אותנו במצות והודיענו דרך עבודתו ויחודו ומה יהיה משפט ע"ז והטועים אחריה. וכאשר קרב עתו של משה רבינו ע"ה להפטר הודיעו הקב"ה מה יהיה לנו באחרית הימים בגלותנו, ואמר לו, הנך שוכב עם אבותיך וקם העם הזה וזנה אחרי אלהי נכר הארץ אשר הוא בא שמה בקרבו וכו'. וכך אירע לנו שהטעו אותנו לסור מאחרי ה' אלהינו ולעבוד אלהים אחרים, ולהאמין באלוהות רבים על ידי ספר הזהר הנמסר לנו מיד מלך ממלכי העכו"ם היושבים במזרח, ויחסו הספר הלזה לרשב"י התנא בשקר ומן הוא והלאה נשתנית אמונתינו מעט מעט עד שנהפכה לזרא. ורוב חכמי ישראל לא נתנו לב לדעת מקורו ומבטן מי יצאו הדברים רק האמינ ולמשאות שוא ומדוחים הממונים בחובו, כאלו מרשב"י התנא הקדוש יצאו הדברים, ולכן לא שמו אל לבם למרורת הפתנים אשר בקרבו כי לקחם בחלק שפתיו והדיחם, ונעשית הקבלה המזוייפת לגבירה! והקבלה האמתית המסורה לנו איש מפי איש ממשה רבי' עד רבי ועד רב אשי ורבינא ור' יוחנן מרה דארעא דישראל נעשית לשפחה בזויה ומאוסה ח"ו ותחתיה רגזה ארץ כמבואר בזהר ויובא לקמן (סי' ס"ז וז"ח).

| id | word | punctuation after |
|---|---|---|
| 5:24.1 | נ | ) |
| 5:24.2 | וכ"כ |  |
| 5:24.3 | הרמב"ם |  |
| 5:24.4 | פרק |  |
| 5:24.5 | א | ' |
| 5:24.6 | מהל | ' |
| 5:24.7 | ע"ז |  |
| 5:24.8 | וז"ל |  |
| 5:24.9 | כיון |  |
| 5:24.10 | שנגמל |  |
| 5:24.11 | איתן |  |
| 5:24.12 | זה |  |
| 5:24.13 | התחיל |  |
| 5:24.14 | לשוטט |  |
| 5:24.15 | בדעתו |  |
| 5:24.16 | והוא |  |
| 5:24.17 | קטן |  |
| 5:24.18 | ולחשב |  |
| 5:24.19 | ביום |  |
| 5:24.20 | ובלילה | , |
| 5:24.21 | והיה |  |
| 5:24.22 | תמיה |  |
| 5:24.23 | היאך |  |
| 5:24.24 | אפשר |  |
| 5:24.25 | שיהיה |  |
| 5:24.26 | הגלגל |  |
| 5:24.27 | הזה |  |
| 5:24.28 | נוהג |  |
| 5:24.29 | תמיד |  |
| 5:24.30 | ואין |  |
| 5:24.31 | לו |  |
| 5:24.32 | מנהיג | , |
| 5:24.33 | ומי |  |
| 5:24.34 | יסבב |  |
| 5:24.35 | אותו |  |
| 5:24.36 | לפי |  |
| 5:24.37 | שאי |  |
| 5:24.38 | אפשר |  |
| 5:24.39 | שיסבב |  |
| 5:24.40 | את |  |
| 5:24.41 | עצמו | . |
| 5:24.42 | ולא |  |
| 5:24.43 | היה |  |
| 5:24.44 | לו |  |
| 5:24.45 | מלמד |  |
| 5:24.46 | ולא |  |
| 5:24.47 | מודיע |  |
| 5:24.48 | לו |  |
| 5:24.49 | דבר |  |
| 5:24.50 | אלא |  |
| 5:24.51 | מושקע |  |
| 5:24.52 | באור |  |
| 5:24.53 | כשדים |  |
| 5:24.54 | בין |  |
| 5:24.55 | עובדי |  |
| 5:24.56 | עבודה |  |
| 5:24.57 | זרה |  |
| 5:24.58 | הטפשים | . |
| 5:24.59 | ואביו |  |
| 5:24.60 | ואמו |  |
| 5:24.61 | וכל |  |
| 5:24.62 | העם |  |
| 5:24.63 | עובדים |  |
| 5:24.64 | ע"ז |  |
| 5:24.65 | והוא |  |
| 5:24.66 | היה |  |
| 5:24.67 | עובד |  |
| 5:24.68 | עמהם |  |
| 5:24.69 | ולבו |  |
| 5:24.70 | משוטט |  |
| 5:24.71 | ומבין |  |
| 5:24.72 | עד |  |
| 5:24.73 | שהשיג |  |
| 5:24.74 | דרך |  |
| 5:24.75 | האמת |  |
| 5:24.76 | והבין |  |
| 5:24.77 | קו |  |
| 5:24.78 | הצדק |  |
| 5:24.79 | מדעתו |  |
| 5:24.80 | הנכונה |  |
| 5:24.81 | והבין |  |
| 5:24.82 | וידע |  |
| 5:24.83 | שיש |  |
| 5:24.84 | שם |  |
| 5:24.85 | אלוה |  |
| 5:24.86 | אחד |  |
| 5:24.87 | והוא |  |
| 5:24.88 | מנהיג |  |
| 5:24.89 | את |  |
| 5:24.90 | הגלגל | . |
| 5:24.91 | והוא |  |
| 5:24.92 | האל |  |
| 5:24.93 | בורא |  |
| 5:24.94 | הכל |  |
| 5:24.95 | ואין |  |
| 5:24.96 | בכל |  |
| 5:24.97 | הנמצאים |  |
| 5:24.98 | אלוה |  |
| 5:24.99 | חוץ |  |
| 5:24.100 | ממנו | , |
| 5:24.101 | וידע |  |
| 5:24.102 | שכל |  |
| 5:24.103 | העם |  |
| 5:24.104 | טועים | , |
| 5:24.105 | ודבר |  |
| 5:24.106 | שגרם |  |
| 5:24.107 | להם |  |
| 5:24.108 | טעות |  |
| 5:24.109 | זו |  |
| 5:24.110 | שעובדים |  |
| 5:24.111 | את |  |
| 5:24.112 | הכוכבים |  |
| 5:24.113 | ואת |  |
| 5:24.114 | הצורות |  |
| 5:24.115 | עד |  |
| 5:24.116 | שאבד |  |
| 5:24.117 | האמת |  |
| 5:24.118 | מדעתן | . |
| 5:24.119 | ובן |  |
| 5:24.120 | ארבעים |  |
| 5:24.121 | שנה |  |
| 5:24.122 | הכיר |  |
| 5:24.123 | אברהם |  |
| 5:24.124 | את |  |
| 5:24.125 | בוראו |  |
| 5:24.126 | כיון |  |
| 5:24.127 | שהכיר |  |
| 5:24.128 | וידע |  |
| 5:24.129 | התחיל |  |
| 5:24.130 | להשיב |  |
| 5:24.131 | תשובות |  |
| 5:24.132 | על |  |
| 5:24.133 | בני |  |
| 5:24.134 | אור |  |
| 5:24.135 | כשדים |  |
| 5:24.136 | ולערוך |  |
| 5:24.137 | דין |  |
| 5:24.138 | עמהם |  |
| 5:24.139 | ולומר |  |
| 5:24.140 | שאין |  |
| 5:24.141 | זו |  |
| 5:24.142 | דרך |  |
| 5:24.143 | האמת |  |
| 5:24.144 | שאתם |  |
| 5:24.145 | הולכים |  |
| 5:24.146 | בה |  |
| 5:24.147 | וכו | ' |
| 5:24.148 | וכיון |  |
| 5:24.149 | שהיו |  |
| 5:24.150 | העם |  |
| 5:24.151 | מתקבצין |  |
| 5:24.152 | אליו |  |
| 5:24.153 | היה |  |
| 5:24.154 | מודיע |  |
| 5:24.155 | לכל |  |
| 5:24.156 | אחד |  |
| 5:24.157 | ואחד |  |
| 5:24.158 | כפי |  |
| 5:24.159 | דעתו |  |
| 5:24.160 | עד |  |
| 5:24.161 | שיחזירוהו |  |
| 5:24.162 | למוטב |  |
| 5:24.163 | לדרך |  |
| 5:24.164 | האמת |  |
| 5:24.165 | עד |  |
| 5:24.166 | שנתקבצו |  |
| 5:24.167 | אליו |  |
| 5:24.168 | אלפים |  |
| 5:24.169 | ורבבות |  |
| 5:24.170 | והם |  |
| 5:24.171 | אנשי |  |
| 5:24.172 | בית |  |
| 5:24.173 | אברהם |  |
| 5:24.174 | ושתל |  |
| 5:24.175 | בלבן |  |
| 5:24.176 | העיקר |  |
| 5:24.177 | הגדול |  |
| 5:24.178 | הזה |  |
| 5:24.179 | וחיבר |  |
| 5:24.180 | בו |  |
| 5:24.181 | ספרים |  |
| 5:24.182 | והודיע |  |
| 5:24.183 | ליצחק |  |
| 5:24.184 | בנו |  |
| 5:24.185 | ויצחק |  |
| 5:24.186 | ישב |  |
| 5:24.187 | ללמד |  |
| 5:24.188 | ומחזיר |  |
| 5:24.189 | ומזהיר | , |
| 5:24.190 | ויצחק |  |
| 5:24.191 | ליעקב |  |
| 5:24.192 | ומינהו |  |
| 5:24.193 | ללמד |  |
| 5:24.194 | ומחזיר |  |
| 5:24.195 | כל |  |
| 5:24.196 | הנלוים |  |
| 5:24.197 | אליו | . |
| 5:24.198 | ויעקב |  |
| 5:24.199 | אבינו |  |
| 5:24.200 | לימד |  |
| 5:24.201 | בניו |  |
| 5:24.202 | כולן | , |
| 5:24.203 | והבדיל |  |
| 5:24.204 | לוי |  |
| 5:24.205 | ומינהו |  |
| 5:24.206 | ראש |  |
| 5:24.207 | והושיבו |  |
| 5:24.208 | בישיבה |  |
| 5:24.209 | ללמד |  |
| 5:24.210 | דרך |  |
| 5:24.211 | ה | ' |
| 5:24.212 | ולשמור |  |
| 5:24.213 | מצות |  |
| 5:24.214 | אברהם | , |
| 5:24.215 | וצוה |  |
| 5:24.216 | את |  |
| 5:24.217 | בניו |  |
| 5:24.218 | שלא |  |
| 5:24.219 | יפסיקו |  |
| 5:24.220 | מבני |  |
| 5:24.221 | לוי |  |
| 5:24.222 | ממונה |  |
| 5:24.223 | כדי |  |
| 5:24.224 | שלא |  |
| 5:24.225 | ישתכח |  |
| 5:24.226 | הלמוד | , |
| 5:24.227 | והיה |  |
| 5:24.228 | הדבר |  |
| 5:24.229 | הולך |  |
| 5:24.230 | ומתגבר |  |
| 5:24.231 | בבני |  |
| 5:24.232 | יעקב |  |
| 5:24.233 | והנלוים |  |
| 5:24.234 | אליהם |  |
| 5:24.235 | ונעשית |  |
| 5:24.236 | בעולם |  |
| 5:24.237 | אומה |  |
| 5:24.238 | שהיא |  |
| 5:24.239 | יודעת |  |
| 5:24.240 | את |  |
| 5:24.241 | השם |  |
| 5:24.242 | עד |  |
| 5:24.243 | שארכו |  |
| 5:24.244 | הימים |  |
| 5:24.245 | לישראל |  |
| 5:24.246 | במצרים | , |
| 5:24.247 | וחזרו |  |
| 5:24.248 | ללמוד |  |
| 5:24.249 | מעשיהם |  |
| 5:24.250 | ולעבוד |  |
| 5:24.251 | ע"ז |  |
| 5:24.252 | כמותן |  |
| 5:24.253 | חוץ |  |
| 5:24.254 | משבט |  |
| 5:24.255 | לוי |  |
| 5:24.256 | שעמד |  |
| 5:24.257 | במצות |  |
| 5:24.258 | אבותיו |  |
| 5:24.259 | ומעולם |  |
| 5:24.260 | לא |  |
| 5:24.261 | עבד |  |
| 5:24.262 | שבט |  |
| 5:24.263 | לוי |  |
| 5:24.264 | ע"ז |  |
| 5:24.265 | עכ"ל | . |
| 5:24.266 | ובהל | ' |
| 5:24.267 | קריאת |  |
| 5:24.268 | שמע |  |
| 5:24.269 | כתוב |  |
| 5:24.270 | כשקרבו |  |
| 5:24.271 | ימי |  |
| 5:24.272 | ישראל |  |
| 5:24.273 | למות |  |
| 5:24.274 | קבץ |  |
| 5:24.275 | בניו |  |
| 5:24.276 | וצוום |  |
| 5:24.277 | וזריזם |  |
| 5:24.278 | על |  |
| 5:24.279 | יחוד |  |
| 5:24.280 | ה | ' |
| 5:24.281 | ועל |  |
| 5:24.282 | דרך |  |
| 5:24.283 | ה | ' |
| 5:24.284 | שהלך |  |
| 5:24.285 | בה |  |
| 5:24.286 | אברהם |  |
| 5:24.287 | ויצחק |  |
| 5:24.288 | אביו | , |
| 5:24.289 | ואמר |  |
| 5:24.290 | להם |  |
| 5:24.291 | בני |  |
| 5:24.292 | שמא |  |
| 5:24.293 | יש |  |
| 5:24.294 | בכם |  |
| 5:24.295 | פסלות |  |
| 5:24.296 | מי |  |
| 5:24.297 | שאינו |  |
| 5:24.298 | עומד |  |
| 5:24.299 | עמו |  |
| 5:24.300 | ביחוד |  |
| 5:24.301 | אדון |  |
| 5:24.302 | העולם |  |
| 5:24.303 | ענו |  |
| 5:24.304 | כולם |  |
| 5:24.305 | ואמרו |  |
| 5:24.306 | שמע |  |
| 5:24.307 | ישראל |  |
| 5:24.308 | ה | ' |
| 5:24.309 | אלהינו |  |
| 5:24.310 | ה | ' |
| 5:24.311 | אחד |  |
| 5:24.312 | כלומר |  |
| 5:24.313 | שמע |  |
| 5:24.314 | ממנו |  |
| 5:24.315 | אבינו |  |
| 5:24.316 | ישראל |  |
| 5:24.317 | ה | ' |
| 5:24.318 | אלהינו |  |
| 5:24.319 | ה | ' |
| 5:24.320 | אחד | ) |
| 5:24.321 | פתח |  |
| 5:24.322 | הזקן |  |
| 5:24.323 | ואמר |  |
| 5:24.324 | ברוך |  |
| 5:24.325 | שם |  |
| 5:24.326 | כבוד |  |
| 5:24.327 | מלכותו |  |
| 5:24.328 | לעולם |  |
| 5:24.329 | ועד | , |
| 5:24.330 | לפיכך |  |
| 5:24.331 | וכו | ' |
| 5:24.332 | ע"כ | . |
| 5:24.333 | דבר |  |
| 5:24.334 | הלמד |  |
| 5:24.335 | מכל |  |
| 5:24.336 | האמור |  |
| 5:24.337 | שאברהם |  |
| 5:24.338 | אבינו |  |
| 5:24.339 | ע"ה |  |
| 5:24.340 | לימד |  |
| 5:24.341 | לכל |  |
| 5:24.342 | העולם |  |
| 5:24.343 | אמונת |  |
| 5:24.344 | יחוד |  |
| 5:24.345 | השם |  |
| 5:24.346 | וחיבר |  |
| 5:24.347 | בה |  |
| 5:24.348 | ספרים |  |
| 5:24.349 | והודיעם |  |
| 5:24.350 | לבניו |  |
| 5:24.351 | ולכל |  |
| 5:24.352 | הנמשכים |  |
| 5:24.353 | אחריו |  |
| 5:24.354 | וישמעאל |  |
| 5:24.355 | בנו |  |
| 5:24.356 | בכללם |  |
| 5:24.357 | דהא |  |
| 5:24.358 | קי"ל |  |
| 5:24.359 | שישמעאל |  |
| 5:24.360 | עשה |  |
| 5:24.361 | תשובה | . |
| 5:24.362 | ואפשר |  |
| 5:24.363 | שבשכר |  |
| 5:24.364 | שעשה |  |
| 5:24.365 | ישמעאל |  |
| 5:24.366 | תשובה |  |
| 5:24.367 | זכו |  |
| 5:24.368 | בניו |  |
| 5:24.369 | לחזור |  |
| 5:24.370 | לאמונת |  |
| 5:24.371 | יחוד |  |
| 5:24.372 | ה | ' |
| 5:24.373 | אחרי |  |
| 5:24.374 | ימים |  |
| 5:24.375 | רבים |  |
| 5:24.376 | על |  |
| 5:24.377 | ידי |  |
| 5:24.378 | מחמד |  |
| 5:24.379 | שלמד |  |
| 5:24.380 | אצל |  |
| 5:24.381 | ישראל |  |
| 5:24.382 | כידוע | , |
| 5:24.383 | אך |  |
| 5:24.384 | בזרע |  |
| 5:24.385 | יעקב |  |
| 5:24.386 | לא |  |
| 5:24.387 | פסקה | , |
| 5:24.388 | ונשארה |  |
| 5:24.389 | בזרע |  |
| 5:24.390 | לוי |  |
| 5:24.391 | ובשרידים |  |
| 5:24.392 | משאר |  |
| 5:24.393 | שבטים |  |
| 5:24.394 | עד |  |
| 5:24.395 | שבא |  |
| 5:24.396 | מחוקקינו |  |
| 5:24.397 | אבי |  |
| 5:24.398 | כל |  |
| 5:24.399 | הנביאים |  |
| 5:24.400 | וראש |  |
| 5:24.401 | לחכמים |  |
| 5:24.402 | משה |  |
| 5:24.403 | רבינו |  |
| 5:24.404 | ע"ה |  |
| 5:24.405 | ונגלו |  |
| 5:24.406 | לו |  |
| 5:24.407 | מלאך |  |
| 5:24.408 | ה | ' |
| 5:24.409 | בסנה |  |
| 5:24.410 | בשמרו |  |
| 5:24.411 | את |  |
| 5:24.412 | השבועה |  |
| 5:24.413 | שנשבע |  |
| 5:24.414 | לאברהם |  |
| 5:24.415 | אבינו |  |
| 5:24.416 | וגאל |  |
| 5:24.417 | את |  |
| 5:24.418 | אבותינו |  |
| 5:24.419 | ממצרים |  |
| 5:24.420 | על |  |
| 5:24.421 | ידו | , |
| 5:24.422 | וקירבנו |  |
| 5:24.423 | לפני |  |
| 5:24.424 | הר |  |
| 5:24.425 | סיני | , |
| 5:24.426 | והכתיר |  |
| 5:24.427 | אותנו |  |
| 5:24.428 | במצות |  |
| 5:24.429 | והודיענו |  |
| 5:24.430 | דרך |  |
| 5:24.431 | עבודתו |  |
| 5:24.432 | ויחודו |  |
| 5:24.433 | ומה |  |
| 5:24.434 | יהיה |  |
| 5:24.435 | משפט |  |
| 5:24.436 | ע"ז |  |
| 5:24.437 | והטועים |  |
| 5:24.438 | אחריה | . |
| 5:24.439 | וכאשר |  |
| 5:24.440 | קרב |  |
| 5:24.441 | עתו |  |
| 5:24.442 | של |  |
| 5:24.443 | משה |  |
| 5:24.444 | רבינו |  |
| 5:24.445 | ע"ה |  |
| 5:24.446 | להפטר |  |
| 5:24.447 | הודיעו |  |
| 5:24.448 | הקב"ה |  |
| 5:24.449 | מה |  |
| 5:24.450 | יהיה |  |
| 5:24.451 | לנו |  |
| 5:24.452 | באחרית |  |
| 5:24.453 | הימים |  |
| 5:24.454 | בגלותנו | , |
| 5:24.455 | ואמר |  |
| 5:24.456 | לו | , |
| 5:24.457 | הנך |  |
| 5:24.458 | שוכב |  |
| 5:24.459 | עם |  |
| 5:24.460 | אבותיך |  |
| 5:24.461 | וקם |  |
| 5:24.462 | העם |  |
| 5:24.463 | הזה |  |
| 5:24.464 | וזנה |  |
| 5:24.465 | אחרי |  |
| 5:24.466 | אלהי |  |
| 5:24.467 | נכר |  |
| 5:24.468 | הארץ |  |
| 5:24.469 | אשר |  |
| 5:24.470 | הוא |  |
| 5:24.471 | בא |  |
| 5:24.472 | שמה |  |
| 5:24.473 | בקרבו |  |
| 5:24.474 | וכו | '. |
| 5:24.475 | וכך |  |
| 5:24.476 | אירע |  |
| 5:24.477 | לנו |  |
| 5:24.478 | שהטעו |  |
| 5:24.479 | אותנו |  |
| 5:24.480 | לסור |  |
| 5:24.481 | מאחרי |  |
| 5:24.482 | ה | ' |
| 5:24.483 | אלהינו |  |
| 5:24.484 | ולעבוד |  |
| 5:24.485 | אלהים |  |
| 5:24.486 | אחרים | , |
| 5:24.487 | ולהאמין |  |
| 5:24.488 | באלוהות |  |
| 5:24.489 | רבים |  |
| 5:24.490 | על |  |
| 5:24.491 | ידי |  |
| 5:24.492 | ספר |  |
| 5:24.493 | הזהר |  |
| 5:24.494 | הנמסר |  |
| 5:24.495 | לנו |  |
| 5:24.496 | מיד |  |
| 5:24.497 | מלך |  |
| 5:24.498 | ממלכי |  |
| 5:24.499 | העכו"ם |  |
| 5:24.500 | היושבים |  |
| 5:24.501 | במזרח | , |
| 5:24.502 | ויחסו |  |
| 5:24.503 | הספר |  |
| 5:24.504 | הלזה |  |
| 5:24.505 | לרשב"י |  |
| 5:24.506 | התנא |  |
| 5:24.507 | בשקר |  |
| 5:24.508 | ומן |  |
| 5:24.509 | הוא |  |
| 5:24.510 | והלאה |  |
| 5:24.511 | נשתנית |  |
| 5:24.512 | אמונתינו |  |
| 5:24.513 | מעט |  |
| 5:24.514 | מעט |  |
| 5:24.515 | עד |  |
| 5:24.516 | שנהפכה |  |
| 5:24.517 | לזרא | . |
| 5:24.518 | ורוב |  |
| 5:24.519 | חכמי |  |
| 5:24.520 | ישראל |  |
| 5:24.521 | לא |  |
| 5:24.522 | נתנו |  |
| 5:24.523 | לב |  |
| 5:24.524 | לדעת |  |
| 5:24.525 | מקורו |  |
| 5:24.526 | ומבטן |  |
| 5:24.527 | מי |  |
| 5:24.528 | יצאו |  |
| 5:24.529 | הדברים |  |
| 5:24.530 | רק |  |
| 5:24.531 | האמינ |  |
| 5:24.532 | ולמשאות |  |
| 5:24.533 | שוא |  |
| 5:24.534 | ומדוחים |  |
| 5:24.535 | הממונים |  |
| 5:24.536 | בחובו | , |
| 5:24.537 | כאלו |  |
| 5:24.538 | מרשב"י |  |
| 5:24.539 | התנא |  |
| 5:24.540 | הקדוש |  |
| 5:24.541 | יצאו |  |
| 5:24.542 | הדברים | , |
| 5:24.543 | ולכן |  |
| 5:24.544 | לא |  |
| 5:24.545 | שמו |  |
| 5:24.546 | אל |  |
| 5:24.547 | לבם |  |
| 5:24.548 | למרורת |  |
| 5:24.549 | הפתנים |  |
| 5:24.550 | אשר |  |
| 5:24.551 | בקרבו |  |
| 5:24.552 | כי |  |
| 5:24.553 | לקחם |  |
| 5:24.554 | בחלק |  |
| 5:24.555 | שפתיו |  |
| 5:24.556 | והדיחם | , |
| 5:24.557 | ונעשית |  |
| 5:24.558 | הקבלה |  |
| 5:24.559 | המזוייפת |  |
| 5:24.560 | לגבירה | ! |
| 5:24.561 | והקבלה |  |
| 5:24.562 | האמתית |  |
| 5:24.563 | המסורה |  |
| 5:24.564 | לנו |  |
| 5:24.565 | איש |  |
| 5:24.566 | מפי |  |
| 5:24.567 | איש |  |
| 5:24.568 | ממשה |  |
| 5:24.569 | רבי | ' |
| 5:24.570 | עד |  |
| 5:24.571 | רבי |  |
| 5:24.572 | ועד |  |
| 5:24.573 | רב |  |
| 5:24.574 | אשי |  |
| 5:24.575 | ורבינא |  |
| 5:24.576 | ור | ' |
| 5:24.577 | יוחנן |  |
| 5:24.578 | מרה |  |
| 5:24.579 | דארעא |  |
| 5:24.580 | דישראל |  |
| 5:24.581 | נעשית |  |
| 5:24.582 | לשפחה |  |
| 5:24.583 | בזויה |  |
| 5:24.584 | ומאוסה |  |
| 5:24.585 | ח"ו |  |
| 5:24.586 | ותחתיה |  |
| 5:24.587 | רגזה |  |
| 5:24.588 | ארץ |  |
| 5:24.589 | כמבואר |  |
| 5:24.590 | בזהר |  |
| 5:24.591 | ויובא |  |
| 5:24.592 | לקמן |  |
| 5:24.593 | סי | ' |
| 5:24.594 | ס"ז |  |
| 5:24.595 | וז"ח | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

