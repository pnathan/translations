# Translation pack: milhamot-hashem perek-9-06 (Milchamot ha-Shem 9:7–9:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-06`.

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

- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
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
- מעמד: "assembly" or "occasion" by context; "the Standing at Sinai" where the
  author names that event.
- והוא שלא: "provided that" (fixed).
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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


## Milchamot ha-Shem 9:7

Hebrew (Wikisource): קכה) וכל הדברים האלה קמו עמדו להאמין בהם בלי קולות וברקים ולפידים וענן וקול שופר ואפי' אות ומופת כמו שאמר הכתוב בנביא השקר ונתן אליך אות או מופת לא שמענו ולא ראינו ולא בת קול מן השמים קראה באזנינו רק דיו שחור על נייר לבן סיפורי פלאות בדויים משם עתיקא דקישא שנראה לבית מדרשו של רשב"י המשקר (ולא לרשב"י תנא דידן ח"ו על כבודו) ושאל לו על פסוק ויאמר אלהים נעשה אדם וכו' כדאיתא בזוהר בראשית כ"ב. היעלה על לב ישראל לשאול על אחד מעשרה מאמרות שבהן נברא העולם מי האומר מאמר זה ומי האומר מאמר זה או על דעתו לומר כי האומר ראו עתה כי אני אני הוא ואין אלהים עמדי ולא זהו שאמר כל העשרה מאמרות? וגם הרבה ספורים משם אליהו ורעיא מהימנא ואחרים מהתנאים והאמוראים והאחרונים אשר כל משכיל בתורה התמימה ודברי נביאינו וחז"ל ישום וישרוק ועמד משתאה ומשתומם איך נוציא ישן אשר שמענו במעמד הר סיני הנבחר ויקם עדות ביעקב ותורה שם בישראל להודיעם לבניהם בקולות וברקים ולפידים וענן כבד וקול ה' בכח מדבר מתוך האש אשר לא עשה כן לכל גוי ובקול עוז ותעצומות דבר עמנו פנים בפנים אנכי ה' אלהיך אשר הוציאתיך וכו' ונאמין שאין לי אב ולא אח ולא בן ראו עתה כי אני אני הוא ואין אלהים עמדי וכו' ונאמין למלך העכו"ם שהביא לנו תורה שבעל פה חדשה האומרת כי האלהים ברא מעצמותו חמשת פרצופין והשליטם על מעשה בארית והם השליטו הזעיר אנפין על כל הברואים וצוו את הכל לעבדו והוא אלהינו ואנחנו עמו ונחלתו?

