# Translation pack: milhamot-hashem perek-6-12 (Milchamot ha-Shem 6:18–6:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-12`.

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
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
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
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 6:18

Hebrew (Wikisource): ועוד לא שבע להטיל קלון ופחיתות במשנה שהיא המצוה שקיבל משה מסיני כדאיתא בברכות והובא לעיל, וקורא לה שפחה, סלע אחרא, דאתקריאת קליפה, משנה. נוקבא דעבד נער, עלה אתמר בדברים לא יוסר עבד אלא דמחן ומתברין מנה כמה פסקות בלא נביעו דחכמתא. ומשבח הוא את תורתו החדשה המלאה אלוהות רבים וקורא לה ברתא דמלכא, דבגינה אמתר ודברתם אל הסלע בדבור ופיוס כברתא דמלכא, ומספר שם שמשה רבינו התודה את חטאתו זאת שנתן לנו את המשנה והתלמוד שהם סותרים ומכים על קדקד אמונתו באלהים רבים שהוא קורא לה ברתא דמלכא. ושעל כן נגזר עליו מיתה. דמאן דסריב למטרוניתא חייב מיתה, כל שכן מאן דמחי לברתא דמלכא ושעל כן נגזר עליו שלא יכנס לארץ ישראל. ושיקבר בארץ נכריה חוצה לארץ.

| id | word | punctuation after |
|---|---|---|
| 6:18.1 | ועוד |  |
| 6:18.2 | לא |  |
| 6:18.3 | שבע |  |
| 6:18.4 | להטיל |  |
| 6:18.5 | קלון |  |
| 6:18.6 | ופחיתות |  |
| 6:18.7 | במשנה |  |
| 6:18.8 | שהיא |  |
| 6:18.9 | המצוה |  |
| 6:18.10 | שקיבל |  |
| 6:18.11 | משה |  |
| 6:18.12 | מסיני |  |
| 6:18.13 | כדאיתא |  |
| 6:18.14 | בברכות |  |
| 6:18.15 | והובא |  |
| 6:18.16 | לעיל | , |
| 6:18.17 | וקורא |  |
| 6:18.18 | לה |  |
| 6:18.19 | שפחה | , |
| 6:18.20 | סלע |  |
| 6:18.21 | אחרא | , |
| 6:18.22 | דאתקריאת |  |
| 6:18.23 | קליפה | , |
| 6:18.24 | משנה | . |
| 6:18.25 | נוקבא |  |
| 6:18.26 | דעבד |  |
| 6:18.27 | נער | , |
| 6:18.28 | עלה |  |
| 6:18.29 | אתמר |  |
| 6:18.30 | בדברים |  |
| 6:18.31 | לא |  |
| 6:18.32 | יוסר |  |
| 6:18.33 | עבד |  |
| 6:18.34 | אלא |  |
| 6:18.35 | דמחן |  |
| 6:18.36 | ומתברין |  |
| 6:18.37 | מנה |  |
| 6:18.38 | כמה |  |
| 6:18.39 | פסקות |  |
| 6:18.40 | בלא |  |
| 6:18.41 | נביעו |  |
| 6:18.42 | דחכמתא | . |
| 6:18.43 | ומשבח |  |
| 6:18.44 | הוא |  |
| 6:18.45 | את |  |
| 6:18.46 | תורתו |  |
| 6:18.47 | החדשה |  |
| 6:18.48 | המלאה |  |
| 6:18.49 | אלוהות |  |
| 6:18.50 | רבים |  |
| 6:18.51 | וקורא |  |
| 6:18.52 | לה |  |
| 6:18.53 | ברתא |  |
| 6:18.54 | דמלכא | , |
| 6:18.55 | דבגינה |  |
| 6:18.56 | אמתר |  |
| 6:18.57 | ודברתם |  |
| 6:18.58 | אל |  |
| 6:18.59 | הסלע |  |
| 6:18.60 | בדבור |  |
| 6:18.61 | ופיוס |  |
| 6:18.62 | כברתא |  |
| 6:18.63 | דמלכא | , |
| 6:18.64 | ומספר |  |
| 6:18.65 | שם |  |
| 6:18.66 | שמשה |  |
| 6:18.67 | רבינו |  |
| 6:18.68 | התודה |  |
| 6:18.69 | את |  |
| 6:18.70 | חטאתו |  |
| 6:18.71 | זאת |  |
| 6:18.72 | שנתן |  |
| 6:18.73 | לנו |  |
| 6:18.74 | את |  |
| 6:18.75 | המשנה |  |
| 6:18.76 | והתלמוד |  |
| 6:18.77 | שהם |  |
| 6:18.78 | סותרים |  |
| 6:18.79 | ומכים |  |
| 6:18.80 | על |  |
| 6:18.81 | קדקד |  |
| 6:18.82 | אמונתו |  |
| 6:18.83 | באלהים |  |
| 6:18.84 | רבים |  |
| 6:18.85 | שהוא |  |
| 6:18.86 | קורא |  |
| 6:18.87 | לה |  |
| 6:18.88 | ברתא |  |
| 6:18.89 | דמלכא | . |
| 6:18.90 | ושעל |  |
| 6:18.91 | כן |  |
| 6:18.92 | נגזר |  |
| 6:18.93 | עליו |  |
| 6:18.94 | מיתה | . |
| 6:18.95 | דמאן |  |
| 6:18.96 | דסריב |  |
| 6:18.97 | למטרוניתא |  |
| 6:18.98 | חייב |  |
| 6:18.99 | מיתה | , |
| 6:18.100 | כל |  |
| 6:18.101 | שכן |  |
| 6:18.102 | מאן |  |
| 6:18.103 | דמחי |  |
| 6:18.104 | לברתא |  |
| 6:18.105 | דמלכא |  |
| 6:18.106 | ושעל |  |
| 6:18.107 | כן |  |
| 6:18.108 | נגזר |  |
| 6:18.109 | עליו |  |
| 6:18.110 | שלא |  |
| 6:18.111 | יכנס |  |
| 6:18.112 | לארץ |  |
| 6:18.113 | ישראל | . |
| 6:18.114 | ושיקבר |  |
| 6:18.115 | בארץ |  |
| 6:18.116 | נכריה |  |
| 6:18.117 | חוצה |  |
| 6:18.118 | לארץ | . |


