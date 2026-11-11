# Translation pack: milhamot-hashem perek-1-01 (Milchamot ha-Shem 1:1–1:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-01`.

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


## Milchamot ha-Shem 1:1

Hebrew (Wikisource): הקדמת המחבר.

| id | word | punctuation after |
|---|---|---|
| 1:1.1 | הקדמת |  |
| 1:1.2 | המחבר | . |


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


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

