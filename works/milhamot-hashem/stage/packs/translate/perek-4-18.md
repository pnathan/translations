# Translation pack: milhamot-hashem perek-4-18 (Milchamot ha-Shem 4:30–4:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-18`.

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
- רשב"י → "Rabbi Shimon bar Yochai", always written out. רשב"י התנא →
  "Rabbi Shimon bar Yochai the Tanna". The author separates the Tanna from
  the writer who "named himself" after him (1:4: "אשר כינה את עצמו בשם רשב״י
  התנא"). Keep exactly the qualifier he gives. Do not add "pseudo-" or
  "the so-called", and do not settle in the English which of the two he
  means.
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
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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

- אבות: lower-case "fathers", "patriarchs" in a token note. Never a title.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 4:30

Hebrew (Wikisource): לו) כבר ידעת ידידי הקורא מאמר רז"ל בפ"ק דשבת ובחגיגה פ' אין דורשין, אמר רב יהודה אמר רב ברם זכור אותו האיש לטוב וחנניה בן חזקיה שמו שאלמלא הוא נגנז ספר יחזקאל שהיו דבריו סותרין דברי תורה, מה עשה העלו לו שלש מאות גרבי שמן וישב בעליה ודרשן ע"כ, ופי' רש"י ז"ל סותרין דברי תורה כגון נבלה וטרפה לא יאכלו הכוהנים הא ישראל יאכלו? וכגון וכן תעשה בשבעה בחדש, היכן נרמז קרבן זה בתורה ודרשן כדדרשינן להני קראי במנחות דמשום דאישתרי מליקה לכהנים בחטאת העוף הוצרך להזהירן על מליקת חולין שהוא נבלה עכ"ל. ודון מינה ואוקי באתרין ומה ספר יחזקאל שכבר הוחזק לנביא אמת והחיה מתים בבקעת דורא וספרו מכלל ספרי הקדש ומצילין אותו מן הדליקה בשבת. ועם כל זה בקשו חכמים לגנזו מפני שהיה נראה כסותר דברי תורה באיסור לאו? כל שכן קל וחומר ס' הזהר שנראה וניכר בו סתירה לתורה שבכתב ושבעל פה באיסור ע"ז שהיא בסקילה וכרת. והוא ככופר בכל התורה כולה. ולא עוד אלא שהוא מבזה את המשנה והתלמוד ומכנה אותם בשם קליפה וסלע אחרא וכו' כמו שיתבאר לפנינו סי' ס"ז וס"ח. על אחת כמה וכמה שאסור לקרות בו כלל וכלל. כיון שכבר נתברר ונתפרסם זיופו לחכמים ונבונים ושדבריו בענין האלוהות סותרים דברי תורה ודברי סופרים המסורים לנו מרבותינו מעתיקי השמועה והגאונים והפוסקים שקבלו מהם. ועוד גרסינן בפ' במה מדליקין אמר רב יהודה בריה דרב יהודה בר שילת משמיה דרב ביקשו חכמים לגנוז ספר קהלת מפני שדבריו סותרים זה את זה וכו' באבות דרבי נתן בתחלה היו קורין משלי בלבד. ושיר השירים וקהלת היו גנוזים עד שבאו אנשי כנסת הגדולה ופירשום וקבעום בין הכתובים ע"כ הנה גם מזה תדין קל וחומר. ומה אם שלמה המלך ע"ה שנ' בו ויחכם מכל האדם. ובנה בית המקדש ושרתה בו שכינה ונראה אליו ה' כמבואר בכתוב. ואמרינן נמי בעירובין בשעה שתיקן עירובין ונטילת ידים. יצאה בת קול ואמרה בני אם חכם לבך ישמח לבי גם אני. ואומר חכם בני ושמח לבי ואשיבה חורפי דבר, ואמרינן נמי התם ואזן וחקר תיקן משלים הרבה, אמר ר"א בתחלה היתה תורה דומה לכפיפה שאין לה אזנים עד שבא שלמה ועשה לה אזנים! ועם כל השבח הזה לא נשאו לו חכמים פנים ובקשו לגנוז ספריו על שהיו דברים נראים סותרים זה את זה, עד שפירשום אנשי כנסת הגדולה! ספרים המבוארים שדבריהם סותרים זה את זה וגם סותרים דברי תורה ועוקרים יסודותיה. וגם לא נודע בבירור מי חברן. ומבטן מי יצאו הדברים הנאמרים בהם. ומיד בן נכר באו לידינו. רק שאותם אנשים שבאו לידם יחסו אותו לר' שמעון בן יוחאי התנא. ויש ידים מוכיחות דבשקר נתיחסו לרשב"י עאכ"ו שאסור לקרות בהן כלל.

