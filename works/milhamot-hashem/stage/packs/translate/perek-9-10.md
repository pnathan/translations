# Translation pack: milhamot-hashem perek-9-10 (Milchamot ha-Shem 9:12–9:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-10`.

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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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


## Milchamot ha-Shem 9:12

Hebrew (Wikisource): קכט) מי לא ישכיל ויבין לומר שהשאלה הנכרית ששאל השואל את החסיד היתה זרה מאד כעין שאלת פלימו לר', ולמה זה יחשוב על החסיד ההוא מחשבות זרות ונכריות כי הוא מזלזל ח"ו בהעסק בדיני הקרבנות וטומאה וטהרה ותרומה וכיוצא? אלא ודאי רוע לב יש בכאן ושנאת חנם קננה בלבם על חכמי האמת המתכוננים במעשה השי"ת בעולמות שברא כדי להכירו מתוך מעשיו כמו שאמר דוד המלך ע"ה כי אראה שמיך וכו' מה רבו מעשיך ה', ונאמר שאו מרום עיניכם וראו מי ברא אלה לא שיראה אותם כשור וכחמור ח"ו אלא שיתבונן בתקופתן ומהלכיהן במסלותם, ומתוכן יכיר גדולת יוצרם וחכמתו כי נפלאת היא עד מאד, כמו שאמר דוד נפלאים מעשיך וכו'. וכל אותן הראיות שמביא הר"ר מנחם ישאם רוח, ואין כאן מקום להאריך להשיב על חטאעתיו בתרמית האמת המעוור את הפתי שהרי ביאר בעל חובת הלבבות דבריו שם בהקדמה שבתחלה ילמוד האדם ויאמין בכל מצות התורה על ידי חכמי התורה והסמך ויסמוך עד דברי קבלתם, ואחר כך ישיב אל דעתו וישמש בשכלו אחר שיעמוד עליו מצד הקבלה שיהא כוללת את מצות התורה ושרשיהן ופרקיהן ותחקור עליו בשכלך עליו עד שיתברר לך אמתתן וידחה השקר, כדכתיב וידעת היום והשבות אל לבבך וכו' ויצא זה מן הכלל דהיינו יחוד השם ללמד על הכלל כולו, כמו שאמרו רז"ל דבר שהיה בכלל ויוצא מן הכלל ללמד לא ללמד על עצמו יצא אלא ללמד על הכלל יצא עכ"ל. אלמלא שהרב בעל החובות מזהיר שכל מצות התורה צריך ללמוד שרשיהן וחילוקיהן כמו שקבלו חכמים ז"ל, ואח"כ יתבונן בהם עד שיתברר לו אמתתן ככל מה דאפשר, ואשר לא יוכל להשיג כפרה אדומה ושעטנז ובשר בחלב וכיוצא בהם מהחקים יסמוך על הקבלה שכל מצות התורה ניתנו לתועלתינו לטוב לנו כל הימים, אע"פ שלא ידענו מעמם ותכלית הטוב הנמשך מהם אע"פ שהשטן מקטרג עליהן, ואומות העולם משיבין עליהן אזליגן בתר רובא שכיון שנגלה לנו התכלית הטוב ברוב המצות כמו שאמרו רז"ל גם באלה החקים אשר לא נדע תכליתן נסמוך על נותן התורה, כמו שאמרו רז"ל חקה חקקתי גזרה גזרתי אין לך רשות להרהר בהן. קל) וצרעת ממארת ומכה טריה פשתה ורבתה בעמנו היא שנאת החכמה והמדע אשר בה ידע אברהם אבינו והכיר את בוראו טרם הגלותו עליו כמו שהעירו רז"ל במדרשיהם שהיה חוקר ומתבונן במעשי השי"ת בשמים ובארץ ובכל הנמצא בהן עד שהשיג שישי לעולם מנהיג ואז נגלה אליו הקב"ה ואמר לו אני בעל הבירה. והמקובלים החדשים שונאי החכמה לא ידעו להפריש בין פילסופיא טמאה של הטעאות ושבושים לפילסופיא הבנויה על השכל הישר אשר חנן אלהים את האדם והיתה שם פילסופיא בפיהם לשמצה לגנות בה את החכמים היודעים חהבין ולהשכיל במעשה השי"ת ולדעת אותו ידיעה ברורה ולתת מופתים ברורים על מציאותו ית' ויחודו והשגחתו הנפלאה על כל ברואיו מקרני רמים ועד ביצי כנים. ויתמהו איש אל רעהו תמיהה קיימת אצלם על דברי החסיד הר' בחיי בעל חובת הלבבות על שקנה מן החכמה אבל זולתם עובדים זולתו. ותמיהתם הלזו שבורה כשבר נבל יוצרים ותמיהתם ותלונתם אל חיקם תשוב איך הם מתלונים על החסיד השלם בעל חובת הלבבות וקבלו דברי הזהר הבא להם מיד העכו"ם שאמר האלהים סבות רבות כנזכר לעיל פעמים רבות ושהיראה שצונו השי"ת היא מאשתו אבל העבודה היא לזעיר אנפין ולא לסבות שקדמו לו היש עבודת זולתו יותר מזה? ואלו נתנו לבם בדעת ותבונה לדברי הרס"ג והרמב"ם וכו' לא עבדו זולתו.

