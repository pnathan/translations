# Review pack: milhamot-hashem perek-2-04 (Milchamot ha-Shem 2:9–2:9)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-2-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-04`.

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
- **The names of the partzufim**: transliterated, the literal sense in a note
  first in a unit: Ze'ir Anpin (Small Face; ז"א → "Ze'ir Anpin"; זעיר alone →
  "Ze'ir"); Arikh Anpin (Long Face); Abba; Imma; Nukva (נוקביה → "his Nukva");
  Atik; עתיקא קדישא → "Atika Kadisha"; עתיקא סתימא → "Atika Setima"; עתיק
  יומין → "Atik Yomin" (Dan 7:9: "the Ancient of Days"). Translated: סבא דסבין
  → "the Elder of Elders"; עלת העלות → "the Cause of causes"; הסבה הראשונה →
  "the First Cause"; כתר עליון → "Keter Elyon (the Supreme Crown)". An
  argument from a name's meaning goes in the note; the name stays.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 2:9

Hebrew (Wikisource): אחלי אדוני הודיעני ג"כ, מי מנהיג כל העולמות הנעלמים אשר לא דברו בהם המקובלים ומי זנם ומפרנסם. אם זעיר אנפין זה באצילות הוא הזנם ומפרנסם, ולו הם עובדים ומשתחוים, או האין סוף הוא אלהיהם ולו לבדו עובדים, ואנחנו לז"א עובדים, כי הוא הזן אותנו והוא אלהינו כמו שאמרו המקובלים. ונמצא לפי זה, כי אין סוף הוא אלהי העליונים, וז"א אלהי התחתונים, והרי כאן שתי רשויות ח"ו. ובהדיא אמרו מקצתן: את כל מעשה ה' הגדול מכלל דאיכא קטן. ורבותינו האמתיים אמרו הובא במנורת המאור פ' קמ"ג שבשעת מתן תורה קרע להם שבעה רקיעים ופתח ז' ארצות ואמר להם לישראל: בינו בעליונים ובינו בתחתונים ודעו שאין אלוה מבלעדי לא בעליונים ולא בתחתונים, אנכי ה' אלהיך. ובמדרש רבה פרשת יתרו אמרו: אנכי ה' אלהיך. אמר ר' אבהו משל למלך בשר ודם מולך, יש לו אב או אח או בן. אמר הקב"ה: אני איני כן, ראשון ואני אחרון ומבלעדי אין אלהים. אני ראשון שאין לי אב. ואני אחרון שאין לי בן. ומבלעדי אין אלהים שאין לי אח. עכ"ל המדרש וכו' הר' עץ יוסף ז"ל ר"ל שמלך ב"ו אין מלכותו שלימה לפי שיש לו אב וצריך לנהוג בו כבוד ומשתתף עמו במלכות. או יש לו אח שיש לו חלק בכבוד מלכות. או יש לו בן שנוהג שררה עמו, ואחריו יורש עצר. ופי' אין לו אב שהוא הסבה הראשונה. ואין לו אח כי הוא אחד ואין שני לו. ואין לו בן שלא נמשכו העלולים ממנו בטבע כבן מן האב, כמו שחשבו קצת (היא דעת המקובלים שאמרו שז"א בן לאו"א והוא אב לעלולים שאחריו וכו' ויש לו אחות או אחיות נוקבין דיליה) כי לפי הדעת ההיא אין היכולת בידו לשנות שום דבר מן הדברים ואין מלכותו שלימה. אבל הקב"ה אמר: אנכי ה' אלהיך, כי זה כחו להוציאם מארץ מצרים לפי שמלכותו שלימה ולית דימחי בידיה עכ"ל:

