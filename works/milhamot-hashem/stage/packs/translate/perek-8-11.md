# Translation pack: milhamot-hashem perek-8-11 (Milchamot ha-Shem 8:13–8:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-11`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
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
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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


## Milchamot ha-Shem 8:13

Hebrew (Wikisource): קח) והנה כבר ביארנו כמה פעמים שדעת המקובלים החדשים שהזעיר אנפין הוא שנגלה על אבותינו בהר סיני ואמר אנכי ה' אלהיך, ואם כן איך הביאו רז"ל בשמות רבה פסוק אני ראשון ואני אחרון וכו' שאין לי אב, ולא בן ולא אח, על אנכי ה' אלהיך שאמרו המקובלים החדשים שהזעיר אנפין אמרו, ומה לו לאלהים העליון להשתבח בעצמו במקום זה תוך דברי הזעיר אנפין שיש לו אב ואם ובן ואחות? אלא ודאי שרז"ל בב"ר ובשמות רבה ודברים רבה, באו להעיר אותנו שלא נטעה באלו הדעות, ולבאר לנו, כי זה שאמר אנכי ה' אלהיך הוא שאמר אני ראשון שאין לי אב, ואני אחרון שאין לי בן, ומבלעדי אין אלהים שאין לי אח, הוא ה' אלהינו הסבה הראשונה הוא ולא אחר, והוא בכבודו ובעצמו מן השמים השמיענו את קולו, והזהרנו בדבור השני לא יהיה לך אלהים אחרים על פני וכו'! והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי! ודלא כדעת הזהר המסית דס"ל שהאלוה שאמר אני ראשון וכו' הוא האין סוף, ושאמר ראו עתה וכו' הוא אדם קדמון, ושאמר יהי אור וכו' הוא אבא שאמר לאימא, ושאמר נעשה אדם וכו' היא אימא שאמרה לאבא וכו' כדלעיל, ומכל מקום לא יחסו כל עבודותינו ותפלותינו וכל בקשותינו רק לקצר אפים (זעיר אנפין) ובצירוף פרצופים שלמעלה ממנו כמו שנת' לעיל ובקראינו אליו כאילו קראינו לאלהים עליון. כי האלהים העליון אין לו שם. ואם נקראיהו בשם יהוה לא יענה. כי לא נתיחסו שמות רק לפרצופים הנאצלים ונבראים ממנו. ושם יהוה ביחוד לפי דבריהם הוא האלוה קצר אפים דוקא. ועל האמונות הזרות האלה הרחוקות מדרכי תורתינו הקדושה וקבלת רז"ל קרא ירמיה הנביא ע"ה בדבר ה' לאמר (ירמיהו ה׳:י״ב) כחשו בה' ויאמרו לוא הוא, כחשו בה' הנאמר בו אלהים בצאתך לפני עמך בצעדך בישימון סלה ארץ ועשה אף שמים נטפו מפני אלהים זה סיני מפני אלהים אלהי ישראל, ונאמר ה' מסיני בא וזרח משעיר וכו'. ויאמר לא הוא, אלא הזעיר אנפין הוא שירד על הר סיני ובא להתחבר עם אשתו שכבר קדמה לו והוא ירד אחריה להתחבר ולהזדווג לה, לוא כתיב מלא וי"ו ואלף שחשבו מציאותו יתברך כאלו אינו, וגם קראהו אין, כלומר לא יש, ואמרו שהקורא אליו לא יענהו, כי הוא אין, לפי שידוע שאותיות השלילה והמניעה וההעדר אשר ישתמשו בהם הם לא ואין כמו שכתב הרמב"ם בספר המצות שורש ח', והערביים משתמשים להורות על השלילה וההעדר הנקרא בערבי נפי במלות לם, וליס, ולא, ומא, והמקובלים החדשים קראו לשי"ת אין שהוא לשון שלילה והעדר, ולכן אמרו שהוא יתברך אין לו שום שם ונקודה שהגבילהו כי אין הוא, וצירפו לפעמים גם מלת סוף, ואמרו אין סוף, והיא מליצה הדומה בעניינה לתרגום אין חקר לתבונתו, שתרגם יונתן בן עוזיאל לית סוף לסוכלתנותיה. והשמיטו מלת תבונתו שאמר הנביא ע"ה לבאר לנו שהיש המצוי הזה הוא ה' אלהינו אין חקר וסוף לתבונתו שהמציא כל נמצא בחכמה נפלאה. כמאמר דוד המלך ע"ה מה רבו מעשיך ה' כלם בחכמה עשית. שאינו ניכר אלא מתוך מעשיו. כמאמר הנביא שאו מרום עיניכם וראו מי ברא אלה וכו' וכן אמר דוד כי אראה שמיך מעשה אצבעותיך! וחשבו שכיון שאי אפשר להשיג מהותו והיאך הוא. אם כן אין הוא ח"ו! ולכן יחסו כל השמות הנזכרים בתורה ובנביאים ובלשון רז"ל הקב"ה או רבונו ש"ע הכל על הפרצופים שהמציאו בדמיונם שהם משתלשלים ונמשכים מהאין וקראון גופא דמלכא, ואלו הפרצופים הם עיקר האלהות, כי הם מושגים ומצויירים בשכל, אלא שהם אורות ואינם בעלי בשר ועצמות כמו שכ' הרמב"ם בפתיחת ספר המורה! ונתאוו לאלוהות הרבה כאבותינו אנשי דוד המדבר שאמרו אלה אלהיך ישראל כמאמר רשב"י ע"ה כדגרסינן בפרק ד' מיתות. אמר רב יהודה אלמלא וי"ו שבהעלוך נחייבו שונאיהם של ישראל כלייה. אמר לו רשב"י והלא כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתו לה' לבדו אלא מה ת"ל אשר העלוך מלמד שאיוו לאלהות הרבה, ועיין לעיל סי' נ"ח.

