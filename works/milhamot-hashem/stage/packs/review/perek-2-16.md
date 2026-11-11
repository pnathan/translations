# Review pack: milhamot-hashem perek-2-16 (Milchamot ha-Shem 2:32–2:32)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-16`.

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
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
- ס' before a quoted verse is "verse" (סימן = "section" in the author's
  numbered sections); read it from the context and note the choice.
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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


## Milchamot ha-Shem 2:32

Hebrew (Wikisource): והנה אהובי כאחי ז"ל ממשל ארץ ישראל הנמשל, כי המהלך הרב אל מקום השם בהשגת הסוד הוא קוצר המשיג ועומק המושג, ורבוי ההצעות שהם סבות מונעות בכל השגת חכמה, ובפרט בחכמת הקבלה מכמה פנים, ועל כולם החטא שהוא מוסך ומבדיל בין השכל והמושכל וכו' והנה אף כי בעל פרדס רמונים לגודל קדושתו וחסידותו נזהר מאד בכמה מקומות ומזהיר מאד על מחשבה זרה שלפירוד וגבול וממשל משלים וכו' ואף אנו נאמר מי יתן ויערה עלינו רוח ממרום שנשכיל ונבין הדברים בלי נטייה מן האמת ומשרשי הדת וכו'. ובאלו נקוט האיי כללא בידך שכל דברים האסורים בספרי המקובלים שזכרו בלי מחלוקת בדבר ובלי חקירה, ויובנו פשטן של דברים וחיבור הדברים קשרם על לוח לבן ויהיו לטטפות בין עיניך. ואף כי יקרא לפניך דברים הנראים תמוהים מצד עצמן או מצד שנראה חלילה שיש פירוד ומספר וגשמות ושינוי, תקבלם באמונת אומן, ותאמר בלבבך כי קצרה דעתך וידיעתך בעומק הבנה מבלי שתסתפק ח"ו באחד משרשי הדת האמורים ביגדל ובכמה דוכתי, שהרי גם אותם העיקרי אי אפשר לירד לשרשם מצד שכל אנושי וכו'. ומעתה צא ולמד שאסור לחדש ולהבין דבר מתוך הנכתב בספר פן לא יכוין אל האמת, אמנם לבאר פסוקים על פי מה שאמרו המקובלים שהקדמה מוסכמת רק שעל זה מבאר הפ' או אומר שנרמז בו על פי הגימטריא או ראשי וסופי תיבות. טוב ויפה הוא שרק יאמרנו בדרך אפשר לא להחליט הדבר וכו' רק לא יכלה זמנו בהם, ואשר נתן ה' בלבבו והוא מוטעם נכון בלי דוחק אל ישליך אחרי גיוו רק יכתבנו לזכרון ויהיה בעיניו במוצא שלל רב, כי הוא סימן לנשמתו חלק אלוה ממעל שיש לו יניקה ומבוא במקום גבוה, וקדוש יאמר לו, ועתה הנה גליתי לך מצפוני לבי. ואתה המלך בעיונך ודעתך הקדושה והרחבה הודיעני נא את דרכיך, ואבטל דעתי מפני דעתך. יאיר חיים בכרך.

