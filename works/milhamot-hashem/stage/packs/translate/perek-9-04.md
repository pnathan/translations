# Translation pack: milhamot-hashem perek-9-04 (Milchamot ha-Shem 9:4–9:4)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-04`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
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


## Milchamot ha-Shem 9:4

Hebrew (Wikisource): קכג) ובנדון דידן איכא קל וחומר גדול ועצום, ומה אם להחליף תורה או מצוה מכל המצות לא נשמעו לו כל שכן שלא נשמע למלך העכו"ם שמסר לנו את הספר הנזכר המסית להחליף ולהמיר את ה' אלהינו אחד בחמשה פרצופים, וביחוד לזעיר אנפין ופילגשו ולמר בפנינו ולחשוב בדעתינו שהוא אלהינו, והוא המלך המהולל בכל התושבחות, וכל המצות שנעשה נעשם לשם האלוה קצר האפים. ונשיב אל לבינו בשום שכל, ונדע בבירור כי מנסה ה' אלהינו אותנו כמו שאמר השי"ת בתורה כי יקום בקרבך נביא וכו' ובא האות והמופת וכו' כי מנסה ה' אלהיכם אתכם. לדעת הישכם אוהבים את ה' אלהיכם, וכמו שביאר הרבה המורה בפרק כ"ד משלישי שרוצה לומר כי בשיקום מתפאר בנבואה, ותראו אותותיו ומופתיו המביאות לחשוב אמת בדבריו. דעו שהוא ענין שירצהו השם ית' להודיע באומות שיעור האמינכם באמתת תורתו והשגתם אמתתו ית' ושאינכם ניסיתם להסתת מסית ולא תפסד אמונתכם בשם ית', למען יכוין אליה כל מבקש אמת. ויבקש מן האמונות מה שהוא אמת ויציב, וקיים קיום שאין לפנות עמו עוד לעשיית אות ומופת, שהמפתה לעשות הפכו באות או מופת. הוא קורא להאמין בנמנעות, כי לוא יועיל המופת רק למתנבא בדבר אפשרי כמו שביארנו במשנה תורה עכ"ל. אוי נא לנו כי נסה ה' אלהינו אותנו בהמציאו לנו את ספר הנזכר הזהר והתיקונים על ידי המלך העכו"ם ולא עמדנו בנסיון, ונפתינו להסתת המסית הנזכר לעבודת אלהים אחרים רבים, המשתתפים עם הבן קצר אפים (ז"א) ופילגשו על ידי הבטחתו שהוא מבטיח את קוראיו, שהזעיר אנפין מושל השפע מאצל העתיק כמו שכתב הרב יושר לבב וכסא אליהו! וכדי לפתותינו הוא משבח את סחורתו בכל דיבור במלות זכאה חולקיהן או זכאה הולקיה דמאן דעביד כך וכך, ווי ליה מאן דלא אזדהר בכך וכך, וכמתלהלה היורה זיקים חצים ומות, והא מפחית ומשפיל ומוריד את כבוד המשנה והתלמוד כלאחר יד, ומגדיל ומשבח את תורתו החדשה, ומתנבא של עתיד בסוף הימים, יעזבו את המשנה והתלמוד ויתפרנסון מחבורא דיליה, ועל ידי הלמוד בספריו יזכו להגאל, ובדברים האלו לבדן בלי שום אות ומופת צלחה לו להטות לב רבים וגדולים מישראל לעבוד אלהים חדשים בעלי גוף וגויה רמ"ח איברים ושס"ה גידים גדולים ועצומים, זכים ובהירים מאירים ומזהירים יתארם הקורא בשרד שכלו, ובמחוגת בינתו, אך ורק שהחומר שלהם איננו מדם ובשר _כמו שכתב הרב המורה בפרק א' מראשון) וקורא להם בשמותיו של הקב"ה.

