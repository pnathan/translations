# Review pack: milhamot-hashem perek-2-18 (Milchamot ha-Shem 2:34–2:34)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-18`.

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
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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


## Milchamot ha-Shem 2:34

Hebrew (Wikisource): תשובה אשאלה מאלהי ינחני במעגלי יושר ויצילנו משפת שקר, מלשון רמיה. הוריני ה' דרכיך אהלך באמתך. יחד לבבי ליראה שמך ושלח אורך ואמתך המה ינחוני להשיב דבר דבור על אפניו. גל מעלי חרפה ובוז כי עדותיך נצרתי. גם כי ישבו שרים בי נדברו, עבדך ישיח בחוקיך. אז לא אבוש כאשר לא אסור מן הדרך אשר למדונו רז"ל בעלי המשנה והתלמוד והמדרשים האמתיים והפוסקים ההולכים בעקבותיהם. אשר כל בית ישראל נשענים עליהם. כי המה בעלי הקבלה האמתית אשר קבלו ממרע"ה כללותיה ופרטותיה של תורה אשר עליהם עמד בסיני ארבעים יום וארבעים לילה ומסרן ליהושע ויהושע לזקנים וכו' כדאיתא במס' אבות. א) וחוב קדוש מוטל עלינו ללכת בעקבותיהם בכל מידי דנפקא מינה לדינא בכל ענייני איסור והתר וטמא וטהור ופסול וכשר. וכ"ש בענייני יחוד השי"ת שהוא העיקר הגדול ואחריו קבלת כל המצות וכיצד עשייתן כמאמר התנא ע"ה למה קדמה פרשת שמע לוהיה אם שמוע וכו'! וכל מחבר ספר אשר יבוא להוסיף או לגרוע מדבריהם וכ"ש להטותינו מאמונת תוה"ק כפי קבלת רבותינו בעלי המשנה והתלמוד. לא נשמע לו. ואף אם יעש אותות ומופתים בשמים ובארץ לא נסור מאחרי ה' אלהינו על פי קבלת רז"ל בעלי המשנה והתלמוד והפוסקים הבאים אחריהם ההולכים בעקבותיהם, וגם אם מן השמים תקרא בת קול שמעו אליו לבור מקבלתם לא משגחינן בה כמאמר ר' יהושע אין משגיחין בבת קול כי לא בשמים היא.

