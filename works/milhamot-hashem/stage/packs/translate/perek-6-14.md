# Translation pack: milhamot-hashem perek-6-14 (Milchamot ha-Shem 6:22–6:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-14`.

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
- עשתונות איש ירושלים (a title): translate the words, "the thoughts of the man
  of Jerusalem", with a note; no identification beyond the text.
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
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
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
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


## Milchamot ha-Shem 6:22

Hebrew (Wikisource): ע) והפילסוף מחבר הזהר אסף את כל הדיעות האלה וטפל אותם וחברם לאמונת תורתינו הקדושה בהטעאת מליצת הכתובים ודברי חכמים שלא כדרך הצעת הלשון ולא כקבלת חכמים. והוא מאמת מציאות אלהים אחרים אשר לא אלהים המה היפך הכתובים המעידים שאין צור מבלעדי אלהינו. וכמבואר בדברי רז"ל בתלמוד ובמדרשיהם האמתיים. ודברי הפילסוף הזה שטפל על משה רבינו עליו השלום שקר שנענש על שנתן לישראל את המשנה הזאת ושהוא בעצמו התודה על חטאתו זאת כמו שנתבאר לעיל. המה נגד דברי רז"ל בגמרא ובמדרש רבה בכמה מקומות ובילקוט פ' ואתחנן וז"ל ויקרא פרשת רבה פרשה ל"א עם פירוש עץ יוסף, ר' יהודה מושלו משל למה הדבר דומה למלך שגזר ואמר כל מי שילקוט ויאכל מפגי פירות שביעית יהו מחזרין אותו בקומפון, הלכה אשה אחת בת טובים ולקטה ואכלה פגי פירות שביעית, התחילו מחזרין אותה בקומפון והיתה צוחת ואומרת בבקשה ממך אדני המלך תלה את הפגין האלו בצוארי כדי שלא יהו הבריות אומרין דומה לנו שנמצא בה דבר של ערוה או דבר של כשפים, אלא מתוך שרואין את הפגין בצוארי הן יודעין שבשבילן אני מחזרת, כך אמר משה לפני הקב"ה. רבון העולם כתוב בתורתך מפני מה איני נכנס לארץ, שלא יהו ישראל אומרין, דומה לנו שזייף משה בתורה (פי' בתורה שבכתב) או אמר דבר שלא הצטוה (פי' בתורה שבעל פה), אמר לו הקב"ה חייך שאני כותב בתורה שלא היתה אלא על המים, הדא הוא דכתיב כאשר מריתם פי במדבר צין וכו' ועיין יומא דף פ"ו.

