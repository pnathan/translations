# Review pack: milhamot-hashem perek-1-01 (Milchamot ha-Shem 1:1–1:2)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-1-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-01`.

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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- משכילים / השכלה: "the educated" / "enlightenment" in the sense of the
  Haskalah only where the paragraph speaks of that movement; otherwise
  "those of understanding".
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- מרגל לרגל: "from festival to festival" (Num 20 ff. sense), not "foot".
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- דברי תימה: "astonishing words" (words that astonish).
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


## Milchamot ha-Shem 1:1

Hebrew (Wikisource): הקדמת המחבר.

| id | word | punctuation after |
|---|---|---|
| 1:1.1 | הקדמת |  |
| 1:1.2 | המחבר | . |


### Translation under review (1:1)

English: The Author's Introduction.

Literal: introduction-of the-author

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:1.1 | introduction of |  |  | Construct state: 'the introduction of the author'. |
| 1:1.2 | the author |  |  |  |

Phrases:
- 1:1.1 1:1.2 | The Author's Introduction. | introduction-of the-author

Grammar notes:
- הקדמת המחבר is a construct phrase, 'the introduction of the author', and it is the heading of the section.
- The Hebrew heading ends with a full stop, which the English keeps.

Translator's note: Heading of the author's introduction, as the decisions give it.

Commentary rulings:

## Milchamot ha-Shem 1:2

Hebrew (Wikisource): יתבונן המשכיל ויתאונן הירא והחרד אל דבר ה', על מצב עם בני ישראל בראותו אשר כל עם ועם מכל מפלגות הגוים מחבבים את לשונם ומגדילים את שפתם אשר חנכו עליה. ואשר הוחקה בה אמונתם ודתם איש ללשונו ואנחנו עם בני ישראל מעת אשר גלינו מארצנו עזבנו את לשונינו הקדושה אשר דבר בה אדם הראשון ובה ניתנה תורתינו הקדושה בהגלות כבודו ית' על הר סיני ואמר לעמו אנכי ה' אלהיך וכו' ולא עשינו כמצווה עלינו ולמדתם אותם את בניכם לדבר בם. וכמו שפי' רש״י ז״ל בפרשת עקב שמכאן אמרו כשהתינוק מתחיל לדבר אביו משיח עמו בלשון הקדש ומלמדו תורה עכ״ל, רק כל שיחתנו עם בנינו בלשון העם אשר אנו יושבים בקרבם ומלמדי תינוקות גם הם רובם לא יכינו לדבר כל צרכם בלשון הקדש אף שיודעים אותו כי לא נסו באלה ואין סדרים ללמוד תינוקות של בית רבן. וידוע הוא ומפורסם חמלת האב על בניו וכמה טרחות ויגיעות ייגעו ההורים בגידול בניהם. עם רוב השקפתם על כל צרכי ילדיהם במאכלו ומשקיו הניאותים לו ומצעו ומלבושו! אך כאשר הגיע חתור להכניסו לבית הספר ישליכהו אצל מלמד תינוקות בבתים צרים וחשוכים על שיירי מחצלאות על הארץ בלי נקיון כי לא ינקה המלמד את מעון התינוקות רק מרגל לרגל אחר. וכל היום יושבים צפופים בלי סדרים ושואפים אויר סרוח ומעופש. ושותים מים מלוחים ובגלל זה כאשר יעברו איזה חדשים מעת שהובא אל בית הספר יעלו בראשו צמחים ושחין רע. ואבותיהם לא ישגיחו עליהם להביא להם מים צלולים ומתוקים לשתות ולא מצע נקי לישב עליו וברוב הימים יושבים על האדמה מאין מחצלת! ובתקופת העת החדשה משנת התר״ן לפרט גדול ליצירה התעוררו רבים מהמשכילים ויבקשו הסתדרות בתי הספר והסתדרות ילדיו לתלמוד תורה ולימוד השפות העברית והערבית והתורקית המתהלכות! וחתרו להושיב מלמדים משכילים בתורת ה' לנהל את ילדי בית הספר בחנוך מוסרי ובנקיון מעון התלמוד בבתים מאירים ובחלונות מרווחים כדי לשאוף רוח צח ונקי! וכל יגיעם של המשכילים ההם היה לריק ולא יצאה מחשבתם לפועל ותינוקות של בית רבן נשארו כשהיו אצל מלמד בחדרים חשוכים ומעופשים יושבים ברבוביא זה לומד אלפא ביתא וזה לומד מקרא וזה לומד עם תרגום אונקלוס ותרגום רס״ג בערבי בלי הבין. גם קטני קטנים שאין ראוים ללמוד כלל כבן שתים ושלש שנים מביאים שם. ופעמים רבות התינוק נפנה שם. והמלמד שולח לאמו שתבוא ותקנח צואת בנה! ולקול הקורא במכתבי עתים בימים ההם מאת ה' אליעזר בן יהודה לתחית השפה העברית הקדושה נכספה נפשינו אך ורק בדברים ושיחות בעלמא ולא באה לידי התגשמות עד שנת תר״ע, אשר אז הואילה ממשלתינו הרוממה ממשלת תוגרמה לפתוח בתי ספר לילדי בני ישראל, ונעשה שם הסתדרות יפה לילדי בני ישראל בעליות מרווחים וחלונות פתוחים ומצעות של צמר, ויריעות עזים וגם שולחנות לילדים גדולים ללמוד ולכתוב עליהם! ואז אמרנו הנה בא העת להוציא לאור עשתונות איש ירושלים ת״ו, אשר בטרם תחיל ילדה האדם הגדול בתחית השפה העברית הוא ניהו אליעזר בן יהודה, ואפריין נמטיה על אשר פעל ועשה ובכל כחו עמל לטובת עם ישראל וישוב ארץ ישראל ותחית השפה העברית ועבודת הקדש הלזו עבד בזריזות וחריצות נפלאה ועשה והצליח להקימה מעפר, ואחרי בלותה היתה לה עדנה, ובגלל אהבתו לשפה הקדושה נתפשטו קמטיה, ותואר פניה הציץ ופרח, ותמצא חן וכבוד בעיני שוחריה יודעיה ומכיריה ויעורו משנתם להחלץ חושים להרחיבה ולהעשירה כבימי נעוריה לפני גלות עם ישראל! ובגלל מצוה זו תהי נפשו צרורה בצרור החיים את ה' אלהינו ומפרי עמלו תשבע נפשו בעדן גן אלהים!

