# Translation pack: milhamot-hashem perek-6-12 (Milchamot ha-Shem 6:18–6:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-12`.

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
## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- משכילים / השכלה: "the educated" / "enlightenment" in the sense of the
  Haskalah only where the paragraph speaks of that movement; otherwise
  "those of understanding".
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


## Milchamot ha-Shem 6:18

Hebrew (Wikisource): ועוד לא שבע להטיל קלון ופחיתות במשנה שהיא המצוה שקיבל משה מסיני כדאיתא בברכות והובא לעיל, וקורא לה שפחה, סלע אחרא, דאתקריאת קליפה, משנה. נוקבא דעבד נער, עלה אתמר בדברים לא יוסר עבד אלא דמחן ומתברין מנה כמה פסקות בלא נביעו דחכמתא. ומשבח הוא את תורתו החדשה המלאה אלוהות רבים וקורא לה ברתא דמלכא, דבגינה אמתר ודברתם אל הסלע בדבור ופיוס כברתא דמלכא, ומספר שם שמשה רבינו התודה את חטאתו זאת שנתן לנו את המשנה והתלמוד שהם סותרים ומכים על קדקד אמונתו באלהים רבים שהוא קורא לה ברתא דמלכא. ושעל כן נגזר עליו מיתה. דמאן דסריב למטרוניתא חייב מיתה, כל שכן מאן דמחי לברתא דמלכא ושעל כן נגזר עליו שלא יכנס לארץ ישראל. ושיקבר בארץ נכריה חוצה לארץ.

| id | word | punctuation after |
|---|---|---|
| 6:18.1 | ועוד |  |
| 6:18.2 | לא |  |
| 6:18.3 | שבע |  |
| 6:18.4 | להטיל |  |
| 6:18.5 | קלון |  |
| 6:18.6 | ופחיתות |  |
| 6:18.7 | במשנה |  |
| 6:18.8 | שהיא |  |
| 6:18.9 | המצוה |  |
| 6:18.10 | שקיבל |  |
| 6:18.11 | משה |  |
| 6:18.12 | מסיני |  |
| 6:18.13 | כדאיתא |  |
| 6:18.14 | בברכות |  |
| 6:18.15 | והובא |  |
| 6:18.16 | לעיל | , |
| 6:18.17 | וקורא |  |
| 6:18.18 | לה |  |
| 6:18.19 | שפחה | , |
| 6:18.20 | סלע |  |
| 6:18.21 | אחרא | , |
| 6:18.22 | דאתקריאת |  |
| 6:18.23 | קליפה | , |
| 6:18.24 | משנה | . |
| 6:18.25 | נוקבא |  |
| 6:18.26 | דעבד |  |
| 6:18.27 | נער | , |
| 6:18.28 | עלה |  |
| 6:18.29 | אתמר |  |
| 6:18.30 | בדברים |  |
| 6:18.31 | לא |  |
| 6:18.32 | יוסר |  |
| 6:18.33 | עבד |  |
| 6:18.34 | אלא |  |
| 6:18.35 | דמחן |  |
| 6:18.36 | ומתברין |  |
| 6:18.37 | מנה |  |
| 6:18.38 | כמה |  |
| 6:18.39 | פסקות |  |
| 6:18.40 | בלא |  |
| 6:18.41 | נביעו |  |
| 6:18.42 | דחכמתא | . |
| 6:18.43 | ומשבח |  |
| 6:18.44 | הוא |  |
| 6:18.45 | את |  |
| 6:18.46 | תורתו |  |
| 6:18.47 | החדשה |  |
| 6:18.48 | המלאה |  |
| 6:18.49 | אלוהות |  |
| 6:18.50 | רבים |  |
| 6:18.51 | וקורא |  |
| 6:18.52 | לה |  |
| 6:18.53 | ברתא |  |
| 6:18.54 | דמלכא | , |
| 6:18.55 | דבגינה |  |
| 6:18.56 | אמתר |  |
| 6:18.57 | ודברתם |  |
| 6:18.58 | אל |  |
| 6:18.59 | הסלע |  |
| 6:18.60 | בדבור |  |
| 6:18.61 | ופיוס |  |
| 6:18.62 | כברתא |  |
| 6:18.63 | דמלכא | , |
| 6:18.64 | ומספר |  |
| 6:18.65 | שם |  |
| 6:18.66 | שמשה |  |
| 6:18.67 | רבינו |  |
| 6:18.68 | התודה |  |
| 6:18.69 | את |  |
| 6:18.70 | חטאתו |  |
| 6:18.71 | זאת |  |
| 6:18.72 | שנתן |  |
| 6:18.73 | לנו |  |
| 6:18.74 | את |  |
| 6:18.75 | המשנה |  |
| 6:18.76 | והתלמוד |  |
| 6:18.77 | שהם |  |
| 6:18.78 | סותרים |  |
| 6:18.79 | ומכים |  |
| 6:18.80 | על |  |
| 6:18.81 | קדקד |  |
| 6:18.82 | אמונתו |  |
| 6:18.83 | באלהים |  |
| 6:18.84 | רבים |  |
| 6:18.85 | שהוא |  |
| 6:18.86 | קורא |  |
| 6:18.87 | לה |  |
| 6:18.88 | ברתא |  |
| 6:18.89 | דמלכא | . |
| 6:18.90 | ושעל |  |
| 6:18.91 | כן |  |
| 6:18.92 | נגזר |  |
| 6:18.93 | עליו |  |
| 6:18.94 | מיתה | . |
| 6:18.95 | דמאן |  |
| 6:18.96 | דסריב |  |
| 6:18.97 | למטרוניתא |  |
| 6:18.98 | חייב |  |
| 6:18.99 | מיתה | , |
| 6:18.100 | כל |  |
| 6:18.101 | שכן |  |
| 6:18.102 | מאן |  |
| 6:18.103 | דמחי |  |
| 6:18.104 | לברתא |  |
| 6:18.105 | דמלכא |  |
| 6:18.106 | ושעל |  |
| 6:18.107 | כן |  |
| 6:18.108 | נגזר |  |
| 6:18.109 | עליו |  |
| 6:18.110 | שלא |  |
| 6:18.111 | יכנס |  |
| 6:18.112 | לארץ |  |
| 6:18.113 | ישראל | . |
| 6:18.114 | ושיקבר |  |
| 6:18.115 | בארץ |  |
| 6:18.116 | נכריה |  |
| 6:18.117 | חוצה |  |
| 6:18.118 | לארץ | . |


