# Translation pack: milhamot-hashem perek-6-11 (Milchamot ha-Shem 6:17–6:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-11`.

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
- **הוי"ה / הוי'ה / הויה** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG). אדנות as a name → "Adonai" (note: "lordship").
  עשר הויות → "ten havayot" (note).
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **כמתלהמים** → "ravers" (Prov 26:18 in `apparatus`).
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


## Milchamot ha-Shem 6:17

Hebrew (Wikisource): סח) ואח"כ (דף הנז"ל ע"ב) הוסיף להרבות דברים בדברי נרגן כמתלהמים, וכמתלהלה היורה זיקים לשבח סודותיו ולקרות תורתו עץ חיים היא למחזיקים בה, ואיהי מטרוניתא מלכות. והעוסקים בה בני מלכים, ואת המשנה שקבל משה מפי הגבורה ושנאה לאהרן ובניו וזקני ישראל ולכל ישראל כדגרסינן בעירובין פ' כיצד מעברין. ת"ר כיצד סדר משנה משה למד מפי הגבורה (לא מפי מטטרון) נכנס אהרן שנה לו משה פרקו, נסתלק אהרון וכו', קרא לה הפילסוף הנז' שפחה ואמר שאלו זכו ישראל היה משה נותן להם משנה שאין בה אסור וטמא ופסול, אלא הכל טהור ומותר וכשר, ועל שנתן להם משה רבינו המשנה הזאת שיש בה איסור והתר, טמא וטהור, פסול וכשר, נענש להקבר בחוצה לארץ, ולא ידע איש את קבורתו, כלומר סבת עונש קבורתו, וסבת עונש קבורתו שם, משנה זו שנתן לישראל. והעוסקים בה ומורים לעם את הדרך ילכו בה קרא להם ערב רב, לפי שהמשנה מרחקת אמונתו באלהים רבים שהוא מרמזם באותיות יהוה, ולכן קרא למשנה שפחה דשלטא על מטרוניתא שהיא גורמת להפריד בין מלכא ומטרוניתא, באמרו שם ומלכא ומטרוניתא מתפרשן מבעלה בגין כן תחת שלש רגזה ארץ, תחת עבד כי ימלך, דא איהו עבדא ידיעא (מטטרון שנתן המשנה הזאת למשה) ושפחה כי תירש גברתה דא משנה, ונבל כי ישבע לחם (מלחמה של משנה) דא ערב רב. (העוסקים במשנה) עם נבל ולא חכם, ואין יחוד לפרצופים עד דיתמחון ערב רב, העוסקים בה שהם גורמים פירוד בין המלכים ופילגשיהם. ובפרשת כי תצא הוא משבח את תורתו החדשה שהוא קורא אותה קבלה, והעוסקים בה קורא להם אדם, אבל העוסקים במשנה ובתלמוד הוא קורא להם בשם דגים ועופות השמים, והעוסקים בקבלתו למעלה מהם ושעליהם נאמר וירדו בדגת הים ובעוף השמים דאינון מארי מתניתין, ומרבה עוד דברים לגנות התנאים ואמוראים ולהפחית כבודם, ולשבח העוסקים בקבלתו החדשה בדברי חלקות היורדים חדרי בטן קלי הדעת וחלושי השכל, כמאמר שלמה ע״ה דברי נרגן כמתלהמים והם ירדו חדרי בטן, כאשר יראה המעיין המשכיל על דבר אמת הבקי בכתובים ובדברי רז"ל בתלמוד ומדרשים האמתיים. ושם בדף רע"ט גילה פנים בחוצפא יתירא בשם רעיא מהימנא לומר אית סלע ואית סלע, אית אבן ואית אבן, אית אבן דשמא דיהוה (היא תורתו החדשה המלאה פרצופים וצורות הנקראים בשם יהוה ואדנות ושאר שמות הקדש נמחקים ואין נמחקים ומרמזם בשם הוי"ה) עלה אמתר ואבנא די מחת לצלמא הות לטור רב, ואית אבן דאיהי אבן משכית, דלית תמן נביעו דחכמתא ולא דיבור אלא אבן דאיהי (חסר הפסוק בכאן ולא חפץ לפרסם כוונתו כל כך בבזוי המשנה והיה רוצה לומר אבן נגף ואבן מכשול). ראה והבן בדבריו איך הוא מחשיב לתורתו החדשה וקורא לה אבנא די מחת לצלמא, ובמלת צלמא כוונתו על המשנה והתלמוד שהוא חושב אותם כצלמי האלילים, ומבזה אותה וקורא לה אבן משכית דלית בה חכמתא, ובדבריו אלה הוא רומז לפרוש מן המשנה כמאמר הכתוב ואבן משכית לא תתנו בארצכם להשתחות עליה (פי' להשתחות וליחד השי"ת על פי המשנה המרחקת כל רבוי באלהות. רק על פי תורתו החדשה שהשי"ת מחולק לכמה פרצופים זכרים ונקבות וצריך לקשרם ולאחדם ולזווגם יחד ר"ל). ושוב חזר הפילסוף הנז' לדבריו שבראש עמוד ב' שקורא לתורתו סלע הנובע חכמה, ואת המשנה ואת התלמוד שקבל משה מסיני הוא מדמה אותן לקליפה הקשה של אגוז שכל חכמי הדורות יגעים בה ולית לון יכולת לאפקא מינה מיא אלא טיפין.

