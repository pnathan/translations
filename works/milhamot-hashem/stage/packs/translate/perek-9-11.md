# Translation pack: milhamot-hashem perek-9-11 (Milchamot ha-Shem 9:13–9:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-11`.

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
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
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


## Milchamot ha-Shem 9:13

Hebrew (Wikisource): קלא) ולפי שנשתבשה אצלם אמונת יחוד השי"ת שהיא לפי האמת אחידות שאין כמוה בכל שאר האחדים כמו שאמרו רז"ל במדרש רבה (הובא לעיל סי' פ"ב) לפי' העלימו עין ממנה והיה להם לקבוע דיני יחוד השם בש"ע ונושאי כליו ובספרי האחרונים בסימנים וסעיפים מיוחדים ולהאריך בהם יותר ממה שהאריך בהם רבי' הרמב"ם בהל' יסוה"ת ובמורה וכמה דינים היו מסתעפים ומתחדשים באמונת היחוד לפי שאין אמונת היחוד איך יאמר אלא יאמין ויקבל עליו עול מלכות שמים באמונת אחדות אחת שאין כמוה בשאר כל האחדים שבעולם בלי שום שתוף עלה אחרת מכל הפרצופים הנבראים וגם להאמין אמונה שלימה במשנה ובתלמוד המקובלים בידינו שמשה רבי' קבלם מפי הגבורה ולא ממטטרון ושאין בהם כל רע כלל לא כדברי הזהר שאמר שהם מעורבים מטוב ורע ושנקראו אבן משכית וקליפה ומוץ ותבן וסלע אחרא וכיוצא בזלזולים אלו אלא הכל מפי הגבורה טוב שאין בו רע כלל והוא חיינו ואורך ימינו. הנה זו רעה חולה עליה יהיו כל הדוים איך נעזוב את ה' שהוא אלהים חיים ומלך עולם ונשמע לדברי הספר המוצא מתוך קברים על ידי מלך ממלכי העכו"ם לעבוד אלהים חדשים פרצופים נבראים בשתוף הדבר האלהי המעורב בהם לפי דעתם שהוא נשמתם וכל מצותינו שנעשה לא לשם השי"ת שאמר ונעשה רצונו רק לקשט בהם עם ברתא שהם זעיר ומלכות, וידוע הוא ומפורש בכתוב שעובדי האלילים הקדומים היו מקשטים אליליהם בכסף וזהב ואבנים יקרות, ואלה המקובלים החדשים מקשטים את אליליהם בתפלותיהם ומעשה מצותיהם והם הנותנים להם חיות ומזון בהתפללם אליהם להמשיך להם שפע וחיות מהאלהים העליון בכדי שיוכלו להשפיע לנו משלהם וחושבים שהתפלה והעבודה היא צורך גבוה כמו שכתב הר"מ גבאי דהיינו צורך הפרצופים הנעבדים, ולא צרכי גופינו היפך דברי רשב"י בירושלמי כדגר' ההם פ"ק דשבת מפסיקין לקרית שמע ואין מפסיקין לתפלה. אמר רבי אחא קרית דבר תורה, ותפלה אינה דבר תורה. אמר ר' בא קרית שמע זמנה קבוע ותפלה אין זמנה קבוע, אמר ר' יוסה קרית שמע אינה צריכה כונה, ותפלה צריכה כונה, אמר ר' מנא קשיתה קומי דר' יוסה ואפי' תימא ק"ש אינה צריכה כונה, שלשה פיסוקים הראשונים צריכים כונה מן גו דאינון ציבחן הוא מכוין ר' יוחנן בשם רשב"י כגון אנו שעסוקין בתלמוד תורה אפי לק"ש אין אנו מפסיקין. רבי יוחנן אמר על גרמיה כגון אנו שאין אנו עסוקין בתלמוד תורה אפי' לתפלה אנו מפסיקין. דין כדעתיה ודין כדעתיה, רבי יוחנן כדעתיה דר' יוחנן אמר הלואי מתפללין כל היום למה שאין תפלה מפסדת, רשב"י כדעתיה, דרשב"י אמר אלו הוינא קאים על טורא דסיני בשעתא דאתיהיבת אוריתא לישראל. הוינא מתבע קומי רחמנא דאתברי להדין לבר נשא תרין פומין. חד דיהוי לעי באוריתא. וחד דיתעביד ביה כל תורכוי, חזר ואמר, מה אין חד הוא, לית עלמא יכיל קאים ביה מן דלטוריא דיליה. אלו הויין תרין על אחת כמה וכמה. אמר ר' יוסא קומי ר' ירמיה, אתיא דר' יוחנן כר' חנניה בן עקביה דתני כותבי ספרים תפלי ומזוזות מפסיקין לק"ש ואין מפסיקין לתפלה. ר' חנניה בו עקביה אומר כשם שמפסיקין לק"ש כך מפסיקין לתפלה ולתפלין ולשאר כל מצותיה של תורה. ופריך ולא מודי רשב"י שמפסיקין לעשות סוכה ולעשות לולב, ולית ליה לרשב"י הלמד לעשות. לא הלמד שלא לעשות? שהלמד שלא לעשות נוח לו אלו לא נברא. ואמר ר' יוחנן הלמד שלא לעשות נוח לו אלו נתהפכה שליתו על פניו ולא יצא לעולם? ומשני טעמיה דרשב"י זה שינון וזה שינון. ואין מפסיקין שינון מפני שינון. ופריך והתנינן הקורא מכאן ואילך לא הפסיד כאשדם שהוא קורא בתורה, הא בעונתה חביבה היא יותר מדברי תורה? ומשני אמר ר' יודן רשב"י על ידי שהיה חדיר בדברי תורה, לפיכך אינה חביבה עליו יותר מדברי תורה ע"כ. שמע מינה מדקאמר רשב"י כדעתיה וכו' שהתנא החסיד רשב"י ס"ל שהתפלה היא תביעת צרכי הגוף ואין לפסוק מדברי תורה שהיא עולם כדי להתפלל ולתבוע צרכי הגוף ואין לפסוק מדברי תורה שהיא עולם כדי להתפלל ותבוע צרכיו שהם חיי שעה ולכן אמר אלו הוינא קאים וכו' הוינא מתבע תרי פומי. חד לעסוק בתורה, וחד לתפלה ולשאר עסקי הגוף. כי התפלה גם היא מעסק הגוף. שמבקש בה רפואה ופרנסה ושאר צרכים. וכדגר' בפ' יציאות השבת. רבא אשכחיה לרב המנונא דקא מאריך בצלותא, אמר מניחין חי עולם ועוסקין בחיי שעה, והוא סבר זמן תורה לחוד, וזמן תפלה לחוד. עוד שם רבי ירמיה הזה יתיב קמי ר' זירא וקא גריס. מטא זמן תפלה וקא מסרהיב קרי עליה רבי זירא מסיר אזנו משמוע צורה גם תפלתו תועבה ע"כ. אלמלא דעסק תורה היה חביב עליהם יותר מן התפלה. כי הוא חיינו ואורך ימינו לעולם שכולו ארוך לעולם שכולו טוב. אלא דבתר הכי מסקינן דלא שנו דאין מפסיקין לתפלה אלא כגון רשב"י וחביריו שתורתן אומנותן. אבל כגון אנו מפסיקין בין לק"ש בין לתפלה, וכן פסקו הרי"ף והרמב"ם כר' יוחנן, והוא כדאמרינן בעלמא הרבה עשו כרשב"י ולא עלתה בידם.

| id | word | punctuation after |
|---|---|---|
| 9:13.1 | קלא | ) |
| 9:13.2 | ולפי |  |
| 9:13.3 | שנשתבשה |  |
| 9:13.4 | אצלם |  |
| 9:13.5 | אמונת |  |
| 9:13.6 | יחוד |  |
| 9:13.7 | השי"ת |  |
| 9:13.8 | שהיא |  |
| 9:13.9 | לפי |  |
| 9:13.10 | האמת |  |
| 9:13.11 | אחידות |  |
| 9:13.12 | שאין |  |
| 9:13.13 | כמוה |  |
| 9:13.14 | בכל |  |
| 9:13.15 | שאר |  |
| 9:13.16 | האחדים |  |
| 9:13.17 | כמו |  |
| 9:13.18 | שאמרו |  |
| 9:13.19 | רז"ל |  |
| 9:13.20 | במדרש |  |
| 9:13.21 | רבה |  |
| 9:13.22 | הובא |  |
| 9:13.23 | לעיל |  |
| 9:13.24 | סי | ' |
| 9:13.25 | פ"ב | ) |
| 9:13.26 | לפי | ' |
| 9:13.27 | העלימו |  |
| 9:13.28 | עין |  |
| 9:13.29 | ממנה |  |
| 9:13.30 | והיה |  |
| 9:13.31 | להם |  |
| 9:13.32 | לקבוע |  |
| 9:13.33 | דיני |  |
| 9:13.34 | יחוד |  |
| 9:13.35 | השם |  |
| 9:13.36 | בש"ע |  |
| 9:13.37 | ונושאי |  |
| 9:13.38 | כליו |  |
| 9:13.39 | ובספרי |  |
| 9:13.40 | האחרונים |  |
| 9:13.41 | בסימנים |  |
| 9:13.42 | וסעיפים |  |
| 9:13.43 | מיוחדים |  |
| 9:13.44 | ולהאריך |  |
| 9:13.45 | בהם |  |
| 9:13.46 | יותר |  |
| 9:13.47 | ממה |  |
| 9:13.48 | שהאריך |  |
| 9:13.49 | בהם |  |
| 9:13.50 | רבי | ' |
| 9:13.51 | הרמב"ם |  |
| 9:13.52 | בהל | ' |
| 9:13.53 | יסוה"ת |  |
| 9:13.54 | ובמורה |  |
| 9:13.55 | וכמה |  |
| 9:13.56 | דינים |  |
| 9:13.57 | היו |  |
| 9:13.58 | מסתעפים |  |
| 9:13.59 | ומתחדשים |  |
| 9:13.60 | באמונת |  |
| 9:13.61 | היחוד |  |
| 9:13.62 | לפי |  |
| 9:13.63 | שאין |  |
| 9:13.64 | אמונת |  |
| 9:13.65 | היחוד |  |
| 9:13.66 | איך |  |
| 9:13.67 | יאמר |  |
| 9:13.68 | אלא |  |
| 9:13.69 | יאמין |  |
| 9:13.70 | ויקבל |  |
| 9:13.71 | עליו |  |
| 9:13.72 | עול |  |
| 9:13.73 | מלכות |  |
| 9:13.74 | שמים |  |
| 9:13.75 | באמונת |  |
| 9:13.76 | אחדות |  |
| 9:13.77 | אחת |  |
| 9:13.78 | שאין |  |
| 9:13.79 | כמוה |  |
| 9:13.80 | בשאר |  |
| 9:13.81 | כל |  |
| 9:13.82 | האחדים |  |
| 9:13.83 | שבעולם |  |
| 9:13.84 | בלי |  |
| 9:13.85 | שום |  |
| 9:13.86 | שתוף |  |
| 9:13.87 | עלה |  |
| 9:13.88 | אחרת |  |
| 9:13.89 | מכל |  |
| 9:13.90 | הפרצופים |  |
| 9:13.91 | הנבראים |  |
| 9:13.92 | וגם |  |
| 9:13.93 | להאמין |  |
| 9:13.94 | אמונה |  |
| 9:13.95 | שלימה |  |
| 9:13.96 | במשנה |  |
| 9:13.97 | ובתלמוד |  |
| 9:13.98 | המקובלים |  |
| 9:13.99 | בידינו |  |
| 9:13.100 | שמשה |  |
| 9:13.101 | רבי | ' |
| 9:13.102 | קבלם |  |
| 9:13.103 | מפי |  |
| 9:13.104 | הגבורה |  |
| 9:13.105 | ולא |  |
| 9:13.106 | ממטטרון |  |
| 9:13.107 | ושאין |  |
| 9:13.108 | בהם |  |
| 9:13.109 | כל |  |
| 9:13.110 | רע |  |
| 9:13.111 | כלל |  |
| 9:13.112 | לא |  |
| 9:13.113 | כדברי |  |
| 9:13.114 | הזהר |  |
| 9:13.115 | שאמר |  |
| 9:13.116 | שהם |  |
| 9:13.117 | מעורבים |  |
| 9:13.118 | מטוב |  |
| 9:13.119 | ורע |  |
| 9:13.120 | ושנקראו |  |
| 9:13.121 | אבן |  |
| 9:13.122 | משכית |  |
| 9:13.123 | וקליפה |  |
| 9:13.124 | ומוץ |  |
| 9:13.125 | ותבן |  |
| 9:13.126 | וסלע |  |
| 9:13.127 | אחרא |  |
| 9:13.128 | וכיוצא |  |
| 9:13.129 | בזלזולים |  |
| 9:13.130 | אלו |  |
| 9:13.131 | אלא |  |
| 9:13.132 | הכל |  |
| 9:13.133 | מפי |  |
| 9:13.134 | הגבורה |  |
| 9:13.135 | טוב |  |
| 9:13.136 | שאין |  |
| 9:13.137 | בו |  |
| 9:13.138 | רע |  |
| 9:13.139 | כלל |  |
| 9:13.140 | והוא |  |
| 9:13.141 | חיינו |  |
| 9:13.142 | ואורך |  |
| 9:13.143 | ימינו | . |
| 9:13.144 | הנה |  |
| 9:13.145 | זו |  |
| 9:13.146 | רעה |  |
| 9:13.147 | חולה |  |
| 9:13.148 | עליה |  |
| 9:13.149 | יהיו |  |
| 9:13.150 | כל |  |
| 9:13.151 | הדוים |  |
| 9:13.152 | איך |  |
| 9:13.153 | נעזוב |  |
| 9:13.154 | את |  |
| 9:13.155 | ה | ' |
| 9:13.156 | שהוא |  |
| 9:13.157 | אלהים |  |
| 9:13.158 | חיים |  |
| 9:13.159 | ומלך |  |
| 9:13.160 | עולם |  |
| 9:13.161 | ונשמע |  |
| 9:13.162 | לדברי |  |
| 9:13.163 | הספר |  |
| 9:13.164 | המוצא |  |
| 9:13.165 | מתוך |  |
| 9:13.166 | קברים |  |
| 9:13.167 | על |  |
| 9:13.168 | ידי |  |
| 9:13.169 | מלך |  |
| 9:13.170 | ממלכי |  |
| 9:13.171 | העכו"ם |  |
| 9:13.172 | לעבוד |  |
| 9:13.173 | אלהים |  |
| 9:13.174 | חדשים |  |
| 9:13.175 | פרצופים |  |
| 9:13.176 | נבראים |  |
| 9:13.177 | בשתוף |  |
| 9:13.178 | הדבר |  |
| 9:13.179 | האלהי |  |
| 9:13.180 | המעורב |  |
| 9:13.181 | בהם |  |
| 9:13.182 | לפי |  |
| 9:13.183 | דעתם |  |
| 9:13.184 | שהוא |  |
| 9:13.185 | נשמתם |  |
| 9:13.186 | וכל |  |
| 9:13.187 | מצותינו |  |
| 9:13.188 | שנעשה |  |
| 9:13.189 | לא |  |
| 9:13.190 | לשם |  |
| 9:13.191 | השי"ת |  |
| 9:13.192 | שאמר |  |
| 9:13.193 | ונעשה |  |
| 9:13.194 | רצונו |  |
| 9:13.195 | רק |  |
| 9:13.196 | לקשט |  |
| 9:13.197 | בהם |  |
| 9:13.198 | עם |  |
| 9:13.199 | ברתא |  |
| 9:13.200 | שהם |  |
| 9:13.201 | זעיר |  |
| 9:13.202 | ומלכות | , |
| 9:13.203 | וידוע |  |
| 9:13.204 | הוא |  |
| 9:13.205 | ומפורש |  |
| 9:13.206 | בכתוב |  |
| 9:13.207 | שעובדי |  |
| 9:13.208 | האלילים |  |
| 9:13.209 | הקדומים |  |
| 9:13.210 | היו |  |
| 9:13.211 | מקשטים |  |
| 9:13.212 | אליליהם |  |
| 9:13.213 | בכסף |  |
| 9:13.214 | וזהב |  |
| 9:13.215 | ואבנים |  |
| 9:13.216 | יקרות | , |
| 9:13.217 | ואלה |  |
| 9:13.218 | המקובלים |  |
| 9:13.219 | החדשים |  |
| 9:13.220 | מקשטים |  |
| 9:13.221 | את |  |
| 9:13.222 | אליליהם |  |
| 9:13.223 | בתפלותיהם |  |
| 9:13.224 | ומעשה |  |
| 9:13.225 | מצותיהם |  |
| 9:13.226 | והם |  |
| 9:13.227 | הנותנים |  |
| 9:13.228 | להם |  |
| 9:13.229 | חיות |  |
| 9:13.230 | ומזון |  |
| 9:13.231 | בהתפללם |  |
| 9:13.232 | אליהם |  |
| 9:13.233 | להמשיך |  |
| 9:13.234 | להם |  |
| 9:13.235 | שפע |  |
| 9:13.236 | וחיות |  |
| 9:13.237 | מהאלהים |  |
| 9:13.238 | העליון |  |
| 9:13.239 | בכדי |  |
| 9:13.240 | שיוכלו |  |
| 9:13.241 | להשפיע |  |
| 9:13.242 | לנו |  |
| 9:13.243 | משלהם |  |
| 9:13.244 | וחושבים |  |
| 9:13.245 | שהתפלה |  |
| 9:13.246 | והעבודה |  |
| 9:13.247 | היא |  |
| 9:13.248 | צורך |  |
| 9:13.249 | גבוה |  |
| 9:13.250 | כמו |  |
| 9:13.251 | שכתב |  |
| 9:13.252 | הר"מ |  |
| 9:13.253 | גבאי |  |
| 9:13.254 | דהיינו |  |
| 9:13.255 | צורך |  |
| 9:13.256 | הפרצופים |  |
| 9:13.257 | הנעבדים | , |
| 9:13.258 | ולא |  |
| 9:13.259 | צרכי |  |
| 9:13.260 | גופינו |  |
| 9:13.261 | היפך |  |
| 9:13.262 | דברי |  |
| 9:13.263 | רשב"י |  |
| 9:13.264 | בירושלמי |  |
| 9:13.265 | כדגר | ' |
| 9:13.266 | ההם |  |
| 9:13.267 | פ"ק |  |
| 9:13.268 | דשבת |  |
| 9:13.269 | מפסיקין |  |
| 9:13.270 | לקרית |  |
| 9:13.271 | שמע |  |
| 9:13.272 | ואין |  |
| 9:13.273 | מפסיקין |  |
| 9:13.274 | לתפלה | . |
| 9:13.275 | אמר |  |
| 9:13.276 | רבי |  |
| 9:13.277 | אחא |  |
| 9:13.278 | קרית |  |
| 9:13.279 | דבר |  |
| 9:13.280 | תורה | , |
| 9:13.281 | ותפלה |  |
| 9:13.282 | אינה |  |
| 9:13.283 | דבר |  |
| 9:13.284 | תורה | . |
| 9:13.285 | אמר |  |
| 9:13.286 | ר | ' |
| 9:13.287 | בא |  |
| 9:13.288 | קרית |  |
| 9:13.289 | שמע |  |
| 9:13.290 | זמנה |  |
| 9:13.291 | קבוע |  |
| 9:13.292 | ותפלה |  |
| 9:13.293 | אין |  |
| 9:13.294 | זמנה |  |
| 9:13.295 | קבוע | , |
| 9:13.296 | אמר |  |
| 9:13.297 | ר | ' |
| 9:13.298 | יוסה |  |
| 9:13.299 | קרית |  |
| 9:13.300 | שמע |  |
| 9:13.301 | אינה |  |
| 9:13.302 | צריכה |  |
| 9:13.303 | כונה | , |
| 9:13.304 | ותפלה |  |
| 9:13.305 | צריכה |  |
| 9:13.306 | כונה | , |
| 9:13.307 | אמר |  |
| 9:13.308 | ר | ' |
| 9:13.309 | מנא |  |
| 9:13.310 | קשיתה |  |
| 9:13.311 | קומי |  |
| 9:13.312 | דר | ' |
| 9:13.313 | יוסה |  |
| 9:13.314 | ואפי | ' |
| 9:13.315 | תימא |  |
| 9:13.316 | ק"ש |  |
| 9:13.317 | אינה |  |
| 9:13.318 | צריכה |  |
| 9:13.319 | כונה | , |
| 9:13.320 | שלשה |  |
| 9:13.321 | פיסוקים |  |
| 9:13.322 | הראשונים |  |
| 9:13.323 | צריכים |  |
| 9:13.324 | כונה |  |
| 9:13.325 | מן |  |
| 9:13.326 | גו |  |
| 9:13.327 | דאינון |  |
| 9:13.328 | ציבחן |  |
| 9:13.329 | הוא |  |
| 9:13.330 | מכוין |  |
| 9:13.331 | ר | ' |
| 9:13.332 | יוחנן |  |
| 9:13.333 | בשם |  |
| 9:13.334 | רשב"י |  |
| 9:13.335 | כגון |  |
| 9:13.336 | אנו |  |
| 9:13.337 | שעסוקין |  |
| 9:13.338 | בתלמוד |  |
| 9:13.339 | תורה |  |
| 9:13.340 | אפי |  |
| 9:13.341 | לק"ש |  |
| 9:13.342 | אין |  |
| 9:13.343 | אנו |  |
| 9:13.344 | מפסיקין | . |
| 9:13.345 | רבי |  |
| 9:13.346 | יוחנן |  |
| 9:13.347 | אמר |  |
| 9:13.348 | על |  |
| 9:13.349 | גרמיה |  |
| 9:13.350 | כגון |  |
| 9:13.351 | אנו |  |
| 9:13.352 | שאין |  |
| 9:13.353 | אנו |  |
| 9:13.354 | עסוקין |  |
| 9:13.355 | בתלמוד |  |
| 9:13.356 | תורה |  |
| 9:13.357 | אפי | ' |
| 9:13.358 | לתפלה |  |
| 9:13.359 | אנו |  |
| 9:13.360 | מפסיקין | . |
| 9:13.361 | דין |  |
| 9:13.362 | כדעתיה |  |
| 9:13.363 | ודין |  |
| 9:13.364 | כדעתיה | , |
| 9:13.365 | רבי |  |
| 9:13.366 | יוחנן |  |
| 9:13.367 | כדעתיה |  |
| 9:13.368 | דר | ' |
| 9:13.369 | יוחנן |  |
| 9:13.370 | אמר |  |
| 9:13.371 | הלואי |  |
| 9:13.372 | מתפללין |  |
| 9:13.373 | כל |  |
| 9:13.374 | היום |  |
| 9:13.375 | למה |  |
| 9:13.376 | שאין |  |
| 9:13.377 | תפלה |  |
| 9:13.378 | מפסדת | , |
| 9:13.379 | רשב"י |  |
| 9:13.380 | כדעתיה | , |
| 9:13.381 | דרשב"י |  |
| 9:13.382 | אמר |  |
| 9:13.383 | אלו |  |
| 9:13.384 | הוינא |  |
| 9:13.385 | קאים |  |
| 9:13.386 | על |  |
| 9:13.387 | טורא |  |
| 9:13.388 | דסיני |  |
| 9:13.389 | בשעתא |  |
| 9:13.390 | דאתיהיבת |  |
| 9:13.391 | אוריתא |  |
| 9:13.392 | לישראל | . |
| 9:13.393 | הוינא |  |
| 9:13.394 | מתבע |  |
| 9:13.395 | קומי |  |
| 9:13.396 | רחמנא |  |
| 9:13.397 | דאתברי |  |
| 9:13.398 | להדין |  |
| 9:13.399 | לבר |  |
| 9:13.400 | נשא |  |
| 9:13.401 | תרין |  |
| 9:13.402 | פומין | . |
| 9:13.403 | חד |  |
| 9:13.404 | דיהוי |  |
| 9:13.405 | לעי |  |
| 9:13.406 | באוריתא | . |
| 9:13.407 | וחד |  |
| 9:13.408 | דיתעביד |  |
| 9:13.409 | ביה |  |
| 9:13.410 | כל |  |
| 9:13.411 | תורכוי | , |
| 9:13.412 | חזר |  |
| 9:13.413 | ואמר | , |
| 9:13.414 | מה |  |
| 9:13.415 | אין |  |
| 9:13.416 | חד |  |
| 9:13.417 | הוא | , |
| 9:13.418 | לית |  |
| 9:13.419 | עלמא |  |
| 9:13.420 | יכיל |  |
| 9:13.421 | קאים |  |
| 9:13.422 | ביה |  |
| 9:13.423 | מן |  |
| 9:13.424 | דלטוריא |  |
| 9:13.425 | דיליה | . |
| 9:13.426 | אלו |  |
| 9:13.427 | הויין |  |
| 9:13.428 | תרין |  |
| 9:13.429 | על |  |
| 9:13.430 | אחת |  |
| 9:13.431 | כמה |  |
| 9:13.432 | וכמה | . |
| 9:13.433 | אמר |  |
| 9:13.434 | ר | ' |
| 9:13.435 | יוסא |  |
| 9:13.436 | קומי |  |
| 9:13.437 | ר | ' |
| 9:13.438 | ירמיה | , |
| 9:13.439 | אתיא |  |
| 9:13.440 | דר | ' |
| 9:13.441 | יוחנן |  |
| 9:13.442 | כר | ' |
| 9:13.443 | חנניה |  |
| 9:13.444 | בן |  |
| 9:13.445 | עקביה |  |
| 9:13.446 | דתני |  |
| 9:13.447 | כותבי |  |
| 9:13.448 | ספרים |  |
| 9:13.449 | תפלי |  |
| 9:13.450 | ומזוזות |  |
| 9:13.451 | מפסיקין |  |
| 9:13.452 | לק"ש |  |
| 9:13.453 | ואין |  |
| 9:13.454 | מפסיקין |  |
| 9:13.455 | לתפלה | . |
| 9:13.456 | ר | ' |
| 9:13.457 | חנניה |  |
| 9:13.458 | בו |  |
| 9:13.459 | עקביה |  |
| 9:13.460 | אומר |  |
| 9:13.461 | כשם |  |
| 9:13.462 | שמפסיקין |  |
| 9:13.463 | לק"ש |  |
| 9:13.464 | כך |  |
| 9:13.465 | מפסיקין |  |
| 9:13.466 | לתפלה |  |
| 9:13.467 | ולתפלין |  |
| 9:13.468 | ולשאר |  |
| 9:13.469 | כל |  |
| 9:13.470 | מצותיה |  |
| 9:13.471 | של |  |
| 9:13.472 | תורה | . |
| 9:13.473 | ופריך |  |
| 9:13.474 | ולא |  |
| 9:13.475 | מודי |  |
| 9:13.476 | רשב"י |  |
| 9:13.477 | שמפסיקין |  |
| 9:13.478 | לעשות |  |
| 9:13.479 | סוכה |  |
| 9:13.480 | ולעשות |  |
| 9:13.481 | לולב | , |
| 9:13.482 | ולית |  |
| 9:13.483 | ליה |  |
| 9:13.484 | לרשב"י |  |
| 9:13.485 | הלמד |  |
| 9:13.486 | לעשות | . |
| 9:13.487 | לא |  |
| 9:13.488 | הלמד |  |
| 9:13.489 | שלא |  |
| 9:13.490 | לעשות | ? |
| 9:13.491 | שהלמד |  |
| 9:13.492 | שלא |  |
| 9:13.493 | לעשות |  |
| 9:13.494 | נוח |  |
| 9:13.495 | לו |  |
| 9:13.496 | אלו |  |
| 9:13.497 | לא |  |
| 9:13.498 | נברא | . |
| 9:13.499 | ואמר |  |
| 9:13.500 | ר | ' |
| 9:13.501 | יוחנן |  |
| 9:13.502 | הלמד |  |
| 9:13.503 | שלא |  |
| 9:13.504 | לעשות |  |
| 9:13.505 | נוח |  |
| 9:13.506 | לו |  |
| 9:13.507 | אלו |  |
| 9:13.508 | נתהפכה |  |
| 9:13.509 | שליתו |  |
| 9:13.510 | על |  |
| 9:13.511 | פניו |  |
| 9:13.512 | ולא |  |
| 9:13.513 | יצא |  |
| 9:13.514 | לעולם | ? |
| 9:13.515 | ומשני |  |
| 9:13.516 | טעמיה |  |
| 9:13.517 | דרשב"י |  |
| 9:13.518 | זה |  |
| 9:13.519 | שינון |  |
| 9:13.520 | וזה |  |
| 9:13.521 | שינון | . |
| 9:13.522 | ואין |  |
| 9:13.523 | מפסיקין |  |
| 9:13.524 | שינון |  |
| 9:13.525 | מפני |  |
| 9:13.526 | שינון | . |
| 9:13.527 | ופריך |  |
| 9:13.528 | והתנינן |  |
| 9:13.529 | הקורא |  |
| 9:13.530 | מכאן |  |
| 9:13.531 | ואילך |  |
| 9:13.532 | לא |  |
| 9:13.533 | הפסיד |  |
| 9:13.534 | כאשדם |  |
| 9:13.535 | שהוא |  |
| 9:13.536 | קורא |  |
| 9:13.537 | בתורה | , |
| 9:13.538 | הא |  |
| 9:13.539 | בעונתה |  |
| 9:13.540 | חביבה |  |
| 9:13.541 | היא |  |
| 9:13.542 | יותר |  |
| 9:13.543 | מדברי |  |
| 9:13.544 | תורה | ? |
| 9:13.545 | ומשני |  |
| 9:13.546 | אמר |  |
| 9:13.547 | ר | ' |
| 9:13.548 | יודן |  |
| 9:13.549 | רשב"י |  |
| 9:13.550 | על |  |
| 9:13.551 | ידי |  |
| 9:13.552 | שהיה |  |
| 9:13.553 | חדיר |  |
| 9:13.554 | בדברי |  |
| 9:13.555 | תורה | , |
| 9:13.556 | לפיכך |  |
| 9:13.557 | אינה |  |
| 9:13.558 | חביבה |  |
| 9:13.559 | עליו |  |
| 9:13.560 | יותר |  |
| 9:13.561 | מדברי |  |
| 9:13.562 | תורה |  |
| 9:13.563 | ע"כ | . |
| 9:13.564 | שמע |  |
| 9:13.565 | מינה |  |
| 9:13.566 | מדקאמר |  |
| 9:13.567 | רשב"י |  |
| 9:13.568 | כדעתיה |  |
| 9:13.569 | וכו | ' |
| 9:13.570 | שהתנא |  |
| 9:13.571 | החסיד |  |
| 9:13.572 | רשב"י |  |
| 9:13.573 | ס"ל |  |
| 9:13.574 | שהתפלה |  |
| 9:13.575 | היא |  |
| 9:13.576 | תביעת |  |
| 9:13.577 | צרכי |  |
| 9:13.578 | הגוף |  |
| 9:13.579 | ואין |  |
| 9:13.580 | לפסוק |  |
| 9:13.581 | מדברי |  |
| 9:13.582 | תורה |  |
| 9:13.583 | שהיא |  |
| 9:13.584 | עולם |  |
| 9:13.585 | כדי |  |
| 9:13.586 | להתפלל |  |
| 9:13.587 | ולתבוע |  |
| 9:13.588 | צרכי |  |
| 9:13.589 | הגוף |  |
| 9:13.590 | ואין |  |
| 9:13.591 | לפסוק |  |
| 9:13.592 | מדברי |  |
| 9:13.593 | תורה |  |
| 9:13.594 | שהיא |  |
| 9:13.595 | עולם |  |
| 9:13.596 | כדי |  |
| 9:13.597 | להתפלל |  |
| 9:13.598 | ותבוע |  |
| 9:13.599 | צרכיו |  |
| 9:13.600 | שהם |  |
| 9:13.601 | חיי |  |
| 9:13.602 | שעה |  |
| 9:13.603 | ולכן |  |
| 9:13.604 | אמר |  |
| 9:13.605 | אלו |  |
| 9:13.606 | הוינא |  |
| 9:13.607 | קאים |  |
| 9:13.608 | וכו | ' |
| 9:13.609 | הוינא |  |
| 9:13.610 | מתבע |  |
| 9:13.611 | תרי |  |
| 9:13.612 | פומי | . |
| 9:13.613 | חד |  |
| 9:13.614 | לעסוק |  |
| 9:13.615 | בתורה | , |
| 9:13.616 | וחד |  |
| 9:13.617 | לתפלה |  |
| 9:13.618 | ולשאר |  |
| 9:13.619 | עסקי |  |
| 9:13.620 | הגוף | . |
| 9:13.621 | כי |  |
| 9:13.622 | התפלה |  |
| 9:13.623 | גם |  |
| 9:13.624 | היא |  |
| 9:13.625 | מעסק |  |
| 9:13.626 | הגוף | . |
| 9:13.627 | שמבקש |  |
| 9:13.628 | בה |  |
| 9:13.629 | רפואה |  |
| 9:13.630 | ופרנסה |  |
| 9:13.631 | ושאר |  |
| 9:13.632 | צרכים | . |
| 9:13.633 | וכדגר | ' |
| 9:13.634 | בפ | ' |
| 9:13.635 | יציאות |  |
| 9:13.636 | השבת | . |
| 9:13.637 | רבא |  |
| 9:13.638 | אשכחיה |  |
| 9:13.639 | לרב |  |
| 9:13.640 | המנונא |  |
| 9:13.641 | דקא |  |
| 9:13.642 | מאריך |  |
| 9:13.643 | בצלותא | , |
| 9:13.644 | אמר |  |
| 9:13.645 | מניחין |  |
| 9:13.646 | חי |  |
| 9:13.647 | עולם |  |
| 9:13.648 | ועוסקין |  |
| 9:13.649 | בחיי |  |
| 9:13.650 | שעה | , |
| 9:13.651 | והוא |  |
| 9:13.652 | סבר |  |
| 9:13.653 | זמן |  |
| 9:13.654 | תורה |  |
| 9:13.655 | לחוד | , |
| 9:13.656 | וזמן |  |
| 9:13.657 | תפלה |  |
| 9:13.658 | לחוד | . |
| 9:13.659 | עוד |  |
| 9:13.660 | שם |  |
| 9:13.661 | רבי |  |
| 9:13.662 | ירמיה |  |
| 9:13.663 | הזה |  |
| 9:13.664 | יתיב |  |
| 9:13.665 | קמי |  |
| 9:13.666 | ר | ' |
| 9:13.667 | זירא |  |
| 9:13.668 | וקא |  |
| 9:13.669 | גריס | . |
| 9:13.670 | מטא |  |
| 9:13.671 | זמן |  |
| 9:13.672 | תפלה |  |
| 9:13.673 | וקא |  |
| 9:13.674 | מסרהיב |  |
| 9:13.675 | קרי |  |
| 9:13.676 | עליה |  |
| 9:13.677 | רבי |  |
| 9:13.678 | זירא |  |
| 9:13.679 | מסיר |  |
| 9:13.680 | אזנו |  |
| 9:13.681 | משמוע |  |
| 9:13.682 | צורה |  |
| 9:13.683 | גם |  |
| 9:13.684 | תפלתו |  |
| 9:13.685 | תועבה |  |
| 9:13.686 | ע"כ | . |
| 9:13.687 | אלמלא |  |
| 9:13.688 | דעסק |  |
| 9:13.689 | תורה |  |
| 9:13.690 | היה |  |
| 9:13.691 | חביב |  |
| 9:13.692 | עליהם |  |
| 9:13.693 | יותר |  |
| 9:13.694 | מן |  |
| 9:13.695 | התפלה | . |
| 9:13.696 | כי |  |
| 9:13.697 | הוא |  |
| 9:13.698 | חיינו |  |
| 9:13.699 | ואורך |  |
| 9:13.700 | ימינו |  |
| 9:13.701 | לעולם |  |
| 9:13.702 | שכולו |  |
| 9:13.703 | ארוך |  |
| 9:13.704 | לעולם |  |
| 9:13.705 | שכולו |  |
| 9:13.706 | טוב | . |
| 9:13.707 | אלא |  |
| 9:13.708 | דבתר |  |
| 9:13.709 | הכי |  |
| 9:13.710 | מסקינן |  |
| 9:13.711 | דלא |  |
| 9:13.712 | שנו |  |
| 9:13.713 | דאין |  |
| 9:13.714 | מפסיקין |  |
| 9:13.715 | לתפלה |  |
| 9:13.716 | אלא |  |
| 9:13.717 | כגון |  |
| 9:13.718 | רשב"י |  |
| 9:13.719 | וחביריו |  |
| 9:13.720 | שתורתן |  |
| 9:13.721 | אומנותן | . |
| 9:13.722 | אבל |  |
| 9:13.723 | כגון |  |
| 9:13.724 | אנו |  |
| 9:13.725 | מפסיקין |  |
| 9:13.726 | בין |  |
| 9:13.727 | לק"ש |  |
| 9:13.728 | בין |  |
| 9:13.729 | לתפלה | , |
| 9:13.730 | וכן |  |
| 9:13.731 | פסקו |  |
| 9:13.732 | הרי"ף |  |
| 9:13.733 | והרמב"ם |  |
| 9:13.734 | כר | ' |
| 9:13.735 | יוחנן | , |
| 9:13.736 | והוא |  |
| 9:13.737 | כדאמרינן |  |
| 9:13.738 | בעלמא |  |
| 9:13.739 | הרבה |  |
| 9:13.740 | עשו |  |
| 9:13.741 | כרשב"י |  |
| 9:13.742 | ולא |  |
| 9:13.743 | עלתה |  |
| 9:13.744 | בידם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

