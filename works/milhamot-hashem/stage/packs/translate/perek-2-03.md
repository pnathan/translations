# Translation pack: milhamot-hashem perek-2-03 (Milchamot ha-Shem 2:8–2:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-03`.

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
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **אין לנו להחמיר** → "we may not be stringent".
- **דחלות** (Aramaic, once, unclear): the best reading by grammar, the others in
  `readings`.
- **ברותא** (for ברייתא) → "berota" (note).
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


## Milchamot ha-Shem 2:8

Hebrew (Wikisource): תשובה מהשואל. איותה נפשי לשמוע אמרות מוזקקות כראי מוצקות, משמי מרום בוערות. כי אני מכל דעה אביון, שואף לצל הקיקיון. ותחת אשר ענותך תרבני וממים רבים כמר מר לי תשקני. עשיתני כמלונה במקשה, ונפשי אנושה בתהלות מרמה (כברכת בלעם הרשע) דחלות. ובדברי נרגן סיימת ולאיש פרוע מוסר ודוב שכול דמית. ובאמת לא זוהי דרכה של תורה. תורתנו הקדושה דרכיה דרכי נועם. ואהבת לרעך כמוך. מה דעלך סני לחברך לא תעביד. כן אמר הלל נשיא ישראל! ריש מלין אמר, מה שאני פערתי פי לבזות חכמי ישראל והוצאתי עליהם מלין אשר לא כדת. חלילה לי לבזות שום אדם אפילו גוי, כל שכן ישראל. ולא ידעתי מקום בדברי לזה שאתה טוען שיש בהם בזיון. ואם להקשות עליהם מדברי רז"ל ולגלות מקום הסתירה, אין בזה שום איסור. וגם בתלמוד מקשים אפילו על גדולים מהם ומביאים ראיה לסתור דבריהם ואומרים תיובתא דר' פלוני. הא דר' פלוני בדותא היא (גירסת הערוך ברותא) ובסנהדרין ובברכות אמרינן: הרואה את הקשת צריך לכרוע דכתיב כמראה הקשת אשר יהיה בענן וכו' ואראה ואפול על פני. לייטי עלה במערבא וכו' מגדף בה ר' אבהו וכו'. והנה במכתבך זה תרעומת מצאתי, תשובה לא מצאתי. וכי אני הוא שקראתים לאלו הפרצופים אלוהות? הלא המקובלים הם עצמם שאמרו כן. וצא למד מכוונת הברכות והתפלה מי הוא אלהינו ומי הוא אלהי אבותינו. ובזהר שואל מאן אלהים דא כך וכך. ועוד שהרי המקובלים ביארו שכמה עולמות אלפי אלפים ורבי רבבות נאצלו ונבראו ונוצרו ונעשו, רק שלא שלחו בהם יד לבארם לרוב העלמם. כי הם אורות דקים ביותר. וכל עולם ועולם כלול מעשר ספירות. וכל ספירה וספירה כלולה כל אחת מעשר פרטי פרטות כמו שאמרנו בעולם דאצילות. וכולן בבחינת עיגולים, זה בתוך זה וזה לפנים מזה. וכן בבחינת ספירות דיושר שהם בציור אדם בעל קומה זקופה ברמ"ח איברים וכו' כל אלו העולמות כל א' מהם כולל שתי בחינות עיגולים ויושר בכלל ובפרט. וכל דברי המקובלים וביאוריהם, אינם רק בעולם דאצילות שכבר נתעבה ונתגלה יותר, ובו ביארו כל ענייני הפרצופים הכוללים שהם חמשה: אריך אנפין, ואבא ואימא, וזעיר ונוקביה. ואחר שכבר נתברר לך מדבריהם שאומרים שנתעבו ונתגלו שהם קורים את הפרצופים הנז' אלוהות באמרם אלהינו ואלהי אבותינו, בתפלה ובברכות, ומכוונים למדה ידועה בכל תיבה. לא נשאר עלי תלונתך לא בטענת גשמות ולא טענת אלוהות, כי אם עליהם, כי הם אמרו והם אמרו. וקושיא חזרה למקומה, איך נניח האין סוף עם כל הספירות אשר בעולמות העליונים עם אריך אנפין ואבא ואימא, ונעבוד לז"א ונאמר שהוא אלהינו ואנחנו עמו ועבדיו הוא עשנו ולו אנחנו. ובמכתבך זה (שאמרת לפי שידוע שהמידות הפנימיות מהם דין ומהם רחמים וכו') כבר הודית הודאה גמורה שכן היא דעתכם שאין שום עבודה ולא תפילה שייכא בעלת על כל העלות, כי אם לז"א, וכמו שכ' הרב מקדש מלך וכס"א ומהר"י לופיס בשער האמונה האמתית. ולכן צריכים אנחנו ממך ליישב ולתרץ לנו איך תסכים דעת המקובלים הלזו שכל עבודתינו ותפלותינו לז"א ונשמתו, עם תורתינו הקדושה שבכתב ושבעל פה. ואשר חרה אפך עלי ואמרת שאני מגשים בספירות, גם זה לא ידעתי לו מקום בדברי, לפי שלא אמרתי דבר שלא מצאתי בספריהם. והם אמרו שעולם האצילות יש בו גוף ונשמה ומלבוש. ולא נעלם ממך כי האור והאויר והאש גופים דקים הם וקלים. ועוד שהרי ביארו דבריהם הם עצמם ואמרו כי העולמות העליונים והפרצופים שלמעלה מאריך אנפין לא דברו בהם, רק מא"א ולמטה שכבר נתעבו ונתגלו יותר. הרי לך בהדיא מדבריהם שהספירות שמדברים בהם נתעבו ונתגלו יותר משל מעלה מהם ומודים הם בעצמם שהספירות שמדברים בהם נתעבו.

