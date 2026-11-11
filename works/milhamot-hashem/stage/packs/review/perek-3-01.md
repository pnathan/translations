# Review pack: milhamot-hashem perek-3-01 (Milchamot ha-Shem 3:1–3:2)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-01`.

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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
- Fixed forms: "Roke'ach" (the Rokeach); "Ein Sof" never "without end" for the
  noun (but לאין סוף "without end" in the Rambam passage when it is not the
  name).
- ר"ל with a named authority as subject: "means to say".
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
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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


## Milchamot ha-Shem 3:1

Hebrew (Wikisource): ב) וכ"ש שלא נאמין לו באמרו שנגלה אליו אליהו או סבא דסבין, שהוא עתיקא קדישא הנקרא כתר עליון, דודאי נביא שקר הוא ובמיתתו בחנק. כמבואה בדברי הרמב"ם בפי' המשנה בהקדמתו לסדר זרעים. ובגמ' פסחים (לד) גרס' אמתר הסח הדעת ר' יוחנן אמר פסול טומאה, שאם יבוא אליהו ויטהרנה ויאמר שלא נגעה טומאה בהן שומעין לו, ור"ל אמר פסול הגוף הוי, שאם יבוא אליהו ויטהרנה אין שומעין לו עכ"ל, ופי' רשי ז"ל פיסול הגוף מעלה היא בקדשים היא שנפסלין בכך, ושם פיסול הוא לעצמו, ואפי' אם יבוא אליהו ויאמר שלא נטמאו כל רותן הימים שהסיח דעתו מהן אין שומעין לו ע"כ, אלמלא כל הבא לשנות ולהתיר את האסור, אפי' איסור דרבנן אין שומעין לו.

| id | word | punctuation after |
|---|---|---|
| 3:1.1 | ב | ) |
| 3:1.2 | וכ"ש |  |
| 3:1.3 | שלא |  |
| 3:1.4 | נאמין |  |
| 3:1.5 | לו |  |
| 3:1.6 | באמרו |  |
| 3:1.7 | שנגלה |  |
| 3:1.8 | אליו |  |
| 3:1.9 | אליהו |  |
| 3:1.10 | או |  |
| 3:1.11 | סבא |  |
| 3:1.12 | דסבין | , |
| 3:1.13 | שהוא |  |
| 3:1.14 | עתיקא |  |
| 3:1.15 | קדישא |  |
| 3:1.16 | הנקרא |  |
| 3:1.17 | כתר |  |
| 3:1.18 | עליון | , |
| 3:1.19 | דודאי |  |
| 3:1.20 | נביא |  |
| 3:1.21 | שקר |  |
| 3:1.22 | הוא |  |
| 3:1.23 | ובמיתתו |  |
| 3:1.24 | בחנק | . |
| 3:1.25 | כמבואה |  |
| 3:1.26 | בדברי |  |
| 3:1.27 | הרמב"ם |  |
| 3:1.28 | בפי | ' |
| 3:1.29 | המשנה |  |
| 3:1.30 | בהקדמתו |  |
| 3:1.31 | לסדר |  |
| 3:1.32 | זרעים | . |
| 3:1.33 | ובגמ | ' |
| 3:1.34 | פסחים |  |
| 3:1.35 | לד | ) |
| 3:1.36 | גרס | ' |
| 3:1.37 | אמתר |  |
| 3:1.38 | הסח |  |
| 3:1.39 | הדעת |  |
| 3:1.40 | ר | ' |
| 3:1.41 | יוחנן |  |
| 3:1.42 | אמר |  |
| 3:1.43 | פסול |  |
| 3:1.44 | טומאה | , |
| 3:1.45 | שאם |  |
| 3:1.46 | יבוא |  |
| 3:1.47 | אליהו |  |
| 3:1.48 | ויטהרנה |  |
| 3:1.49 | ויאמר |  |
| 3:1.50 | שלא |  |
| 3:1.51 | נגעה |  |
| 3:1.52 | טומאה |  |
| 3:1.53 | בהן |  |
| 3:1.54 | שומעין |  |
| 3:1.55 | לו | , |
| 3:1.56 | ור"ל |  |
| 3:1.57 | אמר |  |
| 3:1.58 | פסול |  |
| 3:1.59 | הגוף |  |
| 3:1.60 | הוי | , |
| 3:1.61 | שאם |  |
| 3:1.62 | יבוא |  |
| 3:1.63 | אליהו |  |
| 3:1.64 | ויטהרנה |  |
| 3:1.65 | אין |  |
| 3:1.66 | שומעין |  |
| 3:1.67 | לו |  |
| 3:1.68 | עכ"ל | , |
| 3:1.69 | ופי | ' |
| 3:1.70 | רשי |  |
| 3:1.71 | ז"ל |  |
| 3:1.72 | פיסול |  |
| 3:1.73 | הגוף |  |
| 3:1.74 | מעלה |  |
| 3:1.75 | היא |  |
| 3:1.76 | בקדשים |  |
| 3:1.77 | היא |  |
| 3:1.78 | שנפסלין |  |
| 3:1.79 | בכך | , |
| 3:1.80 | ושם |  |
| 3:1.81 | פיסול |  |
| 3:1.82 | הוא |  |
| 3:1.83 | לעצמו | , |
| 3:1.84 | ואפי | ' |
| 3:1.85 | אם |  |
| 3:1.86 | יבוא |  |
| 3:1.87 | אליהו |  |
| 3:1.88 | ויאמר |  |
| 3:1.89 | שלא |  |
| 3:1.90 | נטמאו |  |
| 3:1.91 | כל |  |
| 3:1.92 | רותן |  |
| 3:1.93 | הימים |  |
| 3:1.94 | שהסיח |  |
| 3:1.95 | דעתו |  |
| 3:1.96 | מהן |  |
| 3:1.97 | אין |  |
| 3:1.98 | שומעין |  |
| 3:1.99 | לו |  |
| 3:1.100 | ע"כ | , |
| 3:1.101 | אלמלא |  |
| 3:1.102 | כל |  |
| 3:1.103 | הבא |  |
| 3:1.104 | לשנות |  |
| 3:1.105 | ולהתיר |  |
| 3:1.106 | את |  |
| 3:1.107 | האסור | , |
| 3:1.108 | אפי | ' |
| 3:1.109 | איסור |  |
| 3:1.110 | דרבנן |  |
| 3:1.111 | אין |  |
| 3:1.112 | שומעין |  |
| 3:1.113 | לו | . |


### Translation under review (3:1)

