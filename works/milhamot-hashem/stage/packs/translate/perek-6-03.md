# Translation pack: milhamot-hashem perek-6-03 (Milchamot ha-Shem 6:4–6:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-03`.

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
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
- An initial before a surname that fixes the person (ר"מ ריקנטי): expand and
  name him in a note ("Rabbi Menachem Rikanti"); an initial no surname fixes
  stays an initial with a note, as with R. A. and ר"י חיים ("Rabbi Y. Chayyim").
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


## Milchamot ha-Shem 6:4

Hebrew (Wikisource): ומדברי הר' הסמ"ג הרוחנו שעד ימיו היו כל ישראל תקועים באמונת יחוד השי"ת שהוא אחד ומיוחד מכל שאר האחדים כד' הרס"ג. והיא דעת תורתנו הקדושה שבכתב ושבעל פה שקבלו רז"ל גמרא ובמדרשים הידועים ומפורסמים לרז"ל בלי שום ספק.

| id | word | punctuation after |
|---|---|---|
| 6:4.1 | ומדברי |  |
| 6:4.2 | הר | ' |
| 6:4.3 | הסמ"ג |  |
| 6:4.4 | הרוחנו |  |
| 6:4.5 | שעד |  |
| 6:4.6 | ימיו |  |
| 6:4.7 | היו |  |
| 6:4.8 | כל |  |
| 6:4.9 | ישראל |  |
| 6:4.10 | תקועים |  |
| 6:4.11 | באמונת |  |
| 6:4.12 | יחוד |  |
| 6:4.13 | השי"ת |  |
| 6:4.14 | שהוא |  |
| 6:4.15 | אחד |  |
| 6:4.16 | ומיוחד |  |
| 6:4.17 | מכל |  |
| 6:4.18 | שאר |  |
| 6:4.19 | האחדים |  |
| 6:4.20 | כד | ' |
| 6:4.21 | הרס"ג | . |
| 6:4.22 | והיא |  |
| 6:4.23 | דעת |  |
| 6:4.24 | תורתנו |  |
| 6:4.25 | הקדושה |  |
| 6:4.26 | שבכתב |  |
| 6:4.27 | ושבעל |  |
| 6:4.28 | פה |  |
| 6:4.29 | שקבלו |  |
| 6:4.30 | רז"ל |  |
| 6:4.31 | גמרא |  |
| 6:4.32 | ובמדרשים |  |
| 6:4.33 | הידועים |  |
| 6:4.34 | ומפורסמים |  |
| 6:4.35 | לרז"ל |  |
| 6:4.36 | בלי |  |
| 6:4.37 | שום |  |
| 6:4.38 | ספק | . |


## Milchamot ha-Shem 6:5

Hebrew (Wikisource): ושלא כדעת הקבלה החדשה אשר לא נודע מבטן מי יצאה, ומיד בן נכר הגיע ליד ישראל כמו שכתב הר' חיים יוסף דוד אזולאי בספרו שם הגדולים ויובא לקמן והיא כמעט מינות גמורה, כי רק בלשונם יכזבו לו יתברך שהוא אחד, ובלבם יחשבו כמה עילות משתלשלים זה אחר זה וכל עילה נקרא בשם הקב"ה.

| id | word | punctuation after |
|---|---|---|
| 6:5.1 | ושלא |  |
| 6:5.2 | כדעת |  |
| 6:5.3 | הקבלה |  |
| 6:5.4 | החדשה |  |
| 6:5.5 | אשר |  |
| 6:5.6 | לא |  |
| 6:5.7 | נודע |  |
| 6:5.8 | מבטן |  |
| 6:5.9 | מי |  |
| 6:5.10 | יצאה | , |
| 6:5.11 | ומיד |  |
| 6:5.12 | בן |  |
| 6:5.13 | נכר |  |
| 6:5.14 | הגיע |  |
| 6:5.15 | ליד |  |
| 6:5.16 | ישראל |  |
| 6:5.17 | כמו |  |
| 6:5.18 | שכתב |  |
| 6:5.19 | הר | ' |
| 6:5.20 | חיים |  |
| 6:5.21 | יוסף |  |
| 6:5.22 | דוד |  |
| 6:5.23 | אזולאי |  |
| 6:5.24 | בספרו |  |
| 6:5.25 | שם |  |
| 6:5.26 | הגדולים |  |
| 6:5.27 | ויובא |  |
| 6:5.28 | לקמן |  |
| 6:5.29 | והיא |  |
| 6:5.30 | כמעט |  |
| 6:5.31 | מינות |  |
| 6:5.32 | גמורה | , |
| 6:5.33 | כי |  |
| 6:5.34 | רק |  |
| 6:5.35 | בלשונם |  |
| 6:5.36 | יכזבו |  |
| 6:5.37 | לו |  |
| 6:5.38 | יתברך |  |
| 6:5.39 | שהוא |  |
| 6:5.40 | אחד | , |
| 6:5.41 | ובלבם |  |
| 6:5.42 | יחשבו |  |
| 6:5.43 | כמה |  |
| 6:5.44 | עילות |  |
| 6:5.45 | משתלשלים |  |
| 6:5.46 | זה |  |
| 6:5.47 | אחר |  |
| 6:5.48 | זה |  |
| 6:5.49 | וכל |  |
| 6:5.50 | עילה |  |
| 6:5.51 | נקרא |  |
| 6:5.52 | בשם |  |
| 6:5.53 | הקב"ה | . |


## Milchamot ha-Shem 6:6

Hebrew (Wikisource): סג) והר"מ בעל הגדרים בשורש אמונה כתב וז"ל, אמונה הוא הענין המצוייר בנפש כשיאמינו בו כך כמו שצויר, יהיה אמת מה שצייר או בלתי אמת, כי היא אמונה אצלו רצויה אצל המאמין, והאמונה האמתית היא ההאמנה במה שצייר שהוא חוץ לשכל כפי מה שהו מצויר בשכל, כי ההאמנה אינה ענין הנאמר בפה, כי אין אמונה אלא אחר ציור. ואם יהיה עם זאת האמונה שאי אפשר חילוף זאת האמונה בשום פנים. אז תהיה האמונה ההיא אמתית וקיימת, כי תסכים למציאות והאמות כאשר יהיה בצודק, והוא לא יאמין במקבילו, היה אמות אמרתי ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:6.1 | סג | ) |
| 6:6.2 | והר"מ |  |
| 6:6.3 | בעל |  |
| 6:6.4 | הגדרים |  |
| 6:6.5 | בשורש |  |
| 6:6.6 | אמונה |  |
| 6:6.7 | כתב |  |
| 6:6.8 | וז"ל | , |
| 6:6.9 | אמונה |  |
| 6:6.10 | הוא |  |
| 6:6.11 | הענין |  |
| 6:6.12 | המצוייר |  |
| 6:6.13 | בנפש |  |
| 6:6.14 | כשיאמינו |  |
| 6:6.15 | בו |  |
| 6:6.16 | כך |  |
| 6:6.17 | כמו |  |
| 6:6.18 | שצויר | , |
| 6:6.19 | יהיה |  |
| 6:6.20 | אמת |  |
| 6:6.21 | מה |  |
| 6:6.22 | שצייר |  |
| 6:6.23 | או |  |
| 6:6.24 | בלתי |  |
| 6:6.25 | אמת | , |
| 6:6.26 | כי |  |
| 6:6.27 | היא |  |
| 6:6.28 | אמונה |  |
| 6:6.29 | אצלו |  |
| 6:6.30 | רצויה |  |
| 6:6.31 | אצל |  |
| 6:6.32 | המאמין | , |
| 6:6.33 | והאמונה |  |
| 6:6.34 | האמתית |  |
| 6:6.35 | היא |  |
| 6:6.36 | ההאמנה |  |
| 6:6.37 | במה |  |
| 6:6.38 | שצייר |  |
| 6:6.39 | שהוא |  |
| 6:6.40 | חוץ |  |
| 6:6.41 | לשכל |  |
| 6:6.42 | כפי |  |
| 6:6.43 | מה |  |
| 6:6.44 | שהו |  |
| 6:6.45 | מצויר |  |
| 6:6.46 | בשכל | , |
| 6:6.47 | כי |  |
| 6:6.48 | ההאמנה |  |
| 6:6.49 | אינה |  |
| 6:6.50 | ענין |  |
| 6:6.51 | הנאמר |  |
| 6:6.52 | בפה | , |
| 6:6.53 | כי |  |
| 6:6.54 | אין |  |
| 6:6.55 | אמונה |  |
| 6:6.56 | אלא |  |
| 6:6.57 | אחר |  |
| 6:6.58 | ציור | . |
| 6:6.59 | ואם |  |
| 6:6.60 | יהיה |  |
| 6:6.61 | עם |  |
| 6:6.62 | זאת |  |
| 6:6.63 | האמונה |  |
| 6:6.64 | שאי |  |
| 6:6.65 | אפשר |  |
| 6:6.66 | חילוף |  |
| 6:6.67 | זאת |  |
| 6:6.68 | האמונה |  |
| 6:6.69 | בשום |  |
| 6:6.70 | פנים | . |
| 6:6.71 | אז |  |
| 6:6.72 | תהיה |  |
| 6:6.73 | האמונה |  |
| 6:6.74 | ההיא |  |
| 6:6.75 | אמתית |  |
| 6:6.76 | וקיימת | , |
| 6:6.77 | כי |  |
| 6:6.78 | תסכים |  |
| 6:6.79 | למציאות |  |
| 6:6.80 | והאמות |  |
| 6:6.81 | כאשר |  |
| 6:6.82 | יהיה |  |
| 6:6.83 | בצודק | , |
| 6:6.84 | והוא |  |
| 6:6.85 | לא |  |
| 6:6.86 | יאמין |  |
| 6:6.87 | במקבילו | , |
| 6:6.88 | היה |  |
| 6:6.89 | אמות |  |
| 6:6.90 | אמרתי |  |
| 6:6.91 | ע"כ | . |


