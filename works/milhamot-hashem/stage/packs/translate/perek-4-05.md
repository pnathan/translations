# Translation pack: milhamot-hashem perek-4-05 (Milchamot ha-Shem 4:8–4:9)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-05`.

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
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- והוא שלא: "provided that" (fixed).
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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


## Milchamot ha-Shem 4:8

Hebrew (Wikisource): כה) ולפי דברי הזהר ומקדש מלך למה שינו השבעים ושנים זקנים לתלמי המלך כשהעתיקו לו את התורה כדגרסינן בפרק קמא דמגלה מעשה בתלמי המלך שכינס ע"ב זקינים והושיבם בע"ב בתים ולא גילה להם על מה כינסם. ונכנס אצל כל אחד ואחד מהם. ואמר להם כתבו לי תורת משה רבכם. ונתן הקב"ה עצה בלב כל אחד ואחד והסכימה דעתם לדעת אחת, וכתבו לו אלהים ברא בראשית את השמים ואת הארץ ופי' רש"י ז"ל שלא יאמר בראשית שם הוא, ושתי רשויות הן, האחד ברא את השני, וכן כתבו אעשה אדם בצלם ובדמות. וכן כ' הרב מהרש"א בפי' רש"י שלא יאמר שתי רשויות הן והוא לא יקבל מה שאמרו בסנהדרין כל מקום שפקרו המינים תשובתן בצדן. נעשה אדם, ויברא אלהים, הבה נרדה, וירד ה', ונכתבו מעיקרא בלשון רבים דאמר ר' יוחנן אין הקב"ה עושה דבר עד שנמלך בפימליא שלמעלה וכו' וענינו כמו שכ' רש"י בחומש שהוא ממדת הענוה שיהא הגדול נמלך בקטן והוא על פי בראשית רבה וכו'.