| id | word | punctuation after |
|---|---|---|
| 2:32.1 | והנה |  |
| 2:32.2 | אהובי |  |
| 2:32.3 | כאחי |  |
| 2:32.4 | ז"ל |  |
| 2:32.5 | ממשל |  |
| 2:32.6 | ארץ |  |
| 2:32.7 | ישראל |  |
| 2:32.8 | הנמשל | , |
| 2:32.9 | כי |  |
| 2:32.10 | המהלך |  |
| 2:32.11 | הרב |  |
| 2:32.12 | אל |  |
| 2:32.13 | מקום |  |
| 2:32.14 | השם |  |
| 2:32.15 | בהשגת |  |
| 2:32.16 | הסוד |  |
| 2:32.17 | הוא |  |
| 2:32.18 | קוצר |  |
| 2:32.19 | המשיג |  |
| 2:32.20 | ועומק |  |
| 2:32.21 | המושג | , |
| 2:32.22 | ורבוי |  |
| 2:32.23 | ההצעות |  |
| 2:32.24 | שהם |  |
| 2:32.25 | סבות |  |
| 2:32.26 | מונעות |  |
| 2:32.27 | בכל |  |
| 2:32.28 | השגת |  |
| 2:32.29 | חכמה | , |
| 2:32.30 | ובפרט |  |
| 2:32.31 | בחכמת |  |
| 2:32.32 | הקבלה |  |
| 2:32.33 | מכמה |  |
| 2:32.34 | פנים | , |
| 2:32.35 | ועל |  |
| 2:32.36 | כולם |  |
| 2:32.37 | החטא |  |
| 2:32.38 | שהוא |  |
| 2:32.39 | מוסך |  |
| 2:32.40 | ומבדיל |  |
| 2:32.41 | בין |  |
| 2:32.42 | השכל |  |
| 2:32.43 | והמושכל |  |
| 2:32.44 | וכו | ' |
| 2:32.45 | והנה |  |
| 2:32.46 | אף |  |
| 2:32.47 | כי |  |
| 2:32.48 | בעל |  |
| 2:32.49 | פרדס |  |
| 2:32.50 | רמונים |  |
| 2:32.51 | לגודל |  |
| 2:32.52 | קדושתו |  |
| 2:32.53 | וחסידותו |  |
| 2:32.54 | נזהר |  |
| 2:32.55 | מאד |  |
| 2:32.56 | בכמה |  |
| 2:32.57 | מקומות |  |
| 2:32.58 | ומזהיר |  |
| 2:32.59 | מאד |  |
| 2:32.60 | על |  |
| 2:32.61 | מחשבה |  |
| 2:32.62 | זרה |  |
| 2:32.63 | שלפירוד |  |
| 2:32.64 | וגבול |  |
| 2:32.65 | וממשל |  |
| 2:32.66 | משלים |  |
| 2:32.67 | וכו | ' |
| 2:32.68 | ואף |  |
| 2:32.69 | אנו |  |
| 2:32.70 | נאמר |  |
| 2:32.71 | מי |  |
| 2:32.72 | יתן |  |
| 2:32.73 | ויערה |  |
| 2:32.74 | עלינו |  |
| 2:32.75 | רוח |  |
| 2:32.76 | ממרום |  |
| 2:32.77 | שנשכיל |  |
| 2:32.78 | ונבין |  |
| 2:32.79 | הדברים |  |
| 2:32.80 | בלי |  |
| 2:32.81 | נטייה |  |
| 2:32.82 | מן |  |
| 2:32.83 | האמת |  |
| 2:32.84 | ומשרשי |  |
| 2:32.85 | הדת |  |
| 2:32.86 | וכו | '. |
| 2:32.87 | ובאלו |  |
| 2:32.88 | נקוט |  |
| 2:32.89 | האיי |  |
| 2:32.90 | כללא |  |
| 2:32.91 | בידך |  |
| 2:32.92 | שכל |  |
| 2:32.93 | דברים |  |
| 2:32.94 | האסורים |  |
| 2:32.95 | בספרי |  |
| 2:32.96 | המקובלים |  |
| 2:32.97 | שזכרו |  |
| 2:32.98 | בלי |  |
| 2:32.99 | מחלוקת |  |
| 2:32.100 | בדבר |  |
| 2:32.101 | ובלי |  |
| 2:32.102 | חקירה | , |
| 2:32.103 | ויובנו |  |
| 2:32.104 | פשטן |  |
| 2:32.105 | של |  |
| 2:32.106 | דברים |  |
| 2:32.107 | וחיבור |  |
| 2:32.108 | הדברים |  |
| 2:32.109 | קשרם |  |
| 2:32.110 | על |  |
| 2:32.111 | לוח |  |
| 2:32.112 | לבן |  |
| 2:32.113 | ויהיו |  |
| 2:32.114 | לטטפות |  |
| 2:32.115 | בין |  |
| 2:32.116 | עיניך | . |
| 2:32.117 | ואף |  |
| 2:32.118 | כי |  |
| 2:32.119 | יקרא |  |
| 2:32.120 | לפניך |  |
| 2:32.121 | דברים |  |
| 2:32.122 | הנראים |  |
| 2:32.123 | תמוהים |  |
| 2:32.124 | מצד |  |
| 2:32.125 | עצמן |  |
| 2:32.126 | או |  |
| 2:32.127 | מצד |  |
| 2:32.128 | שנראה |  |
| 2:32.129 | חלילה |  |
| 2:32.130 | שיש |  |
| 2:32.131 | פירוד |  |
| 2:32.132 | ומספר |  |
| 2:32.133 | וגשמות |  |
| 2:32.134 | ושינוי | , |
| 2:32.135 | תקבלם |  |
| 2:32.136 | באמונת |  |
| 2:32.137 | אומן | , |
| 2:32.138 | ותאמר |  |
| 2:32.139 | בלבבך |  |
| 2:32.140 | כי |  |
| 2:32.141 | קצרה |  |
| 2:32.142 | דעתך |  |
| 2:32.143 | וידיעתך |  |
| 2:32.144 | בעומק |  |
| 2:32.145 | הבנה |  |
| 2:32.146 | מבלי |  |
| 2:32.147 | שתסתפק |  |
| 2:32.148 | ח"ו |  |
| 2:32.149 | באחד |  |
| 2:32.150 | משרשי |  |
| 2:32.151 | הדת |  |
| 2:32.152 | האמורים |  |
| 2:32.153 | ביגדל |  |
| 2:32.154 | ובכמה |  |
| 2:32.155 | דוכתי | , |
| 2:32.156 | שהרי |  |
| 2:32.157 | גם |  |
| 2:32.158 | אותם |  |
| 2:32.159 | העיקרי |  |
| 2:32.160 | אי |  |
| 2:32.161 | אפשר |  |
| 2:32.162 | לירד |  |
| 2:32.163 | לשרשם |  |
| 2:32.164 | מצד |  |
| 2:32.165 | שכל |  |
| 2:32.166 | אנושי |  |
| 2:32.167 | וכו | '. |
| 2:32.168 | ומעתה |  |
| 2:32.169 | צא |  |
| 2:32.170 | ולמד |  |
| 2:32.171 | שאסור |  |
| 2:32.172 | לחדש |  |
| 2:32.173 | ולהבין |  |
| 2:32.174 | דבר |  |
| 2:32.175 | מתוך |  |
| 2:32.176 | הנכתב |  |
| 2:32.177 | בספר |  |
| 2:32.178 | פן |  |
| 2:32.179 | לא |  |
| 2:32.180 | יכוין |  |
| 2:32.181 | אל |  |
| 2:32.182 | האמת | , |
| 2:32.183 | אמנם |  |
| 2:32.184 | לבאר |  |
| 2:32.185 | פסוקים |  |
| 2:32.186 | על |  |
| 2:32.187 | פי |  |
| 2:32.188 | מה |  |
| 2:32.189 | שאמרו |  |
| 2:32.190 | המקובלים |  |
| 2:32.191 | שהקדמה |  |
| 2:32.192 | מוסכמת |  |
| 2:32.193 | רק |  |
| 2:32.194 | שעל |  |
| 2:32.195 | זה |  |
| 2:32.196 | מבאר |  |
| 2:32.197 | הפ | ' |
| 2:32.198 | או |  |
| 2:32.199 | אומר |  |
| 2:32.200 | שנרמז |  |
| 2:32.201 | בו |  |
| 2:32.202 | על |  |
| 2:32.203 | פי |  |
| 2:32.204 | הגימטריא |  |
| 2:32.205 | או |  |
| 2:32.206 | ראשי |  |
| 2:32.207 | וסופי |  |
| 2:32.208 | תיבות | . |
| 2:32.209 | טוב |  |
| 2:32.210 | ויפה |  |
| 2:32.211 | הוא |  |
| 2:32.212 | שרק |  |
| 2:32.213 | יאמרנו |  |
| 2:32.214 | בדרך |  |
| 2:32.215 | אפשר |  |
| 2:32.216 | לא |  |
| 2:32.217 | להחליט |  |
| 2:32.218 | הדבר |  |
| 2:32.219 | וכו | ' |
| 2:32.220 | רק |  |
| 2:32.221 | לא |  |
| 2:32.222 | יכלה |  |
| 2:32.223 | זמנו |  |
| 2:32.224 | בהם | , |
| 2:32.225 | ואשר |  |
| 2:32.226 | נתן |  |
| 2:32.227 | ה | ' |
| 2:32.228 | בלבבו |  |
| 2:32.229 | והוא |  |
| 2:32.230 | מוטעם |  |
| 2:32.231 | נכון |  |
| 2:32.232 | בלי |  |
| 2:32.233 | דוחק |  |
| 2:32.234 | אל |  |
| 2:32.235 | ישליך |  |
| 2:32.236 | אחרי |  |
| 2:32.237 | גיוו |  |
| 2:32.238 | רק |  |
| 2:32.239 | יכתבנו |  |
| 2:32.240 | לזכרון |  |
| 2:32.241 | ויהיה |  |
| 2:32.242 | בעיניו |  |
| 2:32.243 | במוצא |  |
| 2:32.244 | שלל |  |
| 2:32.245 | רב | , |
| 2:32.246 | כי |  |
| 2:32.247 | הוא |  |
| 2:32.248 | סימן |  |
| 2:32.249 | לנשמתו |  |
| 2:32.250 | חלק |  |
| 2:32.251 | אלוה |  |
| 2:32.252 | ממעל |  |
| 2:32.253 | שיש |  |
| 2:32.254 | לו |  |
| 2:32.255 | יניקה |  |
| 2:32.256 | ומבוא |  |
| 2:32.257 | במקום |  |
| 2:32.258 | גבוה | , |
| 2:32.259 | וקדוש |  |
| 2:32.260 | יאמר |  |
| 2:32.261 | לו | , |
| 2:32.262 | ועתה |  |
| 2:32.263 | הנה |  |
| 2:32.264 | גליתי |  |
| 2:32.265 | לך |  |
| 2:32.266 | מצפוני |  |
| 2:32.267 | לבי | . |
| 2:32.268 | ואתה |  |
| 2:32.269 | המלך |  |
| 2:32.270 | בעיונך |  |
| 2:32.271 | ודעתך |  |
| 2:32.272 | הקדושה |  |
| 2:32.273 | והרחבה |  |
| 2:32.274 | הודיעני |  |
| 2:32.275 | נא |  |
| 2:32.276 | את |  |
| 2:32.277 | דרכיך | , |
| 2:32.278 | ואבטל |  |
| 2:32.279 | דעתי |  |
| 2:32.280 | מפני |  |
| 2:32.281 | דעתך | . |
| 2:32.282 | יאיר |  |
| 2:32.283 | חיים |  |
| 2:32.284 | בכרך | . |


