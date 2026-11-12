# Translation pack: milhamot-hashem perek-8-12 (Milchamot ha-Shem 8:14–8:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-12`.

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
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 8:14

Hebrew (Wikisource): קט) וכיון שאיוו לאלוהות הרבה קיימו וקיבלו שהמציא האלוה העליון הנקרא אין סיבות ועילות רבים המשלתשלים ונשפעים זה מזה, וכל אחד מהם נקרא בשם הוי'ה או אדנות או אלהים והקב"ה וכל השמות, ודרשו את הכתוב מה שמו ומה שם בנו כי תדע, מה שמו ה' צבאות שמו, (פי' אבא) ומה שם בנו, זעיר אנפין שמו וכו' כדלעיל סי' כ"ו. וכל עבודתינו ותפלותינו לפי דבריהם אל האלהות האלו הנמשכים מהאלוה העליון שקראוהו אין, כי האין לא יענה את הקורא אליו, כי אם האלהים התחתונים המה יענו, ובלבד שניחד תפלתינו אל האלוה קצר האפים (ז"א) ונשתף במחשבתינו גם פרצופים שלמעלה ממנו. יוחשב לנו עבודת קצר האפים הנברא הזה כאלו אנחנו עובדים וקורים אל האלוה העליון, כי נשמתו היא חלק אלוה ממעל.

| id | word | punctuation after |
|---|---|---|
| 8:14.1 | קט | ) |
| 8:14.2 | וכיון |  |
| 8:14.3 | שאיוו |  |
| 8:14.4 | לאלוהות |  |
| 8:14.5 | הרבה |  |
| 8:14.6 | קיימו |  |
| 8:14.7 | וקיבלו |  |
| 8:14.8 | שהמציא |  |
| 8:14.9 | האלוה |  |
| 8:14.10 | העליון |  |
| 8:14.11 | הנקרא |  |
| 8:14.12 | אין |  |
| 8:14.13 | סיבות |  |
| 8:14.14 | ועילות |  |
| 8:14.15 | רבים |  |
| 8:14.16 | המשלתשלים |  |
| 8:14.17 | ונשפעים |  |
| 8:14.18 | זה |  |
| 8:14.19 | מזה | , |
| 8:14.20 | וכל |  |
| 8:14.21 | אחד |  |
| 8:14.22 | מהם |  |
| 8:14.23 | נקרא |  |
| 8:14.24 | בשם |  |
| 8:14.25 | הוי'ה |  |
| 8:14.26 | או |  |
| 8:14.27 | אדנות |  |
| 8:14.28 | או |  |
| 8:14.29 | אלהים |  |
| 8:14.30 | והקב"ה |  |
| 8:14.31 | וכל |  |
| 8:14.32 | השמות | , |
| 8:14.33 | ודרשו |  |
| 8:14.34 | את |  |
| 8:14.35 | הכתוב |  |
| 8:14.36 | מה |  |
| 8:14.37 | שמו |  |
| 8:14.38 | ומה |  |
| 8:14.39 | שם |  |
| 8:14.40 | בנו |  |
| 8:14.41 | כי |  |
| 8:14.42 | תדע | , |
| 8:14.43 | מה |  |
| 8:14.44 | שמו |  |
| 8:14.45 | ה | ' |
| 8:14.46 | צבאות |  |
| 8:14.47 | שמו | , |
| 8:14.48 | פי | ' |
| 8:14.49 | אבא | ) |
| 8:14.50 | ומה |  |
| 8:14.51 | שם |  |
| 8:14.52 | בנו | , |
| 8:14.53 | זעיר |  |
| 8:14.54 | אנפין |  |
| 8:14.55 | שמו |  |
| 8:14.56 | וכו | ' |
| 8:14.57 | כדלעיל |  |
| 8:14.58 | סי | ' |
| 8:14.59 | כ"ו | . |
| 8:14.60 | וכל |  |
| 8:14.61 | עבודתינו |  |
| 8:14.62 | ותפלותינו |  |
| 8:14.63 | לפי |  |
| 8:14.64 | דבריהם |  |
| 8:14.65 | אל |  |
| 8:14.66 | האלהות |  |
| 8:14.67 | האלו |  |
| 8:14.68 | הנמשכים |  |
| 8:14.69 | מהאלוה |  |
| 8:14.70 | העליון |  |
| 8:14.71 | שקראוהו |  |
| 8:14.72 | אין | , |
| 8:14.73 | כי |  |
| 8:14.74 | האין |  |
| 8:14.75 | לא |  |
| 8:14.76 | יענה |  |
| 8:14.77 | את |  |
| 8:14.78 | הקורא |  |
| 8:14.79 | אליו | , |
| 8:14.80 | כי |  |
| 8:14.81 | אם |  |
| 8:14.82 | האלהים |  |
| 8:14.83 | התחתונים |  |
| 8:14.84 | המה |  |
| 8:14.85 | יענו | , |
| 8:14.86 | ובלבד |  |
| 8:14.87 | שניחד |  |
| 8:14.88 | תפלתינו |  |
| 8:14.89 | אל |  |
| 8:14.90 | האלוה |  |
| 8:14.91 | קצר |  |
| 8:14.92 | האפים |  |
| 8:14.93 | ז"א | ) |
| 8:14.94 | ונשתף |  |
| 8:14.95 | במחשבתינו |  |
| 8:14.96 | גם |  |
| 8:14.97 | פרצופים |  |
| 8:14.98 | שלמעלה |  |
| 8:14.99 | ממנו | . |
| 8:14.100 | יוחשב |  |
| 8:14.101 | לנו |  |
| 8:14.102 | עבודת |  |
| 8:14.103 | קצר |  |
| 8:14.104 | האפים |  |
| 8:14.105 | הנברא |  |
| 8:14.106 | הזה |  |
| 8:14.107 | כאלו |  |
| 8:14.108 | אנחנו |  |
| 8:14.109 | עובדים |  |
| 8:14.110 | וקורים |  |
| 8:14.111 | אל |  |
| 8:14.112 | האלוה |  |
| 8:14.113 | העליון | , |
| 8:14.114 | כי |  |
| 8:14.115 | נשמתו |  |
| 8:14.116 | היא |  |
| 8:14.117 | חלק |  |
| 8:14.118 | אלוה |  |
| 8:14.119 | ממעל | . |


