# Translation pack: milhamot-hashem perek-4-19 (Milchamot ha-Shem 4:31–4:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-19`.

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
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
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
- כלא חד / כולא חד: the same English at every occurrence (1:5-1:7); the
  owner's ruling on the sense is pending; until then, the alternative is in
  `apparatus`.
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
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


## Milchamot ha-Shem 4:31

Hebrew (Wikisource): לז) ומעתה תן דעתך והתבונן מאד בדברי ר"ת בן יחיא המובא בס' רביד הזהב והעתקיו לעיל סי' י"ג שאמר שהמשנה והתלמוד היא הקבלה האמתית והמוסכמת מכל צד שרוצה לומר שאין בה שום ספק. מה שאין כן ספרי הקבלה החדשה יש לספק בהן כמה ספקות. וכבר יצאו עליהם עסיקין כידוע לכל מבין בקי בספרות ישראל. שאין לסמוך עליהם בשום דין מדיני התורה היפך התלמוד והפוסקים כמו שכתבו בעלי הכללים בדרכי הפוסקים. כל שכן וקל וחומר ביחוד השי"ת ועבודתו אשר על זה הוזהרנו בתורה ובנביאים ובכתובים וכדברי רז"ל בכמה אזהרות כמו שנתבאר לעיל. וכדתנן למה קדמה שמע לוהיה אם שמוע וכו' וכמו שכ' הרמב"ם ז"ל בפ"א מהל' קריאת שמע שמקדימין לקרות פרשת שמע מפני שיש בה ציווי על יחוד השם ואהבתו ותלמודו שהוא העיקר הגדול שהכל תלוי בו ע"כ. וכבר כתבתי לעיל סי' ל"ב שרז"ל הקפידו מאד על כל מלתא דנפיק מינה להאמין בשתי רשויות ואמרו שמשתקין אותו. ומה יתן ומה יוסיף עוד אמרו וכלא הד. אחר שכבר אמר וחשב לנו כמה עילות שנוטלין רשות כל אחד מן העילה שלמעלה ממנו. כאלו נצטוינו לומר אחד בפינו. אבל בליבנו הם הרבה אלוהות. וזה לשון הרמב"ם במורה פרק נ' לראשון, דע אתה המעיין במאמרי זה, כי האמונה אינה הענין הנאמר בפה אבל הוא הענים המצוייר בנפש, כשיאמינו בו שהוא כן כמו שצוייר, ואם יספיק לך מן הדעות האמתיות או הנחשבות אמתתם אצלך כשתספרם במאמר מבלתי שתציירם ותאמין בהם. כל שכן כשתבקש מהם האמת. הוא קל מאוד, כמו שתמצא הרבה מן הפתאין ישמרו אמונות לא יציירו להם ענין בשום פנים. אבל אם מלאך לבך לעלות לזאת המדרגה העליונה שהיא מדגרת העיון. ושיתאמת לך שהשם ית' אחד האחדות האמתית עד שלא תמצא לו הרכבה כלל. ושאין לחשוב שום מחשבת חלוק בשום פנים. דע שאין לו ית' תואר עצמי בשום פנים ולא בשום ענין. וכמו שנמנע היותו גשם כן נמנע היותו בעל תואר עצמי. אמנם מי שהאמין שהוא אחד בעל תארים רבים כבר אמר שהוא אחד במלתו. והאמין שהוא רבים במחשבתו וזה כמאמר הנצרים הוא אחד. אבל הוא שלשה והשלשה אחד. כן מאמר האומר הוא אחד, אבל הוא בעל תארים רבים, והוא ותאריו אחד. עם הסתלק הגשמות, והאמנת הפשטות הגמורה כאלו כוונתנו וחיפושנו איך נאמר, לא איך נאמין, ואין האמונה אלא ארת הציור. כי האמונה היא האמונה במה שיצוייר בשכל. ואם יהיה עם זאת האמונה שאי אפשר חלוף זאת האמונה בשום פנים. ולא ימצא בשכל מקום דחייה לאמונה ההיא ולא לשער אפשרות חלופה תהיה אמתית. וכשתפשוט מעליך התאוות והמנהגים, ותהיה בעל תבונה, ותתבונן מה שאומר אותו באלה הפרקים הבאים בהרחקת התארים יתאמת לך מה שאמרנו בהכרח. ותהיה אז ממי שיצייר יחוד השי"ת. לא ממי שיאמר אותו בפיו ולא יצייר לו עניין. ויהיה מכת הנאמר עליהם קרוב אתה בפיהם ורחוק מכליותיהם. אבל צריך שיהיה האדם מכת מי שיצייר האמת וישיגהו אע"פ שלא ידבר בו. כמו שצוו החשובים ונאמר להם אמרו בלבבכם על משכבכם ודומו סלה עכ"ל הרב המורה ז"ל.

