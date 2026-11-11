# Translation pack: milhamot-hashem perek-6-06 (Milchamot ha-Shem 6:11–6:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-06`.

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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
- כלא חד / כולא חד: the same English at every occurrence (1:5-1:7); the
  owner's ruling on the sense is pending; until then, the alternative is in
  `apparatus`.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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


## Milchamot ha-Shem 6:11

Hebrew (Wikisource): סה) עוד כ' בעל הגדרים בשורש יחוד וז"ל יחוד נגזר מן יחיד שמשו בו להורות על האחדות האמתי אשר אין בו הרכבה כלל. אשר לא יבינהו רק החכם הרגיל בשלילת הגשם ומשיגיו. ואיש בער לא ידע שלילה. וכסיל לא יבינהו רק החכם והנה האחדות האמתי היא מצוה כתובה בתורה. שמע ישראל ה' אלהינו ה' אחד. ומהמבואר שאין לפסוק שום מצוה מן התורה שבכתב לבד כי אם מן התורה שבעל פה שהיא ביאור המצות! והנה מענין מצות האחדות לא מצאנו שדברי חכמים ז"ל בביאורה. ובלי ספק שהיו מוסרין ביאורו איש לאיש בקבלה. כי לא יתכן שיעזבו ראשית המצוה ויסוד האמונה האלהית בלי ביאור! אך עתה שהביאור ההוא נשכח מאתנו עלינו לקחת תורה מפי הפילסופים (התורניים) למען לא נתעה בלי דעת להאמין בשניות ובגשמות (כהמקובלים החדשים) ונעבוד ללא אלהי אמת. והוא הדבר שאמר בעל החובות ז"ל לא יוכל לעבוד עלת העלות אלא נביא הדור או החכם המובהק במה שקנהו מן החכמה. וזולתם עובדים זולתו. ובאמת דבר החכם כי כמה דברים מורכבים שנים יתוארו בשם אחד. על דרך ההסכמה ע"כ. ביאור דבריו כגון שם כירים שהוא מקום שפיתת שתי קדרות, ונקרא בדברי רבותינו בשם אחד כירה. וכן רחים שהוא שם השכב העליון רכב, ונקראו שניהם בשם אחד טחון. בחורים טחון נשאו ובערבי רחא אע"פ שהם זוג. וכן מספרים נקרא בערבי מקץ. וכן כל כיוצא בהם כדתנן בפרק י"ג ממסכת כלים קלגריפון, ומכחול, ומכתב, וזומא ליסטרא, אע"פ שהם שנים מחוברים ומורכבים נקראים בשם אחד על דרך ההסכמה וההעברה. ולא כן היא אחדותו ית' שהוא אחד אמיתי שאין בו שום הרכבה וצירוף, שלא כדעת המקובלים החדשים שאומרים שיש לו כמה פרצופים שנאצלו ממנו. וכל אחד נקרא בשם בפני עצמו, והם מחברים ומשתפים אותם בפיהם ובדעתם עם האלוה המאצילם ואומרים וכלא חד. ועובדים אותם יחד גם הגוף גם הנשמה, ומשקרים לומר כי את נשמתו הם עובדים, וזה לשון עוז לאלהים בבית קדש הקדשים פ' י"ח דף נ"ו ע"ב ואם תאמר שאהיה עובד לנשמה בלא גופא ח"ו, חלילה, אלא כבר אמרנו לעיל שזה הפרצוף דמלכא קדישא ושכינתיה הוא אדוק בנשמה. לפי שהוא עתיקא קדישא דכל קדישין. והוא ונשמתא חד אינון כמא דאמרינן לעיל הוא ושמיה חד אינון ולא מתפרש מנהון לעלמין. ואם יעבוד לנשמה בלי מלכא קדישא יהיה פירוד בין נשמה לגופא. וכן מבואר מדברי גליון הזהר ומשנת חסידים הובא לעיל סי' ט"ל. ועל מה שכ' בעל הגדרים שלא יתכן שיעזבו רז"ל ראשית המצוה ויסודה וכו' כבר העירותי על דבריו לעיל סי' ס"ד שאדרבה נהפוך הוא שכל דבר המפורש בכתובים והוא פשוט ומפורסם בישראל לא ידברו בו חכמי המשנה. ובעלי התלמוד ג"כ לא שמו מגמתם רק לבאר ולישא וליתן בדבריה ולבררה וללבנה, ולהודיע דינים ושאלות שנתחדשו בימיהם, וגזרות ותקנות שתיקנו, אבל דברים הפשוטים אין צורך להזכירם כלל. ודעות ורעיונות הקבלה החדשה עוד לא נודעו ולא נשמעו בימיהם בין היהודים כלל, וזרים המה לאמונת תורתינו הקדושה. כי רק מתחלת האלף הששי בסוף המאה הראשונה החזיקו ברעיונות האלה קצת מחברים שנמשכו אחר דעות הפילסוף הרמאי שדבק בדעת השונים והשלישי הוא מחבר ספר הזהר והתקונים. וכוון בו להטעות את ישראל ולהדיחם מעל ה' אלהינו. ותלה רעיונותיו ברשב"י התנא וחביריו, בפרט בשני דברים, מידי דעל רשב"י תני במאמרו הקצר כל המשתף שם שמים ודבר אחר נעקר מן העולם. והשני בביזוי המשנה שקלא וטריא דבתלמודא שהם עיקר ויסוד תורה שבעל פה שהוא מבזה אותן כמו שיתבאר להלן (דלרשב"י שתורתו אומנותו היא כקריאת שמע בעונתה ירוש' שבת פ"א). וטמן את הספר הנז' בארון וקברו בקרקע כדי לישנו כאלו הוא ישן נושן מימי רשב"י. ואמר למלך ממלכי מזרח לחפור שם בקרקע למצוא מטמון של כסף וזהב, וצוה המלך הנז' לחפור על עסקי ממון ונמצא שם ארון ובו ס' הזהר וכו' ושלח אותו לטוליטולא וכו' ומשם נתפשטה הקבלה החדשה בישראל כמו שכ' הר' חיד"א בספרו הנחמד שם הגדולים במערכת ספרים ערך זהר.

