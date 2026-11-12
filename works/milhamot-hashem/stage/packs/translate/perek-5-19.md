# Translation pack: milhamot-hashem perek-5-19 (Milchamot ha-Shem 5:28–5:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-19`.

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
- **אצילות / באצילות / בריאה**: fixed, translated, the verbs linked:
  "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; נאצל → "emanated"; המאציל → "the Emanator". The four worlds
  (אבי"ע): "Emanation, Creation, Formation, Making" (verbs: "emanated,
  created, formed, made"). נאצל and נברא stay distinct ("נברא ולא בורא" →
  "created, not creator").
- **צמצום**: fixed, "contraction"; צמצם / נצטמצם → "contracted"; "tzimtzum" in a
  note first in a unit.
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **ריבה / לרבות** → "include", never "multiply".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
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


## Milchamot ha-Shem 5:28

Hebrew (Wikisource): נד) וזוהי אמונתינו וקבלתינו העיקרית שקבלנו מאבותינו דור אחר דור עד משה רבי' ועד אברהם אבינו. והיא שוה ומתאמת לקבלת הישמעאלים המקובלת אצלם ממחוקקם שקיבל מישראל כנז' לעיל. כי ממקור אחד מאברהם אבינו וכו' כנז"ל. וחיבה יתירה נודעת לנו. שקירבנו הקב"ה לפני הר סיני, ונגלה כבודו לנו שם, ונתן בקולו קול עוז לכל קהלינו לאמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים על פני. והנחילנו תורה ומצות וחקים. ואין להביא ראיה מדורות ימי הבינים שהאמינו בקבלה החדשה לפי שהיו לומדים ס' הזהר בגרסא בעלמא מבלי הבין כידוע. וקיימו בעצמם לגרוס איניש לבד, אבל והדר ליסבר לא קיימו, ולא העמיקו חקר בדברי הספר הנזכר ולא בשאר ספרי הקבלה החדשה להבין שהם הורסים פנות תורתינו הקדושה ועיקריה ויסודותיה שביארם הרמב"ם בפי' המשנה. ולפיכך הם אומרים יגדל אלהים חי בכל ליל שבת קודש ועוסקים בגירסת ס' הזהר ולא ידעו שהם סותרים זה את זה כמו שכ' רבי' תם ן' יחיא ז"ל. ואלו ידעו והבינו שהוא כן היו מרחיקים עצמם ממנה תכלית ההרחקה. כי עין רואה ואזן שומעת שהם מגשימים אפי' את האין סוף שהוא (לפי דבריהם) לעילא לעילא, ומציירים אותו ככדור עיגול וחלול באמצעיתו פנוי וריקם המתהוה על ידי הצמצום הנעשה בו לצדדין שוה בהשואה א'. כדי להמציא מקום שיוכלו להיות שם במקום הריקן כל הנאצלים והנבראים והנוצרים והנעשים. ואז נמשך מאור האין סוף קו דק מן האור העגול של האין סוף מלמעלה למטה ונעשו ספירות העיגולים מאותו הקו הנמשך ממנו עגול בתוך עגול. וגם נשתלשל ונמשך הקו ההוא ביושר תוך החלל ונעשו ספירות דיושר. אשר קראום מדות וכלים והם עיקר האלהות, וגם נשמתם שהוא האור הפנימי גם מלבושם שהוא האור המקיף הכל נמשך ונשתלשל מעצם האין סוף המקיף את הכל, ממנו גוף, ממנו נשמה, ממנו מלבוש. כלם נמשכים מעצמות האור ההוא כמו שכתבתי לעיל סי' ט"ל בשם שושן סודות ומשנת חסידים שגם הכלים אלהות גמור כי מחומר אחד קורצו. והסבות הללו נמשכים ונעשה כל אחד סבה למה שלמטה ממנה. זה סבה לזה וזה סבה לזה. הסבה הראשונה מפרצופי היושר ופרצופים הכוללים אדם קדמאה, ואחריו אדם קדמון, ואחריו עתיקא קדישא, ואחריו אריך אנפין, וממנו נחיו אבא ואימא, ומהם נהיו זעיר ונוקביה, והוא הנקרא בן לאבא ואימא. כל הפרצופים הללו נשתלשלו זה מזה, ונקרא כל אחד מהם עילה וסבה לפרצוף העלול ממנו אחריו. ואבא ואימא אעטרו לז"א בתרין עטרין והשליטוהו על כל הברואים וצוו (לפי דבריהם) לעבדו ליראה את אשתו, שריבה אותה הכתוב ליראה באמרו את ה' אלהיך תירא, את לרבות נוקבא דיליה ליראה אותה כמו שנתבאר לעיל סי' כ"ו כ"ז כ"ח, מה שלא עלה על דעת שמעון העמסוני ואמרי לה נחמיה העמסוני ור' עקיבה. ויחסו העבודה לזעיר אנפין דוקא כמו שנ' ואותו תעבוד והוא נקרא בפיהם יהוה אלהינו. והוא הסבה והפרצוף האחרון שבפרצופים האצילות. ואמרו שהוא הוא אלהינו ואנחנו עמו ועבדיו כמבואר בכסא אליהו וס' הברית ויושר לבב דף כ"א ול"א ול"ב. וענין השתלשלותם בעץ חיים שער עיגולים ויושר וכסא אלהיו.

