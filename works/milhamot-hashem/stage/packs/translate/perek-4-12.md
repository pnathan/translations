# Translation pack: milhamot-hashem perek-4-12 (Milchamot ha-Shem 4:20–4:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-12`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **הוי"ה / הוי'ה / הויה** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG). אדנות as a name → "Adonai" (note: "lordship").
  עשר הויות → "ten havayot" (note).
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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **מרגל לרגל** → "from festival to festival", not "foot".
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 4:20

Hebrew (Wikisource): מה מאד יחרד לבו ויתר ממקומו בשמעו את מערכת האלהות הרבים, אשר פרו ורבו בישראל מתחלת אלף הששי אצל קצת המחברים, ומדור לדור מזמן הגז"ל הלכה האמונה הזאת הלך וגדול מאד להאמין בעילות רבות זו למעלה מזו, וברצות אחת מהעילות לברוא דבר הוא נמלך ונוטל רשות מן העילה שלמעלה ממנה כמבואר בזהר בהדיא (בראשית כ"ב) שכל אחת מן העלות נוטלת רשות מן העילה שעליה, מלכות מזעיר, וזעיר מאימא, ואימא מאבא, ואבא מאריך, ואריך מעתיק, ועתיק מאדם קדמון. כי הוא הראש לפרצופי האצילות והוא הוא לבדו שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאינו צריך ליטול רשות מאדם קדמאה שלמעלה ממנו. ובכל מעשה בראשית היה המלך הבורא אבא. וקורא לו בזהר מלכה עלאה, ואימא אומנא ובעת בריאת אדם הראשון לא חפץ אבא לבראתו, כיון שהוא עתיד לחטוא, ואימא השיבתו הואיל וחטאו תלוי בי, כדכתיב, ובן כסיל תוגת אמו, לא איכפת לך בזה כמבואר שם ובמקדש מלך.