## Milchamot ha-Shem 6:19

Hebrew (Wikisource): סט) ראה והתבונן אחי איך בדבריו אלה הוא מחלל את הקדשים כללותיה ופרטותיה של תורה המקובלים ביד כל ישראל הכתובים במשנה ובתלמוד שקיבל משה מסיני ומסרם ליהושע וכו' הקדושים לכלל האומה הישראלית, ומחבר הזהר מחללם ומורה לעזבם ולפרשו מהם כאבן משכית, וקורא להם בשמו של גנאי שפחה כי תירש גברתה, אשר תחתיה רגזה ארץ, ובשם קליפה קשה אשר ילאה כל אדם לשברה ולהוציא ממנה אוכל מועט כנטפי מים מועטים אשר לא ירוו לנפש עיף וצמא. כה עשה הפילסוף הרמאי מחבר הזהר לגרע נטפי מי ים המשנה והתלמוד בכמה מיני מגרעות אשר שת עליהם בזהר ובתיקונים. ובמחשבה רעה זו חשב לפסול ולהשכיח את המשנה והתלמוד מישראל ולסתום את פי באר מים חיים אשר ממנו יקרו כל הליכות תורה שבעל פה. ומחשבתו זאת בפירוש אמרה בכמה מקומות, וכמו שכ' ר' יצחק דלטאש בפסק המודפס בתחלת ס' הזהר הבנוי על דעת מחבר הזהר שלא כרצון רז"ל בעלי התלמוד, ומהם גם כן בתיקוני זהר חדש בתחלתו דבור המתחיל והמשכילים, שכ' שם דאסתכם קב"ה ושכינתיה למעבר האי הבורא וכו' (והמשכיל יבין איך אפשר שיסכים הקב"ה לשנות מה שכתוב בתורה כי שאל נא לימים ראשונים וכו' וכמו שביארו רז"ל).

