# Translation pack: milhamot-hashem perek-8-05 (Milchamot ha-Shem 8:6–8:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-05`.

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
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
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


## Milchamot ha-Shem 8:6

Hebrew (Wikisource): ושוב אחר כן בגליון הנזכר העתיקו דבר הרב בעל ספר הברית המחזק את לבנו, ומאמץ את רוחינו בל נירא ונפחד מכבוד ה' ומהדר גאונו כי נגשימהו, כי הבקי בחכמת הניתוח, אורו מעליו יפשיט. ולא יחוס על כבוד קונו. ולא יעצרהו הגשם (אשר לדעת הראב"ד ובעל העיקרים מי שיאמין שהשי"ת בעל תואר כתואר םני המלך כגוף וגויה מפני שהוא טועה בעיונו אין לקרותו מין (עיקרים מאמר א' פ' ב') ואע"פ שאינו כן לפי האמת רק שהוא טועה בפי' הכתובים) כלומד הקבלה החדשה אע"פ שיצייר במחשבתו ובפיו ובספר יחוק כל משיגי הגוף ומקריו, וייחסם לשי"ת. הנה אחר כך ישוב ויפשיטם ויסירם. היו כלא היו, וחזר הענין לכמו שהיה, כלעומת שבא כן ילך. ולא הרויח כלום, רק הפסיד ועבר על הכתוב ועל דברי רז"ל ולא חס על כבוד קונו.

