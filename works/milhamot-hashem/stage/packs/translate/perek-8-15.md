# Translation pack: milhamot-hashem perek-8-15 (Milchamot ha-Shem 8:18–8:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-15`.

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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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
- The words of a verse quoted in the plural (Job 32:16, עמדו): plural, no
  subject added.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 8:18

Hebrew (Wikisource): קיב) וכשעלו מן הגולה עמדו אנשי כנסת הגדולה והרבו תפלה ותחנונים לפני הקב"ה וצעקו בקול גדול בייא בייא היינו דאחרביה לבי מקדשא וכו' ובקשו מהשם ית' לסייעם לעקור יצרא דע"ז מישראל שהוא ניהו דאחרביה לבי מקדשא וקלייה להיכלא וכו' ונמסר בידם. היינו שניתן להם סיוע מן השמים וקבעו לישראל בתי מדרשות הרבה. ומינו בהם אנשים חכמים ונבונים ללמד לעם דרכי תורתינו הקדושה, ולהרחיקם מהאמונות הזרות הרחוקות מדרכי התורה. ונפק כי גוריא דנורא, שיצאה וסרה מהם ההתלהבות והחשק והתשוקה לרמזי האמונות הכוזבות שהיו מעמיסים בבית קדש הקדשים. היא תורתינו הקדושה שהיא לנו קדש קדשים המלמדת לאדם דעת ובינה והשכל לדעת את ה' ולאהבה אותו וליראה ממנו, ולעבדו לבדו, לפי שמתחלה היו ממציאים להאמונות הטפלות רמזים וסמוכות דלות ורעות. ותולים אותם בדרכים שונים הרחוקים מקבלת רז"ל שקבלו ממשה רבינו ע"ה דרכי האמונה האמתית שמחוייב כל בעל דת ישראל להאמין כי הבורא ית' שמו הוא אחד ומיוחד מכל שאר האחדים. הוא ראשון והוא אחרון שאיננו נופל תחת הזמן והמקום. ושלא ישיגוהו משיגי הגוף. ושהוא קדמון ונצחי בלתי בעל תכלית. הן מצד הקודם, הן מצד המתאחר כאשר קרא הנביא ישעיה בשם ה' אני ראשון ואני אחרון. וכן תרגם אונקלוס ואמרתי חי אנכי לעולם. קיים אנא לעלמין בלשון רבים ולא תרגם לעלם בלשון יחיד, להורות שהוא ית' קיים תמיד על ענין ואופן אחד בלי שום שינוי ורבוי כלל חלילה. קודם שנברא העולם ואחר שיכלה אם יחפוץ השי"ת (עיין במורה פרק י"ז לשני והקדמה כוללת לרבי נתן אדלר בעל נתינה לגר). והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי. שר"ל אני הוא אשר הייתי ואני אשר אהיה עד בלתי שמים. תמיד אני קיים בתואר אחד בלי שום שינוי. ושהוא ית' הסבה הראשונה ולו לבדו הוא שראוי לעבוד ולהתפלל, ולהודיע גדולתו, ולעשות מצותיו ושלא יעשה כן לשום נברא מכל הנבראים גשמי או רוחני. אלו הן הדעות הרמות המזוקקות צרופות וטהורות. כראי מוצק חזקות. לא כדעת המקובלים החדשים, אשר בעונות הרבים עמודי הדת ירופפו ואדני תורה שבעל פה יתפלצון על ידי הקבלה המזוייפת. לעבוד ולהתפלל לסבות נבראות לקצר אפים (זעיר אנפין) אף על פי שהוא נברא כמבואר בזהר ומפרשיו ובספר הברית ויושר לבב ונחלת יוסף ומאמר האמונה האמתית למהר"י לופיס וזולתם.