| id | word | punctuation after |
|---|---|---|
| 6:22.1 | ע | ) |
| 6:22.2 | והפילסוף |  |
| 6:22.3 | מחבר |  |
| 6:22.4 | הזהר |  |
| 6:22.5 | אסף |  |
| 6:22.6 | את |  |
| 6:22.7 | כל |  |
| 6:22.8 | הדיעות |  |
| 6:22.9 | האלה |  |
| 6:22.10 | וטפל |  |
| 6:22.11 | אותם |  |
| 6:22.12 | וחברם |  |
| 6:22.13 | לאמונת |  |
| 6:22.14 | תורתינו |  |
| 6:22.15 | הקדושה |  |
| 6:22.16 | בהטעאת |  |
| 6:22.17 | מליצת |  |
| 6:22.18 | הכתובים |  |
| 6:22.19 | ודברי |  |
| 6:22.20 | חכמים |  |
| 6:22.21 | שלא |  |
| 6:22.22 | כדרך |  |
| 6:22.23 | הצעת |  |
| 6:22.24 | הלשון |  |
| 6:22.25 | ולא |  |
| 6:22.26 | כקבלת |  |
| 6:22.27 | חכמים | . |
| 6:22.28 | והוא |  |
| 6:22.29 | מאמת |  |
| 6:22.30 | מציאות |  |
| 6:22.31 | אלהים |  |
| 6:22.32 | אחרים |  |
| 6:22.33 | אשר |  |
| 6:22.34 | לא |  |
| 6:22.35 | אלהים |  |
| 6:22.36 | המה |  |
| 6:22.37 | היפך |  |
| 6:22.38 | הכתובים |  |
| 6:22.39 | המעידים |  |
| 6:22.40 | שאין |  |
| 6:22.41 | צור |  |
| 6:22.42 | מבלעדי |  |
| 6:22.43 | אלהינו | . |
| 6:22.44 | וכמבואר |  |
| 6:22.45 | בדברי |  |
| 6:22.46 | רז"ל |  |
| 6:22.47 | בתלמוד |  |
| 6:22.48 | ובמדרשיהם |  |
| 6:22.49 | האמתיים | . |
| 6:22.50 | ודברי |  |
| 6:22.51 | הפילסוף |  |
| 6:22.52 | הזה |  |
| 6:22.53 | שטפל |  |
| 6:22.54 | על |  |
| 6:22.55 | משה |  |
| 6:22.56 | רבינו |  |
| 6:22.57 | עליו |  |
| 6:22.58 | השלום |  |
| 6:22.59 | שקר |  |
| 6:22.60 | שנענש |  |
| 6:22.61 | על |  |
| 6:22.62 | שנתן |  |
| 6:22.63 | לישראל |  |
| 6:22.64 | את |  |
| 6:22.65 | המשנה |  |
| 6:22.66 | הזאת |  |
| 6:22.67 | ושהוא |  |
| 6:22.68 | בעצמו |  |
| 6:22.69 | התודה |  |
| 6:22.70 | על |  |
| 6:22.71 | חטאתו |  |
| 6:22.72 | זאת |  |
| 6:22.73 | כמו |  |
| 6:22.74 | שנתבאר |  |
| 6:22.75 | לעיל | . |
| 6:22.76 | המה |  |
| 6:22.77 | נגד |  |
| 6:22.78 | דברי |  |
| 6:22.79 | רז"ל |  |
| 6:22.80 | בגמרא |  |
| 6:22.81 | ובמדרש |  |
| 6:22.82 | רבה |  |
| 6:22.83 | בכמה |  |
| 6:22.84 | מקומות |  |
| 6:22.85 | ובילקוט |  |
| 6:22.86 | פ | ' |
| 6:22.87 | ואתחנן |  |
| 6:22.88 | וז"ל |  |
| 6:22.89 | ויקרא |  |
| 6:22.90 | פרשת |  |
| 6:22.91 | רבה |  |
| 6:22.92 | פרשה |  |
| 6:22.93 | ל"א |  |
| 6:22.94 | עם |  |
| 6:22.95 | פירוש |  |
| 6:22.96 | עץ |  |
| 6:22.97 | יוסף | , |
| 6:22.98 | ר | ' |
| 6:22.99 | יהודה |  |
| 6:22.100 | מושלו |  |
| 6:22.101 | משל |  |
| 6:22.102 | למה |  |
| 6:22.103 | הדבר |  |
| 6:22.104 | דומה |  |
| 6:22.105 | למלך |  |
| 6:22.106 | שגזר |  |
| 6:22.107 | ואמר |  |
| 6:22.108 | כל |  |
| 6:22.109 | מי |  |
| 6:22.110 | שילקוט |  |
| 6:22.111 | ויאכל |  |
| 6:22.112 | מפגי |  |
| 6:22.113 | פירות |  |
| 6:22.114 | שביעית |  |
| 6:22.115 | יהו |  |
| 6:22.116 | מחזרין |  |
| 6:22.117 | אותו |  |
| 6:22.118 | בקומפון | , |
| 6:22.119 | הלכה |  |
| 6:22.120 | אשה |  |
| 6:22.121 | אחת |  |
| 6:22.122 | בת |  |
| 6:22.123 | טובים |  |
| 6:22.124 | ולקטה |  |
| 6:22.125 | ואכלה |  |
| 6:22.126 | פגי |  |
| 6:22.127 | פירות |  |
| 6:22.128 | שביעית | , |
| 6:22.129 | התחילו |  |
| 6:22.130 | מחזרין |  |
| 6:22.131 | אותה |  |
| 6:22.132 | בקומפון |  |
| 6:22.133 | והיתה |  |
| 6:22.134 | צוחת |  |
| 6:22.135 | ואומרת |  |
| 6:22.136 | בבקשה |  |
| 6:22.137 | ממך |  |
| 6:22.138 | אדני |  |
| 6:22.139 | המלך |  |
| 6:22.140 | תלה |  |
| 6:22.141 | את |  |
| 6:22.142 | הפגין |  |
| 6:22.143 | האלו |  |
| 6:22.144 | בצוארי |  |
| 6:22.145 | כדי |  |
| 6:22.146 | שלא |  |
| 6:22.147 | יהו |  |
| 6:22.148 | הבריות |  |
| 6:22.149 | אומרין |  |
| 6:22.150 | דומה |  |
| 6:22.151 | לנו |  |
| 6:22.152 | שנמצא |  |
| 6:22.153 | בה |  |
| 6:22.154 | דבר |  |
| 6:22.155 | של |  |
| 6:22.156 | ערוה |  |
| 6:22.157 | או |  |
| 6:22.158 | דבר |  |
| 6:22.159 | של |  |
| 6:22.160 | כשפים | , |
| 6:22.161 | אלא |  |
| 6:22.162 | מתוך |  |
| 6:22.163 | שרואין |  |
| 6:22.164 | את |  |
| 6:22.165 | הפגין |  |
| 6:22.166 | בצוארי |  |
| 6:22.167 | הן |  |
| 6:22.168 | יודעין |  |
| 6:22.169 | שבשבילן |  |
| 6:22.170 | אני |  |
| 6:22.171 | מחזרת | , |
| 6:22.172 | כך |  |
| 6:22.173 | אמר |  |
| 6:22.174 | משה |  |
| 6:22.175 | לפני |  |
| 6:22.176 | הקב"ה | . |
| 6:22.177 | רבון |  |
| 6:22.178 | העולם |  |
| 6:22.179 | כתוב |  |
| 6:22.180 | בתורתך |  |
| 6:22.181 | מפני |  |
| 6:22.182 | מה |  |
| 6:22.183 | איני |  |
| 6:22.184 | נכנס |  |
| 6:22.185 | לארץ | , |
| 6:22.186 | שלא |  |
| 6:22.187 | יהו |  |
| 6:22.188 | ישראל |  |
| 6:22.189 | אומרין | , |
| 6:22.190 | דומה |  |
| 6:22.191 | לנו |  |
| 6:22.192 | שזייף |  |
| 6:22.193 | משה |  |
| 6:22.194 | בתורה |  |
| 6:22.195 | פי | ' |
| 6:22.196 | בתורה |  |
| 6:22.197 | שבכתב | ) |
| 6:22.198 | או |  |
| 6:22.199 | אמר |  |
| 6:22.200 | דבר |  |
| 6:22.201 | שלא |  |
| 6:22.202 | הצטוה |  |
| 6:22.203 | פי | ' |
| 6:22.204 | בתורה |  |
| 6:22.205 | שבעל |  |
| 6:22.206 | פה | ), |
| 6:22.207 | אמר |  |
| 6:22.208 | לו |  |
| 6:22.209 | הקב"ה |  |
| 6:22.210 | חייך |  |
| 6:22.211 | שאני |  |
| 6:22.212 | כותב |  |
| 6:22.213 | בתורה |  |
| 6:22.214 | שלא |  |
| 6:22.215 | היתה |  |
| 6:22.216 | אלא |  |
| 6:22.217 | על |  |
| 6:22.218 | המים | , |
| 6:22.219 | הדא |  |
| 6:22.220 | הוא |  |
| 6:22.221 | דכתיב |  |
| 6:22.222 | כאשר |  |
| 6:22.223 | מריתם |  |
| 6:22.224 | פי |  |
| 6:22.225 | במדבר |  |
| 6:22.226 | צין |  |
| 6:22.227 | וכו | ' |
| 6:22.228 | ועיין |  |
| 6:22.229 | יומא |  |
| 6:22.230 | דף |  |
| 6:22.231 | פ"ו | . |


