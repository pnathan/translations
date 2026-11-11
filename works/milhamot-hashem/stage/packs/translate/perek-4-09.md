# Translation pack: milhamot-hashem perek-4-09 (Milchamot ha-Shem 4:14–4:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-09`.

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
- **קבלה / המקובלים / בעלי הקבלה**: context.
  - The lore: "Kabbalah" (חכמת הקבלה → "the wisdom of Kabbalah"; הקבלה החדשה
    → "the new Kabbalah", always "new", never "later" or "Lurianic").
  - The people: המקובלים → "the Kabbalists"; המקובלים החדשים → "the new
    Kabbalists"; "רבותינו המקובלים" / "רבותינו בעלי הקבלה החדשה" → "our
    rabbis the Kabbalists" / "our rabbis, the masters of the new Kabbalah".
    Keep the courtesy as written.
  - The author's claim on the word for the rabbinic chain: "בעלי הקבלה
    האמתית אשר קבלו ממרע"ה" (2:34) → "the masters of the true tradition, who
    received from Moses our teacher". Add a note each time: the Hebrew is the
    word for Kabbalah, and he claims it for the Mishnah and Talmud. "אינה
    קבלה ממשה" (3:18) → "it is not a tradition from Moses".
  - The ordinary participle: "וכמקובל אצל כל איש ישראל" (1:11) → "as is
    accepted by every Jew"; "המקובל בידינו" (2:3) → "what we have received".
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

- ריבה / לרבות (the derashic ribbui): "include", never "multiply".
- מורא: "fear" (so that "the fear of Heaven" and Deut 6:13 connect).
- ד' + number after a cited book: keep the numerals ("4, 12") with a note
  when the author writes דף in full nearby.
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
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- Numbers: translate the number as written (ל"ג = thirty-three), whatever the
  tradition counts; the usual count goes in a note (the thirty-two rules of
  R. Yose the Galilean). This overrides the "32" in the מדות entry for 1:6.
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
- עמי הארץ: "the unlearned", as in the project's Avot (עם הארץ, "an unlearned
  person"), the Hebrew term in a note; not "common folk".
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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


## Milchamot ha-Shem 4:14

Hebrew (Wikisource): כז) והוא שדעת הזהר והמקובלים הנמשכים אחריו שכל עבודתינו ותפלותינו לזעיר אנפין דוקא, וז"ל הזהר וירא דף קיב (ח"א קיב, א) תא חזי לה' אלהיך תירא לא כתיב אלא את ה' אלהיך תירא, מאי את דא דרגא קדמאה (פירש בגליון מלכות רמוזה בתיבת את ה' אלהיך תירא) אחר דחלא דקב"ה, ובגין כך כתיב תירא, דתמן בעי בר נש לדחלא קמי מריה בגין דאיהו בי דינא, ואותו תעבוד דא דרגא עלאה דקימא על האיי דרגא תתאה ולא מתפרשין לעלמין, את ואותו דא בדא דבקין ולא אתפרשן, מאי ואותו דא אתר ברית קדישא אות לעלמין, דהא פולחנא לא שריא באת ולא איהו למיפלח אלא למדחל, אבל פולחנא איהו לעילא, ובגין כך ואותו תעבוד (פי' אבל העבודה היא לזעיר אנפין כמו שנ' ואותו תעבוד ותייחד עמיה שכנתיה ועיין בס' המצות בתפלת אליהו הנביא מפורש באורך) ע"כ לשון הזהר עם פי' הרש"ב בגליון. וכונת הזהר לדרוש מדלא כתיב לה' אלהיך תירא ותעבוד, וכתיב את לרבות, ריבה מדת המלכות שהיא נוקבא דזעיר אנפין ליראה, ומדכתיב ואותו תעבוד מיעט אותה מעבודה דדוקא אותו תעבוד ולא אותה (דהיינו מלכות) הרי ריבה אותה ליראה ומיעט אותה מעבודה שהעבודה לזעיר אנפין דוקא.

