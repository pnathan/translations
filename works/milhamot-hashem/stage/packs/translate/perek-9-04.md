# Translation pack: milhamot-hashem perek-9-04 (Milchamot ha-Shem 9:4–9:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-04`.

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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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


## Milchamot ha-Shem 9:4

Hebrew (Wikisource): קכג) ובנדון דידן איכא קל וחומר גדול ועצום, ומה אם להחליף תורה או מצוה מכל המצות לא נשמעו לו כל שכן שלא נשמע למלך העכו"ם שמסר לנו את הספר הנזכר המסית להחליף ולהמיר את ה' אלהינו אחד בחמשה פרצופים, וביחוד לזעיר אנפין ופילגשו ולמר בפנינו ולחשוב בדעתינו שהוא אלהינו, והוא המלך המהולל בכל התושבחות, וכל המצות שנעשה נעשם לשם האלוה קצר האפים. ונשיב אל לבינו בשום שכל, ונדע בבירור כי מנסה ה' אלהינו אותנו כמו שאמר השי"ת בתורה כי יקום בקרבך נביא וכו' ובא האות והמופת וכו' כי מנסה ה' אלהיכם אתכם. לדעת הישכם אוהבים את ה' אלהיכם, וכמו שביאר הרבה המורה בפרק כ"ד משלישי שרוצה לומר כי בשיקום מתפאר בנבואה, ותראו אותותיו ומופתיו המביאות לחשוב אמת בדבריו. דעו שהוא ענין שירצהו השם ית' להודיע באומות שיעור האמינכם באמתת תורתו והשגתם אמתתו ית' ושאינכם ניסיתם להסתת מסית ולא תפסד אמונתכם בשם ית', למען יכוין אליה כל מבקש אמת. ויבקש מן האמונות מה שהוא אמת ויציב, וקיים קיום שאין לפנות עמו עוד לעשיית אות ומופת, שהמפתה לעשות הפכו באות או מופת. הוא קורא להאמין בנמנעות, כי לוא יועיל המופת רק למתנבא בדבר אפשרי כמו שביארנו במשנה תורה עכ"ל. אוי נא לנו כי נסה ה' אלהינו אותנו בהמציאו לנו את ספר הנזכר הזהר והתיקונים על ידי המלך העכו"ם ולא עמדנו בנסיון, ונפתינו להסתת המסית הנזכר לעבודת אלהים אחרים רבים, המשתתפים עם הבן קצר אפים (ז"א) ופילגשו על ידי הבטחתו שהוא מבטיח את קוראיו, שהזעיר אנפין מושל השפע מאצל העתיק כמו שכתב הרב יושר לבב וכסא אליהו! וכדי לפתותינו הוא משבח את סחורתו בכל דיבור במלות זכאה חולקיהן או זכאה הולקיה דמאן דעביד כך וכך, ווי ליה מאן דלא אזדהר בכך וכך, וכמתלהלה היורה זיקים חצים ומות, והא מפחית ומשפיל ומוריד את כבוד המשנה והתלמוד כלאחר יד, ומגדיל ומשבח את תורתו החדשה, ומתנבא של עתיד בסוף הימים, יעזבו את המשנה והתלמוד ויתפרנסון מחבורא דיליה, ועל ידי הלמוד בספריו יזכו להגאל, ובדברים האלו לבדן בלי שום אות ומופת צלחה לו להטות לב רבים וגדולים מישראל לעבוד אלהים חדשים בעלי גוף וגויה רמ"ח איברים ושס"ה גידים גדולים ועצומים, זכים ובהירים מאירים ומזהירים יתארם הקורא בשרד שכלו, ובמחוגת בינתו, אך ורק שהחומר שלהם איננו מדם ובשר _כמו שכתב הרב המורה בפרק א' מראשון) וקורא להם בשמותיו של הקב"ה.

