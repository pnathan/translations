# Translation pack: milhamot-hashem perek-9-16 (Milchamot ha-Shem 9:19–9:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-16`.

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
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
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
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
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


## Milchamot ha-Shem 9:19

Hebrew (Wikisource): קלז) ובמדרש רבה פרשת קרח דרשו רז"ל ששים המה מלכות, אלו ס' מסכתות ושמנים פילגשים אלו שמנים בתי מדרשות שהיו בירושלם כנגד פתחיה, ועלמות אין מספר, משנה החיצונה, ופירש הרב עץ יוסף ומת"כ דקרי לבתי מדרשות פילגשים על שם שנושאין ונותנים בהם וחולקים בהם וחולקים זה עם זה כמו פילגשם (ומחבר הזהר כשנאתו למשנה ולתלמוד דימה שקלא וטריא דרבנן לכלביך דצוחון הב הב ושם אותם מיורשי גהינם האומרת הב הב) ופי' משנה החיצונה מה שלמדו היחידים חוץ לאלו בתי מדרשות, וביד יוסף פי' אלו הברייתות, ע"כ, וכן מצאתי בשיר השירים רבה וז"ל ששים המה מלכות, אלו ששים מסכתות של הלכות ושמנים פילגשים אלו שמנים פרשיות שבתורת כהנים. ועלמות אין מספר, אין קץ לתוספתות, אחת היא הן חולקין אלו עם אלו וכולהון דורשין מטעם אחד, ומהלכה אחת מגזרה שוה מקל וחומר ע"כ. ראה ידידי הקורא כמה מחבבים רז"ל את המשנה והתלמוד וקראו להן מלכות ולא שפחות, כמאמר שלמה כי מלכים ימלוכו, והמה בתרה של תורה כי לא נוכל לעשות דבר ממצות התורה כי אם על פי קבלת רז"ל בעלי המשנה והתלמודף ולא נסור ימין ושמאל מקבלתם בפירוש המצות היאך נעשה אותם כמשפטם, וכ"ש וק"ו באמונת אחדות הבורא ית' ושלילת מאורעות הגוף ומקריו ממנו ית'. מריה דאברהם, לא לנו ה' לא לנו לחשוב רבוי באלהותך כאשר עשו המקובלים החדשים על פי הזהר והתקונים שמיחסים אותם לרשב"י בשקר, כי נתחו (לפי עניות דעתם) את ה' אלהים אמת האל הגדול הגבור והנורא לנתחים וצורות ופרצופים רבים שונים זה מזה באיכותן וכמותן וחומרן וטבען, ויערכו את נתחיו לעין כל קורא בספריהם, וקורא להם מערכת האלהות, ויקראו לכל נתח בשם, כשם הגדולים אשר בארץ וזהו סדר עריכתם כמבואר בספרים, אין סוף העליון אשר אין למעלה ממנו, אדם קדמאה, אדם קדמון, עתיק ונוקביה, אריך אנפין ונוקביה. אבא ואימא עילאין, ישראל סבא ותבונה, יעקב ולאה זעיר ונוקביה, והוא נקרא ישראל והיא רחל, (עץ חיים שער המלכים פ"ט). ובשער אבא ואימא פ"ב כ' וכבר נודע מ"ש בענין לאה ורחל בהיותן עומדין באחורי זעיר אנפין שאז העקביים של לאה נכנסים תוך ראש רחל העומדת תחתיה, וכן הענין בכאן שעקביים של בינה נכנסין תוך ראש תבונה עכ"ל. וישסעו את אריה דבי עלאי אריה שאג מי לא יירא ה' אלהים דבר, כשסע הגדי ידים ורגלים, ועקביים וחזה ושוק וטבור, ושאר איברים נסתרים כגיד הערוה ותרין ביעין לכל אחד מן הפרצופים! ויראו לאחרים את ערות אבותיהם שבשמים (לפי דעתם שחושבים רבוי פרצופים כאל אחד) ויכניסו אותם במכחול בשפופרת ברחם הנקבות אשר בדו מלבם, כמבואר בעץ חיים בכמה מקומות ושם מבואר שהיסוד כולו הוא בחינת פנים ואין לו אחוריים לפי שכל היסוד של זכר נכנס צוך היסוד של נקבה, וכולו בחינת פנים, מה שאין כן בשאר הגוף שיש בו בחינת פנים הנחזרים נגד פני הנקבה! אוי לאותה בושה ואני לאותה כלמה לספר דברים מכוערים כאלה באלהינו שבשמים אשר מלא כל הארץ כבודו, ולקרוא לו בשם מגונה קצר אפים (ז"א). וכבר הודעתיך בסי' צ"ב דברי רבינו הגדול הרמב"ם ז"ל בספרו הנכבד מורה נבוכים שתורתינו הקדושה וכל הנביאים וחכמינו ז"ל הרחיקו מהשי"ת. שני חושים מהחמשה חושים שבאדם, והן הטעם והמשוש, וכ"ש המשגל שהוא הפחות והמגונה וכו'. וכבר ידעת מה שדרשו רז"ל על פסוק ותמת שם מרים, תקבר שם דאמר ר' אלעזר אף מרים מתה בנשיקה. איתא שם שם, ומפני מה לא נאמר בה על פי ה', שגנאי הדבר לומר ואם כל כך הקפיד הכתוב שלא לומר במרים הנביאה על פי ה' מטעם שגנאי הדבר לומר כן. כל שכן קל וחומר ליחס לשי"ת זכרות ונקבות וזווגין והטפת טפין ולומר שמקום חיבור גיד הערוה הוא למעלה מתרין פרקין נצח והוד כמו שהוא באדם כמו שכתב מהרח"ו בעץ חיים שער נקודים פ"ג. וכבר הבאתי לעיל סי' צ"ד דברי בראשית רבה ודברי הירושלמי על פסוק תאלמנה שפתי שקר הדוברות על צדיק עתק בגאוה ובוז, יתחרשן יתפרכן ישתתקן וכו' הדוברות על צדיק עתק. על צדיקו של עולם דברים שהעתיק מבריותיו בגאוה ובוז, שהוא מתגאה לומר אני דורש במעשה בראשית, סבור שהוא כמגאה ואינו אלא ככבזה. אמר ר' חנינה המתכבד בקלון חבירו אין לו חלק לעולם הבא. המתכבד בכבוד חי העולמים לא כל שכן. מה כתיב בתריה מה רב טובך אשר צפנת ליריאיך, אל כמה רב טובך ע"כ. ומה הדורש במעשה בראשית כל, הדורש במעשה מרכבה עאכ"ו, וכ"ש וק"ו הקורש ברוכב ערבות שהוא למעלה ממעשה מרכבה שלא יהיה לו במה רב טובך.

