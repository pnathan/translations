# Translation pack: milhamot-hashem perek-7-22 (Milchamot ha-Shem 7:27–7:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-22`.

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
- **אצילות / באצילות / בריאה**: fixed, translated, the verbs linked:
  "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; נאצל → "emanated"; המאציל → "the Emanator". The four worlds
  (אבי"ע): "Emanation, Creation, Formation, Making" (verbs: "emanated,
  created, formed, made"). נאצל and נברא stay distinct ("נברא ולא בורא" →
  "created, not creator").
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 7:27

Hebrew (Wikisource): צט) עוד כתבו שם הגדולים והגליון הנז' משם הרב ישר בס' מצרף לחכמה ששמע מפי מגידי אמת, שבשנת ש"ף לאלך השישי כשבזזו הספרדים עיר הירילברג"ה לקחו מהאקימי"א כמה אלפים ספרים וביניהם ספרי הקדש על קלף. ובכלל ספרי הזהר על כלל הכ"ד ספרים משא סבל ושולחו הספרים לחשמנים ודוכוסים וכו'. גם שמועה זו ששמע הרב יש"ר שקר היא שהרי בשנת שי"ח כבר נדפס ספר הזהר ונפשט. ומתיא זו דאיתי בידיה לאו שמה מתיא, ועוד דמ"מ מידי עכו"ם יצא. והגאון ר' יעקב עמדין בספרו מטפחת ספרים כ' ב' רבי יצחק דמן עכו וז"ל שנת נ"ג ה"ר משה מליאון שוכן באויל"א (והוא בואר אלחגרה) היה בזמן הזה, ולמען תספרו לדור אחרון אודיעכם מה שמצאתי כתוב שר' יצחק דמן עכו הלך לחקור על ספר הזהר שכנראה היו דבריו מפליאים ישאבו ממקור עליון. ואשר הוא לשון ירושלמי נר' שדבריו דברי קבלה ואמת. ואשר הם לשון הקדש הם דברי זיוף. אמר ושאלתי את הלמידים שנמצא בידיהם בפר הזהר מאין בא להם. ולא ראיתי דבריהים מכוונים. יש מי שאמר שהרמב"ן מצא אותו בארץ ישראל ושלחו לקאטולוניא, והביאו לארגון. ונפל ביד רבי משה ליאון. וי"א שמעולם לא היה חבור של רשב"י. רק שר' משה ידע שם הכותב, ובכחו יכתוב דברים מפליאים. ולמען יקח בהם מחיר רב. היה תולה עצמו באילן גדול שהוא רשב"י וחביריו ובנו ר' אלעזר. ואבוא אל עיר יאליאדול"ר. ואמצא שם רבי משה דיליאן וישבע לי שספר רשב"י הנו בבית ישיבתו ואמצא שם חכם זקן ושמו ר' דוד רסאן קרובו ואמצא חן בעיני, ואשביעהו לאמר אם היה יודע אמתת ס' הזהר אם היה אמת אם לא. ויאמר החכם כי נתבאר לו בודאי שר' משה בשם הכותב היה בוראם. וסבת האמתות הוא כי ר' משה הנז' היה כותב סודות ונפלאות לעשירים אשר במלכות ההיא ולוקח בהם מתנות רבות זהב וכסף. ואותו היום כל מה שהיה מרויח היה מפזר עד שביום מותו לא נשאר לו אפי' פרוטה ואשתו ובניו ברעב ובצמא ובעירום ובחסד כל וכשמעי מיתתו קמתי והלכתי אל רבי יוסף דיאוילא שהיה עשיר גדול, ובחייו לקח ממנו ממון, ואומר לו הנה עתך עת דודים, תזכה לספר הגדול, ותקח הספר מיד אשתו במפה יפה שתשלח בידי אשתך. שהיא כעת ערומה, וכן עשה שלח ביד אשתו מתנות אל אשת ר' משה הנז' ותשבע לה אשת ר' משה הנז', שמעולם לא היה ספר לאישי רק מלבו ודעתו היה כותב. ואומר לאישי למה אתה תולה עיונך בחכם אחר? הלא טוב היה לך לומר שמשכלך אתה כותב. וידוך כי תיטיב לך. ואישי ענני שאם הייתי אומר שמדעתי הם לא יחשיבו אותם ולא יתנו מתנות בהם. ולכן אני תולה אותם ברשב"י וחביריו, ואחר דברתי עם בתו ואמצא דבריהם מכוונים. אלה הדברים אשר דבר אלי החכם הזקן רבי דוד, ואסע מאויל"א ואבוא אל טלאבירה מצאתי שם זקן מופלא שמו רבי יוסף הלוי בן ר' יוסף טורדוס המקובל, ואחקורה שמו ולדעת מה שידו מגעת בידיעת ספר הזהר, ויאמר אלי דע באמת כי היה ביד ר' משה ספר הזהר אמתי, וסבת אמתות דברי הוא כי היה נותן לי קונדריסים ואני לבחון הספק הזה. לקחתי קונדריס אחד מבין הקונדריסים ואמרתי לו שנאבד ממני שיתן לי ממלא מקומו, ויאמר לי הראה לי אחרית הקונדריס הקודם לו, וראשית הקונדריס המתאחר, ויכתוב לי קונדריס. וארא והנה אין ביניהם שינוי כלל אפילו תיבה אחת, רק שפה א' ודברים אחדים. היש נסיון גדול מזה? אחר זה באתי לטוליטולא ואמרו לי שבחינת ר' יוסף הלוי אינה בחינה, שאולי קודם לחזות הקונדריסים לאדם היה מעתיק אחד לעצמו, ואחר כך מתוך אותם הקונדריסים אשר בביתו היה כותב לאחרים. ע"כ הדברים שהביאם בס' יוחסין. ושם הגיה המביא ספר יוחסין דברי מה בכך נגד הדברים האלה, וכבר השיב על דבריו הגאון יעב"ץ! והוא מאריך עור לפשר לקיים מקצת ולבטל מקצת יעויין שם באורך. והנה מזה תבין שלא נודע לספר הזהר ומקור שממנו יצא ונתפשט. כשאר חיבורי רז"ל המפורסמים מיום התגלותם. המשנה והתלמוד בבלי וירושלמי ותוספתא וספרא וספרי שכולם ידוע ומפורסם מקורן ומי הוא מחברן. אבל ספר הזהר המסית זה אומר בכה וזה אומר בכה. וקורא אני עליו מאמר רז"ל בירושלמי שבת פ' ר"א דמילה ופסחים פרק אלו דברים שאמרו שם רבי זעירא בשם רבי אלעזר כל תורה שאין לה בית אב אינה תורה עכ"ל. ואל תפנה למתק שפתיו במה שהוא מלהיב את לב קוראיו ומעוררם ליראת שמים ולהשתדל בעסק תורתו החדשה שהוא קורא לה בת מלך. לפי שאין כוונתו רק להביא את הקורא לאמונה נכרית להאמין באלהים רבים משפחות משפחות. משפחת האלוהות שבעולם האצילות, שהם אדם קדמון, ועתיק יומין, ואריך אנפין, ואבא ואימא, וזעיר אנפין ונוקביה. וכנגדן בעולם הבריאה, וכמוהם בעולם היצירה, וכן בעולם העשייה, ואת כולם הם מקבלים אותם לאלוהות ומשבחים ומפארים אותם בשעת התפלה באמרם אין כאלהינו וכו' כמו שכתבתי לעיל סי' ע"ט בשם סידור חסד לאברהם. דודאי דרושים כאלה אסור לשמען וכל שכן לקרותן ולהאמין בהם.

