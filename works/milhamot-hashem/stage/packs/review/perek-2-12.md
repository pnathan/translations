# Review pack: milhamot-hashem perek-2-12 (Milchamot ha-Shem 2:28–2:28)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-12`.

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
- **המשרף** (pun on Metzaref ha-Emunah) → "the Mesharef", a note on the pun.
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
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **תשבורת** → "mathematics"; **ותוכנים** → "astronomy".
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


## Milchamot ha-Shem 2:28

Hebrew (Wikisource): תשובה לאדם מערכי לב ומה' מענה לשון. ויהי בראותי את מכתבך ישבתי משתומם ומתבהל. איך מלאך לבך לדרוש ולחקור אחר חבל גאונים, אשר הם בתורה נאמנים גדולי היחס כזרע האיתנים, ואחר כמה פסקנים וגם דרשנים ראשונים עם אחרונים חכמי מערב ואשכנז וספרד ותימנים. ושאר קהלות זרע קדש בני אמונים, זאת היא חמדתם מפז ורב פנינים. וכי תבוא תקלה מיד צדיקים נאמנים יודעי מדע חכמים ונבונים יראי אלהים וסרים מרע באונים. בחכמתם הפילו כמה צרים ומונים העומדים עליהם בכל זמנים. וכי קבלתם מרוקחת בהבל ושגיונים אשר לרוב חכמתם אושרו בפי שרים ורוזנים התיצבו לפני מלכים ושלטונים וכי ילכו בדרך לא טוב אבות ובנים? זה כמה מאות שנים? וכמה הם מופלגים בכל חכמה וספרנים מקובלים עם פילוסופיא והגיונים, רפואה ותשבורת ותוכנים. הוגים יומם ולילה שקטים ושאננים, אוכלים פרי מעשיהם הטוב דשנים ורעננים. חלילה! זה לא יאמן ולא יעלה על לב כי אם לפתיים, לכל דבר מאמינים, אשר על שכלם נשענים. ומן הראוי אין עלי להשיבך. כי מה אני להכנס לפני ולפנים מה שלא הורשתי להרהר כל שכן להשיב שואלים. אבל משום עת לעשות לה' אני משיבך ממה שמצאתי כתוב בדברי חז"ל. וזכותם יעמד לי שלא יארע שום תקלה על ידי. הנה מה שכתבת שאינך מבזה חכמי ישראל כבר יש עליך כמה עדים, לא תוכל להכחיש. ומה שהביא אותך לידי כך שאתה מיחס להם טעות ומינות. אלא מפני שהם בזויים בעיניך. אבל האדם יחוש לעצמו פן יכוה בגחלתן כי גדולים צדיקים במיתתן יותר מבחייהן כמו שאמרו בתלמוד. ומה שכתבת שאם לגלות מקום הסתירה וגם התלמוד מביאים ראיה לסתור דברי גדולים מהם, אין הנדון דומה לראיה, דהיכן מצינו שמביאים ראיה לסתור הלכה למשה מסיני? וחכמי הקבלה אומרים כך קיבלו מפי משה, ומי הוא זה ואיפה הוא שנתן לך רשות לחקור ולדרוש אחר דבריהם? ומה שכתבת שהם קורים לפרצופים אלוהות ח"ו שהם אומרים כך. ומה שכתבת צא ולמד מכוונת הברכות והתפלות. כבר ביארו שעילת כל העלות נתכנה בספירה פלונית בשם כן. ובספירה פלונית בשם כך. ומה שכתבת שביארו המקובלים שכמה עולמות יש אלפי אלפים וכו' אפילו יהיו מליוני מליונים כיון שהם אומ' שבוראם אחד והוא מחיה את כולם מה יש הפסד מזה? ומה שכתבת שהם בציור אדם בעל קומה, כבר אמרו גדול כחן של נביאים שמדמין דמות גבורה של מעלה לצורת אדם. ומה שכתבת ברמ"ח איברים כבר ביארו מה הן האיברים. כי הם רי"ו אותיות של שם בן ע"ב ול"ב נתיבות שהם רומזים לכוחות אלהיים. ומה שכתבת שכבר הודיתי שדעתי שאין שום תפלה שייכא בעילת כל העילות שקר אתה דובר. והלא מתחלת שאלתך אותי על דברי המשרף הישבתיך עיין בכס"א. אבל קבלתינו שהתפלה לא"ס ולהשפיע אל עשר ספירות כמו שכ' לחם שלמה. ומה שאמרתי לך לעיין בכסא אליהו כדי שתדע שדברי המשרף דברי כפירה באמת וראוי לבערו. וח"ו שנמצא בשום ספר כדבריו, אבל דברי כס"א אינם כדבריו ח"ו. כי דברי כסא אליהו אין הזעיר אנפין כי אם כלי לפעולת הבורא. ומה שכ' בדף כ"ה שבקראינו לז"א אנחנו קורין אל כל הפרצופים עם נשמה המסתתרת בהם שהוא אין סוף ב"ה. אין רוצה לומר שאין התפילה לא"ס חלילה, אלא רוצה לומר שעליית התפלה אל הא"ס ב"ה דרך זאת המדה. והראיה על זה מה שכ' בדף י"ז ע"ב וז"ל ומה שכתוב בכוונות התפלה והברכות לכוון בכל ברכה כוונה מיוחדת לספי' מיוחדת לא ח"ו לעצמות הספירה. כי הוא קיצוץ ח"ו, וחלילה לכוון לשום כח פרטי וספירה מיוחדת. אלא הכל לעצמות יחיד א"ס ב"ה כלל כל הכחות כולם. ועוד כ' בדף כ"ה ע"א וכל תפלותינו הם לאין סוף ית'. אך שאין אנחנו יכולים לכנותו בשום כנוי ושם. לזה אנחנו מתפללים על ידי הספירות. כי שם שייכות כל כל הכינוים והשמות. כי כל הכנויים והשמות הם שמות עצמותו המתפשט בספירות. וג"כ כ' בדף כ"ו ואם לא יעשו כן ולא ייחדו תפלתם אלא לז"א לא יהיו נענים. וכל המקובלים כיילי כללא דא וז"ל ומה שתמצא באיזה ספר מהם התפלה לשום ספירה חלילה שיהיה כן. אלא הכוונה שיכוין להעלותם אל הא"ס דרך אותה ספירה.

