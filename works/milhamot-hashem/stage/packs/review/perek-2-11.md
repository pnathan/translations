# Review pack: milhamot-hashem perek-2-11 (Milchamot ha-Shem 2:23–2:27)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-11`.

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


### Translation under review (2:23)

English: "The fifth: the people who dwell in Asia, if a written command from the king comes to them, they take it in their hands, kiss it with their mouths and place it on their heads. How much more so shall we honor the words of the living God, King over all the earth.

Literal: fifth the-people the-dwelling in-Asia if comes to-them written-command from-the-king in-their-hands they-seize-it, with-their-mouths they-kiss-it, and-on their-heads they-place-it, how-much-more so we-shall-honor the-words-of God living King over all the-earth

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:23.1 | fifth |  |  | ordinal of a numbered series: 2:24 is the sixth, 2:25 the seventh |
| 2:23.2 | the people |  |  |  |
| 2:23.3 | dwelling |  |  |  |
| 2:23.4 | in Asia |  |  |  |
| 2:23.5 | if |  |  |  |
| 2:23.6 | comes |  |  | yiqtol after אם: the protasis of a conditional; the verbs that follow are yiqtol in a habitual sense |
| 2:23.7 | to them |  |  |  |
| 2:23.8 | written |  |  | מכתב, 'letter, writing' |
| 2:23.9 | command |  |  | פקודה, an order |
| 2:23.10 | from the king |  |  |  |
| 2:23.11 | in their hands |  |  |  |
| 2:23.12 | they take it |  |  | יתפשוהו: root תפש with the suffix ־הו, 'it'; the command is the object |
| 2:23.13 | with their mouths |  |  |  |
| 2:23.14 | they kiss it |  |  |  |
| 2:23.15 | and on |  |  |  |
| 2:23.16 | their heads |  |  |  |
| 2:23.17 | they place it |  |  |  |
| 2:23.18 | how much more |  |  | אף כי: 'how much more so', the apodosis of an argument from the lesser to the greater |
| 2:23.19 | so |  |  |  |
| 2:23.20 | we |  |  |  |
| 2:23.21 | shall honor |  |  | נכבד: cohortative, 'let us honor' |
| 2:23.22 | - |  |  | object marker את, no English |
| 2:23.23 | the words of |  |  |  |
| 2:23.24 | God |  |  | אלהים: the God of Israel, here 'the living God' (חיים) |
| 2:23.25 | living |  |  |  |
| 2:23.26 | King |  |  |  |
| 2:23.27 | over |  |  |  |
| 2:23.28 | all |  |  |  |
| 2:23.29 | the earth |  |  |  |

Phrases:
- 2:23.1 2:23.2 2:23.3 2:23.4 | The fifth: the people who dwell in Asia, | fifth the-people the-dwelling in-Asia
- 2:23.5 2:23.6 2:23.7 2:23.8 2:23.9 2:23.10 | if a written command from the king comes to them, | if comes to-them written-command from-the-king
- 2:23.11 2:23.12 | they take it in their hands, | in-their-hands they-seize-it
- 2:23.13 2:23.14 | kiss it with their mouths, | with-their-mouths they-kiss-it
- 2:23.15 2:23.16 2:23.17 | and place it on their heads. | and-on their-heads they-place-it
- 2:23.18 2:23.19 | How much more so | how-much-more so
- 2:23.20 2:23.21 | shall we honor | we shall-honor
- 2:23.22 | - | object marker, no English
- 2:23.23 2:23.24 2:23.25 | the words of the living God, | the-words-of God living
- 2:23.26 2:23.27 2:23.28 2:23.29 | King over all the earth. | King over all the-earth

Grammar notes:
- The paragraph is the fifth point of a numbered series, and its number opens the English as it opens the Hebrew.
- The conditional אם takes yiqtol verbs (they take, kiss, place): these describe a habitual practice, so the English uses the present.
- The argument is from the lesser to the greater (qal wa-chomer): if the people honor a king's written command, 'how much more so' the words of God. The cohortative נכבד keeps the 'we shall'.
- The three acts (seizing in the hands, kissing with the mouth, placing on the head) are parallel clauses, and the English keeps their order.

Commentary rulings:
- apparatus: Base text from Wikisource; the 1931 scan is not yet collated, so the paragraph is translated as the base text (decisions, Q2).
- apparatus: Quotation: the double quotation marks open at this segment because the quoted letter is closed by עכ"ל at 2:25. The opening formula is not in this unit.
- apparatus: Allusions, the author's wording and not a quotation: 'the living God' (Jer 10:10) and 'King over all the earth' (cf. Zech 14:9).

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


### Translation under review (2:24)

English: The sixth: when the seventy-two elders brought the Torah to Ptolemy, king of Egypt, he rose from his throne and bowed down seven times to meet it, and if a worshipper of idols who has not put his neck under the yoke of our Torah, which is so heavy, how much more so we who enter into its covenant and to whom it was given as an inheritance.

