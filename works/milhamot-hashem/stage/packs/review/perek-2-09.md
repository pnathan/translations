# Review pack: milhamot-hashem perek-2-09 (Milchamot ha-Shem 2:17–2:18)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-09`.

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
- **עבודה זרה / ע"ז** → "idolatry", with the literal "strange worship" in a note
  once. As a tractate: "Avodah Zarah". עובדי ע"ז → "worshippers of idols".
  **עכו"ם** → "idolaters", with the note "worshippers of stars and
  constellations" (an acronym of the censorship era). Never "gentiles" or
  "Christians", even where the referent is a Christian king (9:10) or the
  Trinity (5:25); a note may say so. אליל → "idol"; האלילה (1:4) → "the
  she-idol"; אל זר → "a strange god"; אל אחר → "another god"; אלהים אחרים →
  "other gods".
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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- אלוה (singular): "God" as in Job/Deut 32:15; not "the first God".
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
- פ' + a mishnah/tractate title (פ' ד' מיתות): the chapter's title, "Four Death
  Penalties"; no numeral added.
- מנחה in הגשת מנחה where the author argues from Gen 43 / 1 Sam 10: "gift";
  "meal offering" only in a sacrificial context.
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
- והוא שלא: "provided that" (fixed).
- אלוה of a man who makes himself an object of worship (Haman): "a god".
- Ambiguous לעבד / לעבד כ-: choose by the verse or context; the other reading
  in `readings`.
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
- A word with no function in its context (סדר): "Seder" as printed, with a
  note; the other reading in `readings`.
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


## Milchamot ha-Shem 2:17

Hebrew (Wikisource): וגדולה מזו שעושים גם כן לזבוח לפני מלך או שר צאן ובקר והיא מעבודות שבפנים שהעושה לאחד ממיני ע"ז אף שאין עבודתה בכך חייב כדתנן בסנהדרין פ' ד' מיתות העובד ע"ז אחד העובד, ואחד המזבח, ואחד המקטר ואחד המנסך, ואחד המשתחוה וכו' ואמרינן בגמ' עלה מאי אחד העובד? אמר ר' ירמיה הכי קאמר אחד העובד כדרכה, ואחד המזבח ואחד המקטר ואחד המנסך ואחד המשתחוה, ואפי' שלא כדרכה. וכן פסקו כל הפוסקים דארבע עבודות חייב עליהן אפי' שלא כדרכה. ודבר זה נהוג עד היום אצל אחינו בני הכפרים והעירות שאם היה לאחד מהם אויב שהוא מבקש רעתו הוא מביא כבש ושוחטו לפני השר שלו להושיעו מאויבו. וקרא נמי כתיב: הקריבהו נא לפחתך וכו'. קדמונינו היה עליהם ג"כ החק הזה בבוא המלך אל המדינה לזבוח לפניו שור ונוטלים את דמי השור מקופה של צבור. ולא שמעו שום מפקפק בזה משום ע"ז, לפי שאין מכוין לעבדו בזה ולקבלו עליו באלוה, רק אות כבוד הוא עושה לו. ויעקב אבינו ע"ה אמר לבניו: קחו מזמרת הארץ בכליכם והורידו לאיש מנחה, ונאמ': ויביאו לו את המנחה אשר בידם הביתה וישתחו לו אפים וכו'. ובשאול בחיר ה' נאמר: ובני בליעל אמרו מה יושיענו זה ולא הביאו לו מנחה. ובנדרים שנינו: נדר בקרבן ואמר לא נדרתי אלא בקרבנות מלכים. אלמא שהשתחויה וזיבוח והגשת מנחה וכל ארבע עבודות פנים מותר לעשותן לאדם שאינו נעבד כהמן, שאם הוא נעבד כהמן אסור, שכן נמנע מרדכי להשתחות להמן מפני שעשה עצמו ע"ז.