| id | word | punctuation after |
|---|---|---|
| 2:28.1 | תשובה |  |
| 2:28.2 | לאדם |  |
| 2:28.3 | מערכי |  |
| 2:28.4 | לב |  |
| 2:28.5 | ומה | ' |
| 2:28.6 | מענה |  |
| 2:28.7 | לשון | . |
| 2:28.8 | ויהי |  |
| 2:28.9 | בראותי |  |
| 2:28.10 | את |  |
| 2:28.11 | מכתבך |  |
| 2:28.12 | ישבתי |  |
| 2:28.13 | משתומם |  |
| 2:28.14 | ומתבהל | . |
| 2:28.15 | איך |  |
| 2:28.16 | מלאך |  |
| 2:28.17 | לבך |  |
| 2:28.18 | לדרוש |  |
| 2:28.19 | ולחקור |  |
| 2:28.20 | אחר |  |
| 2:28.21 | חבל |  |
| 2:28.22 | גאונים | , |
| 2:28.23 | אשר |  |
| 2:28.24 | הם |  |
| 2:28.25 | בתורה |  |
| 2:28.26 | נאמנים |  |
| 2:28.27 | גדולי |  |
| 2:28.28 | היחס |  |
| 2:28.29 | כזרע |  |
| 2:28.30 | האיתנים | , |
| 2:28.31 | ואחר |  |
| 2:28.32 | כמה |  |
| 2:28.33 | פסקנים |  |
| 2:28.34 | וגם |  |
| 2:28.35 | דרשנים |  |
| 2:28.36 | ראשונים |  |
| 2:28.37 | עם |  |
| 2:28.38 | אחרונים |  |
| 2:28.39 | חכמי |  |
| 2:28.40 | מערב |  |
| 2:28.41 | ואשכנז |  |
| 2:28.42 | וספרד |  |
| 2:28.43 | ותימנים | . |
| 2:28.44 | ושאר |  |
| 2:28.45 | קהלות |  |
| 2:28.46 | זרע |  |
| 2:28.47 | קדש |  |
| 2:28.48 | בני |  |
| 2:28.49 | אמונים | , |
| 2:28.50 | זאת |  |
| 2:28.51 | היא |  |
| 2:28.52 | חמדתם |  |
| 2:28.53 | מפז |  |
| 2:28.54 | ורב |  |
| 2:28.55 | פנינים | . |
| 2:28.56 | וכי |  |
| 2:28.57 | תבוא |  |
| 2:28.58 | תקלה |  |
| 2:28.59 | מיד |  |
| 2:28.60 | צדיקים |  |
| 2:28.61 | נאמנים |  |
| 2:28.62 | יודעי |  |
| 2:28.63 | מדע |  |
| 2:28.64 | חכמים |  |
| 2:28.65 | ונבונים |  |
| 2:28.66 | יראי |  |
| 2:28.67 | אלהים |  |
| 2:28.68 | וסרים |  |
| 2:28.69 | מרע |  |
| 2:28.70 | באונים | . |
| 2:28.71 | בחכמתם |  |
| 2:28.72 | הפילו |  |
| 2:28.73 | כמה |  |
| 2:28.74 | צרים |  |
| 2:28.75 | ומונים |  |
| 2:28.76 | העומדים |  |
| 2:28.77 | עליהם |  |
| 2:28.78 | בכל |  |
| 2:28.79 | זמנים | . |
| 2:28.80 | וכי |  |
| 2:28.81 | קבלתם |  |
| 2:28.82 | מרוקחת |  |
| 2:28.83 | בהבל |  |
| 2:28.84 | ושגיונים |  |
| 2:28.85 | אשר |  |
| 2:28.86 | לרוב |  |
| 2:28.87 | חכמתם |  |
| 2:28.88 | אושרו |  |
| 2:28.89 | בפי |  |
| 2:28.90 | שרים |  |
| 2:28.91 | ורוזנים |  |
| 2:28.92 | התיצבו |  |
| 2:28.93 | לפני |  |
| 2:28.94 | מלכים |  |
| 2:28.95 | ושלטונים |  |
| 2:28.96 | וכי |  |
| 2:28.97 | ילכו |  |
| 2:28.98 | בדרך |  |
| 2:28.99 | לא |  |
| 2:28.100 | טוב |  |
| 2:28.101 | אבות |  |
| 2:28.102 | ובנים | ? |
| 2:28.103 | זה |  |
| 2:28.104 | כמה |  |
| 2:28.105 | מאות |  |
| 2:28.106 | שנים | ? |
| 2:28.107 | וכמה |  |
| 2:28.108 | הם |  |
| 2:28.109 | מופלגים |  |
| 2:28.110 | בכל |  |
| 2:28.111 | חכמה |  |
| 2:28.112 | וספרנים |  |
| 2:28.113 | מקובלים |  |
| 2:28.114 | עם |  |
| 2:28.115 | פילוסופיא |  |
| 2:28.116 | והגיונים | , |
| 2:28.117 | רפואה |  |
| 2:28.118 | ותשבורת |  |
| 2:28.119 | ותוכנים | . |
| 2:28.120 | הוגים |  |
| 2:28.121 | יומם |  |
| 2:28.122 | ולילה |  |
| 2:28.123 | שקטים |  |
| 2:28.124 | ושאננים | , |
| 2:28.125 | אוכלים |  |
| 2:28.126 | פרי |  |
| 2:28.127 | מעשיהם |  |
| 2:28.128 | הטוב |  |
| 2:28.129 | דשנים |  |
| 2:28.130 | ורעננים | . |
| 2:28.131 | חלילה | ! |
| 2:28.132 | זה |  |
| 2:28.133 | לא |  |
| 2:28.134 | יאמן |  |
| 2:28.135 | ולא |  |
| 2:28.136 | יעלה |  |
| 2:28.137 | על |  |
| 2:28.138 | לב |  |
| 2:28.139 | כי |  |
| 2:28.140 | אם |  |
| 2:28.141 | לפתיים | , |
| 2:28.142 | לכל |  |
| 2:28.143 | דבר |  |
| 2:28.144 | מאמינים | , |
| 2:28.145 | אשר |  |
| 2:28.146 | על |  |
| 2:28.147 | שכלם |  |
| 2:28.148 | נשענים | . |
| 2:28.149 | ומן |  |
| 2:28.150 | הראוי |  |
| 2:28.151 | אין |  |
| 2:28.152 | עלי |  |
| 2:28.153 | להשיבך | . |
| 2:28.154 | כי |  |
| 2:28.155 | מה |  |
| 2:28.156 | אני |  |
| 2:28.157 | להכנס |  |
| 2:28.158 | לפני |  |
| 2:28.159 | ולפנים |  |
| 2:28.160 | מה |  |
| 2:28.161 | שלא |  |
| 2:28.162 | הורשתי |  |
| 2:28.163 | להרהר |  |
| 2:28.164 | כל |  |
| 2:28.165 | שכן |  |
| 2:28.166 | להשיב |  |
| 2:28.167 | שואלים | . |
| 2:28.168 | אבל |  |
| 2:28.169 | משום |  |
| 2:28.170 | עת |  |
| 2:28.171 | לעשות |  |
| 2:28.172 | לה | ' |
| 2:28.173 | אני |  |
| 2:28.174 | משיבך |  |
| 2:28.175 | ממה |  |
| 2:28.176 | שמצאתי |  |
| 2:28.177 | כתוב |  |
| 2:28.178 | בדברי |  |
| 2:28.179 | חז"ל | . |
| 2:28.180 | וזכותם |  |
| 2:28.181 | יעמד |  |
| 2:28.182 | לי |  |
| 2:28.183 | שלא |  |
| 2:28.184 | יארע |  |
| 2:28.185 | שום |  |
| 2:28.186 | תקלה |  |
| 2:28.187 | על |  |
| 2:28.188 | ידי | . |
| 2:28.189 | הנה |  |
| 2:28.190 | מה |  |
| 2:28.191 | שכתבת |  |
| 2:28.192 | שאינך |  |
| 2:28.193 | מבזה |  |
| 2:28.194 | חכמי |  |
| 2:28.195 | ישראל |  |
| 2:28.196 | כבר |  |
| 2:28.197 | יש |  |
| 2:28.198 | עליך |  |
| 2:28.199 | כמה |  |
| 2:28.200 | עדים | , |
| 2:28.201 | לא |  |
| 2:28.202 | תוכל |  |
| 2:28.203 | להכחיש | . |
| 2:28.204 | ומה |  |
| 2:28.205 | שהביא |  |
| 2:28.206 | אותך |  |
| 2:28.207 | לידי |  |
| 2:28.208 | כך |  |
| 2:28.209 | שאתה |  |
| 2:28.210 | מיחס |  |
| 2:28.211 | להם |  |
| 2:28.212 | טעות |  |
| 2:28.213 | ומינות | . |
| 2:28.214 | אלא |  |
| 2:28.215 | מפני |  |
| 2:28.216 | שהם |  |
| 2:28.217 | בזויים |  |
| 2:28.218 | בעיניך | . |
| 2:28.219 | אבל |  |
| 2:28.220 | האדם |  |
| 2:28.221 | יחוש |  |
| 2:28.222 | לעצמו |  |
| 2:28.223 | פן |  |
| 2:28.224 | יכוה |  |
| 2:28.225 | בגחלתן |  |
| 2:28.226 | כי |  |
| 2:28.227 | גדולים |  |
| 2:28.228 | צדיקים |  |
| 2:28.229 | במיתתן |  |
| 2:28.230 | יותר |  |
| 2:28.231 | מבחייהן |  |
| 2:28.232 | כמו |  |
| 2:28.233 | שאמרו |  |
| 2:28.234 | בתלמוד | . |
| 2:28.235 | ומה |  |
| 2:28.236 | שכתבת |  |
| 2:28.237 | שאם |  |
| 2:28.238 | לגלות |  |
| 2:28.239 | מקום |  |
| 2:28.240 | הסתירה |  |
| 2:28.241 | וגם |  |
| 2:28.242 | התלמוד |  |
| 2:28.243 | מביאים |  |
| 2:28.244 | ראיה |  |
| 2:28.245 | לסתור |  |
| 2:28.246 | דברי |  |
| 2:28.247 | גדולים |  |
| 2:28.248 | מהם | , |
| 2:28.249 | אין |  |
| 2:28.250 | הנדון |  |
| 2:28.251 | דומה |  |
| 2:28.252 | לראיה | , |
| 2:28.253 | דהיכן |  |
| 2:28.254 | מצינו |  |
| 2:28.255 | שמביאים |  |
| 2:28.256 | ראיה |  |
| 2:28.257 | לסתור |  |
| 2:28.258 | הלכה |  |
| 2:28.259 | למשה |  |
| 2:28.260 | מסיני | ? |
| 2:28.261 | וחכמי |  |
| 2:28.262 | הקבלה |  |
| 2:28.263 | אומרים |  |
| 2:28.264 | כך |  |
| 2:28.265 | קיבלו |  |
| 2:28.266 | מפי |  |
| 2:28.267 | משה | , |
| 2:28.268 | ומי |  |
| 2:28.269 | הוא |  |
| 2:28.270 | זה |  |
| 2:28.271 | ואיפה |  |
| 2:28.272 | הוא |  |
| 2:28.273 | שנתן |  |
| 2:28.274 | לך |  |
| 2:28.275 | רשות |  |
| 2:28.276 | לחקור |  |
| 2:28.277 | ולדרוש |  |
| 2:28.278 | אחר |  |
| 2:28.279 | דבריהם | ? |
| 2:28.280 | ומה |  |
| 2:28.281 | שכתבת |  |
| 2:28.282 | שהם |  |
| 2:28.283 | קורים |  |
| 2:28.284 | לפרצופים |  |
| 2:28.285 | אלוהות |  |
| 2:28.286 | ח"ו |  |
| 2:28.287 | שהם |  |
| 2:28.288 | אומרים |  |
| 2:28.289 | כך | . |
| 2:28.290 | ומה |  |
| 2:28.291 | שכתבת |  |
| 2:28.292 | צא |  |
| 2:28.293 | ולמד |  |
| 2:28.294 | מכוונת |  |
| 2:28.295 | הברכות |  |
| 2:28.296 | והתפלות | . |
| 2:28.297 | כבר |  |
| 2:28.298 | ביארו |  |
| 2:28.299 | שעילת |  |
| 2:28.300 | כל |  |
| 2:28.301 | העלות |  |
| 2:28.302 | נתכנה |  |
| 2:28.303 | בספירה |  |
| 2:28.304 | פלונית |  |
| 2:28.305 | בשם |  |
| 2:28.306 | כן | . |
| 2:28.307 | ובספירה |  |
| 2:28.308 | פלונית |  |
| 2:28.309 | בשם |  |
| 2:28.310 | כך | . |
| 2:28.311 | ומה |  |
| 2:28.312 | שכתבת |  |
| 2:28.313 | שביארו |  |
| 2:28.314 | המקובלים |  |
| 2:28.315 | שכמה |  |
| 2:28.316 | עולמות |  |
| 2:28.317 | יש |  |
| 2:28.318 | אלפי |  |
| 2:28.319 | אלפים |  |
| 2:28.320 | וכו | ' |
| 2:28.321 | אפילו |  |
| 2:28.322 | יהיו |  |
| 2:28.323 | מליוני |  |
| 2:28.324 | מליונים |  |
| 2:28.325 | כיון |  |
| 2:28.326 | שהם |  |
| 2:28.327 | אומ | ' |
| 2:28.328 | שבוראם |  |
| 2:28.329 | אחד |  |
| 2:28.330 | והוא |  |
| 2:28.331 | מחיה |  |
| 2:28.332 | את |  |
| 2:28.333 | כולם |  |
| 2:28.334 | מה |  |
| 2:28.335 | יש |  |
| 2:28.336 | הפסד |  |
| 2:28.337 | מזה | ? |
| 2:28.338 | ומה |  |
| 2:28.339 | שכתבת |  |
| 2:28.340 | שהם |  |
| 2:28.341 | בציור |  |
| 2:28.342 | אדם |  |
| 2:28.343 | בעל |  |
| 2:28.344 | קומה | , |
| 2:28.345 | כבר |  |
| 2:28.346 | אמרו |  |
| 2:28.347 | גדול |  |
| 2:28.348 | כחן |  |
| 2:28.349 | של |  |
| 2:28.350 | נביאים |  |
| 2:28.351 | שמדמין |  |
| 2:28.352 | דמות |  |
| 2:28.353 | גבורה |  |
| 2:28.354 | של |  |
| 2:28.355 | מעלה |  |
| 2:28.356 | לצורת |  |
| 2:28.357 | אדם | . |
| 2:28.358 | ומה |  |
| 2:28.359 | שכתבת |  |
| 2:28.360 | ברמ"ח |  |
| 2:28.361 | איברים |  |
| 2:28.362 | כבר |  |
| 2:28.363 | ביארו |  |
| 2:28.364 | מה |  |
| 2:28.365 | הן |  |
| 2:28.366 | האיברים | . |
| 2:28.367 | כי |  |
| 2:28.368 | הם |  |
| 2:28.369 | רי"ו |  |
| 2:28.370 | אותיות |  |
| 2:28.371 | של |  |
| 2:28.372 | שם |  |
| 2:28.373 | בן |  |
| 2:28.374 | ע"ב |  |
| 2:28.375 | ול"ב |  |
| 2:28.376 | נתיבות |  |
| 2:28.377 | שהם |  |
| 2:28.378 | רומזים |  |
| 2:28.379 | לכוחות |  |
| 2:28.380 | אלהיים | . |
| 2:28.381 | ומה |  |
| 2:28.382 | שכתבת |  |
| 2:28.383 | שכבר |  |
| 2:28.384 | הודיתי |  |
| 2:28.385 | שדעתי |  |
| 2:28.386 | שאין |  |
| 2:28.387 | שום |  |
| 2:28.388 | תפלה |  |
| 2:28.389 | שייכא |  |
| 2:28.390 | בעילת |  |
| 2:28.391 | כל |  |
| 2:28.392 | העילות |  |
| 2:28.393 | שקר |  |
| 2:28.394 | אתה |  |
| 2:28.395 | דובר | . |
| 2:28.396 | והלא |  |
| 2:28.397 | מתחלת |  |
| 2:28.398 | שאלתך |  |
| 2:28.399 | אותי |  |
| 2:28.400 | על |  |
| 2:28.401 | דברי |  |
| 2:28.402 | המשרף |  |
| 2:28.403 | הישבתיך |  |
| 2:28.404 | עיין |  |
| 2:28.405 | בכס"א | . |
| 2:28.406 | אבל |  |
| 2:28.407 | קבלתינו |  |
| 2:28.408 | שהתפלה |  |
| 2:28.409 | לא"ס |  |
| 2:28.410 | ולהשפיע |  |
| 2:28.411 | אל |  |
| 2:28.412 | עשר |  |
| 2:28.413 | ספירות |  |
| 2:28.414 | כמו |  |
| 2:28.415 | שכ | ' |
| 2:28.416 | לחם |  |
| 2:28.417 | שלמה | . |
| 2:28.418 | ומה |  |
| 2:28.419 | שאמרתי |  |
| 2:28.420 | לך |  |
| 2:28.421 | לעיין |  |
| 2:28.422 | בכסא |  |
| 2:28.423 | אליהו |  |
| 2:28.424 | כדי |  |
| 2:28.425 | שתדע |  |
| 2:28.426 | שדברי |  |
| 2:28.427 | המשרף |  |
| 2:28.428 | דברי |  |
| 2:28.429 | כפירה |  |
| 2:28.430 | באמת |  |
| 2:28.431 | וראוי |  |
| 2:28.432 | לבערו | . |
| 2:28.433 | וח"ו |  |
| 2:28.434 | שנמצא |  |
| 2:28.435 | בשום |  |
| 2:28.436 | ספר |  |
| 2:28.437 | כדבריו | , |
| 2:28.438 | אבל |  |
| 2:28.439 | דברי |  |
| 2:28.440 | כס"א |  |
| 2:28.441 | אינם |  |
| 2:28.442 | כדבריו |  |
| 2:28.443 | ח"ו | . |
| 2:28.444 | כי |  |
| 2:28.445 | דברי |  |
| 2:28.446 | כסא |  |
| 2:28.447 | אליהו |  |
| 2:28.448 | אין |  |
| 2:28.449 | הזעיר |  |
| 2:28.450 | אנפין |  |
| 2:28.451 | כי |  |
| 2:28.452 | אם |  |
| 2:28.453 | כלי |  |
| 2:28.454 | לפעולת |  |
| 2:28.455 | הבורא | . |
| 2:28.456 | ומה |  |
| 2:28.457 | שכ | ' |
| 2:28.458 | בדף |  |
| 2:28.459 | כ"ה |  |
| 2:28.460 | שבקראינו |  |
| 2:28.461 | לז"א |  |
| 2:28.462 | אנחנו |  |
| 2:28.463 | קורין |  |
| 2:28.464 | אל |  |
| 2:28.465 | כל |  |
| 2:28.466 | הפרצופים |  |
| 2:28.467 | עם |  |
| 2:28.468 | נשמה |  |
| 2:28.469 | המסתתרת |  |
| 2:28.470 | בהם |  |
| 2:28.471 | שהוא |  |
| 2:28.472 | אין |  |
| 2:28.473 | סוף |  |
| 2:28.474 | ב"ה | . |
| 2:28.475 | אין |  |
| 2:28.476 | רוצה |  |
| 2:28.477 | לומר |  |
| 2:28.478 | שאין |  |
| 2:28.479 | התפילה |  |
| 2:28.480 | לא"ס |  |
| 2:28.481 | חלילה | , |
| 2:28.482 | אלא |  |
| 2:28.483 | רוצה |  |
| 2:28.484 | לומר |  |
| 2:28.485 | שעליית |  |
| 2:28.486 | התפלה |  |
| 2:28.487 | אל |  |
| 2:28.488 | הא"ס |  |
| 2:28.489 | ב"ה |  |
| 2:28.490 | דרך |  |
| 2:28.491 | זאת |  |
| 2:28.492 | המדה | . |
| 2:28.493 | והראיה |  |
| 2:28.494 | על |  |
| 2:28.495 | זה |  |
| 2:28.496 | מה |  |
| 2:28.497 | שכ | ' |
| 2:28.498 | בדף |  |
| 2:28.499 | י"ז |  |
| 2:28.500 | ע"ב |  |
| 2:28.501 | וז"ל |  |
| 2:28.502 | ומה |  |
| 2:28.503 | שכתוב |  |
| 2:28.504 | בכוונות |  |
| 2:28.505 | התפלה |  |
| 2:28.506 | והברכות |  |
| 2:28.507 | לכוון |  |
| 2:28.508 | בכל |  |
| 2:28.509 | ברכה |  |
| 2:28.510 | כוונה |  |
| 2:28.511 | מיוחדת |  |
| 2:28.512 | לספי | ' |
| 2:28.513 | מיוחדת |  |
| 2:28.514 | לא |  |
| 2:28.515 | ח"ו |  |
| 2:28.516 | לעצמות |  |
| 2:28.517 | הספירה | . |
| 2:28.518 | כי |  |
| 2:28.519 | הוא |  |
| 2:28.520 | קיצוץ |  |
| 2:28.521 | ח"ו | , |
| 2:28.522 | וחלילה |  |
| 2:28.523 | לכוון |  |
| 2:28.524 | לשום |  |
| 2:28.525 | כח |  |
| 2:28.526 | פרטי |  |
| 2:28.527 | וספירה |  |
| 2:28.528 | מיוחדת | . |
| 2:28.529 | אלא |  |
| 2:28.530 | הכל |  |
| 2:28.531 | לעצמות |  |
| 2:28.532 | יחיד |  |
| 2:28.533 | א"ס |  |
| 2:28.534 | ב"ה |  |
| 2:28.535 | כלל |  |
| 2:28.536 | כל |  |
| 2:28.537 | הכחות |  |
| 2:28.538 | כולם | . |
| 2:28.539 | ועוד |  |
| 2:28.540 | כ | ' |
| 2:28.541 | בדף |  |
| 2:28.542 | כ"ה |  |
| 2:28.543 | ע"א |  |
| 2:28.544 | וכל |  |
| 2:28.545 | תפלותינו |  |
| 2:28.546 | הם |  |
| 2:28.547 | לאין |  |
| 2:28.548 | סוף |  |
| 2:28.549 | ית | '. |
| 2:28.550 | אך |  |
| 2:28.551 | שאין |  |
| 2:28.552 | אנחנו |  |
| 2:28.553 | יכולים |  |
| 2:28.554 | לכנותו |  |
| 2:28.555 | בשום |  |
| 2:28.556 | כנוי |  |
| 2:28.557 | ושם | . |
| 2:28.558 | לזה |  |
| 2:28.559 | אנחנו |  |
| 2:28.560 | מתפללים |  |
| 2:28.561 | על |  |
| 2:28.562 | ידי |  |
| 2:28.563 | הספירות | . |
| 2:28.564 | כי |  |
| 2:28.565 | שם |  |
| 2:28.566 | שייכות |  |
| 2:28.567 | כל |  |
| 2:28.568 | כל |  |
| 2:28.569 | הכינוים |  |
| 2:28.570 | והשמות | . |
| 2:28.571 | כי |  |
| 2:28.572 | כל |  |
| 2:28.573 | הכנויים |  |
| 2:28.574 | והשמות |  |
| 2:28.575 | הם |  |
| 2:28.576 | שמות |  |
| 2:28.577 | עצמותו |  |
| 2:28.578 | המתפשט |  |
| 2:28.579 | בספירות | . |
| 2:28.580 | וג"כ |  |
| 2:28.581 | כ | ' |
| 2:28.582 | בדף |  |
| 2:28.583 | כ"ו |  |
| 2:28.584 | ואם |  |
| 2:28.585 | לא |  |
| 2:28.586 | יעשו |  |
| 2:28.587 | כן |  |
| 2:28.588 | ולא |  |
| 2:28.589 | ייחדו |  |
| 2:28.590 | תפלתם |  |
| 2:28.591 | אלא |  |
| 2:28.592 | לז"א |  |
| 2:28.593 | לא |  |
| 2:28.594 | יהיו |  |
| 2:28.595 | נענים | . |
| 2:28.596 | וכל |  |
| 2:28.597 | המקובלים |  |
| 2:28.598 | כיילי |  |
| 2:28.599 | כללא |  |
| 2:28.600 | דא |  |
| 2:28.601 | וז"ל |  |
| 2:28.602 | ומה |  |
| 2:28.603 | שתמצא |  |
| 2:28.604 | באיזה |  |
| 2:28.605 | ספר |  |
| 2:28.606 | מהם |  |
| 2:28.607 | התפלה |  |
| 2:28.608 | לשום |  |
| 2:28.609 | ספירה |  |
| 2:28.610 | חלילה |  |
| 2:28.611 | שיהיה |  |
| 2:28.612 | כן | . |
| 2:28.613 | אלא |  |
| 2:28.614 | הכוונה |  |
| 2:28.615 | שיכוין |  |
| 2:28.616 | להעלותם |  |
| 2:28.617 | אל |  |
| 2:28.618 | הא"ס |  |
| 2:28.619 | דרך |  |
| 2:28.620 | אותה |  |
| 2:28.621 | ספירה | . |


