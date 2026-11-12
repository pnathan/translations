# Review pack: milhamot-hashem perek-3-06 (Milchamot ha-Shem 3:10–3:11)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-06`.

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
- **המשרף** (pun on Metzaref ha-Emunah) → "the Mesharef", a note on the pun.
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
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
- **מורא** → "fear".
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **מרגל לרגל** → "from festival to festival", not "foot".
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **כמתלהמים** → "ravers" (Prov 26:18 in `apparatus`).
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


## Milchamot ha-Shem 3:10

Hebrew (Wikisource): י) ומעתה בעזר השי"ת נשוב לעיקר שאילתנו ששאלנו לך להודיענו על פי שיטת המקובלים, למי אנו קורין בתפלותנו ומברכים כל ברכותנו ועבודתנו וממי אנו מבקשים סליחה וכפרה. והיתה תשובתך בדברי נרגן כמתלהמים טענו חטים והודה לו בשעורים והכחשת את הידוע ומפורסם בספריהם לכל, כנשבע על האיש שהוא אשה, ועל האשה שהיא איש, ועל עמוד של שיש שהוא זהב. וחרה אפך על ספר מצרף האמונה, והלעגת עליו ובזית אותו, וקראת אותו ספר המשרף וכו' ובזה גלית לכל שאתה הוא המבזה את החכמים, ועל שאלותינו כיון שזעיר אנפין לפי דבריהם הוא אלוהינו, והוא הזן אותנו ומפרנסינו, ואותו נעבוד, אם כן כל העולמות שלמעלה מעולם האצילות, מי זן אותם ומפרנסם? אם זורק הוא שפעם ופרנסתם מלמטה למעלה, והם עובדים ומשתחוים לו, או שמא אלהיהם הזנם ומפרנסם, הוא עלה הראשונה, ולו הם עובדים, ואנחנו עובדים לזעיר אנפין הזנינו ומפרנסינו? ואם כן אלהי העליונים אינו אלהי התחתונים ח"ו? וזהו מקום הסתירה אשר בקשנו ממך להודיענו ולהורותנו, ועל זה באה תשובתך שאין הנדון דומה לראיה וכי מביאין ראיה לסתור הלכה למשה מסיני? ושחכמי הקבלה אומרים כך קיבלו ממשה.

| id | word | punctuation after |
|---|---|---|
| 3:10.1 | י | ) |
| 3:10.2 | ומעתה |  |
| 3:10.3 | בעזר |  |
| 3:10.4 | השי"ת |  |
| 3:10.5 | נשוב |  |
| 3:10.6 | לעיקר |  |
| 3:10.7 | שאילתנו |  |
| 3:10.8 | ששאלנו |  |
| 3:10.9 | לך |  |
| 3:10.10 | להודיענו |  |
| 3:10.11 | על |  |
| 3:10.12 | פי |  |
| 3:10.13 | שיטת |  |
| 3:10.14 | המקובלים | , |
| 3:10.15 | למי |  |
| 3:10.16 | אנו |  |
| 3:10.17 | קורין |  |
| 3:10.18 | בתפלותנו |  |
| 3:10.19 | ומברכים |  |
| 3:10.20 | כל |  |
| 3:10.21 | ברכותנו |  |
| 3:10.22 | ועבודתנו |  |
| 3:10.23 | וממי |  |
| 3:10.24 | אנו |  |
| 3:10.25 | מבקשים |  |
| 3:10.26 | סליחה |  |
| 3:10.27 | וכפרה | . |
| 3:10.28 | והיתה |  |
| 3:10.29 | תשובתך |  |
| 3:10.30 | בדברי |  |
| 3:10.31 | נרגן |  |
| 3:10.32 | כמתלהמים |  |
| 3:10.33 | טענו |  |
| 3:10.34 | חטים |  |
| 3:10.35 | והודה |  |
| 3:10.36 | לו |  |
| 3:10.37 | בשעורים |  |
| 3:10.38 | והכחשת |  |
| 3:10.39 | את |  |
| 3:10.40 | הידוע |  |
| 3:10.41 | ומפורסם |  |
| 3:10.42 | בספריהם |  |
| 3:10.43 | לכל | , |
| 3:10.44 | כנשבע |  |
| 3:10.45 | על |  |
| 3:10.46 | האיש |  |
| 3:10.47 | שהוא |  |
| 3:10.48 | אשה | , |
| 3:10.49 | ועל |  |
| 3:10.50 | האשה |  |
| 3:10.51 | שהיא |  |
| 3:10.52 | איש | , |
| 3:10.53 | ועל |  |
| 3:10.54 | עמוד |  |
| 3:10.55 | של |  |
| 3:10.56 | שיש |  |
| 3:10.57 | שהוא |  |
| 3:10.58 | זהב | . |
| 3:10.59 | וחרה |  |
| 3:10.60 | אפך |  |
| 3:10.61 | על |  |
| 3:10.62 | ספר |  |
| 3:10.63 | מצרף |  |
| 3:10.64 | האמונה | , |
| 3:10.65 | והלעגת |  |
| 3:10.66 | עליו |  |
| 3:10.67 | ובזית |  |
| 3:10.68 | אותו | , |
| 3:10.69 | וקראת |  |
| 3:10.70 | אותו |  |
| 3:10.71 | ספר |  |
| 3:10.72 | המשרף |  |
| 3:10.73 | וכו | ' |
| 3:10.74 | ובזה |  |
| 3:10.75 | גלית |  |
| 3:10.76 | לכל |  |
| 3:10.77 | שאתה |  |
| 3:10.78 | הוא |  |
| 3:10.79 | המבזה |  |
| 3:10.80 | את |  |
| 3:10.81 | החכמים | , |
| 3:10.82 | ועל |  |
| 3:10.83 | שאלותינו |  |
| 3:10.84 | כיון |  |
| 3:10.85 | שזעיר |  |
| 3:10.86 | אנפין |  |
| 3:10.87 | לפי |  |
| 3:10.88 | דבריהם |  |
| 3:10.89 | הוא |  |
| 3:10.90 | אלוהינו | , |
| 3:10.91 | והוא |  |
| 3:10.92 | הזן |  |
| 3:10.93 | אותנו |  |
| 3:10.94 | ומפרנסינו | , |
| 3:10.95 | ואותו |  |
| 3:10.96 | נעבוד | , |
| 3:10.97 | אם |  |
| 3:10.98 | כן |  |
| 3:10.99 | כל |  |
| 3:10.100 | העולמות |  |
| 3:10.101 | שלמעלה |  |
| 3:10.102 | מעולם |  |
| 3:10.103 | האצילות | , |
| 3:10.104 | מי |  |
| 3:10.105 | זן |  |
| 3:10.106 | אותם |  |
| 3:10.107 | ומפרנסם | ? |
| 3:10.108 | אם |  |
| 3:10.109 | זורק |  |
| 3:10.110 | הוא |  |
| 3:10.111 | שפעם |  |
| 3:10.112 | ופרנסתם |  |
| 3:10.113 | מלמטה |  |
| 3:10.114 | למעלה | , |
| 3:10.115 | והם |  |
| 3:10.116 | עובדים |  |
| 3:10.117 | ומשתחוים |  |
| 3:10.118 | לו | , |
| 3:10.119 | או |  |
| 3:10.120 | שמא |  |
| 3:10.121 | אלהיהם |  |
| 3:10.122 | הזנם |  |
| 3:10.123 | ומפרנסם | , |
| 3:10.124 | הוא |  |
| 3:10.125 | עלה |  |
| 3:10.126 | הראשונה | , |
| 3:10.127 | ולו |  |
| 3:10.128 | הם |  |
| 3:10.129 | עובדים | , |
| 3:10.130 | ואנחנו |  |
| 3:10.131 | עובדים |  |
| 3:10.132 | לזעיר |  |
| 3:10.133 | אנפין |  |
| 3:10.134 | הזנינו |  |
| 3:10.135 | ומפרנסינו | ? |
| 3:10.136 | ואם |  |
| 3:10.137 | כן |  |
| 3:10.138 | אלהי |  |
| 3:10.139 | העליונים |  |
| 3:10.140 | אינו |  |
| 3:10.141 | אלהי |  |
| 3:10.142 | התחתונים |  |
| 3:10.143 | ח"ו | ? |
| 3:10.144 | וזהו |  |
| 3:10.145 | מקום |  |
| 3:10.146 | הסתירה |  |
| 3:10.147 | אשר |  |
| 3:10.148 | בקשנו |  |
| 3:10.149 | ממך |  |
| 3:10.150 | להודיענו |  |
| 3:10.151 | ולהורותנו | , |
| 3:10.152 | ועל |  |
| 3:10.153 | זה |  |
| 3:10.154 | באה |  |
| 3:10.155 | תשובתך |  |
| 3:10.156 | שאין |  |
| 3:10.157 | הנדון |  |
| 3:10.158 | דומה |  |
| 3:10.159 | לראיה |  |
| 3:10.160 | וכי |  |
| 3:10.161 | מביאין |  |
| 3:10.162 | ראיה |  |
| 3:10.163 | לסתור |  |
| 3:10.164 | הלכה |  |
| 3:10.165 | למשה |  |
| 3:10.166 | מסיני | ? |
| 3:10.167 | ושחכמי |  |
| 3:10.168 | הקבלה |  |
| 3:10.169 | אומרים |  |
| 3:10.170 | כך |  |
| 3:10.171 | קיבלו |  |
| 3:10.172 | ממשה | . |


