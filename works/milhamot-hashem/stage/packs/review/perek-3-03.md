# Review pack: milhamot-hashem perek-3-03 (Milchamot ha-Shem 3:5–3:6)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-03`.

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
- A Biblical quotation whose wording differs from the verse follows the
  author's wording; the difference goes in `apparatus`.
- `literal` on a long paragraph: the glosses in Hebrew order, hyphens joining
  the parts of one word. No reasoning in it; reasons go in `note`.
- `context` may carry a short note on a story the author cites (the Pardes
  story), without a reference unless the author gives one.

## Rulings from the pilot reviews, by the editor

- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- גופי תורה: "the essentials of the Torah", with "bodies of the Torah" in a
  token note.
- אלהים / אלהי- and the case of "God": the English does not editorialise by
  case. "God" where the Hebrew names the God of Israel, also inside the
  Kabbalists' quoted words (their "our God"); "god(s)" only where the Hebrew
  itself marks foreignness or plurality (אלהים אחרים, אל זר, אלילים) or where
  the author calls the Kabbalists' object of worship a god in his own voice
  (his charge, in his words). In a doubtful case use "God" and put the
  alternative in a token note.
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
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- דברי תימה: "astonishing words" (words that astonish).
- כלי: "vessel" only for the Kabbalistic כלים; "tool" for a craftsman's object.
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
- אורה זו תורה (the exegetic formula, Esth 8:16): "'orah' means Torah".
- "World to Come" in the author's mouth: "world to come" (as in Avot); the Hebrew
  for בחיי alone is "Bachya".
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


## Milchamot ha-Shem 3:5

Hebrew (Wikisource): ובפעם השלישית היתה תשובתכם כשאלת אחד מארבעה בנים מה העבודה הזאת לכם, לכם ולא לו, כאילו לא נצטוינו בתורה וידעת היום והשבות אל לבבך כי ה' הוא האלהים וכו' וכמה מקראות באו בתורה ובנביאים ובכתובים על זה וכל שכן בדברי רז"ל. גם חזרת להיות מכת מספרי לשון הרע וממקבליו באמרך שכבר יש עלינו כמה עדים שאנחנו מבזים את החכמים ולא נוכל להכחיש עוד וכו'. והן ידוע ומפורסם שאין מקבלין עדות אלא בפני בעל דין. ומן התורה אחד דיני ממונות ואחד דיני נפשות בדרישה וחקירה כמו שכ' הרמב"ם ז"ל בפרק ג' מהלכות עדות אבל אמרו חכמים כדי שלא תנעול דלת בפני לוין אין עדי ממון צריכים דרישה וחקירה וכו'. במה דברים אמורים בהודאות והלואות, מתנות ומכירות, אבל בדיני קנסות צריכים דרישה וחקירה, ואין צריך לומר במלקיות ובגלות, ובדין מרומה אף בדיני ממונות צריך דרישה וחקירה.

| id | word | punctuation after |
|---|---|---|
| 3:5.1 | ובפעם |  |
| 3:5.2 | השלישית |  |
| 3:5.3 | היתה |  |
| 3:5.4 | תשובתכם |  |
| 3:5.5 | כשאלת |  |
| 3:5.6 | אחד |  |
| 3:5.7 | מארבעה |  |
| 3:5.8 | בנים |  |
| 3:5.9 | מה |  |
| 3:5.10 | העבודה |  |
| 3:5.11 | הזאת |  |
| 3:5.12 | לכם | , |
| 3:5.13 | לכם |  |
| 3:5.14 | ולא |  |
| 3:5.15 | לו | , |
| 3:5.16 | כאילו |  |
| 3:5.17 | לא |  |
| 3:5.18 | נצטוינו |  |
| 3:5.19 | בתורה |  |
| 3:5.20 | וידעת |  |
| 3:5.21 | היום |  |
| 3:5.22 | והשבות |  |
| 3:5.23 | אל |  |
| 3:5.24 | לבבך |  |
| 3:5.25 | כי |  |
| 3:5.26 | ה | ' |
| 3:5.27 | הוא |  |
| 3:5.28 | האלהים |  |
| 3:5.29 | וכו | ' |
| 3:5.30 | וכמה |  |
| 3:5.31 | מקראות |  |
| 3:5.32 | באו |  |
| 3:5.33 | בתורה |  |
| 3:5.34 | ובנביאים |  |
| 3:5.35 | ובכתובים |  |
| 3:5.36 | על |  |
| 3:5.37 | זה |  |
| 3:5.38 | וכל |  |
| 3:5.39 | שכן |  |
| 3:5.40 | בדברי |  |
| 3:5.41 | רז"ל | . |
| 3:5.42 | גם |  |
| 3:5.43 | חזרת |  |
| 3:5.44 | להיות |  |
| 3:5.45 | מכת |  |
| 3:5.46 | מספרי |  |
| 3:5.47 | לשון |  |
| 3:5.48 | הרע |  |
| 3:5.49 | וממקבליו |  |
| 3:5.50 | באמרך |  |
| 3:5.51 | שכבר |  |
| 3:5.52 | יש |  |
| 3:5.53 | עלינו |  |
| 3:5.54 | כמה |  |
| 3:5.55 | עדים |  |
| 3:5.56 | שאנחנו |  |
| 3:5.57 | מבזים |  |
| 3:5.58 | את |  |
| 3:5.59 | החכמים |  |
| 3:5.60 | ולא |  |
| 3:5.61 | נוכל |  |
| 3:5.62 | להכחיש |  |
| 3:5.63 | עוד |  |
| 3:5.64 | וכו | '. |
| 3:5.65 | והן |  |
| 3:5.66 | ידוע |  |
| 3:5.67 | ומפורסם |  |
| 3:5.68 | שאין |  |
| 3:5.69 | מקבלין |  |
| 3:5.70 | עדות |  |
| 3:5.71 | אלא |  |
| 3:5.72 | בפני |  |
| 3:5.73 | בעל |  |
| 3:5.74 | דין | . |
| 3:5.75 | ומן |  |
| 3:5.76 | התורה |  |
| 3:5.77 | אחד |  |
| 3:5.78 | דיני |  |
| 3:5.79 | ממונות |  |
| 3:5.80 | ואחד |  |
| 3:5.81 | דיני |  |
| 3:5.82 | נפשות |  |
| 3:5.83 | בדרישה |  |
| 3:5.84 | וחקירה |  |
| 3:5.85 | כמו |  |
| 3:5.86 | שכ | ' |
| 3:5.87 | הרמב"ם |  |
| 3:5.88 | ז"ל |  |
| 3:5.89 | בפרק |  |
| 3:5.90 | ג | ' |
| 3:5.91 | מהלכות |  |
| 3:5.92 | עדות |  |
| 3:5.93 | אבל |  |
| 3:5.94 | אמרו |  |
| 3:5.95 | חכמים |  |
| 3:5.96 | כדי |  |
| 3:5.97 | שלא |  |
| 3:5.98 | תנעול |  |
| 3:5.99 | דלת |  |
| 3:5.100 | בפני |  |
| 3:5.101 | לוין |  |
| 3:5.102 | אין |  |
| 3:5.103 | עדי |  |
| 3:5.104 | ממון |  |
| 3:5.105 | צריכים |  |
| 3:5.106 | דרישה |  |
| 3:5.107 | וחקירה |  |
| 3:5.108 | וכו | '. |
| 3:5.109 | במה |  |
| 3:5.110 | דברים |  |
| 3:5.111 | אמורים |  |
| 3:5.112 | בהודאות |  |
| 3:5.113 | והלואות | , |
| 3:5.114 | מתנות |  |
| 3:5.115 | ומכירות | , |
| 3:5.116 | אבל |  |
| 3:5.117 | בדיני |  |
| 3:5.118 | קנסות |  |
| 3:5.119 | צריכים |  |
| 3:5.120 | דרישה |  |
| 3:5.121 | וחקירה | , |
| 3:5.122 | ואין |  |
| 3:5.123 | צריך |  |
| 3:5.124 | לומר |  |
| 3:5.125 | במלקיות |  |
| 3:5.126 | ובגלות | , |
| 3:5.127 | ובדין |  |
| 3:5.128 | מרומה |  |
| 3:5.129 | אף |  |
| 3:5.130 | בדיני |  |
| 3:5.131 | ממונות |  |
| 3:5.132 | צריך |  |
| 3:5.133 | דרישה |  |
| 3:5.134 | וחקירה | . |


