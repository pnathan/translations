# Translation pack: milhamot-hashem perek-5-04 (Milchamot ha-Shem 5:4–5:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-04`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
- **אין סוף / א"ס**: fixed, "Ein Sof", as a name with no article: "האין סוף
  העיקרי" → "the principal Ein Sof"; "כמה מיני אין סוף" (5:2) → "several
  kinds of Ein Sof"; "א"ס ב"ה" → "Ein Sof, blessed be He". Not "the
  Infinite": that makes it a philosophical attribute the author accepts,
  where in the book it is "the First Cause, which they call Ein Sof"
  ("הסבה הראשונה שהוא נקרא אצלם א"ס", 2:11). Give the literal "without end"
  in a note at the first occurrence in a unit.
- **פרצוף / פרצופים**: "partzuf / partzufim" where it means the Kabbalistic
  entities: "פרצוף אחרון שבעולם האצילות" (2:6) → "the last partzuf of the
  World of Emanation"; "חמשה פרצופים" → "five partzufim". In the generic
  sense, paired with צורה or תמונה, it is "face": "אלילי תמונות ופרצופים
  בעלי מפלצות" (1:4) → "idols of images and faces of monstrous form"; "שאין
  לעבוד שום פרצוף וצורה" (2:11) → "that no face or form is to be worshipped".
  Where the author plays on both senses, add a note. Never "configuration" or
  "countenance-structure": that is the defenders' word for them.
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
- ולא אמרינן כלא חד (1:7) and באגרונו (1:12) are for the owner: the main
  English stands as translated; the reviewer's alternative is in `readings`
  / `apparatus`.

## Rulings from wave 1 (chapter 1), by the editor

- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
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
- כלא חד / כולא חד: the same English at every occurrence (1:5-1:7); the
  owner's ruling on the sense is pending; until then, the alternative is in
  `apparatus`.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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


## Milchamot ha-Shem 5:4

Hebrew (Wikisource): מי לא ידע להזהר בכל אלה האזהרות שהזהירו אותנו קדמונינו. פן יתפתה לבבינו אל דברי הכתות החיצונות. כמאמר רז"ל מאן דבעי למהוי חסידא ליקיים מילי דאבות להזהר מכל אמונה חיצונית כוזבת כאלה האמונות שהזכיר ר' חטר ז"ל והזהיר מהן.

| id | word | punctuation after |
|---|---|---|
| 5:4.1 | מי |  |
| 5:4.2 | לא |  |
| 5:4.3 | ידע |  |
| 5:4.4 | להזהר |  |
| 5:4.5 | בכל |  |
| 5:4.6 | אלה |  |
| 5:4.7 | האזהרות |  |
| 5:4.8 | שהזהירו |  |
| 5:4.9 | אותנו |  |
| 5:4.10 | קדמונינו | . |
| 5:4.11 | פן |  |
| 5:4.12 | יתפתה |  |
| 5:4.13 | לבבינו |  |
| 5:4.14 | אל |  |
| 5:4.15 | דברי |  |
| 5:4.16 | הכתות |  |
| 5:4.17 | החיצונות | . |
| 5:4.18 | כמאמר |  |
| 5:4.19 | רז"ל |  |
| 5:4.20 | מאן |  |
| 5:4.21 | דבעי |  |
| 5:4.22 | למהוי |  |
| 5:4.23 | חסידא |  |
| 5:4.24 | ליקיים |  |
| 5:4.25 | מילי |  |
| 5:4.26 | דאבות |  |
| 5:4.27 | להזהר |  |
| 5:4.28 | מכל |  |
| 5:4.29 | אמונה |  |
| 5:4.30 | חיצונית |  |
| 5:4.31 | כוזבת |  |
| 5:4.32 | כאלה |  |
| 5:4.33 | האמונות |  |
| 5:4.34 | שהזכיר |  |
| 5:4.35 | ר | ' |
| 5:4.36 | חטר |  |
| 5:4.37 | ז"ל |  |
| 5:4.38 | והזהיר |  |
| 5:4.39 | מהן | . |


## Milchamot ha-Shem 5:5

Hebrew (Wikisource): מי לא יתן לב להבין ולהשכיל כי הפילסוף נביא השקר מחבר הזהר לקח לו מאלה האמונות החיצונות ובלל ועירב אמונת העכו"ם מאמיני השלוש באמונת תורתינו הקדושה בשמותם שלא כסדר תולדותם אצלם, וסתר פנים ישים להם אב ובן ורוח הקדש בפסוק שמע ישראל, וז"ל הזהר פ' ואתחנן דף רס"ג (ח"ג רסג, א) עם פי' מקדש מלך והגליון. שמע ישראל אמר ר' ייסא ישראל סבא, ר' יצחק אמר ע' רברבא, לאכללא שבעין שמהן מהדותא דכלא (סוד הבינה שיש בה שבעים שמות והכוונה שהמלכות מקבלת מן ע') שמע ישראל כמה דכתיב שמעו שמים וכתיה האזינו השמים. אוף הכא שמע ישראל וכלא חד מלה הוא (פי' דעיר) ידוד רישא דכלא בנהירו דעתיקא קדישא. והאי הוא דאתקרי אב (פי' אבא מקבל מן יוד דאריך) אלהינו עתיקתא דנחלין ומבועין דנפקין וגנדין לכלא (פי' אימא) ידוד גופא דאילנא שלימו דשרשין (פי' זעיר) אחד כנסת ישראל (פי' א"ח הוא תשעה ספירות דזעיר ודל"ת מלכות וזעיר) וכלא חד שלימותא ואתקשר דא בדא ולא אשתבח פירודא אלא כלא חד עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:5.1 | מי |  |
| 5:5.2 | לא |  |
| 5:5.3 | יתן |  |
| 5:5.4 | לב |  |
| 5:5.5 | להבין |  |
| 5:5.6 | ולהשכיל |  |
| 5:5.7 | כי |  |
| 5:5.8 | הפילסוף |  |
| 5:5.9 | נביא |  |
| 5:5.10 | השקר |  |
| 5:5.11 | מחבר |  |
| 5:5.12 | הזהר |  |
| 5:5.13 | לקח |  |
| 5:5.14 | לו |  |
| 5:5.15 | מאלה |  |
| 5:5.16 | האמונות |  |
| 5:5.17 | החיצונות |  |
| 5:5.18 | ובלל |  |
| 5:5.19 | ועירב |  |
| 5:5.20 | אמונת |  |
| 5:5.21 | העכו"ם |  |
| 5:5.22 | מאמיני |  |
| 5:5.23 | השלוש |  |
| 5:5.24 | באמונת |  |
| 5:5.25 | תורתינו |  |
| 5:5.26 | הקדושה |  |
| 5:5.27 | בשמותם |  |
| 5:5.28 | שלא |  |
| 5:5.29 | כסדר |  |
| 5:5.30 | תולדותם |  |
| 5:5.31 | אצלם | , |
| 5:5.32 | וסתר |  |
| 5:5.33 | פנים |  |
| 5:5.34 | ישים |  |
| 5:5.35 | להם |  |
| 5:5.36 | אב |  |
| 5:5.37 | ובן |  |
| 5:5.38 | ורוח |  |
| 5:5.39 | הקדש |  |
| 5:5.40 | בפסוק |  |
| 5:5.41 | שמע |  |
| 5:5.42 | ישראל | , |
| 5:5.43 | וז"ל |  |
| 5:5.44 | הזהר |  |
| 5:5.45 | פ | ' |
| 5:5.46 | ואתחנן |  |
| 5:5.47 | דף |  |
| 5:5.48 | רס"ג |  |
| 5:5.49 | ח"ג |  |
| 5:5.50 | רסג | , |
| 5:5.51 | א | ) |
| 5:5.52 | עם |  |
| 5:5.53 | פי | ' |
| 5:5.54 | מקדש |  |
| 5:5.55 | מלך |  |
| 5:5.56 | והגליון | . |
| 5:5.57 | שמע |  |
| 5:5.58 | ישראל |  |
| 5:5.59 | אמר |  |
| 5:5.60 | ר | ' |
| 5:5.61 | ייסא |  |
| 5:5.62 | ישראל |  |
| 5:5.63 | סבא | , |
| 5:5.64 | ר | ' |
| 5:5.65 | יצחק |  |
| 5:5.66 | אמר |  |
| 5:5.67 | ע | ' |
| 5:5.68 | רברבא | , |
| 5:5.69 | לאכללא |  |
| 5:5.70 | שבעין |  |
| 5:5.71 | שמהן |  |
| 5:5.72 | מהדותא |  |
| 5:5.73 | דכלא |  |
| 5:5.74 | סוד |  |
| 5:5.75 | הבינה |  |
| 5:5.76 | שיש |  |
| 5:5.77 | בה |  |
| 5:5.78 | שבעים |  |
| 5:5.79 | שמות |  |
| 5:5.80 | והכוונה |  |
| 5:5.81 | שהמלכות |  |
| 5:5.82 | מקבלת |  |
| 5:5.83 | מן |  |
| 5:5.84 | ע | ') |
| 5:5.85 | שמע |  |
| 5:5.86 | ישראל |  |
| 5:5.87 | כמה |  |
| 5:5.88 | דכתיב |  |
| 5:5.89 | שמעו |  |
| 5:5.90 | שמים |  |
| 5:5.91 | וכתיה |  |
| 5:5.92 | האזינו |  |
| 5:5.93 | השמים | . |
| 5:5.94 | אוף |  |
| 5:5.95 | הכא |  |
| 5:5.96 | שמע |  |
| 5:5.97 | ישראל |  |
| 5:5.98 | וכלא |  |
| 5:5.99 | חד |  |
| 5:5.100 | מלה |  |
| 5:5.101 | הוא |  |
| 5:5.102 | פי | ' |
| 5:5.103 | דעיר | ) |
| 5:5.104 | ידוד |  |
| 5:5.105 | רישא |  |
| 5:5.106 | דכלא |  |
| 5:5.107 | בנהירו |  |
| 5:5.108 | דעתיקא |  |
| 5:5.109 | קדישא | . |
| 5:5.110 | והאי |  |
| 5:5.111 | הוא |  |
| 5:5.112 | דאתקרי |  |
| 5:5.113 | אב |  |
| 5:5.114 | פי | ' |
| 5:5.115 | אבא |  |
| 5:5.116 | מקבל |  |
| 5:5.117 | מן |  |
| 5:5.118 | יוד |  |
| 5:5.119 | דאריך | ) |
| 5:5.120 | אלהינו |  |
| 5:5.121 | עתיקתא |  |
| 5:5.122 | דנחלין |  |
| 5:5.123 | ומבועין |  |
| 5:5.124 | דנפקין |  |
| 5:5.125 | וגנדין |  |
| 5:5.126 | לכלא |  |
| 5:5.127 | פי | ' |
| 5:5.128 | אימא | ) |
| 5:5.129 | ידוד |  |
| 5:5.130 | גופא |  |
| 5:5.131 | דאילנא |  |
| 5:5.132 | שלימו |  |
| 5:5.133 | דשרשין |  |
| 5:5.134 | פי | ' |
| 5:5.135 | זעיר | ) |
| 5:5.136 | אחד |  |
| 5:5.137 | כנסת |  |
| 5:5.138 | ישראל |  |
| 5:5.139 | פי | ' |
| 5:5.140 | א"ח |  |
| 5:5.141 | הוא |  |
| 5:5.142 | תשעה |  |
| 5:5.143 | ספירות |  |
| 5:5.144 | דזעיר |  |
| 5:5.145 | ודל"ת |  |
| 5:5.146 | מלכות |  |
| 5:5.147 | וזעיר | ) |
| 5:5.148 | וכלא |  |
| 5:5.149 | חד |  |
| 5:5.150 | שלימותא |  |
| 5:5.151 | ואתקשר |  |
| 5:5.152 | דא |  |
| 5:5.153 | בדא |  |
| 5:5.154 | ולא |  |
| 5:5.155 | אשתבח |  |
| 5:5.156 | פירודא |  |
| 5:5.157 | אלא |  |
| 5:5.158 | כלא |  |
| 5:5.159 | חד |  |
| 5:5.160 | עכ"ל | . |