## Milchamot ha-Shem 6:19

Hebrew (Wikisource): סט) ראה והתבונן אחי איך בדבריו אלה הוא מחלל את הקדשים כללותיה ופרטותיה של תורה המקובלים ביד כל ישראל הכתובים במשנה ובתלמוד שקיבל משה מסיני ומסרם ליהושע וכו' הקדושים לכלל האומה הישראלית, ומחבר הזהר מחללם ומורה לעזבם ולפרשו מהם כאבן משכית, וקורא להם בשמו של גנאי שפחה כי תירש גברתה, אשר תחתיה רגזה ארץ, ובשם קליפה קשה אשר ילאה כל אדם לשברה ולהוציא ממנה אוכל מועט כנטפי מים מועטים אשר לא ירוו לנפש עיף וצמא. כה עשה הפילסוף הרמאי מחבר הזהר לגרע נטפי מי ים המשנה והתלמוד בכמה מיני מגרעות אשר שת עליהם בזהר ובתיקונים. ובמחשבה רעה זו חשב לפסול ולהשכיח את המשנה והתלמוד מישראל ולסתום את פי באר מים חיים אשר ממנו יקרו כל הליכות תורה שבעל פה. ומחשבתו זאת בפירוש אמרה בכמה מקומות, וכמו שכ' ר' יצחק דלטאש בפסק המודפס בתחלת ס' הזהר הבנוי על דעת מחבר הזהר שלא כרצון רז"ל בעלי התלמוד, ומהם גם כן בתיקוני זהר חדש בתחלתו דבור המתחיל והמשכילים, שכ' שם דאסתכם קב"ה ושכינתיה למעבר האי הבורא וכו' (והמשכיל יבין איך אפשר שיסכים הקב"ה לשנות מה שכתוב בתורה כי שאל נא לימים ראשונים וכו' וכמו שביארו רז"ל).