### Translation under review (3:5)

English: And for the third time your answer was like the question of one of the four sons: "What is this service to you? To you and not to him." As if we had not been commanded in the Torah, "And know this day and take it to your heart, that the LORD, He is God, etc." And many verses came in the Torah, in the Prophets and in the Writings on this, and all the more so in the words of our Rabbis of blessed memory. You have also returned to be of the company of tellers of evil talk and of those who receive it, when you say that already there are several witnesses against us, that we despise the Sages and can no longer deny it, etc. And it is known and notorious that testimony is accepted only in the presence of the party to the suit. From the Torah, one kind of civil case and one kind of capital case require inquiry and examination, as the Rambam of blessed memory wrote in chapter 3 of the Laws of Testimony. But the Sages said, "So that you do not lock the door before borrowers, monetary witnesses do not require inquiry and examination, etc." In what matters is this said? In confessions and loans, gifts and sales. But in fines, inquiry and examination are required, and there is no need to say in lashes and exile, and in a case of deception, even in civil cases, inquiry and examination are required.

Literal: and for the third time was your reply like the question of one of the four sons what the service this to you to you and not to him as if not we were commanded in the Torah and know this day and take to your heart that the LORD He is God etc. And many verses came in the Torah and in the Prophets and in the Writings on this and all the more so in the words of our Rabbis of blessed memory also you have returned to be of the company of the tellers of evil talk and of those who receive it when you say that already there are against us several witnesses that we despise (object marker) the Sages and cannot deny any longer etc. And indeed it is known and notorious that there is no accept testimony except in the presence of the party to the suit And from the Torah one of the laws of civil matters and one of the laws of capital cases in inquiry and examination as wrote the Rambam of blessed memory in chapter 3 of the Laws of Testimony But said the Sages so that not you lock the door before borrowers there is no witnesses of money need inquiry and examination etc. In what matters are said in confessions and loans gifts and sales But in the cases of fines require inquiry and examination and there is no need to say in lashes and in exile and in a case of deception even in the cases of civil requires inquiry and examination

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:5.1 | and for the |  |  |  |
| 3:5.2 | third time |  |  |  |
| 3:5.3 | was |  |  |  |
| 3:5.4 | your reply |  |  |  |
| 3:5.5 | like the question of |  |  |  |
| 3:5.6 | one |  |  |  |
| 3:5.7 | of the four |  |  |  |
| 3:5.8 | sons |  |  |  |
| 3:5.9 | what |  |  | What: the wicked son's question in the Passover Haggadah (Exod 12:26). |
| 3:5.10 | the service |  |  |  |
| 3:5.11 | this |  |  |  |
| 3:5.12 | to you |  |  | לכם (to you) is repeated as in the Haggadah's wording; the dative of interest. |
| 3:5.13 | to you |  |  |  |
| 3:5.14 | and not |  |  |  |
| 3:5.15 | to him |  |  |  |
| 3:5.16 | as if |  |  |  |
| 3:5.17 | not |  |  |  |
| 3:5.18 | we were commanded |  |  |  |
| 3:5.19 | in the Torah |  |  |  |
| 3:5.20 | and know |  |  |  |
| 3:5.21 | this day |  |  |  |
| 3:5.22 | and take |  |  | והשבות אל לבבך is the idiom 'take it to heart' (the Hiphil of שוב, 'bring back, return to'). |
| 3:5.23 | to |  |  |  |
| 3:5.24 | your heart |  |  |  |
| 3:5.25 | that |  |  |  |
| 3:5.26 | the LORD |  |  |  |
| 3:5.27 | He |  |  |  |
| 3:5.28 | is God |  |  |  |
| 3:5.29 | etc. |  |  | וכו' marks a truncated verse; only the words on the page are rendered, and the verse goes on (see apparatus). |
| 3:5.30 | And many |  |  |  |
| 3:5.31 | verses |  |  |  |
| 3:5.32 | came |  |  |  |
| 3:5.33 | in the Torah |  |  |  |
| 3:5.34 | and in the Prophets |  |  |  |
| 3:5.35 | and in the Writings |  |  |  |
| 3:5.36 | on |  |  |  |
| 3:5.37 | this |  |  |  |
| 3:5.38 | and all |  |  |  |
| 3:5.39 | the more so |  |  |  |
| 3:5.40 | in the words of |  |  |  |
| 3:5.41 | our Rabbis of blessed memory |  |  | רז"ל: our Rabbis of blessed memory (honorific). |
| 3:5.42 | also |  |  |  |
| 3:5.43 | you have returned |  |  |  |
| 3:5.44 | to be |  |  |  |
| 3:5.45 | of the company |  |  | מכת: 'company, band'; the author speaks ironically of the company of tellers of evil talk. |
| 3:5.46 | of the tellers |  |  |  |
| 3:5.47 | of |  |  |  |
| 3:5.48 | evil talk |  |  |  |
| 3:5.49 | and of those who receive it |  |  |  |
| 3:5.50 | when you say |  |  |  |
| 3:5.51 | that already |  |  |  |
| 3:5.52 | there are |  |  |  |
| 3:5.53 | against us |  |  |  |
| 3:5.54 | several |  |  |  |
| 3:5.55 | witnesses |  |  |  |
| 3:5.56 | that we |  |  |  |
| 3:5.57 | despise |  |  |  |
| 3:5.58 | (object marker) |  |  |  |
| 3:5.59 | the Sages |  |  |  |
| 3:5.60 | and |  |  |  |
| 3:5.61 | cannot |  |  |  |
| 3:5.62 | deny |  |  |  |
| 3:5.63 | any longer |  |  |  |
| 3:5.64 | etc. |  |  |  |
| 3:5.65 | And indeed |  |  | והן: 'and indeed', introducing a statement presented as known. |
| 3:5.66 | it is known |  |  |  |
| 3:5.67 | and notorious |  |  |  |
| 3:5.68 | that there is no |  |  |  |
| 3:5.69 | accept |  |  |  |
| 3:5.70 | testimony |  |  |  |
| 3:5.71 | except |  |  |  |
| 3:5.72 | in the presence of |  |  |  |
| 3:5.73 | the party |  |  |  |
| 3:5.74 | to the suit |  |  |  |
| 3:5.75 | And from |  |  |  |
| 3:5.76 | the Torah |  |  |  |
| 3:5.77 | one |  |  |  |
| 3:5.78 | of the laws of |  |  |  |
| 3:5.79 | civil matters |  |  |  |
| 3:5.80 | and one |  |  |  |
| 3:5.81 | of the laws of |  |  |  |
| 3:5.82 | capital cases |  |  |  |
| 3:5.83 | in inquiry |  |  |  |
| 3:5.84 | and examination |  |  |  |
| 3:5.85 | as |  |  |  |
| 3:5.86 | wrote |  |  | שכ' is שכתב, 'wrote'. |
| 3:5.87 | the Rambam |  |  |  |
| 3:5.88 | of blessed memory |  |  |  |
| 3:5.89 | in chapter |  |  |  |
| 3:5.90 | 3 |  |  | ג' is the numeral 3 (chapter 3). |
| 3:5.91 | of the Laws of |  |  | Mishneh Torah (the Rambam's code), Laws of Testimony (Hilkhot Edut), chapter 3; the one work named in this unit. |
| 3:5.92 | Testimony |  |  |  |
| 3:5.93 | But |  |  |  |
| 3:5.94 | said |  |  |  |
| 3:5.95 | the Sages |  |  |  |
| 3:5.96 | so that |  |  |  |
| 3:5.97 | not |  |  |  |
| 3:5.98 | you lock |  |  |  |
| 3:5.99 | the door |  |  |  |
| 3:5.100 | before |  |  |  |
| 3:5.101 | borrowers |  |  | לוין: borrowers (plural of לוה). |
| 3:5.102 | there is no |  |  |  |
| 3:5.103 | witnesses |  |  |  |
| 3:5.104 | of money |  |  | עדי ממון: witnesses in monetary matters. |
| 3:5.105 | need |  |  |  |
| 3:5.106 | inquiry |  |  |  |
| 3:5.107 | and examination |  |  |  |
| 3:5.108 | etc. |  |  |  |
| 3:5.109 | In what |  |  |  |
| 3:5.110 | matters |  |  |  |
| 3:5.111 | are said |  |  |  |
| 3:5.112 | in confessions |  |  |  |
| 3:5.113 | and loans |  |  |  |
| 3:5.114 | gifts |  |  |  |
| 3:5.115 | and sales |  |  |  |
| 3:5.116 | But |  |  |  |
| 3:5.117 | in the cases of |  |  |  |
| 3:5.118 | fines |  |  |  |
| 3:5.119 | require |  |  |  |
| 3:5.120 | inquiry |  |  |  |
| 3:5.121 | and examination |  |  |  |
| 3:5.122 | and there is no |  |  |  |
| 3:5.123 | need |  |  |  |
| 3:5.124 | to say |  |  |  |
| 3:5.125 | in lashes |  |  |  |
| 3:5.126 | and in exile |  |  |  |
| 3:5.127 | and in a case of |  |  | ובדין מרומה: 'and in a case of deception'. |
| 3:5.128 | deception |  |  | מרומה: a case of deception or fraud, a technical term in the Talmud; the author gives no further sense. |
| 3:5.129 | even |  |  |  |
| 3:5.130 | in the cases of |  |  |  |
| 3:5.131 | civil |  |  |  |
| 3:5.132 | requires |  |  |  |
| 3:5.133 | inquiry |  |  |  |
| 3:5.134 | and examination |  |  |  |

Phrases:
- 3:5.1 3:5.2 3:5.3 3:5.4 3:5.5 3:5.6 3:5.7 3:5.8 | And for the third time your answer was like the question of one of the four sons | and for the third time was your reply like the question of one of the four sons
- 3:5.9 3:5.10 3:5.11 3:5.12 3:5.13 | What is this service to you? To you | what the service this to you to you
- 3:5.14 3:5.15 | and not to him | and not to him
- 3:5.16 3:5.17 3:5.18 3:5.19 | As if we had not been commanded in the Torah | as if not we were commanded in the Torah
- 3:5.20 3:5.21 3:5.22 3:5.23 3:5.24 3:5.25 3:5.26 3:5.27 3:5.28 3:5.29 | And know this day and take it to your heart, that the LORD, He is God, etc | and know this day and take to your heart that the LORD He is God etc.
- 3:5.30 3:5.31 3:5.32 3:5.33 3:5.34 3:5.35 3:5.36 3:5.37 3:5.38 3:5.39 3:5.40 3:5.41 | And many verses came in the Torah, in the Prophets and in the Writings on this, and all the more so in the words of our Rabbis of blessed memory | And many verses came in the Torah and in the Prophets and in the Writings on this and all the more so in the words of our Rabbis of blessed memory
- 3:5.42 3:5.43 3:5.44 3:5.45 3:5.46 3:5.47 3:5.48 3:5.49 3:5.50 | You have also returned to be of the company of tellers of evil talk and of those who receive it, when you say | also you have returned to be of the company of the tellers of evil talk and of those who receive it when you say
- 3:5.51 3:5.52 3:5.53 3:5.54 3:5.55 3:5.56 3:5.57 3:5.58 3:5.59 3:5.60 3:5.61 3:5.62 3:5.63 3:5.64 | that already there are several witnesses against us, that we despise the Sages and can no longer deny it, etc | that already there are against us several witnesses that we despise (object marker) the Sages and cannot deny any longer etc.
- 3:5.65 3:5.66 3:5.67 3:5.68 3:5.69 3:5.70 3:5.71 3:5.72 3:5.73 3:5.74 | And it is known and notorious that testimony is accepted only in the presence of the party to the suit | And indeed it is known and notorious that there is no accept testimony except in the presence of the party to the suit
- 3:5.75 3:5.76 3:5.77 3:5.78 3:5.79 3:5.80 3:5.81 3:5.82 3:5.83 3:5.84 | From the Torah, one kind of civil case and one kind of capital case require inquiry and examination | And from the Torah one of the laws of civil matters and one of the laws of capital cases in inquiry and examination
- 3:5.85 3:5.86 3:5.87 3:5.88 3:5.89 3:5.90 3:5.91 3:5.92 | as the Rambam of blessed memory wrote in chapter 3 of the Laws of Testimony | as wrote the Rambam of blessed memory in chapter 3 of the Laws of Testimony
- 3:5.93 3:5.94 3:5.95 3:5.96 3:5.97 3:5.98 3:5.99 3:5.100 3:5.101 3:5.102 3:5.103 3:5.104 3:5.105 3:5.106 3:5.107 | But the Sages said, "So that you do not lock the door before borrowers, monetary witnesses do not require inquiry and examination | But said the Sages so that not you lock the door before borrowers there is no witnesses of money need inquiry and examination
- 3:5.108 | etc. | etc.
- 3:5.109 3:5.110 3:5.111 3:5.112 3:5.113 3:5.114 3:5.115 | In what matters is this said? In confessions and loans, gifts and sales | In what matters are said in confessions and loans gifts and sales
- 3:5.116 3:5.117 3:5.118 3:5.119 3:5.120 3:5.121 | But in fines, inquiry and examination are required | But in the cases of fines require inquiry and examination
- 3:5.122 3:5.123 3:5.124 3:5.125 3:5.126 | and there is no need to say in lashes and exile | and there is no need to say in lashes and in exile
- 3:5.127 3:5.128 3:5.129 3:5.130 3:5.131 3:5.132 3:5.133 3:5.134 | and in a case of deception, even in civil cases, inquiry and examination are required | and in a case of deception even in the cases of civil requires inquiry and examination

Grammar notes:
- The Haggadah's wicked son asks 'What is this service to you?' (מה העבודה הזאת לכם, Exod 12:26); the author quotes the words, which need no quotation marks as an allusion.
- In 'וידעת היום והשבות אל לבבך' the verbs are second-person forms of Deut 4:39; 'כי ה' הוא האלהים' has no copula in Hebrew, so 'is God' supplies it.
- 'ומן התורה' opens a sentence with no finite verb; 'require' is supplied from 'צריכים' in the sentence below, and from the sense of 'בדרישה וחקירה'.
- 'אמרו חכמים' introduces a saying of the Sages with no source given by the author; no reference is added.
- 'במה דברים אמורים' is an idiom: 'in what case is this said?'

Commentary rulings:
- apparatus: Exod 12:26: 'מה העבודה הזאת לכם' (the wicked son of the Passover Haggadah).
- apparatus: Deut 4:39: 'וידעת היום והשבות אל לבבך כי ה' הוא האלהים'; the author stops at וכו', and the verse continues beyond the words quoted.
- apparatus: Wikisource base text (Q2): translated as it stands; flagged until the 1931 scan is collated.

## Milchamot ha-Shem 3:6

Hebrew (Wikisource): ו) ובנדון דידן כולהו איתינהו ביה, ואיך מלאך לבך לקבל לשון הרע מאנשים סכלים ההולכים יחפים כמנודים לשמים, ולחשוב אותה לעדות גמורה, להחזיק את עמיתך בתורה ובמצות דאוריתא ודרבנן לאפיקורסים ומינים שהם מכלל אותם שמורידין ולא מעלין? וידוע הוא ומפורסם, כי זה הדור, דור שוא ושקר, והיא בנשיאי ישראל, וכל אחד מתכבד בקלון חברו ורבו אשר העמידו בקרו אורה זו תורה. ובמשימי לילות כימים ועבדי לתורה שלא תשתכח מישראל כפי יכולתם שפכתם בוז, ותשימו את לומדי התורה ומחזיקים בה נבזים ושפלים ככלי אין חפץ בו. וכל עושי רע טוב בעיניכם, ומכובד יותר מהם, ועל ידי דבות שקר שמתם לומדי תורה לשמצה בעיני כל רואיהם עד שבאו למדרגה פחותה מאד, שעמי הארץ משתבחים בעצמם לומר אשרינו מה טוב חלקנו וערבה שנתינו עד עלות השמש וכבוד ננחל יותר מאלה העומדים בבית ה' בלילות, לקיים מה שנאמר ואם בחקותי תמאסו, אם במלמדי חקותי תמאסו, כמו שכ הרמב"ם ז"ל בהל' תלמוד תורה.