### Translation under review (2:28)

English: Reply. To man the arrangements of the heart, and from the LORD the answer of the tongue. And it was when I saw your letter, I sat astonished and dismayed. How your heart came to inquire and investigate after the band of the Geonim, who are faithful in the Torah, distinguished in lineage, like the seed of the mighty, and after several decisors and also expounders, the early and the later, the sages of the West, Ashkenaz, Sefarad and the Yemenites. and the rest of the congregations, the holy seed, the faithful children, this is their treasure, more than gold and abundant pearls. Can a stumbling come from the hand of righteous and faithful ones, who know knowledge, the wise and discerning, who fear God and turn aside from evil with their might. By their wisdom they felled many foes and multitudes who stood against them in every age. Can it be that their Kabbalah is compounded in vanity and error, which, because of their great wisdom, were approved by the mouths of princes and nobles and stood before kings and rulers, and can fathers and sons walk on a path that is not good? How many hundreds of years is this? And how many of them are distinguished in every wisdom, and scholars, Kabbalists with philosophy and logic, medicine, tishburet and computations. meditating day and night, quiet and at ease, eating the fruit of their deeds, the good, fat and lush. Heaven forbid! This is not to be believed and does not come to mind but only for the simple, who believe every matter, who lean on their understanding. And properly, it is not incumbent on me to answer you. For what am I, to enter before and within, into what I was not permitted to ponder, all the more so to answer questioners. But because it is a time to act for the LORD, I answer you from what I found written in the words of our Sages of blessed memory. May their merit stand for me, that no stumbling befall through me. Now, as for what you wrote, that you do not despise the sages of Israel, there are already several witnesses against you; you cannot deny it. And what brought you to this that you attribute to them error and heresy. Rather, because they are despised in your eyes. But a person should take heed for himself, lest he be burned by their coals, for the righteous are greater in their death than in their life, as they said in the Talmud. And as for what you wrote, that if one reveals the place of contradiction, and the Talmud also brings proof to refute the words of those greater than them, the matter is not comparable to a proof, for where do we find that they bring proof to refute a halakhah given to Moses at Sinai? And the masters of Kabbalah say that they received thus from the mouth of Moses, and who is this one, and where is he, who gave you permission to inquire and investigate after their words? And as for what you wrote, that they call the partzufim divinities, God forbid, that they say thus. And as for what you wrote, go and learn from the intentions of the blessings and the prayers. They already explained that the cause of all causes is designated in a certain sefirah by such a name. And in a certain sefirah by that name. And as for what you wrote, that the Kabbalists explained that there are several worlds, thousands of thousands, etc., even if they were millions of millions, since they say that their Creator is one, and He gives life to all of them: what loss is there from this? And as for what you wrote, that they are in the image of a man of stature, already they said, "Great is the power of the prophets, who liken the form of might above to the shape of a man." And as for what you wrote about 248 limbs, they already explained what the limbs are. For they are 216 letters of the Name of 72 and 32 paths, which allude to divine powers. And as for what you wrote, that I already conceded that my opinion is that there is no prayer applicable to the Cause of all causes, you speak falsehood. But surely, from the beginning of your question to me about the words of the Mesharef, I answered you, see in Kisse Eliyahu. But our Kabbalah holds that prayer is to Ein Sof and bestows upon the ten sefirot, as Lechem Shelomo wrote. And as for what I told you, to consult Kisse Eliyahu, so that you would know that the words of the Mesharef are words of heresy, in truth, and fitting to root it out. And God forbid that there be found in any book words like his, but the words of Kisse Eliyahu are not like his, God forbid. For the words of Kisse Eliyahu hold that Ze'ir Anpin is nothing but a vessel for the action of the Creator. And as for what he wrote on folio 25, that when we call Ze'ir Anpin God, we call all the partzufim God, with a soul hidden in them, which is Ein Sof, blessed be He. This does not mean that prayer is not directed to Ein Sof, Heaven forbid, but it means that the ascent of prayer to Ein Sof, blessed be He, is by way of this attribute. And the proof for this is what he wrote on folio 17b, and these are his words: "As for what is written about the intentions of prayer and the blessings, to direct in each blessing a particular intention, to a particular sefirah, God forbid, not to the essence of the sefirah. For it is a cutting off, God forbid, and Heaven forbid to direct it to any particular power or particular sefirah. But all is to the essence of the one Ein Sof, blessed be He, comprising all the powers, all of them." And furthermore he wrote on folio 25a: "And all our prayers are to Ein Sof, may He be blessed. But that we cannot call Him by any appellation and name. Therefore we pray through the sefirot. For the name pertaining to all, all the appellations and names. For all the appellations and names are names of His essence, extending through the sefirot." Also he wrote on folio 26: "if they do not do so, and do not unify their prayer but to Ze'ir Anpin, they will not be answered." And all the Kabbalists hold this principle, and these are their words: "As for whatever you find in some book of theirs, that prayer is addressed to any sefirah, Heaven forbid that it be so. But the intention is that one direct them to raise them to Ein Sof by way of that sefirah."

