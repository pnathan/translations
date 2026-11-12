# Review pack: milhamot-hashem perek-2-03 (Milchamot ha-Shem 2:8–2:8)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-03`.

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


### Translation under review (2:8)

English: Reply from the questioner. My soul longed to hear words refined like cast mirrors, from the heights of heaven, burning. For I am poor in all knowledge, longing for the shade of the gourd. And instead of your answer enriching me, you gave me to drink many waters, bitter as gall. You made me like a booth in a cucumber field, and my soul is incurable by praises of deceit (like the blessing of the wicked Balaam), fearful. And with the words of a slanderer you concluded, and you likened to an unrestrained man who casts off discipline, and to a bereaved bear. And in truth this is not the way of the Torah. Our holy Torah, its ways are ways of pleasantness. And you shall love your neighbor as yourself. "What is hateful to you, do not do to your fellow." Thus said Hillel, Prince of Israel! He began his speech and said: as for what I opened my mouth to disparage the sages of Israel, and brought out against them words that are not according to the law. Far be it from me to disparage any person, even a gentile, all the more so an Israelite. And I know no place in my words for what you claim, that there is disparagement in them. And if one raises difficulties against them from the words of our Rabbis of blessed memory, and reveals a place of contradiction, there is no prohibition in this at all. And also in the Talmud they raise difficulties even against the great among them, and bring proof to refute their words, and say, "Refutation of Rabbi So-and-so." That of Rabbi So-and-so is a fabrication (the Arukh's reading: a baraita), and in Sanhedrin and in Berakhot we say: "The one who sees the rainbow must bow down, for it is written: 'Like the appearance of the bow that is in the cloud, etc., and I saw and fell on my face.'" "There is none against it in the West, etc. Rabbi Abahu blasphemes against it, etc." And behold, in this your letter I found a complaint; a reply I did not find. Is it I who called these partzufim divinities? Did not the Kabbalists themselves say so. Go and learn from the intention of the blessings and the prayer: who is our God, and who is the God of our fathers. And in the Zohar he asks: "Who is this Elohim, such and such?" Moreover, for the Kabbalists explained that many worlds, thousands of thousands and myriads of myriads, were emanated, created, formed and made, only they did not put a hand to explain them because of their great concealment. For they are very subtle lights. And each world is comprised of ten sefirot. And each sefirah is comprised of ten details of details, as we said in the World of Emanation. And all of them are in the aspect of circles, this within this, and this inside that. And likewise in the aspect of the sefirot of straightness, which are in the form of a man of upright stature with 248 limbs, etc. All these worlds, each one of them comprises two aspects, of circles and of straightness, in general and in particular. And all the words of the Kabbalists and their explanations are only in the World of Emanation, which has already been coarsened and revealed more; and in it they explained all the matters of the comprehensive partzufim, which are five: Arikh Anpin, and Abba, and Imma, and Ze'ir, and his Nukva. Since it has already been made clear to you from their words that they say the partzufim were coarsened and revealed, and that they call the aforementioned partzufim divinities, when they say our God and the God of our fathers in prayer and in blessings, and that they intend a known attribute in every word. Your complaint does not stand against me, neither as a charge of corporeality nor as a charge of divinity, but against them, for they said it, and they said it. And the difficulty returns to its place: how shall we posit Ein Sof together with all the sefirot that are in the upper worlds, together with Arikh Anpin and Abba and Imma; and we worship Ze'ir Anpin and say that He is our God, and we are His people and His servants; He made us, and to Him we are. In your letter (where you said, "since it is known that the inner attributes are, some of them judgment and some of them mercy, etc.") you have already made a complete admission that this is your opinion: that no worship and no prayer applies to the Cause over all the causes, but only to Ze'ir Anpin, as the Rav of Mikdash Melekh wrote, and Kisse Eliyahu, and Mahari Lopes in the Gate of the True Faith. Therefore we need you to resolve and answer for us how this opinion of the Kabbalists, that all our worship and our prayers are to Ze'ir Anpin and his soul, will agree with our holy Torah, the Written Torah and the Oral Torah. And as for your anger that burned against me, and your saying that I ascribe corporeality to the sefirot, I also know no place for this in my words, for I said nothing that I did not find in their books. They said that the World of Emanation has in it a body, a soul and a garment. And it is not hidden from you that light, air and fire are subtle and light bodies. Moreover, they themselves explained their words, and said that the upper worlds and the partzufim above Arikh Anpin they did not speak of, only from Arikh Anpin and downward, since they had already been coarsened and revealed more. Behold, it is plain to you from their words that the sefirot they speak of were coarsened and revealed more than those above them, and they themselves admit that the sefirot they speak of were coarsened.

Literal: Reply from the questioner longed my soul to hear words refined like mirrors cast from the heavens of height burning for I of all knowledge poor aspiring to the shade the gourd and instead of that your answer increased me and with waters many like gall bitter to me you gave me to drink you made me like a booth in a cucumber field and my soul is incurable by praises of deceit like the blessing of Balaam the wicked fearful and with the words of a slanderer you concluded and to a man unrestrained casting off discipline and a bear bereaved you likened and in truth not this is the way of the Torah our Torah the holy its ways are ways of pleasantness and you shall love your neighbor as yourself what to you hateful to your fellow do not do thus said Hillel prince of Israel he began his speech and said as for what I opened my mouth to disparage the sages of Israel and brought out against them words that not according to the law far be it from me to disparage any person even a gentile all the more so an Israelite and not I knew a place in my words for this that you claim that there is in them disparagement and if to raise difficulties against them from the words of our Rabbis of blessed memory and to reveal a place of contradiction there is no in this any prohibition and also in the Talmud they raise difficulties even against the great ones among them and bring proof to refute their words and say refutation of Rabbi So-and-so that Rabbi So-and-so a fabrication is the reading of the Arukh a baraita and in Sanhedrin and in Berakhot we say one who sees - the rainbow must bow down for it is written like the appearance of the bow that is in the cloud etc and I saw and fell on my face there is none against it in the West etc blasphemes against it Rabbi Abahu etc and behold in your letter this a complaint I found a reply not I found is it I the one who called them these the partzufim (faces) divinities did not the Kabbalists they themselves who said so and go and learn from the intention of the blessings and the prayer who is our God and who is the God of our fathers and in the Zohar he asks who Elohim this thus and thus moreover for the Kabbalists explained that many worlds thousands of thousands and myriads of myriads were emanated and created and formed and made only that they put in them a hand to explain them because of their great concealment for they lights subtle exceedingly and every world and world is comprised of ten sefirot and every sefirah and sefirah is comprised each one of ten details of details as we said in the world of Emanation and all of them in the aspect of circles this within this and this inside of that and likewise in the aspect of sefirot of straightness which are in the form of a man of stature upright with 248 limbs etc all these the worlds each one of them comprises two aspects of circles and of straightness in general and in particular and all the words of the Kabbalists and their explanations are only in the world of Emanation which already has been coarsened and revealed more and in it they explained all the matters of the partzufim the comprehensive which are five Arikh Anpin and Abba and Imma and Ze'ir and his Nukva since already has been made clear to you from their words that they say that they were coarsened and revealed that they call - the partzufim the aforementioned divinities when they say our God and the God of our fathers in prayer and in blessings and they intend to an attribute known in every word no remains against me your complaint not in the charge of corporeality nor the charge of divinity for but against them for they said and they said and the difficulty returns to its place how shall we posit Ein Sof together with all the sefirot that in the worlds the upper together with Arikh Anpin and Abba and Imma and we worship to Ze'ir Anpin and we say that He is our God and we are His people and His servants He made us and to Him we are in your letter this you said since it is known that the attributes inner some of them judgment and some of them mercy etc already you admitted an admission complete that this is your opinion that there is no any worship nor prayer applies to the Cause over all the causes but only to Ze'ir Anpin as wrote the Rav Mikdash Melekh and Kisse Eliyahu and Mahari Lopes in the Gate of the faith the true therefore we need we from you to resolve and to answer for us how will agree the opinion of the Kabbalists this that all our worship and our prayers to Ze'ir Anpin and his soul with our Torah the holy the Written and the Oral Torah and as for burned your anger against me and your saying that I ascribe corporeality to the sefirot also this not I know for it any place in my words for that not I said a thing that not I found in their books and they said that the world of Emanation there is in it a body and a soul and a garment and not hidden from you that light and air and fire bodies subtle they are and light moreover for they explained their words they themselves and said that the worlds the upper and the partzufim above Arikh Anpin they did not speak of them only from Arikh Anpin and downward since already were coarsened and revealed more behold to you plainly from their words that the sefirot they speak of in them were coarsened and revealed more than those of above them and they admit they themselves that the sefirot they speak of in them were coarsened

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:8.1 | Reply |  |  |  |
| 2:8.2 | from the questioner |  |  |  |
| 2:8.3 | longed |  |  |  |
| 2:8.4 | my soul |  |  |  |
| 2:8.5 | to hear |  |  |  |
| 2:8.6 | words |  |  |  |
| 2:8.7 | refined |  |  |  |
| 2:8.8 | like mirrors |  |  |  |
| 2:8.9 | cast |  |  |  |
| 2:8.10 | from the heavens of |  |  |  |
| 2:8.11 | height |  |  |  |
| 2:8.12 | burning |  |  |  |
| 2:8.13 | for |  |  |  |
| 2:8.14 | I |  |  |  |
| 2:8.15 | of all |  |  |  |
| 2:8.16 | knowledge |  |  |  |
| 2:8.17 | poor |  |  |  |
| 2:8.18 | aspiring |  |  |  |
| 2:8.19 | to the shade |  |  |  |
| 2:8.20 | the gourd |  |  |  |
| 2:8.21 | and instead of |  |  |  |
| 2:8.22 | that |  |  |  |
| 2:8.23 | your answer |  |  |  |
| 2:8.24 | increased me |  |  |  |
| 2:8.25 | and with waters |  |  |  |
| 2:8.26 | many |  |  |  |
| 2:8.27 | like gall |  |  |  |
| 2:8.28 | bitter |  |  |  |
| 2:8.29 | to me |  |  |  |
| 2:8.30 | you gave me to drink |  |  |  |
| 2:8.31 | you made me |  |  |  |
| 2:8.32 | like a booth |  |  |  |
| 2:8.33 | in a cucumber field |  |  |  |
| 2:8.34 | and my soul |  |  |  |
| 2:8.35 | is incurable |  |  |  |
| 2:8.36 | by praises |  |  |  |
| 2:8.37 | of deceit |  |  |  |
| 2:8.38 | like the blessing |  |  |  |
| 2:8.39 | of Balaam |  |  |  |
| 2:8.40 | the wicked |  |  |  |
| 2:8.41 | fearful |  |  |  |
| 2:8.42 | and with the words of |  |  |  |
| 2:8.43 | a slanderer |  |  |  |
| 2:8.44 | you concluded |  |  |  |
| 2:8.45 | and to a man |  |  |  |
| 2:8.46 | unrestrained |  |  |  |
| 2:8.47 | casting off discipline |  |  |  |
| 2:8.48 | and a bear |  |  |  |
| 2:8.49 | bereaved |  |  |  |
| 2:8.50 | you likened |  |  |  |
| 2:8.51 | and in truth |  |  |  |
| 2:8.52 | not |  |  |  |
| 2:8.53 | this is |  |  |  |
| 2:8.54 | the way |  |  |  |
| 2:8.55 | of |  |  |  |
| 2:8.56 | the Torah |  |  |  |
| 2:8.57 | our Torah |  |  |  |
| 2:8.58 | the holy |  |  |  |
| 2:8.59 | its ways |  |  |  |
| 2:8.60 | are ways of |  |  |  |
| 2:8.61 | pleasantness |  |  |  |
| 2:8.62 | and you shall love |  |  |  |
| 2:8.63 | your neighbor |  |  |  |
| 2:8.64 | as yourself |  |  |  |
| 2:8.65 | what |  |  |  |
| 2:8.66 | to you |  |  |  |
| 2:8.67 | hateful |  |  |  |
| 2:8.68 | to your fellow |  |  |  |
| 2:8.69 | do not |  |  |  |
| 2:8.70 | do |  |  |  |
| 2:8.71 | thus |  |  |  |
| 2:8.72 | said |  |  |  |
| 2:8.73 | Hillel |  |  |  |
| 2:8.74 | prince |  |  |  |
| 2:8.75 | of Israel |  |  |  |
| 2:8.76 | he began |  |  |  |
| 2:8.77 | his speech |  |  |  |
| 2:8.78 | and said |  |  |  |
| 2:8.79 | as for what |  |  |  |
| 2:8.80 | I |  |  |  |
| 2:8.81 | opened |  |  |  |
| 2:8.82 | my mouth |  |  |  |
| 2:8.83 | to disparage |  |  |  |
| 2:8.84 | the sages of |  |  |  |
| 2:8.85 | Israel |  |  |  |
| 2:8.86 | and brought out |  |  |  |
| 2:8.87 | against them |  |  |  |
| 2:8.88 | words |  |  |  |
| 2:8.89 | that |  |  |  |
| 2:8.90 | not |  |  |  |
| 2:8.91 | according to the law |  |  |  |
| 2:8.92 | far be it |  |  |  |
| 2:8.93 | from me |  |  |  |
| 2:8.94 | to disparage |  |  |  |
| 2:8.95 | any |  |  |  |
| 2:8.96 | person |  |  |  |
| 2:8.97 | even |  |  |  |
| 2:8.98 | a gentile |  |  |  |
| 2:8.99 | all the more |  |  |  |
| 2:8.100 | so |  |  |  |
| 2:8.101 | an Israelite |  |  |  |
| 2:8.102 | and not |  |  |  |
| 2:8.103 | I knew |  |  |  |
| 2:8.104 | a place |  |  |  |
| 2:8.105 | in my words |  |  |  |
| 2:8.106 | for this |  |  |  |
| 2:8.107 | that you |  |  |  |
| 2:8.108 | claim |  |  |  |
| 2:8.109 | that there is |  |  |  |
| 2:8.110 | in them |  |  |  |
| 2:8.111 | disparagement |  |  |  |
| 2:8.112 | and if |  |  |  |
| 2:8.113 | to raise difficulties |  |  |  |
| 2:8.114 | against them |  |  |  |
| 2:8.115 | from the words of |  |  |  |
| 2:8.116 | our Rabbis of blessed memory |  |  | rz"l: the Sages, with the honorific given in the decisions. |
| 2:8.117 | and to reveal |  |  |  |
| 2:8.118 | a place |  |  |  |
| 2:8.119 | of contradiction |  |  |  |
| 2:8.120 | there is no |  |  |  |
| 2:8.121 | in this |  |  |  |
| 2:8.122 | any |  |  |  |
| 2:8.123 | prohibition |  |  |  |
| 2:8.124 | and also |  |  |  |
| 2:8.125 | in the Talmud |  |  |  |
| 2:8.126 | they raise difficulties |  |  |  |
| 2:8.127 | even |  |  |  |
| 2:8.128 | against |  |  |  |
| 2:8.129 | the great ones |  |  |  |
| 2:8.130 | among them |  |  |  |
| 2:8.131 | and bring |  |  |  |
| 2:8.132 | proof |  |  |  |
| 2:8.133 | to refute |  |  |  |
| 2:8.134 | their words |  |  |  |
| 2:8.135 | and say |  |  |  |
| 2:8.136 | refutation of |  |  |  |
| 2:8.137 | Rabbi |  |  | dr' with the apostrophe; the Talmud's usual abbreviation for Rabbi. |
| 2:8.138 | So-and-so |  |  | a placeholder for an unnamed Rabbi, not a real name. |
| 2:8.139 | that |  |  |  |
| 2:8.140 | Rabbi |  |  |  |
| 2:8.141 | So-and-so |  |  |  |
| 2:8.142 | a fabrication |  |  |  |
| 2:8.143 | is |  |  |  |
| 2:8.144 | the reading of |  |  |  |
| 2:8.145 | the Arukh |  |  |  |
| 2:8.146 | a baraita |  |  |  |
| 2:8.147 | and in Sanhedrin |  |  |  |
| 2:8.148 | and in Berakhot |  |  |  |
| 2:8.149 | we say |  |  |  |
| 2:8.150 | one who sees |  |  |  |
| 2:8.151 | - |  |  |  |
| 2:8.152 | the rainbow |  |  |  |
| 2:8.153 | must |  |  |  |
| 2:8.154 | bow down |  |  |  |
| 2:8.155 | for it is written |  |  |  |
| 2:8.156 | like the appearance of |  |  |  |
| 2:8.157 | the bow |  |  |  |
| 2:8.158 | that |  |  |  |
| 2:8.159 | is |  |  |  |
| 2:8.160 | in the cloud |  |  |  |
| 2:8.161 | etc |  |  |  |
| 2:8.162 | and I saw |  |  |  |
| 2:8.163 | and fell |  |  |  |
| 2:8.164 | on |  |  |  |
| 2:8.165 | my face |  |  |  |
| 2:8.166 | there is none |  |  |  |
| 2:8.167 | against it |  |  |  |
| 2:8.168 | in the West |  |  |  |
| 2:8.169 | etc |  |  |  |
| 2:8.170 | blasphemes |  |  |  |
| 2:8.171 | against it |  |  |  |
| 2:8.172 | Rabbi |  |  |  |
| 2:8.173 | Abahu |  |  |  |
| 2:8.174 | etc |  |  |  |
| 2:8.175 | and behold |  |  |  |
| 2:8.176 | in your letter |  |  |  |
| 2:8.177 | this |  |  |  |
| 2:8.178 | a complaint |  |  |  |
| 2:8.179 | I found |  |  |  |
| 2:8.180 | a reply |  |  |  |
| 2:8.181 | not |  |  |  |
| 2:8.182 | I found |  |  |  |
| 2:8.183 | is it |  |  |  |
| 2:8.184 | I |  |  |  |
| 2:8.185 | the one |  |  |  |
| 2:8.186 | who called them |  |  |  |
| 2:8.187 | these |  |  |  |
| 2:8.188 | the partzufim (faces) |  |  | partzufim: the Kabbalistic entities, singular partzuf; literally 'faces'. |
| 2:8.189 | divinities |  |  |  |
| 2:8.190 | did not |  |  |  |
| 2:8.191 | the Kabbalists |  |  |  |
| 2:8.192 | they |  |  |  |
| 2:8.193 | themselves |  |  |  |
| 2:8.194 | who said |  |  |  |
| 2:8.195 | so |  |  |  |
| 2:8.196 | and go |  |  |  |
| 2:8.197 | and learn |  |  |  |
| 2:8.198 | from the intention of |  |  |  |
| 2:8.199 | the blessings |  |  |  |
| 2:8.200 | and the prayer |  |  |  |
| 2:8.201 | who |  |  |  |
| 2:8.202 | is |  |  |  |
| 2:8.203 | our God |  |  |  |
| 2:8.204 | and who |  |  |  |
| 2:8.205 | is |  |  |  |
| 2:8.206 | the God of |  |  |  |
| 2:8.207 | our fathers |  |  |  |
| 2:8.208 | and in the Zohar |  |  |  |
| 2:8.209 | he asks |  |  |  |
| 2:8.210 | who |  |  |  |
| 2:8.211 | Elohim |  |  |  |
| 2:8.212 | this |  |  |  |
| 2:8.213 | thus |  |  |  |
| 2:8.214 | and thus |  |  |  |
| 2:8.215 | moreover |  |  |  |
| 2:8.216 | for |  |  |  |
| 2:8.217 | the Kabbalists |  |  |  |
| 2:8.218 | explained |  |  |  |
| 2:8.219 | that many |  |  |  |
| 2:8.220 | worlds |  |  |  |
| 2:8.221 | thousands of |  |  |  |
| 2:8.222 | thousands |  |  |  |
| 2:8.223 | and myriads of |  |  |  |
| 2:8.224 | myriads |  |  |  |
| 2:8.225 | were emanated |  |  | nitsalu: emanated. The four worlds are Emanation, Creation, Formation and Making. |
| 2:8.226 | and created |  |  |  |
| 2:8.227 | and formed |  |  |  |
| 2:8.228 | and made |  |  |  |
| 2:8.229 | only |  |  |  |
| 2:8.230 | that |  |  |  |
| 2:8.231 | they put |  |  |  |
| 2:8.232 | in them |  |  |  |
| 2:8.233 | a hand |  |  |  |
| 2:8.234 | to explain them |  |  |  |
| 2:8.235 | because of |  |  |  |
| 2:8.236 | their great concealment |  |  |  |
| 2:8.237 | for |  |  |  |
| 2:8.238 | they |  |  |  |
| 2:8.239 | lights |  |  |  |
| 2:8.240 | subtle |  |  |  |
| 2:8.241 | exceedingly |  |  |  |
| 2:8.242 | and every |  |  |  |
| 2:8.243 | world |  |  |  |
| 2:8.244 | and world |  |  |  |
| 2:8.245 | is comprised |  |  |  |
| 2:8.246 | of ten |  |  |  |
| 2:8.247 | sefirot |  |  |  |
| 2:8.248 | and every |  |  |  |
| 2:8.249 | sefirah |  |  |  |
| 2:8.250 | and sefirah |  |  |  |
| 2:8.251 | is comprised |  |  |  |
| 2:8.252 | each |  |  |  |
| 2:8.253 | one |  |  |  |
| 2:8.254 | of ten |  |  |  |
| 2:8.255 | details of |  |  |  |
| 2:8.256 | details |  |  | literally 'details of details'. |
| 2:8.257 | as |  |  |  |
| 2:8.258 | we said |  |  |  |
| 2:8.259 | in the world |  |  |  |
| 2:8.260 | of Emanation |  |  | olam ha-atzilut: the World of Emanation, the first of the four worlds. |
| 2:8.261 | and all of them |  |  |  |
| 2:8.262 | in the aspect of |  |  |  |
| 2:8.263 | circles |  |  |  |
| 2:8.264 | this |  |  |  |
| 2:8.265 | within |  |  |  |
| 2:8.266 | this |  |  |  |
| 2:8.267 | and this |  |  |  |
| 2:8.268 | inside |  |  |  |
| 2:8.269 | of that |  |  |  |
| 2:8.270 | and likewise |  |  |  |
| 2:8.271 | in the aspect of |  |  |  |
| 2:8.272 | sefirot |  |  |  |
| 2:8.273 | of straightness |  |  |  |
| 2:8.274 | which are |  |  |  |
| 2:8.275 | in the form of |  |  |  |
| 2:8.276 | a man |  |  |  |
| 2:8.277 | of |  |  |  |
| 2:8.278 | stature |  |  |  |
| 2:8.279 | upright |  |  |  |
| 2:8.280 | with 248 |  |  | The letter numeral for 248, the number of the limbs of the body. |
| 2:8.281 | limbs |  |  |  |
| 2:8.282 | etc |  |  |  |
| 2:8.283 | all |  |  |  |
| 2:8.284 | these |  |  |  |
| 2:8.285 | the worlds |  |  |  |
| 2:8.286 | each |  |  |  |
| 2:8.287 | one |  |  |  |
| 2:8.288 | of them |  |  |  |
| 2:8.289 | comprises |  |  |  |
| 2:8.290 | two |  |  |  |
| 2:8.291 | aspects |  |  |  |
| 2:8.292 | of circles |  |  |  |
| 2:8.293 | and of straightness |  |  |  |
| 2:8.294 | in general |  |  |  |
| 2:8.295 | and in particular |  |  |  |
| 2:8.296 | and all |  |  |  |
| 2:8.297 | the words of |  |  |  |
| 2:8.298 | the Kabbalists |  |  |  |
| 2:8.299 | and their explanations |  |  |  |
| 2:8.300 | are |  |  |  |
| 2:8.301 | only |  |  |  |
| 2:8.302 | in the world |  |  |  |
| 2:8.303 | of Emanation |  |  |  |
| 2:8.304 | which already |  |  |  |
| 2:8.305 | has been coarsened |  |  |  |
| 2:8.306 | and revealed |  |  |  |
| 2:8.307 | more |  |  |  |
| 2:8.308 | and in it |  |  |  |
| 2:8.309 | they explained |  |  |  |
| 2:8.310 | all |  |  |  |
| 2:8.311 | the matters of |  |  |  |
| 2:8.312 | the partzufim |  |  |  |
| 2:8.313 | the comprehensive |  |  |  |
| 2:8.314 | which are |  |  |  |
| 2:8.315 | five |  |  |  |
| 2:8.316 | Arikh |  |  | Arikh Anpin: 'Long Face'. |
| 2:8.317 | Anpin |  |  |  |
| 2:8.318 | and Abba |  |  | Abba: 'Father'. |
| 2:8.319 | and Imma |  |  | Imma: 'Mother'. |
| 2:8.320 | and Ze'ir |  |  | Ze'ir Anpin: 'Small Face'. |
| 2:8.321 | and his Nukva |  |  | Nukva: the female, the partner of Ze'ir Anpin. |
| 2:8.322 | since |  |  |  |
| 2:8.323 | already |  |  |  |
| 2:8.324 | has been made clear |  |  |  |
| 2:8.325 | to you |  |  |  |
| 2:8.326 | from their words |  |  |  |
| 2:8.327 | that they say |  |  |  |
| 2:8.328 | that they were coarsened |  |  |  |
| 2:8.329 | and revealed |  |  |  |
| 2:8.330 | that they |  |  |  |
| 2:8.331 | call |  |  |  |
| 2:8.332 | - |  |  |  |
| 2:8.333 | the partzufim |  |  |  |
| 2:8.334 | the aforementioned |  |  |  |
| 2:8.335 | divinities |  |  |  |
| 2:8.336 | when they say |  |  |  |
| 2:8.337 | our God |  |  |  |
| 2:8.338 | and the God of |  |  |  |
| 2:8.339 | our fathers |  |  |  |
| 2:8.340 | in prayer |  |  |  |
| 2:8.341 | and in blessings |  |  |  |
| 2:8.342 | and they intend |  |  |  |
| 2:8.343 | to an attribute |  |  | middah: attribute, literally 'measure'. |
| 2:8.344 | known |  |  |  |
| 2:8.345 | in every |  |  |  |
| 2:8.346 | word |  |  |  |
| 2:8.347 | no |  |  |  |
| 2:8.348 | remains |  |  |  |
| 2:8.349 | against me |  |  |  |
| 2:8.350 | your complaint |  |  |  |
| 2:8.351 | not |  |  |  |
| 2:8.352 | in the charge of |  |  |  |
| 2:8.353 | corporeality |  |  |  |
| 2:8.354 | nor |  |  |  |
| 2:8.355 | the charge of |  |  |  |
| 2:8.356 | divinity |  |  |  |
| 2:8.357 | for |  |  |  |
| 2:8.358 | but |  |  |  |
| 2:8.359 | against them |  |  |  |
| 2:8.360 | for |  |  |  |
| 2:8.361 | they |  |  |  |
| 2:8.362 | said |  |  |  |
| 2:8.363 | and they |  |  |  |
| 2:8.364 | said |  |  |  |
| 2:8.365 | and the difficulty |  |  |  |
| 2:8.366 | returns |  |  |  |
| 2:8.367 | to its place |  |  |  |
| 2:8.368 | how |  |  |  |
| 2:8.369 | shall we posit |  |  |  |
| 2:8.370 | Ein |  |  | Ein Sof: literally 'without end'. |
| 2:8.371 | Sof |  |  |  |
| 2:8.372 | together with |  |  |  |
| 2:8.373 | all |  |  |  |
| 2:8.374 | the sefirot |  |  |  |
| 2:8.375 | that |  |  |  |
| 2:8.376 | in the worlds |  |  |  |
| 2:8.377 | the upper |  |  |  |
| 2:8.378 | together with |  |  |  |
| 2:8.379 | Arikh |  |  |  |
| 2:8.380 | Anpin |  |  |  |
| 2:8.381 | and Abba |  |  |  |
| 2:8.382 | and Imma |  |  |  |
| 2:8.383 | and we worship |  |  |  |
| 2:8.384 | to Ze'ir Anpin |  |  | Ze'ir Anpin: 'Small Face'; the author writes Z"A. |
| 2:8.385 | and we say |  |  |  |
| 2:8.386 | that He is |  |  |  |
| 2:8.387 | our God |  |  |  |
| 2:8.388 | and we are |  |  |  |
| 2:8.389 | His people |  |  |  |
| 2:8.390 | and His servants |  |  |  |
| 2:8.391 | He |  |  |  |
| 2:8.392 | made us |  |  |  |
| 2:8.393 | and to Him |  |  |  |
| 2:8.394 | we are |  |  |  |
| 2:8.395 | in your letter |  |  |  |
| 2:8.396 | this |  |  |  |
| 2:8.397 | you said |  |  |  |
| 2:8.398 | since |  |  |  |
| 2:8.399 | it is known |  |  |  |
| 2:8.400 | that the attributes |  |  |  |
| 2:8.401 | inner |  |  |  |
| 2:8.402 | some of them |  |  |  |
| 2:8.403 | judgment |  |  |  |
| 2:8.404 | and some of them |  |  |  |
| 2:8.405 | mercy |  |  |  |
| 2:8.406 | etc |  |  |  |
| 2:8.407 | already |  |  |  |
| 2:8.408 | you admitted |  |  |  |
| 2:8.409 | an admission |  |  |  |
| 2:8.410 | complete |  |  |  |
| 2:8.411 | that |  |  |  |
| 2:8.412 | this |  |  |  |
| 2:8.413 | is your opinion |  |  |  |
| 2:8.414 | that there is no |  |  |  |
| 2:8.415 | any |  |  |  |
| 2:8.416 | worship |  |  |  |
| 2:8.417 | nor |  |  |  |
| 2:8.418 | prayer |  |  |  |
| 2:8.419 | applies |  |  |  |
| 2:8.420 | to the Cause |  |  |  |
| 2:8.421 | over |  |  |  |
| 2:8.422 | all |  |  |  |
| 2:8.423 | the causes |  |  |  |
| 2:8.424 | but |  |  |  |
| 2:8.425 | only |  |  |  |
| 2:8.426 | to Ze'ir Anpin |  |  |  |
| 2:8.427 | as |  |  |  |
| 2:8.428 | wrote |  |  | sh' = shakhtav: 'as he wrote'. |
| 2:8.429 | the Rav |  |  |  |
| 2:8.430 | Mikdash |  |  | Mikdash Melekh: 'Temple of the King'. |
| 2:8.431 | Melekh |  |  |  |
| 2:8.432 | and Kisse Eliyahu |  |  | Kisse Eliyahu: 'Throne of Elijah'. |
| 2:8.433 | and Mahari |  |  | Mahari: the honorific acronym for 'our teacher, Rabbi'; the given name is not in the text. |
| 2:8.434 | Lopes |  |  |  |
| 2:8.435 | in the Gate |  |  | Sha'ar ha-Emunah ha-Amitit: 'The Gate of the True Faith'. |
| 2:8.436 | of the faith |  |  |  |
| 2:8.437 | the true |  |  |  |
| 2:8.438 | therefore |  |  |  |
| 2:8.439 | we need |  |  |  |
| 2:8.440 | we |  |  |  |
| 2:8.441 | from you |  |  |  |
| 2:8.442 | to resolve |  |  |  |
| 2:8.443 | and to answer |  |  |  |
| 2:8.444 | for us |  |  |  |
| 2:8.445 | how |  |  |  |
| 2:8.446 | will agree |  |  |  |
| 2:8.447 | the opinion of |  |  |  |
| 2:8.448 | the Kabbalists |  |  |  |
| 2:8.449 | this |  |  |  |
| 2:8.450 | that all |  |  |  |
| 2:8.451 | our worship |  |  |  |
| 2:8.452 | and our prayers |  |  |  |
| 2:8.453 | to Ze'ir Anpin |  |  |  |
| 2:8.454 | and his soul |  |  |  |
| 2:8.455 | with |  |  |  |
| 2:8.456 | our Torah |  |  |  |
| 2:8.457 | the holy |  |  |  |
| 2:8.458 | the Written |  |  |  |
| 2:8.459 | and the Oral |  |  |  |
| 2:8.460 | Torah |  |  |  |
| 2:8.461 | and as for |  |  |  |
| 2:8.462 | burned |  |  |  |
| 2:8.463 | your anger |  |  |  |
| 2:8.464 | against me |  |  |  |
| 2:8.465 | and your saying |  |  |  |
| 2:8.466 | that I |  |  |  |
| 2:8.467 | ascribe corporeality |  |  |  |
| 2:8.468 | to the sefirot |  |  |  |
| 2:8.469 | also |  |  |  |
| 2:8.470 | this |  |  |  |
| 2:8.471 | not |  |  |  |
| 2:8.472 | I know |  |  |  |
| 2:8.473 | for it |  |  |  |
| 2:8.474 | any place |  |  |  |
| 2:8.475 | in my words |  |  |  |
| 2:8.476 | for |  |  |  |
| 2:8.477 | that not |  |  |  |
| 2:8.478 | I said |  |  |  |
| 2:8.479 | a thing |  |  |  |
| 2:8.480 | that not |  |  |  |
| 2:8.481 | I found |  |  |  |
| 2:8.482 | in their books |  |  |  |
| 2:8.483 | and they |  |  |  |
| 2:8.484 | said |  |  |  |
| 2:8.485 | that the world |  |  |  |
| 2:8.486 | of Emanation |  |  |  |
| 2:8.487 | there is |  |  |  |
| 2:8.488 | in it |  |  |  |
| 2:8.489 | a body |  |  |  |
| 2:8.490 | and a soul |  |  |  |
| 2:8.491 | and a garment |  |  |  |
| 2:8.492 | and not |  |  |  |
| 2:8.493 | hidden |  |  |  |
| 2:8.494 | from you |  |  |  |
| 2:8.495 | that |  |  |  |
| 2:8.496 | light |  |  |  |
| 2:8.497 | and air |  |  |  |
| 2:8.498 | and fire |  |  |  |
| 2:8.499 | bodies |  |  |  |
| 2:8.500 | subtle |  |  |  |
| 2:8.501 | they are |  |  |  |
| 2:8.502 | and light |  |  |  |
| 2:8.503 | moreover |  |  |  |
| 2:8.504 | for |  |  |  |
| 2:8.505 | they explained |  |  |  |
| 2:8.506 | their words |  |  |  |
| 2:8.507 | they |  |  |  |
| 2:8.508 | themselves |  |  |  |
| 2:8.509 | and said |  |  |  |
| 2:8.510 | that |  |  |  |
| 2:8.511 | the worlds |  |  |  |
| 2:8.512 | the upper |  |  |  |
| 2:8.513 | and the partzufim |  |  |  |
| 2:8.514 | above |  |  |  |
| 2:8.515 | Arikh |  |  |  |
| 2:8.516 | Anpin |  |  |  |
| 2:8.517 | they did not |  |  |  |
| 2:8.518 | speak |  |  |  |
| 2:8.519 | of them |  |  |  |
| 2:8.520 | only |  |  |  |
| 2:8.521 | from Arikh Anpin |  |  | A"A: Arikh Anpin. The abbreviation is read as 'from Arikh Anpin'. |
| 2:8.522 | and downward |  |  |  |
| 2:8.523 | since already |  |  |  |
| 2:8.524 | were coarsened |  |  |  |
| 2:8.525 | and revealed |  |  |  |
| 2:8.526 | more |  |  |  |
| 2:8.527 | behold |  |  |  |
| 2:8.528 | to you |  |  |  |
| 2:8.529 | plainly |  |  |  |
| 2:8.530 | from their words |  |  |  |
| 2:8.531 | that the sefirot |  |  |  |
| 2:8.532 | they speak of |  |  |  |
| 2:8.533 | in them |  |  |  |
| 2:8.534 | were coarsened |  |  |  |
| 2:8.535 | and revealed |  |  |  |
| 2:8.536 | more |  |  |  |
| 2:8.537 | than those of |  |  |  |
| 2:8.538 | above |  |  |  |
| 2:8.539 | them |  |  |  |
| 2:8.540 | and they admit |  |  |  |
| 2:8.541 | they |  |  |  |
| 2:8.542 | themselves |  |  |  |
| 2:8.543 | that the sefirot |  |  |  |
| 2:8.544 | they speak of |  |  |  |
| 2:8.545 | in them |  |  |  |
| 2:8.546 | were coarsened |  |  |  |

Phrases:
- 2:8.1 2:8.2 | Reply from the questioner. | Reply from the questioner
- 2:8.3 2:8.4 2:8.5 2:8.6 2:8.7 2:8.8 2:8.9 2:8.10 2:8.11 2:8.12 | My soul longed to hear words refined like cast mirrors, from the heights of heaven, burning. | longed my soul to hear words refined like mirrors cast from the heavens of height burning
- 2:8.13 2:8.14 2:8.15 2:8.16 2:8.17 2:8.18 2:8.19 2:8.20 | For I am poor in all knowledge, longing for the shade of the gourd. | for I of all knowledge poor aspiring to the shade the gourd
- 2:8.21 2:8.22 2:8.23 2:8.24 2:8.25 2:8.26 2:8.27 2:8.28 2:8.29 2:8.30 | And instead of your answer enriching me, you gave me to drink many waters, bitter as gall. | and instead of that your answer increased me and with waters many like gall bitter to me you gave me to drink
- 2:8.31 2:8.32 2:8.33 | You made me like a booth in a cucumber field, | you made me like a booth in a cucumber field
- 2:8.34 2:8.35 2:8.36 2:8.37 2:8.38 2:8.39 2:8.40 2:8.41 | and my soul is incurable by praises of deceit (like the blessing of the wicked Balaam), fearful. | and my soul is incurable by praises of deceit like the blessing of Balaam the wicked fearful
- 2:8.42 2:8.43 2:8.44 2:8.45 2:8.46 2:8.47 2:8.48 2:8.49 2:8.50 | And with the words of a slanderer you concluded, and you likened to an unrestrained man who casts off discipline, and to a bereaved bear. | and with the words of a slanderer you concluded and to a man unrestrained casting off discipline and a bear bereaved you likened
- 2:8.51 2:8.52 2:8.53 2:8.54 2:8.55 2:8.56 | And in truth this is not the way of the Torah. | and in truth not this is the way of the Torah
- 2:8.57 2:8.58 2:8.59 2:8.60 2:8.61 | Our holy Torah, its ways are ways of pleasantness. | our Torah the holy its ways are ways of pleasantness
- 2:8.62 2:8.63 2:8.64 | And you shall love your neighbor as yourself. | and you shall love your neighbor as yourself
- 2:8.65 2:8.66 2:8.67 2:8.68 2:8.69 2:8.70 | "What is hateful to you, do not do to your fellow." | what to you hateful to your fellow do not do
- 2:8.71 2:8.72 2:8.73 2:8.74 2:8.75 | Thus said Hillel, Prince of Israel! | thus said Hillel prince of Israel
- 2:8.76 2:8.77 2:8.78 2:8.79 2:8.80 2:8.81 2:8.82 2:8.83 2:8.84 2:8.85 2:8.86 2:8.87 2:8.88 2:8.89 2:8.90 2:8.91 | He began his speech and said: as for what I opened my mouth to disparage the sages of Israel, and brought out against them words that are not according to the law. | he began his speech and said as for what I opened my mouth to disparage the sages of Israel and brought out against them words that not according to the law
- 2:8.92 2:8.93 2:8.94 2:8.95 2:8.96 2:8.97 2:8.98 2:8.99 2:8.100 2:8.101 | Far be it from me to disparage any person, even a gentile, all the more so an Israelite. | far be it from me to disparage any person even a gentile all the more so an Israelite
- 2:8.102 2:8.103 2:8.104 2:8.105 2:8.106 2:8.107 2:8.108 2:8.109 2:8.110 2:8.111 | And I know no place in my words for what you claim, that there is disparagement in them. | and not I knew a place in my words for this that you claim that there is in them disparagement
- 2:8.112 2:8.113 2:8.114 2:8.115 2:8.116 2:8.117 2:8.118 2:8.119 2:8.120 2:8.121 2:8.122 2:8.123 | And if one raises difficulties against them from the words of our Rabbis of blessed memory, and reveals a place of contradiction, there is no prohibition in this at all. | and if to raise difficulties against them from the words of our Rabbis of blessed memory and to reveal a place of contradiction there is no in this any prohibition
- 2:8.124 2:8.125 2:8.126 2:8.127 2:8.128 2:8.129 2:8.130 2:8.131 2:8.132 2:8.133 2:8.134 2:8.135 2:8.136 2:8.137 2:8.138 | And also in the Talmud they raise difficulties even against the great among them, and bring proof to refute their words, and say, "Refutation of Rabbi So-and-so." | and also in the Talmud they raise difficulties even against the great ones among them and bring proof to refute their words and say refutation of Rabbi So-and-so
- 2:8.139 2:8.140 2:8.141 2:8.142 2:8.143 2:8.144 2:8.145 2:8.146 2:8.147 2:8.148 2:8.149 | That of Rabbi So-and-so is a fabrication (the Arukh's reading: a baraita), and in Sanhedrin and in Berakhot we say: | that Rabbi So-and-so a fabrication is the reading of the Arukh a baraita and in Sanhedrin and in Berakhot we say
- 2:8.150 2:8.151 2:8.152 2:8.153 2:8.154 2:8.155 2:8.156 2:8.157 2:8.158 2:8.159 2:8.160 2:8.161 2:8.162 2:8.163 2:8.164 2:8.165 | "The one who sees the rainbow must bow down, for it is written: 'Like the appearance of the bow that is in the cloud, etc., and I saw and fell on my face.'" | one who sees - the rainbow must bow down for it is written like the appearance of the bow that is in the cloud etc and I saw and fell on my face
- 2:8.166 2:8.167 2:8.168 2:8.169 2:8.170 2:8.171 2:8.172 2:8.173 2:8.174 | "There is none against it in the West, etc. Rabbi Abahu blasphemes against it, etc." | there is none against it in the West etc blasphemes against it Rabbi Abahu etc
- 2:8.175 2:8.176 2:8.177 2:8.178 2:8.179 2:8.180 2:8.181 2:8.182 | And behold, in this your letter I found a complaint; a reply I did not find. | and behold in your letter this a complaint I found a reply not I found
- 2:8.183 2:8.184 2:8.185 2:8.186 2:8.187 2:8.188 2:8.189 | Is it I who called these partzufim divinities? | is it I the one who called them these the partzufim (faces) divinities
- 2:8.190 2:8.191 2:8.192 2:8.193 2:8.194 2:8.195 | Did not the Kabbalists themselves say so. | did not the Kabbalists they themselves who said so
- 2:8.196 2:8.197 2:8.198 2:8.199 2:8.200 2:8.201 2:8.202 2:8.203 2:8.204 2:8.205 2:8.206 2:8.207 | Go and learn from the intention of the blessings and the prayer: who is our God, and who is the God of our fathers. | and go and learn from the intention of the blessings and the prayer who is our God and who is the God of our fathers
- 2:8.208 2:8.209 2:8.210 2:8.211 2:8.212 2:8.213 2:8.214 | And in the Zohar he asks: "Who is this Elohim, such and such?" | and in the Zohar he asks who Elohim this thus and thus
- 2:8.215 2:8.216 2:8.217 2:8.218 2:8.219 2:8.220 2:8.221 2:8.222 2:8.223 2:8.224 2:8.225 2:8.226 2:8.227 2:8.228 2:8.229 2:8.230 2:8.231 2:8.232 2:8.233 2:8.234 2:8.235 2:8.236 | Moreover, for the Kabbalists explained that many worlds, thousands of thousands and myriads of myriads, were emanated, created, formed and made, only they did not put a hand to explain them because of their great concealment. | moreover for the Kabbalists explained that many worlds thousands of thousands and myriads of myriads were emanated and created and formed and made only that they put in them a hand to explain them because of their great concealment
- 2:8.237 2:8.238 2:8.239 2:8.240 2:8.241 | For they are very subtle lights. | for they lights subtle exceedingly
- 2:8.242 2:8.243 2:8.244 2:8.245 2:8.246 2:8.247 | And each world is comprised of ten sefirot. | and every world and world is comprised of ten sefirot
- 2:8.248 2:8.249 2:8.250 2:8.251 2:8.252 2:8.253 2:8.254 2:8.255 2:8.256 2:8.257 2:8.258 2:8.259 2:8.260 | And each sefirah is comprised of ten details of details, as we said in the World of Emanation. | and every sefirah and sefirah is comprised each one of ten details of details as we said in the world of Emanation
- 2:8.261 2:8.262 2:8.263 2:8.264 2:8.265 2:8.266 2:8.267 2:8.268 2:8.269 | And all of them are in the aspect of circles, this within this, and this inside that. | and all of them in the aspect of circles this within this and this inside of that
- 2:8.270 2:8.271 2:8.272 2:8.273 2:8.274 2:8.275 2:8.276 2:8.277 2:8.278 2:8.279 2:8.280 2:8.281 2:8.282 | And likewise in the aspect of the sefirot of straightness, which are in the form of a man of upright stature with 248 limbs, etc. | and likewise in the aspect of sefirot of straightness which are in the form of a man of stature upright with 248 limbs etc
- 2:8.283 2:8.284 2:8.285 2:8.286 2:8.287 2:8.288 2:8.289 2:8.290 2:8.291 2:8.292 2:8.293 2:8.294 2:8.295 | All these worlds, each one of them comprises two aspects, of circles and of straightness, in general and in particular. | all these the worlds each one of them comprises two aspects of circles and of straightness in general and in particular
- 2:8.296 2:8.297 2:8.298 2:8.299 2:8.300 2:8.301 2:8.302 2:8.303 2:8.304 2:8.305 2:8.306 2:8.307 | And all the words of the Kabbalists and their explanations are only in the World of Emanation, which has already been coarsened and revealed more; | and all the words of the Kabbalists and their explanations are only in the world of Emanation which already has been coarsened and revealed more
- 2:8.308 2:8.309 2:8.310 2:8.311 2:8.312 2:8.313 2:8.314 2:8.315 2:8.316 2:8.317 2:8.318 2:8.319 2:8.320 2:8.321 | and in it they explained all the matters of the comprehensive partzufim, which are five: Arikh Anpin, and Abba, and Imma, and Ze'ir, and his Nukva. | and in it they explained all the matters of the partzufim the comprehensive which are five Arikh Anpin and Abba and Imma and Ze'ir and his Nukva
- 2:8.322 2:8.323 2:8.324 2:8.325 2:8.326 2:8.327 2:8.328 2:8.329 2:8.330 2:8.331 2:8.332 2:8.333 2:8.334 2:8.335 | Since it has already been made clear to you from their words that they say the partzufim were coarsened and revealed, and that they call the aforementioned partzufim divinities, | since already has been made clear to you from their words that they say that they were coarsened and revealed that they call - the partzufim the aforementioned divinities
- 2:8.336 2:8.337 2:8.338 2:8.339 2:8.340 2:8.341 2:8.342 2:8.343 2:8.344 2:8.345 2:8.346 | when they say our God and the God of our fathers in prayer and in blessings, and that they intend a known attribute in every word. | when they say our God and the God of our fathers in prayer and in blessings and they intend to an attribute known in every word
- 2:8.347 2:8.348 2:8.349 2:8.350 2:8.351 2:8.352 2:8.353 2:8.354 2:8.355 2:8.356 2:8.357 2:8.358 2:8.359 2:8.360 2:8.361 2:8.362 2:8.363 2:8.364 | Your complaint does not stand against me, neither as a charge of corporeality nor as a charge of divinity, but against them, for they said it, and they said it. | no remains against me your complaint not in the charge of corporeality nor the charge of divinity for but against them for they said and they said
- 2:8.365 2:8.366 2:8.367 2:8.368 2:8.369 2:8.370 2:8.371 2:8.372 2:8.373 2:8.374 2:8.375 2:8.376 2:8.377 2:8.378 2:8.379 2:8.380 2:8.381 2:8.382 | And the difficulty returns to its place: how shall we posit Ein Sof together with all the sefirot that are in the upper worlds, together with Arikh Anpin and Abba and Imma; | and the difficulty returns to its place how shall we posit Ein Sof together with all the sefirot that in the worlds the upper together with Arikh Anpin and Abba and Imma
- 2:8.383 2:8.384 2:8.385 2:8.386 2:8.387 2:8.388 2:8.389 2:8.390 2:8.391 2:8.392 2:8.393 2:8.394 | and we worship Ze'ir Anpin and say that He is our God, and we are His people and His servants; He made us, and to Him we are. | and we worship to Ze'ir Anpin and we say that He is our God and we are His people and His servants He made us and to Him we are
- 2:8.395 2:8.396 2:8.397 2:8.398 2:8.399 2:8.400 2:8.401 2:8.402 2:8.403 2:8.404 2:8.405 2:8.406 | In your letter (where you said, "since it is known that the inner attributes are, some of them judgment and some of them mercy, etc.") | in your letter this you said since it is known that the attributes inner some of them judgment and some of them mercy etc
- 2:8.407 2:8.408 2:8.409 2:8.410 2:8.411 2:8.412 2:8.413 2:8.414 2:8.415 2:8.416 2:8.417 2:8.418 2:8.419 2:8.420 2:8.421 2:8.422 2:8.423 2:8.424 2:8.425 2:8.426 2:8.427 2:8.428 2:8.429 2:8.430 2:8.431 2:8.432 2:8.433 2:8.434 2:8.435 2:8.436 2:8.437 | you have already made a complete admission that this is your opinion: that no worship and no prayer applies to the Cause over all the causes, but only to Ze'ir Anpin, as the Rav of Mikdash Melekh wrote, and Kisse Eliyahu, and Mahari Lopes in the Gate of the True Faith. | already you admitted an admission complete that this is your opinion that there is no any worship nor prayer applies to the Cause over all the causes but only to Ze'ir Anpin as wrote the Rav Mikdash Melekh and Kisse Eliyahu and Mahari Lopes in the Gate of the faith the true
- 2:8.438 2:8.439 2:8.440 2:8.441 2:8.442 2:8.443 2:8.444 2:8.445 2:8.446 2:8.447 2:8.448 2:8.449 2:8.450 2:8.451 2:8.452 2:8.453 2:8.454 2:8.455 2:8.456 2:8.457 2:8.458 2:8.459 2:8.460 | Therefore we need you to resolve and answer for us how this opinion of the Kabbalists, that all our worship and our prayers are to Ze'ir Anpin and his soul, will agree with our holy Torah, the Written Torah and the Oral Torah. | therefore we need we from you to resolve and to answer for us how will agree the opinion of the Kabbalists this that all our worship and our prayers to Ze'ir Anpin and his soul with our Torah the holy the Written and the Oral Torah
- 2:8.461 2:8.462 2:8.463 2:8.464 2:8.465 2:8.466 2:8.467 2:8.468 2:8.469 2:8.470 2:8.471 2:8.472 2:8.473 2:8.474 2:8.475 2:8.476 2:8.477 2:8.478 2:8.479 2:8.480 2:8.481 2:8.482 | And as for your anger that burned against me, and your saying that I ascribe corporeality to the sefirot, I also know no place for this in my words, for I said nothing that I did not find in their books. | and as for burned your anger against me and your saying that I ascribe corporeality to the sefirot also this not I know for it any place in my words for that not I said a thing that not I found in their books
- 2:8.483 2:8.484 2:8.485 2:8.486 2:8.487 2:8.488 2:8.489 2:8.490 2:8.491 | They said that the World of Emanation has in it a body, a soul and a garment. | and they said that the world of Emanation there is in it a body and a soul and a garment
- 2:8.492 2:8.493 2:8.494 2:8.495 2:8.496 2:8.497 2:8.498 2:8.499 2:8.500 2:8.501 2:8.502 | And it is not hidden from you that light, air and fire are subtle and light bodies. | and not hidden from you that light and air and fire bodies subtle they are and light
- 2:8.503 2:8.504 2:8.505 2:8.506 2:8.507 2:8.508 2:8.509 2:8.510 2:8.511 2:8.512 2:8.513 2:8.514 2:8.515 2:8.516 2:8.517 2:8.518 2:8.519 2:8.520 2:8.521 2:8.522 2:8.523 2:8.524 2:8.525 2:8.526 | Moreover, they themselves explained their words, and said that the upper worlds and the partzufim above Arikh Anpin they did not speak of, only from Arikh Anpin and downward, since they had already been coarsened and revealed more. | moreover for they explained their words they themselves and said that the worlds the upper and the partzufim above Arikh Anpin they did not speak of them only from Arikh Anpin and downward since already were coarsened and revealed more
- 2:8.527 2:8.528 2:8.529 2:8.530 2:8.531 2:8.532 2:8.533 2:8.534 2:8.535 2:8.536 2:8.537 2:8.538 2:8.539 2:8.540 2:8.541 2:8.542 2:8.543 2:8.544 2:8.545 2:8.546 | Behold, it is plain to you from their words that the sefirot they speak of were coarsened and revealed more than those above them, and they themselves admit that the sefirot they speak of were coarsened. | behold to you plainly from their words that the sefirot they speak of in them were coarsened and revealed more than those of above them and they admit they themselves that the sefirot they speak of in them were coarsened

Grammar notes:
- The paragraph is the questioner's reply in chapter 2 (תשובה מהשואל): 'I' is the questioner and 'you' is the respondent, who is addressed throughout.
- In 'איותה נפשי לשמוע', the soul is the subject of the verb of longing and the infinitive is its object; 'מוזקקות' and 'מוצקות' agree with 'אמרות' (words).
- 'ובדברי נרגן סיימת ... דמית': the second-person verbs 'סיימת' and 'דמית' have no object written, and the English leaves it out as the Hebrew does.
- Aramaic: 'מה דעלך סני לחברך' is Hillel's saying; 'הא דר פלוני', 'לייטי עלה' and the Zohar's 'מאן אלהים דא כך וכך' are Aramaic (Bavli and Zohar Aramaic). 'אמרינן' ('we say') introduces the Talmudic statement.
- 'ריש מלין אמר' is an idiom ('began his speech and said'); the words after it are the questioner's denial, in the first person.

Translator's note: A letter in the correspondence of chapter 2. The questioner answers the respondent's charge that he disparaged the sages and the Kabbalists; the English keeps the questioner's voice and the respondent's 'you'.

Commentary rulings:
- apparatus: Job 37:18 (embedded allusion): 'like cast mirrors' (כראי מוצקות), from 'like a cast mirror'.
- apparatus: Jonah 4:6 (allusion): 'the shade of the gourd' (לצל הקיקיון).
- apparatus: Isa 1:8 (embedded): the verse reads 'like a booth in a vineyard' (כמלונה בכרם); the author writes 'in a cucumber field' (במקשה).
- apparatus: Hos 13:8 and Prov 17:12 (embedded): 'a bereaved bear' (דוב שכול).
- apparatus: Lev 19:18: 'You shall love your neighbor as yourself' (ואהבת לרעך כמוך).
- apparatus: Ezek 1:28 (quoted in part): 'like the appearance of the bow that is in the cloud' and 'I saw and fell on my face' (ואראה ואפול על פני).
- apparatus: Ps 100:3 (embedded): the author writes 'His servants' and 'and to Him we are'; the verse reads 'and the flock of His pasture' and 'and not we ourselves'. The English follows the author's wording.
- apparatus: The saying 'What is hateful to you, do not do to your fellow' is attributed to Hillel by the author, who gives no reference.
- apparatus: Berakhot and Arukh: the printed text reads 'fabrication' (בדותא), the Arukh's reading 'a baraita' (ברותא); the English follows the printed text and the note gives the variant.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

