# Review pack: milhamot-hashem perek-3-05 (Milchamot ha-Shem 3:9–3:9)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-05`.

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
- יש מאין / יש מיש: translate, "something from nothing" / "something from
  something"; no transliteration.
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
- A printed word that echoes a verse's likelier word (נבוש for נמוש, Job 23:12):
  translate the printed word, record the likelier one in the token note and
  `apparatus`.
- A quotation that runs over several paragraphs (Rabbi Menasseh ben Israel,
  Teshu'at Yisrael, 2:18-2:25, opened by כ' וז"ל, closed by עכ"ל): each paragraph
  opens with a double mark, and only the last closes it.
- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
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
- קורין אל: "call upon"; חכמי הקבלה: "sages of Kabbalah"; גבורה as a name of God:
  "the Power".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
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


## Milchamot ha-Shem 3:9

Hebrew (Wikisource): ט) ואתה ידידנו לא דייך שלא עשית אחת מכל אלה, אלא שהאמנת לדבריהם וקבלת שקריהם, בעד אחד ודיין אחד, ושלא בפני בעל דין, ובלא חקירה ובלא דרישה, ובלא בדיקה, ונהנית בדבריהם, וגם סייעת לחזק שקריהם ונרשם בדעתך כאמת לאמתו, עד שכתבת שאיני יכול להכחיש עוד וכו'. ולא שמת אל לבך שאין מקבלין עדות אלא בפני בעל דין, ושאתה מחויב לדרוש ולחקור אותם בשבע חקירות ובדרישות ובדיקות כמצווה עליך. כל שכן שעשו במחשך מעשיהם, והחליקו לשונם אליך בסתר דבריהם, למצוא עון לשנוא והית מחויב לשתקם בנזיפה. אין זה כי אם רוע לב שאתה נהנה ומשתוקק ברכילותם ומלשינותם. רחמנא ליצלן מהאי דעתא! ואפשר שגם אתה מוסיף משלך מלשינות שקר כדי להתכבד בקלון חבריך. ובפרק ערבי פסחים (קי"ה) גרסינן אמר רב ששת משום ראב"ע כל המספר לשון הרע, והמעיד עדות שקר בחברו ראוי להשליכו לכלבים שנאמר לכלב תשליכון אותו, וכתיב בהדיא לא תשא שמע שוא וקרי ביה לא תשיא, ועוד כמה דרשות מצינו לרז"ל בעונש המספר לשון הרע. ובמחזור ויטרי מסכת אבות על הא דתנן על תדון את חבירך עד שתגיע חמקומו כ' עש שתבא לידך אותה עבירה ותנצל ותשמט ממנה כמו שמצינו בירבעם בן נבט, שבשעה שנשא שלמה את בת פרעה הכניסה לו כל מיני זמר שבעולם ואמרה לו כך עושין לע"ז פלונית פעם אחת הלך לישן והעריב היום הלכה וקראה לחרשי העץ ולבונים, העושים במלאכת חרש וחושב, ועשו על מטתו כעין רקיע ובו חמה ולבנה וכוכבים ומזלות. ובבקר העיר משנתו לפתוח דלתות ההיכל, שהיו מפתחות ההיכל תחת מראשותיו ראה שהרקיע בהיר בכוכבים וצבא, כסבור עדיין לא פנה לילה, חזר לישן. וישן עוד ארבע שעות, עד שהבין והכיר מעצמו שאין הלילה ארוך כל כך. ועל אותה שעה שנינו עדות על תמיד של שחר שקרב בארבע שעות, בבחירתא. ובטרם הקיץ משנתו הלך ירבעם בן נבט על פתח ביתו, ויעבר קול, רשע, עד מתי תישן ותבטל תמיד של שחר, מה לך נרדם, קום קרא אל אלהיך! יצתה בת קול ואמרה לו רשע בן רשע עתיד אתה לבטל מישראל כמה וכמה קרבנות במזיד. ואתה מחייבו על השוגג כמזיד. ועל האונס כרצון! והיינו דכתיב "כדבר אפרים רתת, נשא הוא בישראל, ויאשם בבעל וימות" בדבר ירבעם שהוא משבט אפרים, רתת קטיגורין על שלמה, הוא ירבעם נשיא בישראל ומלך, ובת קול השיבתו שיאשם בבעל וימות עכ"ל. ואתה הגדלת והוספת לעשות רע יותר מירבעם בן נבט. כי אחרי אשר הבאתם אותנו אל בית הכלא במלשינות שקר אצל ממשלתינו הרוממה! הזעקת את כל העם לצאת אל בית עלמין, וצוית להם לעפר בעפר על ראשיהם על לא דבר. לא על מצוה קלה שבטלנו ולא על מצוה חמורה ח"ו כי אם על אשר הלכנו כמנהגי קדמונינו וקבענו למודינו במשנה ובתלמוד ורמב"ם וש"ע ללמוד וללמד, לשמור ולעשות. ולא למדנו עוד בספר מזויף ספר הזהר כמנהג שלכם.