Literal: sixth when-they-brought seventy and-two elders (object) the-Torah to-Ptolemy king of-Egypt rose from-his-throne and-bowed-down seven times to-meet-it, and-if worshipper-of idols who not put his-neck into-the-yoke of-our-Torah is-heavy so much, how-much-more so we entering in-its-covenant and-to-whom it-was-given to-us as-inheritance

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:24.1 | sixth |  |  |  |
| 2:24.2 | when they brought |  |  | כשהביאו: 'when' with a qatal; the verb comes before its subject |
| 2:24.3 | seventy |  |  |  |
| 2:24.4 | and two |  |  | שבעים ושנים: seventy-two, written 'seventy and two' |
| 2:24.5 | elders |  |  |  |
| 2:24.6 | - |  |  | object marker את, no English |
| 2:24.7 | the Torah |  |  |  |
| 2:24.8 | to Ptolemy |  |  | תלמי, the Hebrew spelling of the king's name |
| 2:24.9 | king |  |  |  |
| 2:24.10 | of Egypt |  |  |  |
| 2:24.11 | rose |  |  |  |
| 2:24.12 | from his throne |  |  | מכסאו: 'from his throne' |
| 2:24.13 | and bowed down |  |  | והשתחוה: the act of reverence; the verb is the same as in 2:26 |
| 2:24.14 | seven |  |  | ז׳ is the numeral seven, not a letter |
| 2:24.15 | times |  |  |  |
| 2:24.16 | to meet it |  |  | לקראתה: lit. 'to meet it', the greeting of a sacred object |
| 2:24.17 | and if |  |  |  |
| 2:24.18 | a worshipper |  |  | עובד אלילים: 'worshipper of idols' |
| 2:24.19 | of idols |  |  |  |
| 2:24.20 | who |  |  |  |
| 2:24.21 | has not |  |  |  |
| 2:24.22 | put |  |  | הביא צוארו בעול: lit. 'brought his neck into the yoke', an idiom of accepting a burden |
| 2:24.23 | his neck |  |  |  |
| 2:24.24 | under the yoke |  |  | בעול: the yoke, the image of taking on the Torah |
| 2:24.25 | of our Torah |  |  |  |
| 2:24.26 | which is heavy |  |  | כבדה: feminine, agreeing with תורתינו; the Torah is what is heavy, not the yoke |
| 2:24.27 | so |  |  |  |
| 2:24.28 | much |  |  |  |
| 2:24.29 | how much more |  |  | אף כי: 'how much more so', as in 2:23 |
| 2:24.30 | so |  |  |  |
| 2:24.31 | we |  |  |  |
| 2:24.32 | who enter |  |  | הבאים: participle, 'who enter' |
| 2:24.33 | into its covenant |  |  | בבריתה: the covenant of the Torah; 'its' is the Torah |
| 2:24.34 | and to whom |  |  |  |
| 2:24.35 | it was given |  |  |  |
| 2:24.36 | to us |  |  |  |
| 2:24.37 | as an inheritance |  |  | למורשה: 'as an inheritance'; see apparatus |

Phrases:
- 2:24.1 | The sixth: | sixth
- 2:24.2 2:24.3 2:24.4 2:24.5 | when the seventy-two elders brought | when-brought seventy and-two elders
- 2:24.6 2:24.7 | the Torah | (object) the-Torah
- 2:24.8 2:24.9 2:24.10 | to Ptolemy, king of Egypt, | to-Ptolemy king of-Egypt
- 2:24.11 2:24.12 | he rose from his throne | rose from-his-throne
- 2:24.13 2:24.14 2:24.15 2:24.16 | and bowed down seven times to meet it, | and-bowed-down seven times to-meet-it
- 2:24.17 2:24.18 2:24.19 | and if a worshipper of idols | and-if worshipper-of idols
- 2:24.20 2:24.21 2:24.22 2:24.23 2:24.24 2:24.25 | who has not put his neck under the yoke of our Torah, | who not brought his-neck into-the-yoke of-our-Torah
- 2:24.26 2:24.27 2:24.28 | which is so heavy, | is-heavy so much
- 2:24.29 2:24.30 | how much more so | how-much-more so
- 2:24.31 2:24.32 2:24.33 | we who enter into its covenant | we entering in-its-covenant
- 2:24.34 2:24.35 2:24.36 2:24.37 | and to whom it was given as an inheritance. | and-to-whom it-was-given to-us as-inheritance

Grammar notes:
- The temporal clause כשהביאו puts the verb first and the subject after it: 'when the seventy-two elders brought the Torah'.
- קם ... והשתחוה are two past acts in sequence (qatal with vav); the English keeps the sequence, 'rose ... and bowed down'.
- The clause introduced by ואם has no verb. The Hebrew leaves the king's act to be understood, and the English keeps the ellipsis; the apodosis 'how much more so we' answers it.
- כבדה is feminine and agrees with תורתינו, so 'which is so heavy' describes the Torah, not the yoke.

Commentary rulings:
- apparatus: Allusion, not quoted: Deut 33:4 for 'למורשה' ('an inheritance').
- apparatus: The number seventy-two and the translation story are the author's (see context).

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


### Translation under review (2:25)

English: "Seventh: the foundation and principle of our religion is the existence of one, unique, eternal Creator, who is not a body, nor a force in a body, who gives the Torah to His people Israel through the faithful lawgiver Moses, servant of the LORD, master of all the prophets. He, may He be blessed, looks on and exercises providence over the worlds that He created. And His eyes roam over all, He searches and examines the hearts of the sons of man and judges justly according to the deeds, the steps of a man. From Him alone we seek help in our distress, and He will bring forth a shoot from the stock of Jesse, who will gather the outcasts of Israel from the four corners of the earth, and from all the cut-off corners to which they scattered them there, the dry bones of their dead will live, and their corpses will rise. Indeed, all these are the principles and roots of our religion. And what trace of idolatry is found in them? Even concerning those whose opinions and beliefs differ from our opinions and from our faith." Thus far his words.

