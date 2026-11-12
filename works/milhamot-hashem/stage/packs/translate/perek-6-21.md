# Translation pack: milhamot-hashem perek-6-21 (Milchamot ha-Shem 6:32–6:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-21.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-21`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
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


## Milchamot ha-Shem 6:32

Hebrew (Wikisource): עה) ודבריו אלה שכתב שנקראו אלהים אחרים לפי שאין להם קיום מעצמן רק על ידי אחרים, נובעים ממקור נאמן מדרש רבה והילקוט פרשת ויצא ופ' מקץ מדרש הגדול בויהי מקץ, וז"ל מדרש רבה עפ פי' יפה תואר ועץ יוסף, אמר ר"י הרשעים מתקיימין על אלהיהם, ופרעה חלם והנה עומד על היאור, אבל הצדיקים אלהיהם מתקיים עליהם והנה ה' נצב עליו, (פי' ההתיצבות היא ההשגחה להצילו מכל נזק ופגע וזהו שאמר ר"י הרשעים צריכים לעמוד על אלהיהם ולשמרם. אבל הצדיקים אלהיהם נצב עליהם לשמרם ולהצילם ופי' נצב משגיח ע"כ) ובמדרש הגדול גורס במקום מתקיימין עומדים. ובילקוט פר' יתרו. אלהים אחרים. וכי אלהות הן והלא כבר נאמר ונתן אלהיהם באש כי לא אלהים המה. ומה תלמוד לומר אלהים אחרים. שהם מאחרים לעובדיהם ובן הוא אומר הן יצעק אליו ולא יענה ומצרתו לא יושיעו, ד"א אלהים אחרים שהם עושים את עובדיהן אחרים. ר' יוסי אומר אלהים אחרים למה נאמר שלא ליתן פתחון פה לאומות העולם לומר אלו נקראו בשמו כבר היה בהם צורך הרי נקראו בשמו ואין בהם צורך. אימתי נקראו בשמו, בימי אנוש בן שת שנ' אז הוחל לקרא בשם ה'. באותה שעה אמר הקב"ה אתם עשיתם מעשה חדש וקראתם אותם אלוהות, אני אעשה מעשה חדש ואקרא עצמי ה', באותה שעה קלה עלה אוקיינוס והציף שלישו של עולם וכן הוא אומר הקורא למי הים וישפכם על פני הארץ ה' שמו, ר' אלעזר המודעי אומר אלהים אחרים שהם מחדשים להם אלוהות בכל יום, הא כיצד היה לו של זהב והנצרך לו עושהו של כסף. נצרך לו עושהו של נחשת, וכן ברזל, וכן בדיל ועופרת, שני חדשים מקרוב באו וכו' ע"ש. והנה מכל זה תבין שגגת רבותינו בעלי הקבלה החדשה שהאמינו לדברי המסית מחבר הזהר במציאות אלהים אחרים בעלי עשר ספירות וחמשה פרצופים דקליפה נגד עשר ספירות ופרצופין של קדושה, ואמרו נבראת לעומת זה עשה האלהים, גם היכלות הטומאה עומת היכלות הקדושה, בכל פרצוף שלה (משנת חסידים מסכת היכלות הקליפה וס' מחברת הקדש שער ר"ח ד' ט"ז ע"ב). וסוברים שכמו שאצל וברא ויצר ועשה אלהים קדושים כן אצל וברא ויצר ועשה אלהים אחרים וכמו שיש בקדושה אדם קדמון ועתיק ואריך, ואבא ואימא, וזעיר ונוקביה דקדושה כן בשמאל אדם בליעל סטרא אחרא מסאבא, זה לעומת זה כמו שיראה המעיין בס' חיים למהרח"ו, ומשנת חסידים שער הקליפות. ואם תעמיק חקר תבין ותשכיל שסביא זו באה ונמשה להם מאמונת השונים שסוברים שיש בעולם שני פועלים, אחד פועל הטוב, ואחד פועל הרע. כמו שכתבו הראשונים לדחותה בשתי ידים, והמקובלים הרדשים תפשוה וציירו אותה בפנים אחרים, ונדחקו מאד לכסות את נבלותה, ולהלבישה באמונת תוח"ק. ובאמת זרה היא עד מאד לאמונת יהוד השי"ת, ולא תבוא בקהל ה'. ולאו קבלה היא כלל. כמו שכ' לעיל מדברי רז"ל. כי המקובלים החדשים מאמתים מציאות אל אחד בעולם היפך דברי רז"ל. והכתובים מוכיחים כדעת חכמים ז"ל. שנ' כי הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד מלבדו. אתה הוא ה' לבדך. וכמה כתובים מכריזים שאין עוד אלוה אחר זולתי ה' אלהינו. ורבותינו מתקני התפלות התקינו לומר בכל יום אנו יודעים ומעידים ומגידים, שאין אלוה זולתך וכו' והרבה לשונות בתפלה מורים שאין מציאות אל אחר בעולם כלל כלל. ורק לפי מחשבת ובדיהן. אמרה כי תורה כי לא תשתחוה לאל אחר. וכן לא יהיה לך אלהים אחרים על פני. הכל לפי מחשבת עובידהן שטועין וחושבים אותם אלוהות. לא שהשי"ת המציא שני מיני אלוהות אלה קדושים ואלה טמאים ח"ו. כי אין אלוה זולתו. וכמה פסורים בתורה ובנביאים ובכתובים מוכיחים כן.

