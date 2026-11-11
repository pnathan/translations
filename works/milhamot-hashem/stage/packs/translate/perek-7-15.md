# Translation pack: milhamot-hashem perek-7-15 (Milchamot ha-Shem 7:20–7:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-15`.

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
- Two verses the author joins with a comma and no וגו': render the comma;
  an ellipsis only where he omits words and marks it (וכו', וגו').
- והוא שלא: "provided that" (fixed).
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
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
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


## Milchamot ha-Shem 7:20

Hebrew (Wikisource): צב) ולא זו בלבד. אלא אף גם זאת שייחסו לו החוש המגונה והפחות שבחמשה חושים, והוא המשגל והזיווג, כי ידוע מה שכתב הרמב"ם ז"ל בפרק מ"ז חלק א'. שמהחמשה חושים שהם הראייה, והשמע, והריח, וישמע ה', וירא ה', וירח ה', ולא הושאל לו חוש הטעם והמשוש לומר ויטעם ה', וימשש ה'. והסבה לזה הוא כמו שנתבאר לעיל שכל מה שידמוהו ההמון חסרון אי אפשר לציירו בבורא ית', ולכן לא השתמשו בזה בספרי הנבואה עכ"ל. וכן כתב הרב שבילי אמונה בנתיב הא'. ובדרך הזה דרכו רז"ל בעלי המשנה והתלמוד ומתקני התפלות והפייטנים כרבי' סעדיה גאון ור"י הלוי בפיוטיהם הנעימים המעוררים את האדם לאהבת השי"ת וליראתו ולדבקה בו. ומה מאד מתקו דברי רבי יהודה הלוי בסליחה המתחלת יה שמך ארוממך, אמר בסופה, השתונן והכונן, והתבונן בסודיך, והבטת מה אתה, ומאין יסודיך, ומי הבינך, ומי הכינך, וכח מי ינידיך, והבט אל גבורות אל, והעירה כבודיך. חקור פעליו, רק אליו אל תשלח ידיך, כי תדרוש בסוף וראש במופלה ובמכוסה ע"כ. אשר על כן אלו שמרו המקובלים החדשים אזהרת רז"ל ולא נפתו לדברי נביא השקר המסית מחבר הזהר ולא שלחו רעיונותיהם במופלא ובמכוסה מהם לא באו לכל זה. ובפרק ח' לשלישי כתב הר' המורה שחוש המשוש הוא חרפה וכלמה לנו, אשר בעבורו נתאוה למאכל ומשתה ומשגל, והאיש המשכיל מריך למעט בהם מה שאפשר לו (כמו שביאר הרב בהלכות דעות) ולא ירחיב הדבור בהם, וישים תכליתו תכלית האדם מאשר הוא אדם, והוא ציור המושכלות, אשר בהם ישיג ידיעת הבורא, אבל האכילה והמשגל הם חרפתינו וכלימתינו, כי הושמו באדם לצורך קיומו וקיום המין, והוא דומה בהם לבהמה, וכל שכן המשגל שאין ראוי לדבר בו יותר כמו שביארתי בפרקי אבות ממה שבא בתורתינו הקדושה ממאוס בו, ואסור לזכרו ולדבר בו כלל, וכבר ידעת אמרם על אלישע הנביא ע"ה שנקרא קדוש מפני שלא היה חושב בו עש שלא היה רואה קרי, וגם ידעת גודל האיסור שבא אצלינו בענין נבלות הפה (כדגרסינן בכתובות, על כן על בחוריו לא ישמח ה' ואת יתומיו ואת אלמנותיו לא ירחם כי כלו חנף ומרע וכל פה דובר נבלה בכל זאת לא שב אפו ועוד ידו נטויה, מאי ועוד ידו נטויה אמר רב חנן ב"א הכל יודעים למה נכנסה כלה לחופה אלא כל המנבל פיו ומוציא דבר נבלה מפיו אפילו נחתם לו גזר דין של שבעים שנה לטובה נהפך עליו לרעה ע"כ. ובפרק במה מדליקים גרסינן בעון נבלות הפה צרות רבות ודעות וגזרות קשות ותחדשות בעולם ובחורי ישראל מתים ואלמנות ויתומים צועקים ואינם נענים שנאמר על כן על בחוריו לא ישמח ה' וכו' כי כלו חנף ומרע כל פה דובר נבלה וכו' מאי ועוד ידו נטויה א"ר חנן בר אבא א"ר חסדא כל המנבל פיו מעמיקים לו גיהנום שנ' שוחה עמוקה פי זרות זעום ה' יפל שם ע"כ). והטעם לפי שהדיבור בלשון הוא מסגולת האדם וטוב השם אשר גמר עליו להבדילו בו משאר בעלי החיים שנאמר מי שם פה לאדם, ואמר הנביא ה' אלהים נתן לי לשון למודים, ואין ראוי שנשתמש בטובה אשר גמלנו השי"ת לשלמות ללמוד וללמד, בדבר פחות ומגונה ובחרפה הגמורה שיש בנו ונדמה לגוים הסכלים הזונים בשיריהם שירי ענבים, ופחיתות אשר יתפארו בהם לרוב סכלות ופחיתותם (כדאמור רבנן תפארת גוים עבירה. גוי מפערה מעי) לא כאלה חלק יעקב, שנאמר בהם ואתר תהיו לי ממלכת כהנים וגוי קדוש, וכל מי שמשתמש במחשבה או בדבריו מענייני המשגל אשר הוא חרפה לנו ביותר מן הצריך לו (כמאמר רז"ל מגיד לאדם מה שיחו אפילו שיחה יתירה שבין איש לאשתו מגידין לו) וכ"ש אם יאמר שירים, הנה נשתמש בטובה אשר גמלו השי"צ בדברי מרי וסכלות גמורה ועבר על מצות הבורא ית'. והוא כמי שנאמר בהם וכסף הרביתי להם וזהב עשו לבעל, זהו תורף דברי הר' המורה ז"ל. גם הרס"ג בס' האמונות מאמר עשירי פרק ו' דבר בגנות רבוי המשגל ולכלוכו וטנופו והחרפה והבוז אשר בו, ושאין התואה הזאת באדם כי אם להקים בה זרע כמצות השי"ת ואתם פרו ורבו וכו' ובהצנע בעת הראוי ויעצור בה מהרבות ע"כ. ובעיני ראיתי אחד מהמון העם שאל לתינוק קטן ואמר האם אביך אוהב את אמך ומנשקה, וגער בו הדיין כמה"ר סלי' צאלח ז"ל ואמר לו כלך מגבלות הפה כי ענשו חמור.

