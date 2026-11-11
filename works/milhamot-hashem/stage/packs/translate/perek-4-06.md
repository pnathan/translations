# Translation pack: milhamot-hashem perek-4-06 (Milchamot ha-Shem 4:10–4:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-06`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
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
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
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


## Milchamot ha-Shem 4:10

Hebrew (Wikisource): עניך הרואות איך רז"ל הרחיקו ליחס לשי"ת שהוא צריך לברואיו ונעזר בהם, או ליחס למלאכים העושים שליחותו שום בריאה, אלא הכל שב אל השי"ת, לא כמו שעשה הפילסוף מחבר הזהר שקרא לאבא בורא, ולאימא שם אותה לאומן העושה רצון בעל הבית. כמו שנתבאר לעיל, ולא יוחס הבריאה לסבה הראשונה רק לאבא, או או לחד פירוש לאריך בכח אבא כאלו הבורא חלש והם מסתייעים זה בזה בבריאה: ולפי דבריך שהמקובלים קיבלו כן ממשה רבי' ע"ה היה להם לע"ב זקנים להעתיק לתלמי המלך את הכתובים כמה שהם ולפרש לו כי בראשית שהוא אמר אבא ברא את אלהים שהיא בינה, וכן נעשה אדם היה להם לבאר שעל שנים נאמר, ולמה קלקלו את הכתובים בחנם, והחשיבו עצה במלין בלי דעת תורה שהיא הלכה למשה מסיני? ראוי היה להם להניח את הכתובים כמו שהם, ולפרש לו משמעות הכתובים בצורתן שהיא הלכה למשה לפי דבריך, והיו דבריהם נכונים ומתקבלים אצלו כפי דעתו באמונת אלוהות רבים? ולפרש לו ג"כ כי בראשית שם הוא כמו שפי' בזהר שהוא אבא והוא אלהים שאמר יהי אור ויהי רקיע, ויקוו המים, וגם יפרשו לו את השמים ואת הארץ. שמות אלוהות הן, כמו שפי' הפילוסוף בזהר שהם זעיר ונוקביה. ורז"ל הרחיקו הדעות האלה תכלית ההרחקה כדגרסינן בחגיגה פרק אין דורשין שאל ר' ישמעאל את ר' עקיבה כשהיו מהלכין בדרך. אמר לו אתה ששמשת את נחום איש גם זו כ"ב שנה שהיה דורש כל אתין שבתורה, את השמים ואת הארץ מאי דרש בהן אמר לו אלו נאמר שמים וארץ הייתי אומר שמותיה של הקב"ה הן. עכשיו שנ' את השמים ואת הארץ. שמים שמים ממש ארץ ארץ ממש. ובבראשית רבה הענין מתוקן יותר וז"ל ר' ישמעאל שאל את ר' עקיבה אמר לו בשביל ששמשת את נחום איש גם זו כ"ב שנה. אכין ורקין מעוטין. אתין וגמין ריבויין. הדין את דכתיב הכא מהו, אמר לו אלו נאמר בראשית ברא אלהים שמים וארץ היינו אומרים שמים וארץ אלוהות הן ע"כ. ועיין מהרש"א ועץ יוסף ותבין שזהו ממש דעת הזהר כי שמים וארץ הנז' כאן הם זעיר ונוקביה שהוא לפי דעתם ה' אלהינו וכמו שיתבאר לקמן.