| id | word | punctuation after |
|---|---|---|
| 3:6.1 | ו | ) |
| 3:6.2 | ובנדון |  |
| 3:6.3 | דידן |  |
| 3:6.4 | כולהו |  |
| 3:6.5 | איתינהו |  |
| 3:6.6 | ביה | , |
| 3:6.7 | ואיך |  |
| 3:6.8 | מלאך |  |
| 3:6.9 | לבך |  |
| 3:6.10 | לקבל |  |
| 3:6.11 | לשון |  |
| 3:6.12 | הרע |  |
| 3:6.13 | מאנשים |  |
| 3:6.14 | סכלים |  |
| 3:6.15 | ההולכים |  |
| 3:6.16 | יחפים |  |
| 3:6.17 | כמנודים |  |
| 3:6.18 | לשמים | , |
| 3:6.19 | ולחשוב |  |
| 3:6.20 | אותה |  |
| 3:6.21 | לעדות |  |
| 3:6.22 | גמורה | , |
| 3:6.23 | להחזיק |  |
| 3:6.24 | את |  |
| 3:6.25 | עמיתך |  |
| 3:6.26 | בתורה |  |
| 3:6.27 | ובמצות |  |
| 3:6.28 | דאוריתא |  |
| 3:6.29 | ודרבנן |  |
| 3:6.30 | לאפיקורסים |  |
| 3:6.31 | ומינים |  |
| 3:6.32 | שהם |  |
| 3:6.33 | מכלל |  |
| 3:6.34 | אותם |  |
| 3:6.35 | שמורידין |  |
| 3:6.36 | ולא |  |
| 3:6.37 | מעלין | ? |
| 3:6.38 | וידוע |  |
| 3:6.39 | הוא |  |
| 3:6.40 | ומפורסם | , |
| 3:6.41 | כי |  |
| 3:6.42 | זה |  |
| 3:6.43 | הדור | , |
| 3:6.44 | דור |  |
| 3:6.45 | שוא |  |
| 3:6.46 | ושקר | , |
| 3:6.47 | והיא |  |
| 3:6.48 | בנשיאי |  |
| 3:6.49 | ישראל | , |
| 3:6.50 | וכל |  |
| 3:6.51 | אחד |  |
| 3:6.52 | מתכבד |  |
| 3:6.53 | בקלון |  |
| 3:6.54 | חברו |  |
| 3:6.55 | ורבו |  |
| 3:6.56 | אשר |  |
| 3:6.57 | העמידו |  |
| 3:6.58 | בקרו |  |
| 3:6.59 | אורה |  |
| 3:6.60 | זו |  |
| 3:6.61 | תורה | . |
| 3:6.62 | ובמשימי |  |
| 3:6.63 | לילות |  |
| 3:6.64 | כימים |  |
| 3:6.65 | ועבדי |  |
| 3:6.66 | לתורה |  |
| 3:6.67 | שלא |  |
| 3:6.68 | תשתכח |  |
| 3:6.69 | מישראל |  |
| 3:6.70 | כפי |  |
| 3:6.71 | יכולתם |  |
| 3:6.72 | שפכתם |  |
| 3:6.73 | בוז | , |
| 3:6.74 | ותשימו |  |
| 3:6.75 | את |  |
| 3:6.76 | לומדי |  |
| 3:6.77 | התורה |  |
| 3:6.78 | ומחזיקים |  |
| 3:6.79 | בה |  |
| 3:6.80 | נבזים |  |
| 3:6.81 | ושפלים |  |
| 3:6.82 | ככלי |  |
| 3:6.83 | אין |  |
| 3:6.84 | חפץ |  |
| 3:6.85 | בו | . |
| 3:6.86 | וכל |  |
| 3:6.87 | עושי |  |
| 3:6.88 | רע |  |
| 3:6.89 | טוב |  |
| 3:6.90 | בעיניכם | , |
| 3:6.91 | ומכובד |  |
| 3:6.92 | יותר |  |
| 3:6.93 | מהם | , |
| 3:6.94 | ועל |  |
| 3:6.95 | ידי |  |
| 3:6.96 | דבות |  |
| 3:6.97 | שקר |  |
| 3:6.98 | שמתם |  |
| 3:6.99 | לומדי |  |
| 3:6.100 | תורה |  |
| 3:6.101 | לשמצה |  |
| 3:6.102 | בעיני |  |
| 3:6.103 | כל |  |
| 3:6.104 | רואיהם |  |
| 3:6.105 | עד |  |
| 3:6.106 | שבאו |  |
| 3:6.107 | למדרגה |  |
| 3:6.108 | פחותה |  |
| 3:6.109 | מאד | , |
| 3:6.110 | שעמי |  |
| 3:6.111 | הארץ |  |
| 3:6.112 | משתבחים |  |
| 3:6.113 | בעצמם |  |
| 3:6.114 | לומר |  |
| 3:6.115 | אשרינו |  |
| 3:6.116 | מה |  |
| 3:6.117 | טוב |  |
| 3:6.118 | חלקנו |  |
| 3:6.119 | וערבה |  |
| 3:6.120 | שנתינו |  |
| 3:6.121 | עד |  |
| 3:6.122 | עלות |  |
| 3:6.123 | השמש |  |
| 3:6.124 | וכבוד |  |
| 3:6.125 | ננחל |  |
| 3:6.126 | יותר |  |
| 3:6.127 | מאלה |  |
| 3:6.128 | העומדים |  |
| 3:6.129 | בבית |  |
| 3:6.130 | ה | ' |
| 3:6.131 | בלילות | , |
| 3:6.132 | לקיים |  |
| 3:6.133 | מה |  |
| 3:6.134 | שנאמר |  |
| 3:6.135 | ואם |  |
| 3:6.136 | בחקותי |  |
| 3:6.137 | תמאסו | , |
| 3:6.138 | אם |  |
| 3:6.139 | במלמדי |  |
| 3:6.140 | חקותי |  |
| 3:6.141 | תמאסו | , |
| 3:6.142 | כמו |  |
| 3:6.143 | שכ |  |
| 3:6.144 | הרמב"ם |  |
| 3:6.145 | ז"ל |  |
| 3:6.146 | בהל | ' |
| 3:6.147 | תלמוד |  |
| 3:6.148 | תורה | . |