## Milchamot ha-Shem 5:6

Hebrew (Wikisource): הרי לך שהזהר חשב השלשה שמות לשלשה פרצופים מחולקים ידוד ראשונה אבא, אלהינו אמא, ידוד אחרונה זעיר אנפין. אחד כולל זעיר ונוקביה שהם אחד. ולא מתפרשין וכלא חד וכו' שחמה פרצופין הם אחד, ולשיטתיה אזיל שבפרשת בראשית הוא חושב עצם הבורא אבא, ואימא אומנא, וזעיר אנפין בן לאבא ואמא.

| id | word | punctuation after |
|---|---|---|
| 5:6.1 | הרי |  |
| 5:6.2 | לך |  |
| 5:6.3 | שהזהר |  |
| 5:6.4 | חשב |  |
| 5:6.5 | השלשה |  |
| 5:6.6 | שמות |  |
| 5:6.7 | לשלשה |  |
| 5:6.8 | פרצופים |  |
| 5:6.9 | מחולקים |  |
| 5:6.10 | ידוד |  |
| 5:6.11 | ראשונה |  |
| 5:6.12 | אבא | , |
| 5:6.13 | אלהינו |  |
| 5:6.14 | אמא | , |
| 5:6.15 | ידוד |  |
| 5:6.16 | אחרונה |  |
| 5:6.17 | זעיר |  |
| 5:6.18 | אנפין | . |
| 5:6.19 | אחד |  |
| 5:6.20 | כולל |  |
| 5:6.21 | זעיר |  |
| 5:6.22 | ונוקביה |  |
| 5:6.23 | שהם |  |
| 5:6.24 | אחד | . |
| 5:6.25 | ולא |  |
| 5:6.26 | מתפרשין |  |
| 5:6.27 | וכלא |  |
| 5:6.28 | חד |  |
| 5:6.29 | וכו | ' |
| 5:6.30 | שחמה |  |
| 5:6.31 | פרצופין |  |
| 5:6.32 | הם |  |
| 5:6.33 | אחד | , |
| 5:6.34 | ולשיטתיה |  |
| 5:6.35 | אזיל |  |
| 5:6.36 | שבפרשת |  |
| 5:6.37 | בראשית |  |
| 5:6.38 | הוא |  |
| 5:6.39 | חושב |  |
| 5:6.40 | עצם |  |
| 5:6.41 | הבורא |  |
| 5:6.42 | אבא | , |
| 5:6.43 | ואימא |  |
| 5:6.44 | אומנא | , |
| 5:6.45 | וזעיר |  |
| 5:6.46 | אנפין |  |
| 5:6.47 | בן |  |
| 5:6.48 | לאבא |  |
| 5:6.49 | ואמא | . |