| id | word | punctuation after |
|---|---|---|
| 7:27.1 | צט | ) |
| 7:27.2 | עוד |  |
| 7:27.3 | כתבו |  |
| 7:27.4 | שם |  |
| 7:27.5 | הגדולים |  |
| 7:27.6 | והגליון |  |
| 7:27.7 | הנז | ' |
| 7:27.8 | משם |  |
| 7:27.9 | הרב |  |
| 7:27.10 | ישר |  |
| 7:27.11 | בס | ' |
| 7:27.12 | מצרף |  |
| 7:27.13 | לחכמה |  |
| 7:27.14 | ששמע |  |
| 7:27.15 | מפי |  |
| 7:27.16 | מגידי |  |
| 7:27.17 | אמת | , |
| 7:27.18 | שבשנת |  |
| 7:27.19 | ש"ף |  |
| 7:27.20 | לאלך |  |
| 7:27.21 | השישי |  |
| 7:27.22 | כשבזזו |  |
| 7:27.23 | הספרדים |  |
| 7:27.24 | עיר |  |
| 7:27.25 | הירילברג"ה |  |
| 7:27.26 | לקחו |  |
| 7:27.27 | מהאקימי"א |  |
| 7:27.28 | כמה |  |
| 7:27.29 | אלפים |  |
| 7:27.30 | ספרים |  |
| 7:27.31 | וביניהם |  |
| 7:27.32 | ספרי |  |
| 7:27.33 | הקדש |  |
| 7:27.34 | על |  |
| 7:27.35 | קלף | . |
| 7:27.36 | ובכלל |  |
| 7:27.37 | ספרי |  |
| 7:27.38 | הזהר |  |
| 7:27.39 | על |  |
| 7:27.40 | כלל |  |
| 7:27.41 | הכ"ד |  |
| 7:27.42 | ספרים |  |
| 7:27.43 | משא |  |
| 7:27.44 | סבל |  |
| 7:27.45 | ושולחו |  |
| 7:27.46 | הספרים |  |
| 7:27.47 | לחשמנים |  |
| 7:27.48 | ודוכוסים |  |
| 7:27.49 | וכו | '. |
| 7:27.50 | גם |  |
| 7:27.51 | שמועה |  |
| 7:27.52 | זו |  |
| 7:27.53 | ששמע |  |
| 7:27.54 | הרב |  |
| 7:27.55 | יש"ר |  |
| 7:27.56 | שקר |  |
| 7:27.57 | היא |  |
| 7:27.58 | שהרי |  |
| 7:27.59 | בשנת |  |
| 7:27.60 | שי"ח |  |
| 7:27.61 | כבר |  |
| 7:27.62 | נדפס |  |
| 7:27.63 | ספר |  |
| 7:27.64 | הזהר |  |
| 7:27.65 | ונפשט | . |
| 7:27.66 | ומתיא |  |
| 7:27.67 | זו |  |
| 7:27.68 | דאיתי |  |
| 7:27.69 | בידיה |  |
| 7:27.70 | לאו |  |
| 7:27.71 | שמה |  |
| 7:27.72 | מתיא | , |
| 7:27.73 | ועוד |  |
| 7:27.74 | דמ"מ |  |
| 7:27.75 | מידי |  |
| 7:27.76 | עכו"ם |  |
| 7:27.77 | יצא | . |
| 7:27.78 | והגאון |  |
| 7:27.79 | ר | ' |
| 7:27.80 | יעקב |  |
| 7:27.81 | עמדין |  |
| 7:27.82 | בספרו |  |
| 7:27.83 | מטפחת |  |
| 7:27.84 | ספרים |  |
| 7:27.85 | כ | ' |
| 7:27.86 | ב | ' |
| 7:27.87 | רבי |  |
| 7:27.88 | יצחק |  |
| 7:27.89 | דמן |  |
| 7:27.90 | עכו |  |
| 7:27.91 | וז"ל |  |
| 7:27.92 | שנת |  |
| 7:27.93 | נ"ג |  |
| 7:27.94 | ה"ר |  |
| 7:27.95 | משה |  |
| 7:27.96 | מליאון |  |
| 7:27.97 | שוכן |  |
| 7:27.98 | באויל"א |  |
| 7:27.99 | והוא |  |
| 7:27.100 | בואר |  |
| 7:27.101 | אלחגרה | ) |
| 7:27.102 | היה |  |
| 7:27.103 | בזמן |  |
| 7:27.104 | הזה | , |
| 7:27.105 | ולמען |  |
| 7:27.106 | תספרו |  |
| 7:27.107 | לדור |  |
| 7:27.108 | אחרון |  |
| 7:27.109 | אודיעכם |  |
| 7:27.110 | מה |  |
| 7:27.111 | שמצאתי |  |
| 7:27.112 | כתוב |  |
| 7:27.113 | שר | ' |
| 7:27.114 | יצחק |  |
| 7:27.115 | דמן |  |
| 7:27.116 | עכו |  |
| 7:27.117 | הלך |  |
| 7:27.118 | לחקור |  |
| 7:27.119 | על |  |
| 7:27.120 | ספר |  |
| 7:27.121 | הזהר |  |
| 7:27.122 | שכנראה |  |
| 7:27.123 | היו |  |
| 7:27.124 | דבריו |  |
| 7:27.125 | מפליאים |  |
| 7:27.126 | ישאבו |  |
| 7:27.127 | ממקור |  |
| 7:27.128 | עליון | . |
| 7:27.129 | ואשר |  |
| 7:27.130 | הוא |  |
| 7:27.131 | לשון |  |
| 7:27.132 | ירושלמי |  |
| 7:27.133 | נר | ' |
| 7:27.134 | שדבריו |  |
| 7:27.135 | דברי |  |
| 7:27.136 | קבלה |  |
| 7:27.137 | ואמת | . |
| 7:27.138 | ואשר |  |
| 7:27.139 | הם |  |
| 7:27.140 | לשון |  |
| 7:27.141 | הקדש |  |
| 7:27.142 | הם |  |
| 7:27.143 | דברי |  |
| 7:27.144 | זיוף | . |
| 7:27.145 | אמר |  |
| 7:27.146 | ושאלתי |  |
| 7:27.147 | את |  |
| 7:27.148 | הלמידים |  |
| 7:27.149 | שנמצא |  |
| 7:27.150 | בידיהם |  |
| 7:27.151 | בפר |  |
| 7:27.152 | הזהר |  |
| 7:27.153 | מאין |  |
| 7:27.154 | בא |  |
| 7:27.155 | להם | . |
| 7:27.156 | ולא |  |
| 7:27.157 | ראיתי |  |
| 7:27.158 | דבריהים |  |
| 7:27.159 | מכוונים | . |
| 7:27.160 | יש |  |
| 7:27.161 | מי |  |
| 7:27.162 | שאמר |  |
| 7:27.163 | שהרמב"ן |  |
| 7:27.164 | מצא |  |
| 7:27.165 | אותו |  |
| 7:27.166 | בארץ |  |
| 7:27.167 | ישראל |  |
| 7:27.168 | ושלחו |  |
| 7:27.169 | לקאטולוניא | , |
| 7:27.170 | והביאו |  |
| 7:27.171 | לארגון | . |
| 7:27.172 | ונפל |  |
| 7:27.173 | ביד |  |
| 7:27.174 | רבי |  |
| 7:27.175 | משה |  |
| 7:27.176 | ליאון | . |
| 7:27.177 | וי"א |  |
| 7:27.178 | שמעולם |  |
| 7:27.179 | לא |  |
| 7:27.180 | היה |  |
| 7:27.181 | חבור |  |
| 7:27.182 | של |  |
| 7:27.183 | רשב"י | . |
| 7:27.184 | רק |  |
| 7:27.185 | שר | ' |
| 7:27.186 | משה |  |
| 7:27.187 | ידע |  |
| 7:27.188 | שם |  |
| 7:27.189 | הכותב | , |
| 7:27.190 | ובכחו |  |
| 7:27.191 | יכתוב |  |
| 7:27.192 | דברים |  |
| 7:27.193 | מפליאים | . |
| 7:27.194 | ולמען |  |
| 7:27.195 | יקח |  |
| 7:27.196 | בהם |  |
| 7:27.197 | מחיר |  |
| 7:27.198 | רב | . |
| 7:27.199 | היה |  |
| 7:27.200 | תולה |  |
| 7:27.201 | עצמו |  |
| 7:27.202 | באילן |  |
| 7:27.203 | גדול |  |
| 7:27.204 | שהוא |  |
| 7:27.205 | רשב"י |  |
| 7:27.206 | וחביריו |  |
| 7:27.207 | ובנו |  |
| 7:27.208 | ר | ' |
| 7:27.209 | אלעזר | . |
| 7:27.210 | ואבוא |  |
| 7:27.211 | אל |  |
| 7:27.212 | עיר |  |
| 7:27.213 | יאליאדול"ר | . |
| 7:27.214 | ואמצא |  |
| 7:27.215 | שם |  |
| 7:27.216 | רבי |  |
| 7:27.217 | משה |  |
| 7:27.218 | דיליאן |  |
| 7:27.219 | וישבע |  |
| 7:27.220 | לי |  |
| 7:27.221 | שספר |  |
| 7:27.222 | רשב"י |  |
| 7:27.223 | הנו |  |
| 7:27.224 | בבית |  |
| 7:27.225 | ישיבתו |  |
| 7:27.226 | ואמצא |  |
| 7:27.227 | שם |  |
| 7:27.228 | חכם |  |
| 7:27.229 | זקן |  |
| 7:27.230 | ושמו |  |
| 7:27.231 | ר | ' |
| 7:27.232 | דוד |  |
| 7:27.233 | רסאן |  |
| 7:27.234 | קרובו |  |
| 7:27.235 | ואמצא |  |
| 7:27.236 | חן |  |
| 7:27.237 | בעיני | , |
| 7:27.238 | ואשביעהו |  |
| 7:27.239 | לאמר |  |
| 7:27.240 | אם |  |
| 7:27.241 | היה |  |
| 7:27.242 | יודע |  |
| 7:27.243 | אמתת |  |
| 7:27.244 | ס | ' |
| 7:27.245 | הזהר |  |
| 7:27.246 | אם |  |
| 7:27.247 | היה |  |
| 7:27.248 | אמת |  |
| 7:27.249 | אם |  |
| 7:27.250 | לא | . |
| 7:27.251 | ויאמר |  |
| 7:27.252 | החכם |  |
| 7:27.253 | כי |  |
| 7:27.254 | נתבאר |  |
| 7:27.255 | לו |  |
| 7:27.256 | בודאי |  |
| 7:27.257 | שר | ' |
| 7:27.258 | משה |  |
| 7:27.259 | בשם |  |
| 7:27.260 | הכותב |  |
| 7:27.261 | היה |  |
| 7:27.262 | בוראם | . |
| 7:27.263 | וסבת |  |
| 7:27.264 | האמתות |  |
| 7:27.265 | הוא |  |
| 7:27.266 | כי |  |
| 7:27.267 | ר | ' |
| 7:27.268 | משה |  |
| 7:27.269 | הנז | ' |
| 7:27.270 | היה |  |
| 7:27.271 | כותב |  |
| 7:27.272 | סודות |  |
| 7:27.273 | ונפלאות |  |
| 7:27.274 | לעשירים |  |
| 7:27.275 | אשר |  |
| 7:27.276 | במלכות |  |
| 7:27.277 | ההיא |  |
| 7:27.278 | ולוקח |  |
| 7:27.279 | בהם |  |
| 7:27.280 | מתנות |  |
| 7:27.281 | רבות |  |
| 7:27.282 | זהב |  |
| 7:27.283 | וכסף | . |
| 7:27.284 | ואותו |  |
| 7:27.285 | היום |  |
| 7:27.286 | כל |  |
| 7:27.287 | מה |  |
| 7:27.288 | שהיה |  |
| 7:27.289 | מרויח |  |
| 7:27.290 | היה |  |
| 7:27.291 | מפזר |  |
| 7:27.292 | עד |  |
| 7:27.293 | שביום |  |
| 7:27.294 | מותו |  |
| 7:27.295 | לא |  |
| 7:27.296 | נשאר |  |
| 7:27.297 | לו |  |
| 7:27.298 | אפי | ' |
| 7:27.299 | פרוטה |  |
| 7:27.300 | ואשתו |  |
| 7:27.301 | ובניו |  |
| 7:27.302 | ברעב |  |
| 7:27.303 | ובצמא |  |
| 7:27.304 | ובעירום |  |
| 7:27.305 | ובחסד |  |
| 7:27.306 | כל |  |
| 7:27.307 | וכשמעי |  |
| 7:27.308 | מיתתו |  |
| 7:27.309 | קמתי |  |
| 7:27.310 | והלכתי |  |
| 7:27.311 | אל |  |
| 7:27.312 | רבי |  |
| 7:27.313 | יוסף |  |
| 7:27.314 | דיאוילא |  |
| 7:27.315 | שהיה |  |
| 7:27.316 | עשיר |  |
| 7:27.317 | גדול | , |
| 7:27.318 | ובחייו |  |
| 7:27.319 | לקח |  |
| 7:27.320 | ממנו |  |
| 7:27.321 | ממון | , |
| 7:27.322 | ואומר |  |
| 7:27.323 | לו |  |
| 7:27.324 | הנה |  |
| 7:27.325 | עתך |  |
| 7:27.326 | עת |  |
| 7:27.327 | דודים | , |
| 7:27.328 | תזכה |  |
| 7:27.329 | לספר |  |
| 7:27.330 | הגדול | , |
| 7:27.331 | ותקח |  |
| 7:27.332 | הספר |  |
| 7:27.333 | מיד |  |
| 7:27.334 | אשתו |  |
| 7:27.335 | במפה |  |
| 7:27.336 | יפה |  |
| 7:27.337 | שתשלח |  |
| 7:27.338 | בידי |  |
| 7:27.339 | אשתך | . |
| 7:27.340 | שהיא |  |
| 7:27.341 | כעת |  |
| 7:27.342 | ערומה | , |
| 7:27.343 | וכן |  |
| 7:27.344 | עשה |  |
| 7:27.345 | שלח |  |
| 7:27.346 | ביד |  |
| 7:27.347 | אשתו |  |
| 7:27.348 | מתנות |  |
| 7:27.349 | אל |  |
| 7:27.350 | אשת |  |
| 7:27.351 | ר | ' |
| 7:27.352 | משה |  |
| 7:27.353 | הנז | ' |
| 7:27.354 | ותשבע |  |
| 7:27.355 | לה |  |
| 7:27.356 | אשת |  |
| 7:27.357 | ר | ' |
| 7:27.358 | משה |  |
| 7:27.359 | הנז | ', |
| 7:27.360 | שמעולם |  |
| 7:27.361 | לא |  |
| 7:27.362 | היה |  |
| 7:27.363 | ספר |  |
| 7:27.364 | לאישי |  |
| 7:27.365 | רק |  |
| 7:27.366 | מלבו |  |
| 7:27.367 | ודעתו |  |
| 7:27.368 | היה |  |
| 7:27.369 | כותב | . |
| 7:27.370 | ואומר |  |
| 7:27.371 | לאישי |  |
| 7:27.372 | למה |  |
| 7:27.373 | אתה |  |
| 7:27.374 | תולה |  |
| 7:27.375 | עיונך |  |
| 7:27.376 | בחכם |  |
| 7:27.377 | אחר | ? |
| 7:27.378 | הלא |  |
| 7:27.379 | טוב |  |
| 7:27.380 | היה |  |
| 7:27.381 | לך |  |
| 7:27.382 | לומר |  |
| 7:27.383 | שמשכלך |  |
| 7:27.384 | אתה |  |
| 7:27.385 | כותב | . |
| 7:27.386 | וידוך |  |
| 7:27.387 | כי |  |
| 7:27.388 | תיטיב |  |
| 7:27.389 | לך | . |
| 7:27.390 | ואישי |  |
| 7:27.391 | ענני |  |
| 7:27.392 | שאם |  |
| 7:27.393 | הייתי |  |
| 7:27.394 | אומר |  |
| 7:27.395 | שמדעתי |  |
| 7:27.396 | הם |  |
| 7:27.397 | לא |  |
| 7:27.398 | יחשיבו |  |
| 7:27.399 | אותם |  |
| 7:27.400 | ולא |  |
| 7:27.401 | יתנו |  |
| 7:27.402 | מתנות |  |
| 7:27.403 | בהם | . |
| 7:27.404 | ולכן |  |
| 7:27.405 | אני |  |
| 7:27.406 | תולה |  |
| 7:27.407 | אותם |  |
| 7:27.408 | ברשב"י |  |
| 7:27.409 | וחביריו | , |
| 7:27.410 | ואחר |  |
| 7:27.411 | דברתי |  |
| 7:27.412 | עם |  |
| 7:27.413 | בתו |  |
| 7:27.414 | ואמצא |  |
| 7:27.415 | דבריהם |  |
| 7:27.416 | מכוונים | . |
| 7:27.417 | אלה |  |
| 7:27.418 | הדברים |  |
| 7:27.419 | אשר |  |
| 7:27.420 | דבר |  |
| 7:27.421 | אלי |  |
| 7:27.422 | החכם |  |
| 7:27.423 | הזקן |  |
| 7:27.424 | רבי |  |
| 7:27.425 | דוד | , |
| 7:27.426 | ואסע |  |
| 7:27.427 | מאויל"א |  |
| 7:27.428 | ואבוא |  |
| 7:27.429 | אל |  |
| 7:27.430 | טלאבירה |  |
| 7:27.431 | מצאתי |  |
| 7:27.432 | שם |  |
| 7:27.433 | זקן |  |
| 7:27.434 | מופלא |  |
| 7:27.435 | שמו |  |
| 7:27.436 | רבי |  |
| 7:27.437 | יוסף |  |
| 7:27.438 | הלוי |  |
| 7:27.439 | בן |  |
| 7:27.440 | ר | ' |
| 7:27.441 | יוסף |  |
| 7:27.442 | טורדוס |  |
| 7:27.443 | המקובל | , |
| 7:27.444 | ואחקורה |  |
| 7:27.445 | שמו |  |
| 7:27.446 | ולדעת |  |
| 7:27.447 | מה |  |
| 7:27.448 | שידו |  |
| 7:27.449 | מגעת |  |
| 7:27.450 | בידיעת |  |
| 7:27.451 | ספר |  |
| 7:27.452 | הזהר | , |
| 7:27.453 | ויאמר |  |
| 7:27.454 | אלי |  |
| 7:27.455 | דע |  |
| 7:27.456 | באמת |  |
| 7:27.457 | כי |  |
| 7:27.458 | היה |  |
| 7:27.459 | ביד |  |
| 7:27.460 | ר | ' |
| 7:27.461 | משה |  |
| 7:27.462 | ספר |  |
| 7:27.463 | הזהר |  |
| 7:27.464 | אמתי | , |
| 7:27.465 | וסבת |  |
| 7:27.466 | אמתות |  |
| 7:27.467 | דברי |  |
| 7:27.468 | הוא |  |
| 7:27.469 | כי |  |
| 7:27.470 | היה |  |
| 7:27.471 | נותן |  |
| 7:27.472 | לי |  |
| 7:27.473 | קונדריסים |  |
| 7:27.474 | ואני |  |
| 7:27.475 | לבחון |  |
| 7:27.476 | הספק |  |
| 7:27.477 | הזה | . |
| 7:27.478 | לקחתי |  |
| 7:27.479 | קונדריס |  |
| 7:27.480 | אחד |  |
| 7:27.481 | מבין |  |
| 7:27.482 | הקונדריסים |  |
| 7:27.483 | ואמרתי |  |
| 7:27.484 | לו |  |
| 7:27.485 | שנאבד |  |
| 7:27.486 | ממני |  |
| 7:27.487 | שיתן |  |
| 7:27.488 | לי |  |
| 7:27.489 | ממלא |  |
| 7:27.490 | מקומו | , |
| 7:27.491 | ויאמר |  |
| 7:27.492 | לי |  |
| 7:27.493 | הראה |  |
| 7:27.494 | לי |  |
| 7:27.495 | אחרית |  |
| 7:27.496 | הקונדריס |  |
| 7:27.497 | הקודם |  |
| 7:27.498 | לו | , |
| 7:27.499 | וראשית |  |
| 7:27.500 | הקונדריס |  |
| 7:27.501 | המתאחר | , |
| 7:27.502 | ויכתוב |  |
| 7:27.503 | לי |  |
| 7:27.504 | קונדריס | . |
| 7:27.505 | וארא |  |
| 7:27.506 | והנה |  |
| 7:27.507 | אין |  |
| 7:27.508 | ביניהם |  |
| 7:27.509 | שינוי |  |
| 7:27.510 | כלל |  |
| 7:27.511 | אפילו |  |
| 7:27.512 | תיבה |  |
| 7:27.513 | אחת | , |
| 7:27.514 | רק |  |
| 7:27.515 | שפה |  |
| 7:27.516 | א | ' |
| 7:27.517 | ודברים |  |
| 7:27.518 | אחדים | . |
| 7:27.519 | היש |  |
| 7:27.520 | נסיון |  |
| 7:27.521 | גדול |  |
| 7:27.522 | מזה | ? |
| 7:27.523 | אחר |  |
| 7:27.524 | זה |  |
| 7:27.525 | באתי |  |
| 7:27.526 | לטוליטולא |  |
| 7:27.527 | ואמרו |  |
| 7:27.528 | לי |  |
| 7:27.529 | שבחינת |  |
| 7:27.530 | ר | ' |
| 7:27.531 | יוסף |  |
| 7:27.532 | הלוי |  |
| 7:27.533 | אינה |  |
| 7:27.534 | בחינה | , |
| 7:27.535 | שאולי |  |
| 7:27.536 | קודם |  |
| 7:27.537 | לחזות |  |
| 7:27.538 | הקונדריסים |  |
| 7:27.539 | לאדם |  |
| 7:27.540 | היה |  |
| 7:27.541 | מעתיק |  |
| 7:27.542 | אחד |  |
| 7:27.543 | לעצמו | , |
| 7:27.544 | ואחר |  |
| 7:27.545 | כך |  |
| 7:27.546 | מתוך |  |
| 7:27.547 | אותם |  |
| 7:27.548 | הקונדריסים |  |
| 7:27.549 | אשר |  |
| 7:27.550 | בביתו |  |
| 7:27.551 | היה |  |
| 7:27.552 | כותב |  |
| 7:27.553 | לאחרים | . |
| 7:27.554 | ע"כ |  |
| 7:27.555 | הדברים |  |
| 7:27.556 | שהביאם |  |
| 7:27.557 | בס | ' |
| 7:27.558 | יוחסין | . |
| 7:27.559 | ושם |  |
| 7:27.560 | הגיה |  |
| 7:27.561 | המביא |  |
| 7:27.562 | ספר |  |
| 7:27.563 | יוחסין |  |
| 7:27.564 | דברי |  |
| 7:27.565 | מה |  |
| 7:27.566 | בכך |  |
| 7:27.567 | נגד |  |
| 7:27.568 | הדברים |  |
| 7:27.569 | האלה | , |
| 7:27.570 | וכבר |  |
| 7:27.571 | השיב |  |
| 7:27.572 | על |  |
| 7:27.573 | דבריו |  |
| 7:27.574 | הגאון |  |
| 7:27.575 | יעב"ץ | ! |
| 7:27.576 | והוא |  |
| 7:27.577 | מאריך |  |
| 7:27.578 | עור |  |
| 7:27.579 | לפשר |  |
| 7:27.580 | לקיים |  |
| 7:27.581 | מקצת |  |
| 7:27.582 | ולבטל |  |
| 7:27.583 | מקצת |  |
| 7:27.584 | יעויין |  |
| 7:27.585 | שם |  |
| 7:27.586 | באורך | . |
| 7:27.587 | והנה |  |
| 7:27.588 | מזה |  |
| 7:27.589 | תבין |  |
| 7:27.590 | שלא |  |
| 7:27.591 | נודע |  |
| 7:27.592 | לספר |  |
| 7:27.593 | הזהר |  |
| 7:27.594 | ומקור |  |
| 7:27.595 | שממנו |  |
| 7:27.596 | יצא |  |
| 7:27.597 | ונתפשט | . |
| 7:27.598 | כשאר |  |
| 7:27.599 | חיבורי |  |
| 7:27.600 | רז"ל |  |
| 7:27.601 | המפורסמים |  |
| 7:27.602 | מיום |  |
| 7:27.603 | התגלותם | . |
| 7:27.604 | המשנה |  |
| 7:27.605 | והתלמוד |  |
| 7:27.606 | בבלי |  |
| 7:27.607 | וירושלמי |  |
| 7:27.608 | ותוספתא |  |
| 7:27.609 | וספרא |  |
| 7:27.610 | וספרי |  |
| 7:27.611 | שכולם |  |
| 7:27.612 | ידוע |  |
| 7:27.613 | ומפורסם |  |
| 7:27.614 | מקורן |  |
| 7:27.615 | ומי |  |
| 7:27.616 | הוא |  |
| 7:27.617 | מחברן | . |
| 7:27.618 | אבל |  |
| 7:27.619 | ספר |  |
| 7:27.620 | הזהר |  |
| 7:27.621 | המסית |  |
| 7:27.622 | זה |  |
| 7:27.623 | אומר |  |
| 7:27.624 | בכה |  |
| 7:27.625 | וזה |  |
| 7:27.626 | אומר |  |
| 7:27.627 | בכה | . |
| 7:27.628 | וקורא |  |
| 7:27.629 | אני |  |
| 7:27.630 | עליו |  |
| 7:27.631 | מאמר |  |
| 7:27.632 | רז"ל |  |
| 7:27.633 | בירושלמי |  |
| 7:27.634 | שבת |  |
| 7:27.635 | פ | ' |
| 7:27.636 | ר"א |  |
| 7:27.637 | דמילה |  |
| 7:27.638 | ופסחים |  |
| 7:27.639 | פרק |  |
| 7:27.640 | אלו |  |
| 7:27.641 | דברים |  |
| 7:27.642 | שאמרו |  |
| 7:27.643 | שם |  |
| 7:27.644 | רבי |  |
| 7:27.645 | זעירא |  |
| 7:27.646 | בשם |  |
| 7:27.647 | רבי |  |
| 7:27.648 | אלעזר |  |
| 7:27.649 | כל |  |
| 7:27.650 | תורה |  |
| 7:27.651 | שאין |  |
| 7:27.652 | לה |  |
| 7:27.653 | בית |  |
| 7:27.654 | אב |  |
| 7:27.655 | אינה |  |
| 7:27.656 | תורה |  |
| 7:27.657 | עכ"ל | . |
| 7:27.658 | ואל |  |
| 7:27.659 | תפנה |  |
| 7:27.660 | למתק |  |
| 7:27.661 | שפתיו |  |
| 7:27.662 | במה |  |
| 7:27.663 | שהוא |  |
| 7:27.664 | מלהיב |  |
| 7:27.665 | את |  |
| 7:27.666 | לב |  |
| 7:27.667 | קוראיו |  |
| 7:27.668 | ומעוררם |  |
| 7:27.669 | ליראת |  |
| 7:27.670 | שמים |  |
| 7:27.671 | ולהשתדל |  |
| 7:27.672 | בעסק |  |
| 7:27.673 | תורתו |  |
| 7:27.674 | החדשה |  |
| 7:27.675 | שהוא |  |
| 7:27.676 | קורא |  |
| 7:27.677 | לה |  |
| 7:27.678 | בת |  |
| 7:27.679 | מלך | . |
| 7:27.680 | לפי |  |
| 7:27.681 | שאין |  |
| 7:27.682 | כוונתו |  |
| 7:27.683 | רק |  |
| 7:27.684 | להביא |  |
| 7:27.685 | את |  |
| 7:27.686 | הקורא |  |
| 7:27.687 | לאמונה |  |
| 7:27.688 | נכרית |  |
| 7:27.689 | להאמין |  |
| 7:27.690 | באלהים |  |
| 7:27.691 | רבים |  |
| 7:27.692 | משפחות |  |
| 7:27.693 | משפחות | . |
| 7:27.694 | משפחת |  |
| 7:27.695 | האלוהות |  |
| 7:27.696 | שבעולם |  |
| 7:27.697 | האצילות | , |
| 7:27.698 | שהם |  |
| 7:27.699 | אדם |  |
| 7:27.700 | קדמון | , |
| 7:27.701 | ועתיק |  |
| 7:27.702 | יומין | , |
| 7:27.703 | ואריך |  |
| 7:27.704 | אנפין | , |
| 7:27.705 | ואבא |  |
| 7:27.706 | ואימא | , |
| 7:27.707 | וזעיר |  |
| 7:27.708 | אנפין |  |
| 7:27.709 | ונוקביה | . |
| 7:27.710 | וכנגדן |  |
| 7:27.711 | בעולם |  |
| 7:27.712 | הבריאה | , |
| 7:27.713 | וכמוהם |  |
| 7:27.714 | בעולם |  |
| 7:27.715 | היצירה | , |
| 7:27.716 | וכן |  |
| 7:27.717 | בעולם |  |
| 7:27.718 | העשייה | , |
| 7:27.719 | ואת |  |
| 7:27.720 | כולם |  |
| 7:27.721 | הם |  |
| 7:27.722 | מקבלים |  |
| 7:27.723 | אותם |  |
| 7:27.724 | לאלוהות |  |
| 7:27.725 | ומשבחים |  |
| 7:27.726 | ומפארים |  |
| 7:27.727 | אותם |  |
| 7:27.728 | בשעת |  |
| 7:27.729 | התפלה |  |
| 7:27.730 | באמרם |  |
| 7:27.731 | אין |  |
| 7:27.732 | כאלהינו |  |
| 7:27.733 | וכו | ' |
| 7:27.734 | כמו |  |
| 7:27.735 | שכתבתי |  |
| 7:27.736 | לעיל |  |
| 7:27.737 | סי | ' |
| 7:27.738 | ע"ט |  |
| 7:27.739 | בשם |  |
| 7:27.740 | סידור |  |
| 7:27.741 | חסד |  |
| 7:27.742 | לאברהם | . |
| 7:27.743 | דודאי |  |
| 7:27.744 | דרושים |  |
| 7:27.745 | כאלה |  |
| 7:27.746 | אסור |  |
| 7:27.747 | לשמען |  |
| 7:27.748 | וכל |  |
| 7:27.749 | שכן |  |
| 7:27.750 | לקרותן |  |
| 7:27.751 | ולהאמין |  |
| 7:27.752 | בהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

