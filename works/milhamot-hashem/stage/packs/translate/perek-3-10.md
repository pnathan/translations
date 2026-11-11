# Translation pack: milhamot-hashem perek-3-10 (Milchamot ha-Shem 3:16–3:16)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-10`.

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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
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
- דרך ארץ: "worldly occupation", as in the project's Avot.
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
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


## Milchamot ha-Shem 3:16

Hebrew (Wikisource): טז) ואשר טענת עלי ואמרת מי נתן לך רשות לדרוש ולחקור אחר דבריהם, תמה אני מאד על דבריך אלה? וכי נתנה תורה לקרות בה כעוף המצפצף או כבהמה גועה ואינו יודע מה מדבר? והלא העושה כן נקרא רשע ערום כדגרסינן בסוטה (דף כב.) עולא אמר זה שקרא ושנה ולא שמש ת"ח, ר' אליעזר אמר ה"ז עם הארץ, ר' שמואל בר נחמני אמר הרי זה בור, ר' ינאי אמר ה"ז כותי, ר' אחא בר יעקב אמר ה"ז אמגוש, אמר רב נחמן בר יצחק מסתברא כרב אחא בר יעקב דאמרי אינשי רטין מגושא ולא ידע מה רטין, תני תנא ולא ידע מה תני ע"כ. ולפי דבריך עסקי הגוף ומכתבי מסחר שהם מעסקי חיי הגופות עדיפי מעסק התורה שהם חיי הנפש, אשר בהם נזכה לחיי העולם הבא ונקביל בהם פני אלהינו מלכנו? שהרי אם יביאו לך שני מכתבים מעסקי מסחר שהם סותרים זה את זה, אתה נותן לב בהם להבין כדי שלא תפסיד ואינך מקבלם באמונת אומן! ובתורה שהיא חיי הנפש תקבל את הסותר אמונתך שאמונת אומן אם טוב ואם רע בלי הבנה? הלא עיקר עסק התורה לדעת ולהבין מצותיה וחקותיה! ודוד המלך ע"ה אמר מתחנן לפני הקב"ה הבינני ואלמדה מצותיך. ואמר עוד טוב טעם ודעת למדני כי במצותיך האמנתי וקל וחומר הוא, ומה בשאר מצות של תורה צריך לב להבין שלא יטעה להתיר את האסור, או אפילו לאסור את המותר וכמאמר רז"ל הוי זהיר בתלמוד ששגגת תלמוד עולה זדון? במצות ידיעת השי"ת ויחודו על פי האמת והצדק שצונו השי"ת בתורה כמו שכתוב וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד. ודוד המלך ע"ה צוה את בנו לאמר דע את אלהי אביך ועבדהו. והנביא ירמיהו מכריז ואמור אל יתהלל חכם בחכמתו, ואל יתהלל הגבור בגבורתו, אל יתהלל עשיר בעשרו, כי אם בזאת יתהלל המתהלל השכל וידוע אותי כי אני ה' עושה חסדו וכו'. על אחת כמה וכמה שאנחנו צריכים לידע את מי נעבוד.

