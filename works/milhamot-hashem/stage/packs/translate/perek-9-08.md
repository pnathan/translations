# Translation pack: milhamot-hashem perek-9-08 (Milchamot ha-Shem 9:10–9:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-08`.

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
- ב"ע: "in Aramaic" (בארמית) where it is set against בערבי; the other readings in
  `readings`.
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


## Milchamot ha-Shem 9:10

Hebrew (Wikisource): קכז) לבי לחוקקי ישראל המתנבים בעם להבין ולהשכיל בתורת ה' בחכמה ותבונה. ובדעת ישרה על פי קבלת רז"ל בעלי המשנה והתלמוד והפוסקים הדולים ומשקים לכל צמא ממקור מים חיים, משה קבל תורה מסיני ומסרה ליהושע וכו' והם מלאו אחר ה' ותורתו הקדושה שלא לפנות אל האלילים, איך בכל מצותיה של תורה האריכו הרחיבו לבאר חילוקי דיניהם בכמה פרטים, וחומרות יתירות ממה שנזכר בתלמוד! ובדיני יחוד השי"ת שהוא העיקר הגדול, העלימו עיניהם ועזבוהו ולא דברו מאומה אפי' בפרט אחד מכל דיני יחודו ית'! ולמה זה ועל מה זה הניחוהו כעיר פרוצה אין חומה, והוא נוהג בכל זמן ובכל מקום בין בארץ בין בחוצא לארץ, בין בפני הבית בין שלא בפני הבית, כאלו נשתכחה מישראל תורת יחוד השי"ת היוצאת מדברי התלמוד בבלי וירושלמי ומדרשי רז"ל, אחרי מות ר' בחיי בעל חובת הלבבות ורבי יהודה הלוי והרס"ג והרמב"ם והסמ"ג ובעל העיקרים וזולתם! עד שבעבור זה נפלו קצת אחרונים בפח הנסיון אשר נסה ה' אלהינו אותנו על ידי המלך העכו"ם בשלחו ספר הזהר לחכמי טוליטולא, והמה שמחו בו שמחה גדולה, כי חפצו להוציא ישן תורת המשנה והתלמוד, מפני חדש תורת הזהר המבזה את המשנה והתלמוד כמבואר לעיל, ולא נתנו אל לבם למנוע מלעבוד אלהים חדשים, מקרבו באו, אשר הזהירנו ה' אלהינו לבלתי עבוד זולתו, אפי' נקראהו בשמו כמעשה דורו של אנוש. וחכמי טוליטולא לא עשו כן אדרבא שלחו מתנות רבות למלך העכו"ם בעבורו והגדילו והחשיבו דברי הפילסוף מחבר הזהר המתנבא בשקר בשם רשב"י שנגלה לו עתיק, יותר ממ"ח נביאים ושבע נביאות שנתנבאו להם לישראל, שמ"ח נביאים ושבע נביאות לא פיחתו ולא הותירו על מה שכתוב בתורה אפי' אות אחת חוץ ממקרא מגלה, ומחבר הזהר הוסיף אמונת אלהים אחרים חדשים אשר לא שערום אבותיהם. ואמר שמי שלא יעסוק בתורתו החדשה לית ליה רשו למיעל קמי מריה. ובספר אצהאר אלחק לאחד מחכמי הערביים בטענתו נגד אמונת השלוש שהיא אמונת הנצרים כתב בדף רפ"א ורפ"ב וז"ל. ועלמא אליהוד מן ודן מוסי' אלי' זמאננא הדא לא יעתרפון אתתלית (תלת קשרי דמהימנותא) וצאהר אן דאת אללה וצפאתה אלכמאליה קדימא גיה מתגיירה אזלא' ואבדא'. פלו כאן אלתתלית חקא', לכאן אלואגב עלי' מוסי' עליה אסלאם ואנביא בני אסראיל אן יביינוה חק אלביאן, פאלעגב כל אלעגב אן תכון אלשריעה אלמוסויה אלתי כאנת ואנבת אלאטאעה בגמיע אלאנביא אלי' עהד עיסי' כאלה עם ביאן הדה אלעקידה אלתי הי מדאר אלנגאה (היא מרכז ההצלה והוא מאמר הזהר דלית ליה רשו למיעל קמי מריה) עלי' זעם אהל אלתתלית ולא ימכן נגאת אחדכדונהא נבייא כאן או גיר נבי, ולא יביין מוסי' ע"ס. ולא נבי מן אלאנביא אלאסראלייה הדה אלעקידה בביאן ואצח בחית תפהם מנה הדה אלעקידה צראחה, ולא יבקש פיהא שך ויביין מוסי' אלאחכאם אלתי הי ענד מקרס אהל אלתתלית צעיפה נאקצה גדא באלתשריח אלאתם. ויכררהא מרה בעד אולי וכרה בעד אכרא, ויאכד ען מחאפצהתא תאכידא בלינא, ויוגב אלקתל למן לם יעתקדהא עכ"ל. (העתק ועיין בסוף הספר) וטענתו זאת בעצמה אמיתית צודקת על הקבלה החדשה, ומדבריו מוכח שלא פשטה עדיין אמונת הקבלה החדשה בישראל.