## Milchamot ha-Shem 8:15

Hebrew (Wikisource): ולכן הוכיח ירמיה הנביא את ישראל ואמר כחשו בה' שאמר למשה כה תאמר אל בני ישראל יהוה אלהי אבותיכם אלהי אברהם, אלהי יצחק ואלהי יעקב שלחני אליכם זה שמי לעולם וזה זכרי לדר דר. ויאמרו לא הוא המטיב ומשגיח עלינו, ולא אותו נעבוד אלא קצר האפים הוא המשגיח עלינו, ואליו נקרא ואותו נעבוד, והוא המלך הקדוש המהולל בכל התשבחות, כי לו ניתן הכח והממשלה מאביו ואמו, ושהאלהים העליון הוא המאחד ומדבק ומקשר את כל האלהות התחתונים ופרצופיהם להיות כולם נחשבים אחד, אע"פ שהם באמת חלוקים. ובקראינו אל הזעיר אנפין כאלו קראנו אל האלהים העליון, ולכך כתיב לוא מלא וא"ו ואלף קרי ביה לו, וקרי ביה לא, לא אותו נעבוד, כי אם את קצר האפים, וקרי ביה לו בוא"ו כי בעבודתינו לז"א כאלו לו אנחנו עובדים, כי זהו רצונו של מלך העליון שנעבוד את המלך התחתון (זעיר אנפין) שמינהו תחתיו והשליטו על כל הברואים התחתונים וגם ליחד אשתו עמו. ובאמרינו שמע ישראל ה' אלהינו ה' אחד אנו מיחדים לפי דבריהם גם את אבא ואימא עמו כי ה' הראשון הוא אבא, אלהינו אימא, ה' דא זעיר אנפין, אחד דא מלכות, גי א"ת ט' ספירות, ודל"ת כולל הכל להיות אחד, ובכוונה זאת אנו מאחדים ואוגדים את כולם להיות אחד, מלכא עלאה אבא ואימא, ומלכא תתהה זעיר ונוקביה, לקשרא דא בדא ולא ישתכח פירודא אלא כלא חד כדלעיל סי' מ"ב. הה! אוי! ואבוי! על רוע מצב אמונת העם אשר בחר לו יה ישראל לסגולתו. ונתן לנו תורת אמת. היא חותמו של הקב"ה אשר קבל משה רבינו עליו השלום בסיני בכתב ובעל פה. ועמד עליה מ' יום וארבעים לילה ללמוד דקדוקיה וטעמיה וכללותיה, ומסרה ליהושע וכו', ואנשי כנסת הגדולה תפסו ליצרא דע"ז ושדויה בדודא דאברא וחפיוה לפומיה באברא. הן עתה הוסר כסוי העופרת מעל פי האיפה על יד המלך העכו"ם במסרו את בפר הזהר לחכמי טוליטולה, יוצא השטן וישב וירקד בנינו, ויהי לריק כל יגיעם של אנשי כנסת הגדולה אשר טרחו ויגעו. וצעקו לפני השי"ת על יצרא דע"ז. ומן השמיים סייעו אותם ונמסר בידם. ורוב חכמתם ורוח בינתם ביטלו יצרא דע"ז על ידי עסק התורה וביאור המבינים לעם בה על פי האמת והצדק כמו שהוא ערוך לפנינו בב' התלמודים ובמדרשים האמתיים. שג' ויקראו בו לפני הרחוב מאור הבקר עד מחצית היום נגד האנשים והנשים המבינים. ונאמר הלוים מבינים את העם לתורה ויקראו בס' בתורת האלהים מפורש ושום שכל ויבינו במקרא. ועל ידי תקון נוסח התפלות והברכות אשר תקנו לכל ישראל. ועל ידי השתדלותם להעמיד תלמידים הרבה כדאיתא בפרקא קמא דאבות! על ידי כל אלה עלתה בידם לעקור יצרא דע"ז, והחזירו עטרה לישנה, ונמחה שם האלילים וכל האמונות הכוזבות המביאות לידי מינות מישראל, כדגרסינן ביומא פרק בא לו, ויקם על מעלה הלוים ישוע ובני וקדמיאל שבניה בני שרביה בני כנני ויזעקו בקול גדול אל ה' אלהיהם, מאי אמור אמר רב יהודה אמר רב ואי תימא ר' יוחנן בייא בייא היינו האיי דאחרביה לבי מקדשא וקלייה להיכלא וקטלניהו לצדיקיא ואגלינהו לישראל מארעהון ועדיין מרקיד בינא, כלום יהבתיה ניהלן אלא לקבולי ביה אגרא. לא איהו בעינן ולא אגריה, בעינן, נפל להו פתקא מרקיעא דכתיב ביה אמת. אמר רבי חנינה שמע מינה חותמו של הקב"ה אמת, יתיבו בתעניתא תלתה יומין ותלתא לילי מסריה ניהלינו נפק אתא כי גוריא דנורא מבית קדש הקדשים. אמרו להו נביא לישראל היינו הוא יצרא דע"ז שנאמר ויאמר זאת הרשעה וגומר בהדי דתפסי ליה, אשתמיט ביניתא ממזיה ורמא קלא, אזל קליה ארבע מאה פרסי אמר היכי נעביד ליה, דלמא מרחמי עליה מן שמיא. אמר להו נביא שדויה בדורא דאברא והפויה לפומא באברא דשאיב קלא, שנאמר זאת הרשעה וישלך אותה אל תוך האיפה וישלך את אבן העופרת אל פיה. אמרי הואיל ועת רצון היא, נבעי רחמי על יצרא דעבירה. בעו רחמי ואמסר בידיהו, אמר להו נביא הזו ודאי קטלתון לההוא כליא עלמא, חבשוה תלתא יומי. איבעי ביעתא בת יומא לחולה בכל ארץ ישראל ולא אשתכח. אמרי היכי נעביד, ניקטליה איחרב כולי עלמא, נבעי רחמי לפלגא, פלגא מרקיעא לא יהבי. כחלינהו לעיניה ושבקוה, ואהני דלא תגרי ביה באיניש בקרובתא ע"כ. ועיין עץ יוסף, ובענף שכתב שהמאמר תמוה הוא מאד, ויצרא דע"ז בבית קדש הקדשים מאי עבידתיה? וכו'.