### Translation under review (3:10)

English: 10) And now, with the help of the Name, may He be blessed, let us return to the main point of our question, which we asked you to inform us according to the system of the Kabbalists, to whom do we call upon in our prayers, and bless all our blessings and our worship, and from whom do we seek forgiveness and atonement. And your answer was in the words of a grumbler, like the flustered who claim wheat and concede barley to him, and you deny what is known and famous in their books to all, like one who swears about the man that he is a woman, and about the woman that she is a man, and about a pillar of marble that it is gold. And your anger burned against the book Metzaref ha-Emunah, and you mocked it and scorned it, and you called it the book the Mesharef etc., and thereby you revealed to all that you are the one who scorns the sages; And as to our questions: since Ze'ir Anpin, according to their words, is our God, and he feeds us and sustains us, and Him we worship; If so, all the worlds above the World of Emanation, who feeds and sustains them? If he casts their flow down, and their sustenance is from below upward, and they serve and bow down to him, or perhaps their gods fed and sustained them, he being the First Cause, and to him they serve, and we serve Ze'ir Anpin, who feeds us and sustains us? And if so, is the God of the upper ones not the God of the lower ones, God forbid? And this is the point of contradiction which we asked you to inform us and to instruct us, and to this your answer came: that the case is not like proof, and do they bring proof to overturn a halakha given to Moses at Sinai? And the sages of Kabbalah say thus: they received it from Moses.