| id | word | punctuation after |
|---|---|---|
| 8:13.1 | קח | ) |
| 8:13.2 | והנה |  |
| 8:13.3 | כבר |  |
| 8:13.4 | ביארנו |  |
| 8:13.5 | כמה |  |
| 8:13.6 | פעמים |  |
| 8:13.7 | שדעת |  |
| 8:13.8 | המקובלים |  |
| 8:13.9 | החדשים |  |
| 8:13.10 | שהזעיר |  |
| 8:13.11 | אנפין |  |
| 8:13.12 | הוא |  |
| 8:13.13 | שנגלה |  |
| 8:13.14 | על |  |
| 8:13.15 | אבותינו |  |
| 8:13.16 | בהר |  |
| 8:13.17 | סיני |  |
| 8:13.18 | ואמר |  |
| 8:13.19 | אנכי |  |
| 8:13.20 | ה | ' |
| 8:13.21 | אלהיך | , |
| 8:13.22 | ואם |  |
| 8:13.23 | כן |  |
| 8:13.24 | איך |  |
| 8:13.25 | הביאו |  |
| 8:13.26 | רז"ל |  |
| 8:13.27 | בשמות |  |
| 8:13.28 | רבה |  |
| 8:13.29 | פסוק |  |
| 8:13.30 | אני |  |
| 8:13.31 | ראשון |  |
| 8:13.32 | ואני |  |
| 8:13.33 | אחרון |  |
| 8:13.34 | וכו | ' |
| 8:13.35 | שאין |  |
| 8:13.36 | לי |  |
| 8:13.37 | אב | , |
| 8:13.38 | ולא |  |
| 8:13.39 | בן |  |
| 8:13.40 | ולא |  |
| 8:13.41 | אח | , |
| 8:13.42 | על |  |
| 8:13.43 | אנכי |  |
| 8:13.44 | ה | ' |
| 8:13.45 | אלהיך |  |
| 8:13.46 | שאמרו |  |
| 8:13.47 | המקובלים |  |
| 8:13.48 | החדשים |  |
| 8:13.49 | שהזעיר |  |
| 8:13.50 | אנפין |  |
| 8:13.51 | אמרו | , |
| 8:13.52 | ומה |  |
| 8:13.53 | לו |  |
| 8:13.54 | לאלהים |  |
| 8:13.55 | העליון |  |
| 8:13.56 | להשתבח |  |
| 8:13.57 | בעצמו |  |
| 8:13.58 | במקום |  |
| 8:13.59 | זה |  |
| 8:13.60 | תוך |  |
| 8:13.61 | דברי |  |
| 8:13.62 | הזעיר |  |
| 8:13.63 | אנפין |  |
| 8:13.64 | שיש |  |
| 8:13.65 | לו |  |
| 8:13.66 | אב |  |
| 8:13.67 | ואם |  |
| 8:13.68 | ובן |  |
| 8:13.69 | ואחות | ? |
| 8:13.70 | אלא |  |
| 8:13.71 | ודאי |  |
| 8:13.72 | שרז"ל |  |
| 8:13.73 | בב"ר |  |
| 8:13.74 | ובשמות |  |
| 8:13.75 | רבה |  |
| 8:13.76 | ודברים |  |
| 8:13.77 | רבה | , |
| 8:13.78 | באו |  |
| 8:13.79 | להעיר |  |
| 8:13.80 | אותנו |  |
| 8:13.81 | שלא |  |
| 8:13.82 | נטעה |  |
| 8:13.83 | באלו |  |
| 8:13.84 | הדעות | , |
| 8:13.85 | ולבאר |  |
| 8:13.86 | לנו | , |
| 8:13.87 | כי |  |
| 8:13.88 | זה |  |
| 8:13.89 | שאמר |  |
| 8:13.90 | אנכי |  |
| 8:13.91 | ה | ' |
| 8:13.92 | אלהיך |  |
| 8:13.93 | הוא |  |
| 8:13.94 | שאמר |  |
| 8:13.95 | אני |  |
| 8:13.96 | ראשון |  |
| 8:13.97 | שאין |  |
| 8:13.98 | לי |  |
| 8:13.99 | אב | , |
| 8:13.100 | ואני |  |
| 8:13.101 | אחרון |  |
| 8:13.102 | שאין |  |
| 8:13.103 | לי |  |
| 8:13.104 | בן | , |
| 8:13.105 | ומבלעדי |  |
| 8:13.106 | אין |  |
| 8:13.107 | אלהים |  |
| 8:13.108 | שאין |  |
| 8:13.109 | לי |  |
| 8:13.110 | אח | , |
| 8:13.111 | הוא |  |
| 8:13.112 | ה | ' |
| 8:13.113 | אלהינו |  |
| 8:13.114 | הסבה |  |
| 8:13.115 | הראשונה |  |
| 8:13.116 | הוא |  |
| 8:13.117 | ולא |  |
| 8:13.118 | אחר | , |
| 8:13.119 | והוא |  |
| 8:13.120 | בכבודו |  |
| 8:13.121 | ובעצמו |  |
| 8:13.122 | מן |  |
| 8:13.123 | השמים |  |
| 8:13.124 | השמיענו |  |
| 8:13.125 | את |  |
| 8:13.126 | קולו | , |
| 8:13.127 | והזהרנו |  |
| 8:13.128 | בדבור |  |
| 8:13.129 | השני |  |
| 8:13.130 | לא |  |
| 8:13.131 | יהיה |  |
| 8:13.132 | לך |  |
| 8:13.133 | אלהים |  |
| 8:13.134 | אחרים |  |
| 8:13.135 | על |  |
| 8:13.136 | פני |  |
| 8:13.137 | וכו | '! |
| 8:13.138 | והוא |  |
| 8:13.139 | שאמר |  |
| 8:13.140 | ראו |  |
| 8:13.141 | עתה |  |
| 8:13.142 | כי |  |
| 8:13.143 | אני |  |
| 8:13.144 | אני |  |
| 8:13.145 | הוא |  |
| 8:13.146 | ואין |  |
| 8:13.147 | אלהים |  |
| 8:13.148 | עמדי | ! |
| 8:13.149 | ודלא |  |
| 8:13.150 | כדעת |  |
| 8:13.151 | הזהר |  |
| 8:13.152 | המסית |  |
| 8:13.153 | דס"ל |  |
| 8:13.154 | שהאלוה |  |
| 8:13.155 | שאמר |  |
| 8:13.156 | אני |  |
| 8:13.157 | ראשון |  |
| 8:13.158 | וכו | ' |
| 8:13.159 | הוא |  |
| 8:13.160 | האין |  |
| 8:13.161 | סוף | , |
| 8:13.162 | ושאמר |  |
| 8:13.163 | ראו |  |
| 8:13.164 | עתה |  |
| 8:13.165 | וכו | ' |
| 8:13.166 | הוא |  |
| 8:13.167 | אדם |  |
| 8:13.168 | קדמון | , |
| 8:13.169 | ושאמר |  |
| 8:13.170 | יהי |  |
| 8:13.171 | אור |  |
| 8:13.172 | וכו | ' |
| 8:13.173 | הוא |  |
| 8:13.174 | אבא |  |
| 8:13.175 | שאמר |  |
| 8:13.176 | לאימא | , |
| 8:13.177 | ושאמר |  |
| 8:13.178 | נעשה |  |
| 8:13.179 | אדם |  |
| 8:13.180 | וכו | ' |
| 8:13.181 | היא |  |
| 8:13.182 | אימא |  |
| 8:13.183 | שאמרה |  |
| 8:13.184 | לאבא |  |
| 8:13.185 | וכו | ' |
| 8:13.186 | כדלעיל | , |
| 8:13.187 | ומכל |  |
| 8:13.188 | מקום |  |
| 8:13.189 | לא |  |
| 8:13.190 | יחסו |  |
| 8:13.191 | כל |  |
| 8:13.192 | עבודותינו |  |
| 8:13.193 | ותפלותינו |  |
| 8:13.194 | וכל |  |
| 8:13.195 | בקשותינו |  |
| 8:13.196 | רק |  |
| 8:13.197 | לקצר |  |
| 8:13.198 | אפים |  |
| 8:13.199 | זעיר |  |
| 8:13.200 | אנפין | ) |
| 8:13.201 | ובצירוף |  |
| 8:13.202 | פרצופים |  |
| 8:13.203 | שלמעלה |  |
| 8:13.204 | ממנו |  |
| 8:13.205 | כמו |  |
| 8:13.206 | שנת | ' |
| 8:13.207 | לעיל |  |
| 8:13.208 | ובקראינו |  |
| 8:13.209 | אליו |  |
| 8:13.210 | כאילו |  |
| 8:13.211 | קראינו |  |
| 8:13.212 | לאלהים |  |
| 8:13.213 | עליון | . |
| 8:13.214 | כי |  |
| 8:13.215 | האלהים |  |
| 8:13.216 | העליון |  |
| 8:13.217 | אין |  |
| 8:13.218 | לו |  |
| 8:13.219 | שם | . |
| 8:13.220 | ואם |  |
| 8:13.221 | נקראיהו |  |
| 8:13.222 | בשם |  |
| 8:13.223 | יהוה |  |
| 8:13.224 | לא |  |
| 8:13.225 | יענה | . |
| 8:13.226 | כי |  |
| 8:13.227 | לא |  |
| 8:13.228 | נתיחסו |  |
| 8:13.229 | שמות |  |
| 8:13.230 | רק |  |
| 8:13.231 | לפרצופים |  |
| 8:13.232 | הנאצלים |  |
| 8:13.233 | ונבראים |  |
| 8:13.234 | ממנו | . |
| 8:13.235 | ושם |  |
| 8:13.236 | יהוה |  |
| 8:13.237 | ביחוד |  |
| 8:13.238 | לפי |  |
| 8:13.239 | דבריהם |  |
| 8:13.240 | הוא |  |
| 8:13.241 | האלוה |  |
| 8:13.242 | קצר |  |
| 8:13.243 | אפים |  |
| 8:13.244 | דוקא | . |
| 8:13.245 | ועל |  |
| 8:13.246 | האמונות |  |
| 8:13.247 | הזרות |  |
| 8:13.248 | האלה |  |
| 8:13.249 | הרחוקות |  |
| 8:13.250 | מדרכי |  |
| 8:13.251 | תורתינו |  |
| 8:13.252 | הקדושה |  |
| 8:13.253 | וקבלת |  |
| 8:13.254 | רז"ל |  |
| 8:13.255 | קרא |  |
| 8:13.256 | ירמיה |  |
| 8:13.257 | הנביא |  |
| 8:13.258 | ע"ה |  |
| 8:13.259 | בדבר |  |
| 8:13.260 | ה | ' |
| 8:13.261 | לאמר |  |
| 8:13.262 | ירמיהו |  |
| 8:13.263 | ה׳:י״ב | ) |
| 8:13.264 | כחשו |  |
| 8:13.265 | בה | ' |
| 8:13.266 | ויאמרו |  |
| 8:13.267 | לוא |  |
| 8:13.268 | הוא | , |
| 8:13.269 | כחשו |  |
| 8:13.270 | בה | ' |
| 8:13.271 | הנאמר |  |
| 8:13.272 | בו |  |
| 8:13.273 | אלהים |  |
| 8:13.274 | בצאתך |  |
| 8:13.275 | לפני |  |
| 8:13.276 | עמך |  |
| 8:13.277 | בצעדך |  |
| 8:13.278 | בישימון |  |
| 8:13.279 | סלה |  |
| 8:13.280 | ארץ |  |
| 8:13.281 | ועשה |  |
| 8:13.282 | אף |  |
| 8:13.283 | שמים |  |
| 8:13.284 | נטפו |  |
| 8:13.285 | מפני |  |
| 8:13.286 | אלהים |  |
| 8:13.287 | זה |  |
| 8:13.288 | סיני |  |
| 8:13.289 | מפני |  |
| 8:13.290 | אלהים |  |
| 8:13.291 | אלהי |  |
| 8:13.292 | ישראל | , |
| 8:13.293 | ונאמר |  |
| 8:13.294 | ה | ' |
| 8:13.295 | מסיני |  |
| 8:13.296 | בא |  |
| 8:13.297 | וזרח |  |
| 8:13.298 | משעיר |  |
| 8:13.299 | וכו | '. |
| 8:13.300 | ויאמר |  |
| 8:13.301 | לא |  |
| 8:13.302 | הוא | , |
| 8:13.303 | אלא |  |
| 8:13.304 | הזעיר |  |
| 8:13.305 | אנפין |  |
| 8:13.306 | הוא |  |
| 8:13.307 | שירד |  |
| 8:13.308 | על |  |
| 8:13.309 | הר |  |
| 8:13.310 | סיני |  |
| 8:13.311 | ובא |  |
| 8:13.312 | להתחבר |  |
| 8:13.313 | עם |  |
| 8:13.314 | אשתו |  |
| 8:13.315 | שכבר |  |
| 8:13.316 | קדמה |  |
| 8:13.317 | לו |  |
| 8:13.318 | והוא |  |
| 8:13.319 | ירד |  |
| 8:13.320 | אחריה |  |
| 8:13.321 | להתחבר |  |
| 8:13.322 | ולהזדווג |  |
| 8:13.323 | לה | , |
| 8:13.324 | לוא |  |
| 8:13.325 | כתיב |  |
| 8:13.326 | מלא |  |
| 8:13.327 | וי"ו |  |
| 8:13.328 | ואלף |  |
| 8:13.329 | שחשבו |  |
| 8:13.330 | מציאותו |  |
| 8:13.331 | יתברך |  |
| 8:13.332 | כאלו |  |
| 8:13.333 | אינו | , |
| 8:13.334 | וגם |  |
| 8:13.335 | קראהו |  |
| 8:13.336 | אין | , |
| 8:13.337 | כלומר |  |
| 8:13.338 | לא |  |
| 8:13.339 | יש | , |
| 8:13.340 | ואמרו |  |
| 8:13.341 | שהקורא |  |
| 8:13.342 | אליו |  |
| 8:13.343 | לא |  |
| 8:13.344 | יענהו | , |
| 8:13.345 | כי |  |
| 8:13.346 | הוא |  |
| 8:13.347 | אין | , |
| 8:13.348 | לפי |  |
| 8:13.349 | שידוע |  |
| 8:13.350 | שאותיות |  |
| 8:13.351 | השלילה |  |
| 8:13.352 | והמניעה |  |
| 8:13.353 | וההעדר |  |
| 8:13.354 | אשר |  |
| 8:13.355 | ישתמשו |  |
| 8:13.356 | בהם |  |
| 8:13.357 | הם |  |
| 8:13.358 | לא |  |
| 8:13.359 | ואין |  |
| 8:13.360 | כמו |  |
| 8:13.361 | שכתב |  |
| 8:13.362 | הרמב"ם |  |
| 8:13.363 | בספר |  |
| 8:13.364 | המצות |  |
| 8:13.365 | שורש |  |
| 8:13.366 | ח | ', |
| 8:13.367 | והערביים |  |
| 8:13.368 | משתמשים |  |
| 8:13.369 | להורות |  |
| 8:13.370 | על |  |
| 8:13.371 | השלילה |  |
| 8:13.372 | וההעדר |  |
| 8:13.373 | הנקרא |  |
| 8:13.374 | בערבי |  |
| 8:13.375 | נפי |  |
| 8:13.376 | במלות |  |
| 8:13.377 | לם | , |
| 8:13.378 | וליס | , |
| 8:13.379 | ולא | , |
| 8:13.380 | ומא | , |
| 8:13.381 | והמקובלים |  |
| 8:13.382 | החדשים |  |
| 8:13.383 | קראו |  |
| 8:13.384 | לשי"ת |  |
| 8:13.385 | אין |  |
| 8:13.386 | שהוא |  |
| 8:13.387 | לשון |  |
| 8:13.388 | שלילה |  |
| 8:13.389 | והעדר | , |
| 8:13.390 | ולכן |  |
| 8:13.391 | אמרו |  |
| 8:13.392 | שהוא |  |
| 8:13.393 | יתברך |  |
| 8:13.394 | אין |  |
| 8:13.395 | לו |  |
| 8:13.396 | שום |  |
| 8:13.397 | שם |  |
| 8:13.398 | ונקודה |  |
| 8:13.399 | שהגבילהו |  |
| 8:13.400 | כי |  |
| 8:13.401 | אין |  |
| 8:13.402 | הוא | , |
| 8:13.403 | וצירפו |  |
| 8:13.404 | לפעמים |  |
| 8:13.405 | גם |  |
| 8:13.406 | מלת |  |
| 8:13.407 | סוף | , |
| 8:13.408 | ואמרו |  |
| 8:13.409 | אין |  |
| 8:13.410 | סוף | , |
| 8:13.411 | והיא |  |
| 8:13.412 | מליצה |  |
| 8:13.413 | הדומה |  |
| 8:13.414 | בעניינה |  |
| 8:13.415 | לתרגום |  |
| 8:13.416 | אין |  |
| 8:13.417 | חקר |  |
| 8:13.418 | לתבונתו | , |
| 8:13.419 | שתרגם |  |
| 8:13.420 | יונתן |  |
| 8:13.421 | בן |  |
| 8:13.422 | עוזיאל |  |
| 8:13.423 | לית |  |
| 8:13.424 | סוף |  |
| 8:13.425 | לסוכלתנותיה | . |
| 8:13.426 | והשמיטו |  |
| 8:13.427 | מלת |  |
| 8:13.428 | תבונתו |  |
| 8:13.429 | שאמר |  |
| 8:13.430 | הנביא |  |
| 8:13.431 | ע"ה |  |
| 8:13.432 | לבאר |  |
| 8:13.433 | לנו |  |
| 8:13.434 | שהיש |  |
| 8:13.435 | המצוי |  |
| 8:13.436 | הזה |  |
| 8:13.437 | הוא |  |
| 8:13.438 | ה | ' |
| 8:13.439 | אלהינו |  |
| 8:13.440 | אין |  |
| 8:13.441 | חקר |  |
| 8:13.442 | וסוף |  |
| 8:13.443 | לתבונתו |  |
| 8:13.444 | שהמציא |  |
| 8:13.445 | כל |  |
| 8:13.446 | נמצא |  |
| 8:13.447 | בחכמה |  |
| 8:13.448 | נפלאה | . |
| 8:13.449 | כמאמר |  |
| 8:13.450 | דוד |  |
| 8:13.451 | המלך |  |
| 8:13.452 | ע"ה |  |
| 8:13.453 | מה |  |
| 8:13.454 | רבו |  |
| 8:13.455 | מעשיך |  |
| 8:13.456 | ה | ' |
| 8:13.457 | כלם |  |
| 8:13.458 | בחכמה |  |
| 8:13.459 | עשית | . |
| 8:13.460 | שאינו |  |
| 8:13.461 | ניכר |  |
| 8:13.462 | אלא |  |
| 8:13.463 | מתוך |  |
| 8:13.464 | מעשיו | . |
| 8:13.465 | כמאמר |  |
| 8:13.466 | הנביא |  |
| 8:13.467 | שאו |  |
| 8:13.468 | מרום |  |
| 8:13.469 | עיניכם |  |
| 8:13.470 | וראו |  |
| 8:13.471 | מי |  |
| 8:13.472 | ברא |  |
| 8:13.473 | אלה |  |
| 8:13.474 | וכו | ' |
| 8:13.475 | וכן |  |
| 8:13.476 | אמר |  |
| 8:13.477 | דוד |  |
| 8:13.478 | כי |  |
| 8:13.479 | אראה |  |
| 8:13.480 | שמיך |  |
| 8:13.481 | מעשה |  |
| 8:13.482 | אצבעותיך | ! |
| 8:13.483 | וחשבו |  |
| 8:13.484 | שכיון |  |
| 8:13.485 | שאי |  |
| 8:13.486 | אפשר |  |
| 8:13.487 | להשיג |  |
| 8:13.488 | מהותו |  |
| 8:13.489 | והיאך |  |
| 8:13.490 | הוא | . |
| 8:13.491 | אם |  |
| 8:13.492 | כן |  |
| 8:13.493 | אין |  |
| 8:13.494 | הוא |  |
| 8:13.495 | ח"ו | ! |
| 8:13.496 | ולכן |  |
| 8:13.497 | יחסו |  |
| 8:13.498 | כל |  |
| 8:13.499 | השמות |  |
| 8:13.500 | הנזכרים |  |
| 8:13.501 | בתורה |  |
| 8:13.502 | ובנביאים |  |
| 8:13.503 | ובלשון |  |
| 8:13.504 | רז"ל |  |
| 8:13.505 | הקב"ה |  |
| 8:13.506 | או |  |
| 8:13.507 | רבונו |  |
| 8:13.508 | ש"ע |  |
| 8:13.509 | הכל |  |
| 8:13.510 | על |  |
| 8:13.511 | הפרצופים |  |
| 8:13.512 | שהמציאו |  |
| 8:13.513 | בדמיונם |  |
| 8:13.514 | שהם |  |
| 8:13.515 | משתלשלים |  |
| 8:13.516 | ונמשכים |  |
| 8:13.517 | מהאין |  |
| 8:13.518 | וקראון |  |
| 8:13.519 | גופא |  |
| 8:13.520 | דמלכא | , |
| 8:13.521 | ואלו |  |
| 8:13.522 | הפרצופים |  |
| 8:13.523 | הם |  |
| 8:13.524 | עיקר |  |
| 8:13.525 | האלהות | , |
| 8:13.526 | כי |  |
| 8:13.527 | הם |  |
| 8:13.528 | מושגים |  |
| 8:13.529 | ומצויירים |  |
| 8:13.530 | בשכל | , |
| 8:13.531 | אלא |  |
| 8:13.532 | שהם |  |
| 8:13.533 | אורות |  |
| 8:13.534 | ואינם |  |
| 8:13.535 | בעלי |  |
| 8:13.536 | בשר |  |
| 8:13.537 | ועצמות |  |
| 8:13.538 | כמו |  |
| 8:13.539 | שכ | ' |
| 8:13.540 | הרמב"ם |  |
| 8:13.541 | בפתיחת |  |
| 8:13.542 | ספר |  |
| 8:13.543 | המורה | ! |
| 8:13.544 | ונתאוו |  |
| 8:13.545 | לאלוהות |  |
| 8:13.546 | הרבה |  |
| 8:13.547 | כאבותינו |  |
| 8:13.548 | אנשי |  |
| 8:13.549 | דוד |  |
| 8:13.550 | המדבר |  |
| 8:13.551 | שאמרו |  |
| 8:13.552 | אלה |  |
| 8:13.553 | אלהיך |  |
| 8:13.554 | ישראל |  |
| 8:13.555 | כמאמר |  |
| 8:13.556 | רשב"י |  |
| 8:13.557 | ע"ה |  |
| 8:13.558 | כדגרסינן |  |
| 8:13.559 | בפרק |  |
| 8:13.560 | ד | ' |
| 8:13.561 | מיתות | . |
| 8:13.562 | אמר |  |
| 8:13.563 | רב |  |
| 8:13.564 | יהודה |  |
| 8:13.565 | אלמלא |  |
| 8:13.566 | וי"ו |  |
| 8:13.567 | שבהעלוך |  |
| 8:13.568 | נחייבו |  |
| 8:13.569 | שונאיהם |  |
| 8:13.570 | של |  |
| 8:13.571 | ישראל |  |
| 8:13.572 | כלייה | . |
| 8:13.573 | אמר |  |
| 8:13.574 | לו |  |
| 8:13.575 | רשב"י |  |
| 8:13.576 | והלא |  |
| 8:13.577 | כל |  |
| 8:13.578 | המשתף |  |
| 8:13.579 | שם |  |
| 8:13.580 | שמים |  |
| 8:13.581 | ודבר |  |
| 8:13.582 | אחר |  |
| 8:13.583 | נעקר |  |
| 8:13.584 | מן |  |
| 8:13.585 | העולם |  |
| 8:13.586 | שנ | ' |
| 8:13.587 | בלתו |  |
| 8:13.588 | לה | ' |
| 8:13.589 | לבדו |  |
| 8:13.590 | אלא |  |
| 8:13.591 | מה |  |
| 8:13.592 | ת"ל |  |
| 8:13.593 | אשר |  |
| 8:13.594 | העלוך |  |
| 8:13.595 | מלמד |  |
| 8:13.596 | שאיוו |  |
| 8:13.597 | לאלהות |  |
| 8:13.598 | הרבה | , |
| 8:13.599 | ועיין |  |
| 8:13.600 | לעיל |  |
| 8:13.601 | סי | ' |
| 8:13.602 | נ"ח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

