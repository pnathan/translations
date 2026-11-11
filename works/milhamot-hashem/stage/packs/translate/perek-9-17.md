# Translation pack: milhamot-hashem perek-9-17 (Milchamot ha-Shem 9:20–9:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-17`.

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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
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


## Milchamot ha-Shem 9:20

Hebrew (Wikisource): קלח) את כל אלה עשתה יד הפילסוף המסית מחבר הזהר להלביש את קוראיו גאה וגאון אף שאינם יודעים מכל דיני תורתינו הקדושה מאומה, ואף מכל מה שלומדים בו שכן הוא הבטיח להם שכר רב בלמדם אותו אף בלי הבין, ולכן הם חושבים עצמם בני עלייה, ומבזים המה את העוסקים במשנה ובתלמוד שהם כתרה של תורה כעדות רז"ל במס' מגלה והבאתי לעיל סי' קל"ו, ובעיניהם המשנה והתלמוד חשובים לתבן וקש, והעוסקים בהם כבהמות יער האוכלות חציר ותבן, וכעופות כדאיתא בזהר פרשת תצא הובא לעיל סי' ס"ח, וכן חשב מהרח"ו בשבחיו דף י"ב, ולא ידעו ולא יבינו כי בחשכה יתהלכו, ולא ידעו להזהר באיסורים דאוריתא ודרבנן ונקל להם לדבר על ריעיהם לה"ר ורכילות, ולמסור גופם וממונם לאנסים ולהעיד עליהם עדות שקר אף בעש"ג כי הם מינים ואפיקורסים על שפורשים מדרכי הקבלה החדשה אשר רבו בה הבעלים והעשתרות והם האלוהות הזכרים והנקבות, אשר כל עבודת עובדיהם לתקנם ולקשט הנקבות כדי למצוא חן בעיני האלוהות הזכרים שיזדווגו להן כדי להוריד שפע רוחניותם גם לעולם התחתון, כדרשת רז"ל על והנה עמד על היאור לעיל סי' ע"ה, שהאלהים האחרים אלמלא עובדיהם שומרים אותם ומשגיחים עליהם בכל צרכיהם ומחזקים בדקיהם וקלקולם, אזי כרעו ונפלו וחלפו ועברו ובטלו, וכזה היא דעת המקובלים החדשים שכל מעשה מצותיהם עושים לשם האלילים לתקנם כמבואר בעץ חיים שער נקודים סוף פ"ו שקודם אדה"ר היו זעיר ונוקביה אחור באחור וכו' וכשנולד אדה"ר, על ידי מעשיו הטובים החזירם פנים בפנים (ודבריו אלה היפך זהר בראשית כ"ב) אבל אבא ואימא לא הוצרכו למעשיו כי על ידי עצמן חזרו פנים בפנים, (ובזהר שם אמר שנתגרשה אימא על ידי מעשיו). ובס' חיי שלום בשם ס' הלקוטים כתב כי עד אברהם אבינו עדיין לא נתקן הזעיר אנפין וכו' ע"ש, הכלל העולה כי האלוה קצר האפים, שעובדים אותו המקובלים החדשים הוא צריך תיקון בכל עת וזמן עם שתי נשיו רחל ולאה כאשר יראה המעיין בסדר תיקון חצות בספר משנת חסידים ובס' פרי עץ חיים לקשטם לו כדי שיתחמם הזעיר ויתאוה להם (משנת חסידים תיקון חצות פ"ד משנה ג'). אבל אלהינו בהו"ש אינו צריך לבריותיו כלל, רק הם צריכים לו, והוא שומר את יראיו וחסידיו מכל רע, כמבואר בכמה מקומות בכתובים ובדברי רז"ל, ואין צורך להאריך בהם כי די בעיקר שכתב הרמב"ם וכמ"ש רז"ל והנה ה' נצב עליו לשמרו וכו'.