| id | word | punctuation after |
|---|---|---|
| 9:7.1 | קכה | ) |
| 9:7.2 | וכל |  |
| 9:7.3 | הדברים |  |
| 9:7.4 | האלה |  |
| 9:7.5 | קמו |  |
| 9:7.6 | עמדו |  |
| 9:7.7 | להאמין |  |
| 9:7.8 | בהם |  |
| 9:7.9 | בלי |  |
| 9:7.10 | קולות |  |
| 9:7.11 | וברקים |  |
| 9:7.12 | ולפידים |  |
| 9:7.13 | וענן |  |
| 9:7.14 | וקול |  |
| 9:7.15 | שופר |  |
| 9:7.16 | ואפי | ' |
| 9:7.17 | אות |  |
| 9:7.18 | ומופת |  |
| 9:7.19 | כמו |  |
| 9:7.20 | שאמר |  |
| 9:7.21 | הכתוב |  |
| 9:7.22 | בנביא |  |
| 9:7.23 | השקר |  |
| 9:7.24 | ונתן |  |
| 9:7.25 | אליך |  |
| 9:7.26 | אות |  |
| 9:7.27 | או |  |
| 9:7.28 | מופת |  |
| 9:7.29 | לא |  |
| 9:7.30 | שמענו |  |
| 9:7.31 | ולא |  |
| 9:7.32 | ראינו |  |
| 9:7.33 | ולא |  |
| 9:7.34 | בת |  |
| 9:7.35 | קול |  |
| 9:7.36 | מן |  |
| 9:7.37 | השמים |  |
| 9:7.38 | קראה |  |
| 9:7.39 | באזנינו |  |
| 9:7.40 | רק |  |
| 9:7.41 | דיו |  |
| 9:7.42 | שחור |  |
| 9:7.43 | על |  |
| 9:7.44 | נייר |  |
| 9:7.45 | לבן |  |
| 9:7.46 | סיפורי |  |
| 9:7.47 | פלאות |  |
| 9:7.48 | בדויים |  |
| 9:7.49 | משם |  |
| 9:7.50 | עתיקא |  |
| 9:7.51 | דקישא |  |
| 9:7.52 | שנראה |  |
| 9:7.53 | לבית |  |
| 9:7.54 | מדרשו |  |
| 9:7.55 | של |  |
| 9:7.56 | רשב"י |  |
| 9:7.57 | המשקר |  |
| 9:7.58 | ולא |  |
| 9:7.59 | לרשב"י |  |
| 9:7.60 | תנא |  |
| 9:7.61 | דידן |  |
| 9:7.62 | ח"ו |  |
| 9:7.63 | על |  |
| 9:7.64 | כבודו | ) |
| 9:7.65 | ושאל |  |
| 9:7.66 | לו |  |
| 9:7.67 | על |  |
| 9:7.68 | פסוק |  |
| 9:7.69 | ויאמר |  |
| 9:7.70 | אלהים |  |
| 9:7.71 | נעשה |  |
| 9:7.72 | אדם |  |
| 9:7.73 | וכו | ' |
| 9:7.74 | כדאיתא |  |
| 9:7.75 | בזוהר |  |
| 9:7.76 | בראשית |  |
| 9:7.77 | כ"ב | . |
| 9:7.78 | היעלה |  |
| 9:7.79 | על |  |
| 9:7.80 | לב |  |
| 9:7.81 | ישראל |  |
| 9:7.82 | לשאול |  |
| 9:7.83 | על |  |
| 9:7.84 | אחד |  |
| 9:7.85 | מעשרה |  |
| 9:7.86 | מאמרות |  |
| 9:7.87 | שבהן |  |
| 9:7.88 | נברא |  |
| 9:7.89 | העולם |  |
| 9:7.90 | מי |  |
| 9:7.91 | האומר |  |
| 9:7.92 | מאמר |  |
| 9:7.93 | זה |  |
| 9:7.94 | ומי |  |
| 9:7.95 | האומר |  |
| 9:7.96 | מאמר |  |
| 9:7.97 | זה |  |
| 9:7.98 | או |  |
| 9:7.99 | על |  |
| 9:7.100 | דעתו |  |
| 9:7.101 | לומר |  |
| 9:7.102 | כי |  |
| 9:7.103 | האומר |  |
| 9:7.104 | ראו |  |
| 9:7.105 | עתה |  |
| 9:7.106 | כי |  |
| 9:7.107 | אני |  |
| 9:7.108 | אני |  |
| 9:7.109 | הוא |  |
| 9:7.110 | ואין |  |
| 9:7.111 | אלהים |  |
| 9:7.112 | עמדי |  |
| 9:7.113 | ולא |  |
| 9:7.114 | זהו |  |
| 9:7.115 | שאמר |  |
| 9:7.116 | כל |  |
| 9:7.117 | העשרה |  |
| 9:7.118 | מאמרות | ? |
| 9:7.119 | וגם |  |
| 9:7.120 | הרבה |  |
| 9:7.121 | ספורים |  |
| 9:7.122 | משם |  |
| 9:7.123 | אליהו |  |
| 9:7.124 | ורעיא |  |
| 9:7.125 | מהימנא |  |
| 9:7.126 | ואחרים |  |
| 9:7.127 | מהתנאים |  |
| 9:7.128 | והאמוראים |  |
| 9:7.129 | והאחרונים |  |
| 9:7.130 | אשר |  |
| 9:7.131 | כל |  |
| 9:7.132 | משכיל |  |
| 9:7.133 | בתורה |  |
| 9:7.134 | התמימה |  |
| 9:7.135 | ודברי |  |
| 9:7.136 | נביאינו |  |
| 9:7.137 | וחז"ל |  |
| 9:7.138 | ישום |  |
| 9:7.139 | וישרוק |  |
| 9:7.140 | ועמד |  |
| 9:7.141 | משתאה |  |
| 9:7.142 | ומשתומם |  |
| 9:7.143 | איך |  |
| 9:7.144 | נוציא |  |
| 9:7.145 | ישן |  |
| 9:7.146 | אשר |  |
| 9:7.147 | שמענו |  |
| 9:7.148 | במעמד |  |
| 9:7.149 | הר |  |
| 9:7.150 | סיני |  |
| 9:7.151 | הנבחר |  |
| 9:7.152 | ויקם |  |
| 9:7.153 | עדות |  |
| 9:7.154 | ביעקב |  |
| 9:7.155 | ותורה |  |
| 9:7.156 | שם |  |
| 9:7.157 | בישראל |  |
| 9:7.158 | להודיעם |  |
| 9:7.159 | לבניהם |  |
| 9:7.160 | בקולות |  |
| 9:7.161 | וברקים |  |
| 9:7.162 | ולפידים |  |
| 9:7.163 | וענן |  |
| 9:7.164 | כבד |  |
| 9:7.165 | וקול |  |
| 9:7.166 | ה | ' |
| 9:7.167 | בכח |  |
| 9:7.168 | מדבר |  |
| 9:7.169 | מתוך |  |
| 9:7.170 | האש |  |
| 9:7.171 | אשר |  |
| 9:7.172 | לא |  |
| 9:7.173 | עשה |  |
| 9:7.174 | כן |  |
| 9:7.175 | לכל |  |
| 9:7.176 | גוי |  |
| 9:7.177 | ובקול |  |
| 9:7.178 | עוז |  |
| 9:7.179 | ותעצומות |  |
| 9:7.180 | דבר |  |
| 9:7.181 | עמנו |  |
| 9:7.182 | פנים |  |
| 9:7.183 | בפנים |  |
| 9:7.184 | אנכי |  |
| 9:7.185 | ה | ' |
| 9:7.186 | אלהיך |  |
| 9:7.187 | אשר |  |
| 9:7.188 | הוציאתיך |  |
| 9:7.189 | וכו | ' |
| 9:7.190 | ונאמין |  |
| 9:7.191 | שאין |  |
| 9:7.192 | לי |  |
| 9:7.193 | אב |  |
| 9:7.194 | ולא |  |
| 9:7.195 | אח |  |
| 9:7.196 | ולא |  |
| 9:7.197 | בן |  |
| 9:7.198 | ראו |  |
| 9:7.199 | עתה |  |
| 9:7.200 | כי |  |
| 9:7.201 | אני |  |
| 9:7.202 | אני |  |
| 9:7.203 | הוא |  |
| 9:7.204 | ואין |  |
| 9:7.205 | אלהים |  |
| 9:7.206 | עמדי |  |
| 9:7.207 | וכו | ' |
| 9:7.208 | ונאמין |  |
| 9:7.209 | למלך |  |
| 9:7.210 | העכו"ם |  |
| 9:7.211 | שהביא |  |
| 9:7.212 | לנו |  |
| 9:7.213 | תורה |  |
| 9:7.214 | שבעל |  |
| 9:7.215 | פה |  |
| 9:7.216 | חדשה |  |
| 9:7.217 | האומרת |  |
| 9:7.218 | כי |  |
| 9:7.219 | האלהים |  |
| 9:7.220 | ברא |  |
| 9:7.221 | מעצמותו |  |
| 9:7.222 | חמשת |  |
| 9:7.223 | פרצופין |  |
| 9:7.224 | והשליטם |  |
| 9:7.225 | על |  |
| 9:7.226 | מעשה |  |
| 9:7.227 | בארית |  |
| 9:7.228 | והם |  |
| 9:7.229 | השליטו |  |
| 9:7.230 | הזעיר |  |
| 9:7.231 | אנפין |  |
| 9:7.232 | על |  |
| 9:7.233 | כל |  |
| 9:7.234 | הברואים |  |
| 9:7.235 | וצוו |  |
| 9:7.236 | את |  |
| 9:7.237 | הכל |  |
| 9:7.238 | לעבדו |  |
| 9:7.239 | והוא |  |
| 9:7.240 | אלהינו |  |
| 9:7.241 | ואנחנו |  |
| 9:7.242 | עמו |  |
| 9:7.243 | ונחלתו | ? |


