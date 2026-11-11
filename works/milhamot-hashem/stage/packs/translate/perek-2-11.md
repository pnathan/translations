# Translation pack: milhamot-hashem perek-2-11 (Milchamot ha-Shem 2:23–2:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-11`.

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
- **קדמון / קדמות**: context. Of God or the world, in the philosophical sense:
  "eternal / eternity", with the note "without beginning": "בורא אחד יחיד
  קדמון" (2:25) → "one, unique, eternal Creator"; "אם העולם קדמון" (6:31) →
  "if the world is eternal". Keep נצחי / נצחיות → "everlasting /
  everlastingness" distinct from it (6:10). Of people: קדמונינו → "our
  forebears" (fixed: the author's appeal to Yemenite tradition, e.g.
  "באמונת קדמונינו ביחוד השי"ת", 5:3); הקדמונים → "the ancients" ("הפילסופים
  הקדמונים" → "the ancient philosophers"). The verb קדמו → "preceded".
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
- Poetic phrases quoted from a verse (הר המור): translate the words
  ("the mountain of myrrh"); a note names the verse's referent.
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
- A variant spelling of a word the text does not gloss (ברותא for ברייתא):
  transliterate it ("berota"), with a note.
- יש מאין / יש מיש: translate, "something from nothing" / "something from
  something"; no transliteration.