Literal: Reply to-man the-arrangements-of the-heart and-from-the-LORD the-answer-of the-tongue and-it-was when-I-saw (object-marker) your-letter I-sat astonished and-dismayed how has-filled your-heart to-inquire and-to-investigate after band of-the-Geonim who they-are in-the-Torah faithful distinguished-of lineage like-seed-of the-mighty-ones and-after several decisors and-also expounders the-early with the-later sages-of the-West and-Ashkenaz and-Sefarad and-the-Yemenites and-the-rest-of congregations seed holy children-of faithful this is their-treasure more-than-gold and-abundant pearls can-it-be-that comes a-stumbling from-the-hand-of righteous-ones faithful-ones who-know knowledge the-wise and-discerning who-fear God and-who-turn-aside from-evil with-their-might by-their-wisdom they-felled many foes and-multitudes who-stood against-them in-every age can-it-be-that their-Kabbalah compounded in-vanity and-error which because-of-the-great their-wisdom were-approved by-the-mouths-of princes and-nobles they-stood before kings and-rulers can-it-be-that they-walk on-the-path not good fathers and-sons this how-many hundreds of-years and-how-many they are-distinguished in-every wisdom and-scholars Kabbalists with philosophy and-logic medicine and-tishburet and-computations meditating by-day and-by-night quiet and-at-ease eating the-fruit-of their-deeds the-good fat and-lush Heaven-forbid this is-not to-be-believed and-not come-up upon the-heart but only for-the-simple to-every matter who-believe who upon their-understanding lean and-properly what-is-fitting there-is-not upon-me to-answer-you for what am-I to-enter before and-within what that-not I-was-permitted to-ponder all the-more-so to-answer questioners but because-of a-time to-act for-the-LORD I answer-you from-what I-found written in-the-words-of our-Sages-of-blessed-memory and-their-merit will-stand for-me that-not will-befall any stumbling through my-hand behold as-for-what you-wrote that-you-do-not despise the-sages-of Israel already there-are against-you several witnesses not you-can to-deny and-what brought you to this that-you attribute to-them error and-heresy but because that-they-are despised in-your-eyes but the-person should-take-heed for-himself lest he-be-burned by-their-coals for great-ones righteous-ones in-their-death more than-in-their-life as they-said in-the-Talmud and-as-for-what you-wrote that-if to-reveal the-place-of the-contradiction and-also the-Talmud brings proof to-refute the-words-of greater than-them there-is-not the-matter comparable to-a-proof for-where do-we-find that-they-bring proof to-refute a-halakhah given-to-Moses at-Sinai and-the-masters-of Kabbalah say thus they-received from-the-mouth-of Moses and-who is-he this-one and-where is-he who-gave you permission to-inquire and-to-investigate after their-words and-as-for-what you-wrote that-they call the-partzufim divinities God-forbid that-they say thus and-as-for-what you-wrote go-out and-learn from-the-intentions-of the-blessings and-the-prayers already they-explained that-the-cause-of all the-causes is-designated in-a-sefirah a-certain by-the-name such and-in-a-sefirah a-certain by-the-name that and-as-for-what you-wrote that-they-explained the-Kabbalists that-several worlds there-are thousands-of thousands etc. even-if they-were millions-of millions since that-they say that-their-Creator is-one and-He gives-life-to (object-marker) all-of-them what is-there loss from-this and-as-for-what you-wrote that-they in-the-image-of a-man of stature already they-said great is-their-power of the-prophets who-liken the-form of-might of above to-the-shape-of a-man and-as-for-what you-wrote about-248 limbs already they-explained what they-are the-limbs for they-are 216 letters of the-Name of 72 and-32 paths which-are they-allude-to the-powers-of divine and-as-for-what you-wrote that-already I-conceded that-my-opinion that-there-is-not any prayer applicable to-the-cause-of all causes falsehood you speak but-surely from-the-beginning-of your-question to-me about the-words-of the-Mesharef I-answered-you see in-Kisse-Eliyahu but our-Kabbalah that-prayer to-Ein-Sof and-to-bestow upon ten sefirot as he-wrote Lechem Shelomo and-as-for-what I-said to-you to-consult in-the-Throne of-Elijah so-that you-would-know that-the-words-of the-Mesharef are-the-words-of heresy in-truth and-fitting to-root-it-out and-God-forbid that-there-be-found in-any book like-his-words but the-words-of Kisse-Eliyahu are-not like-his-words God-forbid for the-words-of the-Throne of-Elijah is-nothing Ze'ir Anpin but only a-vessel for-the-action-of the-Creator and-as-for-what he-wrote on-folio 25 that-when-we call-Ze'ir-Anpin we call God all the-partzufim with a-soul hidden in-them which-is Ein Sof blessed-be-He this-does-not mean to-say that-not prayer to-Ein-Sof Heaven-forbid but it-means to-say the-ascent-of prayer to the-Ein-Sof blessed-be-He by-way-of this attribute and-the-proof for this is-what he-wrote on-folio 17 b and-these-are-his-words: as-for-what is-written about-the-intentions-of prayer and-the-blessings to-direct in-each blessing an-intention particular to-a-sefirah particular not God-forbid to-the-essence-of the-sefirah for it a-cutting-off God-forbid and-Heaven-forbid to-direct to-any power particular and-sefirah particular but all to-the-essence-of the-one Ein-Sof blessed-be-He comprising all the-powers all-of-them and-furthermore he-wrote on-folio 25 a and-all our-prayers are to-Ein Sof may-He-be-blessed but that-not we can call-Him by-any appellation and-name therefore we pray by-means-of the-hand-of the-sefirot for the-name pertaining all all the-appellations and-names for all the-appellations and-the-names are names of-His-essence that-extends through-the-sefirot also he-wrote on-folio 26 if not they-do so and-not unify their-prayer but to-Ze'ir-Anpin not they-will-be answered and-all the-Kabbalists all-hold the-principle this and-these-are-their-words: as-for-whatever you-find in-some book of-theirs prayer to-any sefirah Heaven-forbid that-it-be so but the-intention that-he-direct to-raise-them to Ein-Sof by-way-of that sefirah

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:28.1 | Reply |  |  | תשובה: the heading of a reply in the correspondence of chapter 2. |
| 2:28.2 | to man |  |  |  |
| 2:28.3 | the arrangements of |  |  | מערכי: construct plural, "arrangements of"; the opening alludes to Prov 16:1 (see apparatus). |
| 2:28.4 | the heart |  |  |  |
| 2:28.5 | and from the LORD |  |  | ומה': "and from the LORD"; the apostrophe is the Tetragrammaton's mark in the punctuation column. |
| 2:28.6 | the answer of |  |  |  |
| 2:28.7 | the tongue |  |  |  |
| 2:28.8 | and it was |  |  | ויהי: a wayyiqtol opening the sentence; the narrative runs on in the verbs that follow. |
| 2:28.9 | when I saw |  |  |  |
| 2:28.10 | (object marker) |  |  | את: the object marker, with no English counterpart. |
| 2:28.11 | your letter |  |  |  |
| 2:28.12 | I sat |  |  |  |
| 2:28.13 | astonished |  |  |  |
| 2:28.14 | and dismayed |  |  |  |
| 2:28.15 | how |  |  |  |
| 2:28.16 | has filled |  |  |  |
| 2:28.17 | your heart |  |  |  |
| 2:28.18 | to inquire |  |  |  |
| 2:28.19 | and to investigate |  |  |  |
| 2:28.20 | after |  |  |  |
| 2:28.21 | band |  |  | חבל: "cord, band"; here a company or fellowship. |
| 2:28.22 | of the Geonim |  |  | גאונים: the Geonim, the heads of the Babylonian academies. |
| 2:28.23 | who |  |  |  |
| 2:28.24 | they are |  |  |  |
| 2:28.25 | in the Torah |  |  |  |
| 2:28.26 | faithful |  |  |  |
| 2:28.27 | distinguished of |  |  |  |
| 2:28.28 | lineage |  |  | היחס: "lineage, pedigree"; גדולי היחס: men of distinguished lineage. |
| 2:28.29 | like seed of |  |  |  |
| 2:28.30 | the mighty ones |  |  |  |
| 2:28.31 | and after |  |  |  |
| 2:28.32 | several |  |  |  |
| 2:28.33 | decisors |  |  |  |
| 2:28.34 | and also |  |  |  |
| 2:28.35 | expounders |  |  |  |
| 2:28.36 | the early |  |  |  |
| 2:28.37 | with |  |  |  |
| 2:28.38 | the later |  |  |  |
| 2:28.39 | sages of |  |  |  |
| 2:28.40 | the West |  |  | מערב: literally "the West"; the place is not named further. |
| 2:28.41 | and Ashkenaz |  |  |  |
| 2:28.42 | and Sefarad |  |  |  |
| 2:28.43 | and the Yemenites |  |  |  |
| 2:28.44 | and the rest of |  |  |  |
| 2:28.45 | congregations |  |  |  |
| 2:28.46 | seed |  |  |  |
| 2:28.47 | holy |  |  |  |
| 2:28.48 | children of |  |  |  |
| 2:28.49 | faithful |  |  |  |
| 2:28.50 | this |  |  |  |
| 2:28.51 | is |  |  |  |
| 2:28.52 | their treasure |  |  | חמדתם: "their delight, treasure", the noun of desire. |
| 2:28.53 | more than gold |  |  |  |
| 2:28.54 | and abundant |  |  |  |
| 2:28.55 | pearls |  |  |  |
| 2:28.56 | can it be that |  |  | וכי: the interrogative of a rhetorical question: "Can it be that ...". |
| 2:28.57 | comes |  |  |  |
| 2:28.58 | a stumbling |  |  |  |
| 2:28.59 | from the hand of |  |  |  |
| 2:28.60 | righteous ones |  |  |  |
| 2:28.61 | faithful ones |  |  |  |
| 2:28.62 | who know |  |  |  |
| 2:28.63 | knowledge |  |  |  |
| 2:28.64 | the wise |  |  |  |
| 2:28.65 | and discerning |  |  |  |
| 2:28.66 | who fear |  |  |  |
| 2:28.67 | God |  |  | אלהים: "God"; the allusion is to Job 1:1 (see apparatus). |
| 2:28.68 | and who turn aside |  |  |  |
| 2:28.69 | from evil |  |  |  |
| 2:28.70 | with their might |  |  |  |
| 2:28.71 | by their wisdom |  |  |  |
| 2:28.72 | they felled |  |  |  |
| 2:28.73 | many |  |  |  |
| 2:28.74 | foes |  |  |  |
| 2:28.75 | and multitudes |  |  | ומונים: sense uncertain; rendered "and multitudes". |
| 2:28.76 | who stood |  |  |  |
| 2:28.77 | against them |  |  |  |
| 2:28.78 | in every |  |  |  |
| 2:28.79 | age |  |  |  |
| 2:28.80 | can it be that |  |  |  |
| 2:28.81 | their Kabbalah |  |  | קבלתם: "their Kabbalah", the lore of the Kabbalists, not the rabbinic chain of tradition. |
| 2:28.82 | compounded |  |  |  |
| 2:28.83 | in vanity |  |  |  |
| 2:28.84 | and error |  |  |  |
| 2:28.85 | which |  |  |  |
| 2:28.86 | because of the great |  |  |  |
| 2:28.87 | their wisdom |  |  |  |
| 2:28.88 | were approved |  |  |  |
| 2:28.89 | by the mouths of |  |  |  |
| 2:28.90 | princes |  |  |  |
| 2:28.91 | and nobles |  |  |  |
| 2:28.92 | they stood |  |  |  |
| 2:28.93 | before |  |  |  |
| 2:28.94 | kings |  |  |  |
| 2:28.95 | and rulers |  |  |  |
| 2:28.96 | can it be that |  |  |  |
| 2:28.97 | they walk |  |  |  |
| 2:28.98 | on the path |  |  |  |
| 2:28.99 | not |  |  |  |
| 2:28.100 | good |  |  |  |
| 2:28.101 | fathers |  |  |  |
| 2:28.102 | and sons |  |  |  |
| 2:28.103 | this |  |  |  |
| 2:28.104 | how many |  |  |  |
| 2:28.105 | hundreds |  |  |  |
| 2:28.106 | of years |  |  |  |
| 2:28.107 | and how many |  |  |  |
| 2:28.108 | they |  |  |  |
| 2:28.109 | are distinguished |  |  |  |
| 2:28.110 | in every |  |  |  |
| 2:28.111 | wisdom |  |  |  |
| 2:28.112 | and scholars |  |  |  |
| 2:28.113 | Kabbalists |  |  |  |
| 2:28.114 | with |  |  |  |
| 2:28.115 | philosophy |  |  |  |
| 2:28.116 | and logic |  |  |  |
| 2:28.117 | medicine |  |  |  |
| 2:28.118 | and tishburet |  |  | ותשבורת: a technical term, transliterated; its sense is not settled here. |
| 2:28.119 | and computations |  |  | ותוכנים: read as "computations"; the sense is uncertain. |
| 2:28.120 | meditating |  |  |  |
| 2:28.121 | by day |  |  |  |
| 2:28.122 | and by night |  |  |  |
| 2:28.123 | quiet |  |  |  |
| 2:28.124 | and at ease |  |  |  |
| 2:28.125 | eating |  |  |  |
| 2:28.126 | the fruit of |  |  |  |
| 2:28.127 | their deeds |  |  |  |
| 2:28.128 | the good |  |  |  |
| 2:28.129 | fat |  |  |  |
| 2:28.130 | and lush |  |  |  |
| 2:28.131 | Heaven forbid |  |  | חלילה: "Heaven forbid", the full form of ח"ו. |
| 2:28.132 | this |  |  |  |
| 2:28.133 | is not |  |  |  |
| 2:28.134 | to be believed |  |  |  |
| 2:28.135 | and not |  |  |  |
| 2:28.136 | come up |  |  |  |
| 2:28.137 | upon |  |  |  |
| 2:28.138 | the heart |  |  |  |
| 2:28.139 | but |  |  |  |
| 2:28.140 | only |  |  |  |
| 2:28.141 | for the simple |  |  |  |
| 2:28.142 | to every |  |  |  |
| 2:28.143 | matter |  |  |  |
| 2:28.144 | who believe |  |  |  |
| 2:28.145 | who |  |  |  |
| 2:28.146 | upon |  |  |  |
| 2:28.147 | their understanding |  |  |  |
| 2:28.148 | lean |  |  |  |
| 2:28.149 | and properly |  |  |  |
| 2:28.150 | what is fitting |  |  |  |
| 2:28.151 | there is not |  |  |  |
| 2:28.152 | upon me |  |  |  |
| 2:28.153 | to answer you |  |  |  |
| 2:28.154 | for |  |  |  |
| 2:28.155 | what |  |  |  |
| 2:28.156 | am I |  |  |  |
| 2:28.157 | to enter |  |  |  |
| 2:28.158 | before |  |  |  |
| 2:28.159 | and within |  |  |  |
| 2:28.160 | what |  |  |  |
| 2:28.161 | that not |  |  |  |
| 2:28.162 | I was permitted |  |  |  |
| 2:28.163 | to ponder |  |  |  |
| 2:28.164 | all |  |  |  |
| 2:28.165 | the more so |  |  |  |
| 2:28.166 | to answer |  |  |  |
| 2:28.167 | questioners |  |  |  |
| 2:28.168 | but |  |  |  |
| 2:28.169 | because of |  |  |  |
| 2:28.170 | a time |  |  |  |
| 2:28.171 | to act |  |  |  |
| 2:28.172 | for the LORD |  |  | לה: "for the LORD", the Tetragrammaton; the words are Ps 119:126. |
| 2:28.173 | I |  |  |  |
| 2:28.174 | answer you |  |  |  |
| 2:28.175 | from what |  |  |  |
| 2:28.176 | I found |  |  |  |
| 2:28.177 | written |  |  |  |
| 2:28.178 | in the words of |  |  |  |
| 2:28.179 | our Sages of blessed memory |  |  | חז"ל: "our Sages of blessed memory". |
| 2:28.180 | and their merit |  |  |  |
| 2:28.181 | will stand |  |  |  |
| 2:28.182 | for me |  |  |  |
| 2:28.183 | that not |  |  |  |
| 2:28.184 | will befall |  |  |  |
| 2:28.185 | any |  |  |  |
| 2:28.186 | stumbling |  |  |  |
| 2:28.187 | through |  |  |  |
| 2:28.188 | my hand |  |  |  |
| 2:28.189 | behold |  |  |  |
| 2:28.190 | as for what |  |  |  |
| 2:28.191 | you wrote |  |  |  |
| 2:28.192 | that you do not |  |  |  |
| 2:28.193 | despise |  |  |  |
| 2:28.194 | the sages of |  |  |  |
| 2:28.195 | Israel |  |  |  |
| 2:28.196 | already |  |  |  |
| 2:28.197 | there are |  |  |  |
| 2:28.198 | against you |  |  |  |
| 2:28.199 | several |  |  |  |
| 2:28.200 | witnesses |  |  |  |
| 2:28.201 | not |  |  |  |
| 2:28.202 | you can |  |  |  |
| 2:28.203 | to deny |  |  |  |
| 2:28.204 | and what |  |  |  |
| 2:28.205 | brought |  |  |  |
| 2:28.206 | you |  |  |  |
| 2:28.207 | to |  |  |  |
| 2:28.208 | this |  |  |  |
| 2:28.209 | that you |  |  |  |
| 2:28.210 | attribute |  |  |  |
| 2:28.211 | to them |  |  |  |
| 2:28.212 | error |  |  |  |
| 2:28.213 | and heresy |  |  |  |
| 2:28.214 | but |  |  |  |
| 2:28.215 | because |  |  |  |
| 2:28.216 | that they are |  |  |  |
| 2:28.217 | despised |  |  |  |
| 2:28.218 | in your eyes |  |  |  |
| 2:28.219 | but |  |  |  |
| 2:28.220 | the person |  |  |  |
| 2:28.221 | should take heed |  |  |  |
| 2:28.222 | for himself |  |  |  |
| 2:28.223 | lest |  |  |  |
| 2:28.224 | he be burned |  |  |  |
| 2:28.225 | by their coals |  |  | בגחלתן: "in their coals", after Avot 2:10 (see apparatus). |
| 2:28.226 | for |  |  |  |
| 2:28.227 | great ones |  |  |  |
| 2:28.228 | righteous ones |  |  |  |
| 2:28.229 | in their death |  |  |  |
| 2:28.230 | more |  |  |  |
| 2:28.231 | than in their life |  |  |  |
| 2:28.232 | as |  |  |  |
| 2:28.233 | they said |  |  |  |
| 2:28.234 | in the Talmud |  |  |  |
| 2:28.235 | and as for what |  |  |  |
| 2:28.236 | you wrote |  |  |  |
| 2:28.237 | that if |  |  |  |
| 2:28.238 | to reveal |  |  |  |
| 2:28.239 | the place of |  |  |  |
| 2:28.240 | the contradiction |  |  |  |
| 2:28.241 | and also |  |  |  |
| 2:28.242 | the Talmud |  |  |  |
| 2:28.243 | brings |  |  |  |
| 2:28.244 | proof |  |  |  |
| 2:28.245 | to refute |  |  |  |
| 2:28.246 | the words of |  |  |  |
| 2:28.247 | greater |  |  |  |
| 2:28.248 | than them |  |  |  |
| 2:28.249 | there is not |  |  |  |
| 2:28.250 | the matter |  |  |  |
| 2:28.251 | comparable |  |  |  |
| 2:28.252 | to a proof |  |  |  |
| 2:28.253 | for where |  |  |  |
| 2:28.254 | do we find |  |  |  |
| 2:28.255 | that they bring |  |  |  |
| 2:28.256 | proof |  |  |  |
| 2:28.257 | to refute |  |  |  |
| 2:28.258 | a halakhah |  |  |  |
| 2:28.259 | given to Moses |  |  |  |
| 2:28.260 | at Sinai |  |  |  |
| 2:28.261 | and the masters of |  |  |  |
| 2:28.262 | Kabbalah |  |  | הקבלה: "the Kabbalah", the Kabbalists' lore; the author claims the same word for the rabbinic chain. |
| 2:28.263 | say |  |  |  |
| 2:28.264 | thus |  |  |  |
| 2:28.265 | they received |  |  | קיבלו: "they received", from the root of Kabbalah. |
| 2:28.266 | from the mouth of |  |  |  |
| 2:28.267 | Moses |  |  |  |
| 2:28.268 | and who |  |  |  |
| 2:28.269 | is he |  |  |  |
| 2:28.270 | this one |  |  |  |
| 2:28.271 | and where |  |  |  |
| 2:28.272 | is he |  |  |  |
| 2:28.273 | who gave |  |  |  |
| 2:28.274 | you |  |  |  |
| 2:28.275 | permission |  |  |  |
| 2:28.276 | to inquire |  |  |  |
| 2:28.277 | and to investigate |  |  |  |
| 2:28.278 | after |  |  |  |
| 2:28.279 | their words |  |  |  |
| 2:28.280 | and as for what |  |  |  |
| 2:28.281 | you wrote |  |  |  |
| 2:28.282 | that they |  |  |  |
| 2:28.283 | call |  |  |  |
| 2:28.284 | the partzufim |  |  |  |
| 2:28.285 | divinities |  |  |  |
| 2:28.286 | God forbid |  |  | ח"ו: "God forbid". |
| 2:28.287 | that they |  |  |  |
| 2:28.288 | say |  |  |  |
| 2:28.289 | thus |  |  |  |
| 2:28.290 | and as for what |  |  |  |
| 2:28.291 | you wrote |  |  |  |
| 2:28.292 | go out |  |  |  |
| 2:28.293 | and learn |  |  |  |
| 2:28.294 | from the intentions of |  |  |  |
| 2:28.295 | the blessings |  |  |  |
| 2:28.296 | and the prayers |  |  |  |
| 2:28.297 | already |  |  |  |
| 2:28.298 | they explained |  |  |  |
| 2:28.299 | that the cause of |  |  |  |
| 2:28.300 | all |  |  |  |
| 2:28.301 | the causes |  |  |  |
| 2:28.302 | is designated |  |  | נתכנה: "is designated", "is called". |
| 2:28.303 | in a sefirah |  |  |  |
| 2:28.304 | a certain |  |  |  |
| 2:28.305 | by the name |  |  |  |
| 2:28.306 | such |  |  |  |
| 2:28.307 | and in a sefirah |  |  |  |
| 2:28.308 | a certain |  |  |  |
| 2:28.309 | by the name |  |  |  |
| 2:28.310 | that |  |  |  |
| 2:28.311 | and as for what |  |  |  |
| 2:28.312 | you wrote |  |  |  |
| 2:28.313 | that they explained |  |  |  |
| 2:28.314 | the Kabbalists |  |  |  |
| 2:28.315 | that several |  |  |  |
| 2:28.316 | worlds |  |  |  |
| 2:28.317 | there are |  |  |  |
| 2:28.318 | thousands of |  |  |  |
| 2:28.319 | thousands |  |  |  |
| 2:28.320 | etc. |  |  |  |
| 2:28.321 | even if |  |  |  |
| 2:28.322 | they were |  |  |  |
| 2:28.323 | millions of |  |  |  |
| 2:28.324 | millions |  |  |  |
| 2:28.325 | since |  |  |  |
| 2:28.326 | that they |  |  |  |
| 2:28.327 | say |  |  |  |
| 2:28.328 | that their Creator |  |  |  |
| 2:28.329 | is one |  |  |  |
| 2:28.330 | and He |  |  |  |
| 2:28.331 | gives life to |  |  |  |
| 2:28.332 | (object marker) |  |  |  |
| 2:28.333 | all of them |  |  |  |
| 2:28.334 | what |  |  |  |
| 2:28.335 | is there |  |  |  |
| 2:28.336 | loss |  |  |  |
| 2:28.337 | from this |  |  |  |
| 2:28.338 | and as for what |  |  |  |
| 2:28.339 | you wrote |  |  |  |
| 2:28.340 | that they |  |  |  |
| 2:28.341 | in the image of |  |  |  |
| 2:28.342 | a man |  |  |  |
| 2:28.343 | of |  |  |  |
| 2:28.344 | stature |  |  |  |
| 2:28.345 | already |  |  |  |
| 2:28.346 | they said |  |  |  |
| 2:28.347 | great |  |  |  |
| 2:28.348 | is their power |  |  |  |
| 2:28.349 | of |  |  |  |
| 2:28.350 | the prophets |  |  |  |
| 2:28.351 | who liken |  |  |  |
| 2:28.352 | the form |  |  |  |
| 2:28.353 | of might |  |  |  |
| 2:28.354 | of |  |  |  |
| 2:28.355 | above |  |  |  |
| 2:28.356 | to the shape of |  |  |  |
| 2:28.357 | a man |  |  |  |
| 2:28.358 | and as for what |  |  |  |
| 2:28.359 | you wrote |  |  |  |
| 2:28.360 | about 248 |  |  | ברמ"ח: the letters רמ"ח are 248 in letter numerals, the traditional number of the limbs. |
| 2:28.361 | limbs |  |  |  |
| 2:28.362 | already |  |  |  |
| 2:28.363 | they explained |  |  |  |
| 2:28.364 | what |  |  |  |
| 2:28.365 | they are |  |  |  |
| 2:28.366 | the limbs |  |  |  |
| 2:28.367 | for |  |  |  |
| 2:28.368 | they are |  |  |  |
| 2:28.369 | 216 |  |  | רי"ו: 216 in letter numerals. |
| 2:28.370 | letters |  |  |  |
| 2:28.371 | of |  |  |  |
| 2:28.372 | the Name |  |  |  |
| 2:28.373 | of |  |  |  |
| 2:28.374 | 72 |  |  | ע"ב: 72 in letter numerals; the Name of 72 is AB in the gematria names. |
| 2:28.375 | and 32 |  |  | ול"ב: 32 in letter numerals. |
| 2:28.376 | paths |  |  |  |
| 2:28.377 | which are |  |  |  |
| 2:28.378 | they allude to |  |  |  |
| 2:28.379 | the powers of |  |  |  |
| 2:28.380 | divine |  |  |  |
| 2:28.381 | and as for what |  |  |  |
| 2:28.382 | you wrote |  |  |  |
| 2:28.383 | that already |  |  |  |
| 2:28.384 | I conceded |  |  |  |
| 2:28.385 | that my opinion |  |  |  |
| 2:28.386 | that there is not |  |  |  |
| 2:28.387 | any |  |  |  |
| 2:28.388 | prayer |  |  |  |
| 2:28.389 | applicable |  |  |  |
| 2:28.390 | to the cause of |  |  |  |
| 2:28.391 | all |  |  |  |
| 2:28.392 | causes |  |  |  |
| 2:28.393 | falsehood |  |  |  |
| 2:28.394 | you |  |  |  |
| 2:28.395 | speak |  |  |  |
| 2:28.396 | but surely |  |  |  |
| 2:28.397 | from the beginning of |  |  |  |
| 2:28.398 | your question |  |  |  |
| 2:28.399 | to me |  |  |  |
| 2:28.400 | about |  |  |  |
| 2:28.401 | the words of |  |  |  |
| 2:28.402 | the Mesharef |  |  | המשרף: "the one who burns"; the reply does not name him. |
| 2:28.403 | I answered you |  |  |  |
| 2:28.404 | see |  |  |  |
| 2:28.405 | in Kisse Eliyahu |  |  | בכס"א: Kisse Eliyahu, "The Throne of Elijah", a book title; the meaning is given at its first occurrence in the unit. |
| 2:28.406 | but |  |  |  |
| 2:28.407 | our Kabbalah |  |  |  |
| 2:28.408 | that prayer |  |  |  |
| 2:28.409 | to Ein Sof |  |  | לא"ס: "to Ein Sof"; Ein Sof is literally "without end". |
| 2:28.410 | and to bestow |  |  |  |
| 2:28.411 | upon |  |  |  |
| 2:28.412 | ten |  |  |  |
| 2:28.413 | sefirot |  |  |  |
| 2:28.414 | as |  |  |  |
| 2:28.415 | he wrote |  |  |  |
| 2:28.416 | Lechem |  |  | לחם: Lechem Shelomo, a book title, literally "the bread of Solomon". |
| 2:28.417 | Shelomo |  |  |  |
| 2:28.418 | and as for what |  |  |  |
| 2:28.419 | I said |  |  |  |
| 2:28.420 | to you |  |  |  |
| 2:28.421 | to consult |  |  |  |
| 2:28.422 | in the Throne |  |  |  |
| 2:28.423 | of Elijah |  |  |  |
| 2:28.424 | so that |  |  |  |
| 2:28.425 | you would know |  |  |  |
| 2:28.426 | that the words of |  |  |  |
| 2:28.427 | the Mesharef |  |  |  |
| 2:28.428 | are the words of |  |  |  |
| 2:28.429 | heresy |  |  |  |
| 2:28.430 | in truth |  |  |  |
| 2:28.431 | and fitting |  |  |  |
| 2:28.432 | to root it out |  |  |  |
| 2:28.433 | and God forbid |  |  |  |
| 2:28.434 | that there be found |  |  |  |
| 2:28.435 | in any |  |  |  |
| 2:28.436 | book |  |  |  |
| 2:28.437 | like his words |  |  |  |
| 2:28.438 | but |  |  |  |
| 2:28.439 | the words of |  |  |  |
| 2:28.440 | Kisse Eliyahu |  |  |  |
| 2:28.441 | are not |  |  |  |
| 2:28.442 | like his words |  |  |  |
| 2:28.443 | God forbid |  |  |  |
| 2:28.444 | for |  |  |  |
| 2:28.445 | the words of |  |  |  |
| 2:28.446 | the Throne |  |  |  |
| 2:28.447 | of Elijah |  |  |  |
| 2:28.448 | is nothing |  |  |  |
| 2:28.449 | Ze'ir |  |  | הזעיר: "Ze'ir", from Ze'ir Anpin, literally "Small Face" (Aramaic). |
| 2:28.450 | Anpin |  |  |  |
| 2:28.451 | but |  |  |  |
| 2:28.452 | only |  |  |  |
| 2:28.453 | a vessel |  |  | כלי: "vessel", the Kabbalistic term. |
| 2:28.454 | for the action of |  |  |  |
| 2:28.455 | the Creator |  |  |  |
| 2:28.456 | and as for what |  |  |  |
| 2:28.457 | he wrote |  |  |  |
| 2:28.458 | on folio |  |  |  |
| 2:28.459 | 25 |  |  | כ"ה: 25, folio 25 (דף כ"ה). |
| 2:28.460 | that when we |  |  |  |
| 2:28.461 | call Ze'ir Anpin |  |  |  |
| 2:28.462 | we |  |  |  |
| 2:28.463 | call |  |  |  |
| 2:28.464 | God |  |  | אל: the bare name El, here of the partzufim; rendered "God" (the alternative is "El"). |
| 2:28.465 | all |  |  |  |
| 2:28.466 | the partzufim |  |  |  |
| 2:28.467 | with |  |  |  |
| 2:28.468 | a soul |  |  |  |
| 2:28.469 | hidden |  |  |  |
| 2:28.470 | in them |  |  |  |
| 2:28.471 | which is |  |  |  |
| 2:28.472 | Ein |  |  |  |
| 2:28.473 | Sof |  |  |  |
| 2:28.474 | blessed be He |  |  | ב"ה: "blessed be He". |
| 2:28.475 | this does not |  |  |  |
| 2:28.476 | mean |  |  |  |
| 2:28.477 | to say |  |  |  |
| 2:28.478 | that not |  |  |  |
| 2:28.479 | prayer |  |  |  |
| 2:28.480 | to Ein Sof |  |  |  |
| 2:28.481 | Heaven forbid |  |  |  |
| 2:28.482 | but |  |  |  |
| 2:28.483 | it means |  |  |  |
| 2:28.484 | to say |  |  |  |
| 2:28.485 | the ascent of |  |  |  |
| 2:28.486 | prayer |  |  |  |
| 2:28.487 | to |  |  |  |
| 2:28.488 | the Ein Sof |  |  |  |
| 2:28.489 | blessed be He |  |  |  |
| 2:28.490 | by way of |  |  |  |
| 2:28.491 | this |  |  |  |
| 2:28.492 | attribute |  |  |  |
| 2:28.493 | and the proof |  |  |  |
| 2:28.494 | for |  |  |  |
| 2:28.495 | this |  |  |  |
| 2:28.496 | is what |  |  |  |
| 2:28.497 | he wrote |  |  |  |
| 2:28.498 | on folio |  |  |  |
| 2:28.499 | 17 |  |  |  |
| 2:28.500 | b |  |  | ע"ב: "b", the second side of the folio (17b). |
| 2:28.501 | and these are his words: |  |  | וז"ל: "and these are his words". |
| 2:28.502 | as for what |  |  |  |
| 2:28.503 | is written |  |  |  |
| 2:28.504 | about the intentions of |  |  |  |
| 2:28.505 | prayer |  |  |  |
| 2:28.506 | and the blessings |  |  |  |
| 2:28.507 | to direct |  |  |  |
| 2:28.508 | in each |  |  |  |
| 2:28.509 | blessing |  |  |  |
| 2:28.510 | an intention |  |  |  |
| 2:28.511 | particular |  |  |  |
| 2:28.512 | to a sefirah |  |  |  |
| 2:28.513 | particular |  |  |  |
| 2:28.514 | not |  |  |  |
| 2:28.515 | God forbid |  |  |  |
| 2:28.516 | to the essence of |  |  |  |
| 2:28.517 | the sefirah |  |  |  |
| 2:28.518 | for |  |  |  |
| 2:28.519 | it |  |  |  |
| 2:28.520 | a cutting off |  |  |  |
| 2:28.521 | God forbid |  |  |  |
| 2:28.522 | and Heaven forbid |  |  |  |
| 2:28.523 | to direct |  |  |  |
| 2:28.524 | to any |  |  |  |
| 2:28.525 | power |  |  |  |
| 2:28.526 | particular |  |  |  |
| 2:28.527 | and sefirah |  |  |  |
| 2:28.528 | particular |  |  |  |
| 2:28.529 | but |  |  |  |
| 2:28.530 | all |  |  |  |
| 2:28.531 | to the essence of |  |  |  |
| 2:28.532 | the one |  |  |  |
| 2:28.533 | Ein Sof |  |  |  |
| 2:28.534 | blessed be He |  |  |  |
| 2:28.535 | comprising |  |  |  |
| 2:28.536 | all |  |  |  |
| 2:28.537 | the powers |  |  |  |
| 2:28.538 | all of them |  |  |  |
| 2:28.539 | and furthermore |  |  |  |
| 2:28.540 | he wrote |  |  |  |
| 2:28.541 | on folio |  |  |  |
| 2:28.542 | 25 |  |  |  |
| 2:28.543 | a |  |  |  |
| 2:28.544 | and all |  |  |  |
| 2:28.545 | our prayers |  |  |  |
| 2:28.546 | are |  |  |  |
| 2:28.547 | to Ein |  |  |  |
| 2:28.548 | Sof |  |  |  |
| 2:28.549 | may He be blessed |  |  |  |
| 2:28.550 | but |  |  |  |
| 2:28.551 | that not |  |  |  |
| 2:28.552 | we |  |  |  |
| 2:28.553 | can |  |  |  |
| 2:28.554 | call Him |  |  |  |
| 2:28.555 | by any |  |  |  |
| 2:28.556 | appellation |  |  |  |
| 2:28.557 | and name |  |  |  |
| 2:28.558 | therefore |  |  |  |
| 2:28.559 | we |  |  |  |
| 2:28.560 | pray |  |  |  |
| 2:28.561 | by means of |  |  |  |
| 2:28.562 | the hand of |  |  |  |
| 2:28.563 | the sefirot |  |  |  |
| 2:28.564 | for |  |  |  |
| 2:28.565 | the name |  |  |  |
| 2:28.566 | pertaining |  |  |  |
| 2:28.567 | all |  |  |  |
| 2:28.568 | all |  |  | כל: the doubled "כל כל" is as printed. |
| 2:28.569 | the appellations |  |  |  |
| 2:28.570 | and names |  |  |  |
| 2:28.571 | for |  |  |  |
| 2:28.572 | all |  |  |  |
| 2:28.573 | the appellations |  |  |  |
| 2:28.574 | and the names |  |  |  |
| 2:28.575 | are |  |  |  |
| 2:28.576 | names |  |  |  |
| 2:28.577 | of His essence |  |  |  |
| 2:28.578 | that extends |  |  |  |
| 2:28.579 | through the sefirot |  |  |  |
| 2:28.580 | also |  |  |  |
| 2:28.581 | he wrote |  |  |  |
| 2:28.582 | on folio |  |  |  |
| 2:28.583 | 26 |  |  |  |
| 2:28.584 | if |  |  |  |
| 2:28.585 | not |  |  |  |
| 2:28.586 | they do |  |  |  |
| 2:28.587 | so |  |  |  |
| 2:28.588 | and not |  |  |  |
| 2:28.589 | unify |  |  |  |
| 2:28.590 | their prayer |  |  |  |
| 2:28.591 | but |  |  |  |
| 2:28.592 | to Ze'ir Anpin |  |  |  |
| 2:28.593 | not |  |  |  |
| 2:28.594 | they will be |  |  |  |
| 2:28.595 | answered |  |  |  |
| 2:28.596 | and all |  |  |  |
| 2:28.597 | the Kabbalists |  |  |  |
| 2:28.598 | all hold |  |  | כיילי: Aramaic, "all of them hold" (see grammar). |
| 2:28.599 | the principle |  |  |  |
| 2:28.600 | this |  |  |  |
| 2:28.601 | and these are their words: |  |  |  |
| 2:28.602 | as for whatever |  |  |  |
| 2:28.603 | you find |  |  |  |
| 2:28.604 | in some |  |  |  |
| 2:28.605 | book |  |  |  |
| 2:28.606 | of theirs |  |  |  |
| 2:28.607 | prayer |  |  |  |
| 2:28.608 | to any |  |  |  |
| 2:28.609 | sefirah |  |  |  |
| 2:28.610 | Heaven forbid |  |  |  |
| 2:28.611 | that it be |  |  |  |
| 2:28.612 | so |  |  |  |
| 2:28.613 | but |  |  |  |
| 2:28.614 | the intention |  |  |  |
| 2:28.615 | that he direct |  |  |  |
| 2:28.616 | to raise them |  |  |  |
| 2:28.617 | to |  |  |  |
| 2:28.618 | Ein Sof |  |  |  |
| 2:28.619 | by way of |  |  |  |
| 2:28.620 | that |  |  |  |
| 2:28.621 | sefirah |  |  |  |

