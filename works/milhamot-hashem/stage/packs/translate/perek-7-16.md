# Translation pack: milhamot-hashem perek-7-16 (Milchamot ha-Shem 7:21–7:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-16`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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


## Milchamot ha-Shem 7:21

Hebrew (Wikisource): צג) מכל זה למדנו גנות המספר בעניינו המשגל ונבלותו ועונשו החמור אף אם מספר על אדם פחות ומזולזל. או על עצמו איך הוא נזקק לאשתו. ואימתי הוא נזקק לה. וכיצד נזקק לה ואיך הוא מחבקה ומנשקה, כ"ש וק"ו אם מספר ומגלה קלון מלך בשר ודם אימתי משמש ונזקק לאשתו ואיך הוא מחבק לה ומנשקה, בודאי שהמספר ומגלה דברים כאלה, גנאי גדול הוא למלך, וכתוב וערות מלכא לא אריך נא למחזי, ופי' רש"י והראב"ע ז"ל וערות מלכא בזיון המלך וגנותו, לא אריך לנא למחזי, אין הגון לנו לראות ע"כ. ותא חזי מאי סלקא ביה בחם אבי כנען על ידי שראה ערות אביו וקלונו והגיד לשני אחיו שם ויפת, והדברים קל וחומר, ומה זה שספר לאחיו גנות אביו כך, המספר במלכו של עולם עאכ"ו. ובילקוט רמז ס"א, ויתגל בתוך אהלה, אהלה כתיב, בתוך אהלה של אשת נח, כשיצא נח מן התבה הכישו הארי ושברו, ובא לו לשמש מטתו ונתפזר זרעו, וירא חם אבי כנען ויגד, אמר להון ואגיד להון. אמר לון, אדם הראשון היו לו שני בנים ועמד אחד והרג את חבירו, וזה יש לו שלשה בנים, והוא מבקש לעשות ארבעה? מה טעם העבד יוצא בשן ועין, מן הכא וירא ויגד, ויקח שם ויפת את השמלה, שם התחיל במצוה, ובא יפת ונשמע לו לפי' זכה שם לטלית ויפת לפלוניא, ממשמע שנ' וילכו אחורנית איני יודע שערות אביהם לא ראו, אלא מלמד נתנו ידיהם על פניהם והיו מהלכין לאחריהם ונהגו בו כבוד כמורא האב על הבן, אמר הקב"ה לשם, אתה כסית את ערות אביך. חייך שאני פורעה לבניך, שנ' באדין גובריא אליך כפיתו בסרבליהון, אמר הקב"ה ליפת אתה כסית ערות אביך חייך שאני פורע לך שנ' ביום ההוא אתן לגוג מקום קבר בישראל. אמר הקב"ה להם אתה בזית עריתא דאבוך, חייך שאני פורע לך שנ' כן ינהג מלך אשור את שבי מצרים ואת גלות כוש נערים וזקנים ערום ויחף וחשופי שת ערות מצרים עכ"ל. ראה תראה ותדין ק"ו כמו שכתבתי לעיל, ובודאי שהרואה בזיון המלך זגגותו ואינו מוחה, וכ"ש המספר בגנות המלך מלכו של עולם. ומיחס לו עניני תשמיש וזיווג ולידה, עאכ"ו שהוא נעקר מן העה"ז ומן העה"ב. ואין לו חלק באלהי ישראל. כי איככה ניחס לאלוהינו אשר אותו נעבוד פילגש ואיברי משגל. יסוד דהינו גיד הערוה וביצי זכר לבשל הזרע. ונקב אחוריים להוציא הפסולת להיות ניזון ממנה אל אחר? הלא נבוש ונכלם לספר עניינים כאלה בבשר ודם כי אם ברמז לעת הצורך כגון שיש מחלה באיברי המשגל או בפי הטבעת, ועם כל זה אין מזכירים אותם אלא ברמז וכינוי ובהצנע ולא בפרסום אבל שלא לצורך הם מלהזכיר כלל. כי זהו גבלות הפה וגנות גדולה שלמדה אותנו תוה"ק להשמר ממנו כמאמר רז"ל בפ"ק דפסחים אמר ר' יהושע בן לוי לעולם אל יוציא אדם דבר מגונה מפיו שהרי עיקם הכ' שמנה אותיות ולא הוציא דבר מגונה מפיו שנ' ומן הבהמה אשר איננה טהורה רב אשי אמר תשע שנ' כי יהיה בך איש אשר לא יהיה טהור מקרה לילה. רבינא אמר עשר וי"ו דטהור. רב אמר שש עשרה שנ' כי אמר מקרה הוא בלתי טהור הוא כי לא טהור. תנא דבי ר' ישמעאל לעולם יספר אדם בלשון נקייה שהרי בזב וכו'. ואומר ותבחר לשון ערומים. ואומר וזעת שפתי ברור מללו וכו'. ובכל התורה ובדברי הנביאים ורז"ל בתלמוד ובמדרשים לא מצינו שכינו להשפעת השי"ת בשם זכר ונקבה וזיווג אלא בלשון נתינה או פתיחה ויתן לך האלהים. נותן לך כח לעשות חיל. יפתח ה' לך את אוצרו הטוב נותן לחם לכל בשר. פותח ידיך ומשביע לכל חי רצון. ואתה מחיה את כלם וכיוצא בלשונות אלו אן מספר. ובדברי רז"ל יושב הקב"ה וזן את העולם כו'. זן ומפרניס לכל. עורך שולחן, נותן טרף וכיוצא. ולמה זה נעזוב מליצת ספרי הקדש שדברו בלשון נקייה. ונשתמש בשלון מגונה סכרים ונקבות רבות וזווגים והטפת טיפי הזרע ממבושי הזכר לרחמה של נקבה וקליטה וקפאון הזרע ברחמה כמפורסם כל זה בספרי הקבלה החדשה הנכריה. ולא די כל זה אלא שהמסית מחבר הזהר העיז פניו לכנות המבושים אשר חשב באדון העולמים בשם בית אלהים. ושער השמים, כדאיתא בזהר ויצא דף ק"ן ע"ב.