Literal: 10) and-now in-the-help-of the-Name-may-He-be-blessed we-return to-the-point-of our-question which we-asked you to-inform-us according-to the-system-of the-Kabbalists, to-whom we call-upon in-our-prayers and-bless all our-blessings and-our-worship and-from-whom we seek forgiveness and-atonement. and-was your-answer in-the-words-of a-grumbler like-the-flustered who-claim wheat and-concede to-him barley, and-you-deny (object) what-is-known and-famous in-their-books to-all, like-one-who-swears about the-man that-he-is a-woman, and-about the-woman that-she-is a-man, and-about a-pillar of marble that-it-is gold. and-burned your-anger against the-book Metzaref ha-Emunah, and-you-mocked at-it and-scorned it, and-you-called it the-book the-Mesharef etc., and-thereby you-revealed to-all that-you (emphatic) the-one-who-scorns (object) the-sages; and-as-to our-questions since Ze'ir Anpin according-to their-words is our-God, and-he feeds-us and-sustains-us, and-Him we-worship; if-so all the-worlds that-are-above from-the-World-of Emanation who-feeds-them and-sustains-them? if he-casts their-flow, and-their-sustenance from-below upward, and-they serve and-bow-down to-him, or perhaps their-gods fed-them and-sustained-them, he (is) the-cause, the-first, and-to-him they serve, and-we serve to-Ze'ir Anpin, who-feeds-us and-sustains-us? and-if-so the-God-of the-upper-ones is-not the-God-of the-lower-ones, God-forbid? and-this the-point-of contradiction which we-asked of-you to-inform-us and-to-instruct-us, and-to this your-answer came: that-not the-case is-like proof, and-do they-bring proof to-overturn a-halakha given-to Moses at-Sinai? and-the-sages-of Kabbalah say thus: they-received it from-Moses.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:10.1 | 10) |  |  | printed paragraph number, kept as printed |
| 3:10.2 | and now |  |  |  |
| 3:10.3 | with the help of |  |  |  |
| 3:10.4 | the Name, may He be blessed |  |  | the Name, may He be blessed: the project's fixed form for the Hebrew abbreviation |
| 3:10.5 | let us return |  |  |  |
| 3:10.6 | to the main point of |  |  |  |
| 3:10.7 | our question |  |  |  |
| 3:10.8 | which we asked |  |  |  |
| 3:10.9 | you |  |  |  |
| 3:10.10 | to inform us |  |  |  |
| 3:10.11 | according to |  |  |  |
| 3:10.12 | (lit. on the mouth of) |  |  |  |
| 3:10.13 | the system of |  |  |  |
| 3:10.14 | the Kabbalists |  |  |  |
| 3:10.15 | to whom |  |  |  |
| 3:10.16 | do we |  |  |  |
| 3:10.17 | call upon |  |  |  |
| 3:10.18 | in our prayers |  |  |  |
| 3:10.19 | and bless |  |  |  |
| 3:10.20 | all |  |  |  |
| 3:10.21 | our blessings |  |  |  |
| 3:10.22 | and our worship |  |  |  |
| 3:10.23 | and from whom |  |  |  |
| 3:10.24 | do we |  |  |  |
| 3:10.25 | seek |  |  |  |
| 3:10.26 | forgiveness |  |  |  |
| 3:10.27 | and atonement |  |  |  |
| 3:10.28 | and was |  |  |  |
| 3:10.29 | your answer |  |  |  |
| 3:10.30 | in the words of |  |  |  |
| 3:10.31 | a grumbler |  |  |  |
| 3:10.32 | like the flustered |  |  | the Hebrew word is obscure; 'like the flustered' is the sense the context gives |
| 3:10.33 | who claim |  |  |  |
| 3:10.34 | wheat |  |  |  |
| 3:10.35 | and concede |  |  |  |
| 3:10.36 | to him |  |  |  |
| 3:10.37 | barley |  |  | the image 'claim wheat, concede barley' is an idiom for a disputant who gives way on the small point; the English keeps the image |
| 3:10.38 | and you deny |  |  |  |
| 3:10.39 | (object marker) |  |  |  |
| 3:10.40 | what is known |  |  |  |
| 3:10.41 | and famous |  |  |  |
| 3:10.42 | in their books |  |  |  |
| 3:10.43 | to all |  |  |  |
| 3:10.44 | like one who swears |  |  |  |
| 3:10.45 | about |  |  |  |
| 3:10.46 | the man |  |  |  |
| 3:10.47 | that he is |  |  |  |
| 3:10.48 | a woman |  |  |  |
| 3:10.49 | and about |  |  |  |
| 3:10.50 | the woman |  |  |  |
| 3:10.51 | that she is |  |  |  |
| 3:10.52 | a man |  |  |  |
| 3:10.53 | and about |  |  |  |
| 3:10.54 | a pillar |  |  |  |
| 3:10.55 | of |  |  |  |
| 3:10.56 | marble |  |  |  |
| 3:10.57 | that it is |  |  |  |
| 3:10.58 | gold |  |  |  |
| 3:10.59 | and burned |  |  |  |
| 3:10.60 | your anger |  |  |  |
| 3:10.61 | against |  |  |  |
| 3:10.62 | the book |  |  |  |
| 3:10.63 | Metzaref |  |  | Metzaref ha-Emunah: 'the refining of faith', a book title, transliterated |
| 3:10.64 | ha-Emunah |  |  |  |
| 3:10.65 | and you mocked |  |  |  |
| 3:10.66 | at it |  |  |  |
| 3:10.67 | and scorned |  |  |  |
| 3:10.68 | it |  |  |  |
| 3:10.69 | and you called |  |  |  |
| 3:10.70 | it |  |  |  |
| 3:10.71 | the book |  |  |  |
| 3:10.72 | the Mesharef |  |  | the Mesharef: a pun on Metzaref ha-Emunah; 'mesharef' (a burner) for 'metzaref' (a refiner) |
| 3:10.73 | etc. |  |  | etc.: the author's own mark that the quotation is cut short |
| 3:10.74 | and thereby |  |  |  |
| 3:10.75 | you revealed |  |  |  |
| 3:10.76 | to all |  |  |  |
| 3:10.77 | that you |  |  |  |
| 3:10.78 | (emphatic: you are) |  |  |  |
| 3:10.79 | the one who scorns |  |  |  |
| 3:10.80 | (object marker) |  |  |  |
| 3:10.81 | the sages |  |  |  |
| 3:10.82 | and as to |  |  |  |
| 3:10.83 | our questions |  |  |  |
| 3:10.84 | since |  |  |  |
| 3:10.85 | that Ze'ir |  |  | Ze'ir Anpin: 'Small Face', the name of a partzuf; the name is transliterated and its sense given here at its first occurrence in the unit |
| 3:10.86 | Anpin |  |  |  |
| 3:10.87 | according to |  |  |  |
| 3:10.88 | their words |  |  |  |
| 3:10.89 | is |  |  |  |
| 3:10.90 | our God |  |  | 'our God' is the author's report of the Kabbalists' view ('according to their words') |
| 3:10.91 | and he |  |  |  |
| 3:10.92 | feeds |  |  |  |
| 3:10.93 | us |  |  |  |
| 3:10.94 | and sustains us |  |  |  |
| 3:10.95 | and Him |  |  |  |
| 3:10.96 | we worship |  |  |  |
| 3:10.97 | if |  |  |  |
| 3:10.98 | so |  |  |  |
| 3:10.99 | all |  |  |  |
| 3:10.100 | the worlds |  |  |  |
| 3:10.101 | that are above |  |  |  |
| 3:10.102 | the world of |  |  |  |
| 3:10.103 | Emanation |  |  |  |
| 3:10.104 | who |  |  |  |
| 3:10.105 | feeds |  |  |  |
| 3:10.106 | them |  |  |  |
| 3:10.107 | and sustains them |  |  |  |
| 3:10.108 | if |  |  |  |
| 3:10.109 | he casts |  |  |  |
| 3:10.110 | he |  |  | he: the question opens with the pronoun, 'he casts' |
| 3:10.111 | their flow |  |  | the flow (השפע) with the suffix 'their' |
| 3:10.112 | and their sustenance |  |  |  |
| 3:10.113 | from below |  |  |  |
| 3:10.114 | upward |  |  |  |
| 3:10.115 | and they |  |  |  |
| 3:10.116 | serve |  |  |  |
| 3:10.117 | and bow down |  |  |  |
| 3:10.118 | to him |  |  |  |
| 3:10.119 | or |  |  |  |
| 3:10.120 | perhaps |  |  |  |
| 3:10.121 | their gods |  |  | plural 'their gods', in the author's voice; the alternative is 'their God' |
| 3:10.122 | fed them |  |  |  |
| 3:10.123 | and sustained them |  |  |  |
| 3:10.124 | he |  |  |  |
| 3:10.125 | cause |  |  | cause: the First Cause (עלה הראשונה) |
| 3:10.126 | first |  |  |  |
| 3:10.127 | and to him |  |  |  |
| 3:10.128 | they |  |  |  |
| 3:10.129 | serve |  |  |  |
| 3:10.130 | and we |  |  |  |
| 3:10.131 | serve |  |  |  |
| 3:10.132 | Ze'ir |  |  |  |
| 3:10.133 | Anpin |  |  |  |
| 3:10.134 | who feeds us |  |  |  |
| 3:10.135 | and sustains us |  |  |  |
| 3:10.136 | and if |  |  |  |
| 3:10.137 | so |  |  |  |
| 3:10.138 | the God of |  |  | plural construct 'elohei' with the singular verb 'eino'; 'God' by the ruling; the alternative is 'gods' |
| 3:10.139 | the upper ones |  |  |  |
| 3:10.140 | is not |  |  |  |
| 3:10.141 | the God of |  |  |  |
| 3:10.142 | the lower ones |  |  |  |
| 3:10.143 | God forbid |  |  | ח"ו: God forbid |
| 3:10.144 | and this |  |  |  |
| 3:10.145 | the point |  |  |  |
| 3:10.146 | of contradiction |  |  |  |
| 3:10.147 | which |  |  |  |
| 3:10.148 | we asked |  |  |  |
| 3:10.149 | of you |  |  |  |
| 3:10.150 | to inform us |  |  |  |
| 3:10.151 | and to instruct us |  |  |  |
| 3:10.152 | and to |  |  |  |
| 3:10.153 | this |  |  |  |
| 3:10.154 | came |  |  |  |
| 3:10.155 | your answer |  |  |  |
| 3:10.156 | that not |  |  |  |
| 3:10.157 | the case |  |  |  |
| 3:10.158 | is like |  |  |  |
| 3:10.159 | proof |  |  |  |
| 3:10.160 | and do |  |  |  |
| 3:10.161 | they bring |  |  |  |
| 3:10.162 | proof |  |  |  |
| 3:10.163 | to overturn |  |  |  |
| 3:10.164 | a halakha |  |  |  |
| 3:10.165 | given to Moses |  |  |  |
| 3:10.166 | at Sinai |  |  |  |
| 3:10.167 | And the sages of |  |  |  |
| 3:10.168 | Kabbalah |  |  | Kabbalah: the word for the tradition the Kabbalists claim; here it is 'the sages of Kabbalah' |
| 3:10.169 | say |  |  |  |
| 3:10.170 | thus |  |  |  |
| 3:10.171 | they received |  |  |  |
| 3:10.172 | it from Moses |  |  |  |