Literal: seventh foundation and-principle of-our-religion are existence of-Creator one unique eternal who is-not body nor force in-body giving the-Torah to-his-people Israel through the-hand-of the-lawgiver the-faithful Moses servant of-LORD master of-all the-prophets he may-he-be-blessed who-looks-on and-exercises-providence over the-worlds he-created and-his-eyes roam over-all and-searches and-examines the-hearts of-sons of-man and-judges justly according-to the-deeds the-steps-of a-man from-him alone we-seek help in-distress our and-he will-bear-fruit a-shoot from-stock-of Jesse who will-gather outcasts-of Israel from-four corners-of the-earth and-from all the-cut-off corners to-which they-scattered them there the-bones of-their-dead dry will-live and-their-corpses will-rise indeed all these the-principles and-roots-of our-religion and-what trace idolatry is-found in-them even concerning those whose opinions and-beliefs differ from-our-opinions and-from-our-faith thus far his-words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:25.1 | seventh |  |  |  |
| 2:25.2 | foundation |  |  |  |
| 2:25.3 | and principle |  |  |  |
| 2:25.4 | of our religion |  |  |  |
| 2:25.5 | is |  |  |  |
| 2:25.6 | existence |  |  |  |
| 2:25.7 | of Creator |  |  |  |
| 2:25.8 | one |  |  |  |
| 2:25.9 | unique |  |  |  |
| 2:25.10 | eternal |  |  | קדמון: eternal, 'without beginning' |
| 2:25.11 | who |  |  |  |
| 2:25.12 | is not |  |  |  |
| 2:25.13 | a body |  |  |  |
| 2:25.14 | nor |  |  |  |
| 2:25.15 | a force |  |  |  |
| 2:25.16 | in a body |  |  | גשם ולא כח בגשם: the form of the Rambam's principle, 'a body or a force in a body' |
| 2:25.17 | gives |  |  |  |
| 2:25.18 | the Torah |  |  |  |
| 2:25.19 | to His people |  |  |  |
| 2:25.20 | Israel |  |  |  |
| 2:25.21 | through |  |  |  |
| 2:25.22 | the hand of |  |  | על ידי: lit. 'upon the hand of', 'through' |
| 2:25.23 | the lawgiver |  |  |  |
| 2:25.24 | the faithful |  |  |  |
| 2:25.25 | Moses |  |  |  |
| 2:25.26 | servant |  |  |  |
| 2:25.27 | LORD |  |  | ה׳: the Tetragrammaton |
| 2:25.28 | master |  |  |  |
| 2:25.29 | of all |  |  |  |
| 2:25.30 | the prophets |  |  |  |
| 2:25.31 | He |  |  |  |
| 2:25.32 | may He be blessed |  |  | יתברך: 'may He be blessed' |
| 2:25.33 | who looks on |  |  | המציץ: participle, 'who looks on' |
| 2:25.34 | and exercises providence over |  |  | ומשגיח: participle; השגחה is 'providence' |
| 2:25.35 | the worlds |  |  |  |
| 2:25.36 | He created |  |  |  |
| 2:25.37 | and His eyes |  |  |  |
| 2:25.38 | roam |  |  |  |
| 2:25.39 | over all |  |  |  |
| 2:25.40 | and searches |  |  |  |
| 2:25.41 | and examines |  |  |  |
| 2:25.42 | the hearts |  |  |  |
| 2:25.43 | of the sons |  |  |  |
| 2:25.44 | of man |  |  |  |
| 2:25.45 | and judges |  |  |  |
| 2:25.46 | justly |  |  |  |
| 2:25.47 | according to |  |  |  |
| 2:25.48 | the deeds |  |  |  |
| 2:25.49 | the steps of |  |  |  |
| 2:25.50 | a man |  |  |  |
| 2:25.51 | from Him |  |  |  |
| 2:25.52 | alone |  |  |  |
| 2:25.53 | we seek |  |  |  |
| 2:25.54 | help |  |  |  |
| 2:25.55 | in distress |  |  |  |
| 2:25.56 | our |  |  |  |
| 2:25.57 | and He |  |  |  |
| 2:25.58 | will bring forth |  |  | יפרה: future; Isa 11:1 has this verb intransitively, 'bear fruit' |
| 2:25.59 | a shoot |  |  | נצר: 'shoot' or 'branch' |
| 2:25.60 | from the stock of |  |  |  |
| 2:25.61 | Jesse |  |  |  |
| 2:25.62 | who |  |  |  |
| 2:25.63 | will gather |  |  |  |
| 2:25.64 | the outcasts of |  |  | נדחי: 'the dispersed' |
| 2:25.65 | Israel |  |  |  |
| 2:25.66 | from the four |  |  |  |
| 2:25.67 | corners |  |  | כנפות: lit. 'wings', 'corners' |
| 2:25.68 | of the earth |  |  |  |
| 2:25.69 | and from all |  |  |  |
| 2:25.70 | the cut-off |  |  | קצוצי פאה: see context |
| 2:25.71 | corners |  |  |  |
| 2:25.72 | to which |  |  |  |
| 2:25.73 | they scattered |  |  |  |
| 2:25.74 | them |  |  |  |
| 2:25.75 | there |  |  |  |
| 2:25.76 | the bones |  |  |  |
| 2:25.77 | of their dead |  |  |  |
| 2:25.78 | dry |  |  |  |
| 2:25.79 | will live |  |  | תחינה: fem. plural yiqtol, the bones as subject |
| 2:25.80 | and their corpses |  |  |  |
| 2:25.81 | will rise |  |  |  |
| 2:25.82 | indeed |  |  |  |
| 2:25.83 | all |  |  |  |
| 2:25.84 | these |  |  |  |
| 2:25.85 | the principles |  |  |  |
| 2:25.86 | and roots of |  |  |  |
| 2:25.87 | our religion |  |  |  |
| 2:25.88 | and what |  |  |  |
| 2:25.89 | trace |  |  |  |
| 2:25.90 | idolatry |  |  | ע״ז = עבודה זרה, lit. 'strange worship' |
| 2:25.91 | is found |  |  |  |
| 2:25.92 | in them |  |  |  |
| 2:25.93 | even |  |  |  |
| 2:25.94 | concerning |  |  | לדעת: an infinitive used as 'concerning'; the sentence has no verb |
| 2:25.95 | those |  |  |  |
| 2:25.96 | whose |  |  |  |
| 2:25.97 | opinions |  |  |  |
| 2:25.98 | and beliefs |  |  |  |
| 2:25.99 | differ |  |  |  |
| 2:25.100 | from our opinions |  |  |  |
| 2:25.101 | and from our faith |  |  |  |
| 2:25.102 | thus far his words |  |  | עכ״ל: the closing formula of the quotation |