### Translation under review (2:32)

English: And behold, my beloved, like my brother of blessed memory, the ruler of the land of Israel, the Nimshal, for the going far toward the place of the Name in the grasp of the secret is the shortness of the one who grasps and the depth of what is grasped, and the multitude of propositions, which are preventing causes in every grasp of wisdom, and especially in the wisdom of Kabbalah, in several respects, and over all of them is the sin, which obstructs and separates between the intellect and the intelligible, etc., and indeed, even though the author of Pardes Rimonim, because of the greatness of his holiness and his piety, is very careful in many places, and warns greatly against a strange thought that is of division, and limit, and dominion of similitudes, etc., we too shall say: would that a spirit from on high be poured out upon us, that we might grasp and understand the matters without inclining from the truth and from the roots of the religion, etc. And by these, take this rule in your hand: all things that are forbidden in the books of the Kabbalists, which they mention without dispute on the matter and without inquiry, and their plain sense will be understood, and the joining of the things: bind them on a white tablet, and let them be frontlets between your eyes. And even when things are read before you that appear astonishing in themselves, or from the side that it seems, God forbid, that there is division, number, corporeality and change, you shall accept them in the faith of an artisan, and say in your heart that your mind is short, and your knowledge in the depth of understanding, without your doubting, God forbid, one of the roots of the religion that are stated in Yigdal and in several places, for even those fundamental ones, it is not possible to descend to their roots by human reason, etc. And now go and learn that it is forbidden to innovate and to understand a matter from what is written in the book, lest he not aim at the truth, however, to explain verses according to what the Kabbalists said, that the premise is agreed; only that on this he explains the verse, or says that there is an allusion in it, by means of gematria, or the initial and final letters of words. It is good and fine that he only say it as a possibility, without deciding the matter, etc., only let him not waste his time on them, and whatever the LORD has put in his heart, and it is correctly tasted without strain, he should not cast it behind his back, but rather let him write it down for memory, and let it be in his eyes as one who finds great spoil, for it is a sign for his soul, a portion of God from on high, which has nourishment and access in a high place, and a holy one will say to him: and now, behold, I have revealed to you the secrets of my heart. And you, the king, in your reflection and your knowledge, holy and broad, tell me, please, your ways, and I will set aside my opinion before your opinion. Yair Chayyim Bakhrakh.

