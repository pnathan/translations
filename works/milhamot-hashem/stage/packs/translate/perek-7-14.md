# Translation pack: milhamot-hashem perek-7-14 (Milchamot ha-Shem 7:19–7:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-14`.

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
- **קצר אפים**: the author's Hebrew calque of זעיר אנפין, set against ארך
  אפים in Exod 34:6. Write "quick to anger", and render ארך אפים "slow to
  anger" everywhere, so that the pair shows in English. "באל זר קצר אפים
  (ז״א)" (1:4) → "a strange god, quick to anger (Ze'ir Anpin)"; "הנקרא (קצר
  אפים) זעיר אנפין" (3:3) → "called (quick to anger) Ze'ir Anpin". Add a note
  each time: Aramaic "small face"; Prov 14:17. "קצא אפים" (6:23) is translated
  as קצר, with the spelling recorded in `apparatus`.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".

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


## Milchamot ha-Shem 7:19

Hebrew (Wikisource): צא) ומזה תבין ותשכיל גודל השגיאה והטעות שטעו ושגגו רבותינו בעלי הקבלה החדשה הנכריה. שמלבד שמנו באלהות כמה מדות ומדרגות זו למעלה מזו במקום ובזמן. וגם במעלה עד שאמרו שזה נוטל רשות מזה, והאלהים הנקבה נוטלת רשות מאלהים הזכר, לפי דעתם כמבואר בזהר בראשית דף כ"ב. ויחסו כל עבודותינו וכל תפלותינו וברכותינו למדרגה התחתונה מכולם, וקראו לו בשמו של הקב"ה ואמרו שהוא ה' אלהינו ואנחנו עמו וכו' ואותו נעבוד וכינו אותו בשם "זעיר אנפין", שענינו בלשון הכתוב קצר אפים. והנה כבר ידעת ונתבאר לך מכתבי הקדש גנות ופחיתות המדה המגונה הזאת, ואיך יתארו את האלהים אשר יעבדו אותו במדה המגונה הזאת, ועליהם אמר אליהו בן ברכאל הבוזי במליצתו, האמור לממך בליעל רשע אל נדיבים. ופי' רש"י ז"ל וכי הגון וראוי לומר למלך בליעל? ולבאר רשע אל נדיבים? אשר לא נשא פני שרים, כלומר למלך עולם אשר אין משוא פנים לפניו עכ"ל. וכן פירש רבי אברהם פריצול ז"ל האמור למלך וכו' אשר לא נשא פני שרים, האם ראוי לומר למלך הממונה על הנהגה כזאת שיהיה בליעל? ואם ראוי לומר רשע למי שהוא נדיב מנדיבי עמים ונודע גזעו וטובו? ואיך ראוי ליחס עול למלך צדיק וישר כזה שהוא ה' אלהינו אשר לא נשא פני שרים בענין הדין לכבדם ולשאת את פניהם, ולא נכר לפניו השוע והעשיר בהיותו טוען בדין עם העני לחוס על כבודו, בעבור כי מעשה ידיו כולם, ואין הבדל ביניהם מצד יצירתם לפני הקב"ה עכ"ל. ודון מינה ואוקי הכא, כיון שהם אומרים שזעיר אנפין הוא ה' אלהינו איך יקראו לו בשם קצר אפים אשר יעשה אולת? כי לא יאות לתאר את מלכו של עולם אלא בתארי שבח ועילוי, לא תארי אולת וכסילות כתאר קצר אפים (זעיר אנפין) שהוא עושה אולת, כמו שאי אפשר לומר למלכו של עולם בליעל כמאמר אליהו. ומזה תבין ותשכיל כי לא קבלה אמתית היא. רק נכריה היא. ולא מרז"ל בעלי המשנה והתלמוד יצאו הדברים האלה, חלילה לכל בעל אמונת תורת משה רע"ה וקבלת רבותינו ז"ל להאמין בדברים כאלה. כי כבר נתברר אצלינו גדולת חכמינו וטוב שכלם ועומק עיונם בחכמה, והם השיגו את האמת השגה גמורה ואמתית, ואמרו כל המשתף שם שמים ודבר אחר נעקר מן העולם.