Phrases:
- 3:10.1 3:10.2 3:10.3 3:10.4 3:10.5 3:10.6 3:10.7 | 10) And now, with the help of the Name, may He be blessed, let us return to the main point of our question, | 10) and-now in-the-help-of the-Name-may-He-be-blessed we-return to-the-point-of our-question
- 3:10.8 3:10.9 3:10.10 | which we asked you to inform us | which we-asked you to-inform-us
- 3:10.11 3:10.12 3:10.13 3:10.14 3:10.15 3:10.16 3:10.17 3:10.18 3:10.19 3:10.20 3:10.21 3:10.22 3:10.23 3:10.24 3:10.25 3:10.26 3:10.27 | according to the system of the Kabbalists, to whom do we call upon in our prayers, and bless all our blessings and our worship, and from whom do we seek forgiveness and atonement. | according-to the-system-of the-Kabbalists, to-whom we call-upon in-our-prayers and-bless all our-blessings and-our-worship and-from-whom we seek forgiveness and-atonement.
- 3:10.28 3:10.29 3:10.30 3:10.31 3:10.32 3:10.33 3:10.34 3:10.35 3:10.36 3:10.37 | And your answer was in the words of a grumbler, like the flustered who claim wheat and concede barley to him, | and-was your-answer in-the-words-of a-grumbler like-the-flustered who-claim wheat and-concede to-him barley,
- 3:10.38 3:10.39 3:10.40 3:10.41 3:10.42 3:10.43 | and you deny what is known and famous in their books to all, | and-you-deny (object) what-is-known and-famous in-their-books to-all,
- 3:10.44 3:10.45 3:10.46 3:10.47 3:10.48 | like one who swears about the man that he is a woman, | like-one-who-swears about the-man that-he-is a-woman,
- 3:10.49 3:10.50 3:10.51 3:10.52 | and about the woman that she is a man, | and-about the-woman that-she-is a-man,
- 3:10.53 3:10.54 3:10.55 3:10.56 3:10.57 3:10.58 | and about a pillar of marble that it is gold. | and-about a-pillar of marble that-it-is gold.
- 3:10.59 3:10.60 3:10.61 3:10.62 3:10.63 3:10.64 | And your anger burned against the book Metzaref ha-Emunah, | and-burned your-anger against the-book Metzaref ha-Emunah,
- 3:10.65 3:10.66 3:10.67 3:10.68 | and you mocked it and scorned it, | and-you-mocked at-it and-scorned it,
- 3:10.69 3:10.70 3:10.71 3:10.72 3:10.73 | and you called it the book the Mesharef etc., | and-you-called it the-book the-Mesharef etc.,
- 3:10.74 3:10.75 3:10.76 3:10.77 3:10.78 3:10.79 3:10.80 3:10.81 | and thereby you revealed to all that you are the one who scorns the sages; | and-thereby you-revealed to-all that-you (emphatic) the-one-who-scorns (object) the-sages;
- 3:10.82 3:10.83 3:10.84 3:10.85 3:10.86 3:10.87 3:10.88 3:10.89 3:10.90 3:10.91 3:10.92 3:10.93 3:10.94 3:10.95 3:10.96 | And as to our questions: since Ze'ir Anpin, according to their words, is our God, and he feeds us and sustains us, and Him we worship; | and-as-to our-questions since Ze'ir Anpin according-to their-words is our-God, and-he feeds-us and-sustains-us, and-Him we-worship;
- 3:10.97 3:10.98 3:10.99 3:10.100 3:10.101 3:10.102 3:10.103 3:10.104 3:10.105 3:10.106 3:10.107 | If so, all the worlds above the World of Emanation, who feeds and sustains them? | if-so all the-worlds that-are-above from-the-World-of Emanation who-feeds-them and-sustains-them?
- 3:10.108 3:10.109 3:10.110 3:10.111 3:10.112 3:10.113 3:10.114 3:10.115 3:10.116 3:10.117 3:10.118 | If he casts their flow down, and their sustenance is from below upward, and they serve and bow down to him, | if he-casts their-flow, and-their-sustenance from-below upward, and-they serve and-bow-down to-him,
- 3:10.119 3:10.120 3:10.121 3:10.122 3:10.123 | or perhaps their gods fed and sustained them, | or perhaps their-gods fed-them and-sustained-them,
- 3:10.124 3:10.125 3:10.126 | he being the First Cause, | he (is) the-cause, the-first,
- 3:10.127 3:10.128 3:10.129 3:10.130 3:10.131 3:10.132 3:10.133 3:10.134 3:10.135 | and to him they serve, and we serve Ze'ir Anpin, who feeds us and sustains us? | and-to-him they serve, and-we serve to-Ze'ir Anpin, who-feeds-us and-sustains-us?
- 3:10.136 3:10.137 3:10.138 3:10.139 3:10.140 3:10.141 3:10.142 3:10.143 | And if so, is the God of the upper ones not the God of the lower ones, God forbid? | and-if-so the-God-of the-upper-ones is-not the-God-of the-lower-ones, God-forbid?
- 3:10.144 3:10.145 3:10.146 3:10.147 3:10.148 3:10.149 3:10.150 3:10.151 | And this is the point of contradiction which we asked you to inform us and to instruct us, | and-this the-point-of contradiction which we-asked of-you to-inform-us and-to-instruct-us,
- 3:10.152 3:10.153 3:10.154 3:10.155 3:10.156 3:10.157 3:10.158 3:10.159 | and to this your answer came: that the case is not like proof, | and-to this your-answer came: that-not the-case is-like proof,
- 3:10.160 3:10.161 3:10.162 3:10.163 3:10.164 3:10.165 3:10.166 | and do they bring proof to overturn a halakha given to Moses at Sinai? | and-do they-bring proof to-overturn a-halakha given-to Moses at-Sinai?
- 3:10.167 3:10.168 3:10.169 3:10.170 3:10.171 3:10.172 | And the sages of Kabbalah say thus: they received it from Moses. | and-the-sages-of Kabbalah say thus: they-received it from-Moses.

Grammar notes:
- והיתה (3:10.28) is a vav-prefixed perfect in the consecutive sense, 'and your answer was'; the subject 'your answer' follows the verb.
- The verbs of 3:10.59-72 (וחרה, והלעגת, ובזית, וקראת) are a series of waw-prefixed perfects, each a separate verb of the sentence, so the English keeps them as a series.
- ואם כן (3:10.136) opens a question. אלהי העליונים is a plural construct with the singular verb אינו; the English reads 'the God of' (see the token note on 3:10.138).
- ועל שאלותינו (3:10.82) has no verb of its own: it is a topic phrase, and the main clause is the question that begins אם כן (3:10.97).
- The simile כמתלהמים טענו חטים והודה לו בשעורים (3:10.32-37) is kept by its image; the English gives its sense, not a word-for-word match.

