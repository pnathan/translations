# Translation pack: milhamot-hashem perek-1-04 (Milchamot ha-Shem 1:5–1:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-04`.

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

- **מלחמות ה׳** → *The Wars of the LORD* (Num 21:14; Q1); transliterated
  *Milchamot ha-Shem*.
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
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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


## Milchamot ha-Shem 1:5

Hebrew (Wikisource): וקראתי שם הספר מלחמות ה' על יחודו כי הפילסוף מחבר הספר הזהר קלקל עלינו את אמיתת ייחוד השי״ת, ונמשכו אחריו כמה רבנים וגדולים בתורה בעצימת עינים ועליהם אמר שלמה בחכמתו, כי רבים חללים הפילה ועצומים כל הרוגיה כמו שדרשו רז״ל (סנהדרין רי״ף) כי רבים חללים הפילה זה תלמיד שלא הגיע להוראה ומורה (כמו שעינינו רואות כמה מחברים דמעיילי פילא בקופא דמחטא ממיעוט בקיאותם) ועצומים כל הרוגיה זה תלמיד שהגיע להוראה ואינו מורה, (שאינו נותן לב להורות לאחרים כהלכה, ועצומים לשון ועוצם עיניו רש״י ז״ל) וכל שכן אם משדל אחרים לעסוק בספ' החיצון ס' הזהר ואינו נותן לב להבין הסתירה רק הולך בעצימת עין השכל ומכשיל אחרים באמונה כוזבת שעליו נאמר ועצומים כל הרוגיה ועובר על ולפני עור לא תתן מכשול, ושגגת תלמוד עולה זדון, והוא חוטא ומחטיא את הרבים וחטא הרבים תלוי בו.

