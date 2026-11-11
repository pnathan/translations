# Review pack: milhamot-hashem perek-2-10 (Milchamot ha-Shem 2:19–2:22)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-10`.

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
- **עבודה / לעבוד** (of God or gods): fixed, "worship / to worship", in prose
  and in quoted verses alike ("ואותו תעבוד" → "and Him you shall worship"), so
  that the author's question, to whom worship belongs, reads the same
  throughout. Use "serve" only for a human master.
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- עבודה אחרת: "other worship" (the Talmudic euphemism for idolatry), noted at
  first occurrence in the unit; with אל אחר = "another god".
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
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
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ב"ד: "court" ("the Great Court" for ב"ד הגדול); "Beit Din" only as a title the
  author gives in transliteration.
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
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
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


## Milchamot ha-Shem 2:19

Hebrew (Wikisource): האחד הפרש גדול בין העובד ומתפלל לאיזה דבר, ובין המכבד איזה דבר לבד. להתפלל ולעבוד עבודה לאיזה דבר אסור לנו, יהיה אדם או מלאך, גשם או רוחני אבל לחלוק כבוד כמו שיכבד הנער את הזקן, העבד את אדוניו, והנקלה לנכבד מותר. וכן מצאנו באברהם אבינו אשר האמין בה' רצץ ושבר מצבות אלילים וכיחש בם, השתחווה ארצה לפני שלשה מלאכים הנצבים עליו אף כי נדמו לו לאנשים. וכן עשה יהושע ביריחו בראותו איש עומד לנגדו ויחשבהו למלאך נפל לפניו ארצה ואם אנשי לבב כמוהם אשר בדרכיהם נלך וממצות שפתותיהם לא נבוש עשו זאת ולא נחשב להם לעון ע"ז ברור הוא כי לכבד ספר תורה אינה עבודת גילולים.

| id | word | punctuation after |
|---|---|---|
| 2:19.1 | האחד |  |
| 2:19.2 | הפרש |  |
| 2:19.3 | גדול |  |
| 2:19.4 | בין |  |
| 2:19.5 | העובד |  |
| 2:19.6 | ומתפלל |  |
| 2:19.7 | לאיזה |  |
| 2:19.8 | דבר | , |
| 2:19.9 | ובין |  |
| 2:19.10 | המכבד |  |
| 2:19.11 | איזה |  |
| 2:19.12 | דבר |  |
| 2:19.13 | לבד | . |
| 2:19.14 | להתפלל |  |
| 2:19.15 | ולעבוד |  |
| 2:19.16 | עבודה |  |
| 2:19.17 | לאיזה |  |
| 2:19.18 | דבר |  |
| 2:19.19 | אסור |  |
| 2:19.20 | לנו | , |
| 2:19.21 | יהיה |  |
| 2:19.22 | אדם |  |
| 2:19.23 | או |  |
| 2:19.24 | מלאך | , |
| 2:19.25 | גשם |  |
| 2:19.26 | או |  |
| 2:19.27 | רוחני |  |
| 2:19.28 | אבל |  |
| 2:19.29 | לחלוק |  |
| 2:19.30 | כבוד |  |
| 2:19.31 | כמו |  |
| 2:19.32 | שיכבד |  |
| 2:19.33 | הנער |  |
| 2:19.34 | את |  |
| 2:19.35 | הזקן | , |
| 2:19.36 | העבד |  |
| 2:19.37 | את |  |
| 2:19.38 | אדוניו | , |
| 2:19.39 | והנקלה |  |
| 2:19.40 | לנכבד |  |
| 2:19.41 | מותר | . |
| 2:19.42 | וכן |  |
| 2:19.43 | מצאנו |  |
| 2:19.44 | באברהם |  |
| 2:19.45 | אבינו |  |
| 2:19.46 | אשר |  |
| 2:19.47 | האמין |  |
| 2:19.48 | בה | ' |
| 2:19.49 | רצץ |  |
| 2:19.50 | ושבר |  |
| 2:19.51 | מצבות |  |
| 2:19.52 | אלילים |  |
| 2:19.53 | וכיחש |  |
| 2:19.54 | בם | , |
| 2:19.55 | השתחווה |  |
| 2:19.56 | ארצה |  |
| 2:19.57 | לפני |  |
| 2:19.58 | שלשה |  |
| 2:19.59 | מלאכים |  |
| 2:19.60 | הנצבים |  |
| 2:19.61 | עליו |  |
| 2:19.62 | אף |  |
| 2:19.63 | כי |  |
| 2:19.64 | נדמו |  |
| 2:19.65 | לו |  |
| 2:19.66 | לאנשים | . |
| 2:19.67 | וכן |  |
| 2:19.68 | עשה |  |
| 2:19.69 | יהושע |  |
| 2:19.70 | ביריחו |  |
| 2:19.71 | בראותו |  |
| 2:19.72 | איש |  |
| 2:19.73 | עומד |  |
| 2:19.74 | לנגדו |  |
| 2:19.75 | ויחשבהו |  |
| 2:19.76 | למלאך |  |
| 2:19.77 | נפל |  |
| 2:19.78 | לפניו |  |
| 2:19.79 | ארצה |  |
| 2:19.80 | ואם |  |
| 2:19.81 | אנשי |  |
| 2:19.82 | לבב |  |
| 2:19.83 | כמוהם |  |
| 2:19.84 | אשר |  |
| 2:19.85 | בדרכיהם |  |
| 2:19.86 | נלך |  |
| 2:19.87 | וממצות |  |
| 2:19.88 | שפתותיהם |  |
| 2:19.89 | לא |  |
| 2:19.90 | נבוש |  |
| 2:19.91 | עשו |  |
| 2:19.92 | זאת |  |
| 2:19.93 | ולא |  |
| 2:19.94 | נחשב |  |
| 2:19.95 | להם |  |
| 2:19.96 | לעון |  |
| 2:19.97 | ע"ז |  |
| 2:19.98 | ברור |  |
| 2:19.99 | הוא |  |
| 2:19.100 | כי |  |
| 2:19.101 | לכבד |  |
| 2:19.102 | ספר |  |
| 2:19.103 | תורה |  |
| 2:19.104 | אינה |  |
| 2:19.105 | עבודת |  |
| 2:19.106 | גילולים | . |


