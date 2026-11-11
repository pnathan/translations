# Review pack: milhamot-hashem perek-3-09 (Milchamot ha-Shem 3:14–3:15)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-09`.

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
- מהר"י + surname, like other titles formed as acronyms (Rambam, Mahari): keep
  the acronym as the author's name for the man ("Mahari Lopes"), with a note
  expanding it (morenu ha-Rav Rabbi Y.) at its first occurrence in the unit.
  Same for הרב מקדש מלך: "the Rav Mikdash Melekh", noting who is meant.
- A variant spelling of a word the text does not gloss (ברותא for ברייתא):
  transliterate it ("berota"), with a note.
- יש מאין / יש מיש: translate, "something from nothing" / "something from
  something"; no transliteration.
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
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
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".
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


## Milchamot ha-Shem 3:14

Hebrew (Wikisource): יד) ראה והבן ידידנו קורא נעלם איך כמה"ר תם הנז"ל צוח ככרוכיא על העוסקים בקבלה החדשה, כי פרצו גדריה של תורה, וערו ערו יסודותיה ושרשיה ביחוד השי"ת כמו שיתבאר להלן, ותחת אשר חשבו להתקרב לשי"ת תעו מן הדרך ובאו והאמין באלהים רבים אחרים ולעבדם! וכדברי ר' תם ן' יחיא כ' גם כן מהר"י אלבו בס' העיקרים פרק כ"ח משני. ועוד האריך בספר רביד הזהב והביא ג"כ דברי הר' הגדול מהרש"ל בתשובותיו ובחבור ים של שלמה על המתיהרים בחידושים שלומדים בקבלה, וכתב דמאחר שנחתם התלמוד אין לנו להחמיר נגד הגמרא והוי כמו מינות כמבואר בחק יעקב סי' תפ"ט ס"ק י"א באורך. וכל זה הוא לפי ההנחה והיחוס שייחסו הזהר לרשב"י התנא, אך אחרי הבקורת שבקרו וחקרו חכמי לב אחריו מבטן מי יצאו הדברים הנזכרים בו ואימתי יצא החבור הנזכר ומיד מי יצא. ונתבאר באר היטב כי הוא חיבור אחד האחרונים ולא יצאו אותן הרעיונים מרשב"י התנא ע"ה כלל, כי נמצא בו כמה דברים בעיקרי אמונת תורתינו הקדושה היפך דברי רשב"י בתלמוד ובמדרשי רז"ל המפורסמים בישראל, ודאי שאין לסמוך עליו בשום דבר הנוגע למנהגי ישראל שנהגו בהם קדמונינו. ולא יפה עשו חכמי ימי הבינים ששינו עלינו כמה דברים בסידורי התפלות על פי הזהר והמקובלים, כי שינויים ההם יש מהם מביאים לידי מינות להאמין ברשויות הרבה ולעבוד אלהים אחרים כמו שיתבאר להלן.

| id | word | punctuation after |
|---|---|---|
| 3:14.1 | יד | ) |
| 3:14.2 | ראה |  |
| 3:14.3 | והבן |  |
| 3:14.4 | ידידנו |  |
| 3:14.5 | קורא |  |
| 3:14.6 | נעלם |  |
| 3:14.7 | איך |  |
| 3:14.8 | כמה"ר |  |
| 3:14.9 | תם |  |
| 3:14.10 | הנז"ל |  |
| 3:14.11 | צוח |  |
| 3:14.12 | ככרוכיא |  |
| 3:14.13 | על |  |
| 3:14.14 | העוסקים |  |
| 3:14.15 | בקבלה |  |
| 3:14.16 | החדשה | , |
| 3:14.17 | כי |  |
| 3:14.18 | פרצו |  |
| 3:14.19 | גדריה |  |
| 3:14.20 | של |  |
| 3:14.21 | תורה | , |
| 3:14.22 | וערו |  |
| 3:14.23 | ערו |  |
| 3:14.24 | יסודותיה |  |
| 3:14.25 | ושרשיה |  |
| 3:14.26 | ביחוד |  |
| 3:14.27 | השי"ת |  |
| 3:14.28 | כמו |  |
| 3:14.29 | שיתבאר |  |
| 3:14.30 | להלן | , |
| 3:14.31 | ותחת |  |
| 3:14.32 | אשר |  |
| 3:14.33 | חשבו |  |
| 3:14.34 | להתקרב |  |
| 3:14.35 | לשי"ת |  |
| 3:14.36 | תעו |  |
| 3:14.37 | מן |  |
| 3:14.38 | הדרך |  |
| 3:14.39 | ובאו |  |
| 3:14.40 | והאמין |  |
| 3:14.41 | באלהים |  |
| 3:14.42 | רבים |  |
| 3:14.43 | אחרים |  |
| 3:14.44 | ולעבדם | ! |
| 3:14.45 | וכדברי |  |
| 3:14.46 | ר | ' |
| 3:14.47 | תם |  |
| 3:14.48 | ן | ' |
| 3:14.49 | יחיא |  |
| 3:14.50 | כ | ' |
| 3:14.51 | גם |  |
| 3:14.52 | כן |  |
| 3:14.53 | מהר"י |  |
| 3:14.54 | אלבו |  |
| 3:14.55 | בס | ' |
| 3:14.56 | העיקרים |  |
| 3:14.57 | פרק |  |
| 3:14.58 | כ"ח |  |
| 3:14.59 | משני | . |
| 3:14.60 | ועוד |  |
| 3:14.61 | האריך |  |
| 3:14.62 | בספר |  |
| 3:14.63 | רביד |  |
| 3:14.64 | הזהב |  |
| 3:14.65 | והביא |  |
| 3:14.66 | ג"כ |  |
| 3:14.67 | דברי |  |
| 3:14.68 | הר | ' |
| 3:14.69 | הגדול |  |
| 3:14.70 | מהרש"ל |  |
| 3:14.71 | בתשובותיו |  |
| 3:14.72 | ובחבור |  |
| 3:14.73 | ים |  |
| 3:14.74 | של |  |
| 3:14.75 | שלמה |  |
| 3:14.76 | על |  |
| 3:14.77 | המתיהרים |  |
| 3:14.78 | בחידושים |  |
| 3:14.79 | שלומדים |  |
| 3:14.80 | בקבלה | , |
| 3:14.81 | וכתב |  |
| 3:14.82 | דמאחר |  |
| 3:14.83 | שנחתם |  |
| 3:14.84 | התלמוד |  |
| 3:14.85 | אין |  |
| 3:14.86 | לנו |  |
| 3:14.87 | להחמיר |  |
| 3:14.88 | נגד |  |
| 3:14.89 | הגמרא |  |
| 3:14.90 | והוי |  |
| 3:14.91 | כמו |  |
| 3:14.92 | מינות |  |
| 3:14.93 | כמבואר |  |
| 3:14.94 | בחק |  |
| 3:14.95 | יעקב |  |
| 3:14.96 | סי | ' |
| 3:14.97 | תפ"ט |  |
| 3:14.98 | ס"ק |  |
| 3:14.99 | י"א |  |
| 3:14.100 | באורך | . |
| 3:14.101 | וכל |  |
| 3:14.102 | זה |  |
| 3:14.103 | הוא |  |
| 3:14.104 | לפי |  |
| 3:14.105 | ההנחה |  |
| 3:14.106 | והיחוס |  |
| 3:14.107 | שייחסו |  |
| 3:14.108 | הזהר |  |
| 3:14.109 | לרשב"י |  |
| 3:14.110 | התנא | , |
| 3:14.111 | אך |  |
| 3:14.112 | אחרי |  |
| 3:14.113 | הבקורת |  |
| 3:14.114 | שבקרו |  |
| 3:14.115 | וחקרו |  |
| 3:14.116 | חכמי |  |
| 3:14.117 | לב |  |
| 3:14.118 | אחריו |  |
| 3:14.119 | מבטן |  |
| 3:14.120 | מי |  |
| 3:14.121 | יצאו |  |
| 3:14.122 | הדברים |  |
| 3:14.123 | הנזכרים |  |
| 3:14.124 | בו |  |
| 3:14.125 | ואימתי |  |
| 3:14.126 | יצא |  |
| 3:14.127 | החבור |  |
| 3:14.128 | הנזכר |  |
| 3:14.129 | ומיד |  |
| 3:14.130 | מי |  |
| 3:14.131 | יצא | . |
| 3:14.132 | ונתבאר |  |
| 3:14.133 | באר |  |
| 3:14.134 | היטב |  |
| 3:14.135 | כי |  |
| 3:14.136 | הוא |  |
| 3:14.137 | חיבור |  |
| 3:14.138 | אחד |  |
| 3:14.139 | האחרונים |  |
| 3:14.140 | ולא |  |
| 3:14.141 | יצאו |  |
| 3:14.142 | אותן |  |
| 3:14.143 | הרעיונים |  |
| 3:14.144 | מרשב"י |  |
| 3:14.145 | התנא |  |
| 3:14.146 | ע"ה |  |
| 3:14.147 | כלל | , |
| 3:14.148 | כי |  |
| 3:14.149 | נמצא |  |
| 3:14.150 | בו |  |
| 3:14.151 | כמה |  |
| 3:14.152 | דברים |  |
| 3:14.153 | בעיקרי |  |
| 3:14.154 | אמונת |  |
| 3:14.155 | תורתינו |  |
| 3:14.156 | הקדושה |  |
| 3:14.157 | היפך |  |
| 3:14.158 | דברי |  |
| 3:14.159 | רשב"י |  |
| 3:14.160 | בתלמוד |  |
| 3:14.161 | ובמדרשי |  |
| 3:14.162 | רז"ל |  |
| 3:14.163 | המפורסמים |  |
| 3:14.164 | בישראל | , |
| 3:14.165 | ודאי |  |
| 3:14.166 | שאין |  |
| 3:14.167 | לסמוך |  |
| 3:14.168 | עליו |  |
| 3:14.169 | בשום |  |
| 3:14.170 | דבר |  |
| 3:14.171 | הנוגע |  |
| 3:14.172 | למנהגי |  |
| 3:14.173 | ישראל |  |
| 3:14.174 | שנהגו |  |
| 3:14.175 | בהם |  |
| 3:14.176 | קדמונינו | . |
| 3:14.177 | ולא |  |
| 3:14.178 | יפה |  |
| 3:14.179 | עשו |  |
| 3:14.180 | חכמי |  |
| 3:14.181 | ימי |  |
| 3:14.182 | הבינים |  |
| 3:14.183 | ששינו |  |
| 3:14.184 | עלינו |  |
| 3:14.185 | כמה |  |
| 3:14.186 | דברים |  |
| 3:14.187 | בסידורי |  |
| 3:14.188 | התפלות |  |
| 3:14.189 | על |  |
| 3:14.190 | פי |  |
| 3:14.191 | הזהר |  |
| 3:14.192 | והמקובלים | , |
| 3:14.193 | כי |  |
| 3:14.194 | שינויים |  |
| 3:14.195 | ההם |  |
| 3:14.196 | יש |  |
| 3:14.197 | מהם |  |
| 3:14.198 | מביאים |  |
| 3:14.199 | לידי |  |
| 3:14.200 | מינות |  |
| 3:14.201 | להאמין |  |
| 3:14.202 | ברשויות |  |
| 3:14.203 | הרבה |  |
| 3:14.204 | ולעבוד |  |
| 3:14.205 | אלהים |  |
| 3:14.206 | אחרים |  |
| 3:14.207 | כמו |  |
| 3:14.208 | שיתבאר |  |
| 3:14.209 | להלן | . |