| id | word | punctuation after |
|---|---|---|
| 6:19.1 | סט | ) |
| 6:19.2 | ראה |  |
| 6:19.3 | והתבונן |  |
| 6:19.4 | אחי |  |
| 6:19.5 | איך |  |
| 6:19.6 | בדבריו |  |
| 6:19.7 | אלה |  |
| 6:19.8 | הוא |  |
| 6:19.9 | מחלל |  |
| 6:19.10 | את |  |
| 6:19.11 | הקדשים |  |
| 6:19.12 | כללותיה |  |
| 6:19.13 | ופרטותיה |  |
| 6:19.14 | של |  |
| 6:19.15 | תורה |  |
| 6:19.16 | המקובלים |  |
| 6:19.17 | ביד |  |
| 6:19.18 | כל |  |
| 6:19.19 | ישראל |  |
| 6:19.20 | הכתובים |  |
| 6:19.21 | במשנה |  |
| 6:19.22 | ובתלמוד |  |
| 6:19.23 | שקיבל |  |
| 6:19.24 | משה |  |
| 6:19.25 | מסיני |  |
| 6:19.26 | ומסרם |  |
| 6:19.27 | ליהושע |  |
| 6:19.28 | וכו | ' |
| 6:19.29 | הקדושים |  |
| 6:19.30 | לכלל |  |
| 6:19.31 | האומה |  |
| 6:19.32 | הישראלית | , |
| 6:19.33 | ומחבר |  |
| 6:19.34 | הזהר |  |
| 6:19.35 | מחללם |  |
| 6:19.36 | ומורה |  |
| 6:19.37 | לעזבם |  |
| 6:19.38 | ולפרשו |  |
| 6:19.39 | מהם |  |
| 6:19.40 | כאבן |  |
| 6:19.41 | משכית | , |
| 6:19.42 | וקורא |  |
| 6:19.43 | להם |  |
| 6:19.44 | בשמו |  |
| 6:19.45 | של |  |
| 6:19.46 | גנאי |  |
| 6:19.47 | שפחה |  |
| 6:19.48 | כי |  |
| 6:19.49 | תירש |  |
| 6:19.50 | גברתה | , |
| 6:19.51 | אשר |  |
| 6:19.52 | תחתיה |  |
| 6:19.53 | רגזה |  |
| 6:19.54 | ארץ | , |
| 6:19.55 | ובשם |  |
| 6:19.56 | קליפה |  |
| 6:19.57 | קשה |  |
| 6:19.58 | אשר |  |
| 6:19.59 | ילאה |  |
| 6:19.60 | כל |  |
| 6:19.61 | אדם |  |
| 6:19.62 | לשברה |  |
| 6:19.63 | ולהוציא |  |
| 6:19.64 | ממנה |  |
| 6:19.65 | אוכל |  |
| 6:19.66 | מועט |  |
| 6:19.67 | כנטפי |  |
| 6:19.68 | מים |  |
| 6:19.69 | מועטים |  |
| 6:19.70 | אשר |  |
| 6:19.71 | לא |  |
| 6:19.72 | ירוו |  |
| 6:19.73 | לנפש |  |
| 6:19.74 | עיף |  |
| 6:19.75 | וצמא | . |
| 6:19.76 | כה |  |
| 6:19.77 | עשה |  |
| 6:19.78 | הפילסוף |  |
| 6:19.79 | הרמאי |  |
| 6:19.80 | מחבר |  |
| 6:19.81 | הזהר |  |
| 6:19.82 | לגרע |  |
| 6:19.83 | נטפי |  |
| 6:19.84 | מי |  |
| 6:19.85 | ים |  |
| 6:19.86 | המשנה |  |
| 6:19.87 | והתלמוד |  |
| 6:19.88 | בכמה |  |
| 6:19.89 | מיני |  |
| 6:19.90 | מגרעות |  |
| 6:19.91 | אשר |  |
| 6:19.92 | שת |  |
| 6:19.93 | עליהם |  |
| 6:19.94 | בזהר |  |
| 6:19.95 | ובתיקונים | . |
| 6:19.96 | ובמחשבה |  |
| 6:19.97 | רעה |  |
| 6:19.98 | זו |  |
| 6:19.99 | חשב |  |
| 6:19.100 | לפסול |  |
| 6:19.101 | ולהשכיח |  |
| 6:19.102 | את |  |
| 6:19.103 | המשנה |  |
| 6:19.104 | והתלמוד |  |
| 6:19.105 | מישראל |  |
| 6:19.106 | ולסתום |  |
| 6:19.107 | את |  |
| 6:19.108 | פי |  |
| 6:19.109 | באר |  |
| 6:19.110 | מים |  |
| 6:19.111 | חיים |  |
| 6:19.112 | אשר |  |
| 6:19.113 | ממנו |  |
| 6:19.114 | יקרו |  |
| 6:19.115 | כל |  |
| 6:19.116 | הליכות |  |
| 6:19.117 | תורה |  |
| 6:19.118 | שבעל |  |
| 6:19.119 | פה | . |
| 6:19.120 | ומחשבתו |  |
| 6:19.121 | זאת |  |
| 6:19.122 | בפירוש |  |
| 6:19.123 | אמרה |  |
| 6:19.124 | בכמה |  |
| 6:19.125 | מקומות | , |
| 6:19.126 | וכמו |  |
| 6:19.127 | שכ | ' |
| 6:19.128 | ר | ' |
| 6:19.129 | יצחק |  |
| 6:19.130 | דלטאש |  |
| 6:19.131 | בפסק |  |
| 6:19.132 | המודפס |  |
| 6:19.133 | בתחלת |  |
| 6:19.134 | ס | ' |
| 6:19.135 | הזהר |  |
| 6:19.136 | הבנוי |  |
| 6:19.137 | על |  |
| 6:19.138 | דעת |  |
| 6:19.139 | מחבר |  |
| 6:19.140 | הזהר |  |
| 6:19.141 | שלא |  |
| 6:19.142 | כרצון |  |
| 6:19.143 | רז"ל |  |
| 6:19.144 | בעלי |  |
| 6:19.145 | התלמוד | , |
| 6:19.146 | ומהם |  |
| 6:19.147 | גם |  |
| 6:19.148 | כן |  |
| 6:19.149 | בתיקוני |  |
| 6:19.150 | זהר |  |
| 6:19.151 | חדש |  |
| 6:19.152 | בתחלתו |  |
| 6:19.153 | דבור |  |
| 6:19.154 | המתחיל |  |
| 6:19.155 | והמשכילים | , |
| 6:19.156 | שכ | ' |
| 6:19.157 | שם |  |
| 6:19.158 | דאסתכם |  |
| 6:19.159 | קב"ה |  |
| 6:19.160 | ושכינתיה |  |
| 6:19.161 | למעבר |  |
| 6:19.162 | האי |  |
| 6:19.163 | הבורא |  |
| 6:19.164 | וכו | ' |
| 6:19.165 | והמשכיל |  |
| 6:19.166 | יבין |  |
| 6:19.167 | איך |  |
| 6:19.168 | אפשר |  |
| 6:19.169 | שיסכים |  |
| 6:19.170 | הקב"ה |  |
| 6:19.171 | לשנות |  |
| 6:19.172 | מה |  |
| 6:19.173 | שכתוב |  |
| 6:19.174 | בתורה |  |
| 6:19.175 | כי |  |
| 6:19.176 | שאל |  |
| 6:19.177 | נא |  |
| 6:19.178 | לימים |  |
| 6:19.179 | ראשונים |  |
| 6:19.180 | וכו | ' |
| 6:19.181 | וכמו |  |
| 6:19.182 | שביארו |  |
| 6:19.183 | רז"ל | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

