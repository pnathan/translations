# Translation pack: milhamot-hashem perek-3-11 (Milchamot ha-Shem 3:17–3:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-11`.

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
- הוי"ה / הוי'ה / הויה (the Name as a name: Kabbalistic usage, and the
  author's own in 3:17) → "YHVH". Never "being". Examples: "ד' מילוי הויה ע"ב
  ס"ג מ"ה ב"ן" (2:33) → "the four fillings of YHVH, 72, 63, 45, 52" (in the
  note: these are the gematria names AB, SaG, MaH, BaN). "שהוא קורא לכל אחד
  מהם בשם הוי'ה ואדנות ואלהים" (8:24) → "who calls each of them by the names
  YHVH, Adonai and Elohim". אדנות as a name → "Adonai", with a note that the
  word is literally "lordship". עשר הויות (2:33) → "ten havayot", with a note.
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
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
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


## Milchamot ha-Shem 3:17

Hebrew (Wikisource): יז) ומן התורה שבכתב ושבעל פה הנמסרת לנו איש מפי איש שמשה רבינו ע"ה נביא ראיות חזקות ואמתיות שה' אלהינו הנקרא בשם הוי'ה ואדנות, ובדברי רז"ל בשם הקב"ה הוא הסבה הראשונה, והוא לבדו מחויב המציאות וכמו שיתבאר להלן שלא כמו שכ' בעל ספר הברית בשם יושר לבב ובשם האר"י שנמשכו אחרי דברי הפילוסוף מחבר הזהר ואמרו שהוא זעיר אנפין! מן התורה דכתיב מעונה אלהי קדם ענינו כי ה' אלהינו הוא הסבה הראשונה לכל ברואי מעלה ומטה והוא אשר לו לבדו נעבוד, כמו שכתב הרמב"ם ז"ל בפי' המשנה ובחיבור הרב סעדיה גאון ור' יהודה הלוי וזולתם. ונאמר אני ה' ראשון ומבלעדי אין אלהים. ודוד המלך ע"ה אמר כי מי אל מבלעדי ה' ומי צור מבלעדי אלהינו, כי הוא יתברך הסבה וההתחלה הפועלת לכל אשר זולתו כמו שביאר הרב המורה בפרק ט"ז מראשון עי"ש: אין סבה אחרת קודמת לו. וז"ל הר' החסיד בעל חובת הלבבות בשער היחוד, מן הדין שיהיה לעולם תחלה. אין תחלה לפניו, וראשון שאין לו ראשון, והוא אשר יצרו וחדשו מאין דבר, לא בדבר ולא על דבר כמו שנאמר אני ה' עשה כל נוטה שמים לבדי רוקע הארץ מי אתי. ואמר איוב נטה צפן על תהו תולה ארץ על בלי מה וכו'. הוא הקדמון הראשון אשר אין ראשית לראשיתו. ואין תכלית לקדמותו כמו שאמר הכתוב אני ראשון ואני אחרון ואומר אני ה' ראשון ואת אחרונים אני הוא עכ"ל. גם בתיקוני התפילות שתקנו אנשי כנה"ג נאמר אמת אתה הוא ראשון ואתה הוא אחרון וכו'.

