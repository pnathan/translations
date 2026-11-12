# Translation pack: milhamot-hashem perek-6-26 (Milchamot ha-Shem 6:39–6:40)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-26.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-26`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
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


## Milchamot ha-Shem 6:39

Hebrew (Wikisource): עוד כתב וז"ל ואל יעלה בדעתך כי תוכל להשיגו ית' מצד היותך דורש וחוקר בידיעת מהותו ועצם אמתתו. כי דבר זה איננו אפשרי. ואדרבה הוא מחרף ומגדף מי שחוקר בזה. שאין כח בכל מה שזולתו ית' להשיגו. וכל שכן כל מי שהוא בעל חומר. ועל זה הזהירו רבותינו ז"ל באמרם במופלא ממך בל תדרוש ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:39.1 | עוד |  |
| 6:39.2 | כתב |  |
| 6:39.3 | וז"ל |  |
| 6:39.4 | ואל |  |
| 6:39.5 | יעלה |  |
| 6:39.6 | בדעתך |  |
| 6:39.7 | כי |  |
| 6:39.8 | תוכל |  |
| 6:39.9 | להשיגו |  |
| 6:39.10 | ית | ' |
| 6:39.11 | מצד |  |
| 6:39.12 | היותך |  |
| 6:39.13 | דורש |  |
| 6:39.14 | וחוקר |  |
| 6:39.15 | בידיעת |  |
| 6:39.16 | מהותו |  |
| 6:39.17 | ועצם |  |
| 6:39.18 | אמתתו | . |
| 6:39.19 | כי |  |
| 6:39.20 | דבר |  |
| 6:39.21 | זה |  |
| 6:39.22 | איננו |  |
| 6:39.23 | אפשרי | . |
| 6:39.24 | ואדרבה |  |
| 6:39.25 | הוא |  |
| 6:39.26 | מחרף |  |
| 6:39.27 | ומגדף |  |
| 6:39.28 | מי |  |
| 6:39.29 | שחוקר |  |
| 6:39.30 | בזה | . |
| 6:39.31 | שאין |  |
| 6:39.32 | כח |  |
| 6:39.33 | בכל |  |
| 6:39.34 | מה |  |
| 6:39.35 | שזולתו |  |
| 6:39.36 | ית | ' |
| 6:39.37 | להשיגו | . |
| 6:39.38 | וכל |  |
| 6:39.39 | שכן |  |
| 6:39.40 | כל |  |
| 6:39.41 | מי |  |
| 6:39.42 | שהוא |  |
| 6:39.43 | בעל |  |
| 6:39.44 | חומר | . |
| 6:39.45 | ועל |  |
| 6:39.46 | זה |  |
| 6:39.47 | הזהירו |  |
| 6:39.48 | רבותינו |  |
| 6:39.49 | ז"ל |  |
| 6:39.50 | באמרם |  |
| 6:39.51 | במופלא |  |
| 6:39.52 | ממך |  |
| 6:39.53 | בל |  |
| 6:39.54 | תדרוש |  |
| 6:39.55 | ע"כ | . |


## Milchamot ha-Shem 6:40

