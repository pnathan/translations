# Translation pack: milhamot-hashem perek-7-13 (Milchamot ha-Shem 7:18–7:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-13`.

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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
- A word that does not fit its context but is close to a likelier one (יכינו
  for יבינו / יוכלו): translate by the sense of the context, and give the
  printed word and the likelier one in a token note and `apparatus`.
- Allusions to Isaiah and other prophets: when a phrase echoes a verse
  (אל דבר ה', בטרם תחיל), the verse's wording governs the sense of the
  words; list the verse in `apparatus`.
- Keep every word of a cited title: מחבר הספר הזהר = "the author of the book the
  Zohar"; the fixed "the author of the Zohar" is for מחבר הזהר only.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- An unclear abbreviation (לי"ח): the conjecture in a token note, the Hebrew
  letters kept in `apparatus`; the English carries the conjecture only if it
  is the sole sense that fits.
- בעל + a book title (המקובל בעל היכל הברכה): "the Kabbalist, the author of
  Heikhal ha-Berakhah" (the title translated once in a note: "Palace of
  Blessing"). Identify no author the text does not name.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 7:18

Hebrew (Wikisource): ורז"ל אמרו בסנהדרין כשעלה משה למרום מצאו להקב"ה שיושב וכותב ארך אפים. אמר לפניו רבונו של עולם לצדיקים? אמר לו אף לרשעים. אמר לפניו רשעים יאבדו אמר ליה השתא חזית דמבעי לך. בשעה שחטאו ישראל אמרו לו לא כך אמרת ועתה יגדל נא כח ה' כאשר דברת לאמר ה' ארך אפים וכו'. וכן כל הנביאים מתארים לשי"ת בתאר ארך אפים. דוד המלך ע"ה גם במזמור פ"ו אמר ואתה ה' אל רחום וחנון ארך אפים ורב חסד ואמת, ובמזמור קמ"ה אמר חנון ורחום ה' ארך אפים וגדל חסד. הנביא נחום אמר ה' ארך אפים וגדל כח ונקה לא ינקה. יונה אמר בתפלתו כי ידעתי כי אתה אל חנון ורחום ארך אפים ורב חסד וניחם על הרעה. יואל אמר ושובו של ה' אלהיכם כי חנון ורחום הוא ארך אפים ורב חסד וניחם על הרעה, אנשי כנסת הגדולה אמרו בתפלתם ואתה אלוה סליחות חנון ורחום ארך אפים ורב חסד ולא עזבתם (נחמיה ט׳:י״ז) ושלמה המלך ע"ה אמר טוב ארך רוח מגבה רוח (קהלת ז׳:ח׳), ופירש רש"י על ארך רוח המאריך רוגזו ואינו ממהר לריב. ובמשלי אמר קצר אפים יעשה אולת, ואיש מזמות ישנא, ופירש רש"י קצר אפים הממהר לנקום כעסו, ואיש מזמות. מחשבות עצה רעה ישנא, ובפי' קב ונקי פי' קצר אפים הוא היפך ארך אפים, האדם שהוא נמהר לשפוך חמתו ולנקום נקמתו יעשה כל מעשיו באולת כסילות, ועיין עוד במפרשים ובס' הנז' משבח ארך אפים ואמר ארך אפים רב תבונות, וקצר רוח מרים אולת, ופירשו המפרשים מי שהוא ארך אפים שאינו ממהר לכעוס והוא מושל ברוחו הוא רב תבונות, ומי שהוא קצר אפים וקצר רוח שלא ימשול ברוחו הוא מרים האולת מן הארץ, ומגביהה ומניחה על ראשו, ומגלה אותה לאשר לא ידעה, או יאמר שהוא מעורר האולת. וי"מ מרים אולת שהוא גורם לכסילים להרים בו יד. כלומר למרוד בו ע"כ. ושם סי' ט"ו איש חמה יגרה מדון, וארך אפים ישקיט ריב, ופי' מי שהוא בעל חמה דהיינו קצר אפים הוא יבעיר המדון והקטטה ויגדילנו אע"פ שהריב קטן יעלנו ויגדילנו בחמתו, והאדם המאריך אפו ישקיט ריב ויניחנו. ובסי' י"ו טוב ארך אפים מגבור ומושל ברוחו מלוכד עיר וכ' המפרשים שטוב ארך אפים המושל ביצרו וכובש תאותו מלהנקם, מן הגבור המתנקם באויבו. כי הכובש את יצרו, אויבו שהוא היצר מבפנים, והוא מתגבר עליו וכובשו, וזה הגבור אויבו מבחוץ והוא מתגבר על האויב החיצון, והמושל ברוחו הוא מושל על האויב הפנימי. המורם מכל האמור, כי תואר קצר אפים הוא תואר אולת וכסילות, וכבר נודע הדבר אשר דברו רז"ל בעל חובת הלבבות ור"י הלוי והרמב"ם במורה פרק כ"ו ומ"ו ומ"ז לראשון ורס"ג בהאמונות והדעות מאמר האחדות ושבילי אמונה וזולתם, שלא יתואר השי"ת אלא בתארים נעלים, המורים על השלימות, להורות שהור ית' שלם בכל מיני שלימות. ולא יתואר בשום תואר חסרון או העדר, כל שכן בתאר פחיתות וסכילות וסכלות.

