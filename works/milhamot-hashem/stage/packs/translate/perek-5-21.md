# Translation pack: milhamot-hashem perek-5-21 (Milchamot ha-Shem 5:30–5:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-21`.

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

- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
- הל' X after the Rambam: "Laws of X", noting once per unit that it is the
  Mishneh Torah. No reference the author lacks.
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
- A quotation cut short in the author's text (ואלהי in Lev 19:4): translate what
  is written ("and the gods of"), add nothing, say so in a note.
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
- פ' + numeral after a book title: "chapter" when the numeral exceeds the
  parashot of that book (Midrash Rabbah 143 is a section/chapter number as
  the author gives it); otherwise "parashah". Note the choice.
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


## Milchamot ha-Shem 5:30

Hebrew (Wikisource): נו) וז"ל הרמב"ם בפרק ב' מהל' ע"ז עיקר הציווי בע"ש שלא לעבוד אחד מכל הנבראים הנבראים מהם. ואעפ"י שהעובד יודע שה' הוא האלהים והוא עובד הנברא הזה על דרך שעבד אנוש ואנשי דורו תחלה הרי זה עובד ע"ז. ובריש פרק א' כ' בימי אנוש טעו בני אדם טעות גדולה וכו' אמרו הואיל והאל ברא כוכבים וגלגלים אלו להנהיג את העולם ונתנן במרום וחלק להם כבוד. והן שמשים שמשמשים לפניו. ראוים הם לשבחם ולפארם ולחלוק להם כבוד. וזה הוא רצון האל ב"ה לגדל ולכבד מי שגדלו וכבדו. כמו שהמלך רוצה לכבד עבדיו והעומדים לפניו וזהו כבודו של מלך. עכ"ל ובילקוט פרשת יתרו לא יהיה לך אלהים אחרים ר' יוסי אומר אלהים אחרים למה נאמר. שלא ליתן פתחון פה לאומות העולם לומר אלו נקראו בשמו כבר היה בהן צורך. הרי נקראו בשמו ואין בהן צורך. אימתי נקראו בשמו בימי אנוש בן שת שנ' אז הוחל לקרא בשם ה'. החל לקרוא ע"ז בשם המקום באותה שעה עלה אוקינוס והציף שלישו של עולם וכו'. ובמדרש הגדול פר' בראשית אנוש, שהיתה דעתו אנושה ונבערה, והתחיל בע"ז בעולם שה' אז הוחל לקרא בשפ ה'. החל לקרוא ע"ז בשם המקום באותה שעה עלה אוקיינוס וכו'. ובמדרש חפץ לר' זכריה הרופא פר' יתרו ובילקוט לא תעשה לך פסל יכול לא תעשה גלומה (ובילקוט גלופה) אבל יעשה אטומה ת"ל כל צמונה. יכול מטע, ת"ל לא תטע לך אשרה. יכול יעשה של עץ ת"ל כל עץ. יכול יעשה של אבן ת"ל ואבן משכית. יכול של זהב ת"ל ואלהי זהב. יכול יעשה של נחשת בדיל ברזל ועופרת ת"ל אלהי מסכה לא תעשה לך. יכול וכו' ומסיים יכול לא יעשה, אבל יעלה בדעתו ת"ל לא יהיה לך אלהים אחרים על פני ע"כ. וכן ב' הרמב"ם ז"ל בפ"א מהל' יסודי התורה וז"ל וכל המעלה בדעתו שיש שם אלוה אחר חוץ מזה עובר בלא תעשה שנ' לא יהיה לך אלהים אחרים על פני וכפר בעיקר שזהו העיקר הגדול שהכל תלוי בו ע"כ. מבואר מכל זה בהדיא שאע"פ שאינו עושה שום מעשה רק שמאמין בדעתו שיש אלוה אחר למטה ממנו מהלאהים העליון ה"ז מין.