| id | word | punctuation after |
|---|---|---|
| 1:2.1 | יתבונן |  |
| 1:2.2 | המשכיל |  |
| 1:2.3 | ויתאונן |  |
| 1:2.4 | הירא |  |
| 1:2.5 | והחרד |  |
| 1:2.6 | אל |  |
| 1:2.7 | דבר |  |
| 1:2.8 | ה | ', |
| 1:2.9 | על |  |
| 1:2.10 | מצב |  |
| 1:2.11 | עם |  |
| 1:2.12 | בני |  |
| 1:2.13 | ישראל |  |
| 1:2.14 | בראותו |  |
| 1:2.15 | אשר |  |
| 1:2.16 | כל |  |
| 1:2.17 | עם |  |
| 1:2.18 | ועם |  |
| 1:2.19 | מכל |  |
| 1:2.20 | מפלגות |  |
| 1:2.21 | הגוים |  |
| 1:2.22 | מחבבים |  |
| 1:2.23 | את |  |
| 1:2.24 | לשונם |  |
| 1:2.25 | ומגדילים |  |
| 1:2.26 | את |  |
| 1:2.27 | שפתם |  |
| 1:2.28 | אשר |  |
| 1:2.29 | חנכו |  |
| 1:2.30 | עליה | . |
| 1:2.31 | ואשר |  |
| 1:2.32 | הוחקה |  |
| 1:2.33 | בה |  |
| 1:2.34 | אמונתם |  |
| 1:2.35 | ודתם |  |
| 1:2.36 | איש |  |
| 1:2.37 | ללשונו |  |
| 1:2.38 | ואנחנו |  |
| 1:2.39 | עם |  |
| 1:2.40 | בני |  |
| 1:2.41 | ישראל |  |
| 1:2.42 | מעת |  |
| 1:2.43 | אשר |  |
| 1:2.44 | גלינו |  |
| 1:2.45 | מארצנו |  |
| 1:2.46 | עזבנו |  |
| 1:2.47 | את |  |
| 1:2.48 | לשונינו |  |
| 1:2.49 | הקדושה |  |
| 1:2.50 | אשר |  |
| 1:2.51 | דבר |  |
| 1:2.52 | בה |  |
| 1:2.53 | אדם |  |
| 1:2.54 | הראשון |  |
| 1:2.55 | ובה |  |
| 1:2.56 | ניתנה |  |
| 1:2.57 | תורתינו |  |
| 1:2.58 | הקדושה |  |
| 1:2.59 | בהגלות |  |
| 1:2.60 | כבודו |  |
| 1:2.61 | ית | ' |
| 1:2.62 | על |  |
| 1:2.63 | הר |  |
| 1:2.64 | סיני |  |
| 1:2.65 | ואמר |  |
| 1:2.66 | לעמו |  |
| 1:2.67 | אנכי |  |
| 1:2.68 | ה | ' |
| 1:2.69 | אלהיך |  |
| 1:2.70 | וכו | ' |
| 1:2.71 | ולא |  |
| 1:2.72 | עשינו |  |
| 1:2.73 | כמצווה |  |
| 1:2.74 | עלינו |  |
| 1:2.75 | ולמדתם |  |
| 1:2.76 | אותם |  |
| 1:2.77 | את |  |
| 1:2.78 | בניכם |  |
| 1:2.79 | לדבר |  |
| 1:2.80 | בם | . |
| 1:2.81 | וכמו |  |
| 1:2.82 | שפי | ' |
| 1:2.83 | רש״י |  |
| 1:2.84 | ז״ל |  |
| 1:2.85 | בפרשת |  |
| 1:2.86 | עקב |  |
| 1:2.87 | שמכאן |  |
| 1:2.88 | אמרו |  |
| 1:2.89 | כשהתינוק |  |
| 1:2.90 | מתחיל |  |
| 1:2.91 | לדבר |  |
| 1:2.92 | אביו |  |
| 1:2.93 | משיח |  |
| 1:2.94 | עמו |  |
| 1:2.95 | בלשון |  |
| 1:2.96 | הקדש |  |
| 1:2.97 | ומלמדו |  |
| 1:2.98 | תורה |  |
| 1:2.99 | עכ״ל | , |
| 1:2.100 | רק |  |
| 1:2.101 | כל |  |
| 1:2.102 | שיחתנו |  |
| 1:2.103 | עם |  |
| 1:2.104 | בנינו |  |
| 1:2.105 | בלשון |  |
| 1:2.106 | העם |  |
| 1:2.107 | אשר |  |
| 1:2.108 | אנו |  |
| 1:2.109 | יושבים |  |
| 1:2.110 | בקרבם |  |
| 1:2.111 | ומלמדי |  |
| 1:2.112 | תינוקות |  |
| 1:2.113 | גם |  |
| 1:2.114 | הם |  |
| 1:2.115 | רובם |  |
| 1:2.116 | לא |  |
| 1:2.117 | יכינו |  |
| 1:2.118 | לדבר |  |
| 1:2.119 | כל |  |
| 1:2.120 | צרכם |  |
| 1:2.121 | בלשון |  |
| 1:2.122 | הקדש |  |
| 1:2.123 | אף |  |
| 1:2.124 | שיודעים |  |
| 1:2.125 | אותו |  |
| 1:2.126 | כי |  |
| 1:2.127 | לא |  |
| 1:2.128 | נסו |  |
| 1:2.129 | באלה |  |
| 1:2.130 | ואין |  |
| 1:2.131 | סדרים |  |
| 1:2.132 | ללמוד |  |
| 1:2.133 | תינוקות |  |
| 1:2.134 | של |  |
| 1:2.135 | בית |  |
| 1:2.136 | רבן | . |
| 1:2.137 | וידוע |  |
| 1:2.138 | הוא |  |
| 1:2.139 | ומפורסם |  |
| 1:2.140 | חמלת |  |
| 1:2.141 | האב |  |
| 1:2.142 | על |  |
| 1:2.143 | בניו |  |
| 1:2.144 | וכמה |  |
| 1:2.145 | טרחות |  |
| 1:2.146 | ויגיעות |  |
| 1:2.147 | ייגעו |  |
| 1:2.148 | ההורים |  |
| 1:2.149 | בגידול |  |
| 1:2.150 | בניהם | . |
| 1:2.151 | עם |  |
| 1:2.152 | רוב |  |
| 1:2.153 | השקפתם |  |
| 1:2.154 | על |  |
| 1:2.155 | כל |  |
| 1:2.156 | צרכי |  |
| 1:2.157 | ילדיהם |  |
| 1:2.158 | במאכלו |  |
| 1:2.159 | ומשקיו |  |
| 1:2.160 | הניאותים |  |
| 1:2.161 | לו |  |
| 1:2.162 | ומצעו |  |
| 1:2.163 | ומלבושו | ! |
| 1:2.164 | אך |  |
| 1:2.165 | כאשר |  |
| 1:2.166 | הגיע |  |
| 1:2.167 | חתור |  |
| 1:2.168 | להכניסו |  |
| 1:2.169 | לבית |  |
| 1:2.170 | הספר |  |
| 1:2.171 | ישליכהו |  |
| 1:2.172 | אצל |  |
| 1:2.173 | מלמד |  |
| 1:2.174 | תינוקות |  |
| 1:2.175 | בבתים |  |
| 1:2.176 | צרים |  |
| 1:2.177 | וחשוכים |  |
| 1:2.178 | על |  |
| 1:2.179 | שיירי |  |
| 1:2.180 | מחצלאות |  |
| 1:2.181 | על |  |
| 1:2.182 | הארץ |  |
| 1:2.183 | בלי |  |
| 1:2.184 | נקיון |  |
| 1:2.185 | כי |  |
| 1:2.186 | לא |  |
| 1:2.187 | ינקה |  |
| 1:2.188 | המלמד |  |
| 1:2.189 | את |  |
| 1:2.190 | מעון |  |
| 1:2.191 | התינוקות |  |
| 1:2.192 | רק |  |
| 1:2.193 | מרגל |  |
| 1:2.194 | לרגל |  |
| 1:2.195 | אחר | . |
| 1:2.196 | וכל |  |
| 1:2.197 | היום |  |
| 1:2.198 | יושבים |  |
| 1:2.199 | צפופים |  |
| 1:2.200 | בלי |  |
| 1:2.201 | סדרים |  |
| 1:2.202 | ושואפים |  |
| 1:2.203 | אויר |  |
| 1:2.204 | סרוח |  |
| 1:2.205 | ומעופש | . |
| 1:2.206 | ושותים |  |
| 1:2.207 | מים |  |
| 1:2.208 | מלוחים |  |
| 1:2.209 | ובגלל |  |
| 1:2.210 | זה |  |
| 1:2.211 | כאשר |  |
| 1:2.212 | יעברו |  |
| 1:2.213 | איזה |  |
| 1:2.214 | חדשים |  |
| 1:2.215 | מעת |  |
| 1:2.216 | שהובא |  |
| 1:2.217 | אל |  |
| 1:2.218 | בית |  |
| 1:2.219 | הספר |  |
| 1:2.220 | יעלו |  |
| 1:2.221 | בראשו |  |
| 1:2.222 | צמחים |  |
| 1:2.223 | ושחין |  |
| 1:2.224 | רע | . |
| 1:2.225 | ואבותיהם |  |
| 1:2.226 | לא |  |
| 1:2.227 | ישגיחו |  |
| 1:2.228 | עליהם |  |
| 1:2.229 | להביא |  |
| 1:2.230 | להם |  |
| 1:2.231 | מים |  |
| 1:2.232 | צלולים |  |
| 1:2.233 | ומתוקים |  |
| 1:2.234 | לשתות |  |
| 1:2.235 | ולא |  |
| 1:2.236 | מצע |  |
| 1:2.237 | נקי |  |
| 1:2.238 | לישב |  |
| 1:2.239 | עליו |  |
| 1:2.240 | וברוב |  |
| 1:2.241 | הימים |  |
| 1:2.242 | יושבים |  |
| 1:2.243 | על |  |
| 1:2.244 | האדמה |  |
| 1:2.245 | מאין |  |
| 1:2.246 | מחצלת | ! |
| 1:2.247 | ובתקופת |  |
| 1:2.248 | העת |  |
| 1:2.249 | החדשה |  |
| 1:2.250 | משנת |  |
| 1:2.251 | התר״ן |  |
| 1:2.252 | לפרט |  |
| 1:2.253 | גדול |  |
| 1:2.254 | ליצירה |  |
| 1:2.255 | התעוררו |  |
| 1:2.256 | רבים |  |
| 1:2.257 | מהמשכילים |  |
| 1:2.258 | ויבקשו |  |
| 1:2.259 | הסתדרות |  |
| 1:2.260 | בתי |  |
| 1:2.261 | הספר |  |
| 1:2.262 | והסתדרות |  |
| 1:2.263 | ילדיו |  |
| 1:2.264 | לתלמוד |  |
| 1:2.265 | תורה |  |
| 1:2.266 | ולימוד |  |
| 1:2.267 | השפות |  |
| 1:2.268 | העברית |  |
| 1:2.269 | והערבית |  |
| 1:2.270 | והתורקית |  |
| 1:2.271 | המתהלכות | ! |
| 1:2.272 | וחתרו |  |
| 1:2.273 | להושיב |  |
| 1:2.274 | מלמדים |  |
| 1:2.275 | משכילים |  |
| 1:2.276 | בתורת |  |
| 1:2.277 | ה | ' |
| 1:2.278 | לנהל |  |
| 1:2.279 | את |  |
| 1:2.280 | ילדי |  |
| 1:2.281 | בית |  |
| 1:2.282 | הספר |  |
| 1:2.283 | בחנוך |  |
| 1:2.284 | מוסרי |  |
| 1:2.285 | ובנקיון |  |
| 1:2.286 | מעון |  |
| 1:2.287 | התלמוד |  |
| 1:2.288 | בבתים |  |
| 1:2.289 | מאירים |  |
| 1:2.290 | ובחלונות |  |
| 1:2.291 | מרווחים |  |
| 1:2.292 | כדי |  |
| 1:2.293 | לשאוף |  |
| 1:2.294 | רוח |  |
| 1:2.295 | צח |  |
| 1:2.296 | ונקי | ! |
| 1:2.297 | וכל |  |
| 1:2.298 | יגיעם |  |
| 1:2.299 | של |  |
| 1:2.300 | המשכילים |  |
| 1:2.301 | ההם |  |
| 1:2.302 | היה |  |
| 1:2.303 | לריק |  |
| 1:2.304 | ולא |  |
| 1:2.305 | יצאה |  |
| 1:2.306 | מחשבתם |  |
| 1:2.307 | לפועל |  |
| 1:2.308 | ותינוקות |  |
| 1:2.309 | של |  |
| 1:2.310 | בית |  |
| 1:2.311 | רבן |  |
| 1:2.312 | נשארו |  |
| 1:2.313 | כשהיו |  |
| 1:2.314 | אצל |  |
| 1:2.315 | מלמד |  |
| 1:2.316 | בחדרים |  |
| 1:2.317 | חשוכים |  |
| 1:2.318 | ומעופשים |  |
| 1:2.319 | יושבים |  |
| 1:2.320 | ברבוביא |  |
| 1:2.321 | זה |  |
| 1:2.322 | לומד |  |
| 1:2.323 | אלפא |  |
| 1:2.324 | ביתא |  |
| 1:2.325 | וזה |  |
| 1:2.326 | לומד |  |
| 1:2.327 | מקרא |  |
| 1:2.328 | וזה |  |
| 1:2.329 | לומד |  |
| 1:2.330 | עם |  |
| 1:2.331 | תרגום |  |
| 1:2.332 | אונקלוס |  |
| 1:2.333 | ותרגום |  |
| 1:2.334 | רס״ג |  |
| 1:2.335 | בערבי |  |
| 1:2.336 | בלי |  |
| 1:2.337 | הבין | . |
| 1:2.338 | גם |  |
| 1:2.339 | קטני |  |
| 1:2.340 | קטנים |  |
| 1:2.341 | שאין |  |
| 1:2.342 | ראוים |  |
| 1:2.343 | ללמוד |  |
| 1:2.344 | כלל |  |
| 1:2.345 | כבן |  |
| 1:2.346 | שתים |  |
| 1:2.347 | ושלש |  |
| 1:2.348 | שנים |  |
| 1:2.349 | מביאים |  |
| 1:2.350 | שם | . |
| 1:2.351 | ופעמים |  |
| 1:2.352 | רבות |  |
| 1:2.353 | התינוק |  |
| 1:2.354 | נפנה |  |
| 1:2.355 | שם | . |
| 1:2.356 | והמלמד |  |
| 1:2.357 | שולח |  |
| 1:2.358 | לאמו |  |
| 1:2.359 | שתבוא |  |
| 1:2.360 | ותקנח |  |
| 1:2.361 | צואת |  |
| 1:2.362 | בנה | ! |
| 1:2.363 | ולקול |  |
| 1:2.364 | הקורא |  |
| 1:2.365 | במכתבי |  |
| 1:2.366 | עתים |  |
| 1:2.367 | בימים |  |
| 1:2.368 | ההם |  |
| 1:2.369 | מאת |  |
| 1:2.370 | ה | ' |
| 1:2.371 | אליעזר |  |
| 1:2.372 | בן |  |
| 1:2.373 | יהודה |  |
| 1:2.374 | לתחית |  |
| 1:2.375 | השפה |  |
| 1:2.376 | העברית |  |
| 1:2.377 | הקדושה |  |
| 1:2.378 | נכספה |  |
| 1:2.379 | נפשינו |  |
| 1:2.380 | אך |  |
| 1:2.381 | ורק |  |
| 1:2.382 | בדברים |  |
| 1:2.383 | ושיחות |  |
| 1:2.384 | בעלמא |  |
| 1:2.385 | ולא |  |
| 1:2.386 | באה |  |
| 1:2.387 | לידי |  |
| 1:2.388 | התגשמות |  |
| 1:2.389 | עד |  |
| 1:2.390 | שנת |  |
| 1:2.391 | תר״ע | , |
| 1:2.392 | אשר |  |
| 1:2.393 | אז |  |
| 1:2.394 | הואילה |  |
| 1:2.395 | ממשלתינו |  |
| 1:2.396 | הרוממה |  |
| 1:2.397 | ממשלת |  |
| 1:2.398 | תוגרמה |  |
| 1:2.399 | לפתוח |  |
| 1:2.400 | בתי |  |
| 1:2.401 | ספר |  |
| 1:2.402 | לילדי |  |
| 1:2.403 | בני |  |
| 1:2.404 | ישראל | , |
| 1:2.405 | ונעשה |  |
| 1:2.406 | שם |  |
| 1:2.407 | הסתדרות |  |
| 1:2.408 | יפה |  |
| 1:2.409 | לילדי |  |
| 1:2.410 | בני |  |
| 1:2.411 | ישראל |  |
| 1:2.412 | בעליות |  |
| 1:2.413 | מרווחים |  |
| 1:2.414 | וחלונות |  |
| 1:2.415 | פתוחים |  |
| 1:2.416 | ומצעות |  |
| 1:2.417 | של |  |
| 1:2.418 | צמר | , |
| 1:2.419 | ויריעות |  |
| 1:2.420 | עזים |  |
| 1:2.421 | וגם |  |
| 1:2.422 | שולחנות |  |
| 1:2.423 | לילדים |  |
| 1:2.424 | גדולים |  |
| 1:2.425 | ללמוד |  |
| 1:2.426 | ולכתוב |  |
| 1:2.427 | עליהם | ! |
| 1:2.428 | ואז |  |
| 1:2.429 | אמרנו |  |
| 1:2.430 | הנה |  |
| 1:2.431 | בא |  |
| 1:2.432 | העת |  |
| 1:2.433 | להוציא |  |
| 1:2.434 | לאור |  |
| 1:2.435 | עשתונות |  |
| 1:2.436 | איש |  |
| 1:2.437 | ירושלים |  |
| 1:2.438 | ת״ו | , |
| 1:2.439 | אשר |  |
| 1:2.440 | בטרם |  |
| 1:2.441 | תחיל |  |
| 1:2.442 | ילדה |  |
| 1:2.443 | האדם |  |
| 1:2.444 | הגדול |  |
| 1:2.445 | בתחית |  |
| 1:2.446 | השפה |  |
| 1:2.447 | העברית |  |
| 1:2.448 | הוא |  |
| 1:2.449 | ניהו |  |
| 1:2.450 | אליעזר |  |
| 1:2.451 | בן |  |
| 1:2.452 | יהודה | , |
| 1:2.453 | ואפריין |  |
| 1:2.454 | נמטיה |  |
| 1:2.455 | על |  |
| 1:2.456 | אשר |  |
| 1:2.457 | פעל |  |
| 1:2.458 | ועשה |  |
| 1:2.459 | ובכל |  |
| 1:2.460 | כחו |  |
| 1:2.461 | עמל |  |
| 1:2.462 | לטובת |  |
| 1:2.463 | עם |  |
| 1:2.464 | ישראל |  |
| 1:2.465 | וישוב |  |
| 1:2.466 | ארץ |  |
| 1:2.467 | ישראל |  |
| 1:2.468 | ותחית |  |
| 1:2.469 | השפה |  |
| 1:2.470 | העברית |  |
| 1:2.471 | ועבודת |  |
| 1:2.472 | הקדש |  |
| 1:2.473 | הלזו |  |
| 1:2.474 | עבד |  |
| 1:2.475 | בזריזות |  |
| 1:2.476 | וחריצות |  |
| 1:2.477 | נפלאה |  |
| 1:2.478 | ועשה |  |
| 1:2.479 | והצליח |  |
| 1:2.480 | להקימה |  |
| 1:2.481 | מעפר | , |
| 1:2.482 | ואחרי |  |
| 1:2.483 | בלותה |  |
| 1:2.484 | היתה |  |
| 1:2.485 | לה |  |
| 1:2.486 | עדנה | , |
| 1:2.487 | ובגלל |  |
| 1:2.488 | אהבתו |  |
| 1:2.489 | לשפה |  |
| 1:2.490 | הקדושה |  |
| 1:2.491 | נתפשטו |  |
| 1:2.492 | קמטיה | , |
| 1:2.493 | ותואר |  |
| 1:2.494 | פניה |  |
| 1:2.495 | הציץ |  |
| 1:2.496 | ופרח | , |
| 1:2.497 | ותמצא |  |
| 1:2.498 | חן |  |
| 1:2.499 | וכבוד |  |
| 1:2.500 | בעיני |  |
| 1:2.501 | שוחריה |  |
| 1:2.502 | יודעיה |  |
| 1:2.503 | ומכיריה |  |
| 1:2.504 | ויעורו |  |
| 1:2.505 | משנתם |  |
| 1:2.506 | להחלץ |  |
| 1:2.507 | חושים |  |
| 1:2.508 | להרחיבה |  |
| 1:2.509 | ולהעשירה |  |
| 1:2.510 | כבימי |  |
| 1:2.511 | נעוריה |  |
| 1:2.512 | לפני |  |
| 1:2.513 | גלות |  |
| 1:2.514 | עם |  |
| 1:2.515 | ישראל | ! |
| 1:2.516 | ובגלל |  |
| 1:2.517 | מצוה |  |
| 1:2.518 | זו |  |
| 1:2.519 | תהי |  |
| 1:2.520 | נפשו |  |
| 1:2.521 | צרורה |  |
| 1:2.522 | בצרור |  |
| 1:2.523 | החיים |  |
| 1:2.524 | את |  |
| 1:2.525 | ה | ' |
| 1:2.526 | אלהינו |  |
| 1:2.527 | ומפרי |  |
| 1:2.528 | עמלו |  |
| 1:2.529 | תשבע |  |
| 1:2.530 | נפשו |  |
| 1:2.531 | בעדן |  |
| 1:2.532 | גן |  |
| 1:2.533 | אלהים | ! |