| id | word | punctuation after |
|---|---|---|
| 3:17.1 | יז | ) |
| 3:17.2 | ומן |  |
| 3:17.3 | התורה |  |
| 3:17.4 | שבכתב |  |
| 3:17.5 | ושבעל |  |
| 3:17.6 | פה |  |
| 3:17.7 | הנמסרת |  |
| 3:17.8 | לנו |  |
| 3:17.9 | איש |  |
| 3:17.10 | מפי |  |
| 3:17.11 | איש |  |
| 3:17.12 | שמשה |  |
| 3:17.13 | רבינו |  |
| 3:17.14 | ע"ה |  |
| 3:17.15 | נביא |  |
| 3:17.16 | ראיות |  |
| 3:17.17 | חזקות |  |
| 3:17.18 | ואמתיות |  |
| 3:17.19 | שה | ' |
| 3:17.20 | אלהינו |  |
| 3:17.21 | הנקרא |  |
| 3:17.22 | בשם |  |
| 3:17.23 | הוי'ה |  |
| 3:17.24 | ואדנות | , |
| 3:17.25 | ובדברי |  |
| 3:17.26 | רז"ל |  |
| 3:17.27 | בשם |  |
| 3:17.28 | הקב"ה |  |
| 3:17.29 | הוא |  |
| 3:17.30 | הסבה |  |
| 3:17.31 | הראשונה | , |
| 3:17.32 | והוא |  |
| 3:17.33 | לבדו |  |
| 3:17.34 | מחויב |  |
| 3:17.35 | המציאות |  |
| 3:17.36 | וכמו |  |
| 3:17.37 | שיתבאר |  |
| 3:17.38 | להלן |  |
| 3:17.39 | שלא |  |
| 3:17.40 | כמו |  |
| 3:17.41 | שכ | ' |
| 3:17.42 | בעל |  |
| 3:17.43 | ספר |  |
| 3:17.44 | הברית |  |
| 3:17.45 | בשם |  |
| 3:17.46 | יושר |  |
| 3:17.47 | לבב |  |
| 3:17.48 | ובשם |  |
| 3:17.49 | האר"י |  |
| 3:17.50 | שנמשכו |  |
| 3:17.51 | אחרי |  |
| 3:17.52 | דברי |  |
| 3:17.53 | הפילוסוף |  |
| 3:17.54 | מחבר |  |
| 3:17.55 | הזהר |  |
| 3:17.56 | ואמרו |  |
| 3:17.57 | שהוא |  |
| 3:17.58 | זעיר |  |
| 3:17.59 | אנפין | ! |
| 3:17.60 | מן |  |
| 3:17.61 | התורה |  |
| 3:17.62 | דכתיב |  |
| 3:17.63 | מעונה |  |
| 3:17.64 | אלהי |  |
| 3:17.65 | קדם |  |
| 3:17.66 | ענינו |  |
| 3:17.67 | כי |  |
| 3:17.68 | ה | ' |
| 3:17.69 | אלהינו |  |
| 3:17.70 | הוא |  |
| 3:17.71 | הסבה |  |
| 3:17.72 | הראשונה |  |
| 3:17.73 | לכל |  |
| 3:17.74 | ברואי |  |
| 3:17.75 | מעלה |  |
| 3:17.76 | ומטה |  |
| 3:17.77 | והוא |  |
| 3:17.78 | אשר |  |
| 3:17.79 | לו |  |
| 3:17.80 | לבדו |  |
| 3:17.81 | נעבוד | , |
| 3:17.82 | כמו |  |
| 3:17.83 | שכתב |  |
| 3:17.84 | הרמב"ם |  |
| 3:17.85 | ז"ל |  |
| 3:17.86 | בפי | ' |
| 3:17.87 | המשנה |  |
| 3:17.88 | ובחיבור |  |
| 3:17.89 | הרב |  |
| 3:17.90 | סעדיה |  |
| 3:17.91 | גאון |  |
| 3:17.92 | ור | ' |
| 3:17.93 | יהודה |  |
| 3:17.94 | הלוי |  |
| 3:17.95 | וזולתם | . |
| 3:17.96 | ונאמר |  |
| 3:17.97 | אני |  |
| 3:17.98 | ה | ' |
| 3:17.99 | ראשון |  |
| 3:17.100 | ומבלעדי |  |
| 3:17.101 | אין |  |
| 3:17.102 | אלהים | . |
| 3:17.103 | ודוד |  |
| 3:17.104 | המלך |  |
| 3:17.105 | ע"ה |  |
| 3:17.106 | אמר |  |
| 3:17.107 | כי |  |
| 3:17.108 | מי |  |
| 3:17.109 | אל |  |
| 3:17.110 | מבלעדי |  |
| 3:17.111 | ה | ' |
| 3:17.112 | ומי |  |
| 3:17.113 | צור |  |
| 3:17.114 | מבלעדי |  |
| 3:17.115 | אלהינו | , |
| 3:17.116 | כי |  |
| 3:17.117 | הוא |  |
| 3:17.118 | יתברך |  |
| 3:17.119 | הסבה |  |
| 3:17.120 | וההתחלה |  |
| 3:17.121 | הפועלת |  |
| 3:17.122 | לכל |  |
| 3:17.123 | אשר |  |
| 3:17.124 | זולתו |  |
| 3:17.125 | כמו |  |
| 3:17.126 | שביאר |  |
| 3:17.127 | הרב |  |
| 3:17.128 | המורה |  |
| 3:17.129 | בפרק |  |
| 3:17.130 | ט"ז |  |
| 3:17.131 | מראשון |  |
| 3:17.132 | עי"ש | : |
| 3:17.133 | אין |  |
| 3:17.134 | סבה |  |
| 3:17.135 | אחרת |  |
| 3:17.136 | קודמת |  |
| 3:17.137 | לו | . |
| 3:17.138 | וז"ל |  |
| 3:17.139 | הר | ' |
| 3:17.140 | החסיד |  |
| 3:17.141 | בעל |  |
| 3:17.142 | חובת |  |
| 3:17.143 | הלבבות |  |
| 3:17.144 | בשער |  |
| 3:17.145 | היחוד | , |
| 3:17.146 | מן |  |
| 3:17.147 | הדין |  |
| 3:17.148 | שיהיה |  |
| 3:17.149 | לעולם |  |
| 3:17.150 | תחלה | . |
| 3:17.151 | אין |  |
| 3:17.152 | תחלה |  |
| 3:17.153 | לפניו | , |
| 3:17.154 | וראשון |  |
| 3:17.155 | שאין |  |
| 3:17.156 | לו |  |
| 3:17.157 | ראשון | , |
| 3:17.158 | והוא |  |
| 3:17.159 | אשר |  |
| 3:17.160 | יצרו |  |
| 3:17.161 | וחדשו |  |
| 3:17.162 | מאין |  |
| 3:17.163 | דבר | , |
| 3:17.164 | לא |  |
| 3:17.165 | בדבר |  |
| 3:17.166 | ולא |  |
| 3:17.167 | על |  |
| 3:17.168 | דבר |  |
| 3:17.169 | כמו |  |
| 3:17.170 | שנאמר |  |
| 3:17.171 | אני |  |
| 3:17.172 | ה | ' |
| 3:17.173 | עשה |  |
| 3:17.174 | כל |  |
| 3:17.175 | נוטה |  |
| 3:17.176 | שמים |  |
| 3:17.177 | לבדי |  |
| 3:17.178 | רוקע |  |
| 3:17.179 | הארץ |  |
| 3:17.180 | מי |  |
| 3:17.181 | אתי | . |
| 3:17.182 | ואמר |  |
| 3:17.183 | איוב |  |
| 3:17.184 | נטה |  |
| 3:17.185 | צפן |  |
| 3:17.186 | על |  |
| 3:17.187 | תהו |  |
| 3:17.188 | תולה |  |
| 3:17.189 | ארץ |  |
| 3:17.190 | על |  |
| 3:17.191 | בלי |  |
| 3:17.192 | מה |  |
| 3:17.193 | וכו | '. |
| 3:17.194 | הוא |  |
| 3:17.195 | הקדמון |  |
| 3:17.196 | הראשון |  |
| 3:17.197 | אשר |  |
| 3:17.198 | אין |  |
| 3:17.199 | ראשית |  |
| 3:17.200 | לראשיתו | . |
| 3:17.201 | ואין |  |
| 3:17.202 | תכלית |  |
| 3:17.203 | לקדמותו |  |
| 3:17.204 | כמו |  |
| 3:17.205 | שאמר |  |
| 3:17.206 | הכתוב |  |
| 3:17.207 | אני |  |
| 3:17.208 | ראשון |  |
| 3:17.209 | ואני |  |
| 3:17.210 | אחרון |  |
| 3:17.211 | ואומר |  |
| 3:17.212 | אני |  |
| 3:17.213 | ה | ' |
| 3:17.214 | ראשון |  |
| 3:17.215 | ואת |  |
| 3:17.216 | אחרונים |  |
| 3:17.217 | אני |  |
| 3:17.218 | הוא |  |
| 3:17.219 | עכ"ל | . |
| 3:17.220 | גם |  |
| 3:17.221 | בתיקוני |  |
| 3:17.222 | התפילות |  |
| 3:17.223 | שתקנו |  |
| 3:17.224 | אנשי |  |
| 3:17.225 | כנה"ג |  |
| 3:17.226 | נאמר |  |
| 3:17.227 | אמת |  |
| 3:17.228 | אתה |  |
| 3:17.229 | הוא |  |
| 3:17.230 | ראשון |  |
| 3:17.231 | ואתה |  |
| 3:17.232 | הוא |  |
| 3:17.233 | אחרון |  |
| 3:17.234 | וכו | '. |