### Translation under review (3:6)

English: 6) And in our case all of them are in it, and how can your heart bring itself to accept evil talk from foolish men who walk barefoot, like outcasts by heaven, and to consider it complete testimony, to hold your fellow in the Torah and in the commandments of Scriptural and rabbinic law as apikorsim and heretics, who are among those whom one lowers and does not raise? And it is known and notorious that this generation, a generation of vanity and falsehood, and it is among the princes of Israel, and each one takes pride in the disgrace of his fellow and of his teacher, whom they set up in their examination; this light is the Torah. And those who make nights as days, and those who devote themselves to the Torah, so that it not be forgotten from Israel according to their ability, you poured out contempt, and you made the students of the Torah and those who hold to it despised and lowly, like a vessel for which there is no delight in it. And all who do evil are good in your eyes, and more honored than them, and by means of false rumors you made the students of Torah a disgrace in the eyes of all who see them, until they came to a very low rank, that the common folk boast of themselves, saying, "Happy are we, how good is our portion and pleasant our sleep until the rising of the sun, and we shall inherit honor more than those who stand in the house of the LORD at night"; to fulfill what was said, "If you despise My statutes, if you despise those who teach My statutes," as the Rambam of blessed memory wrote in the Laws of the Study of Torah.

Literal: 6) and in the case of our all of them are in it and how can bring itself your heart to accept the talk of evil from men foolish who walk barefoot like outcasts by heaven and to consider it as complete testimony to hold (object marker) your fellow in the Torah and in the commandments of Scriptural law and rabbinic as apikorsim and heretics who are among those whom one lowers and not raise And it is known and notorious that this generation a generation of vanity and falsehood and it is among the princes of Israel and each one takes pride in the disgrace of his fellow and his teacher whom they set up in their examination light this Torah And those who make nights as days and those who devote themselves to the Torah so that not be forgotten from Israel according to their ability you poured out contempt and you made (object marker) the students of the Torah and those who hold to it despised and lowly like a vessel for which there is no delight in it And all those who do evil good in your eyes and honored more than them and by means of rumors of falsehood you made the students of Torah a disgrace in the eyes of all who see them until they came to a rank low very that the common folk boast of themselves saying Happy are we how good our portion and pleasant our sleep until the rising of the sun and honor we shall inherit more than those who stand in the house of the LORD at night to fulfill what was said If in My statutes you despise if in those who teach My statutes you despise as wrote the Rambam of blessed memory in the Laws of the study of Torah

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:6.1 | 6) |  |  | ו) is the printed list marker (the sixth point), rendered as the author's own '6)'. |
| 3:6.2 | and in the case of |  |  |  |
| 3:6.3 | our |  |  |  |
| 3:6.4 | all of them |  |  |  |
| 3:6.5 | are |  |  |  |
| 3:6.6 | in it |  |  |  |
| 3:6.7 | and how |  |  |  |
| 3:6.8 | can bring itself |  |  | מלאך (here 'bring itself', from the sense 'fill'): the heart is the subject of a rhetorical question. |
| 3:6.9 | your heart |  |  |  |
| 3:6.10 | to accept |  |  |  |
| 3:6.11 | the talk of |  |  |  |
| 3:6.12 | evil |  |  |  |
| 3:6.13 | from men |  |  |  |
| 3:6.14 | foolish |  |  |  |
| 3:6.15 | who walk |  |  |  |
| 3:6.16 | barefoot |  |  |  |
| 3:6.17 | like outcasts |  |  |  |
| 3:6.18 | by heaven |  |  |  |
| 3:6.19 | and to consider |  |  |  |
| 3:6.20 | it |  |  |  |
| 3:6.21 | as |  |  |  |
| 3:6.22 | complete testimony |  |  |  |
| 3:6.23 | to hold |  |  |  |
| 3:6.24 | (object marker) |  |  |  |
| 3:6.25 | your fellow |  |  |  |
| 3:6.26 | in the Torah |  |  |  |
| 3:6.27 | and in the commandments |  |  |  |
| 3:6.28 | of Scriptural law |  |  | דאוריתא: Aramaic 'de-oraita', of the Torah (Scriptural law); ודרבנן: 'and rabbinic' (de-rabbanan). |
| 3:6.29 | and rabbinic |  |  |  |
| 3:6.30 | as apikorsim |  |  | אפיקורסים: transliterated (fixed term 'apikorsim'). |
| 3:6.31 | and heretics |  |  |  |
| 3:6.32 | who are |  |  |  |
| 3:6.33 | among |  |  |  |
| 3:6.34 | those |  |  |  |
| 3:6.35 | whom one lowers |  |  |  |
| 3:6.36 | and not |  |  |  |
| 3:6.37 | raise |  |  | The sentence ends with a question mark in the Hebrew. |
| 3:6.38 | And it is |  |  |  |
| 3:6.39 | known |  |  |  |
| 3:6.40 | and notorious |  |  |  |
| 3:6.41 | that |  |  |  |
| 3:6.42 | this |  |  |  |
| 3:6.43 | generation |  |  |  |
| 3:6.44 | a generation of |  |  |  |
| 3:6.45 | vanity |  |  |  |
| 3:6.46 | and falsehood |  |  |  |
| 3:6.47 | and it is |  |  |  |
| 3:6.48 | among the princes of |  |  |  |
| 3:6.49 | Israel |  |  |  |
| 3:6.50 | and each |  |  |  |
| 3:6.51 | one |  |  |  |
| 3:6.52 | takes pride |  |  |  |
| 3:6.53 | in the disgrace of |  |  |  |
| 3:6.54 | his fellow |  |  |  |
| 3:6.55 | and his teacher |  |  |  |
| 3:6.56 | whom |  |  |  |
| 3:6.57 | they set up |  |  |  |
| 3:6.58 | in their examination |  |  | בקרו is obscure; read here as 'in their examination' (from בקר). The sense is uncertain. |
| 3:6.59 | light |  |  | אורה: 'light'. The phrase 'אורה זו תורה' is an allusive saying: 'this light is the Torah'. |
| 3:6.60 | this |  |  |  |
| 3:6.61 | Torah |  |  |  |
| 3:6.62 | And those who make |  |  | ובמשימי: participle 'those who make' with לילות כימים as its object: 'nights as days'. |
| 3:6.63 | nights |  |  |  |
| 3:6.64 | as days |  |  |  |
| 3:6.65 | and those who devote themselves |  |  |  |
| 3:6.66 | to the Torah |  |  |  |
| 3:6.67 | so that not |  |  |  |
| 3:6.68 | be forgotten |  |  |  |
| 3:6.69 | from Israel |  |  |  |
| 3:6.70 | according to |  |  |  |
| 3:6.71 | their ability |  |  |  |
| 3:6.72 | you poured out |  |  |  |
| 3:6.73 | contempt |  |  |  |
| 3:6.74 | and you made |  |  |  |
| 3:6.75 | (object marker) |  |  |  |
| 3:6.76 | the students of |  |  |  |
| 3:6.77 | the Torah |  |  |  |
| 3:6.78 | and those who hold |  |  |  |
| 3:6.79 | to it |  |  |  |
| 3:6.80 | despised |  |  |  |
| 3:6.81 | and lowly |  |  |  |
| 3:6.82 | like a vessel |  |  |  |
| 3:6.83 | for which there is |  |  |  |
| 3:6.84 | no delight |  |  |  |
| 3:6.85 | in it |  |  |  |
| 3:6.86 | And all |  |  |  |
| 3:6.87 | those who do |  |  |  |
| 3:6.88 | evil |  |  |  |
| 3:6.89 | good |  |  |  |
| 3:6.90 | in your eyes |  |  |  |
| 3:6.91 | and honored |  |  |  |
| 3:6.92 | more |  |  |  |
| 3:6.93 | than them |  |  |  |
| 3:6.94 | and by |  |  |  |
| 3:6.95 | means of |  |  |  |
| 3:6.96 | rumors |  |  |  |
| 3:6.97 | of falsehood |  |  |  |
| 3:6.98 | you made |  |  |  |
| 3:6.99 | the students of |  |  |  |
| 3:6.100 | Torah |  |  |  |
| 3:6.101 | a disgrace |  |  |  |
| 3:6.102 | in the eyes of |  |  |  |
| 3:6.103 | all |  |  |  |
| 3:6.104 | who see them |  |  |  |
| 3:6.105 | until |  |  |  |
| 3:6.106 | they came |  |  |  |
| 3:6.107 | to a rank |  |  |  |
| 3:6.108 | low |  |  |  |
| 3:6.109 | very |  |  |  |
| 3:6.110 | that the |  |  |  |
| 3:6.111 | common folk |  |  |  |
| 3:6.112 | boast |  |  |  |
| 3:6.113 | of themselves |  |  |  |
| 3:6.114 | saying |  |  |  |
| 3:6.115 | Happy are we |  |  |  |
| 3:6.116 | how |  |  |  |
| 3:6.117 | good |  |  |  |
| 3:6.118 | our portion |  |  |  |
| 3:6.119 | and pleasant |  |  |  |
| 3:6.120 | our sleep |  |  |  |
| 3:6.121 | until |  |  |  |
| 3:6.122 | the rising of |  |  |  |
| 3:6.123 | the sun |  |  |  |
| 3:6.124 | and honor |  |  |  |
| 3:6.125 | we shall inherit |  |  |  |
| 3:6.126 | more |  |  |  |
| 3:6.127 | than those |  |  |  |
| 3:6.128 | who stand |  |  |  |
| 3:6.129 | in the house of |  |  |  |
| 3:6.130 | the LORD |  |  | ה': the Tetragrammaton written as ה' (the LORD), as in Ps 134:1. |
| 3:6.131 | at night |  |  |  |
| 3:6.132 | to fulfill |  |  |  |
| 3:6.133 | what |  |  |  |
| 3:6.134 | was said |  |  |  |
| 3:6.135 | If |  |  |  |
| 3:6.136 | in My statutes |  |  |  |
| 3:6.137 | you despise |  |  |  |
| 3:6.138 | if |  |  |  |
| 3:6.139 | in those who teach |  |  |  |
| 3:6.140 | My statutes |  |  |  |
| 3:6.141 | you despise |  |  | תמאסו: 'you despise', second person plural of the verb of Lev 26:15; the author writes בחקותי. |
| 3:6.142 | as |  |  |  |
| 3:6.143 | wrote |  |  |  |
| 3:6.144 | the Rambam |  |  |  |
| 3:6.145 | of blessed memory |  |  |  |
| 3:6.146 | in the Laws of |  |  | בהל': Mishneh Torah, Laws of Torah Study (Hilkhot Talmud Torah), the same work as in 3:5.91. |
| 3:6.147 | the study of |  |  |  |
| 3:6.148 | Torah |  |  |  |