### Translation under review (1:2)

English: Let the enlightened man reflect, and let the fearful and trembling man lament over the word of the LORD, over the state of the people of Israel, upon seeing that every people and each people, of all the divisions of the nations, cherish their tongue and magnify their language, in which they were educated. Their faith and their religion were engraved in it, each man in his own tongue. And we, the people of Israel, since the time that we were exiled from our land, have abandoned our holy tongue, in which the first Adam spoke and in which our holy Torah was given, when His glory, may He be blessed, was revealed on Mount Sinai, and He said to His people, I am the LORD your God, etc., and we have not done as we were commanded, and you shall teach them to your children, to speak of them. And as Rashi of blessed memory explained in Parashat Ekev, “from here they said: ‘When the child begins to speak, his father speaks with him in the holy tongue and teaches him Torah.’” Thus far his words, but all our conversation with our children is in the language of the people among whom we dwell, and most of the children’s teachers too, are not able to speak all their needs in the holy tongue, even though they know it, for they have not tried these, and there are no systems for teaching the children of the elementary school. It is known and famous, the compassion of the father for his sons, and how many toils and exertions the parents exert in raising their children. With most of their attention on all their children’s needs: his food and his drink, suited to him, his bedding, and his clothing! But when he comes, eager to bring him into school, he casts him off with a teacher of children, in narrow and dark houses, on remnants of mats on the ground, without cleanliness, for the teacher does not clean the children’s abode, only from one foot to the next. All day they sit crowded, without order, and breathe foul and musty air. And they drink salty water, and because of this, when some months pass since he was brought to school, growths and a bad boil come up on his head. And their fathers do not look after them to bring them clear, sweet water to drink, nor clean bedding to sit on, and on most days they sit on the ground without a mat! In the new era, from the year 5650 of the full count of creation, many of the enlightened were stirred and sought an organisation of the schools and an organisation of their children for the study of Torah and the study of the languages Hebrew, Arabic and Turkish, which are in use! And they strove to seat enlightened teachers in the Torah of the LORD, to manage the children of the school in moral upbringing and cleanliness of the place of learning, in bright houses and spacious windows in order to breathe fresh and clean air! And all the labour of those enlightened men came to nothing, and their plan did not come to fruition, and the children of the elementary school remained as they were, with a teacher, in dark and musty rooms, sitting in a crowd: this one learns the alef-bet, and this one learns Scripture, and this one learns with the Targum of Onkelos and the translation of Rabbi Saadia Gaon in Arabic, without understanding. Also the very smallest, who are not at all fit to learn, like a child of two and three years, are brought there. And many times the child relieves himself there. And the teacher sends to his mother that she should come and wipe her son’s excrement! And at the call of the reader of the periodicals in those days, from Mr. Eliezer Ben-Yehuda, for the revival of the holy Hebrew language, our soul yearned only for words and mere conversations, and it did not come to realization until the year 5670, when our exalted government, the government of Turkey, then deigned to open schools for the children of Israel, and there was made there a fine organisation for the children of Israel, in spacious upper rooms with open windows, and bedding of wool, and cloths of goat-hair, and also tables for the older children to learn and to write on them! Then we said: behold, the time has come to bring out Ashtonot Ish Yerushalayim, may it be rebuilt and established, which before she began to bear the great man in the revival of the Hebrew language, namely Eliezer Ben-Yehuda, and we acknowledge our debt to him for what he did and made, and with all his strength he toiled for the good of the people of Israel and the settlement of the land of Israel and the revival of the Hebrew language; this holy work he performed with wondrous alacrity and diligence; he made it succeed and raised it from the dust, and after its decay it had a new bloom, and because of his love for the holy tongue its wrinkles were smoothed out, and the appearance of its face showed itself and blossomed, and it found grace and honour in the eyes of its seekers, its knowers, and its acquaintances, and they awoke from their sleep to gird up their senses, to widen it and to enrich it as in the days of its youth, before the exile of the people of Israel! And because of this mitzvah, may his soul be bound up in the bundle of life with the LORD our God, and from the fruit of his labour may his soul be satisfied in Eden, the garden of God!

