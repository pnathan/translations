# Translation pack: milhamot-hashem perek-7-09 (Milchamot ha-Shem 7:13–7:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-09`.

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
- **יחוד / ייחוד** (both spellings, one rendering):
  - The doctrine: fixed, "unity". "ייחוד השי"ת" → "the unity of the Name,
    may He be blessed"; "אמונת היחוד" → "the belief in the unity"; "שער
    היחוד" → "the Gate of Unity". Not "monotheism" and not "oneness".
    אחדות → "oneness", kept distinct.
  - The Kabbalistic act: "unification". "לשם יחוד קב"ה ושכינתיה" (6:21) →
    "for the sake of the unification of the Holy One, blessed be He, and His
    Shekhinah"; יחודים → "unifications".
  - The verb: ליחד with God as object, in the rabbinic sense → "to declare
    one": "הזהירנו ליחדו" (5:26) → "commanded us to declare Him one".
    "מיחדים את השי"ת" → "declare the Name, may He be blessed, one". With the
    Kabbalistic entities as object → "to unify": "ליחד זוגתו עמו" (7:5) →
    "to unify his consort with him".
  - The author plays on the two senses ("עיקר היחוד אצלם לחבר ולקשר", 6:1;
    "סוד היחוד הוא ליחד את השם בכל עשר ספירות", 7:25). There, keep "unity" for
    the noun and "unify" for the verb, so that the shared stem shows, and
    explain the play in a note.
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
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
- A reply introduced by ויאמר inside a quotation takes double marks; the
  author's report of a sage's question stays unmarked.
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


## Milchamot ha-Shem 7:13

