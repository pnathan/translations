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
- `grammar`: 2 to 6 notes per paragraph (in a long one, on the hardest
  sentences). `literal`: word for word; on a long paragraph the glosses in
  Hebrew order, hyphens joining the parts of one word; no reasoning (that goes
  in `note`).
- `phrases` divide every token of the paragraph, in order; in a long
  paragraph a phrase is a clause, not a word; each phrase's `english` must
  occur in the paragraph's `english`.
- Quotations follow the Quotations section; a Bible verse is translated from
  the Hebrew as the author quotes it, the reference in `apparatus`.
- Every Hebrew abbreviation is expanded in `gloss` (the reading line carries
  the full word).

## Names and titles

- **Authorities**: acronyms and book-names used as names stay, transliterated;
  first in a unit a note gives the expansion where certain (else says so) and
  who is meant; identify no one further. ר' + name → "Rabbi". An initial a
  surname fixes is expanded, else kept (a lone כ' → "K."), noted. Surnames as
  printed, one spelling across the work.
- **Honorifics**: translate every one, opponents' included; add none.
- **Book titles**: transliterated, the meaning in a note first in a unit;
  abbreviations written out; every word kept. Bible books in English;
  tractates transliterated. A misspelt title as printed, the likelier in the
  note and `apparatus`.
- ה׳: the Tetragrammaton → "the LORD"; cited as a name ("זעיר אנפין דאיקרי
  ה'") → "who is called 'the LORD'"; before a personal name → "Mr." ("מאת ה'
  אליעזר בן יהודה" → "from Mr. Eliezer Ben-Yehuda"); a numeral or page letter
  → the number. The pack splits it into ה plus ': read it in context.
- **מלחמות ה׳** → *The Wars of the LORD* (Num 21:14; Q1); transliterated
  *Milchamot ha-Shem*.
- **הצעיר יחיה / אלקאפח**: "הצעיר יחיה בן סלי' אלקאפח יצ״ו" → "The humble
  Yichya son of Sali' al-Kafich, may his Rock keep him and give him life"
  (patronymic as written, Q6).
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
- **מחבר הזהר / מחבר הספר הזהר**: "the author of the Zohar", never named (no
  "Moses de León"); מחבר הספר הזהר → "the author of the book the Zohar".
- **הפילסוף / הפילוסוף** (epithet of the Zohar's author, who is never named) →
  "the philosopher".
- נביא **השקר** (epithet of the Zohar's author, never named): נביא השקר → "the
  false prophet".
- **המסית / המדיח** (epithets of the Zohar's author, never named): המסית → "the
  enticer"; המדיח → "the seducer" (legal categories of Deut 13 and 18, on
  which he builds a case): never "deceiver" or "charlatan".
- **הוי"ה / הוי'ה / הויה / הויות** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG).
  עשר הויות → "ten havayot" (note).
- **אדנות** as a name → "Adonai" (note: "lordship").
- **הקב"ה / קב"ה / קודשא בריך** הוא → "the Holy One, blessed be He"; as a name
  discussed → "the name 'the Holy One, blessed be He'".
- **השי"ת** → "the Name, may He be blessed" (Q3).
- **ית׳ / ית' / יתברך** → "may He be blessed", where the Hebrew places it.
- **ב"ה** → "blessed be He".
- השם **הנכבד** → "the glorious Name".
- שם **שמים** → "the name of Heaven".
- **אלהים / אלהי / אלוה / קבלו באלוה**: no editorialising by case. "God" for the
  God of Israel, also in the Kabbalists' quoted words and for a plural אלהים
  with a singular verb; singular אלוה → "God" (Job, Deut 32:15), not "the
  first God". "god(s)" only where the Hebrew marks foreignness or plurality
  (אלהים אחרים, אל זר, אלילים), where the author in his own voice calls the
  Kabbalists' object of worship a god, and for אלוה as what a human being
  worships or makes himself (קבלו באלוה; Haman) → "a god". In doubt: "God",
  the alternative in a note.
- **האלהי** before a name → "the divine" (note); never dropped.
- **ז"ל** → "of blessed memory" (honorific).
- **ע"ה** → "peace be upon him" (honorific).
- **רז"ל** → "our Rabbis of blessed memory".
- **חז"ל** → "our Sages of blessed memory".
- **זלה"ה** → "may his memory be for the life of the world to come".
- **זצוק"ל** → "may the memory of the righteous and holy be for a blessing".
- **ת"ו** (Jerusalem) → "may it be rebuilt and established".
- **יר"ה** (the Imam) → "may his majesty be exalted".
- **הרמב"ם / רמב"ם** → "the Rambam" (authority, transliterated; note first in
  a unit).
- **המורה / רבינו**: הר' המורה / רבינו meaning the Rambam → "the Master, the
  Guide" / "our Master"; רבינו + name → "Rabbeinu" ("Rabbeinu Bachya").
- **בחיי** alone → "Bachya"; רבינו בחיי → "Rabbeinu Bachya"; בר' בחיי → "Rabbi
  Bachya".
- **הרמב"ן / רמב"ן** → "Ramban" (authority, transliterated; note first in a
  unit).
- **רש"י** → "Rashi" (authority, transliterated; note).
- **האר"י** → "Ari" (authority, transliterated; note).
- **ריב"ש** → "Ribash" (authority, transliterated; note).
- **רשב"א** → "Rashba" (authority, transliterated; note).
- **הרא"ש** → "Rosh" (authority, transliterated; note).
- **ראב"ד** → "Ravad" (authority, transliterated; note).
- **רמ"א** → "Rema" (authority, transliterated; note).
- **רדב"ז** → "Radbaz" (authority, transliterated; note).
- **מהרש"ל** → "Maharshal" (authority, transliterated; note).
- **מהר"ם** אלשקר → "Maharam al-Shakar" (authority, transliterated; note).
- **רא"ם** → "Ra'em" (authority, transliterated; note).
- **מרן** → "Maran" (authority, transliterated; note).
- **רוקח** → "Roke'ach" (authority, transliterated; note).
- **לחם** משנה → "Lechem Mishneh" (authority; הרב לחם משנה → "the Rav Lechem
  Mishneh").
- **חוות יאיר** → "Chavot Yair" (authority, transliterated; note).
- **פתחי תשובה** → "Pitchei Teshuvah" (authority, transliterated; note).
- בית **יוסף** → "Beit Yosef" (authority, transliterated; note).
- כנסת **הגדולה** → "Kenesset ha-Gedolah" (authority, transliterated; note).
- **מהר"י** + surname → "Mahari Lopes" (the surname as printed).
- **רס"ג** → "Rabbi Saadia Gaon", written out.
- **הרח"ו / רח"ו** → "Rabbi Chayyim Vital", written out.
- **ר"ת** → "Rabbeinu Tam".
- **רבי'** before the lawgiver's name → "Moses our teacher".
- **הגאון** alone → "the Gaon".
- בעל **חובת הלבבות** → "the author of Chovot ha-Levavot".
- בעל **העיקרים** → "the author of the Ikkarim".
- **סמ"ג** → "the Semag (Sefer Mitzvot Gadol)".
- **ראב"ע** → "Rabbi Eleazar ben Azariah" in a Talmudic passage, "Ibn Ezra" for
  the commentator (note).
- **ר"א** → "Rabbi Eleazar" only where the paragraph names him, else "R. A."
  (note).
- **ריקנטי**: ר"מ ריקנטי → "Rabbi Menachem Rikanti" (the surname fixes the
  initial).
- **ר"י** חיים → "Rabbi Y. Chayyim" (initial kept, with a note).
- יש"ר **מקנדי"א** → "Yish'ar of Candia" (note: conjectural).
- **אלקבולי** → "Alkaboli", Alkabetz in the note; one spelling across the work.
- **אפנדי / ביי / ביק / פחה / פאשא** (Turkish and Arabic names): the scheme as
  spelled; keep Effendi, Bey, Pasha; identify no one.
- **כסא אליהו / כס"א** → *Kisse Eliyahu* (book title; the meaning in a note
  first in a unit; כס"א written out).
- **יושר** לבב → *Yosher Levav* (book title; meaning in a note first in a unit).
- **עץ חיים / ע"ח** → *Etz Chayyim* (book title; meaning in a note).
- **מקדש** מלך → *Mikdash Melekh* (book title; הרב מקדש מלך → "the Rav Mikdash
  Melekh").
- **אמונות והדעות / אמונות ודעות** → *Emunot ve-De'ot* (book title; meaning in
  a note).
- **מורה נבוכים / המורה** → *Moreh Nevukhim*, also for "Guide for the
  Perplexed" (book title; meaning in a note first in a unit).
- **הכוזרי / כוזרי** → "the Kuzari" (book title; meaning in a note).
- ספר **הברית** → *Sefer ha-Berit* (book title; meaning in a note).
- **תיקוני / תקוני** זהר (הזהר) → *Tikkunei Zohar* (book title; meaning in a note).
- **זהר חדש** → *Zohar Chadash* (book title; meaning in a note).
- **אדרא / אידרא** of Naso → "Idra of Naso" (book title; meaning in a note).
- **ממלכי** → "of the kings of".
- **אלצלאם**: ונור אלצלאם (a misspelt title) as printed ("Nur al-Tzalam"), the
  likelier (Nur al-Zulam, Me'or ha-Afelah) in the note and `apparatus`.
- **הר' / הרב** + book title → "the Rav [title]" ("the Rav Lechem Mishneh",
  "the Rav Mikdash Melekh"; הר' כסא אליהו, הר' עץ יוסף), never dropping הרב,
  note on who is meant; a title is never a personal name; "the author of
  [title]" only for בעל / מחבר.
- **בעל / מחבר / היכל** + book title → "the author of [title]" (only for בעל /
  מחבר; "המקובל בעל היכל הברכה" → "the Kabbalist, the author of Heikhal
  ha-Berakhah", title translated in a note); name no author the text does not
  name.
- **עה"ת** → "on the Torah"; after an unidentified book's name → "a book on the
  Torah" (do not guess Etz Chayyim).
- **עשתונות** איש ירושלים → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **בבחירתא** → "in the Bechirta" (note: "chosen").
- **המשרף** (pun on Metzaref ha-Emunah) → "the Mesharef", a note on the pun.
- **תימן** → "Yemen".
- **צנעא** → "Sanaa".
- **תוגרמה**: ממשלת תוגרמה → "the government of Turkey" (note: Togarmah).
- **הישמעאלים / ישמעאלים** → "the Ishmaelites" (note: Muslims).
- **המושלמנים** → "the Muslims".
- **הנוצרים / הנצרים** → "the Christians".
- **השונים** → "the Dualists".
- **האמגושים** → "the Magians".
- מאמיני **השלוש** → "the believers in the Trinity".
- **מערב** beside Ashkenaz, Sefarad, Yemen → "the Maghreb" (note: "the West").
- **בא"י** → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
- **מדה / מדות / מידות / מדת**: by context, never "sefirot". Hermeneutical
  rules → "rules" (note: literally "measures"). Of God or the sefirot:
  "attribute" ("מדת הדין" → "the attribute of judgment"). Also "measure", and
  "quality / trait / virtue" of character. מסכת מדות → "tractate Middot".
- **ישמעאל / הגלילי**: "הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the
  thirteen rules of Rabbi Ishmael and the thirty-three rules of Rabbi Yose the
  Galilean" (the number as written; the usual thirty-two in a note; note:
  literally "measures").
- **רחמים**: "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes
  of mercy", a different thirteen from the hermeneutical rules.
- אין **סוף / א"ס**: fixed, "Ein Sof", a name with no article ("א"ס ב"ה" → "Ein
  Sof, blessed be He"); never "the Infinite" or "without end" for the noun;
  "without end" in a note first in a unit. לאין סוף where not the name (the
  Rambam passage) → "without end".
- **אצילות / באצילות / האצילות / המאציל**: fixed, translated, the verbs
  linked: "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; המאציל → "the Emanator".
- **נאצל / נברא**: נאצל → "emanated"; נאצל and נברא stay distinct ("נברא ולא
  בורא" → "created, not creator").
- **אבי"ע / בריאה / יצירה / עשיה**: the four worlds (אבי"ע): "Emanation,
  Creation, Formation, Making" (verbs: "emanated, created, formed, made").
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted"; "tzimtzum" in a
  note first in a unit.
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure". The names of the partzufim: transliterated, the
  literal sense in a note first in a unit; an argument from a name's meaning
  goes in the note; the name stays.
- **זעיר אנפין / ז"א / זעיר** → "Ze'ir Anpin" (partzuf; note first in a unit:
  "Small Face"); ז"א → "Ze'ir Anpin"; זעיר alone → "Ze'ir". An argument from
  the name's meaning goes in the note; the name stays.
- **אריך** אנפין → "Arikh Anpin" (partzuf; note first in a unit: "Long Face").
- **אבא / אמא** (the partzufim) → "Abba" / "Imma" (note first in a unit).
- **נוקבא / נוקביה** → "Nukva" (partzuf; note first in a unit); נוקביה → "his
  Nukva".
- **עתיק / עתיקא / יומין**: עתיק → "Atik"; עתיקא קדישא → "Atika Kadisha"; עתיקא
  סתימא → "Atika Setima"; עתיק יומין → "Atik Yomin" (Dan 7:9: "the Ancient of
  Days"); the literal sense in a note first in a unit.
- סבא **דסבין** → "the Elder of Elders".
- עלת **העלות** → "the Cause of causes".
- הסבה **הראשונה** → "the First Cause".
- כתר **עליון** → "Keter Elyon (the Supreme Crown)".
- קצר **אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- אדם **קדמון / קדמאה**: אדם קדמון → "Adam Kadmon" (note: "Primordial Man");
  אדם קדמאה → "Adam Kadma'ah": two entities, never merged.
- **קדמון / קדמות / נצחי / נצחיות**: of God or the world → "eternal / eternity"
  (note: "without beginning"); נצחי / נצחיות → "everlasting /
  everlastingness", distinct. קדמונינו → "our forebears" (fixed); הקדמונים →
  "the ancients"; קדמו → "preceded".
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
- **יחוד / ייחוד / ייחדו / אחדות**: the doctrine → "unity", fixed, not
  "monotheism" or "oneness" (אחדות → "oneness"). The Kabbalistic act →
  "unification" (יחודים → "unifications"). The verb, God as object → "to
  declare one" ("ליחדו" → "to declare Him one"); Kabbalistic entities as
  object → "to unify"; of dedicating a thing to one → "dedicate ... solely to"
  (note: "unify"). Where he plays on the senses: "unity" for the noun, "unify"
  for the verb, the play in a note.
- **כלא חד / כולא חד / באגרונו** (also ולא אמרינן כלא חד): one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **שיתוף / שתוף / לשתף / משותף / המשתף**: associating another with God →
  fixed, "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever
  associates the name of Heaven with something else"). The Maimonidean term
  for a word said of God and creatures → "equivocation / equivocal" ("בשתוף
  השם בלבד" → "by equivocation of the term only"), never "association", Guide
  I in the note; "shared" only for a thing held in common.
- **צירוף**: "conjunction" of God; "permutation" in letter mysticism.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז**: עבודה זרה → "idolatry" (note once: "strange worship");
  the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of idols".
- **עכו"ם** → "idolaters" (note: "worshippers of stars and constellations"),
  never "gentiles" or "Christians", even of a Christian king or the Trinity.
- עבודה **אחרת** → "other worship" (note first in a unit).
- **אליל / האלילה / זר**: אליל → "idol"; האלילה → "the she-idol"; אל זר → "a
  strange god".
- אל **אחר** → "another god".
- אלהים **אחרים** → "other gods".
- **אלהות / אלוהות**: "divinity"; counted or plural → "divinities". Never
  "gods".
- **גשמות / הגשמה / מגשים / להגשים / התגשמות**: "corporeality"; להגשים → "to
  ascribe corporeality to". Not technical: "לידי התגשמות" → "to realization".
  "גוף ולא כח בגוף" → "a body, or a force in a body".
- **קבלה / הקבלה / המקובלים**: the lore → "Kabbalah"; הקבלה החדשה → "the new
  Kabbalah" (never "later" or "Lurianic"). המקובלים → "the Kabbalists"; חכמי
  הקבלה → "sages of Kabbalah"; "רבותינו בעלי הקבלה החדשה" → "our rabbis, the
  masters of the new Kabbalah" (courtesy kept).
- **האמתית / האמיתית / ממשה** (his claim on the word קבלה for the rabbinic
  chain, Q5): "בעלי הקבלה האמתית" → "the masters of the true tradition", a
  note each time that the Hebrew is the word for Kabbalah; "אינה קבלה ממשה" →
  "it is not a tradition from Moses".
- **וכמקובל / בידינו** (the participle of קבל): "וכמקובל" → "as is accepted";
  "המקובל בידינו" → "what we have received".
- **השפע** → "the flow".
- **בחינה** → "aspect".
- **השתלשלות** → "the chain of descent".
- **עלה / עלול** → "cause / effect".
- **מחויב** המציאות → "the Necessary Existent".
- **כלים / כלי**: Kabbalistic כלים → "vessels"; a craftsman's כלי → "tool".
- אורות **דקים** → "subtle lights".
- גוף **דק** → "a subtle body".
- **כוונה** (in prayer) → "intention".
- **זווג** → "coupling".
- **השגחה** → "providence".
- לית מחשבה **תפיסה** ביה → "no thought grasps Him".
- **מין / מינות** → "heretic / heresy".
- **אפיקורוס / אפיקורסים** → "apikoros / apikorsim".
- כופר **בעיקר** → "denies the fundamental principle".
- ספר **חיצון / חיצונים** → "an external book" (note: m. Sanh. 10:1).
- תורה **שבעל** פה → "the Oral Torah".
- ימי **הבינים** → "the Middle Ages".
- אין לו **חלק** לעולם הבא → "has no share in the world to come" (lower case).
- בת **קול** → "bat kol" (note: "daughter of a voice").
- **Abbreviations**: ambiguous: by context, the expansion in the note.
  Unruled (ורבי'): by his usage, noted; of two expansions the shorter, the
  other in `apparatus`. Unclear (לי"ח): the conjecture in a note, the letters
  in `apparatus`, in the English only if the sole sense that fits.
- **ר"ל** → "that is to say" (with a named authority as subject: "means to
  say"), "Resh Lakish" ("אמר ר"ל"), or "the Merciful One save us" (at a
  sentence end); by context, the expansion in the note.
- **ע"כ**: by context, the expansion in the note: "thus far" (after a
  quotation; as the closing formula, outside the marks: "Thus far."),
  "therefore", "perforce" (על כרחך), "and upon all" ("וע"כ עמו ישראל").
- **ע"ש** → "see there" or "on account of" (על שם), by context; כ' ע"ש → "he
  writes, see there:" (על שם in `readings`).
- **ב"ו** → "blessed be He" before הוא, else "flesh and blood".
- **מ"ו** → "our teacher" in יודיענו מורנו.
- **ב"ע** → "in Aramaic" against בערבי (others in `readings`).
- **ח"ו** → "God forbid".
- **וכו' / וכ'** → "etc." (always kept; וכ' closing a quotation is וכו').
- **ג"כ** → "also".
- **כ"ש / ק"ו** → "all the more so".
- **ת"ל** → "Scripture says".
- **א"כ** → "if so".
- **סתו"מ** → "Torah scrolls, tefillin and mezuzot" (note).
- **תוה"ק** → "the holy Torah".
- **ביהכנ"ס** → "synagogue".
- **הינ"ץ** → "H.N.Ts." (note: perhaps ה + ינ"ץ, "may his Rock protect him", of
  the Imam); no person named.
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **ריבה / לרבות** → "include", never "multiply".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות** של בית רבן → "schoolchildren" (note: "the house of the master").
- ובפתרון **הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **מרגל** לרגל → "from festival to festival", not "foot".
- **גופי** תורה → "the essentials of the Torah" (note: "bodies").
- **מאין / מיש**: יש מאין / יש מיש → "something from nothing / from
  something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **אנדרטא** → "statue".
- אשם **תלוי** → "suspended guilt offering".
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- דברי **תימה** → "astonishing words".
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- הגשת **מנחה** (argued from Gen 43 / 1 Sam 10) → "gift"; "meal offering" only
  in a sacrificial context.
- **והוא שלא** → "provided that".
- **לשבר** את האזן → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **גיזומיו** → "exaggerations" of claims, else "trimmings"; the other in a note.
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **המתפלספים** → "would-be philosophers".
- **תשבורת** → "mathematics".
- **ותוכנים** → "astronomy".
- **קורין** אל → "call upon".
- **גבורה** as a name of God → "the Power".
- אחד ... **ואחד** of two kinds of case → "alike".
- ואין **צריך** לומר → "needless to say".
- **אורה** זו תורה (Esth 8:16) → "'orah' means Torah".
- **צחצחות** → "tzachtzachot".
- **חמותי** (Isa 44:16) → "I am warm".
- **נדנוד** → "qualm".
- **נתקע** → "got involved".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **אלפין** → "alefs".
- **רשע** בן רשע → "Wicked one, son of a wicked one".
- **רוע** לב → "wickedness of heart".
- **קטיגורין** → "accuser" (transliteration only in the note).
- **דין / נקבל**: in "אם הלכה נקבל ואם לדין יש תשובה" דין → "reasoning"; בקש →
  "in haste" (note: the conjecture "with straw").
- **כמתלהמים** → "ravers" (Prov 26:18 in `apparatus`).
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- אין לנו **להחמיר** → "we may not be stringent".
- הגיע **חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
- **דחלות** (Aramaic, once, unclear): the best reading by grammar, the others in
  `readings`.
- **ברותא** (for ברייתא) → "berota" (note).
- **ס' / סי' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in
  his numbered sections; by context, the choice noted. A cited book's סי' →
  "section 172", never "§"; his own "לעיל סי' מ"ב" / "לקמן סי' נ"ח" → "above,
  §42" / "below, §58".
- **סכלות** והוללות → "folly and madness".
- לפי **דבריו** (his disclaimer, kept) → "(according to his words)".
- לפי **דעתם** (his disclaimer, kept) → "(in their opinion)".
- **רחמנא ליצלן** → "the Merciful One save us".
- **ידידי** → "my friend" (second person kept).
- קורא **נעים / נעלם** → "pleasant reader".
- **אחי** → "my brother" (second person kept).
- **שאול / תשובה / השואל / המשיב** (the correspondence of chapter 2): שאול
  שאל → "Question"; תשובה → "Reply"; חזר השואל והשיב → "The questioner replied
  again"; המשיב → "the respondent"; each letter in its writer's voice, the
  parties identified no further than the text.
- בית **הספר** (non-technical 1931 Hebrew) → "school".
- **הסתדרות** (non-technical 1931 Hebrew) → "organisation".
- מכתבי **עתים** (non-technical 1931 Hebrew) → "newspapers".
- **לפרט**: "התר"ן לפרט גדול" → "5650 of the full count".
- **חסר / חטר**: "(חסר)" → "(missing)", a lacuna sign; "(חטר)" is a variant of
  it, recorded.

## Quotations

- **Scope.** All quoted text (Zohar, Bavli, Yerushalmi, Targum, Biblical
  Aramaic, Midrash, medieval, Judeo-Arabic) from the author's words; no Hebrew,
  Aramaic or Arabic in `english`.
- **Marking.** A source's words introduced by a formula (וז"ל, כתב, אמרו,
  גרסינן, שנאמר, כדכתיב, שנינו, הכי קאמר, אמר ר' X) or closed by עכ"ל / ע"כ:
  double quotation marks; a paraphrase or report ("כתב ש…"): none; no
  formula: marks if the source is certain, else none and a note. Nesting:
  double, single, double. A reply by ויאמר inside a quotation: double marks;
  his report of a sage's question: unmarked. Over several paragraphs
  (2:18-2:25): each opens with a mark, only the last closes. Plain text (no
  italics, no labels); the language in `grammar` ("Zohar Aramaic",
  "Judeo-Arabic"). Formulas are translated, outside the marks; glosses inside
  a quotation stay in parentheses where they stand, translated.
- **Exegesis on a Hebrew word** (את; אם אין → "Ayin"): the word stays
  transliterated ("not 'to the LORD your God' but 'et the LORD your God'"),
  explained in `grammar`.
- **Cut or changed quotation.** Only what is on the page: supply nothing וכו'
  omits, complete no verse (no "before Me"), restore no source; a verse's
  fixed wording covers only the words quoted. His comma joining two verses
  stays; an ellipsis only where he marks one (וכו', וגו'). His wording against
  the source: follow him, the source in `apparatus`.
- **Bible verses**, quoted or alluded to (*melitzah*, unmarked): from the
  Hebrew as it stands, plain modern English, not King James; each in
  `apparatus`. An echoed verse (אל דבר ה', בטרם תחיל) governs the sense of its
  words; a poetic phrase
  translated, a note naming the referent; a plural stays plural, no subject
  added; ketiv translated, the qere noted.
- **Idioms**: by sense, the literal in `literal`; a saying he glosses
  himself: both translated, unmerged.
- **References** as he gives them, English form, Arabic numerals; add,
  complete, convert or correct nothing; never the 2015 PDF's; none for
  unreferenced Zohar, Mishnah or Talmud, even in `apparatus` (Bible verses
  only). פ' + numeral → "chapter" for a book without parashot or a numeral
  beyond them, else "parashah" (noted). ד' + number after a book: the
  numerals ("4, 12"), noted when he writes דף nearby.
- **Section numbers.** His sections (ב) at 3:1 to קמד) at 10:6): the label →
  "§2."; irregular ones as written ("קיה)" → "§115", letters noted; the gap
  קכט)–קלא) noted, not repaired). List markers (א), ב)) → "1)", "2)" (note).