| id | word | punctuation after |
|---|---|---|
| 8:6.1 | ושוב |  |
| 8:6.2 | אחר |  |
| 8:6.3 | כן |  |
| 8:6.4 | בגליון |  |
| 8:6.5 | הנזכר |  |
| 8:6.6 | העתיקו |  |
| 8:6.7 | דבר |  |
| 8:6.8 | הרב |  |
| 8:6.9 | בעל |  |
| 8:6.10 | ספר |  |
| 8:6.11 | הברית |  |
| 8:6.12 | המחזק |  |
| 8:6.13 | את |  |
| 8:6.14 | לבנו | , |
| 8:6.15 | ומאמץ |  |
| 8:6.16 | את |  |
| 8:6.17 | רוחינו |  |
| 8:6.18 | בל |  |
| 8:6.19 | נירא |  |
| 8:6.20 | ונפחד |  |
| 8:6.21 | מכבוד |  |
| 8:6.22 | ה | ' |
| 8:6.23 | ומהדר |  |
| 8:6.24 | גאונו |  |
| 8:6.25 | כי |  |
| 8:6.26 | נגשימהו | , |
| 8:6.27 | כי |  |
| 8:6.28 | הבקי |  |
| 8:6.29 | בחכמת |  |
| 8:6.30 | הניתוח | , |
| 8:6.31 | אורו |  |
| 8:6.32 | מעליו |  |
| 8:6.33 | יפשיט | . |
| 8:6.34 | ולא |  |
| 8:6.35 | יחוס |  |
| 8:6.36 | על |  |
| 8:6.37 | כבוד |  |
| 8:6.38 | קונו | . |
| 8:6.39 | ולא |  |
| 8:6.40 | יעצרהו |  |
| 8:6.41 | הגשם |  |
| 8:6.42 | אשר |  |
| 8:6.43 | לדעת |  |
| 8:6.44 | הראב"ד |  |
| 8:6.45 | ובעל |  |
| 8:6.46 | העיקרים |  |
| 8:6.47 | מי |  |
| 8:6.48 | שיאמין |  |
| 8:6.49 | שהשי"ת |  |
| 8:6.50 | בעל |  |
| 8:6.51 | תואר |  |
| 8:6.52 | כתואר |  |
| 8:6.53 | םני |  |
| 8:6.54 | המלך |  |
| 8:6.55 | כגוף |  |
| 8:6.56 | וגויה |  |
| 8:6.57 | מפני |  |
| 8:6.58 | שהוא |  |
| 8:6.59 | טועה |  |
| 8:6.60 | בעיונו |  |
| 8:6.61 | אין |  |
| 8:6.62 | לקרותו |  |
| 8:6.63 | מין |  |
| 8:6.64 | עיקרים |  |
| 8:6.65 | מאמר |  |
| 8:6.66 | א | ' |
| 8:6.67 | פ | ' |
| 8:6.68 | ב | ') |
| 8:6.69 | ואע"פ |  |
| 8:6.70 | שאינו |  |
| 8:6.71 | כן |  |
| 8:6.72 | לפי |  |
| 8:6.73 | האמת |  |
| 8:6.74 | רק |  |
| 8:6.75 | שהוא |  |
| 8:6.76 | טועה |  |
| 8:6.77 | בפי | ' |
| 8:6.78 | הכתובים | ) |
| 8:6.79 | כלומד |  |
| 8:6.80 | הקבלה |  |
| 8:6.81 | החדשה |  |
| 8:6.82 | אע"פ |  |
| 8:6.83 | שיצייר |  |
| 8:6.84 | במחשבתו |  |
| 8:6.85 | ובפיו |  |
| 8:6.86 | ובספר |  |
| 8:6.87 | יחוק |  |
| 8:6.88 | כל |  |
| 8:6.89 | משיגי |  |
| 8:6.90 | הגוף |  |
| 8:6.91 | ומקריו | , |
| 8:6.92 | וייחסם |  |
| 8:6.93 | לשי"ת | . |
| 8:6.94 | הנה |  |
| 8:6.95 | אחר |  |
| 8:6.96 | כך |  |
| 8:6.97 | ישוב |  |
| 8:6.98 | ויפשיטם |  |
| 8:6.99 | ויסירם | . |
| 8:6.100 | היו |  |
| 8:6.101 | כלא |  |
| 8:6.102 | היו | , |
| 8:6.103 | וחזר |  |
| 8:6.104 | הענין |  |
| 8:6.105 | לכמו |  |
| 8:6.106 | שהיה | , |
| 8:6.107 | כלעומת |  |
| 8:6.108 | שבא |  |
| 8:6.109 | כן |  |
| 8:6.110 | ילך | . |
| 8:6.111 | ולא |  |
| 8:6.112 | הרויח |  |
| 8:6.113 | כלום | , |
| 8:6.114 | רק |  |
| 8:6.115 | הפסיד |  |
| 8:6.116 | ועבר |  |
| 8:6.117 | על |  |
| 8:6.118 | הכתוב |  |
| 8:6.119 | ועל |  |
| 8:6.120 | דברי |  |
| 8:6.121 | רז"ל |  |
| 8:6.122 | ולא |  |
| 8:6.123 | חס |  |
| 8:6.124 | על |  |
| 8:6.125 | כבוד |  |
| 8:6.126 | קונו | . |


## Milchamot ha-Shem 8:7