Hebrew (Wikisource): פו) והעיקר השלישי מס' שבילי אמונה להרחיק ממנו ית' מדת הגוף, כי זה אלוהינו האחד אשר זכרנו אינו גוף. ולא כח בגוף, ולא ישיגוהו משיגי הגוף ומקריו, כגון התנועה והעמידה, לא בעצם ולא במקרה, ועל כן הרחיקו רז"ל ממנו ית' שם החיבור והפירוד, ואמרו אין למעלה לא ישיבה ולא עמידה ולא עורף ולא עיפוי, עורף הוא פירוד, עיפוי הוא חיבור, נגזר מן ועפו בכתף פלשתים, רוצה לומר יהדפום בהדבקם בהם, ויאמר הנביא ואל מי תדמיוני ואשוה, ואלו היה גוף היה דומה לשאר הגופים וכל מה שנמצא בכתוב אשר ידמוהו בדמות הגופות כגון התנועה והעמידה כל זה דרך העברה כמו שאמרו דברה תורה כלשון בני אדם, וזה העיקר השלישי הוא האמור בכתוב כי לא ראיתם כל תמונה, רוצה לומר שלא השגתם לו ית' שום תמונה. כמו שאמרנו שאינו גוף ולא כח בגוף העכ"ל. וז"ל הרמב"ם בפ"א מהל' יסודי התורה המצוי הזה הוא אלהי העולם אדון כל הארץ, והוא המנהיג את הגלגל בכח שאין לו קץ ותכלית, בכח שאין לו הפסק וכו' וידיעת דבר זה מצות עשה שנ' אנכי ה' אלהיך. וכל המעלה על דעתו שיש שם אלוה אחר חוץ מזה, עובר בלא תעשה שנ' לא יהיה לך אלהים אחרים על פני. וכופר בעיקר, שזהו העיקר הגדול שהכל תלוי בו. אלוה זה אחד הוא, ואינו שנים, ולא יותר על שנים, אלא אחד שאין כיחודו אחד מן האחדים הנמצאים בעולם, לא אחד כמין שהוא כולל אחדים רבים ולא כאחד כגוף שהוא נחלק למחלקות ולקצוות, אלא יחוד שאין יחוד אחר בעולם כמותו. אלו היו אלוהות הרבה, היו גופים וגויות, מפני שאין הנמנין השוים במציאותן נפרדין זה מזה אלא במאורעין שיארעו לגופות ולגוית. ואלו היה היוצר גוף וגויה, היה לו קץ ותכלית, שאי אפשר להיות גוף שאין לו קץ, וכל שיש לגופו קץ ותכלית, יש לכחו קץ וסוף, ואלהינו ברוך שמו הואיל וכחו אין לו קץ ואינו פוסק שהרי הגלגל סובב תמיד, אין כחו כח בגוף. והואיל ואינו גוף, לא יארעו לו מאורעות הגופים כדי שיהיה נחלק ונפרד מאחר. לפיכך אי אפשר שיהיה אלא אחד, וידיעת דבר זה מצות עשה שנ' שמע ישראל ה' אלהינו ה' אחד. הרי מפורש בתורה ובנביאים שאין הקב"ה גוף וגויה שנ' כי ה' אלהיכם הוא אלהים שבשמים ממעל ועל הארץ מתחת. והגוף לא יהיה בשני מקומות ונאמר כי לא ראיתם כל תמונה ונאמר ואל מי תדמיוני ואשוה, ואלו היה גוף היה דומה לשאר גופים. אם כן מהו זה שכתוב בתורה ותחת רגליו, כתובים באצבע אלהים, יד ה', עיני ה', אזני ה', וכיצוא בדברים האלו? הכל לפי דעתן שלבני אדם שאין מכירין אלא הגופות, ודברה תורה כלשון בני אדם. והכל כנוים הם שנ' אם שנותי ברק חרבי וכי חרב יש לו ובחרב הוא הורג. אלא משל, והכל משל, ראיה לדבר זה שנביא אחד אומר שראה הקב"ה לבושיה כתלג חיור. ואחר ראהו חמוץ בגדים. משה רבינו עצמו ראהו על הים כגבור עושה מלחמה, ובסיני כשליח צבור עטוף. לומר שאין לו דמות וצורה. אלא הכל במראה הנבואה ובמחזה. אבל אמתת הדבר אין דעתו של אדם מבין ולא יכולה להשיגו ולחקרו, וזהו שאמר הכתוב החקר אלוה תמצא, אם עד תכלית שדי תמצא וכו'. וכיון שנתברר שאינו גוף וגוייה יתברר שלא יארעו מאורעות הגופים, לא חיבור ולא פירוד. לא מקום ולא מידה, לא עלייה ולא ירידה, ולא ימין ולא שמאל, לא פנים ולא אחור, ולא ישיבה ולא עמידה, ואינו מצוי בזמן עד שיהיה לו ראשית ואחרית ומנין שנים. ואינו משתנה, שאין לו דבר שיגרום לו שינוי, ואין לו לא מות ולא חיים כחיי הגוף, ולא סכלות ולא חכמת כחכמת האיש החכם, לא שינה ולא יקיצה, ולא כעס ולא שחוק, ולא שמחה ולא עצבות, ולא שתיקה ולא דיבור כדיבור בני אדם, וכך אמרו חכמים אין למעלה לא ישיבה ולא עמידה לא עורף ולא עיפוי, עד כאן ל' הרמב"ם ז"ל.