| id | word | punctuation after |
|---|---|---|
| 9:19.1 | קלז | ) |
| 9:19.2 | ובמדרש |  |
| 9:19.3 | רבה |  |
| 9:19.4 | פרשת |  |
| 9:19.5 | קרח |  |
| 9:19.6 | דרשו |  |
| 9:19.7 | רז"ל |  |
| 9:19.8 | ששים |  |
| 9:19.9 | המה |  |
| 9:19.10 | מלכות | , |
| 9:19.11 | אלו |  |
| 9:19.12 | ס | ' |
| 9:19.13 | מסכתות |  |
| 9:19.14 | ושמנים |  |
| 9:19.15 | פילגשים |  |
| 9:19.16 | אלו |  |
| 9:19.17 | שמנים |  |
| 9:19.18 | בתי |  |
| 9:19.19 | מדרשות |  |
| 9:19.20 | שהיו |  |
| 9:19.21 | בירושלם |  |
| 9:19.22 | כנגד |  |
| 9:19.23 | פתחיה | , |
| 9:19.24 | ועלמות |  |
| 9:19.25 | אין |  |
| 9:19.26 | מספר | , |
| 9:19.27 | משנה |  |
| 9:19.28 | החיצונה | , |
| 9:19.29 | ופירש |  |
| 9:19.30 | הרב |  |
| 9:19.31 | עץ |  |
| 9:19.32 | יוסף |  |
| 9:19.33 | ומת"כ |  |
| 9:19.34 | דקרי |  |
| 9:19.35 | לבתי |  |
| 9:19.36 | מדרשות |  |
| 9:19.37 | פילגשים |  |
| 9:19.38 | על |  |
| 9:19.39 | שם |  |
| 9:19.40 | שנושאין |  |
| 9:19.41 | ונותנים |  |
| 9:19.42 | בהם |  |
| 9:19.43 | וחולקים |  |
| 9:19.44 | בהם |  |
| 9:19.45 | וחולקים |  |
| 9:19.46 | זה |  |
| 9:19.47 | עם |  |
| 9:19.48 | זה |  |
| 9:19.49 | כמו |  |
| 9:19.50 | פילגשם |  |
| 9:19.51 | ומחבר |  |
| 9:19.52 | הזהר |  |
| 9:19.53 | כשנאתו |  |
| 9:19.54 | למשנה |  |
| 9:19.55 | ולתלמוד |  |
| 9:19.56 | דימה |  |
| 9:19.57 | שקלא |  |
| 9:19.58 | וטריא |  |
| 9:19.59 | דרבנן |  |
| 9:19.60 | לכלביך |  |
| 9:19.61 | דצוחון |  |
| 9:19.62 | הב |  |
| 9:19.63 | הב |  |
| 9:19.64 | ושם |  |
| 9:19.65 | אותם |  |
| 9:19.66 | מיורשי |  |
| 9:19.67 | גהינם |  |
| 9:19.68 | האומרת |  |
| 9:19.69 | הב |  |
| 9:19.70 | הב | ) |
| 9:19.71 | ופי | ' |
| 9:19.72 | משנה |  |
| 9:19.73 | החיצונה |  |
| 9:19.74 | מה |  |
| 9:19.75 | שלמדו |  |
| 9:19.76 | היחידים |  |
| 9:19.77 | חוץ |  |
| 9:19.78 | לאלו |  |
| 9:19.79 | בתי |  |
| 9:19.80 | מדרשות | , |
| 9:19.81 | וביד |  |
| 9:19.82 | יוסף |  |
| 9:19.83 | פי | ' |
| 9:19.84 | אלו |  |
| 9:19.85 | הברייתות | , |
| 9:19.86 | ע"כ | , |
| 9:19.87 | וכן |  |
| 9:19.88 | מצאתי |  |
| 9:19.89 | בשיר |  |
| 9:19.90 | השירים |  |
| 9:19.91 | רבה |  |
| 9:19.92 | וז"ל |  |
| 9:19.93 | ששים |  |
| 9:19.94 | המה |  |
| 9:19.95 | מלכות | , |
| 9:19.96 | אלו |  |
| 9:19.97 | ששים |  |
| 9:19.98 | מסכתות |  |
| 9:19.99 | של |  |
| 9:19.100 | הלכות |  |
| 9:19.101 | ושמנים |  |
| 9:19.102 | פילגשים |  |
| 9:19.103 | אלו |  |
| 9:19.104 | שמנים |  |
| 9:19.105 | פרשיות |  |
| 9:19.106 | שבתורת |  |
| 9:19.107 | כהנים | . |
| 9:19.108 | ועלמות |  |
| 9:19.109 | אין |  |
| 9:19.110 | מספר | , |
| 9:19.111 | אין |  |
| 9:19.112 | קץ |  |
| 9:19.113 | לתוספתות | , |
| 9:19.114 | אחת |  |
| 9:19.115 | היא |  |
| 9:19.116 | הן |  |
| 9:19.117 | חולקין |  |
| 9:19.118 | אלו |  |
| 9:19.119 | עם |  |
| 9:19.120 | אלו |  |
| 9:19.121 | וכולהון |  |
| 9:19.122 | דורשין |  |
| 9:19.123 | מטעם |  |
| 9:19.124 | אחד | , |
| 9:19.125 | ומהלכה |  |
| 9:19.126 | אחת |  |
| 9:19.127 | מגזרה |  |
| 9:19.128 | שוה |  |
| 9:19.129 | מקל |  |
| 9:19.130 | וחומר |  |
| 9:19.131 | ע"כ | . |
| 9:19.132 | ראה |  |
| 9:19.133 | ידידי |  |
| 9:19.134 | הקורא |  |
| 9:19.135 | כמה |  |
| 9:19.136 | מחבבים |  |
| 9:19.137 | רז"ל |  |
| 9:19.138 | את |  |
| 9:19.139 | המשנה |  |
| 9:19.140 | והתלמוד |  |
| 9:19.141 | וקראו |  |
| 9:19.142 | להן |  |
| 9:19.143 | מלכות |  |
| 9:19.144 | ולא |  |
| 9:19.145 | שפחות | , |
| 9:19.146 | כמאמר |  |
| 9:19.147 | שלמה |  |
| 9:19.148 | כי |  |
| 9:19.149 | מלכים |  |
| 9:19.150 | ימלוכו | , |
| 9:19.151 | והמה |  |
| 9:19.152 | בתרה |  |
| 9:19.153 | של |  |
| 9:19.154 | תורה |  |
| 9:19.155 | כי |  |
| 9:19.156 | לא |  |
| 9:19.157 | נוכל |  |
| 9:19.158 | לעשות |  |
| 9:19.159 | דבר |  |
| 9:19.160 | ממצות |  |
| 9:19.161 | התורה |  |
| 9:19.162 | כי |  |
| 9:19.163 | אם |  |
| 9:19.164 | על |  |
| 9:19.165 | פי |  |
| 9:19.166 | קבלת |  |
| 9:19.167 | רז"ל |  |
| 9:19.168 | בעלי |  |
| 9:19.169 | המשנה |  |
| 9:19.170 | והתלמודף |  |
| 9:19.171 | ולא |  |
| 9:19.172 | נסור |  |
| 9:19.173 | ימין |  |
| 9:19.174 | ושמאל |  |
| 9:19.175 | מקבלתם |  |
| 9:19.176 | בפירוש |  |
| 9:19.177 | המצות |  |
| 9:19.178 | היאך |  |
| 9:19.179 | נעשה |  |
| 9:19.180 | אותם |  |
| 9:19.181 | כמשפטם | , |
| 9:19.182 | וכ"ש |  |
| 9:19.183 | וק"ו |  |
| 9:19.184 | באמונת |  |
| 9:19.185 | אחדות |  |
| 9:19.186 | הבורא |  |
| 9:19.187 | ית | ' |
| 9:19.188 | ושלילת |  |
| 9:19.189 | מאורעות |  |
| 9:19.190 | הגוף |  |
| 9:19.191 | ומקריו |  |
| 9:19.192 | ממנו |  |
| 9:19.193 | ית | '. |
| 9:19.194 | מריה |  |
| 9:19.195 | דאברהם | , |
| 9:19.196 | לא |  |
| 9:19.197 | לנו |  |
| 9:19.198 | ה | ' |
| 9:19.199 | לא |  |
| 9:19.200 | לנו |  |
| 9:19.201 | לחשוב |  |
| 9:19.202 | רבוי |  |
| 9:19.203 | באלהותך |  |
| 9:19.204 | כאשר |  |
| 9:19.205 | עשו |  |
| 9:19.206 | המקובלים |  |
| 9:19.207 | החדשים |  |
| 9:19.208 | על |  |
| 9:19.209 | פי |  |
| 9:19.210 | הזהר |  |
| 9:19.211 | והתקונים |  |
| 9:19.212 | שמיחסים |  |
| 9:19.213 | אותם |  |
| 9:19.214 | לרשב"י |  |
| 9:19.215 | בשקר | , |
| 9:19.216 | כי |  |
| 9:19.217 | נתחו |  |
| 9:19.218 | לפי |  |
| 9:19.219 | עניות |  |
| 9:19.220 | דעתם | ) |
| 9:19.221 | את |  |
| 9:19.222 | ה | ' |
| 9:19.223 | אלהים |  |
| 9:19.224 | אמת |  |
| 9:19.225 | האל |  |
| 9:19.226 | הגדול |  |
| 9:19.227 | הגבור |  |
| 9:19.228 | והנורא |  |
| 9:19.229 | לנתחים |  |
| 9:19.230 | וצורות |  |
| 9:19.231 | ופרצופים |  |
| 9:19.232 | רבים |  |
| 9:19.233 | שונים |  |
| 9:19.234 | זה |  |
| 9:19.235 | מזה |  |
| 9:19.236 | באיכותן |  |
| 9:19.237 | וכמותן |  |
| 9:19.238 | וחומרן |  |
| 9:19.239 | וטבען | , |
| 9:19.240 | ויערכו |  |
| 9:19.241 | את |  |
| 9:19.242 | נתחיו |  |
| 9:19.243 | לעין |  |
| 9:19.244 | כל |  |
| 9:19.245 | קורא |  |
| 9:19.246 | בספריהם | , |
| 9:19.247 | וקורא |  |
| 9:19.248 | להם |  |
| 9:19.249 | מערכת |  |
| 9:19.250 | האלהות | , |
| 9:19.251 | ויקראו |  |
| 9:19.252 | לכל |  |
| 9:19.253 | נתח |  |
| 9:19.254 | בשם | , |
| 9:19.255 | כשם |  |
| 9:19.256 | הגדולים |  |
| 9:19.257 | אשר |  |
| 9:19.258 | בארץ |  |
| 9:19.259 | וזהו |  |
| 9:19.260 | סדר |  |
| 9:19.261 | עריכתם |  |
| 9:19.262 | כמבואר |  |
| 9:19.263 | בספרים | , |
| 9:19.264 | אין |  |
| 9:19.265 | סוף |  |
| 9:19.266 | העליון |  |
| 9:19.267 | אשר |  |
| 9:19.268 | אין |  |
| 9:19.269 | למעלה |  |
| 9:19.270 | ממנו | , |
| 9:19.271 | אדם |  |
| 9:19.272 | קדמאה | , |
| 9:19.273 | אדם |  |
| 9:19.274 | קדמון | , |
| 9:19.275 | עתיק |  |
| 9:19.276 | ונוקביה | , |
| 9:19.277 | אריך |  |
| 9:19.278 | אנפין |  |
| 9:19.279 | ונוקביה | . |
| 9:19.280 | אבא |  |
| 9:19.281 | ואימא |  |
| 9:19.282 | עילאין | , |
| 9:19.283 | ישראל |  |
| 9:19.284 | סבא |  |
| 9:19.285 | ותבונה | , |
| 9:19.286 | יעקב |  |
| 9:19.287 | ולאה |  |
| 9:19.288 | זעיר |  |
| 9:19.289 | ונוקביה | , |
| 9:19.290 | והוא |  |
| 9:19.291 | נקרא |  |
| 9:19.292 | ישראל |  |
| 9:19.293 | והיא |  |
| 9:19.294 | רחל | , |
| 9:19.295 | עץ |  |
| 9:19.296 | חיים |  |
| 9:19.297 | שער |  |
| 9:19.298 | המלכים |  |
| 9:19.299 | פ"ט | ). |
| 9:19.300 | ובשער |  |
| 9:19.301 | אבא |  |
| 9:19.302 | ואימא |  |
| 9:19.303 | פ"ב |  |
| 9:19.304 | כ | ' |
| 9:19.305 | וכבר |  |
| 9:19.306 | נודע |  |
| 9:19.307 | מ"ש |  |
| 9:19.308 | בענין |  |
| 9:19.309 | לאה |  |
| 9:19.310 | ורחל |  |
| 9:19.311 | בהיותן |  |
| 9:19.312 | עומדין |  |
| 9:19.313 | באחורי |  |
| 9:19.314 | זעיר |  |
| 9:19.315 | אנפין |  |
| 9:19.316 | שאז |  |
| 9:19.317 | העקביים |  |
| 9:19.318 | של |  |
| 9:19.319 | לאה |  |
| 9:19.320 | נכנסים |  |
| 9:19.321 | תוך |  |
| 9:19.322 | ראש |  |
| 9:19.323 | רחל |  |
| 9:19.324 | העומדת |  |
| 9:19.325 | תחתיה | , |
| 9:19.326 | וכן |  |
| 9:19.327 | הענין |  |
| 9:19.328 | בכאן |  |
| 9:19.329 | שעקביים |  |
| 9:19.330 | של |  |
| 9:19.331 | בינה |  |
| 9:19.332 | נכנסין |  |
| 9:19.333 | תוך |  |
| 9:19.334 | ראש |  |
| 9:19.335 | תבונה |  |
| 9:19.336 | עכ"ל | . |
| 9:19.337 | וישסעו |  |
| 9:19.338 | את |  |
| 9:19.339 | אריה |  |
| 9:19.340 | דבי |  |
| 9:19.341 | עלאי |  |
| 9:19.342 | אריה |  |
| 9:19.343 | שאג |  |
| 9:19.344 | מי |  |
| 9:19.345 | לא |  |
| 9:19.346 | יירא |  |
| 9:19.347 | ה | ' |
| 9:19.348 | אלהים |  |
| 9:19.349 | דבר | , |
| 9:19.350 | כשסע |  |
| 9:19.351 | הגדי |  |
| 9:19.352 | ידים |  |
| 9:19.353 | ורגלים | , |
| 9:19.354 | ועקביים |  |
| 9:19.355 | וחזה |  |
| 9:19.356 | ושוק |  |
| 9:19.357 | וטבור | , |
| 9:19.358 | ושאר |  |
| 9:19.359 | איברים |  |
| 9:19.360 | נסתרים |  |
| 9:19.361 | כגיד |  |
| 9:19.362 | הערוה |  |
| 9:19.363 | ותרין |  |
| 9:19.364 | ביעין |  |
| 9:19.365 | לכל |  |
| 9:19.366 | אחד |  |
| 9:19.367 | מן |  |
| 9:19.368 | הפרצופים | ! |
| 9:19.369 | ויראו |  |
| 9:19.370 | לאחרים |  |
| 9:19.371 | את |  |
| 9:19.372 | ערות |  |
| 9:19.373 | אבותיהם |  |
| 9:19.374 | שבשמים |  |
| 9:19.375 | לפי |  |
| 9:19.376 | דעתם |  |
| 9:19.377 | שחושבים |  |
| 9:19.378 | רבוי |  |
| 9:19.379 | פרצופים |  |
| 9:19.380 | כאל |  |
| 9:19.381 | אחד | ) |
| 9:19.382 | ויכניסו |  |
| 9:19.383 | אותם |  |
| 9:19.384 | במכחול |  |
| 9:19.385 | בשפופרת |  |
| 9:19.386 | ברחם |  |
| 9:19.387 | הנקבות |  |
| 9:19.388 | אשר |  |
| 9:19.389 | בדו |  |
| 9:19.390 | מלבם | , |
| 9:19.391 | כמבואר |  |
| 9:19.392 | בעץ |  |
| 9:19.393 | חיים |  |
| 9:19.394 | בכמה |  |
| 9:19.395 | מקומות |  |
| 9:19.396 | ושם |  |
| 9:19.397 | מבואר |  |
| 9:19.398 | שהיסוד |  |
| 9:19.399 | כולו |  |
| 9:19.400 | הוא |  |
| 9:19.401 | בחינת |  |
| 9:19.402 | פנים |  |
| 9:19.403 | ואין |  |
| 9:19.404 | לו |  |
| 9:19.405 | אחוריים |  |
| 9:19.406 | לפי |  |
| 9:19.407 | שכל |  |
| 9:19.408 | היסוד |  |
| 9:19.409 | של |  |
| 9:19.410 | זכר |  |
| 9:19.411 | נכנס |  |
| 9:19.412 | צוך |  |
| 9:19.413 | היסוד |  |
| 9:19.414 | של |  |
| 9:19.415 | נקבה | , |
| 9:19.416 | וכולו |  |
| 9:19.417 | בחינת |  |
| 9:19.418 | פנים | , |
| 9:19.419 | מה |  |
| 9:19.420 | שאין |  |
| 9:19.421 | כן |  |
| 9:19.422 | בשאר |  |
| 9:19.423 | הגוף |  |
| 9:19.424 | שיש |  |
| 9:19.425 | בו |  |
| 9:19.426 | בחינת |  |
| 9:19.427 | פנים |  |
| 9:19.428 | הנחזרים |  |
| 9:19.429 | נגד |  |
| 9:19.430 | פני |  |
| 9:19.431 | הנקבה | ! |
| 9:19.432 | אוי |  |
| 9:19.433 | לאותה |  |
| 9:19.434 | בושה |  |
| 9:19.435 | ואני |  |
| 9:19.436 | לאותה |  |
| 9:19.437 | כלמה |  |
| 9:19.438 | לספר |  |
| 9:19.439 | דברים |  |
| 9:19.440 | מכוערים |  |
| 9:19.441 | כאלה |  |
| 9:19.442 | באלהינו |  |
| 9:19.443 | שבשמים |  |
| 9:19.444 | אשר |  |
| 9:19.445 | מלא |  |
| 9:19.446 | כל |  |
| 9:19.447 | הארץ |  |
| 9:19.448 | כבודו | , |
| 9:19.449 | ולקרוא |  |
| 9:19.450 | לו |  |
| 9:19.451 | בשם |  |
| 9:19.452 | מגונה |  |
| 9:19.453 | קצר |  |
| 9:19.454 | אפים |  |
| 9:19.455 | ז"א | ). |
| 9:19.456 | וכבר |  |
| 9:19.457 | הודעתיך |  |
| 9:19.458 | בסי | ' |
| 9:19.459 | צ"ב |  |
| 9:19.460 | דברי |  |
| 9:19.461 | רבינו |  |
| 9:19.462 | הגדול |  |
| 9:19.463 | הרמב"ם |  |
| 9:19.464 | ז"ל |  |
| 9:19.465 | בספרו |  |
| 9:19.466 | הנכבד |  |
| 9:19.467 | מורה |  |
| 9:19.468 | נבוכים |  |
| 9:19.469 | שתורתינו |  |
| 9:19.470 | הקדושה |  |
| 9:19.471 | וכל |  |
| 9:19.472 | הנביאים |  |
| 9:19.473 | וחכמינו |  |
| 9:19.474 | ז"ל |  |
| 9:19.475 | הרחיקו |  |
| 9:19.476 | מהשי"ת | . |
| 9:19.477 | שני |  |
| 9:19.478 | חושים |  |
| 9:19.479 | מהחמשה |  |
| 9:19.480 | חושים |  |
| 9:19.481 | שבאדם | , |
| 9:19.482 | והן |  |
| 9:19.483 | הטעם |  |
| 9:19.484 | והמשוש | , |
| 9:19.485 | וכ"ש |  |
| 9:19.486 | המשגל |  |
| 9:19.487 | שהוא |  |
| 9:19.488 | הפחות |  |
| 9:19.489 | והמגונה |  |
| 9:19.490 | וכו | '. |
| 9:19.491 | וכבר |  |
| 9:19.492 | ידעת |  |
| 9:19.493 | מה |  |
| 9:19.494 | שדרשו |  |
| 9:19.495 | רז"ל |  |
| 9:19.496 | על |  |
| 9:19.497 | פסוק |  |
| 9:19.498 | ותמת |  |
| 9:19.499 | שם |  |
| 9:19.500 | מרים | , |
| 9:19.501 | תקבר |  |
| 9:19.502 | שם |  |
| 9:19.503 | דאמר |  |
| 9:19.504 | ר | ' |
| 9:19.505 | אלעזר |  |
| 9:19.506 | אף |  |
| 9:19.507 | מרים |  |
| 9:19.508 | מתה |  |
| 9:19.509 | בנשיקה | . |
| 9:19.510 | איתא |  |
| 9:19.511 | שם |  |
| 9:19.512 | שם | , |
| 9:19.513 | ומפני |  |
| 9:19.514 | מה |  |
| 9:19.515 | לא |  |
| 9:19.516 | נאמר |  |
| 9:19.517 | בה |  |
| 9:19.518 | על |  |
| 9:19.519 | פי |  |
| 9:19.520 | ה | ', |
| 9:19.521 | שגנאי |  |
| 9:19.522 | הדבר |  |
| 9:19.523 | לומר |  |
| 9:19.524 | ואם |  |
| 9:19.525 | כל |  |
| 9:19.526 | כך |  |
| 9:19.527 | הקפיד |  |
| 9:19.528 | הכתוב |  |
| 9:19.529 | שלא |  |
| 9:19.530 | לומר |  |
| 9:19.531 | במרים |  |
| 9:19.532 | הנביאה |  |
| 9:19.533 | על |  |
| 9:19.534 | פי |  |
| 9:19.535 | ה | ' |
| 9:19.536 | מטעם |  |
| 9:19.537 | שגנאי |  |
| 9:19.538 | הדבר |  |
| 9:19.539 | לומר |  |
| 9:19.540 | כן | . |
| 9:19.541 | כל |  |
| 9:19.542 | שכן |  |
| 9:19.543 | קל |  |
| 9:19.544 | וחומר |  |
| 9:19.545 | ליחס |  |
| 9:19.546 | לשי"ת |  |
| 9:19.547 | זכרות |  |
| 9:19.548 | ונקבות |  |
| 9:19.549 | וזווגין |  |
| 9:19.550 | והטפת |  |
| 9:19.551 | טפין |  |
| 9:19.552 | ולומר |  |
| 9:19.553 | שמקום |  |
| 9:19.554 | חיבור |  |
| 9:19.555 | גיד |  |
| 9:19.556 | הערוה |  |
| 9:19.557 | הוא |  |
| 9:19.558 | למעלה |  |
| 9:19.559 | מתרין |  |
| 9:19.560 | פרקין |  |
| 9:19.561 | נצח |  |
| 9:19.562 | והוד |  |
| 9:19.563 | כמו |  |
| 9:19.564 | שהוא |  |
| 9:19.565 | באדם |  |
| 9:19.566 | כמו |  |
| 9:19.567 | שכתב |  |
| 9:19.568 | מהרח"ו |  |
| 9:19.569 | בעץ |  |
| 9:19.570 | חיים |  |
| 9:19.571 | שער |  |
| 9:19.572 | נקודים |  |
| 9:19.573 | פ"ג | . |
| 9:19.574 | וכבר |  |
| 9:19.575 | הבאתי |  |
| 9:19.576 | לעיל |  |
| 9:19.577 | סי | ' |
| 9:19.578 | צ"ד |  |
| 9:19.579 | דברי |  |
| 9:19.580 | בראשית |  |
| 9:19.581 | רבה |  |
| 9:19.582 | ודברי |  |
| 9:19.583 | הירושלמי |  |
| 9:19.584 | על |  |
| 9:19.585 | פסוק |  |
| 9:19.586 | תאלמנה |  |
| 9:19.587 | שפתי |  |
| 9:19.588 | שקר |  |
| 9:19.589 | הדוברות |  |
| 9:19.590 | על |  |
| 9:19.591 | צדיק |  |
| 9:19.592 | עתק |  |
| 9:19.593 | בגאוה |  |
| 9:19.594 | ובוז | , |
| 9:19.595 | יתחרשן |  |
| 9:19.596 | יתפרכן |  |
| 9:19.597 | ישתתקן |  |
| 9:19.598 | וכו | ' |
| 9:19.599 | הדוברות |  |
| 9:19.600 | על |  |
| 9:19.601 | צדיק |  |
| 9:19.602 | עתק | . |
| 9:19.603 | על |  |
| 9:19.604 | צדיקו |  |
| 9:19.605 | של |  |
| 9:19.606 | עולם |  |
| 9:19.607 | דברים |  |
| 9:19.608 | שהעתיק |  |
| 9:19.609 | מבריותיו |  |
| 9:19.610 | בגאוה |  |
| 9:19.611 | ובוז | , |
| 9:19.612 | שהוא |  |
| 9:19.613 | מתגאה |  |
| 9:19.614 | לומר |  |
| 9:19.615 | אני |  |
| 9:19.616 | דורש |  |
| 9:19.617 | במעשה |  |
| 9:19.618 | בראשית | , |
| 9:19.619 | סבור |  |
| 9:19.620 | שהוא |  |
| 9:19.621 | כמגאה |  |
| 9:19.622 | ואינו |  |
| 9:19.623 | אלא |  |
| 9:19.624 | ככבזה | . |
| 9:19.625 | אמר |  |
| 9:19.626 | ר | ' |
| 9:19.627 | חנינה |  |
| 9:19.628 | המתכבד |  |
| 9:19.629 | בקלון |  |
| 9:19.630 | חבירו |  |
| 9:19.631 | אין |  |
| 9:19.632 | לו |  |
| 9:19.633 | חלק |  |
| 9:19.634 | לעולם |  |
| 9:19.635 | הבא | . |
| 9:19.636 | המתכבד |  |
| 9:19.637 | בכבוד |  |
| 9:19.638 | חי |  |
| 9:19.639 | העולמים |  |
| 9:19.640 | לא |  |
| 9:19.641 | כל |  |
| 9:19.642 | שכן | . |
| 9:19.643 | מה |  |
| 9:19.644 | כתיב |  |
| 9:19.645 | בתריה |  |
| 9:19.646 | מה |  |
| 9:19.647 | רב |  |
| 9:19.648 | טובך |  |
| 9:19.649 | אשר |  |
| 9:19.650 | צפנת |  |
| 9:19.651 | ליריאיך | , |
| 9:19.652 | אל |  |
| 9:19.653 | כמה |  |
| 9:19.654 | רב |  |
| 9:19.655 | טובך |  |
| 9:19.656 | ע"כ | . |
| 9:19.657 | ומה |  |
| 9:19.658 | הדורש |  |
| 9:19.659 | במעשה |  |
| 9:19.660 | בראשית |  |
| 9:19.661 | כל | , |
| 9:19.662 | הדורש |  |
| 9:19.663 | במעשה |  |
| 9:19.664 | מרכבה |  |
| 9:19.665 | עאכ"ו | , |
| 9:19.666 | וכ"ש |  |
| 9:19.667 | וק"ו |  |
| 9:19.668 | הקורש |  |
| 9:19.669 | ברוכב |  |
| 9:19.670 | ערבות |  |
| 9:19.671 | שהוא |  |
| 9:19.672 | למעלה |  |
| 9:19.673 | ממעשה |  |
| 9:19.674 | מרכבה |  |
| 9:19.675 | שלא |  |
| 9:19.676 | יהיה |  |
| 9:19.677 | לו |  |
| 9:19.678 | במה |  |
| 9:19.679 | רב |  |
| 9:19.680 | טובך | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