| id | word | punctuation after |
|---|---|---|
| 1:5.1 | וקראתי |  |
| 1:5.2 | שם |  |
| 1:5.3 | הספר |  |
| 1:5.4 | מלחמות |  |
| 1:5.5 | ה | ' |
| 1:5.6 | על |  |
| 1:5.7 | יחודו |  |
| 1:5.8 | כי |  |
| 1:5.9 | הפילסוף |  |
| 1:5.10 | מחבר |  |
| 1:5.11 | הספר |  |
| 1:5.12 | הזהר |  |
| 1:5.13 | קלקל |  |
| 1:5.14 | עלינו |  |
| 1:5.15 | את |  |
| 1:5.16 | אמיתת |  |
| 1:5.17 | ייחוד |  |
| 1:5.18 | השי״ת | , |
| 1:5.19 | ונמשכו |  |
| 1:5.20 | אחריו |  |
| 1:5.21 | כמה |  |
| 1:5.22 | רבנים |  |
| 1:5.23 | וגדולים |  |
| 1:5.24 | בתורה |  |
| 1:5.25 | בעצימת |  |
| 1:5.26 | עינים |  |
| 1:5.27 | ועליהם |  |
| 1:5.28 | אמר |  |
| 1:5.29 | שלמה |  |
| 1:5.30 | בחכמתו | , |
| 1:5.31 | כי |  |
| 1:5.32 | רבים |  |
| 1:5.33 | חללים |  |
| 1:5.34 | הפילה |  |
| 1:5.35 | ועצומים |  |
| 1:5.36 | כל |  |
| 1:5.37 | הרוגיה |  |
| 1:5.38 | כמו |  |
| 1:5.39 | שדרשו |  |
| 1:5.40 | רז״ל |  |
| 1:5.41 | סנהדרין |  |
| 1:5.42 | רי״ף | ) |
| 1:5.43 | כי |  |
| 1:5.44 | רבים |  |
| 1:5.45 | חללים |  |
| 1:5.46 | הפילה |  |
| 1:5.47 | זה |  |
| 1:5.48 | תלמיד |  |
| 1:5.49 | שלא |  |
| 1:5.50 | הגיע |  |
| 1:5.51 | להוראה |  |
| 1:5.52 | ומורה |  |
| 1:5.53 | כמו |  |
| 1:5.54 | שעינינו |  |
| 1:5.55 | רואות |  |
| 1:5.56 | כמה |  |
| 1:5.57 | מחברים |  |
| 1:5.58 | דמעיילי |  |
| 1:5.59 | פילא |  |
| 1:5.60 | בקופא |  |
| 1:5.61 | דמחטא |  |
| 1:5.62 | ממיעוט |  |
| 1:5.63 | בקיאותם | ) |
| 1:5.64 | ועצומים |  |
| 1:5.65 | כל |  |
| 1:5.66 | הרוגיה |  |
| 1:5.67 | זה |  |
| 1:5.68 | תלמיד |  |
| 1:5.69 | שהגיע |  |
| 1:5.70 | להוראה |  |
| 1:5.71 | ואינו |  |
| 1:5.72 | מורה | , |
| 1:5.73 | שאינו |  |
| 1:5.74 | נותן |  |
| 1:5.75 | לב |  |
| 1:5.76 | להורות |  |
| 1:5.77 | לאחרים |  |
| 1:5.78 | כהלכה | , |
| 1:5.79 | ועצומים |  |
| 1:5.80 | לשון |  |
| 1:5.81 | ועוצם |  |
| 1:5.82 | עיניו |  |
| 1:5.83 | רש״י |  |
| 1:5.84 | ז״ל | ) |
| 1:5.85 | וכל |  |
| 1:5.86 | שכן |  |
| 1:5.87 | אם |  |
| 1:5.88 | משדל |  |
| 1:5.89 | אחרים |  |
| 1:5.90 | לעסוק |  |
| 1:5.91 | בספ | ' |
| 1:5.92 | החיצון |  |
| 1:5.93 | ס | ' |
| 1:5.94 | הזהר |  |
| 1:5.95 | ואינו |  |
| 1:5.96 | נותן |  |
| 1:5.97 | לב |  |
| 1:5.98 | להבין |  |
| 1:5.99 | הסתירה |  |
| 1:5.100 | רק |  |
| 1:5.101 | הולך |  |
| 1:5.102 | בעצימת |  |
| 1:5.103 | עין |  |
| 1:5.104 | השכל |  |
| 1:5.105 | ומכשיל |  |
| 1:5.106 | אחרים |  |
| 1:5.107 | באמונה |  |
| 1:5.108 | כוזבת |  |
| 1:5.109 | שעליו |  |
| 1:5.110 | נאמר |  |
| 1:5.111 | ועצומים |  |
| 1:5.112 | כל |  |
| 1:5.113 | הרוגיה |  |
| 1:5.114 | ועובר |  |
| 1:5.115 | על |  |
| 1:5.116 | ולפני |  |
| 1:5.117 | עור |  |
| 1:5.118 | לא |  |
| 1:5.119 | תתן |  |
| 1:5.120 | מכשול | , |
| 1:5.121 | ושגגת |  |
| 1:5.122 | תלמוד |  |
| 1:5.123 | עולה |  |
| 1:5.124 | זדון | , |
| 1:5.125 | והוא |  |
| 1:5.126 | חוטא |  |
| 1:5.127 | ומחטיא |  |
| 1:5.128 | את |  |
| 1:5.129 | הרבים |  |
| 1:5.130 | וחטא |  |
| 1:5.131 | הרבים |  |
| 1:5.132 | תלוי |  |
| 1:5.133 | בו | . |


## Milchamot ha-Shem 1:6

