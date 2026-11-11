# Translation pack: milhamot-hashem perek-4-03 (Milchamot ha-Shem 4:5–4:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-03`.

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
- **אלהות / אלוהות**: "divinity" for the abstract ("רבוי באלהות" → "a
  multiplicity in divinity"). "divinities" where the word is counted or
  plural ("כמה אלוהות", "קורים לפרצופים אלוהות" → "call the partzufim
  divinities"). Never "gods": that is אלהים אחרים.
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
- ר"ל with a named authority as subject: "means to say".
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
- ר' before a name: "Rabbi" (בר' בחיי: "Rabbi Bachya"); רבינו = "Rabbeinu".
- A Hebrew period with no main clause that is completed in the next segment:
  keep the Hebrew full stops and record the continuation in the segment `note`;
  join with semicolons only where the English becomes unreadable.
- A section number of a book the author cites: "section 172", not "§"; "§" is
  only for the author's own numbered sections.
- פ' + number after a book with no parashot (Menorat ha-Ma'or): "chapter".
- Lev 19:16: translate what the author writes; "stand idly by" adds "idly": use
  "stand by the blood of your fellow".
- וכ' read as וכו' ("etc.") where it closes a quotation.
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


## Milchamot ha-Shem 4:5

Hebrew (Wikisource): כב) חרדה גדולה עד מאד יחרד האיש הישראלי באמת מדברי הפילוסוף הנז' על הדברים האלה חילוף האלוהות הנז' פה, וחלוק עשרה מאמרות שבהם נברא העולם לאלוהות שונים זה מזה! ובסנהדרין פרק אחד דיני ממונות תנן לפיכך נברא אדם יחידי וכו' ושלא יהיו המינים אומרים הרבה רשויות בשמים וכו' פי' רש"י וכל אחד ברא את שלו ע"כ, ובבריתא ת"ר אדם נברא יחידי מפני מה שלא יהיו המינים אומרים הרבה רשיות בשמים, עוד שם אמר רב יהודה א"ר בשעה שבקש הקב"ה לבראת את האדם ברא כת אחת של מלאכי השרת אמר להם רצונכם נעשה אדם בצלמינו, אמרו לפניו רבונו של עולם מה מעשיו אמר להם מעשיו כך וכך אמרו לפניו רבונו של עולם מה אנוש כי תזכרנו, הושיט אצבעו הקטנה ושרפם, וכן כת שנייה. כת שלישית אמרה לפניו רבש"ע ראשונים שאמרו לפניך מה הועילו? כל העולם כולו שלך הוא, כל מה שאתה רוצה לעשות בעולמך עשה, כיון שהגיע לאנשי דור המבול ואנשי דור הפלגה. אמרו לפני רבש"ע לא יפה אמרו הראשונים לפניך? אמר להם ועד זקנה אני הוא ועד שיבה אני אסבל. ועיין מתרש"א ועיין לקמן סי' כ"ה.

