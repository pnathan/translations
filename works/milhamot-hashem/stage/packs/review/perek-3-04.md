# Review pack: milhamot-hashem perek-3-04 (Milchamot ha-Shem 3:7–3:8)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-04`.

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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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
- וז"ל followed by a name: "and these are the words of X:"; the named speaker
  stands outside the quotation.
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
- The words of a verse quoted in the plural (Job 32:16, עמדו): plural, no
  subject added.
- כ' ע"ש: "he writes, see there:"; the quotation opens after it; ע"ש as על שם in
  `readings`.
- bat kol: "bat kol", "daughter of a voice" in the note.
- The Hebrew for "Machzor" follows the scheme (ח = ch): "Machzor Vitry".
- רשע בן רשע: "Wicked one, son of a wicked one"; רוע לב: "wickedness of heart".

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


## Milchamot ha-Shem 3:7

Hebrew (Wikisource): ז) וז"ל המקובל בעל היכל הברכה, מצות עשה על הב"ד לדרוש ולחקור את העדים בשבע חקירות וכו' עוד יש בדיקות שהוא שאלות מה היה לבוש, עפר הארץ שהיה עומד עליה שחור או חול, או חמר, ואם כחשו זה לזה בטל העדות. משרשי המצוה הזאת הכלה הכלולה עולימתא שפירתא דלית בה עינין, הנהגת העולם הזה שיהא בחירה חפשית עיניו כיונים אונאה לבני נשא. ועינין דילה בהיכל זכות, עיני העדה, שדנין בידיעה על פי שנים עדים, תרין עטרין, והדעת הוא סוד עינים כנודע, ואור זה מתפשט בשבע בחינות וכ' ע"ש. ובמצות לא תלך רכיל כתב וז"ל, שלא להגיד לחברו, שכנגדו דבר עליו רע, שנאמר, לא תלך רכיל. ומדיני המצוה אף שמדבר אמת, ויש עון גדול בזה. וגם בכלל לאו זה שלא ידבר כגנות חברו אף אמת והוא כאילו כפר בעיקר וכו'. והאריך עוד בעונש המספר לשון הרע ובשכר הנשמר ממנו, יעו"ש.

| id | word | punctuation after |
|---|---|---|
| 3:7.1 | ז | ) |
| 3:7.2 | וז"ל |  |
| 3:7.3 | המקובל |  |
| 3:7.4 | בעל |  |
| 3:7.5 | היכל |  |
| 3:7.6 | הברכה | , |
| 3:7.7 | מצות |  |
| 3:7.8 | עשה |  |
| 3:7.9 | על |  |
| 3:7.10 | הב"ד |  |
| 3:7.11 | לדרוש |  |
| 3:7.12 | ולחקור |  |
| 3:7.13 | את |  |
| 3:7.14 | העדים |  |
| 3:7.15 | בשבע |  |
| 3:7.16 | חקירות |  |
| 3:7.17 | וכו | ' |
| 3:7.18 | עוד |  |
| 3:7.19 | יש |  |
| 3:7.20 | בדיקות |  |
| 3:7.21 | שהוא |  |
| 3:7.22 | שאלות |  |
| 3:7.23 | מה |  |
| 3:7.24 | היה |  |
| 3:7.25 | לבוש | , |
| 3:7.26 | עפר |  |
| 3:7.27 | הארץ |  |
| 3:7.28 | שהיה |  |
| 3:7.29 | עומד |  |
| 3:7.30 | עליה |  |
| 3:7.31 | שחור |  |
| 3:7.32 | או |  |
| 3:7.33 | חול | , |
| 3:7.34 | או |  |
| 3:7.35 | חמר | , |
| 3:7.36 | ואם |  |
| 3:7.37 | כחשו |  |
| 3:7.38 | זה |  |
| 3:7.39 | לזה |  |
| 3:7.40 | בטל |  |
| 3:7.41 | העדות | . |
| 3:7.42 | משרשי |  |
| 3:7.43 | המצוה |  |
| 3:7.44 | הזאת |  |
| 3:7.45 | הכלה |  |
| 3:7.46 | הכלולה |  |
| 3:7.47 | עולימתא |  |
| 3:7.48 | שפירתא |  |
| 3:7.49 | דלית |  |
| 3:7.50 | בה |  |
| 3:7.51 | עינין | , |
| 3:7.52 | הנהגת |  |
| 3:7.53 | העולם |  |
| 3:7.54 | הזה |  |
| 3:7.55 | שיהא |  |
| 3:7.56 | בחירה |  |
| 3:7.57 | חפשית |  |
| 3:7.58 | עיניו |  |
| 3:7.59 | כיונים |  |
| 3:7.60 | אונאה |  |
| 3:7.61 | לבני |  |
| 3:7.62 | נשא | . |
| 3:7.63 | ועינין |  |
| 3:7.64 | דילה |  |
| 3:7.65 | בהיכל |  |
| 3:7.66 | זכות | , |
| 3:7.67 | עיני |  |
| 3:7.68 | העדה | , |
| 3:7.69 | שדנין |  |
| 3:7.70 | בידיעה |  |
| 3:7.71 | על |  |
| 3:7.72 | פי |  |
| 3:7.73 | שנים |  |
| 3:7.74 | עדים | , |
| 3:7.75 | תרין |  |
| 3:7.76 | עטרין | , |
| 3:7.77 | והדעת |  |
| 3:7.78 | הוא |  |
| 3:7.79 | סוד |  |
| 3:7.80 | עינים |  |
| 3:7.81 | כנודע | , |
| 3:7.82 | ואור |  |
| 3:7.83 | זה |  |
| 3:7.84 | מתפשט |  |
| 3:7.85 | בשבע |  |
| 3:7.86 | בחינות |  |
| 3:7.87 | וכ | ' |
| 3:7.88 | ע"ש | . |
| 3:7.89 | ובמצות |  |
| 3:7.90 | לא |  |
| 3:7.91 | תלך |  |
| 3:7.92 | רכיל |  |
| 3:7.93 | כתב |  |
| 3:7.94 | וז"ל | , |
| 3:7.95 | שלא |  |
| 3:7.96 | להגיד |  |
| 3:7.97 | לחברו | , |
| 3:7.98 | שכנגדו |  |
| 3:7.99 | דבר |  |
| 3:7.100 | עליו |  |
| 3:7.101 | רע | , |
| 3:7.102 | שנאמר | , |
| 3:7.103 | לא |  |
| 3:7.104 | תלך |  |
| 3:7.105 | רכיל | . |
| 3:7.106 | ומדיני |  |
| 3:7.107 | המצוה |  |
| 3:7.108 | אף |  |
| 3:7.109 | שמדבר |  |
| 3:7.110 | אמת | , |
| 3:7.111 | ויש |  |
| 3:7.112 | עון |  |
| 3:7.113 | גדול |  |
| 3:7.114 | בזה | . |
| 3:7.115 | וגם |  |
| 3:7.116 | בכלל |  |
| 3:7.117 | לאו |  |
| 3:7.118 | זה |  |
| 3:7.119 | שלא |  |
| 3:7.120 | ידבר |  |
| 3:7.121 | כגנות |  |
| 3:7.122 | חברו |  |
| 3:7.123 | אף |  |
| 3:7.124 | אמת |  |
| 3:7.125 | והוא |  |
| 3:7.126 | כאילו |  |
| 3:7.127 | כפר |  |
| 3:7.128 | בעיקר |  |
| 3:7.129 | וכו | '. |
| 3:7.130 | והאריך |  |
| 3:7.131 | עוד |  |
| 3:7.132 | בעונש |  |
| 3:7.133 | המספר |  |
| 3:7.134 | לשון |  |
| 3:7.135 | הרע |  |
| 3:7.136 | ובשכר |  |
| 3:7.137 | הנשמר |  |
| 3:7.138 | ממנו | , |
| 3:7.139 | יעו"ש | . |