Commentary rulings:
- apparatus: Base text: the Wikisource transcription of the paragraph; not yet collated against the 1931 scan (Q2 default: flagged until collated).

## Milchamot ha-Shem 3:11

Hebrew (Wikisource): יא) מאד אנו תמהים ומתפלאים על דבריך ועל תשובתך זאת, וכי הלכה למשה מסיני להאמין באלוהות רבים? ולעבוד אלהים אחרים מבלעדי ה' אלהינו הראשון וצור לכל? והוא סבת כל הסבות כולם, ובמרמה את בא בתשובותיך אלה לאנשים עזי לב אשר התפארת עליהם והראית להם דבריך אלה לאמר. ראו כי סתרתי את דבריהם, ודחיתי את ראיותיהם בקש, וחוזר קונדריסך על כל המסובין, ואחר אשר קילסוך ושבחוך שאתה חריף ושנון הבאת אותה לנו. והנה כולה הטעאות כאלו לא ידעת על מה תסוב שאלתינו, ועל טעותך הכוזבת הלזו נשוב עוד, כי ידוע ומפורסם בכללי הפוסקים (כנה"ג ורדב"ז ושה"ז) שכל מקום שהקבלה חולקת על הפסק הלך אחרי הפסק, והכלל הזה מפורסם בישראל. ואם דבריך שהקבלה הלזו היא הלכה למשה מסיני, וכי תדחה הלכה למשה מפני סברת איזה תנא או אמורא, וכ"ש פוסק אחרון (עיין לקמן סי' ט"ו) ועליך להביא ראיה ולהודיענו היכן מצינו זה ששום תנא או אמורא יעיז פניו לחלוק ולהורות היפך הלכה המקובלת ממשה רבי' ע"ה. הלא שוטה או רשע וגס רוח יקרא? ובכמה מקומות בתלמוד מרגלא בפומיהו דרבנן לומר אם הלכה נקבל, ואם לדין יש תשובה.

| id | word | punctuation after |
|---|---|---|
| 3:11.1 | יא | ) |
| 3:11.2 | מאד |  |
| 3:11.3 | אנו |  |
| 3:11.4 | תמהים |  |
| 3:11.5 | ומתפלאים |  |
| 3:11.6 | על |  |
| 3:11.7 | דבריך |  |
| 3:11.8 | ועל |  |
| 3:11.9 | תשובתך |  |
| 3:11.10 | זאת | , |
| 3:11.11 | וכי |  |
| 3:11.12 | הלכה |  |
| 3:11.13 | למשה |  |
| 3:11.14 | מסיני |  |
| 3:11.15 | להאמין |  |
| 3:11.16 | באלוהות |  |
| 3:11.17 | רבים | ? |
| 3:11.18 | ולעבוד |  |
| 3:11.19 | אלהים |  |
| 3:11.20 | אחרים |  |
| 3:11.21 | מבלעדי |  |
| 3:11.22 | ה | ' |
| 3:11.23 | אלהינו |  |
| 3:11.24 | הראשון |  |
| 3:11.25 | וצור |  |
| 3:11.26 | לכל | ? |
| 3:11.27 | והוא |  |
| 3:11.28 | סבת |  |
| 3:11.29 | כל |  |
| 3:11.30 | הסבות |  |
| 3:11.31 | כולם | , |
| 3:11.32 | ובמרמה |  |
| 3:11.33 | את |  |
| 3:11.34 | בא |  |
| 3:11.35 | בתשובותיך |  |
| 3:11.36 | אלה |  |
| 3:11.37 | לאנשים |  |
| 3:11.38 | עזי |  |
| 3:11.39 | לב |  |
| 3:11.40 | אשר |  |
| 3:11.41 | התפארת |  |
| 3:11.42 | עליהם |  |
| 3:11.43 | והראית |  |
| 3:11.44 | להם |  |
| 3:11.45 | דבריך |  |
| 3:11.46 | אלה |  |
| 3:11.47 | לאמר | . |
| 3:11.48 | ראו |  |
| 3:11.49 | כי |  |
| 3:11.50 | סתרתי |  |
| 3:11.51 | את |  |
| 3:11.52 | דבריהם | , |
| 3:11.53 | ודחיתי |  |
| 3:11.54 | את |  |
| 3:11.55 | ראיותיהם |  |
| 3:11.56 | בקש | , |
| 3:11.57 | וחוזר |  |
| 3:11.58 | קונדריסך |  |
| 3:11.59 | על |  |
| 3:11.60 | כל |  |
| 3:11.61 | המסובין | , |
| 3:11.62 | ואחר |  |
| 3:11.63 | אשר |  |
| 3:11.64 | קילסוך |  |
| 3:11.65 | ושבחוך |  |
| 3:11.66 | שאתה |  |
| 3:11.67 | חריף |  |
| 3:11.68 | ושנון |  |
| 3:11.69 | הבאת |  |
| 3:11.70 | אותה |  |
| 3:11.71 | לנו | . |
| 3:11.72 | והנה |  |
| 3:11.73 | כולה |  |
| 3:11.74 | הטעאות |  |
| 3:11.75 | כאלו |  |
| 3:11.76 | לא |  |
| 3:11.77 | ידעת |  |
| 3:11.78 | על |  |
| 3:11.79 | מה |  |
| 3:11.80 | תסוב |  |
| 3:11.81 | שאלתינו | , |
| 3:11.82 | ועל |  |
| 3:11.83 | טעותך |  |
| 3:11.84 | הכוזבת |  |
| 3:11.85 | הלזו |  |
| 3:11.86 | נשוב |  |
| 3:11.87 | עוד | , |
| 3:11.88 | כי |  |
| 3:11.89 | ידוע |  |
| 3:11.90 | ומפורסם |  |
| 3:11.91 | בכללי |  |
| 3:11.92 | הפוסקים |  |
| 3:11.93 | כנה"ג |  |
| 3:11.94 | ורדב"ז |  |
| 3:11.95 | ושה"ז | ) |
| 3:11.96 | שכל |  |
| 3:11.97 | מקום |  |
| 3:11.98 | שהקבלה |  |
| 3:11.99 | חולקת |  |
| 3:11.100 | על |  |
| 3:11.101 | הפסק |  |
| 3:11.102 | הלך |  |
| 3:11.103 | אחרי |  |
| 3:11.104 | הפסק | , |
| 3:11.105 | והכלל |  |
| 3:11.106 | הזה |  |
| 3:11.107 | מפורסם |  |
| 3:11.108 | בישראל | . |
| 3:11.109 | ואם |  |
| 3:11.110 | דבריך |  |
| 3:11.111 | שהקבלה |  |
| 3:11.112 | הלזו |  |
| 3:11.113 | היא |  |
| 3:11.114 | הלכה |  |
| 3:11.115 | למשה |  |
| 3:11.116 | מסיני | , |
| 3:11.117 | וכי |  |
| 3:11.118 | תדחה |  |
| 3:11.119 | הלכה |  |
| 3:11.120 | למשה |  |
| 3:11.121 | מפני |  |
| 3:11.122 | סברת |  |
| 3:11.123 | איזה |  |
| 3:11.124 | תנא |  |
| 3:11.125 | או |  |
| 3:11.126 | אמורא | , |
| 3:11.127 | וכ"ש |  |
| 3:11.128 | פוסק |  |
| 3:11.129 | אחרון |  |
| 3:11.130 | עיין |  |
| 3:11.131 | לקמן |  |
| 3:11.132 | סי | ' |
| 3:11.133 | ט"ו | ) |
| 3:11.134 | ועליך |  |
| 3:11.135 | להביא |  |
| 3:11.136 | ראיה |  |
| 3:11.137 | ולהודיענו |  |
| 3:11.138 | היכן |  |
| 3:11.139 | מצינו |  |
| 3:11.140 | זה |  |
| 3:11.141 | ששום |  |
| 3:11.142 | תנא |  |
| 3:11.143 | או |  |
| 3:11.144 | אמורא |  |
| 3:11.145 | יעיז |  |
| 3:11.146 | פניו |  |
| 3:11.147 | לחלוק |  |
| 3:11.148 | ולהורות |  |
| 3:11.149 | היפך |  |
| 3:11.150 | הלכה |  |
| 3:11.151 | המקובלת |  |
| 3:11.152 | ממשה |  |
| 3:11.153 | רבי | ' |
| 3:11.154 | ע"ה | . |
| 3:11.155 | הלא |  |
| 3:11.156 | שוטה |  |
| 3:11.157 | או |  |
| 3:11.158 | רשע |  |
| 3:11.159 | וגס |  |
| 3:11.160 | רוח |  |
| 3:11.161 | יקרא | ? |
| 3:11.162 | ובכמה |  |
| 3:11.163 | מקומות |  |
| 3:11.164 | בתלמוד |  |
| 3:11.165 | מרגלא |  |
| 3:11.166 | בפומיהו |  |
| 3:11.167 | דרבנן |  |
| 3:11.168 | לומר |  |
| 3:11.169 | אם |  |
| 3:11.170 | הלכה |  |
| 3:11.171 | נקבל | , |
| 3:11.172 | ואם |  |
| 3:11.173 | לדין |  |
| 3:11.174 | יש |  |
| 3:11.175 | תשובה | . |


