# Translation pack: milhamot-hashem perek-5-11 (Milchamot ha-Shem 5:18–5:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-11`.

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

- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
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
- **מורא** → "fear".
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
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
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


## Milchamot ha-Shem 5:18

Hebrew (Wikisource): הנה די מזה באר שהם עקרו ארבעה עיקרים הראשונים שביגדל אשר ביארתי לך בקונדריס שקדם אליך, ואל דבריהם אלה רמז ר' תם ן' יחיא באמרו כי תצא תקלה מזה לפרוץ גדרי התורה וחומותיה ויבא לידי מינות, והוא מה שנמנע הר' האר"י ושאר מקובלים מלומר יגדל, כי חשבו הקבלה החדשה לקבלה אמתית, ואת הקבלה אמתית של רז"ל הפשוטה בכל ישראל כמו שכתב הסמ"ג, והיא המבוארת ופשוטה גם כן בכתובים ובדברי רז"ל, חשבוה לטפלה וכוזבת. ואשר הביאם לידי מחשבה זו לעשות העיקר טפל והטפל עיקר הוא בעבור שראו חכמי ישראל כהרס"ג והרמב"ם ור' יהודה הלוי ור' בחיי בעל חובת הלבבות ושבילי אמונה וכיוצא בהם שהביאו ראיות חזקות על אמתת יחוד השי"ת בדברי הפילוסופים כדי לסתום את פי דוברי השקר הכופרים הטוענים נגד תורתינו הקדושה. ולכן חשבו אותה שהיא לקוחה מדברי הפילוסופים ואינה הקבלה האמיתית. ואת הרעיונות החדשים שהמציא הפילוסוף מחבר הזהר לפרש הכתובים שהם על דרך הרחבת לשון או מושאלים או משותפים כמו שכתב הרס"ג בספר האמונות והדעות והרמב"ם במורה נבוכים חלק א'. והוא נטה לפרשם כהנחתן הראשונה הפשוטה. יד ה', עיני ה', אף ה', אזני ה'. וכן: בנים אתם, אחי ורעי, בדמותי כצלמי, וכיוצא בהם, כהנחתן ומשמעותן הראשונה. וייחס רעיונות האלה בשקר לרשב"י התנא ושאר תנאים ואמוראים שלא היו בימי רשב"י אלא אחריו. ולכן חשבוה בטעות לקבלה אמתית. ולפיכך מיאנו לומר יגדל אלהים חי, כי האמינו וחשבו ליסוד מוסד אצלם כי העבודה והתפלה לאלוה קצר אפים (זעיר אנפין) וקודם הבריאה עוד לא נמצא הזעיר אנפין כמו שכתבתי לעיל בשם ספר הברית וכמבואר גם כן בעץ חיים וכסא אליהו ועז לאלהים. וגם אחר הבריאה עד אברהם אבינו עדיין לא נתקן הזעיר אנפין והיו כל אותן הדורות עד אברהם אבינו ניזונים ומקבלים שפעם מאבא ואימא כמו שכתב מה"ר יחיא הכהן בספרו חיי שלום פרשת חיי שרה בשם ליקוטי תורה פרשת לך לך. ולכך לא חפץ האר"י לומר יגדל אלהים חי כי איך יאמר נמצא ואין עת אל מציאותו, והרי יש עת אל מציאותו דהיינו מהבריאה ואילך ולא קודם לכן, או מאברהם ואילך שמהעת ההיא נשלם ונגמר תיקונו, כי קודם אברהם שעדיין לא נגמר תיקון הזעיר אנפין עדיין מקבלים שפעם מאבא ואימא ואין ראוי לעבוד לאלוה חסר התיקון, ולפי זה הרי יש עת אל מציאותו, דהיינו מהבריאה ואילך או מאברהם ואילך, והוה ליה שבח של שקר אם יאמר נמצא ואין עת אל מציאותו, שהרי קודם זה עדיין לא נמצא שהרי עדיין חסר תיקון ואין ראוי לעבודה. גם לא ערב לו לומר אחד ואין יחוד כיחודו, שהרי גם זה שקר אצלם כמו שנתבאר לעיל ממשל הבית שיש בו כמה חדרים ומקומות ועצים ואבנים וכל אבן חלוקה ומקובצת מכמה צרורות ועפר וסיד, והבנאי הוא שחיברם בחכמתו ועשאם דבר אחד והוא נקרא בשם אחד בית או כותל, נמצא שייחוד השי"ת אצלם כשאר האחדים שהם מחוברים מכמה חלקים והם נקראים בשם אחד, ואיך יאמר אחד ואין יחוד כיחודו! גם נמנע מלומר נעלם שהרי ביארו שיש כמה עולמות נאצלו ונבראו ונוצרו ונעשו למעלה מעולם האצילות, אכן לרוב העלמם לא שלחו בהם יד לבארם ולדבר בהם. וגם באדם קדמון כתבו שלא דיברו בו אלא ברמז קצת, והכלל אצלם שהפרצופים שלמעלה מעולם האצילות לרוב העלמם שהם זכים ודקים ונעלמים מאד לא שייכא תפילה וקריאה אליהם, ורק ביארו הפרצופים שכבר נתעבו ונראו קצת. ובהם דוקא שייך עבודה ותפלה וקרבנות. וזהו הטעם דלא שייך עבודה ותפלה באין סוף ובאדם קדמאה וכו' שבעולמות שלמעלה, ולא באדם קדמון דאצילות, וגם לא בשאר פרצופים שבאצילות לרוב העלמם ודקותם ורק שייך עבודה דוקא בזעיר אנפין פרצוף האחרון שבהם שכבר נתעבה ונתגלה יותר ונעשה בחינה המושגת כמו שכתב ספר הברית וכמו שכתב הרח"ו בעץ חיים ושער עיגולים ויושר ושער סדר אבי"ע ובשער השתלשלות העשר ספירות, ובכסא אליהו שער יראת ה' ועיין לקמן סי' פ"ד מה שכתבתי בשם משנת חסידים כי על כן רצו להשמיט נעלם ואין לו סוף לאחדותו שאיננו שבח אמתי לפי דעתם. גם הרחיקו מלומר אין לו דמות הגוף ואינו גוף. שהרי לפי דעתם יש לו דמות הגוף ברמ"ח איברים ושס"ה גידים כמבואר בכסא אליהו שער ראש חדש דף ל"ב ע"ב כי החמשה פרצופים כל אחד מהם פרצוף גמור בגילוי חתוך איברים רמ"ח איברים ושס"ה גידיים והביא ראיות לזה מדברי הזהר, וכן כתב הרח"ו בספר עץ חיים שער עיגולים ויושר כל הקו המתפשט מאין סוף בבחינת יושר הוא בציור אדם ישר כלול מרמ"ח איברים בסוד צלם אלהים. ועליה רמז הכתוב ויברא אלהים את האדם בצלמו בצלם אלהים ברא אותו, וכל דברי הזהר והתקונים בבחינת דיושר ע"ש, וכן כתב ספר הברית בחלק דברי אמת (מאמר א' פרק ז' וח') וכן כתב הראב"ד בפרק ד' מהלכות תשובה על מה שכתב הרמב"ם: "האומר שיש שם ריבון אחד אבל שהוא גוף ובעל תמונה", כתב הראב"ד: "ולמה קרא לזה מין וכמה גדולים וטובים ממנו הלכו בזו המחשבה לפי מה שראו במקראות" וכו'.

