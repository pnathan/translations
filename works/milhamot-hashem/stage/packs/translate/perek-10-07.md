# Translation pack: milhamot-hashem perek-10-07 (Milchamot ha-Shem 10:8–10:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-07`.

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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
- מ"ו = "our teacher" (מורנו) in the formula יודיענו מורנו; ב"ו = "blessed be
  He" only before הוא elsewhere ב"ו = בשר ודם ("flesh and blood").
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- ר"ל with a named authority as subject: "means to say".
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
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- מערב in a list with Ashkenaz, Sefarad, Yemen: "the Maghreb", "the West" in the
  note.
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


## Milchamot ha-Shem 10:8

Hebrew (Wikisource): העתקת לשון ר' חטר בסי' מ״ב מערבי ללשון הקדש היסוד השני יחודו ית' וזהו לידע שזה האחד שהוא סבת הכל, אחד לא כאחד הסוג, ולא כאחד המין, ולא כאחד המורכב אשר יחולק לאחדים רבים, ולא כגוף הפשוט שהוא אחד במנין (כמו המים שהוא יסוד פשוט) (חטר) ומתחלק לאין קץ. אבל הוא ית' אחד אחדות שאין כמוה אחדות בשום פנים, וזה היסוד השני הוא הנלמד מאמרו שמע ישראל ה' אלהינו ה' אחד! ביאור היסוד הזה הואיל וכבר ביאר רבי' מצויאותו ית'. הוצרך לבאר אחדותו. כי אלו השני יסודות הם עיקר תורת משה רבינו ע״ה. והמה ראשית כל תורה אלהית, וסוף כל תורה נימוסית. כי התורה האלהית מתחלתאמן האלהיות וגומרת בגשמיות. והתורה הנימוסית הפך זה. ביאור זה שהאדם אם ידע מציאות כל הנמצאות מראש ועד סוף רצוני לומר ממעלה מעלה עד מטה מטה. כל עוד שירד בידיעתו מטה נתמעטה האחדות ונתוסף הרבוי עד שמגיע לעולם השפל שבלי ספק לא ימד ולא יספר סוגי מיניו כל שכן מיני אישי כל מין ומין וכו'. אבל אם בא לעלות בנמצאות ממטה למעלה, הנה כל מה שיוסיף לעלות ימעט הרבוי ותתחזק האחדות עד האחד המיוחד שממנו נמצאו כל האחדים שלכל הנמצאות. והוא ית' אחד מיוחד. שלוש אמונות יש, השונים. והאמגושים. והנצרים. והנני מבאר כל אחת מהם. ומי שיעיען בספריהם יתבאר לו זה. השונים אומרים שיש לעולם שני פועלים אחד מהם האור. והוא פועל הטוב בנמצאים. ולא יוכל לעשות רע. והפועל הב' החשך והוא פועל הרע בעולם ולא יוכל לעשות טוב, וחמה חלוקים. יש מהם אומרים שהאור והחשך חיים. ויש אומרים שהם מתים. אמנם האמגוש הנם אומרים שיש לעולם פועל חכם. גבור מנהיג ואינו גוף. והרחיקו ממנו מעשות כל רע. וקראוהו יזדאן. ויחסו כל הרע והצער ההוים בעולם לשטן. וקראוהו אַהְרְמִן בלשונם והמה חלוקים יש מהם אומרים שהוא קדמון ואינו גוף. ומהם אומרים שהוא מחודש וגוף. ונחלקו ג״כ בסברת חדושו ואמרול שאירע ליזדאן מחשבה ואמר בלבו איך יהיה עניינו. ומי זה יחלוק עלי (יריב עמי) ואמרו שהשטן נהיה מזאת המחשבה. אחר כן השלים עם יזדאן ונתפשרו יחד עד גמירא. וזכרו עוד דברים שאין כאן מקום לזכרם. והטוב אצל בעלי שתי האמונות ר״ל השונים והאמגוש. הוא מה שתתאוה לו הנפש. והרע הוא מה שתתעבהו הנפש. אמנם הנצרים אומרים שהשי״ת ג' אקאנים כפי האמת (אפשר שהוא כמו ספירות) אקנום האב והוא עצם הבורא אצלם (אריך) ואקנום הבן והוא אלעלם (חכמה) ואקנום רוח הקדש, והוא החיים (בינה אלהים חיים). והשלשה אחד, והם כת מישראל. וראייתם שהם ג' ונחשבים אחד מאמרו ית' שמע ישראל ה' אלהינו אחד. אמרו הרי אלו שלשה שמות לשלשה נקראים בשם. אמנם השלשה חשובים אחד. ולפי הנראה שהם מאמינים עצם אחד ושני תארים קדומים. ואחר שזכרתי לך אמונות החיצונים החולקים בייחוד כדי שתזהר מהם אשוב לבאר עניין היסוד הזה וכו' עד כאן לשון רבי חטר ז״ל