| id | word | punctuation after |
|---|---|---|
| 4:30.1 | לו | ) |
| 4:30.2 | כבר |  |
| 4:30.3 | ידעת |  |
| 4:30.4 | ידידי |  |
| 4:30.5 | הקורא |  |
| 4:30.6 | מאמר |  |
| 4:30.7 | רז"ל |  |
| 4:30.8 | בפ"ק |  |
| 4:30.9 | דשבת |  |
| 4:30.10 | ובחגיגה |  |
| 4:30.11 | פ | ' |
| 4:30.12 | אין |  |
| 4:30.13 | דורשין | , |
| 4:30.14 | אמר |  |
| 4:30.15 | רב |  |
| 4:30.16 | יהודה |  |
| 4:30.17 | אמר |  |
| 4:30.18 | רב |  |
| 4:30.19 | ברם |  |
| 4:30.20 | זכור |  |
| 4:30.21 | אותו |  |
| 4:30.22 | האיש |  |
| 4:30.23 | לטוב |  |
| 4:30.24 | וחנניה |  |
| 4:30.25 | בן |  |
| 4:30.26 | חזקיה |  |
| 4:30.27 | שמו |  |
| 4:30.28 | שאלמלא |  |
| 4:30.29 | הוא |  |
| 4:30.30 | נגנז |  |
| 4:30.31 | ספר |  |
| 4:30.32 | יחזקאל |  |
| 4:30.33 | שהיו |  |
| 4:30.34 | דבריו |  |
| 4:30.35 | סותרין |  |
| 4:30.36 | דברי |  |
| 4:30.37 | תורה | , |
| 4:30.38 | מה |  |
| 4:30.39 | עשה |  |
| 4:30.40 | העלו |  |
| 4:30.41 | לו |  |
| 4:30.42 | שלש |  |
| 4:30.43 | מאות |  |
| 4:30.44 | גרבי |  |
| 4:30.45 | שמן |  |
| 4:30.46 | וישב |  |
| 4:30.47 | בעליה |  |
| 4:30.48 | ודרשן |  |
| 4:30.49 | ע"כ | , |
| 4:30.50 | ופי | ' |
| 4:30.51 | רש"י |  |
| 4:30.52 | ז"ל |  |
| 4:30.53 | סותרין |  |
| 4:30.54 | דברי |  |
| 4:30.55 | תורה |  |
| 4:30.56 | כגון |  |
| 4:30.57 | נבלה |  |
| 4:30.58 | וטרפה |  |
| 4:30.59 | לא |  |
| 4:30.60 | יאכלו |  |
| 4:30.61 | הכוהנים |  |
| 4:30.62 | הא |  |
| 4:30.63 | ישראל |  |
| 4:30.64 | יאכלו | ? |
| 4:30.65 | וכגון |  |
| 4:30.66 | וכן |  |
| 4:30.67 | תעשה |  |
| 4:30.68 | בשבעה |  |
| 4:30.69 | בחדש | , |
| 4:30.70 | היכן |  |
| 4:30.71 | נרמז |  |
| 4:30.72 | קרבן |  |
| 4:30.73 | זה |  |
| 4:30.74 | בתורה |  |
| 4:30.75 | ודרשן |  |
| 4:30.76 | כדדרשינן |  |
| 4:30.77 | להני |  |
| 4:30.78 | קראי |  |
| 4:30.79 | במנחות |  |
| 4:30.80 | דמשום |  |
| 4:30.81 | דאישתרי |  |
| 4:30.82 | מליקה |  |
| 4:30.83 | לכהנים |  |
| 4:30.84 | בחטאת |  |
| 4:30.85 | העוף |  |
| 4:30.86 | הוצרך |  |
| 4:30.87 | להזהירן |  |
| 4:30.88 | על |  |
| 4:30.89 | מליקת |  |
| 4:30.90 | חולין |  |
| 4:30.91 | שהוא |  |
| 4:30.92 | נבלה |  |
| 4:30.93 | עכ"ל | . |
| 4:30.94 | ודון |  |
| 4:30.95 | מינה |  |
| 4:30.96 | ואוקי |  |
| 4:30.97 | באתרין |  |
| 4:30.98 | ומה |  |
| 4:30.99 | ספר |  |
| 4:30.100 | יחזקאל |  |
| 4:30.101 | שכבר |  |
| 4:30.102 | הוחזק |  |
| 4:30.103 | לנביא |  |
| 4:30.104 | אמת |  |
| 4:30.105 | והחיה |  |
| 4:30.106 | מתים |  |
| 4:30.107 | בבקעת |  |
| 4:30.108 | דורא |  |
| 4:30.109 | וספרו |  |
| 4:30.110 | מכלל |  |
| 4:30.111 | ספרי |  |
| 4:30.112 | הקדש |  |
| 4:30.113 | ומצילין |  |
| 4:30.114 | אותו |  |
| 4:30.115 | מן |  |
| 4:30.116 | הדליקה |  |
| 4:30.117 | בשבת | . |
| 4:30.118 | ועם |  |
| 4:30.119 | כל |  |
| 4:30.120 | זה |  |
| 4:30.121 | בקשו |  |
| 4:30.122 | חכמים |  |
| 4:30.123 | לגנזו |  |
| 4:30.124 | מפני |  |
| 4:30.125 | שהיה |  |
| 4:30.126 | נראה |  |
| 4:30.127 | כסותר |  |
| 4:30.128 | דברי |  |
| 4:30.129 | תורה |  |
| 4:30.130 | באיסור |  |
| 4:30.131 | לאו | ? |
| 4:30.132 | כל |  |
| 4:30.133 | שכן |  |
| 4:30.134 | קל |  |
| 4:30.135 | וחומר |  |
| 4:30.136 | ס | ' |
| 4:30.137 | הזהר |  |
| 4:30.138 | שנראה |  |
| 4:30.139 | וניכר |  |
| 4:30.140 | בו |  |
| 4:30.141 | סתירה |  |
| 4:30.142 | לתורה |  |
| 4:30.143 | שבכתב |  |
| 4:30.144 | ושבעל |  |
| 4:30.145 | פה |  |
| 4:30.146 | באיסור |  |
| 4:30.147 | ע"ז |  |
| 4:30.148 | שהיא |  |
| 4:30.149 | בסקילה |  |
| 4:30.150 | וכרת | . |
| 4:30.151 | והוא |  |
| 4:30.152 | ככופר |  |
| 4:30.153 | בכל |  |
| 4:30.154 | התורה |  |
| 4:30.155 | כולה | . |
| 4:30.156 | ולא |  |
| 4:30.157 | עוד |  |
| 4:30.158 | אלא |  |
| 4:30.159 | שהוא |  |
| 4:30.160 | מבזה |  |
| 4:30.161 | את |  |
| 4:30.162 | המשנה |  |
| 4:30.163 | והתלמוד |  |
| 4:30.164 | ומכנה |  |
| 4:30.165 | אותם |  |
| 4:30.166 | בשם |  |
| 4:30.167 | קליפה |  |
| 4:30.168 | וסלע |  |
| 4:30.169 | אחרא |  |
| 4:30.170 | וכו | ' |
| 4:30.171 | כמו |  |
| 4:30.172 | שיתבאר |  |
| 4:30.173 | לפנינו |  |
| 4:30.174 | סי | ' |
| 4:30.175 | ס"ז |  |
| 4:30.176 | וס"ח | . |
| 4:30.177 | על |  |
| 4:30.178 | אחת |  |
| 4:30.179 | כמה |  |
| 4:30.180 | וכמה |  |
| 4:30.181 | שאסור |  |
| 4:30.182 | לקרות |  |
| 4:30.183 | בו |  |
| 4:30.184 | כלל |  |
| 4:30.185 | וכלל | . |
| 4:30.186 | כיון |  |
| 4:30.187 | שכבר |  |
| 4:30.188 | נתברר |  |
| 4:30.189 | ונתפרסם |  |
| 4:30.190 | זיופו |  |
| 4:30.191 | לחכמים |  |
| 4:30.192 | ונבונים |  |
| 4:30.193 | ושדבריו |  |
| 4:30.194 | בענין |  |
| 4:30.195 | האלוהות |  |
| 4:30.196 | סותרים |  |
| 4:30.197 | דברי |  |
| 4:30.198 | תורה |  |
| 4:30.199 | ודברי |  |
| 4:30.200 | סופרים |  |
| 4:30.201 | המסורים |  |
| 4:30.202 | לנו |  |
| 4:30.203 | מרבותינו |  |
| 4:30.204 | מעתיקי |  |
| 4:30.205 | השמועה |  |
| 4:30.206 | והגאונים |  |
| 4:30.207 | והפוסקים |  |
| 4:30.208 | שקבלו |  |
| 4:30.209 | מהם | . |
| 4:30.210 | ועוד |  |
| 4:30.211 | גרסינן |  |
| 4:30.212 | בפ | ' |
| 4:30.213 | במה |  |
| 4:30.214 | מדליקין |  |
| 4:30.215 | אמר |  |
| 4:30.216 | רב |  |
| 4:30.217 | יהודה |  |
| 4:30.218 | בריה |  |
| 4:30.219 | דרב |  |
| 4:30.220 | יהודה |  |
| 4:30.221 | בר |  |
| 4:30.222 | שילת |  |
| 4:30.223 | משמיה |  |
| 4:30.224 | דרב |  |
| 4:30.225 | ביקשו |  |
| 4:30.226 | חכמים |  |
| 4:30.227 | לגנוז |  |
| 4:30.228 | ספר |  |
| 4:30.229 | קהלת |  |
| 4:30.230 | מפני |  |
| 4:30.231 | שדבריו |  |
| 4:30.232 | סותרים |  |
| 4:30.233 | זה |  |
| 4:30.234 | את |  |
| 4:30.235 | זה |  |
| 4:30.236 | וכו | ' |
| 4:30.237 | באבות |  |
| 4:30.238 | דרבי |  |
| 4:30.239 | נתן |  |
| 4:30.240 | בתחלה |  |
| 4:30.241 | היו |  |
| 4:30.242 | קורין |  |
| 4:30.243 | משלי |  |
| 4:30.244 | בלבד | . |
| 4:30.245 | ושיר |  |
| 4:30.246 | השירים |  |
| 4:30.247 | וקהלת |  |
| 4:30.248 | היו |  |
| 4:30.249 | גנוזים |  |
| 4:30.250 | עד |  |
| 4:30.251 | שבאו |  |
| 4:30.252 | אנשי |  |
| 4:30.253 | כנסת |  |
| 4:30.254 | הגדולה |  |
| 4:30.255 | ופירשום |  |
| 4:30.256 | וקבעום |  |
| 4:30.257 | בין |  |
| 4:30.258 | הכתובים |  |
| 4:30.259 | ע"כ |  |
| 4:30.260 | הנה |  |
| 4:30.261 | גם |  |
| 4:30.262 | מזה |  |
| 4:30.263 | תדין |  |
| 4:30.264 | קל |  |
| 4:30.265 | וחומר | . |
| 4:30.266 | ומה |  |
| 4:30.267 | אם |  |
| 4:30.268 | שלמה |  |
| 4:30.269 | המלך |  |
| 4:30.270 | ע"ה |  |
| 4:30.271 | שנ | ' |
| 4:30.272 | בו |  |
| 4:30.273 | ויחכם |  |
| 4:30.274 | מכל |  |
| 4:30.275 | האדם | . |
| 4:30.276 | ובנה |  |
| 4:30.277 | בית |  |
| 4:30.278 | המקדש |  |
| 4:30.279 | ושרתה |  |
| 4:30.280 | בו |  |
| 4:30.281 | שכינה |  |
| 4:30.282 | ונראה |  |
| 4:30.283 | אליו |  |
| 4:30.284 | ה | ' |
| 4:30.285 | כמבואר |  |
| 4:30.286 | בכתוב | . |
| 4:30.287 | ואמרינן |  |
| 4:30.288 | נמי |  |
| 4:30.289 | בעירובין |  |
| 4:30.290 | בשעה |  |
| 4:30.291 | שתיקן |  |
| 4:30.292 | עירובין |  |
| 4:30.293 | ונטילת |  |
| 4:30.294 | ידים | . |
| 4:30.295 | יצאה |  |
| 4:30.296 | בת |  |
| 4:30.297 | קול |  |
| 4:30.298 | ואמרה |  |
| 4:30.299 | בני |  |
| 4:30.300 | אם |  |
| 4:30.301 | חכם |  |
| 4:30.302 | לבך |  |
| 4:30.303 | ישמח |  |
| 4:30.304 | לבי |  |
| 4:30.305 | גם |  |
| 4:30.306 | אני | . |
| 4:30.307 | ואומר |  |
| 4:30.308 | חכם |  |
| 4:30.309 | בני |  |
| 4:30.310 | ושמח |  |
| 4:30.311 | לבי |  |
| 4:30.312 | ואשיבה |  |
| 4:30.313 | חורפי |  |
| 4:30.314 | דבר | , |
| 4:30.315 | ואמרינן |  |
| 4:30.316 | נמי |  |
| 4:30.317 | התם |  |
| 4:30.318 | ואזן |  |
| 4:30.319 | וחקר |  |
| 4:30.320 | תיקן |  |
| 4:30.321 | משלים |  |
| 4:30.322 | הרבה | , |
| 4:30.323 | אמר |  |
| 4:30.324 | ר"א |  |
| 4:30.325 | בתחלה |  |
| 4:30.326 | היתה |  |
| 4:30.327 | תורה |  |
| 4:30.328 | דומה |  |
| 4:30.329 | לכפיפה |  |
| 4:30.330 | שאין |  |
| 4:30.331 | לה |  |
| 4:30.332 | אזנים |  |
| 4:30.333 | עד |  |
| 4:30.334 | שבא |  |
| 4:30.335 | שלמה |  |
| 4:30.336 | ועשה |  |
| 4:30.337 | לה |  |
| 4:30.338 | אזנים | ! |
| 4:30.339 | ועם |  |
| 4:30.340 | כל |  |
| 4:30.341 | השבח |  |
| 4:30.342 | הזה |  |
| 4:30.343 | לא |  |
| 4:30.344 | נשאו |  |
| 4:30.345 | לו |  |
| 4:30.346 | חכמים |  |
| 4:30.347 | פנים |  |
| 4:30.348 | ובקשו |  |
| 4:30.349 | לגנוז |  |
| 4:30.350 | ספריו |  |
| 4:30.351 | על |  |
| 4:30.352 | שהיו |  |
| 4:30.353 | דברים |  |
| 4:30.354 | נראים |  |
| 4:30.355 | סותרים |  |
| 4:30.356 | זה |  |
| 4:30.357 | את |  |
| 4:30.358 | זה | , |
| 4:30.359 | עד |  |
| 4:30.360 | שפירשום |  |
| 4:30.361 | אנשי |  |
| 4:30.362 | כנסת |  |
| 4:30.363 | הגדולה | ! |
| 4:30.364 | ספרים |  |
| 4:30.365 | המבוארים |  |
| 4:30.366 | שדבריהם |  |
| 4:30.367 | סותרים |  |
| 4:30.368 | זה |  |
| 4:30.369 | את |  |
| 4:30.370 | זה |  |
| 4:30.371 | וגם |  |
| 4:30.372 | סותרים |  |
| 4:30.373 | דברי |  |
| 4:30.374 | תורה |  |
| 4:30.375 | ועוקרים |  |
| 4:30.376 | יסודותיה | . |
| 4:30.377 | וגם |  |
| 4:30.378 | לא |  |
| 4:30.379 | נודע |  |
| 4:30.380 | בבירור |  |
| 4:30.381 | מי |  |
| 4:30.382 | חברן | . |
| 4:30.383 | ומבטן |  |
| 4:30.384 | מי |  |
| 4:30.385 | יצאו |  |
| 4:30.386 | הדברים |  |
| 4:30.387 | הנאמרים |  |
| 4:30.388 | בהם | . |
| 4:30.389 | ומיד |  |
| 4:30.390 | בן |  |
| 4:30.391 | נכר |  |
| 4:30.392 | באו |  |
| 4:30.393 | לידינו | . |
| 4:30.394 | רק |  |
| 4:30.395 | שאותם |  |
| 4:30.396 | אנשים |  |
| 4:30.397 | שבאו |  |
| 4:30.398 | לידם |  |
| 4:30.399 | יחסו |  |
| 4:30.400 | אותו |  |
| 4:30.401 | לר | ' |
| 4:30.402 | שמעון |  |
| 4:30.403 | בן |  |
| 4:30.404 | יוחאי |  |
| 4:30.405 | התנא | . |
| 4:30.406 | ויש |  |
| 4:30.407 | ידים |  |
| 4:30.408 | מוכיחות |  |
| 4:30.409 | דבשקר |  |
| 4:30.410 | נתיחסו |  |
| 4:30.411 | לרשב"י |  |
| 4:30.412 | עאכ"ו |  |
| 4:30.413 | שאסור |  |
| 4:30.414 | לקרות |  |
| 4:30.415 | בהן |  |
| 4:30.416 | כלל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

