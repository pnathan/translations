# Translation pack: milhamot-hashem perek-6-07 (Milchamot ha-Shem 6:12–6:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-07`.

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

- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
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


## Milchamot ha-Shem 6:12

Hebrew (Wikisource): ועל ידי הרבנים הפתאים שקבלוהו לתומם באמונת אומן שהוא חבור רשב"י והחזיקו בו על שמו ולא שמו אל לבם לחקור ולדרוש מבטן מי יצא, ומי התיר לדרוש ולחקור במופלא ומכוסא, ושהוא סותר אמונתינו הטהורה שקבלנו מאבותינו ומרז"ל באל אחד מיוחד מכל שאר האחדים, וזה לסבת שעירב בו הרבה דברים מדברי רז"ל, ומזהיר מאד בקיום המצות של תורה והעסק בתורה שבכתב ותורתו החדשה שהמציאה חלף המשנה והתלמוד שהם מעורבים מטוב ורע, וקורא לתורתו תורה שבע"פ שכולה טוב ורק שיהיו לשם אלהים אחרים, קב"ה ושכינתיה עלאין (אבא ואימא) וקב"ה ושכינתיה התאין (זעיר ונוקביה) ולא שנה במצות התורה רק בקצת פרטים שכיוון לרמוז אמונתו באלהים רבים, ולשום עיקר העבודה לבן שהוא זעיר אנפין, וברוב המצות מכוינים בהם לקשט הנשים שכינתא עלאה ושכינתא תתאה, כדי למצוא חן בעיני בעליהן לעוררם לתשמיש (כמו שכ' הר' משנת חסידים בתיקון חצות וז"ל וכיצד יעשה ישב על הארץ אצל המזוזה שבפתח שרומזת להח' ורגליו יחפים ממנעל, ויתן אפר מקלה וכו' ועל ידי זה מתחמם הזעיר אנפין ומתאוה לה וכו') זה להוליד נשמות למלאכים. וזה להוליד נשמות חדשות לישראל (זהר יתרו דף פ"ט ע"כ) שלא כרצון חכמים בבראשית רבה (הבאתיו לעיל סי' ל"א כל שיש לו תולדות וכו') ויחס עיקר העבודה וקיום המצות והברכות לבן כנז"ל. ועוד יתבאר לקמן. ועל ידי רעיונותיהם העולים על רוחם חשבו לפרש הכתובים שנאמרו על דרך המליצה והרחבת הלשון או בדרך חיבה כמו שאמרו רז"ל לא זז מחבבה וכו' כמשמען והנחתן הראשונה יד ה', עיני ה', יעשן אף ה' עלה עשן באפו, בצלמנו כדמותנו, בנים אתם לה' אלהיכם, אחי ורעי, ויראו את אלוהי ישראל, עורה למה תישן ה', הנה לא ינום ולא יישן שומר ישראל, וכיוצא בהן, יפרשו אותן על ידי רעיונות אלו כהנחתן הראשונה, היפך מאמר רז"ל המתרגם פסוק כצורתו הרי זה בדאי. וכמו שכ' הרס"ג בהאמונות והדעות מאמר האחדות. והרמב"ם במורה ח"א וזולתם.

| id | word | punctuation after |
|---|---|---|
| 6:12.1 | ועל |  |
| 6:12.2 | ידי |  |
| 6:12.3 | הרבנים |  |
| 6:12.4 | הפתאים |  |
| 6:12.5 | שקבלוהו |  |
| 6:12.6 | לתומם |  |
| 6:12.7 | באמונת |  |
| 6:12.8 | אומן |  |
| 6:12.9 | שהוא |  |
| 6:12.10 | חבור |  |
| 6:12.11 | רשב"י |  |
| 6:12.12 | והחזיקו |  |
| 6:12.13 | בו |  |
| 6:12.14 | על |  |
| 6:12.15 | שמו |  |
| 6:12.16 | ולא |  |
| 6:12.17 | שמו |  |
| 6:12.18 | אל |  |
| 6:12.19 | לבם |  |
| 6:12.20 | לחקור |  |
| 6:12.21 | ולדרוש |  |
| 6:12.22 | מבטן |  |
| 6:12.23 | מי |  |
| 6:12.24 | יצא | , |
| 6:12.25 | ומי |  |
| 6:12.26 | התיר |  |
| 6:12.27 | לדרוש |  |
| 6:12.28 | ולחקור |  |
| 6:12.29 | במופלא |  |
| 6:12.30 | ומכוסא | , |
| 6:12.31 | ושהוא |  |
| 6:12.32 | סותר |  |
| 6:12.33 | אמונתינו |  |
| 6:12.34 | הטהורה |  |
| 6:12.35 | שקבלנו |  |
| 6:12.36 | מאבותינו |  |
| 6:12.37 | ומרז"ל |  |
| 6:12.38 | באל |  |
| 6:12.39 | אחד |  |
| 6:12.40 | מיוחד |  |
| 6:12.41 | מכל |  |
| 6:12.42 | שאר |  |
| 6:12.43 | האחדים | , |
| 6:12.44 | וזה |  |
| 6:12.45 | לסבת |  |
| 6:12.46 | שעירב |  |
| 6:12.47 | בו |  |
| 6:12.48 | הרבה |  |
| 6:12.49 | דברים |  |
| 6:12.50 | מדברי |  |
| 6:12.51 | רז"ל | , |
| 6:12.52 | ומזהיר |  |
| 6:12.53 | מאד |  |
| 6:12.54 | בקיום |  |
| 6:12.55 | המצות |  |
| 6:12.56 | של |  |
| 6:12.57 | תורה |  |
| 6:12.58 | והעסק |  |
| 6:12.59 | בתורה |  |
| 6:12.60 | שבכתב |  |
| 6:12.61 | ותורתו |  |
| 6:12.62 | החדשה |  |
| 6:12.63 | שהמציאה |  |
| 6:12.64 | חלף |  |
| 6:12.65 | המשנה |  |
| 6:12.66 | והתלמוד |  |
| 6:12.67 | שהם |  |
| 6:12.68 | מעורבים |  |
| 6:12.69 | מטוב |  |
| 6:12.70 | ורע | , |
| 6:12.71 | וקורא |  |
| 6:12.72 | לתורתו |  |
| 6:12.73 | תורה |  |
| 6:12.74 | שבע"פ |  |
| 6:12.75 | שכולה |  |
| 6:12.76 | טוב |  |
| 6:12.77 | ורק |  |
| 6:12.78 | שיהיו |  |
| 6:12.79 | לשם |  |
| 6:12.80 | אלהים |  |
| 6:12.81 | אחרים | , |
| 6:12.82 | קב"ה |  |
| 6:12.83 | ושכינתיה |  |
| 6:12.84 | עלאין |  |
| 6:12.85 | אבא |  |
| 6:12.86 | ואימא | ) |
| 6:12.87 | וקב"ה |  |
| 6:12.88 | ושכינתיה |  |
| 6:12.89 | התאין |  |
| 6:12.90 | זעיר |  |
| 6:12.91 | ונוקביה | ) |
| 6:12.92 | ולא |  |
| 6:12.93 | שנה |  |
| 6:12.94 | במצות |  |
| 6:12.95 | התורה |  |
| 6:12.96 | רק |  |
| 6:12.97 | בקצת |  |
| 6:12.98 | פרטים |  |
| 6:12.99 | שכיוון |  |
| 6:12.100 | לרמוז |  |
| 6:12.101 | אמונתו |  |
| 6:12.102 | באלהים |  |
| 6:12.103 | רבים | , |
| 6:12.104 | ולשום |  |
| 6:12.105 | עיקר |  |
| 6:12.106 | העבודה |  |
| 6:12.107 | לבן |  |
| 6:12.108 | שהוא |  |
| 6:12.109 | זעיר |  |
| 6:12.110 | אנפין | , |
| 6:12.111 | וברוב |  |
| 6:12.112 | המצות |  |
| 6:12.113 | מכוינים |  |
| 6:12.114 | בהם |  |
| 6:12.115 | לקשט |  |
| 6:12.116 | הנשים |  |
| 6:12.117 | שכינתא |  |
| 6:12.118 | עלאה |  |
| 6:12.119 | ושכינתא |  |
| 6:12.120 | תתאה | , |
| 6:12.121 | כדי |  |
| 6:12.122 | למצוא |  |
| 6:12.123 | חן |  |
| 6:12.124 | בעיני |  |
| 6:12.125 | בעליהן |  |
| 6:12.126 | לעוררם |  |
| 6:12.127 | לתשמיש |  |
| 6:12.128 | כמו |  |
| 6:12.129 | שכ | ' |
| 6:12.130 | הר | ' |
| 6:12.131 | משנת |  |
| 6:12.132 | חסידים |  |
| 6:12.133 | בתיקון |  |
| 6:12.134 | חצות |  |
| 6:12.135 | וז"ל |  |
| 6:12.136 | וכיצד |  |
| 6:12.137 | יעשה |  |
| 6:12.138 | ישב |  |
| 6:12.139 | על |  |
| 6:12.140 | הארץ |  |
| 6:12.141 | אצל |  |
| 6:12.142 | המזוזה |  |
| 6:12.143 | שבפתח |  |
| 6:12.144 | שרומזת |  |
| 6:12.145 | להח | ' |
| 6:12.146 | ורגליו |  |
| 6:12.147 | יחפים |  |
| 6:12.148 | ממנעל | , |
| 6:12.149 | ויתן |  |
| 6:12.150 | אפר |  |
| 6:12.151 | מקלה |  |
| 6:12.152 | וכו | ' |
| 6:12.153 | ועל |  |
| 6:12.154 | ידי |  |
| 6:12.155 | זה |  |
| 6:12.156 | מתחמם |  |
| 6:12.157 | הזעיר |  |
| 6:12.158 | אנפין |  |
| 6:12.159 | ומתאוה |  |
| 6:12.160 | לה |  |
| 6:12.161 | וכו | ') |
| 6:12.162 | זה |  |
| 6:12.163 | להוליד |  |
| 6:12.164 | נשמות |  |
| 6:12.165 | למלאכים | . |
| 6:12.166 | וזה |  |
| 6:12.167 | להוליד |  |
| 6:12.168 | נשמות |  |
| 6:12.169 | חדשות |  |
| 6:12.170 | לישראל |  |
| 6:12.171 | זהר |  |
| 6:12.172 | יתרו |  |
| 6:12.173 | דף |  |
| 6:12.174 | פ"ט |  |
| 6:12.175 | ע"כ | ) |
| 6:12.176 | שלא |  |
| 6:12.177 | כרצון |  |
| 6:12.178 | חכמים |  |
| 6:12.179 | בבראשית |  |
| 6:12.180 | רבה |  |
| 6:12.181 | הבאתיו |  |
| 6:12.182 | לעיל |  |
| 6:12.183 | סי | ' |
| 6:12.184 | ל"א |  |
| 6:12.185 | כל |  |
| 6:12.186 | שיש |  |
| 6:12.187 | לו |  |
| 6:12.188 | תולדות |  |
| 6:12.189 | וכו | ') |
| 6:12.190 | ויחס |  |
| 6:12.191 | עיקר |  |
| 6:12.192 | העבודה |  |
| 6:12.193 | וקיום |  |
| 6:12.194 | המצות |  |
| 6:12.195 | והברכות |  |
| 6:12.196 | לבן |  |
| 6:12.197 | כנז"ל | . |
| 6:12.198 | ועוד |  |
| 6:12.199 | יתבאר |  |
| 6:12.200 | לקמן | . |
| 6:12.201 | ועל |  |
| 6:12.202 | ידי |  |
| 6:12.203 | רעיונותיהם |  |
| 6:12.204 | העולים |  |
| 6:12.205 | על |  |
| 6:12.206 | רוחם |  |
| 6:12.207 | חשבו |  |
| 6:12.208 | לפרש |  |
| 6:12.209 | הכתובים |  |
| 6:12.210 | שנאמרו |  |
| 6:12.211 | על |  |
| 6:12.212 | דרך |  |
| 6:12.213 | המליצה |  |
| 6:12.214 | והרחבת |  |
| 6:12.215 | הלשון |  |
| 6:12.216 | או |  |
| 6:12.217 | בדרך |  |
| 6:12.218 | חיבה |  |
| 6:12.219 | כמו |  |
| 6:12.220 | שאמרו |  |
| 6:12.221 | רז"ל |  |
| 6:12.222 | לא |  |
| 6:12.223 | זז |  |
| 6:12.224 | מחבבה |  |
| 6:12.225 | וכו | ' |
| 6:12.226 | כמשמען |  |
| 6:12.227 | והנחתן |  |
| 6:12.228 | הראשונה |  |
| 6:12.229 | יד |  |
| 6:12.230 | ה | ', |
| 6:12.231 | עיני |  |
| 6:12.232 | ה | ', |
| 6:12.233 | יעשן |  |
| 6:12.234 | אף |  |
| 6:12.235 | ה | ' |
| 6:12.236 | עלה |  |
| 6:12.237 | עשן |  |
| 6:12.238 | באפו | , |
| 6:12.239 | בצלמנו |  |
| 6:12.240 | כדמותנו | , |
| 6:12.241 | בנים |  |
| 6:12.242 | אתם |  |
| 6:12.243 | לה | ' |
| 6:12.244 | אלהיכם | , |
| 6:12.245 | אחי |  |
| 6:12.246 | ורעי | , |
| 6:12.247 | ויראו |  |
| 6:12.248 | את |  |
| 6:12.249 | אלוהי |  |
| 6:12.250 | ישראל | , |
| 6:12.251 | עורה |  |
| 6:12.252 | למה |  |
| 6:12.253 | תישן |  |
| 6:12.254 | ה | ', |
| 6:12.255 | הנה |  |
| 6:12.256 | לא |  |
| 6:12.257 | ינום |  |
| 6:12.258 | ולא |  |
| 6:12.259 | יישן |  |
| 6:12.260 | שומר |  |
| 6:12.261 | ישראל | , |
| 6:12.262 | וכיוצא |  |
| 6:12.263 | בהן | , |
| 6:12.264 | יפרשו |  |
| 6:12.265 | אותן |  |
| 6:12.266 | על |  |
| 6:12.267 | ידי |  |
| 6:12.268 | רעיונות |  |
| 6:12.269 | אלו |  |
| 6:12.270 | כהנחתן |  |
| 6:12.271 | הראשונה | , |
| 6:12.272 | היפך |  |
| 6:12.273 | מאמר |  |
| 6:12.274 | רז"ל |  |
| 6:12.275 | המתרגם |  |
| 6:12.276 | פסוק |  |
| 6:12.277 | כצורתו |  |
| 6:12.278 | הרי |  |
| 6:12.279 | זה |  |
| 6:12.280 | בדאי | . |
| 6:12.281 | וכמו |  |
| 6:12.282 | שכ | ' |
| 6:12.283 | הרס"ג |  |
| 6:12.284 | בהאמונות |  |
| 6:12.285 | והדעות |  |
| 6:12.286 | מאמר |  |
| 6:12.287 | האחדות | . |
| 6:12.288 | והרמב"ם |  |
| 6:12.289 | במורה |  |
| 6:12.290 | ח"א |  |
| 6:12.291 | וזולתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