| id | word | punctuation after |
|---|---|---|
| 4:20.1 | מה |  |
| 4:20.2 | מאד |  |
| 4:20.3 | יחרד |  |
| 4:20.4 | לבו |  |
| 4:20.5 | ויתר |  |
| 4:20.6 | ממקומו |  |
| 4:20.7 | בשמעו |  |
| 4:20.8 | את |  |
| 4:20.9 | מערכת |  |
| 4:20.10 | האלהות |  |
| 4:20.11 | הרבים | , |
| 4:20.12 | אשר |  |
| 4:20.13 | פרו |  |
| 4:20.14 | ורבו |  |
| 4:20.15 | בישראל |  |
| 4:20.16 | מתחלת |  |
| 4:20.17 | אלף |  |
| 4:20.18 | הששי |  |
| 4:20.19 | אצל |  |
| 4:20.20 | קצת |  |
| 4:20.21 | המחברים | , |
| 4:20.22 | ומדור |  |
| 4:20.23 | לדור |  |
| 4:20.24 | מזמן |  |
| 4:20.25 | הגז"ל |  |
| 4:20.26 | הלכה |  |
| 4:20.27 | האמונה |  |
| 4:20.28 | הזאת |  |
| 4:20.29 | הלך |  |
| 4:20.30 | וגדול |  |
| 4:20.31 | מאד |  |
| 4:20.32 | להאמין |  |
| 4:20.33 | בעילות |  |
| 4:20.34 | רבות |  |
| 4:20.35 | זו |  |
| 4:20.36 | למעלה |  |
| 4:20.37 | מזו | , |
| 4:20.38 | וברצות |  |
| 4:20.39 | אחת |  |
| 4:20.40 | מהעילות |  |
| 4:20.41 | לברוא |  |
| 4:20.42 | דבר |  |
| 4:20.43 | הוא |  |
| 4:20.44 | נמלך |  |
| 4:20.45 | ונוטל |  |
| 4:20.46 | רשות |  |
| 4:20.47 | מן |  |
| 4:20.48 | העילה |  |
| 4:20.49 | שלמעלה |  |
| 4:20.50 | ממנה |  |
| 4:20.51 | כמבואר |  |
| 4:20.52 | בזהר |  |
| 4:20.53 | בהדיא |  |
| 4:20.54 | בראשית |  |
| 4:20.55 | כ"ב | ) |
| 4:20.56 | שכל |  |
| 4:20.57 | אחת |  |
| 4:20.58 | מן |  |
| 4:20.59 | העלות |  |
| 4:20.60 | נוטלת |  |
| 4:20.61 | רשות |  |
| 4:20.62 | מן |  |
| 4:20.63 | העילה |  |
| 4:20.64 | שעליה | , |
| 4:20.65 | מלכות |  |
| 4:20.66 | מזעיר | , |
| 4:20.67 | וזעיר |  |
| 4:20.68 | מאימא | , |
| 4:20.69 | ואימא |  |
| 4:20.70 | מאבא | , |
| 4:20.71 | ואבא |  |
| 4:20.72 | מאריך | , |
| 4:20.73 | ואריך |  |
| 4:20.74 | מעתיק | , |
| 4:20.75 | ועתיק |  |
| 4:20.76 | מאדם |  |
| 4:20.77 | קדמון | . |
| 4:20.78 | כי |  |
| 4:20.79 | הוא |  |
| 4:20.80 | הראש |  |
| 4:20.81 | לפרצופי |  |
| 4:20.82 | האצילות |  |
| 4:20.83 | והוא |  |
| 4:20.84 | הוא |  |
| 4:20.85 | לבדו |  |
| 4:20.86 | שאמר |  |
| 4:20.87 | ראו |  |
| 4:20.88 | עתה |  |
| 4:20.89 | כי |  |
| 4:20.90 | אני |  |
| 4:20.91 | אני |  |
| 4:20.92 | הוא |  |
| 4:20.93 | ואין |  |
| 4:20.94 | אלהים |  |
| 4:20.95 | עמדי | , |
| 4:20.96 | שאינו |  |
| 4:20.97 | צריך |  |
| 4:20.98 | ליטול |  |
| 4:20.99 | רשות |  |
| 4:20.100 | מאדם |  |
| 4:20.101 | קדמאה |  |
| 4:20.102 | שלמעלה |  |
| 4:20.103 | ממנו | . |
| 4:20.104 | ובכל |  |
| 4:20.105 | מעשה |  |
| 4:20.106 | בראשית |  |
| 4:20.107 | היה |  |
| 4:20.108 | המלך |  |
| 4:20.109 | הבורא |  |
| 4:20.110 | אבא | . |
| 4:20.111 | וקורא |  |
| 4:20.112 | לו |  |
| 4:20.113 | בזהר |  |
| 4:20.114 | מלכה |  |
| 4:20.115 | עלאה | , |
| 4:20.116 | ואימא |  |
| 4:20.117 | אומנא |  |
| 4:20.118 | ובעת |  |
| 4:20.119 | בריאת |  |
| 4:20.120 | אדם |  |
| 4:20.121 | הראשון |  |
| 4:20.122 | לא |  |
| 4:20.123 | חפץ |  |
| 4:20.124 | אבא |  |
| 4:20.125 | לבראתו | , |
| 4:20.126 | כיון |  |
| 4:20.127 | שהוא |  |
| 4:20.128 | עתיד |  |
| 4:20.129 | לחטוא | , |
| 4:20.130 | ואימא |  |
| 4:20.131 | השיבתו |  |
| 4:20.132 | הואיל |  |
| 4:20.133 | וחטאו |  |
| 4:20.134 | תלוי |  |
| 4:20.135 | בי | , |
| 4:20.136 | כדכתיב | , |
| 4:20.137 | ובן |  |
| 4:20.138 | כסיל |  |
| 4:20.139 | תוגת |  |
| 4:20.140 | אמו | , |
| 4:20.141 | לא |  |
| 4:20.142 | איכפת |  |
| 4:20.143 | לך |  |
| 4:20.144 | בזה |  |
| 4:20.145 | כמבואר |  |
| 4:20.146 | שם |  |
| 4:20.147 | ובמקדש |  |
| 4:20.148 | מלך | . |


## Milchamot ha-Shem 4:21