| id | word | punctuation after |
|---|---|---|
| 2:17.1 | וגדולה |  |
| 2:17.2 | מזו |  |
| 2:17.3 | שעושים |  |
| 2:17.4 | גם |  |
| 2:17.5 | כן |  |
| 2:17.6 | לזבוח |  |
| 2:17.7 | לפני |  |
| 2:17.8 | מלך |  |
| 2:17.9 | או |  |
| 2:17.10 | שר |  |
| 2:17.11 | צאן |  |
| 2:17.12 | ובקר |  |
| 2:17.13 | והיא |  |
| 2:17.14 | מעבודות |  |
| 2:17.15 | שבפנים |  |
| 2:17.16 | שהעושה |  |
| 2:17.17 | לאחד |  |
| 2:17.18 | ממיני |  |
| 2:17.19 | ע"ז |  |
| 2:17.20 | אף |  |
| 2:17.21 | שאין |  |
| 2:17.22 | עבודתה |  |
| 2:17.23 | בכך |  |
| 2:17.24 | חייב |  |
| 2:17.25 | כדתנן |  |
| 2:17.26 | בסנהדרין |  |
| 2:17.27 | פ | ' |
| 2:17.28 | ד | ' |
| 2:17.29 | מיתות |  |
| 2:17.30 | העובד |  |
| 2:17.31 | ע"ז |  |
| 2:17.32 | אחד |  |
| 2:17.33 | העובד | , |
| 2:17.34 | ואחד |  |
| 2:17.35 | המזבח | , |
| 2:17.36 | ואחד |  |
| 2:17.37 | המקטר |  |
| 2:17.38 | ואחד |  |
| 2:17.39 | המנסך | , |
| 2:17.40 | ואחד |  |
| 2:17.41 | המשתחוה |  |
| 2:17.42 | וכו | ' |
| 2:17.43 | ואמרינן |  |
| 2:17.44 | בגמ | ' |
| 2:17.45 | עלה |  |
| 2:17.46 | מאי |  |
| 2:17.47 | אחד |  |
| 2:17.48 | העובד | ? |
| 2:17.49 | אמר |  |
| 2:17.50 | ר | ' |
| 2:17.51 | ירמיה |  |
| 2:17.52 | הכי |  |
| 2:17.53 | קאמר |  |
| 2:17.54 | אחד |  |
| 2:17.55 | העובד |  |
| 2:17.56 | כדרכה | , |
| 2:17.57 | ואחד |  |
| 2:17.58 | המזבח |  |
| 2:17.59 | ואחד |  |
| 2:17.60 | המקטר |  |
| 2:17.61 | ואחד |  |
| 2:17.62 | המנסך |  |
| 2:17.63 | ואחד |  |
| 2:17.64 | המשתחוה | , |
| 2:17.65 | ואפי | ' |
| 2:17.66 | שלא |  |
| 2:17.67 | כדרכה | . |
| 2:17.68 | וכן |  |
| 2:17.69 | פסקו |  |
| 2:17.70 | כל |  |
| 2:17.71 | הפוסקים |  |
| 2:17.72 | דארבע |  |
| 2:17.73 | עבודות |  |
| 2:17.74 | חייב |  |
| 2:17.75 | עליהן |  |
| 2:17.76 | אפי | ' |
| 2:17.77 | שלא |  |
| 2:17.78 | כדרכה | . |
| 2:17.79 | ודבר |  |
| 2:17.80 | זה |  |
| 2:17.81 | נהוג |  |
| 2:17.82 | עד |  |
| 2:17.83 | היום |  |
| 2:17.84 | אצל |  |
| 2:17.85 | אחינו |  |
| 2:17.86 | בני |  |
| 2:17.87 | הכפרים |  |
| 2:17.88 | והעירות |  |
| 2:17.89 | שאם |  |
| 2:17.90 | היה |  |
| 2:17.91 | לאחד |  |
| 2:17.92 | מהם |  |
| 2:17.93 | אויב |  |
| 2:17.94 | שהוא |  |
| 2:17.95 | מבקש |  |
| 2:17.96 | רעתו |  |
| 2:17.97 | הוא |  |
| 2:17.98 | מביא |  |
| 2:17.99 | כבש |  |
| 2:17.100 | ושוחטו |  |
| 2:17.101 | לפני |  |
| 2:17.102 | השר |  |
| 2:17.103 | שלו |  |
| 2:17.104 | להושיעו |  |
| 2:17.105 | מאויבו | . |
| 2:17.106 | וקרא |  |
| 2:17.107 | נמי |  |
| 2:17.108 | כתיב | : |
| 2:17.109 | הקריבהו |  |
| 2:17.110 | נא |  |
| 2:17.111 | לפחתך |  |
| 2:17.112 | וכו | '. |
| 2:17.113 | קדמונינו |  |
| 2:17.114 | היה |  |
| 2:17.115 | עליהם |  |
| 2:17.116 | ג"כ |  |
| 2:17.117 | החק |  |
| 2:17.118 | הזה |  |
| 2:17.119 | בבוא |  |
| 2:17.120 | המלך |  |
| 2:17.121 | אל |  |
| 2:17.122 | המדינה |  |
| 2:17.123 | לזבוח |  |
| 2:17.124 | לפניו |  |
| 2:17.125 | שור |  |
| 2:17.126 | ונוטלים |  |
| 2:17.127 | את |  |
| 2:17.128 | דמי |  |
| 2:17.129 | השור |  |
| 2:17.130 | מקופה |  |
| 2:17.131 | של |  |
| 2:17.132 | צבור | . |
| 2:17.133 | ולא |  |
| 2:17.134 | שמעו |  |
| 2:17.135 | שום |  |
| 2:17.136 | מפקפק |  |
| 2:17.137 | בזה |  |
| 2:17.138 | משום |  |
| 2:17.139 | ע"ז | , |
| 2:17.140 | לפי |  |
| 2:17.141 | שאין |  |
| 2:17.142 | מכוין |  |
| 2:17.143 | לעבדו |  |
| 2:17.144 | בזה |  |
| 2:17.145 | ולקבלו |  |
| 2:17.146 | עליו |  |
| 2:17.147 | באלוה | , |
| 2:17.148 | רק |  |
| 2:17.149 | אות |  |
| 2:17.150 | כבוד |  |
| 2:17.151 | הוא |  |
| 2:17.152 | עושה |  |
| 2:17.153 | לו | . |
| 2:17.154 | ויעקב |  |
| 2:17.155 | אבינו |  |
| 2:17.156 | ע"ה |  |
| 2:17.157 | אמר |  |
| 2:17.158 | לבניו | : |
| 2:17.159 | קחו |  |
| 2:17.160 | מזמרת |  |
| 2:17.161 | הארץ |  |
| 2:17.162 | בכליכם |  |
| 2:17.163 | והורידו |  |
| 2:17.164 | לאיש |  |
| 2:17.165 | מנחה | , |
| 2:17.166 | ונאמ | ': |
| 2:17.167 | ויביאו |  |
| 2:17.168 | לו |  |
| 2:17.169 | את |  |
| 2:17.170 | המנחה |  |
| 2:17.171 | אשר |  |
| 2:17.172 | בידם |  |
| 2:17.173 | הביתה |  |
| 2:17.174 | וישתחו |  |
| 2:17.175 | לו |  |
| 2:17.176 | אפים |  |
| 2:17.177 | וכו | '. |
| 2:17.178 | ובשאול |  |
| 2:17.179 | בחיר |  |
| 2:17.180 | ה | ' |
| 2:17.181 | נאמר | : |
| 2:17.182 | ובני |  |
| 2:17.183 | בליעל |  |
| 2:17.184 | אמרו |  |
| 2:17.185 | מה |  |
| 2:17.186 | יושיענו |  |
| 2:17.187 | זה |  |
| 2:17.188 | ולא |  |
| 2:17.189 | הביאו |  |
| 2:17.190 | לו |  |
| 2:17.191 | מנחה | . |
| 2:17.192 | ובנדרים |  |
| 2:17.193 | שנינו | : |
| 2:17.194 | נדר |  |
| 2:17.195 | בקרבן |  |
| 2:17.196 | ואמר |  |
| 2:17.197 | לא |  |
| 2:17.198 | נדרתי |  |
| 2:17.199 | אלא |  |
| 2:17.200 | בקרבנות |  |
| 2:17.201 | מלכים | . |
| 2:17.202 | אלמא |  |
| 2:17.203 | שהשתחויה |  |
| 2:17.204 | וזיבוח |  |
| 2:17.205 | והגשת |  |
| 2:17.206 | מנחה |  |
| 2:17.207 | וכל |  |
| 2:17.208 | ארבע |  |
| 2:17.209 | עבודות |  |
| 2:17.210 | פנים |  |
| 2:17.211 | מותר |  |
| 2:17.212 | לעשותן |  |
| 2:17.213 | לאדם |  |
| 2:17.214 | שאינו |  |
| 2:17.215 | נעבד |  |
| 2:17.216 | כהמן | , |
| 2:17.217 | שאם |  |
| 2:17.218 | הוא |  |
| 2:17.219 | נעבד |  |
| 2:17.220 | כהמן |  |
| 2:17.221 | אסור | , |
| 2:17.222 | שכן |  |
| 2:17.223 | נמנע |  |
| 2:17.224 | מרדכי |  |
| 2:17.225 | להשתחות |  |
| 2:17.226 | להמן |  |
| 2:17.227 | מפני |  |
| 2:17.228 | שעשה |  |
| 2:17.229 | עצמו |  |
| 2:17.230 | ע"ז | . |


### Translation under review (2:17)

English: And greater than this is what they also do: they sacrifice before a king or a minister sheep and cattle. It is one of the inner worships. Whoever does it for one of the kinds of idolatry, even though its worship is not in that way, is liable, as we learned in Sanhedrin, chapter four, on the four death penalties: “The worshipper of idolatry: one who worships, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, etc.” And we say in the Gemara: On this: “What is ‘one who worships’?” Rabbi Yirmiyah said: “This is what it means: one who worships in the usual way, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, even not in the usual way.” And so all the decisors ruled that for the four worships one is liable even if not in the usual way. And this practice is still observed to this day among our brethren, the people of the villages and the towns: that if one of them had an enemy who seeks his harm, he brings a lamb and slaughters it before his minister, to save him from his enemy. And Scripture also says: “Offer it now to your governor, etc.” Our forebears also had this custom: when the king came to the province, they sacrificed an ox before him, and they took the price of the ox from the communal chest. And no one was heard to doubt this on account of idolatry, since one does not intend to serve him by this or to accept him upon oneself as a god; it is only a sign of honor that one makes to him. Our father Jacob, peace be upon him, said to his sons: “Take of the choicest of the land's produce in your vessels, and bring down to the man a gift,” and it is said: “And they brought him the gift that was in their hands to the house, and bowed down to him, faces, etc.” And concerning Saul, the LORD's chosen one, it is said: “And sons of Belial said, ‘How can this one save us?’ and did not bring him a gift.” And in Nedarim we learned: “One who vows by an offering and says, ‘I vowed only by the offerings of kings.’” So it is clear that bowing down, sacrifice, bringing a meal offering, and all four inner worships are permitted to do for a man who is not worshipped like Haman; but if he is worshipped like Haman it is forbidden, for Mordechai refused to bow down to Haman, because he made himself an object of idolatry.