| id | word | punctuation after |
|---|---|---|
| 3:9.1 | ט | ) |
| 3:9.2 | ואתה |  |
| 3:9.3 | ידידנו |  |
| 3:9.4 | לא |  |
| 3:9.5 | דייך |  |
| 3:9.6 | שלא |  |
| 3:9.7 | עשית |  |
| 3:9.8 | אחת |  |
| 3:9.9 | מכל |  |
| 3:9.10 | אלה | , |
| 3:9.11 | אלא |  |
| 3:9.12 | שהאמנת |  |
| 3:9.13 | לדבריהם |  |
| 3:9.14 | וקבלת |  |
| 3:9.15 | שקריהם | , |
| 3:9.16 | בעד |  |
| 3:9.17 | אחד |  |
| 3:9.18 | ודיין |  |
| 3:9.19 | אחד | , |
| 3:9.20 | ושלא |  |
| 3:9.21 | בפני |  |
| 3:9.22 | בעל |  |
| 3:9.23 | דין | , |
| 3:9.24 | ובלא |  |
| 3:9.25 | חקירה |  |
| 3:9.26 | ובלא |  |
| 3:9.27 | דרישה | , |
| 3:9.28 | ובלא |  |
| 3:9.29 | בדיקה | , |
| 3:9.30 | ונהנית |  |
| 3:9.31 | בדבריהם | , |
| 3:9.32 | וגם |  |
| 3:9.33 | סייעת |  |
| 3:9.34 | לחזק |  |
| 3:9.35 | שקריהם |  |
| 3:9.36 | ונרשם |  |
| 3:9.37 | בדעתך |  |
| 3:9.38 | כאמת |  |
| 3:9.39 | לאמתו | , |
| 3:9.40 | עד |  |
| 3:9.41 | שכתבת |  |
| 3:9.42 | שאיני |  |
| 3:9.43 | יכול |  |
| 3:9.44 | להכחיש |  |
| 3:9.45 | עוד |  |
| 3:9.46 | וכו | '. |
| 3:9.47 | ולא |  |
| 3:9.48 | שמת |  |
| 3:9.49 | אל |  |
| 3:9.50 | לבך |  |
| 3:9.51 | שאין |  |
| 3:9.52 | מקבלין |  |
| 3:9.53 | עדות |  |
| 3:9.54 | אלא |  |
| 3:9.55 | בפני |  |
| 3:9.56 | בעל |  |
| 3:9.57 | דין | , |
| 3:9.58 | ושאתה |  |
| 3:9.59 | מחויב |  |
| 3:9.60 | לדרוש |  |
| 3:9.61 | ולחקור |  |
| 3:9.62 | אותם |  |
| 3:9.63 | בשבע |  |
| 3:9.64 | חקירות |  |
| 3:9.65 | ובדרישות |  |
| 3:9.66 | ובדיקות |  |
| 3:9.67 | כמצווה |  |
| 3:9.68 | עליך | . |
| 3:9.69 | כל |  |
| 3:9.70 | שכן |  |
| 3:9.71 | שעשו |  |
| 3:9.72 | במחשך |  |
| 3:9.73 | מעשיהם | , |
| 3:9.74 | והחליקו |  |
| 3:9.75 | לשונם |  |
| 3:9.76 | אליך |  |
| 3:9.77 | בסתר |  |
| 3:9.78 | דבריהם | , |
| 3:9.79 | למצוא |  |
| 3:9.80 | עון |  |
| 3:9.81 | לשנוא |  |
| 3:9.82 | והית |  |
| 3:9.83 | מחויב |  |
| 3:9.84 | לשתקם |  |
| 3:9.85 | בנזיפה | . |
| 3:9.86 | אין |  |
| 3:9.87 | זה |  |
| 3:9.88 | כי |  |
| 3:9.89 | אם |  |
| 3:9.90 | רוע |  |
| 3:9.91 | לב |  |
| 3:9.92 | שאתה |  |
| 3:9.93 | נהנה |  |
| 3:9.94 | ומשתוקק |  |
| 3:9.95 | ברכילותם |  |
| 3:9.96 | ומלשינותם | . |
| 3:9.97 | רחמנא |  |
| 3:9.98 | ליצלן |  |
| 3:9.99 | מהאי |  |
| 3:9.100 | דעתא | ! |
| 3:9.101 | ואפשר |  |
| 3:9.102 | שגם |  |
| 3:9.103 | אתה |  |
| 3:9.104 | מוסיף |  |
| 3:9.105 | משלך |  |
| 3:9.106 | מלשינות |  |
| 3:9.107 | שקר |  |
| 3:9.108 | כדי |  |
| 3:9.109 | להתכבד |  |
| 3:9.110 | בקלון |  |
| 3:9.111 | חבריך | . |
| 3:9.112 | ובפרק |  |
| 3:9.113 | ערבי |  |
| 3:9.114 | פסחים |  |
| 3:9.115 | קי"ה | ) |
| 3:9.116 | גרסינן |  |
| 3:9.117 | אמר |  |
| 3:9.118 | רב |  |
| 3:9.119 | ששת |  |
| 3:9.120 | משום |  |
| 3:9.121 | ראב"ע |  |
| 3:9.122 | כל |  |
| 3:9.123 | המספר |  |
| 3:9.124 | לשון |  |
| 3:9.125 | הרע | , |
| 3:9.126 | והמעיד |  |
| 3:9.127 | עדות |  |
| 3:9.128 | שקר |  |
| 3:9.129 | בחברו |  |
| 3:9.130 | ראוי |  |
| 3:9.131 | להשליכו |  |
| 3:9.132 | לכלבים |  |
| 3:9.133 | שנאמר |  |
| 3:9.134 | לכלב |  |
| 3:9.135 | תשליכון |  |
| 3:9.136 | אותו | , |
| 3:9.137 | וכתיב |  |
| 3:9.138 | בהדיא |  |
| 3:9.139 | לא |  |
| 3:9.140 | תשא |  |
| 3:9.141 | שמע |  |
| 3:9.142 | שוא |  |
| 3:9.143 | וקרי |  |
| 3:9.144 | ביה |  |
| 3:9.145 | לא |  |
| 3:9.146 | תשיא | , |
| 3:9.147 | ועוד |  |
| 3:9.148 | כמה |  |
| 3:9.149 | דרשות |  |
| 3:9.150 | מצינו |  |
| 3:9.151 | לרז"ל |  |
| 3:9.152 | בעונש |  |
| 3:9.153 | המספר |  |
| 3:9.154 | לשון |  |
| 3:9.155 | הרע | . |
| 3:9.156 | ובמחזור |  |
| 3:9.157 | ויטרי |  |
| 3:9.158 | מסכת |  |
| 3:9.159 | אבות |  |
| 3:9.160 | על |  |
| 3:9.161 | הא |  |
| 3:9.162 | דתנן |  |
| 3:9.163 | על |  |
| 3:9.164 | תדון |  |
| 3:9.165 | את |  |
| 3:9.166 | חבירך |  |
| 3:9.167 | עד |  |
| 3:9.168 | שתגיע |  |
| 3:9.169 | חמקומו |  |
| 3:9.170 | כ | ' |
| 3:9.171 | עש |  |
| 3:9.172 | שתבא |  |
| 3:9.173 | לידך |  |
| 3:9.174 | אותה |  |
| 3:9.175 | עבירה |  |
| 3:9.176 | ותנצל |  |
| 3:9.177 | ותשמט |  |
| 3:9.178 | ממנה |  |
| 3:9.179 | כמו |  |
| 3:9.180 | שמצינו |  |
| 3:9.181 | בירבעם |  |
| 3:9.182 | בן |  |
| 3:9.183 | נבט | , |
| 3:9.184 | שבשעה |  |
| 3:9.185 | שנשא |  |
| 3:9.186 | שלמה |  |
| 3:9.187 | את |  |
| 3:9.188 | בת |  |
| 3:9.189 | פרעה |  |
| 3:9.190 | הכניסה |  |
| 3:9.191 | לו |  |
| 3:9.192 | כל |  |
| 3:9.193 | מיני |  |
| 3:9.194 | זמר |  |
| 3:9.195 | שבעולם |  |
| 3:9.196 | ואמרה |  |
| 3:9.197 | לו |  |
| 3:9.198 | כך |  |
| 3:9.199 | עושין |  |
| 3:9.200 | לע"ז |  |
| 3:9.201 | פלונית |  |
| 3:9.202 | פעם |  |
| 3:9.203 | אחת |  |
| 3:9.204 | הלך |  |
| 3:9.205 | לישן |  |
| 3:9.206 | והעריב |  |
| 3:9.207 | היום |  |
| 3:9.208 | הלכה |  |
| 3:9.209 | וקראה |  |
| 3:9.210 | לחרשי |  |
| 3:9.211 | העץ |  |
| 3:9.212 | ולבונים | , |
| 3:9.213 | העושים |  |
| 3:9.214 | במלאכת |  |
| 3:9.215 | חרש |  |
| 3:9.216 | וחושב | , |
| 3:9.217 | ועשו |  |
| 3:9.218 | על |  |
| 3:9.219 | מטתו |  |
| 3:9.220 | כעין |  |
| 3:9.221 | רקיע |  |
| 3:9.222 | ובו |  |
| 3:9.223 | חמה |  |
| 3:9.224 | ולבנה |  |
| 3:9.225 | וכוכבים |  |
| 3:9.226 | ומזלות | . |
| 3:9.227 | ובבקר |  |
| 3:9.228 | העיר |  |
| 3:9.229 | משנתו |  |
| 3:9.230 | לפתוח |  |
| 3:9.231 | דלתות |  |
| 3:9.232 | ההיכל | , |
| 3:9.233 | שהיו |  |
| 3:9.234 | מפתחות |  |
| 3:9.235 | ההיכל |  |
| 3:9.236 | תחת |  |
| 3:9.237 | מראשותיו |  |
| 3:9.238 | ראה |  |
| 3:9.239 | שהרקיע |  |
| 3:9.240 | בהיר |  |
| 3:9.241 | בכוכבים |  |
| 3:9.242 | וצבא | , |
| 3:9.243 | כסבור |  |
| 3:9.244 | עדיין |  |
| 3:9.245 | לא |  |
| 3:9.246 | פנה |  |
| 3:9.247 | לילה | , |
| 3:9.248 | חזר |  |
| 3:9.249 | לישן | . |
| 3:9.250 | וישן |  |
| 3:9.251 | עוד |  |
| 3:9.252 | ארבע |  |
| 3:9.253 | שעות | , |
| 3:9.254 | עד |  |
| 3:9.255 | שהבין |  |
| 3:9.256 | והכיר |  |
| 3:9.257 | מעצמו |  |
| 3:9.258 | שאין |  |
| 3:9.259 | הלילה |  |
| 3:9.260 | ארוך |  |
| 3:9.261 | כל |  |
| 3:9.262 | כך | . |
| 3:9.263 | ועל |  |
| 3:9.264 | אותה |  |
| 3:9.265 | שעה |  |
| 3:9.266 | שנינו |  |
| 3:9.267 | עדות |  |
| 3:9.268 | על |  |
| 3:9.269 | תמיד |  |
| 3:9.270 | של |  |
| 3:9.271 | שחר |  |
| 3:9.272 | שקרב |  |
| 3:9.273 | בארבע |  |
| 3:9.274 | שעות | , |
| 3:9.275 | בבחירתא | . |
| 3:9.276 | ובטרם |  |
| 3:9.277 | הקיץ |  |
| 3:9.278 | משנתו |  |
| 3:9.279 | הלך |  |
| 3:9.280 | ירבעם |  |
| 3:9.281 | בן |  |
| 3:9.282 | נבט |  |
| 3:9.283 | על |  |
| 3:9.284 | פתח |  |
| 3:9.285 | ביתו | , |
| 3:9.286 | ויעבר |  |
| 3:9.287 | קול | , |
| 3:9.288 | רשע | , |
| 3:9.289 | עד |  |
| 3:9.290 | מתי |  |
| 3:9.291 | תישן |  |
| 3:9.292 | ותבטל |  |
| 3:9.293 | תמיד |  |
| 3:9.294 | של |  |
| 3:9.295 | שחר | , |
| 3:9.296 | מה |  |
| 3:9.297 | לך |  |
| 3:9.298 | נרדם | , |
| 3:9.299 | קום |  |
| 3:9.300 | קרא |  |
| 3:9.301 | אל |  |
| 3:9.302 | אלהיך | ! |
| 3:9.303 | יצתה |  |
| 3:9.304 | בת |  |
| 3:9.305 | קול |  |
| 3:9.306 | ואמרה |  |
| 3:9.307 | לו |  |
| 3:9.308 | רשע |  |
| 3:9.309 | בן |  |
| 3:9.310 | רשע |  |
| 3:9.311 | עתיד |  |
| 3:9.312 | אתה |  |
| 3:9.313 | לבטל |  |
| 3:9.314 | מישראל |  |
| 3:9.315 | כמה |  |
| 3:9.316 | וכמה |  |
| 3:9.317 | קרבנות |  |
| 3:9.318 | במזיד | . |
| 3:9.319 | ואתה |  |
| 3:9.320 | מחייבו |  |
| 3:9.321 | על |  |
| 3:9.322 | השוגג |  |
| 3:9.323 | כמזיד | . |
| 3:9.324 | ועל |  |
| 3:9.325 | האונס |  |
| 3:9.326 | כרצון | ! |
| 3:9.327 | והיינו |  |
| 3:9.328 | דכתיב |  |
| 3:9.329 | כדבר |  |
| 3:9.330 | אפרים |  |
| 3:9.331 | רתת | , |
| 3:9.332 | נשא |  |
| 3:9.333 | הוא |  |
| 3:9.334 | בישראל | , |
| 3:9.335 | ויאשם |  |
| 3:9.336 | בבעל |  |
| 3:9.337 | וימות | " |
| 3:9.338 | בדבר |  |
| 3:9.339 | ירבעם |  |
| 3:9.340 | שהוא |  |
| 3:9.341 | משבט |  |
| 3:9.342 | אפרים | , |
| 3:9.343 | רתת |  |
| 3:9.344 | קטיגורין |  |
| 3:9.345 | על |  |
| 3:9.346 | שלמה | , |
| 3:9.347 | הוא |  |
| 3:9.348 | ירבעם |  |
| 3:9.349 | נשיא |  |
| 3:9.350 | בישראל |  |
| 3:9.351 | ומלך | , |
| 3:9.352 | ובת |  |
| 3:9.353 | קול |  |
| 3:9.354 | השיבתו |  |
| 3:9.355 | שיאשם |  |
| 3:9.356 | בבעל |  |
| 3:9.357 | וימות |  |
| 3:9.358 | עכ"ל | . |
| 3:9.359 | ואתה |  |
| 3:9.360 | הגדלת |  |
| 3:9.361 | והוספת |  |
| 3:9.362 | לעשות |  |
| 3:9.363 | רע |  |
| 3:9.364 | יותר |  |
| 3:9.365 | מירבעם |  |
| 3:9.366 | בן |  |
| 3:9.367 | נבט | . |
| 3:9.368 | כי |  |
| 3:9.369 | אחרי |  |
| 3:9.370 | אשר |  |
| 3:9.371 | הבאתם |  |
| 3:9.372 | אותנו |  |
| 3:9.373 | אל |  |
| 3:9.374 | בית |  |
| 3:9.375 | הכלא |  |
| 3:9.376 | במלשינות |  |
| 3:9.377 | שקר |  |
| 3:9.378 | אצל |  |
| 3:9.379 | ממשלתינו |  |
| 3:9.380 | הרוממה | ! |
| 3:9.381 | הזעקת |  |
| 3:9.382 | את |  |
| 3:9.383 | כל |  |
| 3:9.384 | העם |  |
| 3:9.385 | לצאת |  |
| 3:9.386 | אל |  |
| 3:9.387 | בית |  |
| 3:9.388 | עלמין | , |
| 3:9.389 | וצוית |  |
| 3:9.390 | להם |  |
| 3:9.391 | לעפר |  |
| 3:9.392 | בעפר |  |
| 3:9.393 | על |  |
| 3:9.394 | ראשיהם |  |
| 3:9.395 | על |  |
| 3:9.396 | לא |  |
| 3:9.397 | דבר | . |
| 3:9.398 | לא |  |
| 3:9.399 | על |  |
| 3:9.400 | מצוה |  |
| 3:9.401 | קלה |  |
| 3:9.402 | שבטלנו |  |
| 3:9.403 | ולא |  |
| 3:9.404 | על |  |
| 3:9.405 | מצוה |  |
| 3:9.406 | חמורה |  |
| 3:9.407 | ח"ו |  |
| 3:9.408 | כי |  |
| 3:9.409 | אם |  |
| 3:9.410 | על |  |
| 3:9.411 | אשר |  |
| 3:9.412 | הלכנו |  |
| 3:9.413 | כמנהגי |  |
| 3:9.414 | קדמונינו |  |
| 3:9.415 | וקבענו |  |
| 3:9.416 | למודינו |  |
| 3:9.417 | במשנה |  |
| 3:9.418 | ובתלמוד |  |
| 3:9.419 | ורמב"ם |  |
| 3:9.420 | וש"ע |  |
| 3:9.421 | ללמוד |  |
| 3:9.422 | וללמד | , |
| 3:9.423 | לשמור |  |
| 3:9.424 | ולעשות | . |
| 3:9.425 | ולא |  |
| 3:9.426 | למדנו |  |
| 3:9.427 | עוד |  |
| 3:9.428 | בספר |  |
| 3:9.429 | מזויף |  |
| 3:9.430 | ספר |  |
| 3:9.431 | הזהר |  |
| 3:9.432 | כמנהג |  |
| 3:9.433 | שלכם | . |