### Translation under review (3:7)

English: §7. And these are his words: "The Kabbalist, Master of the Palace of Blessing: a positive commandment on the court to inquire into and to examine the witnesses with seven inquiries, etc. There are also examinations, which are questions: what clothing he wore; the dust of the ground on which he stood, whether black, or sand, or clay; and if they contradict one another, the testimony is void. From the roots of this commandment: the included bride, the beautiful young maiden who has no eyes, the governance of this world, that there be free choice, whose eyes are like doves, a deception to the sons of man. Her eyes in the Palace of Merit: the eyes of the congregation, who judge with knowledge on the testimony of two witnesses, two crowns; and the Da'at is the secret of the eyes, as is known; and this light spreads in seven aspects." And he wrote, see there. And concerning the commandment "You shall not go about as a talebearer," he wrote: " that one should not tell his fellow that his adversary has spoken evil about him, as it is said, 'You shall not go about as a talebearer.' And the laws of the commandment: even one who speaks the truth, there is a great sin in this. And also within the scope of this prohibition: that one should not speak in disparagement of his fellow, even if true, and this is as if he denied the fundamental principle, etc." And he expanded further on the punishment of the one who tells slander and on the reward of the one who guards from it; see there.

Literal: section-7 and-these-are-his-words the-Kabbalist master-of the-palace of-blessing a-commandment positive incumbent-on the-court-(beit-din) to-inquire and-to-examine - the-witnesses with-seven inquiries etc. also there-are examinations which-are questions what was clothing the-dust of-the-ground that-was standing upon-it black or sand or clay and-if they-contradicted one the-other is-void the-testimony from-the-roots-of the-commandment this the-bride the-included the-young-maiden beautiful who-has-not in-her eyes the-governance-of the-world this that-there-be free-choice free its-eyes like-doves a-deception for-the-sons-of man and-her-eyes of-hers in-the-palace of-merit the-eyes-of the-congregation who-judge with-knowledge on the-testimony-of two witnesses two crowns and-the-Da'at is the-secret of-the-eyes as-is-known and-the-light this spreads in-seven aspects and-he-wrote see-there and-concerning not go-about as-a-talebearer he-wrote these-are-his-words that-one-should-not tell his-fellow his-adversary spoke about-him evil as-it-is-said not go-about as-a-talebearer and-the-laws-of the-commandment even one-who-speaks the-truth and-there-is a-sin great in-this and-also within-the-scope-of the-prohibition this that-one-should-not speak in-disparagement-of his-fellow even true and-this as-if he-denied the-fundamental-principle etc. and-he-expanded further on-the-punishment-of the-one-who-tells slander evil and-on-the-reward-of the-one-who-guards-himself from-it see-there

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:7.1 | section 7 |  |  | The printed paragraph number ז) is 7; rendered as the section number, as in the decisions. |
| 3:7.2 | and these are his words |  |  | Opening formula of the first quotation; it stands outside the opening quotation mark. |
| 3:7.3 | the Kabbalist |  |  |  |
| 3:7.4 | master of |  |  |  |
| 3:7.5 | the palace |  |  |  |
| 3:7.6 | of blessing |  |  |  |
| 3:7.7 | a commandment |  |  |  |
| 3:7.8 | positive |  |  |  |
| 3:7.9 | incumbent on |  |  |  |
| 3:7.10 | the court (beit din) |  |  | The abbreviation הב"ד is expanded: בית דין, the court. |
| 3:7.11 | to inquire |  |  |  |
| 3:7.12 | and to examine |  |  |  |
| 3:7.13 | - |  |  |  |
| 3:7.14 | the witnesses |  |  |  |
| 3:7.15 | with seven |  |  |  |
| 3:7.16 | inquiries |  |  |  |
| 3:7.17 | etc. |  |  |  |
| 3:7.18 | also |  |  |  |
| 3:7.19 | there are |  |  |  |
| 3:7.20 | examinations |  |  |  |
| 3:7.21 | which are |  |  |  |
| 3:7.22 | questions |  |  |  |
| 3:7.23 | what |  |  |  |
| 3:7.24 | was |  |  |  |
| 3:7.25 | clothing |  |  |  |
| 3:7.26 | the dust |  |  |  |
| 3:7.27 | of the ground |  |  |  |
| 3:7.28 | that was |  |  |  |
| 3:7.29 | standing |  |  |  |
| 3:7.30 | upon it |  |  |  |
| 3:7.31 | black |  |  |  |
| 3:7.32 | or |  |  |  |
| 3:7.33 | sand |  |  |  |
| 3:7.34 | or |  |  |  |
| 3:7.35 | clay |  |  |  |
| 3:7.36 | and if |  |  |  |
| 3:7.37 | they contradicted |  |  |  |
| 3:7.38 | one |  |  |  |
| 3:7.39 | the other |  |  |  |
| 3:7.40 | is void |  |  | בטל: is annulled, is void. |
| 3:7.41 | the testimony |  |  |  |
| 3:7.42 | from the roots of |  |  |  |
| 3:7.43 | the commandment |  |  |  |
| 3:7.44 | this |  |  |  |
| 3:7.45 | the bride |  |  |  |
| 3:7.46 | the included |  |  |  |
| 3:7.47 | the young maiden |  |  | Aramaic words; rendered by sense. |
| 3:7.48 | beautiful |  |  | Aramaic: 'beautiful'. |
| 3:7.49 | who has not |  |  | Aramaic: 'who has not' (דלית, 'there is not in her'). |
| 3:7.50 | in her |  |  |  |
| 3:7.51 | eyes |  |  |  |
| 3:7.52 | the governance of |  |  |  |
| 3:7.53 | the world |  |  |  |
| 3:7.54 | this |  |  |  |
| 3:7.55 | that there be |  |  |  |
| 3:7.56 | free choice |  |  |  |
| 3:7.57 | free |  |  |  |
| 3:7.58 | its eyes |  |  |  |
| 3:7.59 | like doves |  |  |  |
| 3:7.60 | a deception |  |  |  |
| 3:7.61 | for the sons of |  |  |  |
| 3:7.62 | man |  |  |  |
| 3:7.63 | and her eyes |  |  |  |
| 3:7.64 | of hers |  |  |  |
| 3:7.65 | in the palace |  |  |  |
| 3:7.66 | of merit |  |  |  |
| 3:7.67 | the eyes of |  |  |  |
| 3:7.68 | the congregation |  |  |  |
| 3:7.69 | who judge |  |  |  |
| 3:7.70 | with knowledge |  |  |  |
| 3:7.71 | on |  |  |  |
| 3:7.72 | the testimony of |  |  |  |
| 3:7.73 | two |  |  |  |
| 3:7.74 | witnesses |  |  |  |
| 3:7.75 | two |  |  |  |
| 3:7.76 | crowns |  |  |  |
| 3:7.77 | and the Da'at |  |  | Da'at: a Kabbalistic term, transliterated. |
| 3:7.78 | is |  |  |  |
| 3:7.79 | the secret |  |  |  |
| 3:7.80 | of the eyes |  |  |  |
| 3:7.81 | as is known |  |  |  |
| 3:7.82 | and the light |  |  |  |
| 3:7.83 | this |  |  |  |
| 3:7.84 | spreads |  |  |  |
| 3:7.85 | in seven |  |  |  |
| 3:7.86 | aspects |  |  |  |
| 3:7.87 | and he wrote |  |  |  |
| 3:7.88 | see there |  |  | עיין שם: 'see there' (the ע"ש of the author). |
| 3:7.89 | and concerning |  |  |  |
| 3:7.90 | not |  |  |  |
| 3:7.91 | go about |  |  |  |
| 3:7.92 | as a talebearer |  |  |  |
| 3:7.93 | he wrote |  |  | כתב: 'he wrote'; opens the second quotation. |
| 3:7.94 | these are his words |  |  |  |
| 3:7.95 | that one should not |  |  |  |
| 3:7.96 | tell |  |  |  |
| 3:7.97 | his fellow |  |  |  |
| 3:7.98 | his adversary |  |  | שכנגדו: 'his adversary', the one who speaks against him. |
| 3:7.99 | spoke |  |  |  |
| 3:7.100 | about him |  |  |  |
| 3:7.101 | evil |  |  |  |
| 3:7.102 | as it is said |  |  | שנאמר: formula introducing Lev 19:16. |
| 3:7.103 | not |  |  |  |
| 3:7.104 | go about |  |  |  |
| 3:7.105 | as a talebearer |  |  |  |
| 3:7.106 | and the laws of |  |  |  |
| 3:7.107 | the commandment |  |  |  |
| 3:7.108 | even |  |  |  |
| 3:7.109 | one who speaks |  |  |  |
| 3:7.110 | the truth |  |  |  |
| 3:7.111 | and there is |  |  |  |
| 3:7.112 | a sin |  |  |  |
| 3:7.113 | great |  |  |  |
| 3:7.114 | in this |  |  |  |
| 3:7.115 | and also |  |  |  |
| 3:7.116 | within the scope of |  |  |  |
| 3:7.117 | the prohibition |  |  | לאו: a prohibition (a negative commandment). |
| 3:7.118 | this |  |  |  |
| 3:7.119 | that one should not |  |  |  |
| 3:7.120 | speak |  |  |  |
| 3:7.121 | in disparagement of |  |  |  |
| 3:7.122 | his fellow |  |  |  |
| 3:7.123 | even |  |  |  |
| 3:7.124 | true |  |  |  |
| 3:7.125 | and this |  |  |  |
| 3:7.126 | as if |  |  |  |
| 3:7.127 | he denied |  |  |  |
| 3:7.128 | the fundamental principle |  |  |  |
| 3:7.129 | etc. |  |  | וכו' marks a truncated quotation and is kept as 'etc.'. |
| 3:7.130 | and he expanded |  |  |  |
| 3:7.131 | further |  |  |  |
| 3:7.132 | on the punishment of |  |  |  |
| 3:7.133 | the one who tells |  |  |  |
| 3:7.134 | slander |  |  |  |
| 3:7.135 | evil |  |  | לשון הרע is a fixed term: slander. |
| 3:7.136 | and on the reward of |  |  |  |
| 3:7.137 | the one who guards himself |  |  |  |
| 3:7.138 | from it |  |  |  |
| 3:7.139 | see there |  |  | יעו"ש = יעיין שם: 'see there'. |

Phrases:
- 3:7.1 | §7. | section-7
- 3:7.2 | And these are his words: | and-these-are-his-words
- 3:7.3 3:7.4 3:7.5 3:7.6 3:7.7 3:7.8 3:7.9 3:7.10 3:7.11 3:7.12 3:7.13 3:7.14 3:7.15 3:7.16 3:7.17 | "The Kabbalist, Master of the Palace of Blessing: a positive commandment on the court to inquire into and to examine the witnesses with seven inquiries, etc. | the-Kabbalist master-of the-palace of-blessing a-commandment positive incumbent-on the-court-(beit-din) to-inquire and-to-examine - the-witnesses with-seven inquiries etc.
- 3:7.18 3:7.19 3:7.20 3:7.21 3:7.22 3:7.23 3:7.24 3:7.25 | There are also examinations, which are questions: what clothing he wore; | also there-are examinations which-are questions what was clothing
- 3:7.26 3:7.27 3:7.28 3:7.29 3:7.30 3:7.31 3:7.32 3:7.33 3:7.34 3:7.35 | the dust of the ground on which he stood, whether black, or sand, or clay; | the-dust of-the-ground that-was standing upon-it black or sand or clay
- 3:7.36 3:7.37 3:7.38 3:7.39 3:7.40 3:7.41 | and if they contradict one another, the testimony is void. | and-if they-contradicted one the-other is-void the-testimony
- 3:7.42 3:7.43 3:7.44 3:7.45 3:7.46 3:7.47 3:7.48 3:7.49 3:7.50 3:7.51 | From the roots of this commandment: the included bride, the beautiful young maiden who has no eyes, | from-the-roots-of the-commandment this the-bride the-included the-young-maiden beautiful who-has-not in-her eyes
- 3:7.52 3:7.53 3:7.54 3:7.55 3:7.56 3:7.57 3:7.58 3:7.59 3:7.60 3:7.61 3:7.62 | the governance of this world, that there be free choice, whose eyes are like doves, a deception to the sons of man. | the-governance-of the-world this that-there-be free-choice free its-eyes like-doves a-deception for-the-sons-of man
- 3:7.63 3:7.64 3:7.65 3:7.66 3:7.67 3:7.68 | Her eyes in the Palace of Merit: the eyes of the congregation, | and-her-eyes of-hers in-the-palace of-merit the-eyes-of the-congregation
- 3:7.69 3:7.70 3:7.71 3:7.72 3:7.73 3:7.74 | who judge with knowledge on the testimony of two witnesses, | who-judge with-knowledge on the-testimony-of two witnesses
- 3:7.75 3:7.76 | two crowns; | two crowns
- 3:7.77 3:7.78 3:7.79 3:7.80 3:7.81 | and the Da'at is the secret of the eyes, as is known; | and-the-Da'at is the-secret of-the-eyes as-is-known
- 3:7.82 3:7.83 3:7.84 3:7.85 3:7.86 | and this light spreads in seven aspects." | and-the-light this spreads in-seven aspects
- 3:7.87 3:7.88 | And he wrote, see there. | and-he-wrote see-there
- 3:7.89 3:7.90 3:7.91 3:7.92 3:7.93 3:7.94 | And concerning the commandment "You shall not go about as a talebearer," he wrote: " | and-concerning not go-about as-a-talebearer he-wrote these-are-his-words
- 3:7.95 3:7.96 3:7.97 3:7.98 3:7.99 3:7.100 3:7.101 | that one should not tell his fellow that his adversary has spoken evil about him, | that-one-should-not tell his-fellow his-adversary spoke about-him evil
- 3:7.102 3:7.103 3:7.104 3:7.105 | as it is said, 'You shall not go about as a talebearer.' | as-it-is-said not go-about as-a-talebearer
- 3:7.106 3:7.107 3:7.108 3:7.109 3:7.110 | And the laws of the commandment: even one who speaks the truth, | and-the-laws-of the-commandment even one-who-speaks the-truth
- 3:7.111 3:7.112 3:7.113 3:7.114 | there is a great sin in this. | and-there-is a-sin great in-this
- 3:7.115 3:7.116 3:7.117 3:7.118 | And also within the scope of this prohibition: | and-also within-the-scope-of the-prohibition this
- 3:7.119 3:7.120 3:7.121 3:7.122 3:7.123 3:7.124 | that one should not speak in disparagement of his fellow, even if true, | that-one-should-not speak in-disparagement-of his-fellow even true
- 3:7.125 3:7.126 3:7.127 3:7.128 3:7.129 | and this is as if he denied the fundamental principle, etc." | and-this as-if he-denied the-fundamental-principle etc.
- 3:7.130 3:7.131 3:7.132 3:7.133 3:7.134 3:7.135 3:7.136 3:7.137 3:7.138 | And he expanded further on the punishment of the one who tells slander and on the reward of the one who guards from it; | and-he-expanded further on-the-punishment-of the-one-who-tells slander evil and-on-the-reward-of the-one-who-guards-himself from-it
- 3:7.139 | see there. | see-there

Grammar notes:
- The paragraph is two quotations. The first runs from וז"ל (3:7.2) to the author's own formula וכ' ע"ש; the words עולימתא שפירתא דלית בה עינין are Aramaic and are rendered by sense. The source is not named in the text.
- מצות עשה על הב"ד לדרוש ולחקור: 'a positive commandment on the court to inquire and examine'. על with a noun reads 'incumbent upon'; the infinitives give the content of the commandment.
- ואם כחשו זה לזה בטל העדות: a conditional with a perfect verb; the result clause is בטל, 'is void'. The two witnesses are the subject of the conflict (כחשו).
- The second quotation begins at כתב וז"ל (3:7.93-3:7.94) and ends at וכ' (וכו'); the closing words והאריך עוד are the author's own report and take no quotation marks.

Translator's note: The first quotation closes after 'seven aspects'; the author's ע"ש stands outside it. The second quotation has a Bible verse inside it, in single marks.

Commentary rulings:
- apparatus: Wikisource transcription. The 1931 scan is not yet collated; Wikisource additions are translated as base text (decisions, Q2).
- apparatus: Lev 19:16: 'לא תלך רכיל' (3:7.90-3:7.92 and 3:7.102-3:7.105). The quotation stops before בעמיך, so only the words on the page are translated.

## Milchamot ha-Shem 3:8

Hebrew (Wikisource): ח) ובאזהרת לא תעמוד על דם רעך, כ' אפי' עמידה כל דהו. לא תראה בצערו של חברך שאחד מלבין פניו ואתה שותק, עמדו ולא ענו עוד, אלא תהיה מוחה בכל יכלתך במי שמבזה חברך, או מדבר עליו רכילות ולשון הרע, אתה לא תהיה שומע ושותק, אלא תאמר לו רשע עוכר ישראל, למה תדבר סרה על דמות יוצרנו ועל שכינת עוזינו שהיא היו"ד שבעמיך, ומסתלקת בעון זה. ולכך נסמכה זה לזה, לא תלך רכיל ותאמר כיון שאיני מדבר אשמע ואשתוק ונהנה במה שחברו רשע מדבר, גם זה השומע רשע גמור. ועובר על לא תעמוד על דם רעך. ואין לו חלק באלהי ישראל כששומע רשע מדבר סרה על צדיק, או לוחש מלחש בשקרים על חברו רק תכף תדחה אותו ותאמר לו רשע לך מעלי, איני שומע לך, כי אם הוא שומע ושותק, וכל שכן אם נהנה, סופו נידון באש לא נופח, ולא יזכה לעמוד במחיצת הצדיקים, כשומע, כעונה, כמדבר, כולם יספו מן העולם בזה ובבא עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:8.1 | ח | ) |
| 3:8.2 | ובאזהרת |  |
| 3:8.3 | לא |  |
| 3:8.4 | תעמוד |  |
| 3:8.5 | על |  |
| 3:8.6 | דם |  |
| 3:8.7 | רעך | , |
| 3:8.8 | כ | ' |
| 3:8.9 | אפי | ' |
| 3:8.10 | עמידה |  |
| 3:8.11 | כל |  |
| 3:8.12 | דהו | . |
| 3:8.13 | לא |  |
| 3:8.14 | תראה |  |
| 3:8.15 | בצערו |  |
| 3:8.16 | של |  |
| 3:8.17 | חברך |  |
| 3:8.18 | שאחד |  |
| 3:8.19 | מלבין |  |
| 3:8.20 | פניו |  |
| 3:8.21 | ואתה |  |
| 3:8.22 | שותק | , |
| 3:8.23 | עמדו |  |
| 3:8.24 | ולא |  |
| 3:8.25 | ענו |  |
| 3:8.26 | עוד | , |
| 3:8.27 | אלא |  |
| 3:8.28 | תהיה |  |
| 3:8.29 | מוחה |  |
| 3:8.30 | בכל |  |
| 3:8.31 | יכלתך |  |
| 3:8.32 | במי |  |
| 3:8.33 | שמבזה |  |
| 3:8.34 | חברך | , |
| 3:8.35 | או |  |
| 3:8.36 | מדבר |  |
| 3:8.37 | עליו |  |
| 3:8.38 | רכילות |  |
| 3:8.39 | ולשון |  |
| 3:8.40 | הרע | , |
| 3:8.41 | אתה |  |
| 3:8.42 | לא |  |
| 3:8.43 | תהיה |  |
| 3:8.44 | שומע |  |
| 3:8.45 | ושותק | , |
| 3:8.46 | אלא |  |
| 3:8.47 | תאמר |  |
| 3:8.48 | לו |  |
| 3:8.49 | רשע |  |
| 3:8.50 | עוכר |  |
| 3:8.51 | ישראל | , |
| 3:8.52 | למה |  |
| 3:8.53 | תדבר |  |
| 3:8.54 | סרה |  |
| 3:8.55 | על |  |
| 3:8.56 | דמות |  |
| 3:8.57 | יוצרנו |  |
| 3:8.58 | ועל |  |
| 3:8.59 | שכינת |  |
| 3:8.60 | עוזינו |  |
| 3:8.61 | שהיא |  |
| 3:8.62 | היו"ד |  |
| 3:8.63 | שבעמיך | , |
| 3:8.64 | ומסתלקת |  |
| 3:8.65 | בעון |  |
| 3:8.66 | זה | . |
| 3:8.67 | ולכך |  |
| 3:8.68 | נסמכה |  |
| 3:8.69 | זה |  |
| 3:8.70 | לזה | , |
| 3:8.71 | לא |  |
| 3:8.72 | תלך |  |
| 3:8.73 | רכיל |  |
| 3:8.74 | ותאמר |  |
| 3:8.75 | כיון |  |
| 3:8.76 | שאיני |  |
| 3:8.77 | מדבר |  |
| 3:8.78 | אשמע |  |
| 3:8.79 | ואשתוק |  |
| 3:8.80 | ונהנה |  |
| 3:8.81 | במה |  |
| 3:8.82 | שחברו |  |
| 3:8.83 | רשע |  |
| 3:8.84 | מדבר | , |
| 3:8.85 | גם |  |
| 3:8.86 | זה |  |
| 3:8.87 | השומע |  |
| 3:8.88 | רשע |  |
| 3:8.89 | גמור | . |
| 3:8.90 | ועובר |  |
| 3:8.91 | על |  |
| 3:8.92 | לא |  |
| 3:8.93 | תעמוד |  |
| 3:8.94 | על |  |
| 3:8.95 | דם |  |
| 3:8.96 | רעך | . |
| 3:8.97 | ואין |  |
| 3:8.98 | לו |  |
| 3:8.99 | חלק |  |
| 3:8.100 | באלהי |  |
| 3:8.101 | ישראל |  |
| 3:8.102 | כששומע |  |
| 3:8.103 | רשע |  |
| 3:8.104 | מדבר |  |
| 3:8.105 | סרה |  |
| 3:8.106 | על |  |
| 3:8.107 | צדיק | , |
| 3:8.108 | או |  |
| 3:8.109 | לוחש |  |
| 3:8.110 | מלחש |  |
| 3:8.111 | בשקרים |  |
| 3:8.112 | על |  |
| 3:8.113 | חברו |  |
| 3:8.114 | רק |  |
| 3:8.115 | תכף |  |
| 3:8.116 | תדחה |  |
| 3:8.117 | אותו |  |
| 3:8.118 | ותאמר |  |
| 3:8.119 | לו |  |
| 3:8.120 | רשע |  |
| 3:8.121 | לך |  |
| 3:8.122 | מעלי | , |
| 3:8.123 | איני |  |
| 3:8.124 | שומע |  |
| 3:8.125 | לך | , |
| 3:8.126 | כי |  |
| 3:8.127 | אם |  |
| 3:8.128 | הוא |  |
| 3:8.129 | שומע |  |
| 3:8.130 | ושותק | , |
| 3:8.131 | וכל |  |
| 3:8.132 | שכן |  |
| 3:8.133 | אם |  |
| 3:8.134 | נהנה | , |
| 3:8.135 | סופו |  |
| 3:8.136 | נידון |  |
| 3:8.137 | באש |  |
| 3:8.138 | לא |  |
| 3:8.139 | נופח | , |
| 3:8.140 | ולא |  |
| 3:8.141 | יזכה |  |
| 3:8.142 | לעמוד |  |
| 3:8.143 | במחיצת |  |
| 3:8.144 | הצדיקים | , |
| 3:8.145 | כשומע | , |
| 3:8.146 | כעונה | , |
| 3:8.147 | כמדבר | , |
| 3:8.148 | כולם |  |
| 3:8.149 | יספו |  |
| 3:8.150 | מן |  |
| 3:8.151 | העולם |  |
| 3:8.152 | בזה |  |
| 3:8.153 | ובבא |  |
| 3:8.154 | עכ"ל | . |


