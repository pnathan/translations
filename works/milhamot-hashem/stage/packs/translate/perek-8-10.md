# Translation pack: milhamot-hashem perek-8-10 (Milchamot ha-Shem 8:12–8:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-10`.

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
- **שיתוף / שתוף / לשתף** (there is no שותפות in the text): context.
  - Associating another with God: fixed, "associate / association". "כל
    המשתף שם שמים ודבר אחר נעקר מן העולם" (Sanh. 63a; 1:7, 5:29, 5:32) →
    "whoever associates the name of Heaven with something else is uprooted
    from the world"; "בשיתוף שאר הפרצופים" (5:1) → "in association with the
    other partzufim"; "בלי חיבור וצירוף ושיתוף" (5:21) → "without
    composition, conjunction or association".
  - The Maimonidean term for equivocal words: "בשתוף השם בלבד" (5:10) → "by
    equivocation of the term only"; "מושאלים או משותפים" (5:18) →
    "borrowed or equivocal". Never "association" here. The note gives
    Guide I.
  - צירוף: "conjunction" when used of God; "permutation" in letter mysticism
    ("י"ב צירופי הוי'ה", 2:33 → "the twelve permutations of YHVH").
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 8:12

Hebrew (Wikisource): קז) ולא זו בלבד, אלא גם אל ירושת"ו הרבה לעשות מכתבי און ומרמה ושקרים, עד להעלות אף וחמת החכמים בעיניהם, ויחרימו על פי מלשינותו את לומדי התורה ומלמדיה, ובמקהלות עם רב השמיע ופרסם את מכתב החרם של טיפשי ירושת"ו על ידי אחרים, שהלומדים משנה ותלמוד ופוסקים ואינם עוסקים בספר הזהר והתקונים לאקמא שכינתא מעפרא, הרי הם מוחרמים ומובדלים מחכמי ורבני ירושת"ו ואסור לאכול משחיטתם, גם אסף את כל הקצבים אל ביתו, ויזהירם וישביעם בה' אלהי ישראל לבל יזיד איש מהם לנקר להם בשר! ולילות עמל מנו לו, לחשוב על משכבו, ובבקר השכם יתיצב על פתחי בתי שופטי האימאם יר"ה להלשין את הלומדים, ולמנוע אותם לבל יתחברו עוד ללמוד. כי משחיתים הם את הדת ואת האמונה. ובחזרתו אל ביתו נקיט נפשיה בקצירי, ועל משענתו ילך הנה והנה, כאל אין בו כח לעמוד. כל זאת באתנו ולא שכחנו שם אלהינו האל הגדול הגבור והנורא, לפרוש כפינו לאל זר קצר אפים וזוגתו! אשר לא שערום אבותינו ואבות אבותינו, מיום גלותם אל ארץ התימן עד היום הזה, זולתם קצת חכמים בעיניהם שנמשכו אחר דעות מחבר הזהר כידוע. אך כיון שלא מחינו בכבוד שמו יתברך בהיות דברינו נשמעים עליה נתפשנו, כדגרסינן בפ"ק דע"ז ת"ר כשנתפס רבי אליעזר למינות העלוהו לגרדום לידון, אמר לו אותו הגמון זקן שכמותך יתעסק בדברים בטלים הללו, אמר לו נאמן עלי הדיין. כסבור אותו הגמון עליו הוא אומר. והוא לא אמר אלא כנגד אביו שבשמים. אמר לו האמנתי. דימוס פטור אתה. וכשבא אל ביתו נכנסו תלמידיו לנחמו, ולא קבל מהם תנחומין. אמר לו רבי עקיבה רבי הרשיני לומר לפניך דבר אחד ממה שלמדתני. אמר לו אמור. אמר לו שמא דבר מינות בא לידך והנאך הדבר ועליו נתפסת למינות. אמר לו עקיבה הזכרתני, פעם אחת הייתי מהלך בשוק העליון של ציפורי ומצאני אחד ויעקב איש כפר סכניא שמו. אמר לי כתוב בתורתכם לא תביא אתנן זוגה ומחיר כלב בית ה' אלהיך לכל נדר. מהו לעשות ממנו בית הכסא לכהן גדול, ולא אמרתי לו כלום. אמר לי כך לימדי (ישוע) כי מאתנן זוגה קבצה עד אתנן זוגה ישובו, והנאני הדבר, ועל ידי זה נתפסתי למינות ועברתי על מה שכתוב בתורה הרחק מעליה דרכך זו מינות, ואל תקרב אל פתח ביתה זו זנות ועד כמה אמר רב חסדא עד ארבע אמות וכו'. וכ' הרש"א הזהיר הכתוב להרחיק כדי שלא יהנה וימשך אחריה כדקאמר והנאני הדבר ועי"ז נתפסתי למינות ועברתי. דכששאל ממנו ולא השיב לו כלום היה לו להרחיק ולא היה לו לשמוע ממנו תשובתו עכ"ל. כדברים האלה שאירעו לר"א ע"ה אירעו לוי גם אנחנו על אשר לא גילינו את אמונת תורתינו הקדושה ביחוד השי"ת ועבודתו בלי שתוף וצירוף דבר אחר כאשר צונו, והיינו שומעים דברי מינות מאנשים שהם הולכים בעקבות הזהר והמקובלים החדשים העובדים את האלהים החדשים קצר האפים וזוגתו והוריו אבא ואימא וזקיניו אריך אנפין ועתיקא קדישא, ולא דברנו להם מאומה, ועל ידי זה נתפשנו למינות! ואלו לא בטחנו על ה' אלהינו האל הגדול הגבור והנורא בין יעשה גם בין לא יעשה, כי אז היו מוציאין אותנו בקולרין מבית הכלא להגלותינו אל ארץ אחרת, ארץ אוכלת יושביה על ידי מלשינות שקר. אמנם אלף אלפים תודות לאלהינו מלכינו שהוציא אותנו מבית הכלא לשלום פעם ראשונה ושנייה, והצילנו מן השלישית ולא מעדו רגלינו תלי"ת.

