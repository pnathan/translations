# Translation pack: milhamot-hashem perek-6-18 (Milchamot ha-Shem 6:28–6:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-18`.

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
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 6:28

Hebrew (Wikisource): עג) ובירושלמי פ' הרואה גרסינן המינין שאלו את ר' שמלאי כמה אלוהות בראו את העולם, אמר להם ולי אתם שולאים, לכו ושאלו את אדם הראשון, שנ' כי שאל נא לימים ראשונים וגו' אשר ברא אלהים אדם על הארץ, אין כתיב כאן אלא למן היום אשר ברא אלהים אדם על הארץ, אמרו ליה והא כתיב בראשית ברא אלהים, אמר להו וכי כתיב בראו, אין כתיב אלא ברא, אמר ר' שמלאי כל מקום שפקרו המינין תשובתן בצדן, חזרו ושאלו אותו מה ההן דכתיב נעשה אדם בצלמינו כדמותינו, אמר להם ויבראו אלהים את האדם בצלמם אין כתיב כאן, אלא ויברא אלהים את האדם בצלמו, אמרו לו תלמידיו לאלה דחית בקנה, לנו מה אתה משיב, אמר להם לשעבר אדם נברא מן העפר, וחוה נבראת מן האדם, מאדם ואילך בצלמינו כדמותינו, אי אפשר לאיש בלא אשה, ואי אפשר לאשה בלא איש, אי אפשר לשניהם בלא שכינה, וחזרו ושאלו אותן, מה ההן דכתיב אל אלהים ה', אל אלהים ה' הוא יודע, אמר להם הם יודעים אין כתיב כאן. אלא הוא יודע כתיב, אמרו לו תלמידי הרבי לאלו דחית בקנה, לנו מה אתה משיב, אמר להן שלשתן שם אחד כאיניש דאמר בסיליוס קיסר אנושתום (אדום מושל ושליט). חזרו ושאלו אותו, מהו דכתיב אל אלהים ה' דבר ויקרא ארץ, אמר להן וכי דברו-ויקראו כתוב כאן, אין כתיב אלא דבר ויקראו ארץ, אמרו לו תלמידיו רבי לאלה דחית בקנה, לנו מה אתה משיב אמר להם שלשתן שם אחד הן כאיניש דאמר אומנון בניןן ארכיטקטנן (נ"א ארכי קטונדס ערוך אמן). חזרו ושאלו אותו מהו דכתיב כי אלהים קדושים הוא אמר להם קדושים המה אין כתוב כאן אלא הוא אל קנא הוא, אמרו לו תלמידיו רבי לאלו דחית בקנה. ולנו מה אתה משיב, אמר להם קדוש הוא בכל מיני קדושות, דאמר רבי יודן בשם רבי אחא הקב"ה דרכו בקדושה, דיבורו בקדושה, ישובו בקדושה, אלהים בקדש דרכך, הילוכו בקדושה, הליכות אל מלכי בקדש, מושבו בקדושה אלהים ישב על כסא קדשו, דיבורו בקדושה אלהים דיבר בקדשו, חשיפת דרועו בקדושה, חשף ה' את זרוע קדשו נורא ואדיר בקדושה, מי כמוך נאדר בקדש, חזרו ושאלו אותו מהו ההן דכתיב מי גוי גדול אשר לו אלהים קרובים אליו. אמר להם כה' אלהינו בכל קראינו אליהם אין כתיב כאן, אלא בכל קראינו אליו. אמרו לו תלמידיו רבי לאלו דחית בקנה. לנו מה אתה משיב, אמר להם קרוב הוא בכל מיני קריבות דאמר ר' פנחס בשפ ר' יהודה בר' סימון. ע"ז נראית קרובה ואינה אלא רחוקה. מה טעם ישאוהו על כתף יסבלוהו וגו' סוף דבר אלוהו עמו בבית והוא צועק עד שימות ולא ישמע ולא יושיע מצרתו. אבל הקב"ה נראה רחוק ואין קרוב ממנו, דאמר רבי לוי מהארץ עד לרקיע מהלך חמש מאות שנה. וכן לכל רקיע ורקיע, ואמר רבי חלבו ורבי ברכיה בשם רבי אבא סמוקא אף טלפי החיות מהלך חמש מאות חמש עשרה מנין ישרה, ראה כמה גבוה הוא מעולמו, ואדם נכנס לבית הכנסת ועומד אחורי העמוד ומתפלל בלחישה והקב"ה מאזין את תפלתו שנאמר וחנה היא מדברת אל לבה רק שפתיה נעות והאזין הקב"ה תפלתה. וכן כל בריותיו שנ' תפלה לעני כי יעטף כאדם המשיח באזן חבירו והוא שומע, וכי יש לך אלוה קרוב משה שהוא קרוב לבריותיו מפה לאזן ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:28.1 | עג | ) |
| 6:28.2 | ובירושלמי |  |
| 6:28.3 | פ | ' |
| 6:28.4 | הרואה |  |
| 6:28.5 | גרסינן |  |
| 6:28.6 | המינין |  |
| 6:28.7 | שאלו |  |
| 6:28.8 | את |  |
| 6:28.9 | ר | ' |
| 6:28.10 | שמלאי |  |
| 6:28.11 | כמה |  |
| 6:28.12 | אלוהות |  |
| 6:28.13 | בראו |  |
| 6:28.14 | את |  |
| 6:28.15 | העולם | , |
| 6:28.16 | אמר |  |
| 6:28.17 | להם |  |
| 6:28.18 | ולי |  |
| 6:28.19 | אתם |  |
| 6:28.20 | שולאים | , |
| 6:28.21 | לכו |  |
| 6:28.22 | ושאלו |  |
| 6:28.23 | את |  |
| 6:28.24 | אדם |  |
| 6:28.25 | הראשון | , |
| 6:28.26 | שנ | ' |
| 6:28.27 | כי |  |
| 6:28.28 | שאל |  |
| 6:28.29 | נא |  |
| 6:28.30 | לימים |  |
| 6:28.31 | ראשונים |  |
| 6:28.32 | וגו | ' |
| 6:28.33 | אשר |  |
| 6:28.34 | ברא |  |
| 6:28.35 | אלהים |  |
| 6:28.36 | אדם |  |
| 6:28.37 | על |  |
| 6:28.38 | הארץ | , |
| 6:28.39 | אין |  |
| 6:28.40 | כתיב |  |
| 6:28.41 | כאן |  |
| 6:28.42 | אלא |  |
| 6:28.43 | למן |  |
| 6:28.44 | היום |  |
| 6:28.45 | אשר |  |
| 6:28.46 | ברא |  |
| 6:28.47 | אלהים |  |
| 6:28.48 | אדם |  |
| 6:28.49 | על |  |
| 6:28.50 | הארץ | , |
| 6:28.51 | אמרו |  |
| 6:28.52 | ליה |  |
| 6:28.53 | והא |  |
| 6:28.54 | כתיב |  |
| 6:28.55 | בראשית |  |
| 6:28.56 | ברא |  |
| 6:28.57 | אלהים | , |
| 6:28.58 | אמר |  |
| 6:28.59 | להו |  |
| 6:28.60 | וכי |  |
| 6:28.61 | כתיב |  |
| 6:28.62 | בראו | , |
| 6:28.63 | אין |  |
| 6:28.64 | כתיב |  |
| 6:28.65 | אלא |  |
| 6:28.66 | ברא | , |
| 6:28.67 | אמר |  |
| 6:28.68 | ר | ' |
| 6:28.69 | שמלאי |  |
| 6:28.70 | כל |  |
| 6:28.71 | מקום |  |
| 6:28.72 | שפקרו |  |
| 6:28.73 | המינין |  |
| 6:28.74 | תשובתן |  |
| 6:28.75 | בצדן | , |
| 6:28.76 | חזרו |  |
| 6:28.77 | ושאלו |  |
| 6:28.78 | אותו |  |
| 6:28.79 | מה |  |
| 6:28.80 | ההן |  |
| 6:28.81 | דכתיב |  |
| 6:28.82 | נעשה |  |
| 6:28.83 | אדם |  |
| 6:28.84 | בצלמינו |  |
| 6:28.85 | כדמותינו | , |
| 6:28.86 | אמר |  |
| 6:28.87 | להם |  |
| 6:28.88 | ויבראו |  |
| 6:28.89 | אלהים |  |
| 6:28.90 | את |  |
| 6:28.91 | האדם |  |
| 6:28.92 | בצלמם |  |
| 6:28.93 | אין |  |
| 6:28.94 | כתיב |  |
| 6:28.95 | כאן | , |
| 6:28.96 | אלא |  |
| 6:28.97 | ויברא |  |
| 6:28.98 | אלהים |  |
| 6:28.99 | את |  |
| 6:28.100 | האדם |  |
| 6:28.101 | בצלמו | , |
| 6:28.102 | אמרו |  |
| 6:28.103 | לו |  |
| 6:28.104 | תלמידיו |  |
| 6:28.105 | לאלה |  |
| 6:28.106 | דחית |  |
| 6:28.107 | בקנה | , |
| 6:28.108 | לנו |  |
| 6:28.109 | מה |  |
| 6:28.110 | אתה |  |
| 6:28.111 | משיב | , |
| 6:28.112 | אמר |  |
| 6:28.113 | להם |  |
| 6:28.114 | לשעבר |  |
| 6:28.115 | אדם |  |
| 6:28.116 | נברא |  |
| 6:28.117 | מן |  |
| 6:28.118 | העפר | , |
| 6:28.119 | וחוה |  |
| 6:28.120 | נבראת |  |
| 6:28.121 | מן |  |
| 6:28.122 | האדם | , |
| 6:28.123 | מאדם |  |
| 6:28.124 | ואילך |  |
| 6:28.125 | בצלמינו |  |
| 6:28.126 | כדמותינו | , |
| 6:28.127 | אי |  |
| 6:28.128 | אפשר |  |
| 6:28.129 | לאיש |  |
| 6:28.130 | בלא |  |
| 6:28.131 | אשה | , |
| 6:28.132 | ואי |  |
| 6:28.133 | אפשר |  |
| 6:28.134 | לאשה |  |
| 6:28.135 | בלא |  |
| 6:28.136 | איש | , |
| 6:28.137 | אי |  |
| 6:28.138 | אפשר |  |
| 6:28.139 | לשניהם |  |
| 6:28.140 | בלא |  |
| 6:28.141 | שכינה | , |
| 6:28.142 | וחזרו |  |
| 6:28.143 | ושאלו |  |
| 6:28.144 | אותן | , |
| 6:28.145 | מה |  |
| 6:28.146 | ההן |  |
| 6:28.147 | דכתיב |  |
| 6:28.148 | אל |  |
| 6:28.149 | אלהים |  |
| 6:28.150 | ה | ', |
| 6:28.151 | אל |  |
| 6:28.152 | אלהים |  |
| 6:28.153 | ה | ' |
| 6:28.154 | הוא |  |
| 6:28.155 | יודע | , |
| 6:28.156 | אמר |  |
| 6:28.157 | להם |  |
| 6:28.158 | הם |  |
| 6:28.159 | יודעים |  |
| 6:28.160 | אין |  |
| 6:28.161 | כתיב |  |
| 6:28.162 | כאן | . |
| 6:28.163 | אלא |  |
| 6:28.164 | הוא |  |
| 6:28.165 | יודע |  |
| 6:28.166 | כתיב | , |
| 6:28.167 | אמרו |  |
| 6:28.168 | לו |  |
| 6:28.169 | תלמידי |  |
| 6:28.170 | הרבי |  |
| 6:28.171 | לאלו |  |
| 6:28.172 | דחית |  |
| 6:28.173 | בקנה | , |
| 6:28.174 | לנו |  |
| 6:28.175 | מה |  |
| 6:28.176 | אתה |  |
| 6:28.177 | משיב | , |
| 6:28.178 | אמר |  |
| 6:28.179 | להן |  |
| 6:28.180 | שלשתן |  |
| 6:28.181 | שם |  |
| 6:28.182 | אחד |  |
| 6:28.183 | כאיניש |  |
| 6:28.184 | דאמר |  |
| 6:28.185 | בסיליוס |  |
| 6:28.186 | קיסר |  |
| 6:28.187 | אנושתום |  |
| 6:28.188 | אדום |  |
| 6:28.189 | מושל |  |
| 6:28.190 | ושליט | ). |
| 6:28.191 | חזרו |  |
| 6:28.192 | ושאלו |  |
| 6:28.193 | אותו | , |
| 6:28.194 | מהו |  |
| 6:28.195 | דכתיב |  |
| 6:28.196 | אל |  |
| 6:28.197 | אלהים |  |
| 6:28.198 | ה | ' |
| 6:28.199 | דבר |  |
| 6:28.200 | ויקרא |  |
| 6:28.201 | ארץ | , |
| 6:28.202 | אמר |  |
| 6:28.203 | להן |  |
| 6:28.204 | וכי |  |
| 6:28.205 | דברו-ויקראו |  |
| 6:28.206 | כתוב |  |
| 6:28.207 | כאן | , |
| 6:28.208 | אין |  |
| 6:28.209 | כתיב |  |
| 6:28.210 | אלא |  |
| 6:28.211 | דבר |  |
| 6:28.212 | ויקראו |  |
| 6:28.213 | ארץ | , |
| 6:28.214 | אמרו |  |
| 6:28.215 | לו |  |
| 6:28.216 | תלמידיו |  |
| 6:28.217 | רבי |  |
| 6:28.218 | לאלה |  |
| 6:28.219 | דחית |  |
| 6:28.220 | בקנה | , |
| 6:28.221 | לנו |  |
| 6:28.222 | מה |  |
| 6:28.223 | אתה |  |
| 6:28.224 | משיב |  |
| 6:28.225 | אמר |  |
| 6:28.226 | להם |  |
| 6:28.227 | שלשתן |  |
| 6:28.228 | שם |  |
| 6:28.229 | אחד |  |
| 6:28.230 | הן |  |
| 6:28.231 | כאיניש |  |
| 6:28.232 | דאמר |  |
| 6:28.233 | אומנון |  |
| 6:28.234 | בניןן |  |
| 6:28.235 | ארכיטקטנן |  |
| 6:28.236 | נ"א |  |
| 6:28.237 | ארכי |  |
| 6:28.238 | קטונדס |  |
| 6:28.239 | ערוך |  |
| 6:28.240 | אמן | ). |
| 6:28.241 | חזרו |  |
| 6:28.242 | ושאלו |  |
| 6:28.243 | אותו |  |
| 6:28.244 | מהו |  |
| 6:28.245 | דכתיב |  |
| 6:28.246 | כי |  |
| 6:28.247 | אלהים |  |
| 6:28.248 | קדושים |  |
| 6:28.249 | הוא |  |
| 6:28.250 | אמר |  |
| 6:28.251 | להם |  |
| 6:28.252 | קדושים |  |
| 6:28.253 | המה |  |
| 6:28.254 | אין |  |
| 6:28.255 | כתוב |  |
| 6:28.256 | כאן |  |
| 6:28.257 | אלא |  |
| 6:28.258 | הוא |  |
| 6:28.259 | אל |  |
| 6:28.260 | קנא |  |
| 6:28.261 | הוא | , |
| 6:28.262 | אמרו |  |
| 6:28.263 | לו |  |
| 6:28.264 | תלמידיו |  |
| 6:28.265 | רבי |  |
| 6:28.266 | לאלו |  |
| 6:28.267 | דחית |  |
| 6:28.268 | בקנה | . |
| 6:28.269 | ולנו |  |
| 6:28.270 | מה |  |
| 6:28.271 | אתה |  |
| 6:28.272 | משיב | , |
| 6:28.273 | אמר |  |
| 6:28.274 | להם |  |
| 6:28.275 | קדוש |  |
| 6:28.276 | הוא |  |
| 6:28.277 | בכל |  |
| 6:28.278 | מיני |  |
| 6:28.279 | קדושות | , |
| 6:28.280 | דאמר |  |
| 6:28.281 | רבי |  |
| 6:28.282 | יודן |  |
| 6:28.283 | בשם |  |
| 6:28.284 | רבי |  |
| 6:28.285 | אחא |  |
| 6:28.286 | הקב"ה |  |
| 6:28.287 | דרכו |  |
| 6:28.288 | בקדושה | , |
| 6:28.289 | דיבורו |  |
| 6:28.290 | בקדושה | , |
| 6:28.291 | ישובו |  |
| 6:28.292 | בקדושה | , |
| 6:28.293 | אלהים |  |
| 6:28.294 | בקדש |  |
| 6:28.295 | דרכך | , |
| 6:28.296 | הילוכו |  |
| 6:28.297 | בקדושה | , |
| 6:28.298 | הליכות |  |
| 6:28.299 | אל |  |
| 6:28.300 | מלכי |  |
| 6:28.301 | בקדש | , |
| 6:28.302 | מושבו |  |
| 6:28.303 | בקדושה |  |
| 6:28.304 | אלהים |  |
| 6:28.305 | ישב |  |
| 6:28.306 | על |  |
| 6:28.307 | כסא |  |
| 6:28.308 | קדשו | , |
| 6:28.309 | דיבורו |  |
| 6:28.310 | בקדושה |  |
| 6:28.311 | אלהים |  |
| 6:28.312 | דיבר |  |
| 6:28.313 | בקדשו | , |
| 6:28.314 | חשיפת |  |
| 6:28.315 | דרועו |  |
| 6:28.316 | בקדושה | , |
| 6:28.317 | חשף |  |
| 6:28.318 | ה | ' |
| 6:28.319 | את |  |
| 6:28.320 | זרוע |  |
| 6:28.321 | קדשו |  |
| 6:28.322 | נורא |  |
| 6:28.323 | ואדיר |  |
| 6:28.324 | בקדושה | , |
| 6:28.325 | מי |  |
| 6:28.326 | כמוך |  |
| 6:28.327 | נאדר |  |
| 6:28.328 | בקדש | , |
| 6:28.329 | חזרו |  |
| 6:28.330 | ושאלו |  |
| 6:28.331 | אותו |  |
| 6:28.332 | מהו |  |
| 6:28.333 | ההן |  |
| 6:28.334 | דכתיב |  |
| 6:28.335 | מי |  |
| 6:28.336 | גוי |  |
| 6:28.337 | גדול |  |
| 6:28.338 | אשר |  |
| 6:28.339 | לו |  |
| 6:28.340 | אלהים |  |
| 6:28.341 | קרובים |  |
| 6:28.342 | אליו | . |
| 6:28.343 | אמר |  |
| 6:28.344 | להם |  |
| 6:28.345 | כה | ' |
| 6:28.346 | אלהינו |  |
| 6:28.347 | בכל |  |
| 6:28.348 | קראינו |  |
| 6:28.349 | אליהם |  |
| 6:28.350 | אין |  |
| 6:28.351 | כתיב |  |
| 6:28.352 | כאן | , |
| 6:28.353 | אלא |  |
| 6:28.354 | בכל |  |
| 6:28.355 | קראינו |  |
| 6:28.356 | אליו | . |
| 6:28.357 | אמרו |  |
| 6:28.358 | לו |  |
| 6:28.359 | תלמידיו |  |
| 6:28.360 | רבי |  |
| 6:28.361 | לאלו |  |
| 6:28.362 | דחית |  |
| 6:28.363 | בקנה | . |
| 6:28.364 | לנו |  |
| 6:28.365 | מה |  |
| 6:28.366 | אתה |  |
| 6:28.367 | משיב | , |
| 6:28.368 | אמר |  |
| 6:28.369 | להם |  |
| 6:28.370 | קרוב |  |
| 6:28.371 | הוא |  |
| 6:28.372 | בכל |  |
| 6:28.373 | מיני |  |
| 6:28.374 | קריבות |  |
| 6:28.375 | דאמר |  |
| 6:28.376 | ר | ' |
| 6:28.377 | פנחס |  |
| 6:28.378 | בשפ |  |
| 6:28.379 | ר | ' |
| 6:28.380 | יהודה |  |
| 6:28.381 | בר | ' |
| 6:28.382 | סימון | . |
| 6:28.383 | ע"ז |  |
| 6:28.384 | נראית |  |
| 6:28.385 | קרובה |  |
| 6:28.386 | ואינה |  |
| 6:28.387 | אלא |  |
| 6:28.388 | רחוקה | . |
| 6:28.389 | מה |  |
| 6:28.390 | טעם |  |
| 6:28.391 | ישאוהו |  |
| 6:28.392 | על |  |
| 6:28.393 | כתף |  |
| 6:28.394 | יסבלוהו |  |
| 6:28.395 | וגו | ' |
| 6:28.396 | סוף |  |
| 6:28.397 | דבר |  |
| 6:28.398 | אלוהו |  |
| 6:28.399 | עמו |  |
| 6:28.400 | בבית |  |
| 6:28.401 | והוא |  |
| 6:28.402 | צועק |  |
| 6:28.403 | עד |  |
| 6:28.404 | שימות |  |
| 6:28.405 | ולא |  |
| 6:28.406 | ישמע |  |
| 6:28.407 | ולא |  |
| 6:28.408 | יושיע |  |
| 6:28.409 | מצרתו | . |
| 6:28.410 | אבל |  |
| 6:28.411 | הקב"ה |  |
| 6:28.412 | נראה |  |
| 6:28.413 | רחוק |  |
| 6:28.414 | ואין |  |
| 6:28.415 | קרוב |  |
| 6:28.416 | ממנו | , |
| 6:28.417 | דאמר |  |
| 6:28.418 | רבי |  |
| 6:28.419 | לוי |  |
| 6:28.420 | מהארץ |  |
| 6:28.421 | עד |  |
| 6:28.422 | לרקיע |  |
| 6:28.423 | מהלך |  |
| 6:28.424 | חמש |  |
| 6:28.425 | מאות |  |
| 6:28.426 | שנה | . |
| 6:28.427 | וכן |  |
| 6:28.428 | לכל |  |
| 6:28.429 | רקיע |  |
| 6:28.430 | ורקיע | , |
| 6:28.431 | ואמר |  |
| 6:28.432 | רבי |  |
| 6:28.433 | חלבו |  |
| 6:28.434 | ורבי |  |
| 6:28.435 | ברכיה |  |
| 6:28.436 | בשם |  |
| 6:28.437 | רבי |  |
| 6:28.438 | אבא |  |
| 6:28.439 | סמוקא |  |
| 6:28.440 | אף |  |
| 6:28.441 | טלפי |  |
| 6:28.442 | החיות |  |
| 6:28.443 | מהלך |  |
| 6:28.444 | חמש |  |
| 6:28.445 | מאות |  |
| 6:28.446 | חמש |  |
| 6:28.447 | עשרה |  |
| 6:28.448 | מנין |  |
| 6:28.449 | ישרה | , |
| 6:28.450 | ראה |  |
| 6:28.451 | כמה |  |
| 6:28.452 | גבוה |  |
| 6:28.453 | הוא |  |
| 6:28.454 | מעולמו | , |
| 6:28.455 | ואדם |  |
| 6:28.456 | נכנס |  |
| 6:28.457 | לבית |  |
| 6:28.458 | הכנסת |  |
| 6:28.459 | ועומד |  |
| 6:28.460 | אחורי |  |
| 6:28.461 | העמוד |  |
| 6:28.462 | ומתפלל |  |
| 6:28.463 | בלחישה |  |
| 6:28.464 | והקב"ה |  |
| 6:28.465 | מאזין |  |
| 6:28.466 | את |  |
| 6:28.467 | תפלתו |  |
| 6:28.468 | שנאמר |  |
| 6:28.469 | וחנה |  |
| 6:28.470 | היא |  |
| 6:28.471 | מדברת |  |
| 6:28.472 | אל |  |
| 6:28.473 | לבה |  |
| 6:28.474 | רק |  |
| 6:28.475 | שפתיה |  |
| 6:28.476 | נעות |  |
| 6:28.477 | והאזין |  |
| 6:28.478 | הקב"ה |  |
| 6:28.479 | תפלתה | . |
| 6:28.480 | וכן |  |
| 6:28.481 | כל |  |
| 6:28.482 | בריותיו |  |
| 6:28.483 | שנ | ' |
| 6:28.484 | תפלה |  |
| 6:28.485 | לעני |  |
| 6:28.486 | כי |  |
| 6:28.487 | יעטף |  |
| 6:28.488 | כאדם |  |
| 6:28.489 | המשיח |  |
| 6:28.490 | באזן |  |
| 6:28.491 | חבירו |  |
| 6:28.492 | והוא |  |
| 6:28.493 | שומע | , |
| 6:28.494 | וכי |  |
| 6:28.495 | יש |  |
| 6:28.496 | לך |  |
| 6:28.497 | אלוה |  |
| 6:28.498 | קרוב |  |
| 6:28.499 | משה |  |
| 6:28.500 | שהוא |  |
| 6:28.501 | קרוב |  |
| 6:28.502 | לבריותיו |  |
| 6:28.503 | מפה |  |
| 6:28.504 | לאזן |  |
| 6:28.505 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

