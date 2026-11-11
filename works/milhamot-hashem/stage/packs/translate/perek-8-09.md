# Translation pack: milhamot-hashem perek-8-09 (Milchamot ha-Shem 8:11–8:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-09`.

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
- הר' / הרב before a book title (הר' כסא אליהו, הר' עץ יוסף): "the Rabbi, the author
  of [title]", with a note; the title is never a personal name.
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
- האלהי before a name: "the divine", with a note; never dropped.
- דרך ארץ: "worldly occupation", as in the project's Avot.
- Mishnah/Talmud references the author lacks are not added to `apparatus`.
- Names of works: "Guide for the Perplexed" is rendered as the Moreh
  Nevukhim; ממלכי as "of the kings of".
- Closing formulas (וחתם דבריו, עכ"ל, ע"כ) stand outside the closing quotation
  mark; opening formulas (וז"ל, כ' וז"ל) stand outside the opening one.
- לדעת before a noun: "according to the understanding of" when the sentence
  has no infinitive sense; the infinitive in `readings`.
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
- תשבורת: "mathematics" (mensuration); ותוכנים: "astronomy" (plural of תכונה).
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
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
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


## Milchamot ha-Shem 8:11

Hebrew (Wikisource): ומ"מ אפשר שלא ירדו לעמקה לדעת את מרורות הפתנים אשר בקרבה, רק עשו כאשר עשינו אנחנו שנים רבות שפירשנו ממנה מפני שקרא עליה ערער. ולא דברנו אודותיה כמצווה עלינו להודיע את חקי האלהים ואת תורותיו ולמי ראוי לעבוד! ולכן חרה אף השי"ת בנו, ויענישינו על מריינו, שהיינו שומעים דברים המביאים לידי טעות, ועל שלא מחינו בכבוד קונינו. הסית השי"ת בנו את ארך הלשון אשר כחזקתו גבה לבו עד להשחית. וישיאו לבו ויצרו הרע להתגולל להוציא דבה רעה על לומדי תורה לשמה על אמתתה במשנה ובתלמוד. שהם מינים גמורים ושבבית שני תיקן רבן גמליאל על ידי שמואל הקטן ברכת המינים, ושכל התלמידים הלומדים משנה והרמב"ם יצאו לתרבות רעה על שאין לומדים זהר ותיקונים! ומיום ליום יביע אומר מלשינות ושקר לפני שרי הממשלה, שהמה אלה מכתבי עמל. ורצונם למרוד בממשלה הרוממה ולחסות תחת ממשלה אנגליז, ולהמריד את יושבי הארץ לחסות בממשלה אחרת נכריה ועל ידי רכילותו זאת בקש הואלי (המושל) לנתשם מעל אדמתם אל ארץ אחרת, ויתן אותנו בבית הסהר יותר מחדש ימים כשלשים איש, ויפל פחד ורעדה בלב העם מהדר גאונו ועריצותו! ושוב המציא במזימותיו רכילות אחרת במכתבים אל האימאם יחיא אלמתוכל יר"ה. שקצת מן היהודים רוצים להתנשא על המושלמנים וללבוש בגדי כבוד, ולרכוב על סוסים. ומבקשים בתי ספר ללמוד שפה תורכית ושאר לשונות כדי להתגדל על האיסלאם. ואלמלא אני מלמדם ומדריכם בדרכי כניעת היהדות ואיך הם מחוייבים לקבל עליהם עול האיסלאם. כי עתה פרצו ועלו על האיסלאם. ועל ידי זה זכה ליתן לו רשיון מן האימאם יר"ה לרדות בעם ה' אלה ולעשות כחפצו, אך לימודי שפת תוגרמה וידיעת גלילות הארץ וגבולותיה וחשבון ותשבורת וכיוצא לא אסר אותן האימאם יר"ה, כי הוא איש משכיל ונבון, ובמכתבו כתב שאין למנוע הלימודים האלה מנערי בני ישראל. ובעל הלשון הרחיב עוד את פיו לפני אחד השופטים של האימאם בפנינו ג"כ להלשין עלינו. כי אנחנו עוזבי דת הישראלית, ולומדים בספרי פילסופי יון, ומכחישים אמונת מציאות השדים. ובאמת לא נמצא אצלינו שום ספר מספרי יון, ולא זולתם מספרי הפילסופים החיצונים כלל! ורק זדון לבו השיאו להלשין עלינו שקר, והרבה שנאת חנם במחנה העברים על לא דבר! ויוסף עוד לחטוא לשלח מכתבים אל כל כפרי ועיירות התימן להוציא ולפרסם את הרבה הרעה הזאת עלינו בשם מינים! וערום יערים לעשות נוסחי דבות רעות לאמר, הן כבר עשינו משפט ברשעים האלה ונביא אותם אל בית הכלא, ועליכם המצוה למהר במכתבי קובלנא אל האימאם יר"ה, כי אלה מקלקלי דתינו ומפסידי אמונתינו, מהרו, חושו, חושו, להגדיל המדורה עליהם! ולא היה כותבם בכתב ידו, רק מוסרם לילדים מקרוביו כדי להעתיקם בכתב ידם, ושולחם לכל ערי התימן שלא בכתב ידו, כדי שיאמר לא פעלתי און.