| id | word | punctuation after |
|---|---|---|
| 5:30.1 | נו | ) |
| 5:30.2 | וז"ל |  |
| 5:30.3 | הרמב"ם |  |
| 5:30.4 | בפרק |  |
| 5:30.5 | ב | ' |
| 5:30.6 | מהל | ' |
| 5:30.7 | ע"ז |  |
| 5:30.8 | עיקר |  |
| 5:30.9 | הציווי |  |
| 5:30.10 | בע"ש |  |
| 5:30.11 | שלא |  |
| 5:30.12 | לעבוד |  |
| 5:30.13 | אחד |  |
| 5:30.14 | מכל |  |
| 5:30.15 | הנבראים |  |
| 5:30.16 | הנבראים |  |
| 5:30.17 | מהם | . |
| 5:30.18 | ואעפ"י |  |
| 5:30.19 | שהעובד |  |
| 5:30.20 | יודע |  |
| 5:30.21 | שה | ' |
| 5:30.22 | הוא |  |
| 5:30.23 | האלהים |  |
| 5:30.24 | והוא |  |
| 5:30.25 | עובד |  |
| 5:30.26 | הנברא |  |
| 5:30.27 | הזה |  |
| 5:30.28 | על |  |
| 5:30.29 | דרך |  |
| 5:30.30 | שעבד |  |
| 5:30.31 | אנוש |  |
| 5:30.32 | ואנשי |  |
| 5:30.33 | דורו |  |
| 5:30.34 | תחלה |  |
| 5:30.35 | הרי |  |
| 5:30.36 | זה |  |
| 5:30.37 | עובד |  |
| 5:30.38 | ע"ז | . |
| 5:30.39 | ובריש |  |
| 5:30.40 | פרק |  |
| 5:30.41 | א | ' |
| 5:30.42 | כ | ' |
| 5:30.43 | בימי |  |
| 5:30.44 | אנוש |  |
| 5:30.45 | טעו |  |
| 5:30.46 | בני |  |
| 5:30.47 | אדם |  |
| 5:30.48 | טעות |  |
| 5:30.49 | גדולה |  |
| 5:30.50 | וכו | ' |
| 5:30.51 | אמרו |  |
| 5:30.52 | הואיל |  |
| 5:30.53 | והאל |  |
| 5:30.54 | ברא |  |
| 5:30.55 | כוכבים |  |
| 5:30.56 | וגלגלים |  |
| 5:30.57 | אלו |  |
| 5:30.58 | להנהיג |  |
| 5:30.59 | את |  |
| 5:30.60 | העולם |  |
| 5:30.61 | ונתנן |  |
| 5:30.62 | במרום |  |
| 5:30.63 | וחלק |  |
| 5:30.64 | להם |  |
| 5:30.65 | כבוד | . |
| 5:30.66 | והן |  |
| 5:30.67 | שמשים |  |
| 5:30.68 | שמשמשים |  |
| 5:30.69 | לפניו | . |
| 5:30.70 | ראוים |  |
| 5:30.71 | הם |  |
| 5:30.72 | לשבחם |  |
| 5:30.73 | ולפארם |  |
| 5:30.74 | ולחלוק |  |
| 5:30.75 | להם |  |
| 5:30.76 | כבוד | . |
| 5:30.77 | וזה |  |
| 5:30.78 | הוא |  |
| 5:30.79 | רצון |  |
| 5:30.80 | האל |  |
| 5:30.81 | ב"ה |  |
| 5:30.82 | לגדל |  |
| 5:30.83 | ולכבד |  |
| 5:30.84 | מי |  |
| 5:30.85 | שגדלו |  |
| 5:30.86 | וכבדו | . |
| 5:30.87 | כמו |  |
| 5:30.88 | שהמלך |  |
| 5:30.89 | רוצה |  |
| 5:30.90 | לכבד |  |
| 5:30.91 | עבדיו |  |
| 5:30.92 | והעומדים |  |
| 5:30.93 | לפניו |  |
| 5:30.94 | וזהו |  |
| 5:30.95 | כבודו |  |
| 5:30.96 | של |  |
| 5:30.97 | מלך | . |
| 5:30.98 | עכ"ל |  |
| 5:30.99 | ובילקוט |  |
| 5:30.100 | פרשת |  |
| 5:30.101 | יתרו |  |
| 5:30.102 | לא |  |
| 5:30.103 | יהיה |  |
| 5:30.104 | לך |  |
| 5:30.105 | אלהים |  |
| 5:30.106 | אחרים |  |
| 5:30.107 | ר | ' |
| 5:30.108 | יוסי |  |
| 5:30.109 | אומר |  |
| 5:30.110 | אלהים |  |
| 5:30.111 | אחרים |  |
| 5:30.112 | למה |  |
| 5:30.113 | נאמר | . |
| 5:30.114 | שלא |  |
| 5:30.115 | ליתן |  |
| 5:30.116 | פתחון |  |
| 5:30.117 | פה |  |
| 5:30.118 | לאומות |  |
| 5:30.119 | העולם |  |
| 5:30.120 | לומר |  |
| 5:30.121 | אלו |  |
| 5:30.122 | נקראו |  |
| 5:30.123 | בשמו |  |
| 5:30.124 | כבר |  |
| 5:30.125 | היה |  |
| 5:30.126 | בהן |  |
| 5:30.127 | צורך | . |
| 5:30.128 | הרי |  |
| 5:30.129 | נקראו |  |
| 5:30.130 | בשמו |  |
| 5:30.131 | ואין |  |
| 5:30.132 | בהן |  |
| 5:30.133 | צורך | . |
| 5:30.134 | אימתי |  |
| 5:30.135 | נקראו |  |
| 5:30.136 | בשמו |  |
| 5:30.137 | בימי |  |
| 5:30.138 | אנוש |  |
| 5:30.139 | בן |  |
| 5:30.140 | שת |  |
| 5:30.141 | שנ | ' |
| 5:30.142 | אז |  |
| 5:30.143 | הוחל |  |
| 5:30.144 | לקרא |  |
| 5:30.145 | בשם |  |
| 5:30.146 | ה | '. |
| 5:30.147 | החל |  |
| 5:30.148 | לקרוא |  |
| 5:30.149 | ע"ז |  |
| 5:30.150 | בשם |  |
| 5:30.151 | המקום |  |
| 5:30.152 | באותה |  |
| 5:30.153 | שעה |  |
| 5:30.154 | עלה |  |
| 5:30.155 | אוקינוס |  |
| 5:30.156 | והציף |  |
| 5:30.157 | שלישו |  |
| 5:30.158 | של |  |
| 5:30.159 | עולם |  |
| 5:30.160 | וכו | '. |
| 5:30.161 | ובמדרש |  |
| 5:30.162 | הגדול |  |
| 5:30.163 | פר | ' |
| 5:30.164 | בראשית |  |
| 5:30.165 | אנוש | , |
| 5:30.166 | שהיתה |  |
| 5:30.167 | דעתו |  |
| 5:30.168 | אנושה |  |
| 5:30.169 | ונבערה | , |
| 5:30.170 | והתחיל |  |
| 5:30.171 | בע"ז |  |
| 5:30.172 | בעולם |  |
| 5:30.173 | שה | ' |
| 5:30.174 | אז |  |
| 5:30.175 | הוחל |  |
| 5:30.176 | לקרא |  |
| 5:30.177 | בשפ |  |
| 5:30.178 | ה | '. |
| 5:30.179 | החל |  |
| 5:30.180 | לקרוא |  |
| 5:30.181 | ע"ז |  |
| 5:30.182 | בשם |  |
| 5:30.183 | המקום |  |
| 5:30.184 | באותה |  |
| 5:30.185 | שעה |  |
| 5:30.186 | עלה |  |
| 5:30.187 | אוקיינוס |  |
| 5:30.188 | וכו | '. |
| 5:30.189 | ובמדרש |  |
| 5:30.190 | חפץ |  |
| 5:30.191 | לר | ' |
| 5:30.192 | זכריה |  |
| 5:30.193 | הרופא |  |
| 5:30.194 | פר | ' |
| 5:30.195 | יתרו |  |
| 5:30.196 | ובילקוט |  |
| 5:30.197 | לא |  |
| 5:30.198 | תעשה |  |
| 5:30.199 | לך |  |
| 5:30.200 | פסל |  |
| 5:30.201 | יכול |  |
| 5:30.202 | לא |  |
| 5:30.203 | תעשה |  |
| 5:30.204 | גלומה |  |
| 5:30.205 | ובילקוט |  |
| 5:30.206 | גלופה | ) |
| 5:30.207 | אבל |  |
| 5:30.208 | יעשה |  |
| 5:30.209 | אטומה |  |
| 5:30.210 | ת"ל |  |
| 5:30.211 | כל |  |
| 5:30.212 | צמונה | . |
| 5:30.213 | יכול |  |
| 5:30.214 | מטע | , |
| 5:30.215 | ת"ל |  |
| 5:30.216 | לא |  |
| 5:30.217 | תטע |  |
| 5:30.218 | לך |  |
| 5:30.219 | אשרה | . |
| 5:30.220 | יכול |  |
| 5:30.221 | יעשה |  |
| 5:30.222 | של |  |
| 5:30.223 | עץ |  |
| 5:30.224 | ת"ל |  |
| 5:30.225 | כל |  |
| 5:30.226 | עץ | . |
| 5:30.227 | יכול |  |
| 5:30.228 | יעשה |  |
| 5:30.229 | של |  |
| 5:30.230 | אבן |  |
| 5:30.231 | ת"ל |  |
| 5:30.232 | ואבן |  |
| 5:30.233 | משכית | . |
| 5:30.234 | יכול |  |
| 5:30.235 | של |  |
| 5:30.236 | זהב |  |
| 5:30.237 | ת"ל |  |
| 5:30.238 | ואלהי |  |
| 5:30.239 | זהב | . |
| 5:30.240 | יכול |  |
| 5:30.241 | יעשה |  |
| 5:30.242 | של |  |
| 5:30.243 | נחשת |  |
| 5:30.244 | בדיל |  |
| 5:30.245 | ברזל |  |
| 5:30.246 | ועופרת |  |
| 5:30.247 | ת"ל |  |
| 5:30.248 | אלהי |  |
| 5:30.249 | מסכה |  |
| 5:30.250 | לא |  |
| 5:30.251 | תעשה |  |
| 5:30.252 | לך | . |
| 5:30.253 | יכול |  |
| 5:30.254 | וכו | ' |
| 5:30.255 | ומסיים |  |
| 5:30.256 | יכול |  |
| 5:30.257 | לא |  |
| 5:30.258 | יעשה | , |
| 5:30.259 | אבל |  |
| 5:30.260 | יעלה |  |
| 5:30.261 | בדעתו |  |
| 5:30.262 | ת"ל |  |
| 5:30.263 | לא |  |
| 5:30.264 | יהיה |  |
| 5:30.265 | לך |  |
| 5:30.266 | אלהים |  |
| 5:30.267 | אחרים |  |
| 5:30.268 | על |  |
| 5:30.269 | פני |  |
| 5:30.270 | ע"כ | . |
| 5:30.271 | וכן |  |
| 5:30.272 | ב | ' |
| 5:30.273 | הרמב"ם |  |
| 5:30.274 | ז"ל |  |
| 5:30.275 | בפ"א |  |
| 5:30.276 | מהל | ' |
| 5:30.277 | יסודי |  |
| 5:30.278 | התורה |  |
| 5:30.279 | וז"ל |  |
| 5:30.280 | וכל |  |
| 5:30.281 | המעלה |  |
| 5:30.282 | בדעתו |  |
| 5:30.283 | שיש |  |
| 5:30.284 | שם |  |
| 5:30.285 | אלוה |  |
| 5:30.286 | אחר |  |
| 5:30.287 | חוץ |  |
| 5:30.288 | מזה |  |
| 5:30.289 | עובר |  |
| 5:30.290 | בלא |  |
| 5:30.291 | תעשה |  |
| 5:30.292 | שנ | ' |
| 5:30.293 | לא |  |
| 5:30.294 | יהיה |  |
| 5:30.295 | לך |  |
| 5:30.296 | אלהים |  |
| 5:30.297 | אחרים |  |
| 5:30.298 | על |  |
| 5:30.299 | פני |  |
| 5:30.300 | וכפר |  |
| 5:30.301 | בעיקר |  |
| 5:30.302 | שזהו |  |
| 5:30.303 | העיקר |  |
| 5:30.304 | הגדול |  |
| 5:30.305 | שהכל |  |
| 5:30.306 | תלוי |  |
| 5:30.307 | בו |  |
| 5:30.308 | ע"כ | . |
| 5:30.309 | מבואר |  |
| 5:30.310 | מכל |  |
| 5:30.311 | זה |  |
| 5:30.312 | בהדיא |  |
| 5:30.313 | שאע"פ |  |
| 5:30.314 | שאינו |  |
| 5:30.315 | עושה |  |
| 5:30.316 | שום |  |
| 5:30.317 | מעשה |  |
| 5:30.318 | רק |  |
| 5:30.319 | שמאמין |  |
| 5:30.320 | בדעתו |  |
| 5:30.321 | שיש |  |
| 5:30.322 | אלוה |  |
| 5:30.323 | אחר |  |
| 5:30.324 | למטה |  |
| 5:30.325 | ממנו |  |
| 5:30.326 | מהלאהים |  |
| 5:30.327 | העליון |  |
| 5:30.328 | ה"ז |  |
| 5:30.329 | מין | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

