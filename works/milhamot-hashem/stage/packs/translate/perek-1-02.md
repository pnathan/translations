# Translation pack: milhamot-hashem perek-1-02 (Milchamot ha-Shem 1:3–1:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-02`.

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

- סתו"מ: always "Torah scrolls, tefillin and mezuzot", with a note.
- Turkish and Arabic personal names and titles: transliterate by the project
  scheme as the text spells them; keep the titles (Effendi, Bey, Pasha);
  identify no one.
- ה' as an honorific before a personal name: "Mr." (1:2). Elsewhere ה' is
  the LORD.
- בית רבן in "תינוקות של בית רבן": "schoolchildren" (literally "the house of
  the master/teacher"), with a note.
- משכילים / השכלה: "the educated" / "enlightenment" in the sense of the
  Haskalah only where the paragraph speaks of that movement; otherwise
  "those of understanding".
- ובפתרון הכתובים: "the verses" (of the Torah) when the example is from the
  Torah; "the Writings" only for כתובים as the third part of the Bible.
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
- A sentence opening שכל / שהוא after a full stop: keep the full stop and
  supply a main clause by the author's syntax, with a token note on the ש.
- דברי תימה: "astonishing words" (words that astonish).
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
- בא"י: "in the Land of Israel"; תוה"ק: "the holy Torah".
- Acronym names of persons with a place (יש"ר מקנדי"א): expand by the usual form
  ("Yish'ar of Candia"), noting that the expansion is conjectural if the
  surname does not fix it.
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


## Milchamot ha-Shem 1:3

Hebrew (Wikisource): ומעת אשר פתחה הממשלה את בית הספר היו הילדים לומדים מקרא ותרגומיו אונקלוס ורס״ג, משנה והרמב״ם, ודקדוק ופי' המלות, הכל נעשה על נכון לימוד התורה נביאים וכתובים בפי' ברור ספר דקדוק לשון הקדש! ולימוד השפה התורקית על ידי המלמד התורקי ישר הלבב חאפץ אפנדי. ואחריו זיא ביי בחריצות נפלאה! והרגלנום לדבר זה עם זה בשפה העברית והתורקית כדי שיהיו מורגלים בפיהם היטב, ולתקופת השנה באו הבוחנים מאת הממשלה יר״ה והסדירו הילדים את לימודם בשפה התורקית והראו להם כתיבתם ויראו הבוחנים את כשרונותם מתנוססת בלימודים, ויהללו אותם כי משכילים הם! גם הבוחנים העברים באו ביום של אחריו והסדירו הילדים לפניהם את לימודיהם בעברית ובערבית וימצאו חן בעיניהם בכשרון לימודיהם ויהללו אותם! ולשנה האחרת כן עשו הילדים ללמוד בחריצות וזריזות, ובצאת השנה באו הבוחנים התורקים ובחנו אותם ביושר לבב, ושיבחו אותם שהם עושים חיל בלימודים, וזירזו אותם בדברים ובירכו אותם! אך הבוחנים העברים הבאים אחריהם כבר התחילה להתנוצץ בקרבם שנאת ההשכלה ולימוד שפות. והתחילו לקנתרם ולשאול אותם במה שעדיין לא למדו, וכל שאינו יכול להשיב דבר בכל מקום ששואלין אותו כדין חכם גמור לא נחשב בעיניהם למאומה. כדי להפחית כבודם ולהעציבם ולהרפות ידיהם מן הלימודים ויבקשו חסרונות ומומים לכל אחד ואחד, ומרוע לב לא פנו אל הלמודים והידיעות שרכשו הילדים. רק לקפחם במה שעדיין לא למדו, ולשאול אותם בסדר נזיקים ועדיין עומדים בסדר מועד, ובהראות להם את יופי כתיבתם כתב אשורי מפואר ומסודר בקשו להם חסרון בדיבוק כחוט השערה בקצת אותיות כאלו הם כותבים סתו״מ. ולא הביטו ליופי הכתב וסדורו. ובפתרון הכתובים בקשו מהם לפתור להם מקראות חמורים כגון אסרי לגפן עירה, וכיוצא ויצאו מלפניהם בפחי נפש טובים דוים ועצבים ורפויי ידים.

| id | word | punctuation after |
|---|---|---|
| 1:3.1 | ומעת |  |
| 1:3.2 | אשר |  |
| 1:3.3 | פתחה |  |
| 1:3.4 | הממשלה |  |
| 1:3.5 | את |  |
| 1:3.6 | בית |  |
| 1:3.7 | הספר |  |
| 1:3.8 | היו |  |
| 1:3.9 | הילדים |  |
| 1:3.10 | לומדים |  |
| 1:3.11 | מקרא |  |
| 1:3.12 | ותרגומיו |  |
| 1:3.13 | אונקלוס |  |
| 1:3.14 | ורס״ג | , |
| 1:3.15 | משנה |  |
| 1:3.16 | והרמב״ם | , |
| 1:3.17 | ודקדוק |  |
| 1:3.18 | ופי | ' |
| 1:3.19 | המלות | , |
| 1:3.20 | הכל |  |
| 1:3.21 | נעשה |  |
| 1:3.22 | על |  |
| 1:3.23 | נכון |  |
| 1:3.24 | לימוד |  |
| 1:3.25 | התורה |  |
| 1:3.26 | נביאים |  |
| 1:3.27 | וכתובים |  |
| 1:3.28 | בפי | ' |
| 1:3.29 | ברור |  |
| 1:3.30 | ספר |  |
| 1:3.31 | דקדוק |  |
| 1:3.32 | לשון |  |
| 1:3.33 | הקדש | ! |
| 1:3.34 | ולימוד |  |
| 1:3.35 | השפה |  |
| 1:3.36 | התורקית |  |
| 1:3.37 | על |  |
| 1:3.38 | ידי |  |
| 1:3.39 | המלמד |  |
| 1:3.40 | התורקי |  |
| 1:3.41 | ישר |  |
| 1:3.42 | הלבב |  |
| 1:3.43 | חאפץ |  |
| 1:3.44 | אפנדי | . |
| 1:3.45 | ואחריו |  |
| 1:3.46 | זיא |  |
| 1:3.47 | ביי |  |
| 1:3.48 | בחריצות |  |
| 1:3.49 | נפלאה | ! |
| 1:3.50 | והרגלנום |  |
| 1:3.51 | לדבר |  |
| 1:3.52 | זה |  |
| 1:3.53 | עם |  |
| 1:3.54 | זה |  |
| 1:3.55 | בשפה |  |
| 1:3.56 | העברית |  |
| 1:3.57 | והתורקית |  |
| 1:3.58 | כדי |  |
| 1:3.59 | שיהיו |  |
| 1:3.60 | מורגלים |  |
| 1:3.61 | בפיהם |  |
| 1:3.62 | היטב | , |
| 1:3.63 | ולתקופת |  |
| 1:3.64 | השנה |  |
| 1:3.65 | באו |  |
| 1:3.66 | הבוחנים |  |
| 1:3.67 | מאת |  |
| 1:3.68 | הממשלה |  |
| 1:3.69 | יר״ה |  |
| 1:3.70 | והסדירו |  |
| 1:3.71 | הילדים |  |
| 1:3.72 | את |  |
| 1:3.73 | לימודם |  |
| 1:3.74 | בשפה |  |
| 1:3.75 | התורקית |  |
| 1:3.76 | והראו |  |
| 1:3.77 | להם |  |
| 1:3.78 | כתיבתם |  |
| 1:3.79 | ויראו |  |
| 1:3.80 | הבוחנים |  |
| 1:3.81 | את |  |
| 1:3.82 | כשרונותם |  |
| 1:3.83 | מתנוססת |  |
| 1:3.84 | בלימודים | , |
| 1:3.85 | ויהללו |  |
| 1:3.86 | אותם |  |
| 1:3.87 | כי |  |
| 1:3.88 | משכילים |  |
| 1:3.89 | הם | ! |
| 1:3.90 | גם |  |
| 1:3.91 | הבוחנים |  |
| 1:3.92 | העברים |  |
| 1:3.93 | באו |  |
| 1:3.94 | ביום |  |
| 1:3.95 | של |  |
| 1:3.96 | אחריו |  |
| 1:3.97 | והסדירו |  |
| 1:3.98 | הילדים |  |
| 1:3.99 | לפניהם |  |
| 1:3.100 | את |  |
| 1:3.101 | לימודיהם |  |
| 1:3.102 | בעברית |  |
| 1:3.103 | ובערבית |  |
| 1:3.104 | וימצאו |  |
| 1:3.105 | חן |  |
| 1:3.106 | בעיניהם |  |
| 1:3.107 | בכשרון |  |
| 1:3.108 | לימודיהם |  |
| 1:3.109 | ויהללו |  |
| 1:3.110 | אותם | ! |
| 1:3.111 | ולשנה |  |
| 1:3.112 | האחרת |  |
| 1:3.113 | כן |  |
| 1:3.114 | עשו |  |
| 1:3.115 | הילדים |  |
| 1:3.116 | ללמוד |  |
| 1:3.117 | בחריצות |  |
| 1:3.118 | וזריזות | , |
| 1:3.119 | ובצאת |  |
| 1:3.120 | השנה |  |
| 1:3.121 | באו |  |
| 1:3.122 | הבוחנים |  |
| 1:3.123 | התורקים |  |
| 1:3.124 | ובחנו |  |
| 1:3.125 | אותם |  |
| 1:3.126 | ביושר |  |
| 1:3.127 | לבב | , |
| 1:3.128 | ושיבחו |  |
| 1:3.129 | אותם |  |
| 1:3.130 | שהם |  |
| 1:3.131 | עושים |  |
| 1:3.132 | חיל |  |
| 1:3.133 | בלימודים | , |
| 1:3.134 | וזירזו |  |
| 1:3.135 | אותם |  |
| 1:3.136 | בדברים |  |
| 1:3.137 | ובירכו |  |
| 1:3.138 | אותם | ! |
| 1:3.139 | אך |  |
| 1:3.140 | הבוחנים |  |
| 1:3.141 | העברים |  |
| 1:3.142 | הבאים |  |
| 1:3.143 | אחריהם |  |
| 1:3.144 | כבר |  |
| 1:3.145 | התחילה |  |
| 1:3.146 | להתנוצץ |  |
| 1:3.147 | בקרבם |  |
| 1:3.148 | שנאת |  |
| 1:3.149 | ההשכלה |  |
| 1:3.150 | ולימוד |  |
| 1:3.151 | שפות | . |
| 1:3.152 | והתחילו |  |
| 1:3.153 | לקנתרם |  |
| 1:3.154 | ולשאול |  |
| 1:3.155 | אותם |  |
| 1:3.156 | במה |  |
| 1:3.157 | שעדיין |  |
| 1:3.158 | לא |  |
| 1:3.159 | למדו | , |
| 1:3.160 | וכל |  |
| 1:3.161 | שאינו |  |
| 1:3.162 | יכול |  |
| 1:3.163 | להשיב |  |
| 1:3.164 | דבר |  |
| 1:3.165 | בכל |  |
| 1:3.166 | מקום |  |
| 1:3.167 | ששואלין |  |
| 1:3.168 | אותו |  |
| 1:3.169 | כדין |  |
| 1:3.170 | חכם |  |
| 1:3.171 | גמור |  |
| 1:3.172 | לא |  |
| 1:3.173 | נחשב |  |
| 1:3.174 | בעיניהם |  |
| 1:3.175 | למאומה | . |
| 1:3.176 | כדי |  |
| 1:3.177 | להפחית |  |
| 1:3.178 | כבודם |  |
| 1:3.179 | ולהעציבם |  |
| 1:3.180 | ולהרפות |  |
| 1:3.181 | ידיהם |  |
| 1:3.182 | מן |  |
| 1:3.183 | הלימודים |  |
| 1:3.184 | ויבקשו |  |
| 1:3.185 | חסרונות |  |
| 1:3.186 | ומומים |  |
| 1:3.187 | לכל |  |
| 1:3.188 | אחד |  |
| 1:3.189 | ואחד | , |
| 1:3.190 | ומרוע |  |
| 1:3.191 | לב |  |
| 1:3.192 | לא |  |
| 1:3.193 | פנו |  |
| 1:3.194 | אל |  |
| 1:3.195 | הלמודים |  |
| 1:3.196 | והידיעות |  |
| 1:3.197 | שרכשו |  |
| 1:3.198 | הילדים | . |
| 1:3.199 | רק |  |
| 1:3.200 | לקפחם |  |
| 1:3.201 | במה |  |
| 1:3.202 | שעדיין |  |
| 1:3.203 | לא |  |
| 1:3.204 | למדו | , |
| 1:3.205 | ולשאול |  |
| 1:3.206 | אותם |  |
| 1:3.207 | בסדר |  |
| 1:3.208 | נזיקים |  |
| 1:3.209 | ועדיין |  |
| 1:3.210 | עומדים |  |
| 1:3.211 | בסדר |  |
| 1:3.212 | מועד | , |
| 1:3.213 | ובהראות |  |
| 1:3.214 | להם |  |
| 1:3.215 | את |  |
| 1:3.216 | יופי |  |
| 1:3.217 | כתיבתם |  |
| 1:3.218 | כתב |  |
| 1:3.219 | אשורי |  |
| 1:3.220 | מפואר |  |
| 1:3.221 | ומסודר |  |
| 1:3.222 | בקשו |  |
| 1:3.223 | להם |  |
| 1:3.224 | חסרון |  |
| 1:3.225 | בדיבוק |  |
| 1:3.226 | כחוט |  |
| 1:3.227 | השערה |  |
| 1:3.228 | בקצת |  |
| 1:3.229 | אותיות |  |
| 1:3.230 | כאלו |  |
| 1:3.231 | הם |  |
| 1:3.232 | כותבים |  |
| 1:3.233 | סתו״מ | . |
| 1:3.234 | ולא |  |
| 1:3.235 | הביטו |  |
| 1:3.236 | ליופי |  |
| 1:3.237 | הכתב |  |
| 1:3.238 | וסדורו | . |
| 1:3.239 | ובפתרון |  |
| 1:3.240 | הכתובים |  |
| 1:3.241 | בקשו |  |
| 1:3.242 | מהם |  |
| 1:3.243 | לפתור |  |
| 1:3.244 | להם |  |
| 1:3.245 | מקראות |  |
| 1:3.246 | חמורים |  |
| 1:3.247 | כגון |  |
| 1:3.248 | אסרי |  |
| 1:3.249 | לגפן |  |
| 1:3.250 | עירה | , |
| 1:3.251 | וכיוצא |  |
| 1:3.252 | ויצאו |  |
| 1:3.253 | מלפניהם |  |
| 1:3.254 | בפחי |  |
| 1:3.255 | נפש |  |
| 1:3.256 | טובים |  |
| 1:3.257 | דוים |  |
| 1:3.258 | ועצבים |  |
| 1:3.259 | ורפויי |  |
| 1:3.260 | ידים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

