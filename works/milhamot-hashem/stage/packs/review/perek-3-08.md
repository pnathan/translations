# Review pack: milhamot-hashem perek-3-08 (Milchamot ha-Shem 3:13–3:13)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-08`.

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
- An infinitive after an adjective of motion (הגיע חתור): follow the
  grammar of the Hebrew as written; the reviewer's reading may be recorded in
  `readings`.
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
- ס' before a quoted verse is "verse" (סימן = "section" in the author's
  numbered sections); read it from the context and note the choice.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- והוא שלא: "provided that" (fixed).
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
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


## Milchamot ha-Shem 3:13

Hebrew (Wikisource): יג) למדנו מדברי הגאון כי הרבה ספרים יש שמיחסין אותם על שם גדולי חכמי ישראל (בשקר) בשביל איזה תועלת המגיע להם, ולמדנו ג"כ שכל שאינו מפורש במשנה ובגמרא וראינו שהוא היפך משמעות התלמוד לאו קבלה הוא, ולא סמכינן עליה. ומטעם זה אין פוסקים מן המדרשות במקום שהם חולקים על התלמוד, ואפי' תלמוד ירושלמי אין חוששין לו במקום שהוא חולק על הבבלי אשר כל בית ישראל נשענים עליו בין להקל בין להחמיר כמבואר כל זה בדברי רבותינו הפוסקים בכמה מקומות ועיין לקמן (סימן מ"ב) מה שכתבתי שם בשם הרא"ש, וצא ולמד ממה שכתב ר"ת ן' יחיא הובא בס' רביד הזהב בהל' פסח וז"ל ובזמנינו זה אפס המדע וסתרי תורה נעלמו מעיני כל חי, לא ידע אנוש סדר החכמה הזאת, וכל המחזיק זה יקרה לו מההזק והסכנה חלק רב ( ר"ל הזק וסכנת מינות שאין לו חלק באלהי ישראל). ועל כיוצא בזה הזהירנו רז"ל במופלא ממך אל תדרוש ובמכוסה ממך אל תחקור. ודי לנו במה שהורונו חכמי הקבלה האמתית והמוסכמת מכל צד בפי' התורה ומצותיה. ע"פ המסורה איש מפי איש עד ימות משה והיא הנקראת תורה שבע"פ אשר בגללה עלה משה למרום ועמד שם ארבעים יום ללמוד כללותיה ופרטיה במדות הנמסרות לו מפי הגבורה. ובדברים אלו אנו מחוייבים להתעסק בהם כי הם חיינו ואורך ימינו וא' מהם לא נשגה. אחת דתו להמית אשר ימרה את פיהם וכל העובר על דברי חכמים חייב מיתה. אך יתר הרמזים אשר בהם לא כלו ימיהם חכמי הקבלה עם היותם מסתרי תורה ומעניני המושכלות לא נצטוינו עליהם. כי אלו נשארו להתעסק בהם יחידי סגולה בזמנים אשר היו לפנינו וכו'. ולא כל הרוצה ליטול את השם יטול. כי דברים שהם כבשונו שלעולם הם. ובזמנינו זה תמה זכות אבות, ואין מחזיק בצדק זאת החכמה: אדרבה שהם הורסים פינות התורה ועמודיה יתפלצון מקול מחצצים העברים העורים לא ידעו דרך ישכן אור, ימששו כעורים, וילאו למצוא הפתח. ובמקום שחשבו להרויח פתע ישברו ואין מרפא. ורבים מעמי הארץ מתיהרים באמרם סוד ה' לנו, מורשה לנו המים העליונים, לא לכם בעלי התלמוד ההולכים בחשך. ולבם בל עמם, טח מראות עיניהם, והדבר אין בהם, חכמים להרע לנפשם. כי החזיקו במה שקצרה יד שכלם. ורוצים לעלות במעלות אין ספק בידם להגיע אל התכלית. ויעלו חרס בידם להתגרד בו. ויהיו להם למוקש וכו' כי תצא מהם תקלה לפרוץ גדר התורה ולהרוס חומותיה. יתפאר במה שלא יבין. ויטעה בהתחלות התורה ושרשיה ויבוא לידי מינות. ויתרחק ממחוז חפצו והיא קרבת אלהים, נוח לו שלא נברא, כל שכן לשנות מדרש ולבטל הלכה עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:13.1 | יג | ) |
| 3:13.2 | למדנו |  |
| 3:13.3 | מדברי |  |
| 3:13.4 | הגאון |  |
| 3:13.5 | כי |  |
| 3:13.6 | הרבה |  |
| 3:13.7 | ספרים |  |
| 3:13.8 | יש |  |
| 3:13.9 | שמיחסין |  |
| 3:13.10 | אותם |  |
| 3:13.11 | על |  |
| 3:13.12 | שם |  |
| 3:13.13 | גדולי |  |
| 3:13.14 | חכמי |  |
| 3:13.15 | ישראל |  |
| 3:13.16 | בשקר | ) |
| 3:13.17 | בשביל |  |
| 3:13.18 | איזה |  |
| 3:13.19 | תועלת |  |
| 3:13.20 | המגיע |  |
| 3:13.21 | להם | , |
| 3:13.22 | ולמדנו |  |
| 3:13.23 | ג"כ |  |
| 3:13.24 | שכל |  |
| 3:13.25 | שאינו |  |
| 3:13.26 | מפורש |  |
| 3:13.27 | במשנה |  |
| 3:13.28 | ובגמרא |  |
| 3:13.29 | וראינו |  |
| 3:13.30 | שהוא |  |
| 3:13.31 | היפך |  |
| 3:13.32 | משמעות |  |
| 3:13.33 | התלמוד |  |
| 3:13.34 | לאו |  |
| 3:13.35 | קבלה |  |
| 3:13.36 | הוא | , |
| 3:13.37 | ולא |  |
| 3:13.38 | סמכינן |  |
| 3:13.39 | עליה | . |
| 3:13.40 | ומטעם |  |
| 3:13.41 | זה |  |
| 3:13.42 | אין |  |
| 3:13.43 | פוסקים |  |
| 3:13.44 | מן |  |
| 3:13.45 | המדרשות |  |
| 3:13.46 | במקום |  |
| 3:13.47 | שהם |  |
| 3:13.48 | חולקים |  |
| 3:13.49 | על |  |
| 3:13.50 | התלמוד | , |
| 3:13.51 | ואפי | ' |
| 3:13.52 | תלמוד |  |
| 3:13.53 | ירושלמי |  |
| 3:13.54 | אין |  |
| 3:13.55 | חוששין |  |
| 3:13.56 | לו |  |
| 3:13.57 | במקום |  |
| 3:13.58 | שהוא |  |
| 3:13.59 | חולק |  |
| 3:13.60 | על |  |
| 3:13.61 | הבבלי |  |
| 3:13.62 | אשר |  |
| 3:13.63 | כל |  |
| 3:13.64 | בית |  |
| 3:13.65 | ישראל |  |
| 3:13.66 | נשענים |  |
| 3:13.67 | עליו |  |
| 3:13.68 | בין |  |
| 3:13.69 | להקל |  |
| 3:13.70 | בין |  |
| 3:13.71 | להחמיר |  |
| 3:13.72 | כמבואר |  |
| 3:13.73 | כל |  |
| 3:13.74 | זה |  |
| 3:13.75 | בדברי |  |
| 3:13.76 | רבותינו |  |
| 3:13.77 | הפוסקים |  |
| 3:13.78 | בכמה |  |
| 3:13.79 | מקומות |  |
| 3:13.80 | ועיין |  |
| 3:13.81 | לקמן |  |
| 3:13.82 | סימן |  |
| 3:13.83 | מ"ב | ) |
| 3:13.84 | מה |  |
| 3:13.85 | שכתבתי |  |
| 3:13.86 | שם |  |
| 3:13.87 | בשם |  |
| 3:13.88 | הרא"ש | , |
| 3:13.89 | וצא |  |
| 3:13.90 | ולמד |  |
| 3:13.91 | ממה |  |
| 3:13.92 | שכתב |  |
| 3:13.93 | ר"ת |  |
| 3:13.94 | ן | ' |
| 3:13.95 | יחיא |  |
| 3:13.96 | הובא |  |
| 3:13.97 | בס | ' |
| 3:13.98 | רביד |  |
| 3:13.99 | הזהב |  |
| 3:13.100 | בהל | ' |
| 3:13.101 | פסח |  |
| 3:13.102 | וז"ל |  |
| 3:13.103 | ובזמנינו |  |
| 3:13.104 | זה |  |
| 3:13.105 | אפס |  |
| 3:13.106 | המדע |  |
| 3:13.107 | וסתרי |  |
| 3:13.108 | תורה |  |
| 3:13.109 | נעלמו |  |
| 3:13.110 | מעיני |  |
| 3:13.111 | כל |  |
| 3:13.112 | חי | , |
| 3:13.113 | לא |  |
| 3:13.114 | ידע |  |
| 3:13.115 | אנוש |  |
| 3:13.116 | סדר |  |
| 3:13.117 | החכמה |  |
| 3:13.118 | הזאת | , |
| 3:13.119 | וכל |  |
| 3:13.120 | המחזיק |  |
| 3:13.121 | זה |  |
| 3:13.122 | יקרה |  |
| 3:13.123 | לו |  |
| 3:13.124 | מההזק |  |
| 3:13.125 | והסכנה |  |
| 3:13.126 | חלק |  |
| 3:13.127 | רב | ( |
| 3:13.128 | ר"ל |  |
| 3:13.129 | הזק |  |
| 3:13.130 | וסכנת |  |
| 3:13.131 | מינות |  |
| 3:13.132 | שאין |  |
| 3:13.133 | לו |  |
| 3:13.134 | חלק |  |
| 3:13.135 | באלהי |  |
| 3:13.136 | ישראל | ). |
| 3:13.137 | ועל |  |
| 3:13.138 | כיוצא |  |
| 3:13.139 | בזה |  |
| 3:13.140 | הזהירנו |  |
| 3:13.141 | רז"ל |  |
| 3:13.142 | במופלא |  |
| 3:13.143 | ממך |  |
| 3:13.144 | אל |  |
| 3:13.145 | תדרוש |  |
| 3:13.146 | ובמכוסה |  |
| 3:13.147 | ממך |  |
| 3:13.148 | אל |  |
| 3:13.149 | תחקור | . |
| 3:13.150 | ודי |  |
| 3:13.151 | לנו |  |
| 3:13.152 | במה |  |
| 3:13.153 | שהורונו |  |
| 3:13.154 | חכמי |  |
| 3:13.155 | הקבלה |  |
| 3:13.156 | האמתית |  |
| 3:13.157 | והמוסכמת |  |
| 3:13.158 | מכל |  |
| 3:13.159 | צד |  |
| 3:13.160 | בפי | ' |
| 3:13.161 | התורה |  |
| 3:13.162 | ומצותיה | . |
| 3:13.163 | ע"פ |  |
| 3:13.164 | המסורה |  |
| 3:13.165 | איש |  |
| 3:13.166 | מפי |  |
| 3:13.167 | איש |  |
| 3:13.168 | עד |  |
| 3:13.169 | ימות |  |
| 3:13.170 | משה |  |
| 3:13.171 | והיא |  |
| 3:13.172 | הנקראת |  |
| 3:13.173 | תורה |  |
| 3:13.174 | שבע"פ |  |
| 3:13.175 | אשר |  |
| 3:13.176 | בגללה |  |
| 3:13.177 | עלה |  |
| 3:13.178 | משה |  |
| 3:13.179 | למרום |  |
| 3:13.180 | ועמד |  |
| 3:13.181 | שם |  |
| 3:13.182 | ארבעים |  |
| 3:13.183 | יום |  |
| 3:13.184 | ללמוד |  |
| 3:13.185 | כללותיה |  |
| 3:13.186 | ופרטיה |  |
| 3:13.187 | במדות |  |
| 3:13.188 | הנמסרות |  |
| 3:13.189 | לו |  |
| 3:13.190 | מפי |  |
| 3:13.191 | הגבורה | . |
| 3:13.192 | ובדברים |  |
| 3:13.193 | אלו |  |
| 3:13.194 | אנו |  |
| 3:13.195 | מחוייבים |  |
| 3:13.196 | להתעסק |  |
| 3:13.197 | בהם |  |
| 3:13.198 | כי |  |
| 3:13.199 | הם |  |
| 3:13.200 | חיינו |  |
| 3:13.201 | ואורך |  |
| 3:13.202 | ימינו |  |
| 3:13.203 | וא | ' |
| 3:13.204 | מהם |  |
| 3:13.205 | לא |  |
| 3:13.206 | נשגה | . |
| 3:13.207 | אחת |  |
| 3:13.208 | דתו |  |
| 3:13.209 | להמית |  |
| 3:13.210 | אשר |  |
| 3:13.211 | ימרה |  |
| 3:13.212 | את |  |
| 3:13.213 | פיהם |  |
| 3:13.214 | וכל |  |
| 3:13.215 | העובר |  |
| 3:13.216 | על |  |
| 3:13.217 | דברי |  |
| 3:13.218 | חכמים |  |
| 3:13.219 | חייב |  |
| 3:13.220 | מיתה | . |
| 3:13.221 | אך |  |
| 3:13.222 | יתר |  |
| 3:13.223 | הרמזים |  |
| 3:13.224 | אשר |  |
| 3:13.225 | בהם |  |
| 3:13.226 | לא |  |
| 3:13.227 | כלו |  |
| 3:13.228 | ימיהם |  |
| 3:13.229 | חכמי |  |
| 3:13.230 | הקבלה |  |
| 3:13.231 | עם |  |
| 3:13.232 | היותם |  |
| 3:13.233 | מסתרי |  |
| 3:13.234 | תורה |  |
| 3:13.235 | ומעניני |  |
| 3:13.236 | המושכלות |  |
| 3:13.237 | לא |  |
| 3:13.238 | נצטוינו |  |
| 3:13.239 | עליהם | . |
| 3:13.240 | כי |  |
| 3:13.241 | אלו |  |
| 3:13.242 | נשארו |  |
| 3:13.243 | להתעסק |  |
| 3:13.244 | בהם |  |
| 3:13.245 | יחידי |  |
| 3:13.246 | סגולה |  |
| 3:13.247 | בזמנים |  |
| 3:13.248 | אשר |  |
| 3:13.249 | היו |  |
| 3:13.250 | לפנינו |  |
| 3:13.251 | וכו | '. |
| 3:13.252 | ולא |  |
| 3:13.253 | כל |  |
| 3:13.254 | הרוצה |  |
| 3:13.255 | ליטול |  |
| 3:13.256 | את |  |
| 3:13.257 | השם |  |
| 3:13.258 | יטול | . |
| 3:13.259 | כי |  |
| 3:13.260 | דברים |  |
| 3:13.261 | שהם |  |
| 3:13.262 | כבשונו |  |
| 3:13.263 | שלעולם |  |
| 3:13.264 | הם | . |
| 3:13.265 | ובזמנינו |  |
| 3:13.266 | זה |  |
| 3:13.267 | תמה |  |
| 3:13.268 | זכות |  |
| 3:13.269 | אבות | , |
| 3:13.270 | ואין |  |
| 3:13.271 | מחזיק |  |
| 3:13.272 | בצדק |  |
| 3:13.273 | זאת |  |
| 3:13.274 | החכמה | : |
| 3:13.275 | אדרבה |  |
| 3:13.276 | שהם |  |
| 3:13.277 | הורסים |  |
| 3:13.278 | פינות |  |
| 3:13.279 | התורה |  |
| 3:13.280 | ועמודיה |  |
| 3:13.281 | יתפלצון |  |
| 3:13.282 | מקול |  |
| 3:13.283 | מחצצים |  |
| 3:13.284 | העברים |  |
| 3:13.285 | העורים |  |
| 3:13.286 | לא |  |
| 3:13.287 | ידעו |  |
| 3:13.288 | דרך |  |
| 3:13.289 | ישכן |  |
| 3:13.290 | אור | , |
| 3:13.291 | ימששו |  |
| 3:13.292 | כעורים | , |
| 3:13.293 | וילאו |  |
| 3:13.294 | למצוא |  |
| 3:13.295 | הפתח | . |
| 3:13.296 | ובמקום |  |
| 3:13.297 | שחשבו |  |
| 3:13.298 | להרויח |  |
| 3:13.299 | פתע |  |
| 3:13.300 | ישברו |  |
| 3:13.301 | ואין |  |
| 3:13.302 | מרפא | . |
| 3:13.303 | ורבים |  |
| 3:13.304 | מעמי |  |
| 3:13.305 | הארץ |  |
| 3:13.306 | מתיהרים |  |
| 3:13.307 | באמרם |  |
| 3:13.308 | סוד |  |
| 3:13.309 | ה | ' |
| 3:13.310 | לנו | , |
| 3:13.311 | מורשה |  |
| 3:13.312 | לנו |  |
| 3:13.313 | המים |  |
| 3:13.314 | העליונים | , |
| 3:13.315 | לא |  |
| 3:13.316 | לכם |  |
| 3:13.317 | בעלי |  |
| 3:13.318 | התלמוד |  |
| 3:13.319 | ההולכים |  |
| 3:13.320 | בחשך | . |
| 3:13.321 | ולבם |  |
| 3:13.322 | בל |  |
| 3:13.323 | עמם | , |
| 3:13.324 | טח |  |
| 3:13.325 | מראות |  |
| 3:13.326 | עיניהם | , |
| 3:13.327 | והדבר |  |
| 3:13.328 | אין |  |
| 3:13.329 | בהם | , |
| 3:13.330 | חכמים |  |
| 3:13.331 | להרע |  |
| 3:13.332 | לנפשם | . |
| 3:13.333 | כי |  |
| 3:13.334 | החזיקו |  |
| 3:13.335 | במה |  |
| 3:13.336 | שקצרה |  |
| 3:13.337 | יד |  |
| 3:13.338 | שכלם | . |
| 3:13.339 | ורוצים |  |
| 3:13.340 | לעלות |  |
| 3:13.341 | במעלות |  |
| 3:13.342 | אין |  |
| 3:13.343 | ספק |  |
| 3:13.344 | בידם |  |
| 3:13.345 | להגיע |  |
| 3:13.346 | אל |  |
| 3:13.347 | התכלית | . |
| 3:13.348 | ויעלו |  |
| 3:13.349 | חרס |  |
| 3:13.350 | בידם |  |
| 3:13.351 | להתגרד |  |
| 3:13.352 | בו | . |
| 3:13.353 | ויהיו |  |
| 3:13.354 | להם |  |
| 3:13.355 | למוקש |  |
| 3:13.356 | וכו | ' |
| 3:13.357 | כי |  |
| 3:13.358 | תצא |  |
| 3:13.359 | מהם |  |
| 3:13.360 | תקלה |  |
| 3:13.361 | לפרוץ |  |
| 3:13.362 | גדר |  |
| 3:13.363 | התורה |  |
| 3:13.364 | ולהרוס |  |
| 3:13.365 | חומותיה | . |
| 3:13.366 | יתפאר |  |
| 3:13.367 | במה |  |
| 3:13.368 | שלא |  |
| 3:13.369 | יבין | . |
| 3:13.370 | ויטעה |  |
| 3:13.371 | בהתחלות |  |
| 3:13.372 | התורה |  |
| 3:13.373 | ושרשיה |  |
| 3:13.374 | ויבוא |  |
| 3:13.375 | לידי |  |
| 3:13.376 | מינות | . |
| 3:13.377 | ויתרחק |  |
| 3:13.378 | ממחוז |  |
| 3:13.379 | חפצו |  |
| 3:13.380 | והיא |  |
| 3:13.381 | קרבת |  |
| 3:13.382 | אלהים | , |
| 3:13.383 | נוח |  |
| 3:13.384 | לו |  |
| 3:13.385 | שלא |  |
| 3:13.386 | נברא | , |
| 3:13.387 | כל |  |
| 3:13.388 | שכן |  |
| 3:13.389 | לשנות |  |
| 3:13.390 | מדרש |  |
| 3:13.391 | ולבטל |  |
| 3:13.392 | הלכה |  |
| 3:13.393 | עכ"ל | . |


