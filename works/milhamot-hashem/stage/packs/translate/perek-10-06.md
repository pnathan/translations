# Translation pack: milhamot-hashem perek-10-06 (Milchamot ha-Shem 10:7–10:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-10-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-10-06`.

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
- השי"ת → "the Name, may He be blessed". ית׳ / יתברך → "may He be blessed",
  placed where the Hebrew places it ("אחדותו ית'" → "His oneness, may He be
  blessed"). ב"ה → "blessed be He". השם הנכבד → "the glorious Name". שם שמים
  → "the name of Heaven".
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

- **מדה / מדות / מידות**: context. Do not render it "sefirot" even where that
  is the referent.
  - The hermeneutical rules: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי"
    (1:6) → "the thirteen rules of Rabbi Ishmael and the thirty-two rules of
    Rabbi Yose the Galilean"; "שלש עשרה מדות שהתורה נדרשת בהם" (10:5) → "the
    thirteen rules by which the Torah is expounded". Add a note: literally
    "measures".
  - Exod 34:6–7: "שלש עשרה מדות של רחמים" (9:8) → "the thirteen attributes of
    mercy". These are a different thirteen from the rules. The author uses
    both, so never let one rendering cover the other.
  - Kabbalistic and philosophical attributes, God's or the sefirot's: "מדת
    הדין" → "the attribute of judgment"; "המדות העליונות" (2:11) → "the
    upper attributes"; "המידות הפנימיות" (2:7) → "the inner attributes";
    "מדת המלכות" → "the attribute of Malkhut"; "מלובש במידותיו" (5:15) →
    "clothed in His attributes"; "ממדותיו" (7:17, Exod 34:6) → "of His
    attributes".
  - Measure: "עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה" (4:32) → "ten
    measures and ten sefirot, because they have measure and limit"; "לא מקום
    ולא מידה" (7:13) → "neither place nor measure"; "במדה שמדדתם" →
    "with the measure with which you measured".
  - A quality or trait: "ממדת הענוה" (4:8) → "of the quality of humility";
    "המדה המגונה הזאת" (7:19) → "this disgraceful trait"; "אין לך מדה גדולה
    מזו" (6:14, Talmud) → "there is no greater virtue than this".
  - A proper name: מסכת מדות → "tractate Middot".
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
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דמיון: "analogy" for an argument, "likeness" for resemblance.
- דברי תימה: "astonishing words" (words that astonish).
- והוא שלא: "provided that" (fixed).
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
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- Authorities cited by acronym or title (the Rema, Radbaz, Maharshal, Maharam
  al-Shakar, Beit Yosef, Kenesset ha-Gedolah, Maran, the Ra'em): the plain
  transliteration, a note at first occurrence in the unit saying what the
  acronym stands for where the text or the usual expansion is certain; no
  identification beyond that.
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


## Milchamot ha-Shem 10:7

Hebrew (Wikisource): העתק לשון ערבי בסי צ״ו ל' הרס״ג. העתקתו ללשון הקודש. וכמה רחוק ההקש שיהיה זה השטן מלאך, לפי שמוסכם אצלנו כת המיחדים שהבורא ית' ברא מלאכיו קרובים, לפי שגלוי לפניו שלא ימרו את דברו, כמו שבחר בנביאו ששלח, בידעו שלא ימרוהו, והקנאה בצדיקים להרע מהעבירות הגדולות, ומי שחשב שזה המשטין לאיוב היה מלאך, הנה שם המלאכים ממרים, היפך המוסכם אצלינו, ותדע יעזרך השי״ת שיש לנפש ג' כחות, ההשכלה, והקנאה, היא אשר בה יתעורר וישתדל להנקם מאויביו ולמלא מאויו במותרות והתאוה, היא אשר בה ישתוקק הבעל חי למזון ולמשגל ושאר הענינים, ותדע עם זה כי צדקת המשכיל הוא שתהינה שתי המדות הללו רצוני לומר הקנאה והתאוה נכנעים לשכל, ואם יגבר אחד מהם על השכל הנה זה עול וחטאת, והנה מי שחשב שאפשר למלאכים לקטרג שהוא אחד ממיני הקנאה, יתחייב אצלו ליחס להם פעולות התאות, והם האכילה ושתיה והמשגל ושאר הדברים הפחותים ומזולזלים וטנופים, שכבר חסך השי״ת מלאכיו מהם, בידעו שלא ימרוהו, ואי אפשר להיות היפך ממה שקדם בידיעתו ית' שאלו יהיה הענין היפך מה שהנחנו, כבר היה גלוי לפניו ית' שכן יהיה, וכבר הגיעני מקצת החיצונים שחלק בזה עד שחשב שזה השטן מלאך, ושהקנאה והרכילות אפשר להיות בפועל מן המלאכים כמו שפי' במעשה איוב וקנאת השטן לו, וכאשר ראה שאמונתו בזה לא תכלול האכילה והשתיה במלאכים קיים בדעתו ג״כ אפשרותו בהם, שוב ראה עוד שכמו שאלו העניינים אפשר היותם בהםמ יחוייב בהם התאוה והמשגל וקיימם בהם ג״כ, ואמר בענין ויראו בני האלהים את בנות האדם כי טבת הנה, הם המלאכים הבאים על בנות האדם בנפלם מן השמים והיו ככל האדם, והביא דמיון ראיה לדבריו ממה שנא' אני אמרתי אלהים אתם ובני עליון כלכם אכן כאדם תמותון וכאחד השרים תפלו ! ולא השאיר מדה פחותה ומגונה עד שיחסה להם ! כל זה הולך על פי שטתו, והרכיב כל אלה הדעות הזרות והנפסדות ממעוט הבנתו בפירוש הדברים בלה״ק, והנני מבאר טעותו בכל מה שהניח בע״ה, דע יעזרך השם, שמה שהביא לזה המתחיל המגיס דעתו הוא וכיוצא בו לומר שהשטן הוא מלאך, הוא אמרו ית' לשטן הנה כל אשר לו בידיך רק אליו אל תשלח ידך, וגם אמרו אך את נפשו שמור, וגם אמרו ויך את איוב בשחין רע, וחשב שכל זה מיוחס אל השטן, והנני מבאר ביאור מספיק שבו יסורו אלו הספקות הג' אשר דמה ואומר כי הרצון והחפץ בלשון בני ישראל יקרא יד, כמו שאמר דוד לאשה הקובלת לפניו היד יואב איתך בכל זאת ענינו חפץ ורצון יואב וכו' ועל פי זה יהיה אמרו כאן הנה כל אשר לו בידיך בחפצך ורצונך וכו' ובסור ג' ספיקות הללו אומר שטעות זה המפרש בענין ויראו בני האלהים את בנות האדם יותר זר ממה שאמר בענין איוב, והוא שלא ידע כלל ממה שמשתמשין ישראל בלשון בענין עיני העדה וההמוּן, שדרכם לומר מעשה הגדולים, כך ומעשה המון העם כך, וכולם אנשים מבני אדם, ואמנם הבדילם לגדוללתם ומעלתם, ועל פי זה יהיה אמרו ויראו בני האלהים אנשים ג״כ, אבל הבדילם מן ההמון לחשיבותם ומעלתם, וזה ביאור יפה לשלול מהקרובים לשו״ת הפחיתות והזרות, ומה שיחזק זה הפירוש אמרו אחר זה וירא ה' כי רבה רעת האדם, ולא גינה המלאכים. וגם אמרו בסמוך אמחה את האדם ולא ענש המלאכים עמהם ע״כ.

| id | word | punctuation after |
|---|---|---|
| 10:7.1 | העתק |  |
| 10:7.2 | לשון |  |
| 10:7.3 | ערבי |  |
| 10:7.4 | בסי |  |
| 10:7.5 | צ״ו |  |
| 10:7.6 | ל | ' |
| 10:7.7 | הרס״ג | . |
| 10:7.8 | העתקתו |  |
| 10:7.9 | ללשון |  |
| 10:7.10 | הקודש | . |
| 10:7.11 | וכמה |  |
| 10:7.12 | רחוק |  |
| 10:7.13 | ההקש |  |
| 10:7.14 | שיהיה |  |
| 10:7.15 | זה |  |
| 10:7.16 | השטן |  |
| 10:7.17 | מלאך | , |
| 10:7.18 | לפי |  |
| 10:7.19 | שמוסכם |  |
| 10:7.20 | אצלנו |  |
| 10:7.21 | כת |  |
| 10:7.22 | המיחדים |  |
| 10:7.23 | שהבורא |  |
| 10:7.24 | ית | ' |
| 10:7.25 | ברא |  |
| 10:7.26 | מלאכיו |  |
| 10:7.27 | קרובים | , |
| 10:7.28 | לפי |  |
| 10:7.29 | שגלוי |  |
| 10:7.30 | לפניו |  |
| 10:7.31 | שלא |  |
| 10:7.32 | ימרו |  |
| 10:7.33 | את |  |
| 10:7.34 | דברו | , |
| 10:7.35 | כמו |  |
| 10:7.36 | שבחר |  |
| 10:7.37 | בנביאו |  |
| 10:7.38 | ששלח | , |
| 10:7.39 | בידעו |  |
| 10:7.40 | שלא |  |
| 10:7.41 | ימרוהו | , |
| 10:7.42 | והקנאה |  |
| 10:7.43 | בצדיקים |  |
| 10:7.44 | להרע |  |
| 10:7.45 | מהעבירות |  |
| 10:7.46 | הגדולות | , |
| 10:7.47 | ומי |  |
| 10:7.48 | שחשב |  |
| 10:7.49 | שזה |  |
| 10:7.50 | המשטין |  |
| 10:7.51 | לאיוב |  |
| 10:7.52 | היה |  |
| 10:7.53 | מלאך | , |
| 10:7.54 | הנה |  |
| 10:7.55 | שם |  |
| 10:7.56 | המלאכים |  |
| 10:7.57 | ממרים | , |
| 10:7.58 | היפך |  |
| 10:7.59 | המוסכם |  |
| 10:7.60 | אצלינו | , |
| 10:7.61 | ותדע |  |
| 10:7.62 | יעזרך |  |
| 10:7.63 | השי״ת |  |
| 10:7.64 | שיש |  |
| 10:7.65 | לנפש |  |
| 10:7.66 | ג | ' |
| 10:7.67 | כחות | , |
| 10:7.68 | ההשכלה | , |
| 10:7.69 | והקנאה | , |
| 10:7.70 | היא |  |
| 10:7.71 | אשר |  |
| 10:7.72 | בה |  |
| 10:7.73 | יתעורר |  |
| 10:7.74 | וישתדל |  |
| 10:7.75 | להנקם |  |
| 10:7.76 | מאויביו |  |
| 10:7.77 | ולמלא |  |
| 10:7.78 | מאויו |  |
| 10:7.79 | במותרות |  |
| 10:7.80 | והתאוה | , |
| 10:7.81 | היא |  |
| 10:7.82 | אשר |  |
| 10:7.83 | בה |  |
| 10:7.84 | ישתוקק |  |
| 10:7.85 | הבעל |  |
| 10:7.86 | חי |  |
| 10:7.87 | למזון |  |
| 10:7.88 | ולמשגל |  |
| 10:7.89 | ושאר |  |
| 10:7.90 | הענינים | , |
| 10:7.91 | ותדע |  |
| 10:7.92 | עם |  |
| 10:7.93 | זה |  |
| 10:7.94 | כי |  |
| 10:7.95 | צדקת |  |
| 10:7.96 | המשכיל |  |
| 10:7.97 | הוא |  |
| 10:7.98 | שתהינה |  |
| 10:7.99 | שתי |  |
| 10:7.100 | המדות |  |
| 10:7.101 | הללו |  |
| 10:7.102 | רצוני |  |
| 10:7.103 | לומר |  |
| 10:7.104 | הקנאה |  |
| 10:7.105 | והתאוה |  |
| 10:7.106 | נכנעים |  |
| 10:7.107 | לשכל | , |
| 10:7.108 | ואם |  |
| 10:7.109 | יגבר |  |
| 10:7.110 | אחד |  |
| 10:7.111 | מהם |  |
| 10:7.112 | על |  |
| 10:7.113 | השכל |  |
| 10:7.114 | הנה |  |
| 10:7.115 | זה |  |
| 10:7.116 | עול |  |
| 10:7.117 | וחטאת | , |
| 10:7.118 | והנה |  |
| 10:7.119 | מי |  |
| 10:7.120 | שחשב |  |
| 10:7.121 | שאפשר |  |
| 10:7.122 | למלאכים |  |
| 10:7.123 | לקטרג |  |
| 10:7.124 | שהוא |  |
| 10:7.125 | אחד |  |
| 10:7.126 | ממיני |  |
| 10:7.127 | הקנאה | , |
| 10:7.128 | יתחייב |  |
| 10:7.129 | אצלו |  |
| 10:7.130 | ליחס |  |
| 10:7.131 | להם |  |
| 10:7.132 | פעולות |  |
| 10:7.133 | התאות | , |
| 10:7.134 | והם |  |
| 10:7.135 | האכילה |  |
| 10:7.136 | ושתיה |  |
| 10:7.137 | והמשגל |  |
| 10:7.138 | ושאר |  |
| 10:7.139 | הדברים |  |
| 10:7.140 | הפחותים |  |
| 10:7.141 | ומזולזלים |  |
| 10:7.142 | וטנופים | , |
| 10:7.143 | שכבר |  |
| 10:7.144 | חסך |  |
| 10:7.145 | השי״ת |  |
| 10:7.146 | מלאכיו |  |
| 10:7.147 | מהם | , |
| 10:7.148 | בידעו |  |
| 10:7.149 | שלא |  |
| 10:7.150 | ימרוהו | , |
| 10:7.151 | ואי |  |
| 10:7.152 | אפשר |  |
| 10:7.153 | להיות |  |
| 10:7.154 | היפך |  |
| 10:7.155 | ממה |  |
| 10:7.156 | שקדם |  |
| 10:7.157 | בידיעתו |  |
| 10:7.158 | ית | ' |
| 10:7.159 | שאלו |  |
| 10:7.160 | יהיה |  |
| 10:7.161 | הענין |  |
| 10:7.162 | היפך |  |
| 10:7.163 | מה |  |
| 10:7.164 | שהנחנו | , |
| 10:7.165 | כבר |  |
| 10:7.166 | היה |  |
| 10:7.167 | גלוי |  |
| 10:7.168 | לפניו |  |
| 10:7.169 | ית | ' |
| 10:7.170 | שכן |  |
| 10:7.171 | יהיה | , |
| 10:7.172 | וכבר |  |
| 10:7.173 | הגיעני |  |
| 10:7.174 | מקצת |  |
| 10:7.175 | החיצונים |  |
| 10:7.176 | שחלק |  |
| 10:7.177 | בזה |  |
| 10:7.178 | עד |  |
| 10:7.179 | שחשב |  |
| 10:7.180 | שזה |  |
| 10:7.181 | השטן |  |
| 10:7.182 | מלאך | , |
| 10:7.183 | ושהקנאה |  |
| 10:7.184 | והרכילות |  |
| 10:7.185 | אפשר |  |
| 10:7.186 | להיות |  |
| 10:7.187 | בפועל |  |
| 10:7.188 | מן |  |
| 10:7.189 | המלאכים |  |
| 10:7.190 | כמו |  |
| 10:7.191 | שפי | ' |
| 10:7.192 | במעשה |  |
| 10:7.193 | איוב |  |
| 10:7.194 | וקנאת |  |
| 10:7.195 | השטן |  |
| 10:7.196 | לו | , |
| 10:7.197 | וכאשר |  |
| 10:7.198 | ראה |  |
| 10:7.199 | שאמונתו |  |
| 10:7.200 | בזה |  |
| 10:7.201 | לא |  |
| 10:7.202 | תכלול |  |
| 10:7.203 | האכילה |  |
| 10:7.204 | והשתיה |  |
| 10:7.205 | במלאכים |  |
| 10:7.206 | קיים |  |
| 10:7.207 | בדעתו |  |
| 10:7.208 | ג״כ |  |
| 10:7.209 | אפשרותו |  |
| 10:7.210 | בהם | , |
| 10:7.211 | שוב |  |
| 10:7.212 | ראה |  |
| 10:7.213 | עוד |  |
| 10:7.214 | שכמו |  |
| 10:7.215 | שאלו |  |
| 10:7.216 | העניינים |  |
| 10:7.217 | אפשר |  |
| 10:7.218 | היותם |  |
| 10:7.219 | בהםמ |  |
| 10:7.220 | יחוייב |  |
| 10:7.221 | בהם |  |
| 10:7.222 | התאוה |  |
| 10:7.223 | והמשגל |  |
| 10:7.224 | וקיימם |  |
| 10:7.225 | בהם |  |
| 10:7.226 | ג״כ | , |
| 10:7.227 | ואמר |  |
| 10:7.228 | בענין |  |
| 10:7.229 | ויראו |  |
| 10:7.230 | בני |  |
| 10:7.231 | האלהים |  |
| 10:7.232 | את |  |
| 10:7.233 | בנות |  |
| 10:7.234 | האדם |  |
| 10:7.235 | כי |  |
| 10:7.236 | טבת |  |
| 10:7.237 | הנה | , |
| 10:7.238 | הם |  |
| 10:7.239 | המלאכים |  |
| 10:7.240 | הבאים |  |
| 10:7.241 | על |  |
| 10:7.242 | בנות |  |
| 10:7.243 | האדם |  |
| 10:7.244 | בנפלם |  |
| 10:7.245 | מן |  |
| 10:7.246 | השמים |  |
| 10:7.247 | והיו |  |
| 10:7.248 | ככל |  |
| 10:7.249 | האדם | , |
| 10:7.250 | והביא |  |
| 10:7.251 | דמיון |  |
| 10:7.252 | ראיה |  |
| 10:7.253 | לדבריו |  |
| 10:7.254 | ממה |  |
| 10:7.255 | שנא | ' |
| 10:7.256 | אני |  |
| 10:7.257 | אמרתי |  |
| 10:7.258 | אלהים |  |
| 10:7.259 | אתם |  |
| 10:7.260 | ובני |  |
| 10:7.261 | עליון |  |
| 10:7.262 | כלכם |  |
| 10:7.263 | אכן |  |
| 10:7.264 | כאדם |  |
| 10:7.265 | תמותון |  |
| 10:7.266 | וכאחד |  |
| 10:7.267 | השרים |  |
| 10:7.268 | תפלו | ! |
| 10:7.269 | ולא |  |
| 10:7.270 | השאיר |  |
| 10:7.271 | מדה |  |
| 10:7.272 | פחותה |  |
| 10:7.273 | ומגונה |  |
| 10:7.274 | עד |  |
| 10:7.275 | שיחסה |  |
| 10:7.276 | להם | ! |
| 10:7.277 | כל |  |
| 10:7.278 | זה |  |
| 10:7.279 | הולך |  |
| 10:7.280 | על |  |
| 10:7.281 | פי |  |
| 10:7.282 | שטתו | , |
| 10:7.283 | והרכיב |  |
| 10:7.284 | כל |  |
| 10:7.285 | אלה |  |
| 10:7.286 | הדעות |  |
| 10:7.287 | הזרות |  |
| 10:7.288 | והנפסדות |  |
| 10:7.289 | ממעוט |  |
| 10:7.290 | הבנתו |  |
| 10:7.291 | בפירוש |  |
| 10:7.292 | הדברים |  |
| 10:7.293 | בלה״ק | , |
| 10:7.294 | והנני |  |
| 10:7.295 | מבאר |  |
| 10:7.296 | טעותו |  |
| 10:7.297 | בכל |  |
| 10:7.298 | מה |  |
| 10:7.299 | שהניח |  |
| 10:7.300 | בע״ה | , |
| 10:7.301 | דע |  |
| 10:7.302 | יעזרך |  |
| 10:7.303 | השם | , |
| 10:7.304 | שמה |  |
| 10:7.305 | שהביא |  |
| 10:7.306 | לזה |  |
| 10:7.307 | המתחיל |  |
| 10:7.308 | המגיס |  |
| 10:7.309 | דעתו |  |
| 10:7.310 | הוא |  |
| 10:7.311 | וכיוצא |  |
| 10:7.312 | בו |  |
| 10:7.313 | לומר |  |
| 10:7.314 | שהשטן |  |
| 10:7.315 | הוא |  |
| 10:7.316 | מלאך | , |
| 10:7.317 | הוא |  |
| 10:7.318 | אמרו |  |
| 10:7.319 | ית | ' |
| 10:7.320 | לשטן |  |
| 10:7.321 | הנה |  |
| 10:7.322 | כל |  |
| 10:7.323 | אשר |  |
| 10:7.324 | לו |  |
| 10:7.325 | בידיך |  |
| 10:7.326 | רק |  |
| 10:7.327 | אליו |  |
| 10:7.328 | אל |  |
| 10:7.329 | תשלח |  |
| 10:7.330 | ידך | , |
| 10:7.331 | וגם |  |
| 10:7.332 | אמרו |  |
| 10:7.333 | אך |  |
| 10:7.334 | את |  |
| 10:7.335 | נפשו |  |
| 10:7.336 | שמור | , |
| 10:7.337 | וגם |  |
| 10:7.338 | אמרו |  |
| 10:7.339 | ויך |  |
| 10:7.340 | את |  |
| 10:7.341 | איוב |  |
| 10:7.342 | בשחין |  |
| 10:7.343 | רע | , |
| 10:7.344 | וחשב |  |
| 10:7.345 | שכל |  |
| 10:7.346 | זה |  |
| 10:7.347 | מיוחס |  |
| 10:7.348 | אל |  |
| 10:7.349 | השטן | , |
| 10:7.350 | והנני |  |
| 10:7.351 | מבאר |  |
| 10:7.352 | ביאור |  |
| 10:7.353 | מספיק |  |
| 10:7.354 | שבו |  |
| 10:7.355 | יסורו |  |
| 10:7.356 | אלו |  |
| 10:7.357 | הספקות |  |
| 10:7.358 | הג | ' |
| 10:7.359 | אשר |  |
| 10:7.360 | דמה |  |
| 10:7.361 | ואומר |  |
| 10:7.362 | כי |  |
| 10:7.363 | הרצון |  |
| 10:7.364 | והחפץ |  |
| 10:7.365 | בלשון |  |
| 10:7.366 | בני |  |
| 10:7.367 | ישראל |  |
| 10:7.368 | יקרא |  |
| 10:7.369 | יד | , |
| 10:7.370 | כמו |  |
| 10:7.371 | שאמר |  |
| 10:7.372 | דוד |  |
| 10:7.373 | לאשה |  |
| 10:7.374 | הקובלת |  |
| 10:7.375 | לפניו |  |
| 10:7.376 | היד |  |
| 10:7.377 | יואב |  |
| 10:7.378 | איתך |  |
| 10:7.379 | בכל |  |
| 10:7.380 | זאת |  |
| 10:7.381 | ענינו |  |
| 10:7.382 | חפץ |  |
| 10:7.383 | ורצון |  |
| 10:7.384 | יואב |  |
| 10:7.385 | וכו | ' |
| 10:7.386 | ועל |  |
| 10:7.387 | פי |  |
| 10:7.388 | זה |  |
| 10:7.389 | יהיה |  |
| 10:7.390 | אמרו |  |
| 10:7.391 | כאן |  |
| 10:7.392 | הנה |  |
| 10:7.393 | כל |  |
| 10:7.394 | אשר |  |
| 10:7.395 | לו |  |
| 10:7.396 | בידיך |  |
| 10:7.397 | בחפצך |  |
| 10:7.398 | ורצונך |  |
| 10:7.399 | וכו | ' |
| 10:7.400 | ובסור |  |
| 10:7.401 | ג | ' |
| 10:7.402 | ספיקות |  |
| 10:7.403 | הללו |  |
| 10:7.404 | אומר |  |
| 10:7.405 | שטעות |  |
| 10:7.406 | זה |  |
| 10:7.407 | המפרש |  |
| 10:7.408 | בענין |  |
| 10:7.409 | ויראו |  |
| 10:7.410 | בני |  |
| 10:7.411 | האלהים |  |
| 10:7.412 | את |  |
| 10:7.413 | בנות |  |
| 10:7.414 | האדם |  |
| 10:7.415 | יותר |  |
| 10:7.416 | זר |  |
| 10:7.417 | ממה |  |
| 10:7.418 | שאמר |  |
| 10:7.419 | בענין |  |
| 10:7.420 | איוב | , |
| 10:7.421 | והוא |  |
| 10:7.422 | שלא |  |
| 10:7.423 | ידע |  |
| 10:7.424 | כלל |  |
| 10:7.425 | ממה |  |
| 10:7.426 | שמשתמשין |  |
| 10:7.427 | ישראל |  |
| 10:7.428 | בלשון |  |
| 10:7.429 | בענין |  |
| 10:7.430 | עיני |  |
| 10:7.431 | העדה |  |
| 10:7.432 | וההמוּן | , |
| 10:7.433 | שדרכם |  |
| 10:7.434 | לומר |  |
| 10:7.435 | מעשה |  |
| 10:7.436 | הגדולים | , |
| 10:7.437 | כך |  |
| 10:7.438 | ומעשה |  |
| 10:7.439 | המון |  |
| 10:7.440 | העם |  |
| 10:7.441 | כך | , |
| 10:7.442 | וכולם |  |
| 10:7.443 | אנשים |  |
| 10:7.444 | מבני |  |
| 10:7.445 | אדם | , |
| 10:7.446 | ואמנם |  |
| 10:7.447 | הבדילם |  |
| 10:7.448 | לגדוללתם |  |
| 10:7.449 | ומעלתם | , |
| 10:7.450 | ועל |  |
| 10:7.451 | פי |  |
| 10:7.452 | זה |  |
| 10:7.453 | יהיה |  |
| 10:7.454 | אמרו |  |
| 10:7.455 | ויראו |  |
| 10:7.456 | בני |  |
| 10:7.457 | האלהים |  |
| 10:7.458 | אנשים |  |
| 10:7.459 | ג״כ | , |
| 10:7.460 | אבל |  |
| 10:7.461 | הבדילם |  |
| 10:7.462 | מן |  |
| 10:7.463 | ההמון |  |
| 10:7.464 | לחשיבותם |  |
| 10:7.465 | ומעלתם | , |
| 10:7.466 | וזה |  |
| 10:7.467 | ביאור |  |
| 10:7.468 | יפה |  |
| 10:7.469 | לשלול |  |
| 10:7.470 | מהקרובים |  |
| 10:7.471 | לשו״ת |  |
| 10:7.472 | הפחיתות |  |
| 10:7.473 | והזרות | , |
| 10:7.474 | ומה |  |
| 10:7.475 | שיחזק |  |
| 10:7.476 | זה |  |
| 10:7.477 | הפירוש |  |
| 10:7.478 | אמרו |  |
| 10:7.479 | אחר |  |
| 10:7.480 | זה |  |
| 10:7.481 | וירא |  |
| 10:7.482 | ה | ' |
| 10:7.483 | כי |  |
| 10:7.484 | רבה |  |
| 10:7.485 | רעת |  |
| 10:7.486 | האדם | , |
| 10:7.487 | ולא |  |
| 10:7.488 | גינה |  |
| 10:7.489 | המלאכים | . |
| 10:7.490 | וגם |  |
| 10:7.491 | אמרו |  |
| 10:7.492 | בסמוך |  |
| 10:7.493 | אמחה |  |
| 10:7.494 | את |  |
| 10:7.495 | האדם |  |
| 10:7.496 | ולא |  |
| 10:7.497 | ענש |  |
| 10:7.498 | המלאכים |  |
| 10:7.499 | עמהם |  |
| 10:7.500 | ע״כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