| id | word | punctuation after |
|---|---|---|
| 2:34.1 | תשובה |  |
| 2:34.2 | אשאלה |  |
| 2:34.3 | מאלהי |  |
| 2:34.4 | ינחני |  |
| 2:34.5 | במעגלי |  |
| 2:34.6 | יושר |  |
| 2:34.7 | ויצילנו |  |
| 2:34.8 | משפת |  |
| 2:34.9 | שקר | , |
| 2:34.10 | מלשון |  |
| 2:34.11 | רמיה | . |
| 2:34.12 | הוריני |  |
| 2:34.13 | ה | ' |
| 2:34.14 | דרכיך |  |
| 2:34.15 | אהלך |  |
| 2:34.16 | באמתך | . |
| 2:34.17 | יחד |  |
| 2:34.18 | לבבי |  |
| 2:34.19 | ליראה |  |
| 2:34.20 | שמך |  |
| 2:34.21 | ושלח |  |
| 2:34.22 | אורך |  |
| 2:34.23 | ואמתך |  |
| 2:34.24 | המה |  |
| 2:34.25 | ינחוני |  |
| 2:34.26 | להשיב |  |
| 2:34.27 | דבר |  |
| 2:34.28 | דבור |  |
| 2:34.29 | על |  |
| 2:34.30 | אפניו | . |
| 2:34.31 | גל |  |
| 2:34.32 | מעלי |  |
| 2:34.33 | חרפה |  |
| 2:34.34 | ובוז |  |
| 2:34.35 | כי |  |
| 2:34.36 | עדותיך |  |
| 2:34.37 | נצרתי | . |
| 2:34.38 | גם |  |
| 2:34.39 | כי |  |
| 2:34.40 | ישבו |  |
| 2:34.41 | שרים |  |
| 2:34.42 | בי |  |
| 2:34.43 | נדברו | , |
| 2:34.44 | עבדך |  |
| 2:34.45 | ישיח |  |
| 2:34.46 | בחוקיך | . |
| 2:34.47 | אז |  |
| 2:34.48 | לא |  |
| 2:34.49 | אבוש |  |
| 2:34.50 | כאשר |  |
| 2:34.51 | לא |  |
| 2:34.52 | אסור |  |
| 2:34.53 | מן |  |
| 2:34.54 | הדרך |  |
| 2:34.55 | אשר |  |
| 2:34.56 | למדונו |  |
| 2:34.57 | רז"ל |  |
| 2:34.58 | בעלי |  |
| 2:34.59 | המשנה |  |
| 2:34.60 | והתלמוד |  |
| 2:34.61 | והמדרשים |  |
| 2:34.62 | האמתיים |  |
| 2:34.63 | והפוסקים |  |
| 2:34.64 | ההולכים |  |
| 2:34.65 | בעקבותיהם | . |
| 2:34.66 | אשר |  |
| 2:34.67 | כל |  |
| 2:34.68 | בית |  |
| 2:34.69 | ישראל |  |
| 2:34.70 | נשענים |  |
| 2:34.71 | עליהם | . |
| 2:34.72 | כי |  |
| 2:34.73 | המה |  |
| 2:34.74 | בעלי |  |
| 2:34.75 | הקבלה |  |
| 2:34.76 | האמתית |  |
| 2:34.77 | אשר |  |
| 2:34.78 | קבלו |  |
| 2:34.79 | ממרע"ה |  |
| 2:34.80 | כללותיה |  |
| 2:34.81 | ופרטותיה |  |
| 2:34.82 | של |  |
| 2:34.83 | תורה |  |
| 2:34.84 | אשר |  |
| 2:34.85 | עליהם |  |
| 2:34.86 | עמד |  |
| 2:34.87 | בסיני |  |
| 2:34.88 | ארבעים |  |
| 2:34.89 | יום |  |
| 2:34.90 | וארבעים |  |
| 2:34.91 | לילה |  |
| 2:34.92 | ומסרן |  |
| 2:34.93 | ליהושע |  |
| 2:34.94 | ויהושע |  |
| 2:34.95 | לזקנים |  |
| 2:34.96 | וכו | ' |
| 2:34.97 | כדאיתא |  |
| 2:34.98 | במס | ' |
| 2:34.99 | אבות | . |
| 2:34.100 | א | ) |
| 2:34.101 | וחוב |  |
| 2:34.102 | קדוש |  |
| 2:34.103 | מוטל |  |
| 2:34.104 | עלינו |  |
| 2:34.105 | ללכת |  |
| 2:34.106 | בעקבותיהם |  |
| 2:34.107 | בכל |  |
| 2:34.108 | מידי |  |
| 2:34.109 | דנפקא |  |
| 2:34.110 | מינה |  |
| 2:34.111 | לדינא |  |
| 2:34.112 | בכל |  |
| 2:34.113 | ענייני |  |
| 2:34.114 | איסור |  |
| 2:34.115 | והתר |  |
| 2:34.116 | וטמא |  |
| 2:34.117 | וטהור |  |
| 2:34.118 | ופסול |  |
| 2:34.119 | וכשר | . |
| 2:34.120 | וכ"ש |  |
| 2:34.121 | בענייני |  |
| 2:34.122 | יחוד |  |
| 2:34.123 | השי"ת |  |
| 2:34.124 | שהוא |  |
| 2:34.125 | העיקר |  |
| 2:34.126 | הגדול |  |
| 2:34.127 | ואחריו |  |
| 2:34.128 | קבלת |  |
| 2:34.129 | כל |  |
| 2:34.130 | המצות |  |
| 2:34.131 | וכיצד |  |
| 2:34.132 | עשייתן |  |
| 2:34.133 | כמאמר |  |
| 2:34.134 | התנא |  |
| 2:34.135 | ע"ה |  |
| 2:34.136 | למה |  |
| 2:34.137 | קדמה |  |
| 2:34.138 | פרשת |  |
| 2:34.139 | שמע |  |
| 2:34.140 | לוהיה |  |
| 2:34.141 | אם |  |
| 2:34.142 | שמוע |  |
| 2:34.143 | וכו | '! |
| 2:34.144 | וכל |  |
| 2:34.145 | מחבר |  |
| 2:34.146 | ספר |  |
| 2:34.147 | אשר |  |
| 2:34.148 | יבוא |  |
| 2:34.149 | להוסיף |  |
| 2:34.150 | או |  |
| 2:34.151 | לגרוע |  |
| 2:34.152 | מדבריהם |  |
| 2:34.153 | וכ"ש |  |
| 2:34.154 | להטותינו |  |
| 2:34.155 | מאמונת |  |
| 2:34.156 | תוה"ק |  |
| 2:34.157 | כפי |  |
| 2:34.158 | קבלת |  |
| 2:34.159 | רבותינו |  |
| 2:34.160 | בעלי |  |
| 2:34.161 | המשנה |  |
| 2:34.162 | והתלמוד | . |
| 2:34.163 | לא |  |
| 2:34.164 | נשמע |  |
| 2:34.165 | לו | . |
| 2:34.166 | ואף |  |
| 2:34.167 | אם |  |
| 2:34.168 | יעש |  |
| 2:34.169 | אותות |  |
| 2:34.170 | ומופתים |  |
| 2:34.171 | בשמים |  |
| 2:34.172 | ובארץ |  |
| 2:34.173 | לא |  |
| 2:34.174 | נסור |  |
| 2:34.175 | מאחרי |  |
| 2:34.176 | ה | ' |
| 2:34.177 | אלהינו |  |
| 2:34.178 | על |  |
| 2:34.179 | פי |  |
| 2:34.180 | קבלת |  |
| 2:34.181 | רז"ל |  |
| 2:34.182 | בעלי |  |
| 2:34.183 | המשנה |  |
| 2:34.184 | והתלמוד |  |
| 2:34.185 | והפוסקים |  |
| 2:34.186 | הבאים |  |
| 2:34.187 | אחריהם |  |
| 2:34.188 | ההולכים |  |
| 2:34.189 | בעקבותיהם | , |
| 2:34.190 | וגם |  |
| 2:34.191 | אם |  |
| 2:34.192 | מן |  |
| 2:34.193 | השמים |  |
| 2:34.194 | תקרא |  |
| 2:34.195 | בת |  |
| 2:34.196 | קול |  |
| 2:34.197 | שמעו |  |
| 2:34.198 | אליו |  |
| 2:34.199 | לבור |  |
| 2:34.200 | מקבלתם |  |
| 2:34.201 | לא |  |
| 2:34.202 | משגחינן |  |
| 2:34.203 | בה |  |
| 2:34.204 | כמאמר |  |
| 2:34.205 | ר | ' |
| 2:34.206 | יהושע |  |
| 2:34.207 | אין |  |
| 2:34.208 | משגיחין |  |
| 2:34.209 | בבת |  |
| 2:34.210 | קול |  |
| 2:34.211 | כי |  |
| 2:34.212 | לא |  |
| 2:34.213 | בשמים |  |
| 2:34.214 | היא | . |