| id | word | punctuation after |
|---|---|---|
| 9:4.1 | קכג | ) |
| 9:4.2 | ובנדון |  |
| 9:4.3 | דידן |  |
| 9:4.4 | איכא |  |
| 9:4.5 | קל |  |
| 9:4.6 | וחומר |  |
| 9:4.7 | גדול |  |
| 9:4.8 | ועצום | , |
| 9:4.9 | ומה |  |
| 9:4.10 | אם |  |
| 9:4.11 | להחליף |  |
| 9:4.12 | תורה |  |
| 9:4.13 | או |  |
| 9:4.14 | מצוה |  |
| 9:4.15 | מכל |  |
| 9:4.16 | המצות |  |
| 9:4.17 | לא |  |
| 9:4.18 | נשמעו |  |
| 9:4.19 | לו |  |
| 9:4.20 | כל |  |
| 9:4.21 | שכן |  |
| 9:4.22 | שלא |  |
| 9:4.23 | נשמע |  |
| 9:4.24 | למלך |  |
| 9:4.25 | העכו"ם |  |
| 9:4.26 | שמסר |  |
| 9:4.27 | לנו |  |
| 9:4.28 | את |  |
| 9:4.29 | הספר |  |
| 9:4.30 | הנזכר |  |
| 9:4.31 | המסית |  |
| 9:4.32 | להחליף |  |
| 9:4.33 | ולהמיר |  |
| 9:4.34 | את |  |
| 9:4.35 | ה | ' |
| 9:4.36 | אלהינו |  |
| 9:4.37 | אחד |  |
| 9:4.38 | בחמשה |  |
| 9:4.39 | פרצופים | , |
| 9:4.40 | וביחוד |  |
| 9:4.41 | לזעיר |  |
| 9:4.42 | אנפין |  |
| 9:4.43 | ופילגשו |  |
| 9:4.44 | ולמר |  |
| 9:4.45 | בפנינו |  |
| 9:4.46 | ולחשוב |  |
| 9:4.47 | בדעתינו |  |
| 9:4.48 | שהוא |  |
| 9:4.49 | אלהינו | , |
| 9:4.50 | והוא |  |
| 9:4.51 | המלך |  |
| 9:4.52 | המהולל |  |
| 9:4.53 | בכל |  |
| 9:4.54 | התושבחות | , |
| 9:4.55 | וכל |  |
| 9:4.56 | המצות |  |
| 9:4.57 | שנעשה |  |
| 9:4.58 | נעשם |  |
| 9:4.59 | לשם |  |
| 9:4.60 | האלוה |  |
| 9:4.61 | קצר |  |
| 9:4.62 | האפים | . |
| 9:4.63 | ונשיב |  |
| 9:4.64 | אל |  |
| 9:4.65 | לבינו |  |
| 9:4.66 | בשום |  |
| 9:4.67 | שכל | , |
| 9:4.68 | ונדע |  |
| 9:4.69 | בבירור |  |
| 9:4.70 | כי |  |
| 9:4.71 | מנסה |  |
| 9:4.72 | ה | ' |
| 9:4.73 | אלהינו |  |
| 9:4.74 | אותנו |  |
| 9:4.75 | כמו |  |
| 9:4.76 | שאמר |  |
| 9:4.77 | השי"ת |  |
| 9:4.78 | בתורה |  |
| 9:4.79 | כי |  |
| 9:4.80 | יקום |  |
| 9:4.81 | בקרבך |  |
| 9:4.82 | נביא |  |
| 9:4.83 | וכו | ' |
| 9:4.84 | ובא |  |
| 9:4.85 | האות |  |
| 9:4.86 | והמופת |  |
| 9:4.87 | וכו | ' |
| 9:4.88 | כי |  |
| 9:4.89 | מנסה |  |
| 9:4.90 | ה | ' |
| 9:4.91 | אלהיכם |  |
| 9:4.92 | אתכם | . |
| 9:4.93 | לדעת |  |
| 9:4.94 | הישכם |  |
| 9:4.95 | אוהבים |  |
| 9:4.96 | את |  |
| 9:4.97 | ה | ' |
| 9:4.98 | אלהיכם | , |
| 9:4.99 | וכמו |  |
| 9:4.100 | שביאר |  |
| 9:4.101 | הרבה |  |
| 9:4.102 | המורה |  |
| 9:4.103 | בפרק |  |
| 9:4.104 | כ"ד |  |
| 9:4.105 | משלישי |  |
| 9:4.106 | שרוצה |  |
| 9:4.107 | לומר |  |
| 9:4.108 | כי |  |
| 9:4.109 | בשיקום |  |
| 9:4.110 | מתפאר |  |
| 9:4.111 | בנבואה | , |
| 9:4.112 | ותראו |  |
| 9:4.113 | אותותיו |  |
| 9:4.114 | ומופתיו |  |
| 9:4.115 | המביאות |  |
| 9:4.116 | לחשוב |  |
| 9:4.117 | אמת |  |
| 9:4.118 | בדבריו | . |
| 9:4.119 | דעו |  |
| 9:4.120 | שהוא |  |
| 9:4.121 | ענין |  |
| 9:4.122 | שירצהו |  |
| 9:4.123 | השם |  |
| 9:4.124 | ית | ' |
| 9:4.125 | להודיע |  |
| 9:4.126 | באומות |  |
| 9:4.127 | שיעור |  |
| 9:4.128 | האמינכם |  |
| 9:4.129 | באמתת |  |
| 9:4.130 | תורתו |  |
| 9:4.131 | והשגתם |  |
| 9:4.132 | אמתתו |  |
| 9:4.133 | ית | ' |
| 9:4.134 | ושאינכם |  |
| 9:4.135 | ניסיתם |  |
| 9:4.136 | להסתת |  |
| 9:4.137 | מסית |  |
| 9:4.138 | ולא |  |
| 9:4.139 | תפסד |  |
| 9:4.140 | אמונתכם |  |
| 9:4.141 | בשם |  |
| 9:4.142 | ית | ', |
| 9:4.143 | למען |  |
| 9:4.144 | יכוין |  |
| 9:4.145 | אליה |  |
| 9:4.146 | כל |  |
| 9:4.147 | מבקש |  |
| 9:4.148 | אמת | . |
| 9:4.149 | ויבקש |  |
| 9:4.150 | מן |  |
| 9:4.151 | האמונות |  |
| 9:4.152 | מה |  |
| 9:4.153 | שהוא |  |
| 9:4.154 | אמת |  |
| 9:4.155 | ויציב | , |
| 9:4.156 | וקיים |  |
| 9:4.157 | קיום |  |
| 9:4.158 | שאין |  |
| 9:4.159 | לפנות |  |
| 9:4.160 | עמו |  |
| 9:4.161 | עוד |  |
| 9:4.162 | לעשיית |  |
| 9:4.163 | אות |  |
| 9:4.164 | ומופת | , |
| 9:4.165 | שהמפתה |  |
| 9:4.166 | לעשות |  |
| 9:4.167 | הפכו |  |
| 9:4.168 | באות |  |
| 9:4.169 | או |  |
| 9:4.170 | מופת | . |
| 9:4.171 | הוא |  |
| 9:4.172 | קורא |  |
| 9:4.173 | להאמין |  |
| 9:4.174 | בנמנעות | , |
| 9:4.175 | כי |  |
| 9:4.176 | לוא |  |
| 9:4.177 | יועיל |  |
| 9:4.178 | המופת |  |
| 9:4.179 | רק |  |
| 9:4.180 | למתנבא |  |
| 9:4.181 | בדבר |  |
| 9:4.182 | אפשרי |  |
| 9:4.183 | כמו |  |
| 9:4.184 | שביארנו |  |
| 9:4.185 | במשנה |  |
| 9:4.186 | תורה |  |
| 9:4.187 | עכ"ל | . |
| 9:4.188 | אוי |  |
| 9:4.189 | נא |  |
| 9:4.190 | לנו |  |
| 9:4.191 | כי |  |
| 9:4.192 | נסה |  |
| 9:4.193 | ה | ' |
| 9:4.194 | אלהינו |  |
| 9:4.195 | אותנו |  |
| 9:4.196 | בהמציאו |  |
| 9:4.197 | לנו |  |
| 9:4.198 | את |  |
| 9:4.199 | ספר |  |
| 9:4.200 | הנזכר |  |
| 9:4.201 | הזהר |  |
| 9:4.202 | והתיקונים |  |
| 9:4.203 | על |  |
| 9:4.204 | ידי |  |
| 9:4.205 | המלך |  |
| 9:4.206 | העכו"ם |  |
| 9:4.207 | ולא |  |
| 9:4.208 | עמדנו |  |
| 9:4.209 | בנסיון | , |
| 9:4.210 | ונפתינו |  |
| 9:4.211 | להסתת |  |
| 9:4.212 | המסית |  |
| 9:4.213 | הנזכר |  |
| 9:4.214 | לעבודת |  |
| 9:4.215 | אלהים |  |
| 9:4.216 | אחרים |  |
| 9:4.217 | רבים | , |
| 9:4.218 | המשתתפים |  |
| 9:4.219 | עם |  |
| 9:4.220 | הבן |  |
| 9:4.221 | קצר |  |
| 9:4.222 | אפים |  |
| 9:4.223 | ז"א | ) |
| 9:4.224 | ופילגשו |  |
| 9:4.225 | על |  |
| 9:4.226 | ידי |  |
| 9:4.227 | הבטחתו |  |
| 9:4.228 | שהוא |  |
| 9:4.229 | מבטיח |  |
| 9:4.230 | את |  |
| 9:4.231 | קוראיו | , |
| 9:4.232 | שהזעיר |  |
| 9:4.233 | אנפין |  |
| 9:4.234 | מושל |  |
| 9:4.235 | השפע |  |
| 9:4.236 | מאצל |  |
| 9:4.237 | העתיק |  |
| 9:4.238 | כמו |  |
| 9:4.239 | שכתב |  |
| 9:4.240 | הרב |  |
| 9:4.241 | יושר |  |
| 9:4.242 | לבב |  |
| 9:4.243 | וכסא |  |
| 9:4.244 | אליהו | ! |
| 9:4.245 | וכדי |  |
| 9:4.246 | לפתותינו |  |
| 9:4.247 | הוא |  |
| 9:4.248 | משבח |  |
| 9:4.249 | את |  |
| 9:4.250 | סחורתו |  |
| 9:4.251 | בכל |  |
| 9:4.252 | דיבור |  |
| 9:4.253 | במלות |  |
| 9:4.254 | זכאה |  |
| 9:4.255 | חולקיהן |  |
| 9:4.256 | או |  |
| 9:4.257 | זכאה |  |
| 9:4.258 | הולקיה |  |
| 9:4.259 | דמאן |  |
| 9:4.260 | דעביד |  |
| 9:4.261 | כך |  |
| 9:4.262 | וכך | , |
| 9:4.263 | ווי |  |
| 9:4.264 | ליה |  |
| 9:4.265 | מאן |  |
| 9:4.266 | דלא |  |
| 9:4.267 | אזדהר |  |
| 9:4.268 | בכך |  |
| 9:4.269 | וכך | , |
| 9:4.270 | וכמתלהלה |  |
| 9:4.271 | היורה |  |
| 9:4.272 | זיקים |  |
| 9:4.273 | חצים |  |
| 9:4.274 | ומות | , |
| 9:4.275 | והא |  |
| 9:4.276 | מפחית |  |
| 9:4.277 | ומשפיל |  |
| 9:4.278 | ומוריד |  |
| 9:4.279 | את |  |
| 9:4.280 | כבוד |  |
| 9:4.281 | המשנה |  |
| 9:4.282 | והתלמוד |  |
| 9:4.283 | כלאחר |  |
| 9:4.284 | יד | , |
| 9:4.285 | ומגדיל |  |
| 9:4.286 | ומשבח |  |
| 9:4.287 | את |  |
| 9:4.288 | תורתו |  |
| 9:4.289 | החדשה | , |
| 9:4.290 | ומתנבא |  |
| 9:4.291 | של |  |
| 9:4.292 | עתיד |  |
| 9:4.293 | בסוף |  |
| 9:4.294 | הימים | , |
| 9:4.295 | יעזבו |  |
| 9:4.296 | את |  |
| 9:4.297 | המשנה |  |
| 9:4.298 | והתלמוד |  |
| 9:4.299 | ויתפרנסון |  |
| 9:4.300 | מחבורא |  |
| 9:4.301 | דיליה | , |
| 9:4.302 | ועל |  |
| 9:4.303 | ידי |  |
| 9:4.304 | הלמוד |  |
| 9:4.305 | בספריו |  |
| 9:4.306 | יזכו |  |
| 9:4.307 | להגאל | , |
| 9:4.308 | ובדברים |  |
| 9:4.309 | האלו |  |
| 9:4.310 | לבדן |  |
| 9:4.311 | בלי |  |
| 9:4.312 | שום |  |
| 9:4.313 | אות |  |
| 9:4.314 | ומופת |  |
| 9:4.315 | צלחה |  |
| 9:4.316 | לו |  |
| 9:4.317 | להטות |  |
| 9:4.318 | לב |  |
| 9:4.319 | רבים |  |
| 9:4.320 | וגדולים |  |
| 9:4.321 | מישראל |  |
| 9:4.322 | לעבוד |  |
| 9:4.323 | אלהים |  |
| 9:4.324 | חדשים |  |
| 9:4.325 | בעלי |  |
| 9:4.326 | גוף |  |
| 9:4.327 | וגויה |  |
| 9:4.328 | רמ"ח |  |
| 9:4.329 | איברים |  |
| 9:4.330 | ושס"ה |  |
| 9:4.331 | גידים |  |
| 9:4.332 | גדולים |  |
| 9:4.333 | ועצומים | , |
| 9:4.334 | זכים |  |
| 9:4.335 | ובהירים |  |
| 9:4.336 | מאירים |  |
| 9:4.337 | ומזהירים |  |
| 9:4.338 | יתארם |  |
| 9:4.339 | הקורא |  |
| 9:4.340 | בשרד |  |
| 9:4.341 | שכלו | , |
| 9:4.342 | ובמחוגת |  |
| 9:4.343 | בינתו | , |
| 9:4.344 | אך |  |
| 9:4.345 | ורק |  |
| 9:4.346 | שהחומר |  |
| 9:4.347 | שלהם |  |
| 9:4.348 | איננו |  |
| 9:4.349 | מדם |  |
| 9:4.350 | ובשר |  |
| 9:4.351 | כמו |  |
| 9:4.352 | שכתב |  |
| 9:4.353 | הרב |  |
| 9:4.354 | המורה |  |
| 9:4.355 | בפרק |  |
| 9:4.356 | א | ' |
| 9:4.357 | מראשון | ) |
| 9:4.358 | וקורא |  |
| 9:4.359 | להם |  |
| 9:4.360 | בשמותיו |  |
| 9:4.361 | של |  |
| 9:4.362 | הקב"ה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