Literal: and-great from-this that-they-do also so to-sacrifice before king or minister sheep and-cattle and-it of-worships of-the-inside that-the-one-who-does for-one of-kinds-of idolatry even though-not its-worship in-that liable as-we-learned in-Sanhedrin chapter four deaths the-worshipper-of idolatry one who-worships and-one the-altar and-one the-burner and-one the-pourer and-one the-bower etc and-we-say in-the-Gemara on-this what-is one who-worships said Rabbi Yirmiyah this-is what-it-means one who-worships in-its-way and-one the-altar and-one the-burner and-one the-pourer and-one the-bower even not in-its-way and-so all the-decisors ruled that-for the-four worships liable-on-them even not in-its-way and-this practised until the-day among our-brethren the-people-of the-villages and-the-towns that-if to-one of-them was enemy who-seeks his-harm he brings lamb and-slaughters-it before the-minister his to-save-him from-his-enemy and-Scripture also says offer-it now to-your-governor etc our-forebears had upon-them also the-custom this when came the-king to the-province to-sacrifice before-him ox and-they-took the-bloods-of-the-ox from-chest of-community and-not heard any doubter in-this on-account-of idolatry since not intends to-serve-him by-this and-to-accept-him upon-him as-a-god only sign of-honor it making to-him and-Jacob our-father peace-upon-him said to-his-sons take the-choicest of-the-land in-your-vessels and-bring-down to-the-man gift and-it-is-said and-they-brought to-him the-gift that in-their-hands to-the-house and-bowed to-him faces etc and-concerning-Saul chosen of-the-LORD it-is-said and-sons-of Belial said how will-save-us this and-not brought to-him gift and-in-Nedarim we-learned one vows by-offering and-says not I-vowed only by-offerings of-kings so-clear that-bowing and-sacrifice and-bringing meal-offering and-all the-four worships inner are-permitted to-do-them for-man who-is-not worshipped like-Haman but-if he is-worshipped like-Haman forbidden for refrained Mordechai to-bow-down to-Haman because he-made himself an-object-of idolatry

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:17.1 | and greater |  |  |  |
| 2:17.2 | than this |  |  |  |
| 2:17.3 | that they do |  |  |  |
| 2:17.4 | also |  |  |  |
| 2:17.5 | likewise |  |  |  |
| 2:17.6 | to sacrifice |  |  |  |
| 2:17.7 | before |  |  |  |
| 2:17.8 | a king |  |  |  |
| 2:17.9 | or |  |  |  |
| 2:17.10 | a minister |  |  |  |
| 2:17.11 | sheep |  |  |  |
| 2:17.12 | and cattle |  |  |  |
| 2:17.13 | and it is |  |  |  |
| 2:17.14 | one of the worships |  |  |  |
| 2:17.15 | of the inner kind |  |  | avodot penim, the 'inner' worships: the four acts the Mishnah names below; literally 'worships of the inside'. |
| 2:17.16 | that the one who does |  |  |  |
| 2:17.17 | for one |  |  |  |
| 2:17.18 | of the kinds of |  |  |  |
| 2:17.19 | idolatry |  |  | Literally 'strange worship' (avodah zarah); rendered 'idolatry' throughout. |
| 2:17.20 | even |  |  |  |
| 2:17.21 | though |  |  |  |
| 2:17.22 | its worship |  |  |  |
| 2:17.23 | in that way |  |  |  |
| 2:17.24 | is liable |  |  |  |
| 2:17.25 | as we learned |  |  |  |
| 2:17.26 | in Sanhedrin |  |  | Tractate Sanhedrin of the Mishnah; its chapter is given as the author writes it. |
| 2:17.27 | chapter |  |  |  |
| 2:17.28 | four |  |  |  |
| 2:17.29 | death penalties |  |  |  |
| 2:17.30 | the worshipper of |  |  |  |
| 2:17.31 | idolatry |  |  |  |
| 2:17.32 | one |  |  |  |
| 2:17.33 | who worships |  |  |  |
| 2:17.34 | and one |  |  |  |
| 2:17.35 | the altar |  |  | The Hebrew names the altar; the sense is the one who offers on it. |
| 2:17.36 | and one |  |  |  |
| 2:17.37 | who burns incense |  |  |  |
| 2:17.38 | and one |  |  |  |
| 2:17.39 | who pours libation |  |  |  |
| 2:17.40 | and one |  |  |  |
| 2:17.41 | who bows down |  |  |  |
| 2:17.42 | etc. |  |  |  |
| 2:17.43 | and we say |  |  |  |
| 2:17.44 | in the Gemara |  |  | The Gemara is the Talmud's discussion of the Mishnah. |
| 2:17.45 | on this |  |  | The Gemara asks what the Mishnah's first 'one who worships' means. |
| 2:17.46 | what is |  |  |  |
| 2:17.47 | one |  |  |  |
| 2:17.48 | who worships |  |  |  |
| 2:17.49 | said |  |  |  |
| 2:17.50 | Rabbi |  |  |  |
| 2:17.51 | Yirmiyah |  |  |  |
| 2:17.52 | this is |  |  |  |
| 2:17.53 | what it means |  |  |  |
| 2:17.54 | one |  |  |  |
| 2:17.55 | who worships |  |  |  |
| 2:17.56 | in the usual way |  |  | Literally 'in its way': in the usual manner. |
| 2:17.57 | and one |  |  |  |
| 2:17.58 | who sacrifices at the altar |  |  |  |
| 2:17.59 | and one |  |  |  |
| 2:17.60 | who burns incense |  |  |  |
| 2:17.61 | and one |  |  |  |
| 2:17.62 | who pours libation |  |  |  |
| 2:17.63 | and one |  |  |  |
| 2:17.64 | who bows down |  |  |  |
| 2:17.65 | even |  |  |  |
| 2:17.66 | not |  |  |  |
| 2:17.67 | in the usual way |  |  |  |
| 2:17.68 | and so |  |  |  |
| 2:17.69 | ruled |  |  |  |
| 2:17.70 | all |  |  |  |
| 2:17.71 | the decisors |  |  |  |
| 2:17.72 | that for the four |  |  |  |
| 2:17.73 | worships |  |  |  |
| 2:17.74 | one is liable |  |  |  |
| 2:17.75 | for them |  |  |  |
| 2:17.76 | even |  |  |  |
| 2:17.77 | not |  |  |  |
| 2:17.78 | in the usual way |  |  |  |
| 2:17.79 | and this |  |  |  |
| 2:17.80 | this |  |  |  |
| 2:17.81 | is practised |  |  |  |
| 2:17.82 | to |  |  |  |
| 2:17.83 | this day |  |  |  |
| 2:17.84 | among |  |  |  |
| 2:17.85 | our brethren |  |  |  |
| 2:17.86 | the people of |  |  |  |
| 2:17.87 | the villages |  |  |  |
| 2:17.88 | and the towns |  |  |  |
| 2:17.89 | that if |  |  |  |
| 2:17.90 | had |  |  |  |
| 2:17.91 | one of |  |  |  |
| 2:17.92 | them |  |  |  |
| 2:17.93 | an enemy |  |  |  |
| 2:17.94 | who |  |  |  |
| 2:17.95 | seeks |  |  |  |
| 2:17.96 | his harm |  |  |  |
| 2:17.97 | he |  |  |  |
| 2:17.98 | brings |  |  |  |
| 2:17.99 | a lamb |  |  |  |
| 2:17.100 | and slaughters it |  |  |  |
| 2:17.101 | before |  |  |  |
| 2:17.102 | the minister |  |  |  |
| 2:17.103 | his |  |  |  |
| 2:17.104 | to save him |  |  |  |
| 2:17.105 | from his enemy |  |  |  |
| 2:17.106 | and Scripture |  |  |  |
| 2:17.107 | also |  |  |  |
| 2:17.108 | says |  |  |  |
| 2:17.109 | offer it |  |  |  |
| 2:17.110 | now |  |  |  |
| 2:17.111 | to your governor |  |  |  |
| 2:17.112 | etc. |  |  |  |
| 2:17.113 | our forebears |  |  |  |
| 2:17.114 | had |  |  |  |
| 2:17.115 | upon them |  |  |  |
| 2:17.116 | also |  |  |  |
| 2:17.117 | the custom |  |  |  |
| 2:17.118 | this |  |  |  |
| 2:17.119 | when |  |  |  |
| 2:17.120 | the king |  |  |  |
| 2:17.121 | to |  |  |  |
| 2:17.122 | the province |  |  | medinah: a province or district. |
| 2:17.123 | to sacrifice |  |  |  |
| 2:17.124 | before him |  |  |  |
| 2:17.125 | an ox |  |  |  |
| 2:17.126 | and they took |  |  |  |
| 2:17.127 | - |  |  |  |
| 2:17.128 | the price |  |  | dmei: literally 'bloods'; here the price of the ox, which is the sense the sentence needs. The sense 'blood' is not taken. |
| 2:17.129 | of the ox |  |  |  |
| 2:17.130 | from the chest |  |  |  |
| 2:17.131 | of |  |  |  |
| 2:17.132 | the community |  |  |  |
| 2:17.133 | and no one |  |  |  |
| 2:17.134 | was heard |  |  |  |
| 2:17.135 | any |  |  |  |
| 2:17.136 | doubter |  |  |  |
| 2:17.137 | about this |  |  |  |
| 2:17.138 | on account of |  |  |  |
| 2:17.139 | idolatry |  |  |  |
| 2:17.140 | since |  |  |  |
| 2:17.141 | one does not |  |  |  |
| 2:17.142 | intend |  |  |  |
| 2:17.143 | to serve him |  |  |  |
| 2:17.144 | by this |  |  |  |
| 2:17.145 | or to accept |  |  |  |
| 2:17.146 | upon himself |  |  |  |
| 2:17.147 | as a god |  |  | be-eloah, 'as a god': the author's own term, applied here to the king in his argument. |
| 2:17.148 | only |  |  |  |
| 2:17.149 | a sign |  |  |  |
| 2:17.150 | of honor |  |  |  |
| 2:17.151 | it |  |  |  |
| 2:17.152 | one makes |  |  |  |
| 2:17.153 | to him |  |  |  |
| 2:17.154 | and Jacob |  |  |  |
| 2:17.155 | our father |  |  |  |
| 2:17.156 | peace be upon him |  |  |  |
| 2:17.157 | said |  |  |  |
| 2:17.158 | to his sons |  |  |  |
| 2:17.159 | take |  |  |  |
| 2:17.160 | the choicest |  |  | zimrat ha-aretz (Gen 43:11): the land's best produce. |
| 2:17.161 | of the land |  |  |  |
| 2:17.162 | in your vessels |  |  |  |
| 2:17.163 | and bring down |  |  |  |
| 2:17.164 | to the man |  |  |  |
| 2:17.165 | a gift |  |  |  |
| 2:17.166 | and it is said |  |  |  |
| 2:17.167 | and they brought |  |  |  |
| 2:17.168 | to him |  |  |  |
| 2:17.169 | - |  |  |  |
| 2:17.170 | the gift |  |  |  |
| 2:17.171 | that |  |  |  |
| 2:17.172 | in their hands |  |  |  |
| 2:17.173 | to the house |  |  |  |
| 2:17.174 | and bowed down |  |  |  |
| 2:17.175 | to him |  |  |  |
| 2:17.176 | faces |  |  | apayim, 'faces': the author's wording; see apparatus. |
| 2:17.177 | etc. |  |  |  |
| 2:17.178 | and concerning |  |  |  |
| 2:17.179 | the chosen |  |  |  |
| 2:17.180 | of the LORD |  |  |  |
| 2:17.181 | it is said |  |  |  |
| 2:17.182 | and sons of |  |  |  |
| 2:17.183 | Belial |  |  | Belial: worthless men, as in 1 Sam 10:27. |
| 2:17.184 | said |  |  |  |
| 2:17.185 | how |  |  |  |
| 2:17.186 | can save us |  |  |  |
| 2:17.187 | this one |  |  |  |
| 2:17.188 | and not |  |  |  |
| 2:17.189 | bring |  |  |  |
| 2:17.190 | to him |  |  |  |
| 2:17.191 | a gift |  |  |  |
| 2:17.192 | and in Nedarim |  |  |  |
| 2:17.193 | we learned |  |  |  |
| 2:17.194 | one vows |  |  |  |
| 2:17.195 | by an offering |  |  |  |
| 2:17.196 | and says |  |  |  |
| 2:17.197 | not |  |  |  |
| 2:17.198 | I vowed |  |  |  |
| 2:17.199 | only |  |  |  |
| 2:17.200 | by offerings |  |  |  |
| 2:17.201 | of kings |  |  |  |
| 2:17.202 | so it is clear |  |  |  |
| 2:17.203 | that bowing down |  |  |  |
| 2:17.204 | and sacrifice |  |  |  |
| 2:17.205 | and bringing |  |  |  |
| 2:17.206 | a meal offering |  |  | minchah here is a meal offering, the grain offering; the same word is 'gift' above. |
| 2:17.207 | and all |  |  |  |
| 2:17.208 | the four |  |  |  |
| 2:17.209 | worships |  |  |  |
| 2:17.210 | inner |  |  |  |
| 2:17.211 | are permitted |  |  |  |
| 2:17.212 | to do them |  |  |  |
| 2:17.213 | for a man |  |  |  |
| 2:17.214 | who is not |  |  |  |
| 2:17.215 | worshipped |  |  |  |
| 2:17.216 | like Haman |  |  | Haman, the villain of Esther. |
| 2:17.217 | but if |  |  |  |
| 2:17.218 | he |  |  |  |
| 2:17.219 | is worshipped |  |  |  |
| 2:17.220 | like Haman |  |  |  |
| 2:17.221 | forbidden |  |  |  |
| 2:17.222 | for |  |  |  |
| 2:17.223 | refrained |  |  |  |
| 2:17.224 | Mordechai |  |  | Mordechai, who refused to bow down to Haman (Esther 3:2). |
| 2:17.225 | to bow down |  |  |  |
| 2:17.226 | to Haman |  |  |  |
| 2:17.227 | because |  |  |  |
| 2:17.228 | he made |  |  |  |
| 2:17.229 | himself |  |  |  |
| 2:17.230 | an object of idolatry |  |  | Literally 'made himself strange worship'. |

