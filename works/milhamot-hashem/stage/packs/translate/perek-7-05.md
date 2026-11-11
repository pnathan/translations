# Translation pack: milhamot-hashem perek-7-05 (Milchamot ha-Shem 7:6–7:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-05`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
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
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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


## Milchamot ha-Shem 7:6

Hebrew (Wikisource): פג) עוד שם במדרש על פסוק ירא ה' את בני ומלך, מהו מלך, המליכהו עליך ועם שונים אל תתערב. עם אלו שאומרים יש אלוה שני אל תתערב וכו'. אמר ר' אחא כעס הקב"ה על שלמה כשאמר הפסוק הזה. אמר לו דבר של קדוש שמי היית אומר בלשון נוטריקון עם שונים אל תתערב מיד חזר ופירש את הדבר. יש אחד ואין שני גם בן ואח אין לו. אין לו אח ולא בן. אלא שמע ישראל ה' אלהינו ה' אחד עכ"ל ופי' הר' עץ יוסף. יש אלוה שני. כי המינים אומרים איך יתכן פועל טוב ורע יוצא ממקור אחד? לכן הם מאמינים בשניות וכו'. ובילקוט קהלת רמז תתקע"א. ובקהלת רבותי יש אחד ואין שני. אחד זה הקב"ה דכתיב ביה שמע ישראל ה' אלהינו ה' אחד. ואין שני גם בן ואח אין לו. אח מהיכן יש לו? אלא שחיבב את ישראל וקרא אותם בנים ואחים. ואין קץ לכל עמלו שעמל בששת ימי בראשית (פי' עמל הנז' כאן איננו לשון טורח וידיעה. שאין יגיעה וטורח לפני מי שאמר והיה העולם. אלא הוא לשון פעולה, שכן בלשון ישמעאל קורין לפעולה עמל. וכן תפרש לשון הפיטן מתקן הסליחות דאשמורות שאמר חוסה על עמלך כמו פעלך) ולמי אני עמל ומחסר את נפשי מטובה, לא להדבק בו וללכת בדרכיו? וכל מי שאינו עושה כן. גם זה הבל וענין רע הוא עכ"ל לעניינינו.