### Translation under review (3:9)

English: 9) And you, our friend, it was not enough for you that you did not do any one of all these but you believed their words and accepted their lies on the testimony of one witness and one judge and not in the presence of the litigant without inquiry, without investigation, without examination and you enjoyed their words and you also helped to strengthen their lies and it was fixed in your mind as the truth itself until you wrote that I am no longer able to deny it, etc. And you did not take to heart that testimony is accepted only in the presence of the litigant and that you are obliged to inquire into them with seven inquiries, investigations and examinations, as you are commanded All the more so, since they did their deeds in the dark and smoothed their tongues to you in secret with their words to find a fault to hate and you were obliged to silence them with a rebuke. This is nothing but a wicked heart: that you enjoy and long for their gossip and their informing. the Merciful One save us from this opinion! And it is possible that you too add of your own false informing, in order to honor yourself at the disgrace of your friends. And in the chapter Arvei Pesachim (115) we read: "Rav Sheshet said in the name of R. A.: whoever tells evil talk and whoever bears false witness against his fellow deserves to be thrown to the dogs, as it is said: 'to the dog you shall throw it' and it is written explicitly: 'you shall not carry a false report'; and it is read in it: 'you shall not cause to bear'" And we find several more homilies of our Sages of blessed memory on the punishment of one who tells evil talk. "And in Mahzor Vitry, tractate Avot, on what we learned: 'do not judge your fellow until you reach his place' as he writes there: that the same transgression may come into your hand and you may be saved and escape from it as we find with Jeroboam son of Nebat that when Solomon married the daughter of Pharaoh, she brought in to him every kind of song in the world and said to him: 'So they do for such-and-such idolatry' One time, he went to sleep, and the day grew late she went and called the wood-carvers and the builders, who work at carving and weaving and they made over his bed the likeness of a firmament, and in it the sun, the moon, the stars and the constellations In the morning he woke from his sleep to open the doors of the Temple, for the keys of the Temple were under his pillow he saw that the firmament was bright with stars and the host thinking that the night had not yet turned he went back to sleep And he slept four more hours, until he understood and recognised for himself that the night was not so long And concerning that hour we learned a testimony about the morning daily offering, which is offered at four hours, in the Bechirta. Before he woke from his sleep, Jeroboam son of Nebat went to the door of his house and a voice went out: 'Wicked one how long will you sleep and neglect the morning daily offering What is it to you, sleeper? Arise, call upon your God!' A bat kol came out and said to him: 'Wicked son of a wicked man, you are destined to abolish from Israel many, many offerings wilfully. And you hold him liable for the inadvertent as for the wilful. And for the forced one as for the voluntary!' And this is what is written: 'When Ephraim spoke there was trembling; he was exalted in Israel, but he incurred guilt through Baal and died' This is said of Jeroboam, who was of the tribe of Ephraim, the trembling was the accuser against Solomon, he, Jeroboam, was prince in Israel and king and the bat kol answered him that he would become guilty through Baal and die." Thus far his words. And you have magnified and added to doing evil more than Jeroboam son of Nebat. For after you brought us to the prison by false informing before our exalted government! You called out all the people to go out to the cemetery, and commanded them to dust themselves with dust upon their heads for no reason. Not for a light commandment that we neglected, and not for a weighty commandment, God forbid, but because we walked in the custom of our forefathers and set our learning in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch to learn and to teach, to keep and to do. And we did not study further from the forged book, the book the Zohar, as is your custom.

