# Translation pack: milhamot-hashem perek-9-18 (Milchamot ha-Shem 9:21–9:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-18`.

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
- **מורא** → "fear".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **מעמד** → "assembly" or "occasion"; "the Standing at Sinai" where he names it.
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 9:21

Hebrew (Wikisource): קלט) לפיכך כל מי שלבו שלם עם ה' אלהינו שהוא הסבה הראשונה לכל הנמצאים בלי שום שתוף כח אחר, אל ישמע לדברי המסית מחבר הזהר והנמשכים אחריו! כי ה' אלהינו הוא לבדו הנמציא כל הנמצאים ברוב טובו וחסדו, והוא היחיד ומיוחד מכל שאר האחדים, והוא לבדו ראשון שאין ראשית לראשיתו, והוא שהוציאנו ממצרים, והוא שנגלה בסיני ונתן לנו את התורה בקולות וברקים ולפידים ורול שופר, ורעשו כל ההרים ונזדעזע כל העולם לקול נבורתו, ונתן בקולו קול עוז אנכי ה' אלהיך אשר הוציאתיך, לא היה לך אלהים אחרים על פני, וכל עשרה דברים הנאמרים במעמד הנבחר ההוא בפרסום גדול, עד שכל מלכי מזרח ומערב חרדו ותמהו, והלכו אצל בלעם נביאם להודיעם מה זה, כמו שאמרו רז"ל במדרש, ואבותינו כולם ששים רבוא איש רגלי מלבד הזקינים והנשים והטף כולם נצבים בתחתית הההר הר סיני לקבל תורה ומצות, וענו כולם ואמרו כל אשר דבר ה' נעשה ונשמע, ומיראה ופחד גדול מקול ה' בכח, אמרו למשה רבינו דבר אתה עמנו ונשמעה וכו'. וכל זה עשה הקב"ה עמנו כדי להאמין במשה בחירו שתורת אמת נתן לנו, ולא יחליפה ולא יחליפה ולא ימירה וגם אמר את ה' האמרת וכו' וה' האמירך היום להיות לו לעם סגלה, ולא נאמין לשום נביא או קוסם שיפתה אותנו ל עבוד אלוה אחר מבלעדיו, רק אותו לבדו נעבוד שלי שום שתוף דבר אחר עמו, והוא הודיע דרכיו למשה. כי הוא חנון ורחום ארף אפים וכו' ולא קצר אפים ואם כן איך נעזוב את אלהינו ארך אפים ונעבוד לאלוה קצא אפים? וגם אם יעשה נביא השקר אותות ומופתים לעינינו בשמים ובארץ. כמו שנאמר ונתן אליך אות או מופת וכי' ובא האות והמופת אשר דבר אליך לאמר נלכה אחרי אלהים אחרים אשר לא ידעתם (לא בנביאנו ולא מרז"ל) ונעבדם. לא תשמע אל דברי הנביא ההוא וכו' כל שכן וק"ו שלא נשמע לדברי הפילסוף מחבר הזהר, אשר זה מקרוב צמח נן הארץ ולא נודע לו בית אב מי הוא זה מחברו באמת, כשאר ספרי רז"ל הידועים ומפורסמים מחבריהם מיום שחוברו על ידי גדולי ישראל מעתיקי השמועה ועד עתה. וגדול ועצום הוא הק"ו הזה. ומה אם נביא העושה אותות ומופתים לענינינו בשמים ובארץ כמו שפירשו רז"ל ומביאו רש"י ז"ל, אמרה תורה לא תשמע אל דברי הנביא ההוא כי מנסה ה' אלהיכם אתכם? ספר הזהר הבא לנו מיד עכו"ם, ומשמיע גדולות ונוראות משם תנאים ואמוראים ונשמות צדיקים שקדמו לרשב"י כמה דורות. ואמוראים המאוחרים מן רב מאד אחר רשב"י, ואליהו ועתיקא קדישא שנגלה לנו בבית מדרשו (בין חביריו ותלמידיו) וכיוצא בספורים אלו הרשומים בדיו שחור על נייר לבן, אשר לא ראינו בעינינו ולא אבותינו ספרו לנו כל אותן הדברים והמעשיות הרשומים שם! כ"ש שלא נשמע בליו לעבוד שום פרצוף מכל הפרצופים, ובפרט קצר אפים (ז"א) כי אם את אלהינו בהו"ש אשר קבלנו מאבותינו, ותוה"ק ספרה לנו שאין לו דמות הגוף, ואין ראשית לראשיתו, והוא אחד ואין יחוד כיחודו לפי שכל הפרצופים והעגולים כולם יש ראשית לראשיתם, ואין לעבוד רק לצור הכל בלי שום שתוף נברא או נאצל מכל הנבראים והנאצלים, וגם אין ליחד ולקשר עם אחדותו שום אחד מהם, כמאמר רשב"י כל המשתף וכו'.

| id | word | punctuation after |
|---|---|---|
| 9:21.1 | קלט | ) |
| 9:21.2 | לפיכך |  |
| 9:21.3 | כל |  |
| 9:21.4 | מי |  |
| 9:21.5 | שלבו |  |
| 9:21.6 | שלם |  |
| 9:21.7 | עם |  |
| 9:21.8 | ה | ' |
| 9:21.9 | אלהינו |  |
| 9:21.10 | שהוא |  |
| 9:21.11 | הסבה |  |
| 9:21.12 | הראשונה |  |
| 9:21.13 | לכל |  |
| 9:21.14 | הנמצאים |  |
| 9:21.15 | בלי |  |
| 9:21.16 | שום |  |
| 9:21.17 | שתוף |  |
| 9:21.18 | כח |  |
| 9:21.19 | אחר | , |
| 9:21.20 | אל |  |
| 9:21.21 | ישמע |  |
| 9:21.22 | לדברי |  |
| 9:21.23 | המסית |  |
| 9:21.24 | מחבר |  |
| 9:21.25 | הזהר |  |
| 9:21.26 | והנמשכים |  |
| 9:21.27 | אחריו | ! |
| 9:21.28 | כי |  |
| 9:21.29 | ה | ' |
| 9:21.30 | אלהינו |  |
| 9:21.31 | הוא |  |
| 9:21.32 | לבדו |  |
| 9:21.33 | הנמציא |  |
| 9:21.34 | כל |  |
| 9:21.35 | הנמצאים |  |
| 9:21.36 | ברוב |  |
| 9:21.37 | טובו |  |
| 9:21.38 | וחסדו | , |
| 9:21.39 | והוא |  |
| 9:21.40 | היחיד |  |
| 9:21.41 | ומיוחד |  |
| 9:21.42 | מכל |  |
| 9:21.43 | שאר |  |
| 9:21.44 | האחדים | , |
| 9:21.45 | והוא |  |
| 9:21.46 | לבדו |  |
| 9:21.47 | ראשון |  |
| 9:21.48 | שאין |  |
| 9:21.49 | ראשית |  |
| 9:21.50 | לראשיתו | , |
| 9:21.51 | והוא |  |
| 9:21.52 | שהוציאנו |  |
| 9:21.53 | ממצרים | , |
| 9:21.54 | והוא |  |
| 9:21.55 | שנגלה |  |
| 9:21.56 | בסיני |  |
| 9:21.57 | ונתן |  |
| 9:21.58 | לנו |  |
| 9:21.59 | את |  |
| 9:21.60 | התורה |  |
| 9:21.61 | בקולות |  |
| 9:21.62 | וברקים |  |
| 9:21.63 | ולפידים |  |
| 9:21.64 | ורול |  |
| 9:21.65 | שופר | , |
| 9:21.66 | ורעשו |  |
| 9:21.67 | כל |  |
| 9:21.68 | ההרים |  |
| 9:21.69 | ונזדעזע |  |
| 9:21.70 | כל |  |
| 9:21.71 | העולם |  |
| 9:21.72 | לקול |  |
| 9:21.73 | נבורתו | , |
| 9:21.74 | ונתן |  |
| 9:21.75 | בקולו |  |
| 9:21.76 | קול |  |
| 9:21.77 | עוז |  |
| 9:21.78 | אנכי |  |
| 9:21.79 | ה | ' |
| 9:21.80 | אלהיך |  |
| 9:21.81 | אשר |  |
| 9:21.82 | הוציאתיך | , |
| 9:21.83 | לא |  |
| 9:21.84 | היה |  |
| 9:21.85 | לך |  |
| 9:21.86 | אלהים |  |
| 9:21.87 | אחרים |  |
| 9:21.88 | על |  |
| 9:21.89 | פני | , |
| 9:21.90 | וכל |  |
| 9:21.91 | עשרה |  |
| 9:21.92 | דברים |  |
| 9:21.93 | הנאמרים |  |
| 9:21.94 | במעמד |  |
| 9:21.95 | הנבחר |  |
| 9:21.96 | ההוא |  |
| 9:21.97 | בפרסום |  |
| 9:21.98 | גדול | , |
| 9:21.99 | עד |  |
| 9:21.100 | שכל |  |
| 9:21.101 | מלכי |  |
| 9:21.102 | מזרח |  |
| 9:21.103 | ומערב |  |
| 9:21.104 | חרדו |  |
| 9:21.105 | ותמהו | , |
| 9:21.106 | והלכו |  |
| 9:21.107 | אצל |  |
| 9:21.108 | בלעם |  |
| 9:21.109 | נביאם |  |
| 9:21.110 | להודיעם |  |
| 9:21.111 | מה |  |
| 9:21.112 | זה | , |
| 9:21.113 | כמו |  |
| 9:21.114 | שאמרו |  |
| 9:21.115 | רז"ל |  |
| 9:21.116 | במדרש | , |
| 9:21.117 | ואבותינו |  |
| 9:21.118 | כולם |  |
| 9:21.119 | ששים |  |
| 9:21.120 | רבוא |  |
| 9:21.121 | איש |  |
| 9:21.122 | רגלי |  |
| 9:21.123 | מלבד |  |
| 9:21.124 | הזקינים |  |
| 9:21.125 | והנשים |  |
| 9:21.126 | והטף |  |
| 9:21.127 | כולם |  |
| 9:21.128 | נצבים |  |
| 9:21.129 | בתחתית |  |
| 9:21.130 | הההר |  |
| 9:21.131 | הר |  |
| 9:21.132 | סיני |  |
| 9:21.133 | לקבל |  |
| 9:21.134 | תורה |  |
| 9:21.135 | ומצות | , |
| 9:21.136 | וענו |  |
| 9:21.137 | כולם |  |
| 9:21.138 | ואמרו |  |
| 9:21.139 | כל |  |
| 9:21.140 | אשר |  |
| 9:21.141 | דבר |  |
| 9:21.142 | ה | ' |
| 9:21.143 | נעשה |  |
| 9:21.144 | ונשמע | , |
| 9:21.145 | ומיראה |  |
| 9:21.146 | ופחד |  |
| 9:21.147 | גדול |  |
| 9:21.148 | מקול |  |
| 9:21.149 | ה | ' |
| 9:21.150 | בכח | , |
| 9:21.151 | אמרו |  |
| 9:21.152 | למשה |  |
| 9:21.153 | רבינו |  |
| 9:21.154 | דבר |  |
| 9:21.155 | אתה |  |
| 9:21.156 | עמנו |  |
| 9:21.157 | ונשמעה |  |
| 9:21.158 | וכו | '. |
| 9:21.159 | וכל |  |
| 9:21.160 | זה |  |
| 9:21.161 | עשה |  |
| 9:21.162 | הקב"ה |  |
| 9:21.163 | עמנו |  |
| 9:21.164 | כדי |  |
| 9:21.165 | להאמין |  |
| 9:21.166 | במשה |  |
| 9:21.167 | בחירו |  |
| 9:21.168 | שתורת |  |
| 9:21.169 | אמת |  |
| 9:21.170 | נתן |  |
| 9:21.171 | לנו | , |
| 9:21.172 | ולא |  |
| 9:21.173 | יחליפה |  |
| 9:21.174 | ולא |  |
| 9:21.175 | יחליפה |  |
| 9:21.176 | ולא |  |
| 9:21.177 | ימירה |  |
| 9:21.178 | וגם |  |
| 9:21.179 | אמר |  |
| 9:21.180 | את |  |
| 9:21.181 | ה | ' |
| 9:21.182 | האמרת |  |
| 9:21.183 | וכו | ' |
| 9:21.184 | וה | ' |
| 9:21.185 | האמירך |  |
| 9:21.186 | היום |  |
| 9:21.187 | להיות |  |
| 9:21.188 | לו |  |
| 9:21.189 | לעם |  |
| 9:21.190 | סגלה | , |
| 9:21.191 | ולא |  |
| 9:21.192 | נאמין |  |
| 9:21.193 | לשום |  |
| 9:21.194 | נביא |  |
| 9:21.195 | או |  |
| 9:21.196 | קוסם |  |
| 9:21.197 | שיפתה |  |
| 9:21.198 | אותנו |  |
| 9:21.199 | ל |  |
| 9:21.200 | עבוד |  |
| 9:21.201 | אלוה |  |
| 9:21.202 | אחר |  |
| 9:21.203 | מבלעדיו | , |
| 9:21.204 | רק |  |
| 9:21.205 | אותו |  |
| 9:21.206 | לבדו |  |
| 9:21.207 | נעבוד |  |
| 9:21.208 | שלי |  |
| 9:21.209 | שום |  |
| 9:21.210 | שתוף |  |
| 9:21.211 | דבר |  |
| 9:21.212 | אחר |  |
| 9:21.213 | עמו | , |
| 9:21.214 | והוא |  |
| 9:21.215 | הודיע |  |
| 9:21.216 | דרכיו |  |
| 9:21.217 | למשה | . |
| 9:21.218 | כי |  |
| 9:21.219 | הוא |  |
| 9:21.220 | חנון |  |
| 9:21.221 | ורחום |  |
| 9:21.222 | ארף |  |
| 9:21.223 | אפים |  |
| 9:21.224 | וכו | ' |
| 9:21.225 | ולא |  |
| 9:21.226 | קצר |  |
| 9:21.227 | אפים |  |
| 9:21.228 | ואם |  |
| 9:21.229 | כן |  |
| 9:21.230 | איך |  |
| 9:21.231 | נעזוב |  |
| 9:21.232 | את |  |
| 9:21.233 | אלהינו |  |
| 9:21.234 | ארך |  |
| 9:21.235 | אפים |  |
| 9:21.236 | ונעבוד |  |
| 9:21.237 | לאלוה |  |
| 9:21.238 | קצא |  |
| 9:21.239 | אפים | ? |
| 9:21.240 | וגם |  |
| 9:21.241 | אם |  |
| 9:21.242 | יעשה |  |
| 9:21.243 | נביא |  |
| 9:21.244 | השקר |  |
| 9:21.245 | אותות |  |
| 9:21.246 | ומופתים |  |
| 9:21.247 | לעינינו |  |
| 9:21.248 | בשמים |  |
| 9:21.249 | ובארץ | . |
| 9:21.250 | כמו |  |
| 9:21.251 | שנאמר |  |
| 9:21.252 | ונתן |  |
| 9:21.253 | אליך |  |
| 9:21.254 | אות |  |
| 9:21.255 | או |  |
| 9:21.256 | מופת |  |
| 9:21.257 | וכי | ' |
| 9:21.258 | ובא |  |
| 9:21.259 | האות |  |
| 9:21.260 | והמופת |  |
| 9:21.261 | אשר |  |
| 9:21.262 | דבר |  |
| 9:21.263 | אליך |  |
| 9:21.264 | לאמר |  |
| 9:21.265 | נלכה |  |
| 9:21.266 | אחרי |  |
| 9:21.267 | אלהים |  |
| 9:21.268 | אחרים |  |
| 9:21.269 | אשר |  |
| 9:21.270 | לא |  |
| 9:21.271 | ידעתם |  |
| 9:21.272 | לא |  |
| 9:21.273 | בנביאנו |  |
| 9:21.274 | ולא |  |
| 9:21.275 | מרז"ל | ) |
| 9:21.276 | ונעבדם | . |
| 9:21.277 | לא |  |
| 9:21.278 | תשמע |  |
| 9:21.279 | אל |  |
| 9:21.280 | דברי |  |
| 9:21.281 | הנביא |  |
| 9:21.282 | ההוא |  |
| 9:21.283 | וכו | ' |
| 9:21.284 | כל |  |
| 9:21.285 | שכן |  |
| 9:21.286 | וק"ו |  |
| 9:21.287 | שלא |  |
| 9:21.288 | נשמע |  |
| 9:21.289 | לדברי |  |
| 9:21.290 | הפילסוף |  |
| 9:21.291 | מחבר |  |
| 9:21.292 | הזהר | , |
| 9:21.293 | אשר |  |
| 9:21.294 | זה |  |
| 9:21.295 | מקרוב |  |
| 9:21.296 | צמח |  |
| 9:21.297 | נן |  |
| 9:21.298 | הארץ |  |
| 9:21.299 | ולא |  |
| 9:21.300 | נודע |  |
| 9:21.301 | לו |  |
| 9:21.302 | בית |  |
| 9:21.303 | אב |  |
| 9:21.304 | מי |  |
| 9:21.305 | הוא |  |
| 9:21.306 | זה |  |
| 9:21.307 | מחברו |  |
| 9:21.308 | באמת | , |
| 9:21.309 | כשאר |  |
| 9:21.310 | ספרי |  |
| 9:21.311 | רז"ל |  |
| 9:21.312 | הידועים |  |
| 9:21.313 | ומפורסמים |  |
| 9:21.314 | מחבריהם |  |
| 9:21.315 | מיום |  |
| 9:21.316 | שחוברו |  |
| 9:21.317 | על |  |
| 9:21.318 | ידי |  |
| 9:21.319 | גדולי |  |
| 9:21.320 | ישראל |  |
| 9:21.321 | מעתיקי |  |
| 9:21.322 | השמועה |  |
| 9:21.323 | ועד |  |
| 9:21.324 | עתה | . |
| 9:21.325 | וגדול |  |
| 9:21.326 | ועצום |  |
| 9:21.327 | הוא |  |
| 9:21.328 | הק"ו |  |
| 9:21.329 | הזה | . |
| 9:21.330 | ומה |  |
| 9:21.331 | אם |  |
| 9:21.332 | נביא |  |
| 9:21.333 | העושה |  |
| 9:21.334 | אותות |  |
| 9:21.335 | ומופתים |  |
| 9:21.336 | לענינינו |  |
| 9:21.337 | בשמים |  |
| 9:21.338 | ובארץ |  |
| 9:21.339 | כמו |  |
| 9:21.340 | שפירשו |  |
| 9:21.341 | רז"ל |  |
| 9:21.342 | ומביאו |  |
| 9:21.343 | רש"י |  |
| 9:21.344 | ז"ל | , |
| 9:21.345 | אמרה |  |
| 9:21.346 | תורה |  |
| 9:21.347 | לא |  |
| 9:21.348 | תשמע |  |
| 9:21.349 | אל |  |
| 9:21.350 | דברי |  |
| 9:21.351 | הנביא |  |
| 9:21.352 | ההוא |  |
| 9:21.353 | כי |  |
| 9:21.354 | מנסה |  |
| 9:21.355 | ה | ' |
| 9:21.356 | אלהיכם |  |
| 9:21.357 | אתכם | ? |
| 9:21.358 | ספר |  |
| 9:21.359 | הזהר |  |
| 9:21.360 | הבא |  |
| 9:21.361 | לנו |  |
| 9:21.362 | מיד |  |
| 9:21.363 | עכו"ם | , |
| 9:21.364 | ומשמיע |  |
| 9:21.365 | גדולות |  |
| 9:21.366 | ונוראות |  |
| 9:21.367 | משם |  |
| 9:21.368 | תנאים |  |
| 9:21.369 | ואמוראים |  |
| 9:21.370 | ונשמות |  |
| 9:21.371 | צדיקים |  |
| 9:21.372 | שקדמו |  |
| 9:21.373 | לרשב"י |  |
| 9:21.374 | כמה |  |
| 9:21.375 | דורות | . |
| 9:21.376 | ואמוראים |  |
| 9:21.377 | המאוחרים |  |
| 9:21.378 | מן |  |
| 9:21.379 | רב |  |
| 9:21.380 | מאד |  |
| 9:21.381 | אחר |  |
| 9:21.382 | רשב"י | , |
| 9:21.383 | ואליהו |  |
| 9:21.384 | ועתיקא |  |
| 9:21.385 | קדישא |  |
| 9:21.386 | שנגלה |  |
| 9:21.387 | לנו |  |
| 9:21.388 | בבית |  |
| 9:21.389 | מדרשו |  |
| 9:21.390 | בין |  |
| 9:21.391 | חביריו |  |
| 9:21.392 | ותלמידיו | ) |
| 9:21.393 | וכיוצא |  |
| 9:21.394 | בספורים |  |
| 9:21.395 | אלו |  |
| 9:21.396 | הרשומים |  |
| 9:21.397 | בדיו |  |
| 9:21.398 | שחור |  |
| 9:21.399 | על |  |
| 9:21.400 | נייר |  |
| 9:21.401 | לבן | , |
| 9:21.402 | אשר |  |
| 9:21.403 | לא |  |
| 9:21.404 | ראינו |  |
| 9:21.405 | בעינינו |  |
| 9:21.406 | ולא |  |
| 9:21.407 | אבותינו |  |
| 9:21.408 | ספרו |  |
| 9:21.409 | לנו |  |
| 9:21.410 | כל |  |
| 9:21.411 | אותן |  |
| 9:21.412 | הדברים |  |
| 9:21.413 | והמעשיות |  |
| 9:21.414 | הרשומים |  |
| 9:21.415 | שם | ! |
| 9:21.416 | כ"ש |  |
| 9:21.417 | שלא |  |
| 9:21.418 | נשמע |  |
| 9:21.419 | בליו |  |
| 9:21.420 | לעבוד |  |
| 9:21.421 | שום |  |
| 9:21.422 | פרצוף |  |
| 9:21.423 | מכל |  |
| 9:21.424 | הפרצופים | , |
| 9:21.425 | ובפרט |  |
| 9:21.426 | קצר |  |
| 9:21.427 | אפים |  |
| 9:21.428 | ז"א | ) |
| 9:21.429 | כי |  |
| 9:21.430 | אם |  |
| 9:21.431 | את |  |
| 9:21.432 | אלהינו |  |
| 9:21.433 | בהו"ש |  |
| 9:21.434 | אשר |  |
| 9:21.435 | קבלנו |  |
| 9:21.436 | מאבותינו | , |
| 9:21.437 | ותוה"ק |  |
| 9:21.438 | ספרה |  |
| 9:21.439 | לנו |  |
| 9:21.440 | שאין |  |
| 9:21.441 | לו |  |
| 9:21.442 | דמות |  |
| 9:21.443 | הגוף | , |
| 9:21.444 | ואין |  |
| 9:21.445 | ראשית |  |
| 9:21.446 | לראשיתו | , |
| 9:21.447 | והוא |  |
| 9:21.448 | אחד |  |
| 9:21.449 | ואין |  |
| 9:21.450 | יחוד |  |
| 9:21.451 | כיחודו |  |
| 9:21.452 | לפי |  |
| 9:21.453 | שכל |  |
| 9:21.454 | הפרצופים |  |
| 9:21.455 | והעגולים |  |
| 9:21.456 | כולם |  |
| 9:21.457 | יש |  |
| 9:21.458 | ראשית |  |
| 9:21.459 | לראשיתם | , |
| 9:21.460 | ואין |  |
| 9:21.461 | לעבוד |  |
| 9:21.462 | רק |  |
| 9:21.463 | לצור |  |
| 9:21.464 | הכל |  |
| 9:21.465 | בלי |  |
| 9:21.466 | שום |  |
| 9:21.467 | שתוף |  |
| 9:21.468 | נברא |  |
| 9:21.469 | או |  |
| 9:21.470 | נאצל |  |
| 9:21.471 | מכל |  |
| 9:21.472 | הנבראים |  |
| 9:21.473 | והנאצלים | , |
| 9:21.474 | וגם |  |
| 9:21.475 | אין |  |
| 9:21.476 | ליחד |  |
| 9:21.477 | ולקשר |  |
| 9:21.478 | עם |  |
| 9:21.479 | אחדותו |  |
| 9:21.480 | שום |  |
| 9:21.481 | אחד |  |
| 9:21.482 | מהם | , |
| 9:21.483 | כמאמר |  |
| 9:21.484 | רשב"י |  |
| 9:21.485 | כל |  |
| 9:21.486 | המשתף |  |
| 9:21.487 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