### Translation under review (3:14)

English: §14. See and understand, our friend the hidden reader, how our honored Rabbi Tam, the aforementioned, cried out like a herald against those who engage in the new Kabbalah, for they breached the fences of the Torah, and laid bare, laid bare, its foundations and its roots in the unity of the Name, may He be blessed, as will be explained below. Instead of thinking to draw near to the Name, may He be blessed, they strayed from the way and came to believe in many other gods, and to worship them! As the words of Rabbi Tam ben Yachya, Mahari Albo also wrote in the Ikkarim, chapter 28, of the second. And further he expounded at length in the book Ravid ha-Zahav, and also brought the words of the great Rabbi Maharshal, in his responses and in the work Yam shel Shlomo, against those who boast of the novelties they study in Kabbalah. And he wrote that since the Talmud was sealed, there is no need for us to be stringent against the Gemara, and it is like heresy, as explained at length in Chok Ya'akov, section 489, subsection 11. All this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai, the Tanna, But after the criticism which the wise of heart examined and investigated after him, from whose womb the things mentioned in it came out, and when the work mentioned came out, and from whose hand it came out. And it was shown clearly and well that it is a single work of the latter ones, and those ideas did not come out of Rabbi Shimon bar Yochai, the Tanna, peace be upon him, at all. for in it are found many things in the fundamentals of the faith of our holy Torah contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory, famous in Israel, certainly one must not rely on it in any matter touching the customs of Israel that our forebears practised. And the sages of the Middle Ages did not do well, who changed upon us several things in the prayer books, according to the Zohar and the Kabbalists, for some of those changes lead to heresy, to believe in many powers and to worship other gods, as will be explained below.