| id | word | punctuation after |
|---|---|---|
| 2:9.1 | אחלי |  |
| 2:9.2 | אדוני |  |
| 2:9.3 | הודיעני |  |
| 2:9.4 | ג"כ | , |
| 2:9.5 | מי |  |
| 2:9.6 | מנהיג |  |
| 2:9.7 | כל |  |
| 2:9.8 | העולמות |  |
| 2:9.9 | הנעלמים |  |
| 2:9.10 | אשר |  |
| 2:9.11 | לא |  |
| 2:9.12 | דברו |  |
| 2:9.13 | בהם |  |
| 2:9.14 | המקובלים |  |
| 2:9.15 | ומי |  |
| 2:9.16 | זנם |  |
| 2:9.17 | ומפרנסם | . |
| 2:9.18 | אם |  |
| 2:9.19 | זעיר |  |
| 2:9.20 | אנפין |  |
| 2:9.21 | זה |  |
| 2:9.22 | באצילות |  |
| 2:9.23 | הוא |  |
| 2:9.24 | הזנם |  |
| 2:9.25 | ומפרנסם | , |
| 2:9.26 | ולו |  |
| 2:9.27 | הם |  |
| 2:9.28 | עובדים |  |
| 2:9.29 | ומשתחוים | , |
| 2:9.30 | או |  |
| 2:9.31 | האין |  |
| 2:9.32 | סוף |  |
| 2:9.33 | הוא |  |
| 2:9.34 | אלהיהם |  |
| 2:9.35 | ולו |  |
| 2:9.36 | לבדו |  |
| 2:9.37 | עובדים | , |
| 2:9.38 | ואנחנו |  |
| 2:9.39 | לז"א |  |
| 2:9.40 | עובדים | , |
| 2:9.41 | כי |  |
| 2:9.42 | הוא |  |
| 2:9.43 | הזן |  |
| 2:9.44 | אותנו |  |
| 2:9.45 | והוא |  |
| 2:9.46 | אלהינו |  |
| 2:9.47 | כמו |  |
| 2:9.48 | שאמרו |  |
| 2:9.49 | המקובלים | . |
| 2:9.50 | ונמצא |  |
| 2:9.51 | לפי |  |
| 2:9.52 | זה | , |
| 2:9.53 | כי |  |
| 2:9.54 | אין |  |
| 2:9.55 | סוף |  |
| 2:9.56 | הוא |  |
| 2:9.57 | אלהי |  |
| 2:9.58 | העליונים | , |
| 2:9.59 | וז"א |  |
| 2:9.60 | אלהי |  |
| 2:9.61 | התחתונים | , |
| 2:9.62 | והרי |  |
| 2:9.63 | כאן |  |
| 2:9.64 | שתי |  |
| 2:9.65 | רשויות |  |
| 2:9.66 | ח"ו | . |
| 2:9.67 | ובהדיא |  |
| 2:9.68 | אמרו |  |
| 2:9.69 | מקצתן | : |
| 2:9.70 | את |  |
| 2:9.71 | כל |  |
| 2:9.72 | מעשה |  |
| 2:9.73 | ה | ' |
| 2:9.74 | הגדול |  |
| 2:9.75 | מכלל |  |
| 2:9.76 | דאיכא |  |
| 2:9.77 | קטן | . |
| 2:9.78 | ורבותינו |  |
| 2:9.79 | האמתיים |  |
| 2:9.80 | אמרו |  |
| 2:9.81 | הובא |  |
| 2:9.82 | במנורת |  |
| 2:9.83 | המאור |  |
| 2:9.84 | פ | ' |
| 2:9.85 | קמ"ג |  |
| 2:9.86 | שבשעת |  |
| 2:9.87 | מתן |  |
| 2:9.88 | תורה |  |
| 2:9.89 | קרע |  |
| 2:9.90 | להם |  |
| 2:9.91 | שבעה |  |
| 2:9.92 | רקיעים |  |
| 2:9.93 | ופתח |  |
| 2:9.94 | ז | ' |
| 2:9.95 | ארצות |  |
| 2:9.96 | ואמר |  |
| 2:9.97 | להם |  |
| 2:9.98 | לישראל | : |
| 2:9.99 | בינו |  |
| 2:9.100 | בעליונים |  |
| 2:9.101 | ובינו |  |
| 2:9.102 | בתחתונים |  |
| 2:9.103 | ודעו |  |
| 2:9.104 | שאין |  |
| 2:9.105 | אלוה |  |
| 2:9.106 | מבלעדי |  |
| 2:9.107 | לא |  |
| 2:9.108 | בעליונים |  |
| 2:9.109 | ולא |  |
| 2:9.110 | בתחתונים | , |
| 2:9.111 | אנכי |  |
| 2:9.112 | ה | ' |
| 2:9.113 | אלהיך | . |
| 2:9.114 | ובמדרש |  |
| 2:9.115 | רבה |  |
| 2:9.116 | פרשת |  |
| 2:9.117 | יתרו |  |
| 2:9.118 | אמרו | : |
| 2:9.119 | אנכי |  |
| 2:9.120 | ה | ' |
| 2:9.121 | אלהיך | . |
| 2:9.122 | אמר |  |
| 2:9.123 | ר | ' |
| 2:9.124 | אבהו |  |
| 2:9.125 | משל |  |
| 2:9.126 | למלך |  |
| 2:9.127 | בשר |  |
| 2:9.128 | ודם |  |
| 2:9.129 | מולך | , |
| 2:9.130 | יש |  |
| 2:9.131 | לו |  |
| 2:9.132 | אב |  |
| 2:9.133 | או |  |
| 2:9.134 | אח |  |
| 2:9.135 | או |  |
| 2:9.136 | בן | . |
| 2:9.137 | אמר |  |
| 2:9.138 | הקב"ה | : |
| 2:9.139 | אני |  |
| 2:9.140 | איני |  |
| 2:9.141 | כן | , |
| 2:9.142 | ראשון |  |
| 2:9.143 | ואני |  |
| 2:9.144 | אחרון |  |
| 2:9.145 | ומבלעדי |  |
| 2:9.146 | אין |  |
| 2:9.147 | אלהים | . |
| 2:9.148 | אני |  |
| 2:9.149 | ראשון |  |
| 2:9.150 | שאין |  |
| 2:9.151 | לי |  |
| 2:9.152 | אב | . |
| 2:9.153 | ואני |  |
| 2:9.154 | אחרון |  |
| 2:9.155 | שאין |  |
| 2:9.156 | לי |  |
| 2:9.157 | בן | . |
| 2:9.158 | ומבלעדי |  |
| 2:9.159 | אין |  |
| 2:9.160 | אלהים |  |
| 2:9.161 | שאין |  |
| 2:9.162 | לי |  |
| 2:9.163 | אח | . |
| 2:9.164 | עכ"ל |  |
| 2:9.165 | המדרש |  |
| 2:9.166 | וכו | ' |
| 2:9.167 | הר | ' |
| 2:9.168 | עץ |  |
| 2:9.169 | יוסף |  |
| 2:9.170 | ז"ל |  |
| 2:9.171 | ר"ל |  |
| 2:9.172 | שמלך |  |
| 2:9.173 | ב"ו |  |
| 2:9.174 | אין |  |
| 2:9.175 | מלכותו |  |
| 2:9.176 | שלימה |  |
| 2:9.177 | לפי |  |
| 2:9.178 | שיש |  |
| 2:9.179 | לו |  |
| 2:9.180 | אב |  |
| 2:9.181 | וצריך |  |
| 2:9.182 | לנהוג |  |
| 2:9.183 | בו |  |
| 2:9.184 | כבוד |  |
| 2:9.185 | ומשתתף |  |
| 2:9.186 | עמו |  |
| 2:9.187 | במלכות | . |
| 2:9.188 | או |  |
| 2:9.189 | יש |  |
| 2:9.190 | לו |  |
| 2:9.191 | אח |  |
| 2:9.192 | שיש |  |
| 2:9.193 | לו |  |
| 2:9.194 | חלק |  |
| 2:9.195 | בכבוד |  |
| 2:9.196 | מלכות | . |
| 2:9.197 | או |  |
| 2:9.198 | יש |  |
| 2:9.199 | לו |  |
| 2:9.200 | בן |  |
| 2:9.201 | שנוהג |  |
| 2:9.202 | שררה |  |
| 2:9.203 | עמו | , |
| 2:9.204 | ואחריו |  |
| 2:9.205 | יורש |  |
| 2:9.206 | עצר | . |
| 2:9.207 | ופי | ' |
| 2:9.208 | אין |  |
| 2:9.209 | לו |  |
| 2:9.210 | אב |  |
| 2:9.211 | שהוא |  |
| 2:9.212 | הסבה |  |
| 2:9.213 | הראשונה | . |
| 2:9.214 | ואין |  |
| 2:9.215 | לו |  |
| 2:9.216 | אח |  |
| 2:9.217 | כי |  |
| 2:9.218 | הוא |  |
| 2:9.219 | אחד |  |
| 2:9.220 | ואין |  |
| 2:9.221 | שני |  |
| 2:9.222 | לו | . |
| 2:9.223 | ואין |  |
| 2:9.224 | לו |  |
| 2:9.225 | בן |  |
| 2:9.226 | שלא |  |
| 2:9.227 | נמשכו |  |
| 2:9.228 | העלולים |  |
| 2:9.229 | ממנו |  |
| 2:9.230 | בטבע |  |
| 2:9.231 | כבן |  |
| 2:9.232 | מן |  |
| 2:9.233 | האב | , |
| 2:9.234 | כמו |  |
| 2:9.235 | שחשבו |  |
| 2:9.236 | קצת |  |
| 2:9.237 | היא |  |
| 2:9.238 | דעת |  |
| 2:9.239 | המקובלים |  |
| 2:9.240 | שאמרו |  |
| 2:9.241 | שז"א |  |
| 2:9.242 | בן |  |
| 2:9.243 | לאו"א |  |
| 2:9.244 | והוא |  |
| 2:9.245 | אב |  |
| 2:9.246 | לעלולים |  |
| 2:9.247 | שאחריו |  |
| 2:9.248 | וכו | ' |
| 2:9.249 | ויש |  |
| 2:9.250 | לו |  |
| 2:9.251 | אחות |  |
| 2:9.252 | או |  |
| 2:9.253 | אחיות |  |
| 2:9.254 | נוקבין |  |
| 2:9.255 | דיליה | ) |
| 2:9.256 | כי |  |
| 2:9.257 | לפי |  |
| 2:9.258 | הדעת |  |
| 2:9.259 | ההיא |  |
| 2:9.260 | אין |  |
| 2:9.261 | היכולת |  |
| 2:9.262 | בידו |  |
| 2:9.263 | לשנות |  |
| 2:9.264 | שום |  |
| 2:9.265 | דבר |  |
| 2:9.266 | מן |  |
| 2:9.267 | הדברים |  |
| 2:9.268 | ואין |  |
| 2:9.269 | מלכותו |  |
| 2:9.270 | שלימה | . |
| 2:9.271 | אבל |  |
| 2:9.272 | הקב"ה |  |
| 2:9.273 | אמר | : |
| 2:9.274 | אנכי |  |
| 2:9.275 | ה | ' |
| 2:9.276 | אלהיך | , |
| 2:9.277 | כי |  |
| 2:9.278 | זה |  |
| 2:9.279 | כחו |  |
| 2:9.280 | להוציאם |  |
| 2:9.281 | מארץ |  |
| 2:9.282 | מצרים |  |
| 2:9.283 | לפי |  |
| 2:9.284 | שמלכותו |  |
| 2:9.285 | שלימה |  |
| 2:9.286 | ולית |  |
| 2:9.287 | דימחי |  |
| 2:9.288 | בידיה |  |
| 2:9.289 | עכ"ל | : |