- **וז"ל** → "and these are his words:"; וז"ל + a name → "and these are the
  words of X:" (outside the marks).
- **עכ"ל** → "Thus far his words." (outside the marks).
- **בגליון**: "ע"כ לשון הזהר עם פי' הרש"ב בגליון" → "Thus far the wording of
  the Zohar, with the explanation of the RaShB in the margin."; "(פירש בגליון
  …)" inside a quotation → "(the margin explains: …)".
- **פי'** inside a quotation: "(פי' …)" → "(explanation: …)".
- **זכאה / דכאה / חולקיה / הולקיה / חולקיהון / חולקיהן**: זכאה חולקיה /
  חולקיהון → "happy is his / their portion"; זכאה חולקיהון דישראל → "happy is
  the portion of Israel". Quoted as sales talk (9:4, 9:16): in quotation
  marks, no irony added. Variant spellings ("הולקיה", "דכאה"): the evident
  word, recorded in `apparatus`.
- **ווי** ליה מאן דלא → "woe to him who does not".
- דמאן **דעביד** כך וכך → "of one who does such-and-such".
- **תא חזי** → "come and see".
- **רזא דמלה** → "the secret of the matter".
- פתח ר' **שמעון** → "Rabbi Shimon opened".
- **ואלהי** in a cut verse → "and the gods of" (note); the verse is not
  completed.
