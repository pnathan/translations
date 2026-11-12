# Review pack: milhamot-hashem perek-2-08 (Milchamot ha-Shem 2:16–2:16)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-08`.

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
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **אנדרטא** → "statue"; **אשם תלוי** → "suspended guilt offering".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **אין לנו להחמיר** → "we may not be stringent".
- **ס' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in his
  numbered sections; by context, the choice noted.

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


## Milchamot ha-Shem 2:16

Hebrew (Wikisource): ומה שכתבת: ויותר מזה מה שכ' על כבוד ה' שהוא אור נברא אם כן כיצד היו ישראל יוצאים פתח אהליהם ומשתחוים לנברא? לא לכבוד הנראה משתחוים, רק למי שהשרה כבודו שם. והכבוד הנראה שם לאות וסימן להם, כי השי"ת איווה לאוהל מועד למושב לו, כמו שאנחנו משתחוים ומתפללים לשי"ת נכח המקדש וירושלם, כי בחר ה' בציון אוה למושב לו. וכבר ביאר הרמב"ם בפ' ס"ד מראשון, כי כבוד ה', פעמים רוצים בו האור הנברא אשר ישכינהו הק' במקום, כמו: וישכון כבוד ה' על הר סיני ויכסהו וכו', ופעמים רוצים בו עצמותו ית' ואמתתו, כמו "הראיני נא את כבודך", ובא המענה: "כי לא יראני האדם וחי", מורה כי הכבוד הנאמר פה עצמו ית', ואמר "כבודך" להגדיל וכו'. וכיוצא בזה מצינו בדברי רז"ל ששם אלהים משותף לדיינים ולמלאכים ולשי"ת, גם שם שכינה משותף למה שהזכרתי למעלה. וראיתי להעירך פה שאין איסור להשתחוות לאדם כל שאין דעתו לקבלו באלוה וגם אינו נעבד, רק נכנע לו ומשפיל עצמו לפניו, או דרך ההודאה על טובה שעשה לו. שהרי אברהם אבינו השתחוה לבני חת על החסד שעשו עמו ליתן לו אחזת קבר. וכן תרגם רבי סעדיה גאון: וישתחו לעם הארץ לבני חת, וסגד לאהל אלבלד שברא. ויהושוע השתחווה למלאך. ונתן הנביא ובת שבע השתחוו לדוד. והרבה מצינו בכתובים שהשתחוו לנברא. וגרסינן בסנהדרין פ' ארבע מיתות (סנהדרין ס"א ע"ב), אתמר: העובד ע"ז מאהבה ומיראה, אביי אמר חייב, דהא פלחה. רבא אמר פטור, אי קבליה עליה באלוה אין, אי לא לא. ואמר אביי מנא אמינא לה, דתנן העובד ע"ז אחד העובד ואחד המזבח ואחד המקטיר ואחד המנסך ואחד המשתחוה ואחד המקבלו עליו באלוה והאומר לו אלי אתה. מאי לאו אחד העובד מאהבה או מיראה? ורבא אמר לך, לא, אלא כדמתרץ ר' ירמיה (פירש רש"י: אחד העובד כל עבודה שהיא דרכה בכך, ואחד המזבח וכו' שלא כדרכה, ולעולם רישא וסיפא לשם אלהות). אמר אביי: מנא אמינא לה? דתניא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך. יכול אפי' לעבד כהמן? ת"ל ולא תעבדם. והא המן מיראה הוא נעבד! ורבא? כהמן ולא כהמן. כהמן, דאיהו גופיה ע"ז; ולא כהמן, דאלו המן מיראה, והכא לאו מיראה. ואמר אביי: מנא אמינא לה? דתניא: כהן משיח בע"ז. ר' אומר בשגגת מעשה, וחכמים אומרים בהעלם דבר, ושווין דבשעירה כיחיד, ושווין שאין מביא אשם תלוי. האי שגגת מעשה היכי דמי? אי קסבר ביהכנ"ס והשתחוה לו, הרי לבו לשמים. אלא דחזא אנדרטא והשתחוה לו (פירוש, שראה דמות המלך שהיו רגילים לעשות כדמות המלך והרואה אותו משתחוה לו לכבוד המלך, ופעמים שעובדין אותו, וזה ראה אחד שנעבד והשתחוה לו), אי קבליה עליה באלוה מזיד הוא, ואי לא קבליה עליה באלוה לאו כלום הוא (פירש רש"י: אלא לכבוד המלך השתחוה לו, לאו כלום הוא דהא לא איכוון לע"ז). אלא לאו מאהבה ומיראה? ורבא אמר לך: באומר מותר. אומר מותר היינו העלם דבר? באומר מותר לגמרי, העלם דבר קיום מקצת וביטול מקצת, עכ"ל הגמ' עם קצת פי' רש"י. ודבר נלמד, דלכל נברא אסור להשתחוות, זולת האדם, והוא שלא נעבד כהמן שעשה עצמו ע"ז, כדאמירן במגלה (יט), דאי לאו הכי לא היה מרדכי נמנע מלהשתחות לו, כדקתני בבריתא: לא תשתחוה להם, להם אי אתה משתחוה, אבל אתה משתחוה לאדם כמותך, והוא שלא יהא נעבד כהמן שעשה עצמו אלוה, רק משתחוה לו לאות הכבוד ולא לאלהות, ולכך השתחוה יהושע למלאך הואיל ונגלה לו בדמות אדם, כמ"ש: והנה איש עומד לנגדו, ולכן השתחוה לו לאות הכבוד ולא לעבדו. ועיין בס' העיקרים מאמר ב' פ' כ"ח שהאריך בביאור הכתובים שקורין לכבוד הנראה לנביאים בשם הוי"ה, וגם לארון, כמו שגאמר: "ויהי בנסע הארון וכו' קומה ה', ובנחה יאמר שובה ה'", ויהושע קרא לארון "אדון כל הארץ" ע"ש. ולדעתו שמה שהשתחוה יהושע למלאך, לפי שההשתחויה ההיא היתה לכבוד השי"ת, ולהיות זה שלוחו וכו' ע"ש. ורז"ל אמרו תעלא בעידניה סגיד ליה, ר"ל: אדם פחות ובזוי כשועל, והשעה עומדת לו באיזה משרה שבידו, סגיד ליה. וכן הוא נהוג אצלנו ואצל הערביים להשתחוות ולנשק אדם גדול המעלה, ואין בזה שום איסור, שאינו מקבלו עליו באלוה, רק לאות כבוד הוא משתחוה לו ומנשק את רגליו.