### Translation under review (2:19)

English: The first: there is a great difference between one who worships and prays to something, and one who only honours something. To pray to something and to worship something with worship is forbidden to us, whether it be a man or an angel, corporeal or spiritual; but to bestow honour, as a youth honours an elder and a servant his master, and it is permitted for the lowly to the honoured. So too we find of our father Abraham, who believed in the LORD, that he shattered the idol-monuments and disowned them, and bowed down to the ground before three angels standing over him, even though they appeared to him as men. So too Joshua did at Jericho: when he saw a man standing before him and took him for an angel, he fell to the ground before him. And if men of heart like them, in whose ways we walk and whose commandments of their lips we do not shrink from, did this, and it was not counted to them as a sin of idolatry; it is clear that to honour a Torah scroll is not the worship of idols.

Literal: the-first the-difference great between the-worshipper and-the-prayer to-some-thing, and-between the-honourer some-thing only. to-pray and-to-worship worship to-some-thing forbidden to-us, be-it man or angel, corporeal or spiritual; but to-bestow honour, as the-youth honours the-elder, the-servant [object] his-master, and-the-lowly to-the-honoured is-permitted. and-so we-found in-Abraham our-father who believed in-the-LORD, shattered and-broke monuments-of idols and-disowned them, bowed-down to-the-ground before three angels the-standing over-him, even-though they-appeared to-him as-men. and-so Joshua did at-Jericho when-he-saw a-man standing before-him and-took-him for-an-angel, fell before-him to-the-ground. and-if men-of heart like-them in-whose-ways we-walk and-by-commandments-of their-lips not we-shrink-from did this, and-not was-counted to-them as-sin idolatry; clear it-is that to-honour a-scroll-of the-Torah is-not the-worship-of idols.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:19.1 | the first |  |  | the first of four points: 2:20 'second', 2:21 'third', 2:22 'fourth' |
| 2:19.2 | difference |  |  |  |
| 2:19.3 | great |  |  |  |
| 2:19.4 | between |  |  |  |
| 2:19.5 | the one who worships |  |  |  |
| 2:19.6 | and who prays |  |  |  |
| 2:19.7 | to |  |  |  |
| 2:19.8 | some thing |  |  |  |
| 2:19.9 | and between |  |  |  |
| 2:19.10 | the one who honours |  |  |  |
| 2:19.11 | some |  |  |  |
| 2:19.12 | thing |  |  |  |
| 2:19.13 | only |  |  |  |
| 2:19.14 | to pray |  |  |  |
| 2:19.15 | and to worship |  |  |  |
| 2:19.16 | with worship |  |  | the verb and its noun share one root: 'to worship with worship' |
| 2:19.17 | to |  |  |  |
| 2:19.18 | some thing |  |  |  |
| 2:19.19 | forbidden |  |  |  |
| 2:19.20 | to us |  |  |  |
| 2:19.21 | whether it be |  |  |  |
| 2:19.22 | a man |  |  |  |
| 2:19.23 | or |  |  |  |
| 2:19.24 | an angel |  |  |  |
| 2:19.25 | corporeal |  |  |  |
| 2:19.26 | or |  |  |  |
| 2:19.27 | spiritual |  |  |  |
| 2:19.28 | but |  |  |  |
| 2:19.29 | to bestow |  |  |  |
| 2:19.30 | honour |  |  |  |
| 2:19.31 | as |  |  |  |
| 2:19.32 | honours |  |  |  |
| 2:19.33 | the youth |  |  |  |
| 2:19.34 | (object marker) |  |  | the object marker has no English counterpart |
| 2:19.35 | the elder |  |  |  |
| 2:19.36 | the servant |  |  |  |
| 2:19.37 | (object marker) |  |  | the object marker has no English counterpart |
| 2:19.38 | his master |  |  |  |
| 2:19.39 | and the lowly |  |  |  |
| 2:19.40 | to the honoured |  |  |  |
| 2:19.41 | is permitted |  |  |  |
| 2:19.42 | and so |  |  |  |
| 2:19.43 | we found |  |  |  |
| 2:19.44 | in Abraham |  |  |  |
| 2:19.45 | our father |  |  |  |
| 2:19.46 | who |  |  |  |
| 2:19.47 | believed |  |  |  |
| 2:19.48 | in the LORD |  |  | the Name written as ה׳ (the Tetragrammaton), rendered 'the LORD' |
| 2:19.49 | shattered |  |  |  |
| 2:19.50 | and broke |  |  |  |
| 2:19.51 | the monuments of |  |  | מצבות: standing stones or monuments |
| 2:19.52 | idols |  |  |  |
| 2:19.53 | and disowned |  |  |  |
| 2:19.54 | them |  |  |  |
| 2:19.55 | bowed down |  |  |  |
| 2:19.56 | to the ground |  |  |  |
| 2:19.57 | before |  |  |  |
| 2:19.58 | three |  |  |  |
| 2:19.59 | angels |  |  |  |
| 2:19.60 | who stood |  |  |  |
| 2:19.61 | over him |  |  |  |
| 2:19.62 | even |  |  |  |
| 2:19.63 | though |  |  |  |
| 2:19.64 | appeared |  |  |  |
| 2:19.65 | to him |  |  |  |
| 2:19.66 | as men |  |  |  |
| 2:19.67 | and so |  |  |  |
| 2:19.68 | did |  |  |  |
| 2:19.69 | Joshua |  |  |  |
| 2:19.70 | at Jericho |  |  |  |
| 2:19.71 | when he saw |  |  |  |
| 2:19.72 | a man |  |  |  |
| 2:19.73 | standing |  |  |  |
| 2:19.74 | before him |  |  |  |
| 2:19.75 | and took him |  |  |  |
| 2:19.76 | for an angel |  |  |  |
| 2:19.77 | fell |  |  |  |
| 2:19.78 | before him |  |  |  |
| 2:19.79 | to the ground |  |  |  |
| 2:19.80 | and if |  |  |  |
| 2:19.81 | men of |  |  |  |
| 2:19.82 | heart |  |  |  |
| 2:19.83 | like them |  |  |  |
| 2:19.84 | who |  |  |  |
| 2:19.85 | in their ways |  |  |  |
| 2:19.86 | we walk |  |  |  |
| 2:19.87 | and by the commandments of |  |  |  |
| 2:19.88 | their lips |  |  |  |
| 2:19.89 | we do not |  |  |  |
| 2:19.90 | shrink from |  |  | נבוש: to be ashamed; here 'shrink from' |
| 2:19.91 | did |  |  |  |
| 2:19.92 | this |  |  |  |
| 2:19.93 | and not |  |  |  |
| 2:19.94 | was counted |  |  |  |
| 2:19.95 | to them |  |  |  |
| 2:19.96 | as a sin |  |  |  |
| 2:19.97 | idolatry |  |  | ע"ז, idolatry; literally 'strange worship' |
| 2:19.98 | clear |  |  |  |
| 2:19.99 | it is |  |  |  |
| 2:19.100 | that |  |  |  |
| 2:19.101 | to honour |  |  |  |
| 2:19.102 | a scroll of |  |  |  |
| 2:19.103 | the Torah |  |  |  |
| 2:19.104 | is not |  |  |  |
| 2:19.105 | the worship of |  |  |  |
| 2:19.106 | idols |  |  | גילולים: idols; the root means dung, an insulting word for idols |

