# Translation pack: milhamot-hashem perek-6-24 (Milchamot ha-Shem 6:37–6:37)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-24.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-24`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
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
- **יחוד / ייחוד / ייחדו**: the doctrine → "unity", fixed, not "monotheism" or
  "oneness" (אחדות → "oneness"). The Kabbalistic act → "unification" (יחודים
  → "unifications"). The verb, God as object → "to declare one" ("ליחדו" → "to
  declare Him one"); Kabbalistic entities as object → "to unify"; of
  dedicating a thing to one → "dedicate ... solely to" (note: "unify"). Where
  he plays on the senses: "unity" for the noun, "unify" for the verb, the play
  in a note.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to". Not technical: "לידי התגשמות" → "to realization". "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 6:37

Hebrew (Wikisource): וכדי שלא יטעה הפתי הסכל מסכלותו ופתיותיו ומיעוט הבנתו ויחשוב שהשי"ת הוא על דרך הנראה ממליצת הספר שיש לו גשמות, הזהירנו בתורה שנשמר מאד מלחשוב שיש לבורא ית' צורה ולא דמיון. והוא אמרו ונשמרתם מאד לנפשותיכם כי לא ראיתם כל תמונה. והנביאים הזהירונו להמנע מלדמותו בשכל, ומלהמשילו לרעיון כאמרו ואל מי תדמיון אל, ומה דמות תערכו לו וכתוב ואל מי תדמיוני ואשוה. וכל זה צריך בעל אמונתינו לידע, ואז יוכל להבדיל בין הבורא וזולתו, ויהיה מיחד את השי"ת יחוד גמור כי מי שאינו יודע דרכי ההבדל אינו יכול ליחדו ית' יחוד אמתי. אמנם כאשר ידענו על מתכונתו אז ותקרב לאמתת בוראו ותתקרב נפשו להשגת המציאה. ותתעלה בהשגתה על מה כונתה כשאין אדם זוכה לידע אמתת בוראו אפי' יהיה מיחד אותו בלבו הרי הוא כאלו יש מחיצה בינו ובין בוראו ית'. כי זה הוא דבר אשר יגיע לו הפתי או הקטן. אשר איננו יודע דרכי האמונה ועניינה. ואין אמתתה קבועה בלבו. אבל כל עוד שמתחיל לחקור לידע אמתתו ית'. אזי המחיצה הולכת ומתמעטת. והנפש הולכת ומתקרבת אל בוראה יתעלה. ותתעלה נפשו ושכלו בתענו ג השגתו ית'. כי בהיות הנפש מתעסקת בדברים שכליים אז היא נמשכת אחריהם. לפי שהיא גם כן ממקור השכל נצמאת ולכן היא משתוקקת ונכספה להדמות לבוראה ית' ואינה יכולה להבדל ממנו עד שנמצאת דבוקה בו. ועל כיוצא בענין זה אמר ית' ובו תדבק. וגם על זה אמר הכתוב קרוב ה' לכל קוראיו. כי מאחר שהבורא ית' אינו גוף ולא כח בגוף אי אפשר היותו קרוב או רחוק מן הנפש. אלא רחוק מצד המחיצה המבדלת שהיא הסכלות כשאין לנפש ידיעה באמתתו ית'. וקרוב לנפש מצד הידיעה וההשגה. ולכן כל בעל נפש צריך לתור ולבקש על דרך ידיעת אמיתת מציאות הבורא ית'. כי הוא עיקר ויסוד גדול לאמונה. ועל דבר זה צוה אותנו בתורה ואמר וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:37.1 | וכדי |  |
| 6:37.2 | שלא |  |
| 6:37.3 | יטעה |  |
| 6:37.4 | הפתי |  |
| 6:37.5 | הסכל |  |
| 6:37.6 | מסכלותו |  |
| 6:37.7 | ופתיותיו |  |
| 6:37.8 | ומיעוט |  |
| 6:37.9 | הבנתו |  |
| 6:37.10 | ויחשוב |  |
| 6:37.11 | שהשי"ת |  |
| 6:37.12 | הוא |  |
| 6:37.13 | על |  |
| 6:37.14 | דרך |  |
| 6:37.15 | הנראה |  |
| 6:37.16 | ממליצת |  |
| 6:37.17 | הספר |  |
| 6:37.18 | שיש |  |
| 6:37.19 | לו |  |
| 6:37.20 | גשמות | , |
| 6:37.21 | הזהירנו |  |
| 6:37.22 | בתורה |  |
| 6:37.23 | שנשמר |  |
| 6:37.24 | מאד |  |
| 6:37.25 | מלחשוב |  |
| 6:37.26 | שיש |  |
| 6:37.27 | לבורא |  |
| 6:37.28 | ית | ' |
| 6:37.29 | צורה |  |
| 6:37.30 | ולא |  |
| 6:37.31 | דמיון | . |
| 6:37.32 | והוא |  |
| 6:37.33 | אמרו |  |
| 6:37.34 | ונשמרתם |  |
| 6:37.35 | מאד |  |
| 6:37.36 | לנפשותיכם |  |
| 6:37.37 | כי |  |
| 6:37.38 | לא |  |
| 6:37.39 | ראיתם |  |
| 6:37.40 | כל |  |
| 6:37.41 | תמונה | . |
| 6:37.42 | והנביאים |  |
| 6:37.43 | הזהירונו |  |
| 6:37.44 | להמנע |  |
| 6:37.45 | מלדמותו |  |
| 6:37.46 | בשכל | , |
| 6:37.47 | ומלהמשילו |  |
| 6:37.48 | לרעיון |  |
| 6:37.49 | כאמרו |  |
| 6:37.50 | ואל |  |
| 6:37.51 | מי |  |
| 6:37.52 | תדמיון |  |
| 6:37.53 | אל | , |
| 6:37.54 | ומה |  |
| 6:37.55 | דמות |  |
| 6:37.56 | תערכו |  |
| 6:37.57 | לו |  |
| 6:37.58 | וכתוב |  |
| 6:37.59 | ואל |  |
| 6:37.60 | מי |  |
| 6:37.61 | תדמיוני |  |
| 6:37.62 | ואשוה | . |
| 6:37.63 | וכל |  |
| 6:37.64 | זה |  |
| 6:37.65 | צריך |  |
| 6:37.66 | בעל |  |
| 6:37.67 | אמונתינו |  |
| 6:37.68 | לידע | , |
| 6:37.69 | ואז |  |
| 6:37.70 | יוכל |  |
| 6:37.71 | להבדיל |  |
| 6:37.72 | בין |  |
| 6:37.73 | הבורא |  |
| 6:37.74 | וזולתו | , |
| 6:37.75 | ויהיה |  |
| 6:37.76 | מיחד |  |
| 6:37.77 | את |  |
| 6:37.78 | השי"ת |  |
| 6:37.79 | יחוד |  |
| 6:37.80 | גמור |  |
| 6:37.81 | כי |  |
| 6:37.82 | מי |  |
| 6:37.83 | שאינו |  |
| 6:37.84 | יודע |  |
| 6:37.85 | דרכי |  |
| 6:37.86 | ההבדל |  |
| 6:37.87 | אינו |  |
| 6:37.88 | יכול |  |
| 6:37.89 | ליחדו |  |
| 6:37.90 | ית | ' |
| 6:37.91 | יחוד |  |
| 6:37.92 | אמתי | . |
| 6:37.93 | אמנם |  |
| 6:37.94 | כאשר |  |
| 6:37.95 | ידענו |  |
| 6:37.96 | על |  |
| 6:37.97 | מתכונתו |  |
| 6:37.98 | אז |  |
| 6:37.99 | ותקרב |  |
| 6:37.100 | לאמתת |  |
| 6:37.101 | בוראו |  |
| 6:37.102 | ותתקרב |  |
| 6:37.103 | נפשו |  |
| 6:37.104 | להשגת |  |
| 6:37.105 | המציאה | . |
| 6:37.106 | ותתעלה |  |
| 6:37.107 | בהשגתה |  |
| 6:37.108 | על |  |
| 6:37.109 | מה |  |
| 6:37.110 | כונתה |  |
| 6:37.111 | כשאין |  |
| 6:37.112 | אדם |  |
| 6:37.113 | זוכה |  |
| 6:37.114 | לידע |  |
| 6:37.115 | אמתת |  |
| 6:37.116 | בוראו |  |
| 6:37.117 | אפי | ' |
| 6:37.118 | יהיה |  |
| 6:37.119 | מיחד |  |
| 6:37.120 | אותו |  |
| 6:37.121 | בלבו |  |
| 6:37.122 | הרי |  |
| 6:37.123 | הוא |  |
| 6:37.124 | כאלו |  |
| 6:37.125 | יש |  |
| 6:37.126 | מחיצה |  |
| 6:37.127 | בינו |  |
| 6:37.128 | ובין |  |
| 6:37.129 | בוראו |  |
| 6:37.130 | ית | '. |
| 6:37.131 | כי |  |
| 6:37.132 | זה |  |
| 6:37.133 | הוא |  |
| 6:37.134 | דבר |  |
| 6:37.135 | אשר |  |
| 6:37.136 | יגיע |  |
| 6:37.137 | לו |  |
| 6:37.138 | הפתי |  |
| 6:37.139 | או |  |
| 6:37.140 | הקטן | . |
| 6:37.141 | אשר |  |
| 6:37.142 | איננו |  |
| 6:37.143 | יודע |  |
| 6:37.144 | דרכי |  |
| 6:37.145 | האמונה |  |
| 6:37.146 | ועניינה | . |
| 6:37.147 | ואין |  |
| 6:37.148 | אמתתה |  |
| 6:37.149 | קבועה |  |
| 6:37.150 | בלבו | . |
| 6:37.151 | אבל |  |
| 6:37.152 | כל |  |
| 6:37.153 | עוד |  |
| 6:37.154 | שמתחיל |  |
| 6:37.155 | לחקור |  |
| 6:37.156 | לידע |  |
| 6:37.157 | אמתתו |  |
| 6:37.158 | ית | '. |
| 6:37.159 | אזי |  |
| 6:37.160 | המחיצה |  |
| 6:37.161 | הולכת |  |
| 6:37.162 | ומתמעטת | . |
| 6:37.163 | והנפש |  |
| 6:37.164 | הולכת |  |
| 6:37.165 | ומתקרבת |  |
| 6:37.166 | אל |  |
| 6:37.167 | בוראה |  |
| 6:37.168 | יתעלה | . |
| 6:37.169 | ותתעלה |  |
| 6:37.170 | נפשו |  |
| 6:37.171 | ושכלו |  |
| 6:37.172 | בתענו |  |
| 6:37.173 | ג |  |
| 6:37.174 | השגתו |  |
| 6:37.175 | ית | '. |
| 6:37.176 | כי |  |
| 6:37.177 | בהיות |  |
| 6:37.178 | הנפש |  |
| 6:37.179 | מתעסקת |  |
| 6:37.180 | בדברים |  |
| 6:37.181 | שכליים |  |
| 6:37.182 | אז |  |
| 6:37.183 | היא |  |
| 6:37.184 | נמשכת |  |
| 6:37.185 | אחריהם | . |
| 6:37.186 | לפי |  |
| 6:37.187 | שהיא |  |
| 6:37.188 | גם |  |
| 6:37.189 | כן |  |
| 6:37.190 | ממקור |  |
| 6:37.191 | השכל |  |
| 6:37.192 | נצמאת |  |
| 6:37.193 | ולכן |  |
| 6:37.194 | היא |  |
| 6:37.195 | משתוקקת |  |
| 6:37.196 | ונכספה |  |
| 6:37.197 | להדמות |  |
| 6:37.198 | לבוראה |  |
| 6:37.199 | ית | ' |
| 6:37.200 | ואינה |  |
| 6:37.201 | יכולה |  |
| 6:37.202 | להבדל |  |
| 6:37.203 | ממנו |  |
| 6:37.204 | עד |  |
| 6:37.205 | שנמצאת |  |
| 6:37.206 | דבוקה |  |
| 6:37.207 | בו | . |
| 6:37.208 | ועל |  |
| 6:37.209 | כיוצא |  |
| 6:37.210 | בענין |  |
| 6:37.211 | זה |  |
| 6:37.212 | אמר |  |
| 6:37.213 | ית | ' |
| 6:37.214 | ובו |  |
| 6:37.215 | תדבק | . |
| 6:37.216 | וגם |  |
| 6:37.217 | על |  |
| 6:37.218 | זה |  |
| 6:37.219 | אמר |  |
| 6:37.220 | הכתוב |  |
| 6:37.221 | קרוב |  |
| 6:37.222 | ה | ' |
| 6:37.223 | לכל |  |
| 6:37.224 | קוראיו | . |
| 6:37.225 | כי |  |
| 6:37.226 | מאחר |  |
| 6:37.227 | שהבורא |  |
| 6:37.228 | ית | ' |
| 6:37.229 | אינו |  |
| 6:37.230 | גוף |  |
| 6:37.231 | ולא |  |
| 6:37.232 | כח |  |
| 6:37.233 | בגוף |  |
| 6:37.234 | אי |  |
| 6:37.235 | אפשר |  |
| 6:37.236 | היותו |  |
| 6:37.237 | קרוב |  |
| 6:37.238 | או |  |
| 6:37.239 | רחוק |  |
| 6:37.240 | מן |  |
| 6:37.241 | הנפש | . |
| 6:37.242 | אלא |  |
| 6:37.243 | רחוק |  |
| 6:37.244 | מצד |  |
| 6:37.245 | המחיצה |  |
| 6:37.246 | המבדלת |  |
| 6:37.247 | שהיא |  |
| 6:37.248 | הסכלות |  |
| 6:37.249 | כשאין |  |
| 6:37.250 | לנפש |  |
| 6:37.251 | ידיעה |  |
| 6:37.252 | באמתתו |  |
| 6:37.253 | ית | '. |
| 6:37.254 | וקרוב |  |
| 6:37.255 | לנפש |  |
| 6:37.256 | מצד |  |
| 6:37.257 | הידיעה |  |
| 6:37.258 | וההשגה | . |
| 6:37.259 | ולכן |  |
| 6:37.260 | כל |  |
| 6:37.261 | בעל |  |
| 6:37.262 | נפש |  |
| 6:37.263 | צריך |  |
| 6:37.264 | לתור |  |
| 6:37.265 | ולבקש |  |
| 6:37.266 | על |  |
| 6:37.267 | דרך |  |
| 6:37.268 | ידיעת |  |
| 6:37.269 | אמיתת |  |
| 6:37.270 | מציאות |  |
| 6:37.271 | הבורא |  |
| 6:37.272 | ית | '. |
| 6:37.273 | כי |  |
| 6:37.274 | הוא |  |
| 6:37.275 | עיקר |  |
| 6:37.276 | ויסוד |  |
| 6:37.277 | גדול |  |
| 6:37.278 | לאמונה | . |
| 6:37.279 | ועל |  |
| 6:37.280 | דבר |  |
| 6:37.281 | זה |  |
| 6:37.282 | צוה |  |
| 6:37.283 | אותנו |  |
| 6:37.284 | בתורה |  |
| 6:37.285 | ואמר |  |
| 6:37.286 | וידעת |  |
| 6:37.287 | היום |  |
| 6:37.288 | והשבות |  |
| 6:37.289 | אל |  |
| 6:37.290 | לבבך |  |
| 6:37.291 | כי |  |
| 6:37.292 | ה | ' |
| 6:37.293 | הוא |  |
| 6:37.294 | האלהים |  |
| 6:37.295 | בשמים |  |
| 6:37.296 | ממעל |  |
| 6:37.297 | ועל |  |
| 6:37.298 | הארץ |  |
| 6:37.299 | מתחת |  |
| 6:37.300 | אין |  |
| 6:37.301 | עוד |  |
| 6:37.302 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