Phrases:
- 2:17.1 2:17.2 2:17.3 2:17.4 2:17.5 2:17.6 2:17.7 2:17.8 2:17.9 2:17.10 2:17.11 2:17.12 | And greater than this is what they also do: they sacrifice before a king or a minister sheep and cattle. | and-great from-this that-they-do also so to-sacrifice before king or minister sheep and-cattle
- 2:17.13 2:17.14 2:17.15 | It is one of the inner worships. | and-it of-worships of-the-inside
- 2:17.16 2:17.17 2:17.18 2:17.19 2:17.20 2:17.21 2:17.22 2:17.23 2:17.24 | Whoever does it for one of the kinds of idolatry, even though its worship is not in that way, is liable, | that-the-one-who-does for-one of-kinds-of idolatry even though-not its-worship in-that liable
- 2:17.25 2:17.26 2:17.27 2:17.28 2:17.29 | as we learned in Sanhedrin, chapter four, on the four death penalties: | as-we-learned in-Sanhedrin chapter four deaths
- 2:17.30 2:17.31 2:17.32 2:17.33 2:17.34 2:17.35 2:17.36 2:17.37 2:17.38 2:17.39 2:17.40 2:17.41 2:17.42 | “The worshipper of idolatry: one who worships, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, etc.” | the-worshipper-of idolatry one who-worships and-one the-altar and-one the-burner and-one the-pourer and-one the-bower etc
- 2:17.43 2:17.44 | And we say in the Gemara: | and-we-say in-the-Gemara
- 2:17.45 2:17.46 2:17.47 2:17.48 | On this: “What is ‘one who worships’?” | on-this what-is one who-worships
- 2:17.49 2:17.50 2:17.51 2:17.52 2:17.53 | Rabbi Yirmiyah said: “This is what it means: | said Rabbi Yirmiyah this-is what-it-means
- 2:17.54 2:17.55 2:17.56 2:17.57 2:17.58 2:17.59 2:17.60 2:17.61 2:17.62 2:17.63 2:17.64 2:17.65 2:17.66 2:17.67 | one who worships in the usual way, and one who sacrifices at the altar, and one who burns incense, and one who pours libation, and one who bows down, even not in the usual way.” | one who-worships in-its-way and-one the-altar and-one the-burner and-one the-pourer and-one the-bower even not in-its-way
- 2:17.68 2:17.69 2:17.70 2:17.71 2:17.72 2:17.73 2:17.74 2:17.75 2:17.76 2:17.77 2:17.78 | And so all the decisors ruled that for the four worships one is liable even if not in the usual way. | and-so all the-decisors ruled that-for the-four worships liable-on-them even not in-its-way
- 2:17.79 2:17.80 2:17.81 2:17.82 2:17.83 2:17.84 2:17.85 2:17.86 2:17.87 2:17.88 | And this practice is still observed to this day among our brethren, the people of the villages and the towns: | and-this practised until the-day among our-brethren the-people-of the-villages and-the-towns
- 2:17.89 2:17.90 2:17.91 2:17.92 2:17.93 2:17.94 2:17.95 2:17.96 2:17.97 2:17.98 2:17.99 2:17.100 2:17.101 2:17.102 2:17.103 2:17.104 2:17.105 | that if one of them had an enemy who seeks his harm, he brings a lamb and slaughters it before his minister, to save him from his enemy. | that-if to-one of-them was enemy who-seeks his-harm he brings lamb and-slaughters-it before the-minister his to-save-him from-his-enemy
- 2:17.106 2:17.107 2:17.108 2:17.109 2:17.110 2:17.111 2:17.112 | And Scripture also says: “Offer it now to your governor, etc.” | and-Scripture also says offer-it now to-your-governor etc
- 2:17.113 2:17.114 2:17.115 2:17.116 2:17.117 2:17.118 2:17.119 2:17.120 2:17.121 2:17.122 2:17.123 2:17.124 2:17.125 | Our forebears also had this custom: when the king came to the province, they sacrificed an ox before him, | our-forebears had upon-them also the-custom this when came the-king to the-province to-sacrifice before-him ox
- 2:17.126 2:17.127 2:17.128 2:17.129 2:17.130 2:17.131 2:17.132 | and they took the price of the ox from the communal chest. | and-they-took the-bloods-of-the-ox from-chest of-community
- 2:17.133 2:17.134 2:17.135 2:17.136 2:17.137 2:17.138 2:17.139 | And no one was heard to doubt this on account of idolatry, | and-not heard any doubter in-this on-account-of idolatry
- 2:17.140 2:17.141 2:17.142 2:17.143 2:17.144 2:17.145 2:17.146 2:17.147 | since one does not intend to serve him by this or to accept him upon oneself as a god; | since not intends to-serve-him by-this and-to-accept-him upon-him as-a-god
- 2:17.148 2:17.149 2:17.150 2:17.151 2:17.152 2:17.153 | it is only a sign of honor that one makes to him. | only sign of-honor it making to-him
- 2:17.154 2:17.155 2:17.156 2:17.157 2:17.158 2:17.159 2:17.160 2:17.161 2:17.162 2:17.163 2:17.164 2:17.165 2:17.166 | Our father Jacob, peace be upon him, said to his sons: “Take of the choicest of the land's produce in your vessels, and bring down to the man a gift,” and it is said: | and-Jacob our-father peace-upon-him said to-his-sons take the-choicest of-the-land in-your-vessels and-bring-down to-the-man gift and-it-is-said
- 2:17.167 2:17.168 2:17.169 2:17.170 2:17.171 2:17.172 2:17.173 2:17.174 2:17.175 2:17.176 2:17.177 | “And they brought him the gift that was in their hands to the house, and bowed down to him, faces, etc.” | and-they-brought to-him the-gift that in-their-hands to-the-house and-bowed to-him faces etc
- 2:17.178 2:17.179 2:17.180 2:17.181 | And concerning Saul, the LORD's chosen one, it is said: | and-concerning-Saul chosen of-the-LORD it-is-said
- 2:17.182 2:17.183 2:17.184 2:17.185 2:17.186 2:17.187 2:17.188 2:17.189 2:17.190 2:17.191 | “And sons of Belial said, ‘How can this one save us?’ and did not bring him a gift.” | and-sons-of Belial said how will-save-us this and-not brought to-him gift
- 2:17.192 2:17.193 2:17.194 2:17.195 2:17.196 2:17.197 2:17.198 2:17.199 2:17.200 2:17.201 | And in Nedarim we learned: “One who vows by an offering and says, ‘I vowed only by the offerings of kings.’” | and-in-Nedarim we-learned one vows by-offering and-says not I-vowed only by-offerings of-kings
- 2:17.202 2:17.203 2:17.204 2:17.205 2:17.206 2:17.207 2:17.208 2:17.209 2:17.210 2:17.211 2:17.212 | So it is clear that bowing down, sacrifice, bringing a meal offering, and all four inner worships are permitted to do | so-clear that-bowing and-sacrifice and-bringing meal-offering and-all the-four worships inner are-permitted to-do-them
- 2:17.213 2:17.214 2:17.215 2:17.216 | for a man who is not worshipped like Haman; | for-man who-is-not worshipped like-Haman
- 2:17.217 2:17.218 2:17.219 2:17.220 2:17.221 | but if he is worshipped like Haman it is forbidden, | but-if he is-worshipped like-Haman forbidden
- 2:17.222 2:17.223 2:17.224 2:17.225 2:17.226 2:17.227 2:17.228 2:17.229 2:17.230 | for Mordechai refused to bow down to Haman, because he made himself an object of idolatry. | for refrained Mordechai to-bow-down to-Haman because he-made himself an-object-of idolatry

