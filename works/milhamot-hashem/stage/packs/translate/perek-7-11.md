# Translation pack: milhamot-hashem perek-7-11 (Milchamot ha-Shem 7:15–7:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-11`.

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

- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
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
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
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
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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


## Milchamot ha-Shem 7:15

Hebrew (Wikisource): פח) העיקר הרביעי הקדמות. צריך להאמין כי זה האחד האו קדמון במלה מוחלטת וכל נמצא זולתו אינו קדמון כשנעריכהו אליו. והראיות על זה העיקר רבו בספרים, וזה העיקר הרביעי נזכר בתורה מעונה אלהי קדם ע"כ. כבר ביארתי לעיל סי' מ"ח שגם זה העיקר מיאנו בו כת המקובלים החדשים כמו שנתבאר בענין פיוט יגדל אלהים חי שמיאנו המקובלים לאמרו, ואל תשמי לדברי האומרים שנמנעו מלאמרו מטעם שאין לעשות לתורה עיקרים לפי שכל מצוה ומצוה של תורה היא עיקר לעצמה וכולה עיקרים, כי לא מחכמה אמרו ובאין מבין בעיקרים דברו כן שהרי זה נכלל בעיקר השמיני, ומובאר בהדיא שם שכל התורה כולב מפי הגבורה, ומשה רבינו כמו סופר, הקב"ה אומר, והוא אומר וכותב כל מה שאמר לו, מספר שני הדורות שעברו, וספוריה, ומצותיה, ולפיכך נקרא משה רבינו מחוקק, ואין הפרש בין בני חם וכוש ומצרים וכו' ושם אשתו מהטבאל בת מטרד ותמנע היתה פילגש, ובין אנכי ה' אלהיך, ושמע ישראל וכו' הכל מפי הגבורה, הכל תורת ה' תמימה טהורה קדושה אמת, ורז"ך היה אצלם משנה כופר ומכחש יותר מכל מין וכופר על שחשב שיש בתורה מוח וקליפה, ושחשבו השנים שבה והספורים אין בהם תועלת, ושהם מפי משה עצמו, והוא ענין אין תורה מן השמים, וכך אמרו חכמים, האומר כל שתורה כולה מפי הגבורה חוץ מפסוק זה שלא אמרו הקב"ה אלא משה מפי עצמו, זהו דבר ה' בזה אלא כל ענייניה ומצותיה מפי הגבורה, ומהם ישיג מי שזכהו השי"ת עניינים נפלעים, כמאמר דוד המלך ע"ה כל עיני ואביטה נפלאות מתורתיך והלכות סוכה ולולב ושופר וציצית ותפלין עם כל חילוקי דיניהם הם וזולתם, הם בעצמם ההלכות שקיבל משה רבי' ע"ה מהקב"ה בסיני. והוא אמרם לנו. והוא נאמן בשליחותו. והמאמר המורה על זה הוא מה שנאמר בזאת תדעון כי ה' שלחני לעשות את כל המעשים האלה כי לא מלבי ע"כ. הא למדת שטענתם זו כוזבת. שהרי רבינו מודה בזה וכלל כל מצותיה של תורה וספוריה בזה העיקר השמיני באמרו שכל מצות התורה וחילוקי דיניהם וסיפוריה כולם מפי הגבורה. וכל הכופר בשום חלק מכל מצות התורה או אפי' מספוריה הרי הוא מין וכופר. וכולם נכללים תחת סוג תורת אמת נתן לעמו אל. על יד נביאו, ואין דבר בה לבטלה. וכל מי שאינו מודה באחת מהן הוא מין וכופר ואפיקירוס כמנשה וסיעתו. והראיה הגדולה על מה שאמרתי שאינם מודים בארבעה עיקרים הראשונים שהרי ביאר הרב חמדת ימים שלא חרה להאר"י רק על ארבעה בתים ראשונים שביגדל שהם סותרים אמונתם. אבל מהנו אדון עולם ואילך שפיר דמי לדידהו לאמרו כנז' לעיל סי' מ"ח.