### Translation under review (2:9)

English: Would that, my lord, you would inform me also, who governs all the hidden worlds, of which the Kabbalists did not speak, and who feeds and sustains them. Whether Ze'ir Anpin, this one in the World of Emanation, is the one who feeds and sustains them, and to him they worship and bow down, or Ein Sof is their god, and to Him alone they worship; and we worship Ze'ir Anpin, for he feeds us, and he is our god, as the Kabbalists said. It follows from this that Ein Sof is the god of the upper ones and Ze'ir Anpin the god of the lower ones, and here we have two authorities, God forbid. and explicitly, some of them said: "every great deed of the LORD," by implication there is a small one. And our true Rabbis said, as cited in Menorat ha-Maor, chapter 143, that at the time of the giving of the Torah He tore open for them seven heavens and opened seven lands, and said to them, to Israel: "Understand the upper ones and understand the lower ones, and know that there is no God besides Me, neither among the upper ones nor among the lower ones: I am the LORD your God." In Midrash Rabbah, Parashat Yitro, they said: "I am the LORD your God." Rabbi Abbahu said: "A parable of a king of flesh and blood who reigns; he has a father, or a brother, or a son. The Holy One, blessed be He, said: I am not so. I am first and I am last, and besides Me there is no God. I am first, for I have no father. I am last, for I have no son. And besides Me there is no God, for I have no brother." Thus far the Midrash, etc. The Rabbi Etz Yosef, of blessed memory, that is to say, "that a king of flesh and blood, his kingship is not complete, because he has a father, whom he must honor, and who shares the kingship with him, or he has a brother who has a share in the honor of kingship, or he has a son who exercises authority with him, and after him is heir to the throne. And the explanation: he has no father, which is the First Cause. And he has no brother, for He is one, and has no second. And he has no son, in that the effects were not drawn from Him by nature, as a son is drawn from a father, as some thought (this is the view of the Kabbalists, who said that Ze'ir Anpin is a son of Abba and Imma, and he is a father to the effects that follow him, etc., and he has a sister or sisters, his nukvin) for according to that view he has no power to change any of the things, and his kingship is not complete. But the Holy One, blessed be He, said: 'I am the LORD your God,' for this is His power to bring them out of the land of Egypt, since His kingship is complete and there is none who can thwart His hand." Thus far his words.