English: §2. And all the more so we should not believe him when he says that Elijah was revealed to him, or the Elder of Elders, who is Atika Kadisha, which is called Keter Elyon (the Supreme Crown), for surely he is a false prophet, and he is put to death by strangling. As is shown in the words of the Rambam in his commentary on the Mishnah, in his introduction to the Order of Zera'im. And in the Gemara, Pesachim 34, it reads: "Amtar, the mind was distracted. Rabbi Yochanan said: a disqualification by impurity, that if Elijah should come and declare it clean and say that impurity had not touched them, we listen to him, and Resh Lakish said: a disqualification of the body, is that if Elijah should come and declare it clean, we do not listen to him." Thus far his words. And Rashi, of blessed memory, explains: "bodily disqualification is a merit in sacred offerings, in that they are disqualified by it; and the disqualification there is in itself; even if Elijah should come and say that all their [unclear] days were not defiled, since his mind was distracted from them, we do not listen to him." Thus far. Indeed, everyone who comes to change and to permit the forbidden, even a rabbinic prohibition, we do not listen to him.

Literal: section 2 and all the more so that not we believe to him in his saying that he was revealed to him Elijah or the Elder of Elders who is Atika Kadisha which is called Keter Elyon surely prophet false he and by his death by strangling as is shown in the words of the Rambam in his commentary on the Mishnah in his introduction to the Order Zera'im and in the Gemara Pesachim 34 it reads Amtar was distracted the mind Rabbi Yochanan said disqualification by impurity that if should come Elijah and declare it clean and say that not touched impurity them we listen to him and Resh Lakish said disqualification of the body is that if should come Elijah and declare it clean we do not listen to him thus far his words and explains Rashi of blessed memory disqualification of the body a merit is in sacred offerings that they are disqualified by it and there disqualification is in itself and even if should come Elijah and say that not were defiled all their [unclear] days since he distracted his mind from them we do not listen to him thus far indeed (literally: were it not) everyone who comes to change and to permit (object marker) the forbidden even a prohibition rabbinic we do not listen to him

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:1.1 | section 2 |  |  | the label ב is the numeral 2 (section 2) |
| 3:1.2 | and all the more so |  |  | כ"ש = כל שכן, 'all the more so' |
| 3:1.3 | that not |  |  |  |
| 3:1.4 | we believe |  |  |  |
| 3:1.5 | to him |  |  |  |
| 3:1.6 | in his saying |  |  |  |
| 3:1.7 | that he was revealed |  |  |  |
| 3:1.8 | to him |  |  |  |
| 3:1.9 | Elijah |  |  |  |
| 3:1.10 | or |  |  |  |
| 3:1.11 | the Elder |  |  | Aramaic סבא, 'elder'; the title is the author's, not a name |
| 3:1.12 | of Elders |  |  |  |
| 3:1.13 | who is |  |  |  |
| 3:1.14 | Atika |  |  | Aramaic 'holy ancient one'; the Hebrew title is kept |
| 3:1.15 | Kadisha |  |  |  |
| 3:1.16 | which is called |  |  |  |
| 3:1.17 | Keter |  |  | כתר עליון, 'Supreme Crown' |
| 3:1.18 | Elyon |  |  |  |
| 3:1.19 | surely |  |  |  |
| 3:1.20 | prophet |  |  |  |
| 3:1.21 | false |  |  |  |
| 3:1.22 | he |  |  |  |
| 3:1.23 | and by his death |  |  |  |
| 3:1.24 | by strangling |  |  |  |
| 3:1.25 | as is shown |  |  |  |
| 3:1.26 | in the words of |  |  |  |
| 3:1.27 | the Rambam |  |  |  |
| 3:1.28 | in his commentary |  |  |  |
| 3:1.29 | on the Mishnah |  |  |  |
| 3:1.30 | in his introduction |  |  |  |
| 3:1.31 | to the Order |  |  |  |
| 3:1.32 | Zera'im |  |  |  |
| 3:1.33 | and in the Gemara |  |  |  |
| 3:1.34 | Pesachim |  |  |  |
| 3:1.35 | 34 |  |  |  |
| 3:1.36 | it reads |  |  |  |
| 3:1.37 | Amtar |  |  | amtar: Aramaic word, sense uncertain in this place; see apparatus |
| 3:1.38 | was distracted |  |  |  |
| 3:1.39 | the mind |  |  |  |
| 3:1.40 | Rabbi |  |  |  |
| 3:1.41 | Yochanan |  |  |  |
| 3:1.42 | said |  |  |  |
| 3:1.43 | disqualification |  |  |  |
| 3:1.44 | by impurity |  |  |  |
| 3:1.45 | that if |  |  |  |
| 3:1.46 | should come |  |  |  |
| 3:1.47 | Elijah |  |  |  |
| 3:1.48 | and declare it clean |  |  |  |
| 3:1.49 | and say |  |  |  |
| 3:1.50 | that not |  |  |  |
| 3:1.51 | touched |  |  |  |
| 3:1.52 | impurity |  |  |  |
| 3:1.53 | them |  |  |  |
| 3:1.54 | we listen |  |  |  |
| 3:1.55 | to him |  |  |  |
| 3:1.56 | and Resh Lakish |  |  | ר"ל = Resh Lakish |
| 3:1.57 | said |  |  |  |
| 3:1.58 | disqualification |  |  |  |
| 3:1.59 | of the body |  |  |  |
| 3:1.60 | is |  |  |  |
| 3:1.61 | that if |  |  |  |
| 3:1.62 | should come |  |  |  |
| 3:1.63 | Elijah |  |  |  |
| 3:1.64 | and declare it clean |  |  |  |
| 3:1.65 | we do not |  |  |  |
| 3:1.66 | listen |  |  |  |
| 3:1.67 | to him |  |  |  |
| 3:1.68 | thus far his words |  |  |  |
| 3:1.69 | and explains |  |  |  |
| 3:1.70 | Rashi |  |  |  |
| 3:1.71 | of blessed memory |  |  |  |
| 3:1.72 | disqualification |  |  |  |
| 3:1.73 | of the body |  |  |  |
| 3:1.74 | a merit |  |  |  |
| 3:1.75 | is |  |  |  |
| 3:1.76 | in sacred offerings |  |  |  |
| 3:1.77 | that |  |  |  |
| 3:1.78 | they are disqualified |  |  |  |
| 3:1.79 | by it |  |  |  |
| 3:1.80 | and there |  |  |  |
| 3:1.81 | disqualification |  |  |  |
| 3:1.82 | is |  |  |  |
| 3:1.83 | in itself |  |  |  |
| 3:1.84 | and even |  |  |  |
| 3:1.85 | if |  |  |  |
| 3:1.86 | should come |  |  |  |
| 3:1.87 | Elijah |  |  |  |
| 3:1.88 | and say |  |  |  |
| 3:1.89 | that not |  |  |  |
| 3:1.90 | were defiled |  |  |  |
| 3:1.91 | all |  |  |  |
| 3:1.92 | their [unclear] |  |  | word not clear to me in this sense; see apparatus |
| 3:1.93 | days |  |  |  |
| 3:1.94 | since he distracted |  |  |  |
| 3:1.95 | his mind |  |  |  |
| 3:1.96 | from them |  |  |  |
| 3:1.97 | we do not |  |  |  |
| 3:1.98 | listen |  |  |  |
| 3:1.99 | to him |  |  |  |
| 3:1.100 | thus far |  |  |  |
| 3:1.101 | indeed (literally: were it not) |  |  | literally 'were it not'; taken as the opening of the author's own conclusion |
| 3:1.102 | everyone |  |  |  |
| 3:1.103 | who comes |  |  |  |
| 3:1.104 | to change |  |  |  |
| 3:1.105 | and to permit |  |  |  |
| 3:1.106 | (object marker) |  |  | את: object marker, no English counterpart |
| 3:1.107 | the forbidden |  |  |  |
| 3:1.108 | even |  |  |  |
| 3:1.109 | a prohibition |  |  |  |
| 3:1.110 | rabbinic |  |  |  |
| 3:1.111 | we do not |  |  |  |
| 3:1.112 | listen |  |  |  |
| 3:1.113 | to him |  |  |  |