Literal: §14 see and understand our friend the reader hidden how as the honored Rabbi Tam the aforementioned cried out like a herald against those who engage in Kabbalah the new for breached the fences of the Torah and laid bare laid bare its foundations and its roots in the unity the Name, may He be blessed as will be explained below and instead of thought to draw near to the Name, may He be blessed strayed from the way and came and believed in gods many other and to worship them and as the words of Rabbi Tam ben Yachya wrote also likewise Mahari Albo in the book Ikkarim chapter 28 of the second and further expounded at length in the book Ravid ha-Zahav and brought also the words the Rabbi great Maharshal in his responses and in the work Yam of Shlomo against those who boast of novelties which they study in Kabbalah and he wrote that since was sealed the Talmud there is no for us to be stringent against the Gemara and it is like heresy as explained in Chok Ya'akov section 489 subsection 11 at length and all this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai the Tanna but after the criticism which they examined and investigated the wise of heart after him from the womb of whose came out the things mentioned in it and when came out the work mentioned and from the hand of whom came out and it was shown clearly well that it a work single of the latter ones and not came out those ideas from Rabbi Shimon bar Yochai the Tanna peace be upon him at all for is found in it many things in the fundamentals of the faith of our Torah holy contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory famous in Israel certainly one must not rely on it in any matter touching the customs of Israel which they practised in our forebears and not well did the sages of the days of the Middle Ages who changed upon us several things in the prayer books the prayers according to the word of the Zohar and the Kabbalists for changes those there are of them that lead to heresy to believe in powers many and to worship gods other as will be explained below

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:14.1 | §14 |  |  |  |
| 3:14.2 | see |  |  |  |
| 3:14.3 | and understand |  |  |  |
| 3:14.4 | our friend |  |  |  |
| 3:14.5 | the reader |  |  |  |
| 3:14.6 | hidden |  |  |  |
| 3:14.7 | how |  |  |  |
| 3:14.8 | as the honored Rabbi |  |  | כמוהר"ר: an honorific; the author gives it to Rabbi Tam. |
| 3:14.9 | Tam |  |  |  |
| 3:14.10 | the aforementioned |  |  |  |
| 3:14.11 | cried out |  |  |  |
| 3:14.12 | like a herald |  |  | ככרוכיא: a herald or town crier (karukya). |
| 3:14.13 | against |  |  |  |
| 3:14.14 | those who engage |  |  |  |
| 3:14.15 | in Kabbalah |  |  |  |
| 3:14.16 | the new |  |  |  |
| 3:14.17 | for |  |  |  |
| 3:14.18 | breached |  |  |  |
| 3:14.19 | the fences |  |  |  |
| 3:14.20 | of |  |  |  |
| 3:14.21 | the Torah |  |  |  |
| 3:14.22 | and laid bare |  |  | וערו ערו: a doubled verb; the English keeps it doubled. |
| 3:14.23 | laid bare |  |  |  |
| 3:14.24 | its foundations |  |  |  |
| 3:14.25 | and its roots |  |  |  |
| 3:14.26 | in the unity |  |  |  |
| 3:14.27 | the Name, may He be blessed |  |  | השי"ת: the Name, may He be blessed (fixed form). |
| 3:14.28 | as |  |  |  |
| 3:14.29 | will be explained |  |  |  |
| 3:14.30 | below |  |  |  |
| 3:14.31 | and instead |  |  |  |
| 3:14.32 | of |  |  |  |
| 3:14.33 | thought |  |  |  |
| 3:14.34 | to draw near |  |  |  |
| 3:14.35 | to the Name, may He be blessed |  |  |  |
| 3:14.36 | strayed |  |  |  |
| 3:14.37 | from |  |  |  |
| 3:14.38 | the way |  |  |  |
| 3:14.39 | and came |  |  |  |
| 3:14.40 | and believed |  |  |  |
| 3:14.41 | in gods |  |  |  |
| 3:14.42 | many |  |  |  |
| 3:14.43 | other |  |  |  |
| 3:14.44 | and to worship them |  |  |  |
| 3:14.45 | and as the words of |  |  |  |
| 3:14.46 | Rabbi |  |  |  |
| 3:14.47 | Tam |  |  |  |
| 3:14.48 | ben |  |  |  |
| 3:14.49 | Yachya |  |  |  |
| 3:14.50 | wrote |  |  |  |
| 3:14.51 | also |  |  |  |
| 3:14.52 | likewise |  |  |  |
| 3:14.53 | Mahari |  |  | מהר"י אלבו: Morenu ha-Rav Rabbi Y. Albo, the author of the Ikkarim; the text names him only by the acronym. |
| 3:14.54 | Albo |  |  |  |
| 3:14.55 | in the book |  |  |  |
| 3:14.56 | Ikkarim |  |  |  |
| 3:14.57 | chapter |  |  |  |
| 3:14.58 | 28 |  |  |  |
| 3:14.59 | of the second |  |  | משני: the second treatise of the Ikkarim. |
| 3:14.60 | and further |  |  |  |
| 3:14.61 | expounded at length |  |  |  |
| 3:14.62 | in the book |  |  |  |
| 3:14.63 | Ravid |  |  |  |
| 3:14.64 | ha-Zahav |  |  |  |
| 3:14.65 | and brought |  |  |  |
| 3:14.66 | also |  |  |  |
| 3:14.67 | the words |  |  |  |
| 3:14.68 | the Rabbi |  |  |  |
| 3:14.69 | great |  |  |  |
| 3:14.70 | Maharshal |  |  |  |
| 3:14.71 | in his responses |  |  |  |
| 3:14.72 | and in the work |  |  |  |
| 3:14.73 | Yam |  |  |  |
| 3:14.74 | of |  |  |  |
| 3:14.75 | Shlomo |  |  |  |
| 3:14.76 | against |  |  |  |
| 3:14.77 | those who boast |  |  |  |
| 3:14.78 | of novelties |  |  |  |
| 3:14.79 | which they study |  |  |  |
| 3:14.80 | in Kabbalah |  |  |  |
| 3:14.81 | and he wrote |  |  |  |
| 3:14.82 | that since |  |  |  |
| 3:14.83 | was sealed |  |  |  |
| 3:14.84 | the Talmud |  |  |  |
| 3:14.85 | there is no |  |  |  |
| 3:14.86 | for us |  |  |  |
| 3:14.87 | to be stringent |  |  |  |
| 3:14.88 | against |  |  |  |
| 3:14.89 | the Gemara |  |  |  |
| 3:14.90 | and it is |  |  |  |
| 3:14.91 | like |  |  |  |
| 3:14.92 | heresy |  |  |  |
| 3:14.93 | as explained |  |  |  |
| 3:14.94 | in Chok |  |  |  |
| 3:14.95 | Ya'akov |  |  |  |
| 3:14.96 | section |  |  |  |
| 3:14.97 | 489 |  |  | תפ"ט: the letters tav, pe, tet, which make 489. |
| 3:14.98 | subsection |  |  |  |
| 3:14.99 | 11 |  |  | י"א: the subsection number, 11. |
| 3:14.100 | at length |  |  |  |
| 3:14.101 | and all |  |  |  |
| 3:14.102 | this |  |  |  |
| 3:14.103 | is |  |  |  |
| 3:14.104 | according to |  |  |  |
| 3:14.105 | the assumption |  |  |  |
| 3:14.106 | and the ascription |  |  |  |
| 3:14.107 | which they ascribed |  |  |  |
| 3:14.108 | the Zohar |  |  |  |
| 3:14.109 | to Rabbi Shimon bar Yochai |  |  |  |
| 3:14.110 | the Tanna |  |  |  |
| 3:14.111 | but |  |  |  |
| 3:14.112 | after |  |  |  |
| 3:14.113 | the criticism |  |  |  |
| 3:14.114 | which they examined |  |  |  |
| 3:14.115 | and investigated |  |  |  |
| 3:14.116 | the wise |  |  |  |
| 3:14.117 | of heart |  |  |  |
| 3:14.118 | after him |  |  |  |
| 3:14.119 | from the womb of |  |  |  |
| 3:14.120 | whose |  |  |  |
| 3:14.121 | came out |  |  |  |
| 3:14.122 | the things |  |  |  |
| 3:14.123 | mentioned |  |  |  |
| 3:14.124 | in it |  |  |  |
| 3:14.125 | and when |  |  |  |
| 3:14.126 | came out |  |  |  |
| 3:14.127 | the work |  |  |  |
| 3:14.128 | mentioned |  |  |  |
| 3:14.129 | and from the hand of |  |  |  |
| 3:14.130 | whom |  |  |  |
| 3:14.131 | came out |  |  |  |
| 3:14.132 | and it was shown |  |  |  |
| 3:14.133 | clearly |  |  |  |
| 3:14.134 | well |  |  |  |
| 3:14.135 | that |  |  |  |
| 3:14.136 | it |  |  |  |
| 3:14.137 | a work |  |  |  |
| 3:14.138 | single |  |  |  |
| 3:14.139 | of the latter ones |  |  |  |
| 3:14.140 | and not |  |  |  |
| 3:14.141 | came out |  |  |  |
| 3:14.142 | those |  |  |  |
| 3:14.143 | ideas |  |  |  |
| 3:14.144 | from Rabbi Shimon bar Yochai |  |  |  |
| 3:14.145 | the Tanna |  |  |  |
| 3:14.146 | peace be upon him |  |  | ע"ה: peace be upon him. |
| 3:14.147 | at all |  |  |  |
| 3:14.148 | for |  |  |  |
| 3:14.149 | is found |  |  |  |
| 3:14.150 | in it |  |  |  |
| 3:14.151 | many |  |  |  |
| 3:14.152 | things |  |  |  |
| 3:14.153 | in the fundamentals of |  |  |  |
| 3:14.154 | the faith of |  |  |  |
| 3:14.155 | our Torah |  |  |  |
| 3:14.156 | holy |  |  |  |
| 3:14.157 | contrary to |  |  |  |
| 3:14.158 | the words of |  |  |  |
| 3:14.159 | Rabbi Shimon bar Yochai |  |  |  |
| 3:14.160 | in the Talmud |  |  |  |
| 3:14.161 | and in the Midrashim of |  |  |  |
| 3:14.162 | our Rabbis of blessed memory |  |  |  |
| 3:14.163 | famous |  |  |  |
| 3:14.164 | in Israel |  |  |  |
| 3:14.165 | certainly |  |  |  |
| 3:14.166 | one must not |  |  |  |
| 3:14.167 | rely |  |  |  |
| 3:14.168 | on it |  |  |  |
| 3:14.169 | in any |  |  |  |
| 3:14.170 | matter |  |  |  |
| 3:14.171 | touching |  |  |  |
| 3:14.172 | the customs of |  |  |  |
| 3:14.173 | Israel |  |  |  |
| 3:14.174 | which they practised |  |  |  |
| 3:14.175 | in |  |  |  |
| 3:14.176 | our forebears |  |  |  |
| 3:14.177 | and not |  |  |  |
| 3:14.178 | well |  |  |  |
| 3:14.179 | did |  |  |  |
| 3:14.180 | the sages of |  |  |  |
| 3:14.181 | the days of |  |  |  |
| 3:14.182 | the Middle Ages |  |  | ימי הבינים: the Middle Ages, the author's period term. |
| 3:14.183 | who changed |  |  |  |
| 3:14.184 | upon us |  |  |  |
| 3:14.185 | several |  |  |  |
| 3:14.186 | things |  |  |  |
| 3:14.187 | in the prayer books |  |  |  |
| 3:14.188 | the prayers |  |  |  |
| 3:14.189 | according to |  |  |  |
| 3:14.190 | the word of |  |  |  |
| 3:14.191 | the Zohar |  |  |  |
| 3:14.192 | and the Kabbalists |  |  |  |
| 3:14.193 | for |  |  |  |
| 3:14.194 | changes |  |  |  |
| 3:14.195 | those |  |  |  |
| 3:14.196 | there are |  |  |  |
| 3:14.197 | of them |  |  |  |
| 3:14.198 | that lead |  |  |  |
| 3:14.199 | to |  |  |  |
| 3:14.200 | heresy |  |  |  |
| 3:14.201 | to believe |  |  |  |
| 3:14.202 | in powers |  |  |  |
| 3:14.203 | many |  |  |  |
| 3:14.204 | and to worship |  |  |  |
| 3:14.205 | gods |  |  |  |
| 3:14.206 | other |  |  |  |
| 3:14.207 | as |  |  |  |
| 3:14.208 | will be explained |  |  |  |
| 3:14.209 | below |  |  |  |