Literal: 9) and you our friend not enough for you that not you did one of all these but that you believed to their words and you accepted their lies on the testimony of one and a judge one and not in the presence of the litigant of the case and without inquiry and without investigation and without examination and you enjoyed in their words and also you helped to strengthen their lies and was fixed in your mind as truth to its truth until that you wrote that I am able to deny any longer etc. and not you put to your heart that there is not they accept testimony but only in the presence of a litigant of the case and that you are obliged to inquire and to investigate them with seven inquiries and with investigations and with examinations as commanded of you all the more so that they did in the dark their deeds and smoothed their tongue to you in secret with their words to find a fault to hate and you were obliged to silence them with a rebuke it is not this that except wickedness of heart that you enjoy and long for their gossip and their informing the Merciful One save us from this opinion and it is possible that also you add of your own informing false in order to honor yourself at the disgrace of your friends and in the chapter Arvei Pesachim (folio 115) we read said Rav Sheshet in the name of R. A. whoever tells evil talk and whoever testifies testimony false against his fellow is fit to cast him to the dogs as it is said to the dog you shall throw it and it is written explicitly not you shall carry report false and it is read in it not you shall cause to bear and further several homilies we find our Sages of blessed memory on the punishment of the one who tells evil talk and in Mahzor Vitry tractate Avot on that which we learned on (al: do not) judge - your fellow until you reach his place as he writes there that may come into your hand that same transgression and be saved and escape from it as we find with Jeroboam son of Nebat that when he married Solomon - daughter of Pharaoh she brought in to him every kind of song in the world and she said to him thus they do for idolatry such-and-such once time he went to sleep and became late the day she went and called the wood-carvers of wood and the builders who do the craft of carving and weaving and they made on his bed the likeness of a firmament and in it the sun and the moon and the stars and the constellations and in the morning he woke from his sleep to open the doors of the Temple for there were keys of the Temple under his pillow he saw that the firmament was bright with stars and the host thinking still not had turned night he returned to sleep and he slept more four hours until he understood and recognised for himself that there is not the night long so thus and concerning that hour we learned testimony about the daily offering of morning which is offered at four hours in the Bechirta and before he woke from his sleep went Jeroboam son of Nebat to the door of his house and went out a voice wicked one until when you sleep and neglect the daily offering of morning what to you sleeper arise call upon your God came out bat (a heavenly) voice and said to him wicked son of wicked destined you to nullify from Israel many and many offerings wilfully and you hold him liable for the inadvertent as the wilful and for the forced as the voluntary and this is what is written as the speaking of Ephraim trembling he was exalted he in Israel and became guilty through Baal and died in the matter of Jeroboam who is of the tribe of Ephraim rattat (trembled) katigorin (accuser) against Solomon he Jeroboam prince in Israel and king and a bat kol answered him that he would become guilty through Baal and die thus far his words and you have magnified and added to do evil more than Jeroboam son of Nebat for after that you brought us to the house of the prison by false informing before our government exalted you called out - all the people to go out to the house of the cemetery and you commanded them to dust themselves with dust on their heads for no thing not for a commandment light that we neglected and not for a commandment weighty God forbid but only for that we walked in the custom of our forefathers and we set our learning in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch to learn and to teach to keep and to do and not we studied further from the book forged the book the Zohar as is the custom of you

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:9.1 | 9) |  |  | printed list marker "9)", kept as the author's own |
| 3:9.2 | and you |  |  |  |
| 3:9.3 | our friend |  |  | ידידנו: "our friend", the ironic form of address, kept |
| 3:9.4 | not |  |  |  |
| 3:9.5 | enough for you |  |  | לא דייך: "it was not enough for you"; the idiom governs the clause to 3:9.10 |
| 3:9.6 | that not |  |  |  |
| 3:9.7 | you did |  |  |  |
| 3:9.8 | one |  |  |  |
| 3:9.9 | of all |  |  |  |
| 3:9.10 | these |  |  |  |
| 3:9.11 | but |  |  |  |
| 3:9.12 | that you believed |  |  |  |
| 3:9.13 | to their words |  |  |  |
| 3:9.14 | and you accepted |  |  |  |
| 3:9.15 | their lies |  |  |  |
| 3:9.16 | on the testimony of |  |  |  |
| 3:9.17 | one |  |  |  |
| 3:9.18 | and a judge |  |  |  |
| 3:9.19 | one |  |  |  |
| 3:9.20 | and not |  |  |  |
| 3:9.21 | in the presence of |  |  |  |
| 3:9.22 | the litigant |  |  |  |
| 3:9.23 | of the case |  |  |  |
| 3:9.24 | and without |  |  |  |
| 3:9.25 | inquiry |  |  |  |
| 3:9.26 | and without |  |  |  |
| 3:9.27 | investigation |  |  |  |
| 3:9.28 | and without |  |  |  |
| 3:9.29 | examination |  |  |  |
| 3:9.30 | and you enjoyed |  |  |  |
| 3:9.31 | in their words |  |  |  |
| 3:9.32 | and also |  |  |  |
| 3:9.33 | you helped |  |  |  |
| 3:9.34 | to strengthen |  |  |  |
| 3:9.35 | their lies |  |  |  |
| 3:9.36 | and was fixed |  |  |  |
| 3:9.37 | in your mind |  |  |  |
| 3:9.38 | as truth |  |  |  |
| 3:9.39 | to its truth |  |  |  |
| 3:9.40 | until |  |  |  |
| 3:9.41 | that you wrote |  |  |  |
| 3:9.42 | that I |  |  |  |
| 3:9.43 | am able |  |  |  |
| 3:9.44 | to deny |  |  |  |
| 3:9.45 | any longer |  |  |  |
| 3:9.46 | etc. |  |  | וכו': marks a truncated quotation; kept as "etc." |
| 3:9.47 | and not |  |  |  |
| 3:9.48 | you put |  |  |  |
| 3:9.49 | to |  |  |  |
| 3:9.50 | your heart |  |  |  |
| 3:9.51 | that there is not |  |  |  |
| 3:9.52 | they accept |  |  |  |
| 3:9.53 | testimony |  |  |  |
| 3:9.54 | but only |  |  |  |
| 3:9.55 | in the presence of |  |  |  |
| 3:9.56 | a litigant |  |  |  |
| 3:9.57 | of the case |  |  |  |
| 3:9.58 | and that you |  |  |  |
| 3:9.59 | are obliged |  |  |  |
| 3:9.60 | to inquire |  |  |  |
| 3:9.61 | and to investigate |  |  |  |
| 3:9.62 | them |  |  |  |
| 3:9.63 | with seven |  |  |  |
| 3:9.64 | inquiries |  |  |  |
| 3:9.65 | and with investigations |  |  |  |
| 3:9.66 | and with examinations |  |  |  |
| 3:9.67 | as commanded |  |  |  |
| 3:9.68 | of you |  |  |  |
| 3:9.69 | all |  |  |  |
| 3:9.70 | the more so |  |  |  |
| 3:9.71 | that they did |  |  |  |
| 3:9.72 | in the dark |  |  |  |
| 3:9.73 | their deeds |  |  |  |
| 3:9.74 | and smoothed |  |  |  |
| 3:9.75 | their tongue |  |  |  |
| 3:9.76 | to you |  |  |  |
| 3:9.77 | in secret |  |  |  |
| 3:9.78 | with their words |  |  |  |
| 3:9.79 | to find |  |  |  |
| 3:9.80 | a fault |  |  |  |
| 3:9.81 | to hate |  |  |  |
| 3:9.82 | and you were |  |  | printed "והית" for "והיית" (and you were); the defective spelling is kept as printed |
| 3:9.83 | obliged |  |  |  |
| 3:9.84 | to silence them |  |  |  |
| 3:9.85 | with a rebuke |  |  |  |
| 3:9.86 | it is not |  |  |  |
| 3:9.87 | this |  |  |  |
| 3:9.88 | that |  |  |  |
| 3:9.89 | except |  |  |  |
| 3:9.90 | wickedness |  |  |  |
| 3:9.91 | of heart |  |  |  |
| 3:9.92 | that you |  |  |  |
| 3:9.93 | enjoy |  |  |  |
| 3:9.94 | and long for |  |  |  |
| 3:9.95 | their gossip |  |  |  |
| 3:9.96 | and their informing |  |  |  |
| 3:9.97 | the Merciful One |  |  |  |
| 3:9.98 | save us |  |  |  |
| 3:9.99 | from this |  |  |  |
| 3:9.100 | opinion |  |  |  |
| 3:9.101 | and it is possible |  |  |  |
| 3:9.102 | that also |  |  |  |
| 3:9.103 | you |  |  |  |
| 3:9.104 | add |  |  |  |
| 3:9.105 | of your own |  |  |  |
| 3:9.106 | informing |  |  |  |
| 3:9.107 | false |  |  |  |
| 3:9.108 | in order |  |  |  |
| 3:9.109 | to honor yourself |  |  |  |
| 3:9.110 | at the disgrace |  |  |  |
| 3:9.111 | of your friends |  |  |  |
| 3:9.112 | and in the chapter |  |  |  |
| 3:9.113 | Arvei |  |  |  |
| 3:9.114 | Pesachim |  |  |  |
| 3:9.115 | (folio 115) |  |  | folio 115 of Pesachim, the reference as the author gives it (קי"ה) |
| 3:9.116 | we read |  |  |  |
| 3:9.117 | said |  |  |  |
| 3:9.118 | Rav |  |  |  |
| 3:9.119 | Sheshet |  |  |  |
| 3:9.120 | in the name of |  |  |  |
| 3:9.121 | R. A. |  |  | ראב"ע: not expanded in the printed text; read as "R. A."; the full name is uncertain |
| 3:9.122 | whoever |  |  |  |
| 3:9.123 | tells |  |  |  |
| 3:9.124 | evil |  |  |  |
| 3:9.125 | talk |  |  |  |
| 3:9.126 | and whoever testifies |  |  |  |
| 3:9.127 | testimony |  |  |  |
| 3:9.128 | false |  |  |  |
| 3:9.129 | against his fellow |  |  |  |
| 3:9.130 | is fit |  |  |  |
| 3:9.131 | to cast him |  |  |  |
| 3:9.132 | to the dogs |  |  |  |
| 3:9.133 | as it is said |  |  |  |
| 3:9.134 | to the dog |  |  |  |
| 3:9.135 | you shall throw |  |  |  |
| 3:9.136 | it |  |  |  |
| 3:9.137 | and it is written |  |  |  |
| 3:9.138 | explicitly |  |  |  |
| 3:9.139 | not |  |  |  |
| 3:9.140 | you shall carry |  |  |  |
| 3:9.141 | report |  |  |  |
| 3:9.142 | false |  |  |  |
| 3:9.143 | and it is read |  |  |  |
| 3:9.144 | in it |  |  |  |
| 3:9.145 | not |  |  |  |
| 3:9.146 | you shall cause to bear |  |  | תשיא: the reading the Talmud gives ("וקרי ביה"), set against the written לא תשא of Exod 23:1 |
| 3:9.147 | and further |  |  |  |
| 3:9.148 | several |  |  |  |
| 3:9.149 | homilies |  |  |  |
| 3:9.150 | we find |  |  |  |
| 3:9.151 | our Sages of blessed memory |  |  | לרז"ל: "our Sages of blessed memory" |
| 3:9.152 | on the punishment |  |  |  |
| 3:9.153 | of the one who tells |  |  |  |
| 3:9.154 | evil |  |  |  |
| 3:9.155 | talk |  |  |  |
| 3:9.156 | and in Mahzor |  |  |  |
| 3:9.157 | Vitry |  |  |  |
| 3:9.158 | tractate |  |  |  |
| 3:9.159 | Avot |  |  |  |
| 3:9.160 | on |  |  |  |
| 3:9.161 | that |  |  |  |
| 3:9.162 | which we learned |  |  |  |
| 3:9.163 | on (al: do not) |  |  | על: the Mishnah's "al" (אל), the negative of "do not judge" |
| 3:9.164 | judge |  |  |  |
| 3:9.165 | - |  |  |  |
| 3:9.166 | your fellow |  |  |  |
| 3:9.167 | until |  |  |  |
| 3:9.168 | you reach |  |  |  |
| 3:9.169 | his place |  |  |  |
| 3:9.170 | as he writes |  |  | כ' ע"ש: "as he writes there"; the expansion is uncertain |
| 3:9.171 | there |  |  |  |
| 3:9.172 | that may come |  |  |  |
| 3:9.173 | into your hand |  |  |  |
| 3:9.174 | that same |  |  |  |
| 3:9.175 | transgression |  |  |  |
| 3:9.176 | and be saved |  |  |  |
| 3:9.177 | and escape |  |  |  |
| 3:9.178 | from it |  |  |  |
| 3:9.179 | as |  |  |  |
| 3:9.180 | we find |  |  |  |
| 3:9.181 | with Jeroboam |  |  |  |
| 3:9.182 | son of |  |  |  |
| 3:9.183 | Nebat |  |  |  |
| 3:9.184 | that when |  |  |  |
| 3:9.185 | he married |  |  |  |
| 3:9.186 | Solomon |  |  |  |
| 3:9.187 | - |  |  |  |
| 3:9.188 | daughter |  |  |  |
| 3:9.189 | of Pharaoh |  |  |  |
| 3:9.190 | she brought in |  |  |  |
| 3:9.191 | to him |  |  |  |
| 3:9.192 | every |  |  |  |
| 3:9.193 | kind of |  |  |  |
| 3:9.194 | song |  |  |  |
| 3:9.195 | in the world |  |  |  |
| 3:9.196 | and she said |  |  |  |
| 3:9.197 | to him |  |  |  |
| 3:9.198 | thus |  |  |  |
| 3:9.199 | they do |  |  |  |
| 3:9.200 | for idolatry |  |  | לע"ז: abbreviation of עבודה זרה, "idolatry" |
| 3:9.201 | such-and-such |  |  |  |
| 3:9.202 | once |  |  |  |
| 3:9.203 | time |  |  |  |
| 3:9.204 | he went |  |  |  |
| 3:9.205 | to sleep |  |  |  |
| 3:9.206 | and became late |  |  |  |
| 3:9.207 | the day |  |  |  |
| 3:9.208 | she went |  |  |  |
| 3:9.209 | and called |  |  |  |
| 3:9.210 | the wood-carvers |  |  |  |
| 3:9.211 | of wood |  |  |  |
| 3:9.212 | and the builders |  |  |  |
| 3:9.213 | who do |  |  |  |
| 3:9.214 | the craft |  |  |  |
| 3:9.215 | of carving |  |  |  |
| 3:9.216 | and weaving |  |  |  |
| 3:9.217 | and they made |  |  |  |
| 3:9.218 | on |  |  |  |
| 3:9.219 | his bed |  |  |  |
| 3:9.220 | the likeness of |  |  |  |
| 3:9.221 | a firmament |  |  |  |
| 3:9.222 | and in it |  |  |  |
| 3:9.223 | the sun |  |  |  |
| 3:9.224 | and the moon |  |  |  |
| 3:9.225 | and the stars |  |  |  |
| 3:9.226 | and the constellations |  |  |  |
| 3:9.227 | and in the morning |  |  |  |
| 3:9.228 | he woke |  |  |  |
| 3:9.229 | from his sleep |  |  |  |
| 3:9.230 | to open |  |  |  |
| 3:9.231 | the doors |  |  |  |
| 3:9.232 | of the Temple |  |  |  |
| 3:9.233 | for there were |  |  |  |
| 3:9.234 | keys |  |  |  |
| 3:9.235 | of the Temple |  |  |  |
| 3:9.236 | under |  |  |  |
| 3:9.237 | his pillow |  |  |  |
| 3:9.238 | he saw |  |  |  |
| 3:9.239 | that the firmament |  |  |  |
| 3:9.240 | was bright |  |  |  |
| 3:9.241 | with stars |  |  |  |
| 3:9.242 | and the host |  |  |  |
| 3:9.243 | thinking |  |  |  |
| 3:9.244 | still |  |  |  |
| 3:9.245 | not |  |  |  |
| 3:9.246 | had turned |  |  |  |
| 3:9.247 | night |  |  |  |
| 3:9.248 | he returned |  |  |  |
| 3:9.249 | to sleep |  |  |  |
| 3:9.250 | and he slept |  |  |  |
| 3:9.251 | more |  |  |  |
| 3:9.252 | four |  |  |  |
| 3:9.253 | hours |  |  |  |
| 3:9.254 | until |  |  |  |
| 3:9.255 | he understood |  |  |  |
| 3:9.256 | and recognised |  |  |  |
| 3:9.257 | for himself |  |  |  |
| 3:9.258 | that there is not |  |  |  |
| 3:9.259 | the night |  |  |  |
| 3:9.260 | long |  |  |  |
| 3:9.261 | so |  |  |  |
| 3:9.262 | thus |  |  |  |
| 3:9.263 | and concerning |  |  |  |
| 3:9.264 | that |  |  |  |
| 3:9.265 | hour |  |  |  |
| 3:9.266 | we learned |  |  |  |
| 3:9.267 | testimony |  |  |  |
| 3:9.268 | about |  |  |  |
| 3:9.269 | the daily offering |  |  |  |
| 3:9.270 | of |  |  |  |
| 3:9.271 | morning |  |  |  |
| 3:9.272 | which is offered |  |  |  |
| 3:9.273 | at four |  |  |  |
| 3:9.274 | hours |  |  |  |
| 3:9.275 | in the Bechirta |  |  | בבחירתא: transliterated; the source is not named in the pack |
| 3:9.276 | and before |  |  |  |
| 3:9.277 | he woke |  |  |  |
| 3:9.278 | from his sleep |  |  |  |
| 3:9.279 | went |  |  |  |
| 3:9.280 | Jeroboam |  |  |  |
| 3:9.281 | son of |  |  |  |
| 3:9.282 | Nebat |  |  |  |
| 3:9.283 | to |  |  |  |
| 3:9.284 | the door |  |  |  |
| 3:9.285 | of his house |  |  |  |
| 3:9.286 | and went out |  |  |  |
| 3:9.287 | a voice |  |  |  |
| 3:9.288 | wicked one |  |  |  |
| 3:9.289 | until |  |  |  |
| 3:9.290 | when |  |  |  |
| 3:9.291 | you sleep |  |  |  |
| 3:9.292 | and neglect |  |  |  |
| 3:9.293 | the daily offering |  |  |  |
| 3:9.294 | of |  |  |  |
| 3:9.295 | morning |  |  |  |
| 3:9.296 | what |  |  |  |
| 3:9.297 | to you |  |  |  |
| 3:9.298 | sleeper |  |  |  |
| 3:9.299 | arise |  |  |  |
| 3:9.300 | call |  |  |  |
| 3:9.301 | upon |  |  |  |
| 3:9.302 | your God |  |  | אלהיך: God of Israel, in the voice's words |
| 3:9.303 | came out |  |  |  |
| 3:9.304 | bat (a heavenly) |  |  |  |
| 3:9.305 | voice |  |  |  |
| 3:9.306 | and said |  |  |  |
| 3:9.307 | to him |  |  |  |
| 3:9.308 | wicked |  |  |  |
| 3:9.309 | son of |  |  |  |
| 3:9.310 | wicked |  |  |  |
| 3:9.311 | destined |  |  |  |
| 3:9.312 | you |  |  |  |
| 3:9.313 | to nullify |  |  |  |
| 3:9.314 | from Israel |  |  |  |
| 3:9.315 | many |  |  |  |
| 3:9.316 | and many |  |  |  |
| 3:9.317 | offerings |  |  |  |
| 3:9.318 | wilfully |  |  |  |
| 3:9.319 | and you |  |  |  |
| 3:9.320 | hold him liable |  |  | מחייבו: the pointing does not fix the subject and object; both are kept as in the Hebrew |
| 3:9.321 | for |  |  |  |
| 3:9.322 | the inadvertent |  |  |  |
| 3:9.323 | as the wilful |  |  |  |
| 3:9.324 | and for |  |  |  |
| 3:9.325 | the forced |  |  |  |
| 3:9.326 | as the voluntary |  |  |  |
| 3:9.327 | and this is |  |  |  |
| 3:9.328 | what is written |  |  |  |
| 3:9.329 | as the speaking of |  |  |  |
| 3:9.330 | Ephraim |  |  |  |
| 3:9.331 | trembling |  |  |  |
| 3:9.332 | he was exalted |  |  |  |
| 3:9.333 | he |  |  |  |
| 3:9.334 | in Israel |  |  |  |
| 3:9.335 | and became guilty |  |  |  |
| 3:9.336 | through Baal |  |  |  |
| 3:9.337 | and died |  |  |  |
| 3:9.338 | in the matter of |  |  |  |
| 3:9.339 | Jeroboam |  |  |  |
| 3:9.340 | who is |  |  |  |
| 3:9.341 | of the tribe of |  |  |  |
| 3:9.342 | Ephraim |  |  |  |
| 3:9.343 | rattat (trembled) |  |  | רתת: "trembled", the verse's word (Hos 13:1), read here by the author as the name of the accuser |
| 3:9.344 | katigorin (accuser) |  |  | קטיגורין: an accuser (the Talmudic term) |
| 3:9.345 | against |  |  |  |
| 3:9.346 | Solomon |  |  |  |
| 3:9.347 | he |  |  |  |
| 3:9.348 | Jeroboam |  |  |  |
| 3:9.349 | prince |  |  |  |
| 3:9.350 | in Israel |  |  |  |
| 3:9.351 | and king |  |  |  |
| 3:9.352 | and a bat |  |  |  |
| 3:9.353 | kol |  |  |  |
| 3:9.354 | answered him |  |  |  |
| 3:9.355 | that he would become guilty |  |  |  |
| 3:9.356 | through Baal |  |  |  |
| 3:9.357 | and die |  |  |  |
| 3:9.358 | thus far his words |  |  |  |
| 3:9.359 | and you |  |  |  |
| 3:9.360 | have magnified |  |  |  |
| 3:9.361 | and added |  |  |  |
| 3:9.362 | to do |  |  |  |
| 3:9.363 | evil |  |  |  |
| 3:9.364 | more |  |  |  |
| 3:9.365 | than Jeroboam |  |  |  |
| 3:9.366 | son of |  |  |  |
| 3:9.367 | Nebat |  |  |  |
| 3:9.368 | for |  |  |  |
| 3:9.369 | after |  |  |  |
| 3:9.370 | that |  |  |  |
| 3:9.371 | you brought |  |  |  |
| 3:9.372 | us |  |  |  |
| 3:9.373 | to |  |  |  |
| 3:9.374 | the house of |  |  |  |
| 3:9.375 | the prison |  |  |  |
| 3:9.376 | by |  |  |  |
| 3:9.377 | false informing |  |  |  |
| 3:9.378 | before |  |  |  |
| 3:9.379 | our government |  |  |  |
| 3:9.380 | exalted |  |  |  |
| 3:9.381 | you called out |  |  |  |
| 3:9.382 | - |  |  |  |
| 3:9.383 | all |  |  |  |
| 3:9.384 | the people |  |  |  |
| 3:9.385 | to go out |  |  |  |
| 3:9.386 | to |  |  |  |
| 3:9.387 | the house of |  |  |  |
| 3:9.388 | the cemetery |  |  |  |
| 3:9.389 | and you commanded |  |  |  |
| 3:9.390 | them |  |  |  |
| 3:9.391 | to dust themselves |  |  |  |
| 3:9.392 | with dust |  |  |  |
| 3:9.393 | on |  |  |  |
| 3:9.394 | their heads |  |  |  |
| 3:9.395 | for |  |  |  |
| 3:9.396 | no |  |  |  |
| 3:9.397 | thing |  |  |  |
| 3:9.398 | not |  |  |  |
| 3:9.399 | for |  |  |  |
| 3:9.400 | a commandment |  |  |  |
| 3:9.401 | light |  |  |  |
| 3:9.402 | that we neglected |  |  |  |
| 3:9.403 | and not |  |  |  |
| 3:9.404 | for |  |  |  |
| 3:9.405 | a commandment |  |  |  |
| 3:9.406 | weighty |  |  |  |
| 3:9.407 | God forbid |  |  | ח"ו: "God forbid" |
| 3:9.408 | but |  |  |  |
| 3:9.409 | only |  |  |  |
| 3:9.410 | for |  |  |  |
| 3:9.411 | that |  |  |  |
| 3:9.412 | we walked |  |  |  |
| 3:9.413 | in the custom of |  |  |  |
| 3:9.414 | our forefathers |  |  |  |
| 3:9.415 | and we set |  |  |  |
| 3:9.416 | our learning |  |  |  |
| 3:9.417 | in the Mishnah |  |  |  |
| 3:9.418 | and in the Talmud |  |  |  |
| 3:9.419 | and the Rambam |  |  | ורמב"ם: the Rambam |
| 3:9.420 | and Shulchan Aruch |  |  | וש"ע: the Shulchan Aruch |
| 3:9.421 | to learn |  |  |  |
| 3:9.422 | and to teach |  |  |  |
| 3:9.423 | to keep |  |  |  |
| 3:9.424 | and to do |  |  |  |
| 3:9.425 | and not |  |  |  |
| 3:9.426 | we studied |  |  |  |
| 3:9.427 | further |  |  |  |
| 3:9.428 | from the book |  |  |  |
| 3:9.429 | forged |  |  | מזויף: "forged", the author's epithet for the Zohar |
| 3:9.430 | the book |  |  |  |
| 3:9.431 | the Zohar |  |  | הזהר: "the book the Zohar", the title kept as the author writes it |
| 3:9.432 | as is the custom |  |  |  |
| 3:9.433 | of you |  |  |  |

Phrases:
- 3:9.1 | 9) | 9)
- 3:9.2 3:9.3 3:9.4 3:9.5 3:9.6 3:9.7 3:9.8 3:9.9 3:9.10 | And you, our friend, it was not enough for you that you did not do any one of all these | and you our friend not enough for you that not you did one of all these
- 3:9.11 3:9.12 3:9.13 3:9.14 3:9.15 | but you believed their words and accepted their lies | but that you believed to their words and you accepted their lies
- 3:9.16 3:9.17 3:9.18 3:9.19 | on the testimony of one witness and one judge | on the testimony of one and a judge one
- 3:9.20 3:9.21 3:9.22 3:9.23 | and not in the presence of the litigant | and not in the presence of the litigant of the case
- 3:9.24 3:9.25 3:9.26 3:9.27 3:9.28 3:9.29 | without inquiry, without investigation, without examination | and without inquiry and without investigation and without examination
- 3:9.30 3:9.31 | and you enjoyed their words | and you enjoyed in their words
- 3:9.32 3:9.33 3:9.34 3:9.35 | and you also helped to strengthen their lies | and also you helped to strengthen their lies
- 3:9.36 3:9.37 3:9.38 3:9.39 | and it was fixed in your mind as the truth itself | and was fixed in your mind as truth to its truth
- 3:9.40 3:9.41 3:9.42 3:9.43 3:9.44 3:9.45 3:9.46 | until you wrote that I am no longer able to deny it, etc. | until that you wrote that I am able to deny any longer etc.
- 3:9.47 3:9.48 3:9.49 3:9.50 3:9.51 3:9.52 3:9.53 3:9.54 3:9.55 3:9.56 3:9.57 | And you did not take to heart that testimony is accepted only in the presence of the litigant | and not you put to your heart that there is not they accept testimony but only in the presence of a litigant of the case
- 3:9.58 3:9.59 3:9.60 3:9.61 3:9.62 3:9.63 3:9.64 3:9.65 3:9.66 3:9.67 3:9.68 | and that you are obliged to inquire into them with seven inquiries, investigations and examinations, as you are commanded | and that you are obliged to inquire and to investigate them with seven inquiries and with investigations and with examinations as commanded of you
- 3:9.69 3:9.70 3:9.71 3:9.72 3:9.73 | All the more so, since they did their deeds in the dark | all the more so that they did in the dark their deeds
- 3:9.74 3:9.75 3:9.76 3:9.77 3:9.78 | and smoothed their tongues to you in secret with their words | and smoothed their tongue to you in secret with their words
- 3:9.79 3:9.80 3:9.81 | to find a fault to hate | to find a fault to hate
- 3:9.82 3:9.83 3:9.84 3:9.85 | and you were obliged to silence them with a rebuke. | and you were obliged to silence them with a rebuke
- 3:9.86 3:9.87 3:9.88 3:9.89 3:9.90 3:9.91 3:9.92 3:9.93 3:9.94 3:9.95 3:9.96 | This is nothing but a wicked heart: that you enjoy and long for their gossip and their informing. | it is not this that except wickedness of heart that you enjoy and long for their gossip and their informing
- 3:9.97 3:9.98 3:9.99 3:9.100 | the Merciful One save us from this opinion! | the Merciful One save us from this opinion
- 3:9.101 3:9.102 3:9.103 3:9.104 3:9.105 3:9.106 3:9.107 3:9.108 3:9.109 3:9.110 3:9.111 | And it is possible that you too add of your own false informing, in order to honor yourself at the disgrace of your friends. | and it is possible that also you add of your own informing false in order to honor yourself at the disgrace of your friends
- 3:9.112 3:9.113 3:9.114 3:9.115 3:9.116 | And in the chapter Arvei Pesachim (115) we read: | and in the chapter Arvei Pesachim (folio 115) we read
- 3:9.117 3:9.118 3:9.119 3:9.120 3:9.121 3:9.122 3:9.123 3:9.124 3:9.125 | "Rav Sheshet said in the name of R. A.: whoever tells evil talk | said Rav Sheshet in the name of R. A. whoever tells evil talk
- 3:9.126 3:9.127 3:9.128 3:9.129 3:9.130 3:9.131 3:9.132 3:9.133 | and whoever bears false witness against his fellow deserves to be thrown to the dogs, as it is said: | and whoever testifies testimony false against his fellow is fit to cast him to the dogs as it is said
- 3:9.134 3:9.135 3:9.136 | 'to the dog you shall throw it' | to the dog you shall throw it
- 3:9.137 3:9.138 3:9.139 3:9.140 3:9.141 3:9.142 3:9.143 3:9.144 3:9.145 3:9.146 | and it is written explicitly: 'you shall not carry a false report'; and it is read in it: 'you shall not cause to bear'" | and it is written explicitly not you shall carry report false and it is read in it not you shall cause to bear
- 3:9.147 3:9.148 3:9.149 3:9.150 3:9.151 3:9.152 3:9.153 3:9.154 3:9.155 | And we find several more homilies of our Sages of blessed memory on the punishment of one who tells evil talk. | and further several homilies we find our Sages of blessed memory on the punishment of the one who tells evil talk
- 3:9.156 3:9.157 3:9.158 3:9.159 3:9.160 3:9.161 3:9.162 | "And in Mahzor Vitry, tractate Avot, on what we learned: | and in Mahzor Vitry tractate Avot on that which we learned
- 3:9.163 3:9.164 3:9.165 3:9.166 3:9.167 3:9.168 3:9.169 | 'do not judge your fellow until you reach his place' | on (al: do not) judge - your fellow until you reach his place
- 3:9.170 3:9.171 3:9.172 3:9.173 3:9.174 3:9.175 | as he writes there: that the same transgression may come into your hand | as he writes there that may come into your hand that same transgression
- 3:9.176 3:9.177 3:9.178 | and you may be saved and escape from it | and be saved and escape from it
- 3:9.179 3:9.180 3:9.181 3:9.182 3:9.183 | as we find with Jeroboam son of Nebat | as we find with Jeroboam son of Nebat
- 3:9.184 3:9.185 3:9.186 3:9.187 3:9.188 3:9.189 | that when Solomon married the daughter of Pharaoh, | that when he married Solomon - daughter of Pharaoh
- 3:9.190 3:9.191 3:9.192 3:9.193 3:9.194 3:9.195 | she brought in to him every kind of song in the world | she brought in to him every kind of song in the world
- 3:9.196 3:9.197 3:9.198 3:9.199 3:9.200 3:9.201 | and said to him: 'So they do for such-and-such idolatry' | and she said to him thus they do for idolatry such-and-such
- 3:9.202 3:9.203 3:9.204 3:9.205 3:9.206 3:9.207 | One time, he went to sleep, and the day grew late | once time he went to sleep and became late the day
- 3:9.208 3:9.209 3:9.210 3:9.211 3:9.212 3:9.213 3:9.214 3:9.215 3:9.216 | she went and called the wood-carvers and the builders, who work at carving and weaving | she went and called the wood-carvers of wood and the builders who do the craft of carving and weaving
- 3:9.217 3:9.218 3:9.219 3:9.220 3:9.221 3:9.222 3:9.223 3:9.224 3:9.225 3:9.226 | and they made over his bed the likeness of a firmament, and in it the sun, the moon, the stars and the constellations | and they made on his bed the likeness of a firmament and in it the sun and the moon and the stars and the constellations
- 3:9.227 3:9.228 3:9.229 3:9.230 3:9.231 3:9.232 | In the morning he woke from his sleep to open the doors of the Temple, | and in the morning he woke from his sleep to open the doors of the Temple
- 3:9.233 3:9.234 3:9.235 3:9.236 3:9.237 | for the keys of the Temple were under his pillow | for there were keys of the Temple under his pillow
- 3:9.238 3:9.239 3:9.240 3:9.241 3:9.242 | he saw that the firmament was bright with stars and the host | he saw that the firmament was bright with stars and the host
- 3:9.243 3:9.244 3:9.245 3:9.246 3:9.247 | thinking that the night had not yet turned | thinking still not had turned night
- 3:9.248 3:9.249 | he went back to sleep | he returned to sleep
- 3:9.250 3:9.251 3:9.252 3:9.253 | And he slept four more hours, | and he slept more four hours
- 3:9.254 3:9.255 3:9.256 3:9.257 3:9.258 3:9.259 3:9.260 3:9.261 3:9.262 | until he understood and recognised for himself that the night was not so long | until he understood and recognised for himself that there is not the night long so thus
- 3:9.263 3:9.264 3:9.265 3:9.266 3:9.267 | And concerning that hour we learned a testimony | and concerning that hour we learned testimony
- 3:9.268 3:9.269 3:9.270 3:9.271 3:9.272 3:9.273 3:9.274 3:9.275 | about the morning daily offering, which is offered at four hours, in the Bechirta. | about the daily offering of morning which is offered at four hours in the Bechirta
- 3:9.276 3:9.277 3:9.278 3:9.279 3:9.280 3:9.281 3:9.282 | Before he woke from his sleep, Jeroboam son of Nebat went | and before he woke from his sleep went Jeroboam son of Nebat
- 3:9.283 3:9.284 3:9.285 | to the door of his house | to the door of his house
- 3:9.286 3:9.287 3:9.288 | and a voice went out: 'Wicked one | and went out a voice wicked one
- 3:9.289 3:9.290 3:9.291 3:9.292 3:9.293 3:9.294 3:9.295 | how long will you sleep and neglect the morning daily offering | until when you sleep and neglect the daily offering of morning
- 3:9.296 3:9.297 3:9.298 | What is it to you, sleeper? | what to you sleeper
- 3:9.299 3:9.300 3:9.301 3:9.302 | Arise, call upon your God!' | arise call upon your God
- 3:9.303 3:9.304 3:9.305 3:9.306 3:9.307 | A bat kol came out and said to him: | came out bat (a heavenly) voice and said to him
- 3:9.308 3:9.309 3:9.310 3:9.311 3:9.312 3:9.313 3:9.314 3:9.315 3:9.316 3:9.317 3:9.318 | 'Wicked son of a wicked man, you are destined to abolish from Israel many, many offerings wilfully. | wicked son of wicked destined you to nullify from Israel many and many offerings wilfully
- 3:9.319 3:9.320 3:9.321 3:9.322 3:9.323 | And you hold him liable for the inadvertent as for the wilful. | and you hold him liable for the inadvertent as the wilful
- 3:9.324 3:9.325 3:9.326 | And for the forced one as for the voluntary!' | and for the forced as the voluntary
- 3:9.327 3:9.328 3:9.329 3:9.330 | And this is what is written: 'When Ephraim spoke | and this is what is written as the speaking of Ephraim
- 3:9.331 3:9.332 3:9.333 3:9.334 3:9.335 3:9.336 3:9.337 | there was trembling; he was exalted in Israel, but he incurred guilt through Baal and died' | trembling he was exalted he in Israel and became guilty through Baal and died
- 3:9.338 3:9.339 3:9.340 3:9.341 3:9.342 | This is said of Jeroboam, who was of the tribe of Ephraim, | in the matter of Jeroboam who is of the tribe of Ephraim
- 3:9.343 3:9.344 3:9.345 3:9.346 | the trembling was the accuser against Solomon, | rattat (trembled) katigorin (accuser) against Solomon
- 3:9.347 3:9.348 3:9.349 3:9.350 3:9.351 | he, Jeroboam, was prince in Israel and king | he Jeroboam prince in Israel and king
- 3:9.352 3:9.353 3:9.354 3:9.355 3:9.356 3:9.357 3:9.358 | and the bat kol answered him that he would become guilty through Baal and die." Thus far his words. | and a bat kol answered him that he would become guilty through Baal and die thus far his words
- 3:9.359 3:9.360 3:9.361 3:9.362 3:9.363 3:9.364 3:9.365 3:9.366 3:9.367 | And you have magnified and added to doing evil more than Jeroboam son of Nebat. | and you have magnified and added to do evil more than Jeroboam son of Nebat
- 3:9.368 3:9.369 3:9.370 3:9.371 3:9.372 3:9.373 3:9.374 3:9.375 | For after you brought us to the prison | for after that you brought us to the house of the prison
- 3:9.376 3:9.377 3:9.378 3:9.379 3:9.380 | by false informing before our exalted government! | by false informing before our government exalted
- 3:9.381 3:9.382 3:9.383 3:9.384 3:9.385 3:9.386 3:9.387 3:9.388 | You called out all the people to go out to the cemetery, | you called out - all the people to go out to the house of the cemetery
- 3:9.389 3:9.390 3:9.391 3:9.392 3:9.393 3:9.394 | and commanded them to dust themselves with dust upon their heads | and you commanded them to dust themselves with dust on their heads
- 3:9.395 3:9.396 3:9.397 | for no reason. | for no thing
- 3:9.398 3:9.399 3:9.400 3:9.401 3:9.402 | Not for a light commandment that we neglected, | not for a commandment light that we neglected
- 3:9.403 3:9.404 3:9.405 3:9.406 3:9.407 | and not for a weighty commandment, God forbid, | and not for a commandment weighty God forbid
- 3:9.408 3:9.409 3:9.410 3:9.411 3:9.412 | but because we walked | but only for that we walked
- 3:9.413 3:9.414 3:9.415 3:9.416 | in the custom of our forefathers and set our learning | in the custom of our forefathers and we set our learning
- 3:9.417 3:9.418 3:9.419 3:9.420 | in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch | in the Mishnah and in the Talmud and the Rambam and Shulchan Aruch
- 3:9.421 3:9.422 3:9.423 3:9.424 | to learn and to teach, to keep and to do. | to learn and to teach to keep and to do
- 3:9.425 3:9.426 3:9.427 3:9.428 3:9.429 3:9.430 3:9.431 3:9.432 3:9.433 | And we did not study further from the forged book, the book the Zohar, as is your custom. | and not we studied further from the book forged the book the Zohar as is the custom of you

