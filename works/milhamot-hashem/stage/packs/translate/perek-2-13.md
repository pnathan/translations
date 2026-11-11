# Translation pack: milhamot-hashem perek-2-13 (Milchamot ha-Shem 2:29–2:29)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-13`.

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

- **ספירה / ספירות**: fixed, "sefirah / sefirot". Not "emanations",
  "spheres" or "numbers". The names of the sefirot are transliterated, with a
  gloss at the first occurrence in a unit: Keter (Crown), Chokhmah,
  Binah, Chesed, Gevurah, Tif'eret, Netzach, Hod, Yesod, Malkhut. "לספירה
  הנקראת חסד שהיא מדת רחמים" (2:30) → "to the sefirah called Chesed, which is
  the attribute of mercy".
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
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to"; "כל המגשים בספירות" (2:7) → "whoever ascribes corporeality to the
  sefirot"; "מגשימים אפי' את האין סוף" (5:28) → "they ascribe corporeality
  even to Ein Sof". Not technical: "ולא באה לידי התגשמות" (1:2) → "and it
  did not come to realization". The Rambam's principle: "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- האלהי before a name: "the divine", with a note; never dropped.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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


## Milchamot ha-Shem 2:29

Hebrew (Wikisource): והארכתי כל זה לישב לך דברי כס"א, אבל קבלתי כבר זכרתיה למעלה. ויש הוכחה מכאן, שמה שאתה שואלני אינו מטעם בקשתך האמת כמו שאמרת, כי אם לבקש איזה עלילה. ומה אני שאתה מחזיק אותי למקובל, כי אני בער ולא אדע. וכל תפלותי אפי' כצפצוף העופות אינם, כי העופות ניזונים שלא בצער ואין עליהם לא עול מלכות ולא עול דרך ארץ, ואפשר צפצוף בכוונה אבל אני כמה עול הזמן וטרדת המזונות מבלבלים דעתי ומחשבותי. ועוד שלא קבלתי מחכם מקובל מפה אל פה. ומי יתן ויניח לי ה' מטרדות הזמן והמזונות ויזמין לי רב מקובל יניחני במעגלי יושר, ואזכה להיות מבני עליה. ומה שהשבתי לך פעם שניה הוא דברי כס"א שהעלמת עין מהם, ומה שכתבת שאינך מגשים בספירות, והלא כתבת א"כ יהיה כח בגוף. ובפעם הזאת אתה אומר כי האור גוף דק. וכבר כתבתי לך כי מה שכינו אותן בשם אור אינו אלא כדי לשבר את האזן פעם תגשים ופעם תתנצל ולמה עולים בדעתך כל אלו הספקות אלא מפני שאתה מגשים. ומה שכתבת מי הוא המנהיג כל אלו העולמות הנעלמים, הממציאם הוא מחיה אותם. ומה שכתבת שנמצא הז"א אלהי התחתונים. הלא כתבתי לך שאינו אלא כלי לפעולות הבורא. ומה שכתבת שמקצתם אמרו את כל מעשה ה' הגדול מכלל דאיכא קטן? האומרים כן כבר ביארו הם שכוונתם על מטטרון ששמו כשם רבו, כמו שאמרו בתלמוד ואל משה אמר עלה אל ה', אלי מבעי ליה אלא זה מטטרון ששמו כשם רבו וכ"ה בילקוט. ומה שכתבת שרבותינו האמתיים אמרו, הובא במנורת המאור, שבשעת מתן תורה קרע הקב"ה ז' רקיעים וכו'. ובמדרש רבה אמרו אנכי ה' אלהיך, אני ראשון ואני אחרון. אני ראשון שאין לי אב ואני אחרון שאין לי בן ומבלעדי אין אלהים שאין לי אח. מי הוא האומר שיש לו אב או בם או אח? אלא אתה מוציא דבה על החכמים שהם אומרים כן, והם כבר ביארו דבריהם ז"ל, שכל הכנויים והתארים הם בספירות לשבר את האזן. לך נא וראה מה שכתב מהר"י גיקטליא בהקדמת שערי אורה, והאלהי ר"מ ריקנטי, ור"י חיים וז"ל, סוף דבר כל מקום שתראה דברים שאין ראוי לאמרם בבורא ית' כגון שיעור קומה והדומים להם הכל נאמר על הספירות וכשתראה שידברו דרך שבח והודיה, הכל נאמר על הבורא ית' שבתוכן ובחוצה להן כי אין דבר שיגבילהו ועל כן אין לומר בבורא ית' לא ימין ולא שמאל לא פנים ולא אחור וכו' עי"ש. ומה שכתבת שדעת המקובלים שהספירות משתלשלים בטבע כבן מן האב, חלילה שהם אמרו כן, אדרבה אמרו חלילה להאמין ולהעלות על לב שהספירות הם חלק מהאין סוף שיצאו ממנו ונשתלשל מעילה לעלול כי הוא עון פלילי. שהאין סוף אינו מתחלק לחלקים ח"ו, ואינו מקבל תוספת ולא מגרעת אלא מציאותו קיים תמיד בלי שום שינוי כלל, והוא חידשם חידוש גמור.