Phrases:
- 2:19.1 2:19.2 2:19.3 2:19.4 2:19.5 2:19.6 2:19.7 2:19.8 2:19.9 2:19.10 2:19.11 2:19.12 2:19.13 | The first: there is a great difference between one who worships and prays to something, and one who only honours something. | the-first the-difference great between the-worshipper and-the-prayer to-some-thing and-between the-honourer some-thing only
- 2:19.14 2:19.15 2:19.16 2:19.17 2:19.18 2:19.19 2:19.20 2:19.21 2:19.22 2:19.23 2:19.24 2:19.25 2:19.26 2:19.27 | To pray to something and to worship something with worship is forbidden to us, whether it be a man or an angel, corporeal or spiritual; | to-pray and-to-worship worship to-some-thing forbidden to-us, be-it man or angel, corporeal or spiritual
- 2:19.28 2:19.29 2:19.30 2:19.31 2:19.32 2:19.33 2:19.34 2:19.35 2:19.36 2:19.37 2:19.38 2:19.39 2:19.40 2:19.41 | but to bestow honour, as a youth honours an elder and a servant his master, and it is permitted for the lowly to the honoured. | but to-bestow honour, as the-youth honours the-elder, the-servant [object] his-master, and-the-lowly to-the-honoured is-permitted
- 2:19.42 2:19.43 2:19.44 2:19.45 2:19.46 2:19.47 2:19.48 2:19.49 2:19.50 2:19.51 2:19.52 2:19.53 2:19.54 | So too we find of our father Abraham, who believed in the LORD, that he shattered the idol-monuments and disowned them, | and-so we-found in-Abraham our-father who believed in-the-LORD, shattered and-broke the-monuments-of idols and-disowned them
- 2:19.55 2:19.56 2:19.57 2:19.58 2:19.59 2:19.60 2:19.61 2:19.62 2:19.63 2:19.64 2:19.65 2:19.66 | and bowed down to the ground before three angels standing over him, even though they appeared to him as men. | bowed-down to-the-ground before three angels the-standing over-him, even-though they-appeared to-him as-men
- 2:19.67 2:19.68 2:19.69 2:19.70 2:19.71 2:19.72 2:19.73 2:19.74 2:19.75 2:19.76 2:19.77 2:19.78 2:19.79 | So too Joshua did at Jericho: when he saw a man standing before him and took him for an angel, he fell to the ground before him. | and-so Joshua did at-Jericho when-he-saw a-man standing before-him and-took-him for-an-angel, fell before-him to-the-ground
- 2:19.80 2:19.81 2:19.82 2:19.83 2:19.84 2:19.85 2:19.86 2:19.87 2:19.88 2:19.89 2:19.90 2:19.91 2:19.92 | And if men of heart like them, in whose ways we walk and whose commandments of their lips we do not shrink from, did this, | and-if men-of heart like-them in-whose ways we-walk and-by-commandments-of their-lips not we-shrink-from did this
- 2:19.93 2:19.94 2:19.95 2:19.96 2:19.97 2:19.98 2:19.99 2:19.100 2:19.101 2:19.102 2:19.103 2:19.104 2:19.105 2:19.106 | and it was not counted to them as a sin of idolatry; it is clear that to honour a Torah scroll is not the worship of idols. | and-not was-counted to-them as-sin idolatry; clear it-is that to-honour a-scroll-of the-Torah is-not the-worship-of idols

Grammar notes:
- בין ... ובין frames the contrast: the second half is elliptical, so 'one who only honours' is set against 'one who worships and prays'.
- ולעבוד עבודה: the verb and its noun share a root; the English keeps 'worship' for both, as the Hebrew does.
- כמו שיכבד הנער את הזקן: the comparison gives the model of honour the author permits; the paragraph's argument turns on the difference between the two.
- ואם אנשי לבב ... עשו זאת ... ברור הוא כי: the 'if' clause gives the case of the patriarchs and Joshua; the main clause is 'it is clear that'.