| id | word | punctuation after |
|---|---|---|
| 5:28.1 | נד | ) |
| 5:28.2 | וזוהי |  |
| 5:28.3 | אמונתינו |  |
| 5:28.4 | וקבלתינו |  |
| 5:28.5 | העיקרית |  |
| 5:28.6 | שקבלנו |  |
| 5:28.7 | מאבותינו |  |
| 5:28.8 | דור |  |
| 5:28.9 | אחר |  |
| 5:28.10 | דור |  |
| 5:28.11 | עד |  |
| 5:28.12 | משה |  |
| 5:28.13 | רבי | ' |
| 5:28.14 | ועד |  |
| 5:28.15 | אברהם |  |
| 5:28.16 | אבינו | . |
| 5:28.17 | והיא |  |
| 5:28.18 | שוה |  |
| 5:28.19 | ומתאמת |  |
| 5:28.20 | לקבלת |  |
| 5:28.21 | הישמעאלים |  |
| 5:28.22 | המקובלת |  |
| 5:28.23 | אצלם |  |
| 5:28.24 | ממחוקקם |  |
| 5:28.25 | שקיבל |  |
| 5:28.26 | מישראל |  |
| 5:28.27 | כנז | ' |
| 5:28.28 | לעיל | . |
| 5:28.29 | כי |  |
| 5:28.30 | ממקור |  |
| 5:28.31 | אחד |  |
| 5:28.32 | מאברהם |  |
| 5:28.33 | אבינו |  |
| 5:28.34 | וכו | ' |
| 5:28.35 | כנז"ל | . |
| 5:28.36 | וחיבה |  |
| 5:28.37 | יתירה |  |
| 5:28.38 | נודעת |  |
| 5:28.39 | לנו | . |
| 5:28.40 | שקירבנו |  |
| 5:28.41 | הקב"ה |  |
| 5:28.42 | לפני |  |
| 5:28.43 | הר |  |
| 5:28.44 | סיני | , |
| 5:28.45 | ונגלה |  |
| 5:28.46 | כבודו |  |
| 5:28.47 | לנו |  |
| 5:28.48 | שם | , |
| 5:28.49 | ונתן |  |
| 5:28.50 | בקולו |  |
| 5:28.51 | קול |  |
| 5:28.52 | עוז |  |
| 5:28.53 | לכל |  |
| 5:28.54 | קהלינו |  |
| 5:28.55 | לאמר |  |
| 5:28.56 | אנכי |  |
| 5:28.57 | ה | ' |
| 5:28.58 | אלהיך | , |
| 5:28.59 | לא |  |
| 5:28.60 | יהיה |  |
| 5:28.61 | לך |  |
| 5:28.62 | אלהים |  |
| 5:28.63 | אחרים |  |
| 5:28.64 | על |  |
| 5:28.65 | פני | . |
| 5:28.66 | והנחילנו |  |
| 5:28.67 | תורה |  |
| 5:28.68 | ומצות |  |
| 5:28.69 | וחקים | . |
| 5:28.70 | ואין |  |
| 5:28.71 | להביא |  |
| 5:28.72 | ראיה |  |
| 5:28.73 | מדורות |  |
| 5:28.74 | ימי |  |
| 5:28.75 | הבינים |  |
| 5:28.76 | שהאמינו |  |
| 5:28.77 | בקבלה |  |
| 5:28.78 | החדשה |  |
| 5:28.79 | לפי |  |
| 5:28.80 | שהיו |  |
| 5:28.81 | לומדים |  |
| 5:28.82 | ס | ' |
| 5:28.83 | הזהר |  |
| 5:28.84 | בגרסא |  |
| 5:28.85 | בעלמא |  |
| 5:28.86 | מבלי |  |
| 5:28.87 | הבין |  |
| 5:28.88 | כידוע | . |
| 5:28.89 | וקיימו |  |
| 5:28.90 | בעצמם |  |
| 5:28.91 | לגרוס |  |
| 5:28.92 | איניש |  |
| 5:28.93 | לבד | , |
| 5:28.94 | אבל |  |
| 5:28.95 | והדר |  |
| 5:28.96 | ליסבר |  |
| 5:28.97 | לא |  |
| 5:28.98 | קיימו | , |
| 5:28.99 | ולא |  |
| 5:28.100 | העמיקו |  |
| 5:28.101 | חקר |  |
| 5:28.102 | בדברי |  |
| 5:28.103 | הספר |  |
| 5:28.104 | הנזכר |  |
| 5:28.105 | ולא |  |
| 5:28.106 | בשאר |  |
| 5:28.107 | ספרי |  |
| 5:28.108 | הקבלה |  |
| 5:28.109 | החדשה |  |
| 5:28.110 | להבין |  |
| 5:28.111 | שהם |  |
| 5:28.112 | הורסים |  |
| 5:28.113 | פנות |  |
| 5:28.114 | תורתינו |  |
| 5:28.115 | הקדושה |  |
| 5:28.116 | ועיקריה |  |
| 5:28.117 | ויסודותיה |  |
| 5:28.118 | שביארם |  |
| 5:28.119 | הרמב"ם |  |
| 5:28.120 | בפי | ' |
| 5:28.121 | המשנה | . |
| 5:28.122 | ולפיכך |  |
| 5:28.123 | הם |  |
| 5:28.124 | אומרים |  |
| 5:28.125 | יגדל |  |
| 5:28.126 | אלהים |  |
| 5:28.127 | חי |  |
| 5:28.128 | בכל |  |
| 5:28.129 | ליל |  |
| 5:28.130 | שבת |  |
| 5:28.131 | קודש |  |
| 5:28.132 | ועוסקים |  |
| 5:28.133 | בגירסת |  |
| 5:28.134 | ס | ' |
| 5:28.135 | הזהר |  |
| 5:28.136 | ולא |  |
| 5:28.137 | ידעו |  |
| 5:28.138 | שהם |  |
| 5:28.139 | סותרים |  |
| 5:28.140 | זה |  |
| 5:28.141 | את |  |
| 5:28.142 | זה |  |
| 5:28.143 | כמו |  |
| 5:28.144 | שכ | ' |
| 5:28.145 | רבי | ' |
| 5:28.146 | תם |  |
| 5:28.147 | ן | ' |
| 5:28.148 | יחיא |  |
| 5:28.149 | ז"ל | . |
| 5:28.150 | ואלו |  |
| 5:28.151 | ידעו |  |
| 5:28.152 | והבינו |  |
| 5:28.153 | שהוא |  |
| 5:28.154 | כן |  |
| 5:28.155 | היו |  |
| 5:28.156 | מרחיקים |  |
| 5:28.157 | עצמם |  |
| 5:28.158 | ממנה |  |
| 5:28.159 | תכלית |  |
| 5:28.160 | ההרחקה | . |
| 5:28.161 | כי |  |
| 5:28.162 | עין |  |
| 5:28.163 | רואה |  |
| 5:28.164 | ואזן |  |
| 5:28.165 | שומעת |  |
| 5:28.166 | שהם |  |
| 5:28.167 | מגשימים |  |
| 5:28.168 | אפי | ' |
| 5:28.169 | את |  |
| 5:28.170 | האין |  |
| 5:28.171 | סוף |  |
| 5:28.172 | שהוא |  |
| 5:28.173 | לפי |  |
| 5:28.174 | דבריהם | ) |
| 5:28.175 | לעילא |  |
| 5:28.176 | לעילא | , |
| 5:28.177 | ומציירים |  |
| 5:28.178 | אותו |  |
| 5:28.179 | ככדור |  |
| 5:28.180 | עיגול |  |
| 5:28.181 | וחלול |  |
| 5:28.182 | באמצעיתו |  |
| 5:28.183 | פנוי |  |
| 5:28.184 | וריקם |  |
| 5:28.185 | המתהוה |  |
| 5:28.186 | על |  |
| 5:28.187 | ידי |  |
| 5:28.188 | הצמצום |  |
| 5:28.189 | הנעשה |  |
| 5:28.190 | בו |  |
| 5:28.191 | לצדדין |  |
| 5:28.192 | שוה |  |
| 5:28.193 | בהשואה |  |
| 5:28.194 | א | '. |
| 5:28.195 | כדי |  |
| 5:28.196 | להמציא |  |
| 5:28.197 | מקום |  |
| 5:28.198 | שיוכלו |  |
| 5:28.199 | להיות |  |
| 5:28.200 | שם |  |
| 5:28.201 | במקום |  |
| 5:28.202 | הריקן |  |
| 5:28.203 | כל |  |
| 5:28.204 | הנאצלים |  |
| 5:28.205 | והנבראים |  |
| 5:28.206 | והנוצרים |  |
| 5:28.207 | והנעשים | . |
| 5:28.208 | ואז |  |
| 5:28.209 | נמשך |  |
| 5:28.210 | מאור |  |
| 5:28.211 | האין |  |
| 5:28.212 | סוף |  |
| 5:28.213 | קו |  |
| 5:28.214 | דק |  |
| 5:28.215 | מן |  |
| 5:28.216 | האור |  |
| 5:28.217 | העגול |  |
| 5:28.218 | של |  |
| 5:28.219 | האין |  |
| 5:28.220 | סוף |  |
| 5:28.221 | מלמעלה |  |
| 5:28.222 | למטה |  |
| 5:28.223 | ונעשו |  |
| 5:28.224 | ספירות |  |
| 5:28.225 | העיגולים |  |
| 5:28.226 | מאותו |  |
| 5:28.227 | הקו |  |
| 5:28.228 | הנמשך |  |
| 5:28.229 | ממנו |  |
| 5:28.230 | עגול |  |
| 5:28.231 | בתוך |  |
| 5:28.232 | עגול | . |
| 5:28.233 | וגם |  |
| 5:28.234 | נשתלשל |  |
| 5:28.235 | ונמשך |  |
| 5:28.236 | הקו |  |
| 5:28.237 | ההוא |  |
| 5:28.238 | ביושר |  |
| 5:28.239 | תוך |  |
| 5:28.240 | החלל |  |
| 5:28.241 | ונעשו |  |
| 5:28.242 | ספירות |  |
| 5:28.243 | דיושר | . |
| 5:28.244 | אשר |  |
| 5:28.245 | קראום |  |
| 5:28.246 | מדות |  |
| 5:28.247 | וכלים |  |
| 5:28.248 | והם |  |
| 5:28.249 | עיקר |  |
| 5:28.250 | האלהות | , |
| 5:28.251 | וגם |  |
| 5:28.252 | נשמתם |  |
| 5:28.253 | שהוא |  |
| 5:28.254 | האור |  |
| 5:28.255 | הפנימי |  |
| 5:28.256 | גם |  |
| 5:28.257 | מלבושם |  |
| 5:28.258 | שהוא |  |
| 5:28.259 | האור |  |
| 5:28.260 | המקיף |  |
| 5:28.261 | הכל |  |
| 5:28.262 | נמשך |  |
| 5:28.263 | ונשתלשל |  |
| 5:28.264 | מעצם |  |
| 5:28.265 | האין |  |
| 5:28.266 | סוף |  |
| 5:28.267 | המקיף |  |
| 5:28.268 | את |  |
| 5:28.269 | הכל | , |
| 5:28.270 | ממנו |  |
| 5:28.271 | גוף | , |
| 5:28.272 | ממנו |  |
| 5:28.273 | נשמה | , |
| 5:28.274 | ממנו |  |
| 5:28.275 | מלבוש | . |
| 5:28.276 | כלם |  |
| 5:28.277 | נמשכים |  |
| 5:28.278 | מעצמות |  |
| 5:28.279 | האור |  |
| 5:28.280 | ההוא |  |
| 5:28.281 | כמו |  |
| 5:28.282 | שכתבתי |  |
| 5:28.283 | לעיל |  |
| 5:28.284 | סי | ' |
| 5:28.285 | ט"ל |  |
| 5:28.286 | בשם |  |
| 5:28.287 | שושן |  |
| 5:28.288 | סודות |  |
| 5:28.289 | ומשנת |  |
| 5:28.290 | חסידים |  |
| 5:28.291 | שגם |  |
| 5:28.292 | הכלים |  |
| 5:28.293 | אלהות |  |
| 5:28.294 | גמור |  |
| 5:28.295 | כי |  |
| 5:28.296 | מחומר |  |
| 5:28.297 | אחד |  |
| 5:28.298 | קורצו | . |
| 5:28.299 | והסבות |  |
| 5:28.300 | הללו |  |
| 5:28.301 | נמשכים |  |
| 5:28.302 | ונעשה |  |
| 5:28.303 | כל |  |
| 5:28.304 | אחד |  |
| 5:28.305 | סבה |  |
| 5:28.306 | למה |  |
| 5:28.307 | שלמטה |  |
| 5:28.308 | ממנה | . |
| 5:28.309 | זה |  |
| 5:28.310 | סבה |  |
| 5:28.311 | לזה |  |
| 5:28.312 | וזה |  |
| 5:28.313 | סבה |  |
| 5:28.314 | לזה | . |
| 5:28.315 | הסבה |  |
| 5:28.316 | הראשונה |  |
| 5:28.317 | מפרצופי |  |
| 5:28.318 | היושר |  |
| 5:28.319 | ופרצופים |  |
| 5:28.320 | הכוללים |  |
| 5:28.321 | אדם |  |
| 5:28.322 | קדמאה | , |
| 5:28.323 | ואחריו |  |
| 5:28.324 | אדם |  |
| 5:28.325 | קדמון | , |
| 5:28.326 | ואחריו |  |
| 5:28.327 | עתיקא |  |
| 5:28.328 | קדישא | , |
| 5:28.329 | ואחריו |  |
| 5:28.330 | אריך |  |
| 5:28.331 | אנפין | , |
| 5:28.332 | וממנו |  |
| 5:28.333 | נחיו |  |
| 5:28.334 | אבא |  |
| 5:28.335 | ואימא | , |
| 5:28.336 | ומהם |  |
| 5:28.337 | נהיו |  |
| 5:28.338 | זעיר |  |
| 5:28.339 | ונוקביה | , |
| 5:28.340 | והוא |  |
| 5:28.341 | הנקרא |  |
| 5:28.342 | בן |  |
| 5:28.343 | לאבא |  |
| 5:28.344 | ואימא | . |
| 5:28.345 | כל |  |
| 5:28.346 | הפרצופים |  |
| 5:28.347 | הללו |  |
| 5:28.348 | נשתלשלו |  |
| 5:28.349 | זה |  |
| 5:28.350 | מזה | , |
| 5:28.351 | ונקרא |  |
| 5:28.352 | כל |  |
| 5:28.353 | אחד |  |
| 5:28.354 | מהם |  |
| 5:28.355 | עילה |  |
| 5:28.356 | וסבה |  |
| 5:28.357 | לפרצוף |  |
| 5:28.358 | העלול |  |
| 5:28.359 | ממנו |  |
| 5:28.360 | אחריו | . |
| 5:28.361 | ואבא |  |
| 5:28.362 | ואימא |  |
| 5:28.363 | אעטרו |  |
| 5:28.364 | לז"א |  |
| 5:28.365 | בתרין |  |
| 5:28.366 | עטרין |  |
| 5:28.367 | והשליטוהו |  |
| 5:28.368 | על |  |
| 5:28.369 | כל |  |
| 5:28.370 | הברואים |  |
| 5:28.371 | וצוו |  |
| 5:28.372 | לפי |  |
| 5:28.373 | דבריהם | ) |
| 5:28.374 | לעבדו |  |
| 5:28.375 | ליראה |  |
| 5:28.376 | את |  |
| 5:28.377 | אשתו | , |
| 5:28.378 | שריבה |  |
| 5:28.379 | אותה |  |
| 5:28.380 | הכתוב |  |
| 5:28.381 | ליראה |  |
| 5:28.382 | באמרו |  |
| 5:28.383 | את |  |
| 5:28.384 | ה | ' |
| 5:28.385 | אלהיך |  |
| 5:28.386 | תירא | , |
| 5:28.387 | את |  |
| 5:28.388 | לרבות |  |
| 5:28.389 | נוקבא |  |
| 5:28.390 | דיליה |  |
| 5:28.391 | ליראה |  |
| 5:28.392 | אותה |  |
| 5:28.393 | כמו |  |
| 5:28.394 | שנתבאר |  |
| 5:28.395 | לעיל |  |
| 5:28.396 | סי | ' |
| 5:28.397 | כ"ו |  |
| 5:28.398 | כ"ז |  |
| 5:28.399 | כ"ח | , |
| 5:28.400 | מה |  |
| 5:28.401 | שלא |  |
| 5:28.402 | עלה |  |
| 5:28.403 | על |  |
| 5:28.404 | דעת |  |
| 5:28.405 | שמעון |  |
| 5:28.406 | העמסוני |  |
| 5:28.407 | ואמרי |  |
| 5:28.408 | לה |  |
| 5:28.409 | נחמיה |  |
| 5:28.410 | העמסוני |  |
| 5:28.411 | ור | ' |
| 5:28.412 | עקיבה | . |
| 5:28.413 | ויחסו |  |
| 5:28.414 | העבודה |  |
| 5:28.415 | לזעיר |  |
| 5:28.416 | אנפין |  |
| 5:28.417 | דוקא |  |
| 5:28.418 | כמו |  |
| 5:28.419 | שנ | ' |
| 5:28.420 | ואותו |  |
| 5:28.421 | תעבוד |  |
| 5:28.422 | והוא |  |
| 5:28.423 | נקרא |  |
| 5:28.424 | בפיהם |  |
| 5:28.425 | יהוה |  |
| 5:28.426 | אלהינו | . |
| 5:28.427 | והוא |  |
| 5:28.428 | הסבה |  |
| 5:28.429 | והפרצוף |  |
| 5:28.430 | האחרון |  |
| 5:28.431 | שבפרצופים |  |
| 5:28.432 | האצילות | . |
| 5:28.433 | ואמרו |  |
| 5:28.434 | שהוא |  |
| 5:28.435 | הוא |  |
| 5:28.436 | אלהינו |  |
| 5:28.437 | ואנחנו |  |
| 5:28.438 | עמו |  |
| 5:28.439 | ועבדיו |  |
| 5:28.440 | כמבואר |  |
| 5:28.441 | בכסא |  |
| 5:28.442 | אליהו |  |
| 5:28.443 | וס | ' |
| 5:28.444 | הברית |  |
| 5:28.445 | ויושר |  |
| 5:28.446 | לבב |  |
| 5:28.447 | דף |  |
| 5:28.448 | כ"א |  |
| 5:28.449 | ול"א |  |
| 5:28.450 | ול"ב | . |
| 5:28.451 | וענין |  |
| 5:28.452 | השתלשלותם |  |
| 5:28.453 | בעץ |  |
| 5:28.454 | חיים |  |
| 5:28.455 | שער |  |
| 5:28.456 | עיגולים |  |
| 5:28.457 | ויושר |  |
| 5:28.458 | וכסא |  |
| 5:28.459 | אלהיו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