Hebrew (Wikisource): עח) עיניך הרואות קורא נעים, איך רבו' המקובלים החדשים בהמשכם בעצימת עינים אחר המסית, ששגו לדבר במהותו יתברך לומר שהוא אור צח ומצוחצח, ונתנו לו ית' תבנית ותמונה, ושנויים רבים, יחסו לו באמרם כי בתחלה היה ממלא כל החלל הזה, ואחר כך צמצם את עצמו לצדדין סביב להמציא מקום וחלל לברואיו, ונעשה ככדור עגול וחלול באמצעיתו, ואחר כך נשתלשל ממנו קו וצנור דק, והיה יורד מעט ומתעגל ונעשו ספירות דעיגולים, ואחר כך נמשך ממנו דרך אותו צינור ונעשו גם ספירות דיושר בצורת אדם שלם בעל קומה ברמ"ח איברים ושס"ה גידים כמו שכתבתי לעיל בקיצור, והרוצה לעמוד על ביאורן יעיין בעץ חיים להרח"ו ושער ההקדמות וכסא אליהו ועז לאלהים ושאר ספרי קבלה. ולא הספיק להם דמות ותבנית וצורה ותמונה אחת (כאשר אמרו רז"ל מי לי בשמים. ועמך לא חפצתי בארץ ויובא לקמן סי' פ"ה) אלא פרצופים רבים בראו להם בדימיונם זה למעלה מזה וכל אחד נוטל רשות מאלוה שלמעלה ממנו. כי עשו את כל אחד מהם עלה וסבה לאלוה שלמטה ממנו. ועשו הסבה הראשונה אדם קדמון. ואמרו עליו שהוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין למעלה ממנו ליטול רשות ממנו. ואחריו עתיקא קדישא כלל זכר ונקבה בחד פרצוף. ואחריו אריך אנפין כולל ג"כ זכר ונקבה בפרצוף אחד. אלא שזה הזכר והנקבה פנים ואחור, וזה ימין ושמאל, כמבואר הכל בדבריהם. ואחריו אבא ואימא ג"כ זכר ונקבה בשני פרצופים נפרדים זה מזה, ואחר אבא ואימא זעיר ונוקביה שני פרצופים ג"כ נפרדים וגם אלו הפרצופים נחלקים עוד לשנים עשר פרצופים פרטיים, וכל אחד מן הפרצופים הנזכרים נקרא בשם הקב"ה ובשם הוי'ה ואדנות. ואלו הן השנים עשר פרצופים. עתיק ונוקביה, אריך אנפין ונוקביה, אבא ואימא עילאין, ישראל סבא ותבונה, זעיר ונוקביה, יעקב ולאה (עץ חיים שער התיקון פרק א'). גם נתנו שיעור וקצב לאלו הפרצופים שזה מגיע עד החזה של פרצוף למעלה ממנו, וזה מגיע עד הטבור וכיוצא באלו העניינים שהם ממקרי הגוף ומשיגיו, והנה שכחו באמרם שאדם קדמון הוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי, שאין לו ממי ליטול רשות, ואיה כל הפרצופים שבעולמות העליונים שהם אלפים ורבבות, וכל אחד כלול מאבי"ע? ולמה הוא אין צריך ליטול רשות ממה שלמעלה ממנו? ולמה לא הניחו רשות ושליטה לאדם קדמאה שהוא למעלה ממנו, ולא לאין סוף העולה על כולם? האם לא נמצאו מחרפים ומגדפים באמרם כי מהותו אור כמ"ש הר' הנז'? ולא באים לידי מינות בקראם לכל עלה מאלה העלות בשם הקב"ה והוי'ה ואלהים? וכיון שהם מאמינים באלוהות רבים ואומרים שהכל אחד, מה הפרש יש בין אמונתם לאמונת העכו"ם שהם שלשה והשלשה אחד? ורז"ל הרחיקו לאמונה זו כדגרסינן בע"ז והבאתיו לעיל סי' ס"ד, ולפי דבריהם למה לא הניחו ר' ישמעאל לבן דמא להתרפאת מיעקב איש כפר סכניא כיון שאומרים שהשלשה אחד? ואיין עוד בסי' ס"ד מה שכתבתי בש' ירושלמי שבת פרק ח' שרצים.

| id | word | punctuation after |
|---|---|---|
| 6:40.1 | עח | ) |
| 6:40.2 | עיניך |  |
| 6:40.3 | הרואות |  |
| 6:40.4 | קורא |  |
| 6:40.5 | נעים | , |
| 6:40.6 | איך |  |
| 6:40.7 | רבו | ' |
| 6:40.8 | המקובלים |  |
| 6:40.9 | החדשים |  |
| 6:40.10 | בהמשכם |  |
| 6:40.11 | בעצימת |  |
| 6:40.12 | עינים |  |
| 6:40.13 | אחר |  |
| 6:40.14 | המסית | , |
| 6:40.15 | ששגו |  |
| 6:40.16 | לדבר |  |
| 6:40.17 | במהותו |  |
| 6:40.18 | יתברך |  |
| 6:40.19 | לומר |  |
| 6:40.20 | שהוא |  |
| 6:40.21 | אור |  |
| 6:40.22 | צח |  |
| 6:40.23 | ומצוחצח | , |
| 6:40.24 | ונתנו |  |
| 6:40.25 | לו |  |
| 6:40.26 | ית | ' |
| 6:40.27 | תבנית |  |
| 6:40.28 | ותמונה | , |
| 6:40.29 | ושנויים |  |
| 6:40.30 | רבים | , |
| 6:40.31 | יחסו |  |
| 6:40.32 | לו |  |
| 6:40.33 | באמרם |  |
| 6:40.34 | כי |  |
| 6:40.35 | בתחלה |  |
| 6:40.36 | היה |  |
| 6:40.37 | ממלא |  |
| 6:40.38 | כל |  |
| 6:40.39 | החלל |  |
| 6:40.40 | הזה | , |
| 6:40.41 | ואחר |  |
| 6:40.42 | כך |  |
| 6:40.43 | צמצם |  |
| 6:40.44 | את |  |
| 6:40.45 | עצמו |  |
| 6:40.46 | לצדדין |  |
| 6:40.47 | סביב |  |
| 6:40.48 | להמציא |  |
| 6:40.49 | מקום |  |
| 6:40.50 | וחלל |  |
| 6:40.51 | לברואיו | , |
| 6:40.52 | ונעשה |  |
| 6:40.53 | ככדור |  |
| 6:40.54 | עגול |  |
| 6:40.55 | וחלול |  |
| 6:40.56 | באמצעיתו | , |
| 6:40.57 | ואחר |  |
| 6:40.58 | כך |  |
| 6:40.59 | נשתלשל |  |
| 6:40.60 | ממנו |  |
| 6:40.61 | קו |  |
| 6:40.62 | וצנור |  |
| 6:40.63 | דק | , |
| 6:40.64 | והיה |  |
| 6:40.65 | יורד |  |
| 6:40.66 | מעט |  |
| 6:40.67 | ומתעגל |  |
| 6:40.68 | ונעשו |  |
| 6:40.69 | ספירות |  |
| 6:40.70 | דעיגולים | , |
| 6:40.71 | ואחר |  |
| 6:40.72 | כך |  |
| 6:40.73 | נמשך |  |
| 6:40.74 | ממנו |  |
| 6:40.75 | דרך |  |
| 6:40.76 | אותו |  |
| 6:40.77 | צינור |  |
| 6:40.78 | ונעשו |  |
| 6:40.79 | גם |  |
| 6:40.80 | ספירות |  |
| 6:40.81 | דיושר |  |
| 6:40.82 | בצורת |  |
| 6:40.83 | אדם |  |
| 6:40.84 | שלם |  |
| 6:40.85 | בעל |  |
| 6:40.86 | קומה |  |
| 6:40.87 | ברמ"ח |  |
| 6:40.88 | איברים |  |
| 6:40.89 | ושס"ה |  |
| 6:40.90 | גידים |  |
| 6:40.91 | כמו |  |
| 6:40.92 | שכתבתי |  |
| 6:40.93 | לעיל |  |
| 6:40.94 | בקיצור | , |
| 6:40.95 | והרוצה |  |
| 6:40.96 | לעמוד |  |
| 6:40.97 | על |  |
| 6:40.98 | ביאורן |  |
| 6:40.99 | יעיין |  |
| 6:40.100 | בעץ |  |
| 6:40.101 | חיים |  |
| 6:40.102 | להרח"ו |  |
| 6:40.103 | ושער |  |
| 6:40.104 | ההקדמות |  |
| 6:40.105 | וכסא |  |
| 6:40.106 | אליהו |  |
| 6:40.107 | ועז |  |
| 6:40.108 | לאלהים |  |
| 6:40.109 | ושאר |  |
| 6:40.110 | ספרי |  |
| 6:40.111 | קבלה | . |
| 6:40.112 | ולא |  |
| 6:40.113 | הספיק |  |
| 6:40.114 | להם |  |
| 6:40.115 | דמות |  |
| 6:40.116 | ותבנית |  |
| 6:40.117 | וצורה |  |
| 6:40.118 | ותמונה |  |
| 6:40.119 | אחת |  |
| 6:40.120 | כאשר |  |
| 6:40.121 | אמרו |  |
| 6:40.122 | רז"ל |  |
| 6:40.123 | מי |  |
| 6:40.124 | לי |  |
| 6:40.125 | בשמים | . |
| 6:40.126 | ועמך |  |
| 6:40.127 | לא |  |
| 6:40.128 | חפצתי |  |
| 6:40.129 | בארץ |  |
| 6:40.130 | ויובא |  |
| 6:40.131 | לקמן |  |
| 6:40.132 | סי | ' |
| 6:40.133 | פ"ה | ) |
| 6:40.134 | אלא |  |
| 6:40.135 | פרצופים |  |
| 6:40.136 | רבים |  |
| 6:40.137 | בראו |  |
| 6:40.138 | להם |  |
| 6:40.139 | בדימיונם |  |
| 6:40.140 | זה |  |
| 6:40.141 | למעלה |  |
| 6:40.142 | מזה |  |
| 6:40.143 | וכל |  |
| 6:40.144 | אחד |  |
| 6:40.145 | נוטל |  |
| 6:40.146 | רשות |  |
| 6:40.147 | מאלוה |  |
| 6:40.148 | שלמעלה |  |
| 6:40.149 | ממנו | . |
| 6:40.150 | כי |  |
| 6:40.151 | עשו |  |
| 6:40.152 | את |  |
| 6:40.153 | כל |  |
| 6:40.154 | אחד |  |
| 6:40.155 | מהם |  |
| 6:40.156 | עלה |  |
| 6:40.157 | וסבה |  |
| 6:40.158 | לאלוה |  |
| 6:40.159 | שלמטה |  |
| 6:40.160 | ממנו | . |
| 6:40.161 | ועשו |  |
| 6:40.162 | הסבה |  |
| 6:40.163 | הראשונה |  |
| 6:40.164 | אדם |  |
| 6:40.165 | קדמון | . |
| 6:40.166 | ואמרו |  |
| 6:40.167 | עליו |  |
| 6:40.168 | שהוא |  |
| 6:40.169 | שאמר |  |
| 6:40.170 | ראו |  |
| 6:40.171 | עתה |  |
| 6:40.172 | כי |  |
| 6:40.173 | אני |  |
| 6:40.174 | אני |  |
| 6:40.175 | הוא |  |
| 6:40.176 | ואין |  |
| 6:40.177 | אלהים |  |
| 6:40.178 | עמדי | , |
| 6:40.179 | שאין |  |
| 6:40.180 | למעלה |  |
| 6:40.181 | ממנו |  |
| 6:40.182 | ליטול |  |
| 6:40.183 | רשות |  |
| 6:40.184 | ממנו | . |
| 6:40.185 | ואחריו |  |
| 6:40.186 | עתיקא |  |
| 6:40.187 | קדישא |  |
| 6:40.188 | כלל |  |
| 6:40.189 | זכר |  |
| 6:40.190 | ונקבה |  |
| 6:40.191 | בחד |  |
| 6:40.192 | פרצוף | . |
| 6:40.193 | ואחריו |  |
| 6:40.194 | אריך |  |
| 6:40.195 | אנפין |  |
| 6:40.196 | כולל |  |
| 6:40.197 | ג"כ |  |
| 6:40.198 | זכר |  |
| 6:40.199 | ונקבה |  |
| 6:40.200 | בפרצוף |  |
| 6:40.201 | אחד | . |
| 6:40.202 | אלא |  |
| 6:40.203 | שזה |  |
| 6:40.204 | הזכר |  |
| 6:40.205 | והנקבה |  |
| 6:40.206 | פנים |  |
| 6:40.207 | ואחור | , |
| 6:40.208 | וזה |  |
| 6:40.209 | ימין |  |
| 6:40.210 | ושמאל | , |
| 6:40.211 | כמבואר |  |
| 6:40.212 | הכל |  |
| 6:40.213 | בדבריהם | . |
| 6:40.214 | ואחריו |  |
| 6:40.215 | אבא |  |
| 6:40.216 | ואימא |  |
| 6:40.217 | ג"כ |  |
| 6:40.218 | זכר |  |
| 6:40.219 | ונקבה |  |
| 6:40.220 | בשני |  |
| 6:40.221 | פרצופים |  |
| 6:40.222 | נפרדים |  |
| 6:40.223 | זה |  |
| 6:40.224 | מזה | , |
| 6:40.225 | ואחר |  |
| 6:40.226 | אבא |  |
| 6:40.227 | ואימא |  |
| 6:40.228 | זעיר |  |
| 6:40.229 | ונוקביה |  |
| 6:40.230 | שני |  |
| 6:40.231 | פרצופים |  |
| 6:40.232 | ג"כ |  |
| 6:40.233 | נפרדים |  |
| 6:40.234 | וגם |  |
| 6:40.235 | אלו |  |
| 6:40.236 | הפרצופים |  |
| 6:40.237 | נחלקים |  |
| 6:40.238 | עוד |  |
| 6:40.239 | לשנים |  |
| 6:40.240 | עשר |  |
| 6:40.241 | פרצופים |  |
| 6:40.242 | פרטיים | , |
| 6:40.243 | וכל |  |
| 6:40.244 | אחד |  |
| 6:40.245 | מן |  |
| 6:40.246 | הפרצופים |  |
| 6:40.247 | הנזכרים |  |
| 6:40.248 | נקרא |  |
| 6:40.249 | בשם |  |
| 6:40.250 | הקב"ה |  |
| 6:40.251 | ובשם |  |
| 6:40.252 | הוי'ה |  |
| 6:40.253 | ואדנות | . |
| 6:40.254 | ואלו |  |
| 6:40.255 | הן |  |
| 6:40.256 | השנים |  |
| 6:40.257 | עשר |  |
| 6:40.258 | פרצופים | . |
| 6:40.259 | עתיק |  |
| 6:40.260 | ונוקביה | , |
| 6:40.261 | אריך |  |
| 6:40.262 | אנפין |  |
| 6:40.263 | ונוקביה | , |
| 6:40.264 | אבא |  |
| 6:40.265 | ואימא |  |
| 6:40.266 | עילאין | , |
| 6:40.267 | ישראל |  |
| 6:40.268 | סבא |  |
| 6:40.269 | ותבונה | , |
| 6:40.270 | זעיר |  |
| 6:40.271 | ונוקביה | , |
| 6:40.272 | יעקב |  |
| 6:40.273 | ולאה |  |
| 6:40.274 | עץ |  |
| 6:40.275 | חיים |  |
| 6:40.276 | שער |  |
| 6:40.277 | התיקון |  |
| 6:40.278 | פרק |  |
| 6:40.279 | א | '). |
| 6:40.280 | גם |  |
| 6:40.281 | נתנו |  |
| 6:40.282 | שיעור |  |
| 6:40.283 | וקצב |  |
| 6:40.284 | לאלו |  |
| 6:40.285 | הפרצופים |  |
| 6:40.286 | שזה |  |
| 6:40.287 | מגיע |  |
| 6:40.288 | עד |  |
| 6:40.289 | החזה |  |
| 6:40.290 | של |  |
| 6:40.291 | פרצוף |  |
| 6:40.292 | למעלה |  |
| 6:40.293 | ממנו | , |
| 6:40.294 | וזה |  |
| 6:40.295 | מגיע |  |
| 6:40.296 | עד |  |
| 6:40.297 | הטבור |  |
| 6:40.298 | וכיוצא |  |
| 6:40.299 | באלו |  |
| 6:40.300 | העניינים |  |
| 6:40.301 | שהם |  |
| 6:40.302 | ממקרי |  |
| 6:40.303 | הגוף |  |
| 6:40.304 | ומשיגיו | , |
| 6:40.305 | והנה |  |
| 6:40.306 | שכחו |  |
| 6:40.307 | באמרם |  |
| 6:40.308 | שאדם |  |
| 6:40.309 | קדמון |  |
| 6:40.310 | הוא |  |
| 6:40.311 | שאמר |  |
| 6:40.312 | ראו |  |
| 6:40.313 | עתה |  |
| 6:40.314 | כי |  |
| 6:40.315 | אני |  |
| 6:40.316 | אני |  |
| 6:40.317 | הוא |  |
| 6:40.318 | ואין |  |
| 6:40.319 | אלהים |  |
| 6:40.320 | עמדי | , |
| 6:40.321 | שאין |  |
| 6:40.322 | לו |  |
| 6:40.323 | ממי |  |
| 6:40.324 | ליטול |  |
| 6:40.325 | רשות | , |
| 6:40.326 | ואיה |  |
| 6:40.327 | כל |  |
| 6:40.328 | הפרצופים |  |
| 6:40.329 | שבעולמות |  |
| 6:40.330 | העליונים |  |
| 6:40.331 | שהם |  |
| 6:40.332 | אלפים |  |
| 6:40.333 | ורבבות | , |
| 6:40.334 | וכל |  |
| 6:40.335 | אחד |  |
| 6:40.336 | כלול |  |
| 6:40.337 | מאבי"ע | ? |
| 6:40.338 | ולמה |  |
| 6:40.339 | הוא |  |
| 6:40.340 | אין |  |
| 6:40.341 | צריך |  |
| 6:40.342 | ליטול |  |
| 6:40.343 | רשות |  |
| 6:40.344 | ממה |  |
| 6:40.345 | שלמעלה |  |
| 6:40.346 | ממנו | ? |
| 6:40.347 | ולמה |  |
| 6:40.348 | לא |  |
| 6:40.349 | הניחו |  |
| 6:40.350 | רשות |  |
| 6:40.351 | ושליטה |  |
| 6:40.352 | לאדם |  |
| 6:40.353 | קדמאה |  |
| 6:40.354 | שהוא |  |
| 6:40.355 | למעלה |  |
| 6:40.356 | ממנו | , |
| 6:40.357 | ולא |  |
| 6:40.358 | לאין |  |
| 6:40.359 | סוף |  |
| 6:40.360 | העולה |  |
| 6:40.361 | על |  |
| 6:40.362 | כולם | ? |
| 6:40.363 | האם |  |
| 6:40.364 | לא |  |
| 6:40.365 | נמצאו |  |
| 6:40.366 | מחרפים |  |
| 6:40.367 | ומגדפים |  |
| 6:40.368 | באמרם |  |
| 6:40.369 | כי |  |
| 6:40.370 | מהותו |  |
| 6:40.371 | אור |  |
| 6:40.372 | כמ"ש |  |
| 6:40.373 | הר | ' |
| 6:40.374 | הנז | '? |
| 6:40.375 | ולא |  |
| 6:40.376 | באים |  |
| 6:40.377 | לידי |  |
| 6:40.378 | מינות |  |
| 6:40.379 | בקראם |  |
| 6:40.380 | לכל |  |
| 6:40.381 | עלה |  |
| 6:40.382 | מאלה |  |
| 6:40.383 | העלות |  |
| 6:40.384 | בשם |  |
| 6:40.385 | הקב"ה |  |
| 6:40.386 | והוי'ה |  |
| 6:40.387 | ואלהים | ? |
| 6:40.388 | וכיון |  |
| 6:40.389 | שהם |  |
| 6:40.390 | מאמינים |  |
| 6:40.391 | באלוהות |  |
| 6:40.392 | רבים |  |
| 6:40.393 | ואומרים |  |
| 6:40.394 | שהכל |  |
| 6:40.395 | אחד | , |
| 6:40.396 | מה |  |
| 6:40.397 | הפרש |  |
| 6:40.398 | יש |  |
| 6:40.399 | בין |  |
| 6:40.400 | אמונתם |  |
| 6:40.401 | לאמונת |  |
| 6:40.402 | העכו"ם |  |
| 6:40.403 | שהם |  |
| 6:40.404 | שלשה |  |
| 6:40.405 | והשלשה |  |
| 6:40.406 | אחד | ? |
| 6:40.407 | ורז"ל |  |
| 6:40.408 | הרחיקו |  |
| 6:40.409 | לאמונה |  |
| 6:40.410 | זו |  |
| 6:40.411 | כדגרסינן |  |
| 6:40.412 | בע"ז |  |
| 6:40.413 | והבאתיו |  |
| 6:40.414 | לעיל |  |
| 6:40.415 | סי | ' |
| 6:40.416 | ס"ד | , |
| 6:40.417 | ולפי |  |
| 6:40.418 | דבריהם |  |
| 6:40.419 | למה |  |
| 6:40.420 | לא |  |
| 6:40.421 | הניחו |  |
| 6:40.422 | ר | ' |
| 6:40.423 | ישמעאל |  |
| 6:40.424 | לבן |  |
| 6:40.425 | דמא |  |
| 6:40.426 | להתרפאת |  |
| 6:40.427 | מיעקב |  |
| 6:40.428 | איש |  |
| 6:40.429 | כפר |  |
| 6:40.430 | סכניא |  |
| 6:40.431 | כיון |  |
| 6:40.432 | שאומרים |  |
| 6:40.433 | שהשלשה |  |
| 6:40.434 | אחד | ? |
| 6:40.435 | ואיין |  |
| 6:40.436 | עוד |  |
| 6:40.437 | בסי | ' |
| 6:40.438 | ס"ד |  |
| 6:40.439 | מה |  |
| 6:40.440 | שכתבתי |  |
| 6:40.441 | בש | ' |
| 6:40.442 | ירושלמי |  |
| 6:40.443 | שבת |  |
| 6:40.444 | פרק |  |
| 6:40.445 | ח | ' |
| 6:40.446 | שרצים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