### Translation under review (3:13)

English: 13) We learned from the words of the Gaon that there are many books that are ascribed to the names of the great sages of Israel (falsely), for the sake of some benefit that accrues to them, and we learned also that whatever is not stated explicitly in the Mishnah and in the Gemara, and that we see is the opposite of the plain sense of the Talmud, is not a tradition, and we do not rely on it. And for this reason no decisors rule from the midrashim where they disagree with the Talmud, and even the Jerusalem Talmud we do not heed where it disagrees with the Babylonian, on which all the house of Israel rely, both for leniency and for stringency, as all this is made clear in the words of our rabbis the decisors in several places. And see below, §42, what I wrote there in the name of Rabbeinu Asher, And go and learn from what Rabbi Tam ben Yichya wrote, cited in the book Rivid ha-Zahav in the laws of Passover, and these are his words: "In our time knowledge has ceased, and the hidden matters of the Torah have been hidden from the eyes of every living thing; no man knows the order of this wisdom; and whoever holds to this, there befalls him a great share of the harm and the danger (that is to say, the harm and the danger of heresy, in that there is no share for him in the God of Israel). And concerning such things our Rabbis of blessed memory warned us, 'Do not seek what is too wondrous for you, and do not inquire into what is hidden from you.' And it is enough for us what the sages of the true tradition, agreed upon from every side, have taught us in the explanation of the Torah and its commandments. According to the tradition, from man to man, until the days of Moses; and it is called the Oral Torah, on account of which Moses went up on high and stood there forty days to learn its general principles and its particulars by the rules transmitted to him from the mouth of the Power. And with these matters we are bound to occupy ourselves, for they are our life and the length of our days, and of them not one shall we err. His one law is to put to death whoever defies their word, and whoever transgresses the words of the Sages is liable to death. But the remaining allusions in them, which the sages of Kabbalah did not exhaust in their lifetimes, although they are hidden matters of the Torah and matters of the intellect; we were not commanded concerning them. For these remained for the select few to be occupied with them, in the times that were before us, etc. And not everyone who wishes to take up the Name takes it. For these are matters that are like its furnace; they are so forever. And in our time the merit of the fathers has ended, and no one holds this wisdom justly: on the contrary, they destroy the corners of the Torah and its pillars; they shudder at the sound of the trumpeters; the blind, who pass over, do not know the way where light dwells; they grope like the blind, and weary themselves seeking the door. And where they thought to gain, suddenly they shall be broken, and there is no healing. And many of the unlearned boast, saying: 'The secret of the LORD is ours; the upper waters are our heritage, not yours, you masters of the Talmud who walk in darkness.' Their heart is not with their people; their eyes are plastered over, and there is no substance in them; they are wise to do harm to themselves. For they hold fast to what the hand of their intellect is too short to reach. And they want to climb the rungs; there is no doubt that it is in their power to reach the goal. And they bring up a potsherd in their hands to scrape with it. And they become a snare to them, etc. For there comes out from them a stumbling-block, to breach the fence of the Torah and to tear down its walls. One boasts of what he does not understand, and errs in the beginnings of the Torah and its roots, and comes to heresy. And he is distanced from the place of his desire, which is the nearness of God; it would have been better for him had he not been created, all the more so to change a midrash and to annul a halakhah." Thus far his words.

