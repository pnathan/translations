# Translation pack: milhamot-hashem perek-6-20 (Milchamot ha-Shem 6:30–6:31)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-20`.

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
- **הוי"ה / הוי'ה / הויה** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG). אדנות as a name → "Adonai" (note: "lordship").
  עשר הויות → "ten havayot" (note).
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
- **אצילות / באצילות / בריאה**: fixed, translated, the verbs linked:
  "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; נאצל → "emanated"; המאציל → "the Emanator". The four worlds
  (אבי"ע): "Emanation, Creation, Formation, Making" (verbs: "emanated,
  created, formed, made"). נאצל and נברא stay distinct ("נברא ולא בורא" →
  "created, not creator").
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 6:30

Hebrew (Wikisource): בהא סלקינן ובהא נחתינן. כיון שהמסית מחבר הזהר בא להסתינו ולהדיחנו מעל ה' אלוהינו. ופגע בכבוד אבינו מלכינו וביסודי תורתינו הקדושה ושרשיה הטהורים ביחוד השי"ת שהוא הסבה הראשונה. ובא ליסד לנו דת ותורה חדשה להאמין בסבות רבות שהם משתלשלות זו מזו. וכל אחת נקראת בשם הוי'ה ואדנות ואלהים והקב"ה ויחס התפלה והעבודה והקרבנות שמקריבין ישראל בזמן שבית המקדש קיים לנברא קצר אפים (זעיר אנפין) והוא הסבה האחרונה שבפרצופי האצילות. אע"פ שאמר שנגלה לו עתיקא קדישא או אליהו ור"מ וכיוצא לא נאמין לו, כי שקר בפיו. שכן קבלנו מרז"ל שאין נביא רשאי לחדש דבר מעתה. ואף אם בת קול תכריז מן השמים לאמר שמעו אליו, לא נשמע לו לסור מן הדרך אשר צוה ה' אלהינו לזלזל בשום עיקר מעיקרי תורתינו. כל שכן זה שהוא עוקר ארבעה עיקרים מיסודי תוה"ק שמנאם הרמב"ם ז"ל בפירוש המשנה, כמו שנתבאר לעיל.

