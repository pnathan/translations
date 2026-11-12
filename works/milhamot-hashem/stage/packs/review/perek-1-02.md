# Review pack: milhamot-hashem perek-1-02 (Milchamot ha-Shem 1:3–1:3)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-1-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-02`.

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


### Translation under review (1:3)

English: From the time the government opened the school, the children were studying Scripture and its translations, Onkelos and Rabbi Saadia Gaon, Mishnah and the Rambam, and grammar and the explanation of the words, All was done properly the study of the Torah, Prophets and Writings with clear explanation, a grammar book of the Holy Tongue! and the study of the Turkish language through the Turkish teacher, upright of heart, Chafetz Effendi. And after him Zia Bey with wonderful diligence! and we accustomed them to speak one with another in Hebrew and in Turkish, so that they would be well accustomed in their mouths, And at the turn of the year the examiners came from the government, may his majesty be exalted, and the children arranged their studies in Turkish and they showed them their writing, and the examiners saw their talents shining in their studies, and they praised them, that they are educated! Also the Hebrew examiners came on the day after and the children arranged their studies before them in Hebrew and in Arabic, and they found favour in their eyes for the ability of their studies, and praised them! And the following year the children did so, studying with diligence and alacrity, And at the end of the year the Turkish examiners came and examined them with uprightness of heart, and praised them, that they were making progress in their studies, and urged them with words and blessed them! But the Hebrew examiners who came after them had already begun to flare up among them the hatred of enlightenment and the study of languages. And they began to provoke them and to ask them about what they had not yet studied, and whoever could not answer anything anywhere they asked him, by the standard due to a complete sage, was counted by them as nothing at all. In order to diminish their honour, to sadden them and to slacken their hands from the studies, and they sought faults and blemishes in each and every one, and out of malice they did not turn to the studies and the knowledge that the children had acquired. Only to trip them up over what they had not yet studied, and to question them on the order of Nezikin, while they were still at the order of Moed, and when they showed them the beauty of their writing, the splendid and orderly Ashuri script, they sought a fault in the joining, as fine as a hair's thread, in some of the letters, as if they were writing S-T-V-M. And they did not look at the beauty of the script and its order. And in the interpretation of the Writings they asked them to solve difficult verses, such as "Binding his foal to the vine", and the like, and they went out from before them in low spirits, good, grieved, sad and with limp hands.

