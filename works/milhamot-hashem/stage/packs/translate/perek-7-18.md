# Translation pack: milhamot-hashem perek-7-18 (Milchamot ha-Shem 7:23–7:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-18`.

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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.

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


## Milchamot ha-Shem 7:23

Hebrew (Wikisource): צה) וכל אלה השגיאות נמשכו ונשתרבבו בישראל כשיני עוג מלך הבשן מהאנשים הפתאים אוהבי הנפלאות והנמנעות שנפתו למלך הנכרי שרימא אותם והלכו אחר דעותיו לפרש פסוקי התורה על פי רעיונותיהם העולים על רוחם שלא על דרך רז"ל בעלי התלמוד שהם בעלי הקבלה האמתית. וחדשו מנהגים ופסקי הלכות על פי הקבלה הנכרית החדשה, ונדפסו בסדורי התפלות. עד שנעשית כתורה חדשה מתנגדת לתוה"ק שבע"פ. המבוארת בתלמוד ובפוסקים הראשונים בענין האלהות ויחוד השם ית'. אמנם כאשר נשים לב לעיין יפה בדברי רז"ל ובדברי הזהר נוכל להוכיח בצדק. כי לא מפי רשב"י יצאו הדברים הנזכרים בזהר, שהרי בבראשית רבה פר' כ"ו אמרו, וז"ל ויראו בני האלהים. רשב"י קרא להו בנו דייניא. רשב"י הוה מקלל לכל מאן דקרי להו בני אלהיא תני רשב"י כל פרצה שאינה מן הגדולים אינה פרצה. כומרים גנבו אלהיא מאן מומי ביה או מאן מקריב ליה. ולמה קורא אותן בני האלהים ר' חנינה ור' שמעון בן לקיש תרויחון אמרין שהרבו ימים בלא צער ובלא יסורין וכו' ע"כ. ראה ידידי כמה הוא מרחיק רשב"י ושונא ומקלל את האומר שהם מלאכים? ובזהר מחבב סברא זו וז"ל (בראשית כ"ג) דבר אחר נעשה אדם בצלמינו כדמותינו, הא אוקמוה חבריא על מלאכי השרת דאינון אמרי להאי קרא, א"ל בתר דהוו ידעין מה דהוה, ומה דעתיד להוי ואינון הוו יודעין דעתיד למחטי אמאי בעו למעבד ליה, ולא עוד אלא דעז"א ועזא"ל הוו מקטרגי עליה בזמנא דאמר שכינתא לקב"ה נעשה אדם, אמרו מה אדם ותדעהו, מה את בעי למברי אדם. ותדעהו דעתיד למחטי באתתא דיליה דאיהי חשך, דאור איהו דכורא, וחשך נוקבא שמאלא דבריאה, בההוא זמנא אמר שכינתא לון, בהאי דאתון מקטרגין אתון עתידין למנפל. כדכתיב ויראו בני האלהים את בנות האדם כי טבת הנה. חשקו בהון וטעו בהון. ואפיל להון שכינתא מקדושה דילהון. אמרו חבריא רבי רבי אי הכי עזא ועזאל לא הוו משקרי במלוליהו. דודאי בנוקבא עתיד אדם למחטי. אמר להו הכי אמרה שכינתא. אתון דאזדמנתן לקטרגא קדמי יתיר מחילי מרומא. אי אתון חזיתון שפירין מאדם בעובדיכון יאות לכון לקטרגא עליה, אבל איהו עתיד למחטי באתתא חדא. אתון בנשי סגיאי חבתון יתיר מבני נשא דכתיב ויראו בני האלהים את בנות האדם כי טבת הנה וכו'. גם בדף נ"ה חזר לשנותה בשם ר' יצחק. ובדף נ"ח חזר ושלשה בשם ר"י ור"ח. הרי שמחבר הזהר מחבב סברא זו ולפיכך אמרה ושנאה ושילשה לרוב חבתה. והרי בבראשית רבה מבורא שרשב"י שונא סברא זו עד שהיה מקלל את האומר שהם בני אלהיא. אלא ודאי לא מפי רשב"י יצאו דברים אלו האמורים בזהר. ובשקר יחסוהו לרשב"י.

