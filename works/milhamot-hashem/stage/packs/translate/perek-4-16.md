# Translation pack: milhamot-hashem perek-4-16 (Milchamot ha-Shem 4:26–4:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-16`.

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
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- והוא שלא: "provided that" (fixed).
- An unexplained word in a quoted Judeo-Arabic passage (שברא): translate what
  can be read, mark the rest "[unclear]" in the English, and give the Hebrew
  letters in `apparatus`.
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
- Wikisource additions (Q2): the pack marks none; the translator does not
  add a flag per word. A general line in `apparatus` on the unit's first
  segment is enough until the 1931 scan is collated.
- תשבורת: "mathematics" (mensuration); ותוכנים: "astronomy" (plural of תכונה).
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


## Milchamot ha-Shem 4:26

Hebrew (Wikisource): לד) הרח"ו בספר עץ חיים הורה היתר לעצמו לדבר ולדרוש במה לפנים ולבאר בה ראשי פרקים. וחשב ליתן טעם למה נברא העולם בזמן זה ולא קודם לכן, שבזמן הקדום נתעסק הקב"ה בבריאת עולמות של מעלה, ולא הספיק לגמרן אלא עד שהגיע זמן בריאת העה"ז. ולא מצא לו הבורא ח"ו עת וזמן פנוי לברוא את העה"ז, כי כל משך הזמן הקדום היה ח"ו טרוד בהתעסקו בברית העולמות העליונים לפי שבהכרח צריך עדן ועדנין לבריאת העולם ע"ש.

| id | word | punctuation after |
|---|---|---|
| 4:26.1 | לד | ) |
| 4:26.2 | הרח"ו |  |
| 4:26.3 | בספר |  |
| 4:26.4 | עץ |  |
| 4:26.5 | חיים |  |
| 4:26.6 | הורה |  |
| 4:26.7 | היתר |  |
| 4:26.8 | לעצמו |  |
| 4:26.9 | לדבר |  |
| 4:26.10 | ולדרוש |  |
| 4:26.11 | במה |  |
| 4:26.12 | לפנים |  |
| 4:26.13 | ולבאר |  |
| 4:26.14 | בה |  |
| 4:26.15 | ראשי |  |
| 4:26.16 | פרקים | . |
| 4:26.17 | וחשב |  |
| 4:26.18 | ליתן |  |
| 4:26.19 | טעם |  |
| 4:26.20 | למה |  |
| 4:26.21 | נברא |  |
| 4:26.22 | העולם |  |
| 4:26.23 | בזמן |  |
| 4:26.24 | זה |  |
| 4:26.25 | ולא |  |
| 4:26.26 | קודם |  |
| 4:26.27 | לכן | , |
| 4:26.28 | שבזמן |  |
| 4:26.29 | הקדום |  |
| 4:26.30 | נתעסק |  |
| 4:26.31 | הקב"ה |  |
| 4:26.32 | בבריאת |  |
| 4:26.33 | עולמות |  |
| 4:26.34 | של |  |
| 4:26.35 | מעלה | , |
| 4:26.36 | ולא |  |
| 4:26.37 | הספיק |  |
| 4:26.38 | לגמרן |  |
| 4:26.39 | אלא |  |
| 4:26.40 | עד |  |
| 4:26.41 | שהגיע |  |
| 4:26.42 | זמן |  |
| 4:26.43 | בריאת |  |
| 4:26.44 | העה"ז | . |
| 4:26.45 | ולא |  |
| 4:26.46 | מצא |  |
| 4:26.47 | לו |  |
| 4:26.48 | הבורא |  |
| 4:26.49 | ח"ו |  |
| 4:26.50 | עת |  |
| 4:26.51 | וזמן |  |
| 4:26.52 | פנוי |  |
| 4:26.53 | לברוא |  |
| 4:26.54 | את |  |
| 4:26.55 | העה"ז | , |
| 4:26.56 | כי |  |
| 4:26.57 | כל |  |
| 4:26.58 | משך |  |
| 4:26.59 | הזמן |  |
| 4:26.60 | הקדום |  |
| 4:26.61 | היה |  |
| 4:26.62 | ח"ו |  |
| 4:26.63 | טרוד |  |
| 4:26.64 | בהתעסקו |  |
| 4:26.65 | בברית |  |
| 4:26.66 | העולמות |  |
| 4:26.67 | העליונים |  |
| 4:26.68 | לפי |  |
| 4:26.69 | שבהכרח |  |
| 4:26.70 | צריך |  |
| 4:26.71 | עדן |  |
| 4:26.72 | ועדנין |  |
| 4:26.73 | לבריאת |  |
| 4:26.74 | העולם |  |
| 4:26.75 | ע"ש | . |