Literal: and-from-the-time when opened the-government (object-marker) house-of the-book were the-children studying Scripture and-its-translations Onkelos and-Rabbi-Saadia-Gaon Mishnah and-the-Rambam and-grammar and-the-explanation-of the-words all-of-it was-done on right study-of the-Torah Prophets and-Writings with clear-explanation a-book-of grammar language the-Holy and-the-study-of the-language Turkish through the-hand-of the-teacher the-Turkish upright of-heart Chafetz Effendi and-after-him Zia Bey with-diligence wonderful and-we-accustomed-them to-speak one with another in-the-language Hebrew and-Turkish in-order-that they-would-be accustomed in-their-mouths well and-at-the-turn-of the-year came the-examiners from the-government may-his-majesty-be-exalted and-arranged the-children (object-marker) their-studies in-the-language Turkish and-showed to-them their-writing and-saw the-examiners (object-marker) their-talents shining in-their-studies and-praised them that educated they also the-examiners the-Hebrew came on-the-day of the-next and-arranged the-children before-them (object-marker) their-studies in-Hebrew and-in-Arabic and-found favour in-their-eyes for-the-ability of-their-studies and-praised them and-in-the-year the-following so did the-children to-study with-diligence and-alacrity and-at-the-end-of the-year came the-examiners the-Turks and-examined them with-uprightness of-heart and-praised them that-they are-making progress in-their-studies and-urged them with-words and-blessed them but the-examiners the-Hebrew who-came after-them already had-begun to-flare-up among-them the-hatred-of enlightenment and-the-study-of languages and-they-began to-provoke and-to-ask them about-what still not they-had-studied and-whoever who-is-not able to-answer anything in-every place when-they-asked him by-the-standard-due-to a-sage complete not was-counted in-their-eyes for-nothing-at-all in-order to-diminish their-honour and-to-sadden-them and-to-slacken their-hands from the-studies and-they-sought faults and-blemishes for-each one and-one and-out-of-malice heart they-did-not turn to the-studies and-the-knowledge that-they-had-acquired the-children only to-trip-them-up over-what still not they-had-studied and-to-ask them on-the-order-of Nezikin and-still standing in-the-order-of Moed and-when-they-showed them (object-marker) the-beauty of-their-writing the-script Ashuri splendid and-orderly they-sought for-them a-fault in-the-joining as-fine-as-the-thread of-hair in-some letters as-if they were-writing S-T-V-M and-not they-looked at-the-beauty of-the-script and-its-order and-in-the-interpretation-of the-Writings they-asked of-them to-solve for-them verses difficult such-as binding to-the-vine his-foal and-the-like and-they-went-out from-before-them in-the-dejection-of soul good grieved and-sad and-limp hands

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:3.1 | and from the time |  |  |  |
| 1:3.2 | when |  |  | אשר is a relative here: 'when' (the time at which the government opened the school). |
| 1:3.3 | opened |  |  |  |
| 1:3.4 | the government |  |  |  |
| 1:3.5 | (object marker) |  |  | את before a definite object has no English word. |
| 1:3.6 | house of |  |  | בית הספר: 'house of the book', the word for school; the English gives 'the school'. |
| 1:3.7 | the book |  |  |  |
| 1:3.8 | were |  |  |  |
| 1:3.9 | the children |  |  |  |
| 1:3.10 | studying |  |  |  |
| 1:3.11 | Scripture |  |  |  |
| 1:3.12 | and its translations |  |  |  |
| 1:3.13 | Onkelos |  |  |  |
| 1:3.14 | and Rabbi Saadia Gaon |  |  | ורס״ג: Rabbi Saadia Gaon (רב סעדיה גאון), the abbreviation expanded as the decisions require. |
| 1:3.15 | Mishnah |  |  |  |
| 1:3.16 | and the Rambam |  |  | הרמב״ם: the Rambam, Maimonides. |
| 1:3.17 | and grammar |  |  |  |
| 1:3.18 | and the explanation of |  |  | ופי׳ is ופירוש, 'explanation'. |
| 1:3.19 | the words |  |  |  |
| 1:3.20 | all of it |  |  |  |
| 1:3.21 | was done |  |  |  |
| 1:3.22 | on |  |  |  |
| 1:3.23 | right |  |  |  |
| 1:3.24 | study of |  |  |  |
| 1:3.25 | the Torah |  |  |  |
| 1:3.26 | Prophets |  |  |  |
| 1:3.27 | and Writings |  |  |  |
| 1:3.28 | with |  |  | בפי׳ is בפירוש, 'with explanation'. |
| 1:3.29 | clear explanation |  |  |  |
| 1:3.30 | a book of |  |  |  |
| 1:3.31 | grammar |  |  |  |
| 1:3.32 | language |  |  |  |
| 1:3.33 | the Holy |  |  | לשון הקדש: the Hebrew language, 'the holy tongue'. |
| 1:3.34 | and the study of |  |  |  |
| 1:3.35 | the language |  |  |  |
| 1:3.36 | Turkish |  |  | התורקית: Turkish, the language of the government. |
| 1:3.37 | through |  |  |  |
| 1:3.38 | the hand of |  |  |  |
| 1:3.39 | the teacher |  |  |  |
| 1:3.40 | the Turkish |  |  |  |
| 1:3.41 | upright |  |  |  |
| 1:3.42 | of heart |  |  |  |
| 1:3.43 | Chafetz |  |  | The name is given as the text spells it (חאפץ), in the project's scheme as Chafetz; no identification is made. |
| 1:3.44 | Effendi |  |  | אפנדי: the Turkish title Effendi. |
| 1:3.45 | and after him |  |  |  |
| 1:3.46 | Zia |  |  | זיא: the name as the text spells it; no identification is made. |
| 1:3.47 | Bey |  |  | ביי: the Turkish title Bey. |
| 1:3.48 | with diligence |  |  |  |
| 1:3.49 | wonderful |  |  |  |
| 1:3.50 | and we accustomed them |  |  | הרגלנום: a hif'il perfect, first person plural, with the object suffix ם: 'we accustomed them'. |
| 1:3.51 | to speak |  |  |  |
| 1:3.52 | one |  |  |  |
| 1:3.53 | with |  |  |  |
| 1:3.54 | another |  |  |  |
| 1:3.55 | in the language |  |  |  |
| 1:3.56 | Hebrew |  |  |  |
| 1:3.57 | and Turkish |  |  |  |
| 1:3.58 | in order that |  |  |  |
| 1:3.59 | they would be |  |  |  |
| 1:3.60 | accustomed |  |  |  |
| 1:3.61 | in their mouths |  |  |  |
| 1:3.62 | well |  |  |  |
| 1:3.63 | and at the turn of |  |  |  |
| 1:3.64 | the year |  |  |  |
| 1:3.65 | came |  |  |  |
| 1:3.66 | the examiners |  |  |  |
| 1:3.67 | from |  |  |  |
| 1:3.68 | the government |  |  |  |
| 1:3.69 | may his majesty be exalted |  |  | יר״ה: the honorific of the government, rendered as the decisions give it. |
| 1:3.70 | and arranged |  |  |  |
| 1:3.71 | the children |  |  |  |
| 1:3.72 | (object marker) |  |  |  |
| 1:3.73 | their studies |  |  |  |
| 1:3.74 | in the language |  |  |  |
| 1:3.75 | Turkish |  |  |  |
| 1:3.76 | and showed |  |  |  |
| 1:3.77 | to them |  |  |  |
| 1:3.78 | their writing |  |  |  |
| 1:3.79 | and saw |  |  |  |
| 1:3.80 | the examiners |  |  |  |
| 1:3.81 | (object marker) |  |  |  |
| 1:3.82 | their talents |  |  |  |
| 1:3.83 | shining |  |  |  |
| 1:3.84 | in their studies |  |  |  |
| 1:3.85 | and praised |  |  |  |
| 1:3.86 | them |  |  |  |
| 1:3.87 | that |  |  |  |
| 1:3.88 | educated |  |  | משכילים: 'educated' here; the same word is the name the Haskalah gave its adherents, and the English does not settle that sense. |
| 1:3.89 | they |  |  |  |
| 1:3.90 | also |  |  |  |
| 1:3.91 | the examiners |  |  |  |
| 1:3.92 | the Hebrew |  |  | הבוחנים העברים: the Hebrew examiners; the text does not say who appointed them. |
| 1:3.93 | came |  |  |  |
| 1:3.94 | on the day |  |  |  |
| 1:3.95 | of |  |  |  |
| 1:3.96 | the next |  |  |  |
| 1:3.97 | and arranged |  |  |  |
| 1:3.98 | the children |  |  |  |
| 1:3.99 | before them |  |  |  |
| 1:3.100 | (object marker) |  |  |  |
| 1:3.101 | their studies |  |  |  |
| 1:3.102 | in Hebrew |  |  |  |
| 1:3.103 | and in Arabic |  |  |  |
| 1:3.104 | and found |  |  |  |
| 1:3.105 | favour |  |  |  |
| 1:3.106 | in their eyes |  |  |  |
| 1:3.107 | for the ability |  |  |  |
| 1:3.108 | of their studies |  |  |  |
| 1:3.109 | and praised |  |  |  |
| 1:3.110 | them |  |  |  |
| 1:3.111 | and in the year |  |  |  |
| 1:3.112 | the following |  |  |  |
| 1:3.113 | so |  |  |  |
| 1:3.114 | did |  |  |  |
| 1:3.115 | the children |  |  |  |
| 1:3.116 | to study |  |  |  |
| 1:3.117 | with diligence |  |  |  |
| 1:3.118 | and alacrity |  |  |  |
| 1:3.119 | and at the end of |  |  |  |
| 1:3.120 | the year |  |  |  |
| 1:3.121 | came |  |  |  |
| 1:3.122 | the examiners |  |  |  |
| 1:3.123 | the Turks |  |  | התורקים: the Turks, the examiners of the government. |
| 1:3.124 | and examined |  |  |  |
| 1:3.125 | them |  |  |  |
| 1:3.126 | with uprightness |  |  |  |
| 1:3.127 | of heart |  |  |  |
| 1:3.128 | and praised |  |  |  |
| 1:3.129 | them |  |  |  |
| 1:3.130 | that they |  |  |  |
| 1:3.131 | are making |  |  |  |
| 1:3.132 | progress |  |  | חיל: literally 'strength'; in עושים חיל בלימודים it means 'making progress in the studies'. |
| 1:3.133 | in their studies |  |  |  |
| 1:3.134 | and urged |  |  |  |
| 1:3.135 | them |  |  |  |
| 1:3.136 | with words |  |  |  |
| 1:3.137 | and blessed |  |  |  |
| 1:3.138 | them |  |  |  |
| 1:3.139 | but |  |  |  |
| 1:3.140 | the examiners |  |  |  |
| 1:3.141 | the Hebrew |  |  |  |
| 1:3.142 | who came |  |  |  |
| 1:3.143 | after them |  |  |  |
| 1:3.144 | already |  |  |  |
| 1:3.145 | had begun |  |  |  |
| 1:3.146 | to flare up |  |  |  |
| 1:3.147 | among them |  |  |  |
| 1:3.148 | the hatred of |  |  |  |
| 1:3.149 | enlightenment |  |  | ההשכלה: enlightenment, the movement called the Haskalah. |
| 1:3.150 | and the study of |  |  |  |
| 1:3.151 | languages |  |  |  |
| 1:3.152 | and they began |  |  |  |
| 1:3.153 | to provoke |  |  | לקנתרם: to provoke, to vex. |
| 1:3.154 | and to ask |  |  |  |
| 1:3.155 | them |  |  |  |
| 1:3.156 | about what |  |  |  |
| 1:3.157 | still |  |  |  |
| 1:3.158 | not |  |  |  |
| 1:3.159 | they had studied |  |  |  |
| 1:3.160 | and whoever |  |  |  |
| 1:3.161 | who is not |  |  |  |
| 1:3.162 | able |  |  |  |
| 1:3.163 | to answer |  |  |  |
| 1:3.164 | anything |  |  |  |
| 1:3.165 | in every |  |  |  |
| 1:3.166 | place |  |  |  |
| 1:3.167 | when they asked |  |  |  |
| 1:3.168 | him |  |  |  |
| 1:3.169 | by the standard due to |  |  |  |
| 1:3.170 | a sage |  |  |  |
| 1:3.171 | complete |  |  |  |
| 1:3.172 | not |  |  |  |
| 1:3.173 | was counted |  |  |  |
| 1:3.174 | in their eyes |  |  |  |
| 1:3.175 | for nothing at all |  |  |  |
| 1:3.176 | in order |  |  |  |
| 1:3.177 | to diminish |  |  |  |
| 1:3.178 | their honour |  |  |  |
| 1:3.179 | and to sadden them |  |  |  |
| 1:3.180 | and to slacken |  |  |  |
| 1:3.181 | their hands |  |  |  |
| 1:3.182 | from |  |  |  |
| 1:3.183 | the studies |  |  |  |
| 1:3.184 | and they sought |  |  |  |
| 1:3.185 | faults |  |  |  |
| 1:3.186 | and blemishes |  |  |  |
| 1:3.187 | for each |  |  |  |
| 1:3.188 | one |  |  |  |
| 1:3.189 | and one |  |  |  |
| 1:3.190 | and out of malice |  |  |  |
| 1:3.191 | heart |  |  |  |
| 1:3.192 | they did not |  |  |  |
| 1:3.193 | turn |  |  |  |
| 1:3.194 | to |  |  |  |
| 1:3.195 | the studies |  |  | The witness spells הלמודים; it is read as הלימודים (see apparatus). |
| 1:3.196 | and the knowledge |  |  |  |
| 1:3.197 | that they had acquired |  |  |  |
| 1:3.198 | the children |  |  |  |
| 1:3.199 | only |  |  |  |
| 1:3.200 | to trip them up |  |  | לקפחם: to put them down, to trip them up; literally 'to diminish them'. |
| 1:3.201 | over what |  |  |  |
| 1:3.202 | still |  |  |  |
| 1:3.203 | not |  |  |  |
| 1:3.204 | they had studied |  |  |  |
| 1:3.205 | and to ask |  |  |  |
| 1:3.206 | them |  |  |  |
| 1:3.207 | on the order of |  |  |  |
| 1:3.208 | Nezikin |  |  | נזיקים: Nezikin, the order of the Mishnah on damages. |
| 1:3.209 | and still |  |  |  |
| 1:3.210 | standing |  |  |  |
| 1:3.211 | in the order of |  |  |  |
| 1:3.212 | Moed |  |  | מועד: Moed, the order of the Mishnah on festivals. |
| 1:3.213 | and when they showed |  |  |  |
| 1:3.214 | them |  |  |  |
| 1:3.215 | (object marker) |  |  |  |
| 1:3.216 | the beauty |  |  |  |
| 1:3.217 | of their writing |  |  |  |
| 1:3.218 | the script |  |  | כתב אשורי: the Ashuri script, the square Hebrew letters. |
| 1:3.219 | Ashuri |  |  |  |
| 1:3.220 | splendid |  |  |  |
| 1:3.221 | and orderly |  |  |  |
| 1:3.222 | they sought |  |  |  |
| 1:3.223 | for them |  |  |  |
| 1:3.224 | a fault |  |  |  |
| 1:3.225 | in the joining |  |  | דיבוק: the joining of letters; the fault sought is in the joining. |
| 1:3.226 | as fine as the thread |  |  |  |
| 1:3.227 | of hair |  |  |  |
| 1:3.228 | in some |  |  |  |
| 1:3.229 | letters |  |  |  |
| 1:3.230 | as if |  |  |  |
| 1:3.231 | they |  |  |  |
| 1:3.232 | were writing |  |  |  |
| 1:3.233 | S-T-V-M |  |  | סתו״מ is an abbreviation not expanded in decisions.md; it is kept as its letters, and its expansion is open. |
| 1:3.234 | and not |  |  |  |
| 1:3.235 | they looked |  |  |  |
| 1:3.236 | at the beauty |  |  |  |
| 1:3.237 | of the script |  |  |  |
| 1:3.238 | and its order |  |  |  |
| 1:3.239 | and in the interpretation of |  |  |  |
| 1:3.240 | the Writings |  |  |  |
| 1:3.241 | they asked |  |  |  |
| 1:3.242 | of them |  |  |  |
| 1:3.243 | to solve |  |  |  |
| 1:3.244 | for them |  |  |  |
| 1:3.245 | verses |  |  |  |
| 1:3.246 | difficult |  |  |  |
| 1:3.247 | such as |  |  |  |
| 1:3.248 | binding |  |  | אסרי לגפן עירה: Gen 49:11, quoted. |
| 1:3.249 | to the vine |  |  |  |
| 1:3.250 | his foal |  |  |  |
| 1:3.251 | and the like |  |  |  |
| 1:3.252 | and they went out |  |  |  |
| 1:3.253 | from before them |  |  |  |
| 1:3.254 | in the dejection of |  |  |  |
| 1:3.255 | soul |  |  |  |
| 1:3.256 | good |  |  |  |
| 1:3.257 | grieved |  |  |  |
| 1:3.258 | and sad |  |  |  |
| 1:3.259 | and limp |  |  |  |
| 1:3.260 | hands |  |  |  |