### Translation under review (2:34)

English: An answer I ask of my God, that He lead me in paths of straightness and rescue us from a lying lip, from a deceitful tongue. Teach me, LORD, Your ways; I will walk in Your truth. Unite my heart to fear Your name. And send Your light and Your truth; they will lead me to answer a word spoken in its settings. Roll reproach and scorn away from me, for I have kept Your testimonies. Even though princes sit and speak against me, Your servant meditates on Your statutes. Then I shall not be ashamed, when I do not turn aside from the way that our Rabbis of blessed memory taught us, the masters of the Mishnah and the Talmud and the true midrashim and the decisors who follow in their footsteps, on whom all the house of Israel lean. For they are the masters of the true tradition, who received from Moses our teacher the generalities and particulars of the Torah, upon which he stood at Sinai forty days and forty nights, and handed them to Joshua, and Joshua to the elders, etc., as is stated in tractate Avot. 1) And it is a holy duty incumbent on us to follow in their footsteps in everything that yields a practical difference in law: in all matters of forbidden and permitted, impure and pure, unfit and fit. All the more so in matters of the unity of the Name, may He be blessed, which is the great principle, and after it the tradition of all the commandments and how they are to be done, as the Tanna, peace be upon him, said: "Why did the section of Shema precede 'And it shall be, if you hearken, etc.!'" Any author of a book who comes to add to or subtract from their words, and all the more so to turn us from the faith of the holy Torah according to the tradition of our rabbis, the masters of the Mishnah and the Talmud, we will not hearken to him. Even if he performs signs and wonders in heaven and on earth, we will not turn aside from following the LORD our God, according to the tradition of our Sages of blessed memory, the masters of the Mishnah and the Talmud and the decisors who come after them and follow in their footsteps. And even if from heaven a bat kol calls, hearken to him, to depart from their tradition, we shall not heed it, as Rabbi Yehoshua said: "One does not heed a bat kol, for it is not in heaven."