| id | word | punctuation after |
|---|---|---|
| 8:11.1 | ומ"מ |  |
| 8:11.2 | אפשר |  |
| 8:11.3 | שלא |  |
| 8:11.4 | ירדו |  |
| 8:11.5 | לעמקה |  |
| 8:11.6 | לדעת |  |
| 8:11.7 | את |  |
| 8:11.8 | מרורות |  |
| 8:11.9 | הפתנים |  |
| 8:11.10 | אשר |  |
| 8:11.11 | בקרבה | , |
| 8:11.12 | רק |  |
| 8:11.13 | עשו |  |
| 8:11.14 | כאשר |  |
| 8:11.15 | עשינו |  |
| 8:11.16 | אנחנו |  |
| 8:11.17 | שנים |  |
| 8:11.18 | רבות |  |
| 8:11.19 | שפירשנו |  |
| 8:11.20 | ממנה |  |
| 8:11.21 | מפני |  |
| 8:11.22 | שקרא |  |
| 8:11.23 | עליה |  |
| 8:11.24 | ערער | . |
| 8:11.25 | ולא |  |
| 8:11.26 | דברנו |  |
| 8:11.27 | אודותיה |  |
| 8:11.28 | כמצווה |  |
| 8:11.29 | עלינו |  |
| 8:11.30 | להודיע |  |
| 8:11.31 | את |  |
| 8:11.32 | חקי |  |
| 8:11.33 | האלהים |  |
| 8:11.34 | ואת |  |
| 8:11.35 | תורותיו |  |
| 8:11.36 | ולמי |  |
| 8:11.37 | ראוי |  |
| 8:11.38 | לעבוד | ! |
| 8:11.39 | ולכן |  |
| 8:11.40 | חרה |  |
| 8:11.41 | אף |  |
| 8:11.42 | השי"ת |  |
| 8:11.43 | בנו | , |
| 8:11.44 | ויענישינו |  |
| 8:11.45 | על |  |
| 8:11.46 | מריינו | , |
| 8:11.47 | שהיינו |  |
| 8:11.48 | שומעים |  |
| 8:11.49 | דברים |  |
| 8:11.50 | המביאים |  |
| 8:11.51 | לידי |  |
| 8:11.52 | טעות | , |
| 8:11.53 | ועל |  |
| 8:11.54 | שלא |  |
| 8:11.55 | מחינו |  |
| 8:11.56 | בכבוד |  |
| 8:11.57 | קונינו | . |
| 8:11.58 | הסית |  |
| 8:11.59 | השי"ת |  |
| 8:11.60 | בנו |  |
| 8:11.61 | את |  |
| 8:11.62 | ארך |  |
| 8:11.63 | הלשון |  |
| 8:11.64 | אשר |  |
| 8:11.65 | כחזקתו |  |
| 8:11.66 | גבה |  |
| 8:11.67 | לבו |  |
| 8:11.68 | עד |  |
| 8:11.69 | להשחית | . |
| 8:11.70 | וישיאו |  |
| 8:11.71 | לבו |  |
| 8:11.72 | ויצרו |  |
| 8:11.73 | הרע |  |
| 8:11.74 | להתגולל |  |
| 8:11.75 | להוציא |  |
| 8:11.76 | דבה |  |
| 8:11.77 | רעה |  |
| 8:11.78 | על |  |
| 8:11.79 | לומדי |  |
| 8:11.80 | תורה |  |
| 8:11.81 | לשמה |  |
| 8:11.82 | על |  |
| 8:11.83 | אמתתה |  |
| 8:11.84 | במשנה |  |
| 8:11.85 | ובתלמוד | . |
| 8:11.86 | שהם |  |
| 8:11.87 | מינים |  |
| 8:11.88 | גמורים |  |
| 8:11.89 | ושבבית |  |
| 8:11.90 | שני |  |
| 8:11.91 | תיקן |  |
| 8:11.92 | רבן |  |
| 8:11.93 | גמליאל |  |
| 8:11.94 | על |  |
| 8:11.95 | ידי |  |
| 8:11.96 | שמואל |  |
| 8:11.97 | הקטן |  |
| 8:11.98 | ברכת |  |
| 8:11.99 | המינים | , |
| 8:11.100 | ושכל |  |
| 8:11.101 | התלמידים |  |
| 8:11.102 | הלומדים |  |
| 8:11.103 | משנה |  |
| 8:11.104 | והרמב"ם |  |
| 8:11.105 | יצאו |  |
| 8:11.106 | לתרבות |  |
| 8:11.107 | רעה |  |
| 8:11.108 | על |  |
| 8:11.109 | שאין |  |
| 8:11.110 | לומדים |  |
| 8:11.111 | זהר |  |
| 8:11.112 | ותיקונים | ! |
| 8:11.113 | ומיום |  |
| 8:11.114 | ליום |  |
| 8:11.115 | יביע |  |
| 8:11.116 | אומר |  |
| 8:11.117 | מלשינות |  |
| 8:11.118 | ושקר |  |
| 8:11.119 | לפני |  |
| 8:11.120 | שרי |  |
| 8:11.121 | הממשלה | , |
| 8:11.122 | שהמה |  |
| 8:11.123 | אלה |  |
| 8:11.124 | מכתבי |  |
| 8:11.125 | עמל | . |
| 8:11.126 | ורצונם |  |
| 8:11.127 | למרוד |  |
| 8:11.128 | בממשלה |  |
| 8:11.129 | הרוממה |  |
| 8:11.130 | ולחסות |  |
| 8:11.131 | תחת |  |
| 8:11.132 | ממשלה |  |
| 8:11.133 | אנגליז | , |
| 8:11.134 | ולהמריד |  |
| 8:11.135 | את |  |
| 8:11.136 | יושבי |  |
| 8:11.137 | הארץ |  |
| 8:11.138 | לחסות |  |
| 8:11.139 | בממשלה |  |
| 8:11.140 | אחרת |  |
| 8:11.141 | נכריה |  |
| 8:11.142 | ועל |  |
| 8:11.143 | ידי |  |
| 8:11.144 | רכילותו |  |
| 8:11.145 | זאת |  |
| 8:11.146 | בקש |  |
| 8:11.147 | הואלי |  |
| 8:11.148 | המושל | ) |
| 8:11.149 | לנתשם |  |
| 8:11.150 | מעל |  |
| 8:11.151 | אדמתם |  |
| 8:11.152 | אל |  |
| 8:11.153 | ארץ |  |
| 8:11.154 | אחרת | , |
| 8:11.155 | ויתן |  |
| 8:11.156 | אותנו |  |
| 8:11.157 | בבית |  |
| 8:11.158 | הסהר |  |
| 8:11.159 | יותר |  |
| 8:11.160 | מחדש |  |
| 8:11.161 | ימים |  |
| 8:11.162 | כשלשים |  |
| 8:11.163 | איש | , |
| 8:11.164 | ויפל |  |
| 8:11.165 | פחד |  |
| 8:11.166 | ורעדה |  |
| 8:11.167 | בלב |  |
| 8:11.168 | העם |  |
| 8:11.169 | מהדר |  |
| 8:11.170 | גאונו |  |
| 8:11.171 | ועריצותו | ! |
| 8:11.172 | ושוב |  |
| 8:11.173 | המציא |  |
| 8:11.174 | במזימותיו |  |
| 8:11.175 | רכילות |  |
| 8:11.176 | אחרת |  |
| 8:11.177 | במכתבים |  |
| 8:11.178 | אל |  |
| 8:11.179 | האימאם |  |
| 8:11.180 | יחיא |  |
| 8:11.181 | אלמתוכל |  |
| 8:11.182 | יר"ה | . |
| 8:11.183 | שקצת |  |
| 8:11.184 | מן |  |
| 8:11.185 | היהודים |  |
| 8:11.186 | רוצים |  |
| 8:11.187 | להתנשא |  |
| 8:11.188 | על |  |
| 8:11.189 | המושלמנים |  |
| 8:11.190 | וללבוש |  |
| 8:11.191 | בגדי |  |
| 8:11.192 | כבוד | , |
| 8:11.193 | ולרכוב |  |
| 8:11.194 | על |  |
| 8:11.195 | סוסים | . |
| 8:11.196 | ומבקשים |  |
| 8:11.197 | בתי |  |
| 8:11.198 | ספר |  |
| 8:11.199 | ללמוד |  |
| 8:11.200 | שפה |  |
| 8:11.201 | תורכית |  |
| 8:11.202 | ושאר |  |
| 8:11.203 | לשונות |  |
| 8:11.204 | כדי |  |
| 8:11.205 | להתגדל |  |
| 8:11.206 | על |  |
| 8:11.207 | האיסלאם | . |
| 8:11.208 | ואלמלא |  |
| 8:11.209 | אני |  |
| 8:11.210 | מלמדם |  |
| 8:11.211 | ומדריכם |  |
| 8:11.212 | בדרכי |  |
| 8:11.213 | כניעת |  |
| 8:11.214 | היהדות |  |
| 8:11.215 | ואיך |  |
| 8:11.216 | הם |  |
| 8:11.217 | מחוייבים |  |
| 8:11.218 | לקבל |  |
| 8:11.219 | עליהם |  |
| 8:11.220 | עול |  |
| 8:11.221 | האיסלאם | . |
| 8:11.222 | כי |  |
| 8:11.223 | עתה |  |
| 8:11.224 | פרצו |  |
| 8:11.225 | ועלו |  |
| 8:11.226 | על |  |
| 8:11.227 | האיסלאם | . |
| 8:11.228 | ועל |  |
| 8:11.229 | ידי |  |
| 8:11.230 | זה |  |
| 8:11.231 | זכה |  |
| 8:11.232 | ליתן |  |
| 8:11.233 | לו |  |
| 8:11.234 | רשיון |  |
| 8:11.235 | מן |  |
| 8:11.236 | האימאם |  |
| 8:11.237 | יר"ה |  |
| 8:11.238 | לרדות |  |
| 8:11.239 | בעם |  |
| 8:11.240 | ה | ' |
| 8:11.241 | אלה |  |
| 8:11.242 | ולעשות |  |
| 8:11.243 | כחפצו | , |
| 8:11.244 | אך |  |
| 8:11.245 | לימודי |  |
| 8:11.246 | שפת |  |
| 8:11.247 | תוגרמה |  |
| 8:11.248 | וידיעת |  |
| 8:11.249 | גלילות |  |
| 8:11.250 | הארץ |  |
| 8:11.251 | וגבולותיה |  |
| 8:11.252 | וחשבון |  |
| 8:11.253 | ותשבורת |  |
| 8:11.254 | וכיוצא |  |
| 8:11.255 | לא |  |
| 8:11.256 | אסר |  |
| 8:11.257 | אותן |  |
| 8:11.258 | האימאם |  |
| 8:11.259 | יר"ה | , |
| 8:11.260 | כי |  |
| 8:11.261 | הוא |  |
| 8:11.262 | איש |  |
| 8:11.263 | משכיל |  |
| 8:11.264 | ונבון | , |
| 8:11.265 | ובמכתבו |  |
| 8:11.266 | כתב |  |
| 8:11.267 | שאין |  |
| 8:11.268 | למנוע |  |
| 8:11.269 | הלימודים |  |
| 8:11.270 | האלה |  |
| 8:11.271 | מנערי |  |
| 8:11.272 | בני |  |
| 8:11.273 | ישראל | . |
| 8:11.274 | ובעל |  |
| 8:11.275 | הלשון |  |
| 8:11.276 | הרחיב |  |
| 8:11.277 | עוד |  |
| 8:11.278 | את |  |
| 8:11.279 | פיו |  |
| 8:11.280 | לפני |  |
| 8:11.281 | אחד |  |
| 8:11.282 | השופטים |  |
| 8:11.283 | של |  |
| 8:11.284 | האימאם |  |
| 8:11.285 | בפנינו |  |
| 8:11.286 | ג"כ |  |
| 8:11.287 | להלשין |  |
| 8:11.288 | עלינו | . |
| 8:11.289 | כי |  |
| 8:11.290 | אנחנו |  |
| 8:11.291 | עוזבי |  |
| 8:11.292 | דת |  |
| 8:11.293 | הישראלית | , |
| 8:11.294 | ולומדים |  |
| 8:11.295 | בספרי |  |
| 8:11.296 | פילסופי |  |
| 8:11.297 | יון | , |
| 8:11.298 | ומכחישים |  |
| 8:11.299 | אמונת |  |
| 8:11.300 | מציאות |  |
| 8:11.301 | השדים | . |
| 8:11.302 | ובאמת |  |
| 8:11.303 | לא |  |
| 8:11.304 | נמצא |  |
| 8:11.305 | אצלינו |  |
| 8:11.306 | שום |  |
| 8:11.307 | ספר |  |
| 8:11.308 | מספרי |  |
| 8:11.309 | יון | , |
| 8:11.310 | ולא |  |
| 8:11.311 | זולתם |  |
| 8:11.312 | מספרי |  |
| 8:11.313 | הפילסופים |  |
| 8:11.314 | החיצונים |  |
| 8:11.315 | כלל | ! |
| 8:11.316 | ורק |  |
| 8:11.317 | זדון |  |
| 8:11.318 | לבו |  |
| 8:11.319 | השיאו |  |
| 8:11.320 | להלשין |  |
| 8:11.321 | עלינו |  |
| 8:11.322 | שקר | , |
| 8:11.323 | והרבה |  |
| 8:11.324 | שנאת |  |
| 8:11.325 | חנם |  |
| 8:11.326 | במחנה |  |
| 8:11.327 | העברים |  |
| 8:11.328 | על |  |
| 8:11.329 | לא |  |
| 8:11.330 | דבר | ! |
| 8:11.331 | ויוסף |  |
| 8:11.332 | עוד |  |
| 8:11.333 | לחטוא |  |
| 8:11.334 | לשלח |  |
| 8:11.335 | מכתבים |  |
| 8:11.336 | אל |  |
| 8:11.337 | כל |  |
| 8:11.338 | כפרי |  |
| 8:11.339 | ועיירות |  |
| 8:11.340 | התימן |  |
| 8:11.341 | להוציא |  |
| 8:11.342 | ולפרסם |  |
| 8:11.343 | את |  |
| 8:11.344 | הרבה |  |
| 8:11.345 | הרעה |  |
| 8:11.346 | הזאת |  |
| 8:11.347 | עלינו |  |
| 8:11.348 | בשם |  |
| 8:11.349 | מינים | ! |
| 8:11.350 | וערום |  |
| 8:11.351 | יערים |  |
| 8:11.352 | לעשות |  |
| 8:11.353 | נוסחי |  |
| 8:11.354 | דבות |  |
| 8:11.355 | רעות |  |
| 8:11.356 | לאמר | , |
| 8:11.357 | הן |  |
| 8:11.358 | כבר |  |
| 8:11.359 | עשינו |  |
| 8:11.360 | משפט |  |
| 8:11.361 | ברשעים |  |
| 8:11.362 | האלה |  |
| 8:11.363 | ונביא |  |
| 8:11.364 | אותם |  |
| 8:11.365 | אל |  |
| 8:11.366 | בית |  |
| 8:11.367 | הכלא | , |
| 8:11.368 | ועליכם |  |
| 8:11.369 | המצוה |  |
| 8:11.370 | למהר |  |
| 8:11.371 | במכתבי |  |
| 8:11.372 | קובלנא |  |
| 8:11.373 | אל |  |
| 8:11.374 | האימאם |  |
| 8:11.375 | יר"ה | , |
| 8:11.376 | כי |  |
| 8:11.377 | אלה |  |
| 8:11.378 | מקלקלי |  |
| 8:11.379 | דתינו |  |
| 8:11.380 | ומפסידי |  |
| 8:11.381 | אמונתינו | , |
| 8:11.382 | מהרו | , |
| 8:11.383 | חושו | , |
| 8:11.384 | חושו | , |
| 8:11.385 | להגדיל |  |
| 8:11.386 | המדורה |  |
| 8:11.387 | עליהם | ! |
| 8:11.388 | ולא |  |
| 8:11.389 | היה |  |
| 8:11.390 | כותבם |  |
| 8:11.391 | בכתב |  |
| 8:11.392 | ידו | , |
| 8:11.393 | רק |  |
| 8:11.394 | מוסרם |  |
| 8:11.395 | לילדים |  |
| 8:11.396 | מקרוביו |  |
| 8:11.397 | כדי |  |
| 8:11.398 | להעתיקם |  |
| 8:11.399 | בכתב |  |
| 8:11.400 | ידם | , |
| 8:11.401 | ושולחם |  |
| 8:11.402 | לכל |  |
| 8:11.403 | ערי |  |
| 8:11.404 | התימן |  |
| 8:11.405 | שלא |  |
| 8:11.406 | בכתב |  |
| 8:11.407 | ידו | , |
| 8:11.408 | כדי |  |
| 8:11.409 | שיאמר |  |
| 8:11.410 | לא |  |
| 8:11.411 | פעלתי |  |
| 8:11.412 | און | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