Hebrew (Wikisource): לא) והדבר הזה קשה עד מאד למה אדם קדמון אין צריך ליטול רשות מאדם קדמאה ולא מא"ס שהוא למעלה ממנו? וכבר הבאתי לעיל לשון ספר הברית משם יושר לבב שאמר ששם הקב"ה המורגל בפינו ראי על זעיר אנפין. וכן שם הוי'ה, וכ"כ מצרף האמונה. רק שלא העתקתי דבריו, כיון שאתה מחלל אותו וחושבו למין וכופר מחמת חסרון בקיאותך. ובודאי שטעית בזה כמו שכתבתי לעיל שכן סוברים כל המקובלים. ובבראשית רבה פרשה י"ב הובא ג"כ בילקוט רמז י"ח וז"ל לכל יש תולדות שמים וארץ יש להן תולדות שנ' אלה תולדות השמים והארץ בהבראם. הרים יש להם תולדות שנ' בטרם הרים יולדו. מטר יש לו תולדות שנ' היש למטר אב. טל יש לו תולדות שנ' או מי הוליד אגלי טל. ר' שמעון בן לקיש אמר אלו מרגליות של טל. תני כל מי שיש לו תולדות מת ובולה ונברא ואני בורא וכל מי שאין לו תולדות אינו לא מת ולא בולה. ובורא ואינו נברא ר' עזריה בשם ר' אמר כלפי מעלה הדבר אמור ע"כ. וזה כלל אמתי וברור. ופי' הר' עץ יוסף מת ובולה הם שני דברים האחד ההפסד והשני הבלייה בעודן קיימין. שכל הנמצאים תמיד חסרים והולכים בטבעם לולי שהקב"ה מחדש בטובו בכל יום תמיד מעשה בראשית וכל מי שאין לו תולדות פי' אע"ג דליכא אלא הקב"ה קאמר כל איידי דתני נושא כל . וכמו שפי' ר' עזריה בשם רבי כלפי מעלה הדבר אמור ע"כ. וממוצא דבר אתה למד שכיון שיש לפרצופים הנזכרים בדברי הזהר והמקובלים תולדות ועיבור ויניקה. שקר הוא ומן הנמנע לקרות לאבא או לזולתו מכל הפרצופים בורא. אי נמי אפשר לפרש מי שיש לו תולדות דהיינו שיש סבה להוייתו שהוא אפשרי המציאות. וכל אפשרי המציאות הוא אפשר בו ההפסד אבל אלהינו ברוך הוא. אין סבה אחרת קודמת לו והוא לבדו סבה א' לכל הברואים. כמו שכ' הרמב"ם ז"ל בהל' יסודי התורה ע"ש ובמפרש, וכן כתבו קדמונינו על מה שכתב רבי' בעיקר הרביעי שנאמין שזה האחד הוא קדמון בהחלט וכל נמצא זולתו בלתי קדמון וכו' כתבו הר' ר' חטר ורבי זכריה הרופא וז"ל הקדמון הוא אשר אין ראשית לו והמחודש הוא שבראו והמציאו זולתו. הרי נתברך שהאלוה הוא הראשון בהחלט ע"ש. והכי מסתברא לפרש דברי המדרש רבה כנראה ממאי דקאמר בהרים ובמטר ובטל. וממילא שקר הוא לקרות שום אחד מן הפרצופים אלוה, או בורא, כיון שכל א' יש סבה אחרת קדמה לו. ובמשנת חסידים מס' ליל פסח פרק ז' משנה כ"ג בפי' הגדת ליל פסח כ' וז"ל ויוציאנו ה' אלהינו שהם אבא ואימא וכו' ואלו לא גאל הקב"ה, שהוא אריך אנפין את אבותינו שהם אבא ואימא ממצרים וכו' ובס' הנדרים כ' בערך אדם כי המקובלים קורין להקב"ה אדם קדמון. ואם כן, כל אחד ואחד מן הפרצופים נקרא ה' צבאות והוי'ה ואדנות והקב"ה ושאר השמות והכנויים המיוחדים לשי"ת. כאשר תחזינה עיניך פה וכ"ש אם תלמוד בספריהם בשום שכל. וגם קורין לאבא ואבא אבותינו שנגאלו ממצרים על ידי אריך אנפין מיד פרעה (הוא הסטרא אחרא הנקרא אל אחר) שנאלם אריך אנפין מידו וזה נמשך לפי שטתם שהם מאמינים במציאות אל אחר היפך דברי רז"ל וגם האמינו שהאלהים האחרים המדומים שלטו בהאלוהים הקדושים שהם אבא ואימא וזעיר ונוקביה, וגאלם אריך אנפין מהם.