Grammar notes:
- לא דייך שלא עשית (3:9.5-7): an idiom, "it was not enough for you that you did not do". The irony turns on the reversal: the charge is not only the omission but the belief and acceptance that follow ("but", אלא).
- בעד אחד ודיין אחד, ושלא בפני בעל דין, ובלא חקירה ובלא דרישה, ובלא בדיקה (3:9.16-29): the grounds of acceptance are listed in order, each negative with its own ובלא; the English keeps the repetition.
- The Pesachim passage, after גרסינן, is Talmudic Hebrew with its own formulas (אמר רב ששת, כתיב בהדיא, וקרי ביה). It cites Exod 22:30 and Exod 23:1; the Talmud's reading (לא תשיא) is set against the written form (לא תשא).
- The Mahzor Vitry passage runs from על הא דתנן to עכ"ל (thus far his words). The bat kol's speech sits inside it in single marks, and the verse of Hos 13:1 is its proof. The pronoun in ואתה מחייבו is left open, as the Hebrew leaves it.
- The last part turns from the story to the addressee: ואתה הגדלת (and you have magnified) and כי אחרי אשר (for after) make the charge that his informing brought the author to prison.

Translator's note: One paragraph (Wikisource, 1931 print). The printed list marker 9) is kept; the Hebrew abbreviations are expanded in the token glosses.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated, so the reading is flagged here until it is (Q2 default).
- apparatus: Exod 22:30 (לכלב תשליכון אותו), quoted as written: "to the dog you shall throw it".
- apparatus: Exod 23:1 (לא תשא שמע שוא), quoted as written; the Talmud's reading לא תשיא is given as read (וקרי).
- apparatus: Hos 13:1 (כדבר אפרים רתת נשא הוא בישראל ויאשם בבעל וימות), quoted by the author; translated from the Hebrew.
- apparatus: The Mishnah quoted is named only as tractate Avot; the author writes על for אל. No chapter reference is added.
- apparatus: Pesachim: the author gives folio 115 only.
- apparatus: ראב"ע and כ' ע"ש are not expanded with certainty; the English gives the likeliest sense.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