Literal: an-answer I-ask of-my-God lead-me in-the-paths-of straightness and-rescue-us from-the-lip-of lying from-the-tongue-of deceit teach-me LORD Your-ways I-will-walk in-Your-truth unite my-heart to-fear Your-name and-send Your-light and-Your-truth they will-lead-me to-answer a-word spoken in its-settings roll away-from-me reproach and-scorn for Your-testimonies I-have-kept even though they-sit princes against-me spoke Your-servant meditates on-Your-statutes then not I-will-be-ashamed when not I-turn-aside from the-way that taught-us our-Rabbis-of-blessed-memory the-masters-of the-Mishnah and-the-Talmud and-the-midrashim the-true and-the-decisors who-follow in-their-footsteps on-whom all the-house-of Israel lean on-them for they the-masters-of the-tradition true who received from-Moses-our-teacher its-generalities and-its-particulars of the-Torah which upon-them he-stood at-Sinai forty days and-forty nights and-handed-them to-Joshua and-Joshua to-the-elders etc. as-is-stated in-tractate Avot 1 and-a-duty holy is-laid upon-us to-walk in-their-footsteps in-every anything that-yields a-practical-difference in-law in-all matters-of forbidden and-permitted impure and-pure unfit and-fit all-the-more-so in-matters-of unity the-Name,-may-He-be-blessed which-is the-principle great and-after-it the-tradition-of all the-commandments and-how their-doing as-the-saying the-Tanna peace-be-upon-him why preceded the-section Shema to-'and-it-shall-be' if you-hearken etc.! and-any author of-a-book who comes to-add or to-subtract from-their-words and-all-the-more-so to-turn-us from-the-faith-of the-holy-Torah according-to the-tradition-of our-rabbis the-masters-of the-Mishnah and-the-Talmud not we-will-hearken to-him and-even if he-performs signs and-wonders in-heaven and-on-earth we-will-not turn-aside from-following the-LORD our-God according-to by-the-mouth-of the-tradition-of our-Sages-of-blessed-memory the-masters-of the-Mishnah and-the-Talmud and-the-decisors who-come after-them who-follow in-their-footsteps and-even if from heaven calls daughter-of a-voice hearken to-him to-depart from-their-tradition we-do-not heed it as-said Rabbi Yehoshua one-does-not heed daughter-of a-voice for it-is-not in-heaven it

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:34.1 | an answer |  |  |  |
| 2:34.2 | I ask |  |  | the paragraph opens as a prayer: 'an answer I ask of my God' |
| 2:34.3 | of my God |  |  |  |
| 2:34.4 | lead me |  |  |  |
| 2:34.5 | in the paths of |  |  |  |
| 2:34.6 | straightness |  |  |  |
| 2:34.7 | and rescue us |  |  |  |
| 2:34.8 | from the lip of |  |  |  |
| 2:34.9 | lying |  |  |  |
| 2:34.10 | from the tongue of |  |  |  |
| 2:34.11 | deceit |  |  |  |
| 2:34.12 | teach me |  |  |  |
| 2:34.13 | LORD |  |  | the Tetragrammaton as written (ה׳ in Ps 86:11), read in context as the LORD |
| 2:34.14 | Your ways |  |  |  |
| 2:34.15 | I will walk |  |  |  |
| 2:34.16 | in Your truth |  |  |  |
| 2:34.17 | unite |  |  |  |
| 2:34.18 | my heart |  |  |  |
| 2:34.19 | to fear |  |  |  |
| 2:34.20 | Your name |  |  |  |
| 2:34.21 | and send |  |  |  |
| 2:34.22 | Your light |  |  |  |
| 2:34.23 | and Your truth |  |  |  |
| 2:34.24 | they |  |  |  |
| 2:34.25 | will lead me |  |  |  |
| 2:34.26 | to answer |  |  |  |
| 2:34.27 | a word |  |  |  |
| 2:34.28 | spoken |  |  |  |
| 2:34.29 | in |  |  |  |
| 2:34.30 | its settings |  |  | printed אפניו; Prov 25:11 has אופניו (see apparatus) |
| 2:34.31 | roll |  |  |  |
| 2:34.32 | away from me |  |  |  |
| 2:34.33 | reproach |  |  |  |
| 2:34.34 | and scorn |  |  |  |
| 2:34.35 | for |  |  |  |
| 2:34.36 | Your testimonies |  |  |  |
| 2:34.37 | I have kept |  |  |  |
| 2:34.38 | even |  |  |  |
| 2:34.39 | though |  |  |  |
| 2:34.40 | they sit |  |  |  |
| 2:34.41 | princes |  |  |  |
| 2:34.42 | against me |  |  |  |
| 2:34.43 | spoke |  |  |  |
| 2:34.44 | Your servant |  |  |  |
| 2:34.45 | meditates |  |  |  |
| 2:34.46 | on Your statutes |  |  |  |
| 2:34.47 | then |  |  |  |
| 2:34.48 | not |  |  |  |
| 2:34.49 | I will be ashamed |  |  |  |
| 2:34.50 | when |  |  |  |
| 2:34.51 | not |  |  |  |
| 2:34.52 | I turn aside |  |  |  |
| 2:34.53 | from |  |  |  |
| 2:34.54 | the way |  |  |  |
| 2:34.55 | that |  |  |  |
| 2:34.56 | taught us |  |  |  |
| 2:34.57 | our Rabbis of blessed memory |  |  | רז"ל = רבותינו זכרונם לברכה |
| 2:34.58 | the masters of |  |  |  |
| 2:34.59 | the Mishnah |  |  |  |
| 2:34.60 | and the Talmud |  |  |  |
| 2:34.61 | and the midrashim |  |  |  |
| 2:34.62 | the true |  |  |  |
| 2:34.63 | and the decisors |  |  |  |
| 2:34.64 | who follow |  |  |  |
| 2:34.65 | in their footsteps |  |  |  |
| 2:34.66 | on whom |  |  |  |
| 2:34.67 | all |  |  |  |
| 2:34.68 | the house of |  |  |  |
| 2:34.69 | Israel |  |  |  |
| 2:34.70 | lean |  |  |  |
| 2:34.71 | on them |  |  |  |
| 2:34.72 | for |  |  |  |
| 2:34.73 | they |  |  |  |
| 2:34.74 | the masters of |  |  |  |
| 2:34.75 | the tradition |  |  | הקבלה: the word the Kabbalists use for their own lore; the author claims it here for the Mishnah and Talmud |
| 2:34.76 | true |  |  |  |
| 2:34.77 | who |  |  |  |
| 2:34.78 | received |  |  |  |
| 2:34.79 | from Moses our teacher |  |  | ממרע"ה = משה רבנו עליו השלום |
| 2:34.80 | its generalities |  |  |  |
| 2:34.81 | and its particulars |  |  |  |
| 2:34.82 | of |  |  |  |
| 2:34.83 | the Torah |  |  |  |
| 2:34.84 | which |  |  |  |
| 2:34.85 | upon them |  |  |  |
| 2:34.86 | he stood |  |  |  |
| 2:34.87 | at Sinai |  |  |  |
| 2:34.88 | forty |  |  |  |
| 2:34.89 | days |  |  |  |
| 2:34.90 | and forty |  |  |  |
| 2:34.91 | nights |  |  |  |
| 2:34.92 | and handed them |  |  |  |
| 2:34.93 | to Joshua |  |  |  |
| 2:34.94 | and Joshua |  |  |  |
| 2:34.95 | to the elders |  |  |  |
| 2:34.96 | etc. |  |  |  |
| 2:34.97 | as is stated |  |  |  |
| 2:34.98 | in tractate |  |  |  |
| 2:34.99 | Avot |  |  |  |
| 2:34.100 | 1 |  |  | printed marker א) (a section numeral), rendered 1) |
| 2:34.101 | and a duty |  |  |  |
| 2:34.102 | holy |  |  |  |
| 2:34.103 | is laid |  |  |  |
| 2:34.104 | upon us |  |  |  |
| 2:34.105 | to walk |  |  |  |
| 2:34.106 | in their footsteps |  |  |  |
| 2:34.107 | in every |  |  |  |
| 2:34.108 | anything |  |  |  |
| 2:34.109 | that yields |  |  | Aramaic idiom מידי דנפקא מינה לדינא, 'anything with a practical difference in law' |
| 2:34.110 | a practical difference |  |  |  |
| 2:34.111 | in law |  |  |  |
| 2:34.112 | in all |  |  |  |
| 2:34.113 | matters of |  |  |  |
| 2:34.114 | forbidden |  |  |  |
| 2:34.115 | and permitted |  |  |  |
| 2:34.116 | impure |  |  |  |
| 2:34.117 | and pure |  |  |  |
| 2:34.118 | unfit |  |  |  |
| 2:34.119 | and fit |  |  |  |
| 2:34.120 | all the more so |  |  | וכ"ש = וכל שכן, 'all the more so' |
| 2:34.121 | in matters of |  |  |  |
| 2:34.122 | unity |  |  |  |
| 2:34.123 | the Name, may He be blessed |  |  | השי"ת = the Name, may He be blessed |
| 2:34.124 | which is |  |  |  |
| 2:34.125 | the principle |  |  |  |
| 2:34.126 | great |  |  |  |
| 2:34.127 | and after it |  |  |  |
| 2:34.128 | the tradition of |  |  |  |
| 2:34.129 | all |  |  |  |
| 2:34.130 | the commandments |  |  |  |
| 2:34.131 | and how |  |  |  |
| 2:34.132 | their doing |  |  |  |
| 2:34.133 | as the saying |  |  |  |
| 2:34.134 | the Tanna |  |  |  |
| 2:34.135 | peace be upon him |  |  | ע"ה = peace be upon him |
| 2:34.136 | why |  |  |  |
| 2:34.137 | preceded |  |  |  |
| 2:34.138 | the section |  |  |  |
| 2:34.139 | Shema |  |  |  |
| 2:34.140 | to 'and it shall be' |  |  | ל- before והיה: 'to And it shall be', the opening words of the second paragraph of the Shema (Deut 11:13) |
| 2:34.141 | if |  |  |  |
| 2:34.142 | you hearken |  |  |  |
| 2:34.143 | etc.! |  |  | וכו' truncates the quotation; 'etc.' marks it and nothing is added |
| 2:34.144 | and any |  |  |  |
| 2:34.145 | author |  |  |  |
| 2:34.146 | of a book |  |  |  |
| 2:34.147 | who |  |  |  |
| 2:34.148 | comes |  |  |  |
| 2:34.149 | to add |  |  |  |
| 2:34.150 | or |  |  |  |
| 2:34.151 | to subtract |  |  |  |
| 2:34.152 | from their words |  |  |  |
| 2:34.153 | and all the more so |  |  |  |
| 2:34.154 | to turn us |  |  |  |
| 2:34.155 | from the faith of |  |  |  |
| 2:34.156 | the holy Torah |  |  | תוה"ק = תורה הקדושה, 'the holy Torah' |
| 2:34.157 | according to |  |  |  |
| 2:34.158 | the tradition of |  |  |  |
| 2:34.159 | our rabbis |  |  |  |
| 2:34.160 | the masters of |  |  |  |
| 2:34.161 | the Mishnah |  |  |  |
| 2:34.162 | and the Talmud |  |  |  |
| 2:34.163 | not |  |  |  |
| 2:34.164 | we will hearken |  |  |  |
| 2:34.165 | to him |  |  |  |
| 2:34.166 | and even |  |  |  |
| 2:34.167 | if |  |  |  |
| 2:34.168 | he performs |  |  | יעש: a contracted form of יעשה, 'he will do' |
| 2:34.169 | signs |  |  |  |
| 2:34.170 | and wonders |  |  |  |
| 2:34.171 | in heaven |  |  |  |
| 2:34.172 | and on earth |  |  |  |
| 2:34.173 | we will not |  |  |  |
| 2:34.174 | turn aside |  |  |  |
| 2:34.175 | from following |  |  |  |
| 2:34.176 | the LORD |  |  | ה׳ as the Tetragrammaton: the LORD |
| 2:34.177 | our God |  |  |  |
| 2:34.178 | according to |  |  |  |
| 2:34.179 | by the mouth of |  |  |  |
| 2:34.180 | the tradition of |  |  |  |
| 2:34.181 | our Sages of blessed memory |  |  |  |
| 2:34.182 | the masters of |  |  |  |
| 2:34.183 | the Mishnah |  |  |  |
| 2:34.184 | and the Talmud |  |  |  |
| 2:34.185 | and the decisors |  |  |  |
| 2:34.186 | who come |  |  |  |
| 2:34.187 | after them |  |  |  |
| 2:34.188 | who follow |  |  |  |
| 2:34.189 | in their footsteps |  |  |  |
| 2:34.190 | and even |  |  |  |
| 2:34.191 | if |  |  |  |
| 2:34.192 | from |  |  |  |
| 2:34.193 | heaven |  |  |  |
| 2:34.194 | calls |  |  |  |
| 2:34.195 | daughter of |  |  | בת קול, literally 'daughter of a voice' |
| 2:34.196 | a voice |  |  | a heavenly voice (see context) |
| 2:34.197 | hearken |  |  |  |
| 2:34.198 | to him |  |  |  |
| 2:34.199 | to depart |  |  | לבור: 'to choose'; here 'to depart from' |
| 2:34.200 | from their tradition |  |  |  |
| 2:34.201 | we do not |  |  |  |
| 2:34.202 | heed |  |  |  |
| 2:34.203 | it |  |  |  |
| 2:34.204 | as said |  |  |  |
| 2:34.205 | Rabbi |  |  | ר׳ = Rabbi |
| 2:34.206 | Yehoshua |  |  | Rabbi Yehoshua, as the author names him |
| 2:34.207 | one does not |  |  |  |
| 2:34.208 | heed |  |  |  |
| 2:34.209 | daughter of |  |  |  |
| 2:34.210 | a voice |  |  |  |
| 2:34.211 | for |  |  |  |
| 2:34.212 | it is not |  |  |  |
| 2:34.213 | in heaven |  |  |  |
| 2:34.214 | it |  |  |  |