Phrases:
- 2:25.1 2:25.2 2:25.3 2:25.4 | Seventh: the foundation and principle of our religion | seventh foundation and-principle of-our-religion
- 2:25.5 2:25.6 2:25.7 2:25.8 2:25.9 2:25.10 | is the existence of one, unique, eternal Creator, | are existence of-Creator one unique eternal
- 2:25.11 2:25.12 2:25.13 2:25.14 2:25.15 2:25.16 | who is not a body, nor a force in a body, | who is-not body nor force in-body
- 2:25.17 2:25.18 2:25.19 2:25.20 | who gives the Torah to His people Israel | giving the-Torah to-his-people Israel
- 2:25.21 2:25.22 2:25.23 2:25.24 2:25.25 | through the faithful lawgiver Moses, | through the-hand-of the-lawgiver the-faithful Moses
- 2:25.26 2:25.27 2:25.28 2:25.29 2:25.30 | servant of the LORD, master of all the prophets. | servant of LORD master of-all the-prophets
- 2:25.31 2:25.32 2:25.33 2:25.34 2:25.35 2:25.36 | He, may He be blessed, looks on and exercises providence over the worlds that He created. | he may-he-be-blessed who-looks-on and-exercises-providence over the-worlds he-created
- 2:25.37 2:25.38 2:25.39 | And His eyes roam over all, | and-his-eyes roam over-all
- 2:25.40 2:25.41 2:25.42 2:25.43 2:25.44 2:25.45 2:25.46 | He searches and examines the hearts of the sons of man and judges justly | searches and-examines the-hearts of-sons of-man and-judges justly
- 2:25.47 2:25.48 2:25.49 2:25.50 | according to the deeds, the steps of a man. | according-to the-deeds the-steps-of a-man
- 2:25.51 2:25.52 2:25.53 2:25.54 2:25.55 2:25.56 | From Him alone we seek help in our distress, | from-Him alone we-seek help in-distress our
- 2:25.57 2:25.58 2:25.59 2:25.60 2:25.61 | and He will bring forth a shoot from the stock of Jesse, | and-he will-bear-fruit a-shoot from-stock-of Jesse
- 2:25.62 2:25.63 2:25.64 2:25.65 2:25.66 2:25.67 2:25.68 | who will gather the outcasts of Israel from the four corners of the earth, | who will-gather outcasts-of Israel from-four corners-of the-earth
- 2:25.69 2:25.70 2:25.71 2:25.72 2:25.73 2:25.74 2:25.75 | and from all the cut-off corners to which they scattered them there, | and-from all the-cut-off corners to-which they-scattered them there
- 2:25.76 2:25.77 2:25.78 2:25.79 2:25.80 2:25.81 | the dry bones of their dead will live, and their corpses will rise. | the-bones of-their-dead dry will-live and-their-corpses will-rise
- 2:25.82 2:25.83 2:25.84 2:25.85 2:25.86 2:25.87 | Indeed, all these are the principles and roots of our religion. | indeed all these the-principles and-roots-of our-religion
- 2:25.88 2:25.89 2:25.90 2:25.91 2:25.92 | And what trace of idolatry is found in them? | and-what trace idolatry is-found in-them
- 2:25.93 2:25.94 2:25.95 | Even concerning those | even concerning those
- 2:25.96 2:25.97 2:25.98 2:25.99 2:25.100 2:25.101 | whose opinions and beliefs differ from our opinions and from our faith. | whose opinions and-beliefs differ from-our-opinions and-from-our-faith
- 2:25.102 | Thus far his words. | thus far

Grammar notes:
- The paragraph is one sentence: 'the foundation ... is the existence of ...'. Its relative clauses (אשר, נותן) continue from the Creator, and the English keeps the sentence long.
- The verbs of providence (המציץ, ומשגיח, חוקר, ובוחן, ושופט) are participles with no finite verb; the English makes them finite with 'He'.
- והוא יפרה נצר: Isa 11:1 has this verb intransitively ('shall bear fruit'); here it takes נצר as its object, and the English follows the author ('bring forth a shoot').
- The last sentence has no verb ('even concerning those ...'); the English keeps the fragment, as the Hebrew does.

