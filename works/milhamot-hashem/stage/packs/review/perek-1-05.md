# Review pack: milhamot-hashem perek-1-05 (Milchamot ha-Shem 1:7–1:12)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-1-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-05`.

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
- **שיתוף / שתוף / לשתף** (there is no שותפות in the text): context.
  - Associating another with God: fixed, "associate / association". "כל
    המשתף שם שמים ודבר אחר נעקר מן העולם" (Sanh. 63a; 1:7, 5:29, 5:32) →
    "whoever associates the name of Heaven with something else is uprooted
    from the world"; "בשיתוף שאר הפרצופים" (5:1) → "in association with the
    other partzufim"; "בלי חיבור וצירוף ושיתוף" (5:21) → "without
    composition, conjunction or association".
  - The Maimonidean term for equivocal words: "בשתוף השם בלבד" (5:10) → "by
    equivocation of the term only"; "מושאלים או משותפים" (5:18) →
    "borrowed or equivocal". Never "association" here. The note gives
    Guide I.
  - צירוף: "conjunction" when used of God; "permutation" in letter mysticism
    ("י"ב צירופי הוי'ה", 2:33 → "the twelve permutations of YHVH").
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
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
- An abbreviation with an apostrophe that no ruling covers (ורבי'): expand it
  by the author's usage, with a note; if two expansions are possible, the
  shorter, and the other in `apparatus`.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
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
- מ"ו = "our teacher" (מורנו) in the formula יודיענו מורנו; ב"ו = "blessed be
  He" only before הוא elsewhere ב"ו = בשר ודם ("flesh and blood").
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
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
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
- A fragment after a Hebrew full stop (ושהם, והפטרה, ולדעת, ורבינו): supply a main
  clause when it continues the verb of the sentence before; join with a comma
  when it continues a list; say which in a token note.
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


## Milchamot ha-Shem 1:7

Hebrew (Wikisource): כל שכן וקל וחומר באלהינו ברוך הוא וברוך שמו שתורה הקפידה שלא לטפול ולשתף עמו שום דבר מכל הנבראים בין גשמי בין רוחני כמו שנאמר זובח לאלהים יחרם בלתי לה' לבדו, ואמר רשב״י התנא ע״ה כל המשתף שם שמים ודבר אחר נעקר מן העולם, ולא אמרינן כלא חד כיון שהתורה ביארה בהדיא בכמה מקומות שהוא אחד מיוחד מכל שאר האחדים הנה כן המזכירים את ה' אל דמי לכם מלהודיע אחדותו ית' לבניכם ולתלמידיכם כמצווה עליכם.

| id | word | punctuation after |
|---|---|---|
| 1:7.1 | כל |  |
| 1:7.2 | שכן |  |
| 1:7.3 | וקל |  |
| 1:7.4 | וחומר |  |
| 1:7.5 | באלהינו |  |
| 1:7.6 | ברוך |  |
| 1:7.7 | הוא |  |
| 1:7.8 | וברוך |  |
| 1:7.9 | שמו |  |
| 1:7.10 | שתורה |  |
| 1:7.11 | הקפידה |  |
| 1:7.12 | שלא |  |
| 1:7.13 | לטפול |  |
| 1:7.14 | ולשתף |  |
| 1:7.15 | עמו |  |
| 1:7.16 | שום |  |
| 1:7.17 | דבר |  |
| 1:7.18 | מכל |  |
| 1:7.19 | הנבראים |  |
| 1:7.20 | בין |  |
| 1:7.21 | גשמי |  |
| 1:7.22 | בין |  |
| 1:7.23 | רוחני |  |
| 1:7.24 | כמו |  |
| 1:7.25 | שנאמר |  |
| 1:7.26 | זובח |  |
| 1:7.27 | לאלהים |  |
| 1:7.28 | יחרם |  |
| 1:7.29 | בלתי |  |
| 1:7.30 | לה | ' |
| 1:7.31 | לבדו | , |
| 1:7.32 | ואמר |  |
| 1:7.33 | רשב״י |  |
| 1:7.34 | התנא |  |
| 1:7.35 | ע״ה |  |
| 1:7.36 | כל |  |
| 1:7.37 | המשתף |  |
| 1:7.38 | שם |  |
| 1:7.39 | שמים |  |
| 1:7.40 | ודבר |  |
| 1:7.41 | אחר |  |
| 1:7.42 | נעקר |  |
| 1:7.43 | מן |  |
| 1:7.44 | העולם | , |
| 1:7.45 | ולא |  |
| 1:7.46 | אמרינן |  |
| 1:7.47 | כלא |  |
| 1:7.48 | חד |  |
| 1:7.49 | כיון |  |
| 1:7.50 | שהתורה |  |
| 1:7.51 | ביארה |  |
| 1:7.52 | בהדיא |  |
| 1:7.53 | בכמה |  |
| 1:7.54 | מקומות |  |
| 1:7.55 | שהוא |  |
| 1:7.56 | אחד |  |
| 1:7.57 | מיוחד |  |
| 1:7.58 | מכל |  |
| 1:7.59 | שאר |  |
| 1:7.60 | האחדים |  |
| 1:7.61 | הנה |  |
| 1:7.62 | כן |  |
| 1:7.63 | המזכירים |  |
| 1:7.64 | את |  |
| 1:7.65 | ה | ' |
| 1:7.66 | אל |  |
| 1:7.67 | דמי |  |
| 1:7.68 | לכם |  |
| 1:7.69 | מלהודיע |  |
| 1:7.70 | אחדותו |  |
| 1:7.71 | ית | ' |
| 1:7.72 | לבניכם |  |
| 1:7.73 | ולתלמידיכם |  |
| 1:7.74 | כמצווה |  |
| 1:7.75 | עליכם | . |


### Translation under review (1:7)

English: All the more so, a fortiori, regarding our God, blessed be He, and blessed be His name, that the Torah insisted that one should not attach to Him or associate with Him anything of all created things, whether corporeal or spiritual, as it is said: "He who sacrifices to the gods shall be destroyed, save to the LORD alone." And Rabbi Shimon bar Yochai the Tanna, peace be upon him, said: whoever associates the name of Heaven with something else is uprooted from the world. and we should not say that He is one as the others are one, since the Torah has made clear plainly in many places that He is one, unique, apart from all the other ones. Thus, you who mention the LORD should not think it right to refrain from making known His oneness, may He be blessed, to your children and your disciples, as you are commanded.

Literal: all the-more-so and-light-and-heavy in-our-God blessed He and-blessed His-name that-the-Torah insisted that-not to-attach and-to-associate with-Him any thing of-all the-created-beings whether corporeal or spiritual as it-is-said he-who-sacrifices to-gods shall-be-devoted-to-destruction except to-the-LORD alone and-he-said Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him everyone-who associates the-name the-name-of-Heaven with-something else is-uprooted from the-world and-not we-say as-one one since the-Torah has-made-clear plainly in-many places that-He one unique from-all the-other the-ones so thus you-who-mention (object) the-LORD let-it-not seem to-you from-making-known His-oneness may-He-be-blessed to-your-children and-to-your-disciples as-is-commanded upon-you

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:7.1 | all |  |  | The two Talmudic forms of the a fortiori argument: כל שכן ('all the more so') and וקל וחומר ('light and heavy'). |
| 1:7.2 | the more so |  |  |  |
| 1:7.3 | and light |  |  |  |
| 1:7.4 | and heavy |  |  |  |
| 1:7.5 | in our God |  |  |  |
| 1:7.6 | blessed |  |  |  |
| 1:7.7 | He |  |  |  |
| 1:7.8 | and blessed |  |  |  |
| 1:7.9 | His name |  |  |  |
| 1:7.10 | that the Torah |  |  |  |
| 1:7.11 | insisted |  |  |  |
| 1:7.12 | that not |  |  |  |
| 1:7.13 | to attach |  |  | The verb is paired with לשתף below; the English renders it as attaching one thing to another. |
| 1:7.14 | and to associate |  |  |  |
| 1:7.15 | with Him |  |  |  |
| 1:7.16 | any |  |  |  |
| 1:7.17 | thing |  |  |  |
| 1:7.18 | of all |  |  |  |
| 1:7.19 | the created beings |  |  |  |
| 1:7.20 | whether |  |  |  |
| 1:7.21 | corporeal |  |  |  |
| 1:7.22 | or |  |  |  |
| 1:7.23 | spiritual |  |  |  |
| 1:7.24 | as |  |  |  |
| 1:7.25 | it is said |  |  |  |
| 1:7.26 | he who sacrifices |  |  |  |
| 1:7.27 | to gods |  |  |  |
| 1:7.28 | shall be devoted to destruction |  |  |  |
| 1:7.29 | except |  |  |  |
| 1:7.30 | to the LORD |  |  |  |
| 1:7.31 | alone |  |  |  |
| 1:7.32 | and he said |  |  |  |
| 1:7.33 | Rabbi Shimon bar Yochai |  |  |  |
| 1:7.34 | the Tanna |  |  |  |
| 1:7.35 | peace be upon him |  |  |  |
| 1:7.36 | everyone who |  |  |  |
| 1:7.37 | associates |  |  |  |
| 1:7.38 | the name |  |  |  |
| 1:7.39 | of Heaven |  |  |  |
| 1:7.40 | with something |  |  |  |
| 1:7.41 | else |  |  |  |
| 1:7.42 | is uprooted |  |  |  |
| 1:7.43 | from |  |  |  |
| 1:7.44 | the world |  |  |  |
| 1:7.45 | and not |  |  |  |
| 1:7.46 | we say |  |  |  |
| 1:7.47 | as one |  |  | כלא חד is an Aramaic-style idiom, literally 'as one'; rendered by its sense in the context. |
| 1:7.48 | one |  |  |  |
| 1:7.49 | since |  |  |  |
| 1:7.50 | the Torah |  |  |  |
| 1:7.51 | has made clear |  |  |  |
| 1:7.52 | plainly |  |  |  |
| 1:7.53 | in many |  |  |  |
| 1:7.54 | places |  |  |  |
| 1:7.55 | that He |  |  |  |
| 1:7.56 | one |  |  |  |
| 1:7.57 | unique |  |  |  |
| 1:7.58 | from all |  |  |  |
| 1:7.59 | other |  |  |  |
| 1:7.60 | the ones |  |  |  |
| 1:7.61 | so |  |  |  |
| 1:7.62 | thus |  |  |  |
| 1:7.63 | you who mention |  |  |  |
| 1:7.64 | (object marker) |  |  |  |
| 1:7.65 | the LORD |  |  | The Tetragrammaton as written, ה׳: 'the LORD'. |
| 1:7.66 | let it not |  |  |  |
| 1:7.67 | seem |  |  |  |
| 1:7.68 | to you |  |  |  |
| 1:7.69 | from making known |  |  | להודיע: to make known. The clause אל דמי לכם מלהודיע is 'let it not seem right to you to refrain from making known'. |
| 1:7.70 | His oneness |  |  |  |
| 1:7.71 | may He be blessed |  |  |  |
| 1:7.72 | to your children |  |  |  |
| 1:7.73 | and to your disciples |  |  |  |
| 1:7.74 | as is commanded |  |  |  |
| 1:7.75 | upon you |  |  |  |

Phrases:
- 1:7.1 1:7.2 1:7.3 1:7.4 | All the more so, a fortiori, | all the-more-so and-light-and-heavy
- 1:7.5 1:7.6 1:7.7 | regarding our God, blessed be He, | in-our-God blessed He
- 1:7.8 1:7.9 | and blessed be His name, | and-blessed His-name
- 1:7.10 1:7.11 1:7.12 1:7.13 1:7.14 1:7.15 | that the Torah insisted that one should not attach to Him or associate with Him | that-the-Torah insisted that-not to-attach and-to-associate with-Him
- 1:7.16 1:7.17 1:7.18 1:7.19 | anything of all created things, | any thing of-all the-created-beings
- 1:7.20 1:7.21 1:7.22 1:7.23 | whether corporeal or spiritual, | whether corporeal or spiritual
- 1:7.24 1:7.25 1:7.26 1:7.27 1:7.28 1:7.29 1:7.30 1:7.31 | as it is said: "He who sacrifices to the gods shall be destroyed, save to the LORD alone." | as it-is-said he-who-sacrifices to-gods shall-be-devoted-to-destruction except to-the-LORD alone
- 1:7.32 1:7.33 1:7.34 1:7.35 | And Rabbi Shimon bar Yochai the Tanna, peace be upon him, said: | and-he-said Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him
- 1:7.36 1:7.37 1:7.38 1:7.39 1:7.40 1:7.41 1:7.42 1:7.43 1:7.44 | whoever associates the name of Heaven with something else is uprooted from the world. | everyone-who associates the-name the-name-of-Heaven with-something else is-uprooted from the-world
- 1:7.45 1:7.46 1:7.47 1:7.48 | and we should not say that He is one as the others are one, | and-not we-say as-one one
- 1:7.49 1:7.50 1:7.51 1:7.52 1:7.53 1:7.54 1:7.55 1:7.56 1:7.57 1:7.58 1:7.59 1:7.60 | since the Torah has made clear plainly in many places that He is one, unique, apart from all the other ones. | since the-Torah has-made-clear plainly in-many places that-He one unique from-all the-other the-ones
- 1:7.61 1:7.62 1:7.63 1:7.64 1:7.65 1:7.66 1:7.67 1:7.68 | Thus, you who mention the LORD should not think it right | so thus you-who-mention (object) the-LORD let-it-not seem to-you
- 1:7.69 1:7.70 1:7.71 | to refrain from making known His oneness, may He be blessed, | from-making-known His-oneness may-He-be-blessed
- 1:7.72 1:7.73 | to your children and your disciples, | to-your-children and-to-your-disciples
- 1:7.74 1:7.75 | as you are commanded. | as-is-commanded upon-you

Grammar notes:
- כל שכן and וקל וחומר are the two Talmudic forms of the a fortiori argument, and the English names the argument once, at the start.
- The שלא clause depends on הקפידה: the Torah insisted that one not attach to God or associate with Him anything created.
- Exod 22:19 is quoted with שנאמר, an explicit quotation, so it takes double quotation marks; the verse itself is in the apparatus.
- ולא אמרינן כלא חד כיון ש… is the reasoning: since the Torah has made clear that He is one, we may not speak of Him as one of the ones.
- הנה כן המזכירים את ה׳ is an address in the second person to the reader; אל דמי לכם is 'it should not seem to you'.

Translator's note: Rendered by sense: כלא חד is the one clause whose sense is taken from its context; the literal is 'as one'.

Commentary rulings:
- apparatus: Exod 22:19, quoted with שנאמר: זובח לאלהים יחרם בלתי לה׳ לבדו. The author writes לה׳ where the verse has ליהוה; the English gives 'the LORD'.

## Milchamot ha-Shem 1:8

Hebrew (Wikisource): ולפי שתחלת החבור הזה נעשה בראשונה על ידי שאלת תלמיד אחד את מי נעבוד על כן הנחתי התחלתו שאלת התלמיד ותשובתו מאחד הלומדים המתעקש פעם ראשונה ושנייה ושלישית.

| id | word | punctuation after |
|---|---|---|
| 1:8.1 | ולפי |  |
| 1:8.2 | שתחלת |  |
| 1:8.3 | החבור |  |
| 1:8.4 | הזה |  |
| 1:8.5 | נעשה |  |
| 1:8.6 | בראשונה |  |
| 1:8.7 | על |  |
| 1:8.8 | ידי |  |
| 1:8.9 | שאלת |  |
| 1:8.10 | תלמיד |  |
| 1:8.11 | אחד |  |
| 1:8.12 | את |  |
| 1:8.13 | מי |  |
| 1:8.14 | נעבוד |  |
| 1:8.15 | על |  |
| 1:8.16 | כן |  |
| 1:8.17 | הנחתי |  |
| 1:8.18 | התחלתו |  |
| 1:8.19 | שאלת |  |
| 1:8.20 | התלמיד |  |
| 1:8.21 | ותשובתו |  |
| 1:8.22 | מאחד |  |
| 1:8.23 | הלומדים |  |
| 1:8.24 | המתעקש |  |
| 1:8.25 | פעם |  |
| 1:8.26 | ראשונה |  |
| 1:8.27 | ושנייה |  |
| 1:8.28 | ושלישית | . |


### Translation under review (1:8)

English: Because the beginning of this composition was made at first by the question of one student: whom shall we worship? Therefore I placed at its opening the question of the student, and its reply from one of the learners who persisted the first, second and third time.

Literal: and-since that-the-beginning-of this-composition this was-made at-first by-way-of the-hand-of the-question-of a-student one (object) whom we-shall-worship on-so I-placed its-beginning the-question-of the-student and-its-reply from-one of-the-learners who-persisted the-time first and-second and-third

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:8.1 | and since |  |  |  |
| 1:8.2 | that the beginning of |  |  |  |
| 1:8.3 | the composition |  |  |  |
| 1:8.4 | this |  |  |  |
| 1:8.5 | was made |  |  |  |
| 1:8.6 | at first |  |  |  |
| 1:8.7 | by way of |  |  |  |
| 1:8.8 | the hand of |  |  |  |
| 1:8.9 | the question of |  |  |  |
| 1:8.10 | a student |  |  |  |
| 1:8.11 | one |  |  |  |
| 1:8.12 | (object marker) |  |  |  |
| 1:8.13 | whom |  |  | The student's question in direct speech: the Hebrew has no question mark, and the English adds one. |
| 1:8.14 | we shall worship |  |  | את מי נעבוד is the question itself, 'whom shall we worship?' |
| 1:8.15 | therefore |  |  | על כן, two words: 'therefore'. |
| 1:8.16 | so |  |  |  |
| 1:8.17 | I placed |  |  |  |
| 1:8.18 | its beginning |  |  |  |
| 1:8.19 | the question of |  |  |  |
| 1:8.20 | the student |  |  |  |
| 1:8.21 | and its reply |  |  |  |
| 1:8.22 | from one |  |  |  |
| 1:8.23 | of the learners |  |  |  |
| 1:8.24 | who persisted |  |  |  |
| 1:8.25 | the time |  |  |  |
| 1:8.26 | first |  |  |  |
| 1:8.27 | and second |  |  |  |
| 1:8.28 | and third |  |  |  |

Phrases:
- 1:8.1 1:8.2 1:8.3 1:8.4 1:8.5 1:8.6 | Because the beginning of this composition was made at first | and-since that-the-beginning-of this-composition this was-made at-first
- 1:8.7 1:8.8 1:8.9 1:8.10 1:8.11 | by the question of one student: | by-way-of the-hand-of the-question-of a-student one
- 1:8.12 1:8.13 1:8.14 | whom shall we worship? | (object) whom we-shall-worship
- 1:8.15 1:8.16 | Therefore | on-so
- 1:8.17 1:8.18 | I placed at its opening | I-placed its-beginning
- 1:8.19 1:8.20 | the question of the student, | the-question-of the-student
- 1:8.21 1:8.22 1:8.23 1:8.24 | and its reply from one of the learners who persisted | and-its-reply from-one of-the-learners who-persisted
- 1:8.25 1:8.26 1:8.27 1:8.28 | the first, second and third time. | the-time first and-second and-third

Grammar notes:
- ולפי שתחלת החבור הזה נעשה… על ידי שאלת תלמיד אחד: the clause gives the reason for the book's form, a dialogue opened by a student's question.
- The student's question את מי נעבוד is a direct question inside the author's sentence; it is kept as a question.
- על כן הנחתי התחלתו: 'therefore I placed its beginning', in the author's first person.
- פעם ראשונה ושנייה ושלישית is an adverbial string: 'the first, second and third time'.

Translator's note: The student's question is kept as a question; the English places the question mark, which the Hebrew does not write.

Commentary rulings:

## Milchamot ha-Shem 1:9

Hebrew (Wikisource): הצעיר יחיה בן סלי' אלקאפח יצ״ו

| id | word | punctuation after |
|---|---|---|
| 1:9.1 | הצעיר |  |
| 1:9.2 | יחיה |  |
| 1:9.3 | בן |  |
| 1:9.4 | סלי | ' |
| 1:9.5 | אלקאפח |  |
| 1:9.6 | יצ״ו |  |


### Translation under review (1:9)

English: The humble Yichya, son of Sali' al-Kafich, may his Rock keep him and give him life.

Literal: the-humble Yichya son-of Sali' al-Kafich may-his-Rock-keep-him-and-give-him-life

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:9.1 | the humble |  |  |  |
| 1:9.2 | Yichya |  |  |  |
| 1:9.3 | son of |  |  |  |
| 1:9.4 | Sali' |  |  | The patronymic is abbreviated in the text and is kept as written. |
| 1:9.5 | al-Kafich |  |  |  |
| 1:9.6 | may his Rock keep him and give him life |  |  | An abbreviated blessing: the letters stand for words of blessing for the author. |

Phrases:
- 1:9.1 1:9.2 | The humble Yichya, | the-humble Yichya
- 1:9.3 1:9.4 1:9.5 | son of Sali' al-Kafich, | son-of Sali' al-Kafich
- 1:9.6 | may his Rock keep him and give him life. | may-his-Rock-keep-him-and-give-him-life

Grammar notes:
- הצעיר is an adjective in apposition to יחיה: the author's self-description, 'the humble'.
- בן סלי׳ אלקאפח is the patronymic and the surname; the abbreviated form is kept as written.
- יצ״ו is an abbreviated blessing in the optative: 'may his Rock keep him and give him life'.

Translator's note: The signature as in the decisions for this work.

Commentary rulings:

## Milchamot ha-Shem 1:10

Hebrew (Wikisource): הקדמת המו"ל.

| id | word | punctuation after |
|---|---|---|
| 1:10.1 | הקדמת |  |
| 1:10.2 | המו"ל | . |


### Translation under review (1:10)

English: The Publisher's Introduction.

Literal: the-introduction of-the-publisher

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:10.1 | the introduction |  |  |  |
| 1:10.2 | the publisher |  |  | המו״ל: the publisher. The voice of the publisher runs through 1:10 to 1:13. |

Phrases:
- 1:10.1 1:10.2 | The Publisher's Introduction. | the-introduction of-the-publisher

Grammar notes:
- הקדמת המו״ל is a construct chain: 'the introduction of the publisher'.
- The heading is in the publisher's voice, not the author's (see the decisions).

Commentary rulings:

## Milchamot ha-Shem 1:11

Hebrew (Wikisource): ספר מלחמת ה' המודיע דרך אמתות יחודו יתברך עפ"י התורה ומפרשיה רז"ל במשנה ובתלמוד בבלי וירושלמי ומדרשי רז"ל וכדעת הגאונים רס"ג ורבינו בחיי בחובת הלבבות ור' יהודה הלוי בשיריו ובכוזרי הרמב"ם וסמ"ג ובעל העיקרים ז"ל וכמקובל אצל כל איש ישראל ובו אזהרה חמורה שלא לעבוד שום נברא כפי קבלת כל נביאינו וכל חכמינו התנאים והאמוראים והגאונים והפוסקים הראשונים מפי משה ואהרן וזקנים לא לזוז מהם ולא להוסיף עליהם.

| id | word | punctuation after |
|---|---|---|
| 1:11.1 | ספר |  |
| 1:11.2 | מלחמת |  |
| 1:11.3 | ה | ' |
| 1:11.4 | המודיע |  |
| 1:11.5 | דרך |  |
| 1:11.6 | אמתות |  |
| 1:11.7 | יחודו |  |
| 1:11.8 | יתברך |  |
| 1:11.9 | עפ"י |  |
| 1:11.10 | התורה |  |
| 1:11.11 | ומפרשיה |  |
| 1:11.12 | רז"ל |  |
| 1:11.13 | במשנה |  |
| 1:11.14 | ובתלמוד |  |
| 1:11.15 | בבלי |  |
| 1:11.16 | וירושלמי |  |
| 1:11.17 | ומדרשי |  |
| 1:11.18 | רז"ל |  |
| 1:11.19 | וכדעת |  |
| 1:11.20 | הגאונים |  |
| 1:11.21 | רס"ג |  |
| 1:11.22 | ורבינו |  |
| 1:11.23 | בחיי |  |
| 1:11.24 | בחובת |  |
| 1:11.25 | הלבבות |  |
| 1:11.26 | ור | ' |
| 1:11.27 | יהודה |  |
| 1:11.28 | הלוי |  |
| 1:11.29 | בשיריו |  |
| 1:11.30 | ובכוזרי |  |
| 1:11.31 | הרמב"ם |  |
| 1:11.32 | וסמ"ג |  |
| 1:11.33 | ובעל |  |
| 1:11.34 | העיקרים |  |
| 1:11.35 | ז"ל |  |
| 1:11.36 | וכמקובל |  |
| 1:11.37 | אצל |  |
| 1:11.38 | כל |  |
| 1:11.39 | איש |  |
| 1:11.40 | ישראל |  |
| 1:11.41 | ובו |  |
| 1:11.42 | אזהרה |  |
| 1:11.43 | חמורה |  |
| 1:11.44 | שלא |  |
| 1:11.45 | לעבוד |  |
| 1:11.46 | שום |  |
| 1:11.47 | נברא |  |
| 1:11.48 | כפי |  |
| 1:11.49 | קבלת |  |
| 1:11.50 | כל |  |
| 1:11.51 | נביאינו |  |
| 1:11.52 | וכל |  |
| 1:11.53 | חכמינו |  |
| 1:11.54 | התנאים |  |
| 1:11.55 | והאמוראים |  |
| 1:11.56 | והגאונים |  |
| 1:11.57 | והפוסקים |  |
| 1:11.58 | הראשונים |  |
| 1:11.59 | מפי |  |
| 1:11.60 | משה |  |
| 1:11.61 | ואהרן |  |
| 1:11.62 | וזקנים |  |
| 1:11.63 | לא |  |
| 1:11.64 | לזוז |  |
| 1:11.65 | מהם |  |
| 1:11.66 | ולא |  |
| 1:11.67 | להוסיף |  |
| 1:11.68 | עליהם | . |


### Translation under review (1:11)

English: This book, The Wars of the LORD, makes known the way of the truths of His unity, may He be blessed, according to the Torah and its interpreters, our Rabbis of blessed memory, in the Mishnah, in the Bavli and the Yerushalmi, and the Midrashim of our Rabbis of blessed memory, and in accordance with the view of the Geonim, Rabbi Saadia Gaon, and our Rabbi Bahya in Chovot ha-Levavot, and Rabbi Yehuda Halevi in his poems, and in the Kuzari, the Rambam, the SeMaG and the author of the Ikkarim, of blessed memory, and as is accepted by every Jew. In it is a severe warning not to worship any created thing, in accordance with the tradition of all our prophets and all our sages, the Tannaim, the Amoraim, the Geonim and the early decisors, who received it from the mouth of Moses, Aaron and the elders, not to deviate from them and not to add to them.

Literal: the-book the-wars-of the-LORD that-makes-known the-way-of the-truths-of His-unity may-He-be-blessed according-to the-Torah and-its-interpreters our-Rabbis-of-blessed-memory in-the-Mishnah and-in-the-Talmud Bavli and-Yerushalmi and-the-Midrashim-of our-Rabbis-of-blessed-memory and-in-accordance-with the-Geonim Rabbi-Saadia-Gaon and-our-Rabbi Bahya in-Chovot-ha-Levavot and-Rabbi Yehuda Halevi in-his-poems and-in-the-Kuzari the-Rambam and-the-SeMaG and-the-author-of the-Ikkarim of-blessed-memory and-as-is-accepted by every man-of Israel and-in-it a-warning severe that-not to-worship any created-thing in-accordance-with the-tradition-of all our-prophets and-all our-sages the-Tannaim and-the-Amoraim and-the-Geonim and-the-decisors the-early from-the-mouth-of Moses and-Aaron and-the-elders not to-deviate from-them and-not to-add to-them

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:11.1 | the book |  |  | The title: The Wars of the LORD (Num 21:14), as the title page gives it. |
| 1:11.2 | the wars of |  |  |  |
| 1:11.3 | the LORD |  |  |  |
| 1:11.4 | that makes known |  |  |  |
| 1:11.5 | the way of |  |  |  |
| 1:11.6 | the truths of |  |  |  |
| 1:11.7 | His unity |  |  |  |
| 1:11.8 | may He be blessed |  |  |  |
| 1:11.9 | according to |  |  |  |
| 1:11.10 | the Torah |  |  |  |
| 1:11.11 | and its interpreters |  |  |  |
| 1:11.12 | our Rabbis of blessed memory |  |  |  |
| 1:11.13 | in the Mishnah |  |  |  |
| 1:11.14 | and in the Talmud |  |  |  |
| 1:11.15 | Bavli |  |  | Bavli: the Babylonian Talmud. |
| 1:11.16 | and the Yerushalmi |  |  | Yerushalmi: the Jerusalem Talmud. |
| 1:11.17 | and the Midrashim of |  |  |  |
| 1:11.18 | our Rabbis of blessed memory |  |  |  |
| 1:11.19 | and in accordance with |  |  |  |
| 1:11.20 | the Geonim |  |  |  |
| 1:11.21 | Rabbi Saadia Gaon |  |  |  |
| 1:11.22 | and our Rabbi |  |  |  |
| 1:11.23 | Bahya |  |  |  |
| 1:11.24 | in Chovot |  |  |  |
| 1:11.25 | ha-Levavot |  |  | Chovot ha-Levavot, 'Duties of the Hearts', by Bahya. |
| 1:11.26 | and Rabbi |  |  |  |
| 1:11.27 | Yehuda |  |  |  |
| 1:11.28 | Halevi |  |  | Yehuda Halevi, poet and author of the Kuzari. |
| 1:11.29 | in his poems |  |  |  |
| 1:11.30 | and in the Kuzari |  |  | The Kuzari, Halevi's dialogue set at the court of the Khazar king. |
| 1:11.31 | the Rambam |  |  |  |
| 1:11.32 | and the SeMaG |  |  | SeMaG: an abbreviation for the medieval code Sefer Mitzvot Gadol. |
| 1:11.33 | and the author of |  |  |  |
| 1:11.34 | the Ikkarim |  |  | Ikkarim: 'Principles', the book of the principles of faith. |
| 1:11.35 | of blessed memory |  |  |  |
| 1:11.36 | and as is accepted |  |  |  |
| 1:11.37 | by |  |  |  |
| 1:11.38 | every |  |  |  |
| 1:11.39 | man |  |  |  |
| 1:11.40 | Israel |  |  |  |
| 1:11.41 | and in it |  |  |  |
| 1:11.42 | a warning |  |  |  |
| 1:11.43 | severe |  |  |  |
| 1:11.44 | that not |  |  |  |
| 1:11.45 | to worship |  |  |  |
| 1:11.46 | any |  |  |  |
| 1:11.47 | created thing |  |  |  |
| 1:11.48 | in accordance with |  |  |  |
| 1:11.49 | the tradition of |  |  | The Hebrew is the word for Kabbalah. The author claims it here, for the chain of tradition from the prophets and sages, not for the Kabbalists. |
| 1:11.50 | all |  |  |  |
| 1:11.51 | our prophets |  |  |  |
| 1:11.52 | and all |  |  |  |
| 1:11.53 | our sages |  |  |  |
| 1:11.54 | the Tannaim |  |  |  |
| 1:11.55 | the Amoraim |  |  |  |
| 1:11.56 | and the Geonim |  |  |  |
| 1:11.57 | and the decisors |  |  |  |
| 1:11.58 | the early |  |  |  |
| 1:11.59 | from the mouth of |  |  | From the mouth of: the chain Moses, Aaron and the elders. |
| 1:11.60 | Moses |  |  |  |
| 1:11.61 | and Aaron |  |  |  |
| 1:11.62 | and the elders |  |  |  |
| 1:11.63 | not |  |  |  |
| 1:11.64 | to deviate |  |  |  |
| 1:11.65 | from them |  |  |  |
| 1:11.66 | and not |  |  |  |
| 1:11.67 | to add |  |  |  |
| 1:11.68 | to them |  |  |  |

Phrases:
- 1:11.1 1:11.2 1:11.3 | This book, The Wars of the LORD, | the-book the-wars-of the-LORD
- 1:11.4 1:11.5 1:11.6 1:11.7 1:11.8 | makes known the way of the truths of His unity, may He be blessed, | that-makes-known the-way-of the-truths-of His-unity may-He-be-blessed
- 1:11.9 1:11.10 1:11.11 1:11.12 | according to the Torah and its interpreters, our Rabbis of blessed memory, | according-to the-Torah and-its-interpreters our-Rabbis-of-blessed-memory
- 1:11.13 1:11.14 1:11.15 1:11.16 | in the Mishnah, in the Bavli and the Yerushalmi, | in-the-Mishnah and-in-the-Talmud Bavli and-Yerushalmi
- 1:11.17 1:11.18 | and the Midrashim of our Rabbis of blessed memory, | and-the-Midrashim-of our-Rabbis-of-blessed-memory
- 1:11.19 1:11.20 1:11.21 | and in accordance with the view of the Geonim, Rabbi Saadia Gaon, | and-in-accordance-with the-Geonim Rabbi-Saadia-Gaon
- 1:11.22 1:11.23 1:11.24 1:11.25 | and our Rabbi Bahya in Chovot ha-Levavot, | and-our-Rabbi Bahya in-Chovot-ha-Levavot
- 1:11.26 1:11.27 1:11.28 1:11.29 | and Rabbi Yehuda Halevi in his poems, | and-Rabbi Yehuda Halevi in-his-poems
- 1:11.30 1:11.31 | and in the Kuzari, the Rambam, | and-in-the-Kuzari the-Rambam
- 1:11.32 1:11.33 1:11.34 1:11.35 | the SeMaG and the author of the Ikkarim, of blessed memory, | and-the-SeMaG and-the-author-of the-Ikkarim of-blessed-memory
- 1:11.36 1:11.37 1:11.38 1:11.39 1:11.40 | and as is accepted by every Jew. | and-as-is-accepted by every man-of Israel
- 1:11.41 1:11.42 1:11.43 1:11.44 1:11.45 1:11.46 1:11.47 | In it is a severe warning not to worship any created thing, | and-in-it a-warning severe that-not to-worship any created-thing
- 1:11.48 1:11.49 1:11.50 1:11.51 | in accordance with the tradition of all our prophets | in-accordance-with the-tradition-of all our-prophets
- 1:11.52 1:11.53 1:11.54 1:11.55 1:11.56 1:11.57 1:11.58 | and all our sages, the Tannaim, the Amoraim, the Geonim and the early decisors, | and-all our-sages the-Tannaim and-the-Amoraim and-the-Geonim and-the-decisors the-early
- 1:11.59 1:11.60 1:11.61 1:11.62 | who received it from the mouth of Moses, Aaron and the elders, | from-the-mouth-of Moses and-Aaron and-the-elders
- 1:11.63 1:11.64 1:11.65 1:11.66 1:11.67 1:11.68 | not to deviate from them and not to add to them. | not to-deviate from-them and-not to-add to-them

Grammar notes:
- The sentence has the subject ספר with the participle המודיע and the object דרך אמתות יחודו; the rest is a list of sources.
- The list is joined by ו and by ב (in): 'in the Mishnah', 'in the Bavli', 'in the Kuzari'.
- ובו אזהרה חמורה שלא לעבוד: 'and in it a severe warning not to worship'; the שלא clause gives the content of the warning.
- כפי קבלת … מפי משה ואהרן וזקנים לא לזוז מהם ולא להוסיף עליהם: the chain of tradition ends in two infinitives, 'not to deviate' and 'not to add'.

Translator's note: Title per the decisions (Q1). The Geonim and the Kuzari are given as context.

Commentary rulings:
- apparatus: Num 21:14, the title: ספר מלחמת ה׳ (The Wars of the LORD), as the title page gives it.

## Milchamot ha-Shem 1:12

Hebrew (Wikisource): ודע לך קורא נכבד כי לא תשיג כונת הספר הזה אלא אחרי שתקרא ותשנה את כלו כי מחברו עשה לילות כימים ויגע ודקדק בכמה ספרים ודלה מימים ונהרים ראיות כראי מוצק חזקים ולא מדעתו יצאו הדברים ולכן אבקשך קורא נכבד שלא תבוא באגרונו של עשו בהתולים אלא בקולו של יעקב ולהתבונן בדברים, ומי שהוא חסר אונים יאלם מדברים וידום ואל שדי יתן לכם רחמים וזכות המחבר תגן עלינו וע"כ עמו ישראל.

| id | word | punctuation after |
|---|---|---|
| 1:12.1 | ודע |  |
| 1:12.2 | לך |  |
| 1:12.3 | קורא |  |
| 1:12.4 | נכבד |  |
| 1:12.5 | כי |  |
| 1:12.6 | לא |  |
| 1:12.7 | תשיג |  |
| 1:12.8 | כונת |  |
| 1:12.9 | הספר |  |
| 1:12.10 | הזה |  |
| 1:12.11 | אלא |  |
| 1:12.12 | אחרי |  |
| 1:12.13 | שתקרא |  |
| 1:12.14 | ותשנה |  |
| 1:12.15 | את |  |
| 1:12.16 | כלו |  |
| 1:12.17 | כי |  |
| 1:12.18 | מחברו |  |
| 1:12.19 | עשה |  |
| 1:12.20 | לילות |  |
| 1:12.21 | כימים |  |
| 1:12.22 | ויגע |  |
| 1:12.23 | ודקדק |  |
| 1:12.24 | בכמה |  |
| 1:12.25 | ספרים |  |
| 1:12.26 | ודלה |  |
| 1:12.27 | מימים |  |
| 1:12.28 | ונהרים |  |
| 1:12.29 | ראיות |  |
| 1:12.30 | כראי |  |
| 1:12.31 | מוצק |  |
| 1:12.32 | חזקים |  |
| 1:12.33 | ולא |  |
| 1:12.34 | מדעתו |  |
| 1:12.35 | יצאו |  |
| 1:12.36 | הדברים |  |
| 1:12.37 | ולכן |  |
| 1:12.38 | אבקשך |  |
| 1:12.39 | קורא |  |
| 1:12.40 | נכבד |  |
| 1:12.41 | שלא |  |
| 1:12.42 | תבוא |  |
| 1:12.43 | באגרונו |  |
| 1:12.44 | של |  |
| 1:12.45 | עשו |  |
| 1:12.46 | בהתולים |  |
| 1:12.47 | אלא |  |
| 1:12.48 | בקולו |  |
| 1:12.49 | של |  |
| 1:12.50 | יעקב |  |
| 1:12.51 | ולהתבונן |  |
| 1:12.52 | בדברים | , |
| 1:12.53 | ומי |  |
| 1:12.54 | שהוא |  |
| 1:12.55 | חסר |  |
| 1:12.56 | אונים |  |
| 1:12.57 | יאלם |  |
| 1:12.58 | מדברים |  |
| 1:12.59 | וידום |  |
| 1:12.60 | ואל |  |
| 1:12.61 | שדי |  |
| 1:12.62 | יתן |  |
| 1:12.63 | לכם |  |
| 1:12.64 | רחמים |  |
| 1:12.65 | וזכות |  |
| 1:12.66 | המחבר |  |
| 1:12.67 | תגן |  |
| 1:12.68 | עלינו |  |
| 1:12.69 | וע"כ |  |
| 1:12.70 | עמו |  |
| 1:12.71 | ישראל | . |


### Translation under review (1:12)

English: Know, esteemed reader, that you will not grasp the intention of this book, unless you read and study all of it, for its author made nights as days, and labored and scrutinized many books, and drew from seas and rivers proofs strong as a cast mirror, and the matters did not come from his own opinion. Therefore I ask you, esteemed reader, that you not come with the fist of Esau, in mockery, but with the voice of Jacob, and to consider the matters, and whoever is lacking in strength, let him be silent from speech and be still, and may El Shaddai give you mercy, and the merit of the author will protect us, and upon all His people Israel.

Literal: and-know to-you reader esteemed that not you-will-grasp the-intention-of the-book this unless after you-read and-you-study (object) all-of-it for its-author made nights as-days and-labored and-examined in-many books and-drew from-seas and-rivers proofs like-a-mirror molten strong and-not from-his-mind came-out the-matters and-therefore I-ask-you reader esteemed that-not you-come with-his-fist of Esau in-mockery but with-his-voice of Jacob and-to-consider the-matters and-whoever who-is lacking of-strength let-him-be-silent from-speech and-be-still and-El Shaddai may-he-grant to-you mercy and-the-merit-of the-author shall-shield us and-upon-all His-people Israel

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:12.1 | and know |  |  | Imperative 'know', with 'to you'; the author addresses the reader. |
| 1:12.2 | to you |  |  |  |
| 1:12.3 | reader |  |  |  |
| 1:12.4 | esteemed |  |  |  |
| 1:12.5 | that |  |  |  |
| 1:12.6 | not |  |  |  |
| 1:12.7 | you will grasp |  |  | Second-person future with לא: 'you will not grasp'. |
| 1:12.8 | the intention of |  |  |  |
| 1:12.9 | the book |  |  |  |
| 1:12.10 | this |  |  |  |
| 1:12.11 | unless |  |  |  |
| 1:12.12 | after |  |  |  |
| 1:12.13 | you read |  |  |  |
| 1:12.14 | and you study |  |  |  |
| 1:12.15 | (object marker) |  |  |  |
| 1:12.16 | all of it |  |  |  |
| 1:12.17 | for |  |  |  |
| 1:12.18 | its author |  |  |  |
| 1:12.19 | made |  |  |  |
| 1:12.20 | nights |  |  |  |
| 1:12.21 | as days |  |  |  |
| 1:12.22 | and labored |  |  |  |
| 1:12.23 | and examined |  |  |  |
| 1:12.24 | many |  |  |  |
| 1:12.25 | books |  |  |  |
| 1:12.26 | and drew |  |  |  |
| 1:12.27 | from seas |  |  |  |
| 1:12.28 | and rivers |  |  |  |
| 1:12.29 | proofs |  |  |  |
| 1:12.30 | like a mirror |  |  |  |
| 1:12.31 | molten |  |  | מוצק: molten or cast metal; the mirror was of cast metal. |
| 1:12.32 | strong |  |  |  |
| 1:12.33 | and not |  |  |  |
| 1:12.34 | from his mind |  |  |  |
| 1:12.35 | came out |  |  |  |
| 1:12.36 | the matters |  |  |  |
| 1:12.37 | and therefore |  |  |  |
| 1:12.38 | I ask you |  |  |  |
| 1:12.39 | reader |  |  |  |
| 1:12.40 | esteemed |  |  |  |
| 1:12.41 | that not |  |  |  |
| 1:12.42 | you come |  |  |  |
| 1:12.43 | with his fist |  |  | A fist, the image for the hands of Esau (Gen 27:22; see the apparatus). |
| 1:12.44 | of |  |  |  |
| 1:12.45 | Esau |  |  |  |
| 1:12.46 | in mockery |  |  |  |
| 1:12.47 | but |  |  |  |
| 1:12.48 | with his voice |  |  |  |
| 1:12.49 | of |  |  |  |
| 1:12.50 | Jacob |  |  |  |
| 1:12.51 | and to consider |  |  |  |
| 1:12.52 | the matters |  |  |  |
| 1:12.53 | and whoever |  |  |  |
| 1:12.54 | who is |  |  |  |
| 1:12.55 | lacking |  |  |  |
| 1:12.56 | of strength |  |  |  |
| 1:12.57 | let him be silent |  |  |  |
| 1:12.58 | from speech |  |  |  |
| 1:12.59 | and be still |  |  |  |
| 1:12.60 | and El |  |  | El Shaddai, a divine name. |
| 1:12.61 | Shaddai |  |  |  |
| 1:12.62 | may he grant |  |  |  |
| 1:12.63 | to you |  |  |  |
| 1:12.64 | mercy |  |  |  |
| 1:12.65 | and the merit of |  |  |  |
| 1:12.66 | the author |  |  |  |
| 1:12.67 | shall shield |  |  |  |
| 1:12.68 | us |  |  |  |
| 1:12.69 | and upon all |  |  | ע״כ is read as ועל כל, 'and upon all', here. Elsewhere it closes a quotation or means 'thus far'. |
| 1:12.70 | His people |  |  |  |
| 1:12.71 | Israel |  |  |  |

Phrases:
- 1:12.1 1:12.2 1:12.3 1:12.4 | Know, esteemed reader, | and-know to-you reader esteemed
- 1:12.5 1:12.6 1:12.7 1:12.8 1:12.9 1:12.10 | that you will not grasp the intention of this book, | that not you-will-grasp the-intention-of the-book this
- 1:12.11 1:12.12 1:12.13 1:12.14 1:12.15 1:12.16 | unless you read and study all of it, | unless after you-read and-you-study (object) all-of-it
- 1:12.17 1:12.18 1:12.19 1:12.20 1:12.21 | for its author made nights as days, | for its-author made nights as-days
- 1:12.22 1:12.23 1:12.24 1:12.25 | and labored and scrutinized many books, | and-labored and-examined in-many books
- 1:12.26 1:12.27 1:12.28 1:12.29 1:12.30 1:12.31 1:12.32 | and drew from seas and rivers proofs strong as a cast mirror, | and-drew from-seas and-rivers proofs like-a-mirror molten strong
- 1:12.33 1:12.34 1:12.35 1:12.36 | and the matters did not come from his own opinion. | and-not from-his-mind came-out the-matters
- 1:12.37 1:12.38 1:12.39 1:12.40 | Therefore I ask you, esteemed reader, | and-therefore I-ask-you reader esteemed
- 1:12.41 1:12.42 1:12.43 1:12.44 1:12.45 1:12.46 | that you not come with the fist of Esau, in mockery, | that-not you-come with-his-fist of Esau in-mockery
- 1:12.47 1:12.48 1:12.49 1:12.50 1:12.51 1:12.52 | but with the voice of Jacob, and to consider the matters, | but with-his-voice of Jacob and-to-consider the-matters
- 1:12.53 1:12.54 1:12.55 1:12.56 1:12.57 1:12.58 1:12.59 | and whoever is lacking in strength, let him be silent from speech and be still, | and-whoever who-is lacking of-strength let-him-be-silent from-speech and-be-still
- 1:12.60 1:12.61 1:12.62 1:12.63 1:12.64 | and may El Shaddai give you mercy, | and-El Shaddai may-he-grant to-you mercy
- 1:12.65 1:12.66 1:12.67 1:12.68 | and the merit of the author will protect us, | and-the-merit-of the-author shall-shield us
- 1:12.69 1:12.70 1:12.71 | and upon all His people Israel. | and-upon-all His-people Israel

Grammar notes:
- ודע לך: an imperative with 'to you', 'know'; the author addresses the reader in the second person throughout.
- לא תשיג כונת הספר הזה אלא אחרי שתקרא: 'you will not grasp… unless after you read…'.
- כי מחברו עשה לילות כימים: the second כי is causal, 'for its author made nights as days'.
- ולכן אבקשך שלא תבוא באגרונו של עשו בהתולים אלא בקולו של יעקב: the request contrasts the fist of Esau with the voice of Jacob.

Translator's note: The allusion to Gen 27:22 is kept without quotation marks; the author's 'fist' is kept where the verse has 'hands'.

Commentary rulings:
- apparatus: Gen 27:22, an allusion without quotation marks: 'the voice is the voice of Jacob, but the hands are the hands of Esau'. The author writes אגרון (fist) where the verse has ידים (hands).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

