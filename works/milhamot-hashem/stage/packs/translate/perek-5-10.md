# Translation pack: milhamot-hashem perek-5-10 (Milchamot ha-Shem 5:16–5:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-10`.

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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 5:16

Hebrew (Wikisource): ובספר הברית האריך עוד במאמר עשרים פ' ט"ו בשם הרבה מקובלים (להפוך קערה על פיה) שהמאמין ועובד לאלהים עליון רם על כל רמים אשר לא נשתנה אחר הבריאה להשתתף ולהצטרף עם ברואיו שהם נאצלים ומשתלשלים ממנו לפי דעתם, אלא כדעת רבותינו בעלי התלמוד ואנשי כנסת הגדולה מתקני התפלות והברכות אשר לא ראו את הספר הנכרי הזהר, והגאונים הר' סעדיה גאון ור' יהודה הלוי ובעל חובת הלבבות ובעל העיקרים, והרוקח, וספר מצוות גדול ור' לוי בן גרשם, ור' אברהם בן עזרא ורבים וגדולים מחכמי ישראל שקיבלו איש מפי איש עד משה רבינו ע"ה ועד אברהם אבינו הקבלה האמתית, הרי יצא מן הכלל וכפר בעיקר הוא השם הנכבד והנורא וח"ו נאמר עליו (דברים כ"ט י"ט): "ומחה ה' את שמו מתחת השמים" עיין שם ועיין לעיל סי' כ"ט.

| id | word | punctuation after |
|---|---|---|
| 5:16.1 | ובספר |  |
| 5:16.2 | הברית |  |
| 5:16.3 | האריך |  |
| 5:16.4 | עוד |  |
| 5:16.5 | במאמר |  |
| 5:16.6 | עשרים |  |
| 5:16.7 | פ | ' |
| 5:16.8 | ט"ו |  |
| 5:16.9 | בשם |  |
| 5:16.10 | הרבה |  |
| 5:16.11 | מקובלים |  |
| 5:16.12 | להפוך |  |
| 5:16.13 | קערה |  |
| 5:16.14 | על |  |
| 5:16.15 | פיה | ) |
| 5:16.16 | שהמאמין |  |
| 5:16.17 | ועובד |  |
| 5:16.18 | לאלהים |  |
| 5:16.19 | עליון |  |
| 5:16.20 | רם |  |
| 5:16.21 | על |  |
| 5:16.22 | כל |  |
| 5:16.23 | רמים |  |
| 5:16.24 | אשר |  |
| 5:16.25 | לא |  |
| 5:16.26 | נשתנה |  |
| 5:16.27 | אחר |  |
| 5:16.28 | הבריאה |  |
| 5:16.29 | להשתתף |  |
| 5:16.30 | ולהצטרף |  |
| 5:16.31 | עם |  |
| 5:16.32 | ברואיו |  |
| 5:16.33 | שהם |  |
| 5:16.34 | נאצלים |  |
| 5:16.35 | ומשתלשלים |  |
| 5:16.36 | ממנו |  |
| 5:16.37 | לפי |  |
| 5:16.38 | דעתם | , |
| 5:16.39 | אלא |  |
| 5:16.40 | כדעת |  |
| 5:16.41 | רבותינו |  |
| 5:16.42 | בעלי |  |
| 5:16.43 | התלמוד |  |
| 5:16.44 | ואנשי |  |
| 5:16.45 | כנסת |  |
| 5:16.46 | הגדולה |  |
| 5:16.47 | מתקני |  |
| 5:16.48 | התפלות |  |
| 5:16.49 | והברכות |  |
| 5:16.50 | אשר |  |
| 5:16.51 | לא |  |
| 5:16.52 | ראו |  |
| 5:16.53 | את |  |
| 5:16.54 | הספר |  |
| 5:16.55 | הנכרי |  |
| 5:16.56 | הזהר | , |
| 5:16.57 | והגאונים |  |
| 5:16.58 | הר | ' |
| 5:16.59 | סעדיה |  |
| 5:16.60 | גאון |  |
| 5:16.61 | ור | ' |
| 5:16.62 | יהודה |  |
| 5:16.63 | הלוי |  |
| 5:16.64 | ובעל |  |
| 5:16.65 | חובת |  |
| 5:16.66 | הלבבות |  |
| 5:16.67 | ובעל |  |
| 5:16.68 | העיקרים | , |
| 5:16.69 | והרוקח | , |
| 5:16.70 | וספר |  |
| 5:16.71 | מצוות |  |
| 5:16.72 | גדול |  |
| 5:16.73 | ור | ' |
| 5:16.74 | לוי |  |
| 5:16.75 | בן |  |
| 5:16.76 | גרשם | , |
| 5:16.77 | ור | ' |
| 5:16.78 | אברהם |  |
| 5:16.79 | בן |  |
| 5:16.80 | עזרא |  |
| 5:16.81 | ורבים |  |
| 5:16.82 | וגדולים |  |
| 5:16.83 | מחכמי |  |
| 5:16.84 | ישראל |  |
| 5:16.85 | שקיבלו |  |
| 5:16.86 | איש |  |
| 5:16.87 | מפי |  |
| 5:16.88 | איש |  |
| 5:16.89 | עד |  |
| 5:16.90 | משה |  |
| 5:16.91 | רבינו |  |
| 5:16.92 | ע"ה |  |
| 5:16.93 | ועד |  |
| 5:16.94 | אברהם |  |
| 5:16.95 | אבינו |  |
| 5:16.96 | הקבלה |  |
| 5:16.97 | האמתית | , |
| 5:16.98 | הרי |  |
| 5:16.99 | יצא |  |
| 5:16.100 | מן |  |
| 5:16.101 | הכלל |  |
| 5:16.102 | וכפר |  |
| 5:16.103 | בעיקר |  |
| 5:16.104 | הוא |  |
| 5:16.105 | השם |  |
| 5:16.106 | הנכבד |  |
| 5:16.107 | והנורא |  |
| 5:16.108 | וח"ו |  |
| 5:16.109 | נאמר |  |
| 5:16.110 | עליו |  |
| 5:16.111 | דברים |  |
| 5:16.112 | כ"ט |  |
| 5:16.113 | י"ט | ): |
| 5:16.114 | ומחה |  |
| 5:16.115 | ה | ' |
| 5:16.116 | את |  |
| 5:16.117 | שמו |  |
| 5:16.118 | מתחת |  |
| 5:16.119 | השמים | " |
| 5:16.120 | עיין |  |
| 5:16.121 | שם |  |
| 5:16.122 | ועיין |  |
| 5:16.123 | לעיל |  |
| 5:16.124 | סי | ' |
| 5:16.125 | כ"ט | . |


