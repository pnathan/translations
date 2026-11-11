# Translation pack: milhamot-hashem perek-9-09 (Milchamot ha-Shem 9:11–9:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-09`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- האלהי before a name: "the divine", with a note; never dropped.
- An initial before a surname that fixes the person (ר"מ ריקנטי): expand and
  name him in a note ("Rabbi Menachem Rikanti"); an initial no surname fixes
  stays an initial with a note, as with R. A. and ר"י חיים ("Rabbi Y. Chayyim").
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
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
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 9:11

Hebrew (Wikisource): קכח) עוד זאת רעה חולה, שכל הלומדים בזהר ושאר ספרי קבלה חדשה נקל בעיניהם לעולל עלילות ולשם תהלה במאורי ישראל מורי צדק המאירים עיני ישראל בכל מקצועות התורה אשר יעשה אותם האדם וחי בהם, ובפרט בעיקרי תורתינו איך נאמין ונקיים מצותיה! וביחוד העיקר והעמוד שהכל תלוי בו, והוא יחוד השי"ת לא שנאמר בפינו אחד, ובלבבינו רבים, כמו שאומרים המקובלים כי האלהים העליון דלא שייך ביה עבודה המקשר ומחבר ומזויג את האלהים הרבים התחתונים אשר אותם נעבוד להיותם אחד! אוי לעינים שכך רואות, ואוי לאזנים שכך שומעות, ודברים כאלה הן שאסור לשמען ואסור להרהר בהן אפילו בבית הכסא, כמו שנאמר ולא תתורו אחרי לבבכם וכו' כמו שפירשו רז"ל. ומאמיני הקבלה החדשה הזו אשר נחלוה מיד המלך העכו"ם, ישמיו תהלה במאורי ישראל ומלאכי ה', הרס"ג ובעל חובת הלבבות ורבי יהודה הלוי והרמב"ם ובעל העיקרים כאשר יראה המעיין בספר דרך חיים לרבי מנחם דילונאזנו איך תלה בבעל חה"ל בוקי סריקי מה שלא ניתכוון להם החסיד בעל חובת הלבבות! עוד הוסיף לנטות ידו לבזות את החסיד על מה שכתב בהקדמה שחכם אחד נשאל שאלה נכרית מענייני גירושין, והחכם ההוא השיבו אתה האיש השואל על מה שלא יזיקנו אם לא ידענו הידעת כל מה שאתה חייב לדעתו מן המצות אשר אינך ראשי להתעלם מהן עד שנפנית לחשוב בשאלות נכריות אשר לא תקנה בהם מעלה יתירה בתורתך ואמונתך וכו'. והרחיב הר"מ לונזאנו את פיו לדבר תועה על מלאך ה' צבאות החסיד בעל חובת הלבבות, ולא זכר שר מאי דגרסינן במנחות שרבי גזר על פלימו שיקבל עליו שמתא או יגלה על ששאל לו שאלה נכרית לענין תפלין, דהכי גרסינן בפ' הקומץ, בעא מניה פלימו מר' מי שיש לו שני ראשים באיזה מהם מניח תפילין אמר ליה או קום גלי או קביל עלך שמתא, ואלמלא חס הקב"ה על כבוד צורבא מרבנן פלימו ואתא ההוא נוברא ואמר אתיליד לי ינוקא דאית ליה תרין רישין כמא בעינא למיתב לכהן וכו'! ודברים של הר"מ לונזאנו הי"ל לדבר על ר' שהשיב על שאלת פלימו או קום גלי או קביל עלך שמתא? אלמלא דכל השואל שאלות זרות מאד מנדין אותו, שגורם לבייש את החכם בחנם, וכ"כ מרן בכ"מ סוף פ"ו מהל' ת"ת שהשואל דבר שאי אפשר הוא בכלל מזלזל בד"ת או המבזה את החכם.