Literal: and-behold my-beloved like-my-brother of-blessed-memory the-ruler-of the-land-of Israel the-Nimshal for the-going far toward the-place-of the-Name in-the-grasp-of the-secret is the-shortness-of the-one-who-grasps and-the-depth-of what-is-grasped and-the-multitude-of the-propositions which-are causes preventing in-every grasp-of wisdom and-especially in-the-wisdom-of Kabbalah in-several respects and-over all-of-them the-sin which-is obstructing and-separating between the-intellect and-the-intelligible etc. and-indeed even though the-author-of Pardes Rimonim because-of-the-greatness-of his-holiness and-his-piety is-careful very in-many places and-warns greatly against a-thought strange that-is-of-division and-limit and-dominion of-similitudes etc. and-even we shall-say who will-give and-be-poured-out upon-us a-spirit from-on-high that-we-might-grasp and-understand the-matters without inclining from the-truth and-from-the-roots-of the-religion etc. and-by-these take this rule in-your-hand all things that-are-forbidden in-the-books-of the-Kabbalists which-they-mention without dispute on-the-matter and-without inquiry will-be-understood their-plain-sense of the-things and-the-joining-of the-things bind-them on a-tablet white and-let-them-be frontlets between your-eyes and-even when are-read before-you things that-appear astonishing from-the-side-of themselves or from-the-side-of that-it-seems God-forbid that-there-is division and-number and-corporeality and-change you-shall-accept-them in-the-faith-of an-artisan and-say in-your-heart that is-short your-mind and-your-knowledge in-the-depth-of understanding without your-doubting God-forbid one-of the-roots-of the-religion that-are-stated in-Yigdal and-in-several places for even those the-fundamental-ones it-is-not possible to-descend to-their-roots from-the-side-of reason human etc. and-now go and-learn that-it-is-forbidden to-innovate and-to-understand a-matter from-within what-is-written in-the-book lest not aim at the-truth however to-explain verses by the-mouth-of what they-said the-Kabbalists that-the-premise is-agreed only that-on this he-explains the-verse or says that-there-is-an-allusion in-it by means-of gematria or the-initials-of and-the-finals-of words it-is-good and-fine it-is that-only he-say-it as-a-way a-possibility without deciding the-matter etc. only not let-him-waste his-time on-them and-whatever has-put the-LORD in-his-heart and-it is-correctly-tasted rightly without strain he-should-not cast-it behind his-back but-rather let-him-write-it for-memory and-let-it-be in-his-eyes as-one-who-finds spoil great for it is-a-sign for-his-soul a-portion of-God from-on-high which-has to-it nourishment and-access in-a-place high and-a-holy-one will-say to-him and-now behold I-have-revealed to-you the-secrets-of my-heart and-you the-king in-your-reflection and-your-knowledge holy and-broad tell-me please (object-marker) your-ways and-I-will-set-aside my-opinion before your-opinion Yair Chayyim Bakhrakh

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:32.1 | and behold |  |  | והנה: 'and behold'; at 2:32.45 the same word opens the next clause, 'and indeed'. |
| 2:32.2 | my beloved |  |  |  |
| 2:32.3 | like my brother |  |  |  |
| 2:32.4 | of blessed memory |  |  |  |
| 2:32.5 | the ruler of |  |  | ממשל: a noun 'rule' or the word mashal (parable). The epithet plays on mashal and nimshal. Read as 'ruler' here; the parable reading is in readings. |
| 2:32.6 | the land of |  |  |  |
| 2:32.7 | Israel |  |  |  |
| 2:32.8 | the Nimshal |  |  | הנמשל: the referent of a parable (nimshal). The epithet pairs it with ממשל. |
| 2:32.9 | for |  |  |  |
| 2:32.10 | the going |  |  |  |
| 2:32.11 | far |  |  | הרב: 'much, many', here 'far'. |
| 2:32.12 | toward |  |  |  |
| 2:32.13 | the place of |  |  |  |
| 2:32.14 | the Name |  |  |  |
| 2:32.15 | in the grasp of |  |  |  |
| 2:32.16 | the secret |  |  |  |
| 2:32.17 | is |  |  |  |
| 2:32.18 | the shortness of |  |  |  |
| 2:32.19 | the one who grasps |  |  |  |
| 2:32.20 | and the depth of |  |  |  |
| 2:32.21 | what is grasped |  |  |  |
| 2:32.22 | and the multitude of |  |  |  |
| 2:32.23 | the propositions |  |  | הצעות: propositions or premises put forward for the mind. |
| 2:32.24 | which are |  |  |  |
| 2:32.25 | causes |  |  | סבות: causes. The following adjective, מונעות, is 'preventing'. |
| 2:32.26 | preventing |  |  |  |
| 2:32.27 | in every |  |  |  |
| 2:32.28 | grasp of |  |  |  |
| 2:32.29 | wisdom |  |  |  |
| 2:32.30 | and especially |  |  |  |
| 2:32.31 | in the wisdom of |  |  |  |
| 2:32.32 | Kabbalah |  |  |  |
| 2:32.33 | in several |  |  |  |
| 2:32.34 | respects |  |  | פנים: literally 'faces'; here 'respects' (aspects). |
| 2:32.35 | and over |  |  |  |
| 2:32.36 | all of them |  |  |  |
| 2:32.37 | the sin |  |  |  |
| 2:32.38 | which is |  |  |  |
| 2:32.39 | obstructing |  |  | מוסך: obstructs, stands in the way. |
| 2:32.40 | and separating |  |  |  |
| 2:32.41 | between |  |  |  |
| 2:32.42 | the intellect |  |  |  |
| 2:32.43 | and the intelligible |  |  |  |
| 2:32.44 | etc. |  |  | וכו': 'etc.' marks an omitted continuation in the author's own text. |
| 2:32.45 | and indeed |  |  |  |
| 2:32.46 | even |  |  |  |
| 2:32.47 | though |  |  |  |
| 2:32.48 | the author of |  |  |  |
| 2:32.49 | Pardes |  |  | פרדס רמונים: Pardes Rimonim, 'Pomegranate Orchard', the title of a Kabbalistic book. בעל: 'the author of'. The text does not name him. |
| 2:32.50 | Rimonim |  |  |  |
| 2:32.51 | because of the greatness of |  |  |  |
| 2:32.52 | his holiness |  |  |  |
| 2:32.53 | and his piety |  |  |  |
| 2:32.54 | is careful |  |  |  |
| 2:32.55 | very |  |  |  |
| 2:32.56 | in many |  |  |  |
| 2:32.57 | places |  |  |  |
| 2:32.58 | and warns |  |  |  |
| 2:32.59 | greatly |  |  |  |
| 2:32.60 | against |  |  |  |
| 2:32.61 | a thought |  |  |  |
| 2:32.62 | strange |  |  | זרה: 'strange, foreign'. |
| 2:32.63 | that is of division |  |  | פירוד: division, separation. |
| 2:32.64 | and limit |  |  |  |
| 2:32.65 | and dominion |  |  |  |
| 2:32.66 | of similitudes |  |  |  |
| 2:32.67 | etc. |  |  | וכו': 'etc.' as at 2:32.44. |
| 2:32.68 | and even |  |  |  |
| 2:32.69 | we |  |  |  |
| 2:32.70 | shall say |  |  |  |
| 2:32.71 | who |  |  | מי יתן: an optative idiom, literally 'who will give', rendered 'would that'. |
| 2:32.72 | will give |  |  |  |
| 2:32.73 | and be poured out |  |  | ויערה: 'and be poured out', a yiqtol with wish sense. Allusion to Isa 32:15 (apparatus). |
| 2:32.74 | upon us |  |  |  |
| 2:32.75 | a spirit |  |  |  |
| 2:32.76 | from on high |  |  |  |
| 2:32.77 | that we might grasp |  |  |  |
| 2:32.78 | and understand |  |  |  |
| 2:32.79 | the matters |  |  |  |
| 2:32.80 | without |  |  |  |
| 2:32.81 | inclining |  |  |  |
| 2:32.82 | from |  |  |  |
| 2:32.83 | the truth |  |  |  |
| 2:32.84 | and from the roots of |  |  |  |
| 2:32.85 | the religion |  |  |  |
| 2:32.86 | etc. |  |  | וכו': 'etc.', and the full stop follows in the Hebrew. |
| 2:32.87 | and by these |  |  |  |
| 2:32.88 | take |  |  |  |
| 2:32.89 | this |  |  | האיי: the spelling of האי, 'this'. |
| 2:32.90 | rule |  |  | כללא: the Aramaic word for 'rule' (kelala). |
| 2:32.91 | in your hand |  |  |  |
| 2:32.92 | all |  |  | שכל: the Aramaic word for 'all' (kol). |
| 2:32.93 | things |  |  |  |
| 2:32.94 | that are forbidden |  |  |  |
| 2:32.95 | in the books of |  |  |  |
| 2:32.96 | the Kabbalists |  |  |  |
| 2:32.97 | which they mention |  |  |  |
| 2:32.98 | without |  |  |  |
| 2:32.99 | dispute |  |  |  |
| 2:32.100 | on the matter |  |  |  |
| 2:32.101 | and without |  |  |  |
| 2:32.102 | inquiry |  |  |  |
| 2:32.103 | will be understood |  |  | ויובנו: passive, 'will be understood'. |
| 2:32.104 | their plain sense |  |  |  |
| 2:32.105 | of |  |  |  |
| 2:32.106 | the things |  |  |  |
| 2:32.107 | and the joining of |  |  |  |
| 2:32.108 | the things |  |  |  |
| 2:32.109 | bind them |  |  | קשרם: imperative, 'bind them'. |
| 2:32.110 | on |  |  |  |
| 2:32.111 | a tablet |  |  |  |
| 2:32.112 | white |  |  | לבן: literally 'white'. Prov 3:3 reads 'the tablet of your heart' (see apparatus); translated as written. |
| 2:32.113 | and let them be |  |  | ויהיו: 'and let them be', a jussive. |
| 2:32.114 | frontlets |  |  |  |
| 2:32.115 | between |  |  |  |
| 2:32.116 | your eyes |  |  |  |
| 2:32.117 | and even |  |  |  |
| 2:32.118 | when |  |  |  |
| 2:32.119 | are read |  |  | יקרא: 'is read', passive. |
| 2:32.120 | before you |  |  |  |
| 2:32.121 | things |  |  |  |
| 2:32.122 | that appear |  |  |  |
| 2:32.123 | astonishing |  |  |  |
| 2:32.124 | from the side of |  |  |  |
| 2:32.125 | themselves |  |  |  |
| 2:32.126 | or |  |  |  |
| 2:32.127 | from the side of |  |  |  |
| 2:32.128 | that it seems |  |  |  |
| 2:32.129 | God forbid |  |  | חלילה: an oath formula, 'far be it'; rendered 'God forbid' as at 2:32.148. |
| 2:32.130 | that there is |  |  |  |
| 2:32.131 | division |  |  |  |
| 2:32.132 | and number |  |  |  |
| 2:32.133 | and corporeality |  |  |  |
| 2:32.134 | and change |  |  |  |
| 2:32.135 | you shall accept them |  |  |  |
| 2:32.136 | in the faith of |  |  | באמונת אומן: literally 'with the faith of an artisan'; here, accepting on trust without proof. |
| 2:32.137 | an artisan |  |  |  |
| 2:32.138 | and say |  |  |  |
| 2:32.139 | in your heart |  |  |  |
| 2:32.140 | that |  |  |  |
| 2:32.141 | is short |  |  |  |
| 2:32.142 | your mind |  |  |  |
| 2:32.143 | and your knowledge |  |  |  |
| 2:32.144 | in the depth of |  |  |  |
| 2:32.145 | understanding |  |  |  |
| 2:32.146 | without |  |  |  |
| 2:32.147 | your doubting |  |  |  |
| 2:32.148 | God forbid |  |  |  |
| 2:32.149 | one of |  |  |  |
| 2:32.150 | the roots of |  |  |  |
| 2:32.151 | the religion |  |  |  |
| 2:32.152 | that are stated |  |  |  |
| 2:32.153 | in Yigdal |  |  | ביגדל: Yigdal, the hymn that lists the principles of faith. |
| 2:32.154 | and in several |  |  |  |
| 2:32.155 | places |  |  |  |
| 2:32.156 | for |  |  |  |
| 2:32.157 | even |  |  |  |
| 2:32.158 | those |  |  |  |
| 2:32.159 | the fundamental ones |  |  | העיקרי: 'the fundamental ones', the principles. |
| 2:32.160 | it is not |  |  |  |
| 2:32.161 | possible |  |  |  |
| 2:32.162 | to descend |  |  |  |
| 2:32.163 | to their roots |  |  |  |
| 2:32.164 | from the side of |  |  |  |
| 2:32.165 | reason |  |  |  |
| 2:32.166 | human |  |  |  |
| 2:32.167 | etc. |  |  |  |
| 2:32.168 | and now |  |  |  |
| 2:32.169 | go |  |  |  |
| 2:32.170 | and learn |  |  |  |
| 2:32.171 | that it is forbidden |  |  |  |
| 2:32.172 | to innovate |  |  |  |
| 2:32.173 | and to understand |  |  |  |
| 2:32.174 | a matter |  |  |  |
| 2:32.175 | from within |  |  |  |
| 2:32.176 | what is written |  |  | הנכתב: passive participle, 'what is written'. |
| 2:32.177 | in the book |  |  |  |
| 2:32.178 | lest |  |  |  |
| 2:32.179 | not |  |  |  |
| 2:32.180 | aim |  |  |  |
| 2:32.181 | at |  |  |  |
| 2:32.182 | the truth |  |  |  |
| 2:32.183 | however |  |  |  |
| 2:32.184 | to explain |  |  |  |
| 2:32.185 | verses |  |  |  |
| 2:32.186 | by |  |  |  |
| 2:32.187 | the mouth of |  |  | פי: literally 'mouth of'; with על, the idiom 'according to'. |
| 2:32.188 | what |  |  |  |
| 2:32.189 | they said |  |  |  |
| 2:32.190 | the Kabbalists |  |  |  |
| 2:32.191 | that the premise |  |  |  |
| 2:32.192 | is agreed |  |  |  |
| 2:32.193 | only |  |  |  |
| 2:32.194 | that on |  |  |  |
| 2:32.195 | this |  |  |  |
| 2:32.196 | he explains |  |  |  |
| 2:32.197 | the verse |  |  | הפ': the verse (Hebrew פסוק). The author does not name the verse. |
| 2:32.198 | or |  |  |  |
| 2:32.199 | says |  |  |  |
| 2:32.200 | that there is an allusion |  |  | שנרמז: 'is alluded to', from remez. |
| 2:32.201 | in it |  |  |  |
| 2:32.202 | by |  |  |  |
| 2:32.203 | means of |  |  | על פי: 'by means of' with gematria. |
| 2:32.204 | gematria |  |  |  |
| 2:32.205 | or |  |  |  |
| 2:32.206 | the initials of |  |  | ראשי: 'the initials of'. |
| 2:32.207 | and the finals of |  |  |  |
| 2:32.208 | words |  |  | תיבות: words. |
| 2:32.209 | it is good |  |  |  |
| 2:32.210 | and fine |  |  |  |
| 2:32.211 | it is |  |  |  |
| 2:32.212 | that only |  |  |  |
| 2:32.213 | he say it |  |  |  |
| 2:32.214 | as a way |  |  |  |
| 2:32.215 | a possibility |  |  | אפשר: 'a possibility'. |
| 2:32.216 | without |  |  |  |
| 2:32.217 | deciding |  |  |  |
| 2:32.218 | the matter |  |  |  |
| 2:32.219 | etc. |  |  |  |
| 2:32.220 | only |  |  |  |
| 2:32.221 | not |  |  |  |
| 2:32.222 | let him waste |  |  | יכלה: 'waste, use up'. |
| 2:32.223 | his time |  |  |  |
| 2:32.224 | on them |  |  |  |
| 2:32.225 | and whatever |  |  |  |
| 2:32.226 | has put |  |  |  |
| 2:32.227 | the LORD |  |  | ה': the Tetragrammaton as written, here the LORD. |
| 2:32.228 | in his heart |  |  |  |
| 2:32.229 | and it |  |  |  |
| 2:32.230 | is correctly tasted |  |  |  |
| 2:32.231 | rightly |  |  |  |
| 2:32.232 | without |  |  |  |
| 2:32.233 | strain |  |  |  |
| 2:32.234 | he should not |  |  |  |
| 2:32.235 | cast it |  |  |  |
| 2:32.236 | behind |  |  |  |
| 2:32.237 | his back |  |  | גיוו: literally 'his back' or 'his inner part'. 'Behind his back' is the idiom for ignoring something. See apparatus. |
| 2:32.238 | but rather |  |  |  |
| 2:32.239 | let him write it |  |  |  |
| 2:32.240 | for memory |  |  |  |
| 2:32.241 | and let it be |  |  |  |
| 2:32.242 | in his eyes |  |  |  |
| 2:32.243 | as one who finds |  |  | במוצא: 'as one who finds'. Allusion to Ps 119:162 (apparatus). |
| 2:32.244 | spoil |  |  |  |
| 2:32.245 | great |  |  |  |
| 2:32.246 | for |  |  |  |
| 2:32.247 | it |  |  |  |
| 2:32.248 | is a sign |  |  |  |
| 2:32.249 | for his soul |  |  |  |
| 2:32.250 | a portion |  |  |  |
| 2:32.251 | of God |  |  | אלוה: 'God' (singular). Allusion to Job 31:2 (apparatus). |
| 2:32.252 | from on high |  |  |  |
| 2:32.253 | which has |  |  |  |
| 2:32.254 | to it |  |  |  |
| 2:32.255 | nourishment |  |  |  |
| 2:32.256 | and access |  |  |  |
| 2:32.257 | in a place |  |  |  |
| 2:32.258 | high |  |  |  |
| 2:32.259 | and a holy one |  |  |  |
| 2:32.260 | will say |  |  |  |
| 2:32.261 | to him |  |  |  |
| 2:32.262 | and now |  |  |  |
| 2:32.263 | behold |  |  |  |
| 2:32.264 | I have revealed |  |  |  |
| 2:32.265 | to you |  |  |  |
| 2:32.266 | the secrets of |  |  |  |
| 2:32.267 | my heart |  |  |  |
| 2:32.268 | and you |  |  |  |
| 2:32.269 | the king |  |  |  |
| 2:32.270 | in your reflection |  |  |  |
| 2:32.271 | and your knowledge |  |  |  |
| 2:32.272 | holy |  |  |  |
| 2:32.273 | and broad |  |  |  |
| 2:32.274 | tell me |  |  |  |
| 2:32.275 | please |  |  |  |
| 2:32.276 | (object marker) |  |  | את: the object marker, with no English counterpart. |
| 2:32.277 | your ways |  |  |  |
| 2:32.278 | and I will set aside |  |  |  |
| 2:32.279 | my opinion |  |  |  |
| 2:32.280 | before |  |  |  |
| 2:32.281 | your opinion |  |  |  |
| 2:32.282 | Yair |  |  |  |
| 2:32.283 | Chayyim |  |  |  |
| 2:32.284 | Bakhrakh |  |  | בכרך: the surname is unvocalized in the text; the English is a transliteration and its vowels are a guess. |

