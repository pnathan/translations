# Translation pack: milhamot-hashem perek-8-01 (Milchamot ha-Shem 8:1–8:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-01`.

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
- ס' before a quoted verse is "verse" (סימן = "section" in the author's
  numbered sections); read it from the context and note the choice.
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


## Milchamot ha-Shem 8:1

Hebrew (Wikisource): ק) וראיתי להרב נחלת יוסף לר' שמואל איש עדן שהביא קצור אמונת הקבלה החדשה, וטרח לישב קצת סתירות במשלים שאינם מתקבלים על השכל הישר. ואין בהם כדאי לאוהב את יחיד השי"ת על דרך האמת אשר הורונו רז"ל על פי תורתינו הקדושה, כמו שיראה המעיין שם במשל שהביא שם! וגם הוא סובר שכל עבודתינו ותפלותינו וברכותינו הם לאלוה קצר אפים (זעיר אנפין) והוא אלהינו ואנחנו עמו ונחלתו, והוא המלך הקדוש המהולל בכל התושבחות. השולט על כל הברואים וזנם ומפרנסם ומנהיגם. ושזעיר אנפין הוא שירד על הר סיני להתחבר עם המלכות שנקראת אש. שכבר קדמה וירדה על הר סיני. ואמר אנכי ה' אלהיך אשר הוציאתיך מארץ מצרים. היא אימא עילאה בינה. כי היא הוציאתם ממצרים על די המלכות שנתנה לה גבורותיה, ומלכות היא המכה במצרים בסיועת הבינה, וכן בסיני על ידי הבור בינה עם תפארת ומלכות נאמר אנכי ה' אלהיך. והתורה יצאה מחכמה עלאה עם כללות הבינה. והכל בכח אין סוף ב"ה וכו' ע"ש דף ס"א וס"ב. ודבריו ג"ד מכיינים לדברי מהר"י לופיס לעיל סי' כ"ו שזעיר אנפין הוא שנגלה על אבותינו בסיני ואמר אנכי ה' אלהיך.