| id | word | punctuation after |
|---|---|---|
| 9:11.1 | קכח | ) |
| 9:11.2 | עוד |  |
| 9:11.3 | זאת |  |
| 9:11.4 | רעה |  |
| 9:11.5 | חולה | , |
| 9:11.6 | שכל |  |
| 9:11.7 | הלומדים |  |
| 9:11.8 | בזהר |  |
| 9:11.9 | ושאר |  |
| 9:11.10 | ספרי |  |
| 9:11.11 | קבלה |  |
| 9:11.12 | חדשה |  |
| 9:11.13 | נקל |  |
| 9:11.14 | בעיניהם |  |
| 9:11.15 | לעולל |  |
| 9:11.16 | עלילות |  |
| 9:11.17 | ולשם |  |
| 9:11.18 | תהלה |  |
| 9:11.19 | במאורי |  |
| 9:11.20 | ישראל |  |
| 9:11.21 | מורי |  |
| 9:11.22 | צדק |  |
| 9:11.23 | המאירים |  |
| 9:11.24 | עיני |  |
| 9:11.25 | ישראל |  |
| 9:11.26 | בכל |  |
| 9:11.27 | מקצועות |  |
| 9:11.28 | התורה |  |
| 9:11.29 | אשר |  |
| 9:11.30 | יעשה |  |
| 9:11.31 | אותם |  |
| 9:11.32 | האדם |  |
| 9:11.33 | וחי |  |
| 9:11.34 | בהם | , |
| 9:11.35 | ובפרט |  |
| 9:11.36 | בעיקרי |  |
| 9:11.37 | תורתינו |  |
| 9:11.38 | איך |  |
| 9:11.39 | נאמין |  |
| 9:11.40 | ונקיים |  |
| 9:11.41 | מצותיה | ! |
| 9:11.42 | וביחוד |  |
| 9:11.43 | העיקר |  |
| 9:11.44 | והעמוד |  |
| 9:11.45 | שהכל |  |
| 9:11.46 | תלוי |  |
| 9:11.47 | בו | , |
| 9:11.48 | והוא |  |
| 9:11.49 | יחוד |  |
| 9:11.50 | השי"ת |  |
| 9:11.51 | לא |  |
| 9:11.52 | שנאמר |  |
| 9:11.53 | בפינו |  |
| 9:11.54 | אחד | , |
| 9:11.55 | ובלבבינו |  |
| 9:11.56 | רבים | , |
| 9:11.57 | כמו |  |
| 9:11.58 | שאומרים |  |
| 9:11.59 | המקובלים |  |
| 9:11.60 | כי |  |
| 9:11.61 | האלהים |  |
| 9:11.62 | העליון |  |
| 9:11.63 | דלא |  |
| 9:11.64 | שייך |  |
| 9:11.65 | ביה |  |
| 9:11.66 | עבודה |  |
| 9:11.67 | המקשר |  |
| 9:11.68 | ומחבר |  |
| 9:11.69 | ומזויג |  |
| 9:11.70 | את |  |
| 9:11.71 | האלהים |  |
| 9:11.72 | הרבים |  |
| 9:11.73 | התחתונים |  |
| 9:11.74 | אשר |  |
| 9:11.75 | אותם |  |
| 9:11.76 | נעבוד |  |
| 9:11.77 | להיותם |  |
| 9:11.78 | אחד | ! |
| 9:11.79 | אוי |  |
| 9:11.80 | לעינים |  |
| 9:11.81 | שכך |  |
| 9:11.82 | רואות | , |
| 9:11.83 | ואוי |  |
| 9:11.84 | לאזנים |  |
| 9:11.85 | שכך |  |
| 9:11.86 | שומעות | , |
| 9:11.87 | ודברים |  |
| 9:11.88 | כאלה |  |
| 9:11.89 | הן |  |
| 9:11.90 | שאסור |  |
| 9:11.91 | לשמען |  |
| 9:11.92 | ואסור |  |
| 9:11.93 | להרהר |  |
| 9:11.94 | בהן |  |
| 9:11.95 | אפילו |  |
| 9:11.96 | בבית |  |
| 9:11.97 | הכסא | , |
| 9:11.98 | כמו |  |
| 9:11.99 | שנאמר |  |
| 9:11.100 | ולא |  |
| 9:11.101 | תתורו |  |
| 9:11.102 | אחרי |  |
| 9:11.103 | לבבכם |  |
| 9:11.104 | וכו | ' |
| 9:11.105 | כמו |  |
| 9:11.106 | שפירשו |  |
| 9:11.107 | רז"ל | . |
| 9:11.108 | ומאמיני |  |
| 9:11.109 | הקבלה |  |
| 9:11.110 | החדשה |  |
| 9:11.111 | הזו |  |
| 9:11.112 | אשר |  |
| 9:11.113 | נחלוה |  |
| 9:11.114 | מיד |  |
| 9:11.115 | המלך |  |
| 9:11.116 | העכו"ם | , |
| 9:11.117 | ישמיו |  |
| 9:11.118 | תהלה |  |
| 9:11.119 | במאורי |  |
| 9:11.120 | ישראל |  |
| 9:11.121 | ומלאכי |  |
| 9:11.122 | ה | ', |
| 9:11.123 | הרס"ג |  |
| 9:11.124 | ובעל |  |
| 9:11.125 | חובת |  |
| 9:11.126 | הלבבות |  |
| 9:11.127 | ורבי |  |
| 9:11.128 | יהודה |  |
| 9:11.129 | הלוי |  |
| 9:11.130 | והרמב"ם |  |
| 9:11.131 | ובעל |  |
| 9:11.132 | העיקרים |  |
| 9:11.133 | כאשר |  |
| 9:11.134 | יראה |  |
| 9:11.135 | המעיין |  |
| 9:11.136 | בספר |  |
| 9:11.137 | דרך |  |
| 9:11.138 | חיים |  |
| 9:11.139 | לרבי |  |
| 9:11.140 | מנחם |  |
| 9:11.141 | דילונאזנו |  |
| 9:11.142 | איך |  |
| 9:11.143 | תלה |  |
| 9:11.144 | בבעל |  |
| 9:11.145 | חה"ל |  |
| 9:11.146 | בוקי |  |
| 9:11.147 | סריקי |  |
| 9:11.148 | מה |  |
| 9:11.149 | שלא |  |
| 9:11.150 | ניתכוון |  |
| 9:11.151 | להם |  |
| 9:11.152 | החסיד |  |
| 9:11.153 | בעל |  |
| 9:11.154 | חובת |  |
| 9:11.155 | הלבבות | ! |
| 9:11.156 | עוד |  |
| 9:11.157 | הוסיף |  |
| 9:11.158 | לנטות |  |
| 9:11.159 | ידו |  |
| 9:11.160 | לבזות |  |
| 9:11.161 | את |  |
| 9:11.162 | החסיד |  |
| 9:11.163 | על |  |
| 9:11.164 | מה |  |
| 9:11.165 | שכתב |  |
| 9:11.166 | בהקדמה |  |
| 9:11.167 | שחכם |  |
| 9:11.168 | אחד |  |
| 9:11.169 | נשאל |  |
| 9:11.170 | שאלה |  |
| 9:11.171 | נכרית |  |
| 9:11.172 | מענייני |  |
| 9:11.173 | גירושין | , |
| 9:11.174 | והחכם |  |
| 9:11.175 | ההוא |  |
| 9:11.176 | השיבו |  |
| 9:11.177 | אתה |  |
| 9:11.178 | האיש |  |
| 9:11.179 | השואל |  |
| 9:11.180 | על |  |
| 9:11.181 | מה |  |
| 9:11.182 | שלא |  |
| 9:11.183 | יזיקנו |  |
| 9:11.184 | אם |  |
| 9:11.185 | לא |  |
| 9:11.186 | ידענו |  |
| 9:11.187 | הידעת |  |
| 9:11.188 | כל |  |
| 9:11.189 | מה |  |
| 9:11.190 | שאתה |  |
| 9:11.191 | חייב |  |
| 9:11.192 | לדעתו |  |
| 9:11.193 | מן |  |
| 9:11.194 | המצות |  |
| 9:11.195 | אשר |  |
| 9:11.196 | אינך |  |
| 9:11.197 | ראשי |  |
| 9:11.198 | להתעלם |  |
| 9:11.199 | מהן |  |
| 9:11.200 | עד |  |
| 9:11.201 | שנפנית |  |
| 9:11.202 | לחשוב |  |
| 9:11.203 | בשאלות |  |
| 9:11.204 | נכריות |  |
| 9:11.205 | אשר |  |
| 9:11.206 | לא |  |
| 9:11.207 | תקנה |  |
| 9:11.208 | בהם |  |
| 9:11.209 | מעלה |  |
| 9:11.210 | יתירה |  |
| 9:11.211 | בתורתך |  |
| 9:11.212 | ואמונתך |  |
| 9:11.213 | וכו | '. |
| 9:11.214 | והרחיב |  |
| 9:11.215 | הר"מ |  |
| 9:11.216 | לונזאנו |  |
| 9:11.217 | את |  |
| 9:11.218 | פיו |  |
| 9:11.219 | לדבר |  |
| 9:11.220 | תועה |  |
| 9:11.221 | על |  |
| 9:11.222 | מלאך |  |
| 9:11.223 | ה | ' |
| 9:11.224 | צבאות |  |
| 9:11.225 | החסיד |  |
| 9:11.226 | בעל |  |
| 9:11.227 | חובת |  |
| 9:11.228 | הלבבות | , |
| 9:11.229 | ולא |  |
| 9:11.230 | זכר |  |
| 9:11.231 | שר |  |
| 9:11.232 | מאי |  |
| 9:11.233 | דגרסינן |  |
| 9:11.234 | במנחות |  |
| 9:11.235 | שרבי |  |
| 9:11.236 | גזר |  |
| 9:11.237 | על |  |
| 9:11.238 | פלימו |  |
| 9:11.239 | שיקבל |  |
| 9:11.240 | עליו |  |
| 9:11.241 | שמתא |  |
| 9:11.242 | או |  |
| 9:11.243 | יגלה |  |
| 9:11.244 | על |  |
| 9:11.245 | ששאל |  |
| 9:11.246 | לו |  |
| 9:11.247 | שאלה |  |
| 9:11.248 | נכרית |  |
| 9:11.249 | לענין |  |
| 9:11.250 | תפלין | , |
| 9:11.251 | דהכי |  |
| 9:11.252 | גרסינן |  |
| 9:11.253 | בפ | ' |
| 9:11.254 | הקומץ | , |
| 9:11.255 | בעא |  |
| 9:11.256 | מניה |  |
| 9:11.257 | פלימו |  |
| 9:11.258 | מר | ' |
| 9:11.259 | מי |  |
| 9:11.260 | שיש |  |
| 9:11.261 | לו |  |
| 9:11.262 | שני |  |
| 9:11.263 | ראשים |  |
| 9:11.264 | באיזה |  |
| 9:11.265 | מהם |  |
| 9:11.266 | מניח |  |
| 9:11.267 | תפילין |  |
| 9:11.268 | אמר |  |
| 9:11.269 | ליה |  |
| 9:11.270 | או |  |
| 9:11.271 | קום |  |
| 9:11.272 | גלי |  |
| 9:11.273 | או |  |
| 9:11.274 | קביל |  |
| 9:11.275 | עלך |  |
| 9:11.276 | שמתא | , |
| 9:11.277 | ואלמלא |  |
| 9:11.278 | חס |  |
| 9:11.279 | הקב"ה |  |
| 9:11.280 | על |  |
| 9:11.281 | כבוד |  |
| 9:11.282 | צורבא |  |
| 9:11.283 | מרבנן |  |
| 9:11.284 | פלימו |  |
| 9:11.285 | ואתא |  |
| 9:11.286 | ההוא |  |
| 9:11.287 | נוברא |  |
| 9:11.288 | ואמר |  |
| 9:11.289 | אתיליד |  |
| 9:11.290 | לי |  |
| 9:11.291 | ינוקא |  |
| 9:11.292 | דאית |  |
| 9:11.293 | ליה |  |
| 9:11.294 | תרין |  |
| 9:11.295 | רישין |  |
| 9:11.296 | כמא |  |
| 9:11.297 | בעינא |  |
| 9:11.298 | למיתב |  |
| 9:11.299 | לכהן |  |
| 9:11.300 | וכו | '! |
| 9:11.301 | ודברים |  |
| 9:11.302 | של |  |
| 9:11.303 | הר"מ |  |
| 9:11.304 | לונזאנו |  |
| 9:11.305 | הי"ל |  |
| 9:11.306 | לדבר |  |
| 9:11.307 | על |  |
| 9:11.308 | ר | ' |
| 9:11.309 | שהשיב |  |
| 9:11.310 | על |  |
| 9:11.311 | שאלת |  |
| 9:11.312 | פלימו |  |
| 9:11.313 | או |  |
| 9:11.314 | קום |  |
| 9:11.315 | גלי |  |
| 9:11.316 | או |  |
| 9:11.317 | קביל |  |
| 9:11.318 | עלך |  |
| 9:11.319 | שמתא | ? |
| 9:11.320 | אלמלא |  |
| 9:11.321 | דכל |  |
| 9:11.322 | השואל |  |
| 9:11.323 | שאלות |  |
| 9:11.324 | זרות |  |
| 9:11.325 | מאד |  |
| 9:11.326 | מנדין |  |
| 9:11.327 | אותו | , |
| 9:11.328 | שגורם |  |
| 9:11.329 | לבייש |  |
| 9:11.330 | את |  |
| 9:11.331 | החכם |  |
| 9:11.332 | בחנם | , |
| 9:11.333 | וכ"כ |  |
| 9:11.334 | מרן |  |
| 9:11.335 | בכ"מ |  |
| 9:11.336 | סוף |  |
| 9:11.337 | פ"ו |  |
| 9:11.338 | מהל | ' |
| 9:11.339 | ת"ת |  |
| 9:11.340 | שהשואל |  |
| 9:11.341 | דבר |  |
| 9:11.342 | שאי |  |
| 9:11.343 | אפשר |  |
| 9:11.344 | הוא |  |
| 9:11.345 | בכלל |  |
| 9:11.346 | מזלזל |  |
| 9:11.347 | בד"ת |  |
| 9:11.348 | או |  |
| 9:11.349 | המבזה |  |
| 9:11.350 | את |  |
| 9:11.351 | החכם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

