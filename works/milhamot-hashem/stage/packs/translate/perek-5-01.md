# Translation pack: milhamot-hashem perek-5-01 (Milchamot ha-Shem 5:1–5:1)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-01`.

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
- **שיתוף / שתוף / לשתף / משותף**: associating another with God → fixed,
  "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever associates
  the name of Heaven with something else"). The Maimonidean term for a word said of God and creatures →
  "equivocation / equivocal" ("בשתוף השם בלבד" → "by equivocation of the term
  only"), never "association", Guide I in the note; "shared" only for a thing
  held in common. צירוף: "conjunction" of God; "permutation" in letter
  mysticism.
- **עבודה זרה / ע"ז / עכו"ם / עבודה אחרת**: עבודה זרה → "idolatry" (note once:
  "strange worship"); the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of
  idols". עכו"ם → "idolaters" (note: "worshippers of stars and
  constellations"), never "gentiles" or "Christians", even of a Christian king
  or the Trinity. עבודה אחרת → "other worship" (note first in a unit). אליל →
  "idol"; האלילה → "the she-idol"; אל זר → "a strange god"; אל אחר → "another
  god"; אלהים אחרים → "other gods".
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **המתפלספים** → "would-be philosophers".
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


## Milchamot ha-Shem 5:1

Hebrew (Wikisource): מ) והמקובלים החדשים שציירו את האין סוף המקיף אחר הצמצום ככדור עגול וחלול באמצעו ובתוכו עיגולים לאלפים ורבבות הנמשכים מעצמותו יתברך, גם פרצופים דיושר אין קץ למספרם שנמשכו ונשתלשלו מעצמותו דרך הקו המתפשט ממנו מאדם קדמאה שהוא למעלה מעולם האצילות ועד סוף פרצופי כל האבי"ע, עליהם אמר מורנו ורבנו הרמב"ם ז"ל בפרק נ' מראשון: "קרוב אתה בפיהם ורחוק מכליותיהם" (ע"פ ירמיהו י"ב ב'). כי בפיהם אומרים שהוא אחד ובכל לבבם ובכל נפשם ומחשבותם - שהוא נחלק לכמה פרצופים ומדרגות זה למעלה מזה, וצריכים הם ליטול רשות זה מזה כמו שנתבאר לעיל. וכמעט שאמונה זו גרועה מאמונת העכו"ם וכמו שכתב הריב"ש בשם אחד מן המתפלספים. כי הנוצרים אומרים שהם שלושה והשלושה אחד, והמקובלים אומרים שהם חמישה בעולם האצילות והחמישה נחלקים לי"ב, והכל אחד. וזוכר אני בספר בית יהודה כתב שהדברים הנזכרים בריב"ש בשם אחד מהמתפלספים מילתא דנפשיה היא, רק שתלה אותם באחד מן המתפלספים פן יפגעו בו הרודפים אחר הכבוד ויזיקוהו, לכן תלה הדבר באחרים. ולכך סיים שאינו תוקע עצמו ללמוד הקבלה החדשה, כי די לו בקבלה הישנה שהיא המשנה והתלמוד. והמקובלים החדשים המשילו מציאותו יתברך ואמתתו לגוף האדם שהוא מורכב מרמ"ח אברים ושס"ה גידים והוא נקרא בשם אחד ראובן, או לבית המחובר מאבנים וסיד ועצים וטיט ויש בו כמה חדרים, והוא הנקרא בשם אחד בית. ולא זו בלבד אלא גם בסיבה הראשונה הנקרא אצלם אין סוף עשו בו ריבוי כמ"ש מהרח"ו בספר ארבע מאות שקל כסף דף ע"ב וז"ל: "ראוי שתדע כי הלא כל המציאות כולו האין סוף אשר אנו קורין אותו עתיקא דכל עתיקין האמור כדי שיתגלה לנו קצת הארה ממנו, והוא גנוז וטמיר בכתר כנודע דאין סוף גנוז בכתר. וראוי שתדע כי יש עדיין אין סוף אשר הוא לעילא לעילא, כי אין שום מחשבה תפיס ביה והוא עמוק מי ימצאנו, ולכן תמצא בזהר כמה מיני אין סוף בכמה דוכתי אשר כמעט יסתכן האדם בנפשו, יען מי שאין לו הקדמות אינו יכול לשוט בימא רבה. ולכן על זה האין סוף למטה למדו בזהר אדרת נשא. ואז בהיותו שם בסוד תלת רישין אז הוא נקרא עתיקא דכל עתיקין, וז"ס מה שאמרו האי עתיקא קדישא אשתכח בתלת רישין. נמצא כי עיקר עתיקא קדישא הוא האין סוף, אמנם ע"ש התלבשותו בתלת רישין נקראים התלת רישין עתיקא קדישא", עכ"ל. מתבאר בהדיא מדבריו כי הא"ס הראשון שהוא עמוק עמוק ומקיף את כל העולמות אין לכוון תפילתו אליו כלל, כי הוא עמוק עמוק ורחוק מאד ולית מחשבה תפיסא ביה כלל וכלל. וכל תפילותנו אל הזעיר אנפין להמשיך השפע מאבא ואמא ומאין סוף הגנוז בכתר בלבד יתגלה לנו הארתו. אבל אין סוף הקדום הנעלה לית מחשבה תפיסה ביה כלל. ומי שלא ידע זה יבוא לכוון תפילותיו אליו, ויסתכן בנפשו שלא תקובל תפילתו. ואדרבה ענוש ייענש כמ"ש לעיל בשם ספר כסא אליהו. ודברים אלה היפך דבריך שכתבת בשם לחם שלמה. ואין ספר לחם שלמה מצוי אצלנו לעיין בו ואפשר שאתה מטעה אותנו וכותב לנו בשמו היפך כוונתו כמו שעשית בכוונת כסא אליהו ואינו כן. כי המוסכם על בעלי הקבלה החדשה שהאין סוף שהוא הסיבה הראשונה אין שום תפילה וברכה שייכא ביה לדעתם כלל כי הוא עמוק ורחוק מאד. ורק כל תפילותנו להמשיך מאין סוף הגנוז בכתר בלבד כדי שיתגלה לנו הארתו. ולמה זה לא נכוון בתפילותנו אל האין סוף הקדום המקיף את כל המציאות? ולמה יקצוף עלינו בעת אשר נקרא אליו? וגם למה לא הודיע כל זאת ע"י נאמן ביתו משה איש האלהים, לא בתושב"כ ולא בתושבע"פ שהיא המשנה והתלמוד המקובל מפה אל פה? וכן כל הנביאים שעמדו להם לישראל מימות משה ועד האחרון שבהם - מלאכי - המזהיר אותנו (מלאכי ג' כ"ב): "זכרו תורת משה עבדי אשר צויתי אותו בחורב על כל ישראל חוקים ומשפטים" - ולא צירופים ותיקון עולמות. ולמה לא הודיעונו כ"ז? וילכו כל הדורות שעברו שולל (לפי דבריהם) להתפלל ח"ו אל אל לא יושיע? היה להם להודיע לישראל כי לא אל אל עליון נורא יתפללו כי אינו מושג. ורק אל קצר האפים (הזעיר אנפין) שכבר נתעבה והושג בחוש יתפללו. והן עתה נכפלה קושיתנו וחזקה. כי הרבה מיני אין סוף הם והרבה פרצופים יש שכ"א עושה פעולה, וזה הוא אמונת הרבה רשויות בשמים, וכבא הבאתי לעיל דברי כסא אליהו להודיע לך טעותך שחשבת לישב דבריו. וכל דבריך הבל וריק ישאם רוח. וברור הוא שדעתו שצריך שיתפלל לזעיר אנפין בשיתוף שאר הפרצופים שלמעלה ממנו.