| id | word | punctuation after |
|---|---|---|
| 8:1.1 | ק | ) |
| 8:1.2 | וראיתי |  |
| 8:1.3 | להרב |  |
| 8:1.4 | נחלת |  |
| 8:1.5 | יוסף |  |
| 8:1.6 | לר | ' |
| 8:1.7 | שמואל |  |
| 8:1.8 | איש |  |
| 8:1.9 | עדן |  |
| 8:1.10 | שהביא |  |
| 8:1.11 | קצור |  |
| 8:1.12 | אמונת |  |
| 8:1.13 | הקבלה |  |
| 8:1.14 | החדשה | , |
| 8:1.15 | וטרח |  |
| 8:1.16 | לישב |  |
| 8:1.17 | קצת |  |
| 8:1.18 | סתירות |  |
| 8:1.19 | במשלים |  |
| 8:1.20 | שאינם |  |
| 8:1.21 | מתקבלים |  |
| 8:1.22 | על |  |
| 8:1.23 | השכל |  |
| 8:1.24 | הישר | . |
| 8:1.25 | ואין |  |
| 8:1.26 | בהם |  |
| 8:1.27 | כדאי |  |
| 8:1.28 | לאוהב |  |
| 8:1.29 | את |  |
| 8:1.30 | יחיד |  |
| 8:1.31 | השי"ת |  |
| 8:1.32 | על |  |
| 8:1.33 | דרך |  |
| 8:1.34 | האמת |  |
| 8:1.35 | אשר |  |
| 8:1.36 | הורונו |  |
| 8:1.37 | רז"ל |  |
| 8:1.38 | על |  |
| 8:1.39 | פי |  |
| 8:1.40 | תורתינו |  |
| 8:1.41 | הקדושה | , |
| 8:1.42 | כמו |  |
| 8:1.43 | שיראה |  |
| 8:1.44 | המעיין |  |
| 8:1.45 | שם |  |
| 8:1.46 | במשל |  |
| 8:1.47 | שהביא |  |
| 8:1.48 | שם | ! |
| 8:1.49 | וגם |  |
| 8:1.50 | הוא |  |
| 8:1.51 | סובר |  |
| 8:1.52 | שכל |  |
| 8:1.53 | עבודתינו |  |
| 8:1.54 | ותפלותינו |  |
| 8:1.55 | וברכותינו |  |
| 8:1.56 | הם |  |
| 8:1.57 | לאלוה |  |
| 8:1.58 | קצר |  |
| 8:1.59 | אפים |  |
| 8:1.60 | זעיר |  |
| 8:1.61 | אנפין | ) |
| 8:1.62 | והוא |  |
| 8:1.63 | אלהינו |  |
| 8:1.64 | ואנחנו |  |
| 8:1.65 | עמו |  |
| 8:1.66 | ונחלתו | , |
| 8:1.67 | והוא |  |
| 8:1.68 | המלך |  |
| 8:1.69 | הקדוש |  |
| 8:1.70 | המהולל |  |
| 8:1.71 | בכל |  |
| 8:1.72 | התושבחות | . |
| 8:1.73 | השולט |  |
| 8:1.74 | על |  |
| 8:1.75 | כל |  |
| 8:1.76 | הברואים |  |
| 8:1.77 | וזנם |  |
| 8:1.78 | ומפרנסם |  |
| 8:1.79 | ומנהיגם | . |
| 8:1.80 | ושזעיר |  |
| 8:1.81 | אנפין |  |
| 8:1.82 | הוא |  |
| 8:1.83 | שירד |  |
| 8:1.84 | על |  |
| 8:1.85 | הר |  |
| 8:1.86 | סיני |  |
| 8:1.87 | להתחבר |  |
| 8:1.88 | עם |  |
| 8:1.89 | המלכות |  |
| 8:1.90 | שנקראת |  |
| 8:1.91 | אש | . |
| 8:1.92 | שכבר |  |
| 8:1.93 | קדמה |  |
| 8:1.94 | וירדה |  |
| 8:1.95 | על |  |
| 8:1.96 | הר |  |
| 8:1.97 | סיני | . |
| 8:1.98 | ואמר |  |
| 8:1.99 | אנכי |  |
| 8:1.100 | ה | ' |
| 8:1.101 | אלהיך |  |
| 8:1.102 | אשר |  |
| 8:1.103 | הוציאתיך |  |
| 8:1.104 | מארץ |  |
| 8:1.105 | מצרים | . |
| 8:1.106 | היא |  |
| 8:1.107 | אימא |  |
| 8:1.108 | עילאה |  |
| 8:1.109 | בינה | . |
| 8:1.110 | כי |  |
| 8:1.111 | היא |  |
| 8:1.112 | הוציאתם |  |
| 8:1.113 | ממצרים |  |
| 8:1.114 | על |  |
| 8:1.115 | די |  |
| 8:1.116 | המלכות |  |
| 8:1.117 | שנתנה |  |
| 8:1.118 | לה |  |
| 8:1.119 | גבורותיה | , |
| 8:1.120 | ומלכות |  |
| 8:1.121 | היא |  |
| 8:1.122 | המכה |  |
| 8:1.123 | במצרים |  |
| 8:1.124 | בסיועת |  |
| 8:1.125 | הבינה | , |
| 8:1.126 | וכן |  |
| 8:1.127 | בסיני |  |
| 8:1.128 | על |  |
| 8:1.129 | ידי |  |
| 8:1.130 | הבור |  |
| 8:1.131 | בינה |  |
| 8:1.132 | עם |  |
| 8:1.133 | תפארת |  |
| 8:1.134 | ומלכות |  |
| 8:1.135 | נאמר |  |
| 8:1.136 | אנכי |  |
| 8:1.137 | ה | ' |
| 8:1.138 | אלהיך | . |
| 8:1.139 | והתורה |  |
| 8:1.140 | יצאה |  |
| 8:1.141 | מחכמה |  |
| 8:1.142 | עלאה |  |
| 8:1.143 | עם |  |
| 8:1.144 | כללות |  |
| 8:1.145 | הבינה | . |
| 8:1.146 | והכל |  |
| 8:1.147 | בכח |  |
| 8:1.148 | אין |  |
| 8:1.149 | סוף |  |
| 8:1.150 | ב"ה |  |
| 8:1.151 | וכו | ' |
| 8:1.152 | ע"ש |  |
| 8:1.153 | דף |  |
| 8:1.154 | ס"א |  |
| 8:1.155 | וס"ב | . |
| 8:1.156 | ודבריו |  |
| 8:1.157 | ג"ד |  |
| 8:1.158 | מכיינים |  |
| 8:1.159 | לדברי |  |
| 8:1.160 | מהר"י |  |
| 8:1.161 | לופיס |  |
| 8:1.162 | לעיל |  |
| 8:1.163 | סי | ' |
| 8:1.164 | כ"ו |  |
| 8:1.165 | שזעיר |  |
| 8:1.166 | אנפין |  |
| 8:1.167 | הוא |  |
| 8:1.168 | שנגלה |  |
| 8:1.169 | על |  |
| 8:1.170 | אבותינו |  |
| 8:1.171 | בסיני |  |
| 8:1.172 | ואמר |  |
| 8:1.173 | אנכי |  |
| 8:1.174 | ה | ' |
| 8:1.175 | אלהיך | . |