Commentary rulings:
- apparatus: Gen 18:2: the verse behind 'השתחווה ארצה לפני שלשה מלאכים' (Abraham bows to the ground before the three men).
- apparatus: Josh 5:13-14: the verse behind Joshua's falling to the ground before the man who stands before him.

## Milchamot ha-Shem 2:20

Hebrew (Wikisource): שנית היהודים נזהרים מאד בענייני עבודה אחרת, ואף להראות כמכבדה אסור להם. כמו שאמרו בתלמוד וברמב"ם אם לפת איש ארחות דרכו אצל בית ע"ז וישב לו קוץ ברגלו לא ישח אף לא יכרע וישתחוה להוציאו שם למען לא יתראה בעיני העוברים כמשתחוה ומכבד לע"ז, ואם היתה ההשתחויה לפני הארון כעבודה אחרת לא היו היהודים מחזיקים במנהג זה, וכיון שהחזיקו בו מימי אבותיהם אות אמת הוא כי אינה משובשת.

| id | word | punctuation after |
|---|---|---|
| 2:20.1 | שנית |  |
| 2:20.2 | היהודים |  |
| 2:20.3 | נזהרים |  |
| 2:20.4 | מאד |  |
| 2:20.5 | בענייני |  |
| 2:20.6 | עבודה |  |
| 2:20.7 | אחרת | , |
| 2:20.8 | ואף |  |
| 2:20.9 | להראות |  |
| 2:20.10 | כמכבדה |  |
| 2:20.11 | אסור |  |
| 2:20.12 | להם | . |
| 2:20.13 | כמו |  |
| 2:20.14 | שאמרו |  |
| 2:20.15 | בתלמוד |  |
| 2:20.16 | וברמב"ם |  |
| 2:20.17 | אם |  |
| 2:20.18 | לפת |  |
| 2:20.19 | איש |  |
| 2:20.20 | ארחות |  |
| 2:20.21 | דרכו |  |
| 2:20.22 | אצל |  |
| 2:20.23 | בית |  |
| 2:20.24 | ע"ז |  |
| 2:20.25 | וישב |  |
| 2:20.26 | לו |  |
| 2:20.27 | קוץ |  |
| 2:20.28 | ברגלו |  |
| 2:20.29 | לא |  |
| 2:20.30 | ישח |  |
| 2:20.31 | אף |  |
| 2:20.32 | לא |  |
| 2:20.33 | יכרע |  |
| 2:20.34 | וישתחוה |  |
| 2:20.35 | להוציאו |  |
| 2:20.36 | שם |  |
| 2:20.37 | למען |  |
| 2:20.38 | לא |  |
| 2:20.39 | יתראה |  |
| 2:20.40 | בעיני |  |
| 2:20.41 | העוברים |  |
| 2:20.42 | כמשתחוה |  |
| 2:20.43 | ומכבד |  |
| 2:20.44 | לע"ז | , |
| 2:20.45 | ואם |  |
| 2:20.46 | היתה |  |
| 2:20.47 | ההשתחויה |  |
| 2:20.48 | לפני |  |
| 2:20.49 | הארון |  |
| 2:20.50 | כעבודה |  |
| 2:20.51 | אחרת |  |
| 2:20.52 | לא |  |
| 2:20.53 | היו |  |
| 2:20.54 | היהודים |  |
| 2:20.55 | מחזיקים |  |
| 2:20.56 | במנהג |  |
| 2:20.57 | זה | , |
| 2:20.58 | וכיון |  |
| 2:20.59 | שהחזיקו |  |
| 2:20.60 | בו |  |
| 2:20.61 | מימי |  |
| 2:20.62 | אבותיהם |  |
| 2:20.63 | אות |  |
| 2:20.64 | אמת |  |
| 2:20.65 | הוא |  |
| 2:20.66 | כי |  |
| 2:20.67 | אינה |  |
| 2:20.68 | משובשת | . |


### Translation under review (2:20)

English: Second: the Jews are very careful in matters of other worship, and even to appear as one who honours it is forbidden to them. As they said in the Talmud and in the Rambam: "If a man turned aside from the road he was taking near a house of idolatry, and a thorn lodged in his foot, he may not bend, nor even kneel, nor bow down to remove it there, so that he should not appear in the eyes of passers-by as one who bows down and honours idolatry." And if the prostration before the Ark were like worship of another kind, the Jews would not have kept this custom; and since they kept it from the days of their fathers, it is a sign of truth that it is not corrupted.