| id | word | punctuation after |
|---|---|---|
| 4:14.1 | כז | ) |
| 4:14.2 | והוא |  |
| 4:14.3 | שדעת |  |
| 4:14.4 | הזהר |  |
| 4:14.5 | והמקובלים |  |
| 4:14.6 | הנמשכים |  |
| 4:14.7 | אחריו |  |
| 4:14.8 | שכל |  |
| 4:14.9 | עבודתינו |  |
| 4:14.10 | ותפלותינו |  |
| 4:14.11 | לזעיר |  |
| 4:14.12 | אנפין |  |
| 4:14.13 | דוקא | , |
| 4:14.14 | וז"ל |  |
| 4:14.15 | הזהר |  |
| 4:14.16 | וירא |  |
| 4:14.17 | דף |  |
| 4:14.18 | קיב |  |
| 4:14.19 | ח"א |  |
| 4:14.20 | קיב | , |
| 4:14.21 | א | ) |
| 4:14.22 | תא |  |
| 4:14.23 | חזי |  |
| 4:14.24 | לה | ' |
| 4:14.25 | אלהיך |  |
| 4:14.26 | תירא |  |
| 4:14.27 | לא |  |
| 4:14.28 | כתיב |  |
| 4:14.29 | אלא |  |
| 4:14.30 | את |  |
| 4:14.31 | ה | ' |
| 4:14.32 | אלהיך |  |
| 4:14.33 | תירא | , |
| 4:14.34 | מאי |  |
| 4:14.35 | את |  |
| 4:14.36 | דא |  |
| 4:14.37 | דרגא |  |
| 4:14.38 | קדמאה |  |
| 4:14.39 | פירש |  |
| 4:14.40 | בגליון |  |
| 4:14.41 | מלכות |  |
| 4:14.42 | רמוזה |  |
| 4:14.43 | בתיבת |  |
| 4:14.44 | את |  |
| 4:14.45 | ה | ' |
| 4:14.46 | אלהיך |  |
| 4:14.47 | תירא | ) |
| 4:14.48 | אחר |  |
| 4:14.49 | דחלא |  |
| 4:14.50 | דקב"ה | , |
| 4:14.51 | ובגין |  |
| 4:14.52 | כך |  |
| 4:14.53 | כתיב |  |
| 4:14.54 | תירא | , |
| 4:14.55 | דתמן |  |
| 4:14.56 | בעי |  |
| 4:14.57 | בר |  |
| 4:14.58 | נש |  |
| 4:14.59 | לדחלא |  |
| 4:14.60 | קמי |  |
| 4:14.61 | מריה |  |
| 4:14.62 | בגין |  |
| 4:14.63 | דאיהו |  |
| 4:14.64 | בי |  |
| 4:14.65 | דינא | , |
| 4:14.66 | ואותו |  |
| 4:14.67 | תעבוד |  |
| 4:14.68 | דא |  |
| 4:14.69 | דרגא |  |
| 4:14.70 | עלאה |  |
| 4:14.71 | דקימא |  |
| 4:14.72 | על |  |
| 4:14.73 | האיי |  |
| 4:14.74 | דרגא |  |
| 4:14.75 | תתאה |  |
| 4:14.76 | ולא |  |
| 4:14.77 | מתפרשין |  |
| 4:14.78 | לעלמין | , |
| 4:14.79 | את |  |
| 4:14.80 | ואותו |  |
| 4:14.81 | דא |  |
| 4:14.82 | בדא |  |
| 4:14.83 | דבקין |  |
| 4:14.84 | ולא |  |
| 4:14.85 | אתפרשן | , |
| 4:14.86 | מאי |  |
| 4:14.87 | ואותו |  |
| 4:14.88 | דא |  |
| 4:14.89 | אתר |  |
| 4:14.90 | ברית |  |
| 4:14.91 | קדישא |  |
| 4:14.92 | אות |  |
| 4:14.93 | לעלמין | , |
| 4:14.94 | דהא |  |
| 4:14.95 | פולחנא |  |
| 4:14.96 | לא |  |
| 4:14.97 | שריא |  |
| 4:14.98 | באת |  |
| 4:14.99 | ולא |  |
| 4:14.100 | איהו |  |
| 4:14.101 | למיפלח |  |
| 4:14.102 | אלא |  |
| 4:14.103 | למדחל | , |
| 4:14.104 | אבל |  |
| 4:14.105 | פולחנא |  |
| 4:14.106 | איהו |  |
| 4:14.107 | לעילא | , |
| 4:14.108 | ובגין |  |
| 4:14.109 | כך |  |
| 4:14.110 | ואותו |  |
| 4:14.111 | תעבוד |  |
| 4:14.112 | פי | ' |
| 4:14.113 | אבל |  |
| 4:14.114 | העבודה |  |
| 4:14.115 | היא |  |
| 4:14.116 | לזעיר |  |
| 4:14.117 | אנפין |  |
| 4:14.118 | כמו |  |
| 4:14.119 | שנ | ' |
| 4:14.120 | ואותו |  |
| 4:14.121 | תעבוד |  |
| 4:14.122 | ותייחד |  |
| 4:14.123 | עמיה |  |
| 4:14.124 | שכנתיה |  |
| 4:14.125 | ועיין |  |
| 4:14.126 | בס | ' |
| 4:14.127 | המצות |  |
| 4:14.128 | בתפלת |  |
| 4:14.129 | אליהו |  |
| 4:14.130 | הנביא |  |
| 4:14.131 | מפורש |  |
| 4:14.132 | באורך | ) |
| 4:14.133 | ע"כ |  |
| 4:14.134 | לשון |  |
| 4:14.135 | הזהר |  |
| 4:14.136 | עם |  |
| 4:14.137 | פי | ' |
| 4:14.138 | הרש"ב |  |
| 4:14.139 | בגליון | . |
| 4:14.140 | וכונת |  |
| 4:14.141 | הזהר |  |
| 4:14.142 | לדרוש |  |
| 4:14.143 | מדלא |  |
| 4:14.144 | כתיב |  |
| 4:14.145 | לה | ' |
| 4:14.146 | אלהיך |  |
| 4:14.147 | תירא |  |
| 4:14.148 | ותעבוד | , |
| 4:14.149 | וכתיב |  |
| 4:14.150 | את |  |
| 4:14.151 | לרבות | , |
| 4:14.152 | ריבה |  |
| 4:14.153 | מדת |  |
| 4:14.154 | המלכות |  |
| 4:14.155 | שהיא |  |
| 4:14.156 | נוקבא |  |
| 4:14.157 | דזעיר |  |
| 4:14.158 | אנפין |  |
| 4:14.159 | ליראה | , |
| 4:14.160 | ומדכתיב |  |
| 4:14.161 | ואותו |  |
| 4:14.162 | תעבוד |  |
| 4:14.163 | מיעט |  |
| 4:14.164 | אותה |  |
| 4:14.165 | מעבודה |  |
| 4:14.166 | דדוקא |  |
| 4:14.167 | אותו |  |
| 4:14.168 | תעבוד |  |
| 4:14.169 | ולא |  |
| 4:14.170 | אותה |  |
| 4:14.171 | דהיינו |  |
| 4:14.172 | מלכות | ) |
| 4:14.173 | הרי |  |
| 4:14.174 | ריבה |  |
| 4:14.175 | אותה |  |
| 4:14.176 | ליראה |  |
| 4:14.177 | ומיעט |  |
| 4:14.178 | אותה |  |
| 4:14.179 | מעבודה |  |
| 4:14.180 | שהעבודה |  |
| 4:14.181 | לזעיר |  |
| 4:14.182 | אנפין |  |
| 4:14.183 | דוקא | . |


