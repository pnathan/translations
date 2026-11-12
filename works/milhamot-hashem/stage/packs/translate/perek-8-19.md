# Translation pack: milhamot-hashem perek-8-19 (Milchamot ha-Shem 8:23–8:23)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-19`.

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
- **ריבה / לרבות** → "include", never "multiply".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 8:23

Hebrew (Wikisource): קיו) וכבר מלאו אזניך ורעיונך, כי המקובלים עובדים לנברא כמו שנת' לעיל סי' מ"ח בשם יושר לבב דף ג' ע"ב וז"ל דע את אלהי אביך עבדהו, דע את אלהי אביך כולל ה' פרצופין, ועבדהו לז"א אע"פ שהוא נברא, כי נשמתו את העובד בזה כי מבלעדיו אין כל נמצא, משום הכי לא אמר דע אלהי אביך ולא את אלהי אביך ע"כ. וכוונתו כי במלת את לרבות נשמת זעיר אנפין. והנה בכל תורה לא מצינו שיהא דבר הבא מרבוי עוקר את העיקר שנזכר בכתוב בהדיא, רק שיהא מצורף אל העיקר ומשתתף באיסור עמו כמו שדרשו רז"ל בכתובות כבד את אביך ואת אמך, את אביך זו אשת אביך, ואת אמך זה בעל אמך, שיהא מכבד את אשת אביו עם אביו, ואת בעל אמו עם אמו, וכן מה שדרשו רז"ל וערלתם ערלתו את פריו את הטפל לפריו, דהיינו הקליפין והגרעינין שהן אסורין בערלה כפרי עצמו ובנדון זה עשו המקובלים החדשים דבר הטפל הבא מרבוי את לעיקר ואמרו שהם עובדים לנשמתו, והגוף שהוא נברא טפל. ועובדים וגוף ולנשמה, כמו שביארנו לעיל סי' ט"ל וסי' ס"ה! ונמצאו משתפים שם שמים ודבר נברא, ועוברים על דברי רשב"י שאמר כל המשתף שם שמים ודבר אחר נעקר מן העולם. והן עתה אוסיף להעירך על קצת מאמרים מהזהר דמוכח מנייהו שזעיר אנפין נברא עיין בזהר בראית דף ט"ו ומקדש מלך שם מבואר דז"א נברא הוא וכבר הבאתי קצת מזה בסי' כ"ד, גם בדף כ"ו ע"א ויטע ה' אלהים, אבא ואימא, גן דא שכינתא תתאה, עדן דא אימא עלאה, את האדם דא עמודא דאמצעיתא (שהוא ז"א) עוד שם ויצמח ה' אלהים אבא ואימא, כל עץ נחמד דא צדיק (יסוד) וטוב למאכל, דא עמודא דאמצעיתא וכו' ועץ החיים דהוא אילנא דחי יהי נטיע בנו גינתא וכו', והראיה היותר ברורא הוא האמור באדרא רבה דף קל"ח ע"ב, ה' פעלך בקרב שנים חייהו. האי לעתיק יומין אתאמר, מאן פעלך, זעיר אנפין, חייהו למאן? לזעיר אנפין דכל נהירו דיליה מאינון שנים קדמוניות! מזה יתבאר לך בבירור שזעיר אנפין נברא, וכבר ידעת דעת הפילסוף המסית מחבר הזהר ורוב המקובלים הנמשכים אחריו, שלא תכון תפלה ועבודה רק לזעיר אנפין הנברא הלזה, ולדעת בעל עז לאלהים אין העבודה אלא למלכא קדישא דכל קדישין הנקרא לבא דכל ליבין והוא מתלבש תוך זעיר אנפין כמבואר שם בבית קדש הקדשים פרק י"ט דף נ"ט ע"ד, ואע"פ שיצא ממעי אבא ואימא הוא העיקר ואף אבא ואימא צריכים לכבדו כמלך בשר ודם שמולך שכל העם מכבדים אותו. ואף אביו ואמו צריכים לכבדו, שהוא המלך השליט, ואליו תפלותינו והוא מלכינו ואלהינו, אבל המתפלל לאל עליון שהוא הסבה הראשונה אין תפלתו תפלה, כי איך נאמר ברוך אתה ה', ממי יתברך ויבוא לו השפע? (בית קדש הקדשים דף כ"ו ע"ד). ועיין בזהר רעיא מהימנא דף ק"ט ע"ב שמבאר שם, כי האלהים העליון גלית ליה גופא ולא איברים ולא נוקבא, והוא רק מחבר ומקשר את האלהים התחתונים בעלי גוף ואיברים זכרים ונקיבות עם נקבותיהם להיות אחד, כטיט המחבר אבני הבית, או כחשוקי ברזל ונחשת הקושרים ומחברים את לוחות חביות של עץ להיות אחד. יתעלה השי"ת ויתרומם מכל מה שחשבו בו הטפשים והסכלים! ולכן אמרו המקובלים החדשים שהמתפלל אליו אין תפלתו תפלה, וכל הקורא אליו לא יענהו, והוא מרומם ומסולק מכל ברכה שאין לו ממי להתברך, ולא שם לו שיגבילהו, כי הכח והממשלה להאלהים שלמטה ממנו שהם החמשה פרצופים הכוללים, והמה נקראים בשם הוי'ה ואדנות ואלהים, ועיקר שם המיוחד שם הוי'ה נקרא להאלוה קצר אפים (זעיר אנפין) עם זוגתו, והוא השולט על כל הברואים כמו שנת' לעיל בכמה מקומות.

| id | word | punctuation after |
|---|---|---|
| 8:23.1 | קיו | ) |
| 8:23.2 | וכבר |  |
| 8:23.3 | מלאו |  |
| 8:23.4 | אזניך |  |
| 8:23.5 | ורעיונך | , |
| 8:23.6 | כי |  |
| 8:23.7 | המקובלים |  |
| 8:23.8 | עובדים |  |
| 8:23.9 | לנברא |  |
| 8:23.10 | כמו |  |
| 8:23.11 | שנת | ' |
| 8:23.12 | לעיל |  |
| 8:23.13 | סי | ' |
| 8:23.14 | מ"ח |  |
| 8:23.15 | בשם |  |
| 8:23.16 | יושר |  |
| 8:23.17 | לבב |  |
| 8:23.18 | דף |  |
| 8:23.19 | ג | ' |
| 8:23.20 | ע"ב |  |
| 8:23.21 | וז"ל |  |
| 8:23.22 | דע |  |
| 8:23.23 | את |  |
| 8:23.24 | אלהי |  |
| 8:23.25 | אביך |  |
| 8:23.26 | עבדהו | , |
| 8:23.27 | דע |  |
| 8:23.28 | את |  |
| 8:23.29 | אלהי |  |
| 8:23.30 | אביך |  |
| 8:23.31 | כולל |  |
| 8:23.32 | ה | ' |
| 8:23.33 | פרצופין | , |
| 8:23.34 | ועבדהו |  |
| 8:23.35 | לז"א |  |
| 8:23.36 | אע"פ |  |
| 8:23.37 | שהוא |  |
| 8:23.38 | נברא | , |
| 8:23.39 | כי |  |
| 8:23.40 | נשמתו |  |
| 8:23.41 | את |  |
| 8:23.42 | העובד |  |
| 8:23.43 | בזה |  |
| 8:23.44 | כי |  |
| 8:23.45 | מבלעדיו |  |
| 8:23.46 | אין |  |
| 8:23.47 | כל |  |
| 8:23.48 | נמצא | , |
| 8:23.49 | משום |  |
| 8:23.50 | הכי |  |
| 8:23.51 | לא |  |
| 8:23.52 | אמר |  |
| 8:23.53 | דע |  |
| 8:23.54 | אלהי |  |
| 8:23.55 | אביך |  |
| 8:23.56 | ולא |  |
| 8:23.57 | את |  |
| 8:23.58 | אלהי |  |
| 8:23.59 | אביך |  |
| 8:23.60 | ע"כ | . |
| 8:23.61 | וכוונתו |  |
| 8:23.62 | כי |  |
| 8:23.63 | במלת |  |
| 8:23.64 | את |  |
| 8:23.65 | לרבות |  |
| 8:23.66 | נשמת |  |
| 8:23.67 | זעיר |  |
| 8:23.68 | אנפין | . |
| 8:23.69 | והנה |  |
| 8:23.70 | בכל |  |
| 8:23.71 | תורה |  |
| 8:23.72 | לא |  |
| 8:23.73 | מצינו |  |
| 8:23.74 | שיהא |  |
| 8:23.75 | דבר |  |
| 8:23.76 | הבא |  |
| 8:23.77 | מרבוי |  |
| 8:23.78 | עוקר |  |
| 8:23.79 | את |  |
| 8:23.80 | העיקר |  |
| 8:23.81 | שנזכר |  |
| 8:23.82 | בכתוב |  |
| 8:23.83 | בהדיא | , |
| 8:23.84 | רק |  |
| 8:23.85 | שיהא |  |
| 8:23.86 | מצורף |  |
| 8:23.87 | אל |  |
| 8:23.88 | העיקר |  |
| 8:23.89 | ומשתתף |  |
| 8:23.90 | באיסור |  |
| 8:23.91 | עמו |  |
| 8:23.92 | כמו |  |
| 8:23.93 | שדרשו |  |
| 8:23.94 | רז"ל |  |
| 8:23.95 | בכתובות |  |
| 8:23.96 | כבד |  |
| 8:23.97 | את |  |
| 8:23.98 | אביך |  |
| 8:23.99 | ואת |  |
| 8:23.100 | אמך | , |
| 8:23.101 | את |  |
| 8:23.102 | אביך |  |
| 8:23.103 | זו |  |
| 8:23.104 | אשת |  |
| 8:23.105 | אביך | , |
| 8:23.106 | ואת |  |
| 8:23.107 | אמך |  |
| 8:23.108 | זה |  |
| 8:23.109 | בעל |  |
| 8:23.110 | אמך | , |
| 8:23.111 | שיהא |  |
| 8:23.112 | מכבד |  |
| 8:23.113 | את |  |
| 8:23.114 | אשת |  |
| 8:23.115 | אביו |  |
| 8:23.116 | עם |  |
| 8:23.117 | אביו | , |
| 8:23.118 | ואת |  |
| 8:23.119 | בעל |  |
| 8:23.120 | אמו |  |
| 8:23.121 | עם |  |
| 8:23.122 | אמו | , |
| 8:23.123 | וכן |  |
| 8:23.124 | מה |  |
| 8:23.125 | שדרשו |  |
| 8:23.126 | רז"ל |  |
| 8:23.127 | וערלתם |  |
| 8:23.128 | ערלתו |  |
| 8:23.129 | את |  |
| 8:23.130 | פריו |  |
| 8:23.131 | את |  |
| 8:23.132 | הטפל |  |
| 8:23.133 | לפריו | , |
| 8:23.134 | דהיינו |  |
| 8:23.135 | הקליפין |  |
| 8:23.136 | והגרעינין |  |
| 8:23.137 | שהן |  |
| 8:23.138 | אסורין |  |
| 8:23.139 | בערלה |  |
| 8:23.140 | כפרי |  |
| 8:23.141 | עצמו |  |
| 8:23.142 | ובנדון |  |
| 8:23.143 | זה |  |
| 8:23.144 | עשו |  |
| 8:23.145 | המקובלים |  |
| 8:23.146 | החדשים |  |
| 8:23.147 | דבר |  |
| 8:23.148 | הטפל |  |
| 8:23.149 | הבא |  |
| 8:23.150 | מרבוי |  |
| 8:23.151 | את |  |
| 8:23.152 | לעיקר |  |
| 8:23.153 | ואמרו |  |
| 8:23.154 | שהם |  |
| 8:23.155 | עובדים |  |
| 8:23.156 | לנשמתו | , |
| 8:23.157 | והגוף |  |
| 8:23.158 | שהוא |  |
| 8:23.159 | נברא |  |
| 8:23.160 | טפל | . |
| 8:23.161 | ועובדים |  |
| 8:23.162 | וגוף |  |
| 8:23.163 | ולנשמה | , |
| 8:23.164 | כמו |  |
| 8:23.165 | שביארנו |  |
| 8:23.166 | לעיל |  |
| 8:23.167 | סי | ' |
| 8:23.168 | ט"ל |  |
| 8:23.169 | וסי | ' |
| 8:23.170 | ס"ה | ! |
| 8:23.171 | ונמצאו |  |
| 8:23.172 | משתפים |  |
| 8:23.173 | שם |  |
| 8:23.174 | שמים |  |
| 8:23.175 | ודבר |  |
| 8:23.176 | נברא | , |
| 8:23.177 | ועוברים |  |
| 8:23.178 | על |  |
| 8:23.179 | דברי |  |
| 8:23.180 | רשב"י |  |
| 8:23.181 | שאמר |  |
| 8:23.182 | כל |  |
| 8:23.183 | המשתף |  |
| 8:23.184 | שם |  |
| 8:23.185 | שמים |  |
| 8:23.186 | ודבר |  |
| 8:23.187 | אחר |  |
| 8:23.188 | נעקר |  |
| 8:23.189 | מן |  |
| 8:23.190 | העולם | . |
| 8:23.191 | והן |  |
| 8:23.192 | עתה |  |
| 8:23.193 | אוסיף |  |
| 8:23.194 | להעירך |  |
| 8:23.195 | על |  |
| 8:23.196 | קצת |  |
| 8:23.197 | מאמרים |  |
| 8:23.198 | מהזהר |  |
| 8:23.199 | דמוכח |  |
| 8:23.200 | מנייהו |  |
| 8:23.201 | שזעיר |  |
| 8:23.202 | אנפין |  |
| 8:23.203 | נברא |  |
| 8:23.204 | עיין |  |
| 8:23.205 | בזהר |  |
| 8:23.206 | בראית |  |
| 8:23.207 | דף |  |
| 8:23.208 | ט"ו |  |
| 8:23.209 | ומקדש |  |
| 8:23.210 | מלך |  |
| 8:23.211 | שם |  |
| 8:23.212 | מבואר |  |
| 8:23.213 | דז"א |  |
| 8:23.214 | נברא |  |
| 8:23.215 | הוא |  |
| 8:23.216 | וכבר |  |
| 8:23.217 | הבאתי |  |
| 8:23.218 | קצת |  |
| 8:23.219 | מזה |  |
| 8:23.220 | בסי | ' |
| 8:23.221 | כ"ד | , |
| 8:23.222 | גם |  |
| 8:23.223 | בדף |  |
| 8:23.224 | כ"ו |  |
| 8:23.225 | ע"א |  |
| 8:23.226 | ויטע |  |
| 8:23.227 | ה | ' |
| 8:23.228 | אלהים | , |
| 8:23.229 | אבא |  |
| 8:23.230 | ואימא | , |
| 8:23.231 | גן |  |
| 8:23.232 | דא |  |
| 8:23.233 | שכינתא |  |
| 8:23.234 | תתאה | , |
| 8:23.235 | עדן |  |
| 8:23.236 | דא |  |
| 8:23.237 | אימא |  |
| 8:23.238 | עלאה | , |
| 8:23.239 | את |  |
| 8:23.240 | האדם |  |
| 8:23.241 | דא |  |
| 8:23.242 | עמודא |  |
| 8:23.243 | דאמצעיתא |  |
| 8:23.244 | שהוא |  |
| 8:23.245 | ז"א | ) |
| 8:23.246 | עוד |  |
| 8:23.247 | שם |  |
| 8:23.248 | ויצמח |  |
| 8:23.249 | ה | ' |
| 8:23.250 | אלהים |  |
| 8:23.251 | אבא |  |
| 8:23.252 | ואימא | , |
| 8:23.253 | כל |  |
| 8:23.254 | עץ |  |
| 8:23.255 | נחמד |  |
| 8:23.256 | דא |  |
| 8:23.257 | צדיק |  |
| 8:23.258 | יסוד | ) |
| 8:23.259 | וטוב |  |
| 8:23.260 | למאכל | , |
| 8:23.261 | דא |  |
| 8:23.262 | עמודא |  |
| 8:23.263 | דאמצעיתא |  |
| 8:23.264 | וכו | ' |
| 8:23.265 | ועץ |  |
| 8:23.266 | החיים |  |
| 8:23.267 | דהוא |  |
| 8:23.268 | אילנא |  |
| 8:23.269 | דחי |  |
| 8:23.270 | יהי |  |
| 8:23.271 | נטיע |  |
| 8:23.272 | בנו |  |
| 8:23.273 | גינתא |  |
| 8:23.274 | וכו | ', |
| 8:23.275 | והראיה |  |
| 8:23.276 | היותר |  |
| 8:23.277 | ברורא |  |
| 8:23.278 | הוא |  |
| 8:23.279 | האמור |  |
| 8:23.280 | באדרא |  |
| 8:23.281 | רבה |  |
| 8:23.282 | דף |  |
| 8:23.283 | קל"ח |  |
| 8:23.284 | ע"ב | , |
| 8:23.285 | ה | ' |
| 8:23.286 | פעלך |  |
| 8:23.287 | בקרב |  |
| 8:23.288 | שנים |  |
| 8:23.289 | חייהו | . |
| 8:23.290 | האי |  |
| 8:23.291 | לעתיק |  |
| 8:23.292 | יומין |  |
| 8:23.293 | אתאמר | , |
| 8:23.294 | מאן |  |
| 8:23.295 | פעלך | , |
| 8:23.296 | זעיר |  |
| 8:23.297 | אנפין | , |
| 8:23.298 | חייהו |  |
| 8:23.299 | למאן | ? |
| 8:23.300 | לזעיר |  |
| 8:23.301 | אנפין |  |
| 8:23.302 | דכל |  |
| 8:23.303 | נהירו |  |
| 8:23.304 | דיליה |  |
| 8:23.305 | מאינון |  |
| 8:23.306 | שנים |  |
| 8:23.307 | קדמוניות | ! |
| 8:23.308 | מזה |  |
| 8:23.309 | יתבאר |  |
| 8:23.310 | לך |  |
| 8:23.311 | בבירור |  |
| 8:23.312 | שזעיר |  |
| 8:23.313 | אנפין |  |
| 8:23.314 | נברא | , |
| 8:23.315 | וכבר |  |
| 8:23.316 | ידעת |  |
| 8:23.317 | דעת |  |
| 8:23.318 | הפילסוף |  |
| 8:23.319 | המסית |  |
| 8:23.320 | מחבר |  |
| 8:23.321 | הזהר |  |
| 8:23.322 | ורוב |  |
| 8:23.323 | המקובלים |  |
| 8:23.324 | הנמשכים |  |
| 8:23.325 | אחריו | , |
| 8:23.326 | שלא |  |
| 8:23.327 | תכון |  |
| 8:23.328 | תפלה |  |
| 8:23.329 | ועבודה |  |
| 8:23.330 | רק |  |
| 8:23.331 | לזעיר |  |
| 8:23.332 | אנפין |  |
| 8:23.333 | הנברא |  |
| 8:23.334 | הלזה | , |
| 8:23.335 | ולדעת |  |
| 8:23.336 | בעל |  |
| 8:23.337 | עז |  |
| 8:23.338 | לאלהים |  |
| 8:23.339 | אין |  |
| 8:23.340 | העבודה |  |
| 8:23.341 | אלא |  |
| 8:23.342 | למלכא |  |
| 8:23.343 | קדישא |  |
| 8:23.344 | דכל |  |
| 8:23.345 | קדישין |  |
| 8:23.346 | הנקרא |  |
| 8:23.347 | לבא |  |
| 8:23.348 | דכל |  |
| 8:23.349 | ליבין |  |
| 8:23.350 | והוא |  |
| 8:23.351 | מתלבש |  |
| 8:23.352 | תוך |  |
| 8:23.353 | זעיר |  |
| 8:23.354 | אנפין |  |
| 8:23.355 | כמבואר |  |
| 8:23.356 | שם |  |
| 8:23.357 | בבית |  |
| 8:23.358 | קדש |  |
| 8:23.359 | הקדשים |  |
| 8:23.360 | פרק |  |
| 8:23.361 | י"ט |  |
| 8:23.362 | דף |  |
| 8:23.363 | נ"ט |  |
| 8:23.364 | ע"ד | , |
| 8:23.365 | ואע"פ |  |
| 8:23.366 | שיצא |  |
| 8:23.367 | ממעי |  |
| 8:23.368 | אבא |  |
| 8:23.369 | ואימא |  |
| 8:23.370 | הוא |  |
| 8:23.371 | העיקר |  |
| 8:23.372 | ואף |  |
| 8:23.373 | אבא |  |
| 8:23.374 | ואימא |  |
| 8:23.375 | צריכים |  |
| 8:23.376 | לכבדו |  |
| 8:23.377 | כמלך |  |
| 8:23.378 | בשר |  |
| 8:23.379 | ודם |  |
| 8:23.380 | שמולך |  |
| 8:23.381 | שכל |  |
| 8:23.382 | העם |  |
| 8:23.383 | מכבדים |  |
| 8:23.384 | אותו | . |
| 8:23.385 | ואף |  |
| 8:23.386 | אביו |  |
| 8:23.387 | ואמו |  |
| 8:23.388 | צריכים |  |
| 8:23.389 | לכבדו | , |
| 8:23.390 | שהוא |  |
| 8:23.391 | המלך |  |
| 8:23.392 | השליט | , |
| 8:23.393 | ואליו |  |
| 8:23.394 | תפלותינו |  |
| 8:23.395 | והוא |  |
| 8:23.396 | מלכינו |  |
| 8:23.397 | ואלהינו | , |
| 8:23.398 | אבל |  |
| 8:23.399 | המתפלל |  |
| 8:23.400 | לאל |  |
| 8:23.401 | עליון |  |
| 8:23.402 | שהוא |  |
| 8:23.403 | הסבה |  |
| 8:23.404 | הראשונה |  |
| 8:23.405 | אין |  |
| 8:23.406 | תפלתו |  |
| 8:23.407 | תפלה | , |
| 8:23.408 | כי |  |
| 8:23.409 | איך |  |
| 8:23.410 | נאמר |  |
| 8:23.411 | ברוך |  |
| 8:23.412 | אתה |  |
| 8:23.413 | ה | ', |
| 8:23.414 | ממי |  |
| 8:23.415 | יתברך |  |
| 8:23.416 | ויבוא |  |
| 8:23.417 | לו |  |
| 8:23.418 | השפע | ? |
| 8:23.419 | בית |  |
| 8:23.420 | קדש |  |
| 8:23.421 | הקדשים |  |
| 8:23.422 | דף |  |
| 8:23.423 | כ"ו |  |
| 8:23.424 | ע"ד | ). |
| 8:23.425 | ועיין |  |
| 8:23.426 | בזהר |  |
| 8:23.427 | רעיא |  |
| 8:23.428 | מהימנא |  |
| 8:23.429 | דף |  |
| 8:23.430 | ק"ט |  |
| 8:23.431 | ע"ב |  |
| 8:23.432 | שמבאר |  |
| 8:23.433 | שם | , |
| 8:23.434 | כי |  |
| 8:23.435 | האלהים |  |
| 8:23.436 | העליון |  |
| 8:23.437 | גלית |  |
| 8:23.438 | ליה |  |
| 8:23.439 | גופא |  |
| 8:23.440 | ולא |  |
| 8:23.441 | איברים |  |
| 8:23.442 | ולא |  |
| 8:23.443 | נוקבא | , |
| 8:23.444 | והוא |  |
| 8:23.445 | רק |  |
| 8:23.446 | מחבר |  |
| 8:23.447 | ומקשר |  |
| 8:23.448 | את |  |
| 8:23.449 | האלהים |  |
| 8:23.450 | התחתונים |  |
| 8:23.451 | בעלי |  |
| 8:23.452 | גוף |  |
| 8:23.453 | ואיברים |  |
| 8:23.454 | זכרים |  |
| 8:23.455 | ונקיבות |  |
| 8:23.456 | עם |  |
| 8:23.457 | נקבותיהם |  |
| 8:23.458 | להיות |  |
| 8:23.459 | אחד | , |
| 8:23.460 | כטיט |  |
| 8:23.461 | המחבר |  |
| 8:23.462 | אבני |  |
| 8:23.463 | הבית | , |
| 8:23.464 | או |  |
| 8:23.465 | כחשוקי |  |
| 8:23.466 | ברזל |  |
| 8:23.467 | ונחשת |  |
| 8:23.468 | הקושרים |  |
| 8:23.469 | ומחברים |  |
| 8:23.470 | את |  |
| 8:23.471 | לוחות |  |
| 8:23.472 | חביות |  |
| 8:23.473 | של |  |
| 8:23.474 | עץ |  |
| 8:23.475 | להיות |  |
| 8:23.476 | אחד | . |
| 8:23.477 | יתעלה |  |
| 8:23.478 | השי"ת |  |
| 8:23.479 | ויתרומם |  |
| 8:23.480 | מכל |  |
| 8:23.481 | מה |  |
| 8:23.482 | שחשבו |  |
| 8:23.483 | בו |  |
| 8:23.484 | הטפשים |  |
| 8:23.485 | והסכלים | ! |
| 8:23.486 | ולכן |  |
| 8:23.487 | אמרו |  |
| 8:23.488 | המקובלים |  |
| 8:23.489 | החדשים |  |
| 8:23.490 | שהמתפלל |  |
| 8:23.491 | אליו |  |
| 8:23.492 | אין |  |
| 8:23.493 | תפלתו |  |
| 8:23.494 | תפלה | , |
| 8:23.495 | וכל |  |
| 8:23.496 | הקורא |  |
| 8:23.497 | אליו |  |
| 8:23.498 | לא |  |
| 8:23.499 | יענהו | , |
| 8:23.500 | והוא |  |
| 8:23.501 | מרומם |  |
| 8:23.502 | ומסולק |  |
| 8:23.503 | מכל |  |
| 8:23.504 | ברכה |  |
| 8:23.505 | שאין |  |
| 8:23.506 | לו |  |
| 8:23.507 | ממי |  |
| 8:23.508 | להתברך | , |
| 8:23.509 | ולא |  |
| 8:23.510 | שם |  |
| 8:23.511 | לו |  |
| 8:23.512 | שיגבילהו | , |
| 8:23.513 | כי |  |
| 8:23.514 | הכח |  |
| 8:23.515 | והממשלה |  |
| 8:23.516 | להאלהים |  |
| 8:23.517 | שלמטה |  |
| 8:23.518 | ממנו |  |
| 8:23.519 | שהם |  |
| 8:23.520 | החמשה |  |
| 8:23.521 | פרצופים |  |
| 8:23.522 | הכוללים | , |
| 8:23.523 | והמה |  |
| 8:23.524 | נקראים |  |
| 8:23.525 | בשם |  |
| 8:23.526 | הוי'ה |  |
| 8:23.527 | ואדנות |  |
| 8:23.528 | ואלהים | , |
| 8:23.529 | ועיקר |  |
| 8:23.530 | שם |  |
| 8:23.531 | המיוחד |  |
| 8:23.532 | שם |  |
| 8:23.533 | הוי'ה |  |
| 8:23.534 | נקרא |  |
| 8:23.535 | להאלוה |  |
| 8:23.536 | קצר |  |
| 8:23.537 | אפים |  |
| 8:23.538 | זעיר |  |
| 8:23.539 | אנפין | ) |
| 8:23.540 | עם |  |
| 8:23.541 | זוגתו | , |
| 8:23.542 | והוא |  |
| 8:23.543 | השולט |  |
| 8:23.544 | על |  |
| 8:23.545 | כל |  |
| 8:23.546 | הברואים |  |
| 8:23.547 | כמו |  |
| 8:23.548 | שנת | ' |
| 8:23.549 | לעיל |  |
| 8:23.550 | בכמה |  |
| 8:23.551 | מקומות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