| id | word | punctuation after |
|---|---|---|
| 7:19.1 | צא | ) |
| 7:19.2 | ומזה |  |
| 7:19.3 | תבין |  |
| 7:19.4 | ותשכיל |  |
| 7:19.5 | גודל |  |
| 7:19.6 | השגיאה |  |
| 7:19.7 | והטעות |  |
| 7:19.8 | שטעו |  |
| 7:19.9 | ושגגו |  |
| 7:19.10 | רבותינו |  |
| 7:19.11 | בעלי |  |
| 7:19.12 | הקבלה |  |
| 7:19.13 | החדשה |  |
| 7:19.14 | הנכריה | . |
| 7:19.15 | שמלבד |  |
| 7:19.16 | שמנו |  |
| 7:19.17 | באלהות |  |
| 7:19.18 | כמה |  |
| 7:19.19 | מדות |  |
| 7:19.20 | ומדרגות |  |
| 7:19.21 | זו |  |
| 7:19.22 | למעלה |  |
| 7:19.23 | מזו |  |
| 7:19.24 | במקום |  |
| 7:19.25 | ובזמן | . |
| 7:19.26 | וגם |  |
| 7:19.27 | במעלה |  |
| 7:19.28 | עד |  |
| 7:19.29 | שאמרו |  |
| 7:19.30 | שזה |  |
| 7:19.31 | נוטל |  |
| 7:19.32 | רשות |  |
| 7:19.33 | מזה | , |
| 7:19.34 | והאלהים |  |
| 7:19.35 | הנקבה |  |
| 7:19.36 | נוטלת |  |
| 7:19.37 | רשות |  |
| 7:19.38 | מאלהים |  |
| 7:19.39 | הזכר | , |
| 7:19.40 | לפי |  |
| 7:19.41 | דעתם |  |
| 7:19.42 | כמבואר |  |
| 7:19.43 | בזהר |  |
| 7:19.44 | בראשית |  |
| 7:19.45 | דף |  |
| 7:19.46 | כ"ב | . |
| 7:19.47 | ויחסו |  |
| 7:19.48 | כל |  |
| 7:19.49 | עבודותינו |  |
| 7:19.50 | וכל |  |
| 7:19.51 | תפלותינו |  |
| 7:19.52 | וברכותינו |  |
| 7:19.53 | למדרגה |  |
| 7:19.54 | התחתונה |  |
| 7:19.55 | מכולם | , |
| 7:19.56 | וקראו |  |
| 7:19.57 | לו |  |
| 7:19.58 | בשמו |  |
| 7:19.59 | של |  |
| 7:19.60 | הקב"ה |  |
| 7:19.61 | ואמרו |  |
| 7:19.62 | שהוא |  |
| 7:19.63 | ה | ' |
| 7:19.64 | אלהינו |  |
| 7:19.65 | ואנחנו |  |
| 7:19.66 | עמו |  |
| 7:19.67 | וכו | ' |
| 7:19.68 | ואותו |  |
| 7:19.69 | נעבוד |  |
| 7:19.70 | וכינו |  |
| 7:19.71 | אותו |  |
| 7:19.72 | בשם |  |
| 7:19.73 | זעיר |  |
| 7:19.74 | אנפין | ", |
| 7:19.75 | שענינו |  |
| 7:19.76 | בלשון |  |
| 7:19.77 | הכתוב |  |
| 7:19.78 | קצר |  |
| 7:19.79 | אפים | . |
| 7:19.80 | והנה |  |
| 7:19.81 | כבר |  |
| 7:19.82 | ידעת |  |
| 7:19.83 | ונתבאר |  |
| 7:19.84 | לך |  |
| 7:19.85 | מכתבי |  |
| 7:19.86 | הקדש |  |
| 7:19.87 | גנות |  |
| 7:19.88 | ופחיתות |  |
| 7:19.89 | המדה |  |
| 7:19.90 | המגונה |  |
| 7:19.91 | הזאת | , |
| 7:19.92 | ואיך |  |
| 7:19.93 | יתארו |  |
| 7:19.94 | את |  |
| 7:19.95 | האלהים |  |
| 7:19.96 | אשר |  |
| 7:19.97 | יעבדו |  |
| 7:19.98 | אותו |  |
| 7:19.99 | במדה |  |
| 7:19.100 | המגונה |  |
| 7:19.101 | הזאת | , |
| 7:19.102 | ועליהם |  |
| 7:19.103 | אמר |  |
| 7:19.104 | אליהו |  |
| 7:19.105 | בן |  |
| 7:19.106 | ברכאל |  |
| 7:19.107 | הבוזי |  |
| 7:19.108 | במליצתו | , |
| 7:19.109 | האמור |  |
| 7:19.110 | לממך |  |
| 7:19.111 | בליעל |  |
| 7:19.112 | רשע |  |
| 7:19.113 | אל |  |
| 7:19.114 | נדיבים | . |
| 7:19.115 | ופי | ' |
| 7:19.116 | רש"י |  |
| 7:19.117 | ז"ל |  |
| 7:19.118 | וכי |  |
| 7:19.119 | הגון |  |
| 7:19.120 | וראוי |  |
| 7:19.121 | לומר |  |
| 7:19.122 | למלך |  |
| 7:19.123 | בליעל | ? |
| 7:19.124 | ולבאר |  |
| 7:19.125 | רשע |  |
| 7:19.126 | אל |  |
| 7:19.127 | נדיבים | ? |
| 7:19.128 | אשר |  |
| 7:19.129 | לא |  |
| 7:19.130 | נשא |  |
| 7:19.131 | פני |  |
| 7:19.132 | שרים | , |
| 7:19.133 | כלומר |  |
| 7:19.134 | למלך |  |
| 7:19.135 | עולם |  |
| 7:19.136 | אשר |  |
| 7:19.137 | אין |  |
| 7:19.138 | משוא |  |
| 7:19.139 | פנים |  |
| 7:19.140 | לפניו |  |
| 7:19.141 | עכ"ל | . |
| 7:19.142 | וכן |  |
| 7:19.143 | פירש |  |
| 7:19.144 | רבי |  |
| 7:19.145 | אברהם |  |
| 7:19.146 | פריצול |  |
| 7:19.147 | ז"ל |  |
| 7:19.148 | האמור |  |
| 7:19.149 | למלך |  |
| 7:19.150 | וכו | ' |
| 7:19.151 | אשר |  |
| 7:19.152 | לא |  |
| 7:19.153 | נשא |  |
| 7:19.154 | פני |  |
| 7:19.155 | שרים | , |
| 7:19.156 | האם |  |
| 7:19.157 | ראוי |  |
| 7:19.158 | לומר |  |
| 7:19.159 | למלך |  |
| 7:19.160 | הממונה |  |
| 7:19.161 | על |  |
| 7:19.162 | הנהגה |  |
| 7:19.163 | כזאת |  |
| 7:19.164 | שיהיה |  |
| 7:19.165 | בליעל | ? |
| 7:19.166 | ואם |  |
| 7:19.167 | ראוי |  |
| 7:19.168 | לומר |  |
| 7:19.169 | רשע |  |
| 7:19.170 | למי |  |
| 7:19.171 | שהוא |  |
| 7:19.172 | נדיב |  |
| 7:19.173 | מנדיבי |  |
| 7:19.174 | עמים |  |
| 7:19.175 | ונודע |  |
| 7:19.176 | גזעו |  |
| 7:19.177 | וטובו | ? |
| 7:19.178 | ואיך |  |
| 7:19.179 | ראוי |  |
| 7:19.180 | ליחס |  |
| 7:19.181 | עול |  |
| 7:19.182 | למלך |  |
| 7:19.183 | צדיק |  |
| 7:19.184 | וישר |  |
| 7:19.185 | כזה |  |
| 7:19.186 | שהוא |  |
| 7:19.187 | ה | ' |
| 7:19.188 | אלהינו |  |
| 7:19.189 | אשר |  |
| 7:19.190 | לא |  |
| 7:19.191 | נשא |  |
| 7:19.192 | פני |  |
| 7:19.193 | שרים |  |
| 7:19.194 | בענין |  |
| 7:19.195 | הדין |  |
| 7:19.196 | לכבדם |  |
| 7:19.197 | ולשאת |  |
| 7:19.198 | את |  |
| 7:19.199 | פניהם | , |
| 7:19.200 | ולא |  |
| 7:19.201 | נכר |  |
| 7:19.202 | לפניו |  |
| 7:19.203 | השוע |  |
| 7:19.204 | והעשיר |  |
| 7:19.205 | בהיותו |  |
| 7:19.206 | טוען |  |
| 7:19.207 | בדין |  |
| 7:19.208 | עם |  |
| 7:19.209 | העני |  |
| 7:19.210 | לחוס |  |
| 7:19.211 | על |  |
| 7:19.212 | כבודו | , |
| 7:19.213 | בעבור |  |
| 7:19.214 | כי |  |
| 7:19.215 | מעשה |  |
| 7:19.216 | ידיו |  |
| 7:19.217 | כולם | , |
| 7:19.218 | ואין |  |
| 7:19.219 | הבדל |  |
| 7:19.220 | ביניהם |  |
| 7:19.221 | מצד |  |
| 7:19.222 | יצירתם |  |
| 7:19.223 | לפני |  |
| 7:19.224 | הקב"ה |  |
| 7:19.225 | עכ"ל | . |
| 7:19.226 | ודון |  |
| 7:19.227 | מינה |  |
| 7:19.228 | ואוקי |  |
| 7:19.229 | הכא | , |
| 7:19.230 | כיון |  |
| 7:19.231 | שהם |  |
| 7:19.232 | אומרים |  |
| 7:19.233 | שזעיר |  |
| 7:19.234 | אנפין |  |
| 7:19.235 | הוא |  |
| 7:19.236 | ה | ' |
| 7:19.237 | אלהינו |  |
| 7:19.238 | איך |  |
| 7:19.239 | יקראו |  |
| 7:19.240 | לו |  |
| 7:19.241 | בשם |  |
| 7:19.242 | קצר |  |
| 7:19.243 | אפים |  |
| 7:19.244 | אשר |  |
| 7:19.245 | יעשה |  |
| 7:19.246 | אולת | ? |
| 7:19.247 | כי |  |
| 7:19.248 | לא |  |
| 7:19.249 | יאות |  |
| 7:19.250 | לתאר |  |
| 7:19.251 | את |  |
| 7:19.252 | מלכו |  |
| 7:19.253 | של |  |
| 7:19.254 | עולם |  |
| 7:19.255 | אלא |  |
| 7:19.256 | בתארי |  |
| 7:19.257 | שבח |  |
| 7:19.258 | ועילוי | , |
| 7:19.259 | לא |  |
| 7:19.260 | תארי |  |
| 7:19.261 | אולת |  |
| 7:19.262 | וכסילות |  |
| 7:19.263 | כתאר |  |
| 7:19.264 | קצר |  |
| 7:19.265 | אפים |  |
| 7:19.266 | זעיר |  |
| 7:19.267 | אנפין | ) |
| 7:19.268 | שהוא |  |
| 7:19.269 | עושה |  |
| 7:19.270 | אולת | , |
| 7:19.271 | כמו |  |
| 7:19.272 | שאי |  |
| 7:19.273 | אפשר |  |
| 7:19.274 | לומר |  |
| 7:19.275 | למלכו |  |
| 7:19.276 | של |  |
| 7:19.277 | עולם |  |
| 7:19.278 | בליעל |  |
| 7:19.279 | כמאמר |  |
| 7:19.280 | אליהו | . |
| 7:19.281 | ומזה |  |
| 7:19.282 | תבין |  |
| 7:19.283 | ותשכיל |  |
| 7:19.284 | כי |  |
| 7:19.285 | לא |  |
| 7:19.286 | קבלה |  |
| 7:19.287 | אמתית |  |
| 7:19.288 | היא | . |
| 7:19.289 | רק |  |
| 7:19.290 | נכריה |  |
| 7:19.291 | היא | . |
| 7:19.292 | ולא |  |
| 7:19.293 | מרז"ל |  |
| 7:19.294 | בעלי |  |
| 7:19.295 | המשנה |  |
| 7:19.296 | והתלמוד |  |
| 7:19.297 | יצאו |  |
| 7:19.298 | הדברים |  |
| 7:19.299 | האלה | , |
| 7:19.300 | חלילה |  |
| 7:19.301 | לכל |  |
| 7:19.302 | בעל |  |
| 7:19.303 | אמונת |  |
| 7:19.304 | תורת |  |
| 7:19.305 | משה |  |
| 7:19.306 | רע"ה |  |
| 7:19.307 | וקבלת |  |
| 7:19.308 | רבותינו |  |
| 7:19.309 | ז"ל |  |
| 7:19.310 | להאמין |  |
| 7:19.311 | בדברים |  |
| 7:19.312 | כאלה | . |
| 7:19.313 | כי |  |
| 7:19.314 | כבר |  |
| 7:19.315 | נתברר |  |
| 7:19.316 | אצלינו |  |
| 7:19.317 | גדולת |  |
| 7:19.318 | חכמינו |  |
| 7:19.319 | וטוב |  |
| 7:19.320 | שכלם |  |
| 7:19.321 | ועומק |  |
| 7:19.322 | עיונם |  |
| 7:19.323 | בחכמה | , |
| 7:19.324 | והם |  |
| 7:19.325 | השיגו |  |
| 7:19.326 | את |  |
| 7:19.327 | האמת |  |
| 7:19.328 | השגה |  |
| 7:19.329 | גמורה |  |
| 7:19.330 | ואמתית | , |
| 7:19.331 | ואמרו |  |
| 7:19.332 | כל |  |
| 7:19.333 | המשתף |  |
| 7:19.334 | שם |  |
| 7:19.335 | שמים |  |
| 7:19.336 | ודבר |  |
| 7:19.337 | אחר |  |
| 7:19.338 | נעקר |  |
| 7:19.339 | מן |  |
| 7:19.340 | העולם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