## Milchamot ha-Shem 6:23

Hebrew (Wikisource): הרי מכל זה מתבאר שתורתו של משה רבינו ומשנתו אמת מפי הגבורה אמרם ולא שינה לקבל מטטרון משנה שיש בה טוב ורע. והרי מצינו שאף עדת קרח החולקים על משה רבינו ע"ה וירדו חיים שאולה מתודים את חטאם ואומרים משה אמת ותורתו אמת והם בדאים. כדאיתה בפרק חלק ובפרק הספינה. ואם כן חלילה לכל אוהב תורת משה רבינו ע"ה וירדו חיים שאולה מתודים את חטאם ואומרים משה אמת ותורתו אמת והם בדאים. כדאיתה בפרק חלק ובפרק הספינה. ואם כן חלילה לכל אוהב תורת משה רבינו ומאמין בקבלת רז"ל במשנה ובתלמוד, להאמין לשקריו של מחבר הזהר הוא והמלך הנכרי אשר על ידו רמה את ישראל וגילה את חיבורו הזהר וכיוון לטפול שקר על תורת משה שבעל פה המקובלת איש מפי איש בלי שופ ספק וליתן בה דופי שהיא מעורבת מטוב ורע, ושטוב בעיניו לשוב בתשובה ולעזוב את המשנה והתלמוד, וללמוד בספריו החיצוניים הזהר והתיקונים. כדי לשכח מישראל כללותיה ופרטותיה של תורה, וגם לשכח מהם אמתת יחודו של הקב"ה. ולטעת במקומה אמונת אלהים רבים המשתלשלים ומסתעפים זה מזה. וקרא זה אל זה ואמר יהי כך וכך, והאחר עושה. או התחתון יטול רשות מהעליון, ואם לא יחפוץ העליון, יאמר לו התחתון מה לי ולך, אם יחטא לי יחטא ולא לך. כמו שאמרה אימא לאבא בעת בריאת אדם הראשון לפי דבריו אשר בדא מלבו כדלעיל. ועליו אמר שלמה המלך ע"ה ובכן ראיתי (ספרים) קבורים ובאו (לנו מיד בן נכר) וממקום קדוש יהלכו (כאלו מהלכה למשה מסיני נאמרו) וישתכחו בעיר שכן קבלו מרשב"י התנא, אל תקרי וישתכחו אלא וישתבחו כמו שדרש רז"ל, אשר כן עשו להטעות את ישראל ולהרחיקם מעל ה' אלהיהם שהוא הסבה הראשונה. ולעבוד סבות רבות מחודשות, ובייחוד לאלוה קצא אפים (זעיר אנפין ונוקביה) וליחס כל תפלתם ועבודתם אליהם, ולקרות להן בשמותיו של הקב"ה כמעשה דורו של אנוש שקראו לעצבים בשמותיו של הקב"ה עד שהציף השי"ת מי אוקינוס ואיבד שלישו של עולם כמו שאמרו רז"ל והבאתיו לעיל בסי' נ"ו.