Phrases:
- 3:1.1 3:1.2 | §2. And all the more so | section 2 and all the more so
- 3:1.3 3:1.4 3:1.5 3:1.6 3:1.7 3:1.8 3:1.9 | we should not believe him when he says that Elijah was revealed to him, | that not we believe to him in his saying that he was revealed to him Elijah
- 3:1.10 3:1.11 3:1.12 | or the Elder of Elders, | or the Elder of Elders
- 3:1.13 3:1.14 3:1.15 3:1.16 | who is Atika Kadisha, which is called | who is Atika Kadisha which is called
- 3:1.17 3:1.18 | Keter Elyon (the Supreme Crown), | Keter Elyon
- 3:1.19 3:1.20 3:1.21 3:1.22 | for surely he is a false prophet, | surely prophet false he
- 3:1.23 3:1.24 | and he is put to death by strangling. | and by his death by strangling
- 3:1.25 3:1.26 | As is shown in the words of | as is shown in the words of
- 3:1.27 | the Rambam | the Rambam
- 3:1.28 3:1.29 3:1.30 3:1.31 3:1.32 | in his commentary on the Mishnah, in his introduction to the Order of Zera'im. | in his commentary on the Mishnah in his introduction to the Order Zera'im
- 3:1.33 3:1.34 3:1.35 3:1.36 | And in the Gemara, Pesachim 34, it reads: | and in the Gemara Pesachim 34 it reads
- 3:1.37 3:1.38 3:1.39 | "Amtar, the mind was distracted. | Amtar was distracted the mind
- 3:1.40 3:1.41 3:1.42 | Rabbi Yochanan said: | Rabbi Yochanan said
- 3:1.43 3:1.44 | a disqualification by impurity, | disqualification by impurity
- 3:1.45 3:1.46 3:1.47 3:1.48 3:1.49 3:1.50 3:1.51 3:1.52 3:1.53 3:1.54 3:1.55 | that if Elijah should come and declare it clean and say that impurity had not touched them, we listen to him, | that if should come Elijah and declare it clean and say that not touched impurity them we listen to him
- 3:1.56 3:1.57 | and Resh Lakish said: | and Resh Lakish said
- 3:1.58 3:1.59 | a disqualification of the body, | disqualification of the body
- 3:1.60 3:1.61 3:1.62 3:1.63 3:1.64 3:1.65 3:1.66 3:1.67 3:1.68 | is that if Elijah should come and declare it clean, we do not listen to him." Thus far his words. | is that if should come Elijah and declare it clean we do not listen to him thus far his words
- 3:1.69 3:1.70 3:1.71 | And Rashi, of blessed memory, explains: | and explains Rashi of blessed memory
- 3:1.72 3:1.73 3:1.74 3:1.75 3:1.76 3:1.77 3:1.78 3:1.79 | "bodily disqualification is a merit in sacred offerings, in that they are disqualified by it; | disqualification of the body a merit is in sacred offerings that they are disqualified by it
- 3:1.80 3:1.81 3:1.82 3:1.83 | and the disqualification there is in itself; | and there disqualification is in itself
- 3:1.84 3:1.85 3:1.86 3:1.87 3:1.88 3:1.89 3:1.90 3:1.91 3:1.92 3:1.93 3:1.94 3:1.95 3:1.96 3:1.97 3:1.98 3:1.99 3:1.100 | even if Elijah should come and say that all their [unclear] days were not defiled, since his mind was distracted from them, we do not listen to him." Thus far. | and even if should come Elijah and say that not were defiled all their [unclear] days since he distracted his mind from them we do not listen to him thus far
- 3:1.101 3:1.102 3:1.103 3:1.104 3:1.105 3:1.106 3:1.107 3:1.108 3:1.109 3:1.110 3:1.111 3:1.112 3:1.113 | Indeed, everyone who comes to change and to permit the forbidden, even a rabbinic prohibition, we do not listen to him. | indeed (literally: were it not) everyone who comes to change and to permit (object marker) the forbidden even a prohibition rabbinic we do not listen to him

Grammar notes:
- שלא נאמין לו: 'that we should not believe him', a clause with שלא and the imperfect after כ"ש, 'all the more so'.
- סבא דסבין is Aramaic: 'the Elder of Elders', with the Aramaic relative ד and the title given as a name-like epithet.
- The Talmud passage is quoted with 'גרס', 'it reads'; the closing 'עכ"ל' ends it, and 'ע"כ' ends Rashi's words after it.
- ויטהרנה: the verb with the object suffix -נה; the English 'declare it clean' takes the suffix as 'it'.

Translator's note: The paragraph has three parts: the refusal to believe the claimed revelation, a Talmud passage with Resh Lakish's view, and Rashi's explanation with the conclusion. The Hebrew is the Wikisource base text.

Commentary rulings:
- apparatus: Base text: Wikisource edition, translated as base text and flagged until the 1931 scan is collated (decisions, Q2).
- apparatus: Rambam: his commentary on the Mishnah, introduction to Zera'im, as the author gives it (no chapter or page).
- apparatus: Pesachim (לד): the author's '34' as written; the quotation as he gives it, not checked against the Talmud text.
- apparatus: amtar: Aramaic word of uncertain sense here; the English reads it as the word the Gemara uses to introduce the case, and the alternative is to leave it untranslated.
- apparatus: רותן: word not established in this sense; translated 'their [unclear]' and marked in the token notes.

## Milchamot ha-Shem 3:2