| id | word | punctuation after |
|---|---|---|
| 9:10.1 | קכז | ) |
| 9:10.2 | לבי |  |
| 9:10.3 | לחוקקי |  |
| 9:10.4 | ישראל |  |
| 9:10.5 | המתנבים |  |
| 9:10.6 | בעם |  |
| 9:10.7 | להבין |  |
| 9:10.8 | ולהשכיל |  |
| 9:10.9 | בתורת |  |
| 9:10.10 | ה | ' |
| 9:10.11 | בחכמה |  |
| 9:10.12 | ותבונה | . |
| 9:10.13 | ובדעת |  |
| 9:10.14 | ישרה |  |
| 9:10.15 | על |  |
| 9:10.16 | פי |  |
| 9:10.17 | קבלת |  |
| 9:10.18 | רז"ל |  |
| 9:10.19 | בעלי |  |
| 9:10.20 | המשנה |  |
| 9:10.21 | והתלמוד |  |
| 9:10.22 | והפוסקים |  |
| 9:10.23 | הדולים |  |
| 9:10.24 | ומשקים |  |
| 9:10.25 | לכל |  |
| 9:10.26 | צמא |  |
| 9:10.27 | ממקור |  |
| 9:10.28 | מים |  |
| 9:10.29 | חיים | , |
| 9:10.30 | משה |  |
| 9:10.31 | קבל |  |
| 9:10.32 | תורה |  |
| 9:10.33 | מסיני |  |
| 9:10.34 | ומסרה |  |
| 9:10.35 | ליהושע |  |
| 9:10.36 | וכו | ' |
| 9:10.37 | והם |  |
| 9:10.38 | מלאו |  |
| 9:10.39 | אחר |  |
| 9:10.40 | ה | ' |
| 9:10.41 | ותורתו |  |
| 9:10.42 | הקדושה |  |
| 9:10.43 | שלא |  |
| 9:10.44 | לפנות |  |
| 9:10.45 | אל |  |
| 9:10.46 | האלילים | , |
| 9:10.47 | איך |  |
| 9:10.48 | בכל |  |
| 9:10.49 | מצותיה |  |
| 9:10.50 | של |  |
| 9:10.51 | תורה |  |
| 9:10.52 | האריכו |  |
| 9:10.53 | הרחיבו |  |
| 9:10.54 | לבאר |  |
| 9:10.55 | חילוקי |  |
| 9:10.56 | דיניהם |  |
| 9:10.57 | בכמה |  |
| 9:10.58 | פרטים | , |
| 9:10.59 | וחומרות |  |
| 9:10.60 | יתירות |  |
| 9:10.61 | ממה |  |
| 9:10.62 | שנזכר |  |
| 9:10.63 | בתלמוד | ! |
| 9:10.64 | ובדיני |  |
| 9:10.65 | יחוד |  |
| 9:10.66 | השי"ת |  |
| 9:10.67 | שהוא |  |
| 9:10.68 | העיקר |  |
| 9:10.69 | הגדול | , |
| 9:10.70 | העלימו |  |
| 9:10.71 | עיניהם |  |
| 9:10.72 | ועזבוהו |  |
| 9:10.73 | ולא |  |
| 9:10.74 | דברו |  |
| 9:10.75 | מאומה |  |
| 9:10.76 | אפי | ' |
| 9:10.77 | בפרט |  |
| 9:10.78 | אחד |  |
| 9:10.79 | מכל |  |
| 9:10.80 | דיני |  |
| 9:10.81 | יחודו |  |
| 9:10.82 | ית | '! |
| 9:10.83 | ולמה |  |
| 9:10.84 | זה |  |
| 9:10.85 | ועל |  |
| 9:10.86 | מה |  |
| 9:10.87 | זה |  |
| 9:10.88 | הניחוהו |  |
| 9:10.89 | כעיר |  |
| 9:10.90 | פרוצה |  |
| 9:10.91 | אין |  |
| 9:10.92 | חומה | , |
| 9:10.93 | והוא |  |
| 9:10.94 | נוהג |  |
| 9:10.95 | בכל |  |
| 9:10.96 | זמן |  |
| 9:10.97 | ובכל |  |
| 9:10.98 | מקום |  |
| 9:10.99 | בין |  |
| 9:10.100 | בארץ |  |
| 9:10.101 | בין |  |
| 9:10.102 | בחוצא |  |
| 9:10.103 | לארץ | , |
| 9:10.104 | בין |  |
| 9:10.105 | בפני |  |
| 9:10.106 | הבית |  |
| 9:10.107 | בין |  |
| 9:10.108 | שלא |  |
| 9:10.109 | בפני |  |
| 9:10.110 | הבית | , |
| 9:10.111 | כאלו |  |
| 9:10.112 | נשתכחה |  |
| 9:10.113 | מישראל |  |
| 9:10.114 | תורת |  |
| 9:10.115 | יחוד |  |
| 9:10.116 | השי"ת |  |
| 9:10.117 | היוצאת |  |
| 9:10.118 | מדברי |  |
| 9:10.119 | התלמוד |  |
| 9:10.120 | בבלי |  |
| 9:10.121 | וירושלמי |  |
| 9:10.122 | ומדרשי |  |
| 9:10.123 | רז"ל | , |
| 9:10.124 | אחרי |  |
| 9:10.125 | מות |  |
| 9:10.126 | ר | ' |
| 9:10.127 | בחיי |  |
| 9:10.128 | בעל |  |
| 9:10.129 | חובת |  |
| 9:10.130 | הלבבות |  |
| 9:10.131 | ורבי |  |
| 9:10.132 | יהודה |  |
| 9:10.133 | הלוי |  |
| 9:10.134 | והרס"ג |  |
| 9:10.135 | והרמב"ם |  |
| 9:10.136 | והסמ"ג |  |
| 9:10.137 | ובעל |  |
| 9:10.138 | העיקרים |  |
| 9:10.139 | וזולתם | ! |
| 9:10.140 | עד |  |
| 9:10.141 | שבעבור |  |
| 9:10.142 | זה |  |
| 9:10.143 | נפלו |  |
| 9:10.144 | קצת |  |
| 9:10.145 | אחרונים |  |
| 9:10.146 | בפח |  |
| 9:10.147 | הנסיון |  |
| 9:10.148 | אשר |  |
| 9:10.149 | נסה |  |
| 9:10.150 | ה | ' |
| 9:10.151 | אלהינו |  |
| 9:10.152 | אותנו |  |
| 9:10.153 | על |  |
| 9:10.154 | ידי |  |
| 9:10.155 | המלך |  |
| 9:10.156 | העכו"ם |  |
| 9:10.157 | בשלחו |  |
| 9:10.158 | ספר |  |
| 9:10.159 | הזהר |  |
| 9:10.160 | לחכמי |  |
| 9:10.161 | טוליטולא | , |
| 9:10.162 | והמה |  |
| 9:10.163 | שמחו |  |
| 9:10.164 | בו |  |
| 9:10.165 | שמחה |  |
| 9:10.166 | גדולה | , |
| 9:10.167 | כי |  |
| 9:10.168 | חפצו |  |
| 9:10.169 | להוציא |  |
| 9:10.170 | ישן |  |
| 9:10.171 | תורת |  |
| 9:10.172 | המשנה |  |
| 9:10.173 | והתלמוד | , |
| 9:10.174 | מפני |  |
| 9:10.175 | חדש |  |
| 9:10.176 | תורת |  |
| 9:10.177 | הזהר |  |
| 9:10.178 | המבזה |  |
| 9:10.179 | את |  |
| 9:10.180 | המשנה |  |
| 9:10.181 | והתלמוד |  |
| 9:10.182 | כמבואר |  |
| 9:10.183 | לעיל | , |
| 9:10.184 | ולא |  |
| 9:10.185 | נתנו |  |
| 9:10.186 | אל |  |
| 9:10.187 | לבם |  |
| 9:10.188 | למנוע |  |
| 9:10.189 | מלעבוד |  |
| 9:10.190 | אלהים |  |
| 9:10.191 | חדשים | , |
| 9:10.192 | מקרבו |  |
| 9:10.193 | באו | , |
| 9:10.194 | אשר |  |
| 9:10.195 | הזהירנו |  |
| 9:10.196 | ה | ' |
| 9:10.197 | אלהינו |  |
| 9:10.198 | לבלתי |  |
| 9:10.199 | עבוד |  |
| 9:10.200 | זולתו | , |
| 9:10.201 | אפי | ' |
| 9:10.202 | נקראהו |  |
| 9:10.203 | בשמו |  |
| 9:10.204 | כמעשה |  |
| 9:10.205 | דורו |  |
| 9:10.206 | של |  |
| 9:10.207 | אנוש | . |
| 9:10.208 | וחכמי |  |
| 9:10.209 | טוליטולא |  |
| 9:10.210 | לא |  |
| 9:10.211 | עשו |  |
| 9:10.212 | כן |  |
| 9:10.213 | אדרבא |  |
| 9:10.214 | שלחו |  |
| 9:10.215 | מתנות |  |
| 9:10.216 | רבות |  |
| 9:10.217 | למלך |  |
| 9:10.218 | העכו"ם |  |
| 9:10.219 | בעבורו |  |
| 9:10.220 | והגדילו |  |
| 9:10.221 | והחשיבו |  |
| 9:10.222 | דברי |  |
| 9:10.223 | הפילסוף |  |
| 9:10.224 | מחבר |  |
| 9:10.225 | הזהר |  |
| 9:10.226 | המתנבא |  |
| 9:10.227 | בשקר |  |
| 9:10.228 | בשם |  |
| 9:10.229 | רשב"י |  |
| 9:10.230 | שנגלה |  |
| 9:10.231 | לו |  |
| 9:10.232 | עתיק | , |
| 9:10.233 | יותר |  |
| 9:10.234 | ממ"ח |  |
| 9:10.235 | נביאים |  |
| 9:10.236 | ושבע |  |
| 9:10.237 | נביאות |  |
| 9:10.238 | שנתנבאו |  |
| 9:10.239 | להם |  |
| 9:10.240 | לישראל | , |
| 9:10.241 | שמ"ח |  |
| 9:10.242 | נביאים |  |
| 9:10.243 | ושבע |  |
| 9:10.244 | נביאות |  |
| 9:10.245 | לא |  |
| 9:10.246 | פיחתו |  |
| 9:10.247 | ולא |  |
| 9:10.248 | הותירו |  |
| 9:10.249 | על |  |
| 9:10.250 | מה |  |
| 9:10.251 | שכתוב |  |
| 9:10.252 | בתורה |  |
| 9:10.253 | אפי | ' |
| 9:10.254 | אות |  |
| 9:10.255 | אחת |  |
| 9:10.256 | חוץ |  |
| 9:10.257 | ממקרא |  |
| 9:10.258 | מגלה | , |
| 9:10.259 | ומחבר |  |
| 9:10.260 | הזהר |  |
| 9:10.261 | הוסיף |  |
| 9:10.262 | אמונת |  |
| 9:10.263 | אלהים |  |
| 9:10.264 | אחרים |  |
| 9:10.265 | חדשים |  |
| 9:10.266 | אשר |  |
| 9:10.267 | לא |  |
| 9:10.268 | שערום |  |
| 9:10.269 | אבותיהם | . |
| 9:10.270 | ואמר |  |
| 9:10.271 | שמי |  |
| 9:10.272 | שלא |  |
| 9:10.273 | יעסוק |  |
| 9:10.274 | בתורתו |  |
| 9:10.275 | החדשה |  |
| 9:10.276 | לית |  |
| 9:10.277 | ליה |  |
| 9:10.278 | רשו |  |
| 9:10.279 | למיעל |  |
| 9:10.280 | קמי |  |
| 9:10.281 | מריה | . |
| 9:10.282 | ובספר |  |
| 9:10.283 | אצהאר |  |
| 9:10.284 | אלחק |  |
| 9:10.285 | לאחד |  |
| 9:10.286 | מחכמי |  |
| 9:10.287 | הערביים |  |
| 9:10.288 | בטענתו |  |
| 9:10.289 | נגד |  |
| 9:10.290 | אמונת |  |
| 9:10.291 | השלוש |  |
| 9:10.292 | שהיא |  |
| 9:10.293 | אמונת |  |
| 9:10.294 | הנצרים |  |
| 9:10.295 | כתב |  |
| 9:10.296 | בדף |  |
| 9:10.297 | רפ"א |  |
| 9:10.298 | ורפ"ב |  |
| 9:10.299 | וז"ל | . |
| 9:10.300 | ועלמא |  |
| 9:10.301 | אליהוד |  |
| 9:10.302 | מן |  |
| 9:10.303 | ודן |  |
| 9:10.304 | מוסי | ' |
| 9:10.305 | אלי | ' |
| 9:10.306 | זמאננא |  |
| 9:10.307 | הדא |  |
| 9:10.308 | לא |  |
| 9:10.309 | יעתרפון |  |
| 9:10.310 | אתתלית |  |
| 9:10.311 | תלת |  |
| 9:10.312 | קשרי |  |
| 9:10.313 | דמהימנותא | ) |
| 9:10.314 | וצאהר |  |
| 9:10.315 | אן |  |
| 9:10.316 | דאת |  |
| 9:10.317 | אללה |  |
| 9:10.318 | וצפאתה |  |
| 9:10.319 | אלכמאליה |  |
| 9:10.320 | קדימא |  |
| 9:10.321 | גיה |  |
| 9:10.322 | מתגיירה |  |
| 9:10.323 | אזלא | ' |
| 9:10.324 | ואבדא | '. |
| 9:10.325 | פלו |  |
| 9:10.326 | כאן |  |
| 9:10.327 | אלתתלית |  |
| 9:10.328 | חקא | ', |
| 9:10.329 | לכאן |  |
| 9:10.330 | אלואגב |  |
| 9:10.331 | עלי | ' |
| 9:10.332 | מוסי | ' |
| 9:10.333 | עליה |  |
| 9:10.334 | אסלאם |  |
| 9:10.335 | ואנביא |  |
| 9:10.336 | בני |  |
| 9:10.337 | אסראיל |  |
| 9:10.338 | אן |  |
| 9:10.339 | יביינוה |  |
| 9:10.340 | חק |  |
| 9:10.341 | אלביאן | , |
| 9:10.342 | פאלעגב |  |
| 9:10.343 | כל |  |
| 9:10.344 | אלעגב |  |
| 9:10.345 | אן |  |
| 9:10.346 | תכון |  |
| 9:10.347 | אלשריעה |  |
| 9:10.348 | אלמוסויה |  |
| 9:10.349 | אלתי |  |
| 9:10.350 | כאנת |  |
| 9:10.351 | ואנבת |  |
| 9:10.352 | אלאטאעה |  |
| 9:10.353 | בגמיע |  |
| 9:10.354 | אלאנביא |  |
| 9:10.355 | אלי | ' |
| 9:10.356 | עהד |  |
| 9:10.357 | עיסי | ' |
| 9:10.358 | כאלה |  |
| 9:10.359 | עם |  |
| 9:10.360 | ביאן |  |
| 9:10.361 | הדה |  |
| 9:10.362 | אלעקידה |  |
| 9:10.363 | אלתי |  |
| 9:10.364 | הי |  |
| 9:10.365 | מדאר |  |
| 9:10.366 | אלנגאה |  |
| 9:10.367 | היא |  |
| 9:10.368 | מרכז |  |
| 9:10.369 | ההצלה |  |
| 9:10.370 | והוא |  |
| 9:10.371 | מאמר |  |
| 9:10.372 | הזהר |  |
| 9:10.373 | דלית |  |
| 9:10.374 | ליה |  |
| 9:10.375 | רשו |  |
| 9:10.376 | למיעל |  |
| 9:10.377 | קמי |  |
| 9:10.378 | מריה | ) |
| 9:10.379 | עלי | ' |
| 9:10.380 | זעם |  |
| 9:10.381 | אהל |  |
| 9:10.382 | אלתתלית |  |
| 9:10.383 | ולא |  |
| 9:10.384 | ימכן |  |
| 9:10.385 | נגאת |  |
| 9:10.386 | אחדכדונהא |  |
| 9:10.387 | נבייא |  |
| 9:10.388 | כאן |  |
| 9:10.389 | או |  |
| 9:10.390 | גיר |  |
| 9:10.391 | נבי | , |
| 9:10.392 | ולא |  |
| 9:10.393 | יביין |  |
| 9:10.394 | מוסי | ' |
| 9:10.395 | ע"ס | . |
| 9:10.396 | ולא |  |
| 9:10.397 | נבי |  |
| 9:10.398 | מן |  |
| 9:10.399 | אלאנביא |  |
| 9:10.400 | אלאסראלייה |  |
| 9:10.401 | הדה |  |
| 9:10.402 | אלעקידה |  |
| 9:10.403 | בביאן |  |
| 9:10.404 | ואצח |  |
| 9:10.405 | בחית |  |
| 9:10.406 | תפהם |  |
| 9:10.407 | מנה |  |
| 9:10.408 | הדה |  |
| 9:10.409 | אלעקידה |  |
| 9:10.410 | צראחה | , |
| 9:10.411 | ולא |  |
| 9:10.412 | יבקש |  |
| 9:10.413 | פיהא |  |
| 9:10.414 | שך |  |
| 9:10.415 | ויביין |  |
| 9:10.416 | מוסי | ' |
| 9:10.417 | אלאחכאם |  |
| 9:10.418 | אלתי |  |
| 9:10.419 | הי |  |
| 9:10.420 | ענד |  |
| 9:10.421 | מקרס |  |
| 9:10.422 | אהל |  |
| 9:10.423 | אלתתלית |  |
| 9:10.424 | צעיפה |  |
| 9:10.425 | נאקצה |  |
| 9:10.426 | גדא |  |
| 9:10.427 | באלתשריח |  |
| 9:10.428 | אלאתם | . |
| 9:10.429 | ויכררהא |  |
| 9:10.430 | מרה |  |
| 9:10.431 | בעד |  |
| 9:10.432 | אולי |  |
| 9:10.433 | וכרה |  |
| 9:10.434 | בעד |  |
| 9:10.435 | אכרא | , |
| 9:10.436 | ויאכד |  |
| 9:10.437 | ען |  |
| 9:10.438 | מחאפצהתא |  |
| 9:10.439 | תאכידא |  |
| 9:10.440 | בלינא | , |
| 9:10.441 | ויוגב |  |
| 9:10.442 | אלקתל |  |
| 9:10.443 | למן |  |
| 9:10.444 | לם |  |
| 9:10.445 | יעתקדהא |  |
| 9:10.446 | עכ"ל | . |
| 9:10.447 | העתק |  |
| 9:10.448 | ועיין |  |
| 9:10.449 | בסוף |  |
| 9:10.450 | הספר | ) |
| 9:10.451 | וטענתו |  |
| 9:10.452 | זאת |  |
| 9:10.453 | בעצמה |  |
| 9:10.454 | אמיתית |  |
| 9:10.455 | צודקת |  |
| 9:10.456 | על |  |
| 9:10.457 | הקבלה |  |
| 9:10.458 | החדשה | , |
| 9:10.459 | ומדבריו |  |
| 9:10.460 | מוכח |  |
| 9:10.461 | שלא |  |
| 9:10.462 | פשטה |  |
| 9:10.463 | עדיין |  |
| 9:10.464 | אמונת |  |
| 9:10.465 | הקבלה |  |
| 9:10.466 | החדשה |  |
| 9:10.467 | בישראל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