| id | word | punctuation after |
|---|---|---|
| 6:32.1 | עה | ) |
| 6:32.2 | ודבריו |  |
| 6:32.3 | אלה |  |
| 6:32.4 | שכתב |  |
| 6:32.5 | שנקראו |  |
| 6:32.6 | אלהים |  |
| 6:32.7 | אחרים |  |
| 6:32.8 | לפי |  |
| 6:32.9 | שאין |  |
| 6:32.10 | להם |  |
| 6:32.11 | קיום |  |
| 6:32.12 | מעצמן |  |
| 6:32.13 | רק |  |
| 6:32.14 | על |  |
| 6:32.15 | ידי |  |
| 6:32.16 | אחרים | , |
| 6:32.17 | נובעים |  |
| 6:32.18 | ממקור |  |
| 6:32.19 | נאמן |  |
| 6:32.20 | מדרש |  |
| 6:32.21 | רבה |  |
| 6:32.22 | והילקוט |  |
| 6:32.23 | פרשת |  |
| 6:32.24 | ויצא |  |
| 6:32.25 | ופ | ' |
| 6:32.26 | מקץ |  |
| 6:32.27 | מדרש |  |
| 6:32.28 | הגדול |  |
| 6:32.29 | בויהי |  |
| 6:32.30 | מקץ | , |
| 6:32.31 | וז"ל |  |
| 6:32.32 | מדרש |  |
| 6:32.33 | רבה |  |
| 6:32.34 | עפ |  |
| 6:32.35 | פי | ' |
| 6:32.36 | יפה |  |
| 6:32.37 | תואר |  |
| 6:32.38 | ועץ |  |
| 6:32.39 | יוסף | , |
| 6:32.40 | אמר |  |
| 6:32.41 | ר"י |  |
| 6:32.42 | הרשעים |  |
| 6:32.43 | מתקיימין |  |
| 6:32.44 | על |  |
| 6:32.45 | אלהיהם | , |
| 6:32.46 | ופרעה |  |
| 6:32.47 | חלם |  |
| 6:32.48 | והנה |  |
| 6:32.49 | עומד |  |
| 6:32.50 | על |  |
| 6:32.51 | היאור | , |
| 6:32.52 | אבל |  |
| 6:32.53 | הצדיקים |  |
| 6:32.54 | אלהיהם |  |
| 6:32.55 | מתקיים |  |
| 6:32.56 | עליהם |  |
| 6:32.57 | והנה |  |
| 6:32.58 | ה | ' |
| 6:32.59 | נצב |  |
| 6:32.60 | עליו | , |
| 6:32.61 | פי | ' |
| 6:32.62 | ההתיצבות |  |
| 6:32.63 | היא |  |
| 6:32.64 | ההשגחה |  |
| 6:32.65 | להצילו |  |
| 6:32.66 | מכל |  |
| 6:32.67 | נזק |  |
| 6:32.68 | ופגע |  |
| 6:32.69 | וזהו |  |
| 6:32.70 | שאמר |  |
| 6:32.71 | ר"י |  |
| 6:32.72 | הרשעים |  |
| 6:32.73 | צריכים |  |
| 6:32.74 | לעמוד |  |
| 6:32.75 | על |  |
| 6:32.76 | אלהיהם |  |
| 6:32.77 | ולשמרם | . |
| 6:32.78 | אבל |  |
| 6:32.79 | הצדיקים |  |
| 6:32.80 | אלהיהם |  |
| 6:32.81 | נצב |  |
| 6:32.82 | עליהם |  |
| 6:32.83 | לשמרם |  |
| 6:32.84 | ולהצילם |  |
| 6:32.85 | ופי | ' |
| 6:32.86 | נצב |  |
| 6:32.87 | משגיח |  |
| 6:32.88 | ע"כ | ) |
| 6:32.89 | ובמדרש |  |
| 6:32.90 | הגדול |  |
| 6:32.91 | גורס |  |
| 6:32.92 | במקום |  |
| 6:32.93 | מתקיימין |  |
| 6:32.94 | עומדים | . |
| 6:32.95 | ובילקוט |  |
| 6:32.96 | פר | ' |
| 6:32.97 | יתרו | . |
| 6:32.98 | אלהים |  |
| 6:32.99 | אחרים | . |
| 6:32.100 | וכי |  |
| 6:32.101 | אלהות |  |
| 6:32.102 | הן |  |
| 6:32.103 | והלא |  |
| 6:32.104 | כבר |  |
| 6:32.105 | נאמר |  |
| 6:32.106 | ונתן |  |
| 6:32.107 | אלהיהם |  |
| 6:32.108 | באש |  |
| 6:32.109 | כי |  |
| 6:32.110 | לא |  |
| 6:32.111 | אלהים |  |
| 6:32.112 | המה | . |
| 6:32.113 | ומה |  |
| 6:32.114 | תלמוד |  |
| 6:32.115 | לומר |  |
| 6:32.116 | אלהים |  |
| 6:32.117 | אחרים | . |
| 6:32.118 | שהם |  |
| 6:32.119 | מאחרים |  |
| 6:32.120 | לעובדיהם |  |
| 6:32.121 | ובן |  |
| 6:32.122 | הוא |  |
| 6:32.123 | אומר |  |
| 6:32.124 | הן |  |
| 6:32.125 | יצעק |  |
| 6:32.126 | אליו |  |
| 6:32.127 | ולא |  |
| 6:32.128 | יענה |  |
| 6:32.129 | ומצרתו |  |
| 6:32.130 | לא |  |
| 6:32.131 | יושיעו | , |
| 6:32.132 | ד"א |  |
| 6:32.133 | אלהים |  |
| 6:32.134 | אחרים |  |
| 6:32.135 | שהם |  |
| 6:32.136 | עושים |  |
| 6:32.137 | את |  |
| 6:32.138 | עובדיהן |  |
| 6:32.139 | אחרים | . |
| 6:32.140 | ר | ' |
| 6:32.141 | יוסי |  |
| 6:32.142 | אומר |  |
| 6:32.143 | אלהים |  |
| 6:32.144 | אחרים |  |
| 6:32.145 | למה |  |
| 6:32.146 | נאמר |  |
| 6:32.147 | שלא |  |
| 6:32.148 | ליתן |  |
| 6:32.149 | פתחון |  |
| 6:32.150 | פה |  |
| 6:32.151 | לאומות |  |
| 6:32.152 | העולם |  |
| 6:32.153 | לומר |  |
| 6:32.154 | אלו |  |
| 6:32.155 | נקראו |  |
| 6:32.156 | בשמו |  |
| 6:32.157 | כבר |  |
| 6:32.158 | היה |  |
| 6:32.159 | בהם |  |
| 6:32.160 | צורך |  |
| 6:32.161 | הרי |  |
| 6:32.162 | נקראו |  |
| 6:32.163 | בשמו |  |
| 6:32.164 | ואין |  |
| 6:32.165 | בהם |  |
| 6:32.166 | צורך | . |
| 6:32.167 | אימתי |  |
| 6:32.168 | נקראו |  |
| 6:32.169 | בשמו | , |
| 6:32.170 | בימי |  |
| 6:32.171 | אנוש |  |
| 6:32.172 | בן |  |
| 6:32.173 | שת |  |
| 6:32.174 | שנ | ' |
| 6:32.175 | אז |  |
| 6:32.176 | הוחל |  |
| 6:32.177 | לקרא |  |
| 6:32.178 | בשם |  |
| 6:32.179 | ה | '. |
| 6:32.180 | באותה |  |
| 6:32.181 | שעה |  |
| 6:32.182 | אמר |  |
| 6:32.183 | הקב"ה |  |
| 6:32.184 | אתם |  |
| 6:32.185 | עשיתם |  |
| 6:32.186 | מעשה |  |
| 6:32.187 | חדש |  |
| 6:32.188 | וקראתם |  |
| 6:32.189 | אותם |  |
| 6:32.190 | אלוהות | , |
| 6:32.191 | אני |  |
| 6:32.192 | אעשה |  |
| 6:32.193 | מעשה |  |
| 6:32.194 | חדש |  |
| 6:32.195 | ואקרא |  |
| 6:32.196 | עצמי |  |
| 6:32.197 | ה | ', |
| 6:32.198 | באותה |  |
| 6:32.199 | שעה |  |
| 6:32.200 | קלה |  |
| 6:32.201 | עלה |  |
| 6:32.202 | אוקיינוס |  |
| 6:32.203 | והציף |  |
| 6:32.204 | שלישו |  |
| 6:32.205 | של |  |
| 6:32.206 | עולם |  |
| 6:32.207 | וכן |  |
| 6:32.208 | הוא |  |
| 6:32.209 | אומר |  |
| 6:32.210 | הקורא |  |
| 6:32.211 | למי |  |
| 6:32.212 | הים |  |
| 6:32.213 | וישפכם |  |
| 6:32.214 | על |  |
| 6:32.215 | פני |  |
| 6:32.216 | הארץ |  |
| 6:32.217 | ה | ' |
| 6:32.218 | שמו | , |
| 6:32.219 | ר | ' |
| 6:32.220 | אלעזר |  |
| 6:32.221 | המודעי |  |
| 6:32.222 | אומר |  |
| 6:32.223 | אלהים |  |
| 6:32.224 | אחרים |  |
| 6:32.225 | שהם |  |
| 6:32.226 | מחדשים |  |
| 6:32.227 | להם |  |
| 6:32.228 | אלוהות |  |
| 6:32.229 | בכל |  |
| 6:32.230 | יום | , |
| 6:32.231 | הא |  |
| 6:32.232 | כיצד |  |
| 6:32.233 | היה |  |
| 6:32.234 | לו |  |
| 6:32.235 | של |  |
| 6:32.236 | זהב |  |
| 6:32.237 | והנצרך |  |
| 6:32.238 | לו |  |
| 6:32.239 | עושהו |  |
| 6:32.240 | של |  |
| 6:32.241 | כסף | . |
| 6:32.242 | נצרך |  |
| 6:32.243 | לו |  |
| 6:32.244 | עושהו |  |
| 6:32.245 | של |  |
| 6:32.246 | נחשת | , |
| 6:32.247 | וכן |  |
| 6:32.248 | ברזל | , |
| 6:32.249 | וכן |  |
| 6:32.250 | בדיל |  |
| 6:32.251 | ועופרת | , |
| 6:32.252 | שני |  |
| 6:32.253 | חדשים |  |
| 6:32.254 | מקרוב |  |
| 6:32.255 | באו |  |
| 6:32.256 | וכו | ' |
| 6:32.257 | ע"ש | . |
| 6:32.258 | והנה |  |
| 6:32.259 | מכל |  |
| 6:32.260 | זה |  |
| 6:32.261 | תבין |  |
| 6:32.262 | שגגת |  |
| 6:32.263 | רבותינו |  |
| 6:32.264 | בעלי |  |
| 6:32.265 | הקבלה |  |
| 6:32.266 | החדשה |  |
| 6:32.267 | שהאמינו |  |
| 6:32.268 | לדברי |  |
| 6:32.269 | המסית |  |
| 6:32.270 | מחבר |  |
| 6:32.271 | הזהר |  |
| 6:32.272 | במציאות |  |
| 6:32.273 | אלהים |  |
| 6:32.274 | אחרים |  |
| 6:32.275 | בעלי |  |
| 6:32.276 | עשר |  |
| 6:32.277 | ספירות |  |
| 6:32.278 | וחמשה |  |
| 6:32.279 | פרצופים |  |
| 6:32.280 | דקליפה |  |
| 6:32.281 | נגד |  |
| 6:32.282 | עשר |  |
| 6:32.283 | ספירות |  |
| 6:32.284 | ופרצופין |  |
| 6:32.285 | של |  |
| 6:32.286 | קדושה | , |
| 6:32.287 | ואמרו |  |
| 6:32.288 | נבראת |  |
| 6:32.289 | לעומת |  |
| 6:32.290 | זה |  |
| 6:32.291 | עשה |  |
| 6:32.292 | האלהים | , |
| 6:32.293 | גם |  |
| 6:32.294 | היכלות |  |
| 6:32.295 | הטומאה |  |
| 6:32.296 | עומת |  |
| 6:32.297 | היכלות |  |
| 6:32.298 | הקדושה | , |
| 6:32.299 | בכל |  |
| 6:32.300 | פרצוף |  |
| 6:32.301 | שלה |  |
| 6:32.302 | משנת |  |
| 6:32.303 | חסידים |  |
| 6:32.304 | מסכת |  |
| 6:32.305 | היכלות |  |
| 6:32.306 | הקליפה |  |
| 6:32.307 | וס | ' |
| 6:32.308 | מחברת |  |
| 6:32.309 | הקדש |  |
| 6:32.310 | שער |  |
| 6:32.311 | ר"ח |  |
| 6:32.312 | ד | ' |
| 6:32.313 | ט"ז |  |
| 6:32.314 | ע"ב | ). |
| 6:32.315 | וסוברים |  |
| 6:32.316 | שכמו |  |
| 6:32.317 | שאצל |  |
| 6:32.318 | וברא |  |
| 6:32.319 | ויצר |  |
| 6:32.320 | ועשה |  |
| 6:32.321 | אלהים |  |
| 6:32.322 | קדושים |  |
| 6:32.323 | כן |  |
| 6:32.324 | אצל |  |
| 6:32.325 | וברא |  |
| 6:32.326 | ויצר |  |
| 6:32.327 | ועשה |  |
| 6:32.328 | אלהים |  |
| 6:32.329 | אחרים |  |
| 6:32.330 | וכמו |  |
| 6:32.331 | שיש |  |
| 6:32.332 | בקדושה |  |
| 6:32.333 | אדם |  |
| 6:32.334 | קדמון |  |
| 6:32.335 | ועתיק |  |
| 6:32.336 | ואריך | , |
| 6:32.337 | ואבא |  |
| 6:32.338 | ואימא | , |
| 6:32.339 | וזעיר |  |
| 6:32.340 | ונוקביה |  |
| 6:32.341 | דקדושה |  |
| 6:32.342 | כן |  |
| 6:32.343 | בשמאל |  |
| 6:32.344 | אדם |  |
| 6:32.345 | בליעל |  |
| 6:32.346 | סטרא |  |
| 6:32.347 | אחרא |  |
| 6:32.348 | מסאבא | , |
| 6:32.349 | זה |  |
| 6:32.350 | לעומת |  |
| 6:32.351 | זה |  |
| 6:32.352 | כמו |  |
| 6:32.353 | שיראה |  |
| 6:32.354 | המעיין |  |
| 6:32.355 | בס | ' |
| 6:32.356 | חיים |  |
| 6:32.357 | למהרח"ו | , |
| 6:32.358 | ומשנת |  |
| 6:32.359 | חסידים |  |
| 6:32.360 | שער |  |
| 6:32.361 | הקליפות | . |
| 6:32.362 | ואם |  |
| 6:32.363 | תעמיק |  |
| 6:32.364 | חקר |  |
| 6:32.365 | תבין |  |
| 6:32.366 | ותשכיל |  |
| 6:32.367 | שסביא |  |
| 6:32.368 | זו |  |
| 6:32.369 | באה |  |
| 6:32.370 | ונמשה |  |
| 6:32.371 | להם |  |
| 6:32.372 | מאמונת |  |
| 6:32.373 | השונים |  |
| 6:32.374 | שסוברים |  |
| 6:32.375 | שיש |  |
| 6:32.376 | בעולם |  |
| 6:32.377 | שני |  |
| 6:32.378 | פועלים | , |
| 6:32.379 | אחד |  |
| 6:32.380 | פועל |  |
| 6:32.381 | הטוב | , |
| 6:32.382 | ואחד |  |
| 6:32.383 | פועל |  |
| 6:32.384 | הרע | . |
| 6:32.385 | כמו |  |
| 6:32.386 | שכתבו |  |
| 6:32.387 | הראשונים |  |
| 6:32.388 | לדחותה |  |
| 6:32.389 | בשתי |  |
| 6:32.390 | ידים | , |
| 6:32.391 | והמקובלים |  |
| 6:32.392 | הרדשים |  |
| 6:32.393 | תפשוה |  |
| 6:32.394 | וציירו |  |
| 6:32.395 | אותה |  |
| 6:32.396 | בפנים |  |
| 6:32.397 | אחרים | , |
| 6:32.398 | ונדחקו |  |
| 6:32.399 | מאד |  |
| 6:32.400 | לכסות |  |
| 6:32.401 | את |  |
| 6:32.402 | נבלותה | , |
| 6:32.403 | ולהלבישה |  |
| 6:32.404 | באמונת |  |
| 6:32.405 | תוח"ק | . |
| 6:32.406 | ובאמת |  |
| 6:32.407 | זרה |  |
| 6:32.408 | היא |  |
| 6:32.409 | עד |  |
| 6:32.410 | מאד |  |
| 6:32.411 | לאמונת |  |
| 6:32.412 | יהוד |  |
| 6:32.413 | השי"ת | , |
| 6:32.414 | ולא |  |
| 6:32.415 | תבוא |  |
| 6:32.416 | בקהל |  |
| 6:32.417 | ה | '. |
| 6:32.418 | ולאו |  |
| 6:32.419 | קבלה |  |
| 6:32.420 | היא |  |
| 6:32.421 | כלל | . |
| 6:32.422 | כמו |  |
| 6:32.423 | שכ | ' |
| 6:32.424 | לעיל |  |
| 6:32.425 | מדברי |  |
| 6:32.426 | רז"ל | . |
| 6:32.427 | כי |  |
| 6:32.428 | המקובלים |  |
| 6:32.429 | החדשים |  |
| 6:32.430 | מאמתים |  |
| 6:32.431 | מציאות |  |
| 6:32.432 | אל |  |
| 6:32.433 | אחד |  |
| 6:32.434 | בעולם |  |
| 6:32.435 | היפך |  |
| 6:32.436 | דברי |  |
| 6:32.437 | רז"ל | . |
| 6:32.438 | והכתובים |  |
| 6:32.439 | מוכיחים |  |
| 6:32.440 | כדעת |  |
| 6:32.441 | חכמים |  |
| 6:32.442 | ז"ל | . |
| 6:32.443 | שנ | ' |
| 6:32.444 | כי |  |
| 6:32.445 | הוא |  |
| 6:32.446 | האלהים |  |
| 6:32.447 | בשמים |  |
| 6:32.448 | ממעל |  |
| 6:32.449 | ועל |  |
| 6:32.450 | הארץ |  |
| 6:32.451 | מתחת |  |
| 6:32.452 | אין |  |
| 6:32.453 | עוד |  |
| 6:32.454 | מלבדו | . |
| 6:32.455 | אתה |  |
| 6:32.456 | הוא |  |
| 6:32.457 | ה | ' |
| 6:32.458 | לבדך | . |
| 6:32.459 | וכמה |  |
| 6:32.460 | כתובים |  |
| 6:32.461 | מכריזים |  |
| 6:32.462 | שאין |  |
| 6:32.463 | עוד |  |
| 6:32.464 | אלוה |  |
| 6:32.465 | אחר |  |
| 6:32.466 | זולתי |  |
| 6:32.467 | ה | ' |
| 6:32.468 | אלהינו | . |
| 6:32.469 | ורבותינו |  |
| 6:32.470 | מתקני |  |
| 6:32.471 | התפלות |  |
| 6:32.472 | התקינו |  |
| 6:32.473 | לומר |  |
| 6:32.474 | בכל |  |
| 6:32.475 | יום |  |
| 6:32.476 | אנו |  |
| 6:32.477 | יודעים |  |
| 6:32.478 | ומעידים |  |
| 6:32.479 | ומגידים | , |
| 6:32.480 | שאין |  |
| 6:32.481 | אלוה |  |
| 6:32.482 | זולתך |  |
| 6:32.483 | וכו | ' |
| 6:32.484 | והרבה |  |
| 6:32.485 | לשונות |  |
| 6:32.486 | בתפלה |  |
| 6:32.487 | מורים |  |
| 6:32.488 | שאין |  |
| 6:32.489 | מציאות |  |
| 6:32.490 | אל |  |
| 6:32.491 | אחר |  |
| 6:32.492 | בעולם |  |
| 6:32.493 | כלל |  |
| 6:32.494 | כלל | . |
| 6:32.495 | ורק |  |
| 6:32.496 | לפי |  |
| 6:32.497 | מחשבת |  |
| 6:32.498 | ובדיהן | . |
| 6:32.499 | אמרה |  |
| 6:32.500 | כי |  |
| 6:32.501 | תורה |  |
| 6:32.502 | כי |  |
| 6:32.503 | לא |  |
| 6:32.504 | תשתחוה |  |
| 6:32.505 | לאל |  |
| 6:32.506 | אחר | . |
| 6:32.507 | וכן |  |
| 6:32.508 | לא |  |
| 6:32.509 | יהיה |  |
| 6:32.510 | לך |  |
| 6:32.511 | אלהים |  |
| 6:32.512 | אחרים |  |
| 6:32.513 | על |  |
| 6:32.514 | פני | . |
| 6:32.515 | הכל |  |
| 6:32.516 | לפי |  |
| 6:32.517 | מחשבת |  |
| 6:32.518 | עובידהן |  |
| 6:32.519 | שטועין |  |
| 6:32.520 | וחושבים |  |
| 6:32.521 | אותם |  |
| 6:32.522 | אלוהות | . |
| 6:32.523 | לא |  |
| 6:32.524 | שהשי"ת |  |
| 6:32.525 | המציא |  |
| 6:32.526 | שני |  |
| 6:32.527 | מיני |  |
| 6:32.528 | אלוהות |  |
| 6:32.529 | אלה |  |
| 6:32.530 | קדושים |  |
| 6:32.531 | ואלה |  |
| 6:32.532 | טמאים |  |
| 6:32.533 | ח"ו | . |
| 6:32.534 | כי |  |
| 6:32.535 | אין |  |
| 6:32.536 | אלוה |  |
| 6:32.537 | זולתו | . |
| 6:32.538 | וכמה |  |
| 6:32.539 | פסורים |  |
| 6:32.540 | בתורה |  |
| 6:32.541 | ובנביאים |  |
| 6:32.542 | ובכתובים |  |
| 6:32.543 | מוכיחים |  |
| 6:32.544 | כן | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