## Milchamot ha-Shem 3:18

Hebrew (Wikisource): יח) ועוד ראיה שאינה קבלה ממשה שהרי מצינו מחלוקת בין המקובלים באיכותה ועצמותה זה אומר בכה וזה אומר בכה, וזו ראיה גמורה שאינה הלכה למשה מסיני כמו שעלה על דעתך, שהרי כתב הרמב"ם ז"ל בהל' ממרים פ"א שדברי קבלה אין בהם מחלוקת כלל. וכל דבר שתמצא בו מחלוקת בידוע שאינו קבלה ממשה רבי' ע"כ, וכיון שכל ראיות שאנו מביאים הם מדברי רז"ל בעלי המשנה והתלמוד והמדרשים האמתיים והגאונים והפוסקים שהם בודאי בעלי הקבלה האמתית ממשה רבי' ע"ה, וכל האומר היפך דבריהם הוא מין ואפיקורוס, ואם כן איך תוכל לומר מי נתן לך לדרוש ולחקור אחר דבריהם של בעלי הקבלה החדשה? וכבר העיד אחד מיוחד מחכמי ימי הבינים בארץ התימן הוא מהרי"י צאהרי בדרשות בפרשת לך לך, וז"ל ואני באתי להגיד לך מעט מזער דברי הקבלה חדשים מקרוב באו ע"כ. ועדותו זאת מועלת לנו לחזק דברינו, כי חדשות המה לא שערום רבותינו קדמונינו. ופליטת ספריהם יעידון יגידון האמת הלא המה מצויים בידינו במדרש הגדול לר' דוד בן עמרם איש עדן כמו שכ' מהר"י בשירי בפרשת בראשית ובתחלת אלה הדברים. וכ"כ בספר סגולת ישראל, וספר נור אל צאלם לר' נתנאל בן ישע ומדרש החפץ לר' יחיא הרופא זלה"ה.