Literal: 13) list marker learned-we from-words-of the-Gaon that many books there-are that-ascribe them to the-name-of the-great sages-of Israel falsely for-the-sake-of some benefit that-accrues to-them and-learned-we also that-whatever is-not stated-explicitly in-the-Mishnah and-in-the-Gemara and-we-see that-it the-opposite-of the-plain-sense-of the-Talmud is-not tradition it and-not we-rely on-it and-for this-reason no decisors from the-midrashim where they disagree with the-Talmud and-even the-Talmud Jerusalem not-we heed it where it disagrees with the-Babylonian which all the-house-of Israel rely on-it both for-leniency and for-stringency as-is-made-clear all this in-the-words-of our-rabbis the-decisors in-several places and-see below section 42 what I-wrote there in-the-name-of Rabbeinu Asher and-go-and learn from-what wrote Rabbi-Tam ben Yichya cited in-the-book Rivid ha-Zahav in-the-laws-of Passover and-these are-his-words in-our-time this has-ceased knowledge and-hidden-matters-of the-Torah have-been-hidden from-the-eyes-of every living-thing no knows man the-order-of the-wisdom this and-whoever holds-to this befalls him of-the-harm and-the-danger a-share great that-is-to-say the-harm and-the-danger-of heresy in-that-there-is-no for-him share in-the-God-of Israel and-concerning such things warned-us our-Rabbis-of-blessed-memory what-is-too-wondrous for-you do-not seek and-what-is-hidden from-you do-not inquire and-enough for-us what that-taught-us the-sages-of tradition true agreed-upon from-every side in-the-explanation-of the-Torah and-its-commandments according-to the-tradition man from-the-mouth-of man until the-days-of Moses and-it is-called the Torah of-the-mouth which on-account-of-it Moses went-up on-high and-stood there forty days to-learn its-general-principles and-its-particulars by-the-measures transmitted to-him from-the-mouth-of the-Power and-with these we are-bound to-occupy-ourselves with-them for they our-life and-the-length-of our-days and-one of-them not we-shall-err his-one law to-put-to-death who defies the-mouth-of-them and-whoever transgresses against the-words-of the-Sages is-liable death but the-remaining allusions which in-them not exhausted their-days the-sages-of Kabbalah although they-are hidden-matters-of the-Torah and-matters-of the-intellect we were-not commanded concerning-them for these remained to-be-occupied with-them the-select few in-the-times that were before-us etc and-not everyone who-wishes to-take the-Name takes for matters that like-its-furnace forever they are and-in-our-time this has-ended the-merit of-the-fathers and-no-one holds justly this wisdom on-the-contrary they destroy the-corners-of the-Torah and-its-pillars they-shudder at-the-sound-of the-trumpeters the-passers-over the-blind not know the-way where dwells light they-grope like-the-blind and-they-weary seeking the-door and-where they-thought to-gain suddenly they-shall-be-broken and-there-is no healing and-many of-the-unlearned boast saying the-secret-of the-LORD for-us a-heritage for-us the-waters the-upper not for-you the-masters-of the-Talmud who walk in-darkness and-their-heart not with-their-people plastered-over the-sight of-their-eyes and-substance not in-them wise to-do-harm to-themselves for they hold-fast to-what is-short the-hand of-their-intellect and-they-want to-climb up-the-rungs there-is no doubt in-their-power to-reach to the-goal and-they-bring-up a-potsherd in-their-hands to-scrape with-it and-they-become to-them a-snare etc for comes-out from-them a-stumbling-block to-breach the-fence-of the-Torah and-to-tear-down its-walls one-boasts of-what that-he-not understands and-errs in-the-beginnings-of the-Torah and-its-roots and-comes to heresy and-he-is-distanced from-the-place-of his-desire which is the-nearness-of God it-would-have-been-better for-him that-he-not been-created all the-more-so to-change a-midrash and-to-annul a-halakhah thus far his words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:13.1 | 13) |  |  | Printed list marker, kept as the author's own (decisions, wave 1 rulings). |
| 3:13.2 | we learned |  |  |  |
| 3:13.3 | from the words of |  |  |  |
| 3:13.4 | the Gaon |  |  | The paragraph does not name the Gaon, and the English does not identify him. |
| 3:13.5 | that |  |  |  |
| 3:13.6 | many |  |  |  |
| 3:13.7 | books |  |  |  |
| 3:13.8 | there are |  |  |  |
| 3:13.9 | that are ascribed |  |  |  |
| 3:13.10 | them |  |  |  |
| 3:13.11 | to |  |  |  |
| 3:13.12 | the name of |  |  |  |
| 3:13.13 | the great |  |  |  |
| 3:13.14 | sages of |  |  |  |
| 3:13.15 | Israel |  |  |  |
| 3:13.16 | falsely |  |  |  |
| 3:13.17 | for the sake of |  |  |  |
| 3:13.18 | some |  |  |  |
| 3:13.19 | benefit |  |  |  |
| 3:13.20 | that accrues |  |  |  |
| 3:13.21 | to them |  |  |  |
| 3:13.22 | and we learned |  |  |  |
| 3:13.23 | also |  |  |  |
| 3:13.24 | that whatever |  |  |  |
| 3:13.25 | is not |  |  |  |
| 3:13.26 | stated explicitly |  |  |  |
| 3:13.27 | in the Mishnah |  |  |  |
| 3:13.28 | and in the Gemara |  |  |  |
| 3:13.29 | and we see |  |  |  |
| 3:13.30 | that it |  |  |  |
| 3:13.31 | the opposite of |  |  |  |
| 3:13.32 | the plain sense of |  |  |  |
| 3:13.33 | the Talmud |  |  |  |
| 3:13.34 | is not |  |  |  |
| 3:13.35 | a tradition |  |  | The Hebrew is קבלה, the word for Kabbalah. The author claims it for the Mishnah and Talmud, and here he denies that a teaching outside them is a tradition. |
| 3:13.36 | it |  |  |  |
| 3:13.37 | and we do not |  |  |  |
| 3:13.38 | rely |  |  |  |
| 3:13.39 | on it |  |  |  |
| 3:13.40 | And for |  |  |  |
| 3:13.41 | this reason |  |  |  |
| 3:13.42 | no |  |  |  |
| 3:13.43 | decisors |  |  |  |
| 3:13.44 | from |  |  |  |
| 3:13.45 | the midrashim |  |  |  |
| 3:13.46 | where |  |  |  |
| 3:13.47 | they |  |  |  |
| 3:13.48 | disagree |  |  |  |
| 3:13.49 | with |  |  |  |
| 3:13.50 | the Talmud |  |  |  |
| 3:13.51 | and even |  |  |  |
| 3:13.52 | the Talmud |  |  |  |
| 3:13.53 | Jerusalem |  |  |  |
| 3:13.54 | we do not |  |  |  |
| 3:13.55 | heed |  |  |  |
| 3:13.56 | it |  |  |  |
| 3:13.57 | where |  |  |  |
| 3:13.58 | it |  |  |  |
| 3:13.59 | disagrees |  |  |  |
| 3:13.60 | with |  |  |  |
| 3:13.61 | the Babylonian |  |  |  |
| 3:13.62 | which |  |  |  |
| 3:13.63 | all |  |  |  |
| 3:13.64 | the house of |  |  |  |
| 3:13.65 | Israel |  |  |  |
| 3:13.66 | rely |  |  |  |
| 3:13.67 | on it |  |  |  |
| 3:13.68 | both |  |  |  |
| 3:13.69 | for leniency |  |  |  |
| 3:13.70 | and |  |  |  |
| 3:13.71 | for stringency |  |  |  |
| 3:13.72 | as is made clear |  |  |  |
| 3:13.73 | all |  |  |  |
| 3:13.74 | this |  |  |  |
| 3:13.75 | in the words of |  |  |  |
| 3:13.76 | our rabbis |  |  |  |
| 3:13.77 | the decisors |  |  |  |
| 3:13.78 | in several |  |  |  |
| 3:13.79 | places |  |  |  |
| 3:13.80 | And see |  |  |  |
| 3:13.81 | below |  |  |  |
| 3:13.82 | section |  |  |  |
| 3:13.83 | 42 |  |  |  |
| 3:13.84 | what |  |  |  |
| 3:13.85 | I wrote |  |  |  |
| 3:13.86 | there |  |  |  |
| 3:13.87 | in the name of |  |  |  |
| 3:13.88 | Rabbeinu Asher |  |  | The Hebrew is Rabbeinu Asher, the Rosh. |
| 3:13.89 | And go and |  |  |  |
| 3:13.90 | learn |  |  |  |
| 3:13.91 | from what |  |  |  |
| 3:13.92 | wrote |  |  |  |
| 3:13.93 | Rabbi Tam |  |  | ר"ת is expanded as 'Rabbi Tam', the shorter reading. 'Rabbeinu Tam' is the other expansion, given in apparatus. |
| 3:13.94 | ben |  |  |  |
| 3:13.95 | Yichya |  |  |  |
| 3:13.96 | cited |  |  |  |
| 3:13.97 | in the book |  |  |  |
| 3:13.98 | Rivid |  |  | Rivid ha-Zahav: 'the golden chain'. The title is transliterated; rivid is a chain or necklace. |
| 3:13.99 | ha-Zahav |  |  |  |
| 3:13.100 | in the laws of |  |  |  |
| 3:13.101 | Passover |  |  |  |
| 3:13.102 | and these are his words: |  |  |  |
| 3:13.103 | In our time |  |  |  |
| 3:13.104 | - |  |  |  |
| 3:13.105 | has ceased |  |  |  |
| 3:13.106 | knowledge |  |  |  |
| 3:13.107 | and the hidden matters of |  |  |  |
| 3:13.108 | the Torah |  |  |  |
| 3:13.109 | have been hidden |  |  |  |
| 3:13.110 | from the eyes of |  |  |  |
| 3:13.111 | every |  |  |  |
| 3:13.112 | living thing |  |  |  |
| 3:13.113 | no |  |  |  |
| 3:13.114 | knows |  |  |  |
| 3:13.115 | man |  |  |  |
| 3:13.116 | the order of |  |  |  |
| 3:13.117 | the wisdom |  |  |  |
| 3:13.118 | this |  |  |  |
| 3:13.119 | and whoever |  |  |  |
| 3:13.120 | holds to |  |  |  |
| 3:13.121 | this |  |  |  |
| 3:13.122 | befalls |  |  |  |
| 3:13.123 | him |  |  |  |
| 3:13.124 | of the harm |  |  |  |
| 3:13.125 | and the danger |  |  |  |
| 3:13.126 | a share |  |  |  |
| 3:13.127 | great |  |  |  |
| 3:13.128 | that is to say |  |  | ר"ל: 'that is to say' (רוצה לומר). |
| 3:13.129 | the harm |  |  |  |
| 3:13.130 | and the danger of |  |  |  |
| 3:13.131 | heresy |  |  |  |
| 3:13.132 | in that there is no |  |  |  |
| 3:13.133 | for him |  |  |  |
| 3:13.134 | share |  |  |  |
| 3:13.135 | in the God of |  |  |  |
| 3:13.136 | Israel |  |  |  |
| 3:13.137 | And concerning |  |  |  |
| 3:13.138 | such |  |  |  |
| 3:13.139 | things |  |  |  |
| 3:13.140 | warned us |  |  |  |
| 3:13.141 | our Rabbis of blessed memory |  |  | רז"ל: 'our Rabbis of blessed memory'. |
| 3:13.142 | what is too wondrous |  |  |  |
| 3:13.143 | for you |  |  |  |
| 3:13.144 | do not |  |  |  |
| 3:13.145 | seek |  |  |  |
| 3:13.146 | and what is hidden |  |  |  |
| 3:13.147 | from you |  |  |  |
| 3:13.148 | do not |  |  |  |
| 3:13.149 | inquire into |  |  |  |
| 3:13.150 | And it is enough |  |  |  |
| 3:13.151 | for us |  |  |  |
| 3:13.152 | what |  |  |  |
| 3:13.153 | that taught us |  |  |  |
| 3:13.154 | the sages of |  |  |  |
| 3:13.155 | tradition |  |  | קבלה with האמתית: 'the true tradition' (owner's default, Q5), with the word Kabbalah in the Hebrew. |
| 3:13.156 | true |  |  |  |
| 3:13.157 | agreed upon |  |  |  |
| 3:13.158 | from every |  |  |  |
| 3:13.159 | side |  |  |  |
| 3:13.160 | in the explanation of |  |  |  |
| 3:13.161 | the Torah |  |  |  |
| 3:13.162 | and its commandments |  |  |  |
| 3:13.163 | according to |  |  | ע"פ: 'according to' (על פי). |
| 3:13.164 | the tradition |  |  |  |
| 3:13.165 | man |  |  |  |
| 3:13.166 | from the mouth of |  |  |  |
| 3:13.167 | man |  |  |  |
| 3:13.168 | until |  |  |  |
| 3:13.169 | the days of |  |  |  |
| 3:13.170 | Moses |  |  |  |
| 3:13.171 | and it is |  |  |  |
| 3:13.172 | called the |  |  |  |
| 3:13.173 | Torah |  |  |  |
| 3:13.174 | Oral |  |  | שבע"פ: 'of the mouth'; the Oral Torah. |
| 3:13.175 | which |  |  | The Hebrew is אשר בגללה, literally 'which on account of it'. |
| 3:13.176 | on account of it |  |  |  |
| 3:13.177 | went up |  |  |  |
| 3:13.178 | Moses |  |  |  |
| 3:13.179 | on high |  |  |  |
| 3:13.180 | and stood |  |  |  |
| 3:13.181 | there |  |  |  |
| 3:13.182 | forty |  |  |  |
| 3:13.183 | days |  |  |  |
| 3:13.184 | to learn |  |  |  |
| 3:13.185 | its general principles |  |  |  |
| 3:13.186 | and its particulars |  |  |  |
| 3:13.187 | by the rules |  |  | The Hebrew word is מדות, literally 'measures'. Here it is the rules transmitted to Moses. |
| 3:13.188 | transmitted |  |  |  |
| 3:13.189 | to him |  |  |  |
| 3:13.190 | from the mouth of |  |  |  |
| 3:13.191 | the Power |  |  | הגבורה is 'the Power', a name of God (decisions). |
| 3:13.192 | And with |  |  |  |
| 3:13.193 | these |  |  |  |
| 3:13.194 | we |  |  |  |
| 3:13.195 | are bound |  |  |  |
| 3:13.196 | to occupy ourselves |  |  |  |
| 3:13.197 | with them |  |  |  |
| 3:13.198 | for |  |  |  |
| 3:13.199 | they |  |  |  |
| 3:13.200 | are our life |  |  |  |
| 3:13.201 | and the length of |  |  |  |
| 3:13.202 | our days |  |  |  |
| 3:13.203 | and not one |  |  | וא' is ואחד, 'and one'; it is read with מהם and the negative in 205. |
| 3:13.204 | of them |  |  |  |
| 3:13.205 | not |  |  |  |
| 3:13.206 | we shall err |  |  |  |
| 3:13.207 | His one |  |  | אחת דתו is a set phrase, literally 'his one law'; the English keeps it. |
| 3:13.208 | law |  |  |  |
| 3:13.209 | is to put to death |  |  |  |
| 3:13.210 | whoever |  |  |  |
| 3:13.211 | defies |  |  |  |
| 3:13.212 | - |  |  |  |
| 3:13.213 | their word |  |  | Literally 'their mouth'; here it means what they say. |
| 3:13.214 | and whoever |  |  |  |
| 3:13.215 | transgresses |  |  |  |
| 3:13.216 | against |  |  |  |
| 3:13.217 | the words of |  |  |  |
| 3:13.218 | the Sages |  |  |  |
| 3:13.219 | is liable to |  |  |  |
| 3:13.220 | death |  |  |  |
| 3:13.221 | But |  |  |  |
| 3:13.222 | the remaining |  |  |  |
| 3:13.223 | allusions |  |  |  |
| 3:13.224 | which |  |  |  |
| 3:13.225 | in them |  |  |  |
| 3:13.226 | not |  |  |  |
| 3:13.227 | exhausted |  |  |  |
| 3:13.228 | in their lifetimes |  |  |  |
| 3:13.229 | the sages of |  |  |  |
| 3:13.230 | Kabbalah |  |  |  |
| 3:13.231 | although |  |  |  |
| 3:13.232 | they are |  |  |  |
| 3:13.233 | hidden matters of |  |  |  |
| 3:13.234 | the Torah |  |  |  |
| 3:13.235 | and matters of |  |  |  |
| 3:13.236 | the intellect |  |  |  |
| 3:13.237 | we were not |  |  |  |
| 3:13.238 | commanded |  |  |  |
| 3:13.239 | concerning them |  |  |  |
| 3:13.240 | For |  |  |  |
| 3:13.241 | these |  |  |  |
| 3:13.242 | remained |  |  |  |
| 3:13.243 | to be occupied |  |  |  |
| 3:13.244 | with them |  |  |  |
| 3:13.245 | the select |  |  |  |
| 3:13.246 | few |  |  |  |
| 3:13.247 | in the times |  |  |  |
| 3:13.248 | that |  |  |  |
| 3:13.249 | were |  |  |  |
| 3:13.250 | before us |  |  |  |
| 3:13.251 | etc. |  |  |  |
| 3:13.252 | And not |  |  |  |
| 3:13.253 | everyone |  |  |  |
| 3:13.254 | who wishes |  |  |  |
| 3:13.255 | to take |  |  |  |
| 3:13.256 | - |  |  |  |
| 3:13.257 | the Name |  |  |  |
| 3:13.258 | takes |  |  |  |
| 3:13.259 | For |  |  |  |
| 3:13.260 | matters |  |  |  |
| 3:13.261 | that |  |  |  |
| 3:13.262 | like its furnace |  |  | כבשונו is literally 'like its furnace'. The paragraph does not explain the figure, and the English keeps it. |
| 3:13.263 | forever |  |  |  |
| 3:13.264 | they are |  |  |  |
| 3:13.265 | And in our time |  |  |  |
| 3:13.266 | - |  |  |  |
| 3:13.267 | has ended |  |  |  |
| 3:13.268 | the merit |  |  |  |
| 3:13.269 | of the fathers |  |  |  |
| 3:13.270 | and no one |  |  |  |
| 3:13.271 | holds |  |  |  |
| 3:13.272 | justly |  |  |  |
| 3:13.273 | this |  |  |  |
| 3:13.274 | wisdom |  |  |  |
| 3:13.275 | On the contrary |  |  |  |
| 3:13.276 | they |  |  |  |
| 3:13.277 | destroy |  |  |  |
| 3:13.278 | the corners of |  |  |  |
| 3:13.279 | the Torah |  |  |  |
| 3:13.280 | and its pillars |  |  |  |
| 3:13.281 | they shudder |  |  |  |
| 3:13.282 | at the sound of |  |  |  |
| 3:13.283 | the trumpeters |  |  | מחצצים: 'trumpeters' (blowers of a trumpet). The sense of this clause is uncertain; the English keeps the words as read. |
| 3:13.284 | who pass over |  |  | העברים: 'those who pass over', literally 'the passers'. |
| 3:13.285 | the blind |  |  |  |
| 3:13.286 | do not |  |  |  |
| 3:13.287 | know |  |  |  |
| 3:13.288 | the way |  |  |  |
| 3:13.289 | where dwells |  |  |  |
| 3:13.290 | light |  |  |  |
| 3:13.291 | they grope |  |  |  |
| 3:13.292 | like the blind |  |  |  |
| 3:13.293 | and they weary themselves |  |  |  |
| 3:13.294 | seeking |  |  |  |
| 3:13.295 | the door |  |  |  |
| 3:13.296 | And where |  |  |  |
| 3:13.297 | they thought |  |  |  |
| 3:13.298 | to gain |  |  |  |
| 3:13.299 | suddenly |  |  |  |
| 3:13.300 | they will be broken |  |  |  |
| 3:13.301 | and there is |  |  |  |
| 3:13.302 | no healing |  |  |  |
| 3:13.303 | And many |  |  | The Hebrew is עמי הארץ, 'the unlearned' (decisions). |
| 3:13.304 | of the |  |  |  |
| 3:13.305 | unlearned |  |  |  |
| 3:13.306 | boast |  |  |  |
| 3:13.307 | saying |  |  |  |
| 3:13.308 | The secret of |  |  |  |
| 3:13.309 | the LORD |  |  |  |
| 3:13.310 | is ours |  |  |  |
| 3:13.311 | a heritage |  |  |  |
| 3:13.312 | for us |  |  |  |
| 3:13.313 | the waters |  |  |  |
| 3:13.314 | the upper |  |  |  |
| 3:13.315 | not |  |  |  |
| 3:13.316 | yours |  |  |  |
| 3:13.317 | you masters of |  |  |  |
| 3:13.318 | the Talmud |  |  |  |
| 3:13.319 | who walk |  |  |  |
| 3:13.320 | in darkness |  |  |  |
| 3:13.321 | And their heart |  |  |  |
| 3:13.322 | is not |  |  |  |
| 3:13.323 | with their people |  |  |  |
| 3:13.324 | is plastered over |  |  | טח: 'plastered over' (Isa 44:18, see apparatus). |
| 3:13.325 | the sight |  |  |  |
| 3:13.326 | their eyes |  |  |  |
| 3:13.327 | and substance |  |  |  |
| 3:13.328 | is not |  |  |  |
| 3:13.329 | in them |  |  |  |
| 3:13.330 | wise |  |  |  |
| 3:13.331 | to do harm |  |  |  |
| 3:13.332 | to themselves |  |  |  |
| 3:13.333 | For |  |  |  |
| 3:13.334 | they hold fast |  |  |  |
| 3:13.335 | to what |  |  |  |
| 3:13.336 | is too short |  |  |  |
| 3:13.337 | the hand |  |  |  |
| 3:13.338 | of their intellect |  |  | Literally 'the hand of their intellect is short'. |
| 3:13.339 | And they want |  |  |  |
| 3:13.340 | to climb |  |  |  |
| 3:13.341 | the rungs |  |  |  |
| 3:13.342 | there is no |  |  |  |
| 3:13.343 | doubt |  |  |  |
| 3:13.344 | in their power |  |  |  |
| 3:13.345 | to reach |  |  | The clause reads as affirmative. The irony lies in the next clause (the potsherd, Job 2:8), and the English keeps the words without adding any. |
| 3:13.346 | to |  |  |  |
| 3:13.347 | the goal |  |  |  |
| 3:13.348 | And they bring up |  |  |  |
| 3:13.349 | a potsherd |  |  |  |
| 3:13.350 | in their hands |  |  |  |
| 3:13.351 | to scrape |  |  |  |
| 3:13.352 | with it |  |  |  |
| 3:13.353 | And they become |  |  |  |
| 3:13.354 | to them |  |  |  |
| 3:13.355 | a snare |  |  |  |
| 3:13.356 | etc. |  |  |  |
| 3:13.357 | For |  |  |  |
| 3:13.358 | there comes out |  |  |  |
| 3:13.359 | from them |  |  |  |
| 3:13.360 | a stumbling-block |  |  |  |
| 3:13.361 | to breach |  |  |  |
| 3:13.362 | the fence of |  |  |  |
| 3:13.363 | the Torah |  |  |  |
| 3:13.364 | and to tear down |  |  |  |
| 3:13.365 | its walls |  |  |  |
| 3:13.366 | One boasts |  |  |  |
| 3:13.367 | of what |  |  |  |
| 3:13.368 | that he does not |  |  |  |
| 3:13.369 | understand |  |  |  |
| 3:13.370 | And he errs |  |  |  |
| 3:13.371 | in the beginnings of |  |  |  |
| 3:13.372 | the Torah |  |  |  |
| 3:13.373 | and its roots |  |  |  |
| 3:13.374 | and comes |  |  |  |
| 3:13.375 | to |  |  |  |
| 3:13.376 | heresy |  |  |  |
| 3:13.377 | And he is distanced |  |  |  |
| 3:13.378 | from the place |  |  |  |
| 3:13.379 | of his desire |  |  |  |
| 3:13.380 | which is |  |  |  |
| 3:13.381 | the nearness of |  |  |  |
| 3:13.382 | God |  |  | אלהים: 'God' (decisions). |
| 3:13.383 | it would have been better |  |  |  |
| 3:13.384 | for him |  |  |  |
| 3:13.385 | that he had not |  |  |  |
| 3:13.386 | been created |  |  |  |
| 3:13.387 | All |  |  |  |
| 3:13.388 | the more so |  |  |  |
| 3:13.389 | to change |  |  |  |
| 3:13.390 | a midrash |  |  |  |
| 3:13.391 | and to annul |  |  |  |
| 3:13.392 | a halakhah |  |  |  |
| 3:13.393 | Thus far his words. |  |  |  |

Phrases:
- 3:13.1 | 13) | 13) list marker
- 3:13.2 3:13.3 3:13.4 3:13.5 3:13.6 3:13.7 3:13.8 3:13.9 3:13.10 3:13.11 3:13.12 3:13.13 3:13.14 3:13.15 3:13.16 3:13.17 3:13.18 3:13.19 3:13.20 3:13.21 | We learned from the words of the Gaon that there are many books that are ascribed to the names of the great sages of Israel (falsely), for the sake of some benefit that accrues to them, | learned-we from-words-of the-Gaon that many books there-are that-ascribe them to the-name-of the-great sages-of Israel falsely for-the-sake-of some benefit that-accrues to-them
- 3:13.22 3:13.23 3:13.24 3:13.25 3:13.26 3:13.27 3:13.28 3:13.29 3:13.30 3:13.31 3:13.32 3:13.33 3:13.34 3:13.35 3:13.36 | and we learned also that whatever is not stated explicitly in the Mishnah and in the Gemara, and that we see is the opposite of the plain sense of the Talmud, is not a tradition, | and-learned-we also that-whatever is-not stated-explicitly in-the-Mishnah and-in-the-Gemara and-we-see that-it the-opposite-of the-plain-sense-of the-Talmud is-not tradition it
- 3:13.37 3:13.38 3:13.39 | and we do not rely on it. | and-not we-rely on-it
- 3:13.40 3:13.41 3:13.42 3:13.43 3:13.44 3:13.45 3:13.46 3:13.47 3:13.48 3:13.49 3:13.50 | And for this reason no decisors rule from the midrashim where they disagree with the Talmud, | and-for this-reason no decisors from the-midrashim where they disagree with the-Talmud
- 3:13.51 3:13.52 3:13.53 3:13.54 3:13.55 3:13.56 3:13.57 3:13.58 3:13.59 3:13.60 3:13.61 | and even the Jerusalem Talmud we do not heed where it disagrees with the Babylonian, | and-even the-Talmud Jerusalem not-we heed it where it disagrees with the-Babylonian
- 3:13.62 3:13.63 3:13.64 3:13.65 3:13.66 3:13.67 3:13.68 3:13.69 3:13.70 3:13.71 3:13.72 3:13.73 3:13.74 3:13.75 3:13.76 3:13.77 3:13.78 3:13.79 | on which all the house of Israel rely, both for leniency and for stringency, as all this is made clear in the words of our rabbis the decisors in several places. | which all the-house-of Israel rely on-it both for-leniency and for-stringency as-is-made-clear all this in-the-words-of our-rabbis the-decisors in-several places
- 3:13.80 3:13.81 3:13.82 3:13.83 3:13.84 3:13.85 3:13.86 3:13.87 3:13.88 | And see below, §42, what I wrote there in the name of Rabbeinu Asher, | and-see below section 42 what I-wrote there in-the-name-of Rabbeinu Asher
- 3:13.89 3:13.90 3:13.91 3:13.92 3:13.93 3:13.94 3:13.95 3:13.96 3:13.97 3:13.98 3:13.99 3:13.100 3:13.101 3:13.102 | And go and learn from what Rabbi Tam ben Yichya wrote, cited in the book Rivid ha-Zahav in the laws of Passover, and these are his words: | and-go-and learn from-what wrote Rabbi-Tam ben Yichya cited in-the-book Rivid ha-Zahav in-the-laws-of Passover and-these are-his-words
- 3:13.103 3:13.104 3:13.105 3:13.106 3:13.107 3:13.108 3:13.109 3:13.110 3:13.111 3:13.112 3:13.113 3:13.114 3:13.115 | "In our time knowledge has ceased, and the hidden matters of the Torah have been hidden from the eyes of every living thing; no man | in-our-time this has-ceased knowledge and-hidden-matters-of the-Torah have-been-hidden from-the-eyes-of every living-thing no knows man
- 3:13.116 3:13.117 3:13.118 | knows the order of this wisdom; | the-order-of the-wisdom this
- 3:13.119 3:13.120 3:13.121 3:13.122 3:13.123 3:13.124 3:13.125 3:13.126 3:13.127 | and whoever holds to this, there befalls him a great share of the harm and the danger | and-whoever holds-to this befalls him of-the-harm and-the-danger a-share great
- 3:13.128 3:13.129 3:13.130 3:13.131 3:13.132 3:13.133 3:13.134 3:13.135 3:13.136 | (that is to say, the harm and the danger of heresy, in that there is no share for him in the God of Israel). | that-is-to-say the-harm and-the-danger-of heresy in-that-there-is-no for-him share in-the-God-of Israel
- 3:13.137 3:13.138 3:13.139 3:13.140 3:13.141 3:13.142 3:13.143 3:13.144 3:13.145 3:13.146 3:13.147 3:13.148 3:13.149 | And concerning such things our Rabbis of blessed memory warned us, 'Do not seek what is too wondrous for you, and do not inquire into what is hidden from you.' | and-concerning such things warned-us our-Rabbis-of-blessed-memory what-is-too-wondrous for-you do-not seek and-what-is-hidden from-you do-not inquire
- 3:13.150 3:13.151 3:13.152 3:13.153 3:13.154 3:13.155 3:13.156 3:13.157 3:13.158 3:13.159 3:13.160 3:13.161 3:13.162 | And it is enough for us what the sages of the true tradition, agreed upon from every side, have taught us in the explanation of the Torah and its commandments. | and-enough for-us what that-taught-us the-sages-of tradition true agreed-upon from-every side in-the-explanation-of the-Torah and-its-commandments
- 3:13.163 3:13.164 3:13.165 3:13.166 3:13.167 3:13.168 3:13.169 3:13.170 | According to the tradition, from man to man, until the days of Moses; | according-to the-tradition man from-the-mouth-of man until the-days-of Moses
- 3:13.171 3:13.172 3:13.173 3:13.174 | and it is called the Oral Torah, | and-it is-called the Torah of-the-mouth
- 3:13.175 3:13.176 3:13.177 3:13.178 3:13.179 3:13.180 3:13.181 3:13.182 3:13.183 3:13.184 3:13.185 3:13.186 | on account of which Moses went up on high and stood there forty days to learn its general principles and its particulars | which on-account-of-it Moses went-up on-high and-stood there forty days to-learn its-general-principles and-its-particulars
- 3:13.187 3:13.188 3:13.189 3:13.190 3:13.191 | by the rules transmitted to him from the mouth of the Power. | by-the-measures transmitted to-him from-the-mouth-of the-Power
- 3:13.192 3:13.193 3:13.194 3:13.195 3:13.196 3:13.197 3:13.198 3:13.199 3:13.200 3:13.201 3:13.202 3:13.203 3:13.204 3:13.205 3:13.206 | And with these matters we are bound to occupy ourselves, for they are our life and the length of our days, and of them not one shall we err. | and-with these we are-bound to-occupy-ourselves with-them for they our-life and-the-length-of our-days and-one of-them not we-shall-err
- 3:13.207 3:13.208 3:13.209 3:13.210 3:13.211 3:13.212 3:13.213 3:13.214 3:13.215 3:13.216 3:13.217 3:13.218 3:13.219 3:13.220 | His one law is to put to death whoever defies their word, and whoever transgresses the words of the Sages is liable to death. | his-one law to-put-to-death who defies the-mouth-of-them and-whoever transgresses against the-words-of the-Sages is-liable death
- 3:13.221 3:13.222 3:13.223 3:13.224 3:13.225 3:13.226 3:13.227 3:13.228 3:13.229 3:13.230 | But the remaining allusions in them, which the sages of Kabbalah did not exhaust in their lifetimes, | but the-remaining allusions which in-them not exhausted their-days the-sages-of Kabbalah
- 3:13.231 3:13.232 3:13.233 3:13.234 3:13.235 3:13.236 3:13.237 3:13.238 3:13.239 | although they are hidden matters of the Torah and matters of the intellect; we were not commanded concerning them. | although they-are hidden-matters-of the-Torah and-matters-of the-intellect we were-not commanded concerning-them
- 3:13.240 3:13.241 3:13.242 3:13.243 3:13.244 3:13.245 3:13.246 3:13.247 3:13.248 3:13.249 3:13.250 3:13.251 | For these remained for the select few to be occupied with them, in the times that were before us, etc. | for these remained to-be-occupied with-them the-select few in-the-times that were before-us etc
- 3:13.252 3:13.253 3:13.254 3:13.255 3:13.256 3:13.257 3:13.258 | And not everyone who wishes to take up the Name takes it. | and-not everyone who-wishes to-take the-Name takes
- 3:13.259 3:13.260 3:13.261 3:13.262 3:13.263 3:13.264 | For these are matters that are like its furnace; they are so forever. | for matters that like-its-furnace forever they are
- 3:13.265 3:13.266 3:13.267 3:13.268 3:13.269 3:13.270 3:13.271 3:13.272 3:13.273 3:13.274 | And in our time the merit of the fathers has ended, and no one holds this wisdom justly: | and-in-our-time this has-ended the-merit of-the-fathers and-no-one holds justly this wisdom
- 3:13.275 3:13.276 3:13.277 3:13.278 3:13.279 3:13.280 | on the contrary, they destroy the corners of the Torah and its pillars; | on-the-contrary they destroy the-corners-of the-Torah and-its-pillars
- 3:13.281 3:13.282 3:13.283 3:13.284 3:13.285 3:13.286 3:13.287 3:13.288 3:13.289 3:13.290 | they shudder at the sound of the trumpeters; the blind, who pass over, do not know the way where light dwells; | they-shudder at-the-sound-of the-trumpeters the-passers-over the-blind not know the-way where dwells light
- 3:13.291 3:13.292 3:13.293 3:13.294 3:13.295 | they grope like the blind, and weary themselves seeking the door. | they-grope like-the-blind and-they-weary seeking the-door
- 3:13.296 3:13.297 3:13.298 3:13.299 3:13.300 3:13.301 3:13.302 | And where they thought to gain, suddenly they shall be broken, and there is no healing. | and-where they-thought to-gain suddenly they-shall-be-broken and-there-is no healing
- 3:13.303 3:13.304 3:13.305 3:13.306 3:13.307 3:13.308 3:13.309 3:13.310 | And many of the unlearned boast, saying: 'The secret of the LORD is ours; | and-many of-the-unlearned boast saying the-secret-of the-LORD for-us
- 3:13.311 3:13.312 3:13.313 3:13.314 3:13.315 3:13.316 3:13.317 3:13.318 3:13.319 3:13.320 | the upper waters are our heritage, not yours, you masters of the Talmud who walk in darkness.' | a-heritage for-us the-waters the-upper not for-you the-masters-of the-Talmud who walk in-darkness
- 3:13.321 3:13.322 3:13.323 3:13.324 3:13.325 3:13.326 | Their heart is not with their people; their eyes are plastered over, | and-their-heart not with-their-people plastered-over the-sight of-their-eyes
- 3:13.327 3:13.328 3:13.329 3:13.330 3:13.331 3:13.332 | and there is no substance in them; they are wise to do harm to themselves. | and-substance not in-them wise to-do-harm to-themselves
- 3:13.333 3:13.334 3:13.335 3:13.336 3:13.337 3:13.338 | For they hold fast to what the hand of their intellect is too short to reach. | for they hold-fast to-what is-short the-hand of-their-intellect
- 3:13.339 3:13.340 3:13.341 3:13.342 3:13.343 3:13.344 3:13.345 3:13.346 3:13.347 | And they want to climb the rungs; there is no doubt that it is in their power to reach the goal. | and-they-want to-climb up-the-rungs there-is no doubt in-their-power to-reach to the-goal
- 3:13.348 3:13.349 3:13.350 3:13.351 3:13.352 | And they bring up a potsherd in their hands to scrape with it. | and-they-bring-up a-potsherd in-their-hands to-scrape with-it
- 3:13.353 3:13.354 3:13.355 3:13.356 | And they become a snare to them, etc. | and-they-become to-them a-snare etc
- 3:13.357 3:13.358 3:13.359 3:13.360 3:13.361 3:13.362 3:13.363 3:13.364 3:13.365 | For there comes out from them a stumbling-block, to breach the fence of the Torah and to tear down its walls. | for comes-out from-them a-stumbling-block to-breach the-fence-of the-Torah and-to-tear-down its-walls
- 3:13.366 3:13.367 3:13.368 3:13.369 3:13.370 3:13.371 3:13.372 3:13.373 3:13.374 3:13.375 3:13.376 | One boasts of what he does not understand, and errs in the beginnings of the Torah and its roots, and comes to heresy. | one-boasts of-what that-he-not understands and-errs in-the-beginnings-of the-Torah and-its-roots and-comes to heresy
- 3:13.377 3:13.378 3:13.379 3:13.380 3:13.381 3:13.382 3:13.383 3:13.384 3:13.385 3:13.386 | And he is distanced from the place of his desire, which is the nearness of God; it would have been better for him had he not been created, | and-he-is-distanced from-the-place-of his-desire which is the-nearness-of God it-would-have-been-better for-him that-he-not been-created
- 3:13.387 3:13.388 3:13.389 3:13.390 3:13.391 3:13.392 | all the more so to change a midrash and to annul a halakhah." | all the-more-so to-change a-midrash and-to-annul a-halakhah
- 3:13.393 | Thus far his words. | thus far his words