### Translation under review (3:8)

English: §8. And in the warning, "You shall not stand idly by the blood of your fellow," he wrote: "even standing by, to any degree. You shall not see your fellow's distress, that one puts him to shame, and you are silent; he stood and did not answer again. But you shall protest with all your power against him who disgraces your fellow, or speaks about him with talebearing and evil talk, You shall not be a listener and silent, but you shall say to him: 'Wicked one, troubler of Israel, why do you speak defiance against the likeness of our Maker, and against the Shekhinah, our strength, which is the yod in "among your people," and departs through this sin.' Therefore this was juxtaposed to that, 'You shall not go about as a talebearer,' and one might say: since I do not speak, I will listen and be silent, and take pleasure in what his fellow, the wicked one, speaks. This listener too is an utterly wicked person. and transgresses 'You shall not stand idly by the blood of your fellow.' He has no portion in the God of Israel when he hears a wicked one speak defiance against a righteous person, or whispers a whisper of lies about his fellow; rather, you shall immediately push him away and say to him: 'Wicked one, go away from me, I do not listen to you.' But if he hears and is silent, and all the more so if he takes pleasure, his end is to be judged in unblown fire, and he will not merit to stand in the company of the righteous, as listener, as answerer, as speaker. All of them shall perish from the world in this one, and in the world to come." Thus far his words.