Hebrew (Wikisource): ועל אמונה החיצונה הלזו של הפילסוף מחבר הזהר דרשו רז״ל (ילקוט פרשת קדושים ומדרש הגדול שם ורמב״ם בס' המצות לאוין י') אל תפנו אל האלילים ואלהי אם אתה פונה אחרים (לכבדם ספורנו) סופך לעשותם אלוהות ע״כ, וכך אירע באמונת הפילסוף מחבר הזהר כי הלומדים הראשונים אשר למדו בו בחשבם שהוא לרשב״י התנא, חשבו לישב דבריו ולקרות את הפרצופים כלים, והחשיבום וכיבדו אותם והבאים אחריהם עשו אותם אלוהות כמחשבת הפילסוף המחבר וקיבלום עליהם באלוהות כמו שיתבאר בפנים וכן דרשו רז״ל בפסוק יין ושכר אל תשת וז״ל ר' חייא פתח פקודי ה' ישרים וכתיב ביין יתהלך במישרים וכתיב משמחי לב וכתיב ויין ישמח לבב אנוש מה היין הזה משמח את הלב כך דברי תורה משמחין את הלב מה היין הזה השותה ממנו יותר מצרכו לסוף מקיא את הכל כך דברי תורה העוסק בהן ומחשב בהן יותר מדאי סוף שהוא יוצא למינות ע״כ (מדרש הגדול פרשת שמיני פסוק יין ושכר) כדבר הזה עשה הפילסוף נביא השקר מחבר הזהר לתור אחרי לבבו לחשב בדברי תורה יותר מדאי בדרכים הרחוקים מן האמת, ורחוקים מאד מן המדות שקבלו חכמים איש מפי איש שהתורה נדרשת בהן והן הי״ג מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי דאגדתא כמו שיראה המעיין הנבון והסית והדיח לחכמים ישרי לב ההולכים לתומם, להאמינו לשקריו שנגלה לו הקב״ה הנקרא אצלו סבא דסבין ומשה רבינו ‏ע״ה ואליהו וגילו לו סתרים של מינות שאסור לחשב בהן אפי' בבית הכסא והבטיח לעוסק בתורתו בלי מבין לשכר רב יותר מגופי תורה כדי להעור עיני חכמים ולסתום דרך העיון בדבריו שלא להבין את נכליו ולהגדיל מעלת העוסקים בתורתו החדשה המלאה אלילים שונים בעצמותם ופעוליהם מבלי להקפיד על תוספת ורבוי באלהינו המיוחד מכל שאר האחדים ולא יתנו לב איך הקפידה תורה על בל תוסיף בין בתפילין בין במזוזה בין בארבעה מינין שבלולב וכן בשאר מצות אע״פ שנותן חמש הפרשיות בעור אחד והוא מחברן להיות אחד, וכן אם מוסיף במזוזה פרשה אחת אע״פ שקלף מחברן לחיות אחד לא אמרינן כלא חד וכן בציצית ובאגודת ארבעה מינין אם הוסיף מין חמישי באגודה פסל ולא אמרינן כלא חד.