## Milchamot ha-Shem 5:7

Hebrew (Wikisource): והעכו"ם חשבו עצם הבירא אריך, וקראו לו אב, וספירת החכמה קראו ל הבן והוא אלעלם בלשון ערבי. ולספירת המינה שהיא אימא קראו לה רוח הקודש. והיא רוח החיים הנקרא בערבי אלחיאה, ובדברי מחבר הזהר אלהים חיים.

| id | word | punctuation after |
|---|---|---|
| 5:7.1 | והעכו"ם |  |
| 5:7.2 | חשבו |  |
| 5:7.3 | עצם |  |
| 5:7.4 | הבירא |  |
| 5:7.5 | אריך | , |
| 5:7.6 | וקראו |  |
| 5:7.7 | לו |  |
| 5:7.8 | אב | , |
| 5:7.9 | וספירת |  |
| 5:7.10 | החכמה |  |
| 5:7.11 | קראו |  |
| 5:7.12 | ל |  |
| 5:7.13 | הבן |  |
| 5:7.14 | והוא |  |
| 5:7.15 | אלעלם |  |
| 5:7.16 | בלשון |  |
| 5:7.17 | ערבי | . |
| 5:7.18 | ולספירת |  |
| 5:7.19 | המינה |  |
| 5:7.20 | שהיא |  |
| 5:7.21 | אימא |  |
| 5:7.22 | קראו |  |
| 5:7.23 | לה |  |
| 5:7.24 | רוח |  |
| 5:7.25 | הקודש | . |
| 5:7.26 | והיא |  |
| 5:7.27 | רוח |  |
| 5:7.28 | החיים |  |
| 5:7.29 | הנקרא |  |
| 5:7.30 | בערבי |  |
| 5:7.31 | אלחיאה | , |
| 5:7.32 | ובדברי |  |
| 5:7.33 | מחבר |  |
| 5:7.34 | הזהר |  |
| 5:7.35 | אלהים |  |
| 5:7.36 | חיים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