Hebrew (Wikisource): ג) ובאמת שהחלת דברי פיך סכלות והוללות רעה שאמרת איך מלאני לבי לדרוש ולחקור וכו' ולמה לא? אדרבא איפכא מסתברא (כיון שאומרים לעבוד לז"א שהוא נברא יושר לבב דף ג' ע"ב ועוז לאלוהים). חוב גמור עלינו ועל כל אחד ואחד בישראל לדעת את ה' שהוא בורא ולא נברא ולעבדו ככתוב וידעת היום והשבות אל לבבך. וידעת כי ה' אלהיך הוא האלהים וכו' ואומר דע את אלהי אביך ועבדהו. והתירו חכמים לאפרושי מאיסורא אפי' בבית הכסא ובבית המרחץ שאסור להרהר בהם בדברי תורה. ולאפרושי מאיסורא שרי. ועל תמיהתך איך יטעו? אין זה פלא. ולא עדיפי כת המקובלים האחרונים אשר הלכו לתומם אחרי צו של הפילוסוף מחבר הזהר שתלאו ברשב"י ור' אלעזר בנו וחביריו מסנהדרי גדולה היושבים בלשכת הגזית שעליהם אמר הכתוב כי יפלא ממך דבר למשפט וכו' וקמת ועלית אל המקום אשר יבחר ה' אלוהיך וכו' ועשית על פי הדבר אשר יגידו לך וכו'. ועם כל זה אמרה תורה שאם שגגו ב"ד הגדול בהוראה ועשו העם על פיהם מביאים קרבן על שגגם. אם בע"ז שגגו והורו מביאים פר לעולה ושעיר לחטאת מכל שבט ושבט. והוא האמור בפרשת שלח והיה אם מעיני העדה נעשתה לשגגה, ואם בשאר כריתות שגגו והורו מביא כל שבט ושבט פר לחטאת והוא האמור בפרשת ויקרא. ואם כל עדת ישראל ישגו ונעלם דבר מעיני הקהל.

| id | word | punctuation after |
|---|---|---|
| 3:2.1 | ג | ) |
| 3:2.2 | ובאמת |  |
| 3:2.3 | שהחלת |  |
| 3:2.4 | דברי |  |
| 3:2.5 | פיך |  |
| 3:2.6 | סכלות |  |
| 3:2.7 | והוללות |  |
| 3:2.8 | רעה |  |
| 3:2.9 | שאמרת |  |
| 3:2.10 | איך |  |
| 3:2.11 | מלאני |  |
| 3:2.12 | לבי |  |
| 3:2.13 | לדרוש |  |
| 3:2.14 | ולחקור |  |
| 3:2.15 | וכו | ' |
| 3:2.16 | ולמה |  |
| 3:2.17 | לא | ? |
| 3:2.18 | אדרבא |  |
| 3:2.19 | איפכא |  |
| 3:2.20 | מסתברא |  |
| 3:2.21 | כיון |  |
| 3:2.22 | שאומרים |  |
| 3:2.23 | לעבוד |  |
| 3:2.24 | לז"א |  |
| 3:2.25 | שהוא |  |
| 3:2.26 | נברא |  |
| 3:2.27 | יושר |  |
| 3:2.28 | לבב |  |
| 3:2.29 | דף |  |
| 3:2.30 | ג | ' |
| 3:2.31 | ע"ב |  |
| 3:2.32 | ועוז |  |
| 3:2.33 | לאלוהים | ). |
| 3:2.34 | חוב |  |
| 3:2.35 | גמור |  |
| 3:2.36 | עלינו |  |
| 3:2.37 | ועל |  |
| 3:2.38 | כל |  |
| 3:2.39 | אחד |  |
| 3:2.40 | ואחד |  |
| 3:2.41 | בישראל |  |
| 3:2.42 | לדעת |  |
| 3:2.43 | את |  |
| 3:2.44 | ה | ' |
| 3:2.45 | שהוא |  |
| 3:2.46 | בורא |  |
| 3:2.47 | ולא |  |
| 3:2.48 | נברא |  |
| 3:2.49 | ולעבדו |  |
| 3:2.50 | ככתוב |  |
| 3:2.51 | וידעת |  |
| 3:2.52 | היום |  |
| 3:2.53 | והשבות |  |
| 3:2.54 | אל |  |
| 3:2.55 | לבבך | . |
| 3:2.56 | וידעת |  |
| 3:2.57 | כי |  |
| 3:2.58 | ה | ' |
| 3:2.59 | אלהיך |  |
| 3:2.60 | הוא |  |
| 3:2.61 | האלהים |  |
| 3:2.62 | וכו | ' |
| 3:2.63 | ואומר |  |
| 3:2.64 | דע |  |
| 3:2.65 | את |  |
| 3:2.66 | אלהי |  |
| 3:2.67 | אביך |  |
| 3:2.68 | ועבדהו | . |
| 3:2.69 | והתירו |  |
| 3:2.70 | חכמים |  |
| 3:2.71 | לאפרושי |  |
| 3:2.72 | מאיסורא |  |
| 3:2.73 | אפי | ' |
| 3:2.74 | בבית |  |
| 3:2.75 | הכסא |  |
| 3:2.76 | ובבית |  |
| 3:2.77 | המרחץ |  |
| 3:2.78 | שאסור |  |
| 3:2.79 | להרהר |  |
| 3:2.80 | בהם |  |
| 3:2.81 | בדברי |  |
| 3:2.82 | תורה | . |
| 3:2.83 | ולאפרושי |  |
| 3:2.84 | מאיסורא |  |
| 3:2.85 | שרי | . |
| 3:2.86 | ועל |  |
| 3:2.87 | תמיהתך |  |
| 3:2.88 | איך |  |
| 3:2.89 | יטעו | ? |
| 3:2.90 | אין |  |
| 3:2.91 | זה |  |
| 3:2.92 | פלא | . |
| 3:2.93 | ולא |  |
| 3:2.94 | עדיפי |  |
| 3:2.95 | כת |  |
| 3:2.96 | המקובלים |  |
| 3:2.97 | האחרונים |  |
| 3:2.98 | אשר |  |
| 3:2.99 | הלכו |  |
| 3:2.100 | לתומם |  |
| 3:2.101 | אחרי |  |
| 3:2.102 | צו |  |
| 3:2.103 | של |  |
| 3:2.104 | הפילוסוף |  |
| 3:2.105 | מחבר |  |
| 3:2.106 | הזהר |  |
| 3:2.107 | שתלאו |  |
| 3:2.108 | ברשב"י |  |
| 3:2.109 | ור | ' |
| 3:2.110 | אלעזר |  |
| 3:2.111 | בנו |  |
| 3:2.112 | וחביריו |  |
| 3:2.113 | מסנהדרי |  |
| 3:2.114 | גדולה |  |
| 3:2.115 | היושבים |  |
| 3:2.116 | בלשכת |  |
| 3:2.117 | הגזית |  |
| 3:2.118 | שעליהם |  |
| 3:2.119 | אמר |  |
| 3:2.120 | הכתוב |  |
| 3:2.121 | כי |  |
| 3:2.122 | יפלא |  |
| 3:2.123 | ממך |  |
| 3:2.124 | דבר |  |
| 3:2.125 | למשפט |  |
| 3:2.126 | וכו | ' |
| 3:2.127 | וקמת |  |
| 3:2.128 | ועלית |  |
| 3:2.129 | אל |  |
| 3:2.130 | המקום |  |
| 3:2.131 | אשר |  |
| 3:2.132 | יבחר |  |
| 3:2.133 | ה | ' |
| 3:2.134 | אלוהיך |  |
| 3:2.135 | וכו | ' |
| 3:2.136 | ועשית |  |
| 3:2.137 | על |  |
| 3:2.138 | פי |  |
| 3:2.139 | הדבר |  |
| 3:2.140 | אשר |  |
| 3:2.141 | יגידו |  |
| 3:2.142 | לך |  |
| 3:2.143 | וכו | '. |
| 3:2.144 | ועם |  |
| 3:2.145 | כל |  |
| 3:2.146 | זה |  |
| 3:2.147 | אמרה |  |
| 3:2.148 | תורה |  |
| 3:2.149 | שאם |  |
| 3:2.150 | שגגו |  |
| 3:2.151 | ב"ד |  |
| 3:2.152 | הגדול |  |
| 3:2.153 | בהוראה |  |
| 3:2.154 | ועשו |  |
| 3:2.155 | העם |  |
| 3:2.156 | על |  |
| 3:2.157 | פיהם |  |
| 3:2.158 | מביאים |  |
| 3:2.159 | קרבן |  |
| 3:2.160 | על |  |
| 3:2.161 | שגגם | . |
| 3:2.162 | אם |  |
| 3:2.163 | בע"ז |  |
| 3:2.164 | שגגו |  |
| 3:2.165 | והורו |  |
| 3:2.166 | מביאים |  |
| 3:2.167 | פר |  |
| 3:2.168 | לעולה |  |
| 3:2.169 | ושעיר |  |
| 3:2.170 | לחטאת |  |
| 3:2.171 | מכל |  |
| 3:2.172 | שבט |  |
| 3:2.173 | ושבט | . |
| 3:2.174 | והוא |  |
| 3:2.175 | האמור |  |
| 3:2.176 | בפרשת |  |
| 3:2.177 | שלח |  |
| 3:2.178 | והיה |  |
| 3:2.179 | אם |  |
| 3:2.180 | מעיני |  |
| 3:2.181 | העדה |  |
| 3:2.182 | נעשתה |  |
| 3:2.183 | לשגגה | , |
| 3:2.184 | ואם |  |
| 3:2.185 | בשאר |  |
| 3:2.186 | כריתות |  |
| 3:2.187 | שגגו |  |
| 3:2.188 | והורו |  |
| 3:2.189 | מביא |  |
| 3:2.190 | כל |  |
| 3:2.191 | שבט |  |
| 3:2.192 | ושבט |  |
| 3:2.193 | פר |  |
| 3:2.194 | לחטאת |  |
| 3:2.195 | והוא |  |
| 3:2.196 | האמור |  |
| 3:2.197 | בפרשת |  |
| 3:2.198 | ויקרא | . |
| 3:2.199 | ואם |  |
| 3:2.200 | כל |  |
| 3:2.201 | עדת |  |
| 3:2.202 | ישראל |  |
| 3:2.203 | ישגו |  |
| 3:2.204 | ונעלם |  |
| 3:2.205 | דבר |  |
| 3:2.206 | מעיני |  |
| 3:2.207 | הקהל | . |