Phrases:
- 3:6.1 | 6) | 6)
- 3:6.2 3:6.3 3:6.4 3:6.5 3:6.6 | And in our case all of them are in it | and in the case of our all of them are in it
- 3:6.7 3:6.8 3:6.9 3:6.10 | and how can your heart bring itself to accept | and how can bring itself your heart to accept
- 3:6.11 3:6.12 3:6.13 3:6.14 3:6.15 3:6.16 3:6.17 3:6.18 | evil talk from foolish men who walk barefoot, like outcasts by heaven | the talk of evil from men foolish who walk barefoot like outcasts by heaven
- 3:6.19 3:6.20 3:6.21 3:6.22 | and to consider it complete testimony | and to consider it as complete testimony
- 3:6.23 3:6.24 3:6.25 3:6.26 3:6.27 3:6.28 3:6.29 3:6.30 3:6.31 | to hold your fellow in the Torah and in the commandments of Scriptural and rabbinic law as apikorsim and heretics | to hold (object marker) your fellow in the Torah and in the commandments of Scriptural law and rabbinic as apikorsim and heretics
- 3:6.32 3:6.33 3:6.34 3:6.35 3:6.36 3:6.37 | who are among those whom one lowers and does not raise | who are among those whom one lowers and not raise
- 3:6.38 3:6.39 3:6.40 | And it is known and notorious | And it is known and notorious
- 3:6.41 3:6.42 3:6.43 3:6.44 3:6.45 3:6.46 | that this generation, a generation of vanity and falsehood | that this generation a generation of vanity and falsehood
- 3:6.47 3:6.48 3:6.49 | and it is among the princes of Israel | and it is among the princes of Israel
- 3:6.50 3:6.51 3:6.52 3:6.53 3:6.54 3:6.55 3:6.56 3:6.57 3:6.58 3:6.59 3:6.60 3:6.61 | and each one takes pride in the disgrace of his fellow and of his teacher, whom they set up in their examination; this light is the Torah | and each one takes pride in the disgrace of his fellow and his teacher whom they set up in their examination light this Torah
- 3:6.62 3:6.63 3:6.64 3:6.65 3:6.66 | And those who make nights as days, and those who devote themselves to the Torah | And those who make nights as days and those who devote themselves to the Torah
- 3:6.67 3:6.68 3:6.69 3:6.70 3:6.71 | so that it not be forgotten from Israel according to their ability | so that not be forgotten from Israel according to their ability
- 3:6.72 3:6.73 | you poured out contempt | you poured out contempt
- 3:6.74 3:6.75 3:6.76 3:6.77 3:6.78 3:6.79 3:6.80 3:6.81 3:6.82 3:6.83 3:6.84 3:6.85 | and you made the students of the Torah and those who hold to it despised and lowly, like a vessel for which there is no delight in it | and you made (object marker) the students of the Torah and those who hold to it despised and lowly like a vessel for which there is no delight in it
- 3:6.86 3:6.87 3:6.88 3:6.89 3:6.90 3:6.91 3:6.92 3:6.93 | And all who do evil are good in your eyes, and more honored than them | And all those who do evil good in your eyes and honored more than them
- 3:6.94 3:6.95 3:6.96 3:6.97 3:6.98 3:6.99 3:6.100 | and by means of false rumors you made the students of Torah | and by means of rumors of falsehood you made the students of Torah
- 3:6.101 3:6.102 3:6.103 3:6.104 | a disgrace in the eyes of all who see them | a disgrace in the eyes of all who see them
- 3:6.105 3:6.106 3:6.107 3:6.108 3:6.109 | until they came to a very low rank | until they came to a rank low very
- 3:6.110 3:6.111 3:6.112 3:6.113 3:6.114 | that the common folk boast of themselves, saying | that the common folk boast of themselves saying
- 3:6.115 3:6.116 3:6.117 3:6.118 3:6.119 3:6.120 3:6.121 3:6.122 3:6.123 | Happy are we, how good is our portion and pleasant our sleep until the rising of the sun | Happy are we how good our portion and pleasant our sleep until the rising of the sun
- 3:6.124 3:6.125 3:6.126 3:6.127 3:6.128 3:6.129 | and we shall inherit honor more than those who stand in the house of | and honor we shall inherit more than those who stand in the house of
- 3:6.130 3:6.131 | the LORD at night | the LORD at night
- 3:6.132 3:6.133 3:6.134 | to fulfill what was said | to fulfill what was said
- 3:6.135 3:6.136 3:6.137 3:6.138 3:6.139 3:6.140 3:6.141 | If you despise My statutes, if you despise those who teach My statutes | If in My statutes you despise if in those who teach My statutes you despise
- 3:6.142 3:6.143 3:6.144 3:6.145 3:6.146 3:6.147 3:6.148 | as the Rambam of blessed memory wrote in the Laws of the Study of Torah | as wrote the Rambam of blessed memory in the Laws of the study of Torah

