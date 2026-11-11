# Translation pack: milhamot-hashem perek-3-12 (Milchamot ha-Shem 3:19–3:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-12`.

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
- סנהדרין רי"ף: "Sanhedrin, [in] the Rif", as written; no volume or folio added.
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
- אלוה as the object of worship of a human being (קבלו באלוה): "a god"; "God"
  elsewhere (this refines the אלוה ruling above).
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


## Milchamot ha-Shem 3:19

Hebrew (Wikisource): יט) וכח זה לחקור ולדרוש בתורה על פי קבלת רבותינו בעלי המשנה והתלמוד מסור לנו מרבותינו כמו שאמר התנא הפוך בה והפוך בה. דכלא בה ובה תחזי. וסיב ובלי בה.! כי אם ילמוד האדם ולא יבין לא יתגלה לו הסוד ולא ישיג דבר. כדגרסינן בברכות פ"ק לעולם ירוץ אדם לדבר מצוה ואפי'בשבת שנאמר אחרי ה' ילכו כאריה ישאג וכו' אמר ר' זירא אגרא דפרקא רהטא. (פי' רש"י אגרא דפרקא עיקר קיבול שכר הבריות הרצים לשמוע דרשה מפי חכם הוא שכר המרוצה שהרי רובם אינם מבינים להעמיד גרסא ולמי השמועה לאחר זמן שיקבלו שכר לימוד) אגרא דכלה דוחקא (שבת שלפני הרגל שהכל נאספים לשמוע הלכוה הרגל) אגרא דשמעתא סברא - פי' שהוא יגע וטורח ומחשב להבין טעמו של דבר) ע"כ. וכסנהדרין גרסי' כי בתחבולות תעשה לך מלחמה אמר ר' אחא בר' חנינא אמר ר' אסי אמר ר' יוחנן במי אתה מוצא מלחמתה של תורה. במי שיש בידו חבילות של משנה וכו' ופי' רש"י מלחמתה של תורה הוראותיה ולעמוד על בוריה ועל עיקרה, לא כאדם המפולפל ומחודד ובעל סברא ולא למד משניות וברייתות הרבה. כי מהיכן יתגלה לו הסוג? אלא בבעל משניות הרבה שאם יצטרך לו טעם בכאן ילמדנו מתוך משנה אחרת. או אם יקשה לו דבר יבין מתוך משניות שבידו. הא מני ר' פלוני היא ששמענוהו במקום אחר אומר כן עכ"ל הגמ' ורש"י ז"ל. וגרסי נמי בפ"ק דחגיגה ואף הוא פתח ודרש דברי חכמים כדרבונות וכמסמרות נטועים בעלי אסופות כולם ניתנו מרועה אחד. למה נמשלו דברי תורה מכוונים לב לומדיהן מדרכי מיתה לדרכי חיים. אי מה דרבן זה וכו' בעלי אסופות אלו תלמידי חכמים שיושבין אסופות אסופות ועוסקין בתורה. הללו מטמאין והללו מטהרין. הללו אוסרין, והללו מתירין, הללו פוסלין והללו מכשירין שמא יאמר היאך אני למד תורה מעתה תלמוד לומר כולם ניתנו מרועה אחד, כולן אל אחד נתנן. ופרנס אחד אמרן, מפי אדון כל המעשים ברוך הוא שנ' וידבר אלהים את כל הדברים האלה לאמר, אף אתה עשה אזנך כאפרכסת וקנה לך לך מבין לשמוע דברי המטהרין ודברי המטמאין, דברי האוסרין ודברי המתירין, דברי הפוסלין ודברי המכשירין ע"כ ופי' רש"י ז"ל אל אחד אמרן, אין לך אחד בבני המחלוקת מביא ראיה מתורת אל אחר אלא מתורתו של הקב"ה, ופרנס אחד אמרן אין לך מביא ראיה מדברי נביא הבא לחלוק על משה רבי' ע"ה, עשה אזנך כאפרכסת, מאחר שכולן לבן לשמים עשה אזן שומעת ולמוד ודע דברי כולן, וכשתדע להבחין איזה יבשר קבע הלכה כמותו, עכ"ל, וכן כ' המהרש"א שזה משל לעניינו, שיכניס לאזנו לשמוע הרבה דברי כולן בין המטמאין בין המטהרין אבל יקנה לו לב מבין להבחין מדעות החלוקות איזה יכשר ויקבע לו הלכה כמותו וכפי' רש"י ז"ל עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:19.1 | יט | ) |
| 3:19.2 | וכח |  |
| 3:19.3 | זה |  |
| 3:19.4 | לחקור |  |
| 3:19.5 | ולדרוש |  |
| 3:19.6 | בתורה |  |
| 3:19.7 | על |  |
| 3:19.8 | פי |  |
| 3:19.9 | קבלת |  |
| 3:19.10 | רבותינו |  |
| 3:19.11 | בעלי |  |
| 3:19.12 | המשנה |  |
| 3:19.13 | והתלמוד |  |
| 3:19.14 | מסור |  |
| 3:19.15 | לנו |  |
| 3:19.16 | מרבותינו |  |
| 3:19.17 | כמו |  |
| 3:19.18 | שאמר |  |
| 3:19.19 | התנא |  |
| 3:19.20 | הפוך |  |
| 3:19.21 | בה |  |
| 3:19.22 | והפוך |  |
| 3:19.23 | בה | . |
| 3:19.24 | דכלא |  |
| 3:19.25 | בה |  |
| 3:19.26 | ובה |  |
| 3:19.27 | תחזי | . |
| 3:19.28 | וסיב |  |
| 3:19.29 | ובלי |  |
| 3:19.30 | בה | .! |
| 3:19.31 | כי |  |
| 3:19.32 | אם |  |
| 3:19.33 | ילמוד |  |
| 3:19.34 | האדם |  |
| 3:19.35 | ולא |  |
| 3:19.36 | יבין |  |
| 3:19.37 | לא |  |
| 3:19.38 | יתגלה |  |
| 3:19.39 | לו |  |
| 3:19.40 | הסוד |  |
| 3:19.41 | ולא |  |
| 3:19.42 | ישיג |  |
| 3:19.43 | דבר | . |
| 3:19.44 | כדגרסינן |  |
| 3:19.45 | בברכות |  |
| 3:19.46 | פ"ק |  |
| 3:19.47 | לעולם |  |
| 3:19.48 | ירוץ |  |
| 3:19.49 | אדם |  |
| 3:19.50 | לדבר |  |
| 3:19.51 | מצוה |  |
| 3:19.52 | ואפי'בשבת |  |
| 3:19.53 | שנאמר |  |
| 3:19.54 | אחרי |  |
| 3:19.55 | ה | ' |
| 3:19.56 | ילכו |  |
| 3:19.57 | כאריה |  |
| 3:19.58 | ישאג |  |
| 3:19.59 | וכו | ' |
| 3:19.60 | אמר |  |
| 3:19.61 | ר | ' |
| 3:19.62 | זירא |  |
| 3:19.63 | אגרא |  |
| 3:19.64 | דפרקא |  |
| 3:19.65 | רהטא | . |
| 3:19.66 | פי | ' |
| 3:19.67 | רש"י |  |
| 3:19.68 | אגרא |  |
| 3:19.69 | דפרקא |  |
| 3:19.70 | עיקר |  |
| 3:19.71 | קיבול |  |
| 3:19.72 | שכר |  |
| 3:19.73 | הבריות |  |
| 3:19.74 | הרצים |  |
| 3:19.75 | לשמוע |  |
| 3:19.76 | דרשה |  |
| 3:19.77 | מפי |  |
| 3:19.78 | חכם |  |
| 3:19.79 | הוא |  |
| 3:19.80 | שכר |  |
| 3:19.81 | המרוצה |  |
| 3:19.82 | שהרי |  |
| 3:19.83 | רובם |  |
| 3:19.84 | אינם |  |
| 3:19.85 | מבינים |  |
| 3:19.86 | להעמיד |  |
| 3:19.87 | גרסא |  |
| 3:19.88 | ולמי |  |
| 3:19.89 | השמועה |  |
| 3:19.90 | לאחר |  |
| 3:19.91 | זמן |  |
| 3:19.92 | שיקבלו |  |
| 3:19.93 | שכר |  |
| 3:19.94 | לימוד | ) |
| 3:19.95 | אגרא |  |
| 3:19.96 | דכלה |  |
| 3:19.97 | דוחקא |  |
| 3:19.98 | שבת |  |
| 3:19.99 | שלפני |  |
| 3:19.100 | הרגל |  |
| 3:19.101 | שהכל |  |
| 3:19.102 | נאספים |  |
| 3:19.103 | לשמוע |  |
| 3:19.104 | הלכוה |  |
| 3:19.105 | הרגל | ) |
| 3:19.106 | אגרא |  |
| 3:19.107 | דשמעתא |  |
| 3:19.108 | סברא | - |
| 3:19.109 | פי | ' |
| 3:19.110 | שהוא |  |
| 3:19.111 | יגע |  |
| 3:19.112 | וטורח |  |
| 3:19.113 | ומחשב |  |
| 3:19.114 | להבין |  |
| 3:19.115 | טעמו |  |
| 3:19.116 | של |  |
| 3:19.117 | דבר | ) |
| 3:19.118 | ע"כ | . |
| 3:19.119 | וכסנהדרין |  |
| 3:19.120 | גרסי | ' |
| 3:19.121 | כי |  |
| 3:19.122 | בתחבולות |  |
| 3:19.123 | תעשה |  |
| 3:19.124 | לך |  |
| 3:19.125 | מלחמה |  |
| 3:19.126 | אמר |  |
| 3:19.127 | ר | ' |
| 3:19.128 | אחא |  |
| 3:19.129 | בר | ' |
| 3:19.130 | חנינא |  |
| 3:19.131 | אמר |  |
| 3:19.132 | ר | ' |
| 3:19.133 | אסי |  |
| 3:19.134 | אמר |  |
| 3:19.135 | ר | ' |
| 3:19.136 | יוחנן |  |
| 3:19.137 | במי |  |
| 3:19.138 | אתה |  |
| 3:19.139 | מוצא |  |
| 3:19.140 | מלחמתה |  |
| 3:19.141 | של |  |
| 3:19.142 | תורה | . |
| 3:19.143 | במי |  |
| 3:19.144 | שיש |  |
| 3:19.145 | בידו |  |
| 3:19.146 | חבילות |  |
| 3:19.147 | של |  |
| 3:19.148 | משנה |  |
| 3:19.149 | וכו | ' |
| 3:19.150 | ופי | ' |
| 3:19.151 | רש"י |  |
| 3:19.152 | מלחמתה |  |
| 3:19.153 | של |  |
| 3:19.154 | תורה |  |
| 3:19.155 | הוראותיה |  |
| 3:19.156 | ולעמוד |  |
| 3:19.157 | על |  |
| 3:19.158 | בוריה |  |
| 3:19.159 | ועל |  |
| 3:19.160 | עיקרה | , |
| 3:19.161 | לא |  |
| 3:19.162 | כאדם |  |
| 3:19.163 | המפולפל |  |
| 3:19.164 | ומחודד |  |
| 3:19.165 | ובעל |  |
| 3:19.166 | סברא |  |
| 3:19.167 | ולא |  |
| 3:19.168 | למד |  |
| 3:19.169 | משניות |  |
| 3:19.170 | וברייתות |  |
| 3:19.171 | הרבה | . |
| 3:19.172 | כי |  |
| 3:19.173 | מהיכן |  |
| 3:19.174 | יתגלה |  |
| 3:19.175 | לו |  |
| 3:19.176 | הסוג | ? |
| 3:19.177 | אלא |  |
| 3:19.178 | בבעל |  |
| 3:19.179 | משניות |  |
| 3:19.180 | הרבה |  |
| 3:19.181 | שאם |  |
| 3:19.182 | יצטרך |  |
| 3:19.183 | לו |  |
| 3:19.184 | טעם |  |
| 3:19.185 | בכאן |  |
| 3:19.186 | ילמדנו |  |
| 3:19.187 | מתוך |  |
| 3:19.188 | משנה |  |
| 3:19.189 | אחרת | . |
| 3:19.190 | או |  |
| 3:19.191 | אם |  |
| 3:19.192 | יקשה |  |
| 3:19.193 | לו |  |
| 3:19.194 | דבר |  |
| 3:19.195 | יבין |  |
| 3:19.196 | מתוך |  |
| 3:19.197 | משניות |  |
| 3:19.198 | שבידו | . |
| 3:19.199 | הא |  |
| 3:19.200 | מני |  |
| 3:19.201 | ר | ' |
| 3:19.202 | פלוני |  |
| 3:19.203 | היא |  |
| 3:19.204 | ששמענוהו |  |
| 3:19.205 | במקום |  |
| 3:19.206 | אחר |  |
| 3:19.207 | אומר |  |
| 3:19.208 | כן |  |
| 3:19.209 | עכ"ל |  |
| 3:19.210 | הגמ | ' |
| 3:19.211 | ורש"י |  |
| 3:19.212 | ז"ל | . |
| 3:19.213 | וגרסי |  |
| 3:19.214 | נמי |  |
| 3:19.215 | בפ"ק |  |
| 3:19.216 | דחגיגה |  |
| 3:19.217 | ואף |  |
| 3:19.218 | הוא |  |
| 3:19.219 | פתח |  |
| 3:19.220 | ודרש |  |
| 3:19.221 | דברי |  |
| 3:19.222 | חכמים |  |
| 3:19.223 | כדרבונות |  |
| 3:19.224 | וכמסמרות |  |
| 3:19.225 | נטועים |  |
| 3:19.226 | בעלי |  |
| 3:19.227 | אסופות |  |
| 3:19.228 | כולם |  |
| 3:19.229 | ניתנו |  |
| 3:19.230 | מרועה |  |
| 3:19.231 | אחד | . |
| 3:19.232 | למה |  |
| 3:19.233 | נמשלו |  |
| 3:19.234 | דברי |  |
| 3:19.235 | תורה |  |
| 3:19.236 | מכוונים |  |
| 3:19.237 | לב |  |
| 3:19.238 | לומדיהן |  |
| 3:19.239 | מדרכי |  |
| 3:19.240 | מיתה |  |
| 3:19.241 | לדרכי |  |
| 3:19.242 | חיים | . |
| 3:19.243 | אי |  |
| 3:19.244 | מה |  |
| 3:19.245 | דרבן |  |
| 3:19.246 | זה |  |
| 3:19.247 | וכו | ' |
| 3:19.248 | בעלי |  |
| 3:19.249 | אסופות |  |
| 3:19.250 | אלו |  |
| 3:19.251 | תלמידי |  |
| 3:19.252 | חכמים |  |
| 3:19.253 | שיושבין |  |
| 3:19.254 | אסופות |  |
| 3:19.255 | אסופות |  |
| 3:19.256 | ועוסקין |  |
| 3:19.257 | בתורה | . |
| 3:19.258 | הללו |  |
| 3:19.259 | מטמאין |  |
| 3:19.260 | והללו |  |
| 3:19.261 | מטהרין | . |
| 3:19.262 | הללו |  |
| 3:19.263 | אוסרין | , |
| 3:19.264 | והללו |  |
| 3:19.265 | מתירין | , |
| 3:19.266 | הללו |  |
| 3:19.267 | פוסלין |  |
| 3:19.268 | והללו |  |
| 3:19.269 | מכשירין |  |
| 3:19.270 | שמא |  |
| 3:19.271 | יאמר |  |
| 3:19.272 | היאך |  |
| 3:19.273 | אני |  |
| 3:19.274 | למד |  |
| 3:19.275 | תורה |  |
| 3:19.276 | מעתה |  |
| 3:19.277 | תלמוד |  |
| 3:19.278 | לומר |  |
| 3:19.279 | כולם |  |
| 3:19.280 | ניתנו |  |
| 3:19.281 | מרועה |  |
| 3:19.282 | אחד | , |
| 3:19.283 | כולן |  |
| 3:19.284 | אל |  |
| 3:19.285 | אחד |  |
| 3:19.286 | נתנן | . |
| 3:19.287 | ופרנס |  |
| 3:19.288 | אחד |  |
| 3:19.289 | אמרן | , |
| 3:19.290 | מפי |  |
| 3:19.291 | אדון |  |
| 3:19.292 | כל |  |
| 3:19.293 | המעשים |  |
| 3:19.294 | ברוך |  |
| 3:19.295 | הוא |  |
| 3:19.296 | שנ | ' |
| 3:19.297 | וידבר |  |
| 3:19.298 | אלהים |  |
| 3:19.299 | את |  |
| 3:19.300 | כל |  |
| 3:19.301 | הדברים |  |
| 3:19.302 | האלה |  |
| 3:19.303 | לאמר | , |
| 3:19.304 | אף |  |
| 3:19.305 | אתה |  |
| 3:19.306 | עשה |  |
| 3:19.307 | אזנך |  |
| 3:19.308 | כאפרכסת |  |
| 3:19.309 | וקנה |  |
| 3:19.310 | לך |  |
| 3:19.311 | לך |  |
| 3:19.312 | מבין |  |
| 3:19.313 | לשמוע |  |
| 3:19.314 | דברי |  |
| 3:19.315 | המטהרין |  |
| 3:19.316 | ודברי |  |
| 3:19.317 | המטמאין | , |
| 3:19.318 | דברי |  |
| 3:19.319 | האוסרין |  |
| 3:19.320 | ודברי |  |
| 3:19.321 | המתירין | , |
| 3:19.322 | דברי |  |
| 3:19.323 | הפוסלין |  |
| 3:19.324 | ודברי |  |
| 3:19.325 | המכשירין |  |
| 3:19.326 | ע"כ |  |
| 3:19.327 | ופי | ' |
| 3:19.328 | רש"י |  |
| 3:19.329 | ז"ל |  |
| 3:19.330 | אל |  |
| 3:19.331 | אחד |  |
| 3:19.332 | אמרן | , |
| 3:19.333 | אין |  |
| 3:19.334 | לך |  |
| 3:19.335 | אחד |  |
| 3:19.336 | בבני |  |
| 3:19.337 | המחלוקת |  |
| 3:19.338 | מביא |  |
| 3:19.339 | ראיה |  |
| 3:19.340 | מתורת |  |
| 3:19.341 | אל |  |
| 3:19.342 | אחר |  |
| 3:19.343 | אלא |  |
| 3:19.344 | מתורתו |  |
| 3:19.345 | של |  |
| 3:19.346 | הקב"ה | , |
| 3:19.347 | ופרנס |  |
| 3:19.348 | אחד |  |
| 3:19.349 | אמרן |  |
| 3:19.350 | אין |  |
| 3:19.351 | לך |  |
| 3:19.352 | מביא |  |
| 3:19.353 | ראיה |  |
| 3:19.354 | מדברי |  |
| 3:19.355 | נביא |  |
| 3:19.356 | הבא |  |
| 3:19.357 | לחלוק |  |
| 3:19.358 | על |  |
| 3:19.359 | משה |  |
| 3:19.360 | רבי | ' |
| 3:19.361 | ע"ה | , |
| 3:19.362 | עשה |  |
| 3:19.363 | אזנך |  |
| 3:19.364 | כאפרכסת | , |
| 3:19.365 | מאחר |  |
| 3:19.366 | שכולן |  |
| 3:19.367 | לבן |  |
| 3:19.368 | לשמים |  |
| 3:19.369 | עשה |  |
| 3:19.370 | אזן |  |
| 3:19.371 | שומעת |  |
| 3:19.372 | ולמוד |  |
| 3:19.373 | ודע |  |
| 3:19.374 | דברי |  |
| 3:19.375 | כולן | , |
| 3:19.376 | וכשתדע |  |
| 3:19.377 | להבחין |  |
| 3:19.378 | איזה |  |
| 3:19.379 | יבשר |  |
| 3:19.380 | קבע |  |
| 3:19.381 | הלכה |  |
| 3:19.382 | כמותו | , |
| 3:19.383 | עכ"ל | , |
| 3:19.384 | וכן |  |
| 3:19.385 | כ | ' |
| 3:19.386 | המהרש"א |  |
| 3:19.387 | שזה |  |
| 3:19.388 | משל |  |
| 3:19.389 | לעניינו | , |
| 3:19.390 | שיכניס |  |
| 3:19.391 | לאזנו |  |
| 3:19.392 | לשמוע |  |
| 3:19.393 | הרבה |  |
| 3:19.394 | דברי |  |
| 3:19.395 | כולן |  |
| 3:19.396 | בין |  |
| 3:19.397 | המטמאין |  |
| 3:19.398 | בין |  |
| 3:19.399 | המטהרין |  |
| 3:19.400 | אבל |  |
| 3:19.401 | יקנה |  |
| 3:19.402 | לו |  |
| 3:19.403 | לב |  |
| 3:19.404 | מבין |  |
| 3:19.405 | להבחין |  |
| 3:19.406 | מדעות |  |
| 3:19.407 | החלוקות |  |
| 3:19.408 | איזה |  |
| 3:19.409 | יכשר |  |
| 3:19.410 | ויקבע |  |
| 3:19.411 | לו |  |
| 3:19.412 | הלכה |  |
| 3:19.413 | כמותו |  |
| 3:19.414 | וכפי | ' |
| 3:19.415 | רש"י |  |
| 3:19.416 | ז"ל |  |
| 3:19.417 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

