# Translation pack: milhamot-hashem perek-7-20 (Milchamot ha-Shem 7:25–7:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-20`.

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
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
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
- **הגיע חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
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


## Milchamot ha-Shem 7:25

Hebrew (Wikisource): צז) וראיתי לאחד קדוש מחכמי ימי הבינים הוא מהר"י צאהרי אשר חי בתימן בעת התפשטות ספרי הזהר והקבלה הנכריה החדשה בארץ התימן כמו שכ' בפרשת לך לך שהן חדשים מקרוב באו. ובפרשת בראשית ביאר שיש שתי דיעות ביחוד השי"ת וז"ל ואחר הקדמתי זאת צריך אני לגלות לך מעט כי בזה הענין (ר"ל יחוד השם) יש שתי דעות, האחת דעת הרמב"ם והנמשכים אחריו שלא הגיע אליהם דבר מדברי הקבלה (הנכריה החדשה). וסברתם ביחוד השם על דרך שלילה רצוני לומר לסלק הספק מלבו ולסלק מהשם ית' כל החסרונות ואז הוא אחד פשוט ועוד אמרו (גמרא ברכות דף י"ג ע"ב) צריך להאריך בדלת של אחד כדי שימליכהו על ארבע פנות העולם כולל הדברים לא תמצא מגמתם מגעת בכל מה שהביאו בדיבריהם אלא עד כאן. והדעת האחרת הם בעלי הקבלה העושים את כל מלאכת הקדש אמרו כי סוד היחוד הוא ליחד את השם בכל עשר ספירות. והקדושה גם כן צריך שיכוין בכל עשר ספירות שהם מערכת האלהות, וכל המצות מתפשטות מן העשר רצוני לומר שכל מצוה ומצוה יש לה סוד באחת מן העשר כמו שאמר בוצנא קדישא המאיר לארץ ולדרים רשב"י בפ' וישמע יתרו, כל פקודי אוריתא מתאחדן בגופא דמלכא מנהון ברישא ומנהון בגופא, ומנהון בידי מלכא ומנהון ברגלוי. ולית מאן דנפיק מגופא דמלכא לבר וכו'. ואתה המעיין אל תתמה במה שאמרו רישא ידא ורגלא שהרי כל אלו רומזים אל הספירות כמו שאצייר לך עתה שיח השדה הקדושה והטהורה העליונה כפי מה שסידר הרב שערי אורה אחת לאחת למצוא חשבון. אבל כי חג' ראשונות רומזים לראש. גדולה גבורה זרועות עולם. תפארת לשדרה וללב. נצח והוד תרין ירכין. ויסוד דא ברית. ומלכות היא מקבלת את הכל ומשפעת לכל התחתונים. וזהו סוד היחוד והקדושה, וצריך לחבר את האהל להיות אחד השני שמות ואהדונה"א. נמצאת למד כשתאמר מלת אחד, תכוין לתשע ספירות, אל"ף וחי"ת, ומלכות רומזת לעשירית, שהיא דל"ת, וחסוד למען יאריך ימים על ממלכתו, על שם הבורא מלאך וגלגל ואופן עכ"ל. ראה אבי גם ראה איך החכם הנז' מהר"י צאהרי הרגיש וידע שהם שתי דעות מתנגדות ביחוד השי"ת, אך ורק כי חשב את הראשונים חכמי התלמוד שאמרו להאריך בדל"ת דאחד כדי להמליכו בארבע פינות העולם לחסידי מדע. שלא ידעו בקבלה הנכריה החדשה ולא הגיע אליהם דבר ממנה, ולא ידעו ליחד את השם עם ברואיו הספירות המדומות שעשו אותם גופים ופרצופים לאלהינו ב"ה והוא נשמה להם. ואדרבא נהפוך הוא שאיסור גמור הוא לכלול ולשתף עם השי"ת שום נברא כמו שנתבאר לעיל סי' ע"ד. וסי' פ"ב בשם מדרש רבה, והטעם כי הוא ית' נעלה ומרומם ושגיא כה, ודי לו בעצמותו ואין צריך לבריותיו להשתתף עם שום אחד מהם. שכל בריותיו צריכים לו, והוא אינו צריך להם להעזר בהם ח"ו כי הוא אמר ויהי הוא לבדו צוה ויעמוד. ולא ירד הרב הנז' לעומק דברי ה רס"ג והרמב"ם שהוא עסוק בהם להעתיק דבריהם והוא מחבבן, ועשה אותן כעין הצעה והקדמה לאמתת יחוד השי"ת, והרגיש שהם שתי סברות. ולא הרגיש הסתירה בפרק ג' לראשון שהעתיק מקצת דבריו, שהם מתנגדים לגמרי לשטת הקבלה הנכריה החדשה, רק חשב אותם לחסידי מדע. שלא ידעו לצרף ולשתף ולכלול את הצורות והפרצופים עם אמתתו ית' וליחדו אמהם דוקא. ולא ידעו ולא יבינו כי הרס"ג והרמב"ם ורבי יהודה הלוי ובעל חובת הלבבות נלחמים המה בחזקת היד עם דעות הנכריות הזרות שהתחילו להנץ בין היהודים מזמן הגאונים, ועדיין היו בעל פה ולא הוחקו על ספר. עד שיצא הספר הנכרי הזהר המבזה את תורה שבעל פה המשנה והתלמוד, וקורא להם אבן משכית, ועץ הדעת טוב ורע. וקליפה, ושפחה כי תירש גברתה, נגד דברי דוד המלך ע"ה שאמר עליהם הנחמדים מזהב ומפז רב ומתוקים מדבש ומנופת צופים, והסית והדיח את הרבים לעבוד לאל אחר קצר אפים נברא, כמו שכתב הר' יושר לבב לעבדו אע"פ שהוא נברא, לפי שהאמינו באמונה שלימה כי מרשב"י יצאו הדברים ההם. והסבה שגרמה להם זה שלא שמו אל לבם ונגד עיניהם דברי הראשונים בעלי הכללים המעמיקים בתלמוד, המה דברי הרא"ש ז"ל שהבאתי לעיל סי' מ"ב שכל תוספתא או בריתא שלא נתפשטה עד אחר סידור הגמרא לאו דסמכא היא. דמסתמא מחברי התלמוד חקרו ודרשו לידע כל הספרים שנכתבו על דעת חכמים וביררו אותם שהם בר סמכא. וכל תוספתא או ברייתא שהיא היפך הגמרא לאו דסמכא היא דודאי משבשתא היא או מזויפת ע"ש.