| id | word | punctuation after |
|---|---|---|
| 9:4.1 | קכג | ) |
| 9:4.2 | ובנדון |  |
| 9:4.3 | דידן |  |
| 9:4.4 | איכא |  |
| 9:4.5 | קל |  |
| 9:4.6 | וחומר |  |
| 9:4.7 | גדול |  |
| 9:4.8 | ועצום | , |
| 9:4.9 | ומה |  |
| 9:4.10 | אם |  |
| 9:4.11 | להחליף |  |
| 9:4.12 | תורה |  |
| 9:4.13 | או |  |
| 9:4.14 | מצוה |  |
| 9:4.15 | מכל |  |
| 9:4.16 | המצות |  |
| 9:4.17 | לא |  |
| 9:4.18 | נשמעו |  |
| 9:4.19 | לו |  |
| 9:4.20 | כל |  |
| 9:4.21 | שכן |  |
| 9:4.22 | שלא |  |
| 9:4.23 | נשמע |  |
| 9:4.24 | למלך |  |
| 9:4.25 | העכו"ם |  |
| 9:4.26 | שמסר |  |
| 9:4.27 | לנו |  |
| 9:4.28 | את |  |
| 9:4.29 | הספר |  |
| 9:4.30 | הנזכר |  |
| 9:4.31 | המסית |  |
| 9:4.32 | להחליף |  |
| 9:4.33 | ולהמיר |  |
| 9:4.34 | את |  |
| 9:4.35 | ה | ' |
| 9:4.36 | אלהינו |  |
| 9:4.37 | אחד |  |
| 9:4.38 | בחמשה |  |
| 9:4.39 | פרצופים | , |
| 9:4.40 | וביחוד |  |
| 9:4.41 | לזעיר |  |
| 9:4.42 | אנפין |  |
| 9:4.43 | ופילגשו |  |
| 9:4.44 | ולמר |  |
| 9:4.45 | בפנינו |  |
| 9:4.46 | ולחשוב |  |
| 9:4.47 | בדעתינו |  |
| 9:4.48 | שהוא |  |
| 9:4.49 | אלהינו | , |
| 9:4.50 | והוא |  |
| 9:4.51 | המלך |  |
| 9:4.52 | המהולל |  |
| 9:4.53 | בכל |  |
| 9:4.54 | התושבחות | , |
| 9:4.55 | וכל |  |
| 9:4.56 | המצות |  |
| 9:4.57 | שנעשה |  |
| 9:4.58 | נעשם |  |
| 9:4.59 | לשם |  |
| 9:4.60 | האלוה |  |
| 9:4.61 | קצר |  |
| 9:4.62 | האפים | . |
| 9:4.63 | ונשיב |  |
| 9:4.64 | אל |  |
| 9:4.65 | לבינו |  |
| 9:4.66 | בשום |  |
| 9:4.67 | שכל | , |
| 9:4.68 | ונדע |  |
| 9:4.69 | בבירור |  |
| 9:4.70 | כי |  |
| 9:4.71 | מנסה |  |
| 9:4.72 | ה | ' |
| 9:4.73 | אלהינו |  |
| 9:4.74 | אותנו |  |
| 9:4.75 | כמו |  |
| 9:4.76 | שאמר |  |
| 9:4.77 | השי"ת |  |
| 9:4.78 | בתורה |  |
| 9:4.79 | כי |  |
| 9:4.80 | יקום |  |
| 9:4.81 | בקרבך |  |
| 9:4.82 | נביא |  |
| 9:4.83 | וכו | ' |
| 9:4.84 | ובא |  |
| 9:4.85 | האות |  |
| 9:4.86 | והמופת |  |
| 9:4.87 | וכו | ' |
| 9:4.88 | כי |  |
| 9:4.89 | מנסה |  |
| 9:4.90 | ה | ' |
| 9:4.91 | אלהיכם |  |
| 9:4.92 | אתכם | . |
| 9:4.93 | לדעת |  |
| 9:4.94 | הישכם |  |
| 9:4.95 | אוהבים |  |
| 9:4.96 | את |  |
| 9:4.97 | ה | ' |
| 9:4.98 | אלהיכם | , |
| 9:4.99 | וכמו |  |
| 9:4.100 | שביאר |  |
| 9:4.101 | הרבה |  |
| 9:4.102 | המורה |  |
| 9:4.103 | בפרק |  |
| 9:4.104 | כ"ד |  |
| 9:4.105 | משלישי |  |
| 9:4.106 | שרוצה |  |
| 9:4.107 | לומר |  |
| 9:4.108 | כי |  |
| 9:4.109 | בשיקום |  |
| 9:4.110 | מתפאר |  |
| 9:4.111 | בנבואה | , |
| 9:4.112 | ותראו |  |
| 9:4.113 | אותותיו |  |
| 9:4.114 | ומופתיו |  |
| 9:4.115 | המביאות |  |
| 9:4.116 | לחשוב |  |
| 9:4.117 | אמת |  |
| 9:4.118 | בדבריו | . |
| 9:4.119 | דעו |  |
| 9:4.120 | שהוא |  |
| 9:4.121 | ענין |  |
| 9:4.122 | שירצהו |  |
| 9:4.123 | השם |  |
| 9:4.124 | ית | ' |
| 9:4.125 | להודיע |  |
| 9:4.126 | באומות |  |
| 9:4.127 | שיעור |  |
| 9:4.128 | האמינכם |  |
| 9:4.129 | באמתת |  |
| 9:4.130 | תורתו |  |
| 9:4.131 | והשגתם |  |
| 9:4.132 | אמתתו |  |
| 9:4.133 | ית | ' |
| 9:4.134 | ושאינכם |  |
| 9:4.135 | ניסיתם |  |
| 9:4.136 | להסתת |  |
| 9:4.137 | מסית |  |
| 9:4.138 | ולא |  |
| 9:4.139 | תפסד |  |
| 9:4.140 | אמונתכם |  |
| 9:4.141 | בשם |  |
| 9:4.142 | ית | ', |
| 9:4.143 | למען |  |
| 9:4.144 | יכוין |  |
| 9:4.145 | אליה |  |
| 9:4.146 | כל |  |
| 9:4.147 | מבקש |  |
| 9:4.148 | אמת | . |
| 9:4.149 | ויבקש |  |
| 9:4.150 | מן |  |
| 9:4.151 | האמונות |  |
| 9:4.152 | מה |  |
| 9:4.153 | שהוא |  |
| 9:4.154 | אמת |  |
| 9:4.155 | ויציב | , |
| 9:4.156 | וקיים |  |
| 9:4.157 | קיום |  |
| 9:4.158 | שאין |  |
| 9:4.159 | לפנות |  |
| 9:4.160 | עמו |  |
| 9:4.161 | עוד |  |
| 9:4.162 | לעשיית |  |
| 9:4.163 | אות |  |
| 9:4.164 | ומופת | , |
| 9:4.165 | שהמפתה |  |
| 9:4.166 | לעשות |  |
| 9:4.167 | הפכו |  |
| 9:4.168 | באות |  |
| 9:4.169 | או |  |
| 9:4.170 | מופת | . |
| 9:4.171 | הוא |  |
| 9:4.172 | קורא |  |
| 9:4.173 | להאמין |  |
| 9:4.174 | בנמנעות | , |
| 9:4.175 | כי |  |
| 9:4.176 | לוא |  |
| 9:4.177 | יועיל |  |
| 9:4.178 | המופת |  |
| 9:4.179 | רק |  |
| 9:4.180 | למתנבא |  |
| 9:4.181 | בדבר |  |
| 9:4.182 | אפשרי |  |
| 9:4.183 | כמו |  |
| 9:4.184 | שביארנו |  |
| 9:4.185 | במשנה |  |
| 9:4.186 | תורה |  |
| 9:4.187 | עכ"ל | . |
| 9:4.188 | אוי |  |
| 9:4.189 | נא |  |
| 9:4.190 | לנו |  |
| 9:4.191 | כי |  |
| 9:4.192 | נסה |  |
| 9:4.193 | ה | ' |
| 9:4.194 | אלהינו |  |
| 9:4.195 | אותנו |  |
| 9:4.196 | בהמציאו |  |
| 9:4.197 | לנו |  |
| 9:4.198 | את |  |
| 9:4.199 | ספר |  |
| 9:4.200 | הנזכר |  |
| 9:4.201 | הזהר |  |
| 9:4.202 | והתיקונים |  |
| 9:4.203 | על |  |
| 9:4.204 | ידי |  |
| 9:4.205 | המלך |  |
| 9:4.206 | העכו"ם |  |
| 9:4.207 | ולא |  |
| 9:4.208 | עמדנו |  |
| 9:4.209 | בנסיון | , |
| 9:4.210 | ונפתינו |  |
| 9:4.211 | להסתת |  |
| 9:4.212 | המסית |  |
| 9:4.213 | הנזכר |  |
| 9:4.214 | לעבודת |  |
| 9:4.215 | אלהים |  |
| 9:4.216 | אחרים |  |
| 9:4.217 | רבים | , |
| 9:4.218 | המשתתפים |  |
| 9:4.219 | עם |  |
| 9:4.220 | הבן |  |
| 9:4.221 | קצר |  |
| 9:4.222 | אפים |  |
| 9:4.223 | ז"א | ) |
| 9:4.224 | ופילגשו |  |
| 9:4.225 | על |  |
| 9:4.226 | ידי |  |
| 9:4.227 | הבטחתו |  |
| 9:4.228 | שהוא |  |
| 9:4.229 | מבטיח |  |
| 9:4.230 | את |  |
| 9:4.231 | קוראיו | , |
| 9:4.232 | שהזעיר |  |
| 9:4.233 | אנפין |  |
| 9:4.234 | מושל |  |
| 9:4.235 | השפע |  |
| 9:4.236 | מאצל |  |
| 9:4.237 | העתיק |  |
| 9:4.238 | כמו |  |
| 9:4.239 | שכתב |  |
| 9:4.240 | הרב |  |
| 9:4.241 | יושר |  |
| 9:4.242 | לבב |  |
| 9:4.243 | וכסא |  |
| 9:4.244 | אליהו | ! |
| 9:4.245 | וכדי |  |
| 9:4.246 | לפתותינו |  |
| 9:4.247 | הוא |  |
| 9:4.248 | משבח |  |
| 9:4.249 | את |  |
| 9:4.250 | סחורתו |  |
| 9:4.251 | בכל |  |
| 9:4.252 | דיבור |  |
| 9:4.253 | במלות |  |
| 9:4.254 | זכאה |  |
| 9:4.255 | חולקיהן |  |
| 9:4.256 | או |  |
| 9:4.257 | זכאה |  |
| 9:4.258 | הולקיה |  |
| 9:4.259 | דמאן |  |
| 9:4.260 | דעביד |  |
| 9:4.261 | כך |  |
| 9:4.262 | וכך | , |
| 9:4.263 | ווי |  |
| 9:4.264 | ליה |  |
| 9:4.265 | מאן |  |
| 9:4.266 | דלא |  |
| 9:4.267 | אזדהר |  |
| 9:4.268 | בכך |  |
| 9:4.269 | וכך | , |
| 9:4.270 | וכמתלהלה |  |
| 9:4.271 | היורה |  |
| 9:4.272 | זיקים |  |
| 9:4.273 | חצים |  |
| 9:4.274 | ומות | , |
| 9:4.275 | והא |  |
| 9:4.276 | מפחית |  |
| 9:4.277 | ומשפיל |  |
| 9:4.278 | ומוריד |  |
| 9:4.279 | את |  |
| 9:4.280 | כבוד |  |
| 9:4.281 | המשנה |  |
| 9:4.282 | והתלמוד |  |
| 9:4.283 | כלאחר |  |
| 9:4.284 | יד | , |
| 9:4.285 | ומגדיל |  |
| 9:4.286 | ומשבח |  |
| 9:4.287 | את |  |
| 9:4.288 | תורתו |  |
| 9:4.289 | החדשה | , |
| 9:4.290 | ומתנבא |  |
| 9:4.291 | של |  |
| 9:4.292 | עתיד |  |
| 9:4.293 | בסוף |  |
| 9:4.294 | הימים | , |
| 9:4.295 | יעזבו |  |
| 9:4.296 | את |  |
| 9:4.297 | המשנה |  |
| 9:4.298 | והתלמוד |  |
| 9:4.299 | ויתפרנסון |  |
| 9:4.300 | מחבורא |  |
| 9:4.301 | דיליה | , |
| 9:4.302 | ועל |  |
| 9:4.303 | ידי |  |
| 9:4.304 | הלמוד |  |
| 9:4.305 | בספריו |  |
| 9:4.306 | יזכו |  |
| 9:4.307 | להגאל | , |
| 9:4.308 | ובדברים |  |
| 9:4.309 | האלו |  |
| 9:4.310 | לבדן |  |
| 9:4.311 | בלי |  |
| 9:4.312 | שום |  |
| 9:4.313 | אות |  |
| 9:4.314 | ומופת |  |
| 9:4.315 | צלחה |  |
| 9:4.316 | לו |  |
| 9:4.317 | להטות |  |
| 9:4.318 | לב |  |
| 9:4.319 | רבים |  |
| 9:4.320 | וגדולים |  |
| 9:4.321 | מישראל |  |
| 9:4.322 | לעבוד |  |
| 9:4.323 | אלהים |  |
| 9:4.324 | חדשים |  |
| 9:4.325 | בעלי |  |
| 9:4.326 | גוף |  |
| 9:4.327 | וגויה |  |
| 9:4.328 | רמ"ח |  |
| 9:4.329 | איברים |  |
| 9:4.330 | ושס"ה |  |
| 9:4.331 | גידים |  |
| 9:4.332 | גדולים |  |
| 9:4.333 | ועצומים | , |
| 9:4.334 | זכים |  |
| 9:4.335 | ובהירים |  |
| 9:4.336 | מאירים |  |
| 9:4.337 | ומזהירים |  |
| 9:4.338 | יתארם |  |
| 9:4.339 | הקורא |  |
| 9:4.340 | בשרד |  |
| 9:4.341 | שכלו | , |
| 9:4.342 | ובמחוגת |  |
| 9:4.343 | בינתו | , |
| 9:4.344 | אך |  |
| 9:4.345 | ורק |  |
| 9:4.346 | שהחומר |  |
| 9:4.347 | שלהם |  |
| 9:4.348 | איננו |  |
| 9:4.349 | מדם |  |
| 9:4.350 | ובשר |  |
| 9:4.351 | כמו |  |
| 9:4.352 | שכתב |  |
| 9:4.353 | הרב |  |
| 9:4.354 | המורה |  |
| 9:4.355 | בפרק |  |
| 9:4.356 | א | ' |
| 9:4.357 | מראשון | ) |
| 9:4.358 | וקורא |  |
| 9:4.359 | להם |  |
| 9:4.360 | בשמותיו |  |
| 9:4.361 | של |  |
| 9:4.362 | הקב"ה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