| id | word | punctuation after |
|---|---|---|
| 1:6.1 | ועל |  |
| 1:6.2 | אמונה |  |
| 1:6.3 | החיצונה |  |
| 1:6.4 | הלזו |  |
| 1:6.5 | של |  |
| 1:6.6 | הפילסוף |  |
| 1:6.7 | מחבר |  |
| 1:6.8 | הזהר |  |
| 1:6.9 | דרשו |  |
| 1:6.10 | רז״ל |  |
| 1:6.11 | ילקוט |  |
| 1:6.12 | פרשת |  |
| 1:6.13 | קדושים |  |
| 1:6.14 | ומדרש |  |
| 1:6.15 | הגדול |  |
| 1:6.16 | שם |  |
| 1:6.17 | ורמב״ם |  |
| 1:6.18 | בס | ' |
| 1:6.19 | המצות |  |
| 1:6.20 | לאוין |  |
| 1:6.21 | י | ') |
| 1:6.22 | אל |  |
| 1:6.23 | תפנו |  |
| 1:6.24 | אל |  |
| 1:6.25 | האלילים |  |
| 1:6.26 | ואלהי |  |
| 1:6.27 | אם |  |
| 1:6.28 | אתה |  |
| 1:6.29 | פונה |  |
| 1:6.30 | אחרים |  |
| 1:6.31 | לכבדם |  |
| 1:6.32 | ספורנו | ) |
| 1:6.33 | סופך |  |
| 1:6.34 | לעשותם |  |
| 1:6.35 | אלוהות |  |
| 1:6.36 | ע״כ | , |
| 1:6.37 | וכך |  |
| 1:6.38 | אירע |  |
| 1:6.39 | באמונת |  |
| 1:6.40 | הפילסוף |  |
| 1:6.41 | מחבר |  |
| 1:6.42 | הזהר |  |
| 1:6.43 | כי |  |
| 1:6.44 | הלומדים |  |
| 1:6.45 | הראשונים |  |
| 1:6.46 | אשר |  |
| 1:6.47 | למדו |  |
| 1:6.48 | בו |  |
| 1:6.49 | בחשבם |  |
| 1:6.50 | שהוא |  |
| 1:6.51 | לרשב״י |  |
| 1:6.52 | התנא | , |
| 1:6.53 | חשבו |  |
| 1:6.54 | לישב |  |
| 1:6.55 | דבריו |  |
| 1:6.56 | ולקרות |  |
| 1:6.57 | את |  |
| 1:6.58 | הפרצופים |  |
| 1:6.59 | כלים | , |
| 1:6.60 | והחשיבום |  |
| 1:6.61 | וכיבדו |  |
| 1:6.62 | אותם |  |
| 1:6.63 | והבאים |  |
| 1:6.64 | אחריהם |  |
| 1:6.65 | עשו |  |
| 1:6.66 | אותם |  |
| 1:6.67 | אלוהות |  |
| 1:6.68 | כמחשבת |  |
| 1:6.69 | הפילסוף |  |
| 1:6.70 | המחבר |  |
| 1:6.71 | וקיבלום |  |
| 1:6.72 | עליהם |  |
| 1:6.73 | באלוהות |  |
| 1:6.74 | כמו |  |
| 1:6.75 | שיתבאר |  |
| 1:6.76 | בפנים |  |
| 1:6.77 | וכן |  |
| 1:6.78 | דרשו |  |
| 1:6.79 | רז״ל |  |
| 1:6.80 | בפסוק |  |
| 1:6.81 | יין |  |
| 1:6.82 | ושכר |  |
| 1:6.83 | אל |  |
| 1:6.84 | תשת |  |
| 1:6.85 | וז״ל |  |
| 1:6.86 | ר | ' |
| 1:6.87 | חייא |  |
| 1:6.88 | פתח |  |
| 1:6.89 | פקודי |  |
| 1:6.90 | ה | ' |
| 1:6.91 | ישרים |  |
| 1:6.92 | וכתיב |  |
| 1:6.93 | ביין |  |
| 1:6.94 | יתהלך |  |
| 1:6.95 | במישרים |  |
| 1:6.96 | וכתיב |  |
| 1:6.97 | משמחי |  |
| 1:6.98 | לב |  |
| 1:6.99 | וכתיב |  |
| 1:6.100 | ויין |  |
| 1:6.101 | ישמח |  |
| 1:6.102 | לבב |  |
| 1:6.103 | אנוש |  |
| 1:6.104 | מה |  |
| 1:6.105 | היין |  |
| 1:6.106 | הזה |  |
| 1:6.107 | משמח |  |
| 1:6.108 | את |  |
| 1:6.109 | הלב |  |
| 1:6.110 | כך |  |
| 1:6.111 | דברי |  |
| 1:6.112 | תורה |  |
| 1:6.113 | משמחין |  |
| 1:6.114 | את |  |
| 1:6.115 | הלב |  |
| 1:6.116 | מה |  |
| 1:6.117 | היין |  |
| 1:6.118 | הזה |  |
| 1:6.119 | השותה |  |
| 1:6.120 | ממנו |  |
| 1:6.121 | יותר |  |
| 1:6.122 | מצרכו |  |
| 1:6.123 | לסוף |  |
| 1:6.124 | מקיא |  |
| 1:6.125 | את |  |
| 1:6.126 | הכל |  |
| 1:6.127 | כך |  |
| 1:6.128 | דברי |  |
| 1:6.129 | תורה |  |
| 1:6.130 | העוסק |  |
| 1:6.131 | בהן |  |
| 1:6.132 | ומחשב |  |
| 1:6.133 | בהן |  |
| 1:6.134 | יותר |  |
| 1:6.135 | מדאי |  |
| 1:6.136 | סוף |  |
| 1:6.137 | שהוא |  |
| 1:6.138 | יוצא |  |
| 1:6.139 | למינות |  |
| 1:6.140 | ע״כ |  |
| 1:6.141 | מדרש |  |
| 1:6.142 | הגדול |  |
| 1:6.143 | פרשת |  |
| 1:6.144 | שמיני |  |
| 1:6.145 | פסוק |  |
| 1:6.146 | יין |  |
| 1:6.147 | ושכר | ) |
| 1:6.148 | כדבר |  |
| 1:6.149 | הזה |  |
| 1:6.150 | עשה |  |
| 1:6.151 | הפילסוף |  |
| 1:6.152 | נביא |  |
| 1:6.153 | השקר |  |
| 1:6.154 | מחבר |  |
| 1:6.155 | הזהר |  |
| 1:6.156 | לתור |  |
| 1:6.157 | אחרי |  |
| 1:6.158 | לבבו |  |
| 1:6.159 | לחשב |  |
| 1:6.160 | בדברי |  |
| 1:6.161 | תורה |  |
| 1:6.162 | יותר |  |
| 1:6.163 | מדאי |  |
| 1:6.164 | בדרכים |  |
| 1:6.165 | הרחוקים |  |
| 1:6.166 | מן |  |
| 1:6.167 | האמת | , |
| 1:6.168 | ורחוקים |  |
| 1:6.169 | מאד |  |
| 1:6.170 | מן |  |
| 1:6.171 | המדות |  |
| 1:6.172 | שקבלו |  |
| 1:6.173 | חכמים |  |
| 1:6.174 | איש |  |
| 1:6.175 | מפי |  |
| 1:6.176 | איש |  |
| 1:6.177 | שהתורה |  |
| 1:6.178 | נדרשת |  |
| 1:6.179 | בהן |  |
| 1:6.180 | והן |  |
| 1:6.181 | הי״ג |  |
| 1:6.182 | מדות |  |
| 1:6.183 | דר | ' |
| 1:6.184 | ישמעאל |  |
| 1:6.185 | ול״ג |  |
| 1:6.186 | מדות |  |
| 1:6.187 | דר | ' |
| 1:6.188 | יוסי |  |
| 1:6.189 | הגלילי |  |
| 1:6.190 | דאגדתא |  |
| 1:6.191 | כמו |  |
| 1:6.192 | שיראה |  |
| 1:6.193 | המעיין |  |
| 1:6.194 | הנבון |  |
| 1:6.195 | והסית |  |
| 1:6.196 | והדיח |  |
| 1:6.197 | לחכמים |  |
| 1:6.198 | ישרי |  |
| 1:6.199 | לב |  |
| 1:6.200 | ההולכים |  |
| 1:6.201 | לתומם | , |
| 1:6.202 | להאמינו |  |
| 1:6.203 | לשקריו |  |
| 1:6.204 | שנגלה |  |
| 1:6.205 | לו |  |
| 1:6.206 | הקב״ה |  |
| 1:6.207 | הנקרא |  |
| 1:6.208 | אצלו |  |
| 1:6.209 | סבא |  |
| 1:6.210 | דסבין |  |
| 1:6.211 | ומשה |  |
| 1:6.212 | רבינו |  |
| 1:6.213 | ע״ה |  |
| 1:6.214 | ואליהו |  |
| 1:6.215 | וגילו |  |
| 1:6.216 | לו |  |
| 1:6.217 | סתרים |  |
| 1:6.218 | של |  |
| 1:6.219 | מינות |  |
| 1:6.220 | שאסור |  |
| 1:6.221 | לחשב |  |
| 1:6.222 | בהן |  |
| 1:6.223 | אפי | ' |
| 1:6.224 | בבית |  |
| 1:6.225 | הכסא |  |
| 1:6.226 | והבטיח |  |
| 1:6.227 | לעוסק |  |
| 1:6.228 | בתורתו |  |
| 1:6.229 | בלי |  |
| 1:6.230 | מבין |  |
| 1:6.231 | לשכר |  |
| 1:6.232 | רב |  |
| 1:6.233 | יותר |  |
| 1:6.234 | מגופי |  |
| 1:6.235 | תורה |  |
| 1:6.236 | כדי |  |
| 1:6.237 | להעור |  |
| 1:6.238 | עיני |  |
| 1:6.239 | חכמים |  |
| 1:6.240 | ולסתום |  |
| 1:6.241 | דרך |  |
| 1:6.242 | העיון |  |
| 1:6.243 | בדבריו |  |
| 1:6.244 | שלא |  |
| 1:6.245 | להבין |  |
| 1:6.246 | את |  |
| 1:6.247 | נכליו |  |
| 1:6.248 | ולהגדיל |  |
| 1:6.249 | מעלת |  |
| 1:6.250 | העוסקים |  |
| 1:6.251 | בתורתו |  |
| 1:6.252 | החדשה |  |
| 1:6.253 | המלאה |  |
| 1:6.254 | אלילים |  |
| 1:6.255 | שונים |  |
| 1:6.256 | בעצמותם |  |
| 1:6.257 | ופעוליהם |  |
| 1:6.258 | מבלי |  |
| 1:6.259 | להקפיד |  |
| 1:6.260 | על |  |
| 1:6.261 | תוספת |  |
| 1:6.262 | ורבוי |  |
| 1:6.263 | באלהינו |  |
| 1:6.264 | המיוחד |  |
| 1:6.265 | מכל |  |
| 1:6.266 | שאר |  |
| 1:6.267 | האחדים |  |
| 1:6.268 | ולא |  |
| 1:6.269 | יתנו |  |
| 1:6.270 | לב |  |
| 1:6.271 | איך |  |
| 1:6.272 | הקפידה |  |
| 1:6.273 | תורה |  |
| 1:6.274 | על |  |
| 1:6.275 | בל |  |
| 1:6.276 | תוסיף |  |
| 1:6.277 | בין |  |
| 1:6.278 | בתפילין |  |
| 1:6.279 | בין |  |
| 1:6.280 | במזוזה |  |
| 1:6.281 | בין |  |
| 1:6.282 | בארבעה |  |
| 1:6.283 | מינין |  |
| 1:6.284 | שבלולב |  |
| 1:6.285 | וכן |  |
| 1:6.286 | בשאר |  |
| 1:6.287 | מצות |  |
| 1:6.288 | אע״פ |  |
| 1:6.289 | שנותן |  |
| 1:6.290 | חמש |  |
| 1:6.291 | הפרשיות |  |
| 1:6.292 | בעור |  |
| 1:6.293 | אחד |  |
| 1:6.294 | והוא |  |
| 1:6.295 | מחברן |  |
| 1:6.296 | להיות |  |
| 1:6.297 | אחד | , |
| 1:6.298 | וכן |  |
| 1:6.299 | אם |  |
| 1:6.300 | מוסיף |  |
| 1:6.301 | במזוזה |  |
| 1:6.302 | פרשה |  |
| 1:6.303 | אחת |  |
| 1:6.304 | אע״פ |  |
| 1:6.305 | שקלף |  |
| 1:6.306 | מחברן |  |
| 1:6.307 | לחיות |  |
| 1:6.308 | אחד |  |
| 1:6.309 | לא |  |
| 1:6.310 | אמרינן |  |
| 1:6.311 | כלא |  |
| 1:6.312 | חד |  |
| 1:6.313 | וכן |  |
| 1:6.314 | בציצית |  |
| 1:6.315 | ובאגודת |  |
| 1:6.316 | ארבעה |  |
| 1:6.317 | מינין |  |
| 1:6.318 | אם |  |
| 1:6.319 | הוסיף |  |
| 1:6.320 | מין |  |
| 1:6.321 | חמישי |  |
| 1:6.322 | באגודה |  |
| 1:6.323 | פסל |  |
| 1:6.324 | ולא |  |
| 1:6.325 | אמרינן |  |
| 1:6.326 | כלא |  |
| 1:6.327 | חד | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

