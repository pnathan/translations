# Translation pack: milhamot-hashem perek-5-07 (Milchamot ha-Shem 5:13–5:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-07`.

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
- כלא חד / כולא חד: the same English at every occurrence (1:5-1:7); the
  owner's ruling on the sense is pending; until then, the alternative is in
  `apparatus`.
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
- דברי תימה: "astonishing words" (words that astonish).
- והוא שלא: "provided that" (fixed).
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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


## Milchamot ha-Shem 5:13

Hebrew (Wikisource): מה) בין תבין ותשכיל ידידנו הקורא בשום לב איך שגו רבותינו המקובלים החדשים להטיל דופי בידיעת צופה ומביט עד סוף כל הדורות דבר שיהיה קודם שיהיה והשוו ידיעתו יתברך לידיעתם היפך דברי קדמונינו וחכמינו אנשי כנסת הגדולה מתקני התפילות והברכות, וגם היפך דברי התנא שאמר (אבות ג' ט"ו): "הכל צפוי (והרשות נתונה)" - עבר והוה ועתיד. והיפך דברי רבינו סעדיה גאון בספר האמונות והדעות והרמב"ם בספר המדע ובפירוש המשנה ובמורה. ולפי דבריהם ידיעתו יתברך מסופקת, לפי שאפשר לו לאדם לבחור בזולת מה שידע השי"ת הואיל והוא בעל בחירה! ומה שהביאם לזה הוא מפני השוואתם ידיעתו יתברך לידיעתם. ולא שמו אל לבם דברי הנביא שאמר (ישעיה נ"ה ח') "כי לא מחשבותי" וכו'. ולא די זה אלא שסלקו מהסיבה הראשונה ההשגחה במעשה התחתונים ואמרו איכה ידע אל בעתיד ואין דעה בעליון (עתיקא קדישא), ולפי דבריהם אין הפרש למעלה בין צדיק לרשע כמו שכתב האר"י, שכמו יעקב כן עשו אין הפרש ביניהם כטוב כחוטא. ורק בזעיר אנפין תלוי ההשגחה, והוא הדן ומשלם לאיש איש כפעלו וכמעלליו ישיב לו. ודבר זה מבואר באדרא פרשת האזינו דף רצ"ב דבזעיר אנפין כתיב כי אל דעות ה' ולו נתכנו עלילות בגין דאיהו בתרי גווני לו נתכנו עלילות. אבל לעתיקא קדישא סתימאה לא נתכנו. מאי טעמא נתכנו להאי בגין דירית תרין חולקין, וכתיב עם חסיד תתחסד וכו' ע"כ. ופירש בגיליון בשם הרש"ב הוא כמו שאמר בפרשת בלק "שולטנו יהיב ליה על כלא דכלא יפלחון ליה בגין דאעטרו ליה בדינא ורחמי. מאן דזכי לדינא לדינא ומאן דזכי לרחמי רחמי". מבואר מזה כי למעלה בנאצלים העליונים, לפי דעת הפילוסוף מחבר הזהר והמקובלים, אין מקפידים על מעשה התחתונים, כל שכן בסיבה הראשונה. ורק זעיר אנפין לבדו לו נתכנו עלילות, אבל למעלה לא נתכנו.