Phrases:
- 2:32.1 2:32.2 2:32.3 2:32.4 2:32.5 2:32.6 2:32.7 2:32.8 | And behold, my beloved, like my brother of blessed memory, the ruler of the land of Israel, the Nimshal, | and-behold my-beloved like-my-brother of-blessed-memory the-ruler-of the-land-of Israel the-Nimshal
- 2:32.9 2:32.10 2:32.11 2:32.12 2:32.13 2:32.14 2:32.15 2:32.16 2:32.17 | for the going far toward the place of the Name in the grasp of the secret is | for the-going far toward the-place-of the-Name in-the-grasp-of the-secret is
- 2:32.18 2:32.19 2:32.20 2:32.21 | the shortness of the one who grasps and the depth of what is grasped, | the-shortness-of the-one-who-grasps and-the-depth-of what-is-grasped
- 2:32.22 2:32.23 2:32.24 2:32.25 2:32.26 2:32.27 2:32.28 2:32.29 | and the multitude of propositions, which are preventing causes in every grasp of wisdom, | and-the-multitude-of the-propositions which-are causes preventing in-every grasp-of wisdom
- 2:32.30 2:32.31 2:32.32 2:32.33 2:32.34 | and especially in the wisdom of Kabbalah, in several respects, | and-especially in-the-wisdom-of Kabbalah in-several respects
- 2:32.35 2:32.36 2:32.37 2:32.38 2:32.39 2:32.40 2:32.41 2:32.42 2:32.43 2:32.44 | and over all of them is the sin, which obstructs and separates between the intellect and the intelligible, etc., | and-over all-of-them the-sin which-is obstructing and-separating between the-intellect and-the-intelligible etc.
- 2:32.45 2:32.46 2:32.47 2:32.48 2:32.49 2:32.50 2:32.51 2:32.52 2:32.53 2:32.54 2:32.55 2:32.56 2:32.57 2:32.58 2:32.59 2:32.60 2:32.61 2:32.62 2:32.63 2:32.64 2:32.65 2:32.66 2:32.67 | and indeed, even though the author of Pardes Rimonim, because of the greatness of his holiness and his piety, is very careful in many places, and warns greatly against a strange thought that is of division, and limit, and dominion of similitudes, etc., | and-indeed even though the-author-of Pardes Rimonim because-of-the-greatness-of his-holiness and-his-piety is-careful very in-many places and-warns greatly against a-thought strange that-is-of-division and-limit and-dominion of-similitudes etc.
- 2:32.68 2:32.69 2:32.70 2:32.71 2:32.72 2:32.73 2:32.74 2:32.75 2:32.76 | we too shall say: would that a spirit from on high be poured out upon us, | and-even we shall-say who will-give and-be-poured-out upon-us a-spirit from-on-high
- 2:32.77 2:32.78 2:32.79 2:32.80 2:32.81 2:32.82 2:32.83 2:32.84 2:32.85 2:32.86 | that we might grasp and understand the matters without inclining from the truth and from the roots of the religion, etc. | that-we-might-grasp and-understand the-matters without inclining from the-truth and-from-the-roots-of the-religion etc.
- 2:32.87 2:32.88 2:32.89 2:32.90 2:32.91 2:32.92 2:32.93 | And by these, take this rule in your hand: all things | and-by-these take this rule in-your-hand all things
- 2:32.94 2:32.95 2:32.96 2:32.97 2:32.98 2:32.99 2:32.100 2:32.101 2:32.102 | that are forbidden in the books of the Kabbalists, which they mention without dispute on the matter and without inquiry, | that-are-forbidden in-the-books-of the-Kabbalists which-they-mention without dispute on-the-matter and-without inquiry
- 2:32.103 2:32.104 2:32.105 2:32.106 2:32.107 2:32.108 | and their plain sense will be understood, and the joining of the things: | will-be-understood their-plain-sense of the-things and-the-joining-of the-things
- 2:32.109 2:32.110 2:32.111 2:32.112 2:32.113 2:32.114 2:32.115 2:32.116 | bind them on a white tablet, and let them be frontlets between your eyes. | bind-them on a-tablet white and-let-them-be frontlets between your-eyes
- 2:32.117 2:32.118 2:32.119 2:32.120 2:32.121 2:32.122 2:32.123 2:32.124 2:32.125 2:32.126 2:32.127 2:32.128 2:32.129 2:32.130 2:32.131 2:32.132 2:32.133 2:32.134 | And even when things are read before you that appear astonishing in themselves, or from the side that it seems, God forbid, that there is division, number, corporeality and change, | and-even when are-read before-you things that-appear astonishing from-the-side-of themselves or from-the-side-of that-it-seems God-forbid that-there-is division and-number and-corporeality and-change
- 2:32.135 2:32.136 2:32.137 2:32.138 2:32.139 | you shall accept them in the faith of an artisan, and say in your heart | you-shall-accept-them in-the-faith-of an-artisan and-say in-your-heart
- 2:32.140 2:32.141 2:32.142 2:32.143 2:32.144 2:32.145 2:32.146 2:32.147 | that your mind is short, and your knowledge in the depth of understanding, without your doubting, | that is-short your-mind and-your-knowledge in-the-depth-of understanding without your-doubting
- 2:32.148 2:32.149 2:32.150 2:32.151 2:32.152 2:32.153 2:32.154 2:32.155 | God forbid, one of the roots of the religion that are stated in Yigdal and in several places, | God-forbid one-of the-roots-of the-religion that-are-stated in-Yigdal and-in-several places
- 2:32.156 2:32.157 2:32.158 2:32.159 2:32.160 2:32.161 2:32.162 2:32.163 2:32.164 2:32.165 2:32.166 2:32.167 | for even those fundamental ones, it is not possible to descend to their roots by human reason, etc. | for even those the-fundamental-ones it-is-not possible to-descend to-their-roots from-the-side-of reason human etc.
- 2:32.168 2:32.169 2:32.170 2:32.171 2:32.172 2:32.173 2:32.174 | And now go and learn that it is forbidden to innovate and to understand a matter | and-now go and-learn that-it-is-forbidden to-innovate and-to-understand a-matter
- 2:32.175 2:32.176 2:32.177 2:32.178 2:32.179 2:32.180 2:32.181 2:32.182 | from what is written in the book, lest he not aim at the truth, | from-within what-is-written in-the-book lest not aim at the-truth
- 2:32.183 2:32.184 2:32.185 2:32.186 2:32.187 2:32.188 2:32.189 2:32.190 | however, to explain verses according to what the Kabbalists said, | however to-explain verses by the-mouth-of what they-said the-Kabbalists
- 2:32.191 2:32.192 2:32.193 2:32.194 2:32.195 2:32.196 2:32.197 | that the premise is agreed; only that on this he explains the verse, | that-the-premise is-agreed only that-on this he-explains the-verse
- 2:32.198 2:32.199 2:32.200 2:32.201 2:32.202 2:32.203 2:32.204 2:32.205 2:32.206 2:32.207 2:32.208 | or says that there is an allusion in it, by means of gematria, or the initial and final letters of words. | or says that-there-is-an-allusion in-it by means-of gematria or the-initials-of and-the-finals-of words
- 2:32.209 2:32.210 2:32.211 2:32.212 2:32.213 2:32.214 2:32.215 2:32.216 2:32.217 2:32.218 2:32.219 | It is good and fine that he only say it as a possibility, without deciding the matter, etc., | it-is-good and-fine it-is that-only he-say-it as-a-way a-possibility without deciding the-matter etc.
- 2:32.220 2:32.221 2:32.222 2:32.223 2:32.224 | only let him not waste his time on them, | only not let-him-waste his-time on-them
- 2:32.225 2:32.226 2:32.227 2:32.228 2:32.229 2:32.230 2:32.231 2:32.232 2:32.233 | and whatever the LORD has put in his heart, and it is correctly tasted without strain, | and-whatever has-put the-LORD in-his-heart and-it is-correctly-tasted rightly without strain
- 2:32.234 2:32.235 2:32.236 2:32.237 | he should not cast it behind his back, | he-should-not cast-it behind his-back
- 2:32.238 2:32.239 2:32.240 2:32.241 2:32.242 2:32.243 2:32.244 2:32.245 | but rather let him write it down for memory, and let it be in his eyes as one who finds great spoil, | but-rather let-him-write-it for-memory and-let-it-be in-his-eyes as-one-who-finds spoil great
- 2:32.246 2:32.247 2:32.248 2:32.249 2:32.250 2:32.251 2:32.252 2:32.253 2:32.254 2:32.255 2:32.256 2:32.257 2:32.258 | for it is a sign for his soul, a portion of God from on high, which has nourishment and access in a high place, | for it is-a-sign for-his-soul a-portion of-God from-on-high which-has to-it nourishment and-access in-a-place high
- 2:32.259 2:32.260 2:32.261 2:32.262 2:32.263 2:32.264 2:32.265 | and a holy one will say to him: and now, behold, I have revealed to you | and-a-holy-one will-say to-him and-now behold I-have-revealed to-you
- 2:32.266 2:32.267 | the secrets of my heart. | the-secrets-of my-heart
- 2:32.268 2:32.269 2:32.270 2:32.271 2:32.272 2:32.273 2:32.274 2:32.275 2:32.276 2:32.277 | And you, the king, in your reflection and your knowledge, holy and broad, tell me, please, your ways, | and-you the-king in-your-reflection and-your-knowledge holy and-broad tell-me please (object-marker) your-ways
- 2:32.278 2:32.279 2:32.280 2:32.281 | and I will set aside my opinion before your opinion. | and-I-will-set-aside my-opinion before your-opinion
- 2:32.282 2:32.283 2:32.284 | Yair Chayyim Bakhrakh. | Yair Chayyim Bakhrakh

