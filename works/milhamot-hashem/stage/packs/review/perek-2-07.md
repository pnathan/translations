# Review pack: milhamot-hashem perek-2-07 (Milchamot ha-Shem 2:13–2:15)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-07`.

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
- **אדם קדמון** → "Adam Kadmon" (the note: "Primordial Man"). **אדם קדמאה**
  → "Adam Kadma'ah". These are two distinct entities in the author's list
  ("א"ס ואדם קדמאה ואדם קדמון ועתיק…", 2:11): never merge them.
- **שכינה**: fixed, "the Shekhinah". Not "the Divine Presence": the author
  argues (2:13–2:15) that the word has three senses (created light,
  revelation, providence), and "Presence" decides the question for him.
  "גילוי שכינה" → "the revelation of the Shekhinah"; "תחת כנפי השכינה" →
  "under the wings of the Shekhinah"; "שכינתיה" in the Kabbalists' formula →
  "His Shekhinah".
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
- שלשה פנים / ל before a person's name as possessive: ל = "of" / "belonging to"
  (לרשב"י = Rabbi Shimon's); never read as the subject.
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
- דברי תימה: "astonishing words" (words that astonish).
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
- The author's own section labels (ו), ז) at a section's head, numbered through
  the book): "§6."; list markers inside a sentence stay "1)", "2)".
- אחד ... ואחד in "one ... and one" about two kinds of case: "alike" ("civil
  cases and capital cases alike").
- ואין צריך לומר: "needless to say"; עבודה (worship) is never "service".
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


## Milchamot ha-Shem 2:13

Hebrew (Wikisource): ועל שאלותיך שכתבת שיש לשאול על מה שכתב הרמב"ם ז"ל: יכווין לבו כנגד השכינה ויתפלל, ובמורה נבוכים כתב שהשכינה אור נברא וכו' הן בשאלתך גילית שאתה חכם בעיניך ולא למדת מוסר. וכי דרך הנשאל לשאול את שואלו? מכל מקום לא אחשוך קולמוסי מלהשיב אעפ"י שעשית לא כהוגן! ואומר כי פשוט הוא וידוע לכל משכיל בקי בספרי ישראל ששם שכינה נאמר על ג' דברים: או אל השי"ת, כמו שכתב רבינו בהל' תשובה פ"ז: גדולה תשובה שמקרבת האדם לשכינה שנאמר: שובה ישראל עד ה' אלהיך וכו', ונאמר: ולא שבתם עדי נאם ה', ונאמר: אם תשוב ישראל אלי תשוב. כלומר, אם תחזור בתשובה, בי תדבק וכו'. עוד שם על מעלת התשובה: אמש היה זה מובדל מה' אלהי ישראל וכו', והיום הוא מודבק בשכינה, שנאמר: ואתם הדבקים בה' אלהיכם, צועק ונענה מיד, שנאמר: והיה טרם יקראו ואני אענה. וכן ענין דבריו בהל' תפלה, שאם היה במקום שאינו יכול לכוון את הרוחות וכו'. וכן העתיקו קדמונינו, ואן כאן פי בוצע ולא יחסן יערף אלגהת יהי"י קלבה אלי כאלקה ויצלי.

| id | word | punctuation after |
|---|---|---|
| 2:13.1 | ועל |  |
| 2:13.2 | שאלותיך |  |
| 2:13.3 | שכתבת |  |
| 2:13.4 | שיש |  |
| 2:13.5 | לשאול |  |
| 2:13.6 | על |  |
| 2:13.7 | מה |  |
| 2:13.8 | שכתב |  |
| 2:13.9 | הרמב"ם |  |
| 2:13.10 | ז"ל | : |
| 2:13.11 | יכווין |  |
| 2:13.12 | לבו |  |
| 2:13.13 | כנגד |  |
| 2:13.14 | השכינה |  |
| 2:13.15 | ויתפלל | , |
| 2:13.16 | ובמורה |  |
| 2:13.17 | נבוכים |  |
| 2:13.18 | כתב |  |
| 2:13.19 | שהשכינה |  |
| 2:13.20 | אור |  |
| 2:13.21 | נברא |  |
| 2:13.22 | וכו | ' |
| 2:13.23 | הן |  |
| 2:13.24 | בשאלתך |  |
| 2:13.25 | גילית |  |
| 2:13.26 | שאתה |  |
| 2:13.27 | חכם |  |
| 2:13.28 | בעיניך |  |
| 2:13.29 | ולא |  |
| 2:13.30 | למדת |  |
| 2:13.31 | מוסר | . |
| 2:13.32 | וכי |  |
| 2:13.33 | דרך |  |
| 2:13.34 | הנשאל |  |
| 2:13.35 | לשאול |  |
| 2:13.36 | את |  |
| 2:13.37 | שואלו | ? |
| 2:13.38 | מכל |  |
| 2:13.39 | מקום |  |
| 2:13.40 | לא |  |
| 2:13.41 | אחשוך |  |
| 2:13.42 | קולמוסי |  |
| 2:13.43 | מלהשיב |  |
| 2:13.44 | אעפ"י |  |
| 2:13.45 | שעשית |  |
| 2:13.46 | לא |  |
| 2:13.47 | כהוגן | ! |
| 2:13.48 | ואומר |  |
| 2:13.49 | כי |  |
| 2:13.50 | פשוט |  |
| 2:13.51 | הוא |  |
| 2:13.52 | וידוע |  |
| 2:13.53 | לכל |  |
| 2:13.54 | משכיל |  |
| 2:13.55 | בקי |  |
| 2:13.56 | בספרי |  |
| 2:13.57 | ישראל |  |
| 2:13.58 | ששם |  |
| 2:13.59 | שכינה |  |
| 2:13.60 | נאמר |  |
| 2:13.61 | על |  |
| 2:13.62 | ג | ' |
| 2:13.63 | דברים | : |
| 2:13.64 | או |  |
| 2:13.65 | אל |  |
| 2:13.66 | השי"ת | , |
| 2:13.67 | כמו |  |
| 2:13.68 | שכתב |  |
| 2:13.69 | רבינו |  |
| 2:13.70 | בהל | ' |
| 2:13.71 | תשובה |  |
| 2:13.72 | פ"ז | : |
| 2:13.73 | גדולה |  |
| 2:13.74 | תשובה |  |
| 2:13.75 | שמקרבת |  |
| 2:13.76 | האדם |  |
| 2:13.77 | לשכינה |  |
| 2:13.78 | שנאמר | : |
| 2:13.79 | שובה |  |
| 2:13.80 | ישראל |  |
| 2:13.81 | עד |  |
| 2:13.82 | ה | ' |
| 2:13.83 | אלהיך |  |
| 2:13.84 | וכו | ', |
| 2:13.85 | ונאמר | : |
| 2:13.86 | ולא |  |
| 2:13.87 | שבתם |  |
| 2:13.88 | עדי |  |
| 2:13.89 | נאם |  |
| 2:13.90 | ה | ', |
| 2:13.91 | ונאמר | : |
| 2:13.92 | אם |  |
| 2:13.93 | תשוב |  |
| 2:13.94 | ישראל |  |
| 2:13.95 | אלי |  |
| 2:13.96 | תשוב | . |
| 2:13.97 | כלומר | , |
| 2:13.98 | אם |  |
| 2:13.99 | תחזור |  |
| 2:13.100 | בתשובה | , |
| 2:13.101 | בי |  |
| 2:13.102 | תדבק |  |
| 2:13.103 | וכו | '. |
| 2:13.104 | עוד |  |
| 2:13.105 | שם |  |
| 2:13.106 | על |  |
| 2:13.107 | מעלת |  |
| 2:13.108 | התשובה | : |
| 2:13.109 | אמש |  |
| 2:13.110 | היה |  |
| 2:13.111 | זה |  |
| 2:13.112 | מובדל |  |
| 2:13.113 | מה | ' |
| 2:13.114 | אלהי |  |
| 2:13.115 | ישראל |  |
| 2:13.116 | וכו | ', |
| 2:13.117 | והיום |  |
| 2:13.118 | הוא |  |
| 2:13.119 | מודבק |  |
| 2:13.120 | בשכינה | , |
| 2:13.121 | שנאמר | : |
| 2:13.122 | ואתם |  |
| 2:13.123 | הדבקים |  |
| 2:13.124 | בה | ' |
| 2:13.125 | אלהיכם | , |
| 2:13.126 | צועק |  |
| 2:13.127 | ונענה |  |
| 2:13.128 | מיד | , |
| 2:13.129 | שנאמר | : |
| 2:13.130 | והיה |  |
| 2:13.131 | טרם |  |
| 2:13.132 | יקראו |  |
| 2:13.133 | ואני |  |
| 2:13.134 | אענה | . |
| 2:13.135 | וכן |  |
| 2:13.136 | ענין |  |
| 2:13.137 | דבריו |  |
| 2:13.138 | בהל | ' |
| 2:13.139 | תפלה | , |
| 2:13.140 | שאם |  |
| 2:13.141 | היה |  |
| 2:13.142 | במקום |  |
| 2:13.143 | שאינו |  |
| 2:13.144 | יכול |  |
| 2:13.145 | לכוון |  |
| 2:13.146 | את |  |
| 2:13.147 | הרוחות |  |
| 2:13.148 | וכו | '. |
| 2:13.149 | וכן |  |
| 2:13.150 | העתיקו |  |
| 2:13.151 | קדמונינו | , |
| 2:13.152 | ואן |  |
| 2:13.153 | כאן |  |
| 2:13.154 | פי |  |
| 2:13.155 | בוצע |  |
| 2:13.156 | ולא |  |
| 2:13.157 | יחסן |  |
| 2:13.158 | יערף |  |
| 2:13.159 | אלגהת |  |
| 2:13.160 | יהי"י |  |
| 2:13.161 | קלבה |  |
| 2:13.162 | אלי |  |
| 2:13.163 | כאלקה |  |
| 2:13.164 | ויצלי | . |


### Translation under review (2:13)

English: As for your questions that you wrote, that one should ask about what the Rambam of blessed memory wrote: "Let him direct his heart toward the Shekhinah and pray," and in Moreh Nevukhim he wrote that the Shekhinah is a created light, etc. Behold, in your question you have revealed that you are wise in your own eyes and have not learned discipline. Is it, then, the way of one who is asked to question his questioner? In any case I will not withhold my quill from replying, even though you have not acted properly! And I say that it is plain and known to every learned man well versed in the books of Israel that the name Shekhinah is said of 3 things: either of the Name, may He be blessed, as our Master wrote in Laws of Repentance, chapter 7: "Great is repentance, which brings a person near to the Shekhinah, as it is said: 'Return, O Israel, to the LORD your God,' etc., and it is said: 'Yet you did not return to Me, says the LORD,' and it is said: 'If you return, O Israel, to Me you shall return.' That is to say, if you turn in repentance, you will cleave to Me, etc." Again there, on the excellence of repentance: "Yesterday this man was separated from the LORD, the God of Israel, etc., and today he is cleaving to the Shekhinah, as it is said: 'And you who cleave to the LORD your God,' he cries out and is answered at once, as it is said: 'And it shall be that before they call, I will answer.'" And so is the sense of his words in Laws of Prayer, that if he was in a place where he cannot determine the directions, etc. And so our forebears rendered it: "And if he is in a place and cannot tell the direction, let him direct his heart to his Creator and pray."

Literal: and-on questions-your that-you-wrote that-there-is to-ask on what that-wrote the-Rambam of-blessed-memory he-shall-direct his-heart toward the-Shekhinah and-he-shall-pray and-in-Moreh Nevukhim he-wrote that-the-Shekhinah light created etc. behold in-your-question you-revealed that-you wise in-your-eyes and-not you-learned discipline and-is the-way-of the-one-asked to-ask [object-marker] his-asker from-every place not I-will-withhold my-quill from-replying although that-you-did not properly and-I-say that plain it and-known to-every learned versed in-books-of Israel that-name Shekhinah is-said on 3 things either to the-Name-may-He-be-blessed as wrote our-Master in-Laws-of Repentance chapter-7 great repentance that-brings-near the-person to-the-Shekhinah as-it-is-said return Israel to the-LORD your-God etc. and-it-is-said and-not you-returned to-me declaration-of the-LORD and-it-is-said if you-return Israel to-me you-shall-return that-is-to-say if you-turn in-repentance in-me you-will-cleave etc. again there on excellence-of the-repentance yesterday was this-one separated from-the-LORD God-of Israel etc. and-today he cleaving to-the-Shekhinah as-it-is-said and-you the-cleaving to-the-LORD your-God crying-out and-answered at-once as-it-is-said and-it-shall-be before they-call and-I I-will-answer and-so sense-of his-words in-Laws-of Prayer that-if he-was in-a-place that-not able to-direct [object-marker] the-directions etc. and-so rendered our-forebears and-if he-is in a-place and-not he-is-able he-knows the-direction he-directs his-heart to his-Creator and-he-prays

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:13.1 | and on |  |  |  |
| 2:13.2 | your questions |  |  |  |
| 2:13.3 | that you wrote |  |  |  |
| 2:13.4 | that there is |  |  | idiom: "there is to ask" means "one should ask" |
| 2:13.5 | to ask |  |  |  |
| 2:13.6 | on |  |  |  |
| 2:13.7 | what |  |  |  |
| 2:13.8 | that wrote |  |  |  |
| 2:13.9 | the Rambam |  |  | Rabbi Moshe ben Maimon; the Rambam is the form English uses for him |
| 2:13.10 | of blessed memory |  |  | abbreviation of the honorific; expanded |
| 2:13.11 | let him direct |  |  | yiqtol used as a rule: "he shall direct" |
| 2:13.12 | his heart |  |  |  |
| 2:13.13 | toward |  |  |  |
| 2:13.14 | the Shekhinah |  |  |  |
| 2:13.15 | and pray |  |  |  |
| 2:13.16 | and in Moreh |  |  | title of the Rambam's book, the Guide of the Perplexed (moreh = guide, nevukhim = the perplexed) |
| 2:13.17 | Nevukhim |  |  |  |
| 2:13.18 | he wrote |  |  |  |
| 2:13.19 | that the Shekhinah |  |  |  |
| 2:13.20 | light |  |  |  |
| 2:13.21 | created |  |  | the author's central word: the Shekhinah is created (nivra) light |
| 2:13.22 | etc. |  |  | the author truncates the citation; nothing is supplied |
| 2:13.23 | behold |  |  |  |
| 2:13.24 | in your question |  |  |  |
| 2:13.25 | you have revealed |  |  |  |
| 2:13.26 | that you |  |  |  |
| 2:13.27 | wise |  |  |  |
| 2:13.28 | in your own eyes |  |  | cf. Prov 3:7 |
| 2:13.29 | and not |  |  |  |
| 2:13.30 | you have learned |  |  |  |
| 2:13.31 | discipline |  |  | mussar: instruction, correction |
| 2:13.32 | and is it |  |  |  |
| 2:13.33 | the way of |  |  |  |
| 2:13.34 | the one asked |  |  |  |
| 2:13.35 | to ask |  |  |  |
| 2:13.36 | (object marker) |  |  |  |
| 2:13.37 | his questioner |  |  |  |
| 2:13.38 | from every |  |  | with the next word, the idiom "in any case" |
| 2:13.39 | place |  |  |  |
| 2:13.40 | not |  |  |  |
| 2:13.41 | I will withhold |  |  | the verb is chasakh, to restrain |
| 2:13.42 | my quill |  |  | kolmos, from Greek kalamos, a reed pen |
| 2:13.43 | from replying |  |  |  |
| 2:13.44 | although |  |  | abbreviation for af al pi she-, expanded |
| 2:13.45 | that you did |  |  |  |
| 2:13.46 | not |  |  |  |
| 2:13.47 | properly |  |  | ke-hogen, "as is fitting" |
| 2:13.48 | and I say |  |  |  |
| 2:13.49 | that |  |  |  |
| 2:13.50 | plain |  |  |  |
| 2:13.51 | it |  |  |  |
| 2:13.52 | and known |  |  |  |
| 2:13.53 | to every |  |  |  |
| 2:13.54 | learned man |  |  |  |
| 2:13.55 | well versed |  |  |  |
| 2:13.56 | in the books of |  |  |  |
| 2:13.57 | Israel |  |  |  |
| 2:13.58 | that the name |  |  |  |
| 2:13.59 | Shekhinah |  |  |  |
| 2:13.60 | is said |  |  |  |
| 2:13.61 | of |  |  |  |
| 2:13.62 | 3 |  |  | letter numeral gimel |
| 2:13.63 | things |  |  |  |
| 2:13.64 | either |  |  |  |
| 2:13.65 | of |  |  | literally "to" |
| 2:13.66 | the Name, may He be blessed |  |  | abbreviation ha-Shem yitbarakh, expanded |
| 2:13.67 | as |  |  |  |
| 2:13.68 | wrote |  |  |  |
| 2:13.69 | our Master |  |  | rabbenu: the Rambam |
| 2:13.70 | in Laws of |  |  | abbreviation of hilkhot |
| 2:13.71 | Repentance |  |  |  |
| 2:13.72 | chapter 7 |  |  | abbreviation: perek zayin |
| 2:13.73 | great |  |  |  |
| 2:13.74 | repentance |  |  |  |
| 2:13.75 | which brings near |  |  | participle |
| 2:13.76 | the person |  |  |  |
| 2:13.77 | to the Shekhinah |  |  |  |
| 2:13.78 | as it is said |  |  | niphal; the formula that introduces a verse |
| 2:13.79 | return |  |  | Hos 14:2 |
| 2:13.80 | Israel |  |  |  |
| 2:13.81 | to |  |  |  |
| 2:13.82 | the LORD |  |  |  |
| 2:13.83 | your God |  |  |  |
| 2:13.84 | etc. |  |  |  |
| 2:13.85 | and it is said |  |  |  |
| 2:13.86 | and not |  |  |  |
| 2:13.87 | you returned |  |  |  |
| 2:13.88 | to Me |  |  | adey, "to me" (Amos 4:6) |
| 2:13.89 | says |  |  |  |
| 2:13.90 | the LORD |  |  |  |
| 2:13.91 | and it is said |  |  |  |
| 2:13.92 | if |  |  |  |
| 2:13.93 | you return |  |  |  |
| 2:13.94 | Israel |  |  |  |
| 2:13.95 | to Me |  |  |  |
| 2:13.96 | you shall return |  |  |  |
| 2:13.97 | that is to say |  |  | kelomar |
| 2:13.98 | if |  |  |  |
| 2:13.99 | you turn |  |  |  |
| 2:13.100 | in repentance |  |  |  |
| 2:13.101 | to Me |  |  | literally "in Me" |
| 2:13.102 | you will cleave |  |  |  |
| 2:13.103 | etc. |  |  |  |
| 2:13.104 | again |  |  |  |
| 2:13.105 | there |  |  |  |
| 2:13.106 | on |  |  |  |
| 2:13.107 | the excellence of |  |  |  |
| 2:13.108 | repentance |  |  |  |
| 2:13.109 | yesterday |  |  |  |
| 2:13.110 | was |  |  |  |
| 2:13.111 | this man |  |  |  |
| 2:13.112 | separated |  |  |  |
| 2:13.113 | from the LORD |  |  |  |
| 2:13.114 | the God of |  |  |  |
| 2:13.115 | Israel |  |  |  |
| 2:13.116 | etc. |  |  |  |
| 2:13.117 | and today |  |  |  |
| 2:13.118 | he |  |  |  |
| 2:13.119 | cleaving |  |  |  |
| 2:13.120 | to the Shekhinah |  |  |  |
| 2:13.121 | as it is said |  |  |  |
| 2:13.122 | and you |  |  |  |
| 2:13.123 | who cleave |  |  |  |
| 2:13.124 | to the LORD |  |  |  |
| 2:13.125 | your God |  |  |  |
| 2:13.126 | he cries out |  |  |  |
| 2:13.127 | and is answered |  |  | niphal participle |
| 2:13.128 | at once |  |  |  |
| 2:13.129 | as it is said |  |  |  |
| 2:13.130 | and it shall be |  |  |  |
| 2:13.131 | before |  |  |  |
| 2:13.132 | they call |  |  |  |
| 2:13.133 | and I |  |  |  |
| 2:13.134 | I will answer |  |  |  |
| 2:13.135 | and so |  |  |  |
| 2:13.136 | the sense of |  |  |  |
| 2:13.137 | his words |  |  |  |
| 2:13.138 | in Laws of |  |  |  |
| 2:13.139 | Prayer |  |  |  |
| 2:13.140 | that if |  |  |  |
| 2:13.141 | he was |  |  |  |
| 2:13.142 | in a place |  |  |  |
| 2:13.143 | where he is not |  |  |  |
| 2:13.144 | able |  |  |  |
| 2:13.145 | to determine |  |  | lekhavven, to align, fix |
| 2:13.146 | (object marker) |  |  |  |
| 2:13.147 | the directions |  |  | literally "the winds", the points of the compass |
| 2:13.148 | etc. |  |  | truncated; nothing supplied |
| 2:13.149 | and so |  |  |  |
| 2:13.150 | rendered |  |  | he'etiku: translated, or copied over into Arabic; either sense suits a rendering of the Rambam's rule |
| 2:13.151 | our forebears |  |  | qadmoneinu, the author's appeal to Yemenite tradition |
| 2:13.152 | and if |  |  | Arabic wa-in |
| 2:13.153 | he is |  |  | Arabic kana |
| 2:13.154 | in |  |  | Arabic fi |
| 2:13.155 | a place |  |  | Arabic; the spelling is probably buqca, "spot, place" |
| 2:13.156 | and not |  |  | Arabic wa-la |
| 2:13.157 | is able |  |  | Arabic yuhsin, "is able to, knows how to" |
| 2:13.158 | know |  |  | Arabic yacrif |
| 2:13.159 | the direction |  |  | Arabic al-jiha |
| 2:13.160 | let him direct |  |  | Arabic, spelled with an abbreviation mark; read as yuhayyi', "prepares, turns" |
| 2:13.161 | his heart |  |  | Arabic qalbahu |
| 2:13.162 | to |  |  | Arabic ila |
| 2:13.163 | his Creator |  |  | Arabic khaliqihi |
| 2:13.164 | and pray |  |  | Arabic wa-yusalli |

Phrases:
- 2:13.1 2:13.2 2:13.3 2:13.4 2:13.5 | As for your questions that you wrote, that one should ask | and-on questions-your that-you-wrote that-there-is to-ask
- 2:13.6 2:13.7 2:13.8 2:13.9 2:13.10 | about what the Rambam of blessed memory wrote: | on what that-wrote the-Rambam of-blessed-memory
- 2:13.11 2:13.12 2:13.13 2:13.14 2:13.15 | "Let him direct his heart toward the Shekhinah and pray," | he-shall-direct his-heart toward the-Shekhinah and-he-shall-pray
- 2:13.16 2:13.17 2:13.18 2:13.19 2:13.20 2:13.21 2:13.22 | and in Moreh Nevukhim he wrote that the Shekhinah is a created light, etc. | and-in-Moreh Nevukhim he-wrote that-the-Shekhinah light created etc.
- 2:13.23 2:13.24 2:13.25 2:13.26 2:13.27 2:13.28 2:13.29 2:13.30 2:13.31 | Behold, in your question you have revealed that you are wise in your own eyes and have not learned discipline. | behold in-your-question you-revealed that-you wise in-your-eyes and-not you-learned discipline
- 2:13.32 2:13.33 2:13.34 2:13.35 2:13.36 2:13.37 | Is it, then, the way of one who is asked to question his questioner? | and-is the-way-of the-one-asked to-ask [object-marker] his-asker
- 2:13.38 2:13.39 2:13.40 2:13.41 2:13.42 2:13.43 | In any case I will not withhold my quill from replying, | from-every place not I-will-withhold my-quill from-replying
- 2:13.44 2:13.45 2:13.46 2:13.47 | even though you have not acted properly! | although that-you-did not properly
- 2:13.48 2:13.49 2:13.50 2:13.51 2:13.52 2:13.53 2:13.54 2:13.55 2:13.56 2:13.57 | And I say that it is plain and known to every learned man well versed in the books of Israel | and-I-say that plain it and-known to-every learned versed in-books-of Israel
- 2:13.58 2:13.59 2:13.60 2:13.61 2:13.62 2:13.63 | that the name Shekhinah is said of 3 things: | that-name Shekhinah is-said on 3 things
- 2:13.64 2:13.65 2:13.66 | either of the Name, may He be blessed, | either to the-Name-may-He-be-blessed
- 2:13.67 2:13.68 2:13.69 2:13.70 2:13.71 2:13.72 | as our Master wrote in Laws of Repentance, chapter 7: | as wrote our-Master in-Laws-of Repentance chapter-7
- 2:13.73 2:13.74 2:13.75 2:13.76 2:13.77 2:13.78 | "Great is repentance, which brings a person near to the Shekhinah, as it is said: | great repentance that-brings-near the-person to-the-Shekhinah as-it-is-said
- 2:13.79 2:13.80 2:13.81 2:13.82 2:13.83 2:13.84 | 'Return, O Israel, to the LORD your God,' etc., | return Israel to the-LORD your-God etc.
- 2:13.85 2:13.86 2:13.87 2:13.88 2:13.89 2:13.90 | and it is said: 'Yet you did not return to Me, says the LORD,' | and-it-is-said and-not you-returned to-me declaration-of the-LORD
- 2:13.91 2:13.92 2:13.93 2:13.94 2:13.95 2:13.96 | and it is said: 'If you return, O Israel, to Me you shall return.' | and-it-is-said if you-return Israel to-me you-shall-return
- 2:13.97 2:13.98 2:13.99 2:13.100 2:13.101 2:13.102 2:13.103 | That is to say, if you turn in repentance, you will cleave to Me, etc." | that-is-to-say if you-turn in-repentance in-me you-will-cleave etc.
- 2:13.104 2:13.105 2:13.106 2:13.107 2:13.108 | Again there, on the excellence of repentance: | again there on excellence-of the-repentance
- 2:13.109 2:13.110 2:13.111 2:13.112 2:13.113 2:13.114 2:13.115 2:13.116 | "Yesterday this man was separated from the LORD, the God of Israel, etc., | yesterday was this-one separated from-the-LORD God-of Israel etc.
- 2:13.117 2:13.118 2:13.119 2:13.120 | and today he is cleaving to the Shekhinah, | and-today he cleaving to-the-Shekhinah
- 2:13.121 2:13.122 2:13.123 2:13.124 2:13.125 | as it is said: 'And you who cleave to the LORD your God,' | as-it-is-said and-you the-cleaving to-the-LORD your-God
- 2:13.126 2:13.127 2:13.128 | he cries out and is answered at once, | crying-out and-answered at-once
- 2:13.129 2:13.130 2:13.131 2:13.132 2:13.133 2:13.134 | as it is said: 'And it shall be that before they call, I will answer.'" | as-it-is-said and-it-shall-be before they-call and-I I-will-answer
- 2:13.135 2:13.136 2:13.137 2:13.138 2:13.139 | And so is the sense of his words in Laws of Prayer, | and-so sense-of his-words in-Laws-of Prayer
- 2:13.140 2:13.141 2:13.142 2:13.143 2:13.144 2:13.145 2:13.146 2:13.147 2:13.148 | that if he was in a place where he cannot determine the directions, etc. | that-if he-was in-a-place that-not able to-direct [object-marker] the-directions etc.
- 2:13.149 2:13.150 2:13.151 | And so our forebears rendered it: | and-so rendered our-forebears
- 2:13.152 2:13.153 2:13.154 2:13.155 2:13.156 2:13.157 2:13.158 2:13.159 2:13.160 2:13.161 2:13.162 2:13.163 2:13.164 | "And if he is in a place and cannot tell the direction, let him direct his heart to his Creator and pray." | and-if he-is in a-place and-not he-is-able he-knows the-direction he-directs his-heart to his-Creator and-he-prays

Grammar notes:
- Opening: the questioner has raised a contradiction between two statements of the Rambam; the reply first rebukes him (the rhetorical question and the exclamation) and only then answers.
- The answer is built on a list: the name Shekhinah is said of 3 things, introduced by the particle for the first one (either ... ), and continued in 2:14 (the second) and 2:15 (the third).
- Quotations inside quotations: the Rambam's Laws of Repentance is quoted with four verses, each introduced by the formula of the niphal, 'as it is said' or 'and it is said'; the author's own gloss on a verse (kelomar, 'that is to say') belongs to the Rambam's text.
- The participle forms of the Rambam's sentences ('cries out and is answered', 'is cleaving') describe a state, not a past event.
- The last sentence is Judeo-Arabic (Hebrew letters, Arabic words): a one-line proverb, introduced by 'so our forebears rendered it'. The first word, wa-in, 'and if', has the Arabic particle and the verb kana, 'is'.

Translator's note: The reply claims that the Rambam uses the word Shekhinah in three senses, which resolves the apparent contradiction between the Laws of Prayer ('toward the Shekhinah') and the Guide ('a created light'). The Arabic proverb is translated from the Arabic as it stands. The word the text writes with the letters bet, vav, tzadi, ayin is read as buqca, 'a place', and the word written with an abbreviation mark as yuhayyi'; both readings are the translator's, since the pack offers no lexicon for Judeo-Arabic. The proverb is presented as the Arabic rendering of the Rambam's rule in the Laws of Prayer, and the sense of rendered (he'etiku) is given in the token note. The Rambam passages are translated from the Hebrew the author quotes.

Commentary rulings:
- apparatus: Verse: Hos 14:2 ('Return, O Israel, to the LORD your God').
- apparatus: Verse: Amos 4:6 ('Yet you did not return to Me, says the LORD').
- apparatus: Verse: Jer 4:1; the author's text omits the words 'says the LORD' that stand after 'O Israel' in the verse, and the English follows his wording.
- apparatus: Verse: Deut 4:4 ('And you who cleave to the LORD your God', cut off at the comma).
- apparatus: Verse: Isa 65:24 ('And it shall be that before they call, I will answer').
- apparatus: Allusion: Prov 3:7 ('Do not be wise in your own eyes'); the English of 'wise in your own eyes' is embedded without quotation marks.
- apparatus: Judeo-Arabic proverb: wa-in kana fi buqca wa-la yuhsin yacrif al-jiha yuhayyi' qalbahu ila khaliqihi wa-yusalli; the spelling of two words (the word for place and the word before 'his heart') is not certain; both are translated as the evident word.

## Milchamot ha-Shem 2:14

Hebrew (Wikisource): והשני במקום שנזכר שם "גילוי שכינה", כגון גבי אברהם אבינו כשהלך לעקוד את יצחק שנאמר בו וירא את המקום מרחוק. ואמרו רז"ל שראה אור מרחוק בהר המוריה, וכמו שאמר הפיטן: וירא דמות כבוד והוד ויקר, ושאל עוד את נעריו, אור הראיתם צץ בראש הר המור וכו'. זהו אור נברא שכתב רבינו. וכן בסנה, וירא והנה הסנה בוער באש, ובהדיא נאמר שם וירא מלאך ה' אליו בלבת אש מתוך הסנה, וזהו אור נברא. וכן כל מקום שנזכר שם "גילוי שכינה" הוא אור נברא.

| id | word | punctuation after |
|---|---|---|
| 2:14.1 | והשני |  |
| 2:14.2 | במקום |  |
| 2:14.3 | שנזכר |  |
| 2:14.4 | שם |  |
| 2:14.5 | גילוי |  |
| 2:14.6 | שכינה | ", |
| 2:14.7 | כגון |  |
| 2:14.8 | גבי |  |
| 2:14.9 | אברהם |  |
| 2:14.10 | אבינו |  |
| 2:14.11 | כשהלך |  |
| 2:14.12 | לעקוד |  |
| 2:14.13 | את |  |
| 2:14.14 | יצחק |  |
| 2:14.15 | שנאמר |  |
| 2:14.16 | בו |  |
| 2:14.17 | וירא |  |
| 2:14.18 | את |  |
| 2:14.19 | המקום |  |
| 2:14.20 | מרחוק | . |
| 2:14.21 | ואמרו |  |
| 2:14.22 | רז"ל |  |
| 2:14.23 | שראה |  |
| 2:14.24 | אור |  |
| 2:14.25 | מרחוק |  |
| 2:14.26 | בהר |  |
| 2:14.27 | המוריה | , |
| 2:14.28 | וכמו |  |
| 2:14.29 | שאמר |  |
| 2:14.30 | הפיטן | : |
| 2:14.31 | וירא |  |
| 2:14.32 | דמות |  |
| 2:14.33 | כבוד |  |
| 2:14.34 | והוד |  |
| 2:14.35 | ויקר | , |
| 2:14.36 | ושאל |  |
| 2:14.37 | עוד |  |
| 2:14.38 | את |  |
| 2:14.39 | נעריו | , |
| 2:14.40 | אור |  |
| 2:14.41 | הראיתם |  |
| 2:14.42 | צץ |  |
| 2:14.43 | בראש |  |
| 2:14.44 | הר |  |
| 2:14.45 | המור |  |
| 2:14.46 | וכו | '. |
| 2:14.47 | זהו |  |
| 2:14.48 | אור |  |
| 2:14.49 | נברא |  |
| 2:14.50 | שכתב |  |
| 2:14.51 | רבינו | . |
| 2:14.52 | וכן |  |
| 2:14.53 | בסנה | , |
| 2:14.54 | וירא |  |
| 2:14.55 | והנה |  |
| 2:14.56 | הסנה |  |
| 2:14.57 | בוער |  |
| 2:14.58 | באש | , |
| 2:14.59 | ובהדיא |  |
| 2:14.60 | נאמר |  |
| 2:14.61 | שם |  |
| 2:14.62 | וירא |  |
| 2:14.63 | מלאך |  |
| 2:14.64 | ה | ' |
| 2:14.65 | אליו |  |
| 2:14.66 | בלבת |  |
| 2:14.67 | אש |  |
| 2:14.68 | מתוך |  |
| 2:14.69 | הסנה | , |
| 2:14.70 | וזהו |  |
| 2:14.71 | אור |  |
| 2:14.72 | נברא | . |
| 2:14.73 | וכן |  |
| 2:14.74 | כל |  |
| 2:14.75 | מקום |  |
| 2:14.76 | שנזכר |  |
| 2:14.77 | שם |  |
| 2:14.78 | גילוי |  |
| 2:14.79 | שכינה | " |
| 2:14.80 | הוא |  |
| 2:14.81 | אור |  |
| 2:14.82 | נברא | . |


### Translation under review (2:14)

English: And the second is in a place where the term "the revelation of the Shekhinah" is mentioned, for example, regarding Abraham our father when he went to bind Isaac, of whom it is said: "And he saw the place from afar." Our Rabbis of blessed memory said that he saw a light from afar on Mount Moriah, and as the liturgical poet said: "And he saw a likeness of glory, splendor and honor, and further asked his young men, 'A light: did you see it shining on the top of Mount ha-Mor?' etc." This is the created light that our Master wrote of. And so at the bush: and he saw, and behold, the bush was burning with fire, and explicitly it is said there: "And the angel of the LORD appeared to him in a flame of fire out of the bush," and this is the created light. And so in every place where the term "the revelation of the Shekhinah" is mentioned, it is a created light.

Literal: and-the-second in-a-place that-is-mentioned there/name revelation-of Shekhinah for-example regarding Abraham our-father when-he-went to-bind [object-marker] Isaac that-is-said of-him and-he-saw [object-marker] the-place from-afar and-said our-Rabbis-of-blessed-memory that-he-saw light from-afar on-mount the-Moriah and-as said the-poet and-he-saw likeness-of glory and-splendor and-honor and-asked further [object-marker] his-young-men light did-you-see shining on-top-of mount the-Mor etc. this-is light created that-wrote our-Master and-so at-the-bush and-he-saw and-behold the-bush burning with-fire and-explicitly it-is-said there and-appeared angel-of the-LORD to-him in-flame-of fire from-within the-bush and-this-is light created and-so every place that-is-mentioned there/name revelation-of Shekhinah it light created

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:14.1 | and the second |  |  |  |
| 2:14.2 | in a place |  |  |  |
| 2:14.3 | where is mentioned |  |  |  |
| 2:14.4 | the term |  |  | shem: "there" or "name"; see the note |
| 2:14.5 | the revelation of |  |  | gilui |
| 2:14.6 | the Shekhinah |  |  |  |
| 2:14.7 | for example |  |  |  |
| 2:14.8 | regarding |  |  |  |
| 2:14.9 | Abraham |  |  |  |
| 2:14.10 | our father |  |  |  |
| 2:14.11 | when he went |  |  |  |
| 2:14.12 | to bind |  |  | la'akod, the verb of the Akedah |
| 2:14.13 | (object marker) |  |  |  |
| 2:14.14 | Isaac |  |  |  |
| 2:14.15 | of whom it is said |  |  |  |
| 2:14.16 | of him |  |  |  |
| 2:14.17 | and he saw |  |  |  |
| 2:14.18 | (object marker) |  |  |  |
| 2:14.19 | the place |  |  |  |
| 2:14.20 | from afar |  |  |  |
| 2:14.21 | and said |  |  |  |
| 2:14.22 | our Rabbis of blessed memory |  |  | abbreviation razal, expanded |
| 2:14.23 | that he saw |  |  |  |
| 2:14.24 | a light |  |  |  |
| 2:14.25 | from afar |  |  |  |
| 2:14.26 | on Mount |  |  |  |
| 2:14.27 | Moriah |  |  |  |
| 2:14.28 | and as |  |  |  |
| 2:14.29 | said |  |  |  |
| 2:14.30 | the liturgical poet |  |  | ha-paytan; the poem is not named |
| 2:14.31 | and he saw |  |  |  |
| 2:14.32 | a likeness of |  |  |  |
| 2:14.33 | glory |  |  |  |
| 2:14.34 | and splendor |  |  |  |
| 2:14.35 | and honor |  |  |  |
| 2:14.36 | and asked |  |  |  |
| 2:14.37 | further |  |  |  |
| 2:14.38 | (object marker) |  |  |  |
| 2:14.39 | his young men |  |  | ne'arav |
| 2:14.40 | a light |  |  |  |
| 2:14.41 | did you see |  |  | the ha- is the interrogative |
| 2:14.42 | shining |  |  | tzatz: flashes, gleams |
| 2:14.43 | on the top of |  |  |  |
| 2:14.44 | Mount |  |  |  |
| 2:14.45 | ha-Mor |  |  | the poet's form for Moriah; in Song 4:6 the same words mean "the mountain of myrrh" |
| 2:14.46 | etc. |  |  |  |
| 2:14.47 | this is |  |  |  |
| 2:14.48 | light |  |  |  |
| 2:14.49 | created |  |  |  |
| 2:14.50 | that wrote |  |  |  |
| 2:14.51 | our Master |  |  | the Rambam |
| 2:14.52 | and so |  |  |  |
| 2:14.53 | at the bush |  |  |  |
| 2:14.54 | and he saw |  |  |  |
| 2:14.55 | and behold |  |  |  |
| 2:14.56 | the bush |  |  |  |
| 2:14.57 | burning |  |  | participle (bo'er); the verse has the perfect form |
| 2:14.58 | with fire |  |  |  |
| 2:14.59 | and explicitly |  |  | Aramaic be-hadya, "openly, expressly" |
| 2:14.60 | it is said |  |  |  |
| 2:14.61 | there |  |  |  |
| 2:14.62 | and appeared |  |  | the same verb as "he saw" in the previous verse; here niphal "appeared" |
| 2:14.63 | angel of |  |  |  |
| 2:14.64 | the LORD |  |  |  |
| 2:14.65 | to him |  |  |  |
| 2:14.66 | in a flame of |  |  |  |
| 2:14.67 | fire |  |  |  |
| 2:14.68 | out of |  |  |  |
| 2:14.69 | the bush |  |  |  |
| 2:14.70 | and this is |  |  |  |
| 2:14.71 | light |  |  |  |
| 2:14.72 | created |  |  |  |
| 2:14.73 | and so |  |  |  |
| 2:14.74 | every |  |  |  |
| 2:14.75 | place |  |  |  |
| 2:14.76 | that is mentioned |  |  |  |
| 2:14.77 | the term |  |  |  |
| 2:14.78 | the revelation of |  |  |  |
| 2:14.79 | the Shekhinah |  |  |  |
| 2:14.80 | it is |  |  |  |
| 2:14.81 | light |  |  |  |
| 2:14.82 | created |  |  |  |

Phrases:
- 2:14.1 2:14.2 2:14.3 2:14.4 2:14.5 2:14.6 | And the second is in a place where the term "the revelation of the Shekhinah" is mentioned, | and-the-second in-a-place that-is-mentioned there/name revelation-of Shekhinah
- 2:14.7 2:14.8 2:14.9 2:14.10 2:14.11 2:14.12 2:14.13 2:14.14 | for example, regarding Abraham our father when he went to bind Isaac, | for-example regarding Abraham our-father when-he-went to-bind [object-marker] Isaac
- 2:14.15 2:14.16 2:14.17 2:14.18 2:14.19 2:14.20 | of whom it is said: "And he saw the place from afar." | that-is-said of-him and-he-saw [object-marker] the-place from-afar
- 2:14.21 2:14.22 2:14.23 2:14.24 2:14.25 2:14.26 2:14.27 | Our Rabbis of blessed memory said that he saw a light from afar on Mount Moriah, | and-said our-Rabbis-of-blessed-memory that-he-saw light from-afar on-mount the-Moriah
- 2:14.28 2:14.29 2:14.30 2:14.31 2:14.32 2:14.33 2:14.34 2:14.35 | and as the liturgical poet said: "And he saw a likeness of glory, splendor and honor, | and-as said the-poet and-he-saw likeness-of glory and-splendor and-honor
- 2:14.36 2:14.37 2:14.38 2:14.39 | and further asked his young men, | and-asked further [object-marker] his-young-men
- 2:14.40 2:14.41 2:14.42 2:14.43 2:14.44 2:14.45 2:14.46 | 'A light: did you see it shining on the top of Mount ha-Mor?' etc." | light did-you-see shining on-top-of mount the-Mor etc.
- 2:14.47 2:14.48 2:14.49 2:14.50 2:14.51 | This is the created light that our Master wrote of. | this-is light created that-wrote our-Master
- 2:14.52 2:14.53 2:14.54 2:14.55 2:14.56 2:14.57 2:14.58 | And so at the bush: and he saw, and behold, the bush was burning with fire, | and-so at-the-bush and-he-saw and-behold the-bush burning with-fire
- 2:14.59 2:14.60 2:14.61 | and explicitly it is said there: | and-explicitly it-is-said there
- 2:14.62 2:14.63 2:14.64 2:14.65 2:14.66 2:14.67 2:14.68 2:14.69 | "And the angel of the LORD appeared to him in a flame of fire out of the bush," | and-appeared angel-of the-LORD to-him in-flame-of fire from-within the-bush
- 2:14.70 2:14.71 2:14.72 | and this is the created light. | and-this-is light created
- 2:14.73 2:14.74 2:14.75 2:14.76 2:14.77 2:14.78 2:14.79 | And so in every place where the term "the revelation of the Shekhinah" is mentioned, | and-so every place that-is-mentioned there/name revelation-of Shekhinah
- 2:14.80 2:14.81 2:14.82 | it is a created light. | it light created

Grammar notes:
- The second sense is identified by a phrase that the text itself sets in quotation marks: the revelation of the Shekhinah. The word shem after the verb 'is mentioned' may be 'there' or 'term, name'; the English takes it as 'term' and keeps the quotation marks of the text.
- Three witnesses for one light: the Rabbis on Mount Moriah, a liturgical poet, and the burning bush. Each is brought by a different formula: 'said', 'as the poet said', and a bare quotation of the verse at the bush.
- The same verb is used in Hebrew for 'saw' and 'appeared' (vayar / vayera): Abraham 'saw' the place; the angel 'appeared' to Moses in the fire.
- The poet's question begins with a ha- that makes it interrogative ('did you see'), and his phrase 'Mount ha-Mor' is a form of Moriah.
- The conclusion is a rule: wherever the revelation of the Shekhinah is mentioned, it is a created light.

Translator's note: The word shem in 'where the term ... is mentioned' is ambiguous (the first occurrence in this paragraph and the last are the same phrase). The Hebrew of the bush verse in the text has the participle 'burning'; the verse in Exod 3:2 has the perfect form. The English follows the author's word. The poet is not named and the quoted lines are not identified.

Commentary rulings:
- apparatus: Verse: Gen 22:4 ('And he saw the place from afar').
- apparatus: Verse: Exod 3:2 ('And he saw, and behold, the bush was burning with fire'; the verse itself has the perfect form of 'burn'), and the second half of the same verse ('And the angel of the LORD appeared to him in a flame of fire out of the bush').
- apparatus: The poetic lines are quoted from an unnamed liturgical poem; no source is given and none is added.

## Milchamot ha-Shem 2:15

Hebrew (Wikisource): והשלישי, השגחת השי"ת עלינו נקרא ג"כ שכינה, כמו שאמרו רז"ל: גלו לבבל שכינה עמהם. כלומר, השגחתו ית' דבקה בנו כמו שהבטיחנו בתורה: "ואף גם זאת בהיותם בארץ אויביהם לא מאסתים ולא געלתים לכלותם להפר בריתי אתם כי אני ה' אלהיהם". וכן כל כיוצא בזה. כל מקום יובן לפי עניינו באחד משלשה פנים הללו הנז'.

| id | word | punctuation after |
|---|---|---|
| 2:15.1 | והשלישי | , |
| 2:15.2 | השגחת |  |
| 2:15.3 | השי"ת |  |
| 2:15.4 | עלינו |  |
| 2:15.5 | נקרא |  |
| 2:15.6 | ג"כ |  |
| 2:15.7 | שכינה | , |
| 2:15.8 | כמו |  |
| 2:15.9 | שאמרו |  |
| 2:15.10 | רז"ל | : |
| 2:15.11 | גלו |  |
| 2:15.12 | לבבל |  |
| 2:15.13 | שכינה |  |
| 2:15.14 | עמהם | . |
| 2:15.15 | כלומר | , |
| 2:15.16 | השגחתו |  |
| 2:15.17 | ית | ' |
| 2:15.18 | דבקה |  |
| 2:15.19 | בנו |  |
| 2:15.20 | כמו |  |
| 2:15.21 | שהבטיחנו |  |
| 2:15.22 | בתורה | : |
| 2:15.23 | ואף |  |
| 2:15.24 | גם |  |
| 2:15.25 | זאת |  |
| 2:15.26 | בהיותם |  |
| 2:15.27 | בארץ |  |
| 2:15.28 | אויביהם |  |
| 2:15.29 | לא |  |
| 2:15.30 | מאסתים |  |
| 2:15.31 | ולא |  |
| 2:15.32 | געלתים |  |
| 2:15.33 | לכלותם |  |
| 2:15.34 | להפר |  |
| 2:15.35 | בריתי |  |
| 2:15.36 | אתם |  |
| 2:15.37 | כי |  |
| 2:15.38 | אני |  |
| 2:15.39 | ה | ' |
| 2:15.40 | אלהיהם | ". |
| 2:15.41 | וכן |  |
| 2:15.42 | כל |  |
| 2:15.43 | כיוצא |  |
| 2:15.44 | בזה | . |
| 2:15.45 | כל |  |
| 2:15.46 | מקום |  |
| 2:15.47 | יובן |  |
| 2:15.48 | לפי |  |
| 2:15.49 | עניינו |  |
| 2:15.50 | באחד |  |
| 2:15.51 | משלשה |  |
| 2:15.52 | פנים |  |
| 2:15.53 | הללו |  |
| 2:15.54 | הנז | '. |


### Translation under review (2:15)

English: And the third: the providence of the Name, may He be blessed, over us is also called Shekhinah, as our Rabbis of blessed memory said: "They went into exile to Babylon; the Shekhinah was with them." That is to say, His providence, may He be blessed, clung to us, as He promised us in the Torah: "And yet even so, when they are in the land of their enemies, I will not reject them and I will not abhor them so as to destroy them, to break My covenant with them, for I am the LORD their God." And so it is with everything of the kind. Every place is to be understood according to its subject in one of these three mentioned aspects.

Literal: and-the-third the-providence-of the-Name-may-He-be-blessed over-us is-called also Shekhinah as said our-Rabbis-of-blessed-memory they-went-into-exile to-Babylon Shekhinah with-them that-is-to-say his-providence may-He-be-blessed clung to-us as He-promised-us in-the-Torah and-even also this when-they-are in-land-of their-enemies not I-rejected-them and-not I-abhorred-them to-destroy-them to-break my-covenant with-them for I the-LORD their-God and-so every like-this every place is-understood according-to its-subject in-one from-three aspects these the-mentioned

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:15.1 | and the third |  |  |  |
| 2:15.2 | the providence of |  |  | hashgachah |
| 2:15.3 | the Name, may He be blessed |  |  | abbreviation, expanded |
| 2:15.4 | over us |  |  |  |
| 2:15.5 | is called |  |  |  |
| 2:15.6 | also |  |  | abbreviation gam ken, expanded |
| 2:15.7 | Shekhinah |  |  |  |
| 2:15.8 | as |  |  |  |
| 2:15.9 | said |  |  |  |
| 2:15.10 | our Rabbis of blessed memory |  |  |  |
| 2:15.11 | they went into exile |  |  | the sentence has no verb for the Shekhinah |
| 2:15.12 | to Babylon |  |  |  |
| 2:15.13 | the Shekhinah |  |  |  |
| 2:15.14 | with them |  |  |  |
| 2:15.15 | that is to say |  |  |  |
| 2:15.16 | His providence |  |  |  |
| 2:15.17 | may He be blessed |  |  | abbreviation yit., expanded |
| 2:15.18 | clung |  |  | davkah, feminine to agree with providence |
| 2:15.19 | to us |  |  |  |
| 2:15.20 | as |  |  |  |
| 2:15.21 | He promised us |  |  |  |
| 2:15.22 | in the Torah |  |  |  |
| 2:15.23 | and yet |  |  |  |
| 2:15.24 | even |  |  |  |
| 2:15.25 | so |  |  |  |
| 2:15.26 | when they are |  |  |  |
| 2:15.27 | in the land of |  |  |  |
| 2:15.28 | their enemies |  |  |  |
| 2:15.29 | not |  |  |  |
| 2:15.30 | I will reject them |  |  | perfect tense, as in the verse, with the force of a promise |
| 2:15.31 | and not |  |  |  |
| 2:15.32 | I will abhor them |  |  |  |
| 2:15.33 | so as to destroy them |  |  |  |
| 2:15.34 | to break |  |  |  |
| 2:15.35 | My covenant |  |  |  |
| 2:15.36 | with them |  |  |  |
| 2:15.37 | for |  |  |  |
| 2:15.38 | I |  |  |  |
| 2:15.39 | the LORD |  |  |  |
| 2:15.40 | their God |  |  |  |
| 2:15.41 | and so |  |  |  |
| 2:15.42 | every |  |  |  |
| 2:15.43 | of the kind |  |  | ke-yotze ba-zeh, "like this" |
| 2:15.44 | of this |  |  |  |
| 2:15.45 | every |  |  |  |
| 2:15.46 | place |  |  | here a place in the text, a passage |
| 2:15.47 | is understood |  |  |  |
| 2:15.48 | according to |  |  |  |
| 2:15.49 | its subject |  |  |  |
| 2:15.50 | in one |  |  |  |
| 2:15.51 | of three |  |  |  |
| 2:15.52 | aspects |  |  | panim, literally "faces" |
| 2:15.53 | these |  |  |  |
| 2:15.54 | mentioned |  |  | abbreviation ha-nizkar, expanded |

Phrases:
- 2:15.1 | And the third: | and-the-third
- 2:15.2 2:15.3 2:15.4 2:15.5 2:15.6 2:15.7 | the providence of the Name, may He be blessed, over us is also called Shekhinah, | the-providence-of the-Name-may-He-be-blessed over-us is-called also Shekhinah
- 2:15.8 2:15.9 2:15.10 | as our Rabbis of blessed memory said: | as said our-Rabbis-of-blessed-memory
- 2:15.11 2:15.12 2:15.13 2:15.14 | "They went into exile to Babylon; the Shekhinah was with them." | they-went-into-exile to-Babylon Shekhinah with-them
- 2:15.15 2:15.16 2:15.17 2:15.18 2:15.19 | That is to say, His providence, may He be blessed, clung to us, | that-is-to-say his-providence may-He-be-blessed clung to-us
- 2:15.20 2:15.21 2:15.22 | as He promised us in the Torah: | as He-promised-us in-the-Torah
- 2:15.23 2:15.24 2:15.25 2:15.26 2:15.27 2:15.28 | "And yet even so, when they are in the land of their enemies, | and-even also this when-they-are in-land-of their-enemies
- 2:15.29 2:15.30 2:15.31 2:15.32 | I will not reject them and I will not abhor them | not I-rejected-them and-not I-abhorred-them
- 2:15.33 2:15.34 2:15.35 2:15.36 | so as to destroy them, to break My covenant with them, | to-destroy-them to-break my-covenant with-them
- 2:15.37 2:15.38 2:15.39 2:15.40 | for I am the LORD their God." | for I the-LORD their-God
- 2:15.41 2:15.42 2:15.43 2:15.44 | And so it is with everything of the kind. | and-so every like-this
- 2:15.45 2:15.46 2:15.47 | Every place is to be understood | every place is-understood
- 2:15.48 2:15.49 | according to its subject | according-to its-subject
- 2:15.50 2:15.51 2:15.52 2:15.53 2:15.54 | in one of these three mentioned aspects. | in-one from-three aspects these the-mentioned

Grammar notes:
- The third sense is God's providence over His people. The proof is a rabbinic saying with no verb for the second clause (the Shekhinah with them) and a verse from the Torah.
- The verb of the promise in the verse is a perfect, 'I have not rejected them', which in this context states a promise about the future, as the English renders it.
- Dabekah in 'clung to us' is feminine to agree with 'providence'; the verb 'clung' reuses the stem of the word the Rambam used for repentance (2:13).
- The paragraph closes with a rule: each occurrence of the word is to be understood according to its subject as one of the three senses.

Translator's note: The rabbinic saying is quoted without a source; none is added. The verse is Lev 26:44 as the author gives it. His spelling of two words differs from the verse; the meaning is not affected.

Commentary rulings:
- apparatus: Verse: Lev 26:44 ('And yet even so, when they are in the land of their enemies, I will not reject them ... for I am the LORD their God'); the author's text spells 'their enemies' and 'to destroy them' with an added vav.
- apparatus: Rabbinic saying: 'They went into exile to Babylon; the Shekhinah was with them'; no source is given.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

