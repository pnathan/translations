# Translation pack: milhamot-hashem perek-5-23 (Milchamot ha-Shem 5:32–5:32)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-23.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-23`.

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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 5:32

Hebrew (Wikisource): נח) ובסנהדרין פ' ד' מיתות גרסינן אמר ר' יהודה אלמלא וי"ו שבהעלוך נתחייבו שונאיהם של ישראל כלייה. אמר לו רשב"י והלא כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתי לה' לבדו אלא מה ת"ל אשר העלוך שאיוו לאלוהות הרבה ע"כ, ומביאו מעשה רקח בפרק ב' מהל' ע"ז. ובפ' לולב וערבה אמתניתין דהיו אומרין ליה ולך מזבח פריך תלמודא והתניא כל המשתף שם שמים ודבר אחד נעקר מן העולם? אלא הכי קאמרי וכו' וכ' הרב מהרש"א בפ' ד' מיתות דודאי לכ"ע גרע טפי כופר בעיקר וחייב כלייה, אלא דרשב"י קאמר שכל המשתף שם שמים ודבר אחר דהיינו המאמין נמי בדבר אחר חייב כלייה ונעקר מן העולם, אלא מה ת"ל אשר העלוך דודאי האמינו בנמצא וכו' ולכך לא נעקרו אלא שאיוו לאלוהות הרבה ואמרו שהאחד הנמצא נתן כח לאלוהות וכוחות רבים עזב ה' את הארץ. ואלו הכוחות והאלוהות יהיו מנהיגים בארץ. וכן כ' הרמב"ם שבה דעת רבים מעובדי ע"ז שעובדים אותה. שאומרים שניתנה ההנהגה בידה. ועל דרך זה יש לכוון מה שאמר משה ראה אתה אומר אלי העל את העם הזה. שבא משה להתנצל בעדם. שאמר משה להקב"ה שהרי אתה אומר אלי העל את העם הזה. וכוונתך שאתה מעמיד אותי למנהיג להם, ותלית העלייה בי, כיוצא בזה כיוונו כלל לע"ז ולשתף שם שמים ודבר אחר עכ"ל וכל זה דוחק גדול. וכ' מעשה רקח דאפשר דשיתוף דקאמר היינו לכל האלוהות עם הקב"ה ע"ש ופשוט הוא. מכל זה נתבאר לך ידידי האוהב את תורת האמת ודברי דודים מעתיקיה ומקבליה ממשה רבי' איש מפי איש, שזה שאמרו המקובלים החדשים כי אלהינו בהו"ש בהיותו בבחינתו הפשוטה לא שייך בו שום עבודה ותפלה וקריאה כלל, כי קצור קצרה ידו ח"ו מפדות ואין בו כח להציל את הקורא ומשוע אליו ורק בהיותו משתתף בפרצופים הנז' וביחוד פרצוף זעיר אנפין אז שייך בו עבודה וברכה. וצריך דוקא לאחדו ולשתפו עם נוקבא דילה ועם שאר החמשה פרצופים, ואז בהיותו משתתף עמהם אז המה עונים את כל הקורא אליהם, וזהו ממש שיתוף שם שמים ודבר אחר, כי לבדו אינו יכול לעשות מאומה, ורק בשיתופו עמהם הוא יכול לפעול את פעולותיו, ואין לך שיתוף שם שמים ודבר אחר יותר מזה. ומאי דעל רשב"י סני בקבלה האמיתית בתלמוד, עשה רשב"י המשקר מחבר ס' הזהר בקבלה המזויפת, רחמנא ליצלן מאמונות זרות כאלה, והאמונה הברורה הטהורה מכל סיג וחלאה היא להאמין באמונה שלימה תמימה כי לו לבדו ית' הכח והממשלה בכל הנבראים שהמציא בחסדו ולו ית' נכוין בכל עבודותינו ותפלותינו כמצווה עלינו בתורה בלתי לה' לבדו שהוא הסבה הראשונה בלי צרוף ושיתוף שום כח אחר לא עתיק ונוקביה, ולא אריך ונוקביה, ולא אבא ואימא, ולא זעיר ונוקביה, כי כל הנמצאים צריכים לו והוא אינו צריך להם, וכן כתב הראב"ע בפי' החומש בלתי לה' לבדו ולא יתערב אחר עמו ע"כ, ושמעתי קצת טפשים אמרו דלא מיקרי משתף שם שמים ודבר אחר אלא דוקא כשמוציא בפיו דומיא דמתניתין דקאמרי ליה ולך מזבח, שעליה הקשו בגמרא והא משתף שם שמים ודבר אחר וכו'. ואין זה אמת דאפי' לכוון במחשבתו בהזכירו את השם בברכה או בתפלה לשום נברא או נאצל אף אם מכיון לשתף כח הניתן בו מאת השי"ת אסור, והרי זה עובד ע"ז, דמחשבת ע"ז הקב"ה מצרפה למעשה כדגרסינן בקידושין ד' ל"ט ובחולין קמ"ב ואלא הא דכתיב למען תפוש את בית ישראל בלבם ההוא כמחשבת ע"ז הוא דכתיב דאפי' חרהור הוי כמעשה. ובנ"ד כיון שמכוין במחשבתו לזעיר אנפין הרי זה כעובד ע"ז. וקרינא ביה מה שכ' הרמב"ם בפרק ג' לראשון קרוב אתה בפיהם ורחוק מכליותיהם.

