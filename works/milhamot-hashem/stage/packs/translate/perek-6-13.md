# Translation pack: milhamot-hashem perek-6-13 (Milchamot ha-Shem 6:20–6:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-13`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
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


## Milchamot ha-Shem 6:20

Hebrew (Wikisource): ובהקדמת תקוניה הזהר דף ג' ע"ב, גם בתחלת הקדמה שנייה ובסוף תיקונא שתיתאה כ' וכמה בני נשא יתפרנסון מהאי חבורה דילך בסוף יומיא וכו'. ומשפתי חלקות מסית ומדיח לזלזל בלימוד משנה ותלמוד, וחשב עסקם לעון אשר חטא עד שצריך תשובה לעזוב אותם ולהתעסק בתורתו החדשה המלאה תארים ותבניות צורות ופרצופים רבים שונים זה מזה אשר עלה על רוחו לעשותם גופים, ולהוריד להם נשמות מעצמות האלהים העליון שאין לו דמות הגוף בידים ורגלים ושאר אברי הגוף רק דמות עגול גדול מקיף את כל החלל שבתוכו לפי דעתו ודרך צינור דק משך כמנו הפילסוף המתעה נשמות ומלבושים לכסות מערומי הצורות והפרצופים כמבואר בעץ חיים להרח"ו, ובשאר ספרי המקבלים ממנו אשר לא שמו אל לבם דברי רז"ל בגמרא ומדרש רבה ותנחומא וכיוצא מספרי רז"ל האמתיים.

| id | word | punctuation after |
|---|---|---|
| 6:20.1 | ובהקדמת |  |
| 6:20.2 | תקוניה |  |
| 6:20.3 | הזהר |  |
| 6:20.4 | דף |  |
| 6:20.5 | ג | ' |
| 6:20.6 | ע"ב | , |
| 6:20.7 | גם |  |
| 6:20.8 | בתחלת |  |
| 6:20.9 | הקדמה |  |
| 6:20.10 | שנייה |  |
| 6:20.11 | ובסוף |  |
| 6:20.12 | תיקונא |  |
| 6:20.13 | שתיתאה |  |
| 6:20.14 | כ | ' |
| 6:20.15 | וכמה |  |
| 6:20.16 | בני |  |
| 6:20.17 | נשא |  |
| 6:20.18 | יתפרנסון |  |
| 6:20.19 | מהאי |  |
| 6:20.20 | חבורה |  |
| 6:20.21 | דילך |  |
| 6:20.22 | בסוף |  |
| 6:20.23 | יומיא |  |
| 6:20.24 | וכו | '. |
| 6:20.25 | ומשפתי |  |
| 6:20.26 | חלקות |  |
| 6:20.27 | מסית |  |
| 6:20.28 | ומדיח |  |
| 6:20.29 | לזלזל |  |
| 6:20.30 | בלימוד |  |
| 6:20.31 | משנה |  |
| 6:20.32 | ותלמוד | , |
| 6:20.33 | וחשב |  |
| 6:20.34 | עסקם |  |
| 6:20.35 | לעון |  |
| 6:20.36 | אשר |  |
| 6:20.37 | חטא |  |
| 6:20.38 | עד |  |
| 6:20.39 | שצריך |  |
| 6:20.40 | תשובה |  |
| 6:20.41 | לעזוב |  |
| 6:20.42 | אותם |  |
| 6:20.43 | ולהתעסק |  |
| 6:20.44 | בתורתו |  |
| 6:20.45 | החדשה |  |
| 6:20.46 | המלאה |  |
| 6:20.47 | תארים |  |
| 6:20.48 | ותבניות |  |
| 6:20.49 | צורות |  |
| 6:20.50 | ופרצופים |  |
| 6:20.51 | רבים |  |
| 6:20.52 | שונים |  |
| 6:20.53 | זה |  |
| 6:20.54 | מזה |  |
| 6:20.55 | אשר |  |
| 6:20.56 | עלה |  |
| 6:20.57 | על |  |
| 6:20.58 | רוחו |  |
| 6:20.59 | לעשותם |  |
| 6:20.60 | גופים | , |
| 6:20.61 | ולהוריד |  |
| 6:20.62 | להם |  |
| 6:20.63 | נשמות |  |
| 6:20.64 | מעצמות |  |
| 6:20.65 | האלהים |  |
| 6:20.66 | העליון |  |
| 6:20.67 | שאין |  |
| 6:20.68 | לו |  |
| 6:20.69 | דמות |  |
| 6:20.70 | הגוף |  |
| 6:20.71 | בידים |  |
| 6:20.72 | ורגלים |  |
| 6:20.73 | ושאר |  |
| 6:20.74 | אברי |  |
| 6:20.75 | הגוף |  |
| 6:20.76 | רק |  |
| 6:20.77 | דמות |  |
| 6:20.78 | עגול |  |
| 6:20.79 | גדול |  |
| 6:20.80 | מקיף |  |
| 6:20.81 | את |  |
| 6:20.82 | כל |  |
| 6:20.83 | החלל |  |
| 6:20.84 | שבתוכו |  |
| 6:20.85 | לפי |  |
| 6:20.86 | דעתו |  |
| 6:20.87 | ודרך |  |
| 6:20.88 | צינור |  |
| 6:20.89 | דק |  |
| 6:20.90 | משך |  |
| 6:20.91 | כמנו |  |
| 6:20.92 | הפילסוף |  |
| 6:20.93 | המתעה |  |
| 6:20.94 | נשמות |  |
| 6:20.95 | ומלבושים |  |
| 6:20.96 | לכסות |  |
| 6:20.97 | מערומי |  |
| 6:20.98 | הצורות |  |
| 6:20.99 | והפרצופים |  |
| 6:20.100 | כמבואר |  |
| 6:20.101 | בעץ |  |
| 6:20.102 | חיים |  |
| 6:20.103 | להרח"ו | , |
| 6:20.104 | ובשאר |  |
| 6:20.105 | ספרי |  |
| 6:20.106 | המקבלים |  |
| 6:20.107 | ממנו |  |
| 6:20.108 | אשר |  |
| 6:20.109 | לא |  |
| 6:20.110 | שמו |  |
| 6:20.111 | אל |  |
| 6:20.112 | לבם |  |
| 6:20.113 | דברי |  |
| 6:20.114 | רז"ל |  |
| 6:20.115 | בגמרא |  |
| 6:20.116 | ומדרש |  |
| 6:20.117 | רבה |  |
| 6:20.118 | ותנחומא |  |
| 6:20.119 | וכיוצא |  |
| 6:20.120 | מספרי |  |
| 6:20.121 | רז"ל |  |
| 6:20.122 | האמתיים | . |