| id | word | punctuation after |
|---|---|---|
| 7:6.1 | פג | ) |
| 7:6.2 | עוד |  |
| 7:6.3 | שם |  |
| 7:6.4 | במדרש |  |
| 7:6.5 | על |  |
| 7:6.6 | פסוק |  |
| 7:6.7 | ירא |  |
| 7:6.8 | ה | ' |
| 7:6.9 | את |  |
| 7:6.10 | בני |  |
| 7:6.11 | ומלך | , |
| 7:6.12 | מהו |  |
| 7:6.13 | מלך | , |
| 7:6.14 | המליכהו |  |
| 7:6.15 | עליך |  |
| 7:6.16 | ועם |  |
| 7:6.17 | שונים |  |
| 7:6.18 | אל |  |
| 7:6.19 | תתערב | . |
| 7:6.20 | עם |  |
| 7:6.21 | אלו |  |
| 7:6.22 | שאומרים |  |
| 7:6.23 | יש |  |
| 7:6.24 | אלוה |  |
| 7:6.25 | שני |  |
| 7:6.26 | אל |  |
| 7:6.27 | תתערב |  |
| 7:6.28 | וכו | '. |
| 7:6.29 | אמר |  |
| 7:6.30 | ר | ' |
| 7:6.31 | אחא |  |
| 7:6.32 | כעס |  |
| 7:6.33 | הקב"ה |  |
| 7:6.34 | על |  |
| 7:6.35 | שלמה |  |
| 7:6.36 | כשאמר |  |
| 7:6.37 | הפסוק |  |
| 7:6.38 | הזה | . |
| 7:6.39 | אמר |  |
| 7:6.40 | לו |  |
| 7:6.41 | דבר |  |
| 7:6.42 | של |  |
| 7:6.43 | קדוש |  |
| 7:6.44 | שמי |  |
| 7:6.45 | היית |  |
| 7:6.46 | אומר |  |
| 7:6.47 | בלשון |  |
| 7:6.48 | נוטריקון |  |
| 7:6.49 | עם |  |
| 7:6.50 | שונים |  |
| 7:6.51 | אל |  |
| 7:6.52 | תתערב |  |
| 7:6.53 | מיד |  |
| 7:6.54 | חזר |  |
| 7:6.55 | ופירש |  |
| 7:6.56 | את |  |
| 7:6.57 | הדבר | . |
| 7:6.58 | יש |  |
| 7:6.59 | אחד |  |
| 7:6.60 | ואין |  |
| 7:6.61 | שני |  |
| 7:6.62 | גם |  |
| 7:6.63 | בן |  |
| 7:6.64 | ואח |  |
| 7:6.65 | אין |  |
| 7:6.66 | לו | . |
| 7:6.67 | אין |  |
| 7:6.68 | לו |  |
| 7:6.69 | אח |  |
| 7:6.70 | ולא |  |
| 7:6.71 | בן | . |
| 7:6.72 | אלא |  |
| 7:6.73 | שמע |  |
| 7:6.74 | ישראל |  |
| 7:6.75 | ה | ' |
| 7:6.76 | אלהינו |  |
| 7:6.77 | ה | ' |
| 7:6.78 | אחד |  |
| 7:6.79 | עכ"ל |  |
| 7:6.80 | ופי | ' |
| 7:6.81 | הר | ' |
| 7:6.82 | עץ |  |
| 7:6.83 | יוסף | . |
| 7:6.84 | יש |  |
| 7:6.85 | אלוה |  |
| 7:6.86 | שני | . |
| 7:6.87 | כי |  |
| 7:6.88 | המינים |  |
| 7:6.89 | אומרים |  |
| 7:6.90 | איך |  |
| 7:6.91 | יתכן |  |
| 7:6.92 | פועל |  |
| 7:6.93 | טוב |  |
| 7:6.94 | ורע |  |
| 7:6.95 | יוצא |  |
| 7:6.96 | ממקור |  |
| 7:6.97 | אחד | ? |
| 7:6.98 | לכן |  |
| 7:6.99 | הם |  |
| 7:6.100 | מאמינים |  |
| 7:6.101 | בשניות |  |
| 7:6.102 | וכו | '. |
| 7:6.103 | ובילקוט |  |
| 7:6.104 | קהלת |  |
| 7:6.105 | רמז |  |
| 7:6.106 | תתקע"א | . |
| 7:6.107 | ובקהלת |  |
| 7:6.108 | רבותי |  |
| 7:6.109 | יש |  |
| 7:6.110 | אחד |  |
| 7:6.111 | ואין |  |
| 7:6.112 | שני | . |
| 7:6.113 | אחד |  |
| 7:6.114 | זה |  |
| 7:6.115 | הקב"ה |  |
| 7:6.116 | דכתיב |  |
| 7:6.117 | ביה |  |
| 7:6.118 | שמע |  |
| 7:6.119 | ישראל |  |
| 7:6.120 | ה | ' |
| 7:6.121 | אלהינו |  |
| 7:6.122 | ה | ' |
| 7:6.123 | אחד | . |
| 7:6.124 | ואין |  |
| 7:6.125 | שני |  |
| 7:6.126 | גם |  |
| 7:6.127 | בן |  |
| 7:6.128 | ואח |  |
| 7:6.129 | אין |  |
| 7:6.130 | לו | . |
| 7:6.131 | אח |  |
| 7:6.132 | מהיכן |  |
| 7:6.133 | יש |  |
| 7:6.134 | לו | ? |
| 7:6.135 | אלא |  |
| 7:6.136 | שחיבב |  |
| 7:6.137 | את |  |
| 7:6.138 | ישראל |  |
| 7:6.139 | וקרא |  |
| 7:6.140 | אותם |  |
| 7:6.141 | בנים |  |
| 7:6.142 | ואחים | . |
| 7:6.143 | ואין |  |
| 7:6.144 | קץ |  |
| 7:6.145 | לכל |  |
| 7:6.146 | עמלו |  |
| 7:6.147 | שעמל |  |
| 7:6.148 | בששת |  |
| 7:6.149 | ימי |  |
| 7:6.150 | בראשית |  |
| 7:6.151 | פי | ' |
| 7:6.152 | עמל |  |
| 7:6.153 | הנז | ' |
| 7:6.154 | כאן |  |
| 7:6.155 | איננו |  |
| 7:6.156 | לשון |  |
| 7:6.157 | טורח |  |
| 7:6.158 | וידיעה | . |
| 7:6.159 | שאין |  |
| 7:6.160 | יגיעה |  |
| 7:6.161 | וטורח |  |
| 7:6.162 | לפני |  |
| 7:6.163 | מי |  |
| 7:6.164 | שאמר |  |
| 7:6.165 | והיה |  |
| 7:6.166 | העולם | . |
| 7:6.167 | אלא |  |
| 7:6.168 | הוא |  |
| 7:6.169 | לשון |  |
| 7:6.170 | פעולה | , |
| 7:6.171 | שכן |  |
| 7:6.172 | בלשון |  |
| 7:6.173 | ישמעאל |  |
| 7:6.174 | קורין |  |
| 7:6.175 | לפעולה |  |
| 7:6.176 | עמל | . |
| 7:6.177 | וכן |  |
| 7:6.178 | תפרש |  |
| 7:6.179 | לשון |  |
| 7:6.180 | הפיטן |  |
| 7:6.181 | מתקן |  |
| 7:6.182 | הסליחות |  |
| 7:6.183 | דאשמורות |  |
| 7:6.184 | שאמר |  |
| 7:6.185 | חוסה |  |
| 7:6.186 | על |  |
| 7:6.187 | עמלך |  |
| 7:6.188 | כמו |  |
| 7:6.189 | פעלך | ) |
| 7:6.190 | ולמי |  |
| 7:6.191 | אני |  |
| 7:6.192 | עמל |  |
| 7:6.193 | ומחסר |  |
| 7:6.194 | את |  |
| 7:6.195 | נפשי |  |
| 7:6.196 | מטובה | , |
| 7:6.197 | לא |  |
| 7:6.198 | להדבק |  |
| 7:6.199 | בו |  |
| 7:6.200 | וללכת |  |
| 7:6.201 | בדרכיו | ? |
| 7:6.202 | וכל |  |
| 7:6.203 | מי |  |
| 7:6.204 | שאינו |  |
| 7:6.205 | עושה |  |
| 7:6.206 | כן | . |
| 7:6.207 | גם |  |
| 7:6.208 | זה |  |
| 7:6.209 | הבל |  |
| 7:6.210 | וענין |  |
| 7:6.211 | רע |  |
| 7:6.212 | הוא |  |
| 7:6.213 | עכ"ל |  |
| 7:6.214 | לעניינינו | . |