| id | word | punctuation after |
|---|---|---|
| 7:13.1 | פו | ) |
| 7:13.2 | והעיקר |  |
| 7:13.3 | השלישי |  |
| 7:13.4 | מס | ' |
| 7:13.5 | שבילי |  |
| 7:13.6 | אמונה |  |
| 7:13.7 | להרחיק |  |
| 7:13.8 | ממנו |  |
| 7:13.9 | ית | ' |
| 7:13.10 | מדת |  |
| 7:13.11 | הגוף | , |
| 7:13.12 | כי |  |
| 7:13.13 | זה |  |
| 7:13.14 | אלוהינו |  |
| 7:13.15 | האחד |  |
| 7:13.16 | אשר |  |
| 7:13.17 | זכרנו |  |
| 7:13.18 | אינו |  |
| 7:13.19 | גוף | . |
| 7:13.20 | ולא |  |
| 7:13.21 | כח |  |
| 7:13.22 | בגוף | , |
| 7:13.23 | ולא |  |
| 7:13.24 | ישיגוהו |  |
| 7:13.25 | משיגי |  |
| 7:13.26 | הגוף |  |
| 7:13.27 | ומקריו | , |
| 7:13.28 | כגון |  |
| 7:13.29 | התנועה |  |
| 7:13.30 | והעמידה | , |
| 7:13.31 | לא |  |
| 7:13.32 | בעצם |  |
| 7:13.33 | ולא |  |
| 7:13.34 | במקרה | , |
| 7:13.35 | ועל |  |
| 7:13.36 | כן |  |
| 7:13.37 | הרחיקו |  |
| 7:13.38 | רז"ל |  |
| 7:13.39 | ממנו |  |
| 7:13.40 | ית | ' |
| 7:13.41 | שם |  |
| 7:13.42 | החיבור |  |
| 7:13.43 | והפירוד | , |
| 7:13.44 | ואמרו |  |
| 7:13.45 | אין |  |
| 7:13.46 | למעלה |  |
| 7:13.47 | לא |  |
| 7:13.48 | ישיבה |  |
| 7:13.49 | ולא |  |
| 7:13.50 | עמידה |  |
| 7:13.51 | ולא |  |
| 7:13.52 | עורף |  |
| 7:13.53 | ולא |  |
| 7:13.54 | עיפוי | , |
| 7:13.55 | עורף |  |
| 7:13.56 | הוא |  |
| 7:13.57 | פירוד | , |
| 7:13.58 | עיפוי |  |
| 7:13.59 | הוא |  |
| 7:13.60 | חיבור | , |
| 7:13.61 | נגזר |  |
| 7:13.62 | מן |  |
| 7:13.63 | ועפו |  |
| 7:13.64 | בכתף |  |
| 7:13.65 | פלשתים | , |
| 7:13.66 | רוצה |  |
| 7:13.67 | לומר |  |
| 7:13.68 | יהדפום |  |
| 7:13.69 | בהדבקם |  |
| 7:13.70 | בהם | , |
| 7:13.71 | ויאמר |  |
| 7:13.72 | הנביא |  |
| 7:13.73 | ואל |  |
| 7:13.74 | מי |  |
| 7:13.75 | תדמיוני |  |
| 7:13.76 | ואשוה | , |
| 7:13.77 | ואלו |  |
| 7:13.78 | היה |  |
| 7:13.79 | גוף |  |
| 7:13.80 | היה |  |
| 7:13.81 | דומה |  |
| 7:13.82 | לשאר |  |
| 7:13.83 | הגופים |  |
| 7:13.84 | וכל |  |
| 7:13.85 | מה |  |
| 7:13.86 | שנמצא |  |
| 7:13.87 | בכתוב |  |
| 7:13.88 | אשר |  |
| 7:13.89 | ידמוהו |  |
| 7:13.90 | בדמות |  |
| 7:13.91 | הגופות |  |
| 7:13.92 | כגון |  |
| 7:13.93 | התנועה |  |
| 7:13.94 | והעמידה |  |
| 7:13.95 | כל |  |
| 7:13.96 | זה |  |
| 7:13.97 | דרך |  |
| 7:13.98 | העברה |  |
| 7:13.99 | כמו |  |
| 7:13.100 | שאמרו |  |
| 7:13.101 | דברה |  |
| 7:13.102 | תורה |  |
| 7:13.103 | כלשון |  |
| 7:13.104 | בני |  |
| 7:13.105 | אדם | , |
| 7:13.106 | וזה |  |
| 7:13.107 | העיקר |  |
| 7:13.108 | השלישי |  |
| 7:13.109 | הוא |  |
| 7:13.110 | האמור |  |
| 7:13.111 | בכתוב |  |
| 7:13.112 | כי |  |
| 7:13.113 | לא |  |
| 7:13.114 | ראיתם |  |
| 7:13.115 | כל |  |
| 7:13.116 | תמונה | , |
| 7:13.117 | רוצה |  |
| 7:13.118 | לומר |  |
| 7:13.119 | שלא |  |
| 7:13.120 | השגתם |  |
| 7:13.121 | לו |  |
| 7:13.122 | ית | ' |
| 7:13.123 | שום |  |
| 7:13.124 | תמונה | . |
| 7:13.125 | כמו |  |
| 7:13.126 | שאמרנו |  |
| 7:13.127 | שאינו |  |
| 7:13.128 | גוף |  |
| 7:13.129 | ולא |  |
| 7:13.130 | כח |  |
| 7:13.131 | בגוף |  |
| 7:13.132 | העכ"ל | . |
| 7:13.133 | וז"ל |  |
| 7:13.134 | הרמב"ם |  |
| 7:13.135 | בפ"א |  |
| 7:13.136 | מהל | ' |
| 7:13.137 | יסודי |  |
| 7:13.138 | התורה |  |
| 7:13.139 | המצוי |  |
| 7:13.140 | הזה |  |
| 7:13.141 | הוא |  |
| 7:13.142 | אלהי |  |
| 7:13.143 | העולם |  |
| 7:13.144 | אדון |  |
| 7:13.145 | כל |  |
| 7:13.146 | הארץ | , |
| 7:13.147 | והוא |  |
| 7:13.148 | המנהיג |  |
| 7:13.149 | את |  |
| 7:13.150 | הגלגל |  |
| 7:13.151 | בכח |  |
| 7:13.152 | שאין |  |
| 7:13.153 | לו |  |
| 7:13.154 | קץ |  |
| 7:13.155 | ותכלית | , |
| 7:13.156 | בכח |  |
| 7:13.157 | שאין |  |
| 7:13.158 | לו |  |
| 7:13.159 | הפסק |  |
| 7:13.160 | וכו | ' |
| 7:13.161 | וידיעת |  |
| 7:13.162 | דבר |  |
| 7:13.163 | זה |  |
| 7:13.164 | מצות |  |
| 7:13.165 | עשה |  |
| 7:13.166 | שנ | ' |
| 7:13.167 | אנכי |  |
| 7:13.168 | ה | ' |
| 7:13.169 | אלהיך | . |
| 7:13.170 | וכל |  |
| 7:13.171 | המעלה |  |
| 7:13.172 | על |  |
| 7:13.173 | דעתו |  |
| 7:13.174 | שיש |  |
| 7:13.175 | שם |  |
| 7:13.176 | אלוה |  |
| 7:13.177 | אחר |  |
| 7:13.178 | חוץ |  |
| 7:13.179 | מזה | , |
| 7:13.180 | עובר |  |
| 7:13.181 | בלא |  |
| 7:13.182 | תעשה |  |
| 7:13.183 | שנ | ' |
| 7:13.184 | לא |  |
| 7:13.185 | יהיה |  |
| 7:13.186 | לך |  |
| 7:13.187 | אלהים |  |
| 7:13.188 | אחרים |  |
| 7:13.189 | על |  |
| 7:13.190 | פני | . |
| 7:13.191 | וכופר |  |
| 7:13.192 | בעיקר | , |
| 7:13.193 | שזהו |  |
| 7:13.194 | העיקר |  |
| 7:13.195 | הגדול |  |
| 7:13.196 | שהכל |  |
| 7:13.197 | תלוי |  |
| 7:13.198 | בו | . |
| 7:13.199 | אלוה |  |
| 7:13.200 | זה |  |
| 7:13.201 | אחד |  |
| 7:13.202 | הוא | , |
| 7:13.203 | ואינו |  |
| 7:13.204 | שנים | , |
| 7:13.205 | ולא |  |
| 7:13.206 | יותר |  |
| 7:13.207 | על |  |
| 7:13.208 | שנים | , |
| 7:13.209 | אלא |  |
| 7:13.210 | אחד |  |
| 7:13.211 | שאין |  |
| 7:13.212 | כיחודו |  |
| 7:13.213 | אחד |  |
| 7:13.214 | מן |  |
| 7:13.215 | האחדים |  |
| 7:13.216 | הנמצאים |  |
| 7:13.217 | בעולם | , |
| 7:13.218 | לא |  |
| 7:13.219 | אחד |  |
| 7:13.220 | כמין |  |
| 7:13.221 | שהוא |  |
| 7:13.222 | כולל |  |
| 7:13.223 | אחדים |  |
| 7:13.224 | רבים |  |
| 7:13.225 | ולא |  |
| 7:13.226 | כאחד |  |
| 7:13.227 | כגוף |  |
| 7:13.228 | שהוא |  |
| 7:13.229 | נחלק |  |
| 7:13.230 | למחלקות |  |
| 7:13.231 | ולקצוות | , |
| 7:13.232 | אלא |  |
| 7:13.233 | יחוד |  |
| 7:13.234 | שאין |  |
| 7:13.235 | יחוד |  |
| 7:13.236 | אחר |  |
| 7:13.237 | בעולם |  |
| 7:13.238 | כמותו | . |
| 7:13.239 | אלו |  |
| 7:13.240 | היו |  |
| 7:13.241 | אלוהות |  |
| 7:13.242 | הרבה | , |
| 7:13.243 | היו |  |
| 7:13.244 | גופים |  |
| 7:13.245 | וגויות | , |
| 7:13.246 | מפני |  |
| 7:13.247 | שאין |  |
| 7:13.248 | הנמנין |  |
| 7:13.249 | השוים |  |
| 7:13.250 | במציאותן |  |
| 7:13.251 | נפרדין |  |
| 7:13.252 | זה |  |
| 7:13.253 | מזה |  |
| 7:13.254 | אלא |  |
| 7:13.255 | במאורעין |  |
| 7:13.256 | שיארעו |  |
| 7:13.257 | לגופות |  |
| 7:13.258 | ולגוית | . |
| 7:13.259 | ואלו |  |
| 7:13.260 | היה |  |
| 7:13.261 | היוצר |  |
| 7:13.262 | גוף |  |
| 7:13.263 | וגויה | , |
| 7:13.264 | היה |  |
| 7:13.265 | לו |  |
| 7:13.266 | קץ |  |
| 7:13.267 | ותכלית | , |
| 7:13.268 | שאי |  |
| 7:13.269 | אפשר |  |
| 7:13.270 | להיות |  |
| 7:13.271 | גוף |  |
| 7:13.272 | שאין |  |
| 7:13.273 | לו |  |
| 7:13.274 | קץ | , |
| 7:13.275 | וכל |  |
| 7:13.276 | שיש |  |
| 7:13.277 | לגופו |  |
| 7:13.278 | קץ |  |
| 7:13.279 | ותכלית | , |
| 7:13.280 | יש |  |
| 7:13.281 | לכחו |  |
| 7:13.282 | קץ |  |
| 7:13.283 | וסוף | , |
| 7:13.284 | ואלהינו |  |
| 7:13.285 | ברוך |  |
| 7:13.286 | שמו |  |
| 7:13.287 | הואיל |  |
| 7:13.288 | וכחו |  |
| 7:13.289 | אין |  |
| 7:13.290 | לו |  |
| 7:13.291 | קץ |  |
| 7:13.292 | ואינו |  |
| 7:13.293 | פוסק |  |
| 7:13.294 | שהרי |  |
| 7:13.295 | הגלגל |  |
| 7:13.296 | סובב |  |
| 7:13.297 | תמיד | , |
| 7:13.298 | אין |  |
| 7:13.299 | כחו |  |
| 7:13.300 | כח |  |
| 7:13.301 | בגוף | . |
| 7:13.302 | והואיל |  |
| 7:13.303 | ואינו |  |
| 7:13.304 | גוף | , |
| 7:13.305 | לא |  |
| 7:13.306 | יארעו |  |
| 7:13.307 | לו |  |
| 7:13.308 | מאורעות |  |
| 7:13.309 | הגופים |  |
| 7:13.310 | כדי |  |
| 7:13.311 | שיהיה |  |
| 7:13.312 | נחלק |  |
| 7:13.313 | ונפרד |  |
| 7:13.314 | מאחר | . |
| 7:13.315 | לפיכך |  |
| 7:13.316 | אי |  |
| 7:13.317 | אפשר |  |
| 7:13.318 | שיהיה |  |
| 7:13.319 | אלא |  |
| 7:13.320 | אחד | , |
| 7:13.321 | וידיעת |  |
| 7:13.322 | דבר |  |
| 7:13.323 | זה |  |
| 7:13.324 | מצות |  |
| 7:13.325 | עשה |  |
| 7:13.326 | שנ | ' |
| 7:13.327 | שמע |  |
| 7:13.328 | ישראל |  |
| 7:13.329 | ה | ' |
| 7:13.330 | אלהינו |  |
| 7:13.331 | ה | ' |
| 7:13.332 | אחד | . |
| 7:13.333 | הרי |  |
| 7:13.334 | מפורש |  |
| 7:13.335 | בתורה |  |
| 7:13.336 | ובנביאים |  |
| 7:13.337 | שאין |  |
| 7:13.338 | הקב"ה |  |
| 7:13.339 | גוף |  |
| 7:13.340 | וגויה |  |
| 7:13.341 | שנ | ' |
| 7:13.342 | כי |  |
| 7:13.343 | ה | ' |
| 7:13.344 | אלהיכם |  |
| 7:13.345 | הוא |  |
| 7:13.346 | אלהים |  |
| 7:13.347 | שבשמים |  |
| 7:13.348 | ממעל |  |
| 7:13.349 | ועל |  |
| 7:13.350 | הארץ |  |
| 7:13.351 | מתחת | . |
| 7:13.352 | והגוף |  |
| 7:13.353 | לא |  |
| 7:13.354 | יהיה |  |
| 7:13.355 | בשני |  |
| 7:13.356 | מקומות |  |
| 7:13.357 | ונאמר |  |
| 7:13.358 | כי |  |
| 7:13.359 | לא |  |
| 7:13.360 | ראיתם |  |
| 7:13.361 | כל |  |
| 7:13.362 | תמונה |  |
| 7:13.363 | ונאמר |  |
| 7:13.364 | ואל |  |
| 7:13.365 | מי |  |
| 7:13.366 | תדמיוני |  |
| 7:13.367 | ואשוה | , |
| 7:13.368 | ואלו |  |
| 7:13.369 | היה |  |
| 7:13.370 | גוף |  |
| 7:13.371 | היה |  |
| 7:13.372 | דומה |  |
| 7:13.373 | לשאר |  |
| 7:13.374 | גופים | . |
| 7:13.375 | אם |  |
| 7:13.376 | כן |  |
| 7:13.377 | מהו |  |
| 7:13.378 | זה |  |
| 7:13.379 | שכתוב |  |
| 7:13.380 | בתורה |  |
| 7:13.381 | ותחת |  |
| 7:13.382 | רגליו | , |
| 7:13.383 | כתובים |  |
| 7:13.384 | באצבע |  |
| 7:13.385 | אלהים | , |
| 7:13.386 | יד |  |
| 7:13.387 | ה | ', |
| 7:13.388 | עיני |  |
| 7:13.389 | ה | ', |
| 7:13.390 | אזני |  |
| 7:13.391 | ה | ', |
| 7:13.392 | וכיצוא |  |
| 7:13.393 | בדברים |  |
| 7:13.394 | האלו | ? |
| 7:13.395 | הכל |  |
| 7:13.396 | לפי |  |
| 7:13.397 | דעתן |  |
| 7:13.398 | שלבני |  |
| 7:13.399 | אדם |  |
| 7:13.400 | שאין |  |
| 7:13.401 | מכירין |  |
| 7:13.402 | אלא |  |
| 7:13.403 | הגופות | , |
| 7:13.404 | ודברה |  |
| 7:13.405 | תורה |  |
| 7:13.406 | כלשון |  |
| 7:13.407 | בני |  |
| 7:13.408 | אדם | . |
| 7:13.409 | והכל |  |
| 7:13.410 | כנוים |  |
| 7:13.411 | הם |  |
| 7:13.412 | שנ | ' |
| 7:13.413 | אם |  |
| 7:13.414 | שנותי |  |
| 7:13.415 | ברק |  |
| 7:13.416 | חרבי |  |
| 7:13.417 | וכי |  |
| 7:13.418 | חרב |  |
| 7:13.419 | יש |  |
| 7:13.420 | לו |  |
| 7:13.421 | ובחרב |  |
| 7:13.422 | הוא |  |
| 7:13.423 | הורג | . |
| 7:13.424 | אלא |  |
| 7:13.425 | משל | , |
| 7:13.426 | והכל |  |
| 7:13.427 | משל | , |
| 7:13.428 | ראיה |  |
| 7:13.429 | לדבר |  |
| 7:13.430 | זה |  |
| 7:13.431 | שנביא |  |
| 7:13.432 | אחד |  |
| 7:13.433 | אומר |  |
| 7:13.434 | שראה |  |
| 7:13.435 | הקב"ה |  |
| 7:13.436 | לבושיה |  |
| 7:13.437 | כתלג |  |
| 7:13.438 | חיור | . |
| 7:13.439 | ואחר |  |
| 7:13.440 | ראהו |  |
| 7:13.441 | חמוץ |  |
| 7:13.442 | בגדים | . |
| 7:13.443 | משה |  |
| 7:13.444 | רבינו |  |
| 7:13.445 | עצמו |  |
| 7:13.446 | ראהו |  |
| 7:13.447 | על |  |
| 7:13.448 | הים |  |
| 7:13.449 | כגבור |  |
| 7:13.450 | עושה |  |
| 7:13.451 | מלחמה | , |
| 7:13.452 | ובסיני |  |
| 7:13.453 | כשליח |  |
| 7:13.454 | צבור |  |
| 7:13.455 | עטוף | . |
| 7:13.456 | לומר |  |
| 7:13.457 | שאין |  |
| 7:13.458 | לו |  |
| 7:13.459 | דמות |  |
| 7:13.460 | וצורה | . |
| 7:13.461 | אלא |  |
| 7:13.462 | הכל |  |
| 7:13.463 | במראה |  |
| 7:13.464 | הנבואה |  |
| 7:13.465 | ובמחזה | . |
| 7:13.466 | אבל |  |
| 7:13.467 | אמתת |  |
| 7:13.468 | הדבר |  |
| 7:13.469 | אין |  |
| 7:13.470 | דעתו |  |
| 7:13.471 | של |  |
| 7:13.472 | אדם |  |
| 7:13.473 | מבין |  |
| 7:13.474 | ולא |  |
| 7:13.475 | יכולה |  |
| 7:13.476 | להשיגו |  |
| 7:13.477 | ולחקרו | , |
| 7:13.478 | וזהו |  |
| 7:13.479 | שאמר |  |
| 7:13.480 | הכתוב |  |
| 7:13.481 | החקר |  |
| 7:13.482 | אלוה |  |
| 7:13.483 | תמצא | , |
| 7:13.484 | אם |  |
| 7:13.485 | עד |  |
| 7:13.486 | תכלית |  |
| 7:13.487 | שדי |  |
| 7:13.488 | תמצא |  |
| 7:13.489 | וכו | '. |
| 7:13.490 | וכיון |  |
| 7:13.491 | שנתברר |  |
| 7:13.492 | שאינו |  |
| 7:13.493 | גוף |  |
| 7:13.494 | וגוייה |  |
| 7:13.495 | יתברר |  |
| 7:13.496 | שלא |  |
| 7:13.497 | יארעו |  |
| 7:13.498 | מאורעות |  |
| 7:13.499 | הגופים | , |
| 7:13.500 | לא |  |
| 7:13.501 | חיבור |  |
| 7:13.502 | ולא |  |
| 7:13.503 | פירוד | . |
| 7:13.504 | לא |  |
| 7:13.505 | מקום |  |
| 7:13.506 | ולא |  |
| 7:13.507 | מידה | , |
| 7:13.508 | לא |  |
| 7:13.509 | עלייה |  |
| 7:13.510 | ולא |  |
| 7:13.511 | ירידה | , |
| 7:13.512 | ולא |  |
| 7:13.513 | ימין |  |
| 7:13.514 | ולא |  |
| 7:13.515 | שמאל | , |
| 7:13.516 | לא |  |
| 7:13.517 | פנים |  |
| 7:13.518 | ולא |  |
| 7:13.519 | אחור | , |
| 7:13.520 | ולא |  |
| 7:13.521 | ישיבה |  |
| 7:13.522 | ולא |  |
| 7:13.523 | עמידה | , |
| 7:13.524 | ואינו |  |
| 7:13.525 | מצוי |  |
| 7:13.526 | בזמן |  |
| 7:13.527 | עד |  |
| 7:13.528 | שיהיה |  |
| 7:13.529 | לו |  |
| 7:13.530 | ראשית |  |
| 7:13.531 | ואחרית |  |
| 7:13.532 | ומנין |  |
| 7:13.533 | שנים | . |
| 7:13.534 | ואינו |  |
| 7:13.535 | משתנה | , |
| 7:13.536 | שאין |  |
| 7:13.537 | לו |  |
| 7:13.538 | דבר |  |
| 7:13.539 | שיגרום |  |
| 7:13.540 | לו |  |
| 7:13.541 | שינוי | , |
| 7:13.542 | ואין |  |
| 7:13.543 | לו |  |
| 7:13.544 | לא |  |
| 7:13.545 | מות |  |
| 7:13.546 | ולא |  |
| 7:13.547 | חיים |  |
| 7:13.548 | כחיי |  |
| 7:13.549 | הגוף | , |
| 7:13.550 | ולא |  |
| 7:13.551 | סכלות |  |
| 7:13.552 | ולא |  |
| 7:13.553 | חכמת |  |
| 7:13.554 | כחכמת |  |
| 7:13.555 | האיש |  |
| 7:13.556 | החכם | , |
| 7:13.557 | לא |  |
| 7:13.558 | שינה |  |
| 7:13.559 | ולא |  |
| 7:13.560 | יקיצה | , |
| 7:13.561 | ולא |  |
| 7:13.562 | כעס |  |
| 7:13.563 | ולא |  |
| 7:13.564 | שחוק | , |
| 7:13.565 | ולא |  |
| 7:13.566 | שמחה |  |
| 7:13.567 | ולא |  |
| 7:13.568 | עצבות | , |
| 7:13.569 | ולא |  |
| 7:13.570 | שתיקה |  |
| 7:13.571 | ולא |  |
| 7:13.572 | דיבור |  |
| 7:13.573 | כדיבור |  |
| 7:13.574 | בני |  |
| 7:13.575 | אדם | , |
| 7:13.576 | וכך |  |
| 7:13.577 | אמרו |  |
| 7:13.578 | חכמים |  |
| 7:13.579 | אין |  |
| 7:13.580 | למעלה |  |
| 7:13.581 | לא |  |
| 7:13.582 | ישיבה |  |
| 7:13.583 | ולא |  |
| 7:13.584 | עמידה |  |
| 7:13.585 | לא |  |
| 7:13.586 | עורף |  |
| 7:13.587 | ולא |  |
| 7:13.588 | עיפוי | , |
| 7:13.589 | עד |  |
| 7:13.590 | כאן |  |
| 7:13.591 | ל | ' |
| 7:13.592 | הרמב"ם |  |
| 7:13.593 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