Commentary rulings:
- apparatus: Quotation: the double quotation marks close here with the formula עכ"ל. The opening is outside this unit (see 2:23 apparatus).
- apparatus: Allusions, the author's wording and not quotations: Jer 17:10 (חוקר לב בוחן) for 'searches and examines the hearts'; Prov 20:24 (מצעדי גבר) for 'the steps of a man'; Isa 11:1 for 'a shoot from the stock of Jesse'; Isa 11:12 (מארבע כנפות הארץ) for 'from the four corners of the earth'; Ezek 37 (the dry bones) and Isa 26:19 (נבלתי יקומון) for the resurrection clause; Jer 9:25 for 'cut-off corners'.

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


### Translation under review (2:26)

English: And according to our way, the way of our Rabbis of blessed memory in the aforementioned baraita, there is no need for all this, since our intention is not that we bow down before the Torah scroll, but to Him who chose the Torah. Explicitly, we say: "Come, let us bow down and kneel, let us bless before the LORD our Maker." As for the bowing down of Abraham our father to the angels, and so Joshua's. If you wish, say: since they appeared to them as men, we apply to them what the Rabbis said in the baraita: "'You shall not bow down to them.' To them you may not bow down, but you bow down to a man like yourself," since they were revealed to them in the form of men. And if you wish, say: the intention of Abraham and Joshua was to bow down to Him who sent them, as the author of the Ikkarim wrote.

Literal: according-to our-way the-way-of our-Rabbis-of-blessed-memory in-the-baraita the-aforementioned there-is-no need for-all this since our-intention is-not that-we bow-down before the-Torah-scroll but to-whom chose the-Torah explicitly we say come let-us-bow-down and-kneel let-us-bless before the-LORD our-maker the-bowing-down of-Abraham our-father to-the-angels and-so Joshua if-you-wish say since they-appeared to-them as-men we-apply to-them what they-said the-Rabbis in-the-baraita you-shall-not bow-down to-them to-them you-may-not you bow-down but you bow-down to-a-man like-yourself since they-were-revealed to-them in-the-form of-men and-if-you-wish say the-intention of-Abraham and-Joshua to-bow-down to-whom they-sent-them as the-author-of the-Ikkarim wrote

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:26.1 | and according to |  |  |  |
| 2:26.2 | our way |  |  |  |
| 2:26.3 | the way of |  |  |  |
| 2:26.4 | our Rabbis of blessed memory |  |  | רז״ל: רבותינו זכרונם לברכה |
| 2:26.5 | the baraita |  |  |  |
| 2:26.6 | aforementioned |  |  | הנז״ל: הנזכר |
| 2:26.7 | there is no |  |  |  |
| 2:26.8 | need |  |  |  |
| 2:26.9 | for all |  |  |  |
| 2:26.10 | this |  |  |  |
| 2:26.11 | since not |  |  |  |
| 2:26.12 | our intention |  |  |  |
| 2:26.13 | that we |  |  |  |
| 2:26.14 | bow down |  |  |  |
| 2:26.15 | before |  |  |  |
| 2:26.16 | the Torah scroll |  |  | ס״ת: Sefer Torah, the Torah scroll |
| 2:26.17 | but |  |  |  |
| 2:26.18 | to Him who |  |  | למי does not name him; the context is worship of God |
| 2:26.19 | chose |  |  |  |
| 2:26.20 | the Torah |  |  |  |
| 2:26.21 | explicitly |  |  |  |
| 2:26.22 | we |  |  |  |
| 2:26.23 | say |  |  |  |
| 2:26.24 | come |  |  |  |
| 2:26.25 | let us bow down |  |  |  |
| 2:26.26 | and kneel |  |  |  |
| 2:26.27 | let us bless |  |  |  |
| 2:26.28 | before |  |  |  |
| 2:26.29 | the LORD |  |  | ה׳: the Tetragrammaton (Ps 95:6) |
| 2:26.30 | our Maker |  |  |  |
| 2:26.31 | the bowing down |  |  |  |
| 2:26.32 | of Abraham |  |  |  |
| 2:26.33 | our father |  |  |  |
| 2:26.34 | to the angels |  |  |  |
| 2:26.35 | and so |  |  |  |
| 2:26.36 | Joshua |  |  |  |
| 2:26.37 | if you wish |  |  |  |
| 2:26.38 | say |  |  |  |
| 2:26.39 | since |  |  |  |
| 2:26.40 | they appeared |  |  |  |
| 2:26.41 | to them |  |  |  |
| 2:26.42 | as men |  |  |  |
| 2:26.43 | we apply |  |  |  |
| 2:26.44 | to them |  |  | קרינא בהו: lit. 'we read them', the Talmudic 'we apply the verse to them' |
| 2:26.45 | what |  |  |  |
| 2:26.46 | they said |  |  |  |
| 2:26.47 | the Rabbis |  |  |  |
| 2:26.48 | in the baraita |  |  |  |
| 2:26.49 | you shall not |  |  |  |
| 2:26.50 | bow down |  |  |  |
| 2:26.51 | to them |  |  |  |
| 2:26.52 | to them |  |  |  |
| 2:26.53 | you may not |  |  | אי אתה משתחוה: the Talmudic 'you may not bow down' |
| 2:26.54 | you |  |  |  |
| 2:26.55 | bow down |  |  |  |
| 2:26.56 | but |  |  |  |
| 2:26.57 | you |  |  |  |
| 2:26.58 | bow down |  |  |  |
| 2:26.59 | to a man |  |  |  |
| 2:26.60 | like yourself |  |  |  |
| 2:26.61 | since |  |  |  |
| 2:26.62 | they were revealed |  |  |  |
| 2:26.63 | to them |  |  |  |
| 2:26.64 | in the form |  |  |  |
| 2:26.65 | of men |  |  |  |
| 2:26.66 | and if you wish |  |  |  |
| 2:26.67 | say |  |  |  |
| 2:26.68 | the intention |  |  |  |
| 2:26.69 | of Abraham |  |  |  |
| 2:26.70 | and Joshua |  |  |  |
| 2:26.71 | to bow down |  |  |  |
| 2:26.72 | to Him |  |  | as the note at 'to Him who' above |
| 2:26.73 | who sent them |  |  |  |
| 2:26.74 | as |  |  |  |
| 2:26.75 | wrote |  |  | שכ׳ = שכתב |
| 2:26.76 | the author of |  |  |  |
| 2:26.77 | the Ikkarim |  |  | Sefer ha-Ikkarim, 'Principles'; the author is named only by the title |