Phrases:
- 2:34.1 2:34.2 2:34.3 | An answer I ask of my God | an-answer I-ask of-my-God; the paragraph opens with the noun, then the verb: an answer asked for, in the first person
- 2:34.4 2:34.5 2:34.6 | that He lead me in paths of straightness | lead-me in-the-paths-of straightness; ינחני is 'He will lead me'; the English makes it a clause of the prayer
- 2:34.7 2:34.8 2:34.9 2:34.10 2:34.11 | and rescue us from a lying lip, from a deceitful tongue | and-rescue-us from-the-lip-of lying from-the-tongue-of deceit; ויצילנו is plural 'us', as written; the psalm (Ps 120:2) has 'my soul'
- 2:34.12 2:34.13 2:34.14 2:34.15 2:34.16 | Teach me, LORD, Your ways; I will walk in Your truth. | teach-me LORD Your-ways I-will-walk in-Your-truth; imperative with the LORD, then first-person 'I will walk' (Ps 86:11)
- 2:34.17 2:34.18 2:34.19 2:34.20 | Unite my heart to fear Your name. | unite my-heart to-fear Your-name; Ps 86:11; יחד is the plain verb 'unite' with the heart as object
- 2:34.21 2:34.22 2:34.23 2:34.24 2:34.25 | And send Your light and Your truth; they will lead me | and-send Your-light and-Your-truth they will-lead-me; imperative ושלח; המה 'they' are light and truth (Ps 43:3)
- 2:34.26 2:34.27 2:34.28 2:34.29 2:34.30 | to answer a word spoken in its settings | to-answer a-word spoken in its-settings; להשיב דבר דבור; על אפניו 'on its settings' (Prov 25:11)
- 2:34.31 2:34.32 2:34.33 2:34.34 | Roll reproach and scorn away from me | roll away-from-me reproach and-scorn; imperative גל, 'roll' (Ps 119:22)
- 2:34.35 2:34.36 2:34.37 | for I have kept Your testimonies | for Your-testimonies I-have-kept
- 2:34.38 2:34.39 2:34.40 2:34.41 2:34.42 2:34.43 | Even though princes sit and speak against me | even though they-sit princes against-me spoke; נדברו בי: 'spoke about me'; the sense in the psalm is hostile
- 2:34.44 2:34.45 2:34.46 | Your servant meditates on Your statutes | Your-servant meditates on-Your-statutes; Ps 119:23
- 2:34.47 2:34.48 2:34.49 2:34.50 2:34.51 2:34.52 | Then I shall not be ashamed, when I do not turn aside | then not I-will-be-ashamed when not I-turn-aside; אז לא אבוש is the main clause; כאשר is 'when'
- 2:34.53 2:34.54 2:34.55 | from the way that | from the-way that; מן הדרך אשר: 'from the way which'
- 2:34.56 2:34.57 | our Rabbis of blessed memory taught us | taught-us our-Rabbis-of-blessed-memory; the Hebrew puts the verb first: 'taught us our Rabbis'
- 2:34.58 2:34.59 2:34.60 2:34.61 2:34.62 2:34.63 2:34.64 2:34.65 | the masters of the Mishnah and the Talmud and the true midrashim and the decisors who follow in their footsteps | the-masters-of the-Mishnah and-the-Talmud and-the-midrashim the-true and-the-decisors who-follow in-their-footsteps
- 2:34.66 2:34.67 2:34.68 2:34.69 2:34.70 2:34.71 | on whom all the house of Israel lean | on-whom all the-house-of Israel lean on-them; נשענים עליהם: 'lean on them'
- 2:34.72 2:34.73 2:34.74 | For they are the masters of | for they the-masters-of; כי המה בעלי: 'for they, the masters of'
- 2:34.75 2:34.76 2:34.77 | the true tradition, who | the-tradition true who; Q5 default: the tradition the author claims for the Mishnah and Talmud
- 2:34.78 2:34.79 | received from Moses our teacher | received from-Moses-our-teacher; ממרע"ה: 'from Moses our teacher'
- 2:34.80 2:34.81 2:34.82 2:34.83 | the generalities and particulars of the Torah | its-generalities and-its-particulars of the-Torah; כללותיה ופרטותיה: 'its generalities and its particulars'
- 2:34.84 2:34.85 2:34.86 2:34.87 | upon which he stood at Sinai | which upon-them he-stood at-Sinai; עליהם עמד: 'stood upon them'
- 2:34.88 2:34.89 2:34.90 2:34.91 | forty days and forty nights | forty days and-forty nights
- 2:34.92 2:34.93 2:34.94 2:34.95 | and handed them to Joshua, and Joshua to the elders | and-handed-them to-Joshua and-Joshua to-the-elders
- 2:34.96 2:34.97 2:34.98 2:34.99 | etc., as is stated in tractate Avot. | etc. as-is-stated in-tractate Avot; וכו' marks the omitted part of the chain; no words are added
- 2:34.100 | 1) | 1; the printed marker א) rendered 1)
- 2:34.101 2:34.102 2:34.103 2:34.104 2:34.105 2:34.106 | And it is a holy duty incumbent on us to follow in their footsteps | and-a-duty holy is-laid upon-us to-walk in-their-footsteps
- 2:34.107 2:34.108 2:34.109 2:34.110 2:34.111 | in everything that yields a practical difference in law | in-every anything that-yields a-practical-difference in-law; Aramaic idiom, translated by its sense
- 2:34.112 2:34.113 2:34.114 2:34.115 2:34.116 2:34.117 2:34.118 2:34.119 | in all matters of forbidden and permitted, impure and pure, unfit and fit | in-all matters-of forbidden and-permitted impure and-pure unfit and-fit
- 2:34.120 2:34.121 | All the more so in matters of | all-the-more-so in-matters-of; וכ"ש: 'all the more so'
- 2:34.122 2:34.123 | the unity of the Name, may He be blessed | unity the-Name,-may-He-be-blessed; יחוד השי"ת: the fixed rendering
- 2:34.124 2:34.125 2:34.126 | which is the great principle | which-is the-principle great
- 2:34.127 2:34.128 2:34.129 2:34.130 | and after it the tradition of all the commandments | and-after-it the-tradition-of all the-commandments
- 2:34.131 2:34.132 | and how they are to be done | and-how their-doing
- 2:34.133 2:34.134 2:34.135 | as the Tanna, peace be upon him | as-the-saying the-Tanna peace-be-upon-him; כמאמר introduces the Tanna's words
- 2:34.136 2:34.137 2:34.138 2:34.139 | Why did the section of Shema precede | why preceded the-section Shema; the question is quoted; למה קדמה: 'why preceded'
- 2:34.140 2:34.141 2:34.142 2:34.143 | And it shall be, if you hearken, etc.! | to-'and-it-shall-be' if you-hearken etc.!; לוהיה: 'to And it shall be'; וכו' truncates the quotation
- 2:34.144 2:34.145 2:34.146 2:34.147 | Any author of a book who | and-any author of-a-book who; וכל מחבר ספר אשר: 'and any author of a book who'
- 2:34.148 2:34.149 2:34.150 2:34.151 | comes to add to or subtract | comes to-add or to-subtract; יבוא להוסיף או לגרוע
- 2:34.152 | from their words | from-their-words
- 2:34.153 2:34.154 2:34.155 2:34.156 | and all the more so to turn us from the faith of the holy Torah | and-all-the-more-so to-turn-us from-the-faith-of the-holy-Torah; וכ"ש להטותינו מאמונת תוה"ק
- 2:34.157 2:34.158 2:34.159 2:34.160 2:34.161 2:34.162 | according to the tradition of our rabbis, the masters of the Mishnah and the Talmud | according-to the-tradition-of our-rabbis the-masters-of the-Mishnah and-the-Talmud; כפי קבלת רבותינו בעלי המשנה והתלמוד
- 2:34.163 2:34.164 2:34.165 | we will not hearken to him. | not we-will-hearken to-him; לא נשמע לו: 'not, we will hearken to him'
- 2:34.166 2:34.167 2:34.168 2:34.169 2:34.170 2:34.171 2:34.172 | Even if he performs signs and wonders in heaven and on earth | and-even if he-performs signs and-wonders in-heaven and-on-earth; ואף אם: 'and even if'
- 2:34.173 2:34.174 2:34.175 2:34.176 2:34.177 | we will not turn aside from following the LORD our God | we-will-not turn-aside from-following the-LORD our-God; לא נסור מאחרי ה' אלהינו
- 2:34.178 2:34.179 2:34.180 2:34.181 | according to the tradition of our Sages of blessed memory | according-to by-the-mouth-of the-tradition-of our-Sages-of-blessed-memory; על פי: 'according to the mouth of'
- 2:34.182 2:34.183 2:34.184 2:34.185 2:34.186 2:34.187 2:34.188 2:34.189 | the masters of the Mishnah and the Talmud and the decisors who come after them and follow in their footsteps | the-masters-of the-Mishnah and-the-Talmud and-the-decisors who-come after-them who-follow in-their-footsteps
- 2:34.190 2:34.191 2:34.192 2:34.193 2:34.194 2:34.195 2:34.196 | And even if from heaven a bat kol calls | and-even if from heaven calls daughter-of a-voice; וגם אם מן השמים תקרא בת קול: 'and even if from heaven a bat kol calls'
- 2:34.197 2:34.198 | hearken to him | hearken to-him; שמעו אליו: 'hear him'; the words of the call
- 2:34.199 2:34.200 | to depart from their tradition | to-depart from-their-tradition; לבור מקבלתם: 'to choose away from their tradition'
- 2:34.201 2:34.202 2:34.203 | we shall not heed it | we-do-not heed it; לא משגחינן בה: 'we do not heed it'
- 2:34.204 2:34.205 2:34.206 | as Rabbi Yehoshua said | as-said Rabbi Yehoshua; כמאמר introduces the quotation
- 2:34.207 2:34.208 2:34.209 2:34.210 | One does not heed a bat kol | one-does-not heed daughter-of a-voice; אין משגיחין בבת קול: 'one does not heed a bat kol'
- 2:34.211 2:34.212 2:34.213 2:34.214 | for it is not in heaven. | for it-is-not in-heaven it; כי לא בשמים היא: 'for not in heaven it'

