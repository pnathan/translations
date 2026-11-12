# Translation pack: milhamot-hashem perek-6-27 (Milchamot ha-Shem 6:41–6:41)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-27.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-27`.

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
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **שלשה פנים** → "three aspects" (note: "faces").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 6:41

Hebrew (Wikisource): עט) ולא די שעשו כן לחשוב כמה אלוהות בעולם האצילות, אלא שגם בעולם הבריאה, ועולם היצירה, ועולם העשייה. המציאו ברעיונותם כל הפרצופים הנז' אצלם בעולם האצילות, ואמר שהאין סוף המתלבש בהם לאו איהו חד בהון כמו בעולם האצילות כמו שכתב הר' כסא אליהו דף כ"א ע"א. ואע"פ שאמרו שהפרצופים שבבי"ע אין האין סוף בהם א' עכ"ז הם משבחים ומפארים להם, ומקבלים עליהם אלהותם כמבואר בסידור חסד חאברהם באמרם אין כאלהינו (מלכות דעשיה) אין כאדונינו (זעיר דעשייה) אין כמלכינו (אימא דעשייה) אין כמושיענו (אבא ואריך דעשייה) מי כאלהינו (מלכות דיצירה) מי כאדונינו (זעיר כיצירה) מי כמלכינו (אימא דיצירה) מי כמושיענו (אבא ואריך דיצירה) נודה לאלהינו (מלכות דבריאה) נודה לאדונינו (סעיר דבריאה) נודה למלכינו (אימא דבריאה) נודה למושיענו (אבא ואריך אנפין דבריאה) ברוך אלהינו (מלכות דאצילות) ברוך אדונינו (זעיר דאצילות) ברוך מלכינו (אימא דאצילות) ברןך מושעינו (אבא ואריך אנפין דאצילות) אצ ההוא אלהינו (כתר דאריך אנפין) אתה הוא אדונינו (ראש עתיק יומין) אתה הוא מלכינו (אדם קדמון) אתה הוא מושיענו (אור פנימי ואור מקיף המתלבש בו) כמו שיראה המעיין כל זה בסידור חסד לאברהם. הרי לך בהדיא שהם מקבלים עליהם מלכות ואדנות ואלהות כל פרצוף ופרצוף. אפי' פרצופים שבעולמות הפירוד לפי דבריהם שהם בריאה, יצירה ועשייה. ופליאה גדולה אנו מתפלאים עליכם שאתם נלחמים עמנו בכל המלחמה הזאת לבזות ולרדוף ולירד עמנו עד לחיינו. ואנו סובלים כל חרפות ובזוים לכבוד אבינו מלכינו יחיד ומיוחד מכל האחדים, בעבור מה? לא בעבור שנקבל ונודה לאמונתכם הלזו ש]ירשנו ממנה ולא חירפנו ולא בזינו אותה? ועם כל זה אתם מסתירים ומכסים את אמונתכם שלא לפרסמה. ומראים בתשובתכם כאלו אינכם מודים בה כעין גנב שמרה נשף לאמר לא תשורני עין. וסתר פנים תשימו לה! לכן הוכרחנו להביא כל זה להוכיח את הכחשות שאתם מכחישים אותנו בתשובתכם ואומרים ומי הוא שאמר שהאין סוף מתחלק? כל זה הוצרכנו להעתיק מספרי המקובלים החדשים. ולהביא מאמרי רז"ל נגדן להוכיח טעותכם או רמאותכם שאתם מסתירים איכות אמונתכם ומכסים אותה בטלית. ומכחישים את הידוע, ואינכם מבקשים האמת. ולמה זה תבושו לומר האמת שכן קבלנו לעבוד אלוהות רבים נבראים רק שאתם רודפים אחר הכבוד המדומה. לקנות שם נגד ההמון אשר לא ידעו ולא יבינו, בחשכה יתהלכו, ולשבח את עצמכם נגדם כאלו אתם נלחמים עם ה' ועם תורתו.