Literal: would-that my-lord inform-me also who governs all the-worlds the-hidden which did-not speak of-them the-Kabbalists and-who feeds-them and-sustains-them whether Ze'ir Anpin this in-the-World-of-Emanation is the-one-who-feeds-them and-sustains-them and-to-him they worship and-bow-down or Ein Sof is their-god and-to-Him alone they-worship and-we to-Ze'ir-Anpin worship for he feeds us and-he is-our-god as they-said the-Kabbalists and-it-follows according-to this that Ein Sof is god-of the-upper-ones and Ze'ir-Anpin god-of the-lower-ones and-here we-have two authorities God-forbid and-explicitly said some-of-them (object marker) every deed the-great of-the-LORD by-implication there-is a-small-one and-our-Rabbis true said cited in-Menorat ha-Maor chapter 143 at-the-time-of the-giving-of the-Torah tore-open for-them seven heavens and-opened seven lands and-said to-them to-Israel understand the-upper-ones and-understand the-lower-ones and-know that-there-is-no God besides-me neither in-the-upper-ones nor in-the-lower-ones I the-LORD your-God and-in Midrash Rabbah Parashat Yitro said I the-LORD your-God said Rabbi Abbahu parable of-king of-flesh and-blood who-reigns he-has to-him father or brother or son said the-Holy-One-blessed-be-He I am-not so first and-I last and-besides-me there-is-no God I first that-there-is-no to-me father and-I last that-there-is-no to-me son and-besides-me there-is-no God that-there-is-no to-me brother thus-far his-words the-Midrash etc. the-Rabbi Etz Yosef of-blessed-memory that-is-to-say that a-king of-flesh-and-blood is-not his-kingship complete because he-has to-him a-father and-must show him honor and-shares with-him in-the-kingship or he-has to-him brother who-has to-him a-share in-the-honor of-kingship or he-has to-him son who-exercises authority with-him and-after-him heir to-the-throne and-the-explanation he-has-no to-him father which-is the-cause the-first and-no to-him brother for He is-one and-there-is-no second to-him and-he-has-no to-him son in-that not were-drawn the-effects from-Him by-nature as a-son from the-father as that-they-thought some this-is the-view of-the-Kabbalists who-said that Ze'ir-Anpin is-a-son of-Abba-and-Imma and-he is-a-father to-the-effects that-follow-him etc. and-he-has to-him sister or sisters nukvin his for according-to that view that there-is-no the-power in-his-hand to-change any thing of the-things and-no his-kingship complete but the-Holy-One-blessed-be-He said I the-LORD your-God for this his-power to-bring-them out-of-the-land-of Egypt since his-kingship complete and-there-is-none who-can-thwart his-hand thus-far his-words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 2:9.1 | would that |  |  |  |
| 2:9.2 | my lord |  |  | The questioner's letter is addressed to 'my lord' (his respondent); the paragraph is in his voice. |
| 2:9.3 | inform me |  |  |  |
| 2:9.4 | also |  |  |  |
| 2:9.5 | who |  |  |  |
| 2:9.6 | governs |  |  |  |
| 2:9.7 | all |  |  |  |
| 2:9.8 | the worlds |  |  |  |
| 2:9.9 | the hidden |  |  |  |
| 2:9.10 | which |  |  |  |
| 2:9.11 | did not |  |  |  |
| 2:9.12 | speak |  |  |  |
| 2:9.13 | of them |  |  |  |
| 2:9.14 | the Kabbalists |  |  |  |
| 2:9.15 | and who |  |  |  |
| 2:9.16 | feeds them |  |  |  |
| 2:9.17 | and sustains them |  |  |  |
| 2:9.18 | whether |  |  |  |
| 2:9.19 | Ze'ir |  |  | Ze'ir Anpin: Aramaic, 'small face'. The name is transliterated; this is its first occurrence in the unit, so the sense is given here. Ze'ir alone is the short form. |
| 2:9.20 | Anpin |  |  |  |
| 2:9.21 | this one |  |  |  |
| 2:9.22 | in the World of Emanation |  |  | The World of Emanation (olam ha-atzilut), the first of the four worlds: Emanation, Creation, Formation, Making. |
| 2:9.23 | is |  |  |  |
| 2:9.24 | the one who feeds them |  |  |  |
| 2:9.25 | and sustains them |  |  |  |
| 2:9.26 | and to him |  |  |  |
| 2:9.27 | they |  |  |  |
| 2:9.28 | worship |  |  |  |
| 2:9.29 | and bow down |  |  |  |
| 2:9.30 | or |  |  |  |
| 2:9.31 | Ein |  |  | Ein Sof, literally 'without end'. A name with no article; the two words are one name. |
| 2:9.32 | Sof |  |  |  |
| 2:9.33 | is |  |  |  |
| 2:9.34 | their god |  |  | The questioner's word for the god of the Kabbalists' view: lower case 'god', not the God of the verses. |
| 2:9.35 | and to Him |  |  |  |
| 2:9.36 | alone |  |  |  |
| 2:9.37 | they worship |  |  |  |
| 2:9.38 | and we |  |  |  |
| 2:9.39 | to Ze'ir Anpin |  |  | Ze'ir Anpin, with the preposition ל (to). |
| 2:9.40 | worship |  |  |  |
| 2:9.41 | for |  |  |  |
| 2:9.42 | he |  |  |  |
| 2:9.43 | feeds |  |  |  |
| 2:9.44 | us |  |  |  |
| 2:9.45 | and he |  |  |  |
| 2:9.46 | is our god |  |  | 'Our god' is the questioner's phrase for Ze'ir Anpin, quoting the Kabbalists' view (see the note on token 34). |
| 2:9.47 | as |  |  |  |
| 2:9.48 | said |  |  |  |
| 2:9.49 | the Kabbalists |  |  |  |
| 2:9.50 | and so |  |  |  |
| 2:9.51 | according to this |  |  |  |
| 2:9.52 | this |  |  |  |
| 2:9.53 | that |  |  |  |
| 2:9.54 | Ein |  |  |  |
| 2:9.55 | Sof |  |  |  |
| 2:9.56 | is |  |  |  |
| 2:9.57 | the god of |  |  |  |
| 2:9.58 | the upper ones |  |  |  |
| 2:9.59 | and Ze'ir Anpin |  |  |  |
| 2:9.60 | the god of |  |  |  |
| 2:9.61 | the lower ones |  |  |  |
| 2:9.62 | and here |  |  |  |
| 2:9.63 | we have |  |  |  |
| 2:9.64 | two |  |  |  |
| 2:9.65 | authorities |  |  | רשויות, literally 'authorities': the two powers that would rule, which the questioner takes to follow. |
| 2:9.66 | God forbid |  |  | ח"ו = חס ושלום, 'God forbid'. |
| 2:9.67 | and explicitly |  |  |  |
| 2:9.68 | said |  |  |  |
| 2:9.69 | some of them |  |  |  |
| 2:9.70 | (object marker) |  |  | את is the object marker and has no English counterpart. |
| 2:9.71 | every |  |  |  |
| 2:9.72 | deed |  |  |  |
| 2:9.73 | of the LORD |  |  |  |
| 2:9.74 | great |  |  |  |
| 2:9.75 | by implication |  |  | מכלל דאיכא קטן, 'by implication there is a small one': the Talmudic inference that the great implies a lesser. |
| 2:9.76 | there is |  |  |  |
| 2:9.77 | a small one |  |  |  |
| 2:9.78 | and our Rabbis |  |  |  |
| 2:9.79 | true |  |  |  |
| 2:9.80 | said |  |  |  |
| 2:9.81 | cited |  |  |  |
| 2:9.82 | in Menorat |  |  | Menorat ha-Maor, 'the candelabrum of light', a book; the title is transliterated and the reference kept as written. |
| 2:9.83 | ha-Maor |  |  |  |
| 2:9.84 | chapter |  |  | פ' is abbreviated and may read 'chapter' or 'section' (פרק or פרשה); read as 'chapter', and kept as the author writes it. |
| 2:9.85 | 143 |  |  | קמ"ג = 143 (the letters 100, 40, 3). |
| 2:9.86 | at the time |  |  |  |
| 2:9.87 | of the giving |  |  |  |
| 2:9.88 | of the Torah |  |  |  |
| 2:9.89 | tore open |  |  |  |
| 2:9.90 | for them |  |  |  |
| 2:9.91 | seven |  |  |  |
| 2:9.92 | heavens |  |  |  |
| 2:9.93 | and opened |  |  |  |
| 2:9.94 | seven |  |  | ז' = seven. |
| 2:9.95 | lands |  |  |  |
| 2:9.96 | and said |  |  |  |
| 2:9.97 | to them |  |  |  |
| 2:9.98 | to Israel |  |  |  |
| 2:9.99 | understand |  |  | Plural imperative, addressed to Israel: 'understand'. |
| 2:9.100 | the upper ones |  |  |  |
| 2:9.101 | and understand |  |  |  |
| 2:9.102 | the lower ones |  |  |  |
| 2:9.103 | and know |  |  |  |
| 2:9.104 | that there is no |  |  |  |
| 2:9.105 | God |  |  |  |
| 2:9.106 | besides Me |  |  |  |
| 2:9.107 | neither |  |  |  |
| 2:9.108 | among the upper ones |  |  |  |
| 2:9.109 | nor |  |  |  |
| 2:9.110 | among the lower ones |  |  |  |
| 2:9.111 | I |  |  |  |
| 2:9.112 | the LORD |  |  | ה' as written is the Tetragrammaton: 'the LORD'. |
| 2:9.113 | your God |  |  |  |
| 2:9.114 | and in Midrash |  |  | Midrash Rabbah: the great midrashic collection on the Torah. |
| 2:9.115 | Rabbah |  |  |  |
| 2:9.116 | Parashat |  |  | Parashat: the parashah of the Torah reading (Exod 18-20). |
| 2:9.117 | Yitro |  |  |  |
| 2:9.118 | said |  |  |  |
| 2:9.119 | I |  |  |  |
| 2:9.120 | the LORD |  |  |  |
| 2:9.121 | your God |  |  |  |
| 2:9.122 | said |  |  |  |
| 2:9.123 | Rabbi |  |  |  |
| 2:9.124 | Abbahu |  |  |  |
| 2:9.125 | a parable |  |  |  |
| 2:9.126 | of a king |  |  |  |
| 2:9.127 | of flesh |  |  |  |
| 2:9.128 | and blood |  |  |  |
| 2:9.129 | who reigns |  |  |  |
| 2:9.130 | he has |  |  |  |
| 2:9.131 | to him |  |  |  |
| 2:9.132 | a father |  |  |  |
| 2:9.133 | or |  |  |  |
| 2:9.134 | a brother |  |  |  |
| 2:9.135 | or |  |  |  |
| 2:9.136 | a son |  |  |  |
| 2:9.137 | said |  |  |  |
| 2:9.138 | the Holy One, blessed be He |  |  | הקב"ה = 'the Holy One, blessed be He'. |
| 2:9.139 | I |  |  |  |
| 2:9.140 | am not |  |  |  |
| 2:9.141 | so |  |  |  |
| 2:9.142 | first |  |  |  |
| 2:9.143 | and I |  |  |  |
| 2:9.144 | last |  |  |  |
| 2:9.145 | and besides Me |  |  |  |
| 2:9.146 | there is no |  |  |  |
| 2:9.147 | God |  |  |  |
| 2:9.148 | I |  |  |  |
| 2:9.149 | first |  |  |  |
| 2:9.150 | that there is no |  |  |  |
| 2:9.151 | to me |  |  |  |
| 2:9.152 | father |  |  |  |
| 2:9.153 | and I |  |  |  |
| 2:9.154 | last |  |  |  |
| 2:9.155 | that there is no |  |  |  |
| 2:9.156 | to me |  |  |  |
| 2:9.157 | son |  |  |  |
| 2:9.158 | and besides Me |  |  |  |
| 2:9.159 | there is no |  |  |  |
| 2:9.160 | God |  |  |  |
| 2:9.161 | that there is no |  |  |  |
| 2:9.162 | to me |  |  |  |
| 2:9.163 | brother |  |  |  |
| 2:9.164 | thus far his words |  |  | עכ"ל: the closing formula of a quotation, 'thus far his words'. |
| 2:9.165 | the Midrash |  |  |  |
| 2:9.166 | etc. |  |  | וכו' is kept as 'etc.', because it marks a truncated quotation. |
| 2:9.167 | the Rabbi |  |  | הר', 'the Rabbi', before the name of the commentary Etz Yosef, which the Hebrew gives in place of a personal name. |
| 2:9.168 | Etz |  |  |  |
| 2:9.169 | Yosef |  |  |  |
| 2:9.170 | of blessed memory |  |  |  |
| 2:9.171 | that is to say |  |  | ר"ל = רוצה לומר, 'that is to say'; it introduces the commentator's sense of the Midrash. |
| 2:9.172 | that a king |  |  |  |
| 2:9.173 | of flesh and blood |  |  | ב"ו = בשר ודם, 'flesh and blood': a human king. |
| 2:9.174 | is not |  |  |  |
| 2:9.175 | his kingship |  |  |  |
| 2:9.176 | complete |  |  |  |
| 2:9.177 | because |  |  |  |
| 2:9.178 | he has |  |  |  |
| 2:9.179 | to him |  |  |  |
| 2:9.180 | a father |  |  |  |
| 2:9.181 | and must |  |  |  |
| 2:9.182 | show |  |  |  |
| 2:9.183 | him |  |  |  |
| 2:9.184 | honor |  |  |  |
| 2:9.185 | and shares |  |  |  |
| 2:9.186 | with him |  |  |  |
| 2:9.187 | in the kingship |  |  |  |
| 2:9.188 | or |  |  |  |
| 2:9.189 | he has |  |  |  |
| 2:9.190 | to him |  |  |  |
| 2:9.191 | a brother |  |  |  |
| 2:9.192 | who has |  |  |  |
| 2:9.193 | to him |  |  |  |
| 2:9.194 | a share |  |  |  |
| 2:9.195 | in the honor |  |  |  |
| 2:9.196 | of kingship |  |  |  |
| 2:9.197 | or |  |  |  |
| 2:9.198 | he has |  |  |  |
| 2:9.199 | to him |  |  |  |
| 2:9.200 | a son |  |  |  |
| 2:9.201 | who exercises |  |  |  |
| 2:9.202 | authority |  |  |  |
| 2:9.203 | with him |  |  |  |
| 2:9.204 | and after him |  |  |  |
| 2:9.205 | heir |  |  |  |
| 2:9.206 | to the throne |  |  |  |
| 2:9.207 | and the explanation |  |  |  |
| 2:9.208 | he has no |  |  |  |
| 2:9.209 | to him |  |  |  |
| 2:9.210 | father |  |  |  |
| 2:9.211 | which is |  |  |  |
| 2:9.212 | the |  |  | הסבה הראשונה, 'the First Cause', a term of the philosophers. |
| 2:9.213 | First Cause |  |  |  |
| 2:9.214 | and no |  |  |  |
| 2:9.215 | to him |  |  |  |
| 2:9.216 | brother |  |  |  |
| 2:9.217 | for |  |  |  |
| 2:9.218 | He |  |  |  |
| 2:9.219 | is one |  |  |  |
| 2:9.220 | and there is no |  |  |  |
| 2:9.221 | second |  |  |  |
| 2:9.222 | to him |  |  |  |
| 2:9.223 | and he has no |  |  |  |
| 2:9.224 | to him |  |  |  |
| 2:9.225 | son |  |  |  |
| 2:9.226 | in that not |  |  |  |
| 2:9.227 | were drawn |  |  |  |
| 2:9.228 | the effects |  |  |  |
| 2:9.229 | from Him |  |  |  |
| 2:9.230 | by nature |  |  |  |
| 2:9.231 | as a son |  |  |  |
| 2:9.232 | from |  |  |  |
| 2:9.233 | the father |  |  |  |
| 2:9.234 | as |  |  |  |
| 2:9.235 | that they thought |  |  |  |
| 2:9.236 | some |  |  |  |
| 2:9.237 | this is |  |  |  |
| 2:9.238 | the view |  |  |  |
| 2:9.239 | of the Kabbalists |  |  |  |
| 2:9.240 | who said |  |  |  |
| 2:9.241 | that Ze'ir Anpin |  |  | שז"א = ש + ז"א: 'that Ze'ir Anpin'. |
| 2:9.242 | is a son |  |  |  |
| 2:9.243 | of Abba and Imma |  |  | לאו"א = לאבא ולאמא: Abba and Imma, the upper partzufim, transliterated. |
| 2:9.244 | and he |  |  |  |
| 2:9.245 | is a father |  |  |  |
| 2:9.246 | to the effects |  |  |  |
| 2:9.247 | that follow him |  |  |  |
| 2:9.248 | etc. |  |  |  |
| 2:9.249 | and he has |  |  |  |
| 2:9.250 | to him |  |  |  |
| 2:9.251 | a sister |  |  |  |
| 2:9.252 | or |  |  |  |
| 2:9.253 | sisters |  |  |  |
| 2:9.254 | nukvin |  |  | Aramaic plural of nukva, 'female'; transliterated as nukvin. |
| 2:9.255 | his |  |  | Aramaic 'of his'. |
| 2:9.256 | for |  |  |  |
| 2:9.257 | according to |  |  |  |
| 2:9.258 | that view |  |  |  |
| 2:9.259 | that |  |  |  |
| 2:9.260 | there is no |  |  |  |
| 2:9.261 | the power |  |  |  |
| 2:9.262 | in his hand |  |  |  |
| 2:9.263 | to change |  |  |  |
| 2:9.264 | any |  |  |  |
| 2:9.265 | thing |  |  |  |
| 2:9.266 | of |  |  |  |
| 2:9.267 | the things |  |  |  |
| 2:9.268 | and no |  |  |  |
| 2:9.269 | his kingship |  |  |  |
| 2:9.270 | complete |  |  |  |
| 2:9.271 | but |  |  |  |
| 2:9.272 | the Holy One, blessed be He |  |  | הקב"ה = 'the Holy One, blessed be He'. |
| 2:9.273 | said |  |  |  |
| 2:9.274 | I |  |  |  |
| 2:9.275 | the LORD |  |  |  |
| 2:9.276 | your God |  |  |  |
| 2:9.277 | for |  |  |  |
| 2:9.278 | this |  |  |  |
| 2:9.279 | his power |  |  |  |
| 2:9.280 | to bring them out |  |  |  |
| 2:9.281 | out of the land of |  |  |  |
| 2:9.282 | Egypt |  |  |  |
| 2:9.283 | since |  |  |  |
| 2:9.284 | his kingship |  |  |  |
| 2:9.285 | complete |  |  |  |
| 2:9.286 | and there is none |  |  |  |
| 2:9.287 | who can thwart |  |  | Aramaic, 'can wipe out, thwart' (dimchi), from the root 'to strike, wipe out'. |
| 2:9.288 | His hand |  |  |  |
| 2:9.289 | thus far his words |  |  | עכ"ל: the closing formula, 'thus far his words'. |