Phrases:
- 1:3.1 1:3.2 1:3.3 1:3.4 1:3.5 1:3.6 1:3.7 | From the time the government opened the school, | and-from-the-time when opened the-government (object-marker) house-of the-book
- 1:3.8 1:3.9 1:3.10 1:3.11 1:3.12 1:3.13 1:3.14 | the children were studying Scripture and its translations, Onkelos and Rabbi Saadia Gaon, | were the-children studying Scripture and-its-translations Onkelos and-Rabbi-Saadia-Gaon
- 1:3.15 1:3.16 | Mishnah and the Rambam, | Mishnah and-the-Rambam
- 1:3.17 1:3.18 1:3.19 | and grammar and the explanation of the words, | and-grammar and-the-explanation-of the-words
- 1:3.20 1:3.21 1:3.22 1:3.23 | All was done properly | all-of-it was-done on right
- 1:3.24 1:3.25 1:3.26 1:3.27 1:3.28 1:3.29 | the study of the Torah, Prophets and Writings with clear explanation, | study-of the-Torah Prophets and-Writings with clear-explanation
- 1:3.30 1:3.31 1:3.32 1:3.33 | a grammar book of the Holy Tongue! | a-book-of grammar language the-Holy
- 1:3.34 1:3.35 1:3.36 | and the study of the Turkish language | and-the-study-of the-language Turkish
- 1:3.37 1:3.38 1:3.39 1:3.40 1:3.41 1:3.42 1:3.43 1:3.44 | through the Turkish teacher, upright of heart, Chafetz Effendi. | through the-hand-of the-teacher the-Turkish upright of-heart Chafetz Effendi
- 1:3.45 1:3.46 1:3.47 1:3.48 1:3.49 | And after him Zia Bey with wonderful diligence! | and-after-him Zia Bey with-diligence wonderful
- 1:3.50 1:3.51 1:3.52 1:3.53 1:3.54 1:3.55 1:3.56 1:3.57 | and we accustomed them to speak one with another in Hebrew and in Turkish, | and-we-accustomed-them to-speak one with another in-the-language Hebrew and-Turkish
- 1:3.58 1:3.59 1:3.60 1:3.61 1:3.62 | so that they would be well accustomed in their mouths, | in-order-that they-would-be accustomed in-their-mouths well
- 1:3.63 1:3.64 1:3.65 1:3.66 | And at the turn of the year the examiners came | and-at-the-turn-of the-year came the-examiners
- 1:3.67 1:3.68 1:3.69 | from the government, may his majesty be exalted, | from the-government may-his-majesty-be-exalted
- 1:3.70 1:3.71 1:3.72 1:3.73 1:3.74 1:3.75 | and the children arranged their studies in Turkish | and-arranged the-children (object-marker) their-studies in-the-language Turkish
- 1:3.76 1:3.77 1:3.78 | and they showed them their writing, | and-showed to-them their-writing
- 1:3.79 1:3.80 1:3.81 1:3.82 1:3.83 1:3.84 | and the examiners saw their talents shining in their studies, | and-saw the-examiners (object-marker) their-talents shining in-their-studies
- 1:3.85 1:3.86 1:3.87 1:3.88 1:3.89 | and they praised them, that they are educated! | and-praised them that educated they
- 1:3.90 1:3.91 1:3.92 1:3.93 | Also the Hebrew examiners came | also the-examiners the-Hebrew came
- 1:3.94 1:3.95 1:3.96 | on the day after | on-the-day of the-next
- 1:3.97 1:3.98 1:3.99 1:3.100 1:3.101 1:3.102 1:3.103 | and the children arranged their studies before them in Hebrew and in Arabic, | and-arranged the-children before-them (object-marker) their-studies in-Hebrew and-in-Arabic
- 1:3.104 1:3.105 1:3.106 1:3.107 1:3.108 1:3.109 1:3.110 | and they found favour in their eyes for the ability of their studies, and praised them! | and-found favour in-their-eyes for-the-ability of-their-studies and-praised them
- 1:3.111 1:3.112 1:3.113 1:3.114 1:3.115 1:3.116 1:3.117 1:3.118 | And the following year the children did so, studying with diligence and alacrity, | and-in-the-year the-following so did the-children to-study with-diligence and-alacrity
- 1:3.119 1:3.120 1:3.121 1:3.122 1:3.123 1:3.124 1:3.125 1:3.126 1:3.127 | And at the end of the year the Turkish examiners came and examined them with uprightness of heart, | and-at-the-end-of the-year came the-examiners the-Turks and-examined them with-uprightness of-heart
- 1:3.128 1:3.129 1:3.130 1:3.131 1:3.132 1:3.133 | and praised them, that they were making progress in their studies, | and-praised them that-they are-making progress in-their-studies
- 1:3.134 1:3.135 1:3.136 1:3.137 1:3.138 | and urged them with words and blessed them! | and-urged them with-words and-blessed them
- 1:3.139 1:3.140 1:3.141 1:3.142 1:3.143 | But the Hebrew examiners who came after them | but the-examiners the-Hebrew who-came after-them
- 1:3.144 1:3.145 1:3.146 1:3.147 1:3.148 1:3.149 1:3.150 1:3.151 | had already begun to flare up among them the hatred of enlightenment and the study of languages. | already had-begun to-flare-up among-them the-hatred-of enlightenment and-the-study-of languages
- 1:3.152 1:3.153 1:3.154 1:3.155 1:3.156 1:3.157 1:3.158 1:3.159 | And they began to provoke them and to ask them about what they had not yet studied, | and-they-began to-provoke and-to-ask them about-what still not they-had-studied
- 1:3.160 1:3.161 1:3.162 1:3.163 1:3.164 1:3.165 1:3.166 1:3.167 1:3.168 1:3.169 1:3.170 1:3.171 1:3.172 1:3.173 1:3.174 1:3.175 | and whoever could not answer anything anywhere they asked him, by the standard due to a complete sage, was counted by them as nothing at all. | and-whoever who-is-not able to-answer anything in-every place when-they-asked him by-the-standard-due-to a-sage complete not was-counted in-their-eyes for-nothing-at-all
- 1:3.176 1:3.177 1:3.178 1:3.179 1:3.180 1:3.181 1:3.182 1:3.183 | In order to diminish their honour, to sadden them and to slacken their hands from the studies, | in-order to-diminish their-honour and-to-sadden-them and-to-slacken their-hands from the-studies
- 1:3.184 1:3.185 1:3.186 1:3.187 1:3.188 1:3.189 | and they sought faults and blemishes in each and every one, | and-they-sought faults and-blemishes for-each one and-one
- 1:3.190 1:3.191 1:3.192 1:3.193 1:3.194 1:3.195 1:3.196 1:3.197 1:3.198 | and out of malice they did not turn to the studies and the knowledge that the children had acquired. | and-out-of-malice heart they-did-not turn to the-studies and-the-knowledge that-they-had-acquired the-children
- 1:3.199 1:3.200 1:3.201 1:3.202 1:3.203 1:3.204 1:3.205 1:3.206 1:3.207 1:3.208 1:3.209 1:3.210 1:3.211 1:3.212 | Only to trip them up over what they had not yet studied, and to question them on the order of Nezikin, while they were still at the order of Moed, | only to-trip-them-up over-what still not they-had-studied and-to-ask them on-the-order-of Nezikin and-still standing in-the-order-of Moed
- 1:3.213 1:3.214 1:3.215 1:3.216 1:3.217 1:3.218 1:3.219 1:3.220 1:3.221 | and when they showed them the beauty of their writing, the splendid and orderly Ashuri script, | and-when-they-showed them (object-marker) the-beauty of-their-writing the-script Ashuri splendid and-orderly
- 1:3.222 1:3.223 1:3.224 1:3.225 1:3.226 1:3.227 1:3.228 1:3.229 1:3.230 1:3.231 1:3.232 1:3.233 | they sought a fault in the joining, as fine as a hair's thread, in some of the letters, as if they were writing S-T-V-M. | they-sought for-them a-fault in-the-joining as-fine-as-the-thread of-hair in-some letters as-if they were-writing S-T-V-M
- 1:3.234 1:3.235 1:3.236 1:3.237 1:3.238 | And they did not look at the beauty of the script and its order. | and-not they-looked at-the-beauty of-the-script and-its-order
- 1:3.239 1:3.240 1:3.241 1:3.242 1:3.243 1:3.244 1:3.245 1:3.246 1:3.247 1:3.248 1:3.249 1:3.250 | And in the interpretation of the Writings they asked them to solve difficult verses, such as "Binding his foal to the vine", | and-in-the-interpretation-of the-Writings they-asked of-them to-solve for-them verses difficult such-as binding to-the-vine his-foal
- 1:3.251 1:3.252 1:3.253 1:3.254 1:3.255 1:3.256 1:3.257 1:3.258 1:3.259 1:3.260 | and the like, and they went out from before them in low spirits, good, grieved, sad and with limp hands. | and-the-like and-they-went-out from-before-them in-the-dejection-of soul good grieved and-sad and-limp hands