| id | word | punctuation after |
|---|---|---|
| 5:1.1 | מ | ) |
| 5:1.2 | והמקובלים |  |
| 5:1.3 | החדשים |  |
| 5:1.4 | שציירו |  |
| 5:1.5 | את |  |
| 5:1.6 | האין |  |
| 5:1.7 | סוף |  |
| 5:1.8 | המקיף |  |
| 5:1.9 | אחר |  |
| 5:1.10 | הצמצום |  |
| 5:1.11 | ככדור |  |
| 5:1.12 | עגול |  |
| 5:1.13 | וחלול |  |
| 5:1.14 | באמצעו |  |
| 5:1.15 | ובתוכו |  |
| 5:1.16 | עיגולים |  |
| 5:1.17 | לאלפים |  |
| 5:1.18 | ורבבות |  |
| 5:1.19 | הנמשכים |  |
| 5:1.20 | מעצמותו |  |
| 5:1.21 | יתברך | , |
| 5:1.22 | גם |  |
| 5:1.23 | פרצופים |  |
| 5:1.24 | דיושר |  |
| 5:1.25 | אין |  |
| 5:1.26 | קץ |  |
| 5:1.27 | למספרם |  |
| 5:1.28 | שנמשכו |  |
| 5:1.29 | ונשתלשלו |  |
| 5:1.30 | מעצמותו |  |
| 5:1.31 | דרך |  |
| 5:1.32 | הקו |  |
| 5:1.33 | המתפשט |  |
| 5:1.34 | ממנו |  |
| 5:1.35 | מאדם |  |
| 5:1.36 | קדמאה |  |
| 5:1.37 | שהוא |  |
| 5:1.38 | למעלה |  |
| 5:1.39 | מעולם |  |
| 5:1.40 | האצילות |  |
| 5:1.41 | ועד |  |
| 5:1.42 | סוף |  |
| 5:1.43 | פרצופי |  |
| 5:1.44 | כל |  |
| 5:1.45 | האבי"ע | , |
| 5:1.46 | עליהם |  |
| 5:1.47 | אמר |  |
| 5:1.48 | מורנו |  |
| 5:1.49 | ורבנו |  |
| 5:1.50 | הרמב"ם |  |
| 5:1.51 | ז"ל |  |
| 5:1.52 | בפרק |  |
| 5:1.53 | נ | ' |
| 5:1.54 | מראשון | : |
| 5:1.55 | קרוב |  |
| 5:1.56 | אתה |  |
| 5:1.57 | בפיהם |  |
| 5:1.58 | ורחוק |  |
| 5:1.59 | מכליותיהם | " |
| 5:1.60 | ע"פ |  |
| 5:1.61 | ירמיהו |  |
| 5:1.62 | י"ב |  |
| 5:1.63 | ב | '). |
| 5:1.64 | כי |  |
| 5:1.65 | בפיהם |  |
| 5:1.66 | אומרים |  |
| 5:1.67 | שהוא |  |
| 5:1.68 | אחד |  |
| 5:1.69 | ובכל |  |
| 5:1.70 | לבבם |  |
| 5:1.71 | ובכל |  |
| 5:1.72 | נפשם |  |
| 5:1.73 | ומחשבותם | - |
| 5:1.74 | שהוא |  |
| 5:1.75 | נחלק |  |
| 5:1.76 | לכמה |  |
| 5:1.77 | פרצופים |  |
| 5:1.78 | ומדרגות |  |
| 5:1.79 | זה |  |
| 5:1.80 | למעלה |  |
| 5:1.81 | מזה | , |
| 5:1.82 | וצריכים |  |
| 5:1.83 | הם |  |
| 5:1.84 | ליטול |  |
| 5:1.85 | רשות |  |
| 5:1.86 | זה |  |
| 5:1.87 | מזה |  |
| 5:1.88 | כמו |  |
| 5:1.89 | שנתבאר |  |
| 5:1.90 | לעיל | . |
| 5:1.91 | וכמעט |  |
| 5:1.92 | שאמונה |  |
| 5:1.93 | זו |  |
| 5:1.94 | גרועה |  |
| 5:1.95 | מאמונת |  |
| 5:1.96 | העכו"ם |  |
| 5:1.97 | וכמו |  |
| 5:1.98 | שכתב |  |
| 5:1.99 | הריב"ש |  |
| 5:1.100 | בשם |  |
| 5:1.101 | אחד |  |
| 5:1.102 | מן |  |
| 5:1.103 | המתפלספים | . |
| 5:1.104 | כי |  |
| 5:1.105 | הנוצרים |  |
| 5:1.106 | אומרים |  |
| 5:1.107 | שהם |  |
| 5:1.108 | שלושה |  |
| 5:1.109 | והשלושה |  |
| 5:1.110 | אחד | , |
| 5:1.111 | והמקובלים |  |
| 5:1.112 | אומרים |  |
| 5:1.113 | שהם |  |
| 5:1.114 | חמישה |  |
| 5:1.115 | בעולם |  |
| 5:1.116 | האצילות |  |
| 5:1.117 | והחמישה |  |
| 5:1.118 | נחלקים |  |
| 5:1.119 | לי"ב | , |
| 5:1.120 | והכל |  |
| 5:1.121 | אחד | . |
| 5:1.122 | וזוכר |  |
| 5:1.123 | אני |  |
| 5:1.124 | בספר |  |
| 5:1.125 | בית |  |
| 5:1.126 | יהודה |  |
| 5:1.127 | כתב |  |
| 5:1.128 | שהדברים |  |
| 5:1.129 | הנזכרים |  |
| 5:1.130 | בריב"ש |  |
| 5:1.131 | בשם |  |
| 5:1.132 | אחד |  |
| 5:1.133 | מהמתפלספים |  |
| 5:1.134 | מילתא |  |
| 5:1.135 | דנפשיה |  |
| 5:1.136 | היא | , |
| 5:1.137 | רק |  |
| 5:1.138 | שתלה |  |
| 5:1.139 | אותם |  |
| 5:1.140 | באחד |  |
| 5:1.141 | מן |  |
| 5:1.142 | המתפלספים |  |
| 5:1.143 | פן |  |
| 5:1.144 | יפגעו |  |
| 5:1.145 | בו |  |
| 5:1.146 | הרודפים |  |
| 5:1.147 | אחר |  |
| 5:1.148 | הכבוד |  |
| 5:1.149 | ויזיקוהו | , |
| 5:1.150 | לכן |  |
| 5:1.151 | תלה |  |
| 5:1.152 | הדבר |  |
| 5:1.153 | באחרים | . |
| 5:1.154 | ולכך |  |
| 5:1.155 | סיים |  |
| 5:1.156 | שאינו |  |
| 5:1.157 | תוקע |  |
| 5:1.158 | עצמו |  |
| 5:1.159 | ללמוד |  |
| 5:1.160 | הקבלה |  |
| 5:1.161 | החדשה | , |
| 5:1.162 | כי |  |
| 5:1.163 | די |  |
| 5:1.164 | לו |  |
| 5:1.165 | בקבלה |  |
| 5:1.166 | הישנה |  |
| 5:1.167 | שהיא |  |
| 5:1.168 | המשנה |  |
| 5:1.169 | והתלמוד | . |
| 5:1.170 | והמקובלים |  |
| 5:1.171 | החדשים |  |
| 5:1.172 | המשילו |  |
| 5:1.173 | מציאותו |  |
| 5:1.174 | יתברך |  |
| 5:1.175 | ואמתתו |  |
| 5:1.176 | לגוף |  |
| 5:1.177 | האדם |  |
| 5:1.178 | שהוא |  |
| 5:1.179 | מורכב |  |
| 5:1.180 | מרמ"ח |  |
| 5:1.181 | אברים |  |
| 5:1.182 | ושס"ה |  |
| 5:1.183 | גידים |  |
| 5:1.184 | והוא |  |
| 5:1.185 | נקרא |  |
| 5:1.186 | בשם |  |
| 5:1.187 | אחד |  |
| 5:1.188 | ראובן | , |
| 5:1.189 | או |  |
| 5:1.190 | לבית |  |
| 5:1.191 | המחובר |  |
| 5:1.192 | מאבנים |  |
| 5:1.193 | וסיד |  |
| 5:1.194 | ועצים |  |
| 5:1.195 | וטיט |  |
| 5:1.196 | ויש |  |
| 5:1.197 | בו |  |
| 5:1.198 | כמה |  |
| 5:1.199 | חדרים | , |
| 5:1.200 | והוא |  |
| 5:1.201 | הנקרא |  |
| 5:1.202 | בשם |  |
| 5:1.203 | אחד |  |
| 5:1.204 | בית | . |
| 5:1.205 | ולא |  |
| 5:1.206 | זו |  |
| 5:1.207 | בלבד |  |
| 5:1.208 | אלא |  |
| 5:1.209 | גם |  |
| 5:1.210 | בסיבה |  |
| 5:1.211 | הראשונה |  |
| 5:1.212 | הנקרא |  |
| 5:1.213 | אצלם |  |
| 5:1.214 | אין |  |
| 5:1.215 | סוף |  |
| 5:1.216 | עשו |  |
| 5:1.217 | בו |  |
| 5:1.218 | ריבוי |  |
| 5:1.219 | כמ"ש |  |
| 5:1.220 | מהרח"ו |  |
| 5:1.221 | בספר |  |
| 5:1.222 | ארבע |  |
| 5:1.223 | מאות |  |
| 5:1.224 | שקל |  |
| 5:1.225 | כסף |  |
| 5:1.226 | דף |  |
| 5:1.227 | ע"ב |  |
| 5:1.228 | וז"ל | : |
| 5:1.229 | ראוי |  |
| 5:1.230 | שתדע |  |
| 5:1.231 | כי |  |
| 5:1.232 | הלא |  |
| 5:1.233 | כל |  |
| 5:1.234 | המציאות |  |
| 5:1.235 | כולו |  |
| 5:1.236 | האין |  |
| 5:1.237 | סוף |  |
| 5:1.238 | אשר |  |
| 5:1.239 | אנו |  |
| 5:1.240 | קורין |  |
| 5:1.241 | אותו |  |
| 5:1.242 | עתיקא |  |
| 5:1.243 | דכל |  |
| 5:1.244 | עתיקין |  |
| 5:1.245 | האמור |  |
| 5:1.246 | כדי |  |
| 5:1.247 | שיתגלה |  |
| 5:1.248 | לנו |  |
| 5:1.249 | קצת |  |
| 5:1.250 | הארה |  |
| 5:1.251 | ממנו | , |
| 5:1.252 | והוא |  |
| 5:1.253 | גנוז |  |
| 5:1.254 | וטמיר |  |
| 5:1.255 | בכתר |  |
| 5:1.256 | כנודע |  |
| 5:1.257 | דאין |  |
| 5:1.258 | סוף |  |
| 5:1.259 | גנוז |  |
| 5:1.260 | בכתר | . |
| 5:1.261 | וראוי |  |
| 5:1.262 | שתדע |  |
| 5:1.263 | כי |  |
| 5:1.264 | יש |  |
| 5:1.265 | עדיין |  |
| 5:1.266 | אין |  |
| 5:1.267 | סוף |  |
| 5:1.268 | אשר |  |
| 5:1.269 | הוא |  |
| 5:1.270 | לעילא |  |
| 5:1.271 | לעילא | , |
| 5:1.272 | כי |  |
| 5:1.273 | אין |  |
| 5:1.274 | שום |  |
| 5:1.275 | מחשבה |  |
| 5:1.276 | תפיס |  |
| 5:1.277 | ביה |  |
| 5:1.278 | והוא |  |
| 5:1.279 | עמוק |  |
| 5:1.280 | מי |  |
| 5:1.281 | ימצאנו | , |
| 5:1.282 | ולכן |  |
| 5:1.283 | תמצא |  |
| 5:1.284 | בזהר |  |
| 5:1.285 | כמה |  |
| 5:1.286 | מיני |  |
| 5:1.287 | אין |  |
| 5:1.288 | סוף |  |
| 5:1.289 | בכמה |  |
| 5:1.290 | דוכתי |  |
| 5:1.291 | אשר |  |
| 5:1.292 | כמעט |  |
| 5:1.293 | יסתכן |  |
| 5:1.294 | האדם |  |
| 5:1.295 | בנפשו | , |
| 5:1.296 | יען |  |
| 5:1.297 | מי |  |
| 5:1.298 | שאין |  |
| 5:1.299 | לו |  |
| 5:1.300 | הקדמות |  |
| 5:1.301 | אינו |  |
| 5:1.302 | יכול |  |
| 5:1.303 | לשוט |  |
| 5:1.304 | בימא |  |
| 5:1.305 | רבה | . |
| 5:1.306 | ולכן |  |
| 5:1.307 | על |  |
| 5:1.308 | זה |  |
| 5:1.309 | האין |  |
| 5:1.310 | סוף |  |
| 5:1.311 | למטה |  |
| 5:1.312 | למדו |  |
| 5:1.313 | בזהר |  |
| 5:1.314 | אדרת |  |
| 5:1.315 | נשא | . |
| 5:1.316 | ואז |  |
| 5:1.317 | בהיותו |  |
| 5:1.318 | שם |  |
| 5:1.319 | בסוד |  |
| 5:1.320 | תלת |  |
| 5:1.321 | רישין |  |
| 5:1.322 | אז |  |
| 5:1.323 | הוא |  |
| 5:1.324 | נקרא |  |
| 5:1.325 | עתיקא |  |
| 5:1.326 | דכל |  |
| 5:1.327 | עתיקין | , |
| 5:1.328 | וז"ס |  |
| 5:1.329 | מה |  |
| 5:1.330 | שאמרו |  |
| 5:1.331 | האי |  |
| 5:1.332 | עתיקא |  |
| 5:1.333 | קדישא |  |
| 5:1.334 | אשתכח |  |
| 5:1.335 | בתלת |  |
| 5:1.336 | רישין | . |
| 5:1.337 | נמצא |  |
| 5:1.338 | כי |  |
| 5:1.339 | עיקר |  |
| 5:1.340 | עתיקא |  |
| 5:1.341 | קדישא |  |
| 5:1.342 | הוא |  |
| 5:1.343 | האין |  |
| 5:1.344 | סוף | , |
| 5:1.345 | אמנם |  |
| 5:1.346 | ע"ש |  |
| 5:1.347 | התלבשותו |  |
| 5:1.348 | בתלת |  |
| 5:1.349 | רישין |  |
| 5:1.350 | נקראים |  |
| 5:1.351 | התלת |  |
| 5:1.352 | רישין |  |
| 5:1.353 | עתיקא |  |
| 5:1.354 | קדישא | ", |
| 5:1.355 | עכ"ל | . |
| 5:1.356 | מתבאר |  |
| 5:1.357 | בהדיא |  |
| 5:1.358 | מדבריו |  |
| 5:1.359 | כי |  |
| 5:1.360 | הא"ס |  |
| 5:1.361 | הראשון |  |
| 5:1.362 | שהוא |  |
| 5:1.363 | עמוק |  |
| 5:1.364 | עמוק |  |
| 5:1.365 | ומקיף |  |
| 5:1.366 | את |  |
| 5:1.367 | כל |  |
| 5:1.368 | העולמות |  |
| 5:1.369 | אין |  |
| 5:1.370 | לכוון |  |
| 5:1.371 | תפילתו |  |
| 5:1.372 | אליו |  |
| 5:1.373 | כלל | , |
| 5:1.374 | כי |  |
| 5:1.375 | הוא |  |
| 5:1.376 | עמוק |  |
| 5:1.377 | עמוק |  |
| 5:1.378 | ורחוק |  |
| 5:1.379 | מאד |  |
| 5:1.380 | ולית |  |
| 5:1.381 | מחשבה |  |
| 5:1.382 | תפיסא |  |
| 5:1.383 | ביה |  |
| 5:1.384 | כלל |  |
| 5:1.385 | וכלל | . |
| 5:1.386 | וכל |  |
| 5:1.387 | תפילותנו |  |
| 5:1.388 | אל |  |
| 5:1.389 | הזעיר |  |
| 5:1.390 | אנפין |  |
| 5:1.391 | להמשיך |  |
| 5:1.392 | השפע |  |
| 5:1.393 | מאבא |  |
| 5:1.394 | ואמא |  |
| 5:1.395 | ומאין |  |
| 5:1.396 | סוף |  |
| 5:1.397 | הגנוז |  |
| 5:1.398 | בכתר |  |
| 5:1.399 | בלבד |  |
| 5:1.400 | יתגלה |  |
| 5:1.401 | לנו |  |
| 5:1.402 | הארתו | . |
| 5:1.403 | אבל |  |
| 5:1.404 | אין |  |
| 5:1.405 | סוף |  |
| 5:1.406 | הקדום |  |
| 5:1.407 | הנעלה |  |
| 5:1.408 | לית |  |
| 5:1.409 | מחשבה |  |
| 5:1.410 | תפיסה |  |
| 5:1.411 | ביה |  |
| 5:1.412 | כלל | . |
| 5:1.413 | ומי |  |
| 5:1.414 | שלא |  |
| 5:1.415 | ידע |  |
| 5:1.416 | זה |  |
| 5:1.417 | יבוא |  |
| 5:1.418 | לכוון |  |
| 5:1.419 | תפילותיו |  |
| 5:1.420 | אליו | , |
| 5:1.421 | ויסתכן |  |
| 5:1.422 | בנפשו |  |
| 5:1.423 | שלא |  |
| 5:1.424 | תקובל |  |
| 5:1.425 | תפילתו | . |
| 5:1.426 | ואדרבה |  |
| 5:1.427 | ענוש |  |
| 5:1.428 | ייענש |  |
| 5:1.429 | כמ"ש |  |
| 5:1.430 | לעיל |  |
| 5:1.431 | בשם |  |
| 5:1.432 | ספר |  |
| 5:1.433 | כסא |  |
| 5:1.434 | אליהו | . |
| 5:1.435 | ודברים |  |
| 5:1.436 | אלה |  |
| 5:1.437 | היפך |  |
| 5:1.438 | דבריך |  |
| 5:1.439 | שכתבת |  |
| 5:1.440 | בשם |  |
| 5:1.441 | לחם |  |
| 5:1.442 | שלמה | . |
| 5:1.443 | ואין |  |
| 5:1.444 | ספר |  |
| 5:1.445 | לחם |  |
| 5:1.446 | שלמה |  |
| 5:1.447 | מצוי |  |
| 5:1.448 | אצלנו |  |
| 5:1.449 | לעיין |  |
| 5:1.450 | בו |  |
| 5:1.451 | ואפשר |  |
| 5:1.452 | שאתה |  |
| 5:1.453 | מטעה |  |
| 5:1.454 | אותנו |  |
| 5:1.455 | וכותב |  |
| 5:1.456 | לנו |  |
| 5:1.457 | בשמו |  |
| 5:1.458 | היפך |  |
| 5:1.459 | כוונתו |  |
| 5:1.460 | כמו |  |
| 5:1.461 | שעשית |  |
| 5:1.462 | בכוונת |  |
| 5:1.463 | כסא |  |
| 5:1.464 | אליהו |  |
| 5:1.465 | ואינו |  |
| 5:1.466 | כן | . |
| 5:1.467 | כי |  |
| 5:1.468 | המוסכם |  |
| 5:1.469 | על |  |
| 5:1.470 | בעלי |  |
| 5:1.471 | הקבלה |  |
| 5:1.472 | החדשה |  |
| 5:1.473 | שהאין |  |
| 5:1.474 | סוף |  |
| 5:1.475 | שהוא |  |
| 5:1.476 | הסיבה |  |
| 5:1.477 | הראשונה |  |
| 5:1.478 | אין |  |
| 5:1.479 | שום |  |
| 5:1.480 | תפילה |  |
| 5:1.481 | וברכה |  |
| 5:1.482 | שייכא |  |
| 5:1.483 | ביה |  |
| 5:1.484 | לדעתם |  |
| 5:1.485 | כלל |  |
| 5:1.486 | כי |  |
| 5:1.487 | הוא |  |
| 5:1.488 | עמוק |  |
| 5:1.489 | ורחוק |  |
| 5:1.490 | מאד | . |
| 5:1.491 | ורק |  |
| 5:1.492 | כל |  |
| 5:1.493 | תפילותנו |  |
| 5:1.494 | להמשיך |  |
| 5:1.495 | מאין |  |
| 5:1.496 | סוף |  |
| 5:1.497 | הגנוז |  |
| 5:1.498 | בכתר |  |
| 5:1.499 | בלבד |  |
| 5:1.500 | כדי |  |
| 5:1.501 | שיתגלה |  |
| 5:1.502 | לנו |  |
| 5:1.503 | הארתו | . |
| 5:1.504 | ולמה |  |
| 5:1.505 | זה |  |
| 5:1.506 | לא |  |
| 5:1.507 | נכוון |  |
| 5:1.508 | בתפילותנו |  |
| 5:1.509 | אל |  |
| 5:1.510 | האין |  |
| 5:1.511 | סוף |  |
| 5:1.512 | הקדום |  |
| 5:1.513 | המקיף |  |
| 5:1.514 | את |  |
| 5:1.515 | כל |  |
| 5:1.516 | המציאות | ? |
| 5:1.517 | ולמה |  |
| 5:1.518 | יקצוף |  |
| 5:1.519 | עלינו |  |
| 5:1.520 | בעת |  |
| 5:1.521 | אשר |  |
| 5:1.522 | נקרא |  |
| 5:1.523 | אליו | ? |
| 5:1.524 | וגם |  |
| 5:1.525 | למה |  |
| 5:1.526 | לא |  |
| 5:1.527 | הודיע |  |
| 5:1.528 | כל |  |
| 5:1.529 | זאת |  |
| 5:1.530 | ע"י |  |
| 5:1.531 | נאמן |  |
| 5:1.532 | ביתו |  |
| 5:1.533 | משה |  |
| 5:1.534 | איש |  |
| 5:1.535 | האלהים | , |
| 5:1.536 | לא |  |
| 5:1.537 | בתושב"כ |  |
| 5:1.538 | ולא |  |
| 5:1.539 | בתושבע"פ |  |
| 5:1.540 | שהיא |  |
| 5:1.541 | המשנה |  |
| 5:1.542 | והתלמוד |  |
| 5:1.543 | המקובל |  |
| 5:1.544 | מפה |  |
| 5:1.545 | אל |  |
| 5:1.546 | פה | ? |
| 5:1.547 | וכן |  |
| 5:1.548 | כל |  |
| 5:1.549 | הנביאים |  |
| 5:1.550 | שעמדו |  |
| 5:1.551 | להם |  |
| 5:1.552 | לישראל |  |
| 5:1.553 | מימות |  |
| 5:1.554 | משה |  |
| 5:1.555 | ועד |  |
| 5:1.556 | האחרון |  |
| 5:1.557 | שבהם | - |
| 5:1.558 | מלאכי | - |
| 5:1.559 | המזהיר |  |
| 5:1.560 | אותנו |  |
| 5:1.561 | מלאכי |  |
| 5:1.562 | ג | ' |
| 5:1.563 | כ"ב | ): |
| 5:1.564 | זכרו |  |
| 5:1.565 | תורת |  |
| 5:1.566 | משה |  |
| 5:1.567 | עבדי |  |
| 5:1.568 | אשר |  |
| 5:1.569 | צויתי |  |
| 5:1.570 | אותו |  |
| 5:1.571 | בחורב |  |
| 5:1.572 | על |  |
| 5:1.573 | כל |  |
| 5:1.574 | ישראל |  |
| 5:1.575 | חוקים |  |
| 5:1.576 | ומשפטים | "- |
| 5:1.577 | ולא |  |
| 5:1.578 | צירופים |  |
| 5:1.579 | ותיקון |  |
| 5:1.580 | עולמות | . |
| 5:1.581 | ולמה |  |
| 5:1.582 | לא |  |
| 5:1.583 | הודיעונו |  |
| 5:1.584 | כ"ז | ? |
| 5:1.585 | וילכו |  |
| 5:1.586 | כל |  |
| 5:1.587 | הדורות |  |
| 5:1.588 | שעברו |  |
| 5:1.589 | שולל |  |
| 5:1.590 | לפי |  |
| 5:1.591 | דבריהם | ) |
| 5:1.592 | להתפלל |  |
| 5:1.593 | ח"ו |  |
| 5:1.594 | אל |  |
| 5:1.595 | אל |  |
| 5:1.596 | לא |  |
| 5:1.597 | יושיע | ? |
| 5:1.598 | היה |  |
| 5:1.599 | להם |  |
| 5:1.600 | להודיע |  |
| 5:1.601 | לישראל |  |
| 5:1.602 | כי |  |
| 5:1.603 | לא |  |
| 5:1.604 | אל |  |
| 5:1.605 | אל |  |
| 5:1.606 | עליון |  |
| 5:1.607 | נורא |  |
| 5:1.608 | יתפללו |  |
| 5:1.609 | כי |  |
| 5:1.610 | אינו |  |
| 5:1.611 | מושג | . |
| 5:1.612 | ורק |  |
| 5:1.613 | אל |  |
| 5:1.614 | קצר |  |
| 5:1.615 | האפים |  |
| 5:1.616 | הזעיר |  |
| 5:1.617 | אנפין | ) |
| 5:1.618 | שכבר |  |
| 5:1.619 | נתעבה |  |
| 5:1.620 | והושג |  |
| 5:1.621 | בחוש |  |
| 5:1.622 | יתפללו | . |
| 5:1.623 | והן |  |
| 5:1.624 | עתה |  |
| 5:1.625 | נכפלה |  |
| 5:1.626 | קושיתנו |  |
| 5:1.627 | וחזקה | . |
| 5:1.628 | כי |  |
| 5:1.629 | הרבה |  |
| 5:1.630 | מיני |  |
| 5:1.631 | אין |  |
| 5:1.632 | סוף |  |
| 5:1.633 | הם |  |
| 5:1.634 | והרבה |  |
| 5:1.635 | פרצופים |  |
| 5:1.636 | יש |  |
| 5:1.637 | שכ"א |  |
| 5:1.638 | עושה |  |
| 5:1.639 | פעולה | , |
| 5:1.640 | וזה |  |
| 5:1.641 | הוא |  |
| 5:1.642 | אמונת |  |
| 5:1.643 | הרבה |  |
| 5:1.644 | רשויות |  |
| 5:1.645 | בשמים | , |
| 5:1.646 | וכבא |  |
| 5:1.647 | הבאתי |  |
| 5:1.648 | לעיל |  |
| 5:1.649 | דברי |  |
| 5:1.650 | כסא |  |
| 5:1.651 | אליהו |  |
| 5:1.652 | להודיע |  |
| 5:1.653 | לך |  |
| 5:1.654 | טעותך |  |
| 5:1.655 | שחשבת |  |
| 5:1.656 | לישב |  |
| 5:1.657 | דבריו | . |
| 5:1.658 | וכל |  |
| 5:1.659 | דבריך |  |
| 5:1.660 | הבל |  |
| 5:1.661 | וריק |  |
| 5:1.662 | ישאם |  |
| 5:1.663 | רוח | . |
| 5:1.664 | וברור |  |
| 5:1.665 | הוא |  |
| 5:1.666 | שדעתו |  |
| 5:1.667 | שצריך |  |
| 5:1.668 | שיתפלל |  |
| 5:1.669 | לזעיר |  |
| 5:1.670 | אנפין |  |
| 5:1.671 | בשיתוף |  |
| 5:1.672 | שאר |  |
| 5:1.673 | הפרצופים |  |
| 5:1.674 | שלמעלה |  |
| 5:1.675 | ממנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