Phrases:
- 2:28.1 2:28.2 2:28.3 2:28.4 2:28.5 2:28.6 2:28.7 | Reply. To man the arrangements of the heart, and from the LORD the answer of the tongue. | 'Reply to-man the-arrangements-of the-heart and-from-the-LORD the-answer-of the-tongue'
- 2:28.8 2:28.9 2:28.10 2:28.11 | And it was when I saw your letter, | 'and-it-was when-I-saw (object-marker) your-letter'
- 2:28.12 2:28.13 2:28.14 | I sat astonished and dismayed. | 'I-sat astonished and-dismayed'
- 2:28.15 2:28.16 2:28.17 2:28.18 2:28.19 | How your heart came to inquire and investigate | 'how has-filled your-heart to-inquire and-to-investigate'
- 2:28.20 2:28.21 2:28.22 | after the band of the Geonim, | 'after band of-the-Geonim'
- 2:28.23 2:28.24 2:28.25 2:28.26 2:28.27 2:28.28 2:28.29 2:28.30 | who are faithful in the Torah, distinguished in lineage, like the seed of the mighty, | 'who they-are in-the-Torah faithful distinguished-of lineage like-seed-of the-mighty-ones'
- 2:28.31 2:28.32 2:28.33 2:28.34 2:28.35 | and after several decisors and also expounders, | 'and-after several decisors and-also expounders'
- 2:28.36 2:28.37 2:28.38 | the early and the later, | 'the-early with the-later'
- 2:28.39 2:28.40 2:28.41 2:28.42 2:28.43 | the sages of the West, Ashkenaz, Sefarad and the Yemenites. | 'sages-of the-West and-Ashkenaz and-Sefarad and-the-Yemenites'
- 2:28.44 2:28.45 2:28.46 2:28.47 2:28.48 2:28.49 | and the rest of the congregations, the holy seed, the faithful children, | 'and-the-rest-of congregations seed holy children-of faithful'
- 2:28.50 2:28.51 2:28.52 2:28.53 2:28.54 2:28.55 | this is their treasure, more than gold and abundant pearls. | 'this is their-treasure more-than-gold and-abundant pearls'
- 2:28.56 2:28.57 2:28.58 2:28.59 2:28.60 2:28.61 | Can a stumbling come from the hand of righteous and faithful ones, | 'can-it-be-that comes a-stumbling from-the-hand-of righteous-ones faithful-ones'
- 2:28.62 2:28.63 2:28.64 2:28.65 2:28.66 2:28.67 2:28.68 2:28.69 2:28.70 | who know knowledge, the wise and discerning, who fear God and turn aside from evil with their might. | 'who-know knowledge the-wise and-discerning who-fear God and-who-turn-aside from-evil with-their-might'
- 2:28.71 2:28.72 2:28.73 2:28.74 2:28.75 2:28.76 2:28.77 2:28.78 2:28.79 | By their wisdom they felled many foes and multitudes who stood against them in every age. | 'by-their-wisdom they-felled many foes and-multitudes who-stood against-them in-every age'
- 2:28.80 2:28.81 2:28.82 2:28.83 2:28.84 | Can it be that their Kabbalah is compounded in vanity and error, | 'can-it-be-that their-Kabbalah compounded in-vanity and-error'
- 2:28.85 2:28.86 2:28.87 2:28.88 2:28.89 2:28.90 2:28.91 | which, because of their great wisdom, were approved by the mouths of princes and nobles | 'which because-of-the-great their-wisdom were-approved by-the-mouths-of princes and-nobles'
- 2:28.92 2:28.93 2:28.94 2:28.95 | and stood before kings and rulers, | 'they-stood before kings and-rulers'
- 2:28.96 2:28.97 2:28.98 2:28.99 2:28.100 2:28.101 2:28.102 | and can fathers and sons walk on a path that is not good? | 'can-it-be-that they-walk on-the-path not good fathers and-sons'
- 2:28.103 2:28.104 2:28.105 2:28.106 | How many hundreds of years is this? | 'this how-many hundreds of-years'
- 2:28.107 2:28.108 2:28.109 2:28.110 2:28.111 2:28.112 | And how many of them are distinguished in every wisdom, and scholars, | 'and-how-many they are-distinguished in-every wisdom and-scholars'
- 2:28.113 2:28.114 2:28.115 2:28.116 | Kabbalists with philosophy and logic, | 'Kabbalists with philosophy and-logic'
- 2:28.117 2:28.118 2:28.119 | medicine, tishburet and computations. | 'medicine and-tishburet and-computations'
- 2:28.120 2:28.121 2:28.122 2:28.123 2:28.124 | meditating day and night, quiet and at ease, | 'meditating by-day and-by-night quiet and-at-ease'
- 2:28.125 2:28.126 2:28.127 2:28.128 2:28.129 2:28.130 | eating the fruit of their deeds, the good, fat and lush. | 'eating the-fruit-of their-deeds the-good fat and-lush'
- 2:28.131 2:28.132 2:28.133 2:28.134 | Heaven forbid! This is not to be believed | 'Heaven-forbid this is-not to-be-believed'
- 2:28.135 2:28.136 2:28.137 2:28.138 | and does not come to mind | 'and-not come-up upon the-heart'
- 2:28.139 2:28.140 2:28.141 | but only for the simple, | 'but only for-the-simple'
- 2:28.142 2:28.143 2:28.144 2:28.145 2:28.146 2:28.147 2:28.148 | who believe every matter, who lean on their understanding. | 'to-every matter who-believe who upon their-understanding lean'
- 2:28.149 2:28.150 2:28.151 2:28.152 2:28.153 | And properly, it is not incumbent on me to answer you. | 'and-properly what-is-fitting there-is-not upon-me to-answer-you'
- 2:28.154 2:28.155 2:28.156 | For what am I, | 'for what am-I'
- 2:28.157 2:28.158 2:28.159 | to enter before and within, | 'to-enter before and-within'
- 2:28.160 2:28.161 2:28.162 2:28.163 | into what I was not permitted to ponder, | 'what that-not I-was-permitted to-ponder'
- 2:28.164 2:28.165 2:28.166 2:28.167 | all the more so to answer questioners. | 'all the-more-so to-answer questioners'
- 2:28.168 2:28.169 2:28.170 2:28.171 2:28.172 | But because it is a time to act for the LORD, | 'but because-of a-time to-act for-the-LORD'
- 2:28.173 2:28.174 2:28.175 2:28.176 2:28.177 2:28.178 2:28.179 | I answer you from what I found written in the words of our Sages of blessed memory. | 'I answer-you from-what I-found written in-the-words-of our-Sages-of-blessed-memory'
- 2:28.180 2:28.181 2:28.182 2:28.183 2:28.184 2:28.185 2:28.186 2:28.187 2:28.188 | May their merit stand for me, that no stumbling befall through me. | 'and-their-merit will-stand for-me that-not will-befall any stumbling through my-hand'
- 2:28.189 2:28.190 2:28.191 2:28.192 2:28.193 2:28.194 2:28.195 | Now, as for what you wrote, that you do not despise the sages of Israel, | 'behold as-for-what you-wrote that-you-do-not despise the-sages-of Israel'
- 2:28.196 2:28.197 2:28.198 2:28.199 2:28.200 | there are already several witnesses against you; | 'already there-are against-you several witnesses'
- 2:28.201 2:28.202 2:28.203 | you cannot deny it. | 'not you-can to-deny'
- 2:28.204 2:28.205 2:28.206 2:28.207 2:28.208 | And what brought you to this | 'and-what brought you to this'
- 2:28.209 2:28.210 2:28.211 2:28.212 2:28.213 | that you attribute to them error and heresy. | 'that-you attribute to-them error and-heresy'
- 2:28.214 2:28.215 2:28.216 2:28.217 2:28.218 | Rather, because they are despised in your eyes. | 'but because that-they-are despised in-your-eyes'
- 2:28.219 2:28.220 2:28.221 2:28.222 2:28.223 2:28.224 2:28.225 | But a person should take heed for himself, lest he be burned by their coals, | 'but the-person should-take-heed for-himself lest he-be-burned by-their-coals'
- 2:28.226 2:28.227 2:28.228 2:28.229 2:28.230 2:28.231 | for the righteous are greater in their death than in their life, | 'for great-ones righteous-ones in-their-death more than-in-their-life'
- 2:28.232 2:28.233 2:28.234 | as they said in the Talmud. | 'as they-said in-the-Talmud'
- 2:28.235 2:28.236 2:28.237 2:28.238 2:28.239 2:28.240 | And as for what you wrote, that if one reveals the place of contradiction, | 'and-as-for-what you-wrote that-if to-reveal the-place-of the-contradiction'
- 2:28.241 2:28.242 2:28.243 2:28.244 2:28.245 2:28.246 2:28.247 2:28.248 | and the Talmud also brings proof to refute the words of those greater than them, | 'and-also the-Talmud brings proof to-refute the-words-of greater than-them'
- 2:28.249 2:28.250 2:28.251 2:28.252 | the matter is not comparable to a proof, | 'there-is-not the-matter comparable to-a-proof'
- 2:28.253 2:28.254 2:28.255 2:28.256 2:28.257 2:28.258 2:28.259 2:28.260 | for where do we find that they bring proof to refute a halakhah given to Moses at Sinai? | 'for-where do-we-find that-they-bring proof to-refute a-halakhah given-to-Moses at-Sinai'
- 2:28.261 2:28.262 2:28.263 2:28.264 2:28.265 2:28.266 2:28.267 | And the masters of Kabbalah say that they received thus from the mouth of Moses, | 'and-the-masters-of Kabbalah say thus they-received from-the-mouth-of Moses'
- 2:28.268 2:28.269 2:28.270 2:28.271 2:28.272 2:28.273 2:28.274 2:28.275 2:28.276 2:28.277 2:28.278 2:28.279 | and who is this one, and where is he, who gave you permission to inquire and investigate after their words? | 'and-who is-he this-one and-where is-he who-gave you permission to-inquire and-to-investigate after their-words'
- 2:28.280 2:28.281 2:28.282 2:28.283 2:28.284 2:28.285 | And as for what you wrote, that they call the partzufim divinities, | 'and-as-for-what you-wrote that-they call the-partzufim divinities'
- 2:28.286 2:28.287 2:28.288 2:28.289 | God forbid, that they say thus. | 'God-forbid that-they say thus'
- 2:28.290 2:28.291 2:28.292 2:28.293 2:28.294 2:28.295 2:28.296 | And as for what you wrote, go and learn from the intentions of the blessings and the prayers. | 'and-as-for-what you-wrote go-out and-learn from-the-intentions-of the-blessings and-the-prayers'
- 2:28.297 2:28.298 2:28.299 2:28.300 2:28.301 2:28.302 | They already explained that the cause of all causes is designated | 'already they-explained that-the-cause-of all the-causes is-designated'
- 2:28.303 2:28.304 2:28.305 2:28.306 | in a certain sefirah by such a name. | 'in-a-sefirah a-certain by-the-name such'
- 2:28.307 2:28.308 2:28.309 2:28.310 | And in a certain sefirah by that name. | 'and-in-a-sefirah a-certain by-the-name that'
- 2:28.311 2:28.312 2:28.313 2:28.314 2:28.315 2:28.316 2:28.317 2:28.318 2:28.319 2:28.320 | And as for what you wrote, that the Kabbalists explained that there are several worlds, thousands of thousands, etc., | 'and-as-for-what you-wrote that-they-explained the-Kabbalists that-several worlds there-are thousands-of thousands etc.'
- 2:28.321 2:28.322 2:28.323 2:28.324 | even if they were millions of millions, | 'even-if they-were millions-of millions'
- 2:28.325 2:28.326 2:28.327 2:28.328 2:28.329 2:28.330 2:28.331 2:28.332 2:28.333 2:28.334 2:28.335 2:28.336 2:28.337 | since they say that their Creator is one, and He gives life to all of them: what loss is there from this? | 'since that-they say that-their-Creator is-one and-He gives-life-to (object-marker) all-of-them what is-there loss from-this'
- 2:28.338 2:28.339 2:28.340 2:28.341 2:28.342 2:28.343 2:28.344 | And as for what you wrote, that they are in the image of a man of stature, | 'and-as-for-what you-wrote that-they in-the-image-of a-man of stature'
- 2:28.345 2:28.346 2:28.347 2:28.348 2:28.349 2:28.350 2:28.351 2:28.352 2:28.353 2:28.354 2:28.355 2:28.356 2:28.357 | already they said, "Great is the power of the prophets, who liken the form of might above to the shape of a man." | 'already they-said great is-their-power of the-prophets who-liken the-form of-might of above to-the-shape-of a-man'
- 2:28.358 2:28.359 2:28.360 2:28.361 | And as for what you wrote about 248 limbs, | 'and-as-for-what you-wrote about-248 limbs'
- 2:28.362 2:28.363 2:28.364 2:28.365 2:28.366 | they already explained what the limbs are. | 'already they-explained what they-are the-limbs'
- 2:28.367 2:28.368 2:28.369 2:28.370 2:28.371 2:28.372 | For they are 216 letters of the Name | 'for they-are 216 letters of the-Name'
- 2:28.373 2:28.374 2:28.375 2:28.376 | of 72 and 32 paths, | 'of 72 and-32 paths'
- 2:28.377 2:28.378 2:28.379 2:28.380 | which allude to divine powers. | 'which-are they-allude-to the-powers-of divine'
- 2:28.381 2:28.382 2:28.383 2:28.384 2:28.385 | And as for what you wrote, that I already conceded that my opinion | 'and-as-for-what you-wrote that-already I-conceded that-my-opinion'
- 2:28.386 2:28.387 2:28.388 2:28.389 2:28.390 2:28.391 2:28.392 | is that there is no prayer applicable to the Cause of all causes, | 'that-there-is-not any prayer applicable to-the-cause-of all causes'
- 2:28.393 2:28.394 2:28.395 | you speak falsehood. | 'falsehood you speak'
- 2:28.396 2:28.397 2:28.398 2:28.399 2:28.400 2:28.401 2:28.402 | But surely, from the beginning of your question to me about the words of the Mesharef, | 'but-surely from-the-beginning-of your-question to-me about the-words-of the-Mesharef'
- 2:28.403 2:28.404 2:28.405 | I answered you, see in Kisse Eliyahu. | 'I-answered-you see in-Kisse-Eliyahu'
- 2:28.406 2:28.407 2:28.408 2:28.409 2:28.410 2:28.411 2:28.412 2:28.413 | But our Kabbalah holds that prayer is to Ein Sof and bestows upon the ten sefirot, | 'but our-Kabbalah that-prayer to-Ein-Sof and-to-bestow upon ten sefirot'
- 2:28.414 2:28.415 2:28.416 2:28.417 | as Lechem Shelomo wrote. | 'as he-wrote Lechem Shelomo'
- 2:28.418 2:28.419 2:28.420 2:28.421 2:28.422 2:28.423 | And as for what I told you, to consult Kisse Eliyahu, | 'and-as-for-what I-said to-you to-consult in-the-Throne of-Elijah'
- 2:28.424 2:28.425 2:28.426 2:28.427 2:28.428 2:28.429 2:28.430 2:28.431 2:28.432 | so that you would know that the words of the Mesharef are words of heresy, in truth, and fitting to root it out. | 'so-that you-would-know that-the-words-of the-Mesharef are-the-words-of heresy in-truth and-fitting to-root-it-out'
- 2:28.433 2:28.434 2:28.435 2:28.436 2:28.437 | And God forbid that there be found in any book words like his, | 'and-God-forbid that-there-be-found in-any book like-his-words'
- 2:28.438 2:28.439 2:28.440 2:28.441 2:28.442 2:28.443 | but the words of Kisse Eliyahu are not like his, God forbid. | 'but the-words-of Kisse-Eliyahu are-not like-his-words God-forbid'
- 2:28.444 2:28.445 2:28.446 2:28.447 | For the words of Kisse Eliyahu | 'for the-words-of the-Throne of-Elijah'
- 2:28.448 2:28.449 2:28.450 2:28.451 2:28.452 2:28.453 2:28.454 2:28.455 | hold that Ze'ir Anpin is nothing but a vessel for the action of the Creator. | 'is-nothing Ze'ir Anpin but only a-vessel for-the-action-of the-Creator'
- 2:28.456 2:28.457 2:28.458 2:28.459 | And as for what he wrote on folio 25, | 'and-as-for-what he-wrote on-folio 25'
- 2:28.460 2:28.461 2:28.462 2:28.463 2:28.464 2:28.465 2:28.466 2:28.467 2:28.468 2:28.469 2:28.470 | that when we call Ze'ir Anpin God, we call all the partzufim God, with a soul hidden in them, | 'that-when-we call-Ze'ir-Anpin we call God all the-partzufim with a-soul hidden in-them'
- 2:28.471 2:28.472 2:28.473 2:28.474 | which is Ein Sof, blessed be He. | 'which-is Ein Sof blessed-be-He'
- 2:28.475 2:28.476 2:28.477 | This does not mean | 'this-does-not mean to-say'
- 2:28.478 2:28.479 2:28.480 2:28.481 | that prayer is not directed to Ein Sof, Heaven forbid, | 'that-not prayer to-Ein-Sof Heaven-forbid'
- 2:28.482 2:28.483 2:28.484 | but it means that | 'but it-means to-say'
- 2:28.485 2:28.486 2:28.487 2:28.488 2:28.489 2:28.490 2:28.491 2:28.492 | the ascent of prayer to Ein Sof, blessed be He, is by way of this attribute. | 'the-ascent-of prayer to the-Ein-Sof blessed-be-He by-way-of this attribute'
- 2:28.493 2:28.494 2:28.495 2:28.496 2:28.497 | And the proof for this is what he wrote | 'and-the-proof for this is-what he-wrote'
- 2:28.498 2:28.499 2:28.500 2:28.501 | on folio 17b, and these are his words: | 'on-folio 17 b and-these-are-his-words:'
- 2:28.502 2:28.503 2:28.504 2:28.505 2:28.506 2:28.507 2:28.508 2:28.509 2:28.510 2:28.511 | "As for what is written about the intentions of prayer and the blessings, to direct in each blessing a particular intention, | 'as-for-what is-written about-the-intentions-of prayer and-the-blessings to-direct in-each blessing an-intention particular'
- 2:28.512 2:28.513 2:28.514 2:28.515 2:28.516 2:28.517 | to a particular sefirah, God forbid, not to the essence of the sefirah. | 'to-a-sefirah particular not God-forbid to-the-essence-of the-sefirah'
- 2:28.518 2:28.519 2:28.520 2:28.521 | For it is a cutting off, God forbid, | 'for it a-cutting-off God-forbid'
- 2:28.522 2:28.523 2:28.524 2:28.525 2:28.526 2:28.527 2:28.528 | and Heaven forbid to direct it to any particular power or particular sefirah. | 'and-Heaven-forbid to-direct to-any power particular and-sefirah particular'
- 2:28.529 2:28.530 2:28.531 2:28.532 2:28.533 2:28.534 2:28.535 2:28.536 2:28.537 2:28.538 | But all is to the essence of the one Ein Sof, blessed be He, comprising all the powers, all of them." | 'but all to-the-essence-of the-one Ein-Sof blessed-be-He comprising all the-powers all-of-them'
- 2:28.539 2:28.540 2:28.541 2:28.542 2:28.543 | And furthermore he wrote on folio 25a: | 'and-furthermore he-wrote on-folio 25 a'
- 2:28.544 2:28.545 2:28.546 2:28.547 2:28.548 2:28.549 | "And all our prayers are to Ein Sof, may He be blessed. | 'and-all our-prayers are to-Ein Sof may-He-be-blessed'
- 2:28.550 2:28.551 2:28.552 2:28.553 2:28.554 2:28.555 2:28.556 2:28.557 | But that we cannot call Him by any appellation and name. | 'but that-not we can call-Him by-any appellation and-name'
- 2:28.558 2:28.559 2:28.560 2:28.561 2:28.562 2:28.563 | Therefore we pray through the sefirot. | 'therefore we pray by-means-of the-hand-of the-sefirot'
- 2:28.564 2:28.565 2:28.566 2:28.567 2:28.568 2:28.569 2:28.570 | For the name pertaining to all, all the appellations and names. | 'for the-name pertaining all all the-appellations and-names'
- 2:28.571 2:28.572 2:28.573 2:28.574 2:28.575 2:28.576 2:28.577 2:28.578 2:28.579 | For all the appellations and names are names of His essence, extending through the sefirot." | 'for all the-appellations and-the-names are names of-His-essence that-extends through-the-sefirot'
- 2:28.580 2:28.581 2:28.582 2:28.583 | Also he wrote on folio 26: | 'also he-wrote on-folio 26'
- 2:28.584 2:28.585 2:28.586 2:28.587 2:28.588 2:28.589 2:28.590 2:28.591 2:28.592 2:28.593 2:28.594 2:28.595 | "if they do not do so, and do not unify their prayer but to Ze'ir Anpin, they will not be answered." | 'if not they-do so and-not unify their-prayer but to-Ze'ir-Anpin not they-will-be answered'
- 2:28.596 2:28.597 2:28.598 2:28.599 2:28.600 | And all the Kabbalists hold this principle, | 'and-all the-Kabbalists all-hold the-principle this'
- 2:28.601 2:28.602 2:28.603 2:28.604 2:28.605 2:28.606 2:28.607 2:28.608 2:28.609 2:28.610 2:28.611 2:28.612 | and these are their words: "As for whatever you find in some book of theirs, that prayer is addressed to any sefirah, Heaven forbid that it be so. | 'and-these-are-their-words: as-for-whatever you-find in-some book of-theirs prayer to-any sefirah Heaven-forbid that-it-be so'
- 2:28.613 2:28.614 2:28.615 2:28.616 2:28.617 2:28.618 2:28.619 2:28.620 2:28.621 | But the intention is that one direct them to raise them to Ein Sof by way of that sefirah." | 'but the-intention that-he-direct to-raise-them to Ein-Sof by-way-of that sefirah'

