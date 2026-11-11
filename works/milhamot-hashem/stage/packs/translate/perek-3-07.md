# Translation pack: milhamot-hashem perek-3-07 (Milchamot ha-Shem 3:12–3:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-07`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- והוא שלא: "provided that" (fixed).
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
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
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
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


## Milchamot ha-Shem 3:12

Hebrew (Wikisource): יב) הנה כי כן בודאי שכל דבריך הפלגות וגיזומים להבהיל את השוטים והמשוגעים אוהבי הנפלאות והנמנעות ולא בדעת תדבר. וכל דבריך לא בהשכל. ובמרמה אתה ממאן לדעת את ה' כמאמר הנביא ירמיה ע"ה בשם ה' שבתך בתוך מרמה במרמה מאנו דעת אותי נאם ה'. וכבר אמר אחד מן הרמאים שבכם ידענו שהאמת אתכם. אך מה לכם לגלות דבר זה ולפרסם לקצת התלמידים. טוב לכם להניחם בטעותם ממש כדברי הרשע שבארבעה בנים האומר מה העבודה הזאת לכם וכו'. והראיה שאתה טח תפל שהרי פרסמת במכתבך זה שהרב בעל מצרף האמונה הוא ח"ו מין וכופר. וקראת את ספרו משרף האמונה. והוא בתומו הולך לא סר מדברי הזהר ומקדש מלך, וכס"א, והרש"ב, ויושר לבב, וספר הברית ור' האר"י ומהרח"ו בס' עץ חיים ונחלת יוסף וסולתם ימין ושמאל. וכמו שיתבאר לקמן, ותחלה אביא לך כמה ראיות. שאין הקבלה הזאת הלכה למשה מסיני שהרי אמרו רז"ל בעירובין דכל מתניתא דלא מיתניא בי רבי אושעיא לאו מתניתא היא ולא תתיבון מינה, דמשבשתא היא. הרי דכל מתניתא שלא נכנסה לבית מדרשו של רבי חייא או של ר' אושעיא לא מיבעיא דלא הויא הלכה למשה מסיני. אלא דאפילו לאותובי מינה על אמורא לא מותבינן. ובירושלמי עירובין פרק קמא בפסקא דלחיים. וז"ל אמר ליה אחא כן אמר ר' שמעון בן לקיש כל משנה שלא נכנסה לחבורה אין סומכין עליה ע"כ, ועיין לקמן סי' מ"ב. ובקונדריס קדום מחבורי קדמונינו מצאתי כתוב משם ר' סעדיה גאון וז"ל כתאב שיעור קומה לם יקע עליה נקל אמנתא אר ליס הוא פי אל משנה ולא פי אל תלמוד, ולא דליל נסתדל בה הל קאלה רבי ישמעאל עלי' אלחקיקה או גירה קאלה פנסבה אליה כמא תנסב כתיר מן אלכתב אלמפתעלה אלי' בעץ אלעלמא אלמשהורין, עכ"ל. (תרגום מאת הרב המחבר: "ספר שיעור קומה אין בידינו קבלה עליו מחכמי אמונתנו, לפי שאינו לא במשנה ולא בתלמוד, וגם אין ראיה ללמוד ממנה אם אמרו ר' ישמעאל באמת או אחרים זולתו אמרו וייחסו לר' ישמעאל כמו שמייחסים הרבה מן הספרים העשויים לקצת מן החכמים המפורסמים").

