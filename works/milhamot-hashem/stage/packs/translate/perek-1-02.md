# Translation pack: milhamot-hashem perek-1-02 (Milchamot ha-Shem 1:3–1:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-02`.

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

- **הצעיר יחיה בן סלי' אלקאפח יצ״ו** → "The humble Yichya son of Sali'
  al-Kafich, may his Rock keep him and give him life" (patronymic as written,
  Q6).
- **מחבר הזהר / מחבר הספר הזהר / הפילסוף / נביא השקר / המסית / המדיח**: "the
  author of the Zohar", never named (no "Moses de León"); מחבר הספר הזהר →
  "the author of the book the Zohar". Fixed epithets: הפילסוף/הפילוסוף →
  "the philosopher"; נביא השקר → "the false prophet"; המסית → "the enticer";
  המדיח → "the seducer" (legal categories of Deut 13 and 18, on which he
  builds a case): never "deceiver" or "charlatan".
- ה׳: the Tetragrammaton → "the LORD"; cited as a name ("זעיר אנפין דאיקרי
  ה'") → "who is called 'the LORD'"; before a personal name → "Mr." ("מאת ה'
  אליעזר בן יהודה" → "from Mr. Eliezer Ben-Yehuda"); a numeral or page letter
  → the number. The pack splits it into ה plus ': read it in context.
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
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

- **The names of the partzufim**: transliterated, the literal sense in a note
  first in a unit: Ze'ir Anpin (Small Face; ז"א → "Ze'ir Anpin"; זעיר alone →
  "Ze'ir"); Arikh Anpin (Long Face); Abba; Imma; Nukva (נוקביה → "his Nukva");
  Atik; עתיקא קדישא → "Atika Kadisha"; עתיקא סתימא → "Atika Setima"; עתיק
  יומין → "Atik Yomin" (Dan 7:9: "the Ancient of Days"). Translated: סבא דסבין
  → "the Elder of Elders"; עלת העלות → "the Cause of causes"; הסבה הראשונה →
  "the First Cause"; כתר עליון → "Keter Elyon (the Supreme Crown)". An
  argument from a name's meaning goes in the note; the name stays.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה זרה / ע"ז / עכו"ם / עבודה אחרת**: עבודה זרה → "idolatry" (note once:
  "strange worship"); the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of
  idols". עכו"ם → "idolaters" (note: "worshippers of stars and
  constellations"), never "gentiles" or "Christians", even of a Christian king
  or the Trinity. עבודה אחרת → "other worship" (note first in a unit). אליל →
  "idol"; האלילה → "the she-idol"; אל זר → "a strange god"; אל אחר → "another
  god"; אלהים אחרים → "other gods".
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
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **אין לנו להחמיר** → "we may not be stringent".
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