| id | word | punctuation after |
|---|---|---|
| 8:15.1 | ולכן |  |
| 8:15.2 | הוכיח |  |
| 8:15.3 | ירמיה |  |
| 8:15.4 | הנביא |  |
| 8:15.5 | את |  |
| 8:15.6 | ישראל |  |
| 8:15.7 | ואמר |  |
| 8:15.8 | כחשו |  |
| 8:15.9 | בה | ' |
| 8:15.10 | שאמר |  |
| 8:15.11 | למשה |  |
| 8:15.12 | כה |  |
| 8:15.13 | תאמר |  |
| 8:15.14 | אל |  |
| 8:15.15 | בני |  |
| 8:15.16 | ישראל |  |
| 8:15.17 | יהוה |  |
| 8:15.18 | אלהי |  |
| 8:15.19 | אבותיכם |  |
| 8:15.20 | אלהי |  |
| 8:15.21 | אברהם | , |
| 8:15.22 | אלהי |  |
| 8:15.23 | יצחק |  |
| 8:15.24 | ואלהי |  |
| 8:15.25 | יעקב |  |
| 8:15.26 | שלחני |  |
| 8:15.27 | אליכם |  |
| 8:15.28 | זה |  |
| 8:15.29 | שמי |  |
| 8:15.30 | לעולם |  |
| 8:15.31 | וזה |  |
| 8:15.32 | זכרי |  |
| 8:15.33 | לדר |  |
| 8:15.34 | דר | . |
| 8:15.35 | ויאמרו |  |
| 8:15.36 | לא |  |
| 8:15.37 | הוא |  |
| 8:15.38 | המטיב |  |
| 8:15.39 | ומשגיח |  |
| 8:15.40 | עלינו | , |
| 8:15.41 | ולא |  |
| 8:15.42 | אותו |  |
| 8:15.43 | נעבוד |  |
| 8:15.44 | אלא |  |
| 8:15.45 | קצר |  |
| 8:15.46 | האפים |  |
| 8:15.47 | הוא |  |
| 8:15.48 | המשגיח |  |
| 8:15.49 | עלינו | , |
| 8:15.50 | ואליו |  |
| 8:15.51 | נקרא |  |
| 8:15.52 | ואותו |  |
| 8:15.53 | נעבוד | , |
| 8:15.54 | והוא |  |
| 8:15.55 | המלך |  |
| 8:15.56 | הקדוש |  |
| 8:15.57 | המהולל |  |
| 8:15.58 | בכל |  |
| 8:15.59 | התשבחות | , |
| 8:15.60 | כי |  |
| 8:15.61 | לו |  |
| 8:15.62 | ניתן |  |
| 8:15.63 | הכח |  |
| 8:15.64 | והממשלה |  |
| 8:15.65 | מאביו |  |
| 8:15.66 | ואמו | , |
| 8:15.67 | ושהאלהים |  |
| 8:15.68 | העליון |  |
| 8:15.69 | הוא |  |
| 8:15.70 | המאחד |  |
| 8:15.71 | ומדבק |  |
| 8:15.72 | ומקשר |  |
| 8:15.73 | את |  |
| 8:15.74 | כל |  |
| 8:15.75 | האלהות |  |
| 8:15.76 | התחתונים |  |
| 8:15.77 | ופרצופיהם |  |
| 8:15.78 | להיות |  |
| 8:15.79 | כולם |  |
| 8:15.80 | נחשבים |  |
| 8:15.81 | אחד | , |
| 8:15.82 | אע"פ |  |
| 8:15.83 | שהם |  |
| 8:15.84 | באמת |  |
| 8:15.85 | חלוקים | . |
| 8:15.86 | ובקראינו |  |
| 8:15.87 | אל |  |
| 8:15.88 | הזעיר |  |
| 8:15.89 | אנפין |  |
| 8:15.90 | כאלו |  |
| 8:15.91 | קראנו |  |
| 8:15.92 | אל |  |
| 8:15.93 | האלהים |  |
| 8:15.94 | העליון | , |
| 8:15.95 | ולכך |  |
| 8:15.96 | כתיב |  |
| 8:15.97 | לוא |  |
| 8:15.98 | מלא |  |
| 8:15.99 | וא"ו |  |
| 8:15.100 | ואלף |  |
| 8:15.101 | קרי |  |
| 8:15.102 | ביה |  |
| 8:15.103 | לו | , |
| 8:15.104 | וקרי |  |
| 8:15.105 | ביה |  |
| 8:15.106 | לא | , |
| 8:15.107 | לא |  |
| 8:15.108 | אותו |  |
| 8:15.109 | נעבוד | , |
| 8:15.110 | כי |  |
| 8:15.111 | אם |  |
| 8:15.112 | את |  |
| 8:15.113 | קצר |  |
| 8:15.114 | האפים | , |
| 8:15.115 | וקרי |  |
| 8:15.116 | ביה |  |
| 8:15.117 | לו |  |
| 8:15.118 | בוא"ו |  |
| 8:15.119 | כי |  |
| 8:15.120 | בעבודתינו |  |
| 8:15.121 | לז"א |  |
| 8:15.122 | כאלו |  |
| 8:15.123 | לו |  |
| 8:15.124 | אנחנו |  |
| 8:15.125 | עובדים | , |
| 8:15.126 | כי |  |
| 8:15.127 | זהו |  |
| 8:15.128 | רצונו |  |
| 8:15.129 | של |  |
| 8:15.130 | מלך |  |
| 8:15.131 | העליון |  |
| 8:15.132 | שנעבוד |  |
| 8:15.133 | את |  |
| 8:15.134 | המלך |  |
| 8:15.135 | התחתון |  |
| 8:15.136 | זעיר |  |
| 8:15.137 | אנפין | ) |
| 8:15.138 | שמינהו |  |
| 8:15.139 | תחתיו |  |
| 8:15.140 | והשליטו |  |
| 8:15.141 | על |  |
| 8:15.142 | כל |  |
| 8:15.143 | הברואים |  |
| 8:15.144 | התחתונים |  |
| 8:15.145 | וגם |  |
| 8:15.146 | ליחד |  |
| 8:15.147 | אשתו |  |
| 8:15.148 | עמו | . |
| 8:15.149 | ובאמרינו |  |
| 8:15.150 | שמע |  |
| 8:15.151 | ישראל |  |
| 8:15.152 | ה | ' |
| 8:15.153 | אלהינו |  |
| 8:15.154 | ה | ' |
| 8:15.155 | אחד |  |
| 8:15.156 | אנו |  |
| 8:15.157 | מיחדים |  |
| 8:15.158 | לפי |  |
| 8:15.159 | דבריהם |  |
| 8:15.160 | גם |  |
| 8:15.161 | את |  |
| 8:15.162 | אבא |  |
| 8:15.163 | ואימא |  |
| 8:15.164 | עמו |  |
| 8:15.165 | כי |  |
| 8:15.166 | ה | ' |
| 8:15.167 | הראשון |  |
| 8:15.168 | הוא |  |
| 8:15.169 | אבא | , |
| 8:15.170 | אלהינו |  |
| 8:15.171 | אימא | , |
| 8:15.172 | ה | ' |
| 8:15.173 | דא |  |
| 8:15.174 | זעיר |  |
| 8:15.175 | אנפין | , |
| 8:15.176 | אחד |  |
| 8:15.177 | דא |  |
| 8:15.178 | מלכות | , |
| 8:15.179 | גי |  |
| 8:15.180 | א"ת |  |
| 8:15.181 | ט | ' |
| 8:15.182 | ספירות | , |
| 8:15.183 | ודל"ת |  |
| 8:15.184 | כולל |  |
| 8:15.185 | הכל |  |
| 8:15.186 | להיות |  |
| 8:15.187 | אחד | , |
| 8:15.188 | ובכוונה |  |
| 8:15.189 | זאת |  |
| 8:15.190 | אנו |  |
| 8:15.191 | מאחדים |  |
| 8:15.192 | ואוגדים |  |
| 8:15.193 | את |  |
| 8:15.194 | כולם |  |
| 8:15.195 | להיות |  |
| 8:15.196 | אחד | , |
| 8:15.197 | מלכא |  |
| 8:15.198 | עלאה |  |
| 8:15.199 | אבא |  |
| 8:15.200 | ואימא | , |
| 8:15.201 | ומלכא |  |
| 8:15.202 | תתהה |  |
| 8:15.203 | זעיר |  |
| 8:15.204 | ונוקביה | , |
| 8:15.205 | לקשרא |  |
| 8:15.206 | דא |  |
| 8:15.207 | בדא |  |
| 8:15.208 | ולא |  |
| 8:15.209 | ישתכח |  |
| 8:15.210 | פירודא |  |
| 8:15.211 | אלא |  |
| 8:15.212 | כלא |  |
| 8:15.213 | חד |  |
| 8:15.214 | כדלעיל |  |
| 8:15.215 | סי | ' |
| 8:15.216 | מ"ב | . |
| 8:15.217 | הה | ! |
| 8:15.218 | אוי | ! |
| 8:15.219 | ואבוי | ! |
| 8:15.220 | על |  |
| 8:15.221 | רוע |  |
| 8:15.222 | מצב |  |
| 8:15.223 | אמונת |  |
| 8:15.224 | העם |  |
| 8:15.225 | אשר |  |
| 8:15.226 | בחר |  |
| 8:15.227 | לו |  |
| 8:15.228 | יה |  |
| 8:15.229 | ישראל |  |
| 8:15.230 | לסגולתו | . |
| 8:15.231 | ונתן |  |
| 8:15.232 | לנו |  |
| 8:15.233 | תורת |  |
| 8:15.234 | אמת | . |
| 8:15.235 | היא |  |
| 8:15.236 | חותמו |  |
| 8:15.237 | של |  |
| 8:15.238 | הקב"ה |  |
| 8:15.239 | אשר |  |
| 8:15.240 | קבל |  |
| 8:15.241 | משה |  |
| 8:15.242 | רבינו |  |
| 8:15.243 | עליו |  |
| 8:15.244 | השלום |  |
| 8:15.245 | בסיני |  |
| 8:15.246 | בכתב |  |
| 8:15.247 | ובעל |  |
| 8:15.248 | פה | . |
| 8:15.249 | ועמד |  |
| 8:15.250 | עליה |  |
| 8:15.251 | מ | ' |
| 8:15.252 | יום |  |
| 8:15.253 | וארבעים |  |
| 8:15.254 | לילה |  |
| 8:15.255 | ללמוד |  |
| 8:15.256 | דקדוקיה |  |
| 8:15.257 | וטעמיה |  |
| 8:15.258 | וכללותיה | , |
| 8:15.259 | ומסרה |  |
| 8:15.260 | ליהושע |  |
| 8:15.261 | וכו | ', |
| 8:15.262 | ואנשי |  |
| 8:15.263 | כנסת |  |
| 8:15.264 | הגדולה |  |
| 8:15.265 | תפסו |  |
| 8:15.266 | ליצרא |  |
| 8:15.267 | דע"ז |  |
| 8:15.268 | ושדויה |  |
| 8:15.269 | בדודא |  |
| 8:15.270 | דאברא |  |
| 8:15.271 | וחפיוה |  |
| 8:15.272 | לפומיה |  |
| 8:15.273 | באברא | . |
| 8:15.274 | הן |  |
| 8:15.275 | עתה |  |
| 8:15.276 | הוסר |  |
| 8:15.277 | כסוי |  |
| 8:15.278 | העופרת |  |
| 8:15.279 | מעל |  |
| 8:15.280 | פי |  |
| 8:15.281 | האיפה |  |
| 8:15.282 | על |  |
| 8:15.283 | יד |  |
| 8:15.284 | המלך |  |
| 8:15.285 | העכו"ם |  |
| 8:15.286 | במסרו |  |
| 8:15.287 | את |  |
| 8:15.288 | בפר |  |
| 8:15.289 | הזהר |  |
| 8:15.290 | לחכמי |  |
| 8:15.291 | טוליטולה | , |
| 8:15.292 | יוצא |  |
| 8:15.293 | השטן |  |
| 8:15.294 | וישב |  |
| 8:15.295 | וירקד |  |
| 8:15.296 | בנינו | , |
| 8:15.297 | ויהי |  |
| 8:15.298 | לריק |  |
| 8:15.299 | כל |  |
| 8:15.300 | יגיעם |  |
| 8:15.301 | של |  |
| 8:15.302 | אנשי |  |
| 8:15.303 | כנסת |  |
| 8:15.304 | הגדולה |  |
| 8:15.305 | אשר |  |
| 8:15.306 | טרחו |  |
| 8:15.307 | ויגעו | . |
| 8:15.308 | וצעקו |  |
| 8:15.309 | לפני |  |
| 8:15.310 | השי"ת |  |
| 8:15.311 | על |  |
| 8:15.312 | יצרא |  |
| 8:15.313 | דע"ז | . |
| 8:15.314 | ומן |  |
| 8:15.315 | השמיים |  |
| 8:15.316 | סייעו |  |
| 8:15.317 | אותם |  |
| 8:15.318 | ונמסר |  |
| 8:15.319 | בידם | . |
| 8:15.320 | ורוב |  |
| 8:15.321 | חכמתם |  |
| 8:15.322 | ורוח |  |
| 8:15.323 | בינתם |  |
| 8:15.324 | ביטלו |  |
| 8:15.325 | יצרא |  |
| 8:15.326 | דע"ז |  |
| 8:15.327 | על |  |
| 8:15.328 | ידי |  |
| 8:15.329 | עסק |  |
| 8:15.330 | התורה |  |
| 8:15.331 | וביאור |  |
| 8:15.332 | המבינים |  |
| 8:15.333 | לעם |  |
| 8:15.334 | בה |  |
| 8:15.335 | על |  |
| 8:15.336 | פי |  |
| 8:15.337 | האמת |  |
| 8:15.338 | והצדק |  |
| 8:15.339 | כמו |  |
| 8:15.340 | שהוא |  |
| 8:15.341 | ערוך |  |
| 8:15.342 | לפנינו |  |
| 8:15.343 | בב | ' |
| 8:15.344 | התלמודים |  |
| 8:15.345 | ובמדרשים |  |
| 8:15.346 | האמתיים | . |
| 8:15.347 | שג | ' |
| 8:15.348 | ויקראו |  |
| 8:15.349 | בו |  |
| 8:15.350 | לפני |  |
| 8:15.351 | הרחוב |  |
| 8:15.352 | מאור |  |
| 8:15.353 | הבקר |  |
| 8:15.354 | עד |  |
| 8:15.355 | מחצית |  |
| 8:15.356 | היום |  |
| 8:15.357 | נגד |  |
| 8:15.358 | האנשים |  |
| 8:15.359 | והנשים |  |
| 8:15.360 | המבינים | . |
| 8:15.361 | ונאמר |  |
| 8:15.362 | הלוים |  |
| 8:15.363 | מבינים |  |
| 8:15.364 | את |  |
| 8:15.365 | העם |  |
| 8:15.366 | לתורה |  |
| 8:15.367 | ויקראו |  |
| 8:15.368 | בס | ' |
| 8:15.369 | בתורת |  |
| 8:15.370 | האלהים |  |
| 8:15.371 | מפורש |  |
| 8:15.372 | ושום |  |
| 8:15.373 | שכל |  |
| 8:15.374 | ויבינו |  |
| 8:15.375 | במקרא | . |
| 8:15.376 | ועל |  |
| 8:15.377 | ידי |  |
| 8:15.378 | תקון |  |
| 8:15.379 | נוסח |  |
| 8:15.380 | התפלות |  |
| 8:15.381 | והברכות |  |
| 8:15.382 | אשר |  |
| 8:15.383 | תקנו |  |
| 8:15.384 | לכל |  |
| 8:15.385 | ישראל | . |
| 8:15.386 | ועל |  |
| 8:15.387 | ידי |  |
| 8:15.388 | השתדלותם |  |
| 8:15.389 | להעמיד |  |
| 8:15.390 | תלמידים |  |
| 8:15.391 | הרבה |  |
| 8:15.392 | כדאיתא |  |
| 8:15.393 | בפרקא |  |
| 8:15.394 | קמא |  |
| 8:15.395 | דאבות | ! |
| 8:15.396 | על |  |
| 8:15.397 | ידי |  |
| 8:15.398 | כל |  |
| 8:15.399 | אלה |  |
| 8:15.400 | עלתה |  |
| 8:15.401 | בידם |  |
| 8:15.402 | לעקור |  |
| 8:15.403 | יצרא |  |
| 8:15.404 | דע"ז | , |
| 8:15.405 | והחזירו |  |
| 8:15.406 | עטרה |  |
| 8:15.407 | לישנה | , |
| 8:15.408 | ונמחה |  |
| 8:15.409 | שם |  |
| 8:15.410 | האלילים |  |
| 8:15.411 | וכל |  |
| 8:15.412 | האמונות |  |
| 8:15.413 | הכוזבות |  |
| 8:15.414 | המביאות |  |
| 8:15.415 | לידי |  |
| 8:15.416 | מינות |  |
| 8:15.417 | מישראל | , |
| 8:15.418 | כדגרסינן |  |
| 8:15.419 | ביומא |  |
| 8:15.420 | פרק |  |
| 8:15.421 | בא |  |
| 8:15.422 | לו | , |
| 8:15.423 | ויקם |  |
| 8:15.424 | על |  |
| 8:15.425 | מעלה |  |
| 8:15.426 | הלוים |  |
| 8:15.427 | ישוע |  |
| 8:15.428 | ובני |  |
| 8:15.429 | וקדמיאל |  |
| 8:15.430 | שבניה |  |
| 8:15.431 | בני |  |
| 8:15.432 | שרביה |  |
| 8:15.433 | בני |  |
| 8:15.434 | כנני |  |
| 8:15.435 | ויזעקו |  |
| 8:15.436 | בקול |  |
| 8:15.437 | גדול |  |
| 8:15.438 | אל |  |
| 8:15.439 | ה | ' |
| 8:15.440 | אלהיהם | , |
| 8:15.441 | מאי |  |
| 8:15.442 | אמור |  |
| 8:15.443 | אמר |  |
| 8:15.444 | רב |  |
| 8:15.445 | יהודה |  |
| 8:15.446 | אמר |  |
| 8:15.447 | רב |  |
| 8:15.448 | ואי |  |
| 8:15.449 | תימא |  |
| 8:15.450 | ר | ' |
| 8:15.451 | יוחנן |  |
| 8:15.452 | בייא |  |
| 8:15.453 | בייא |  |
| 8:15.454 | היינו |  |
| 8:15.455 | האיי |  |
| 8:15.456 | דאחרביה |  |
| 8:15.457 | לבי |  |
| 8:15.458 | מקדשא |  |
| 8:15.459 | וקלייה |  |
| 8:15.460 | להיכלא |  |
| 8:15.461 | וקטלניהו |  |
| 8:15.462 | לצדיקיא |  |
| 8:15.463 | ואגלינהו |  |
| 8:15.464 | לישראל |  |
| 8:15.465 | מארעהון |  |
| 8:15.466 | ועדיין |  |
| 8:15.467 | מרקיד |  |
| 8:15.468 | בינא | , |
| 8:15.469 | כלום |  |
| 8:15.470 | יהבתיה |  |
| 8:15.471 | ניהלן |  |
| 8:15.472 | אלא |  |
| 8:15.473 | לקבולי |  |
| 8:15.474 | ביה |  |
| 8:15.475 | אגרא | . |
| 8:15.476 | לא |  |
| 8:15.477 | איהו |  |
| 8:15.478 | בעינן |  |
| 8:15.479 | ולא |  |
| 8:15.480 | אגריה | , |
| 8:15.481 | בעינן | , |
| 8:15.482 | נפל |  |
| 8:15.483 | להו |  |
| 8:15.484 | פתקא |  |
| 8:15.485 | מרקיעא |  |
| 8:15.486 | דכתיב |  |
| 8:15.487 | ביה |  |
| 8:15.488 | אמת | . |
| 8:15.489 | אמר |  |
| 8:15.490 | רבי |  |
| 8:15.491 | חנינה |  |
| 8:15.492 | שמע |  |
| 8:15.493 | מינה |  |
| 8:15.494 | חותמו |  |
| 8:15.495 | של |  |
| 8:15.496 | הקב"ה |  |
| 8:15.497 | אמת | , |
| 8:15.498 | יתיבו |  |
| 8:15.499 | בתעניתא |  |
| 8:15.500 | תלתה |  |
| 8:15.501 | יומין |  |
| 8:15.502 | ותלתא |  |
| 8:15.503 | לילי |  |
| 8:15.504 | מסריה |  |
| 8:15.505 | ניהלינו |  |
| 8:15.506 | נפק |  |
| 8:15.507 | אתא |  |
| 8:15.508 | כי |  |
| 8:15.509 | גוריא |  |
| 8:15.510 | דנורא |  |
| 8:15.511 | מבית |  |
| 8:15.512 | קדש |  |
| 8:15.513 | הקדשים | . |
| 8:15.514 | אמרו |  |
| 8:15.515 | להו |  |
| 8:15.516 | נביא |  |
| 8:15.517 | לישראל |  |
| 8:15.518 | היינו |  |
| 8:15.519 | הוא |  |
| 8:15.520 | יצרא |  |
| 8:15.521 | דע"ז |  |
| 8:15.522 | שנאמר |  |
| 8:15.523 | ויאמר |  |
| 8:15.524 | זאת |  |
| 8:15.525 | הרשעה |  |
| 8:15.526 | וגומר |  |
| 8:15.527 | בהדי |  |
| 8:15.528 | דתפסי |  |
| 8:15.529 | ליה | , |
| 8:15.530 | אשתמיט |  |
| 8:15.531 | ביניתא |  |
| 8:15.532 | ממזיה |  |
| 8:15.533 | ורמא |  |
| 8:15.534 | קלא | , |
| 8:15.535 | אזל |  |
| 8:15.536 | קליה |  |
| 8:15.537 | ארבע |  |
| 8:15.538 | מאה |  |
| 8:15.539 | פרסי |  |
| 8:15.540 | אמר |  |
| 8:15.541 | היכי |  |
| 8:15.542 | נעביד |  |
| 8:15.543 | ליה | , |
| 8:15.544 | דלמא |  |
| 8:15.545 | מרחמי |  |
| 8:15.546 | עליה |  |
| 8:15.547 | מן |  |
| 8:15.548 | שמיא | . |
| 8:15.549 | אמר |  |
| 8:15.550 | להו |  |
| 8:15.551 | נביא |  |
| 8:15.552 | שדויה |  |
| 8:15.553 | בדורא |  |
| 8:15.554 | דאברא |  |
| 8:15.555 | והפויה |  |
| 8:15.556 | לפומא |  |
| 8:15.557 | באברא |  |
| 8:15.558 | דשאיב |  |
| 8:15.559 | קלא | , |
| 8:15.560 | שנאמר |  |
| 8:15.561 | זאת |  |
| 8:15.562 | הרשעה |  |
| 8:15.563 | וישלך |  |
| 8:15.564 | אותה |  |
| 8:15.565 | אל |  |
| 8:15.566 | תוך |  |
| 8:15.567 | האיפה |  |
| 8:15.568 | וישלך |  |
| 8:15.569 | את |  |
| 8:15.570 | אבן |  |
| 8:15.571 | העופרת |  |
| 8:15.572 | אל |  |
| 8:15.573 | פיה | . |
| 8:15.574 | אמרי |  |
| 8:15.575 | הואיל |  |
| 8:15.576 | ועת |  |
| 8:15.577 | רצון |  |
| 8:15.578 | היא | , |
| 8:15.579 | נבעי |  |
| 8:15.580 | רחמי |  |
| 8:15.581 | על |  |
| 8:15.582 | יצרא |  |
| 8:15.583 | דעבירה | . |
| 8:15.584 | בעו |  |
| 8:15.585 | רחמי |  |
| 8:15.586 | ואמסר |  |
| 8:15.587 | בידיהו | , |
| 8:15.588 | אמר |  |
| 8:15.589 | להו |  |
| 8:15.590 | נביא |  |
| 8:15.591 | הזו |  |
| 8:15.592 | ודאי |  |
| 8:15.593 | קטלתון |  |
| 8:15.594 | לההוא |  |
| 8:15.595 | כליא |  |
| 8:15.596 | עלמא | , |
| 8:15.597 | חבשוה |  |
| 8:15.598 | תלתא |  |
| 8:15.599 | יומי | . |
| 8:15.600 | איבעי |  |
| 8:15.601 | ביעתא |  |
| 8:15.602 | בת |  |
| 8:15.603 | יומא |  |
| 8:15.604 | לחולה |  |
| 8:15.605 | בכל |  |
| 8:15.606 | ארץ |  |
| 8:15.607 | ישראל |  |
| 8:15.608 | ולא |  |
| 8:15.609 | אשתכח | . |
| 8:15.610 | אמרי |  |
| 8:15.611 | היכי |  |
| 8:15.612 | נעביד | , |
| 8:15.613 | ניקטליה |  |
| 8:15.614 | איחרב |  |
| 8:15.615 | כולי |  |
| 8:15.616 | עלמא | , |
| 8:15.617 | נבעי |  |
| 8:15.618 | רחמי |  |
| 8:15.619 | לפלגא | , |
| 8:15.620 | פלגא |  |
| 8:15.621 | מרקיעא |  |
| 8:15.622 | לא |  |
| 8:15.623 | יהבי | . |
| 8:15.624 | כחלינהו |  |
| 8:15.625 | לעיניה |  |
| 8:15.626 | ושבקוה | , |
| 8:15.627 | ואהני |  |
| 8:15.628 | דלא |  |
| 8:15.629 | תגרי |  |
| 8:15.630 | ביה |  |
| 8:15.631 | באיניש |  |
| 8:15.632 | בקרובתא |  |
| 8:15.633 | ע"כ | . |
| 8:15.634 | ועיין |  |
| 8:15.635 | עץ |  |
| 8:15.636 | יוסף | , |
| 8:15.637 | ובענף |  |
| 8:15.638 | שכתב |  |
| 8:15.639 | שהמאמר |  |
| 8:15.640 | תמוה |  |
| 8:15.641 | הוא |  |
| 8:15.642 | מאד | , |
| 8:15.643 | ויצרא |  |
| 8:15.644 | דע"ז |  |
| 8:15.645 | בבית |  |
| 8:15.646 | קדש |  |
| 8:15.647 | הקדשים |  |
| 8:15.648 | מאי |  |
| 8:15.649 | עבידתיה | ? |
| 8:15.650 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