Grammar notes:
- One long period runs from the opening to 'etc.' (2:32.44). Its subject is 'the going far toward the place of the Name in the grasp of the secret'; its predicate is 'is the shortness of the one who grasps and the depth of what is grasped'. The list that follows (the propositions, the sin) hangs on the same clause.
- 'Even when' (2:32.117-118, ואף כי) and 'even though' (2:32.46-47, והנה אף כי) open concessive clauses. The main clause of the first is at 2:32.135-139, and of the second at 2:32.68-70 'we too shall say'. The English keeps the Hebrew's order, and has no full stop where the Hebrew has none (2:32.44, 2:32.67).
- 'Would that' is the optative 'who will give' (2:32.71-72). 'And be poured out' (2:32.73) is the wish, and 'that we might grasp' (2:32.77) is its purpose clause, with the verb in the form שנשכיל.
- Imperatives and jussives: 'take' (2:32.88), 'go and learn' (2:32.169-170), 'let them be' (2:32.113), 'let him write it' (2:32.239). 'He should not cast' (2:32.234-235) is a negative command with אל.
- Passive forms, each rendered passive: 'will be understood' (2:32.103), 'are read' (2:32.119), 'is correctly tasted' (2:32.230), 'what is written' (2:32.176), 'is agreed' (2:32.192).