| id | word | punctuation after |
|---|---|---|
| 8:12.1 | קז | ) |
| 8:12.2 | ולא |  |
| 8:12.3 | זו |  |
| 8:12.4 | בלבד | , |
| 8:12.5 | אלא |  |
| 8:12.6 | גם |  |
| 8:12.7 | אל |  |
| 8:12.8 | ירושת"ו |  |
| 8:12.9 | הרבה |  |
| 8:12.10 | לעשות |  |
| 8:12.11 | מכתבי |  |
| 8:12.12 | און |  |
| 8:12.13 | ומרמה |  |
| 8:12.14 | ושקרים | , |
| 8:12.15 | עד |  |
| 8:12.16 | להעלות |  |
| 8:12.17 | אף |  |
| 8:12.18 | וחמת |  |
| 8:12.19 | החכמים |  |
| 8:12.20 | בעיניהם | , |
| 8:12.21 | ויחרימו |  |
| 8:12.22 | על |  |
| 8:12.23 | פי |  |
| 8:12.24 | מלשינותו |  |
| 8:12.25 | את |  |
| 8:12.26 | לומדי |  |
| 8:12.27 | התורה |  |
| 8:12.28 | ומלמדיה | , |
| 8:12.29 | ובמקהלות |  |
| 8:12.30 | עם |  |
| 8:12.31 | רב |  |
| 8:12.32 | השמיע |  |
| 8:12.33 | ופרסם |  |
| 8:12.34 | את |  |
| 8:12.35 | מכתב |  |
| 8:12.36 | החרם |  |
| 8:12.37 | של |  |
| 8:12.38 | טיפשי |  |
| 8:12.39 | ירושת"ו |  |
| 8:12.40 | על |  |
| 8:12.41 | ידי |  |
| 8:12.42 | אחרים | , |
| 8:12.43 | שהלומדים |  |
| 8:12.44 | משנה |  |
| 8:12.45 | ותלמוד |  |
| 8:12.46 | ופוסקים |  |
| 8:12.47 | ואינם |  |
| 8:12.48 | עוסקים |  |
| 8:12.49 | בספר |  |
| 8:12.50 | הזהר |  |
| 8:12.51 | והתקונים |  |
| 8:12.52 | לאקמא |  |
| 8:12.53 | שכינתא |  |
| 8:12.54 | מעפרא | , |
| 8:12.55 | הרי |  |
| 8:12.56 | הם |  |
| 8:12.57 | מוחרמים |  |
| 8:12.58 | ומובדלים |  |
| 8:12.59 | מחכמי |  |
| 8:12.60 | ורבני |  |
| 8:12.61 | ירושת"ו |  |
| 8:12.62 | ואסור |  |
| 8:12.63 | לאכול |  |
| 8:12.64 | משחיטתם | , |
| 8:12.65 | גם |  |
| 8:12.66 | אסף |  |
| 8:12.67 | את |  |
| 8:12.68 | כל |  |
| 8:12.69 | הקצבים |  |
| 8:12.70 | אל |  |
| 8:12.71 | ביתו | , |
| 8:12.72 | ויזהירם |  |
| 8:12.73 | וישביעם |  |
| 8:12.74 | בה | ' |
| 8:12.75 | אלהי |  |
| 8:12.76 | ישראל |  |
| 8:12.77 | לבל |  |
| 8:12.78 | יזיד |  |
| 8:12.79 | איש |  |
| 8:12.80 | מהם |  |
| 8:12.81 | לנקר |  |
| 8:12.82 | להם |  |
| 8:12.83 | בשר | ! |
| 8:12.84 | ולילות |  |
| 8:12.85 | עמל |  |
| 8:12.86 | מנו |  |
| 8:12.87 | לו | , |
| 8:12.88 | לחשוב |  |
| 8:12.89 | על |  |
| 8:12.90 | משכבו | , |
| 8:12.91 | ובבקר |  |
| 8:12.92 | השכם |  |
| 8:12.93 | יתיצב |  |
| 8:12.94 | על |  |
| 8:12.95 | פתחי |  |
| 8:12.96 | בתי |  |
| 8:12.97 | שופטי |  |
| 8:12.98 | האימאם |  |
| 8:12.99 | יר"ה |  |
| 8:12.100 | להלשין |  |
| 8:12.101 | את |  |
| 8:12.102 | הלומדים | , |
| 8:12.103 | ולמנוע |  |
| 8:12.104 | אותם |  |
| 8:12.105 | לבל |  |
| 8:12.106 | יתחברו |  |
| 8:12.107 | עוד |  |
| 8:12.108 | ללמוד | . |
| 8:12.109 | כי |  |
| 8:12.110 | משחיתים |  |
| 8:12.111 | הם |  |
| 8:12.112 | את |  |
| 8:12.113 | הדת |  |
| 8:12.114 | ואת |  |
| 8:12.115 | האמונה | . |
| 8:12.116 | ובחזרתו |  |
| 8:12.117 | אל |  |
| 8:12.118 | ביתו |  |
| 8:12.119 | נקיט |  |
| 8:12.120 | נפשיה |  |
| 8:12.121 | בקצירי | , |
| 8:12.122 | ועל |  |
| 8:12.123 | משענתו |  |
| 8:12.124 | ילך |  |
| 8:12.125 | הנה |  |
| 8:12.126 | והנה | , |
| 8:12.127 | כאל |  |
| 8:12.128 | אין |  |
| 8:12.129 | בו |  |
| 8:12.130 | כח |  |
| 8:12.131 | לעמוד | . |
| 8:12.132 | כל |  |
| 8:12.133 | זאת |  |
| 8:12.134 | באתנו |  |
| 8:12.135 | ולא |  |
| 8:12.136 | שכחנו |  |
| 8:12.137 | שם |  |
| 8:12.138 | אלהינו |  |
| 8:12.139 | האל |  |
| 8:12.140 | הגדול |  |
| 8:12.141 | הגבור |  |
| 8:12.142 | והנורא | , |
| 8:12.143 | לפרוש |  |
| 8:12.144 | כפינו |  |
| 8:12.145 | לאל |  |
| 8:12.146 | זר |  |
| 8:12.147 | קצר |  |
| 8:12.148 | אפים |  |
| 8:12.149 | וזוגתו | ! |
| 8:12.150 | אשר |  |
| 8:12.151 | לא |  |
| 8:12.152 | שערום |  |
| 8:12.153 | אבותינו |  |
| 8:12.154 | ואבות |  |
| 8:12.155 | אבותינו | , |
| 8:12.156 | מיום |  |
| 8:12.157 | גלותם |  |
| 8:12.158 | אל |  |
| 8:12.159 | ארץ |  |
| 8:12.160 | התימן |  |
| 8:12.161 | עד |  |
| 8:12.162 | היום |  |
| 8:12.163 | הזה | , |
| 8:12.164 | זולתם |  |
| 8:12.165 | קצת |  |
| 8:12.166 | חכמים |  |
| 8:12.167 | בעיניהם |  |
| 8:12.168 | שנמשכו |  |
| 8:12.169 | אחר |  |
| 8:12.170 | דעות |  |
| 8:12.171 | מחבר |  |
| 8:12.172 | הזהר |  |
| 8:12.173 | כידוע | . |
| 8:12.174 | אך |  |
| 8:12.175 | כיון |  |
| 8:12.176 | שלא |  |
| 8:12.177 | מחינו |  |
| 8:12.178 | בכבוד |  |
| 8:12.179 | שמו |  |
| 8:12.180 | יתברך |  |
| 8:12.181 | בהיות |  |
| 8:12.182 | דברינו |  |
| 8:12.183 | נשמעים |  |
| 8:12.184 | עליה |  |
| 8:12.185 | נתפשנו | , |
| 8:12.186 | כדגרסינן |  |
| 8:12.187 | בפ"ק |  |
| 8:12.188 | דע"ז |  |
| 8:12.189 | ת"ר |  |
| 8:12.190 | כשנתפס |  |
| 8:12.191 | רבי |  |
| 8:12.192 | אליעזר |  |
| 8:12.193 | למינות |  |
| 8:12.194 | העלוהו |  |
| 8:12.195 | לגרדום |  |
| 8:12.196 | לידון | , |
| 8:12.197 | אמר |  |
| 8:12.198 | לו |  |
| 8:12.199 | אותו |  |
| 8:12.200 | הגמון |  |
| 8:12.201 | זקן |  |
| 8:12.202 | שכמותך |  |
| 8:12.203 | יתעסק |  |
| 8:12.204 | בדברים |  |
| 8:12.205 | בטלים |  |
| 8:12.206 | הללו | , |
| 8:12.207 | אמר |  |
| 8:12.208 | לו |  |
| 8:12.209 | נאמן |  |
| 8:12.210 | עלי |  |
| 8:12.211 | הדיין | . |
| 8:12.212 | כסבור |  |
| 8:12.213 | אותו |  |
| 8:12.214 | הגמון |  |
| 8:12.215 | עליו |  |
| 8:12.216 | הוא |  |
| 8:12.217 | אומר | . |
| 8:12.218 | והוא |  |
| 8:12.219 | לא |  |
| 8:12.220 | אמר |  |
| 8:12.221 | אלא |  |
| 8:12.222 | כנגד |  |
| 8:12.223 | אביו |  |
| 8:12.224 | שבשמים | . |
| 8:12.225 | אמר |  |
| 8:12.226 | לו |  |
| 8:12.227 | האמנתי | . |
| 8:12.228 | דימוס |  |
| 8:12.229 | פטור |  |
| 8:12.230 | אתה | . |
| 8:12.231 | וכשבא |  |
| 8:12.232 | אל |  |
| 8:12.233 | ביתו |  |
| 8:12.234 | נכנסו |  |
| 8:12.235 | תלמידיו |  |
| 8:12.236 | לנחמו | , |
| 8:12.237 | ולא |  |
| 8:12.238 | קבל |  |
| 8:12.239 | מהם |  |
| 8:12.240 | תנחומין | . |
| 8:12.241 | אמר |  |
| 8:12.242 | לו |  |
| 8:12.243 | רבי |  |
| 8:12.244 | עקיבה |  |
| 8:12.245 | רבי |  |
| 8:12.246 | הרשיני |  |
| 8:12.247 | לומר |  |
| 8:12.248 | לפניך |  |
| 8:12.249 | דבר |  |
| 8:12.250 | אחד |  |
| 8:12.251 | ממה |  |
| 8:12.252 | שלמדתני | . |
| 8:12.253 | אמר |  |
| 8:12.254 | לו |  |
| 8:12.255 | אמור | . |
| 8:12.256 | אמר |  |
| 8:12.257 | לו |  |
| 8:12.258 | שמא |  |
| 8:12.259 | דבר |  |
| 8:12.260 | מינות |  |
| 8:12.261 | בא |  |
| 8:12.262 | לידך |  |
| 8:12.263 | והנאך |  |
| 8:12.264 | הדבר |  |
| 8:12.265 | ועליו |  |
| 8:12.266 | נתפסת |  |
| 8:12.267 | למינות | . |
| 8:12.268 | אמר |  |
| 8:12.269 | לו |  |
| 8:12.270 | עקיבה |  |
| 8:12.271 | הזכרתני | , |
| 8:12.272 | פעם |  |
| 8:12.273 | אחת |  |
| 8:12.274 | הייתי |  |
| 8:12.275 | מהלך |  |
| 8:12.276 | בשוק |  |
| 8:12.277 | העליון |  |
| 8:12.278 | של |  |
| 8:12.279 | ציפורי |  |
| 8:12.280 | ומצאני |  |
| 8:12.281 | אחד |  |
| 8:12.282 | ויעקב |  |
| 8:12.283 | איש |  |
| 8:12.284 | כפר |  |
| 8:12.285 | סכניא |  |
| 8:12.286 | שמו | . |
| 8:12.287 | אמר |  |
| 8:12.288 | לי |  |
| 8:12.289 | כתוב |  |
| 8:12.290 | בתורתכם |  |
| 8:12.291 | לא |  |
| 8:12.292 | תביא |  |
| 8:12.293 | אתנן |  |
| 8:12.294 | זוגה |  |
| 8:12.295 | ומחיר |  |
| 8:12.296 | כלב |  |
| 8:12.297 | בית |  |
| 8:12.298 | ה | ' |
| 8:12.299 | אלהיך |  |
| 8:12.300 | לכל |  |
| 8:12.301 | נדר | . |
| 8:12.302 | מהו |  |
| 8:12.303 | לעשות |  |
| 8:12.304 | ממנו |  |
| 8:12.305 | בית |  |
| 8:12.306 | הכסא |  |
| 8:12.307 | לכהן |  |
| 8:12.308 | גדול | , |
| 8:12.309 | ולא |  |
| 8:12.310 | אמרתי |  |
| 8:12.311 | לו |  |
| 8:12.312 | כלום | . |
| 8:12.313 | אמר |  |
| 8:12.314 | לי |  |
| 8:12.315 | כך |  |
| 8:12.316 | לימדי |  |
| 8:12.317 | ישוע | ) |
| 8:12.318 | כי |  |
| 8:12.319 | מאתנן |  |
| 8:12.320 | זוגה |  |
| 8:12.321 | קבצה |  |
| 8:12.322 | עד |  |
| 8:12.323 | אתנן |  |
| 8:12.324 | זוגה |  |
| 8:12.325 | ישובו | , |
| 8:12.326 | והנאני |  |
| 8:12.327 | הדבר | , |
| 8:12.328 | ועל |  |
| 8:12.329 | ידי |  |
| 8:12.330 | זה |  |
| 8:12.331 | נתפסתי |  |
| 8:12.332 | למינות |  |
| 8:12.333 | ועברתי |  |
| 8:12.334 | על |  |
| 8:12.335 | מה |  |
| 8:12.336 | שכתוב |  |
| 8:12.337 | בתורה |  |
| 8:12.338 | הרחק |  |
| 8:12.339 | מעליה |  |
| 8:12.340 | דרכך |  |
| 8:12.341 | זו |  |
| 8:12.342 | מינות | , |
| 8:12.343 | ואל |  |
| 8:12.344 | תקרב |  |
| 8:12.345 | אל |  |
| 8:12.346 | פתח |  |
| 8:12.347 | ביתה |  |
| 8:12.348 | זו |  |
| 8:12.349 | זנות |  |
| 8:12.350 | ועד |  |
| 8:12.351 | כמה |  |
| 8:12.352 | אמר |  |
| 8:12.353 | רב |  |
| 8:12.354 | חסדא |  |
| 8:12.355 | עד |  |
| 8:12.356 | ארבע |  |
| 8:12.357 | אמות |  |
| 8:12.358 | וכו | '. |
| 8:12.359 | וכ | ' |
| 8:12.360 | הרש"א |  |
| 8:12.361 | הזהיר |  |
| 8:12.362 | הכתוב |  |
| 8:12.363 | להרחיק |  |
| 8:12.364 | כדי |  |
| 8:12.365 | שלא |  |
| 8:12.366 | יהנה |  |
| 8:12.367 | וימשך |  |
| 8:12.368 | אחריה |  |
| 8:12.369 | כדקאמר |  |
| 8:12.370 | והנאני |  |
| 8:12.371 | הדבר |  |
| 8:12.372 | ועי"ז |  |
| 8:12.373 | נתפסתי |  |
| 8:12.374 | למינות |  |
| 8:12.375 | ועברתי | . |
| 8:12.376 | דכששאל |  |
| 8:12.377 | ממנו |  |
| 8:12.378 | ולא |  |
| 8:12.379 | השיב |  |
| 8:12.380 | לו |  |
| 8:12.381 | כלום |  |
| 8:12.382 | היה |  |
| 8:12.383 | לו |  |
| 8:12.384 | להרחיק |  |
| 8:12.385 | ולא |  |
| 8:12.386 | היה |  |
| 8:12.387 | לו |  |
| 8:12.388 | לשמוע |  |
| 8:12.389 | ממנו |  |
| 8:12.390 | תשובתו |  |
| 8:12.391 | עכ"ל | . |
| 8:12.392 | כדברים |  |
| 8:12.393 | האלה |  |
| 8:12.394 | שאירעו |  |
| 8:12.395 | לר"א |  |
| 8:12.396 | ע"ה |  |
| 8:12.397 | אירעו |  |
| 8:12.398 | לוי |  |
| 8:12.399 | גם |  |
| 8:12.400 | אנחנו |  |
| 8:12.401 | על |  |
| 8:12.402 | אשר |  |
| 8:12.403 | לא |  |
| 8:12.404 | גילינו |  |
| 8:12.405 | את |  |
| 8:12.406 | אמונת |  |
| 8:12.407 | תורתינו |  |
| 8:12.408 | הקדושה |  |
| 8:12.409 | ביחוד |  |
| 8:12.410 | השי"ת |  |
| 8:12.411 | ועבודתו |  |
| 8:12.412 | בלי |  |
| 8:12.413 | שתוף |  |
| 8:12.414 | וצירוף |  |
| 8:12.415 | דבר |  |
| 8:12.416 | אחר |  |
| 8:12.417 | כאשר |  |
| 8:12.418 | צונו | , |
| 8:12.419 | והיינו |  |
| 8:12.420 | שומעים |  |
| 8:12.421 | דברי |  |
| 8:12.422 | מינות |  |
| 8:12.423 | מאנשים |  |
| 8:12.424 | שהם |  |
| 8:12.425 | הולכים |  |
| 8:12.426 | בעקבות |  |
| 8:12.427 | הזהר |  |
| 8:12.428 | והמקובלים |  |
| 8:12.429 | החדשים |  |
| 8:12.430 | העובדים |  |
| 8:12.431 | את |  |
| 8:12.432 | האלהים |  |
| 8:12.433 | החדשים |  |
| 8:12.434 | קצר |  |
| 8:12.435 | האפים |  |
| 8:12.436 | וזוגתו |  |
| 8:12.437 | והוריו |  |
| 8:12.438 | אבא |  |
| 8:12.439 | ואימא |  |
| 8:12.440 | וזקיניו |  |
| 8:12.441 | אריך |  |
| 8:12.442 | אנפין |  |
| 8:12.443 | ועתיקא |  |
| 8:12.444 | קדישא | , |
| 8:12.445 | ולא |  |
| 8:12.446 | דברנו |  |
| 8:12.447 | להם |  |
| 8:12.448 | מאומה | , |
| 8:12.449 | ועל |  |
| 8:12.450 | ידי |  |
| 8:12.451 | זה |  |
| 8:12.452 | נתפשנו |  |
| 8:12.453 | למינות | ! |
| 8:12.454 | ואלו |  |
| 8:12.455 | לא |  |
| 8:12.456 | בטחנו |  |
| 8:12.457 | על |  |
| 8:12.458 | ה | ' |
| 8:12.459 | אלהינו |  |
| 8:12.460 | האל |  |
| 8:12.461 | הגדול |  |
| 8:12.462 | הגבור |  |
| 8:12.463 | והנורא |  |
| 8:12.464 | בין |  |
| 8:12.465 | יעשה |  |
| 8:12.466 | גם |  |
| 8:12.467 | בין |  |
| 8:12.468 | לא |  |
| 8:12.469 | יעשה | , |
| 8:12.470 | כי |  |
| 8:12.471 | אז |  |
| 8:12.472 | היו |  |
| 8:12.473 | מוציאין |  |
| 8:12.474 | אותנו |  |
| 8:12.475 | בקולרין |  |
| 8:12.476 | מבית |  |
| 8:12.477 | הכלא |  |
| 8:12.478 | להגלותינו |  |
| 8:12.479 | אל |  |
| 8:12.480 | ארץ |  |
| 8:12.481 | אחרת | , |
| 8:12.482 | ארץ |  |
| 8:12.483 | אוכלת |  |
| 8:12.484 | יושביה |  |
| 8:12.485 | על |  |
| 8:12.486 | ידי |  |
| 8:12.487 | מלשינות |  |
| 8:12.488 | שקר | . |
| 8:12.489 | אמנם |  |
| 8:12.490 | אלף |  |
| 8:12.491 | אלפים |  |
| 8:12.492 | תודות |  |
| 8:12.493 | לאלהינו |  |
| 8:12.494 | מלכינו |  |
| 8:12.495 | שהוציא |  |
| 8:12.496 | אותנו |  |
| 8:12.497 | מבית |  |
| 8:12.498 | הכלא |  |
| 8:12.499 | לשלום |  |
| 8:12.500 | פעם |  |
| 8:12.501 | ראשונה |  |
| 8:12.502 | ושנייה | , |
| 8:12.503 | והצילנו |  |
| 8:12.504 | מן |  |
| 8:12.505 | השלישית |  |
| 8:12.506 | ולא |  |
| 8:12.507 | מעדו |  |
| 8:12.508 | רגלינו |  |
| 8:12.509 | תלי"ת | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