| id | word | punctuation after |
|---|---|---|
| 5:18.1 | הנה |  |
| 5:18.2 | די |  |
| 5:18.3 | מזה |  |
| 5:18.4 | באר |  |
| 5:18.5 | שהם |  |
| 5:18.6 | עקרו |  |
| 5:18.7 | ארבעה |  |
| 5:18.8 | עיקרים |  |
| 5:18.9 | הראשונים |  |
| 5:18.10 | שביגדל |  |
| 5:18.11 | אשר |  |
| 5:18.12 | ביארתי |  |
| 5:18.13 | לך |  |
| 5:18.14 | בקונדריס |  |
| 5:18.15 | שקדם |  |
| 5:18.16 | אליך | , |
| 5:18.17 | ואל |  |
| 5:18.18 | דבריהם |  |
| 5:18.19 | אלה |  |
| 5:18.20 | רמז |  |
| 5:18.21 | ר | ' |
| 5:18.22 | תם |  |
| 5:18.23 | ן | ' |
| 5:18.24 | יחיא |  |
| 5:18.25 | באמרו |  |
| 5:18.26 | כי |  |
| 5:18.27 | תצא |  |
| 5:18.28 | תקלה |  |
| 5:18.29 | מזה |  |
| 5:18.30 | לפרוץ |  |
| 5:18.31 | גדרי |  |
| 5:18.32 | התורה |  |
| 5:18.33 | וחומותיה |  |
| 5:18.34 | ויבא |  |
| 5:18.35 | לידי |  |
| 5:18.36 | מינות | , |
| 5:18.37 | והוא |  |
| 5:18.38 | מה |  |
| 5:18.39 | שנמנע |  |
| 5:18.40 | הר | ' |
| 5:18.41 | האר"י |  |
| 5:18.42 | ושאר |  |
| 5:18.43 | מקובלים |  |
| 5:18.44 | מלומר |  |
| 5:18.45 | יגדל | , |
| 5:18.46 | כי |  |
| 5:18.47 | חשבו |  |
| 5:18.48 | הקבלה |  |
| 5:18.49 | החדשה |  |
| 5:18.50 | לקבלה |  |
| 5:18.51 | אמתית | , |
| 5:18.52 | ואת |  |
| 5:18.53 | הקבלה |  |
| 5:18.54 | אמתית |  |
| 5:18.55 | של |  |
| 5:18.56 | רז"ל |  |
| 5:18.57 | הפשוטה |  |
| 5:18.58 | בכל |  |
| 5:18.59 | ישראל |  |
| 5:18.60 | כמו |  |
| 5:18.61 | שכתב |  |
| 5:18.62 | הסמ"ג | , |
| 5:18.63 | והיא |  |
| 5:18.64 | המבוארת |  |
| 5:18.65 | ופשוטה |  |
| 5:18.66 | גם |  |
| 5:18.67 | כן |  |
| 5:18.68 | בכתובים |  |
| 5:18.69 | ובדברי |  |
| 5:18.70 | רז"ל | , |
| 5:18.71 | חשבוה |  |
| 5:18.72 | לטפלה |  |
| 5:18.73 | וכוזבת | . |
| 5:18.74 | ואשר |  |
| 5:18.75 | הביאם |  |
| 5:18.76 | לידי |  |
| 5:18.77 | מחשבה |  |
| 5:18.78 | זו |  |
| 5:18.79 | לעשות |  |
| 5:18.80 | העיקר |  |
| 5:18.81 | טפל |  |
| 5:18.82 | והטפל |  |
| 5:18.83 | עיקר |  |
| 5:18.84 | הוא |  |
| 5:18.85 | בעבור |  |
| 5:18.86 | שראו |  |
| 5:18.87 | חכמי |  |
| 5:18.88 | ישראל |  |
| 5:18.89 | כהרס"ג |  |
| 5:18.90 | והרמב"ם |  |
| 5:18.91 | ור | ' |
| 5:18.92 | יהודה |  |
| 5:18.93 | הלוי |  |
| 5:18.94 | ור | ' |
| 5:18.95 | בחיי |  |
| 5:18.96 | בעל |  |
| 5:18.97 | חובת |  |
| 5:18.98 | הלבבות |  |
| 5:18.99 | ושבילי |  |
| 5:18.100 | אמונה |  |
| 5:18.101 | וכיוצא |  |
| 5:18.102 | בהם |  |
| 5:18.103 | שהביאו |  |
| 5:18.104 | ראיות |  |
| 5:18.105 | חזקות |  |
| 5:18.106 | על |  |
| 5:18.107 | אמתת |  |
| 5:18.108 | יחוד |  |
| 5:18.109 | השי"ת |  |
| 5:18.110 | בדברי |  |
| 5:18.111 | הפילוסופים |  |
| 5:18.112 | כדי |  |
| 5:18.113 | לסתום |  |
| 5:18.114 | את |  |
| 5:18.115 | פי |  |
| 5:18.116 | דוברי |  |
| 5:18.117 | השקר |  |
| 5:18.118 | הכופרים |  |
| 5:18.119 | הטוענים |  |
| 5:18.120 | נגד |  |
| 5:18.121 | תורתינו |  |
| 5:18.122 | הקדושה | . |
| 5:18.123 | ולכן |  |
| 5:18.124 | חשבו |  |
| 5:18.125 | אותה |  |
| 5:18.126 | שהיא |  |
| 5:18.127 | לקוחה |  |
| 5:18.128 | מדברי |  |
| 5:18.129 | הפילוסופים |  |
| 5:18.130 | ואינה |  |
| 5:18.131 | הקבלה |  |
| 5:18.132 | האמיתית | . |
| 5:18.133 | ואת |  |
| 5:18.134 | הרעיונות |  |
| 5:18.135 | החדשים |  |
| 5:18.136 | שהמציא |  |
| 5:18.137 | הפילוסוף |  |
| 5:18.138 | מחבר |  |
| 5:18.139 | הזהר |  |
| 5:18.140 | לפרש |  |
| 5:18.141 | הכתובים |  |
| 5:18.142 | שהם |  |
| 5:18.143 | על |  |
| 5:18.144 | דרך |  |
| 5:18.145 | הרחבת |  |
| 5:18.146 | לשון |  |
| 5:18.147 | או |  |
| 5:18.148 | מושאלים |  |
| 5:18.149 | או |  |
| 5:18.150 | משותפים |  |
| 5:18.151 | כמו |  |
| 5:18.152 | שכתב |  |
| 5:18.153 | הרס"ג |  |
| 5:18.154 | בספר |  |
| 5:18.155 | האמונות |  |
| 5:18.156 | והדעות |  |
| 5:18.157 | והרמב"ם |  |
| 5:18.158 | במורה |  |
| 5:18.159 | נבוכים |  |
| 5:18.160 | חלק |  |
| 5:18.161 | א | '. |
| 5:18.162 | והוא |  |
| 5:18.163 | נטה |  |
| 5:18.164 | לפרשם |  |
| 5:18.165 | כהנחתן |  |
| 5:18.166 | הראשונה |  |
| 5:18.167 | הפשוטה | . |
| 5:18.168 | יד |  |
| 5:18.169 | ה | ', |
| 5:18.170 | עיני |  |
| 5:18.171 | ה | ', |
| 5:18.172 | אף |  |
| 5:18.173 | ה | ', |
| 5:18.174 | אזני |  |
| 5:18.175 | ה | '. |
| 5:18.176 | וכן | : |
| 5:18.177 | בנים |  |
| 5:18.178 | אתם | , |
| 5:18.179 | אחי |  |
| 5:18.180 | ורעי | , |
| 5:18.181 | בדמותי |  |
| 5:18.182 | כצלמי | , |
| 5:18.183 | וכיוצא |  |
| 5:18.184 | בהם | , |
| 5:18.185 | כהנחתן |  |
| 5:18.186 | ומשמעותן |  |
| 5:18.187 | הראשונה | . |
| 5:18.188 | וייחס |  |
| 5:18.189 | רעיונות |  |
| 5:18.190 | האלה |  |
| 5:18.191 | בשקר |  |
| 5:18.192 | לרשב"י |  |
| 5:18.193 | התנא |  |
| 5:18.194 | ושאר |  |
| 5:18.195 | תנאים |  |
| 5:18.196 | ואמוראים |  |
| 5:18.197 | שלא |  |
| 5:18.198 | היו |  |
| 5:18.199 | בימי |  |
| 5:18.200 | רשב"י |  |
| 5:18.201 | אלא |  |
| 5:18.202 | אחריו | . |
| 5:18.203 | ולכן |  |
| 5:18.204 | חשבוה |  |
| 5:18.205 | בטעות |  |
| 5:18.206 | לקבלה |  |
| 5:18.207 | אמתית | . |
| 5:18.208 | ולפיכך |  |
| 5:18.209 | מיאנו |  |
| 5:18.210 | לומר |  |
| 5:18.211 | יגדל |  |
| 5:18.212 | אלהים |  |
| 5:18.213 | חי | , |
| 5:18.214 | כי |  |
| 5:18.215 | האמינו |  |
| 5:18.216 | וחשבו |  |
| 5:18.217 | ליסוד |  |
| 5:18.218 | מוסד |  |
| 5:18.219 | אצלם |  |
| 5:18.220 | כי |  |
| 5:18.221 | העבודה |  |
| 5:18.222 | והתפלה |  |
| 5:18.223 | לאלוה |  |
| 5:18.224 | קצר |  |
| 5:18.225 | אפים |  |
| 5:18.226 | זעיר |  |
| 5:18.227 | אנפין | ) |
| 5:18.228 | וקודם |  |
| 5:18.229 | הבריאה |  |
| 5:18.230 | עוד |  |
| 5:18.231 | לא |  |
| 5:18.232 | נמצא |  |
| 5:18.233 | הזעיר |  |
| 5:18.234 | אנפין |  |
| 5:18.235 | כמו |  |
| 5:18.236 | שכתבתי |  |
| 5:18.237 | לעיל |  |
| 5:18.238 | בשם |  |
| 5:18.239 | ספר |  |
| 5:18.240 | הברית |  |
| 5:18.241 | וכמבואר |  |
| 5:18.242 | גם |  |
| 5:18.243 | כן |  |
| 5:18.244 | בעץ |  |
| 5:18.245 | חיים |  |
| 5:18.246 | וכסא |  |
| 5:18.247 | אליהו |  |
| 5:18.248 | ועז |  |
| 5:18.249 | לאלהים | . |
| 5:18.250 | וגם |  |
| 5:18.251 | אחר |  |
| 5:18.252 | הבריאה |  |
| 5:18.253 | עד |  |
| 5:18.254 | אברהם |  |
| 5:18.255 | אבינו |  |
| 5:18.256 | עדיין |  |
| 5:18.257 | לא |  |
| 5:18.258 | נתקן |  |
| 5:18.259 | הזעיר |  |
| 5:18.260 | אנפין |  |
| 5:18.261 | והיו |  |
| 5:18.262 | כל |  |
| 5:18.263 | אותן |  |
| 5:18.264 | הדורות |  |
| 5:18.265 | עד |  |
| 5:18.266 | אברהם |  |
| 5:18.267 | אבינו |  |
| 5:18.268 | ניזונים |  |
| 5:18.269 | ומקבלים |  |
| 5:18.270 | שפעם |  |
| 5:18.271 | מאבא |  |
| 5:18.272 | ואימא |  |
| 5:18.273 | כמו |  |
| 5:18.274 | שכתב |  |
| 5:18.275 | מה"ר |  |
| 5:18.276 | יחיא |  |
| 5:18.277 | הכהן |  |
| 5:18.278 | בספרו |  |
| 5:18.279 | חיי |  |
| 5:18.280 | שלום |  |
| 5:18.281 | פרשת |  |
| 5:18.282 | חיי |  |
| 5:18.283 | שרה |  |
| 5:18.284 | בשם |  |
| 5:18.285 | ליקוטי |  |
| 5:18.286 | תורה |  |
| 5:18.287 | פרשת |  |
| 5:18.288 | לך |  |
| 5:18.289 | לך | . |
| 5:18.290 | ולכך |  |
| 5:18.291 | לא |  |
| 5:18.292 | חפץ |  |
| 5:18.293 | האר"י |  |
| 5:18.294 | לומר |  |
| 5:18.295 | יגדל |  |
| 5:18.296 | אלהים |  |
| 5:18.297 | חי |  |
| 5:18.298 | כי |  |
| 5:18.299 | איך |  |
| 5:18.300 | יאמר |  |
| 5:18.301 | נמצא |  |
| 5:18.302 | ואין |  |
| 5:18.303 | עת |  |
| 5:18.304 | אל |  |
| 5:18.305 | מציאותו | , |
| 5:18.306 | והרי |  |
| 5:18.307 | יש |  |
| 5:18.308 | עת |  |
| 5:18.309 | אל |  |
| 5:18.310 | מציאותו |  |
| 5:18.311 | דהיינו |  |
| 5:18.312 | מהבריאה |  |
| 5:18.313 | ואילך |  |
| 5:18.314 | ולא |  |
| 5:18.315 | קודם |  |
| 5:18.316 | לכן | , |
| 5:18.317 | או |  |
| 5:18.318 | מאברהם |  |
| 5:18.319 | ואילך |  |
| 5:18.320 | שמהעת |  |
| 5:18.321 | ההיא |  |
| 5:18.322 | נשלם |  |
| 5:18.323 | ונגמר |  |
| 5:18.324 | תיקונו | , |
| 5:18.325 | כי |  |
| 5:18.326 | קודם |  |
| 5:18.327 | אברהם |  |
| 5:18.328 | שעדיין |  |
| 5:18.329 | לא |  |
| 5:18.330 | נגמר |  |
| 5:18.331 | תיקון |  |
| 5:18.332 | הזעיר |  |
| 5:18.333 | אנפין |  |
| 5:18.334 | עדיין |  |
| 5:18.335 | מקבלים |  |
| 5:18.336 | שפעם |  |
| 5:18.337 | מאבא |  |
| 5:18.338 | ואימא |  |
| 5:18.339 | ואין |  |
| 5:18.340 | ראוי |  |
| 5:18.341 | לעבוד |  |
| 5:18.342 | לאלוה |  |
| 5:18.343 | חסר |  |
| 5:18.344 | התיקון | , |
| 5:18.345 | ולפי |  |
| 5:18.346 | זה |  |
| 5:18.347 | הרי |  |
| 5:18.348 | יש |  |
| 5:18.349 | עת |  |
| 5:18.350 | אל |  |
| 5:18.351 | מציאותו | , |
| 5:18.352 | דהיינו |  |
| 5:18.353 | מהבריאה |  |
| 5:18.354 | ואילך |  |
| 5:18.355 | או |  |
| 5:18.356 | מאברהם |  |
| 5:18.357 | ואילך | , |
| 5:18.358 | והוה |  |
| 5:18.359 | ליה |  |
| 5:18.360 | שבח |  |
| 5:18.361 | של |  |
| 5:18.362 | שקר |  |
| 5:18.363 | אם |  |
| 5:18.364 | יאמר |  |
| 5:18.365 | נמצא |  |
| 5:18.366 | ואין |  |
| 5:18.367 | עת |  |
| 5:18.368 | אל |  |
| 5:18.369 | מציאותו | , |
| 5:18.370 | שהרי |  |
| 5:18.371 | קודם |  |
| 5:18.372 | זה |  |
| 5:18.373 | עדיין |  |
| 5:18.374 | לא |  |
| 5:18.375 | נמצא |  |
| 5:18.376 | שהרי |  |
| 5:18.377 | עדיין |  |
| 5:18.378 | חסר |  |
| 5:18.379 | תיקון |  |
| 5:18.380 | ואין |  |
| 5:18.381 | ראוי |  |
| 5:18.382 | לעבודה | . |
| 5:18.383 | גם |  |
| 5:18.384 | לא |  |
| 5:18.385 | ערב |  |
| 5:18.386 | לו |  |
| 5:18.387 | לומר |  |
| 5:18.388 | אחד |  |
| 5:18.389 | ואין |  |
| 5:18.390 | יחוד |  |
| 5:18.391 | כיחודו | , |
| 5:18.392 | שהרי |  |
| 5:18.393 | גם |  |
| 5:18.394 | זה |  |
| 5:18.395 | שקר |  |
| 5:18.396 | אצלם |  |
| 5:18.397 | כמו |  |
| 5:18.398 | שנתבאר |  |
| 5:18.399 | לעיל |  |
| 5:18.400 | ממשל |  |
| 5:18.401 | הבית |  |
| 5:18.402 | שיש |  |
| 5:18.403 | בו |  |
| 5:18.404 | כמה |  |
| 5:18.405 | חדרים |  |
| 5:18.406 | ומקומות |  |
| 5:18.407 | ועצים |  |
| 5:18.408 | ואבנים |  |
| 5:18.409 | וכל |  |
| 5:18.410 | אבן |  |
| 5:18.411 | חלוקה |  |
| 5:18.412 | ומקובצת |  |
| 5:18.413 | מכמה |  |
| 5:18.414 | צרורות |  |
| 5:18.415 | ועפר |  |
| 5:18.416 | וסיד | , |
| 5:18.417 | והבנאי |  |
| 5:18.418 | הוא |  |
| 5:18.419 | שחיברם |  |
| 5:18.420 | בחכמתו |  |
| 5:18.421 | ועשאם |  |
| 5:18.422 | דבר |  |
| 5:18.423 | אחד |  |
| 5:18.424 | והוא |  |
| 5:18.425 | נקרא |  |
| 5:18.426 | בשם |  |
| 5:18.427 | אחד |  |
| 5:18.428 | בית |  |
| 5:18.429 | או |  |
| 5:18.430 | כותל | , |
| 5:18.431 | נמצא |  |
| 5:18.432 | שייחוד |  |
| 5:18.433 | השי"ת |  |
| 5:18.434 | אצלם |  |
| 5:18.435 | כשאר |  |
| 5:18.436 | האחדים |  |
| 5:18.437 | שהם |  |
| 5:18.438 | מחוברים |  |
| 5:18.439 | מכמה |  |
| 5:18.440 | חלקים |  |
| 5:18.441 | והם |  |
| 5:18.442 | נקראים |  |
| 5:18.443 | בשם |  |
| 5:18.444 | אחד | , |
| 5:18.445 | ואיך |  |
| 5:18.446 | יאמר |  |
| 5:18.447 | אחד |  |
| 5:18.448 | ואין |  |
| 5:18.449 | יחוד |  |
| 5:18.450 | כיחודו | ! |
| 5:18.451 | גם |  |
| 5:18.452 | נמנע |  |
| 5:18.453 | מלומר |  |
| 5:18.454 | נעלם |  |
| 5:18.455 | שהרי |  |
| 5:18.456 | ביארו |  |
| 5:18.457 | שיש |  |
| 5:18.458 | כמה |  |
| 5:18.459 | עולמות |  |
| 5:18.460 | נאצלו |  |
| 5:18.461 | ונבראו |  |
| 5:18.462 | ונוצרו |  |
| 5:18.463 | ונעשו |  |
| 5:18.464 | למעלה |  |
| 5:18.465 | מעולם |  |
| 5:18.466 | האצילות | , |
| 5:18.467 | אכן |  |
| 5:18.468 | לרוב |  |
| 5:18.469 | העלמם |  |
| 5:18.470 | לא |  |
| 5:18.471 | שלחו |  |
| 5:18.472 | בהם |  |
| 5:18.473 | יד |  |
| 5:18.474 | לבארם |  |
| 5:18.475 | ולדבר |  |
| 5:18.476 | בהם | . |
| 5:18.477 | וגם |  |
| 5:18.478 | באדם |  |
| 5:18.479 | קדמון |  |
| 5:18.480 | כתבו |  |
| 5:18.481 | שלא |  |
| 5:18.482 | דיברו |  |
| 5:18.483 | בו |  |
| 5:18.484 | אלא |  |
| 5:18.485 | ברמז |  |
| 5:18.486 | קצת | , |
| 5:18.487 | והכלל |  |
| 5:18.488 | אצלם |  |
| 5:18.489 | שהפרצופים |  |
| 5:18.490 | שלמעלה |  |
| 5:18.491 | מעולם |  |
| 5:18.492 | האצילות |  |
| 5:18.493 | לרוב |  |
| 5:18.494 | העלמם |  |
| 5:18.495 | שהם |  |
| 5:18.496 | זכים |  |
| 5:18.497 | ודקים |  |
| 5:18.498 | ונעלמים |  |
| 5:18.499 | מאד |  |
| 5:18.500 | לא |  |
| 5:18.501 | שייכא |  |
| 5:18.502 | תפילה |  |
| 5:18.503 | וקריאה |  |
| 5:18.504 | אליהם | , |
| 5:18.505 | ורק |  |
| 5:18.506 | ביארו |  |
| 5:18.507 | הפרצופים |  |
| 5:18.508 | שכבר |  |
| 5:18.509 | נתעבו |  |
| 5:18.510 | ונראו |  |
| 5:18.511 | קצת | . |
| 5:18.512 | ובהם |  |
| 5:18.513 | דוקא |  |
| 5:18.514 | שייך |  |
| 5:18.515 | עבודה |  |
| 5:18.516 | ותפלה |  |
| 5:18.517 | וקרבנות | . |
| 5:18.518 | וזהו |  |
| 5:18.519 | הטעם |  |
| 5:18.520 | דלא |  |
| 5:18.521 | שייך |  |
| 5:18.522 | עבודה |  |
| 5:18.523 | ותפלה |  |
| 5:18.524 | באין |  |
| 5:18.525 | סוף |  |
| 5:18.526 | ובאדם |  |
| 5:18.527 | קדמאה |  |
| 5:18.528 | וכו | ' |
| 5:18.529 | שבעולמות |  |
| 5:18.530 | שלמעלה | , |
| 5:18.531 | ולא |  |
| 5:18.532 | באדם |  |
| 5:18.533 | קדמון |  |
| 5:18.534 | דאצילות | , |
| 5:18.535 | וגם |  |
| 5:18.536 | לא |  |
| 5:18.537 | בשאר |  |
| 5:18.538 | פרצופים |  |
| 5:18.539 | שבאצילות |  |
| 5:18.540 | לרוב |  |
| 5:18.541 | העלמם |  |
| 5:18.542 | ודקותם |  |
| 5:18.543 | ורק |  |
| 5:18.544 | שייך |  |
| 5:18.545 | עבודה |  |
| 5:18.546 | דוקא |  |
| 5:18.547 | בזעיר |  |
| 5:18.548 | אנפין |  |
| 5:18.549 | פרצוף |  |
| 5:18.550 | האחרון |  |
| 5:18.551 | שבהם |  |
| 5:18.552 | שכבר |  |
| 5:18.553 | נתעבה |  |
| 5:18.554 | ונתגלה |  |
| 5:18.555 | יותר |  |
| 5:18.556 | ונעשה |  |
| 5:18.557 | בחינה |  |
| 5:18.558 | המושגת |  |
| 5:18.559 | כמו |  |
| 5:18.560 | שכתב |  |
| 5:18.561 | ספר |  |
| 5:18.562 | הברית |  |
| 5:18.563 | וכמו |  |
| 5:18.564 | שכתב |  |
| 5:18.565 | הרח"ו |  |
| 5:18.566 | בעץ |  |
| 5:18.567 | חיים |  |
| 5:18.568 | ושער |  |
| 5:18.569 | עיגולים |  |
| 5:18.570 | ויושר |  |
| 5:18.571 | ושער |  |
| 5:18.572 | סדר |  |
| 5:18.573 | אבי"ע |  |
| 5:18.574 | ובשער |  |
| 5:18.575 | השתלשלות |  |
| 5:18.576 | העשר |  |
| 5:18.577 | ספירות | , |
| 5:18.578 | ובכסא |  |
| 5:18.579 | אליהו |  |
| 5:18.580 | שער |  |
| 5:18.581 | יראת |  |
| 5:18.582 | ה | ' |
| 5:18.583 | ועיין |  |
| 5:18.584 | לקמן |  |
| 5:18.585 | סי | ' |
| 5:18.586 | פ"ד |  |
| 5:18.587 | מה |  |
| 5:18.588 | שכתבתי |  |
| 5:18.589 | בשם |  |
| 5:18.590 | משנת |  |
| 5:18.591 | חסידים |  |
| 5:18.592 | כי |  |
| 5:18.593 | על |  |
| 5:18.594 | כן |  |
| 5:18.595 | רצו |  |
| 5:18.596 | להשמיט |  |
| 5:18.597 | נעלם |  |
| 5:18.598 | ואין |  |
| 5:18.599 | לו |  |
| 5:18.600 | סוף |  |
| 5:18.601 | לאחדותו |  |
| 5:18.602 | שאיננו |  |
| 5:18.603 | שבח |  |
| 5:18.604 | אמתי |  |
| 5:18.605 | לפי |  |
| 5:18.606 | דעתם | . |
| 5:18.607 | גם |  |
| 5:18.608 | הרחיקו |  |
| 5:18.609 | מלומר |  |
| 5:18.610 | אין |  |
| 5:18.611 | לו |  |
| 5:18.612 | דמות |  |
| 5:18.613 | הגוף |  |
| 5:18.614 | ואינו |  |
| 5:18.615 | גוף | . |
| 5:18.616 | שהרי |  |
| 5:18.617 | לפי |  |
| 5:18.618 | דעתם |  |
| 5:18.619 | יש |  |
| 5:18.620 | לו |  |
| 5:18.621 | דמות |  |
| 5:18.622 | הגוף |  |
| 5:18.623 | ברמ"ח |  |
| 5:18.624 | איברים |  |
| 5:18.625 | ושס"ה |  |
| 5:18.626 | גידים |  |
| 5:18.627 | כמבואר |  |
| 5:18.628 | בכסא |  |
| 5:18.629 | אליהו |  |
| 5:18.630 | שער |  |
| 5:18.631 | ראש |  |
| 5:18.632 | חדש |  |
| 5:18.633 | דף |  |
| 5:18.634 | ל"ב |  |
| 5:18.635 | ע"ב |  |
| 5:18.636 | כי |  |
| 5:18.637 | החמשה |  |
| 5:18.638 | פרצופים |  |
| 5:18.639 | כל |  |
| 5:18.640 | אחד |  |
| 5:18.641 | מהם |  |
| 5:18.642 | פרצוף |  |
| 5:18.643 | גמור |  |
| 5:18.644 | בגילוי |  |
| 5:18.645 | חתוך |  |
| 5:18.646 | איברים |  |
| 5:18.647 | רמ"ח |  |
| 5:18.648 | איברים |  |
| 5:18.649 | ושס"ה |  |
| 5:18.650 | גידיים |  |
| 5:18.651 | והביא |  |
| 5:18.652 | ראיות |  |
| 5:18.653 | לזה |  |
| 5:18.654 | מדברי |  |
| 5:18.655 | הזהר | , |
| 5:18.656 | וכן |  |
| 5:18.657 | כתב |  |
| 5:18.658 | הרח"ו |  |
| 5:18.659 | בספר |  |
| 5:18.660 | עץ |  |
| 5:18.661 | חיים |  |
| 5:18.662 | שער |  |
| 5:18.663 | עיגולים |  |
| 5:18.664 | ויושר |  |
| 5:18.665 | כל |  |
| 5:18.666 | הקו |  |
| 5:18.667 | המתפשט |  |
| 5:18.668 | מאין |  |
| 5:18.669 | סוף |  |
| 5:18.670 | בבחינת |  |
| 5:18.671 | יושר |  |
| 5:18.672 | הוא |  |
| 5:18.673 | בציור |  |
| 5:18.674 | אדם |  |
| 5:18.675 | ישר |  |
| 5:18.676 | כלול |  |
| 5:18.677 | מרמ"ח |  |
| 5:18.678 | איברים |  |
| 5:18.679 | בסוד |  |
| 5:18.680 | צלם |  |
| 5:18.681 | אלהים | . |
| 5:18.682 | ועליה |  |
| 5:18.683 | רמז |  |
| 5:18.684 | הכתוב |  |
| 5:18.685 | ויברא |  |
| 5:18.686 | אלהים |  |
| 5:18.687 | את |  |
| 5:18.688 | האדם |  |
| 5:18.689 | בצלמו |  |
| 5:18.690 | בצלם |  |
| 5:18.691 | אלהים |  |
| 5:18.692 | ברא |  |
| 5:18.693 | אותו | , |
| 5:18.694 | וכל |  |
| 5:18.695 | דברי |  |
| 5:18.696 | הזהר |  |
| 5:18.697 | והתקונים |  |
| 5:18.698 | בבחינת |  |
| 5:18.699 | דיושר |  |
| 5:18.700 | ע"ש | , |
| 5:18.701 | וכן |  |
| 5:18.702 | כתב |  |
| 5:18.703 | ספר |  |
| 5:18.704 | הברית |  |
| 5:18.705 | בחלק |  |
| 5:18.706 | דברי |  |
| 5:18.707 | אמת |  |
| 5:18.708 | מאמר |  |
| 5:18.709 | א | ' |
| 5:18.710 | פרק |  |
| 5:18.711 | ז | ' |
| 5:18.712 | וח | ') |
| 5:18.713 | וכן |  |
| 5:18.714 | כתב |  |
| 5:18.715 | הראב"ד |  |
| 5:18.716 | בפרק |  |
| 5:18.717 | ד | ' |
| 5:18.718 | מהלכות |  |
| 5:18.719 | תשובה |  |
| 5:18.720 | על |  |
| 5:18.721 | מה |  |
| 5:18.722 | שכתב |  |
| 5:18.723 | הרמב"ם | : |
| 5:18.724 | האומר |  |
| 5:18.725 | שיש |  |
| 5:18.726 | שם |  |
| 5:18.727 | ריבון |  |
| 5:18.728 | אחד |  |
| 5:18.729 | אבל |  |
| 5:18.730 | שהוא |  |
| 5:18.731 | גוף |  |
| 5:18.732 | ובעל |  |
| 5:18.733 | תמונה | ", |
| 5:18.734 | כתב |  |
| 5:18.735 | הראב"ד | : |
| 5:18.736 | ולמה |  |
| 5:18.737 | קרא |  |
| 5:18.738 | לזה |  |
| 5:18.739 | מין |  |
| 5:18.740 | וכמה |  |
| 5:18.741 | גדולים |  |
| 5:18.742 | וטובים |  |
| 5:18.743 | ממנו |  |
| 5:18.744 | הלכו |  |
| 5:18.745 | בזו |  |
| 5:18.746 | המחשבה |  |
| 5:18.747 | לפי |  |
| 5:18.748 | מה |  |
| 5:18.749 | שראו |  |
| 5:18.750 | במקראות | " |
| 5:18.751 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