## Milchamot ha-Shem 9:8

Hebrew (Wikisource): ואם ח"ו רצונו ית' להמיר כבודו באל אחר קצר אפים ולהשליטו על כל הברואים כמו שאמר הזהר שולטנו אתיהיב ליה וכו' היפך מה שצונו בתחלה תמר בו ואמרו רז"ל אל תמירני בו ואל תמירהו בי (רזה פרשת משפטים) מה לא יט שמים וירד על איזה נר שיבחר וידבר עמנו עוד הפעם לאמר עתה בסוף הימים נלאתי כלכל והמלכתי במקומי את קצר האפים? זאת לא תהיה לא זו הדרך ולא זו העיר ולא יחליף האל ולא ימיר ולא יעשה כן בישראל להמיר את אלהינו אשר הודיע דרכיו למשה כי הוא ארך אפים ורב חסד באל אחר קצר אפים ובעל אשה. ורז"ל אמרו ברית כרותה לשלש עשרה מדות של רחמים שאינן חוזרות ריקם ולכך קבעו בתפלות של פרקים על כל צרה שלא תבוא וגם בימים נוראים להגביר הי"ג מדות של רחמים ואיך תיסק אדעתין ליחסם לזעיר אנפין שהוא קצר אפים הא הוו הני מילי סתרן אחדדי קורים אנו אותו ז"א שעניינו קצר אפים ומשבחים אותו בההיפך שהוא ארך אפים? אתמהא?

