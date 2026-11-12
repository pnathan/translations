# Translation pack: milhamot-hashem perek-4-03 (Milchamot ha-Shem 4:5–4:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-03`.

## Rulings for this unit

Extracted from `works/milhamot-hashem/decisions.md` for the words of this unit; the rulings below are the whole of what you need. Read the full file only if a question is not answered here.

# Milchamot ha-Shem: decisions for every unit

> Owner answered Q4 and Q7 on 2026-11-10; Q1-3, Q5-6 stand at the defaults (see "Owner's answers and defaults" at the end).

Editorial decisions set before translation. A reviewer may question one in a
`question` patch; the project owner decides. Where a ruling was amended, only
the final rule is given. "(note)" means a token `note`; "first in a unit"
means at the first occurrence in the unit.

General rule for technical terms: the English names the thing the words name
and goes no further. Kabbalistic technical names are transliterated, because
every English equivalent takes a side: the Kabbalists' defence
("configuration", "the Infinite", "Divine Presence") or Kafich's charge
("god", "idol"). He makes his charge in his own words ("אליל", "אלוהות", "אל
זר"), and the English translates those where he uses them.

Transliteration: the project's plain scheme, as in Avot and Genesis (Shetach,
Chayei Sarah, Lekh Lekha): no diacritics; ח = ch, כ = kh, צ = tz, ק = k
("Machzor Vitry"). The author: Yichya Kafich (owner, 2026-11-10; "Yihya
Kafich" is named once in the front matter). Book titles with a conventional
English spelling keep it (Kisse Eliyahu, Sefer ha-Kavanot, Machberet
ha-Kodesh); the scheme applies where the project has no fixed form.

## Output format for this work (prose)

Read this whole file before translating. A segment is a paragraph of prose.
The pack has no morphemes, lemmas, lexicon or commentary, so:

- `tokens`: one entry per word: `id`, `gloss`, and a `note` where a learner
  needs one. No `sense`, no `lex`.
- `commentary`: `[]`. `context`: only for a realia a reader needs (a Yemenite
  custom, a book, a person); say what is established; cite no modern work.
  It may carry a short note on a story the author cites (the Pardes story),
  without a reference unless the author gives one.
- `grammar`: 2 to 6 notes per paragraph; in a long paragraph, notes on the
  hardest sentences, not on all. `literal`: a word-for-word line. On a long
  paragraph, `literal` is the glosses in Hebrew order, hyphens joining the
  parts of one word; no reasoning in it (reasons go in `note`).
- `phrases` divide every token of the paragraph, in order. In a long
  paragraph make a phrase a clause, not a word, so that the file stays
  checkable; each phrase's `english` must occur in the paragraph's `english`.
- Quoted passages (Bible, Talmud, Zohar, Judeo-Arabic) follow the Quotations
  section; Bible verses are translated from the Hebrew of the verse as the
  author quotes it and the reference goes in `apparatus`.
- Every Hebrew abbreviation is expanded in `gloss` (the reading line carries
  the full word).

## Names and titles

- ה׳: the Tetragrammaton → "the LORD"; cited as a name ("זעיר אנפין דאיקרי
  ה'") → "who is called 'the LORD'"; before a personal name → "Mr." ("מאת ה'
  אליעזר בן יהודה" → "from Mr. Eliezer Ben-Yehuda"); a numeral or page letter
  → the number. The pack splits it into ה plus ': read it in context.
- **הקב"ה / קב"ה / קודשא בריך הוא** → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **Honorifics**: ז"ל → "of blessed memory"; ע"ה → "peace be upon him"; רז"ל →
  "our Rabbis of blessed memory"; חז"ל → "our Sages of blessed memory"; זלה"ה
  → "may his memory be for the life of the world to come"; זצוק"ל → "may the
  memory of the righteous and holy be for a blessing"; ת"ו (Jerusalem) → "may
  it be rebuilt and established"; יר"ה (the Imam) → "may his majesty be
  exalted". Translate every one, opponents' included; add none.
- **Authorities.** Acronyms and book-names used as names stay, transliterated
  (the Rambam, Ramban, Rashi, Ari, Ribash, Rashba, Rosh (הרא"ש), Ravad, Rema, Radbaz,
  Maharshal, Maharam al-Shakar, Ra'em, Maran, Roke'ach, Lechem Mishneh, Chavot
  Yair, Pitchei Teshuvah, Beit Yosef, Kenesset ha-Gedolah; מהר"י + surname →
  "Mahari Lopes"); first in a unit a note gives the expansion where certain
  (else says so) and who is meant; identify no one further. Written out: רס"ג →
  "Rabbi Saadia Gaon"; הרח"ו → "Rabbi Chayyim Vital"; ר"ת → "Rabbeinu Tam";
  ר' + name → "Rabbi" (בר' בחיי → "Rabbi Bachya"); רבינו + name → "Rabbeinu"
  ("Rabbeinu Bachya"; בחיי alone "Bachya"); הר' המורה / רבינו meaning the
  Rambam → "the Master, the Guide" / "our Master"; רבי' before the lawgiver's
  name → "Moses our teacher"; הגאון alone → "the Gaon"; בעל חובת הלבבות →
  "the author of Chovot ha-Levavot"; בעל העיקרים → "the author of the
  Ikkarim"; סמ"ג → "the Semag (Sefer Mitzvot Gadol)". ראב"ע → "Rabbi Eleazar
  ben Azariah" in a Talmudic passage, "Ibn Ezra" for the commentator (note).
  ר"א → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note). An initial a surname fixes is expanded (ר"מ ריקנטי → "Rabbi Menachem
  Rikanti"), else kept (ר"י חיים → "Rabbi Y. Chayyim"; a lone כ' → "K."), with
  a note; יש"ר מקנדי"א → "Yish'ar of Candia" (note: conjectural). Surnames as printed (אלקבולי → "Alkaboli", Alkabetz
  in the note), one spelling across the work. Turkish and Arabic names: the
  scheme as spelled; keep Effendi, Bey, Pasha; identify no one.
- **Book titles**: transliterated (Kisse Eliyahu, Yosher Levav, Etz Chayyim,
  Mikdash Melekh, Emunot ve-De'ot, Moreh Nevukhim, also for "Guide for the
  Perplexed", the Kuzari, Sefer ha-Berit, Tikkunei Zohar, Zohar Chadash, Idra
  of Naso), the meaning in a
  note first in a unit; abbreviations written out (כס"א → Kisse Eliyahu); every word
  of a cited title kept. Bible books in English; tractates transliterated.
  ממלכי → "of the kings of". A misspelt title (ונור אלצלאם) as printed ("Nur
  al-Tzalam"), the likelier (Nur al-Zulam, Me'or ha-Afelah) in the note and
  `apparatus`.
- **הר' / הרב / הר' כסא אליהו / הר' עץ יוסף / הרב לחם משנה / המקובל בעל היכל הברכה**:
  הר' / הרב + book title → "the Rav [title]" ("the Rav Lechem Mishneh", "the
  Rav Mikdash Melekh"), never dropping הרב, note on who is meant; a title is
  never a personal name. "the author of [title]" only for בעל / מחבר ("the
  Kabbalist, the author of Heikhal ha-Berakhah", title translated in a note);
  name no author the text does not name.
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

- **אין סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **The names of the partzufim**: transliterated, the literal sense in a note
  first in a unit: Ze'ir Anpin (Small Face; ז"א → "Ze'ir Anpin"; זעיר alone →
  "Ze'ir"); Arikh Anpin (Long Face); Abba; Imma; Nukva (נוקביה → "his Nukva");
  Atik; עתיקא קדישא → "Atika Kadisha"; עתיקא סתימא → "Atika Setima"; עתיק
  יומין → "Atik Yomin" (Dan 7:9: "the Ancient of Days"). Translated: סבא דסבין
  → "the Elder of Elders"; עלת העלות → "the Cause of causes"; הסבה הראשונה →
  "the First Cause"; כתר עליון → "Keter Elyon (the Supreme Crown)". An
  argument from a name's meaning goes in the note; the name stays.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **Other recurrent terms**, fixed: השפע → "the flow"; בחינה → "aspect";
  השתלשלות → "the chain of descent"; עלה / עלול → "cause / effect"; מחויב
  המציאות → "the Necessary Existent"; כלים (Kabbalistic) → "vessels" (a
  craftsman's כלי → "tool"); אורות דקים → "subtle lights"; גוף דק → "a subtle
  body"; כוונה (in prayer) → "intention"; זווג → "coupling"; השגחה →
  "providence"; "לית מחשבה תפיסה ביה" → "no thought grasps Him"; מין / מינות
  → "heretic / heresy"; אפיקורוס / אפיקורסים → "apikoros / apikorsim"; כופר
  בעיקר → "denies the fundamental principle"; ספר חיצון → "an external book"
  (note: m. Sanh. 10:1); תורה שבעל פה → "the Oral Torah"; ימי הבינים → "the
  Middle Ages"; "אין לו חלק לעולם הבא" → "has no share in the world to come"
  (lower case); bat kol → "bat kol" (note: "daughter of a voice").
- **Abbreviations.** The ambiguous ones by context, the expansion in the note:
  ר"ל → "that is to say" (with a named authority as subject: "means to say"),
  "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a sentence
  end); ע"כ → "thus far" (after a quotation), "therefore", "perforce" (על
  כרחך), "and upon all" ("וע"כ עמו ישראל"); ע"ש → "see there" or "on account
  of" (על שם), כ' ע"ש → "he writes, see there:" (על שם in `readings`); ב"ו →
  "blessed be He" before הוא, else "flesh and blood"; מ"ו → "our teacher" in
  יודיענו מורנו; ב"ע → "in Aramaic" against בערבי (others in `readings`).
  Fixed: ח"ו → "God forbid"; וכו' → "etc." (always kept; וכ' closing a
  quotation is וכו'); ג"כ → "also"; כ"ש / ק"ו → "all the more so"; ת"ל →
  "Scripture says"; א"כ → "if so"; סתו"מ → "Torah scrolls, tefillin and
  mezuzot" (note); תוה"ק → "the holy Torah"; ביהכנ"ס → "synagogue". Unruled
  (ורבי'): by his usage, with a note; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if it is the sole sense that fits.
  הינ"ץ → "H.N.Ts." (note: perhaps ה + ינ"ץ, "may his Rock protect him", of
  the Imam); no person named.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **אין לנו להחמיר** → "we may not be stringent".
- **הגיע חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
## Quotations

- **Scope.** Everything quoted is translated from the words the author gives
  (Zohar, Bavli, Yerushalmi, Targum, Biblical Aramaic, Midrash, medieval
  authors, Judeo-Arabic). No Hebrew, Aramaic or Arabic in `english`.
- **Marking.** A source's own words introduced by a formula (וז"ל, כתב, אמרו,
  גרסינן, שנאמר, כדכתיב, שנינו, הכי קאמר, אמר ר' X) or closed by עכ"ל / ע"כ
  take double quotation marks; a paraphrase or a report ("כתב ש…") none. No
  formula: marks if the source is certain, else none and a note. Nesting:
  double, single, double. A reply by ויאמר inside a quotation: double marks;
  his report of a sage's question: unmarked. A quotation over several
  paragraphs (2:18-2:25): each paragraph opens with a mark, only the last
  closes. `english` is plain text (no italics, no labels); the language goes
  in `grammar` ("Zohar Aramaic", "Judeo-Arabic"), the word-for-word in
  `literal`.
- **Formulas**, translated, outside the marks: וז"ל → "and these are his
  words:"; וז"ל + a name → "and these are the words of X:"; עכ"ל → "Thus far
  his words."; ע"כ → "Thus far."; "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus
  far the wording of the Zohar, with the explanation of the RaShB in the
  margin."
- **Glosses inside a quotation** ("(פירש בגליון …)", "(פי' …)") stay in
  parentheses where they stand, translated: "(the margin explains: …)",
  "(explanation: …)".
- **Zohar formulas**, fixed: זכאה חולקיה / חולקיהון → "happy is his / their
  portion"; זכאה חולקיהון דישראל → "happy is the portion of Israel"; ווי ליה מאן דלא → "woe to him who does not"; דמאן דעביד כך וכך →
  "of one who does such-and-such"; תא חזי → "come and see"; רזא דמלה → "the
  secret of the matter"; פתח ר' שמעון → "Rabbi Shimon opened". Quoted as
  sales talk (9:4, 9:16): in quotation marks, no irony added. Variant spellings
  ("הולקיה", "דכאה"): the evident word, recorded in `apparatus`.
- **Exegesis on a Hebrew word** (את; אם אין → "Ayin"): the word stays
  transliterated ("not 'to the LORD your God' but 'et the LORD your God'"),
  explained in `grammar`.
- **Cut or changed quotation.** Only what is on the page: do not supply what
  וכו' omits, complete a verse (no "before Me"; ואלהי → "and the gods of",
  note) or restore a source; a verse's fixed wording covers only the words
  quoted. His comma joining two verses stays; an ellipsis only where he marks
  an omission (וכו', וגו'). Where his wording differs from the verse or source
  ("אבותיהם" for "אבתיכם"), follow him; the source in `apparatus`.
- **Bible verses**, quoted or alluded to (*melitzah*, unmarked): from the
  Hebrew as it stands, plain modern English, not King James; every verse in
  `apparatus`. An echoed verse (אל דבר ה', בטרם תחיל) governs the sense of its
  words. A poetic phrase (הר המור) is translated ("the mountain of myrrh"), a
  note naming the referent. A plural stays plural (עמדו, Job 32:16), no
  subject added. Ketiv ("מי אתי") translated, the qere noted. Fixed:
  - "בחלק שפתיו" → "with the smoothness of his lips";
  - "כמתלהלה היורה זיקים חצים ומות" → "like a madman who shoots firebrands,
    arrows and death";
  - "צנומות דקות שדופות קדים" → "withered, thin, blasted by the east wind";
  - "בחבלי השוא … כעבות העגלה" → "with cords of falsehood … as with
    cart-ropes";
  - "כעיר פרוצה אין חומה" → "like a city breached, without a wall";
  - "להוציא ישן … מפני חדש" → "to clear out the old before the new";
  - "מילדי נכרים" → "from the children of strangers"; "הבל וריק" → "vanity
    and emptiness"; תהו ובהו → "formlessness and void";
  - ריש מלין אמר (Dan 7:1) → "he stated the sum of the matter";
  - קצוצי פאה (Jer 9:25) → "those with clipped temples" (context note: desert
    peoples);
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
  - Exod 34:6: "the LORD, the LORD, God merciful and gracious, slow to anger".
  - Jer 10:11: "Thus shall you say to them: the gods that did not make the
    heavens and the earth shall perish from the earth and from under these
    heavens."
  - Lev 19:16: "stand by the blood of your fellow" (no "idly").
- **Idioms in his own Hebrew**, by sense, the literal in `literal` ("דמעיילי
  פילא בקופא דמחטא" → "who bring an elephant through the eye of a needle";
  "ואפריין נמטיה" → "and we acknowledge our debt to him"; "הני מילי דמרפסן
  איגרא" → "words that split roofs"; "בהא סלקינן ובהא נחתינן" → "this is where
  we come out"; "בנדון דידן" → "in our case"). A saying he glosses himself: both translated, unmerged.
- **Medieval authors** (Rambam, Saadia, Halevi, Bachya, Semag, Albo, Roke'ach,
  Shevilei Emunah): from the Hebrew he quotes, not the originals, standard
  editions or published translations; a certain difference may be noted in
  `apparatus`.
- **Judeo-Arabic** (Saadia on Job; Choter ben Shelomo; a Muslim polemic, "ספר
  אצהאר אלחק" = "גילוי האמת"; the proverb in 2:13): from the Arabic as it
  stands, never transliterated in the English. Ambiguous spelling: the evident
  word (note), the other in `apparatus`; אלי' (alif maqsura) is ilā
  "to"; a scribal letter change (ר for ד) noted, not corrected; an unexplained
  word (שברא): "[unclear]", the letters in `apparatus`. A rendering from
  Saadia's translation (אלבלד שברא): as given, Saadia's translation named in
  `apparatus`; his אמנתא / אמונתנו pairs: the Arabic as written, his Hebrew
  gloss in a note. Glosses inside the Arabic ("אתתלית (תלת קשרי דמהימנותא)")
  stay in parentheses, translated as his. The polemic's title as each place
  writes it ("Iẓhar al-Chaqq", "Gilui ha-Emet"), a note that they are one
  book; its author unnamed. His own Hebrew translations (10:7–10:9; "[תרגום
  הרב המחבר: …]") are separate segments, each from its own Hebrew; a
  difference in sense in `apparatus`. "כת המיחדים" = "אהל אלתוחיד" → "the
  upholders of the unity". (Q4.)
- **References** as he gives them, in English form, Arabic numerals:
  דף → "folio"; ע"א / ע"ב → "a / b"; ח"א → "vol. 1"; פ' / פרשת → the
  parashah ("Zohar, Balak, folio 191b"); "זהר בראשית כ״ב" → "Zohar, Bereshit
  22" (no "folio" he does not write); שער / ענף / מאמר → "Gate / Branch /
  Treatise"; "Sanhedrin 99"; סנהדרין רי"ף → "Sanhedrin, [in] the Rif". פ' +
  numeral → "chapter" for a book without parashot or a numeral beyond them,
  else "parashah" (noted); פ' + a chapter title (פ' ד' מיתות) → "Four Death
  Penalties". ד' + number after a book: the numerals ("4, 12"), noted when he
  writes דף nearby. הל' X after the Rambam → "Laws of X" (noted once per unit:
  Mishneh Torah). A cited book's סי' → "section 172", never "§". Add,
  complete, convert or correct nothing ("בשלח דף ס"ד (ח"ב סז, א)" →
  "Be-shallach, folio 64 (vol. 2, 67a)"); never the 2015 PDF's references;
  nothing for unreferenced Zohar, Mishnah or Talmud, even in `apparatus`
  (Bible verses only).
- **Section numbers and markers.** His own sections (ב) at 3:1 to קמד) at
  10:6): the head label → "§2."; "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58". Irregular numbers as written ("קיה)" → "§115", letters
  noted; the gap קכט)–קלא) noted, not repaired). List markers (א), ב)) →
  "1)", "2)" (note).

