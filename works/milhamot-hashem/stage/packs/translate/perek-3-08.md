# Translation pack: milhamot-hashem perek-3-08 (Milchamot ha-Shem 3:13–3:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-08`.

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


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