| id | word | punctuation after |
|---|---|---|
| 6:41.1 | עט | ) |
| 6:41.2 | ולא |  |
| 6:41.3 | די |  |
| 6:41.4 | שעשו |  |
| 6:41.5 | כן |  |
| 6:41.6 | לחשוב |  |
| 6:41.7 | כמה |  |
| 6:41.8 | אלוהות |  |
| 6:41.9 | בעולם |  |
| 6:41.10 | האצילות | , |
| 6:41.11 | אלא |  |
| 6:41.12 | שגם |  |
| 6:41.13 | בעולם |  |
| 6:41.14 | הבריאה | , |
| 6:41.15 | ועולם |  |
| 6:41.16 | היצירה | , |
| 6:41.17 | ועולם |  |
| 6:41.18 | העשייה | . |
| 6:41.19 | המציאו |  |
| 6:41.20 | ברעיונותם |  |
| 6:41.21 | כל |  |
| 6:41.22 | הפרצופים |  |
| 6:41.23 | הנז | ' |
| 6:41.24 | אצלם |  |
| 6:41.25 | בעולם |  |
| 6:41.26 | האצילות | , |
| 6:41.27 | ואמר |  |
| 6:41.28 | שהאין |  |
| 6:41.29 | סוף |  |
| 6:41.30 | המתלבש |  |
| 6:41.31 | בהם |  |
| 6:41.32 | לאו |  |
| 6:41.33 | איהו |  |
| 6:41.34 | חד |  |
| 6:41.35 | בהון |  |
| 6:41.36 | כמו |  |
| 6:41.37 | בעולם |  |
| 6:41.38 | האצילות |  |
| 6:41.39 | כמו |  |
| 6:41.40 | שכתב |  |
| 6:41.41 | הר | ' |
| 6:41.42 | כסא |  |
| 6:41.43 | אליהו |  |
| 6:41.44 | דף |  |
| 6:41.45 | כ"א |  |
| 6:41.46 | ע"א | . |
| 6:41.47 | ואע"פ |  |
| 6:41.48 | שאמרו |  |
| 6:41.49 | שהפרצופים |  |
| 6:41.50 | שבבי"ע |  |
| 6:41.51 | אין |  |
| 6:41.52 | האין |  |
| 6:41.53 | סוף |  |
| 6:41.54 | בהם |  |
| 6:41.55 | א | ' |
| 6:41.56 | עכ"ז |  |
| 6:41.57 | הם |  |
| 6:41.58 | משבחים |  |
| 6:41.59 | ומפארים |  |
| 6:41.60 | להם | , |
| 6:41.61 | ומקבלים |  |
| 6:41.62 | עליהם |  |
| 6:41.63 | אלהותם |  |
| 6:41.64 | כמבואר |  |
| 6:41.65 | בסידור |  |
| 6:41.66 | חסד |  |
| 6:41.67 | חאברהם |  |
| 6:41.68 | באמרם |  |
| 6:41.69 | אין |  |
| 6:41.70 | כאלהינו |  |
| 6:41.71 | מלכות |  |
| 6:41.72 | דעשיה | ) |
| 6:41.73 | אין |  |
| 6:41.74 | כאדונינו |  |
| 6:41.75 | זעיר |  |
| 6:41.76 | דעשייה | ) |
| 6:41.77 | אין |  |
| 6:41.78 | כמלכינו |  |
| 6:41.79 | אימא |  |
| 6:41.80 | דעשייה | ) |
| 6:41.81 | אין |  |
| 6:41.82 | כמושיענו |  |
| 6:41.83 | אבא |  |
| 6:41.84 | ואריך |  |
| 6:41.85 | דעשייה | ) |
| 6:41.86 | מי |  |
| 6:41.87 | כאלהינו |  |
| 6:41.88 | מלכות |  |
| 6:41.89 | דיצירה | ) |
| 6:41.90 | מי |  |
| 6:41.91 | כאדונינו |  |
| 6:41.92 | זעיר |  |
| 6:41.93 | כיצירה | ) |
| 6:41.94 | מי |  |
| 6:41.95 | כמלכינו |  |
| 6:41.96 | אימא |  |
| 6:41.97 | דיצירה | ) |
| 6:41.98 | מי |  |
| 6:41.99 | כמושיענו |  |
| 6:41.100 | אבא |  |
| 6:41.101 | ואריך |  |
| 6:41.102 | דיצירה | ) |
| 6:41.103 | נודה |  |
| 6:41.104 | לאלהינו |  |
| 6:41.105 | מלכות |  |
| 6:41.106 | דבריאה | ) |
| 6:41.107 | נודה |  |
| 6:41.108 | לאדונינו |  |
| 6:41.109 | סעיר |  |
| 6:41.110 | דבריאה | ) |
| 6:41.111 | נודה |  |
| 6:41.112 | למלכינו |  |
| 6:41.113 | אימא |  |
| 6:41.114 | דבריאה | ) |
| 6:41.115 | נודה |  |
| 6:41.116 | למושיענו |  |
| 6:41.117 | אבא |  |
| 6:41.118 | ואריך |  |
| 6:41.119 | אנפין |  |
| 6:41.120 | דבריאה | ) |
| 6:41.121 | ברוך |  |
| 6:41.122 | אלהינו |  |
| 6:41.123 | מלכות |  |
| 6:41.124 | דאצילות | ) |
| 6:41.125 | ברוך |  |
| 6:41.126 | אדונינו |  |
| 6:41.127 | זעיר |  |
| 6:41.128 | דאצילות | ) |
| 6:41.129 | ברוך |  |
| 6:41.130 | מלכינו |  |
| 6:41.131 | אימא |  |
| 6:41.132 | דאצילות | ) |
| 6:41.133 | ברןך |  |
| 6:41.134 | מושעינו |  |
| 6:41.135 | אבא |  |
| 6:41.136 | ואריך |  |
| 6:41.137 | אנפין |  |
| 6:41.138 | דאצילות | ) |
| 6:41.139 | אצ |  |
| 6:41.140 | ההוא |  |
| 6:41.141 | אלהינו |  |
| 6:41.142 | כתר |  |
| 6:41.143 | דאריך |  |
| 6:41.144 | אנפין | ) |
| 6:41.145 | אתה |  |
| 6:41.146 | הוא |  |
| 6:41.147 | אדונינו |  |
| 6:41.148 | ראש |  |
| 6:41.149 | עתיק |  |
| 6:41.150 | יומין | ) |
| 6:41.151 | אתה |  |
| 6:41.152 | הוא |  |
| 6:41.153 | מלכינו |  |
| 6:41.154 | אדם |  |
| 6:41.155 | קדמון | ) |
| 6:41.156 | אתה |  |
| 6:41.157 | הוא |  |
| 6:41.158 | מושיענו |  |
| 6:41.159 | אור |  |
| 6:41.160 | פנימי |  |
| 6:41.161 | ואור |  |
| 6:41.162 | מקיף |  |
| 6:41.163 | המתלבש |  |
| 6:41.164 | בו | ) |
| 6:41.165 | כמו |  |
| 6:41.166 | שיראה |  |
| 6:41.167 | המעיין |  |
| 6:41.168 | כל |  |
| 6:41.169 | זה |  |
| 6:41.170 | בסידור |  |
| 6:41.171 | חסד |  |
| 6:41.172 | לאברהם | . |
| 6:41.173 | הרי |  |
| 6:41.174 | לך |  |
| 6:41.175 | בהדיא |  |
| 6:41.176 | שהם |  |
| 6:41.177 | מקבלים |  |
| 6:41.178 | עליהם |  |
| 6:41.179 | מלכות |  |
| 6:41.180 | ואדנות |  |
| 6:41.181 | ואלהות |  |
| 6:41.182 | כל |  |
| 6:41.183 | פרצוף |  |
| 6:41.184 | ופרצוף | . |
| 6:41.185 | אפי | ' |
| 6:41.186 | פרצופים |  |
| 6:41.187 | שבעולמות |  |
| 6:41.188 | הפירוד |  |
| 6:41.189 | לפי |  |
| 6:41.190 | דבריהם |  |
| 6:41.191 | שהם |  |
| 6:41.192 | בריאה | , |
| 6:41.193 | יצירה |  |
| 6:41.194 | ועשייה | . |
| 6:41.195 | ופליאה |  |
| 6:41.196 | גדולה |  |
| 6:41.197 | אנו |  |
| 6:41.198 | מתפלאים |  |
| 6:41.199 | עליכם |  |
| 6:41.200 | שאתם |  |
| 6:41.201 | נלחמים |  |
| 6:41.202 | עמנו |  |
| 6:41.203 | בכל |  |
| 6:41.204 | המלחמה |  |
| 6:41.205 | הזאת |  |
| 6:41.206 | לבזות |  |
| 6:41.207 | ולרדוף |  |
| 6:41.208 | ולירד |  |
| 6:41.209 | עמנו |  |
| 6:41.210 | עד |  |
| 6:41.211 | לחיינו | . |
| 6:41.212 | ואנו |  |
| 6:41.213 | סובלים |  |
| 6:41.214 | כל |  |
| 6:41.215 | חרפות |  |
| 6:41.216 | ובזוים |  |
| 6:41.217 | לכבוד |  |
| 6:41.218 | אבינו |  |
| 6:41.219 | מלכינו |  |
| 6:41.220 | יחיד |  |
| 6:41.221 | ומיוחד |  |
| 6:41.222 | מכל |  |
| 6:41.223 | האחדים | , |
| 6:41.224 | בעבור |  |
| 6:41.225 | מה | ? |
| 6:41.226 | לא |  |
| 6:41.227 | בעבור |  |
| 6:41.228 | שנקבל |  |
| 6:41.229 | ונודה |  |
| 6:41.230 | לאמונתכם |  |
| 6:41.231 | הלזו |  |
| 6:41.232 | ש]ירשנו |  |
| 6:41.233 | ממנה |  |
| 6:41.234 | ולא |  |
| 6:41.235 | חירפנו |  |
| 6:41.236 | ולא |  |
| 6:41.237 | בזינו |  |
| 6:41.238 | אותה | ? |
| 6:41.239 | ועם |  |
| 6:41.240 | כל |  |
| 6:41.241 | זה |  |
| 6:41.242 | אתם |  |
| 6:41.243 | מסתירים |  |
| 6:41.244 | ומכסים |  |
| 6:41.245 | את |  |
| 6:41.246 | אמונתכם |  |
| 6:41.247 | שלא |  |
| 6:41.248 | לפרסמה | . |
| 6:41.249 | ומראים |  |
| 6:41.250 | בתשובתכם |  |
| 6:41.251 | כאלו |  |
| 6:41.252 | אינכם |  |
| 6:41.253 | מודים |  |
| 6:41.254 | בה |  |
| 6:41.255 | כעין |  |
| 6:41.256 | גנב |  |
| 6:41.257 | שמרה |  |
| 6:41.258 | נשף |  |
| 6:41.259 | לאמר |  |
| 6:41.260 | לא |  |
| 6:41.261 | תשורני |  |
| 6:41.262 | עין | . |
| 6:41.263 | וסתר |  |
| 6:41.264 | פנים |  |
| 6:41.265 | תשימו |  |
| 6:41.266 | לה | ! |
| 6:41.267 | לכן |  |
| 6:41.268 | הוכרחנו |  |
| 6:41.269 | להביא |  |
| 6:41.270 | כל |  |
| 6:41.271 | זה |  |
| 6:41.272 | להוכיח |  |
| 6:41.273 | את |  |
| 6:41.274 | הכחשות |  |
| 6:41.275 | שאתם |  |
| 6:41.276 | מכחישים |  |
| 6:41.277 | אותנו |  |
| 6:41.278 | בתשובתכם |  |
| 6:41.279 | ואומרים |  |
| 6:41.280 | ומי |  |
| 6:41.281 | הוא |  |
| 6:41.282 | שאמר |  |
| 6:41.283 | שהאין |  |
| 6:41.284 | סוף |  |
| 6:41.285 | מתחלק | ? |
| 6:41.286 | כל |  |
| 6:41.287 | זה |  |
| 6:41.288 | הוצרכנו |  |
| 6:41.289 | להעתיק |  |
| 6:41.290 | מספרי |  |
| 6:41.291 | המקובלים |  |
| 6:41.292 | החדשים | . |
| 6:41.293 | ולהביא |  |
| 6:41.294 | מאמרי |  |
| 6:41.295 | רז"ל |  |
| 6:41.296 | נגדן |  |
| 6:41.297 | להוכיח |  |
| 6:41.298 | טעותכם |  |
| 6:41.299 | או |  |
| 6:41.300 | רמאותכם |  |
| 6:41.301 | שאתם |  |
| 6:41.302 | מסתירים |  |
| 6:41.303 | איכות |  |
| 6:41.304 | אמונתכם |  |
| 6:41.305 | ומכסים |  |
| 6:41.306 | אותה |  |
| 6:41.307 | בטלית | . |
| 6:41.308 | ומכחישים |  |
| 6:41.309 | את |  |
| 6:41.310 | הידוע | , |
| 6:41.311 | ואינכם |  |
| 6:41.312 | מבקשים |  |
| 6:41.313 | האמת | . |
| 6:41.314 | ולמה |  |
| 6:41.315 | זה |  |
| 6:41.316 | תבושו |  |
| 6:41.317 | לומר |  |
| 6:41.318 | האמת |  |
| 6:41.319 | שכן |  |
| 6:41.320 | קבלנו |  |
| 6:41.321 | לעבוד |  |
| 6:41.322 | אלוהות |  |
| 6:41.323 | רבים |  |
| 6:41.324 | נבראים |  |
| 6:41.325 | רק |  |
| 6:41.326 | שאתם |  |
| 6:41.327 | רודפים |  |
| 6:41.328 | אחר |  |
| 6:41.329 | הכבוד |  |
| 6:41.330 | המדומה | . |
| 6:41.331 | לקנות |  |
| 6:41.332 | שם |  |
| 6:41.333 | נגד |  |
| 6:41.334 | ההמון |  |
| 6:41.335 | אשר |  |
| 6:41.336 | לא |  |
| 6:41.337 | ידעו |  |
| 6:41.338 | ולא |  |
| 6:41.339 | יבינו | , |
| 6:41.340 | בחשכה |  |
| 6:41.341 | יתהלכו | , |
| 6:41.342 | ולשבח |  |
| 6:41.343 | את |  |
| 6:41.344 | עצמכם |  |
| 6:41.345 | נגדם |  |
| 6:41.346 | כאלו |  |
| 6:41.347 | אתם |  |
| 6:41.348 | נלחמים |  |
| 6:41.349 | עם |  |
| 6:41.350 | ה | ' |
| 6:41.351 | ועם |  |
| 6:41.352 | תורתו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