## Style

- Plain, exact, modern English, the words at their own strength (אליל "idol",
  סכלות והוללות "folly and madness"). No sarcasm of your own: no scare quotes,
  "so-called", "sic", or added "!". Keep every "!" and rhetorical question and
  his disclaimers ("(לפי דבריו)" → "(according to his words)"; "(לפי דעתם)" →
  "(in their opinion)"; "רחמנא ליצלן" → "the Merciful One save us").
- Address and voices: "ידידי" → "my friend"; "קורא נעים" (or "נעלם") →
  "pleasant reader"; "אחי" → "my brother"; second person kept. Chapter 2 is a
  correspondence: שאול שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב →
  "The questioner replied again"; המשיב → "the respondent"; each letter in its
  writer's voice, the parties identified no further than the text.
- Long periods may be divided: no clause dropped, no connective that changes
  the logic, the order kept. Hebrew full stops: שכל / שהוא opening a sentence
  keeps the stop and gets a main clause by his syntax (note on the ש); a
  fragment (ושהם, והפטרה, ולדעת, ורבינו) takes a main clause when it continues the
  verb before, a comma when it continues a list (note); a main clause after a
  stop (ואם ראיתי. הם מגלים): join with a comma (note); אם ... או after a
  stop: as punctuated if the English has a main clause, else a colon to the
  sentence before or a direct alternative question; a period completed in the
  next segment: keep the stops, the continuation in the segment `note`,
  semicolons only where unreadable.