### Translation under review (3:11)

English: 11) We are very astonished and amazed at your words and at this reply of yours, and is it a halakha given to Moses at Sinai to believe in many divinities? and to worship other gods apart from the LORD our God, the First, and Rock of all? And He is the cause of all causes, all of them, and with deceit you came in your replies to these stubborn-hearted people, whom you boasted over, and showed them these words of yours, saying: "See, I have overturned their words, and rejected their proofs in haste, and your pamphlet goes round to all the diners, and after they praised and lauded you, that you are sharp and keen, you brought it to us." And behold, all of it is deceptions, as if you did not know what our question turns on, and to this false error of yours we shall return again, since it is known and famous in the rules of the decisors (Kenesset ha-Gedolah, the Radbaz, and Sh"Z), that wherever the Kabbalah disagrees with the ruling, one goes after the ruling, and this rule is famous in Israel. And if your words are that this Kabbalah is a halakha given to Moses at Sinai, and do you reject a halakha given to Moses on account of the reasoning of some tanna or amora, and all the more so a later decisor (see below, §15), and you must bring a proof and inform us where we have found this, that any tanna or amora would be so bold to dispute and to rule against the halakha received from Moses, our Rabbi, peace be upon him. Would he not be called a fool, or a wicked and arrogant man? And in many places in the Talmud it is a saying in the mouths of the Rabbis to say: "If it is a halakha, we accept it, and if it is a judgment, there is an answer."

Literal: 11) very we are-astonished and-amazed at your-words and-at your-reply, this, and-is-it a-halakha given-to Moses from-Sinai to-believe in-many-divinities? and-to-worship other gods apart-from the-LORD our-God, the-First and-Rock of-all? and-He the-cause of-all the-causes, all-of-them, and-with-deceit you came in-your-replies to these people stubborn of-heart, whom you-boasted over-them, and-showed-them your-words these, saying: "see that I-overturned their-words, and-I-rejected their-proofs in-haste, and-your-pamphlet goes-round to-all the-diners, and-after that they-praised-you and-lauded-you, that you sharp and-keen, you-brought it to-us." and-behold all-of-it deceptions, as-if you knew-not about what turns our-question, and-to your-error the-false this, we-shall-return again, since known and-famous in-the-rules of-the-decisors (Kenesset ha-Gedolah, the-Radbaz, and-Sh"Z), that wherever the-Kabbalah disagrees with the-ruling, one-goes after the-ruling, and-this rule is-famous in-Israel. and-if your-words that this-Kabbalah is a-halakha given-to Moses from-Sinai, and-do you-reject a-halakha given-to Moses on-account-of the-reasoning-of some tanna or amora, and-all-the-more-so a-decisor later (see below, section 15), and-you-must bring proof and-inform-us where we-found this, that-any tanna or amora would-be-so-bold (to-brazen his-face) to-dispute and-to-rule against the-halakha received from-Moses, our-Rabbi, peace-be-upon-him. would-he-not be-called a-fool, or a-wicked-man and-arrogant? and-in many places in-the-Talmud it-is a-saying in-their-mouths of-the-Rabbis to-say: "if halakha we-accept, and-if (as-to) a-judgment there-is an-answer."

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:11.1 | 11) |  |  | printed paragraph number, kept as printed |
| 3:11.2 | very |  |  |  |
| 3:11.3 | we |  |  |  |
| 3:11.4 | are astonished |  |  |  |
| 3:11.5 | and amazed |  |  |  |
| 3:11.6 | at |  |  |  |
| 3:11.7 | your words |  |  |  |
| 3:11.8 | and at |  |  |  |
| 3:11.9 | your reply |  |  |  |
| 3:11.10 | this |  |  |  |
| 3:11.11 | and is it |  |  | rhetorical question, as the Hebrew has it (וכי) |
| 3:11.12 | a halakha |  |  |  |
| 3:11.13 | given to Moses |  |  |  |
| 3:11.14 | at Sinai |  |  |  |
| 3:11.15 | to believe |  |  |  |
| 3:11.16 | in |  |  |  |
| 3:11.17 | many divinities |  |  | divinities: the plural is counted ('many'), so the English is plural |
| 3:11.18 | and to worship |  |  |  |
| 3:11.19 | gods |  |  | other gods: the wording of Exod 20:3 (see apparatus) |
| 3:11.20 | other |  |  |  |
| 3:11.21 | apart from |  |  |  |
| 3:11.22 | the LORD |  |  | the LORD: the Tetragrammaton, written ה' in the text |
| 3:11.23 | our God |  |  |  |
| 3:11.24 | the First |  |  | the First: an allusion to Isa 44:6 (see apparatus) |
| 3:11.25 | and Rock |  |  |  |
| 3:11.26 | of all |  |  |  |
| 3:11.27 | and He |  |  |  |
| 3:11.28 | cause |  |  |  |
| 3:11.29 | of all |  |  |  |
| 3:11.30 | causes |  |  |  |
| 3:11.31 | all of them |  |  |  |
| 3:11.32 | and with deceit |  |  |  |
| 3:11.33 | you |  |  | את here is the second-person pronoun 'you', the subject, not the object marker |
| 3:11.34 | came |  |  |  |
| 3:11.35 | in your replies |  |  |  |
| 3:11.36 | these |  |  |  |
| 3:11.37 | to people |  |  |  |
| 3:11.38 | stubborn |  |  |  |
| 3:11.39 | of heart |  |  |  |
| 3:11.40 | whom |  |  |  |
| 3:11.41 | you boasted |  |  |  |
| 3:11.42 | over them |  |  |  |
| 3:11.43 | and showed |  |  |  |
| 3:11.44 | them |  |  |  |
| 3:11.45 | your words |  |  |  |
| 3:11.46 | these |  |  |  |
| 3:11.47 | saying |  |  | לאמר introduces the quotation, which is given in double marks; the quoted words keep the Hebrew's shift between 'I' and 'you' |
| 3:11.48 | See |  |  |  |
| 3:11.49 | that |  |  |  |
| 3:11.50 | I overturned |  |  |  |
| 3:11.51 | (object marker) |  |  |  |
| 3:11.52 | their words |  |  |  |
| 3:11.53 | and I rejected |  |  |  |
| 3:11.54 | (object marker) |  |  |  |
| 3:11.55 | their proofs |  |  |  |
| 3:11.56 | in haste |  |  | בקש: 'in haste' is one reading; 'in brief' is another; the sense is not settled |
| 3:11.57 | and goes round |  |  |  |
| 3:11.58 | your pamphlet |  |  | קונדריס: a small book, a pamphlet |
| 3:11.59 | to |  |  |  |
| 3:11.60 | all |  |  |  |
| 3:11.61 | the diners |  |  | המסובין: those reclining at a meal |
| 3:11.62 | and after |  |  |  |
| 3:11.63 | that |  |  |  |
| 3:11.64 | they praised you |  |  |  |
| 3:11.65 | and lauded you |  |  |  |
| 3:11.66 | that you are |  |  |  |
| 3:11.67 | sharp |  |  |  |
| 3:11.68 | and keen |  |  |  |
| 3:11.69 | you brought |  |  |  |
| 3:11.70 | it |  |  |  |
| 3:11.71 | to us |  |  |  |
| 3:11.72 | And behold |  |  |  |
| 3:11.73 | all of it |  |  |  |
| 3:11.74 | deceptions |  |  |  |
| 3:11.75 | as if |  |  |  |
| 3:11.76 | not |  |  |  |
| 3:11.77 | you knew |  |  |  |
| 3:11.78 | about |  |  |  |
| 3:11.79 | what |  |  |  |
| 3:11.80 | turns |  |  |  |
| 3:11.81 | our question |  |  |  |
| 3:11.82 | and to |  |  |  |
| 3:11.83 | your error |  |  |  |
| 3:11.84 | the false |  |  |  |
| 3:11.85 | this |  |  |  |
| 3:11.86 | we shall return |  |  |  |
| 3:11.87 | again |  |  |  |
| 3:11.88 | since |  |  |  |
| 3:11.89 | it is known |  |  |  |
| 3:11.90 | and famous |  |  |  |
| 3:11.91 | in the rules of |  |  |  |
| 3:11.92 | the decisors |  |  |  |
| 3:11.93 | Kenesset ha-Gedolah |  |  | כנה"ג: Kenesset ha-Gedolah, written out |
| 3:11.94 | the Radbaz |  |  | ורדב"ז: the Radbaz (David ibn Zimra) |
| 3:11.95 | and Sh"Z |  |  | ושה"ז: the abbreviation is not expanded here, since its expansion is not established; kept as written |
| 3:11.96 | that |  |  |  |
| 3:11.97 | wherever |  |  |  |
| 3:11.98 | the Kabbalah |  |  |  |
| 3:11.99 | disagrees |  |  |  |
| 3:11.100 | with |  |  |  |
| 3:11.101 | the ruling |  |  |  |
| 3:11.102 | one goes |  |  |  |
| 3:11.103 | after |  |  |  |
| 3:11.104 | the ruling |  |  |  |
| 3:11.105 | and this |  |  |  |
| 3:11.106 | rule |  |  |  |
| 3:11.107 | is famous |  |  |  |
| 3:11.108 | in Israel |  |  |  |
| 3:11.109 | and if |  |  |  |
| 3:11.110 | your words |  |  |  |
| 3:11.111 | that |  |  |  |
| 3:11.112 | this Kabbalah |  |  |  |
| 3:11.113 | is |  |  |  |
| 3:11.114 | a halakha |  |  |  |
| 3:11.115 | given to Moses |  |  |  |
| 3:11.116 | at Sinai |  |  |  |
| 3:11.117 | and do |  |  |  |
| 3:11.118 | you reject |  |  |  |
| 3:11.119 | a halakha |  |  |  |
| 3:11.120 | given to Moses |  |  |  |
| 3:11.121 | on account of |  |  |  |
| 3:11.122 | the reasoning of |  |  |  |
| 3:11.123 | some |  |  |  |
| 3:11.124 | tanna |  |  |  |
| 3:11.125 | or |  |  |  |
| 3:11.126 | amora |  |  |  |
| 3:11.127 | and all the more so |  |  |  |
| 3:11.128 | a decisor |  |  |  |
| 3:11.129 | later |  |  |  |
| 3:11.130 | see |  |  |  |
| 3:11.131 | below |  |  |  |
| 3:11.132 | section |  |  | סי': the author's section number, §15 |
| 3:11.133 | 15 |  |  |  |
| 3:11.134 | and you must |  |  |  |
| 3:11.135 | bring |  |  |  |
| 3:11.136 | a proof |  |  |  |
| 3:11.137 | and inform us |  |  |  |
| 3:11.138 | where |  |  |  |
| 3:11.139 | we have found |  |  |  |
| 3:11.140 | this |  |  |  |
| 3:11.141 | that any |  |  |  |
| 3:11.142 | tanna |  |  |  |
| 3:11.143 | or |  |  |  |
| 3:11.144 | amora |  |  |  |
| 3:11.145 | would be so bold |  |  |  |
| 3:11.146 | to brazen (lit. his face) |  |  | יעיז פניו: 'brazen his face', the idiom for daring |
| 3:11.147 | to dispute |  |  |  |
| 3:11.148 | and to rule |  |  |  |
| 3:11.149 | against |  |  |  |
| 3:11.150 | the halakha |  |  |  |
| 3:11.151 | received |  |  |  |
| 3:11.152 | from Moses |  |  |  |
| 3:11.153 | our Rabbi |  |  | רבי': 'our Rabbi', of Moses |
| 3:11.154 | peace be upon him |  |  | ע"ה: peace be upon him |
| 3:11.155 | would he not |  |  |  |
| 3:11.156 | a fool |  |  |  |
| 3:11.157 | or |  |  |  |
| 3:11.158 | a wicked man |  |  |  |
| 3:11.159 | and arrogant |  |  |  |
| 3:11.160 | of spirit |  |  |  |
| 3:11.161 | be called |  |  |  |
| 3:11.162 | And in many |  |  |  |
| 3:11.163 | places |  |  |  |
| 3:11.164 | in the Talmud |  |  |  |
| 3:11.165 | a saying |  |  | מרגלא בפומיהו דרבנן: a common saying in the Aramaic of the Talmud, as the author gives it |
| 3:11.166 | in their mouths |  |  |  |
| 3:11.167 | of the Rabbis |  |  |  |
| 3:11.168 | to say |  |  |  |
| 3:11.169 | if |  |  | the saying is quoted without a reference in the author's text; nothing is added |
| 3:11.170 | halakha |  |  |  |
| 3:11.171 | we accept |  |  |  |
| 3:11.172 | and if |  |  |  |
| 3:11.173 | as to a judgment |  |  |  |
| 3:11.174 | there is |  |  |  |
| 3:11.175 | an answer |  |  |  |