Literal: section-8 and-in-the-warning not stand-idly by the-blood-of your-fellow he-wrote even standing-by any degree not you-see his-distress of your-fellow that-one puts-to-shame his-face and-you are-silent he-stood and-not answer again but you-shall-be protesting with-all your-power against who-disgraces your-fellow or speaks about-him talebearing and-slander evil you not be a-listener and-silent but you-shall-say to-him wicked-one troubler of-Israel why do-you-speak defiance against the-likeness-of our-Maker and-against the-Shekhinah our-strength which-is the-yod in-'among-your-people' and-departs through this-sin therefore was-juxtaposed this to-that not go-about as-a-talebearer and-one-might-say since I-am-not speaking I-will-hear and-be-silent and-take-pleasure in-what his-fellow wicked-one speaks also this listener wicked utterly and-transgresses against not stand by blood-of your-fellow and-he-has-no for-him portion in-the-God-of Israel when-he-hears a-wicked-one speaks defiance against a-righteous-person or whispers a-whisper of-lies about his-fellow but immediately you-push him and-say to-him wicked-one go away-from-me I-do-not listen to-you but if he hears and-is-silent and-all the-more-so if he-takes-pleasure his-end is-judged in-fire not blown and-not he-will-merit to-stand in-the-enclosure of-the-righteous as-listener as-answerer as-speaker all-of-them shall-perish from the-world in-this and-in-the-world-to-come Thus-far-his-words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:8.1 | section 8 |  |  | The printed paragraph number ח) is 8; rendered as the section number. |
| 3:8.2 | and in the warning |  |  |  |
| 3:8.3 | not |  |  |  |
| 3:8.4 | stand idly |  |  |  |
| 3:8.5 | by |  |  |  |
| 3:8.6 | the blood of |  |  |  |
| 3:8.7 | your fellow |  |  |  |
| 3:8.8 | he wrote |  |  | כ' = כתב, 'he wrote'; it opens the quotation that runs to עכ"ל. |
| 3:8.9 | even |  |  |  |
| 3:8.10 | standing by |  |  |  |
| 3:8.11 | any |  |  |  |
| 3:8.12 | degree |  |  |  |
| 3:8.13 | not |  |  |  |
| 3:8.14 | you see |  |  |  |
| 3:8.15 | his distress |  |  |  |
| 3:8.16 | of |  |  |  |
| 3:8.17 | your fellow |  |  |  |
| 3:8.18 | that one |  |  |  |
| 3:8.19 | puts to shame |  |  | Idiom: להלבין פנים, 'to whiten the face', means to shame. |
| 3:8.20 | his face |  |  |  |
| 3:8.21 | and you |  |  |  |
| 3:8.22 | are silent |  |  |  |
| 3:8.23 | he stood |  |  | עמדו is read as 'he stood'; it may also be a noun, 'his standing'. |
| 3:8.24 | and not |  |  |  |
| 3:8.25 | answer |  |  |  |
| 3:8.26 | again |  |  |  |
| 3:8.27 | but |  |  |  |
| 3:8.28 | you shall be |  |  |  |
| 3:8.29 | protesting |  |  |  |
| 3:8.30 | with all |  |  |  |
| 3:8.31 | your power |  |  |  |
| 3:8.32 | against |  |  |  |
| 3:8.33 | who disgraces |  |  |  |
| 3:8.34 | your fellow |  |  |  |
| 3:8.35 | or |  |  |  |
| 3:8.36 | speaks |  |  |  |
| 3:8.37 | about him |  |  |  |
| 3:8.38 | talebearing |  |  |  |
| 3:8.39 | and slander |  |  |  |
| 3:8.40 | evil |  |  |  |
| 3:8.41 | you |  |  |  |
| 3:8.42 | not |  |  |  |
| 3:8.43 | be |  |  |  |
| 3:8.44 | a listener |  |  |  |
| 3:8.45 | and silent |  |  |  |
| 3:8.46 | but |  |  |  |
| 3:8.47 | you shall say |  |  |  |
| 3:8.48 | to him |  |  |  |
| 3:8.49 | wicked one |  |  |  |
| 3:8.50 | troubler |  |  |  |
| 3:8.51 | of Israel |  |  |  |
| 3:8.52 | why |  |  |  |
| 3:8.53 | do you speak |  |  |  |
| 3:8.54 | defiance |  |  |  |
| 3:8.55 | against |  |  |  |
| 3:8.56 | the likeness of |  |  |  |
| 3:8.57 | our Maker |  |  |  |
| 3:8.58 | and against |  |  |  |
| 3:8.59 | the Shekhinah |  |  |  |
| 3:8.60 | our strength |  |  |  |
| 3:8.61 | which is |  |  |  |
| 3:8.62 | the yod |  |  | היו"ד: the letter yod, which the author takes as the Shekhinah. |
| 3:8.63 | in 'among your people' |  |  | The yod is the one in בעמיך of Lev 19:16, which the author cites. |
| 3:8.64 | and departs |  |  |  |
| 3:8.65 | through |  |  |  |
| 3:8.66 | this sin |  |  |  |
| 3:8.67 | therefore |  |  |  |
| 3:8.68 | was juxtaposed |  |  |  |
| 3:8.69 | this |  |  |  |
| 3:8.70 | to that |  |  |  |
| 3:8.71 | not |  |  |  |
| 3:8.72 | go about |  |  |  |
| 3:8.73 | as a talebearer |  |  |  |
| 3:8.74 | and one might say |  |  | ותאמר opens the hypothetical objection, read as a generic 'one might say'. |
| 3:8.75 | since |  |  |  |
| 3:8.76 | I am not |  |  |  |
| 3:8.77 | speaking |  |  |  |
| 3:8.78 | I will hear |  |  |  |
| 3:8.79 | and be silent |  |  |  |
| 3:8.80 | and take pleasure |  |  |  |
| 3:8.81 | in what |  |  |  |
| 3:8.82 | his fellow |  |  |  |
| 3:8.83 | wicked one |  |  |  |
| 3:8.84 | speaks |  |  |  |
| 3:8.85 | also |  |  |  |
| 3:8.86 | this |  |  |  |
| 3:8.87 | listener |  |  |  |
| 3:8.88 | wicked |  |  |  |
| 3:8.89 | utterly |  |  |  |
| 3:8.90 | and transgresses |  |  |  |
| 3:8.91 | against |  |  |  |
| 3:8.92 | not |  |  |  |
| 3:8.93 | stand |  |  |  |
| 3:8.94 | by |  |  |  |
| 3:8.95 | blood of |  |  |  |
| 3:8.96 | your fellow |  |  |  |
| 3:8.97 | and he has no |  |  |  |
| 3:8.98 | for him |  |  |  |
| 3:8.99 | portion |  |  |  |
| 3:8.100 | in the God of |  |  | באלהי ישראל: 'in the God of Israel'. |
| 3:8.101 | Israel |  |  |  |
| 3:8.102 | when he hears |  |  |  |
| 3:8.103 | a wicked one |  |  |  |
| 3:8.104 | speaks |  |  |  |
| 3:8.105 | defiance |  |  |  |
| 3:8.106 | against |  |  |  |
| 3:8.107 | a righteous person |  |  |  |
| 3:8.108 | or |  |  |  |
| 3:8.109 | whispers |  |  |  |
| 3:8.110 | a whisper |  |  |  |
| 3:8.111 | of lies |  |  |  |
| 3:8.112 | about |  |  |  |
| 3:8.113 | his fellow |  |  |  |
| 3:8.114 | but |  |  |  |
| 3:8.115 | immediately |  |  |  |
| 3:8.116 | you push |  |  |  |
| 3:8.117 | him |  |  |  |
| 3:8.118 | and say |  |  |  |
| 3:8.119 | to him |  |  |  |
| 3:8.120 | wicked one |  |  |  |
| 3:8.121 | go |  |  |  |
| 3:8.122 | away from me |  |  |  |
| 3:8.123 | I do not |  |  |  |
| 3:8.124 | listen |  |  |  |
| 3:8.125 | to you |  |  |  |
| 3:8.126 | but |  |  |  |
| 3:8.127 | if |  |  |  |
| 3:8.128 | he |  |  |  |
| 3:8.129 | hears |  |  |  |
| 3:8.130 | and is silent |  |  |  |
| 3:8.131 | and all |  |  | וכל שכן: 'and all the more so', a kal va-chomer. |
| 3:8.132 | the more so |  |  |  |
| 3:8.133 | if |  |  |  |
| 3:8.134 | he takes pleasure |  |  |  |
| 3:8.135 | his end |  |  |  |
| 3:8.136 | is judged |  |  |  |
| 3:8.137 | in fire |  |  |  |
| 3:8.138 | not |  |  |  |
| 3:8.139 | blown |  |  | לא נופח: fire that is not fanned. |
| 3:8.140 | and not |  |  |  |
| 3:8.141 | he will merit |  |  |  |
| 3:8.142 | to stand |  |  |  |
| 3:8.143 | in the enclosure |  |  | מחיצת: the enclosure, the company of the righteous. |
| 3:8.144 | of the righteous |  |  |  |
| 3:8.145 | as listener |  |  |  |
| 3:8.146 | as answerer |  |  |  |
| 3:8.147 | as speaker |  |  |  |
| 3:8.148 | all of them |  |  |  |
| 3:8.149 | shall perish |  |  |  |
| 3:8.150 | from |  |  |  |
| 3:8.151 | the world |  |  |  |
| 3:8.152 | in this |  |  |  |
| 3:8.153 | and in the world to come |  |  | ובבא: 'and in the coming (world)', the world to come. |
| 3:8.154 | Thus far his words |  |  | עכ"ל: closing formula, 'thus far his words', outside the closing quotation mark. |