Literal: let-him-reflect the-enlightened-man and-let-him-lament the-fearful-one and-the-trembling-one concerning the-word-of the-LORD concerning the-state-of the-people the-children-of Israel upon-his-seeing that every people and-each of-all the-divisions of-the-nations cherish (object-marker) their-tongue and-magnify (object-marker) their-language which they-were-educated in-it and-in-which were-engraved in-it their-faith and-their-religion each-man to-his-own-tongue and-we people of-the-children-of Israel from-the-time that we-were-exiled from-our-land we-have-abandoned (object-marker) our-tongue holy in-which spoke in-it Adam the-first and-in-it was-given our-Torah holy when-revealed His-glory may-He-be-blessed on the-mountain Sinai and-He-said to-His-people I-am the-LORD your-God etc. and-not we-did as-commanded upon-us and-you-shall-teach them (object-marker) your-children to-speak of-them and-as explained Rashi of-blessed-memory in-the-parashah-of Ekev from-here they-said when-the-child begins to-speak his-father speaks-with him in-the-language of-the-holy and-teaches-him Torah thus-far-his-words but all our-conversation with our-children in-the-language of-the-people among-whom we dwell among-them and-the-teachers-of children also they most-of-them not are-able to-speak all their-needs in-the-language of-the-holy even though-they-know it for not they-tried these and-there-is-no system for-teaching children of the-house-of the-elementary-school it-is-known it and-famous the-compassion-of the-father for his-sons and-how-many toils and-exertions they-exert-themselves the-parents in-raising their-children with the-great-majority their-attention to all the-needs of-their-children in-his-food and-his-drink suited to-him and-his-bedding and-his-clothing but when he-comes eager to-bring-him-in to-the-house-of school he-casts-him-off to the-teacher-of children in-houses narrow and-dark on remnants-of mats on the-ground without cleanliness for not cleans the-teacher (object-marker) the-abode-of the-children only from-foot to-foot another and-all the-day they-sit crowded without order and-they-breathe air foul and-musty and-they-drink water salty and-because-of this when pass some months from-the-time that-he-was-brought to the-house-of school come-up on-his-head growths and-boils bad and-their-fathers not look-after over-them to-bring for-them water clear and-sweet to-drink nor bedding clean to-sit on-it and-for-most of-the-days they-sit on the-ground without a-mat in-the-period-of the-time the-new from-the-year 5650 of-the full-count of-creation were-stirred many of-the-enlightened and-they-sought an-organisation of-the schools and-an-organisation of-their-children for-the-study-of Torah and-the-study-of the-languages Hebrew and-Arabic and-Turkish in-use and-they-strove to-seat teachers enlightened in-the-Torah-of the-LORD to-manage (object-marker) the-children-of the-house-of school in-education moral and-in-cleanliness the-abode-of the-place-of-learning in-houses bright and-with-windows spacious in-order to-breathe air fresh and-clean and-all their-labour of the-enlightened those was to-nothing and-not came-out their-plan to-fruition and-the-children of the-house-of the-elementary-school remained as-they-were with a-teacher in-rooms dark and-musty sitting in-a-crowd this-one learns alef bet and-this-one learns Scripture and-this-one learns with the-Targum of-Onkelos and-the-translation Rabbi-Saadia-Gaon in-Arabic without understanding also the-smallest-of the-small who-are-not fit to-learn at-all like two and-three years they-bring there and-sometimes many the-child relieves-himself there and-the-teacher sends to-his-mother that-she-come and-wipe the-excrement-of her-son and-at-the-call-of the-reader in-the periodicals in-the-days of-those from Mr. Eliezer ben Yehuda for-the-revival-of the-language Hebrew holy yearned our-soul only and-only for-words and-conversations merely and-not came to realization until the-year 5670 which then deigned our-government the-exalted the-government-of Turkey to-open the-houses-of school for-the-children-of the-sons-of Israel and-it-was-made there an-organisation fine for-the-children-of the-sons-of Israel in-upper-rooms spacious and-windows open and-bedding of wool and-cloths of-goat-hair and-also tables for-the-children older to-learn and-to-write on-them then we-said behold has-come the-time to-bring-out into-light Ashtonot the-man-of Jerusalem may-it-be-rebuilt-and-established which before she-began to-bear the-man the-great in-the-revival-of the-language Hebrew he namely Eliezer ben Yehuda and-we-acknowledge our-debt-to-him for what he-did and-made and-with-all his-strength he-toiled for-the-good-of the-people of-Israel and-the-settlement of-the-land of-Israel and-the-revival-of the-language Hebrew and-the-work-of the-holy this he-performed with-alacrity and-diligence wondrous and-he-made and-succeeded to-raise-it from-dust and-after its-decay it-had for-it a-new-bloom and-because-of his-love for-the-language holy were-smoothed-out its-wrinkles and-the-appearance-of its-face showed-itself and-blossomed and-it-found grace and-honour in-the-eyes-of its-seekers its-knowers and-its-acquaintances and-they-awoke from-their-sleep to-gird-up their-senses to-widen-it and-to-enrich-it as-in-the-days-of its-youth before the-exile-of the-people-of Israel and-because-of mitzvah this may-be his-soul bound in-the-bundle-of life (object-marker) the-LORD our-God and-from-the-fruit-of his-labour shall-be-satisfied his-soul in-Eden the-garden-of God

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:2.1 | let him reflect |  |  | Yitpa'el jussive: 'let him reflect', a wish addressed to the reader, not a plain future. |
| 1:2.2 | the enlightened man |  |  |  |
| 1:2.3 | and let him lament |  |  | Vav with a jussive, continuing the wish. |
| 1:2.4 | the fearful one |  |  |  |
| 1:2.5 | and the trembling one |  |  |  |
| 1:2.6 | concerning |  |  |  |
| 1:2.7 | the word of |  |  |  |
| 1:2.8 | the LORD |  |  | The Name as written (ה׳ with a geresh), read 'the LORD'. |
| 1:2.9 | concerning |  |  |  |
| 1:2.10 | the state of |  |  |  |
| 1:2.11 | the people |  |  |  |
| 1:2.12 | the children of |  |  |  |
| 1:2.13 | Israel |  |  |  |
| 1:2.14 | upon his seeing |  |  |  |
| 1:2.15 | that |  |  |  |
| 1:2.16 | every |  |  |  |
| 1:2.17 | people |  |  |  |
| 1:2.18 | and each |  |  |  |
| 1:2.19 | of all |  |  |  |
| 1:2.20 | the divisions |  |  |  |
| 1:2.21 | of the nations |  |  |  |
| 1:2.22 | cherish |  |  |  |
| 1:2.23 | (object marker) |  |  |  |
| 1:2.24 | their tongue |  |  |  |
| 1:2.25 | and magnify |  |  |  |
| 1:2.26 | (object marker) |  |  |  |
| 1:2.27 | their language |  |  |  |
| 1:2.28 | which |  |  |  |
| 1:2.29 | they were educated |  |  |  |
| 1:2.30 | in it |  |  |  |
| 1:2.31 | and in which |  |  |  |
| 1:2.32 | were engraved |  |  |  |
| 1:2.33 | in it |  |  |  |
| 1:2.34 | their faith |  |  |  |
| 1:2.35 | and their religion |  |  |  |
| 1:2.36 | each man |  |  |  |
| 1:2.37 | to his own tongue |  |  |  |
| 1:2.38 | and we |  |  |  |
| 1:2.39 | people |  |  |  |
| 1:2.40 | of the children of |  |  |  |
| 1:2.41 | Israel |  |  |  |
| 1:2.42 | from the time |  |  |  |
| 1:2.43 | that |  |  |  |
| 1:2.44 | we were exiled |  |  |  |
| 1:2.45 | from our land |  |  |  |
| 1:2.46 | we have abandoned |  |  |  |
| 1:2.47 | (object marker) |  |  |  |
| 1:2.48 | our tongue |  |  |  |
| 1:2.49 | holy |  |  |  |
| 1:2.50 | in which |  |  |  |
| 1:2.51 | spoke |  |  |  |
| 1:2.52 | in it |  |  |  |
| 1:2.53 | Adam |  |  |  |
| 1:2.54 | the first |  |  |  |
| 1:2.55 | and in it |  |  |  |
| 1:2.56 | was given |  |  |  |
| 1:2.57 | our Torah |  |  |  |
| 1:2.58 | holy |  |  |  |
| 1:2.59 | when revealed |  |  |  |
| 1:2.60 | His glory |  |  |  |
| 1:2.61 | may He be blessed |  |  | Abbreviation of yitbarakh, 'may He be blessed'; placed where the Hebrew places it. |
| 1:2.62 | on |  |  |  |
| 1:2.63 | the mountain |  |  |  |
| 1:2.64 | Sinai |  |  |  |
| 1:2.65 | and He said |  |  |  |
| 1:2.66 | to His people |  |  |  |
| 1:2.67 | I am |  |  |  |
| 1:2.68 | the LORD |  |  |  |
| 1:2.69 | your God |  |  |  |
| 1:2.70 | etc. |  |  | The author's 'etc.' marks a truncated quotation of Exod 20:2; see apparatus. |
| 1:2.71 | and not |  |  |  |
| 1:2.72 | we did |  |  |  |
| 1:2.73 | as commanded |  |  |  |
| 1:2.74 | upon us |  |  |  |
| 1:2.75 | and you shall teach |  |  |  |
| 1:2.76 | them |  |  |  |
| 1:2.77 | (object marker) |  |  |  |
| 1:2.78 | your children |  |  |  |
| 1:2.79 | to speak |  |  |  |
| 1:2.80 | of them |  |  |  |
| 1:2.81 | and as |  |  |  |
| 1:2.82 | explained |  |  |  |
| 1:2.83 | Rashi |  |  |  |
| 1:2.84 | of blessed memory |  |  |  |
| 1:2.85 | in the parashah of |  |  |  |
| 1:2.86 | Ekev |  |  |  |
| 1:2.87 | from here |  |  |  |
| 1:2.88 | they said |  |  |  |
| 1:2.89 | when the child |  |  |  |
| 1:2.90 | begins |  |  |  |
| 1:2.91 | to speak |  |  |  |
| 1:2.92 | his father |  |  |  |
| 1:2.93 | speaks with |  |  |  |
| 1:2.94 | him |  |  |  |
| 1:2.95 | in the language |  |  |  |
| 1:2.96 | of the holy |  |  |  |
| 1:2.97 | and teaches him |  |  |  |
| 1:2.98 | Torah |  |  |  |
| 1:2.99 | thus far his words |  |  | Closes Rashi's words; the quotation is Rashi's, citing the Sages. |
| 1:2.100 | but |  |  |  |
| 1:2.101 | all |  |  |  |
| 1:2.102 | our conversation |  |  |  |
| 1:2.103 | with |  |  |  |
| 1:2.104 | our children |  |  |  |
| 1:2.105 | in the language |  |  |  |
| 1:2.106 | of the people |  |  |  |
| 1:2.107 | among whom |  |  |  |
| 1:2.108 | we |  |  |  |
| 1:2.109 | dwell |  |  |  |
| 1:2.110 | among them |  |  |  |
| 1:2.111 | and the teachers of |  |  |  |
| 1:2.112 | children |  |  |  |
| 1:2.113 | also |  |  |  |
| 1:2.114 | they |  |  |  |
| 1:2.115 | most of them |  |  |  |
| 1:2.116 | not |  |  |  |
| 1:2.117 | are able |  |  |  |
| 1:2.118 | to speak |  |  |  |
| 1:2.119 | all |  |  |  |
| 1:2.120 | their needs |  |  |  |
| 1:2.121 | in the language |  |  |  |
| 1:2.122 | of the holy |  |  |  |
| 1:2.123 | even |  |  |  |
| 1:2.124 | though they know |  |  |  |
| 1:2.125 | it |  |  |  |
| 1:2.126 | for |  |  |  |
| 1:2.127 | not |  |  |  |
| 1:2.128 | they tried |  |  |  |
| 1:2.129 | these |  |  |  |
| 1:2.130 | and there is no |  |  |  |
| 1:2.131 | system |  |  |  |
| 1:2.132 | for teaching |  |  |  |
| 1:2.133 | children |  |  |  |
| 1:2.134 | of |  |  |  |
| 1:2.135 | the house of |  |  |  |
| 1:2.136 | the elementary school |  |  | Beit rabban, 'house of the master': the traditional elementary school. |
| 1:2.137 | it is known |  |  |  |
| 1:2.138 | it |  |  |  |
| 1:2.139 | and famous |  |  |  |
| 1:2.140 | the compassion of |  |  |  |
| 1:2.141 | the father |  |  |  |
| 1:2.142 | for |  |  |  |
| 1:2.143 | his sons |  |  |  |
| 1:2.144 | and how many |  |  |  |
| 1:2.145 | toils |  |  |  |
| 1:2.146 | and exertions |  |  |  |
| 1:2.147 | they exert themselves |  |  |  |
| 1:2.148 | the parents |  |  |  |
| 1:2.149 | in raising |  |  |  |
| 1:2.150 | their children |  |  |  |
| 1:2.151 | with |  |  |  |
| 1:2.152 | the great majority |  |  |  |
| 1:2.153 | their attention |  |  |  |
| 1:2.154 | to |  |  |  |
| 1:2.155 | all |  |  |  |
| 1:2.156 | the needs |  |  |  |
| 1:2.157 | of their children |  |  |  |
| 1:2.158 | in his food |  |  |  |
| 1:2.159 | and his drink |  |  |  |
| 1:2.160 | suited |  |  |  |
| 1:2.161 | to him |  |  |  |
| 1:2.162 | and his bedding |  |  |  |
| 1:2.163 | and his clothing |  |  |  |
| 1:2.164 | but |  |  |  |
| 1:2.165 | when |  |  |  |
| 1:2.166 | he comes |  |  |  |
| 1:2.167 | eager |  |  |  |
| 1:2.168 | to bring him in |  |  |  |
| 1:2.169 | to the house of |  |  |  |
| 1:2.170 | school |  |  |  |
| 1:2.171 | he casts him off |  |  |  |
| 1:2.172 | to |  |  |  |
| 1:2.173 | the teacher of |  |  |  |
| 1:2.174 | children |  |  |  |
| 1:2.175 | in houses |  |  |  |
| 1:2.176 | narrow |  |  |  |
| 1:2.177 | and dark |  |  |  |
| 1:2.178 | on |  |  |  |
| 1:2.179 | remnants of |  |  |  |
| 1:2.180 | mats |  |  |  |
| 1:2.181 | on |  |  |  |
| 1:2.182 | the ground |  |  |  |
| 1:2.183 | without |  |  |  |
| 1:2.184 | cleanliness |  |  |  |
| 1:2.185 | for |  |  |  |
| 1:2.186 | not |  |  |  |
| 1:2.187 | cleans |  |  |  |
| 1:2.188 | the teacher |  |  |  |
| 1:2.189 | (object marker) |  |  |  |
| 1:2.190 | the abode of |  |  |  |
| 1:2.191 | the children |  |  |  |
| 1:2.192 | only |  |  |  |
| 1:2.193 | from foot |  |  |  |
| 1:2.194 | to foot |  |  |  |
| 1:2.195 | another |  |  |  |
| 1:2.196 | and all |  |  |  |
| 1:2.197 | the day |  |  |  |
| 1:2.198 | they sit |  |  |  |
| 1:2.199 | crowded |  |  |  |
| 1:2.200 | without |  |  |  |
| 1:2.201 | order |  |  |  |
| 1:2.202 | and they breathe |  |  |  |
| 1:2.203 | air |  |  |  |
| 1:2.204 | foul |  |  |  |
| 1:2.205 | and musty |  |  |  |
| 1:2.206 | and they drink |  |  |  |
| 1:2.207 | water |  |  |  |
| 1:2.208 | salty |  |  |  |
| 1:2.209 | and because of |  |  |  |
| 1:2.210 | this |  |  |  |
| 1:2.211 | when |  |  |  |
| 1:2.212 | pass |  |  |  |
| 1:2.213 | some |  |  |  |
| 1:2.214 | months |  |  |  |
| 1:2.215 | from the time |  |  |  |
| 1:2.216 | that he was brought |  |  |  |
| 1:2.217 | to |  |  |  |
| 1:2.218 | the house of |  |  |  |
| 1:2.219 | school |  |  |  |
| 1:2.220 | come up |  |  |  |
| 1:2.221 | on his head |  |  |  |
| 1:2.222 | growths |  |  |  |
| 1:2.223 | and boils |  |  |  |
| 1:2.224 | bad |  |  |  |
| 1:2.225 | and their fathers |  |  |  |
| 1:2.226 | not |  |  |  |
| 1:2.227 | look after |  |  |  |
| 1:2.228 | over them |  |  |  |
| 1:2.229 | to bring |  |  |  |
| 1:2.230 | for them |  |  |  |
| 1:2.231 | water |  |  |  |
| 1:2.232 | clear |  |  |  |
| 1:2.233 | and sweet |  |  |  |
| 1:2.234 | to drink |  |  |  |
| 1:2.235 | nor |  |  |  |
| 1:2.236 | bedding |  |  |  |
| 1:2.237 | clean |  |  |  |
| 1:2.238 | to sit |  |  |  |
| 1:2.239 | on it |  |  |  |
| 1:2.240 | and for most |  |  |  |
| 1:2.241 | of the days |  |  |  |
| 1:2.242 | they sit |  |  |  |
| 1:2.243 | on |  |  |  |
| 1:2.244 | the ground |  |  |  |
| 1:2.245 | without |  |  |  |
| 1:2.246 | a mat |  |  |  |
| 1:2.247 | in the period of |  |  |  |
| 1:2.248 | the time |  |  |  |
| 1:2.249 | the new |  |  |  |
| 1:2.250 | from the year |  |  |  |
| 1:2.251 | 5650 |  |  | A year without the thousands: 5650 of the era from creation (lefirat gadol). |
| 1:2.252 | of the |  |  |  |
| 1:2.253 | full count |  |  |  |
| 1:2.254 | of creation |  |  |  |
| 1:2.255 | were stirred |  |  |  |
| 1:2.256 | many |  |  |  |
| 1:2.257 | of the enlightened |  |  |  |
| 1:2.258 | and they sought |  |  |  |
| 1:2.259 | an organisation |  |  |  |
| 1:2.260 | of the |  |  |  |
| 1:2.261 | schools |  |  |  |
| 1:2.262 | and an organisation |  |  |  |
| 1:2.263 | of their children |  |  |  |
| 1:2.264 | for the study of |  |  |  |
| 1:2.265 | Torah |  |  |  |
| 1:2.266 | and the study of |  |  |  |
| 1:2.267 | the languages |  |  |  |
| 1:2.268 | Hebrew |  |  |  |
| 1:2.269 | and Arabic |  |  |  |
| 1:2.270 | and Turkish |  |  |  |
| 1:2.271 | in use |  |  |  |
| 1:2.272 | and they strove |  |  |  |
| 1:2.273 | to seat |  |  |  |
| 1:2.274 | teachers |  |  |  |
| 1:2.275 | enlightened |  |  |  |
| 1:2.276 | in the Torah of |  |  |  |
| 1:2.277 | the LORD |  |  |  |
| 1:2.278 | to manage |  |  |  |
| 1:2.279 | (object marker) |  |  |  |
| 1:2.280 | the children of |  |  |  |
| 1:2.281 | the house of |  |  |  |
| 1:2.282 | school |  |  |  |
| 1:2.283 | in education |  |  |  |
| 1:2.284 | moral |  |  |  |
| 1:2.285 | and in cleanliness |  |  |  |
| 1:2.286 | the abode of |  |  |  |
| 1:2.287 | the place of learning |  |  |  |
| 1:2.288 | in houses |  |  |  |
| 1:2.289 | bright |  |  |  |
| 1:2.290 | and with windows |  |  |  |
| 1:2.291 | spacious |  |  |  |
| 1:2.292 | in order |  |  |  |
| 1:2.293 | to breathe |  |  |  |
| 1:2.294 | air |  |  |  |
| 1:2.295 | fresh |  |  |  |
| 1:2.296 | and clean |  |  |  |
| 1:2.297 | and all |  |  |  |
| 1:2.298 | their labour |  |  |  |
| 1:2.299 | of |  |  |  |
| 1:2.300 | the enlightened |  |  |  |
| 1:2.301 | those |  |  |  |
| 1:2.302 | was |  |  |  |
| 1:2.303 | to nothing |  |  |  |
| 1:2.304 | and not |  |  |  |
| 1:2.305 | came out |  |  |  |
| 1:2.306 | their plan |  |  |  |
| 1:2.307 | to fruition |  |  |  |
| 1:2.308 | and the children |  |  |  |
| 1:2.309 | of |  |  |  |
| 1:2.310 | the house of |  |  |  |
| 1:2.311 | the elementary school |  |  |  |
| 1:2.312 | remained |  |  |  |
| 1:2.313 | as they were |  |  |  |
| 1:2.314 | with |  |  |  |
| 1:2.315 | a teacher |  |  |  |
| 1:2.316 | in rooms |  |  |  |
| 1:2.317 | dark |  |  |  |
| 1:2.318 | and musty |  |  |  |
| 1:2.319 | sitting |  |  |  |
| 1:2.320 | in a crowd |  |  | The word is obscure; 'in a crowd' is by sense. |
| 1:2.321 | this one |  |  |  |
| 1:2.322 | learns |  |  |  |
| 1:2.323 | alef |  |  | Aramaic form of the alphabet's first two letters, alef and bet. |
| 1:2.324 | bet |  |  |  |
| 1:2.325 | and this one |  |  |  |
| 1:2.326 | learns |  |  |  |
| 1:2.327 | Scripture |  |  |  |
| 1:2.328 | and this one |  |  |  |
| 1:2.329 | learns |  |  |  |
| 1:2.330 | with |  |  |  |
| 1:2.331 | the Targum |  |  |  |
| 1:2.332 | of Onkelos |  |  | Onkelos's Aramaic Targum of the Torah. |
| 1:2.333 | and the translation |  |  |  |
| 1:2.334 | Rabbi Saadia Gaon |  |  | Saadia Gaon (882-942), his Arabic translation of Scripture. |
| 1:2.335 | in Arabic |  |  |  |
| 1:2.336 | without |  |  |  |
| 1:2.337 | understanding |  |  |  |
| 1:2.338 | also |  |  |  |
| 1:2.339 | the smallest of |  |  |  |
| 1:2.340 | the small |  |  |  |
| 1:2.341 | who are not |  |  |  |
| 1:2.342 | fit |  |  |  |
| 1:2.343 | to learn |  |  |  |
| 1:2.344 | at all |  |  |  |
| 1:2.345 | like |  |  |  |
| 1:2.346 | two |  |  |  |
| 1:2.347 | and three |  |  |  |
| 1:2.348 | years |  |  |  |
| 1:2.349 | they bring |  |  |  |
| 1:2.350 | there |  |  |  |
| 1:2.351 | and sometimes |  |  |  |
| 1:2.352 | many |  |  |  |
| 1:2.353 | the child |  |  |  |
| 1:2.354 | relieves himself |  |  |  |
| 1:2.355 | there |  |  |  |
| 1:2.356 | and the teacher |  |  |  |
| 1:2.357 | sends |  |  |  |
| 1:2.358 | to his mother |  |  |  |
| 1:2.359 | that she come |  |  |  |
| 1:2.360 | and wipe |  |  |  |
| 1:2.361 | the excrement of |  |  |  |
| 1:2.362 | her son |  |  |  |
| 1:2.363 | and at the call of |  |  |  |
| 1:2.364 | the reader |  |  |  |
| 1:2.365 | in the |  |  | Lit. 'letters of the times', the periodicals. |
| 1:2.366 | periodicals |  |  |  |
| 1:2.367 | in the days |  |  |  |
| 1:2.368 | of those |  |  |  |
| 1:2.369 | from |  |  |  |
| 1:2.370 | Mr. |  |  | Mr., an honorific abbreviation here, not the Name, though written ה׳. |
| 1:2.371 | Eliezer |  |  |  |
| 1:2.372 | ben |  |  |  |
| 1:2.373 | Yehuda |  |  |  |
| 1:2.374 | for the revival of |  |  |  |
| 1:2.375 | the language |  |  |  |
| 1:2.376 | Hebrew |  |  |  |
| 1:2.377 | holy |  |  |  |
| 1:2.378 | yearned |  |  |  |
| 1:2.379 | our soul |  |  |  |
| 1:2.380 | only |  |  |  |
| 1:2.381 | and only |  |  |  |
| 1:2.382 | for words |  |  |  |
| 1:2.383 | and conversations |  |  |  |
| 1:2.384 | merely |  |  |  |
| 1:2.385 | and not |  |  |  |
| 1:2.386 | came |  |  |  |
| 1:2.387 | to |  |  |  |
| 1:2.388 | realization |  |  |  |
| 1:2.389 | until |  |  |  |
| 1:2.390 | the year |  |  |  |
| 1:2.391 | 5670 |  |  | 5670 of the era from creation; the civil year is in apparatus. |
| 1:2.392 | which |  |  |  |
| 1:2.393 | then |  |  |  |
| 1:2.394 | deigned |  |  |  |
| 1:2.395 | our government |  |  |  |
| 1:2.396 | the exalted |  |  |  |
| 1:2.397 | the government of |  |  |  |
| 1:2.398 | Turkey |  |  | Togarmah, the Hebrew name the author uses for Turkey. |
| 1:2.399 | to open |  |  |  |
| 1:2.400 | the houses of |  |  |  |
| 1:2.401 | school |  |  |  |
| 1:2.402 | for the children of |  |  |  |
| 1:2.403 | the sons of |  |  |  |
| 1:2.404 | Israel |  |  |  |
| 1:2.405 | and it was made |  |  |  |
| 1:2.406 | there |  |  |  |
| 1:2.407 | an organisation |  |  |  |
| 1:2.408 | fine |  |  |  |
| 1:2.409 | for the children of |  |  |  |
| 1:2.410 | the sons of |  |  |  |
| 1:2.411 | Israel |  |  |  |
| 1:2.412 | in upper rooms |  |  |  |
| 1:2.413 | spacious |  |  |  |
| 1:2.414 | and windows |  |  |  |
| 1:2.415 | open |  |  |  |
| 1:2.416 | and bedding |  |  |  |
| 1:2.417 | of |  |  |  |
| 1:2.418 | wool |  |  |  |
| 1:2.419 | and cloths |  |  |  |
| 1:2.420 | of goat-hair |  |  |  |
| 1:2.421 | and also |  |  |  |
| 1:2.422 | tables |  |  |  |
| 1:2.423 | for the children |  |  |  |
| 1:2.424 | older |  |  |  |
| 1:2.425 | to learn |  |  |  |
| 1:2.426 | and to write |  |  |  |
| 1:2.427 | on them |  |  |  |
| 1:2.428 | then |  |  |  |
| 1:2.429 | we said |  |  |  |
| 1:2.430 | behold |  |  |  |
| 1:2.431 | has come |  |  |  |
| 1:2.432 | the time |  |  |  |
| 1:2.433 | to bring out |  |  |  |
| 1:2.434 | into light |  |  |  |
| 1:2.435 | Ashtonot |  |  | Ashtonot, lit. 'reflections' or 'thoughts', the title. |
| 1:2.436 | the man of |  |  |  |
| 1:2.437 | Jerusalem |  |  |  |
| 1:2.438 | may it be rebuilt and established |  |  | Tav-vav: may it be rebuilt and established (Jerusalem). |
| 1:2.439 | which |  |  |  |
| 1:2.440 | before |  |  |  |
| 1:2.441 | she began |  |  |  |
| 1:2.442 | to bear |  |  |  |
| 1:2.443 | the man |  |  |  |
| 1:2.444 | the great |  |  |  |
| 1:2.445 | in the revival of |  |  |  |
| 1:2.446 | the language |  |  |  |
| 1:2.447 | Hebrew |  |  |  |
| 1:2.448 | he |  |  |  |
| 1:2.449 | namely |  |  |  |
| 1:2.450 | Eliezer |  |  |  |
| 1:2.451 | ben |  |  |  |
| 1:2.452 | Yehuda |  |  |  |
| 1:2.453 | and we acknowledge |  |  | Aramaic idiom, rendered by its sense: a public acknowledgement of debt. |
| 1:2.454 | our debt to him |  |  |  |
| 1:2.455 | for |  |  |  |
| 1:2.456 | what |  |  |  |
| 1:2.457 | he did |  |  |  |
| 1:2.458 | and made |  |  |  |
| 1:2.459 | and with all |  |  |  |
| 1:2.460 | his strength |  |  |  |
| 1:2.461 | he toiled |  |  |  |
| 1:2.462 | for the good of |  |  |  |
| 1:2.463 | the people |  |  |  |
| 1:2.464 | of Israel |  |  |  |
| 1:2.465 | and the settlement |  |  |  |
| 1:2.466 | of the land |  |  |  |
| 1:2.467 | of Israel |  |  |  |
| 1:2.468 | and the revival of |  |  |  |
| 1:2.469 | the language |  |  |  |
| 1:2.470 | Hebrew |  |  |  |
| 1:2.471 | and the work of |  |  |  |
| 1:2.472 | the holy |  |  |  |
| 1:2.473 | this |  |  |  |
| 1:2.474 | he performed |  |  |  |
| 1:2.475 | with alacrity |  |  |  |
| 1:2.476 | and diligence |  |  |  |
| 1:2.477 | wondrous |  |  |  |
| 1:2.478 | and he made |  |  |  |
| 1:2.479 | and succeeded |  |  |  |
| 1:2.480 | to raise it |  |  |  |
| 1:2.481 | from dust |  |  |  |
| 1:2.482 | and after |  |  |  |
| 1:2.483 | its decay |  |  |  |
| 1:2.484 | it had |  |  |  |
| 1:2.485 | for it |  |  |  |
| 1:2.486 | a new bloom |  |  |  |
| 1:2.487 | and because of |  |  |  |
| 1:2.488 | his love |  |  |  |
| 1:2.489 | for the language |  |  |  |
| 1:2.490 | holy |  |  |  |
| 1:2.491 | were smoothed out |  |  |  |
| 1:2.492 | its wrinkles |  |  |  |
| 1:2.493 | and the appearance of |  |  |  |
| 1:2.494 | its face |  |  |  |
| 1:2.495 | showed itself |  |  |  |
| 1:2.496 | and blossomed |  |  |  |
| 1:2.497 | and it found |  |  |  |
| 1:2.498 | grace |  |  |  |
| 1:2.499 | and honour |  |  |  |
| 1:2.500 | in the eyes of |  |  |  |
| 1:2.501 | its seekers |  |  |  |
| 1:2.502 | its knowers |  |  |  |
| 1:2.503 | and its acquaintances |  |  |  |
| 1:2.504 | and they awoke |  |  |  |
| 1:2.505 | from their sleep |  |  |  |
| 1:2.506 | to gird up |  |  |  |
| 1:2.507 | their senses |  |  |  |
| 1:2.508 | to widen it |  |  |  |
| 1:2.509 | and to enrich it |  |  |  |
| 1:2.510 | as in the days of |  |  |  |
| 1:2.511 | its youth |  |  |  |
| 1:2.512 | before |  |  |  |
| 1:2.513 | the exile of |  |  |  |
| 1:2.514 | the people of |  |  |  |
| 1:2.515 | Israel |  |  |  |
| 1:2.516 | and because of |  |  |  |
| 1:2.517 | mitzvah |  |  |  |
| 1:2.518 | this |  |  |  |
| 1:2.519 | may be |  |  |  |
| 1:2.520 | his soul |  |  |  |
| 1:2.521 | bound |  |  |  |
| 1:2.522 | in the bundle of |  |  |  |
| 1:2.523 | life |  |  |  |
| 1:2.524 | (object marker) |  |  |  |
| 1:2.525 | the LORD |  |  |  |
| 1:2.526 | our God |  |  |  |
| 1:2.527 | and from the fruit of |  |  |  |
| 1:2.528 | his labour |  |  |  |
| 1:2.529 | shall be satisfied |  |  |  |
| 1:2.530 | his soul |  |  |  |
| 1:2.531 | in Eden |  |  |  |
| 1:2.532 | the garden of |  |  |  |
| 1:2.533 | God |  |  | Gan Elohim, the garden of God. |