| id | word | punctuation after |
|---|---|---|
| 6:30.1 | בהא |  |
| 6:30.2 | סלקינן |  |
| 6:30.3 | ובהא |  |
| 6:30.4 | נחתינן | . |
| 6:30.5 | כיון |  |
| 6:30.6 | שהמסית |  |
| 6:30.7 | מחבר |  |
| 6:30.8 | הזהר |  |
| 6:30.9 | בא |  |
| 6:30.10 | להסתינו |  |
| 6:30.11 | ולהדיחנו |  |
| 6:30.12 | מעל |  |
| 6:30.13 | ה | ' |
| 6:30.14 | אלוהינו | . |
| 6:30.15 | ופגע |  |
| 6:30.16 | בכבוד |  |
| 6:30.17 | אבינו |  |
| 6:30.18 | מלכינו |  |
| 6:30.19 | וביסודי |  |
| 6:30.20 | תורתינו |  |
| 6:30.21 | הקדושה |  |
| 6:30.22 | ושרשיה |  |
| 6:30.23 | הטהורים |  |
| 6:30.24 | ביחוד |  |
| 6:30.25 | השי"ת |  |
| 6:30.26 | שהוא |  |
| 6:30.27 | הסבה |  |
| 6:30.28 | הראשונה | . |
| 6:30.29 | ובא |  |
| 6:30.30 | ליסד |  |
| 6:30.31 | לנו |  |
| 6:30.32 | דת |  |
| 6:30.33 | ותורה |  |
| 6:30.34 | חדשה |  |
| 6:30.35 | להאמין |  |
| 6:30.36 | בסבות |  |
| 6:30.37 | רבות |  |
| 6:30.38 | שהם |  |
| 6:30.39 | משתלשלות |  |
| 6:30.40 | זו |  |
| 6:30.41 | מזו | . |
| 6:30.42 | וכל |  |
| 6:30.43 | אחת |  |
| 6:30.44 | נקראת |  |
| 6:30.45 | בשם |  |
| 6:30.46 | הוי'ה |  |
| 6:30.47 | ואדנות |  |
| 6:30.48 | ואלהים |  |
| 6:30.49 | והקב"ה |  |
| 6:30.50 | ויחס |  |
| 6:30.51 | התפלה |  |
| 6:30.52 | והעבודה |  |
| 6:30.53 | והקרבנות |  |
| 6:30.54 | שמקריבין |  |
| 6:30.55 | ישראל |  |
| 6:30.56 | בזמן |  |
| 6:30.57 | שבית |  |
| 6:30.58 | המקדש |  |
| 6:30.59 | קיים |  |
| 6:30.60 | לנברא |  |
| 6:30.61 | קצר |  |
| 6:30.62 | אפים |  |
| 6:30.63 | זעיר |  |
| 6:30.64 | אנפין | ) |
| 6:30.65 | והוא |  |
| 6:30.66 | הסבה |  |
| 6:30.67 | האחרונה |  |
| 6:30.68 | שבפרצופי |  |
| 6:30.69 | האצילות | . |
| 6:30.70 | אע"פ |  |
| 6:30.71 | שאמר |  |
| 6:30.72 | שנגלה |  |
| 6:30.73 | לו |  |
| 6:30.74 | עתיקא |  |
| 6:30.75 | קדישא |  |
| 6:30.76 | או |  |
| 6:30.77 | אליהו |  |
| 6:30.78 | ור"מ |  |
| 6:30.79 | וכיוצא |  |
| 6:30.80 | לא |  |
| 6:30.81 | נאמין |  |
| 6:30.82 | לו | , |
| 6:30.83 | כי |  |
| 6:30.84 | שקר |  |
| 6:30.85 | בפיו | . |
| 6:30.86 | שכן |  |
| 6:30.87 | קבלנו |  |
| 6:30.88 | מרז"ל |  |
| 6:30.89 | שאין |  |
| 6:30.90 | נביא |  |
| 6:30.91 | רשאי |  |
| 6:30.92 | לחדש |  |
| 6:30.93 | דבר |  |
| 6:30.94 | מעתה | . |
| 6:30.95 | ואף |  |
| 6:30.96 | אם |  |
| 6:30.97 | בת |  |
| 6:30.98 | קול |  |
| 6:30.99 | תכריז |  |
| 6:30.100 | מן |  |
| 6:30.101 | השמים |  |
| 6:30.102 | לאמר |  |
| 6:30.103 | שמעו |  |
| 6:30.104 | אליו | , |
| 6:30.105 | לא |  |
| 6:30.106 | נשמע |  |
| 6:30.107 | לו |  |
| 6:30.108 | לסור |  |
| 6:30.109 | מן |  |
| 6:30.110 | הדרך |  |
| 6:30.111 | אשר |  |
| 6:30.112 | צוה |  |
| 6:30.113 | ה | ' |
| 6:30.114 | אלהינו |  |
| 6:30.115 | לזלזל |  |
| 6:30.116 | בשום |  |
| 6:30.117 | עיקר |  |
| 6:30.118 | מעיקרי |  |
| 6:30.119 | תורתינו | . |
| 6:30.120 | כל |  |
| 6:30.121 | שכן |  |
| 6:30.122 | זה |  |
| 6:30.123 | שהוא |  |
| 6:30.124 | עוקר |  |
| 6:30.125 | ארבעה |  |
| 6:30.126 | עיקרים |  |
| 6:30.127 | מיסודי |  |
| 6:30.128 | תוה"ק |  |
| 6:30.129 | שמנאם |  |
| 6:30.130 | הרמב"ם |  |
| 6:30.131 | ז"ל |  |
| 6:30.132 | בפירוש |  |
| 6:30.133 | המשנה | , |
| 6:30.134 | כמו |  |
| 6:30.135 | שנתבאר |  |
| 6:30.136 | לעיל | . |


## Milchamot ha-Shem 6:31

