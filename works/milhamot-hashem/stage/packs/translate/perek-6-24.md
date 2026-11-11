# Translation pack: milhamot-hashem perek-6-24 (Milchamot ha-Shem 6:37–6:37)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-24.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-24`.

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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- יש מאין / יש מיש: translate, "something from nothing" / "something from
  something"; no transliteration.
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- Quotation formulas also include שנינו, הכי קאמר, אמר ר' X and the like: a passage
  so introduced takes quotation marks when the author gives a source's own
  words; a paraphrase takes none.
- דמי: "blood", or "price/value" where the context is a sale (דמי השור); the
  other sense in a token note.
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דמיון: "analogy" for an argument, "likeness" for resemblance.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- והוא שלא: "provided that" (fixed).
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


## Milchamot ha-Shem 6:37

Hebrew (Wikisource): וכדי שלא יטעה הפתי הסכל מסכלותו ופתיותיו ומיעוט הבנתו ויחשוב שהשי"ת הוא על דרך הנראה ממליצת הספר שיש לו גשמות, הזהירנו בתורה שנשמר מאד מלחשוב שיש לבורא ית' צורה ולא דמיון. והוא אמרו ונשמרתם מאד לנפשותיכם כי לא ראיתם כל תמונה. והנביאים הזהירונו להמנע מלדמותו בשכל, ומלהמשילו לרעיון כאמרו ואל מי תדמיון אל, ומה דמות תערכו לו וכתוב ואל מי תדמיוני ואשוה. וכל זה צריך בעל אמונתינו לידע, ואז יוכל להבדיל בין הבורא וזולתו, ויהיה מיחד את השי"ת יחוד גמור כי מי שאינו יודע דרכי ההבדל אינו יכול ליחדו ית' יחוד אמתי. אמנם כאשר ידענו על מתכונתו אז ותקרב לאמתת בוראו ותתקרב נפשו להשגת המציאה. ותתעלה בהשגתה על מה כונתה כשאין אדם זוכה לידע אמתת בוראו אפי' יהיה מיחד אותו בלבו הרי הוא כאלו יש מחיצה בינו ובין בוראו ית'. כי זה הוא דבר אשר יגיע לו הפתי או הקטן. אשר איננו יודע דרכי האמונה ועניינה. ואין אמתתה קבועה בלבו. אבל כל עוד שמתחיל לחקור לידע אמתתו ית'. אזי המחיצה הולכת ומתמעטת. והנפש הולכת ומתקרבת אל בוראה יתעלה. ותתעלה נפשו ושכלו בתענו ג השגתו ית'. כי בהיות הנפש מתעסקת בדברים שכליים אז היא נמשכת אחריהם. לפי שהיא גם כן ממקור השכל נצמאת ולכן היא משתוקקת ונכספה להדמות לבוראה ית' ואינה יכולה להבדל ממנו עד שנמצאת דבוקה בו. ועל כיוצא בענין זה אמר ית' ובו תדבק. וגם על זה אמר הכתוב קרוב ה' לכל קוראיו. כי מאחר שהבורא ית' אינו גוף ולא כח בגוף אי אפשר היותו קרוב או רחוק מן הנפש. אלא רחוק מצד המחיצה המבדלת שהיא הסכלות כשאין לנפש ידיעה באמתתו ית'. וקרוב לנפש מצד הידיעה וההשגה. ולכן כל בעל נפש צריך לתור ולבקש על דרך ידיעת אמיתת מציאות הבורא ית'. כי הוא עיקר ויסוד גדול לאמונה. ועל דבר זה צוה אותנו בתורה ואמר וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:37.1 | וכדי |  |
| 6:37.2 | שלא |  |
| 6:37.3 | יטעה |  |
| 6:37.4 | הפתי |  |
| 6:37.5 | הסכל |  |
| 6:37.6 | מסכלותו |  |
| 6:37.7 | ופתיותיו |  |
| 6:37.8 | ומיעוט |  |
| 6:37.9 | הבנתו |  |
| 6:37.10 | ויחשוב |  |
| 6:37.11 | שהשי"ת |  |
| 6:37.12 | הוא |  |
| 6:37.13 | על |  |
| 6:37.14 | דרך |  |
| 6:37.15 | הנראה |  |
| 6:37.16 | ממליצת |  |
| 6:37.17 | הספר |  |
| 6:37.18 | שיש |  |
| 6:37.19 | לו |  |
| 6:37.20 | גשמות | , |
| 6:37.21 | הזהירנו |  |
| 6:37.22 | בתורה |  |
| 6:37.23 | שנשמר |  |
| 6:37.24 | מאד |  |
| 6:37.25 | מלחשוב |  |
| 6:37.26 | שיש |  |
| 6:37.27 | לבורא |  |
| 6:37.28 | ית | ' |
| 6:37.29 | צורה |  |
| 6:37.30 | ולא |  |
| 6:37.31 | דמיון | . |
| 6:37.32 | והוא |  |
| 6:37.33 | אמרו |  |
| 6:37.34 | ונשמרתם |  |
| 6:37.35 | מאד |  |
| 6:37.36 | לנפשותיכם |  |
| 6:37.37 | כי |  |
| 6:37.38 | לא |  |
| 6:37.39 | ראיתם |  |
| 6:37.40 | כל |  |
| 6:37.41 | תמונה | . |
| 6:37.42 | והנביאים |  |
| 6:37.43 | הזהירונו |  |
| 6:37.44 | להמנע |  |
| 6:37.45 | מלדמותו |  |
| 6:37.46 | בשכל | , |
| 6:37.47 | ומלהמשילו |  |
| 6:37.48 | לרעיון |  |
| 6:37.49 | כאמרו |  |
| 6:37.50 | ואל |  |
| 6:37.51 | מי |  |
| 6:37.52 | תדמיון |  |
| 6:37.53 | אל | , |
| 6:37.54 | ומה |  |
| 6:37.55 | דמות |  |
| 6:37.56 | תערכו |  |
| 6:37.57 | לו |  |
| 6:37.58 | וכתוב |  |
| 6:37.59 | ואל |  |
| 6:37.60 | מי |  |
| 6:37.61 | תדמיוני |  |
| 6:37.62 | ואשוה | . |
| 6:37.63 | וכל |  |
| 6:37.64 | זה |  |
| 6:37.65 | צריך |  |
| 6:37.66 | בעל |  |
| 6:37.67 | אמונתינו |  |
| 6:37.68 | לידע | , |
| 6:37.69 | ואז |  |
| 6:37.70 | יוכל |  |
| 6:37.71 | להבדיל |  |
| 6:37.72 | בין |  |
| 6:37.73 | הבורא |  |
| 6:37.74 | וזולתו | , |
| 6:37.75 | ויהיה |  |
| 6:37.76 | מיחד |  |
| 6:37.77 | את |  |
| 6:37.78 | השי"ת |  |
| 6:37.79 | יחוד |  |
| 6:37.80 | גמור |  |
| 6:37.81 | כי |  |
| 6:37.82 | מי |  |
| 6:37.83 | שאינו |  |
| 6:37.84 | יודע |  |
| 6:37.85 | דרכי |  |
| 6:37.86 | ההבדל |  |
| 6:37.87 | אינו |  |
| 6:37.88 | יכול |  |
| 6:37.89 | ליחדו |  |
| 6:37.90 | ית | ' |
| 6:37.91 | יחוד |  |
| 6:37.92 | אמתי | . |
| 6:37.93 | אמנם |  |
| 6:37.94 | כאשר |  |
| 6:37.95 | ידענו |  |
| 6:37.96 | על |  |
| 6:37.97 | מתכונתו |  |
| 6:37.98 | אז |  |
| 6:37.99 | ותקרב |  |
| 6:37.100 | לאמתת |  |
| 6:37.101 | בוראו |  |
| 6:37.102 | ותתקרב |  |
| 6:37.103 | נפשו |  |
| 6:37.104 | להשגת |  |
| 6:37.105 | המציאה | . |
| 6:37.106 | ותתעלה |  |
| 6:37.107 | בהשגתה |  |
| 6:37.108 | על |  |
| 6:37.109 | מה |  |
| 6:37.110 | כונתה |  |
| 6:37.111 | כשאין |  |
| 6:37.112 | אדם |  |
| 6:37.113 | זוכה |  |
| 6:37.114 | לידע |  |
| 6:37.115 | אמתת |  |
| 6:37.116 | בוראו |  |
| 6:37.117 | אפי | ' |
| 6:37.118 | יהיה |  |
| 6:37.119 | מיחד |  |
| 6:37.120 | אותו |  |
| 6:37.121 | בלבו |  |
| 6:37.122 | הרי |  |
| 6:37.123 | הוא |  |
| 6:37.124 | כאלו |  |
| 6:37.125 | יש |  |
| 6:37.126 | מחיצה |  |
| 6:37.127 | בינו |  |
| 6:37.128 | ובין |  |
| 6:37.129 | בוראו |  |
| 6:37.130 | ית | '. |
| 6:37.131 | כי |  |
| 6:37.132 | זה |  |
| 6:37.133 | הוא |  |
| 6:37.134 | דבר |  |
| 6:37.135 | אשר |  |
| 6:37.136 | יגיע |  |
| 6:37.137 | לו |  |
| 6:37.138 | הפתי |  |
| 6:37.139 | או |  |
| 6:37.140 | הקטן | . |
| 6:37.141 | אשר |  |
| 6:37.142 | איננו |  |
| 6:37.143 | יודע |  |
| 6:37.144 | דרכי |  |
| 6:37.145 | האמונה |  |
| 6:37.146 | ועניינה | . |
| 6:37.147 | ואין |  |
| 6:37.148 | אמתתה |  |
| 6:37.149 | קבועה |  |
| 6:37.150 | בלבו | . |
| 6:37.151 | אבל |  |
| 6:37.152 | כל |  |
| 6:37.153 | עוד |  |
| 6:37.154 | שמתחיל |  |
| 6:37.155 | לחקור |  |
| 6:37.156 | לידע |  |
| 6:37.157 | אמתתו |  |
| 6:37.158 | ית | '. |
| 6:37.159 | אזי |  |
| 6:37.160 | המחיצה |  |
| 6:37.161 | הולכת |  |
| 6:37.162 | ומתמעטת | . |
| 6:37.163 | והנפש |  |
| 6:37.164 | הולכת |  |
| 6:37.165 | ומתקרבת |  |
| 6:37.166 | אל |  |
| 6:37.167 | בוראה |  |
| 6:37.168 | יתעלה | . |
| 6:37.169 | ותתעלה |  |
| 6:37.170 | נפשו |  |
| 6:37.171 | ושכלו |  |
| 6:37.172 | בתענו |  |
| 6:37.173 | ג |  |
| 6:37.174 | השגתו |  |
| 6:37.175 | ית | '. |
| 6:37.176 | כי |  |
| 6:37.177 | בהיות |  |
| 6:37.178 | הנפש |  |
| 6:37.179 | מתעסקת |  |
| 6:37.180 | בדברים |  |
| 6:37.181 | שכליים |  |
| 6:37.182 | אז |  |
| 6:37.183 | היא |  |
| 6:37.184 | נמשכת |  |
| 6:37.185 | אחריהם | . |
| 6:37.186 | לפי |  |
| 6:37.187 | שהיא |  |
| 6:37.188 | גם |  |
| 6:37.189 | כן |  |
| 6:37.190 | ממקור |  |
| 6:37.191 | השכל |  |
| 6:37.192 | נצמאת |  |
| 6:37.193 | ולכן |  |
| 6:37.194 | היא |  |
| 6:37.195 | משתוקקת |  |
| 6:37.196 | ונכספה |  |
| 6:37.197 | להדמות |  |
| 6:37.198 | לבוראה |  |
| 6:37.199 | ית | ' |
| 6:37.200 | ואינה |  |
| 6:37.201 | יכולה |  |
| 6:37.202 | להבדל |  |
| 6:37.203 | ממנו |  |
| 6:37.204 | עד |  |
| 6:37.205 | שנמצאת |  |
| 6:37.206 | דבוקה |  |
| 6:37.207 | בו | . |
| 6:37.208 | ועל |  |
| 6:37.209 | כיוצא |  |
| 6:37.210 | בענין |  |
| 6:37.211 | זה |  |
| 6:37.212 | אמר |  |
| 6:37.213 | ית | ' |
| 6:37.214 | ובו |  |
| 6:37.215 | תדבק | . |
| 6:37.216 | וגם |  |
| 6:37.217 | על |  |
| 6:37.218 | זה |  |
| 6:37.219 | אמר |  |
| 6:37.220 | הכתוב |  |
| 6:37.221 | קרוב |  |
| 6:37.222 | ה | ' |
| 6:37.223 | לכל |  |
| 6:37.224 | קוראיו | . |
| 6:37.225 | כי |  |
| 6:37.226 | מאחר |  |
| 6:37.227 | שהבורא |  |
| 6:37.228 | ית | ' |
| 6:37.229 | אינו |  |
| 6:37.230 | גוף |  |
| 6:37.231 | ולא |  |
| 6:37.232 | כח |  |
| 6:37.233 | בגוף |  |
| 6:37.234 | אי |  |
| 6:37.235 | אפשר |  |
| 6:37.236 | היותו |  |
| 6:37.237 | קרוב |  |
| 6:37.238 | או |  |
| 6:37.239 | רחוק |  |
| 6:37.240 | מן |  |
| 6:37.241 | הנפש | . |
| 6:37.242 | אלא |  |
| 6:37.243 | רחוק |  |
| 6:37.244 | מצד |  |
| 6:37.245 | המחיצה |  |
| 6:37.246 | המבדלת |  |
| 6:37.247 | שהיא |  |
| 6:37.248 | הסכלות |  |
| 6:37.249 | כשאין |  |
| 6:37.250 | לנפש |  |
| 6:37.251 | ידיעה |  |
| 6:37.252 | באמתתו |  |
| 6:37.253 | ית | '. |
| 6:37.254 | וקרוב |  |
| 6:37.255 | לנפש |  |
| 6:37.256 | מצד |  |
| 6:37.257 | הידיעה |  |
| 6:37.258 | וההשגה | . |
| 6:37.259 | ולכן |  |
| 6:37.260 | כל |  |
| 6:37.261 | בעל |  |
| 6:37.262 | נפש |  |
| 6:37.263 | צריך |  |
| 6:37.264 | לתור |  |
| 6:37.265 | ולבקש |  |
| 6:37.266 | על |  |
| 6:37.267 | דרך |  |
| 6:37.268 | ידיעת |  |
| 6:37.269 | אמיתת |  |
| 6:37.270 | מציאות |  |
| 6:37.271 | הבורא |  |
| 6:37.272 | ית | '. |
| 6:37.273 | כי |  |
| 6:37.274 | הוא |  |
| 6:37.275 | עיקר |  |
| 6:37.276 | ויסוד |  |
| 6:37.277 | גדול |  |
| 6:37.278 | לאמונה | . |
| 6:37.279 | ועל |  |
| 6:37.280 | דבר |  |
| 6:37.281 | זה |  |
| 6:37.282 | צוה |  |
| 6:37.283 | אותנו |  |
| 6:37.284 | בתורה |  |
| 6:37.285 | ואמר |  |
| 6:37.286 | וידעת |  |
| 6:37.287 | היום |  |
| 6:37.288 | והשבות |  |
| 6:37.289 | אל |  |
| 6:37.290 | לבבך |  |
| 6:37.291 | כי |  |
| 6:37.292 | ה | ' |
| 6:37.293 | הוא |  |
| 6:37.294 | האלהים |  |
| 6:37.295 | בשמים |  |
| 6:37.296 | ממעל |  |
| 6:37.297 | ועל |  |
| 6:37.298 | הארץ |  |
| 6:37.299 | מתחת |  |
| 6:37.300 | אין |  |
| 6:37.301 | עוד |  |
| 6:37.302 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

