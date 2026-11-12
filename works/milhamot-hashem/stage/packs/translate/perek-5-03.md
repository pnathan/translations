# Translation pack: milhamot-hashem perek-5-03 (Milchamot ha-Shem 5:3–5:3)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-03`.

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
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

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
- **שיתוף / שתוף / לשתף / משותף**: associating another with God → fixed,
  "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever associates
  the name of Heaven with something else"). The Maimonidean term for a word said of God and creatures →
  "equivocation / equivocal" ("בשתוף השם בלבד" → "by equivocation of the term
  only"), never "association", Guide I in the note; "shared" only for a thing
  held in common. צירוף: "conjunction" of God; "permutation" in letter
  mysticism.
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
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
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
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 5:3

Hebrew (Wikisource): מב) נטה לך קורא נעים מהאמונות המדומות האלה. וראה מה שכ' אחד מעמודי הוראה אשר בית ישראל נשענים עליהם הוא רבי' הראש ז"ל במס' חולין פ' השוחט והובאו דבריו במבוא התלמוד וז"ל ומסתברא שכל תוספתא שלא נתפשטה עד אחר סידור הגמרא לאו דסמכא היא דמסתמא כיון שרצו חכמי ישראל לעשות חבור קיימא אמת. חקרו ודרשו וחיזרו לידע כל הספרים שנכתבו על דברי חכמים וביררו אותם שהם בר סמכא עכ"ל. ועיין לעיל סי' י"ב וי"ג כ' הר' סעדיה גאון שדבריו מסכימים לדברי הרא"ש. ומה שכ' המגיה במבוא התלמוד בגיליון לעיין בהרא"ש דהרמב"ם לא ס"ל הכי כבר ישב הר' ב"י דברי הרמב"ם שפסק בהתוספתא ושגם מסדר הגמרא סבירא ליה כהתוספתא ושגירסת הרמב"ם בגמרא שהה כמיעוט סימנים וכבר הביאה אבי' לעיל שהה כדי שיגביה הבהמה וירביצנה וישחוט מיעוט הסימנים. וכן היא הגירסא בהרמב"ם כ"י ישנים. ונמצא שגם הרמב"ם הכי ס"ל, וכבר העיר על זה ג"כ המ"מ הנדפס מחדש. ולפי זה כולי עלמא מודו שכל חיבור המיוחס לשום תנא או אמורא שלא נתפשט רק אחר סידור הגמ' לאו דסמכא הוא. וכמו שאמרו בירושלמי שכל משנה שלא נכנסה לחבורה אינה משנה. כמו שכתבתי לעיל סי' י"ב. ואם כך היא אמונתך שכל תפלותינו ועבודותינו לסבה הראשונה שהוא ה' אלהים אמת אשר אין ראשית לראשיתו והוא שהוציא את אבותינו ממצרים וכבודו נגלה עליהם בהר סיני בתתו לנו את התורה ואמר אנכי ה' אלהיך, לא יהיה לך אלהים אחרים על פני וכו' א"כ למה תבזה את בעלי אמונתך העוסקים יומם ולילה במשנה ובתלמוד ותקרא להם מינים ואפיקורסים ותחזיק ידי זרים לאמונתך אשר בילדי נכרים ישפיקו הוא ספר הזהר הנמסר לישראל מיד גוי ויחסוהו לרשב"י התנא בשקר. והטוב והנאות לנו להחזיק באמונת קדמונינו ביחוד השי"ת בלי צירוף ושתוף. לקיים ואל תטוש תורת אמך, ואל תבוז כי זקנה אמך. רצוני לומר שלא תבוז ותלעיג על אמונת קדמונינו ביחוד השי"ת ולא תזלזל בספריהם אף על פי שבלו והזקינו ונתישנו. לא כאשר עשו בימי הבינים שזלזלו בהם ואיבדו אותן ועשו אותן כריכות לספרים חדשים המודפסים. ויש שמכרו אותן בזול מאד. ובעיני ראיתי אחד מבני עמנו שהחליף את ספר מורה הנבוכים כ"י בלשון ערבי השוה לנו עתה שנים ושלשה זהובים בספר קטן הכמות והאיכות מודפס שוה חצי ריאל. והרבה מהמונינו מכרו ספרים ישנים שירשו מאבותיהם לערלים שבאו לארץ התימן כמשה בן נתנאל שישב בבית מ"ו אברהם צאלח. והוא בעצמו מכר לו אז נביאים וכתובים עם תרגום יונתן בן עוזיאל על נביאים. ותר' ר' סעדיה גאון ערבי על ישעיה ועל כתובים. גם משנה עם פי' הרמב"ם בערבי ס' קדשים בכתב מפואר עם ציורי בית קדשינו ותפארתינו וכלי הקדש במסכת מדות. ובעיני ראיתיה בידו אני והמנוח מו"ר יחיא קרח זלה"ה. ואין לאל ידינו להוציאה מתחת ידו. ועד עתה חסר לנו סדר קדשים ונשים מהמשנה כ"י עם פי' הרמב"ם ז"ל. ולא אלו בלבד מכרו לו. כי הרבה ספרים כתבי יד עתיקים נמכרו לו גם לזולתו בזול. כי היו בידם ככלי אין חפץ בו ולא היו יודעים לקרות בהם ואנן הוא דעבדינן לתורת אבותינו שלא תשתכח מישראל בעיר תהלה צנעא, כי קבענו למודינו במשנה עם פי' הרמב"ם בערבי. ופירושי רס"ג במקרא עם תלמוד ופוסקים. ומקודם לא היו לומדים בכל בתי כנסיות בלילות כי אם בספר הזהר בלי הבין כמפורסם. והיו משתבחים בעצמם שלמדו באותה לילה כך וכך דפים. ולימוד המשנה והתלמוד היה יקר בימים ההם אין חזון נפרץ. ורק בבקר לומדים מעט משנה לפני מו"ר יוסף קרח ז"ל עם פי' ר' עובדיה בבית הכנסת שרעבי. ותחלת לימודנו היה שם, וגם אחרי כן לפני בנו הרה"ג כמו"ר חיים קרח ז"ל בלילה. ורבים ממשכילי עמנו היו באים ללמוד. ואחרי כן התחלנו ללמוד פירוש הרמב"ם בערבי בלילי שבת מחצות. ואז נשתדל ידינו המנוח מו"ר חיים אלכסאר לכתוב לו משנה בפי' הרמב"ם ז"ל. ואחרים מהמשכילים גם הם נשתדלו לכתבם וללמוד בהם. ונמשך הדבר שגם בקצת בתי כנסיות נתקנאו לעשות כמעשנו כמאמר רז"ל קנאת סופרים תרבה חכמה. ולמדו גם כן הקדמת הרמב"ם לסדר זרעים ושלש עשרה עיקרים שכ' רבינו בפרק עשירי דסנהדרין בשום שכל ויבינו בהם ובדברי ר' סעדיה גאון וחובת הלבבות שהם עומדים נגד דברי הזהר והמקובלים החדשים. וראיתי להביא פה העיקר השני מהי"ג עיקרים עם קצת מדברי אחד מקדמונינו הר"ר חטר בן שלמה בן אלמעלם בפי' העיקר הזה וז"ל. ואלקאעדה אלתאניה וחדתה תעלי ודאלך אן הדא עלה אלכל ואחד, ליס כואחד אלגנס, ולא כואחד אלנוע. ולא כאלשכץ אלואחד אלמרכב אלדי ינקסם לאחאד כתירה ולא ואחד כאלגסם אלביסט אלואחד באלעדר אלדי יקבל אלאנקסאם ואלתגזי אלי לא נהאה. בל הו תע' ואחד בוחדה ליס כמתלהא וחדה בוגה, והדה אלקאעדה אלתאניה הי אלמדלול עליהא בקולה שמע ישראל ה' אלהינו ה' אחד. שרחהא אנה למא ביין לנא ונודה תע' אפתקר לביאן וחדאנייתה, אד האתין אלקאעדתין המא אצל שריעה משה רבי ע"ה, והמא אוול כל שריעה אלאהייה. ואכר כל שריעה נאמוסייה, לאן אלשריעה אלאלאהייה תבתדי מן אלאלאהיאת ותנתהי אלי אלנסטאנייאת, ואלשריעה אלנמוסייה עכס דאלך. ביאן דאך אן אלאנסאן אלא עלם אלונוד מן אבתדאיה אלי אנתהאיה אעני מם אעלא אעאליין אלי אספל אלסאפלין, פכיך מא נזל ואכר פי אלהבוט ספלא קלת אלוחדה צרורה ואזאדת ואלכהרה וכו. ואמא אדא אכד אלאנסאן אן יתרקא פי אלוגוד מן אספל אלי אלעא, פכיך מא אזאד פי אלתרקא עלווא צעפת אלכתרה ותקוות אלוהדה אלי אלוהחד אלאחד אלדי נשת מנה כל וחדה אחדיהה פי אלוגוד, והו ואחד אלוהדאנייה, והא אנא אדבר תלאתה מדאהב, אתנוייה, ואמגוס, ואנצארא והא אנא אביין כל מדהב מנהם ומן אטלע עלי כתבהם פצל דאלך, אמא אלתנוייה פאנהם יקולון אן לאלעאלם צאנעאן דקימאן, אחדהמא אלנור, והו יפעל אלכיר פי אלוגוד ולא יקדר עלי פעל אלשר! ואלתאני אלצלמה, והו יפעל אלשר פי אלוגוד ולא יקדר עלי פעל אלכיר, והם מכתלפין, מנהם מן יקול אן אלנור ואלצלמה חיאן, ומנחם מן יקול מייתאן. ואמא אלמגוס פאנהם יקולין אן לאלעאלם צאנעא חכימא קדירא מדברא גיר גסם, תזהוה ען פעל אלשדור וסמוה יזדאן, תם אנהם אצאפו היה אלאלאם אלכיאנה פי אלעאלם לאלשיטאן אהרמין בלגתהם, והם איצא מכתלפין, מנאהם מן יקול אנה מהדת וגסם, ומנאהם מאן יקול בקרם אהרמין ואנא גיר גסם, תם אכתלפו בסבב חדותה איצא, פקאלו אנה חדת ליזדאן פכרה, פקאל פי נפסה, פכיף יכון חאלי וחדי או מן ינאזעני, וקאלו אן אלשיטאן תולד מן תלך אלפכרה, תם אנה חדאת יזדאן וצאלחה אלי מרה, ודכרו אשיא ליס הנא מוצע דכרהא, ואלכיר ענדהם גניע אעני אלתנוייה ואלמגוס הו מא תשתהיה אלנפוס, ואלשר הו מא תכרהה אלנפוס! ואמא אלנצרא פאנהם יקולו אם אללה תע ואחד (עלי אלחקיקה) תלאתה אקאנים עלי אלחקיקה, אקנים אלאב, והו דאת אלבארי ענדהם, ואקנוס אלאכן והו אלעלם (חכמה) ואוקנום רוח אלקודש והו אלחיה (בינה הנקרא אלהים חיים) ואלתלאתה ורגעו ואחד. והם פריק מן ישראל. ואסתדלאלהם עלי אנה אלתלאתה ירגעו ואחד, מן קולח תע שמע ישראל ידוד אלהינו ידוד אחד, קאלו הדה תלתלאה אסמא לתלאתה מסמיין, וכלן אלתלאתה ירגעו ואחד, כמא קאל תע ידוד אלהינו ידוד אחד, ואלתאלב אנהם יעתקדו פי אלתוחיד לתהדר מן דאלך ארגע אלי גרץ אלקאעה וכו' עכ"ל הר' חטר ז"ל ע"ש. [תרגום הרב המחבר: היסוד השני יחודו יתברך וזהו לידע שזה האחד שהוא סיבת הכול, אחד לא כאחד הסוג, ולא כאחד המין, ולא כאחד המורכב אשר יחולק לחלקים רבים, ולא כגוף הפשוט שהוא אחד במנין (כמו המים שהוא יסוד פשוט) (חסר) ומתחלק לאין קץ. אבל הוא יתברך אחד, אחדות שאין כמוה אחדות בשום פנים, וזה היסוד השני הוא הנלמד מאומרו: "שמע ישראל ה"א ה' אחד!". ביאור היסוד הזה, הואיל וכבר באר רבנו מציאותו ית', הוצרך לבאר אחדותו. כי אלו השני יסודות הם עיקר תורת משרע"ה. והמה ראשית לכל תורה אלוהית וסוף לכל תורה נימוסית. כי התורה האלוהית (מתחלתאמן האלהיות) [מתחלת באלוהיות] וגומרת בגשמיות והתורה הנימוסית היפך זה. ביאור זה שהאדם אם ידע מציאות כל הנמצאים מראש ועד סוף, ר"ך ממעלה מעלה ועד מטה מטה, כל עוד שירד בידיעתו מטה נתמעטה האחדות ונתוסף הריבוי עד שמגיע לעולם השפל שבלי ספק לא ימד ולא יספר סוגי מיניו כ"ש אישי כל מין ומין וכו'. אבל אם בא לעלות בנמצאים ממטה למעלה, הנה כל מה שיוסיף לעלות יתמעט הריבוי ותתחזק האחדות עד האחד המיוחד שממנו נמצאו כל האחדים שלכל הנמצאים. והוא אחד מיוחד. שלש אמונות יש: השונים, והאמגושים והנוצרים, והנני מבאר כ"א מהם ומי שיעיין בספריהם יתבאר לו זה. השונים אומרים שיש לעולם שני פועלים: אחד מהם האור, והוא פועל הטוב בנמצאים ולא יוכל לעשות רע, והפועל השני החושך, והוא פועל הרע בעולם ולא יוכל לעשות טוב, והמה חלוקים: יש מהם אומרים שהאור והחושך חיים וי"א שהם מתים. אמנם האמגוש הנם אומרים שיש לעולם פועל חכם גיבור מנהיג ואינו גוף, והרחיקו ממנו מעשות כל רע, וקראוהו "יזדאן". וייחסו כל הרע והצער ההווים בעולם לשטן, וקראוהו "אהרמן" בלשונם. והמה חלוקים:

| id | word | punctuation after |
|---|---|---|
| 5:3.1 | מב | ) |
| 5:3.2 | נטה |  |
| 5:3.3 | לך |  |
| 5:3.4 | קורא |  |
| 5:3.5 | נעים |  |
| 5:3.6 | מהאמונות |  |
| 5:3.7 | המדומות |  |
| 5:3.8 | האלה | . |
| 5:3.9 | וראה |  |
| 5:3.10 | מה |  |
| 5:3.11 | שכ | ' |
| 5:3.12 | אחד |  |
| 5:3.13 | מעמודי |  |
| 5:3.14 | הוראה |  |
| 5:3.15 | אשר |  |
| 5:3.16 | בית |  |
| 5:3.17 | ישראל |  |
| 5:3.18 | נשענים |  |
| 5:3.19 | עליהם |  |
| 5:3.20 | הוא |  |
| 5:3.21 | רבי | ' |
| 5:3.22 | הראש |  |
| 5:3.23 | ז"ל |  |
| 5:3.24 | במס | ' |
| 5:3.25 | חולין |  |
| 5:3.26 | פ | ' |
| 5:3.27 | השוחט |  |
| 5:3.28 | והובאו |  |
| 5:3.29 | דבריו |  |
| 5:3.30 | במבוא |  |
| 5:3.31 | התלמוד |  |
| 5:3.32 | וז"ל |  |
| 5:3.33 | ומסתברא |  |
| 5:3.34 | שכל |  |
| 5:3.35 | תוספתא |  |
| 5:3.36 | שלא |  |
| 5:3.37 | נתפשטה |  |
| 5:3.38 | עד |  |
| 5:3.39 | אחר |  |
| 5:3.40 | סידור |  |
| 5:3.41 | הגמרא |  |
| 5:3.42 | לאו |  |
| 5:3.43 | דסמכא |  |
| 5:3.44 | היא |  |
| 5:3.45 | דמסתמא |  |
| 5:3.46 | כיון |  |
| 5:3.47 | שרצו |  |
| 5:3.48 | חכמי |  |
| 5:3.49 | ישראל |  |
| 5:3.50 | לעשות |  |
| 5:3.51 | חבור |  |
| 5:3.52 | קיימא |  |
| 5:3.53 | אמת | . |
| 5:3.54 | חקרו |  |
| 5:3.55 | ודרשו |  |
| 5:3.56 | וחיזרו |  |
| 5:3.57 | לידע |  |
| 5:3.58 | כל |  |
| 5:3.59 | הספרים |  |
| 5:3.60 | שנכתבו |  |
| 5:3.61 | על |  |
| 5:3.62 | דברי |  |
| 5:3.63 | חכמים |  |
| 5:3.64 | וביררו |  |
| 5:3.65 | אותם |  |
| 5:3.66 | שהם |  |
| 5:3.67 | בר |  |
| 5:3.68 | סמכא |  |
| 5:3.69 | עכ"ל | . |
| 5:3.70 | ועיין |  |
| 5:3.71 | לעיל |  |
| 5:3.72 | סי | ' |
| 5:3.73 | י"ב |  |
| 5:3.74 | וי"ג |  |
| 5:3.75 | כ | ' |
| 5:3.76 | הר | ' |
| 5:3.77 | סעדיה |  |
| 5:3.78 | גאון |  |
| 5:3.79 | שדבריו |  |
| 5:3.80 | מסכימים |  |
| 5:3.81 | לדברי |  |
| 5:3.82 | הרא"ש | . |
| 5:3.83 | ומה |  |
| 5:3.84 | שכ | ' |
| 5:3.85 | המגיה |  |
| 5:3.86 | במבוא |  |
| 5:3.87 | התלמוד |  |
| 5:3.88 | בגיליון |  |
| 5:3.89 | לעיין |  |
| 5:3.90 | בהרא"ש |  |
| 5:3.91 | דהרמב"ם |  |
| 5:3.92 | לא |  |
| 5:3.93 | ס"ל |  |
| 5:3.94 | הכי |  |
| 5:3.95 | כבר |  |
| 5:3.96 | ישב |  |
| 5:3.97 | הר | ' |
| 5:3.98 | ב"י |  |
| 5:3.99 | דברי |  |
| 5:3.100 | הרמב"ם |  |
| 5:3.101 | שפסק |  |
| 5:3.102 | בהתוספתא |  |
| 5:3.103 | ושגם |  |
| 5:3.104 | מסדר |  |
| 5:3.105 | הגמרא |  |
| 5:3.106 | סבירא |  |
| 5:3.107 | ליה |  |
| 5:3.108 | כהתוספתא |  |
| 5:3.109 | ושגירסת |  |
| 5:3.110 | הרמב"ם |  |
| 5:3.111 | בגמרא |  |
| 5:3.112 | שהה |  |
| 5:3.113 | כמיעוט |  |
| 5:3.114 | סימנים |  |
| 5:3.115 | וכבר |  |
| 5:3.116 | הביאה |  |
| 5:3.117 | אבי | ' |
| 5:3.118 | לעיל |  |
| 5:3.119 | שהה |  |
| 5:3.120 | כדי |  |
| 5:3.121 | שיגביה |  |
| 5:3.122 | הבהמה |  |
| 5:3.123 | וירביצנה |  |
| 5:3.124 | וישחוט |  |
| 5:3.125 | מיעוט |  |
| 5:3.126 | הסימנים | . |
| 5:3.127 | וכן |  |
| 5:3.128 | היא |  |
| 5:3.129 | הגירסא |  |
| 5:3.130 | בהרמב"ם |  |
| 5:3.131 | כ"י |  |
| 5:3.132 | ישנים | . |
| 5:3.133 | ונמצא |  |
| 5:3.134 | שגם |  |
| 5:3.135 | הרמב"ם |  |
| 5:3.136 | הכי |  |
| 5:3.137 | ס"ל | , |
| 5:3.138 | וכבר |  |
| 5:3.139 | העיר |  |
| 5:3.140 | על |  |
| 5:3.141 | זה |  |
| 5:3.142 | ג"כ |  |
| 5:3.143 | המ"מ |  |
| 5:3.144 | הנדפס |  |
| 5:3.145 | מחדש | . |
| 5:3.146 | ולפי |  |
| 5:3.147 | זה |  |
| 5:3.148 | כולי |  |
| 5:3.149 | עלמא |  |
| 5:3.150 | מודו |  |
| 5:3.151 | שכל |  |
| 5:3.152 | חיבור |  |
| 5:3.153 | המיוחס |  |
| 5:3.154 | לשום |  |
| 5:3.155 | תנא |  |
| 5:3.156 | או |  |
| 5:3.157 | אמורא |  |
| 5:3.158 | שלא |  |
| 5:3.159 | נתפשט |  |
| 5:3.160 | רק |  |
| 5:3.161 | אחר |  |
| 5:3.162 | סידור |  |
| 5:3.163 | הגמ | ' |
| 5:3.164 | לאו |  |
| 5:3.165 | דסמכא |  |
| 5:3.166 | הוא | . |
| 5:3.167 | וכמו |  |
| 5:3.168 | שאמרו |  |
| 5:3.169 | בירושלמי |  |
| 5:3.170 | שכל |  |
| 5:3.171 | משנה |  |
| 5:3.172 | שלא |  |
| 5:3.173 | נכנסה |  |
| 5:3.174 | לחבורה |  |
| 5:3.175 | אינה |  |
| 5:3.176 | משנה | . |
| 5:3.177 | כמו |  |
| 5:3.178 | שכתבתי |  |
| 5:3.179 | לעיל |  |
| 5:3.180 | סי | ' |
| 5:3.181 | י"ב | . |
| 5:3.182 | ואם |  |
| 5:3.183 | כך |  |
| 5:3.184 | היא |  |
| 5:3.185 | אמונתך |  |
| 5:3.186 | שכל |  |
| 5:3.187 | תפלותינו |  |
| 5:3.188 | ועבודותינו |  |
| 5:3.189 | לסבה |  |
| 5:3.190 | הראשונה |  |
| 5:3.191 | שהוא |  |
| 5:3.192 | ה | ' |
| 5:3.193 | אלהים |  |
| 5:3.194 | אמת |  |
| 5:3.195 | אשר |  |
| 5:3.196 | אין |  |
| 5:3.197 | ראשית |  |
| 5:3.198 | לראשיתו |  |
| 5:3.199 | והוא |  |
| 5:3.200 | שהוציא |  |
| 5:3.201 | את |  |
| 5:3.202 | אבותינו |  |
| 5:3.203 | ממצרים |  |
| 5:3.204 | וכבודו |  |
| 5:3.205 | נגלה |  |
| 5:3.206 | עליהם |  |
| 5:3.207 | בהר |  |
| 5:3.208 | סיני |  |
| 5:3.209 | בתתו |  |
| 5:3.210 | לנו |  |
| 5:3.211 | את |  |
| 5:3.212 | התורה |  |
| 5:3.213 | ואמר |  |
| 5:3.214 | אנכי |  |
| 5:3.215 | ה | ' |
| 5:3.216 | אלהיך | , |
| 5:3.217 | לא |  |
| 5:3.218 | יהיה |  |
| 5:3.219 | לך |  |
| 5:3.220 | אלהים |  |
| 5:3.221 | אחרים |  |
| 5:3.222 | על |  |
| 5:3.223 | פני |  |
| 5:3.224 | וכו | ' |
| 5:3.225 | א"כ |  |
| 5:3.226 | למה |  |
| 5:3.227 | תבזה |  |
| 5:3.228 | את |  |
| 5:3.229 | בעלי |  |
| 5:3.230 | אמונתך |  |
| 5:3.231 | העוסקים |  |
| 5:3.232 | יומם |  |
| 5:3.233 | ולילה |  |
| 5:3.234 | במשנה |  |
| 5:3.235 | ובתלמוד |  |
| 5:3.236 | ותקרא |  |
| 5:3.237 | להם |  |
| 5:3.238 | מינים |  |
| 5:3.239 | ואפיקורסים |  |
| 5:3.240 | ותחזיק |  |
| 5:3.241 | ידי |  |
| 5:3.242 | זרים |  |
| 5:3.243 | לאמונתך |  |
| 5:3.244 | אשר |  |
| 5:3.245 | בילדי |  |
| 5:3.246 | נכרים |  |
| 5:3.247 | ישפיקו |  |
| 5:3.248 | הוא |  |
| 5:3.249 | ספר |  |
| 5:3.250 | הזהר |  |
| 5:3.251 | הנמסר |  |
| 5:3.252 | לישראל |  |
| 5:3.253 | מיד |  |
| 5:3.254 | גוי |  |
| 5:3.255 | ויחסוהו |  |
| 5:3.256 | לרשב"י |  |
| 5:3.257 | התנא |  |
| 5:3.258 | בשקר | . |
| 5:3.259 | והטוב |  |
| 5:3.260 | והנאות |  |
| 5:3.261 | לנו |  |
| 5:3.262 | להחזיק |  |
| 5:3.263 | באמונת |  |
| 5:3.264 | קדמונינו |  |
| 5:3.265 | ביחוד |  |
| 5:3.266 | השי"ת |  |
| 5:3.267 | בלי |  |
| 5:3.268 | צירוף |  |
| 5:3.269 | ושתוף | . |
| 5:3.270 | לקיים |  |
| 5:3.271 | ואל |  |
| 5:3.272 | תטוש |  |
| 5:3.273 | תורת |  |
| 5:3.274 | אמך | , |
| 5:3.275 | ואל |  |
| 5:3.276 | תבוז |  |
| 5:3.277 | כי |  |
| 5:3.278 | זקנה |  |
| 5:3.279 | אמך | . |
| 5:3.280 | רצוני |  |
| 5:3.281 | לומר |  |
| 5:3.282 | שלא |  |
| 5:3.283 | תבוז |  |
| 5:3.284 | ותלעיג |  |
| 5:3.285 | על |  |
| 5:3.286 | אמונת |  |
| 5:3.287 | קדמונינו |  |
| 5:3.288 | ביחוד |  |
| 5:3.289 | השי"ת |  |
| 5:3.290 | ולא |  |
| 5:3.291 | תזלזל |  |
| 5:3.292 | בספריהם |  |
| 5:3.293 | אף |  |
| 5:3.294 | על |  |
| 5:3.295 | פי |  |
| 5:3.296 | שבלו |  |
| 5:3.297 | והזקינו |  |
| 5:3.298 | ונתישנו | . |
| 5:3.299 | לא |  |
| 5:3.300 | כאשר |  |
| 5:3.301 | עשו |  |
| 5:3.302 | בימי |  |
| 5:3.303 | הבינים |  |
| 5:3.304 | שזלזלו |  |
| 5:3.305 | בהם |  |
| 5:3.306 | ואיבדו |  |
| 5:3.307 | אותן |  |
| 5:3.308 | ועשו |  |
| 5:3.309 | אותן |  |
| 5:3.310 | כריכות |  |
| 5:3.311 | לספרים |  |
| 5:3.312 | חדשים |  |
| 5:3.313 | המודפסים | . |
| 5:3.314 | ויש |  |
| 5:3.315 | שמכרו |  |
| 5:3.316 | אותן |  |
| 5:3.317 | בזול |  |
| 5:3.318 | מאד | . |
| 5:3.319 | ובעיני |  |
| 5:3.320 | ראיתי |  |
| 5:3.321 | אחד |  |
| 5:3.322 | מבני |  |
| 5:3.323 | עמנו |  |
| 5:3.324 | שהחליף |  |
| 5:3.325 | את |  |
| 5:3.326 | ספר |  |
| 5:3.327 | מורה |  |
| 5:3.328 | הנבוכים |  |
| 5:3.329 | כ"י |  |
| 5:3.330 | בלשון |  |
| 5:3.331 | ערבי |  |
| 5:3.332 | השוה |  |
| 5:3.333 | לנו |  |
| 5:3.334 | עתה |  |
| 5:3.335 | שנים |  |
| 5:3.336 | ושלשה |  |
| 5:3.337 | זהובים |  |
| 5:3.338 | בספר |  |
| 5:3.339 | קטן |  |
| 5:3.340 | הכמות |  |
| 5:3.341 | והאיכות |  |
| 5:3.342 | מודפס |  |
| 5:3.343 | שוה |  |
| 5:3.344 | חצי |  |
| 5:3.345 | ריאל | . |
| 5:3.346 | והרבה |  |
| 5:3.347 | מהמונינו |  |
| 5:3.348 | מכרו |  |
| 5:3.349 | ספרים |  |
| 5:3.350 | ישנים |  |
| 5:3.351 | שירשו |  |
| 5:3.352 | מאבותיהם |  |
| 5:3.353 | לערלים |  |
| 5:3.354 | שבאו |  |
| 5:3.355 | לארץ |  |
| 5:3.356 | התימן |  |
| 5:3.357 | כמשה |  |
| 5:3.358 | בן |  |
| 5:3.359 | נתנאל |  |
| 5:3.360 | שישב |  |
| 5:3.361 | בבית |  |
| 5:3.362 | מ"ו |  |
| 5:3.363 | אברהם |  |
| 5:3.364 | צאלח | . |
| 5:3.365 | והוא |  |
| 5:3.366 | בעצמו |  |
| 5:3.367 | מכר |  |
| 5:3.368 | לו |  |
| 5:3.369 | אז |  |
| 5:3.370 | נביאים |  |
| 5:3.371 | וכתובים |  |
| 5:3.372 | עם |  |
| 5:3.373 | תרגום |  |
| 5:3.374 | יונתן |  |
| 5:3.375 | בן |  |
| 5:3.376 | עוזיאל |  |
| 5:3.377 | על |  |
| 5:3.378 | נביאים | . |
| 5:3.379 | ותר | ' |
| 5:3.380 | ר | ' |
| 5:3.381 | סעדיה |  |
| 5:3.382 | גאון |  |
| 5:3.383 | ערבי |  |
| 5:3.384 | על |  |
| 5:3.385 | ישעיה |  |
| 5:3.386 | ועל |  |
| 5:3.387 | כתובים | . |
| 5:3.388 | גם |  |
| 5:3.389 | משנה |  |
| 5:3.390 | עם |  |
| 5:3.391 | פי | ' |
| 5:3.392 | הרמב"ם |  |
| 5:3.393 | בערבי |  |
| 5:3.394 | ס | ' |
| 5:3.395 | קדשים |  |
| 5:3.396 | בכתב |  |
| 5:3.397 | מפואר |  |
| 5:3.398 | עם |  |
| 5:3.399 | ציורי |  |
| 5:3.400 | בית |  |
| 5:3.401 | קדשינו |  |
| 5:3.402 | ותפארתינו |  |
| 5:3.403 | וכלי |  |
| 5:3.404 | הקדש |  |
| 5:3.405 | במסכת |  |
| 5:3.406 | מדות | . |
| 5:3.407 | ובעיני |  |
| 5:3.408 | ראיתיה |  |
| 5:3.409 | בידו |  |
| 5:3.410 | אני |  |
| 5:3.411 | והמנוח |  |
| 5:3.412 | מו"ר |  |
| 5:3.413 | יחיא |  |
| 5:3.414 | קרח |  |
| 5:3.415 | זלה"ה | . |
| 5:3.416 | ואין |  |
| 5:3.417 | לאל |  |
| 5:3.418 | ידינו |  |
| 5:3.419 | להוציאה |  |
| 5:3.420 | מתחת |  |
| 5:3.421 | ידו | . |
| 5:3.422 | ועד |  |
| 5:3.423 | עתה |  |
| 5:3.424 | חסר |  |
| 5:3.425 | לנו |  |
| 5:3.426 | סדר |  |
| 5:3.427 | קדשים |  |
| 5:3.428 | ונשים |  |
| 5:3.429 | מהמשנה |  |
| 5:3.430 | כ"י |  |
| 5:3.431 | עם |  |
| 5:3.432 | פי | ' |
| 5:3.433 | הרמב"ם |  |
| 5:3.434 | ז"ל | . |
| 5:3.435 | ולא |  |
| 5:3.436 | אלו |  |
| 5:3.437 | בלבד |  |
| 5:3.438 | מכרו |  |
| 5:3.439 | לו | . |
| 5:3.440 | כי |  |
| 5:3.441 | הרבה |  |
| 5:3.442 | ספרים |  |
| 5:3.443 | כתבי |  |
| 5:3.444 | יד |  |
| 5:3.445 | עתיקים |  |
| 5:3.446 | נמכרו |  |
| 5:3.447 | לו |  |
| 5:3.448 | גם |  |
| 5:3.449 | לזולתו |  |
| 5:3.450 | בזול | . |
| 5:3.451 | כי |  |
| 5:3.452 | היו |  |
| 5:3.453 | בידם |  |
| 5:3.454 | ככלי |  |
| 5:3.455 | אין |  |
| 5:3.456 | חפץ |  |
| 5:3.457 | בו |  |
| 5:3.458 | ולא |  |
| 5:3.459 | היו |  |
| 5:3.460 | יודעים |  |
| 5:3.461 | לקרות |  |
| 5:3.462 | בהם |  |
| 5:3.463 | ואנן |  |
| 5:3.464 | הוא |  |
| 5:3.465 | דעבדינן |  |
| 5:3.466 | לתורת |  |
| 5:3.467 | אבותינו |  |
| 5:3.468 | שלא |  |
| 5:3.469 | תשתכח |  |
| 5:3.470 | מישראל |  |
| 5:3.471 | בעיר |  |
| 5:3.472 | תהלה |  |
| 5:3.473 | צנעא | , |
| 5:3.474 | כי |  |
| 5:3.475 | קבענו |  |
| 5:3.476 | למודינו |  |
| 5:3.477 | במשנה |  |
| 5:3.478 | עם |  |
| 5:3.479 | פי | ' |
| 5:3.480 | הרמב"ם |  |
| 5:3.481 | בערבי | . |
| 5:3.482 | ופירושי |  |
| 5:3.483 | רס"ג |  |
| 5:3.484 | במקרא |  |
| 5:3.485 | עם |  |
| 5:3.486 | תלמוד |  |
| 5:3.487 | ופוסקים | . |
| 5:3.488 | ומקודם |  |
| 5:3.489 | לא |  |
| 5:3.490 | היו |  |
| 5:3.491 | לומדים |  |
| 5:3.492 | בכל |  |
| 5:3.493 | בתי |  |
| 5:3.494 | כנסיות |  |
| 5:3.495 | בלילות |  |
| 5:3.496 | כי |  |
| 5:3.497 | אם |  |
| 5:3.498 | בספר |  |
| 5:3.499 | הזהר |  |
| 5:3.500 | בלי |  |
| 5:3.501 | הבין |  |
| 5:3.502 | כמפורסם | . |
| 5:3.503 | והיו |  |
| 5:3.504 | משתבחים |  |
| 5:3.505 | בעצמם |  |
| 5:3.506 | שלמדו |  |
| 5:3.507 | באותה |  |
| 5:3.508 | לילה |  |
| 5:3.509 | כך |  |
| 5:3.510 | וכך |  |
| 5:3.511 | דפים | . |
| 5:3.512 | ולימוד |  |
| 5:3.513 | המשנה |  |
| 5:3.514 | והתלמוד |  |
| 5:3.515 | היה |  |
| 5:3.516 | יקר |  |
| 5:3.517 | בימים |  |
| 5:3.518 | ההם |  |
| 5:3.519 | אין |  |
| 5:3.520 | חזון |  |
| 5:3.521 | נפרץ | . |
| 5:3.522 | ורק |  |
| 5:3.523 | בבקר |  |
| 5:3.524 | לומדים |  |
| 5:3.525 | מעט |  |
| 5:3.526 | משנה |  |
| 5:3.527 | לפני |  |
| 5:3.528 | מו"ר |  |
| 5:3.529 | יוסף |  |
| 5:3.530 | קרח |  |
| 5:3.531 | ז"ל |  |
| 5:3.532 | עם |  |
| 5:3.533 | פי | ' |
| 5:3.534 | ר | ' |
| 5:3.535 | עובדיה |  |
| 5:3.536 | בבית |  |
| 5:3.537 | הכנסת |  |
| 5:3.538 | שרעבי | . |
| 5:3.539 | ותחלת |  |
| 5:3.540 | לימודנו |  |
| 5:3.541 | היה |  |
| 5:3.542 | שם | , |
| 5:3.543 | וגם |  |
| 5:3.544 | אחרי |  |
| 5:3.545 | כן |  |
| 5:3.546 | לפני |  |
| 5:3.547 | בנו |  |
| 5:3.548 | הרה"ג |  |
| 5:3.549 | כמו"ר |  |
| 5:3.550 | חיים |  |
| 5:3.551 | קרח |  |
| 5:3.552 | ז"ל |  |
| 5:3.553 | בלילה | . |
| 5:3.554 | ורבים |  |
| 5:3.555 | ממשכילי |  |
| 5:3.556 | עמנו |  |
| 5:3.557 | היו |  |
| 5:3.558 | באים |  |
| 5:3.559 | ללמוד | . |
| 5:3.560 | ואחרי |  |
| 5:3.561 | כן |  |
| 5:3.562 | התחלנו |  |
| 5:3.563 | ללמוד |  |
| 5:3.564 | פירוש |  |
| 5:3.565 | הרמב"ם |  |
| 5:3.566 | בערבי |  |
| 5:3.567 | בלילי |  |
| 5:3.568 | שבת |  |
| 5:3.569 | מחצות | . |
| 5:3.570 | ואז |  |
| 5:3.571 | נשתדל |  |
| 5:3.572 | ידינו |  |
| 5:3.573 | המנוח |  |
| 5:3.574 | מו"ר |  |
| 5:3.575 | חיים |  |
| 5:3.576 | אלכסאר |  |
| 5:3.577 | לכתוב |  |
| 5:3.578 | לו |  |
| 5:3.579 | משנה |  |
| 5:3.580 | בפי | ' |
| 5:3.581 | הרמב"ם |  |
| 5:3.582 | ז"ל | . |
| 5:3.583 | ואחרים |  |
| 5:3.584 | מהמשכילים |  |
| 5:3.585 | גם |  |
| 5:3.586 | הם |  |
| 5:3.587 | נשתדלו |  |
| 5:3.588 | לכתבם |  |
| 5:3.589 | וללמוד |  |
| 5:3.590 | בהם | . |
| 5:3.591 | ונמשך |  |
| 5:3.592 | הדבר |  |
| 5:3.593 | שגם |  |
| 5:3.594 | בקצת |  |
| 5:3.595 | בתי |  |
| 5:3.596 | כנסיות |  |
| 5:3.597 | נתקנאו |  |
| 5:3.598 | לעשות |  |
| 5:3.599 | כמעשנו |  |
| 5:3.600 | כמאמר |  |
| 5:3.601 | רז"ל |  |
| 5:3.602 | קנאת |  |
| 5:3.603 | סופרים |  |
| 5:3.604 | תרבה |  |
| 5:3.605 | חכמה | . |
| 5:3.606 | ולמדו |  |
| 5:3.607 | גם |  |
| 5:3.608 | כן |  |
| 5:3.609 | הקדמת |  |
| 5:3.610 | הרמב"ם |  |
| 5:3.611 | לסדר |  |
| 5:3.612 | זרעים |  |
| 5:3.613 | ושלש |  |
| 5:3.614 | עשרה |  |
| 5:3.615 | עיקרים |  |
| 5:3.616 | שכ | ' |
| 5:3.617 | רבינו |  |
| 5:3.618 | בפרק |  |
| 5:3.619 | עשירי |  |
| 5:3.620 | דסנהדרין |  |
| 5:3.621 | בשום |  |
| 5:3.622 | שכל |  |
| 5:3.623 | ויבינו |  |
| 5:3.624 | בהם |  |
| 5:3.625 | ובדברי |  |
| 5:3.626 | ר | ' |
| 5:3.627 | סעדיה |  |
| 5:3.628 | גאון |  |
| 5:3.629 | וחובת |  |
| 5:3.630 | הלבבות |  |
| 5:3.631 | שהם |  |
| 5:3.632 | עומדים |  |
| 5:3.633 | נגד |  |
| 5:3.634 | דברי |  |
| 5:3.635 | הזהר |  |
| 5:3.636 | והמקובלים |  |
| 5:3.637 | החדשים | . |
| 5:3.638 | וראיתי |  |
| 5:3.639 | להביא |  |
| 5:3.640 | פה |  |
| 5:3.641 | העיקר |  |
| 5:3.642 | השני |  |
| 5:3.643 | מהי"ג |  |
| 5:3.644 | עיקרים |  |
| 5:3.645 | עם |  |
| 5:3.646 | קצת |  |
| 5:3.647 | מדברי |  |
| 5:3.648 | אחד |  |
| 5:3.649 | מקדמונינו |  |
| 5:3.650 | הר"ר |  |
| 5:3.651 | חטר |  |
| 5:3.652 | בן |  |
| 5:3.653 | שלמה |  |
| 5:3.654 | בן |  |
| 5:3.655 | אלמעלם |  |
| 5:3.656 | בפי | ' |
| 5:3.657 | העיקר |  |
| 5:3.658 | הזה |  |
| 5:3.659 | וז"ל | . |
| 5:3.660 | ואלקאעדה |  |
| 5:3.661 | אלתאניה |  |
| 5:3.662 | וחדתה |  |
| 5:3.663 | תעלי |  |
| 5:3.664 | ודאלך |  |
| 5:3.665 | אן |  |
| 5:3.666 | הדא |  |
| 5:3.667 | עלה |  |
| 5:3.668 | אלכל |  |
| 5:3.669 | ואחד | , |
| 5:3.670 | ליס |  |
| 5:3.671 | כואחד |  |
| 5:3.672 | אלגנס | , |
| 5:3.673 | ולא |  |
| 5:3.674 | כואחד |  |
| 5:3.675 | אלנוע | . |
| 5:3.676 | ולא |  |
| 5:3.677 | כאלשכץ |  |
| 5:3.678 | אלואחד |  |
| 5:3.679 | אלמרכב |  |
| 5:3.680 | אלדי |  |
| 5:3.681 | ינקסם |  |
| 5:3.682 | לאחאד |  |
| 5:3.683 | כתירה |  |
| 5:3.684 | ולא |  |
| 5:3.685 | ואחד |  |
| 5:3.686 | כאלגסם |  |
| 5:3.687 | אלביסט |  |
| 5:3.688 | אלואחד |  |
| 5:3.689 | באלעדר |  |
| 5:3.690 | אלדי |  |
| 5:3.691 | יקבל |  |
| 5:3.692 | אלאנקסאם |  |
| 5:3.693 | ואלתגזי |  |
| 5:3.694 | אלי |  |
| 5:3.695 | לא |  |
| 5:3.696 | נהאה | . |
| 5:3.697 | בל |  |
| 5:3.698 | הו |  |
| 5:3.699 | תע | ' |
| 5:3.700 | ואחד |  |
| 5:3.701 | בוחדה |  |
| 5:3.702 | ליס |  |
| 5:3.703 | כמתלהא |  |
| 5:3.704 | וחדה |  |
| 5:3.705 | בוגה | , |
| 5:3.706 | והדה |  |
| 5:3.707 | אלקאעדה |  |
| 5:3.708 | אלתאניה |  |
| 5:3.709 | הי |  |
| 5:3.710 | אלמדלול |  |
| 5:3.711 | עליהא |  |
| 5:3.712 | בקולה |  |
| 5:3.713 | שמע |  |
| 5:3.714 | ישראל |  |
| 5:3.715 | ה | ' |
| 5:3.716 | אלהינו |  |
| 5:3.717 | ה | ' |
| 5:3.718 | אחד | . |
| 5:3.719 | שרחהא |  |
| 5:3.720 | אנה |  |
| 5:3.721 | למא |  |
| 5:3.722 | ביין |  |
| 5:3.723 | לנא |  |
| 5:3.724 | ונודה |  |
| 5:3.725 | תע | ' |
| 5:3.726 | אפתקר |  |
| 5:3.727 | לביאן |  |
| 5:3.728 | וחדאנייתה | , |
| 5:3.729 | אד |  |
| 5:3.730 | האתין |  |
| 5:3.731 | אלקאעדתין |  |
| 5:3.732 | המא |  |
| 5:3.733 | אצל |  |
| 5:3.734 | שריעה |  |
| 5:3.735 | משה |  |
| 5:3.736 | רבי |  |
| 5:3.737 | ע"ה | , |
| 5:3.738 | והמא |  |
| 5:3.739 | אוול |  |
| 5:3.740 | כל |  |
| 5:3.741 | שריעה |  |
| 5:3.742 | אלאהייה | . |
| 5:3.743 | ואכר |  |
| 5:3.744 | כל |  |
| 5:3.745 | שריעה |  |
| 5:3.746 | נאמוסייה | , |
| 5:3.747 | לאן |  |
| 5:3.748 | אלשריעה |  |
| 5:3.749 | אלאלאהייה |  |
| 5:3.750 | תבתדי |  |
| 5:3.751 | מן |  |
| 5:3.752 | אלאלאהיאת |  |
| 5:3.753 | ותנתהי |  |
| 5:3.754 | אלי |  |
| 5:3.755 | אלנסטאנייאת | , |
| 5:3.756 | ואלשריעה |  |
| 5:3.757 | אלנמוסייה |  |
| 5:3.758 | עכס |  |
| 5:3.759 | דאלך | . |
| 5:3.760 | ביאן |  |
| 5:3.761 | דאך |  |
| 5:3.762 | אן |  |
| 5:3.763 | אלאנסאן |  |
| 5:3.764 | אלא |  |
| 5:3.765 | עלם |  |
| 5:3.766 | אלונוד |  |
| 5:3.767 | מן |  |
| 5:3.768 | אבתדאיה |  |
| 5:3.769 | אלי |  |
| 5:3.770 | אנתהאיה |  |
| 5:3.771 | אעני |  |
| 5:3.772 | מם |  |
| 5:3.773 | אעלא |  |
| 5:3.774 | אעאליין |  |
| 5:3.775 | אלי |  |
| 5:3.776 | אספל |  |
| 5:3.777 | אלסאפלין | , |
| 5:3.778 | פכיך |  |
| 5:3.779 | מא |  |
| 5:3.780 | נזל |  |
| 5:3.781 | ואכר |  |
| 5:3.782 | פי |  |
| 5:3.783 | אלהבוט |  |
| 5:3.784 | ספלא |  |
| 5:3.785 | קלת |  |
| 5:3.786 | אלוחדה |  |
| 5:3.787 | צרורה |  |
| 5:3.788 | ואזאדת |  |
| 5:3.789 | ואלכהרה |  |
| 5:3.790 | וכו | . |
| 5:3.791 | ואמא |  |
| 5:3.792 | אדא |  |
| 5:3.793 | אכד |  |
| 5:3.794 | אלאנסאן |  |
| 5:3.795 | אן |  |
| 5:3.796 | יתרקא |  |
| 5:3.797 | פי |  |
| 5:3.798 | אלוגוד |  |
| 5:3.799 | מן |  |
| 5:3.800 | אספל |  |
| 5:3.801 | אלי |  |
| 5:3.802 | אלעא | , |
| 5:3.803 | פכיך |  |
| 5:3.804 | מא |  |
| 5:3.805 | אזאד |  |
| 5:3.806 | פי |  |
| 5:3.807 | אלתרקא |  |
| 5:3.808 | עלווא |  |
| 5:3.809 | צעפת |  |
| 5:3.810 | אלכתרה |  |
| 5:3.811 | ותקוות |  |
| 5:3.812 | אלוהדה |  |
| 5:3.813 | אלי |  |
| 5:3.814 | אלוהחד |  |
| 5:3.815 | אלאחד |  |
| 5:3.816 | אלדי |  |
| 5:3.817 | נשת |  |
| 5:3.818 | מנה |  |
| 5:3.819 | כל |  |
| 5:3.820 | וחדה |  |
| 5:3.821 | אחדיהה |  |
| 5:3.822 | פי |  |
| 5:3.823 | אלוגוד | , |
| 5:3.824 | והו |  |
| 5:3.825 | ואחד |  |
| 5:3.826 | אלוהדאנייה | , |
| 5:3.827 | והא |  |
| 5:3.828 | אנא |  |
| 5:3.829 | אדבר |  |
| 5:3.830 | תלאתה |  |
| 5:3.831 | מדאהב | , |
| 5:3.832 | אתנוייה | , |
| 5:3.833 | ואמגוס | , |
| 5:3.834 | ואנצארא |  |
| 5:3.835 | והא |  |
| 5:3.836 | אנא |  |
| 5:3.837 | אביין |  |
| 5:3.838 | כל |  |
| 5:3.839 | מדהב |  |
| 5:3.840 | מנהם |  |
| 5:3.841 | ומן |  |
| 5:3.842 | אטלע |  |
| 5:3.843 | עלי |  |
| 5:3.844 | כתבהם |  |
| 5:3.845 | פצל |  |
| 5:3.846 | דאלך | , |
| 5:3.847 | אמא |  |
| 5:3.848 | אלתנוייה |  |
| 5:3.849 | פאנהם |  |
| 5:3.850 | יקולון |  |
| 5:3.851 | אן |  |
| 5:3.852 | לאלעאלם |  |
| 5:3.853 | צאנעאן |  |
| 5:3.854 | דקימאן | , |
| 5:3.855 | אחדהמא |  |
| 5:3.856 | אלנור | , |
| 5:3.857 | והו |  |
| 5:3.858 | יפעל |  |
| 5:3.859 | אלכיר |  |
| 5:3.860 | פי |  |
| 5:3.861 | אלוגוד |  |
| 5:3.862 | ולא |  |
| 5:3.863 | יקדר |  |
| 5:3.864 | עלי |  |
| 5:3.865 | פעל |  |
| 5:3.866 | אלשר | ! |
| 5:3.867 | ואלתאני |  |
| 5:3.868 | אלצלמה | , |
| 5:3.869 | והו |  |
| 5:3.870 | יפעל |  |
| 5:3.871 | אלשר |  |
| 5:3.872 | פי |  |
| 5:3.873 | אלוגוד |  |
| 5:3.874 | ולא |  |
| 5:3.875 | יקדר |  |
| 5:3.876 | עלי |  |
| 5:3.877 | פעל |  |
| 5:3.878 | אלכיר | , |
| 5:3.879 | והם |  |
| 5:3.880 | מכתלפין | , |
| 5:3.881 | מנהם |  |
| 5:3.882 | מן |  |
| 5:3.883 | יקול |  |
| 5:3.884 | אן |  |
| 5:3.885 | אלנור |  |
| 5:3.886 | ואלצלמה |  |
| 5:3.887 | חיאן | , |
| 5:3.888 | ומנחם |  |
| 5:3.889 | מן |  |
| 5:3.890 | יקול |  |
| 5:3.891 | מייתאן | . |
| 5:3.892 | ואמא |  |
| 5:3.893 | אלמגוס |  |
| 5:3.894 | פאנהם |  |
| 5:3.895 | יקולין |  |
| 5:3.896 | אן |  |
| 5:3.897 | לאלעאלם |  |
| 5:3.898 | צאנעא |  |
| 5:3.899 | חכימא |  |
| 5:3.900 | קדירא |  |
| 5:3.901 | מדברא |  |
| 5:3.902 | גיר |  |
| 5:3.903 | גסם | , |
| 5:3.904 | תזהוה |  |
| 5:3.905 | ען |  |
| 5:3.906 | פעל |  |
| 5:3.907 | אלשדור |  |
| 5:3.908 | וסמוה |  |
| 5:3.909 | יזדאן | , |
| 5:3.910 | תם |  |
| 5:3.911 | אנהם |  |
| 5:3.912 | אצאפו |  |
| 5:3.913 | היה |  |
| 5:3.914 | אלאלאם |  |
| 5:3.915 | אלכיאנה |  |
| 5:3.916 | פי |  |
| 5:3.917 | אלעאלם |  |
| 5:3.918 | לאלשיטאן |  |
| 5:3.919 | אהרמין |  |
| 5:3.920 | בלגתהם | , |
| 5:3.921 | והם |  |
| 5:3.922 | איצא |  |
| 5:3.923 | מכתלפין | , |
| 5:3.924 | מנאהם |  |
| 5:3.925 | מן |  |
| 5:3.926 | יקול |  |
| 5:3.927 | אנה |  |
| 5:3.928 | מהדת |  |
| 5:3.929 | וגסם | , |
| 5:3.930 | ומנאהם |  |
| 5:3.931 | מאן |  |
| 5:3.932 | יקול |  |
| 5:3.933 | בקרם |  |
| 5:3.934 | אהרמין |  |
| 5:3.935 | ואנא |  |
| 5:3.936 | גיר |  |
| 5:3.937 | גסם | , |
| 5:3.938 | תם |  |
| 5:3.939 | אכתלפו |  |
| 5:3.940 | בסבב |  |
| 5:3.941 | חדותה |  |
| 5:3.942 | איצא | , |
| 5:3.943 | פקאלו |  |
| 5:3.944 | אנה |  |
| 5:3.945 | חדת |  |
| 5:3.946 | ליזדאן |  |
| 5:3.947 | פכרה | , |
| 5:3.948 | פקאל |  |
| 5:3.949 | פי |  |
| 5:3.950 | נפסה | , |
| 5:3.951 | פכיף |  |
| 5:3.952 | יכון |  |
| 5:3.953 | חאלי |  |
| 5:3.954 | וחדי |  |
| 5:3.955 | או |  |
| 5:3.956 | מן |  |
| 5:3.957 | ינאזעני | , |
| 5:3.958 | וקאלו |  |
| 5:3.959 | אן |  |
| 5:3.960 | אלשיטאן |  |
| 5:3.961 | תולד |  |
| 5:3.962 | מן |  |
| 5:3.963 | תלך |  |
| 5:3.964 | אלפכרה | , |
| 5:3.965 | תם |  |
| 5:3.966 | אנה |  |
| 5:3.967 | חדאת |  |
| 5:3.968 | יזדאן |  |
| 5:3.969 | וצאלחה |  |
| 5:3.970 | אלי |  |
| 5:3.971 | מרה | , |
| 5:3.972 | ודכרו |  |
| 5:3.973 | אשיא |  |
| 5:3.974 | ליס |  |
| 5:3.975 | הנא |  |
| 5:3.976 | מוצע |  |
| 5:3.977 | דכרהא | , |
| 5:3.978 | ואלכיר |  |
| 5:3.979 | ענדהם |  |
| 5:3.980 | גניע |  |
| 5:3.981 | אעני |  |
| 5:3.982 | אלתנוייה |  |
| 5:3.983 | ואלמגוס |  |
| 5:3.984 | הו |  |
| 5:3.985 | מא |  |
| 5:3.986 | תשתהיה |  |
| 5:3.987 | אלנפוס | , |
| 5:3.988 | ואלשר |  |
| 5:3.989 | הו |  |
| 5:3.990 | מא |  |
| 5:3.991 | תכרהה |  |
| 5:3.992 | אלנפוס | ! |
| 5:3.993 | ואמא |  |
| 5:3.994 | אלנצרא |  |
| 5:3.995 | פאנהם |  |
| 5:3.996 | יקולו |  |
| 5:3.997 | אם |  |
| 5:3.998 | אללה |  |
| 5:3.999 | תע |  |
| 5:3.1000 | ואחד |  |
| 5:3.1001 | עלי |  |
| 5:3.1002 | אלחקיקה | ) |
| 5:3.1003 | תלאתה |  |
| 5:3.1004 | אקאנים |  |
| 5:3.1005 | עלי |  |
| 5:3.1006 | אלחקיקה | , |
| 5:3.1007 | אקנים |  |
| 5:3.1008 | אלאב | , |
| 5:3.1009 | והו |  |
| 5:3.1010 | דאת |  |
| 5:3.1011 | אלבארי |  |
| 5:3.1012 | ענדהם | , |
| 5:3.1013 | ואקנוס |  |
| 5:3.1014 | אלאכן |  |
| 5:3.1015 | והו |  |
| 5:3.1016 | אלעלם |  |
| 5:3.1017 | חכמה | ) |
| 5:3.1018 | ואוקנום |  |
| 5:3.1019 | רוח |  |
| 5:3.1020 | אלקודש |  |
| 5:3.1021 | והו |  |
| 5:3.1022 | אלחיה |  |
| 5:3.1023 | בינה |  |
| 5:3.1024 | הנקרא |  |
| 5:3.1025 | אלהים |  |
| 5:3.1026 | חיים | ) |
| 5:3.1027 | ואלתלאתה |  |
| 5:3.1028 | ורגעו |  |
| 5:3.1029 | ואחד | . |
| 5:3.1030 | והם |  |
| 5:3.1031 | פריק |  |
| 5:3.1032 | מן |  |
| 5:3.1033 | ישראל | . |
| 5:3.1034 | ואסתדלאלהם |  |
| 5:3.1035 | עלי |  |
| 5:3.1036 | אנה |  |
| 5:3.1037 | אלתלאתה |  |
| 5:3.1038 | ירגעו |  |
| 5:3.1039 | ואחד | , |
| 5:3.1040 | מן |  |
| 5:3.1041 | קולח |  |
| 5:3.1042 | תע |  |
| 5:3.1043 | שמע |  |
| 5:3.1044 | ישראל |  |
| 5:3.1045 | ידוד |  |
| 5:3.1046 | אלהינו |  |
| 5:3.1047 | ידוד |  |
| 5:3.1048 | אחד | , |
| 5:3.1049 | קאלו |  |
| 5:3.1050 | הדה |  |
| 5:3.1051 | תלתלאה |  |
| 5:3.1052 | אסמא |  |
| 5:3.1053 | לתלאתה |  |
| 5:3.1054 | מסמיין | , |
| 5:3.1055 | וכלן |  |
| 5:3.1056 | אלתלאתה |  |
| 5:3.1057 | ירגעו |  |
| 5:3.1058 | ואחד | , |
| 5:3.1059 | כמא |  |
| 5:3.1060 | קאל |  |
| 5:3.1061 | תע |  |
| 5:3.1062 | ידוד |  |
| 5:3.1063 | אלהינו |  |
| 5:3.1064 | ידוד |  |
| 5:3.1065 | אחד | , |
| 5:3.1066 | ואלתאלב |  |
| 5:3.1067 | אנהם |  |
| 5:3.1068 | יעתקדו |  |
| 5:3.1069 | פי |  |
| 5:3.1070 | אלתוחיד |  |
| 5:3.1071 | לתהדר |  |
| 5:3.1072 | מן |  |
| 5:3.1073 | דאלך |  |
| 5:3.1074 | ארגע |  |
| 5:3.1075 | אלי |  |
| 5:3.1076 | גרץ |  |
| 5:3.1077 | אלקאעה |  |
| 5:3.1078 | וכו | ' |
| 5:3.1079 | עכ"ל |  |
| 5:3.1080 | הר | ' |
| 5:3.1081 | חטר |  |
| 5:3.1082 | ז"ל |  |
| 5:3.1083 | ע"ש | . |
| 5:3.1084 | תרגום |  |
| 5:3.1085 | הרב |  |
| 5:3.1086 | המחבר | : |
| 5:3.1087 | היסוד |  |
| 5:3.1088 | השני |  |
| 5:3.1089 | יחודו |  |
| 5:3.1090 | יתברך |  |
| 5:3.1091 | וזהו |  |
| 5:3.1092 | לידע |  |
| 5:3.1093 | שזה |  |
| 5:3.1094 | האחד |  |
| 5:3.1095 | שהוא |  |
| 5:3.1096 | סיבת |  |
| 5:3.1097 | הכול | , |
| 5:3.1098 | אחד |  |
| 5:3.1099 | לא |  |
| 5:3.1100 | כאחד |  |
| 5:3.1101 | הסוג | , |
| 5:3.1102 | ולא |  |
| 5:3.1103 | כאחד |  |
| 5:3.1104 | המין | , |
| 5:3.1105 | ולא |  |
| 5:3.1106 | כאחד |  |
| 5:3.1107 | המורכב |  |
| 5:3.1108 | אשר |  |
| 5:3.1109 | יחולק |  |
| 5:3.1110 | לחלקים |  |
| 5:3.1111 | רבים | , |
| 5:3.1112 | ולא |  |
| 5:3.1113 | כגוף |  |
| 5:3.1114 | הפשוט |  |
| 5:3.1115 | שהוא |  |
| 5:3.1116 | אחד |  |
| 5:3.1117 | במנין |  |
| 5:3.1118 | כמו |  |
| 5:3.1119 | המים |  |
| 5:3.1120 | שהוא |  |
| 5:3.1121 | יסוד |  |
| 5:3.1122 | פשוט | ) |
| 5:3.1123 | חסר | ) |
| 5:3.1124 | ומתחלק |  |
| 5:3.1125 | לאין |  |
| 5:3.1126 | קץ | . |
| 5:3.1127 | אבל |  |
| 5:3.1128 | הוא |  |
| 5:3.1129 | יתברך |  |
| 5:3.1130 | אחד | , |
| 5:3.1131 | אחדות |  |
| 5:3.1132 | שאין |  |
| 5:3.1133 | כמוה |  |
| 5:3.1134 | אחדות |  |
| 5:3.1135 | בשום |  |
| 5:3.1136 | פנים | , |
| 5:3.1137 | וזה |  |
| 5:3.1138 | היסוד |  |
| 5:3.1139 | השני |  |
| 5:3.1140 | הוא |  |
| 5:3.1141 | הנלמד |  |
| 5:3.1142 | מאומרו | : |
| 5:3.1143 | שמע |  |
| 5:3.1144 | ישראל |  |
| 5:3.1145 | ה"א |  |
| 5:3.1146 | ה | ' |
| 5:3.1147 | אחד | !". |
| 5:3.1148 | ביאור |  |
| 5:3.1149 | היסוד |  |
| 5:3.1150 | הזה | , |
| 5:3.1151 | הואיל |  |
| 5:3.1152 | וכבר |  |
| 5:3.1153 | באר |  |
| 5:3.1154 | רבנו |  |
| 5:3.1155 | מציאותו |  |
| 5:3.1156 | ית | ', |
| 5:3.1157 | הוצרך |  |
| 5:3.1158 | לבאר |  |
| 5:3.1159 | אחדותו | . |
| 5:3.1160 | כי |  |
| 5:3.1161 | אלו |  |
| 5:3.1162 | השני |  |
| 5:3.1163 | יסודות |  |
| 5:3.1164 | הם |  |
| 5:3.1165 | עיקר |  |
| 5:3.1166 | תורת |  |
| 5:3.1167 | משרע"ה | . |
| 5:3.1168 | והמה |  |
| 5:3.1169 | ראשית |  |
| 5:3.1170 | לכל |  |
| 5:3.1171 | תורה |  |
| 5:3.1172 | אלוהית |  |
| 5:3.1173 | וסוף |  |
| 5:3.1174 | לכל |  |
| 5:3.1175 | תורה |  |
| 5:3.1176 | נימוסית | . |
| 5:3.1177 | כי |  |
| 5:3.1178 | התורה |  |
| 5:3.1179 | האלוהית |  |
| 5:3.1180 | מתחלתאמן |  |
| 5:3.1181 | האלהיות | ) |
| 5:3.1182 | מתחלת |  |
| 5:3.1183 | באלוהיות | ] |
| 5:3.1184 | וגומרת |  |
| 5:3.1185 | בגשמיות |  |
| 5:3.1186 | והתורה |  |
| 5:3.1187 | הנימוסית |  |
| 5:3.1188 | היפך |  |
| 5:3.1189 | זה | . |
| 5:3.1190 | ביאור |  |
| 5:3.1191 | זה |  |
| 5:3.1192 | שהאדם |  |
| 5:3.1193 | אם |  |
| 5:3.1194 | ידע |  |
| 5:3.1195 | מציאות |  |
| 5:3.1196 | כל |  |
| 5:3.1197 | הנמצאים |  |
| 5:3.1198 | מראש |  |
| 5:3.1199 | ועד |  |
| 5:3.1200 | סוף | , |
| 5:3.1201 | ר"ך |  |
| 5:3.1202 | ממעלה |  |
| 5:3.1203 | מעלה |  |
| 5:3.1204 | ועד |  |
| 5:3.1205 | מטה |  |
| 5:3.1206 | מטה | , |
| 5:3.1207 | כל |  |
| 5:3.1208 | עוד |  |
| 5:3.1209 | שירד |  |
| 5:3.1210 | בידיעתו |  |
| 5:3.1211 | מטה |  |
| 5:3.1212 | נתמעטה |  |
| 5:3.1213 | האחדות |  |
| 5:3.1214 | ונתוסף |  |
| 5:3.1215 | הריבוי |  |
| 5:3.1216 | עד |  |
| 5:3.1217 | שמגיע |  |
| 5:3.1218 | לעולם |  |
| 5:3.1219 | השפל |  |
| 5:3.1220 | שבלי |  |
| 5:3.1221 | ספק |  |
| 5:3.1222 | לא |  |
| 5:3.1223 | ימד |  |
| 5:3.1224 | ולא |  |
| 5:3.1225 | יספר |  |
| 5:3.1226 | סוגי |  |
| 5:3.1227 | מיניו |  |
| 5:3.1228 | כ"ש |  |
| 5:3.1229 | אישי |  |
| 5:3.1230 | כל |  |
| 5:3.1231 | מין |  |
| 5:3.1232 | ומין |  |
| 5:3.1233 | וכו | '. |
| 5:3.1234 | אבל |  |
| 5:3.1235 | אם |  |
| 5:3.1236 | בא |  |
| 5:3.1237 | לעלות |  |
| 5:3.1238 | בנמצאים |  |
| 5:3.1239 | ממטה |  |
| 5:3.1240 | למעלה | , |
| 5:3.1241 | הנה |  |
| 5:3.1242 | כל |  |
| 5:3.1243 | מה |  |
| 5:3.1244 | שיוסיף |  |
| 5:3.1245 | לעלות |  |
| 5:3.1246 | יתמעט |  |
| 5:3.1247 | הריבוי |  |
| 5:3.1248 | ותתחזק |  |
| 5:3.1249 | האחדות |  |
| 5:3.1250 | עד |  |
| 5:3.1251 | האחד |  |
| 5:3.1252 | המיוחד |  |
| 5:3.1253 | שממנו |  |
| 5:3.1254 | נמצאו |  |
| 5:3.1255 | כל |  |
| 5:3.1256 | האחדים |  |
| 5:3.1257 | שלכל |  |
| 5:3.1258 | הנמצאים | . |
| 5:3.1259 | והוא |  |
| 5:3.1260 | אחד |  |
| 5:3.1261 | מיוחד | . |
| 5:3.1262 | שלש |  |
| 5:3.1263 | אמונות |  |
| 5:3.1264 | יש | : |
| 5:3.1265 | השונים | , |
| 5:3.1266 | והאמגושים |  |
| 5:3.1267 | והנוצרים | , |
| 5:3.1268 | והנני |  |
| 5:3.1269 | מבאר |  |
| 5:3.1270 | כ"א |  |
| 5:3.1271 | מהם |  |
| 5:3.1272 | ומי |  |
| 5:3.1273 | שיעיין |  |
| 5:3.1274 | בספריהם |  |
| 5:3.1275 | יתבאר |  |
| 5:3.1276 | לו |  |
| 5:3.1277 | זה | . |
| 5:3.1278 | השונים |  |
| 5:3.1279 | אומרים |  |
| 5:3.1280 | שיש |  |
| 5:3.1281 | לעולם |  |
| 5:3.1282 | שני |  |
| 5:3.1283 | פועלים | : |
| 5:3.1284 | אחד |  |
| 5:3.1285 | מהם |  |
| 5:3.1286 | האור | , |
| 5:3.1287 | והוא |  |
| 5:3.1288 | פועל |  |
| 5:3.1289 | הטוב |  |
| 5:3.1290 | בנמצאים |  |
| 5:3.1291 | ולא |  |
| 5:3.1292 | יוכל |  |
| 5:3.1293 | לעשות |  |
| 5:3.1294 | רע | , |
| 5:3.1295 | והפועל |  |
| 5:3.1296 | השני |  |
| 5:3.1297 | החושך | , |
| 5:3.1298 | והוא |  |
| 5:3.1299 | פועל |  |
| 5:3.1300 | הרע |  |
| 5:3.1301 | בעולם |  |
| 5:3.1302 | ולא |  |
| 5:3.1303 | יוכל |  |
| 5:3.1304 | לעשות |  |
| 5:3.1305 | טוב | , |
| 5:3.1306 | והמה |  |
| 5:3.1307 | חלוקים | : |
| 5:3.1308 | יש |  |
| 5:3.1309 | מהם |  |
| 5:3.1310 | אומרים |  |
| 5:3.1311 | שהאור |  |
| 5:3.1312 | והחושך |  |
| 5:3.1313 | חיים |  |
| 5:3.1314 | וי"א |  |
| 5:3.1315 | שהם |  |
| 5:3.1316 | מתים | . |
| 5:3.1317 | אמנם |  |
| 5:3.1318 | האמגוש |  |
| 5:3.1319 | הנם |  |
| 5:3.1320 | אומרים |  |
| 5:3.1321 | שיש |  |
| 5:3.1322 | לעולם |  |
| 5:3.1323 | פועל |  |
| 5:3.1324 | חכם |  |
| 5:3.1325 | גיבור |  |
| 5:3.1326 | מנהיג |  |
| 5:3.1327 | ואינו |  |
| 5:3.1328 | גוף | , |
| 5:3.1329 | והרחיקו |  |
| 5:3.1330 | ממנו |  |
| 5:3.1331 | מעשות |  |
| 5:3.1332 | כל |  |
| 5:3.1333 | רע | , |
| 5:3.1334 | וקראוהו |  |
| 5:3.1335 | יזדאן | ". |
| 5:3.1336 | וייחסו |  |
| 5:3.1337 | כל |  |
| 5:3.1338 | הרע |  |
| 5:3.1339 | והצער |  |
| 5:3.1340 | ההווים |  |
| 5:3.1341 | בעולם |  |
| 5:3.1342 | לשטן | , |
| 5:3.1343 | וקראוהו |  |
| 5:3.1344 | אהרמן | " |
| 5:3.1345 | בלשונם | . |
| 5:3.1346 | והמה |  |
| 5:3.1347 | חלוקים | : |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