| id | word | punctuation after |
|---|---|---|
| 6:23.1 | הרי |  |
| 6:23.2 | מכל |  |
| 6:23.3 | זה |  |
| 6:23.4 | מתבאר |  |
| 6:23.5 | שתורתו |  |
| 6:23.6 | של |  |
| 6:23.7 | משה |  |
| 6:23.8 | רבינו |  |
| 6:23.9 | ומשנתו |  |
| 6:23.10 | אמת |  |
| 6:23.11 | מפי |  |
| 6:23.12 | הגבורה |  |
| 6:23.13 | אמרם |  |
| 6:23.14 | ולא |  |
| 6:23.15 | שינה |  |
| 6:23.16 | לקבל |  |
| 6:23.17 | מטטרון |  |
| 6:23.18 | משנה |  |
| 6:23.19 | שיש |  |
| 6:23.20 | בה |  |
| 6:23.21 | טוב |  |
| 6:23.22 | ורע | . |
| 6:23.23 | והרי |  |
| 6:23.24 | מצינו |  |
| 6:23.25 | שאף |  |
| 6:23.26 | עדת |  |
| 6:23.27 | קרח |  |
| 6:23.28 | החולקים |  |
| 6:23.29 | על |  |
| 6:23.30 | משה |  |
| 6:23.31 | רבינו |  |
| 6:23.32 | ע"ה |  |
| 6:23.33 | וירדו |  |
| 6:23.34 | חיים |  |
| 6:23.35 | שאולה |  |
| 6:23.36 | מתודים |  |
| 6:23.37 | את |  |
| 6:23.38 | חטאם |  |
| 6:23.39 | ואומרים |  |
| 6:23.40 | משה |  |
| 6:23.41 | אמת |  |
| 6:23.42 | ותורתו |  |
| 6:23.43 | אמת |  |
| 6:23.44 | והם |  |
| 6:23.45 | בדאים | . |
| 6:23.46 | כדאיתה |  |
| 6:23.47 | בפרק |  |
| 6:23.48 | חלק |  |
| 6:23.49 | ובפרק |  |
| 6:23.50 | הספינה | . |
| 6:23.51 | ואם |  |
| 6:23.52 | כן |  |
| 6:23.53 | חלילה |  |
| 6:23.54 | לכל |  |
| 6:23.55 | אוהב |  |
| 6:23.56 | תורת |  |
| 6:23.57 | משה |  |
| 6:23.58 | רבינו |  |
| 6:23.59 | ע"ה |  |
| 6:23.60 | וירדו |  |
| 6:23.61 | חיים |  |
| 6:23.62 | שאולה |  |
| 6:23.63 | מתודים |  |
| 6:23.64 | את |  |
| 6:23.65 | חטאם |  |
| 6:23.66 | ואומרים |  |
| 6:23.67 | משה |  |
| 6:23.68 | אמת |  |
| 6:23.69 | ותורתו |  |
| 6:23.70 | אמת |  |
| 6:23.71 | והם |  |
| 6:23.72 | בדאים | . |
| 6:23.73 | כדאיתה |  |
| 6:23.74 | בפרק |  |
| 6:23.75 | חלק |  |
| 6:23.76 | ובפרק |  |
| 6:23.77 | הספינה | . |
| 6:23.78 | ואם |  |
| 6:23.79 | כן |  |
| 6:23.80 | חלילה |  |
| 6:23.81 | לכל |  |
| 6:23.82 | אוהב |  |
| 6:23.83 | תורת |  |
| 6:23.84 | משה |  |
| 6:23.85 | רבינו |  |
| 6:23.86 | ומאמין |  |
| 6:23.87 | בקבלת |  |
| 6:23.88 | רז"ל |  |
| 6:23.89 | במשנה |  |
| 6:23.90 | ובתלמוד | , |
| 6:23.91 | להאמין |  |
| 6:23.92 | לשקריו |  |
| 6:23.93 | של |  |
| 6:23.94 | מחבר |  |
| 6:23.95 | הזהר |  |
| 6:23.96 | הוא |  |
| 6:23.97 | והמלך |  |
| 6:23.98 | הנכרי |  |
| 6:23.99 | אשר |  |
| 6:23.100 | על |  |
| 6:23.101 | ידו |  |
| 6:23.102 | רמה |  |
| 6:23.103 | את |  |
| 6:23.104 | ישראל |  |
| 6:23.105 | וגילה |  |
| 6:23.106 | את |  |
| 6:23.107 | חיבורו |  |
| 6:23.108 | הזהר |  |
| 6:23.109 | וכיוון |  |
| 6:23.110 | לטפול |  |
| 6:23.111 | שקר |  |
| 6:23.112 | על |  |
| 6:23.113 | תורת |  |
| 6:23.114 | משה |  |
| 6:23.115 | שבעל |  |
| 6:23.116 | פה |  |
| 6:23.117 | המקובלת |  |
| 6:23.118 | איש |  |
| 6:23.119 | מפי |  |
| 6:23.120 | איש |  |
| 6:23.121 | בלי |  |
| 6:23.122 | שופ |  |
| 6:23.123 | ספק |  |
| 6:23.124 | וליתן |  |
| 6:23.125 | בה |  |
| 6:23.126 | דופי |  |
| 6:23.127 | שהיא |  |
| 6:23.128 | מעורבת |  |
| 6:23.129 | מטוב |  |
| 6:23.130 | ורע | , |
| 6:23.131 | ושטוב |  |
| 6:23.132 | בעיניו |  |
| 6:23.133 | לשוב |  |
| 6:23.134 | בתשובה |  |
| 6:23.135 | ולעזוב |  |
| 6:23.136 | את |  |
| 6:23.137 | המשנה |  |
| 6:23.138 | והתלמוד | , |
| 6:23.139 | וללמוד |  |
| 6:23.140 | בספריו |  |
| 6:23.141 | החיצוניים |  |
| 6:23.142 | הזהר |  |
| 6:23.143 | והתיקונים | . |
| 6:23.144 | כדי |  |
| 6:23.145 | לשכח |  |
| 6:23.146 | מישראל |  |
| 6:23.147 | כללותיה |  |
| 6:23.148 | ופרטותיה |  |
| 6:23.149 | של |  |
| 6:23.150 | תורה | , |
| 6:23.151 | וגם |  |
| 6:23.152 | לשכח |  |
| 6:23.153 | מהם |  |
| 6:23.154 | אמתת |  |
| 6:23.155 | יחודו |  |
| 6:23.156 | של |  |
| 6:23.157 | הקב"ה | . |
| 6:23.158 | ולטעת |  |
| 6:23.159 | במקומה |  |
| 6:23.160 | אמונת |  |
| 6:23.161 | אלהים |  |
| 6:23.162 | רבים |  |
| 6:23.163 | המשתלשלים |  |
| 6:23.164 | ומסתעפים |  |
| 6:23.165 | זה |  |
| 6:23.166 | מזה | . |
| 6:23.167 | וקרא |  |
| 6:23.168 | זה |  |
| 6:23.169 | אל |  |
| 6:23.170 | זה |  |
| 6:23.171 | ואמר |  |
| 6:23.172 | יהי |  |
| 6:23.173 | כך |  |
| 6:23.174 | וכך | , |
| 6:23.175 | והאחר |  |
| 6:23.176 | עושה | . |
| 6:23.177 | או |  |
| 6:23.178 | התחתון |  |
| 6:23.179 | יטול |  |
| 6:23.180 | רשות |  |
| 6:23.181 | מהעליון | , |
| 6:23.182 | ואם |  |
| 6:23.183 | לא |  |
| 6:23.184 | יחפוץ |  |
| 6:23.185 | העליון | , |
| 6:23.186 | יאמר |  |
| 6:23.187 | לו |  |
| 6:23.188 | התחתון |  |
| 6:23.189 | מה |  |
| 6:23.190 | לי |  |
| 6:23.191 | ולך | , |
| 6:23.192 | אם |  |
| 6:23.193 | יחטא |  |
| 6:23.194 | לי |  |
| 6:23.195 | יחטא |  |
| 6:23.196 | ולא |  |
| 6:23.197 | לך | . |
| 6:23.198 | כמו |  |
| 6:23.199 | שאמרה |  |
| 6:23.200 | אימא |  |
| 6:23.201 | לאבא |  |
| 6:23.202 | בעת |  |
| 6:23.203 | בריאת |  |
| 6:23.204 | אדם |  |
| 6:23.205 | הראשון |  |
| 6:23.206 | לפי |  |
| 6:23.207 | דבריו |  |
| 6:23.208 | אשר |  |
| 6:23.209 | בדא |  |
| 6:23.210 | מלבו |  |
| 6:23.211 | כדלעיל | . |
| 6:23.212 | ועליו |  |
| 6:23.213 | אמר |  |
| 6:23.214 | שלמה |  |
| 6:23.215 | המלך |  |
| 6:23.216 | ע"ה |  |
| 6:23.217 | ובכן |  |
| 6:23.218 | ראיתי |  |
| 6:23.219 | ספרים | ) |
| 6:23.220 | קבורים |  |
| 6:23.221 | ובאו |  |
| 6:23.222 | לנו |  |
| 6:23.223 | מיד |  |
| 6:23.224 | בן |  |
| 6:23.225 | נכר | ) |
| 6:23.226 | וממקום |  |
| 6:23.227 | קדוש |  |
| 6:23.228 | יהלכו |  |
| 6:23.229 | כאלו |  |
| 6:23.230 | מהלכה |  |
| 6:23.231 | למשה |  |
| 6:23.232 | מסיני |  |
| 6:23.233 | נאמרו | ) |
| 6:23.234 | וישתכחו |  |
| 6:23.235 | בעיר |  |
| 6:23.236 | שכן |  |
| 6:23.237 | קבלו |  |
| 6:23.238 | מרשב"י |  |
| 6:23.239 | התנא | , |
| 6:23.240 | אל |  |
| 6:23.241 | תקרי |  |
| 6:23.242 | וישתכחו |  |
| 6:23.243 | אלא |  |
| 6:23.244 | וישתבחו |  |
| 6:23.245 | כמו |  |
| 6:23.246 | שדרש |  |
| 6:23.247 | רז"ל | , |
| 6:23.248 | אשר |  |
| 6:23.249 | כן |  |
| 6:23.250 | עשו |  |
| 6:23.251 | להטעות |  |
| 6:23.252 | את |  |
| 6:23.253 | ישראל |  |
| 6:23.254 | ולהרחיקם |  |
| 6:23.255 | מעל |  |
| 6:23.256 | ה | ' |
| 6:23.257 | אלהיהם |  |
| 6:23.258 | שהוא |  |
| 6:23.259 | הסבה |  |
| 6:23.260 | הראשונה | . |
| 6:23.261 | ולעבוד |  |
| 6:23.262 | סבות |  |
| 6:23.263 | רבות |  |
| 6:23.264 | מחודשות | , |
| 6:23.265 | ובייחוד |  |
| 6:23.266 | לאלוה |  |
| 6:23.267 | קצא |  |
| 6:23.268 | אפים |  |
| 6:23.269 | זעיר |  |
| 6:23.270 | אנפין |  |
| 6:23.271 | ונוקביה | ) |
| 6:23.272 | וליחס |  |
| 6:23.273 | כל |  |
| 6:23.274 | תפלתם |  |
| 6:23.275 | ועבודתם |  |
| 6:23.276 | אליהם | , |
| 6:23.277 | ולקרות |  |
| 6:23.278 | להן |  |
| 6:23.279 | בשמותיו |  |
| 6:23.280 | של |  |
| 6:23.281 | הקב"ה |  |
| 6:23.282 | כמעשה |  |
| 6:23.283 | דורו |  |
| 6:23.284 | של |  |
| 6:23.285 | אנוש |  |
| 6:23.286 | שקראו |  |
| 6:23.287 | לעצבים |  |
| 6:23.288 | בשמותיו |  |
| 6:23.289 | של |  |
| 6:23.290 | הקב"ה |  |
| 6:23.291 | עד |  |
| 6:23.292 | שהציף |  |
| 6:23.293 | השי"ת |  |
| 6:23.294 | מי |  |
| 6:23.295 | אוקינוס |  |
| 6:23.296 | ואיבד |  |
| 6:23.297 | שלישו |  |
| 6:23.298 | של |  |
| 6:23.299 | עולם |  |
| 6:23.300 | כמו |  |
| 6:23.301 | שאמרו |  |
| 6:23.302 | רז"ל |  |
| 6:23.303 | והבאתיו |  |
| 6:23.304 | לעיל |  |
| 6:23.305 | בסי | ' |
| 6:23.306 | נ"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