## Milchamot ha-Shem 6:7

Hebrew (Wikisource): ונכתבו בצדו מהמפרש וז"ל אמונה הוא הענין המצויר בנפש, לאפוקי אמונת הסותרים שאין לה ציור בנפש, לכן לא תכנס בגדר אמונה. והאמונה האמתית וכו' זה הגדר לקוח מדברי הרב המורה, ושם נאמר שהאמונה האמתית לה שם בשלושה מסכימים בקשר הסכמה אחת, שהם הציור המחשבי, ואמונת הלב, והמציאות. כמו שמצייר במחשבתו צורת הארי ומאמין בלבו שהוא נמצא כן במציאות חוץ לנפש. ואחר שהנפש לא תתרצה באמונתה כי אם בהסכמת השכל, לכן עלינו זאת המצוה לחזק אמונת ודעות התורה במחקר השכלי והסכמת החכמה עם התורה, כי בהצמדם יחד תחזק הנפש באמונתה. והיתה האמת אזור מתנה, והאמונה אזור חלציה. ויחדיו יהיו תמים, בל יזוזו מלבו כל הימים. וכן אמרו רז"ל ושאינו יודע לשאול את פתח לו וכו' וכן הורו הר' בעל החובות וה' בעל בינה לעתים וסמכה אקרא דכ' דרך אמונה בחרתי משפטיך שויתי, פי' האמונה והדעות האמנתי אחר בחירת השכל באמתתן, משפטיך שויתי, אך בדברים הנוגעים במשפט וחוקים שויתי הבלתי מבוררים בשכל עם המבוררים וקצת חכמים בעיניהם יחשבו שהאמונה הנבחרת היא המתנגדת לשכל, ונמשכו בזה אחר דעת בעלי השלישים. ולא כן חלק יעקב, כי אמונתינו לא יסור אותם השכל, ככתוב אל תדמה ישראל אל עמים לא עלו מתוך ההפכה, וריב להם עם הדעת (אסף חלק א' מאמר א') וכן יאמר הכוזרי מאמר ו' סי' ל"א חלילה לאל מן השקר ושיבא בתורה מה שהשכל מרחיק אותו וישימהו שקר, ומי חכם יבין דעת, כי רובי מצות התורה לא ניתנו לדרוש מפי התורה שבכתב כ"א כפי ביאור חכמים ז"ל בתורה שבעל פה.