## Milchamot ha-Shem 5:17

Hebrew (Wikisource): מח) ראה תראה קורא נעים איך מלאם לבם לדבר רז"ל על סרה ועל גדולי חכמי ישראל הנזכרים לעיל. ועקרו ארבעה עיקרים מיסודי ועיקרי תורתינו הקדושה האמורים ביגדל אלהים חי, כמבואר בספר הכוונות שר' האר"י לא היה אומר יגדל אלהים חי וישתבח ע"ש במחברת הקדש דף כ"ח ע"ב. והרב חמדת ימים שהוא מתלמידי האר"י ביאר כוונתו בסוף פרק ו' מתקוני שבת וז"ל: "אחר שהעדנו ביכולו בחידוש העולם נהגו לומר פיוט יגדל אלהים בכל שבת. ונכון להיות כל איש שורר אותו בביתו בכוונה זו, אף שאמרו על האר"י שלא היה חפץ בו, והיטב חרה לו על הד' בתים הראשונים שאינם מכוונים על דרך האמת. ובזהר פרשת תרומה אמר שירות שאית בהון קבלה אמתית ולכן אין ראוי לאמרם. אכן מהנו אדון עולם ואילך אין למנוע מלאמרם וכן אני נוהג עם החברים, ולפעמים היה אומר מורי אדון עולם במקומו", עד כאן לשון חמדת ימים.