Literal: second the-Jews are-careful very in-matters-of worship other, and-even to-appear as-honouring-it forbidden to-them. as they-said in-the-Talmud and-in-the-Rambam: if turned-aside a-man the-road of his-way near a-house-of idolatry, and-lodged in-him a-thorn in-his-foot, not bend, nor not kneel, and bow-down to-remove-it there, so-that not appear in-the-eyes-of the-passers-by as-bowing-down and-honouring idolatry; and-if was the-prostration before the-Ark like-worship of-another; not would-be the-Jews keeping in-the-custom of-this, and-since they-kept it from-the-days-of their-fathers, a-sign of-truth it that it-is-not corrupted.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:20.1 | second |  |  |  |
| 2:20.2 | the Jews |  |  |  |
| 2:20.3 | are careful |  |  |  |
| 2:20.4 | very |  |  |  |
| 2:20.5 | in matters of |  |  |  |
| 2:20.6 | worship |  |  |  |
| 2:20.7 | other |  |  |  |
| 2:20.8 | and even |  |  |  |
| 2:20.9 | to appear |  |  |  |
| 2:20.10 | as one who honours it |  |  |  |
| 2:20.11 | forbidden |  |  |  |
| 2:20.12 | to them |  |  |  |
| 2:20.13 | as |  |  |  |
| 2:20.14 | they said |  |  |  |
| 2:20.15 | in the Talmud |  |  |  |
| 2:20.16 | and in the Rambam |  |  |  |
| 2:20.17 | if |  |  |  |
| 2:20.18 | turned aside |  |  |  |
| 2:20.19 | a man |  |  |  |
| 2:20.20 | the paths of |  |  |  |
| 2:20.21 | his way |  |  |  |
| 2:20.22 | near |  |  |  |
| 2:20.23 | a house of |  |  |  |
| 2:20.24 | idolatry |  |  | ע"ז, idolatry: a house of idolatry is a house for the idol's worship |
| 2:20.25 | and lodged |  |  | וישב: from ישב, to sit or settle; a thorn 'lodged' in the foot |
| 2:20.26 | in him |  |  |  |
| 2:20.27 | a thorn |  |  |  |
| 2:20.28 | in his foot |  |  |  |
| 2:20.29 | not |  |  |  |
| 2:20.30 | bend |  |  |  |
| 2:20.31 | nor |  |  |  |
| 2:20.32 | not |  |  |  |
| 2:20.33 | kneel |  |  |  |
| 2:20.34 | and bow down |  |  |  |
| 2:20.35 | to remove it |  |  |  |
| 2:20.36 | there |  |  |  |
| 2:20.37 | so that |  |  |  |
| 2:20.38 | not |  |  |  |
| 2:20.39 | he should appear |  |  |  |
| 2:20.40 | in the eyes of |  |  |  |
| 2:20.41 | the passers-by |  |  |  |
| 2:20.42 | as one who bows down |  |  |  |
| 2:20.43 | and honours |  |  |  |
| 2:20.44 | idolatry |  |  |  |
| 2:20.45 | and if |  |  |  |
| 2:20.46 | were |  |  |  |
| 2:20.47 | the prostration |  |  |  |
| 2:20.48 | before |  |  |  |
| 2:20.49 | the Ark |  |  | the Ark: the cabinet in the synagogue in which the Torah scrolls are kept |
| 2:20.50 | like worship |  |  |  |
| 2:20.51 | of another kind |  |  |  |
| 2:20.52 | not |  |  |  |
| 2:20.53 | would |  |  |  |
| 2:20.54 | the Jews |  |  |  |
| 2:20.55 | keep |  |  |  |
| 2:20.56 | in the custom |  |  |  |
| 2:20.57 | of this |  |  | זה: 'this', the custom of bowing before the Ark |
| 2:20.58 | and since |  |  |  |
| 2:20.59 | they kept |  |  |  |
| 2:20.60 | it |  |  |  |
| 2:20.61 | from the days of |  |  |  |
| 2:20.62 | their fathers |  |  |  |
| 2:20.63 | a sign |  |  |  |
| 2:20.64 | of truth |  |  |  |
| 2:20.65 | it is |  |  |  |
| 2:20.66 | that |  |  |  |
| 2:20.67 | it is not |  |  |  |
| 2:20.68 | corrupted |  |  |  |

Phrases:
- 2:20.1 2:20.2 2:20.3 2:20.4 2:20.5 2:20.6 2:20.7 2:20.8 2:20.9 2:20.10 2:20.11 2:20.12 | Second: the Jews are very careful in matters of other worship, and even to appear as one who honours it is forbidden to them. | second the-Jews are-careful very in-matters-of worship other, and-even to-appear as-honouring-it forbidden to-them
- 2:20.13 2:20.14 2:20.15 2:20.16 | As they said in the Talmud and in the Rambam: | as they-said in-the-Talmud and-in-the-Rambam:
- 2:20.17 2:20.18 2:20.19 2:20.20 2:20.21 2:20.22 2:20.23 2:20.24 2:20.25 2:20.26 2:20.27 2:20.28 | "If a man turned aside from the road he was taking near a house of idolatry, and a thorn lodged in his foot, | if a-man turned-aside the-road of his-way near a-house-of idolatry, and-lodged in-him a-thorn in-his-foot
- 2:20.29 2:20.30 2:20.31 2:20.32 2:20.33 2:20.34 2:20.35 2:20.36 2:20.37 2:20.38 2:20.39 2:20.40 2:20.41 2:20.42 2:20.43 2:20.44 | he may not bend, nor even kneel, nor bow down to remove it there, so that he should not appear in the eyes of passers-by as one who bows down and honours idolatry." | not bend, nor not kneel, and bow-down to-remove-it there, so-that not appear in-the-eyes-of the-passers-by as-bowing-down and-honouring idolatry
- 2:20.45 2:20.46 2:20.47 2:20.48 2:20.49 2:20.50 2:20.51 | And if the prostration before the Ark were like worship of another kind, | and-if was the-prostration before the-Ark like-worship of-another
- 2:20.52 2:20.53 2:20.54 2:20.55 2:20.56 2:20.57 | the Jews would not have kept this custom; | not would-be the-Jews keeping in-the-custom of-this
- 2:20.58 2:20.59 2:20.60 2:20.61 2:20.62 2:20.63 2:20.64 2:20.65 2:20.66 2:20.67 2:20.68 | and since they kept it from the days of their fathers, it is a sign of truth that it is not corrupted. | and-since they-kept it from-the-days-of their-fathers, a-sign of-truth it that it-is-not corrupted

Grammar notes:
- כמו שאמרו בתלמוד וברמב"ם: the formula introduces a passage from the Talmud and the Rambam, given in the author's own wording; no reference is given.
- לא ישח אף לא יכרע וישתחוה להוציאו: three verbs of bowing, two negated, the last joined to the purpose clause 'so that he should not appear'.
- ואם היתה ... לא היו היהודים מחזיקים: a counterfactual, both clauses in the past; the author argues from what the Jews would not have done.
- וכיון שהחזיקו בו מימי אבותיהם אות אמת הוא כי אינה משובשת: the argument from custom; 'since they kept it' is causal, and 'that it is not corrupted' gives the sign's content.