- **אבותיהם** for "אבתיכם": follow his wording; the verse in `apparatus`.
- בחלק **שפתיו** → "with the smoothness of his lips".
- **כמתלהלה** היורה זיקים חצים ומות → "like a madman who shoots firebrands,
  arrows and death".
- **צנומות** דקות שדופות קדים → "withered, thin, blasted by the east wind".
- **בחבלי** השוא … כעבות העגלה → "with cords of falsehood … as with
  cart-ropes".
- כעיר **פרוצה** אין חומה → "like a city breached, without a wall".
- להוציא **ישן** … מפני חדש → "to clear out the old before the new".
- **מילדי** נכרים → "from the children of strangers".
- הבל **וריק** → "vanity and emptiness".
- תהו **ובהו** → "formlessness and void".
- ריש **מלין** אמר (Dan 7:1) → "he stated the sum of the matter".
- **קצוצי** פאה (Jer 9:25) → "those with clipped temples" (context note: desert
  peoples).
- הר **המור** → "the mountain of myrrh" (a poetic phrase, a note naming the
  referent).
- מי **אתי** (ketiv): translated, the qere noted.
- **עמדו** (Job 32:16): a plural stays plural, no subject added.
- **שמע** ישראל (Deut 6:4) → "Hear, O Israel: the LORD our God, the LORD is
  one." (Four units in order, which the Zohar reading in 6:1 needs.)