| id | word | punctuation after |
|---|---|---|
| 6:17.1 | סח | ) |
| 6:17.2 | ואח"כ |  |
| 6:17.3 | דף |  |
| 6:17.4 | הנז"ל |  |
| 6:17.5 | ע"ב | ) |
| 6:17.6 | הוסיף |  |
| 6:17.7 | להרבות |  |
| 6:17.8 | דברים |  |
| 6:17.9 | בדברי |  |
| 6:17.10 | נרגן |  |
| 6:17.11 | כמתלהמים | , |
| 6:17.12 | וכמתלהלה |  |
| 6:17.13 | היורה |  |
| 6:17.14 | זיקים |  |
| 6:17.15 | לשבח |  |
| 6:17.16 | סודותיו |  |
| 6:17.17 | ולקרות |  |
| 6:17.18 | תורתו |  |
| 6:17.19 | עץ |  |
| 6:17.20 | חיים |  |
| 6:17.21 | היא |  |
| 6:17.22 | למחזיקים |  |
| 6:17.23 | בה | , |
| 6:17.24 | ואיהי |  |
| 6:17.25 | מטרוניתא |  |
| 6:17.26 | מלכות | . |
| 6:17.27 | והעוסקים |  |
| 6:17.28 | בה |  |
| 6:17.29 | בני |  |
| 6:17.30 | מלכים | , |
| 6:17.31 | ואת |  |
| 6:17.32 | המשנה |  |
| 6:17.33 | שקבל |  |
| 6:17.34 | משה |  |
| 6:17.35 | מפי |  |
| 6:17.36 | הגבורה |  |
| 6:17.37 | ושנאה |  |
| 6:17.38 | לאהרן |  |
| 6:17.39 | ובניו |  |
| 6:17.40 | וזקני |  |
| 6:17.41 | ישראל |  |
| 6:17.42 | ולכל |  |
| 6:17.43 | ישראל |  |
| 6:17.44 | כדגרסינן |  |
| 6:17.45 | בעירובין |  |
| 6:17.46 | פ | ' |
| 6:17.47 | כיצד |  |
| 6:17.48 | מעברין | . |
| 6:17.49 | ת"ר |  |
| 6:17.50 | כיצד |  |
| 6:17.51 | סדר |  |
| 6:17.52 | משנה |  |
| 6:17.53 | משה |  |
| 6:17.54 | למד |  |
| 6:17.55 | מפי |  |
| 6:17.56 | הגבורה |  |
| 6:17.57 | לא |  |
| 6:17.58 | מפי |  |
| 6:17.59 | מטטרון | ) |
| 6:17.60 | נכנס |  |
| 6:17.61 | אהרן |  |
| 6:17.62 | שנה |  |
| 6:17.63 | לו |  |
| 6:17.64 | משה |  |
| 6:17.65 | פרקו | , |
| 6:17.66 | נסתלק |  |
| 6:17.67 | אהרון |  |
| 6:17.68 | וכו | ', |
| 6:17.69 | קרא |  |
| 6:17.70 | לה |  |
| 6:17.71 | הפילסוף |  |
| 6:17.72 | הנז | ' |
| 6:17.73 | שפחה |  |
| 6:17.74 | ואמר |  |
| 6:17.75 | שאלו |  |
| 6:17.76 | זכו |  |
| 6:17.77 | ישראל |  |
| 6:17.78 | היה |  |
| 6:17.79 | משה |  |
| 6:17.80 | נותן |  |
| 6:17.81 | להם |  |
| 6:17.82 | משנה |  |
| 6:17.83 | שאין |  |
| 6:17.84 | בה |  |
| 6:17.85 | אסור |  |
| 6:17.86 | וטמא |  |
| 6:17.87 | ופסול | , |
| 6:17.88 | אלא |  |
| 6:17.89 | הכל |  |
| 6:17.90 | טהור |  |
| 6:17.91 | ומותר |  |
| 6:17.92 | וכשר | , |
| 6:17.93 | ועל |  |
| 6:17.94 | שנתן |  |
| 6:17.95 | להם |  |
| 6:17.96 | משה |  |
| 6:17.97 | רבינו |  |
| 6:17.98 | המשנה |  |
| 6:17.99 | הזאת |  |
| 6:17.100 | שיש |  |
| 6:17.101 | בה |  |
| 6:17.102 | איסור |  |
| 6:17.103 | והתר | , |
| 6:17.104 | טמא |  |
| 6:17.105 | וטהור | , |
| 6:17.106 | פסול |  |
| 6:17.107 | וכשר | , |
| 6:17.108 | נענש |  |
| 6:17.109 | להקבר |  |
| 6:17.110 | בחוצה |  |
| 6:17.111 | לארץ | , |
| 6:17.112 | ולא |  |
| 6:17.113 | ידע |  |
| 6:17.114 | איש |  |
| 6:17.115 | את |  |
| 6:17.116 | קבורתו | , |
| 6:17.117 | כלומר |  |
| 6:17.118 | סבת |  |
| 6:17.119 | עונש |  |
| 6:17.120 | קבורתו | , |
| 6:17.121 | וסבת |  |
| 6:17.122 | עונש |  |
| 6:17.123 | קבורתו |  |
| 6:17.124 | שם | , |
| 6:17.125 | משנה |  |
| 6:17.126 | זו |  |
| 6:17.127 | שנתן |  |
| 6:17.128 | לישראל | . |
| 6:17.129 | והעוסקים |  |
| 6:17.130 | בה |  |
| 6:17.131 | ומורים |  |
| 6:17.132 | לעם |  |
| 6:17.133 | את |  |
| 6:17.134 | הדרך |  |
| 6:17.135 | ילכו |  |
| 6:17.136 | בה |  |
| 6:17.137 | קרא |  |
| 6:17.138 | להם |  |
| 6:17.139 | ערב |  |
| 6:17.140 | רב | , |
| 6:17.141 | לפי |  |
| 6:17.142 | שהמשנה |  |
| 6:17.143 | מרחקת |  |
| 6:17.144 | אמונתו |  |
| 6:17.145 | באלהים |  |
| 6:17.146 | רבים |  |
| 6:17.147 | שהוא |  |
| 6:17.148 | מרמזם |  |
| 6:17.149 | באותיות |  |
| 6:17.150 | יהוה | , |
| 6:17.151 | ולכן |  |
| 6:17.152 | קרא |  |
| 6:17.153 | למשנה |  |
| 6:17.154 | שפחה |  |
| 6:17.155 | דשלטא |  |
| 6:17.156 | על |  |
| 6:17.157 | מטרוניתא |  |
| 6:17.158 | שהיא |  |
| 6:17.159 | גורמת |  |
| 6:17.160 | להפריד |  |
| 6:17.161 | בין |  |
| 6:17.162 | מלכא |  |
| 6:17.163 | ומטרוניתא | , |
| 6:17.164 | באמרו |  |
| 6:17.165 | שם |  |
| 6:17.166 | ומלכא |  |
| 6:17.167 | ומטרוניתא |  |
| 6:17.168 | מתפרשן |  |
| 6:17.169 | מבעלה |  |
| 6:17.170 | בגין |  |
| 6:17.171 | כן |  |
| 6:17.172 | תחת |  |
| 6:17.173 | שלש |  |
| 6:17.174 | רגזה |  |
| 6:17.175 | ארץ | , |
| 6:17.176 | תחת |  |
| 6:17.177 | עבד |  |
| 6:17.178 | כי |  |
| 6:17.179 | ימלך | , |
| 6:17.180 | דא |  |
| 6:17.181 | איהו |  |
| 6:17.182 | עבדא |  |
| 6:17.183 | ידיעא |  |
| 6:17.184 | מטטרון |  |
| 6:17.185 | שנתן |  |
| 6:17.186 | המשנה |  |
| 6:17.187 | הזאת |  |
| 6:17.188 | למשה | ) |
| 6:17.189 | ושפחה |  |
| 6:17.190 | כי |  |
| 6:17.191 | תירש |  |
| 6:17.192 | גברתה |  |
| 6:17.193 | דא |  |
| 6:17.194 | משנה | , |
| 6:17.195 | ונבל |  |
| 6:17.196 | כי |  |
| 6:17.197 | ישבע |  |
| 6:17.198 | לחם |  |
| 6:17.199 | מלחמה |  |
| 6:17.200 | של |  |
| 6:17.201 | משנה | ) |
| 6:17.202 | דא |  |
| 6:17.203 | ערב |  |
| 6:17.204 | רב | . |
| 6:17.205 | העוסקים |  |
| 6:17.206 | במשנה | ) |
| 6:17.207 | עם |  |
| 6:17.208 | נבל |  |
| 6:17.209 | ולא |  |
| 6:17.210 | חכם | , |
| 6:17.211 | ואין |  |
| 6:17.212 | יחוד |  |
| 6:17.213 | לפרצופים |  |
| 6:17.214 | עד |  |
| 6:17.215 | דיתמחון |  |
| 6:17.216 | ערב |  |
| 6:17.217 | רב | , |
| 6:17.218 | העוסקים |  |
| 6:17.219 | בה |  |
| 6:17.220 | שהם |  |
| 6:17.221 | גורמים |  |
| 6:17.222 | פירוד |  |
| 6:17.223 | בין |  |
| 6:17.224 | המלכים |  |
| 6:17.225 | ופילגשיהם | . |
| 6:17.226 | ובפרשת |  |
| 6:17.227 | כי |  |
| 6:17.228 | תצא |  |
| 6:17.229 | הוא |  |
| 6:17.230 | משבח |  |
| 6:17.231 | את |  |
| 6:17.232 | תורתו |  |
| 6:17.233 | החדשה |  |
| 6:17.234 | שהוא |  |
| 6:17.235 | קורא |  |
| 6:17.236 | אותה |  |
| 6:17.237 | קבלה | , |
| 6:17.238 | והעוסקים |  |
| 6:17.239 | בה |  |
| 6:17.240 | קורא |  |
| 6:17.241 | להם |  |
| 6:17.242 | אדם | , |
| 6:17.243 | אבל |  |
| 6:17.244 | העוסקים |  |
| 6:17.245 | במשנה |  |
| 6:17.246 | ובתלמוד |  |
| 6:17.247 | הוא |  |
| 6:17.248 | קורא |  |
| 6:17.249 | להם |  |
| 6:17.250 | בשם |  |
| 6:17.251 | דגים |  |
| 6:17.252 | ועופות |  |
| 6:17.253 | השמים | , |
| 6:17.254 | והעוסקים |  |
| 6:17.255 | בקבלתו |  |
| 6:17.256 | למעלה |  |
| 6:17.257 | מהם |  |
| 6:17.258 | ושעליהם |  |
| 6:17.259 | נאמר |  |
| 6:17.260 | וירדו |  |
| 6:17.261 | בדגת |  |
| 6:17.262 | הים |  |
| 6:17.263 | ובעוף |  |
| 6:17.264 | השמים |  |
| 6:17.265 | דאינון |  |
| 6:17.266 | מארי |  |
| 6:17.267 | מתניתין | , |
| 6:17.268 | ומרבה |  |
| 6:17.269 | עוד |  |
| 6:17.270 | דברים |  |
| 6:17.271 | לגנות |  |
| 6:17.272 | התנאים |  |
| 6:17.273 | ואמוראים |  |
| 6:17.274 | ולהפחית |  |
| 6:17.275 | כבודם | , |
| 6:17.276 | ולשבח |  |
| 6:17.277 | העוסקים |  |
| 6:17.278 | בקבלתו |  |
| 6:17.279 | החדשה |  |
| 6:17.280 | בדברי |  |
| 6:17.281 | חלקות |  |
| 6:17.282 | היורדים |  |
| 6:17.283 | חדרי |  |
| 6:17.284 | בטן |  |
| 6:17.285 | קלי |  |
| 6:17.286 | הדעת |  |
| 6:17.287 | וחלושי |  |
| 6:17.288 | השכל | , |
| 6:17.289 | כמאמר |  |
| 6:17.290 | שלמה |  |
| 6:17.291 | ע״ה |  |
| 6:17.292 | דברי |  |
| 6:17.293 | נרגן |  |
| 6:17.294 | כמתלהמים |  |
| 6:17.295 | והם |  |
| 6:17.296 | ירדו |  |
| 6:17.297 | חדרי |  |
| 6:17.298 | בטן | , |
| 6:17.299 | כאשר |  |
| 6:17.300 | יראה |  |
| 6:17.301 | המעיין |  |
| 6:17.302 | המשכיל |  |
| 6:17.303 | על |  |
| 6:17.304 | דבר |  |
| 6:17.305 | אמת |  |
| 6:17.306 | הבקי |  |
| 6:17.307 | בכתובים |  |
| 6:17.308 | ובדברי |  |
| 6:17.309 | רז"ל |  |
| 6:17.310 | בתלמוד |  |
| 6:17.311 | ומדרשים |  |
| 6:17.312 | האמתיים | . |
| 6:17.313 | ושם |  |
| 6:17.314 | בדף |  |
| 6:17.315 | רע"ט |  |
| 6:17.316 | גילה |  |
| 6:17.317 | פנים |  |
| 6:17.318 | בחוצפא |  |
| 6:17.319 | יתירא |  |
| 6:17.320 | בשם |  |
| 6:17.321 | רעיא |  |
| 6:17.322 | מהימנא |  |
| 6:17.323 | לומר |  |
| 6:17.324 | אית |  |
| 6:17.325 | סלע |  |
| 6:17.326 | ואית |  |
| 6:17.327 | סלע | , |
| 6:17.328 | אית |  |
| 6:17.329 | אבן |  |
| 6:17.330 | ואית |  |
| 6:17.331 | אבן | , |
| 6:17.332 | אית |  |
| 6:17.333 | אבן |  |
| 6:17.334 | דשמא |  |
| 6:17.335 | דיהוה |  |
| 6:17.336 | היא |  |
| 6:17.337 | תורתו |  |
| 6:17.338 | החדשה |  |
| 6:17.339 | המלאה |  |
| 6:17.340 | פרצופים |  |
| 6:17.341 | וצורות |  |
| 6:17.342 | הנקראים |  |
| 6:17.343 | בשם |  |
| 6:17.344 | יהוה |  |
| 6:17.345 | ואדנות |  |
| 6:17.346 | ושאר |  |
| 6:17.347 | שמות |  |
| 6:17.348 | הקדש |  |
| 6:17.349 | נמחקים |  |
| 6:17.350 | ואין |  |
| 6:17.351 | נמחקים |  |
| 6:17.352 | ומרמזם |  |
| 6:17.353 | בשם |  |
| 6:17.354 | הוי"ה | ) |
| 6:17.355 | עלה |  |
| 6:17.356 | אמתר |  |
| 6:17.357 | ואבנא |  |
| 6:17.358 | די |  |
| 6:17.359 | מחת |  |
| 6:17.360 | לצלמא |  |
| 6:17.361 | הות |  |
| 6:17.362 | לטור |  |
| 6:17.363 | רב | , |
| 6:17.364 | ואית |  |
| 6:17.365 | אבן |  |
| 6:17.366 | דאיהי |  |
| 6:17.367 | אבן |  |
| 6:17.368 | משכית | , |
| 6:17.369 | דלית |  |
| 6:17.370 | תמן |  |
| 6:17.371 | נביעו |  |
| 6:17.372 | דחכמתא |  |
| 6:17.373 | ולא |  |
| 6:17.374 | דיבור |  |
| 6:17.375 | אלא |  |
| 6:17.376 | אבן |  |
| 6:17.377 | דאיהי |  |
| 6:17.378 | חסר |  |
| 6:17.379 | הפסוק |  |
| 6:17.380 | בכאן |  |
| 6:17.381 | ולא |  |
| 6:17.382 | חפץ |  |
| 6:17.383 | לפרסם |  |
| 6:17.384 | כוונתו |  |
| 6:17.385 | כל |  |
| 6:17.386 | כך |  |
| 6:17.387 | בבזוי |  |
| 6:17.388 | המשנה |  |
| 6:17.389 | והיה |  |
| 6:17.390 | רוצה |  |
| 6:17.391 | לומר |  |
| 6:17.392 | אבן |  |
| 6:17.393 | נגף |  |
| 6:17.394 | ואבן |  |
| 6:17.395 | מכשול | ). |
| 6:17.396 | ראה |  |
| 6:17.397 | והבן |  |
| 6:17.398 | בדבריו |  |
| 6:17.399 | איך |  |
| 6:17.400 | הוא |  |
| 6:17.401 | מחשיב |  |
| 6:17.402 | לתורתו |  |
| 6:17.403 | החדשה |  |
| 6:17.404 | וקורא |  |
| 6:17.405 | לה |  |
| 6:17.406 | אבנא |  |
| 6:17.407 | די |  |
| 6:17.408 | מחת |  |
| 6:17.409 | לצלמא | , |
| 6:17.410 | ובמלת |  |
| 6:17.411 | צלמא |  |
| 6:17.412 | כוונתו |  |
| 6:17.413 | על |  |
| 6:17.414 | המשנה |  |
| 6:17.415 | והתלמוד |  |
| 6:17.416 | שהוא |  |
| 6:17.417 | חושב |  |
| 6:17.418 | אותם |  |
| 6:17.419 | כצלמי |  |
| 6:17.420 | האלילים | , |
| 6:17.421 | ומבזה |  |
| 6:17.422 | אותה |  |
| 6:17.423 | וקורא |  |
| 6:17.424 | לה |  |
| 6:17.425 | אבן |  |
| 6:17.426 | משכית |  |
| 6:17.427 | דלית |  |
| 6:17.428 | בה |  |
| 6:17.429 | חכמתא | , |
| 6:17.430 | ובדבריו |  |
| 6:17.431 | אלה |  |
| 6:17.432 | הוא |  |
| 6:17.433 | רומז |  |
| 6:17.434 | לפרוש |  |
| 6:17.435 | מן |  |
| 6:17.436 | המשנה |  |
| 6:17.437 | כמאמר |  |
| 6:17.438 | הכתוב |  |
| 6:17.439 | ואבן |  |
| 6:17.440 | משכית |  |
| 6:17.441 | לא |  |
| 6:17.442 | תתנו |  |
| 6:17.443 | בארצכם |  |
| 6:17.444 | להשתחות |  |
| 6:17.445 | עליה |  |
| 6:17.446 | פי | ' |
| 6:17.447 | להשתחות |  |
| 6:17.448 | וליחד |  |
| 6:17.449 | השי"ת |  |
| 6:17.450 | על |  |
| 6:17.451 | פי |  |
| 6:17.452 | המשנה |  |
| 6:17.453 | המרחקת |  |
| 6:17.454 | כל |  |
| 6:17.455 | רבוי |  |
| 6:17.456 | באלהות | . |
| 6:17.457 | רק |  |
| 6:17.458 | על |  |
| 6:17.459 | פי |  |
| 6:17.460 | תורתו |  |
| 6:17.461 | החדשה |  |
| 6:17.462 | שהשי"ת |  |
| 6:17.463 | מחולק |  |
| 6:17.464 | לכמה |  |
| 6:17.465 | פרצופים |  |
| 6:17.466 | זכרים |  |
| 6:17.467 | ונקבות |  |
| 6:17.468 | וצריך |  |
| 6:17.469 | לקשרם |  |
| 6:17.470 | ולאחדם |  |
| 6:17.471 | ולזווגם |  |
| 6:17.472 | יחד |  |
| 6:17.473 | ר"ל | ). |
| 6:17.474 | ושוב |  |
| 6:17.475 | חזר |  |
| 6:17.476 | הפילסוף |  |
| 6:17.477 | הנז | ' |
| 6:17.478 | לדבריו |  |
| 6:17.479 | שבראש |  |
| 6:17.480 | עמוד |  |
| 6:17.481 | ב | ' |
| 6:17.482 | שקורא |  |
| 6:17.483 | לתורתו |  |
| 6:17.484 | סלע |  |
| 6:17.485 | הנובע |  |
| 6:17.486 | חכמה | , |
| 6:17.487 | ואת |  |
| 6:17.488 | המשנה |  |
| 6:17.489 | ואת |  |
| 6:17.490 | התלמוד |  |
| 6:17.491 | שקבל |  |
| 6:17.492 | משה |  |
| 6:17.493 | מסיני |  |
| 6:17.494 | הוא |  |
| 6:17.495 | מדמה |  |
| 6:17.496 | אותן |  |
| 6:17.497 | לקליפה |  |
| 6:17.498 | הקשה |  |
| 6:17.499 | של |  |
| 6:17.500 | אגוז |  |
| 6:17.501 | שכל |  |
| 6:17.502 | חכמי |  |
| 6:17.503 | הדורות |  |
| 6:17.504 | יגעים |  |
| 6:17.505 | בה |  |
| 6:17.506 | ולית |  |
| 6:17.507 | לון |  |
| 6:17.508 | יכולת |  |
| 6:17.509 | לאפקא |  |
| 6:17.510 | מינה |  |
| 6:17.511 | מיא |  |
| 6:17.512 | אלא |  |
| 6:17.513 | טיפין | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