Commentary rulings:
- apparatus: The Talmud and Rambam passage is given in the author's own paraphrase; he names no tractate or chapter for it.

## Milchamot ha-Shem 2:21

Hebrew (Wikisource): שלישית אחד ממנהגי הכבוד אשר כבדו לפנים את הע"ז היה הנשיקה וכן מצינו במקרא והשארתי בישראל שבעת אלפים כל הברכים אשר לא כרעו לבעל וכל הפה אשר לא נשק לו, ונאמר זובחי אדם עגלים ישקון. ואם כן הוא יהיו כל המנשקים את ספר התורה עובדי ע"ז? אבל לא כן הדבר, כי אין זה רק כבוד שיכבדו את תורתם ולא עבודה, ולכן ההשתחויה מול הארון הוא ג"כ רק כבוד ולא עבודה.

| id | word | punctuation after |
|---|---|---|
| 2:21.1 | שלישית |  |
| 2:21.2 | אחד |  |
| 2:21.3 | ממנהגי |  |
| 2:21.4 | הכבוד |  |
| 2:21.5 | אשר |  |
| 2:21.6 | כבדו |  |
| 2:21.7 | לפנים |  |
| 2:21.8 | את |  |
| 2:21.9 | הע"ז |  |
| 2:21.10 | היה |  |
| 2:21.11 | הנשיקה |  |
| 2:21.12 | וכן |  |
| 2:21.13 | מצינו |  |
| 2:21.14 | במקרא |  |
| 2:21.15 | והשארתי |  |
| 2:21.16 | בישראל |  |
| 2:21.17 | שבעת |  |
| 2:21.18 | אלפים |  |
| 2:21.19 | כל |  |
| 2:21.20 | הברכים |  |
| 2:21.21 | אשר |  |
| 2:21.22 | לא |  |
| 2:21.23 | כרעו |  |
| 2:21.24 | לבעל |  |
| 2:21.25 | וכל |  |
| 2:21.26 | הפה |  |
| 2:21.27 | אשר |  |
| 2:21.28 | לא |  |
| 2:21.29 | נשק |  |
| 2:21.30 | לו | , |
| 2:21.31 | ונאמר |  |
| 2:21.32 | זובחי |  |
| 2:21.33 | אדם |  |
| 2:21.34 | עגלים |  |
| 2:21.35 | ישקון | . |
| 2:21.36 | ואם |  |
| 2:21.37 | כן |  |
| 2:21.38 | הוא |  |
| 2:21.39 | יהיו |  |
| 2:21.40 | כל |  |
| 2:21.41 | המנשקים |  |
| 2:21.42 | את |  |
| 2:21.43 | ספר |  |
| 2:21.44 | התורה |  |
| 2:21.45 | עובדי |  |
| 2:21.46 | ע"ז | ? |
| 2:21.47 | אבל |  |
| 2:21.48 | לא |  |
| 2:21.49 | כן |  |
| 2:21.50 | הדבר | , |
| 2:21.51 | כי |  |
| 2:21.52 | אין |  |
| 2:21.53 | זה |  |
| 2:21.54 | רק |  |
| 2:21.55 | כבוד |  |
| 2:21.56 | שיכבדו |  |
| 2:21.57 | את |  |
| 2:21.58 | תורתם |  |
| 2:21.59 | ולא |  |
| 2:21.60 | עבודה | , |
| 2:21.61 | ולכן |  |
| 2:21.62 | ההשתחויה |  |
| 2:21.63 | מול |  |
| 2:21.64 | הארון |  |
| 2:21.65 | הוא |  |
| 2:21.66 | ג"כ |  |
| 2:21.67 | רק |  |
| 2:21.68 | כבוד |  |
| 2:21.69 | ולא |  |
| 2:21.70 | עבודה | . |


### Translation under review (2:21)

English: Third: one of the customs of honour by which they once honoured idolatry was kissing. So too we find in Scripture: "Yet I have left in Israel seven thousand, all the knees that have not bowed to Baal, and every mouth that has not kissed him." And it is said: "Those who sacrifice men kiss calves." If so, will all who kiss the Torah scroll be worshippers of idols? But the matter is not so, for this is only honour, that they honour their Torah, and not worship; and therefore the prostration before the Ark is also only honour and not worship.