Grammar notes:
- שעושים לזבוח: the relative clause שעושים ('that they do') governs the infinitive לזבוח ('to sacrifice'); the object is 'sheep and cattle' and the place is 'before a king or a minister'.
- והיא מעבודות שבפנים: והיא refers back to the act in the sentence before; עבודות פנים is the Mishnah's term for the inner worships.
- עלה מאי אחד העובד: the Gemara asks what the Mishnah's first 'one who worships' means; Rabbi Yirmiyah answers with הכי קאמר ('this is what it means').
- The Mishnah's list and Rabbi Yirmiyah's list each close with 'even not in the usual way' (ואפי' שלא כדרכה): the English repeats the phrase at both points.
- נעבד (passive, 'is worshipped'), נמנע ('refrained') and the adjective אסור ('forbidden') carry the last sentences; the English keeps the passive.
- שאם היה לאחד מהם אויב ... הוא מביא: a past verb in the conditional clause and a present verb in the main clause, with a habitual sense: 'if one of them had an enemy, he brings'.

Translator's note: The Mishnah's list and Rabbi Yirmiyah's gloss are given in full, as the author quotes them. 'Idolatry' and 'inner worships' follow the fixed terms. The chapter reference is the author's and is kept.

Commentary rulings:
- apparatus: Mal 1:8: 'הקריבהו נא לפחתך' (Offer it now to your governor), quoted by the author with וכו'.
- apparatus: Gen 43:11: 'קחו מזמרת הארץ בכליכם והורידו לאיש מנחה' (Take of the choicest of the land's produce in your vessels and bring down to the man a gift).
- apparatus: Gen 43:26: 'ויביאו לו את המנחה אשר בידם הביתה'; the author continues 'וישתחו לו אפים וכו'', and the English follows his wording, not the verse's continuation.
- apparatus: 1 Sam 10:27: 'ובני בליעל אמרו מה יושיענו זה ולא הביאו לו מנחה'; the verse has an extra clause before 'ולא הביאו' that the author omits, and the English omits it too.
- apparatus: Esth 3:2: Mordechai's refusal to bow to Haman, given by the author in his own words.
- apparatus: The Mishnah of Sanhedrin is cited as 'chapter four' (פ' ד'); the reference is kept as written.
- apparatus: The quotation from Nedarim gives no reference, and none is added.

## Milchamot ha-Shem 2:18

Hebrew (Wikisource): ובס' תשועת ישראל לר' מנשה בן ישראל כ' וז"ל: נדרשתי להודיע מנהגי הכבוד והכניעות אשר יכבדו בם היהודים את ספר התורה, כי צרינו אשר אינם יודעים מה זאת מכעיסים אותנו לאמר כי עובדי אלילים אנחנו, וכי נשתחוה לארון עץ וליריעות עור אשר בתוכה, לכן אערוך תשובתי בסדר הניאות כיד ה' הטובה עלי והוא כל איש ישראל יראה את עצמו מחוייב לקום ולעמוד בהוצאת ספר תורה מארון הקודש עדי תונח על התיבה להראותה לכל העם ולקרות בה. ומצינו כי כן עשו בימי נחמיה ככתוב וכפתחו עמדו כל העם. וזאת יעשו מפני היראה והכבוד אשר יכבדו דברי אלהים חיים הכתובים בה, ומאותו טעם יכופו ראשם בנשאם אותה מהארון עד התיבה, ולא משום דבר אחר מן הטעמים אשר נזכיר.

| id | word | punctuation after |
|---|---|---|
| 2:18.1 | ובס | ' |
| 2:18.2 | תשועת |  |
| 2:18.3 | ישראל |  |
| 2:18.4 | לר | ' |
| 2:18.5 | מנשה |  |
| 2:18.6 | בן |  |
| 2:18.7 | ישראל |  |
| 2:18.8 | כ | ' |
| 2:18.9 | וז"ל | : |
| 2:18.10 | נדרשתי |  |
| 2:18.11 | להודיע |  |
| 2:18.12 | מנהגי |  |
| 2:18.13 | הכבוד |  |
| 2:18.14 | והכניעות |  |
| 2:18.15 | אשר |  |
| 2:18.16 | יכבדו |  |
| 2:18.17 | בם |  |
| 2:18.18 | היהודים |  |
| 2:18.19 | את |  |
| 2:18.20 | ספר |  |
| 2:18.21 | התורה | , |
| 2:18.22 | כי |  |
| 2:18.23 | צרינו |  |
| 2:18.24 | אשר |  |
| 2:18.25 | אינם |  |
| 2:18.26 | יודעים |  |
| 2:18.27 | מה |  |
| 2:18.28 | זאת |  |
| 2:18.29 | מכעיסים |  |
| 2:18.30 | אותנו |  |
| 2:18.31 | לאמר |  |
| 2:18.32 | כי |  |
| 2:18.33 | עובדי |  |
| 2:18.34 | אלילים |  |
| 2:18.35 | אנחנו | , |
| 2:18.36 | וכי |  |
| 2:18.37 | נשתחוה |  |
| 2:18.38 | לארון |  |
| 2:18.39 | עץ |  |
| 2:18.40 | וליריעות |  |
| 2:18.41 | עור |  |
| 2:18.42 | אשר |  |
| 2:18.43 | בתוכה | , |
| 2:18.44 | לכן |  |
| 2:18.45 | אערוך |  |
| 2:18.46 | תשובתי |  |
| 2:18.47 | בסדר |  |
| 2:18.48 | הניאות |  |
| 2:18.49 | כיד |  |
| 2:18.50 | ה | ' |
| 2:18.51 | הטובה |  |
| 2:18.52 | עלי |  |
| 2:18.53 | והוא |  |
| 2:18.54 | כל |  |
| 2:18.55 | איש |  |
| 2:18.56 | ישראל |  |
| 2:18.57 | יראה |  |
| 2:18.58 | את |  |
| 2:18.59 | עצמו |  |
| 2:18.60 | מחוייב |  |
| 2:18.61 | לקום |  |
| 2:18.62 | ולעמוד |  |
| 2:18.63 | בהוצאת |  |
| 2:18.64 | ספר |  |
| 2:18.65 | תורה |  |
| 2:18.66 | מארון |  |
| 2:18.67 | הקודש |  |
| 2:18.68 | עדי |  |
| 2:18.69 | תונח |  |
| 2:18.70 | על |  |
| 2:18.71 | התיבה |  |
| 2:18.72 | להראותה |  |
| 2:18.73 | לכל |  |
| 2:18.74 | העם |  |
| 2:18.75 | ולקרות |  |
| 2:18.76 | בה | . |
| 2:18.77 | ומצינו |  |
| 2:18.78 | כי |  |
| 2:18.79 | כן |  |
| 2:18.80 | עשו |  |
| 2:18.81 | בימי |  |
| 2:18.82 | נחמיה |  |
| 2:18.83 | ככתוב |  |
| 2:18.84 | וכפתחו |  |
| 2:18.85 | עמדו |  |
| 2:18.86 | כל |  |
| 2:18.87 | העם | . |
| 2:18.88 | וזאת |  |
| 2:18.89 | יעשו |  |
| 2:18.90 | מפני |  |
| 2:18.91 | היראה |  |
| 2:18.92 | והכבוד |  |
| 2:18.93 | אשר |  |
| 2:18.94 | יכבדו |  |
| 2:18.95 | דברי |  |
| 2:18.96 | אלהים |  |
| 2:18.97 | חיים |  |
| 2:18.98 | הכתובים |  |
| 2:18.99 | בה | , |
| 2:18.100 | ומאותו |  |
| 2:18.101 | טעם |  |
| 2:18.102 | יכופו |  |
| 2:18.103 | ראשם |  |
| 2:18.104 | בנשאם |  |
| 2:18.105 | אותה |  |
| 2:18.106 | מהארון |  |
| 2:18.107 | עד |  |
| 2:18.108 | התיבה | , |
| 2:18.109 | ולא |  |
| 2:18.110 | משום |  |
| 2:18.111 | דבר |  |
| 2:18.112 | אחר |  |
| 2:18.113 | מן |  |
| 2:18.114 | הטעמים |  |
| 2:18.115 | אשר |  |
| 2:18.116 | נזכיר | . |


### Translation under review (2:18)

English: And in the book Teshu'at Yisrael by Rabbi Menasseh ben Israel he wrote, and these are his words: “I have been required to make known the customs of honor and submission with which the Jews honor the Torah scroll, for our enemies, who do not know what this is, provoke us, saying that we are worshippers of idols, and that we bow down to a wooden ark and to the sheets of skin within it. Therefore I shall set out my reply in proper order, as the good hand of the LORD is upon me. namely, every man of Israel will see himself obliged to rise and stand at the removal of a Torah scroll from the holy ark until it is placed on the reading platform, to show it to all the people and to read from it. And we find that they did so in the days of Nehemiah, as it is written: ‘And when he opened, all the people stood.’ And this they do out of the fear and honor with which they honor the words of the living God written in it, and for that same reason they bow their heads when carrying it from the ark to the reading platform, and not because of any other of the reasons that we shall mention.”

Literal: and-in-the-book Teshu'at Yisrael by-Rabbi Menasseh son-of Israel wrote and-these-his-words I-have-been-required to-make-known customs-of honor and-submission with-which-they-honor with-them the-Jews object scroll-of the-Torah for our-enemies who not know what this provoke us saying that worshippers-of idols we and-that we-bow to-ark of-wood and-to-sheets of-skin which within-it therefore I-shall-arrange my-reply in-order proper as-hand-of the-LORD good upon-me and-it every man of-Israel will-see himself obliged to-rise and-stand at-removal-of scroll-of Torah from-ark the-holy until is-placed on the-reading-platform to-show-it to-all the-people and-to-read from-it and-we-find that-so they-did in-days-of Nehemiah as-it-is-written and-when-he-opened stood all the-people and-this they-do from the-fear and-the-honor which they-honor the-words-of God living written in-it and-for-that reason they-bow their-heads when-carrying-it from-the-ark until the-reading-platform and-not because-of other-thing of the-reasons which we-shall-mention

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:18.1 | and in the book |  |  |  |
| 2:18.2 | Teshu'at |  |  | Teshu'at Yisrael: literally 'the salvation of Israel', the title of the book. |
| 2:18.3 | Yisrael |  |  |  |
| 2:18.4 | Rabbi |  |  |  |
| 2:18.5 | Menasseh |  |  |  |
| 2:18.6 | ben |  |  |  |
| 2:18.7 | Israel |  |  |  |
| 2:18.8 | wrote |  |  |  |
| 2:18.9 | and these are his words: |  |  | vezal: the formula that opens the quotation, 'and these are his words'. |
| 2:18.10 | I have been required |  |  |  |
| 2:18.11 | to make known |  |  |  |
| 2:18.12 | the customs of |  |  |  |
| 2:18.13 | honor |  |  |  |
| 2:18.14 | and submission |  |  |  |
| 2:18.15 | with which |  |  |  |
| 2:18.16 | they honor |  |  |  |
| 2:18.17 | with them |  |  |  |
| 2:18.18 | the Jews |  |  |  |
| 2:18.19 | - |  |  |  |
| 2:18.20 | the scroll of |  |  |  |
| 2:18.21 | the Torah |  |  |  |
| 2:18.22 | for |  |  |  |
| 2:18.23 | our enemies |  |  |  |
| 2:18.24 | who |  |  |  |
| 2:18.25 | do not |  |  |  |
| 2:18.26 | know |  |  |  |
| 2:18.27 | what |  |  |  |
| 2:18.28 | this is |  |  |  |
| 2:18.29 | provoke |  |  |  |
| 2:18.30 | us |  |  |  |
| 2:18.31 | saying |  |  |  |
| 2:18.32 | that |  |  |  |
| 2:18.33 | worshippers of |  |  |  |
| 2:18.34 | idols |  |  |  |
| 2:18.35 | we are |  |  |  |
| 2:18.36 | and that |  |  |  |
| 2:18.37 | we bow down |  |  |  |
| 2:18.38 | to the ark |  |  |  |
| 2:18.39 | of wood |  |  |  |
| 2:18.40 | and to the sheets |  |  |  |
| 2:18.41 | of skin |  |  |  |
| 2:18.42 | which |  |  |  |
| 2:18.43 | within it |  |  |  |
| 2:18.44 | Therefore |  |  |  |
| 2:18.45 | I shall set in order |  |  |  |
| 2:18.46 | my reply |  |  |  |
| 2:18.47 | in order |  |  |  |
| 2:18.48 | proper |  |  |  |
| 2:18.49 | as the hand of |  |  |  |
| 2:18.50 | the LORD |  |  |  |
| 2:18.51 | good |  |  |  |
| 2:18.52 | upon me |  |  |  |
| 2:18.53 | namely |  |  |  |
| 2:18.54 | every |  |  |  |
| 2:18.55 | man |  |  |  |
| 2:18.56 | of Israel |  |  |  |
| 2:18.57 | will see |  |  |  |
| 2:18.58 | - |  |  |  |
| 2:18.59 | himself |  |  |  |
| 2:18.60 | obliged |  |  |  |
| 2:18.61 | to rise |  |  |  |
| 2:18.62 | and stand |  |  |  |
| 2:18.63 | at the removal of |  |  |  |
| 2:18.64 | a scroll of |  |  |  |
| 2:18.65 | Torah |  |  |  |
| 2:18.66 | from the ark |  |  |  |
| 2:18.67 | the holy |  |  |  |
| 2:18.68 | until |  |  |  |
| 2:18.69 | it is placed |  |  |  |
| 2:18.70 | on |  |  |  |
| 2:18.71 | the reading platform |  |  | tevah: the reading platform in the synagogue, from which the Torah is read. |
| 2:18.72 | to show it |  |  |  |
| 2:18.73 | to all |  |  |  |
| 2:18.74 | the people |  |  |  |
| 2:18.75 | and to read |  |  |  |
| 2:18.76 | from it |  |  |  |
| 2:18.77 | we find |  |  |  |
| 2:18.78 | that |  |  |  |
| 2:18.79 | so |  |  |  |
| 2:18.80 | they did |  |  |  |
| 2:18.81 | in the days of |  |  |  |
| 2:18.82 | Nehemiah |  |  |  |
| 2:18.83 | as it is written |  |  |  |
| 2:18.84 | and when he opened |  |  | Neh 8:5, in the author's wording. |
| 2:18.85 | stood |  |  |  |
| 2:18.86 | all |  |  |  |
| 2:18.87 | the people |  |  |  |
| 2:18.88 | And this |  |  |  |
| 2:18.89 | they do |  |  |  |
| 2:18.90 | out of |  |  |  |
| 2:18.91 | the fear |  |  |  |
| 2:18.92 | and the honor |  |  |  |
| 2:18.93 | with which |  |  |  |
| 2:18.94 | they honor |  |  |  |
| 2:18.95 | the words of |  |  |  |
| 2:18.96 | God |  |  | Elohim chayyim, 'the living God'; an allusion, see apparatus. |
| 2:18.97 | living |  |  |  |
| 2:18.98 | written |  |  |  |
| 2:18.99 | in it |  |  |  |
| 2:18.100 | and for that same |  |  |  |
| 2:18.101 | reason |  |  |  |
| 2:18.102 | they bow |  |  |  |
| 2:18.103 | their heads |  |  |  |
| 2:18.104 | when carrying |  |  |  |
| 2:18.105 | it |  |  |  |
| 2:18.106 | from the ark |  |  |  |
| 2:18.107 | until |  |  |  |
| 2:18.108 | the reading platform |  |  | tevah: the reading platform, as in token 71. |
| 2:18.109 | and not |  |  |  |
| 2:18.110 | because of |  |  |  |
| 2:18.111 | any |  |  |  |
| 2:18.112 | other |  |  |  |
| 2:18.113 | of |  |  |  |
| 2:18.114 | the reasons |  |  |  |
| 2:18.115 | that |  |  |  |
| 2:18.116 | we shall mention |  |  |  |

Phrases:
- 2:18.1 2:18.2 2:18.3 2:18.4 2:18.5 2:18.6 2:18.7 2:18.8 2:18.9 | And in the book Teshu'at Yisrael by Rabbi Menasseh ben Israel he wrote, and these are his words: | and-in-the-book Teshu'at Yisrael by-Rabbi Menasseh son-of Israel wrote and-these-his-words
- 2:18.10 2:18.11 2:18.12 2:18.13 2:18.14 2:18.15 2:18.16 2:18.17 2:18.18 | “I have been required to make known the customs of honor and submission with which the Jews honor | I-have-been-required to-make-known customs-of honor and-submission with-which-they-honor with-them the-Jews
- 2:18.19 2:18.20 2:18.21 | the Torah scroll, | object scroll-of the-Torah
- 2:18.22 2:18.23 2:18.24 2:18.25 2:18.26 2:18.27 2:18.28 2:18.29 2:18.30 2:18.31 2:18.32 2:18.33 2:18.34 2:18.35 | for our enemies, who do not know what this is, provoke us, saying that we are worshippers of idols, | for our-enemies who not know what this provoke us saying that worshippers-of idols we
- 2:18.36 2:18.37 2:18.38 2:18.39 2:18.40 2:18.41 2:18.42 2:18.43 | and that we bow down to a wooden ark and to the sheets of skin within it. | and-that we-bow to-ark of-wood and-to-sheets of-skin which within-it
- 2:18.44 2:18.45 2:18.46 2:18.47 2:18.48 2:18.49 2:18.50 2:18.51 2:18.52 | Therefore I shall set out my reply in proper order, as the good hand of the LORD is upon me. | therefore I-shall-arrange my-reply in-order proper as-hand-of the-LORD good upon-me
- 2:18.53 2:18.54 2:18.55 2:18.56 2:18.57 2:18.58 2:18.59 2:18.60 2:18.61 2:18.62 2:18.63 2:18.64 2:18.65 2:18.66 2:18.67 2:18.68 2:18.69 2:18.70 2:18.71 2:18.72 2:18.73 2:18.74 2:18.75 2:18.76 | namely, every man of Israel will see himself obliged to rise and stand at the removal of a Torah scroll from the holy ark until it is placed on the reading platform, to show it to all the people and to read from it. | and-it every man of-Israel will-see himself obliged to-rise and-stand at-removal-of scroll-of Torah from-ark the-holy until is-placed on the-reading-platform to-show-it to-all the-people and-to-read from-it
- 2:18.77 2:18.78 2:18.79 2:18.80 2:18.81 2:18.82 2:18.83 2:18.84 2:18.85 2:18.86 2:18.87 | And we find that they did so in the days of Nehemiah, as it is written: ‘And when he opened, all the people stood.’ | and-we-find that-so they-did in-days-of Nehemiah as-it-is-written and-when-he-opened stood all the-people
- 2:18.88 2:18.89 2:18.90 2:18.91 2:18.92 2:18.93 2:18.94 2:18.95 2:18.96 2:18.97 2:18.98 2:18.99 | And this they do out of the fear and honor with which they honor the words of the living God written in it, | and-this they-do from the-fear and-the-honor which they-honor the-words-of God living written in-it
- 2:18.100 2:18.101 2:18.102 2:18.103 2:18.104 2:18.105 2:18.106 2:18.107 2:18.108 | and for that same reason they bow their heads when carrying it from the ark to the reading platform, | and-for-that reason they-bow their-heads when-carrying-it from-the-ark until the-reading-platform
- 2:18.109 2:18.110 2:18.111 2:18.112 2:18.113 2:18.114 2:18.115 2:18.116 | and not because of any other of the reasons that we shall mention.” | and-not because-of other-thing of the-reasons which we-shall-mention

Grammar notes:
- נדרשתי להודיע: a niphal perfect, 'I have been required', with the infinitive להודיע, 'to make known'; the first person is the letter-writer's, quoted by the author.
- ככתוב וכפתחו עמדו כל העם: וכפתחו is a temporal clause ('when he opened'); עמדו is third plural, 'they stood'; the quotation is inside the quoted letter, so it takes single marks.
- והוא כל איש ישראל יראה: והוא introduces the content ('namely'); יראה is a yiqtol with a future sense, 'will see himself obliged'.
- ולא משום דבר אחר מן הטעמים אשר נזכיר: the last clause denies every other reason, so the reasons given are the only ones; the witness has no closing formula, and the quotation runs to the end of the paragraph.

Translator's note: The letter is the one Rabbi Menasseh ben Israel is quoted as writing; 'the reading platform' renders tevah. The witness has no closing formula, so the quotation runs to the end of the paragraph.

Commentary rulings:
- apparatus: Allusion to Ezra 7:28 ('the hand of the LORD upon me'); the English follows the author's wording.
- apparatus: Allusion to Jer 10:10 ('the living God').
- apparatus: Neh 8:5: the author writes 'וכפתחו עמדו כל העם'; his wording is followed.
- apparatus: The title of the book is transliterated per the project scheme; its meaning is at token 2.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