| id | word | punctuation after |
|---|---|---|
| 5:13.1 | מה | ) |
| 5:13.2 | בין |  |
| 5:13.3 | תבין |  |
| 5:13.4 | ותשכיל |  |
| 5:13.5 | ידידנו |  |
| 5:13.6 | הקורא |  |
| 5:13.7 | בשום |  |
| 5:13.8 | לב |  |
| 5:13.9 | איך |  |
| 5:13.10 | שגו |  |
| 5:13.11 | רבותינו |  |
| 5:13.12 | המקובלים |  |
| 5:13.13 | החדשים |  |
| 5:13.14 | להטיל |  |
| 5:13.15 | דופי |  |
| 5:13.16 | בידיעת |  |
| 5:13.17 | צופה |  |
| 5:13.18 | ומביט |  |
| 5:13.19 | עד |  |
| 5:13.20 | סוף |  |
| 5:13.21 | כל |  |
| 5:13.22 | הדורות |  |
| 5:13.23 | דבר |  |
| 5:13.24 | שיהיה |  |
| 5:13.25 | קודם |  |
| 5:13.26 | שיהיה |  |
| 5:13.27 | והשוו |  |
| 5:13.28 | ידיעתו |  |
| 5:13.29 | יתברך |  |
| 5:13.30 | לידיעתם |  |
| 5:13.31 | היפך |  |
| 5:13.32 | דברי |  |
| 5:13.33 | קדמונינו |  |
| 5:13.34 | וחכמינו |  |
| 5:13.35 | אנשי |  |
| 5:13.36 | כנסת |  |
| 5:13.37 | הגדולה |  |
| 5:13.38 | מתקני |  |
| 5:13.39 | התפילות |  |
| 5:13.40 | והברכות | , |
| 5:13.41 | וגם |  |
| 5:13.42 | היפך |  |
| 5:13.43 | דברי |  |
| 5:13.44 | התנא |  |
| 5:13.45 | שאמר |  |
| 5:13.46 | אבות |  |
| 5:13.47 | ג | ' |
| 5:13.48 | ט"ו | ): |
| 5:13.49 | הכל |  |
| 5:13.50 | צפוי |  |
| 5:13.51 | והרשות |  |
| 5:13.52 | נתונה | )"- |
| 5:13.53 | עבר |  |
| 5:13.54 | והוה |  |
| 5:13.55 | ועתיד | . |
| 5:13.56 | והיפך |  |
| 5:13.57 | דברי |  |
| 5:13.58 | רבינו |  |
| 5:13.59 | סעדיה |  |
| 5:13.60 | גאון |  |
| 5:13.61 | בספר |  |
| 5:13.62 | האמונות |  |
| 5:13.63 | והדעות |  |
| 5:13.64 | והרמב"ם |  |
| 5:13.65 | בספר |  |
| 5:13.66 | המדע |  |
| 5:13.67 | ובפירוש |  |
| 5:13.68 | המשנה |  |
| 5:13.69 | ובמורה | . |
| 5:13.70 | ולפי |  |
| 5:13.71 | דבריהם |  |
| 5:13.72 | ידיעתו |  |
| 5:13.73 | יתברך |  |
| 5:13.74 | מסופקת | , |
| 5:13.75 | לפי |  |
| 5:13.76 | שאפשר |  |
| 5:13.77 | לו |  |
| 5:13.78 | לאדם |  |
| 5:13.79 | לבחור |  |
| 5:13.80 | בזולת |  |
| 5:13.81 | מה |  |
| 5:13.82 | שידע |  |
| 5:13.83 | השי"ת |  |
| 5:13.84 | הואיל |  |
| 5:13.85 | והוא |  |
| 5:13.86 | בעל |  |
| 5:13.87 | בחירה | ! |
| 5:13.88 | ומה |  |
| 5:13.89 | שהביאם |  |
| 5:13.90 | לזה |  |
| 5:13.91 | הוא |  |
| 5:13.92 | מפני |  |
| 5:13.93 | השוואתם |  |
| 5:13.94 | ידיעתו |  |
| 5:13.95 | יתברך |  |
| 5:13.96 | לידיעתם | . |
| 5:13.97 | ולא |  |
| 5:13.98 | שמו |  |
| 5:13.99 | אל |  |
| 5:13.100 | לבם |  |
| 5:13.101 | דברי |  |
| 5:13.102 | הנביא |  |
| 5:13.103 | שאמר |  |
| 5:13.104 | ישעיה |  |
| 5:13.105 | נ"ה |  |
| 5:13.106 | ח | ') |
| 5:13.107 | כי |  |
| 5:13.108 | לא |  |
| 5:13.109 | מחשבותי | " |
| 5:13.110 | וכו | '. |
| 5:13.111 | ולא |  |
| 5:13.112 | די |  |
| 5:13.113 | זה |  |
| 5:13.114 | אלא |  |
| 5:13.115 | שסלקו |  |
| 5:13.116 | מהסיבה |  |
| 5:13.117 | הראשונה |  |
| 5:13.118 | ההשגחה |  |
| 5:13.119 | במעשה |  |
| 5:13.120 | התחתונים |  |
| 5:13.121 | ואמרו |  |
| 5:13.122 | איכה |  |
| 5:13.123 | ידע |  |
| 5:13.124 | אל |  |
| 5:13.125 | בעתיד |  |
| 5:13.126 | ואין |  |
| 5:13.127 | דעה |  |
| 5:13.128 | בעליון |  |
| 5:13.129 | עתיקא |  |
| 5:13.130 | קדישא | ), |
| 5:13.131 | ולפי |  |
| 5:13.132 | דבריהם |  |
| 5:13.133 | אין |  |
| 5:13.134 | הפרש |  |
| 5:13.135 | למעלה |  |
| 5:13.136 | בין |  |
| 5:13.137 | צדיק |  |
| 5:13.138 | לרשע |  |
| 5:13.139 | כמו |  |
| 5:13.140 | שכתב |  |
| 5:13.141 | האר"י | , |
| 5:13.142 | שכמו |  |
| 5:13.143 | יעקב |  |
| 5:13.144 | כן |  |
| 5:13.145 | עשו |  |
| 5:13.146 | אין |  |
| 5:13.147 | הפרש |  |
| 5:13.148 | ביניהם |  |
| 5:13.149 | כטוב |  |
| 5:13.150 | כחוטא | . |
| 5:13.151 | ורק |  |
| 5:13.152 | בזעיר |  |
| 5:13.153 | אנפין |  |
| 5:13.154 | תלוי |  |
| 5:13.155 | ההשגחה | , |
| 5:13.156 | והוא |  |
| 5:13.157 | הדן |  |
| 5:13.158 | ומשלם |  |
| 5:13.159 | לאיש |  |
| 5:13.160 | איש |  |
| 5:13.161 | כפעלו |  |
| 5:13.162 | וכמעלליו |  |
| 5:13.163 | ישיב |  |
| 5:13.164 | לו | . |
| 5:13.165 | ודבר |  |
| 5:13.166 | זה |  |
| 5:13.167 | מבואר |  |
| 5:13.168 | באדרא |  |
| 5:13.169 | פרשת |  |
| 5:13.170 | האזינו |  |
| 5:13.171 | דף |  |
| 5:13.172 | רצ"ב |  |
| 5:13.173 | דבזעיר |  |
| 5:13.174 | אנפין |  |
| 5:13.175 | כתיב |  |
| 5:13.176 | כי |  |
| 5:13.177 | אל |  |
| 5:13.178 | דעות |  |
| 5:13.179 | ה | ' |
| 5:13.180 | ולו |  |
| 5:13.181 | נתכנו |  |
| 5:13.182 | עלילות |  |
| 5:13.183 | בגין |  |
| 5:13.184 | דאיהו |  |
| 5:13.185 | בתרי |  |
| 5:13.186 | גווני |  |
| 5:13.187 | לו |  |
| 5:13.188 | נתכנו |  |
| 5:13.189 | עלילות | . |
| 5:13.190 | אבל |  |
| 5:13.191 | לעתיקא |  |
| 5:13.192 | קדישא |  |
| 5:13.193 | סתימאה |  |
| 5:13.194 | לא |  |
| 5:13.195 | נתכנו | . |
| 5:13.196 | מאי |  |
| 5:13.197 | טעמא |  |
| 5:13.198 | נתכנו |  |
| 5:13.199 | להאי |  |
| 5:13.200 | בגין |  |
| 5:13.201 | דירית |  |
| 5:13.202 | תרין |  |
| 5:13.203 | חולקין | , |
| 5:13.204 | וכתיב |  |
| 5:13.205 | עם |  |
| 5:13.206 | חסיד |  |
| 5:13.207 | תתחסד |  |
| 5:13.208 | וכו | ' |
| 5:13.209 | ע"כ | . |
| 5:13.210 | ופירש |  |
| 5:13.211 | בגיליון |  |
| 5:13.212 | בשם |  |
| 5:13.213 | הרש"ב |  |
| 5:13.214 | הוא |  |
| 5:13.215 | כמו |  |
| 5:13.216 | שאמר |  |
| 5:13.217 | בפרשת |  |
| 5:13.218 | בלק |  |
| 5:13.219 | שולטנו |  |
| 5:13.220 | יהיב |  |
| 5:13.221 | ליה |  |
| 5:13.222 | על |  |
| 5:13.223 | כלא |  |
| 5:13.224 | דכלא |  |
| 5:13.225 | יפלחון |  |
| 5:13.226 | ליה |  |
| 5:13.227 | בגין |  |
| 5:13.228 | דאעטרו |  |
| 5:13.229 | ליה |  |
| 5:13.230 | בדינא |  |
| 5:13.231 | ורחמי | . |
| 5:13.232 | מאן |  |
| 5:13.233 | דזכי |  |
| 5:13.234 | לדינא |  |
| 5:13.235 | לדינא |  |
| 5:13.236 | ומאן |  |
| 5:13.237 | דזכי |  |
| 5:13.238 | לרחמי |  |
| 5:13.239 | רחמי | ". |
| 5:13.240 | מבואר |  |
| 5:13.241 | מזה |  |
| 5:13.242 | כי |  |
| 5:13.243 | למעלה |  |
| 5:13.244 | בנאצלים |  |
| 5:13.245 | העליונים | , |
| 5:13.246 | לפי |  |
| 5:13.247 | דעת |  |
| 5:13.248 | הפילוסוף |  |
| 5:13.249 | מחבר |  |
| 5:13.250 | הזהר |  |
| 5:13.251 | והמקובלים | , |
| 5:13.252 | אין |  |
| 5:13.253 | מקפידים |  |
| 5:13.254 | על |  |
| 5:13.255 | מעשה |  |
| 5:13.256 | התחתונים | , |
| 5:13.257 | כל |  |
| 5:13.258 | שכן |  |
| 5:13.259 | בסיבה |  |
| 5:13.260 | הראשונה | . |
| 5:13.261 | ורק |  |
| 5:13.262 | זעיר |  |
| 5:13.263 | אנפין |  |
| 5:13.264 | לבדו |  |
| 5:13.265 | לו |  |
| 5:13.266 | נתכנו |  |
| 5:13.267 | עלילות | , |
| 5:13.268 | אבל |  |
| 5:13.269 | למעלה |  |
| 5:13.270 | לא |  |
| 5:13.271 | נתכנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