### Translation under review (3:2)

English: §3. And indeed, that you have begun the words of your mouth, folly and madness, evil that you said, "How my heart urged me to inquire and to examine, etc." And why not? On the contrary, the reverse is more reasonable (since they say to worship Ze'ir Anpin, who is created, upright of heart, folio 3b, and strength to God). It is a complete obligation upon us and upon each and every one of Israel to know the LORD, that He is Creator and not created, and to worship Him, as it is written: "And you shall know this day and take it to your heart." "And you shall know that the LORD your God is God, etc." And it says: "Know the God of your father and worship Him." The Sages permitted abstaining from the forbidden, even in the privy and in the bathhouse, where it is forbidden to meditate on them in words of Torah. but to keep away from the forbidden is permitted. As to your astonishment, how could they err? There is no wonder in this. Nor is the sect of the later Kabbalists any better, those who went in their simplicity after the command of the philosopher, the author of the Zohar who attributed it to Rabbi Shimon bar Yochai, and Rabbi Eleazar his son, and their companions of the Great Sanhedrin who sat in the chamber of the Hewn Stone about whom Scripture said: "If a matter is too hard for you in judgment, etc." "And you shall rise and go up to the place that the LORD your God will choose, etc." "And you shall act according to the word that they tell you, etc." And yet, for all this, the Torah said that if the Great Court erred in a ruling and the people acted on their word, they bring an offering for their error. If they erred in idolatry and ruled, they bring a bull as a burnt offering and a goat as a sin offering, from each tribe. This is what is stated in the parashah of Shelach: "And it shall be, if from the eyes of the congregation it was done in error." And if in the other offences punishable by karet they erred and ruled, each tribe brings a bull for a sin offering. This is what is stated in the parashah of Vayikra. And if "all the congregation of Israel errs and the matter is hidden from the eyes of the assembly."

Literal: section 3 and indeed that you have begun words of your mouth folly and madness evil that you said how urged me my heart to inquire and to examine etc. and why not on the contrary the reverse is more reasonable since they say to worship Ze'ir Anpin who is created upright of heart folio 3 b and strength to God obligation complete upon us and upon every one and one in Israel to know (object marker) the LORD that he is Creator and not created and to worship him as it is written and you shall know this day and take to heart to your heart and you shall know that the LORD your God he God etc. and it says know (object marker) the God of your father and worship him and they permitted the Sages to keep away from the forbidden even in the privy and in the bathhouse for it is forbidden to meditate on them in the words of Torah but to keep away from the forbidden it is permitted as to your astonishment how could they err there is no this wonder nor are better sect the Kabbalists the later ones who went in their simplicity after the command of the philosopher the author of the Zohar who attributed to Rabbi Shimon bar Yochai and Rabbi Eleazar his son and their companions of the Sanhedrin the Great who sat in the chamber of the Hewn Stone about whom said Scripture if is too hard for you a matter for judgment etc. and you shall rise and go up to the place which will choose the LORD your God etc. and you shall do according to the mouth of the word which they tell you etc. and yet all this said the Torah that if they erred Beit Din (the court) the Great in a ruling and acted the people on their word they bring an offering for their error if in idolatry (avodah zarah) they erred and ruled they bring a bull for a burnt offering and a goat for a sin offering from each tribe (repeated) this is what is stated in the parashah of Shelach and it shall be if from the eyes of the congregation it was done in error and if in the other offences punishable by karet they erred and ruled he brings every tribe and each bull for a sin offering this is what is stated in the parashah of Vayikra and if all the congregation of Israel errs and is hidden matter from the eyes of the assembly

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:2.1 | section 3 |  |  |  |
| 3:2.2 | and indeed |  |  |  |
| 3:2.3 | that you have begun |  |  |  |
| 3:2.4 | words |  |  |  |
| 3:2.5 | of your mouth |  |  |  |
| 3:2.6 | folly |  |  |  |
| 3:2.7 | and madness |  |  |  |
| 3:2.8 | evil |  |  |  |
| 3:2.9 | that you said |  |  |  |
| 3:2.10 | how |  |  |  |
| 3:2.11 | urged me |  |  |  |
| 3:2.12 | my heart |  |  |  |
| 3:2.13 | to inquire |  |  |  |
| 3:2.14 | and to examine |  |  |  |
| 3:2.15 | etc. |  |  |  |
| 3:2.16 | and why |  |  |  |
| 3:2.17 | not |  |  |  |
| 3:2.18 | on the contrary |  |  |  |
| 3:2.19 | the reverse |  |  |  |
| 3:2.20 | is more reasonable |  |  |  |
| 3:2.21 | since |  |  |  |
| 3:2.22 | they say |  |  |  |
| 3:2.23 | to worship |  |  |  |
| 3:2.24 | Ze'ir Anpin |  |  | לז"א = לזעיר אנפין, 'to Ze'ir Anpin' (the Small Face, a Kabbalistic partzuf); first occurrence in this unit |
| 3:2.25 | who is |  |  |  |
| 3:2.26 | created |  |  |  |
| 3:2.27 | upright |  |  |  |
| 3:2.28 | of heart |  |  |  |
| 3:2.29 | folio |  |  |  |
| 3:2.30 | 3 |  |  |  |
| 3:2.31 | b |  |  | ע"ב = side b of the folio |
| 3:2.32 | and strength |  |  |  |
| 3:2.33 | to God |  |  | the author's spelling אלוהים for אלהים; 'to God' |
| 3:2.34 | obligation |  |  |  |
| 3:2.35 | complete |  |  |  |
| 3:2.36 | upon us |  |  |  |
| 3:2.37 | and upon |  |  |  |
| 3:2.38 | every |  |  |  |
| 3:2.39 | one |  |  |  |
| 3:2.40 | and one |  |  |  |
| 3:2.41 | in Israel |  |  |  |
| 3:2.42 | to know |  |  |  |
| 3:2.43 | (object marker) |  |  |  |
| 3:2.44 | the LORD |  |  |  |
| 3:2.45 | that he is |  |  |  |
| 3:2.46 | Creator |  |  |  |
| 3:2.47 | and not |  |  |  |
| 3:2.48 | created |  |  |  |
| 3:2.49 | and to worship him |  |  |  |
| 3:2.50 | as it is written |  |  |  |
| 3:2.51 | and you shall know |  |  |  |
| 3:2.52 | this day |  |  |  |
| 3:2.53 | and take to heart |  |  |  |
| 3:2.54 | to |  |  |  |
| 3:2.55 | your heart |  |  |  |
| 3:2.56 | and you shall know |  |  |  |
| 3:2.57 | that |  |  |  |
| 3:2.58 | the LORD |  |  |  |
| 3:2.59 | your God |  |  |  |
| 3:2.60 | he |  |  |  |
| 3:2.61 | God |  |  |  |
| 3:2.62 | etc. |  |  |  |
| 3:2.63 | and it says |  |  |  |
| 3:2.64 | know |  |  |  |
| 3:2.65 | (object marker) |  |  |  |
| 3:2.66 | the God of |  |  |  |
| 3:2.67 | your father |  |  |  |
| 3:2.68 | and worship him |  |  |  |
| 3:2.69 | and they permitted |  |  |  |
| 3:2.70 | the Sages |  |  |  |
| 3:2.71 | to keep away |  |  |  |
| 3:2.72 | from the forbidden |  |  |  |
| 3:2.73 | even |  |  |  |
| 3:2.74 | in the |  |  |  |
| 3:2.75 | privy |  |  |  |
| 3:2.76 | and in the |  |  |  |
| 3:2.77 | bathhouse |  |  |  |
| 3:2.78 | for it is forbidden |  |  |  |
| 3:2.79 | to meditate |  |  |  |
| 3:2.80 | on them |  |  |  |
| 3:2.81 | in the words of |  |  |  |
| 3:2.82 | Torah |  |  |  |
| 3:2.83 | but to keep away |  |  |  |
| 3:2.84 | from the forbidden |  |  |  |
| 3:2.85 | it is permitted |  |  |  |
| 3:2.86 | as to |  |  |  |
| 3:2.87 | your astonishment |  |  |  |
| 3:2.88 | how |  |  |  |
| 3:2.89 | could they err |  |  |  |
| 3:2.90 | there is no |  |  |  |
| 3:2.91 | this |  |  |  |
| 3:2.92 | wonder |  |  |  |
| 3:2.93 | nor |  |  |  |
| 3:2.94 | are better |  |  |  |
| 3:2.95 | sect |  |  | כת, 'sect', a group |
| 3:2.96 | the Kabbalists |  |  |  |
| 3:2.97 | the later ones |  |  | האחרונים, 'the later ones' |
| 3:2.98 | who |  |  |  |
| 3:2.99 | went |  |  |  |
| 3:2.100 | in their simplicity |  |  |  |
| 3:2.101 | after |  |  |  |
| 3:2.102 | the command |  |  |  |
| 3:2.103 | of |  |  |  |
| 3:2.104 | the philosopher |  |  | הפילוסוף, the author's epithet for the author of the Zohar (decisions) |
| 3:2.105 | the author of |  |  |  |
| 3:2.106 | the Zohar |  |  |  |
| 3:2.107 | who attributed |  |  |  |
| 3:2.108 | to Rabbi Shimon bar Yochai |  |  | שתלאו, literally 'hung on'; the verb means to attribute |
| 3:2.109 | and Rabbi |  |  |  |
| 3:2.110 | Eleazar |  |  |  |
| 3:2.111 | his son |  |  |  |
| 3:2.112 | and their companions |  |  |  |
| 3:2.113 | of the Sanhedrin |  |  |  |
| 3:2.114 | the Great |  |  |  |
| 3:2.115 | who sat |  |  |  |
| 3:2.116 | in the chamber of |  |  |  |
| 3:2.117 | the Hewn Stone |  |  |  |
| 3:2.118 | about whom |  |  |  |
| 3:2.119 | said |  |  |  |
| 3:2.120 | Scripture |  |  |  |
| 3:2.121 | if |  |  |  |
| 3:2.122 | is too hard |  |  |  |
| 3:2.123 | for you |  |  |  |
| 3:2.124 | a matter |  |  |  |
| 3:2.125 | for judgment |  |  |  |
| 3:2.126 | etc. |  |  |  |
| 3:2.127 | and you shall rise |  |  |  |
| 3:2.128 | and go up |  |  |  |
| 3:2.129 | to |  |  |  |
| 3:2.130 | the place |  |  |  |
| 3:2.131 | which |  |  |  |
| 3:2.132 | will choose |  |  |  |
| 3:2.133 | the LORD |  |  |  |
| 3:2.134 | your God |  |  |  |
| 3:2.135 | etc. |  |  |  |
| 3:2.136 | and you shall do |  |  |  |
| 3:2.137 | according to |  |  |  |
| 3:2.138 | the mouth of |  |  |  |
| 3:2.139 | the word |  |  |  |
| 3:2.140 | which |  |  |  |
| 3:2.141 | they tell |  |  |  |
| 3:2.142 | you |  |  |  |
| 3:2.143 | etc. |  |  |  |
| 3:2.144 | and yet |  |  |  |
| 3:2.145 | all |  |  |  |
| 3:2.146 | this |  |  |  |
| 3:2.147 | said |  |  |  |
| 3:2.148 | the Torah |  |  |  |
| 3:2.149 | that if |  |  |  |
| 3:2.150 | they erred |  |  |  |
| 3:2.151 | Beit Din (the court) |  |  | ב"ד הגדול = בית דין הגדול, 'the Great Court' |
| 3:2.152 | the Great |  |  |  |
| 3:2.153 | in a ruling |  |  |  |
| 3:2.154 | and acted |  |  |  |
| 3:2.155 | the people |  |  |  |
| 3:2.156 | on |  |  |  |
| 3:2.157 | their word |  |  |  |
| 3:2.158 | they bring |  |  |  |
| 3:2.159 | an offering |  |  |  |
| 3:2.160 | for |  |  |  |
| 3:2.161 | their error |  |  |  |
| 3:2.162 | if |  |  |  |
| 3:2.163 | in idolatry (avodah zarah) |  |  | ע"ז = עבודה זרה, 'idolatry' (literally 'strange worship') |
| 3:2.164 | they erred |  |  |  |
| 3:2.165 | and ruled |  |  |  |
| 3:2.166 | they bring |  |  |  |
| 3:2.167 | a bull |  |  |  |
| 3:2.168 | for a burnt offering |  |  |  |
| 3:2.169 | and a goat |  |  |  |
| 3:2.170 | for a sin offering |  |  |  |
| 3:2.171 | from |  |  |  |
| 3:2.172 | each |  |  |  |
| 3:2.173 | tribe (repeated) |  |  |  |
| 3:2.174 | this is |  |  |  |
| 3:2.175 | what is stated |  |  |  |
| 3:2.176 | in the parashah of |  |  |  |
| 3:2.177 | Shelach |  |  | Shelach: the parashah of Numbers 13-15 |
| 3:2.178 | and it shall be |  |  |  |
| 3:2.179 | if |  |  |  |
| 3:2.180 | from the eyes of |  |  |  |
| 3:2.181 | the congregation |  |  |  |
| 3:2.182 | it was done |  |  |  |
| 3:2.183 | in error |  |  |  |
| 3:2.184 | and if |  |  |  |
| 3:2.185 | in the other |  |  |  |
| 3:2.186 | offences punishable by karet |  |  | כריתות: offences punished by karet |
| 3:2.187 | they erred |  |  |  |
| 3:2.188 | and ruled |  |  |  |
| 3:2.189 | he brings |  |  |  |
| 3:2.190 | every |  |  |  |
| 3:2.191 | tribe |  |  |  |
| 3:2.192 | and each |  |  |  |
| 3:2.193 | bull |  |  |  |
| 3:2.194 | for a sin offering |  |  |  |
| 3:2.195 | this is |  |  |  |
| 3:2.196 | what is stated |  |  |  |
| 3:2.197 | in the parashah of |  |  |  |
| 3:2.198 | Vayikra |  |  | ויקרא: the parashah of Vayikra (Leviticus) |
| 3:2.199 | and if |  |  |  |
| 3:2.200 | all |  |  |  |
| 3:2.201 | the congregation of |  |  |  |
| 3:2.202 | Israel |  |  |  |
| 3:2.203 | errs |  |  |  |
| 3:2.204 | and is hidden |  |  |  |
| 3:2.205 | matter |  |  |  |
| 3:2.206 | from the eyes of |  |  |  |
| 3:2.207 | the assembly |  |  |  |

Phrases:
- 3:2.1 3:2.2 | §3. And indeed, | section 3 and indeed
- 3:2.3 3:2.4 3:2.5 3:2.6 3:2.7 3:2.8 3:2.9 | that you have begun the words of your mouth, folly and madness, evil that you said, | that you have begun words of your mouth folly and madness evil that you said
- 3:2.10 3:2.11 3:2.12 3:2.13 3:2.14 | "How my heart urged me to inquire and to examine, | how urged me my heart to inquire and to examine
- 3:2.15 | etc." | etc.
- 3:2.16 3:2.17 | And why not? | and why not
- 3:2.18 3:2.19 3:2.20 | On the contrary, the reverse is more reasonable | on the contrary the reverse is more reasonable
- 3:2.21 3:2.22 3:2.23 3:2.24 3:2.25 3:2.26 3:2.27 3:2.28 3:2.29 3:2.30 3:2.31 3:2.32 3:2.33 | (since they say to worship Ze'ir Anpin, who is created, upright of heart, folio 3b, and strength to God). | since they say to worship Ze'ir Anpin who is created upright of heart folio 3 b and strength to God
- 3:2.34 3:2.35 3:2.36 | It is a complete obligation upon us | obligation complete upon us
- 3:2.37 3:2.38 3:2.39 3:2.40 3:2.41 | and upon each and every one of Israel | and upon every one and one in Israel
- 3:2.42 3:2.43 3:2.44 3:2.45 3:2.46 3:2.47 3:2.48 3:2.49 3:2.50 | to know the LORD, that He is Creator and not created, and to worship Him, as it is written: | to know (object marker) the LORD that he is Creator and not created and to worship him as it is written
- 3:2.51 3:2.52 3:2.53 3:2.54 3:2.55 | "And you shall know this day and take it to your heart." | and you shall know this day and take to heart to your heart
- 3:2.56 3:2.57 3:2.58 3:2.59 3:2.60 3:2.61 3:2.62 | "And you shall know that the LORD your God is God, etc." | and you shall know that the LORD your God he God etc.
- 3:2.63 3:2.64 3:2.65 3:2.66 3:2.67 3:2.68 | And it says: "Know the God of your father and worship Him." | and it says know (object marker) the God of your father and worship him
- 3:2.69 3:2.70 3:2.71 3:2.72 | The Sages permitted abstaining from the forbidden, | and they permitted the Sages to keep away from the forbidden
- 3:2.73 3:2.74 3:2.75 3:2.76 3:2.77 | even in the privy and in the bathhouse, | even in the privy and in the bathhouse
- 3:2.78 3:2.79 3:2.80 3:2.81 3:2.82 | where it is forbidden to meditate on them in words of Torah. | for it is forbidden to meditate on them in the words of Torah
- 3:2.83 3:2.84 3:2.85 | but to keep away from the forbidden is permitted. | but to keep away from the forbidden it is permitted
- 3:2.86 3:2.87 3:2.88 3:2.89 | As to your astonishment, how could they err? | as to your astonishment how could they err
- 3:2.90 3:2.91 3:2.92 | There is no wonder in this. | there is no this wonder
- 3:2.93 3:2.94 3:2.95 3:2.96 3:2.97 3:2.98 3:2.99 3:2.100 | Nor is the sect of the later Kabbalists any better, those who went in their simplicity | nor are better sect the Kabbalists the later ones who went in their simplicity
- 3:2.101 3:2.102 3:2.103 3:2.104 3:2.105 3:2.106 | after the command of the philosopher, the author of the Zohar | after the command of the philosopher the author of the Zohar
- 3:2.107 3:2.108 3:2.109 3:2.110 3:2.111 3:2.112 3:2.113 3:2.114 | who attributed it to Rabbi Shimon bar Yochai, and Rabbi Eleazar his son, and their companions of the Great Sanhedrin | who attributed to Rabbi Shimon bar Yochai and Rabbi Eleazar his son and their companions of the Sanhedrin the Great
- 3:2.115 3:2.116 3:2.117 | who sat in the chamber of the Hewn Stone | who sat in the chamber of the Hewn Stone
- 3:2.118 3:2.119 3:2.120 3:2.121 3:2.122 3:2.123 3:2.124 3:2.125 3:2.126 | about whom Scripture said: "If a matter is too hard for you in judgment, etc." | about whom said Scripture if is too hard for you a matter for judgment etc.
- 3:2.127 3:2.128 3:2.129 3:2.130 3:2.131 3:2.132 3:2.133 3:2.134 3:2.135 | "And you shall rise and go up to the place that the LORD your God will choose, etc." | and you shall rise and go up to the place which will choose the LORD your God etc.
- 3:2.136 3:2.137 3:2.138 3:2.139 3:2.140 3:2.141 3:2.142 3:2.143 | "And you shall act according to the word that they tell you, etc." | and you shall do according to the mouth of the word which they tell you etc.
- 3:2.144 3:2.145 3:2.146 3:2.147 3:2.148 | And yet, for all this, the Torah said | and yet all this said the Torah
- 3:2.149 3:2.150 3:2.151 3:2.152 3:2.153 3:2.154 3:2.155 3:2.156 3:2.157 3:2.158 3:2.159 3:2.160 3:2.161 | that if the Great Court erred in a ruling and the people acted on their word, they bring an offering for their error. | that if they erred Beit Din (the court) the Great in a ruling and acted the people on their word they bring an offering for their error
- 3:2.162 3:2.163 3:2.164 3:2.165 3:2.166 3:2.167 3:2.168 3:2.169 3:2.170 3:2.171 3:2.172 3:2.173 | If they erred in idolatry and ruled, they bring a bull as a burnt offering and a goat as a sin offering, from each tribe. | if in idolatry (avodah zarah) they erred and ruled they bring a bull for a burnt offering and a goat for a sin offering from each tribe (repeated)
- 3:2.174 3:2.175 3:2.176 3:2.177 3:2.178 3:2.179 3:2.180 3:2.181 3:2.182 3:2.183 | This is what is stated in the parashah of Shelach: "And it shall be, if from the eyes of the congregation it was done in error." | this is what is stated in the parashah of Shelach and it shall be if from the eyes of the congregation it was done in error
- 3:2.184 3:2.185 3:2.186 3:2.187 3:2.188 3:2.189 3:2.190 3:2.191 3:2.192 3:2.193 3:2.194 3:2.195 3:2.196 3:2.197 3:2.198 | And if in the other offences punishable by karet they erred and ruled, each tribe brings a bull for a sin offering. This is what is stated in the parashah of Vayikra. | and if in the other offences punishable by karet they erred and ruled he brings every tribe and each bull for a sin offering this is what is stated in the parashah of Vayikra
- 3:2.199 3:2.200 3:2.201 3:2.202 3:2.203 3:2.204 3:2.205 3:2.206 3:2.207 | And if "all the congregation of Israel errs and the matter is hidden from the eyes of the assembly." | and if all the congregation of Israel errs and is hidden matter from the eyes of the assembly

Grammar notes:
- שהחלת דברי פיך: a perfect with the second-person ending, 'that you have begun the words of your mouth'; the reader's words are cited as 'שאמרת'.
- איך מלאני לבי לדרוש ולחקור: the infinitives לדרוש and לחקור follow 'urged me'; the reader's question is closed by 'וכו'' and the author's reply begins with 'ולמה לא'.
- ועוז לאלוהים sits in the author's parenthesis and gives no formula, so it is not marked as quotation; the parenthesis is kept as written.
- ואומר: 'and it says' introduces a verse; 'ועבדהו' is the verb 'worship' with the object suffix -הו, 'worship him'.
- The conditional clauses (אם שגגו ... מביאים; ואם שגגו ... מביא) are built on perfects of 'to err' with present-tense apodosis, as the verses have them.

Commentary rulings:
- apparatus: Deut 4:39: 'וידעת היום והשבות אל לבבך', quoted with 'ככתוב'.
- apparatus: Deut 7:9: 'וידעת כי ה' אלהיך הוא האלהים', quoted with 'וכו''.
- apparatus: 1 Chr 28:9: 'דע את אלהי אביך ועבדהו', quoted with 'ואומר'.
- apparatus: ועוז לאלוהים in the parenthesis: possibly Ps 68:35; not verified, and no formula marks it.
- apparatus: Deut 17:8: 'כי יפלא ממך דבר למשפט' and 'וקמת ועלית אל המקום אשר יבחר ה' אלהיך'. Deut 17:10: 'ועשית על פי הדבר אשר יגידו לך'. The author writes 'אלוהיך' for 'אלהיך'; the English follows the verse.
- apparatus: Num 15:24: 'והיה אם מעיני העדה נעשתה לשגגה'.
- apparatus: Lev 4:13: 'ואם כל עדת ישראל ישגו ונעלם דבר מעיני הקהל'.
- apparatus: The author's 'folio 3b' is given as he writes it; no Zohar text is quoted here.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