Phrases:
- 1:2.1 1:2.2 1:2.3 1:2.4 1:2.5 | Let the enlightened man reflect, and let the fearful and trembling man lament | let-him-reflect the-enlightened-man and-let-him-lament the-fearful-one and-the-trembling-one
- 1:2.6 1:2.7 1:2.8 | over the word of the LORD, | concerning the-word-of the-LORD
- 1:2.9 1:2.10 1:2.11 1:2.12 1:2.13 | over the state of the people of Israel, | concerning the-state-of the-people the-children-of Israel
- 1:2.14 1:2.15 | upon seeing that | upon-his-seeing that
- 1:2.16 1:2.17 1:2.18 1:2.19 1:2.20 1:2.21 | every people and each people, of all the divisions of the nations, | every people and-each of-all the-divisions of-the-nations
- 1:2.22 1:2.23 1:2.24 1:2.25 1:2.26 1:2.27 1:2.28 1:2.29 1:2.30 | cherish their tongue and magnify their language, in which they were educated. | cherish (object-marker) their-tongue and-magnify (object-marker) their-language which they-were-educated in-it
- 1:2.31 1:2.32 1:2.33 1:2.34 1:2.35 1:2.36 1:2.37 | Their faith and their religion were engraved in it, each man in his own tongue. | and-in-which were-engraved in-it their-faith and-their-religion each-man to-his-own-tongue
- 1:2.38 1:2.39 1:2.40 1:2.41 1:2.42 1:2.43 1:2.44 1:2.45 1:2.46 | And we, the people of Israel, since the time that we were exiled from our land, have abandoned | and-we people of-the-children-of Israel from-the-time that we-were-exiled from-our-land we-have-abandoned
- 1:2.47 1:2.48 1:2.49 1:2.50 | our holy tongue, in which | (object-marker) our-tongue holy in-which
- 1:2.51 1:2.52 1:2.53 1:2.54 | the first Adam spoke | spoke in-it Adam the-first
- 1:2.55 1:2.56 1:2.57 1:2.58 | and in which our holy Torah was given, | and-in-it was-given our-Torah holy
- 1:2.59 1:2.60 1:2.61 1:2.62 1:2.63 1:2.64 | when His glory, may He be blessed, was revealed on Mount Sinai, | when-revealed His-glory may-He-be-blessed on the-mountain Sinai
- 1:2.65 1:2.66 1:2.67 1:2.68 1:2.69 1:2.70 | and He said to His people, I am the LORD your God, etc., | and-He-said to-His-people I-am the-LORD your-God etc.
- 1:2.71 1:2.72 1:2.73 1:2.74 1:2.75 1:2.76 1:2.77 1:2.78 1:2.79 1:2.80 | and we have not done as we were commanded, and you shall teach them to your children, to speak of them. | and-not we-did as-commanded upon-us and-you-shall-teach them (object-marker) your-children to-speak of-them
- 1:2.81 1:2.82 1:2.83 1:2.84 1:2.85 1:2.86 1:2.87 1:2.88 | And as Rashi of blessed memory explained in Parashat Ekev, “from here they said: | and-as explained Rashi of-blessed-memory in-the-parashah-of Ekev from-here they-said
- 1:2.89 1:2.90 1:2.91 1:2.92 1:2.93 1:2.94 1:2.95 1:2.96 1:2.97 1:2.98 | ‘When the child begins to speak, his father speaks with him in the holy tongue and teaches him Torah.’” | when-the-child begins to-speak his-father speaks-with him in-the-language of-the-holy and-teaches-him Torah
- 1:2.99 | Thus far his words, | thus-far-his-words
- 1:2.100 1:2.101 1:2.102 1:2.103 1:2.104 1:2.105 1:2.106 | but all our conversation with our children is in the language of the people | but all our-conversation with our-children in-the-language of-the-people
- 1:2.107 1:2.108 1:2.109 1:2.110 | among whom we dwell, | among-whom we dwell among-them
- 1:2.111 1:2.112 1:2.113 1:2.114 1:2.115 | and most of the children’s teachers too, | and-the-teachers-of children also they most-of-them
- 1:2.116 1:2.117 1:2.118 | are not able to speak | not are-able to-speak
- 1:2.119 1:2.120 1:2.121 1:2.122 | all their needs in the holy tongue, | all their-needs in-the-language of-the-holy
- 1:2.123 1:2.124 1:2.125 | even though they know it, | even though-they-know it
- 1:2.126 1:2.127 1:2.128 1:2.129 | for they have not tried these, | for not they-tried these
- 1:2.130 1:2.131 1:2.132 1:2.133 1:2.134 1:2.135 1:2.136 | and there are no systems for teaching the children of the elementary school. | and-there-is-no system for-teaching children of the-house-of the-elementary-school
- 1:2.137 1:2.138 1:2.139 1:2.140 1:2.141 1:2.142 1:2.143 1:2.144 1:2.145 1:2.146 1:2.147 1:2.148 1:2.149 1:2.150 | It is known and famous, the compassion of the father for his sons, and how many toils and exertions the parents exert in raising their children. | it-is-known it and-famous the-compassion-of the-father for his-sons and-how-many toils and-exertions they-exert-themselves the-parents in-raising their-children
- 1:2.151 1:2.152 1:2.153 1:2.154 1:2.155 1:2.156 1:2.157 1:2.158 1:2.159 1:2.160 1:2.161 1:2.162 1:2.163 | With most of their attention on all their children’s needs: his food and his drink, suited to him, his bedding, and his clothing! | with the-great-majority their-attention to all the-needs of-their-children in-his-food and-his-drink suited to-him and-his-bedding and-his-clothing
- 1:2.164 1:2.165 1:2.166 1:2.167 1:2.168 1:2.169 1:2.170 | But when he comes, eager to bring him into school, | but when he-comes eager to-bring-him-in to-the-house-of school
- 1:2.171 1:2.172 1:2.173 1:2.174 | he casts him off with a teacher of children, | he-casts-him-off to the-teacher-of children
- 1:2.175 1:2.176 1:2.177 | in narrow and dark houses, | in-houses narrow and-dark
- 1:2.178 1:2.179 1:2.180 1:2.181 1:2.182 | on remnants of mats on the ground, | on remnants-of mats on the-ground
- 1:2.183 1:2.184 | without cleanliness, | without cleanliness
- 1:2.185 1:2.186 1:2.187 1:2.188 1:2.189 1:2.190 1:2.191 1:2.192 1:2.193 1:2.194 1:2.195 | for the teacher does not clean the children’s abode, only from one foot to the next. | for not cleans the-teacher (object-marker) the-abode-of the-children only from-foot to-foot another
- 1:2.196 1:2.197 1:2.198 1:2.199 1:2.200 1:2.201 1:2.202 1:2.203 1:2.204 1:2.205 | All day they sit crowded, without order, and breathe foul and musty air. | and-all the-day they-sit crowded without order and-they-breathe air foul and-musty
- 1:2.206 1:2.207 1:2.208 | And they drink salty water, | and-they-drink water salty
- 1:2.209 1:2.210 1:2.211 1:2.212 1:2.213 1:2.214 1:2.215 1:2.216 1:2.217 1:2.218 1:2.219 1:2.220 1:2.221 1:2.222 1:2.223 1:2.224 | and because of this, when some months pass since he was brought to school, growths and a bad boil come up on his head. | and-because-of this when pass some months from-the-time that-he-was-brought to the-house-of school come-up on-his-head growths and-boils bad
- 1:2.225 1:2.226 1:2.227 1:2.228 1:2.229 1:2.230 1:2.231 1:2.232 1:2.233 1:2.234 1:2.235 1:2.236 1:2.237 1:2.238 1:2.239 | And their fathers do not look after them to bring them clear, sweet water to drink, nor clean bedding to sit on, | and-their-fathers not look-after over-them to-bring for-them water clear and-sweet to-drink nor bedding clean to-sit on-it
- 1:2.240 1:2.241 1:2.242 1:2.243 1:2.244 1:2.245 1:2.246 | and on most days they sit on the ground without a mat! | and-for-most of-the-days they-sit on the-ground without a-mat
- 1:2.247 1:2.248 1:2.249 1:2.250 1:2.251 1:2.252 1:2.253 | In the new era, from the year 5650 of the full count | in-the-period-of the-time the-new from-the-year 5650 of-the full-count
- 1:2.254 1:2.255 1:2.256 1:2.257 | of creation, many of the enlightened were stirred | of-creation were-stirred many of-the-enlightened
- 1:2.258 1:2.259 1:2.260 1:2.261 1:2.262 1:2.263 1:2.264 1:2.265 1:2.266 1:2.267 1:2.268 1:2.269 1:2.270 1:2.271 | and sought an organisation of the schools and an organisation of their children for the study of Torah and the study of the languages Hebrew, Arabic and Turkish, which are in use! | and-they-sought an-organisation of-the schools and-an-organisation of-their-children for-the-study-of Torah and-the-study-of the-languages Hebrew and-Arabic and-Turkish in-use
- 1:2.272 1:2.273 1:2.274 1:2.275 1:2.276 1:2.277 | And they strove to seat enlightened teachers in the Torah of the LORD, | and-they-strove to-seat teachers enlightened in-the-Torah-of the-LORD
- 1:2.278 1:2.279 1:2.280 1:2.281 1:2.282 | to manage the children of the school | to-manage (object-marker) the-children-of the-house-of school
- 1:2.283 1:2.284 1:2.285 | in moral upbringing and cleanliness | in-education moral and-in-cleanliness
- 1:2.286 1:2.287 | of the place of learning, | the-abode-of the-place-of-learning
- 1:2.288 1:2.289 1:2.290 1:2.291 | in bright houses and spacious windows | in-houses bright and-with-windows spacious
- 1:2.292 1:2.293 1:2.294 1:2.295 1:2.296 | in order to breathe fresh and clean air! | in-order to-breathe air fresh and-clean
- 1:2.297 1:2.298 1:2.299 1:2.300 1:2.301 1:2.302 1:2.303 | And all the labour of those enlightened men came to nothing, | and-all their-labour of the-enlightened those was to-nothing
- 1:2.304 1:2.305 1:2.306 1:2.307 | and their plan did not come to fruition, | and-not came-out their-plan to-fruition
- 1:2.308 1:2.309 1:2.310 1:2.311 1:2.312 | and the children of the elementary school remained | and-the-children of the-house-of the-elementary-school remained
- 1:2.313 1:2.314 1:2.315 1:2.316 1:2.317 1:2.318 | as they were, with a teacher, in dark and musty rooms, | as-they-were with a-teacher in-rooms dark and-musty
- 1:2.319 1:2.320 | sitting in a crowd: | sitting in-a-crowd
- 1:2.321 1:2.322 1:2.323 1:2.324 | this one learns the alef-bet, | this-one learns alef bet
- 1:2.325 1:2.326 1:2.327 | and this one learns Scripture, | and-this-one learns Scripture
- 1:2.328 1:2.329 1:2.330 1:2.331 1:2.332 | and this one learns with the Targum of Onkelos | and-this-one learns with the-Targum of-Onkelos
- 1:2.333 1:2.334 1:2.335 | and the translation of Rabbi Saadia Gaon in Arabic, | and-the-translation Rabbi-Saadia-Gaon in-Arabic
- 1:2.336 1:2.337 | without understanding. | without understanding
- 1:2.338 1:2.339 1:2.340 1:2.341 1:2.342 1:2.343 1:2.344 1:2.345 1:2.346 1:2.347 1:2.348 1:2.349 1:2.350 | Also the very smallest, who are not at all fit to learn, like a child of two and three years, are brought there. | also the-smallest-of the-small who-are-not fit to-learn at-all like two and-three years they-bring there
- 1:2.351 1:2.352 1:2.353 1:2.354 1:2.355 | And many times the child relieves himself there. | and-sometimes many the-child relieves-himself there
- 1:2.356 1:2.357 1:2.358 1:2.359 1:2.360 1:2.361 1:2.362 | And the teacher sends to his mother that she should come and wipe her son’s excrement! | and-the-teacher sends to-his-mother that-she-come and-wipe the-excrement-of her-son
- 1:2.363 1:2.364 1:2.365 1:2.366 1:2.367 1:2.368 | And at the call of the reader of the periodicals in those days, | and-at-the-call-of the-reader in-the periodicals in-the-days of-those
- 1:2.369 1:2.370 1:2.371 1:2.372 1:2.373 | from Mr. Eliezer Ben-Yehuda, | from Mr. Eliezer ben Yehuda
- 1:2.374 1:2.375 1:2.376 1:2.377 1:2.378 1:2.379 | for the revival of the holy Hebrew language, our soul yearned | for-the-revival-of the-language Hebrew holy yearned our-soul
- 1:2.380 1:2.381 1:2.382 1:2.383 1:2.384 1:2.385 1:2.386 1:2.387 1:2.388 | only for words and mere conversations, and it did not come to realization | only and-only for-words and-conversations merely and-not came to realization
- 1:2.389 1:2.390 1:2.391 | until the year 5670, | until the-year 5670
- 1:2.392 1:2.393 1:2.394 1:2.395 1:2.396 1:2.397 1:2.398 | when our exalted government, the government of Turkey, then deigned | which then deigned our-government the-exalted the-government-of Turkey
- 1:2.399 1:2.400 1:2.401 1:2.402 1:2.403 1:2.404 | to open schools for the children of Israel, | to-open the-houses-of school for-the-children-of the-sons-of Israel
- 1:2.405 1:2.406 1:2.407 1:2.408 1:2.409 1:2.410 1:2.411 | and there was made there a fine organisation for the children of Israel, | and-it-was-made there an-organisation fine for-the-children-of the-sons-of Israel
- 1:2.412 1:2.413 1:2.414 1:2.415 1:2.416 1:2.417 1:2.418 1:2.419 1:2.420 | in spacious upper rooms with open windows, and bedding of wool, and cloths of goat-hair, | in-upper-rooms spacious and-windows open and-bedding of wool and-cloths of-goat-hair
- 1:2.421 1:2.422 1:2.423 1:2.424 1:2.425 1:2.426 1:2.427 | and also tables for the older children to learn and to write on them! | and-also tables for-the-children older to-learn and-to-write on-them
- 1:2.428 1:2.429 1:2.430 1:2.431 1:2.432 1:2.433 1:2.434 1:2.435 1:2.436 1:2.437 1:2.438 1:2.439 1:2.440 1:2.441 1:2.442 1:2.443 1:2.444 1:2.445 1:2.446 1:2.447 1:2.448 1:2.449 1:2.450 1:2.451 1:2.452 | Then we said: behold, the time has come to bring out Ashtonot Ish Yerushalayim, may it be rebuilt and established, which before she began to bear the great man in the revival of the Hebrew language, namely Eliezer Ben-Yehuda, | then we-said behold has-come the-time to-bring-out into-light Ashtonot the-man-of Jerusalem may-it-be-rebuilt-and-established which before she-began to-bear the-man the-great in-the-revival-of the-language Hebrew he namely Eliezer ben Yehuda
- 1:2.453 1:2.454 1:2.455 1:2.456 1:2.457 1:2.458 | and we acknowledge our debt to him for what he did and made, | and-we-acknowledge our-debt-to-him for what he-did and-made
- 1:2.459 1:2.460 1:2.461 1:2.462 1:2.463 1:2.464 1:2.465 1:2.466 1:2.467 1:2.468 1:2.469 1:2.470 1:2.471 1:2.472 1:2.473 1:2.474 1:2.475 1:2.476 1:2.477 1:2.478 1:2.479 1:2.480 1:2.481 | and with all his strength he toiled for the good of the people of Israel and the settlement of the land of Israel and the revival of the Hebrew language; this holy work he performed with wondrous alacrity and diligence; he made it succeed and raised it from the dust, | and-with-all his-strength he-toiled for-the-good-of the-people of-Israel and-the-settlement of-the-land of-Israel and-the-revival-of the-language Hebrew and-the-work-of the-holy this he-performed with-alacrity and-diligence wondrous and-he-made and-succeeded to-raise-it from-dust
- 1:2.482 1:2.483 1:2.484 1:2.485 1:2.486 | and after its decay it had a new bloom, | and-after its-decay it-had for-it a-new-bloom
- 1:2.487 1:2.488 1:2.489 1:2.490 1:2.491 1:2.492 1:2.493 1:2.494 1:2.495 1:2.496 | and because of his love for the holy tongue its wrinkles were smoothed out, and the appearance of its face showed itself and blossomed, | and-because-of his-love for-the-language holy were-smoothed-out its-wrinkles and-the-appearance-of its-face showed-itself and-blossomed
- 1:2.497 1:2.498 1:2.499 1:2.500 1:2.501 1:2.502 1:2.503 | and it found grace and honour in the eyes of its seekers, its knowers, and its acquaintances, | and-it-found grace and-honour in-the-eyes-of its-seekers its-knowers and-its-acquaintances
- 1:2.504 1:2.505 1:2.506 1:2.507 1:2.508 1:2.509 1:2.510 1:2.511 1:2.512 1:2.513 1:2.514 1:2.515 | and they awoke from their sleep to gird up their senses, to widen it and to enrich it as in the days of its youth, before the exile of the people of Israel! | and-they-awoke from-their-sleep to-gird-up their-senses to-widen-it and-to-enrich-it as-in-the-days-of its-youth before the-exile-of the-people-of Israel
- 1:2.516 1:2.517 1:2.518 | And because of this mitzvah, | and-because-of mitzvah this
- 1:2.519 1:2.520 1:2.521 1:2.522 1:2.523 | may his soul be bound up in the bundle of life | may-be his-soul bound in-the-bundle-of life
- 1:2.524 1:2.525 1:2.526 | with the LORD our God, | (object-marker) the-LORD our-God
- 1:2.527 1:2.528 1:2.529 1:2.530 1:2.531 1:2.532 1:2.533 | and from the fruit of his labour may his soul be satisfied in Eden, the garden of God! | and-from-the-fruit-of his-labour shall-be-satisfied his-soul in-Eden the-garden-of God

