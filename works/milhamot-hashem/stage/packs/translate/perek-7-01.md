# Translation pack: milhamot-hashem perek-7-01 (Milchamot ha-Shem 7:1–7:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-01`.

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
- **הוי"ה / הוי'ה / הויה** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG). אדנות as a name → "Adonai" (note: "lordship").
  עשר הויות → "ten havayot" (note).
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
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
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


## Milchamot ha-Shem 7:1

Hebrew (Wikisource): פ) ואעתיק פה קצת העיקרים הנצרכים לעניינו מי"ג עיקרים שכתב הרמב"ם בפי' המשנה פרק חלק ומביאם הר' שבילי אמונה וז"ל. העיקר הראשון להאמין כי הבורא ית' נמצא מציאות גמורה. והוא עלת מציאות כל נצא. וממנו ימשך כח קיומם. ואילו נעביר על ליבנו מניעת מציאותו. מיד בטלה מציאות כל נמצא, ואין מציאות לשום נמצא. ואלו יעלה על ליבנו בטול מציאות הנמצאים כולם זולתו. לא בטלה מציאותו בשבילם ולא חסרה, כי הבורא ית' עשיר כי אין לו צורך במציאותו לזולתו. וכל בלעדיו מן הדיעות וגופי הגלגלים ומה שיש בתוכם, כולם יצטרכו במציאותם אליו ית'. הוא אינו צריך להם. וזה העיקר נרמז במלת אנכי ה' אלהיך. והעיקר הב' יחודו ית', והוא לדעת כי עלת הכל אחד לא כאחד המין הכולל. כמו מין האדם שכולל אישים הרבה. ולא כאחד המין הפרטי כמו שתאמר אדם אחר בפרט. ואתה כולל בו רמ"ח איברים. ולא כאחד המורכב אשר חלק לאחדים רבים. ולא כגוף הפשוט אשר הוא אחד במספרו ויקבל חילוק עד לאין תכלית. כי לכל אחד מאלו יקרא אחד על דרך העברה. מפני שהדברים ההם אשר יקבץ אותם השם ההוא הם שוין בענין אחד. אבל אינו אחד על דרך אמתי. אמנם האחד האמתי הוא אחדות הבורא ית'. אשר אין אחדות כמהו. בשום פנים. והוא אמרו ית' שמע ישראל ה' אלהינו ה' אחד עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 7:1.1 | פ | ) |
| 7:1.2 | ואעתיק |  |
| 7:1.3 | פה |  |
| 7:1.4 | קצת |  |
| 7:1.5 | העיקרים |  |
| 7:1.6 | הנצרכים |  |
| 7:1.7 | לעניינו |  |
| 7:1.8 | מי"ג |  |
| 7:1.9 | עיקרים |  |
| 7:1.10 | שכתב |  |
| 7:1.11 | הרמב"ם |  |
| 7:1.12 | בפי | ' |
| 7:1.13 | המשנה |  |
| 7:1.14 | פרק |  |
| 7:1.15 | חלק |  |
| 7:1.16 | ומביאם |  |
| 7:1.17 | הר | ' |
| 7:1.18 | שבילי |  |
| 7:1.19 | אמונה |  |
| 7:1.20 | וז"ל | . |
| 7:1.21 | העיקר |  |
| 7:1.22 | הראשון |  |
| 7:1.23 | להאמין |  |
| 7:1.24 | כי |  |
| 7:1.25 | הבורא |  |
| 7:1.26 | ית | ' |
| 7:1.27 | נמצא |  |
| 7:1.28 | מציאות |  |
| 7:1.29 | גמורה | . |
| 7:1.30 | והוא |  |
| 7:1.31 | עלת |  |
| 7:1.32 | מציאות |  |
| 7:1.33 | כל |  |
| 7:1.34 | נצא | . |
| 7:1.35 | וממנו |  |
| 7:1.36 | ימשך |  |
| 7:1.37 | כח |  |
| 7:1.38 | קיומם | . |
| 7:1.39 | ואילו |  |
| 7:1.40 | נעביר |  |
| 7:1.41 | על |  |
| 7:1.42 | ליבנו |  |
| 7:1.43 | מניעת |  |
| 7:1.44 | מציאותו | . |
| 7:1.45 | מיד |  |
| 7:1.46 | בטלה |  |
| 7:1.47 | מציאות |  |
| 7:1.48 | כל |  |
| 7:1.49 | נמצא | , |
| 7:1.50 | ואין |  |
| 7:1.51 | מציאות |  |
| 7:1.52 | לשום |  |
| 7:1.53 | נמצא | . |
| 7:1.54 | ואלו |  |
| 7:1.55 | יעלה |  |
| 7:1.56 | על |  |
| 7:1.57 | ליבנו |  |
| 7:1.58 | בטול |  |
| 7:1.59 | מציאות |  |
| 7:1.60 | הנמצאים |  |
| 7:1.61 | כולם |  |
| 7:1.62 | זולתו | . |
| 7:1.63 | לא |  |
| 7:1.64 | בטלה |  |
| 7:1.65 | מציאותו |  |
| 7:1.66 | בשבילם |  |
| 7:1.67 | ולא |  |
| 7:1.68 | חסרה | , |
| 7:1.69 | כי |  |
| 7:1.70 | הבורא |  |
| 7:1.71 | ית | ' |
| 7:1.72 | עשיר |  |
| 7:1.73 | כי |  |
| 7:1.74 | אין |  |
| 7:1.75 | לו |  |
| 7:1.76 | צורך |  |
| 7:1.77 | במציאותו |  |
| 7:1.78 | לזולתו | . |
| 7:1.79 | וכל |  |
| 7:1.80 | בלעדיו |  |
| 7:1.81 | מן |  |
| 7:1.82 | הדיעות |  |
| 7:1.83 | וגופי |  |
| 7:1.84 | הגלגלים |  |
| 7:1.85 | ומה |  |
| 7:1.86 | שיש |  |
| 7:1.87 | בתוכם | , |
| 7:1.88 | כולם |  |
| 7:1.89 | יצטרכו |  |
| 7:1.90 | במציאותם |  |
| 7:1.91 | אליו |  |
| 7:1.92 | ית | '. |
| 7:1.93 | הוא |  |
| 7:1.94 | אינו |  |
| 7:1.95 | צריך |  |
| 7:1.96 | להם | . |
| 7:1.97 | וזה |  |
| 7:1.98 | העיקר |  |
| 7:1.99 | נרמז |  |
| 7:1.100 | במלת |  |
| 7:1.101 | אנכי |  |
| 7:1.102 | ה | ' |
| 7:1.103 | אלהיך | . |
| 7:1.104 | והעיקר |  |
| 7:1.105 | הב | ' |
| 7:1.106 | יחודו |  |
| 7:1.107 | ית | ', |
| 7:1.108 | והוא |  |
| 7:1.109 | לדעת |  |
| 7:1.110 | כי |  |
| 7:1.111 | עלת |  |
| 7:1.112 | הכל |  |
| 7:1.113 | אחד |  |
| 7:1.114 | לא |  |
| 7:1.115 | כאחד |  |
| 7:1.116 | המין |  |
| 7:1.117 | הכולל | . |
| 7:1.118 | כמו |  |
| 7:1.119 | מין |  |
| 7:1.120 | האדם |  |
| 7:1.121 | שכולל |  |
| 7:1.122 | אישים |  |
| 7:1.123 | הרבה | . |
| 7:1.124 | ולא |  |
| 7:1.125 | כאחד |  |
| 7:1.126 | המין |  |
| 7:1.127 | הפרטי |  |
| 7:1.128 | כמו |  |
| 7:1.129 | שתאמר |  |
| 7:1.130 | אדם |  |
| 7:1.131 | אחר |  |
| 7:1.132 | בפרט | . |
| 7:1.133 | ואתה |  |
| 7:1.134 | כולל |  |
| 7:1.135 | בו |  |
| 7:1.136 | רמ"ח |  |
| 7:1.137 | איברים | . |
| 7:1.138 | ולא |  |
| 7:1.139 | כאחד |  |
| 7:1.140 | המורכב |  |
| 7:1.141 | אשר |  |
| 7:1.142 | חלק |  |
| 7:1.143 | לאחדים |  |
| 7:1.144 | רבים | . |
| 7:1.145 | ולא |  |
| 7:1.146 | כגוף |  |
| 7:1.147 | הפשוט |  |
| 7:1.148 | אשר |  |
| 7:1.149 | הוא |  |
| 7:1.150 | אחד |  |
| 7:1.151 | במספרו |  |
| 7:1.152 | ויקבל |  |
| 7:1.153 | חילוק |  |
| 7:1.154 | עד |  |
| 7:1.155 | לאין |  |
| 7:1.156 | תכלית | . |
| 7:1.157 | כי |  |
| 7:1.158 | לכל |  |
| 7:1.159 | אחד |  |
| 7:1.160 | מאלו |  |
| 7:1.161 | יקרא |  |
| 7:1.162 | אחד |  |
| 7:1.163 | על |  |
| 7:1.164 | דרך |  |
| 7:1.165 | העברה | . |
| 7:1.166 | מפני |  |
| 7:1.167 | שהדברים |  |
| 7:1.168 | ההם |  |
| 7:1.169 | אשר |  |
| 7:1.170 | יקבץ |  |
| 7:1.171 | אותם |  |
| 7:1.172 | השם |  |
| 7:1.173 | ההוא |  |
| 7:1.174 | הם |  |
| 7:1.175 | שוין |  |
| 7:1.176 | בענין |  |
| 7:1.177 | אחד | . |
| 7:1.178 | אבל |  |
| 7:1.179 | אינו |  |
| 7:1.180 | אחד |  |
| 7:1.181 | על |  |
| 7:1.182 | דרך |  |
| 7:1.183 | אמתי | . |
| 7:1.184 | אמנם |  |
| 7:1.185 | האחד |  |
| 7:1.186 | האמתי |  |
| 7:1.187 | הוא |  |
| 7:1.188 | אחדות |  |
| 7:1.189 | הבורא |  |
| 7:1.190 | ית | '. |
| 7:1.191 | אשר |  |
| 7:1.192 | אין |  |
| 7:1.193 | אחדות |  |
| 7:1.194 | כמהו | . |
| 7:1.195 | בשום |  |
| 7:1.196 | פנים | . |
| 7:1.197 | והוא |  |
| 7:1.198 | אמרו |  |
| 7:1.199 | ית | ' |
| 7:1.200 | שמע |  |
| 7:1.201 | ישראל |  |
| 7:1.202 | ה | ' |
| 7:1.203 | אלהינו |  |
| 7:1.204 | ה | ' |
| 7:1.205 | אחד |  |
| 7:1.206 | עכ"ל | . |