Phrases:
- 3:14.1 3:14.2 3:14.3 3:14.4 3:14.5 3:14.6 3:14.7 3:14.8 | §14. See and understand, our friend the hidden reader, how our honored Rabbi | §14 see and understand our friend the reader hidden how as the honored Rabbi
- 3:14.9 3:14.10 | Tam, the aforementioned, | Tam the aforementioned
- 3:14.11 3:14.12 3:14.13 3:14.14 3:14.15 3:14.16 | cried out like a herald against those who engage in the new Kabbalah, | cried out like a herald against those who engage in Kabbalah the new
- 3:14.17 3:14.18 3:14.19 3:14.20 3:14.21 | for they breached the fences of the Torah, | for breached the fences of the Torah
- 3:14.22 3:14.23 3:14.24 3:14.25 | and laid bare, laid bare, its foundations and its roots | and laid bare laid bare its foundations and its roots
- 3:14.26 3:14.27 3:14.28 3:14.29 3:14.30 | in the unity of the Name, may He be blessed, as will be explained below. | in the unity the Name, may He be blessed as will be explained below
- 3:14.31 3:14.32 | Instead of | and instead of
- 3:14.33 3:14.34 3:14.35 3:14.36 3:14.37 3:14.38 | thinking to draw near to the Name, may He be blessed, they strayed from the way | thought to draw near to the Name, may He be blessed strayed from the way
- 3:14.39 3:14.40 3:14.41 3:14.42 3:14.43 | and came to believe in many other gods, | and came and believed in gods many other
- 3:14.44 | and to worship them! | and to worship them
- 3:14.45 3:14.46 3:14.47 3:14.48 3:14.49 3:14.50 3:14.51 3:14.52 3:14.53 3:14.54 | As the words of Rabbi Tam ben Yachya, Mahari Albo also wrote | and as the words of Rabbi Tam ben Yachya wrote also likewise Mahari Albo
- 3:14.55 3:14.56 3:14.57 3:14.58 3:14.59 | in the Ikkarim, chapter 28, of the second. | in the book Ikkarim chapter 28 of the second
- 3:14.60 3:14.61 3:14.62 3:14.63 3:14.64 | And further he expounded at length in the book Ravid ha-Zahav, | and further expounded at length in the book Ravid ha-Zahav
- 3:14.65 3:14.66 3:14.67 3:14.68 3:14.69 3:14.70 | and also brought the words of the great Rabbi Maharshal, | and brought also the words the Rabbi great Maharshal
- 3:14.71 3:14.72 3:14.73 3:14.74 3:14.75 | in his responses and in the work Yam shel Shlomo, | in his responses and in the work Yam of Shlomo
- 3:14.76 3:14.77 3:14.78 3:14.79 3:14.80 | against those who boast of the novelties they study in Kabbalah. | against those who boast of novelties which they study in Kabbalah
- 3:14.81 3:14.82 3:14.83 3:14.84 3:14.85 3:14.86 3:14.87 3:14.88 3:14.89 3:14.90 | And he wrote that since the Talmud was sealed, there is no need for us to be stringent against the Gemara, and it is | and he wrote that since was sealed the Talmud there is no for us to be stringent against the Gemara and it is
- 3:14.91 3:14.92 3:14.93 3:14.94 3:14.95 3:14.96 3:14.97 3:14.98 3:14.99 3:14.100 | like heresy, as explained at length in Chok Ya'akov, section 489, subsection 11. | like heresy as explained in Chok Ya'akov section 489 subsection 11 at length
- 3:14.101 3:14.102 3:14.103 3:14.104 3:14.105 3:14.106 3:14.107 3:14.108 3:14.109 3:14.110 | All this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai, the Tanna, | and all this is according to the assumption and the ascription which they ascribed the Zohar to Rabbi Shimon bar Yochai the Tanna
- 3:14.111 3:14.112 3:14.113 3:14.114 3:14.115 3:14.116 3:14.117 3:14.118 3:14.119 3:14.120 | But after the criticism which the wise of heart examined and investigated after him, from whose womb | but after the criticism which they examined and investigated the wise of heart after him from the womb of whose
- 3:14.121 3:14.122 3:14.123 3:14.124 3:14.125 3:14.126 3:14.127 3:14.128 3:14.129 3:14.130 3:14.131 | the things mentioned in it came out, and when the work mentioned came out, and from whose hand it came out. | came out the things mentioned in it and when came out the work mentioned and from the hand of whom came out
- 3:14.132 3:14.133 3:14.134 3:14.135 3:14.136 | And it was shown clearly and well that it | and it was shown clearly well that it
- 3:14.137 3:14.138 3:14.139 | is a single work of the latter ones, | a work single of the latter ones
- 3:14.140 3:14.141 3:14.142 3:14.143 3:14.144 3:14.145 3:14.146 3:14.147 | and those ideas did not come out of Rabbi Shimon bar Yochai, the Tanna, peace be upon him, at all. | and not came out those ideas from Rabbi Shimon bar Yochai the Tanna peace be upon him at all
- 3:14.148 3:14.149 3:14.150 3:14.151 3:14.152 3:14.153 3:14.154 3:14.155 3:14.156 3:14.157 3:14.158 3:14.159 3:14.160 3:14.161 3:14.162 3:14.163 3:14.164 | for in it are found many things in the fundamentals of the faith of our holy Torah contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory, famous in Israel, | for is found in it many things in the fundamentals of the faith of our Torah holy contrary to the words of Rabbi Shimon bar Yochai in the Talmud and in the Midrashim of our Rabbis of blessed memory famous in Israel
- 3:14.165 3:14.166 3:14.167 3:14.168 3:14.169 3:14.170 3:14.171 3:14.172 3:14.173 3:14.174 3:14.175 3:14.176 | certainly one must not rely on it in any matter touching the customs of Israel that our forebears practised. | certainly one must not rely on it in any matter touching the customs of Israel which they practised in our forebears
- 3:14.177 3:14.178 3:14.179 3:14.180 3:14.181 3:14.182 3:14.183 3:14.184 3:14.185 3:14.186 | And the sages of the Middle Ages did not do well, who changed upon us several things | and not well did the sages of the days of the Middle Ages who changed upon us several things
- 3:14.187 3:14.188 3:14.189 3:14.190 3:14.191 | in the prayer books, according to the Zohar | in the prayer books the prayers according to the word of the Zohar
- 3:14.192 3:14.193 3:14.194 3:14.195 3:14.196 3:14.197 3:14.198 3:14.199 3:14.200 | and the Kabbalists, for some of those changes lead to heresy, | and the Kabbalists for changes those there are of them that lead to heresy
- 3:14.201 3:14.202 3:14.203 3:14.204 3:14.205 3:14.206 3:14.207 3:14.208 3:14.209 | to believe in many powers and to worship other gods, as will be explained below. | to believe in powers many and to worship gods other as will be explained below