| id | word | punctuation after |
|---|---|---|
| 7:21.1 | צג | ) |
| 7:21.2 | מכל |  |
| 7:21.3 | זה |  |
| 7:21.4 | למדנו |  |
| 7:21.5 | גנות |  |
| 7:21.6 | המספר |  |
| 7:21.7 | בעניינו |  |
| 7:21.8 | המשגל |  |
| 7:21.9 | ונבלותו |  |
| 7:21.10 | ועונשו |  |
| 7:21.11 | החמור |  |
| 7:21.12 | אף |  |
| 7:21.13 | אם |  |
| 7:21.14 | מספר |  |
| 7:21.15 | על |  |
| 7:21.16 | אדם |  |
| 7:21.17 | פחות |  |
| 7:21.18 | ומזולזל | . |
| 7:21.19 | או |  |
| 7:21.20 | על |  |
| 7:21.21 | עצמו |  |
| 7:21.22 | איך |  |
| 7:21.23 | הוא |  |
| 7:21.24 | נזקק |  |
| 7:21.25 | לאשתו | . |
| 7:21.26 | ואימתי |  |
| 7:21.27 | הוא |  |
| 7:21.28 | נזקק |  |
| 7:21.29 | לה | . |
| 7:21.30 | וכיצד |  |
| 7:21.31 | נזקק |  |
| 7:21.32 | לה |  |
| 7:21.33 | ואיך |  |
| 7:21.34 | הוא |  |
| 7:21.35 | מחבקה |  |
| 7:21.36 | ומנשקה | , |
| 7:21.37 | כ"ש |  |
| 7:21.38 | וק"ו |  |
| 7:21.39 | אם |  |
| 7:21.40 | מספר |  |
| 7:21.41 | ומגלה |  |
| 7:21.42 | קלון |  |
| 7:21.43 | מלך |  |
| 7:21.44 | בשר |  |
| 7:21.45 | ודם |  |
| 7:21.46 | אימתי |  |
| 7:21.47 | משמש |  |
| 7:21.48 | ונזקק |  |
| 7:21.49 | לאשתו |  |
| 7:21.50 | ואיך |  |
| 7:21.51 | הוא |  |
| 7:21.52 | מחבק |  |
| 7:21.53 | לה |  |
| 7:21.54 | ומנשקה | , |
| 7:21.55 | בודאי |  |
| 7:21.56 | שהמספר |  |
| 7:21.57 | ומגלה |  |
| 7:21.58 | דברים |  |
| 7:21.59 | כאלה | , |
| 7:21.60 | גנאי |  |
| 7:21.61 | גדול |  |
| 7:21.62 | הוא |  |
| 7:21.63 | למלך | , |
| 7:21.64 | וכתוב |  |
| 7:21.65 | וערות |  |
| 7:21.66 | מלכא |  |
| 7:21.67 | לא |  |
| 7:21.68 | אריך |  |
| 7:21.69 | נא |  |
| 7:21.70 | למחזי | , |
| 7:21.71 | ופי | ' |
| 7:21.72 | רש"י |  |
| 7:21.73 | והראב"ע |  |
| 7:21.74 | ז"ל |  |
| 7:21.75 | וערות |  |
| 7:21.76 | מלכא |  |
| 7:21.77 | בזיון |  |
| 7:21.78 | המלך |  |
| 7:21.79 | וגנותו | , |
| 7:21.80 | לא |  |
| 7:21.81 | אריך |  |
| 7:21.82 | לנא |  |
| 7:21.83 | למחזי | , |
| 7:21.84 | אין |  |
| 7:21.85 | הגון |  |
| 7:21.86 | לנו |  |
| 7:21.87 | לראות |  |
| 7:21.88 | ע"כ | . |
| 7:21.89 | ותא |  |
| 7:21.90 | חזי |  |
| 7:21.91 | מאי |  |
| 7:21.92 | סלקא |  |
| 7:21.93 | ביה |  |
| 7:21.94 | בחם |  |
| 7:21.95 | אבי |  |
| 7:21.96 | כנען |  |
| 7:21.97 | על |  |
| 7:21.98 | ידי |  |
| 7:21.99 | שראה |  |
| 7:21.100 | ערות |  |
| 7:21.101 | אביו |  |
| 7:21.102 | וקלונו |  |
| 7:21.103 | והגיד |  |
| 7:21.104 | לשני |  |
| 7:21.105 | אחיו |  |
| 7:21.106 | שם |  |
| 7:21.107 | ויפת | , |
| 7:21.108 | והדברים |  |
| 7:21.109 | קל |  |
| 7:21.110 | וחומר | , |
| 7:21.111 | ומה |  |
| 7:21.112 | זה |  |
| 7:21.113 | שספר |  |
| 7:21.114 | לאחיו |  |
| 7:21.115 | גנות |  |
| 7:21.116 | אביו |  |
| 7:21.117 | כך | , |
| 7:21.118 | המספר |  |
| 7:21.119 | במלכו |  |
| 7:21.120 | של |  |
| 7:21.121 | עולם |  |
| 7:21.122 | עאכ"ו | . |
| 7:21.123 | ובילקוט |  |
| 7:21.124 | רמז |  |
| 7:21.125 | ס"א | , |
| 7:21.126 | ויתגל |  |
| 7:21.127 | בתוך |  |
| 7:21.128 | אהלה | , |
| 7:21.129 | אהלה |  |
| 7:21.130 | כתיב | , |
| 7:21.131 | בתוך |  |
| 7:21.132 | אהלה |  |
| 7:21.133 | של |  |
| 7:21.134 | אשת |  |
| 7:21.135 | נח | , |
| 7:21.136 | כשיצא |  |
| 7:21.137 | נח |  |
| 7:21.138 | מן |  |
| 7:21.139 | התבה |  |
| 7:21.140 | הכישו |  |
| 7:21.141 | הארי |  |
| 7:21.142 | ושברו | , |
| 7:21.143 | ובא |  |
| 7:21.144 | לו |  |
| 7:21.145 | לשמש |  |
| 7:21.146 | מטתו |  |
| 7:21.147 | ונתפזר |  |
| 7:21.148 | זרעו | , |
| 7:21.149 | וירא |  |
| 7:21.150 | חם |  |
| 7:21.151 | אבי |  |
| 7:21.152 | כנען |  |
| 7:21.153 | ויגד | , |
| 7:21.154 | אמר |  |
| 7:21.155 | להון |  |
| 7:21.156 | ואגיד |  |
| 7:21.157 | להון | . |
| 7:21.158 | אמר |  |
| 7:21.159 | לון | , |
| 7:21.160 | אדם |  |
| 7:21.161 | הראשון |  |
| 7:21.162 | היו |  |
| 7:21.163 | לו |  |
| 7:21.164 | שני |  |
| 7:21.165 | בנים |  |
| 7:21.166 | ועמד |  |
| 7:21.167 | אחד |  |
| 7:21.168 | והרג |  |
| 7:21.169 | את |  |
| 7:21.170 | חבירו | , |
| 7:21.171 | וזה |  |
| 7:21.172 | יש |  |
| 7:21.173 | לו |  |
| 7:21.174 | שלשה |  |
| 7:21.175 | בנים | , |
| 7:21.176 | והוא |  |
| 7:21.177 | מבקש |  |
| 7:21.178 | לעשות |  |
| 7:21.179 | ארבעה | ? |
| 7:21.180 | מה |  |
| 7:21.181 | טעם |  |
| 7:21.182 | העבד |  |
| 7:21.183 | יוצא |  |
| 7:21.184 | בשן |  |
| 7:21.185 | ועין | , |
| 7:21.186 | מן |  |
| 7:21.187 | הכא |  |
| 7:21.188 | וירא |  |
| 7:21.189 | ויגד | , |
| 7:21.190 | ויקח |  |
| 7:21.191 | שם |  |
| 7:21.192 | ויפת |  |
| 7:21.193 | את |  |
| 7:21.194 | השמלה | , |
| 7:21.195 | שם |  |
| 7:21.196 | התחיל |  |
| 7:21.197 | במצוה | , |
| 7:21.198 | ובא |  |
| 7:21.199 | יפת |  |
| 7:21.200 | ונשמע |  |
| 7:21.201 | לו |  |
| 7:21.202 | לפי | ' |
| 7:21.203 | זכה |  |
| 7:21.204 | שם |  |
| 7:21.205 | לטלית |  |
| 7:21.206 | ויפת |  |
| 7:21.207 | לפלוניא | , |
| 7:21.208 | ממשמע |  |
| 7:21.209 | שנ | ' |
| 7:21.210 | וילכו |  |
| 7:21.211 | אחורנית |  |
| 7:21.212 | איני |  |
| 7:21.213 | יודע |  |
| 7:21.214 | שערות |  |
| 7:21.215 | אביהם |  |
| 7:21.216 | לא |  |
| 7:21.217 | ראו | , |
| 7:21.218 | אלא |  |
| 7:21.219 | מלמד |  |
| 7:21.220 | נתנו |  |
| 7:21.221 | ידיהם |  |
| 7:21.222 | על |  |
| 7:21.223 | פניהם |  |
| 7:21.224 | והיו |  |
| 7:21.225 | מהלכין |  |
| 7:21.226 | לאחריהם |  |
| 7:21.227 | ונהגו |  |
| 7:21.228 | בו |  |
| 7:21.229 | כבוד |  |
| 7:21.230 | כמורא |  |
| 7:21.231 | האב |  |
| 7:21.232 | על |  |
| 7:21.233 | הבן | , |
| 7:21.234 | אמר |  |
| 7:21.235 | הקב"ה |  |
| 7:21.236 | לשם | , |
| 7:21.237 | אתה |  |
| 7:21.238 | כסית |  |
| 7:21.239 | את |  |
| 7:21.240 | ערות |  |
| 7:21.241 | אביך | . |
| 7:21.242 | חייך |  |
| 7:21.243 | שאני |  |
| 7:21.244 | פורעה |  |
| 7:21.245 | לבניך | , |
| 7:21.246 | שנ | ' |
| 7:21.247 | באדין |  |
| 7:21.248 | גובריא |  |
| 7:21.249 | אליך |  |
| 7:21.250 | כפיתו |  |
| 7:21.251 | בסרבליהון | , |
| 7:21.252 | אמר |  |
| 7:21.253 | הקב"ה |  |
| 7:21.254 | ליפת |  |
| 7:21.255 | אתה |  |
| 7:21.256 | כסית |  |
| 7:21.257 | ערות |  |
| 7:21.258 | אביך |  |
| 7:21.259 | חייך |  |
| 7:21.260 | שאני |  |
| 7:21.261 | פורע |  |
| 7:21.262 | לך |  |
| 7:21.263 | שנ | ' |
| 7:21.264 | ביום |  |
| 7:21.265 | ההוא |  |
| 7:21.266 | אתן |  |
| 7:21.267 | לגוג |  |
| 7:21.268 | מקום |  |
| 7:21.269 | קבר |  |
| 7:21.270 | בישראל | . |
| 7:21.271 | אמר |  |
| 7:21.272 | הקב"ה |  |
| 7:21.273 | להם |  |
| 7:21.274 | אתה |  |
| 7:21.275 | בזית |  |
| 7:21.276 | עריתא |  |
| 7:21.277 | דאבוך | , |
| 7:21.278 | חייך |  |
| 7:21.279 | שאני |  |
| 7:21.280 | פורע |  |
| 7:21.281 | לך |  |
| 7:21.282 | שנ | ' |
| 7:21.283 | כן |  |
| 7:21.284 | ינהג |  |
| 7:21.285 | מלך |  |
| 7:21.286 | אשור |  |
| 7:21.287 | את |  |
| 7:21.288 | שבי |  |
| 7:21.289 | מצרים |  |
| 7:21.290 | ואת |  |
| 7:21.291 | גלות |  |
| 7:21.292 | כוש |  |
| 7:21.293 | נערים |  |
| 7:21.294 | וזקנים |  |
| 7:21.295 | ערום |  |
| 7:21.296 | ויחף |  |
| 7:21.297 | וחשופי |  |
| 7:21.298 | שת |  |
| 7:21.299 | ערות |  |
| 7:21.300 | מצרים |  |
| 7:21.301 | עכ"ל | . |
| 7:21.302 | ראה |  |
| 7:21.303 | תראה |  |
| 7:21.304 | ותדין |  |
| 7:21.305 | ק"ו |  |
| 7:21.306 | כמו |  |
| 7:21.307 | שכתבתי |  |
| 7:21.308 | לעיל | , |
| 7:21.309 | ובודאי |  |
| 7:21.310 | שהרואה |  |
| 7:21.311 | בזיון |  |
| 7:21.312 | המלך |  |
| 7:21.313 | זגגותו |  |
| 7:21.314 | ואינו |  |
| 7:21.315 | מוחה | , |
| 7:21.316 | וכ"ש |  |
| 7:21.317 | המספר |  |
| 7:21.318 | בגנות |  |
| 7:21.319 | המלך |  |
| 7:21.320 | מלכו |  |
| 7:21.321 | של |  |
| 7:21.322 | עולם | . |
| 7:21.323 | ומיחס |  |
| 7:21.324 | לו |  |
| 7:21.325 | עניני |  |
| 7:21.326 | תשמיש |  |
| 7:21.327 | וזיווג |  |
| 7:21.328 | ולידה | , |
| 7:21.329 | עאכ"ו |  |
| 7:21.330 | שהוא |  |
| 7:21.331 | נעקר |  |
| 7:21.332 | מן |  |
| 7:21.333 | העה"ז |  |
| 7:21.334 | ומן |  |
| 7:21.335 | העה"ב | . |
| 7:21.336 | ואין |  |
| 7:21.337 | לו |  |
| 7:21.338 | חלק |  |
| 7:21.339 | באלהי |  |
| 7:21.340 | ישראל | . |
| 7:21.341 | כי |  |
| 7:21.342 | איככה |  |
| 7:21.343 | ניחס |  |
| 7:21.344 | לאלוהינו |  |
| 7:21.345 | אשר |  |
| 7:21.346 | אותו |  |
| 7:21.347 | נעבוד |  |
| 7:21.348 | פילגש |  |
| 7:21.349 | ואיברי |  |
| 7:21.350 | משגל | . |
| 7:21.351 | יסוד |  |
| 7:21.352 | דהינו |  |
| 7:21.353 | גיד |  |
| 7:21.354 | הערוה |  |
| 7:21.355 | וביצי |  |
| 7:21.356 | זכר |  |
| 7:21.357 | לבשל |  |
| 7:21.358 | הזרע | . |
| 7:21.359 | ונקב |  |
| 7:21.360 | אחוריים |  |
| 7:21.361 | להוציא |  |
| 7:21.362 | הפסולת |  |
| 7:21.363 | להיות |  |
| 7:21.364 | ניזון |  |
| 7:21.365 | ממנה |  |
| 7:21.366 | אל |  |
| 7:21.367 | אחר | ? |
| 7:21.368 | הלא |  |
| 7:21.369 | נבוש |  |
| 7:21.370 | ונכלם |  |
| 7:21.371 | לספר |  |
| 7:21.372 | עניינים |  |
| 7:21.373 | כאלה |  |
| 7:21.374 | בבשר |  |
| 7:21.375 | ודם |  |
| 7:21.376 | כי |  |
| 7:21.377 | אם |  |
| 7:21.378 | ברמז |  |
| 7:21.379 | לעת |  |
| 7:21.380 | הצורך |  |
| 7:21.381 | כגון |  |
| 7:21.382 | שיש |  |
| 7:21.383 | מחלה |  |
| 7:21.384 | באיברי |  |
| 7:21.385 | המשגל |  |
| 7:21.386 | או |  |
| 7:21.387 | בפי |  |
| 7:21.388 | הטבעת | , |
| 7:21.389 | ועם |  |
| 7:21.390 | כל |  |
| 7:21.391 | זה |  |
| 7:21.392 | אין |  |
| 7:21.393 | מזכירים |  |
| 7:21.394 | אותם |  |
| 7:21.395 | אלא |  |
| 7:21.396 | ברמז |  |
| 7:21.397 | וכינוי |  |
| 7:21.398 | ובהצנע |  |
| 7:21.399 | ולא |  |
| 7:21.400 | בפרסום |  |
| 7:21.401 | אבל |  |
| 7:21.402 | שלא |  |
| 7:21.403 | לצורך |  |
| 7:21.404 | הם |  |
| 7:21.405 | מלהזכיר |  |
| 7:21.406 | כלל | . |
| 7:21.407 | כי |  |
| 7:21.408 | זהו |  |
| 7:21.409 | גבלות |  |
| 7:21.410 | הפה |  |
| 7:21.411 | וגנות |  |
| 7:21.412 | גדולה |  |
| 7:21.413 | שלמדה |  |
| 7:21.414 | אותנו |  |
| 7:21.415 | תוה"ק |  |
| 7:21.416 | להשמר |  |
| 7:21.417 | ממנו |  |
| 7:21.418 | כמאמר |  |
| 7:21.419 | רז"ל |  |
| 7:21.420 | בפ"ק |  |
| 7:21.421 | דפסחים |  |
| 7:21.422 | אמר |  |
| 7:21.423 | ר | ' |
| 7:21.424 | יהושע |  |
| 7:21.425 | בן |  |
| 7:21.426 | לוי |  |
| 7:21.427 | לעולם |  |
| 7:21.428 | אל |  |
| 7:21.429 | יוציא |  |
| 7:21.430 | אדם |  |
| 7:21.431 | דבר |  |
| 7:21.432 | מגונה |  |
| 7:21.433 | מפיו |  |
| 7:21.434 | שהרי |  |
| 7:21.435 | עיקם |  |
| 7:21.436 | הכ | ' |
| 7:21.437 | שמנה |  |
| 7:21.438 | אותיות |  |
| 7:21.439 | ולא |  |
| 7:21.440 | הוציא |  |
| 7:21.441 | דבר |  |
| 7:21.442 | מגונה |  |
| 7:21.443 | מפיו |  |
| 7:21.444 | שנ | ' |
| 7:21.445 | ומן |  |
| 7:21.446 | הבהמה |  |
| 7:21.447 | אשר |  |
| 7:21.448 | איננה |  |
| 7:21.449 | טהורה |  |
| 7:21.450 | רב |  |
| 7:21.451 | אשי |  |
| 7:21.452 | אמר |  |
| 7:21.453 | תשע |  |
| 7:21.454 | שנ | ' |
| 7:21.455 | כי |  |
| 7:21.456 | יהיה |  |
| 7:21.457 | בך |  |
| 7:21.458 | איש |  |
| 7:21.459 | אשר |  |
| 7:21.460 | לא |  |
| 7:21.461 | יהיה |  |
| 7:21.462 | טהור |  |
| 7:21.463 | מקרה |  |
| 7:21.464 | לילה | . |
| 7:21.465 | רבינא |  |
| 7:21.466 | אמר |  |
| 7:21.467 | עשר |  |
| 7:21.468 | וי"ו |  |
| 7:21.469 | דטהור | . |
| 7:21.470 | רב |  |
| 7:21.471 | אמר |  |
| 7:21.472 | שש |  |
| 7:21.473 | עשרה |  |
| 7:21.474 | שנ | ' |
| 7:21.475 | כי |  |
| 7:21.476 | אמר |  |
| 7:21.477 | מקרה |  |
| 7:21.478 | הוא |  |
| 7:21.479 | בלתי |  |
| 7:21.480 | טהור |  |
| 7:21.481 | הוא |  |
| 7:21.482 | כי |  |
| 7:21.483 | לא |  |
| 7:21.484 | טהור | . |
| 7:21.485 | תנא |  |
| 7:21.486 | דבי |  |
| 7:21.487 | ר | ' |
| 7:21.488 | ישמעאל |  |
| 7:21.489 | לעולם |  |
| 7:21.490 | יספר |  |
| 7:21.491 | אדם |  |
| 7:21.492 | בלשון |  |
| 7:21.493 | נקייה |  |
| 7:21.494 | שהרי |  |
| 7:21.495 | בזב |  |
| 7:21.496 | וכו | '. |
| 7:21.497 | ואומר |  |
| 7:21.498 | ותבחר |  |
| 7:21.499 | לשון |  |
| 7:21.500 | ערומים | . |
| 7:21.501 | ואומר |  |
| 7:21.502 | וזעת |  |
| 7:21.503 | שפתי |  |
| 7:21.504 | ברור |  |
| 7:21.505 | מללו |  |
| 7:21.506 | וכו | '. |
| 7:21.507 | ובכל |  |
| 7:21.508 | התורה |  |
| 7:21.509 | ובדברי |  |
| 7:21.510 | הנביאים |  |
| 7:21.511 | ורז"ל |  |
| 7:21.512 | בתלמוד |  |
| 7:21.513 | ובמדרשים |  |
| 7:21.514 | לא |  |
| 7:21.515 | מצינו |  |
| 7:21.516 | שכינו |  |
| 7:21.517 | להשפעת |  |
| 7:21.518 | השי"ת |  |
| 7:21.519 | בשם |  |
| 7:21.520 | זכר |  |
| 7:21.521 | ונקבה |  |
| 7:21.522 | וזיווג |  |
| 7:21.523 | אלא |  |
| 7:21.524 | בלשון |  |
| 7:21.525 | נתינה |  |
| 7:21.526 | או |  |
| 7:21.527 | פתיחה |  |
| 7:21.528 | ויתן |  |
| 7:21.529 | לך |  |
| 7:21.530 | האלהים | . |
| 7:21.531 | נותן |  |
| 7:21.532 | לך |  |
| 7:21.533 | כח |  |
| 7:21.534 | לעשות |  |
| 7:21.535 | חיל | . |
| 7:21.536 | יפתח |  |
| 7:21.537 | ה | ' |
| 7:21.538 | לך |  |
| 7:21.539 | את |  |
| 7:21.540 | אוצרו |  |
| 7:21.541 | הטוב |  |
| 7:21.542 | נותן |  |
| 7:21.543 | לחם |  |
| 7:21.544 | לכל |  |
| 7:21.545 | בשר | . |
| 7:21.546 | פותח |  |
| 7:21.547 | ידיך |  |
| 7:21.548 | ומשביע |  |
| 7:21.549 | לכל |  |
| 7:21.550 | חי |  |
| 7:21.551 | רצון | . |
| 7:21.552 | ואתה |  |
| 7:21.553 | מחיה |  |
| 7:21.554 | את |  |
| 7:21.555 | כלם |  |
| 7:21.556 | וכיוצא |  |
| 7:21.557 | בלשונות |  |
| 7:21.558 | אלו |  |
| 7:21.559 | אן |  |
| 7:21.560 | מספר | . |
| 7:21.561 | ובדברי |  |
| 7:21.562 | רז"ל |  |
| 7:21.563 | יושב |  |
| 7:21.564 | הקב"ה |  |
| 7:21.565 | וזן |  |
| 7:21.566 | את |  |
| 7:21.567 | העולם |  |
| 7:21.568 | כו | '. |
| 7:21.569 | זן |  |
| 7:21.570 | ומפרניס |  |
| 7:21.571 | לכל | . |
| 7:21.572 | עורך |  |
| 7:21.573 | שולחן | , |
| 7:21.574 | נותן |  |
| 7:21.575 | טרף |  |
| 7:21.576 | וכיוצא | . |
| 7:21.577 | ולמה |  |
| 7:21.578 | זה |  |
| 7:21.579 | נעזוב |  |
| 7:21.580 | מליצת |  |
| 7:21.581 | ספרי |  |
| 7:21.582 | הקדש |  |
| 7:21.583 | שדברו |  |
| 7:21.584 | בלשון |  |
| 7:21.585 | נקייה | . |
| 7:21.586 | ונשתמש |  |
| 7:21.587 | בשלון |  |
| 7:21.588 | מגונה |  |
| 7:21.589 | סכרים |  |
| 7:21.590 | ונקבות |  |
| 7:21.591 | רבות |  |
| 7:21.592 | וזווגים |  |
| 7:21.593 | והטפת |  |
| 7:21.594 | טיפי |  |
| 7:21.595 | הזרע |  |
| 7:21.596 | ממבושי |  |
| 7:21.597 | הזכר |  |
| 7:21.598 | לרחמה |  |
| 7:21.599 | של |  |
| 7:21.600 | נקבה |  |
| 7:21.601 | וקליטה |  |
| 7:21.602 | וקפאון |  |
| 7:21.603 | הזרע |  |
| 7:21.604 | ברחמה |  |
| 7:21.605 | כמפורסם |  |
| 7:21.606 | כל |  |
| 7:21.607 | זה |  |
| 7:21.608 | בספרי |  |
| 7:21.609 | הקבלה |  |
| 7:21.610 | החדשה |  |
| 7:21.611 | הנכריה | . |
| 7:21.612 | ולא |  |
| 7:21.613 | די |  |
| 7:21.614 | כל |  |
| 7:21.615 | זה |  |
| 7:21.616 | אלא |  |
| 7:21.617 | שהמסית |  |
| 7:21.618 | מחבר |  |
| 7:21.619 | הזהר |  |
| 7:21.620 | העיז |  |
| 7:21.621 | פניו |  |
| 7:21.622 | לכנות |  |
| 7:21.623 | המבושים |  |
| 7:21.624 | אשר |  |
| 7:21.625 | חשב |  |
| 7:21.626 | באדון |  |
| 7:21.627 | העולמים |  |
| 7:21.628 | בשם |  |
| 7:21.629 | בית |  |
| 7:21.630 | אלהים | . |
| 7:21.631 | ושער |  |
| 7:21.632 | השמים | , |
| 7:21.633 | כדאיתא |  |
| 7:21.634 | בזהר |  |
| 7:21.635 | ויצא |  |
| 7:21.636 | דף |  |
| 7:21.637 | ק"ן |  |
| 7:21.638 | ע"ב | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