| id | word | punctuation after |
|---|---|---|
| 9:8.1 | ואם |  |
| 9:8.2 | ח"ו |  |
| 9:8.3 | רצונו |  |
| 9:8.4 | ית | ' |
| 9:8.5 | להמיר |  |
| 9:8.6 | כבודו |  |
| 9:8.7 | באל |  |
| 9:8.8 | אחר |  |
| 9:8.9 | קצר |  |
| 9:8.10 | אפים |  |
| 9:8.11 | ולהשליטו |  |
| 9:8.12 | על |  |
| 9:8.13 | כל |  |
| 9:8.14 | הברואים |  |
| 9:8.15 | כמו |  |
| 9:8.16 | שאמר |  |
| 9:8.17 | הזהר |  |
| 9:8.18 | שולטנו |  |
| 9:8.19 | אתיהיב |  |
| 9:8.20 | ליה |  |
| 9:8.21 | וכו | ' |
| 9:8.22 | היפך |  |
| 9:8.23 | מה |  |
| 9:8.24 | שצונו |  |
| 9:8.25 | בתחלה |  |
| 9:8.26 | תמר |  |
| 9:8.27 | בו |  |
| 9:8.28 | ואמרו |  |
| 9:8.29 | רז"ל |  |
| 9:8.30 | אל |  |
| 9:8.31 | תמירני |  |
| 9:8.32 | בו |  |
| 9:8.33 | ואל |  |
| 9:8.34 | תמירהו |  |
| 9:8.35 | בי |  |
| 9:8.36 | רזה |  |
| 9:8.37 | פרשת |  |
| 9:8.38 | משפטים | ) |
| 9:8.39 | מה |  |
| 9:8.40 | לא |  |
| 9:8.41 | יט |  |
| 9:8.42 | שמים |  |
| 9:8.43 | וירד |  |
| 9:8.44 | על |  |
| 9:8.45 | איזה |  |
| 9:8.46 | נר |  |
| 9:8.47 | שיבחר |  |
| 9:8.48 | וידבר |  |
| 9:8.49 | עמנו |  |
| 9:8.50 | עוד |  |
| 9:8.51 | הפעם |  |
| 9:8.52 | לאמר |  |
| 9:8.53 | עתה |  |
| 9:8.54 | בסוף |  |
| 9:8.55 | הימים |  |
| 9:8.56 | נלאתי |  |
| 9:8.57 | כלכל |  |
| 9:8.58 | והמלכתי |  |
| 9:8.59 | במקומי |  |
| 9:8.60 | את |  |
| 9:8.61 | קצר |  |
| 9:8.62 | האפים | ? |
| 9:8.63 | זאת |  |
| 9:8.64 | לא |  |
| 9:8.65 | תהיה |  |
| 9:8.66 | לא |  |
| 9:8.67 | זו |  |
| 9:8.68 | הדרך |  |
| 9:8.69 | ולא |  |
| 9:8.70 | זו |  |
| 9:8.71 | העיר |  |
| 9:8.72 | ולא |  |
| 9:8.73 | יחליף |  |
| 9:8.74 | האל |  |
| 9:8.75 | ולא |  |
| 9:8.76 | ימיר |  |
| 9:8.77 | ולא |  |
| 9:8.78 | יעשה |  |
| 9:8.79 | כן |  |
| 9:8.80 | בישראל |  |
| 9:8.81 | להמיר |  |
| 9:8.82 | את |  |
| 9:8.83 | אלהינו |  |
| 9:8.84 | אשר |  |
| 9:8.85 | הודיע |  |
| 9:8.86 | דרכיו |  |
| 9:8.87 | למשה |  |
| 9:8.88 | כי |  |
| 9:8.89 | הוא |  |
| 9:8.90 | ארך |  |
| 9:8.91 | אפים |  |
| 9:8.92 | ורב |  |
| 9:8.93 | חסד |  |
| 9:8.94 | באל |  |
| 9:8.95 | אחר |  |
| 9:8.96 | קצר |  |
| 9:8.97 | אפים |  |
| 9:8.98 | ובעל |  |
| 9:8.99 | אשה | . |
| 9:8.100 | ורז"ל |  |
| 9:8.101 | אמרו |  |
| 9:8.102 | ברית |  |
| 9:8.103 | כרותה |  |
| 9:8.104 | לשלש |  |
| 9:8.105 | עשרה |  |
| 9:8.106 | מדות |  |
| 9:8.107 | של |  |
| 9:8.108 | רחמים |  |
| 9:8.109 | שאינן |  |
| 9:8.110 | חוזרות |  |
| 9:8.111 | ריקם |  |
| 9:8.112 | ולכך |  |
| 9:8.113 | קבעו |  |
| 9:8.114 | בתפלות |  |
| 9:8.115 | של |  |
| 9:8.116 | פרקים |  |
| 9:8.117 | על |  |
| 9:8.118 | כל |  |
| 9:8.119 | צרה |  |
| 9:8.120 | שלא |  |
| 9:8.121 | תבוא |  |
| 9:8.122 | וגם |  |
| 9:8.123 | בימים |  |
| 9:8.124 | נוראים |  |
| 9:8.125 | להגביר |  |
| 9:8.126 | הי"ג |  |
| 9:8.127 | מדות |  |
| 9:8.128 | של |  |
| 9:8.129 | רחמים |  |
| 9:8.130 | ואיך |  |
| 9:8.131 | תיסק |  |
| 9:8.132 | אדעתין |  |
| 9:8.133 | ליחסם |  |
| 9:8.134 | לזעיר |  |
| 9:8.135 | אנפין |  |
| 9:8.136 | שהוא |  |
| 9:8.137 | קצר |  |
| 9:8.138 | אפים |  |
| 9:8.139 | הא |  |
| 9:8.140 | הוו |  |
| 9:8.141 | הני |  |
| 9:8.142 | מילי |  |
| 9:8.143 | סתרן |  |
| 9:8.144 | אחדדי |  |
| 9:8.145 | קורים |  |
| 9:8.146 | אנו |  |
| 9:8.147 | אותו |  |
| 9:8.148 | ז"א |  |
| 9:8.149 | שעניינו |  |
| 9:8.150 | קצר |  |
| 9:8.151 | אפים |  |
| 9:8.152 | ומשבחים |  |
| 9:8.153 | אותו |  |
| 9:8.154 | בההיפך |  |
| 9:8.155 | שהוא |  |
| 9:8.156 | ארך |  |
| 9:8.157 | אפים | ? |
| 9:8.158 | אתמהא | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