Literal: third one of-the-customs-of the-honour which they-honoured formerly [object] the-idolatry was the-kissing. so-too we-found in-Scripture: I-have-left in-Israel seven-thousand, all the-knees which not bowed to-Baal, and-every the-mouth which not kissed to-him; and-it-is-said: sacrificers-of men calves kiss. and-if so then will-be all the-kissers [object] scroll-of the-Torah worshippers-of idols? but the-matter not so, for not is this only honour that they-honour [object] their-Torah, and-not worship; and-therefore the-prostration before the-Ark it also only honour and-not worship.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:21.1 | third |  |  |  |
| 2:21.2 | one |  |  |  |
| 2:21.3 | of the customs of |  |  |  |
| 2:21.4 | honour |  |  |  |
| 2:21.5 | by which |  |  |  |
| 2:21.6 | they honoured |  |  |  |
| 2:21.7 | formerly |  |  |  |
| 2:21.8 | (object marker) |  |  |  |
| 2:21.9 | idolatry |  |  | ע"ז, idolatry |
| 2:21.10 | was |  |  |  |
| 2:21.11 | kissing |  |  |  |
| 2:21.12 | so too |  |  |  |
| 2:21.13 | we find |  |  |  |
| 2:21.14 | in Scripture |  |  |  |
| 2:21.15 | I have left |  |  | והשארתי: 'I have left', God's words to Elijah in the verse; the author takes it as a sign of kissing Baal |
| 2:21.16 | in Israel |  |  |  |
| 2:21.17 | seven |  |  |  |
| 2:21.18 | thousand |  |  |  |
| 2:21.19 | all |  |  |  |
| 2:21.20 | the knees |  |  |  |
| 2:21.21 | that |  |  |  |
| 2:21.22 | not |  |  |  |
| 2:21.23 | bowed |  |  |  |
| 2:21.24 | to Baal |  |  | Baal: the god of the Canaanites, whose worship Elijah opposed (1 Kgs 18) |
| 2:21.25 | and every |  |  |  |
| 2:21.26 | mouth |  |  |  |
| 2:21.27 | that |  |  |  |
| 2:21.28 | not |  |  |  |
| 2:21.29 | kissed |  |  |  |
| 2:21.30 | him |  |  |  |
| 2:21.31 | and it is said |  |  |  |
| 2:21.32 | those who sacrifice |  |  |  |
| 2:21.33 | men |  |  |  |
| 2:21.34 | calves |  |  |  |
| 2:21.35 | kiss |  |  |  |
| 2:21.36 | and if |  |  |  |
| 2:21.37 | so |  |  |  |
| 2:21.38 | then |  |  | ואם כן: 'if so'; 'then' introduces the consequence |
| 2:21.39 | will be |  |  |  |
| 2:21.40 | all |  |  |  |
| 2:21.41 | those who kiss |  |  |  |
| 2:21.42 | (object marker) |  |  |  |
| 2:21.43 | a scroll of |  |  |  |
| 2:21.44 | the Torah |  |  |  |
| 2:21.45 | worshippers of |  |  |  |
| 2:21.46 | idols |  |  | עובדי ע"ז: worshippers of idols |
| 2:21.47 | but |  |  |  |
| 2:21.48 | not |  |  |  |
| 2:21.49 | so |  |  |  |
| 2:21.50 | the matter |  |  |  |
| 2:21.51 | for |  |  |  |
| 2:21.52 | there is not |  |  |  |
| 2:21.53 | this |  |  |  |
| 2:21.54 | only |  |  |  |
| 2:21.55 | honour |  |  |  |
| 2:21.56 | that they honour |  |  |  |
| 2:21.57 | (object marker) |  |  |  |
| 2:21.58 | their Torah |  |  |  |
| 2:21.59 | and not |  |  |  |
| 2:21.60 | worship |  |  |  |
| 2:21.61 | and therefore |  |  |  |
| 2:21.62 | the prostration |  |  |  |
| 2:21.63 | before |  |  |  |
| 2:21.64 | the Ark |  |  |  |
| 2:21.65 | is |  |  |  |
| 2:21.66 | also |  |  | ג"כ: 'also' (גם כן) |
| 2:21.67 | only |  |  |  |
| 2:21.68 | honour |  |  |  |
| 2:21.69 | and not |  |  |  |
| 2:21.70 | worship |  |  |  |

Phrases:
- 2:21.1 2:21.2 2:21.3 2:21.4 2:21.5 2:21.6 2:21.7 2:21.8 2:21.9 2:21.10 2:21.11 | Third: one of the customs of honour by which they once honoured idolatry was kissing. | third one of-the-customs-of the-honour by-which they-honoured formerly [object] idolatry was the-kissing
- 2:21.12 2:21.13 2:21.14 2:21.15 2:21.16 2:21.17 2:21.18 2:21.19 2:21.20 2:21.21 2:21.22 2:21.23 2:21.24 2:21.25 2:21.26 2:21.27 2:21.28 2:21.29 2:21.30 | So too we find in Scripture: "Yet I have left in Israel seven thousand, all the knees that have not bowed to Baal, and every mouth that has not kissed him." | and-so we-found in-Scripture: I-have-left in-Israel seven-thousand, all the-knees which not bowed to-Baal, and-every the-mouth which not kissed to-him
- 2:21.31 2:21.32 2:21.33 2:21.34 2:21.35 | And it is said: "Those who sacrifice men kiss calves." | and-it-is-said: sacrificers-of men calves kiss
- 2:21.36 2:21.37 2:21.38 2:21.39 2:21.40 2:21.41 2:21.42 2:21.43 2:21.44 2:21.45 2:21.46 | If so, will all who kiss the Torah scroll be worshippers of idols? | and-if so then will-be all the-kissers [object] scroll-of the-Torah worshippers-of idols?
- 2:21.47 2:21.48 2:21.49 2:21.50 2:21.51 2:21.52 2:21.53 2:21.54 2:21.55 2:21.56 2:21.57 2:21.58 2:21.59 2:21.60 | But the matter is not so, for this is only honour, that they honour their Torah, and not worship; | but the-matter not so, for not is this only honour that they-honour [object] their-Torah and-not worship
- 2:21.61 2:21.62 2:21.63 2:21.64 2:21.65 2:21.66 2:21.67 2:21.68 2:21.69 2:21.70 | and therefore the prostration before the Ark is also only honour and not worship. | and-therefore the-prostration before the-Ark it also only honour and-not worship

Grammar notes:
- וכן מצינו במקרא: the verse quoted is 1 Kgs 19:18, in the author's wording; 'we find' introduces it without a formula.
- ונאמר זובחי אדם עגלים ישקון: the second verse, Hos 13:2; the subject is the sacrificers, and the verb is the plural 'kiss'.
- ואם כן הוא יהיו כל המנשקים ... עובדי ע"ז?: a question about the future, answered by 'but this is not so'.
- כי אין זה רק כבוד שיכבדו את תורתם ולא עבודה: 'for this is only honour, that they honour their Torah and not worship' explains 'this'.

Commentary rulings:
- apparatus: 1 Kgs 19:18: the verse quoted in the English, in the author's wording.
- apparatus: Hos 13:2: the verse quoted in the English.

## Milchamot ha-Shem 2:22

