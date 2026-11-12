# Review pack: milhamot-hashem perek-2-07 (Milchamot ha-Shem 2:13–2:15)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-07`.

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
- **השי"ת** → "the Name, may He be blessed" (Q3). ית׳ / יתברך → "may He be
  blessed", where the Hebrew places it. ב"ה → "blessed be He"; השם הנכבד → "the glorious Name"; שם שמים
  → "the name of Heaven".
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
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
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
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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