Grammar notes:
- The paragraph is one long sentence of questions, statements and lists, with its only question marked '?' after 'ולא מעלין'.
- 'ואיך מלאך לבך לקבל' is rhetorical: 'how can your heart bring itself to accept'; the verb 'מלאך' is reflexive in sense.
- 'לאפיקורסים ומינים' is the object of 'להחזיק' with a double predicate: 'to hold him as apikorsim and heretics'.
- 'כי זה הדור, דור שוא ושקר' repeats 'דור' as an apposition: 'this generation, a generation of vanity and falsehood'.
- The quotations are 'Happy are we' (the speech of the common folk) and 'If you despise My statutes' (Lev 26:15 with the author's variant), introduced by 'לומר' and 'שנאמר'.

Commentary rulings:
- apparatus: Lev 26:15: 'ואם בחקתי תמאסו'; the author writes בחקותי and adds 'במלמדי' in his own sentence.
- apparatus: Ps 134:1: 'העומדים בבית ה' בלילות'.
- apparatus: Jer 22:28: 'כלי אין חפץ בו' (the allusion in 'ככלי אין חפץ בו').
- apparatus: 'בקרו' (3:6.58) is uncertain; see the token note.
- apparatus: Wikisource base text (Q2): translated as it stands; flagged until the 1931 scan is collated.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