Grammar notes:
- יתבונן and ויתאונן are jussives (hitpa'el): 'let him reflect' and 'and let him lament', a wish addressed to the enlightened and the fearful, not a plain report.
- בראותו is a construct infinitive with a pronominal suffix, 'upon his seeing'; the clause that follows, אשר כל עם ועם ..., states what he sees.
- Rashi's quotation runs from כשהתינוק to תורה, introduced by שמכאן אמרו ('from here they said') and closed by עכ״ל. The words before it are the author's own.
- ולמדתם אותם את בניכם לדבר בם is Deut 11:19 in the second person plural, embedded in the first person plural of ולא עשינו: the author speaks for the community, and the verse's 'you' is the people it addresses.

Translator's note: This paragraph is in the author's own voice (the publisher's voice begins at 1:10). The Hebrew full stops are kept as the English sentence breaks, and the long periods are divided into sentences without dropping a clause.

Commentary rulings:
- apparatus: Deut 11:19, 'ולמדתם אותם את בניכם לדבר בם', embedded without quotation marks (tokens 75-80).
- apparatus: Exod 20:2, 'אנכי ה' אלהיך', the opening of the Decalogue, cited with the author's וכו' ('etc.'), so the English keeps 'etc.'.
- apparatus: 1 Sam 25:29, 'צרורה בצרור החיים את ה' אלהיך' ('the bundle of life'), with the author's אלהינו for אלהיך.
- apparatus: Ezek 31:9, 'גן האלהים' ('the garden of God'), compared with the author's 'בעדן גן אלהים'.
- apparatus: The Rashi reference is the author's: 'בפרשת עקב' (Parashat Ekev); no verse is given.
- apparatus: Dates: 5650 is about 1889-1890 and 5670 is about 1909-1910 of the civil calendar; the text gives no civil year.
- apparatus: Wikisource text: no collation with the 1931 scan has been made, so any Wikisource addition is translated as base text (Q2).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