Grammar notes:
- The paragraph is the respondent's reply in chapter 2. "You" is the questioner; the reply takes up each point of his letter in turn, introduced by ומה שכתבת ("as for what you wrote") and answered after it.
- The rhetorical questions opened by וכי ("Can it be that ...", 2:28.56, 2:28.80, 2:28.96) and the sentences ending in a question mark (2:28.102, 2:28.106, 2:28.260, 2:28.279, 2:28.337) keep their question form; the text gives no answer to them within the paragraph.
- Quotations of the Kabbalists are in double quotation marks. They open after the formula וז"ל (2:28.501, 2:28.601) or after כ' with a folio (2:28.539, 2:28.580), and close at the end of each passage, since the text gives no closing formula. The quoted words are Hebrew, except כיילי כללא דא (2:28.598-600), which is Aramaic and is rendered by its sense.
- The verb ייחדו (2:28.589) takes the prayer as its object and is rendered "unify", the Kabbalistic act; לז"א (2:28.592) is the entity the prayer is directed to.
- Two Biblical allusions carry the opening and the close of the argument: Prov 16:1 at the start (2:28.1-7), and the word לה' from Ps 119:126 (2:28.172), which the author reads as "for the LORD".

Translator's note: Prose paragraph: the respondent's reply in chapter 2. Each word is glossed. Letter numerals are given as Arabic numerals (רמ"ח 248, רי"ו 216, ע"ב 72, ל"ב 32, כ"ה 25, י"ז ע"ב 17b). Kisse Eliyahu and Lechem Shelomo are book titles; Ein Sof and Ze'ir Anpin are kept as names.

Commentary rulings:
- apparatus: Base text: Wikisource, as printed; the 1931 scan is not yet collated (Q2), so no reading is added or corrected here.
- apparatus: Prov 16:1 (allusion at the opening, 2:28.1-7): the verse reads "The arrangements of the heart belong to man, and the answer of the tongue is from the LORD." The author's Hebrew is the verse's words with תשובה prefixed.
- apparatus: Job 1:1 (allusion, 2:28.66-70): "וִירֵא אֱלֹהִים וְסָר מֵרָע", rendered "fear God and turn from evil", with the author's plural form.
- apparatus: Isa 6:13 (allusion, 2:28.46-47): "זֶרַע קֹדֶשׁ", "the holy seed".
- apparatus: Isa 65:2 (allusion, 2:28.97-100): "בְּדֶרֶךְ לֹא־טוֹב", "in a way that is not good".
- apparatus: Ps 92:15 (allusion, 2:28.129-130): "דְּשֵׁנִים וְרַעֲנַנִּים", "fat and luxuriant".
- apparatus: Avot 2:10 (allusion, 2:28.224-225): the author's "פן יכוה בגחלתן" recalls the saying to beware of the coals of the sages; the words are his, not a quotation.
- apparatus: Ps 119:126 (allusion, 2:28.171-172): "עֵת לַעֲשׂוֹת לַה'", "It is a time to act for the LORD".
- apparatus: Folio references as the author gives them: folio 17b (2:28.498-500), folio 25 (2:28.458-459, 2:28.541-543), folio 26 (2:28.582-583). No reference is added.
- apparatus: The doubled "כל כל" (2:28.567-568) is as printed; both words are translated.
- apparatus: ומונים (2:28.75): the word is uncertain; rendered "and multitudes".
- apparatus: המשרף (2:28.402, 2:28.427): the reply does not name him; the transliteration "Mesharef" is kept.
- apparatus: ותשבורת and ותוכנים (2:28.118-119): technical terms; "tishburet" is kept, and "computations" is a guess at the second.
- apparatus: The Talmudic sayings (2:28.233-234, 2:28.241-260, 2:28.258-260) are given without a reference, since the author gives none.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