Phrases:
- 2:26.1 2:26.2 2:26.3 2:26.4 2:26.5 2:26.6 | And according to our way, the way of our Rabbis of blessed memory in the aforementioned baraita, | according-to our-way the-way-of our-Rabbis in-the-baraita the-aforementioned
- 2:26.7 2:26.8 2:26.9 2:26.10 | there is no need for all this, | there-is-no need for-all this
- 2:26.11 2:26.12 2:26.13 2:26.14 2:26.15 2:26.16 | since our intention is not that we bow down before the Torah scroll, | since our-intention is-not that-we bow-down before the-Torah-scroll
- 2:26.17 2:26.18 2:26.19 2:26.20 | but to Him who chose the Torah. | but to-whom chose the-Torah
- 2:26.21 2:26.22 2:26.23 | Explicitly, we say: | explicitly we say
- 2:26.24 2:26.25 2:26.26 2:26.27 2:26.28 2:26.29 2:26.30 | "Come, let us bow down and kneel, let us bless before the LORD our Maker." | come let-us-bow-down and-kneel let-us-bless before the-LORD our-maker
- 2:26.31 2:26.32 2:26.33 2:26.34 2:26.35 2:26.36 | As for the bowing down of Abraham our father to the angels, and so Joshua's. | the-bowing-down of-Abraham our-father to-the-angels and-so Joshua
- 2:26.37 2:26.38 | If you wish, say: | if-you-wish say
- 2:26.39 2:26.40 2:26.41 2:26.42 2:26.43 2:26.44 | since they appeared to them as men, we apply to them | since they-appeared to-them as-men we-apply to-them
- 2:26.45 2:26.46 2:26.47 2:26.48 | what the Rabbis said in the baraita: | what they-said the-Rabbis in-the-baraita
- 2:26.49 2:26.50 2:26.51 | "'You shall not bow down to them.' | you-shall-not bow-down to-them
- 2:26.52 2:26.53 2:26.54 2:26.55 | To them you may not bow down, | to-them you-may-not you bow-down
- 2:26.56 2:26.57 2:26.58 2:26.59 2:26.60 | but you bow down to a man like yourself." | but you bow-down to-a-man like-yourself
- 2:26.61 2:26.62 2:26.63 2:26.64 2:26.65 | since they were revealed to them in the form of men. | since they-were-revealed to-them in-the-form of-men
- 2:26.66 2:26.67 | And if you wish, say: | and-if-you-wish say
- 2:26.68 2:26.69 2:26.70 2:26.71 2:26.72 2:26.73 | the intention of Abraham and Joshua was to bow down to Him who sent them, | the-intention of-Abraham and-Joshua to-bow-down to-whom they-sent-them
- 2:26.74 2:26.75 2:26.76 2:26.77 | as the author of the Ikkarim wrote. | as the-author-of the-Ikkarim wrote

Grammar notes:
- בואו נשתחוה ונכרעה נברכה is Ps 95:6, given as the people's words: the cohortatives ('let us') stay in the English, and the verse is set in quotation marks.
- והשתחוית is a noun ('the bowing down'), so 'As for the bowing down' is a topic phrase; the English needs no verb until 'and so Joshua's'.
- איבעית אימא ('if you wish, say') introduces a second explanation. The author gives two, the first (the angels appeared as men) and the second (the intention was toward the One who sent them); the English keeps both in order.
- The baraita is quoted inside the sentence ('what the Rabbis said in the baraita: ...'), and the English returns to the author's syntax with 'since they were revealed to them'.
- אי אתה משתחוה is the Talmudic negative of permission, 'you may not bow down'; the English keeps the modal.

Commentary rulings:
- apparatus: Ps 95:6 for 'בואו נשתחוה ונכרעה נברכה לפני ה׳ עושנו'.
- apparatus: Exod 20:5 for 'לא תשתחוה להם' ('you shall not bow down to them'), the words of the baraita.
- apparatus: Cf. Gen 18:2 and Josh 5:14 for the bowing of Abraham and Joshua; the author alludes to them and does not quote them.

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


### Translation under review (2:27)