## Milchamot ha-Shem 7:7

Hebrew (Wikisource): ומדבריהם למדנו דפשיטא להו לרבנן כי יחוד השם, ית' בשמע ישראל הוא מבלי לכלול עמו אבא ואימא וזעיר ונוקביה, ודלא כמו שאמר בזהר ואתחנן המובא לעיל סי' מ"ב ועיין שם בשם ר' חטר שהמכוין כן דעת חיצונית היא ולא כאלה חלק יעקוב, וכשתתבונן בדברי הקבלה החדשה תמצא שלא זו בלבד שהם מאמינים במציאות אלהים רבים קדושים כי אף גם זאת שהם מאמינים במציאות אלהים אחרים טמאים בעלי עשר ספירות, ונקראים ספירות דקליפה. ופרצופים רבים להם כמו בקדושה כדכתבינן לעיל. ונקראים אצלם סטרא אחרא מסאבא כמבואר כל זה בספריהם והם הפועלים את הרע בעולם. ולפעמים שולטים גם כן על האלהים הקדושים אשר יעבדו אותם המקובלים. ומשתעבדים בהם בחזקת היד כמו שכתב בעל משנת חסידים מסכת ליל פסח פרק ז' משנה ג' וז"ל ואלו לא גאל הקב"ה (אריך אנפין) את אבותינו (אבא ואימא) ממצרים (אדם בליעל) האחוז בגרון דאריך אנפין ונוטל השפע ומונעו מרדת לזעיר אנפין והיה זעיר אנפין נשאר בגלות וצער גדול כולד המעובר במעי אמו. עיין שם גם במשנה ד' שם, והנה כבר הבאתי ס' י"ג וי"ד מקצת מדברי הר' הגדול מהרש"ל בתשובותיו ובס' ים של שלמה הרבה מהמדקדקים דקדוקי עניות להניח תפילין של יד מיושב. דע אהובי כי חדשים מקרוב באו ורוצים להיות מכת המקובלים וממדרשי הנעלמים ומחלושי הראות. לא יביטו באור הזהר ולא ידעו מצאו ומובאו וכוונתו. אלא שכך מצאו בספרי הרשב"י ודע שכל רבותי הקדושים ששמשו גאוני עולם לא נהגו כך אלא כדברי התלמוד והפוסקים. ועתה פקח עיניך וראה אם יש בזה ממש אי אפשר שלא הזכירו שום מחבר. או גם לא אשתמיט רמז בתלמוד בבלי או ירושלמי ספרא וספרי ושאלתות ופסיקתא. ומי שלא ידע להשיג סודה על נכון יבוא לקצץ בנטיעות. לכך אהובי אל תלך בדרך אתם. ואין לך עסק בנסתרות המתיהרי' בחידושים. כאלו הם יודעים ומבינים רזי תורה וספוניה. והלואי שידעו הגלוים. אבל דע שכל המשנה ידו על התחתונה. וידינו על העליונה עכ"ל. עוד כ' רביד הזהב בשמו וז"ל והמורה היפך התלמוד חייב מיתה. וכמו שכ' בס' צח ואדום בשם רש"ל בתשו' החדשות וז"ל שכל המחמיר שלא לאכול בשר אחד גבינה בלא גריעותא דאית ליה בשינוהי. כי המחמיר חייב מיתה וכו' ובס' ים שלמה כ' דמאחר שנחתם התלמוד אין להחמיר נגד התלמוד והוי כמו מינות ע"ש.

