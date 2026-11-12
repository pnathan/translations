# Translation pack: milhamot-hashem perek-4-20 (Milchamot ha-Shem 4:32–4:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-20`.

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
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted"; "tzimtzum" in a
  note first in a unit.
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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 4:32

Hebrew (Wikisource): לח) הנה גם מזה נתבאר לך ידידי שגגת רבותינו המקובלים החדשים שציירו ותיארו את השי"ת בכמה ציורים ותארים שונים. עגולים ויושר בכמה פרצופים שונים זה מזה, זה בתוך זה, וזה למעלה מזה, למרות קבלת רז"ל שקבלו שאסור לחשוב ולהרהר בדברים אלו שהם בכלל מה למעלה מה למטה מה לפנים ומה לאחור. ואמרו עוד המקובלים החדשים שבתחלה היה השי"ת ממלא את חלל כל העולם. ואחר כך צמצם את עצמו ונצטמק אל הצדדים סביב, ונשתנה מכמות שהיה כדי להמציא מקום וחלל לכל העולם. וזה לשון מהרח"ו בספר עץ חיים שער עיגולים ויושר ענף ב' והנה אחר הצמצום אשר אז נשאר מקום החלל ואויר פנוי וריקני באמצעות האור שלאין סוף ממש, הנה כבר היה מקום לשיוכלו להיות שם הנאצלים והנבראים והיצורים והנעשים, ואז המשיך מן אור אין סוף קו אחד ישר מן האור העגול שלו מלמעלה למטה ומשתלשל ויורד תוך החלל ההוא, וראש העליון שלהקו נמשך מן הא"ס עצמו ונוגע בו, אמנם סיום הקו ההוא למטה בסופו אינו נוגע באור האין סוף למטה, ובמקום החלל ההוא האציל וברא ויצר ועשה כל העולמות כולם, וקו זה כעין צנור דק אחד אשר דרך בו מתפשט ונמשך מימי האור העליון של אין סוף אל כל העולמות אשר במקום האויר והחלל ההוא, ונבאר עתה ענין חקירת המקובלים לדעת איך יש ראש וסוף בספירות, אמנם בהיות הקו ההוא ראשו נוגע באור אין סוף מצד העליון, וסופו אינו נמשך למטה עד מקום אין סוף הסובב תחת העולמות ואינו דבוק בו, לכן אז יצדק בו ראש וסוף, כי אם היו דרך שני קצוות מקבלים שפע האין סוף היו שני הקצוות בחינת ראשים שווים זה לזה ולא היה אז בחינת מעלה ומטה, וכן אם היה האין סוף נמשך מכל צדי המקום החלל ההוא, לא היה לא מעלה ולא מטה לא פנים ולא אחור. לא מזרח ולא מערב לא צפון ולא דרום, אך בהיות אור האין סוף נמשך דרך קו אחד וצנור דק בלבד יצדק מעלה ומטה פנים ואחור מזרח ומערב (דבריו אלה לא בהשכל דבשלמא ראש וסוף שפיר אך פנים ואחור מזרח ומערב איך יצדק?) והנה בהיות אור האין סוף נמשך בבחינת קו ישר תוך החלל, לא נמשך ונתפשט למטה תיכף, אלא היה מתפשט לאט לאט. רצוני לומר כי בתחלה התחיל קו האור להתפשט שם, ותכף בתחלת התפשטותו בסוד קו נתפשט ונמשך ונעשה כעין גלגל עגול מסביב. והעגול הזה היה בלתי דבוק אל האין סוף הסובב עליו מכל צדדיו, שאם יתדבק בו יחזור הדבר לכמות שהיה, ויתבטל באור אין סוף ולא יתראה כחו כלל, ויהיה הכל אור אין סוף לבד כבראשונה, (מדבריו אלה יש להקשות על דבריו שיזכיר להלן כי דרך הקו הזה יורד אור פנימי שהוא נשמת הספירות ואור מקיף שהוא מלבוש לספירות. למה אינו מתערב ומתדבק בגוף הספירות ויחזור עצם הספירות עם מלבושן ונשמתן כולם גולם אחד ויתבטל הכלי ויתערב באור פנימי ואור חיצון?) לכן העגול הזה סמוך אל האין סוף ובלתי מתדבק בו. וכל עיקר התקשרות ודביקות הנאצל ההוא עם אין סוף המאציל הוא ע"י הקו ההוא אשר דרך בו יורד ונמשך אור מן האין סוף המשפיע בעיגול ההוא. והאין סוף סובב ומקיף עליו מכל צדדיו ורחוק ממנו בהשואה אחת מכל צדדיו. כי גם הוא בבחינת עיגול סביב עליו. כי הוא מוכרח שהארת א"ס בנאצלים תהיה דרך הקו ההוא לבד. כי אם היה האור ההוא נמשך להם גם מכל סביבותיהם היו הנאצלים בבחינת המאציל עצמו בלתי גבול וקצבה, ולא עוד אלא אפילו הקו ההוא הוא דק מאד ולא עבה כדי שיהיה האור ההוא הנמשך אל הנאצלים במדה וקצבה אשר לסבה זאת נקראים הנאצלים עשר מדות ועשר ספירות לפי שיש להם מדה וקצבה ומספר קצוב. והנה העיגול הראשון הזה היותר דבוק עם האין סוף הוא הנקרא כתר דאדם קדמון. ואחר כך נתפשט הקו הזה ונמשך מעט וחזר להתעגל ונעשה עיגול שני תוך עיגול ראשון. וזה נקרא עיגול החכמה דאדם קדמון, עוד נתפשט יותר למטה וחזר להתעגל ונעשה עיגול שלישי תוך עיגול שני הזה ונקרא עיגול בינה דאדם קדמון, ועל דרך זה היה הולך ומתעגל עד עיגול העשירי הנקרא עיגול מלכות דאדם קדמון וכו'. אמנם מבואר ופשוט הוא שכמה מיני עולמות נאצלו ונבראו ונוצרו ונעשו אלף אלפים ורבוא רבבות, וכולם כאחד תוך המקום החלל הנז"ל ואין דבר חוצה לו. וכל עולם ועולם יש בו עשר ספירות פרטיות. וכל ספירה וספירה שבכל עולם ועולם כלולה מעשר ספירות פרטי פרטות. ועתה נבאר בחינה שנית שיש בעשר ספירות, היא בחינת היושר בצורת אדם העליון. הנה דרך הקו הנז' לעיל המתפשט מלמעלה למטה אשר ממנו מתפשטים העיגולים, גם הקו ההוא מתפשט ביושר ממעלה למטה מראש גג העליון שלעיגול העליון עד למטה מתחתית סיום כל העיגולים ממש מלמעלה למטה כלול מעשר ספירות בסוד צלם אדם ישר בעל קומה זקופה כלול מרמ"ח איברים וכו'. והנה בחינה זו השנית נקראת צלם אלהים ועליה רמז הכתוב באמרו ויברא אלהים את האדם בצלמו בצלם אלהים ברא אותו וכו' ושכמעט רוב דברי הזהר והתיקונים מתעסקים בבחינה הזאת השנית שהיא בחינת היושר עכ"ל. ובענף ג' וד' כפל עוד הדברים שאין העשר ספירות דעולם האצילות הם ראשונים וגבוהים מכל מה שנאצלו. כי כמה עולמות אבי"ע קדמו להם. ולרוב העלמם לא שלחו בהם יד להזכירם בזהר ובתיקונים אלא ברמז נפלא. ושם ביאר עוד עד היכן מגיעים רגלי אדם קדמון דיושר. ועד היכן מגיעים רגלי עתיק יומין, ורגלי אריך אנפין, ורגלי אבא ואימא, ורגלי נוקבא דיליה. ושאבא ואימא קצרי קומה, ואין שיעור קומתם רק מהגרון דאריך אנפין עד הטבור שלאריך אנפין, וכדברים האלה מבואר ג"כ בשער ההקדמות ובספר כסא אליהו ומקדש מלך פרשת בראשית עיין שם ושם.