Phrases:
- 2:9.1 2:9.2 2:9.3 2:9.4 | Would that, my lord, you would inform me also, | would-that my-lord inform-me also
- 2:9.5 2:9.6 2:9.7 2:9.8 2:9.9 2:9.10 2:9.11 2:9.12 2:9.13 2:9.14 2:9.15 2:9.16 2:9.17 | who governs all the hidden worlds, of which the Kabbalists did not speak, and who feeds and sustains them. | who governs all the-worlds the-hidden which did-not speak of-them the-Kabbalists and-who feeds-them and-sustains-them
- 2:9.18 2:9.19 2:9.20 2:9.21 2:9.22 2:9.23 2:9.24 2:9.25 2:9.26 2:9.27 2:9.28 2:9.29 | Whether Ze'ir Anpin, this one in the World of Emanation, is the one who feeds and sustains them, and to him they worship and bow down, | whether Ze'ir Anpin this in-the-World-of-Emanation is the-one-who-feeds-them and-sustains-them and-to-him they worship and-bow-down
- 2:9.30 2:9.31 2:9.32 2:9.33 2:9.34 2:9.35 2:9.36 2:9.37 | or Ein Sof is their god, and to Him alone they worship; | or Ein Sof is their-god and-to-Him alone they-worship
- 2:9.38 2:9.39 2:9.40 2:9.41 2:9.42 2:9.43 2:9.44 2:9.45 2:9.46 2:9.47 2:9.48 2:9.49 | and we worship Ze'ir Anpin, for he feeds us, and he is our god, as the Kabbalists said. | and-we to-Ze'ir-Anpin worship for he feeds us and-he is-our-god as they-said the-Kabbalists
- 2:9.50 2:9.51 2:9.52 2:9.53 2:9.54 2:9.55 2:9.56 2:9.57 2:9.58 2:9.59 2:9.60 2:9.61 2:9.62 2:9.63 2:9.64 2:9.65 2:9.66 | It follows from this that Ein Sof is the god of the upper ones and Ze'ir Anpin the god of the lower ones, and here we have two authorities, God forbid. | and-it-follows according-to this that Ein Sof is god-of the-upper-ones and Ze'ir-Anpin god-of the-lower-ones and-here we-have two authorities God-forbid
- 2:9.67 2:9.68 2:9.69 | and explicitly, some of them said: | and-explicitly said some-of-them
- 2:9.70 | - | (object marker)
- 2:9.71 2:9.72 2:9.73 2:9.74 2:9.75 2:9.76 2:9.77 | "every great deed of the LORD," by implication there is a small one. | every deed the-great of-the-LORD by-implication there-is a-small-one
- 2:9.78 2:9.79 2:9.80 2:9.81 2:9.82 2:9.83 2:9.84 2:9.85 2:9.86 2:9.87 2:9.88 | And our true Rabbis said, as cited in Menorat ha-Maor, chapter 143, that at the time of the giving of the Torah | and-our-Rabbis true said cited in-Menorat ha-Maor chapter 143 at-the-time-of the-giving-of the-Torah
- 2:9.89 2:9.90 2:9.91 2:9.92 2:9.93 2:9.94 2:9.95 2:9.96 2:9.97 2:9.98 | He tore open for them seven heavens and opened seven lands, and said to them, to Israel: | tore-open for-them seven heavens and-opened seven lands and-said to-them to-Israel
- 2:9.99 2:9.100 2:9.101 2:9.102 2:9.103 2:9.104 2:9.105 2:9.106 2:9.107 2:9.108 2:9.109 2:9.110 2:9.111 2:9.112 2:9.113 | "Understand the upper ones and understand the lower ones, and know that there is no God besides Me, neither among the upper ones nor among the lower ones: I am the LORD your God." | understand the-upper-ones and-understand the-lower-ones and-know that-there-is-no God besides-me neither in-the-upper-ones nor in-the-lower-ones I the-LORD your-God
- 2:9.114 2:9.115 2:9.116 2:9.117 2:9.118 2:9.119 2:9.120 2:9.121 | In Midrash Rabbah, Parashat Yitro, they said: "I am the LORD your God." | and-in Midrash Rabbah Parashat Yitro said I the-LORD your-God
- 2:9.122 2:9.123 2:9.124 | Rabbi Abbahu said: | said Rabbi Abbahu
- 2:9.125 2:9.126 2:9.127 2:9.128 2:9.129 2:9.130 2:9.131 2:9.132 2:9.133 2:9.134 2:9.135 2:9.136 | "A parable of a king of flesh and blood who reigns; he has a father, or a brother, or a son. | parable of-king of-flesh and-blood who-reigns he-has to-him father or brother or son
- 2:9.137 2:9.138 2:9.139 2:9.140 2:9.141 2:9.142 2:9.143 2:9.144 2:9.145 2:9.146 2:9.147 | The Holy One, blessed be He, said: I am not so. I am first and I am last, and besides Me there is no God. | said the-Holy-One-blessed-be-He I am-not so first and-I last and-besides-me there-is-no God
- 2:9.148 2:9.149 2:9.150 2:9.151 2:9.152 | I am first, for I have no father. | I first that-there-is-no to-me father
- 2:9.153 2:9.154 2:9.155 2:9.156 2:9.157 | I am last, for I have no son. | and-I last that-there-is-no to-me son
- 2:9.158 2:9.159 2:9.160 2:9.161 2:9.162 2:9.163 | And besides Me there is no God, for I have no brother." | and-besides-me there-is-no God that-there-is-no to-me brother
- 2:9.164 2:9.165 2:9.166 | Thus far the Midrash, etc. | thus-far his-words the-Midrash etc.
- 2:9.167 2:9.168 2:9.169 2:9.170 | The Rabbi Etz Yosef, of blessed memory, | the-Rabbi Etz Yosef of-blessed-memory
- 2:9.171 | that is to say, "that | that-is-to-say
- 2:9.172 2:9.173 2:9.174 2:9.175 2:9.176 | a king of flesh and blood, his kingship is not complete, | that a-king of-flesh-and-blood is-not his-kingship complete
- 2:9.177 2:9.178 2:9.179 2:9.180 2:9.181 2:9.182 2:9.183 2:9.184 2:9.185 2:9.186 2:9.187 | because he has a father, whom he must honor, and who shares the kingship with him, | because he-has to-him a-father and-must show him honor and-shares with-him in-the-kingship
- 2:9.188 2:9.189 2:9.190 2:9.191 2:9.192 2:9.193 2:9.194 2:9.195 2:9.196 | or he has a brother who has a share in the honor of kingship, | or he-has to-him brother who-has to-him a-share in-the-honor of-kingship
- 2:9.197 2:9.198 2:9.199 2:9.200 2:9.201 2:9.202 2:9.203 2:9.204 2:9.205 2:9.206 | or he has a son who exercises authority with him, and after him is heir to the throne. | or he-has to-him son who-exercises authority with-him and-after-him heir to-the-throne
- 2:9.207 2:9.208 2:9.209 2:9.210 2:9.211 2:9.212 2:9.213 | And the explanation: he has no father, which is the First Cause. | and-the-explanation he-has-no to-him father which-is the-cause the-first
- 2:9.214 2:9.215 2:9.216 2:9.217 2:9.218 2:9.219 2:9.220 2:9.221 2:9.222 | And he has no brother, for He is one, and has no second. | and-no to-him brother for He is-one and-there-is-no second to-him
- 2:9.223 2:9.224 2:9.225 2:9.226 2:9.227 2:9.228 2:9.229 2:9.230 2:9.231 2:9.232 2:9.233 | And he has no son, in that the effects were not drawn from Him by nature, as a son is drawn from a father, | and-he-has-no to-him son in-that not were-drawn the-effects from-Him by-nature as a-son from the-father
- 2:9.234 2:9.235 2:9.236 | as some thought | as that-they-thought some
- 2:9.237 2:9.238 2:9.239 2:9.240 2:9.241 2:9.242 2:9.243 2:9.244 2:9.245 2:9.246 2:9.247 2:9.248 2:9.249 2:9.250 2:9.251 2:9.252 2:9.253 2:9.254 2:9.255 | (this is the view of the Kabbalists, who said that Ze'ir Anpin is a son of Abba and Imma, and he is a father to the effects that follow him, etc., and he has a sister or sisters, his nukvin) | this-is the-view of-the-Kabbalists who-said that Ze'ir-Anpin is-a-son of-Abba-and-Imma and-he is-a-father to-the-effects that-follow-him etc. and-he-has to-him sister or sisters nukvin his
- 2:9.256 2:9.257 2:9.258 2:9.259 2:9.260 2:9.261 2:9.262 2:9.263 2:9.264 2:9.265 2:9.266 2:9.267 2:9.268 2:9.269 2:9.270 | for according to that view he has no power to change any of the things, and his kingship is not complete. | for according-to that view that there-is-no the-power in-his-hand to-change any thing of the-things and-no his-kingship complete
- 2:9.271 2:9.272 2:9.273 | But the Holy One, blessed be He, said: | but the-Holy-One-blessed-be-He said
- 2:9.274 2:9.275 2:9.276 | 'I am the LORD your God,' | I the-LORD your-God
- 2:9.277 2:9.278 2:9.279 2:9.280 2:9.281 2:9.282 | for this is His power to bring them out of the land of Egypt, | for this his-power to-bring-them out-of-the-land-of Egypt
- 2:9.283 2:9.284 2:9.285 | since His kingship is complete | since his-kingship complete
- 2:9.286 2:9.287 2:9.288 | and there is none who can thwart His hand." | and-there-is-none who-can-thwart his-hand
- 2:9.289 | Thus far his words. | thus-far his-words

Grammar notes:
- אחלי (2:9.1) is a wish, 'would that'. The paragraph is the questioner's letter, addressed to 'my lord' (2:9.2), and is kept in his voice.
- The sentence 'אם ... או' (2:9.18, 2:9.30) gives two alternatives, 'whether ... or', in one statement that ends with a period, not a question; each alternative has its own verbs (עובדים, ומשתחוים; לבדו עובדים).
- ונמצא לפי זה (2:9.50-52) is the questioner's conclusion from the two alternatives, and הרי כאן שתי רשויות ח"ו (2:9.62-66) is his objection to it.
- From ובהדיא אמרו (2:9.67) the paragraph quotes its sources: the Mnorat ha-Maor, the Midrash Rabbah with its parable, and the commentary Etz Yosef. Each quotation is in quotation marks; the Etz Yosef passage runs from ר"ל (2:9.171) to the closing עכ"ל (2:9.289).
- Aramaic: נוקבין דיליה (2:9.254-255) and ולית דימחי בידיה (2:9.286-288); the English gives the sense of each.

Translator's note: The questioner's letter in his voice (chapter 2 correspondence). The English keeps each sentence and the quotation boundaries. Names and terms are transliterated per the decisions. 'God' is used for the God of the verses and of the Midrash's direct speech; 'god' in lower case is used for the divinities of the Kabbalists' view, as the questioner states it.

Commentary rulings:
- apparatus: Base text: Wikisource, not yet collated against the 1931 scan (Q2).
- apparatus: Deut 11:7: 'every great deed of the LORD' renders the words the author quotes from את כל מעשה ה' הגדול אשר עשה; the verse continues 'that He did', which the author does not quote.
- apparatus: Exod 20:2: 'I am the LORD your God' (אנכי ה' אלהיך), quoted in the Mnorat ha-Maor passage, the Midrash, and the Etz Yosef passage.
- apparatus: Isa 44:6: the Midrash's אני ראשון ... ומבלעדי אין אלהים echoes Isa 44:6; the English keeps the author's wording, 'I am first ... and besides Me there is no God'.
- apparatus: The reference 'פ' קמ"ג' to Menorat ha-Maor is kept as written and has not been checked.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