- **עמדי** (Deut 32:39) → "See now that I, I am He, and there is no god with
  Me."
- **מבלעדי / אחרון** (Isa 44:6) → "I am the first and I am the last, and
  besides Me there is no God."
- **אנכי / יהיה** (Exod 20:2–3) → "I am the LORD your God … You shall have no
  other gods before Me."
- **יחרם** (Exod 22:19) → "He who sacrifices to the gods shall be destroyed,
  save to the LORD alone."
- **תירא / תעבוד / תעבד** (Deut 6:13) → "The LORD your God you shall fear, and
  Him you shall worship."
- **רחום / וחנון** (Exod 34:6) → "the LORD, the LORD, God merciful and
  gracious, slow to anger".
- **כדנה / אלהיא / יאבדו** (Jer 10:11) → "Thus shall you say to them: the gods
  that did not make the heavens and the earth shall perish from the earth and
  from under these heavens."
- **תעמוד / רעך** (Lev 19:16) → "stand by the blood of your fellow" (no
  "idly").
- **פילא**: "דמעיילי פילא בקופא דמחטא" → "who bring an elephant through the eye
  of a needle" (idiom: by sense, the literal in `literal`).
- **נמטיה**: "ואפריין נמטיה" → "and we acknowledge our debt to him".
- **דמרפסן**: "הני מילי דמרפסן איגרא" → "words that split roofs".
- **סלקינן**: "בהא סלקינן ובהא נחתינן" → "this is where we come out".
- **בנדון** דידן → "in our case".
- **רמב"ם / רס"ג / הכוזרי / הלוי / בחיי / סמ"ג / אלבו / העיקרים / רוקח / שבילי**
  (medieval authors quoted): from the Hebrew he quotes, not the originals,
  standard editions or published translations; a certain difference may be
  noted in `apparatus`.