| id | word | punctuation after |
|---|---|---|
| 4:5.1 | כב | ) |
| 4:5.2 | חרדה |  |
| 4:5.3 | גדולה |  |
| 4:5.4 | עד |  |
| 4:5.5 | מאד |  |
| 4:5.6 | יחרד |  |
| 4:5.7 | האיש |  |
| 4:5.8 | הישראלי |  |
| 4:5.9 | באמת |  |
| 4:5.10 | מדברי |  |
| 4:5.11 | הפילוסוף |  |
| 4:5.12 | הנז | ' |
| 4:5.13 | על |  |
| 4:5.14 | הדברים |  |
| 4:5.15 | האלה |  |
| 4:5.16 | חילוף |  |
| 4:5.17 | האלוהות |  |
| 4:5.18 | הנז | ' |
| 4:5.19 | פה | , |
| 4:5.20 | וחלוק |  |
| 4:5.21 | עשרה |  |
| 4:5.22 | מאמרות |  |
| 4:5.23 | שבהם |  |
| 4:5.24 | נברא |  |
| 4:5.25 | העולם |  |
| 4:5.26 | לאלוהות |  |
| 4:5.27 | שונים |  |
| 4:5.28 | זה |  |
| 4:5.29 | מזה | ! |
| 4:5.30 | ובסנהדרין |  |
| 4:5.31 | פרק |  |
| 4:5.32 | אחד |  |
| 4:5.33 | דיני |  |
| 4:5.34 | ממונות |  |
| 4:5.35 | תנן |  |
| 4:5.36 | לפיכך |  |
| 4:5.37 | נברא |  |
| 4:5.38 | אדם |  |
| 4:5.39 | יחידי |  |
| 4:5.40 | וכו | ' |
| 4:5.41 | ושלא |  |
| 4:5.42 | יהיו |  |
| 4:5.43 | המינים |  |
| 4:5.44 | אומרים |  |
| 4:5.45 | הרבה |  |
| 4:5.46 | רשויות |  |
| 4:5.47 | בשמים |  |
| 4:5.48 | וכו | ' |
| 4:5.49 | פי | ' |
| 4:5.50 | רש"י |  |
| 4:5.51 | וכל |  |
| 4:5.52 | אחד |  |
| 4:5.53 | ברא |  |
| 4:5.54 | את |  |
| 4:5.55 | שלו |  |
| 4:5.56 | ע"כ | , |
| 4:5.57 | ובבריתא |  |
| 4:5.58 | ת"ר |  |
| 4:5.59 | אדם |  |
| 4:5.60 | נברא |  |
| 4:5.61 | יחידי |  |
| 4:5.62 | מפני |  |
| 4:5.63 | מה |  |
| 4:5.64 | שלא |  |
| 4:5.65 | יהיו |  |
| 4:5.66 | המינים |  |
| 4:5.67 | אומרים |  |
| 4:5.68 | הרבה |  |
| 4:5.69 | רשיות |  |
| 4:5.70 | בשמים | , |
| 4:5.71 | עוד |  |
| 4:5.72 | שם |  |
| 4:5.73 | אמר |  |
| 4:5.74 | רב |  |
| 4:5.75 | יהודה |  |
| 4:5.76 | א"ר |  |
| 4:5.77 | בשעה |  |
| 4:5.78 | שבקש |  |
| 4:5.79 | הקב"ה |  |
| 4:5.80 | לבראת |  |
| 4:5.81 | את |  |
| 4:5.82 | האדם |  |
| 4:5.83 | ברא |  |
| 4:5.84 | כת |  |
| 4:5.85 | אחת |  |
| 4:5.86 | של |  |
| 4:5.87 | מלאכי |  |
| 4:5.88 | השרת |  |
| 4:5.89 | אמר |  |
| 4:5.90 | להם |  |
| 4:5.91 | רצונכם |  |
| 4:5.92 | נעשה |  |
| 4:5.93 | אדם |  |
| 4:5.94 | בצלמינו | , |
| 4:5.95 | אמרו |  |
| 4:5.96 | לפניו |  |
| 4:5.97 | רבונו |  |
| 4:5.98 | של |  |
| 4:5.99 | עולם |  |
| 4:5.100 | מה |  |
| 4:5.101 | מעשיו |  |
| 4:5.102 | אמר |  |
| 4:5.103 | להם |  |
| 4:5.104 | מעשיו |  |
| 4:5.105 | כך |  |
| 4:5.106 | וכך |  |
| 4:5.107 | אמרו |  |
| 4:5.108 | לפניו |  |
| 4:5.109 | רבונו |  |
| 4:5.110 | של |  |
| 4:5.111 | עולם |  |
| 4:5.112 | מה |  |
| 4:5.113 | אנוש |  |
| 4:5.114 | כי |  |
| 4:5.115 | תזכרנו | , |
| 4:5.116 | הושיט |  |
| 4:5.117 | אצבעו |  |
| 4:5.118 | הקטנה |  |
| 4:5.119 | ושרפם | , |
| 4:5.120 | וכן |  |
| 4:5.121 | כת |  |
| 4:5.122 | שנייה | . |
| 4:5.123 | כת |  |
| 4:5.124 | שלישית |  |
| 4:5.125 | אמרה |  |
| 4:5.126 | לפניו |  |
| 4:5.127 | רבש"ע |  |
| 4:5.128 | ראשונים |  |
| 4:5.129 | שאמרו |  |
| 4:5.130 | לפניך |  |
| 4:5.131 | מה |  |
| 4:5.132 | הועילו | ? |
| 4:5.133 | כל |  |
| 4:5.134 | העולם |  |
| 4:5.135 | כולו |  |
| 4:5.136 | שלך |  |
| 4:5.137 | הוא | , |
| 4:5.138 | כל |  |
| 4:5.139 | מה |  |
| 4:5.140 | שאתה |  |
| 4:5.141 | רוצה |  |
| 4:5.142 | לעשות |  |
| 4:5.143 | בעולמך |  |
| 4:5.144 | עשה | , |
| 4:5.145 | כיון |  |
| 4:5.146 | שהגיע |  |
| 4:5.147 | לאנשי |  |
| 4:5.148 | דור |  |
| 4:5.149 | המבול |  |
| 4:5.150 | ואנשי |  |
| 4:5.151 | דור |  |
| 4:5.152 | הפלגה | . |
| 4:5.153 | אמרו |  |
| 4:5.154 | לפני |  |
| 4:5.155 | רבש"ע |  |
| 4:5.156 | לא |  |
| 4:5.157 | יפה |  |
| 4:5.158 | אמרו |  |
| 4:5.159 | הראשונים |  |
| 4:5.160 | לפניך | ? |
| 4:5.161 | אמר |  |
| 4:5.162 | להם |  |
| 4:5.163 | ועד |  |
| 4:5.164 | זקנה |  |
| 4:5.165 | אני |  |
| 4:5.166 | הוא |  |
| 4:5.167 | ועד |  |
| 4:5.168 | שיבה |  |
| 4:5.169 | אני |  |
| 4:5.170 | אסבל | . |
| 4:5.171 | ועיין |  |
| 4:5.172 | מתרש"א |  |
| 4:5.173 | ועיין |  |
| 4:5.174 | לקמן |  |
| 4:5.175 | סי | ' |
| 4:5.176 | כ"ה | . |