## Milchamot ha-Shem 8:2

Hebrew (Wikisource): והנה כבר הודעתיך בסי' מ' שכל אין סוף שהזכירו המקובלים החדשים הוא החלק הגנוז בכתר הנקרא עתיקא דכל עתיקין בלבד, אבל עדיין יש אין סוף שהוא לעילא לעילא. והוא המקיף את כל אשר בתוכו מכל הנבראים ואין דבר חוצה לו, ורק אל האין סוף אשר בכתר מליהם ירמזון, וממנו הם ממשיכים בהתפללם אל קצר אפים (זעיר אנפין). ואל עתיקא קדישא לפי דבריהם אמר הנביא, ה' שהוא עתיק, פעלך שהוא זעיר אנפין חייהו וכו' כמו שנתבאר לעיל סוף סימן מ"ח ב' האדרא רבה דף קל"ח. והעולה מדברי הרב נח"י הוא הדבר שכבר זכרתי לך פעמים רבות. שדעת המקובלים החדשים שכל עבודתינו ותפלותינו וברכותינו להאלוה קצר אפים, והזעיר אנפין הוא שנגלה על אבותינו בהר סיני ונתן לנו את התורה, ומפוי שמעו אנכי ה' אלהיך, הוא ולא אבא ואימא, ולא אריך אנפין, ולא עתיק יומין, ולא אדם קדמון, וכ"ש שאינו האין סוף שהוא לעילא לעילא.