| id | word | punctuation after |
|---|---|---|
| 2:16.1 | ומה |  |
| 2:16.2 | שכתבת | : |
| 2:16.3 | ויותר |  |
| 2:16.4 | מזה |  |
| 2:16.5 | מה |  |
| 2:16.6 | שכ | ' |
| 2:16.7 | על |  |
| 2:16.8 | כבוד |  |
| 2:16.9 | ה | ' |
| 2:16.10 | שהוא |  |
| 2:16.11 | אור |  |
| 2:16.12 | נברא |  |
| 2:16.13 | אם |  |
| 2:16.14 | כן |  |
| 2:16.15 | כיצד |  |
| 2:16.16 | היו |  |
| 2:16.17 | ישראל |  |
| 2:16.18 | יוצאים |  |
| 2:16.19 | פתח |  |
| 2:16.20 | אהליהם |  |
| 2:16.21 | ומשתחוים |  |
| 2:16.22 | לנברא | ? |
| 2:16.23 | לא |  |
| 2:16.24 | לכבוד |  |
| 2:16.25 | הנראה |  |
| 2:16.26 | משתחוים | , |
| 2:16.27 | רק |  |
| 2:16.28 | למי |  |
| 2:16.29 | שהשרה |  |
| 2:16.30 | כבודו |  |
| 2:16.31 | שם | . |
| 2:16.32 | והכבוד |  |
| 2:16.33 | הנראה |  |
| 2:16.34 | שם |  |
| 2:16.35 | לאות |  |
| 2:16.36 | וסימן |  |
| 2:16.37 | להם | , |
| 2:16.38 | כי |  |
| 2:16.39 | השי"ת |  |
| 2:16.40 | איווה |  |
| 2:16.41 | לאוהל |  |
| 2:16.42 | מועד |  |
| 2:16.43 | למושב |  |
| 2:16.44 | לו | , |
| 2:16.45 | כמו |  |
| 2:16.46 | שאנחנו |  |
| 2:16.47 | משתחוים |  |
| 2:16.48 | ומתפללים |  |
| 2:16.49 | לשי"ת |  |
| 2:16.50 | נכח |  |
| 2:16.51 | המקדש |  |
| 2:16.52 | וירושלם | , |
| 2:16.53 | כי |  |
| 2:16.54 | בחר |  |
| 2:16.55 | ה | ' |
| 2:16.56 | בציון |  |
| 2:16.57 | אוה |  |
| 2:16.58 | למושב |  |
| 2:16.59 | לו | . |
| 2:16.60 | וכבר |  |
| 2:16.61 | ביאר |  |
| 2:16.62 | הרמב"ם |  |
| 2:16.63 | בפ | ' |
| 2:16.64 | ס"ד |  |
| 2:16.65 | מראשון | , |
| 2:16.66 | כי |  |
| 2:16.67 | כבוד |  |
| 2:16.68 | ה | ', |
| 2:16.69 | פעמים |  |
| 2:16.70 | רוצים |  |
| 2:16.71 | בו |  |
| 2:16.72 | האור |  |
| 2:16.73 | הנברא |  |
| 2:16.74 | אשר |  |
| 2:16.75 | ישכינהו |  |
| 2:16.76 | הק | ' |
| 2:16.77 | במקום | , |
| 2:16.78 | כמו | : |
| 2:16.79 | וישכון |  |
| 2:16.80 | כבוד |  |
| 2:16.81 | ה | ' |
| 2:16.82 | על |  |
| 2:16.83 | הר |  |
| 2:16.84 | סיני |  |
| 2:16.85 | ויכסהו |  |
| 2:16.86 | וכו | ', |
| 2:16.87 | ופעמים |  |
| 2:16.88 | רוצים |  |
| 2:16.89 | בו |  |
| 2:16.90 | עצמותו |  |
| 2:16.91 | ית | ' |
| 2:16.92 | ואמתתו | , |
| 2:16.93 | כמו |  |
| 2:16.94 | הראיני |  |
| 2:16.95 | נא |  |
| 2:16.96 | את |  |
| 2:16.97 | כבודך | ", |
| 2:16.98 | ובא |  |
| 2:16.99 | המענה | : |
| 2:16.100 | כי |  |
| 2:16.101 | לא |  |
| 2:16.102 | יראני |  |
| 2:16.103 | האדם |  |
| 2:16.104 | וחי | ", |
| 2:16.105 | מורה |  |
| 2:16.106 | כי |  |
| 2:16.107 | הכבוד |  |
| 2:16.108 | הנאמר |  |
| 2:16.109 | פה |  |
| 2:16.110 | עצמו |  |
| 2:16.111 | ית | ', |
| 2:16.112 | ואמר |  |
| 2:16.113 | כבודך | " |
| 2:16.114 | להגדיל |  |
| 2:16.115 | וכו | '. |
| 2:16.116 | וכיוצא |  |
| 2:16.117 | בזה |  |
| 2:16.118 | מצינו |  |
| 2:16.119 | בדברי |  |
| 2:16.120 | רז"ל |  |
| 2:16.121 | ששם |  |
| 2:16.122 | אלהים |  |
| 2:16.123 | משותף |  |
| 2:16.124 | לדיינים |  |
| 2:16.125 | ולמלאכים |  |
| 2:16.126 | ולשי"ת | , |
| 2:16.127 | גם |  |
| 2:16.128 | שם |  |
| 2:16.129 | שכינה |  |
| 2:16.130 | משותף |  |
| 2:16.131 | למה |  |
| 2:16.132 | שהזכרתי |  |
| 2:16.133 | למעלה | . |
| 2:16.134 | וראיתי |  |
| 2:16.135 | להעירך |  |
| 2:16.136 | פה |  |
| 2:16.137 | שאין |  |
| 2:16.138 | איסור |  |
| 2:16.139 | להשתחוות |  |
| 2:16.140 | לאדם |  |
| 2:16.141 | כל |  |
| 2:16.142 | שאין |  |
| 2:16.143 | דעתו |  |
| 2:16.144 | לקבלו |  |
| 2:16.145 | באלוה |  |
| 2:16.146 | וגם |  |
| 2:16.147 | אינו |  |
| 2:16.148 | נעבד | , |
| 2:16.149 | רק |  |
| 2:16.150 | נכנע |  |
| 2:16.151 | לו |  |
| 2:16.152 | ומשפיל |  |
| 2:16.153 | עצמו |  |
| 2:16.154 | לפניו | , |
| 2:16.155 | או |  |
| 2:16.156 | דרך |  |
| 2:16.157 | ההודאה |  |
| 2:16.158 | על |  |
| 2:16.159 | טובה |  |
| 2:16.160 | שעשה |  |
| 2:16.161 | לו | . |
| 2:16.162 | שהרי |  |
| 2:16.163 | אברהם |  |
| 2:16.164 | אבינו |  |
| 2:16.165 | השתחוה |  |
| 2:16.166 | לבני |  |
| 2:16.167 | חת |  |
| 2:16.168 | על |  |
| 2:16.169 | החסד |  |
| 2:16.170 | שעשו |  |
| 2:16.171 | עמו |  |
| 2:16.172 | ליתן |  |
| 2:16.173 | לו |  |
| 2:16.174 | אחזת |  |
| 2:16.175 | קבר | . |
| 2:16.176 | וכן |  |
| 2:16.177 | תרגם |  |
| 2:16.178 | רבי |  |
| 2:16.179 | סעדיה |  |
| 2:16.180 | גאון | : |
| 2:16.181 | וישתחו |  |
| 2:16.182 | לעם |  |
| 2:16.183 | הארץ |  |
| 2:16.184 | לבני |  |
| 2:16.185 | חת | , |
| 2:16.186 | וסגד |  |
| 2:16.187 | לאהל |  |
| 2:16.188 | אלבלד |  |
| 2:16.189 | שברא | . |
| 2:16.190 | ויהושוע |  |
| 2:16.191 | השתחווה |  |
| 2:16.192 | למלאך | . |
| 2:16.193 | ונתן |  |
| 2:16.194 | הנביא |  |
| 2:16.195 | ובת |  |
| 2:16.196 | שבע |  |
| 2:16.197 | השתחוו |  |
| 2:16.198 | לדוד | . |
| 2:16.199 | והרבה |  |
| 2:16.200 | מצינו |  |
| 2:16.201 | בכתובים |  |
| 2:16.202 | שהשתחוו |  |
| 2:16.203 | לנברא | . |
| 2:16.204 | וגרסינן |  |
| 2:16.205 | בסנהדרין |  |
| 2:16.206 | פ | ' |
| 2:16.207 | ארבע |  |
| 2:16.208 | מיתות |  |
| 2:16.209 | סנהדרין |  |
| 2:16.210 | ס"א |  |
| 2:16.211 | ע"ב | ), |
| 2:16.212 | אתמר | : |
| 2:16.213 | העובד |  |
| 2:16.214 | ע"ז |  |
| 2:16.215 | מאהבה |  |
| 2:16.216 | ומיראה | , |
| 2:16.217 | אביי |  |
| 2:16.218 | אמר |  |
| 2:16.219 | חייב | , |
| 2:16.220 | דהא |  |
| 2:16.221 | פלחה | . |
| 2:16.222 | רבא |  |
| 2:16.223 | אמר |  |
| 2:16.224 | פטור | , |
| 2:16.225 | אי |  |
| 2:16.226 | קבליה |  |
| 2:16.227 | עליה |  |
| 2:16.228 | באלוה |  |
| 2:16.229 | אין | , |
| 2:16.230 | אי |  |
| 2:16.231 | לא |  |
| 2:16.232 | לא | . |
| 2:16.233 | ואמר |  |
| 2:16.234 | אביי |  |
| 2:16.235 | מנא |  |
| 2:16.236 | אמינא |  |
| 2:16.237 | לה | , |
| 2:16.238 | דתנן |  |
| 2:16.239 | העובד |  |
| 2:16.240 | ע"ז |  |
| 2:16.241 | אחד |  |
| 2:16.242 | העובד |  |
| 2:16.243 | ואחד |  |
| 2:16.244 | המזבח |  |
| 2:16.245 | ואחד |  |
| 2:16.246 | המקטיר |  |
| 2:16.247 | ואחד |  |
| 2:16.248 | המנסך |  |
| 2:16.249 | ואחד |  |
| 2:16.250 | המשתחוה |  |
| 2:16.251 | ואחד |  |
| 2:16.252 | המקבלו |  |
| 2:16.253 | עליו |  |
| 2:16.254 | באלוה |  |
| 2:16.255 | והאומר |  |
| 2:16.256 | לו |  |
| 2:16.257 | אלי |  |
| 2:16.258 | אתה | . |
| 2:16.259 | מאי |  |
| 2:16.260 | לאו |  |
| 2:16.261 | אחד |  |
| 2:16.262 | העובד |  |
| 2:16.263 | מאהבה |  |
| 2:16.264 | או |  |
| 2:16.265 | מיראה | ? |
| 2:16.266 | ורבא |  |
| 2:16.267 | אמר |  |
| 2:16.268 | לך | , |
| 2:16.269 | לא | , |
| 2:16.270 | אלא |  |
| 2:16.271 | כדמתרץ |  |
| 2:16.272 | ר | ' |
| 2:16.273 | ירמיה |  |
| 2:16.274 | פירש |  |
| 2:16.275 | רש"י | : |
| 2:16.276 | אחד |  |
| 2:16.277 | העובד |  |
| 2:16.278 | כל |  |
| 2:16.279 | עבודה |  |
| 2:16.280 | שהיא |  |
| 2:16.281 | דרכה |  |
| 2:16.282 | בכך | , |
| 2:16.283 | ואחד |  |
| 2:16.284 | המזבח |  |
| 2:16.285 | וכו | ' |
| 2:16.286 | שלא |  |
| 2:16.287 | כדרכה | , |
| 2:16.288 | ולעולם |  |
| 2:16.289 | רישא |  |
| 2:16.290 | וסיפא |  |
| 2:16.291 | לשם |  |
| 2:16.292 | אלהות | ). |
| 2:16.293 | אמר |  |
| 2:16.294 | אביי | : |
| 2:16.295 | מנא |  |
| 2:16.296 | אמינא |  |
| 2:16.297 | לה | ? |
| 2:16.298 | דתניא | : |
| 2:16.299 | לא |  |
| 2:16.300 | תשתחוה |  |
| 2:16.301 | להם | , |
| 2:16.302 | להם |  |
| 2:16.303 | אי |  |
| 2:16.304 | אתה |  |
| 2:16.305 | משתחוה | , |
| 2:16.306 | אבל |  |
| 2:16.307 | אתה |  |
| 2:16.308 | משתחוה |  |
| 2:16.309 | לאדם |  |
| 2:16.310 | כמותך | . |
| 2:16.311 | יכול |  |
| 2:16.312 | אפי | ' |
| 2:16.313 | לעבד |  |
| 2:16.314 | כהמן | ? |
| 2:16.315 | ת"ל |  |
| 2:16.316 | ולא |  |
| 2:16.317 | תעבדם | . |
| 2:16.318 | והא |  |
| 2:16.319 | המן |  |
| 2:16.320 | מיראה |  |
| 2:16.321 | הוא |  |
| 2:16.322 | נעבד | ! |
| 2:16.323 | ורבא | ? |
| 2:16.324 | כהמן |  |
| 2:16.325 | ולא |  |
| 2:16.326 | כהמן | . |
| 2:16.327 | כהמן | , |
| 2:16.328 | דאיהו |  |
| 2:16.329 | גופיה |  |
| 2:16.330 | ע"ז | ; |
| 2:16.331 | ולא |  |
| 2:16.332 | כהמן | , |
| 2:16.333 | דאלו |  |
| 2:16.334 | המן |  |
| 2:16.335 | מיראה | , |
| 2:16.336 | והכא |  |
| 2:16.337 | לאו |  |
| 2:16.338 | מיראה | . |
| 2:16.339 | ואמר |  |
| 2:16.340 | אביי | : |
| 2:16.341 | מנא |  |
| 2:16.342 | אמינא |  |
| 2:16.343 | לה | ? |
| 2:16.344 | דתניא | : |
| 2:16.345 | כהן |  |
| 2:16.346 | משיח |  |
| 2:16.347 | בע"ז | . |
| 2:16.348 | ר | ' |
| 2:16.349 | אומר |  |
| 2:16.350 | בשגגת |  |
| 2:16.351 | מעשה | , |
| 2:16.352 | וחכמים |  |
| 2:16.353 | אומרים |  |
| 2:16.354 | בהעלם |  |
| 2:16.355 | דבר | , |
| 2:16.356 | ושווין |  |
| 2:16.357 | דבשעירה |  |
| 2:16.358 | כיחיד | , |
| 2:16.359 | ושווין |  |
| 2:16.360 | שאין |  |
| 2:16.361 | מביא |  |
| 2:16.362 | אשם |  |
| 2:16.363 | תלוי | . |
| 2:16.364 | האי |  |
| 2:16.365 | שגגת |  |
| 2:16.366 | מעשה |  |
| 2:16.367 | היכי |  |
| 2:16.368 | דמי | ? |
| 2:16.369 | אי |  |
| 2:16.370 | קסבר |  |
| 2:16.371 | ביהכנ"ס |  |
| 2:16.372 | והשתחוה |  |
| 2:16.373 | לו | , |
| 2:16.374 | הרי |  |
| 2:16.375 | לבו |  |
| 2:16.376 | לשמים | . |
| 2:16.377 | אלא |  |
| 2:16.378 | דחזא |  |
| 2:16.379 | אנדרטא |  |
| 2:16.380 | והשתחוה |  |
| 2:16.381 | לו |  |
| 2:16.382 | פירוש | , |
| 2:16.383 | שראה |  |
| 2:16.384 | דמות |  |
| 2:16.385 | המלך |  |
| 2:16.386 | שהיו |  |
| 2:16.387 | רגילים |  |
| 2:16.388 | לעשות |  |
| 2:16.389 | כדמות |  |
| 2:16.390 | המלך |  |
| 2:16.391 | והרואה |  |
| 2:16.392 | אותו |  |
| 2:16.393 | משתחוה |  |
| 2:16.394 | לו |  |
| 2:16.395 | לכבוד |  |
| 2:16.396 | המלך | , |
| 2:16.397 | ופעמים |  |
| 2:16.398 | שעובדין |  |
| 2:16.399 | אותו | , |
| 2:16.400 | וזה |  |
| 2:16.401 | ראה |  |
| 2:16.402 | אחד |  |
| 2:16.403 | שנעבד |  |
| 2:16.404 | והשתחוה |  |
| 2:16.405 | לו | ), |
| 2:16.406 | אי |  |
| 2:16.407 | קבליה |  |
| 2:16.408 | עליה |  |
| 2:16.409 | באלוה |  |
| 2:16.410 | מזיד |  |
| 2:16.411 | הוא | , |
| 2:16.412 | ואי |  |
| 2:16.413 | לא |  |
| 2:16.414 | קבליה |  |
| 2:16.415 | עליה |  |
| 2:16.416 | באלוה |  |
| 2:16.417 | לאו |  |
| 2:16.418 | כלום |  |
| 2:16.419 | הוא |  |
| 2:16.420 | פירש |  |
| 2:16.421 | רש"י | : |
| 2:16.422 | אלא |  |
| 2:16.423 | לכבוד |  |
| 2:16.424 | המלך |  |
| 2:16.425 | השתחוה |  |
| 2:16.426 | לו | , |
| 2:16.427 | לאו |  |
| 2:16.428 | כלום |  |
| 2:16.429 | הוא |  |
| 2:16.430 | דהא |  |
| 2:16.431 | לא |  |
| 2:16.432 | איכוון |  |
| 2:16.433 | לע"ז | ). |
| 2:16.434 | אלא |  |
| 2:16.435 | לאו |  |
| 2:16.436 | מאהבה |  |
| 2:16.437 | ומיראה | ? |
| 2:16.438 | ורבא |  |
| 2:16.439 | אמר |  |
| 2:16.440 | לך | : |
| 2:16.441 | באומר |  |
| 2:16.442 | מותר | . |
| 2:16.443 | אומר |  |
| 2:16.444 | מותר |  |
| 2:16.445 | היינו |  |
| 2:16.446 | העלם |  |
| 2:16.447 | דבר | ? |
| 2:16.448 | באומר |  |
| 2:16.449 | מותר |  |
| 2:16.450 | לגמרי | , |
| 2:16.451 | העלם |  |
| 2:16.452 | דבר |  |
| 2:16.453 | קיום |  |
| 2:16.454 | מקצת |  |
| 2:16.455 | וביטול |  |
| 2:16.456 | מקצת | , |
| 2:16.457 | עכ"ל |  |
| 2:16.458 | הגמ | ' |
| 2:16.459 | עם |  |
| 2:16.460 | קצת |  |
| 2:16.461 | פי | ' |
| 2:16.462 | רש"י | . |
| 2:16.463 | ודבר |  |
| 2:16.464 | נלמד | , |
| 2:16.465 | דלכל |  |
| 2:16.466 | נברא |  |
| 2:16.467 | אסור |  |
| 2:16.468 | להשתחוות | , |
| 2:16.469 | זולת |  |
| 2:16.470 | האדם | , |
| 2:16.471 | והוא |  |
| 2:16.472 | שלא |  |
| 2:16.473 | נעבד |  |
| 2:16.474 | כהמן |  |
| 2:16.475 | שעשה |  |
| 2:16.476 | עצמו |  |
| 2:16.477 | ע"ז | , |
| 2:16.478 | כדאמירן |  |
| 2:16.479 | במגלה |  |
| 2:16.480 | יט | ), |
| 2:16.481 | דאי |  |
| 2:16.482 | לאו |  |
| 2:16.483 | הכי |  |
| 2:16.484 | לא |  |
| 2:16.485 | היה |  |
| 2:16.486 | מרדכי |  |
| 2:16.487 | נמנע |  |
| 2:16.488 | מלהשתחות |  |
| 2:16.489 | לו | , |
| 2:16.490 | כדקתני |  |
| 2:16.491 | בבריתא | : |
| 2:16.492 | לא |  |
| 2:16.493 | תשתחוה |  |
| 2:16.494 | להם | , |
| 2:16.495 | להם |  |
| 2:16.496 | אי |  |
| 2:16.497 | אתה |  |
| 2:16.498 | משתחוה | , |
| 2:16.499 | אבל |  |
| 2:16.500 | אתה |  |
| 2:16.501 | משתחוה |  |
| 2:16.502 | לאדם |  |
| 2:16.503 | כמותך | , |
| 2:16.504 | והוא |  |
| 2:16.505 | שלא |  |
| 2:16.506 | יהא |  |
| 2:16.507 | נעבד |  |
| 2:16.508 | כהמן |  |
| 2:16.509 | שעשה |  |
| 2:16.510 | עצמו |  |
| 2:16.511 | אלוה | , |
| 2:16.512 | רק |  |
| 2:16.513 | משתחוה |  |
| 2:16.514 | לו |  |
| 2:16.515 | לאות |  |
| 2:16.516 | הכבוד |  |
| 2:16.517 | ולא |  |
| 2:16.518 | לאלהות | , |
| 2:16.519 | ולכך |  |
| 2:16.520 | השתחוה |  |
| 2:16.521 | יהושע |  |
| 2:16.522 | למלאך |  |
| 2:16.523 | הואיל |  |
| 2:16.524 | ונגלה |  |
| 2:16.525 | לו |  |
| 2:16.526 | בדמות |  |
| 2:16.527 | אדם | , |
| 2:16.528 | כמ"ש | : |
| 2:16.529 | והנה |  |
| 2:16.530 | איש |  |
| 2:16.531 | עומד |  |
| 2:16.532 | לנגדו | , |
| 2:16.533 | ולכן |  |
| 2:16.534 | השתחוה |  |
| 2:16.535 | לו |  |
| 2:16.536 | לאות |  |
| 2:16.537 | הכבוד |  |
| 2:16.538 | ולא |  |
| 2:16.539 | לעבדו | . |
| 2:16.540 | ועיין |  |
| 2:16.541 | בס | ' |
| 2:16.542 | העיקרים |  |
| 2:16.543 | מאמר |  |
| 2:16.544 | ב | ' |
| 2:16.545 | פ | ' |
| 2:16.546 | כ"ח |  |
| 2:16.547 | שהאריך |  |
| 2:16.548 | בביאור |  |
| 2:16.549 | הכתובים |  |
| 2:16.550 | שקורין |  |
| 2:16.551 | לכבוד |  |
| 2:16.552 | הנראה |  |
| 2:16.553 | לנביאים |  |
| 2:16.554 | בשם |  |
| 2:16.555 | הוי"ה | , |
| 2:16.556 | וגם |  |
| 2:16.557 | לארון | , |
| 2:16.558 | כמו |  |
| 2:16.559 | שגאמר | : |
| 2:16.560 | ויהי |  |
| 2:16.561 | בנסע |  |
| 2:16.562 | הארון |  |
| 2:16.563 | וכו | ' |
| 2:16.564 | קומה |  |
| 2:16.565 | ה | ', |
| 2:16.566 | ובנחה |  |
| 2:16.567 | יאמר |  |
| 2:16.568 | שובה |  |
| 2:16.569 | ה | '", |
| 2:16.570 | ויהושע |  |
| 2:16.571 | קרא |  |
| 2:16.572 | לארון |  |
| 2:16.573 | אדון |  |
| 2:16.574 | כל |  |
| 2:16.575 | הארץ | " |
| 2:16.576 | ע"ש | . |
| 2:16.577 | ולדעתו |  |
| 2:16.578 | שמה |  |
| 2:16.579 | שהשתחוה |  |
| 2:16.580 | יהושע |  |
| 2:16.581 | למלאך | , |
| 2:16.582 | לפי |  |
| 2:16.583 | שההשתחויה |  |
| 2:16.584 | ההיא |  |
| 2:16.585 | היתה |  |
| 2:16.586 | לכבוד |  |
| 2:16.587 | השי"ת | , |
| 2:16.588 | ולהיות |  |
| 2:16.589 | זה |  |
| 2:16.590 | שלוחו |  |
| 2:16.591 | וכו | ' |
| 2:16.592 | ע"ש | . |
| 2:16.593 | ורז"ל |  |
| 2:16.594 | אמרו |  |
| 2:16.595 | תעלא |  |
| 2:16.596 | בעידניה |  |
| 2:16.597 | סגיד |  |
| 2:16.598 | ליה | , |
| 2:16.599 | ר"ל | : |
| 2:16.600 | אדם |  |
| 2:16.601 | פחות |  |
| 2:16.602 | ובזוי |  |
| 2:16.603 | כשועל | , |
| 2:16.604 | והשעה |  |
| 2:16.605 | עומדת |  |
| 2:16.606 | לו |  |
| 2:16.607 | באיזה |  |
| 2:16.608 | משרה |  |
| 2:16.609 | שבידו | , |
| 2:16.610 | סגיד |  |
| 2:16.611 | ליה | . |
| 2:16.612 | וכן |  |
| 2:16.613 | הוא |  |
| 2:16.614 | נהוג |  |
| 2:16.615 | אצלנו |  |
| 2:16.616 | ואצל |  |
| 2:16.617 | הערביים |  |
| 2:16.618 | להשתחוות |  |
| 2:16.619 | ולנשק |  |
| 2:16.620 | אדם |  |
| 2:16.621 | גדול |  |
| 2:16.622 | המעלה | , |
| 2:16.623 | ואין |  |
| 2:16.624 | בזה |  |
| 2:16.625 | שום |  |
| 2:16.626 | איסור | , |
| 2:16.627 | שאינו |  |
| 2:16.628 | מקבלו |  |
| 2:16.629 | עליו |  |
| 2:16.630 | באלוה | , |
| 2:16.631 | רק |  |
| 2:16.632 | לאות |  |
| 2:16.633 | כבוד |  |
| 2:16.634 | הוא |  |
| 2:16.635 | משתחוה |  |
| 2:16.636 | לו |  |
| 2:16.637 | ומנשק |  |
| 2:16.638 | את |  |
| 2:16.639 | רגליו | . |