## Milchamot ha-Shem 4:6

Hebrew (Wikisource): כג) ועוד גרסינן שם אמר ר' יוחנן כל מקום שפקרו המינים תשובתן בצדן נעשה אדם בצלמינו כדמותינו, ויברא אלהים את האדם בצלמו, הבה נרדה ונבלה שם שפתם, וירד ה' לראות את העיר וכו' ופי' רש"י כל מקום שפקרו המינים שמביאין ראיה להפקירן מן התורה ( ר"ל שמביאין ראיה לרבוי אלוהות) תשובתן בצדן ששם באותו ענין מפורש בתורה שהוא אחד, וכמ"ש הסמ"ג משם ר' סעדיה גאון יובא לקמן סי' ס"א. ובזהר שהבאתי דבריו לעיל אמר נעשה ודאי על תרין אתמר ומפרש שאמא אמרה לאבא נעשה אדם. ועשתה כחפצה ובראה את אדם שלא בהסכמתו כמו שכתבתי לעיל בשם שביבי נגה. וכי לא זו היא דעת הפוקרים ומביאים ראיה להפקירן שאלהות הרבה הם וכל אחד עושה כחפצו?

| id | word | punctuation after |
|---|---|---|
| 4:6.1 | כג | ) |
| 4:6.2 | ועוד |  |
| 4:6.3 | גרסינן |  |
| 4:6.4 | שם |  |
| 4:6.5 | אמר |  |
| 4:6.6 | ר | ' |
| 4:6.7 | יוחנן |  |
| 4:6.8 | כל |  |
| 4:6.9 | מקום |  |
| 4:6.10 | שפקרו |  |
| 4:6.11 | המינים |  |
| 4:6.12 | תשובתן |  |
| 4:6.13 | בצדן |  |
| 4:6.14 | נעשה |  |
| 4:6.15 | אדם |  |
| 4:6.16 | בצלמינו |  |
| 4:6.17 | כדמותינו | , |
| 4:6.18 | ויברא |  |
| 4:6.19 | אלהים |  |
| 4:6.20 | את |  |
| 4:6.21 | האדם |  |
| 4:6.22 | בצלמו | , |
| 4:6.23 | הבה |  |
| 4:6.24 | נרדה |  |
| 4:6.25 | ונבלה |  |
| 4:6.26 | שם |  |
| 4:6.27 | שפתם | , |
| 4:6.28 | וירד |  |
| 4:6.29 | ה | ' |
| 4:6.30 | לראות |  |
| 4:6.31 | את |  |
| 4:6.32 | העיר |  |
| 4:6.33 | וכו | ' |
| 4:6.34 | ופי | ' |
| 4:6.35 | רש"י |  |
| 4:6.36 | כל |  |
| 4:6.37 | מקום |  |
| 4:6.38 | שפקרו |  |
| 4:6.39 | המינים |  |
| 4:6.40 | שמביאין |  |
| 4:6.41 | ראיה |  |
| 4:6.42 | להפקירן |  |
| 4:6.43 | מן |  |
| 4:6.44 | התורה | ( |
| 4:6.45 | ר"ל |  |
| 4:6.46 | שמביאין |  |
| 4:6.47 | ראיה |  |
| 4:6.48 | לרבוי |  |
| 4:6.49 | אלוהות | ) |
| 4:6.50 | תשובתן |  |
| 4:6.51 | בצדן |  |
| 4:6.52 | ששם |  |
| 4:6.53 | באותו |  |
| 4:6.54 | ענין |  |
| 4:6.55 | מפורש |  |
| 4:6.56 | בתורה |  |
| 4:6.57 | שהוא |  |
| 4:6.58 | אחד | , |
| 4:6.59 | וכמ"ש |  |
| 4:6.60 | הסמ"ג |  |
| 4:6.61 | משם |  |
| 4:6.62 | ר | ' |
| 4:6.63 | סעדיה |  |
| 4:6.64 | גאון |  |
| 4:6.65 | יובא |  |
| 4:6.66 | לקמן |  |
| 4:6.67 | סי | ' |
| 4:6.68 | ס"א | . |
| 4:6.69 | ובזהר |  |
| 4:6.70 | שהבאתי |  |
| 4:6.71 | דבריו |  |
| 4:6.72 | לעיל |  |
| 4:6.73 | אמר |  |
| 4:6.74 | נעשה |  |
| 4:6.75 | ודאי |  |
| 4:6.76 | על |  |
| 4:6.77 | תרין |  |
| 4:6.78 | אתמר |  |
| 4:6.79 | ומפרש |  |
| 4:6.80 | שאמא |  |
| 4:6.81 | אמרה |  |
| 4:6.82 | לאבא |  |
| 4:6.83 | נעשה |  |
| 4:6.84 | אדם | . |
| 4:6.85 | ועשתה |  |
| 4:6.86 | כחפצה |  |
| 4:6.87 | ובראה |  |
| 4:6.88 | את |  |
| 4:6.89 | אדם |  |
| 4:6.90 | שלא |  |
| 4:6.91 | בהסכמתו |  |
| 4:6.92 | כמו |  |
| 4:6.93 | שכתבתי |  |
| 4:6.94 | לעיל |  |
| 4:6.95 | בשם |  |
| 4:6.96 | שביבי |  |
| 4:6.97 | נגה | . |
| 4:6.98 | וכי |  |
| 4:6.99 | לא |  |
| 4:6.100 | זו |  |
| 4:6.101 | היא |  |
| 4:6.102 | דעת |  |
| 4:6.103 | הפוקרים |  |
| 4:6.104 | ומביאים |  |
| 4:6.105 | ראיה |  |
| 4:6.106 | להפקירן |  |
| 4:6.107 | שאלהות |  |
| 4:6.108 | הרבה |  |
| 4:6.109 | הם |  |
| 4:6.110 | וכל |  |
| 4:6.111 | אחד |  |
| 4:6.112 | עושה |  |
| 4:6.113 | כחפצו | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