## Milchamot ha-Shem 4:27

Hebrew (Wikisource): ונעלם מעיני הרח"ו מה שכתב הרמב"ם ז"ל במורה סוף פרק ט"ו משני וז"ל כי אלו תאמר על דרך משל כי הבורא ברא עולמות רבות קודם זה העולם כפי מנין מלוא רוחב גלגל העליון חרדל, וכל עולם ועולם מהן עמד שנים רבות בלי מספר כמלוא רוחב גלגל העליון (חרדל), היה זה כשנעריכהו אל מציאותו יתברך שאין לו תכלית כמו אלו, אמרת כי האל ברא אתמול העולם, כי כשנקיים התחלת מציאות אחר העדר (ג"א דאפיסה) אין הפרש בין שתשים זה קודם מאות אלפים שנים או מזמן קרוב מאוד ע"כ. ומזה תבין חסרון חכמה מהרח"ו שחשב להחכם מבלי חכמה. ובספר עוז לאלהים בבית קדש הקדשים פ' כ"א ד' ס"ז ע"ד העיז פניו להשיב וז"ל ואם לא תתרצה בתשובת הבחירה נשיב לך שמעת שנתהוה מלכא קדישא ברא העולמות, ובזה לא תוכל להקשות למה בראם עכשו ולא קודם, כי מעת שנתהוה בראם עכ"ל. ונמצא לפי דבריו שאין הבורא יתברך גדל מאדם הראשון אלא בחמשה ימים בלבד. ולדבריו ג"כ איך יצוייר מה שדרשו ז"ל על פסוק ואהיה אצלו אמון, שהתורה קדמה לבריאת העולם תתעק"ד דורות, כי לדבריו היתה התורה קודם התהוות הבורא, רחמנא לצלן ממחשבות זרות כאלה. וז"ל מהר'י צהרי בשם ר' סעדיה גאון פרשה בראשית הענין השביעי בא להודיענו, שאעפ"י שתראו השמים גדולי הערך במדתם כפי מה שנתבאר במופת לבעלי חכמת התשבורת בשמים ובארץ, אל תחשבו שהוצרכו זמן רב לבריאתם, אלא קורא אני להם יעמדו יחדו, ורוצה לומר בראשית בזמן קטן שבזמנים בלי עמל ובלי טורח ובלי יגיעה, וכן הוא אומר לא ייעף ולא ייגע אין חקר לתבונתו עכ"ל. ובבראשית רבה פ' א' וג' ובילקוט בראשית אמז ה' וישעיה שכ"ב, ותילים תתל"ו שאלו רז"ל אימתי נבראו המלאכים. ר' יוחנן אמר בשני נבראו הדה הוא דכ' המקרה במים עליותיו וכו' וכתיב בתריה עושה מלאכיו רוחות. ר' חנינה אמר בחמשי נבראו המלאכים הדה הא דכתיב ועוף יעופף על הארץ וכתיב ובשתים יעופף ר' לוליאני בשם ר' יצחק אמר בין על דעתיה דר' יוחנן בין על דעתיה דר' חנינה, הכל מודים שלא נברא ביום א' כלום. שלא יאמרו מיכאל היה מותח בדרומה של רקיע וגבריאל בצפונו והקב"ה ממדד באמצע אלא אנכי ה' עושה כל נוטה שמים לבדי רוקע הארץ מיאתי. מי אתי כתיב, מי היה שותף עמי בבריאתו של עולם ע"כ. וכן כ' רש"י ז"ל בפ' יו"ט של ר"ה שלא נבראו המלאכים עד יום ב' וכר' יוחנן. וכ"כ הסמ"ג בהקדמת עשין וז"ל כשברא הקב"ה את העולם בראו ביום הראשון והמלאכים נבראו ביום השני שנאמר המקרה במים עליותיו וכתיב בתריה עושה מלאכיו רוחות ע"כ ואתיא כר' יוחנן. ובילקוט תילים ויהי ערב ויהי בקר יום אחד יום יחידו של עולם שלא היה בעולם אלא הוא, הוי אומר אתה הוא ה' האלהים לבדך ע"כ. ד"א כי גדול אתה ועושה נפלאות בנוהג שבעולם מלך בשר ודם מתקלס במדינה וגדולי המדינה מתקלסין עמו לפי שנושאין עמו במשאו. אבל הקב"ה אינו כן אלא הוא לבדו ברא את העולם הוא לבדו מתקלס בעולם, הוא לבדו מתהדר בעולמו, אמר ר' תנחומא כי גדול אתה ועושה נפלאות למה? כי אתה אלהים לבדך אתה לבדך בראת את העולם ע"כ.