| id | word | punctuation after |
|---|---|---|
| 7:18.1 | ורז"ל |  |
| 7:18.2 | אמרו |  |
| 7:18.3 | בסנהדרין |  |
| 7:18.4 | כשעלה |  |
| 7:18.5 | משה |  |
| 7:18.6 | למרום |  |
| 7:18.7 | מצאו |  |
| 7:18.8 | להקב"ה |  |
| 7:18.9 | שיושב |  |
| 7:18.10 | וכותב |  |
| 7:18.11 | ארך |  |
| 7:18.12 | אפים | . |
| 7:18.13 | אמר |  |
| 7:18.14 | לפניו |  |
| 7:18.15 | רבונו |  |
| 7:18.16 | של |  |
| 7:18.17 | עולם |  |
| 7:18.18 | לצדיקים | ? |
| 7:18.19 | אמר |  |
| 7:18.20 | לו |  |
| 7:18.21 | אף |  |
| 7:18.22 | לרשעים | . |
| 7:18.23 | אמר |  |
| 7:18.24 | לפניו |  |
| 7:18.25 | רשעים |  |
| 7:18.26 | יאבדו |  |
| 7:18.27 | אמר |  |
| 7:18.28 | ליה |  |
| 7:18.29 | השתא |  |
| 7:18.30 | חזית |  |
| 7:18.31 | דמבעי |  |
| 7:18.32 | לך | . |
| 7:18.33 | בשעה |  |
| 7:18.34 | שחטאו |  |
| 7:18.35 | ישראל |  |
| 7:18.36 | אמרו |  |
| 7:18.37 | לו |  |
| 7:18.38 | לא |  |
| 7:18.39 | כך |  |
| 7:18.40 | אמרת |  |
| 7:18.41 | ועתה |  |
| 7:18.42 | יגדל |  |
| 7:18.43 | נא |  |
| 7:18.44 | כח |  |
| 7:18.45 | ה | ' |
| 7:18.46 | כאשר |  |
| 7:18.47 | דברת |  |
| 7:18.48 | לאמר |  |
| 7:18.49 | ה | ' |
| 7:18.50 | ארך |  |
| 7:18.51 | אפים |  |
| 7:18.52 | וכו | '. |
| 7:18.53 | וכן |  |
| 7:18.54 | כל |  |
| 7:18.55 | הנביאים |  |
| 7:18.56 | מתארים |  |
| 7:18.57 | לשי"ת |  |
| 7:18.58 | בתאר |  |
| 7:18.59 | ארך |  |
| 7:18.60 | אפים | . |
| 7:18.61 | דוד |  |
| 7:18.62 | המלך |  |
| 7:18.63 | ע"ה |  |
| 7:18.64 | גם |  |
| 7:18.65 | במזמור |  |
| 7:18.66 | פ"ו |  |
| 7:18.67 | אמר |  |
| 7:18.68 | ואתה |  |
| 7:18.69 | ה | ' |
| 7:18.70 | אל |  |
| 7:18.71 | רחום |  |
| 7:18.72 | וחנון |  |
| 7:18.73 | ארך |  |
| 7:18.74 | אפים |  |
| 7:18.75 | ורב |  |
| 7:18.76 | חסד |  |
| 7:18.77 | ואמת | , |
| 7:18.78 | ובמזמור |  |
| 7:18.79 | קמ"ה |  |
| 7:18.80 | אמר |  |
| 7:18.81 | חנון |  |
| 7:18.82 | ורחום |  |
| 7:18.83 | ה | ' |
| 7:18.84 | ארך |  |
| 7:18.85 | אפים |  |
| 7:18.86 | וגדל |  |
| 7:18.87 | חסד | . |
| 7:18.88 | הנביא |  |
| 7:18.89 | נחום |  |
| 7:18.90 | אמר |  |
| 7:18.91 | ה | ' |
| 7:18.92 | ארך |  |
| 7:18.93 | אפים |  |
| 7:18.94 | וגדל |  |
| 7:18.95 | כח |  |
| 7:18.96 | ונקה |  |
| 7:18.97 | לא |  |
| 7:18.98 | ינקה | . |
| 7:18.99 | יונה |  |
| 7:18.100 | אמר |  |
| 7:18.101 | בתפלתו |  |
| 7:18.102 | כי |  |
| 7:18.103 | ידעתי |  |
| 7:18.104 | כי |  |
| 7:18.105 | אתה |  |
| 7:18.106 | אל |  |
| 7:18.107 | חנון |  |
| 7:18.108 | ורחום |  |
| 7:18.109 | ארך |  |
| 7:18.110 | אפים |  |
| 7:18.111 | ורב |  |
| 7:18.112 | חסד |  |
| 7:18.113 | וניחם |  |
| 7:18.114 | על |  |
| 7:18.115 | הרעה | . |
| 7:18.116 | יואל |  |
| 7:18.117 | אמר |  |
| 7:18.118 | ושובו |  |
| 7:18.119 | של |  |
| 7:18.120 | ה | ' |
| 7:18.121 | אלהיכם |  |
| 7:18.122 | כי |  |
| 7:18.123 | חנון |  |
| 7:18.124 | ורחום |  |
| 7:18.125 | הוא |  |
| 7:18.126 | ארך |  |
| 7:18.127 | אפים |  |
| 7:18.128 | ורב |  |
| 7:18.129 | חסד |  |
| 7:18.130 | וניחם |  |
| 7:18.131 | על |  |
| 7:18.132 | הרעה | , |
| 7:18.133 | אנשי |  |
| 7:18.134 | כנסת |  |
| 7:18.135 | הגדולה |  |
| 7:18.136 | אמרו |  |
| 7:18.137 | בתפלתם |  |
| 7:18.138 | ואתה |  |
| 7:18.139 | אלוה |  |
| 7:18.140 | סליחות |  |
| 7:18.141 | חנון |  |
| 7:18.142 | ורחום |  |
| 7:18.143 | ארך |  |
| 7:18.144 | אפים |  |
| 7:18.145 | ורב |  |
| 7:18.146 | חסד |  |
| 7:18.147 | ולא |  |
| 7:18.148 | עזבתם |  |
| 7:18.149 | נחמיה |  |
| 7:18.150 | ט׳:י״ז | ) |
| 7:18.151 | ושלמה |  |
| 7:18.152 | המלך |  |
| 7:18.153 | ע"ה |  |
| 7:18.154 | אמר |  |
| 7:18.155 | טוב |  |
| 7:18.156 | ארך |  |
| 7:18.157 | רוח |  |
| 7:18.158 | מגבה |  |
| 7:18.159 | רוח |  |
| 7:18.160 | קהלת |  |
| 7:18.161 | ז׳:ח׳ | ), |
| 7:18.162 | ופירש |  |
| 7:18.163 | רש"י |  |
| 7:18.164 | על |  |
| 7:18.165 | ארך |  |
| 7:18.166 | רוח |  |
| 7:18.167 | המאריך |  |
| 7:18.168 | רוגזו |  |
| 7:18.169 | ואינו |  |
| 7:18.170 | ממהר |  |
| 7:18.171 | לריב | . |
| 7:18.172 | ובמשלי |  |
| 7:18.173 | אמר |  |
| 7:18.174 | קצר |  |
| 7:18.175 | אפים |  |
| 7:18.176 | יעשה |  |
| 7:18.177 | אולת | , |
| 7:18.178 | ואיש |  |
| 7:18.179 | מזמות |  |
| 7:18.180 | ישנא | , |
| 7:18.181 | ופירש |  |
| 7:18.182 | רש"י |  |
| 7:18.183 | קצר |  |
| 7:18.184 | אפים |  |
| 7:18.185 | הממהר |  |
| 7:18.186 | לנקום |  |
| 7:18.187 | כעסו | , |
| 7:18.188 | ואיש |  |
| 7:18.189 | מזמות | . |
| 7:18.190 | מחשבות |  |
| 7:18.191 | עצה |  |
| 7:18.192 | רעה |  |
| 7:18.193 | ישנא | , |
| 7:18.194 | ובפי | ' |
| 7:18.195 | קב |  |
| 7:18.196 | ונקי |  |
| 7:18.197 | פי | ' |
| 7:18.198 | קצר |  |
| 7:18.199 | אפים |  |
| 7:18.200 | הוא |  |
| 7:18.201 | היפך |  |
| 7:18.202 | ארך |  |
| 7:18.203 | אפים | , |
| 7:18.204 | האדם |  |
| 7:18.205 | שהוא |  |
| 7:18.206 | נמהר |  |
| 7:18.207 | לשפוך |  |
| 7:18.208 | חמתו |  |
| 7:18.209 | ולנקום |  |
| 7:18.210 | נקמתו |  |
| 7:18.211 | יעשה |  |
| 7:18.212 | כל |  |
| 7:18.213 | מעשיו |  |
| 7:18.214 | באולת |  |
| 7:18.215 | כסילות | , |
| 7:18.216 | ועיין |  |
| 7:18.217 | עוד |  |
| 7:18.218 | במפרשים |  |
| 7:18.219 | ובס | ' |
| 7:18.220 | הנז | ' |
| 7:18.221 | משבח |  |
| 7:18.222 | ארך |  |
| 7:18.223 | אפים |  |
| 7:18.224 | ואמר |  |
| 7:18.225 | ארך |  |
| 7:18.226 | אפים |  |
| 7:18.227 | רב |  |
| 7:18.228 | תבונות | , |
| 7:18.229 | וקצר |  |
| 7:18.230 | רוח |  |
| 7:18.231 | מרים |  |
| 7:18.232 | אולת | , |
| 7:18.233 | ופירשו |  |
| 7:18.234 | המפרשים |  |
| 7:18.235 | מי |  |
| 7:18.236 | שהוא |  |
| 7:18.237 | ארך |  |
| 7:18.238 | אפים |  |
| 7:18.239 | שאינו |  |
| 7:18.240 | ממהר |  |
| 7:18.241 | לכעוס |  |
| 7:18.242 | והוא |  |
| 7:18.243 | מושל |  |
| 7:18.244 | ברוחו |  |
| 7:18.245 | הוא |  |
| 7:18.246 | רב |  |
| 7:18.247 | תבונות | , |
| 7:18.248 | ומי |  |
| 7:18.249 | שהוא |  |
| 7:18.250 | קצר |  |
| 7:18.251 | אפים |  |
| 7:18.252 | וקצר |  |
| 7:18.253 | רוח |  |
| 7:18.254 | שלא |  |
| 7:18.255 | ימשול |  |
| 7:18.256 | ברוחו |  |
| 7:18.257 | הוא |  |
| 7:18.258 | מרים |  |
| 7:18.259 | האולת |  |
| 7:18.260 | מן |  |
| 7:18.261 | הארץ | , |
| 7:18.262 | ומגביהה |  |
| 7:18.263 | ומניחה |  |
| 7:18.264 | על |  |
| 7:18.265 | ראשו | , |
| 7:18.266 | ומגלה |  |
| 7:18.267 | אותה |  |
| 7:18.268 | לאשר |  |
| 7:18.269 | לא |  |
| 7:18.270 | ידעה | , |
| 7:18.271 | או |  |
| 7:18.272 | יאמר |  |
| 7:18.273 | שהוא |  |
| 7:18.274 | מעורר |  |
| 7:18.275 | האולת | . |
| 7:18.276 | וי"מ |  |
| 7:18.277 | מרים |  |
| 7:18.278 | אולת |  |
| 7:18.279 | שהוא |  |
| 7:18.280 | גורם |  |
| 7:18.281 | לכסילים |  |
| 7:18.282 | להרים |  |
| 7:18.283 | בו |  |
| 7:18.284 | יד | . |
| 7:18.285 | כלומר |  |
| 7:18.286 | למרוד |  |
| 7:18.287 | בו |  |
| 7:18.288 | ע"כ | . |
| 7:18.289 | ושם |  |
| 7:18.290 | סי | ' |
| 7:18.291 | ט"ו |  |
| 7:18.292 | איש |  |
| 7:18.293 | חמה |  |
| 7:18.294 | יגרה |  |
| 7:18.295 | מדון | , |
| 7:18.296 | וארך |  |
| 7:18.297 | אפים |  |
| 7:18.298 | ישקיט |  |
| 7:18.299 | ריב | , |
| 7:18.300 | ופי | ' |
| 7:18.301 | מי |  |
| 7:18.302 | שהוא |  |
| 7:18.303 | בעל |  |
| 7:18.304 | חמה |  |
| 7:18.305 | דהיינו |  |
| 7:18.306 | קצר |  |
| 7:18.307 | אפים |  |
| 7:18.308 | הוא |  |
| 7:18.309 | יבעיר |  |
| 7:18.310 | המדון |  |
| 7:18.311 | והקטטה |  |
| 7:18.312 | ויגדילנו |  |
| 7:18.313 | אע"פ |  |
| 7:18.314 | שהריב |  |
| 7:18.315 | קטן |  |
| 7:18.316 | יעלנו |  |
| 7:18.317 | ויגדילנו |  |
| 7:18.318 | בחמתו | , |
| 7:18.319 | והאדם |  |
| 7:18.320 | המאריך |  |
| 7:18.321 | אפו |  |
| 7:18.322 | ישקיט |  |
| 7:18.323 | ריב |  |
| 7:18.324 | ויניחנו | . |
| 7:18.325 | ובסי | ' |
| 7:18.326 | י"ו |  |
| 7:18.327 | טוב |  |
| 7:18.328 | ארך |  |
| 7:18.329 | אפים |  |
| 7:18.330 | מגבור |  |
| 7:18.331 | ומושל |  |
| 7:18.332 | ברוחו |  |
| 7:18.333 | מלוכד |  |
| 7:18.334 | עיר |  |
| 7:18.335 | וכ | ' |
| 7:18.336 | המפרשים |  |
| 7:18.337 | שטוב |  |
| 7:18.338 | ארך |  |
| 7:18.339 | אפים |  |
| 7:18.340 | המושל |  |
| 7:18.341 | ביצרו |  |
| 7:18.342 | וכובש |  |
| 7:18.343 | תאותו |  |
| 7:18.344 | מלהנקם | , |
| 7:18.345 | מן |  |
| 7:18.346 | הגבור |  |
| 7:18.347 | המתנקם |  |
| 7:18.348 | באויבו | . |
| 7:18.349 | כי |  |
| 7:18.350 | הכובש |  |
| 7:18.351 | את |  |
| 7:18.352 | יצרו | , |
| 7:18.353 | אויבו |  |
| 7:18.354 | שהוא |  |
| 7:18.355 | היצר |  |
| 7:18.356 | מבפנים | , |
| 7:18.357 | והוא |  |
| 7:18.358 | מתגבר |  |
| 7:18.359 | עליו |  |
| 7:18.360 | וכובשו | , |
| 7:18.361 | וזה |  |
| 7:18.362 | הגבור |  |
| 7:18.363 | אויבו |  |
| 7:18.364 | מבחוץ |  |
| 7:18.365 | והוא |  |
| 7:18.366 | מתגבר |  |
| 7:18.367 | על |  |
| 7:18.368 | האויב |  |
| 7:18.369 | החיצון | , |
| 7:18.370 | והמושל |  |
| 7:18.371 | ברוחו |  |
| 7:18.372 | הוא |  |
| 7:18.373 | מושל |  |
| 7:18.374 | על |  |
| 7:18.375 | האויב |  |
| 7:18.376 | הפנימי | . |
| 7:18.377 | המורם |  |
| 7:18.378 | מכל |  |
| 7:18.379 | האמור | , |
| 7:18.380 | כי |  |
| 7:18.381 | תואר |  |
| 7:18.382 | קצר |  |
| 7:18.383 | אפים |  |
| 7:18.384 | הוא |  |
| 7:18.385 | תואר |  |
| 7:18.386 | אולת |  |
| 7:18.387 | וכסילות | , |
| 7:18.388 | וכבר |  |
| 7:18.389 | נודע |  |
| 7:18.390 | הדבר |  |
| 7:18.391 | אשר |  |
| 7:18.392 | דברו |  |
| 7:18.393 | רז"ל |  |
| 7:18.394 | בעל |  |
| 7:18.395 | חובת |  |
| 7:18.396 | הלבבות |  |
| 7:18.397 | ור"י |  |
| 7:18.398 | הלוי |  |
| 7:18.399 | והרמב"ם |  |
| 7:18.400 | במורה |  |
| 7:18.401 | פרק |  |
| 7:18.402 | כ"ו |  |
| 7:18.403 | ומ"ו |  |
| 7:18.404 | ומ"ז |  |
| 7:18.405 | לראשון |  |
| 7:18.406 | ורס"ג |  |
| 7:18.407 | בהאמונות |  |
| 7:18.408 | והדעות |  |
| 7:18.409 | מאמר |  |
| 7:18.410 | האחדות |  |
| 7:18.411 | ושבילי |  |
| 7:18.412 | אמונה |  |
| 7:18.413 | וזולתם | , |
| 7:18.414 | שלא |  |
| 7:18.415 | יתואר |  |
| 7:18.416 | השי"ת |  |
| 7:18.417 | אלא |  |
| 7:18.418 | בתארים |  |
| 7:18.419 | נעלים | , |
| 7:18.420 | המורים |  |
| 7:18.421 | על |  |
| 7:18.422 | השלימות | , |
| 7:18.423 | להורות |  |
| 7:18.424 | שהור |  |
| 7:18.425 | ית | ' |
| 7:18.426 | שלם |  |
| 7:18.427 | בכל |  |
| 7:18.428 | מיני |  |
| 7:18.429 | שלימות | . |
| 7:18.430 | ולא |  |
| 7:18.431 | יתואר |  |
| 7:18.432 | בשום |  |
| 7:18.433 | תואר |  |
| 7:18.434 | חסרון |  |
| 7:18.435 | או |  |
| 7:18.436 | העדר | , |
| 7:18.437 | כל |  |
| 7:18.438 | שכן |  |
| 7:18.439 | בתאר |  |
| 7:18.440 | פחיתות |  |
| 7:18.441 | וסכילות |  |
| 7:18.442 | וסכלות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