Grammar notes:
- תשובה אשאלה מאלהי ינחני: the noun תשובה is the object of the first-person 'I ask'; the paragraph is a prayer, the author's own voice, addressed to the LORD.
- The Psalm verses are run into the author's sentences without a formula: הוריני (Ps 86:11), ושלח (Ps 43:3), גל (Ps 119:22) and גם כי ישבו (Ps 119:23), so they take no quotation marks.
- אז לא אבוש כאשר לא אסור מן הדרך: the main clause is 'then I shall not be ashamed'; כאשר is a temporal clause, 'when I do not turn aside'.
- אשר למדונו רז"ל ... בעלי המשנה ... ההולכים בעקבותיהם is a relative clause on הדרך: the way that the Sages taught us, with the authorities in apposition; אשר כל בית ישראל נשענים עליהם is a second relative clause on the same authorities.
- The Tanna's question, 'למה קדמה פרשת שמע לוהיה אם שמוע וכו'!', is quoted in the author's words and truncated with וכו'; the author then argues a fortiori (וכ"ש) from matters of law to the unity of the Name, and from the unity to the rejection of an author who alters the tradition.

Translator's note: The author's prayer opens the paragraph, then his claim for the tradition of the Mishnah, Talmud and decisors. The English keeps 'us' where the Hebrew has ויצילנו; the Psalm allusions are in apparatus.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated (Wikisource additions, Q2).
- apparatus: Ps 120:2 (embedded): משפת שקר, מלשון רמיה. The psalm has 'my soul' where the author writes ויצילנו (us).
- apparatus: Ps 86:11 (embedded): הוריני ה' דרכיך אהלך באמתך; יחד לבבי ליראה שמך.
- apparatus: Ps 43:3 (embedded): ושלח אורך ואמתך המה ינחוני.
- apparatus: Prov 25:11 (embedded): דבר דבור על אופניו. The author prints אפניו; the verse has אופניו.
- apparatus: Ps 119:22: גל מעלי חרפה ובוז כי עדותיך נצרתי.
- apparatus: Ps 119:23: גם כי ישבו שרים בי נדברו, עבדך ישיח בחוקיך.
- apparatus: Ps 119:6 (embedded): אז לא אבוש.
- apparatus: Deut 11:13: והיה אם שמוע תשמעו, the opening of the second paragraph of the Shema; the author writes לוהיה.
- apparatus: Avot (tractate Avot, Mishnah 1:1, the chain Moses, Joshua, the elders): the author gives no chapter, and the words after וכו' are not supplied.
- apparatus: Deut 13:2-4 (embedded allusion): the sign and the wonder, אות ומופת.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