Phrases:
- 3:11.1 3:11.2 3:11.3 3:11.4 3:11.5 3:11.6 3:11.7 3:11.8 3:11.9 3:11.10 | 11) We are very astonished and amazed at your words and at this reply of yours, | 11) very we are-astonished and-amazed at your-words and-at your-reply, this,
- 3:11.11 3:11.12 3:11.13 3:11.14 3:11.15 3:11.16 3:11.17 | and is it a halakha given to Moses at Sinai to believe in many divinities? | and-is-it a-halakha given-to Moses from-Sinai to-believe in-many-divinities?
- 3:11.18 3:11.19 3:11.20 3:11.21 3:11.22 3:11.23 3:11.24 3:11.25 3:11.26 | and to worship other gods apart from the LORD our God, the First, and Rock of all? | and-to-worship other gods apart-from the-LORD our-God, the-First and-Rock of-all?
- 3:11.27 3:11.28 3:11.29 3:11.30 3:11.31 | And He is the cause of all causes, all of them, | and-He the-cause of-all the-causes, all-of-them,
- 3:11.32 3:11.33 3:11.34 3:11.35 3:11.36 3:11.37 3:11.38 3:11.39 3:11.40 3:11.41 3:11.42 3:11.43 3:11.44 3:11.45 3:11.46 3:11.47 | and with deceit you came in your replies to these stubborn-hearted people, whom you boasted over, and showed them these words of yours, saying: | and-with-deceit you came in-your-replies to these people stubborn of-heart, whom you-boasted over-them, and-showed-them your-words these, saying:
- 3:11.48 3:11.49 3:11.50 3:11.51 3:11.52 3:11.53 3:11.54 3:11.55 3:11.56 3:11.57 3:11.58 3:11.59 3:11.60 3:11.61 | "See, I have overturned their words, and rejected their proofs in haste, and your pamphlet goes round to all the diners, | "see that I-overturned their-words, and-I-rejected their-proofs in-haste, and-your-pamphlet goes-round to-all the-diners,
- 3:11.62 3:11.63 3:11.64 3:11.65 3:11.66 3:11.67 3:11.68 3:11.69 3:11.70 3:11.71 | and after they praised and lauded you, that you are sharp and keen, you brought it to us." | and-after that they-praised-you and-lauded-you, that you sharp and-keen, you-brought it to-us."
- 3:11.72 3:11.73 3:11.74 3:11.75 3:11.76 3:11.77 3:11.78 3:11.79 3:11.80 3:11.81 | And behold, all of it is deceptions, as if you did not know what our question turns on, | and-behold all-of-it deceptions, as-if you knew-not about what turns our-question,
- 3:11.82 3:11.83 3:11.84 3:11.85 3:11.86 3:11.87 | and to this false error of yours we shall return again, | and-to your-error the-false this, we-shall-return again,
- 3:11.88 3:11.89 3:11.90 3:11.91 3:11.92 3:11.93 3:11.94 3:11.95 | since it is known and famous in the rules of the decisors (Kenesset ha-Gedolah, the Radbaz, and Sh"Z), | since known and-famous in-the-rules of-the-decisors (Kenesset ha-Gedolah, the-Radbaz, and-Sh"Z),
- 3:11.96 3:11.97 3:11.98 3:11.99 3:11.100 3:11.101 3:11.102 3:11.103 3:11.104 3:11.105 3:11.106 3:11.107 3:11.108 | that wherever the Kabbalah disagrees with the ruling, one goes after the ruling, and this rule is famous in Israel. | that wherever the-Kabbalah disagrees with the-ruling, one-goes after the-ruling, and-this rule is-famous in-Israel.
- 3:11.109 3:11.110 3:11.111 3:11.112 3:11.113 3:11.114 3:11.115 3:11.116 | And if your words are that this Kabbalah is a halakha given to Moses at Sinai, | and-if your-words that this-Kabbalah is a-halakha given-to Moses from-Sinai,
- 3:11.117 3:11.118 3:11.119 3:11.120 3:11.121 3:11.122 3:11.123 3:11.124 3:11.125 3:11.126 3:11.127 3:11.128 3:11.129 3:11.130 3:11.131 3:11.132 3:11.133 | and do you reject a halakha given to Moses on account of the reasoning of some tanna or amora, and all the more so a later decisor (see below, §15), | and-do you-reject a-halakha given-to Moses on-account-of the-reasoning-of some tanna or amora, and-all-the-more-so a-decisor later (see below, section 15),
- 3:11.134 3:11.135 3:11.136 3:11.137 3:11.138 3:11.139 3:11.140 3:11.141 3:11.142 3:11.143 3:11.144 3:11.145 3:11.146 | and you must bring a proof and inform us where we have found this, that any tanna or amora would be so bold | and-you-must bring proof and-inform-us where we-found this, that-any tanna or amora would-be-so-bold (to-brazen his-face)
- 3:11.147 3:11.148 3:11.149 3:11.150 3:11.151 3:11.152 3:11.153 3:11.154 | to dispute and to rule against the halakha received from Moses, our Rabbi, peace be upon him. | to-dispute and-to-rule against the-halakha received from-Moses, our-Rabbi, peace-be-upon-him.
- 3:11.155 3:11.156 3:11.157 3:11.158 3:11.159 3:11.160 3:11.161 | Would he not be called a fool, or a wicked and arrogant man? | would-he-not be-called a-fool, or a-wicked-man and-arrogant?
- 3:11.162 3:11.163 3:11.164 3:11.165 3:11.166 3:11.167 3:11.168 | And in many places in the Talmud it is a saying in the mouths of the Rabbis to say: | and-in many places in-the-Talmud it-is a-saying in-their-mouths of-the-Rabbis to-say:
- 3:11.169 3:11.170 3:11.171 3:11.172 3:11.173 3:11.174 3:11.175 | "If it is a halakha, we accept it, and if it is a judgment, there is an answer." | "if halakha we-accept, and-if (as-to) a-judgment there-is an-answer."

Grammar notes:
- הלכה למשה מסיני להאמין (3:11.11-16) and ולעבוד (3:11.18) are infinitives after the question 'is it a halakha': 'to believe' and 'to worship' are its content.
- את in ובמרמה את בא (3:11.32-34) is the second-person pronoun 'you', the subject, not the object marker; the English reads 'you came'.
- ראו כי סתרתי... (3:11.48-71) is the quotation that follows לאמר (3:11.47); its shifts between 'I' and 'you' are kept as in the Hebrew.
- מרגלא בפומיהו דרבנן (3:11.165-168) is Bavli Aramaic, the frame of the saying; the saying itself (3:11.169-175) is in Hebrew and Aramaic words, quoted without a reference.

Commentary rulings:
- apparatus: Allusion to Isa 44:6 ('I am the first and I am the last, and besides Me there is no God'): 'the First' (3:11.24).
- apparatus: Allusion to Isa 44:8 ('Is there a God besides Me? There is no Rock; I know none'): 'and Rock' (3:11.25), and 'other' with מבלעדי (3:11.21).
- apparatus: Allusion to Exod 20:3 ('You shall have no other gods before Me'): 'other gods' (3:11.19-20).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