- Authority names by their books or acronyms: Ravad (הראב"ד), Lechem Mishneh
  (לח"מ), Chavot Yair (חו"י, a note if the expansion is uncertain), Pitchei
  Teshuvah: the project's plain transliteration, a note at first occurrence.
- הרב before a book title (הרב לחם משנה): "the Rav Lechem Mishneh", never drop
  הרב. (Applies with the "author of" ruling: use "the Rabbi, the author of
  [title]" only where the Hebrew says בעל / מחבר.)
- Fixed terms: אפיקורסים = "apikorsim"; בריאה in the four worlds = "Creation";
  תהו ובהו = "formlessness and void" (Gen 1:2), listed in `apparatus`.
- A verse the author cuts short or continues with his own comma: render only
  the words on the page; no completion ("before Me") added; the fixed wording of
  the verse applies to the words that are quoted.
- A Judeo-Arabic rendering the author quotes from Saadia (אלבלד שברא): translate
  the Arabic as he gives it and name Saadia's translation in `apparatus`; the
  Arabic is not transliterated in the English line.
- דברי תימה: "astonishing words" (words that astonish).
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
- קצוצי פאה (Jer 9:25): the literal "those with clipped temples"; a context
  note that in the verse the phrase names desert peoples.

- לשבר את האזן: "to accommodate the ear" (literal "to break the ear" in a
  note); never "shock".
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
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


## Milchamot ha-Shem 2:23

Hebrew (Wikisource): חמישית העם היושבים באזיא אם יבוא אליהם מכתב פקודה מהמלך בכפם יתפשוהו, בפיהם ישקוהו, ועל ראשם יניחוהו, אף כי אנו נכבד את דברי אלהים חיים מלך על כל הארץ.

| id | word | punctuation after |
|---|---|---|
| 2:23.1 | חמישית |  |
| 2:23.2 | העם |  |
| 2:23.3 | היושבים |  |
| 2:23.4 | באזיא |  |
| 2:23.5 | אם |  |
| 2:23.6 | יבוא |  |
| 2:23.7 | אליהם |  |
| 2:23.8 | מכתב |  |
| 2:23.9 | פקודה |  |
| 2:23.10 | מהמלך |  |
| 2:23.11 | בכפם |  |
| 2:23.12 | יתפשוהו | , |
| 2:23.13 | בפיהם |  |
| 2:23.14 | ישקוהו | , |
| 2:23.15 | ועל |  |
| 2:23.16 | ראשם |  |
| 2:23.17 | יניחוהו | , |
| 2:23.18 | אף |  |
| 2:23.19 | כי |  |
| 2:23.20 | אנו |  |
| 2:23.21 | נכבד |  |
| 2:23.22 | את |  |
| 2:23.23 | דברי |  |
| 2:23.24 | אלהים |  |
| 2:23.25 | חיים |  |
| 2:23.26 | מלך |  |
| 2:23.27 | על |  |
| 2:23.28 | כל |  |
| 2:23.29 | הארץ | . |


## Milchamot ha-Shem 2:24

Hebrew (Wikisource): ששית כשהביאו שבעים ושנים זקנים את התורה לתלמי מלך מצרים, קם מכסאו והשתחוה ז' פעמים לקראתה, ואם עובד אלילים אשר לא הביא צוארו בעול תורתינו כבדה כל כך, אף כי אנחנו הבאים בבריתה ואשר ניתנה לנו למורשה.

| id | word | punctuation after |
|---|---|---|
| 2:24.1 | ששית |  |
| 2:24.2 | כשהביאו |  |
| 2:24.3 | שבעים |  |
| 2:24.4 | ושנים |  |
| 2:24.5 | זקנים |  |
| 2:24.6 | את |  |
| 2:24.7 | התורה |  |
| 2:24.8 | לתלמי |  |
| 2:24.9 | מלך |  |
| 2:24.10 | מצרים | , |
| 2:24.11 | קם |  |
| 2:24.12 | מכסאו |  |
| 2:24.13 | והשתחוה |  |
| 2:24.14 | ז | ' |
| 2:24.15 | פעמים |  |
| 2:24.16 | לקראתה | , |
| 2:24.17 | ואם |  |
| 2:24.18 | עובד |  |
| 2:24.19 | אלילים |  |
| 2:24.20 | אשר |  |
| 2:24.21 | לא |  |
| 2:24.22 | הביא |  |
| 2:24.23 | צוארו |  |
| 2:24.24 | בעול |  |
| 2:24.25 | תורתינו |  |
| 2:24.26 | כבדה |  |
| 2:24.27 | כל |  |
| 2:24.28 | כך | , |
| 2:24.29 | אף |  |
| 2:24.30 | כי |  |
| 2:24.31 | אנחנו |  |
| 2:24.32 | הבאים |  |
| 2:24.33 | בבריתה |  |
| 2:24.34 | ואשר |  |
| 2:24.35 | ניתנה |  |
| 2:24.36 | לנו |  |
| 2:24.37 | למורשה | . |


## Milchamot ha-Shem 2:25

Hebrew (Wikisource): שביעית יסודי ועיקרי דתינו הם מציאות בורא אחד יחיד קדמון אשר איננו גשם ולא כח בגשם, נותן התורה לעמו ישראל על ידי המחוקק הנאמן משה עבד ה' אדון כל הנביאים. הוא יתברך המציץ ומשגיח בעולמות שברא. ועיניו משוטטות בכל, וחוקר ובוחן לבות בני אדם ושופט בצדק על עלילות מצעדי גבר. ממנו לבד נבקש עזרה בצר לנו, והוא יפרה נצר מגזע ישי אשר יקבץ נדחי ישראל מארבע כנפות הארץ, ומכל קצוצי פאה אשר זרו אותם שמה, עצמות מתיהם היבשות תחינה ונבלתם יקומון. הן כל אלה עקרי ושרשי דתינו. ומה שמץ ע"ז נמצא בם? אף לדעת אותם אשר דעותיהם ואמונותיהם נבדלו מדעותינו ומאמונתנו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 2:25.1 | שביעית |  |
| 2:25.2 | יסודי |  |
| 2:25.3 | ועיקרי |  |
| 2:25.4 | דתינו |  |
| 2:25.5 | הם |  |
| 2:25.6 | מציאות |  |
| 2:25.7 | בורא |  |
| 2:25.8 | אחד |  |
| 2:25.9 | יחיד |  |
| 2:25.10 | קדמון |  |
| 2:25.11 | אשר |  |
| 2:25.12 | איננו |  |
| 2:25.13 | גשם |  |
| 2:25.14 | ולא |  |
| 2:25.15 | כח |  |
| 2:25.16 | בגשם | , |
| 2:25.17 | נותן |  |
| 2:25.18 | התורה |  |
| 2:25.19 | לעמו |  |
| 2:25.20 | ישראל |  |
| 2:25.21 | על |  |
| 2:25.22 | ידי |  |
| 2:25.23 | המחוקק |  |
| 2:25.24 | הנאמן |  |
| 2:25.25 | משה |  |
| 2:25.26 | עבד |  |
| 2:25.27 | ה | ' |
| 2:25.28 | אדון |  |
| 2:25.29 | כל |  |
| 2:25.30 | הנביאים | . |
| 2:25.31 | הוא |  |
| 2:25.32 | יתברך |  |
| 2:25.33 | המציץ |  |
| 2:25.34 | ומשגיח |  |
| 2:25.35 | בעולמות |  |
| 2:25.36 | שברא | . |
| 2:25.37 | ועיניו |  |
| 2:25.38 | משוטטות |  |
| 2:25.39 | בכל | , |
| 2:25.40 | וחוקר |  |
| 2:25.41 | ובוחן |  |
| 2:25.42 | לבות |  |
| 2:25.43 | בני |  |
| 2:25.44 | אדם |  |
| 2:25.45 | ושופט |  |
| 2:25.46 | בצדק |  |
| 2:25.47 | על |  |
| 2:25.48 | עלילות |  |
| 2:25.49 | מצעדי |  |
| 2:25.50 | גבר | . |
| 2:25.51 | ממנו |  |
| 2:25.52 | לבד |  |
| 2:25.53 | נבקש |  |
| 2:25.54 | עזרה |  |
| 2:25.55 | בצר |  |
| 2:25.56 | לנו | , |
| 2:25.57 | והוא |  |
| 2:25.58 | יפרה |  |
| 2:25.59 | נצר |  |
| 2:25.60 | מגזע |  |
| 2:25.61 | ישי |  |
| 2:25.62 | אשר |  |
| 2:25.63 | יקבץ |  |
| 2:25.64 | נדחי |  |
| 2:25.65 | ישראל |  |
| 2:25.66 | מארבע |  |
| 2:25.67 | כנפות |  |
| 2:25.68 | הארץ | , |
| 2:25.69 | ומכל |  |
| 2:25.70 | קצוצי |  |
| 2:25.71 | פאה |  |
| 2:25.72 | אשר |  |
| 2:25.73 | זרו |  |
| 2:25.74 | אותם |  |
| 2:25.75 | שמה | , |
| 2:25.76 | עצמות |  |
| 2:25.77 | מתיהם |  |
| 2:25.78 | היבשות |  |
| 2:25.79 | תחינה |  |
| 2:25.80 | ונבלתם |  |
| 2:25.81 | יקומון | . |
| 2:25.82 | הן |  |
| 2:25.83 | כל |  |
| 2:25.84 | אלה |  |
| 2:25.85 | עקרי |  |
| 2:25.86 | ושרשי |  |
| 2:25.87 | דתינו | . |
| 2:25.88 | ומה |  |
| 2:25.89 | שמץ |  |
| 2:25.90 | ע"ז |  |
| 2:25.91 | נמצא |  |
| 2:25.92 | בם | ? |
| 2:25.93 | אף |  |
| 2:25.94 | לדעת |  |
| 2:25.95 | אותם |  |
| 2:25.96 | אשר |  |
| 2:25.97 | דעותיהם |  |
| 2:25.98 | ואמונותיהם |  |
| 2:25.99 | נבדלו |  |
| 2:25.100 | מדעותינו |  |
| 2:25.101 | ומאמונתנו |  |
| 2:25.102 | עכ"ל | . |


## Milchamot ha-Shem 2:26

Hebrew (Wikisource): ולפי דרכינו דרך רז"ל בבריתא הנז"ל אין צורך לכל זה שאין כוונתנו שאנו משתחוים מול הס"ת אלא למי שבחר בתורה, ובהדיא אנו אומרים בואו נשתחוה ונכרעה נברכה לפני ה' עושנו. והשתחוית אברהם אבינו למלאכים וכן יהושע, איבעית אימא כיוון שנדמו להם כאנשים קרינא בהו מאי דאמור רבנן בבריתא לא תשתחוה להם, להם אי אתה משתחוה אבל אתה משתחוה לאדם שכמותך כיון שנגלו להם בדמות אנשים ואיבעית אימא שכוונת אברהם ויהושע להשתחות למי ששלחם וכמו שכ' בעל העיקרים.

| id | word | punctuation after |
|---|---|---|
| 2:26.1 | ולפי |  |
| 2:26.2 | דרכינו |  |
| 2:26.3 | דרך |  |
| 2:26.4 | רז"ל |  |
| 2:26.5 | בבריתא |  |
| 2:26.6 | הנז"ל |  |
| 2:26.7 | אין |  |
| 2:26.8 | צורך |  |
| 2:26.9 | לכל |  |
| 2:26.10 | זה |  |
| 2:26.11 | שאין |  |
| 2:26.12 | כוונתנו |  |
| 2:26.13 | שאנו |  |
| 2:26.14 | משתחוים |  |
| 2:26.15 | מול |  |
| 2:26.16 | הס"ת |  |
| 2:26.17 | אלא |  |
| 2:26.18 | למי |  |
| 2:26.19 | שבחר |  |
| 2:26.20 | בתורה | , |
| 2:26.21 | ובהדיא |  |
| 2:26.22 | אנו |  |
| 2:26.23 | אומרים |  |
| 2:26.24 | בואו |  |
| 2:26.25 | נשתחוה |  |
| 2:26.26 | ונכרעה |  |
| 2:26.27 | נברכה |  |
| 2:26.28 | לפני |  |
| 2:26.29 | ה | ' |
| 2:26.30 | עושנו | . |
| 2:26.31 | והשתחוית |  |
| 2:26.32 | אברהם |  |
| 2:26.33 | אבינו |  |
| 2:26.34 | למלאכים |  |
| 2:26.35 | וכן |  |
| 2:26.36 | יהושע | , |
| 2:26.37 | איבעית |  |
| 2:26.38 | אימא |  |
| 2:26.39 | כיוון |  |
| 2:26.40 | שנדמו |  |
| 2:26.41 | להם |  |
| 2:26.42 | כאנשים |  |
| 2:26.43 | קרינא |  |
| 2:26.44 | בהו |  |
| 2:26.45 | מאי |  |
| 2:26.46 | דאמור |  |
| 2:26.47 | רבנן |  |
| 2:26.48 | בבריתא |  |
| 2:26.49 | לא |  |
| 2:26.50 | תשתחוה |  |
| 2:26.51 | להם | , |
| 2:26.52 | להם |  |
| 2:26.53 | אי |  |
| 2:26.54 | אתה |  |
| 2:26.55 | משתחוה |  |
| 2:26.56 | אבל |  |
| 2:26.57 | אתה |  |
| 2:26.58 | משתחוה |  |
| 2:26.59 | לאדם |  |
| 2:26.60 | שכמותך |  |
| 2:26.61 | כיון |  |
| 2:26.62 | שנגלו |  |
| 2:26.63 | להם |  |
| 2:26.64 | בדמות |  |
| 2:26.65 | אנשים |  |
| 2:26.66 | ואיבעית |  |
| 2:26.67 | אימא |  |
| 2:26.68 | שכוונת |  |
| 2:26.69 | אברהם |  |
| 2:26.70 | ויהושע |  |
| 2:26.71 | להשתחות |  |
| 2:26.72 | למי |  |
| 2:26.73 | ששלחם |  |
| 2:26.74 | וכמו |  |
| 2:26.75 | שכ | ' |
| 2:26.76 | בעל |  |
| 2:26.77 | העיקרים | . |


## Milchamot ha-Shem 2:27

Hebrew (Wikisource): ושאלתך מענין הנבואה אעפ"י שאינו שייך לעניינינו ולא שאלת רק להטעותנו ולהשיאנו לדבר אחר, הנה זה מבואר מדרגות הנבואה ברמב"ם הל' יסודי התורה, ובחלק שני מספר המורה, ובספר האמונות והדעות להרס"ג. ולפי שמצאתיך חכם בעיניך אחזור לבקש ממך להשיב לי תשובה ניצחת מדברי רז"ל בתלמוד, כי רק אתה לבדך (לפי דעתך) לך נגלו כל תעלומות, ובינה יתירה ורוח חדש ניתן בקרבך. הורינו אני וחבירי איך הדברים יעלו בקנה אחד מלאות וטובות. כי נפשינו חשקה לשמוע לימודיך, ולהשתעשע בתורתך התמימה, ולך תהיה צדקה, וגם כבוד גדול תנחל מאיתנו.

| id | word | punctuation after |
|---|---|---|
| 2:27.1 | ושאלתך |  |
| 2:27.2 | מענין |  |
| 2:27.3 | הנבואה |  |
| 2:27.4 | אעפ"י |  |
| 2:27.5 | שאינו |  |
| 2:27.6 | שייך |  |
| 2:27.7 | לעניינינו |  |
| 2:27.8 | ולא |  |
| 2:27.9 | שאלת |  |
| 2:27.10 | רק |  |
| 2:27.11 | להטעותנו |  |
| 2:27.12 | ולהשיאנו |  |
| 2:27.13 | לדבר |  |
| 2:27.14 | אחר | , |
| 2:27.15 | הנה |  |
| 2:27.16 | זה |  |
| 2:27.17 | מבואר |  |
| 2:27.18 | מדרגות |  |
| 2:27.19 | הנבואה |  |
| 2:27.20 | ברמב"ם |  |
| 2:27.21 | הל | ' |
| 2:27.22 | יסודי |  |
| 2:27.23 | התורה | , |
| 2:27.24 | ובחלק |  |
| 2:27.25 | שני |  |
| 2:27.26 | מספר |  |
| 2:27.27 | המורה | , |
| 2:27.28 | ובספר |  |
| 2:27.29 | האמונות |  |
| 2:27.30 | והדעות |  |
| 2:27.31 | להרס"ג | . |
| 2:27.32 | ולפי |  |
| 2:27.33 | שמצאתיך |  |
| 2:27.34 | חכם |  |
| 2:27.35 | בעיניך |  |
| 2:27.36 | אחזור |  |
| 2:27.37 | לבקש |  |
| 2:27.38 | ממך |  |
| 2:27.39 | להשיב |  |
| 2:27.40 | לי |  |
| 2:27.41 | תשובה |  |
| 2:27.42 | ניצחת |  |
| 2:27.43 | מדברי |  |
| 2:27.44 | רז"ל |  |
| 2:27.45 | בתלמוד | , |
| 2:27.46 | כי |  |
| 2:27.47 | רק |  |
| 2:27.48 | אתה |  |
| 2:27.49 | לבדך |  |
| 2:27.50 | לפי |  |
| 2:27.51 | דעתך | ) |
| 2:27.52 | לך |  |
| 2:27.53 | נגלו |  |
| 2:27.54 | כל |  |
| 2:27.55 | תעלומות | , |
| 2:27.56 | ובינה |  |
| 2:27.57 | יתירה |  |
| 2:27.58 | ורוח |  |
| 2:27.59 | חדש |  |
| 2:27.60 | ניתן |  |
| 2:27.61 | בקרבך | . |
| 2:27.62 | הורינו |  |
| 2:27.63 | אני |  |
| 2:27.64 | וחבירי |  |
| 2:27.65 | איך |  |
| 2:27.66 | הדברים |  |
| 2:27.67 | יעלו |  |
| 2:27.68 | בקנה |  |
| 2:27.69 | אחד |  |
| 2:27.70 | מלאות |  |
| 2:27.71 | וטובות | . |
| 2:27.72 | כי |  |
| 2:27.73 | נפשינו |  |
| 2:27.74 | חשקה |  |
| 2:27.75 | לשמוע |  |
| 2:27.76 | לימודיך | , |
| 2:27.77 | ולהשתעשע |  |
| 2:27.78 | בתורתך |  |
| 2:27.79 | התמימה | , |
| 2:27.80 | ולך |  |
| 2:27.81 | תהיה |  |
| 2:27.82 | צדקה | , |
| 2:27.83 | וגם |  |
| 2:27.84 | כבוד |  |
| 2:27.85 | גדול |  |
| 2:27.86 | תנחל |  |
| 2:27.87 | מאיתנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