| id | word | punctuation after |
|---|---|---|
| 4:31.1 | לז | ) |
| 4:31.2 | ומעתה |  |
| 4:31.3 | תן |  |
| 4:31.4 | דעתך |  |
| 4:31.5 | והתבונן |  |
| 4:31.6 | מאד |  |
| 4:31.7 | בדברי |  |
| 4:31.8 | ר"ת |  |
| 4:31.9 | בן |  |
| 4:31.10 | יחיא |  |
| 4:31.11 | המובא |  |
| 4:31.12 | בס | ' |
| 4:31.13 | רביד |  |
| 4:31.14 | הזהב |  |
| 4:31.15 | והעתקיו |  |
| 4:31.16 | לעיל |  |
| 4:31.17 | סי | ' |
| 4:31.18 | י"ג |  |
| 4:31.19 | שאמר |  |
| 4:31.20 | שהמשנה |  |
| 4:31.21 | והתלמוד |  |
| 4:31.22 | היא |  |
| 4:31.23 | הקבלה |  |
| 4:31.24 | האמתית |  |
| 4:31.25 | והמוסכמת |  |
| 4:31.26 | מכל |  |
| 4:31.27 | צד |  |
| 4:31.28 | שרוצה |  |
| 4:31.29 | לומר |  |
| 4:31.30 | שאין |  |
| 4:31.31 | בה |  |
| 4:31.32 | שום |  |
| 4:31.33 | ספק | . |
| 4:31.34 | מה |  |
| 4:31.35 | שאין |  |
| 4:31.36 | כן |  |
| 4:31.37 | ספרי |  |
| 4:31.38 | הקבלה |  |
| 4:31.39 | החדשה |  |
| 4:31.40 | יש |  |
| 4:31.41 | לספק |  |
| 4:31.42 | בהן |  |
| 4:31.43 | כמה |  |
| 4:31.44 | ספקות | . |
| 4:31.45 | וכבר |  |
| 4:31.46 | יצאו |  |
| 4:31.47 | עליהם |  |
| 4:31.48 | עסיקין |  |
| 4:31.49 | כידוע |  |
| 4:31.50 | לכל |  |
| 4:31.51 | מבין |  |
| 4:31.52 | בקי |  |
| 4:31.53 | בספרות |  |
| 4:31.54 | ישראל | . |
| 4:31.55 | שאין |  |
| 4:31.56 | לסמוך |  |
| 4:31.57 | עליהם |  |
| 4:31.58 | בשום |  |
| 4:31.59 | דין |  |
| 4:31.60 | מדיני |  |
| 4:31.61 | התורה |  |
| 4:31.62 | היפך |  |
| 4:31.63 | התלמוד |  |
| 4:31.64 | והפוסקים |  |
| 4:31.65 | כמו |  |
| 4:31.66 | שכתבו |  |
| 4:31.67 | בעלי |  |
| 4:31.68 | הכללים |  |
| 4:31.69 | בדרכי |  |
| 4:31.70 | הפוסקים | . |
| 4:31.71 | כל |  |
| 4:31.72 | שכן |  |
| 4:31.73 | וקל |  |
| 4:31.74 | וחומר |  |
| 4:31.75 | ביחוד |  |
| 4:31.76 | השי"ת |  |
| 4:31.77 | ועבודתו |  |
| 4:31.78 | אשר |  |
| 4:31.79 | על |  |
| 4:31.80 | זה |  |
| 4:31.81 | הוזהרנו |  |
| 4:31.82 | בתורה |  |
| 4:31.83 | ובנביאים |  |
| 4:31.84 | ובכתובים |  |
| 4:31.85 | וכדברי |  |
| 4:31.86 | רז"ל |  |
| 4:31.87 | בכמה |  |
| 4:31.88 | אזהרות |  |
| 4:31.89 | כמו |  |
| 4:31.90 | שנתבאר |  |
| 4:31.91 | לעיל | . |
| 4:31.92 | וכדתנן |  |
| 4:31.93 | למה |  |
| 4:31.94 | קדמה |  |
| 4:31.95 | שמע |  |
| 4:31.96 | לוהיה |  |
| 4:31.97 | אם |  |
| 4:31.98 | שמוע |  |
| 4:31.99 | וכו | ' |
| 4:31.100 | וכמו |  |
| 4:31.101 | שכ | ' |
| 4:31.102 | הרמב"ם |  |
| 4:31.103 | ז"ל |  |
| 4:31.104 | בפ"א |  |
| 4:31.105 | מהל | ' |
| 4:31.106 | קריאת |  |
| 4:31.107 | שמע |  |
| 4:31.108 | שמקדימין |  |
| 4:31.109 | לקרות |  |
| 4:31.110 | פרשת |  |
| 4:31.111 | שמע |  |
| 4:31.112 | מפני |  |
| 4:31.113 | שיש |  |
| 4:31.114 | בה |  |
| 4:31.115 | ציווי |  |
| 4:31.116 | על |  |
| 4:31.117 | יחוד |  |
| 4:31.118 | השם |  |
| 4:31.119 | ואהבתו |  |
| 4:31.120 | ותלמודו |  |
| 4:31.121 | שהוא |  |
| 4:31.122 | העיקר |  |
| 4:31.123 | הגדול |  |
| 4:31.124 | שהכל |  |
| 4:31.125 | תלוי |  |
| 4:31.126 | בו |  |
| 4:31.127 | ע"כ | . |
| 4:31.128 | וכבר |  |
| 4:31.129 | כתבתי |  |
| 4:31.130 | לעיל |  |
| 4:31.131 | סי | ' |
| 4:31.132 | ל"ב |  |
| 4:31.133 | שרז"ל |  |
| 4:31.134 | הקפידו |  |
| 4:31.135 | מאד |  |
| 4:31.136 | על |  |
| 4:31.137 | כל |  |
| 4:31.138 | מלתא |  |
| 4:31.139 | דנפיק |  |
| 4:31.140 | מינה |  |
| 4:31.141 | להאמין |  |
| 4:31.142 | בשתי |  |
| 4:31.143 | רשויות |  |
| 4:31.144 | ואמרו |  |
| 4:31.145 | שמשתקין |  |
| 4:31.146 | אותו | . |
| 4:31.147 | ומה |  |
| 4:31.148 | יתן |  |
| 4:31.149 | ומה |  |
| 4:31.150 | יוסיף |  |
| 4:31.151 | עוד |  |
| 4:31.152 | אמרו |  |
| 4:31.153 | וכלא |  |
| 4:31.154 | הד | . |
| 4:31.155 | אחר |  |
| 4:31.156 | שכבר |  |
| 4:31.157 | אמר |  |
| 4:31.158 | וחשב |  |
| 4:31.159 | לנו |  |
| 4:31.160 | כמה |  |
| 4:31.161 | עילות |  |
| 4:31.162 | שנוטלין |  |
| 4:31.163 | רשות |  |
| 4:31.164 | כל |  |
| 4:31.165 | אחד |  |
| 4:31.166 | מן |  |
| 4:31.167 | העילה |  |
| 4:31.168 | שלמעלה |  |
| 4:31.169 | ממנו | . |
| 4:31.170 | כאלו |  |
| 4:31.171 | נצטוינו |  |
| 4:31.172 | לומר |  |
| 4:31.173 | אחד |  |
| 4:31.174 | בפינו | . |
| 4:31.175 | אבל |  |
| 4:31.176 | בליבנו |  |
| 4:31.177 | הם |  |
| 4:31.178 | הרבה |  |
| 4:31.179 | אלוהות | . |
| 4:31.180 | וזה |  |
| 4:31.181 | לשון |  |
| 4:31.182 | הרמב"ם |  |
| 4:31.183 | במורה |  |
| 4:31.184 | פרק |  |
| 4:31.185 | נ | ' |
| 4:31.186 | לראשון | , |
| 4:31.187 | דע |  |
| 4:31.188 | אתה |  |
| 4:31.189 | המעיין |  |
| 4:31.190 | במאמרי |  |
| 4:31.191 | זה | , |
| 4:31.192 | כי |  |
| 4:31.193 | האמונה |  |
| 4:31.194 | אינה |  |
| 4:31.195 | הענין |  |
| 4:31.196 | הנאמר |  |
| 4:31.197 | בפה |  |
| 4:31.198 | אבל |  |
| 4:31.199 | הוא |  |
| 4:31.200 | הענים |  |
| 4:31.201 | המצוייר |  |
| 4:31.202 | בנפש | , |
| 4:31.203 | כשיאמינו |  |
| 4:31.204 | בו |  |
| 4:31.205 | שהוא |  |
| 4:31.206 | כן |  |
| 4:31.207 | כמו |  |
| 4:31.208 | שצוייר | , |
| 4:31.209 | ואם |  |
| 4:31.210 | יספיק |  |
| 4:31.211 | לך |  |
| 4:31.212 | מן |  |
| 4:31.213 | הדעות |  |
| 4:31.214 | האמתיות |  |
| 4:31.215 | או |  |
| 4:31.216 | הנחשבות |  |
| 4:31.217 | אמתתם |  |
| 4:31.218 | אצלך |  |
| 4:31.219 | כשתספרם |  |
| 4:31.220 | במאמר |  |
| 4:31.221 | מבלתי |  |
| 4:31.222 | שתציירם |  |
| 4:31.223 | ותאמין |  |
| 4:31.224 | בהם | . |
| 4:31.225 | כל |  |
| 4:31.226 | שכן |  |
| 4:31.227 | כשתבקש |  |
| 4:31.228 | מהם |  |
| 4:31.229 | האמת | . |
| 4:31.230 | הוא |  |
| 4:31.231 | קל |  |
| 4:31.232 | מאוד | , |
| 4:31.233 | כמו |  |
| 4:31.234 | שתמצא |  |
| 4:31.235 | הרבה |  |
| 4:31.236 | מן |  |
| 4:31.237 | הפתאין |  |
| 4:31.238 | ישמרו |  |
| 4:31.239 | אמונות |  |
| 4:31.240 | לא |  |
| 4:31.241 | יציירו |  |
| 4:31.242 | להם |  |
| 4:31.243 | ענין |  |
| 4:31.244 | בשום |  |
| 4:31.245 | פנים | . |
| 4:31.246 | אבל |  |
| 4:31.247 | אם |  |
| 4:31.248 | מלאך |  |
| 4:31.249 | לבך |  |
| 4:31.250 | לעלות |  |
| 4:31.251 | לזאת |  |
| 4:31.252 | המדרגה |  |
| 4:31.253 | העליונה |  |
| 4:31.254 | שהיא |  |
| 4:31.255 | מדגרת |  |
| 4:31.256 | העיון | . |
| 4:31.257 | ושיתאמת |  |
| 4:31.258 | לך |  |
| 4:31.259 | שהשם |  |
| 4:31.260 | ית | ' |
| 4:31.261 | אחד |  |
| 4:31.262 | האחדות |  |
| 4:31.263 | האמתית |  |
| 4:31.264 | עד |  |
| 4:31.265 | שלא |  |
| 4:31.266 | תמצא |  |
| 4:31.267 | לו |  |
| 4:31.268 | הרכבה |  |
| 4:31.269 | כלל | . |
| 4:31.270 | ושאין |  |
| 4:31.271 | לחשוב |  |
| 4:31.272 | שום |  |
| 4:31.273 | מחשבת |  |
| 4:31.274 | חלוק |  |
| 4:31.275 | בשום |  |
| 4:31.276 | פנים | . |
| 4:31.277 | דע |  |
| 4:31.278 | שאין |  |
| 4:31.279 | לו |  |
| 4:31.280 | ית | ' |
| 4:31.281 | תואר |  |
| 4:31.282 | עצמי |  |
| 4:31.283 | בשום |  |
| 4:31.284 | פנים |  |
| 4:31.285 | ולא |  |
| 4:31.286 | בשום |  |
| 4:31.287 | ענין | . |
| 4:31.288 | וכמו |  |
| 4:31.289 | שנמנע |  |
| 4:31.290 | היותו |  |
| 4:31.291 | גשם |  |
| 4:31.292 | כן |  |
| 4:31.293 | נמנע |  |
| 4:31.294 | היותו |  |
| 4:31.295 | בעל |  |
| 4:31.296 | תואר |  |
| 4:31.297 | עצמי | . |
| 4:31.298 | אמנם |  |
| 4:31.299 | מי |  |
| 4:31.300 | שהאמין |  |
| 4:31.301 | שהוא |  |
| 4:31.302 | אחד |  |
| 4:31.303 | בעל |  |
| 4:31.304 | תארים |  |
| 4:31.305 | רבים |  |
| 4:31.306 | כבר |  |
| 4:31.307 | אמר |  |
| 4:31.308 | שהוא |  |
| 4:31.309 | אחד |  |
| 4:31.310 | במלתו | . |
| 4:31.311 | והאמין |  |
| 4:31.312 | שהוא |  |
| 4:31.313 | רבים |  |
| 4:31.314 | במחשבתו |  |
| 4:31.315 | וזה |  |
| 4:31.316 | כמאמר |  |
| 4:31.317 | הנצרים |  |
| 4:31.318 | הוא |  |
| 4:31.319 | אחד | . |
| 4:31.320 | אבל |  |
| 4:31.321 | הוא |  |
| 4:31.322 | שלשה |  |
| 4:31.323 | והשלשה |  |
| 4:31.324 | אחד | . |
| 4:31.325 | כן |  |
| 4:31.326 | מאמר |  |
| 4:31.327 | האומר |  |
| 4:31.328 | הוא |  |
| 4:31.329 | אחד | , |
| 4:31.330 | אבל |  |
| 4:31.331 | הוא |  |
| 4:31.332 | בעל |  |
| 4:31.333 | תארים |  |
| 4:31.334 | רבים | , |
| 4:31.335 | והוא |  |
| 4:31.336 | ותאריו |  |
| 4:31.337 | אחד | . |
| 4:31.338 | עם |  |
| 4:31.339 | הסתלק |  |
| 4:31.340 | הגשמות | , |
| 4:31.341 | והאמנת |  |
| 4:31.342 | הפשטות |  |
| 4:31.343 | הגמורה |  |
| 4:31.344 | כאלו |  |
| 4:31.345 | כוונתנו |  |
| 4:31.346 | וחיפושנו |  |
| 4:31.347 | איך |  |
| 4:31.348 | נאמר | , |
| 4:31.349 | לא |  |
| 4:31.350 | איך |  |
| 4:31.351 | נאמין | , |
| 4:31.352 | ואין |  |
| 4:31.353 | האמונה |  |
| 4:31.354 | אלא |  |
| 4:31.355 | ארת |  |
| 4:31.356 | הציור | . |
| 4:31.357 | כי |  |
| 4:31.358 | האמונה |  |
| 4:31.359 | היא |  |
| 4:31.360 | האמונה |  |
| 4:31.361 | במה |  |
| 4:31.362 | שיצוייר |  |
| 4:31.363 | בשכל | . |
| 4:31.364 | ואם |  |
| 4:31.365 | יהיה |  |
| 4:31.366 | עם |  |
| 4:31.367 | זאת |  |
| 4:31.368 | האמונה |  |
| 4:31.369 | שאי |  |
| 4:31.370 | אפשר |  |
| 4:31.371 | חלוף |  |
| 4:31.372 | זאת |  |
| 4:31.373 | האמונה |  |
| 4:31.374 | בשום |  |
| 4:31.375 | פנים | . |
| 4:31.376 | ולא |  |
| 4:31.377 | ימצא |  |
| 4:31.378 | בשכל |  |
| 4:31.379 | מקום |  |
| 4:31.380 | דחייה |  |
| 4:31.381 | לאמונה |  |
| 4:31.382 | ההיא |  |
| 4:31.383 | ולא |  |
| 4:31.384 | לשער |  |
| 4:31.385 | אפשרות |  |
| 4:31.386 | חלופה |  |
| 4:31.387 | תהיה |  |
| 4:31.388 | אמתית | . |
| 4:31.389 | וכשתפשוט |  |
| 4:31.390 | מעליך |  |
| 4:31.391 | התאוות |  |
| 4:31.392 | והמנהגים | , |
| 4:31.393 | ותהיה |  |
| 4:31.394 | בעל |  |
| 4:31.395 | תבונה | , |
| 4:31.396 | ותתבונן |  |
| 4:31.397 | מה |  |
| 4:31.398 | שאומר |  |
| 4:31.399 | אותו |  |
| 4:31.400 | באלה |  |
| 4:31.401 | הפרקים |  |
| 4:31.402 | הבאים |  |
| 4:31.403 | בהרחקת |  |
| 4:31.404 | התארים |  |
| 4:31.405 | יתאמת |  |
| 4:31.406 | לך |  |
| 4:31.407 | מה |  |
| 4:31.408 | שאמרנו |  |
| 4:31.409 | בהכרח | . |
| 4:31.410 | ותהיה |  |
| 4:31.411 | אז |  |
| 4:31.412 | ממי |  |
| 4:31.413 | שיצייר |  |
| 4:31.414 | יחוד |  |
| 4:31.415 | השי"ת | . |
| 4:31.416 | לא |  |
| 4:31.417 | ממי |  |
| 4:31.418 | שיאמר |  |
| 4:31.419 | אותו |  |
| 4:31.420 | בפיו |  |
| 4:31.421 | ולא |  |
| 4:31.422 | יצייר |  |
| 4:31.423 | לו |  |
| 4:31.424 | עניין | . |
| 4:31.425 | ויהיה |  |
| 4:31.426 | מכת |  |
| 4:31.427 | הנאמר |  |
| 4:31.428 | עליהם |  |
| 4:31.429 | קרוב |  |
| 4:31.430 | אתה |  |
| 4:31.431 | בפיהם |  |
| 4:31.432 | ורחוק |  |
| 4:31.433 | מכליותיהם | . |
| 4:31.434 | אבל |  |
| 4:31.435 | צריך |  |
| 4:31.436 | שיהיה |  |
| 4:31.437 | האדם |  |
| 4:31.438 | מכת |  |
| 4:31.439 | מי |  |
| 4:31.440 | שיצייר |  |
| 4:31.441 | האמת |  |
| 4:31.442 | וישיגהו |  |
| 4:31.443 | אע"פ |  |
| 4:31.444 | שלא |  |
| 4:31.445 | ידבר |  |
| 4:31.446 | בו | . |
| 4:31.447 | כמו |  |
| 4:31.448 | שצוו |  |
| 4:31.449 | החשובים |  |
| 4:31.450 | ונאמר |  |
| 4:31.451 | להם |  |
| 4:31.452 | אמרו |  |
| 4:31.453 | בלבבכם |  |
| 4:31.454 | על |  |
| 4:31.455 | משכבכם |  |
| 4:31.456 | ודומו |  |
| 4:31.457 | סלה |  |
| 4:31.458 | עכ"ל |  |
| 4:31.459 | הרב |  |
| 4:31.460 | המורה |  |
| 4:31.461 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