- **אללה / הדה / אלגהת** (any Judeo-Arabic passage: Saadia on Job; Choter ben
  Shelomo; a Muslim polemic; the proverb in 2:13): from the Arabic as it
  stands, never transliterated in the English. Ambiguous spelling: the evident
  word (note), the other in `apparatus`; אלי' (alif maqsura) is ilā "to"; a
  scribal letter change (ר for ד) noted, not corrected; an unexplained word
  (שברא): "[unclear]", the letters in `apparatus`. A rendering from Saadia's
  translation (אלבלד שברא): as given, Saadia's translation named in
  `apparatus`; his אמנתא / אמונתנו pairs: the Arabic as written, his Hebrew
  gloss in a note. Glosses inside the Arabic ("אתתלית (תלת קשרי דמהימנותא)")
  stay in parentheses, translated as his. (Q4.)
- **אצהאר / גילוי**: ספר אצהאר אלחק = גילוי האמת; the title as each place writes
  it ("Iẓhar al-Chaqq", "Gilui ha-Emet"), a note that they are one book; its
  author unnamed.
- **העתקתו / תרגום**: his own Hebrew translations (10:7–10:9; "[תרגום הרב
  המחבר: …]") are separate segments, each from its own Hebrew; a difference
  in sense in `apparatus`.
