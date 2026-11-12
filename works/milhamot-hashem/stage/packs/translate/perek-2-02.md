# Translation pack: milhamot-hashem perek-2-02 (Milchamot ha-Shem 2:5–2:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-02.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-02`.

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
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות**: by context, never "sefirot". Hermeneutical rules: "הי״ג
  מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the thirteen rules of Rabbi
  Ishmael and the thirty-three rules of Rabbi Yose the Galilean" (the number
  as written; the usual thirty-two in a note; note: literally "measures").
  "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes of mercy",
  a different thirteen. Of God or the sefirot: "attribute" ("מדת הדין" → "the
  attribute of judgment"). Also "measure", and "quality / trait / virtue" of
  character. מסכת מדות → "tractate Middot".
- **אין סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / בריאה**: fixed, translated, the verbs linked:
  "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; נאצל → "emanated"; המאציל → "the Emanator". The four worlds
  (אבי"ע): "Emanation, Creation, Formation, Making" (verbs: "emanated,
  created, formed, made"). נאצל and נברא stay distinct ("נברא ולא בורא" →
  "created, not creator").
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure".
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז / עכו"ם / עבודה אחרת**: עבודה זרה → "idolatry" (note once:
  "strange worship"); the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of
  idols". עכו"ם → "idolaters" (note: "worshippers of stars and
  constellations"), never "gentiles" or "Christians", even of a Christian king
  or the Trinity. עבודה אחרת → "other worship" (note first in a unit). אליל →
  "idol"; האלילה → "the she-idol"; אל זר → "a strange god"; אל אחר → "another
  god"; אלהים אחרים → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to". Not technical: "לידי התגשמות" → "to realization". "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
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


## Milchamot ha-Shem 2:5

Hebrew (Wikisource): ואיך נניח כל האלוהות אשר למעלה מעולם דאצילות ונעבוד דוקא לזעיר ונוקביה שבעולם האצילות שכבר קדמו לו כמה עתיקין ואריכי אנפין ואו"א וזעיר ונוקביה בעולמות העליונים אין מספר הקרובים לאין סוף יותר?

| id | word | punctuation after |
|---|---|---|
| 2:5.1 | ואיך |  |
| 2:5.2 | נניח |  |
| 2:5.3 | כל |  |
| 2:5.4 | האלוהות |  |
| 2:5.5 | אשר |  |
| 2:5.6 | למעלה |  |
| 2:5.7 | מעולם |  |
| 2:5.8 | דאצילות |  |
| 2:5.9 | ונעבוד |  |
| 2:5.10 | דוקא |  |
| 2:5.11 | לזעיר |  |
| 2:5.12 | ונוקביה |  |
| 2:5.13 | שבעולם |  |
| 2:5.14 | האצילות |  |
| 2:5.15 | שכבר |  |
| 2:5.16 | קדמו |  |
| 2:5.17 | לו |  |
| 2:5.18 | כמה |  |
| 2:5.19 | עתיקין |  |
| 2:5.20 | ואריכי |  |
| 2:5.21 | אנפין |  |
| 2:5.22 | ואו"א |  |
| 2:5.23 | וזעיר |  |
| 2:5.24 | ונוקביה |  |
| 2:5.25 | בעולמות |  |
| 2:5.26 | העליונים |  |
| 2:5.27 | אין |  |
| 2:5.28 | מספר |  |
| 2:5.29 | הקרובים |  |
| 2:5.30 | לאין |  |
| 2:5.31 | סוף |  |
| 2:5.32 | יותר | ? |


## Milchamot ha-Shem 2:6

Hebrew (Wikisource): ורבותינו הראשונים אמרו שאין ראוי לעבוד ולהודות ולגדל ולפאר רק לאלוה הראשון לכל הנמצאים מבלעדיו! ולמה נניח האין סוף העיקרי המצטמצם למעלה מכולם עם כל הפרצופים הקרובים אליו, ונעבוד לזה הרחוק ממנו מאוד שהוא פרצוף אחרון שבעולם האצילות? את כל אלה הביננו אדונינו נלכה נרדפה את ה' יוצרינו יוצר בראשית. ואל תצוק בנו דבר בליעל ח"ו. כי מי אסף חן ושכל טוב בחפניו השכל וידוע את ה' בחקירת העניינים האלהיים בעיקרי האמונה כמוך. מי יעלה בהר ה' ומי יקום במקום קדשו כמוך, כי בער אנכי מאיש ולא בינת אדם לי. וצור יודע תעלומות לב, בוחן לבבות כי לא נתתי שינה לעיני ונגזלה שנתי מעיני עד אמצא פתרונים והבא לטהר מסייעין אותו, להמציא לו מורה דרך הקדש בשכל ומדע ככלכל ודרדע, ובין חכמי האמת חכם ואומן, כאיתן והימן, שאולי הרבנים החסרים מן החכמה הזאת לנו משאות שוא ומדוחים ח"ו. ואתה אדני לך נגלו תעלומותיה וצפונך תמלא בטנם של תלמידיך, ישבע ויתענגו מטובך, ועלינו לשלם שכר פעולתך. ואתה אל תמנע טוב מבעליו. והנני משביעך באלהי השמים האר פניך אלינו לעות את יעף דבר דבור על אפניו והיינו עבדים לאדוני. ושכר רב טובך תבענה שפתינו, ואז לא נבוש לא בעה"ז ולא בעה"ב ונשובה נרעה בשושנים מסוד חכמים ונבונים.

| id | word | punctuation after |
|---|---|---|
| 2:6.1 | ורבותינו |  |
| 2:6.2 | הראשונים |  |
| 2:6.3 | אמרו |  |
| 2:6.4 | שאין |  |
| 2:6.5 | ראוי |  |
| 2:6.6 | לעבוד |  |
| 2:6.7 | ולהודות |  |
| 2:6.8 | ולגדל |  |
| 2:6.9 | ולפאר |  |
| 2:6.10 | רק |  |
| 2:6.11 | לאלוה |  |
| 2:6.12 | הראשון |  |
| 2:6.13 | לכל |  |
| 2:6.14 | הנמצאים |  |
| 2:6.15 | מבלעדיו | ! |
| 2:6.16 | ולמה |  |
| 2:6.17 | נניח |  |
| 2:6.18 | האין |  |
| 2:6.19 | סוף |  |
| 2:6.20 | העיקרי |  |
| 2:6.21 | המצטמצם |  |
| 2:6.22 | למעלה |  |
| 2:6.23 | מכולם |  |
| 2:6.24 | עם |  |
| 2:6.25 | כל |  |
| 2:6.26 | הפרצופים |  |
| 2:6.27 | הקרובים |  |
| 2:6.28 | אליו | , |
| 2:6.29 | ונעבוד |  |
| 2:6.30 | לזה |  |
| 2:6.31 | הרחוק |  |
| 2:6.32 | ממנו |  |
| 2:6.33 | מאוד |  |
| 2:6.34 | שהוא |  |
| 2:6.35 | פרצוף |  |
| 2:6.36 | אחרון |  |
| 2:6.37 | שבעולם |  |
| 2:6.38 | האצילות | ? |
| 2:6.39 | את |  |
| 2:6.40 | כל |  |
| 2:6.41 | אלה |  |
| 2:6.42 | הביננו |  |
| 2:6.43 | אדונינו |  |
| 2:6.44 | נלכה |  |
| 2:6.45 | נרדפה |  |
| 2:6.46 | את |  |
| 2:6.47 | ה | ' |
| 2:6.48 | יוצרינו |  |
| 2:6.49 | יוצר |  |
| 2:6.50 | בראשית | . |
| 2:6.51 | ואל |  |
| 2:6.52 | תצוק |  |
| 2:6.53 | בנו |  |
| 2:6.54 | דבר |  |
| 2:6.55 | בליעל |  |
| 2:6.56 | ח"ו | . |
| 2:6.57 | כי |  |
| 2:6.58 | מי |  |
| 2:6.59 | אסף |  |
| 2:6.60 | חן |  |
| 2:6.61 | ושכל |  |
| 2:6.62 | טוב |  |
| 2:6.63 | בחפניו |  |
| 2:6.64 | השכל |  |
| 2:6.65 | וידוע |  |
| 2:6.66 | את |  |
| 2:6.67 | ה | ' |
| 2:6.68 | בחקירת |  |
| 2:6.69 | העניינים |  |
| 2:6.70 | האלהיים |  |
| 2:6.71 | בעיקרי |  |
| 2:6.72 | האמונה |  |
| 2:6.73 | כמוך | . |
| 2:6.74 | מי |  |
| 2:6.75 | יעלה |  |
| 2:6.76 | בהר |  |
| 2:6.77 | ה | ' |
| 2:6.78 | ומי |  |
| 2:6.79 | יקום |  |
| 2:6.80 | במקום |  |
| 2:6.81 | קדשו |  |
| 2:6.82 | כמוך | , |
| 2:6.83 | כי |  |
| 2:6.84 | בער |  |
| 2:6.85 | אנכי |  |
| 2:6.86 | מאיש |  |
| 2:6.87 | ולא |  |
| 2:6.88 | בינת |  |
| 2:6.89 | אדם |  |
| 2:6.90 | לי | . |
| 2:6.91 | וצור |  |
| 2:6.92 | יודע |  |
| 2:6.93 | תעלומות |  |
| 2:6.94 | לב | , |
| 2:6.95 | בוחן |  |
| 2:6.96 | לבבות |  |
| 2:6.97 | כי |  |
| 2:6.98 | לא |  |
| 2:6.99 | נתתי |  |
| 2:6.100 | שינה |  |
| 2:6.101 | לעיני |  |
| 2:6.102 | ונגזלה |  |
| 2:6.103 | שנתי |  |
| 2:6.104 | מעיני |  |
| 2:6.105 | עד |  |
| 2:6.106 | אמצא |  |
| 2:6.107 | פתרונים |  |
| 2:6.108 | והבא |  |
| 2:6.109 | לטהר |  |
| 2:6.110 | מסייעין |  |
| 2:6.111 | אותו | , |
| 2:6.112 | להמציא |  |
| 2:6.113 | לו |  |
| 2:6.114 | מורה |  |
| 2:6.115 | דרך |  |
| 2:6.116 | הקדש |  |
| 2:6.117 | בשכל |  |
| 2:6.118 | ומדע |  |
| 2:6.119 | ככלכל |  |
| 2:6.120 | ודרדע | , |
| 2:6.121 | ובין |  |
| 2:6.122 | חכמי |  |
| 2:6.123 | האמת |  |
| 2:6.124 | חכם |  |
| 2:6.125 | ואומן | , |
| 2:6.126 | כאיתן |  |
| 2:6.127 | והימן | , |
| 2:6.128 | שאולי |  |
| 2:6.129 | הרבנים |  |
| 2:6.130 | החסרים |  |
| 2:6.131 | מן |  |
| 2:6.132 | החכמה |  |
| 2:6.133 | הזאת |  |
| 2:6.134 | לנו |  |
| 2:6.135 | משאות |  |
| 2:6.136 | שוא |  |
| 2:6.137 | ומדוחים |  |
| 2:6.138 | ח"ו | . |
| 2:6.139 | ואתה |  |
| 2:6.140 | אדני |  |
| 2:6.141 | לך |  |
| 2:6.142 | נגלו |  |
| 2:6.143 | תעלומותיה |  |
| 2:6.144 | וצפונך |  |
| 2:6.145 | תמלא |  |
| 2:6.146 | בטנם |  |
| 2:6.147 | של |  |
| 2:6.148 | תלמידיך | , |
| 2:6.149 | ישבע |  |
| 2:6.150 | ויתענגו |  |
| 2:6.151 | מטובך | , |
| 2:6.152 | ועלינו |  |
| 2:6.153 | לשלם |  |
| 2:6.154 | שכר |  |
| 2:6.155 | פעולתך | . |
| 2:6.156 | ואתה |  |
| 2:6.157 | אל |  |
| 2:6.158 | תמנע |  |
| 2:6.159 | טוב |  |
| 2:6.160 | מבעליו | . |
| 2:6.161 | והנני |  |
| 2:6.162 | משביעך |  |
| 2:6.163 | באלהי |  |
| 2:6.164 | השמים |  |
| 2:6.165 | האר |  |
| 2:6.166 | פניך |  |
| 2:6.167 | אלינו |  |
| 2:6.168 | לעות |  |
| 2:6.169 | את |  |
| 2:6.170 | יעף |  |
| 2:6.171 | דבר |  |
| 2:6.172 | דבור |  |
| 2:6.173 | על |  |
| 2:6.174 | אפניו |  |
| 2:6.175 | והיינו |  |
| 2:6.176 | עבדים |  |
| 2:6.177 | לאדוני | . |
| 2:6.178 | ושכר |  |
| 2:6.179 | רב |  |
| 2:6.180 | טובך |  |
| 2:6.181 | תבענה |  |
| 2:6.182 | שפתינו | , |
| 2:6.183 | ואז |  |
| 2:6.184 | לא |  |
| 2:6.185 | נבוש |  |
| 2:6.186 | לא |  |
| 2:6.187 | בעה"ז |  |
| 2:6.188 | ולא |  |
| 2:6.189 | בעה"ב |  |
| 2:6.190 | ונשובה |  |
| 2:6.191 | נרעה |  |
| 2:6.192 | בשושנים |  |
| 2:6.193 | מסוד |  |
| 2:6.194 | חכמים |  |
| 2:6.195 | ונבונים | . |


## Milchamot ha-Shem 2:7

Hebrew (Wikisource): תשובה שנית ראיתי את כל הרשום בכתב האלוף, אילן שענפיו מרובים, רומו כגובה ארזים, אחר המחילה על השהייה כי טרדת הזמן נטויה, וכל עכבה לטובה. בהשקיפי על המכתב והנה פערת פיך בהזכירך אלוהות בחשבך שהוא כנגדינו. ונחנו מה כי תלינו עלינו ובזיונך אינו מגיע אלינו כי אם אל חכמי ישראל אשר קטנם עבה ממתניך וכל ישראל נחים אליהם ובזכותם העולם עומד. אי עליך כי הוצאת עליהם דברים אשר לא כדת, בתומה ותחתיה ערמה. אבל כבר נתבאר דינך בפי כל בעלי חכמה. ומה שכתבת על כסא אליהו כי בדף ג' מדמה אחדותו ית' לשאר אחדים. כבר אתה הולך חשכים, הלא תחילת דבריו דף עשרים שכל דמיונות אלו על אורות קדושים אשר אין להם תפיסה וצורה ודמיון, לא בשכל ולא בדעת כלל ועיקר, אלא כדי שיוכל השכל לקבל אותם הורשינו לדמות. ובדף ד' ע"ב כתב שכל הכנויים והתארים הם בדרך דמיון בספירות ואודות קדושים הנאצלים, אבל בו ית' אין בו שום דמיון וצורה ותפישה במחשבה כלל ועיקר. ומה שכתבת איך נניח שאר הפרצופים ונעבוד לזעיר אנפין עם כח שבו א"ס ב"ה. כבר ביאר הר' כס"א בדף כ"ה ע"ב שבקראינו לז"א אנחנו קורין לכל הפרצופים כאחד עם נשמה המסתתרת בהם. ומה שכתבת למה נניח הפרצופים הקרובים לא"ס ונקרא לז"א. לפי שידוע שהמידות הפנימיות יש מהם דין ויש מהם רחמים, וזה כלול מכולם, ולזה רצה המאציל לפעול פעולותיו על ידי כלי זה. ועד שאתה מקשה על דמיונות אודות עליונים שאין לך רשות להרהר בהן במכל שכן ללגלג, תקשה על ידך הגופנית שהיא עושה כל המלאכות, והראש שלמעלה אינו עושה כלום. ג"כ כח הדבור שהוא פארך על כל בעל חי למה לא ניתן למעלה מעיניך. אלא כך גזרה חכמתו ית'. ואם יש לך עצה הוגנת תלך אצל בוראך ותיעצנו. ומה שכתבת שדברי כס"א נגד הרמב"ם וכו' שכתבו שאינו כח בגוף, ולפי דברי כס"א הוא כח בגוף. הלא כל המקובלים פה אחד עונים ואומרים שכל המגשים בספירות כמוך הרי זה כופר ואין לו חלק באלהי ישראל. ומה שכינו אותם אורות דקים אינו אלא לשכך את האזן! יש לנו לשאול על מה שכתב הרמב"ם: יכוין לבו כנגד השכינה ויתפלל. ובמורה כ' שהשכינה אור נברא. ויותר מזה מה שכתב על כבוד ה' שהוא נברא. אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחווים לנברא. ויותר על זה מה שכתב שכל צורה שהיו רואים הנביאים ומדברים עמה היא אור נברא. א"כ אין לנו שום נבואה ע"י הבורא, אלא כולם מצד הנבראים. ואם נאמר שהיה הנברא דבק בבורא, א"כ לדעתך הוא כח בגוף! ותשובתך מהרה תצמח, ס' מאנו דעת אותי נאם ה'.

| id | word | punctuation after |
|---|---|---|
| 2:7.1 | תשובה |  |
| 2:7.2 | שנית |  |
| 2:7.3 | ראיתי |  |
| 2:7.4 | את |  |
| 2:7.5 | כל |  |
| 2:7.6 | הרשום |  |
| 2:7.7 | בכתב |  |
| 2:7.8 | האלוף | , |
| 2:7.9 | אילן |  |
| 2:7.10 | שענפיו |  |
| 2:7.11 | מרובים | , |
| 2:7.12 | רומו |  |
| 2:7.13 | כגובה |  |
| 2:7.14 | ארזים | , |
| 2:7.15 | אחר |  |
| 2:7.16 | המחילה |  |
| 2:7.17 | על |  |
| 2:7.18 | השהייה |  |
| 2:7.19 | כי |  |
| 2:7.20 | טרדת |  |
| 2:7.21 | הזמן |  |
| 2:7.22 | נטויה | , |
| 2:7.23 | וכל |  |
| 2:7.24 | עכבה |  |
| 2:7.25 | לטובה | . |
| 2:7.26 | בהשקיפי |  |
| 2:7.27 | על |  |
| 2:7.28 | המכתב |  |
| 2:7.29 | והנה |  |
| 2:7.30 | פערת |  |
| 2:7.31 | פיך |  |
| 2:7.32 | בהזכירך |  |
| 2:7.33 | אלוהות |  |
| 2:7.34 | בחשבך |  |
| 2:7.35 | שהוא |  |
| 2:7.36 | כנגדינו | . |
| 2:7.37 | ונחנו |  |
| 2:7.38 | מה |  |
| 2:7.39 | כי |  |
| 2:7.40 | תלינו |  |
| 2:7.41 | עלינו |  |
| 2:7.42 | ובזיונך |  |
| 2:7.43 | אינו |  |
| 2:7.44 | מגיע |  |
| 2:7.45 | אלינו |  |
| 2:7.46 | כי |  |
| 2:7.47 | אם |  |
| 2:7.48 | אל |  |
| 2:7.49 | חכמי |  |
| 2:7.50 | ישראל |  |
| 2:7.51 | אשר |  |
| 2:7.52 | קטנם |  |
| 2:7.53 | עבה |  |
| 2:7.54 | ממתניך |  |
| 2:7.55 | וכל |  |
| 2:7.56 | ישראל |  |
| 2:7.57 | נחים |  |
| 2:7.58 | אליהם |  |
| 2:7.59 | ובזכותם |  |
| 2:7.60 | העולם |  |
| 2:7.61 | עומד | . |
| 2:7.62 | אי |  |
| 2:7.63 | עליך |  |
| 2:7.64 | כי |  |
| 2:7.65 | הוצאת |  |
| 2:7.66 | עליהם |  |
| 2:7.67 | דברים |  |
| 2:7.68 | אשר |  |
| 2:7.69 | לא |  |
| 2:7.70 | כדת | , |
| 2:7.71 | בתומה |  |
| 2:7.72 | ותחתיה |  |
| 2:7.73 | ערמה | . |
| 2:7.74 | אבל |  |
| 2:7.75 | כבר |  |
| 2:7.76 | נתבאר |  |
| 2:7.77 | דינך |  |
| 2:7.78 | בפי |  |
| 2:7.79 | כל |  |
| 2:7.80 | בעלי |  |
| 2:7.81 | חכמה | . |
| 2:7.82 | ומה |  |
| 2:7.83 | שכתבת |  |
| 2:7.84 | על |  |
| 2:7.85 | כסא |  |
| 2:7.86 | אליהו |  |
| 2:7.87 | כי |  |
| 2:7.88 | בדף |  |
| 2:7.89 | ג | ' |
| 2:7.90 | מדמה |  |
| 2:7.91 | אחדותו |  |
| 2:7.92 | ית | ' |
| 2:7.93 | לשאר |  |
| 2:7.94 | אחדים | . |
| 2:7.95 | כבר |  |
| 2:7.96 | אתה |  |
| 2:7.97 | הולך |  |
| 2:7.98 | חשכים | , |
| 2:7.99 | הלא |  |
| 2:7.100 | תחילת |  |
| 2:7.101 | דבריו |  |
| 2:7.102 | דף |  |
| 2:7.103 | עשרים |  |
| 2:7.104 | שכל |  |
| 2:7.105 | דמיונות |  |
| 2:7.106 | אלו |  |
| 2:7.107 | על |  |
| 2:7.108 | אורות |  |
| 2:7.109 | קדושים |  |
| 2:7.110 | אשר |  |
| 2:7.111 | אין |  |
| 2:7.112 | להם |  |
| 2:7.113 | תפיסה |  |
| 2:7.114 | וצורה |  |
| 2:7.115 | ודמיון | , |
| 2:7.116 | לא |  |
| 2:7.117 | בשכל |  |
| 2:7.118 | ולא |  |
| 2:7.119 | בדעת |  |
| 2:7.120 | כלל |  |
| 2:7.121 | ועיקר | , |
| 2:7.122 | אלא |  |
| 2:7.123 | כדי |  |
| 2:7.124 | שיוכל |  |
| 2:7.125 | השכל |  |
| 2:7.126 | לקבל |  |
| 2:7.127 | אותם |  |
| 2:7.128 | הורשינו |  |
| 2:7.129 | לדמות | . |
| 2:7.130 | ובדף |  |
| 2:7.131 | ד | ' |
| 2:7.132 | ע"ב |  |
| 2:7.133 | כתב |  |
| 2:7.134 | שכל |  |
| 2:7.135 | הכנויים |  |
| 2:7.136 | והתארים |  |
| 2:7.137 | הם |  |
| 2:7.138 | בדרך |  |
| 2:7.139 | דמיון |  |
| 2:7.140 | בספירות |  |
| 2:7.141 | ואודות |  |
| 2:7.142 | קדושים |  |
| 2:7.143 | הנאצלים | , |
| 2:7.144 | אבל |  |
| 2:7.145 | בו |  |
| 2:7.146 | ית | ' |
| 2:7.147 | אין |  |
| 2:7.148 | בו |  |
| 2:7.149 | שום |  |
| 2:7.150 | דמיון |  |
| 2:7.151 | וצורה |  |
| 2:7.152 | ותפישה |  |
| 2:7.153 | במחשבה |  |
| 2:7.154 | כלל |  |
| 2:7.155 | ועיקר | . |
| 2:7.156 | ומה |  |
| 2:7.157 | שכתבת |  |
| 2:7.158 | איך |  |
| 2:7.159 | נניח |  |
| 2:7.160 | שאר |  |
| 2:7.161 | הפרצופים |  |
| 2:7.162 | ונעבוד |  |
| 2:7.163 | לזעיר |  |
| 2:7.164 | אנפין |  |
| 2:7.165 | עם |  |
| 2:7.166 | כח |  |
| 2:7.167 | שבו |  |
| 2:7.168 | א"ס |  |
| 2:7.169 | ב"ה | . |
| 2:7.170 | כבר |  |
| 2:7.171 | ביאר |  |
| 2:7.172 | הר | ' |
| 2:7.173 | כס"א |  |
| 2:7.174 | בדף |  |
| 2:7.175 | כ"ה |  |
| 2:7.176 | ע"ב |  |
| 2:7.177 | שבקראינו |  |
| 2:7.178 | לז"א |  |
| 2:7.179 | אנחנו |  |
| 2:7.180 | קורין |  |
| 2:7.181 | לכל |  |
| 2:7.182 | הפרצופים |  |
| 2:7.183 | כאחד |  |
| 2:7.184 | עם |  |
| 2:7.185 | נשמה |  |
| 2:7.186 | המסתתרת |  |
| 2:7.187 | בהם | . |
| 2:7.188 | ומה |  |
| 2:7.189 | שכתבת |  |
| 2:7.190 | למה |  |
| 2:7.191 | נניח |  |
| 2:7.192 | הפרצופים |  |
| 2:7.193 | הקרובים |  |
| 2:7.194 | לא"ס |  |
| 2:7.195 | ונקרא |  |
| 2:7.196 | לז"א | . |
| 2:7.197 | לפי |  |
| 2:7.198 | שידוע |  |
| 2:7.199 | שהמידות |  |
| 2:7.200 | הפנימיות |  |
| 2:7.201 | יש |  |
| 2:7.202 | מהם |  |
| 2:7.203 | דין |  |
| 2:7.204 | ויש |  |
| 2:7.205 | מהם |  |
| 2:7.206 | רחמים | , |
| 2:7.207 | וזה |  |
| 2:7.208 | כלול |  |
| 2:7.209 | מכולם | , |
| 2:7.210 | ולזה |  |
| 2:7.211 | רצה |  |
| 2:7.212 | המאציל |  |
| 2:7.213 | לפעול |  |
| 2:7.214 | פעולותיו |  |
| 2:7.215 | על |  |
| 2:7.216 | ידי |  |
| 2:7.217 | כלי |  |
| 2:7.218 | זה | . |
| 2:7.219 | ועד |  |
| 2:7.220 | שאתה |  |
| 2:7.221 | מקשה |  |
| 2:7.222 | על |  |
| 2:7.223 | דמיונות |  |
| 2:7.224 | אודות |  |
| 2:7.225 | עליונים |  |
| 2:7.226 | שאין |  |
| 2:7.227 | לך |  |
| 2:7.228 | רשות |  |
| 2:7.229 | להרהר |  |
| 2:7.230 | בהן |  |
| 2:7.231 | במכל |  |
| 2:7.232 | שכן |  |
| 2:7.233 | ללגלג | , |
| 2:7.234 | תקשה |  |
| 2:7.235 | על |  |
| 2:7.236 | ידך |  |
| 2:7.237 | הגופנית |  |
| 2:7.238 | שהיא |  |
| 2:7.239 | עושה |  |
| 2:7.240 | כל |  |
| 2:7.241 | המלאכות | , |
| 2:7.242 | והראש |  |
| 2:7.243 | שלמעלה |  |
| 2:7.244 | אינו |  |
| 2:7.245 | עושה |  |
| 2:7.246 | כלום | . |
| 2:7.247 | ג"כ |  |
| 2:7.248 | כח |  |
| 2:7.249 | הדבור |  |
| 2:7.250 | שהוא |  |
| 2:7.251 | פארך |  |
| 2:7.252 | על |  |
| 2:7.253 | כל |  |
| 2:7.254 | בעל |  |
| 2:7.255 | חי |  |
| 2:7.256 | למה |  |
| 2:7.257 | לא |  |
| 2:7.258 | ניתן |  |
| 2:7.259 | למעלה |  |
| 2:7.260 | מעיניך | . |
| 2:7.261 | אלא |  |
| 2:7.262 | כך |  |
| 2:7.263 | גזרה |  |
| 2:7.264 | חכמתו |  |
| 2:7.265 | ית | '. |
| 2:7.266 | ואם |  |
| 2:7.267 | יש |  |
| 2:7.268 | לך |  |
| 2:7.269 | עצה |  |
| 2:7.270 | הוגנת |  |
| 2:7.271 | תלך |  |
| 2:7.272 | אצל |  |
| 2:7.273 | בוראך |  |
| 2:7.274 | ותיעצנו | . |
| 2:7.275 | ומה |  |
| 2:7.276 | שכתבת |  |
| 2:7.277 | שדברי |  |
| 2:7.278 | כס"א |  |
| 2:7.279 | נגד |  |
| 2:7.280 | הרמב"ם |  |
| 2:7.281 | וכו | ' |
| 2:7.282 | שכתבו |  |
| 2:7.283 | שאינו |  |
| 2:7.284 | כח |  |
| 2:7.285 | בגוף | , |
| 2:7.286 | ולפי |  |
| 2:7.287 | דברי |  |
| 2:7.288 | כס"א |  |
| 2:7.289 | הוא |  |
| 2:7.290 | כח |  |
| 2:7.291 | בגוף | . |
| 2:7.292 | הלא |  |
| 2:7.293 | כל |  |
| 2:7.294 | המקובלים |  |
| 2:7.295 | פה |  |
| 2:7.296 | אחד |  |
| 2:7.297 | עונים |  |
| 2:7.298 | ואומרים |  |
| 2:7.299 | שכל |  |
| 2:7.300 | המגשים |  |
| 2:7.301 | בספירות |  |
| 2:7.302 | כמוך |  |
| 2:7.303 | הרי |  |
| 2:7.304 | זה |  |
| 2:7.305 | כופר |  |
| 2:7.306 | ואין |  |
| 2:7.307 | לו |  |
| 2:7.308 | חלק |  |
| 2:7.309 | באלהי |  |
| 2:7.310 | ישראל | . |
| 2:7.311 | ומה |  |
| 2:7.312 | שכינו |  |
| 2:7.313 | אותם |  |
| 2:7.314 | אורות |  |
| 2:7.315 | דקים |  |
| 2:7.316 | אינו |  |
| 2:7.317 | אלא |  |
| 2:7.318 | לשכך |  |
| 2:7.319 | את |  |
| 2:7.320 | האזן | ! |
| 2:7.321 | יש |  |
| 2:7.322 | לנו |  |
| 2:7.323 | לשאול |  |
| 2:7.324 | על |  |
| 2:7.325 | מה |  |
| 2:7.326 | שכתב |  |
| 2:7.327 | הרמב"ם | : |
| 2:7.328 | יכוין |  |
| 2:7.329 | לבו |  |
| 2:7.330 | כנגד |  |
| 2:7.331 | השכינה |  |
| 2:7.332 | ויתפלל | . |
| 2:7.333 | ובמורה |  |
| 2:7.334 | כ | ' |
| 2:7.335 | שהשכינה |  |
| 2:7.336 | אור |  |
| 2:7.337 | נברא | . |
| 2:7.338 | ויותר |  |
| 2:7.339 | מזה |  |
| 2:7.340 | מה |  |
| 2:7.341 | שכתב |  |
| 2:7.342 | על |  |
| 2:7.343 | כבוד |  |
| 2:7.344 | ה | ' |
| 2:7.345 | שהוא |  |
| 2:7.346 | נברא | . |
| 2:7.347 | אם |  |
| 2:7.348 | כן |  |
| 2:7.349 | כיצד |  |
| 2:7.350 | היו |  |
| 2:7.351 | ישראל |  |
| 2:7.352 | יוצאים |  |
| 2:7.353 | פתח |  |
| 2:7.354 | אהליהם |  |
| 2:7.355 | ומשתחווים |  |
| 2:7.356 | לנברא | . |
| 2:7.357 | ויותר |  |
| 2:7.358 | על |  |
| 2:7.359 | זה |  |
| 2:7.360 | מה |  |
| 2:7.361 | שכתב |  |
| 2:7.362 | שכל |  |
| 2:7.363 | צורה |  |
| 2:7.364 | שהיו |  |
| 2:7.365 | רואים |  |
| 2:7.366 | הנביאים |  |
| 2:7.367 | ומדברים |  |
| 2:7.368 | עמה |  |
| 2:7.369 | היא |  |
| 2:7.370 | אור |  |
| 2:7.371 | נברא | . |
| 2:7.372 | א"כ |  |
| 2:7.373 | אין |  |
| 2:7.374 | לנו |  |
| 2:7.375 | שום |  |
| 2:7.376 | נבואה |  |
| 2:7.377 | ע"י |  |
| 2:7.378 | הבורא | , |
| 2:7.379 | אלא |  |
| 2:7.380 | כולם |  |
| 2:7.381 | מצד |  |
| 2:7.382 | הנבראים | . |
| 2:7.383 | ואם |  |
| 2:7.384 | נאמר |  |
| 2:7.385 | שהיה |  |
| 2:7.386 | הנברא |  |
| 2:7.387 | דבק |  |
| 2:7.388 | בבורא | , |
| 2:7.389 | א"כ |  |
| 2:7.390 | לדעתך |  |
| 2:7.391 | הוא |  |
| 2:7.392 | כח |  |
| 2:7.393 | בגוף | ! |
| 2:7.394 | ותשובתך |  |
| 2:7.395 | מהרה |  |
| 2:7.396 | תצמח | , |
| 2:7.397 | ס | ' |
| 2:7.398 | מאנו |  |
| 2:7.399 | דעת |  |
| 2:7.400 | אותי |  |
| 2:7.401 | נאם |  |
| 2:7.402 | ה | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