Phrases:
- 3:8.1 | §8. | section-8
- 3:8.2 3:8.3 3:8.4 3:8.5 3:8.6 3:8.7 | And in the warning, "You shall not stand idly by the blood of your fellow," | and-in-the-warning not stand-idly by the-blood-of your-fellow
- 3:8.8 3:8.9 3:8.10 3:8.11 3:8.12 | he wrote: "even standing by, to any degree. | he-wrote even standing-by any degree
- 3:8.13 3:8.14 3:8.15 3:8.16 3:8.17 3:8.18 3:8.19 3:8.20 | You shall not see your fellow's distress, that one puts him to shame, | not you-see his-distress of your-fellow that-one puts-to-shame his-face
- 3:8.21 3:8.22 | and you are silent; | and-you are-silent
- 3:8.23 3:8.24 3:8.25 3:8.26 | he stood and did not answer again. | he-stood and-not answer again
- 3:8.27 3:8.28 3:8.29 3:8.30 3:8.31 | But you shall protest with all your power | but you-shall-be protesting with-all your-power
- 3:8.32 3:8.33 3:8.34 | against him who disgraces your fellow, | against who-disgraces your-fellow
- 3:8.35 3:8.36 3:8.37 3:8.38 3:8.39 3:8.40 | or speaks about him with talebearing and evil talk, | or speaks about-him talebearing and-slander evil
- 3:8.41 3:8.42 3:8.43 3:8.44 3:8.45 | You shall not be a listener and silent, | you not be a-listener and-silent
- 3:8.46 3:8.47 3:8.48 3:8.49 3:8.50 3:8.51 | but you shall say to him: 'Wicked one, troubler of Israel, | but you-shall-say to-him wicked-one troubler of-Israel
- 3:8.52 3:8.53 3:8.54 3:8.55 3:8.56 3:8.57 | why do you speak defiance against the likeness of our Maker, | why do-you-speak defiance against the-likeness-of our-Maker
- 3:8.58 3:8.59 3:8.60 3:8.61 3:8.62 3:8.63 | and against the Shekhinah, our strength, which is the yod in "among your people," | and-against the-Shekhinah our-strength which-is the-yod in-'among-your-people'
- 3:8.64 3:8.65 3:8.66 | and departs through this sin.' | and-departs through this-sin
- 3:8.67 3:8.68 3:8.69 3:8.70 | Therefore this was juxtaposed to that, | therefore was-juxtaposed this to-that
- 3:8.71 3:8.72 3:8.73 | 'You shall not go about as a talebearer,' | not go-about as-a-talebearer
- 3:8.74 3:8.75 3:8.76 3:8.77 | and one might say: since I do not speak, | and-one-might-say since I-am-not speaking
- 3:8.78 3:8.79 3:8.80 3:8.81 3:8.82 3:8.83 3:8.84 | I will listen and be silent, and take pleasure in what his fellow, the wicked one, speaks. | I-will-hear and-be-silent and-take-pleasure in-what his-fellow wicked-one speaks
- 3:8.85 3:8.86 3:8.87 3:8.88 3:8.89 | This listener too is an utterly wicked person. | also this listener wicked utterly
- 3:8.90 3:8.91 3:8.92 3:8.93 3:8.94 3:8.95 3:8.96 | and transgresses 'You shall not stand idly by the blood of your fellow.' | and-transgresses against not stand by blood-of your-fellow
- 3:8.97 3:8.98 3:8.99 3:8.100 3:8.101 | He has no portion in the God of Israel | and-he-has-no for-him portion in-the-God-of Israel
- 3:8.102 3:8.103 3:8.104 3:8.105 3:8.106 3:8.107 | when he hears a wicked one speak defiance against a righteous person, | when-he-hears a-wicked-one speaks defiance against a-righteous-person
- 3:8.108 3:8.109 3:8.110 3:8.111 3:8.112 3:8.113 | or whispers a whisper of lies about his fellow; | or whispers a-whisper of-lies about his-fellow
- 3:8.114 3:8.115 3:8.116 3:8.117 | rather, you shall immediately push him away | but immediately you-push him
- 3:8.118 3:8.119 3:8.120 3:8.121 3:8.122 | and say to him: 'Wicked one, go away from me, | and-say to-him wicked-one go away-from-me
- 3:8.123 3:8.124 3:8.125 | I do not listen to you.' | I-do-not listen to-you
- 3:8.126 3:8.127 3:8.128 3:8.129 3:8.130 | But if he hears and is silent, | but if he hears and-is-silent
- 3:8.131 3:8.132 3:8.133 3:8.134 | and all the more so if he takes pleasure, | and-all the-more-so if he-takes-pleasure
- 3:8.135 3:8.136 3:8.137 3:8.138 3:8.139 | his end is to be judged in unblown fire, | his-end is-judged in-fire not blown
- 3:8.140 3:8.141 3:8.142 3:8.143 3:8.144 | and he will not merit to stand in the company of the righteous, | and-not he-will-merit to-stand in-the-enclosure of-the-righteous
- 3:8.145 3:8.146 3:8.147 | as listener, as answerer, as speaker. | as-listener as-answerer as-speaker
- 3:8.148 3:8.149 3:8.150 3:8.151 3:8.152 | All of them shall perish from the world in this one, | all-of-them shall-perish from the-world in-this
- 3:8.153 | and in the world to come." | and-in-the-world-to-come
- 3:8.154 | Thus far his words. | Thus-far-his-words

Grammar notes:
- The quotation opens with כ' (he wrote) after the verse לא תעמוד על דם רעך, and closes at עכ"ל. Within it, direct address is in single marks, and the verse לא תלך רכיל is a single-marked quotation inside the outer quotation.
- ולכך נסמכה זה לזה: 'for this reason this was placed next to that'. The verse לא תלך רכיל is the second member of the pair; ותאמר begins the objection that the author answers.
- וכל שכן ('all the more so') introduces the consequence. The three participles כשומע, כעונה, כמדבר ('as listener, as answerer, as speaker') are the ways of taking part, and כולם יספו ('all of them shall perish') refers back to them.

Translator's note: The opening verse is in double marks because it stands outside the quotation; the verse quoted inside the quotation (3:8.90-3:8.96) is in single marks.

Commentary rulings:
- apparatus: Lev 19:16: 'לא תעמוד על דם רעך' (3:8.2-3:8.7 and 3:8.90-3:8.96), quoted as the author gives it.
- apparatus: Lev 19:16, 'בעמיך', is the source of the yod of the Shekhinah (3:8.62-3:8.63); the author's reading, not a translation of the verse's sense.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

