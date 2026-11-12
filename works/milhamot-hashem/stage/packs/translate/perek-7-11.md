# Translation pack: milhamot-hashem perek-7-11 (Milchamot ha-Shem 7:15–7:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-11`.

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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
- **קבלה / המקובלים / בעלי הקבלה / חכמי הקבלה**: the lore → "Kabbalah"; הקבלה
  החדשה → "the new Kabbalah" (never "later" or "Lurianic"). המקובלים → "the
  Kabbalists"; חכמי הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" →
  "our rabbis, the masters of the new Kabbalah" (courtesy kept). His claim on the word for the rabbinic chain (Q5): "בעלי
  הקבלה האמתית" → "the masters of the true tradition", a note each time that
  the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" → "it is not a
  tradition from Moses". The participle: "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
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
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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


## Milchamot ha-Shem 7:15

Hebrew (Wikisource): פח) העיקר הרביעי הקדמות. צריך להאמין כי זה האחד האו קדמון במלה מוחלטת וכל נמצא זולתו אינו קדמון כשנעריכהו אליו. והראיות על זה העיקר רבו בספרים, וזה העיקר הרביעי נזכר בתורה מעונה אלהי קדם ע"כ. כבר ביארתי לעיל סי' מ"ח שגם זה העיקר מיאנו בו כת המקובלים החדשים כמו שנתבאר בענין פיוט יגדל אלהים חי שמיאנו המקובלים לאמרו, ואל תשמי לדברי האומרים שנמנעו מלאמרו מטעם שאין לעשות לתורה עיקרים לפי שכל מצוה ומצוה של תורה היא עיקר לעצמה וכולה עיקרים, כי לא מחכמה אמרו ובאין מבין בעיקרים דברו כן שהרי זה נכלל בעיקר השמיני, ומובאר בהדיא שם שכל התורה כולב מפי הגבורה, ומשה רבינו כמו סופר, הקב"ה אומר, והוא אומר וכותב כל מה שאמר לו, מספר שני הדורות שעברו, וספוריה, ומצותיה, ולפיכך נקרא משה רבינו מחוקק, ואין הפרש בין בני חם וכוש ומצרים וכו' ושם אשתו מהטבאל בת מטרד ותמנע היתה פילגש, ובין אנכי ה' אלהיך, ושמע ישראל וכו' הכל מפי הגבורה, הכל תורת ה' תמימה טהורה קדושה אמת, ורז"ך היה אצלם משנה כופר ומכחש יותר מכל מין וכופר על שחשב שיש בתורה מוח וקליפה, ושחשבו השנים שבה והספורים אין בהם תועלת, ושהם מפי משה עצמו, והוא ענין אין תורה מן השמים, וכך אמרו חכמים, האומר כל שתורה כולה מפי הגבורה חוץ מפסוק זה שלא אמרו הקב"ה אלא משה מפי עצמו, זהו דבר ה' בזה אלא כל ענייניה ומצותיה מפי הגבורה, ומהם ישיג מי שזכהו השי"ת עניינים נפלעים, כמאמר דוד המלך ע"ה כל עיני ואביטה נפלאות מתורתיך והלכות סוכה ולולב ושופר וציצית ותפלין עם כל חילוקי דיניהם הם וזולתם, הם בעצמם ההלכות שקיבל משה רבי' ע"ה מהקב"ה בסיני. והוא אמרם לנו. והוא נאמן בשליחותו. והמאמר המורה על זה הוא מה שנאמר בזאת תדעון כי ה' שלחני לעשות את כל המעשים האלה כי לא מלבי ע"כ. הא למדת שטענתם זו כוזבת. שהרי רבינו מודה בזה וכלל כל מצותיה של תורה וספוריה בזה העיקר השמיני באמרו שכל מצות התורה וחילוקי דיניהם וסיפוריה כולם מפי הגבורה. וכל הכופר בשום חלק מכל מצות התורה או אפי' מספוריה הרי הוא מין וכופר. וכולם נכללים תחת סוג תורת אמת נתן לעמו אל. על יד נביאו, ואין דבר בה לבטלה. וכל מי שאינו מודה באחת מהן הוא מין וכופר ואפיקירוס כמנשה וסיעתו. והראיה הגדולה על מה שאמרתי שאינם מודים בארבעה עיקרים הראשונים שהרי ביאר הרב חמדת ימים שלא חרה להאר"י רק על ארבעה בתים ראשונים שביגדל שהם סותרים אמונתם. אבל מהנו אדון עולם ואילך שפיר דמי לדידהו לאמרו כנז' לעיל סי' מ"ח.