English: And your question concerning prophecy, although it does not belong to our subject, and you did not ask only to mislead us and to divert us to another matter. Behold, this is set out in the degrees of prophecy in the Rambam, Laws of the Foundations of the Torah, and in part two of the book Moreh Nevukhim, and in the book Emunot ve-De'ot of Rabbi Saadia Gaon. Since I have found you wise in your eyes, I return to ask you to answer me a decisive answer from the words of our Rabbis of blessed memory in the Talmud, for you alone (in your opinion) all mysteries have been revealed to you, and greater understanding, and a new spirit has been given within you. Instruct us, I and my companions, how the matters can agree fully and well. For our soul has longed to hear your teachings, and to delight in your perfect Torah, and it will be to you a merit, and also great honor you will inherit from us.

Literal: and-your-question concerning prophecy although is-not relevant to-our-subject and-not you-asked only to-mislead-us and-to-divert-us to-a-matter other. behold this is-set-out the-degrees-of prophecy in-the-Rambam Laws-of the-Foundations-of the-Torah, and-in part two of-the-book the-Moreh, and-in-the-book the-Emunot and-the-De'ot of-Rabbi-Saadia-Gaon. and-since I-found-you wise in-your-eyes, I-return to-ask of-you to-answer to-me an-answer decisive from-the-words-of our-Rabbis in-the-Talmud, for you alone (in-your-opinion) to-you were-revealed all mysteries, and-understanding greater and-spirit new has-been-given within-you. instruct-us I and-my-companions how the-matters will-agree in-one-reed in-full and-good. for our-soul longed to-hear your-teachings, and-to-delight in-your-Torah the-perfect, and-to-you will-be righteousness, and-also honor great you-will-inherit from-us

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:27.1 | and your question |  |  |  |
| 2:27.2 | concerning |  |  |  |
| 2:27.3 | prophecy |  |  |  |
| 2:27.4 | although |  |  | אעפ״י = אף על פי |
| 2:27.5 | is not |  |  |  |
| 2:27.6 | relevant |  |  |  |
| 2:27.7 | to our subject |  |  |  |
| 2:27.8 | and not |  |  |  |
| 2:27.9 | you asked |  |  |  |
| 2:27.10 | only |  |  |  |
| 2:27.11 | to mislead us |  |  |  |
| 2:27.12 | and to divert us |  |  |  |
| 2:27.13 | to a matter |  |  |  |
| 2:27.14 | other |  |  |  |
| 2:27.15 | behold |  |  |  |
| 2:27.16 | this |  |  |  |
| 2:27.17 | is set out |  |  |  |
| 2:27.18 | the degrees of |  |  |  |
| 2:27.19 | prophecy |  |  |  |
| 2:27.20 | in the Rambam |  |  |  |
| 2:27.21 | Laws of |  |  | הל׳ = הלכות: the Rambam's Mishneh Torah (the code); its first book is 'Laws of the Foundations of the Torah' |
| 2:27.22 | of the Foundations |  |  |  |
| 2:27.23 | of the Torah |  |  |  |
| 2:27.24 | and in part |  |  |  |
| 2:27.25 | two |  |  |  |
| 2:27.26 | of the book |  |  |  |
| 2:27.27 | Moreh Nevukhim |  |  | Moreh Nevukhim, the Guide of the Perplexed; the Hebrew is 'the book of the Moreh' |
| 2:27.28 | and in the book |  |  |  |
| 2:27.29 | Emunot |  |  | Emunot ve-De'ot, 'Beliefs and Opinions' |
| 2:27.30 | and De'ot |  |  |  |
| 2:27.31 | of Rabbi Saadia Gaon |  |  | רס״ג: Rabbi Saadia Gaon |
| 2:27.32 | and since |  |  |  |
| 2:27.33 | I have found you |  |  |  |
| 2:27.34 | wise |  |  |  |
| 2:27.35 | in your eyes |  |  | the idiom 'wise in one's own eyes' (Prov 3:7); the Hebrew keeps both readings |
| 2:27.36 | I return |  |  |  |
| 2:27.37 | to ask |  |  |  |
| 2:27.38 | of you |  |  |  |
| 2:27.39 | to answer |  |  |  |
| 2:27.40 | me |  |  |  |
| 2:27.41 | an answer |  |  |  |
| 2:27.42 | decisive |  |  |  |
| 2:27.43 | from the words of |  |  |  |
| 2:27.44 | our Rabbis of blessed memory |  |  | רז״ל: רבותינו זכרונם לברכה |
| 2:27.45 | in the Talmud |  |  |  |
| 2:27.46 | for |  |  |  |
| 2:27.47 | only |  |  |  |
| 2:27.48 | you |  |  |  |
| 2:27.49 | alone |  |  |  |
| 2:27.50 | in |  |  |  |
| 2:27.51 | your opinion |  |  | the parentheses are the author's and are kept |
| 2:27.52 | to you |  |  |  |
| 2:27.53 | were revealed |  |  |  |
| 2:27.54 | all |  |  |  |
| 2:27.55 | mysteries |  |  |  |
| 2:27.56 | and understanding |  |  |  |
| 2:27.57 | greater |  |  |  |
| 2:27.58 | and a spirit |  |  |  |
| 2:27.59 | new |  |  |  |
| 2:27.60 | has been given |  |  |  |
| 2:27.61 | within you |  |  |  |
| 2:27.62 | instruct us |  |  | imperative, with 'us' as object |
| 2:27.63 | I |  |  |  |
| 2:27.64 | and my companions |  |  |  |
| 2:27.65 | how |  |  |  |
| 2:27.66 | the matters |  |  |  |
| 2:27.67 | can agree |  |  | יעלו בקנה אחד: the idiom 'to agree' |
| 2:27.68 | in agreement |  |  | lit. 'in one reed' |
| 2:27.69 | one |  |  |  |
| 2:27.70 | fully |  |  |  |
| 2:27.71 | and well |  |  |  |
| 2:27.72 | for |  |  |  |
| 2:27.73 | our soul |  |  |  |
| 2:27.74 | has longed |  |  |  |
| 2:27.75 | to hear |  |  |  |
| 2:27.76 | your teachings |  |  |  |
| 2:27.77 | and to delight |  |  |  |
| 2:27.78 | in your Torah |  |  |  |
| 2:27.79 | perfect |  |  |  |
| 2:27.80 | and to you |  |  |  |
| 2:27.81 | will be |  |  |  |
| 2:27.82 | a merit |  |  | צדקה: 'righteousness', here the merit that the act earns |
| 2:27.83 | and also |  |  |  |
| 2:27.84 | honor |  |  |  |
| 2:27.85 | great |  |  |  |
| 2:27.86 | you will inherit |  |  |  |
| 2:27.87 | from us |  |  |  |