| id | word | punctuation after |
|---|---|---|
| 9:20.1 | קלח | ) |
| 9:20.2 | את |  |
| 9:20.3 | כל |  |
| 9:20.4 | אלה |  |
| 9:20.5 | עשתה |  |
| 9:20.6 | יד |  |
| 9:20.7 | הפילסוף |  |
| 9:20.8 | המסית |  |
| 9:20.9 | מחבר |  |
| 9:20.10 | הזהר |  |
| 9:20.11 | להלביש |  |
| 9:20.12 | את |  |
| 9:20.13 | קוראיו |  |
| 9:20.14 | גאה |  |
| 9:20.15 | וגאון |  |
| 9:20.16 | אף |  |
| 9:20.17 | שאינם |  |
| 9:20.18 | יודעים |  |
| 9:20.19 | מכל |  |
| 9:20.20 | דיני |  |
| 9:20.21 | תורתינו |  |
| 9:20.22 | הקדושה |  |
| 9:20.23 | מאומה | , |
| 9:20.24 | ואף |  |
| 9:20.25 | מכל |  |
| 9:20.26 | מה |  |
| 9:20.27 | שלומדים |  |
| 9:20.28 | בו |  |
| 9:20.29 | שכן |  |
| 9:20.30 | הוא |  |
| 9:20.31 | הבטיח |  |
| 9:20.32 | להם |  |
| 9:20.33 | שכר |  |
| 9:20.34 | רב |  |
| 9:20.35 | בלמדם |  |
| 9:20.36 | אותו |  |
| 9:20.37 | אף |  |
| 9:20.38 | בלי |  |
| 9:20.39 | הבין | , |
| 9:20.40 | ולכן |  |
| 9:20.41 | הם |  |
| 9:20.42 | חושבים |  |
| 9:20.43 | עצמם |  |
| 9:20.44 | בני |  |
| 9:20.45 | עלייה | , |
| 9:20.46 | ומבזים |  |
| 9:20.47 | המה |  |
| 9:20.48 | את |  |
| 9:20.49 | העוסקים |  |
| 9:20.50 | במשנה |  |
| 9:20.51 | ובתלמוד |  |
| 9:20.52 | שהם |  |
| 9:20.53 | כתרה |  |
| 9:20.54 | של |  |
| 9:20.55 | תורה |  |
| 9:20.56 | כעדות |  |
| 9:20.57 | רז"ל |  |
| 9:20.58 | במס | ' |
| 9:20.59 | מגלה |  |
| 9:20.60 | והבאתי |  |
| 9:20.61 | לעיל |  |
| 9:20.62 | סי | ' |
| 9:20.63 | קל"ו | , |
| 9:20.64 | ובעיניהם |  |
| 9:20.65 | המשנה |  |
| 9:20.66 | והתלמוד |  |
| 9:20.67 | חשובים |  |
| 9:20.68 | לתבן |  |
| 9:20.69 | וקש | , |
| 9:20.70 | והעוסקים |  |
| 9:20.71 | בהם |  |
| 9:20.72 | כבהמות |  |
| 9:20.73 | יער |  |
| 9:20.74 | האוכלות |  |
| 9:20.75 | חציר |  |
| 9:20.76 | ותבן | , |
| 9:20.77 | וכעופות |  |
| 9:20.78 | כדאיתא |  |
| 9:20.79 | בזהר |  |
| 9:20.80 | פרשת |  |
| 9:20.81 | תצא |  |
| 9:20.82 | הובא |  |
| 9:20.83 | לעיל |  |
| 9:20.84 | סי | ' |
| 9:20.85 | ס"ח | , |
| 9:20.86 | וכן |  |
| 9:20.87 | חשב |  |
| 9:20.88 | מהרח"ו |  |
| 9:20.89 | בשבחיו |  |
| 9:20.90 | דף |  |
| 9:20.91 | י"ב | , |
| 9:20.92 | ולא |  |
| 9:20.93 | ידעו |  |
| 9:20.94 | ולא |  |
| 9:20.95 | יבינו |  |
| 9:20.96 | כי |  |
| 9:20.97 | בחשכה |  |
| 9:20.98 | יתהלכו | , |
| 9:20.99 | ולא |  |
| 9:20.100 | ידעו |  |
| 9:20.101 | להזהר |  |
| 9:20.102 | באיסורים |  |
| 9:20.103 | דאוריתא |  |
| 9:20.104 | ודרבנן |  |
| 9:20.105 | ונקל |  |
| 9:20.106 | להם |  |
| 9:20.107 | לדבר |  |
| 9:20.108 | על |  |
| 9:20.109 | ריעיהם |  |
| 9:20.110 | לה"ר |  |
| 9:20.111 | ורכילות | , |
| 9:20.112 | ולמסור |  |
| 9:20.113 | גופם |  |
| 9:20.114 | וממונם |  |
| 9:20.115 | לאנסים |  |
| 9:20.116 | ולהעיד |  |
| 9:20.117 | עליהם |  |
| 9:20.118 | עדות |  |
| 9:20.119 | שקר |  |
| 9:20.120 | אף |  |
| 9:20.121 | בעש"ג |  |
| 9:20.122 | כי |  |
| 9:20.123 | הם |  |
| 9:20.124 | מינים |  |
| 9:20.125 | ואפיקורסים |  |
| 9:20.126 | על |  |
| 9:20.127 | שפורשים |  |
| 9:20.128 | מדרכי |  |
| 9:20.129 | הקבלה |  |
| 9:20.130 | החדשה |  |
| 9:20.131 | אשר |  |
| 9:20.132 | רבו |  |
| 9:20.133 | בה |  |
| 9:20.134 | הבעלים |  |
| 9:20.135 | והעשתרות |  |
| 9:20.136 | והם |  |
| 9:20.137 | האלוהות |  |
| 9:20.138 | הזכרים |  |
| 9:20.139 | והנקבות | , |
| 9:20.140 | אשר |  |
| 9:20.141 | כל |  |
| 9:20.142 | עבודת |  |
| 9:20.143 | עובדיהם |  |
| 9:20.144 | לתקנם |  |
| 9:20.145 | ולקשט |  |
| 9:20.146 | הנקבות |  |
| 9:20.147 | כדי |  |
| 9:20.148 | למצוא |  |
| 9:20.149 | חן |  |
| 9:20.150 | בעיני |  |
| 9:20.151 | האלוהות |  |
| 9:20.152 | הזכרים |  |
| 9:20.153 | שיזדווגו |  |
| 9:20.154 | להן |  |
| 9:20.155 | כדי |  |
| 9:20.156 | להוריד |  |
| 9:20.157 | שפע |  |
| 9:20.158 | רוחניותם |  |
| 9:20.159 | גם |  |
| 9:20.160 | לעולם |  |
| 9:20.161 | התחתון | , |
| 9:20.162 | כדרשת |  |
| 9:20.163 | רז"ל |  |
| 9:20.164 | על |  |
| 9:20.165 | והנה |  |
| 9:20.166 | עמד |  |
| 9:20.167 | על |  |
| 9:20.168 | היאור |  |
| 9:20.169 | לעיל |  |
| 9:20.170 | סי | ' |
| 9:20.171 | ע"ה | , |
| 9:20.172 | שהאלהים |  |
| 9:20.173 | האחרים |  |
| 9:20.174 | אלמלא |  |
| 9:20.175 | עובדיהם |  |
| 9:20.176 | שומרים |  |
| 9:20.177 | אותם |  |
| 9:20.178 | ומשגיחים |  |
| 9:20.179 | עליהם |  |
| 9:20.180 | בכל |  |
| 9:20.181 | צרכיהם |  |
| 9:20.182 | ומחזקים |  |
| 9:20.183 | בדקיהם |  |
| 9:20.184 | וקלקולם | , |
| 9:20.185 | אזי |  |
| 9:20.186 | כרעו |  |
| 9:20.187 | ונפלו |  |
| 9:20.188 | וחלפו |  |
| 9:20.189 | ועברו |  |
| 9:20.190 | ובטלו | , |
| 9:20.191 | וכזה |  |
| 9:20.192 | היא |  |
| 9:20.193 | דעת |  |
| 9:20.194 | המקובלים |  |
| 9:20.195 | החדשים |  |
| 9:20.196 | שכל |  |
| 9:20.197 | מעשה |  |
| 9:20.198 | מצותיהם |  |
| 9:20.199 | עושים |  |
| 9:20.200 | לשם |  |
| 9:20.201 | האלילים |  |
| 9:20.202 | לתקנם |  |
| 9:20.203 | כמבואר |  |
| 9:20.204 | בעץ |  |
| 9:20.205 | חיים |  |
| 9:20.206 | שער |  |
| 9:20.207 | נקודים |  |
| 9:20.208 | סוף |  |
| 9:20.209 | פ"ו |  |
| 9:20.210 | שקודם |  |
| 9:20.211 | אדה"ר |  |
| 9:20.212 | היו |  |
| 9:20.213 | זעיר |  |
| 9:20.214 | ונוקביה |  |
| 9:20.215 | אחור |  |
| 9:20.216 | באחור |  |
| 9:20.217 | וכו | ' |
| 9:20.218 | וכשנולד |  |
| 9:20.219 | אדה"ר | , |
| 9:20.220 | על |  |
| 9:20.221 | ידי |  |
| 9:20.222 | מעשיו |  |
| 9:20.223 | הטובים |  |
| 9:20.224 | החזירם |  |
| 9:20.225 | פנים |  |
| 9:20.226 | בפנים |  |
| 9:20.227 | ודבריו |  |
| 9:20.228 | אלה |  |
| 9:20.229 | היפך |  |
| 9:20.230 | זהר |  |
| 9:20.231 | בראשית |  |
| 9:20.232 | כ"ב | ) |
| 9:20.233 | אבל |  |
| 9:20.234 | אבא |  |
| 9:20.235 | ואימא |  |
| 9:20.236 | לא |  |
| 9:20.237 | הוצרכו |  |
| 9:20.238 | למעשיו |  |
| 9:20.239 | כי |  |
| 9:20.240 | על |  |
| 9:20.241 | ידי |  |
| 9:20.242 | עצמן |  |
| 9:20.243 | חזרו |  |
| 9:20.244 | פנים |  |
| 9:20.245 | בפנים | , |
| 9:20.246 | ובזהר |  |
| 9:20.247 | שם |  |
| 9:20.248 | אמר |  |
| 9:20.249 | שנתגרשה |  |
| 9:20.250 | אימא |  |
| 9:20.251 | על |  |
| 9:20.252 | ידי |  |
| 9:20.253 | מעשיו | ). |
| 9:20.254 | ובס | ' |
| 9:20.255 | חיי |  |
| 9:20.256 | שלום |  |
| 9:20.257 | בשם |  |
| 9:20.258 | ס | ' |
| 9:20.259 | הלקוטים |  |
| 9:20.260 | כתב |  |
| 9:20.261 | כי |  |
| 9:20.262 | עד |  |
| 9:20.263 | אברהם |  |
| 9:20.264 | אבינו |  |
| 9:20.265 | עדיין |  |
| 9:20.266 | לא |  |
| 9:20.267 | נתקן |  |
| 9:20.268 | הזעיר |  |
| 9:20.269 | אנפין |  |
| 9:20.270 | וכו | ' |
| 9:20.271 | ע"ש | , |
| 9:20.272 | הכלל |  |
| 9:20.273 | העולה |  |
| 9:20.274 | כי |  |
| 9:20.275 | האלוה |  |
| 9:20.276 | קצר |  |
| 9:20.277 | האפים | , |
| 9:20.278 | שעובדים |  |
| 9:20.279 | אותו |  |
| 9:20.280 | המקובלים |  |
| 9:20.281 | החדשים |  |
| 9:20.282 | הוא |  |
| 9:20.283 | צריך |  |
| 9:20.284 | תיקון |  |
| 9:20.285 | בכל |  |
| 9:20.286 | עת |  |
| 9:20.287 | וזמן |  |
| 9:20.288 | עם |  |
| 9:20.289 | שתי |  |
| 9:20.290 | נשיו |  |
| 9:20.291 | רחל |  |
| 9:20.292 | ולאה |  |
| 9:20.293 | כאשר |  |
| 9:20.294 | יראה |  |
| 9:20.295 | המעיין |  |
| 9:20.296 | בסדר |  |
| 9:20.297 | תיקון |  |
| 9:20.298 | חצות |  |
| 9:20.299 | בספר |  |
| 9:20.300 | משנת |  |
| 9:20.301 | חסידים |  |
| 9:20.302 | ובס | ' |
| 9:20.303 | פרי |  |
| 9:20.304 | עץ |  |
| 9:20.305 | חיים |  |
| 9:20.306 | לקשטם |  |
| 9:20.307 | לו |  |
| 9:20.308 | כדי |  |
| 9:20.309 | שיתחמם |  |
| 9:20.310 | הזעיר |  |
| 9:20.311 | ויתאוה |  |
| 9:20.312 | להם |  |
| 9:20.313 | משנת |  |
| 9:20.314 | חסידים |  |
| 9:20.315 | תיקון |  |
| 9:20.316 | חצות |  |
| 9:20.317 | פ"ד |  |
| 9:20.318 | משנה |  |
| 9:20.319 | ג | '). |
| 9:20.320 | אבל |  |
| 9:20.321 | אלהינו |  |
| 9:20.322 | בהו"ש |  |
| 9:20.323 | אינו |  |
| 9:20.324 | צריך |  |
| 9:20.325 | לבריותיו |  |
| 9:20.326 | כלל | , |
| 9:20.327 | רק |  |
| 9:20.328 | הם |  |
| 9:20.329 | צריכים |  |
| 9:20.330 | לו | , |
| 9:20.331 | והוא |  |
| 9:20.332 | שומר |  |
| 9:20.333 | את |  |
| 9:20.334 | יראיו |  |
| 9:20.335 | וחסידיו |  |
| 9:20.336 | מכל |  |
| 9:20.337 | רע | , |
| 9:20.338 | כמבואר |  |
| 9:20.339 | בכמה |  |
| 9:20.340 | מקומות |  |
| 9:20.341 | בכתובים |  |
| 9:20.342 | ובדברי |  |
| 9:20.343 | רז"ל | , |
| 9:20.344 | ואין |  |
| 9:20.345 | צורך |  |
| 9:20.346 | להאריך |  |
| 9:20.347 | בהם |  |
| 9:20.348 | כי |  |
| 9:20.349 | די |  |
| 9:20.350 | בעיקר |  |
| 9:20.351 | שכתב |  |
| 9:20.352 | הרמב"ם |  |
| 9:20.353 | וכמ"ש |  |
| 9:20.354 | רז"ל |  |
| 9:20.355 | והנה |  |
| 9:20.356 | ה | ' |
| 9:20.357 | נצב |  |
| 9:20.358 | עליו |  |
| 9:20.359 | לשמרו |  |
| 9:20.360 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