| id | word | punctuation after |
|---|---|---|
| 4:10.1 | עניך |  |
| 4:10.2 | הרואות |  |
| 4:10.3 | איך |  |
| 4:10.4 | רז"ל |  |
| 4:10.5 | הרחיקו |  |
| 4:10.6 | ליחס |  |
| 4:10.7 | לשי"ת |  |
| 4:10.8 | שהוא |  |
| 4:10.9 | צריך |  |
| 4:10.10 | לברואיו |  |
| 4:10.11 | ונעזר |  |
| 4:10.12 | בהם | , |
| 4:10.13 | או |  |
| 4:10.14 | ליחס |  |
| 4:10.15 | למלאכים |  |
| 4:10.16 | העושים |  |
| 4:10.17 | שליחותו |  |
| 4:10.18 | שום |  |
| 4:10.19 | בריאה | , |
| 4:10.20 | אלא |  |
| 4:10.21 | הכל |  |
| 4:10.22 | שב |  |
| 4:10.23 | אל |  |
| 4:10.24 | השי"ת | , |
| 4:10.25 | לא |  |
| 4:10.26 | כמו |  |
| 4:10.27 | שעשה |  |
| 4:10.28 | הפילסוף |  |
| 4:10.29 | מחבר |  |
| 4:10.30 | הזהר |  |
| 4:10.31 | שקרא |  |
| 4:10.32 | לאבא |  |
| 4:10.33 | בורא | , |
| 4:10.34 | ולאימא |  |
| 4:10.35 | שם |  |
| 4:10.36 | אותה |  |
| 4:10.37 | לאומן |  |
| 4:10.38 | העושה |  |
| 4:10.39 | רצון |  |
| 4:10.40 | בעל |  |
| 4:10.41 | הבית | . |
| 4:10.42 | כמו |  |
| 4:10.43 | שנתבאר |  |
| 4:10.44 | לעיל | , |
| 4:10.45 | ולא |  |
| 4:10.46 | יוחס |  |
| 4:10.47 | הבריאה |  |
| 4:10.48 | לסבה |  |
| 4:10.49 | הראשונה |  |
| 4:10.50 | רק |  |
| 4:10.51 | לאבא | , |
| 4:10.52 | או |  |
| 4:10.53 | או |  |
| 4:10.54 | לחד |  |
| 4:10.55 | פירוש |  |
| 4:10.56 | לאריך |  |
| 4:10.57 | בכח |  |
| 4:10.58 | אבא |  |
| 4:10.59 | כאלו |  |
| 4:10.60 | הבורא |  |
| 4:10.61 | חלש |  |
| 4:10.62 | והם |  |
| 4:10.63 | מסתייעים |  |
| 4:10.64 | זה |  |
| 4:10.65 | בזה |  |
| 4:10.66 | בבריאה | : |
| 4:10.67 | ולפי |  |
| 4:10.68 | דבריך |  |
| 4:10.69 | שהמקובלים |  |
| 4:10.70 | קיבלו |  |
| 4:10.71 | כן |  |
| 4:10.72 | ממשה |  |
| 4:10.73 | רבי | ' |
| 4:10.74 | ע"ה |  |
| 4:10.75 | היה |  |
| 4:10.76 | להם |  |
| 4:10.77 | לע"ב |  |
| 4:10.78 | זקנים |  |
| 4:10.79 | להעתיק |  |
| 4:10.80 | לתלמי |  |
| 4:10.81 | המלך |  |
| 4:10.82 | את |  |
| 4:10.83 | הכתובים |  |
| 4:10.84 | כמה |  |
| 4:10.85 | שהם |  |
| 4:10.86 | ולפרש |  |
| 4:10.87 | לו |  |
| 4:10.88 | כי |  |
| 4:10.89 | בראשית |  |
| 4:10.90 | שהוא |  |
| 4:10.91 | אמר |  |
| 4:10.92 | אבא |  |
| 4:10.93 | ברא |  |
| 4:10.94 | את |  |
| 4:10.95 | אלהים |  |
| 4:10.96 | שהיא |  |
| 4:10.97 | בינה | , |
| 4:10.98 | וכן |  |
| 4:10.99 | נעשה |  |
| 4:10.100 | אדם |  |
| 4:10.101 | היה |  |
| 4:10.102 | להם |  |
| 4:10.103 | לבאר |  |
| 4:10.104 | שעל |  |
| 4:10.105 | שנים |  |
| 4:10.106 | נאמר | , |
| 4:10.107 | ולמה |  |
| 4:10.108 | קלקלו |  |
| 4:10.109 | את |  |
| 4:10.110 | הכתובים |  |
| 4:10.111 | בחנם | , |
| 4:10.112 | והחשיבו |  |
| 4:10.113 | עצה |  |
| 4:10.114 | במלין |  |
| 4:10.115 | בלי |  |
| 4:10.116 | דעת |  |
| 4:10.117 | תורה |  |
| 4:10.118 | שהיא |  |
| 4:10.119 | הלכה |  |
| 4:10.120 | למשה |  |
| 4:10.121 | מסיני | ? |
| 4:10.122 | ראוי |  |
| 4:10.123 | היה |  |
| 4:10.124 | להם |  |
| 4:10.125 | להניח |  |
| 4:10.126 | את |  |
| 4:10.127 | הכתובים |  |
| 4:10.128 | כמו |  |
| 4:10.129 | שהם | , |
| 4:10.130 | ולפרש |  |
| 4:10.131 | לו |  |
| 4:10.132 | משמעות |  |
| 4:10.133 | הכתובים |  |
| 4:10.134 | בצורתן |  |
| 4:10.135 | שהיא |  |
| 4:10.136 | הלכה |  |
| 4:10.137 | למשה |  |
| 4:10.138 | לפי |  |
| 4:10.139 | דבריך | , |
| 4:10.140 | והיו |  |
| 4:10.141 | דבריהם |  |
| 4:10.142 | נכונים |  |
| 4:10.143 | ומתקבלים |  |
| 4:10.144 | אצלו |  |
| 4:10.145 | כפי |  |
| 4:10.146 | דעתו |  |
| 4:10.147 | באמונת |  |
| 4:10.148 | אלוהות |  |
| 4:10.149 | רבים | ? |
| 4:10.150 | ולפרש |  |
| 4:10.151 | לו |  |
| 4:10.152 | ג"כ |  |
| 4:10.153 | כי |  |
| 4:10.154 | בראשית |  |
| 4:10.155 | שם |  |
| 4:10.156 | הוא |  |
| 4:10.157 | כמו |  |
| 4:10.158 | שפי | ' |
| 4:10.159 | בזהר |  |
| 4:10.160 | שהוא |  |
| 4:10.161 | אבא |  |
| 4:10.162 | והוא |  |
| 4:10.163 | אלהים |  |
| 4:10.164 | שאמר |  |
| 4:10.165 | יהי |  |
| 4:10.166 | אור |  |
| 4:10.167 | ויהי |  |
| 4:10.168 | רקיע | , |
| 4:10.169 | ויקוו |  |
| 4:10.170 | המים | , |
| 4:10.171 | וגם |  |
| 4:10.172 | יפרשו |  |
| 4:10.173 | לו |  |
| 4:10.174 | את |  |
| 4:10.175 | השמים |  |
| 4:10.176 | ואת |  |
| 4:10.177 | הארץ | . |
| 4:10.178 | שמות |  |
| 4:10.179 | אלוהות |  |
| 4:10.180 | הן | , |
| 4:10.181 | כמו |  |
| 4:10.182 | שפי | ' |
| 4:10.183 | הפילוסוף |  |
| 4:10.184 | בזהר |  |
| 4:10.185 | שהם |  |
| 4:10.186 | זעיר |  |
| 4:10.187 | ונוקביה | . |
| 4:10.188 | ורז"ל |  |
| 4:10.189 | הרחיקו |  |
| 4:10.190 | הדעות |  |
| 4:10.191 | האלה |  |
| 4:10.192 | תכלית |  |
| 4:10.193 | ההרחקה |  |
| 4:10.194 | כדגרסינן |  |
| 4:10.195 | בחגיגה |  |
| 4:10.196 | פרק |  |
| 4:10.197 | אין |  |
| 4:10.198 | דורשין |  |
| 4:10.199 | שאל |  |
| 4:10.200 | ר | ' |
| 4:10.201 | ישמעאל |  |
| 4:10.202 | את |  |
| 4:10.203 | ר | ' |
| 4:10.204 | עקיבה |  |
| 4:10.205 | כשהיו |  |
| 4:10.206 | מהלכין |  |
| 4:10.207 | בדרך | . |
| 4:10.208 | אמר |  |
| 4:10.209 | לו |  |
| 4:10.210 | אתה |  |
| 4:10.211 | ששמשת |  |
| 4:10.212 | את |  |
| 4:10.213 | נחום |  |
| 4:10.214 | איש |  |
| 4:10.215 | גם |  |
| 4:10.216 | זו |  |
| 4:10.217 | כ"ב |  |
| 4:10.218 | שנה |  |
| 4:10.219 | שהיה |  |
| 4:10.220 | דורש |  |
| 4:10.221 | כל |  |
| 4:10.222 | אתין |  |
| 4:10.223 | שבתורה | , |
| 4:10.224 | את |  |
| 4:10.225 | השמים |  |
| 4:10.226 | ואת |  |
| 4:10.227 | הארץ |  |
| 4:10.228 | מאי |  |
| 4:10.229 | דרש |  |
| 4:10.230 | בהן |  |
| 4:10.231 | אמר |  |
| 4:10.232 | לו |  |
| 4:10.233 | אלו |  |
| 4:10.234 | נאמר |  |
| 4:10.235 | שמים |  |
| 4:10.236 | וארץ |  |
| 4:10.237 | הייתי |  |
| 4:10.238 | אומר |  |
| 4:10.239 | שמותיה |  |
| 4:10.240 | של |  |
| 4:10.241 | הקב"ה |  |
| 4:10.242 | הן | . |
| 4:10.243 | עכשיו |  |
| 4:10.244 | שנ | ' |
| 4:10.245 | את |  |
| 4:10.246 | השמים |  |
| 4:10.247 | ואת |  |
| 4:10.248 | הארץ | . |
| 4:10.249 | שמים |  |
| 4:10.250 | שמים |  |
| 4:10.251 | ממש |  |
| 4:10.252 | ארץ |  |
| 4:10.253 | ארץ |  |
| 4:10.254 | ממש | . |
| 4:10.255 | ובבראשית |  |
| 4:10.256 | רבה |  |
| 4:10.257 | הענין |  |
| 4:10.258 | מתוקן |  |
| 4:10.259 | יותר |  |
| 4:10.260 | וז"ל |  |
| 4:10.261 | ר | ' |
| 4:10.262 | ישמעאל |  |
| 4:10.263 | שאל |  |
| 4:10.264 | את |  |
| 4:10.265 | ר | ' |
| 4:10.266 | עקיבה |  |
| 4:10.267 | אמר |  |
| 4:10.268 | לו |  |
| 4:10.269 | בשביל |  |
| 4:10.270 | ששמשת |  |
| 4:10.271 | את |  |
| 4:10.272 | נחום |  |
| 4:10.273 | איש |  |
| 4:10.274 | גם |  |
| 4:10.275 | זו |  |
| 4:10.276 | כ"ב |  |
| 4:10.277 | שנה | . |
| 4:10.278 | אכין |  |
| 4:10.279 | ורקין |  |
| 4:10.280 | מעוטין | . |
| 4:10.281 | אתין |  |
| 4:10.282 | וגמין |  |
| 4:10.283 | ריבויין | . |
| 4:10.284 | הדין |  |
| 4:10.285 | את |  |
| 4:10.286 | דכתיב |  |
| 4:10.287 | הכא |  |
| 4:10.288 | מהו | , |
| 4:10.289 | אמר |  |
| 4:10.290 | לו |  |
| 4:10.291 | אלו |  |
| 4:10.292 | נאמר |  |
| 4:10.293 | בראשית |  |
| 4:10.294 | ברא |  |
| 4:10.295 | אלהים |  |
| 4:10.296 | שמים |  |
| 4:10.297 | וארץ |  |
| 4:10.298 | היינו |  |
| 4:10.299 | אומרים |  |
| 4:10.300 | שמים |  |
| 4:10.301 | וארץ |  |
| 4:10.302 | אלוהות |  |
| 4:10.303 | הן |  |
| 4:10.304 | ע"כ | . |
| 4:10.305 | ועיין |  |
| 4:10.306 | מהרש"א |  |
| 4:10.307 | ועץ |  |
| 4:10.308 | יוסף |  |
| 4:10.309 | ותבין |  |
| 4:10.310 | שזהו |  |
| 4:10.311 | ממש |  |
| 4:10.312 | דעת |  |
| 4:10.313 | הזהר |  |
| 4:10.314 | כי |  |
| 4:10.315 | שמים |  |
| 4:10.316 | וארץ |  |
| 4:10.317 | הנז | ' |
| 4:10.318 | כאן |  |
| 4:10.319 | הם |  |
| 4:10.320 | זעיר |  |
| 4:10.321 | ונוקביה |  |
| 4:10.322 | שהוא |  |
| 4:10.323 | לפי |  |
| 4:10.324 | דעתם |  |
| 4:10.325 | ה | ' |
| 4:10.326 | אלהינו |  |
| 4:10.327 | וכמו |  |
| 4:10.328 | שיתבאר |  |
| 4:10.329 | לקמן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