| id | word | punctuation after |
|---|---|---|
| 4:8.1 | כה | ) |
| 4:8.2 | ולפי |  |
| 4:8.3 | דברי |  |
| 4:8.4 | הזהר |  |
| 4:8.5 | ומקדש |  |
| 4:8.6 | מלך |  |
| 4:8.7 | למה |  |
| 4:8.8 | שינו |  |
| 4:8.9 | השבעים |  |
| 4:8.10 | ושנים |  |
| 4:8.11 | זקנים |  |
| 4:8.12 | לתלמי |  |
| 4:8.13 | המלך |  |
| 4:8.14 | כשהעתיקו |  |
| 4:8.15 | לו |  |
| 4:8.16 | את |  |
| 4:8.17 | התורה |  |
| 4:8.18 | כדגרסינן |  |
| 4:8.19 | בפרק |  |
| 4:8.20 | קמא |  |
| 4:8.21 | דמגלה |  |
| 4:8.22 | מעשה |  |
| 4:8.23 | בתלמי |  |
| 4:8.24 | המלך |  |
| 4:8.25 | שכינס |  |
| 4:8.26 | ע"ב |  |
| 4:8.27 | זקינים |  |
| 4:8.28 | והושיבם |  |
| 4:8.29 | בע"ב |  |
| 4:8.30 | בתים |  |
| 4:8.31 | ולא |  |
| 4:8.32 | גילה |  |
| 4:8.33 | להם |  |
| 4:8.34 | על |  |
| 4:8.35 | מה |  |
| 4:8.36 | כינסם | . |
| 4:8.37 | ונכנס |  |
| 4:8.38 | אצל |  |
| 4:8.39 | כל |  |
| 4:8.40 | אחד |  |
| 4:8.41 | ואחד |  |
| 4:8.42 | מהם | . |
| 4:8.43 | ואמר |  |
| 4:8.44 | להם |  |
| 4:8.45 | כתבו |  |
| 4:8.46 | לי |  |
| 4:8.47 | תורת |  |
| 4:8.48 | משה |  |
| 4:8.49 | רבכם | . |
| 4:8.50 | ונתן |  |
| 4:8.51 | הקב"ה |  |
| 4:8.52 | עצה |  |
| 4:8.53 | בלב |  |
| 4:8.54 | כל |  |
| 4:8.55 | אחד |  |
| 4:8.56 | ואחד |  |
| 4:8.57 | והסכימה |  |
| 4:8.58 | דעתם |  |
| 4:8.59 | לדעת |  |
| 4:8.60 | אחת | , |
| 4:8.61 | וכתבו |  |
| 4:8.62 | לו |  |
| 4:8.63 | אלהים |  |
| 4:8.64 | ברא |  |
| 4:8.65 | בראשית |  |
| 4:8.66 | את |  |
| 4:8.67 | השמים |  |
| 4:8.68 | ואת |  |
| 4:8.69 | הארץ |  |
| 4:8.70 | ופי | ' |
| 4:8.71 | רש"י |  |
| 4:8.72 | ז"ל |  |
| 4:8.73 | שלא |  |
| 4:8.74 | יאמר |  |
| 4:8.75 | בראשית |  |
| 4:8.76 | שם |  |
| 4:8.77 | הוא | , |
| 4:8.78 | ושתי |  |
| 4:8.79 | רשויות |  |
| 4:8.80 | הן | , |
| 4:8.81 | האחד |  |
| 4:8.82 | ברא |  |
| 4:8.83 | את |  |
| 4:8.84 | השני | , |
| 4:8.85 | וכן |  |
| 4:8.86 | כתבו |  |
| 4:8.87 | אעשה |  |
| 4:8.88 | אדם |  |
| 4:8.89 | בצלם |  |
| 4:8.90 | ובדמות | . |
| 4:8.91 | וכן |  |
| 4:8.92 | כ | ' |
| 4:8.93 | הרב |  |
| 4:8.94 | מהרש"א |  |
| 4:8.95 | בפי | ' |
| 4:8.96 | רש"י |  |
| 4:8.97 | שלא |  |
| 4:8.98 | יאמר |  |
| 4:8.99 | שתי |  |
| 4:8.100 | רשויות |  |
| 4:8.101 | הן |  |
| 4:8.102 | והוא |  |
| 4:8.103 | לא |  |
| 4:8.104 | יקבל |  |
| 4:8.105 | מה |  |
| 4:8.106 | שאמרו |  |
| 4:8.107 | בסנהדרין |  |
| 4:8.108 | כל |  |
| 4:8.109 | מקום |  |
| 4:8.110 | שפקרו |  |
| 4:8.111 | המינים |  |
| 4:8.112 | תשובתן |  |
| 4:8.113 | בצדן | . |
| 4:8.114 | נעשה |  |
| 4:8.115 | אדם | , |
| 4:8.116 | ויברא |  |
| 4:8.117 | אלהים | , |
| 4:8.118 | הבה |  |
| 4:8.119 | נרדה | , |
| 4:8.120 | וירד |  |
| 4:8.121 | ה | ', |
| 4:8.122 | ונכתבו |  |
| 4:8.123 | מעיקרא |  |
| 4:8.124 | בלשון |  |
| 4:8.125 | רבים |  |
| 4:8.126 | דאמר |  |
| 4:8.127 | ר | ' |
| 4:8.128 | יוחנן |  |
| 4:8.129 | אין |  |
| 4:8.130 | הקב"ה |  |
| 4:8.131 | עושה |  |
| 4:8.132 | דבר |  |
| 4:8.133 | עד |  |
| 4:8.134 | שנמלך |  |
| 4:8.135 | בפימליא |  |
| 4:8.136 | שלמעלה |  |
| 4:8.137 | וכו | ' |
| 4:8.138 | וענינו |  |
| 4:8.139 | כמו |  |
| 4:8.140 | שכ | ' |
| 4:8.141 | רש"י |  |
| 4:8.142 | בחומש |  |
| 4:8.143 | שהוא |  |
| 4:8.144 | ממדת |  |
| 4:8.145 | הענוה |  |
| 4:8.146 | שיהא |  |
| 4:8.147 | הגדול |  |
| 4:8.148 | נמלך |  |
| 4:8.149 | בקטן |  |
| 4:8.150 | והוא |  |
| 4:8.151 | על |  |
| 4:8.152 | פי |  |
| 4:8.153 | בראשית |  |
| 4:8.154 | רבה |  |
| 4:8.155 | וכו | '. |