| id | word | punctuation after |
|---|---|---|
| 6:19.1 | סט | ) |
| 6:19.2 | ראה |  |
| 6:19.3 | והתבונן |  |
| 6:19.4 | אחי |  |
| 6:19.5 | איך |  |
| 6:19.6 | בדבריו |  |
| 6:19.7 | אלה |  |
| 6:19.8 | הוא |  |
| 6:19.9 | מחלל |  |
| 6:19.10 | את |  |
| 6:19.11 | הקדשים |  |
| 6:19.12 | כללותיה |  |
| 6:19.13 | ופרטותיה |  |
| 6:19.14 | של |  |
| 6:19.15 | תורה |  |
| 6:19.16 | המקובלים |  |
| 6:19.17 | ביד |  |
| 6:19.18 | כל |  |
| 6:19.19 | ישראל |  |
| 6:19.20 | הכתובים |  |
| 6:19.21 | במשנה |  |
| 6:19.22 | ובתלמוד |  |
| 6:19.23 | שקיבל |  |
| 6:19.24 | משה |  |
| 6:19.25 | מסיני |  |
| 6:19.26 | ומסרם |  |
| 6:19.27 | ליהושע |  |
| 6:19.28 | וכו | ' |
| 6:19.29 | הקדושים |  |
| 6:19.30 | לכלל |  |
| 6:19.31 | האומה |  |
| 6:19.32 | הישראלית | , |
| 6:19.33 | ומחבר |  |
| 6:19.34 | הזהר |  |
| 6:19.35 | מחללם |  |
| 6:19.36 | ומורה |  |
| 6:19.37 | לעזבם |  |
| 6:19.38 | ולפרשו |  |
| 6:19.39 | מהם |  |
| 6:19.40 | כאבן |  |
| 6:19.41 | משכית | , |
| 6:19.42 | וקורא |  |
| 6:19.43 | להם |  |
| 6:19.44 | בשמו |  |
| 6:19.45 | של |  |
| 6:19.46 | גנאי |  |
| 6:19.47 | שפחה |  |
| 6:19.48 | כי |  |
| 6:19.49 | תירש |  |
| 6:19.50 | גברתה | , |
| 6:19.51 | אשר |  |
| 6:19.52 | תחתיה |  |
| 6:19.53 | רגזה |  |
| 6:19.54 | ארץ | , |
| 6:19.55 | ובשם |  |
| 6:19.56 | קליפה |  |
| 6:19.57 | קשה |  |
| 6:19.58 | אשר |  |
| 6:19.59 | ילאה |  |
| 6:19.60 | כל |  |
| 6:19.61 | אדם |  |
| 6:19.62 | לשברה |  |
| 6:19.63 | ולהוציא |  |
| 6:19.64 | ממנה |  |
| 6:19.65 | אוכל |  |
| 6:19.66 | מועט |  |
| 6:19.67 | כנטפי |  |
| 6:19.68 | מים |  |
| 6:19.69 | מועטים |  |
| 6:19.70 | אשר |  |
| 6:19.71 | לא |  |
| 6:19.72 | ירוו |  |
| 6:19.73 | לנפש |  |
| 6:19.74 | עיף |  |
| 6:19.75 | וצמא | . |
| 6:19.76 | כה |  |
| 6:19.77 | עשה |  |
| 6:19.78 | הפילסוף |  |
| 6:19.79 | הרמאי |  |
| 6:19.80 | מחבר |  |
| 6:19.81 | הזהר |  |
| 6:19.82 | לגרע |  |
| 6:19.83 | נטפי |  |
| 6:19.84 | מי |  |
| 6:19.85 | ים |  |
| 6:19.86 | המשנה |  |
| 6:19.87 | והתלמוד |  |
| 6:19.88 | בכמה |  |
| 6:19.89 | מיני |  |
| 6:19.90 | מגרעות |  |
| 6:19.91 | אשר |  |
| 6:19.92 | שת |  |
| 6:19.93 | עליהם |  |
| 6:19.94 | בזהר |  |
| 6:19.95 | ובתיקונים | . |
| 6:19.96 | ובמחשבה |  |
| 6:19.97 | רעה |  |
| 6:19.98 | זו |  |
| 6:19.99 | חשב |  |
| 6:19.100 | לפסול |  |
| 6:19.101 | ולהשכיח |  |
| 6:19.102 | את |  |
| 6:19.103 | המשנה |  |
| 6:19.104 | והתלמוד |  |
| 6:19.105 | מישראל |  |
| 6:19.106 | ולסתום |  |
| 6:19.107 | את |  |
| 6:19.108 | פי |  |
| 6:19.109 | באר |  |
| 6:19.110 | מים |  |
| 6:19.111 | חיים |  |
| 6:19.112 | אשר |  |
| 6:19.113 | ממנו |  |
| 6:19.114 | יקרו |  |
| 6:19.115 | כל |  |
| 6:19.116 | הליכות |  |
| 6:19.117 | תורה |  |
| 6:19.118 | שבעל |  |
| 6:19.119 | פה | . |
| 6:19.120 | ומחשבתו |  |
| 6:19.121 | זאת |  |
| 6:19.122 | בפירוש |  |
| 6:19.123 | אמרה |  |
| 6:19.124 | בכמה |  |
| 6:19.125 | מקומות | , |
| 6:19.126 | וכמו |  |
| 6:19.127 | שכ | ' |
| 6:19.128 | ר | ' |
| 6:19.129 | יצחק |  |
| 6:19.130 | דלטאש |  |
| 6:19.131 | בפסק |  |
| 6:19.132 | המודפס |  |
| 6:19.133 | בתחלת |  |
| 6:19.134 | ס | ' |
| 6:19.135 | הזהר |  |
| 6:19.136 | הבנוי |  |
| 6:19.137 | על |  |
| 6:19.138 | דעת |  |
| 6:19.139 | מחבר |  |
| 6:19.140 | הזהר |  |
| 6:19.141 | שלא |  |
| 6:19.142 | כרצון |  |
| 6:19.143 | רז"ל |  |
| 6:19.144 | בעלי |  |
| 6:19.145 | התלמוד | , |
| 6:19.146 | ומהם |  |
| 6:19.147 | גם |  |
| 6:19.148 | כן |  |
| 6:19.149 | בתיקוני |  |
| 6:19.150 | זהר |  |
| 6:19.151 | חדש |  |
| 6:19.152 | בתחלתו |  |
| 6:19.153 | דבור |  |
| 6:19.154 | המתחיל |  |
| 6:19.155 | והמשכילים | , |
| 6:19.156 | שכ | ' |
| 6:19.157 | שם |  |
| 6:19.158 | דאסתכם |  |
| 6:19.159 | קב"ה |  |
| 6:19.160 | ושכינתיה |  |
| 6:19.161 | למעבר |  |
| 6:19.162 | האי |  |
| 6:19.163 | הבורא |  |
| 6:19.164 | וכו | ' |
| 6:19.165 | והמשכיל |  |
| 6:19.166 | יבין |  |
| 6:19.167 | איך |  |
| 6:19.168 | אפשר |  |
| 6:19.169 | שיסכים |  |
| 6:19.170 | הקב"ה |  |
| 6:19.171 | לשנות |  |
| 6:19.172 | מה |  |
| 6:19.173 | שכתוב |  |
| 6:19.174 | בתורה |  |
| 6:19.175 | כי |  |
| 6:19.176 | שאל |  |
| 6:19.177 | נא |  |
| 6:19.178 | לימים |  |
| 6:19.179 | ראשונים |  |
| 6:19.180 | וכו | ' |
| 6:19.181 | וכמו |  |
| 6:19.182 | שביארו |  |
| 6:19.183 | רז"ל | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