| id | word | punctuation after |
|---|---|---|
| 8:2.1 | והנה |  |
| 8:2.2 | כבר |  |
| 8:2.3 | הודעתיך |  |
| 8:2.4 | בסי | ' |
| 8:2.5 | מ | ' |
| 8:2.6 | שכל |  |
| 8:2.7 | אין |  |
| 8:2.8 | סוף |  |
| 8:2.9 | שהזכירו |  |
| 8:2.10 | המקובלים |  |
| 8:2.11 | החדשים |  |
| 8:2.12 | הוא |  |
| 8:2.13 | החלק |  |
| 8:2.14 | הגנוז |  |
| 8:2.15 | בכתר |  |
| 8:2.16 | הנקרא |  |
| 8:2.17 | עתיקא |  |
| 8:2.18 | דכל |  |
| 8:2.19 | עתיקין |  |
| 8:2.20 | בלבד | , |
| 8:2.21 | אבל |  |
| 8:2.22 | עדיין |  |
| 8:2.23 | יש |  |
| 8:2.24 | אין |  |
| 8:2.25 | סוף |  |
| 8:2.26 | שהוא |  |
| 8:2.27 | לעילא |  |
| 8:2.28 | לעילא | . |
| 8:2.29 | והוא |  |
| 8:2.30 | המקיף |  |
| 8:2.31 | את |  |
| 8:2.32 | כל |  |
| 8:2.33 | אשר |  |
| 8:2.34 | בתוכו |  |
| 8:2.35 | מכל |  |
| 8:2.36 | הנבראים |  |
| 8:2.37 | ואין |  |
| 8:2.38 | דבר |  |
| 8:2.39 | חוצה |  |
| 8:2.40 | לו | , |
| 8:2.41 | ורק |  |
| 8:2.42 | אל |  |
| 8:2.43 | האין |  |
| 8:2.44 | סוף |  |
| 8:2.45 | אשר |  |
| 8:2.46 | בכתר |  |
| 8:2.47 | מליהם |  |
| 8:2.48 | ירמזון | , |
| 8:2.49 | וממנו |  |
| 8:2.50 | הם |  |
| 8:2.51 | ממשיכים |  |
| 8:2.52 | בהתפללם |  |
| 8:2.53 | אל |  |
| 8:2.54 | קצר |  |
| 8:2.55 | אפים |  |
| 8:2.56 | זעיר |  |
| 8:2.57 | אנפין | ). |
| 8:2.58 | ואל |  |
| 8:2.59 | עתיקא |  |
| 8:2.60 | קדישא |  |
| 8:2.61 | לפי |  |
| 8:2.62 | דבריהם |  |
| 8:2.63 | אמר |  |
| 8:2.64 | הנביא | , |
| 8:2.65 | ה | ' |
| 8:2.66 | שהוא |  |
| 8:2.67 | עתיק | , |
| 8:2.68 | פעלך |  |
| 8:2.69 | שהוא |  |
| 8:2.70 | זעיר |  |
| 8:2.71 | אנפין |  |
| 8:2.72 | חייהו |  |
| 8:2.73 | וכו | ' |
| 8:2.74 | כמו |  |
| 8:2.75 | שנתבאר |  |
| 8:2.76 | לעיל |  |
| 8:2.77 | סוף |  |
| 8:2.78 | סימן |  |
| 8:2.79 | מ"ח |  |
| 8:2.80 | ב | ' |
| 8:2.81 | האדרא |  |
| 8:2.82 | רבה |  |
| 8:2.83 | דף |  |
| 8:2.84 | קל"ח | . |
| 8:2.85 | והעולה |  |
| 8:2.86 | מדברי |  |
| 8:2.87 | הרב |  |
| 8:2.88 | נח"י |  |
| 8:2.89 | הוא |  |
| 8:2.90 | הדבר |  |
| 8:2.91 | שכבר |  |
| 8:2.92 | זכרתי |  |
| 8:2.93 | לך |  |
| 8:2.94 | פעמים |  |
| 8:2.95 | רבות | . |
| 8:2.96 | שדעת |  |
| 8:2.97 | המקובלים |  |
| 8:2.98 | החדשים |  |
| 8:2.99 | שכל |  |
| 8:2.100 | עבודתינו |  |
| 8:2.101 | ותפלותינו |  |
| 8:2.102 | וברכותינו |  |
| 8:2.103 | להאלוה |  |
| 8:2.104 | קצר |  |
| 8:2.105 | אפים | , |
| 8:2.106 | והזעיר |  |
| 8:2.107 | אנפין |  |
| 8:2.108 | הוא |  |
| 8:2.109 | שנגלה |  |
| 8:2.110 | על |  |
| 8:2.111 | אבותינו |  |
| 8:2.112 | בהר |  |
| 8:2.113 | סיני |  |
| 8:2.114 | ונתן |  |
| 8:2.115 | לנו |  |
| 8:2.116 | את |  |
| 8:2.117 | התורה | , |
| 8:2.118 | ומפוי |  |
| 8:2.119 | שמעו |  |
| 8:2.120 | אנכי |  |
| 8:2.121 | ה | ' |
| 8:2.122 | אלהיך | , |
| 8:2.123 | הוא |  |
| 8:2.124 | ולא |  |
| 8:2.125 | אבא |  |
| 8:2.126 | ואימא | , |
| 8:2.127 | ולא |  |
| 8:2.128 | אריך |  |
| 8:2.129 | אנפין | , |
| 8:2.130 | ולא |  |
| 8:2.131 | עתיק |  |
| 8:2.132 | יומין | , |
| 8:2.133 | ולא |  |
| 8:2.134 | אדם |  |
| 8:2.135 | קדמון | , |
| 8:2.136 | וכ"ש |  |
| 8:2.137 | שאינו |  |
| 8:2.138 | האין |  |
| 8:2.139 | סוף |  |
| 8:2.140 | שהוא |  |
| 8:2.141 | לעילא |  |
| 8:2.142 | לעילא | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