| id | word | punctuation after |
|---|---|---|
| 4:32.1 | לח | ) |
| 4:32.2 | הנה |  |
| 4:32.3 | גם |  |
| 4:32.4 | מזה |  |
| 4:32.5 | נתבאר |  |
| 4:32.6 | לך |  |
| 4:32.7 | ידידי |  |
| 4:32.8 | שגגת |  |
| 4:32.9 | רבותינו |  |
| 4:32.10 | המקובלים |  |
| 4:32.11 | החדשים |  |
| 4:32.12 | שציירו |  |
| 4:32.13 | ותיארו |  |
| 4:32.14 | את |  |
| 4:32.15 | השי"ת |  |
| 4:32.16 | בכמה |  |
| 4:32.17 | ציורים |  |
| 4:32.18 | ותארים |  |
| 4:32.19 | שונים | . |
| 4:32.20 | עגולים |  |
| 4:32.21 | ויושר |  |
| 4:32.22 | בכמה |  |
| 4:32.23 | פרצופים |  |
| 4:32.24 | שונים |  |
| 4:32.25 | זה |  |
| 4:32.26 | מזה | , |
| 4:32.27 | זה |  |
| 4:32.28 | בתוך |  |
| 4:32.29 | זה | , |
| 4:32.30 | וזה |  |
| 4:32.31 | למעלה |  |
| 4:32.32 | מזה | , |
| 4:32.33 | למרות |  |
| 4:32.34 | קבלת |  |
| 4:32.35 | רז"ל |  |
| 4:32.36 | שקבלו |  |
| 4:32.37 | שאסור |  |
| 4:32.38 | לחשוב |  |
| 4:32.39 | ולהרהר |  |
| 4:32.40 | בדברים |  |
| 4:32.41 | אלו |  |
| 4:32.42 | שהם |  |
| 4:32.43 | בכלל |  |
| 4:32.44 | מה |  |
| 4:32.45 | למעלה |  |
| 4:32.46 | מה |  |
| 4:32.47 | למטה |  |
| 4:32.48 | מה |  |
| 4:32.49 | לפנים |  |
| 4:32.50 | ומה |  |
| 4:32.51 | לאחור | . |
| 4:32.52 | ואמרו |  |
| 4:32.53 | עוד |  |
| 4:32.54 | המקובלים |  |
| 4:32.55 | החדשים |  |
| 4:32.56 | שבתחלה |  |
| 4:32.57 | היה |  |
| 4:32.58 | השי"ת |  |
| 4:32.59 | ממלא |  |
| 4:32.60 | את |  |
| 4:32.61 | חלל |  |
| 4:32.62 | כל |  |
| 4:32.63 | העולם | . |
| 4:32.64 | ואחר |  |
| 4:32.65 | כך |  |
| 4:32.66 | צמצם |  |
| 4:32.67 | את |  |
| 4:32.68 | עצמו |  |
| 4:32.69 | ונצטמק |  |
| 4:32.70 | אל |  |
| 4:32.71 | הצדדים |  |
| 4:32.72 | סביב | , |
| 4:32.73 | ונשתנה |  |
| 4:32.74 | מכמות |  |
| 4:32.75 | שהיה |  |
| 4:32.76 | כדי |  |
| 4:32.77 | להמציא |  |
| 4:32.78 | מקום |  |
| 4:32.79 | וחלל |  |
| 4:32.80 | לכל |  |
| 4:32.81 | העולם | . |
| 4:32.82 | וזה |  |
| 4:32.83 | לשון |  |
| 4:32.84 | מהרח"ו |  |
| 4:32.85 | בספר |  |
| 4:32.86 | עץ |  |
| 4:32.87 | חיים |  |
| 4:32.88 | שער |  |
| 4:32.89 | עיגולים |  |
| 4:32.90 | ויושר |  |
| 4:32.91 | ענף |  |
| 4:32.92 | ב | ' |
| 4:32.93 | והנה |  |
| 4:32.94 | אחר |  |
| 4:32.95 | הצמצום |  |
| 4:32.96 | אשר |  |
| 4:32.97 | אז |  |
| 4:32.98 | נשאר |  |
| 4:32.99 | מקום |  |
| 4:32.100 | החלל |  |
| 4:32.101 | ואויר |  |
| 4:32.102 | פנוי |  |
| 4:32.103 | וריקני |  |
| 4:32.104 | באמצעות |  |
| 4:32.105 | האור |  |
| 4:32.106 | שלאין |  |
| 4:32.107 | סוף |  |
| 4:32.108 | ממש | , |
| 4:32.109 | הנה |  |
| 4:32.110 | כבר |  |
| 4:32.111 | היה |  |
| 4:32.112 | מקום |  |
| 4:32.113 | לשיוכלו |  |
| 4:32.114 | להיות |  |
| 4:32.115 | שם |  |
| 4:32.116 | הנאצלים |  |
| 4:32.117 | והנבראים |  |
| 4:32.118 | והיצורים |  |
| 4:32.119 | והנעשים | , |
| 4:32.120 | ואז |  |
| 4:32.121 | המשיך |  |
| 4:32.122 | מן |  |
| 4:32.123 | אור |  |
| 4:32.124 | אין |  |
| 4:32.125 | סוף |  |
| 4:32.126 | קו |  |
| 4:32.127 | אחד |  |
| 4:32.128 | ישר |  |
| 4:32.129 | מן |  |
| 4:32.130 | האור |  |
| 4:32.131 | העגול |  |
| 4:32.132 | שלו |  |
| 4:32.133 | מלמעלה |  |
| 4:32.134 | למטה |  |
| 4:32.135 | ומשתלשל |  |
| 4:32.136 | ויורד |  |
| 4:32.137 | תוך |  |
| 4:32.138 | החלל |  |
| 4:32.139 | ההוא | , |
| 4:32.140 | וראש |  |
| 4:32.141 | העליון |  |
| 4:32.142 | שלהקו |  |
| 4:32.143 | נמשך |  |
| 4:32.144 | מן |  |
| 4:32.145 | הא"ס |  |
| 4:32.146 | עצמו |  |
| 4:32.147 | ונוגע |  |
| 4:32.148 | בו | , |
| 4:32.149 | אמנם |  |
| 4:32.150 | סיום |  |
| 4:32.151 | הקו |  |
| 4:32.152 | ההוא |  |
| 4:32.153 | למטה |  |
| 4:32.154 | בסופו |  |
| 4:32.155 | אינו |  |
| 4:32.156 | נוגע |  |
| 4:32.157 | באור |  |
| 4:32.158 | האין |  |
| 4:32.159 | סוף |  |
| 4:32.160 | למטה | , |
| 4:32.161 | ובמקום |  |
| 4:32.162 | החלל |  |
| 4:32.163 | ההוא |  |
| 4:32.164 | האציל |  |
| 4:32.165 | וברא |  |
| 4:32.166 | ויצר |  |
| 4:32.167 | ועשה |  |
| 4:32.168 | כל |  |
| 4:32.169 | העולמות |  |
| 4:32.170 | כולם | , |
| 4:32.171 | וקו |  |
| 4:32.172 | זה |  |
| 4:32.173 | כעין |  |
| 4:32.174 | צנור |  |
| 4:32.175 | דק |  |
| 4:32.176 | אחד |  |
| 4:32.177 | אשר |  |
| 4:32.178 | דרך |  |
| 4:32.179 | בו |  |
| 4:32.180 | מתפשט |  |
| 4:32.181 | ונמשך |  |
| 4:32.182 | מימי |  |
| 4:32.183 | האור |  |
| 4:32.184 | העליון |  |
| 4:32.185 | של |  |
| 4:32.186 | אין |  |
| 4:32.187 | סוף |  |
| 4:32.188 | אל |  |
| 4:32.189 | כל |  |
| 4:32.190 | העולמות |  |
| 4:32.191 | אשר |  |
| 4:32.192 | במקום |  |
| 4:32.193 | האויר |  |
| 4:32.194 | והחלל |  |
| 4:32.195 | ההוא | , |
| 4:32.196 | ונבאר |  |
| 4:32.197 | עתה |  |
| 4:32.198 | ענין |  |
| 4:32.199 | חקירת |  |
| 4:32.200 | המקובלים |  |
| 4:32.201 | לדעת |  |
| 4:32.202 | איך |  |
| 4:32.203 | יש |  |
| 4:32.204 | ראש |  |
| 4:32.205 | וסוף |  |
| 4:32.206 | בספירות | , |
| 4:32.207 | אמנם |  |
| 4:32.208 | בהיות |  |
| 4:32.209 | הקו |  |
| 4:32.210 | ההוא |  |
| 4:32.211 | ראשו |  |
| 4:32.212 | נוגע |  |
| 4:32.213 | באור |  |
| 4:32.214 | אין |  |
| 4:32.215 | סוף |  |
| 4:32.216 | מצד |  |
| 4:32.217 | העליון | , |
| 4:32.218 | וסופו |  |
| 4:32.219 | אינו |  |
| 4:32.220 | נמשך |  |
| 4:32.221 | למטה |  |
| 4:32.222 | עד |  |
| 4:32.223 | מקום |  |
| 4:32.224 | אין |  |
| 4:32.225 | סוף |  |
| 4:32.226 | הסובב |  |
| 4:32.227 | תחת |  |
| 4:32.228 | העולמות |  |
| 4:32.229 | ואינו |  |
| 4:32.230 | דבוק |  |
| 4:32.231 | בו | , |
| 4:32.232 | לכן |  |
| 4:32.233 | אז |  |
| 4:32.234 | יצדק |  |
| 4:32.235 | בו |  |
| 4:32.236 | ראש |  |
| 4:32.237 | וסוף | , |
| 4:32.238 | כי |  |
| 4:32.239 | אם |  |
| 4:32.240 | היו |  |
| 4:32.241 | דרך |  |
| 4:32.242 | שני |  |
| 4:32.243 | קצוות |  |
| 4:32.244 | מקבלים |  |
| 4:32.245 | שפע |  |
| 4:32.246 | האין |  |
| 4:32.247 | סוף |  |
| 4:32.248 | היו |  |
| 4:32.249 | שני |  |
| 4:32.250 | הקצוות |  |
| 4:32.251 | בחינת |  |
| 4:32.252 | ראשים |  |
| 4:32.253 | שווים |  |
| 4:32.254 | זה |  |
| 4:32.255 | לזה |  |
| 4:32.256 | ולא |  |
| 4:32.257 | היה |  |
| 4:32.258 | אז |  |
| 4:32.259 | בחינת |  |
| 4:32.260 | מעלה |  |
| 4:32.261 | ומטה | , |
| 4:32.262 | וכן |  |
| 4:32.263 | אם |  |
| 4:32.264 | היה |  |
| 4:32.265 | האין |  |
| 4:32.266 | סוף |  |
| 4:32.267 | נמשך |  |
| 4:32.268 | מכל |  |
| 4:32.269 | צדי |  |
| 4:32.270 | המקום |  |
| 4:32.271 | החלל |  |
| 4:32.272 | ההוא | , |
| 4:32.273 | לא |  |
| 4:32.274 | היה |  |
| 4:32.275 | לא |  |
| 4:32.276 | מעלה |  |
| 4:32.277 | ולא |  |
| 4:32.278 | מטה |  |
| 4:32.279 | לא |  |
| 4:32.280 | פנים |  |
| 4:32.281 | ולא |  |
| 4:32.282 | אחור | . |
| 4:32.283 | לא |  |
| 4:32.284 | מזרח |  |
| 4:32.285 | ולא |  |
| 4:32.286 | מערב |  |
| 4:32.287 | לא |  |
| 4:32.288 | צפון |  |
| 4:32.289 | ולא |  |
| 4:32.290 | דרום | , |
| 4:32.291 | אך |  |
| 4:32.292 | בהיות |  |
| 4:32.293 | אור |  |
| 4:32.294 | האין |  |
| 4:32.295 | סוף |  |
| 4:32.296 | נמשך |  |
| 4:32.297 | דרך |  |
| 4:32.298 | קו |  |
| 4:32.299 | אחד |  |
| 4:32.300 | וצנור |  |
| 4:32.301 | דק |  |
| 4:32.302 | בלבד |  |
| 4:32.303 | יצדק |  |
| 4:32.304 | מעלה |  |
| 4:32.305 | ומטה |  |
| 4:32.306 | פנים |  |
| 4:32.307 | ואחור |  |
| 4:32.308 | מזרח |  |
| 4:32.309 | ומערב |  |
| 4:32.310 | דבריו |  |
| 4:32.311 | אלה |  |
| 4:32.312 | לא |  |
| 4:32.313 | בהשכל |  |
| 4:32.314 | דבשלמא |  |
| 4:32.315 | ראש |  |
| 4:32.316 | וסוף |  |
| 4:32.317 | שפיר |  |
| 4:32.318 | אך |  |
| 4:32.319 | פנים |  |
| 4:32.320 | ואחור |  |
| 4:32.321 | מזרח |  |
| 4:32.322 | ומערב |  |
| 4:32.323 | איך |  |
| 4:32.324 | יצדק | ?) |
| 4:32.325 | והנה |  |
| 4:32.326 | בהיות |  |
| 4:32.327 | אור |  |
| 4:32.328 | האין |  |
| 4:32.329 | סוף |  |
| 4:32.330 | נמשך |  |
| 4:32.331 | בבחינת |  |
| 4:32.332 | קו |  |
| 4:32.333 | ישר |  |
| 4:32.334 | תוך |  |
| 4:32.335 | החלל | , |
| 4:32.336 | לא |  |
| 4:32.337 | נמשך |  |
| 4:32.338 | ונתפשט |  |
| 4:32.339 | למטה |  |
| 4:32.340 | תיכף | , |
| 4:32.341 | אלא |  |
| 4:32.342 | היה |  |
| 4:32.343 | מתפשט |  |
| 4:32.344 | לאט |  |
| 4:32.345 | לאט | . |
| 4:32.346 | רצוני |  |
| 4:32.347 | לומר |  |
| 4:32.348 | כי |  |
| 4:32.349 | בתחלה |  |
| 4:32.350 | התחיל |  |
| 4:32.351 | קו |  |
| 4:32.352 | האור |  |
| 4:32.353 | להתפשט |  |
| 4:32.354 | שם | , |
| 4:32.355 | ותכף |  |
| 4:32.356 | בתחלת |  |
| 4:32.357 | התפשטותו |  |
| 4:32.358 | בסוד |  |
| 4:32.359 | קו |  |
| 4:32.360 | נתפשט |  |
| 4:32.361 | ונמשך |  |
| 4:32.362 | ונעשה |  |
| 4:32.363 | כעין |  |
| 4:32.364 | גלגל |  |
| 4:32.365 | עגול |  |
| 4:32.366 | מסביב | . |
| 4:32.367 | והעגול |  |
| 4:32.368 | הזה |  |
| 4:32.369 | היה |  |
| 4:32.370 | בלתי |  |
| 4:32.371 | דבוק |  |
| 4:32.372 | אל |  |
| 4:32.373 | האין |  |
| 4:32.374 | סוף |  |
| 4:32.375 | הסובב |  |
| 4:32.376 | עליו |  |
| 4:32.377 | מכל |  |
| 4:32.378 | צדדיו | , |
| 4:32.379 | שאם |  |
| 4:32.380 | יתדבק |  |
| 4:32.381 | בו |  |
| 4:32.382 | יחזור |  |
| 4:32.383 | הדבר |  |
| 4:32.384 | לכמות |  |
| 4:32.385 | שהיה | , |
| 4:32.386 | ויתבטל |  |
| 4:32.387 | באור |  |
| 4:32.388 | אין |  |
| 4:32.389 | סוף |  |
| 4:32.390 | ולא |  |
| 4:32.391 | יתראה |  |
| 4:32.392 | כחו |  |
| 4:32.393 | כלל | , |
| 4:32.394 | ויהיה |  |
| 4:32.395 | הכל |  |
| 4:32.396 | אור |  |
| 4:32.397 | אין |  |
| 4:32.398 | סוף |  |
| 4:32.399 | לבד |  |
| 4:32.400 | כבראשונה | , |
| 4:32.401 | מדבריו |  |
| 4:32.402 | אלה |  |
| 4:32.403 | יש |  |
| 4:32.404 | להקשות |  |
| 4:32.405 | על |  |
| 4:32.406 | דבריו |  |
| 4:32.407 | שיזכיר |  |
| 4:32.408 | להלן |  |
| 4:32.409 | כי |  |
| 4:32.410 | דרך |  |
| 4:32.411 | הקו |  |
| 4:32.412 | הזה |  |
| 4:32.413 | יורד |  |
| 4:32.414 | אור |  |
| 4:32.415 | פנימי |  |
| 4:32.416 | שהוא |  |
| 4:32.417 | נשמת |  |
| 4:32.418 | הספירות |  |
| 4:32.419 | ואור |  |
| 4:32.420 | מקיף |  |
| 4:32.421 | שהוא |  |
| 4:32.422 | מלבוש |  |
| 4:32.423 | לספירות | . |
| 4:32.424 | למה |  |
| 4:32.425 | אינו |  |
| 4:32.426 | מתערב |  |
| 4:32.427 | ומתדבק |  |
| 4:32.428 | בגוף |  |
| 4:32.429 | הספירות |  |
| 4:32.430 | ויחזור |  |
| 4:32.431 | עצם |  |
| 4:32.432 | הספירות |  |
| 4:32.433 | עם |  |
| 4:32.434 | מלבושן |  |
| 4:32.435 | ונשמתן |  |
| 4:32.436 | כולם |  |
| 4:32.437 | גולם |  |
| 4:32.438 | אחד |  |
| 4:32.439 | ויתבטל |  |
| 4:32.440 | הכלי |  |
| 4:32.441 | ויתערב |  |
| 4:32.442 | באור |  |
| 4:32.443 | פנימי |  |
| 4:32.444 | ואור |  |
| 4:32.445 | חיצון | ?) |
| 4:32.446 | לכן |  |
| 4:32.447 | העגול |  |
| 4:32.448 | הזה |  |
| 4:32.449 | סמוך |  |
| 4:32.450 | אל |  |
| 4:32.451 | האין |  |
| 4:32.452 | סוף |  |
| 4:32.453 | ובלתי |  |
| 4:32.454 | מתדבק |  |
| 4:32.455 | בו | . |
| 4:32.456 | וכל |  |
| 4:32.457 | עיקר |  |
| 4:32.458 | התקשרות |  |
| 4:32.459 | ודביקות |  |
| 4:32.460 | הנאצל |  |
| 4:32.461 | ההוא |  |
| 4:32.462 | עם |  |
| 4:32.463 | אין |  |
| 4:32.464 | סוף |  |
| 4:32.465 | המאציל |  |
| 4:32.466 | הוא |  |
| 4:32.467 | ע"י |  |
| 4:32.468 | הקו |  |
| 4:32.469 | ההוא |  |
| 4:32.470 | אשר |  |
| 4:32.471 | דרך |  |
| 4:32.472 | בו |  |
| 4:32.473 | יורד |  |
| 4:32.474 | ונמשך |  |
| 4:32.475 | אור |  |
| 4:32.476 | מן |  |
| 4:32.477 | האין |  |
| 4:32.478 | סוף |  |
| 4:32.479 | המשפיע |  |
| 4:32.480 | בעיגול |  |
| 4:32.481 | ההוא | . |
| 4:32.482 | והאין |  |
| 4:32.483 | סוף |  |
| 4:32.484 | סובב |  |
| 4:32.485 | ומקיף |  |
| 4:32.486 | עליו |  |
| 4:32.487 | מכל |  |
| 4:32.488 | צדדיו |  |
| 4:32.489 | ורחוק |  |
| 4:32.490 | ממנו |  |
| 4:32.491 | בהשואה |  |
| 4:32.492 | אחת |  |
| 4:32.493 | מכל |  |
| 4:32.494 | צדדיו | . |
| 4:32.495 | כי |  |
| 4:32.496 | גם |  |
| 4:32.497 | הוא |  |
| 4:32.498 | בבחינת |  |
| 4:32.499 | עיגול |  |
| 4:32.500 | סביב |  |
| 4:32.501 | עליו | . |
| 4:32.502 | כי |  |
| 4:32.503 | הוא |  |
| 4:32.504 | מוכרח |  |
| 4:32.505 | שהארת |  |
| 4:32.506 | א"ס |  |
| 4:32.507 | בנאצלים |  |
| 4:32.508 | תהיה |  |
| 4:32.509 | דרך |  |
| 4:32.510 | הקו |  |
| 4:32.511 | ההוא |  |
| 4:32.512 | לבד | . |
| 4:32.513 | כי |  |
| 4:32.514 | אם |  |
| 4:32.515 | היה |  |
| 4:32.516 | האור |  |
| 4:32.517 | ההוא |  |
| 4:32.518 | נמשך |  |
| 4:32.519 | להם |  |
| 4:32.520 | גם |  |
| 4:32.521 | מכל |  |
| 4:32.522 | סביבותיהם |  |
| 4:32.523 | היו |  |
| 4:32.524 | הנאצלים |  |
| 4:32.525 | בבחינת |  |
| 4:32.526 | המאציל |  |
| 4:32.527 | עצמו |  |
| 4:32.528 | בלתי |  |
| 4:32.529 | גבול |  |
| 4:32.530 | וקצבה | , |
| 4:32.531 | ולא |  |
| 4:32.532 | עוד |  |
| 4:32.533 | אלא |  |
| 4:32.534 | אפילו |  |
| 4:32.535 | הקו |  |
| 4:32.536 | ההוא |  |
| 4:32.537 | הוא |  |
| 4:32.538 | דק |  |
| 4:32.539 | מאד |  |
| 4:32.540 | ולא |  |
| 4:32.541 | עבה |  |
| 4:32.542 | כדי |  |
| 4:32.543 | שיהיה |  |
| 4:32.544 | האור |  |
| 4:32.545 | ההוא |  |
| 4:32.546 | הנמשך |  |
| 4:32.547 | אל |  |
| 4:32.548 | הנאצלים |  |
| 4:32.549 | במדה |  |
| 4:32.550 | וקצבה |  |
| 4:32.551 | אשר |  |
| 4:32.552 | לסבה |  |
| 4:32.553 | זאת |  |
| 4:32.554 | נקראים |  |
| 4:32.555 | הנאצלים |  |
| 4:32.556 | עשר |  |
| 4:32.557 | מדות |  |
| 4:32.558 | ועשר |  |
| 4:32.559 | ספירות |  |
| 4:32.560 | לפי |  |
| 4:32.561 | שיש |  |
| 4:32.562 | להם |  |
| 4:32.563 | מדה |  |
| 4:32.564 | וקצבה |  |
| 4:32.565 | ומספר |  |
| 4:32.566 | קצוב | . |
| 4:32.567 | והנה |  |
| 4:32.568 | העיגול |  |
| 4:32.569 | הראשון |  |
| 4:32.570 | הזה |  |
| 4:32.571 | היותר |  |
| 4:32.572 | דבוק |  |
| 4:32.573 | עם |  |
| 4:32.574 | האין |  |
| 4:32.575 | סוף |  |
| 4:32.576 | הוא |  |
| 4:32.577 | הנקרא |  |
| 4:32.578 | כתר |  |
| 4:32.579 | דאדם |  |
| 4:32.580 | קדמון | . |
| 4:32.581 | ואחר |  |
| 4:32.582 | כך |  |
| 4:32.583 | נתפשט |  |
| 4:32.584 | הקו |  |
| 4:32.585 | הזה |  |
| 4:32.586 | ונמשך |  |
| 4:32.587 | מעט |  |
| 4:32.588 | וחזר |  |
| 4:32.589 | להתעגל |  |
| 4:32.590 | ונעשה |  |
| 4:32.591 | עיגול |  |
| 4:32.592 | שני |  |
| 4:32.593 | תוך |  |
| 4:32.594 | עיגול |  |
| 4:32.595 | ראשון | . |
| 4:32.596 | וזה |  |
| 4:32.597 | נקרא |  |
| 4:32.598 | עיגול |  |
| 4:32.599 | החכמה |  |
| 4:32.600 | דאדם |  |
| 4:32.601 | קדמון | , |
| 4:32.602 | עוד |  |
| 4:32.603 | נתפשט |  |
| 4:32.604 | יותר |  |
| 4:32.605 | למטה |  |
| 4:32.606 | וחזר |  |
| 4:32.607 | להתעגל |  |
| 4:32.608 | ונעשה |  |
| 4:32.609 | עיגול |  |
| 4:32.610 | שלישי |  |
| 4:32.611 | תוך |  |
| 4:32.612 | עיגול |  |
| 4:32.613 | שני |  |
| 4:32.614 | הזה |  |
| 4:32.615 | ונקרא |  |
| 4:32.616 | עיגול |  |
| 4:32.617 | בינה |  |
| 4:32.618 | דאדם |  |
| 4:32.619 | קדמון | , |
| 4:32.620 | ועל |  |
| 4:32.621 | דרך |  |
| 4:32.622 | זה |  |
| 4:32.623 | היה |  |
| 4:32.624 | הולך |  |
| 4:32.625 | ומתעגל |  |
| 4:32.626 | עד |  |
| 4:32.627 | עיגול |  |
| 4:32.628 | העשירי |  |
| 4:32.629 | הנקרא |  |
| 4:32.630 | עיגול |  |
| 4:32.631 | מלכות |  |
| 4:32.632 | דאדם |  |
| 4:32.633 | קדמון |  |
| 4:32.634 | וכו | '. |
| 4:32.635 | אמנם |  |
| 4:32.636 | מבואר |  |
| 4:32.637 | ופשוט |  |
| 4:32.638 | הוא |  |
| 4:32.639 | שכמה |  |
| 4:32.640 | מיני |  |
| 4:32.641 | עולמות |  |
| 4:32.642 | נאצלו |  |
| 4:32.643 | ונבראו |  |
| 4:32.644 | ונוצרו |  |
| 4:32.645 | ונעשו |  |
| 4:32.646 | אלף |  |
| 4:32.647 | אלפים |  |
| 4:32.648 | ורבוא |  |
| 4:32.649 | רבבות | , |
| 4:32.650 | וכולם |  |
| 4:32.651 | כאחד |  |
| 4:32.652 | תוך |  |
| 4:32.653 | המקום |  |
| 4:32.654 | החלל |  |
| 4:32.655 | הנז"ל |  |
| 4:32.656 | ואין |  |
| 4:32.657 | דבר |  |
| 4:32.658 | חוצה |  |
| 4:32.659 | לו | . |
| 4:32.660 | וכל |  |
| 4:32.661 | עולם |  |
| 4:32.662 | ועולם |  |
| 4:32.663 | יש |  |
| 4:32.664 | בו |  |
| 4:32.665 | עשר |  |
| 4:32.666 | ספירות |  |
| 4:32.667 | פרטיות | . |
| 4:32.668 | וכל |  |
| 4:32.669 | ספירה |  |
| 4:32.670 | וספירה |  |
| 4:32.671 | שבכל |  |
| 4:32.672 | עולם |  |
| 4:32.673 | ועולם |  |
| 4:32.674 | כלולה |  |
| 4:32.675 | מעשר |  |
| 4:32.676 | ספירות |  |
| 4:32.677 | פרטי |  |
| 4:32.678 | פרטות | . |
| 4:32.679 | ועתה |  |
| 4:32.680 | נבאר |  |
| 4:32.681 | בחינה |  |
| 4:32.682 | שנית |  |
| 4:32.683 | שיש |  |
| 4:32.684 | בעשר |  |
| 4:32.685 | ספירות | , |
| 4:32.686 | היא |  |
| 4:32.687 | בחינת |  |
| 4:32.688 | היושר |  |
| 4:32.689 | בצורת |  |
| 4:32.690 | אדם |  |
| 4:32.691 | העליון | . |
| 4:32.692 | הנה |  |
| 4:32.693 | דרך |  |
| 4:32.694 | הקו |  |
| 4:32.695 | הנז | ' |
| 4:32.696 | לעיל |  |
| 4:32.697 | המתפשט |  |
| 4:32.698 | מלמעלה |  |
| 4:32.699 | למטה |  |
| 4:32.700 | אשר |  |
| 4:32.701 | ממנו |  |
| 4:32.702 | מתפשטים |  |
| 4:32.703 | העיגולים | , |
| 4:32.704 | גם |  |
| 4:32.705 | הקו |  |
| 4:32.706 | ההוא |  |
| 4:32.707 | מתפשט |  |
| 4:32.708 | ביושר |  |
| 4:32.709 | ממעלה |  |
| 4:32.710 | למטה |  |
| 4:32.711 | מראש |  |
| 4:32.712 | גג |  |
| 4:32.713 | העליון |  |
| 4:32.714 | שלעיגול |  |
| 4:32.715 | העליון |  |
| 4:32.716 | עד |  |
| 4:32.717 | למטה |  |
| 4:32.718 | מתחתית |  |
| 4:32.719 | סיום |  |
| 4:32.720 | כל |  |
| 4:32.721 | העיגולים |  |
| 4:32.722 | ממש |  |
| 4:32.723 | מלמעלה |  |
| 4:32.724 | למטה |  |
| 4:32.725 | כלול |  |
| 4:32.726 | מעשר |  |
| 4:32.727 | ספירות |  |
| 4:32.728 | בסוד |  |
| 4:32.729 | צלם |  |
| 4:32.730 | אדם |  |
| 4:32.731 | ישר |  |
| 4:32.732 | בעל |  |
| 4:32.733 | קומה |  |
| 4:32.734 | זקופה |  |
| 4:32.735 | כלול |  |
| 4:32.736 | מרמ"ח |  |
| 4:32.737 | איברים |  |
| 4:32.738 | וכו | '. |
| 4:32.739 | והנה |  |
| 4:32.740 | בחינה |  |
| 4:32.741 | זו |  |
| 4:32.742 | השנית |  |
| 4:32.743 | נקראת |  |
| 4:32.744 | צלם |  |
| 4:32.745 | אלהים |  |
| 4:32.746 | ועליה |  |
| 4:32.747 | רמז |  |
| 4:32.748 | הכתוב |  |
| 4:32.749 | באמרו |  |
| 4:32.750 | ויברא |  |
| 4:32.751 | אלהים |  |
| 4:32.752 | את |  |
| 4:32.753 | האדם |  |
| 4:32.754 | בצלמו |  |
| 4:32.755 | בצלם |  |
| 4:32.756 | אלהים |  |
| 4:32.757 | ברא |  |
| 4:32.758 | אותו |  |
| 4:32.759 | וכו | ' |
| 4:32.760 | ושכמעט |  |
| 4:32.761 | רוב |  |
| 4:32.762 | דברי |  |
| 4:32.763 | הזהר |  |
| 4:32.764 | והתיקונים |  |
| 4:32.765 | מתעסקים |  |
| 4:32.766 | בבחינה |  |
| 4:32.767 | הזאת |  |
| 4:32.768 | השנית |  |
| 4:32.769 | שהיא |  |
| 4:32.770 | בחינת |  |
| 4:32.771 | היושר |  |
| 4:32.772 | עכ"ל | . |
| 4:32.773 | ובענף |  |
| 4:32.774 | ג | ' |
| 4:32.775 | וד | ' |
| 4:32.776 | כפל |  |
| 4:32.777 | עוד |  |
| 4:32.778 | הדברים |  |
| 4:32.779 | שאין |  |
| 4:32.780 | העשר |  |
| 4:32.781 | ספירות |  |
| 4:32.782 | דעולם |  |
| 4:32.783 | האצילות |  |
| 4:32.784 | הם |  |
| 4:32.785 | ראשונים |  |
| 4:32.786 | וגבוהים |  |
| 4:32.787 | מכל |  |
| 4:32.788 | מה |  |
| 4:32.789 | שנאצלו | . |
| 4:32.790 | כי |  |
| 4:32.791 | כמה |  |
| 4:32.792 | עולמות |  |
| 4:32.793 | אבי"ע |  |
| 4:32.794 | קדמו |  |
| 4:32.795 | להם | . |
| 4:32.796 | ולרוב |  |
| 4:32.797 | העלמם |  |
| 4:32.798 | לא |  |
| 4:32.799 | שלחו |  |
| 4:32.800 | בהם |  |
| 4:32.801 | יד |  |
| 4:32.802 | להזכירם |  |
| 4:32.803 | בזהר |  |
| 4:32.804 | ובתיקונים |  |
| 4:32.805 | אלא |  |
| 4:32.806 | ברמז |  |
| 4:32.807 | נפלא | . |
| 4:32.808 | ושם |  |
| 4:32.809 | ביאר |  |
| 4:32.810 | עוד |  |
| 4:32.811 | עד |  |
| 4:32.812 | היכן |  |
| 4:32.813 | מגיעים |  |
| 4:32.814 | רגלי |  |
| 4:32.815 | אדם |  |
| 4:32.816 | קדמון |  |
| 4:32.817 | דיושר | . |
| 4:32.818 | ועד |  |
| 4:32.819 | היכן |  |
| 4:32.820 | מגיעים |  |
| 4:32.821 | רגלי |  |
| 4:32.822 | עתיק |  |
| 4:32.823 | יומין | , |
| 4:32.824 | ורגלי |  |
| 4:32.825 | אריך |  |
| 4:32.826 | אנפין | , |
| 4:32.827 | ורגלי |  |
| 4:32.828 | אבא |  |
| 4:32.829 | ואימא | , |
| 4:32.830 | ורגלי |  |
| 4:32.831 | נוקבא |  |
| 4:32.832 | דיליה | . |
| 4:32.833 | ושאבא |  |
| 4:32.834 | ואימא |  |
| 4:32.835 | קצרי |  |
| 4:32.836 | קומה | , |
| 4:32.837 | ואין |  |
| 4:32.838 | שיעור |  |
| 4:32.839 | קומתם |  |
| 4:32.840 | רק |  |
| 4:32.841 | מהגרון |  |
| 4:32.842 | דאריך |  |
| 4:32.843 | אנפין |  |
| 4:32.844 | עד |  |
| 4:32.845 | הטבור |  |
| 4:32.846 | שלאריך |  |
| 4:32.847 | אנפין | , |
| 4:32.848 | וכדברים |  |
| 4:32.849 | האלה |  |
| 4:32.850 | מבואר |  |
| 4:32.851 | ג"כ |  |
| 4:32.852 | בשער |  |
| 4:32.853 | ההקדמות |  |
| 4:32.854 | ובספר |  |
| 4:32.855 | כסא |  |
| 4:32.856 | אליהו |  |
| 4:32.857 | ומקדש |  |
| 4:32.858 | מלך |  |
| 4:32.859 | פרשת |  |
| 4:32.860 | בראשית |  |
| 4:32.861 | עיין |  |
| 4:32.862 | שם |  |
| 4:32.863 | ושם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

