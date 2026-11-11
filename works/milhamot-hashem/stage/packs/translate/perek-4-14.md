# Translation pack: milhamot-hashem perek-4-14 (Milchamot ha-Shem 4:23–4:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-14`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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


## Milchamot ha-Shem 4:23

Hebrew (Wikisource): לב) והנה ידוע ומפורסם כמה הקפידו רז"ל מעתיקי השמועה ממשה רבי' ע"ה על האומר מודים מודים שתי פעמים או שמע שמע שתי פעמים שמשתקים אותנו כדתנן בברכות פ' אין עומדין. ובמגלה פרק הקורא האומר מודים מודים משתקים אותו. ובברכות דף ל"ג (ברכות לג, א) גר' אמר רבי זירא כל האומר שמע שמע כאומר מודים מודים דמי.ופי' כל המפרשים דטעמא משום דמחזי כשתי רשויות וכמקבל עליו שני אלוהות. ובברכות דף י"ד (דף יד.) ההוא דנחית קמיה דרבא שמעיה דקאמר אמת אמת תרי זמני. אמר רבא כל אמת אמת תפסיה להאי ע"כ וכ' הכונה בעין יע' כשם שאין לומר אחד אחד שהוא רבוי ולא אחדות. כך אין לומר אמת אמת כי האמת אינו אלא אחד. והוא מאמר רבא רהטא של אמת תפסתו לזה. כלומר הוא סובר לחזק דבריו באמתות השי"ת בהיותו מוסיף ואינו אלא גורע ע"כ. ובסוכה פרק החליל אהא דתנן אנו ליה וליה עינינו גרסינן איני והאמר ר' זירא כל האומר שמע שמע כאלו אומר מודים מודים? (פי' רש"י כאומר מודים מודים דתנן משתקים אותו משום דמחזי כשתי רשויות) ומתרץ אלא הכי אמרי אנו ליה מודים. ועינינו ליה מיחלות ע"כ. אלמא דרבנן קפדי אבל מלתא דמשתמע מינה שיש שתי רשויות. וכן פסקו כל הפוסקים דהאומר שמע שמע או מודים מודים משתקים אותו משום דמחזי כשתי רשויות. ומזה תדון קל וחומר (כפי הכלל המסור בידינו אדם דן ק"ו מעצמו) לנדון הוזכר בזהר (בראשית ל"ד ע"א) וז"ל כתיב ה' מלך ה' מלך ה' ימלוך לעולם ועד. ה' מלך לעילא. ה' מלך באמצעיתא. ה' ימלוך לתתא ופירש המקדש מלך בשם הכוונות ה' מלך הוא אריך. ה' מלך הוא אבא ואימא. ה' ימלוך זעיר ונוקביה ע"כ. אלמא דבאמרו ה' מלך הוא ממליך לאריך אנפין. ובאמרו ה' מלך הוא ממליך לאבא ואימא. ובאמרו ה' ימלוך לעולם ועד ממליך לזעיר ונוקביה. היש אמונת רשויות רבות ומינות גדולה מזו? שכיון שהמליך שלשה מלכים באמרו ה' מלך וכו' איך ישקר לומר אבינו מלכנו אין לנו מלך אלא אתה. והלא כבר המליך שלשה מלכים? ובאיזה מהם הוא בוחר עתה להמליכו. ולהסיר עטרת מלוכה מהאחרים? ובזה וכיוצא בו צדקו דברי הר' בעל העיקרים מאמר שני סוף פ' כ"ח ששם הזהיר מאד שלא ללמוד בספר הזהר ושאר ספרי קבלה וז"ל וכלל אומר לך השמר לך ושמור נפשך מאד פן תנקש אחריהם ותלכד במצודתם. כי הם עוזבים ארחות יושר ללכת בדרכי חשך כי לא ידעו ולא יבינו בחשכה חתהלכו המתעסקים בקבלה מדעת עצמם מזולת קבלה מפי חכם מקובל עכ"ל. ועל כן כל אוהב את השי"ת ודבק בתורתו שבכתב ושבעל פה אשר העתיקו רז"ל בעלי המשנה והתלמוד. ירחיק עצמו מהקבלה החדשה, ואל ידאג למטעמיה פן יפסיד אמונתו הטהורה המזוקקה שבעתים על פי תנאים, ואמוראים ופוסקים.

| id | word | punctuation after |
|---|---|---|
| 4:23.1 | לב | ) |
| 4:23.2 | והנה |  |
| 4:23.3 | ידוע |  |
| 4:23.4 | ומפורסם |  |
| 4:23.5 | כמה |  |
| 4:23.6 | הקפידו |  |
| 4:23.7 | רז"ל |  |
| 4:23.8 | מעתיקי |  |
| 4:23.9 | השמועה |  |
| 4:23.10 | ממשה |  |
| 4:23.11 | רבי | ' |
| 4:23.12 | ע"ה |  |
| 4:23.13 | על |  |
| 4:23.14 | האומר |  |
| 4:23.15 | מודים |  |
| 4:23.16 | מודים |  |
| 4:23.17 | שתי |  |
| 4:23.18 | פעמים |  |
| 4:23.19 | או |  |
| 4:23.20 | שמע |  |
| 4:23.21 | שמע |  |
| 4:23.22 | שתי |  |
| 4:23.23 | פעמים |  |
| 4:23.24 | שמשתקים |  |
| 4:23.25 | אותנו |  |
| 4:23.26 | כדתנן |  |
| 4:23.27 | בברכות |  |
| 4:23.28 | פ | ' |
| 4:23.29 | אין |  |
| 4:23.30 | עומדין | . |
| 4:23.31 | ובמגלה |  |
| 4:23.32 | פרק |  |
| 4:23.33 | הקורא |  |
| 4:23.34 | האומר |  |
| 4:23.35 | מודים |  |
| 4:23.36 | מודים |  |
| 4:23.37 | משתקים |  |
| 4:23.38 | אותו | . |
| 4:23.39 | ובברכות |  |
| 4:23.40 | דף |  |
| 4:23.41 | ל"ג |  |
| 4:23.42 | ברכות |  |
| 4:23.43 | לג | , |
| 4:23.44 | א | ) |
| 4:23.45 | גר | ' |
| 4:23.46 | אמר |  |
| 4:23.47 | רבי |  |
| 4:23.48 | זירא |  |
| 4:23.49 | כל |  |
| 4:23.50 | האומר |  |
| 4:23.51 | שמע |  |
| 4:23.52 | שמע |  |
| 4:23.53 | כאומר |  |
| 4:23.54 | מודים |  |
| 4:23.55 | מודים |  |
| 4:23.56 | דמי.ופי | ' |
| 4:23.57 | כל |  |
| 4:23.58 | המפרשים |  |
| 4:23.59 | דטעמא |  |
| 4:23.60 | משום |  |
| 4:23.61 | דמחזי |  |
| 4:23.62 | כשתי |  |
| 4:23.63 | רשויות |  |
| 4:23.64 | וכמקבל |  |
| 4:23.65 | עליו |  |
| 4:23.66 | שני |  |
| 4:23.67 | אלוהות | . |
| 4:23.68 | ובברכות |  |
| 4:23.69 | דף |  |
| 4:23.70 | י"ד |  |
| 4:23.71 | דף |  |
| 4:23.72 | יד | .) |
| 4:23.73 | ההוא |  |
| 4:23.74 | דנחית |  |
| 4:23.75 | קמיה |  |
| 4:23.76 | דרבא |  |
| 4:23.77 | שמעיה |  |
| 4:23.78 | דקאמר |  |
| 4:23.79 | אמת |  |
| 4:23.80 | אמת |  |
| 4:23.81 | תרי |  |
| 4:23.82 | זמני | . |
| 4:23.83 | אמר |  |
| 4:23.84 | רבא |  |
| 4:23.85 | כל |  |
| 4:23.86 | אמת |  |
| 4:23.87 | אמת |  |
| 4:23.88 | תפסיה |  |
| 4:23.89 | להאי |  |
| 4:23.90 | ע"כ |  |
| 4:23.91 | וכ | ' |
| 4:23.92 | הכונה |  |
| 4:23.93 | בעין |  |
| 4:23.94 | יע | ' |
| 4:23.95 | כשם |  |
| 4:23.96 | שאין |  |
| 4:23.97 | לומר |  |
| 4:23.98 | אחד |  |
| 4:23.99 | אחד |  |
| 4:23.100 | שהוא |  |
| 4:23.101 | רבוי |  |
| 4:23.102 | ולא |  |
| 4:23.103 | אחדות | . |
| 4:23.104 | כך |  |
| 4:23.105 | אין |  |
| 4:23.106 | לומר |  |
| 4:23.107 | אמת |  |
| 4:23.108 | אמת |  |
| 4:23.109 | כי |  |
| 4:23.110 | האמת |  |
| 4:23.111 | אינו |  |
| 4:23.112 | אלא |  |
| 4:23.113 | אחד | . |
| 4:23.114 | והוא |  |
| 4:23.115 | מאמר |  |
| 4:23.116 | רבא |  |
| 4:23.117 | רהטא |  |
| 4:23.118 | של |  |
| 4:23.119 | אמת |  |
| 4:23.120 | תפסתו |  |
| 4:23.121 | לזה | . |
| 4:23.122 | כלומר |  |
| 4:23.123 | הוא |  |
| 4:23.124 | סובר |  |
| 4:23.125 | לחזק |  |
| 4:23.126 | דבריו |  |
| 4:23.127 | באמתות |  |
| 4:23.128 | השי"ת |  |
| 4:23.129 | בהיותו |  |
| 4:23.130 | מוסיף |  |
| 4:23.131 | ואינו |  |
| 4:23.132 | אלא |  |
| 4:23.133 | גורע |  |
| 4:23.134 | ע"כ | . |
| 4:23.135 | ובסוכה |  |
| 4:23.136 | פרק |  |
| 4:23.137 | החליל |  |
| 4:23.138 | אהא |  |
| 4:23.139 | דתנן |  |
| 4:23.140 | אנו |  |
| 4:23.141 | ליה |  |
| 4:23.142 | וליה |  |
| 4:23.143 | עינינו |  |
| 4:23.144 | גרסינן |  |
| 4:23.145 | איני |  |
| 4:23.146 | והאמר |  |
| 4:23.147 | ר | ' |
| 4:23.148 | זירא |  |
| 4:23.149 | כל |  |
| 4:23.150 | האומר |  |
| 4:23.151 | שמע |  |
| 4:23.152 | שמע |  |
| 4:23.153 | כאלו |  |
| 4:23.154 | אומר |  |
| 4:23.155 | מודים |  |
| 4:23.156 | מודים | ? |
| 4:23.157 | פי | ' |
| 4:23.158 | רש"י |  |
| 4:23.159 | כאומר |  |
| 4:23.160 | מודים |  |
| 4:23.161 | מודים |  |
| 4:23.162 | דתנן |  |
| 4:23.163 | משתקים |  |
| 4:23.164 | אותו |  |
| 4:23.165 | משום |  |
| 4:23.166 | דמחזי |  |
| 4:23.167 | כשתי |  |
| 4:23.168 | רשויות | ) |
| 4:23.169 | ומתרץ |  |
| 4:23.170 | אלא |  |
| 4:23.171 | הכי |  |
| 4:23.172 | אמרי |  |
| 4:23.173 | אנו |  |
| 4:23.174 | ליה |  |
| 4:23.175 | מודים | . |
| 4:23.176 | ועינינו |  |
| 4:23.177 | ליה |  |
| 4:23.178 | מיחלות |  |
| 4:23.179 | ע"כ | . |
| 4:23.180 | אלמא |  |
| 4:23.181 | דרבנן |  |
| 4:23.182 | קפדי |  |
| 4:23.183 | אבל |  |
| 4:23.184 | מלתא |  |
| 4:23.185 | דמשתמע |  |
| 4:23.186 | מינה |  |
| 4:23.187 | שיש |  |
| 4:23.188 | שתי |  |
| 4:23.189 | רשויות | . |
| 4:23.190 | וכן |  |
| 4:23.191 | פסקו |  |
| 4:23.192 | כל |  |
| 4:23.193 | הפוסקים |  |
| 4:23.194 | דהאומר |  |
| 4:23.195 | שמע |  |
| 4:23.196 | שמע |  |
| 4:23.197 | או |  |
| 4:23.198 | מודים |  |
| 4:23.199 | מודים |  |
| 4:23.200 | משתקים |  |
| 4:23.201 | אותו |  |
| 4:23.202 | משום |  |
| 4:23.203 | דמחזי |  |
| 4:23.204 | כשתי |  |
| 4:23.205 | רשויות | . |
| 4:23.206 | ומזה |  |
| 4:23.207 | תדון |  |
| 4:23.208 | קל |  |
| 4:23.209 | וחומר |  |
| 4:23.210 | כפי |  |
| 4:23.211 | הכלל |  |
| 4:23.212 | המסור |  |
| 4:23.213 | בידינו |  |
| 4:23.214 | אדם |  |
| 4:23.215 | דן |  |
| 4:23.216 | ק"ו |  |
| 4:23.217 | מעצמו | ) |
| 4:23.218 | לנדון |  |
| 4:23.219 | הוזכר |  |
| 4:23.220 | בזהר |  |
| 4:23.221 | בראשית |  |
| 4:23.222 | ל"ד |  |
| 4:23.223 | ע"א | ) |
| 4:23.224 | וז"ל |  |
| 4:23.225 | כתיב |  |
| 4:23.226 | ה | ' |
| 4:23.227 | מלך |  |
| 4:23.228 | ה | ' |
| 4:23.229 | מלך |  |
| 4:23.230 | ה | ' |
| 4:23.231 | ימלוך |  |
| 4:23.232 | לעולם |  |
| 4:23.233 | ועד | . |
| 4:23.234 | ה | ' |
| 4:23.235 | מלך |  |
| 4:23.236 | לעילא | . |
| 4:23.237 | ה | ' |
| 4:23.238 | מלך |  |
| 4:23.239 | באמצעיתא | . |
| 4:23.240 | ה | ' |
| 4:23.241 | ימלוך |  |
| 4:23.242 | לתתא |  |
| 4:23.243 | ופירש |  |
| 4:23.244 | המקדש |  |
| 4:23.245 | מלך |  |
| 4:23.246 | בשם |  |
| 4:23.247 | הכוונות |  |
| 4:23.248 | ה | ' |
| 4:23.249 | מלך |  |
| 4:23.250 | הוא |  |
| 4:23.251 | אריך | . |
| 4:23.252 | ה | ' |
| 4:23.253 | מלך |  |
| 4:23.254 | הוא |  |
| 4:23.255 | אבא |  |
| 4:23.256 | ואימא | . |
| 4:23.257 | ה | ' |
| 4:23.258 | ימלוך |  |
| 4:23.259 | זעיר |  |
| 4:23.260 | ונוקביה |  |
| 4:23.261 | ע"כ | . |
| 4:23.262 | אלמא |  |
| 4:23.263 | דבאמרו |  |
| 4:23.264 | ה | ' |
| 4:23.265 | מלך |  |
| 4:23.266 | הוא |  |
| 4:23.267 | ממליך |  |
| 4:23.268 | לאריך |  |
| 4:23.269 | אנפין | . |
| 4:23.270 | ובאמרו |  |
| 4:23.271 | ה | ' |
| 4:23.272 | מלך |  |
| 4:23.273 | הוא |  |
| 4:23.274 | ממליך |  |
| 4:23.275 | לאבא |  |
| 4:23.276 | ואימא | . |
| 4:23.277 | ובאמרו |  |
| 4:23.278 | ה | ' |
| 4:23.279 | ימלוך |  |
| 4:23.280 | לעולם |  |
| 4:23.281 | ועד |  |
| 4:23.282 | ממליך |  |
| 4:23.283 | לזעיר |  |
| 4:23.284 | ונוקביה | . |
| 4:23.285 | היש |  |
| 4:23.286 | אמונת |  |
| 4:23.287 | רשויות |  |
| 4:23.288 | רבות |  |
| 4:23.289 | ומינות |  |
| 4:23.290 | גדולה |  |
| 4:23.291 | מזו | ? |
| 4:23.292 | שכיון |  |
| 4:23.293 | שהמליך |  |
| 4:23.294 | שלשה |  |
| 4:23.295 | מלכים |  |
| 4:23.296 | באמרו |  |
| 4:23.297 | ה | ' |
| 4:23.298 | מלך |  |
| 4:23.299 | וכו | ' |
| 4:23.300 | איך |  |
| 4:23.301 | ישקר |  |
| 4:23.302 | לומר |  |
| 4:23.303 | אבינו |  |
| 4:23.304 | מלכנו |  |
| 4:23.305 | אין |  |
| 4:23.306 | לנו |  |
| 4:23.307 | מלך |  |
| 4:23.308 | אלא |  |
| 4:23.309 | אתה | . |
| 4:23.310 | והלא |  |
| 4:23.311 | כבר |  |
| 4:23.312 | המליך |  |
| 4:23.313 | שלשה |  |
| 4:23.314 | מלכים | ? |
| 4:23.315 | ובאיזה |  |
| 4:23.316 | מהם |  |
| 4:23.317 | הוא |  |
| 4:23.318 | בוחר |  |
| 4:23.319 | עתה |  |
| 4:23.320 | להמליכו | . |
| 4:23.321 | ולהסיר |  |
| 4:23.322 | עטרת |  |
| 4:23.323 | מלוכה |  |
| 4:23.324 | מהאחרים | ? |
| 4:23.325 | ובזה |  |
| 4:23.326 | וכיוצא |  |
| 4:23.327 | בו |  |
| 4:23.328 | צדקו |  |
| 4:23.329 | דברי |  |
| 4:23.330 | הר | ' |
| 4:23.331 | בעל |  |
| 4:23.332 | העיקרים |  |
| 4:23.333 | מאמר |  |
| 4:23.334 | שני |  |
| 4:23.335 | סוף |  |
| 4:23.336 | פ | ' |
| 4:23.337 | כ"ח |  |
| 4:23.338 | ששם |  |
| 4:23.339 | הזהיר |  |
| 4:23.340 | מאד |  |
| 4:23.341 | שלא |  |
| 4:23.342 | ללמוד |  |
| 4:23.343 | בספר |  |
| 4:23.344 | הזהר |  |
| 4:23.345 | ושאר |  |
| 4:23.346 | ספרי |  |
| 4:23.347 | קבלה |  |
| 4:23.348 | וז"ל |  |
| 4:23.349 | וכלל |  |
| 4:23.350 | אומר |  |
| 4:23.351 | לך |  |
| 4:23.352 | השמר |  |
| 4:23.353 | לך |  |
| 4:23.354 | ושמור |  |
| 4:23.355 | נפשך |  |
| 4:23.356 | מאד |  |
| 4:23.357 | פן |  |
| 4:23.358 | תנקש |  |
| 4:23.359 | אחריהם |  |
| 4:23.360 | ותלכד |  |
| 4:23.361 | במצודתם | . |
| 4:23.362 | כי |  |
| 4:23.363 | הם |  |
| 4:23.364 | עוזבים |  |
| 4:23.365 | ארחות |  |
| 4:23.366 | יושר |  |
| 4:23.367 | ללכת |  |
| 4:23.368 | בדרכי |  |
| 4:23.369 | חשך |  |
| 4:23.370 | כי |  |
| 4:23.371 | לא |  |
| 4:23.372 | ידעו |  |
| 4:23.373 | ולא |  |
| 4:23.374 | יבינו |  |
| 4:23.375 | בחשכה |  |
| 4:23.376 | חתהלכו |  |
| 4:23.377 | המתעסקים |  |
| 4:23.378 | בקבלה |  |
| 4:23.379 | מדעת |  |
| 4:23.380 | עצמם |  |
| 4:23.381 | מזולת |  |
| 4:23.382 | קבלה |  |
| 4:23.383 | מפי |  |
| 4:23.384 | חכם |  |
| 4:23.385 | מקובל |  |
| 4:23.386 | עכ"ל | . |
| 4:23.387 | ועל |  |
| 4:23.388 | כן |  |
| 4:23.389 | כל |  |
| 4:23.390 | אוהב |  |
| 4:23.391 | את |  |
| 4:23.392 | השי"ת |  |
| 4:23.393 | ודבק |  |
| 4:23.394 | בתורתו |  |
| 4:23.395 | שבכתב |  |
| 4:23.396 | ושבעל |  |
| 4:23.397 | פה |  |
| 4:23.398 | אשר |  |
| 4:23.399 | העתיקו |  |
| 4:23.400 | רז"ל |  |
| 4:23.401 | בעלי |  |
| 4:23.402 | המשנה |  |
| 4:23.403 | והתלמוד | . |
| 4:23.404 | ירחיק |  |
| 4:23.405 | עצמו |  |
| 4:23.406 | מהקבלה |  |
| 4:23.407 | החדשה | , |
| 4:23.408 | ואל |  |
| 4:23.409 | ידאג |  |
| 4:23.410 | למטעמיה |  |
| 4:23.411 | פן |  |
| 4:23.412 | יפסיד |  |
| 4:23.413 | אמונתו |  |
| 4:23.414 | הטהורה |  |
| 4:23.415 | המזוקקה |  |
| 4:23.416 | שבעתים |  |
| 4:23.417 | על |  |
| 4:23.418 | פי |  |
| 4:23.419 | תנאים | , |
| 4:23.420 | ואמוראים |  |
| 4:23.421 | ופוסקים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