| id | word | punctuation after |
|---|---|---|
| 8:18.1 | קיב | ) |
| 8:18.2 | וכשעלו |  |
| 8:18.3 | מן |  |
| 8:18.4 | הגולה |  |
| 8:18.5 | עמדו |  |
| 8:18.6 | אנשי |  |
| 8:18.7 | כנסת |  |
| 8:18.8 | הגדולה |  |
| 8:18.9 | והרבו |  |
| 8:18.10 | תפלה |  |
| 8:18.11 | ותחנונים |  |
| 8:18.12 | לפני |  |
| 8:18.13 | הקב"ה |  |
| 8:18.14 | וצעקו |  |
| 8:18.15 | בקול |  |
| 8:18.16 | גדול |  |
| 8:18.17 | בייא |  |
| 8:18.18 | בייא |  |
| 8:18.19 | היינו |  |
| 8:18.20 | דאחרביה |  |
| 8:18.21 | לבי |  |
| 8:18.22 | מקדשא |  |
| 8:18.23 | וכו | ' |
| 8:18.24 | ובקשו |  |
| 8:18.25 | מהשם |  |
| 8:18.26 | ית | ' |
| 8:18.27 | לסייעם |  |
| 8:18.28 | לעקור |  |
| 8:18.29 | יצרא |  |
| 8:18.30 | דע"ז |  |
| 8:18.31 | מישראל |  |
| 8:18.32 | שהוא |  |
| 8:18.33 | ניהו |  |
| 8:18.34 | דאחרביה |  |
| 8:18.35 | לבי |  |
| 8:18.36 | מקדשא |  |
| 8:18.37 | וקלייה |  |
| 8:18.38 | להיכלא |  |
| 8:18.39 | וכו | ' |
| 8:18.40 | ונמסר |  |
| 8:18.41 | בידם | . |
| 8:18.42 | היינו |  |
| 8:18.43 | שניתן |  |
| 8:18.44 | להם |  |
| 8:18.45 | סיוע |  |
| 8:18.46 | מן |  |
| 8:18.47 | השמים |  |
| 8:18.48 | וקבעו |  |
| 8:18.49 | לישראל |  |
| 8:18.50 | בתי |  |
| 8:18.51 | מדרשות |  |
| 8:18.52 | הרבה | . |
| 8:18.53 | ומינו |  |
| 8:18.54 | בהם |  |
| 8:18.55 | אנשים |  |
| 8:18.56 | חכמים |  |
| 8:18.57 | ונבונים |  |
| 8:18.58 | ללמד |  |
| 8:18.59 | לעם |  |
| 8:18.60 | דרכי |  |
| 8:18.61 | תורתינו |  |
| 8:18.62 | הקדושה | , |
| 8:18.63 | ולהרחיקם |  |
| 8:18.64 | מהאמונות |  |
| 8:18.65 | הזרות |  |
| 8:18.66 | הרחוקות |  |
| 8:18.67 | מדרכי |  |
| 8:18.68 | התורה | . |
| 8:18.69 | ונפק |  |
| 8:18.70 | כי |  |
| 8:18.71 | גוריא |  |
| 8:18.72 | דנורא | , |
| 8:18.73 | שיצאה |  |
| 8:18.74 | וסרה |  |
| 8:18.75 | מהם |  |
| 8:18.76 | ההתלהבות |  |
| 8:18.77 | והחשק |  |
| 8:18.78 | והתשוקה |  |
| 8:18.79 | לרמזי |  |
| 8:18.80 | האמונות |  |
| 8:18.81 | הכוזבות |  |
| 8:18.82 | שהיו |  |
| 8:18.83 | מעמיסים |  |
| 8:18.84 | בבית |  |
| 8:18.85 | קדש |  |
| 8:18.86 | הקדשים | . |
| 8:18.87 | היא |  |
| 8:18.88 | תורתינו |  |
| 8:18.89 | הקדושה |  |
| 8:18.90 | שהיא |  |
| 8:18.91 | לנו |  |
| 8:18.92 | קדש |  |
| 8:18.93 | קדשים |  |
| 8:18.94 | המלמדת |  |
| 8:18.95 | לאדם |  |
| 8:18.96 | דעת |  |
| 8:18.97 | ובינה |  |
| 8:18.98 | והשכל |  |
| 8:18.99 | לדעת |  |
| 8:18.100 | את |  |
| 8:18.101 | ה | ' |
| 8:18.102 | ולאהבה |  |
| 8:18.103 | אותו |  |
| 8:18.104 | וליראה |  |
| 8:18.105 | ממנו | , |
| 8:18.106 | ולעבדו |  |
| 8:18.107 | לבדו | , |
| 8:18.108 | לפי |  |
| 8:18.109 | שמתחלה |  |
| 8:18.110 | היו |  |
| 8:18.111 | ממציאים |  |
| 8:18.112 | להאמונות |  |
| 8:18.113 | הטפלות |  |
| 8:18.114 | רמזים |  |
| 8:18.115 | וסמוכות |  |
| 8:18.116 | דלות |  |
| 8:18.117 | ורעות | . |
| 8:18.118 | ותולים |  |
| 8:18.119 | אותם |  |
| 8:18.120 | בדרכים |  |
| 8:18.121 | שונים |  |
| 8:18.122 | הרחוקים |  |
| 8:18.123 | מקבלת |  |
| 8:18.124 | רז"ל |  |
| 8:18.125 | שקבלו |  |
| 8:18.126 | ממשה |  |
| 8:18.127 | רבינו |  |
| 8:18.128 | ע"ה |  |
| 8:18.129 | דרכי |  |
| 8:18.130 | האמונה |  |
| 8:18.131 | האמתית |  |
| 8:18.132 | שמחוייב |  |
| 8:18.133 | כל |  |
| 8:18.134 | בעל |  |
| 8:18.135 | דת |  |
| 8:18.136 | ישראל |  |
| 8:18.137 | להאמין |  |
| 8:18.138 | כי |  |
| 8:18.139 | הבורא |  |
| 8:18.140 | ית | ' |
| 8:18.141 | שמו |  |
| 8:18.142 | הוא |  |
| 8:18.143 | אחד |  |
| 8:18.144 | ומיוחד |  |
| 8:18.145 | מכל |  |
| 8:18.146 | שאר |  |
| 8:18.147 | האחדים | . |
| 8:18.148 | הוא |  |
| 8:18.149 | ראשון |  |
| 8:18.150 | והוא |  |
| 8:18.151 | אחרון |  |
| 8:18.152 | שאיננו |  |
| 8:18.153 | נופל |  |
| 8:18.154 | תחת |  |
| 8:18.155 | הזמן |  |
| 8:18.156 | והמקום | . |
| 8:18.157 | ושלא |  |
| 8:18.158 | ישיגוהו |  |
| 8:18.159 | משיגי |  |
| 8:18.160 | הגוף | . |
| 8:18.161 | ושהוא |  |
| 8:18.162 | קדמון |  |
| 8:18.163 | ונצחי |  |
| 8:18.164 | בלתי |  |
| 8:18.165 | בעל |  |
| 8:18.166 | תכלית | . |
| 8:18.167 | הן |  |
| 8:18.168 | מצד |  |
| 8:18.169 | הקודם | , |
| 8:18.170 | הן |  |
| 8:18.171 | מצד |  |
| 8:18.172 | המתאחר |  |
| 8:18.173 | כאשר |  |
| 8:18.174 | קרא |  |
| 8:18.175 | הנביא |  |
| 8:18.176 | ישעיה |  |
| 8:18.177 | בשם |  |
| 8:18.178 | ה | ' |
| 8:18.179 | אני |  |
| 8:18.180 | ראשון |  |
| 8:18.181 | ואני |  |
| 8:18.182 | אחרון | . |
| 8:18.183 | וכן |  |
| 8:18.184 | תרגם |  |
| 8:18.185 | אונקלוס |  |
| 8:18.186 | ואמרתי |  |
| 8:18.187 | חי |  |
| 8:18.188 | אנכי |  |
| 8:18.189 | לעולם | . |
| 8:18.190 | קיים |  |
| 8:18.191 | אנא |  |
| 8:18.192 | לעלמין |  |
| 8:18.193 | בלשון |  |
| 8:18.194 | רבים |  |
| 8:18.195 | ולא |  |
| 8:18.196 | תרגם |  |
| 8:18.197 | לעלם |  |
| 8:18.198 | בלשון |  |
| 8:18.199 | יחיד | , |
| 8:18.200 | להורות |  |
| 8:18.201 | שהוא |  |
| 8:18.202 | ית | ' |
| 8:18.203 | קיים |  |
| 8:18.204 | תמיד |  |
| 8:18.205 | על |  |
| 8:18.206 | ענין |  |
| 8:18.207 | ואופן |  |
| 8:18.208 | אחד |  |
| 8:18.209 | בלי |  |
| 8:18.210 | שום |  |
| 8:18.211 | שינוי |  |
| 8:18.212 | ורבוי |  |
| 8:18.213 | כלל |  |
| 8:18.214 | חלילה | . |
| 8:18.215 | קודם |  |
| 8:18.216 | שנברא |  |
| 8:18.217 | העולם |  |
| 8:18.218 | ואחר |  |
| 8:18.219 | שיכלה |  |
| 8:18.220 | אם |  |
| 8:18.221 | יחפוץ |  |
| 8:18.222 | השי"ת |  |
| 8:18.223 | עיין |  |
| 8:18.224 | במורה |  |
| 8:18.225 | פרק |  |
| 8:18.226 | י"ז |  |
| 8:18.227 | לשני |  |
| 8:18.228 | והקדמה |  |
| 8:18.229 | כוללת |  |
| 8:18.230 | לרבי |  |
| 8:18.231 | נתן |  |
| 8:18.232 | אדלר |  |
| 8:18.233 | בעל |  |
| 8:18.234 | נתינה |  |
| 8:18.235 | לגר | ). |
| 8:18.236 | והוא |  |
| 8:18.237 | שאמר |  |
| 8:18.238 | ראו |  |
| 8:18.239 | עתה |  |
| 8:18.240 | כי |  |
| 8:18.241 | אני |  |
| 8:18.242 | אני |  |
| 8:18.243 | הוא |  |
| 8:18.244 | ואין |  |
| 8:18.245 | אלהים |  |
| 8:18.246 | עמדי | . |
| 8:18.247 | שר"ל |  |
| 8:18.248 | אני |  |
| 8:18.249 | הוא |  |
| 8:18.250 | אשר |  |
| 8:18.251 | הייתי |  |
| 8:18.252 | ואני |  |
| 8:18.253 | אשר |  |
| 8:18.254 | אהיה |  |
| 8:18.255 | עד |  |
| 8:18.256 | בלתי |  |
| 8:18.257 | שמים | . |
| 8:18.258 | תמיד |  |
| 8:18.259 | אני |  |
| 8:18.260 | קיים |  |
| 8:18.261 | בתואר |  |
| 8:18.262 | אחד |  |
| 8:18.263 | בלי |  |
| 8:18.264 | שום |  |
| 8:18.265 | שינוי | . |
| 8:18.266 | ושהוא |  |
| 8:18.267 | ית | ' |
| 8:18.268 | הסבה |  |
| 8:18.269 | הראשונה |  |
| 8:18.270 | ולו |  |
| 8:18.271 | לבדו |  |
| 8:18.272 | הוא |  |
| 8:18.273 | שראוי |  |
| 8:18.274 | לעבוד |  |
| 8:18.275 | ולהתפלל | , |
| 8:18.276 | ולהודיע |  |
| 8:18.277 | גדולתו | , |
| 8:18.278 | ולעשות |  |
| 8:18.279 | מצותיו |  |
| 8:18.280 | ושלא |  |
| 8:18.281 | יעשה |  |
| 8:18.282 | כן |  |
| 8:18.283 | לשום |  |
| 8:18.284 | נברא |  |
| 8:18.285 | מכל |  |
| 8:18.286 | הנבראים |  |
| 8:18.287 | גשמי |  |
| 8:18.288 | או |  |
| 8:18.289 | רוחני | . |
| 8:18.290 | אלו |  |
| 8:18.291 | הן |  |
| 8:18.292 | הדעות |  |
| 8:18.293 | הרמות |  |
| 8:18.294 | המזוקקות |  |
| 8:18.295 | צרופות |  |
| 8:18.296 | וטהורות | . |
| 8:18.297 | כראי |  |
| 8:18.298 | מוצק |  |
| 8:18.299 | חזקות | . |
| 8:18.300 | לא |  |
| 8:18.301 | כדעת |  |
| 8:18.302 | המקובלים |  |
| 8:18.303 | החדשים | , |
| 8:18.304 | אשר |  |
| 8:18.305 | בעונות |  |
| 8:18.306 | הרבים |  |
| 8:18.307 | עמודי |  |
| 8:18.308 | הדת |  |
| 8:18.309 | ירופפו |  |
| 8:18.310 | ואדני |  |
| 8:18.311 | תורה |  |
| 8:18.312 | שבעל |  |
| 8:18.313 | פה |  |
| 8:18.314 | יתפלצון |  |
| 8:18.315 | על |  |
| 8:18.316 | ידי |  |
| 8:18.317 | הקבלה |  |
| 8:18.318 | המזוייפת | . |
| 8:18.319 | לעבוד |  |
| 8:18.320 | ולהתפלל |  |
| 8:18.321 | לסבות |  |
| 8:18.322 | נבראות |  |
| 8:18.323 | לקצר |  |
| 8:18.324 | אפים |  |
| 8:18.325 | זעיר |  |
| 8:18.326 | אנפין | ) |
| 8:18.327 | אף |  |
| 8:18.328 | על |  |
| 8:18.329 | פי |  |
| 8:18.330 | שהוא |  |
| 8:18.331 | נברא |  |
| 8:18.332 | כמבואר |  |
| 8:18.333 | בזהר |  |
| 8:18.334 | ומפרשיו |  |
| 8:18.335 | ובספר |  |
| 8:18.336 | הברית |  |
| 8:18.337 | ויושר |  |
| 8:18.338 | לבב |  |
| 8:18.339 | ונחלת |  |
| 8:18.340 | יוסף |  |
| 8:18.341 | ומאמר |  |
| 8:18.342 | האמונה |  |
| 8:18.343 | האמתית |  |
| 8:18.344 | למהר"י |  |
| 8:18.345 | לופיס |  |
| 8:18.346 | וזולתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