Grammar notes:
- ראה והבן ... איך ... צוח: 'see and understand how ... cried out'; the two imperatives open one long period, which runs to the end of the paragraph.
- ותחת אשר חשבו להתקרב ... תעו: 'instead of thinking to draw near, they strayed'; the author sets the intention against the act.
- כ' גם כן מהר"י אלבו: 'wrote also Mahari Albo'; the verb comes first in the Hebrew and the English places it after the name.
- וכתב דמאחר ... והוי כמו מינות: an indirect report with no quotation marks.
- וכל זה הוא לפי ההנחה והיחוס: 'all this is according to the assumption and the ascription'; the ascription to Rabbi Shimon bar Yochai is what is disputed.

Translator's note: The paragraph is a single period with a report of two authorities at its end. The Wikisource base text is used; the 1931 scan is not collated.

Commentary rulings:
- apparatus: Allusion, not quotation: 'אלהים אחרים' echoes Exod 20:3 ('other gods'); the English follows the author's wording.
- apparatus: Base text is Wikisource; a general line until the 1931 scan is collated (open question 2).

## Milchamot ha-Shem 3:15

Hebrew (Wikisource): טו) גם הביא עוד הר' רביד הזהב דברי מרן בב"י דחייש לדברי הזהר במקום שלא נזכר בגמרא היפך מדבריו, וכתב עליו הרמ"א ז"ל שאע"פ שלא נזכר בגמרא היפך דבריו, אין לזה מדברי הפוסקים, ואין לחוש לדברי הזהר ע"ש. עוד הביא דברי הרדב"ז ז"ל שכתב שכל דבר שנזכר בגמרא או באחד מן הפוסקים, או בעלי ההלכות אפי' שיהיה היפך ממה שכתוב בספרי הקבלה אני מורה בו בין להקל בין להחמיר, ולא אחוש למה שכתוב באותן הספרים. ולעצמי אם הוא חומרא אני נוהג אותו ואם קולא לא אחוש לה. ועוד כתב בשם הרא"ם ושאר פוסקים שאין לחוש לדברי המקובלים נגד שום פוסק ע"ש. אמור מעתה שמדברי הפוסקים הנזכרים בתשובת הרב רבי הזהב והם מרן ז"ל ומור"ם והרדב"ז וכנה"ג ומהר"י הלוי כ' מהר"ש הלוי ור' אליהו הלוי, והר"מ אלקבולי ור"ת בן יחיא והר' חק יעקב ומהרש"ל ומהר"ם אלשקר שאומרים שאין לחוש לדברי הזהר והמקובלים שהוא נגד המשנה והתלמוד ואפילו נגד פוסק יחידי. הנה מזה הוכיח לעצמך שכתבת שהם הלכה למשה מסיני, שדבריך בדותא וסכלות גמורה.