| id | word | punctuation after |
|---|---|---|
| 7:20.1 | צב | ) |
| 7:20.2 | ולא |  |
| 7:20.3 | זו |  |
| 7:20.4 | בלבד | . |
| 7:20.5 | אלא |  |
| 7:20.6 | אף |  |
| 7:20.7 | גם |  |
| 7:20.8 | זאת |  |
| 7:20.9 | שייחסו |  |
| 7:20.10 | לו |  |
| 7:20.11 | החוש |  |
| 7:20.12 | המגונה |  |
| 7:20.13 | והפחות |  |
| 7:20.14 | שבחמשה |  |
| 7:20.15 | חושים | , |
| 7:20.16 | והוא |  |
| 7:20.17 | המשגל |  |
| 7:20.18 | והזיווג | , |
| 7:20.19 | כי |  |
| 7:20.20 | ידוע |  |
| 7:20.21 | מה |  |
| 7:20.22 | שכתב |  |
| 7:20.23 | הרמב"ם |  |
| 7:20.24 | ז"ל |  |
| 7:20.25 | בפרק |  |
| 7:20.26 | מ"ז |  |
| 7:20.27 | חלק |  |
| 7:20.28 | א | '. |
| 7:20.29 | שמהחמשה |  |
| 7:20.30 | חושים |  |
| 7:20.31 | שהם |  |
| 7:20.32 | הראייה | , |
| 7:20.33 | והשמע | , |
| 7:20.34 | והריח | , |
| 7:20.35 | וישמע |  |
| 7:20.36 | ה | ', |
| 7:20.37 | וירא |  |
| 7:20.38 | ה | ', |
| 7:20.39 | וירח |  |
| 7:20.40 | ה | ', |
| 7:20.41 | ולא |  |
| 7:20.42 | הושאל |  |
| 7:20.43 | לו |  |
| 7:20.44 | חוש |  |
| 7:20.45 | הטעם |  |
| 7:20.46 | והמשוש |  |
| 7:20.47 | לומר |  |
| 7:20.48 | ויטעם |  |
| 7:20.49 | ה | ', |
| 7:20.50 | וימשש |  |
| 7:20.51 | ה | '. |
| 7:20.52 | והסבה |  |
| 7:20.53 | לזה |  |
| 7:20.54 | הוא |  |
| 7:20.55 | כמו |  |
| 7:20.56 | שנתבאר |  |
| 7:20.57 | לעיל |  |
| 7:20.58 | שכל |  |
| 7:20.59 | מה |  |
| 7:20.60 | שידמוהו |  |
| 7:20.61 | ההמון |  |
| 7:20.62 | חסרון |  |
| 7:20.63 | אי |  |
| 7:20.64 | אפשר |  |
| 7:20.65 | לציירו |  |
| 7:20.66 | בבורא |  |
| 7:20.67 | ית | ', |
| 7:20.68 | ולכן |  |
| 7:20.69 | לא |  |
| 7:20.70 | השתמשו |  |
| 7:20.71 | בזה |  |
| 7:20.72 | בספרי |  |
| 7:20.73 | הנבואה |  |
| 7:20.74 | עכ"ל | . |
| 7:20.75 | וכן |  |
| 7:20.76 | כתב |  |
| 7:20.77 | הרב |  |
| 7:20.78 | שבילי |  |
| 7:20.79 | אמונה |  |
| 7:20.80 | בנתיב |  |
| 7:20.81 | הא | '. |
| 7:20.82 | ובדרך |  |
| 7:20.83 | הזה |  |
| 7:20.84 | דרכו |  |
| 7:20.85 | רז"ל |  |
| 7:20.86 | בעלי |  |
| 7:20.87 | המשנה |  |
| 7:20.88 | והתלמוד |  |
| 7:20.89 | ומתקני |  |
| 7:20.90 | התפלות |  |
| 7:20.91 | והפייטנים |  |
| 7:20.92 | כרבי | ' |
| 7:20.93 | סעדיה |  |
| 7:20.94 | גאון |  |
| 7:20.95 | ור"י |  |
| 7:20.96 | הלוי |  |
| 7:20.97 | בפיוטיהם |  |
| 7:20.98 | הנעימים |  |
| 7:20.99 | המעוררים |  |
| 7:20.100 | את |  |
| 7:20.101 | האדם |  |
| 7:20.102 | לאהבת |  |
| 7:20.103 | השי"ת |  |
| 7:20.104 | וליראתו |  |
| 7:20.105 | ולדבקה |  |
| 7:20.106 | בו | . |
| 7:20.107 | ומה |  |
| 7:20.108 | מאד |  |
| 7:20.109 | מתקו |  |
| 7:20.110 | דברי |  |
| 7:20.111 | רבי |  |
| 7:20.112 | יהודה |  |
| 7:20.113 | הלוי |  |
| 7:20.114 | בסליחה |  |
| 7:20.115 | המתחלת |  |
| 7:20.116 | יה |  |
| 7:20.117 | שמך |  |
| 7:20.118 | ארוממך | , |
| 7:20.119 | אמר |  |
| 7:20.120 | בסופה | , |
| 7:20.121 | השתונן |  |
| 7:20.122 | והכונן | , |
| 7:20.123 | והתבונן |  |
| 7:20.124 | בסודיך | , |
| 7:20.125 | והבטת |  |
| 7:20.126 | מה |  |
| 7:20.127 | אתה | , |
| 7:20.128 | ומאין |  |
| 7:20.129 | יסודיך | , |
| 7:20.130 | ומי |  |
| 7:20.131 | הבינך | , |
| 7:20.132 | ומי |  |
| 7:20.133 | הכינך | , |
| 7:20.134 | וכח |  |
| 7:20.135 | מי |  |
| 7:20.136 | ינידיך | , |
| 7:20.137 | והבט |  |
| 7:20.138 | אל |  |
| 7:20.139 | גבורות |  |
| 7:20.140 | אל | , |
| 7:20.141 | והעירה |  |
| 7:20.142 | כבודיך | . |
| 7:20.143 | חקור |  |
| 7:20.144 | פעליו | , |
| 7:20.145 | רק |  |
| 7:20.146 | אליו |  |
| 7:20.147 | אל |  |
| 7:20.148 | תשלח |  |
| 7:20.149 | ידיך | , |
| 7:20.150 | כי |  |
| 7:20.151 | תדרוש |  |
| 7:20.152 | בסוף |  |
| 7:20.153 | וראש |  |
| 7:20.154 | במופלה |  |
| 7:20.155 | ובמכוסה |  |
| 7:20.156 | ע"כ | . |
| 7:20.157 | אשר |  |
| 7:20.158 | על |  |
| 7:20.159 | כן |  |
| 7:20.160 | אלו |  |
| 7:20.161 | שמרו |  |
| 7:20.162 | המקובלים |  |
| 7:20.163 | החדשים |  |
| 7:20.164 | אזהרת |  |
| 7:20.165 | רז"ל |  |
| 7:20.166 | ולא |  |
| 7:20.167 | נפתו |  |
| 7:20.168 | לדברי |  |
| 7:20.169 | נביא |  |
| 7:20.170 | השקר |  |
| 7:20.171 | המסית |  |
| 7:20.172 | מחבר |  |
| 7:20.173 | הזהר |  |
| 7:20.174 | ולא |  |
| 7:20.175 | שלחו |  |
| 7:20.176 | רעיונותיהם |  |
| 7:20.177 | במופלא |  |
| 7:20.178 | ובמכוסה |  |
| 7:20.179 | מהם |  |
| 7:20.180 | לא |  |
| 7:20.181 | באו |  |
| 7:20.182 | לכל |  |
| 7:20.183 | זה | . |
| 7:20.184 | ובפרק |  |
| 7:20.185 | ח | ' |
| 7:20.186 | לשלישי |  |
| 7:20.187 | כתב |  |
| 7:20.188 | הר | ' |
| 7:20.189 | המורה |  |
| 7:20.190 | שחוש |  |
| 7:20.191 | המשוש |  |
| 7:20.192 | הוא |  |
| 7:20.193 | חרפה |  |
| 7:20.194 | וכלמה |  |
| 7:20.195 | לנו | , |
| 7:20.196 | אשר |  |
| 7:20.197 | בעבורו |  |
| 7:20.198 | נתאוה |  |
| 7:20.199 | למאכל |  |
| 7:20.200 | ומשתה |  |
| 7:20.201 | ומשגל | , |
| 7:20.202 | והאיש |  |
| 7:20.203 | המשכיל |  |
| 7:20.204 | מריך |  |
| 7:20.205 | למעט |  |
| 7:20.206 | בהם |  |
| 7:20.207 | מה |  |
| 7:20.208 | שאפשר |  |
| 7:20.209 | לו |  |
| 7:20.210 | כמו |  |
| 7:20.211 | שביאר |  |
| 7:20.212 | הרב |  |
| 7:20.213 | בהלכות |  |
| 7:20.214 | דעות | ) |
| 7:20.215 | ולא |  |
| 7:20.216 | ירחיב |  |
| 7:20.217 | הדבור |  |
| 7:20.218 | בהם | , |
| 7:20.219 | וישים |  |
| 7:20.220 | תכליתו |  |
| 7:20.221 | תכלית |  |
| 7:20.222 | האדם |  |
| 7:20.223 | מאשר |  |
| 7:20.224 | הוא |  |
| 7:20.225 | אדם | , |
| 7:20.226 | והוא |  |
| 7:20.227 | ציור |  |
| 7:20.228 | המושכלות | , |
| 7:20.229 | אשר |  |
| 7:20.230 | בהם |  |
| 7:20.231 | ישיג |  |
| 7:20.232 | ידיעת |  |
| 7:20.233 | הבורא | , |
| 7:20.234 | אבל |  |
| 7:20.235 | האכילה |  |
| 7:20.236 | והמשגל |  |
| 7:20.237 | הם |  |
| 7:20.238 | חרפתינו |  |
| 7:20.239 | וכלימתינו | , |
| 7:20.240 | כי |  |
| 7:20.241 | הושמו |  |
| 7:20.242 | באדם |  |
| 7:20.243 | לצורך |  |
| 7:20.244 | קיומו |  |
| 7:20.245 | וקיום |  |
| 7:20.246 | המין | , |
| 7:20.247 | והוא |  |
| 7:20.248 | דומה |  |
| 7:20.249 | בהם |  |
| 7:20.250 | לבהמה | , |
| 7:20.251 | וכל |  |
| 7:20.252 | שכן |  |
| 7:20.253 | המשגל |  |
| 7:20.254 | שאין |  |
| 7:20.255 | ראוי |  |
| 7:20.256 | לדבר |  |
| 7:20.257 | בו |  |
| 7:20.258 | יותר |  |
| 7:20.259 | כמו |  |
| 7:20.260 | שביארתי |  |
| 7:20.261 | בפרקי |  |
| 7:20.262 | אבות |  |
| 7:20.263 | ממה |  |
| 7:20.264 | שבא |  |
| 7:20.265 | בתורתינו |  |
| 7:20.266 | הקדושה |  |
| 7:20.267 | ממאוס |  |
| 7:20.268 | בו | , |
| 7:20.269 | ואסור |  |
| 7:20.270 | לזכרו |  |
| 7:20.271 | ולדבר |  |
| 7:20.272 | בו |  |
| 7:20.273 | כלל | , |
| 7:20.274 | וכבר |  |
| 7:20.275 | ידעת |  |
| 7:20.276 | אמרם |  |
| 7:20.277 | על |  |
| 7:20.278 | אלישע |  |
| 7:20.279 | הנביא |  |
| 7:20.280 | ע"ה |  |
| 7:20.281 | שנקרא |  |
| 7:20.282 | קדוש |  |
| 7:20.283 | מפני |  |
| 7:20.284 | שלא |  |
| 7:20.285 | היה |  |
| 7:20.286 | חושב |  |
| 7:20.287 | בו |  |
| 7:20.288 | עש |  |
| 7:20.289 | שלא |  |
| 7:20.290 | היה |  |
| 7:20.291 | רואה |  |
| 7:20.292 | קרי | , |
| 7:20.293 | וגם |  |
| 7:20.294 | ידעת |  |
| 7:20.295 | גודל |  |
| 7:20.296 | האיסור |  |
| 7:20.297 | שבא |  |
| 7:20.298 | אצלינו |  |
| 7:20.299 | בענין |  |
| 7:20.300 | נבלות |  |
| 7:20.301 | הפה |  |
| 7:20.302 | כדגרסינן |  |
| 7:20.303 | בכתובות | , |
| 7:20.304 | על |  |
| 7:20.305 | כן |  |
| 7:20.306 | על |  |
| 7:20.307 | בחוריו |  |
| 7:20.308 | לא |  |
| 7:20.309 | ישמח |  |
| 7:20.310 | ה | ' |
| 7:20.311 | ואת |  |
| 7:20.312 | יתומיו |  |
| 7:20.313 | ואת |  |
| 7:20.314 | אלמנותיו |  |
| 7:20.315 | לא |  |
| 7:20.316 | ירחם |  |
| 7:20.317 | כי |  |
| 7:20.318 | כלו |  |
| 7:20.319 | חנף |  |
| 7:20.320 | ומרע |  |
| 7:20.321 | וכל |  |
| 7:20.322 | פה |  |
| 7:20.323 | דובר |  |
| 7:20.324 | נבלה |  |
| 7:20.325 | בכל |  |
| 7:20.326 | זאת |  |
| 7:20.327 | לא |  |
| 7:20.328 | שב |  |
| 7:20.329 | אפו |  |
| 7:20.330 | ועוד |  |
| 7:20.331 | ידו |  |
| 7:20.332 | נטויה | , |
| 7:20.333 | מאי |  |
| 7:20.334 | ועוד |  |
| 7:20.335 | ידו |  |
| 7:20.336 | נטויה |  |
| 7:20.337 | אמר |  |
| 7:20.338 | רב |  |
| 7:20.339 | חנן |  |
| 7:20.340 | ב"א |  |
| 7:20.341 | הכל |  |
| 7:20.342 | יודעים |  |
| 7:20.343 | למה |  |
| 7:20.344 | נכנסה |  |
| 7:20.345 | כלה |  |
| 7:20.346 | לחופה |  |
| 7:20.347 | אלא |  |
| 7:20.348 | כל |  |
| 7:20.349 | המנבל |  |
| 7:20.350 | פיו |  |
| 7:20.351 | ומוציא |  |
| 7:20.352 | דבר |  |
| 7:20.353 | נבלה |  |
| 7:20.354 | מפיו |  |
| 7:20.355 | אפילו |  |
| 7:20.356 | נחתם |  |
| 7:20.357 | לו |  |
| 7:20.358 | גזר |  |
| 7:20.359 | דין |  |
| 7:20.360 | של |  |
| 7:20.361 | שבעים |  |
| 7:20.362 | שנה |  |
| 7:20.363 | לטובה |  |
| 7:20.364 | נהפך |  |
| 7:20.365 | עליו |  |
| 7:20.366 | לרעה |  |
| 7:20.367 | ע"כ | . |
| 7:20.368 | ובפרק |  |
| 7:20.369 | במה |  |
| 7:20.370 | מדליקים |  |
| 7:20.371 | גרסינן |  |
| 7:20.372 | בעון |  |
| 7:20.373 | נבלות |  |
| 7:20.374 | הפה |  |
| 7:20.375 | צרות |  |
| 7:20.376 | רבות |  |
| 7:20.377 | ודעות |  |
| 7:20.378 | וגזרות |  |
| 7:20.379 | קשות |  |
| 7:20.380 | ותחדשות |  |
| 7:20.381 | בעולם |  |
| 7:20.382 | ובחורי |  |
| 7:20.383 | ישראל |  |
| 7:20.384 | מתים |  |
| 7:20.385 | ואלמנות |  |
| 7:20.386 | ויתומים |  |
| 7:20.387 | צועקים |  |
| 7:20.388 | ואינם |  |
| 7:20.389 | נענים |  |
| 7:20.390 | שנאמר |  |
| 7:20.391 | על |  |
| 7:20.392 | כן |  |
| 7:20.393 | על |  |
| 7:20.394 | בחוריו |  |
| 7:20.395 | לא |  |
| 7:20.396 | ישמח |  |
| 7:20.397 | ה | ' |
| 7:20.398 | וכו | ' |
| 7:20.399 | כי |  |
| 7:20.400 | כלו |  |
| 7:20.401 | חנף |  |
| 7:20.402 | ומרע |  |
| 7:20.403 | כל |  |
| 7:20.404 | פה |  |
| 7:20.405 | דובר |  |
| 7:20.406 | נבלה |  |
| 7:20.407 | וכו | ' |
| 7:20.408 | מאי |  |
| 7:20.409 | ועוד |  |
| 7:20.410 | ידו |  |
| 7:20.411 | נטויה |  |
| 7:20.412 | א"ר |  |
| 7:20.413 | חנן |  |
| 7:20.414 | בר |  |
| 7:20.415 | אבא |  |
| 7:20.416 | א"ר |  |
| 7:20.417 | חסדא |  |
| 7:20.418 | כל |  |
| 7:20.419 | המנבל |  |
| 7:20.420 | פיו |  |
| 7:20.421 | מעמיקים |  |
| 7:20.422 | לו |  |
| 7:20.423 | גיהנום |  |
| 7:20.424 | שנ | ' |
| 7:20.425 | שוחה |  |
| 7:20.426 | עמוקה |  |
| 7:20.427 | פי |  |
| 7:20.428 | זרות |  |
| 7:20.429 | זעום |  |
| 7:20.430 | ה | ' |
| 7:20.431 | יפל |  |
| 7:20.432 | שם |  |
| 7:20.433 | ע"כ | ). |
| 7:20.434 | והטעם |  |
| 7:20.435 | לפי |  |
| 7:20.436 | שהדיבור |  |
| 7:20.437 | בלשון |  |
| 7:20.438 | הוא |  |
| 7:20.439 | מסגולת |  |
| 7:20.440 | האדם |  |
| 7:20.441 | וטוב |  |
| 7:20.442 | השם |  |
| 7:20.443 | אשר |  |
| 7:20.444 | גמר |  |
| 7:20.445 | עליו |  |
| 7:20.446 | להבדילו |  |
| 7:20.447 | בו |  |
| 7:20.448 | משאר |  |
| 7:20.449 | בעלי |  |
| 7:20.450 | החיים |  |
| 7:20.451 | שנאמר |  |
| 7:20.452 | מי |  |
| 7:20.453 | שם |  |
| 7:20.454 | פה |  |
| 7:20.455 | לאדם | , |
| 7:20.456 | ואמר |  |
| 7:20.457 | הנביא |  |
| 7:20.458 | ה | ' |
| 7:20.459 | אלהים |  |
| 7:20.460 | נתן |  |
| 7:20.461 | לי |  |
| 7:20.462 | לשון |  |
| 7:20.463 | למודים | , |
| 7:20.464 | ואין |  |
| 7:20.465 | ראוי |  |
| 7:20.466 | שנשתמש |  |
| 7:20.467 | בטובה |  |
| 7:20.468 | אשר |  |
| 7:20.469 | גמלנו |  |
| 7:20.470 | השי"ת |  |
| 7:20.471 | לשלמות |  |
| 7:20.472 | ללמוד |  |
| 7:20.473 | וללמד | , |
| 7:20.474 | בדבר |  |
| 7:20.475 | פחות |  |
| 7:20.476 | ומגונה |  |
| 7:20.477 | ובחרפה |  |
| 7:20.478 | הגמורה |  |
| 7:20.479 | שיש |  |
| 7:20.480 | בנו |  |
| 7:20.481 | ונדמה |  |
| 7:20.482 | לגוים |  |
| 7:20.483 | הסכלים |  |
| 7:20.484 | הזונים |  |
| 7:20.485 | בשיריהם |  |
| 7:20.486 | שירי |  |
| 7:20.487 | ענבים | , |
| 7:20.488 | ופחיתות |  |
| 7:20.489 | אשר |  |
| 7:20.490 | יתפארו |  |
| 7:20.491 | בהם |  |
| 7:20.492 | לרוב |  |
| 7:20.493 | סכלות |  |
| 7:20.494 | ופחיתותם |  |
| 7:20.495 | כדאמור |  |
| 7:20.496 | רבנן |  |
| 7:20.497 | תפארת |  |
| 7:20.498 | גוים |  |
| 7:20.499 | עבירה | . |
| 7:20.500 | גוי |  |
| 7:20.501 | מפערה |  |
| 7:20.502 | מעי | ) |
| 7:20.503 | לא |  |
| 7:20.504 | כאלה |  |
| 7:20.505 | חלק |  |
| 7:20.506 | יעקב | , |
| 7:20.507 | שנאמר |  |
| 7:20.508 | בהם |  |
| 7:20.509 | ואתר |  |
| 7:20.510 | תהיו |  |
| 7:20.511 | לי |  |
| 7:20.512 | ממלכת |  |
| 7:20.513 | כהנים |  |
| 7:20.514 | וגוי |  |
| 7:20.515 | קדוש | , |
| 7:20.516 | וכל |  |
| 7:20.517 | מי |  |
| 7:20.518 | שמשתמש |  |
| 7:20.519 | במחשבה |  |
| 7:20.520 | או |  |
| 7:20.521 | בדבריו |  |
| 7:20.522 | מענייני |  |
| 7:20.523 | המשגל |  |
| 7:20.524 | אשר |  |
| 7:20.525 | הוא |  |
| 7:20.526 | חרפה |  |
| 7:20.527 | לנו |  |
| 7:20.528 | ביותר |  |
| 7:20.529 | מן |  |
| 7:20.530 | הצריך |  |
| 7:20.531 | לו |  |
| 7:20.532 | כמאמר |  |
| 7:20.533 | רז"ל |  |
| 7:20.534 | מגיד |  |
| 7:20.535 | לאדם |  |
| 7:20.536 | מה |  |
| 7:20.537 | שיחו |  |
| 7:20.538 | אפילו |  |
| 7:20.539 | שיחה |  |
| 7:20.540 | יתירה |  |
| 7:20.541 | שבין |  |
| 7:20.542 | איש |  |
| 7:20.543 | לאשתו |  |
| 7:20.544 | מגידין |  |
| 7:20.545 | לו | ) |
| 7:20.546 | וכ"ש |  |
| 7:20.547 | אם |  |
| 7:20.548 | יאמר |  |
| 7:20.549 | שירים | , |
| 7:20.550 | הנה |  |
| 7:20.551 | נשתמש |  |
| 7:20.552 | בטובה |  |
| 7:20.553 | אשר |  |
| 7:20.554 | גמלו |  |
| 7:20.555 | השי"צ |  |
| 7:20.556 | בדברי |  |
| 7:20.557 | מרי |  |
| 7:20.558 | וסכלות |  |
| 7:20.559 | גמורה |  |
| 7:20.560 | ועבר |  |
| 7:20.561 | על |  |
| 7:20.562 | מצות |  |
| 7:20.563 | הבורא |  |
| 7:20.564 | ית | '. |
| 7:20.565 | והוא |  |
| 7:20.566 | כמי |  |
| 7:20.567 | שנאמר |  |
| 7:20.568 | בהם |  |
| 7:20.569 | וכסף |  |
| 7:20.570 | הרביתי |  |
| 7:20.571 | להם |  |
| 7:20.572 | וזהב |  |
| 7:20.573 | עשו |  |
| 7:20.574 | לבעל | , |
| 7:20.575 | זהו |  |
| 7:20.576 | תורף |  |
| 7:20.577 | דברי |  |
| 7:20.578 | הר | ' |
| 7:20.579 | המורה |  |
| 7:20.580 | ז"ל | . |
| 7:20.581 | גם |  |
| 7:20.582 | הרס"ג |  |
| 7:20.583 | בס | ' |
| 7:20.584 | האמונות |  |
| 7:20.585 | מאמר |  |
| 7:20.586 | עשירי |  |
| 7:20.587 | פרק |  |
| 7:20.588 | ו | ' |
| 7:20.589 | דבר |  |
| 7:20.590 | בגנות |  |
| 7:20.591 | רבוי |  |
| 7:20.592 | המשגל |  |
| 7:20.593 | ולכלוכו |  |
| 7:20.594 | וטנופו |  |
| 7:20.595 | והחרפה |  |
| 7:20.596 | והבוז |  |
| 7:20.597 | אשר |  |
| 7:20.598 | בו | , |
| 7:20.599 | ושאין |  |
| 7:20.600 | התואה |  |
| 7:20.601 | הזאת |  |
| 7:20.602 | באדם |  |
| 7:20.603 | כי |  |
| 7:20.604 | אם |  |
| 7:20.605 | להקים |  |
| 7:20.606 | בה |  |
| 7:20.607 | זרע |  |
| 7:20.608 | כמצות |  |
| 7:20.609 | השי"ת |  |
| 7:20.610 | ואתם |  |
| 7:20.611 | פרו |  |
| 7:20.612 | ורבו |  |
| 7:20.613 | וכו | ' |
| 7:20.614 | ובהצנע |  |
| 7:20.615 | בעת |  |
| 7:20.616 | הראוי |  |
| 7:20.617 | ויעצור |  |
| 7:20.618 | בה |  |
| 7:20.619 | מהרבות |  |
| 7:20.620 | ע"כ | . |
| 7:20.621 | ובעיני |  |
| 7:20.622 | ראיתי |  |
| 7:20.623 | אחד |  |
| 7:20.624 | מהמון |  |
| 7:20.625 | העם |  |
| 7:20.626 | שאל |  |
| 7:20.627 | לתינוק |  |
| 7:20.628 | קטן |  |
| 7:20.629 | ואמר |  |
| 7:20.630 | האם |  |
| 7:20.631 | אביך |  |
| 7:20.632 | אוהב |  |
| 7:20.633 | את |  |
| 7:20.634 | אמך |  |
| 7:20.635 | ומנשקה | , |
| 7:20.636 | וגער |  |
| 7:20.637 | בו |  |
| 7:20.638 | הדיין |  |
| 7:20.639 | כמה"ר |  |
| 7:20.640 | סלי | ' |
| 7:20.641 | צאלח |  |
| 7:20.642 | ז"ל |  |
| 7:20.643 | ואמר |  |
| 7:20.644 | לו |  |
| 7:20.645 | כלך |  |
| 7:20.646 | מגבלות |  |
| 7:20.647 | הפה |  |
| 7:20.648 | כי |  |
| 7:20.649 | ענשו |  |
| 7:20.650 | חמור | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