| id | word | punctuation after |
|---|---|---|
| 7:23.1 | צה | ) |
| 7:23.2 | וכל |  |
| 7:23.3 | אלה |  |
| 7:23.4 | השגיאות |  |
| 7:23.5 | נמשכו |  |
| 7:23.6 | ונשתרבבו |  |
| 7:23.7 | בישראל |  |
| 7:23.8 | כשיני |  |
| 7:23.9 | עוג |  |
| 7:23.10 | מלך |  |
| 7:23.11 | הבשן |  |
| 7:23.12 | מהאנשים |  |
| 7:23.13 | הפתאים |  |
| 7:23.14 | אוהבי |  |
| 7:23.15 | הנפלאות |  |
| 7:23.16 | והנמנעות |  |
| 7:23.17 | שנפתו |  |
| 7:23.18 | למלך |  |
| 7:23.19 | הנכרי |  |
| 7:23.20 | שרימא |  |
| 7:23.21 | אותם |  |
| 7:23.22 | והלכו |  |
| 7:23.23 | אחר |  |
| 7:23.24 | דעותיו |  |
| 7:23.25 | לפרש |  |
| 7:23.26 | פסוקי |  |
| 7:23.27 | התורה |  |
| 7:23.28 | על |  |
| 7:23.29 | פי |  |
| 7:23.30 | רעיונותיהם |  |
| 7:23.31 | העולים |  |
| 7:23.32 | על |  |
| 7:23.33 | רוחם |  |
| 7:23.34 | שלא |  |
| 7:23.35 | על |  |
| 7:23.36 | דרך |  |
| 7:23.37 | רז"ל |  |
| 7:23.38 | בעלי |  |
| 7:23.39 | התלמוד |  |
| 7:23.40 | שהם |  |
| 7:23.41 | בעלי |  |
| 7:23.42 | הקבלה |  |
| 7:23.43 | האמתית | . |
| 7:23.44 | וחדשו |  |
| 7:23.45 | מנהגים |  |
| 7:23.46 | ופסקי |  |
| 7:23.47 | הלכות |  |
| 7:23.48 | על |  |
| 7:23.49 | פי |  |
| 7:23.50 | הקבלה |  |
| 7:23.51 | הנכרית |  |
| 7:23.52 | החדשה | , |
| 7:23.53 | ונדפסו |  |
| 7:23.54 | בסדורי |  |
| 7:23.55 | התפלות | . |
| 7:23.56 | עד |  |
| 7:23.57 | שנעשית |  |
| 7:23.58 | כתורה |  |
| 7:23.59 | חדשה |  |
| 7:23.60 | מתנגדת |  |
| 7:23.61 | לתוה"ק |  |
| 7:23.62 | שבע"פ | . |
| 7:23.63 | המבוארת |  |
| 7:23.64 | בתלמוד |  |
| 7:23.65 | ובפוסקים |  |
| 7:23.66 | הראשונים |  |
| 7:23.67 | בענין |  |
| 7:23.68 | האלהות |  |
| 7:23.69 | ויחוד |  |
| 7:23.70 | השם |  |
| 7:23.71 | ית | '. |
| 7:23.72 | אמנם |  |
| 7:23.73 | כאשר |  |
| 7:23.74 | נשים |  |
| 7:23.75 | לב |  |
| 7:23.76 | לעיין |  |
| 7:23.77 | יפה |  |
| 7:23.78 | בדברי |  |
| 7:23.79 | רז"ל |  |
| 7:23.80 | ובדברי |  |
| 7:23.81 | הזהר |  |
| 7:23.82 | נוכל |  |
| 7:23.83 | להוכיח |  |
| 7:23.84 | בצדק | . |
| 7:23.85 | כי |  |
| 7:23.86 | לא |  |
| 7:23.87 | מפי |  |
| 7:23.88 | רשב"י |  |
| 7:23.89 | יצאו |  |
| 7:23.90 | הדברים |  |
| 7:23.91 | הנזכרים |  |
| 7:23.92 | בזהר | , |
| 7:23.93 | שהרי |  |
| 7:23.94 | בבראשית |  |
| 7:23.95 | רבה |  |
| 7:23.96 | פר | ' |
| 7:23.97 | כ"ו |  |
| 7:23.98 | אמרו | , |
| 7:23.99 | וז"ל |  |
| 7:23.100 | ויראו |  |
| 7:23.101 | בני |  |
| 7:23.102 | האלהים | . |
| 7:23.103 | רשב"י |  |
| 7:23.104 | קרא |  |
| 7:23.105 | להו |  |
| 7:23.106 | בנו |  |
| 7:23.107 | דייניא | . |
| 7:23.108 | רשב"י |  |
| 7:23.109 | הוה |  |
| 7:23.110 | מקלל |  |
| 7:23.111 | לכל |  |
| 7:23.112 | מאן |  |
| 7:23.113 | דקרי |  |
| 7:23.114 | להו |  |
| 7:23.115 | בני |  |
| 7:23.116 | אלהיא |  |
| 7:23.117 | תני |  |
| 7:23.118 | רשב"י |  |
| 7:23.119 | כל |  |
| 7:23.120 | פרצה |  |
| 7:23.121 | שאינה |  |
| 7:23.122 | מן |  |
| 7:23.123 | הגדולים |  |
| 7:23.124 | אינה |  |
| 7:23.125 | פרצה | . |
| 7:23.126 | כומרים |  |
| 7:23.127 | גנבו |  |
| 7:23.128 | אלהיא |  |
| 7:23.129 | מאן |  |
| 7:23.130 | מומי |  |
| 7:23.131 | ביה |  |
| 7:23.132 | או |  |
| 7:23.133 | מאן |  |
| 7:23.134 | מקריב |  |
| 7:23.135 | ליה | . |
| 7:23.136 | ולמה |  |
| 7:23.137 | קורא |  |
| 7:23.138 | אותן |  |
| 7:23.139 | בני |  |
| 7:23.140 | האלהים |  |
| 7:23.141 | ר | ' |
| 7:23.142 | חנינה |  |
| 7:23.143 | ור | ' |
| 7:23.144 | שמעון |  |
| 7:23.145 | בן |  |
| 7:23.146 | לקיש |  |
| 7:23.147 | תרויחון |  |
| 7:23.148 | אמרין |  |
| 7:23.149 | שהרבו |  |
| 7:23.150 | ימים |  |
| 7:23.151 | בלא |  |
| 7:23.152 | צער |  |
| 7:23.153 | ובלא |  |
| 7:23.154 | יסורין |  |
| 7:23.155 | וכו | ' |
| 7:23.156 | ע"כ | . |
| 7:23.157 | ראה |  |
| 7:23.158 | ידידי |  |
| 7:23.159 | כמה |  |
| 7:23.160 | הוא |  |
| 7:23.161 | מרחיק |  |
| 7:23.162 | רשב"י |  |
| 7:23.163 | ושונא |  |
| 7:23.164 | ומקלל |  |
| 7:23.165 | את |  |
| 7:23.166 | האומר |  |
| 7:23.167 | שהם |  |
| 7:23.168 | מלאכים | ? |
| 7:23.169 | ובזהר |  |
| 7:23.170 | מחבב |  |
| 7:23.171 | סברא |  |
| 7:23.172 | זו |  |
| 7:23.173 | וז"ל |  |
| 7:23.174 | בראשית |  |
| 7:23.175 | כ"ג | ) |
| 7:23.176 | דבר |  |
| 7:23.177 | אחר |  |
| 7:23.178 | נעשה |  |
| 7:23.179 | אדם |  |
| 7:23.180 | בצלמינו |  |
| 7:23.181 | כדמותינו | , |
| 7:23.182 | הא |  |
| 7:23.183 | אוקמוה |  |
| 7:23.184 | חבריא |  |
| 7:23.185 | על |  |
| 7:23.186 | מלאכי |  |
| 7:23.187 | השרת |  |
| 7:23.188 | דאינון |  |
| 7:23.189 | אמרי |  |
| 7:23.190 | להאי |  |
| 7:23.191 | קרא | , |
| 7:23.192 | א"ל |  |
| 7:23.193 | בתר |  |
| 7:23.194 | דהוו |  |
| 7:23.195 | ידעין |  |
| 7:23.196 | מה |  |
| 7:23.197 | דהוה | , |
| 7:23.198 | ומה |  |
| 7:23.199 | דעתיד |  |
| 7:23.200 | להוי |  |
| 7:23.201 | ואינון |  |
| 7:23.202 | הוו |  |
| 7:23.203 | יודעין |  |
| 7:23.204 | דעתיד |  |
| 7:23.205 | למחטי |  |
| 7:23.206 | אמאי |  |
| 7:23.207 | בעו |  |
| 7:23.208 | למעבד |  |
| 7:23.209 | ליה | , |
| 7:23.210 | ולא |  |
| 7:23.211 | עוד |  |
| 7:23.212 | אלא |  |
| 7:23.213 | דעז"א |  |
| 7:23.214 | ועזא"ל |  |
| 7:23.215 | הוו |  |
| 7:23.216 | מקטרגי |  |
| 7:23.217 | עליה |  |
| 7:23.218 | בזמנא |  |
| 7:23.219 | דאמר |  |
| 7:23.220 | שכינתא |  |
| 7:23.221 | לקב"ה |  |
| 7:23.222 | נעשה |  |
| 7:23.223 | אדם | , |
| 7:23.224 | אמרו |  |
| 7:23.225 | מה |  |
| 7:23.226 | אדם |  |
| 7:23.227 | ותדעהו | , |
| 7:23.228 | מה |  |
| 7:23.229 | את |  |
| 7:23.230 | בעי |  |
| 7:23.231 | למברי |  |
| 7:23.232 | אדם | . |
| 7:23.233 | ותדעהו |  |
| 7:23.234 | דעתיד |  |
| 7:23.235 | למחטי |  |
| 7:23.236 | באתתא |  |
| 7:23.237 | דיליה |  |
| 7:23.238 | דאיהי |  |
| 7:23.239 | חשך | , |
| 7:23.240 | דאור |  |
| 7:23.241 | איהו |  |
| 7:23.242 | דכורא | , |
| 7:23.243 | וחשך |  |
| 7:23.244 | נוקבא |  |
| 7:23.245 | שמאלא |  |
| 7:23.246 | דבריאה | , |
| 7:23.247 | בההוא |  |
| 7:23.248 | זמנא |  |
| 7:23.249 | אמר |  |
| 7:23.250 | שכינתא |  |
| 7:23.251 | לון | , |
| 7:23.252 | בהאי |  |
| 7:23.253 | דאתון |  |
| 7:23.254 | מקטרגין |  |
| 7:23.255 | אתון |  |
| 7:23.256 | עתידין |  |
| 7:23.257 | למנפל | . |
| 7:23.258 | כדכתיב |  |
| 7:23.259 | ויראו |  |
| 7:23.260 | בני |  |
| 7:23.261 | האלהים |  |
| 7:23.262 | את |  |
| 7:23.263 | בנות |  |
| 7:23.264 | האדם |  |
| 7:23.265 | כי |  |
| 7:23.266 | טבת |  |
| 7:23.267 | הנה | . |
| 7:23.268 | חשקו |  |
| 7:23.269 | בהון |  |
| 7:23.270 | וטעו |  |
| 7:23.271 | בהון | . |
| 7:23.272 | ואפיל |  |
| 7:23.273 | להון |  |
| 7:23.274 | שכינתא |  |
| 7:23.275 | מקדושה |  |
| 7:23.276 | דילהון | . |
| 7:23.277 | אמרו |  |
| 7:23.278 | חבריא |  |
| 7:23.279 | רבי |  |
| 7:23.280 | רבי |  |
| 7:23.281 | אי |  |
| 7:23.282 | הכי |  |
| 7:23.283 | עזא |  |
| 7:23.284 | ועזאל |  |
| 7:23.285 | לא |  |
| 7:23.286 | הוו |  |
| 7:23.287 | משקרי |  |
| 7:23.288 | במלוליהו | . |
| 7:23.289 | דודאי |  |
| 7:23.290 | בנוקבא |  |
| 7:23.291 | עתיד |  |
| 7:23.292 | אדם |  |
| 7:23.293 | למחטי | . |
| 7:23.294 | אמר |  |
| 7:23.295 | להו |  |
| 7:23.296 | הכי |  |
| 7:23.297 | אמרה |  |
| 7:23.298 | שכינתא | . |
| 7:23.299 | אתון |  |
| 7:23.300 | דאזדמנתן |  |
| 7:23.301 | לקטרגא |  |
| 7:23.302 | קדמי |  |
| 7:23.303 | יתיר |  |
| 7:23.304 | מחילי |  |
| 7:23.305 | מרומא | . |
| 7:23.306 | אי |  |
| 7:23.307 | אתון |  |
| 7:23.308 | חזיתון |  |
| 7:23.309 | שפירין |  |
| 7:23.310 | מאדם |  |
| 7:23.311 | בעובדיכון |  |
| 7:23.312 | יאות |  |
| 7:23.313 | לכון |  |
| 7:23.314 | לקטרגא |  |
| 7:23.315 | עליה | , |
| 7:23.316 | אבל |  |
| 7:23.317 | איהו |  |
| 7:23.318 | עתיד |  |
| 7:23.319 | למחטי |  |
| 7:23.320 | באתתא |  |
| 7:23.321 | חדא | . |
| 7:23.322 | אתון |  |
| 7:23.323 | בנשי |  |
| 7:23.324 | סגיאי |  |
| 7:23.325 | חבתון |  |
| 7:23.326 | יתיר |  |
| 7:23.327 | מבני |  |
| 7:23.328 | נשא |  |
| 7:23.329 | דכתיב |  |
| 7:23.330 | ויראו |  |
| 7:23.331 | בני |  |
| 7:23.332 | האלהים |  |
| 7:23.333 | את |  |
| 7:23.334 | בנות |  |
| 7:23.335 | האדם |  |
| 7:23.336 | כי |  |
| 7:23.337 | טבת |  |
| 7:23.338 | הנה |  |
| 7:23.339 | וכו | '. |
| 7:23.340 | גם |  |
| 7:23.341 | בדף |  |
| 7:23.342 | נ"ה |  |
| 7:23.343 | חזר |  |
| 7:23.344 | לשנותה |  |
| 7:23.345 | בשם |  |
| 7:23.346 | ר | ' |
| 7:23.347 | יצחק | . |
| 7:23.348 | ובדף |  |
| 7:23.349 | נ"ח |  |
| 7:23.350 | חזר |  |
| 7:23.351 | ושלשה |  |
| 7:23.352 | בשם |  |
| 7:23.353 | ר"י |  |
| 7:23.354 | ור"ח | . |
| 7:23.355 | הרי |  |
| 7:23.356 | שמחבר |  |
| 7:23.357 | הזהר |  |
| 7:23.358 | מחבב |  |
| 7:23.359 | סברא |  |
| 7:23.360 | זו |  |
| 7:23.361 | ולפיכך |  |
| 7:23.362 | אמרה |  |
| 7:23.363 | ושנאה |  |
| 7:23.364 | ושילשה |  |
| 7:23.365 | לרוב |  |
| 7:23.366 | חבתה | . |
| 7:23.367 | והרי |  |
| 7:23.368 | בבראשית |  |
| 7:23.369 | רבה |  |
| 7:23.370 | מבורא |  |
| 7:23.371 | שרשב"י |  |
| 7:23.372 | שונא |  |
| 7:23.373 | סברא |  |
| 7:23.374 | זו |  |
| 7:23.375 | עד |  |
| 7:23.376 | שהיה |  |
| 7:23.377 | מקלל |  |
| 7:23.378 | את |  |
| 7:23.379 | האומר |  |
| 7:23.380 | שהם |  |
| 7:23.381 | בני |  |
| 7:23.382 | אלהיא | . |
| 7:23.383 | אלא |  |
| 7:23.384 | ודאי |  |
| 7:23.385 | לא |  |
| 7:23.386 | מפי |  |
| 7:23.387 | רשב"י |  |
| 7:23.388 | יצאו |  |
| 7:23.389 | דברים |  |
| 7:23.390 | אלו |  |
| 7:23.391 | האמורים |  |
| 7:23.392 | בזהר | . |
| 7:23.393 | ובשקר |  |
| 7:23.394 | יחסוהו |  |
| 7:23.395 | לרשב"י | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