| id | word | punctuation after |
|---|---|---|
| 2:8.1 | תשובה |  |
| 2:8.2 | מהשואל | . |
| 2:8.3 | איותה |  |
| 2:8.4 | נפשי |  |
| 2:8.5 | לשמוע |  |
| 2:8.6 | אמרות |  |
| 2:8.7 | מוזקקות |  |
| 2:8.8 | כראי |  |
| 2:8.9 | מוצקות | , |
| 2:8.10 | משמי |  |
| 2:8.11 | מרום |  |
| 2:8.12 | בוערות | . |
| 2:8.13 | כי |  |
| 2:8.14 | אני |  |
| 2:8.15 | מכל |  |
| 2:8.16 | דעה |  |
| 2:8.17 | אביון | , |
| 2:8.18 | שואף |  |
| 2:8.19 | לצל |  |
| 2:8.20 | הקיקיון | . |
| 2:8.21 | ותחת |  |
| 2:8.22 | אשר |  |
| 2:8.23 | ענותך |  |
| 2:8.24 | תרבני |  |
| 2:8.25 | וממים |  |
| 2:8.26 | רבים |  |
| 2:8.27 | כמר |  |
| 2:8.28 | מר |  |
| 2:8.29 | לי |  |
| 2:8.30 | תשקני | . |
| 2:8.31 | עשיתני |  |
| 2:8.32 | כמלונה |  |
| 2:8.33 | במקשה | , |
| 2:8.34 | ונפשי |  |
| 2:8.35 | אנושה |  |
| 2:8.36 | בתהלות |  |
| 2:8.37 | מרמה |  |
| 2:8.38 | כברכת |  |
| 2:8.39 | בלעם |  |
| 2:8.40 | הרשע | ) |
| 2:8.41 | דחלות | . |
| 2:8.42 | ובדברי |  |
| 2:8.43 | נרגן |  |
| 2:8.44 | סיימת |  |
| 2:8.45 | ולאיש |  |
| 2:8.46 | פרוע |  |
| 2:8.47 | מוסר |  |
| 2:8.48 | ודוב |  |
| 2:8.49 | שכול |  |
| 2:8.50 | דמית | . |
| 2:8.51 | ובאמת |  |
| 2:8.52 | לא |  |
| 2:8.53 | זוהי |  |
| 2:8.54 | דרכה |  |
| 2:8.55 | של |  |
| 2:8.56 | תורה | . |
| 2:8.57 | תורתנו |  |
| 2:8.58 | הקדושה |  |
| 2:8.59 | דרכיה |  |
| 2:8.60 | דרכי |  |
| 2:8.61 | נועם | . |
| 2:8.62 | ואהבת |  |
| 2:8.63 | לרעך |  |
| 2:8.64 | כמוך | . |
| 2:8.65 | מה |  |
| 2:8.66 | דעלך |  |
| 2:8.67 | סני |  |
| 2:8.68 | לחברך |  |
| 2:8.69 | לא |  |
| 2:8.70 | תעביד | . |
| 2:8.71 | כן |  |
| 2:8.72 | אמר |  |
| 2:8.73 | הלל |  |
| 2:8.74 | נשיא |  |
| 2:8.75 | ישראל | ! |
| 2:8.76 | ריש |  |
| 2:8.77 | מלין |  |
| 2:8.78 | אמר | , |
| 2:8.79 | מה |  |
| 2:8.80 | שאני |  |
| 2:8.81 | פערתי |  |
| 2:8.82 | פי |  |
| 2:8.83 | לבזות |  |
| 2:8.84 | חכמי |  |
| 2:8.85 | ישראל |  |
| 2:8.86 | והוצאתי |  |
| 2:8.87 | עליהם |  |
| 2:8.88 | מלין |  |
| 2:8.89 | אשר |  |
| 2:8.90 | לא |  |
| 2:8.91 | כדת | . |
| 2:8.92 | חלילה |  |
| 2:8.93 | לי |  |
| 2:8.94 | לבזות |  |
| 2:8.95 | שום |  |
| 2:8.96 | אדם |  |
| 2:8.97 | אפילו |  |
| 2:8.98 | גוי | , |
| 2:8.99 | כל |  |
| 2:8.100 | שכן |  |
| 2:8.101 | ישראל | . |
| 2:8.102 | ולא |  |
| 2:8.103 | ידעתי |  |
| 2:8.104 | מקום |  |
| 2:8.105 | בדברי |  |
| 2:8.106 | לזה |  |
| 2:8.107 | שאתה |  |
| 2:8.108 | טוען |  |
| 2:8.109 | שיש |  |
| 2:8.110 | בהם |  |
| 2:8.111 | בזיון | . |
| 2:8.112 | ואם |  |
| 2:8.113 | להקשות |  |
| 2:8.114 | עליהם |  |
| 2:8.115 | מדברי |  |
| 2:8.116 | רז"ל |  |
| 2:8.117 | ולגלות |  |
| 2:8.118 | מקום |  |
| 2:8.119 | הסתירה | , |
| 2:8.120 | אין |  |
| 2:8.121 | בזה |  |
| 2:8.122 | שום |  |
| 2:8.123 | איסור | . |
| 2:8.124 | וגם |  |
| 2:8.125 | בתלמוד |  |
| 2:8.126 | מקשים |  |
| 2:8.127 | אפילו |  |
| 2:8.128 | על |  |
| 2:8.129 | גדולים |  |
| 2:8.130 | מהם |  |
| 2:8.131 | ומביאים |  |
| 2:8.132 | ראיה |  |
| 2:8.133 | לסתור |  |
| 2:8.134 | דבריהם |  |
| 2:8.135 | ואומרים |  |
| 2:8.136 | תיובתא |  |
| 2:8.137 | דר | ' |
| 2:8.138 | פלוני | . |
| 2:8.139 | הא |  |
| 2:8.140 | דר | ' |
| 2:8.141 | פלוני |  |
| 2:8.142 | בדותא |  |
| 2:8.143 | היא |  |
| 2:8.144 | גירסת |  |
| 2:8.145 | הערוך |  |
| 2:8.146 | ברותא | ) |
| 2:8.147 | ובסנהדרין |  |
| 2:8.148 | ובברכות |  |
| 2:8.149 | אמרינן | : |
| 2:8.150 | הרואה |  |
| 2:8.151 | את |  |
| 2:8.152 | הקשת |  |
| 2:8.153 | צריך |  |
| 2:8.154 | לכרוע |  |
| 2:8.155 | דכתיב |  |
| 2:8.156 | כמראה |  |
| 2:8.157 | הקשת |  |
| 2:8.158 | אשר |  |
| 2:8.159 | יהיה |  |
| 2:8.160 | בענן |  |
| 2:8.161 | וכו | ' |
| 2:8.162 | ואראה |  |
| 2:8.163 | ואפול |  |
| 2:8.164 | על |  |
| 2:8.165 | פני | . |
| 2:8.166 | לייטי |  |
| 2:8.167 | עלה |  |
| 2:8.168 | במערבא |  |
| 2:8.169 | וכו | ' |
| 2:8.170 | מגדף |  |
| 2:8.171 | בה |  |
| 2:8.172 | ר | ' |
| 2:8.173 | אבהו |  |
| 2:8.174 | וכו | '. |
| 2:8.175 | והנה |  |
| 2:8.176 | במכתבך |  |
| 2:8.177 | זה |  |
| 2:8.178 | תרעומת |  |
| 2:8.179 | מצאתי | , |
| 2:8.180 | תשובה |  |
| 2:8.181 | לא |  |
| 2:8.182 | מצאתי | . |
| 2:8.183 | וכי |  |
| 2:8.184 | אני |  |
| 2:8.185 | הוא |  |
| 2:8.186 | שקראתים |  |
| 2:8.187 | לאלו |  |
| 2:8.188 | הפרצופים |  |
| 2:8.189 | אלוהות | ? |
| 2:8.190 | הלא |  |
| 2:8.191 | המקובלים |  |
| 2:8.192 | הם |  |
| 2:8.193 | עצמם |  |
| 2:8.194 | שאמרו |  |
| 2:8.195 | כן | . |
| 2:8.196 | וצא |  |
| 2:8.197 | למד |  |
| 2:8.198 | מכוונת |  |
| 2:8.199 | הברכות |  |
| 2:8.200 | והתפלה |  |
| 2:8.201 | מי |  |
| 2:8.202 | הוא |  |
| 2:8.203 | אלהינו |  |
| 2:8.204 | ומי |  |
| 2:8.205 | הוא |  |
| 2:8.206 | אלהי |  |
| 2:8.207 | אבותינו | . |
| 2:8.208 | ובזהר |  |
| 2:8.209 | שואל |  |
| 2:8.210 | מאן |  |
| 2:8.211 | אלהים |  |
| 2:8.212 | דא |  |
| 2:8.213 | כך |  |
| 2:8.214 | וכך | . |
| 2:8.215 | ועוד |  |
| 2:8.216 | שהרי |  |
| 2:8.217 | המקובלים |  |
| 2:8.218 | ביארו |  |
| 2:8.219 | שכמה |  |
| 2:8.220 | עולמות |  |
| 2:8.221 | אלפי |  |
| 2:8.222 | אלפים |  |
| 2:8.223 | ורבי |  |
| 2:8.224 | רבבות |  |
| 2:8.225 | נאצלו |  |
| 2:8.226 | ונבראו |  |
| 2:8.227 | ונוצרו |  |
| 2:8.228 | ונעשו | , |
| 2:8.229 | רק |  |
| 2:8.230 | שלא |  |
| 2:8.231 | שלחו |  |
| 2:8.232 | בהם |  |
| 2:8.233 | יד |  |
| 2:8.234 | לבארם |  |
| 2:8.235 | לרוב |  |
| 2:8.236 | העלמם | . |
| 2:8.237 | כי |  |
| 2:8.238 | הם |  |
| 2:8.239 | אורות |  |
| 2:8.240 | דקים |  |
| 2:8.241 | ביותר | . |
| 2:8.242 | וכל |  |
| 2:8.243 | עולם |  |
| 2:8.244 | ועולם |  |
| 2:8.245 | כלול |  |
| 2:8.246 | מעשר |  |
| 2:8.247 | ספירות | . |
| 2:8.248 | וכל |  |
| 2:8.249 | ספירה |  |
| 2:8.250 | וספירה |  |
| 2:8.251 | כלולה |  |
| 2:8.252 | כל |  |
| 2:8.253 | אחת |  |
| 2:8.254 | מעשר |  |
| 2:8.255 | פרטי |  |
| 2:8.256 | פרטות |  |
| 2:8.257 | כמו |  |
| 2:8.258 | שאמרנו |  |
| 2:8.259 | בעולם |  |
| 2:8.260 | דאצילות | . |
| 2:8.261 | וכולן |  |
| 2:8.262 | בבחינת |  |
| 2:8.263 | עיגולים | , |
| 2:8.264 | זה |  |
| 2:8.265 | בתוך |  |
| 2:8.266 | זה |  |
| 2:8.267 | וזה |  |
| 2:8.268 | לפנים |  |
| 2:8.269 | מזה | . |
| 2:8.270 | וכן |  |
| 2:8.271 | בבחינת |  |
| 2:8.272 | ספירות |  |
| 2:8.273 | דיושר |  |
| 2:8.274 | שהם |  |
| 2:8.275 | בציור |  |
| 2:8.276 | אדם |  |
| 2:8.277 | בעל |  |
| 2:8.278 | קומה |  |
| 2:8.279 | זקופה |  |
| 2:8.280 | ברמ"ח |  |
| 2:8.281 | איברים |  |
| 2:8.282 | וכו | ' |
| 2:8.283 | כל |  |
| 2:8.284 | אלו |  |
| 2:8.285 | העולמות |  |
| 2:8.286 | כל |  |
| 2:8.287 | א | ' |
| 2:8.288 | מהם |  |
| 2:8.289 | כולל |  |
| 2:8.290 | שתי |  |
| 2:8.291 | בחינות |  |
| 2:8.292 | עיגולים |  |
| 2:8.293 | ויושר |  |
| 2:8.294 | בכלל |  |
| 2:8.295 | ובפרט | . |
| 2:8.296 | וכל |  |
| 2:8.297 | דברי |  |
| 2:8.298 | המקובלים |  |
| 2:8.299 | וביאוריהם | , |
| 2:8.300 | אינם |  |
| 2:8.301 | רק |  |
| 2:8.302 | בעולם |  |
| 2:8.303 | דאצילות |  |
| 2:8.304 | שכבר |  |
| 2:8.305 | נתעבה |  |
| 2:8.306 | ונתגלה |  |
| 2:8.307 | יותר | , |
| 2:8.308 | ובו |  |
| 2:8.309 | ביארו |  |
| 2:8.310 | כל |  |
| 2:8.311 | ענייני |  |
| 2:8.312 | הפרצופים |  |
| 2:8.313 | הכוללים |  |
| 2:8.314 | שהם |  |
| 2:8.315 | חמשה | : |
| 2:8.316 | אריך |  |
| 2:8.317 | אנפין | , |
| 2:8.318 | ואבא |  |
| 2:8.319 | ואימא | , |
| 2:8.320 | וזעיר |  |
| 2:8.321 | ונוקביה | . |
| 2:8.322 | ואחר |  |
| 2:8.323 | שכבר |  |
| 2:8.324 | נתברר |  |
| 2:8.325 | לך |  |
| 2:8.326 | מדבריהם |  |
| 2:8.327 | שאומרים |  |
| 2:8.328 | שנתעבו |  |
| 2:8.329 | ונתגלו |  |
| 2:8.330 | שהם |  |
| 2:8.331 | קורים |  |
| 2:8.332 | את |  |
| 2:8.333 | הפרצופים |  |
| 2:8.334 | הנז | ' |
| 2:8.335 | אלוהות |  |
| 2:8.336 | באמרם |  |
| 2:8.337 | אלהינו |  |
| 2:8.338 | ואלהי |  |
| 2:8.339 | אבותינו | , |
| 2:8.340 | בתפלה |  |
| 2:8.341 | ובברכות | , |
| 2:8.342 | ומכוונים |  |
| 2:8.343 | למדה |  |
| 2:8.344 | ידועה |  |
| 2:8.345 | בכל |  |
| 2:8.346 | תיבה | . |
| 2:8.347 | לא |  |
| 2:8.348 | נשאר |  |
| 2:8.349 | עלי |  |
| 2:8.350 | תלונתך |  |
| 2:8.351 | לא |  |
| 2:8.352 | בטענת |  |
| 2:8.353 | גשמות |  |
| 2:8.354 | ולא |  |
| 2:8.355 | טענת |  |
| 2:8.356 | אלוהות | , |
| 2:8.357 | כי |  |
| 2:8.358 | אם |  |
| 2:8.359 | עליהם | , |
| 2:8.360 | כי |  |
| 2:8.361 | הם |  |
| 2:8.362 | אמרו |  |
| 2:8.363 | והם |  |
| 2:8.364 | אמרו | . |
| 2:8.365 | וקושיא |  |
| 2:8.366 | חזרה |  |
| 2:8.367 | למקומה | , |
| 2:8.368 | איך |  |
| 2:8.369 | נניח |  |
| 2:8.370 | האין |  |
| 2:8.371 | סוף |  |
| 2:8.372 | עם |  |
| 2:8.373 | כל |  |
| 2:8.374 | הספירות |  |
| 2:8.375 | אשר |  |
| 2:8.376 | בעולמות |  |
| 2:8.377 | העליונים |  |
| 2:8.378 | עם |  |
| 2:8.379 | אריך |  |
| 2:8.380 | אנפין |  |
| 2:8.381 | ואבא |  |
| 2:8.382 | ואימא | , |
| 2:8.383 | ונעבוד |  |
| 2:8.384 | לז"א |  |
| 2:8.385 | ונאמר |  |
| 2:8.386 | שהוא |  |
| 2:8.387 | אלהינו |  |
| 2:8.388 | ואנחנו |  |
| 2:8.389 | עמו |  |
| 2:8.390 | ועבדיו |  |
| 2:8.391 | הוא |  |
| 2:8.392 | עשנו |  |
| 2:8.393 | ולו |  |
| 2:8.394 | אנחנו | . |
| 2:8.395 | ובמכתבך |  |
| 2:8.396 | זה |  |
| 2:8.397 | שאמרת |  |
| 2:8.398 | לפי |  |
| 2:8.399 | שידוע |  |
| 2:8.400 | שהמידות |  |
| 2:8.401 | הפנימיות |  |
| 2:8.402 | מהם |  |
| 2:8.403 | דין |  |
| 2:8.404 | ומהם |  |
| 2:8.405 | רחמים |  |
| 2:8.406 | וכו | ') |
| 2:8.407 | כבר |  |
| 2:8.408 | הודית |  |
| 2:8.409 | הודאה |  |
| 2:8.410 | גמורה |  |
| 2:8.411 | שכן |  |
| 2:8.412 | היא |  |
| 2:8.413 | דעתכם |  |
| 2:8.414 | שאין |  |
| 2:8.415 | שום |  |
| 2:8.416 | עבודה |  |
| 2:8.417 | ולא |  |
| 2:8.418 | תפילה |  |
| 2:8.419 | שייכא |  |
| 2:8.420 | בעלת |  |
| 2:8.421 | על |  |
| 2:8.422 | כל |  |
| 2:8.423 | העלות | , |
| 2:8.424 | כי |  |
| 2:8.425 | אם |  |
| 2:8.426 | לז"א | , |
| 2:8.427 | וכמו |  |
| 2:8.428 | שכ | ' |
| 2:8.429 | הרב |  |
| 2:8.430 | מקדש |  |
| 2:8.431 | מלך |  |
| 2:8.432 | וכס"א |  |
| 2:8.433 | ומהר"י |  |
| 2:8.434 | לופיס |  |
| 2:8.435 | בשער |  |
| 2:8.436 | האמונה |  |
| 2:8.437 | האמתית | . |
| 2:8.438 | ולכן |  |
| 2:8.439 | צריכים |  |
| 2:8.440 | אנחנו |  |
| 2:8.441 | ממך |  |
| 2:8.442 | ליישב |  |
| 2:8.443 | ולתרץ |  |
| 2:8.444 | לנו |  |
| 2:8.445 | איך |  |
| 2:8.446 | תסכים |  |
| 2:8.447 | דעת |  |
| 2:8.448 | המקובלים |  |
| 2:8.449 | הלזו |  |
| 2:8.450 | שכל |  |
| 2:8.451 | עבודתינו |  |
| 2:8.452 | ותפלותינו |  |
| 2:8.453 | לז"א |  |
| 2:8.454 | ונשמתו | , |
| 2:8.455 | עם |  |
| 2:8.456 | תורתינו |  |
| 2:8.457 | הקדושה |  |
| 2:8.458 | שבכתב |  |
| 2:8.459 | ושבעל |  |
| 2:8.460 | פה | . |
| 2:8.461 | ואשר |  |
| 2:8.462 | חרה |  |
| 2:8.463 | אפך |  |
| 2:8.464 | עלי |  |
| 2:8.465 | ואמרת |  |
| 2:8.466 | שאני |  |
| 2:8.467 | מגשים |  |
| 2:8.468 | בספירות | , |
| 2:8.469 | גם |  |
| 2:8.470 | זה |  |
| 2:8.471 | לא |  |
| 2:8.472 | ידעתי |  |
| 2:8.473 | לו |  |
| 2:8.474 | מקום |  |
| 2:8.475 | בדברי | , |
| 2:8.476 | לפי |  |
| 2:8.477 | שלא |  |
| 2:8.478 | אמרתי |  |
| 2:8.479 | דבר |  |
| 2:8.480 | שלא |  |
| 2:8.481 | מצאתי |  |
| 2:8.482 | בספריהם | . |
| 2:8.483 | והם |  |
| 2:8.484 | אמרו |  |
| 2:8.485 | שעולם |  |
| 2:8.486 | האצילות |  |
| 2:8.487 | יש |  |
| 2:8.488 | בו |  |
| 2:8.489 | גוף |  |
| 2:8.490 | ונשמה |  |
| 2:8.491 | ומלבוש | . |
| 2:8.492 | ולא |  |
| 2:8.493 | נעלם |  |
| 2:8.494 | ממך |  |
| 2:8.495 | כי |  |
| 2:8.496 | האור |  |
| 2:8.497 | והאויר |  |
| 2:8.498 | והאש |  |
| 2:8.499 | גופים |  |
| 2:8.500 | דקים |  |
| 2:8.501 | הם |  |
| 2:8.502 | וקלים | . |
| 2:8.503 | ועוד |  |
| 2:8.504 | שהרי |  |
| 2:8.505 | ביארו |  |
| 2:8.506 | דבריהם |  |
| 2:8.507 | הם |  |
| 2:8.508 | עצמם |  |
| 2:8.509 | ואמרו |  |
| 2:8.510 | כי |  |
| 2:8.511 | העולמות |  |
| 2:8.512 | העליונים |  |
| 2:8.513 | והפרצופים |  |
| 2:8.514 | שלמעלה |  |
| 2:8.515 | מאריך |  |
| 2:8.516 | אנפין |  |
| 2:8.517 | לא |  |
| 2:8.518 | דברו |  |
| 2:8.519 | בהם | , |
| 2:8.520 | רק |  |
| 2:8.521 | מא"א |  |
| 2:8.522 | ולמטה |  |
| 2:8.523 | שכבר |  |
| 2:8.524 | נתעבו |  |
| 2:8.525 | ונתגלו |  |
| 2:8.526 | יותר | . |
| 2:8.527 | הרי |  |
| 2:8.528 | לך |  |
| 2:8.529 | בהדיא |  |
| 2:8.530 | מדבריהם |  |
| 2:8.531 | שהספירות |  |
| 2:8.532 | שמדברים |  |
| 2:8.533 | בהם |  |
| 2:8.534 | נתעבו |  |
| 2:8.535 | ונתגלו |  |
| 2:8.536 | יותר |  |
| 2:8.537 | משל |  |
| 2:8.538 | מעלה |  |
| 2:8.539 | מהם |  |
| 2:8.540 | ומודים |  |
| 2:8.541 | הם |  |
| 2:8.542 | בעצמם |  |
| 2:8.543 | שהספירות |  |
| 2:8.544 | שמדברים |  |
| 2:8.545 | בהם |  |
| 2:8.546 | נתעבו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