- ל before a person's name is possessive (לרשב"י → "Rabbi Shimon's"), never
  the subject.
- Numbers: letter numerals → Arabic, as written whatever the tradition counts
  (the usual count in a note). A year without thousands → full ("תר"ע" →
  "5670"; "התר"ן לפרט גדול" → "5650 of the full count"); the civil year in
  `apparatus`. Where gematria is the point, the number in English, the letters
  in a note.
- Witnesses: square brackets ("[מתחלת באלוהיות]") are the transcriber's
  corrections: translate the corrected reading, record both. "(חסר)" →
  "(missing)", a lacuna sign; "(חטר)" is a variant of it, recorded. A word
  that does not fit but is close to a likelier one (יכינו for יבינו): by the
  context; a printed word echoing a verse's likelier word (נבוש for נמוש):
  the printed word; either way both words in a note and `apparatus`. Rule 3
  covers every Wikisource/PDF difference that bears on the English.
- Non-technical 1931 Hebrew (בית הספר, הסתדרות, מכתבי עתים): "school",
  "organisation", "newspapers".

## Translator's checklist (the pilot reviews found these faults most often)

Before running the check, read your English once against the Hebrew:
1. Every Hebrew word has its English; none is added (no "who received it", no
   parenthetical explanation). A Hebrew full stop is an English full stop, and
   every sentence has a main clause.
2. Recognise Bible verses, even in pieces (Isa 62:6, Job 37:18, Gen 43:14,
   Lev 26:44); translate from the Hebrew, check each verb's form, list the
   verse in `apparatus`.
3. Use the fixed renderings of this file for every term, name and abbreviation.
4. A passive participle is passive; a participle is not a perfect.
5. Capitalise only what the file capitalises. Begin each sentence with a
   capital.

## Owner's answers and defaults

Answered by the owner (2026-11-10):

- Q4: Judeo-Arabic and Aramaic are Sonnet's, not Haiku's. `models.json` lists
  the units Sonnet translates: the four Judeo-Arabic units (5:3, 7:24, 9:10,
  2:13) and the units with at least six Aramaic marker words (15 of 179; a
  heuristic; the Sonnet review covers every unit). The `lang` tags stay "heb".
- Q7: the project's transliteration scheme (above).

Not yet answered; the default stands and can be changed before the review
stage (a change to 1, 3 or 5 is a find-and-patch):

- Q1 title: "The Wars of the LORD" (Num 21:14).
- Q2 Wikisource additions: translated as base text; the pack marks none, the
  translator adds no flag per word, one general line in `apparatus` on the
  unit's first segment until the 1931 scan is collated.
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