| id | word | punctuation after |
|---|---|---|
| 3:16.1 | טז | ) |
| 3:16.2 | ואשר |  |
| 3:16.3 | טענת |  |
| 3:16.4 | עלי |  |
| 3:16.5 | ואמרת |  |
| 3:16.6 | מי |  |
| 3:16.7 | נתן |  |
| 3:16.8 | לך |  |
| 3:16.9 | רשות |  |
| 3:16.10 | לדרוש |  |
| 3:16.11 | ולחקור |  |
| 3:16.12 | אחר |  |
| 3:16.13 | דבריהם | , |
| 3:16.14 | תמה |  |
| 3:16.15 | אני |  |
| 3:16.16 | מאד |  |
| 3:16.17 | על |  |
| 3:16.18 | דבריך |  |
| 3:16.19 | אלה | ? |
| 3:16.20 | וכי |  |
| 3:16.21 | נתנה |  |
| 3:16.22 | תורה |  |
| 3:16.23 | לקרות |  |
| 3:16.24 | בה |  |
| 3:16.25 | כעוף |  |
| 3:16.26 | המצפצף |  |
| 3:16.27 | או |  |
| 3:16.28 | כבהמה |  |
| 3:16.29 | גועה |  |
| 3:16.30 | ואינו |  |
| 3:16.31 | יודע |  |
| 3:16.32 | מה |  |
| 3:16.33 | מדבר | ? |
| 3:16.34 | והלא |  |
| 3:16.35 | העושה |  |
| 3:16.36 | כן |  |
| 3:16.37 | נקרא |  |
| 3:16.38 | רשע |  |
| 3:16.39 | ערום |  |
| 3:16.40 | כדגרסינן |  |
| 3:16.41 | בסוטה |  |
| 3:16.42 | דף |  |
| 3:16.43 | כב | .) |
| 3:16.44 | עולא |  |
| 3:16.45 | אמר |  |
| 3:16.46 | זה |  |
| 3:16.47 | שקרא |  |
| 3:16.48 | ושנה |  |
| 3:16.49 | ולא |  |
| 3:16.50 | שמש |  |
| 3:16.51 | ת"ח | , |
| 3:16.52 | ר | ' |
| 3:16.53 | אליעזר |  |
| 3:16.54 | אמר |  |
| 3:16.55 | ה"ז |  |
| 3:16.56 | עם |  |
| 3:16.57 | הארץ | , |
| 3:16.58 | ר | ' |
| 3:16.59 | שמואל |  |
| 3:16.60 | בר |  |
| 3:16.61 | נחמני |  |
| 3:16.62 | אמר |  |
| 3:16.63 | הרי |  |
| 3:16.64 | זה |  |
| 3:16.65 | בור | , |
| 3:16.66 | ר | ' |
| 3:16.67 | ינאי |  |
| 3:16.68 | אמר |  |
| 3:16.69 | ה"ז |  |
| 3:16.70 | כותי | , |
| 3:16.71 | ר | ' |
| 3:16.72 | אחא |  |
| 3:16.73 | בר |  |
| 3:16.74 | יעקב |  |
| 3:16.75 | אמר |  |
| 3:16.76 | ה"ז |  |
| 3:16.77 | אמגוש | , |
| 3:16.78 | אמר |  |
| 3:16.79 | רב |  |
| 3:16.80 | נחמן |  |
| 3:16.81 | בר |  |
| 3:16.82 | יצחק |  |
| 3:16.83 | מסתברא |  |
| 3:16.84 | כרב |  |
| 3:16.85 | אחא |  |
| 3:16.86 | בר |  |
| 3:16.87 | יעקב |  |
| 3:16.88 | דאמרי |  |
| 3:16.89 | אינשי |  |
| 3:16.90 | רטין |  |
| 3:16.91 | מגושא |  |
| 3:16.92 | ולא |  |
| 3:16.93 | ידע |  |
| 3:16.94 | מה |  |
| 3:16.95 | רטין | , |
| 3:16.96 | תני |  |
| 3:16.97 | תנא |  |
| 3:16.98 | ולא |  |
| 3:16.99 | ידע |  |
| 3:16.100 | מה |  |
| 3:16.101 | תני |  |
| 3:16.102 | ע"כ | . |
| 3:16.103 | ולפי |  |
| 3:16.104 | דבריך |  |
| 3:16.105 | עסקי |  |
| 3:16.106 | הגוף |  |
| 3:16.107 | ומכתבי |  |
| 3:16.108 | מסחר |  |
| 3:16.109 | שהם |  |
| 3:16.110 | מעסקי |  |
| 3:16.111 | חיי |  |
| 3:16.112 | הגופות |  |
| 3:16.113 | עדיפי |  |
| 3:16.114 | מעסק |  |
| 3:16.115 | התורה |  |
| 3:16.116 | שהם |  |
| 3:16.117 | חיי |  |
| 3:16.118 | הנפש | , |
| 3:16.119 | אשר |  |
| 3:16.120 | בהם |  |
| 3:16.121 | נזכה |  |
| 3:16.122 | לחיי |  |
| 3:16.123 | העולם |  |
| 3:16.124 | הבא |  |
| 3:16.125 | ונקביל |  |
| 3:16.126 | בהם |  |
| 3:16.127 | פני |  |
| 3:16.128 | אלהינו |  |
| 3:16.129 | מלכנו | ? |
| 3:16.130 | שהרי |  |
| 3:16.131 | אם |  |
| 3:16.132 | יביאו |  |
| 3:16.133 | לך |  |
| 3:16.134 | שני |  |
| 3:16.135 | מכתבים |  |
| 3:16.136 | מעסקי |  |
| 3:16.137 | מסחר |  |
| 3:16.138 | שהם |  |
| 3:16.139 | סותרים |  |
| 3:16.140 | זה |  |
| 3:16.141 | את |  |
| 3:16.142 | זה | , |
| 3:16.143 | אתה |  |
| 3:16.144 | נותן |  |
| 3:16.145 | לב |  |
| 3:16.146 | בהם |  |
| 3:16.147 | להבין |  |
| 3:16.148 | כדי |  |
| 3:16.149 | שלא |  |
| 3:16.150 | תפסיד |  |
| 3:16.151 | ואינך |  |
| 3:16.152 | מקבלם |  |
| 3:16.153 | באמונת |  |
| 3:16.154 | אומן | ! |
| 3:16.155 | ובתורה |  |
| 3:16.156 | שהיא |  |
| 3:16.157 | חיי |  |
| 3:16.158 | הנפש |  |
| 3:16.159 | תקבל |  |
| 3:16.160 | את |  |
| 3:16.161 | הסותר |  |
| 3:16.162 | אמונתך |  |
| 3:16.163 | שאמונת |  |
| 3:16.164 | אומן |  |
| 3:16.165 | אם |  |
| 3:16.166 | טוב |  |
| 3:16.167 | ואם |  |
| 3:16.168 | רע |  |
| 3:16.169 | בלי |  |
| 3:16.170 | הבנה | ? |
| 3:16.171 | הלא |  |
| 3:16.172 | עיקר |  |
| 3:16.173 | עסק |  |
| 3:16.174 | התורה |  |
| 3:16.175 | לדעת |  |
| 3:16.176 | ולהבין |  |
| 3:16.177 | מצותיה |  |
| 3:16.178 | וחקותיה | ! |
| 3:16.179 | ודוד |  |
| 3:16.180 | המלך |  |
| 3:16.181 | ע"ה |  |
| 3:16.182 | אמר |  |
| 3:16.183 | מתחנן |  |
| 3:16.184 | לפני |  |
| 3:16.185 | הקב"ה |  |
| 3:16.186 | הבינני |  |
| 3:16.187 | ואלמדה |  |
| 3:16.188 | מצותיך | . |
| 3:16.189 | ואמר |  |
| 3:16.190 | עוד |  |
| 3:16.191 | טוב |  |
| 3:16.192 | טעם |  |
| 3:16.193 | ודעת |  |
| 3:16.194 | למדני |  |
| 3:16.195 | כי |  |
| 3:16.196 | במצותיך |  |
| 3:16.197 | האמנתי |  |
| 3:16.198 | וקל |  |
| 3:16.199 | וחומר |  |
| 3:16.200 | הוא | , |
| 3:16.201 | ומה |  |
| 3:16.202 | בשאר |  |
| 3:16.203 | מצות |  |
| 3:16.204 | של |  |
| 3:16.205 | תורה |  |
| 3:16.206 | צריך |  |
| 3:16.207 | לב |  |
| 3:16.208 | להבין |  |
| 3:16.209 | שלא |  |
| 3:16.210 | יטעה |  |
| 3:16.211 | להתיר |  |
| 3:16.212 | את |  |
| 3:16.213 | האסור | , |
| 3:16.214 | או |  |
| 3:16.215 | אפילו |  |
| 3:16.216 | לאסור |  |
| 3:16.217 | את |  |
| 3:16.218 | המותר |  |
| 3:16.219 | וכמאמר |  |
| 3:16.220 | רז"ל |  |
| 3:16.221 | הוי |  |
| 3:16.222 | זהיר |  |
| 3:16.223 | בתלמוד |  |
| 3:16.224 | ששגגת |  |
| 3:16.225 | תלמוד |  |
| 3:16.226 | עולה |  |
| 3:16.227 | זדון | ? |
| 3:16.228 | במצות |  |
| 3:16.229 | ידיעת |  |
| 3:16.230 | השי"ת |  |
| 3:16.231 | ויחודו |  |
| 3:16.232 | על |  |
| 3:16.233 | פי |  |
| 3:16.234 | האמת |  |
| 3:16.235 | והצדק |  |
| 3:16.236 | שצונו |  |
| 3:16.237 | השי"ת |  |
| 3:16.238 | בתורה |  |
| 3:16.239 | כמו |  |
| 3:16.240 | שכתוב |  |
| 3:16.241 | וידעת |  |
| 3:16.242 | היום |  |
| 3:16.243 | והשבות |  |
| 3:16.244 | אל |  |
| 3:16.245 | לבבך |  |
| 3:16.246 | כי |  |
| 3:16.247 | ה | ' |
| 3:16.248 | הוא |  |
| 3:16.249 | האלהים |  |
| 3:16.250 | בשמים |  |
| 3:16.251 | ממעל |  |
| 3:16.252 | ועל |  |
| 3:16.253 | הארץ |  |
| 3:16.254 | מתחת |  |
| 3:16.255 | אין |  |
| 3:16.256 | עוד | . |
| 3:16.257 | ודוד |  |
| 3:16.258 | המלך |  |
| 3:16.259 | ע"ה |  |
| 3:16.260 | צוה |  |
| 3:16.261 | את |  |
| 3:16.262 | בנו |  |
| 3:16.263 | לאמר |  |
| 3:16.264 | דע |  |
| 3:16.265 | את |  |
| 3:16.266 | אלהי |  |
| 3:16.267 | אביך |  |
| 3:16.268 | ועבדהו | . |
| 3:16.269 | והנביא |  |
| 3:16.270 | ירמיהו |  |
| 3:16.271 | מכריז |  |
| 3:16.272 | ואמור |  |
| 3:16.273 | אל |  |
| 3:16.274 | יתהלל |  |
| 3:16.275 | חכם |  |
| 3:16.276 | בחכמתו | , |
| 3:16.277 | ואל |  |
| 3:16.278 | יתהלל |  |
| 3:16.279 | הגבור |  |
| 3:16.280 | בגבורתו | , |
| 3:16.281 | אל |  |
| 3:16.282 | יתהלל |  |
| 3:16.283 | עשיר |  |
| 3:16.284 | בעשרו | , |
| 3:16.285 | כי |  |
| 3:16.286 | אם |  |
| 3:16.287 | בזאת |  |
| 3:16.288 | יתהלל |  |
| 3:16.289 | המתהלל |  |
| 3:16.290 | השכל |  |
| 3:16.291 | וידוע |  |
| 3:16.292 | אותי |  |
| 3:16.293 | כי |  |
| 3:16.294 | אני |  |
| 3:16.295 | ה | ' |
| 3:16.296 | עושה |  |
| 3:16.297 | חסדו |  |
| 3:16.298 | וכו | '. |
| 3:16.299 | על |  |
| 3:16.300 | אחת |  |
| 3:16.301 | כמה |  |
| 3:16.302 | וכמה |  |
| 3:16.303 | שאנחנו |  |
| 3:16.304 | צריכים |  |
| 3:16.305 | לידע |  |
| 3:16.306 | את |  |
| 3:16.307 | מי |  |
| 3:16.308 | נעבוד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