| id | word | punctuation after |
|---|---|---|
| 3:12.1 | יב | ) |
| 3:12.2 | הנה |  |
| 3:12.3 | כי |  |
| 3:12.4 | כן |  |
| 3:12.5 | בודאי |  |
| 3:12.6 | שכל |  |
| 3:12.7 | דבריך |  |
| 3:12.8 | הפלגות |  |
| 3:12.9 | וגיזומים |  |
| 3:12.10 | להבהיל |  |
| 3:12.11 | את |  |
| 3:12.12 | השוטים |  |
| 3:12.13 | והמשוגעים |  |
| 3:12.14 | אוהבי |  |
| 3:12.15 | הנפלאות |  |
| 3:12.16 | והנמנעות |  |
| 3:12.17 | ולא |  |
| 3:12.18 | בדעת |  |
| 3:12.19 | תדבר | . |
| 3:12.20 | וכל |  |
| 3:12.21 | דבריך |  |
| 3:12.22 | לא |  |
| 3:12.23 | בהשכל | . |
| 3:12.24 | ובמרמה |  |
| 3:12.25 | אתה |  |
| 3:12.26 | ממאן |  |
| 3:12.27 | לדעת |  |
| 3:12.28 | את |  |
| 3:12.29 | ה | ' |
| 3:12.30 | כמאמר |  |
| 3:12.31 | הנביא |  |
| 3:12.32 | ירמיה |  |
| 3:12.33 | ע"ה |  |
| 3:12.34 | בשם |  |
| 3:12.35 | ה | ' |
| 3:12.36 | שבתך |  |
| 3:12.37 | בתוך |  |
| 3:12.38 | מרמה |  |
| 3:12.39 | במרמה |  |
| 3:12.40 | מאנו |  |
| 3:12.41 | דעת |  |
| 3:12.42 | אותי |  |
| 3:12.43 | נאם |  |
| 3:12.44 | ה | '. |
| 3:12.45 | וכבר |  |
| 3:12.46 | אמר |  |
| 3:12.47 | אחד |  |
| 3:12.48 | מן |  |
| 3:12.49 | הרמאים |  |
| 3:12.50 | שבכם |  |
| 3:12.51 | ידענו |  |
| 3:12.52 | שהאמת |  |
| 3:12.53 | אתכם | . |
| 3:12.54 | אך |  |
| 3:12.55 | מה |  |
| 3:12.56 | לכם |  |
| 3:12.57 | לגלות |  |
| 3:12.58 | דבר |  |
| 3:12.59 | זה |  |
| 3:12.60 | ולפרסם |  |
| 3:12.61 | לקצת |  |
| 3:12.62 | התלמידים | . |
| 3:12.63 | טוב |  |
| 3:12.64 | לכם |  |
| 3:12.65 | להניחם |  |
| 3:12.66 | בטעותם |  |
| 3:12.67 | ממש |  |
| 3:12.68 | כדברי |  |
| 3:12.69 | הרשע |  |
| 3:12.70 | שבארבעה |  |
| 3:12.71 | בנים |  |
| 3:12.72 | האומר |  |
| 3:12.73 | מה |  |
| 3:12.74 | העבודה |  |
| 3:12.75 | הזאת |  |
| 3:12.76 | לכם |  |
| 3:12.77 | וכו | '. |
| 3:12.78 | והראיה |  |
| 3:12.79 | שאתה |  |
| 3:12.80 | טח |  |
| 3:12.81 | תפל |  |
| 3:12.82 | שהרי |  |
| 3:12.83 | פרסמת |  |
| 3:12.84 | במכתבך |  |
| 3:12.85 | זה |  |
| 3:12.86 | שהרב |  |
| 3:12.87 | בעל |  |
| 3:12.88 | מצרף |  |
| 3:12.89 | האמונה |  |
| 3:12.90 | הוא |  |
| 3:12.91 | ח"ו |  |
| 3:12.92 | מין |  |
| 3:12.93 | וכופר | . |
| 3:12.94 | וקראת |  |
| 3:12.95 | את |  |
| 3:12.96 | ספרו |  |
| 3:12.97 | משרף |  |
| 3:12.98 | האמונה | . |
| 3:12.99 | והוא |  |
| 3:12.100 | בתומו |  |
| 3:12.101 | הולך |  |
| 3:12.102 | לא |  |
| 3:12.103 | סר |  |
| 3:12.104 | מדברי |  |
| 3:12.105 | הזהר |  |
| 3:12.106 | ומקדש |  |
| 3:12.107 | מלך | , |
| 3:12.108 | וכס"א | , |
| 3:12.109 | והרש"ב | , |
| 3:12.110 | ויושר |  |
| 3:12.111 | לבב | , |
| 3:12.112 | וספר |  |
| 3:12.113 | הברית |  |
| 3:12.114 | ור | ' |
| 3:12.115 | האר"י |  |
| 3:12.116 | ומהרח"ו |  |
| 3:12.117 | בס | ' |
| 3:12.118 | עץ |  |
| 3:12.119 | חיים |  |
| 3:12.120 | ונחלת |  |
| 3:12.121 | יוסף |  |
| 3:12.122 | וסולתם |  |
| 3:12.123 | ימין |  |
| 3:12.124 | ושמאל | . |
| 3:12.125 | וכמו |  |
| 3:12.126 | שיתבאר |  |
| 3:12.127 | לקמן | , |
| 3:12.128 | ותחלה |  |
| 3:12.129 | אביא |  |
| 3:12.130 | לך |  |
| 3:12.131 | כמה |  |
| 3:12.132 | ראיות | . |
| 3:12.133 | שאין |  |
| 3:12.134 | הקבלה |  |
| 3:12.135 | הזאת |  |
| 3:12.136 | הלכה |  |
| 3:12.137 | למשה |  |
| 3:12.138 | מסיני |  |
| 3:12.139 | שהרי |  |
| 3:12.140 | אמרו |  |
| 3:12.141 | רז"ל |  |
| 3:12.142 | בעירובין |  |
| 3:12.143 | דכל |  |
| 3:12.144 | מתניתא |  |
| 3:12.145 | דלא |  |
| 3:12.146 | מיתניא |  |
| 3:12.147 | בי |  |
| 3:12.148 | רבי |  |
| 3:12.149 | אושעיא |  |
| 3:12.150 | לאו |  |
| 3:12.151 | מתניתא |  |
| 3:12.152 | היא |  |
| 3:12.153 | ולא |  |
| 3:12.154 | תתיבון |  |
| 3:12.155 | מינה | , |
| 3:12.156 | דמשבשתא |  |
| 3:12.157 | היא | . |
| 3:12.158 | הרי |  |
| 3:12.159 | דכל |  |
| 3:12.160 | מתניתא |  |
| 3:12.161 | שלא |  |
| 3:12.162 | נכנסה |  |
| 3:12.163 | לבית |  |
| 3:12.164 | מדרשו |  |
| 3:12.165 | של |  |
| 3:12.166 | רבי |  |
| 3:12.167 | חייא |  |
| 3:12.168 | או |  |
| 3:12.169 | של |  |
| 3:12.170 | ר | ' |
| 3:12.171 | אושעיא |  |
| 3:12.172 | לא |  |
| 3:12.173 | מיבעיא |  |
| 3:12.174 | דלא |  |
| 3:12.175 | הויא |  |
| 3:12.176 | הלכה |  |
| 3:12.177 | למשה |  |
| 3:12.178 | מסיני | . |
| 3:12.179 | אלא |  |
| 3:12.180 | דאפילו |  |
| 3:12.181 | לאותובי |  |
| 3:12.182 | מינה |  |
| 3:12.183 | על |  |
| 3:12.184 | אמורא |  |
| 3:12.185 | לא |  |
| 3:12.186 | מותבינן | . |
| 3:12.187 | ובירושלמי |  |
| 3:12.188 | עירובין |  |
| 3:12.189 | פרק |  |
| 3:12.190 | קמא |  |
| 3:12.191 | בפסקא |  |
| 3:12.192 | דלחיים | . |
| 3:12.193 | וז"ל |  |
| 3:12.194 | אמר |  |
| 3:12.195 | ליה |  |
| 3:12.196 | אחא |  |
| 3:12.197 | כן |  |
| 3:12.198 | אמר |  |
| 3:12.199 | ר | ' |
| 3:12.200 | שמעון |  |
| 3:12.201 | בן |  |
| 3:12.202 | לקיש |  |
| 3:12.203 | כל |  |
| 3:12.204 | משנה |  |
| 3:12.205 | שלא |  |
| 3:12.206 | נכנסה |  |
| 3:12.207 | לחבורה |  |
| 3:12.208 | אין |  |
| 3:12.209 | סומכין |  |
| 3:12.210 | עליה |  |
| 3:12.211 | ע"כ | , |
| 3:12.212 | ועיין |  |
| 3:12.213 | לקמן |  |
| 3:12.214 | סי | ' |
| 3:12.215 | מ"ב | . |
| 3:12.216 | ובקונדריס |  |
| 3:12.217 | קדום |  |
| 3:12.218 | מחבורי |  |
| 3:12.219 | קדמונינו |  |
| 3:12.220 | מצאתי |  |
| 3:12.221 | כתוב |  |
| 3:12.222 | משם |  |
| 3:12.223 | ר | ' |
| 3:12.224 | סעדיה |  |
| 3:12.225 | גאון |  |
| 3:12.226 | וז"ל |  |
| 3:12.227 | כתאב |  |
| 3:12.228 | שיעור |  |
| 3:12.229 | קומה |  |
| 3:12.230 | לם |  |
| 3:12.231 | יקע |  |
| 3:12.232 | עליה |  |
| 3:12.233 | נקל |  |
| 3:12.234 | אמנתא |  |
| 3:12.235 | אר |  |
| 3:12.236 | ליס |  |
| 3:12.237 | הוא |  |
| 3:12.238 | פי |  |
| 3:12.239 | אל |  |
| 3:12.240 | משנה |  |
| 3:12.241 | ולא |  |
| 3:12.242 | פי |  |
| 3:12.243 | אל |  |
| 3:12.244 | תלמוד | , |
| 3:12.245 | ולא |  |
| 3:12.246 | דליל |  |
| 3:12.247 | נסתדל |  |
| 3:12.248 | בה |  |
| 3:12.249 | הל |  |
| 3:12.250 | קאלה |  |
| 3:12.251 | רבי |  |
| 3:12.252 | ישמעאל |  |
| 3:12.253 | עלי | ' |
| 3:12.254 | אלחקיקה |  |
| 3:12.255 | או |  |
| 3:12.256 | גירה |  |
| 3:12.257 | קאלה |  |
| 3:12.258 | פנסבה |  |
| 3:12.259 | אליה |  |
| 3:12.260 | כמא |  |
| 3:12.261 | תנסב |  |
| 3:12.262 | כתיר |  |
| 3:12.263 | מן |  |
| 3:12.264 | אלכתב |  |
| 3:12.265 | אלמפתעלה |  |
| 3:12.266 | אלי | ' |
| 3:12.267 | בעץ |  |
| 3:12.268 | אלעלמא |  |
| 3:12.269 | אלמשהורין | , |
| 3:12.270 | עכ"ל | . |
| 3:12.271 | תרגום |  |
| 3:12.272 | מאת |  |
| 3:12.273 | הרב |  |
| 3:12.274 | המחבר | : |
| 3:12.275 | ספר |  |
| 3:12.276 | שיעור |  |
| 3:12.277 | קומה |  |
| 3:12.278 | אין |  |
| 3:12.279 | בידינו |  |
| 3:12.280 | קבלה |  |
| 3:12.281 | עליו |  |
| 3:12.282 | מחכמי |  |
| 3:12.283 | אמונתנו | , |
| 3:12.284 | לפי |  |
| 3:12.285 | שאינו |  |
| 3:12.286 | לא |  |
| 3:12.287 | במשנה |  |
| 3:12.288 | ולא |  |
| 3:12.289 | בתלמוד | , |
| 3:12.290 | וגם |  |
| 3:12.291 | אין |  |
| 3:12.292 | ראיה |  |
| 3:12.293 | ללמוד |  |
| 3:12.294 | ממנה |  |
| 3:12.295 | אם |  |
| 3:12.296 | אמרו |  |
| 3:12.297 | ר | ' |
| 3:12.298 | ישמעאל |  |
| 3:12.299 | באמת |  |
| 3:12.300 | או |  |
| 3:12.301 | אחרים |  |
| 3:12.302 | זולתו |  |
| 3:12.303 | אמרו |  |
| 3:12.304 | וייחסו |  |
| 3:12.305 | לר | ' |
| 3:12.306 | ישמעאל |  |
| 3:12.307 | כמו |  |
| 3:12.308 | שמייחסים |  |
| 3:12.309 | הרבה |  |
| 3:12.310 | מן |  |
| 3:12.311 | הספרים |  |
| 3:12.312 | העשויים |  |
| 3:12.313 | לקצת |  |
| 3:12.314 | מן |  |
| 3:12.315 | החכמים |  |
| 3:12.316 | המפורסמים | "). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