| id | word | punctuation after |
|---|---|---|
| 4:21.1 | לא | ) |
| 4:21.2 | והדבר |  |
| 4:21.3 | הזה |  |
| 4:21.4 | קשה |  |
| 4:21.5 | עד |  |
| 4:21.6 | מאד |  |
| 4:21.7 | למה |  |
| 4:21.8 | אדם |  |
| 4:21.9 | קדמון |  |
| 4:21.10 | אין |  |
| 4:21.11 | צריך |  |
| 4:21.12 | ליטול |  |
| 4:21.13 | רשות |  |
| 4:21.14 | מאדם |  |
| 4:21.15 | קדמאה |  |
| 4:21.16 | ולא |  |
| 4:21.17 | מא"ס |  |
| 4:21.18 | שהוא |  |
| 4:21.19 | למעלה |  |
| 4:21.20 | ממנו | ? |
| 4:21.21 | וכבר |  |
| 4:21.22 | הבאתי |  |
| 4:21.23 | לעיל |  |
| 4:21.24 | לשון |  |
| 4:21.25 | ספר |  |
| 4:21.26 | הברית |  |
| 4:21.27 | משם |  |
| 4:21.28 | יושר |  |
| 4:21.29 | לבב |  |
| 4:21.30 | שאמר |  |
| 4:21.31 | ששם |  |
| 4:21.32 | הקב"ה |  |
| 4:21.33 | המורגל |  |
| 4:21.34 | בפינו |  |
| 4:21.35 | ראי |  |
| 4:21.36 | על |  |
| 4:21.37 | זעיר |  |
| 4:21.38 | אנפין | . |
| 4:21.39 | וכן |  |
| 4:21.40 | שם |  |
| 4:21.41 | הוי'ה | , |
| 4:21.42 | וכ"כ |  |
| 4:21.43 | מצרף |  |
| 4:21.44 | האמונה | . |
| 4:21.45 | רק |  |
| 4:21.46 | שלא |  |
| 4:21.47 | העתקתי |  |
| 4:21.48 | דבריו | , |
| 4:21.49 | כיון |  |
| 4:21.50 | שאתה |  |
| 4:21.51 | מחלל |  |
| 4:21.52 | אותו |  |
| 4:21.53 | וחושבו |  |
| 4:21.54 | למין |  |
| 4:21.55 | וכופר |  |
| 4:21.56 | מחמת |  |
| 4:21.57 | חסרון |  |
| 4:21.58 | בקיאותך | . |
| 4:21.59 | ובודאי |  |
| 4:21.60 | שטעית |  |
| 4:21.61 | בזה |  |
| 4:21.62 | כמו |  |
| 4:21.63 | שכתבתי |  |
| 4:21.64 | לעיל |  |
| 4:21.65 | שכן |  |
| 4:21.66 | סוברים |  |
| 4:21.67 | כל |  |
| 4:21.68 | המקובלים | . |
| 4:21.69 | ובבראשית |  |
| 4:21.70 | רבה |  |
| 4:21.71 | פרשה |  |
| 4:21.72 | י"ב |  |
| 4:21.73 | הובא |  |
| 4:21.74 | ג"כ |  |
| 4:21.75 | בילקוט |  |
| 4:21.76 | רמז |  |
| 4:21.77 | י"ח |  |
| 4:21.78 | וז"ל |  |
| 4:21.79 | לכל |  |
| 4:21.80 | יש |  |
| 4:21.81 | תולדות |  |
| 4:21.82 | שמים |  |
| 4:21.83 | וארץ |  |
| 4:21.84 | יש |  |
| 4:21.85 | להן |  |
| 4:21.86 | תולדות |  |
| 4:21.87 | שנ | ' |
| 4:21.88 | אלה |  |
| 4:21.89 | תולדות |  |
| 4:21.90 | השמים |  |
| 4:21.91 | והארץ |  |
| 4:21.92 | בהבראם | . |
| 4:21.93 | הרים |  |
| 4:21.94 | יש |  |
| 4:21.95 | להם |  |
| 4:21.96 | תולדות |  |
| 4:21.97 | שנ | ' |
| 4:21.98 | בטרם |  |
| 4:21.99 | הרים |  |
| 4:21.100 | יולדו | . |
| 4:21.101 | מטר |  |
| 4:21.102 | יש |  |
| 4:21.103 | לו |  |
| 4:21.104 | תולדות |  |
| 4:21.105 | שנ | ' |
| 4:21.106 | היש |  |
| 4:21.107 | למטר |  |
| 4:21.108 | אב | . |
| 4:21.109 | טל |  |
| 4:21.110 | יש |  |
| 4:21.111 | לו |  |
| 4:21.112 | תולדות |  |
| 4:21.113 | שנ | ' |
| 4:21.114 | או |  |
| 4:21.115 | מי |  |
| 4:21.116 | הוליד |  |
| 4:21.117 | אגלי |  |
| 4:21.118 | טל | . |
| 4:21.119 | ר | ' |
| 4:21.120 | שמעון |  |
| 4:21.121 | בן |  |
| 4:21.122 | לקיש |  |
| 4:21.123 | אמר |  |
| 4:21.124 | אלו |  |
| 4:21.125 | מרגליות |  |
| 4:21.126 | של |  |
| 4:21.127 | טל | . |
| 4:21.128 | תני |  |
| 4:21.129 | כל |  |
| 4:21.130 | מי |  |
| 4:21.131 | שיש |  |
| 4:21.132 | לו |  |
| 4:21.133 | תולדות |  |
| 4:21.134 | מת |  |
| 4:21.135 | ובולה |  |
| 4:21.136 | ונברא |  |
| 4:21.137 | ואני |  |
| 4:21.138 | בורא |  |
| 4:21.139 | וכל |  |
| 4:21.140 | מי |  |
| 4:21.141 | שאין |  |
| 4:21.142 | לו |  |
| 4:21.143 | תולדות |  |
| 4:21.144 | אינו |  |
| 4:21.145 | לא |  |
| 4:21.146 | מת |  |
| 4:21.147 | ולא |  |
| 4:21.148 | בולה | . |
| 4:21.149 | ובורא |  |
| 4:21.150 | ואינו |  |
| 4:21.151 | נברא |  |
| 4:21.152 | ר | ' |
| 4:21.153 | עזריה |  |
| 4:21.154 | בשם |  |
| 4:21.155 | ר | ' |
| 4:21.156 | אמר |  |
| 4:21.157 | כלפי |  |
| 4:21.158 | מעלה |  |
| 4:21.159 | הדבר |  |
| 4:21.160 | אמור |  |
| 4:21.161 | ע"כ | . |
| 4:21.162 | וזה |  |
| 4:21.163 | כלל |  |
| 4:21.164 | אמתי |  |
| 4:21.165 | וברור | . |
| 4:21.166 | ופי | ' |
| 4:21.167 | הר | ' |
| 4:21.168 | עץ |  |
| 4:21.169 | יוסף |  |
| 4:21.170 | מת |  |
| 4:21.171 | ובולה |  |
| 4:21.172 | הם |  |
| 4:21.173 | שני |  |
| 4:21.174 | דברים |  |
| 4:21.175 | האחד |  |
| 4:21.176 | ההפסד |  |
| 4:21.177 | והשני |  |
| 4:21.178 | הבלייה |  |
| 4:21.179 | בעודן |  |
| 4:21.180 | קיימין | . |
| 4:21.181 | שכל |  |
| 4:21.182 | הנמצאים |  |
| 4:21.183 | תמיד |  |
| 4:21.184 | חסרים |  |
| 4:21.185 | והולכים |  |
| 4:21.186 | בטבעם |  |
| 4:21.187 | לולי |  |
| 4:21.188 | שהקב"ה |  |
| 4:21.189 | מחדש |  |
| 4:21.190 | בטובו |  |
| 4:21.191 | בכל |  |
| 4:21.192 | יום |  |
| 4:21.193 | תמיד |  |
| 4:21.194 | מעשה |  |
| 4:21.195 | בראשית |  |
| 4:21.196 | וכל |  |
| 4:21.197 | מי |  |
| 4:21.198 | שאין |  |
| 4:21.199 | לו |  |
| 4:21.200 | תולדות |  |
| 4:21.201 | פי | ' |
| 4:21.202 | אע"ג |  |
| 4:21.203 | דליכא |  |
| 4:21.204 | אלא |  |
| 4:21.205 | הקב"ה |  |
| 4:21.206 | קאמר |  |
| 4:21.207 | כל |  |
| 4:21.208 | איידי |  |
| 4:21.209 | דתני |  |
| 4:21.210 | נושא |  |
| 4:21.211 | כל | . |
| 4:21.212 | וכמו |  |
| 4:21.213 | שפי | ' |
| 4:21.214 | ר | ' |
| 4:21.215 | עזריה |  |
| 4:21.216 | בשם |  |
| 4:21.217 | רבי |  |
| 4:21.218 | כלפי |  |
| 4:21.219 | מעלה |  |
| 4:21.220 | הדבר |  |
| 4:21.221 | אמור |  |
| 4:21.222 | ע"כ | . |
| 4:21.223 | וממוצא |  |
| 4:21.224 | דבר |  |
| 4:21.225 | אתה |  |
| 4:21.226 | למד |  |
| 4:21.227 | שכיון |  |
| 4:21.228 | שיש |  |
| 4:21.229 | לפרצופים |  |
| 4:21.230 | הנזכרים |  |
| 4:21.231 | בדברי |  |
| 4:21.232 | הזהר |  |
| 4:21.233 | והמקובלים |  |
| 4:21.234 | תולדות |  |
| 4:21.235 | ועיבור |  |
| 4:21.236 | ויניקה | . |
| 4:21.237 | שקר |  |
| 4:21.238 | הוא |  |
| 4:21.239 | ומן |  |
| 4:21.240 | הנמנע |  |
| 4:21.241 | לקרות |  |
| 4:21.242 | לאבא |  |
| 4:21.243 | או |  |
| 4:21.244 | לזולתו |  |
| 4:21.245 | מכל |  |
| 4:21.246 | הפרצופים |  |
| 4:21.247 | בורא | . |
| 4:21.248 | אי |  |
| 4:21.249 | נמי |  |
| 4:21.250 | אפשר |  |
| 4:21.251 | לפרש |  |
| 4:21.252 | מי |  |
| 4:21.253 | שיש |  |
| 4:21.254 | לו |  |
| 4:21.255 | תולדות |  |
| 4:21.256 | דהיינו |  |
| 4:21.257 | שיש |  |
| 4:21.258 | סבה |  |
| 4:21.259 | להוייתו |  |
| 4:21.260 | שהוא |  |
| 4:21.261 | אפשרי |  |
| 4:21.262 | המציאות | . |
| 4:21.263 | וכל |  |
| 4:21.264 | אפשרי |  |
| 4:21.265 | המציאות |  |
| 4:21.266 | הוא |  |
| 4:21.267 | אפשר |  |
| 4:21.268 | בו |  |
| 4:21.269 | ההפסד |  |
| 4:21.270 | אבל |  |
| 4:21.271 | אלהינו |  |
| 4:21.272 | ברוך |  |
| 4:21.273 | הוא | . |
| 4:21.274 | אין |  |
| 4:21.275 | סבה |  |
| 4:21.276 | אחרת |  |
| 4:21.277 | קודמת |  |
| 4:21.278 | לו |  |
| 4:21.279 | והוא |  |
| 4:21.280 | לבדו |  |
| 4:21.281 | סבה |  |
| 4:21.282 | א | ' |
| 4:21.283 | לכל |  |
| 4:21.284 | הברואים | . |
| 4:21.285 | כמו |  |
| 4:21.286 | שכ | ' |
| 4:21.287 | הרמב"ם |  |
| 4:21.288 | ז"ל |  |
| 4:21.289 | בהל | ' |
| 4:21.290 | יסודי |  |
| 4:21.291 | התורה |  |
| 4:21.292 | ע"ש |  |
| 4:21.293 | ובמפרש | , |
| 4:21.294 | וכן |  |
| 4:21.295 | כתבו |  |
| 4:21.296 | קדמונינו |  |
| 4:21.297 | על |  |
| 4:21.298 | מה |  |
| 4:21.299 | שכתב |  |
| 4:21.300 | רבי | ' |
| 4:21.301 | בעיקר |  |
| 4:21.302 | הרביעי |  |
| 4:21.303 | שנאמין |  |
| 4:21.304 | שזה |  |
| 4:21.305 | האחד |  |
| 4:21.306 | הוא |  |
| 4:21.307 | קדמון |  |
| 4:21.308 | בהחלט |  |
| 4:21.309 | וכל |  |
| 4:21.310 | נמצא |  |
| 4:21.311 | זולתו |  |
| 4:21.312 | בלתי |  |
| 4:21.313 | קדמון |  |
| 4:21.314 | וכו | ' |
| 4:21.315 | כתבו |  |
| 4:21.316 | הר | ' |
| 4:21.317 | ר | ' |
| 4:21.318 | חטר |  |
| 4:21.319 | ורבי |  |
| 4:21.320 | זכריה |  |
| 4:21.321 | הרופא |  |
| 4:21.322 | וז"ל |  |
| 4:21.323 | הקדמון |  |
| 4:21.324 | הוא |  |
| 4:21.325 | אשר |  |
| 4:21.326 | אין |  |
| 4:21.327 | ראשית |  |
| 4:21.328 | לו |  |
| 4:21.329 | והמחודש |  |
| 4:21.330 | הוא |  |
| 4:21.331 | שבראו |  |
| 4:21.332 | והמציאו |  |
| 4:21.333 | זולתו | . |
| 4:21.334 | הרי |  |
| 4:21.335 | נתברך |  |
| 4:21.336 | שהאלוה |  |
| 4:21.337 | הוא |  |
| 4:21.338 | הראשון |  |
| 4:21.339 | בהחלט |  |
| 4:21.340 | ע"ש | . |
| 4:21.341 | והכי |  |
| 4:21.342 | מסתברא |  |
| 4:21.343 | לפרש |  |
| 4:21.344 | דברי |  |
| 4:21.345 | המדרש |  |
| 4:21.346 | רבה |  |
| 4:21.347 | כנראה |  |
| 4:21.348 | ממאי |  |
| 4:21.349 | דקאמר |  |
| 4:21.350 | בהרים |  |
| 4:21.351 | ובמטר |  |
| 4:21.352 | ובטל | . |
| 4:21.353 | וממילא |  |
| 4:21.354 | שקר |  |
| 4:21.355 | הוא |  |
| 4:21.356 | לקרות |  |
| 4:21.357 | שום |  |
| 4:21.358 | אחד |  |
| 4:21.359 | מן |  |
| 4:21.360 | הפרצופים |  |
| 4:21.361 | אלוה | , |
| 4:21.362 | או |  |
| 4:21.363 | בורא | , |
| 4:21.364 | כיון |  |
| 4:21.365 | שכל |  |
| 4:21.366 | א | ' |
| 4:21.367 | יש |  |
| 4:21.368 | סבה |  |
| 4:21.369 | אחרת |  |
| 4:21.370 | קדמה |  |
| 4:21.371 | לו | . |
| 4:21.372 | ובמשנת |  |
| 4:21.373 | חסידים |  |
| 4:21.374 | מס | ' |
| 4:21.375 | ליל |  |
| 4:21.376 | פסח |  |
| 4:21.377 | פרק |  |
| 4:21.378 | ז | ' |
| 4:21.379 | משנה |  |
| 4:21.380 | כ"ג |  |
| 4:21.381 | בפי | ' |
| 4:21.382 | הגדת |  |
| 4:21.383 | ליל |  |
| 4:21.384 | פסח |  |
| 4:21.385 | כ | ' |
| 4:21.386 | וז"ל |  |
| 4:21.387 | ויוציאנו |  |
| 4:21.388 | ה | ' |
| 4:21.389 | אלהינו |  |
| 4:21.390 | שהם |  |
| 4:21.391 | אבא |  |
| 4:21.392 | ואימא |  |
| 4:21.393 | וכו | ' |
| 4:21.394 | ואלו |  |
| 4:21.395 | לא |  |
| 4:21.396 | גאל |  |
| 4:21.397 | הקב"ה | , |
| 4:21.398 | שהוא |  |
| 4:21.399 | אריך |  |
| 4:21.400 | אנפין |  |
| 4:21.401 | את |  |
| 4:21.402 | אבותינו |  |
| 4:21.403 | שהם |  |
| 4:21.404 | אבא |  |
| 4:21.405 | ואימא |  |
| 4:21.406 | ממצרים |  |
| 4:21.407 | וכו | ' |
| 4:21.408 | ובס | ' |
| 4:21.409 | הנדרים |  |
| 4:21.410 | כ | ' |
| 4:21.411 | בערך |  |
| 4:21.412 | אדם |  |
| 4:21.413 | כי |  |
| 4:21.414 | המקובלים |  |
| 4:21.415 | קורין |  |
| 4:21.416 | להקב"ה |  |
| 4:21.417 | אדם |  |
| 4:21.418 | קדמון | . |
| 4:21.419 | ואם |  |
| 4:21.420 | כן | , |
| 4:21.421 | כל |  |
| 4:21.422 | אחד |  |
| 4:21.423 | ואחד |  |
| 4:21.424 | מן |  |
| 4:21.425 | הפרצופים |  |
| 4:21.426 | נקרא |  |
| 4:21.427 | ה | ' |
| 4:21.428 | צבאות |  |
| 4:21.429 | והוי'ה |  |
| 4:21.430 | ואדנות |  |
| 4:21.431 | והקב"ה |  |
| 4:21.432 | ושאר |  |
| 4:21.433 | השמות |  |
| 4:21.434 | והכנויים |  |
| 4:21.435 | המיוחדים |  |
| 4:21.436 | לשי"ת | . |
| 4:21.437 | כאשר |  |
| 4:21.438 | תחזינה |  |
| 4:21.439 | עיניך |  |
| 4:21.440 | פה |  |
| 4:21.441 | וכ"ש |  |
| 4:21.442 | אם |  |
| 4:21.443 | תלמוד |  |
| 4:21.444 | בספריהם |  |
| 4:21.445 | בשום |  |
| 4:21.446 | שכל | . |
| 4:21.447 | וגם |  |
| 4:21.448 | קורין |  |
| 4:21.449 | לאבא |  |
| 4:21.450 | ואבא |  |
| 4:21.451 | אבותינו |  |
| 4:21.452 | שנגאלו |  |
| 4:21.453 | ממצרים |  |
| 4:21.454 | על |  |
| 4:21.455 | ידי |  |
| 4:21.456 | אריך |  |
| 4:21.457 | אנפין |  |
| 4:21.458 | מיד |  |
| 4:21.459 | פרעה |  |
| 4:21.460 | הוא |  |
| 4:21.461 | הסטרא |  |
| 4:21.462 | אחרא |  |
| 4:21.463 | הנקרא |  |
| 4:21.464 | אל |  |
| 4:21.465 | אחר | ) |
| 4:21.466 | שנאלם |  |
| 4:21.467 | אריך |  |
| 4:21.468 | אנפין |  |
| 4:21.469 | מידו |  |
| 4:21.470 | וזה |  |
| 4:21.471 | נמשך |  |
| 4:21.472 | לפי |  |
| 4:21.473 | שטתם |  |
| 4:21.474 | שהם |  |
| 4:21.475 | מאמינים |  |
| 4:21.476 | במציאות |  |
| 4:21.477 | אל |  |
| 4:21.478 | אחר |  |
| 4:21.479 | היפך |  |
| 4:21.480 | דברי |  |
| 4:21.481 | רז"ל |  |
| 4:21.482 | וגם |  |
| 4:21.483 | האמינו |  |
| 4:21.484 | שהאלהים |  |
| 4:21.485 | האחרים |  |
| 4:21.486 | המדומים |  |
| 4:21.487 | שלטו |  |
| 4:21.488 | בהאלוהים |  |
| 4:21.489 | הקדושים |  |
| 4:21.490 | שהם |  |
| 4:21.491 | אבא |  |
| 4:21.492 | ואימא |  |
| 4:21.493 | וזעיר |  |
| 4:21.494 | ונוקביה | , |
| 4:21.495 | וגאלם |  |
| 4:21.496 | אריך |  |
| 4:21.497 | אנפין |  |
| 4:21.498 | מהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