Hebrew (Wikisource): קג) ולא אמרו הנז' שהמאמין שהוא ית' גוף וגויא וכו' אינו אלא במי שהוא מאמין וחושב בדעתו שהשי"ת שהוא הסבא הראשונה הוא גוף ובעל תמונה, ברם דא עקא וחולי שיש בו סכנה ואין לו רפואה כי הוא מיחס העבודה לסבות נבראות מחודשות, הם האלהים הקדושים שהם החמשה פרצופים, ושאין לעבוד לסבה הראשונה רק לפרצוף קצר האפים (זעיר ונוקביה) ולכלול עמו גם שאר פרצופים שלמעלה ממנו, אבל העובד לסבה הראשונה בבחינתו הפשוטה (שאין לו דמות הגוף) טועה הוא ואין תפלתו תפלה, כי לא תאות העבודה רק לחלק היושר המושג (בציור הגוף). אבל העגולים עם האין סוף המקיף, אין להם שום אלהות כלל, וכל עיקר האלהות הם פרצופי היושר, והם הברואים, ועשרה מאמרות שבהם נברא העולם נתחלקו ביניהם, בראשית שהוא מאמר ראשון נתיחס לעתיק ואריך עם אבא ואימא, כמו שנתבאר לעיל סי' כ"ד דלפי' קמא בכח אבא אריך את אלהים שהיא אימא, ולפירוש בתרא בכה אבא ברא אלהים שהיא בינה את האלוהות שאחריה המכונים בשם שמים וארץ, שהם קצר אפים וזוגתו (זעיר אנפין ונוקביה) אשר העבודה שייכא בהם דווקא. ושאר הח' מאמרות נתיחסו לאבא לבדו! ומאמר עשירי שהוא נעשה אדם, נתיחס לאימא. אבל זעיר ונוקביה אשר אותם ונעבוד לפי דבריהם לא נשתתפו במעשה בראשית בשום מאמר מעשרה מאמרות כאשר יראה המעיין בספר הזהר ומפרשיו. וכבר ביארתי לך לעיל סי' כ"ה איך הרחיקו רז"ל מלחשוב כי שמים וארץ אלוהות הן. והנה הבטחתו של הרב ספר הברית שלא נירא מהגשם, לא הועילה להצילנו מאמונת המון אלוהות רבים הנעזרים זה נזה, כי לא יכול לעשות והוא נצרך לכח אחר, וגם מיעץ אותנו לעבוד את קצר אפים (זעיר אנפין) הקדוש, והוא קורא אותו אלהא רברבא, והוא בן בן בן בנו של האלהים העליון שאמר אני ראשון ואני אחרון ומבלעדי אין אלהים. אני ראשון שאין לי אב, ואני אחרון שאין לי בן. ומבלעדי אין אלהים שאין לי אח, ושעל זה האלהים העליון אמרו בזהר (בחר דף ק"ט) ואנת לית לך גופא. ולית לך איברים ולית לך נוקבא אלא חד בלא שינוי וכו' כמבואר שם ב' רב פעלים (והרב עז לאלהים פרק כ"ה כתב על מאמר רז"ל על פסוק יש אחד ואין שני גם בן ואח אין לו, אבל אחות יש לו כו' ע"ש) וכללות דבריהם כי כל העניינים המורים על האחדות הגמורה בכתובים ובדברי רז"ל לא נאמרו רק על האלהים העליון אשר לא נעבוד אותו ולא נקרא אליו. כי לא יענה ח"ו. כמאמר הנביא מוכיח לישראל הקצור קצרה ידי מפדות ואם אין בי כח הציל וכו'. וכביכול לפי דבריהם הוא משתבח בעצמו בלבד ואינו עושה למבקש ממנו כלום. אבל עיקר העבודה והתפלה היא לקצר אפים. שהוא בגוף ואיברים ובעל אשה הוא. והוא אשר יענה לכל קוראיו ברצון אביו ואמו אשר גידלו אותו ורוממוהו ואעטרו ליה בתרין עטרין עד אשר לו נתכנו עלילות! אבל לעתיקא קדישא לא נתכנו עלילות, והזעיר אנפין הוא אשר נקרא בשם הוי'ה ובשם הקב"ה. ואליו תכון עבודתינו ותפלותינו לפי דבריהם אבל האלוה העליון שהוא יחיד בלי גוף ואיברים ואין לו אשה, לא נקרא אליו ולא נתפלל, והעובד אותו ומתפלל אליו אין תפלתו תפלה ולא יענה אותו. ואדרבא ענוש יענש.

| id | word | punctuation after |
|---|---|---|
| 8:7.1 | קג | ) |
| 8:7.2 | ולא |  |
| 8:7.3 | אמרו |  |
| 8:7.4 | הנז | ' |
| 8:7.5 | שהמאמין |  |
| 8:7.6 | שהוא |  |
| 8:7.7 | ית | ' |
| 8:7.8 | גוף |  |
| 8:7.9 | וגויא |  |
| 8:7.10 | וכו | ' |
| 8:7.11 | אינו |  |
| 8:7.12 | אלא |  |
| 8:7.13 | במי |  |
| 8:7.14 | שהוא |  |
| 8:7.15 | מאמין |  |
| 8:7.16 | וחושב |  |
| 8:7.17 | בדעתו |  |
| 8:7.18 | שהשי"ת |  |
| 8:7.19 | שהוא |  |
| 8:7.20 | הסבא |  |
| 8:7.21 | הראשונה |  |
| 8:7.22 | הוא |  |
| 8:7.23 | גוף |  |
| 8:7.24 | ובעל |  |
| 8:7.25 | תמונה | , |
| 8:7.26 | ברם |  |
| 8:7.27 | דא |  |
| 8:7.28 | עקא |  |
| 8:7.29 | וחולי |  |
| 8:7.30 | שיש |  |
| 8:7.31 | בו |  |
| 8:7.32 | סכנה |  |
| 8:7.33 | ואין |  |
| 8:7.34 | לו |  |
| 8:7.35 | רפואה |  |
| 8:7.36 | כי |  |
| 8:7.37 | הוא |  |
| 8:7.38 | מיחס |  |
| 8:7.39 | העבודה |  |
| 8:7.40 | לסבות |  |
| 8:7.41 | נבראות |  |
| 8:7.42 | מחודשות | , |
| 8:7.43 | הם |  |
| 8:7.44 | האלהים |  |
| 8:7.45 | הקדושים |  |
| 8:7.46 | שהם |  |
| 8:7.47 | החמשה |  |
| 8:7.48 | פרצופים | , |
| 8:7.49 | ושאין |  |
| 8:7.50 | לעבוד |  |
| 8:7.51 | לסבה |  |
| 8:7.52 | הראשונה |  |
| 8:7.53 | רק |  |
| 8:7.54 | לפרצוף |  |
| 8:7.55 | קצר |  |
| 8:7.56 | האפים |  |
| 8:7.57 | זעיר |  |
| 8:7.58 | ונוקביה | ) |
| 8:7.59 | ולכלול |  |
| 8:7.60 | עמו |  |
| 8:7.61 | גם |  |
| 8:7.62 | שאר |  |
| 8:7.63 | פרצופים |  |
| 8:7.64 | שלמעלה |  |
| 8:7.65 | ממנו | , |
| 8:7.66 | אבל |  |
| 8:7.67 | העובד |  |
| 8:7.68 | לסבה |  |
| 8:7.69 | הראשונה |  |
| 8:7.70 | בבחינתו |  |
| 8:7.71 | הפשוטה |  |
| 8:7.72 | שאין |  |
| 8:7.73 | לו |  |
| 8:7.74 | דמות |  |
| 8:7.75 | הגוף | ) |
| 8:7.76 | טועה |  |
| 8:7.77 | הוא |  |
| 8:7.78 | ואין |  |
| 8:7.79 | תפלתו |  |
| 8:7.80 | תפלה | , |
| 8:7.81 | כי |  |
| 8:7.82 | לא |  |
| 8:7.83 | תאות |  |
| 8:7.84 | העבודה |  |
| 8:7.85 | רק |  |
| 8:7.86 | לחלק |  |
| 8:7.87 | היושר |  |
| 8:7.88 | המושג |  |
| 8:7.89 | בציור |  |
| 8:7.90 | הגוף | ). |
| 8:7.91 | אבל |  |
| 8:7.92 | העגולים |  |
| 8:7.93 | עם |  |
| 8:7.94 | האין |  |
| 8:7.95 | סוף |  |
| 8:7.96 | המקיף | , |
| 8:7.97 | אין |  |
| 8:7.98 | להם |  |
| 8:7.99 | שום |  |
| 8:7.100 | אלהות |  |
| 8:7.101 | כלל | , |
| 8:7.102 | וכל |  |
| 8:7.103 | עיקר |  |
| 8:7.104 | האלהות |  |
| 8:7.105 | הם |  |
| 8:7.106 | פרצופי |  |
| 8:7.107 | היושר | , |
| 8:7.108 | והם |  |
| 8:7.109 | הברואים | , |
| 8:7.110 | ועשרה |  |
| 8:7.111 | מאמרות |  |
| 8:7.112 | שבהם |  |
| 8:7.113 | נברא |  |
| 8:7.114 | העולם |  |
| 8:7.115 | נתחלקו |  |
| 8:7.116 | ביניהם | , |
| 8:7.117 | בראשית |  |
| 8:7.118 | שהוא |  |
| 8:7.119 | מאמר |  |
| 8:7.120 | ראשון |  |
| 8:7.121 | נתיחס |  |
| 8:7.122 | לעתיק |  |
| 8:7.123 | ואריך |  |
| 8:7.124 | עם |  |
| 8:7.125 | אבא |  |
| 8:7.126 | ואימא | , |
| 8:7.127 | כמו |  |
| 8:7.128 | שנתבאר |  |
| 8:7.129 | לעיל |  |
| 8:7.130 | סי | ' |
| 8:7.131 | כ"ד |  |
| 8:7.132 | דלפי | ' |
| 8:7.133 | קמא |  |
| 8:7.134 | בכח |  |
| 8:7.135 | אבא |  |
| 8:7.136 | אריך |  |
| 8:7.137 | את |  |
| 8:7.138 | אלהים |  |
| 8:7.139 | שהיא |  |
| 8:7.140 | אימא | , |
| 8:7.141 | ולפירוש |  |
| 8:7.142 | בתרא |  |
| 8:7.143 | בכה |  |
| 8:7.144 | אבא |  |
| 8:7.145 | ברא |  |
| 8:7.146 | אלהים |  |
| 8:7.147 | שהיא |  |
| 8:7.148 | בינה |  |
| 8:7.149 | את |  |
| 8:7.150 | האלוהות |  |
| 8:7.151 | שאחריה |  |
| 8:7.152 | המכונים |  |
| 8:7.153 | בשם |  |
| 8:7.154 | שמים |  |
| 8:7.155 | וארץ | , |
| 8:7.156 | שהם |  |
| 8:7.157 | קצר |  |
| 8:7.158 | אפים |  |
| 8:7.159 | וזוגתו |  |
| 8:7.160 | זעיר |  |
| 8:7.161 | אנפין |  |
| 8:7.162 | ונוקביה | ) |
| 8:7.163 | אשר |  |
| 8:7.164 | העבודה |  |
| 8:7.165 | שייכא |  |
| 8:7.166 | בהם |  |
| 8:7.167 | דווקא | . |
| 8:7.168 | ושאר |  |
| 8:7.169 | הח | ' |
| 8:7.170 | מאמרות |  |
| 8:7.171 | נתיחסו |  |
| 8:7.172 | לאבא |  |
| 8:7.173 | לבדו | ! |
| 8:7.174 | ומאמר |  |
| 8:7.175 | עשירי |  |
| 8:7.176 | שהוא |  |
| 8:7.177 | נעשה |  |
| 8:7.178 | אדם | , |
| 8:7.179 | נתיחס |  |
| 8:7.180 | לאימא | . |
| 8:7.181 | אבל |  |
| 8:7.182 | זעיר |  |
| 8:7.183 | ונוקביה |  |
| 8:7.184 | אשר |  |
| 8:7.185 | אותם |  |
| 8:7.186 | ונעבוד |  |
| 8:7.187 | לפי |  |
| 8:7.188 | דבריהם |  |
| 8:7.189 | לא |  |
| 8:7.190 | נשתתפו |  |
| 8:7.191 | במעשה |  |
| 8:7.192 | בראשית |  |
| 8:7.193 | בשום |  |
| 8:7.194 | מאמר |  |
| 8:7.195 | מעשרה |  |
| 8:7.196 | מאמרות |  |
| 8:7.197 | כאשר |  |
| 8:7.198 | יראה |  |
| 8:7.199 | המעיין |  |
| 8:7.200 | בספר |  |
| 8:7.201 | הזהר |  |
| 8:7.202 | ומפרשיו | . |
| 8:7.203 | וכבר |  |
| 8:7.204 | ביארתי |  |
| 8:7.205 | לך |  |
| 8:7.206 | לעיל |  |
| 8:7.207 | סי | ' |
| 8:7.208 | כ"ה |  |
| 8:7.209 | איך |  |
| 8:7.210 | הרחיקו |  |
| 8:7.211 | רז"ל |  |
| 8:7.212 | מלחשוב |  |
| 8:7.213 | כי |  |
| 8:7.214 | שמים |  |
| 8:7.215 | וארץ |  |
| 8:7.216 | אלוהות |  |
| 8:7.217 | הן | . |
| 8:7.218 | והנה |  |
| 8:7.219 | הבטחתו |  |
| 8:7.220 | של |  |
| 8:7.221 | הרב |  |
| 8:7.222 | ספר |  |
| 8:7.223 | הברית |  |
| 8:7.224 | שלא |  |
| 8:7.225 | נירא |  |
| 8:7.226 | מהגשם | , |
| 8:7.227 | לא |  |
| 8:7.228 | הועילה |  |
| 8:7.229 | להצילנו |  |
| 8:7.230 | מאמונת |  |
| 8:7.231 | המון |  |
| 8:7.232 | אלוהות |  |
| 8:7.233 | רבים |  |
| 8:7.234 | הנעזרים |  |
| 8:7.235 | זה |  |
| 8:7.236 | נזה | , |
| 8:7.237 | כי |  |
| 8:7.238 | לא |  |
| 8:7.239 | יכול |  |
| 8:7.240 | לעשות |  |
| 8:7.241 | והוא |  |
| 8:7.242 | נצרך |  |
| 8:7.243 | לכח |  |
| 8:7.244 | אחר | , |
| 8:7.245 | וגם |  |
| 8:7.246 | מיעץ |  |
| 8:7.247 | אותנו |  |
| 8:7.248 | לעבוד |  |
| 8:7.249 | את |  |
| 8:7.250 | קצר |  |
| 8:7.251 | אפים |  |
| 8:7.252 | זעיר |  |
| 8:7.253 | אנפין | ) |
| 8:7.254 | הקדוש | , |
| 8:7.255 | והוא |  |
| 8:7.256 | קורא |  |
| 8:7.257 | אותו |  |
| 8:7.258 | אלהא |  |
| 8:7.259 | רברבא | , |
| 8:7.260 | והוא |  |
| 8:7.261 | בן |  |
| 8:7.262 | בן |  |
| 8:7.263 | בן |  |
| 8:7.264 | בנו |  |
| 8:7.265 | של |  |
| 8:7.266 | האלהים |  |
| 8:7.267 | העליון |  |
| 8:7.268 | שאמר |  |
| 8:7.269 | אני |  |
| 8:7.270 | ראשון |  |
| 8:7.271 | ואני |  |
| 8:7.272 | אחרון |  |
| 8:7.273 | ומבלעדי |  |
| 8:7.274 | אין |  |
| 8:7.275 | אלהים | . |
| 8:7.276 | אני |  |
| 8:7.277 | ראשון |  |
| 8:7.278 | שאין |  |
| 8:7.279 | לי |  |
| 8:7.280 | אב | , |
| 8:7.281 | ואני |  |
| 8:7.282 | אחרון |  |
| 8:7.283 | שאין |  |
| 8:7.284 | לי |  |
| 8:7.285 | בן | . |
| 8:7.286 | ומבלעדי |  |
| 8:7.287 | אין |  |
| 8:7.288 | אלהים |  |
| 8:7.289 | שאין |  |
| 8:7.290 | לי |  |
| 8:7.291 | אח | , |
| 8:7.292 | ושעל |  |
| 8:7.293 | זה |  |
| 8:7.294 | האלהים |  |
| 8:7.295 | העליון |  |
| 8:7.296 | אמרו |  |
| 8:7.297 | בזהר |  |
| 8:7.298 | בחר |  |
| 8:7.299 | דף |  |
| 8:7.300 | ק"ט | ) |
| 8:7.301 | ואנת |  |
| 8:7.302 | לית |  |
| 8:7.303 | לך |  |
| 8:7.304 | גופא | . |
| 8:7.305 | ולית |  |
| 8:7.306 | לך |  |
| 8:7.307 | איברים |  |
| 8:7.308 | ולית |  |
| 8:7.309 | לך |  |
| 8:7.310 | נוקבא |  |
| 8:7.311 | אלא |  |
| 8:7.312 | חד |  |
| 8:7.313 | בלא |  |
| 8:7.314 | שינוי |  |
| 8:7.315 | וכו | ' |
| 8:7.316 | כמבואר |  |
| 8:7.317 | שם |  |
| 8:7.318 | ב | ' |
| 8:7.319 | רב |  |
| 8:7.320 | פעלים |  |
| 8:7.321 | והרב |  |
| 8:7.322 | עז |  |
| 8:7.323 | לאלהים |  |
| 8:7.324 | פרק |  |
| 8:7.325 | כ"ה |  |
| 8:7.326 | כתב |  |
| 8:7.327 | על |  |
| 8:7.328 | מאמר |  |
| 8:7.329 | רז"ל |  |
| 8:7.330 | על |  |
| 8:7.331 | פסוק |  |
| 8:7.332 | יש |  |
| 8:7.333 | אחד |  |
| 8:7.334 | ואין |  |
| 8:7.335 | שני |  |
| 8:7.336 | גם |  |
| 8:7.337 | בן |  |
| 8:7.338 | ואח |  |
| 8:7.339 | אין |  |
| 8:7.340 | לו | , |
| 8:7.341 | אבל |  |
| 8:7.342 | אחות |  |
| 8:7.343 | יש |  |
| 8:7.344 | לו |  |
| 8:7.345 | כו | ' |
| 8:7.346 | ע"ש | ) |
| 8:7.347 | וכללות |  |
| 8:7.348 | דבריהם |  |
| 8:7.349 | כי |  |
| 8:7.350 | כל |  |
| 8:7.351 | העניינים |  |
| 8:7.352 | המורים |  |
| 8:7.353 | על |  |
| 8:7.354 | האחדות |  |
| 8:7.355 | הגמורה |  |
| 8:7.356 | בכתובים |  |
| 8:7.357 | ובדברי |  |
| 8:7.358 | רז"ל |  |
| 8:7.359 | לא |  |
| 8:7.360 | נאמרו |  |
| 8:7.361 | רק |  |
| 8:7.362 | על |  |
| 8:7.363 | האלהים |  |
| 8:7.364 | העליון |  |
| 8:7.365 | אשר |  |
| 8:7.366 | לא |  |
| 8:7.367 | נעבוד |  |
| 8:7.368 | אותו |  |
| 8:7.369 | ולא |  |
| 8:7.370 | נקרא |  |
| 8:7.371 | אליו | . |
| 8:7.372 | כי |  |
| 8:7.373 | לא |  |
| 8:7.374 | יענה |  |
| 8:7.375 | ח"ו | . |
| 8:7.376 | כמאמר |  |
| 8:7.377 | הנביא |  |
| 8:7.378 | מוכיח |  |
| 8:7.379 | לישראל |  |
| 8:7.380 | הקצור |  |
| 8:7.381 | קצרה |  |
| 8:7.382 | ידי |  |
| 8:7.383 | מפדות |  |
| 8:7.384 | ואם |  |
| 8:7.385 | אין |  |
| 8:7.386 | בי |  |
| 8:7.387 | כח |  |
| 8:7.388 | הציל |  |
| 8:7.389 | וכו | '. |
| 8:7.390 | וכביכול |  |
| 8:7.391 | לפי |  |
| 8:7.392 | דבריהם |  |
| 8:7.393 | הוא |  |
| 8:7.394 | משתבח |  |
| 8:7.395 | בעצמו |  |
| 8:7.396 | בלבד |  |
| 8:7.397 | ואינו |  |
| 8:7.398 | עושה |  |
| 8:7.399 | למבקש |  |
| 8:7.400 | ממנו |  |
| 8:7.401 | כלום | . |
| 8:7.402 | אבל |  |
| 8:7.403 | עיקר |  |
| 8:7.404 | העבודה |  |
| 8:7.405 | והתפלה |  |
| 8:7.406 | היא |  |
| 8:7.407 | לקצר |  |
| 8:7.408 | אפים | . |
| 8:7.409 | שהוא |  |
| 8:7.410 | בגוף |  |
| 8:7.411 | ואיברים |  |
| 8:7.412 | ובעל |  |
| 8:7.413 | אשה |  |
| 8:7.414 | הוא | . |
| 8:7.415 | והוא |  |
| 8:7.416 | אשר |  |
| 8:7.417 | יענה |  |
| 8:7.418 | לכל |  |
| 8:7.419 | קוראיו |  |
| 8:7.420 | ברצון |  |
| 8:7.421 | אביו |  |
| 8:7.422 | ואמו |  |
| 8:7.423 | אשר |  |
| 8:7.424 | גידלו |  |
| 8:7.425 | אותו |  |
| 8:7.426 | ורוממוהו |  |
| 8:7.427 | ואעטרו |  |
| 8:7.428 | ליה |  |
| 8:7.429 | בתרין |  |
| 8:7.430 | עטרין |  |
| 8:7.431 | עד |  |
| 8:7.432 | אשר |  |
| 8:7.433 | לו |  |
| 8:7.434 | נתכנו |  |
| 8:7.435 | עלילות | ! |
| 8:7.436 | אבל |  |
| 8:7.437 | לעתיקא |  |
| 8:7.438 | קדישא |  |
| 8:7.439 | לא |  |
| 8:7.440 | נתכנו |  |
| 8:7.441 | עלילות | , |
| 8:7.442 | והזעיר |  |
| 8:7.443 | אנפין |  |
| 8:7.444 | הוא |  |
| 8:7.445 | אשר |  |
| 8:7.446 | נקרא |  |
| 8:7.447 | בשם |  |
| 8:7.448 | הוי'ה |  |
| 8:7.449 | ובשם |  |
| 8:7.450 | הקב"ה | . |
| 8:7.451 | ואליו |  |
| 8:7.452 | תכון |  |
| 8:7.453 | עבודתינו |  |
| 8:7.454 | ותפלותינו |  |
| 8:7.455 | לפי |  |
| 8:7.456 | דבריהם |  |
| 8:7.457 | אבל |  |
| 8:7.458 | האלוה |  |
| 8:7.459 | העליון |  |
| 8:7.460 | שהוא |  |
| 8:7.461 | יחיד |  |
| 8:7.462 | בלי |  |
| 8:7.463 | גוף |  |
| 8:7.464 | ואיברים |  |
| 8:7.465 | ואין |  |
| 8:7.466 | לו |  |
| 8:7.467 | אשה | , |
| 8:7.468 | לא |  |
| 8:7.469 | נקרא |  |
| 8:7.470 | אליו |  |
| 8:7.471 | ולא |  |
| 8:7.472 | נתפלל | , |
| 8:7.473 | והעובד |  |
| 8:7.474 | אותו |  |
| 8:7.475 | ומתפלל |  |
| 8:7.476 | אליו |  |
| 8:7.477 | אין |  |
| 8:7.478 | תפלתו |  |
| 8:7.479 | תפלה |  |
| 8:7.480 | ולא |  |
| 8:7.481 | יענה |  |
| 8:7.482 | אותו | . |
| 8:7.483 | ואדרבא |  |
| 8:7.484 | ענוש |  |
| 8:7.485 | יענש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