- **המיחדים / אלתוחיד**: "כת המיחדים" = "אהל אלתוחיד" → "the upholders of the
  unity".
- **דף** → "folio" (no "folio" he does not write: "זהר בראשית כ״ב" → "Zohar,
  Bereshit 22"); "בשלח דף ס"ד (ח"ב סז, א)" → "Be-shallach, folio 64 (vol. 2,
  67a)" (nothing converted).
- **ע"א / ע"ב** → "a / b".
- **ח"א / ח"ב** → "vol. 1" / "vol. 2".
- **פרשת** → the parashah ("Zohar, Balak, folio 191b").
- **מיתות**: פ' ד' מיתות (פ' + a chapter title) → "Four Death Penalties".
- **שער / ענף / מאמר** → "Gate / Branch / Treatise".
- **סנהדרין** → "Sanhedrin 99"; סנהדרין רי"ף → "Sanhedrin, [in] the Rif".
- **הל'** X after the Rambam → "Laws of X" (noted once per unit: Mishneh
  Torah).

## Style

- Plain, exact, modern English, the words at their own strength (אליל
  "idol"). No sarcasm of your own: no scare quotes, "so-called", "sic", or
  added "!". Keep his every "!", rhetorical question and disclaimer, and the
  second person.
- Long periods may be divided: no clause dropped, no connective that changes
  the logic, the order kept. Hebrew full stops: שכל / שהוא opening a sentence:
  the stop kept, a main clause by his syntax (note on the ש); a fragment
  (ושהם, והפטרה, ולדעת, ורבינו): a main clause when it continues the verb
  before, a comma when it continues a list (note); a main clause after a stop
  (ואם ראיתי. הם מגלים): joined with a comma (note); אם ... או after a stop:
  as punctuated if the English has a main clause, else a colon to the
  sentence before or a direct alternative question; a period completed in the
  next segment: the stops kept, the continuation in the segment `note`,
  semicolons only where unreadable.