### Translation under review (2:16)

English: And as for what you wrote: "And more than this, what he wrote about the glory of the LORD that it is created light, if so, how did Israel go out at the entrance of their tents and bow down to the created thing?" Not to the visible glory do they bow down, but only to him who made His glory dwell there. And the visible glory is there as a sign and a token for them, for the Name, may He be blessed, desired the tent of meeting as His dwelling, just as we bow down and pray to the Name, may He be blessed, toward the Temple and Jerusalem, For the LORD chose Zion; He desired it as His dwelling. And the Rambam had already explained in chapter 64 of the first, that the glory of the LORD sometimes denotes the created light, which the Holy One makes dwell in the place, as "and the glory of the LORD dwelt on Mount Sinai, and covered it, etc." and sometimes by it they mean His essence, may He be blessed, and His reality, as in "show me, please, Your glory," and the answer came: "for no man shall see Me and live," which shows that the glory spoken of here is His very essence, may He be blessed, and says "Your glory," to magnify, etc. And similarly we find in the words of our Rabbis of blessed memory that the name Elohim is shared by judges and by angels, and by the Name, may He be blessed. Also, the name Shekhinah is shared with what I mentioned above. And I saw fit to alert you here. that there is no prohibition on bowing down to a man, so long as his intent is not to accept him as God, and he is not worshipped, but only submits to him and humbles himself before him, or by way of acknowledgment for a kindness that he did for him. For indeed Abraham our father bowed down to the sons of Heth for the kindness they did with him, to give him a burial place. And so Rabbi Saadia Gaon rendered: "and bowed down to the people of the land, to the sons of Heth, and prostrated to the tent of the land that He created." And Joshua bowed down to the angel. And Nathan the prophet and Bathsheba bowed down to David. And we find much in the Writings that they bowed down to the created being. And we read in Sanhedrin, chapter Four Deaths (Sanhedrin 61b): It was stated: "One who worships idolatry out of love or out of fear, Abaye says: he is liable, for he worshipped it. Rava says: he is exempt. If he accepted it upon himself as God, yes; if not, no. And Abaye said: How do I know this? for we learned: 'One who worships idolatry: one who worships, and one who sacrifices, and one who burns incense, and one who pours libations, and one who bows down, and one who accepts it upon himself as God, and one who says to him: You are my God.' What? Is it not that 'one who worships' means out of love or out of fear? And Rava says to you: No, but as Rabbi Yirmiyah resolves it: (Rashi explains: one who worships every worship that is customary for it, and one who sacrifices, etc., not in its customary way, and in every case the first clause and the last are for the sake of divinity). Abaye said: How do I know this? for it was taught: 'You shall not bow down to them; to them you do not bow down, but you bow down to a man like yourself.' One might think: even to worship like Haman? The verse says: 'and you shall not worship them.' But Haman was worshipped out of fear! And Rava? Like Haman, and not like Haman. Like Haman, for he himself is idolatry; and not like Haman, for in the case of Haman it was out of fear, and here it is not out of fear. And Abaye said: How do I know this? for it was taught: 'An anointed priest in idolatry. Rabbi says: in an unwitting act, and the Sages say: in a concealed matter, and they agree that for a he-goat it is like an individual, and they agree that he does not bring a suspended guilt offering.' What is this unwitting act like? If he holds that it is a synagogue and bowed down to it, then his heart is with Heaven. But suppose he saw a statue and bowed down to it, (explanation: that he saw the figure of the king, which they were accustomed to make in the likeness of the king, and one who saw it bows down to it in honor of the king, and sometimes they worship it, and this one saw one who was worshipped and bowed down to it). If he accepted it upon himself as God, it is deliberate; and if he did not accept it upon himself as God, it is nothing. (Rashi explains: but he bowed down to him in honor of the king; it is nothing, for he did not intend idolatry). But is it not out of love and out of fear? And Rava says to you: In the case of one who says 'permitted.' Does 'says permitted' mean concealment of a matter? In the case of one who says permitted entirely; concealment of a matter is upholding part and nullifying part, "Thus far the Gemara, with some of Rashi's explanation." And a matter is learned: that to every created being it is forbidden to bow down, except a man, and this is that he be not worshipped like Haman, who made himself idolatry, (as we said in Megillah 19), for otherwise Mordecai would not have refrained from bowing down to him, as it is taught in the baraita: "You shall not bow down to them; to them you do not bow down, but you bow down to a man like yourself," and this is that he be not worshipped like Haman, who made himself God, but only bows down to him as a sign of honor, and not for divinity, and therefore Joshua bowed down to the angel, since he appeared to him in the form of a man, as it is written: "And behold, a man stood before him, and therefore he bowed down to him as a sign of honor and not to worship him." And see in Sefer ha-Ikkarim, Treatise 2, chapter 28, where he elaborated on the verses that call the glory that appears to the prophets by the name YHVH, and also the ark, as as it says: "And it was when the ark set out, etc. Arise, LORD, and when it rested he said: Return, LORD." And Joshua called the ark "Lord of all the earth"; see there. And in his view, that Joshua bowed down to the angel, because that bowing was in honor of the Name, may He be blessed, and since this one was His messenger, etc.; see there. And our Rabbis of blessed memory said: "The fox at its time bows down to him, that is to say: a man lowly and despised like a fox, when the moment stands for him in some office in his hand, he bows down to him." And so it is customary with us and with the Arabs to bow down and kiss a man of great stature, and there is no prohibition in this, for he does not accept him as God; it is only as a sign of honor that he bows down to him and kisses his feet.