| id | word | punctuation after |
|---|---|---|
| 3:18.1 | יח | ) |
| 3:18.2 | ועוד |  |
| 3:18.3 | ראיה |  |
| 3:18.4 | שאינה |  |
| 3:18.5 | קבלה |  |
| 3:18.6 | ממשה |  |
| 3:18.7 | שהרי |  |
| 3:18.8 | מצינו |  |
| 3:18.9 | מחלוקת |  |
| 3:18.10 | בין |  |
| 3:18.11 | המקובלים |  |
| 3:18.12 | באיכותה |  |
| 3:18.13 | ועצמותה |  |
| 3:18.14 | זה |  |
| 3:18.15 | אומר |  |
| 3:18.16 | בכה |  |
| 3:18.17 | וזה |  |
| 3:18.18 | אומר |  |
| 3:18.19 | בכה | , |
| 3:18.20 | וזו |  |
| 3:18.21 | ראיה |  |
| 3:18.22 | גמורה |  |
| 3:18.23 | שאינה |  |
| 3:18.24 | הלכה |  |
| 3:18.25 | למשה |  |
| 3:18.26 | מסיני |  |
| 3:18.27 | כמו |  |
| 3:18.28 | שעלה |  |
| 3:18.29 | על |  |
| 3:18.30 | דעתך | , |
| 3:18.31 | שהרי |  |
| 3:18.32 | כתב |  |
| 3:18.33 | הרמב"ם |  |
| 3:18.34 | ז"ל |  |
| 3:18.35 | בהל | ' |
| 3:18.36 | ממרים |  |
| 3:18.37 | פ"א |  |
| 3:18.38 | שדברי |  |
| 3:18.39 | קבלה |  |
| 3:18.40 | אין |  |
| 3:18.41 | בהם |  |
| 3:18.42 | מחלוקת |  |
| 3:18.43 | כלל | . |
| 3:18.44 | וכל |  |
| 3:18.45 | דבר |  |
| 3:18.46 | שתמצא |  |
| 3:18.47 | בו |  |
| 3:18.48 | מחלוקת |  |
| 3:18.49 | בידוע |  |
| 3:18.50 | שאינו |  |
| 3:18.51 | קבלה |  |
| 3:18.52 | ממשה |  |
| 3:18.53 | רבי | ' |
| 3:18.54 | ע"כ | , |
| 3:18.55 | וכיון |  |
| 3:18.56 | שכל |  |
| 3:18.57 | ראיות |  |
| 3:18.58 | שאנו |  |
| 3:18.59 | מביאים |  |
| 3:18.60 | הם |  |
| 3:18.61 | מדברי |  |
| 3:18.62 | רז"ל |  |
| 3:18.63 | בעלי |  |
| 3:18.64 | המשנה |  |
| 3:18.65 | והתלמוד |  |
| 3:18.66 | והמדרשים |  |
| 3:18.67 | האמתיים |  |
| 3:18.68 | והגאונים |  |
| 3:18.69 | והפוסקים |  |
| 3:18.70 | שהם |  |
| 3:18.71 | בודאי |  |
| 3:18.72 | בעלי |  |
| 3:18.73 | הקבלה |  |
| 3:18.74 | האמתית |  |
| 3:18.75 | ממשה |  |
| 3:18.76 | רבי | ' |
| 3:18.77 | ע"ה | , |
| 3:18.78 | וכל |  |
| 3:18.79 | האומר |  |
| 3:18.80 | היפך |  |
| 3:18.81 | דבריהם |  |
| 3:18.82 | הוא |  |
| 3:18.83 | מין |  |
| 3:18.84 | ואפיקורוס | , |
| 3:18.85 | ואם |  |
| 3:18.86 | כן |  |
| 3:18.87 | איך |  |
| 3:18.88 | תוכל |  |
| 3:18.89 | לומר |  |
| 3:18.90 | מי |  |
| 3:18.91 | נתן |  |
| 3:18.92 | לך |  |
| 3:18.93 | לדרוש |  |
| 3:18.94 | ולחקור |  |
| 3:18.95 | אחר |  |
| 3:18.96 | דבריהם |  |
| 3:18.97 | של |  |
| 3:18.98 | בעלי |  |
| 3:18.99 | הקבלה |  |
| 3:18.100 | החדשה | ? |
| 3:18.101 | וכבר |  |
| 3:18.102 | העיד |  |
| 3:18.103 | אחד |  |
| 3:18.104 | מיוחד |  |
| 3:18.105 | מחכמי |  |
| 3:18.106 | ימי |  |
| 3:18.107 | הבינים |  |
| 3:18.108 | בארץ |  |
| 3:18.109 | התימן |  |
| 3:18.110 | הוא |  |
| 3:18.111 | מהרי"י |  |
| 3:18.112 | צאהרי |  |
| 3:18.113 | בדרשות |  |
| 3:18.114 | בפרשת |  |
| 3:18.115 | לך |  |
| 3:18.116 | לך | , |
| 3:18.117 | וז"ל |  |
| 3:18.118 | ואני |  |
| 3:18.119 | באתי |  |
| 3:18.120 | להגיד |  |
| 3:18.121 | לך |  |
| 3:18.122 | מעט |  |
| 3:18.123 | מזער |  |
| 3:18.124 | דברי |  |
| 3:18.125 | הקבלה |  |
| 3:18.126 | חדשים |  |
| 3:18.127 | מקרוב |  |
| 3:18.128 | באו |  |
| 3:18.129 | ע"כ | . |
| 3:18.130 | ועדותו |  |
| 3:18.131 | זאת |  |
| 3:18.132 | מועלת |  |
| 3:18.133 | לנו |  |
| 3:18.134 | לחזק |  |
| 3:18.135 | דברינו | , |
| 3:18.136 | כי |  |
| 3:18.137 | חדשות |  |
| 3:18.138 | המה |  |
| 3:18.139 | לא |  |
| 3:18.140 | שערום |  |
| 3:18.141 | רבותינו |  |
| 3:18.142 | קדמונינו | . |
| 3:18.143 | ופליטת |  |
| 3:18.144 | ספריהם |  |
| 3:18.145 | יעידון |  |
| 3:18.146 | יגידון |  |
| 3:18.147 | האמת |  |
| 3:18.148 | הלא |  |
| 3:18.149 | המה |  |
| 3:18.150 | מצויים |  |
| 3:18.151 | בידינו |  |
| 3:18.152 | במדרש |  |
| 3:18.153 | הגדול |  |
| 3:18.154 | לר | ' |
| 3:18.155 | דוד |  |
| 3:18.156 | בן |  |
| 3:18.157 | עמרם |  |
| 3:18.158 | איש |  |
| 3:18.159 | עדן |  |
| 3:18.160 | כמו |  |
| 3:18.161 | שכ | ' |
| 3:18.162 | מהר"י |  |
| 3:18.163 | בשירי |  |
| 3:18.164 | בפרשת |  |
| 3:18.165 | בראשית |  |
| 3:18.166 | ובתחלת |  |
| 3:18.167 | אלה |  |
| 3:18.168 | הדברים | . |
| 3:18.169 | וכ"כ |  |
| 3:18.170 | בספר |  |
| 3:18.171 | סגולת |  |
| 3:18.172 | ישראל | , |
| 3:18.173 | וספר |  |
| 3:18.174 | נור |  |
| 3:18.175 | אל |  |
| 3:18.176 | צאלם |  |
| 3:18.177 | לר | ' |
| 3:18.178 | נתנאל |  |
| 3:18.179 | בן |  |
| 3:18.180 | ישע |  |
| 3:18.181 | ומדרש |  |
| 3:18.182 | החפץ |  |
| 3:18.183 | לר | ' |
| 3:18.184 | יחיא |  |
| 3:18.185 | הרופא |  |
| 3:18.186 | זלה"ה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