- ל before a person's name is possessive (לרשב"י → "Rabbi Shimon's"), never
  the subject.
- Numbers: letter numerals → Arabic, as written whatever the tradition counts
  (the usual count in a note). A year without thousands → full ("תר"ע" →
  "5670"); the civil year in `apparatus`. Where gematria is the point, the
  number in English, the letters in a note.
- Witnesses: square brackets ("[מתחלת באלוהיות]") are the transcriber's
  corrections: translate the corrected reading, record both. A word that does
  not fit but is close to a likelier one (יכינו for יבינו): by the context; a
  printed word echoing a verse's likelier word (נבוש for נמוש): the printed
  word; both words in a note and `apparatus`. Rule 3 covers every
  Wikisource/PDF difference that bears on the English.

## Pending consistency pass (before the PR)

- 2:18 (perek-2-09) closes the Menasseh quotation too early; 2:19-2:22
  (perek-2-10) carry no opening marks. Patch both (multi-paragraph ruling).
- 3:6 (perek-3-03): "the ignorant" for עמי הארץ becomes "the unlearned".
- 3:3 (perek-3-02): "Beit Din" becomes "court" ("the Great Court").
- 3:7-3:8 (perek-3-04): remove the named author (Safrin of Komarno) from the
  note and context on היכל הברכה.
- 3:9 (perek-3-05): check העבודה הזאת (Exod 12:26) against "this service".
- Spell אלקבולי one way across units (search for "Alkabetz" and "Alkaboli").

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

- Q4: Judeo-Arabic and Aramaic are Sonnet's, not Haiku's: `models.json` lists
  the four Judeo-Arabic units (5:3, 7:24, 9:10, 2:13) and the 15 units with at
  least six Aramaic marker words (a heuristic; the Sonnet review covers every
  unit). The `lang` tags stay "heb".
- Q7: the project's transliteration scheme (above).

Not yet answered; the default stands and can be changed before the review
stage (a change to 1, 3 or 5 is a find-and-patch):

- Q1 title: "The Wars of the LORD" (Num 21:14).
- Q2 Wikisource additions: translated as base text, unmarked per word; one
  general line in `apparatus` on the unit's first segment until the 1931 scan
  is collated.
- Q3 השי"ת: "the Name, may He be blessed".
- Q5 קבלה האמתית: "the true tradition", with a note.
- Q6 author's name: patronymic kept as written, "Sali'".