## Milchamot ha-Shem 4:15

Hebrew (Wikisource): כח) ומכאן יש להוכיח שאינה קבלה ממשה רבינו כמו שעלה על דעתך. שאם כן שהיא קבלה ממשה. למה פירש שמעון העמסוני או נחמיה העמסוני מלדרוש האתין שבתורה כשהגיע לאת ה' אלהיך תירא כדגר' בפ' כל שעה, תניא שמעון העמסוני ואמרי לה נחמיה העמסוני היה דורש כל אתין שבתורה, כיון שהגיע לאת ה' אלהיך תירא פירש אמרו לו תלמידיו ר' כל אתין שדרשת מה תהא עליהן. אמר להם כשם שקיבלתי שכר על הדרישה כך מקבל אני שכר על הפרישה, עד שבא ר' עקיבה ודרש את ה' אלהיך תירא לרבות תלמידי חכמים, ופי' רש"י ז"ל כיון שהגיע לאת ה' אלהיך תירא, אמר מה ארבה עמו לירא, פירש מכולן וחזר בו מכל רבויין שדרש, עד שבא ר' עקיבה ודרש את לרבות ת"ח שיהא מורא רבך כמורא שמים ע"כ לשון הגמ' ורש"י ז"ל. ואם כדבריך שהיא הלכה למשה מסיני. מה לו לשמעון העמסוני לפרוש? וכי לא ידע לומר לרבות אשתו של זעיר אנפין שהוא ה' אלהינו לפי דבריהם? וכמו שדרשו רז"ל בכתובות תניא כבד את אביך ואת אמך, את אביך זו אשת אביך ואת אמך זו בעל אמך, וא"ו יתירה לרבות אחיך הגדול ע"כ.