Hebrew (Wikisource): רביעית הנסיון הורה אותנו כי בכל לשונות הגוים נמצאו מנהגי הכבוד יכבדו בם איש את רעהו בהיפגשם וכן יכופו ראשיהם זה לעומת זה, והיה גם זה לנו לאות שאינה עבודת גלולים כי אם מנהגי כבוד.

| id | word | punctuation after |
|---|---|---|
| 2:22.1 | רביעית |  |
| 2:22.2 | הנסיון |  |
| 2:22.3 | הורה |  |
| 2:22.4 | אותנו |  |
| 2:22.5 | כי |  |
| 2:22.6 | בכל |  |
| 2:22.7 | לשונות |  |
| 2:22.8 | הגוים |  |
| 2:22.9 | נמצאו |  |
| 2:22.10 | מנהגי |  |
| 2:22.11 | הכבוד |  |
| 2:22.12 | יכבדו |  |
| 2:22.13 | בם |  |
| 2:22.14 | איש |  |
| 2:22.15 | את |  |
| 2:22.16 | רעהו |  |
| 2:22.17 | בהיפגשם |  |
| 2:22.18 | וכן |  |
| 2:22.19 | יכופו |  |
| 2:22.20 | ראשיהם |  |
| 2:22.21 | זה |  |
| 2:22.22 | לעומת |  |
| 2:22.23 | זה | , |
| 2:22.24 | והיה |  |
| 2:22.25 | גם |  |
| 2:22.26 | זה |  |
| 2:22.27 | לנו |  |
| 2:22.28 | לאות |  |
| 2:22.29 | שאינה |  |
| 2:22.30 | עבודת |  |
| 2:22.31 | גלולים |  |
| 2:22.32 | כי |  |
| 2:22.33 | אם |  |
| 2:22.34 | מנהגי |  |
| 2:22.35 | כבוד | . |


### Translation under review (2:22)

English: Fourth: experience has taught us that in all the languages of the nations customs of honour are found, by which a man honours his fellow when they meet, and so they bow their heads, one opposite the other, and this too has been for us a sign that it is not the worship of idols, but customs of honour.

Literal: fourth experience taught us that in-all languages-of the-nations were-found customs-of the-honour; they-honour by-them a-man [object] his-fellow when-they-meet; and-so they-bow their-heads one opposite the-other; and-it-was this too for-us for-a-sign that-it-is-not the-worship-of idols; but customs-of honour.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:22.1 | fourth |  |  |  |
| 2:22.2 | experience |  |  |  |
| 2:22.3 | has taught |  |  |  |
| 2:22.4 | us |  |  |  |
| 2:22.5 | that |  |  |  |
| 2:22.6 | in all |  |  |  |
| 2:22.7 | languages |  |  |  |
| 2:22.8 | of the nations |  |  |  |
| 2:22.9 | are found |  |  |  |
| 2:22.10 | customs of |  |  |  |
| 2:22.11 | honour |  |  |  |
| 2:22.12 | they honour |  |  |  |
| 2:22.13 | by them |  |  | בם: 'by them', referring back to מנהגי הכבוד ('by which') |
| 2:22.14 | a man |  |  |  |
| 2:22.15 | (object marker) |  |  |  |
| 2:22.16 | his fellow |  |  |  |
| 2:22.17 | when they meet |  |  |  |
| 2:22.18 | and so |  |  |  |
| 2:22.19 | they bow |  |  |  |
| 2:22.20 | their heads |  |  |  |
| 2:22.21 | one |  |  |  |
| 2:22.22 | opposite |  |  |  |
| 2:22.23 | the other |  |  |  |
| 2:22.24 | and it was |  |  |  |
| 2:22.25 | too |  |  |  |
| 2:22.26 | this |  |  |  |
| 2:22.27 | for us |  |  |  |
| 2:22.28 | a sign |  |  |  |
| 2:22.29 | that it is not |  |  |  |
| 2:22.30 | the worship of |  |  |  |
| 2:22.31 | idols |  |  | גלולים: idols; see the note on 2:19.106 |
| 2:22.32 | but |  |  |  |
| 2:22.33 | only |  |  |  |
| 2:22.34 | customs of |  |  |  |
| 2:22.35 | honour |  |  |  |

Phrases:
- 2:22.1 2:22.2 2:22.3 2:22.4 2:22.5 2:22.6 2:22.7 2:22.8 2:22.9 2:22.10 2:22.11 | Fourth: experience has taught us that in all the languages of the nations customs of honour are found, | fourth experience taught us that in-all languages-of the-nations were-found customs-of honour
- 2:22.12 2:22.13 2:22.14 2:22.15 2:22.16 2:22.17 | by which a man honours his fellow when they meet, | they-honour by-them a-man [object] his-fellow in-their-meeting
- 2:22.18 2:22.19 2:22.20 2:22.21 2:22.22 2:22.23 | and so they bow their heads, one opposite the other, | and-so they-bow their-heads one opposite the-other
- 2:22.24 2:22.25 2:22.26 2:22.27 2:22.28 2:22.29 2:22.30 2:22.31 | and this too has been for us a sign that it is not the worship of idols, | and-it-was this too for-us for-a-sign that-it-is-not the-worship-of idols
- 2:22.32 2:22.33 2:22.34 2:22.35 | but customs of honour. | but customs-of honour

Grammar notes:
- הנסיון הורה אותנו: the subject is 'experience' and the verb is hiphil, 'taught us'.
- יכבדו בם איש את רעהו בהיפגשם: 'by them' refers back to the customs; ב plus the infinitive with a suffix gives 'when they meet'.
- וכן יכופו ראשיהם זה לעומת זה: 'one opposite the other' is the reciprocal phrase of the bowing.
- והיה גם זה לנו לאות שאינה עבודת גלולים: the sign (אות) is the clause that follows: 'that it is not the worship of idols'.

Commentary rulings:

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