| id | word | punctuation after |
|---|---|---|
| 10:8.1 | העתקת |  |
| 10:8.2 | לשון |  |
| 10:8.3 | ר | ' |
| 10:8.4 | חטר |  |
| 10:8.5 | בסי | ' |
| 10:8.6 | מ״ב |  |
| 10:8.7 | מערבי |  |
| 10:8.8 | ללשון |  |
| 10:8.9 | הקדש |  |
| 10:8.10 | היסוד |  |
| 10:8.11 | השני |  |
| 10:8.12 | יחודו |  |
| 10:8.13 | ית | ' |
| 10:8.14 | וזהו |  |
| 10:8.15 | לידע |  |
| 10:8.16 | שזה |  |
| 10:8.17 | האחד |  |
| 10:8.18 | שהוא |  |
| 10:8.19 | סבת |  |
| 10:8.20 | הכל | , |
| 10:8.21 | אחד |  |
| 10:8.22 | לא |  |
| 10:8.23 | כאחד |  |
| 10:8.24 | הסוג | , |
| 10:8.25 | ולא |  |
| 10:8.26 | כאחד |  |
| 10:8.27 | המין | , |
| 10:8.28 | ולא |  |
| 10:8.29 | כאחד |  |
| 10:8.30 | המורכב |  |
| 10:8.31 | אשר |  |
| 10:8.32 | יחולק |  |
| 10:8.33 | לאחדים |  |
| 10:8.34 | רבים | , |
| 10:8.35 | ולא |  |
| 10:8.36 | כגוף |  |
| 10:8.37 | הפשוט |  |
| 10:8.38 | שהוא |  |
| 10:8.39 | אחד |  |
| 10:8.40 | במנין |  |
| 10:8.41 | כמו |  |
| 10:8.42 | המים |  |
| 10:8.43 | שהוא |  |
| 10:8.44 | יסוד |  |
| 10:8.45 | פשוט | ) |
| 10:8.46 | חטר | ) |
| 10:8.47 | ומתחלק |  |
| 10:8.48 | לאין |  |
| 10:8.49 | קץ | . |
| 10:8.50 | אבל |  |
| 10:8.51 | הוא |  |
| 10:8.52 | ית | ' |
| 10:8.53 | אחד |  |
| 10:8.54 | אחדות |  |
| 10:8.55 | שאין |  |
| 10:8.56 | כמוה |  |
| 10:8.57 | אחדות |  |
| 10:8.58 | בשום |  |
| 10:8.59 | פנים | , |
| 10:8.60 | וזה |  |
| 10:8.61 | היסוד |  |
| 10:8.62 | השני |  |
| 10:8.63 | הוא |  |
| 10:8.64 | הנלמד |  |
| 10:8.65 | מאמרו |  |
| 10:8.66 | שמע |  |
| 10:8.67 | ישראל |  |
| 10:8.68 | ה | ' |
| 10:8.69 | אלהינו |  |
| 10:8.70 | ה | ' |
| 10:8.71 | אחד | ! |
| 10:8.72 | ביאור |  |
| 10:8.73 | היסוד |  |
| 10:8.74 | הזה |  |
| 10:8.75 | הואיל |  |
| 10:8.76 | וכבר |  |
| 10:8.77 | ביאר |  |
| 10:8.78 | רבי | ' |
| 10:8.79 | מצויאותו |  |
| 10:8.80 | ית | '. |
| 10:8.81 | הוצרך |  |
| 10:8.82 | לבאר |  |
| 10:8.83 | אחדותו | . |
| 10:8.84 | כי |  |
| 10:8.85 | אלו |  |
| 10:8.86 | השני |  |
| 10:8.87 | יסודות |  |
| 10:8.88 | הם |  |
| 10:8.89 | עיקר |  |
| 10:8.90 | תורת |  |
| 10:8.91 | משה |  |
| 10:8.92 | רבינו |  |
| 10:8.93 | ע״ה | . |
| 10:8.94 | והמה |  |
| 10:8.95 | ראשית |  |
| 10:8.96 | כל |  |
| 10:8.97 | תורה |  |
| 10:8.98 | אלהית | , |
| 10:8.99 | וסוף |  |
| 10:8.100 | כל |  |
| 10:8.101 | תורה |  |
| 10:8.102 | נימוסית | . |
| 10:8.103 | כי |  |
| 10:8.104 | התורה |  |
| 10:8.105 | האלהית |  |
| 10:8.106 | מתחלתאמן |  |
| 10:8.107 | האלהיות |  |
| 10:8.108 | וגומרת |  |
| 10:8.109 | בגשמיות | . |
| 10:8.110 | והתורה |  |
| 10:8.111 | הנימוסית |  |
| 10:8.112 | הפך |  |
| 10:8.113 | זה | . |
| 10:8.114 | ביאור |  |
| 10:8.115 | זה |  |
| 10:8.116 | שהאדם |  |
| 10:8.117 | אם |  |
| 10:8.118 | ידע |  |
| 10:8.119 | מציאות |  |
| 10:8.120 | כל |  |
| 10:8.121 | הנמצאות |  |
| 10:8.122 | מראש |  |
| 10:8.123 | ועד |  |
| 10:8.124 | סוף |  |
| 10:8.125 | רצוני |  |
| 10:8.126 | לומר |  |
| 10:8.127 | ממעלה |  |
| 10:8.128 | מעלה |  |
| 10:8.129 | עד |  |
| 10:8.130 | מטה |  |
| 10:8.131 | מטה | . |
| 10:8.132 | כל |  |
| 10:8.133 | עוד |  |
| 10:8.134 | שירד |  |
| 10:8.135 | בידיעתו |  |
| 10:8.136 | מטה |  |
| 10:8.137 | נתמעטה |  |
| 10:8.138 | האחדות |  |
| 10:8.139 | ונתוסף |  |
| 10:8.140 | הרבוי |  |
| 10:8.141 | עד |  |
| 10:8.142 | שמגיע |  |
| 10:8.143 | לעולם |  |
| 10:8.144 | השפל |  |
| 10:8.145 | שבלי |  |
| 10:8.146 | ספק |  |
| 10:8.147 | לא |  |
| 10:8.148 | ימד |  |
| 10:8.149 | ולא |  |
| 10:8.150 | יספר |  |
| 10:8.151 | סוגי |  |
| 10:8.152 | מיניו |  |
| 10:8.153 | כל |  |
| 10:8.154 | שכן |  |
| 10:8.155 | מיני |  |
| 10:8.156 | אישי |  |
| 10:8.157 | כל |  |
| 10:8.158 | מין |  |
| 10:8.159 | ומין |  |
| 10:8.160 | וכו | '. |
| 10:8.161 | אבל |  |
| 10:8.162 | אם |  |
| 10:8.163 | בא |  |
| 10:8.164 | לעלות |  |
| 10:8.165 | בנמצאות |  |
| 10:8.166 | ממטה |  |
| 10:8.167 | למעלה | , |
| 10:8.168 | הנה |  |
| 10:8.169 | כל |  |
| 10:8.170 | מה |  |
| 10:8.171 | שיוסיף |  |
| 10:8.172 | לעלות |  |
| 10:8.173 | ימעט |  |
| 10:8.174 | הרבוי |  |
| 10:8.175 | ותתחזק |  |
| 10:8.176 | האחדות |  |
| 10:8.177 | עד |  |
| 10:8.178 | האחד |  |
| 10:8.179 | המיוחד |  |
| 10:8.180 | שממנו |  |
| 10:8.181 | נמצאו |  |
| 10:8.182 | כל |  |
| 10:8.183 | האחדים |  |
| 10:8.184 | שלכל |  |
| 10:8.185 | הנמצאות | . |
| 10:8.186 | והוא |  |
| 10:8.187 | ית | ' |
| 10:8.188 | אחד |  |
| 10:8.189 | מיוחד | . |
| 10:8.190 | שלוש |  |
| 10:8.191 | אמונות |  |
| 10:8.192 | יש | , |
| 10:8.193 | השונים | . |
| 10:8.194 | והאמגושים | . |
| 10:8.195 | והנצרים | . |
| 10:8.196 | והנני |  |
| 10:8.197 | מבאר |  |
| 10:8.198 | כל |  |
| 10:8.199 | אחת |  |
| 10:8.200 | מהם | . |
| 10:8.201 | ומי |  |
| 10:8.202 | שיעיען |  |
| 10:8.203 | בספריהם |  |
| 10:8.204 | יתבאר |  |
| 10:8.205 | לו |  |
| 10:8.206 | זה | . |
| 10:8.207 | השונים |  |
| 10:8.208 | אומרים |  |
| 10:8.209 | שיש |  |
| 10:8.210 | לעולם |  |
| 10:8.211 | שני |  |
| 10:8.212 | פועלים |  |
| 10:8.213 | אחד |  |
| 10:8.214 | מהם |  |
| 10:8.215 | האור | . |
| 10:8.216 | והוא |  |
| 10:8.217 | פועל |  |
| 10:8.218 | הטוב |  |
| 10:8.219 | בנמצאים | . |
| 10:8.220 | ולא |  |
| 10:8.221 | יוכל |  |
| 10:8.222 | לעשות |  |
| 10:8.223 | רע | . |
| 10:8.224 | והפועל |  |
| 10:8.225 | הב | ' |
| 10:8.226 | החשך |  |
| 10:8.227 | והוא |  |
| 10:8.228 | פועל |  |
| 10:8.229 | הרע |  |
| 10:8.230 | בעולם |  |
| 10:8.231 | ולא |  |
| 10:8.232 | יוכל |  |
| 10:8.233 | לעשות |  |
| 10:8.234 | טוב | , |
| 10:8.235 | וחמה |  |
| 10:8.236 | חלוקים | . |
| 10:8.237 | יש |  |
| 10:8.238 | מהם |  |
| 10:8.239 | אומרים |  |
| 10:8.240 | שהאור |  |
| 10:8.241 | והחשך |  |
| 10:8.242 | חיים | . |
| 10:8.243 | ויש |  |
| 10:8.244 | אומרים |  |
| 10:8.245 | שהם |  |
| 10:8.246 | מתים | . |
| 10:8.247 | אמנם |  |
| 10:8.248 | האמגוש |  |
| 10:8.249 | הנם |  |
| 10:8.250 | אומרים |  |
| 10:8.251 | שיש |  |
| 10:8.252 | לעולם |  |
| 10:8.253 | פועל |  |
| 10:8.254 | חכם | . |
| 10:8.255 | גבור |  |
| 10:8.256 | מנהיג |  |
| 10:8.257 | ואינו |  |
| 10:8.258 | גוף | . |
| 10:8.259 | והרחיקו |  |
| 10:8.260 | ממנו |  |
| 10:8.261 | מעשות |  |
| 10:8.262 | כל |  |
| 10:8.263 | רע | . |
| 10:8.264 | וקראוהו |  |
| 10:8.265 | יזדאן | . |
| 10:8.266 | ויחסו |  |
| 10:8.267 | כל |  |
| 10:8.268 | הרע |  |
| 10:8.269 | והצער |  |
| 10:8.270 | ההוים |  |
| 10:8.271 | בעולם |  |
| 10:8.272 | לשטן | . |
| 10:8.273 | וקראוהו |  |
| 10:8.274 | אַהְרְמִן |  |
| 10:8.275 | בלשונם |  |
| 10:8.276 | והמה |  |
| 10:8.277 | חלוקים |  |
| 10:8.278 | יש |  |
| 10:8.279 | מהם |  |
| 10:8.280 | אומרים |  |
| 10:8.281 | שהוא |  |
| 10:8.282 | קדמון |  |
| 10:8.283 | ואינו |  |
| 10:8.284 | גוף | . |
| 10:8.285 | ומהם |  |
| 10:8.286 | אומרים |  |
| 10:8.287 | שהוא |  |
| 10:8.288 | מחודש |  |
| 10:8.289 | וגוף | . |
| 10:8.290 | ונחלקו |  |
| 10:8.291 | ג״כ |  |
| 10:8.292 | בסברת |  |
| 10:8.293 | חדושו |  |
| 10:8.294 | ואמרול |  |
| 10:8.295 | שאירע |  |
| 10:8.296 | ליזדאן |  |
| 10:8.297 | מחשבה |  |
| 10:8.298 | ואמר |  |
| 10:8.299 | בלבו |  |
| 10:8.300 | איך |  |
| 10:8.301 | יהיה |  |
| 10:8.302 | עניינו | . |
| 10:8.303 | ומי |  |
| 10:8.304 | זה |  |
| 10:8.305 | יחלוק |  |
| 10:8.306 | עלי |  |
| 10:8.307 | יריב |  |
| 10:8.308 | עמי | ) |
| 10:8.309 | ואמרו |  |
| 10:8.310 | שהשטן |  |
| 10:8.311 | נהיה |  |
| 10:8.312 | מזאת |  |
| 10:8.313 | המחשבה | . |
| 10:8.314 | אחר |  |
| 10:8.315 | כן |  |
| 10:8.316 | השלים |  |
| 10:8.317 | עם |  |
| 10:8.318 | יזדאן |  |
| 10:8.319 | ונתפשרו |  |
| 10:8.320 | יחד |  |
| 10:8.321 | עד |  |
| 10:8.322 | גמירא | . |
| 10:8.323 | וזכרו |  |
| 10:8.324 | עוד |  |
| 10:8.325 | דברים |  |
| 10:8.326 | שאין |  |
| 10:8.327 | כאן |  |
| 10:8.328 | מקום |  |
| 10:8.329 | לזכרם | . |
| 10:8.330 | והטוב |  |
| 10:8.331 | אצל |  |
| 10:8.332 | בעלי |  |
| 10:8.333 | שתי |  |
| 10:8.334 | האמונות |  |
| 10:8.335 | ר״ל |  |
| 10:8.336 | השונים |  |
| 10:8.337 | והאמגוש | . |
| 10:8.338 | הוא |  |
| 10:8.339 | מה |  |
| 10:8.340 | שתתאוה |  |
| 10:8.341 | לו |  |
| 10:8.342 | הנפש | . |
| 10:8.343 | והרע |  |
| 10:8.344 | הוא |  |
| 10:8.345 | מה |  |
| 10:8.346 | שתתעבהו |  |
| 10:8.347 | הנפש | . |
| 10:8.348 | אמנם |  |
| 10:8.349 | הנצרים |  |
| 10:8.350 | אומרים |  |
| 10:8.351 | שהשי״ת |  |
| 10:8.352 | ג | ' |
| 10:8.353 | אקאנים |  |
| 10:8.354 | כפי |  |
| 10:8.355 | האמת |  |
| 10:8.356 | אפשר |  |
| 10:8.357 | שהוא |  |
| 10:8.358 | כמו |  |
| 10:8.359 | ספירות | ) |
| 10:8.360 | אקנום |  |
| 10:8.361 | האב |  |
| 10:8.362 | והוא |  |
| 10:8.363 | עצם |  |
| 10:8.364 | הבורא |  |
| 10:8.365 | אצלם |  |
| 10:8.366 | אריך | ) |
| 10:8.367 | ואקנום |  |
| 10:8.368 | הבן |  |
| 10:8.369 | והוא |  |
| 10:8.370 | אלעלם |  |
| 10:8.371 | חכמה | ) |
| 10:8.372 | ואקנום |  |
| 10:8.373 | רוח |  |
| 10:8.374 | הקדש | , |
| 10:8.375 | והוא |  |
| 10:8.376 | החיים |  |
| 10:8.377 | בינה |  |
| 10:8.378 | אלהים |  |
| 10:8.379 | חיים | ). |
| 10:8.380 | והשלשה |  |
| 10:8.381 | אחד | , |
| 10:8.382 | והם |  |
| 10:8.383 | כת |  |
| 10:8.384 | מישראל | . |
| 10:8.385 | וראייתם |  |
| 10:8.386 | שהם |  |
| 10:8.387 | ג | ' |
| 10:8.388 | ונחשבים |  |
| 10:8.389 | אחד |  |
| 10:8.390 | מאמרו |  |
| 10:8.391 | ית | ' |
| 10:8.392 | שמע |  |
| 10:8.393 | ישראל |  |
| 10:8.394 | ה | ' |
| 10:8.395 | אלהינו |  |
| 10:8.396 | אחד | . |
| 10:8.397 | אמרו |  |
| 10:8.398 | הרי |  |
| 10:8.399 | אלו |  |
| 10:8.400 | שלשה |  |
| 10:8.401 | שמות |  |
| 10:8.402 | לשלשה |  |
| 10:8.403 | נקראים |  |
| 10:8.404 | בשם | . |
| 10:8.405 | אמנם |  |
| 10:8.406 | השלשה |  |
| 10:8.407 | חשובים |  |
| 10:8.408 | אחד | . |
| 10:8.409 | ולפי |  |
| 10:8.410 | הנראה |  |
| 10:8.411 | שהם |  |
| 10:8.412 | מאמינים |  |
| 10:8.413 | עצם |  |
| 10:8.414 | אחד |  |
| 10:8.415 | ושני |  |
| 10:8.416 | תארים |  |
| 10:8.417 | קדומים | . |
| 10:8.418 | ואחר |  |
| 10:8.419 | שזכרתי |  |
| 10:8.420 | לך |  |
| 10:8.421 | אמונות |  |
| 10:8.422 | החיצונים |  |
| 10:8.423 | החולקים |  |
| 10:8.424 | בייחוד |  |
| 10:8.425 | כדי |  |
| 10:8.426 | שתזהר |  |
| 10:8.427 | מהם |  |
| 10:8.428 | אשוב |  |
| 10:8.429 | לבאר |  |
| 10:8.430 | עניין |  |
| 10:8.431 | היסוד |  |
| 10:8.432 | הזה |  |
| 10:8.433 | וכו | ' |
| 10:8.434 | עד |  |
| 10:8.435 | כאן |  |
| 10:8.436 | לשון |  |
| 10:8.437 | רבי |  |
| 10:8.438 | חטר |  |
| 10:8.439 | ז״ל |  |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