Hebrew (Wikisource): וכתב בספר שבילי אמונה תחלת הנתיב הראשון וז"ל דרך אמונה בחרתי משפטיך שויתי וכו' וכתבו חכמי המחקר, לא יוכל לעבוד עלת כל העילות אלא נביא הדור בטבעו או החכם המובהק במה שקנה מן החכמה. וכאשר ישיג האדם לזה, אזי יגיע לקבל שכר הטוב ויתקיים בו כוונת הבריאה כאמרו כל הנקרא בשמי ולכבודי בראתיו. ולפי עיקר זה, נאמר בראשית עשרת הדברות אנכי ה' אלהיך כלומר האמן בלבך אמתת מציאותי. ואח"כ תוכל לקבל עול מלכותי. לפי שאין קבלת עול מלכות ראויה. רק אחר אמונת מציאותו. ואמר ה' אלהיך כלומר הוה קדמון. ומאתו היה הכל בחפץ ויכולת. והוא אלהים להם, שחייבים לעבוד אותו. ואמר אשר הוצאתיך מארץ מצרים, כי הוצאתו אותנו משם. מורה על המציאות והחפץ. כי בידיעה ובהשגחה ממנו ית' יצאו. וגם תורה על החדוש. כי אם העולם קדמון לא ישתנה דבר מטבעו. ותורה על היכולת, ותורה על היחוד. כמו שאמר בעבור תדע כי אין כמוני בכל הארץ ומצוה זו לא נאמרה בדברות על דרך הצווי, לפי שאין ראוי לצוות אלא אחר שידע השמוע מי הוא המצווה ויאמין בלבו שהוא נמצא על ענין כך. ולפיכך אמר בדבור זה אנכי ה' אלהיך, ולא אמר על צד הצווי אלא על צד ההאמנה. ואחר שהודיענו אמתת מציאותו ודרכי הנהגתו לעולם, אזי התחיל לצוות אותנו ולהודיע לנו שאין דבר זולתו ראוי לעבודה. לפי שכל הנמצאים זולתו נמצאו במאמרו, ואין ראוי לכלול עמו ובקבלת עול מלכותו אחד מכל העבדים. שלא להשוות כבוד הבורא יתברך לכבוד הנבראים. לפיכך סמך מצוה זו לאמתת מציאותו. לפי שאמונת עבודת זולתו מכחשת מציאותו ועל דבר זה התחיל לצוות ולהודיע שאין ראוי לעבוד לזולתו. ולכן הזכיר ואמר לא יהיה לך אלהים אחרים על פני. ועל יעלה בדעתך ומחשבתך ממה שאמר לא יהיה לך אלהים אחרים שיש מציאות אלוה אחר זולתו. כי מה שאמר אלהים הוא לפי מחשבת עובדיהם. כמו והאנשים רדפו אחריהם. ואמר אחרים, כלומר אלהים נמשכים אחר אחרים, שאין להם כח לעצמן, שקיומן תלוי בדבר אחר זולתן. ומי שאין לו כח קיים מעצמו אלא מזולתו אין ראוי להיעבד. ולא להשבע בשמו. אבל העובדה הראויה היא ליחיד אשר הוא נושא ומקיים כל הבריות הנמשכים זה אחר זה, ועדיו ית' יגיע תכליתן והוא אין לו נושא ואין מקיים ית' שמו ויתעלה זכרו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 6:31.1 | וכתב |  |
| 6:31.2 | בספר |  |
| 6:31.3 | שבילי |  |
| 6:31.4 | אמונה |  |
| 6:31.5 | תחלת |  |
| 6:31.6 | הנתיב |  |
| 6:31.7 | הראשון |  |
| 6:31.8 | וז"ל |  |
| 6:31.9 | דרך |  |
| 6:31.10 | אמונה |  |
| 6:31.11 | בחרתי |  |
| 6:31.12 | משפטיך |  |
| 6:31.13 | שויתי |  |
| 6:31.14 | וכו | ' |
| 6:31.15 | וכתבו |  |
| 6:31.16 | חכמי |  |
| 6:31.17 | המחקר | , |
| 6:31.18 | לא |  |
| 6:31.19 | יוכל |  |
| 6:31.20 | לעבוד |  |
| 6:31.21 | עלת |  |
| 6:31.22 | כל |  |
| 6:31.23 | העילות |  |
| 6:31.24 | אלא |  |
| 6:31.25 | נביא |  |
| 6:31.26 | הדור |  |
| 6:31.27 | בטבעו |  |
| 6:31.28 | או |  |
| 6:31.29 | החכם |  |
| 6:31.30 | המובהק |  |
| 6:31.31 | במה |  |
| 6:31.32 | שקנה |  |
| 6:31.33 | מן |  |
| 6:31.34 | החכמה | . |
| 6:31.35 | וכאשר |  |
| 6:31.36 | ישיג |  |
| 6:31.37 | האדם |  |
| 6:31.38 | לזה | , |
| 6:31.39 | אזי |  |
| 6:31.40 | יגיע |  |
| 6:31.41 | לקבל |  |
| 6:31.42 | שכר |  |
| 6:31.43 | הטוב |  |
| 6:31.44 | ויתקיים |  |
| 6:31.45 | בו |  |
| 6:31.46 | כוונת |  |
| 6:31.47 | הבריאה |  |
| 6:31.48 | כאמרו |  |
| 6:31.49 | כל |  |
| 6:31.50 | הנקרא |  |
| 6:31.51 | בשמי |  |
| 6:31.52 | ולכבודי |  |
| 6:31.53 | בראתיו | . |
| 6:31.54 | ולפי |  |
| 6:31.55 | עיקר |  |
| 6:31.56 | זה | , |
| 6:31.57 | נאמר |  |
| 6:31.58 | בראשית |  |
| 6:31.59 | עשרת |  |
| 6:31.60 | הדברות |  |
| 6:31.61 | אנכי |  |
| 6:31.62 | ה | ' |
| 6:31.63 | אלהיך |  |
| 6:31.64 | כלומר |  |
| 6:31.65 | האמן |  |
| 6:31.66 | בלבך |  |
| 6:31.67 | אמתת |  |
| 6:31.68 | מציאותי | . |
| 6:31.69 | ואח"כ |  |
| 6:31.70 | תוכל |  |
| 6:31.71 | לקבל |  |
| 6:31.72 | עול |  |
| 6:31.73 | מלכותי | . |
| 6:31.74 | לפי |  |
| 6:31.75 | שאין |  |
| 6:31.76 | קבלת |  |
| 6:31.77 | עול |  |
| 6:31.78 | מלכות |  |
| 6:31.79 | ראויה | . |
| 6:31.80 | רק |  |
| 6:31.81 | אחר |  |
| 6:31.82 | אמונת |  |
| 6:31.83 | מציאותו | . |
| 6:31.84 | ואמר |  |
| 6:31.85 | ה | ' |
| 6:31.86 | אלהיך |  |
| 6:31.87 | כלומר |  |
| 6:31.88 | הוה |  |
| 6:31.89 | קדמון | . |
| 6:31.90 | ומאתו |  |
| 6:31.91 | היה |  |
| 6:31.92 | הכל |  |
| 6:31.93 | בחפץ |  |
| 6:31.94 | ויכולת | . |
| 6:31.95 | והוא |  |
| 6:31.96 | אלהים |  |
| 6:31.97 | להם | , |
| 6:31.98 | שחייבים |  |
| 6:31.99 | לעבוד |  |
| 6:31.100 | אותו | . |
| 6:31.101 | ואמר |  |
| 6:31.102 | אשר |  |
| 6:31.103 | הוצאתיך |  |
| 6:31.104 | מארץ |  |
| 6:31.105 | מצרים | , |
| 6:31.106 | כי |  |
| 6:31.107 | הוצאתו |  |
| 6:31.108 | אותנו |  |
| 6:31.109 | משם | . |
| 6:31.110 | מורה |  |
| 6:31.111 | על |  |
| 6:31.112 | המציאות |  |
| 6:31.113 | והחפץ | . |
| 6:31.114 | כי |  |
| 6:31.115 | בידיעה |  |
| 6:31.116 | ובהשגחה |  |
| 6:31.117 | ממנו |  |
| 6:31.118 | ית | ' |
| 6:31.119 | יצאו | . |
| 6:31.120 | וגם |  |
| 6:31.121 | תורה |  |
| 6:31.122 | על |  |
| 6:31.123 | החדוש | . |
| 6:31.124 | כי |  |
| 6:31.125 | אם |  |
| 6:31.126 | העולם |  |
| 6:31.127 | קדמון |  |
| 6:31.128 | לא |  |
| 6:31.129 | ישתנה |  |
| 6:31.130 | דבר |  |
| 6:31.131 | מטבעו | . |
| 6:31.132 | ותורה |  |
| 6:31.133 | על |  |
| 6:31.134 | היכולת | , |
| 6:31.135 | ותורה |  |
| 6:31.136 | על |  |
| 6:31.137 | היחוד | . |
| 6:31.138 | כמו |  |
| 6:31.139 | שאמר |  |
| 6:31.140 | בעבור |  |
| 6:31.141 | תדע |  |
| 6:31.142 | כי |  |
| 6:31.143 | אין |  |
| 6:31.144 | כמוני |  |
| 6:31.145 | בכל |  |
| 6:31.146 | הארץ |  |
| 6:31.147 | ומצוה |  |
| 6:31.148 | זו |  |
| 6:31.149 | לא |  |
| 6:31.150 | נאמרה |  |
| 6:31.151 | בדברות |  |
| 6:31.152 | על |  |
| 6:31.153 | דרך |  |
| 6:31.154 | הצווי | , |
| 6:31.155 | לפי |  |
| 6:31.156 | שאין |  |
| 6:31.157 | ראוי |  |
| 6:31.158 | לצוות |  |
| 6:31.159 | אלא |  |
| 6:31.160 | אחר |  |
| 6:31.161 | שידע |  |
| 6:31.162 | השמוע |  |
| 6:31.163 | מי |  |
| 6:31.164 | הוא |  |
| 6:31.165 | המצווה |  |
| 6:31.166 | ויאמין |  |
| 6:31.167 | בלבו |  |
| 6:31.168 | שהוא |  |
| 6:31.169 | נמצא |  |
| 6:31.170 | על |  |
| 6:31.171 | ענין |  |
| 6:31.172 | כך | . |
| 6:31.173 | ולפיכך |  |
| 6:31.174 | אמר |  |
| 6:31.175 | בדבור |  |
| 6:31.176 | זה |  |
| 6:31.177 | אנכי |  |
| 6:31.178 | ה | ' |
| 6:31.179 | אלהיך | , |
| 6:31.180 | ולא |  |
| 6:31.181 | אמר |  |
| 6:31.182 | על |  |
| 6:31.183 | צד |  |
| 6:31.184 | הצווי |  |
| 6:31.185 | אלא |  |
| 6:31.186 | על |  |
| 6:31.187 | צד |  |
| 6:31.188 | ההאמנה | . |
| 6:31.189 | ואחר |  |
| 6:31.190 | שהודיענו |  |
| 6:31.191 | אמתת |  |
| 6:31.192 | מציאותו |  |
| 6:31.193 | ודרכי |  |
| 6:31.194 | הנהגתו |  |
| 6:31.195 | לעולם | , |
| 6:31.196 | אזי |  |
| 6:31.197 | התחיל |  |
| 6:31.198 | לצוות |  |
| 6:31.199 | אותנו |  |
| 6:31.200 | ולהודיע |  |
| 6:31.201 | לנו |  |
| 6:31.202 | שאין |  |
| 6:31.203 | דבר |  |
| 6:31.204 | זולתו |  |
| 6:31.205 | ראוי |  |
| 6:31.206 | לעבודה | . |
| 6:31.207 | לפי |  |
| 6:31.208 | שכל |  |
| 6:31.209 | הנמצאים |  |
| 6:31.210 | זולתו |  |
| 6:31.211 | נמצאו |  |
| 6:31.212 | במאמרו | , |
| 6:31.213 | ואין |  |
| 6:31.214 | ראוי |  |
| 6:31.215 | לכלול |  |
| 6:31.216 | עמו |  |
| 6:31.217 | ובקבלת |  |
| 6:31.218 | עול |  |
| 6:31.219 | מלכותו |  |
| 6:31.220 | אחד |  |
| 6:31.221 | מכל |  |
| 6:31.222 | העבדים | . |
| 6:31.223 | שלא |  |
| 6:31.224 | להשוות |  |
| 6:31.225 | כבוד |  |
| 6:31.226 | הבורא |  |
| 6:31.227 | יתברך |  |
| 6:31.228 | לכבוד |  |
| 6:31.229 | הנבראים | . |
| 6:31.230 | לפיכך |  |
| 6:31.231 | סמך |  |
| 6:31.232 | מצוה |  |
| 6:31.233 | זו |  |
| 6:31.234 | לאמתת |  |
| 6:31.235 | מציאותו | . |
| 6:31.236 | לפי |  |
| 6:31.237 | שאמונת |  |
| 6:31.238 | עבודת |  |
| 6:31.239 | זולתו |  |
| 6:31.240 | מכחשת |  |
| 6:31.241 | מציאותו |  |
| 6:31.242 | ועל |  |
| 6:31.243 | דבר |  |
| 6:31.244 | זה |  |
| 6:31.245 | התחיל |  |
| 6:31.246 | לצוות |  |
| 6:31.247 | ולהודיע |  |
| 6:31.248 | שאין |  |
| 6:31.249 | ראוי |  |
| 6:31.250 | לעבוד |  |
| 6:31.251 | לזולתו | . |
| 6:31.252 | ולכן |  |
| 6:31.253 | הזכיר |  |
| 6:31.254 | ואמר |  |
| 6:31.255 | לא |  |
| 6:31.256 | יהיה |  |
| 6:31.257 | לך |  |
| 6:31.258 | אלהים |  |
| 6:31.259 | אחרים |  |
| 6:31.260 | על |  |
| 6:31.261 | פני | . |
| 6:31.262 | ועל |  |
| 6:31.263 | יעלה |  |
| 6:31.264 | בדעתך |  |
| 6:31.265 | ומחשבתך |  |
| 6:31.266 | ממה |  |
| 6:31.267 | שאמר |  |
| 6:31.268 | לא |  |
| 6:31.269 | יהיה |  |
| 6:31.270 | לך |  |
| 6:31.271 | אלהים |  |
| 6:31.272 | אחרים |  |
| 6:31.273 | שיש |  |
| 6:31.274 | מציאות |  |
| 6:31.275 | אלוה |  |
| 6:31.276 | אחר |  |
| 6:31.277 | זולתו | . |
| 6:31.278 | כי |  |
| 6:31.279 | מה |  |
| 6:31.280 | שאמר |  |
| 6:31.281 | אלהים |  |
| 6:31.282 | הוא |  |
| 6:31.283 | לפי |  |
| 6:31.284 | מחשבת |  |
| 6:31.285 | עובדיהם | . |
| 6:31.286 | כמו |  |
| 6:31.287 | והאנשים |  |
| 6:31.288 | רדפו |  |
| 6:31.289 | אחריהם | . |
| 6:31.290 | ואמר |  |
| 6:31.291 | אחרים | , |
| 6:31.292 | כלומר |  |
| 6:31.293 | אלהים |  |
| 6:31.294 | נמשכים |  |
| 6:31.295 | אחר |  |
| 6:31.296 | אחרים | , |
| 6:31.297 | שאין |  |
| 6:31.298 | להם |  |
| 6:31.299 | כח |  |
| 6:31.300 | לעצמן | , |
| 6:31.301 | שקיומן |  |
| 6:31.302 | תלוי |  |
| 6:31.303 | בדבר |  |
| 6:31.304 | אחר |  |
| 6:31.305 | זולתן | . |
| 6:31.306 | ומי |  |
| 6:31.307 | שאין |  |
| 6:31.308 | לו |  |
| 6:31.309 | כח |  |
| 6:31.310 | קיים |  |
| 6:31.311 | מעצמו |  |
| 6:31.312 | אלא |  |
| 6:31.313 | מזולתו |  |
| 6:31.314 | אין |  |
| 6:31.315 | ראוי |  |
| 6:31.316 | להיעבד | . |
| 6:31.317 | ולא |  |
| 6:31.318 | להשבע |  |
| 6:31.319 | בשמו | . |
| 6:31.320 | אבל |  |
| 6:31.321 | העובדה |  |
| 6:31.322 | הראויה |  |
| 6:31.323 | היא |  |
| 6:31.324 | ליחיד |  |
| 6:31.325 | אשר |  |
| 6:31.326 | הוא |  |
| 6:31.327 | נושא |  |
| 6:31.328 | ומקיים |  |
| 6:31.329 | כל |  |
| 6:31.330 | הבריות |  |
| 6:31.331 | הנמשכים |  |
| 6:31.332 | זה |  |
| 6:31.333 | אחר |  |
| 6:31.334 | זה | , |
| 6:31.335 | ועדיו |  |
| 6:31.336 | ית | ' |
| 6:31.337 | יגיע |  |
| 6:31.338 | תכליתן |  |
| 6:31.339 | והוא |  |
| 6:31.340 | אין |  |
| 6:31.341 | לו |  |
| 6:31.342 | נושא |  |
| 6:31.343 | ואין |  |
| 6:31.344 | מקיים |  |
| 6:31.345 | ית | ' |
| 6:31.346 | שמו |  |
| 6:31.347 | ויתעלה |  |
| 6:31.348 | זכרו |  |
| 6:31.349 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