## Milchamot ha-Shem 6:21

Hebrew (Wikisource): ובתיקוני זהר חדש (דף ו' וז' ודף י') הרחיב עוד את פיו לספר בגנות העוסקים במשנה ובתלמוד כדי שיקרא רבי או כדי להשיג עושר בעה"ז או שיזכה לחיי העה"ב, שאמרו חכמים עליהם לעולם יעסוק אדם בתורה ובמצות אפי' שלא לשמה, שמתוך וכו' וגם אמרו האומר סלע זו לצדקה בשביל שיחיו בני בשביל שאזכה לחיי העולם הבא הרי זה צדיק גמור. והפילסוף מחבר הזהר שם אותם ליורשי גהינם, ודימה שקלא וטריא דרבנן העוסקים בתורה כדי שיזכו לחיי העולם הבא לכלבים המנבחים ואומרים חף חף הב הב. ואמר שהם יורשים גיהנם האומרת הב הב. כמו שמסיים שם שעליהם נאמר ואלה לחרפות ולדראון עולם. ולעג על מאמר התנא למוד תורה הרבה ויתנו לך שכר הרבה. ואף אם עוסק בתורה להבדיל בין הטהור ובין הטמא כאשר ציונו ה' אלהינו להבדיל בין החיה הנאכלת ובין החיה אשר לא תאכל, ושאר איסורי מאכלות הכל נקרא אצלו עוסק שלא לשמה כמו שאמר שם - לא משתדלין אלא לגרמייהו לדבתא בשרא וכו'. ולא נקרא עוסק לשמה אלא העוסק בתורה כדי לחבר האלהות הזכרים בנקבות, אבא עם אימא, וזעיר עם נוקביה. לזוגם יחד והיו לבשר אחד. וכן בקריאת שמע וציצית ותפילין ושאר כל המצות. ולכן תקנו המקובלים לומר קודם כל מצוה שעושים לשם יחוד קב"ה ושכינתיה וכו', כי תכלית כוונת הפילסוף להכניס בישראל אמונת רבוי באלוהות זכרים ונקבות ולעשות חבורים שלו שהם הזהר והתיקונים עיקר לעסוק בהם חלף המשנה והתלמוד. ולשכח מישראל עסק המשנה והתלמוד שכל כללותיה ופרטותיה של תורה כלולים בהם. ובכמה מקומות הוא משבח את חבורו זה ומהם בתיקוני זוהר חדש דף י"ו ע"ב. ובנפת אשר יטיף הפילסוף הנז' בשפתי אמונה שהיא זרה לאמונת תורתינו הקדושה משך לב רבים בחלקלקותיו, וישיאם לעבוד אלהים אחרים המה הצורות והפרצופים אשר בדה מלבו להטעותנו בפירוש עשר ספירות הנז' בספר יצירה לחשבם אלהות ולעבוד אותם בעבור שנשמותיהם חצובים מהאלהים העליון (לפי דעתו) אשר חלק להם מעצמותו ומהותו הם ועיגוליהם וגופם ונשמתם ולבושם, כולם חוצבו ממקור אחד (הוא הא"ם) אשר לפי דעתם אין לו שם ונקודה ולית מחשבה תפיסה ביה, ולא שייך שום עבודה וקריאה אליו, ולא יענה את הקורא אליו שכבר מינה והפקיד את הפרצופים על כל מלאכת הבריאה. ועשה את אבא בורא, והוא אומר ומצוה לאימא יהי כך וכך, והיא עושה כדבריו, כאומן שעושה מלאכת בעל הבית באמונה. ואבא ואמא גם הם צוו והפקידו את זעיר אנפין ונוקבתיה להשגיח ולרון על מעשה התחתונים. מאן דזכי לדינא לדינא, ומאן דזכי לרחמי לרחמי, הכל נעשה כחפצו ורצונו של זעיר אנפין. וכל דברי המחבר הזה בנוים על הפילוסופיא הקדומה אשר כבר נמר ריחה וטעמה מעת אשר נגלה כבודו ית' על הר סיני ואמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים. רק שאחרי מות יהושע והזקנים חזרו להדיעות האלו הנפסדות ויעבדו את הבעלים והעשתרות האלו ויעשו צלמיהם כדי להוריד השפע מהן על אותן הצורות והצלמים לפי דעתם.

| id | word | punctuation after |
|---|---|---|
| 6:21.1 | ובתיקוני |  |
| 6:21.2 | זהר |  |
| 6:21.3 | חדש |  |
| 6:21.4 | דף |  |
| 6:21.5 | ו | ' |
| 6:21.6 | וז | ' |
| 6:21.7 | ודף |  |
| 6:21.8 | י | ') |
| 6:21.9 | הרחיב |  |
| 6:21.10 | עוד |  |
| 6:21.11 | את |  |
| 6:21.12 | פיו |  |
| 6:21.13 | לספר |  |
| 6:21.14 | בגנות |  |
| 6:21.15 | העוסקים |  |
| 6:21.16 | במשנה |  |
| 6:21.17 | ובתלמוד |  |
| 6:21.18 | כדי |  |
| 6:21.19 | שיקרא |  |
| 6:21.20 | רבי |  |
| 6:21.21 | או |  |
| 6:21.22 | כדי |  |
| 6:21.23 | להשיג |  |
| 6:21.24 | עושר |  |
| 6:21.25 | בעה"ז |  |
| 6:21.26 | או |  |
| 6:21.27 | שיזכה |  |
| 6:21.28 | לחיי |  |
| 6:21.29 | העה"ב | , |
| 6:21.30 | שאמרו |  |
| 6:21.31 | חכמים |  |
| 6:21.32 | עליהם |  |
| 6:21.33 | לעולם |  |
| 6:21.34 | יעסוק |  |
| 6:21.35 | אדם |  |
| 6:21.36 | בתורה |  |
| 6:21.37 | ובמצות |  |
| 6:21.38 | אפי | ' |
| 6:21.39 | שלא |  |
| 6:21.40 | לשמה | , |
| 6:21.41 | שמתוך |  |
| 6:21.42 | וכו | ' |
| 6:21.43 | וגם |  |
| 6:21.44 | אמרו |  |
| 6:21.45 | האומר |  |
| 6:21.46 | סלע |  |
| 6:21.47 | זו |  |
| 6:21.48 | לצדקה |  |
| 6:21.49 | בשביל |  |
| 6:21.50 | שיחיו |  |
| 6:21.51 | בני |  |
| 6:21.52 | בשביל |  |
| 6:21.53 | שאזכה |  |
| 6:21.54 | לחיי |  |
| 6:21.55 | העולם |  |
| 6:21.56 | הבא |  |
| 6:21.57 | הרי |  |
| 6:21.58 | זה |  |
| 6:21.59 | צדיק |  |
| 6:21.60 | גמור | . |
| 6:21.61 | והפילסוף |  |
| 6:21.62 | מחבר |  |
| 6:21.63 | הזהר |  |
| 6:21.64 | שם |  |
| 6:21.65 | אותם |  |
| 6:21.66 | ליורשי |  |
| 6:21.67 | גהינם | , |
| 6:21.68 | ודימה |  |
| 6:21.69 | שקלא |  |
| 6:21.70 | וטריא |  |
| 6:21.71 | דרבנן |  |
| 6:21.72 | העוסקים |  |
| 6:21.73 | בתורה |  |
| 6:21.74 | כדי |  |
| 6:21.75 | שיזכו |  |
| 6:21.76 | לחיי |  |
| 6:21.77 | העולם |  |
| 6:21.78 | הבא |  |
| 6:21.79 | לכלבים |  |
| 6:21.80 | המנבחים |  |
| 6:21.81 | ואומרים |  |
| 6:21.82 | חף |  |
| 6:21.83 | חף |  |
| 6:21.84 | הב |  |
| 6:21.85 | הב | . |
| 6:21.86 | ואמר |  |
| 6:21.87 | שהם |  |
| 6:21.88 | יורשים |  |
| 6:21.89 | גיהנם |  |
| 6:21.90 | האומרת |  |
| 6:21.91 | הב |  |
| 6:21.92 | הב | . |
| 6:21.93 | כמו |  |
| 6:21.94 | שמסיים |  |
| 6:21.95 | שם |  |
| 6:21.96 | שעליהם |  |
| 6:21.97 | נאמר |  |
| 6:21.98 | ואלה |  |
| 6:21.99 | לחרפות |  |
| 6:21.100 | ולדראון |  |
| 6:21.101 | עולם | . |
| 6:21.102 | ולעג |  |
| 6:21.103 | על |  |
| 6:21.104 | מאמר |  |
| 6:21.105 | התנא |  |
| 6:21.106 | למוד |  |
| 6:21.107 | תורה |  |
| 6:21.108 | הרבה |  |
| 6:21.109 | ויתנו |  |
| 6:21.110 | לך |  |
| 6:21.111 | שכר |  |
| 6:21.112 | הרבה | . |
| 6:21.113 | ואף |  |
| 6:21.114 | אם |  |
| 6:21.115 | עוסק |  |
| 6:21.116 | בתורה |  |
| 6:21.117 | להבדיל |  |
| 6:21.118 | בין |  |
| 6:21.119 | הטהור |  |
| 6:21.120 | ובין |  |
| 6:21.121 | הטמא |  |
| 6:21.122 | כאשר |  |
| 6:21.123 | ציונו |  |
| 6:21.124 | ה | ' |
| 6:21.125 | אלהינו |  |
| 6:21.126 | להבדיל |  |
| 6:21.127 | בין |  |
| 6:21.128 | החיה |  |
| 6:21.129 | הנאכלת |  |
| 6:21.130 | ובין |  |
| 6:21.131 | החיה |  |
| 6:21.132 | אשר |  |
| 6:21.133 | לא |  |
| 6:21.134 | תאכל | , |
| 6:21.135 | ושאר |  |
| 6:21.136 | איסורי |  |
| 6:21.137 | מאכלות |  |
| 6:21.138 | הכל |  |
| 6:21.139 | נקרא |  |
| 6:21.140 | אצלו |  |
| 6:21.141 | עוסק |  |
| 6:21.142 | שלא |  |
| 6:21.143 | לשמה |  |
| 6:21.144 | כמו |  |
| 6:21.145 | שאמר |  |
| 6:21.146 | שם | - |
| 6:21.147 | לא |  |
| 6:21.148 | משתדלין |  |
| 6:21.149 | אלא |  |
| 6:21.150 | לגרמייהו |  |
| 6:21.151 | לדבתא |  |
| 6:21.152 | בשרא |  |
| 6:21.153 | וכו | '. |
| 6:21.154 | ולא |  |
| 6:21.155 | נקרא |  |
| 6:21.156 | עוסק |  |
| 6:21.157 | לשמה |  |
| 6:21.158 | אלא |  |
| 6:21.159 | העוסק |  |
| 6:21.160 | בתורה |  |
| 6:21.161 | כדי |  |
| 6:21.162 | לחבר |  |
| 6:21.163 | האלהות |  |
| 6:21.164 | הזכרים |  |
| 6:21.165 | בנקבות | , |
| 6:21.166 | אבא |  |
| 6:21.167 | עם |  |
| 6:21.168 | אימא | , |
| 6:21.169 | וזעיר |  |
| 6:21.170 | עם |  |
| 6:21.171 | נוקביה | . |
| 6:21.172 | לזוגם |  |
| 6:21.173 | יחד |  |
| 6:21.174 | והיו |  |
| 6:21.175 | לבשר |  |
| 6:21.176 | אחד | . |
| 6:21.177 | וכן |  |
| 6:21.178 | בקריאת |  |
| 6:21.179 | שמע |  |
| 6:21.180 | וציצית |  |
| 6:21.181 | ותפילין |  |
| 6:21.182 | ושאר |  |
| 6:21.183 | כל |  |
| 6:21.184 | המצות | . |
| 6:21.185 | ולכן |  |
| 6:21.186 | תקנו |  |
| 6:21.187 | המקובלים |  |
| 6:21.188 | לומר |  |
| 6:21.189 | קודם |  |
| 6:21.190 | כל |  |
| 6:21.191 | מצוה |  |
| 6:21.192 | שעושים |  |
| 6:21.193 | לשם |  |
| 6:21.194 | יחוד |  |
| 6:21.195 | קב"ה |  |
| 6:21.196 | ושכינתיה |  |
| 6:21.197 | וכו | ', |
| 6:21.198 | כי |  |
| 6:21.199 | תכלית |  |
| 6:21.200 | כוונת |  |
| 6:21.201 | הפילסוף |  |
| 6:21.202 | להכניס |  |
| 6:21.203 | בישראל |  |
| 6:21.204 | אמונת |  |
| 6:21.205 | רבוי |  |
| 6:21.206 | באלוהות |  |
| 6:21.207 | זכרים |  |
| 6:21.208 | ונקבות |  |
| 6:21.209 | ולעשות |  |
| 6:21.210 | חבורים |  |
| 6:21.211 | שלו |  |
| 6:21.212 | שהם |  |
| 6:21.213 | הזהר |  |
| 6:21.214 | והתיקונים |  |
| 6:21.215 | עיקר |  |
| 6:21.216 | לעסוק |  |
| 6:21.217 | בהם |  |
| 6:21.218 | חלף |  |
| 6:21.219 | המשנה |  |
| 6:21.220 | והתלמוד | . |
| 6:21.221 | ולשכח |  |
| 6:21.222 | מישראל |  |
| 6:21.223 | עסק |  |
| 6:21.224 | המשנה |  |
| 6:21.225 | והתלמוד |  |
| 6:21.226 | שכל |  |
| 6:21.227 | כללותיה |  |
| 6:21.228 | ופרטותיה |  |
| 6:21.229 | של |  |
| 6:21.230 | תורה |  |
| 6:21.231 | כלולים |  |
| 6:21.232 | בהם | . |
| 6:21.233 | ובכמה |  |
| 6:21.234 | מקומות |  |
| 6:21.235 | הוא |  |
| 6:21.236 | משבח |  |
| 6:21.237 | את |  |
| 6:21.238 | חבורו |  |
| 6:21.239 | זה |  |
| 6:21.240 | ומהם |  |
| 6:21.241 | בתיקוני |  |
| 6:21.242 | זוהר |  |
| 6:21.243 | חדש |  |
| 6:21.244 | דף |  |
| 6:21.245 | י"ו |  |
| 6:21.246 | ע"ב | . |
| 6:21.247 | ובנפת |  |
| 6:21.248 | אשר |  |
| 6:21.249 | יטיף |  |
| 6:21.250 | הפילסוף |  |
| 6:21.251 | הנז | ' |
| 6:21.252 | בשפתי |  |
| 6:21.253 | אמונה |  |
| 6:21.254 | שהיא |  |
| 6:21.255 | זרה |  |
| 6:21.256 | לאמונת |  |
| 6:21.257 | תורתינו |  |
| 6:21.258 | הקדושה |  |
| 6:21.259 | משך |  |
| 6:21.260 | לב |  |
| 6:21.261 | רבים |  |
| 6:21.262 | בחלקלקותיו | , |
| 6:21.263 | וישיאם |  |
| 6:21.264 | לעבוד |  |
| 6:21.265 | אלהים |  |
| 6:21.266 | אחרים |  |
| 6:21.267 | המה |  |
| 6:21.268 | הצורות |  |
| 6:21.269 | והפרצופים |  |
| 6:21.270 | אשר |  |
| 6:21.271 | בדה |  |
| 6:21.272 | מלבו |  |
| 6:21.273 | להטעותנו |  |
| 6:21.274 | בפירוש |  |
| 6:21.275 | עשר |  |
| 6:21.276 | ספירות |  |
| 6:21.277 | הנז | ' |
| 6:21.278 | בספר |  |
| 6:21.279 | יצירה |  |
| 6:21.280 | לחשבם |  |
| 6:21.281 | אלהות |  |
| 6:21.282 | ולעבוד |  |
| 6:21.283 | אותם |  |
| 6:21.284 | בעבור |  |
| 6:21.285 | שנשמותיהם |  |
| 6:21.286 | חצובים |  |
| 6:21.287 | מהאלהים |  |
| 6:21.288 | העליון |  |
| 6:21.289 | לפי |  |
| 6:21.290 | דעתו | ) |
| 6:21.291 | אשר |  |
| 6:21.292 | חלק |  |
| 6:21.293 | להם |  |
| 6:21.294 | מעצמותו |  |
| 6:21.295 | ומהותו |  |
| 6:21.296 | הם |  |
| 6:21.297 | ועיגוליהם |  |
| 6:21.298 | וגופם |  |
| 6:21.299 | ונשמתם |  |
| 6:21.300 | ולבושם | , |
| 6:21.301 | כולם |  |
| 6:21.302 | חוצבו |  |
| 6:21.303 | ממקור |  |
| 6:21.304 | אחד |  |
| 6:21.305 | הוא |  |
| 6:21.306 | הא"ם | ) |
| 6:21.307 | אשר |  |
| 6:21.308 | לפי |  |
| 6:21.309 | דעתם |  |
| 6:21.310 | אין |  |
| 6:21.311 | לו |  |
| 6:21.312 | שם |  |
| 6:21.313 | ונקודה |  |
| 6:21.314 | ולית |  |
| 6:21.315 | מחשבה |  |
| 6:21.316 | תפיסה |  |
| 6:21.317 | ביה | , |
| 6:21.318 | ולא |  |
| 6:21.319 | שייך |  |
| 6:21.320 | שום |  |
| 6:21.321 | עבודה |  |
| 6:21.322 | וקריאה |  |
| 6:21.323 | אליו | , |
| 6:21.324 | ולא |  |
| 6:21.325 | יענה |  |
| 6:21.326 | את |  |
| 6:21.327 | הקורא |  |
| 6:21.328 | אליו |  |
| 6:21.329 | שכבר |  |
| 6:21.330 | מינה |  |
| 6:21.331 | והפקיד |  |
| 6:21.332 | את |  |
| 6:21.333 | הפרצופים |  |
| 6:21.334 | על |  |
| 6:21.335 | כל |  |
| 6:21.336 | מלאכת |  |
| 6:21.337 | הבריאה | . |
| 6:21.338 | ועשה |  |
| 6:21.339 | את |  |
| 6:21.340 | אבא |  |
| 6:21.341 | בורא | , |
| 6:21.342 | והוא |  |
| 6:21.343 | אומר |  |
| 6:21.344 | ומצוה |  |
| 6:21.345 | לאימא |  |
| 6:21.346 | יהי |  |
| 6:21.347 | כך |  |
| 6:21.348 | וכך | , |
| 6:21.349 | והיא |  |
| 6:21.350 | עושה |  |
| 6:21.351 | כדבריו | , |
| 6:21.352 | כאומן |  |
| 6:21.353 | שעושה |  |
| 6:21.354 | מלאכת |  |
| 6:21.355 | בעל |  |
| 6:21.356 | הבית |  |
| 6:21.357 | באמונה | . |
| 6:21.358 | ואבא |  |
| 6:21.359 | ואמא |  |
| 6:21.360 | גם |  |
| 6:21.361 | הם |  |
| 6:21.362 | צוו |  |
| 6:21.363 | והפקידו |  |
| 6:21.364 | את |  |
| 6:21.365 | זעיר |  |
| 6:21.366 | אנפין |  |
| 6:21.367 | ונוקבתיה |  |
| 6:21.368 | להשגיח |  |
| 6:21.369 | ולרון |  |
| 6:21.370 | על |  |
| 6:21.371 | מעשה |  |
| 6:21.372 | התחתונים | . |
| 6:21.373 | מאן |  |
| 6:21.374 | דזכי |  |
| 6:21.375 | לדינא |  |
| 6:21.376 | לדינא | , |
| 6:21.377 | ומאן |  |
| 6:21.378 | דזכי |  |
| 6:21.379 | לרחמי |  |
| 6:21.380 | לרחמי | , |
| 6:21.381 | הכל |  |
| 6:21.382 | נעשה |  |
| 6:21.383 | כחפצו |  |
| 6:21.384 | ורצונו |  |
| 6:21.385 | של |  |
| 6:21.386 | זעיר |  |
| 6:21.387 | אנפין | . |
| 6:21.388 | וכל |  |
| 6:21.389 | דברי |  |
| 6:21.390 | המחבר |  |
| 6:21.391 | הזה |  |
| 6:21.392 | בנוים |  |
| 6:21.393 | על |  |
| 6:21.394 | הפילוסופיא |  |
| 6:21.395 | הקדומה |  |
| 6:21.396 | אשר |  |
| 6:21.397 | כבר |  |
| 6:21.398 | נמר |  |
| 6:21.399 | ריחה |  |
| 6:21.400 | וטעמה |  |
| 6:21.401 | מעת |  |
| 6:21.402 | אשר |  |
| 6:21.403 | נגלה |  |
| 6:21.404 | כבודו |  |
| 6:21.405 | ית | ' |
| 6:21.406 | על |  |
| 6:21.407 | הר |  |
| 6:21.408 | סיני |  |
| 6:21.409 | ואמר |  |
| 6:21.410 | אנכי |  |
| 6:21.411 | ה | ' |
| 6:21.412 | אלהיך | , |
| 6:21.413 | לא |  |
| 6:21.414 | יהיה |  |
| 6:21.415 | לך |  |
| 6:21.416 | אלהים |  |
| 6:21.417 | אחרים | . |
| 6:21.418 | רק |  |
| 6:21.419 | שאחרי |  |
| 6:21.420 | מות |  |
| 6:21.421 | יהושע |  |
| 6:21.422 | והזקנים |  |
| 6:21.423 | חזרו |  |
| 6:21.424 | להדיעות |  |
| 6:21.425 | האלו |  |
| 6:21.426 | הנפסדות |  |
| 6:21.427 | ויעבדו |  |
| 6:21.428 | את |  |
| 6:21.429 | הבעלים |  |
| 6:21.430 | והעשתרות |  |
| 6:21.431 | האלו |  |
| 6:21.432 | ויעשו |  |
| 6:21.433 | צלמיהם |  |
| 6:21.434 | כדי |  |
| 6:21.435 | להוריד |  |
| 6:21.436 | השפע |  |
| 6:21.437 | מהן |  |
| 6:21.438 | על |  |
| 6:21.439 | אותן |  |
| 6:21.440 | הצורות |  |
| 6:21.441 | והצלמים |  |
| 6:21.442 | לפי |  |
| 6:21.443 | דעתם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