Grammar notes:
- The paragraph is the author's own argument in its first half. From וז"ל (3:13.102) to עכ"ל (3:13.393) the words are Rabbi Tam's, as cited in Rivid ha-Zahav, and the quotation is marked in double quotation marks.
- ולא סמכינן עליה (3:13.37-39): סמכינן is the Talmudic Aramaic first-person plural of 'to rely', and the English gives it as the plain 'we do not rely on it'.
- The subject of יתפלצון (3:13.281) is not marked. The English takes the destroyers of the Torah's corners as its subject; a reader may take the blind as the subject instead.
- The speech of the unlearned (3:13.307-320) has no introductory formula beyond באמרם. It is marked as direct speech inside the quotation, in single quotation marks.
- אין ספק בידם להגיע (3:13.342-345) reads as an affirmative. The irony is carried by the next clause, the potsherd of 3:13.349-352.

Translator's note: Paragraph 13 of the book's third chapter. The opening 'we learned' frame is the author's; the central quotation is Rabbi Tam's, and it closes with the author's 'Thus far his words'.

Commentary rulings:
- apparatus: Base text: Wikisource, as translated for the unit under Q2 (the 1931 scan is not yet collated for this paragraph).
- apparatus: ר"ת (3:13.93): 'Rabbeinu Tam' is the other expansion; 'Rabbi Tam' is kept as the shorter.
- apparatus: Isa 44:18: טח מראות עיניהם (their eyes plastered over).
- apparatus: Job 38:19: אי זה הדרך ישכן אור; the author writes דרך ישכן אור.
- apparatus: Isa 59:10: ימששו כעורים (they grope like the blind).
- apparatus: Ps 25:14: סוד ה' ליראיו; the author writes סוד ה' לנו.
- apparatus: Prov 6:15: פתע יבוא ... ישבר ואין מרפא; the author writes פתע ישברו ואין מרפא.
- apparatus: Isa 59:1: הנה לא קצרה יד ה'; echoed in קצרה יד שכלם (3:13.336-338).
- apparatus: Job 2:8: ויקח לו חרש להתגרד בו; echoed in חרס בידם להתגרד בו (3:13.349-351).
- apparatus: Jer 4:22: חכמים המה להרע; echoed in חכמים להרע לנפשם (3:13.330-332).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