## Milchamot ha-Shem 7:2

Hebrew (Wikisource): ובחידושי גאונים בעין יע' פ"ג דנדרים כ' כי כמה מעלות טובותשיושלמו במצות המילה אל על עדת בני ישראל המכוונים במצותיה בחדרי קדש יסודתה היא המשלמת לאדם החלק המכבד שכאדם והוא השכל האמתי בענין שנאמר כי אם בזאת יתהלל המתהלל השכל וידוע אותי כי אני ה' וכו'. ביאורו שהשכל האמתי הוא ידיעת הבורא ודבקותו בו ית' וידיעת שמו הגדול הנרמז בזה המקרא יתהלל המתהלל השכל וידוע בראשי תיבותיהן שם בן ארבע אותיות. באמת יושלם על ידי המילה וסודה כמו שאמרו רז"ל מי יעלה לנו השמימה. ר"ה מילה וסופי תיבות הוי'ה. מה שיורה כי על ידי המילה אנו דבקים בו ית'. ולהוסיף ידיעה בשמו הקדוש ית' ויתעלה. לפי שהמילה היא תקון השכל האמתי. כי באמצעות מילת הבשר החיצונית תשלם מילת הלב הפנימית. אשר לזה הטעם בעצמו ולזאת הסבה בעינה חוייב שיכרתו עליה שלש עשרה בריתות בדוקא לא פחות ולא יותר, מה שיורה להיות מעשה הזה תועלת גדול הערך, כי יעלה ויבא ויגיע לזמן היותו בר עונשין שלש עשרה שנה שאז יקרא גדול, כי אז יבין יראת השם לאהבה אותו ולדבקה בו, ולהכיר בוראו לידע שהוא אחד תכלית היחוד, לא אחד במין, ולא אחד כגוף, שהוא מתחלק למחלוקות ולקצוות ולא אחד במספר, אבל אחד תכלית, היחוד שאין יחוד אחר כמותו בכל העולם, ובזה יעבוד בוראו, מאהבה נוסף על מה שעבדו מיראה, וזהו שנ' בחבלי אדם אמשכם בעבותות אהבה, ועל זה נקרא אברהם אבינו אחד שנ' אחד היה אברהם לטעם מצות המילה, עכ"ל לעניינינו.