| id | word | punctuation after |
|---|---|---|
| 5:32.1 | נח | ) |
| 5:32.2 | ובסנהדרין |  |
| 5:32.3 | פ | ' |
| 5:32.4 | ד | ' |
| 5:32.5 | מיתות |  |
| 5:32.6 | גרסינן |  |
| 5:32.7 | אמר |  |
| 5:32.8 | ר | ' |
| 5:32.9 | יהודה |  |
| 5:32.10 | אלמלא |  |
| 5:32.11 | וי"ו |  |
| 5:32.12 | שבהעלוך |  |
| 5:32.13 | נתחייבו |  |
| 5:32.14 | שונאיהם |  |
| 5:32.15 | של |  |
| 5:32.16 | ישראל |  |
| 5:32.17 | כלייה | . |
| 5:32.18 | אמר |  |
| 5:32.19 | לו |  |
| 5:32.20 | רשב"י |  |
| 5:32.21 | והלא |  |
| 5:32.22 | כל |  |
| 5:32.23 | המשתף |  |
| 5:32.24 | שם |  |
| 5:32.25 | שמים |  |
| 5:32.26 | ודבר |  |
| 5:32.27 | אחר |  |
| 5:32.28 | נעקר |  |
| 5:32.29 | מן |  |
| 5:32.30 | העולם |  |
| 5:32.31 | שנ | ' |
| 5:32.32 | בלתי |  |
| 5:32.33 | לה | ' |
| 5:32.34 | לבדו |  |
| 5:32.35 | אלא |  |
| 5:32.36 | מה |  |
| 5:32.37 | ת"ל |  |
| 5:32.38 | אשר |  |
| 5:32.39 | העלוך |  |
| 5:32.40 | שאיוו |  |
| 5:32.41 | לאלוהות |  |
| 5:32.42 | הרבה |  |
| 5:32.43 | ע"כ | , |
| 5:32.44 | ומביאו |  |
| 5:32.45 | מעשה |  |
| 5:32.46 | רקח |  |
| 5:32.47 | בפרק |  |
| 5:32.48 | ב | ' |
| 5:32.49 | מהל | ' |
| 5:32.50 | ע"ז | . |
| 5:32.51 | ובפ | ' |
| 5:32.52 | לולב |  |
| 5:32.53 | וערבה |  |
| 5:32.54 | אמתניתין |  |
| 5:32.55 | דהיו |  |
| 5:32.56 | אומרין |  |
| 5:32.57 | ליה |  |
| 5:32.58 | ולך |  |
| 5:32.59 | מזבח |  |
| 5:32.60 | פריך |  |
| 5:32.61 | תלמודא |  |
| 5:32.62 | והתניא |  |
| 5:32.63 | כל |  |
| 5:32.64 | המשתף |  |
| 5:32.65 | שם |  |
| 5:32.66 | שמים |  |
| 5:32.67 | ודבר |  |
| 5:32.68 | אחד |  |
| 5:32.69 | נעקר |  |
| 5:32.70 | מן |  |
| 5:32.71 | העולם | ? |
| 5:32.72 | אלא |  |
| 5:32.73 | הכי |  |
| 5:32.74 | קאמרי |  |
| 5:32.75 | וכו | ' |
| 5:32.76 | וכ | ' |
| 5:32.77 | הרב |  |
| 5:32.78 | מהרש"א |  |
| 5:32.79 | בפ | ' |
| 5:32.80 | ד | ' |
| 5:32.81 | מיתות |  |
| 5:32.82 | דודאי |  |
| 5:32.83 | לכ"ע |  |
| 5:32.84 | גרע |  |
| 5:32.85 | טפי |  |
| 5:32.86 | כופר |  |
| 5:32.87 | בעיקר |  |
| 5:32.88 | וחייב |  |
| 5:32.89 | כלייה | , |
| 5:32.90 | אלא |  |
| 5:32.91 | דרשב"י |  |
| 5:32.92 | קאמר |  |
| 5:32.93 | שכל |  |
| 5:32.94 | המשתף |  |
| 5:32.95 | שם |  |
| 5:32.96 | שמים |  |
| 5:32.97 | ודבר |  |
| 5:32.98 | אחר |  |
| 5:32.99 | דהיינו |  |
| 5:32.100 | המאמין |  |
| 5:32.101 | נמי |  |
| 5:32.102 | בדבר |  |
| 5:32.103 | אחר |  |
| 5:32.104 | חייב |  |
| 5:32.105 | כלייה |  |
| 5:32.106 | ונעקר |  |
| 5:32.107 | מן |  |
| 5:32.108 | העולם | , |
| 5:32.109 | אלא |  |
| 5:32.110 | מה |  |
| 5:32.111 | ת"ל |  |
| 5:32.112 | אשר |  |
| 5:32.113 | העלוך |  |
| 5:32.114 | דודאי |  |
| 5:32.115 | האמינו |  |
| 5:32.116 | בנמצא |  |
| 5:32.117 | וכו | ' |
| 5:32.118 | ולכך |  |
| 5:32.119 | לא |  |
| 5:32.120 | נעקרו |  |
| 5:32.121 | אלא |  |
| 5:32.122 | שאיוו |  |
| 5:32.123 | לאלוהות |  |
| 5:32.124 | הרבה |  |
| 5:32.125 | ואמרו |  |
| 5:32.126 | שהאחד |  |
| 5:32.127 | הנמצא |  |
| 5:32.128 | נתן |  |
| 5:32.129 | כח |  |
| 5:32.130 | לאלוהות |  |
| 5:32.131 | וכוחות |  |
| 5:32.132 | רבים |  |
| 5:32.133 | עזב |  |
| 5:32.134 | ה | ' |
| 5:32.135 | את |  |
| 5:32.136 | הארץ | . |
| 5:32.137 | ואלו |  |
| 5:32.138 | הכוחות |  |
| 5:32.139 | והאלוהות |  |
| 5:32.140 | יהיו |  |
| 5:32.141 | מנהיגים |  |
| 5:32.142 | בארץ | . |
| 5:32.143 | וכן |  |
| 5:32.144 | כ | ' |
| 5:32.145 | הרמב"ם |  |
| 5:32.146 | שבה |  |
| 5:32.147 | דעת |  |
| 5:32.148 | רבים |  |
| 5:32.149 | מעובדי |  |
| 5:32.150 | ע"ז |  |
| 5:32.151 | שעובדים |  |
| 5:32.152 | אותה | . |
| 5:32.153 | שאומרים |  |
| 5:32.154 | שניתנה |  |
| 5:32.155 | ההנהגה |  |
| 5:32.156 | בידה | . |
| 5:32.157 | ועל |  |
| 5:32.158 | דרך |  |
| 5:32.159 | זה |  |
| 5:32.160 | יש |  |
| 5:32.161 | לכוון |  |
| 5:32.162 | מה |  |
| 5:32.163 | שאמר |  |
| 5:32.164 | משה |  |
| 5:32.165 | ראה |  |
| 5:32.166 | אתה |  |
| 5:32.167 | אומר |  |
| 5:32.168 | אלי |  |
| 5:32.169 | העל |  |
| 5:32.170 | את |  |
| 5:32.171 | העם |  |
| 5:32.172 | הזה | . |
| 5:32.173 | שבא |  |
| 5:32.174 | משה |  |
| 5:32.175 | להתנצל |  |
| 5:32.176 | בעדם | . |
| 5:32.177 | שאמר |  |
| 5:32.178 | משה |  |
| 5:32.179 | להקב"ה |  |
| 5:32.180 | שהרי |  |
| 5:32.181 | אתה |  |
| 5:32.182 | אומר |  |
| 5:32.183 | אלי |  |
| 5:32.184 | העל |  |
| 5:32.185 | את |  |
| 5:32.186 | העם |  |
| 5:32.187 | הזה | . |
| 5:32.188 | וכוונתך |  |
| 5:32.189 | שאתה |  |
| 5:32.190 | מעמיד |  |
| 5:32.191 | אותי |  |
| 5:32.192 | למנהיג |  |
| 5:32.193 | להם | , |
| 5:32.194 | ותלית |  |
| 5:32.195 | העלייה |  |
| 5:32.196 | בי | , |
| 5:32.197 | כיוצא |  |
| 5:32.198 | בזה |  |
| 5:32.199 | כיוונו |  |
| 5:32.200 | כלל |  |
| 5:32.201 | לע"ז |  |
| 5:32.202 | ולשתף |  |
| 5:32.203 | שם |  |
| 5:32.204 | שמים |  |
| 5:32.205 | ודבר |  |
| 5:32.206 | אחר |  |
| 5:32.207 | עכ"ל |  |
| 5:32.208 | וכל |  |
| 5:32.209 | זה |  |
| 5:32.210 | דוחק |  |
| 5:32.211 | גדול | . |
| 5:32.212 | וכ | ' |
| 5:32.213 | מעשה |  |
| 5:32.214 | רקח |  |
| 5:32.215 | דאפשר |  |
| 5:32.216 | דשיתוף |  |
| 5:32.217 | דקאמר |  |
| 5:32.218 | היינו |  |
| 5:32.219 | לכל |  |
| 5:32.220 | האלוהות |  |
| 5:32.221 | עם |  |
| 5:32.222 | הקב"ה |  |
| 5:32.223 | ע"ש |  |
| 5:32.224 | ופשוט |  |
| 5:32.225 | הוא | . |
| 5:32.226 | מכל |  |
| 5:32.227 | זה |  |
| 5:32.228 | נתבאר |  |
| 5:32.229 | לך |  |
| 5:32.230 | ידידי |  |
| 5:32.231 | האוהב |  |
| 5:32.232 | את |  |
| 5:32.233 | תורת |  |
| 5:32.234 | האמת |  |
| 5:32.235 | ודברי |  |
| 5:32.236 | דודים |  |
| 5:32.237 | מעתיקיה |  |
| 5:32.238 | ומקבליה |  |
| 5:32.239 | ממשה |  |
| 5:32.240 | רבי | ' |
| 5:32.241 | איש |  |
| 5:32.242 | מפי |  |
| 5:32.243 | איש | , |
| 5:32.244 | שזה |  |
| 5:32.245 | שאמרו |  |
| 5:32.246 | המקובלים |  |
| 5:32.247 | החדשים |  |
| 5:32.248 | כי |  |
| 5:32.249 | אלהינו |  |
| 5:32.250 | בהו"ש |  |
| 5:32.251 | בהיותו |  |
| 5:32.252 | בבחינתו |  |
| 5:32.253 | הפשוטה |  |
| 5:32.254 | לא |  |
| 5:32.255 | שייך |  |
| 5:32.256 | בו |  |
| 5:32.257 | שום |  |
| 5:32.258 | עבודה |  |
| 5:32.259 | ותפלה |  |
| 5:32.260 | וקריאה |  |
| 5:32.261 | כלל | , |
| 5:32.262 | כי |  |
| 5:32.263 | קצור |  |
| 5:32.264 | קצרה |  |
| 5:32.265 | ידו |  |
| 5:32.266 | ח"ו |  |
| 5:32.267 | מפדות |  |
| 5:32.268 | ואין |  |
| 5:32.269 | בו |  |
| 5:32.270 | כח |  |
| 5:32.271 | להציל |  |
| 5:32.272 | את |  |
| 5:32.273 | הקורא |  |
| 5:32.274 | ומשוע |  |
| 5:32.275 | אליו |  |
| 5:32.276 | ורק |  |
| 5:32.277 | בהיותו |  |
| 5:32.278 | משתתף |  |
| 5:32.279 | בפרצופים |  |
| 5:32.280 | הנז | ' |
| 5:32.281 | וביחוד |  |
| 5:32.282 | פרצוף |  |
| 5:32.283 | זעיר |  |
| 5:32.284 | אנפין |  |
| 5:32.285 | אז |  |
| 5:32.286 | שייך |  |
| 5:32.287 | בו |  |
| 5:32.288 | עבודה |  |
| 5:32.289 | וברכה | . |
| 5:32.290 | וצריך |  |
| 5:32.291 | דוקא |  |
| 5:32.292 | לאחדו |  |
| 5:32.293 | ולשתפו |  |
| 5:32.294 | עם |  |
| 5:32.295 | נוקבא |  |
| 5:32.296 | דילה |  |
| 5:32.297 | ועם |  |
| 5:32.298 | שאר |  |
| 5:32.299 | החמשה |  |
| 5:32.300 | פרצופים | , |
| 5:32.301 | ואז |  |
| 5:32.302 | בהיותו |  |
| 5:32.303 | משתתף |  |
| 5:32.304 | עמהם |  |
| 5:32.305 | אז |  |
| 5:32.306 | המה |  |
| 5:32.307 | עונים |  |
| 5:32.308 | את |  |
| 5:32.309 | כל |  |
| 5:32.310 | הקורא |  |
| 5:32.311 | אליהם | , |
| 5:32.312 | וזהו |  |
| 5:32.313 | ממש |  |
| 5:32.314 | שיתוף |  |
| 5:32.315 | שם |  |
| 5:32.316 | שמים |  |
| 5:32.317 | ודבר |  |
| 5:32.318 | אחר | , |
| 5:32.319 | כי |  |
| 5:32.320 | לבדו |  |
| 5:32.321 | אינו |  |
| 5:32.322 | יכול |  |
| 5:32.323 | לעשות |  |
| 5:32.324 | מאומה | , |
| 5:32.325 | ורק |  |
| 5:32.326 | בשיתופו |  |
| 5:32.327 | עמהם |  |
| 5:32.328 | הוא |  |
| 5:32.329 | יכול |  |
| 5:32.330 | לפעול |  |
| 5:32.331 | את |  |
| 5:32.332 | פעולותיו | , |
| 5:32.333 | ואין |  |
| 5:32.334 | לך |  |
| 5:32.335 | שיתוף |  |
| 5:32.336 | שם |  |
| 5:32.337 | שמים |  |
| 5:32.338 | ודבר |  |
| 5:32.339 | אחר |  |
| 5:32.340 | יותר |  |
| 5:32.341 | מזה | . |
| 5:32.342 | ומאי |  |
| 5:32.343 | דעל |  |
| 5:32.344 | רשב"י |  |
| 5:32.345 | סני |  |
| 5:32.346 | בקבלה |  |
| 5:32.347 | האמיתית |  |
| 5:32.348 | בתלמוד | , |
| 5:32.349 | עשה |  |
| 5:32.350 | רשב"י |  |
| 5:32.351 | המשקר |  |
| 5:32.352 | מחבר |  |
| 5:32.353 | ס | ' |
| 5:32.354 | הזהר |  |
| 5:32.355 | בקבלה |  |
| 5:32.356 | המזויפת | , |
| 5:32.357 | רחמנא |  |
| 5:32.358 | ליצלן |  |
| 5:32.359 | מאמונות |  |
| 5:32.360 | זרות |  |
| 5:32.361 | כאלה | , |
| 5:32.362 | והאמונה |  |
| 5:32.363 | הברורה |  |
| 5:32.364 | הטהורה |  |
| 5:32.365 | מכל |  |
| 5:32.366 | סיג |  |
| 5:32.367 | וחלאה |  |
| 5:32.368 | היא |  |
| 5:32.369 | להאמין |  |
| 5:32.370 | באמונה |  |
| 5:32.371 | שלימה |  |
| 5:32.372 | תמימה |  |
| 5:32.373 | כי |  |
| 5:32.374 | לו |  |
| 5:32.375 | לבדו |  |
| 5:32.376 | ית | ' |
| 5:32.377 | הכח |  |
| 5:32.378 | והממשלה |  |
| 5:32.379 | בכל |  |
| 5:32.380 | הנבראים |  |
| 5:32.381 | שהמציא |  |
| 5:32.382 | בחסדו |  |
| 5:32.383 | ולו |  |
| 5:32.384 | ית | ' |
| 5:32.385 | נכוין |  |
| 5:32.386 | בכל |  |
| 5:32.387 | עבודותינו |  |
| 5:32.388 | ותפלותינו |  |
| 5:32.389 | כמצווה |  |
| 5:32.390 | עלינו |  |
| 5:32.391 | בתורה |  |
| 5:32.392 | בלתי |  |
| 5:32.393 | לה | ' |
| 5:32.394 | לבדו |  |
| 5:32.395 | שהוא |  |
| 5:32.396 | הסבה |  |
| 5:32.397 | הראשונה |  |
| 5:32.398 | בלי |  |
| 5:32.399 | צרוף |  |
| 5:32.400 | ושיתוף |  |
| 5:32.401 | שום |  |
| 5:32.402 | כח |  |
| 5:32.403 | אחר |  |
| 5:32.404 | לא |  |
| 5:32.405 | עתיק |  |
| 5:32.406 | ונוקביה | , |
| 5:32.407 | ולא |  |
| 5:32.408 | אריך |  |
| 5:32.409 | ונוקביה | , |
| 5:32.410 | ולא |  |
| 5:32.411 | אבא |  |
| 5:32.412 | ואימא | , |
| 5:32.413 | ולא |  |
| 5:32.414 | זעיר |  |
| 5:32.415 | ונוקביה | , |
| 5:32.416 | כי |  |
| 5:32.417 | כל |  |
| 5:32.418 | הנמצאים |  |
| 5:32.419 | צריכים |  |
| 5:32.420 | לו |  |
| 5:32.421 | והוא |  |
| 5:32.422 | אינו |  |
| 5:32.423 | צריך |  |
| 5:32.424 | להם | , |
| 5:32.425 | וכן |  |
| 5:32.426 | כתב |  |
| 5:32.427 | הראב"ע |  |
| 5:32.428 | בפי | ' |
| 5:32.429 | החומש |  |
| 5:32.430 | בלתי |  |
| 5:32.431 | לה | ' |
| 5:32.432 | לבדו |  |
| 5:32.433 | ולא |  |
| 5:32.434 | יתערב |  |
| 5:32.435 | אחר |  |
| 5:32.436 | עמו |  |
| 5:32.437 | ע"כ | , |
| 5:32.438 | ושמעתי |  |
| 5:32.439 | קצת |  |
| 5:32.440 | טפשים |  |
| 5:32.441 | אמרו |  |
| 5:32.442 | דלא |  |
| 5:32.443 | מיקרי |  |
| 5:32.444 | משתף |  |
| 5:32.445 | שם |  |
| 5:32.446 | שמים |  |
| 5:32.447 | ודבר |  |
| 5:32.448 | אחר |  |
| 5:32.449 | אלא |  |
| 5:32.450 | דוקא |  |
| 5:32.451 | כשמוציא |  |
| 5:32.452 | בפיו |  |
| 5:32.453 | דומיא |  |
| 5:32.454 | דמתניתין |  |
| 5:32.455 | דקאמרי |  |
| 5:32.456 | ליה |  |
| 5:32.457 | ולך |  |
| 5:32.458 | מזבח | , |
| 5:32.459 | שעליה |  |
| 5:32.460 | הקשו |  |
| 5:32.461 | בגמרא |  |
| 5:32.462 | והא |  |
| 5:32.463 | משתף |  |
| 5:32.464 | שם |  |
| 5:32.465 | שמים |  |
| 5:32.466 | ודבר |  |
| 5:32.467 | אחר |  |
| 5:32.468 | וכו | '. |
| 5:32.469 | ואין |  |
| 5:32.470 | זה |  |
| 5:32.471 | אמת |  |
| 5:32.472 | דאפי | ' |
| 5:32.473 | לכוון |  |
| 5:32.474 | במחשבתו |  |
| 5:32.475 | בהזכירו |  |
| 5:32.476 | את |  |
| 5:32.477 | השם |  |
| 5:32.478 | בברכה |  |
| 5:32.479 | או |  |
| 5:32.480 | בתפלה |  |
| 5:32.481 | לשום |  |
| 5:32.482 | נברא |  |
| 5:32.483 | או |  |
| 5:32.484 | נאצל |  |
| 5:32.485 | אף |  |
| 5:32.486 | אם |  |
| 5:32.487 | מכיון |  |
| 5:32.488 | לשתף |  |
| 5:32.489 | כח |  |
| 5:32.490 | הניתן |  |
| 5:32.491 | בו |  |
| 5:32.492 | מאת |  |
| 5:32.493 | השי"ת |  |
| 5:32.494 | אסור | , |
| 5:32.495 | והרי |  |
| 5:32.496 | זה |  |
| 5:32.497 | עובד |  |
| 5:32.498 | ע"ז | , |
| 5:32.499 | דמחשבת |  |
| 5:32.500 | ע"ז |  |
| 5:32.501 | הקב"ה |  |
| 5:32.502 | מצרפה |  |
| 5:32.503 | למעשה |  |
| 5:32.504 | כדגרסינן |  |
| 5:32.505 | בקידושין |  |
| 5:32.506 | ד | ' |
| 5:32.507 | ל"ט |  |
| 5:32.508 | ובחולין |  |
| 5:32.509 | קמ"ב |  |
| 5:32.510 | ואלא |  |
| 5:32.511 | הא |  |
| 5:32.512 | דכתיב |  |
| 5:32.513 | למען |  |
| 5:32.514 | תפוש |  |
| 5:32.515 | את |  |
| 5:32.516 | בית |  |
| 5:32.517 | ישראל |  |
| 5:32.518 | בלבם |  |
| 5:32.519 | ההוא |  |
| 5:32.520 | כמחשבת |  |
| 5:32.521 | ע"ז |  |
| 5:32.522 | הוא |  |
| 5:32.523 | דכתיב |  |
| 5:32.524 | דאפי | ' |
| 5:32.525 | חרהור |  |
| 5:32.526 | הוי |  |
| 5:32.527 | כמעשה | . |
| 5:32.528 | ובנ"ד |  |
| 5:32.529 | כיון |  |
| 5:32.530 | שמכוין |  |
| 5:32.531 | במחשבתו |  |
| 5:32.532 | לזעיר |  |
| 5:32.533 | אנפין |  |
| 5:32.534 | הרי |  |
| 5:32.535 | זה |  |
| 5:32.536 | כעובד |  |
| 5:32.537 | ע"ז | . |
| 5:32.538 | וקרינא |  |
| 5:32.539 | ביה |  |
| 5:32.540 | מה |  |
| 5:32.541 | שכ | ' |
| 5:32.542 | הרמב"ם |  |
| 5:32.543 | בפרק |  |
| 5:32.544 | ג | ' |
| 5:32.545 | לראשון |  |
| 5:32.546 | קרוב |  |
| 5:32.547 | אתה |  |
| 5:32.548 | בפיהם |  |
| 5:32.549 | ורחוק |  |
| 5:32.550 | מכליותיהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