| id | word | punctuation after |
|---|---|---|
| 7:15.1 | פח | ) |
| 7:15.2 | העיקר |  |
| 7:15.3 | הרביעי |  |
| 7:15.4 | הקדמות | . |
| 7:15.5 | צריך |  |
| 7:15.6 | להאמין |  |
| 7:15.7 | כי |  |
| 7:15.8 | זה |  |
| 7:15.9 | האחד |  |
| 7:15.10 | האו |  |
| 7:15.11 | קדמון |  |
| 7:15.12 | במלה |  |
| 7:15.13 | מוחלטת |  |
| 7:15.14 | וכל |  |
| 7:15.15 | נמצא |  |
| 7:15.16 | זולתו |  |
| 7:15.17 | אינו |  |
| 7:15.18 | קדמון |  |
| 7:15.19 | כשנעריכהו |  |
| 7:15.20 | אליו | . |
| 7:15.21 | והראיות |  |
| 7:15.22 | על |  |
| 7:15.23 | זה |  |
| 7:15.24 | העיקר |  |
| 7:15.25 | רבו |  |
| 7:15.26 | בספרים | , |
| 7:15.27 | וזה |  |
| 7:15.28 | העיקר |  |
| 7:15.29 | הרביעי |  |
| 7:15.30 | נזכר |  |
| 7:15.31 | בתורה |  |
| 7:15.32 | מעונה |  |
| 7:15.33 | אלהי |  |
| 7:15.34 | קדם |  |
| 7:15.35 | ע"כ | . |
| 7:15.36 | כבר |  |
| 7:15.37 | ביארתי |  |
| 7:15.38 | לעיל |  |
| 7:15.39 | סי | ' |
| 7:15.40 | מ"ח |  |
| 7:15.41 | שגם |  |
| 7:15.42 | זה |  |
| 7:15.43 | העיקר |  |
| 7:15.44 | מיאנו |  |
| 7:15.45 | בו |  |
| 7:15.46 | כת |  |
| 7:15.47 | המקובלים |  |
| 7:15.48 | החדשים |  |
| 7:15.49 | כמו |  |
| 7:15.50 | שנתבאר |  |
| 7:15.51 | בענין |  |
| 7:15.52 | פיוט |  |
| 7:15.53 | יגדל |  |
| 7:15.54 | אלהים |  |
| 7:15.55 | חי |  |
| 7:15.56 | שמיאנו |  |
| 7:15.57 | המקובלים |  |
| 7:15.58 | לאמרו | , |
| 7:15.59 | ואל |  |
| 7:15.60 | תשמי |  |
| 7:15.61 | לדברי |  |
| 7:15.62 | האומרים |  |
| 7:15.63 | שנמנעו |  |
| 7:15.64 | מלאמרו |  |
| 7:15.65 | מטעם |  |
| 7:15.66 | שאין |  |
| 7:15.67 | לעשות |  |
| 7:15.68 | לתורה |  |
| 7:15.69 | עיקרים |  |
| 7:15.70 | לפי |  |
| 7:15.71 | שכל |  |
| 7:15.72 | מצוה |  |
| 7:15.73 | ומצוה |  |
| 7:15.74 | של |  |
| 7:15.75 | תורה |  |
| 7:15.76 | היא |  |
| 7:15.77 | עיקר |  |
| 7:15.78 | לעצמה |  |
| 7:15.79 | וכולה |  |
| 7:15.80 | עיקרים | , |
| 7:15.81 | כי |  |
| 7:15.82 | לא |  |
| 7:15.83 | מחכמה |  |
| 7:15.84 | אמרו |  |
| 7:15.85 | ובאין |  |
| 7:15.86 | מבין |  |
| 7:15.87 | בעיקרים |  |
| 7:15.88 | דברו |  |
| 7:15.89 | כן |  |
| 7:15.90 | שהרי |  |
| 7:15.91 | זה |  |
| 7:15.92 | נכלל |  |
| 7:15.93 | בעיקר |  |
| 7:15.94 | השמיני | , |
| 7:15.95 | ומובאר |  |
| 7:15.96 | בהדיא |  |
| 7:15.97 | שם |  |
| 7:15.98 | שכל |  |
| 7:15.99 | התורה |  |
| 7:15.100 | כולב |  |
| 7:15.101 | מפי |  |
| 7:15.102 | הגבורה | , |
| 7:15.103 | ומשה |  |
| 7:15.104 | רבינו |  |
| 7:15.105 | כמו |  |
| 7:15.106 | סופר | , |
| 7:15.107 | הקב"ה |  |
| 7:15.108 | אומר | , |
| 7:15.109 | והוא |  |
| 7:15.110 | אומר |  |
| 7:15.111 | וכותב |  |
| 7:15.112 | כל |  |
| 7:15.113 | מה |  |
| 7:15.114 | שאמר |  |
| 7:15.115 | לו | , |
| 7:15.116 | מספר |  |
| 7:15.117 | שני |  |
| 7:15.118 | הדורות |  |
| 7:15.119 | שעברו | , |
| 7:15.120 | וספוריה | , |
| 7:15.121 | ומצותיה | , |
| 7:15.122 | ולפיכך |  |
| 7:15.123 | נקרא |  |
| 7:15.124 | משה |  |
| 7:15.125 | רבינו |  |
| 7:15.126 | מחוקק | , |
| 7:15.127 | ואין |  |
| 7:15.128 | הפרש |  |
| 7:15.129 | בין |  |
| 7:15.130 | בני |  |
| 7:15.131 | חם |  |
| 7:15.132 | וכוש |  |
| 7:15.133 | ומצרים |  |
| 7:15.134 | וכו | ' |
| 7:15.135 | ושם |  |
| 7:15.136 | אשתו |  |
| 7:15.137 | מהטבאל |  |
| 7:15.138 | בת |  |
| 7:15.139 | מטרד |  |
| 7:15.140 | ותמנע |  |
| 7:15.141 | היתה |  |
| 7:15.142 | פילגש | , |
| 7:15.143 | ובין |  |
| 7:15.144 | אנכי |  |
| 7:15.145 | ה | ' |
| 7:15.146 | אלהיך | , |
| 7:15.147 | ושמע |  |
| 7:15.148 | ישראל |  |
| 7:15.149 | וכו | ' |
| 7:15.150 | הכל |  |
| 7:15.151 | מפי |  |
| 7:15.152 | הגבורה | , |
| 7:15.153 | הכל |  |
| 7:15.154 | תורת |  |
| 7:15.155 | ה | ' |
| 7:15.156 | תמימה |  |
| 7:15.157 | טהורה |  |
| 7:15.158 | קדושה |  |
| 7:15.159 | אמת | , |
| 7:15.160 | ורז"ך |  |
| 7:15.161 | היה |  |
| 7:15.162 | אצלם |  |
| 7:15.163 | משנה |  |
| 7:15.164 | כופר |  |
| 7:15.165 | ומכחש |  |
| 7:15.166 | יותר |  |
| 7:15.167 | מכל |  |
| 7:15.168 | מין |  |
| 7:15.169 | וכופר |  |
| 7:15.170 | על |  |
| 7:15.171 | שחשב |  |
| 7:15.172 | שיש |  |
| 7:15.173 | בתורה |  |
| 7:15.174 | מוח |  |
| 7:15.175 | וקליפה | , |
| 7:15.176 | ושחשבו |  |
| 7:15.177 | השנים |  |
| 7:15.178 | שבה |  |
| 7:15.179 | והספורים |  |
| 7:15.180 | אין |  |
| 7:15.181 | בהם |  |
| 7:15.182 | תועלת | , |
| 7:15.183 | ושהם |  |
| 7:15.184 | מפי |  |
| 7:15.185 | משה |  |
| 7:15.186 | עצמו | , |
| 7:15.187 | והוא |  |
| 7:15.188 | ענין |  |
| 7:15.189 | אין |  |
| 7:15.190 | תורה |  |
| 7:15.191 | מן |  |
| 7:15.192 | השמים | , |
| 7:15.193 | וכך |  |
| 7:15.194 | אמרו |  |
| 7:15.195 | חכמים | , |
| 7:15.196 | האומר |  |
| 7:15.197 | כל |  |
| 7:15.198 | שתורה |  |
| 7:15.199 | כולה |  |
| 7:15.200 | מפי |  |
| 7:15.201 | הגבורה |  |
| 7:15.202 | חוץ |  |
| 7:15.203 | מפסוק |  |
| 7:15.204 | זה |  |
| 7:15.205 | שלא |  |
| 7:15.206 | אמרו |  |
| 7:15.207 | הקב"ה |  |
| 7:15.208 | אלא |  |
| 7:15.209 | משה |  |
| 7:15.210 | מפי |  |
| 7:15.211 | עצמו | , |
| 7:15.212 | זהו |  |
| 7:15.213 | דבר |  |
| 7:15.214 | ה | ' |
| 7:15.215 | בזה |  |
| 7:15.216 | אלא |  |
| 7:15.217 | כל |  |
| 7:15.218 | ענייניה |  |
| 7:15.219 | ומצותיה |  |
| 7:15.220 | מפי |  |
| 7:15.221 | הגבורה | , |
| 7:15.222 | ומהם |  |
| 7:15.223 | ישיג |  |
| 7:15.224 | מי |  |
| 7:15.225 | שזכהו |  |
| 7:15.226 | השי"ת |  |
| 7:15.227 | עניינים |  |
| 7:15.228 | נפלעים | , |
| 7:15.229 | כמאמר |  |
| 7:15.230 | דוד |  |
| 7:15.231 | המלך |  |
| 7:15.232 | ע"ה |  |
| 7:15.233 | כל |  |
| 7:15.234 | עיני |  |
| 7:15.235 | ואביטה |  |
| 7:15.236 | נפלאות |  |
| 7:15.237 | מתורתיך |  |
| 7:15.238 | והלכות |  |
| 7:15.239 | סוכה |  |
| 7:15.240 | ולולב |  |
| 7:15.241 | ושופר |  |
| 7:15.242 | וציצית |  |
| 7:15.243 | ותפלין |  |
| 7:15.244 | עם |  |
| 7:15.245 | כל |  |
| 7:15.246 | חילוקי |  |
| 7:15.247 | דיניהם |  |
| 7:15.248 | הם |  |
| 7:15.249 | וזולתם | , |
| 7:15.250 | הם |  |
| 7:15.251 | בעצמם |  |
| 7:15.252 | ההלכות |  |
| 7:15.253 | שקיבל |  |
| 7:15.254 | משה |  |
| 7:15.255 | רבי | ' |
| 7:15.256 | ע"ה |  |
| 7:15.257 | מהקב"ה |  |
| 7:15.258 | בסיני | . |
| 7:15.259 | והוא |  |
| 7:15.260 | אמרם |  |
| 7:15.261 | לנו | . |
| 7:15.262 | והוא |  |
| 7:15.263 | נאמן |  |
| 7:15.264 | בשליחותו | . |
| 7:15.265 | והמאמר |  |
| 7:15.266 | המורה |  |
| 7:15.267 | על |  |
| 7:15.268 | זה |  |
| 7:15.269 | הוא |  |
| 7:15.270 | מה |  |
| 7:15.271 | שנאמר |  |
| 7:15.272 | בזאת |  |
| 7:15.273 | תדעון |  |
| 7:15.274 | כי |  |
| 7:15.275 | ה | ' |
| 7:15.276 | שלחני |  |
| 7:15.277 | לעשות |  |
| 7:15.278 | את |  |
| 7:15.279 | כל |  |
| 7:15.280 | המעשים |  |
| 7:15.281 | האלה |  |
| 7:15.282 | כי |  |
| 7:15.283 | לא |  |
| 7:15.284 | מלבי |  |
| 7:15.285 | ע"כ | . |
| 7:15.286 | הא |  |
| 7:15.287 | למדת |  |
| 7:15.288 | שטענתם |  |
| 7:15.289 | זו |  |
| 7:15.290 | כוזבת | . |
| 7:15.291 | שהרי |  |
| 7:15.292 | רבינו |  |
| 7:15.293 | מודה |  |
| 7:15.294 | בזה |  |
| 7:15.295 | וכלל |  |
| 7:15.296 | כל |  |
| 7:15.297 | מצותיה |  |
| 7:15.298 | של |  |
| 7:15.299 | תורה |  |
| 7:15.300 | וספוריה |  |
| 7:15.301 | בזה |  |
| 7:15.302 | העיקר |  |
| 7:15.303 | השמיני |  |
| 7:15.304 | באמרו |  |
| 7:15.305 | שכל |  |
| 7:15.306 | מצות |  |
| 7:15.307 | התורה |  |
| 7:15.308 | וחילוקי |  |
| 7:15.309 | דיניהם |  |
| 7:15.310 | וסיפוריה |  |
| 7:15.311 | כולם |  |
| 7:15.312 | מפי |  |
| 7:15.313 | הגבורה | . |
| 7:15.314 | וכל |  |
| 7:15.315 | הכופר |  |
| 7:15.316 | בשום |  |
| 7:15.317 | חלק |  |
| 7:15.318 | מכל |  |
| 7:15.319 | מצות |  |
| 7:15.320 | התורה |  |
| 7:15.321 | או |  |
| 7:15.322 | אפי | ' |
| 7:15.323 | מספוריה |  |
| 7:15.324 | הרי |  |
| 7:15.325 | הוא |  |
| 7:15.326 | מין |  |
| 7:15.327 | וכופר | . |
| 7:15.328 | וכולם |  |
| 7:15.329 | נכללים |  |
| 7:15.330 | תחת |  |
| 7:15.331 | סוג |  |
| 7:15.332 | תורת |  |
| 7:15.333 | אמת |  |
| 7:15.334 | נתן |  |
| 7:15.335 | לעמו |  |
| 7:15.336 | אל | . |
| 7:15.337 | על |  |
| 7:15.338 | יד |  |
| 7:15.339 | נביאו | , |
| 7:15.340 | ואין |  |
| 7:15.341 | דבר |  |
| 7:15.342 | בה |  |
| 7:15.343 | לבטלה | . |
| 7:15.344 | וכל |  |
| 7:15.345 | מי |  |
| 7:15.346 | שאינו |  |
| 7:15.347 | מודה |  |
| 7:15.348 | באחת |  |
| 7:15.349 | מהן |  |
| 7:15.350 | הוא |  |
| 7:15.351 | מין |  |
| 7:15.352 | וכופר |  |
| 7:15.353 | ואפיקירוס |  |
| 7:15.354 | כמנשה |  |
| 7:15.355 | וסיעתו | . |
| 7:15.356 | והראיה |  |
| 7:15.357 | הגדולה |  |
| 7:15.358 | על |  |
| 7:15.359 | מה |  |
| 7:15.360 | שאמרתי |  |
| 7:15.361 | שאינם |  |
| 7:15.362 | מודים |  |
| 7:15.363 | בארבעה |  |
| 7:15.364 | עיקרים |  |
| 7:15.365 | הראשונים |  |
| 7:15.366 | שהרי |  |
| 7:15.367 | ביאר |  |
| 7:15.368 | הרב |  |
| 7:15.369 | חמדת |  |
| 7:15.370 | ימים |  |
| 7:15.371 | שלא |  |
| 7:15.372 | חרה |  |
| 7:15.373 | להאר"י |  |
| 7:15.374 | רק |  |
| 7:15.375 | על |  |
| 7:15.376 | ארבעה |  |
| 7:15.377 | בתים |  |
| 7:15.378 | ראשונים |  |
| 7:15.379 | שביגדל |  |
| 7:15.380 | שהם |  |
| 7:15.381 | סותרים |  |
| 7:15.382 | אמונתם | . |
| 7:15.383 | אבל |  |
| 7:15.384 | מהנו |  |
| 7:15.385 | אדון |  |
| 7:15.386 | עולם |  |
| 7:15.387 | ואילך |  |
| 7:15.388 | שפיר |  |
| 7:15.389 | דמי |  |
| 7:15.390 | לדידהו |  |
| 7:15.391 | לאמרו |  |
| 7:15.392 | כנז | ' |
| 7:15.393 | לעיל |  |
| 7:15.394 | סי | ' |
| 7:15.395 | מ"ח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