| id | word | punctuation after |
|---|---|---|
| 7:7.1 | ומדבריהם |  |
| 7:7.2 | למדנו |  |
| 7:7.3 | דפשיטא |  |
| 7:7.4 | להו |  |
| 7:7.5 | לרבנן |  |
| 7:7.6 | כי |  |
| 7:7.7 | יחוד |  |
| 7:7.8 | השם | , |
| 7:7.9 | ית | ' |
| 7:7.10 | בשמע |  |
| 7:7.11 | ישראל |  |
| 7:7.12 | הוא |  |
| 7:7.13 | מבלי |  |
| 7:7.14 | לכלול |  |
| 7:7.15 | עמו |  |
| 7:7.16 | אבא |  |
| 7:7.17 | ואימא |  |
| 7:7.18 | וזעיר |  |
| 7:7.19 | ונוקביה | , |
| 7:7.20 | ודלא |  |
| 7:7.21 | כמו |  |
| 7:7.22 | שאמר |  |
| 7:7.23 | בזהר |  |
| 7:7.24 | ואתחנן |  |
| 7:7.25 | המובא |  |
| 7:7.26 | לעיל |  |
| 7:7.27 | סי | ' |
| 7:7.28 | מ"ב |  |
| 7:7.29 | ועיין |  |
| 7:7.30 | שם |  |
| 7:7.31 | בשם |  |
| 7:7.32 | ר | ' |
| 7:7.33 | חטר |  |
| 7:7.34 | שהמכוין |  |
| 7:7.35 | כן |  |
| 7:7.36 | דעת |  |
| 7:7.37 | חיצונית |  |
| 7:7.38 | היא |  |
| 7:7.39 | ולא |  |
| 7:7.40 | כאלה |  |
| 7:7.41 | חלק |  |
| 7:7.42 | יעקוב | , |
| 7:7.43 | וכשתתבונן |  |
| 7:7.44 | בדברי |  |
| 7:7.45 | הקבלה |  |
| 7:7.46 | החדשה |  |
| 7:7.47 | תמצא |  |
| 7:7.48 | שלא |  |
| 7:7.49 | זו |  |
| 7:7.50 | בלבד |  |
| 7:7.51 | שהם |  |
| 7:7.52 | מאמינים |  |
| 7:7.53 | במציאות |  |
| 7:7.54 | אלהים |  |
| 7:7.55 | רבים |  |
| 7:7.56 | קדושים |  |
| 7:7.57 | כי |  |
| 7:7.58 | אף |  |
| 7:7.59 | גם |  |
| 7:7.60 | זאת |  |
| 7:7.61 | שהם |  |
| 7:7.62 | מאמינים |  |
| 7:7.63 | במציאות |  |
| 7:7.64 | אלהים |  |
| 7:7.65 | אחרים |  |
| 7:7.66 | טמאים |  |
| 7:7.67 | בעלי |  |
| 7:7.68 | עשר |  |
| 7:7.69 | ספירות | , |
| 7:7.70 | ונקראים |  |
| 7:7.71 | ספירות |  |
| 7:7.72 | דקליפה | . |
| 7:7.73 | ופרצופים |  |
| 7:7.74 | רבים |  |
| 7:7.75 | להם |  |
| 7:7.76 | כמו |  |
| 7:7.77 | בקדושה |  |
| 7:7.78 | כדכתבינן |  |
| 7:7.79 | לעיל | . |
| 7:7.80 | ונקראים |  |
| 7:7.81 | אצלם |  |
| 7:7.82 | סטרא |  |
| 7:7.83 | אחרא |  |
| 7:7.84 | מסאבא |  |
| 7:7.85 | כמבואר |  |
| 7:7.86 | כל |  |
| 7:7.87 | זה |  |
| 7:7.88 | בספריהם |  |
| 7:7.89 | והם |  |
| 7:7.90 | הפועלים |  |
| 7:7.91 | את |  |
| 7:7.92 | הרע |  |
| 7:7.93 | בעולם | . |
| 7:7.94 | ולפעמים |  |
| 7:7.95 | שולטים |  |
| 7:7.96 | גם |  |
| 7:7.97 | כן |  |
| 7:7.98 | על |  |
| 7:7.99 | האלהים |  |
| 7:7.100 | הקדושים |  |
| 7:7.101 | אשר |  |
| 7:7.102 | יעבדו |  |
| 7:7.103 | אותם |  |
| 7:7.104 | המקובלים | . |
| 7:7.105 | ומשתעבדים |  |
| 7:7.106 | בהם |  |
| 7:7.107 | בחזקת |  |
| 7:7.108 | היד |  |
| 7:7.109 | כמו |  |
| 7:7.110 | שכתב |  |
| 7:7.111 | בעל |  |
| 7:7.112 | משנת |  |
| 7:7.113 | חסידים |  |
| 7:7.114 | מסכת |  |
| 7:7.115 | ליל |  |
| 7:7.116 | פסח |  |
| 7:7.117 | פרק |  |
| 7:7.118 | ז | ' |
| 7:7.119 | משנה |  |
| 7:7.120 | ג | ' |
| 7:7.121 | וז"ל |  |
| 7:7.122 | ואלו |  |
| 7:7.123 | לא |  |
| 7:7.124 | גאל |  |
| 7:7.125 | הקב"ה |  |
| 7:7.126 | אריך |  |
| 7:7.127 | אנפין | ) |
| 7:7.128 | את |  |
| 7:7.129 | אבותינו |  |
| 7:7.130 | אבא |  |
| 7:7.131 | ואימא | ) |
| 7:7.132 | ממצרים |  |
| 7:7.133 | אדם |  |
| 7:7.134 | בליעל | ) |
| 7:7.135 | האחוז |  |
| 7:7.136 | בגרון |  |
| 7:7.137 | דאריך |  |
| 7:7.138 | אנפין |  |
| 7:7.139 | ונוטל |  |
| 7:7.140 | השפע |  |
| 7:7.141 | ומונעו |  |
| 7:7.142 | מרדת |  |
| 7:7.143 | לזעיר |  |
| 7:7.144 | אנפין |  |
| 7:7.145 | והיה |  |
| 7:7.146 | זעיר |  |
| 7:7.147 | אנפין |  |
| 7:7.148 | נשאר |  |
| 7:7.149 | בגלות |  |
| 7:7.150 | וצער |  |
| 7:7.151 | גדול |  |
| 7:7.152 | כולד |  |
| 7:7.153 | המעובר |  |
| 7:7.154 | במעי |  |
| 7:7.155 | אמו | . |
| 7:7.156 | עיין |  |
| 7:7.157 | שם |  |
| 7:7.158 | גם |  |
| 7:7.159 | במשנה |  |
| 7:7.160 | ד | ' |
| 7:7.161 | שם | , |
| 7:7.162 | והנה |  |
| 7:7.163 | כבר |  |
| 7:7.164 | הבאתי |  |
| 7:7.165 | ס | ' |
| 7:7.166 | י"ג |  |
| 7:7.167 | וי"ד |  |
| 7:7.168 | מקצת |  |
| 7:7.169 | מדברי |  |
| 7:7.170 | הר | ' |
| 7:7.171 | הגדול |  |
| 7:7.172 | מהרש"ל |  |
| 7:7.173 | בתשובותיו |  |
| 7:7.174 | ובס | ' |
| 7:7.175 | ים |  |
| 7:7.176 | של |  |
| 7:7.177 | שלמה |  |
| 7:7.178 | הרבה |  |
| 7:7.179 | מהמדקדקים |  |
| 7:7.180 | דקדוקי |  |
| 7:7.181 | עניות |  |
| 7:7.182 | להניח |  |
| 7:7.183 | תפילין |  |
| 7:7.184 | של |  |
| 7:7.185 | יד |  |
| 7:7.186 | מיושב | . |
| 7:7.187 | דע |  |
| 7:7.188 | אהובי |  |
| 7:7.189 | כי |  |
| 7:7.190 | חדשים |  |
| 7:7.191 | מקרוב |  |
| 7:7.192 | באו |  |
| 7:7.193 | ורוצים |  |
| 7:7.194 | להיות |  |
| 7:7.195 | מכת |  |
| 7:7.196 | המקובלים |  |
| 7:7.197 | וממדרשי |  |
| 7:7.198 | הנעלמים |  |
| 7:7.199 | ומחלושי |  |
| 7:7.200 | הראות | . |
| 7:7.201 | לא |  |
| 7:7.202 | יביטו |  |
| 7:7.203 | באור |  |
| 7:7.204 | הזהר |  |
| 7:7.205 | ולא |  |
| 7:7.206 | ידעו |  |
| 7:7.207 | מצאו |  |
| 7:7.208 | ומובאו |  |
| 7:7.209 | וכוונתו | . |
| 7:7.210 | אלא |  |
| 7:7.211 | שכך |  |
| 7:7.212 | מצאו |  |
| 7:7.213 | בספרי |  |
| 7:7.214 | הרשב"י |  |
| 7:7.215 | ודע |  |
| 7:7.216 | שכל |  |
| 7:7.217 | רבותי |  |
| 7:7.218 | הקדושים |  |
| 7:7.219 | ששמשו |  |
| 7:7.220 | גאוני |  |
| 7:7.221 | עולם |  |
| 7:7.222 | לא |  |
| 7:7.223 | נהגו |  |
| 7:7.224 | כך |  |
| 7:7.225 | אלא |  |
| 7:7.226 | כדברי |  |
| 7:7.227 | התלמוד |  |
| 7:7.228 | והפוסקים | . |
| 7:7.229 | ועתה |  |
| 7:7.230 | פקח |  |
| 7:7.231 | עיניך |  |
| 7:7.232 | וראה |  |
| 7:7.233 | אם |  |
| 7:7.234 | יש |  |
| 7:7.235 | בזה |  |
| 7:7.236 | ממש |  |
| 7:7.237 | אי |  |
| 7:7.238 | אפשר |  |
| 7:7.239 | שלא |  |
| 7:7.240 | הזכירו |  |
| 7:7.241 | שום |  |
| 7:7.242 | מחבר | . |
| 7:7.243 | או |  |
| 7:7.244 | גם |  |
| 7:7.245 | לא |  |
| 7:7.246 | אשתמיט |  |
| 7:7.247 | רמז |  |
| 7:7.248 | בתלמוד |  |
| 7:7.249 | בבלי |  |
| 7:7.250 | או |  |
| 7:7.251 | ירושלמי |  |
| 7:7.252 | ספרא |  |
| 7:7.253 | וספרי |  |
| 7:7.254 | ושאלתות |  |
| 7:7.255 | ופסיקתא | . |
| 7:7.256 | ומי |  |
| 7:7.257 | שלא |  |
| 7:7.258 | ידע |  |
| 7:7.259 | להשיג |  |
| 7:7.260 | סודה |  |
| 7:7.261 | על |  |
| 7:7.262 | נכון |  |
| 7:7.263 | יבוא |  |
| 7:7.264 | לקצץ |  |
| 7:7.265 | בנטיעות | . |
| 7:7.266 | לכך |  |
| 7:7.267 | אהובי |  |
| 7:7.268 | אל |  |
| 7:7.269 | תלך |  |
| 7:7.270 | בדרך |  |
| 7:7.271 | אתם | . |
| 7:7.272 | ואין |  |
| 7:7.273 | לך |  |
| 7:7.274 | עסק |  |
| 7:7.275 | בנסתרות |  |
| 7:7.276 | המתיהרי | ' |
| 7:7.277 | בחידושים | . |
| 7:7.278 | כאלו |  |
| 7:7.279 | הם |  |
| 7:7.280 | יודעים |  |
| 7:7.281 | ומבינים |  |
| 7:7.282 | רזי |  |
| 7:7.283 | תורה |  |
| 7:7.284 | וספוניה | . |
| 7:7.285 | והלואי |  |
| 7:7.286 | שידעו |  |
| 7:7.287 | הגלוים | . |
| 7:7.288 | אבל |  |
| 7:7.289 | דע |  |
| 7:7.290 | שכל |  |
| 7:7.291 | המשנה |  |
| 7:7.292 | ידו |  |
| 7:7.293 | על |  |
| 7:7.294 | התחתונה | . |
| 7:7.295 | וידינו |  |
| 7:7.296 | על |  |
| 7:7.297 | העליונה |  |
| 7:7.298 | עכ"ל | . |
| 7:7.299 | עוד |  |
| 7:7.300 | כ | ' |
| 7:7.301 | רביד |  |
| 7:7.302 | הזהב |  |
| 7:7.303 | בשמו |  |
| 7:7.304 | וז"ל |  |
| 7:7.305 | והמורה |  |
| 7:7.306 | היפך |  |
| 7:7.307 | התלמוד |  |
| 7:7.308 | חייב |  |
| 7:7.309 | מיתה | . |
| 7:7.310 | וכמו |  |
| 7:7.311 | שכ | ' |
| 7:7.312 | בס | ' |
| 7:7.313 | צח |  |
| 7:7.314 | ואדום |  |
| 7:7.315 | בשם |  |
| 7:7.316 | רש"ל |  |
| 7:7.317 | בתשו | ' |
| 7:7.318 | החדשות |  |
| 7:7.319 | וז"ל |  |
| 7:7.320 | שכל |  |
| 7:7.321 | המחמיר |  |
| 7:7.322 | שלא |  |
| 7:7.323 | לאכול |  |
| 7:7.324 | בשר |  |
| 7:7.325 | אחד |  |
| 7:7.326 | גבינה |  |
| 7:7.327 | בלא |  |
| 7:7.328 | גריעותא |  |
| 7:7.329 | דאית |  |
| 7:7.330 | ליה |  |
| 7:7.331 | בשינוהי | . |
| 7:7.332 | כי |  |
| 7:7.333 | המחמיר |  |
| 7:7.334 | חייב |  |
| 7:7.335 | מיתה |  |
| 7:7.336 | וכו | ' |
| 7:7.337 | ובס | ' |
| 7:7.338 | ים |  |
| 7:7.339 | שלמה |  |
| 7:7.340 | כ | ' |
| 7:7.341 | דמאחר |  |
| 7:7.342 | שנחתם |  |
| 7:7.343 | התלמוד |  |
| 7:7.344 | אין |  |
| 7:7.345 | להחמיר |  |
| 7:7.346 | נגד |  |
| 7:7.347 | התלמוד |  |
| 7:7.348 | והוי |  |
| 7:7.349 | כמו |  |
| 7:7.350 | מינות |  |
| 7:7.351 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