| id | word | punctuation after |
|---|---|---|
| 3:15.1 | טו | ) |
| 3:15.2 | גם |  |
| 3:15.3 | הביא |  |
| 3:15.4 | עוד |  |
| 3:15.5 | הר | ' |
| 3:15.6 | רביד |  |
| 3:15.7 | הזהב |  |
| 3:15.8 | דברי |  |
| 3:15.9 | מרן |  |
| 3:15.10 | בב"י |  |
| 3:15.11 | דחייש |  |
| 3:15.12 | לדברי |  |
| 3:15.13 | הזהר |  |
| 3:15.14 | במקום |  |
| 3:15.15 | שלא |  |
| 3:15.16 | נזכר |  |
| 3:15.17 | בגמרא |  |
| 3:15.18 | היפך |  |
| 3:15.19 | מדבריו | , |
| 3:15.20 | וכתב |  |
| 3:15.21 | עליו |  |
| 3:15.22 | הרמ"א |  |
| 3:15.23 | ז"ל |  |
| 3:15.24 | שאע"פ |  |
| 3:15.25 | שלא |  |
| 3:15.26 | נזכר |  |
| 3:15.27 | בגמרא |  |
| 3:15.28 | היפך |  |
| 3:15.29 | דבריו | , |
| 3:15.30 | אין |  |
| 3:15.31 | לזה |  |
| 3:15.32 | מדברי |  |
| 3:15.33 | הפוסקים | , |
| 3:15.34 | ואין |  |
| 3:15.35 | לחוש |  |
| 3:15.36 | לדברי |  |
| 3:15.37 | הזהר |  |
| 3:15.38 | ע"ש | . |
| 3:15.39 | עוד |  |
| 3:15.40 | הביא |  |
| 3:15.41 | דברי |  |
| 3:15.42 | הרדב"ז |  |
| 3:15.43 | ז"ל |  |
| 3:15.44 | שכתב |  |
| 3:15.45 | שכל |  |
| 3:15.46 | דבר |  |
| 3:15.47 | שנזכר |  |
| 3:15.48 | בגמרא |  |
| 3:15.49 | או |  |
| 3:15.50 | באחד |  |
| 3:15.51 | מן |  |
| 3:15.52 | הפוסקים | , |
| 3:15.53 | או |  |
| 3:15.54 | בעלי |  |
| 3:15.55 | ההלכות |  |
| 3:15.56 | אפי | ' |
| 3:15.57 | שיהיה |  |
| 3:15.58 | היפך |  |
| 3:15.59 | ממה |  |
| 3:15.60 | שכתוב |  |
| 3:15.61 | בספרי |  |
| 3:15.62 | הקבלה |  |
| 3:15.63 | אני |  |
| 3:15.64 | מורה |  |
| 3:15.65 | בו |  |
| 3:15.66 | בין |  |
| 3:15.67 | להקל |  |
| 3:15.68 | בין |  |
| 3:15.69 | להחמיר | , |
| 3:15.70 | ולא |  |
| 3:15.71 | אחוש |  |
| 3:15.72 | למה |  |
| 3:15.73 | שכתוב |  |
| 3:15.74 | באותן |  |
| 3:15.75 | הספרים | . |
| 3:15.76 | ולעצמי |  |
| 3:15.77 | אם |  |
| 3:15.78 | הוא |  |
| 3:15.79 | חומרא |  |
| 3:15.80 | אני |  |
| 3:15.81 | נוהג |  |
| 3:15.82 | אותו |  |
| 3:15.83 | ואם |  |
| 3:15.84 | קולא |  |
| 3:15.85 | לא |  |
| 3:15.86 | אחוש |  |
| 3:15.87 | לה | . |
| 3:15.88 | ועוד |  |
| 3:15.89 | כתב |  |
| 3:15.90 | בשם |  |
| 3:15.91 | הרא"ם |  |
| 3:15.92 | ושאר |  |
| 3:15.93 | פוסקים |  |
| 3:15.94 | שאין |  |
| 3:15.95 | לחוש |  |
| 3:15.96 | לדברי |  |
| 3:15.97 | המקובלים |  |
| 3:15.98 | נגד |  |
| 3:15.99 | שום |  |
| 3:15.100 | פוסק |  |
| 3:15.101 | ע"ש | . |
| 3:15.102 | אמור |  |
| 3:15.103 | מעתה |  |
| 3:15.104 | שמדברי |  |
| 3:15.105 | הפוסקים |  |
| 3:15.106 | הנזכרים |  |
| 3:15.107 | בתשובת |  |
| 3:15.108 | הרב |  |
| 3:15.109 | רבי |  |
| 3:15.110 | הזהב |  |
| 3:15.111 | והם |  |
| 3:15.112 | מרן |  |
| 3:15.113 | ז"ל |  |
| 3:15.114 | ומור"ם |  |
| 3:15.115 | והרדב"ז |  |
| 3:15.116 | וכנה"ג |  |
| 3:15.117 | ומהר"י |  |
| 3:15.118 | הלוי |  |
| 3:15.119 | כ | ' |
| 3:15.120 | מהר"ש |  |
| 3:15.121 | הלוי |  |
| 3:15.122 | ור | ' |
| 3:15.123 | אליהו |  |
| 3:15.124 | הלוי | , |
| 3:15.125 | והר"מ |  |
| 3:15.126 | אלקבולי |  |
| 3:15.127 | ור"ת |  |
| 3:15.128 | בן |  |
| 3:15.129 | יחיא |  |
| 3:15.130 | והר | ' |
| 3:15.131 | חק |  |
| 3:15.132 | יעקב |  |
| 3:15.133 | ומהרש"ל |  |
| 3:15.134 | ומהר"ם |  |
| 3:15.135 | אלשקר |  |
| 3:15.136 | שאומרים |  |
| 3:15.137 | שאין |  |
| 3:15.138 | לחוש |  |
| 3:15.139 | לדברי |  |
| 3:15.140 | הזהר |  |
| 3:15.141 | והמקובלים |  |
| 3:15.142 | שהוא |  |
| 3:15.143 | נגד |  |
| 3:15.144 | המשנה |  |
| 3:15.145 | והתלמוד |  |
| 3:15.146 | ואפילו |  |
| 3:15.147 | נגד |  |
| 3:15.148 | פוסק |  |
| 3:15.149 | יחידי | . |
| 3:15.150 | הנה |  |
| 3:15.151 | מזה |  |
| 3:15.152 | הוכיח |  |
| 3:15.153 | לעצמך |  |
| 3:15.154 | שכתבת |  |
| 3:15.155 | שהם |  |
| 3:15.156 | הלכה |  |
| 3:15.157 | למשה |  |
| 3:15.158 | מסיני | , |
| 3:15.159 | שדבריך |  |
| 3:15.160 | בדותא |  |
| 3:15.161 | וסכלות |  |
| 3:15.162 | גמורה | . |


### Translation under review (3:15)

English: §15. Also he brought further the Rabbi, the author of Ravid ha-Zahav, the words of Maran in Beit Yosef, who was concerned about the words of the Zohar in a place where the Gemara does not mention anything contrary to them. And he wrote about him, the Rema, of blessed memory, that although the Gemara does not mention anything contrary to its words, there is no basis for this in the words of the decisors, and one need not be concerned with the words of the Zohar, see there. Further he brought the words of the Radbaz, of blessed memory, that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot, even if it be contrary to what is written in the books of Kabbalah, I rule on, both to be lenient and to be strict, and I will not be concerned with what is written in those books. As for myself, if it is a stringency, I follow it and if a leniency, I will not be concerned with it. Further he wrote in the name of the Ra'em and other decisors that one need not be concerned about the words of the Kabbalists against any decisor, see there. Say now, that from the decisors mentioned in the responsum of the Rav, Rabbi ha-Zahav, and they are Maran, of blessed memory, and Moreinu ha-Rav Moshe, and the Radbaz, and the Kenesset ha-Gedolah, and Mahari ha-Levi wrote, Maharash ha-Levi and Rabbi Eliyahu ha-Levi, and Rabbi Moshe Alkabetz, and Rabbi Tam ben Yachya, and the Rabbi, the author of Chok Ya'akov, and Maharshal, and Maharam al-Shakar, who say that one need not be concerned with the words of the Zohar and the Kabbalists, when they are against the Mishnah and the Talmud, and even against a single decisor. Behold, from this prove to yourself that you wrote that they are a law of Moses from Sinai, that your words are a fabrication and utter folly.