| id | word | punctuation after |
|---|---|---|
| 6:7.1 | ונכתבו |  |
| 6:7.2 | בצדו |  |
| 6:7.3 | מהמפרש |  |
| 6:7.4 | וז"ל |  |
| 6:7.5 | אמונה |  |
| 6:7.6 | הוא |  |
| 6:7.7 | הענין |  |
| 6:7.8 | המצויר |  |
| 6:7.9 | בנפש | , |
| 6:7.10 | לאפוקי |  |
| 6:7.11 | אמונת |  |
| 6:7.12 | הסותרים |  |
| 6:7.13 | שאין |  |
| 6:7.14 | לה |  |
| 6:7.15 | ציור |  |
| 6:7.16 | בנפש | , |
| 6:7.17 | לכן |  |
| 6:7.18 | לא |  |
| 6:7.19 | תכנס |  |
| 6:7.20 | בגדר |  |
| 6:7.21 | אמונה | . |
| 6:7.22 | והאמונה |  |
| 6:7.23 | האמתית |  |
| 6:7.24 | וכו | ' |
| 6:7.25 | זה |  |
| 6:7.26 | הגדר |  |
| 6:7.27 | לקוח |  |
| 6:7.28 | מדברי |  |
| 6:7.29 | הרב |  |
| 6:7.30 | המורה | , |
| 6:7.31 | ושם |  |
| 6:7.32 | נאמר |  |
| 6:7.33 | שהאמונה |  |
| 6:7.34 | האמתית |  |
| 6:7.35 | לה |  |
| 6:7.36 | שם |  |
| 6:7.37 | בשלושה |  |
| 6:7.38 | מסכימים |  |
| 6:7.39 | בקשר |  |
| 6:7.40 | הסכמה |  |
| 6:7.41 | אחת | , |
| 6:7.42 | שהם |  |
| 6:7.43 | הציור |  |
| 6:7.44 | המחשבי | , |
| 6:7.45 | ואמונת |  |
| 6:7.46 | הלב | , |
| 6:7.47 | והמציאות | . |
| 6:7.48 | כמו |  |
| 6:7.49 | שמצייר |  |
| 6:7.50 | במחשבתו |  |
| 6:7.51 | צורת |  |
| 6:7.52 | הארי |  |
| 6:7.53 | ומאמין |  |
| 6:7.54 | בלבו |  |
| 6:7.55 | שהוא |  |
| 6:7.56 | נמצא |  |
| 6:7.57 | כן |  |
| 6:7.58 | במציאות |  |
| 6:7.59 | חוץ |  |
| 6:7.60 | לנפש | . |
| 6:7.61 | ואחר |  |
| 6:7.62 | שהנפש |  |
| 6:7.63 | לא |  |
| 6:7.64 | תתרצה |  |
| 6:7.65 | באמונתה |  |
| 6:7.66 | כי |  |
| 6:7.67 | אם |  |
| 6:7.68 | בהסכמת |  |
| 6:7.69 | השכל | , |
| 6:7.70 | לכן |  |
| 6:7.71 | עלינו |  |
| 6:7.72 | זאת |  |
| 6:7.73 | המצוה |  |
| 6:7.74 | לחזק |  |
| 6:7.75 | אמונת |  |
| 6:7.76 | ודעות |  |
| 6:7.77 | התורה |  |
| 6:7.78 | במחקר |  |
| 6:7.79 | השכלי |  |
| 6:7.80 | והסכמת |  |
| 6:7.81 | החכמה |  |
| 6:7.82 | עם |  |
| 6:7.83 | התורה | , |
| 6:7.84 | כי |  |
| 6:7.85 | בהצמדם |  |
| 6:7.86 | יחד |  |
| 6:7.87 | תחזק |  |
| 6:7.88 | הנפש |  |
| 6:7.89 | באמונתה | . |
| 6:7.90 | והיתה |  |
| 6:7.91 | האמת |  |
| 6:7.92 | אזור |  |
| 6:7.93 | מתנה | , |
| 6:7.94 | והאמונה |  |
| 6:7.95 | אזור |  |
| 6:7.96 | חלציה | . |
| 6:7.97 | ויחדיו |  |
| 6:7.98 | יהיו |  |
| 6:7.99 | תמים | , |
| 6:7.100 | בל |  |
| 6:7.101 | יזוזו |  |
| 6:7.102 | מלבו |  |
| 6:7.103 | כל |  |
| 6:7.104 | הימים | . |
| 6:7.105 | וכן |  |
| 6:7.106 | אמרו |  |
| 6:7.107 | רז"ל |  |
| 6:7.108 | ושאינו |  |
| 6:7.109 | יודע |  |
| 6:7.110 | לשאול |  |
| 6:7.111 | את |  |
| 6:7.112 | פתח |  |
| 6:7.113 | לו |  |
| 6:7.114 | וכו | ' |
| 6:7.115 | וכן |  |
| 6:7.116 | הורו |  |
| 6:7.117 | הר | ' |
| 6:7.118 | בעל |  |
| 6:7.119 | החובות |  |
| 6:7.120 | וה | ' |
| 6:7.121 | בעל |  |
| 6:7.122 | בינה |  |
| 6:7.123 | לעתים |  |
| 6:7.124 | וסמכה |  |
| 6:7.125 | אקרא |  |
| 6:7.126 | דכ | ' |
| 6:7.127 | דרך |  |
| 6:7.128 | אמונה |  |
| 6:7.129 | בחרתי |  |
| 6:7.130 | משפטיך |  |
| 6:7.131 | שויתי | , |
| 6:7.132 | פי | ' |
| 6:7.133 | האמונה |  |
| 6:7.134 | והדעות |  |
| 6:7.135 | האמנתי |  |
| 6:7.136 | אחר |  |
| 6:7.137 | בחירת |  |
| 6:7.138 | השכל |  |
| 6:7.139 | באמתתן | , |
| 6:7.140 | משפטיך |  |
| 6:7.141 | שויתי | , |
| 6:7.142 | אך |  |
| 6:7.143 | בדברים |  |
| 6:7.144 | הנוגעים |  |
| 6:7.145 | במשפט |  |
| 6:7.146 | וחוקים |  |
| 6:7.147 | שויתי |  |
| 6:7.148 | הבלתי |  |
| 6:7.149 | מבוררים |  |
| 6:7.150 | בשכל |  |
| 6:7.151 | עם |  |
| 6:7.152 | המבוררים |  |
| 6:7.153 | וקצת |  |
| 6:7.154 | חכמים |  |
| 6:7.155 | בעיניהם |  |
| 6:7.156 | יחשבו |  |
| 6:7.157 | שהאמונה |  |
| 6:7.158 | הנבחרת |  |
| 6:7.159 | היא |  |
| 6:7.160 | המתנגדת |  |
| 6:7.161 | לשכל | , |
| 6:7.162 | ונמשכו |  |
| 6:7.163 | בזה |  |
| 6:7.164 | אחר |  |
| 6:7.165 | דעת |  |
| 6:7.166 | בעלי |  |
| 6:7.167 | השלישים | . |
| 6:7.168 | ולא |  |
| 6:7.169 | כן |  |
| 6:7.170 | חלק |  |
| 6:7.171 | יעקב | , |
| 6:7.172 | כי |  |
| 6:7.173 | אמונתינו |  |
| 6:7.174 | לא |  |
| 6:7.175 | יסור |  |
| 6:7.176 | אותם |  |
| 6:7.177 | השכל | , |
| 6:7.178 | ככתוב |  |
| 6:7.179 | אל |  |
| 6:7.180 | תדמה |  |
| 6:7.181 | ישראל |  |
| 6:7.182 | אל |  |
| 6:7.183 | עמים |  |
| 6:7.184 | לא |  |
| 6:7.185 | עלו |  |
| 6:7.186 | מתוך |  |
| 6:7.187 | ההפכה | , |
| 6:7.188 | וריב |  |
| 6:7.189 | להם |  |
| 6:7.190 | עם |  |
| 6:7.191 | הדעת |  |
| 6:7.192 | אסף |  |
| 6:7.193 | חלק |  |
| 6:7.194 | א | ' |
| 6:7.195 | מאמר |  |
| 6:7.196 | א | ') |
| 6:7.197 | וכן |  |
| 6:7.198 | יאמר |  |
| 6:7.199 | הכוזרי |  |
| 6:7.200 | מאמר |  |
| 6:7.201 | ו | ' |
| 6:7.202 | סי | ' |
| 6:7.203 | ל"א |  |
| 6:7.204 | חלילה |  |
| 6:7.205 | לאל |  |
| 6:7.206 | מן |  |
| 6:7.207 | השקר |  |
| 6:7.208 | ושיבא |  |
| 6:7.209 | בתורה |  |
| 6:7.210 | מה |  |
| 6:7.211 | שהשכל |  |
| 6:7.212 | מרחיק |  |
| 6:7.213 | אותו |  |
| 6:7.214 | וישימהו |  |
| 6:7.215 | שקר | , |
| 6:7.216 | ומי |  |
| 6:7.217 | חכם |  |
| 6:7.218 | יבין |  |
| 6:7.219 | דעת | , |
| 6:7.220 | כי |  |
| 6:7.221 | רובי |  |
| 6:7.222 | מצות |  |
| 6:7.223 | התורה |  |
| 6:7.224 | לא |  |
| 6:7.225 | ניתנו |  |
| 6:7.226 | לדרוש |  |
| 6:7.227 | מפי |  |
| 6:7.228 | התורה |  |
| 6:7.229 | שבכתב |  |
| 6:7.230 | כ"א |  |
| 6:7.231 | כפי |  |
| 6:7.232 | ביאור |  |
| 6:7.233 | חכמים |  |
| 6:7.234 | ז"ל |  |
| 6:7.235 | בתורה |  |
| 6:7.236 | שבעל |  |
| 6:7.237 | פה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