| id | word | punctuation after |
|---|---|---|
| 4:27.1 | ונעלם |  |
| 4:27.2 | מעיני |  |
| 4:27.3 | הרח"ו |  |
| 4:27.4 | מה |  |
| 4:27.5 | שכתב |  |
| 4:27.6 | הרמב"ם |  |
| 4:27.7 | ז"ל |  |
| 4:27.8 | במורה |  |
| 4:27.9 | סוף |  |
| 4:27.10 | פרק |  |
| 4:27.11 | ט"ו |  |
| 4:27.12 | משני |  |
| 4:27.13 | וז"ל |  |
| 4:27.14 | כי |  |
| 4:27.15 | אלו |  |
| 4:27.16 | תאמר |  |
| 4:27.17 | על |  |
| 4:27.18 | דרך |  |
| 4:27.19 | משל |  |
| 4:27.20 | כי |  |
| 4:27.21 | הבורא |  |
| 4:27.22 | ברא |  |
| 4:27.23 | עולמות |  |
| 4:27.24 | רבות |  |
| 4:27.25 | קודם |  |
| 4:27.26 | זה |  |
| 4:27.27 | העולם |  |
| 4:27.28 | כפי |  |
| 4:27.29 | מנין |  |
| 4:27.30 | מלוא |  |
| 4:27.31 | רוחב |  |
| 4:27.32 | גלגל |  |
| 4:27.33 | העליון |  |
| 4:27.34 | חרדל | , |
| 4:27.35 | וכל |  |
| 4:27.36 | עולם |  |
| 4:27.37 | ועולם |  |
| 4:27.38 | מהן |  |
| 4:27.39 | עמד |  |
| 4:27.40 | שנים |  |
| 4:27.41 | רבות |  |
| 4:27.42 | בלי |  |
| 4:27.43 | מספר |  |
| 4:27.44 | כמלוא |  |
| 4:27.45 | רוחב |  |
| 4:27.46 | גלגל |  |
| 4:27.47 | העליון |  |
| 4:27.48 | חרדל | ), |
| 4:27.49 | היה |  |
| 4:27.50 | זה |  |
| 4:27.51 | כשנעריכהו |  |
| 4:27.52 | אל |  |
| 4:27.53 | מציאותו |  |
| 4:27.54 | יתברך |  |
| 4:27.55 | שאין |  |
| 4:27.56 | לו |  |
| 4:27.57 | תכלית |  |
| 4:27.58 | כמו |  |
| 4:27.59 | אלו | , |
| 4:27.60 | אמרת |  |
| 4:27.61 | כי |  |
| 4:27.62 | האל |  |
| 4:27.63 | ברא |  |
| 4:27.64 | אתמול |  |
| 4:27.65 | העולם | , |
| 4:27.66 | כי |  |
| 4:27.67 | כשנקיים |  |
| 4:27.68 | התחלת |  |
| 4:27.69 | מציאות |  |
| 4:27.70 | אחר |  |
| 4:27.71 | העדר |  |
| 4:27.72 | ג"א |  |
| 4:27.73 | דאפיסה | ) |
| 4:27.74 | אין |  |
| 4:27.75 | הפרש |  |
| 4:27.76 | בין |  |
| 4:27.77 | שתשים |  |
| 4:27.78 | זה |  |
| 4:27.79 | קודם |  |
| 4:27.80 | מאות |  |
| 4:27.81 | אלפים |  |
| 4:27.82 | שנים |  |
| 4:27.83 | או |  |
| 4:27.84 | מזמן |  |
| 4:27.85 | קרוב |  |
| 4:27.86 | מאוד |  |
| 4:27.87 | ע"כ | . |
| 4:27.88 | ומזה |  |
| 4:27.89 | תבין |  |
| 4:27.90 | חסרון |  |
| 4:27.91 | חכמה |  |
| 4:27.92 | מהרח"ו |  |
| 4:27.93 | שחשב |  |
| 4:27.94 | להחכם |  |
| 4:27.95 | מבלי |  |
| 4:27.96 | חכמה | . |
| 4:27.97 | ובספר |  |
| 4:27.98 | עוז |  |
| 4:27.99 | לאלהים |  |
| 4:27.100 | בבית |  |
| 4:27.101 | קדש |  |
| 4:27.102 | הקדשים |  |
| 4:27.103 | פ | ' |
| 4:27.104 | כ"א |  |
| 4:27.105 | ד | ' |
| 4:27.106 | ס"ז |  |
| 4:27.107 | ע"ד |  |
| 4:27.108 | העיז |  |
| 4:27.109 | פניו |  |
| 4:27.110 | להשיב |  |
| 4:27.111 | וז"ל |  |
| 4:27.112 | ואם |  |
| 4:27.113 | לא |  |
| 4:27.114 | תתרצה |  |
| 4:27.115 | בתשובת |  |
| 4:27.116 | הבחירה |  |
| 4:27.117 | נשיב |  |
| 4:27.118 | לך |  |
| 4:27.119 | שמעת |  |
| 4:27.120 | שנתהוה |  |
| 4:27.121 | מלכא |  |
| 4:27.122 | קדישא |  |
| 4:27.123 | ברא |  |
| 4:27.124 | העולמות | , |
| 4:27.125 | ובזה |  |
| 4:27.126 | לא |  |
| 4:27.127 | תוכל |  |
| 4:27.128 | להקשות |  |
| 4:27.129 | למה |  |
| 4:27.130 | בראם |  |
| 4:27.131 | עכשו |  |
| 4:27.132 | ולא |  |
| 4:27.133 | קודם | , |
| 4:27.134 | כי |  |
| 4:27.135 | מעת |  |
| 4:27.136 | שנתהוה |  |
| 4:27.137 | בראם |  |
| 4:27.138 | עכ"ל | . |
| 4:27.139 | ונמצא |  |
| 4:27.140 | לפי |  |
| 4:27.141 | דבריו |  |
| 4:27.142 | שאין |  |
| 4:27.143 | הבורא |  |
| 4:27.144 | יתברך |  |
| 4:27.145 | גדל |  |
| 4:27.146 | מאדם |  |
| 4:27.147 | הראשון |  |
| 4:27.148 | אלא |  |
| 4:27.149 | בחמשה |  |
| 4:27.150 | ימים |  |
| 4:27.151 | בלבד | . |
| 4:27.152 | ולדבריו |  |
| 4:27.153 | ג"כ |  |
| 4:27.154 | איך |  |
| 4:27.155 | יצוייר |  |
| 4:27.156 | מה |  |
| 4:27.157 | שדרשו |  |
| 4:27.158 | ז"ל |  |
| 4:27.159 | על |  |
| 4:27.160 | פסוק |  |
| 4:27.161 | ואהיה |  |
| 4:27.162 | אצלו |  |
| 4:27.163 | אמון | , |
| 4:27.164 | שהתורה |  |
| 4:27.165 | קדמה |  |
| 4:27.166 | לבריאת |  |
| 4:27.167 | העולם |  |
| 4:27.168 | תתעק"ד |  |
| 4:27.169 | דורות | , |
| 4:27.170 | כי |  |
| 4:27.171 | לדבריו |  |
| 4:27.172 | היתה |  |
| 4:27.173 | התורה |  |
| 4:27.174 | קודם |  |
| 4:27.175 | התהוות |  |
| 4:27.176 | הבורא | , |
| 4:27.177 | רחמנא |  |
| 4:27.178 | לצלן |  |
| 4:27.179 | ממחשבות |  |
| 4:27.180 | זרות |  |
| 4:27.181 | כאלה | . |
| 4:27.182 | וז"ל |  |
| 4:27.183 | מהר'י |  |
| 4:27.184 | צהרי |  |
| 4:27.185 | בשם |  |
| 4:27.186 | ר | ' |
| 4:27.187 | סעדיה |  |
| 4:27.188 | גאון |  |
| 4:27.189 | פרשה |  |
| 4:27.190 | בראשית |  |
| 4:27.191 | הענין |  |
| 4:27.192 | השביעי |  |
| 4:27.193 | בא |  |
| 4:27.194 | להודיענו | , |
| 4:27.195 | שאעפ"י |  |
| 4:27.196 | שתראו |  |
| 4:27.197 | השמים |  |
| 4:27.198 | גדולי |  |
| 4:27.199 | הערך |  |
| 4:27.200 | במדתם |  |
| 4:27.201 | כפי |  |
| 4:27.202 | מה |  |
| 4:27.203 | שנתבאר |  |
| 4:27.204 | במופת |  |
| 4:27.205 | לבעלי |  |
| 4:27.206 | חכמת |  |
| 4:27.207 | התשבורת |  |
| 4:27.208 | בשמים |  |
| 4:27.209 | ובארץ | , |
| 4:27.210 | אל |  |
| 4:27.211 | תחשבו |  |
| 4:27.212 | שהוצרכו |  |
| 4:27.213 | זמן |  |
| 4:27.214 | רב |  |
| 4:27.215 | לבריאתם | , |
| 4:27.216 | אלא |  |
| 4:27.217 | קורא |  |
| 4:27.218 | אני |  |
| 4:27.219 | להם |  |
| 4:27.220 | יעמדו |  |
| 4:27.221 | יחדו | , |
| 4:27.222 | ורוצה |  |
| 4:27.223 | לומר |  |
| 4:27.224 | בראשית |  |
| 4:27.225 | בזמן |  |
| 4:27.226 | קטן |  |
| 4:27.227 | שבזמנים |  |
| 4:27.228 | בלי |  |
| 4:27.229 | עמל |  |
| 4:27.230 | ובלי |  |
| 4:27.231 | טורח |  |
| 4:27.232 | ובלי |  |
| 4:27.233 | יגיעה | , |
| 4:27.234 | וכן |  |
| 4:27.235 | הוא |  |
| 4:27.236 | אומר |  |
| 4:27.237 | לא |  |
| 4:27.238 | ייעף |  |
| 4:27.239 | ולא |  |
| 4:27.240 | ייגע |  |
| 4:27.241 | אין |  |
| 4:27.242 | חקר |  |
| 4:27.243 | לתבונתו |  |
| 4:27.244 | עכ"ל | . |
| 4:27.245 | ובבראשית |  |
| 4:27.246 | רבה |  |
| 4:27.247 | פ | ' |
| 4:27.248 | א | ' |
| 4:27.249 | וג | ' |
| 4:27.250 | ובילקוט |  |
| 4:27.251 | בראשית |  |
| 4:27.252 | אמז |  |
| 4:27.253 | ה | ' |
| 4:27.254 | וישעיה |  |
| 4:27.255 | שכ"ב | , |
| 4:27.256 | ותילים |  |
| 4:27.257 | תתל"ו |  |
| 4:27.258 | שאלו |  |
| 4:27.259 | רז"ל |  |
| 4:27.260 | אימתי |  |
| 4:27.261 | נבראו |  |
| 4:27.262 | המלאכים | . |
| 4:27.263 | ר | ' |
| 4:27.264 | יוחנן |  |
| 4:27.265 | אמר |  |
| 4:27.266 | בשני |  |
| 4:27.267 | נבראו |  |
| 4:27.268 | הדה |  |
| 4:27.269 | הוא |  |
| 4:27.270 | דכ | ' |
| 4:27.271 | המקרה |  |
| 4:27.272 | במים |  |
| 4:27.273 | עליותיו |  |
| 4:27.274 | וכו | ' |
| 4:27.275 | וכתיב |  |
| 4:27.276 | בתריה |  |
| 4:27.277 | עושה |  |
| 4:27.278 | מלאכיו |  |
| 4:27.279 | רוחות | . |
| 4:27.280 | ר | ' |
| 4:27.281 | חנינה |  |
| 4:27.282 | אמר |  |
| 4:27.283 | בחמשי |  |
| 4:27.284 | נבראו |  |
| 4:27.285 | המלאכים |  |
| 4:27.286 | הדה |  |
| 4:27.287 | הא |  |
| 4:27.288 | דכתיב |  |
| 4:27.289 | ועוף |  |
| 4:27.290 | יעופף |  |
| 4:27.291 | על |  |
| 4:27.292 | הארץ |  |
| 4:27.293 | וכתיב |  |
| 4:27.294 | ובשתים |  |
| 4:27.295 | יעופף |  |
| 4:27.296 | ר | ' |
| 4:27.297 | לוליאני |  |
| 4:27.298 | בשם |  |
| 4:27.299 | ר | ' |
| 4:27.300 | יצחק |  |
| 4:27.301 | אמר |  |
| 4:27.302 | בין |  |
| 4:27.303 | על |  |
| 4:27.304 | דעתיה |  |
| 4:27.305 | דר | ' |
| 4:27.306 | יוחנן |  |
| 4:27.307 | בין |  |
| 4:27.308 | על |  |
| 4:27.309 | דעתיה |  |
| 4:27.310 | דר | ' |
| 4:27.311 | חנינה | , |
| 4:27.312 | הכל |  |
| 4:27.313 | מודים |  |
| 4:27.314 | שלא |  |
| 4:27.315 | נברא |  |
| 4:27.316 | ביום |  |
| 4:27.317 | א | ' |
| 4:27.318 | כלום | . |
| 4:27.319 | שלא |  |
| 4:27.320 | יאמרו |  |
| 4:27.321 | מיכאל |  |
| 4:27.322 | היה |  |
| 4:27.323 | מותח |  |
| 4:27.324 | בדרומה |  |
| 4:27.325 | של |  |
| 4:27.326 | רקיע |  |
| 4:27.327 | וגבריאל |  |
| 4:27.328 | בצפונו |  |
| 4:27.329 | והקב"ה |  |
| 4:27.330 | ממדד |  |
| 4:27.331 | באמצע |  |
| 4:27.332 | אלא |  |
| 4:27.333 | אנכי |  |
| 4:27.334 | ה | ' |
| 4:27.335 | עושה |  |
| 4:27.336 | כל |  |
| 4:27.337 | נוטה |  |
| 4:27.338 | שמים |  |
| 4:27.339 | לבדי |  |
| 4:27.340 | רוקע |  |
| 4:27.341 | הארץ |  |
| 4:27.342 | מיאתי | . |
| 4:27.343 | מי |  |
| 4:27.344 | אתי |  |
| 4:27.345 | כתיב | , |
| 4:27.346 | מי |  |
| 4:27.347 | היה |  |
| 4:27.348 | שותף |  |
| 4:27.349 | עמי |  |
| 4:27.350 | בבריאתו |  |
| 4:27.351 | של |  |
| 4:27.352 | עולם |  |
| 4:27.353 | ע"כ | . |
| 4:27.354 | וכן |  |
| 4:27.355 | כ | ' |
| 4:27.356 | רש"י |  |
| 4:27.357 | ז"ל |  |
| 4:27.358 | בפ | ' |
| 4:27.359 | יו"ט |  |
| 4:27.360 | של |  |
| 4:27.361 | ר"ה |  |
| 4:27.362 | שלא |  |
| 4:27.363 | נבראו |  |
| 4:27.364 | המלאכים |  |
| 4:27.365 | עד |  |
| 4:27.366 | יום |  |
| 4:27.367 | ב | ' |
| 4:27.368 | וכר | ' |
| 4:27.369 | יוחנן | . |
| 4:27.370 | וכ"כ |  |
| 4:27.371 | הסמ"ג |  |
| 4:27.372 | בהקדמת |  |
| 4:27.373 | עשין |  |
| 4:27.374 | וז"ל |  |
| 4:27.375 | כשברא |  |
| 4:27.376 | הקב"ה |  |
| 4:27.377 | את |  |
| 4:27.378 | העולם |  |
| 4:27.379 | בראו |  |
| 4:27.380 | ביום |  |
| 4:27.381 | הראשון |  |
| 4:27.382 | והמלאכים |  |
| 4:27.383 | נבראו |  |
| 4:27.384 | ביום |  |
| 4:27.385 | השני |  |
| 4:27.386 | שנאמר |  |
| 4:27.387 | המקרה |  |
| 4:27.388 | במים |  |
| 4:27.389 | עליותיו |  |
| 4:27.390 | וכתיב |  |
| 4:27.391 | בתריה |  |
| 4:27.392 | עושה |  |
| 4:27.393 | מלאכיו |  |
| 4:27.394 | רוחות |  |
| 4:27.395 | ע"כ |  |
| 4:27.396 | ואתיא |  |
| 4:27.397 | כר | ' |
| 4:27.398 | יוחנן | . |
| 4:27.399 | ובילקוט |  |
| 4:27.400 | תילים |  |
| 4:27.401 | ויהי |  |
| 4:27.402 | ערב |  |
| 4:27.403 | ויהי |  |
| 4:27.404 | בקר |  |
| 4:27.405 | יום |  |
| 4:27.406 | אחד |  |
| 4:27.407 | יום |  |
| 4:27.408 | יחידו |  |
| 4:27.409 | של |  |
| 4:27.410 | עולם |  |
| 4:27.411 | שלא |  |
| 4:27.412 | היה |  |
| 4:27.413 | בעולם |  |
| 4:27.414 | אלא |  |
| 4:27.415 | הוא | , |
| 4:27.416 | הוי |  |
| 4:27.417 | אומר |  |
| 4:27.418 | אתה |  |
| 4:27.419 | הוא |  |
| 4:27.420 | ה | ' |
| 4:27.421 | האלהים |  |
| 4:27.422 | לבדך |  |
| 4:27.423 | ע"כ | . |
| 4:27.424 | ד"א |  |
| 4:27.425 | כי |  |
| 4:27.426 | גדול |  |
| 4:27.427 | אתה |  |
| 4:27.428 | ועושה |  |
| 4:27.429 | נפלאות |  |
| 4:27.430 | בנוהג |  |
| 4:27.431 | שבעולם |  |
| 4:27.432 | מלך |  |
| 4:27.433 | בשר |  |
| 4:27.434 | ודם |  |
| 4:27.435 | מתקלס |  |
| 4:27.436 | במדינה |  |
| 4:27.437 | וגדולי |  |
| 4:27.438 | המדינה |  |
| 4:27.439 | מתקלסין |  |
| 4:27.440 | עמו |  |
| 4:27.441 | לפי |  |
| 4:27.442 | שנושאין |  |
| 4:27.443 | עמו |  |
| 4:27.444 | במשאו | . |
| 4:27.445 | אבל |  |
| 4:27.446 | הקב"ה |  |
| 4:27.447 | אינו |  |
| 4:27.448 | כן |  |
| 4:27.449 | אלא |  |
| 4:27.450 | הוא |  |
| 4:27.451 | לבדו |  |
| 4:27.452 | ברא |  |
| 4:27.453 | את |  |
| 4:27.454 | העולם |  |
| 4:27.455 | הוא |  |
| 4:27.456 | לבדו |  |
| 4:27.457 | מתקלס |  |
| 4:27.458 | בעולם | , |
| 4:27.459 | הוא |  |
| 4:27.460 | לבדו |  |
| 4:27.461 | מתהדר |  |
| 4:27.462 | בעולמו | , |
| 4:27.463 | אמר |  |
| 4:27.464 | ר | ' |
| 4:27.465 | תנחומא |  |
| 4:27.466 | כי |  |
| 4:27.467 | גדול |  |
| 4:27.468 | אתה |  |
| 4:27.469 | ועושה |  |
| 4:27.470 | נפלאות |  |
| 4:27.471 | למה | ? |
| 4:27.472 | כי |  |
| 4:27.473 | אתה |  |
| 4:27.474 | אלהים |  |
| 4:27.475 | לבדך |  |
| 4:27.476 | אתה |  |
| 4:27.477 | לבדך |  |
| 4:27.478 | בראת |  |
| 4:27.479 | את |  |
| 4:27.480 | העולם |  |
| 4:27.481 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