| id | word | punctuation after |
|---|---|---|
| 7:15.1 | פח | ) |
| 7:15.2 | העיקר |  |
| 7:15.3 | הרביעי |  |
| 7:15.4 | הקדמות | . |
| 7:15.5 | צריך |  |
| 7:15.6 | להאמין |  |
| 7:15.7 | כי |  |
| 7:15.8 | זה |  |
| 7:15.9 | האחד |  |
| 7:15.10 | האו |  |
| 7:15.11 | קדמון |  |
| 7:15.12 | במלה |  |
| 7:15.13 | מוחלטת |  |
| 7:15.14 | וכל |  |
| 7:15.15 | נמצא |  |
| 7:15.16 | זולתו |  |
| 7:15.17 | אינו |  |
| 7:15.18 | קדמון |  |
| 7:15.19 | כשנעריכהו |  |
| 7:15.20 | אליו | . |
| 7:15.21 | והראיות |  |
| 7:15.22 | על |  |
| 7:15.23 | זה |  |
| 7:15.24 | העיקר |  |
| 7:15.25 | רבו |  |
| 7:15.26 | בספרים | , |
| 7:15.27 | וזה |  |
| 7:15.28 | העיקר |  |
| 7:15.29 | הרביעי |  |
| 7:15.30 | נזכר |  |
| 7:15.31 | בתורה |  |
| 7:15.32 | מעונה |  |
| 7:15.33 | אלהי |  |
| 7:15.34 | קדם |  |
| 7:15.35 | ע"כ | . |
| 7:15.36 | כבר |  |
| 7:15.37 | ביארתי |  |
| 7:15.38 | לעיל |  |
| 7:15.39 | סי | ' |
| 7:15.40 | מ"ח |  |
| 7:15.41 | שגם |  |
| 7:15.42 | זה |  |
| 7:15.43 | העיקר |  |
| 7:15.44 | מיאנו |  |
| 7:15.45 | בו |  |
| 7:15.46 | כת |  |
| 7:15.47 | המקובלים |  |
| 7:15.48 | החדשים |  |
| 7:15.49 | כמו |  |
| 7:15.50 | שנתבאר |  |
| 7:15.51 | בענין |  |
| 7:15.52 | פיוט |  |
| 7:15.53 | יגדל |  |
| 7:15.54 | אלהים |  |
| 7:15.55 | חי |  |
| 7:15.56 | שמיאנו |  |
| 7:15.57 | המקובלים |  |
| 7:15.58 | לאמרו | , |
| 7:15.59 | ואל |  |
| 7:15.60 | תשמי |  |
| 7:15.61 | לדברי |  |
| 7:15.62 | האומרים |  |
| 7:15.63 | שנמנעו |  |
| 7:15.64 | מלאמרו |  |
| 7:15.65 | מטעם |  |
| 7:15.66 | שאין |  |
| 7:15.67 | לעשות |  |
| 7:15.68 | לתורה |  |
| 7:15.69 | עיקרים |  |
| 7:15.70 | לפי |  |
| 7:15.71 | שכל |  |
| 7:15.72 | מצוה |  |
| 7:15.73 | ומצוה |  |
| 7:15.74 | של |  |
| 7:15.75 | תורה |  |
| 7:15.76 | היא |  |
| 7:15.77 | עיקר |  |
| 7:15.78 | לעצמה |  |
| 7:15.79 | וכולה |  |
| 7:15.80 | עיקרים | , |
| 7:15.81 | כי |  |
| 7:15.82 | לא |  |
| 7:15.83 | מחכמה |  |
| 7:15.84 | אמרו |  |
| 7:15.85 | ובאין |  |
| 7:15.86 | מבין |  |
| 7:15.87 | בעיקרים |  |
| 7:15.88 | דברו |  |
| 7:15.89 | כן |  |
| 7:15.90 | שהרי |  |
| 7:15.91 | זה |  |
| 7:15.92 | נכלל |  |
| 7:15.93 | בעיקר |  |
| 7:15.94 | השמיני | , |
| 7:15.95 | ומובאר |  |
| 7:15.96 | בהדיא |  |
| 7:15.97 | שם |  |
| 7:15.98 | שכל |  |
| 7:15.99 | התורה |  |
| 7:15.100 | כולב |  |
| 7:15.101 | מפי |  |
| 7:15.102 | הגבורה | , |
| 7:15.103 | ומשה |  |
| 7:15.104 | רבינו |  |
| 7:15.105 | כמו |  |
| 7:15.106 | סופר | , |
| 7:15.107 | הקב"ה |  |
| 7:15.108 | אומר | , |
| 7:15.109 | והוא |  |
| 7:15.110 | אומר |  |
| 7:15.111 | וכותב |  |
| 7:15.112 | כל |  |
| 7:15.113 | מה |  |
| 7:15.114 | שאמר |  |
| 7:15.115 | לו | , |
| 7:15.116 | מספר |  |
| 7:15.117 | שני |  |
| 7:15.118 | הדורות |  |
| 7:15.119 | שעברו | , |
| 7:15.120 | וספוריה | , |
| 7:15.121 | ומצותיה | , |
| 7:15.122 | ולפיכך |  |
| 7:15.123 | נקרא |  |
| 7:15.124 | משה |  |
| 7:15.125 | רבינו |  |
| 7:15.126 | מחוקק | , |
| 7:15.127 | ואין |  |
| 7:15.128 | הפרש |  |
| 7:15.129 | בין |  |
| 7:15.130 | בני |  |
| 7:15.131 | חם |  |
| 7:15.132 | וכוש |  |
| 7:15.133 | ומצרים |  |
| 7:15.134 | וכו | ' |
| 7:15.135 | ושם |  |
| 7:15.136 | אשתו |  |
| 7:15.137 | מהטבאל |  |
| 7:15.138 | בת |  |
| 7:15.139 | מטרד |  |
| 7:15.140 | ותמנע |  |
| 7:15.141 | היתה |  |
| 7:15.142 | פילגש | , |
| 7:15.143 | ובין |  |
| 7:15.144 | אנכי |  |
| 7:15.145 | ה | ' |
| 7:15.146 | אלהיך | , |
| 7:15.147 | ושמע |  |
| 7:15.148 | ישראל |  |
| 7:15.149 | וכו | ' |
| 7:15.150 | הכל |  |
| 7:15.151 | מפי |  |
| 7:15.152 | הגבורה | , |
| 7:15.153 | הכל |  |
| 7:15.154 | תורת |  |
| 7:15.155 | ה | ' |
| 7:15.156 | תמימה |  |
| 7:15.157 | טהורה |  |
| 7:15.158 | קדושה |  |
| 7:15.159 | אמת | , |
| 7:15.160 | ורז"ך |  |
| 7:15.161 | היה |  |
| 7:15.162 | אצלם |  |
| 7:15.163 | משנה |  |
| 7:15.164 | כופר |  |
| 7:15.165 | ומכחש |  |
| 7:15.166 | יותר |  |
| 7:15.167 | מכל |  |
| 7:15.168 | מין |  |
| 7:15.169 | וכופר |  |
| 7:15.170 | על |  |
| 7:15.171 | שחשב |  |
| 7:15.172 | שיש |  |
| 7:15.173 | בתורה |  |
| 7:15.174 | מוח |  |
| 7:15.175 | וקליפה | , |
| 7:15.176 | ושחשבו |  |
| 7:15.177 | השנים |  |
| 7:15.178 | שבה |  |
| 7:15.179 | והספורים |  |
| 7:15.180 | אין |  |
| 7:15.181 | בהם |  |
| 7:15.182 | תועלת | , |
| 7:15.183 | ושהם |  |
| 7:15.184 | מפי |  |
| 7:15.185 | משה |  |
| 7:15.186 | עצמו | , |
| 7:15.187 | והוא |  |
| 7:15.188 | ענין |  |
| 7:15.189 | אין |  |
| 7:15.190 | תורה |  |
| 7:15.191 | מן |  |
| 7:15.192 | השמים | , |
| 7:15.193 | וכך |  |
| 7:15.194 | אמרו |  |
| 7:15.195 | חכמים | , |
| 7:15.196 | האומר |  |
| 7:15.197 | כל |  |
| 7:15.198 | שתורה |  |
| 7:15.199 | כולה |  |
| 7:15.200 | מפי |  |
| 7:15.201 | הגבורה |  |
| 7:15.202 | חוץ |  |
| 7:15.203 | מפסוק |  |
| 7:15.204 | זה |  |
| 7:15.205 | שלא |  |
| 7:15.206 | אמרו |  |
| 7:15.207 | הקב"ה |  |
| 7:15.208 | אלא |  |
| 7:15.209 | משה |  |
| 7:15.210 | מפי |  |
| 7:15.211 | עצמו | , |
| 7:15.212 | זהו |  |
| 7:15.213 | דבר |  |
| 7:15.214 | ה | ' |
| 7:15.215 | בזה |  |
| 7:15.216 | אלא |  |
| 7:15.217 | כל |  |
| 7:15.218 | ענייניה |  |
| 7:15.219 | ומצותיה |  |
| 7:15.220 | מפי |  |
| 7:15.221 | הגבורה | , |
| 7:15.222 | ומהם |  |
| 7:15.223 | ישיג |  |
| 7:15.224 | מי |  |
| 7:15.225 | שזכהו |  |
| 7:15.226 | השי"ת |  |
| 7:15.227 | עניינים |  |
| 7:15.228 | נפלעים | , |
| 7:15.229 | כמאמר |  |
| 7:15.230 | דוד |  |
| 7:15.231 | המלך |  |
| 7:15.232 | ע"ה |  |
| 7:15.233 | כל |  |
| 7:15.234 | עיני |  |
| 7:15.235 | ואביטה |  |
| 7:15.236 | נפלאות |  |
| 7:15.237 | מתורתיך |  |
| 7:15.238 | והלכות |  |
| 7:15.239 | סוכה |  |
| 7:15.240 | ולולב |  |
| 7:15.241 | ושופר |  |
| 7:15.242 | וציצית |  |
| 7:15.243 | ותפלין |  |
| 7:15.244 | עם |  |
| 7:15.245 | כל |  |
| 7:15.246 | חילוקי |  |
| 7:15.247 | דיניהם |  |
| 7:15.248 | הם |  |
| 7:15.249 | וזולתם | , |
| 7:15.250 | הם |  |
| 7:15.251 | בעצמם |  |
| 7:15.252 | ההלכות |  |
| 7:15.253 | שקיבל |  |
| 7:15.254 | משה |  |
| 7:15.255 | רבי | ' |
| 7:15.256 | ע"ה |  |
| 7:15.257 | מהקב"ה |  |
| 7:15.258 | בסיני | . |
| 7:15.259 | והוא |  |
| 7:15.260 | אמרם |  |
| 7:15.261 | לנו | . |
| 7:15.262 | והוא |  |
| 7:15.263 | נאמן |  |
| 7:15.264 | בשליחותו | . |
| 7:15.265 | והמאמר |  |
| 7:15.266 | המורה |  |
| 7:15.267 | על |  |
| 7:15.268 | זה |  |
| 7:15.269 | הוא |  |
| 7:15.270 | מה |  |
| 7:15.271 | שנאמר |  |
| 7:15.272 | בזאת |  |
| 7:15.273 | תדעון |  |
| 7:15.274 | כי |  |
| 7:15.275 | ה | ' |
| 7:15.276 | שלחני |  |
| 7:15.277 | לעשות |  |
| 7:15.278 | את |  |
| 7:15.279 | כל |  |
| 7:15.280 | המעשים |  |
| 7:15.281 | האלה |  |
| 7:15.282 | כי |  |
| 7:15.283 | לא |  |
| 7:15.284 | מלבי |  |
| 7:15.285 | ע"כ | . |
| 7:15.286 | הא |  |
| 7:15.287 | למדת |  |
| 7:15.288 | שטענתם |  |
| 7:15.289 | זו |  |
| 7:15.290 | כוזבת | . |
| 7:15.291 | שהרי |  |
| 7:15.292 | רבינו |  |
| 7:15.293 | מודה |  |
| 7:15.294 | בזה |  |
| 7:15.295 | וכלל |  |
| 7:15.296 | כל |  |
| 7:15.297 | מצותיה |  |
| 7:15.298 | של |  |
| 7:15.299 | תורה |  |
| 7:15.300 | וספוריה |  |
| 7:15.301 | בזה |  |
| 7:15.302 | העיקר |  |
| 7:15.303 | השמיני |  |
| 7:15.304 | באמרו |  |
| 7:15.305 | שכל |  |
| 7:15.306 | מצות |  |
| 7:15.307 | התורה |  |
| 7:15.308 | וחילוקי |  |
| 7:15.309 | דיניהם |  |
| 7:15.310 | וסיפוריה |  |
| 7:15.311 | כולם |  |
| 7:15.312 | מפי |  |
| 7:15.313 | הגבורה | . |
| 7:15.314 | וכל |  |
| 7:15.315 | הכופר |  |
| 7:15.316 | בשום |  |
| 7:15.317 | חלק |  |
| 7:15.318 | מכל |  |
| 7:15.319 | מצות |  |
| 7:15.320 | התורה |  |
| 7:15.321 | או |  |
| 7:15.322 | אפי | ' |
| 7:15.323 | מספוריה |  |
| 7:15.324 | הרי |  |
| 7:15.325 | הוא |  |
| 7:15.326 | מין |  |
| 7:15.327 | וכופר | . |
| 7:15.328 | וכולם |  |
| 7:15.329 | נכללים |  |
| 7:15.330 | תחת |  |
| 7:15.331 | סוג |  |
| 7:15.332 | תורת |  |
| 7:15.333 | אמת |  |
| 7:15.334 | נתן |  |
| 7:15.335 | לעמו |  |
| 7:15.336 | אל | . |
| 7:15.337 | על |  |
| 7:15.338 | יד |  |
| 7:15.339 | נביאו | , |
| 7:15.340 | ואין |  |
| 7:15.341 | דבר |  |
| 7:15.342 | בה |  |
| 7:15.343 | לבטלה | . |
| 7:15.344 | וכל |  |
| 7:15.345 | מי |  |
| 7:15.346 | שאינו |  |
| 7:15.347 | מודה |  |
| 7:15.348 | באחת |  |
| 7:15.349 | מהן |  |
| 7:15.350 | הוא |  |
| 7:15.351 | מין |  |
| 7:15.352 | וכופר |  |
| 7:15.353 | ואפיקירוס |  |
| 7:15.354 | כמנשה |  |
| 7:15.355 | וסיעתו | . |
| 7:15.356 | והראיה |  |
| 7:15.357 | הגדולה |  |
| 7:15.358 | על |  |
| 7:15.359 | מה |  |
| 7:15.360 | שאמרתי |  |
| 7:15.361 | שאינם |  |
| 7:15.362 | מודים |  |
| 7:15.363 | בארבעה |  |
| 7:15.364 | עיקרים |  |
| 7:15.365 | הראשונים |  |
| 7:15.366 | שהרי |  |
| 7:15.367 | ביאר |  |
| 7:15.368 | הרב |  |
| 7:15.369 | חמדת |  |
| 7:15.370 | ימים |  |
| 7:15.371 | שלא |  |
| 7:15.372 | חרה |  |
| 7:15.373 | להאר"י |  |
| 7:15.374 | רק |  |
| 7:15.375 | על |  |
| 7:15.376 | ארבעה |  |
| 7:15.377 | בתים |  |
| 7:15.378 | ראשונים |  |
| 7:15.379 | שביגדל |  |
| 7:15.380 | שהם |  |
| 7:15.381 | סותרים |  |
| 7:15.382 | אמונתם | . |
| 7:15.383 | אבל |  |
| 7:15.384 | מהנו |  |
| 7:15.385 | אדון |  |
| 7:15.386 | עולם |  |
| 7:15.387 | ואילך |  |
| 7:15.388 | שפיר |  |
| 7:15.389 | דמי |  |
| 7:15.390 | לדידהו |  |
| 7:15.391 | לאמרו |  |
| 7:15.392 | כנז | ' |
| 7:15.393 | לעיל |  |
| 7:15.394 | סי | ' |
| 7:15.395 | מ"ח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