Literal: §15 also he brought further the Rabbi, the author of Ravid ha-Zahav the words of Maran in Beit Yosef who was concerned about the words of the Zohar in a place where not mentioned in the Gemara contrary to them and he wrote about him the Rema of blessed memory that although that not mentioned in the Gemara contrary to its words there is no for this in the words of the decisors and one need not be concerned about the words of the Zohar see there further he brought the words of the Radbaz of blessed memory that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot even if it be contrary to what is written in the books of Kabbalah I rule on it both to be lenient and to be strict and not I will be concerned with what is written in those books as for myself if it is a stringency I follow it and if a leniency not I will be concerned with it further he wrote in the name of the Ra'em and other decisors that one need not be concerned about words the Kabbalists against any decisor see there say now that from the words of the decisors mentioned in the responsum of the Rav Rabbi ha-Zahav and they are Maran of blessed memory and Moreinu ha-Rav Moshe and the Radbaz and the Kenesset ha-Gedolah and Mahari ha-Levi wrote Maharash ha-Levi and Rabbi Eliyahu ha-Levi and Rabbi Moshe Alkabetz and Rabbi Tam ben Yachya and the Rabbi, the author of Chok Ya'akov and Maharshal and Maharam al-Shakar who say that one need not be concerned about words the Zohar and the Kabbalists when they are against the Mishnah and the Talmud and even against a decisor single behold from this prove to yourself that you wrote that they are a law of Moses from Sinai that your words are a fabrication and folly utter

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:15.1 | §15 |  |  |  |
| 3:15.2 | also |  |  |  |
| 3:15.3 | he brought |  |  |  |
| 3:15.4 | further |  |  |  |
| 3:15.5 | the Rabbi, the author of |  |  | הר' + a book title: 'the Rabbi, the author of Ravid ha-Zahav'; the title is never a personal name. |
| 3:15.6 | Ravid |  |  |  |
| 3:15.7 | ha-Zahav |  |  |  |
| 3:15.8 | the words of |  |  |  |
| 3:15.9 | Maran |  |  | מרן: the author's honorific for the writer of the Beit Yosef; the text names no one. Literally 'our master'. |
| 3:15.10 | in Beit Yosef |  |  | בב"י: Beit Yosef, the commentary on the Tur. |
| 3:15.11 | who was concerned |  |  |  |
| 3:15.12 | about the words of |  |  |  |
| 3:15.13 | the Zohar |  |  |  |
| 3:15.14 | in a place |  |  |  |
| 3:15.15 | where not |  |  |  |
| 3:15.16 | mentioned |  |  |  |
| 3:15.17 | in the Gemara |  |  |  |
| 3:15.18 | contrary |  |  |  |
| 3:15.19 | to them |  |  |  |
| 3:15.20 | and he wrote |  |  |  |
| 3:15.21 | about him |  |  |  |
| 3:15.22 | the Rema |  |  | הרמ"א: the Rema, Rabbi Moshe Isserles. |
| 3:15.23 | of blessed memory |  |  |  |
| 3:15.24 | that although |  |  |  |
| 3:15.25 | that not |  |  |  |
| 3:15.26 | mentioned |  |  |  |
| 3:15.27 | in the Gemara |  |  |  |
| 3:15.28 | contrary |  |  |  |
| 3:15.29 | to its words |  |  |  |
| 3:15.30 | there is no |  |  |  |
| 3:15.31 | for this |  |  |  |
| 3:15.32 | in the words of |  |  |  |
| 3:15.33 | the decisors |  |  |  |
| 3:15.34 | and one need not |  |  |  |
| 3:15.35 | be concerned |  |  |  |
| 3:15.36 | about the words of |  |  |  |
| 3:15.37 | the Zohar |  |  |  |
| 3:15.38 | see there |  |  |  |
| 3:15.39 | further |  |  |  |
| 3:15.40 | he brought |  |  |  |
| 3:15.41 | the words of |  |  |  |
| 3:15.42 | the Radbaz |  |  | הרדב"ז: the Radbaz, Rabbi David ibn Zimra. |
| 3:15.43 | of blessed memory |  |  |  |
| 3:15.44 | that he wrote |  |  |  |
| 3:15.45 | that every |  |  |  |
| 3:15.46 | matter |  |  |  |
| 3:15.47 | mentioned |  |  |  |
| 3:15.48 | in the Gemara |  |  |  |
| 3:15.49 | or |  |  |  |
| 3:15.50 | in one |  |  |  |
| 3:15.51 | of |  |  |  |
| 3:15.52 | the decisors |  |  |  |
| 3:15.53 | or |  |  |  |
| 3:15.54 | the masters of |  |  |  |
| 3:15.55 | the halakhot |  |  |  |
| 3:15.56 | even |  |  |  |
| 3:15.57 | if it be |  |  |  |
| 3:15.58 | contrary |  |  |  |
| 3:15.59 | to what |  |  |  |
| 3:15.60 | is written |  |  |  |
| 3:15.61 | in the books |  |  |  |
| 3:15.62 | of Kabbalah |  |  |  |
| 3:15.63 | I |  |  |  |
| 3:15.64 | rule |  |  |  |
| 3:15.65 | on it |  |  |  |
| 3:15.66 | both |  |  |  |
| 3:15.67 | to be lenient |  |  |  |
| 3:15.68 | and |  |  |  |
| 3:15.69 | to be strict |  |  |  |
| 3:15.70 | and not |  |  |  |
| 3:15.71 | I will be concerned |  |  |  |
| 3:15.72 | with what |  |  |  |
| 3:15.73 | is written |  |  |  |
| 3:15.74 | in those |  |  |  |
| 3:15.75 | books |  |  |  |
| 3:15.76 | as for myself |  |  |  |
| 3:15.77 | if |  |  |  |
| 3:15.78 | it is |  |  |  |
| 3:15.79 | a stringency |  |  |  |
| 3:15.80 | I |  |  |  |
| 3:15.81 | follow |  |  |  |
| 3:15.82 | it |  |  |  |
| 3:15.83 | and if |  |  |  |
| 3:15.84 | a leniency |  |  |  |
| 3:15.85 | not |  |  |  |
| 3:15.86 | I will be concerned |  |  |  |
| 3:15.87 | with it |  |  |  |
| 3:15.88 | further |  |  |  |
| 3:15.89 | he wrote |  |  |  |
| 3:15.90 | in the name of |  |  |  |
| 3:15.91 | the Ra'em |  |  | הרא"ם: the Ra'em. |
| 3:15.92 | and other |  |  |  |
| 3:15.93 | decisors |  |  |  |
| 3:15.94 | that one need not |  |  |  |
| 3:15.95 | be concerned |  |  |  |
| 3:15.96 | about words |  |  |  |
| 3:15.97 | the Kabbalists |  |  |  |
| 3:15.98 | against |  |  |  |
| 3:15.99 | any |  |  |  |
| 3:15.100 | decisor |  |  |  |
| 3:15.101 | see there |  |  |  |
| 3:15.102 | say |  |  |  |
| 3:15.103 | now |  |  |  |
| 3:15.104 | that from the words of |  |  |  |
| 3:15.105 | the decisors |  |  |  |
| 3:15.106 | mentioned |  |  |  |
| 3:15.107 | in the responsum of |  |  |  |
| 3:15.108 | the Rav |  |  |  |
| 3:15.109 | Rabbi |  |  |  |
| 3:15.110 | ha-Zahav |  |  | הזהב: the Hebrew prints 'הרב רבי הזהב' here; see apparatus. |
| 3:15.111 | and they are |  |  |  |
| 3:15.112 | Maran |  |  |  |
| 3:15.113 | of blessed memory |  |  |  |
| 3:15.114 | and Moreinu ha-Rav Moshe |  |  |  |
| 3:15.115 | and the Radbaz |  |  |  |
| 3:15.116 | and the Kenesset ha-Gedolah |  |  | כנה"ג: the Kenesset ha-Gedolah. |
| 3:15.117 | and Mahari |  |  |  |
| 3:15.118 | ha-Levi |  |  |  |
| 3:15.119 | wrote |  |  |  |
| 3:15.120 | Maharash |  |  |  |
| 3:15.121 | ha-Levi |  |  |  |
| 3:15.122 | and Rabbi |  |  |  |
| 3:15.123 | Eliyahu |  |  |  |
| 3:15.124 | ha-Levi |  |  |  |
| 3:15.125 | and Rabbi Moshe |  |  | הר"מ: read as Rabbi Moshe; the abbreviation is not expanded in the text. |
| 3:15.126 | Alkabetz |  |  |  |
| 3:15.127 | and Rabbi Tam |  |  |  |
| 3:15.128 | ben |  |  |  |
| 3:15.129 | Yachya |  |  |  |
| 3:15.130 | and the Rabbi, the author of |  |  |  |
| 3:15.131 | Chok |  |  |  |
| 3:15.132 | Ya'akov |  |  |  |
| 3:15.133 | and Maharshal |  |  |  |
| 3:15.134 | and Maharam |  |  |  |
| 3:15.135 | al-Shakar |  |  | מהר"ם אלשקר: Maharam al-Shakar, as the author writes him. |
| 3:15.136 | who say |  |  |  |
| 3:15.137 | that one need not |  |  |  |
| 3:15.138 | be concerned |  |  |  |
| 3:15.139 | about words |  |  |  |
| 3:15.140 | the Zohar |  |  |  |
| 3:15.141 | and the Kabbalists |  |  |  |
| 3:15.142 | when they are |  |  |  |
| 3:15.143 | against |  |  |  |
| 3:15.144 | the Mishnah |  |  |  |
| 3:15.145 | and the Talmud |  |  |  |
| 3:15.146 | and even |  |  |  |
| 3:15.147 | against |  |  |  |
| 3:15.148 | a decisor |  |  |  |
| 3:15.149 | single |  |  |  |
| 3:15.150 | behold |  |  |  |
| 3:15.151 | from this |  |  |  |
| 3:15.152 | prove |  |  |  |
| 3:15.153 | to yourself |  |  |  |
| 3:15.154 | that you wrote |  |  |  |
| 3:15.155 | that they are |  |  |  |
| 3:15.156 | a law |  |  |  |
| 3:15.157 | of Moses |  |  |  |
| 3:15.158 | from Sinai |  |  |  |
| 3:15.159 | that your words are |  |  |  |
| 3:15.160 | a fabrication |  |  |  |
| 3:15.161 | and folly |  |  |  |
| 3:15.162 | utter |  |  |  |