| id | word | punctuation after |
|---|---|---|
| 5:17.1 | מח | ) |
| 5:17.2 | ראה |  |
| 5:17.3 | תראה |  |
| 5:17.4 | קורא |  |
| 5:17.5 | נעים |  |
| 5:17.6 | איך |  |
| 5:17.7 | מלאם |  |
| 5:17.8 | לבם |  |
| 5:17.9 | לדבר |  |
| 5:17.10 | רז"ל |  |
| 5:17.11 | על |  |
| 5:17.12 | סרה |  |
| 5:17.13 | ועל |  |
| 5:17.14 | גדולי |  |
| 5:17.15 | חכמי |  |
| 5:17.16 | ישראל |  |
| 5:17.17 | הנזכרים |  |
| 5:17.18 | לעיל | . |
| 5:17.19 | ועקרו |  |
| 5:17.20 | ארבעה |  |
| 5:17.21 | עיקרים |  |
| 5:17.22 | מיסודי |  |
| 5:17.23 | ועיקרי |  |
| 5:17.24 | תורתינו |  |
| 5:17.25 | הקדושה |  |
| 5:17.26 | האמורים |  |
| 5:17.27 | ביגדל |  |
| 5:17.28 | אלהים |  |
| 5:17.29 | חי | , |
| 5:17.30 | כמבואר |  |
| 5:17.31 | בספר |  |
| 5:17.32 | הכוונות |  |
| 5:17.33 | שר | ' |
| 5:17.34 | האר"י |  |
| 5:17.35 | לא |  |
| 5:17.36 | היה |  |
| 5:17.37 | אומר |  |
| 5:17.38 | יגדל |  |
| 5:17.39 | אלהים |  |
| 5:17.40 | חי |  |
| 5:17.41 | וישתבח |  |
| 5:17.42 | ע"ש |  |
| 5:17.43 | במחברת |  |
| 5:17.44 | הקדש |  |
| 5:17.45 | דף |  |
| 5:17.46 | כ"ח |  |
| 5:17.47 | ע"ב | . |
| 5:17.48 | והרב |  |
| 5:17.49 | חמדת |  |
| 5:17.50 | ימים |  |
| 5:17.51 | שהוא |  |
| 5:17.52 | מתלמידי |  |
| 5:17.53 | האר"י |  |
| 5:17.54 | ביאר |  |
| 5:17.55 | כוונתו |  |
| 5:17.56 | בסוף |  |
| 5:17.57 | פרק |  |
| 5:17.58 | ו | ' |
| 5:17.59 | מתקוני |  |
| 5:17.60 | שבת |  |
| 5:17.61 | וז"ל | : |
| 5:17.62 | אחר |  |
| 5:17.63 | שהעדנו |  |
| 5:17.64 | ביכולו |  |
| 5:17.65 | בחידוש |  |
| 5:17.66 | העולם |  |
| 5:17.67 | נהגו |  |
| 5:17.68 | לומר |  |
| 5:17.69 | פיוט |  |
| 5:17.70 | יגדל |  |
| 5:17.71 | אלהים |  |
| 5:17.72 | בכל |  |
| 5:17.73 | שבת | . |
| 5:17.74 | ונכון |  |
| 5:17.75 | להיות |  |
| 5:17.76 | כל |  |
| 5:17.77 | איש |  |
| 5:17.78 | שורר |  |
| 5:17.79 | אותו |  |
| 5:17.80 | בביתו |  |
| 5:17.81 | בכוונה |  |
| 5:17.82 | זו | , |
| 5:17.83 | אף |  |
| 5:17.84 | שאמרו |  |
| 5:17.85 | על |  |
| 5:17.86 | האר"י |  |
| 5:17.87 | שלא |  |
| 5:17.88 | היה |  |
| 5:17.89 | חפץ |  |
| 5:17.90 | בו | , |
| 5:17.91 | והיטב |  |
| 5:17.92 | חרה |  |
| 5:17.93 | לו |  |
| 5:17.94 | על |  |
| 5:17.95 | הד | ' |
| 5:17.96 | בתים |  |
| 5:17.97 | הראשונים |  |
| 5:17.98 | שאינם |  |
| 5:17.99 | מכוונים |  |
| 5:17.100 | על |  |
| 5:17.101 | דרך |  |
| 5:17.102 | האמת | . |
| 5:17.103 | ובזהר |  |
| 5:17.104 | פרשת |  |
| 5:17.105 | תרומה |  |
| 5:17.106 | אמר |  |
| 5:17.107 | שירות |  |
| 5:17.108 | שאית |  |
| 5:17.109 | בהון |  |
| 5:17.110 | קבלה |  |
| 5:17.111 | אמתית |  |
| 5:17.112 | ולכן |  |
| 5:17.113 | אין |  |
| 5:17.114 | ראוי |  |
| 5:17.115 | לאמרם | . |
| 5:17.116 | אכן |  |
| 5:17.117 | מהנו |  |
| 5:17.118 | אדון |  |
| 5:17.119 | עולם |  |
| 5:17.120 | ואילך |  |
| 5:17.121 | אין |  |
| 5:17.122 | למנוע |  |
| 5:17.123 | מלאמרם |  |
| 5:17.124 | וכן |  |
| 5:17.125 | אני |  |
| 5:17.126 | נוהג |  |
| 5:17.127 | עם |  |
| 5:17.128 | החברים | , |
| 5:17.129 | ולפעמים |  |
| 5:17.130 | היה |  |
| 5:17.131 | אומר |  |
| 5:17.132 | מורי |  |
| 5:17.133 | אדון |  |
| 5:17.134 | עולם |  |
| 5:17.135 | במקומו | ", |
| 5:17.136 | עד |  |
| 5:17.137 | כאן |  |
| 5:17.138 | לשון |  |
| 5:17.139 | חמדת |  |
| 5:17.140 | ימים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