Phrases:
- 2:27.1 2:27.2 2:27.3 | And your question concerning prophecy, | and-your-question concerning prophecy
- 2:27.4 2:27.5 2:27.6 2:27.7 | although it does not belong to our subject, | although is-not relevant to-our-subject
- 2:27.8 2:27.9 2:27.10 2:27.11 2:27.12 2:27.13 2:27.14 | and you did not ask only to mislead us and to divert us to another matter. | and-not you-asked only to-mislead-us and-to-divert-us to-a-matter other
- 2:27.15 2:27.16 2:27.17 2:27.18 2:27.19 | Behold, this is set out in the degrees of prophecy | behold this is-set-out the-degrees-of prophecy
- 2:27.20 2:27.21 2:27.22 2:27.23 | in the Rambam, Laws of the Foundations of the Torah, | in-the-Rambam Laws-of the-Foundations-of the-Torah
- 2:27.24 2:27.25 2:27.26 2:27.27 | and in part two of the book Moreh Nevukhim, | and-in part two of-the-book the-Moreh
- 2:27.28 2:27.29 2:27.30 2:27.31 | and in the book Emunot ve-De'ot of Rabbi Saadia Gaon. | and-in-the-book the-Emunot and-the-De'ot of-Rabbi-Saadia-Gaon
- 2:27.32 2:27.33 2:27.34 2:27.35 | Since I have found you wise in your eyes, | and-since I-found-you wise in-your-eyes
- 2:27.36 2:27.37 2:27.38 2:27.39 2:27.40 | I return to ask you to answer me | I return to-ask of-you to-answer me
- 2:27.41 2:27.42 2:27.43 2:27.44 2:27.45 | a decisive answer from the words of our Rabbis of blessed memory in the Talmud, | an-answer decisive from-the-words-of our-Rabbis in-the-Talmud
- 2:27.46 2:27.47 2:27.48 2:27.49 2:27.50 2:27.51 | for you alone (in your opinion) | for only you alone (in-your-opinion)
- 2:27.52 2:27.53 2:27.54 2:27.55 | all mysteries have been revealed to you, | to-you were-revealed all mysteries
- 2:27.56 2:27.57 2:27.58 2:27.59 2:27.60 2:27.61 | and greater understanding, and a new spirit has been given within you. | and-understanding greater and-spirit new has-been-given within-you
- 2:27.62 2:27.63 2:27.64 2:27.65 2:27.66 2:27.67 2:27.68 2:27.69 2:27.70 2:27.71 | Instruct us, I and my companions, how the matters can agree fully and well. | instruct-us I and-my-companions how the-matters will-agree in-one-reed in-full and-good
- 2:27.72 2:27.73 2:27.74 2:27.75 2:27.76 | For our soul has longed to hear your teachings, | for our-soul has-longed to-hear your-teachings
- 2:27.77 2:27.78 2:27.79 | and to delight in your perfect Torah, | and-to-delight in-your-Torah the-perfect
- 2:27.80 2:27.81 2:27.82 | and it will be to you a merit, | and-to-you will-be righteousness
- 2:27.83 2:27.84 2:27.85 2:27.86 2:27.87 | and also great honor you will inherit from us. | and-also honor great you-will-inherit from-us

Grammar notes:
- The paragraph is a single letter sentence, and the list of works (the Rambam, the Moreh, Emunot ve-De'ot) is a series of 'in' phrases; the English keeps the series in the Hebrew's order.
- ולא שאלת רק להטעותנו: 'you did not ask except to mislead us'; the English keeps the restriction with 'only'.
- חכם בעיניך is an idiom that can mean 'wise in your own eyes' (Prov 3:7). The English keeps 'in your eyes', as the Hebrew does.
- הורינו אני וחבירי איך הדברים יעלו: the imperative is followed by a 'how' clause, so the English reads 'Instruct us ..., how the matters can agree'.

Commentary rulings:
- apparatus: Cf. Ezek 36:26 for 'ורוח חדש ניתן בקרבך' ('a new spirit ... within you'); the author does not quote it.
- apparatus: Cf. Prov 3:7 for 'חכם בעיניך' (see the token note at 'in your eyes').

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