Phrases:
- 3:15.1 3:15.2 3:15.3 3:15.4 | §15. Also he brought further | §15 also he brought further
- 3:15.5 3:15.6 3:15.7 | the Rabbi, the author of Ravid ha-Zahav, | the Rabbi, the author of Ravid ha-Zahav
- 3:15.8 3:15.9 3:15.10 3:15.11 | the words of Maran in Beit Yosef, who was concerned | the words of Maran in Beit Yosef who was concerned
- 3:15.12 3:15.13 3:15.14 3:15.15 3:15.16 3:15.17 3:15.18 3:15.19 | about the words of the Zohar in a place where the Gemara does not mention anything contrary to them. | about the words of the Zohar in a place where not mentioned in the Gemara contrary to them
- 3:15.20 3:15.21 3:15.22 3:15.23 | And he wrote about him, the Rema, of blessed memory, | and he wrote about him the Rema of blessed memory
- 3:15.24 3:15.25 3:15.26 3:15.27 3:15.28 3:15.29 | that although the Gemara does not mention anything contrary to its words, | that although that not mentioned in the Gemara contrary to its words
- 3:15.30 3:15.31 3:15.32 3:15.33 3:15.34 3:15.35 3:15.36 3:15.37 3:15.38 | there is no basis for this in the words of the decisors, and one need not be concerned with the words of the Zohar, see there. | there is no for this in the words of the decisors and one need not be concerned about the words of the Zohar see there
- 3:15.39 3:15.40 3:15.41 3:15.42 3:15.43 | Further he brought the words of the Radbaz, of blessed memory, | further he brought the words of the Radbaz of blessed memory
- 3:15.44 3:15.45 3:15.46 3:15.47 3:15.48 3:15.49 3:15.50 3:15.51 3:15.52 3:15.53 3:15.54 3:15.55 3:15.56 3:15.57 3:15.58 3:15.59 3:15.60 3:15.61 3:15.62 3:15.63 3:15.64 3:15.65 3:15.66 3:15.67 3:15.68 3:15.69 3:15.70 3:15.71 3:15.72 3:15.73 3:15.74 3:15.75 | that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot, even if it be contrary to what is written in the books of Kabbalah, I rule on, both to be lenient and to be strict, and I will not be concerned with what is written in those books. | that he wrote that every matter mentioned in the Gemara or in one of the decisors or the masters of the halakhot even if it be contrary to what is written in the books of Kabbalah I rule on it both to be lenient and to be strict and not I will be concerned with what is written in those books
- 3:15.76 3:15.77 3:15.78 3:15.79 3:15.80 3:15.81 3:15.82 | As for myself, if it is a stringency, I follow it | as for myself if it is a stringency I follow it
- 3:15.83 3:15.84 3:15.85 3:15.86 3:15.87 | and if a leniency, I will not be concerned with it. | and if a leniency not I will be concerned with it
- 3:15.88 3:15.89 3:15.90 3:15.91 3:15.92 3:15.93 3:15.94 3:15.95 3:15.96 3:15.97 3:15.98 3:15.99 3:15.100 3:15.101 | Further he wrote in the name of the Ra'em and other decisors that one need not be concerned about the words of the Kabbalists against any decisor, see there. | further he wrote in the name of the Ra'em and other decisors that one need not be concerned about words the Kabbalists against any decisor see there
- 3:15.102 3:15.103 3:15.104 3:15.105 3:15.106 3:15.107 3:15.108 3:15.109 3:15.110 | Say now, that from the decisors mentioned in the responsum of the Rav, Rabbi ha-Zahav, | say now that from the words of the decisors mentioned in the responsum of the Rav Rabbi ha-Zahav
- 3:15.111 3:15.112 3:15.113 3:15.114 3:15.115 3:15.116 3:15.117 3:15.118 3:15.119 3:15.120 3:15.121 3:15.122 3:15.123 3:15.124 | and they are Maran, of blessed memory, and Moreinu ha-Rav Moshe, and the Radbaz, and the Kenesset ha-Gedolah, and Mahari ha-Levi wrote, Maharash ha-Levi and Rabbi Eliyahu ha-Levi, | and they are Maran of blessed memory and Moreinu ha-Rav Moshe and the Radbaz and the Kenesset ha-Gedolah and Mahari ha-Levi wrote Maharash ha-Levi and Rabbi Eliyahu ha-Levi
- 3:15.125 3:15.126 3:15.127 3:15.128 3:15.129 | and Rabbi Moshe Alkabetz, and Rabbi Tam ben Yachya, | and Rabbi Moshe Alkabetz and Rabbi Tam ben Yachya
- 3:15.130 3:15.131 3:15.132 3:15.133 3:15.134 3:15.135 | and the Rabbi, the author of Chok Ya'akov, and Maharshal, and Maharam al-Shakar, | and the Rabbi, the author of Chok Ya'akov and Maharshal and Maharam al-Shakar
- 3:15.136 3:15.137 3:15.138 3:15.139 3:15.140 3:15.141 3:15.142 3:15.143 3:15.144 3:15.145 3:15.146 3:15.147 3:15.148 3:15.149 | who say that one need not be concerned with the words of the Zohar and the Kabbalists, when they are against the Mishnah and the Talmud, and even against a single decisor. | who say that one need not be concerned about words the Zohar and the Kabbalists when they are against the Mishnah and the Talmud and even against a decisor single
- 3:15.150 3:15.151 3:15.152 3:15.153 3:15.154 3:15.155 3:15.156 3:15.157 3:15.158 3:15.159 3:15.160 3:15.161 3:15.162 | Behold, from this prove to yourself that you wrote that they are a law of Moses from Sinai, that your words are a fabrication and utter folly. | behold from this prove to yourself that you wrote that they are a law of Moses from Sinai that your words are a fabrication and folly utter

Grammar notes:
- גם הביא עוד ... וכתב עליו ... שכתב שכל דבר: the reporting verbs carry indirect speech throughout the paragraph, so the English has no quotation marks.
- שכתב שכל דבר ... אני מורה בו ... ולעצמי: the Radbaz's words are reported in indirect speech, and the first person ('I rule', 'I follow') is kept.
- הנה מזה הוכיח לעצמך: the closing turn is addressed to a second person, whom the author charges with calling these rulings 'a law of Moses from Sinai'.

Translator's note: The paragraph reports a sequence of decisors and ends with the author's address to his opponent. The base text is Wikisource.

Commentary rulings:
- apparatus: The printed 'הרב רבי הזהב' (3:15.108-110) is read as the author of Ravid ha-Zahav named at 3:15.5-7; the Hebrew is kept, and the reading is not certain.
- apparatus: Base text is Wikisource; a general line until the 1931 scan is collated (open question 2).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