| id | word | punctuation after |
|---|---|---|
| 9:12.1 | קכט | ) |
| 9:12.2 | מי |  |
| 9:12.3 | לא |  |
| 9:12.4 | ישכיל |  |
| 9:12.5 | ויבין |  |
| 9:12.6 | לומר |  |
| 9:12.7 | שהשאלה |  |
| 9:12.8 | הנכרית |  |
| 9:12.9 | ששאל |  |
| 9:12.10 | השואל |  |
| 9:12.11 | את |  |
| 9:12.12 | החסיד |  |
| 9:12.13 | היתה |  |
| 9:12.14 | זרה |  |
| 9:12.15 | מאד |  |
| 9:12.16 | כעין |  |
| 9:12.17 | שאלת |  |
| 9:12.18 | פלימו |  |
| 9:12.19 | לר | ', |
| 9:12.20 | ולמה |  |
| 9:12.21 | זה |  |
| 9:12.22 | יחשוב |  |
| 9:12.23 | על |  |
| 9:12.24 | החסיד |  |
| 9:12.25 | ההוא |  |
| 9:12.26 | מחשבות |  |
| 9:12.27 | זרות |  |
| 9:12.28 | ונכריות |  |
| 9:12.29 | כי |  |
| 9:12.30 | הוא |  |
| 9:12.31 | מזלזל |  |
| 9:12.32 | ח"ו |  |
| 9:12.33 | בהעסק |  |
| 9:12.34 | בדיני |  |
| 9:12.35 | הקרבנות |  |
| 9:12.36 | וטומאה |  |
| 9:12.37 | וטהרה |  |
| 9:12.38 | ותרומה |  |
| 9:12.39 | וכיוצא | ? |
| 9:12.40 | אלא |  |
| 9:12.41 | ודאי |  |
| 9:12.42 | רוע |  |
| 9:12.43 | לב |  |
| 9:12.44 | יש |  |
| 9:12.45 | בכאן |  |
| 9:12.46 | ושנאת |  |
| 9:12.47 | חנם |  |
| 9:12.48 | קננה |  |
| 9:12.49 | בלבם |  |
| 9:12.50 | על |  |
| 9:12.51 | חכמי |  |
| 9:12.52 | האמת |  |
| 9:12.53 | המתכוננים |  |
| 9:12.54 | במעשה |  |
| 9:12.55 | השי"ת |  |
| 9:12.56 | בעולמות |  |
| 9:12.57 | שברא |  |
| 9:12.58 | כדי |  |
| 9:12.59 | להכירו |  |
| 9:12.60 | מתוך |  |
| 9:12.61 | מעשיו |  |
| 9:12.62 | כמו |  |
| 9:12.63 | שאמר |  |
| 9:12.64 | דוד |  |
| 9:12.65 | המלך |  |
| 9:12.66 | ע"ה |  |
| 9:12.67 | כי |  |
| 9:12.68 | אראה |  |
| 9:12.69 | שמיך |  |
| 9:12.70 | וכו | ' |
| 9:12.71 | מה |  |
| 9:12.72 | רבו |  |
| 9:12.73 | מעשיך |  |
| 9:12.74 | ה | ', |
| 9:12.75 | ונאמר |  |
| 9:12.76 | שאו |  |
| 9:12.77 | מרום |  |
| 9:12.78 | עיניכם |  |
| 9:12.79 | וראו |  |
| 9:12.80 | מי |  |
| 9:12.81 | ברא |  |
| 9:12.82 | אלה |  |
| 9:12.83 | לא |  |
| 9:12.84 | שיראה |  |
| 9:12.85 | אותם |  |
| 9:12.86 | כשור |  |
| 9:12.87 | וכחמור |  |
| 9:12.88 | ח"ו |  |
| 9:12.89 | אלא |  |
| 9:12.90 | שיתבונן |  |
| 9:12.91 | בתקופתן |  |
| 9:12.92 | ומהלכיהן |  |
| 9:12.93 | במסלותם | , |
| 9:12.94 | ומתוכן |  |
| 9:12.95 | יכיר |  |
| 9:12.96 | גדולת |  |
| 9:12.97 | יוצרם |  |
| 9:12.98 | וחכמתו |  |
| 9:12.99 | כי |  |
| 9:12.100 | נפלאת |  |
| 9:12.101 | היא |  |
| 9:12.102 | עד |  |
| 9:12.103 | מאד | , |
| 9:12.104 | כמו |  |
| 9:12.105 | שאמר |  |
| 9:12.106 | דוד |  |
| 9:12.107 | נפלאים |  |
| 9:12.108 | מעשיך |  |
| 9:12.109 | וכו | '. |
| 9:12.110 | וכל |  |
| 9:12.111 | אותן |  |
| 9:12.112 | הראיות |  |
| 9:12.113 | שמביא |  |
| 9:12.114 | הר"ר |  |
| 9:12.115 | מנחם |  |
| 9:12.116 | ישאם |  |
| 9:12.117 | רוח | , |
| 9:12.118 | ואין |  |
| 9:12.119 | כאן |  |
| 9:12.120 | מקום |  |
| 9:12.121 | להאריך |  |
| 9:12.122 | להשיב |  |
| 9:12.123 | על |  |
| 9:12.124 | חטאעתיו |  |
| 9:12.125 | בתרמית |  |
| 9:12.126 | האמת |  |
| 9:12.127 | המעוור |  |
| 9:12.128 | את |  |
| 9:12.129 | הפתי |  |
| 9:12.130 | שהרי |  |
| 9:12.131 | ביאר |  |
| 9:12.132 | בעל |  |
| 9:12.133 | חובת |  |
| 9:12.134 | הלבבות |  |
| 9:12.135 | דבריו |  |
| 9:12.136 | שם |  |
| 9:12.137 | בהקדמה |  |
| 9:12.138 | שבתחלה |  |
| 9:12.139 | ילמוד |  |
| 9:12.140 | האדם |  |
| 9:12.141 | ויאמין |  |
| 9:12.142 | בכל |  |
| 9:12.143 | מצות |  |
| 9:12.144 | התורה |  |
| 9:12.145 | על |  |
| 9:12.146 | ידי |  |
| 9:12.147 | חכמי |  |
| 9:12.148 | התורה |  |
| 9:12.149 | והסמך |  |
| 9:12.150 | ויסמוך |  |
| 9:12.151 | עד |  |
| 9:12.152 | דברי |  |
| 9:12.153 | קבלתם | , |
| 9:12.154 | ואחר |  |
| 9:12.155 | כך |  |
| 9:12.156 | ישיב |  |
| 9:12.157 | אל |  |
| 9:12.158 | דעתו |  |
| 9:12.159 | וישמש |  |
| 9:12.160 | בשכלו |  |
| 9:12.161 | אחר |  |
| 9:12.162 | שיעמוד |  |
| 9:12.163 | עליו |  |
| 9:12.164 | מצד |  |
| 9:12.165 | הקבלה |  |
| 9:12.166 | שיהא |  |
| 9:12.167 | כוללת |  |
| 9:12.168 | את |  |
| 9:12.169 | מצות |  |
| 9:12.170 | התורה |  |
| 9:12.171 | ושרשיהן |  |
| 9:12.172 | ופרקיהן |  |
| 9:12.173 | ותחקור |  |
| 9:12.174 | עליו |  |
| 9:12.175 | בשכלך |  |
| 9:12.176 | עליו |  |
| 9:12.177 | עד |  |
| 9:12.178 | שיתברר |  |
| 9:12.179 | לך |  |
| 9:12.180 | אמתתן |  |
| 9:12.181 | וידחה |  |
| 9:12.182 | השקר | , |
| 9:12.183 | כדכתיב |  |
| 9:12.184 | וידעת |  |
| 9:12.185 | היום |  |
| 9:12.186 | והשבות |  |
| 9:12.187 | אל |  |
| 9:12.188 | לבבך |  |
| 9:12.189 | וכו | ' |
| 9:12.190 | ויצא |  |
| 9:12.191 | זה |  |
| 9:12.192 | מן |  |
| 9:12.193 | הכלל |  |
| 9:12.194 | דהיינו |  |
| 9:12.195 | יחוד |  |
| 9:12.196 | השם |  |
| 9:12.197 | ללמד |  |
| 9:12.198 | על |  |
| 9:12.199 | הכלל |  |
| 9:12.200 | כולו | , |
| 9:12.201 | כמו |  |
| 9:12.202 | שאמרו |  |
| 9:12.203 | רז"ל |  |
| 9:12.204 | דבר |  |
| 9:12.205 | שהיה |  |
| 9:12.206 | בכלל |  |
| 9:12.207 | ויוצא |  |
| 9:12.208 | מן |  |
| 9:12.209 | הכלל |  |
| 9:12.210 | ללמד |  |
| 9:12.211 | לא |  |
| 9:12.212 | ללמד |  |
| 9:12.213 | על |  |
| 9:12.214 | עצמו |  |
| 9:12.215 | יצא |  |
| 9:12.216 | אלא |  |
| 9:12.217 | ללמד |  |
| 9:12.218 | על |  |
| 9:12.219 | הכלל |  |
| 9:12.220 | יצא |  |
| 9:12.221 | עכ"ל | . |
| 9:12.222 | אלמלא |  |
| 9:12.223 | שהרב |  |
| 9:12.224 | בעל |  |
| 9:12.225 | החובות |  |
| 9:12.226 | מזהיר |  |
| 9:12.227 | שכל |  |
| 9:12.228 | מצות |  |
| 9:12.229 | התורה |  |
| 9:12.230 | צריך |  |
| 9:12.231 | ללמוד |  |
| 9:12.232 | שרשיהן |  |
| 9:12.233 | וחילוקיהן |  |
| 9:12.234 | כמו |  |
| 9:12.235 | שקבלו |  |
| 9:12.236 | חכמים |  |
| 9:12.237 | ז"ל | , |
| 9:12.238 | ואח"כ |  |
| 9:12.239 | יתבונן |  |
| 9:12.240 | בהם |  |
| 9:12.241 | עד |  |
| 9:12.242 | שיתברר |  |
| 9:12.243 | לו |  |
| 9:12.244 | אמתתן |  |
| 9:12.245 | ככל |  |
| 9:12.246 | מה |  |
| 9:12.247 | דאפשר | , |
| 9:12.248 | ואשר |  |
| 9:12.249 | לא |  |
| 9:12.250 | יוכל |  |
| 9:12.251 | להשיג |  |
| 9:12.252 | כפרה |  |
| 9:12.253 | אדומה |  |
| 9:12.254 | ושעטנז |  |
| 9:12.255 | ובשר |  |
| 9:12.256 | בחלב |  |
| 9:12.257 | וכיוצא |  |
| 9:12.258 | בהם |  |
| 9:12.259 | מהחקים |  |
| 9:12.260 | יסמוך |  |
| 9:12.261 | על |  |
| 9:12.262 | הקבלה |  |
| 9:12.263 | שכל |  |
| 9:12.264 | מצות |  |
| 9:12.265 | התורה |  |
| 9:12.266 | ניתנו |  |
| 9:12.267 | לתועלתינו |  |
| 9:12.268 | לטוב |  |
| 9:12.269 | לנו |  |
| 9:12.270 | כל |  |
| 9:12.271 | הימים | , |
| 9:12.272 | אע"פ |  |
| 9:12.273 | שלא |  |
| 9:12.274 | ידענו |  |
| 9:12.275 | מעמם |  |
| 9:12.276 | ותכלית |  |
| 9:12.277 | הטוב |  |
| 9:12.278 | הנמשך |  |
| 9:12.279 | מהם |  |
| 9:12.280 | אע"פ |  |
| 9:12.281 | שהשטן |  |
| 9:12.282 | מקטרג |  |
| 9:12.283 | עליהן | , |
| 9:12.284 | ואומות |  |
| 9:12.285 | העולם |  |
| 9:12.286 | משיבין |  |
| 9:12.287 | עליהן |  |
| 9:12.288 | אזליגן |  |
| 9:12.289 | בתר |  |
| 9:12.290 | רובא |  |
| 9:12.291 | שכיון |  |
| 9:12.292 | שנגלה |  |
| 9:12.293 | לנו |  |
| 9:12.294 | התכלית |  |
| 9:12.295 | הטוב |  |
| 9:12.296 | ברוב |  |
| 9:12.297 | המצות |  |
| 9:12.298 | כמו |  |
| 9:12.299 | שאמרו |  |
| 9:12.300 | רז"ל |  |
| 9:12.301 | גם |  |
| 9:12.302 | באלה |  |
| 9:12.303 | החקים |  |
| 9:12.304 | אשר |  |
| 9:12.305 | לא |  |
| 9:12.306 | נדע |  |
| 9:12.307 | תכליתן |  |
| 9:12.308 | נסמוך |  |
| 9:12.309 | על |  |
| 9:12.310 | נותן |  |
| 9:12.311 | התורה | , |
| 9:12.312 | כמו |  |
| 9:12.313 | שאמרו |  |
| 9:12.314 | רז"ל |  |
| 9:12.315 | חקה |  |
| 9:12.316 | חקקתי |  |
| 9:12.317 | גזרה |  |
| 9:12.318 | גזרתי |  |
| 9:12.319 | אין |  |
| 9:12.320 | לך |  |
| 9:12.321 | רשות |  |
| 9:12.322 | להרהר |  |
| 9:12.323 | בהן | . |
| 9:12.324 | קל | ) |
| 9:12.325 | וצרעת |  |
| 9:12.326 | ממארת |  |
| 9:12.327 | ומכה |  |
| 9:12.328 | טריה |  |
| 9:12.329 | פשתה |  |
| 9:12.330 | ורבתה |  |
| 9:12.331 | בעמנו |  |
| 9:12.332 | היא |  |
| 9:12.333 | שנאת |  |
| 9:12.334 | החכמה |  |
| 9:12.335 | והמדע |  |
| 9:12.336 | אשר |  |
| 9:12.337 | בה |  |
| 9:12.338 | ידע |  |
| 9:12.339 | אברהם |  |
| 9:12.340 | אבינו |  |
| 9:12.341 | והכיר |  |
| 9:12.342 | את |  |
| 9:12.343 | בוראו |  |
| 9:12.344 | טרם |  |
| 9:12.345 | הגלותו |  |
| 9:12.346 | עליו |  |
| 9:12.347 | כמו |  |
| 9:12.348 | שהעירו |  |
| 9:12.349 | רז"ל |  |
| 9:12.350 | במדרשיהם |  |
| 9:12.351 | שהיה |  |
| 9:12.352 | חוקר |  |
| 9:12.353 | ומתבונן |  |
| 9:12.354 | במעשי |  |
| 9:12.355 | השי"ת |  |
| 9:12.356 | בשמים |  |
| 9:12.357 | ובארץ |  |
| 9:12.358 | ובכל |  |
| 9:12.359 | הנמצא |  |
| 9:12.360 | בהן |  |
| 9:12.361 | עד |  |
| 9:12.362 | שהשיג |  |
| 9:12.363 | שישי |  |
| 9:12.364 | לעולם |  |
| 9:12.365 | מנהיג |  |
| 9:12.366 | ואז |  |
| 9:12.367 | נגלה |  |
| 9:12.368 | אליו |  |
| 9:12.369 | הקב"ה |  |
| 9:12.370 | ואמר |  |
| 9:12.371 | לו |  |
| 9:12.372 | אני |  |
| 9:12.373 | בעל |  |
| 9:12.374 | הבירה | . |
| 9:12.375 | והמקובלים |  |
| 9:12.376 | החדשים |  |
| 9:12.377 | שונאי |  |
| 9:12.378 | החכמה |  |
| 9:12.379 | לא |  |
| 9:12.380 | ידעו |  |
| 9:12.381 | להפריש |  |
| 9:12.382 | בין |  |
| 9:12.383 | פילסופיא |  |
| 9:12.384 | טמאה |  |
| 9:12.385 | של |  |
| 9:12.386 | הטעאות |  |
| 9:12.387 | ושבושים |  |
| 9:12.388 | לפילסופיא |  |
| 9:12.389 | הבנויה |  |
| 9:12.390 | על |  |
| 9:12.391 | השכל |  |
| 9:12.392 | הישר |  |
| 9:12.393 | אשר |  |
| 9:12.394 | חנן |  |
| 9:12.395 | אלהים |  |
| 9:12.396 | את |  |
| 9:12.397 | האדם |  |
| 9:12.398 | והיתה |  |
| 9:12.399 | שם |  |
| 9:12.400 | פילסופיא |  |
| 9:12.401 | בפיהם |  |
| 9:12.402 | לשמצה |  |
| 9:12.403 | לגנות |  |
| 9:12.404 | בה |  |
| 9:12.405 | את |  |
| 9:12.406 | החכמים |  |
| 9:12.407 | היודעים |  |
| 9:12.408 | חהבין |  |
| 9:12.409 | ולהשכיל |  |
| 9:12.410 | במעשה |  |
| 9:12.411 | השי"ת |  |
| 9:12.412 | ולדעת |  |
| 9:12.413 | אותו |  |
| 9:12.414 | ידיעה |  |
| 9:12.415 | ברורה |  |
| 9:12.416 | ולתת |  |
| 9:12.417 | מופתים |  |
| 9:12.418 | ברורים |  |
| 9:12.419 | על |  |
| 9:12.420 | מציאותו |  |
| 9:12.421 | ית | ' |
| 9:12.422 | ויחודו |  |
| 9:12.423 | והשגחתו |  |
| 9:12.424 | הנפלאה |  |
| 9:12.425 | על |  |
| 9:12.426 | כל |  |
| 9:12.427 | ברואיו |  |
| 9:12.428 | מקרני |  |
| 9:12.429 | רמים |  |
| 9:12.430 | ועד |  |
| 9:12.431 | ביצי |  |
| 9:12.432 | כנים | . |
| 9:12.433 | ויתמהו |  |
| 9:12.434 | איש |  |
| 9:12.435 | אל |  |
| 9:12.436 | רעהו |  |
| 9:12.437 | תמיהה |  |
| 9:12.438 | קיימת |  |
| 9:12.439 | אצלם |  |
| 9:12.440 | על |  |
| 9:12.441 | דברי |  |
| 9:12.442 | החסיד |  |
| 9:12.443 | הר | ' |
| 9:12.444 | בחיי |  |
| 9:12.445 | בעל |  |
| 9:12.446 | חובת |  |
| 9:12.447 | הלבבות |  |
| 9:12.448 | על |  |
| 9:12.449 | שקנה |  |
| 9:12.450 | מן |  |
| 9:12.451 | החכמה |  |
| 9:12.452 | אבל |  |
| 9:12.453 | זולתם |  |
| 9:12.454 | עובדים |  |
| 9:12.455 | זולתו | . |
| 9:12.456 | ותמיהתם |  |
| 9:12.457 | הלזו |  |
| 9:12.458 | שבורה |  |
| 9:12.459 | כשבר |  |
| 9:12.460 | נבל |  |
| 9:12.461 | יוצרים |  |
| 9:12.462 | ותמיהתם |  |
| 9:12.463 | ותלונתם |  |
| 9:12.464 | אל |  |
| 9:12.465 | חיקם |  |
| 9:12.466 | תשוב |  |
| 9:12.467 | איך |  |
| 9:12.468 | הם |  |
| 9:12.469 | מתלונים |  |
| 9:12.470 | על |  |
| 9:12.471 | החסיד |  |
| 9:12.472 | השלם |  |
| 9:12.473 | בעל |  |
| 9:12.474 | חובת |  |
| 9:12.475 | הלבבות |  |
| 9:12.476 | וקבלו |  |
| 9:12.477 | דברי |  |
| 9:12.478 | הזהר |  |
| 9:12.479 | הבא |  |
| 9:12.480 | להם |  |
| 9:12.481 | מיד |  |
| 9:12.482 | העכו"ם |  |
| 9:12.483 | שאמר |  |
| 9:12.484 | האלהים |  |
| 9:12.485 | סבות |  |
| 9:12.486 | רבות |  |
| 9:12.487 | כנזכר |  |
| 9:12.488 | לעיל |  |
| 9:12.489 | פעמים |  |
| 9:12.490 | רבות |  |
| 9:12.491 | ושהיראה |  |
| 9:12.492 | שצונו |  |
| 9:12.493 | השי"ת |  |
| 9:12.494 | היא |  |
| 9:12.495 | מאשתו |  |
| 9:12.496 | אבל |  |
| 9:12.497 | העבודה |  |
| 9:12.498 | היא |  |
| 9:12.499 | לזעיר |  |
| 9:12.500 | אנפין |  |
| 9:12.501 | ולא |  |
| 9:12.502 | לסבות |  |
| 9:12.503 | שקדמו |  |
| 9:12.504 | לו |  |
| 9:12.505 | היש |  |
| 9:12.506 | עבודת |  |
| 9:12.507 | זולתו |  |
| 9:12.508 | יותר |  |
| 9:12.509 | מזה | ? |
| 9:12.510 | ואלו |  |
| 9:12.511 | נתנו |  |
| 9:12.512 | לבם |  |
| 9:12.513 | בדעת |  |
| 9:12.514 | ותבונה |  |
| 9:12.515 | לדברי |  |
| 9:12.516 | הרס"ג |  |
| 9:12.517 | והרמב"ם |  |
| 9:12.518 | וכו | ' |
| 9:12.519 | לא |  |
| 9:12.520 | עבדו |  |
| 9:12.521 | זולתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