Grammar notes:
- The paragraph is a chain of clauses joined by ו and by the time phrases that open each school year: ולתקופת השנה (3:63), ולשנה האחרת (3:111), ובצאת השנה (3:119). The sentence that begins with ומעת has its main verb in היו הילדים לומדים (3:8).
- והרגלנום לדבר זה עם זה (3:50) is a hif'il perfect with an object suffix; the purpose clause כדי שיהיו מורגלים בפיהם היטב (3:58) follows it.
- The sentence ואחריו זיא ביי בחריצות נפלאה! (3:45-3:49) has no verb, and the English adds none beyond the words.
- The turn to hostility comes at אך (3:139): the Hebrew examiners who came after them are set against the Turkish examiners, and the verb התחילה (3:145) starts the change. The purpose clause כדי להפחית כבודם (3:176) states the aim.
- The phrase אסרי לגפן עירה (3:248-3:250) is quoted as an example of a hard verse, so the English keeps quotation marks; the source is Gen 49:11.

Translator's note: Translated from the Wikisource witness as base text. The paragraph is ironic in the author's framing: the examiners first praise the children, then the Hebrew examiners set out to shame them. משכילים and ההשכלה are rendered in neutral terms; see the token notes.

Commentary rulings:
- apparatus: Witness spelling הלמודים (1:3.195) is read as the evident word הלימודים.
- apparatus: Bible verse: Gen 49:11, 'Binding his foal to the vine' (אסרי לגפן עירה), quoted as an example of a hard verse (1:3.248-1:3.250).
- apparatus: The abbreviation סתו״מ (1:3.233) is not expanded in decisions.md; it is kept as the letters S-T-V-M.
- apparatus: Base text: the Wikisource witness, not yet collated against the 1931 scan (open question 2).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