| id | word | punctuation after |
|---|---|---|
| 7:25.1 | צז | ) |
| 7:25.2 | וראיתי |  |
| 7:25.3 | לאחד |  |
| 7:25.4 | קדוש |  |
| 7:25.5 | מחכמי |  |
| 7:25.6 | ימי |  |
| 7:25.7 | הבינים |  |
| 7:25.8 | הוא |  |
| 7:25.9 | מהר"י |  |
| 7:25.10 | צאהרי |  |
| 7:25.11 | אשר |  |
| 7:25.12 | חי |  |
| 7:25.13 | בתימן |  |
| 7:25.14 | בעת |  |
| 7:25.15 | התפשטות |  |
| 7:25.16 | ספרי |  |
| 7:25.17 | הזהר |  |
| 7:25.18 | והקבלה |  |
| 7:25.19 | הנכריה |  |
| 7:25.20 | החדשה |  |
| 7:25.21 | בארץ |  |
| 7:25.22 | התימן |  |
| 7:25.23 | כמו |  |
| 7:25.24 | שכ | ' |
| 7:25.25 | בפרשת |  |
| 7:25.26 | לך |  |
| 7:25.27 | לך |  |
| 7:25.28 | שהן |  |
| 7:25.29 | חדשים |  |
| 7:25.30 | מקרוב |  |
| 7:25.31 | באו | . |
| 7:25.32 | ובפרשת |  |
| 7:25.33 | בראשית |  |
| 7:25.34 | ביאר |  |
| 7:25.35 | שיש |  |
| 7:25.36 | שתי |  |
| 7:25.37 | דיעות |  |
| 7:25.38 | ביחוד |  |
| 7:25.39 | השי"ת |  |
| 7:25.40 | וז"ל |  |
| 7:25.41 | ואחר |  |
| 7:25.42 | הקדמתי |  |
| 7:25.43 | זאת |  |
| 7:25.44 | צריך |  |
| 7:25.45 | אני |  |
| 7:25.46 | לגלות |  |
| 7:25.47 | לך |  |
| 7:25.48 | מעט |  |
| 7:25.49 | כי |  |
| 7:25.50 | בזה |  |
| 7:25.51 | הענין |  |
| 7:25.52 | ר"ל |  |
| 7:25.53 | יחוד |  |
| 7:25.54 | השם | ) |
| 7:25.55 | יש |  |
| 7:25.56 | שתי |  |
| 7:25.57 | דעות | , |
| 7:25.58 | האחת |  |
| 7:25.59 | דעת |  |
| 7:25.60 | הרמב"ם |  |
| 7:25.61 | והנמשכים |  |
| 7:25.62 | אחריו |  |
| 7:25.63 | שלא |  |
| 7:25.64 | הגיע |  |
| 7:25.65 | אליהם |  |
| 7:25.66 | דבר |  |
| 7:25.67 | מדברי |  |
| 7:25.68 | הקבלה |  |
| 7:25.69 | הנכריה |  |
| 7:25.70 | החדשה | ). |
| 7:25.71 | וסברתם |  |
| 7:25.72 | ביחוד |  |
| 7:25.73 | השם |  |
| 7:25.74 | על |  |
| 7:25.75 | דרך |  |
| 7:25.76 | שלילה |  |
| 7:25.77 | רצוני |  |
| 7:25.78 | לומר |  |
| 7:25.79 | לסלק |  |
| 7:25.80 | הספק |  |
| 7:25.81 | מלבו |  |
| 7:25.82 | ולסלק |  |
| 7:25.83 | מהשם |  |
| 7:25.84 | ית | ' |
| 7:25.85 | כל |  |
| 7:25.86 | החסרונות |  |
| 7:25.87 | ואז |  |
| 7:25.88 | הוא |  |
| 7:25.89 | אחד |  |
| 7:25.90 | פשוט |  |
| 7:25.91 | ועוד |  |
| 7:25.92 | אמרו |  |
| 7:25.93 | גמרא |  |
| 7:25.94 | ברכות |  |
| 7:25.95 | דף |  |
| 7:25.96 | י"ג |  |
| 7:25.97 | ע"ב | ) |
| 7:25.98 | צריך |  |
| 7:25.99 | להאריך |  |
| 7:25.100 | בדלת |  |
| 7:25.101 | של |  |
| 7:25.102 | אחד |  |
| 7:25.103 | כדי |  |
| 7:25.104 | שימליכהו |  |
| 7:25.105 | על |  |
| 7:25.106 | ארבע |  |
| 7:25.107 | פנות |  |
| 7:25.108 | העולם |  |
| 7:25.109 | כולל |  |
| 7:25.110 | הדברים |  |
| 7:25.111 | לא |  |
| 7:25.112 | תמצא |  |
| 7:25.113 | מגמתם |  |
| 7:25.114 | מגעת |  |
| 7:25.115 | בכל |  |
| 7:25.116 | מה |  |
| 7:25.117 | שהביאו |  |
| 7:25.118 | בדיבריהם |  |
| 7:25.119 | אלא |  |
| 7:25.120 | עד |  |
| 7:25.121 | כאן | . |
| 7:25.122 | והדעת |  |
| 7:25.123 | האחרת |  |
| 7:25.124 | הם |  |
| 7:25.125 | בעלי |  |
| 7:25.126 | הקבלה |  |
| 7:25.127 | העושים |  |
| 7:25.128 | את |  |
| 7:25.129 | כל |  |
| 7:25.130 | מלאכת |  |
| 7:25.131 | הקדש |  |
| 7:25.132 | אמרו |  |
| 7:25.133 | כי |  |
| 7:25.134 | סוד |  |
| 7:25.135 | היחוד |  |
| 7:25.136 | הוא |  |
| 7:25.137 | ליחד |  |
| 7:25.138 | את |  |
| 7:25.139 | השם |  |
| 7:25.140 | בכל |  |
| 7:25.141 | עשר |  |
| 7:25.142 | ספירות | . |
| 7:25.143 | והקדושה |  |
| 7:25.144 | גם |  |
| 7:25.145 | כן |  |
| 7:25.146 | צריך |  |
| 7:25.147 | שיכוין |  |
| 7:25.148 | בכל |  |
| 7:25.149 | עשר |  |
| 7:25.150 | ספירות |  |
| 7:25.151 | שהם |  |
| 7:25.152 | מערכת |  |
| 7:25.153 | האלהות | , |
| 7:25.154 | וכל |  |
| 7:25.155 | המצות |  |
| 7:25.156 | מתפשטות |  |
| 7:25.157 | מן |  |
| 7:25.158 | העשר |  |
| 7:25.159 | רצוני |  |
| 7:25.160 | לומר |  |
| 7:25.161 | שכל |  |
| 7:25.162 | מצוה |  |
| 7:25.163 | ומצוה |  |
| 7:25.164 | יש |  |
| 7:25.165 | לה |  |
| 7:25.166 | סוד |  |
| 7:25.167 | באחת |  |
| 7:25.168 | מן |  |
| 7:25.169 | העשר |  |
| 7:25.170 | כמו |  |
| 7:25.171 | שאמר |  |
| 7:25.172 | בוצנא |  |
| 7:25.173 | קדישא |  |
| 7:25.174 | המאיר |  |
| 7:25.175 | לארץ |  |
| 7:25.176 | ולדרים |  |
| 7:25.177 | רשב"י |  |
| 7:25.178 | בפ | ' |
| 7:25.179 | וישמע |  |
| 7:25.180 | יתרו | , |
| 7:25.181 | כל |  |
| 7:25.182 | פקודי |  |
| 7:25.183 | אוריתא |  |
| 7:25.184 | מתאחדן |  |
| 7:25.185 | בגופא |  |
| 7:25.186 | דמלכא |  |
| 7:25.187 | מנהון |  |
| 7:25.188 | ברישא |  |
| 7:25.189 | ומנהון |  |
| 7:25.190 | בגופא | , |
| 7:25.191 | ומנהון |  |
| 7:25.192 | בידי |  |
| 7:25.193 | מלכא |  |
| 7:25.194 | ומנהון |  |
| 7:25.195 | ברגלוי | . |
| 7:25.196 | ולית |  |
| 7:25.197 | מאן |  |
| 7:25.198 | דנפיק |  |
| 7:25.199 | מגופא |  |
| 7:25.200 | דמלכא |  |
| 7:25.201 | לבר |  |
| 7:25.202 | וכו | '. |
| 7:25.203 | ואתה |  |
| 7:25.204 | המעיין |  |
| 7:25.205 | אל |  |
| 7:25.206 | תתמה |  |
| 7:25.207 | במה |  |
| 7:25.208 | שאמרו |  |
| 7:25.209 | רישא |  |
| 7:25.210 | ידא |  |
| 7:25.211 | ורגלא |  |
| 7:25.212 | שהרי |  |
| 7:25.213 | כל |  |
| 7:25.214 | אלו |  |
| 7:25.215 | רומזים |  |
| 7:25.216 | אל |  |
| 7:25.217 | הספירות |  |
| 7:25.218 | כמו |  |
| 7:25.219 | שאצייר |  |
| 7:25.220 | לך |  |
| 7:25.221 | עתה |  |
| 7:25.222 | שיח |  |
| 7:25.223 | השדה |  |
| 7:25.224 | הקדושה |  |
| 7:25.225 | והטהורה |  |
| 7:25.226 | העליונה |  |
| 7:25.227 | כפי |  |
| 7:25.228 | מה |  |
| 7:25.229 | שסידר |  |
| 7:25.230 | הרב |  |
| 7:25.231 | שערי |  |
| 7:25.232 | אורה |  |
| 7:25.233 | אחת |  |
| 7:25.234 | לאחת |  |
| 7:25.235 | למצוא |  |
| 7:25.236 | חשבון | . |
| 7:25.237 | אבל |  |
| 7:25.238 | כי |  |
| 7:25.239 | חג | ' |
| 7:25.240 | ראשונות |  |
| 7:25.241 | רומזים |  |
| 7:25.242 | לראש | . |
| 7:25.243 | גדולה |  |
| 7:25.244 | גבורה |  |
| 7:25.245 | זרועות |  |
| 7:25.246 | עולם | . |
| 7:25.247 | תפארת |  |
| 7:25.248 | לשדרה |  |
| 7:25.249 | וללב | . |
| 7:25.250 | נצח |  |
| 7:25.251 | והוד |  |
| 7:25.252 | תרין |  |
| 7:25.253 | ירכין | . |
| 7:25.254 | ויסוד |  |
| 7:25.255 | דא |  |
| 7:25.256 | ברית | . |
| 7:25.257 | ומלכות |  |
| 7:25.258 | היא |  |
| 7:25.259 | מקבלת |  |
| 7:25.260 | את |  |
| 7:25.261 | הכל |  |
| 7:25.262 | ומשפעת |  |
| 7:25.263 | לכל |  |
| 7:25.264 | התחתונים | . |
| 7:25.265 | וזהו |  |
| 7:25.266 | סוד |  |
| 7:25.267 | היחוד |  |
| 7:25.268 | והקדושה | , |
| 7:25.269 | וצריך |  |
| 7:25.270 | לחבר |  |
| 7:25.271 | את |  |
| 7:25.272 | האהל |  |
| 7:25.273 | להיות |  |
| 7:25.274 | אחד |  |
| 7:25.275 | השני |  |
| 7:25.276 | שמות |  |
| 7:25.277 | ואהדונה"א | . |
| 7:25.278 | נמצאת |  |
| 7:25.279 | למד |  |
| 7:25.280 | כשתאמר |  |
| 7:25.281 | מלת |  |
| 7:25.282 | אחד | , |
| 7:25.283 | תכוין |  |
| 7:25.284 | לתשע |  |
| 7:25.285 | ספירות | , |
| 7:25.286 | אל"ף |  |
| 7:25.287 | וחי"ת | , |
| 7:25.288 | ומלכות |  |
| 7:25.289 | רומזת |  |
| 7:25.290 | לעשירית | , |
| 7:25.291 | שהיא |  |
| 7:25.292 | דל"ת | , |
| 7:25.293 | וחסוד |  |
| 7:25.294 | למען |  |
| 7:25.295 | יאריך |  |
| 7:25.296 | ימים |  |
| 7:25.297 | על |  |
| 7:25.298 | ממלכתו | , |
| 7:25.299 | על |  |
| 7:25.300 | שם |  |
| 7:25.301 | הבורא |  |
| 7:25.302 | מלאך |  |
| 7:25.303 | וגלגל |  |
| 7:25.304 | ואופן |  |
| 7:25.305 | עכ"ל | . |
| 7:25.306 | ראה |  |
| 7:25.307 | אבי |  |
| 7:25.308 | גם |  |
| 7:25.309 | ראה |  |
| 7:25.310 | איך |  |
| 7:25.311 | החכם |  |
| 7:25.312 | הנז | ' |
| 7:25.313 | מהר"י |  |
| 7:25.314 | צאהרי |  |
| 7:25.315 | הרגיש |  |
| 7:25.316 | וידע |  |
| 7:25.317 | שהם |  |
| 7:25.318 | שתי |  |
| 7:25.319 | דעות |  |
| 7:25.320 | מתנגדות |  |
| 7:25.321 | ביחוד |  |
| 7:25.322 | השי"ת | , |
| 7:25.323 | אך |  |
| 7:25.324 | ורק |  |
| 7:25.325 | כי |  |
| 7:25.326 | חשב |  |
| 7:25.327 | את |  |
| 7:25.328 | הראשונים |  |
| 7:25.329 | חכמי |  |
| 7:25.330 | התלמוד |  |
| 7:25.331 | שאמרו |  |
| 7:25.332 | להאריך |  |
| 7:25.333 | בדל"ת |  |
| 7:25.334 | דאחד |  |
| 7:25.335 | כדי |  |
| 7:25.336 | להמליכו |  |
| 7:25.337 | בארבע |  |
| 7:25.338 | פינות |  |
| 7:25.339 | העולם |  |
| 7:25.340 | לחסידי |  |
| 7:25.341 | מדע | . |
| 7:25.342 | שלא |  |
| 7:25.343 | ידעו |  |
| 7:25.344 | בקבלה |  |
| 7:25.345 | הנכריה |  |
| 7:25.346 | החדשה |  |
| 7:25.347 | ולא |  |
| 7:25.348 | הגיע |  |
| 7:25.349 | אליהם |  |
| 7:25.350 | דבר |  |
| 7:25.351 | ממנה | , |
| 7:25.352 | ולא |  |
| 7:25.353 | ידעו |  |
| 7:25.354 | ליחד |  |
| 7:25.355 | את |  |
| 7:25.356 | השם |  |
| 7:25.357 | עם |  |
| 7:25.358 | ברואיו |  |
| 7:25.359 | הספירות |  |
| 7:25.360 | המדומות |  |
| 7:25.361 | שעשו |  |
| 7:25.362 | אותם |  |
| 7:25.363 | גופים |  |
| 7:25.364 | ופרצופים |  |
| 7:25.365 | לאלהינו |  |
| 7:25.366 | ב"ה |  |
| 7:25.367 | והוא |  |
| 7:25.368 | נשמה |  |
| 7:25.369 | להם | . |
| 7:25.370 | ואדרבא |  |
| 7:25.371 | נהפוך |  |
| 7:25.372 | הוא |  |
| 7:25.373 | שאיסור |  |
| 7:25.374 | גמור |  |
| 7:25.375 | הוא |  |
| 7:25.376 | לכלול |  |
| 7:25.377 | ולשתף |  |
| 7:25.378 | עם |  |
| 7:25.379 | השי"ת |  |
| 7:25.380 | שום |  |
| 7:25.381 | נברא |  |
| 7:25.382 | כמו |  |
| 7:25.383 | שנתבאר |  |
| 7:25.384 | לעיל |  |
| 7:25.385 | סי | ' |
| 7:25.386 | ע"ד | . |
| 7:25.387 | וסי | ' |
| 7:25.388 | פ"ב |  |
| 7:25.389 | בשם |  |
| 7:25.390 | מדרש |  |
| 7:25.391 | רבה | , |
| 7:25.392 | והטעם |  |
| 7:25.393 | כי |  |
| 7:25.394 | הוא |  |
| 7:25.395 | ית | ' |
| 7:25.396 | נעלה |  |
| 7:25.397 | ומרומם |  |
| 7:25.398 | ושגיא |  |
| 7:25.399 | כה | , |
| 7:25.400 | ודי |  |
| 7:25.401 | לו |  |
| 7:25.402 | בעצמותו |  |
| 7:25.403 | ואין |  |
| 7:25.404 | צריך |  |
| 7:25.405 | לבריותיו |  |
| 7:25.406 | להשתתף |  |
| 7:25.407 | עם |  |
| 7:25.408 | שום |  |
| 7:25.409 | אחד |  |
| 7:25.410 | מהם | . |
| 7:25.411 | שכל |  |
| 7:25.412 | בריותיו |  |
| 7:25.413 | צריכים |  |
| 7:25.414 | לו | , |
| 7:25.415 | והוא |  |
| 7:25.416 | אינו |  |
| 7:25.417 | צריך |  |
| 7:25.418 | להם |  |
| 7:25.419 | להעזר |  |
| 7:25.420 | בהם |  |
| 7:25.421 | ח"ו |  |
| 7:25.422 | כי |  |
| 7:25.423 | הוא |  |
| 7:25.424 | אמר |  |
| 7:25.425 | ויהי |  |
| 7:25.426 | הוא |  |
| 7:25.427 | לבדו |  |
| 7:25.428 | צוה |  |
| 7:25.429 | ויעמוד | . |
| 7:25.430 | ולא |  |
| 7:25.431 | ירד |  |
| 7:25.432 | הרב |  |
| 7:25.433 | הנז | ' |
| 7:25.434 | לעומק |  |
| 7:25.435 | דברי |  |
| 7:25.436 | ה |  |
| 7:25.437 | רס"ג |  |
| 7:25.438 | והרמב"ם |  |
| 7:25.439 | שהוא |  |
| 7:25.440 | עסוק |  |
| 7:25.441 | בהם |  |
| 7:25.442 | להעתיק |  |
| 7:25.443 | דבריהם |  |
| 7:25.444 | והוא |  |
| 7:25.445 | מחבבן | , |
| 7:25.446 | ועשה |  |
| 7:25.447 | אותן |  |
| 7:25.448 | כעין |  |
| 7:25.449 | הצעה |  |
| 7:25.450 | והקדמה |  |
| 7:25.451 | לאמתת |  |
| 7:25.452 | יחוד |  |
| 7:25.453 | השי"ת | , |
| 7:25.454 | והרגיש |  |
| 7:25.455 | שהם |  |
| 7:25.456 | שתי |  |
| 7:25.457 | סברות | . |
| 7:25.458 | ולא |  |
| 7:25.459 | הרגיש |  |
| 7:25.460 | הסתירה |  |
| 7:25.461 | בפרק |  |
| 7:25.462 | ג | ' |
| 7:25.463 | לראשון |  |
| 7:25.464 | שהעתיק |  |
| 7:25.465 | מקצת |  |
| 7:25.466 | דבריו | , |
| 7:25.467 | שהם |  |
| 7:25.468 | מתנגדים |  |
| 7:25.469 | לגמרי |  |
| 7:25.470 | לשטת |  |
| 7:25.471 | הקבלה |  |
| 7:25.472 | הנכריה |  |
| 7:25.473 | החדשה | , |
| 7:25.474 | רק |  |
| 7:25.475 | חשב |  |
| 7:25.476 | אותם |  |
| 7:25.477 | לחסידי |  |
| 7:25.478 | מדע | . |
| 7:25.479 | שלא |  |
| 7:25.480 | ידעו |  |
| 7:25.481 | לצרף |  |
| 7:25.482 | ולשתף |  |
| 7:25.483 | ולכלול |  |
| 7:25.484 | את |  |
| 7:25.485 | הצורות |  |
| 7:25.486 | והפרצופים |  |
| 7:25.487 | עם |  |
| 7:25.488 | אמתתו |  |
| 7:25.489 | ית | ' |
| 7:25.490 | וליחדו |  |
| 7:25.491 | אמהם |  |
| 7:25.492 | דוקא | . |
| 7:25.493 | ולא |  |
| 7:25.494 | ידעו |  |
| 7:25.495 | ולא |  |
| 7:25.496 | יבינו |  |
| 7:25.497 | כי |  |
| 7:25.498 | הרס"ג |  |
| 7:25.499 | והרמב"ם |  |
| 7:25.500 | ורבי |  |
| 7:25.501 | יהודה |  |
| 7:25.502 | הלוי |  |
| 7:25.503 | ובעל |  |
| 7:25.504 | חובת |  |
| 7:25.505 | הלבבות |  |
| 7:25.506 | נלחמים |  |
| 7:25.507 | המה |  |
| 7:25.508 | בחזקת |  |
| 7:25.509 | היד |  |
| 7:25.510 | עם |  |
| 7:25.511 | דעות |  |
| 7:25.512 | הנכריות |  |
| 7:25.513 | הזרות |  |
| 7:25.514 | שהתחילו |  |
| 7:25.515 | להנץ |  |
| 7:25.516 | בין |  |
| 7:25.517 | היהודים |  |
| 7:25.518 | מזמן |  |
| 7:25.519 | הגאונים | , |
| 7:25.520 | ועדיין |  |
| 7:25.521 | היו |  |
| 7:25.522 | בעל |  |
| 7:25.523 | פה |  |
| 7:25.524 | ולא |  |
| 7:25.525 | הוחקו |  |
| 7:25.526 | על |  |
| 7:25.527 | ספר | . |
| 7:25.528 | עד |  |
| 7:25.529 | שיצא |  |
| 7:25.530 | הספר |  |
| 7:25.531 | הנכרי |  |
| 7:25.532 | הזהר |  |
| 7:25.533 | המבזה |  |
| 7:25.534 | את |  |
| 7:25.535 | תורה |  |
| 7:25.536 | שבעל |  |
| 7:25.537 | פה |  |
| 7:25.538 | המשנה |  |
| 7:25.539 | והתלמוד | , |
| 7:25.540 | וקורא |  |
| 7:25.541 | להם |  |
| 7:25.542 | אבן |  |
| 7:25.543 | משכית | , |
| 7:25.544 | ועץ |  |
| 7:25.545 | הדעת |  |
| 7:25.546 | טוב |  |
| 7:25.547 | ורע | . |
| 7:25.548 | וקליפה | , |
| 7:25.549 | ושפחה |  |
| 7:25.550 | כי |  |
| 7:25.551 | תירש |  |
| 7:25.552 | גברתה | , |
| 7:25.553 | נגד |  |
| 7:25.554 | דברי |  |
| 7:25.555 | דוד |  |
| 7:25.556 | המלך |  |
| 7:25.557 | ע"ה |  |
| 7:25.558 | שאמר |  |
| 7:25.559 | עליהם |  |
| 7:25.560 | הנחמדים |  |
| 7:25.561 | מזהב |  |
| 7:25.562 | ומפז |  |
| 7:25.563 | רב |  |
| 7:25.564 | ומתוקים |  |
| 7:25.565 | מדבש |  |
| 7:25.566 | ומנופת |  |
| 7:25.567 | צופים | , |
| 7:25.568 | והסית |  |
| 7:25.569 | והדיח |  |
| 7:25.570 | את |  |
| 7:25.571 | הרבים |  |
| 7:25.572 | לעבוד |  |
| 7:25.573 | לאל |  |
| 7:25.574 | אחר |  |
| 7:25.575 | קצר |  |
| 7:25.576 | אפים |  |
| 7:25.577 | נברא | , |
| 7:25.578 | כמו |  |
| 7:25.579 | שכתב |  |
| 7:25.580 | הר | ' |
| 7:25.581 | יושר |  |
| 7:25.582 | לבב |  |
| 7:25.583 | לעבדו |  |
| 7:25.584 | אע"פ |  |
| 7:25.585 | שהוא |  |
| 7:25.586 | נברא | , |
| 7:25.587 | לפי |  |
| 7:25.588 | שהאמינו |  |
| 7:25.589 | באמונה |  |
| 7:25.590 | שלימה |  |
| 7:25.591 | כי |  |
| 7:25.592 | מרשב"י |  |
| 7:25.593 | יצאו |  |
| 7:25.594 | הדברים |  |
| 7:25.595 | ההם | . |
| 7:25.596 | והסבה |  |
| 7:25.597 | שגרמה |  |
| 7:25.598 | להם |  |
| 7:25.599 | זה |  |
| 7:25.600 | שלא |  |
| 7:25.601 | שמו |  |
| 7:25.602 | אל |  |
| 7:25.603 | לבם |  |
| 7:25.604 | ונגד |  |
| 7:25.605 | עיניהם |  |
| 7:25.606 | דברי |  |
| 7:25.607 | הראשונים |  |
| 7:25.608 | בעלי |  |
| 7:25.609 | הכללים |  |
| 7:25.610 | המעמיקים |  |
| 7:25.611 | בתלמוד | , |
| 7:25.612 | המה |  |
| 7:25.613 | דברי |  |
| 7:25.614 | הרא"ש |  |
| 7:25.615 | ז"ל |  |
| 7:25.616 | שהבאתי |  |
| 7:25.617 | לעיל |  |
| 7:25.618 | סי | ' |
| 7:25.619 | מ"ב |  |
| 7:25.620 | שכל |  |
| 7:25.621 | תוספתא |  |
| 7:25.622 | או |  |
| 7:25.623 | בריתא |  |
| 7:25.624 | שלא |  |
| 7:25.625 | נתפשטה |  |
| 7:25.626 | עד |  |
| 7:25.627 | אחר |  |
| 7:25.628 | סידור |  |
| 7:25.629 | הגמרא |  |
| 7:25.630 | לאו |  |
| 7:25.631 | דסמכא |  |
| 7:25.632 | היא | . |
| 7:25.633 | דמסתמא |  |
| 7:25.634 | מחברי |  |
| 7:25.635 | התלמוד |  |
| 7:25.636 | חקרו |  |
| 7:25.637 | ודרשו |  |
| 7:25.638 | לידע |  |
| 7:25.639 | כל |  |
| 7:25.640 | הספרים |  |
| 7:25.641 | שנכתבו |  |
| 7:25.642 | על |  |
| 7:25.643 | דעת |  |
| 7:25.644 | חכמים |  |
| 7:25.645 | וביררו |  |
| 7:25.646 | אותם |  |
| 7:25.647 | שהם |  |
| 7:25.648 | בר |  |
| 7:25.649 | סמכא | . |
| 7:25.650 | וכל |  |
| 7:25.651 | תוספתא |  |
| 7:25.652 | או |  |
| 7:25.653 | ברייתא |  |
| 7:25.654 | שהיא |  |
| 7:25.655 | היפך |  |
| 7:25.656 | הגמרא |  |
| 7:25.657 | לאו |  |
| 7:25.658 | דסמכא |  |
| 7:25.659 | היא |  |
| 7:25.660 | דודאי |  |
| 7:25.661 | משבשתא |  |
| 7:25.662 | היא |  |
| 7:25.663 | או |  |
| 7:25.664 | מזויפת |  |
| 7:25.665 | ע"ש | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