## Milchamot ha-Shem 4:9

Hebrew (Wikisource): ועיין במורה נבוכים פ"ו משני שהביא שתי נוסחאות. האחת שאין הקב"ה עושה דבר עד שמסתכל בפימליא שלמעלה והוה כמאמר אפלטון שהשי"ת יעיין בעולם השכלים וישפיע ממנו המציאות. והנוסחא השנית היא שאמרו במקום אחר אין הקב"ה עושה דבר עד שנמלך בפימיליא שלמעלה. ופימליא הוא החיל בלשון יון. ובבראשית רבה ובמדרש קהלת את אשר כבר עשוהו. עשהו לא נאמר אלא עשוהו. כביכול הוא ובית דינו נמנו על כל אבר ואבר שבך והושיבוהו על כנו שנ' הוא עשך ויכוננך, ובבראשית רבה אמרו ג"כ כל מקום שנ' וה' הוא ובית דנו, ואין הכוונה כאלו המאמרים מה שהשבוהו הפתאים שיש לו ית' דברים או מחשבה או שאלת עצה או התבוננות, והעזר בדעת אחרים, ואיך יעזר הבורא במה שברא, אבל זה כולו עניינו שאפי' חלקי המציאות עד בריאת האברים מבעלי חיים כפי מה שהם עליו. כל זה באמצעית מלאכים, כי הכחות כולם מלאכים. ומה מאד רע עורון הסכלות, ומה מאד מזיק, שאלו תאמר לאחד מאשר יחשבו שהם חכמי ישראל וכו' ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:9.1 | ועיין |  |
| 4:9.2 | במורה |  |
| 4:9.3 | נבוכים |  |
| 4:9.4 | פ"ו |  |
| 4:9.5 | משני |  |
| 4:9.6 | שהביא |  |
| 4:9.7 | שתי |  |
| 4:9.8 | נוסחאות | . |
| 4:9.9 | האחת |  |
| 4:9.10 | שאין |  |
| 4:9.11 | הקב"ה |  |
| 4:9.12 | עושה |  |
| 4:9.13 | דבר |  |
| 4:9.14 | עד |  |
| 4:9.15 | שמסתכל |  |
| 4:9.16 | בפימליא |  |
| 4:9.17 | שלמעלה |  |
| 4:9.18 | והוה |  |
| 4:9.19 | כמאמר |  |
| 4:9.20 | אפלטון |  |
| 4:9.21 | שהשי"ת |  |
| 4:9.22 | יעיין |  |
| 4:9.23 | בעולם |  |
| 4:9.24 | השכלים |  |
| 4:9.25 | וישפיע |  |
| 4:9.26 | ממנו |  |
| 4:9.27 | המציאות | . |
| 4:9.28 | והנוסחא |  |
| 4:9.29 | השנית |  |
| 4:9.30 | היא |  |
| 4:9.31 | שאמרו |  |
| 4:9.32 | במקום |  |
| 4:9.33 | אחר |  |
| 4:9.34 | אין |  |
| 4:9.35 | הקב"ה |  |
| 4:9.36 | עושה |  |
| 4:9.37 | דבר |  |
| 4:9.38 | עד |  |
| 4:9.39 | שנמלך |  |
| 4:9.40 | בפימיליא |  |
| 4:9.41 | שלמעלה | . |
| 4:9.42 | ופימליא |  |
| 4:9.43 | הוא |  |
| 4:9.44 | החיל |  |
| 4:9.45 | בלשון |  |
| 4:9.46 | יון | . |
| 4:9.47 | ובבראשית |  |
| 4:9.48 | רבה |  |
| 4:9.49 | ובמדרש |  |
| 4:9.50 | קהלת |  |
| 4:9.51 | את |  |
| 4:9.52 | אשר |  |
| 4:9.53 | כבר |  |
| 4:9.54 | עשוהו | . |
| 4:9.55 | עשהו |  |
| 4:9.56 | לא |  |
| 4:9.57 | נאמר |  |
| 4:9.58 | אלא |  |
| 4:9.59 | עשוהו | . |
| 4:9.60 | כביכול |  |
| 4:9.61 | הוא |  |
| 4:9.62 | ובית |  |
| 4:9.63 | דינו |  |
| 4:9.64 | נמנו |  |
| 4:9.65 | על |  |
| 4:9.66 | כל |  |
| 4:9.67 | אבר |  |
| 4:9.68 | ואבר |  |
| 4:9.69 | שבך |  |
| 4:9.70 | והושיבוהו |  |
| 4:9.71 | על |  |
| 4:9.72 | כנו |  |
| 4:9.73 | שנ | ' |
| 4:9.74 | הוא |  |
| 4:9.75 | עשך |  |
| 4:9.76 | ויכוננך | , |
| 4:9.77 | ובבראשית |  |
| 4:9.78 | רבה |  |
| 4:9.79 | אמרו |  |
| 4:9.80 | ג"כ |  |
| 4:9.81 | כל |  |
| 4:9.82 | מקום |  |
| 4:9.83 | שנ | ' |
| 4:9.84 | וה | ' |
| 4:9.85 | הוא |  |
| 4:9.86 | ובית |  |
| 4:9.87 | דנו | , |
| 4:9.88 | ואין |  |
| 4:9.89 | הכוונה |  |
| 4:9.90 | כאלו |  |
| 4:9.91 | המאמרים |  |
| 4:9.92 | מה |  |
| 4:9.93 | שהשבוהו |  |
| 4:9.94 | הפתאים |  |
| 4:9.95 | שיש |  |
| 4:9.96 | לו |  |
| 4:9.97 | ית | ' |
| 4:9.98 | דברים |  |
| 4:9.99 | או |  |
| 4:9.100 | מחשבה |  |
| 4:9.101 | או |  |
| 4:9.102 | שאלת |  |
| 4:9.103 | עצה |  |
| 4:9.104 | או |  |
| 4:9.105 | התבוננות | , |
| 4:9.106 | והעזר |  |
| 4:9.107 | בדעת |  |
| 4:9.108 | אחרים | , |
| 4:9.109 | ואיך |  |
| 4:9.110 | יעזר |  |
| 4:9.111 | הבורא |  |
| 4:9.112 | במה |  |
| 4:9.113 | שברא | , |
| 4:9.114 | אבל |  |
| 4:9.115 | זה |  |
| 4:9.116 | כולו |  |
| 4:9.117 | עניינו |  |
| 4:9.118 | שאפי | ' |
| 4:9.119 | חלקי |  |
| 4:9.120 | המציאות |  |
| 4:9.121 | עד |  |
| 4:9.122 | בריאת |  |
| 4:9.123 | האברים |  |
| 4:9.124 | מבעלי |  |
| 4:9.125 | חיים |  |
| 4:9.126 | כפי |  |
| 4:9.127 | מה |  |
| 4:9.128 | שהם |  |
| 4:9.129 | עליו | . |
| 4:9.130 | כל |  |
| 4:9.131 | זה |  |
| 4:9.132 | באמצעית |  |
| 4:9.133 | מלאכים | , |
| 4:9.134 | כי |  |
| 4:9.135 | הכחות |  |
| 4:9.136 | כולם |  |
| 4:9.137 | מלאכים | . |
| 4:9.138 | ומה |  |
| 4:9.139 | מאד |  |
| 4:9.140 | רע |  |
| 4:9.141 | עורון |  |
| 4:9.142 | הסכלות | , |
| 4:9.143 | ומה |  |
| 4:9.144 | מאד |  |
| 4:9.145 | מזיק | , |
| 4:9.146 | שאלו |  |
| 4:9.147 | תאמר |  |
| 4:9.148 | לאחד |  |
| 4:9.149 | מאשר |  |
| 4:9.150 | יחשבו |  |
| 4:9.151 | שהם |  |
| 4:9.152 | חכמי |  |
| 4:9.153 | ישראל |  |
| 4:9.154 | וכו | ' |
| 4:9.155 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