| id | word | punctuation after |
|---|---|---|
| 4:15.1 | כח | ) |
| 4:15.2 | ומכאן |  |
| 4:15.3 | יש |  |
| 4:15.4 | להוכיח |  |
| 4:15.5 | שאינה |  |
| 4:15.6 | קבלה |  |
| 4:15.7 | ממשה |  |
| 4:15.8 | רבינו |  |
| 4:15.9 | כמו |  |
| 4:15.10 | שעלה |  |
| 4:15.11 | על |  |
| 4:15.12 | דעתך | . |
| 4:15.13 | שאם |  |
| 4:15.14 | כן |  |
| 4:15.15 | שהיא |  |
| 4:15.16 | קבלה |  |
| 4:15.17 | ממשה | . |
| 4:15.18 | למה |  |
| 4:15.19 | פירש |  |
| 4:15.20 | שמעון |  |
| 4:15.21 | העמסוני |  |
| 4:15.22 | או |  |
| 4:15.23 | נחמיה |  |
| 4:15.24 | העמסוני |  |
| 4:15.25 | מלדרוש |  |
| 4:15.26 | האתין |  |
| 4:15.27 | שבתורה |  |
| 4:15.28 | כשהגיע |  |
| 4:15.29 | לאת |  |
| 4:15.30 | ה | ' |
| 4:15.31 | אלהיך |  |
| 4:15.32 | תירא |  |
| 4:15.33 | כדגר | ' |
| 4:15.34 | בפ | ' |
| 4:15.35 | כל |  |
| 4:15.36 | שעה | , |
| 4:15.37 | תניא |  |
| 4:15.38 | שמעון |  |
| 4:15.39 | העמסוני |  |
| 4:15.40 | ואמרי |  |
| 4:15.41 | לה |  |
| 4:15.42 | נחמיה |  |
| 4:15.43 | העמסוני |  |
| 4:15.44 | היה |  |
| 4:15.45 | דורש |  |
| 4:15.46 | כל |  |
| 4:15.47 | אתין |  |
| 4:15.48 | שבתורה | , |
| 4:15.49 | כיון |  |
| 4:15.50 | שהגיע |  |
| 4:15.51 | לאת |  |
| 4:15.52 | ה | ' |
| 4:15.53 | אלהיך |  |
| 4:15.54 | תירא |  |
| 4:15.55 | פירש |  |
| 4:15.56 | אמרו |  |
| 4:15.57 | לו |  |
| 4:15.58 | תלמידיו |  |
| 4:15.59 | ר | ' |
| 4:15.60 | כל |  |
| 4:15.61 | אתין |  |
| 4:15.62 | שדרשת |  |
| 4:15.63 | מה |  |
| 4:15.64 | תהא |  |
| 4:15.65 | עליהן | . |
| 4:15.66 | אמר |  |
| 4:15.67 | להם |  |
| 4:15.68 | כשם |  |
| 4:15.69 | שקיבלתי |  |
| 4:15.70 | שכר |  |
| 4:15.71 | על |  |
| 4:15.72 | הדרישה |  |
| 4:15.73 | כך |  |
| 4:15.74 | מקבל |  |
| 4:15.75 | אני |  |
| 4:15.76 | שכר |  |
| 4:15.77 | על |  |
| 4:15.78 | הפרישה | , |
| 4:15.79 | עד |  |
| 4:15.80 | שבא |  |
| 4:15.81 | ר | ' |
| 4:15.82 | עקיבה |  |
| 4:15.83 | ודרש |  |
| 4:15.84 | את |  |
| 4:15.85 | ה | ' |
| 4:15.86 | אלהיך |  |
| 4:15.87 | תירא |  |
| 4:15.88 | לרבות |  |
| 4:15.89 | תלמידי |  |
| 4:15.90 | חכמים | , |
| 4:15.91 | ופי | ' |
| 4:15.92 | רש"י |  |
| 4:15.93 | ז"ל |  |
| 4:15.94 | כיון |  |
| 4:15.95 | שהגיע |  |
| 4:15.96 | לאת |  |
| 4:15.97 | ה | ' |
| 4:15.98 | אלהיך |  |
| 4:15.99 | תירא | , |
| 4:15.100 | אמר |  |
| 4:15.101 | מה |  |
| 4:15.102 | ארבה |  |
| 4:15.103 | עמו |  |
| 4:15.104 | לירא | , |
| 4:15.105 | פירש |  |
| 4:15.106 | מכולן |  |
| 4:15.107 | וחזר |  |
| 4:15.108 | בו |  |
| 4:15.109 | מכל |  |
| 4:15.110 | רבויין |  |
| 4:15.111 | שדרש | , |
| 4:15.112 | עד |  |
| 4:15.113 | שבא |  |
| 4:15.114 | ר | ' |
| 4:15.115 | עקיבה |  |
| 4:15.116 | ודרש |  |
| 4:15.117 | את |  |
| 4:15.118 | לרבות |  |
| 4:15.119 | ת"ח |  |
| 4:15.120 | שיהא |  |
| 4:15.121 | מורא |  |
| 4:15.122 | רבך |  |
| 4:15.123 | כמורא |  |
| 4:15.124 | שמים |  |
| 4:15.125 | ע"כ |  |
| 4:15.126 | לשון |  |
| 4:15.127 | הגמ | ' |
| 4:15.128 | ורש"י |  |
| 4:15.129 | ז"ל | . |
| 4:15.130 | ואם |  |
| 4:15.131 | כדבריך |  |
| 4:15.132 | שהיא |  |
| 4:15.133 | הלכה |  |
| 4:15.134 | למשה |  |
| 4:15.135 | מסיני | . |
| 4:15.136 | מה |  |
| 4:15.137 | לו |  |
| 4:15.138 | לשמעון |  |
| 4:15.139 | העמסוני |  |
| 4:15.140 | לפרוש | ? |
| 4:15.141 | וכי |  |
| 4:15.142 | לא |  |
| 4:15.143 | ידע |  |
| 4:15.144 | לומר |  |
| 4:15.145 | לרבות |  |
| 4:15.146 | אשתו |  |
| 4:15.147 | של |  |
| 4:15.148 | זעיר |  |
| 4:15.149 | אנפין |  |
| 4:15.150 | שהוא |  |
| 4:15.151 | ה | ' |
| 4:15.152 | אלהינו |  |
| 4:15.153 | לפי |  |
| 4:15.154 | דבריהם | ? |
| 4:15.155 | וכמו |  |
| 4:15.156 | שדרשו |  |
| 4:15.157 | רז"ל |  |
| 4:15.158 | בכתובות |  |
| 4:15.159 | תניא |  |
| 4:15.160 | כבד |  |
| 4:15.161 | את |  |
| 4:15.162 | אביך |  |
| 4:15.163 | ואת |  |
| 4:15.164 | אמך | , |
| 4:15.165 | את |  |
| 4:15.166 | אביך |  |
| 4:15.167 | זו |  |
| 4:15.168 | אשת |  |
| 4:15.169 | אביך |  |
| 4:15.170 | ואת |  |
| 4:15.171 | אמך |  |
| 4:15.172 | זו |  |
| 4:15.173 | בעל |  |
| 4:15.174 | אמך | , |
| 4:15.175 | וא"ו |  |
| 4:15.176 | יתירה |  |
| 4:15.177 | לרבות |  |
| 4:15.178 | אחיך |  |
| 4:15.179 | הגדול |  |
| 4:15.180 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