Literal: and-what you-wrote and-more than-this what that-he-wrote about the-glory-of the-LORD that-it-is light created if so how did Israel go-out at-the-entrance-of their-tents and-bow-down to-the-created-thing not to-the-glory-of the-visible they-bow-down but-only to-him-who made-dwell His-glory there and-the-glory visible there as-a-sign and-a-token for-them for the-Name,-may-He-be-blessed desired to-the-tent-of meeting as His-dwelling just-as we bow-down and-pray to-the-Name,-may-He-be-blessed toward the-Temple and-Jerusalem for chose the-LORD Zion desired as His-dwelling and-already explained the-Rambam in-chapter 64 of-the-first that the-glory-of the-LORD sometimes they-mean by-it the-light that-is-created which makes-dwell the-Holy-One in-the-place as and-dwelt the-glory-of the-LORD on the-mountain Sinai and-covered-it etc. and-sometimes they-mean by-it His-essence may-He-be-blessed and-His-reality as-in show-me please (object-marker) Your-glory and-came the-answer that no shall-see-me man and-live shows that the-glory spoken-of here itself may-He-be-blessed and-says Your-glory to-magnify etc. and-similarly in-this we-find in-the-words-of our-Rabbis-of-blessed-memory that-the-name Elohim is-shared by-judges and-by-angels and-by-the-Name,-may-He-be-blessed also the-name Shekhinah is-shared with-what I-mentioned above and-I-saw-fit to-alert-you here that-there-is-not a-prohibition to-bow-down to-a-man so-long-as there-is-not his-intent to-accept-him as-God and-also he-is-not worshipped but-only submits to-him and-humbles himself before-him or by-way-of acknowledgment for a-kindness that-he-did for-him for-indeed Abraham our-father bowed-down to-the-sons-of Heth for the-kindness that-they-did with-him to-give him a-possession-of a-burial-place and-so rendered Rabbi Saadia Gaon and-bowed-down to-the-people-of the-land to-the-sons-of Heth and-prostrated to-the-tent-of the-land that-He-created and-Joshua bowed-down to-the-angel and-Nathan the-prophet and-Bath- sheba bowed-down to-David and-much we-find in-the-Writings that-they-bowed-down to-the-created-being and-we-read in-Sanhedrin chapter Four Deaths Sanhedrin 61 b it-was-stated one-who-worships idolatry out-of-love or-out-of-fear Abaye says liable for worshipped-it Rava says exempt if he-accepted upon-himself as-God yes if not no and-said Abaye from-where do-I-say it as-we-learned one-who-worships idolatry one-who worships and-one-who sacrifices and-one-who burns-incense and-one-who pours-libations and-one-who bows-down and-one-who accepts-it upon-himself as-God and-one-who-says to-him my-God you what is-it-not one-who worships out-of-love or out-of-fear and-Rava said to-you no but as-he-resolves Rabbi Yirmiyah (explains Rashi: one-who worships every worship that-is customary for-it and-one-who sacrifices etc. not in-its-customary-way and-always the-first-clause and-the-last for-the-sake-of divinity said Abaye from-where do-I-say it for-it-was-taught you-shall-not bow-down to-them to-them not you bow-down but you bow-down to-a-man like-yourself one-might-think even to-worship like-Haman the-verse-says and-not you-shall-worship-them but Haman out-of-fear he was-worshipped and-Rava like-Haman and-not like-Haman like-Haman for-he himself idolatry and-not like-Haman for Haman out-of-fear and-here it-is-not out-of-fear and Abaye-said from-where do-I-say it for-it-was-taught a-priest anointed in-idolatry Rabbi says in-an-unwitting act and-the-Sages say in-a-concealed matter and-they-agree that-for-a-he-goat it-is-like-an-individual and-they-agree that-he-does-not bring a-guilt-offering suspended this unwitting act what-is it if he-holds a-synagogue and-bowed-down to-it then his-heart is-with-Heaven but he-saw a-statue and-bowed-down to-it (explanation that-he-saw the-figure-of the-king which-they-were accustomed to-make in-the-likeness-of the-king and-one-who-saw it bows-down to-it in-honor-of the-king, and-sometimes they-worship it and-this-one saw one who-was-worshipped and-bowed-down to-it if he-accepted upon-himself as-God deliberate it-is and-if not he-accepted upon-himself as-God it-is-not anything it-is (explains Rashi: but in-honor-of the-king he-bowed-down to-him it-is-not anything it-is for not he-intended idolatry but is-it-not out-of-love and-out-of-fear and-Rava said to-you in-the-case-of-one-who-says permitted says permitted does-it-mean concealment of-a-matter in-the-case-of-one-who-says permitted entirely concealment of-a-matter upholding a-part and-nullifying a-part thus-far the-Gemara with some of-the-explanation of-Rashi and-a-matter is-learned that-to-every created-being it-is-forbidden to-bow-down except a-man and-this-is that-not be-worshipped like-Haman who-made himself idolatry as-we-said in-Megillah 19 for-otherwise if-not so not would-have-been Mordecai refrained from-bowing-down to-him as-it-is-taught in-the-baraita you-shall-not bow-down to-them to-them not you bow-down but you bow-down to-a-man like-yourself and-this that-not be worshipped like-Haman who-made himself God but-only bows-down to-him as-a-sign of-honor and-not for-divinity and-therefore bowed-down Joshua to-the-angel since appeared to-him in-the-form-of a-man as-it-is-written and-behold a-man stood before-him and-therefore bowed-down to-him as-a-sign of-honor and-not to-worship-him and-see in-Sefer ha-Ikkarim Treatise 2 chapter 28 where-he-elaborated on-the-explanation-of the-verses that-call the-glory that-appears to-the-prophets by-the-name YHVH and-also the-ark as as-it-says and-it-was when-set-out the-ark etc. Arise LORD and-when-it-rested he-said Return LORD and-Joshua called the-ark Lord of-all the-earth see-there and-in-his-view that he-bowed-down Joshua to-the-angel because the-bowing that was in-honor-of the-Name,-may-He-be-blessed and-since this-one is-His-messenger etc. see-there and-our-Rabbis-of-blessed-memory said the-fox at-its-time bows-down to-him that-is-to-say a-man lowly and-despised like-a-fox and-the-moment stands for-him in-some office that-is-in-his-hand he-bows-down to-him and-so it-is customary with-us and-with the-Arabs to-bow-down and-to-kiss a-man of-great stature and-there-is-not in-this any prohibition for-he-does-not accept-him upon-himself as-God only as-a-sign of-honor it-is he-bows-down to-him and-kisses (object-marker) his-feet

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:16.1 | and what |  |  |  |
| 2:16.2 | you wrote |  |  |  |
| 2:16.3 | and more |  |  |  |
| 2:16.4 | than this |  |  |  |
| 2:16.5 | what |  |  |  |
| 2:16.6 | that he wrote |  |  | שכ' abbreviates שכתב, 'that he wrote'; the questioner's source is not named. |
| 2:16.7 | about |  |  |  |
| 2:16.8 | the glory of |  |  |  |
| 2:16.9 | the LORD |  |  | ה' is the Tetragrammaton as written: 'the LORD' (decisions, Names and titles). |
| 2:16.10 | that it is |  |  |  |
| 2:16.11 | light |  |  |  |
| 2:16.12 | created |  |  |  |
| 2:16.13 | if |  |  |  |
| 2:16.14 | so |  |  |  |
| 2:16.15 | how |  |  |  |
| 2:16.16 | did |  |  |  |
| 2:16.17 | Israel |  |  |  |
| 2:16.18 | go out |  |  |  |
| 2:16.19 | at the entrance of |  |  |  |
| 2:16.20 | their tents |  |  |  |
| 2:16.21 | and bow down |  |  |  |
| 2:16.22 | to the created thing |  |  |  |
| 2:16.23 | not |  |  |  |
| 2:16.24 | to the glory of |  |  |  |
| 2:16.25 | the visible |  |  |  |
| 2:16.26 | they bow down |  |  |  |
| 2:16.27 | but only |  |  |  |
| 2:16.28 | to him who |  |  |  |
| 2:16.29 | made dwell |  |  |  |
| 2:16.30 | His glory |  |  |  |
| 2:16.31 | there |  |  |  |
| 2:16.32 | and the glory |  |  |  |
| 2:16.33 | visible |  |  |  |
| 2:16.34 | there |  |  |  |
| 2:16.35 | as a sign |  |  |  |
| 2:16.36 | and a token |  |  |  |
| 2:16.37 | for them |  |  |  |
| 2:16.38 | for |  |  |  |
| 2:16.39 | the Name, may He be blessed |  |  | השי"ת: 'the Name, may He be blessed' (owner's default, Q3). |
| 2:16.40 | desired |  |  |  |
| 2:16.41 | to the tent of |  |  |  |
| 2:16.42 | meeting |  |  |  |
| 2:16.43 | as |  |  |  |
| 2:16.44 | His dwelling |  |  |  |
| 2:16.45 | just as |  |  |  |
| 2:16.46 | we |  |  |  |
| 2:16.47 | bow down |  |  |  |
| 2:16.48 | and pray |  |  |  |
| 2:16.49 | to the Name, may He be blessed |  |  |  |
| 2:16.50 | toward |  |  |  |
| 2:16.51 | the Temple |  |  |  |
| 2:16.52 | and Jerusalem |  |  |  |
| 2:16.53 | for |  |  |  |
| 2:16.54 | chose |  |  |  |
| 2:16.55 | the LORD |  |  |  |
| 2:16.56 | Zion |  |  | בציון is the verse's 'in Zion', rendered as the object of 'chose'. |
| 2:16.57 | desired |  |  |  |
| 2:16.58 | as |  |  |  |
| 2:16.59 | His dwelling |  |  |  |
| 2:16.60 | and already |  |  |  |
| 2:16.61 | explained |  |  |  |
| 2:16.62 | the Rambam |  |  |  |
| 2:16.63 | in chapter |  |  | בפ' is בפרק: 'in chapter'. |
| 2:16.64 | 64 |  |  | ס"ד is 64 in letter numerals. |
| 2:16.65 | of the first |  |  | מראשון: 'of the first'. The author names no book, so none is added. |
| 2:16.66 | that |  |  |  |
| 2:16.67 | the glory of |  |  |  |
| 2:16.68 | the LORD |  |  |  |
| 2:16.69 | sometimes |  |  |  |
| 2:16.70 | they mean |  |  |  |
| 2:16.71 | by it |  |  |  |
| 2:16.72 | the light |  |  |  |
| 2:16.73 | that is created |  |  |  |
| 2:16.74 | which |  |  |  |
| 2:16.75 | makes dwell |  |  |  |
| 2:16.76 | the Holy One |  |  | הק' abbreviates הקדוש: 'the Holy One'; the author gives no 'blessed be He' here. |
| 2:16.77 | in the place |  |  |  |
| 2:16.78 | as |  |  |  |
| 2:16.79 | and dwelt |  |  |  |
| 2:16.80 | the glory of |  |  |  |
| 2:16.81 | the LORD |  |  |  |
| 2:16.82 | on |  |  |  |
| 2:16.83 | the mountain |  |  |  |
| 2:16.84 | Sinai |  |  |  |
| 2:16.85 | and covered it |  |  |  |
| 2:16.86 | etc. |  |  |  |
| 2:16.87 | and sometimes |  |  |  |
| 2:16.88 | they mean |  |  |  |
| 2:16.89 | by it |  |  |  |
| 2:16.90 | His essence |  |  |  |
| 2:16.91 | may He be blessed |  |  | ית' (יתברך) follows 'His essence', where the Hebrew places it. |
| 2:16.92 | and His reality |  |  |  |
| 2:16.93 | as in |  |  |  |
| 2:16.94 | show me |  |  |  |
| 2:16.95 | please |  |  |  |
| 2:16.96 | (object marker) |  |  |  |
| 2:16.97 | Your glory |  |  |  |
| 2:16.98 | and came |  |  |  |
| 2:16.99 | the answer |  |  |  |
| 2:16.100 | that |  |  |  |
| 2:16.101 | no |  |  |  |
| 2:16.102 | shall see me |  |  |  |
| 2:16.103 | man |  |  |  |
| 2:16.104 | and live |  |  |  |
| 2:16.105 | shows |  |  |  |
| 2:16.106 | that |  |  |  |
| 2:16.107 | the glory |  |  |  |
| 2:16.108 | spoken of |  |  |  |
| 2:16.109 | here |  |  |  |
| 2:16.110 | itself |  |  |  |
| 2:16.111 | may He be blessed |  |  |  |
| 2:16.112 | and says |  |  |  |
| 2:16.113 | Your glory |  |  |  |
| 2:16.114 | to magnify |  |  |  |
| 2:16.115 | etc. |  |  |  |
| 2:16.116 | and similarly |  |  |  |
| 2:16.117 | in this |  |  |  |
| 2:16.118 | we find |  |  |  |
| 2:16.119 | in the words of |  |  |  |
| 2:16.120 | our Rabbis of blessed memory |  |  | רז"ל is רבותינו זכרונם לברכה: 'our Rabbis of blessed memory'. |
| 2:16.121 | that the name |  |  |  |
| 2:16.122 | Elohim |  |  | אלהים is discussed as a word shared by three referents, so it is kept as 'Elohim' and not rendered 'God' or 'gods'. |
| 2:16.123 | is shared |  |  |  |
| 2:16.124 | by judges |  |  |  |
| 2:16.125 | and by angels |  |  |  |
| 2:16.126 | and by the Name, may He be blessed |  |  | ולשי"ת: 'and by the Name, may He be blessed', the third referent of the shared word. |
| 2:16.127 | also |  |  |  |
| 2:16.128 | the name |  |  |  |
| 2:16.129 | Shekhinah |  |  |  |
| 2:16.130 | is shared |  |  |  |
| 2:16.131 | with what |  |  |  |
| 2:16.132 | I mentioned |  |  |  |
| 2:16.133 | above |  |  |  |
| 2:16.134 | and I saw fit |  |  |  |
| 2:16.135 | to alert you |  |  |  |
| 2:16.136 | here |  |  |  |
| 2:16.137 | that there is not |  |  |  |
| 2:16.138 | a prohibition |  |  |  |
| 2:16.139 | to bow down |  |  |  |
| 2:16.140 | to a man |  |  |  |
| 2:16.141 | so long as |  |  |  |
| 2:16.142 | there is not |  |  |  |
| 2:16.143 | his intent |  |  |  |
| 2:16.144 | to accept him |  |  |  |
| 2:16.145 | as God |  |  |  |
| 2:16.146 | and also |  |  |  |
| 2:16.147 | he is not |  |  |  |
| 2:16.148 | worshipped |  |  |  |
| 2:16.149 | but only |  |  |  |
| 2:16.150 | submits |  |  |  |
| 2:16.151 | to him |  |  |  |
| 2:16.152 | and humbles |  |  |  |
| 2:16.153 | himself |  |  |  |
| 2:16.154 | before him |  |  |  |
| 2:16.155 | or |  |  |  |
| 2:16.156 | by way of |  |  |  |
| 2:16.157 | acknowledgment |  |  |  |
| 2:16.158 | for |  |  |  |
| 2:16.159 | a kindness |  |  |  |
| 2:16.160 | that he did |  |  |  |
| 2:16.161 | for him |  |  |  |
| 2:16.162 | for indeed |  |  |  |
| 2:16.163 | Abraham |  |  |  |
| 2:16.164 | our father |  |  |  |
| 2:16.165 | bowed down |  |  |  |
| 2:16.166 | to the sons of |  |  |  |
| 2:16.167 | Heth |  |  |  |
| 2:16.168 | for |  |  |  |
| 2:16.169 | the kindness |  |  |  |
| 2:16.170 | that they did |  |  |  |
| 2:16.171 | with him |  |  |  |
| 2:16.172 | to give |  |  |  |
| 2:16.173 | him |  |  |  |
| 2:16.174 | a possession of |  |  |  |
| 2:16.175 | a burial place |  |  |  |
| 2:16.176 | and so |  |  |  |
| 2:16.177 | rendered |  |  |  |
| 2:16.178 | Rabbi |  |  |  |
| 2:16.179 | Saadia |  |  |  |
| 2:16.180 | Gaon |  |  |  |
| 2:16.181 | and bowed down |  |  |  |
| 2:16.182 | to the people of |  |  |  |
| 2:16.183 | the land |  |  |  |
| 2:16.184 | to the sons of |  |  |  |
| 2:16.185 | Heth |  |  |  |
| 2:16.186 | and prostrated |  |  |  |
| 2:16.187 | to the tent of |  |  |  |
| 2:16.188 | the land |  |  | אלבלד is Arabic al-balad, 'the land', in a Judeo-Arabic rendering written in Hebrew letters. |
| 2:16.189 | that He created |  |  | שברא: 'that He created'. The subject of the verb is not marked in the Judeo-Arabic; see apparatus. |
| 2:16.190 | and Joshua |  |  |  |
| 2:16.191 | bowed down |  |  |  |
| 2:16.192 | to the angel |  |  |  |
| 2:16.193 | and Nathan |  |  |  |
| 2:16.194 | the prophet |  |  |  |
| 2:16.195 | and Bath- |  |  |  |
| 2:16.196 | sheba |  |  |  |
| 2:16.197 | bowed down |  |  |  |
| 2:16.198 | to David |  |  |  |
| 2:16.199 | and much |  |  |  |
| 2:16.200 | we find |  |  |  |
| 2:16.201 | in the Writings |  |  |  |
| 2:16.202 | that they bowed down |  |  |  |
| 2:16.203 | to the created being |  |  |  |
| 2:16.204 | and we read |  |  |  |
| 2:16.205 | in Sanhedrin |  |  |  |
| 2:16.206 | chapter |  |  |  |
| 2:16.207 | Four |  |  |  |
| 2:16.208 | Deaths |  |  |  |
| 2:16.209 | Sanhedrin |  |  |  |
| 2:16.210 | 61 |  |  | ס"א is 61 (the folio). |
| 2:16.211 | b |  |  | ע"ב is the side b of the folio: 'Sanhedrin 61b'. |
| 2:16.212 | it was stated |  |  |  |
| 2:16.213 | one who worships |  |  |  |
| 2:16.214 | idolatry |  |  | ע"ז is עבודה זרה, 'strange worship'; here 'idolatry' (decisions). |
| 2:16.215 | out of love |  |  |  |
| 2:16.216 | or out of fear |  |  |  |
| 2:16.217 | Abaye |  |  |  |
| 2:16.218 | says |  |  |  |
| 2:16.219 | liable |  |  |  |
| 2:16.220 | for |  |  |  |
| 2:16.221 | worshipped it |  |  | פלחה is Aramaic for 'he worshipped it'. |
| 2:16.222 | Rava |  |  |  |
| 2:16.223 | says |  |  |  |
| 2:16.224 | exempt |  |  |  |
| 2:16.225 | if |  |  |  |
| 2:16.226 | he accepted |  |  |  |
| 2:16.227 | upon himself |  |  |  |
| 2:16.228 | as God |  |  | באלוה: 'as God'. Under the doubtful-case rule the singular is 'God'; 'as a god' is the alternative. |
| 2:16.229 | yes |  |  |  |
| 2:16.230 | if |  |  |  |
| 2:16.231 | not |  |  |  |
| 2:16.232 | no |  |  |  |
| 2:16.233 | and said |  |  |  |
| 2:16.234 | Abaye |  |  |  |
| 2:16.235 | from where |  |  |  |
| 2:16.236 | do I say |  |  |  |
| 2:16.237 | it |  |  |  |
| 2:16.238 | as we learned |  |  |  |
| 2:16.239 | one who worships |  |  |  |
| 2:16.240 | idolatry |  |  |  |
| 2:16.241 | one who |  |  |  |
| 2:16.242 | worships |  |  |  |
| 2:16.243 | and one who |  |  |  |
| 2:16.244 | sacrifices |  |  |  |
| 2:16.245 | and one who |  |  |  |
| 2:16.246 | burns incense |  |  |  |
| 2:16.247 | and one who |  |  |  |
| 2:16.248 | pours libations |  |  |  |
| 2:16.249 | and one who |  |  |  |
| 2:16.250 | bows down |  |  |  |
| 2:16.251 | and one who |  |  |  |
| 2:16.252 | accepts it |  |  |  |
| 2:16.253 | upon himself |  |  |  |
| 2:16.254 | as God |  |  |  |
| 2:16.255 | and one who says |  |  |  |
| 2:16.256 | to him |  |  |  |
| 2:16.257 | my God |  |  |  |
| 2:16.258 | you |  |  |  |
| 2:16.259 | what |  |  |  |
| 2:16.260 | is it not |  |  |  |
| 2:16.261 | one who |  |  |  |
| 2:16.262 | worships |  |  |  |
| 2:16.263 | out of love |  |  |  |
| 2:16.264 | or |  |  |  |
| 2:16.265 | out of fear |  |  |  |
| 2:16.266 | and Rava |  |  |  |
| 2:16.267 | said |  |  |  |
| 2:16.268 | to you |  |  |  |
| 2:16.269 | no |  |  |  |
| 2:16.270 | but |  |  |  |
| 2:16.271 | as he resolves |  |  |  |
| 2:16.272 | Rabbi |  |  |  |
| 2:16.273 | Yirmiyah |  |  |  |
| 2:16.274 | (explains |  |  |  |
| 2:16.275 | Rashi: |  |  |  |
| 2:16.276 | one who |  |  |  |
| 2:16.277 | worships |  |  |  |
| 2:16.278 | every |  |  |  |
| 2:16.279 | worship |  |  |  |
| 2:16.280 | that is |  |  |  |
| 2:16.281 | customary |  |  |  |
| 2:16.282 | for it |  |  |  |
| 2:16.283 | and one who |  |  |  |
| 2:16.284 | sacrifices |  |  |  |
| 2:16.285 | etc. |  |  |  |
| 2:16.286 | not |  |  |  |
| 2:16.287 | in its customary way |  |  |  |
| 2:16.288 | and always |  |  |  |
| 2:16.289 | the first clause |  |  |  |
| 2:16.290 | and the last |  |  |  |
| 2:16.291 | for the sake of |  |  |  |
| 2:16.292 | divinity |  |  |  |
| 2:16.293 | said |  |  |  |
| 2:16.294 | Abaye |  |  |  |
| 2:16.295 | from where |  |  |  |
| 2:16.296 | do I say |  |  |  |
| 2:16.297 | it |  |  |  |
| 2:16.298 | for it was taught |  |  |  |
| 2:16.299 | you shall not |  |  |  |
| 2:16.300 | bow down |  |  |  |
| 2:16.301 | to them |  |  |  |
| 2:16.302 | to them |  |  |  |
| 2:16.303 | not |  |  |  |
| 2:16.304 | you |  |  |  |
| 2:16.305 | bow down |  |  |  |
| 2:16.306 | but |  |  |  |
| 2:16.307 | you |  |  |  |
| 2:16.308 | bow down |  |  |  |
| 2:16.309 | to a man |  |  |  |
| 2:16.310 | like yourself |  |  |  |
| 2:16.311 | one might think |  |  |  |
| 2:16.312 | even |  |  |  |
| 2:16.313 | to worship |  |  |  |
| 2:16.314 | like Haman |  |  |  |
| 2:16.315 | the verse says |  |  |  |
| 2:16.316 | and not |  |  |  |
| 2:16.317 | you shall worship them |  |  |  |
| 2:16.318 | but |  |  |  |
| 2:16.319 | Haman |  |  |  |
| 2:16.320 | out of fear |  |  |  |
| 2:16.321 | he |  |  |  |
| 2:16.322 | was worshipped |  |  |  |
| 2:16.323 | and Rava |  |  |  |
| 2:16.324 | like Haman |  |  |  |
| 2:16.325 | and not |  |  |  |
| 2:16.326 | like Haman |  |  |  |
| 2:16.327 | like Haman |  |  |  |
| 2:16.328 | for he |  |  |  |
| 2:16.329 | himself |  |  |  |
| 2:16.330 | idolatry |  |  |  |
| 2:16.331 | and not |  |  |  |
| 2:16.332 | like Haman |  |  |  |
| 2:16.333 | for |  |  |  |
| 2:16.334 | Haman |  |  |  |
| 2:16.335 | out of fear |  |  |  |
| 2:16.336 | and here |  |  |  |
| 2:16.337 | it is not |  |  |  |
| 2:16.338 | out of fear |  |  |  |
| 2:16.339 | and |  |  |  |
| 2:16.340 | Abaye said |  |  |  |
| 2:16.341 | from where |  |  |  |
| 2:16.342 | do I say |  |  |  |
| 2:16.343 | it |  |  |  |
| 2:16.344 | for it was taught |  |  |  |
| 2:16.345 | a priest |  |  | כהן משיח is the anointed priest (Lev 4:3). |
| 2:16.346 | anointed |  |  |  |
| 2:16.347 | in idolatry |  |  |  |
| 2:16.348 | Rabbi |  |  |  |
| 2:16.349 | says |  |  |  |
| 2:16.350 | in an unwitting |  |  |  |
| 2:16.351 | act |  |  |  |
| 2:16.352 | and the Sages |  |  |  |
| 2:16.353 | say |  |  |  |
| 2:16.354 | in a concealed |  |  |  |
| 2:16.355 | matter |  |  |  |
| 2:16.356 | and they agree |  |  |  |
| 2:16.357 | that for a he-goat |  |  |  |
| 2:16.358 | it is like an individual |  |  |  |
| 2:16.359 | and they agree |  |  |  |
| 2:16.360 | that he does not |  |  |  |
| 2:16.361 | bring |  |  |  |
| 2:16.362 | a guilt offering |  |  |  |
| 2:16.363 | suspended |  |  |  |
| 2:16.364 | this |  |  |  |
| 2:16.365 | unwitting |  |  |  |
| 2:16.366 | act |  |  |  |
| 2:16.367 | what is |  |  |  |
| 2:16.368 | it |  |  |  |
| 2:16.369 | if |  |  |  |
| 2:16.370 | he holds |  |  |  |
| 2:16.371 | a synagogue |  |  | ביהכנ"ס is בית הכנסת: 'synagogue'. |
| 2:16.372 | and bowed down |  |  |  |
| 2:16.373 | to it |  |  |  |
| 2:16.374 | then |  |  |  |
| 2:16.375 | his heart |  |  |  |
| 2:16.376 | is with Heaven |  |  |  |
| 2:16.377 | but |  |  | אנדרטא: 'a statue'. The decisions give no ruling; the author's Rashi gloss below reads 'a figure of the king'. |
| 2:16.378 | he saw |  |  |  |
| 2:16.379 | a statue |  |  |  |
| 2:16.380 | and bowed down |  |  |  |
| 2:16.381 | to it |  |  |  |
| 2:16.382 | (explanation |  |  |  |
| 2:16.383 | that he saw |  |  |  |
| 2:16.384 | the figure of |  |  |  |
| 2:16.385 | the king |  |  |  |
| 2:16.386 | which they were |  |  |  |
| 2:16.387 | accustomed |  |  |  |
| 2:16.388 | to make |  |  |  |
| 2:16.389 | in the likeness of |  |  |  |
| 2:16.390 | the king |  |  |  |
| 2:16.391 | and one who saw |  |  |  |
| 2:16.392 | it |  |  |  |
| 2:16.393 | bows down |  |  |  |
| 2:16.394 | to it |  |  |  |
| 2:16.395 | in honor of |  |  |  |
| 2:16.396 | the king, |  |  |  |
| 2:16.397 | and sometimes |  |  |  |
| 2:16.398 | they worship |  |  |  |
| 2:16.399 | it |  |  |  |
| 2:16.400 | and this one |  |  |  |
| 2:16.401 | saw |  |  |  |
| 2:16.402 | one |  |  |  |
| 2:16.403 | who was worshipped |  |  |  |
| 2:16.404 | and bowed down |  |  |  |
| 2:16.405 | to it |  |  |  |
| 2:16.406 | if |  |  |  |
| 2:16.407 | he accepted |  |  |  |
| 2:16.408 | upon himself |  |  |  |
| 2:16.409 | as God |  |  |  |
| 2:16.410 | deliberate |  |  |  |
| 2:16.411 | it is |  |  |  |
| 2:16.412 | and if |  |  |  |
| 2:16.413 | not |  |  |  |
| 2:16.414 | he accepted |  |  |  |
| 2:16.415 | upon himself |  |  |  |
| 2:16.416 | as God |  |  |  |
| 2:16.417 | it is not |  |  |  |
| 2:16.418 | anything |  |  |  |
| 2:16.419 | it is |  |  |  |
| 2:16.420 | (explains |  |  |  |
| 2:16.421 | Rashi: |  |  |  |
| 2:16.422 | but |  |  |  |
| 2:16.423 | in honor of |  |  |  |
| 2:16.424 | the king |  |  |  |
| 2:16.425 | he bowed down |  |  |  |
| 2:16.426 | to him |  |  |  |
| 2:16.427 | it is not |  |  |  |
| 2:16.428 | anything |  |  |  |
| 2:16.429 | it is |  |  |  |
| 2:16.430 | for |  |  |  |
| 2:16.431 | not |  |  |  |
| 2:16.432 | he intended |  |  |  |
| 2:16.433 | idolatry |  |  |  |
| 2:16.434 | but |  |  |  |
| 2:16.435 | is it not |  |  |  |
| 2:16.436 | out of love |  |  |  |
| 2:16.437 | and out of fear |  |  |  |
| 2:16.438 | and Rava |  |  |  |
| 2:16.439 | said |  |  |  |
| 2:16.440 | to you |  |  |  |
| 2:16.441 | in the case of one who says |  |  |  |
| 2:16.442 | permitted |  |  |  |
| 2:16.443 | says |  |  |  |
| 2:16.444 | permitted |  |  |  |
| 2:16.445 | does it mean |  |  |  |
| 2:16.446 | concealment |  |  |  |
| 2:16.447 | of a matter |  |  |  |
| 2:16.448 | in the case of one who says |  |  |  |
| 2:16.449 | permitted |  |  |  |
| 2:16.450 | entirely |  |  |  |
| 2:16.451 | concealment |  |  |  |
| 2:16.452 | of a matter |  |  |  |
| 2:16.453 | upholding |  |  |  |
| 2:16.454 | a part |  |  |  |
| 2:16.455 | and nullifying |  |  |  |
| 2:16.456 | a part |  |  |  |
| 2:16.457 | thus far |  |  |  |
| 2:16.458 | the Gemara |  |  |  |
| 2:16.459 | with |  |  |  |
| 2:16.460 | some |  |  |  |
| 2:16.461 | of the explanation |  |  |  |
| 2:16.462 | of Rashi |  |  |  |
| 2:16.463 | and a matter |  |  |  |
| 2:16.464 | is learned |  |  |  |
| 2:16.465 | that to every |  |  |  |
| 2:16.466 | created being |  |  |  |
| 2:16.467 | it is forbidden |  |  |  |
| 2:16.468 | to bow down |  |  |  |
| 2:16.469 | except |  |  |  |
| 2:16.470 | a man |  |  |  |
| 2:16.471 | and this is |  |  |  |
| 2:16.472 | that not |  |  |  |
| 2:16.473 | be worshipped |  |  |  |
| 2:16.474 | like Haman |  |  |  |
| 2:16.475 | who made |  |  |  |
| 2:16.476 | himself |  |  |  |
| 2:16.477 | idolatry |  |  |  |
| 2:16.478 | as we said |  |  |  |
| 2:16.479 | in Megillah |  |  |  |
| 2:16.480 | 19 |  |  | יט is 19: Megillah 19, as the author gives it. |
| 2:16.481 | for otherwise |  |  |  |
| 2:16.482 | if not |  |  |  |
| 2:16.483 | so |  |  |  |
| 2:16.484 | not |  |  |  |
| 2:16.485 | would have been |  |  |  |
| 2:16.486 | Mordecai |  |  |  |
| 2:16.487 | refrained |  |  |  |
| 2:16.488 | from bowing down |  |  |  |
| 2:16.489 | to him |  |  |  |
| 2:16.490 | as it is taught |  |  |  |
| 2:16.491 | in the baraita |  |  |  |
| 2:16.492 | you shall not |  |  |  |
| 2:16.493 | bow down |  |  |  |
| 2:16.494 | to them |  |  |  |
| 2:16.495 | to them |  |  |  |
| 2:16.496 | not |  |  |  |
| 2:16.497 | you |  |  |  |
| 2:16.498 | bow down |  |  |  |
| 2:16.499 | but |  |  |  |
| 2:16.500 | you |  |  |  |
| 2:16.501 | bow down |  |  |  |
| 2:16.502 | to a man |  |  |  |
| 2:16.503 | like yourself |  |  |  |
| 2:16.504 | and this |  |  |  |
| 2:16.505 | that not |  |  |  |
| 2:16.506 | be |  |  |  |
| 2:16.507 | worshipped |  |  |  |
| 2:16.508 | like Haman |  |  |  |
| 2:16.509 | who made |  |  |  |
| 2:16.510 | himself |  |  |  |
| 2:16.511 | God |  |  | אלוה (singular, 'who made himself a god') is 'God' under the doubtful-case rule; 'a god' is the alternative. |
| 2:16.512 | but only |  |  |  |
| 2:16.513 | bows down |  |  |  |
| 2:16.514 | to him |  |  |  |
| 2:16.515 | as a sign |  |  |  |
| 2:16.516 | of honor |  |  |  |
| 2:16.517 | and not |  |  |  |
| 2:16.518 | for divinity |  |  |  |
| 2:16.519 | and therefore |  |  |  |
| 2:16.520 | bowed down |  |  |  |
| 2:16.521 | Joshua |  |  |  |
| 2:16.522 | to the angel |  |  |  |
| 2:16.523 | since |  |  |  |
| 2:16.524 | appeared |  |  |  |
| 2:16.525 | to him |  |  |  |
| 2:16.526 | in the form of |  |  |  |
| 2:16.527 | a man |  |  |  |
| 2:16.528 | as it is written |  |  |  |
| 2:16.529 | and behold |  |  |  |
| 2:16.530 | a man |  |  |  |
| 2:16.531 | stood |  |  |  |
| 2:16.532 | before him |  |  |  |
| 2:16.533 | and therefore |  |  |  |
| 2:16.534 | bowed down |  |  |  |
| 2:16.535 | to him |  |  |  |
| 2:16.536 | as a sign |  |  |  |
| 2:16.537 | of honor |  |  |  |
| 2:16.538 | and not |  |  |  |
| 2:16.539 | to worship him |  |  |  |
| 2:16.540 | and see |  |  |  |
| 2:16.541 | in Sefer |  |  |  |
| 2:16.542 | ha-Ikkarim |  |  |  |
| 2:16.543 | Treatise |  |  |  |
| 2:16.544 | 2 |  |  |  |
| 2:16.545 | chapter |  |  |  |
| 2:16.546 | 28 |  |  |  |
| 2:16.547 | where he elaborated |  |  |  |
| 2:16.548 | on the explanation of |  |  |  |
| 2:16.549 | the verses |  |  |  |
| 2:16.550 | that call |  |  |  |
| 2:16.551 | the glory |  |  |  |
| 2:16.552 | that appears |  |  |  |
| 2:16.553 | to the prophets |  |  |  |
| 2:16.554 | by the name |  |  |  |
| 2:16.555 | YHVH |  |  |  |
| 2:16.556 | and also |  |  |  |
| 2:16.557 | the ark |  |  |  |
| 2:16.558 | as |  |  |  |
| 2:16.559 | as it says |  |  | שגאמר is the witness spelling for שנאמר, recorded here and not corrected. |
| 2:16.560 | and it was |  |  |  |
| 2:16.561 | when set out |  |  |  |
| 2:16.562 | the ark |  |  |  |
| 2:16.563 | etc. |  |  |  |
| 2:16.564 | Arise |  |  |  |
| 2:16.565 | LORD |  |  |  |
| 2:16.566 | and when it rested |  |  |  |
| 2:16.567 | he said |  |  |  |
| 2:16.568 | Return |  |  |  |
| 2:16.569 | LORD |  |  |  |
| 2:16.570 | and Joshua |  |  |  |
| 2:16.571 | called |  |  |  |
| 2:16.572 | the ark |  |  |  |
| 2:16.573 | Lord |  |  |  |
| 2:16.574 | of all |  |  |  |
| 2:16.575 | the earth |  |  |  |
| 2:16.576 | see there |  |  | ע"ש is עיין שם: 'see there'. |
| 2:16.577 | and in his view |  |  |  |
| 2:16.578 | that |  |  |  |
| 2:16.579 | he bowed down |  |  |  |
| 2:16.580 | Joshua |  |  |  |
| 2:16.581 | to the angel |  |  |  |
| 2:16.582 | because |  |  |  |
| 2:16.583 | the bowing |  |  |  |
| 2:16.584 | that |  |  |  |
| 2:16.585 | was |  |  |  |
| 2:16.586 | in honor of |  |  |  |
| 2:16.587 | the Name, may He be blessed |  |  |  |
| 2:16.588 | and since |  |  |  |
| 2:16.589 | this one |  |  |  |
| 2:16.590 | is His messenger |  |  |  |
| 2:16.591 | etc. |  |  |  |
| 2:16.592 | see there |  |  | ע"ש is עיין שם: 'see there'. |
| 2:16.593 | and our Rabbis of blessed memory |  |  |  |
| 2:16.594 | said |  |  |  |
| 2:16.595 | the fox |  |  |  |
| 2:16.596 | at its time |  |  |  |
| 2:16.597 | bows down |  |  |  |
| 2:16.598 | to him |  |  |  |
| 2:16.599 | that is to say |  |  | ר"ל is רוצה לומר: 'that is to say'. |
| 2:16.600 | a man |  |  |  |
| 2:16.601 | lowly |  |  |  |
| 2:16.602 | and despised |  |  |  |
| 2:16.603 | like a fox |  |  |  |
| 2:16.604 | and the moment |  |  | והשעה: 'and the moment', the Aramaic of the saying 'when the fox is at its time'. |
| 2:16.605 | stands |  |  |  |
| 2:16.606 | for him |  |  |  |
| 2:16.607 | in some |  |  |  |
| 2:16.608 | office |  |  |  |
| 2:16.609 | that is in his hand |  |  |  |
| 2:16.610 | he bows down |  |  |  |
| 2:16.611 | to him |  |  |  |
| 2:16.612 | and so |  |  |  |
| 2:16.613 | it is |  |  |  |
| 2:16.614 | customary |  |  |  |
| 2:16.615 | with us |  |  |  |
| 2:16.616 | and with |  |  |  |
| 2:16.617 | the Arabs |  |  |  |
| 2:16.618 | to bow down |  |  |  |
| 2:16.619 | and to kiss |  |  |  |
| 2:16.620 | a man |  |  |  |
| 2:16.621 | of great |  |  |  |
| 2:16.622 | stature |  |  |  |
| 2:16.623 | and there is not |  |  |  |
| 2:16.624 | in this |  |  |  |
| 2:16.625 | any |  |  |  |
| 2:16.626 | prohibition |  |  |  |
| 2:16.627 | for he does not |  |  |  |
| 2:16.628 | accept him |  |  |  |
| 2:16.629 | upon himself |  |  |  |
| 2:16.630 | as God |  |  |  |
| 2:16.631 | only |  |  |  |
| 2:16.632 | as a sign |  |  |  |
| 2:16.633 | of honor |  |  |  |
| 2:16.634 | it is |  |  |  |
| 2:16.635 | he bows down |  |  |  |
| 2:16.636 | to him |  |  |  |
| 2:16.637 | and kisses |  |  |  |
| 2:16.638 | (object marker) |  |  |  |
| 2:16.639 | his feet |  |  |  |

Phrases:
- 2:16.1 2:16.2 | And as for what you wrote: | and-what you-wrote
- 2:16.3 2:16.4 2:16.5 2:16.6 2:16.7 2:16.8 2:16.9 | "And more than this, what he wrote about the glory of the LORD | and-more than-this what that-he-wrote about the-glory-of the-LORD
- 2:16.10 2:16.11 2:16.12 | that it is created light, | that-it-is light created
- 2:16.13 2:16.14 2:16.15 | if so, how | if so how
- 2:16.16 2:16.17 2:16.18 2:16.19 2:16.20 | did Israel go out at the entrance of their tents | did Israel go-out at-the-entrance-of their-tents
- 2:16.21 2:16.22 | and bow down to the created thing?" | and-bow-down to-the-created-thing
- 2:16.23 2:16.24 2:16.25 2:16.26 | Not to the visible glory do they bow down, | not to-the-glory-of the-visible they-bow-down
- 2:16.27 2:16.28 2:16.29 2:16.30 2:16.31 | but only to him who made His glory dwell there. | but-only to-him-who made-dwell His-glory there
- 2:16.32 2:16.33 2:16.34 2:16.35 2:16.36 2:16.37 | And the visible glory is there as a sign and a token for them, | and-the-glory visible there as-a-sign and-a-token for-them
- 2:16.38 2:16.39 2:16.40 2:16.41 2:16.42 2:16.43 2:16.44 | for the Name, may He be blessed, desired the tent of meeting as His dwelling, | for the-Name,-may-He-be-blessed desired to-the-tent-of meeting as His-dwelling
- 2:16.45 2:16.46 2:16.47 2:16.48 2:16.49 | just as we bow down and pray to the Name, may He be blessed, | just-as we bow-down and-pray to-the-Name,-may-He-be-blessed
- 2:16.50 2:16.51 2:16.52 | toward the Temple and Jerusalem, | toward the-Temple and-Jerusalem
- 2:16.53 2:16.54 2:16.55 2:16.56 2:16.57 2:16.58 2:16.59 | For the LORD chose Zion; He desired it as His dwelling. | for chose the-LORD Zion desired as His-dwelling
- 2:16.60 2:16.61 2:16.62 2:16.63 2:16.64 2:16.65 | And the Rambam had already explained in chapter 64 of the first, | and-already explained the-Rambam in-chapter 64 of-the-first
- 2:16.66 2:16.67 2:16.68 2:16.69 2:16.70 2:16.71 2:16.72 2:16.73 | that the glory of the LORD sometimes denotes the created light, | that the-glory-of the-LORD sometimes they-mean by-it the-light that-is-created
- 2:16.74 2:16.75 2:16.76 2:16.77 2:16.78 | which the Holy One makes dwell in the place, as | which makes-dwell the-Holy-One in-the-place as
- 2:16.79 2:16.80 2:16.81 2:16.82 2:16.83 2:16.84 2:16.85 2:16.86 | "and the glory of the LORD dwelt on Mount Sinai, and covered it, etc." | and-dwelt the-glory-of the-LORD on the-mountain Sinai and-covered-it etc.
- 2:16.87 2:16.88 2:16.89 2:16.90 2:16.91 2:16.92 | and sometimes by it they mean His essence, may He be blessed, and His reality, | and-sometimes they-mean by-it His-essence may-He-be-blessed and-His-reality
- 2:16.93 2:16.94 2:16.95 2:16.96 2:16.97 | as in "show me, please, Your glory," | as-in show-me please (object-marker) Your-glory
- 2:16.98 2:16.99 2:16.100 2:16.101 2:16.102 2:16.103 2:16.104 | and the answer came: "for no man shall see Me and live," | and-came the-answer that no shall-see-me man and-live
- 2:16.105 2:16.106 2:16.107 2:16.108 2:16.109 2:16.110 2:16.111 | which shows that the glory spoken of here is His very essence, may He be blessed, | shows that the-glory spoken-of here itself may-He-be-blessed
- 2:16.112 2:16.113 2:16.114 2:16.115 | and says "Your glory," to magnify, etc. | and-says Your-glory to-magnify etc.
- 2:16.116 2:16.117 2:16.118 2:16.119 2:16.120 2:16.121 2:16.122 2:16.123 2:16.124 2:16.125 | And similarly we find in the words of our Rabbis of blessed memory that the name Elohim is shared by judges and by angels, | and-similarly in-this we-find in-the-words-of our-Rabbis-of-blessed-memory that-the-name Elohim is-shared by-judges and-by-angels
- 2:16.126 2:16.127 2:16.128 2:16.129 2:16.130 2:16.131 2:16.132 2:16.133 | and by the Name, may He be blessed. Also, the name Shekhinah is shared with what I mentioned above. | and-by-the-Name,-may-He-be-blessed also the-name Shekhinah is-shared with-what I-mentioned above
- 2:16.134 2:16.135 2:16.136 | And I saw fit to alert you here. | and-I-saw-fit to-alert-you here
- 2:16.137 2:16.138 2:16.139 2:16.140 | that there is no prohibition on bowing down to a man, | that-there-is-not a-prohibition to-bow-down to-a-man
- 2:16.141 2:16.142 2:16.143 2:16.144 2:16.145 | so long as his intent is not to accept him as God, | so-long-as there-is-not his-intent to-accept-him as-God
- 2:16.146 2:16.147 2:16.148 | and he is not worshipped, | and-also he-is-not worshipped
- 2:16.149 2:16.150 2:16.151 2:16.152 2:16.153 2:16.154 | but only submits to him and humbles himself before him, | but-only submits to-him and-humbles himself before-him
- 2:16.155 2:16.156 2:16.157 2:16.158 2:16.159 2:16.160 2:16.161 | or by way of acknowledgment for a kindness that he did for him. | or by-way-of acknowledgment for a-kindness that-he-did for-him
- 2:16.162 2:16.163 2:16.164 2:16.165 2:16.166 2:16.167 2:16.168 2:16.169 2:16.170 2:16.171 2:16.172 2:16.173 2:16.174 2:16.175 | For indeed Abraham our father bowed down to the sons of Heth for the kindness they did with him, to give him a burial place. | for-indeed Abraham our-father bowed-down to-the-sons-of Heth for the-kindness that-they-did with-him to-give him a-possession-of a-burial-place
- 2:16.176 2:16.177 2:16.178 2:16.179 2:16.180 | And so Rabbi Saadia Gaon rendered: | and-so rendered Rabbi Saadia Gaon
- 2:16.181 2:16.182 2:16.183 2:16.184 2:16.185 2:16.186 2:16.187 2:16.188 2:16.189 | "and bowed down to the people of the land, to the sons of Heth, and prostrated to the tent of the land that He created." | and-bowed-down to-the-people-of the-land to-the-sons-of Heth and-prostrated to-the-tent-of the-land that-He-created
- 2:16.190 2:16.191 2:16.192 | And Joshua bowed down to the angel. | and-Joshua bowed-down to-the-angel
- 2:16.193 2:16.194 2:16.195 2:16.196 2:16.197 2:16.198 | And Nathan the prophet and Bathsheba bowed down to David. | and-Nathan the-prophet and-Bath- sheba bowed-down to-David
- 2:16.199 2:16.200 2:16.201 2:16.202 2:16.203 | And we find much in the Writings that they bowed down to the created being. | and-much we-find in-the-Writings that-they-bowed-down to-the-created-being
- 2:16.204 2:16.205 2:16.206 2:16.207 2:16.208 2:16.209 2:16.210 2:16.211 | And we read in Sanhedrin, chapter Four Deaths (Sanhedrin 61b): | and-we-read in-Sanhedrin chapter Four Deaths Sanhedrin 61 b
- 2:16.212 2:16.213 2:16.214 2:16.215 2:16.216 | It was stated: "One who worships idolatry out of love or out of fear, | it-was-stated one-who-worships idolatry out-of-love or-out-of-fear
- 2:16.217 2:16.218 2:16.219 2:16.220 2:16.221 | Abaye says: he is liable, for he worshipped it. | Abaye says liable for worshipped-it
- 2:16.222 2:16.223 2:16.224 | Rava says: he is exempt. | Rava says exempt
- 2:16.225 2:16.226 2:16.227 2:16.228 2:16.229 2:16.230 2:16.231 2:16.232 | If he accepted it upon himself as God, yes; if not, no. | if he-accepted upon-himself as-God yes if not no
- 2:16.233 2:16.234 2:16.235 2:16.236 2:16.237 | And Abaye said: How do I know this? | and-said Abaye from-where do-I-say it
- 2:16.238 2:16.239 2:16.240 2:16.241 2:16.242 | for we learned: 'One who worships idolatry: one who worships, | as-we-learned one-who-worships idolatry one-who worships
- 2:16.243 2:16.244 2:16.245 2:16.246 2:16.247 2:16.248 2:16.249 2:16.250 | and one who sacrifices, and one who burns incense, and one who pours libations, and one who bows down, | and-one-who sacrifices and-one-who burns-incense and-one-who pours-libations and-one-who bows-down
- 2:16.251 2:16.252 2:16.253 2:16.254 2:16.255 2:16.256 2:16.257 2:16.258 | and one who accepts it upon himself as God, and one who says to him: You are my God.' | and-one-who accepts-it upon-himself as-God and-one-who-says to-him my-God you
- 2:16.259 2:16.260 2:16.261 2:16.262 2:16.263 2:16.264 2:16.265 | What? Is it not that 'one who worships' means out of love or out of fear? | what is-it-not one-who worships out-of-love or out-of-fear
- 2:16.266 2:16.267 2:16.268 2:16.269 2:16.270 2:16.271 2:16.272 2:16.273 | And Rava says to you: No, but as Rabbi Yirmiyah resolves it: | and-Rava said to-you no but as-he-resolves Rabbi Yirmiyah
- 2:16.274 2:16.275 | (Rashi explains: | (explains Rashi:
- 2:16.276 2:16.277 2:16.278 2:16.279 2:16.280 2:16.281 2:16.282 | one who worships every worship that is customary for it, | one-who worships every worship that-is customary for-it
- 2:16.283 2:16.284 2:16.285 2:16.286 2:16.287 | and one who sacrifices, etc., not in its customary way, | and-one-who sacrifices etc. not in-its-customary-way
- 2:16.288 2:16.289 2:16.290 2:16.291 2:16.292 | and in every case the first clause and the last are for the sake of divinity). | and-always the-first-clause and-the-last for-the-sake-of divinity
- 2:16.293 2:16.294 2:16.295 2:16.296 2:16.297 | Abaye said: How do I know this? | said Abaye from-where do-I-say it
- 2:16.298 2:16.299 2:16.300 2:16.301 | for it was taught: 'You shall not bow down to them; | for-it-was-taught you-shall-not bow-down to-them
- 2:16.302 2:16.303 2:16.304 2:16.305 | to them you do not bow down, | to-them not you bow-down
- 2:16.306 2:16.307 2:16.308 2:16.309 2:16.310 | but you bow down to a man like yourself.' | but you bow-down to-a-man like-yourself
- 2:16.311 2:16.312 2:16.313 2:16.314 2:16.315 2:16.316 2:16.317 | One might think: even to worship like Haman? The verse says: 'and you shall not worship them.' | one-might-think even to-worship like-Haman the-verse-says and-not you-shall-worship-them
- 2:16.318 2:16.319 2:16.320 2:16.321 2:16.322 | But Haman was worshipped out of fear! | but Haman out-of-fear he was-worshipped
- 2:16.323 2:16.324 2:16.325 2:16.326 | And Rava? Like Haman, and not like Haman. | and-Rava like-Haman and-not like-Haman
- 2:16.327 2:16.328 2:16.329 2:16.330 | Like Haman, for he himself is idolatry; | like-Haman for-he himself idolatry
- 2:16.331 2:16.332 2:16.333 2:16.334 2:16.335 2:16.336 2:16.337 2:16.338 | and not like Haman, for in the case of Haman it was out of fear, and here it is not out of fear. | and-not like-Haman for Haman out-of-fear and-here it-is-not out-of-fear
- 2:16.339 2:16.340 2:16.341 2:16.342 2:16.343 | And Abaye said: How do I know this? | and Abaye-said from-where do-I-say it
- 2:16.344 2:16.345 2:16.346 2:16.347 | for it was taught: 'An anointed priest in idolatry. | for-it-was-taught a-priest anointed in-idolatry
- 2:16.348 2:16.349 2:16.350 2:16.351 | Rabbi says: in an unwitting act, | Rabbi says in-an-unwitting act
- 2:16.352 2:16.353 2:16.354 2:16.355 | and the Sages say: in a concealed matter, | and-the-Sages say in-a-concealed matter
- 2:16.356 2:16.357 2:16.358 | and they agree that for a he-goat it is like an individual, | and-they-agree that-for-a-he-goat it-is-like-an-individual
- 2:16.359 2:16.360 2:16.361 2:16.362 2:16.363 | and they agree that he does not bring a suspended guilt offering.' | and-they-agree that-he-does-not bring a-guilt-offering suspended
- 2:16.364 2:16.365 2:16.366 2:16.367 2:16.368 | What is this unwitting act like? | this unwitting act what-is it
- 2:16.369 2:16.370 2:16.371 2:16.372 2:16.373 2:16.374 2:16.375 2:16.376 | If he holds that it is a synagogue and bowed down to it, then his heart is with Heaven. | if he-holds a-synagogue and-bowed-down to-it then his-heart is-with-Heaven
- 2:16.377 2:16.378 2:16.379 2:16.380 2:16.381 | But suppose he saw a statue and bowed down to it, | but he-saw a-statue and-bowed-down to-it
- 2:16.382 2:16.383 2:16.384 2:16.385 2:16.386 2:16.387 2:16.388 2:16.389 2:16.390 2:16.391 2:16.392 2:16.393 2:16.394 2:16.395 2:16.396 | (explanation: that he saw the figure of the king, which they were accustomed to make in the likeness of the king, and one who saw it bows down to it in honor of the king, | (explanation that-he-saw the-figure-of the-king which-they-were accustomed to-make in-the-likeness-of the-king and-one-who-saw it bows-down to-it in-honor-of the-king,
- 2:16.397 2:16.398 2:16.399 2:16.400 2:16.401 2:16.402 2:16.403 2:16.404 2:16.405 | and sometimes they worship it, and this one saw one who was worshipped and bowed down to it). | and-sometimes they-worship it and-this-one saw one who-was-worshipped and-bowed-down to-it
- 2:16.406 2:16.407 2:16.408 2:16.409 2:16.410 2:16.411 2:16.412 2:16.413 2:16.414 2:16.415 2:16.416 2:16.417 2:16.418 2:16.419 | If he accepted it upon himself as God, it is deliberate; and if he did not accept it upon himself as God, it is nothing. | if he-accepted upon-himself as-God deliberate it-is and-if not he-accepted upon-himself as-God it-is-not anything it-is
- 2:16.420 2:16.421 2:16.422 2:16.423 2:16.424 2:16.425 2:16.426 2:16.427 2:16.428 2:16.429 2:16.430 2:16.431 2:16.432 2:16.433 | (Rashi explains: but he bowed down to him in honor of the king; it is nothing, for he did not intend idolatry). | (explains Rashi: but in-honor-of the-king he-bowed-down to-him it-is-not anything it-is for not he-intended idolatry
- 2:16.434 2:16.435 2:16.436 2:16.437 | But is it not out of love and out of fear? | but is-it-not out-of-love and-out-of-fear
- 2:16.438 2:16.439 2:16.440 2:16.441 2:16.442 | And Rava says to you: In the case of one who says 'permitted.' | and-Rava said to-you in-the-case-of-one-who-says permitted
- 2:16.443 2:16.444 2:16.445 2:16.446 2:16.447 | Does 'says permitted' mean concealment of a matter? | says permitted does-it-mean concealment of-a-matter
- 2:16.448 2:16.449 2:16.450 2:16.451 2:16.452 2:16.453 2:16.454 2:16.455 2:16.456 | In the case of one who says permitted entirely; concealment of a matter is upholding part and nullifying part, | in-the-case-of-one-who-says permitted entirely concealment of-a-matter upholding a-part and-nullifying a-part
- 2:16.457 2:16.458 2:16.459 2:16.460 2:16.461 2:16.462 | "Thus far the Gemara, with some of Rashi's explanation." | thus-far the-Gemara with some of-the-explanation of-Rashi
- 2:16.463 2:16.464 2:16.465 2:16.466 2:16.467 2:16.468 2:16.469 2:16.470 | And a matter is learned: that to every created being it is forbidden to bow down, except a man, | and-a-matter is-learned that-to-every created-being it-is-forbidden to-bow-down except a-man
- 2:16.471 2:16.472 2:16.473 2:16.474 2:16.475 2:16.476 2:16.477 | and this is that he be not worshipped like Haman, who made himself idolatry, | and-this-is that-not be-worshipped like-Haman who-made himself idolatry
- 2:16.478 2:16.479 2:16.480 | (as we said in Megillah 19), | as-we-said in-Megillah 19
- 2:16.481 2:16.482 2:16.483 2:16.484 2:16.485 2:16.486 2:16.487 2:16.488 2:16.489 2:16.490 | for otherwise Mordecai would not have refrained from bowing down to him, as it is taught | for-otherwise if-not so not would-have-been Mordecai refrained from-bowing-down to-him as-it-is-taught
- 2:16.491 2:16.492 2:16.493 2:16.494 2:16.495 2:16.496 2:16.497 2:16.498 2:16.499 2:16.500 2:16.501 2:16.502 2:16.503 | in the baraita: "You shall not bow down to them; to them you do not bow down, but you bow down to a man like yourself," | in-the-baraita you-shall-not bow-down to-them to-them not you bow-down but you bow-down to-a-man like-yourself
- 2:16.504 2:16.505 2:16.506 2:16.507 2:16.508 2:16.509 2:16.510 2:16.511 | and this is that he be not worshipped like Haman, who made himself God, | and-this that-not be worshipped like-Haman who-made himself God
- 2:16.512 2:16.513 2:16.514 2:16.515 2:16.516 | but only bows down to him as a sign of honor, | but-only bows-down to-him as-a-sign of-honor
- 2:16.517 2:16.518 | and not for divinity, | and-not for-divinity
- 2:16.519 2:16.520 2:16.521 2:16.522 2:16.523 2:16.524 2:16.525 2:16.526 2:16.527 | and therefore Joshua bowed down to the angel, since he appeared to him in the form of a man, | and-therefore bowed-down Joshua to-the-angel since appeared to-him in-the-form-of a-man
- 2:16.528 2:16.529 2:16.530 2:16.531 2:16.532 2:16.533 2:16.534 2:16.535 2:16.536 2:16.537 2:16.538 2:16.539 | as it is written: "And behold, a man stood before him, and therefore he bowed down to him as a sign of honor and not to worship him." | as-it-is-written and-behold a-man stood before-him and-therefore bowed-down to-him as-a-sign of-honor and-not to-worship-him
- 2:16.540 2:16.541 2:16.542 2:16.543 2:16.544 2:16.545 2:16.546 | And see in Sefer ha-Ikkarim, Treatise 2, chapter 28, | and-see in-Sefer ha-Ikkarim Treatise 2 chapter 28
- 2:16.547 2:16.548 2:16.549 2:16.550 2:16.551 2:16.552 2:16.553 2:16.554 2:16.555 2:16.556 2:16.557 2:16.558 | where he elaborated on the verses that call the glory that appears to the prophets by the name YHVH, and also the ark, as | where-he-elaborated on-the-explanation-of the-verses that-call the-glory that-appears to-the-prophets by-the-name YHVH and-also the-ark as
- 2:16.559 2:16.560 2:16.561 2:16.562 2:16.563 2:16.564 2:16.565 2:16.566 2:16.567 2:16.568 2:16.569 | as it says: "And it was when the ark set out, etc. Arise, LORD, and when it rested he said: Return, LORD." | as-it-says and-it-was when-set-out the-ark etc. Arise LORD and-when-it-rested he-said Return LORD
- 2:16.570 2:16.571 2:16.572 2:16.573 2:16.574 2:16.575 2:16.576 | And Joshua called the ark "Lord of all the earth"; see there. | and-Joshua called the-ark Lord of-all the-earth see-there
- 2:16.577 2:16.578 2:16.579 2:16.580 2:16.581 2:16.582 2:16.583 2:16.584 2:16.585 2:16.586 2:16.587 2:16.588 2:16.589 2:16.590 2:16.591 2:16.592 | And in his view, that Joshua bowed down to the angel, because that bowing was in honor of the Name, may He be blessed, and since this one was His messenger, etc.; see there. | and-in-his-view that he-bowed-down Joshua to-the-angel because the-bowing that was in-honor-of the-Name,-may-He-be-blessed and-since this-one is-His-messenger etc. see-there
- 2:16.593 2:16.594 2:16.595 2:16.596 2:16.597 2:16.598 | And our Rabbis of blessed memory said: "The fox at its time bows down to him, | and-our-Rabbis-of-blessed-memory said the-fox at-its-time bows-down to-him
- 2:16.599 2:16.600 2:16.601 2:16.602 2:16.603 2:16.604 2:16.605 2:16.606 2:16.607 2:16.608 2:16.609 2:16.610 2:16.611 | that is to say: a man lowly and despised like a fox, when the moment stands for him in some office in his hand, he bows down to him." | that-is-to-say a-man lowly and-despised like-a-fox and-the-moment stands for-him in-some office that-is-in-his-hand he-bows-down to-him
- 2:16.612 2:16.613 2:16.614 2:16.615 2:16.616 2:16.617 2:16.618 2:16.619 2:16.620 2:16.621 2:16.622 2:16.623 2:16.624 2:16.625 2:16.626 | And so it is customary with us and with the Arabs to bow down and kiss a man of great stature, and there is no prohibition in this, | and-so it-is customary with-us and-with the-Arabs to-bow-down and-to-kiss a-man of-great stature and-there-is-not in-this any prohibition
- 2:16.627 2:16.628 2:16.629 2:16.630 2:16.631 2:16.632 2:16.633 2:16.634 2:16.635 2:16.636 2:16.637 2:16.638 2:16.639 | for he does not accept him as God; it is only as a sign of honor that he bows down to him and kisses his feet. | for-he-does-not accept-him upon-himself as-God only as-a-sign of-honor it-is he-bows-down to-him and-kisses (object-marker) his-feet

Grammar notes:
- The paragraph is the respondent's reply in the correspondence of chapter 2. 'What you wrote' answers the questioner, and the question from 'And as for what you wrote' to 'created thing?' is his.
- The main sense turns on the Bavli passage (Sanhedrin 61b), quoted in Aramaic from 'It was stated' to 'Thus far the Gemara'. Its Aramaic forms are glossed in the tokens (פלחה, קבליה, דהא, אי ... אין).
- The Hebrew for 'not to the visible glory do they bow down' puts the negated object first. The English keeps that order, and 'but only' (רק) then answers it.
- שהשרה is causative (made dwell): 'to him who made His glory dwell there'. The verse cited, Exod 24:16, is cut short by the author's 'וכו'', and the English keeps the 'etc.'.
- Two verbs carry the argument: נעבד (passive, 'is worshipped') and נכנע ('submits'). The author separates bowing as worship from bowing as submission or thanks, so the English keeps the two apart.

Translator's note: The reply quotes the questioner and the Gemara, and its Saadia and Joshua examples are kept as the author gives them.

Commentary rulings:
- apparatus: Base text: Wikisource. Its additions are translated as base text and flagged here until the 1931 scan is collated (decisions, Q2).
- apparatus: Ps 132:13 (embedded, no marks): 'כי בחר ה' בציון אוה למושב לו'.
- apparatus: Exod 24:16 ('וישכון כבוד ה' על הר סיני ויכסהו'), cited with 'וכו''.
- apparatus: Exod 33:18 ('הראיני נא את כבודך') and Exod 33:20 ('כי לא יראני האדם וחי').
- apparatus: Num 10:35-36 ('ויהי בנסע הארון ... קומה ה' ... ובנחה יאמר שובה ה''), cited with 'וכו''.
- apparatus: Josh 3:11 ('אדון כל הארץ'), the author's name for the ark, cited with 'ע"ש'.
- apparatus: Josh 5:13 ('והנה איש עומד לנגדו') and Josh 5:14 (Joshua's bowing).
- apparatus: Gen 23:7: 'וישתחו לעם הארץ לבני חת'; Saadia's rendering (Judeo-Arabic in Hebrew letters) is translated from the words as they stand. The subject of 'שברא' is not marked; no other reading is certain.
- apparatus: 1 Kgs 1:16 and 1:23 (Bathsheba and Nathan bowing to David).
- apparatus: Exod 20:5: 'לא תשתחוה להם ולא תעבדם' (the Baraita in the Gemara, and the author's quotation of it).
- apparatus: Gap in the decisions: אשם תלוי ('suspended guilt offering'), ביהכנ"ס ('synagogue') and אנדרטא have no ruling; the English is literal.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