Translator's note: Single prose paragraph in the writer's own voice, addressed to a correspondent ('my beloved'; 'you, the king'). The English follows the Hebrew's order. Where the Hebrew has no full stop (2:32.44, 2:32.67), the English has none. Three choices bear on the words and are recorded in apparatus or readings: ממשל (2:32.5, 'ruler' or the parable), גיוו (2:32.237, 'back' or 'inner part'), and the signature surname (2:32.284), which is unvocalized in the text.

Commentary rulings:
- apparatus: Base text: Wikisource. The 1931 scan is not yet collated, and Wikisource additions are translated as base text and flagged here (Q2).
- apparatus: 'ויערה עלינו רוח ממרום' (2:32.73-76): allusion to Isa 32:15, 'until a spirit be poured out upon us from on high'.
- apparatus: 'לטטפות בין עיניך' (2:32.114-116): allusion to Deut 6:8, 'and they shall be frontlets between your eyes'.
- apparatus: 'קשרם על לוח לבן' (2:32.109-112): allusion to Prov 3:3, which reads 'the tablet of your heart' (לוח לבך). The author writes לבן, 'white', and the English follows his wording.
- apparatus: 'כמוצא שלל רב' (2:32.243-245): allusion to Ps 119:162, 'as one who finds great spoil'.
- apparatus: 'חלק אלוה ממעל' (2:32.250-252): allusion to Job 31:2, 'a portion of God from on high'.
- apparatus: 'ממשל' (2:32.5): read here as 'ruler'. The alternative, 'mashal' (parable), gives 'the parable of the land of Israel, the referent'; it is in readings and is not adopted.
- apparatus: 'גיוו' (2:32.237): 'back' or 'inner part'. The English takes the idiom 'behind his back', to ignore; the literal sense is in the token note.
- apparatus: 'הפ'' (2:32.197): read as 'the verse' by the rule for ס' before a verse. The author does not name the verse.
- apparatus: Signature (2:32.282-284): the surname בכרך is unvocalized. 'Bakhrakh' is a guess at the vowels, not a fixed form. The first two names are as written.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