| id | word | punctuation after |
|---|---|---|
| 2:29.1 | והארכתי |  |
| 2:29.2 | כל |  |
| 2:29.3 | זה |  |
| 2:29.4 | לישב |  |
| 2:29.5 | לך |  |
| 2:29.6 | דברי |  |
| 2:29.7 | כס"א | , |
| 2:29.8 | אבל |  |
| 2:29.9 | קבלתי |  |
| 2:29.10 | כבר |  |
| 2:29.11 | זכרתיה |  |
| 2:29.12 | למעלה | . |
| 2:29.13 | ויש |  |
| 2:29.14 | הוכחה |  |
| 2:29.15 | מכאן | , |
| 2:29.16 | שמה |  |
| 2:29.17 | שאתה |  |
| 2:29.18 | שואלני |  |
| 2:29.19 | אינו |  |
| 2:29.20 | מטעם |  |
| 2:29.21 | בקשתך |  |
| 2:29.22 | האמת |  |
| 2:29.23 | כמו |  |
| 2:29.24 | שאמרת | , |
| 2:29.25 | כי |  |
| 2:29.26 | אם |  |
| 2:29.27 | לבקש |  |
| 2:29.28 | איזה |  |
| 2:29.29 | עלילה | . |
| 2:29.30 | ומה |  |
| 2:29.31 | אני |  |
| 2:29.32 | שאתה |  |
| 2:29.33 | מחזיק |  |
| 2:29.34 | אותי |  |
| 2:29.35 | למקובל | , |
| 2:29.36 | כי |  |
| 2:29.37 | אני |  |
| 2:29.38 | בער |  |
| 2:29.39 | ולא |  |
| 2:29.40 | אדע | . |
| 2:29.41 | וכל |  |
| 2:29.42 | תפלותי |  |
| 2:29.43 | אפי | ' |
| 2:29.44 | כצפצוף |  |
| 2:29.45 | העופות |  |
| 2:29.46 | אינם | , |
| 2:29.47 | כי |  |
| 2:29.48 | העופות |  |
| 2:29.49 | ניזונים |  |
| 2:29.50 | שלא |  |
| 2:29.51 | בצער |  |
| 2:29.52 | ואין |  |
| 2:29.53 | עליהם |  |
| 2:29.54 | לא |  |
| 2:29.55 | עול |  |
| 2:29.56 | מלכות |  |
| 2:29.57 | ולא |  |
| 2:29.58 | עול |  |
| 2:29.59 | דרך |  |
| 2:29.60 | ארץ | , |
| 2:29.61 | ואפשר |  |
| 2:29.62 | צפצוף |  |
| 2:29.63 | בכוונה |  |
| 2:29.64 | אבל |  |
| 2:29.65 | אני |  |
| 2:29.66 | כמה |  |
| 2:29.67 | עול |  |
| 2:29.68 | הזמן |  |
| 2:29.69 | וטרדת |  |
| 2:29.70 | המזונות |  |
| 2:29.71 | מבלבלים |  |
| 2:29.72 | דעתי |  |
| 2:29.73 | ומחשבותי | . |
| 2:29.74 | ועוד |  |
| 2:29.75 | שלא |  |
| 2:29.76 | קבלתי |  |
| 2:29.77 | מחכם |  |
| 2:29.78 | מקובל |  |
| 2:29.79 | מפה |  |
| 2:29.80 | אל |  |
| 2:29.81 | פה | . |
| 2:29.82 | ומי |  |
| 2:29.83 | יתן |  |
| 2:29.84 | ויניח |  |
| 2:29.85 | לי |  |
| 2:29.86 | ה | ' |
| 2:29.87 | מטרדות |  |
| 2:29.88 | הזמן |  |
| 2:29.89 | והמזונות |  |
| 2:29.90 | ויזמין |  |
| 2:29.91 | לי |  |
| 2:29.92 | רב |  |
| 2:29.93 | מקובל |  |
| 2:29.94 | יניחני |  |
| 2:29.95 | במעגלי |  |
| 2:29.96 | יושר | , |
| 2:29.97 | ואזכה |  |
| 2:29.98 | להיות |  |
| 2:29.99 | מבני |  |
| 2:29.100 | עליה | . |
| 2:29.101 | ומה |  |
| 2:29.102 | שהשבתי |  |
| 2:29.103 | לך |  |
| 2:29.104 | פעם |  |
| 2:29.105 | שניה |  |
| 2:29.106 | הוא |  |
| 2:29.107 | דברי |  |
| 2:29.108 | כס"א |  |
| 2:29.109 | שהעלמת |  |
| 2:29.110 | עין |  |
| 2:29.111 | מהם | , |
| 2:29.112 | ומה |  |
| 2:29.113 | שכתבת |  |
| 2:29.114 | שאינך |  |
| 2:29.115 | מגשים |  |
| 2:29.116 | בספירות | , |
| 2:29.117 | והלא |  |
| 2:29.118 | כתבת |  |
| 2:29.119 | א"כ |  |
| 2:29.120 | יהיה |  |
| 2:29.121 | כח |  |
| 2:29.122 | בגוף | . |
| 2:29.123 | ובפעם |  |
| 2:29.124 | הזאת |  |
| 2:29.125 | אתה |  |
| 2:29.126 | אומר |  |
| 2:29.127 | כי |  |
| 2:29.128 | האור |  |
| 2:29.129 | גוף |  |
| 2:29.130 | דק | . |
| 2:29.131 | וכבר |  |
| 2:29.132 | כתבתי |  |
| 2:29.133 | לך |  |
| 2:29.134 | כי |  |
| 2:29.135 | מה |  |
| 2:29.136 | שכינו |  |
| 2:29.137 | אותן |  |
| 2:29.138 | בשם |  |
| 2:29.139 | אור |  |
| 2:29.140 | אינו |  |
| 2:29.141 | אלא |  |
| 2:29.142 | כדי |  |
| 2:29.143 | לשבר |  |
| 2:29.144 | את |  |
| 2:29.145 | האזן |  |
| 2:29.146 | פעם |  |
| 2:29.147 | תגשים |  |
| 2:29.148 | ופעם |  |
| 2:29.149 | תתנצל |  |
| 2:29.150 | ולמה |  |
| 2:29.151 | עולים |  |
| 2:29.152 | בדעתך |  |
| 2:29.153 | כל |  |
| 2:29.154 | אלו |  |
| 2:29.155 | הספקות |  |
| 2:29.156 | אלא |  |
| 2:29.157 | מפני |  |
| 2:29.158 | שאתה |  |
| 2:29.159 | מגשים | . |
| 2:29.160 | ומה |  |
| 2:29.161 | שכתבת |  |
| 2:29.162 | מי |  |
| 2:29.163 | הוא |  |
| 2:29.164 | המנהיג |  |
| 2:29.165 | כל |  |
| 2:29.166 | אלו |  |
| 2:29.167 | העולמות |  |
| 2:29.168 | הנעלמים | , |
| 2:29.169 | הממציאם |  |
| 2:29.170 | הוא |  |
| 2:29.171 | מחיה |  |
| 2:29.172 | אותם | . |
| 2:29.173 | ומה |  |
| 2:29.174 | שכתבת |  |
| 2:29.175 | שנמצא |  |
| 2:29.176 | הז"א |  |
| 2:29.177 | אלהי |  |
| 2:29.178 | התחתונים | . |
| 2:29.179 | הלא |  |
| 2:29.180 | כתבתי |  |
| 2:29.181 | לך |  |
| 2:29.182 | שאינו |  |
| 2:29.183 | אלא |  |
| 2:29.184 | כלי |  |
| 2:29.185 | לפעולות |  |
| 2:29.186 | הבורא | . |
| 2:29.187 | ומה |  |
| 2:29.188 | שכתבת |  |
| 2:29.189 | שמקצתם |  |
| 2:29.190 | אמרו |  |
| 2:29.191 | את |  |
| 2:29.192 | כל |  |
| 2:29.193 | מעשה |  |
| 2:29.194 | ה | ' |
| 2:29.195 | הגדול |  |
| 2:29.196 | מכלל |  |
| 2:29.197 | דאיכא |  |
| 2:29.198 | קטן | ? |
| 2:29.199 | האומרים |  |
| 2:29.200 | כן |  |
| 2:29.201 | כבר |  |
| 2:29.202 | ביארו |  |
| 2:29.203 | הם |  |
| 2:29.204 | שכוונתם |  |
| 2:29.205 | על |  |
| 2:29.206 | מטטרון |  |
| 2:29.207 | ששמו |  |
| 2:29.208 | כשם |  |
| 2:29.209 | רבו | , |
| 2:29.210 | כמו |  |
| 2:29.211 | שאמרו |  |
| 2:29.212 | בתלמוד |  |
| 2:29.213 | ואל |  |
| 2:29.214 | משה |  |
| 2:29.215 | אמר |  |
| 2:29.216 | עלה |  |
| 2:29.217 | אל |  |
| 2:29.218 | ה | ', |
| 2:29.219 | אלי |  |
| 2:29.220 | מבעי |  |
| 2:29.221 | ליה |  |
| 2:29.222 | אלא |  |
| 2:29.223 | זה |  |
| 2:29.224 | מטטרון |  |
| 2:29.225 | ששמו |  |
| 2:29.226 | כשם |  |
| 2:29.227 | רבו |  |
| 2:29.228 | וכ"ה |  |
| 2:29.229 | בילקוט | . |
| 2:29.230 | ומה |  |
| 2:29.231 | שכתבת |  |
| 2:29.232 | שרבותינו |  |
| 2:29.233 | האמתיים |  |
| 2:29.234 | אמרו | , |
| 2:29.235 | הובא |  |
| 2:29.236 | במנורת |  |
| 2:29.237 | המאור | , |
| 2:29.238 | שבשעת |  |
| 2:29.239 | מתן |  |
| 2:29.240 | תורה |  |
| 2:29.241 | קרע |  |
| 2:29.242 | הקב"ה |  |
| 2:29.243 | ז | ' |
| 2:29.244 | רקיעים |  |
| 2:29.245 | וכו | '. |
| 2:29.246 | ובמדרש |  |
| 2:29.247 | רבה |  |
| 2:29.248 | אמרו |  |
| 2:29.249 | אנכי |  |
| 2:29.250 | ה | ' |
| 2:29.251 | אלהיך | , |
| 2:29.252 | אני |  |
| 2:29.253 | ראשון |  |
| 2:29.254 | ואני |  |
| 2:29.255 | אחרון | . |
| 2:29.256 | אני |  |
| 2:29.257 | ראשון |  |
| 2:29.258 | שאין |  |
| 2:29.259 | לי |  |
| 2:29.260 | אב |  |
| 2:29.261 | ואני |  |
| 2:29.262 | אחרון |  |
| 2:29.263 | שאין |  |
| 2:29.264 | לי |  |
| 2:29.265 | בן |  |
| 2:29.266 | ומבלעדי |  |
| 2:29.267 | אין |  |
| 2:29.268 | אלהים |  |
| 2:29.269 | שאין |  |
| 2:29.270 | לי |  |
| 2:29.271 | אח | . |
| 2:29.272 | מי |  |
| 2:29.273 | הוא |  |
| 2:29.274 | האומר |  |
| 2:29.275 | שיש |  |
| 2:29.276 | לו |  |
| 2:29.277 | אב |  |
| 2:29.278 | או |  |
| 2:29.279 | בם |  |
| 2:29.280 | או |  |
| 2:29.281 | אח | ? |
| 2:29.282 | אלא |  |
| 2:29.283 | אתה |  |
| 2:29.284 | מוציא |  |
| 2:29.285 | דבה |  |
| 2:29.286 | על |  |
| 2:29.287 | החכמים |  |
| 2:29.288 | שהם |  |
| 2:29.289 | אומרים |  |
| 2:29.290 | כן | , |
| 2:29.291 | והם |  |
| 2:29.292 | כבר |  |
| 2:29.293 | ביארו |  |
| 2:29.294 | דבריהם |  |
| 2:29.295 | ז"ל | , |
| 2:29.296 | שכל |  |
| 2:29.297 | הכנויים |  |
| 2:29.298 | והתארים |  |
| 2:29.299 | הם |  |
| 2:29.300 | בספירות |  |
| 2:29.301 | לשבר |  |
| 2:29.302 | את |  |
| 2:29.303 | האזן | . |
| 2:29.304 | לך |  |
| 2:29.305 | נא |  |
| 2:29.306 | וראה |  |
| 2:29.307 | מה |  |
| 2:29.308 | שכתב |  |
| 2:29.309 | מהר"י |  |
| 2:29.310 | גיקטליא |  |
| 2:29.311 | בהקדמת |  |
| 2:29.312 | שערי |  |
| 2:29.313 | אורה | , |
| 2:29.314 | והאלהי |  |
| 2:29.315 | ר"מ |  |
| 2:29.316 | ריקנטי | , |
| 2:29.317 | ור"י |  |
| 2:29.318 | חיים |  |
| 2:29.319 | וז"ל | , |
| 2:29.320 | סוף |  |
| 2:29.321 | דבר |  |
| 2:29.322 | כל |  |
| 2:29.323 | מקום |  |
| 2:29.324 | שתראה |  |
| 2:29.325 | דברים |  |
| 2:29.326 | שאין |  |
| 2:29.327 | ראוי |  |
| 2:29.328 | לאמרם |  |
| 2:29.329 | בבורא |  |
| 2:29.330 | ית | ' |
| 2:29.331 | כגון |  |
| 2:29.332 | שיעור |  |
| 2:29.333 | קומה |  |
| 2:29.334 | והדומים |  |
| 2:29.335 | להם |  |
| 2:29.336 | הכל |  |
| 2:29.337 | נאמר |  |
| 2:29.338 | על |  |
| 2:29.339 | הספירות |  |
| 2:29.340 | וכשתראה |  |
| 2:29.341 | שידברו |  |
| 2:29.342 | דרך |  |
| 2:29.343 | שבח |  |
| 2:29.344 | והודיה | , |
| 2:29.345 | הכל |  |
| 2:29.346 | נאמר |  |
| 2:29.347 | על |  |
| 2:29.348 | הבורא |  |
| 2:29.349 | ית | ' |
| 2:29.350 | שבתוכן |  |
| 2:29.351 | ובחוצה |  |
| 2:29.352 | להן |  |
| 2:29.353 | כי |  |
| 2:29.354 | אין |  |
| 2:29.355 | דבר |  |
| 2:29.356 | שיגבילהו |  |
| 2:29.357 | ועל |  |
| 2:29.358 | כן |  |
| 2:29.359 | אין |  |
| 2:29.360 | לומר |  |
| 2:29.361 | בבורא |  |
| 2:29.362 | ית | ' |
| 2:29.363 | לא |  |
| 2:29.364 | ימין |  |
| 2:29.365 | ולא |  |
| 2:29.366 | שמאל |  |
| 2:29.367 | לא |  |
| 2:29.368 | פנים |  |
| 2:29.369 | ולא |  |
| 2:29.370 | אחור |  |
| 2:29.371 | וכו | ' |
| 2:29.372 | עי"ש | . |
| 2:29.373 | ומה |  |
| 2:29.374 | שכתבת |  |
| 2:29.375 | שדעת |  |
| 2:29.376 | המקובלים |  |
| 2:29.377 | שהספירות |  |
| 2:29.378 | משתלשלים |  |
| 2:29.379 | בטבע |  |
| 2:29.380 | כבן |  |
| 2:29.381 | מן |  |
| 2:29.382 | האב | , |
| 2:29.383 | חלילה |  |
| 2:29.384 | שהם |  |
| 2:29.385 | אמרו |  |
| 2:29.386 | כן | , |
| 2:29.387 | אדרבה |  |
| 2:29.388 | אמרו |  |
| 2:29.389 | חלילה |  |
| 2:29.390 | להאמין |  |
| 2:29.391 | ולהעלות |  |
| 2:29.392 | על |  |
| 2:29.393 | לב |  |
| 2:29.394 | שהספירות |  |
| 2:29.395 | הם |  |
| 2:29.396 | חלק |  |
| 2:29.397 | מהאין |  |
| 2:29.398 | סוף |  |
| 2:29.399 | שיצאו |  |
| 2:29.400 | ממנו |  |
| 2:29.401 | ונשתלשל |  |
| 2:29.402 | מעילה |  |
| 2:29.403 | לעלול |  |
| 2:29.404 | כי |  |
| 2:29.405 | הוא |  |
| 2:29.406 | עון |  |
| 2:29.407 | פלילי | . |
| 2:29.408 | שהאין |  |
| 2:29.409 | סוף |  |
| 2:29.410 | אינו |  |
| 2:29.411 | מתחלק |  |
| 2:29.412 | לחלקים |  |
| 2:29.413 | ח"ו | , |
| 2:29.414 | ואינו |  |
| 2:29.415 | מקבל |  |
| 2:29.416 | תוספת |  |
| 2:29.417 | ולא |  |
| 2:29.418 | מגרעת |  |
| 2:29.419 | אלא |  |
| 2:29.420 | מציאותו |  |
| 2:29.421 | קיים |  |
| 2:29.422 | תמיד |  |
| 2:29.423 | בלי |  |
| 2:29.424 | שום |  |
| 2:29.425 | שינוי |  |
| 2:29.426 | כלל | , |
| 2:29.427 | והוא |  |
| 2:29.428 | חידשם |  |
| 2:29.429 | חידוש |  |
| 2:29.430 | גמור | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