| id | word | punctuation after |
|---|---|---|
| 6:11.1 | סה | ) |
| 6:11.2 | עוד |  |
| 6:11.3 | כ | ' |
| 6:11.4 | בעל |  |
| 6:11.5 | הגדרים |  |
| 6:11.6 | בשורש |  |
| 6:11.7 | יחוד |  |
| 6:11.8 | וז"ל |  |
| 6:11.9 | יחוד |  |
| 6:11.10 | נגזר |  |
| 6:11.11 | מן |  |
| 6:11.12 | יחיד |  |
| 6:11.13 | שמשו |  |
| 6:11.14 | בו |  |
| 6:11.15 | להורות |  |
| 6:11.16 | על |  |
| 6:11.17 | האחדות |  |
| 6:11.18 | האמתי |  |
| 6:11.19 | אשר |  |
| 6:11.20 | אין |  |
| 6:11.21 | בו |  |
| 6:11.22 | הרכבה |  |
| 6:11.23 | כלל | . |
| 6:11.24 | אשר |  |
| 6:11.25 | לא |  |
| 6:11.26 | יבינהו |  |
| 6:11.27 | רק |  |
| 6:11.28 | החכם |  |
| 6:11.29 | הרגיל |  |
| 6:11.30 | בשלילת |  |
| 6:11.31 | הגשם |  |
| 6:11.32 | ומשיגיו | . |
| 6:11.33 | ואיש |  |
| 6:11.34 | בער |  |
| 6:11.35 | לא |  |
| 6:11.36 | ידע |  |
| 6:11.37 | שלילה | . |
| 6:11.38 | וכסיל |  |
| 6:11.39 | לא |  |
| 6:11.40 | יבינהו |  |
| 6:11.41 | רק |  |
| 6:11.42 | החכם |  |
| 6:11.43 | והנה |  |
| 6:11.44 | האחדות |  |
| 6:11.45 | האמתי |  |
| 6:11.46 | היא |  |
| 6:11.47 | מצוה |  |
| 6:11.48 | כתובה |  |
| 6:11.49 | בתורה | . |
| 6:11.50 | שמע |  |
| 6:11.51 | ישראל |  |
| 6:11.52 | ה | ' |
| 6:11.53 | אלהינו |  |
| 6:11.54 | ה | ' |
| 6:11.55 | אחד | . |
| 6:11.56 | ומהמבואר |  |
| 6:11.57 | שאין |  |
| 6:11.58 | לפסוק |  |
| 6:11.59 | שום |  |
| 6:11.60 | מצוה |  |
| 6:11.61 | מן |  |
| 6:11.62 | התורה |  |
| 6:11.63 | שבכתב |  |
| 6:11.64 | לבד |  |
| 6:11.65 | כי |  |
| 6:11.66 | אם |  |
| 6:11.67 | מן |  |
| 6:11.68 | התורה |  |
| 6:11.69 | שבעל |  |
| 6:11.70 | פה |  |
| 6:11.71 | שהיא |  |
| 6:11.72 | ביאור |  |
| 6:11.73 | המצות | ! |
| 6:11.74 | והנה |  |
| 6:11.75 | מענין |  |
| 6:11.76 | מצות |  |
| 6:11.77 | האחדות |  |
| 6:11.78 | לא |  |
| 6:11.79 | מצאנו |  |
| 6:11.80 | שדברי |  |
| 6:11.81 | חכמים |  |
| 6:11.82 | ז"ל |  |
| 6:11.83 | בביאורה | . |
| 6:11.84 | ובלי |  |
| 6:11.85 | ספק |  |
| 6:11.86 | שהיו |  |
| 6:11.87 | מוסרין |  |
| 6:11.88 | ביאורו |  |
| 6:11.89 | איש |  |
| 6:11.90 | לאיש |  |
| 6:11.91 | בקבלה | . |
| 6:11.92 | כי |  |
| 6:11.93 | לא |  |
| 6:11.94 | יתכן |  |
| 6:11.95 | שיעזבו |  |
| 6:11.96 | ראשית |  |
| 6:11.97 | המצוה |  |
| 6:11.98 | ויסוד |  |
| 6:11.99 | האמונה |  |
| 6:11.100 | האלהית |  |
| 6:11.101 | בלי |  |
| 6:11.102 | ביאור | ! |
| 6:11.103 | אך |  |
| 6:11.104 | עתה |  |
| 6:11.105 | שהביאור |  |
| 6:11.106 | ההוא |  |
| 6:11.107 | נשכח |  |
| 6:11.108 | מאתנו |  |
| 6:11.109 | עלינו |  |
| 6:11.110 | לקחת |  |
| 6:11.111 | תורה |  |
| 6:11.112 | מפי |  |
| 6:11.113 | הפילסופים |  |
| 6:11.114 | התורניים | ) |
| 6:11.115 | למען |  |
| 6:11.116 | לא |  |
| 6:11.117 | נתעה |  |
| 6:11.118 | בלי |  |
| 6:11.119 | דעת |  |
| 6:11.120 | להאמין |  |
| 6:11.121 | בשניות |  |
| 6:11.122 | ובגשמות |  |
| 6:11.123 | כהמקובלים |  |
| 6:11.124 | החדשים | ) |
| 6:11.125 | ונעבוד |  |
| 6:11.126 | ללא |  |
| 6:11.127 | אלהי |  |
| 6:11.128 | אמת | . |
| 6:11.129 | והוא |  |
| 6:11.130 | הדבר |  |
| 6:11.131 | שאמר |  |
| 6:11.132 | בעל |  |
| 6:11.133 | החובות |  |
| 6:11.134 | ז"ל |  |
| 6:11.135 | לא |  |
| 6:11.136 | יוכל |  |
| 6:11.137 | לעבוד |  |
| 6:11.138 | עלת |  |
| 6:11.139 | העלות |  |
| 6:11.140 | אלא |  |
| 6:11.141 | נביא |  |
| 6:11.142 | הדור |  |
| 6:11.143 | או |  |
| 6:11.144 | החכם |  |
| 6:11.145 | המובהק |  |
| 6:11.146 | במה |  |
| 6:11.147 | שקנהו |  |
| 6:11.148 | מן |  |
| 6:11.149 | החכמה | . |
| 6:11.150 | וזולתם |  |
| 6:11.151 | עובדים |  |
| 6:11.152 | זולתו | . |
| 6:11.153 | ובאמת |  |
| 6:11.154 | דבר |  |
| 6:11.155 | החכם |  |
| 6:11.156 | כי |  |
| 6:11.157 | כמה |  |
| 6:11.158 | דברים |  |
| 6:11.159 | מורכבים |  |
| 6:11.160 | שנים |  |
| 6:11.161 | יתוארו |  |
| 6:11.162 | בשם |  |
| 6:11.163 | אחד | . |
| 6:11.164 | על |  |
| 6:11.165 | דרך |  |
| 6:11.166 | ההסכמה |  |
| 6:11.167 | ע"כ | . |
| 6:11.168 | ביאור |  |
| 6:11.169 | דבריו |  |
| 6:11.170 | כגון |  |
| 6:11.171 | שם |  |
| 6:11.172 | כירים |  |
| 6:11.173 | שהוא |  |
| 6:11.174 | מקום |  |
| 6:11.175 | שפיתת |  |
| 6:11.176 | שתי |  |
| 6:11.177 | קדרות | , |
| 6:11.178 | ונקרא |  |
| 6:11.179 | בדברי |  |
| 6:11.180 | רבותינו |  |
| 6:11.181 | בשם |  |
| 6:11.182 | אחד |  |
| 6:11.183 | כירה | . |
| 6:11.184 | וכן |  |
| 6:11.185 | רחים |  |
| 6:11.186 | שהוא |  |
| 6:11.187 | שם |  |
| 6:11.188 | השכב |  |
| 6:11.189 | העליון |  |
| 6:11.190 | רכב | , |
| 6:11.191 | ונקראו |  |
| 6:11.192 | שניהם |  |
| 6:11.193 | בשם |  |
| 6:11.194 | אחד |  |
| 6:11.195 | טחון | . |
| 6:11.196 | בחורים |  |
| 6:11.197 | טחון |  |
| 6:11.198 | נשאו |  |
| 6:11.199 | ובערבי |  |
| 6:11.200 | רחא |  |
| 6:11.201 | אע"פ |  |
| 6:11.202 | שהם |  |
| 6:11.203 | זוג | . |
| 6:11.204 | וכן |  |
| 6:11.205 | מספרים |  |
| 6:11.206 | נקרא |  |
| 6:11.207 | בערבי |  |
| 6:11.208 | מקץ | . |
| 6:11.209 | וכן |  |
| 6:11.210 | כל |  |
| 6:11.211 | כיוצא |  |
| 6:11.212 | בהם |  |
| 6:11.213 | כדתנן |  |
| 6:11.214 | בפרק |  |
| 6:11.215 | י"ג |  |
| 6:11.216 | ממסכת |  |
| 6:11.217 | כלים |  |
| 6:11.218 | קלגריפון | , |
| 6:11.219 | ומכחול | , |
| 6:11.220 | ומכתב | , |
| 6:11.221 | וזומא |  |
| 6:11.222 | ליסטרא | , |
| 6:11.223 | אע"פ |  |
| 6:11.224 | שהם |  |
| 6:11.225 | שנים |  |
| 6:11.226 | מחוברים |  |
| 6:11.227 | ומורכבים |  |
| 6:11.228 | נקראים |  |
| 6:11.229 | בשם |  |
| 6:11.230 | אחד |  |
| 6:11.231 | על |  |
| 6:11.232 | דרך |  |
| 6:11.233 | ההסכמה |  |
| 6:11.234 | וההעברה | . |
| 6:11.235 | ולא |  |
| 6:11.236 | כן |  |
| 6:11.237 | היא |  |
| 6:11.238 | אחדותו |  |
| 6:11.239 | ית | ' |
| 6:11.240 | שהוא |  |
| 6:11.241 | אחד |  |
| 6:11.242 | אמיתי |  |
| 6:11.243 | שאין |  |
| 6:11.244 | בו |  |
| 6:11.245 | שום |  |
| 6:11.246 | הרכבה |  |
| 6:11.247 | וצירוף | , |
| 6:11.248 | שלא |  |
| 6:11.249 | כדעת |  |
| 6:11.250 | המקובלים |  |
| 6:11.251 | החדשים |  |
| 6:11.252 | שאומרים |  |
| 6:11.253 | שיש |  |
| 6:11.254 | לו |  |
| 6:11.255 | כמה |  |
| 6:11.256 | פרצופים |  |
| 6:11.257 | שנאצלו |  |
| 6:11.258 | ממנו | . |
| 6:11.259 | וכל |  |
| 6:11.260 | אחד |  |
| 6:11.261 | נקרא |  |
| 6:11.262 | בשם |  |
| 6:11.263 | בפני |  |
| 6:11.264 | עצמו | , |
| 6:11.265 | והם |  |
| 6:11.266 | מחברים |  |
| 6:11.267 | ומשתפים |  |
| 6:11.268 | אותם |  |
| 6:11.269 | בפיהם |  |
| 6:11.270 | ובדעתם |  |
| 6:11.271 | עם |  |
| 6:11.272 | האלוה |  |
| 6:11.273 | המאצילם |  |
| 6:11.274 | ואומרים |  |
| 6:11.275 | וכלא |  |
| 6:11.276 | חד | . |
| 6:11.277 | ועובדים |  |
| 6:11.278 | אותם |  |
| 6:11.279 | יחד |  |
| 6:11.280 | גם |  |
| 6:11.281 | הגוף |  |
| 6:11.282 | גם |  |
| 6:11.283 | הנשמה | , |
| 6:11.284 | ומשקרים |  |
| 6:11.285 | לומר |  |
| 6:11.286 | כי |  |
| 6:11.287 | את |  |
| 6:11.288 | נשמתו |  |
| 6:11.289 | הם |  |
| 6:11.290 | עובדים | , |
| 6:11.291 | וזה |  |
| 6:11.292 | לשון |  |
| 6:11.293 | עוז |  |
| 6:11.294 | לאלהים |  |
| 6:11.295 | בבית |  |
| 6:11.296 | קדש |  |
| 6:11.297 | הקדשים |  |
| 6:11.298 | פ | ' |
| 6:11.299 | י"ח |  |
| 6:11.300 | דף |  |
| 6:11.301 | נ"ו |  |
| 6:11.302 | ע"ב |  |
| 6:11.303 | ואם |  |
| 6:11.304 | תאמר |  |
| 6:11.305 | שאהיה |  |
| 6:11.306 | עובד |  |
| 6:11.307 | לנשמה |  |
| 6:11.308 | בלא |  |
| 6:11.309 | גופא |  |
| 6:11.310 | ח"ו | , |
| 6:11.311 | חלילה | , |
| 6:11.312 | אלא |  |
| 6:11.313 | כבר |  |
| 6:11.314 | אמרנו |  |
| 6:11.315 | לעיל |  |
| 6:11.316 | שזה |  |
| 6:11.317 | הפרצוף |  |
| 6:11.318 | דמלכא |  |
| 6:11.319 | קדישא |  |
| 6:11.320 | ושכינתיה |  |
| 6:11.321 | הוא |  |
| 6:11.322 | אדוק |  |
| 6:11.323 | בנשמה | . |
| 6:11.324 | לפי |  |
| 6:11.325 | שהוא |  |
| 6:11.326 | עתיקא |  |
| 6:11.327 | קדישא |  |
| 6:11.328 | דכל |  |
| 6:11.329 | קדישין | . |
| 6:11.330 | והוא |  |
| 6:11.331 | ונשמתא |  |
| 6:11.332 | חד |  |
| 6:11.333 | אינון |  |
| 6:11.334 | כמא |  |
| 6:11.335 | דאמרינן |  |
| 6:11.336 | לעיל |  |
| 6:11.337 | הוא |  |
| 6:11.338 | ושמיה |  |
| 6:11.339 | חד |  |
| 6:11.340 | אינון |  |
| 6:11.341 | ולא |  |
| 6:11.342 | מתפרש |  |
| 6:11.343 | מנהון |  |
| 6:11.344 | לעלמין | . |
| 6:11.345 | ואם |  |
| 6:11.346 | יעבוד |  |
| 6:11.347 | לנשמה |  |
| 6:11.348 | בלי |  |
| 6:11.349 | מלכא |  |
| 6:11.350 | קדישא |  |
| 6:11.351 | יהיה |  |
| 6:11.352 | פירוד |  |
| 6:11.353 | בין |  |
| 6:11.354 | נשמה |  |
| 6:11.355 | לגופא | . |
| 6:11.356 | וכן |  |
| 6:11.357 | מבואר |  |
| 6:11.358 | מדברי |  |
| 6:11.359 | גליון |  |
| 6:11.360 | הזהר |  |
| 6:11.361 | ומשנת |  |
| 6:11.362 | חסידים |  |
| 6:11.363 | הובא |  |
| 6:11.364 | לעיל |  |
| 6:11.365 | סי | ' |
| 6:11.366 | ט"ל | . |
| 6:11.367 | ועל |  |
| 6:11.368 | מה |  |
| 6:11.369 | שכ | ' |
| 6:11.370 | בעל |  |
| 6:11.371 | הגדרים |  |
| 6:11.372 | שלא |  |
| 6:11.373 | יתכן |  |
| 6:11.374 | שיעזבו |  |
| 6:11.375 | רז"ל |  |
| 6:11.376 | ראשית |  |
| 6:11.377 | המצוה |  |
| 6:11.378 | ויסודה |  |
| 6:11.379 | וכו | ' |
| 6:11.380 | כבר |  |
| 6:11.381 | העירותי |  |
| 6:11.382 | על |  |
| 6:11.383 | דבריו |  |
| 6:11.384 | לעיל |  |
| 6:11.385 | סי | ' |
| 6:11.386 | ס"ד |  |
| 6:11.387 | שאדרבה |  |
| 6:11.388 | נהפוך |  |
| 6:11.389 | הוא |  |
| 6:11.390 | שכל |  |
| 6:11.391 | דבר |  |
| 6:11.392 | המפורש |  |
| 6:11.393 | בכתובים |  |
| 6:11.394 | והוא |  |
| 6:11.395 | פשוט |  |
| 6:11.396 | ומפורסם |  |
| 6:11.397 | בישראל |  |
| 6:11.398 | לא |  |
| 6:11.399 | ידברו |  |
| 6:11.400 | בו |  |
| 6:11.401 | חכמי |  |
| 6:11.402 | המשנה | . |
| 6:11.403 | ובעלי |  |
| 6:11.404 | התלמוד |  |
| 6:11.405 | ג"כ |  |
| 6:11.406 | לא |  |
| 6:11.407 | שמו |  |
| 6:11.408 | מגמתם |  |
| 6:11.409 | רק |  |
| 6:11.410 | לבאר |  |
| 6:11.411 | ולישא |  |
| 6:11.412 | וליתן |  |
| 6:11.413 | בדבריה |  |
| 6:11.414 | ולבררה |  |
| 6:11.415 | וללבנה | , |
| 6:11.416 | ולהודיע |  |
| 6:11.417 | דינים |  |
| 6:11.418 | ושאלות |  |
| 6:11.419 | שנתחדשו |  |
| 6:11.420 | בימיהם | , |
| 6:11.421 | וגזרות |  |
| 6:11.422 | ותקנות |  |
| 6:11.423 | שתיקנו | , |
| 6:11.424 | אבל |  |
| 6:11.425 | דברים |  |
| 6:11.426 | הפשוטים |  |
| 6:11.427 | אין |  |
| 6:11.428 | צורך |  |
| 6:11.429 | להזכירם |  |
| 6:11.430 | כלל | . |
| 6:11.431 | ודעות |  |
| 6:11.432 | ורעיונות |  |
| 6:11.433 | הקבלה |  |
| 6:11.434 | החדשה |  |
| 6:11.435 | עוד |  |
| 6:11.436 | לא |  |
| 6:11.437 | נודעו |  |
| 6:11.438 | ולא |  |
| 6:11.439 | נשמעו |  |
| 6:11.440 | בימיהם |  |
| 6:11.441 | בין |  |
| 6:11.442 | היהודים |  |
| 6:11.443 | כלל | , |
| 6:11.444 | וזרים |  |
| 6:11.445 | המה |  |
| 6:11.446 | לאמונת |  |
| 6:11.447 | תורתינו |  |
| 6:11.448 | הקדושה | . |
| 6:11.449 | כי |  |
| 6:11.450 | רק |  |
| 6:11.451 | מתחלת |  |
| 6:11.452 | האלף |  |
| 6:11.453 | הששי |  |
| 6:11.454 | בסוף |  |
| 6:11.455 | המאה |  |
| 6:11.456 | הראשונה |  |
| 6:11.457 | החזיקו |  |
| 6:11.458 | ברעיונות |  |
| 6:11.459 | האלה |  |
| 6:11.460 | קצת |  |
| 6:11.461 | מחברים |  |
| 6:11.462 | שנמשכו |  |
| 6:11.463 | אחר |  |
| 6:11.464 | דעות |  |
| 6:11.465 | הפילסוף |  |
| 6:11.466 | הרמאי |  |
| 6:11.467 | שדבק |  |
| 6:11.468 | בדעת |  |
| 6:11.469 | השונים |  |
| 6:11.470 | והשלישי |  |
| 6:11.471 | הוא |  |
| 6:11.472 | מחבר |  |
| 6:11.473 | ספר |  |
| 6:11.474 | הזהר |  |
| 6:11.475 | והתקונים | . |
| 6:11.476 | וכוון |  |
| 6:11.477 | בו |  |
| 6:11.478 | להטעות |  |
| 6:11.479 | את |  |
| 6:11.480 | ישראל |  |
| 6:11.481 | ולהדיחם |  |
| 6:11.482 | מעל |  |
| 6:11.483 | ה | ' |
| 6:11.484 | אלהינו | . |
| 6:11.485 | ותלה |  |
| 6:11.486 | רעיונותיו |  |
| 6:11.487 | ברשב"י |  |
| 6:11.488 | התנא |  |
| 6:11.489 | וחביריו | , |
| 6:11.490 | בפרט |  |
| 6:11.491 | בשני |  |
| 6:11.492 | דברים | , |
| 6:11.493 | מידי |  |
| 6:11.494 | דעל |  |
| 6:11.495 | רשב"י |  |
| 6:11.496 | תני |  |
| 6:11.497 | במאמרו |  |
| 6:11.498 | הקצר |  |
| 6:11.499 | כל |  |
| 6:11.500 | המשתף |  |
| 6:11.501 | שם |  |
| 6:11.502 | שמים |  |
| 6:11.503 | ודבר |  |
| 6:11.504 | אחר |  |
| 6:11.505 | נעקר |  |
| 6:11.506 | מן |  |
| 6:11.507 | העולם | . |
| 6:11.508 | והשני |  |
| 6:11.509 | בביזוי |  |
| 6:11.510 | המשנה |  |
| 6:11.511 | שקלא |  |
| 6:11.512 | וטריא |  |
| 6:11.513 | דבתלמודא |  |
| 6:11.514 | שהם |  |
| 6:11.515 | עיקר |  |
| 6:11.516 | ויסוד |  |
| 6:11.517 | תורה |  |
| 6:11.518 | שבעל |  |
| 6:11.519 | פה |  |
| 6:11.520 | שהוא |  |
| 6:11.521 | מבזה |  |
| 6:11.522 | אותן |  |
| 6:11.523 | כמו |  |
| 6:11.524 | שיתבאר |  |
| 6:11.525 | להלן |  |
| 6:11.526 | דלרשב"י |  |
| 6:11.527 | שתורתו |  |
| 6:11.528 | אומנותו |  |
| 6:11.529 | היא |  |
| 6:11.530 | כקריאת |  |
| 6:11.531 | שמע |  |
| 6:11.532 | בעונתה |  |
| 6:11.533 | ירוש | ' |
| 6:11.534 | שבת |  |
| 6:11.535 | פ"א | ). |
| 6:11.536 | וטמן |  |
| 6:11.537 | את |  |
| 6:11.538 | הספר |  |
| 6:11.539 | הנז | ' |
| 6:11.540 | בארון |  |
| 6:11.541 | וקברו |  |
| 6:11.542 | בקרקע |  |
| 6:11.543 | כדי |  |
| 6:11.544 | לישנו |  |
| 6:11.545 | כאלו |  |
| 6:11.546 | הוא |  |
| 6:11.547 | ישן |  |
| 6:11.548 | נושן |  |
| 6:11.549 | מימי |  |
| 6:11.550 | רשב"י | . |
| 6:11.551 | ואמר |  |
| 6:11.552 | למלך |  |
| 6:11.553 | ממלכי |  |
| 6:11.554 | מזרח |  |
| 6:11.555 | לחפור |  |
| 6:11.556 | שם |  |
| 6:11.557 | בקרקע |  |
| 6:11.558 | למצוא |  |
| 6:11.559 | מטמון |  |
| 6:11.560 | של |  |
| 6:11.561 | כסף |  |
| 6:11.562 | וזהב | , |
| 6:11.563 | וצוה |  |
| 6:11.564 | המלך |  |
| 6:11.565 | הנז | ' |
| 6:11.566 | לחפור |  |
| 6:11.567 | על |  |
| 6:11.568 | עסקי |  |
| 6:11.569 | ממון |  |
| 6:11.570 | ונמצא |  |
| 6:11.571 | שם |  |
| 6:11.572 | ארון |  |
| 6:11.573 | ובו |  |
| 6:11.574 | ס | ' |
| 6:11.575 | הזהר |  |
| 6:11.576 | וכו | ' |
| 6:11.577 | ושלח |  |
| 6:11.578 | אותו |  |
| 6:11.579 | לטוליטולא |  |
| 6:11.580 | וכו | ' |
| 6:11.581 | ומשם |  |
| 6:11.582 | נתפשטה |  |
| 6:11.583 | הקבלה |  |
| 6:11.584 | החדשה |  |
| 6:11.585 | בישראל |  |
| 6:11.586 | כמו |  |
| 6:11.587 | שכ | ' |
| 6:11.588 | הר | ' |
| 6:11.589 | חיד"א |  |
| 6:11.590 | בספרו |  |
| 6:11.591 | הנחמד |  |
| 6:11.592 | שם |  |
| 6:11.593 | הגדולים |  |
| 6:11.594 | במערכת |  |
| 6:11.595 | ספרים |  |
| 6:11.596 | ערך |  |
| 6:11.597 | זהר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