| id | word | punctuation after |
|---|---|---|
| 7:2.1 | ובחידושי |  |
| 7:2.2 | גאונים |  |
| 7:2.3 | בעין |  |
| 7:2.4 | יע | ' |
| 7:2.5 | פ"ג |  |
| 7:2.6 | דנדרים |  |
| 7:2.7 | כ | ' |
| 7:2.8 | כי |  |
| 7:2.9 | כמה |  |
| 7:2.10 | מעלות |  |
| 7:2.11 | טובותשיושלמו |  |
| 7:2.12 | במצות |  |
| 7:2.13 | המילה |  |
| 7:2.14 | אל |  |
| 7:2.15 | על |  |
| 7:2.16 | עדת |  |
| 7:2.17 | בני |  |
| 7:2.18 | ישראל |  |
| 7:2.19 | המכוונים |  |
| 7:2.20 | במצותיה |  |
| 7:2.21 | בחדרי |  |
| 7:2.22 | קדש |  |
| 7:2.23 | יסודתה |  |
| 7:2.24 | היא |  |
| 7:2.25 | המשלמת |  |
| 7:2.26 | לאדם |  |
| 7:2.27 | החלק |  |
| 7:2.28 | המכבד |  |
| 7:2.29 | שכאדם |  |
| 7:2.30 | והוא |  |
| 7:2.31 | השכל |  |
| 7:2.32 | האמתי |  |
| 7:2.33 | בענין |  |
| 7:2.34 | שנאמר |  |
| 7:2.35 | כי |  |
| 7:2.36 | אם |  |
| 7:2.37 | בזאת |  |
| 7:2.38 | יתהלל |  |
| 7:2.39 | המתהלל |  |
| 7:2.40 | השכל |  |
| 7:2.41 | וידוע |  |
| 7:2.42 | אותי |  |
| 7:2.43 | כי |  |
| 7:2.44 | אני |  |
| 7:2.45 | ה | ' |
| 7:2.46 | וכו | '. |
| 7:2.47 | ביאורו |  |
| 7:2.48 | שהשכל |  |
| 7:2.49 | האמתי |  |
| 7:2.50 | הוא |  |
| 7:2.51 | ידיעת |  |
| 7:2.52 | הבורא |  |
| 7:2.53 | ודבקותו |  |
| 7:2.54 | בו |  |
| 7:2.55 | ית | ' |
| 7:2.56 | וידיעת |  |
| 7:2.57 | שמו |  |
| 7:2.58 | הגדול |  |
| 7:2.59 | הנרמז |  |
| 7:2.60 | בזה |  |
| 7:2.61 | המקרא |  |
| 7:2.62 | יתהלל |  |
| 7:2.63 | המתהלל |  |
| 7:2.64 | השכל |  |
| 7:2.65 | וידוע |  |
| 7:2.66 | בראשי |  |
| 7:2.67 | תיבותיהן |  |
| 7:2.68 | שם |  |
| 7:2.69 | בן |  |
| 7:2.70 | ארבע |  |
| 7:2.71 | אותיות | . |
| 7:2.72 | באמת |  |
| 7:2.73 | יושלם |  |
| 7:2.74 | על |  |
| 7:2.75 | ידי |  |
| 7:2.76 | המילה |  |
| 7:2.77 | וסודה |  |
| 7:2.78 | כמו |  |
| 7:2.79 | שאמרו |  |
| 7:2.80 | רז"ל |  |
| 7:2.81 | מי |  |
| 7:2.82 | יעלה |  |
| 7:2.83 | לנו |  |
| 7:2.84 | השמימה | . |
| 7:2.85 | ר"ה |  |
| 7:2.86 | מילה |  |
| 7:2.87 | וסופי |  |
| 7:2.88 | תיבות |  |
| 7:2.89 | הוי'ה | . |
| 7:2.90 | מה |  |
| 7:2.91 | שיורה |  |
| 7:2.92 | כי |  |
| 7:2.93 | על |  |
| 7:2.94 | ידי |  |
| 7:2.95 | המילה |  |
| 7:2.96 | אנו |  |
| 7:2.97 | דבקים |  |
| 7:2.98 | בו |  |
| 7:2.99 | ית | '. |
| 7:2.100 | ולהוסיף |  |
| 7:2.101 | ידיעה |  |
| 7:2.102 | בשמו |  |
| 7:2.103 | הקדוש |  |
| 7:2.104 | ית | ' |
| 7:2.105 | ויתעלה | . |
| 7:2.106 | לפי |  |
| 7:2.107 | שהמילה |  |
| 7:2.108 | היא |  |
| 7:2.109 | תקון |  |
| 7:2.110 | השכל |  |
| 7:2.111 | האמתי | . |
| 7:2.112 | כי |  |
| 7:2.113 | באמצעות |  |
| 7:2.114 | מילת |  |
| 7:2.115 | הבשר |  |
| 7:2.116 | החיצונית |  |
| 7:2.117 | תשלם |  |
| 7:2.118 | מילת |  |
| 7:2.119 | הלב |  |
| 7:2.120 | הפנימית | . |
| 7:2.121 | אשר |  |
| 7:2.122 | לזה |  |
| 7:2.123 | הטעם |  |
| 7:2.124 | בעצמו |  |
| 7:2.125 | ולזאת |  |
| 7:2.126 | הסבה |  |
| 7:2.127 | בעינה |  |
| 7:2.128 | חוייב |  |
| 7:2.129 | שיכרתו |  |
| 7:2.130 | עליה |  |
| 7:2.131 | שלש |  |
| 7:2.132 | עשרה |  |
| 7:2.133 | בריתות |  |
| 7:2.134 | בדוקא |  |
| 7:2.135 | לא |  |
| 7:2.136 | פחות |  |
| 7:2.137 | ולא |  |
| 7:2.138 | יותר | , |
| 7:2.139 | מה |  |
| 7:2.140 | שיורה |  |
| 7:2.141 | להיות |  |
| 7:2.142 | מעשה |  |
| 7:2.143 | הזה |  |
| 7:2.144 | תועלת |  |
| 7:2.145 | גדול |  |
| 7:2.146 | הערך | , |
| 7:2.147 | כי |  |
| 7:2.148 | יעלה |  |
| 7:2.149 | ויבא |  |
| 7:2.150 | ויגיע |  |
| 7:2.151 | לזמן |  |
| 7:2.152 | היותו |  |
| 7:2.153 | בר |  |
| 7:2.154 | עונשין |  |
| 7:2.155 | שלש |  |
| 7:2.156 | עשרה |  |
| 7:2.157 | שנה |  |
| 7:2.158 | שאז |  |
| 7:2.159 | יקרא |  |
| 7:2.160 | גדול | , |
| 7:2.161 | כי |  |
| 7:2.162 | אז |  |
| 7:2.163 | יבין |  |
| 7:2.164 | יראת |  |
| 7:2.165 | השם |  |
| 7:2.166 | לאהבה |  |
| 7:2.167 | אותו |  |
| 7:2.168 | ולדבקה |  |
| 7:2.169 | בו | , |
| 7:2.170 | ולהכיר |  |
| 7:2.171 | בוראו |  |
| 7:2.172 | לידע |  |
| 7:2.173 | שהוא |  |
| 7:2.174 | אחד |  |
| 7:2.175 | תכלית |  |
| 7:2.176 | היחוד | , |
| 7:2.177 | לא |  |
| 7:2.178 | אחד |  |
| 7:2.179 | במין | , |
| 7:2.180 | ולא |  |
| 7:2.181 | אחד |  |
| 7:2.182 | כגוף | , |
| 7:2.183 | שהוא |  |
| 7:2.184 | מתחלק |  |
| 7:2.185 | למחלוקות |  |
| 7:2.186 | ולקצוות |  |
| 7:2.187 | ולא |  |
| 7:2.188 | אחד |  |
| 7:2.189 | במספר | , |
| 7:2.190 | אבל |  |
| 7:2.191 | אחד |  |
| 7:2.192 | תכלית | , |
| 7:2.193 | היחוד |  |
| 7:2.194 | שאין |  |
| 7:2.195 | יחוד |  |
| 7:2.196 | אחר |  |
| 7:2.197 | כמותו |  |
| 7:2.198 | בכל |  |
| 7:2.199 | העולם | , |
| 7:2.200 | ובזה |  |
| 7:2.201 | יעבוד |  |
| 7:2.202 | בוראו | , |
| 7:2.203 | מאהבה |  |
| 7:2.204 | נוסף |  |
| 7:2.205 | על |  |
| 7:2.206 | מה |  |
| 7:2.207 | שעבדו |  |
| 7:2.208 | מיראה | , |
| 7:2.209 | וזהו |  |
| 7:2.210 | שנ | ' |
| 7:2.211 | בחבלי |  |
| 7:2.212 | אדם |  |
| 7:2.213 | אמשכם |  |
| 7:2.214 | בעבותות |  |
| 7:2.215 | אהבה | , |
| 7:2.216 | ועל |  |
| 7:2.217 | זה |  |
| 7:2.218 | נקרא |  |
| 7:2.219 | אברהם |  |
| 7:2.220 | אבינו |  |
| 7:2.221 | אחד |  |
| 7:2.222 | שנ | ' |
| 7:2.223 | אחד |  |
| 7:2.224 | היה |  |
| 7:2.225 | אברהם |  |
| 7:2.226 | לטעם |  |
| 7:2.227 | מצות |  |
| 7:2.228 | המילה | , |
| 7:2.229 | עכ"ל |  |
| 7:2.230 | לעניינינו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

